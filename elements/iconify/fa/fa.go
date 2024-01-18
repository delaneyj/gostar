package fa

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "4.7.0"
	hAttr          = "1em"
	viewbox        = "0 0 1536 1536"
	fill           = "currentColor"
)

type FaIcon struct {
	*SVGSVGElement
}

type FaIconFn func(children ...ElementRenderer) *FaIcon

var IconLookup = map[string]FaIconFn{
	"addressBook":                      AddressBook,
	"addressBookO":                     AddressBookO,
	"addressCard":                      AddressCard,
	"addressCardO":                     AddressCardO,
	"adjust":                           Adjust,
	"adn":                              Adn,
	"alignCenter":                      AlignCenter,
	"alignJustify":                     AlignJustify,
	"alignLeft":                        AlignLeft,
	"amazon":                           Amazon,
	"ambulance":                        Ambulance,
	"americanSignLanguageInterpreting": AmericanSignLanguageInterpreting,
	"anchor":                           Anchor,
	"android":                          Android,
	"angellist":                        Angellist,
	"angleDoubleLeft":                  AngleDoubleLeft,
	"angleDoubleUp":                    AngleDoubleUp,
	"angleDown":                        AngleDown,
	"angleLeft":                        AngleLeft,
	"angleUp":                          AngleUp,
	"apple":                            Apple,
	"archive":                          Archive,
	"areaChart":                        AreaChart,
	"arrowCircleLeft":                  ArrowCircleLeft,
	"arrowCircleOleft":                 ArrowCircleOleft,
	"arrowCircleOup":                   ArrowCircleOup,
	"arrowCircleUp":                    ArrowCircleUp,
	"arrowLeft":                        ArrowLeft,
	"arrowUp":                          ArrowUp,
	"arrows":                           Arrows,
	"arrowsAlt":                        ArrowsAlt,
	"arrowsH":                          ArrowsH,
	"arrowsV":                          ArrowsV,
	"assistiveListeningSystems":        AssistiveListeningSystems,
	"asterisk":                         Asterisk,
	"at":                               At,
	"audioDescription":                 AudioDescription,
	"automobile":                       Automobile,
	"backward":                         Backward,
	"balanceScale":                     BalanceScale,
	"ban":                              Ban,
	"bandcamp":                         Bandcamp,
	"bank":                             Bank,
	"barChart":                         BarChart,
	"barcode":                          Barcode,
	"bars":                             Bars,
	"bath":                             Bath,
	"battery":                          Battery,
	"batteryOne":                       BatteryOne,
	"batteryThree":                     BatteryThree,
	"batteryTwo":                       BatteryTwo,
	"batteryZero":                      BatteryZero,
	"bed":                              Bed,
	"beer":                             Beer,
	"behance":                          Behance,
	"behanceSquare":                    BehanceSquare,
	"bell":                             Bell,
	"bellO":                            BellO,
	"bellSlash":                        BellSlash,
	"bellSlashO":                       BellSlashO,
	"bicycle":                          Bicycle,
	"binoculars":                       Binoculars,
	"birthdayCake":                     BirthdayCake,
	"bitbucket":                        Bitbucket,
	"bitbucketSquare":                  BitbucketSquare,
	"bitcoin":                          Bitcoin,
	"blackTie":                         BlackTie,
	"blind":                            Blind,
	"bluetooth":                        Bluetooth,
	"bluetoothB":                       BluetoothB,
	"bold":                             Bold,
	"bolt":                             Bolt,
	"bomb":                             Bomb,
	"book":                             Book,
	"bookmark":                         Bookmark,
	"bookmarkO":                        BookmarkO,
	"braille":                          Braille,
	"briefcase":                        Briefcase,
	"bug":                              Bug,
	"building":                         Building,
	"buildingO":                        BuildingO,
	"bullhorn":                         Bullhorn,
	"bullseye":                         Bullseye,
	"bus":                              Bus,
	"buysellads":                       Buysellads,
	"cab":                              Cab,
	"calculator":                       Calculator,
	"calendar":                         Calendar,
	"calendarCheckO":                   CalendarCheckO,
	"calendarMinusO":                   CalendarMinusO,
	"calendarO":                        CalendarO,
	"calendarPlusO":                    CalendarPlusO,
	"calendarTimesO":                   CalendarTimesO,
	"camera":                           Camera,
	"cameraRetro":                      CameraRetro,
	"caretDown":                        CaretDown,
	"caretLeft":                        CaretLeft,
	"caretSquareOleft":                 CaretSquareOleft,
	"caretSquareOup":                   CaretSquareOup,
	"caretUp":                          CaretUp,
	"cartArrowDown":                    CartArrowDown,
	"cartPlus":                         CartPlus,
	"cc":                               Cc,
	"ccAmex":                           CcAmex,
	"ccDinersClub":                     CcDinersClub,
	"ccDiscover":                       CcDiscover,
	"ccJcb":                            CcJcb,
	"ccMastercard":                     CcMastercard,
	"ccPaypal":                         CcPaypal,
	"ccStripe":                         CcStripe,
	"ccVisa":                           CcVisa,
	"certificate":                      Certificate,
	"chain":                            Chain,
	"chainBroken":                      ChainBroken,
	"check":                            Check,
	"checkCircle":                      CheckCircle,
	"checkCircleO":                     CheckCircleO,
	"checkSquare":                      CheckSquare,
	"checkSquareO":                     CheckSquareO,
	"chevronCircleLeft":                ChevronCircleLeft,
	"chevronCircleUp":                  ChevronCircleUp,
	"chevronDown":                      ChevronDown,
	"chevronLeft":                      ChevronLeft,
	"chevronUp":                        ChevronUp,
	"child":                            Child,
	"chrome":                           Chrome,
	"circle":                           Circle,
	"circleO":                          CircleO,
	"circleOnotch":                     CircleOnotch,
	"circleThin":                       CircleThin,
	"clipboard":                        Clipboard,
	"clockO":                           ClockO,
	"clone":                            Clone,
	"close":                            Close,
	"cloud":                            Cloud,
	"cloudDownload":                    CloudDownload,
	"cloudUpload":                      CloudUpload,
	"cny":                              Cny,
	"code":                             Code,
	"codeFork":                         CodeFork,
	"codepen":                          Codepen,
	"codiepie":                         Codiepie,
	"coffee":                           Coffee,
	"cog":                              Cog,
	"cogs":                             Cogs,
	"columns":                          Columns,
	"comment":                          Comment,
	"commentO":                         CommentO,
	"commenting":                       Commenting,
	"commentingO":                      CommentingO,
	"comments":                         Comments,
	"commentsO":                        CommentsO,
	"compass":                          Compass,
	"compress":                         Compress,
	"connectdevelop":                   Connectdevelop,
	"contao":                           Contao,
	"copy":                             Copy,
	"copyright":                        Copyright,
	"creativeCommons":                  CreativeCommons,
	"creditCard":                       CreditCard,
	"creditCardAlt":                    CreditCardAlt,
	"crop":                             Crop,
	"crosshairs":                       Crosshairs,
	"cssThree":                         CssThree,
	"cube":                             Cube,
	"cubes":                            Cubes,
	"cut":                              Cut,
	"cutlery":                          Cutlery,
	"dashboard":                        Dashboard,
	"dashcube":                         Dashcube,
	"database":                         Database,
	"deaf":                             Deaf,
	"dedent":                           Dedent,
	"delicious":                        Delicious,
	"desktop":                          Desktop,
	"deviantart":                       Deviantart,
	"diamond":                          Diamond,
	"digg":                             Digg,
	"dollar":                           Dollar,
	"dotCircleO":                       DotCircleO,
	"download":                         Download,
	"dribbble":                         Dribbble,
	"driversLicense":                   DriversLicense,
	"driversLicenseO":                  DriversLicenseO,
	"dropbox":                          Dropbox,
	"drupal":                           Drupal,
	"edge":                             Edge,
	"edit":                             Edit,
	"eercast":                          Eercast,
	"eject":                            Eject,
	"ellipsisH":                        EllipsisH,
	"ellipsisV":                        EllipsisV,
	"empire":                           Empire,
	"envelope":                         Envelope,
	"envelopeO":                        EnvelopeO,
	"envelopeOpen":                     EnvelopeOpen,
	"envelopeOpenO":                    EnvelopeOpenO,
	"envelopeSquare":                   EnvelopeSquare,
	"envira":                           Envira,
	"eraser":                           Eraser,
	"etsy":                             Etsy,
	"eur":                              Eur,
	"exchange":                         Exchange,
	"exclamation":                      Exclamation,
	"exclamationCircle":                ExclamationCircle,
	"exclamationTriangle":              ExclamationTriangle,
	"expand":                           Expand,
	"expeditedssl":                     Expeditedssl,
	"externalLink":                     ExternalLink,
	"externalLinkSquare":               ExternalLinkSquare,
	"eye":                              Eye,
	"eyeSlash":                         EyeSlash,
	"eyedropper":                       Eyedropper,
	"fa":                               Fa,
	"facebook":                         Facebook,
	"facebookOfficial":                 FacebookOfficial,
	"facebookSquare":                   FacebookSquare,
	"fastBackward":                     FastBackward,
	"fax":                              Fax,
	"feed":                             Feed,
	"female":                           Female,
	"fighterJet":                       FighterJet,
	"file":                             File,
	"fileArchiveO":                     FileArchiveO,
	"fileAudioO":                       FileAudioO,
	"fileCodeO":                        FileCodeO,
	"fileExcelO":                       FileExcelO,
	"fileImageO":                       FileImageO,
	"fileMovieO":                       FileMovieO,
	"fileO":                            FileO,
	"filePdfO":                         FilePdfO,
	"filePowerpointO":                  FilePowerpointO,
	"fileText":                         FileText,
	"fileTextO":                        FileTextO,
	"fileWordO":                        FileWordO,
	"film":                             Film,
	"filter":                           Filter,
	"fire":                             Fire,
	"fireExtinguisher":                 FireExtinguisher,
	"firefox":                          Firefox,
	"firstOrder":                       FirstOrder,
	"fiveHundredPx":                    FiveHundredPx,
	"flag":                             Flag,
	"flagCheckered":                    FlagCheckered,
	"flagO":                            FlagO,
	"flask":                            Flask,
	"flickr":                           Flickr,
	"floppyO":                          FloppyO,
	"folder":                           Folder,
	"folderO":                          FolderO,
	"folderOpen":                       FolderOpen,
	"folderOpenO":                      FolderOpenO,
	"font":                             Font,
	"fonticons":                        Fonticons,
	"fortAwesome":                      FortAwesome,
	"forumbee":                         Forumbee,
	"foursquare":                       Foursquare,
	"freeCodeCamp":                     FreeCodeCamp,
	"frownO":                           FrownO,
	"futbolO":                          FutbolO,
	"gamepad":                          Gamepad,
	"gavel":                            Gavel,
	"gbp":                              Gbp,
	"genderless":                       Genderless,
	"getPocket":                        GetPocket,
	"gg":                               Gg,
	"ggCircle":                         GgCircle,
	"gift":                             Gift,
	"git":                              Git,
	"gitSquare":                        GitSquare,
	"github":                           Github,
	"githubAlt":                        GithubAlt,
	"githubSquare":                     GithubSquare,
	"gitlab":                           Gitlab,
	"gittip":                           Gittip,
	"glass":                            Glass,
	"glide":                            Glide,
	"glideG":                           GlideG,
	"globe":                            Globe,
	"google":                           Google,
	"googlePlus":                       GooglePlus,
	"googlePlusCircle":                 GooglePlusCircle,
	"googlePlusSquare":                 GooglePlusSquare,
	"googleWallet":                     GoogleWallet,
	"graduationCap":                    GraduationCap,
	"grav":                             Grav,
	"groupIcon":                        GroupIcon,
	"hsquare":                          Hsquare,
	"hackerNews":                       HackerNews,
	"handGrabO":                        HandGrabO,
	"handLizardO":                      HandLizardO,
	"handOleft":                        HandOleft,
	"handOup":                          HandOup,
	"handPaperO":                       HandPaperO,
	"handPeaceO":                       HandPeaceO,
	"handPointerO":                     HandPointerO,
	"handScissorsO":                    HandScissorsO,
	"handSpockO":                       HandSpockO,
	"handshakeO":                       HandshakeO,
	"hashtag":                          Hashtag,
	"hddO":                             HddO,
	"header":                           Header,
	"headphones":                       Headphones,
	"heart":                            Heart,
	"heartO":                           HeartO,
	"heartbeat":                        Heartbeat,
	"history":                          History,
	"home":                             Home,
	"hospitalO":                        HospitalO,
	"hourglass":                        Hourglass,
	"hourglassO":                       HourglassO,
	"hourglassOne":                     HourglassOne,
	"hourglassThree":                   HourglassThree,
	"hourglassTwo":                     HourglassTwo,
	"houzz":                            Houzz,
	"htmlFive":                         HtmlFive,
	"icursor":                          Icursor,
	"idBadge":                          IdBadge,
	"ils":                              Ils,
	"image":                            Image,
	"imdb":                             Imdb,
	"inbox":                            Inbox,
	"indent":                           Indent,
	"industry":                         Industry,
	"info":                             Info,
	"infoCircle":                       InfoCircle,
	"inr":                              Inr,
	"instagram":                        Instagram,
	"internetExplorer":                 InternetExplorer,
	"intersex":                         Intersex,
	"ioxhost":                          Ioxhost,
	"italic":                           Italic,
	"joomla":                           Joomla,
	"jsfiddle":                         Jsfiddle,
	"key":                              Key,
	"keyboardO":                        KeyboardO,
	"krw":                              Krw,
	"language":                         Language,
	"laptop":                           Laptop,
	"lastfm":                           Lastfm,
	"lastfmSquare":                     LastfmSquare,
	"leaf":                             Leaf,
	"leanpub":                          Leanpub,
	"lemonO":                           LemonO,
	"levelUp":                          LevelUp,
	"lifeBouy":                         LifeBouy,
	"lightbulbO":                       LightbulbO,
	"lineChart":                        LineChart,
	"linkedin":                         Linkedin,
	"linkedinSquare":                   LinkedinSquare,
	"linode":                           Linode,
	"linux":                            Linux,
	"list":                             List,
	"listAlt":                          ListAlt,
	"listOl":                           ListOl,
	"listUl":                           ListUl,
	"locationArrow":                    LocationArrow,
	"lock":                             Lock,
	"longArrowLeft":                    LongArrowLeft,
	"longArrowUp":                      LongArrowUp,
	"lowVision":                        LowVision,
	"magic":                            Magic,
	"magnet":                           Magnet,
	"mailForward":                      MailForward,
	"mailReply":                        MailReply,
	"mailReplyAll":                     MailReplyAll,
	"male":                             Male,
	"map":                              Map,
	"mapMarker":                        MapMarker,
	"mapO":                             MapO,
	"mapPin":                           MapPin,
	"mapSigns":                         MapSigns,
	"mars":                             Mars,
	"marsDouble":                       MarsDouble,
	"marsStroke":                       MarsStroke,
	"marsStrokeH":                      MarsStrokeH,
	"marsStrokeV":                      MarsStrokeV,
	"maxcdn":                           Maxcdn,
	"meanpath":                         Meanpath,
	"medium":                           Medium,
	"medkit":                           Medkit,
	"meetup":                           Meetup,
	"mehO":                             MehO,
	"mercury":                          Mercury,
	"microchip":                        Microchip,
	"microphone":                       Microphone,
	"microphoneSlash":                  MicrophoneSlash,
	"minus":                            Minus,
	"minusCircle":                      MinusCircle,
	"minusSquare":                      MinusSquare,
	"minusSquareO":                     MinusSquareO,
	"mixcloud":                         Mixcloud,
	"mobile":                           Mobile,
	"modx":                             Modx,
	"money":                            Money,
	"moonO":                            MoonO,
	"motorcycle":                       Motorcycle,
	"mousePointer":                     MousePointer,
	"music":                            Music,
	"neuter":                           Neuter,
	"newspaperO":                       NewspaperO,
	"objectGroup":                      ObjectGroup,
	"objectUngroup":                    ObjectUngroup,
	"odnoklassniki":                    Odnoklassniki,
	"odnoklassnikiSquare":              OdnoklassnikiSquare,
	"opencart":                         Opencart,
	"openid":                           Openid,
	"opera":                            Opera,
	"optinMonster":                     OptinMonster,
	"pagelines":                        Pagelines,
	"paintBrush":                       PaintBrush,
	"paperPlane":                       PaperPlane,
	"paperPlaneO":                      PaperPlaneO,
	"paperclip":                        Paperclip,
	"paragraph":                        Paragraph,
	"pause":                            Pause,
	"pauseCircle":                      PauseCircle,
	"pauseCircleO":                     PauseCircleO,
	"paw":                              Paw,
	"paypal":                           Paypal,
	"pencil":                           Pencil,
	"pencilSquare":                     PencilSquare,
	"percent":                          Percent,
	"phone":                            Phone,
	"phoneSquare":                      PhoneSquare,
	"pieChart":                         PieChart,
	"piedPiper":                        PiedPiper,
	"piedPiperAlt":                     PiedPiperAlt,
	"piedPiperPp":                      PiedPiperPp,
	"pinterest":                        Pinterest,
	"pinterestP":                       PinterestP,
	"pinterestSquare":                  PinterestSquare,
	"plane":                            Plane,
	"play":                             Play,
	"playCircle":                       PlayCircle,
	"playCircleO":                      PlayCircleO,
	"plug":                             Plug,
	"plus":                             Plus,
	"plusCircle":                       PlusCircle,
	"plusSquare":                       PlusSquare,
	"plusSquareO":                      PlusSquareO,
	"podcast":                          Podcast,
	"powerOff":                         PowerOff,
	"print":                            Print,
	"productHunt":                      ProductHunt,
	"puzzlePiece":                      PuzzlePiece,
	"qq":                               Qq,
	"qrcode":                           Qrcode,
	"question":                         Question,
	"questionCircle":                   QuestionCircle,
	"questionCircleO":                  QuestionCircleO,
	"quora":                            Quora,
	"quoteLeft":                        QuoteLeft,
	"quoteRight":                       QuoteRight,
	"ra":                               Ra,
	"random":                           Random,
	"ravelry":                          Ravelry,
	"recycle":                          Recycle,
	"reddit":                           Reddit,
	"redditAlien":                      RedditAlien,
	"redditSquare":                     RedditSquare,
	"refresh":                          Refresh,
	"registered":                       Registered,
	"renren":                           Renren,
	"repeat":                           Repeat,
	"retweet":                          Retweet,
	"road":                             Road,
	"rocket":                           Rocket,
	"rotateLeft":                       RotateLeft,
	"rouble":                           Rouble,
	"rssSquare":                        RssSquare,
	"safari":                           Safari,
	"scribd":                           Scribd,
	"search":                           Search,
	"searchMinus":                      SearchMinus,
	"searchPlus":                       SearchPlus,
	"sellsy":                           Sellsy,
	"server":                           Server,
	"shareAlt":                         ShareAlt,
	"shareAltSquare":                   ShareAltSquare,
	"shareSquare":                      ShareSquare,
	"shareSquareO":                     ShareSquareO,
	"shield":                           Shield,
	"ship":                             Ship,
	"shirtsinbulk":                     Shirtsinbulk,
	"shoppingBag":                      ShoppingBag,
	"shoppingBasket":                   ShoppingBasket,
	"shoppingCart":                     ShoppingCart,
	"shower":                           Shower,
	"signIn":                           SignIn,
	"signLanguage":                     SignLanguage,
	"signOut":                          SignOut,
	"signal":                           Signal,
	"simplybuilt":                      Simplybuilt,
	"sitemap":                          Sitemap,
	"skyatlas":                         Skyatlas,
	"skype":                            Skype,
	"slack":                            Slack,
	"sliders":                          Sliders,
	"slideshare":                       Slideshare,
	"smileO":                           SmileO,
	"snapchat":                         Snapchat,
	"snapchatGhost":                    SnapchatGhost,
	"snapchatSquare":                   SnapchatSquare,
	"snowflakeO":                       SnowflakeO,
	"sort":                             Sort,
	"sortAlphaAsc":                     SortAlphaAsc,
	"sortAlphaDesc":                    SortAlphaDesc,
	"sortAmountAsc":                    SortAmountAsc,
	"sortAmountDesc":                   SortAmountDesc,
	"sortAsc":                          SortAsc,
	"sortNumericAsc":                   SortNumericAsc,
	"sortNumericDesc":                  SortNumericDesc,
	"soundcloud":                       Soundcloud,
	"spaceShuttle":                     SpaceShuttle,
	"spinner":                          Spinner,
	"spoon":                            Spoon,
	"spotify":                          Spotify,
	"square":                           Square,
	"squareO":                          SquareO,
	"stackExchange":                    StackExchange,
	"stackOverflow":                    StackOverflow,
	"star":                             Star,
	"starHalf":                         StarHalf,
	"starHalfEmpty":                    StarHalfEmpty,
	"starO":                            StarO,
	"steam":                            Steam,
	"steamSquare":                      SteamSquare,
	"stepBackward":                     StepBackward,
	"stethoscope":                      Stethoscope,
	"stickyNote":                       StickyNote,
	"stickyNoteO":                      StickyNoteO,
	"stop":                             Stop,
	"stopCircle":                       StopCircle,
	"stopCircleO":                      StopCircleO,
	"streetView":                       StreetView,
	"strikethrough":                    Strikethrough,
	"stumbleupon":                      Stumbleupon,
	"stumbleuponCircle":                StumbleuponCircle,
	"subscript":                        Subscript,
	"subway":                           Subway,
	"suitcase":                         Suitcase,
	"sunO":                             SunO,
	"superpowers":                      Superpowers,
	"superscript":                      Superscript,
	"table":                            Table,
	"tablet":                           Tablet,
	"tag":                              Tag,
	"tags":                             Tags,
	"tasks":                            Tasks,
	"telegram":                         Telegram,
	"television":                       Television,
	"tencentWeibo":                     TencentWeibo,
	"terminal":                         Terminal,
	"textHeight":                       TextHeight,
	"textWidth":                        TextWidth,
	"th":                               Th,
	"thLarge":                          ThLarge,
	"thList":                           ThList,
	"themeisle":                        Themeisle,
	"thermometer":                      Thermometer,
	"thermometerOne":                   ThermometerOne,
	"thermometerThree":                 ThermometerThree,
	"thermometerTwo":                   ThermometerTwo,
	"thermometerZero":                  ThermometerZero,
	"thumbTack":                        ThumbTack,
	"thumbsDown":                       ThumbsDown,
	"thumbsOup":                        ThumbsOup,
	"thumbsUp":                         ThumbsUp,
	"ticket":                           Ticket,
	"timesCircle":                      TimesCircle,
	"timesCircleO":                     TimesCircleO,
	"timesRectangle":                   TimesRectangle,
	"timesRectangleO":                  TimesRectangleO,
	"tint":                             Tint,
	"toggleOff":                        ToggleOff,
	"toggleOn":                         ToggleOn,
	"trademark":                        Trademark,
	"train":                            Train,
	"transgenderAlt":                   TransgenderAlt,
	"trash":                            Trash,
	"trashO":                           TrashO,
	"tree":                             Tree,
	"trello":                           Trello,
	"tripadvisor":                      Tripadvisor,
	"trophy":                           Trophy,
	"truck":                            Truck,
	"try":                              Try,
	"tty":                              Tty,
	"tumblr":                           Tumblr,
	"tumblrSquare":                     TumblrSquare,
	"twitch":                           Twitch,
	"twitter":                          Twitter,
	"twitterSquare":                    TwitterSquare,
	"umbrella":                         Umbrella,
	"underline":                        Underline,
	"universalAccess":                  UniversalAccess,
	"unlock":                           Unlock,
	"unlockAlt":                        UnlockAlt,
	"upload":                           Upload,
	"usb":                              Usb,
	"user":                             User,
	"userCircle":                       UserCircle,
	"userCircleO":                      UserCircleO,
	"userMd":                           UserMd,
	"userO":                            UserO,
	"userPlus":                         UserPlus,
	"userSecret":                       UserSecret,
	"userTimes":                        UserTimes,
	"venus":                            Venus,
	"venusDouble":                      VenusDouble,
	"venusMars":                        VenusMars,
	"viacoin":                          Viacoin,
	"viadeo":                           Viadeo,
	"viadeoSquare":                     ViadeoSquare,
	"videoCamera":                      VideoCamera,
	"vimeo":                            Vimeo,
	"vimeoSquare":                      VimeoSquare,
	"vine":                             Vine,
	"vk":                               Vk,
	"volumeControlPhone":               VolumeControlPhone,
	"volumeDown":                       VolumeDown,
	"volumeOff":                        VolumeOff,
	"volumeUp":                         VolumeUp,
	"wechat":                           Wechat,
	"weibo":                            Weibo,
	"whatsapp":                         Whatsapp,
	"wheelchair":                       Wheelchair,
	"wheelchairAlt":                    WheelchairAlt,
	"wifi":                             Wifi,
	"wikipediaW":                       WikipediaW,
	"windowMaximize":                   WindowMaximize,
	"windowMinimize":                   WindowMinimize,
	"windowRestore":                    WindowRestore,
	"windows":                          Windows,
	"wordpress":                        Wordpress,
	"wpbeginner":                       Wpbeginner,
	"wpexplorer":                       Wpexplorer,
	"wpforms":                          Wpforms,
	"wrench":                           Wrench,
	"xing":                             Xing,
	"xingSquare":                       XingSquare,
	"ycombinator":                      Ycombinator,
	"yahoo":                            Yahoo,
	"yelp":                             Yelp,
	"yoast":                            Yoast,
	"youtube":                          Youtube,
	"youtubePlay":                      YoutubePlay,
	"youtubeSquare":                    YoutubeSquare,
}

func AddressBook(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1201 1238q0-57-5.5-107t-21-100.5t-39.5-86t-64-58t-91-22.5q-6 4-33.5 20.5T904 909t-40.5 20t-49 17t-46.5 5t-46.5-5t-49-17t-40.5-20t-42.5-24.5T556 864q-51 0-91 22.5t-64 58t-39.5 86t-21 100.5t-5.5 107q0 73 42 121.5t103 48.5h576q61 0 103-48.5t42-121.5m-173-594q0-108-76.5-184T768 384t-183.5 76T508 644q0 107 76.5 183T768 903t183.5-76t76.5-183m636 540v192q0 14-9 23t-23 9h-96v224q0 66-47 113t-113 47H160q-66 0-113-47T0 1632V160Q0 94 47 47T160 0h1216q66 0 113 47t47 113v224h96q14 0 23 9t9 23v192q0 14-9 23t-23 9h-96v128h96q14 0 23 9t9 23v192q0 14-9 23t-23 9h-96v128h96q14 0 23 9t9 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddressBookO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1028 644q0 107-76.5 183T768 903t-183.5-76T508 644q0-108 76.5-184T768 384t183.5 76t76.5 184m-48 220q46 0 82.5 17t60 47.5t39.5 67t24 81t11.5 82.5t3.5 79q0 67-39.5 118.5T1056 1408H480q-66 0-105.5-51.5T335 1238q0-48 4.5-93.5T358 1046t36.5-91.5t63-64.5t93.5-26h5q7 4 32 19.5t35.5 21t33 17t37 16t35 9T768 951t39.5-4.5t35-9t37-16t33-17t35.5-21t32-19.5m684-256q0 13-9.5 22.5T1632 640h-96v128h96q13 0 22.5 9.5t9.5 22.5v192q0 13-9.5 22.5t-22.5 9.5h-96v128h96q13 0 22.5 9.5t9.5 22.5v192q0 13-9.5 22.5t-22.5 9.5h-96v224q0 66-47 113t-113 47H160q-66 0-113-47T0 1632V160Q0 94 47 47T160 0h1216q66 0 113 47t47 113v224h96q13 0 22.5 9.5t9.5 22.5zm-256 1024V160q0-13-9.5-22.5T1376 128H160q-13 0-22.5 9.5T128 160v1472q0 13 9.5 22.5t22.5 9.5h1216q13 0 22.5-9.5t9.5-22.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddressCard(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1024 1003q0-64-9-117.5t-29.5-103t-60.5-78t-97-28.5q-6 4-30 18t-37.5 21.5T725 733t-43 14.5t-42 4.5t-42-4.5t-43-14.5t-35.5-17.5T482 694t-30-18q-57 0-97 28.5t-60.5 78t-29.5 103t-9 117.5t37 106.5t91 42.5h512q54 0 91-42.5t37-106.5M867 483q0-94-66.5-160.5T640 256t-160.5 66.5T413 483t66.5 160.5T640 710t160.5-66.5T867 483m925 509v-64q0-14-9-23t-23-9h-576q-14 0-23 9t-9 23v64q0 14 9 23t23 9h576q14 0 23-9t9-23m0-260v-56q0-15-10.5-25.5T1756 640h-568q-15 0-25.5 10.5T1152 676v56q0 15 10.5 25.5T1188 768h568q15 0 25.5-10.5T1792 732m0-252v-64q0-14-9-23t-23-9h-576q-14 0-23 9t-9 23v64q0 14 9 23t23 9h576q14 0 23-9t9-23m256-320v1216q0 66-47 113t-113 47h-352v-96q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v96H640v-96q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v96H160q-66 0-113-47T0 1376V160Q0 94 47 47T160 0h1728q66 0 113 47t47 113"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddressCardO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1024 1003q0 64-37 106.5t-91 42.5H384q-54 0-91-42.5T256 1003t9-117.5t29.5-103t60.5-78t97-28.5q6 4 30 18t37.5 21.5T555 733t43 14.5t42 4.5t42-4.5t43-14.5t35.5-17.5T798 694t30-18q57 0 97 28.5t60.5 78t29.5 103t9 117.5M867 483q0 94-66.5 160.5T640 710t-160.5-66.5T413 483t66.5-160.5T640 256t160.5 66.5T867 483m925 445v64q0 14-9 23t-23 9h-576q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h576q14 0 23 9t9 23m0-252v56q0 15-10.5 25.5T1756 768h-568q-15 0-25.5-10.5T1152 732v-56q0-15 10.5-25.5T1188 640h568q15 0 25.5 10.5T1792 676m0-260v64q0 14-9 23t-23 9h-576q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h576q14 0 23 9t9 23m128 960V160q0-13-9.5-22.5T1888 128H160q-13 0-22.5 9.5T128 160v1216q0 13 9.5 22.5t22.5 9.5h352v-96q0-14 9-23t23-9h64q14 0 23 9t9 23v96h768v-96q0-14 9-23t23-9h64q14 0 23 9t9 23v96h352q13 0 22.5-9.5t9.5-22.5m128-1216v1216q0 66-47 113t-113 47H160q-66 0-113-47T0 1376V160Q0 94 47 47T160 0h1728q66 0 113 47t47 113"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Adjust(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 1312V224q-148 0-273 73T297 495t-73 273t73 273t198 198t273 73m768-544q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Adn(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m768 494l201 306H567zm365 530h94L768 333l-459 691h94l104-160h522zm403-256q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenter(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 1216v128q0 26-19 45t-45 19H64q-26 0-45-19t-19-45v-128q0-26 19-45t45-19h1664q26 0 45 19t19 45m-384-384v128q0 26-19 45t-45 19H448q-26 0-45-19t-19-45V832q0-26 19-45t45-19h896q26 0 45 19t19 45m256-384v128q0 26-19 45t-45 19H192q-26 0-45-19t-19-45V448q0-26 19-45t45-19h1408q26 0 45 19t19 45M1280 64v128q0 26-19 45t-45 19H576q-26 0-45-19t-19-45V64q0-26 19-45t45-19h640q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignJustify(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 1216v128q0 26-19 45t-45 19H64q-26 0-45-19t-19-45v-128q0-26 19-45t45-19h1664q26 0 45 19t19 45m0-384v128q0 26-19 45t-45 19H64q-26 0-45-19T0 960V832q0-26 19-45t45-19h1664q26 0 45 19t19 45m0-384v128q0 26-19 45t-45 19H64q-26 0-45-19T0 576V448q0-26 19-45t45-19h1664q26 0 45 19t19 45m0-384v128q0 26-19 45t-45 19H64q-26 0-45-19T0 192V64q0-26 19-45T64 0h1664q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeft(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 1216v128q0 26-19 45t-45 19H64q-26 0-45-19t-19-45v-128q0-26 19-45t45-19h1664q26 0 45 19t19 45m-384-384v128q0 26-19 45t-45 19H64q-26 0-45-19T0 960V832q0-26 19-45t45-19h1280q26 0 45 19t19 45m256-384v128q0 26-19 45t-45 19H64q-26 0-45-19T0 576V448q0-26 19-45t45-19h1536q26 0 45 19t19 45M1280 64v128q0 26-19 45t-45 19H64q-26 0-45-19T0 192V64q0-26 19-45T64 0h1152q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Amazon(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1551 1476q15-6 26-3t11 17.5t-15 33.5q-13 16-44 43.5t-95.5 68t-141 74t-188 58T875 1792q-119 0-238-31t-209-76.5t-172.5-104t-132.5-105t-84-87.5q-8-9-10-16.5t1-12t8-7t11.5-2T61 1355q192 117 300 166q389 176 799 90q190-40 391-135m207-115q11 16 2.5 69.5T1732 1533q-34 83-85 124q-17 14-26 9t0-24q21-45 44.5-121.5t6.5-98.5q-5-7-15.5-11.5t-27-6t-29.5-2.5t-35 0t-31.5 2t-31 3t-22.5 2q-6 1-13 1.5t-11 1t-8.5 1t-7 .5h-10l-3-.5l-2-1.5l-1.5-3q-6-16 47-40t103-30q46-7 108-1t76 24m-394-443q0 31 13.5 64t32 58t37.5 46t33 32l13 11l-227 224q-40-37-79-75.5t-58-58.5l-19-20q-11-11-25-33q-38 59-97.5 102.5T860 1332t-140 23t-137.5-21t-117.5-65.5t-83-113T351 993q0-84 28-154t72-116.5t106.5-83t122.5-57T810 548t119.5-18.5t99.5-6.5V396q0-65-21-97q-34-53-121-53q-6 0-16.5 1T830 259t-56 29.5t-56 59.5t-48 96l-294-27q0-60 22-119t67-113t108-95t151.5-65.5T915 0q100 0 181 25t129.5 61.5t81 83t45 86T1364 329zm-672 21q0 86 70 133q66 44 139 22q84-25 114-123q14-45 14-101V708q-59 2-111 12t-106.5 33.5t-87 71T692 939"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ambulance(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M576 1280q0-53-37.5-90.5T448 1152t-90.5 37.5T320 1280t37.5 90.5T448 1408t90.5-37.5T576 1280M192 768h384V512H418q-14 2-22 9L201 716q-7 12-9 22zm1280 512q0-53-37.5-90.5T1344 1152t-90.5 37.5t-37.5 90.5t37.5 90.5t90.5 37.5t90.5-37.5t37.5-90.5m128-672V416q0-14-9-23t-23-9h-224V160q0-14-9-23t-23-9h-192q-14 0-23 9t-9 23v224H864q-14 0-23 9t-9 23v192q0 14 9 23t23 9h224v224q0 14 9 23t23 9h192q14 0 23-9t9-23V640h224q14 0 23-9t9-23m256-544v1152q0 26-19 45t-45 19h-192q0 106-75 181t-181 75t-181-75t-75-181H704q0 106-75 181t-181 75t-181-75t-75-181H64q-26 0-45-19t-19-45t19-45t45-19V736q0-26 13-58t32-51l198-198q19-19 51-32t58-13h160V64q0-26 19-45t45-19h1152q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AmericanSignLanguageInterpreting(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1032 960q-59-2-84-55q-17-34-48-53.5T832 832q-53 0-90.5 37.5T704 960q0 56 36 89l10 8q34 31 82 31q37 0 68-19.5t48-53.5q25-53 84-55m568-128q0-56-36-89l-10-8q-34-31-82-31q-37 0-68 19.5t-48 53.5q-25 53-84 55q59 2 84 55q17 34 48 53.5t68 19.5q53 0 90.5-37.5T1600 832m-426-221q-17 35-55 48t-73-4q-62-31-134-31q-51 0-99 17q3 0 9.5-.5t9.5-.5q92 0 170.5 50T1121 823q17 36 3.5 73.5T1075 951q-18 9-39 9q21 0 39 9q36 17 49.5 54.5t-3.5 73.5q-40 83-118.5 133T832 1280h-6q-16-2-44-4l-290-27l-239 120q-14 7-29 7q-40 0-57-35L7 1021q-11-23-4-47.5T32 936l209-119l148-267q17-155 91.5-291.5T676 22Q707-3 746.5.5T811 35t21.5 70t-34.5 65q-70 59-117 128q123-84 267-101q40-5 71.5 19t35.5 64q5 40-19 71.5T972 387q-84 10-159 55q46-10 99-10q115 0 218 50q36 18 49 55.5t-5 73.5m963-160l160 320q11 23 4 47.5t-29 37.5l-209 119l-148 267q-17 155-91.5 291.5T1628 1770q-26 22-61 22q-45 0-74-35q-25-31-21.5-70t34.5-65q70-59 117-128q-123 84-267 101q-4 1-12 1q-36 0-63.5-24t-31.5-60q-5-40 19-71.5t64-35.5q84-10 159-55q-46 10-99 10q-115 0-218-50q-36-18-49-55.5t5-73.5q17-35 55-48t73 4q62 31 134 31q51 0 99-17q-3 0-9.5.5t-9.5.5q-92 0-170.5-50T1183 969q-17-36-3.5-73.5T1229 841q18-9 39-9q-21 0-39-9q-36-17-49.5-54.5t3.5-73.5q40-83 118.5-133t170.5-50h7q14 2 42 4l291 27l239-120q14-7 29-7q40 0 57 35"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Anchor(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M960 256q0-26-19-45t-45-19t-45 19t-19 45t19 45t45 19t45-19t19-45m832 928v352q0 22-20 30q-8 2-12 2q-12 0-23-9l-93-93q-119 143-318.5 226.5T896 1776t-429.5-83.5T148 1466l-93 93q-9 9-23 9q-4 0-12-2q-20-8-20-30v-352q0-14 9-23t23-9h352q22 0 30 20q8 19-7 35l-100 100q67 91 189.5 153.5T768 1543V896H576q-26 0-45-19t-19-45V704q0-26 19-45t45-19h192V477q-58-34-93-92.5T640 256q0-106 75-181T896 0t181 75t75 181q0 70-35 128.5t-93 92.5v163h192q26 0 45 19t19 45v128q0 26-19 45t-45 19h-192v647q149-20 271.5-82.5T1485 1307l-100-100q-15-16-7-35q8-20 30-20h352q14 0 23 9t9 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Android(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M493 355q16 0 27.5-11.5T532 316t-11.5-27.5T493 277t-27 11.5t-11 27.5t11 27.5t27 11.5m422 0q16 0 27-11.5t11-27.5t-11-27.5t-27-11.5t-27.5 11.5T876 316t11.5 27.5T915 355M103 539q42 0 72 30t30 72v430q0 43-29.5 73t-72.5 30t-73-30t-30-73V641q0-42 30-72t73-30m1060 19v666q0 46-32 78t-77 32h-75v227q0 43-30 73t-73 30t-73-30t-30-73v-227H635v227q0 43-30 73t-73 30q-42 0-72-30t-30-73l-1-227h-74q-46 0-78-32t-32-78V558zM931 153q107 55 171 153.5t64 215.5H241q0-117 64-215.5T477 153L406 22q-7-13 5-20q13-6 20 6l72 132q95-42 201-42t201 42L977 8q7-12 20-6q12 7 5 20zm477 488v430q0 43-30 73t-73 30q-42 0-72-30t-30-73V641q0-43 30-72.5t72-29.5q43 0 73 29.5t30 72.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Angellist(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M953 378L839 706l117 21q165-451 165-518q0-56-38-56q-57 0-130 225m-299 687l33 88q37-42 71-67l-33-5.5l-38.5-7zM362 169q0 98 159 521q17-10 49-10q15 0 75 5L524 334q-75-220-123-220q-19 0-29 17.5T362 169m-79 759q0 36 51.5 119T452 1200t100 70q14 0 25.5-13t11.5-27q0-24-32-102q-13-32-32-72t-47.5-89t-61.5-81t-62-32q-20 0-45.5 27T283 928m-158 335q0 41 25 104q59 145 183.5 227t281.5 82q227 0 382-170q152-169 152-427q0-43-1-67t-11.5-62t-30.5-56q-56-49-211.5-75.5T624 792q-37 0-49 11q-12 5-12 35q0 34 21.5 60t55.5 40t77.5 23.5T805 973t85 4t70 0h23q24 0 40 19q15 19 19 55q-28 28-96 54q-61 22-93 46q-64 46-108.5 114T700 1402q0 31 18.5 88.5T737 1578l-3 12q-4 12-4 14q-137-10-146-216q-8 2-41 2q2 7 2 21q0 53-40.5 89.5T410 1537q-82 0-166.5-78T159 1300q0-34 33-67q52 64 60 76q77 104 133 104q12 0 26.5-8.5T426 1384q0-34-87.5-145T222 1128q-43 0-70 44.5t-27 90.5m-114 9q0-101 42.5-163t136.5-88q-28-74-28-104q0-62 61-123t122-61q29 0 70 15q-163-462-163-567q0-80 41-130.5T412 0q131 0 325 581q6 17 8 23q6-16 29-79.5T817.5 406t54-127.5t64.5-123t70.5-86.5t76.5-36q71 0 112 49t41 122q0 108-159 550q61 15 100.5 46t58.5 78t26 93.5t7 110.5q0 150-47 280t-132 225t-211 150t-278 55q-111 0-223-42q-149-57-258-191.5T11 1272"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDoubleLeft(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M595 1120q0 13-10 23l-50 50q-10 10-23 10t-23-10L23 727q-10-10-10-23t10-23l466-466q10-10 23-10t23 10l50 50q10 10 10 23t-10 23L192 704l393 393q10 10 10 23m384 0q0 13-10 23l-50 50q-10 10-23 10t-23-10L407 727q-10-10-10-23t10-23l466-466q10-10 23-10t23 10l50 50q10 10 10 23t-10 23L576 704l393 393q10 10 10 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDoubleUp(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1011 1056q0 13-10 23l-50 50q-10 10-23 10t-23-10L512 736l-393 393q-10 10-23 10t-23-10l-50-50q-10-10-10-23t10-23l466-466q10-10 23-10t23 10l466 466q10 10 10 23m0-384q0 13-10 23l-50 50q-10 10-23 10t-23-10L512 352L119 745q-10 10-23 10t-23-10l-50-50q-10-10-10-23t10-23l466-466q10-10 23-10t23 10l466 466q10 10 10 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDown(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1011 480q0 13-10 23L535 969q-10 10-23 10t-23-10L23 503q-10-10-10-23t10-23l50-50q10-10 23-10t23 10l393 393l393-393q10-10 23-10t23 10l50 50q10 10 10 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleLeft(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M595 288q0 13-10 23L192 704l393 393q10 10 10 23t-10 23l-50 50q-10 10-23 10t-23-10L23 727q-10-10-10-23t10-23l466-466q10-10 23-10t23 10l50 50q10 10 10 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleUp(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1011 928q0 13-10 23l-50 50q-10 10-23 10t-23-10L512 608l-393 393q-10 10-23 10t-23-10l-50-50q-10-10-10-23t10-23l466-466q10-10 23-10t23 10l466 466q10 10 10 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Apple(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1393 1215q-39 125-123 250q-129 196-257 196q-49 0-140-32q-86-32-151-32q-61 0-142 33q-81 34-132 34q-152 0-301-259Q0 1144 0 902q0-228 113-374q113-144 284-144q72 0 177 30q104 30 138 30q45 0 143-34q102-34 173-34q119 0 213 65q52 36 104 100q-79 67-114 118q-65 94-65 207q0 124 69 223t158 126M1017 42q0 61-29 136q-30 75-93 138q-54 54-108 72q-37 11-104 17q3-149 78-257Q835 41 1011 0q1 3 2.5 11t2.5 11q0 4 .5 10t.5 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Archive(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1024 704q0-26-19-45t-45-19H704q-26 0-45 19t-19 45t19 45t45 19h256q26 0 45-19t19-45m576-192v960q0 26-19 45t-45 19H128q-26 0-45-19t-19-45V512q0-26 19-45t45-19h1408q26 0 45 19t19 45m64-448v256q0 26-19 45t-45 19H64q-26 0-45-19T0 320V64q0-26 19-45T64 0h1536q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AreaChart(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2048 1408v128H0V0h128v1408zM1664 384l256 896H256V704l448-576l576 576z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleLeft(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1280 832V704q0-26-19-45t-45-19H714l189-189q19-19 19-45t-19-45l-91-91q-18-18-45-18t-45 18L360 632l-91 91q-18 18-18 45t18 45l91 91l362 362q18 18 45 18t45-18l91-91q18-18 18-45t-18-45L714 896h502q26 0 45-19t19-45m256-64q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleOleft(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1152 672v192q0 13-9.5 22.5T1120 896H768v192q0 14-9 23t-23 9q-12 0-24-10L393 791q-9-9-9-23t9-23l320-320q9-9 23-9q13 0 22.5 9.5T768 448v192h352q13 0 22.5 9.5t9.5 22.5m160 96q0-148-73-273t-198-198t-273-73t-273 73t-198 198t-73 273t73 273t198 198t273 73t273-73t198-198t73-273m224 0q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleOup(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1118 748q-8 20-30 20H896v352q0 14-9 23t-23 9H672q-14 0-23-9t-9-23V768H448q-14 0-23-9t-9-23q0-12 10-24l319-319q11-9 23-9t23 9l320 320q15 16 7 35M768 224q-148 0-273 73T297 495t-73 273t73 273t198 198t273 73t273-73t198-198t73-273t-73-273t-198-198t-273-73m768 544q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleUp(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1284 767q0-27-18-45L904 360l-91-91q-18-18-45-18t-45 18l-91 91l-362 362q-18 18-18 45t18 45l91 91q18 18 45 18t45-18l189-189v502q0 26 19 45t45 19h128q26 0 45-19t19-45V714l189 189q19 19 45 19t45-19l91-91q18-18 18-45m252 1q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1472 736v128q0 53-32.5 90.5T1355 992H651l293 294q38 36 38 90t-38 90l-75 76q-37 37-90 37q-52 0-91-37L37 890Q0 853 0 800q0-52 37-91L688 59q38-38 91-38q52 0 90 38l75 74q38 38 38 91t-38 91L651 608h704q52 0 84.5 37.5T1472 736"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUp(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1579 779q0 51-37 90l-75 75q-38 38-91 38q-54 0-90-38L992 651v704q0 52-37.5 84.5T864 1472H736q-53 0-90.5-32.5T608 1355V651L314 944q-36 38-90 38t-90-38l-75-75q-38-38-38-90q0-53 38-91L710 37q35-37 90-37q54 0 91 37l651 651q37 39 37 91"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arrows(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 896q0 26-19 45l-256 256q-19 19-45 19t-45-19t-19-45v-128h-384v384h128q26 0 45 19t19 45t-19 45l-256 256q-19 19-45 19t-45-19l-256-256q-19-19-19-45t19-45t45-19h128v-384H384v128q0 26-19 45t-45 19t-45-19L19 941Q0 922 0 896t19-45l256-256q19-19 45-19t45 19t19 45v128h384V384H640q-26 0-45-19t-19-45t19-45L851 19q19-19 45-19t45 19l256 256q19 19 19 45t-19 45t-45 19h-128v384h384V640q0-26 19-45t45-19t45 19l256 256q19 19 19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsAlt(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1283 413L928 768l355 355l144-144q29-31 70-14q39 17 39 59v448q0 26-19 45t-45 19h-448q-42 0-59-40q-17-39 14-69l144-144l-355-355l-355 355l144 144q31 30 14 69q-17 40-59 40H64q-26 0-45-19t-19-45v-448q0-42 40-59q39-17 69 14l144 144l355-355l-355-355l-144 144q-19 19-45 19q-12 0-24-5q-40-17-40-59V64q0-26 19-45T64 0h448q42 0 59 40q17 39-14 69L413 253l355 355l355-355l-144-144q-31-30-14-69q17-40 59-40h448q26 0 45 19t19 45v448q0 42-39 59q-13 5-25 5q-26 0-45-19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsH(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 640q0 26-19 45l-256 256q-19 19-45 19t-45-19t-19-45V768H384v128q0 26-19 45t-45 19t-45-19L19 685Q0 666 0 640t19-45l256-256q19-19 45-19t45 19t19 45v128h1024V384q0-26 19-45t45-19t45 19l256 256q19 19 19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsV(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 320q0 26-19 45t-45 19H448v1024h128q26 0 45 19t19 45t-19 45l-256 256q-19 19-45 19t-45-19L19 1517q-19-19-19-45t19-45t45-19h128V384H64q-26 0-45-19T0 320t19-45L275 19q19-19 45-19t45 19l256 256q19 19 19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AssistiveListeningSystems(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 1728q0 26-19 45t-45 19t-45-19t-19-45t19-45t45-19t45 19t19 45m192-192q0 26-19 45t-45 19t-45-19t-19-45t19-45t45-19t45 19t19 45m45-365l256 256l-90 90l-256-256zm339-19q0 26-19 45t-45 19t-45-19t-19-45t19-45t45-19t45 19t19 45m707-320q0 59-11.5 108.5T1362 1034t-44 67.5t-53 64.5q-31 35-45.5 54t-33.5 50t-26.5 64t-7.5 74q0 159-112.5 271.5T768 1792q-26 0-45-19t-19-45t19-45t45-19q106 0 181-75t75-181q0-57 11.5-105.5t37-91t43.5-66.5t52-63q40-46 59.5-72t37.5-74.5t18-103.5q0-185-131.5-316.5T835 384T518.5 515.5T387 832q0 26-19 45t-45 19t-45-19t-19-45q0-117 45.5-223.5t123-184t184-123T835 256t223.5 45.5t184 123t123 184T1411 832M896 960q0 26-19 45t-45 19t-45-19t-19-45t19-45t45-19t45 19t19 45m288-128q0 26-19 45t-45 19t-45-19t-19-45q0-93-65.5-158.5T832 608q-92 0-158 65.5T608 832q0 26-19 45t-45 19t-45-19t-19-45q0-146 103-249t249-103t249 103t103 249m394-289q10 25-1 49t-36 34q-9 4-23 4q-19 0-35.5-11t-23.5-30q-68-178-224-295q-21-16-25-42t12-47q17-21 43-25t47 12q183 137 266 351m210-81q9 25-1.5 49t-35.5 34q-11 4-23 4q-44 0-60-41q-92-238-297-393q-22-16-25.5-42t12.5-47q16-22 42-25.5t47 12.5q235 175 341 449"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Asterisk(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1386 922q46 26 59.5 77.5T1433 1097l-64 110q-26 46-77.5 59.5T1194 1254l-266-153v307q0 52-38 90t-90 38H672q-52 0-90-38t-38-90v-307l-266 153q-46 26-97.5 12.5T103 1207l-64-110q-26-46-12.5-97.5T86 922l266-154L86 614q-46-26-59.5-77.5T39 439l64-110q26-46 77.5-59.5T278 282l266 153V128q0-52 38-90t90-38h128q52 0 90 38t38 90v307l266-153q46-26 97.5-12.5T1369 329l64 110q26 46 12.5 97.5T1386 614l-266 154z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func At(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M972 647q0-108-53.5-169T771 417q-63 0-124 30.5T537 532t-79.5 137T427 849q0 112 53.5 173t150.5 61q96 0 176-66.5t122.5-166T972 647m564 121q0 111-37 197t-98.5 135t-131.5 74.5t-145 27.5q-6 0-15.5.5t-16.5.5q-95 0-142-53q-28-33-33-83q-52 66-131.5 110T612 1221q-161 0-249.5-95.5T274 856q0-157 66-290t179-210.5T765 278q87 0 155 35.5t106 99.5l2-19l11-56q1-6 5.5-12t9.5-6h118q5 0 13 11q5 5 3 16l-120 614q-5 24-5 48q0 39 12.5 52t44.5 13q28-1 57-5.5t73-24t77-50t57-89.5t24-137q0-292-174-466T768 128q-130 0-248.5 51t-204 136.5t-136.5 204T128 768t51 248.5t136.5 204t204 136.5t248.5 51q228 0 405-144q11-9 24-8t21 12l41 49q8 12 7 24q-2 13-12 22q-102 83-227.5 128T768 1536q-156 0-298-61t-245-164t-164-245T0 768t61-298t164-245T470 61T768 0q344 0 556 212t212 556"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AudioDescription(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M504 738h171l-1-265zm1026-99q0-87-50.5-140T1333 446h-54v388h52q91 0 145-57t54-138M956 262l1 756q0 14-9.5 24t-23.5 10H708q-14 0-23.5-10t-9.5-24v-62H384l-55 81q-10 15-28 15H34q-21 0-30.5-18T7 999l556-757q9-14 27-14h332q14 0 24 10t10 24m827 377q0 193-125.5 303T1333 1052h-270q-14 0-24-10t-10-24V262q0-14 10-24t24-10h268q200 0 326 109t126 302m156 1q0 11-.5 29t-8 71.5t-21.5 102t-44.5 108T1791 1053h-51q38-45 66.5-104.5t41.5-112t21-98t9-72.5l1-27q0-8-.5-22.5t-7.5-60t-20-91.5t-41-111.5t-66-124.5h43q41 47 72 107t45.5 111.5t23 96T1938 614zm184 0q0 11-.5 29t-8 71.5t-21.5 102t-45 108t-74 102.5h-51q38-45 66.5-104.5t41.5-112t21-98t9-72.5l1-27q0-8-.5-22.5t-7.5-60t-19.5-91.5t-40.5-111.5t-66-124.5h43q41 47 72 107t45.5 111.5t23 96T2122 614zm181 0q0 11-.5 29t-8 71.5t-21.5 102t-44.5 108T2156 1053h-51q38-45 66-104.5t41-112t21-98t9-72.5l1-27q0-8-.5-22.5t-7.5-60t-19.5-91.5t-40.5-111.5t-66-124.5h43q41 47 72 107t45.5 111.5t23 96t9.5 70.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Automobile(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M480 960q0-66-47-113t-113-47t-113 47t-47 113t47 113t113 47t113-47t47-113m36-320h1016l-89-357q-2-8-14-17.5t-21-9.5H640q-9 0-21 9.5T605 283zm1372 320q0-66-47-113t-113-47t-113 47t-47 113t47 113t113 47t113-47t47-113m160-96v384q0 14-9 23t-23 9h-96v128q0 80-56 136t-136 56t-136-56t-56-136v-128H512v128q0 80-56 136t-136 56t-136-56t-56-136v-128H32q-14 0-23-9t-9-23V864q0-93 65.5-158.5T224 640h28l105-419q23-94 104-157.5T640 0h768q98 0 179 63.5T1691 221l105 419h28q93 0 158.5 65.5T2048 864"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Backward(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1523 13q19-19 32-13t13 32v1472q0 26-13 32t-32-13L813 813q-9-9-13-19v710q0 26-13 32t-32-13L45 813q-19-19-19-45t19-45L755 13q19-19 32-13t13 32v710q4-10 13-19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BalanceScale(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1728 448l-384 704h768zm-1280 0L64 1152h768zm821-192q-14 40-45.5 71.5T1152 373v1291h608q14 0 23 9t9 23v64q0 14-9 23t-23 9H416q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h608V373q-40-14-71.5-45.5T907 256H416q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h491q21-57 70-92.5T1088 0t111 35.5t70 92.5h491q14 0 23 9t9 23v64q0 14-9 23t-23 9zm-181 16q33 0 56.5-23.5T1168 192t-23.5-56.5T1088 112t-56.5 23.5T1008 192t23.5 56.5T1088 272m1088 880q0 73-46.5 131t-117.5 91t-144.5 49.5T1728 1440t-139.5-16.5T1444 1374t-117.5-91t-46.5-131q0-11 35-81t92-174.5T1514 701t102-184t56-100q18-33 56-33t56 33q4 7 56 100t102 184t107 195.5t92 174.5t35 81m-1280 0q0 73-46.5 131T732 1374t-144.5 49.5T448 1440t-139.5-16.5T164 1374t-117.5-91T0 1152q0-11 35-81t92-174.5T234 701t102-184t56-100q18-33 56-33t56 33q4 7 56 100t102 184t107 195.5t92 174.5t35 81"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ban(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1312 797q0-161-87-295l-754 753q137 89 297 89q111 0 211.5-43.5T1153 1184t116-174.5t43-212.5m-999 299l755-754q-135-91-300-91q-148 0-273 73T297 523t-73 274q0 162 89 299m1223-299q0 157-61 300t-163.5 246t-245 164t-298.5 61t-298.5-61t-245-164T61 1097T0 797t61-299.5T224.5 252t245-164T768 27t298.5 61t245 164T1475 497.5t61 299.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bandcamp(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1070 1178l306-564H722l-306 564zm722-282q0 182-71 348t-191 286t-286 191t-348 71t-348-71t-286-191t-191-286T0 896t71-348t191-286T548 71T896 0t348 71t286 191t191 286t71 348"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bank(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m960 0l960 384v128h-128q0 26-20.5 45t-48.5 19H197q-28 0-48.5-19T128 512H0V384zM256 640h256v768h128V640h256v768h128V640h256v768h128V640h256v768h59q28 0 48.5 19t20.5 45v64H128v-64q0-26 20.5-45t48.5-19h59zm1595 960q28 0 48.5 19t20.5 45v128H0v-128q0-26 20.5-45t48.5-19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarChart(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 768v512H384V768zm384-512v1024H768V256zm1024 1152v128H0V0h128v1408zm-640-896v768h-256V512zm384-384v1152h-256V128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Barcode(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M63 1408H0V0h63zm63-1H94V0h32zm94 0h-31V0h31zm157 0h-31V0h31zm157 0h-62V0h62zm126 0h-31V0h31zm63 0h-31V0h31zm63 0h-31V0h31zm157 0h-63V0h63zm157 0h-63V0h63zm126 0h-63V0h63zm126 0h-63V0h63zm94 0h-63V0h63zm189 0h-94V0h94zm63 0h-32V0h32zm94 1h-63V0h63z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bars(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1536 1088v128q0 26-19 45t-45 19H64q-26 0-45-19t-19-45v-128q0-26 19-45t45-19h1408q26 0 45 19t19 45m0-512v128q0 26-19 45t-45 19H64q-26 0-45-19T0 704V576q0-26 19-45t45-19h1408q26 0 45 19t19 45m0-512v128q0 26-19 45t-45 19H64q-26 0-45-19T0 192V64q0-26 19-45T64 0h1408q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bath(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1664 1088v192q0 169-128 286v194q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23v-118q-63 22-128 22H512q-65 0-128-22v110q0 17-9.5 28.5T352 1792h-64q-13 0-22.5-11.5T256 1752v-186q-128-117-128-286v-192zM704 672q0 14-9 23t-23 9t-23-9t-9-23t9-23t23-9t23 9t9 23m64-64q0 14-9 23t-23 9t-23-9t-9-23t9-23t23-9t23 9t9 23m-64-64q0 14-9 23t-23 9t-23-9t-9-23t9-23t23-9t23 9t9 23m128 0q0 14-9 23t-23 9t-23-9t-9-23t9-23t23-9t23 9t9 23m-64-64q0 14-9 23t-23 9t-23-9t-9-23t9-23t23-9t23 9t9 23m-64-64q0 14-9 23t-23 9t-23-9t-9-23t9-23t23-9t23 9t9 23m1088 512v64q0 14-9 23t-23 9H32q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h96V256q0-106 75-181T384 0q108 0 184 78q46-19 98-12t93 39l22-22q11-11 22 0l42 42q11 11 0 22L531 461q-11 11-22 0l-42-42q-11-11 0-22l22-22q-36-46-40.5-104T472 163q-37-35-88-35q-53 0-90.5 37.5T256 256v640h1504q14 0 23 9t9 23M896 480q0 14-9 23t-23 9t-23-9t-9-23t9-23t23-9t23 9t9 23m-64-64q0 14-9 23t-23 9t-23-9t-9-23t9-23t23-9t23 9t9 23m-64-64q0 14-9 23t-23 9t-23-9t-9-23t9-23t23-9t23 9t9 23m192 64q0 14-9 23t-23 9t-23-9t-9-23t9-23t23-9t23 9t9 23m-64-64q0 14-9 23t-23 9t-23-9t-9-23t9-23t23-9t23 9t9 23m-64-64q0 14-9 23t-23 9t-23-9t-9-23t9-23t23-9t23 9t9 23m192 64q0 14-9 23t-23 9t-23-9t-9-23t9-23t23-9t23 9t9 23m-64-64q0 14-9 23t-23 9t-23-9t-9-23t9-23t23-9t23 9t9 23m128 0q0 14-9 23t-23 9t-23-9t-9-23t9-23t23-9t23 9t9 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Battery(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1920 256v768H256V256zm128 576h128V448h-128V160q0-14-9-23t-23-9H160q-14 0-23 9t-9 23v960q0 14 9 23t23 9h1856q14 0 23-9t9-23zm256-384v384q0 53-37.5 90.5T2176 960v160q0 66-47 113t-113 47H160q-66 0-113-47T0 1120V160Q0 94 47 47T160 0h1856q66 0 113 47t47 113v160q53 0 90.5 37.5T2304 448"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryOne(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 1024V256h512v768zm1920-704q53 0 90.5 37.5T2304 448v384q0 53-37.5 90.5T2176 960v160q0 66-47 113t-113 47H160q-66 0-113-47T0 1120V160Q0 94 47 47T160 0h1856q66 0 113 47t47 113zm0 512V448h-128V160q0-14-9-23t-23-9H160q-14 0-23 9t-9 23v960q0 14 9 23t23 9h1856q14 0 23-9t9-23V832z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryThree(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 1024V256h1280v768zm1920-704q53 0 90.5 37.5T2304 448v384q0 53-37.5 90.5T2176 960v160q0 66-47 113t-113 47H160q-66 0-113-47T0 1120V160Q0 94 47 47T160 0h1856q66 0 113 47t47 113zm0 512V448h-128V160q0-14-9-23t-23-9H160q-14 0-23 9t-9 23v960q0 14 9 23t23 9h1856q14 0 23-9t9-23V832z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryTwo(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 1024V256h896v768zm1920-704q53 0 90.5 37.5T2304 448v384q0 53-37.5 90.5T2176 960v160q0 66-47 113t-113 47H160q-66 0-113-47T0 1120V160Q0 94 47 47T160 0h1856q66 0 113 47t47 113zm0 512V448h-128V160q0-14-9-23t-23-9H160q-14 0-23 9t-9 23v960q0 14 9 23t23 9h1856q14 0 23-9t9-23V832z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryZero(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2176 320q53 0 90.5 37.5T2304 448v384q0 53-37.5 90.5T2176 960v160q0 66-47 113t-113 47H160q-66 0-113-47T0 1120V160Q0 94 47 47T160 0h1856q66 0 113 47t47 113zm0 512V448h-128V160q0-14-9-23t-23-9H160q-14 0-23 9t-9 23v960q0 14 9 23t23 9h1856q14 0 23-9t9-23V832z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bed(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 768h1728q26 0 45 19t19 45v448h-256v-256H256v256H0V64q0-26 19-45T64 0h128q26 0 45 19t19 45zm576-320q0-106-75-181t-181-75t-181 75t-75 181t75 181t181 75t181-75t75-181m1216 256v-64q0-159-112.5-271.5T1664 256H960q-26 0-45 19t-19 45v384z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Beer(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M576 768V384H320v256q0 53 37.5 90.5T448 768zm1024 448v192H448v-192l128-192H448q-159 0-271.5-112.5T64 640V320L0 256l32-128h480L544 0h960l32 192l-64 32v800z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Behance(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1848 115h-511v124h511zm-252 426q-90 0-146 52.5T1388 736h408q-18-195-200-195m16 585q63 0 122-32t76-87h221q-100 307-427 307q-214 0-340.5-132T1137 835q0-208 130.5-345.5T1604 352q138 0 240.5 68t153 179t50.5 248q0 17-2 47h-658q0 111 57.5 171.5T1612 1126m-1335-50h296q205 0 205-167q0-180-199-180H277zm0-537h281q78 0 123.5-36.5T727 389q0-144-190-144H277zM0 30h594q87 0 155 14t126.5 47.5t90 96.5T997 342q0 181-172 263q114 32 172 115t58 204q0 75-24.5 136.5t-66 103.5t-98.5 71t-121 42t-134 13H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BehanceSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1248 0q119 0 203.5 84.5T1536 288v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0zM499 367H128v787h382q117 0 197-57.5T787 926q0-158-143-200q107-52 107-164q0-57-19.5-96.5T675 405t-79-29.5t-97-8.5m-22 318H301V501h163q119 0 119 90q0 94-106 94m9 335H301V803h189q124 0 124 113q0 104-128 104m650 32q-68 0-104-38t-36-107h411q1-10 1-30q0-132-74.5-220.5T1130 568q-128 0-210 86t-82 216q0 135 79 217t213 82q205 0 267-191h-138q-11 34-47.5 54t-75.5 20m-10-366q113 0 124 122H996q4-56 39-89t91-33M964 420h319v77H964z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M848 1696q0-16-16-16q-59 0-101.5-42.5T688 1536q0-16-16-16t-16 16q0 73 51.5 124.5T832 1712q16 0 16-16m816-288q0 52-38 90t-90 38h-448q0 106-75 181t-181 75t-181-75t-75-181H128q-52 0-90-38t-38-90q50-42 91-88t85-119.5t74.5-158.5t50-206T320 576q0-152 117-282.5T744 135q-8-19-8-39q0-40 28-68t68-28t68 28t28 68q0 20-8 39q190 28 307 158.5T1344 576q0 139 19.5 260t50 206t74.5 158.5t85 119.5t91 88"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M848 1696q0-16-16-16q-59 0-101.5-42.5T688 1536q0-16-16-16t-16 16q0 73 51.5 124.5T832 1712q16 0 16-16m-666-288h1300q-266-300-266-832q0-51-24-105t-69-103t-121.5-80.5T832 256t-169.5 31.5T541 368t-69 103t-24 105q0 532-266 832m1482 0q0 52-38 90t-90 38h-448q0 106-75 181t-181 75t-181-75t-75-181H128q-52 0-90-38t-38-90q50-42 91-88t85-119.5t74.5-158.5t50-206T320 576q0-152 117-282.5T744 135q-8-19-8-39q0-40 28-68t68-28t68 28t28 68q0 20-8 39q190 28 307 158.5T1344 576q0 139 19.5 260t50 206t74.5 158.5t85 119.5t91 88"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellSlash(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1558 852q61 356 298 556q0 52-38 90t-90 38h-448q0 106-75 181t-181 75t-180.5-74.5T768 1537zm-534 860q16 0 16-16t-16-16q-59 0-101.5-42.5T880 1536q0-16-16-16t-16 16q0 73 51.5 124.5T1024 1712M2026 112q8 10 7.5 23.5T2023 158L151 1780q-10 8-23.5 7t-21.5-11l-84-96q-8-10-7.5-23.5T25 1635l186-161q-19-32-19-66q50-42 91-88t85-119.5t74.5-158.5t50-206T512 576q0-152 117-282.5T936 135q-8-19-8-39q0-40 28-68t68-28t68 28t28 68q0 20-8 39q124 18 219 82.5T1479 375l418-363q10-8 23.5-7t21.5 11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellSlashO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1040 1696q0-16-16-16q-59 0-101.5-42.5T880 1536q0-16-16-16t-16 16q0 73 51.5 124.5T1024 1712q16 0 16-16m-537-475l877-760q-42-88-132.5-146.5T1024 256q-93 0-169.5 31.5T733 368t-69 103t-24 105q0 384-137 645m1353 187q0 52-38 90t-90 38h-448q0 106-75 181t-181 75t-180.5-74.5T768 1537l149-129h757q-166-187-227-459l111-97q61 356 298 556m86-1392l84 96q8 10 7.5 23.5T2023 158L151 1780q-10 8-23.5 7t-21.5-11l-84-96q-8-10-7.5-23.5T25 1635l186-161q-19-32-19-66q50-42 91-88t85-119.5t74.5-158.5t50-206T512 576q0-152 117-282.5T936 135q-8-19-8-39q0-40 28-68t68-28t68 28t28 68q0 20-8 39q124 18 219 82.5T1479 375l418-363q10-8 23.5-7t21.5 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bicycle(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M762 896H448q-40 0-57.5-35t6.5-67l188-251q-65-31-137-31q-132 0-226 94t-94 226t94 226t226 94q115 0 203-72.5T762 896M576 768h186q-18-85-75-148zm480 0l288-384H864l-99 132q105 103 126 252zm1120 64q0-132-94-226t-226-94q-60 0-121 24l174 260q15 23 10 49t-27 40q-15 11-36 11q-35 0-53-29l-174-260q-93 95-93 225q0 132 94 226t226 94t226-94t94-226m128 0q0 185-131.5 316.5T1856 1280t-316.5-131.5T1408 832q0-97 39.5-183.5T1557 499l-65-98l-353 469q-18 26-51 26H891q-23 164-149 274t-294 110q-185 0-316.5-131.5T0 832t131.5-316.5T448 384q114 0 215 55l137-183H576q-26 0-45-19t-19-45t19-45t45-19h384v128h435l-85-128h-222q-26 0-45-19t-19-45t19-45t45-19h256q33 0 53 28l267 400q91-44 192-44q185 0 316.5 131.5T2304 832"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Binoculars(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M704 320v768q0 26-19 45t-45 19v576q0 26-19 45t-45 19H64q-26 0-45-19t-19-45v-512l249-873q7-23 31-23zm320 0v704H768V320zm768 896v512q0 26-19 45t-45 19h-512q-26 0-45-19t-19-45v-576q-26 0-45-19t-19-45V320h424q24 0 31 23zM736 32v224H384V32q0-14 9-23t23-9h288q14 0 23 9t9 23m672 0v224h-352V32q0-14 9-23t23-9h288q14 0 23 9t9 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BirthdayCake(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 1408v384H0v-384q45 0 85-14t59-27.5t47-37.5q30-27 51.5-38t56.5-11q24 0 44 7t31 15t33 27q29 25 47 38t58 27t86 14q45 0 85-14.5t58-27t48-37.5q21-19 32.5-27t31-15t43.5-7q35 0 56.5 11t51.5 38q28 24 47 37.5t59 27.5t85 14t85-14t59-27.5t47-37.5q30-27 51.5-38t56.5-11q34 0 55.5 11t51.5 38q28 24 47 37.5t59 27.5t85 14m0-320v192q-24 0-44-7t-31-15t-33-27q-29-25-47-38t-58-27t-85-14q-46 0-86 14t-58 27t-47 38q-22 19-33 27t-31 15t-44 7q-35 0-56.5-11t-51.5-38q-29-25-47-38t-58-27t-86-14q-45 0-85 14.5t-58 27t-48 37.5q-21 19-32.5 27t-31 15t-43.5 7q-35 0-56.5-11t-51.5-38q-28-24-47-37.5t-59-27.5t-85-14q-46 0-86 14t-58 27t-47 38q-30 27-51.5 38T0 1280v-192q0-80 56-136t136-56h64V448h256v448h256V448h256v448h256V448h256v448h64q80 0 136 56t56 136M512 224q0 77-36 118.5T384 384q-53 0-90.5-37.5T256 256q0-29 9.5-51t23.5-34t31-28t31-31.5T374.5 67T384 0q38 0 83 74t45 150m512 0q0 77-36 118.5T896 384q-53 0-90.5-37.5T768 256q0-29 9.5-51t23.5-34t31-28t31-31.5T886.5 67T896 0q38 0 83 74t45 150m512 0q0 77-36 118.5t-92 41.5q-53 0-90.5-37.5T1280 256q0-29 9.5-51t23.5-34t31-28t31-31.5t23.5-44.5t9.5-67q38 0 83 74t45 150"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bitbucket(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M815 795q8 63-50.5 101T653 902q-39-17-53.5-58t-.5-82t52-58q36-18 72.5-12t64 35.5T815 795m111-21q-14-107-113-164t-197-13q-63 28-100.5 88.5T481 815q4 91 77.5 155t165.5 56q91-8 152-84t50-168m239-542q-20-27-56-44.5t-58-22t-71-12.5q-291-47-566 2q-43 7-66 12t-55 22t-50 43q30 28 76 45.5t73.5 22T480 311q228 29 448 1q63-8 89.5-12t72.5-21.5t75-46.5m57 1035q-8 26-15.5 76.5t-14 84t-28.5 70t-58 56.5q-86 48-189.5 71.5t-202 22T513 1629q-46-8-81.5-18t-76.5-27t-73-43.5t-52-61.5q-25-96-57-292l6-16l18-9q223 148 506.5 148t507.5-148q21 6 24 23t-5 45t-8 37m181-961q-26 167-111 655q-5 30-27 56t-43.5 40t-54.5 31q-252 126-610 88q-248-27-394-139q-15-12-25.5-26.5t-17-35t-9-34t-6-39.5t-5.5-35q-9-50-26.5-150t-28-161.5T22 408T0 250q3-26 17.5-48.5T49 164t45-30t46-22.5T188 93q125-46 313-64q379-37 676 50q155 46 215 122q16 20 16.5 51t-5.5 54"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BitbucketSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M848 742q0-43-41-66t-77-1q-43 20-42.5 72.5T731 818q39 23 81-4t36-72m80-16q8 66-36 121t-110 61t-119-40t-56-113q-2-49 25.5-93t72.5-64q70-31 141.5 10T928 726m172-391q-20 21-53.5 34t-53 16t-63.5 8q-155 20-324 0q-44-6-63-9.5t-52.5-16T436 335q13-19 36-31t40-15.5t47-8.5q198-35 408-1q33 5 51 8.5t43 16t39 31.5m42 746q0-7 5.5-26.5t3-32t-17.5-16.5q-161 106-365 106t-366-106l-12 6l-5 12q26 154 41 210q47 81 204 108q249 46 428-53q34-19 49-51.5t22.5-85.5t12.5-71m130-693q9-53-8-75q-43-55-155-88q-216-63-487-36q-132 12-226 46q-38 15-59.5 25t-47 34t-29.5 54q8 68 19 138t29 171t24 137q1 5 5 31t7 36t12 27t22 28q105 80 284 100q259 28 440-63q24-13 39.5-23t31-29t19.5-40q48-267 80-473m264-100v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bitcoin(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1135 512q18 182-131 258q117 28 175 103t45 214q-7 71-32.5 125t-64.5 89t-97 58.5t-121.5 34.5t-145.5 15v255H609v-251q-80 0-122-1v252H333v-255q-18 0-54-.5t-55-.5H24l31-183h111q50 0 58-51V772h16q-6-1-16-1V484q-13-68-89-68H24V252l212 1q64 0 97-1V0h154v247q82-2 122-2V0h154v252q79 7 140 22.5t113 45t82.5 78T1135 512m-215 545q0-36-15-64t-37-46t-57.5-30.5T745 898t-74-9t-69-3t-64.5 1t-47.5 1v338q8 0 37 .5t48 .5t53-1.5t58.5-4t57-8.5t55.5-14t47.5-21t39.5-30t24.5-40t9.5-51m-71-476q0-33-12.5-58.5t-30.5-42t-48-28t-55-16.5t-61.5-8t-58-2.5t-54 1t-39.5.5v307q5 0 34.5.5t46.5 0t50-2t55-5.5t51.5-11t48.5-18.5t37-27t27-38.5t9-51"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BlackTie(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h1536v1536H0zm1085 1115L864 484l221-297H451l221 297l-221 631l317 304z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blind(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M334 311q-64 0-110-45.5T178 155q0-64 46-109.5T334 0t109.5 45.5T489 155q0 65-45.5 110.5T334 311m551 642q0 50-30 67.5t-63.5 6.5t-47.5-34L377 555q-7-12-14-15.5t-11-1.5l-3 3q-7 8 4 21l122 139l1 354l-161 457q-67 192-92 234q-15 26-28 32q-50 26-103 1q-29-13-41.5-43t-9.5-57q2-17 197-618l5-416l-85 164l35 222q4 24-1 42t-14 27.5t-19 16t-17 7.5l-7 2q-19 3-34.5-3t-24-16t-14-22t-7.5-19.5t-2-9.5L7 757l211-381q23-34 113-34q75 0 107 40l424 521q7 5 14 17l3 3l-1 1q7 13 7 29m-403 150q43 113 88.5 225t69.5 168l24 55q36 93 42 125q11 70-36 97q-35 22-66 16t-51-22t-29-35h-1q-6-16-8-25l-124-351zm824 592q31 49 31 57q0 5-3 7q-9 5-14.5-.5t-15.5-26t-16-30.5q-114-172-423-661q3 1 7-1t7-4l3-2q11-9 11-17z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bluetooth(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m745 1053l148 148l-149 149zm-1-611l149 149l-148 148zM614 1666l464-464l-306-306l306-306l-464-464v611L359 482l-93 93l320 321l-320 321l93 93l255-255zm719-770q0 209-32 365.5t-87.5 257T1073 1681t-181.5 86.5T672 1792t-219.5-24.5T271 1681t-140.5-162.5t-87.5-257T11 896t32-365.5t87.5-257T271 111t181.5-86.5T672 0t219.5 24.5T1073 111t140.5 162.5t87.5 257t32 365.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BluetoothB(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m564 1423l173-172l-173-172zm0-710l173-172l-173-172zm32 183l356 356l-539 540v-711l-297 296L8 1269l372-373L8 523l108-108l297 296V0l539 540z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bold(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M555 1393q74 32 140 32q376 0 376-335q0-114-41-180q-27-44-61.5-74T901 789.5t-80.5-25t-84-10.5t-94.5-2q-73 0-101 10q0 53-.5 159t-.5 158q0 8-1 67.5t-.5 96.5t4.5 83.5t12 66.5m-14-746q42 7 109 7q82 0 143-13t110-44.5t74.5-89.5t25.5-142q0-70-29-122.5t-79-82T787 117t-124-14q-50 0-130 13q0 50 4 151t4 152q0 27-.5 80t-.5 79q0 46 1 69M0 1536l2-94q15-4 85-16t106-27q7-12 12.5-27t8.5-33.5t5.5-32.5t3-37.5t.5-34V1169q0-982-22-1025q-4-8-22-14.5t-44.5-11t-49.5-7t-48.5-4.5T6 104L2 21q98-2 340-11.5T715 0q23 0 68 .5t68 .5q70 0 136.5 13T1116 56t108 71t74 104.5t28 137.5q0 52-16.5 95.5t-39 72T1206 594t-73 45t-84 40q154 35 256.5 134t102.5 248q0 100-35 179.5t-93.5 130.5t-138 85.5T978 1505t-176 14q-44 0-132-3t-132-3q-106 0-307 11T0 1536"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bolt(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M885 438q18 20 7 44L352 1639q-13 25-42 25q-4 0-14-2q-17-5-25.5-19t-4.5-30l197-808L57 906q-4 1-12 1q-18 0-31-11q-18-15-13-39L202 32q4-14 16-23t28-9h328q19 0 32 12.5T619 42q0 8-5 18L443 523l396-98q8-2 12-2q19 0 34 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bomb(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M571 589q-10-25-34-35t-49 0q-108 44-191 127T170 872q-10 25 0 49t35 34q13 5 24 5q42 0 60-40q34-84 98.5-148.5T536 673q25-11 35-35t0-49m942-356l46 46l-244 243l68 68q19 19 19 45.5t-19 45.5l-64 64q89 161 89 343q0 143-55.5 273.5t-150 225t-225 150T704 1792t-273.5-55.5t-225-150t-150-225T0 1088t55.5-273.5t150-225t225-150T704 384q182 0 343 89l64-64q19-19 45.5-19t45.5 19l68 68zm8-56q-10 10-22 10q-13 0-23-10l-91-90q-9-10-9-23t9-23q10-9 23-9t23 9l90 91q10 9 10 22.5t-10 22.5m230 230q-11 9-23 9t-23-9l-90-91q-10-9-10-22.5t10-22.5q9-10 22.5-10t22.5 10l91 90q9 10 9 23t-9 23m41-183q0 14-9 23t-23 9h-96q-14 0-23-9t-9-23t9-23t23-9h96q14 0 23 9t9 23M1600 32v96q0 14-9 23t-23 9t-23-9t-9-23V32q0-14 9-23t23-9t23 9t9 23m151 55l-91 90q-10 10-22 10q-13 0-23-10q-10-9-10-22.5t10-22.5l90-91q10-9 23-9t23 9q9 10 9 23t-9 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Book(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1671 350q40 57 18 129l-275 906q-19 64-76.5 107.5T1215 1536H292q-77 0-148.5-53.5T44 1351q-24-67-2-127q0-4 3-27t4-37q1-8-3-21.5t-3-19.5q2-11 8-21t16.5-23.5T84 1051q23-38 45-91.5t30-91.5q3-10 .5-30t-.5-28q3-11 17-28t17-23q21-36 42-92t25-90q1-9-2.5-32t.5-28q4-13 22-30.5t22-22.5q19-26 42.5-84.5T372 283q1-8-3-25.5t-2-26.5q2-8 9-18t18-23t17-21q8-12 16.5-30.5t15-35t16-36t19.5-32T504.5 12t36-11.5T588 6l-1 3q38-9 51-9h761q74 0 114 56t18 130l-274 906q-36 119-71.5 153.5T1057 1280H188q-27 0-38 15q-11 16-1 43q24 70 144 70h923q29 0 56-15.5t35-41.5l300-987q7-22 5-57q38 15 59 43m-1064 2q-4 13 2 22.5t20 9.5h608q13 0 25.5-9.5T1279 352l21-64q4-13-2-22.5t-20-9.5H670q-13 0-25.5 9.5T628 288zm-83 256q-4 13 2 22.5t20 9.5h608q13 0 25.5-9.5T1196 608l21-64q4-13-2-22.5t-20-9.5H587q-13 0-25.5 9.5T545 544z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1164 0q23 0 44 9q33 13 52.5 41t19.5 62v1289q0 34-19.5 62t-52.5 41q-19 8-44 8q-48 0-83-32l-441-424l-441 424q-36 33-83 33q-23 0-44-9q-33-13-52.5-41T0 1401V112q0-34 19.5-62T72 9q21-9 44-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1152 128H128v1242l423-406l89-85l89 85l423 406zm12-128q23 0 44 9q33 13 52.5 41t19.5 62v1289q0 34-19.5 62t-52.5 41q-19 8-44 8q-48 0-83-32l-441-424l-441 424q-36 33-83 33q-23 0-44-9q-33-13-52.5-41T0 1401V112q0-34 19.5-62T72 9q21-9 44-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Braille(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 1056q-66 0-113 47t-47 113t47 113t113 47t113-47t47-113t-47-113t-113-47m512 0q-66 0-113 47t-47 113t47 113t113 47t113-47t47-113t-47-113t-113-47m0-512q-66 0-113 47t-47 113t47 113t113 47t113-47t47-113t-47-113t-113-47m768 512q-66 0-113 47t-47 113t47 113t113 47t113-47t47-113t-47-113t-113-47m512 0q-66 0-113 47t-47 113t47 113t113 47t113-47t47-113t-47-113t-113-47m-512-512q-66 0-113 47t-47 113t47 113t113 47t113-47t47-113t-47-113t-113-47m512 0q-66 0-113 47t-47 113t47 113t113 47t113-47t47-113t-47-113t-113-47m0-512q-66 0-113 47t-47 113t47 113t113 47t113-47t47-113t-47-113t-113-47M384 1216q0 80-56 136t-136 56t-136-56t-56-136t56-136t136-56t136 56t56 136m512 0q0 80-56 136t-136 56t-136-56t-56-136t56-136t136-56t136 56t56 136M384 704q0 80-56 136t-136 56t-136-56T0 704t56-136t136-56t136 56t56 136m512 0q0 80-56 136t-136 56t-136-56t-56-136t56-136t136-56t136 56t56 136M384 192q0 80-56 136t-136 56t-136-56T0 192T56 56T192 0t136 56t56 136m1280 1024q0 80-56 136t-136 56t-136-56t-56-136t56-136t136-56t136 56t56 136M896 192q0 80-56 136t-136 56t-136-56t-56-136t56-136T704 0t136 56t56 136m1280 1024q0 80-56 136t-136 56t-136-56t-56-136t56-136t136-56t136 56t56 136m-512-512q0 80-56 136t-136 56t-136-56t-56-136t56-136t136-56t136 56t56 136m512 0q0 80-56 136t-136 56t-136-56t-56-136t56-136t136-56t136 56t56 136m-512-512q0 80-56 136t-136 56t-136-56t-56-136t56-136t136-56t136 56t56 136m512 0q0 80-56 136t-136 56t-136-56t-56-136t56-136t136-56t136 56t56 136"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Briefcase(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 256h512V128H640zm1152 640v480q0 66-47 113t-113 47H160q-66 0-113-47T0 1376V896h672v160q0 26 19 45t45 19h320q26 0 45-19t19-45V896zm-768 0v128H768V896zm768-480v384H0V416q0-66 47-113t113-47h352V96q0-40 28-68t68-28h576q40 0 68 28t28 68v160h352q66 0 113 47t47 113"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bug(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1600 896q0 26-19 45t-45 19h-224q0 171-67 290l208 209q19 19 19 45t-19 45q-18 19-45 19t-45-19l-198-197q-5 5-15 13t-42 28.5t-65 36.5t-82 29t-97 13V576H736v896q-51 0-101.5-13.5t-87-33t-66-39T438 1354l-15-14l-183 207q-20 21-48 21q-24 0-43-16q-19-18-20.5-44.5T144 1461l202-227q-58-114-58-274H64q-26 0-45-19T0 896t19-45t45-19h224V538L115 365q-19-19-19-45t19-45t45-19t45 19l173 173h844l173-173q19-19 45-19t45 19t19 45t-19 45l-173 173v294h224q26 0 45 19t19 45m-480-576H480q0-133 93.5-226.5T800 0t226.5 93.5T1120 320"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Building(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1344 0q26 0 45 19t19 45v1664q0 26-19 45t-45 19H64q-26 0-45-19t-19-45V64q0-26 19-45T64 0zM512 288v64q0 14 9 23t23 9h64q14 0 23-9t9-23v-64q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23m0 256v64q0 14 9 23t23 9h64q14 0 23-9t9-23v-64q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23m0 256v64q0 14 9 23t23 9h64q14 0 23-9t9-23v-64q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23m0 256v64q0 14 9 23t23 9h64q14 0 23-9t9-23v-64q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23m-128 320v-64q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v64q0 14 9 23t23 9h64q14 0 23-9t9-23m0-256v-64q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v64q0 14 9 23t23 9h64q14 0 23-9t9-23m0-256v-64q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v64q0 14 9 23t23 9h64q14 0 23-9t9-23m0-256v-64q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v64q0 14 9 23t23 9h64q14 0 23-9t9-23m0-256v-64q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v64q0 14 9 23t23 9h64q14 0 23-9t9-23m512 1280v-192q0-14-9-23t-23-9H544q-14 0-23 9t-9 23v192q0 14 9 23t23 9h320q14 0 23-9t9-23m0-512v-64q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v64q0 14 9 23t23 9h64q14 0 23-9t9-23m0-256v-64q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v64q0 14 9 23t23 9h64q14 0 23-9t9-23m0-256v-64q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v64q0 14 9 23t23 9h64q14 0 23-9t9-23m0-256v-64q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v64q0 14 9 23t23 9h64q14 0 23-9t9-23m256 1024v-64q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v64q0 14 9 23t23 9h64q14 0 23-9t9-23m0-256v-64q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v64q0 14 9 23t23 9h64q14 0 23-9t9-23m0-256v-64q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v64q0 14 9 23t23 9h64q14 0 23-9t9-23m0-256v-64q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v64q0 14 9 23t23 9h64q14 0 23-9t9-23m0-256v-64q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v64q0 14 9 23t23 9h64q14 0 23-9t9-23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 1312v64q0 13-9.5 22.5T352 1408h-64q-13 0-22.5-9.5T256 1376v-64q0-13 9.5-22.5t22.5-9.5h64q13 0 22.5 9.5t9.5 22.5m0-256v64q0 13-9.5 22.5T352 1152h-64q-13 0-22.5-9.5T256 1120v-64q0-13 9.5-22.5t22.5-9.5h64q13 0 22.5 9.5t9.5 22.5m256 0v64q0 13-9.5 22.5T608 1152h-64q-13 0-22.5-9.5T512 1120v-64q0-13 9.5-22.5t22.5-9.5h64q13 0 22.5 9.5t9.5 22.5M384 800v64q0 13-9.5 22.5T352 896h-64q-13 0-22.5-9.5T256 864v-64q0-13 9.5-22.5T288 768h64q13 0 22.5 9.5T384 800m768 512v64q0 13-9.5 22.5t-22.5 9.5h-64q-13 0-22.5-9.5t-9.5-22.5v-64q0-13 9.5-22.5t22.5-9.5h64q13 0 22.5 9.5t9.5 22.5m-256-256v64q0 13-9.5 22.5T864 1152h-64q-13 0-22.5-9.5T768 1120v-64q0-13 9.5-22.5t22.5-9.5h64q13 0 22.5 9.5t9.5 22.5M640 800v64q0 13-9.5 22.5T608 896h-64q-13 0-22.5-9.5T512 864v-64q0-13 9.5-22.5T544 768h64q13 0 22.5 9.5T640 800M384 544v64q0 13-9.5 22.5T352 640h-64q-13 0-22.5-9.5T256 608v-64q0-13 9.5-22.5T288 512h64q13 0 22.5 9.5T384 544m768 512v64q0 13-9.5 22.5t-22.5 9.5h-64q-13 0-22.5-9.5t-9.5-22.5v-64q0-13 9.5-22.5t22.5-9.5h64q13 0 22.5 9.5t9.5 22.5M896 800v64q0 13-9.5 22.5T864 896h-64q-13 0-22.5-9.5T768 864v-64q0-13 9.5-22.5T800 768h64q13 0 22.5 9.5T896 800M640 544v64q0 13-9.5 22.5T608 640h-64q-13 0-22.5-9.5T512 608v-64q0-13 9.5-22.5T544 512h64q13 0 22.5 9.5T640 544M384 288v64q0 13-9.5 22.5T352 384h-64q-13 0-22.5-9.5T256 352v-64q0-13 9.5-22.5T288 256h64q13 0 22.5 9.5T384 288m768 512v64q0 13-9.5 22.5T1120 896h-64q-13 0-22.5-9.5T1024 864v-64q0-13 9.5-22.5t22.5-9.5h64q13 0 22.5 9.5t9.5 22.5M896 544v64q0 13-9.5 22.5T864 640h-64q-13 0-22.5-9.5T768 608v-64q0-13 9.5-22.5T800 512h64q13 0 22.5 9.5T896 544M640 288v64q0 13-9.5 22.5T608 384h-64q-13 0-22.5-9.5T512 352v-64q0-13 9.5-22.5T544 256h64q13 0 22.5 9.5T640 288m512 256v64q0 13-9.5 22.5T1120 640h-64q-13 0-22.5-9.5T1024 608v-64q0-13 9.5-22.5t22.5-9.5h64q13 0 22.5 9.5t9.5 22.5M896 288v64q0 13-9.5 22.5T864 384h-64q-13 0-22.5-9.5T768 352v-64q0-13 9.5-22.5T800 256h64q13 0 22.5 9.5T896 288m256 0v64q0 13-9.5 22.5T1120 384h-64q-13 0-22.5-9.5T1024 352v-64q0-13 9.5-22.5t22.5-9.5h64q13 0 22.5 9.5t9.5 22.5M896 1664h384V128H128v1536h384v-224q0-13 9.5-22.5t22.5-9.5h320q13 0 22.5 9.5t9.5 22.5zM1408 64v1664q0 26-19 45t-45 19H64q-26 0-45-19t-19-45V64q0-26 19-45T64 0h1280q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bullhorn(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1664 512q53 0 90.5 37.5T1792 640t-37.5 90.5T1664 768v384q0 52-38 90t-90 38q-417-347-812-380q-58 19-91 66t-31 100.5t40 92.5q-20 33-23 65.5t6 58t33.5 55t48 50T768 1438q-29 58-111.5 83T488 1532.5T356 1477q-7-23-29.5-87.5t-32-94.5t-23-89t-15-101t3.5-98.5T282 896H160q-66 0-113-47T0 736V544q0-66 47-113t113-47h480q435 0 896-384q52 0 90 38t38 90zm-128 604V162q-394 302-768 343v270q377 42 768 341"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bullseye(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1024 768q0 106-75 181t-181 75t-181-75t-75-181t75-181t181-75t181 75t75 181m128 0q0-159-112.5-271.5T768 384T496.5 496.5T384 768t112.5 271.5T768 1152t271.5-112.5T1152 768m128 0q0 212-150 362t-362 150t-362-150t-150-362t150-362t362-150t362 150t150 362m128 0q0-130-51-248.5t-136.5-204t-204-136.5T768 128t-248.5 51t-204 136.5t-136.5 204T128 768t51 248.5t136.5 204t204 136.5t248.5 51t248.5-51t204-136.5t136.5-204t51-248.5m128 0q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bus(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 1216q0-53-37.5-90.5T256 1088t-90.5 37.5T128 1216t37.5 90.5T256 1344t90.5-37.5T384 1216m1024 0q0-53-37.5-90.5T1280 1088t-90.5 37.5t-37.5 90.5t37.5 90.5t90.5 37.5t90.5-37.5t37.5-90.5m-46-396l-72-384q-5-23-22.5-37.5T1227 384H309q-23 0-40.5 14.5T246 436l-72 384q-5 30 14 53t49 23h1062q30 0 49-23t14-53m-226-612q0-20-14-34t-34-14H448q-20 0-34 14t-14 34t14 34t34 14h640q20 0 34-14t14-34m400 725v603h-128v128q0 53-37.5 90.5T1280 1792t-90.5-37.5t-37.5-90.5v-128H384v128q0 53-37.5 90.5T256 1792t-90.5-37.5T128 1664v-128H0V933q0-112 25-223l103-454q9-78 97.5-137t230-89T768 0t312.5 30t230 89t97.5 137l105 454q23 102 23 223"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Buysellads(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M915 958H621l147-551zm86 322h311L988 256H548L224 1280h311l383-314zm535-992v960q0 118-85 203t-203 85H288q-118 0-203-85T0 1248V288Q0 170 85 85T288 0h960q118 0 203 85t85 203"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cab(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1824 896q93 0 158.5 65.5T2048 1120v384q0 14-9 23t-23 9h-96v64q0 80-56 136t-136 56t-136-56t-56-136v-64H512v64q0 80-56 136t-136 56t-136-56t-56-136v-64H32q-14 0-23-9t-9-23v-384q0-93 65.5-158.5T224 896h28l105-419q23-94 104-157.5T640 256h128V32q0-14 9-23t23-9h448q14 0 23 9t9 23v224h128q98 0 179 63.5T1691 477l105 419zM320 1376q66 0 113-47t47-113t-47-113t-113-47t-113 47t-47 113t47 113t113 47m196-480h1016l-89-357q-2-8-14-17.5t-21-9.5H640q-9 0-21 9.5T605 539zm1212 480q66 0 113-47t47-113t-47-113t-113-47t-113 47t-47 113t47 113t113 47"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calculator(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 1536q0-53-37.5-90.5T256 1408t-90.5 37.5T128 1536t37.5 90.5T256 1664t90.5-37.5T384 1536m384 0q0-53-37.5-90.5T640 1408t-90.5 37.5T512 1536t37.5 90.5T640 1664t90.5-37.5T768 1536m-384-384q0-53-37.5-90.5T256 1024t-90.5 37.5T128 1152t37.5 90.5T256 1280t90.5-37.5T384 1152m768 384q0-53-37.5-90.5T1024 1408t-90.5 37.5T896 1536t37.5 90.5t90.5 37.5t90.5-37.5t37.5-90.5m-384-384q0-53-37.5-90.5T640 1024t-90.5 37.5T512 1152t37.5 90.5T640 1280t90.5-37.5T768 1152M384 768q0-53-37.5-90.5T256 640t-90.5 37.5T128 768t37.5 90.5T256 896t90.5-37.5T384 768m768 384q0-53-37.5-90.5T1024 1024t-90.5 37.5T896 1152t37.5 90.5t90.5 37.5t90.5-37.5t37.5-90.5M768 768q0-53-37.5-90.5T640 640t-90.5 37.5T512 768t37.5 90.5T640 896t90.5-37.5T768 768m768 768v-384q0-52-38-90t-90-38t-90 38t-38 90v384q0 52 38 90t90 38t90-38t38-90m-384-768q0-53-37.5-90.5T1024 640t-90.5 37.5T896 768t37.5 90.5T1024 896t90.5-37.5T1152 768m384-320V192q0-26-19-45t-45-19H192q-26 0-45 19t-19 45v256q0 26 19 45t45 19h1280q26 0 45-19t19-45m0 320q0-53-37.5-90.5T1408 640t-90.5 37.5T1280 768t37.5 90.5T1408 896t90.5-37.5T1536 768m128-640v1536q0 52-38 90t-90 38H128q-52 0-90-38t-38-90V128q0-52 38-90t90-38h1408q52 0 90 38t38 90"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 1664h288v-288H128zm352 0h320v-288H480zm-352-352h288V992H128zm352 0h320V992H480zM128 928h288V640H128zm736 736h320v-288H864zM480 928h320V640H480zm768 736h288v-288h-288zm-384-352h320V992H864zM512 448V160q0-13-9.5-22.5T480 128h-64q-13 0-22.5 9.5T384 160v288q0 13 9.5 22.5T416 480h64q13 0 22.5-9.5T512 448m736 864h288V992h-288zM864 928h320V640H864zm384 0h288V640h-288zm32-480V160q0-13-9.5-22.5T1248 128h-64q-13 0-22.5 9.5T1152 160v288q0 13 9.5 22.5t22.5 9.5h64q13 0 22.5-9.5t9.5-22.5m384-64v1280q0 52-38 90t-90 38H128q-52 0-90-38t-38-90V384q0-52 38-90t90-38h128v-96q0-66 47-113T416 0h64q66 0 113 47t47 113v96h384v-96q0-66 47-113t113-47h64q66 0 113 47t47 113v96h128q52 0 90 38t38 90"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarCheckO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1303 964l-512 512q-10 9-23 9t-23-9l-288-288q-9-10-9-23t9-22l46-46q9-9 22-9t23 9l220 220l444-444q10-9 23-9t22 9l46 46q9 9 9 22t-9 23M128 1664h1408V640H128zM512 448V160q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v288q0 14 9 23t23 9h64q14 0 23-9t9-23m768 0V160q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v288q0 14 9 23t23 9h64q14 0 23-9t9-23m384-64v1280q0 52-38 90t-90 38H128q-52 0-90-38t-38-90V384q0-52 38-90t90-38h128v-96q0-66 47-113T416 0h64q66 0 113 47t47 113v96h384v-96q0-66 47-113t113-47h64q66 0 113 47t47 113v96h128q52 0 90 38t38 90"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarMinusO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1152 1120v64q0 14-9 23t-23 9H544q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h576q14 0 23 9t9 23M128 1664h1408V640H128zM512 448V160q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v288q0 14 9 23t23 9h64q14 0 23-9t9-23m768 0V160q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v288q0 14 9 23t23 9h64q14 0 23-9t9-23m384-64v1280q0 52-38 90t-90 38H128q-52 0-90-38t-38-90V384q0-52 38-90t90-38h128v-96q0-66 47-113T416 0h64q66 0 113 47t47 113v96h384v-96q0-66 47-113t113-47h64q66 0 113 47t47 113v96h128q52 0 90 38t38 90"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 1664h1408V640H128zM512 448V160q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v288q0 14 9 23t23 9h64q14 0 23-9t9-23m768 0V160q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v288q0 14 9 23t23 9h64q14 0 23-9t9-23m384-64v1280q0 52-38 90t-90 38H128q-52 0-90-38t-38-90V384q0-52 38-90t90-38h128v-96q0-66 47-113T416 0h64q66 0 113 47t47 113v96h384v-96q0-66 47-113t113-47h64q66 0 113 47t47 113v96h128q52 0 90 38t38 90"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarPlusO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1536 256q52 0 90 38t38 90v1280q0 52-38 90t-90 38H128q-52 0-90-38t-38-90V384q0-52 38-90t90-38h128v-96q0-66 47-113T416 0h64q66 0 113 47t47 113v96h384v-96q0-66 47-113t113-47h64q66 0 113 47t47 113v96zm-384-96v288q0 14 9 23t23 9h64q14 0 23-9t9-23V160q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23m-768 0v288q0 14 9 23t23 9h64q14 0 23-9t9-23V160q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23m1152 1504V640H128v1024zm-640-576h224q14 0 23 9t9 23v64q0 14-9 23t-23 9H896v224q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23v-224H544q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h224V864q0-14 9-23t23-9h64q14 0 23 9t9 23z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarTimesO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1111 1385l-46 46q-9 9-22 9t-23-9l-188-189l-188 189q-10 9-23 9t-22-9l-46-46q-9-9-9-22t9-23l189-188l-189-188q-9-10-9-23t9-22l46-46q9-9 22-9t23 9l188 188l188-188q10-9 23-9t22 9l46 46q9 9 9 22t-9 23l-188 188l188 188q9 10 9 23t-9 22m-983 279h1408V640H128zM512 448V160q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v288q0 14 9 23t23 9h64q14 0 23-9t9-23m768 0V160q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v288q0 14 9 23t23 9h64q14 0 23-9t9-23m384-64v1280q0 52-38 90t-90 38H128q-52 0-90-38t-38-90V384q0-52 38-90t90-38h128v-96q0-66 47-113T416 0h64q66 0 113 47t47 113v96h384v-96q0-66 47-113t113-47h64q66 0 113 47t47 113v96h128q52 0 90 38t38 90"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M960 672q119 0 203.5 84.5T1248 960t-84.5 203.5T960 1248t-203.5-84.5T672 960t84.5-203.5T960 672m704-416q106 0 181 75t75 181v896q0 106-75 181t-181 75H256q-106 0-181-75T0 1408V512q0-106 75-181t181-75h224l51-136q19-49 69.5-84.5T704 0h512q53 0 103.5 35.5T1389 120l51 136zM960 1408q185 0 316.5-131.5T1408 960t-131.5-316.5T960 512T643.5 643.5T512 960t131.5 316.5T960 1408"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraRetro(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M928 704q0-14-9-23t-23-9q-66 0-113 47t-47 113q0 14 9 23t23 9t23-9t9-23q0-40 28-68t68-28q14 0 23-9t9-23m224 130q0 106-75 181t-181 75t-181-75t-75-181t75-181t181-75t181 75t75 181M128 1408h1536v-128H128zm1152-574q0-159-112.5-271.5T896 450T624.5 562.5T512 834t112.5 271.5T896 1218t271.5-112.5T1280 834M256 192h384V64H256zM128 384h1536V128H836l-64 128H128zm1664-256v1280q0 53-37.5 90.5T1664 1536H128q-53 0-90.5-37.5T0 1408V128q0-53 37.5-90.5T128 0h1536q53 0 90.5 37.5T1792 128"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretDown(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1024 448q0 26-19 45L557 941q-19 19-45 19t-45-19L19 493Q0 474 0 448t19-45t45-19h896q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretLeft(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M576 192v896q0 26-19 45t-45 19t-45-19L19 685Q0 666 0 640t19-45l448-448q19-19 45-19t45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretSquareOleft(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1024 448v640q0 26-19 45t-45 19q-20 0-37-12L475 820q-27-19-27-52t27-52l448-320q17-12 37-12q26 0 45 19t19 45m256 800V288q0-13-9.5-22.5T1248 256H288q-13 0-22.5 9.5T256 288v960q0 13 9.5 22.5t22.5 9.5h960q13 0 22.5-9.5t9.5-22.5m256-960v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretSquareOup(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1145 989q-17 35-57 35H448q-40 0-57-35q-18-35 5-66l320-448q19-27 52-27t52 27l320 448q23 31 5 66m135 259V288q0-13-9.5-22.5T1248 256H288q-13 0-22.5 9.5T256 288v960q0 13 9.5 22.5t22.5 9.5h960q13 0 22.5-9.5t9.5-22.5m256-960v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretUp(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1024 960q0 26-19 45t-45 19H64q-26 0-45-19T0 960t19-45l448-448q19-19 45-19t45 19l448 448q19 19 19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CartArrowDown(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1280 448q0-26-19-45t-45-19t-45 19l-147 146V256q0-26-19-45t-45-19t-45 19t-19 45v293L749 403q-19-19-45-19t-45 19t-19 45t19 45l256 256q19 19 45 19t45-19l256-256q19-19 19-45m-640 832q0 53-37.5 90.5T512 1408t-90.5-37.5T384 1280t37.5-90.5T512 1152t90.5 37.5T640 1280m896 0q0 53-37.5 90.5T1408 1408t-90.5-37.5t-37.5-90.5t37.5-90.5t90.5-37.5t90.5 37.5t37.5 90.5m128-1088v512q0 24-16 42.5t-41 21.5L563 890q1 7 4.5 21.5t6 26.5t2.5 22q0 16-24 64h920q26 0 45 19t19 45t-19 45t-45 19H448q-26 0-45-19t-19-45q0-14 11-39.5t29.5-59.5t20.5-38L268 128H64q-26 0-45-19T0 64t19-45T64 0h256q16 0 28.5 6.5t20 15.5t13 24.5T389 73t5.5 29.5T399 128h1201q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CartPlus(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1216 448q0-26-19-45t-45-19h-128V256q0-26-19-45t-45-19t-45 19t-19 45v128H768q-26 0-45 19t-19 45t19 45t45 19h128v128q0 26 19 45t45 19t45-19t19-45V512h128q26 0 45-19t19-45m-576 832q0 53-37.5 90.5T512 1408t-90.5-37.5T384 1280t37.5-90.5T512 1152t90.5 37.5T640 1280m896 0q0 53-37.5 90.5T1408 1408t-90.5-37.5t-37.5-90.5t37.5-90.5t90.5-37.5t90.5 37.5t37.5 90.5m128-1088v512q0 24-16 42.5t-41 21.5L563 890q1 7 4.5 21.5t6 26.5t2.5 22q0 16-24 64h920q26 0 45 19t19 45t-19 45t-45 19H448q-26 0-45-19t-19-45q0-14 11-39.5t29.5-59.5t20.5-38L268 128H64q-26 0-45-19T0 64t19-45T64 0h256q16 0 28.5 6.5t20 15.5t13 24.5T389 73t5.5 29.5T399 128h1201q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cc(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M785 880h207q-14 158-98.5 248.5T679 1219q-162 0-254.5-116T332 787q0-194 93-311.5T658 358q148 0 232 87t97 247H784q-5-64-35.5-99T667 558q-57 0-88.5 60.5T547 796q0 48 5 84t18 69.5t40 51.5t66 18q95 0 109-139m712 0h206q-14 158-98 248.5t-214 90.5q-162 0-254.5-116T1044 787q0-194 93-311.5T1370 358q148 0 232 87t97 247h-204q-4-64-35-99t-81-35q-57 0-88.5 60.5T1259 796q0 48 5 84t18 69.5t39.5 51.5t65.5 18q49 0 76.5-38t33.5-101m359-119q0-207-15.5-307T1780 293q-6-8-13.5-14t-21.5-15t-16-11q-86-63-697-63q-625 0-710 63q-5 4-17.5 11.5t-21 14T269 293q-45 60-60 159.5T194 761q0 208 15 307.5t60 160.5q6 8 15 15t20.5 14t17.5 12q44 33 239.5 49t470.5 16q610 0 697-65q5-4 17-11t20.5-14t13.5-16q46-60 61-159t15-309M2048 0v1536H0V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CcAmex(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M119 554h89l-45-108zm621 526l74-79l-70-79H581v49h142v55H581v54zm158-78l99 110V895zm288-47q0-33-40-33h-84v69h83q41 0 41-36m289-4q0-29-42-29h-82v61h81q43 0 43-32m-278-466q0-29-42-29h-82v60h81q43 0 43-31m459 69h89l-44-108zM699 399v271h-66V458l-94 212h-57l-94-212v212H256l-25-60H96l-25 60H1l116-271h96l110 257V399h106l85 184l77-184zm556 556q0 20-5.5 35t-14 25t-22.5 16.5t-26 10t-31.5 4.5t-31.5 1t-32.5-.5t-29.5-.5v91H936l-80-90l-83 90H517V866h260l80 89l82-89h207q109 0 109 89M964 614v56H747V399h217v57H812v49h148v55H812v54zm1340 559v229q0 55-38.5 94.5T2172 1536H132q-55 0-93.5-39.5T0 1402V724h111l25-61h55l25 61h218v-46l19 46h113l20-47v47h541v-99l10-1q10 0 10 14v86h279v-23q23 12 55 18t52.5 6.5t63-.5t51.5-1l25-61h56l25 61h227v-58l34 58h182V346h-180v44l-25-44h-185v44l-23-44h-249q-69 0-109 22v-22h-172v22q-24-22-73-22H553l-43 97l-43-97H269v44l-22-44H78L0 525V134q0-55 38.5-94.5T132 0h2040q55 0 93.5 39.5T2304 134v678h-120q-51 0-81 22v-22h-177q-55 0-78 22v-22h-316v22q-31-22-87-22h-209v22q-23-22-91-22H911l-54 58l-50-58H458v378h343l55-59l52 59h211v-89h21q59 0 90-13v102h174v-99h8q8 0 10 2t2 10v87h529q57 0 88-24v24h168q60 0 95-17m-758-234q0 23-12 43t-34 29q25 9 34 26t9 46v54h-65v-45q0-33-12-43.5t-46-10.5h-69v99h-65V866h154q48 0 77 15t29 58m-277-467q0 24-12.5 44t-33.5 29q26 9 34.5 25.5t8.5 46.5v53h-65q0-9 .5-26.5t0-25t-3-18.5t-8.5-16t-17.5-8.5t-29.5-3.5h-70v98h-64V399l153 1q49 0 78 14.5t29 57.5m529 609v56h-216V866h216v56h-151v49h148v55h-148v54zm-426-682v271h-66V399zm693 652q0 86-102 86h-126v-58h126q34 0 34-25q0-16-17-21t-41.5-5t-49.5-3.5t-42-22.5t-17-55q0-39 26-60t66-21h130v57h-119q-36 0-36 25q0 16 17.5 20.5t42 4t49 2.5t42 21.5t17.5 54.5m239-50v101q-24 35-88 35h-125v-58h125q33 0 33-25q0-13-12.5-19t-31-5.5t-40-2t-40-8t-31-24T2082 947q0-39 26.5-60t66.5-21h129v57h-118q-36 0-36 25q0 20 29 22t68.5 5t56.5 26m-165-601v270h-92l-122-203v203h-132l-26-60h-134l-25 60h-75q-129 0-129-133q0-138 133-138h63v59q-7 0-28-1t-28.5-.5t-23 2t-21.5 6.5t-14.5 13.5t-11.5 23t-3 33.5q0 38 13.5 58t49.5 20h29l92-213h97l109 256V400h99l114 188V400z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CcDinersClub(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M858 1113V420q-106 41-172 135.5T620 767t66 211.5T858 1113m504-346q0-117-66-211.5T1124 420v694q106-41 172-135.5t66-211.5m215 0q0 159-78.5 294T1285 1274.5T991 1353q-119 0-227.5-46.5t-187-125t-125-187T405 767q0-159 78.5-294T697 259.5T991 181t294 78.5T1498.5 473t78.5 294m383 7q0-139-55.5-261.5T1757 307t-213.5-131t-252.5-48H990q-176 0-323.5 81t-235 230T344 774q0 171 87 317.5T667 1323t323 85h301q129 0 251.5-50.5t214.5-135t147.5-202.5t55.5-246m344-646v1280q0 52-38 90t-90 38H128q-52 0-90-38t-38-90V128q0-52 38-90t90-38h2048q52 0 90 38t38 90"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CcDiscover(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M313 649q0 51-36 84q-29 26-89 26h-17V539h17q61 0 89 27q36 31 36 83m1776-65q0 52-64 52h-19V535h20q63 0 63 49M380 649q0-74-50-120.5T201 482h-95v333h95q74 0 119-38q60-51 60-128m30 166h65V482h-65zm320-101q0-40-20.5-62T634 610q-29-10-39.5-19T584 568q0-16 13.5-26.5T632 531q29 0 53 27l34-44q-41-37-98-37q-44 0-74 27.5T517 572q0 35 18 55.5t64 36.5q37 13 45 19q19 12 19 34q0 20-14 33.5T613 764q-48 0-71-44l-42 40q44 64 115 64q51 0 83-30.5t32-79.5m278 90v-77q-37 37-78 37q-49 0-80.5-32.5T818 649q0-48 31.5-81.5T927 534q43 0 81 38v-77q-40-20-80-20q-74 0-125.5 50.5T751 649t51 123.5T927 823q42 0 81-19m1232 604V881q-65 40-144.5 84T1858 1082t-329.5 137.5T1111 1354t-504 118h1569q26 0 45-19t19-45m-851-757q0-75-53-128t-128-53t-128 53t-53 128t53 128t128 53t128-53t53-128m152 173l144-342h-71l-90 224l-89-224h-71l142 342zm173-9h184v-56h-119v-90h115v-56h-115v-74h119v-57h-184zm391 0h80l-105-140q76-16 76-94q0-47-31-73t-87-26h-97v333h65V682h9zm199-681v1268q0 56-38.5 95t-93.5 39H132q-55 0-93.5-39T0 1402V134q0-56 38.5-95T132 0h2040q55 0 93.5 39t38.5 95"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CcJcb(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1951 870q0 26-15.5 44.5T1897 938q-8 2-18 2h-153V800h153q10 0 18 2q23 5 38.5 23.5T1951 870m-18-213q0 25-15 42t-38 21q-3 1-15 1h-139V592h139q3 0 8.5.5t6.5.5q23 4 38 21.5t15 42.5M728 821V513H500v308q0 58-38 94.5T357 952q-108 0-229-59v112q53 15 121 23t109 9l42 1q328 0 328-217m714 184V892q-99 52-200 59q-108 8-169-41t-61-142t61-142t169-41q101 7 200 58V531q-48-12-100-19.5t-80-9.5l-28-2q-127-6-218.5 14T875 574t-71 88t-22 106t22 106t71 88t140.5 60t218.5 14q101-4 208-31m734-115q0-54-43-88.5T2024 762v-3q57-8 89-41.5t32-79.5q0-55-41-88t-107-36q-3 0-12-.5t-14-.5h-455v510h491q74 0 121.5-36.5T2176 890m128-762v1280q0 52-38 90t-90 38H128q-52 0-90-38t-38-90V128q0-52 38-90t90-38h2048q52 0 90 38t38 90"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CcMastercard(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1119 213q-128-85-281-85q-103 0-197.5 40.5T478 277T369.5 439T329 636q0 104 40.5 198T478 996t162 108.5t198 40.5q153 0 281-85q-131-107-178-265.5t.5-316.5T1119 213m33 24q-126 99-172 249.5t-.5 300.5t172.5 249q127-99 172.5-249t-.5-300.5T1152 237m33-24q130 107 177.5 265.5t.5 317t-178 264.5q128 85 281 85q104 0 198-40.5T1826 996t108.5-162t40.5-198q0-103-40.5-197T1826 277t-162.5-108.5T1466 128q-153 0-281 85m741 722h7v-3h-17v3h7v17h3zm29 17h4v-20h-5l-6 13l-6-13h-5v20h3v-15l6 13h4l5-13zm-8 440v2h-5v-3h5zm0 9h3l-4-5h2l1-1q1-1 1-3t-1-3l-1-1h-9v13h3v-5h1zm-1262-68q0-19 11-31t30-12q18 0 29 12.5t11 30.5q0 19-11 31t-29 12q-19 0-30-12t-11-31m473-44q30 0 35 32h-70q5-32 35-32m356 44q0-19 11-31t29-12t29.5 12.5t11.5 30.5q0 19-11 31t-30 12q-18 0-29-12t-11-31m272 0q0-18 11.5-30.5t29.5-12.5t29.5 12.5t11.5 30.5q0 19-11.5 31t-29.5 12t-29.5-12.5t-11.5-30.5m158 72q-2 0-4-1q-1 0-3-2t-2-3q-1-2-1-4q0-3 1-4q0-2 2-4l1-1q2 0 2-1q2-1 4-1q3 0 4 1l4 2l2 4v1q1 2 1 3l-1 1v3l-1 1l-1 2q-2 2-4 2q-1 1-4 1m-1345-4h30v-85q0-24-14.5-38.5T575 1262q-32 0-47 24q-14-24-45-24q-24 0-39 20v-16h-30v135h30v-75q0-36 33-36q30 0 30 36v75h29v-75q0-36 33-36q30 0 30 36zm166 0h29v-135h-29v16q-17-20-43-20q-29 0-48 20t-19 51t19 51t48 20q28 0 43-20zm178-41q0-34-47-40l-14-2q-23-4-23-14q0-15 25-15q23 0 43 11l12-24q-22-14-55-14q-26 0-41 12t-15 32q0 33 47 39l13 2q24 4 24 14q0 17-31 17q-25 0-45-14l-13 23q25 17 58 17q29 0 45.5-12t16.5-32m130 34l-8-25q-13 7-26 7q-19 0-19-22v-61h48v-27h-48v-41h-30v41h-28v27h28v61q0 50 47 50q21 0 36-10m86-132q-29 0-48 20t-19 51q0 32 19.5 51.5t49.5 19.5q33 0 55-19l-14-22q-18 15-39 15q-34 0-41-33h101v-12q0-32-18-51.5t-46-19.5m159 0q-23 0-35 20v-16h-30v135h30v-76q0-35 29-35q10 0 18 4l9-28q-9-4-21-4m30 71q0 31 19.5 51t52.5 20q29 0 48-16l-14-24q-18 13-35 12q-18 0-29.5-12t-11.5-31t11.5-31t29.5-12q19 0 35 12l14-24q-20-16-48-16q-33 0-52.5 20t-19.5 51m245 68h30v-135h-30v16q-15-20-42-20q-29 0-48.5 20t-19.5 51t19.5 51t48.5 20q28 0 42-20zm133-139q-23 0-35 20v-16h-29v135h29v-76q0-35 29-35q10 0 18 4l9-28q-8-4-21-4m140 139h29v-190h-29v71q-15-20-43-20t-47.5 20.5t-19.5 50.5t19.5 50.5t47.5 20.5q29 0 43-20zm78-20l-2 1h-3q-2 1-4 3q-3 1-3 4q-1 2-1 6q0 3 1 5q0 2 3 4q2 2 4 3t5 1q4 0 6-1q0-1 2-2l2-1q1-1 3-4q1-2 1-5q0-4-1-6q-1-1-3-4q0-1-2-2l-2-1q-1 0-3-.5t-3-.5m360-1253v1280q0 52-38 90t-90 38H128q-52 0-90-38t-38-90V128q0-52 38-90t90-38h2048q52 0 90 38t38 90"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CcPaypal(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M745 778q0 37-25.5 61.5T657 864q-29 0-46.5-16T593 804q0-37 25-62.5t62-25.5q28 0 46.5 16.5T745 778m785-149q0 42-22 57t-66 15l-32 1l17-107q2-11 13-11h18q22 0 35 2t25 12.5t12 30.5m351 149q0 36-25.5 61t-61.5 25q-29 0-47-16t-18-44q0-37 25-62.5t62-25.5q28 0 46.5 16.5T1881 778M513 607q0-59-38.5-85.5T374 495H214q-19 0-21 19l-65 408q-1 6 3 11t10 5h76q20 0 22-19l18-110q1-8 7-13t15-6.5t17-1.5t19 1t14 1q86 0 135-48.5T513 607m309 312l41-261q1-6-3-11t-10-5h-76q-14 0-17 33q-27-40-95-40q-72 0-122.5 54T489 816q0 59 34.5 94t92.5 35q28 0 58-12t48-32q-4 12-4 21q0 16 13 16h69q19 0 22-19m447-263q0-5-4-9.5t-9-4.5h-77q-11 0-18 10l-106 156l-44-150q-5-16-22-16h-75q-5 0-9 4.5t-4 9.5q0 2 19.5 59t42 123t23.5 70q-82 112-82 120q0 13 13 13h77q11 0 18-10l255-368q2-2 2-7m380-49q0-59-38.5-85.5T1510 495h-159q-20 0-22 19l-65 408q-1 6 3 11t10 5h82q12 0 16-13l18-116q1-8 7-13t15-6.5t17-1.5t19 1t14 1q86 0 135-48.5t49-134.5m309 312l41-261q1-6-3-11t-10-5h-76q-14 0-17 33q-26-40-95-40q-72 0-122.5 54T1625 816q0 59 34.5 94t92.5 35q29 0 59-12t47-32q0 1-2 9t-2 12q0 16 13 16h69q19 0 22-19m218-409v-1q0-14-13-14h-74q-11 0-13 11l-65 416l-1 2q0 5 4 9.5t10 4.5h66q19 0 21-19zM392 644q-5 35-26 46t-60 11l-33 1l17-107q2-11 13-11h19q40 0 58 11.5t12 48.5m1912-516v1280q0 52-38 90t-90 38H128q-52 0-90-38t-38-90V128q0-52 38-90t90-38h2048q52 0 90 38t38 90"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CcStripe(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1597 775q0 69-21 106q-19 35-52 35q-23 0-41-9V683q29-30 57-30q57 0 57 122m438-36h-110q6-98 56-98q51 0 54 98M476 874q0-59-33-91.5T342 725q-36-13-52-24t-16-25q0-26 38-26q58 0 124 33l18-112q-67-32-149-32q-77 0-123 38q-48 39-48 109q0 58 32.5 90.5T266 833q39 14 54.5 25.5T336 886q0 31-48 31q-29 0-70-12.5T146 874l-18 113q72 41 168 41q81 0 129-37q51-41 51-117m295-215l19-111h-96V413l-129 21l-18 114l-46 8l-17 103h62v219q0 84 44 120q38 30 111 30q32 0 79-11V899q-32 7-44 7q-42 0-42-50V659zm316 25V545q-15-3-28-3q-32 0-55.5 16T970 604l-10-56H829v471h150V713q26-31 82-31q16 0 26 2m37 335h150V548h-150zm622-249q0-122-45-179q-40-52-111-52q-64 0-117 56l-8-47h-132v645l150-25v-151q36 11 68 11q83 0 134-56q61-65 61-202m-468-348q0-33-23-56t-56-23t-56 23t-23 56t23 56.5t56 23.5t56-23.5t23-56.5m898 357q0-113-48-176q-50-64-144-64q-96 0-151.5 66T1777 785q0 128 63 188q55 55 161 55q101 0 160-40l-16-103q-57 31-128 31q-43 0-63-19q-23-19-28-66h248q2-14 2-52m128-651v1280q0 52-38 90t-90 38H128q-52 0-90-38t-38-90V128q0-52 38-90t90-38h2048q52 0 90 38t38 90"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CcVisa(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1975 862h-138q14-37 66-179l3-9q4-10 10-26t9-26l12 55zM531 797l-58-295q-11-54-75-54H130l-2 13q311 79 403 336m179-349L548 886l-17-89q-26-70-85-129.5T315 579l135 510h175l261-641zm139 642h166l104-642H953zm768-626q-69-27-149-27q-123 0-201 59t-79 153q-1 102 145 174q48 23 67 41t19 39q0 30-30 46t-69 16q-86 0-156-33l-22-11l-23 144q74 34 185 34q130 1 208.5-59t80.5-160q0-106-140-174q-49-25-71-42t-22-38q0-22 24.5-38.5T1455 571q70-1 124 24l15 8zm425-16h-128q-65 0-87 54l-246 588h174l35-96h212q5 22 20 96h154zm262-320v1280q0 52-38 90t-90 38H128q-52 0-90-38t-38-90V128q0-52 38-90t90-38h2048q52 0 90 38t38 90"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Certificate(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1376 768l138 135q30 28 20 70q-12 41-52 51l-188 48l53 186q12 41-19 70q-29 31-70 19l-186-53l-48 188q-10 40-51 52q-12 2-19 2q-31 0-51-22l-135-138l-135 138q-28 30-70 20q-41-11-51-52l-48-188l-186 53q-41 12-70-19q-31-29-19-70l53-186l-188-48q-40-10-52-51q-10-42 20-70l138-135L22 633q-30-28-20-70q12-41 52-51l188-48l-53-186q-12-41 19-70q29-31 70-19l186 53l48-188q10-41 51-51q41-12 70 19l135 139L903 22q29-30 70-19q41 10 51 51l48 188l186-53q41-12 70 19q31 29 19 70l-53 186l188 48q40 10 52 51q10 42-20 70z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chain(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1456 1216q0-40-28-68l-208-208q-28-28-68-28q-42 0-72 32q3 3 19 18.5t21.5 21.5t15 19t13 25.5t3.5 27.5q0 40-28 68t-68 28q-15 0-27.5-3.5t-25.5-13t-19-15t-21.5-21.5t-18.5-19q-33 31-33 73q0 40 28 68l206 207q27 27 68 27q40 0 68-26l147-146q28-28 28-67M753 511q0-40-28-68L519 236q-28-28-68-28q-39 0-68 27L236 381q-28 28-28 67q0 40 28 68l208 208q27 27 68 27q42 0 72-31q-3-3-19-18.5T543.5 680t-15-19t-13-25.5T512 608q0-40 28-68t68-28q15 0 27.5 3.5t25.5 13t19 15t21.5 21.5t18.5 19q33-31 33-73m895 705q0 120-85 203l-147 146q-83 83-203 83q-121 0-204-85l-206-207q-83-83-83-203q0-123 88-209l-88-88q-86 88-208 88q-120 0-204-84L100 652q-84-84-84-204t85-203L248 99q83-83 203-83q121 0 204 85l206 207q83 83 83 203q0 123-88 209l88 88q86-88 208-88q120 0 204 84l208 208q84 84 84 204"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChainBroken(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m439 1271l-256 256q-11 9-23 9t-23-9q-9-10-9-23t9-23l256-256q10-9 23-9t23 9q9 10 9 23t-9 23m169 41v320q0 14-9 23t-23 9t-23-9t-9-23v-320q0-14 9-23t23-9t23 9t9 23m-224-224q0 14-9 23t-23 9H32q-14 0-23-9t-9-23t9-23t23-9h320q14 0 23 9t9 23m1264 128q0 120-85 203l-147 146q-83 83-203 83q-121 0-204-85l-334-335q-21-21-42-56l239-18l273 274q27 27 68 27.5t68-26.5l147-146q28-28 28-67q0-40-28-68l-274-275l18-239q35 21 56 42l336 336q84 86 84 204m-617-724l-239 18l-273-274q-28-28-68-28q-39 0-68 27L236 381q-28 28-28 67q0 40 28 68l274 274l-18 240q-35-21-56-42L100 652q-84-86-84-204q0-120 85-203L248 99q83-83 203-83q121 0 204 85l334 335q21 21 42 56m633 84q0 14-9 23t-23 9h-320q-14 0-23-9t-9-23t9-23t23-9h320q14 0 23 9t9 23M1120 32v320q0 14-9 23t-23 9t-23-9t-9-23V32q0-14 9-23t23-9t23 9t9 23m407 151l-256 256q-11 9-23 9t-23-9q-9-10-9-23t9-23l256-256q10-9 23-9t23 9q9 10 9 23t-9 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1575 310q0 40-28 68l-724 724l-136 136q-28 28-68 28t-68-28l-136-136L53 740q-28-28-28-68t28-68l136-136q28-28 68-28t68 28l294 295l656-657q28-28 68-28t68 28l136 136q28 28 28 68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircle(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1284 606q0-28-18-46l-91-90q-19-19-45-19t-45 19L677 877L451 651q-19-19-45-19t-45 19l-91 90q-18 18-18 46q0 27 18 45l362 362q19 19 45 19q27 0 46-19l543-543q18-18 18-45m252 162q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircleO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1171 685l-422 422q-19 19-45 19t-45-19L365 813q-19-19-19-45t19-45l102-102q19-19 45-19t45 19l147 147l275-275q19-19 45-19t45 19l102 102q19 19 19 45t-19 45m141 83q0-148-73-273t-198-198t-273-73t-273 73t-198 198t-73 273t73 273t198 198t273 73t273-73t198-198t73-273m224 0q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m685 1171l614-614q19-19 19-45t-19-45l-102-102q-19-19-45-19t-45 19L640 832L429 621q-19-19-45-19t-45 19L237 723q-19 19-19 45t19 45l358 358q19 19 45 19t45-19m851-883v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckSquareO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1408 802v318q0 119-84.5 203.5T1120 1408H288q-119 0-203.5-84.5T0 1120V288Q0 169 84.5 84.5T288 0h832q63 0 117 25q15 7 18 23q3 17-9 29l-49 49q-10 10-23 10q-3 0-9-2q-23-6-45-6H288q-66 0-113 47t-47 113v832q0 66 47 113t113 47h832q66 0 113-47t47-113V866q0-13 9-22l64-64q10-10 23-10q6 0 12 3q20 8 20 29m231-489l-814 814q-24 24-57 24t-57-24L281 697q-24-24-24-57t24-57l110-110q24-24 57-24t57 24l263 263l647-647q24-24 57-24t57 24l110 110q24 24 24 57t-24 57"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronCircleLeft(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m909 1267l102-102q19-19 19-45t-19-45L704 768l307-307q19-19 19-45t-19-45L909 269q-19-19-45-19t-45 19L365 723q-19 19-19 45t19 45l454 454q19 19 45 19t45-19m627-499q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronCircleUp(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1165 1011l102-102q19-19 19-45t-19-45L813 365q-19-19-45-19t-45 19L269 819q-19 19-19 45t19 45l102 102q19 19 45 19t45-19l307-307l307 307q19 19 45 19t45-19m371-243q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDown(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1619 552l-742 741q-19 19-45 19t-45-19L45 552q-19-19-19-45.5T45 461l166-165q19-19 45-19t45 19l531 531l531-531q19-19 45-19t45 19l166 165q19 19 19 45.5t-19 45.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeft(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1043 301L512 832l531 531q19 19 19 45t-19 45l-166 166q-19 19-45 19t-45-19L45 877q-19-19-19-45t19-45L787 45q19-19 45-19t45 19l166 166q19 19 19 45t-19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUp(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1619 1075l-166 165q-19 19-45 19t-45-19L832 709l-531 531q-19 19-45 19t-45-19L45 1075q-19-19-19-45.5T45 984l742-741q19-19 45-19t45 19l742 741q19 19 19 45.5t-19 45.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Child(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1124 388L832 680v824q0 46-33 79t-79 33t-79-33t-33-79v-384h-64v384q0 46-33 79t-79 33t-79-33t-33-79V680L28 388Q0 360 0 320t28-68q29-28 68.5-28t67.5 28l228 228h368l228-228q28-28 68-28t68 28q28 29 28 68.5t-28 67.5M800 224q0 93-65.5 158.5T576 448t-158.5-65.5T352 224t65.5-158.5T576 0t158.5 65.5T800 224"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chrome(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M893 0q240-2 451 120q232 134 352 372l-742-39q-160-9-294 74.5T475 757L199 333Q327 174 510 87.5T893 0M146 405l337 663q72 143 211 217t293 45l-230 451q-212-33-385-157.5t-272.5-316T0 896q0-267 146-491m1586 169q58 150 59.5 310.5t-48.5 306t-153 272t-246 209.5q-230 133-498 119l405-623q88-131 82.5-290.5T1227 600zm-836 20q125 0 213.5 88.5T1198 896t-88.5 213.5T896 1198t-213.5-88.5T594 896t88.5-213.5T896 594"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circle(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1536 768q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 224q-148 0-273 73T297 495t-73 273t73 273t198 198t273 73t273-73t198-198t73-273t-73-273t-198-198t-273-73m768 544q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleOnotch(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1728 864q0 176-68.5 336t-184 275.5t-275.5 184t-336 68.5t-336-68.5t-275.5-184t-184-275.5T0 864q0-213 97-398.5T362 160T736 9v228q-221 45-366.5 221T224 864q0 130 51 248.5t136.5 204t204 136.5t248.5 51t248.5-51t204-136.5t136.5-204t51-248.5q0-230-145.5-406T992 237V9q206 31 374 151t265 305.5t97 398.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleThin(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 128q-130 0-248.5 51t-204 136.5t-136.5 204T128 768t51 248.5t136.5 204t204 136.5t248.5 51t248.5-51t204-136.5t136.5-204t51-248.5t-51-248.5t-136.5-204t-204-136.5T768 128m768 640q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clipboard(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 1664h896v-640h-416q-40 0-68-28t-28-68V512H768zm256-1440v-64q0-13-9.5-22.5T992 128H288q-13 0-22.5 9.5T256 160v64q0 13 9.5 22.5T288 256h704q13 0 22.5-9.5t9.5-22.5m256 672h299l-299-299zm512 128v672q0 40-28 68t-68 28H736q-40 0-68-28t-28-68v-160H96q-40 0-68-28t-28-68V96q0-40 28-68T96 0h1088q40 0 68 28t28 68v328q21 13 36 28l408 408q28 28 48 76t20 88"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M896 416v448q0 14-9 23t-23 9H544q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h224V416q0-14 9-23t23-9h64q14 0 23 9t9 23m416 352q0-148-73-273t-198-198t-273-73t-273 73t-198 198t-73 273t73 273t198 198t273 73t273-73t198-198t73-273m224 0q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clone(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1664 1632V544q0-13-9.5-22.5T1632 512H544q-13 0-22.5 9.5T512 544v1088q0 13 9.5 22.5t22.5 9.5h1088q13 0 22.5-9.5t9.5-22.5m128-1088v1088q0 66-47 113t-113 47H544q-66 0-113-47t-47-113V544q0-66 47-113t113-47h1088q66 0 113 47t47 113m-384-384v160h-128V160q0-13-9.5-22.5T1248 128H160q-13 0-22.5 9.5T128 160v1088q0 13 9.5 22.5t22.5 9.5h160v128H160q-66 0-113-47T0 1248V160Q0 94 47 47T160 0h1088q66 0 113 47t47 113"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Close(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1202 1066q0 40-28 68l-136 136q-28 28-68 28t-68-28L608 976l-294 294q-28 28-68 28t-68-28L42 1134q-28-28-28-68t28-68l294-294L42 410q-28-28-28-68t28-68l136-136q28-28 68-28t68 28l294 294l294-294q28-28 68-28t68 28l136 136q28 28 28 68t-28 68L880 704l294 294q28 28 28 68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1920 1024q0 159-112.5 271.5T1536 1408H448q-185 0-316.5-131.5T0 960q0-132 71-241.5T258 555q-2-28-2-43q0-212 150-362T768 0q158 0 286.5 88T1242 318q70-62 166-62q106 0 181 75t75 181q0 75-41 138q129 30 213 134.5t84 239.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDownload(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1280 800q0-14-9-23t-23-9h-224V416q0-13-9.5-22.5T992 384H800q-13 0-22.5 9.5T768 416v352H544q-13 0-22.5 9.5T512 800q0 14 9 23l352 352q9 9 23 9t23-9l351-351q10-12 10-24m640 224q0 159-112.5 271.5T1536 1408H448q-185 0-316.5-131.5T0 960q0-130 70-240t188-165q-2-30-2-43q0-212 150-362T768 0q156 0 285.5 87T1242 318q71-62 166-62q106 0 181 75t75 181q0 76-41 138q130 31 213.5 135.5T1920 1024"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUpload(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1280 736q0-14-9-23L919 361q-9-9-23-9t-23 9L522 712q-10 12-10 24q0 14 9 23t23 9h224v352q0 13 9.5 22.5t22.5 9.5h192q13 0 22.5-9.5t9.5-22.5V768h224q13 0 22.5-9.5t9.5-22.5m640 288q0 159-112.5 271.5T1536 1408H448q-185 0-316.5-131.5T0 960q0-130 70-240t188-165q-2-30-2-43q0-212 150-362T768 0q156 0 285.5 87T1242 318q71-62 166-62q106 0 181 75t75 181q0 76-41 138q130 31 213.5 135.5T1920 1024"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cny(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M603 1408H431q-13 0-22.5-9t-9.5-23v-330H111q-13 0-22.5-9t-9.5-23V911q0-13 9.5-22.5T111 879h288v-85H111q-13 0-22.5-9T79 762V658q0-13 9.5-22.5T111 626h214L4 48q-8-16 0-32Q14 0 32 0h194q19 0 29 18l215 425q19 38 56 125q10-24 30.5-68t27.5-61L775 19q8-19 29-19h191q17 0 27 16q9 14 1 31L710 626h215q13 0 22.5 9.5T957 658v104q0 14-9.5 23t-22.5 9H635v85h290q13 0 22.5 9.5T957 911v103q0 14-9.5 23t-22.5 9H635v330q0 13-9.5 22.5T603 1408"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Code(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m585 1143l-50 50q-10 10-23 10t-23-10L23 727q-10-10-10-23t10-23l466-466q10-10 23-10t23 10l50 50q10 10 10 23t-10 23L192 704l393 393q10 10 10 23t-10 23M1176 76L803 1367q-4 13-15.5 19.5T764 1389l-62-17q-13-4-19.5-15.5T680 1332L1053 41q4-13 15.5-19.5T1092 19l62 17q13 4 19.5 15.5T1176 76m657 651l-466 466q-10 10-23 10t-23-10l-50-50q-10-10-10-23t10-23l393-393l-393-393q-10-10-10-23t10-23l50-50q10-10 23-10t23 10l466 466q10 10 10 23t-10 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeFork(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M288 1344q0-40-28-68t-68-28t-68 28t-28 68t28 68t68 28t68-28t28-68m0-1152q0-40-28-68t-68-28t-68 28t-28 68t28 68t68 28t68-28t28-68m640 128q0-40-28-68t-68-28t-68 28t-28 68t28 68t68 28t68-28t28-68m96 0q0 52-26 96.5T928 486q-2 287-226 414q-67 38-203 81q-128 40-169.5 71T288 1152v26q44 25 70 69.5t26 96.5q0 80-56 136t-136 56t-136-56t-56-136q0-52 26-96.5t70-69.5V358q-44-25-70-69.5T0 192q0-80 56-136T192 0t136 56t56 136q0 52-26 96.5T288 358v497q54-26 154-57q55-17 87.5-29.5t70.5-31t59-39.5t40.5-51t28-69.5T736 486q-44-25-70-69.5T640 320q0-80 56-136t136-56t136 56t56 136"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Codepen(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m216 1169l603 402v-359L485 989zm-62-144l193-129l-193-129zm819 546l603-402l-269-180l-334 223zm-77-493l272-182l-272-182l-272 182zM485 803l334-223V221L216 623zm960 93l193 129V767zm-138-93l269-180l-603-402v359zm485-180v546q0 41-34 64l-819 546q-21 13-43 13t-43-13L34 1233q-34-23-34-64V623q0-41 34-64L853 13q21-13 43-13t43 13l819 546q34 23 34 64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Codiepie(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1584 1290l-218-111q-74 120-196.5 189T906 1437q-147 0-271-72t-196-196t-72-270q0-110 42.5-209.5t115-172t172-115T906 360q131 0 247.5 60.5T1346 589l215-125q-110-169-286.5-265T896 103q-161 0-308 63T335 335T166 588t-63 308t63 308t169 253t253 169t308 63q213 0 397.5-107t290.5-292m-554-397l693 352q-116 253-334.5 400T896 1792q-182 0-348-71t-286-191t-191-286T0 896t71-348t191-286T548 71T896 0q260 0 470.5 133.5T1702 500zm513 3h-39v160h-96V704h136q32 0 54.5 20t28.5 48t1 56t-27.5 48t-57.5 20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coffee(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1664 384q0-80-56-136t-136-56h-64v384h64q80 0 136-56t56-136M0 1152h1792q0 106-75 181t-181 75H256q-106 0-181-75T0 1152m1856-768q0 159-112.5 271.5T1472 768h-64v32q0 92-66 158t-158 66H480q-92 0-158-66t-66-158V64q0-26 19-45t45-19h1152q159 0 271.5 112.5T1856 384"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cog(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1024 768q0-106-75-181t-181-75t-181 75t-75 181t75 181t181 75t181-75t75-181m512-109v222q0 12-8 23t-20 13l-185 28q-19 54-39 91q35 50 107 138q10 12 10 25t-9 23q-27 37-99 108t-94 71q-12 0-26-9l-138-108q-44 23-91 38q-16 136-29 186q-7 28-36 28H657q-14 0-24.5-8.5T621 1506l-28-184q-49-16-90-37l-141 107q-10 9-25 9q-14 0-25-11q-126-114-165-168q-7-10-7-23q0-12 8-23q15-21 51-66.5t54-70.5q-27-50-41-99L29 913q-13-2-21-12.5T0 877V655q0-12 8-23t19-13l186-28q14-46 39-92q-40-57-107-138q-10-12-10-24q0-10 9-23q26-36 98.5-107.5T337 135q13 0 26 10l138 107q44-23 91-38q16-136 29-186q7-28 36-28h222q14 0 24.5 8.5T915 30l28 184q49 16 90 37l142-107q9-9 24-9q13 0 25 10q129 119 165 170q7 8 7 22q0 12-8 23q-15 21-51 66.5t-54 70.5q26 50 41 98l183 28q13 2 21 12.5t8 23.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cogs(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M896 896q0-106-75-181t-181-75t-181 75t-75 181t75 181t181 75t181-75t75-181m768 512q0-52-38-90t-90-38t-90 38t-38 90q0 53 37.5 90.5t90.5 37.5t90.5-37.5t37.5-90.5m0-1024q0-52-38-90t-90-38t-90 38t-38 90q0 53 37.5 90.5T1536 512t90.5-37.5T1664 384m-384 421v185q0 10-7 19.5t-16 10.5l-155 24q-11 35-32 76q34 48 90 115q7 11 7 20q0 12-7 19q-23 30-82.5 89.5T999 1423q-11 0-21-7l-115-90q-37 19-77 31q-11 108-23 155q-7 24-30 24H547q-11 0-20-7.5t-10-17.5l-23-153q-34-10-75-31l-118 89q-7 7-20 7q-11 0-21-8q-144-133-144-160q0-9 7-19q10-14 41-53t47-61q-23-44-35-82l-152-24q-10-1-17-9.5T0 987V802q0-10 7-19.5T23 772l155-24q11-35 32-76q-34-48-90-115q-7-11-7-20q0-12 7-20q22-30 82-89t79-59q11 0 21 7l115 90q34-18 77-32q11-108 23-154q7-24 30-24h186q11 0 20 7.5t10 17.5l23 153q34 10 75 31l118-89q8-7 20-7q11 0 21 8q144 133 144 160q0 8-7 19q-12 16-42 54t-45 60q23 48 34 82l152 23q10 2 17 10.5t7 19.5m640 533v140q0 16-149 31q-12 27-30 52q51 113 51 138q0 4-4 7q-122 71-124 71q-8 0-46-47t-52-68q-20 2-30 2t-30-2q-14 21-52 68t-46 47q-2 0-124-71q-4-3-4-7q0-25 51-138q-18-25-30-52q-149-15-149-31v-140q0-16 149-31q13-29 30-52q-51-113-51-138q0-4 4-7q4-2 35-20t59-34t30-16q8 0 46 46.5t52 67.5q20-2 30-2t30 2q51-71 92-112l6-2q4 0 124 70q4 3 4 7q0 25-51 138q17 23 30 52q149 15 149 31m0-1024v140q0 16-149 31q-12 27-30 52q51 113 51 138q0 4-4 7q-122 71-124 71q-8 0-46-47t-52-68q-20 2-30 2t-30-2q-14 21-52 68t-46 47q-2 0-124-71q-4-3-4-7q0-25 51-138q-18-25-30-52q-149-15-149-31V314q0-16 149-31q13-29 30-52q-51-113-51-138q0-4 4-7q4-2 35-20t59-34t30-16q8 0 46 46.5t52 67.5q20-2 30-2t30 2q51-71 92-112l6-2q4 0 124 70q4 3 4 7q0 25-51 138q17 23 30 52q149 15 149 31"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Columns(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M160 1408h608V256H128v1120q0 13 9.5 22.5t22.5 9.5m1376-32V256H896v1152h608q13 0 22.5-9.5t9.5-22.5m128-1216v1216q0 66-47 113t-113 47H160q-66 0-113-47T0 1376V160Q0 94 47 47T160 0h1344q66 0 113 47t47 113"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comment(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 640q0 174-120 321.5t-326 233t-450 85.5q-70 0-145-8q-198 175-460 242q-49 14-114 22q-17 2-30.5-9t-17.5-29v-1q-3-4-.5-12t2-10t4.5-9.5l6-9l7-8.5l8-9q7-8 31-34.5t34.5-38t31-39.5t32.5-51t27-59t26-76q-157-89-247.5-220T0 640q0-130 71-248.5T262 187T548 50.5T896 0q244 0 450 85.5t326 233T1792 640"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M896 128q-204 0-381.5 69.5T232.5 385T128 640q0 112 71.5 213.5T401 1029l87 50l-27 96q-24 91-70 172q152-63 275-171l43-38l57 6q69 8 130 8q204 0 381.5-69.5t282-187.5T1664 640t-104.5-255t-282-187.5T896 128m896 512q0 174-120 321.5t-326 233t-450 85.5q-70 0-145-8q-198 175-460 242q-49 14-114 22h-5q-15 0-27-10.5t-16-27.5v-1q-3-4-.5-12t2-10t4.5-9.5l6-9l7-8.5l8-9q7-8 31-34.5t34.5-38t31-39.5t32.5-51t27-59t26-76q-157-89-247.5-220T0 640q0-174 120-321.5t326-233T896 0t450 85.5t326 233T1792 640"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Commenting(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 640q0-53-37.5-90.5T512 512t-90.5 37.5T384 640t37.5 90.5T512 768t90.5-37.5T640 640m384 0q0-53-37.5-90.5T896 512t-90.5 37.5T768 640t37.5 90.5T896 768t90.5-37.5T1024 640m384 0q0-53-37.5-90.5T1280 512t-90.5 37.5T1152 640t37.5 90.5T1280 768t90.5-37.5T1408 640m384 0q0 174-120 321.5t-326 233t-450 85.5q-110 0-211-18q-173 173-435 229q-52 10-86 13q-12 1-22-6t-13-18q-4-15 20-37q5-5 23.5-21.5T198 1398t23.5-25.5t24-31.5t20.5-37t20-48t14.5-57.5T313 1126q-146-90-229.5-216.5T0 640q0-174 120-321.5t326-233T896 0t450 85.5t326 233T1792 640"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentingO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 640q0 53-37.5 90.5T512 768t-90.5-37.5T384 640t37.5-90.5T512 512t90.5 37.5T640 640m384 0q0 53-37.5 90.5T896 768t-90.5-37.5T768 640t37.5-90.5T896 512t90.5 37.5T1024 640m384 0q0 53-37.5 90.5T1280 768t-90.5-37.5T1152 640t37.5-90.5T1280 512t90.5 37.5T1408 640M896 128q-204 0-381.5 69.5T232.5 385T128 640q0 112 71.5 213.5T401 1029l87 50l-27 96q-24 91-70 172q152-63 275-171l43-38l57 6q69 8 130 8q204 0 381.5-69.5t282-187.5T1664 640t-104.5-255t-282-187.5T896 128m896 512q0 174-120 321.5t-326 233t-450 85.5q-70 0-145-8q-198 175-460 242q-49 14-114 22h-5q-15 0-27-10.5t-16-27.5v-1q-3-4-.5-12t2-10t4.5-9.5l6-9l7-8.5l8-9q7-8 31-34.5t34.5-38t31-39.5t32.5-51t27-59t26-76q-157-89-247.5-220T0 640q0-130 71-248.5T262 187T548 50.5T896 0t348 50.5T1530 187t191 204.5t71 248.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comments(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1408 512q0 139-94 257t-256.5 186.5T704 1024q-86 0-176-16q-124 88-278 128q-36 9-86 16h-3q-11 0-20.5-8t-11.5-21q-1-3-1-6.5t.5-6.5t2-6l2.5-5l3.5-5.5l4-5l4.5-5l4-4.5q5-6 23-25t26-29.5t22.5-29t25-38.5t20.5-44Q142 841 71 736T0 512q0-139 94-257T350.5 68.5T704 0t353.5 68.5T1314 255t94 257m384 256q0 120-71 224.5T1526 1169q10 24 20.5 44t25 38.5t22.5 29t26 29.5t23 25q1 1 4 4.5t4.5 5t4 5t3.5 5.5l2.5 5l2 6l.5 6.5l-1 6.5q-3 14-13 22t-22 7q-50-7-86-16q-154-40-278-128q-90 16-176 16q-271 0-472-132q58 4 88 4q161 0 309-45t264-129q125-92 192-212t67-254q0-77-23-152q129 71 204 178t75 230"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentsO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M704 128q-153 0-286 52T206.5 321T128 512q0 82 53 158t149 132l97 56l-35 84q34-20 62-39l44-31l53 10q78 14 153 14q153 0 286-52t211.5-141t78.5-191t-78.5-191T990 180t-286-52m0-128q191 0 353.5 68.5T1314 255t94 257t-94 257t-256.5 186.5T704 1024q-86 0-176-16q-124 88-278 128q-36 9-86 16h-3q-11 0-20.5-8t-11.5-21q-1-3-1-6.5t.5-6.5t2-6l2.5-5l3.5-5.5l4-5l4.5-5l4-4.5q5-6 23-25t26-29.5t22.5-29t25-38.5t20.5-44Q142 841 71 736T0 512q0-139 94-257T350.5 68.5T704 0m822 1169q10 24 20.5 44t25 38.5t22.5 29t26 29.5t23 25q1 1 4 4.5t4.5 5t4 5t3.5 5.5l2.5 5l2 6l.5 6.5l-1 6.5q-3 14-13 22t-22 7q-50-7-86-16q-154-40-278-128q-90 16-176 16q-271 0-472-132q58 4 88 4q161 0 309-45t264-129q125-92 192-212t67-254q0-77-23-152q129 71 204 178t75 230q0 120-71 224.5T1526 1169"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compass(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m640 960l256-128l-256-128zm384-591v542l-512 256V625zm288 399q0-148-73-273t-198-198t-273-73t-273 73t-198 198t-73 273t73 273t198 198t273 73t273-73t198-198t73-273m224 0q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compress(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 832v448q0 26-19 45t-45 19t-45-19l-144-144l-332 332q-10 10-23 10t-23-10L23 1399q-10-10-10-23t10-23l332-332l-144-144q-19-19-19-45t19-45t45-19h448q26 0 45 19t19 45m755-672q0 13-10 23l-332 332l144 144q19 19 19 45t-19 45t-45 19H832q-26 0-45-19t-19-45V256q0-26 19-45t45-19t45 19l144 144l332-332q10-10 23-10t23 10l114 114q10 10 10 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Connectdevelop(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2048 895q0 21-13 36.5t-33 19.5l-205 356q3 9 3 18q0 20-12.5 35.5T1755 1380l-193 337q3 8 3 16q0 23-16.5 40t-40.5 17q-25 0-41-18h-400q-17 20-43 20t-43-20H582q-17 20-43 20q-23 0-40-16.5t-17-40.5q0-8 4-20l-193-335q-20-4-32.5-19.5T248 1325q0-9 3-18L45 951q-20-5-32.5-20.5T0 895q0-21 13.5-36.5T47 839l199-344q0-1-.5-3t-.5-3q0-36 34-51L488 75q-4-10-4-18q0-24 17-40.5T541 0q26 0 44 21h396q16-21 43-21t43 21h398q18-21 44-21q23 0 40 16.5t17 40.5q0 6-4 18l207 358q23 1 39 17.5t16 38.5q0 13-7 27l187 324q19 4 31.5 19.5T2048 895m-985 799h389l-342-354H967l-342 354h360q18-16 39-16t39 16M112 882q1 4 1 13q0 10-2 15l208 360l15 6l188-199V730L335 536q-13 8-29 10zM986 98H598l190 200l554-200h-280q-16 16-38 16t-38-16m703 1212q1-6 5-11l-64-68l-17 79zm-106 0l22-105l-252-266l-296 307l63 64zm-88 368l16-28l65-310h-427l333 343q8-4 13-5m-917 16h5l342-354H552v335l4 6q14 5 22 13m-26-384h402l64-66l-309-321l-157 166zm-193 0h163v-189l-168 177q4 8 5 12m-1-825q0 1 .5 2t.5 2q0 16-8 29l171 177V426zm194-70v311l153 157l297-314l-223-236zm4-304l-4 8v264l205-74l-191-201q-6 2-10 3m891-13h-16L810 322l213 225zm-424 492L726 905l311 319l296-307zM688 902L552 761v284zm350 364l-42 44h85zm336-348l238 251l132-624l-3-5l-1-1zm344-400q-8-13-8-29v-2l-216-376q-5-1-13-5l-437 463l310 327zM522 394V171L359 453zm0 946H359l163 283zm1085 0l-48 227l130-227zm122-70l207-361q-2-10-2-14q0-1 3-16l-171-296l-129 612l77 82q5-3 15-7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Contao(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M138 0h197q-70 64-126 149q-36 56-59 115t-30 125.5t-8.5 120t10.5 132t21 126T171 904q4 19 6 28q51 238 81 329q57 171 152 275H138q-48 0-82-34t-34-82V116q0-48 34-82t82-34m1208 0h308q48 0 82 34t34 82v1304q0 48-34 82t-82 34h-178q212-210 196-565l-469 101q-2 45-12 82t-31 72t-59.5 59.5t-93.5 36.5q-123 26-199-40q-32-27-53-61t-51.5-129T639 834q-35-163-45.5-263T588 432t23-77q20-41 62.5-73T776 237q45-12 83.5-6.5t67 17t54 35t43 48T1058 387l468-100q-68-175-180-287"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copy(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1696 384q40 0 68 28t28 68v1216q0 40-28 68t-68 28H736q-40 0-68-28t-28-68v-288H96q-40 0-68-28t-28-68V640q0-40 20-88t48-76L476 68q28-28 76-48t88-20h416q40 0 68 28t28 68v328q68-40 128-40zm-544 213L853 896h299zM512 213L213 512h299zm196 647l316-316V128H640v416q0 40-28 68t-68 28H128v640h512v-256q0-40 20-88t48-76m956 804V512h-384v416q0 40-28 68t-68 28H768v640z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copyright(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1150 946v109q0 50-36.5 89t-94 60.5t-118 32.5t-117.5 11q-205 0-342.5-139T304 763q0-203 136-339t339-136q34 0 75.5 4.5t93 18t92.5 34t69 56.5t28 81v109q0 16-16 16h-118q-16 0-16-16v-70q0-43-65.5-67.5T784 429q-140 0-228.5 91.5T467 758q0 151 91.5 249.5T792 1106q68 0 138-24t70-66v-70q0-7 4.5-11.5t10.5-4.5h119q6 0 11 4.5t5 11.5M768 128q-130 0-248.5 51t-204 136.5t-136.5 204T128 768t51 248.5t136.5 204t204 136.5t248.5 51t248.5-51t204-136.5t136.5-204t51-248.5t-51-248.5t-136.5-204t-204-136.5T768 128m768 640q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreativeCommons(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M605 1233q153 0 257-104q14-18 3-36l-45-82q-6-13-24-17q-16-2-27 11l-4 3q-4 4-11.5 10t-17.5 13.5t-23.5 14.5t-28.5 13t-33.5 9.5t-37.5 3.5q-76 0-125-50t-49-127q0-76 48-125.5T609 720q37 0 71.5 14t50.5 28l16 14q11 11 26 10q16-2 24-14l53-78q13-20-2-39q-3-4-11-12t-30-23.5t-48.5-28T691 569t-86-10q-148 0-246 96.5T261 896q0 146 97 241.5t247 95.5m630 0q153 0 257-104q14-18 4-36l-45-82q-8-14-25-17q-16-2-27 11l-4 3q-4 4-11.5 10t-17.5 13.5t-23.5 14.5t-28.5 13t-33.5 9.5t-37.5 3.5q-76 0-125-50t-49-127q0-76 48-125.5t122-49.5q37 0 71.5 14t50.5 28l16 14q11 11 26 10q16-2 24-14l53-78q13-20-2-39q-3-4-11-12t-30-23.5t-48.5-28T1321 569t-86-10q-147 0-245.5 96.5T891 896q0 146 97 241.5t247 95.5M896 160q-150 0-286 58.5t-234.5 157t-157 234.5T160 896t58.5 286t157 234.5t234.5 157t286 58.5t286-58.5t234.5-157t157-234.5t58.5-286t-58.5-286t-157-234.5t-234.5-157T896 160m0-160q182 0 348 71t286 191t191 286t71 348t-71 348t-191 286t-286 191t-348 71t-348-71t-286-191t-191-286T0 896t71-348t191-286T548 71T896 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCard(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1760 0q66 0 113 47t47 113v1216q0 66-47 113t-113 47H160q-66 0-113-47T0 1376V160Q0 94 47 47T160 0zM160 128q-13 0-22.5 9.5T128 160v224h1664V160q0-13-9.5-22.5T1760 128zm1600 1280q13 0 22.5-9.5t9.5-22.5V768H128v608q0 13 9.5 22.5t22.5 9.5zM256 1280v-128h256v128zm384 0v-128h384v128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardAlt(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1376V768h2304v608q0 66-47 113t-113 47H160q-66 0-113-47T0 1376m640-224v128h384v-128zm-384 0v128h256v-128zM2144 0q66 0 113 47t47 113v224H0V160Q0 94 47 47T160 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crop(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M557 1152h595V557zm-45-45l595-595H512zm1152 77v192q0 14-9 23t-23 9h-224v224q0 14-9 23t-23 9h-192q-14 0-23-9t-9-23v-224H288q-14 0-23-9t-9-23V512H32q-14 0-23-9t-9-23V288q0-14 9-23t23-9h224V32q0-14 9-23t23-9h192q14 0 23 9t9 23v224h851L1609 9q10-9 23-9t23 9q9 10 9 23t-9 23l-247 246v851h224q14 0 23 9t9 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crosshairs(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1197 896h-109q-26 0-45-19t-19-45V704q0-26 19-45t45-19h109q-32-108-112.5-188.5T896 339v109q0 26-19 45t-45 19H704q-26 0-45-19t-19-45V339q-108 32-188.5 112.5T339 640h109q26 0 45 19t19 45v128q0 26-19 45t-45 19H339q32 108 112.5 188.5T640 1197v-109q0-26 19-45t45-19h128q26 0 45 19t19 45v109q108-32 188.5-112.5T1197 896m339-192v128q0 26-19 45t-45 19h-143q-37 161-154.5 278.5T896 1329v143q0 26-19 45t-45 19H704q-26 0-45-19t-19-45v-143q-161-37-278.5-154.5T207 896H64q-26 0-45-19T0 832V704q0-26 19-45t45-19h143q37-161 154.5-278.5T640 207V64q0-26 19-45t45-19h128q26 0 45 19t19 45v143q161 37 278.5 154.5T1329 640h143q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CssThree(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M275 0h1505l-266 1333l-804 267l-698-267l71-356h297l-29 147l422 161l486-161l68-339H119l58-297h1209l38-191H216z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cube(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m896 1501l640-349V516L896 749zm-64-865l698-254l-698-254l-698 254zm832-252v768q0 35-18 65t-49 47l-704 384q-28 16-61 16t-61-16L67 1264q-31-17-49-47t-18-65V384q0-40 23-73t61-47L788 8q22-8 44-8t44 8l704 256q38 14 61 47t23 73"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cubes(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m640 1632l384-192v-314l-384 164zm-64-454l404-173l-404-173l-404 173zm1088 454l384-192v-314l-384 164zm-64-454l404-173l-404-173l-404 173zm-448-293l384-165V454l-384 164zm-64-379l441-189l-441-189l-441 189zm1088 518v416q0 36-19 67t-52 47l-448 224q-25 14-57 14t-57-14l-448-224q-4-2-7-4q-2 2-7 4l-448 224q-25 14-57 14t-57-14L71 1554q-33-16-52-47t-19-67v-416q0-38 21.5-70T78 906l434-186V320q0-38 21.5-70t56.5-48l448-192q23-10 50-10t50 10l448 192q35 16 56.5 48t21.5 70v400l434 186q36 16 57 48t21 70"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cut(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M960 640q26 0 45 19t19 45t-19 45t-45 19t-45-19t-19-45t19-45t45-19m300 64l507 398q28 20 25 56q-5 35-35 51l-128 64q-13 7-29 7q-17 0-31-8L879 885l-110 66q-8 4-12 5q14 49 10 97q-7 77-56 147.5T579 1324q-132 84-277 84q-136 0-222-78q-90-84-79-207q7-76 56-147t131-124q132-84 278-84q83 0 151 31q9-13 22-22l122-73l-122-73q-13-9-22-22q-68 31-151 31q-146 0-278-84q-82-53-131-124T1 285q-5-59 15.5-113T80 79Q165 0 302 0q145 0 277 84q83 52 132 123t56 148q4 48-10 97q4 1 12 5l110 66l690-387q14-8 31-8q16 0 29 7l128 64q30 16 35 51q3 36-25 56zM579 444q46-42 21-108T494 219q-92-59-192-59q-74 0-113 36q-46 42-21 108t106 117q92 59 192 59q74 0 113-36m-85 745q81-51 106-117t-21-108q-39-36-113-36q-100 0-192 59q-81 51-106 117t21 108q39 36 113 36q100 0 192-59m178-613l96 58v-11q0-36 33-56l14-8l-79-47l-26 26q-3 3-10 11t-12 12q-2 2-4 3.5t-3 2.5zm224 224l96 32l736-576l-128-64l-768 431v113l-160 96l9 8q2 2 7 6q4 4 11 12t11 12l26 26zm704 416l128-64l-520-408l-177 138q-2 3-13 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cutlery(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 64v640q0 61-35.5 111T512 885v779q0 52-38 90t-90 38H256q-52 0-90-38t-38-90V885q-57-20-92.5-70T0 704V64q0-26 19-45T64 0t45 19t19 45v416q0 26 19 45t45 19t45-19t19-45V64q0-26 19-45t45-19t45 19t19 45v416q0 26 19 45t45 19t45-19t19-45V64q0-26 19-45t45-19t45 19t19 45m768 0v1600q0 52-38 90t-90 38h-128q-52 0-90-38t-38-90v-512H800q-13 0-22.5-9.5T768 1120V320q0-132 94-226t226-94h256q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dashboard(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 896q0-53-37.5-90.5T256 768t-90.5 37.5T128 896t37.5 90.5T256 1024t90.5-37.5T384 896m192-448q0-53-37.5-90.5T448 320t-90.5 37.5T320 448t37.5 90.5T448 576t90.5-37.5T576 448m428 481l101-382q6-26-7.5-48.5T1059 469t-48 6.5t-30 39.5L880 897q-60 5-107 43.5t-63 98.5q-20 77 20 146t117 89t146-20t89-117q16-60-6-117t-72-91m660-33q0-53-37.5-90.5T1536 768t-90.5 37.5T1408 896t37.5 90.5t90.5 37.5t90.5-37.5T1664 896m-640-640q0-53-37.5-90.5T896 128t-90.5 37.5T768 256t37.5 90.5T896 384t90.5-37.5T1024 256m448 192q0-53-37.5-90.5T1344 320t-90.5 37.5T1216 448t37.5 90.5T1344 576t90.5-37.5T1472 448m320 448q0 261-141 483q-19 29-54 29H195q-35 0-54-29Q0 1158 0 896q0-182 71-348t191-286T548 71T896 0t348 71t286 191t191 286t71 348"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dashcube(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 680q0-131 91.5-226.5T314 358h742L1408 0v1470q0 132-91.5 227t-222.5 95H314q-131 0-222.5-95T0 1470zm1232 754l-176-180V829q0-46-32-79t-78-33H462q-46 0-78 33t-32 79v492q0 46 32.5 79.5T462 1434z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Database(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 768q237 0 443-43t325-127v170q0 69-103 128t-280 93.5t-385 34.5t-385-34.5T103 896T0 768V598q119 84 325 127t443 43m0 768q237 0 443-43t325-127v170q0 69-103 128t-280 93.5t-385 34.5t-385-34.5t-280-93.5T0 1536v-170q119 84 325 127t443 43m0-384q237 0 443-43t325-127v170q0 69-103 128t-280 93.5t-385 34.5t-385-34.5t-280-93.5T0 1152V982q119 84 325 127t443 43M768 0q208 0 385 34.5t280 93.5t103 128v128q0 69-103 128t-280 93.5T768 640t-385-34.5T103 512T0 384V256q0-69 103-128t280-93.5T768 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Deaf(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1056 832q0 26 19 45t45 19t45-19t19-45q0-146-103-249T832 480T583 583T480 832q0 26 19 45t45 19t45-19t19-45q0-93 66-158.5T832 608t158 65.5t66 158.5M835 256q-117 0-223.5 45.5t-184 123t-123 184T259 832q0 26 19 45t45 19t45-19t19-45q0-185 131.5-316.5T835 384t316.5 131.5T1283 832q0 55-18 103.5t-37.5 74.5t-59.5 72q-34 39-52 63t-43.5 66.5t-37 91T1024 1408q0 106-75 181t-181 75q-26 0-45 19t-19 45t19 45t45 19q159 0 271.5-112.5T1152 1408q0-41 7.5-74t26.5-64t33.5-50t45.5-54q35-41 53-64.5t44-67.5t37.5-93.5T1411 832q0-117-45.5-223.5t-123-184t-184-123T835 256M591 975l226 226l-579 579q-12 12-29 12t-29-12L12 1612q-12-12-12-29t12-29zM1612 12l168 168q12 12 12 29t-12 30l-233 233l-26 25l-71 71q-66-153-195-258l91-91l207-207q13-12 30-12t29 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dedent(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 416v576q0 13-9.5 22.5T352 1024q-14 0-23-9L41 727q-9-9-9-23t9-23l288-288q9-9 23-9q13 0 22.5 9.5T384 416m1408 768v192q0 13-9.5 22.5t-22.5 9.5H32q-13 0-22.5-9.5T0 1376v-192q0-13 9.5-22.5T32 1152h1728q13 0 22.5 9.5t9.5 22.5m0-384v192q0 13-9.5 22.5t-22.5 9.5H672q-13 0-22.5-9.5T640 992V800q0-13 9.5-22.5T672 768h1088q13 0 22.5 9.5t9.5 22.5m0-384v192q0 13-9.5 22.5T1760 640H672q-13 0-22.5-9.5T640 608V416q0-13 9.5-22.5T672 384h1088q13 0 22.5 9.5t9.5 22.5m0-384v192q0 13-9.5 22.5T1760 256H32q-13 0-22.5-9.5T0 224V32Q0 19 9.5 9.5T32 0h1728q13 0 22.5 9.5T1792 32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Delicious(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1472 1248V768H768V64H288q-93 0-158.5 65.5T64 288v480h704v704h480q93 0 158.5-65.5T1472 1248m64-960v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Desktop(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 992V160q0-13-9.5-22.5T1760 128H160q-13 0-22.5 9.5T128 160v832q0 13 9.5 22.5t22.5 9.5h1600q13 0 22.5-9.5t9.5-22.5m128-832v1088q0 66-47 113t-113 47h-544q0 37 16 77.5t32 71t16 43.5q0 26-19 45t-45 19H704q-26 0-45-19t-19-45q0-14 16-44t32-70t16-78H160q-66 0-113-47T0 1248V160Q0 94 47 47T160 0h1600q66 0 113 47t47 113"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Deviantart(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1024 303L721 885l24 31h279v415H517l-44 30l-142 273l-30 30H0v-303l303-583l-24-30H0V333h507l44-30L693 30l30-30h301z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Diamond(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m212 640l623 665l-300-665zm812 772l349-772H675zM538 512l204-384H480L192 512zm675 793l623-665h-323zM683 512h682l-204-384H887zm827 0h346l-288-384h-262zm141-486l384 512q14 18 13 41.5t-17 40.5l-960 1024q-18 20-47 20t-47-20L17 620Q1 603 0 579.5T13 538L397 26q18-26 51-26h1152q33 0 51 26"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Digg(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M328 26h204v983H0V312h328zm0 819V476H205v369zm286-533v697h205V312zm0-286v204h205V26zm287 286h533v942H901v-163h328v-82H901zm328 533V476h-123v369zm287-533h532v942h-532v-163h327v-82h-327zm327 533V476h-123v369z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dollar(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M946 1185q0 153-99.5 263.5T588 1585v175q0 14-9 23t-23 9H421q-13 0-22.5-9.5T389 1760v-175q-66-9-127.5-31T160 1509.5t-74-48t-46.5-37.5t-17.5-18q-17-21-2-41l103-135q7-10 23-12q15-2 24 9l2 2q113 99 243 125q37 8 74 8q81 0 142.5-43t61.5-122q0-28-15-53t-33.5-42t-58.5-37.5t-66-32t-80-32.5q-39-16-61.5-25T317 948.5t-62.5-31T198 882t-53.5-42.5t-43.5-49t-35.5-58t-21-66.5t-8.5-78q0-138 98-242t255-134V32q0-13 9.5-22.5T421 0h135q14 0 23 9t9 23v176q57 6 110.5 23t87 33.5T849 302t39 29t15 14q17 18 5 38l-81 146q-8 15-23 16q-14 3-27-7q-3-3-14.5-12t-39-26.5t-58.5-32t-74.5-26T505 430q-95 0-155 43t-60 111q0 26 8.5 48t29.5 41.5t39.5 33t56 31t60.5 27t70 27.5q53 20 81 31.5t76 35t75.5 42.5t62 50t53 63.5T933 1091t13 94"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotCircleO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1024 768q0 106-75 181t-181 75t-181-75t-75-181t75-181t181-75t181 75t75 181M768 224q-148 0-273 73T297 495t-73 273t73 273t198 198t273 73t273-73t198-198t73-273t-73-273t-198-198t-273-73m768 544q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1280 1344q0-26-19-45t-45-19t-45 19t-19 45t19 45t45 19t45-19t19-45m256 0q0-26-19-45t-45-19t-45 19t-19 45t19 45t45 19t45-19t19-45m128-224v320q0 40-28 68t-68 28H96q-40 0-68-28t-28-68v-320q0-40 28-68t68-28h465l135 136q58 56 136 56t136-56l136-136h464q40 0 68 28t28 68m-325-569q17 41-14 70l-448 448q-18 19-45 19t-45-19L339 621q-31-29-14-70q17-39 59-39h256V64q0-26 19-45t45-19h256q26 0 45 19t19 45v448h256q42 0 59 39"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dribbble(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1024 1372q-42-241-140-498h-2l-2 1q-16 6-43 16.5t-101 49t-137 82T468 1137t-103 148l-15-11q184 150 418 150q132 0 256-52M839 765q-21-49-53-111q-311 93-673 93q-1 7-1 21q0 124 44 236.5T280 1206q50-89 123.5-166.5T546 915t130.5-81t99.5-48l37-13q4-1 13-3.5t13-4.5M732 553Q612 340 488 175q-138 65-234 186T126 633q302 0 606-80m684 319q-210-60-409-29q87 239 128 469q111-75 185-189.5t96-250.5M611 131q-1 0-2 1q1-1 2-1m590 145q-185-164-433-164q-76 0-155 19q131 170 246 382q69-26 130-60.5t96.5-61.5t65.5-57t37.5-40.5zm223 485q-3-232-149-410l-1 1q-9 12-19 24.5t-43.5 44.5t-71 60.5t-100 65T909 611q25 53 44 95q2 5 6.5 17t7.5 17q36-5 74.5-7t73.5-2t69 1.5t64 4t56.5 5.5t48 6.5t36.5 6t25 4.5zm112 7q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DriversLicense(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M896 1084q0-54-7.5-100.5t-24.5-90t-51-68.5t-81-25q-64 64-156 64t-156-64q-47 0-81 25t-51 68.5t-24.5 90T256 1084q0 55 31.5 93.5T363 1216h426q44 0 75.5-38.5T896 1084M768 640q0-80-56-136t-136-56t-136 56t-56 136t56 136t136 56t136-56t56-136m1024 480v-64q0-14-9-23t-23-9h-704q-14 0-23 9t-9 23v64q0 14 9 23t23 9h704q14 0 23-9t9-23m-384-256v-64q0-14-9-23t-23-9h-320q-14 0-23 9t-9 23v64q0 14 9 23t23 9h320q14 0 23-9t9-23m384 0v-64q0-14-9-23t-23-9h-192q-14 0-23 9t-9 23v64q0 14 9 23t23 9h192q14 0 23-9t9-23m0-256v-64q0-14-9-23t-23-9h-704q-14 0-23 9t-9 23v64q0 14 9 23t23 9h704q14 0 23-9t9-23M128 256h1792v-96q0-14-9-23t-23-9H160q-14 0-23 9t-9 23zm1920-96v1216q0 66-47 113t-113 47H160q-66 0-113-47T0 1376V160Q0 94 47 47T160 0h1728q66 0 113 47t47 113"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DriversLicenseO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M896 1084q0 55-31.5 93.5T789 1216H363q-44 0-75.5-38.5T256 1084q0-54 7.5-100.5t24.5-90t51-68.5t81-25q64 64 156 64t156-64q47 0 81 25t51 68.5t24.5 90T896 1084M768 640q0 80-56 136t-136 56t-136-56t-56-136t56-136t136-56t136 56t56 136m1024 416v64q0 14-9 23t-23 9h-704q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h704q14 0 23 9t9 23m-384-256v64q0 14-9 23t-23 9h-320q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h320q14 0 23 9t9 23m384 0v64q0 14-9 23t-23 9h-192q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h192q14 0 23 9t9 23m0-256v64q0 14-9 23t-23 9h-704q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h704q14 0 23 9t9 23m128 832V256H128v1120q0 13 9.5 22.5t22.5 9.5h1728q13 0 22.5-9.5t9.5-22.5m128-1216v1216q0 66-47 113t-113 47H160q-66 0-113-47T0 1376V160Q0 94 47 47T160 0h1728q66 0 113 47t47 113"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dropbox(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m338 611l494 305l-342 285L0 882zm986 555v108l-490 293v1l-1-1l-1 1v-1l-489-293v-108l147 96l342-284v-2l1 1l1-1v2l343 284zM490 22l342 285l-494 304L0 341zm836 589l338 271l-489 319l-343-285zM1175 22l489 319l-338 270l-494-304z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drupal(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1167 1586q-5-19-24-5q-30 22-87 39t-131 17q-129 0-193-49q-5-4-13-4q-11 0-26 12q-7 6-7.5 16t7.5 20q34 32 87.5 46t102.5 12.5t99-4.5q41-4 84.5-20.5t65-30t28.5-20.5q12-12 7-29m-39-115q-19-47-39-61q-23-15-76-15q-47 0-71 10q-29 12-78 56q-26 24-12 44q9 8 17.5 4.5T901 1486q3-2 10.5-8.5t10.5-8.5t10-7t11.5-7t12.5-5t15-4.5t16.5-2.5t20.5-1q27 0 44.5 7.5t23 14.5t13.5 22q10 17 12.5 20t12.5-1q23-12 14-34m355-281q0-22-5-44.5t-16.5-45t-34-36.5t-52.5-14q-33 0-97 41.5t-129 83.5t-101 42q-27 1-63.5-19t-76-49t-83.5-58t-100-49t-111-19q-115 1-197 78.5T333 1280q-2 112 74 164q29 20 62.5 28.5T573 1481q57 0 132-32.5t134-71t120-70.5t93-31q26 1 65 31.5t71.5 67t68 67.5t55.5 32q35 3 58.5-14t55.5-63q28-41 42.5-101t14.5-106m53-160q0 164-62 304.5t-166 236t-242.5 149.5t-290.5 54t-293-57.5t-247.5-157T64 1318T0 1016q0-89 19.5-172.5t49-145.5T139 579.5t78.5-94T296 416t64.5-46.5T403 345q14-8 51-26.5t54.5-28.5t48-30t60.5-44q36-28 58-72.5T705 18q129 155 186 193q44 29 130 68t129 66q21 13 39 25t60.5 46.5t76 70.5t75 95t69 122t47 148.5T1536 1030"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Edge(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 795h1q16-126 58.5-241.5t115-217t167.5-176T570.5 43T847 0q231 0 414 105.5T1555 409q104 187 104 442v188H534q1 111 53.5 192.5T724 1354t189.5 57t213 3t208-46.5T1508 1283v377q-92 55-229.5 92T966 1790t-316-53q-189-73-311.5-249T214 1116q-3-242 111-412t325-268q-48 60-78 125.5T526 721h635q8-77-8-140t-47-101.5t-70.5-66.5t-80.5-41t-75-20.5t-56-8.5l-22-1q-135 5-259.5 44.5T319 491T143 631.5T5 795"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Edit(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m888 1056l116-116l-152-152l-116 116v56h96v96zm440-720q-16-16-33 1L945 687q-17 17-1 33t33-1l350-350q17-17 1-33m80 594v190q0 119-84.5 203.5T1120 1408H288q-119 0-203.5-84.5T0 1120V288Q0 169 84.5 84.5T288 0h832q63 0 117 25q15 7 18 23q3 17-9 29l-49 49q-14 14-32 8q-23-6-45-6H288q-66 0-113 47t-47 113v832q0 66 47 113t113 47h832q66 0 113-47t47-113V994q0-13 9-22l64-64q15-15 35-7t20 29m-96-738l288 288l-672 672H640V864zm444 132l-92 92l-288-288l92-92q28-28 68-28t68 28l152 152q28 28 28 68t-28 68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eercast(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1309 784q29-44-6.5-129.5T1181 512q-58-39-125.5-53.5t-118-4.5t-68.5 37q-12 23-4.5 28t42.5 10q23 3 38.5 5t44.5 9.5t56 17.5q36 13 67.5 31.5t53 37t40 38.5t30.5 38t22 34.5t16.5 28.5t12 18.5t10.5 6t11-9.5m363 574q-52 127-148.5 220T1309 1719.5t-253 60.5t-266-13.5t-251-91T329 1514t-141.5-235.5T141 975q1-41 8.5-84.5t12.5-64t24-80.5t23-73q-51 208 1 397t173 318t291 206t346 83t349-74.5t289-244.5q20-27 18-14q0 4-4 14m-239-449q0 104-40.5 199T1284 1272t-162 109.5t-198 40.5t-198-40.5T564 1272t-108.5-164T415 909t40.5-199T564 546t162-109.5T924 396t198 40.5T1284 546t108.5 164t40.5 199m287-288q-65-147-180.5-251t-253-153.5t-292-53.5t-301 36.5t-275.5 129T198 540T67 837t-10 373Q8 1049 5.5 898.5T41 626t109-227t165.5-180.5t207-126t232-71t242.5-9t236 54T1449 191t178 197q33 50 62 121t31 112m-62 342q12-244-136.5-416T1125 307q-8 0-10-5t24-8q125 4 230 50t173 120t116 168.5t58.5 199t-1 208T1654 1237t-122.5 167t-185 117.5T1098 1568q108-30 201.5-80t174-123t129.5-176.5t55-225.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eject(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 755L724 45q19-19 45-19t45 19l710 710q19 19 13 32t-32 13H33q-26 0-32-13t13-32m1459 557H65q-26 0-45-19t-19-45V992q0-26 19-45t45-19h1408q26 0 45 19t19 45v256q0 26-19 45t-45 19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisH(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 480v192q0 40-28 68t-68 28H96q-40 0-68-28T0 672V480q0-40 28-68t68-28h192q40 0 68 28t28 68m512 0v192q0 40-28 68t-68 28H608q-40 0-68-28t-28-68V480q0-40 28-68t68-28h192q40 0 68 28t28 68m512 0v192q0 40-28 68t-68 28h-192q-40 0-68-28t-28-68V480q0-40 28-68t68-28h192q40 0 68 28t28 68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisV(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 1120v192q0 40-28 68t-68 28H96q-40 0-68-28t-28-68v-192q0-40 28-68t68-28h192q40 0 68 28t28 68m0-512v192q0 40-28 68t-68 28H96q-40 0-68-28T0 800V608q0-40 28-68t68-28h192q40 0 68 28t28 68m0-512v192q0 40-28 68t-68 28H96q-40 0-68-28T0 288V96q0-40 28-68T96 0h192q40 0 68 28t28 68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Empire(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M874 1638v66q-208-6-385-109.5T206 1319l58-34q29 49 73 99l65-57q148 168 368 212l-17 86q65 12 121 13m-598-530l-83 28q22 60 49 112l-57 33q-98-180-98-385t98-385l57 33q-30 56-49 112l82 28q-35 100-35 212q0 109 36 212m1252 177l58 34q-106 172-283 275.5T918 1704v-66q56-1 121-13l-17-86q220-44 368-212l65 57q44-50 73-99m-151-554l-233 80q14 42 14 85t-14 85l232 80q-31 92-98 169l-185-162q-57 67-147 85l48 241q-52 10-98 10t-98-10l48-241q-90-18-147-85l-185 162q-67-77-98-169l232-80q-14-42-14-85t14-85l-233-80q33-93 99-169l185 162q59-68 147-86l-48-240q44-10 98-10t98 10l-48 240q88 18 147 86l185-162q66 76 99 169M874 88v66q-65 2-121 13l17 86q-220 42-368 211l-65-56q-38 42-73 98l-57-33q106-172 282-275.5T874 88m831 808q0 205-98 385l-57-33q27-52 49-112l-83-28q36-103 36-212q0-112-35-212l82-28q-19-56-49-112l57-33q98 180 98 385m-120-423l-57 33q-35-56-73-98l-65 56q-148-169-368-211l17-86q-56-11-121-13V88q209 6 385 109.5T1585 473m163 423q0-173-67.5-331T1499 293t-272-181.5T896 44t-331 67.5T293 293T111.5 565T44 896t67.5 331T293 1499t272 181.5t331 67.5t331-67.5t272-181.5t181.5-272t67.5-331m44 0q0 182-71 348t-191 286t-286 191t-348 71t-348-71t-286-191t-191-286T0 896t71-348t191-286T548 71T896 0t348 71t286 191t191 286t71 348"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Envelope(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 454v794q0 66-47 113t-113 47H160q-66 0-113-47T0 1248V454q44 49 101 87q362 246 497 345q57 42 92.5 65.5t94.5 48t110 24.5h2q51 0 110-24.5t94.5-48T1194 886q170-123 498-345q57-39 100-87m0-294q0 79-49 151t-122 123q-376 261-468 325q-10 7-42.5 30.5t-54 38t-52 32.5t-57.5 27t-50 9h-2q-23 0-50-9t-57.5-27t-52-32.5t-54-38T639 759q-91-64-262-182.5T172 434q-62-42-117-115.5T0 182q0-78 41.5-130T160 0h1472q65 0 112.5 47t47.5 113"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1664 1248V480q-32 36-69 66q-268 206-426 338q-51 43-83 67t-86.5 48.5T897 1024h-2q-48 0-102.5-24.5T706 951t-83-67Q465 752 197 546q-37-30-69-66v768q0 13 9.5 22.5t22.5 9.5h1472q13 0 22.5-9.5t9.5-22.5m0-1051v-24.5l-.5-13l-3-12.5l-5.5-9l-9-7.5l-14-2.5H160q-13 0-22.5 9.5T128 160q0 168 147 284q193 152 401 317q6 5 35 29.5t46 37.5t44.5 31.5T852 887t43 9h2q20 0 43-9t50.5-27.5T1035 828t46-37.5t35-29.5q208-165 401-317q54-43 100.5-115.5T1664 197m128-37v1088q0 66-47 113t-113 47H160q-66 0-113-47T0 1248V160Q0 94 47 47T160 0h1472q66 0 113 47t47 113"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeOpen(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 654v978q0 66-47 113t-113 47H160q-66 0-113-47T0 1632V654q0-15 11-24q8-7 39-34.5t41.5-36T137 522t70-55.5t96-73t143.5-107T639 146q5-4 52.5-40T763 53.5t64-35T896 0t69 18.5t65 35.5t71 52t52 40q110 80 192.5 140.5t143.5 107t96 73t70 55.5t45.5 37.5t41.5 36t39 34.5q11 9 11 24m-564 585q263-191 345-252q11-8 12.5-20.5T1579 943l-38-52q-8-11-21-12.5t-24 6.5q-231 169-343 250q-5 3-52 39t-71.5 52.5t-64.5 35t-69 18.5t-69-18.5t-64.5-35T691 1174t-52-39q-186-134-343-250q-11-8-24-6.5T251 891l-38 52q-8 11-6.5 23.5T219 987q82 61 345 252q10 8 50 38t65 47t64 39.5t77.5 33.5t75.5 11t75.5-11t79-34.5t64.5-39.5t65-47.5t48-36.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeOpenO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1474 913l39 51q8 11 6.5 23.5T1508 1008q-43 34-126.5 98.5t-146.5 113t-67 51.5q-39 32-60 48t-60.5 41t-76.5 36.5t-74 11.5h-2q-37 0-74-11.5t-76-36.5t-61-41.5t-60-47.5q-5-4-65-50.5t-143.5-111T293 1015q-11-8-12.5-20.5T287 971l37-52q8-11 21.5-13t24.5 7q94 73 306 236q5 4 43.5 35t60.5 46.5t56.5 32.5t58.5 17h2q24 0 58.5-17t56.5-32.5t60.5-46.5t43.5-35q258-198 313-242q11-8 24-6.5t21 12.5m190 719V704q-90-83-159-139q-91-74-389-304q-3-2-43-35t-61-48t-56-32.5t-59-17.5h-2q-24 0-59 17.5T780 178t-61 48t-43 35Q461 427 360.5 506.5T231 610.5T149 685q-14 12-21 19v928q0 13 9.5 22.5t22.5 9.5h1472q13 0 22.5-9.5t9.5-22.5m128-928v928q0 66-47 113t-113 47H160q-66 0-113-47T0 1632V704q0-56 41-94q123-114 350-290.5T624 138q36-30 59-47.5t61.5-42t76-36.5T895 0h2q37 0 74.5 12t76 36.5t61.5 42t59 47.5q43 36 156 122t226 177t201 173q41 38 41 94"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1248 0q119 0 203.5 84.5T1536 288v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0zm32 1056V620q-31 35-64 55q-34 22-132.5 85T932 859q-98 69-164 69t-164-69q-47-32-142-92.5T320 674q-12-8-33-27t-31-27v436q0 40 28 68t68 28h832q40 0 68-28t28-68m0-573q0-41-27.5-70t-68.5-29H352q-40 0-68 28t-28 68q0 37 30.5 76.5T354 621q47 32 137.5 89T621 793q3 2 17 11.5t21 14t21 13t23.5 13T725 854t22.5 7.5T768 864t20.5-2.5T811 854t21.5-9.5t23.5-13t21-13t21-14t17-11.5l267-174q35-23 66.5-62.5T1280 483"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Envira(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M896 816Q792 620 736 538Q597 336 389 220q-34-19-70-36q-89-40-94-32t34 38l39 31q62 43 112.5 93.5T505 431t70.5 113T646 675q9 17 13 25q44 84 84 153t98 154t115.5 150t131 123.5T1236 1371q153 66 154 60q1-3-49-37q-53-36-81-57q-77-58-179-211T896 816m-347 543q-76-60-132.5-125t-98-143.5t-71-154.5T189 750t-52-209t-60.5-252T0 0q273 0 497.5 36t379 92t271 144.5T1333 445t110 198.5t56 199.5t12.5 198.5t-9.5 173t-20 143.5t-13 107l323 327h-104l-281-285q-22 2-91.5 14t-121.5 19t-138 6t-160.5-17t-167.5-59t-179-111"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eraser(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m896 1152l336-384H464l-336 384zM1909 75q15 34 9.5 71.5T1888 212L992 1236q-38 44-96 44H128q-38 0-69.5-20.5T11 1205q-15-34-9.5-71.5T32 1068L928 44q38-44 96-44h768q38 0 69.5 20.5T1909 75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Etsy(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M518 183v655q103 1 191.5-1.5T835 831l37-3q68-2 90.5-24.5T1002 709l33-142h103l-14 322l7 319h-103l-29-127q-15-68-45-93t-84-26q-87-8-352-8v556q0 78 43.5 115.5T695 1663h357q35 0 59.5-2t55-7.5t54-18t48.5-32t46-50.5t39-73l93-216h89q-6 37-31.5 252t-30.5 276q-146-5-263.5-8t-162.5-4H376L0 1792v-102l127-25q67-13 91.5-37t25.5-79l8-643q3-402-8-645q-2-61-25.5-84T127 141L0 117V15l376 12h702q139 0 374-27q-6 68-14 194.5T1426 414l-5 92h-93l-32-124q-31-121-74-179.5T1109 144H561q-28 0-35.5 8.5T518 183"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eur(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m976 1179l35 159q3 12-3 22.5t-17 14.5l-5 1q-4 2-10.5 3.5t-16 4.5t-21.5 5.5t-25.5 5t-30 5t-33.5 4.5t-36.5 3t-38.5 1q-234 0-409-130.5T127 926H32q-13 0-22.5-9.5T0 894V781q0-13 9.5-22.5T32 749h66q-2-57 1-105H32q-14 0-23-9t-9-23V498q0-14 9-23t23-9h98q67-210 243.5-338T774 0q102 0 194 23q11 3 20 15q6 11 3 24l-43 159q-3 13-14 19.5t-24 2.5l-4-1q-4-1-11.5-2.5L877 236l-22.5-3.5l-26-3l-29-2.5l-29.5-1q-126 0-226 64T394 466h468q16 0 25 12q10 12 7 26l-24 114q-5 26-32 26H350q-3 37 0 105h459q15 0 25 12q9 12 6 27l-24 112q-2 11-11 18.5t-20 7.5H398q48 117 149.5 185.5T776 1180q18 0 36-1.5t33.5-3.5t29.5-4.5t24.5-5t18.5-4.5l12-3l5-2q13-5 26 2q12 7 15 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Exchange(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 928v192q0 13-9.5 22.5t-22.5 9.5H384v192q0 13-9.5 22.5T352 1376q-12 0-24-10L9 1046q-9-9-9-22q0-14 9-23l320-320q9-9 23-9q13 0 22.5 9.5T384 704v192h1376q13 0 22.5 9.5t9.5 22.5m0-544q0 14-9 23l-320 320q-9 9-23 9q-13 0-22.5-9.5T1408 704V512H32q-13 0-22.5-9.5T0 480V288q0-13 9.5-22.5T32 256h1376V64q0-14 9-23t23-9q12 0 24 10l319 319q9 9 9 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Exclamation(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M416 1120v224q0 26-19 45t-45 19H96q-26 0-45-19t-19-45v-224q0-26 19-45t45-19h256q26 0 45 19t19 45M446 64l-28 768q-1 26-20.5 45T352 896H96q-26 0-45.5-19T30 832L2 64Q1 38 19.5 19T64 0h320q26 0 44.5 19T446 64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationCircle(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 0q209 0 385.5 103T1433 382.5T1536 768t-103 385.5t-279.5 279.5T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0m128 1247v-190q0-14-9-23.5t-22-9.5H673q-13 0-23 10t-10 23v190q0 13 10 23t23 10h192q13 0 22-9.5t9-23.5m-2-344l18-621q0-12-10-18q-10-8-24-8H658q-14 0-24 8q-10 6-10 18l17 621q0 10 10 17.5t24 7.5h185q14 0 23.5-7.5T894 903"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationTriangle(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1056 1375v-190q0-14-9.5-23.5t-22.5-9.5H832q-13 0-22.5 9.5T800 1185v190q0 14 9.5 23.5t22.5 9.5h192q13 0 22.5-9.5t9.5-23.5m-2-374l18-459q0-12-10-19q-13-11-24-11H818q-11 0-24 11q-10 7-10 21l17 457q0 10 10 16.5t24 6.5h185q14 0 23.5-6.5t10.5-16.5m-14-934l768 1408q35 63-2 126q-17 29-46.5 46t-63.5 17H160q-34 0-63.5-17T50 1601q-37-63-2-126L816 67q17-31 47-49t65-18t65 18t47 49"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Expand(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M755 928q0 13-10 23l-332 332l144 144q19 19 19 45t-19 45t-45 19H64q-26 0-45-19t-19-45v-448q0-26 19-45t45-19t45 19l144 144l332-332q10-10 23-10t23 10l114 114q10 10 10 23m781-864v448q0 26-19 45t-45 19t-45-19l-144-144l-332 332q-10 10-23 10t-23-10L791 631q-10-10-10-23t10-23l332-332l-144-144q-19-19-19-45t19-45t45-19h448q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Expeditedssl(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M896 64q-169 0-323 66T307.5 307.5T130 573T64 896t66 323t177.5 265.5T573 1662t323 66t323-66t265.5-177.5T1662 1219t66-323t-66-323t-177.5-265.5T1219 130T896 64m0-64q182 0 348 71t286 191t191 286t71 348t-71 348t-191 286t-286 191t-348 71t-348-71t-286-191t-191-286T0 896t71-348t191-286T548 71T896 0M496 832q16 0 16 16v480q0 16-16 16h-32q-16 0-16-16V848q0-16 16-16zm400 64q53 0 90.5 37.5t37.5 90.5q0 35-17.5 64t-46.5 46v114q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23v-114q-29-17-46.5-46t-17.5-64q0-53 37.5-90.5T896 896m0-768q209 0 385.5 103T1561 510.5T1664 896t-103 385.5t-279.5 279.5T896 1664t-385.5-103T231 1281.5T128 896t103-385.5T510.5 231T896 128M544 608v96q0 14 9 23t23 9h64q14 0 23-9t9-23v-96q0-93 65.5-158.5T896 384t158.5 65.5T1120 608v96q0 14 9 23t23 9h64q14 0 23-9t9-23v-96q0-146-103-249T896 256T647 359T544 608m864 736V832q0-26-19-45t-45-19H448q-26 0-45 19t-19 45v512q0 26 19 45t45 19h896q26 0 45-19t19-45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExternalLink(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1408 928v320q0 119-84.5 203.5T1120 1536H288q-119 0-203.5-84.5T0 1248V416q0-119 84.5-203.5T288 128h704q14 0 23 9t9 23v64q0 14-9 23t-23 9H288q-66 0-113 47t-47 113v832q0 66 47 113t113 47h832q66 0 113-47t47-113V928q0-14 9-23t23-9h64q14 0 23 9t9 23m384-864v512q0 26-19 45t-45 19t-45-19l-176-176l-652 652q-10 10-23 10t-23-10L695 983q-10-10-10-23t10-23l652-652l-176-176q-19-19-19-45t19-45t45-19h512q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExternalLinkSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1280 800V320q0-26-19-45t-45-19H736q-42 0-59 39q-17 41 14 70l144 144l-534 534q-19 19-19 45t19 45l102 102q19 19 45 19t45-19l534-534l144 144q18 19 45 19q12 0 25-5q39-17 39-59m256-512v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1664 704q-152-236-381-353q61 104 61 225q0 185-131.5 316.5T896 1024T579.5 892.5T448 576q0-121 61-225q-229 117-381 353q133 205 333.5 326.5T896 1152t434.5-121.5T1664 704M944 320q0-20-14-34t-34-14q-125 0-214.5 89.5T592 576q0 20 14 34t34 14t34-14t14-34q0-86 61-147t147-61q20 0 34-14t14-34m848 384q0 34-20 69q-140 230-376.5 368.5T896 1280t-499.5-139T20 773Q0 738 0 704t20-69q140-229 376.5-368T896 128t499.5 139T1772 635q20 35 20 69"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeSlash(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m555 1079l78-141q-87-63-136-159t-49-203q0-121 61-225q-229 117-381 353q167 258 427 375m389-759q0-20-14-34t-34-14q-125 0-214.5 89.5T592 576q0 20 14 34t34 14t34-14t14-34q0-86 61-147t147-61q20 0 34-14t14-34m363-191q0 7-1 9q-106 189-316 567t-315 566l-49 89q-10 16-28 16q-12 0-134-70q-16-10-16-28q0-12 44-87q-143-65-263.5-173T20 773Q0 742 0 704t20-69q153-235 380-371t496-136q89 0 180 17l54-97q10-16 28-16q5 0 18 6t31 15.5t33 18.5t31.5 18.5T1291 102q16 10 16 27m37 447q0 139-79 253.5T1056 994l280-502q8 45 8 84m448 128q0 35-20 69q-39 64-109 145q-150 172-347.5 267T896 1280l74-132q212-18 392.5-137T1664 704q-115-179-282-294l63-112q95 64 182.5 153T1772 635q20 34 20 69"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eyedropper(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1698 94q94 94 94 226.5T1698 546l-225 223l104 104q10 10 10 23t-10 23l-210 210q-10 10-23 10t-23-10l-105-105l-603 603q-37 37-90 37H320L64 1792l-64-64l128-256v-203q0-53 37-90l603-603l-105-105q-10-10-10-23t10-23l210-210q10-10 23-10t23 10l104 104l223-225q93-94 225.5-94T1698 94M512 1472l576-576l-192-192l-576 576v192z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fa(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1504 512v839q0 48-49 62q-174 52-338 52q-73 0-215.5-29.5T674 1406q-164 0-370 48v338H144V424q-63-25-101-81T5 219q0-91 64-155T224 0t155 64t64 155q0 68-38 124t-101 81v68q190-44 343-44q99 0 198 15q14 2 111.5 22.5T1106 506q77 0 165-18q11-2 80-21t89-19q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Facebook(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M895 12v264H738q-86 0-116 36t-30 108v189h293l-39 296H592v759H286V905H31V609h255V391q0-186 104-288.5T667 0q147 0 228 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacebookOfficial(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1451 0q35 0 60 25t25 60v1366q0 35-25 60t-60 25h-391V941h199l30-232h-229V561q0-56 23.5-84t91.5-28l122-1V241q-63-9-178-9q-136 0-217.5 80T820 538v171H620v232h200v595H85q-35 0-60-25t-25-60V85q0-35 25-60T85 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacebookSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1248 0q119 0 203.5 84.5T1536 288v960q0 119-84.5 203.5T1248 1536h-188V941h199l30-232h-229V561q0-56 23.5-84t91.5-28l122-1V241q-63-9-178-9q-136 0-217.5 80T820 538v171H620v232h200v595H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastBackward(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1747 13q19-19 32-13t13 32v1472q0 26-13 32t-32-13l-710-710q-9-9-13-19v710q0 26-13 32t-32-13L269 813q-9-9-13-19v678q0 26-19 45t-45 19H64q-26 0-45-19t-19-45V64q0-26 19-45T64 0h128q26 0 45 19t19 45v678q4-10 13-19L979 13q19-19 32-13t13 32v710q4-10 13-19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fax(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M288 384q66 0 113 47t47 113v1088q0 66-47 113t-113 47H160q-66 0-113-47T0 1632V544q0-66 47-113t113-47zm1376 163q58 34 93 93t35 128v768q0 106-75 181t-181 75H672q-66 0-113-47t-47-113V96q0-40 28-68t68-28h672q40 0 88 20t76 48l152 152q28 28 48 76t20 88zm-736 989v-128q0-14-9-23t-23-9H768q-14 0-23 9t-9 23v128q0 14 9 23t23 9h128q14 0 23-9t9-23m0-256v-128q0-14-9-23t-23-9H768q-14 0-23 9t-9 23v128q0 14 9 23t23 9h128q14 0 23-9t9-23m0-256V896q0-14-9-23t-23-9H768q-14 0-23 9t-9 23v128q0 14 9 23t23 9h128q14 0 23-9t9-23m256 512v-128q0-14-9-23t-23-9h-128q-14 0-23 9t-9 23v128q0 14 9 23t23 9h128q14 0 23-9t9-23m0-256v-128q0-14-9-23t-23-9h-128q-14 0-23 9t-9 23v128q0 14 9 23t23 9h128q14 0 23-9t9-23m0-256V896q0-14-9-23t-23-9h-128q-14 0-23 9t-9 23v128q0 14 9 23t23 9h128q14 0 23-9t9-23m256 512v-128q0-14-9-23t-23-9h-128q-14 0-23 9t-9 23v128q0 14 9 23t23 9h128q14 0 23-9t9-23m0-256v-128q0-14-9-23t-23-9h-128q-14 0-23 9t-9 23v128q0 14 9 23t23 9h128q14 0 23-9t9-23m0-256V896q0-14-9-23t-23-9h-128q-14 0-23 9t-9 23v128q0 14 9 23t23 9h128q14 0 23-9t9-23m96-384V384h-160q-40 0-68-28t-28-68V128H640v512z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Feed(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 1216q0 80-56 136t-136 56t-136-56t-56-136t56-136t136-56t136 56t56 136m512 123q2 28-17 48q-18 21-47 21H697q-25 0-43-16.5t-20-41.5q-22-229-184.5-391.5T58 774q-25-2-41.5-20T0 711V576q0-29 21-47q17-17 43-17h5q160 13 306 80.5T634 774q114 113 181.5 259t80.5 306m512 2q2 27-18 47q-18 20-46 20h-143q-26 0-44.5-17.5T1137 1348q-12-215-101-408.5t-231.5-336t-336-231.5T60 270q-25-1-42.5-19.5T0 207V64q0-28 20-46Q38 0 64 0h3q262 13 501.5 120T994 414q187 186 294 425.5t120 501.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Female(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1280 1024q0 40-28 68t-68 28q-51 0-80-43L877 736h-45v132l247 411q9 15 9 33q0 26-19 45t-45 19H832v272q0 46-33 79t-79 33H560q-46 0-79-33t-33-79v-272H256q-26 0-45-19t-19-45q0-18 9-33l247-411V736h-45l-227 341q-29 43-80 43q-40 0-68-28t-28-68q0-29 16-53l256-384q73-107 176-107h384q103 0 176 107l256 384q16 24 16 53M864 224q0 93-65.5 158.5T640 448t-158.5-65.5T416 224t65.5-158.5T640 0t158.5 65.5T864 224"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FighterJet(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1920 704q-1 32-288 96l-352 32l-224 64h-64l-293 352h69q26 0 45 4.5t19 11.5t-19 11.5t-45 4.5H448v-32h64V832H352l-192 224H64l-32-32V832h32v-32h128v-8L0 768V640l192-24v-8H64v-32H32V384l32-32h96l192 224h160V160h-64v-32h320q26 0 45 4.5t19 11.5t-19 11.5t-45 4.5h-69l293 352h64l224 64l352 32q128 28 200 52t80 34z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func File(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1024 512V40q22 14 36 28l408 408q14 14 28 36zm-128 32q0 40 28 68t68 28h544v1056q0 40-28 68t-68 28H96q-40 0-68-28t-28-68V96q0-40 28-68T96 0h800z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileArchiveO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 384V256H512v128zm128 128V384H640v128zM640 640V512H512v128zm128 128V640H640v128zm700-388q28 28 48 76t20 88v1152q0 40-28 68t-68 28H96q-40 0-68-28t-28-68V96q0-40 28-68T96 0h896q40 0 88 20t76 48zm-444-244v376h376q-10-29-22-41l-313-313q-12-12-41-22m384 1528V640H992q-40 0-68-28t-28-68V128H768v128H640V128H128v1536zM781 943l107 349q8 27 8 52q0 83-72.5 137.5T640 1536t-183.5-54.5T384 1344q0-25 8-52q21-63 120-396V768h128v128h79q22 0 39 13t23 34m-141 465q53 0 90.5-19t37.5-45t-37.5-45t-90.5-19t-90.5 19t-37.5 45t37.5 45t90.5 19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileAudioO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1468 380q28 28 48 76t20 88v1152q0 40-28 68t-68 28H96q-40 0-68-28t-28-68V96q0-40 28-68T96 0h896q40 0 88 20t76 48zm-444-244v376h376q-10-29-22-41l-313-313q-12-12-41-22m384 1528V640H992q-40 0-68-28t-28-68V128H128v1536zM620 850q20 8 20 30v544q0 22-20 30q-8 2-12 2q-12 0-23-9l-166-167H288q-14 0-23-9t-9-23v-192q0-14 9-23t23-9h131l166-167q16-15 35-7m417 689q31 0 50-24q129-159 129-363t-129-363q-16-21-43-24t-47 14q-21 17-23.5 43.5T988 870q100 123 100 282t-100 282q-17 21-14.5 47.5T997 1524q18 15 40 15m-211-148q27 0 47-20q87-93 87-219t-87-219q-18-19-45-20t-46 17t-20 44.5t18 46.5q52 57 52 131t-52 131q-19 20-18 46.5t20 44.5q20 17 44 17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileCodeO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1468 380q28 28 48 76t20 88v1152q0 40-28 68t-68 28H96q-40 0-68-28t-28-68V96q0-40 28-68T96 0h896q40 0 88 20t76 48zm-444-244v376h376q-10-29-22-41l-313-313q-12-12-41-22m384 1528V640H992q-40 0-68-28t-28-68V128H128v1536zM480 768q8-11 21-12.5t24 6.5l51 38q11 8 12.5 21t-6.5 24l-182 243l182 243q8 11 6.5 24t-12.5 21l-51 38q-11 8-24 6.5t-21-12.5l-226-301q-14-19 0-38zm802 301q14 19 0 38l-226 301q-8 11-21 12.5t-24-6.5l-51-38q-11-8-12.5-21t6.5-24l182-243l-182-243q-8-11-6.5-24t12.5-21l51-38q11-8 24-6.5t21 12.5zm-620 461q-13-2-20.5-13t-5.5-24l138-831q2-13 13-20.5t24-5.5l63 10q13 2 20.5 13t5.5 24l-138 831q-2 13-13 20.5t-24 5.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileExcelO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1468 380q28 28 48 76t20 88v1152q0 40-28 68t-68 28H96q-40 0-68-28t-28-68V96q0-40 28-68T96 0h896q40 0 88 20t76 48zm-444-244v376h376q-10-29-22-41l-313-313q-12-12-41-22m384 1528V640H992q-40 0-68-28t-28-68V128H128v1536zm-979-234v106h281v-106h-75l103-161q5-7 10-16.5t7.5-13.5t3.5-4h2q1 4 5 10q2 4 4.5 7.5t6 8t6.5 8.5l107 161h-76v106h291v-106h-68l-192-273l195-282h67V768H828v107h74l-103 159q-4 7-10 16.5t-9 13.5l-2 3h-2q-1-4-5-10q-6-11-17-23L648 875h76V768H434v107h68l189 272l-194 283z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileImageO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1468 380q28 28 48 76t20 88v1152q0 40-28 68t-68 28H96q-40 0-68-28t-28-68V96q0-40 28-68T96 0h896q40 0 88 20t76 48zm-444-244v376h376q-10-29-22-41l-313-313q-12-12-41-22m384 1528V640H992q-40 0-68-28t-28-68V128H128v1536zm-128-448v320H256v-192l192-192l128 128l384-384zm-832-192q-80 0-136-56t-56-136t56-136t136-56t136 56t56 136t-56 136t-136 56"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileMovieO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1468 380q28 28 48 76t20 88v1152q0 40-28 68t-68 28H96q-40 0-68-28t-28-68V96q0-40 28-68T96 0h896q40 0 88 20t76 48zm-444-244v376h376q-10-29-22-41l-313-313q-12-12-41-22m384 1528V640H992q-40 0-68-28t-28-68V128H128v1536zM768 768q52 0 90 38t38 90v384q0 52-38 90t-90 38H384q-52 0-90-38t-38-90V896q0-52 38-90t90-38zm492 2q20 8 20 30v576q0 22-20 30q-8 2-12 2q-14 0-23-9l-265-266v-90l265-266q9-9 23-9q4 0 12 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1468 380q28 28 48 76t20 88v1152q0 40-28 68t-68 28H96q-40 0-68-28t-28-68V96q0-40 28-68T96 0h896q40 0 88 20t76 48zm-444-244v376h376q-10-29-22-41l-313-313q-12-12-41-22m384 1528V640H992q-40 0-68-28t-28-68V128H128v1536z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePdfO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1468 380q28 28 48 76t20 88v1152q0 40-28 68t-68 28H96q-40 0-68-28t-28-68V96q0-40 28-68T96 0h896q40 0 88 20t76 48zm-444-244v376h376q-10-29-22-41l-313-313q-12-12-41-22m384 1528V640H992q-40 0-68-28t-28-68V128H128v1536zm-514-593q33 26 84 56q59-7 117-7q147 0 177 49q16 22 2 52q0 1-1 2l-2 2v1q-6 38-71 38q-48 0-115-20t-130-53q-221 24-392 83q-153 262-242 262q-15 0-28-7l-24-12q-1-1-6-5q-10-10-6-36q9-40 56-91.5t132-96.5q14-9 23 6q2 2 2 4q52-85 107-197q68-136 104-262q-24-82-30.5-159.5T657 552q11-40 42-40h22q23 0 35 15q18 21 9 68q-2 6-4 8q1 3 1 8v30q-2 123-14 192q55 164 146 238m-576 411q52-24 137-158q-51 40-87.5 84t-49.5 74m398-920q-15 42-2 132q1-7 7-44q0-3 7-43q1-4 4-8q-1-1-1-2q-1-2-1-3q-1-22-13-36q0 1-1 2zm-124 661q135-54 284-81q-2-1-13-9.5t-16-13.5q-76-67-127-176q-27 86-83 197q-30 56-45 83m646-16q-24-24-140-24q76 28 124 28q14 0 18-1q0-1-2-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePowerpointO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1468 380q28 28 48 76t20 88v1152q0 40-28 68t-68 28H96q-40 0-68-28t-28-68V96q0-40 28-68T96 0h896q40 0 88 20t76 48zm-444-244v376h376q-10-29-22-41l-313-313q-12-12-41-22m384 1528V640H992q-40 0-68-28t-28-68V128H128v1536zm-992-234v106h327v-106h-93v-167h137q76 0 118-15q67-23 106.5-87t39.5-146q0-81-37-141t-100-87q-48-19-130-19H416v107h92v555zm353-280H650V882h120q52 0 83 18q56 33 56 115q0 89-62 120q-31 15-78 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileText(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1468 476q14 14 28 36h-472V40q22 14 36 28zM992 640h544v1056q0 40-28 68t-68 28H96q-40 0-68-28t-28-68V96q0-40 28-68T96 0h800v544q0 40 28 68t68 28m160 736v-64q0-14-9-23t-23-9H416q-14 0-23 9t-9 23v64q0 14 9 23t23 9h704q14 0 23-9t9-23m0-256v-64q0-14-9-23t-23-9H416q-14 0-23 9t-9 23v64q0 14 9 23t23 9h704q14 0 23-9t9-23m0-256v-64q0-14-9-23t-23-9H416q-14 0-23 9t-9 23v64q0 14 9 23t23 9h704q14 0 23-9t9-23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileTextO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1468 380q28 28 48 76t20 88v1152q0 40-28 68t-68 28H96q-40 0-68-28t-28-68V96q0-40 28-68T96 0h896q40 0 88 20t76 48zm-444-244v376h376q-10-29-22-41l-313-313q-12-12-41-22m384 1528V640H992q-40 0-68-28t-28-68V128H128v1536zM384 800q0-14 9-23t23-9h704q14 0 23 9t9 23v64q0 14-9 23t-23 9H416q-14 0-23-9t-9-23zm736 224q14 0 23 9t9 23v64q0 14-9 23t-23 9H416q-14 0-23-9t-9-23v-64q0-14 9-23t23-9zm0 256q14 0 23 9t9 23v64q0 14-9 23t-23 9H416q-14 0-23-9t-9-23v-64q0-14 9-23t23-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileWordO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1468 380q28 28 48 76t20 88v1152q0 40-28 68t-68 28H96q-40 0-68-28t-28-68V96q0-40 28-68T96 0h896q40 0 88 20t76 48zm-444-244v376h376q-10-29-22-41l-313-313q-12-12-41-22m384 1528V640H992q-40 0-68-28t-28-68V128H128v1536zM233 768v107h70l164 661h159l128-485q7-20 10-46q2-16 2-24h4l3 24q1 3 3.5 20t5.5 26l128 485h159l164-661h70V768h-300v107h90l-99 438q-5 20-7 46l-2 21h-4q0-3-.5-6.5t-1.5-8t-1-6.5q-1-5-4-21t-5-25L825 768H711l-144 545q-2 9-4.5 24.5T559 1359l-4 21h-4l-2-21q-2-26-7-46l-99-438h90V768z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Film(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 1472v-128q0-26-19-45t-45-19H192q-26 0-45 19t-19 45v128q0 26 19 45t45 19h128q26 0 45-19t19-45m0-384V960q0-26-19-45t-45-19H192q-26 0-45 19t-19 45v128q0 26 19 45t45 19h128q26 0 45-19t19-45m0-384V576q0-26-19-45t-45-19H192q-26 0-45 19t-19 45v128q0 26 19 45t45 19h128q26 0 45-19t19-45m1024 768V960q0-26-19-45t-45-19H576q-26 0-45 19t-19 45v512q0 26 19 45t45 19h768q26 0 45-19t19-45M384 320V192q0-26-19-45t-45-19H192q-26 0-45 19t-19 45v128q0 26 19 45t45 19h128q26 0 45-19t19-45m1408 1152v-128q0-26-19-45t-45-19h-128q-26 0-45 19t-19 45v128q0 26 19 45t45 19h128q26 0 45-19t19-45m-384-768V192q0-26-19-45t-45-19H576q-26 0-45 19t-19 45v512q0 26 19 45t45 19h768q26 0 45-19t19-45m384 384V960q0-26-19-45t-45-19h-128q-26 0-45 19t-19 45v128q0 26 19 45t45 19h128q26 0 45-19t19-45m0-384V576q0-26-19-45t-45-19h-128q-26 0-45 19t-19 45v128q0 26 19 45t45 19h128q26 0 45-19t19-45m0-384V192q0-26-19-45t-45-19h-128q-26 0-45 19t-19 45v128q0 26 19 45t45 19h128q26 0 45-19t19-45m128-160v1344q0 66-47 113t-113 47H160q-66 0-113-47T0 1504V160Q0 94 47 47T160 0h1600q66 0 113 47t47 113"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filter(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1403 39q17 41-14 70L896 602v742q0 42-39 59q-13 5-25 5q-27 0-45-19l-256-256q-19-19-19-45V602L19 109Q-12 80 5 39Q22 0 64 0h1280q42 0 59 39"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fire(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1408 1696v64q0 13-9.5 22.5t-22.5 9.5H32q-13 0-22.5-9.5T0 1760v-64q0-13 9.5-22.5T32 1664h1344q13 0 22.5 9.5t9.5 22.5M1152 640q0 78-24.5 144t-64 112.5t-87.5 88t-96 77.5t-87.5 72t-64 81.5T704 1312q0 96 67 224l-4-1l1 1q-90-41-160-83t-138.5-100T356 1230.5T283.5 1080T256 896q0-78 24.5-144t64-112.5t87.5-88t96-77.5t87.5-72t64-81.5T704 224q0-94-66-224l3 1l-1-1q90 41 160 83t138.5 100T1052 305.5t72.5 150.5t27.5 184"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FireExtinguisher(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 160q0-26-19-45t-45-19t-45 19t-19 45t19 45t45 19t45-19t19-45m896-32v320q0 16-12 25q-8 7-20 7q-4 0-7-1l-448-96q-11-2-18-11t-7-20H640v102q111 23 183.5 111T896 768v800q0 26-19 45t-45 19H320q-26 0-45-19t-19-45V768q0-106 62.5-190.5T480 463V352h-32q-59 0-115 23.5t-91.5 53t-66 66.5t-40.5 53.5t-14 24.5q-17 35-57 35q-16 0-29-7q-23-12-31.5-37T7 515q5-10 14.5-26T59 435.5t60.5-70t85-67T313 246q-25-42-25-86q0-66 47-113T448 0t113 47t47 113q0 33-14 64h302q0-11 7-20t18-11l448-96q3-1 7-1q12 0 20 7q12 9 12 25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Firefox(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M903 1760q-283 0-504.5-150.5T69 1211Q11 1080 2 910t26-332.5t111-312T318 23l-11 281q11-14 68-15.5t70 15.5q42-81 160.5-138T840 107q-54 45-119.5 148.5T662 419q25 8 62.5 13.5t63 7.5t68 4t50.5 3q15 5 9.5 45.5T885 568q-5 7-16.5 18.5T812 622t-101 34l15 189l-139-67q-18 43-7.5 81.5t36 66.5t65.5 41.5t81 6.5q51-9 98-34.5t83.5-45T1017 877q61 4 89.5 33t19.5 65q-1 2-2.5 5.5T1115 993t-18 15.5t-31.5 10.5t-46.5 1q-60 95-144.5 135.5T665 1185q74 61 162.5 82.5t168.5 6t154.5-52t128-87.5t80.5-104q43-91 39-192.5T1360.5 649T1282 524q87 38 137 79.5t77 112.5q15-170-57.5-343T1229 89q265 77 412 279.5T1792 886q2 127-40.5 255T1628 1379t-189 196t-247.5 135.5T903 1760"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FirstOrder(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1322 896q0 45-5 76l-236-14l224 78q-19 73-58 141l-214-103l177 158q-44 61-107 108l-157-178l103 215q-61 37-140 59l-79-228l14 240q-38 6-76 6t-76-6l14-238l-78 226q-74-19-140-59l103-215l-157 178q-59-43-108-108l178-158l-214 104q-39-69-58-141l224-79l-237 14q-5-42-5-76q0-35 5-77l238 14l-225-79q19-73 58-140l214 104l-177-159q46-61 107-108l158 178l-103-215q67-39 140-58l77 224l-13-236q36-6 75-6q38 0 76 6l-14 237l78-225q74 19 140 59L945 629l158-178q61 47 107 108l-177 159l213-104q37 62 58 141l-224 78l237-14q5 31 5 77m30 0q0-160-78.5-295.5t-213-214T768 308q-119 0-227 46.5t-186.5 125T230 667t-46 229q0 119 46 228t124.5 187.5t186.5 125t227 46.5q158 0 292.5-78.5t213-214T1352 896m73-383v766l-657 383l-657-383V513l657-383zM768 1719l708-412V484L768 73L60 484v823zm768-1271v896l-768 448L0 1344V448L768 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FiveHundredPx(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1401 1547l-6 6q-113 113-259 175q-154 64-317 64q-165 0-317-64q-148-63-259-175q-113-112-175-258q-42-103-54-189q-4-28 48-36q51-8 56 20q1 1 1 4q18 90 46 159q50 124 152 226q98 98 226 152q132 56 276 56q143 0 276-56q128-55 225-152l6-6q10-10 25-6q12 3 33 22q36 37 17 58M929 932l-66 66l63 63q21 21-7 49q-17 17-32 17q-10 0-19-10l-62-61l-66 66q-5 5-15 5q-15 0-31-16l-2-2q-18-15-18-29q0-7 8-17l66-65l-66-66q-16-16 14-45q18-18 31-18q6 0 13 5l65 66l65-65q18-17 48 13q27 27 11 44m471 57q0 118-46 228q-45 105-126 186q-80 80-187 126t-228 46t-228-46t-187-126q-82-82-125-186q-15-33-15-40h-1q-9-27 43-44q50-16 60 12q37 99 97 167h1V971q3-136 102-232q105-103 253-103q147 0 251 103t104 249q0 147-104.5 251T813 1343q-58 0-112-16q-28-11-13-61q16-51 44-43l14 3q14 3 33 6t30 3q104 0 176-71.5t72-174.5q0-101-72-171q-71-71-175-71q-107 0-178 80q-64 72-64 160v413q110 67 242 67q96 0 185-36.5t156-103.5t103.5-155t36.5-183q0-198-141-339q-140-140-339-140q-200 0-340 140q-53 53-77 87l-2 2q-8 11-13 15.5t-21.5 9.5t-38.5-3q-21-5-36.5-16.5T267 718V38q0-15 10.5-26.5T305 0h877q30 0 30 55t-30 55H371v483h1q40-42 102-84t108-61q109-46 231-46q121 0 228 46t187 126q81 81 126 186q46 112 46 229m-31-581q9 8 9 18t-5.5 18t-16.5 21q-26 26-39 26q-9 0-16-7q-106-91-207-133q-128-56-276-56q-133 0-262 49q-27 10-45-37q-9-25-8-38q3-16 16-20q130-57 299-57q164 0 316 64q137 58 235 152"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 128q0 72-64 110v1266q0 13-9.5 22.5T160 1536H96q-13 0-22.5-9.5T64 1504V238Q0 200 0 128q0-53 37.5-90.5T128 0t90.5 37.5T256 128m1472 64v763q0 25-12.5 38.5T1676 1021q-215 116-369 116q-61 0-123.5-22t-108.5-48t-115.5-48T817 997q-192 0-464 146q-17 9-33 9q-26 0-45-19t-19-45V346q0-32 31-55q21-14 79-43q236-120 421-120q107 0 200 29t219 88q38 19 88 19q54 0 117.5-21t110-47t88-47t54.5-21q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagCheckered(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 872V680q-181 16-384 117v185q205-96 384-110m0-418V257q-172 8-384 126v189q215-111 384-118m832 463V733q-235 116-384 71V580q-20-6-39-15q-5-3-33-17t-34.5-17t-31.5-15t-34.5-15.5t-32.5-13t-36-12.5t-35-8.5t-39.5-7.5t-39.5-4t-44-2q-23 0-49 3v222h19q102 0 192.5 29t197.5 82q19 9 39 15v188q42 17 91 17q120 0 293-92m0-427V301q-169 91-306 91q-45 0-78-8v196q148 42 384-90M256 128q0 35-17.5 64T192 238v1266q0 14-9 23t-23 9H96q-14 0-23-9t-9-23V238q-29-17-46.5-46T0 128q0-53 37.5-90.5T128 0t90.5 37.5T256 128m1472 64v763q0 39-35 57q-10 5-17 9q-218 116-369 116q-88 0-158-35l-28-14q-64-33-99-48t-91-29t-114-14q-102 0-235.5 44T353 1143q-15 9-33 9q-16 0-32-8q-32-19-32-56V346q0-35 31-55q35-21 78.5-42.5t114-52T632 147t155-19q112 0 209 31t209 86q38 19 89 19q122 0 310-112q22-12 31-17q31-16 62 2q31 20 31 55"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1600 917V301q-169 91-306 91q-82 0-145-32q-100-49-184-76.5T787 256q-173 0-403 127v599q245-113 433-113q55 0 103.5 7.5t98 26t77 31T1178 973l28 14q44 22 101 22q120 0 293-92M256 128q0 35-17.5 64T192 238v1266q0 14-9 23t-23 9H96q-14 0-23-9t-9-23V238q-29-17-46.5-46T0 128q0-53 37.5-90.5T128 0t90.5 37.5T256 128m1472 64v763q0 39-35 57q-10 5-17 9q-218 116-369 116q-88 0-158-35l-28-14q-64-33-99-48t-91-29t-114-14q-102 0-235.5 44T353 1143q-15 9-33 9q-16 0-32-8q-32-19-32-56V346q0-35 31-55q35-21 78.5-42.5t114-52T632 147t155-19q112 0 209 31t209 86q38 19 89 19q122 0 310-112q22-12 31-17q31-16 62 2q31 20 31 55"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flask(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1431 1320q56 89 21.5 152.5T1312 1536H160q-106 0-140.5-63.5T41 1320l503-793V128h-64q-26 0-45-19t-19-45t19-45t45-19h512q26 0 45 19t19 45t-19 45t-45 19h-64v399zM652 595l-272 429h712L820 595l-20-31V128H672v436z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flickr(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1248 0q119 0 203.5 84.5T1536 288v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0zM698 768q0-88-62-150t-150-62t-150 62t-62 150t62 150t150 62t150-62t62-150m564 0q0-88-62-150t-150-62t-150 62t-62 150t62 150t150 62t150-62t62-150"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FloppyO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 1408h768v-384H384zm896 0h128V512q0-14-10-38.5t-20-34.5l-281-281q-10-10-34-20t-39-10v416q0 40-28 68t-68 28H352q-40 0-68-28t-28-68V128H128v1280h128V992q0-40 28-68t68-28h832q40 0 68 28t28 68zM896 480V160q0-13-9.5-22.5T864 128H672q-13 0-22.5 9.5T640 160v320q0 13 9.5 22.5T672 512h192q13 0 22.5-9.5T896 480m640 32v928q0 40-28 68t-68 28H96q-40 0-68-28t-28-68V96q0-40 28-68T96 0h928q40 0 88 20t76 48l280 280q28 28 48 76t20 88"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1664 480v704q0 92-66 158t-158 66H224q-92 0-158-66T0 1184V224q0-92 66-158T224 0h320q92 0 158 66t66 158v32h672q92 0 158 66t66 158"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1536 1184V480q0-40-28-68t-68-28H736q-40 0-68-28t-28-68v-64q0-40-28-68t-68-28H224q-40 0-68 28t-28 68v960q0 40 28 68t68 28h1216q40 0 68-28t28-68m128-704v704q0 92-66 158t-158 66H224q-92 0-158-66T0 1184V224q0-92 66-158T224 0h320q92 0 158 66t66 158v32h672q92 0 158 66t66 158"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpen(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1879 824q0 31-31 66l-336 396q-43 51-120.5 86.5T1248 1408H160q-34 0-60.5-13T73 1352q0-31 31-66l336-396q43-51 120.5-86.5T704 768h1088q34 0 60.5 13t26.5 43m-343-344v160H704q-94 0-197 47.5T343 807L6 1203l-5 6q0-4-.5-12.5T0 1184V224q0-92 66-158T224 0h320q92 0 158 66t66 158v32h544q92 0 158 66t66 158"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpenO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1781 803q0-35-53-35H640q-40 0-85.5 21.5T483 842l-294 363q-18 24-18 40q0 35 53 35h1088q40 0 86-22t71-53l294-363q18-22 18-39M640 640h768V480q0-40-28-68t-68-28H736q-40 0-68-28t-28-68v-64q0-40-28-68t-68-28H224q-40 0-68 28t-28 68v853l256-315q44-53 116-87.5T640 640m1269 163q0 62-46 120l-295 363q-43 53-116 87.5t-140 34.5H224q-92 0-158-66T0 1184V224q0-92 66-158T224 0h320q92 0 158 66t66 158v32h544q92 0 158 66t66 158v160h192q54 0 99 24.5t67 70.5q15 32 15 68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Font(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M725 431L555 881q33 0 136.5 2t160.5 2q19 0 57-2q-87-253-184-452M0 1536l2-79q23-7 56-12.5t57-10.5t49.5-14.5t44.5-29t31-50.5l237-616L757 0h128q8 14 11 21l205 480q33 78 106 257.5t114 274.5q15 34 58 144.5t72 168.5q20 45 35 57q19 15 88 29.5t84 20.5q6 38 6 57q0 5-.5 13.5t-.5 12.5q-63 0-190-8t-191-8q-76 0-215 7t-178 8q0-43 4-78l131-28q1 0 12.5-2.5t15.5-3.5t14.5-4.5t15-6.5t11-8t9-11t2.5-14q0-16-31-96.5t-72-177.5t-42-100l-450-2q-26 58-76.5 195.5T382 1361q0 22 14 37.5t43.5 24.5t48.5 13.5t57 8.5t41 4q1 19 1 58q0 9-2 27q-58 0-174.5-10T236 1514q-8 0-26.5 4t-21.5 4q-80 14-188 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fonticons(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h1536v1536H0zm908 320l-12 33l75 83l-31 114l25 25l107-57l107 57l25-25l-31-114l75-83l-12-33h-95l-53-96h-32l-53 96zM641 483q32 0 44.5 16t11.5 63l174-21q0-55-17.5-92.5t-50.5-56t-69-25.5t-85-7q-133 0-199 57.5T384 600v72h-96v128h76q20 0 20 8v382q0 14-5 20t-18 7l-73 7v88h448v-86l-149-14q-6-1-8.5-1.5t-3.5-2.5t-.5-4t1-7t.5-10V800h191l38-128H574q-6 0-2-6t4-9v-80q0-27 1.5-40.5t7.5-28t19.5-20T641 483m607 829v-86l-54-9q-7-1-9.5-2.5t-2.5-3t1-7.5t1-12V672H909l-23 101l83 22q23 7 23 27v370q0 14-6 18.5t-20 6.5l-70 9v86z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FortAwesome(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 1008V784q0-16-16-16h-96q-16 0-16 16v224q0 16 16 16h96q16 0 16-16m512 0V784q0-16-16-16h-96q-16 0-16 16v224q0 16 16 16h96q16 0 16-16m512 32v752h-640v-320q0-80-56-136t-136-56t-136 56t-56 136v320H0v-752q0-16 16-16h96q16 0 16 16v112h128V528q0-16 16-16h96q16 0 16 16v112h128V528q0-16 16-16h96q16 0 16 16v112h128V528q0-6 2.5-9.5t8.5-5t9.5-2t11.5 0t9 .5V121q-32-15-32-50q0-23 16.5-39T832 16t38.5 16T887 71q0 35-32 50v17q45-10 83-10q21 0 59.5 7.5t54.5 7.5q17 0 47-7.5t37-7.5q16 0 16 16v210q0 15-35 21.5t-62 6.5q-18 0-54.5-7.5T945 367q-40 0-90 12v133q1 0 9-.5t11.5 0t9.5 2t8.5 5t2.5 9.5v112h128V528q0-16 16-16h96q16 0 16 16v112h128V528q0-16 16-16h96q16 0 16 16v624h128v-112q0-16 16-16h96q16 0 16 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Forumbee(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M934 22Q617 143 378 384.5T20 945Q0 856 0 769q0-208 102.5-384.5t278.5-279T765 3q82 0 169 19m269 119q93 65 164 155q-389 113-674.5 400.5T296 1373q-93-72-155-162q112-386 395-671t667-399M470 1475q115-356 379.5-622T1469 469q40 92 54 195q-292 120-516 345t-343 518q-103-14-194-52m1066 58q-193-50-367-115q-135 84-290 107q109-205 274-370.5T1522 879q-21 152-101 284q65 175 115 370"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Foursquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m968 434l37-194q5-23-9-40t-35-17H249q-23 0-38.5 17T195 237v1101q0 7 6 1l291-352q23-26 38-33.5t48-7.5h239q22 0 37-14.5t18-29.5q24-130 37-191q4-21-11.5-40T861 652H567q-29 0-48-19t-19-48v-42q0-29 19-47.5t48-18.5h346q18 0 35-13.5t20-29.5m227-222q-15 73-53.5 266.5t-69.5 350t-35 173.5q-6 22-9 32.5t-14 32.5t-24.5 33t-38.5 21t-58 10H622q-13 0-22 10q-8 9-426 494q-22 25-58.5 28.5T67 1658q-55-22-55-98V150q0-55 38-102.5T170 0h888q95 0 127 53t10 159m0 0l-158 790q4-17 35-173.5t69.5-350T1195 212"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FreeCodeCamp(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M453 1541q0 21-16 37.5t-37 16.5q-1 0-13-3q-63-15-162-140Q0 1168 0 776q0-341 213-614q39-51 95-103.5T402 6q19 0 35 13.5T453 52q0 27-63 90q-98 102-147 184q-119 199-119 449q0 281 123 491q50 85 136 173q2 3 14.5 16t19.5 21t17 20.5t14.5 23.5t4.5 21m1343-134q0 29-17.5 48.5T1732 1475H651q-26 0-45-19t-19-45q0-29 17.5-48.5T651 1343h1081q26 0 45 19t19 45m-215-611q0 134-67 233q-25 38-69.5 78.5T1361 1168q-16 10-27 10q-7 0-15-6t-8-12q0-9 19-30t42-46t42-67.5t19-88.5q0-76-35-130q-29-42-46-42q-3 0-3 5q0 12 7.5 35.5t7.5 36.5q0 22-21.5 35t-44.5 13q-66 0-66-76q0-15 1.5-44t1.5-44q0-25-10-46q-13-25-42-53.5t-51-28.5q-5 0-7 .5t-3.5 2.5t-1.5 6q0 2 16 26t16 54q0 37-19 68t-46 54t-53.5 46t-45.5 54t-19 68q0 98 42 160q29 43 79 63q16 5 17 10q1 2 1 5q0 16-18 16q-6 0-33-11q-119-43-195-139.5T786 853q0-55 24.5-115.5t60-115T941 514t59.5-113.5T1025 289q0-53-25-94q-29-48-56-64q-19-9-19-21q0-20 41-20q50 0 110 29q41 19 71 44.5t49.5 51T1230 277t22 69t16 80q0 1 3 17.5t4.5 25t5.5 25t9 27t11 21.5t14.5 16.5t18.5 5.5q23 0 37-14t14-37q0-25-20-67t-20-52t10-10q27 0 93 70q72 76 102.5 156t30.5 186m723 29q0 274-138 503q-19 32-48 72t-68 86.5t-81 77t-74 30.5q-16 0-31-15.5t-15-31.5q0-15 29-50.5t68.5-77t48.5-52.5q183-230 183-531q0-131-20.5-235T2085 390q-58-119-163-228q-2-3-13-13.5t-16.5-16.5t-15-17.5t-15-20T1853 76t-4-19q0-19 16-35.5T1900 5q70 0 196 169q98 131 146 273t60 314q2 42 2 64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FrownO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1134 1101q8 25-4 48.5t-37 31.5t-49-4t-32-38q-25-80-92.5-129.5T768 960t-151.5 49.5T524 1139q-8 26-31.5 38t-48.5 4q-26-8-38-31.5t-4-48.5q37-121 138-195t228-74t228 74t138 195M640 512q0 53-37.5 90.5T512 640t-90.5-37.5T384 512t37.5-90.5T512 384t90.5 37.5T640 512m512 0q0 53-37.5 90.5T1024 640t-90.5-37.5T896 512t37.5-90.5T1024 384t90.5 37.5T1152 512m256 256q0-130-51-248.5t-136.5-204t-204-136.5T768 128t-248.5 51t-204 136.5t-136.5 204T128 768t51 248.5t136.5 204t204 136.5t248.5 51t248.5-51t204-136.5t136.5-204t51-248.5m128 0q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FutbolO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m609 816l287-208l287 208l-109 336H719zM896 0q182 0 348 71t286 191t191 286t71 348t-71 348t-191 286t-286 191t-348 71t-348-71t-286-191t-191-286T0 896t71-348t191-286T548 71T896 0m619 1350q149-203 149-454v-3l-102 89l-240-224l63-323l134 12q-150-206-389-282l53 124l-287 159l-287-159l53-124q-239 76-389 282l135-12l62 323l-240 224l-102-89v3q0 251 149 454l30-132l326 40l139 298l-116 69q117 39 240 39t240-39l-116-69l139-298l326-40z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gamepad(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 832V704q0-14-9-23t-23-9H608V480q0-14-9-23t-23-9H448q-14 0-23 9t-9 23v192H224q-14 0-23 9t-9 23v128q0 14 9 23t23 9h192v192q0 14 9 23t23 9h128q14 0 23-9t9-23V864h192q14 0 23-9t9-23m576 64q0-53-37.5-90.5T1280 768t-90.5 37.5T1152 896t37.5 90.5t90.5 37.5t90.5-37.5T1408 896m256-256q0-53-37.5-90.5T1536 512t-90.5 37.5T1408 640t37.5 90.5T1536 768t90.5-37.5T1664 640m256 128q0 212-150 362t-362 150q-192 0-338-128H850q-146 128-338 128q-212 0-362-150T0 768t150-362t362-150h896q212 0 362 150t150 362"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gavel(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1739 1504q0 53-37 90l-107 108q-39 37-91 37q-53 0-90-37l-363-364q-38-36-38-90q0-53 43-96L800 896l-126 126q-14 14-34 14t-34-14q2 2 12.5 12t12.5 13t10 11.5t10 13.5t6 13.5t5.5 16.5t1.5 18q0 38-28 68q-3 3-16.5 18t-19 20.5T582 1243t-22 15.5t-22 9t-26 4.5q-40 0-68-28L36 836Q8 808 8 768q0-13 4.5-26t9-22T37 698t16.5-18.5t20.5-19T92 644q30-28 68-28q10 0 18 1.5t16.5 5.5t13.5 6t13.5 10t11.5 10t13 12.5t12 12.5q-14-14-14-34t14-34l348-348q14-14 34-14t34 14q-2-2-12.5-12T649 233t-10-11.5t-10-13.5t-6-13.5t-5.5-16.5t-1.5-18q0-38 28-68q3-3 16.5-18t19-20.5T698 37t22-15.5t22-9T768 8q40 0 68 28l408 408q28 28 28 68q0 13-4.5 26t-9 22t-15.5 22t-16.5 18.5t-20.5 19t-18 16.5q-30 28-68 28q-10 0-18-1.5t-16.5-5.5t-13.5-6t-13.5-10t-11.5-10t-13-12.5t-12-12.5q14 14 14 34t-14 34L896 800l256 256q43-43 96-43q52 0 91 37l363 363q37 39 37 91"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gbp(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1020 1009v367q0 14-9 23t-23 9H32q-14 0-23-9t-9-23v-150q0-13 9.5-22.5T32 1194h97V811H34q-14 0-23-9.5T2 779V648q0-14 9-23t23-9h95V393q0-171 123.5-282T567 0q185 0 335 125q9 8 10 20.5t-7 22.5L802 295q-9 11-22 12q-13 2-23-7q-5-5-26-19t-69-32t-93-18q-85 0-137 47t-52 123v215h305q13 0 22.5 9t9.5 23v131q0 13-9.5 22.5T685 811H380v379h414v-181q0-13 9-22.5t23-9.5h162q14 0 23 9.5t9 22.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Genderless(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1024 704q0-185-131.5-316.5T576 256T259.5 387.5T128 704t131.5 316.5T576 1152t316.5-131.5T1024 704m128 0q0 117-45.5 223.5t-123 184t-184 123T576 1280t-223.5-45.5t-184-123t-123-184T0 704t45.5-223.5t123-184t184-123T576 128t223.5 45.5t184 123t123 184T1152 704"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GetPocket(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1565 0q65 0 110 45.5t45 110.5v519q0 176-68 336t-182.5 275t-274 182.5T861 1536q-176 0-335.5-67.5T251 1286T68 1011T0 675V156Q0 92 46 46T156 0zM861 1064q47 0 82-33l404-388q37-35 37-85q0-49-34.5-83.5T1266 440q-47 0-82 33L861 783L538 473q-35-33-81-33q-49 0-83.5 34.5T339 558q0 51 36 85l405 388q33 33 81 33"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gg(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m672 576l384 384l-384 384L0 672L672 0l168 168l-96 96l-72-72l-480 480l480 480l193-193l-289-287zM1248 0l672 672l-672 672l-168-168l96-96l72 72l480-480l-480-480l-193 193l289 287l-96 96l-384-384z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GgCircle(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m717 1354l271-271l-279-279l-88 88l192 191l-96 96l-279-279l279-279l40 40l87-87l-127-128l-454 454zm358-8l454-454l-454-454l-271 271l279 279l88-88l-192-191l96-96l279 279l-279 279l-40-40l-87 88zm717-450q0 182-71 348t-191 286t-286 191t-348 71t-348-71t-286-191t-191-286T0 896t71-348t191-286T548 71T896 0t348 71t286 191t191 286t71 348"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gift(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M928 1164V448H608v716q0 25 18 38.5t46 13.5h192q28 0 46-13.5t18-38.5M472 320h195L541 159q-26-31-69-31q-40 0-68 28t-28 68t28 68t68 28m688-96q0-40-28-68t-68-28q-43 0-69 31L870 320h194q40 0 68-28t28-68m376 256v320q0 14-9 23t-23 9h-96v416q0 40-28 68t-68 28H224q-40 0-68-28t-28-68V832H32q-14 0-23-9t-9-23V480q0-14 9-23t23-9h440q-93 0-158.5-65.5T248 224t65.5-158.5T472 0q107 0 168 77l128 165L896 77q61-77 168-77q93 0 158.5 65.5T1288 224t-65.5 158.5T1064 448h440q14 0 23 9t9 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Git(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M531 1514q0-100-165-100q-158 0-158 104q0 101 172 101q151 0 151-105m-59-755q0-61-30-102t-89-41q-124 0-124 145q0 135 124 135q119 0 119-137m269-324v202q-36 12-79 22q16 43 16 84q0 127-73 216.5T408 1072q-40 8-59.5 27t-19.5 58q0 31 22.5 51.5t58 32t78.5 22t86 25.5t78.5 37.5t58 64T733 1488q0 304-363 304q-69 0-130-12.5t-116-41t-87.5-82T4 1529q0-165 182-225v-4q-67-41-67-126q0-109 63-137v-4q-72-24-119.5-108.5T15 759q0-139 95-231.5T345 435q96 0 178 47q98 0 218-47m318 881H837q4-45 4-134V573q0-94-4-128h222q-4 33-4 124v613q0 89 4 134m601-222v196q-71 39-174 39q-62 0-107-20t-70-50t-39.5-78t-18.5-92t-4-103V635h2v-4q-7 0-19-1t-18-1q-21 0-59 6V445h96v-76q0-54-6-89h227q-6 41-6 165h171v190q-15 0-43.5-2t-42.5-2h-85v365q0 131 87 131q61 0 109-33m-576-947q0 58-39 101.5T949 292q-58 0-98-43.5T811 147q0-59 39.5-103T949 0q58 0 96.5 44.5T1084 147"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M582 1180q0 66-93 66q-107 0-107-63q0-64 98-64q102 0 102 61m-36-466q0 85-74 85q-77 0-77-84q0-90 77-90q36 0 55 25.5t19 63.5m166-75V514q-78 29-135 29q-50-29-110-29q-86 0-145 57t-59 143q0 50 29.5 102t73.5 67v3q-38 17-38 85q0 53 41 77v3q-113 37-113 139q0 45 20 78.5t54 51t72 25.5t81 8q224 0 224-188q0-67-48-99t-126-46q-27-5-51.5-20.5T457 960q0-44 49-52q77-15 122-70t45-134q0-24-10-52q37-9 49-13m59 419h137q-2-27-2-82V589q0-46 2-69H771q3 23 3 71v392q0 50-3 75m509-16V921q-30 21-68 21q-53 0-53-82V635h52q9 0 26.5 1t26.5 1V520h-105q0-82 3-102h-140q4 24 4 55v47h-60v117q36-3 37-3q3 0 11 .5t12 .5v2h-2v217q0 37 2.5 64t11.5 56.5t24.5 48.5t43.5 31t66 12q64 0 108-24M924 336q0-36-24-63.5T840 245t-60.5 27t-24.5 64q0 36 25 62.5t60 26.5t59.5-27t24.5-62m612-48v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Github(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 0q209 0 385.5 103T1433 382.5T1536 768q0 251-146.5 451.5T1011 1497q-27 5-40-7t-13-30q0-3 .5-76.5t.5-134.5q0-97-52-142q57-6 102.5-18t94-39t81-66.5t53-105T1258 728q0-119-79-206q37-91-8-204q-28-9-81 11t-92 44l-38 24q-93-26-192-26t-192 26q-16-11-42.5-27T450 331.5T365 318q-45 113-8 204q-79 87-79 206q0 85 20.5 150T351 983t80.5 67t94 39t102.5 18q-39 36-49 103q-21 10-45 15t-57 5t-65.5-21.5T356 1146q-19-32-48.5-52t-49.5-24l-20-3q-21 0-29 4.5t-5 11.5t9 14t13 12l7 5q22 10 43.5 38t31.5 51l10 23q13 38 44 61.5t67 30t69.5 7t55.5-3.5l23-4q0 38 .5 88.5t.5 54.5q0 18-13 30t-40 7q-232-77-378.5-277.5T0 768q0-209 103-385.5T382.5 103T768 0M291 1103q3-7-7-12q-10-3-13 2q-3 7 7 12q9 6 13-2m31 34q7-5-2-16q-10-9-16-3q-7 5 2 16q10 10 16 3m30 45q9-7 0-19q-8-13-17-6q-9 5 0 18t17 7m42 42q8-8-4-19q-12-12-20-3q-9 8 4 19q12 12 20 3m57 25q3-11-13-16q-15-4-19 7t13 15q15 6 19-6m63 5q0-13-17-11q-16 0-16 11q0 13 17 11q16 0 16-11m58-10q-2-11-18-9q-16 3-14 15t18 8t14-14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GithubAlt(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 960q0 40-12.5 82t-43 76t-72.5 34t-72.5-34t-43-76t-12.5-82t12.5-82t43-76t72.5-34t72.5 34t43 76t12.5 82m640 0q0 40-12.5 82t-43 76t-72.5 34t-72.5-34t-43-76t-12.5-82t12.5-82t43-76t72.5-34t72.5 34t43 76t12.5 82m160 0q0-120-69-204t-187-84q-41 0-195 21q-71 11-157 11t-157-11q-152-21-195-21q-118 0-187 84t-69 204q0 88 32 153.5t81 103t122 60t140 29.5t149 7h168q82 0 149-7t140-29.5t122-60t81-103t32-153.5m224-176q0 207-61 331q-38 77-105.5 133t-141 86t-170 47.5t-171.5 22t-167 4.5q-78 0-142-3t-147.5-12.5t-152.5-30t-137-51.5t-121-81t-86-115Q0 992 0 784q0-237 136-396q-27-82-27-170q0-116 51-218q108 0 190 39.5T539 163q147-35 309-35q148 0 280 32q105-82 187-121t189-39q51 102 51 218q0 87-27 168q136 160 136 398"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GithubSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M519 1072q4-6-3-13q-9-7-14-2q-4 6 3 13q9 7 14 2m-28-41q-5-7-12-4q-6 4 0 12q7 8 12 5q6-4 0-13m-41-40q2-4-5-8q-7-2-8 2q-3 5 4 8q8 2 9-2m21 23q2-1 1.5-4.5t-3.5-5.5q-6-7-10-3t1 11q6 6 11 2m86 75q2-7-9-11q-9-3-13 4q-2 7 9 11q9 3 13-4m42 3q0-8-12-8q-10 0-10 8t11 8t11-8m39-7q-2-7-13-5t-9 9q2 8 12 6t10-10m642-317q0-212-150-362T768 256T406 406T256 768q0 167 98 300.5T606 1254q18 3 26.5-5t8.5-20q0-52-1-95q-6 1-15.5 2.5t-35.5 2t-48-4t-43.5-20T468 1073q-23-59-57-74q-2-1-4.5-3.5l-8-8l-7-9.5l4-7.5L415 967q6 0 15 2t30 15.5t33 35.5q16 28 37.5 42t43.5 14t38-3.5t30-9.5q7-47 33-69q-49-6-86-18.5t-73-39t-55.5-76T441 741q0-79 53-137q-24-62 5-136q19-6 54.5 7.5T614 505l26 16q58-17 128-17t128 17q11-7 28.5-18t55.5-26t57-9q29 74 5 136q53 58 53 137q0 57-14 100.5t-35.5 70T992 956t-62.5 26t-68.5 12q35 31 35 95q0 40-.5 89t-.5 51q0 12 8.5 20t26.5 5q154-52 252-185.5t98-300.5m256-480v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gitlab(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m104 642l792 1015l-868-630q-18-13-25-34.5T3 950zm462 0h660L896 1657zM368 30l198 612H104L302 30q8-23 33-23t33 23m1320 612l101 308q7 21 0 42.5t-25 34.5l-868 630zm0 0h-462l198-612q8-23 33-23t33 23z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gittip(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m773 1174l350-473q16-22 24.5-59t-6-85t-61.5-79q-40-26-83-25.5T923.5 470T869 515q-36 40-96 40q-59 0-95-40q-24-28-54.5-45T550 452.5T466 478q-46 31-60.5 79t-6 85t24.5 59zm763-406q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Glass(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1635 58q0 35-43 78L960 768v768h320q26 0 45 19t19 45t-19 45t-45 19H384q-26 0-45-19t-19-45t19-45t45-19h320V768L72 136Q29 93 29 58q0-23 18-36.5T85 4t43-4h1408q23 0 43 4t38 17.5t18 36.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Glide(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M866 387q0 27-13 94q-11 50-31.5 150T791 781q-2 11-4.5 12.5T773 796q-20 2-31 2q-58 0-84-49.5T632 635q0-88 35-174t103-124q28-14 51-14q28 0 36.5 16.5T866 387m486 424q0-14-39-75.5t-52-66.5q-21-8-34-8q-91 0-226 77l-2-2q3-22 27.5-135t24.5-178q0-233-242-233q-24 0-68 6q-94 17-168.5 89.5T461 452t-37 189q0 146 80.5 225T732 945q25 0 25 3t-1 5q-4 34-26 117q-14 52-51.5 101t-82.5 49q-42 0-42-47q0-24 10.5-47.5t25-39.5t29.5-28.5t26-20t11-8.5q0-3-7-10q-24-22-58.5-36.5T525 968q-35 0-63.5 34t-41 75t-12.5 75q0 88 51.5 142t138.5 54q82 0 155-53t117.5-126t65.5-153q6-22 15.5-66.5T966 883q3-12 14-18q118-60 227-60q48 0 127 18q1 1 4 1q5 0 9.5-4.5t4.5-8.5m184-523v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlideG(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M712 305q0-24-2-38.5t-8.5-30t-21-23T643 206q-39 0-78 23q-105 58-159 190.5T352 689q0 44 8.5 85.5T387 855t52.5 62.5T521 941q4 0 18 .5t20 0t16-3t15-8.5t7-16q16-77 48-231.5T693 451q19-91 19-146m754 656q0 7-7.5 13.5T1443 981l-6-1q-22-3-62-11t-72-12.5t-63-4.5q-167 0-351 93q-15 8-21 27q-10 36-24.5 105.5T821 1278q-23 91-70 179.5T638.5 1622T484 1745t-185 47q-135 0-214.5-83.5T5 1489q0-53 19.5-117t63-116.5T185 1203q38 0 120 33.5t83 61.5q0 1-16.5 12.5t-39.5 31t-46 44.5t-39 61t-16 74q0 33 16.5 53t48.5 20q45 0 85-31.5t66.5-78t48-105.5t32.5-107t16-90v-9q0-2-3.5-3.5t-8.5-1.5h-10l-10 .5l-6 .5q-227 0-352-122.5T29 698q0-108 34.5-221t96-210t156-167.5T520 10q52-9 106-9q374 0 374 360q0 98-38 273t-43 211l3 3q101-57 182.5-88t167.5-31q22 0 53 13q19 7 80 102.5t61 116.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globe(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 0q209 0 385.5 103T1433 382.5T1536 768t-103 385.5t-279.5 279.5T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0m274 521q-2 1-9.5 9.5T1019 540q2 0 4.5-5t5-11t3.5-7q6-7 22-15q14-6 52-12q34-8 51 11q-2-2 9.5-13t14.5-12q3-2 15-4.5t15-7.5l2-22q-12 1-17.5-7t-6.5-21q0 2-6 8q0-7-4.5-8t-11.5 1t-9 1q-10-3-15-7.5t-8-16.5t-4-15q-2-5-9.5-11t-9.5-10q-1-2-2.5-5.5t-3-6.5t-4-5.5t-5.5-2.5t-7 5t-7.5 10t-4.5 5q-3-2-6-1.5t-4.5 1t-4.5 3t-5 3.5q-3 2-8.5 3t-8.5 2q15-5-1-11q-10-4-16-3q9-4 7.5-12t-8.5-14h5q-1-4-8.5-8.5T1002 310t-13-6q-8-5-34-9.5t-33-.5q-5 6-4.5 10.5t4 14T925 331q1 6-5.5 13t-6.5 12q0 7 14 15.5t10 21.5q-3 8-16 16t-16 12q-5 8-1.5 18.5T914 456q2 2 1.5 4t-3.5 4.5t-5.5 4t-6.5 3.5l-3 2q-11 5-20.5-6T863 442q-7-25-16-30q-23-8-29 1q-5-13-41-26q-25-9-58-4q6-1 0-15q-7-15-19-12q3-6 4-17.5t1-13.5q3-13 12-23q1-1 7-8.5t9.5-13.5t.5-6q35 4 50-11q5-5 11.5-17t10.5-17q9-6 14-5.5t14.5 5.5t14.5 5q14 1 15.5-11t-7.5-20q12 1 3-17q-4-7-8-9q-12-4-27 5q-8 4 2 8q-1-1-9.5 10.5T801 218t-16-5q-1-1-5.5-13.5T770 186q-8 0-16 15q3-8-11-15t-24-8q19-12-8-27q-7-4-20.5-5t-19.5 4q-5 7-5.5 11.5t5 8T681 175t11.5 4t8.5 3q14 10 8 14q-2 1-8.5 3.5T689 204t-6 4q-3 4 0 14t-2 14q-5-5-9-17.5t-7-16.5q7 9-25 6l-10-1q-4 0-16 2t-20.5 1t-13.5-8q-4-8 0-20q1-4 4-2q-4-3-11-9.5t-10-8.5q-46 15-94 41q6 1 12-1q5-2 13-6.5t10-5.5q34-14 42-7l5-5q14 16 20 25q-7-4-30-1q-20 6-22 12q7 12 5 18q-4-3-11.5-10T498 211t-15-5q-16 0-22 1q-146 80-235 222q7 7 12 8q4 1 5 9t2.5 11t11.5-3q9 8 3 19q1-1 44 27q19 17 21 21q3 11-10 18q-1-2-9-9t-9-4q-3 5 .5 18.5T308 557q-7 0-9.5 16t-2.5 35.5t-1 23.5l2 1q-3 12 5.5 34.5T324 687q-13 3 20 43q6 8 8 9q3 2 12 7.5t15 10t10 10.5q4 5 10 22.5t14 23.5q-2 6 9.5 20t10.5 23q-1 0-2.5 1t-2.5 1q3 7 15.5 14t15.5 13q1 3 2 10t3 11t8 2q2-20-24-62q-15-25-17-29q-3-5-5.5-15.5T421 787q2 0 6 1.5t8.5 3.5t7.5 4t2 3q-3 7 2 17.5t12 18.5t17 19t12 13q6 6 14 19.5t0 13.5q9 0 20 10.5t17 19.5q5 8 8 26t5 24q2 7 8.5 13.5t12.5 9.5l16 8l13 7q5 2 18.5 10.5T642 1040q10 4 16 4t14.5-2.5t13.5-3.5q15-2 29 15t21 21q36 19 55 11q-2 1 .5 7.5t8 15.5t9 14.5t5.5 8.5q5 6 18 15t18 15q6-4 7-9q-3 8 7 20t18 10q14-3 14-32q-31 15-49-18q0-1-2.5-5.5t-4-8.5t-2.5-8.5t0-7.5t5-3q9 0 10-3.5t-2-12.5t-4-13q-1-8-11-20t-12-15q-5 9-16 8t-16-9q0 1-1.5 5.5t-1.5 6.5q-13 0-15-1q1-3 2.5-17.5t3.5-22.5q1-4 5.5-12t7.5-14.5t4-12.5t-4.5-9.5T775 954q-19 1-26 20q-1 3-3 10.5t-5 11.5t-9 7q-7 3-24 2t-24-5q-13-8-22.5-29t-9.5-37q0-10 2.5-26.5t3-25T652 858q3-2 9-9.5t10-10.5q2-1 4.5-1.5t4.5 0t4-1.5t3-6q-1-1-4-3q-3-3-4-3q7 3 28.5-1.5T735 823q15 11 22-2q0-1-2.5-9.5T754 798q5 27 29 9q3 3 15.5 5t17.5 5q3 2 7 5.5t5.5 4.5t5-.5t8.5-6.5q10 14 12 24q11 40 19 44q7 3 11 2t4.5-9.5t0-14T887 854l-1-8v-18l-1-8q-15-3-18.5-12t1.5-18.5t15-18.5q1-1 8-3.5t15.5-6.5t12.5-8q21-19 15-35q7 0 11-9q-1 0-5-3t-7.5-5t-4.5-2q9-5 2-16q5-3 7.5-11t7.5-10q9 12 21 2q8-8 1-16q5-7 20.5-10.5t18.5-9.5q7 2 8-2t1-12t3-12q4-5 15-9t13-5l17-11q3-4 0-4q18 2 31-11q10-11-6-20q3-6-3-9.5t-15-5.5q3-1 11.5-.5t10.5-1.5q15-10-7-16q-17-5-43 12m-163 877q206-36 351-189q-3-3-12.5-4.5t-12.5-3.5q-18-7-24-8q1-7-2.5-13t-8-9t-12.5-8t-11-7q-2-2-7-6t-7-5.5t-7.5-4.5t-8.5-2t-10 1l-3 1q-3 1-5.5 2.5t-5.5 3t-4 3t0 2.5q-21-17-36-22q-5-1-11-5.5t-10.5-7t-10-1.5t-11.5 7q-5 5-6 15t-2 13q-7-5 0-17.5t2-18.5q-3-6-10.5-4.5t-12 4.5t-11.5 8.5t-9 6.5t-8.5 5.5t-8.5 7.5q-3 4-6 12t-5 11q-2-4-11.5-6.5t-9.5-5.5q2 10 4 35t5 38q7 31-12 48q-27 25-29 40q-4 22 12 26q0 7-8 20.5t-7 21.5q0 6 2 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Google(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 658h725q12 67 12 128q0 217-91 387.5T1154.5 1440T768 1536q-157 0-299-60.5T224 1312T60.5 1067T0 768t60.5-299T224 224T469 60.5T768 0q300 0 515 201l-209 201Q951 283 768 283q-129 0-238.5 65T356 524.5T292 768t64 243.5T529.5 1188t238.5 65q87 0 160-24t120-60t82-82t51.5-87t22.5-78H768z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePlus(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1437 753q0 208-87 370.5t-248 254t-369 91.5q-149 0-285-58t-234-156t-156-234T0 736t58-285t156-234T448 61T733 3q286 0 491 192l-199 191Q908 273 733 273q-123 0-227.5 62T340 503.5T279 736t61 232.5T505.5 1137t227.5 62q83 0 152.5-23t114.5-57.5t78.5-78.5t49-83t21.5-74H733V631h692q12 63 12 122m867-122v210h-209v209h-210V841h-209V631h209V422h210v209z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePlusCircle(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M917 777q0-33-6-64H549v132h217q-12 76-74.5 120.5T549 1010q-99 0-169-71.5T310 768t70-170.5T549 526q93 0 153 59l104-101Q698 384 549 384q-160 0-272 112.5T165 768t112 271.5T549 1152q165 0 266.5-105T917 777m345 46h109V713h-109V603h-110v110h-110v110h110v110h110zm274-55q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePlusSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M917 777q0-26-6-64H549v132h217q-3 24-16.5 50T712 948t-66.5 44.5T549 1010q-99 0-169-71t-70-171t70-171t169-71q92 0 153 59l104-101Q698 384 549 384q-160 0-272 112.5T165 768t112 271.5T549 1152q165 0 266.5-105T917 777m345 46h109V713h-109V603h-110v110h-110v110h110v110h110zm274-535v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleWallet(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M409 672q33 0 52 26q266 364 362 774H377Q250 1031 10 723q-12-16-3-33.5T36 672zm559 357q-49 199-125 393q-79-310-256-594q40-221 44-449q211 340 337 650m99-709q235 324 384.5 698.5T1636 1792h-451q-41-665-553-1472zm693 576q0 424-101 812q-67-560-359-1083q-25-301-106-584q-4-16 5.5-28.5T1225 0h359q21 0 38.5 13t22.5 33q115 409 115 850"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraduationCap(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1774 708l18 316q4 69-82 128t-235 93.5t-323 34.5t-323-34.5t-235-93.5t-82-128l18-316l574 181q22 7 48 7t48-7zm530-324q0 23-22 31L1162 767q-4 1-10 1t-10-1L490 561q-43 34-71 111.5T385 851q63 36 63 109q0 69-58 107l58 433q2 14-8 25q-9 11-24 11H224q-15 0-24-11q-10-11-8-25l58-433q-58-38-58-107q0-73 65-111q11-207 98-330L22 415q-22-8-22-31t22-31L1142 1q4-1 10-1t10 1l1120 352q22 8 22 31"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grav(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1291 476q-15-17-35-8.5t-26 28.5t5 38q14 17 40 14.5t34-20.5t-18-52M895 722q-8 8-19.5 8t-18.5-8q-8-8-8-19t8-18q7-8 18.5-8t19.5 8q7 7 7 18t-7 19m165 74l-35 35q-12 13-29.5 13T965 831l-38-38q-12-13-12-30t12-30l35-35q12-12 29.5-12t30.5 12l38 39q12 12 12 29.5t-12 29.5M951 666q-7 8-18.5 8t-19.5-8q-7-8-7-19t7-19q8-8 19-8t19 8t8 19t-8 19m403-98q-34 64-107.5 85.5T1119 637q-38-28-61-66.5t-21-87.5t39-92t75.5-53t70.5 5t70 51q2 2 13 12.5t14.5 13.5t13 13.5T1345 449t10 15.5t8.5 18t4 18.5t1 21t-5 22t-9.5 24m201 482q3-20-8.5-34.5T1519 994t-33-17t-23-20q-40-71-84-98.5T1266 847q19-13 40-18.5t33-4.5l12 1q2-45-34-90q6-20 6.5-40.5T1321 664l-3-10q43-24 71-65t34-91q10-84-43-150.5T1243 271q-60-7-114 18.5t-82 74.5q-30 51-33.5 101t14.5 87t43.5 64t56.5 42q-45-4-88-36t-57-88q-28-108 32-222q-16-21-29-32q-50 0-89 19q19-24 42-37t36-14l13-1q0-50-13-78q-10-21-32.5-28.5t-47 3.5t-37.5 40q2-4 4-7q-7 28-6.5 75.5t19 117T923 492q-25 14-47 36q-35 16-85.5 70.5T706 700l-33 46q-90 34-181 125.5T417 1034q1 16 11 27q-15 12-30 30q-21 25-21 54t21.5 40t63.5-6q41-19 77-49.5t55-60.5q-2-2-6.5-5t-20.5-7.5t-33-3.5q23-5 51-12.5t40-10t27.5-6t26-4t23.5-.5q14 7 22-34q7-37 7-90q0-102-40-150q106 103 101 219q-1 29-15 50t-27 27l-13 6q-4 7-19 32t-26 45.5t-26.5 52t-25 61t-17 63t-6.5 66.5t10 63q-35-54-37-80q-22 24-34.5 39t-33.5 42t-30.5 46t-16.5 41t-.5 38t25.5 27q45 25 144-64t190.5-221.5T957 1070q86-52 145-115.5t86-119.5q47 93 154 178q104 83 167 80q39-2 46-43m239-154q0 182-71 348t-191 286t-286.5 191t-348.5 71t-348.5-71T262 1530T71 1244T0 896t71-348t191-286T548.5 71T897 0t348.5 71T1532 262t191 286t71 348"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GroupIcon(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M593 896q-162 5-265 128H194q-82 0-138-40.5T0 865q0-353 124-353q6 0 43.5 21t97.5 42.5T384 597q67 0 133-23q-5 37-5 66q0 139 81 256m1071 637q0 120-73 189.5t-194 69.5H523q-121 0-194-69.5T256 1533q0-53 3.5-103.5t14-109T300 1212t43-97.5t62-81t85.5-53.5T602 960q10 0 43 21.5t73 48t107 48t135 21.5t135-21.5t107-48t73-48t43-21.5q61 0 111.5 20t85.5 53.5t62 81t43 97.5t26.5 108.5t14 109t3.5 103.5M640 256q0 106-75 181t-181 75t-181-75t-75-181t75-181T384 0t181 75t75 181m704 384q0 159-112.5 271.5T960 1024T688.5 911.5T576 640t112.5-271.5T960 256t271.5 112.5T1344 640m576 225q0 78-56 118.5t-138 40.5h-134q-103-123-265-128q81-117 81-256q0-29-5-66q66 23 133 23q59 0 119-21.5t97.5-42.5t43.5-21q124 0 124 353m-128-609q0 106-75 181t-181 75t-181-75t-75-181t75-181t181-75t181 75t75 181"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hsquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1280 1216V320q0-26-19-45t-45-19h-128q-26 0-45 19t-19 45v320H512V320q0-26-19-45t-45-19H320q-26 0-45 19t-19 45v896q0 26 19 45t45 19h128q26 0 45-19t19-45V896h512v320q0 26 19 45t45 19h128q26 0 45-19t19-45m256-928v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HackerNews(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m809 876l266-499H963L806 689q-24 48-44 92l-42-92l-155-312H445l263 493v324h101zm727-588v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandGrabO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 128q-53 0-90.5 37.5T640 256v128h-32v-93q0-48-32-81.5T496 176q-46 0-79 33t-33 79v429l-32-30V515q0-48-32-81.5T240 400q-46 0-79 33t-33 79v224q0 47 35 82l310 296q39 39 39 102q0 26 19 45t45 19h640q26 0 45-19t19-45v-25q0-41 10-77l108-436q10-36 10-77V355q0-48-32-81.5t-80-33.5q-46 0-79 33t-33 79v32h-32V259q0-40-25-72.5t-64-40.5q-14-2-23-2q-46 0-79 33t-33 79v128h-32V262q0-51-32.5-89.5T781 129q-5-1-13-1m0-128q84 0 149 50q57-34 123-34q59 0 111 27t86 76q27-7 59-7q100 0 170 71.5t70 171.5v246q0 51-13 108l-109 436q-6 24-6 71q0 80-56 136t-136 56H576q-84 0-138-58.5T384 1207L76 911Q0 838 0 736V512q0-99 70.5-169.5T240 272q11 0 16 1q6-95 75.5-160T496 48q52 0 98 21Q666 0 768 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandLizardO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1151 0q61 0 116 28t91 77l572 781q118 159 118 359v355q0 80-56 136t-136 56h-384q-80 0-136-56t-56-136v-177l-286-143H448q-80 0-136-56t-56-136v-32q0-119 84.5-203.5T544 768h420l42-128H320q-100 0-173.5-67.5T65 406Q0 327 0 224v-32q0-80 56-136T192 0zm769 1600v-355q0-157-93-284l-573-781q-39-52-103-52H192q-26 0-45 19t-19 45q0 32 1.5 49.5T139 282t25 43q10-31 35.5-50t56.5-19h832v32H256q-26 0-45 19t-19 45q0 44 3 58q8 44 44 73t81 29h731q40 0 68 28t28 68q0 15-5 30l-64 192q-10 29-35 47.5T987 896H544q-66 0-113 47t-47 113v32q0 26 19 45t45 19h561q16 0 29 7l317 158q24 13 38.5 36t14.5 50v197q0 26 19 45t45 19h384q26 0 45-19t19-45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandOleft(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1376 1280h32V640h-32q-35 0-67.5-12t-62.5-37t-50-46t-49-54q-8-9-12-14q-72-81-112-145q-14-22-38-68q-1-3-10.5-22.5t-18.5-36t-20-35.5t-21.5-30.5T896 128q-71 0-115.5 30.5T736 256q0 43 15 84.5t33 68t33 55t15 48.5H256q-50 0-89 38.5T128 640q0 52 38 90t90 38h331q-15 17-25 47.5T552 871q0 69 53 119q-18 32-18 69t17.5 73.5T652 1185q-4 24-4 56q0 85 48.5 126t135.5 41q84 0 183-32t194-64t167-32m288-64q0-26-19-45t-45-19t-45 19t-19 45t19 45t45 19t45-19t19-45m128-576v640q0 53-37.5 90.5T1664 1408h-288q-59 0-223 59q-190 69-317 69q-142 0-230-77.5T519 1241l1-5q-61-76-61-178q0-22 3-43q-33-57-37-119H256q-105 0-180.5-76T0 639q0-103 76-179t180-76h374q-22-60-22-128q0-122 81.5-189T896 0q38 0 69.5 17.5t55 49.5t40.5 63t37 72t33 62q35 55 100 129q2 3 14 17t19 21.5t20.5 21.5t24 22.5t22.5 18t23.5 14t21.5 4.5h288q53 0 90.5 37.5T1792 640"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandOup(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1280 1600q0-26-19-45t-45-19t-45 19t-19 45t19 45t45 19t45-19t19-45m128-764q0-189-167-189q-26 0-56 5q-16-30-52.5-47.5T1059 587t-69 18q-50-53-119-53q-25 0-55.5 10T768 587V256q0-52-38-90t-90-38q-51 0-89.5 39T512 256v576q-20 0-48.5-15t-55-33t-68-33t-84.5-15q-67 0-97.5 44.5T128 896q0 24 139 90q44 24 65 37q64 40 145 112q81 71 106 101q57 69 57 140v32h640v-32q0-72 32-167t64-193.5t32-179.5m128-5q0 133-69 322q-59 164-59 223v288q0 53-37.5 90.5T1280 1792H640q-53 0-90.5-37.5T512 1664v-288q0-10-4.5-21.5t-14-23.5t-18-22.5t-22.5-24t-21.5-20.5t-21.5-19t-17-14q-74-65-129-100q-21-13-62-33t-72-37t-63-40.5t-49.5-55T0 896q0-125 67-206.5T256 608q68 0 128 22V256q0-104 76-180T639 0q105 0 181 75.5T896 256v169q62 4 119 37q21-3 43-3q101 0 178 60q139-1 219.5 85t80.5 227"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandPaperO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 128q-46 0-79 33t-33 79v656h-32V368q0-46-33-79t-79-33t-79 33t-33 79v784L358 947q-38-51-102-51q-53 0-90.5 37.5T128 1024q0 43 26 77l384 512q38 51 102 51h688q34 0 61-22t34-56l76-405q5-32 5-59V624q0-46-33-79t-79-33t-79 33t-33 79v272h-32V368q0-46-33-79t-79-33t-79 33t-33 79v528h-32V240q0-46-33-79t-79-33m0-128q68 0 125.5 35.5T1094 132q19-4 42-4q99 0 169.5 70.5T1376 368v17q105-6 180.5 64t75.5 175v498q0 40-8 83l-76 404q-14 79-76.5 131t-143.5 52H640q-60 0-114.5-27.5T435 1690L51 1178q-51-68-51-154q0-106 75-181t181-75q78 0 128 34V368q0-99 70.5-169.5T624 128q23 0 42 4q31-61 88.5-96.5T880 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandPeaceO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1288 647q60 0 107 23q141 63 141 226v177q0 94-23 186l-85 339q-21 86-90.5 140t-157.5 54H512q-106 0-181-75t-75-181v-401L17 507Q0 462 0 416q0-106 75-181t181-75q80 0 145.5 45.5T495 325l17 44V256q0-106 75-181T768 0t181 75t75 181v261q27-5 48-5q69 0 127.5 36.5T1288 647m-216-7q-33 0-60.5 18T970 706l-74 163l-71 155h55q50 0 90 31.5t50 80.5l154-338q10-20 10-46q0-46-33-79t-79-33m221 135q-22 0-40.5 8t-29 16t-23.5 29.5t-17 30.5t-17 37l-132 290q-10 20-10 46q0 46 33 79t79 33q33 0 60.5-18t41.5-48l160-352q9-18 9-38q0-50-32-81.5t-82-31.5M128 416q0 22 8 46l248 650v69l102-111q43-46 106-46h198l106-233V256q0-53-37.5-90.5T768 128t-90.5 37.5T640 256v640h-64L376 370q-14-37-47-59.5T256 288q-53 0-90.5 37.5T128 416m1052 1248q44 0 78.5-27t45.5-70l85-339q19-73 19-155v-91l-141 310q-17 38-53 61t-78 23q-53 0-93.5-34.5T994 1255q-44 57-114 57H672v-32h208q46 0 81-33t35-79t-31-79t-77-33H592q-49 0-82 36l-126 136v308q0 53 37.5 90.5T512 1664z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandPointerO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 128q-53 0-90.5 37.5T512 256v896L361 950q-41-54-107-54q-52 0-89 38t-37 90q0 43 26 77l384 512q38 51 102 51h718q22 0 39.5-13.5t22.5-34.5l92-368q24-96 24-194V837q0-41-28-71t-68-30t-68 28t-28 68h-32v-61q0-48-32-81.5t-80-33.5q-46 0-79 33t-33 79v64h-32v-90q0-55-37-94.5T928 608q-53 0-90.5 37.5T800 736v96h-32V262q0-55-37-94.5T640 128m0-128q107 0 181.5 77.5T896 262v220q22-2 32-2q99 0 173 69q47-21 99-21q113 0 184 87q27-7 56-7q94 0 159 67.5t65 161.5v217q0 116-28 225l-92 368q-16 64-68 104.5t-118 40.5H640q-60 0-114.5-27.5T435 1690L51 1178q-51-68-51-154q0-105 74.5-180.5T254 768q71 0 130 35V256q0-106 75-181T640 0m128 1408v-384h-32v384zm256 0v-384h-32v384zm256 0v-384h-32v384z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandScissorsO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1073 1536H896q-163 0-226-141q-23-49-23-102v-5q-62-30-98.5-88.5T512 1072q0-38 5-48H256q-106 0-181-75T0 768t75-181t181-75h113l-44-17q-74-28-119.5-93.5T160 256q0-106 75-181T416 0q46 0 91 17l628 239h401q106 0 181 75t75 181v668q0 88-54 157.5t-140 90.5l-339 85q-92 23-186 23m-49-711l-155 71l-163 74q-30 14-48 41.5t-18 60.5q0 46 33 79t79 33q26 0 46-10l338-154q-49-10-80.5-50t-31.5-90zm320 311q0-46-33-79t-79-33q-26 0-46 10l-290 132q-28 13-37 17t-30.5 17t-29.5 23.5t-16 29t-8 40.5q0 50 31.5 82t81.5 32q20 0 38-9l352-160q30-14 48-41.5t18-60.5m-232-752L462 136q-24-8-46-8q-53 0-90.5 37.5T288 256q0 40 22.5 73t59.5 47l526 200v64H256q-53 0-90.5 37.5T128 768t37.5 90.5T256 896h535l233-106V592q0-63 46-106l111-102zm-39 1024q82 0 155-19l339-85q43-11 70-45.5t27-78.5V512q0-53-37.5-90.5T1536 384h-308l-136 126q-36 33-36 82v296q0 46 33 77t79 31t79-35t33-81V672h32v208q0 70-57 114q52 8 86.5 48.5t34.5 93.5q0 42-23 78t-61 53l-310 141z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandSpockO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M459 1792q-77 0-137.5-47.5T242 1622l-101-401q-13-57-13-108q0-45-5-67L7 569q-7-27-7-57q0-93 62-161t155-78q17-85 82.5-139T452 80q83 0 148 51.5T685 264l83 348l103-428q20-81 85-132.5T1104 0q89 0 155.5 57.5T1340 202q92 10 152 79t60 162q0 24-7 59l-123 512q10-7 37.5-28.5T1498 956t35-23t41-20.5t41.5-11t49.5-5.5q105 0 180 74t75 179q0 62-28.5 118t-78.5 94l-507 380q-68 51-153 51zm645-1664q-38 0-68.5 24T996 214L832 896H705L560 294q-9-38-39.5-62T452 208q-48 0-80 33t-32 80q0 15 3 28l132 547h-26l-99-408q-9-37-40-62.5T241 400q-47 0-80 33t-33 79q0 14 3 26l116 478q7 28 9 86t10 88l100 401q8 32 34 52.5t59 20.5h694q42 0 76-26l507-379q56-43 56-110q0-52-37.5-88.5T1665 1024q-43 0-77 26l-307 230v-227q0-4 32-138t68-282t39-161q4-18 4-29q0-47-32-81t-79-34q-39 0-69.5 24t-39.5 62l-116 482h-26l150-624q3-14 3-28q0-48-31.5-82t-79.5-34"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandshakeO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 896q40 0 56-32t0-64t-56-32t-56 32t0 64t56 32m1473-58q-10-13-38.5-50t-41.5-54t-38-49t-42.5-53t-40.5-47t-45-49l-125 140q-83 94-208.5 92T880 670q-57-69-56.5-158T882 355l177-206q-22-11-51-16.5t-47.5-6t-56.5.5t-49 1q-92 0-158 66L539 352H384v544q5 0 21-.5t22 0t19.5 2T467 902t17.5 8.5T503 924l297 292q115 111 227 111q78 0 125-47q57 20 112.5-8t72.5-85q74 6 127-44q20-18 36-45.5t14-50.5q10 10 43 10q43 0 77-21t49.5-53t12-71.5T1665 838m159 58h96V384h-93l-157-180q-66-76-169-76h-167q-89 0-146 67L979 438q-28 33-28 75t27 75q43 51 110 52t111-49l193-218q25-23 53.5-21.5t47 27t8.5 56.5q16 19 56 63t60 68q29 36 82.5 105.5T1764 756q52 66 60 140m288 0q40 0 56-32t0-64t-56-32t-56 32t0 64t56 32m192-576v640q0 26-19 45t-45 19h-434q-27 65-82 106.5t-125 51.5q-33 48-80.5 81.5T1416 1309q-42 53-104.5 81.5T1183 1415q-60 34-126 39.5t-127.5-14t-117-53.5t-103.5-81l-287-282H64q-26 0-45-19T0 960V288q0-26 19-45t45-19h421q14-14 47-48t47.5-48t44-40T674 50.5T725 25t62-19.5T855 0h117q99 0 181 56q82-56 181-56h167q35 0 67 6t56.5 14.5T1676 47t44.5 31t43 39.5t39 42t41 48T1885 256h355q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hashtag(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m959 896l64-256H769l-64 256zm768-504l-56 224q-7 24-31 24h-327l-64 256h311q15 0 25 12q10 14 6 28l-56 224q-5 24-31 24h-327l-81 328q-7 24-31 24H841q-16 0-26-12q-9-12-6-28l78-312H633l-81 328q-7 24-31 24H296q-15 0-25-12q-9-12-6-28l78-312H32q-15 0-25-12q-9-12-6-28l56-224q7-24 31-24h327l64-256H168q-15 0-25-12q-10-14-6-28l56-224q5-24 31-24h327l81-328q7-24 32-24h224q15 0 25 12q9 12 6 28l-78 312h254l81-328q7-24 32-24h224q15 0 25 12q9 12 6 28l-78 312h311q15 0 25 12q9 12 6 28"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HddO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1040 960q0 33-23.5 56.5T960 1040t-56.5-23.5T880 960t23.5-56.5T960 880t56.5 23.5T1040 960m256 0q0 33-23.5 56.5T1216 1040t-56.5-23.5T1136 960t23.5-56.5T1216 880t56.5 23.5T1296 960m112 160V800q0-13-9.5-22.5T1376 768H160q-13 0-22.5 9.5T128 800v320q0 13 9.5 22.5t22.5 9.5h1216q13 0 22.5-9.5t9.5-22.5M178 640h1180l-157-482q-4-13-16-21.5t-26-8.5H377q-14 0-26 8.5T335 158zm1358 160v320q0 66-47 113t-113 47H160q-66 0-113-47T0 1120V800q0-25 16-75l197-606q17-53 63-86T377 0h782q55 0 101 33t63 86l197 606q16 50 16 75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Header(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1650 1536q-44 0-132.5-3.5T1384 1529q-44 0-132 3.5t-132 3.5q-24 0-37-20.5t-13-45.5q0-31 17-46t39-17t51-7t45-15q33-21 33-140l-1-391q0-21-1-31q-13-4-50-4H528q-38 0-51 4q-1 10-1 31l-1 371q0 142 37 164q16 10 48 13t57 3.5t45 15t20 45.5q0 26-12.5 48t-36.5 22q-47 0-139.5-3.5T355 1529q-43 0-128 3.5t-127 3.5q-23 0-35.5-21T52 1470q0-30 15.5-45t36-17.5t47.5-7.5t42-15q33-23 33-143l-1-57V372q0-3 .5-26t0-36.5T224 271t-3.5-42t-6.5-36.5t-11-31.5t-16-18q-15-10-45-12t-53-2t-41-14t-18-45q0-26 12-48T78 0q46 0 138.5 3.5T355 7q42 0 126.5-3.5T608 0q25 0 37.5 22T658 70q0 30-17 43.5T602.5 128t-49.5 4t-43 13q-35 21-35 160l1 320q0 21 1 32q13 3 39 3h699q25 0 38-3q1-11 1-32l1-320q0-139-35-160q-18-11-58.5-12.5t-66-13T1070 70q0-26 12.5-48T1120 0q44 0 132 3.5T1384 7q43 0 129-3.5T1642 0q25 0 37.5 22t12.5 48q0 30-17.5 44t-40 14.5t-51.5 3t-44 12.5q-35 23-35 161l1 943q0 119 34 140q16 10 46 13.5t53.5 4.5t41.5 15.5t18 44.5q0 26-12 48t-36 22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Headphones(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1664 758q0 166-60 314l-20 49l-185 33q-22 83-90.5 136.5T1152 1344v32q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23V800q0-14 9-23t23-9h64q14 0 23 9t9 23v32q71 0 130 35.5t93 95.5l68-12q29-95 29-193q0-148-88-279t-236.5-209T832 192t-315.5 78T280 479t-88 279q0 98 29 193l68 12q34-60 93-95.5T512 832v-32q0-14 9-23t23-9h64q14 0 23 9t9 23v576q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23v-32q-88 0-156.5-53.5T265 1154l-185-33l-20-49Q0 924 0 758q0-151 67-291t179-242.5T512 61T832 0t320 61t266 163.5T1597 467t67 291"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M896 1536q-26 0-44-18L228 916q-10-8-27.5-26T145 824.5T77 727T23.5 606T0 468q0-220 127-344T478 0q62 0 126.5 21.5t120 58T820 148t76 68q36-36 76-68t95.5-68.5t120-58T1314 0q224 0 351 124t127 344q0 221-229 450l-623 600q-18 18-44 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1664 468q0-81-21.5-143t-55-98.5T1506 167t-94-31t-98-8t-112 25.5t-110.5 64t-86.5 72t-60 61.5q-18 22-49 22t-49-22q-24-28-60-61.5t-86.5-72t-110.5-64T478 128t-98 8t-94 31t-81.5 59.5t-55 98.5T128 468q0 168 187 355l581 560l580-559q188-188 188-356m128 0q0 221-229 450l-623 600q-18 18-44 18t-44-18L228 916q-10-8-27.5-26T145 824.5T77 727T23.5 606T0 468q0-220 127-344T478 0q62 0 126.5 21.5t120 58T820 148t76 68q36-36 76-68t95.5-68.5t120-58T1314 0q224 0 351 124t127 344"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heartbeat(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1280 896h305q-5 6-10 10.5t-9 7.5l-3 4l-623 600q-18 18-44 18t-44-18L228 916q-5-2-21-20h369q22 0 39.5-13.5T638 848l70-281l190 667q6 20 23 33t39 13q21 0 38-13t23-33l146-485l56 112q18 35 57 35m512-428q0 145-103 300h-369l-111-221q-8-17-25.5-27t-36.5-8q-45 5-56 46L962 988L766 302q-6-20-23.5-33T703 256t-39 13.5t-22 34.5L526 768H103Q0 613 0 468q0-220 127-344T478 0q62 0 126.5 21.5t120 58T820 148t76 68q36-36 76-68t95.5-68.5t120-58T1314 0q224 0 351 124t127 344"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func History(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1536 768q0 156-61 298t-164 245t-245 164t-298 61q-172 0-327-72.5T177 1259q-7-10-6.5-22.5t8.5-20.5l137-138q10-9 25-9q16 2 23 12q73 95 179 147t225 52q104 0 198.5-40.5T1130 1130t109.5-163.5T1280 768t-40.5-198.5T1130 406T966.5 296.5T768 256q-98 0-188 35.5T420 393l137 138q31 30 14 69q-17 40-59 40H64q-26 0-45-19T0 576V128q0-42 40-59q39-17 69 14l130 129Q346 111 483.5 55.5T768 0q156 0 298 61t245 164t164 245t61 298M896 480v448q0 14-9 23t-23 9H544q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h224V480q0-14 9-23t23-9h64q14 0 23 9t9 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1408 768v480q0 26-19 45t-45 19H960V928H704v384H320q-26 0-45-19t-19-45V768q0-1 .5-3t.5-3l575-474l575 474q1 2 1 6m223-69l-62 74q-8 9-21 11h-3q-13 0-21-7L832 200L140 777q-12 8-24 7q-13-2-21-11l-62-74q-8-10-7-23.5T37 654L756 55q32-26 76-26t76 26l244 204V64q0-14 9-23t23-9h192q14 0 23 9t9 23v408l219 182q10 8 11 21.5t-7 23.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HospitalO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 1312v64q0 13-9.5 22.5T352 1408h-64q-13 0-22.5-9.5T256 1376v-64q0-13 9.5-22.5t22.5-9.5h64q13 0 22.5 9.5t9.5 22.5m0-256v64q0 13-9.5 22.5T352 1152h-64q-13 0-22.5-9.5T256 1120v-64q0-13 9.5-22.5t22.5-9.5h64q13 0 22.5 9.5t9.5 22.5m256 0v64q0 13-9.5 22.5T608 1152h-64q-13 0-22.5-9.5T512 1120v-64q0-13 9.5-22.5t22.5-9.5h64q13 0 22.5 9.5t9.5 22.5M384 800v64q0 13-9.5 22.5T352 896h-64q-13 0-22.5-9.5T256 864v-64q0-13 9.5-22.5T288 768h64q13 0 22.5 9.5T384 800m768 512v64q0 13-9.5 22.5t-22.5 9.5h-64q-13 0-22.5-9.5t-9.5-22.5v-64q0-13 9.5-22.5t22.5-9.5h64q13 0 22.5 9.5t9.5 22.5m-256-256v64q0 13-9.5 22.5T864 1152h-64q-13 0-22.5-9.5T768 1120v-64q0-13 9.5-22.5t22.5-9.5h64q13 0 22.5 9.5t9.5 22.5M640 800v64q0 13-9.5 22.5T608 896h-64q-13 0-22.5-9.5T512 864v-64q0-13 9.5-22.5T544 768h64q13 0 22.5 9.5T640 800m512 256v64q0 13-9.5 22.5t-22.5 9.5h-64q-13 0-22.5-9.5t-9.5-22.5v-64q0-13 9.5-22.5t22.5-9.5h64q13 0 22.5 9.5t9.5 22.5M896 800v64q0 13-9.5 22.5T864 896h-64q-13 0-22.5-9.5T768 864v-64q0-13 9.5-22.5T800 768h64q13 0 22.5 9.5T896 800m256 0v64q0 13-9.5 22.5T1120 896h-64q-13 0-22.5-9.5T1024 864v-64q0-13 9.5-22.5t22.5-9.5h64q13 0 22.5 9.5t9.5 22.5m-256 864h384V512h-256v32q0 40-28 68t-68 28H480q-40 0-68-28t-28-68v-32H128v1152h384v-224q0-13 9.5-22.5t22.5-9.5h320q13 0 22.5 9.5t9.5 22.5zm0-1184V160q0-13-9.5-22.5T864 128h-64q-13 0-22.5 9.5T768 160v96H640v-96q0-13-9.5-22.5T608 128h-64q-13 0-22.5 9.5T512 160v320q0 13 9.5 22.5T544 512h64q13 0 22.5-9.5T640 480v-96h128v96q0 13 9.5 22.5T800 512h64q13 0 22.5-9.5T896 480m512-32v1280q0 26-19 45t-45 19H64q-26 0-45-19t-19-45V448q0-26 19-45t45-19h320V96q0-40 28-68t68-28h448q40 0 68 28t28 68v288h320q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hourglass(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1504 1600q14 0 23 9t9 23v128q0 14-9 23t-23 9H32q-14 0-23-9t-9-23v-128q0-14 9-23t23-9zm-1374-64q3-55 16-107t30-95t46-87t53.5-76t64.5-69.5t66-60t70.5-55T543 939t65-43q-43-28-65-43t-66.5-47.5t-70.5-55t-66-60t-64.5-69.5t-53.5-76t-46-87t-30-95t-16-107h1276q-3 55-16 107t-30 95t-46 87t-53.5 76t-64.5 69.5t-66 60t-70.5 55T993 853t-65 43q43 28 65 43t66.5 47.5t70.5 55t66 60t64.5 69.5t53.5 76t46 87t30 95t16 107zM1504 0q14 0 23 9t9 23v128q0 14-9 23t-23 9H32q-14 0-23-9t-9-23V32Q0 18 9 9t23-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HourglassO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1408 128q0 261-106.5 461.5T1035 896q160 106 266.5 306.5T1408 1664h96q14 0 23 9t9 23v64q0 14-9 23t-23 9H32q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h96q0-261 106.5-461.5T501 896Q341 790 234.5 589.5T128 128H32q-14 0-23-9T0 96V32Q0 18 9 9t23-9h1472q14 0 23 9t9 23v64q0 14-9 23t-23 9zM874 836q77-29 149-92.5T1152.5 591t92.5-210t35-253H256q0 132 35 253t92.5 210T513 743.5T662 836q19 7 30.5 23.5T704 896t-11.5 36.5T662 956q-77 29-149 92.5T383.5 1201T291 1411t-35 253h1024q0-132-35-253t-92.5-210t-129.5-152.5T874 956q-19-7-30.5-23.5T832 896t11.5-36.5T874 836"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HourglassOne(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1408 128q0 261-106.5 461.5T1035 896q160 106 266.5 306.5T1408 1664h96q14 0 23 9t9 23v64q0 14-9 23t-23 9H32q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h96q0-261 106.5-461.5T501 896Q341 790 234.5 589.5T128 128H32q-14 0-23-9T0 96V32Q0 18 9 9t23-9h1472q14 0 23 9t9 23v64q0 14-9 23t-23 9zm-128 0H256q0 66 9 128h1006q9-61 9-128m0 1536q0-130-34-249.5t-90.5-208t-126.5-152T883 960H653q-76 31-146 94.5t-126.5 152t-90.5 208t-34 249.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HourglassThree(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1408 128q0 261-106.5 461.5T1035 896q160 106 266.5 306.5T1408 1664h96q14 0 23 9t9 23v64q0 14-9 23t-23 9H32q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h96q0-261 106.5-461.5T501 896Q341 790 234.5 589.5T128 128H32q-14 0-23-9T0 96V32Q0 18 9 9t23-9h1472q14 0 23 9t9 23v64q0 14-9 23t-23 9zM874 836q77-29 149-92.5T1152.5 591t92.5-210t35-253H256q0 132 35 253t92.5 210T513 743.5T662 836q19 7 30.5 23.5T704 896t-11.5 36.5T662 956q-137 51-244 196h700q-107-145-244-196q-19-7-30.5-23.5T832 896t11.5-36.5T874 836"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HourglassTwo(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1408 128q0 261-106.5 461.5T1035 896q160 106 266.5 306.5T1408 1664h96q14 0 23 9t9 23v64q0 14-9 23t-23 9H32q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h96q0-261 106.5-461.5T501 896Q341 790 234.5 589.5T128 128H32q-14 0-23-9T0 96V32Q0 18 9 9t23-9h1472q14 0 23 9t9 23v64q0 14-9 23t-23 9zm-128 0H256q0 206 85 384h854q85-178 85-384m-57 1216q-54-141-145.5-241.5T883 960H653q-103 42-194.5 142.5T313 1344z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Houzz(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m512 1191l512-295v591l-512 296zM0 896v591l512-296zM512 9v591L0 896V305zm0 591l512-295v591z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HtmlFive(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1130 469l16-175H262l47 534h612l-22 228l-197 53l-196-53l-13-140H318l22 278l362 100h4v-1l359-99l50-544H471l-15-181zM0 0h1408l-128 1438l-578 162l-574-162z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Icursor(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 128q-320 0-320 224v416h128v128H512v544q0 224 320 224h64v128h-64q-272 0-384-146q-112 146-384 146H0v-128h64q320 0 320-224V896H256V768h128V352q0-224-320-224H0V0h64q272 0 384 146Q560 0 832 0h64v128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IdBadge(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1024 1258q0 64-37 107t-91 43H384q-54 0-91-43t-37-107t9-118t29.5-104t61-78.5T452 929q80 75 188 75t188-75q56 0 96.5 28.5t61 78.5t29.5 104t9 118M870 739q0 94-67.5 160.5T640 966t-162.5-66.5T410 739t67.5-160.5T640 512t162.5 66.5T870 739m282 893V256H128v1376q0 13 9.5 22.5t22.5 9.5h960q13 0 22.5-9.5t9.5-22.5m128-1472v1472q0 66-47 113t-113 47H160q-66 0-113-47T0 1632V160Q0 94 47 47T160 0h352v96q0 14 9 23t23 9h192q14 0 23-9t9-23V0h352q66 0 113 47t47 113"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ils(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M992 496v496q0 14-9 23t-23 9H800q-14 0-23-9t-9-23V496q0-112-80-192t-192-80H224v1152q0 14-9 23t-23 9H32q-14 0-23-9t-9-23V32Q0 18 9 9t23-9h464q135 0 249 66.5T925.5 247T992 496m384-464v880q0 135-66.5 249T1129 1341.5T880 1408H416q-14 0-23-9t-9-23V416q0-14 9-23t23-9h160q14 0 23 9t9 23v768h272q112 0 192-80t80-192V32q0-14 9-23t23-9h160q14 0 23 9t9 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Image(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 448q0 80-56 136t-136 56t-136-56t-56-136t56-136t136-56t136 56t56 136m1024 384v448H256v-192l320-320l160 160l512-512zm96-704H160q-13 0-22.5 9.5T128 160v1216q0 13 9.5 22.5t22.5 9.5h1600q13 0 22.5-9.5t9.5-22.5V160q0-13-9.5-22.5T1760 128m160 32v1216q0 66-47 113t-113 47H160q-66 0-113-47T0 1376V160Q0 94 47 47T160 0h1600q66 0 113 47t47 113"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Imdb(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M922 669v182q0 4 .5 15t0 15l-1.5 12l-3.5 11.5l-6.5 7.5l-11 5.5l-16 1.5V610q9 0 16 1t11 5t6.5 5.5t3.5 9.5t1 10.5zm316 96v121q0 1 .5 12.5t0 15.5t-2.5 11.5t-7.5 10.5t-13.5 3q-9 0-14-9q-4-10-4-165v-24.5l1.5-8.5l3.5-7l5-5.5l8-1.5q6 0 10 1.5t6.5 4.5t4 6t2 8.5t.5 8zM180 1001h122V529H180zm434 0h106V529H561l-28 221q-20-148-32-221H343v472h107V689l45 312h76l43-319zm425-305q0-67-5-90q-3-16-11-28.5t-17-20.5t-25-14t-26.5-8.5t-31-4t-29-1.5H762v472h56q169 1 197-24.5t25-180.5q-1-62-1-100m317 197V760q0-29-2-45t-9.5-33.5t-24.5-25t-46-7.5q-46 0-77 34V529h-117v472h110l7-30q30 36 77 36q50 0 66-30.5t16-83.5m180-733v1216q0 66-47 113t-113 47H160q-66 0-113-47T0 1376V160Q0 94 47 47T160 0h1216q66 0 113 47t47 113"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inbox(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1023 704h316q-1-3-2.5-8.5t-2.5-7.5l-212-496H414L202 688q-1 3-2.5 8.5T197 704h316l95 192h320zm513 30v482q0 26-19 45t-45 19H64q-26 0-45-19t-19-45V734q0-62 25-123L263 59q10-25 36.5-42T352 0h832q26 0 52.5 17t36.5 42l238 552q25 61 25 123"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Indent(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M352 704q0 14-9 23L55 1015q-9 9-23 9q-13 0-22.5-9.5T0 992V416q0-13 9.5-22.5T32 384q14 0 23 9l288 288q9 9 9 23m1440 480v192q0 13-9.5 22.5t-22.5 9.5H32q-13 0-22.5-9.5T0 1376v-192q0-13 9.5-22.5T32 1152h1728q13 0 22.5 9.5t9.5 22.5m0-384v192q0 13-9.5 22.5t-22.5 9.5H672q-13 0-22.5-9.5T640 992V800q0-13 9.5-22.5T672 768h1088q13 0 22.5 9.5t9.5 22.5m0-384v192q0 13-9.5 22.5T1760 640H672q-13 0-22.5-9.5T640 608V416q0-13 9.5-22.5T672 384h1088q13 0 22.5 9.5t9.5 22.5m0-384v192q0 13-9.5 22.5T1760 256H32q-13 0-22.5-9.5T0 224V32Q0 19 9.5 9.5T32 0h1728q13 0 22.5 9.5T1792 32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Industry(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 0q26 0 45 19t19 45v891l536-429q17-14 40-14q26 0 45 19t19 45v379l536-429q17-14 40-14q26 0 45 19t19 45v1152q0 26-19 45t-45 19H64q-26 0-45-19t-19-45V64q0-26 19-45T64 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Info(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 1216v128q0 26-19 45t-45 19H64q-26 0-45-19t-19-45v-128q0-26 19-45t45-19h64V768H64q-26 0-45-19T0 704V576q0-26 19-45t45-19h384q26 0 45 19t19 45v576h64q26 0 45 19t19 45M512 64v192q0 26-19 45t-45 19H192q-26 0-45-19t-19-45V64q0-26 19-45t45-19h256q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoCircle(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1024 1248v-160q0-14-9-23t-23-9h-96V544q0-14-9-23t-23-9H544q-14 0-23 9t-9 23v160q0 14 9 23t23 9h96v320h-96q-14 0-23 9t-9 23v160q0 14 9 23t23 9h448q14 0 23-9t9-23M896 352V192q0-14-9-23t-23-9H672q-14 0-23 9t-9 23v160q0 14 9 23t23 9h192q14 0 23-9t9-23m640 416q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inr(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M898 342v102q0 14-9 23t-23 9H698q-23 144-129 234T293 820q167 178 459 536q14 16 4 34q-8 18-29 18H532q-16 0-25-12Q201 1029 9 825q-9-9-9-22V676q0-13 9.5-22.5T32 644h112q132 0 212.5-43T459 476H32q-14 0-23-9t-9-23V342q0-14 9-23t23-9h413q-57-113-268-113H32q-13 0-22.5-9.5T0 165V32Q0 18 9 9t23-9h832q14 0 23 9t9 23v102q0 14-9 23t-23 9H631q47 61 64 144h171q14 0 23 9t9 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Instagram(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1024 768q0-106-75-181t-181-75t-181 75t-75 181t75 181t181 75t181-75t75-181m138 0q0 164-115 279t-279 115t-279-115t-115-279t115-279t279-115t279 115t115 279m108-410q0 38-27 65t-65 27t-65-27t-27-65t27-65t65-27t65 27t27 65M768 138q-7 0-76.5-.5t-105.5 0t-96.5 3t-103 10T315 169q-50 20-88 58t-58 88q-11 29-18.5 71.5t-10 103t-3 96.5t0 105.5t.5 76.5t-.5 76.5t0 105.5t3 96.5t10 103T169 1221q20 50 58 88t88 58q29 11 71.5 18.5t103 10t96.5 3t105.5 0t76.5-.5t76.5.5t105.5 0t96.5-3t103-10t71.5-18.5q50-20 88-58t58-88q11-29 18.5-71.5t10-103t3-96.5t0-105.5t-.5-76.5t.5-76.5t0-105.5t-3-96.5t-10-103T1367 315q-20-50-58-88t-88-58q-29-11-71.5-18.5t-103-10t-96.5-3t-105.5 0t-76.5.5m768 630q0 229-5 317q-10 208-124 322t-322 124q-88 5-317 5t-317-5q-208-10-322-124T5 1085q-5-88-5-317t5-317q10-208 124-322T451 5q88-5 317-5t317 5q208 10 322 124t124 322q5 88 5 317"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InternetExplorer(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 937q0 56-7 104H634q0 146 109.5 244.5T1001 1384q99 0 185.5-46.5T1323 1207h423q-56 159-170.5 281T1308 1676.5T987 1743q-187 0-356-83q-228 116-394 116q-237 0-237-263q0-115 45-275q17-60 109-229q199-360 475-606q-184 79-427 354q63-274 283.5-449.5T987 132q30 0 45 1q255-117 433-117q64 0 116 13t94.5 40.5T1742 146t24 115q0 116-75 286q101 182 101 390m-70-640q0-83-53-132t-137-49q-108 0-254 70q121 47 222.5 131.5T1671 513q51-135 51-216M128 1534q0 86 48.5 132.5T311 1713q115 0 266-83q-122-72-213.5-183T226 1202q-98 205-98 332m504-713h728q-5-142-113-237t-251-95q-144 0-251.5 95T632 821"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Intersex(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1024 32q0-14 9-23t23-9h288q26 0 45 19t19 45v288q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23V218l-254 255q126 158 126 359q0 221-147.5 384.5T640 1404v132h96q14 0 23 9t9 23v64q0 14-9 23t-23 9h-96v96q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23v-96h-96q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h96v-132q-149-16-270.5-103T55 1077.5T2 786q16-204 160-353.5T509 260q118-14 228 19t198 103l255-254h-134q-14 0-23-9t-9-23zM576 1280q185 0 316.5-131.5T1024 832T892.5 515.5T576 384T259.5 515.5T128 832t131.5 316.5T576 1280"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ioxhost(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1463 832q0 35-25 60.5t-61 25.5H675q-36 0-61-25.5T589 832t25-60.5t61-25.5h702q36 0 61 25.5t25 60.5m214 0q0-86-23-170H672q-36 0-61-25t-25-60q0-36 25-61t61-25h908q-88-143-235-227t-320-84q-177 0-327.5 87.5T459.5 505T372 832q0 86 23 170h982q36 0 61 25t25 60q0 36-25 61t-61 25H469q88 143 235.5 227t320.5 84q132 0 253-51.5t208-139t139-208t52-253.5m371-255q0 35-25 60t-61 25h-131q17 85 17 170q0 167-65.5 319.5t-175.5 263t-262.5 176T1025 1656q-246 0-448.5-133T275 1173H86q-36 0-61-25t-25-61q0-35 25-60t61-25h132q-17-85-17-170q0-167 65.5-319.5t175.5-263t262.5-176T1025 8q245 0 447.5 133T1774 491h188q36 0 61 25t25 61"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Italic(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 1534l17-85q22-7 61.5-16.5t72-19T210 1390q28-35 41-101q1-7 62-289t114-543.5T479 160v-25q-24-13-54.5-18.5t-69.5-8t-58-5.5L316 0q33 2 120 6.5t149.5 7T706 16q48 0 98.5-2.5t121-7T1024 0q-5 39-19 89q-30 10-101.5 28.5T795 151q-8 19-14 42.5t-9 40t-7.5 45.5t-6.5 42q-27 148-87.5 419.5T593 1096q-2 9-13 58t-20 90t-16 83.5t-6 57.5l1 18q17 4 185 31q-3 44-16 99q-11 0-32.5 1.5T643 1536q-29 0-87-10t-86-10q-138-2-206-2q-51 0-143 9T0 1534"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Joomla(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1070 945l-160 160l-151 152l-30 30q-65 64-151.5 87t-171.5 2q-16 70-72 115t-129 45q-85 0-145-60.5T0 1330q0-72 44.5-128t113.5-72q-22-86 1-173t88-152l12-12l151 152l-11 11q-37 37-37 89t37 90q37 37 89 37t89-37l30-30l151-152l161-160zM729 263l12 12l-152 152l-12-12q-37-37-89-37t-89 37t-37 89.5t37 89.5l29 29l152 152l160 160l-151 152l-161-160l-151-152l-30-30q-68-67-90-159.5t5-179.5q-70-15-115-71T2 206q0-85 60-145.5T207 0q76 0 133.5 49T410 172q84-20 169.5 3.5T729 263m807 1067q0 85-60 145.5t-145 60.5q-74 0-131-47t-71-118q-86 28-179.5 6T788 1287l-11-12l151-152l12 12q37 37 89 37t89-37t37-89t-37-89l-30-30l-152-152l-160-160l152-152l160 160l152 152l29 30q64 64 87.5 150.5t2.5 171.5q76 11 126.5 68.5T1536 1330m-2-1124q0 77-51 135t-127 69q26 85 3 176.5T1269 745l-12 12l-151-152l12-12q37-37 37-89t-37-89t-89-37t-89 37l-30 30l-152 152l-160 160l-152-152l161-160l152-152l29-30q67-67 159-89.5t178 3.5q11-75 68.5-126T1329 0q85 0 145 60.5t60 145.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Jsfiddle(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1800 644q111 46 179.5 145.5T2048 1011q0 164-118 280.5T1645 1408q-4 0-11.5-.5t-10.5-.5H406q-170-10-288-125.5T0 1001q0-110 55-203t147-147q-12-39-12-82q0-115 82-196t199-81q95 0 172 58q75-154 222.5-248T1192 8q166 0 306 80.5T1719.5 307t81.5 301q0 6-.5 18t-.5 18M468 910q0 122 84 193t208 71q137 0 240-99q-16-20-47.5-56.5T909 968q-67 65-144 65q-55 0-93.5-33.5T633 912q0-53 38.5-87t91.5-34q44 0 84.5 21t73 55t65 75t69 82t77 75t97 55t121.5 21q121 0 204.5-71.5T1638 913q0-121-84-192t-207-71q-143 0-241 97l93 108q66-64 142-64q52 0 92 33t40 84q0 57-37 91.5t-94 34.5q-43 0-82.5-21t-72-55t-65.5-75t-69.5-82t-77.5-75t-96.5-55T760 650q-122 0-207 70.5T468 910"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Key(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 384q0-80-56-136t-136-56t-136 56t-56 136q0 42 19 83q-41-19-83-19q-80 0-136 56t-56 136t56 136t136 56t136-56t56-136q0-42-19-83q41 19 83 19q80 0 136-56t56-136m851 704q0 17-49 66t-66 49q-9 0-28.5-16t-36.5-33t-38.5-40t-24.5-26l-96 96l220 220q28 28 28 68q0 42-39 81t-81 39q-40 0-68-28L733 893q-176 131-365 131q-163 0-265.5-102.5T0 656q0-160 95-313T343 95T656 0q163 0 265.5 102.5T1024 368q0 189-131 365l355 355l96-96q-3-3-26-24.5t-40-38.5t-33-36.5t-16-28.5q0-17 49-66t66-49q13 0 23 10q6 6 46 44.5t82 79.5t86.5 86t73 78t28.5 41"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyboardO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 912v96q0 16-16 16h-96q-16 0-16-16v-96q0-16 16-16h96q16 0 16 16m128-256v96q0 16-16 16H272q-16 0-16-16v-96q0-16 16-16h224q16 0 16 16M384 400v96q0 16-16 16h-96q-16 0-16-16v-96q0-16 16-16h96q16 0 16 16m1024 512v96q0 16-16 16H528q-16 0-16-16v-96q0-16 16-16h864q16 0 16 16M768 656v96q0 16-16 16h-96q-16 0-16-16v-96q0-16 16-16h96q16 0 16 16M640 400v96q0 16-16 16h-96q-16 0-16-16v-96q0-16 16-16h96q16 0 16 16m384 256v96q0 16-16 16h-96q-16 0-16-16v-96q0-16 16-16h96q16 0 16 16M896 400v96q0 16-16 16h-96q-16 0-16-16v-96q0-16 16-16h96q16 0 16 16m384 256v96q0 16-16 16h-96q-16 0-16-16v-96q0-16 16-16h96q16 0 16 16m384 256v96q0 16-16 16h-96q-16 0-16-16v-96q0-16 16-16h96q16 0 16 16m-512-512v96q0 16-16 16h-96q-16 0-16-16v-96q0-16 16-16h96q16 0 16 16m256 0v96q0 16-16 16h-96q-16 0-16-16v-96q0-16 16-16h96q16 0 16 16m256 0v352q0 16-16 16h-224q-16 0-16-16v-96q0-16 16-16h112V400q0-16 16-16h96q16 0 16 16m128 752V256H128v896zm128-896v896q0 53-37.5 90.5T1792 1280H128q-53 0-90.5-37.5T0 1152V256q0-53 37.5-90.5T128 128h1664q53 0 90.5 37.5T1920 256"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Krw(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m514 1067l81-299H436l75 300q1 1 1 3t1 3q0-1 .5-3.5t.5-3.5m116-427l35-128H373l32 128zm192 0h139l-35-128h-70zm449 428l78-300h-162l81 299q0 1 .5 3.5t1.5 3.5q0-1 .5-3t.5-3m111-428l33-128h-297l34 128zm410 32v64q0 14-9 23t-23 9h-213l-164 616q-7 24-31 24h-159q-24 0-31-24L996 768H787l-167 616q-7 24-31 24H430q-11 0-19.5-7t-10.5-17L240 768H32q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h175l-33-128H32q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h109L52 40q-5-15 5-28Q67 0 83 0h137q26 0 31 24l90 360h359l97-360q7-24 31-24h126q24 0 31 24l98 360h365l93-360q5-24 31-24h137q16 0 26 12q10 13 5 28l-91 344h111q14 0 23 9t9 23v64q0 14-9 23t-23 9h-145l-34 128h179q14 0 23 9t9 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Language(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M654 1078q-1 3-12.5-.5T610 1066l-20-9q-44-20-87-49q-7-5-41-31.5T424 948q-67 103-134 181q-81 95-105 110q-4 2-19.5 4t-18.5 0q6-4 82-92q21-24 85.5-115T393 918q17-30 51-98.5t36-77.5q-8-1-110 33q-8 2-27.5 7.5T308 792t-17 5q-2 2-2 10.5t-1 9.5q-5 10-31 15q-23 7-47 0q-18-4-28-21q-4-6-5-23q6-2 24.5-5t29.5-6q58-16 105-32q100-35 102-35q10-2 43-19.5t44-21.5q9-3 21.5-8t14.5-5.5t6 .5q2 12-1 33q0 2-12.5 27T527 769.5T510 803q-25 50-77 131l64 28q12 6 74.5 32t67.5 28q4 1 10.5 25.5t4.5 30.5M449 592q3 15-4 28q-12 23-50 38q-30 12-60 12q-26-3-49-26q-14-15-18-41l1-3q3 3 19.5 5t26.5 0t58-16q36-12 55-14q17 0 21 17m698 129l63 227l-139-42zM39 1521l694-232V257L39 490zm1241-317l102 31l-181-657l-100-31l-216 536l102 31l45-110l211 65zM777 242l573 184V46zm311 1323l158 13l-54 160l-40-66q-130 83-276 108q-58 12-91 12h-84q-79 0-199.5-39T318 1668q-8-7-8-16q0-8 5-13.5t13-5.5q4 0 18 7.5t30.5 16.5t20.5 11q73 37 159.5 61.5T714 1754q95 0 167-14.5t157-50.5q15-7 30.5-15.5t34-19t28.5-16.5zm448-1079v1079l-774-246q-14 6-375 127.5T19 1568q-13 0-18-13q0-1-1-3V474q3-9 4-10q5-6 20-11q107-36 149-50V19l558 198q2 0 160.5-55t316-108.5T1369 0q20 0 20 21v418z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laptop(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M416 1024q-66 0-113-47t-47-113V160q0-66 47-113T416 0h1088q66 0 113 47t47 113v704q0 66-47 113t-113 47zm-32-864v704q0 13 9.5 22.5T416 896h1088q13 0 22.5-9.5t9.5-22.5V160q0-13-9.5-22.5T1504 128H416q-13 0-22.5 9.5T384 160m1376 928h160v96q0 40-47 68t-113 28H160q-66 0-113-28t-47-68v-96zm-720 96q16 0 16-16t-16-16H880q-16 0-16 16t16 16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lastfm(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1292 448q0 6 10 41q10 29 25 49.5t41 34t44 20t55 16.5q325 91 325 332q0 146-105.5 242.5T1432 1280q-59 0-111.5-18.5T1229 1216t-77-74.5t-63-87.5t-53.5-103.5t-43.5-103T952.5 741T917 646q-32-81-61.5-133.5T782 416t-104-64t-142-20q-96 0-183 55.5T215 532t-51 185q0 160 106.5 279.5T534 1116q177 0 258-95q56-63 83-116l84 152q-15 34-44 70l1 1q-131 152-388 152q-147 0-269.5-79T68 993.5T0 719q0-105 43.5-206t116-176.5t172-121.5T536 169q87 0 159 19t123.5 50t95 80t72.5 99t58.5 117t50.5 124.5t50 130.5t55 127q96 200 233 200q81 0 138.5-48.5T1629 939q0-42-19-72t-50.5-46t-72.5-31.5t-84.5-27t-87.5-34t-81-52t-65-82t-39-122.5q-3-16-3-33q0-110 87.5-192t198.5-78q78 3 120.5 14.5T1624 237h-1q12 11 23 24.5t26 36t19 27.5l-129 99q-26-49-54-70v-1q-23-21-97-21q-49 0-84 33t-35 83"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LastfmSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1432 924q0-173-234-239q-35-10-53-16.5t-38-25t-29-46.5q0-2-2-8.5t-3-12t-1-7.5q0-36 24.5-59.5T1157 486q54 0 71 15h-1q20 15 39 51l93-71q-39-54-49-64q-33-29-67.5-39t-85.5-10q-80 0-142 57.5T953 563q0 7 2 23q16 96 64.5 140t148.5 73q29 8 49 15.5t45 21.5t38.5 34.5T1314 917v5q1 58-40.5 93t-100.5 35q-97 0-167-144q-23-47-51.5-121.5t-48-125.5t-54-110.5t-74-95.5T675 392.5T528 368q-101 0-192 56T192 572t-50 192v1q4 108 50.5 199T326 1111.5t196 56.5q186 0 279-110q20-27 31-51l-60-109q-42 80-99 116t-146 36q-115 0-191-87t-76-204q0-105 82-189t186-84q112 0 170 53.5T802 712q8 21 25.5 68.5T856 857t31.5 74.5t38.5 74t45.5 62.5t55.5 53.5t66 33t80 13.5q107 0 183-69.5t76-174.5m104-636v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Leaf(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1280 576q0-26-19-45t-45-19q-172 0-318 49.5t-259.5 134T403 915q-19 21-19 45q0 26 19 45t45 19q24 0 45-19q27-24 74-71t67-66q137-124 268.5-176t313.5-52q26 0 45-19t19-45m512-198q0 95-20 193q-46 224-184.5 383T1230 1222q-214 108-438 108q-148 0-286-47q-15-5-88-42t-96-37q-16 0-39.5 32t-45 70t-52.5 70t-60 32q-43 0-63.5-17.5T16 1331q-2-4-6-11t-5.5-10t-3-9.5T0 1287q0-35 31-73.5t68-65.5t68-56t31-48q0-4-14-38t-16-44q-9-51-9-104q0-115 43.5-220t119-184.5t170.5-139T696 219q55-18 145-25.5t179.5-9t178.5-6t163.5-24T1476 98l29.5-29.5l29.5-28l27-20l36.5-16L1642 0q39 0 70.5 46t47.5 112t24 124t8 96"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Leanpub(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1893 136l155 1272q-131 0-257-57q-200-91-393-91q-226 0-374 148q-148-148-374-148q-193 0-393 91q-128 57-252 57H0L155 136Q379 9 637 9q233 0 387 106Q1178 9 1411 9q258 0 482 127m-495 987q129 0 232 28.5t260 93.5L1766 224q-171-78-368-78q-224 0-374 141q-150-141-374-141q-197 0-368 78L158 1245q105-43 165.5-65t148.5-39.5t178-17.5q202 0 374 108q172-108 374-108m40-34l-55-907q-211 4-359 155q-152-155-374-155q-176 0-336 66l-114 941q124-51 228.5-76t221.5-25q209 0 374 102q172-107 374-102z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LemonO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1407 698q0-44-7-113.5t-18-96.5q-12-30-17-44t-9-36.5t-4-48.5q0-23 5-68.5t5-67.5q0-37-10-55q-4-1-13-1q-19 0-58 4.5t-59 4.5q-60 0-176-24t-175-24q-43 0-94.5 11.5t-85 23.5t-89.5 34q-137 54-202 103q-96 73-159.5 189.5t-88 236T128 974q0 40 12.5 120t12.5 121q0 23-11 66.5t-11 65.5t12 36.5t34 14.5q24 0 72.5-11t73.5-11q57 0 169.5 15.5T662 1407q181 0 284-36q129-45 235.5-152.5t166-245.5t59.5-275m128-2q0 165-70 327.5t-196 288T988 1492q-124 44-326 44q-57 0-170-14.5T323 1507q-24 0-72.5 14.5T177 1536q-73 0-123.5-55.5T3 1352q0-24 11-68t11-67q0-40-12.5-120.5T0 975q0-111 18-217.5T72.5 548T173 354t150-156q78-59 232-120Q749 0 871 0q60 0 175.5 24T1220 48q19 0 57-5t58-5q81 0 118 50.5t37 134.5q0 23-5 68t-5 68q0 13 2 25t3.5 16.5t7.5 20.5t8 20q16 40 25 118.5t9 136.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LevelUp(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1018 475q-18 37-58 37H768v864q0 14-9 23t-23 9H32q-21 0-29-18q-8-20 4-35l160-192q9-11 25-11h320V512H320q-40 0-58-37q-17-37 9-68L591 23q18-22 49-22t49 22l320 384q27 32 9 68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LifeBouy(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M896 0q182 0 348 71t286 191t191 286t71 348t-71 348t-191 286t-286 191t-348 71t-348-71t-286-191t-191-286T0 896t71-348t191-286T548 71T896 0m0 128q-190 0-361 90l194 194q82-28 167-28t167 28l194-194q-171-90-361-90M218 1257l194-194q-28-82-28-167t28-167L218 535q-90 171-90 361t90 361m678 407q190 0 361-90l-194-194q-82 28-167 28t-167-28l-194 194q171 90 361 90m0-384q159 0 271.5-112.5T1280 896t-112.5-271.5T896 512T624.5 624.5T512 896t112.5 271.5T896 1280m484-217l194 194q90-171 90-361t-90-361l-194 194q28 82 28 167t-28 167"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightbulbO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M736 448q0 13-9.5 22.5T704 480t-22.5-9.5T672 448q0-46-54-71t-106-25q-13 0-22.5-9.5T480 320t9.5-22.5T512 288q50 0 99.5 16t87 54t37.5 90m160 0q0-72-34.5-134t-90-101.5t-123-62T512 128t-136.5 22.5t-123 62t-90 101.5T128 448q0 101 68 180q10 11 30.5 33t30.5 33q128 153 141 298h228q13-145 141-298q10-11 30.5-33t30.5-33q68-79 68-180m128 0q0 155-103 268q-45 49-74.5 87T787 898.5T753 1006q47 28 47 82q0 37-25 64q25 27 25 64q0 52-45 81q13 23 13 47q0 46-31.5 71t-77.5 25q-20 44-60 70t-87 26t-87-26t-60-70q-46 0-77.5-25t-31.5-71q0-24 13-47q-45-29-45-81q0-37 25-64q-25-27-25-64q0-54 47-82q-4-50-34-107.5T177.5 803T103 716Q0 603 0 448q0-99 44.5-184.5t117-142t164-89T512 0t186.5 32.5t164 89t117 142T1024 448"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineChart(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2048 1408v128H0V0h128v1408zM1920 160v435q0 21-19.5 29.5T1865 617l-121-121l-633 633q-10 10-23 10t-23-10L832 896l-416 416l-192-192l585-585q10-10 23-10t23 10l233 233l464-464l-121-121q-16-16-7.5-35.5T1453 128h435q14 0 23 9t9 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linkedin(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M349 497v991H19V497zm21-306q1 73-50.5 122T184 362h-2q-82 0-132-49T0 191q0-74 51.5-122.5T186 20t133 48.5T370 191m1166 729v568h-329V958q0-105-40.5-164.5T1040 734q-63 0-105.5 34.5T871 854q-11 30-11 81v553H531q2-399 2-647t-1-296l-1-48h329v144h-2q20-32 41-56t56.5-52t87-43.5T1157 474q171 0 275 113.5T1536 920"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkedinSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M237 1286h231V592H237zm246-908q-1-52-36-86t-93-34t-94.5 34t-36.5 86q0 51 35.5 85.5T351 498h1q59 0 95-34.5t36-85.5m585 908h231V888q0-154-73-233t-193-79q-136 0-209 117h2V592H595q3 66 0 694h231V898q0-38 7-56q15-35 45-59.5t74-24.5q116 0 116 157zm468-998v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linode(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m330 1535l202 214l-34-236l-216-213zm226 226l274-218l-11-245l-300 215zm-311-638l227 213l-48-327l-245-204zm250 224l317-214l-14-324l-352 200zm348 11l95 80l-2-239l-103-79q0 1 1 8.5t0 12t-5 7.5l-78 52l85 70q7 6 7 88M138 606l256 200l-68-465L47 168zm1035 663l15-234l-230 164l2 240zM417 814l373-194l-19-441l-423 163zm853 365l20-233l-226-142l-2 105l144 95q6 4 4 9l-7 119zm191-139l30-222l-179 128l-20 228zm-188 167l-71-49l-8 117q0 5-4 8l-234 187q-7 5-14 0l-98-83l7 161q0 5-4 8l-293 234q-4 2-6 2q-8-2-8-3l-228-242q-4-4-59-277q-2-7 5-11l61-37q-94-86-95-92l-72-351q-2-7 6-12l94-45Q119 622 117 614L21 148q-2-10 7-13L461 0q5 0 8 1l317 153q6 4 6 9l20 463q0 7-6 10l-118 61l126 85q5 2 5 8l5 123l121-74q5-4 11 0l84 56l3-110q0-6 5-9l206-126q6-3 11 0l245 135q4 4 5 7t-6.5 60t-17.5 124.5t-10 70.5q0 5-4 7l-191 153q-6 5-13 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linux(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M663 411q-11 1-15.5 10.5T639 431q-5 1-5-5q0-12 19-15zm87 14q-4 1-11.5-6.5T721 414q24-11 32 2q3 6-3 9M399 852q-4-1-6 3t-4.5 12.5T383 881t-10 13q-10 11-1 12q4 1 12.5-7t12.5-18q1-3 2-7t2-6t1.5-4.5t.5-4v-3l-1-2.5zm855 359q0-18-55-42q4-15 7.5-27.5t5-26t3-21.5t.5-22.5t-1-19.5t-3.5-22t-4-20.5t-5-25t-5.5-26.5q-10-48-47-103t-72-75q24 20 57 83q87 162 54 278q-11 40-50 42q-31 4-38.5-18.5t-8-83.5t-11.5-107q-9-39-19.5-69t-19.5-45.5t-15.5-24.5t-13-15t-7.5-7q-14-62-31-103t-29.5-56t-23.5-33t-15-40q-4-21 6-53.5t4.5-49.5t-44.5-25q-15-3-44.5-18T792 419q-8-1-11-26t8-51t36-27q37-3 51 30t4 58q-11 19-2 26.5t30 .5q13-4 13-36v-37q-5-30-13.5-50t-21-30.5t-23.5-15t-27-7.5q-107 8-89 134q0 15-1 15q-9-9-29.5-10.5t-33 .5t-15.5-5q1-57-16-90t-45-34q-27-1-41.5 27.5T549 351q-1 15 3.5 37t13 37.5T581 439q10-3 16-14q4-9-7-8q-7 0-15.5-14.5T565 369q-1-22 9-37t34-14q17 0 27 21t9.5 39t-1.5 22q-22 15-31 29q-8 12-27.5 23.5T564 465q-13 14-15.5 27t7.5 18q14 8 25 19.5t16 19t18.5 13T651 568q47 2 102-15q2-1 23-7t34.5-10.5t29.5-13t21-17.5q9-14 20-8q5 3 6.5 8.5t-3 12T868 527q-20 6-56.5 21.5T766 568q-44 19-70 23q-25 5-79-2q-10-2-9 2t17 19q25 23 67 22q17-1 36-7t36-14t33.5-17.5t30-17t24.5-12t17.5-2.5t8.5 11q0 2-1 4.5t-4 5t-6 4.5t-8.5 5t-9 4.5t-10 5t-9.5 4.5q-28 14-67.5 44T696 693t-49 1q-21-11-63-73q-22-31-25-22q-1 3-1 10q0 25-15 56.5T513.5 721t-21 58t11.5 63q-23 6-62.5 90T394 1073q-2 18-1.5 69t-5.5 59q-8 24-29 3q-32-31-36-94q-2-28 4-56q4-19-1-18q-2 1-4 5q-36 65 10 166q5 12 25 28t24 20q20 23 104 90.5t93 76.5q16 15 17.5 38t-14 43t-45.5 23q8 15 29 44.5t28 54t7 70.5q46-24 7-92q-4-8-10.5-16t-9.5-12t-2-6q3-5 13-9.5t20 2.5q46 52 166 36q133-15 177-87q23-38 34-30q12 6 10 52q-1 25-23 92q-9 23-6 37.5t24 15.5q3-19 14.5-77t13.5-90q2-21-6.5-73.5t-7.5-97t23-70.5q15-18 51-18q1-37 34.5-53t72.5-10.5t60 22.5M626 384q3-17-2.5-30T612 339q-9-2-9 7q2 5 5 6q10 0 7 15q-3 20 8 20q3 0 3-3m419 197q-2-8-6.5-11.5t-13-5t-14.5-5.5q-5-3-9.5-8t-7-8t-5.5-6.5t-4-4t-4 1.5q-14 16 7 43.5t39 31.5q9 1 14.5-8t3.5-20M867 368q0-11-5-19.5T851 336t-9-3q-6 0-8 2t0 4t5 3q14 4 18 31q0 3 8-2q2-2 2-3m54-233q0-2-2.5-5t-9-7t-9.5-6q-15-15-24-15q-9 1-11.5 7.5t-1 13t-.5 12.5q-1 4-6 10.5t-6 9t3 8.5q4 3 8 0t11-9t15-9q1-1 9-1t15-2t9-7m565 1341q20 12 31 24.5t12 24t-2.5 22.5t-15.5 22t-23.5 19.5t-30 18.5t-31.5 16.5t-32 15.5t-27 13q-38 19-85.5 56t-75.5 64q-17 16-68 19.5t-89-14.5q-18-9-29.5-23.5T1003 1728t-22-19.5t-47-9.5q-44-1-130-1q-19 0-57 1.5t-58 2.5q-44 1-79.5 15t-53.5 30t-43.5 28.5T459 1787q-29-1-111-31t-146-43q-19-4-51-9.5t-50-9t-39.5-9.5t-33.5-14.5t-17-19.5q-10-23 7-66.5t18-54.5q1-16-4-40t-10-42.5t-4.5-36.5t10.5-27q14-12 57-14t60-12q30-18 42-35t12-51q21 73-32 106q-32 20-83 15q-34-3-43 10q-13 15 5 57q2 6 8 18t8.5 18t4.5 17t1 22q0 15-17 49t-14 48q3 17 37 26q20 6 84.5 18.5T258 1697q24 6 74 22t82.5 23t55.5 4q43-6 64.5-28t23-48t-7.5-58.5t-19-52t-20-36.5q-121-190-169-242q-68-74-113-40q-11 9-15-15q-3-16-2-38q1-29 10-52t24-47t22-42q8-21 26.5-72t29.5-78t30-61t39-54q110-143 124-195q-12-112-16-310q-2-90 24-151.5T631 21Q670 0 735 0q53-1 106 13.5T930 55q57 42 91.5 121.5T1051 324q-5 95 30 214q34 113 133 218q55 59 99.5 163t59.5 191q8 49 5 84.5t-12 55.5t-20 22q-10 2-23.5 19t-27 35.5t-40.5 33.5t-61 14q-18-1-31.5-5t-22.5-13.5t-13.5-15.5t-11.5-20.5t-9-19.5q-22-37-41-30t-28 49t7 97q20 70 1 195q-10 65 18 100.5t73 33t85-35.5q59-49 89.5-66.5T1414 1600q53-18 77-36.5t18.5-34.5t-25-28.5t-51.5-23.5q-33-11-49.5-48t-15-72.5t15.5-47.5q1 31 8 56.5t14.5 40.5t20.5 28.5t21 19t21.5 13t16.5 9.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func List(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 1184v192q0 13-9.5 22.5T224 1408H32q-13 0-22.5-9.5T0 1376v-192q0-13 9.5-22.5T32 1152h192q13 0 22.5 9.5t9.5 22.5m0-384v192q0 13-9.5 22.5T224 1024H32q-13 0-22.5-9.5T0 992V800q0-13 9.5-22.5T32 768h192q13 0 22.5 9.5T256 800m0-384v192q0 13-9.5 22.5T224 640H32q-13 0-22.5-9.5T0 608V416q0-13 9.5-22.5T32 384h192q13 0 22.5 9.5T256 416m1536 768v192q0 13-9.5 22.5t-22.5 9.5H416q-13 0-22.5-9.5T384 1376v-192q0-13 9.5-22.5t22.5-9.5h1344q13 0 22.5 9.5t9.5 22.5M256 32v192q0 13-9.5 22.5T224 256H32q-13 0-22.5-9.5T0 224V32Q0 19 9.5 9.5T32 0h192q13 0 22.5 9.5T256 32m1536 768v192q0 13-9.5 22.5t-22.5 9.5H416q-13 0-22.5-9.5T384 992V800q0-13 9.5-22.5T416 768h1344q13 0 22.5 9.5t9.5 22.5m0-384v192q0 13-9.5 22.5T1760 640H416q-13 0-22.5-9.5T384 608V416q0-13 9.5-22.5T416 384h1344q13 0 22.5 9.5t9.5 22.5m0-384v192q0 13-9.5 22.5T1760 256H416q-13 0-22.5-9.5T384 224V32q0-13 9.5-22.5T416 0h1344q13 0 22.5 9.5T1792 32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListAlt(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 1056v64q0 13-9.5 22.5T352 1152h-64q-13 0-22.5-9.5T256 1120v-64q0-13 9.5-22.5t22.5-9.5h64q13 0 22.5 9.5t9.5 22.5m0-256v64q0 13-9.5 22.5T352 896h-64q-13 0-22.5-9.5T256 864v-64q0-13 9.5-22.5T288 768h64q13 0 22.5 9.5T384 800m0-256v64q0 13-9.5 22.5T352 640h-64q-13 0-22.5-9.5T256 608v-64q0-13 9.5-22.5T288 512h64q13 0 22.5 9.5T384 544m1152 512v64q0 13-9.5 22.5t-22.5 9.5H544q-13 0-22.5-9.5T512 1120v-64q0-13 9.5-22.5t22.5-9.5h960q13 0 22.5 9.5t9.5 22.5m0-256v64q0 13-9.5 22.5T1504 896H544q-13 0-22.5-9.5T512 864v-64q0-13 9.5-22.5T544 768h960q13 0 22.5 9.5t9.5 22.5m0-256v64q0 13-9.5 22.5T1504 640H544q-13 0-22.5-9.5T512 608v-64q0-13 9.5-22.5T544 512h960q13 0 22.5 9.5t9.5 22.5m128 704V416q0-13-9.5-22.5T1632 384H160q-13 0-22.5 9.5T128 416v832q0 13 9.5 22.5t22.5 9.5h1472q13 0 22.5-9.5t9.5-22.5m128-1088v1088q0 66-47 113t-113 47H160q-66 0-113-47T0 1248V160Q0 94 47 47T160 0h1472q66 0 113 47t47 113"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListOl(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M381 1620q0 80-54.5 126T191 1792q-106 0-172-66l57-88q49 45 106 45q29 0 50.5-14.5T254 1626q0-64-105-56l-26-56q8-10 32.5-43.5t42.5-54t37-38.5v-1q-16 0-48.5 1t-48.5 1v53H32v-152h333v88l-95 115q51 12 81 49t30 88m2-627v159H21q-6-36-6-54q0-51 23.5-93T95 937t66-47.5t56.5-43.5t23.5-45q0-25-14.5-38.5T187 749q-46 0-81 58l-85-59q24-51 71.5-79.5T198 640q73 0 123 41.5T371 794q0 50-34 91.5T262 950t-75.5 50.5T151 1053h127v-60zm1409 319v192q0 13-9.5 22.5t-22.5 9.5H544q-13 0-22.5-9.5T512 1504v-192q0-14 9-23t23-9h1216q13 0 22.5 9.5t9.5 22.5M384 413v99H49v-99h107q0-41 .5-121.5T157 170v-12h-2q-8 17-50 54l-71-76L170 9h106v404zm1408 387v192q0 13-9.5 22.5t-22.5 9.5H544q-13 0-22.5-9.5T512 992V800q0-14 9-23t23-9h1216q13 0 22.5 9.5t9.5 22.5m0-512v192q0 13-9.5 22.5T1760 512H544q-13 0-22.5-9.5T512 480V288q0-13 9.5-22.5T544 256h1216q13 0 22.5 9.5t9.5 22.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListUl(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 1216q0 80-56 136t-136 56t-136-56t-56-136t56-136t136-56t136 56t56 136m0-512q0 80-56 136t-136 56t-136-56T0 704t56-136t136-56t136 56t56 136m1408 416v192q0 13-9.5 22.5t-22.5 9.5H544q-13 0-22.5-9.5T512 1312v-192q0-13 9.5-22.5t22.5-9.5h1216q13 0 22.5 9.5t9.5 22.5M384 192q0 80-56 136t-136 56t-136-56T0 192T56 56T192 0t136 56t56 136m1408 416v192q0 13-9.5 22.5T1760 832H544q-13 0-22.5-9.5T512 800V608q0-13 9.5-22.5T544 576h1216q13 0 22.5 9.5t9.5 22.5m0-512v192q0 13-9.5 22.5T1760 320H544q-13 0-22.5-9.5T512 288V96q0-13 9.5-22.5T544 64h1216q13 0 22.5 9.5T1792 96"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationArrow(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1401 93L761 1373q-17 35-57 35q-5 0-15-2q-22-5-35.5-22.5T640 1344V768H64q-22 0-39.5-13.5T2 719t4-42t29-30L1315 7q13-7 29-7q27 0 45 19q15 14 18.5 34.5T1401 93"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M320 640h512V448q0-106-75-181t-181-75t-181 75t-75 181zm832 96v576q0 40-28 68t-68 28H96q-40 0-68-28t-28-68V736q0-40 28-68t68-28h32V448q0-184 132-316T576 0t316 132t132 316v192h32q40 0 68 28t28 68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LongArrowLeft(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1728 544v192q0 14-9 23t-23 9H448v224q0 21-19 29t-35-5L10 666Q0 656 0 643q0-14 10-24l384-354q16-14 35-6q19 9 19 29v224h1248q14 0 23 9t9 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LongArrowUp(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M765 429q-9 19-29 19H512v1248q0 14-9 23t-23 9H288q-14 0-23-9t-9-23V448H32q-21 0-29-19t5-35L358 10q10-10 23-10q14 0 24 10l355 384q13 16 5 35"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LowVision(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M335 1228q-2 0-6-2q-86-57-168.5-145T21 901Q0 871 0 832q0-9 2-19t4-18t7-18t8.5-16T32 744t10-15t12-15.5T65 699q184-251 452-365q-110-198-110-211q0-19 17-29q116-64 128-64q18 0 28 16l124 229q92-19 192-19q266 0 497.5 137.5T1772 763q20 31 20 69t-20 69q-91 142-218.5 253.5T1275 1330q110 198 110 211q0 20-17 29q-116 64-127 64q-19 0-29-16l-124-229l-64-119l-444-820l7-7q-58 24-99 47q3 5 127 234t243 449t119 223q0 7-9 9q-13 3-72 3q-57 0-60-7L380 560q-39 28-82 68q24 43 214 393.5T702 1376q0 10-11 10q-14 0-82.5-22t-72.5-28l-106-197l-224-413q-44 53-78 106q2 3 18 25t23 34l176 327q0 10-10 10m830-102l49 91q273-111 450-385q-180-277-459-389q67 64 103 148.5t36 176.5q0 106-47 200.5T1165 1126M848 512q0 20 14 34t34 14q86 0 147 61t61 147q0 20 14 34t34 14t34-14t14-34q0-126-89-215t-215-89q-20 0-34 14t-14 34m366-65l-9-4l7 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Magic(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1190 581l293-293l-107-107l-293 293zm447-293q0 27-18 45L333 1619q-18 18-45 18t-45-18L45 1421q-18-18-18-45t18-45L1331 45q18-18 45-18t45 18l198 198q18 18 18 45M286 98l98 30l-98 30l-30 98l-30-98l-98-30l98-30l30-98zm350 162l196 60l-196 60l-60 196l-60-196l-196-60l196-60l60-196zm930 478l98 30l-98 30l-30 98l-30-98l-98-30l98-30l30-98zM926 98l98 30l-98 30l-30 98l-30-98l-98-30l98-30l30-98z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Magnet(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1536 704v128q0 201-98.5 362t-274 251.5T768 1536t-395.5-90.5t-274-251.5T0 832V704q0-26 19-45t45-19h384q26 0 45 19t19 45v128q0 52 23.5 90t53.5 57t71 30t64 13t44 2t44-2t64-13t71-30t53.5-57t23.5-90V704q0-26 19-45t45-19h384q26 0 45 19t19 45M512 64v384q0 26-19 45t-45 19H64q-26 0-45-19T0 448V64q0-26 19-45T64 0h384q26 0 45 19t19 45m1024 0v384q0 26-19 45t-45 19h-384q-26 0-45-19t-19-45V64q0-26 19-45t45-19h384q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailForward(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 576q0 26-19 45l-512 512q-19 19-45 19t-45-19t-19-45V832H928q-98 0-175.5 6t-154 21.5t-133 42.5T360 971.5t-80 101t-48.5 138.5t-17.5 181q0 55 5 123q0 6 2.5 23.5t2.5 26.5q0 15-8.5 25t-23.5 10q-16 0-28-17q-7-9-13-22t-13.5-30t-10.5-24Q0 1222 0 1056q0-199 53-333q162-403 875-403h224V64q0-26 19-45t45-19t45 19l512 512q19 19 19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailReply(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 1056q0 166-127 451q-3 7-10.5 24t-13.5 30t-13 22q-12 17-28 17q-15 0-23.5-10t-8.5-25q0-9 2.5-26.5t2.5-23.5q5-68 5-123q0-101-17.5-181t-48.5-138.5t-80-101t-105.5-69.5t-133-42.5t-154-21.5t-175.5-6H640v256q0 26-19 45t-45 19t-45-19L19 621Q0 602 0 576t19-45L531 19q19-19 45-19t45 19t19 45v256h224q713 0 875 403q53 134 53 333"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailReplyAll(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 1018v70q0 42-39 59q-13 5-25 5q-27 0-45-19L19 621Q0 602 0 576t19-45L531 19q29-31 70-14q39 17 39 59v69L243 531q-19 19-19 45t19 45zm1152 38q0 58-17 133.5t-38.5 138t-48 125t-40.5 90.5l-20 40q-8 17-28 17q-6 0-9-1q-25-8-23-34q43-400-106-565q-64-71-170.5-110.5T1024 837v251q0 42-39 59q-13 5-25 5q-27 0-45-19L403 621q-19-19-19-45t19-45L915 19q29-31 70-14q39 17 39 59v262q411 28 599 221q169 173 169 509"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Male(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1024 672v416q0 40-28 68t-68 28t-68-28t-28-68V736h-64v912q0 46-33 79t-79 33t-79-33t-33-79v-464h-64v464q0 46-33 79t-79 33t-79-33t-33-79V736h-64v352q0 40-28 68t-68 28t-68-28t-28-68V672q0-80 56-136t136-56h640q80 0 136 56t56 136M736 224q0 93-65.5 158.5T512 448t-158.5-65.5T288 224t65.5-158.5T512 0t158.5 65.5T736 224"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Map(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 0q13 0 22.5 9.5T544 32v1472q0 20-17 28L47 1788q-7 4-15 4q-13 0-22.5-9.5T0 1760V288q0-20 17-28L497 4q7-4 15-4m1248 0q13 0 22.5 9.5T1792 32v1472q0 20-17 28l-480 256q-7 4-15 4q-13 0-22.5-9.5t-9.5-22.5V288q0-20 17-28L1745 4q7-4 15-4M640 0q8 0 14 3l512 256q18 10 18 29v1472q0 13-9.5 22.5t-22.5 9.5q-8 0-14-3l-512-256q-18-10-18-29V32q0-13 9.5-22.5T640 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarker(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 512q0-106-75-181t-181-75t-181 75t-75 181t75 181t181 75t181-75t75-181m256 0q0 109-33 179l-364 774q-16 33-47.5 52t-67.5 19t-67.5-19t-46.5-52L33 691Q0 621 0 512q0-212 150-362T512 0t362 150t150 362"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2020 11q28 20 28 53v1408q0 20-11 36t-29 23l-640 256q-24 11-48 0l-616-246l-616 246q-10 5-24 5q-19 0-36-11q-28-20-28-53V320q0-20 11-36t29-23L680 5q24-11 48 0l616 246L1960 5q32-13 60 6M736 146v1270l576 230V376zM128 363v1270l544-217V146zm1792 1066V159l-544 217v1270z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapPin(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 1088q66 0 128-15v655q0 26-19 45t-45 19H448q-26 0-45-19t-19-45v-655q62 15 128 15M512 0q212 0 362 150t150 362t-150 362t-362 150t-362-150T0 512t150-362T512 0m0 224q14 0 23-9t9-23t-9-23t-23-9q-146 0-249 103T160 512q0 14 9 23t23 9t23-9t9-23q0-119 84.5-203.5T512 224"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapSigns(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1713 297q10 10 10 23t-10 23l-141 141q-28 28-68 28H160q-26 0-45-19t-19-45V192q0-26 19-45t45-19h576V64q0-26 19-45t45-19h128q26 0 45 19t19 45v64h512q40 0 68 28zm-977 919h256v512q0 26-19 45t-45 19H800q-26 0-45-19t-19-45zm832-448q26 0 45 19t19 45v256q0 26-19 45t-45 19H224q-40 0-68-28L15 983Q5 973 5 960t10-23l141-141q28-28 68-28h512V576h256v192z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mars(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1472 0q26 0 45 19t19 45v416q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23V218l-382 383q126 156 126 359q0 117-45.5 223.5t-123 184t-184 123T576 1536t-223.5-45.5t-184-123t-123-184T0 960t45.5-223.5t123-184t184-123T576 384q203 0 359 126l382-382h-261q-14 0-23-9t-9-23V32q0-14 9-23t23-9zM576 1408q185 0 316.5-131.5T1024 960T892.5 643.5T576 512T259.5 643.5T128 960t131.5 316.5T576 1408"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarsDouble(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1536 416q0-14 9-23t23-9h288q26 0 45 19t19 45v288q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23V602l-254 255q76 95 107.5 214t9.5 247q-31 182-166 312t-318 156q-210 29-384.5-80T545 1406q-117-6-221-57.5t-177.5-133T33 1023T1 793q9-135 78-252t182-191.5T509 260q118-14 227.5 19T935 382l255-254h-134q-14 0-23-9t-9-23V32q0-14 9-23t23-9h288q26 0 45 19t19 45v288q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23V218l-254 255q59 74 93 169q182 9 328 124l255-254h-134q-14 0-23-9t-9-23zm-512 416q0-20-4-58q-162 25-271 150t-109 292q0 20 4 58q162-25 271-150t109-292m-896 0q0 168 111 294t276 149q-3-29-3-59q0-210 135-369.5T985 650q-53-120-163.5-193T576 384q-185 0-316.5 131.5T128 832m960 832q185 0 316.5-131.5T1536 1216q0-168-111-294t-276-149q3 28 3 59q0 210-135 369.5T679 1398q53 120 163.5 193t245.5 73"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarsStroke(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1472 0q26 0 45 19t19 45v416q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23V218l-213 214l140 140q9 10 9 23t-9 22l-46 46q-9 9-22 9t-23-9l-140-141l-78 79q126 156 126 359q0 117-45.5 223.5t-123 184t-184 123T576 1536t-223.5-45.5t-184-123t-123-184T0 960t45.5-223.5t123-184t184-123T576 384q203 0 359 126l78-78l-172-172q-9-10-9-23t9-22l46-46q9-9 22-9t23 9l172 172l213-213h-261q-14 0-23-9t-9-23V32q0-14 9-23t23-9zM576 1408q185 0 316.5-131.5T1024 960T892.5 643.5T576 512T259.5 643.5T128 960t131.5 316.5T576 1408"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarsStrokeH(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1901 659q19 19 19 45t-19 45l-294 294q-9 10-22.5 10t-22.5-10l-45-45q-10-9-10-22.5t10-22.5l185-185h-294v224q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23V768h-132q-24 217-187.5 364.5T576 1280q-167 0-306-87T58 957T4 638q15-133 88-245.5t188-182T529 130q155-12 292 52.5t224 186T1148 640h132V416q0-14 9-23t23-9h64q14 0 23 9t9 23v224h294l-185-185q-10-9-10-22.5t10-22.5l45-45q9-10 22.5-10t22.5 10zM576 1152q185 0 316.5-131.5T1024 704T892.5 387.5T576 256T259.5 387.5T128 704t131.5 316.5T576 1152"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarsStrokeV(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 644q217 24 364.5 187.5T1152 1216q0 167-87 306t-236 212t-319 54q-133-15-245.5-88t-182-188T2 1263q-12-155 52.5-292t186-224T512 644V512H352q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h160V219l-92 92q-10 9-23 9t-22-9l-46-46q-9-9-9-22t9-23L531 19q19-19 45-19t45 19l202 201q9 10 9 23t-9 22l-46 46q-9 9-22 9t-23-9l-92-92v165h160q14 0 23 9t9 23v64q0 14-9 23t-23 9H640zm-64 1020q185 0 316.5-131.5T1024 1216T892.5 899.5T576 768T259.5 899.5T128 1216t131.5 316.5T576 1664"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Maxcdn(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1745 517l-164 763h-334l178-832q13-56-15-88q-27-33-83-33h-169l-204 953H620l204-953H538l-204 953H0l204-953L51 0h1276q101 0 189.5 40.5T1664 154q60 73 81 168.5t0 194.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Meanpath(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1311 714v114q0 24-13.5 38t-37.5 14h-202q-24 0-38-14t-14-38V714q0-24 14-38t38-14h202q24 0 37.5 14t13.5 38M821 944V694q0-53-32.5-85.5T703 576H570q-68 0-96 52q-28-52-96-52H248q-53 0-85.5 32.5T130 694v250q0 22 21 22h55q22 0 22-22V714q0-24 13.5-38t38.5-14h94q24 0 38 14t14 38v230q0 22 21 22h54q22 0 22-22V714q0-24 14-38t38-14h97q24 0 37.5 14t13.5 38v230q0 22 22 22h55q21 0 21-22m589-96V694q0-53-33-85.5t-86-32.5h-264q-53 0-86 32.5T908 694v410q0 21 22 21h55q21 0 21-21V924q31 42 94 42h191q53 0 86-32.5t33-85.5m126-616v1072q0 96-68 164t-164 68H232q-96 0-164-68T0 1304V232q0-96 68-164T232 0h1072q96 0 164 68t68 164"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Medium(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M597 293v1173q0 25-12.5 42.5T548 1526q-17 0-33-8L50 1285q-21-10-35.5-33.5T0 1205V65q0-20 10-34t29-14q14 0 44 15l511 256q3 3 3 5m64 101l534 866l-534-266zm1131 18v1054q0 25-14 40.5t-38 15.5t-47-13l-441-220zm-3-120q0 3-256.5 419.5T1232 1199L842 565l324-527q17-28 52-28q14 0 26 6l541 270q4 2 4 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Medkit(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1280 992V800q0-14-9-23t-23-9h-224V544q0-14-9-23t-23-9H800q-14 0-23 9t-9 23v224H544q-14 0-23 9t-9 23v192q0 14 9 23t23 9h224v224q0 14 9 23t23 9h192q14 0 23-9t9-23v-224h224q14 0 23-9t9-23M640 256h512V128H640zm-384 0v1280h-32q-92 0-158-66T0 1312V480q0-92 66-158t158-66zm1184 0v1280H352V256h160V96q0-40 28-68t68-28h576q40 0 68 28t28 68v160zm352 224v832q0 92-66 158t-158 66h-32V256h32q92 0 158 66t66 158"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Meetup(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1302 1234q-4-24-27.5-34t-49.5-10.5t-48.5-12.5t-25.5-38q-5-47 33-139.5t75-181t32-127.5q-14-101-117-103q-45-1-75 16l-3 2l-5 2.5l-4.5 2l-5 2l-5 .5l-6-1.5l-6-3.5l-6.5-5q-3-2-9-8.5t-9-9t-8.5-7.5t-9.5-7.5t-9.5-5.5t-11-4.5T990 558q-30-5-48 3t-45 31q-1 1-9 8.5t-12.5 11t-15 10T844 627t-17-3q-54-27-84-40q-41-18-94 5t-76 65q-16 28-41 98.5T488.5 885t-40 134t-21.5 73q-22 69 18.5 119t110.5 46q30-2 50.5-15t38.5-46q7-13 79-199.5T801 802q6-11 21.5-18t29.5 0q27 15 21 53q-2 18-51 139.5T772 1109q-6 38 19.5 56.5t60.5 7t55-49.5q4-8 45.5-92t81.5-163.5t46-88.5q20-29 41-28q29 0 25 38q-2 16-65.5 147.5T1010 1096q-12 53 13 103t74 74q17 9 51 15.5t71.5 8t62.5-14t20-48.5m-951 216q3 15-5 27.5t-23 15.5q-14 3-26.5-5t-15.5-23q-3-14 5-27t22-16t27 5t16 23m570 263q12 17 8.5 37.5T909 1783t-37.5 8t-32.5-21q-11-17-7.5-37.5T852 1700t37.5-8t31.5 21M145 901q-18 27-49.5 33T38 921q-26-18-32-50t12-58q18-27 49.5-33t57.5 12q26 19 32 50.5T145 901m1290 677q19 28 13 61.5t-34 52.5t-60.5 13t-51.5-34t-13-61t33-53q28-19 60.5-13t52.5 34m112-604q69 113 42.5 244.5T1455 1426q-90 63-199 60q-20 80-84.5 127t-143.5 44.5t-140-57.5q-12 9-13 10q-103 71-225 48.5T457 1532q-50-73-53-164q-83-14-142.5-70.5t-80.5-128t-2-152T260 879q-36-60-38-128t24.5-125t79.5-98.5T447 477q32-85 99-148t146.5-91.5t168-17T1020 287q72-21 140-17.5t128.5 36t104.5 80t67.5 115T1478 641q52 16 87 57t45.5 89t-5.5 99.5t-58 87.5M423 314q14 20 9.5 44.5T408 397q-19 14-43.5 9.5T327 382q-14-20-9.5-44.5T342 299q19-14 43.5-9.5T423 314M582 33q4 16-5 30.5T551 82t-31-5.5T502 50q-3-17 6.5-31T534 1q17-4 31 5.5T582 33m1186 948q4 20-6.5 37t-30.5 21q-19 4-36-6.5t-21-30.5t6.5-37t30.5-22q20-4 36.5 7.5T1768 981M1104 88q16 27 8.5 58.5T1077 194q-27 16-57.5 8.5T973 168q-16-28-8.5-59T999 61t58-9t47 36m746 656q4 15-4 27.5t-23 16.5q-15 3-27.5-5.5T1780 760q-3-15 5-28t23-16q14-3 26.5 5t15.5 23m-191-241q15 22 10.5 49t-26.5 43q-22 15-49 10t-42-27t-10-49t27-43t48.5-11t41.5 28"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MehO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1152 960q0 26-19 45t-45 19H448q-26 0-45-19t-19-45t19-45t45-19h640q26 0 45 19t19 45M640 512q0 53-37.5 90.5T512 640t-90.5-37.5T384 512t37.5-90.5T512 384t90.5 37.5T640 512m512 0q0 53-37.5 90.5T1024 640t-90.5-37.5T896 512t37.5-90.5T1024 384t90.5 37.5T1152 512m256 256q0-130-51-248.5t-136.5-204t-204-136.5T768 128t-248.5 51t-204 136.5t-136.5 204T128 768t51 248.5t136.5 204t204 136.5t248.5 51t248.5-51t204-136.5t136.5-204t51-248.5m128 0q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mercury(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M830 316q145 72 233.5 210.5T1152 832q0 221-147.5 384.5T640 1404v132h96q14 0 23 9t9 23v64q0 14-9 23t-23 9h-96v96q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23v-96h-96q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h96v-132q-217-24-364.5-187.5T0 832q0-167 88.5-305.5T322 316Q157 220 94 43q-6-16 3.5-29.5T124 0h69q21 0 29 20q44 106 140 171t214 65t214-65T930 20q8-20 37-20h61q17 0 26.5 13.5T1058 43q-63 177-228 273m-254 964q185 0 316.5-131.5T1024 832T892.5 515.5T576 384T259.5 515.5T128 832t131.5 316.5T576 1280"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microchip(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 1280v128H80q-16 0-16-16v-16H16q-16 0-16-16v-32q0-16 16-16h48v-16q0-16 16-16zm0-256v128H80q-16 0-16-16v-16H16q-16 0-16-16v-32q0-16 16-16h48v-16q0-16 16-16zm0-256v128H80q-16 0-16-16v-16H16q-16 0-16-16v-32q0-16 16-16h48v-16q0-16 16-16zm0-256v128H80q-16 0-16-16v-16H16q-16 0-16-16v-32q0-16 16-16h48v-16q0-16 16-16zm0-256v128H80q-16 0-16-16v-16H16q-16 0-16-16v-32q0-16 16-16h48v-16q0-16 16-16zM1280 96v1472q0 40-28 68t-68 28H352q-40 0-68-28t-28-68V96q0-40 28-68t68-28h832q40 0 68 28t28 68m256 1232v32q0 16-16 16h-48v16q0 16-16 16h-112v-128h112q16 0 16 16v16h48q16 0 16 16m0-256v32q0 16-16 16h-48v16q0 16-16 16h-112v-128h112q16 0 16 16v16h48q16 0 16 16m0-256v32q0 16-16 16h-48v16q0 16-16 16h-112V768h112q16 0 16 16v16h48q16 0 16 16m0-256v32q0 16-16 16h-48v16q0 16-16 16h-112V512h112q16 0 16 16v16h48q16 0 16 16m0-256v32q0 16-16 16h-48v16q0 16-16 16h-112V256h112q16 0 16 16v16h48q16 0 16 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microphone(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1152 704v128q0 221-147.5 384.5T640 1404v132h256q26 0 45 19t19 45t-19 45t-45 19H256q-26 0-45-19t-19-45t19-45t45-19h256v-132q-217-24-364.5-187.5T0 832V704q0-26 19-45t45-19t45 19t19 45v128q0 185 131.5 316.5T576 1280t316.5-131.5T1024 832V704q0-26 19-45t45-19t45 19t19 45M896 320v512q0 132-94 226t-226 94t-226-94t-94-226V320q0-132 94-226T576 0t226 94t94 226"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneSlash(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m271 945l-101 101q-42-103-42-214V704q0-26 19-45t45-19t45 19t19 45v128q0 53 15 113m1114-602l-361 361v128q0 132-94 226t-226 94q-55 0-109-19l-96 96q97 51 205 51q185 0 316.5-131.5T1152 832V704q0-26 19-45t45-19t45 19t19 45v128q0 221-147.5 384.5T768 1404v132h256q26 0 45 19t19 45t-19 45t-45 19H384q-26 0-45-19t-19-45t19-45t45-19h256v-132q-125-13-235-81l-254 254q-10 10-23 10t-23-10l-82-82q-10-10-10-23t10-23L1257 215q10-10 23-10t23 10l82 82q10 10 10 23t-10 23m-380-132L384 832V320q0-132 94-226T704 0q102 0 184.5 59T1005 211"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1408 480v192q0 40-28 68t-68 28H96q-40 0-68-28T0 672V480q0-40 28-68t68-28h1216q40 0 68 28t28 68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircle(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1216 832V704q0-26-19-45t-45-19H384q-26 0-45 19t-19 45v128q0 26 19 45t45 19h768q26 0 45-19t19-45m320-64q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1280 832V704q0-26-19-45t-45-19H320q-26 0-45 19t-19 45v128q0 26 19 45t45 19h896q26 0 45-19t19-45m256-544v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSquareO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1152 672v64q0 14-9 23t-23 9H288q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h832q14 0 23 9t9 23m128 448V288q0-66-47-113t-113-47H288q-66 0-113 47t-47 113v832q0 66 47 113t113 47h832q66 0 113-47t47-113m128-832v832q0 119-84.5 203.5T1120 1408H288q-119 0-203.5-84.5T0 1120V288Q0 169 84.5 84.5T288 0h832q119 0 203.5 84.5T1408 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mixcloud(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1645 906q0-59-34-106.5t-87-68.5q-7 45-23 92q-7 24-27.5 38t-44.5 14q-12 0-24-3q-31-10-45-38.5t-4-58.5q23-71 23-143q0-123-61-227.5T1152 239t-228-61q-134 0-247 73T510 445q108 28 188 106q22 23 22 55t-22 54t-54 22t-55-22q-75-75-180-75q-106 0-181 74.5T153 840t75 180.5t181 74.5h1046q79 0 134.5-55.5T1645 906m153 0q0 142-100.5 242T1455 1248H409q-169 0-289-119.5T0 840q0-153 100-267t249-136q62-184 221-298T924 25q235 0 408.5 158.5T1529 573q116 25 192.5 118.5T1798 906m250 0q0 175-97 319q-23 33-64 33q-24 0-43-13q-26-17-32-48.5t12-57.5q71-104 71-233t-71-233q-18-26-12-57t32-49t57.5-11.5T1951 588q97 142 97 318m256 0q0 244-134 443q-23 34-64 34q-23 0-42-13q-26-18-32.5-49t11.5-57q108-164 108-358q0-195-108-357q-18-26-11.5-57.5T2064 443q26-18 57-12t49 33q134 198 134 442"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mobile(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 1152q0-33-23.5-56.5T384 1072t-56.5 23.5T304 1152t23.5 56.5T384 1232t56.5-23.5T464 1152m208-160V288q0-13-9.5-22.5T640 256H128q-13 0-22.5 9.5T96 288v704q0 13 9.5 22.5t22.5 9.5h512q13 0 22.5-9.5T672 992M480 144q0-16-16-16H304q-16 0-16 16t16 16h160q16 0 16-16m288-16v1024q0 52-38 90t-90 38H128q-52 0-90-38t-38-90V128q0-52 38-90t90-38h512q52 0 90 38t38 90"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Modx(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1395 709L781 323l92-151h855zM373 974L189 858V0l1183 743zm1019-135l147 95v858l-532-335zm-37-21l-500 802H0l356-571z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Money(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 896h384v-96h-128V352H910L762 489l77 80q42-37 55-57h2v288H768zm512-256q0 70-21 142t-59.5 134t-101.5 101t-138 39t-138-39t-101.5-101T661 782t-21-142t21-142t59.5-134T822 263t138-39t138 39t101.5 101t59.5 134t21 142m512 256V384q-106 0-181-75t-75-181H384q0 106-75 181t-181 75v512q106 0 181 75t75 181h1152q0-106 75-181t181-75m128-832v1152q0 26-19 45t-45 19H64q-26 0-45-19t-19-45V64q0-26 19-45T64 0h1792q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1262 1175q-54 9-110 9q-182 0-337-90T570 849t-90-337q0-192 104-357q-201 60-328.5 229T128 768q0 130 51 248.5t136.5 204t204 136.5t248.5 51q144 0 273.5-61.5T1262 1175m203-85q-94 203-283.5 324.5T768 1536q-156 0-298-61t-245-164t-164-245T0 768q0-153 57.5-292.5t156-241.5T449 69.5T739 1q44-2 61 39q18 41-15 72q-86 78-131.5 181.5T608 512q0 148 73 273t198 198t273 73q118 0 228-51q41-18 72 13q14 14 17.5 34t-4.5 38"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Motorcycle(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2301 780q12 103-22 198.5t-99 163.5t-158.5 106t-196.5 31q-161-11-279.5-125T1411 880q-12-111 27.5-210.5T1557 499l-71-107q-96 80-151 194t-55 244q0 27-18.5 46.5T1216 896H891q-23 164-149 274t-294 110q-185 0-316.5-131.5T0 832t131.5-316.5T448 384q76 0 152 27l24-45Q501 256 320 256h-64q-26 0-45-19t-19-45t19-45t45-19h128q78 0 145 13.5T645.5 180t71.5 39.5t51 36.5h627l-85-128h-222q-30 0-49-22.5T1025 53q4-23 23-38t43-15h253q33 0 53 28l70 105l114-114q19-19 46-19h101q26 0 45 19t19 45v128q0 26-19 45t-45 19h-179l115 172q131-63 275-36q143 26 244 134.5T2301 780M448 1152q115 0 203-72.5T762 896H448q-35 0-55-31q-18-32-1-63l147-277q-47-13-91-13q-132 0-226 94t-94 226t94 226t226 94m1408 0q132 0 226-94t94-226t-94-226t-226-94q-60 0-121 24l174 260q15 23 10 49t-27 40q-15 11-36 11q-35 0-53-29l-174-260q-93 95-93 225q0 132 94 226t226 94"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MousePointer(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1133 1043q31 30 14 69q-17 40-59 40H706l201 476q10 25 0 49t-34 35l-177 75q-25 10-49 0t-35-34l-191-452l-312 312q-19 19-45 19q-12 0-24-5q-40-17-40-59V64Q0 22 40 5q12-5 24-5q27 0 45 19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Music(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1536 96v1120q0 50-34 89t-86 60.5t-103.5 32t-96.5 10.5t-96.5-10.5t-103.5-32t-86-60.5t-34-89t34-89t86-60.5t103.5-32t96.5-10.5q105 0 192 39V526L640 763v709q0 50-34 89t-86 60.5t-103.5 32T320 1664t-96.5-10.5t-103.5-32t-86-60.5t-34-89t34-89t86-60.5t103.5-32T320 1280q105 0 192 39V352q0-31 19-56.5t49-35.5L1412 4q12-4 28-4q40 0 68 28t28 68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Neuter(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1152 576q0 221-147.5 384.5T640 1148v612q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23v-612q-217-24-364.5-187.5T0 576q0-117 45.5-223.5t123-184t184-123T576 0t223.5 45.5t184 123t123 184T1152 576m-576 448q185 0 316.5-131.5T1024 576T892.5 259.5T576 128T259.5 259.5T128 576t131.5 316.5T576 1024"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NewspaperO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1024 384H640v384h384zm128 640v128H512v-128zm0-768v640H512V256zm640 768v128h-512v-128zm0-256v128h-512V768zm0-256v128h-512V512zm0-256v128h-512V256zM256 1216V256H128v960q0 26 19 45t45 19t45-19t19-45m1664 0V128H384v1088q0 33-11 64h1483q26 0 45-19t19-45M2048 0v1216q0 80-56 136t-136 56H192q-80 0-136-56T0 1216V128h256V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectGroup(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2048 384h-128v1024h128v384h-384v-128H384v128H0v-384h128V384H0V0h384v128h1280V0h384zm-256-256v128h128V128zm-1664 0v128h128V128zm128 1536v-128H128v128zm1408-128v-128h128V384h-128V256H384v128H256v1024h128v128zm256 128v-128h-128v128zM1280 640h384v768H768v-256H384V384h896zm-768 384h640V512H512zm1024 256V768h-256v384H896v128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectUngroup(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2304 768h-128v640h128v384h-384v-128h-896v128H640v-384h128v-128H384v128H0v-384h128V384H0V0h384v128h896V0h384v384h-128v128h384V384h384zm-256-256v128h128V512zm-640-384v128h128V128zm-1280 0v128h128V128zm128 1152v-128H128v128zm1280-128h-128v128h128zm-1152 0h896v-128h128V384h-128V256H384v128H256v640h128zm512 512v-128H768v128zm1280 0v-128h-128v128zm-128-256V768h-128V640h-384v384h128v384h-384v-128H896v128h128v128h896v-128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Odnoklassniki(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M544 907q-188 0-321-133T90 454q0-188 133-321T544 0t321 133t133 321q0 187-133 320T544 907m0-677q-92 0-157.5 65.5T321 454q0 92 65.5 157.5T544 677t157.5-65.5T767 454q0-93-65.5-158.5T544 230m523 732q13 27 15 49.5t-4.5 40.5t-26.5 38.5t-42.5 37T947 1169q-115 73-315 94l73 72l267 267q30 31 30 74t-30 73l-12 13q-31 30-74 30t-74-30q-67-68-267-268l-267 268q-31 30-74 30t-73-30l-12-13q-31-30-31-73t31-74l267-267l72-72q-203-21-317-94q-39-25-61.5-41.5t-42.5-37t-26.5-38.5t-4.5-40.5T21 962q10-20 28-35t42-22t56 2t65 35q5 4 15 11t43 24.5t69 30.5t92 24t113 11q91 0 174-25.5T838 967l38-25q33-26 65-35t56-2t42 22t28 35"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OdnoklassnikiSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M927 452q0 66-46.5 112.5T768 611t-112.5-46.5T609 452t46.5-112.5T768 293t112.5 46.5T927 452m214 363q-10-20-28-32t-47.5-9.5T1005 801q-10 8-29 20t-81 32t-127 20t-124-18t-86-36l-27-18q-31-25-60.5-27.5T423 783t-28 32q-22 45-2 74.5t87 73.5q83 53 226 67l-51 52q-142 142-191 190q-22 22-22 52.5t22 52.5l9 9q22 22 52.5 22t52.5-22l191-191q114 115 191 191q22 22 52.5 22t52.5-22l9-9q22-22 22-52.5t-22-52.5l-191-190l-52-52q141-14 225-67q67-44 87-73.5t-2-74.5m-49-363q0-134-95-229t-229-95t-229 95t-95 229t95 229t229 95t229-95t95-229m444-164v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Opencart(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1524 1497q0 68-48 116t-116 48t-116.5-48t-48.5-116t48.5-116.5T1360 1332t116 48.5t48 116.5m-749 0q0 68-48.5 116T610 1661t-116-48t-48-116t48-116.5t116-48.5t116.5 48.5T775 1497M0 3q57 60 110.5 104.5t121 82t136 63t166 45.5t200 31.5t250 18.5t304 9.5T1660 360q139 0 244.5 5t181 16.5t124 27.5t71 39.5t24 51.5t-19.5 64t-56.5 76.5t-89.5 91T2023 836t-139 119q-185 157-286 247q29-51 76.5-109t94-105.5T1863 889t83-91.5t54-80.5t13-70t-45.5-55.5t-116.5-41t-204-23.5t-304-5q-168 2-314-6t-256-23t-204.5-41T409 400.5T286.5 338T195 271.5T127 200t-50.5-69.5t-40-68T0 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Openid(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1086 0v1536l-272 128q-228-20-414-102t-293-208.5T0 1081q0-140 100.5-263.5t275-205.5T767 504v172q-217 38-356.5 150T271 1081q0 152 154.5 267T814 1493V133zm669 582l37 390l-525-114l147-83q-119-70-280-99V504q277 33 481 157z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Opera(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1493 228q-165-110-359-110q-155 0-293 73T601 391q-75 93-119.5 218T433 875v42q4 141 48.5 266T601 1401q102 127 240 200t293 73q194 0 359-110q-121 108-274.5 168T896 1792q-29 0-43-1q-175-8-333-82t-272-193t-181-281T0 896q0-182 71-348t191-286T548 71T896 0h3q168 1 320.5 60.5T1493 228m299 668q0 192-77 362.5T1502 1555q-104 63-222 63q-137 0-255-84q154-56 253.5-233t99.5-405q0-227-99-404t-253-234q119-83 254-83q119 0 226 65q135 125 210.5 295t75.5 361"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OptinMonster(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M414 1675q-8 16-27 34.5t-37 25.5q-25 9-51.5-3.5T270 1700q-1-22 40-55t68-38q23-4 34 21.5t2 46.5m1341 0q7 16 26 34.5t38 25.5q25 9 51.5-3.5t27.5-31.5q2-22-39.5-55t-68.5-38q-22-4-33 21.5t-2 46.5m48-109q13 27 56.5 59.5t77.5 41.5q45 13 82-4.5t37-50.5q0-46-67.5-100.5T1873 1452q-40-5-63.5 37.5t-6.5 76.5m-1439 0q-13 27-56 59.5t-77 41.5q-45 13-82-4.5t-37-50.5q0-46 67.5-100.5T295 1452q40-5 63 37.5t6 76.5m730-1124h1q-41 0-76 15q27 8 44 30.5t17 49.5q0 35-27 60t-65 25q-52 0-80-43q-5 23-5 42q0 74 56 126.5t135 52.5q80 0 136-52.5t56-126.5t-56-126.5t-136-52.5m304-218q-99-109-220.5-131.5T932 137q27-60 82.5-96.5t118-39.5T1254 18t99.5 74.5T1398 224m750 1239q8 11-11 42q7 23 7 40q1 56-44.5 112.5T1990 1749t-118 37q-48 2-92-21.5t-66-65.5q-687 25-1259 0q-23 41-66.5 65t-92.5 22q-86-3-179.5-80.5T24 1545q2-22 7-40q-19-31-11-42q6-10 31-1q14-22 41-51q-7-29 2-38q11-10 39 4q29-20 59-34q0-29 13-37q23-12 51 16q35-5 61 2q18 4 38 19v-73q-11 0-18-2q-53-10-97-44.5t-55-87.5q-9-38 0-81q15-62 93-95q2-17 19-35.5t36-23.5t33 7.5t19 30.5h13q46 5 60 23q3 3 5 7q10-1 30.5-3.5T524 962q-15-11-30-17q-23-40-91-43q0-6 1-10q-62-2-118.5-18.5T201 826q-32-36-42.5-92T156 622q16-126 90-179q23-16 52-4.5t32 40.5q0 1 1.5 14t2.5 21t3 20t5.5 19t8.5 10q27 14 76 12q48-46 98-74q-40-4-162 14l47-46q61-58 163-111q145-73 282-86q-20-8-41-15.5t-47-14t-42.5-10.5t-47.5-11t-43-10q595-126 904 139q98 84 158 222q85 10 121-9h1q5-3 8.5-10t5.5-19t3-19.5t3-21.5l1-14q3-28 32-40t52 5q73 52 91 178q7 57-3.5 113t-42.5 91q-28 32-83.5 48.5T1769 893v10q-71 2-95 43q-14 5-31 17q11 1 32 3.5t30 3.5q1-5 5-8q16-18 60-23h13q5-18 19-30t33-8t36 23t19 36q79 32 93 95q9 40 1 81q-12 53-56 88t-97 44q-10 2-17 2q0 49-1 73q20-15 38-19q26-7 61-2q28-28 51-16q14 9 14 37q33 16 59 34q27-13 38-4q10 10 2 38q28 30 41 51q23-8 31 1m-275-952q0 29-9 54q82 32 112 132q4-37-9.5-98.5T1925 508q-20-19-36-17t-16 20m-78 100q35 42 47.5 108.5T1842 844q67-13 97-45q13-14 18-28q-3-64-31-114.5t-79-66.5q-15 15-52 21m-37 4q-30 0-44-1q42 115 53 239q21 0 43-3q16-68 1-135t-53-100M194 697q30-100 112-132q-9-25-9-54q0-18-16.5-20T245 508q-28 29-41.5 90.5T194 697m36 102q29 31 97 45q-13-58-.5-124.5T374 611q-37-6-52-21q-51 16-78.5 66T212 771q9 17 18 28m177 54q14-124 73-235q-19 4-55 18l-45 19v-1q-46 89-20 196q25 3 47 3m963 39q8 38 16.5 108.5t11.5 89.5q3 18 9.5 21.5t23.5-4.5q40-20 62-85.5t23-125.5q-24-2-146-4m-282-641q-116 0-199 82.5T806 532q0 117 83 199.5t199 82.5t199-82.5t83-199.5q0-116-83-198.5T1088 251m228 639q-105-2-211 0v-1q-1 27 2.5 86t13.5 66q29 14 93.5 14.5t95.5-10.5q9-3 11-39t-.5-69.5t-4.5-46.5m-268 199q8-4 9.5-48t-.5-88t-4-63v-1q-212 3-214 3q-4 20-7 62t0 83t14 46q34 15 101 16t101-10M654 900q-16 59 4.5 118.5T736 1103q15 8 24 5t12-21q3-16 8-90t10-103q-69 2-136 6m-127 126q3 23-34 36q132 141 271.5 240t305.5 154q172-49 310.5-146t293.5-250q-33-13-30-34q0-2 .5-3.5t1.5-3t1-2.5v-1v1q-17-2-50-5.5t-48-4.5q-26 90-82 132q-51 38-82-1q-5-6-9-14q-7-13-17-62q-2 5-5 9t-7.5 7t-8 5.5t-9.5 4l-10 2.5l-12 2l-12 1.5l-13.5 1l-13.5.5q-106 9-163-11q-4 17-10 26.5t-21 15t-23 7t-36 3.5q-6 1-9 1q-179 17-203-40q-2 63-56 54q-47-8-91-54q-12-13-20-26q-17-29-26-65q-58 6-87 10q1 2 4 10m-84 628q3-14 3-30q-17-71-51-130t-73-70q-41-12-101.5 14.5t-104.5 80T77 1626q35 53 100 93t119 42q51 2 94-28t53-79m3-171q23 63 27 119q195-113 392-174q-98-52-180.5-120T505 1143q-6 4-29 13q0 1-1 4t-1 5q31 18 22 37q-12 23-56 34q-10 13-29 24h-1q-2 83 1 150q19 34 35 73m69 166q532 21 1145 0q-254-147-428-196q-76 35-156 57q-8 3-16 0q-65-21-129-49q-208 60-416 188h-1v1q1 0 1-1m1184-46q4-54 28-120q14-38 33-71l-1 1q3-77 3-153q-15-8-30-25q-42-9-56-33q-9-20 22-38q-2-4-2-9q-16-4-28-12q-204 190-383 284q198 59 414 176m392 23q5-54-39-107.5t-104-80t-102-14.5q-38 11-72.5 70.5T1722 1624q0 16 3 30q10 49 53 79t94 28q54-2 119-42t100-93"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pagelines(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1402 1103q-32 80-76 138t-91 88.5t-99 46.5t-101.5 14.5t-96.5-8.5t-86.5-22t-69.5-27.5t-46-22.5l-17-10q-113 228-289.5 359.5T45 1792q-19 0-32-13t-13-32t13-31.5t32-12.5q173-1 322.5-107.5T619 1301q-36 14-72 23t-83 13t-91-2.5t-93-28.5t-92-59t-84.5-100T29 1001q114-47 214-57t167.5 7.5T535 1008t88.5 77t56.5 82q53-131 79-291q-7 1-18 2.5t-46.5 2.5t-69.5-.5t-81.5-10t-88.5-23t-84-42.5t-75-65t-54.5-94.5T213 518q70-28 133.5-36.5t112.5 1t92 30t73.5 50t56 61t42 63t27.5 56t16 39.5l4 16q12-122 12-195q-8-6-21.5-16t-49-44.5T648 471t-54-93t-33-112.5t12-127T643 0q73 25 127.5 61.5T855 138t48 85t20.5 89t-.5 85.5t-13 76.5t-19 62t-17 42l-7 15q1 4 1 50t-1 72q3-7 10-18.5t30.5-43t50.5-58t71-55.5t91.5-44.5t112-14.5t132.5 24q-2 78-21.5 141.5t-50 104.5t-69.5 71.5t-81.5 45.5t-84.5 24t-80 9.5t-67.5-1T864 896l-17-3q-23 147-73 283q6-7 18-18.5t49.5-41T919 1064t99.5-42t117.5-20t129 23.5t137 77.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaintBrush(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1615 0q70 0 122.5 46.5T1790 163q0 63-45 151q-332 629-465 752q-97 91-218 91q-126 0-216.5-92.5T755 845q0-128 92-212l638-579q59-54 130-54M706 1034q39 76 106.5 130t150.5 76l1 71q4 213-129.5 347T486 1792q-123 0-218-46.5T115.5 1618T29 1435T0 1215q7 5 41 30t62 44.5t59 36.5t46 17q41 0 55-37q25-66 57.5-112.5t69.5-76t88-47.5t103-25.5t125-10.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperPlane(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1764 11q33 24 27 64l-256 1536q-5 29-32 45q-14 8-31 8q-11 0-24-5l-453-185l-242 295q-18 23-49 23q-13 0-22-4q-19-7-30.5-23.5T640 1728v-349l864-1059l-1069 925l-395-162q-37-14-40-55q-2-40 32-59L1696 9q15-9 32-9q20 0 36 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperPlaneO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1764 43q33 24 27 64l-256 1536q-5 29-32 45q-14 8-31 8q-11 0-24-5l-527-215l-298 327q-18 21-47 21q-14 0-23-4q-19-7-30-23.5t-11-36.5v-452L40 1115q-37-14-40-55q-3-39 32-59L1696 41q35-21 68 2m-342 1499l221-1323l-1434 827l336 137l863-639l-478 797z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paperclip(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1404 1257q0 117-79 196t-196 79q-135 0-235-100L117 656Q4 541 4 385q0-159 110-270T383 4q158 0 273 113l605 606q10 10 10 22q0 16-30.5 46.5T1194 822q-13 0-23-10L565 205q-79-77-181-77q-106 0-179 75t-73 181q0 105 76 181l776 777q63 63 145 63q64 0 106-42t42-106q0-82-63-145L633 531q-26-24-60-24q-29 0-48 19t-19 48q0 32 25 59l410 410q10 10 10 22q0 16-31 47t-47 31q-12 0-22-10L441 723q-63-61-63-149q0-82 57-139t139-57q88 0 149 63l581 581q100 98 100 235"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paragraph(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1278 61v73q0 29-18.5 61t-42.5 32q-50 0-54 1q-26 6-32 31q-3 11-3 64v1152q0 25-18 43t-43 18H959q-25 0-43-18t-18-43V257H755v1218q0 25-17.5 43t-43.5 18H586q-26 0-43.5-18t-17.5-43V979q-147-12-245-59q-126-58-192-179q-64-117-64-259q0-166 88-286Q200 78 321 37Q432 0 738 0h479q25 0 43 18t18 43"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1536 64v1408q0 26-19 45t-45 19H960q-26 0-45-19t-19-45V64q0-26 19-45t45-19h512q26 0 45 19t19 45m-896 0v1408q0 26-19 45t-45 19H64q-26 0-45-19t-19-45V64q0-26 19-45T64 0h512q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseCircle(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M704 1056V480q0-14-9-23t-23-9H416q-14 0-23 9t-9 23v576q0 14 9 23t23 9h256q14 0 23-9t9-23m448 0V480q0-14-9-23t-23-9H864q-14 0-23 9t-9 23v576q0 14 9 23t23 9h256q14 0 23-9t9-23m384-288q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseCircleO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 0q209 0 385.5 103T1433 382.5T1536 768t-103 385.5t-279.5 279.5T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0m0 1312q148 0 273-73t198-198t73-273t-73-273t-198-198t-273-73t-273 73t-198 198t-73 273t73 273t198 198t273 73m96-224q-14 0-23-9t-9-23V480q0-14 9-23t23-9h192q14 0 23 9t9 23v576q0 14-9 23t-23 9zm-384 0q-14 0-23-9t-9-23V480q0-14 9-23t23-9h192q14 0 23 9t9 23v576q0 14-9 23t-23 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paw(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M780 344q0 60-19 113.5T698 550t-105 39q-76 0-138-57.5T363 396t-30-151q0-60 19-113.5T415 39T520 0q77 0 138.5 57.5t91.5 135T780 344M438 827q0 80-42 139t-119 59q-76 0-141.5-55.5T35 836T0 684q0-80 42-139.5T161 485q76 0 141.5 55.5t100.5 134T438 827m394-27q118 0 255 97.5t229 237t92 254.5q0 46-17 76.5t-48.5 45t-64.5 20t-76 5.5q-68 0-187.5-45T832 1446q-66 0-192.5 44.5T439 1535q-183 0-183-146q0-86 56-191.5T451.5 1005T639 859t193-59m239-211q-61 0-105-39t-63-92.5T884 344q0-74 30-151.5t91.5-135T1144 0q61 0 105 39t63 92.5t19 113.5q0 73-30 151t-92 135.5t-138 57.5m432-104q77 0 119 59.5t42 139.5q0 74-35 152t-100.5 133.5T1387 1025q-77 0-119-59t-42-139q0-74 35-152.5t100.5-134T1503 485"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paypal(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1519 646q18 84-4 204q-87 444-565 444h-44q-25 0-44 16.5t-24 42.5l-4 19l-55 346l-2 15q-5 26-24.5 42.5T708 1792H457q-21 0-33-15t-9-36q9-56 26.5-168t26.5-168t27-167.5t27-167.5q5-37 43-37h131q133 2 236-21q175-39 287-144q102-95 155-246q24-70 35-133q1-6 2.5-7.5t3.5-1t6 3.5q79 59 98 162m-172-282q0 107-46 236q-80 233-302 315q-113 40-252 42q0 1-90 1l-90-1q-100 0-118 96q-2 8-85 530q-1 10-12 10H57q-22 0-36.5-16.5T9 1538L241 67q5-29 27.5-48T320 0h598q34 0 97.5 13T1127 45q107 41 163.5 123t56.5 196"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pencil(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m363 1408l91-91l-235-235l-91 91v107h128v128zm523-928q0-22-22-22q-10 0-17 7l-542 542q-7 7-7 17q0 22 22 22q10 0 17-7l542-542q7-7 7-17m-54-192l416 416l-832 832H0v-416zm683 96q0 53-37 90l-166 166l-416-416l166-165q36-38 90-38q53 0 91 38l235 234q37 39 37 91"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PencilSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m404 980l152 152l-52 52h-56v-96h-96v-56zm414-390q14 13-3 30L524 911q-17 17-30 3q-14-13 3-30l291-291q17-17 30-3m-274 690l544-544l-288-288l-544 544v288zm608-608l92-92q28-28 28-68t-28-68l-152-152q-28-28-68-28t-68 28l-92 92zm384-384v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Percent(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1280 1152q0-52-38-90t-90-38t-90 38t-38 90t38 90t90 38t90-38t38-90M512 384q0-52-38-90t-90-38t-90 38t-38 90t38 90t90 38t90-38t38-90m1024 768q0 159-112.5 271.5T1152 1536t-271.5-112.5T768 1152t112.5-271.5T1152 768t271.5 112.5T1536 1152M1440 64q0 20-13 38L371 1510q-19 26-51 26H160q-26 0-45-19t-19-45q0-20 13-38L1165 26q19-26 51-26h160q26 0 45 19t19 45M768 384q0 159-112.5 271.5T384 768T112.5 655.5T0 384t112.5-271.5T384 0t271.5 112.5T768 384"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1408 1112q0 27-10 70.5t-21 68.5q-21 50-122 106q-94 51-186 51q-27 0-53-3.5t-57.5-12.5t-47-14.5T856 1357t-49-18q-98-35-175-83q-127-79-264-216T152 776q-48-77-83-175q-3-9-18-49t-20.5-55.5t-14.5-47T3.5 392T0 339q0-92 51-186Q107 52 157 31q25-11 68.5-21T296 0q14 0 21 3q18 6 53 76q11 19 30 54t35 63.5t31 53.5q3 4 17.5 25t21.5 35.5t7 28.5q0 20-28.5 50t-62 55t-62 53t-28.5 46q0 9 5 22.5t8.5 20.5t14 24t11.5 19q76 137 174 235t235 174q2 1 19 11.5t24 14t20.5 8.5t22.5 5q18 0 46-28.5t53-62t55-62t50-28.5q14 0 28.5 7t35.5 21.5t25 17.5q25 15 53.5 31t63.5 35t54 30q70 35 76 53q3 7 3 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1280 1065q0-11-2-16t-18-16.5t-40.5-25T1172 981t-45.5-25t-28.5-15q-5-3-19-13t-25-15t-21-5q-15 0-36.5 20.5t-39.5 45t-38.5 45T885 1039q-7 0-16.5-3.5T853 1029t-17-9.5t-14-8.5q-99-55-170-126.5T525 714q-2-3-8.5-14t-9.5-17t-6.5-15.5T497 651q0-13 20.5-33.5t45-38.5t45-39.5T628 503q0-10-5-21t-15-25t-13-19q-3-6-15-28.5T555 364t-26.5-47.5t-25-40.5t-16.5-18t-16-2q-48 0-101 22q-46 21-80 94.5T256 503q0 16 2.5 34t5 30.5t9 33t10 29.5t12.5 33t11 30q60 164 216.5 320.5T843 1230q6 2 30 11t33 12.5t29.5 10t33 9t30.5 5t34 2.5q57 0 130.5-34t94.5-80q22-53 22-101m256-777v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PieChart(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m768 890l546 546q-106 108-247.5 168T768 1664q-209 0-385.5-103T103 1281.5T0 896t103-385.5T382.5 231T768 128zm187 6h773q0 157-60 298.5T1500 1442zm709-128H896V0q209 0 385.5 103T1561 382.5T1664 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PiedPiper(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2304 0q-69 46-125 92t-89 81t-59.5 71.5T1993 302t-22 44.5t-14 29.5q-10 18-35.5 136.5T1873 677q-15 29-50 60.5t-67.5 50.5t-72.5 41t-48 28q-47 31-151 231q-341-14-630 158q-92 53-303 179q47-16 86-31t55-22l15-7q71-27 163-64.5t133.5-53.5t108-34.5T1254 1181q186-31 465 7q1 0 10 3q11 6 14 17t-3 22l-194 345q-15 29-47 22q-128-24-354-24q-146 0-402 44.5T351 1664q-82 1-149-13t-107-37t-61-40t-33-34l-1-1v-2q0-6 6-6q138 0 371-55q192-366 374.5-524T1135 794q5 0 14.5.5t38 5t55 12T1304 836t63 39.5t54 59t40 82.5l102-177q2-4 21-42.5t44.5-86.5t61-109.5t84-133.5T1874 331q66-82 128-141.5T2123.5 93t92.5-53.5T2304 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PiedPiperAlt(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1190 929q75-3 143.5 20.5t118 58.5t101 94.5t84 108T1712 1331q33 56 78.5 109t75.5 80.5t99 88.5q-48 30-108.5 57.5t-138.5 59t-114 47.5q-44-37-74-115t-43.5-164.5t-33-180.5t-42.5-168.5t-72.5-123T1216 973l-10 2l-6 4q4 5 13 14q6 5 28 23.5t25.5 22t19 18t18 20.5t11.5 21t10.5 27.5t4.5 31t4 40.5l1 33q1 26-2.5 57.5t-7.5 52t-12.5 58.5t-11.5 53q-35-1-101 9.5t-98 10.5q-39 0-72-10q-2-16-2-47q0-74 3-96q2-13 31.5-41.5t57-59t26.5-51.5q-24-2-43 24q-36 53-111.5 99.5T855 1336q-25 0-75.5-63T673 1133.5t-84-96.5q-6-4-27-30q-482 112-513 112q-16 0-28-11t-12-27q0-15 8.5-26.5T40 1040l486-106q-8-14-8-25t5.5-17.5t16-11.5t20-7t23-4.5T601 864q4-1 15.5-7.5T634 850q15 0 28 16t20 33q163-37 172-37q17 0 29.5 11t12.5 28q0 15-8.5 26T864 941l-182 40l-1 16q-1 26 81.5 117.5T867 1206q47 0 119-80t72-129q0-36-23.5-53t-51-18.5t-51-11.5t-23.5-34q0-16 10-34l-68-19q43-44 43-117q0-26-5-58q82-16 144-16q44 0 71.5 1.5t48.5 8.5t31 13.5t20.5 24.5t15.5 33.5t17 47.5t24 60l50-25q-3 40-23 60t-42.5 21t-40 6.5T1189 908zm60-235q-5-5-13.5-15.5t-12-14.5t-10.5-11.5t-10-10.5l-8-8l-8.5-7.5l-8-5l-8.5-4.5q-7-3-14.5-5t-20.5-2.5t-22-.5h-70q-126 0-217 43q16-30 36-46.5t54-29.5t65.5-36t46-36.5t50-55T1122 398q12 9 28 31.5t32 36.5t38 13l12-1v76l22 1q247-95 371-190q28-21 50-39t42.5-37.5t33-31t29.5-34t24-31t24.5-37t23-38t27-47.5t29.5-53l7-9q-2 53-43 139q-79 165-205 264t-306 142q-14 3-42 7.5t-50 9.5t-39 14q3 19 24.5 46t21.5 34q0 11-26 30m-221 921q39-26 131.5-47.5T1307 1546q9 0 22.5 15.5t28 42.5t26 50t24 51t14.5 33q-121 45-244 45q-61 0-125-11zM790 968l48-12l109 177l-73 48zm501 517q3 15 3 16q0 7-17.5 14.5t-46 13t-54 9.5t-53.5 7.5t-32 4.5l-7-43q21-2 60.5-8.5t72-10t60.5-3.5zM834 857l-96 20l-6-17q10-1 32.5-7t34.5-6q19 0 35 10m195 634h31l10 83l-41 12zM1918 1V0zm0 0l-1 5l-2 2l1-3zm0 0l1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PiedPiperPp(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1046 892q0 64-38 109t-91 45q-43 0-70-15V754q28-17 70-17q53 0 91 45.5t38 109.5M703 464q0 64-38 109.5T574 619q-43 0-70-15V327q28-17 70-17q53 0 91 45t38 109m562 431q0-134-88-229t-213-95q-20 0-39 3q-23 78-78 136q-87 95-211 101v636l211-41v-206q51 19 117 19q125 0 213-95t88-229M922 468q0-134-88.5-229T620 144q-74 0-141 36H293v840l211-41V773q55 19 116 19q125 0 213.5-95T922 468m614-180v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pinterest(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1536 768q0 209-103 385.5T1153.5 1433T768 1536q-111 0-218-32q59-93 78-164q9-34 54-211q20 39 73 67.5t114 28.5q121 0 216-68.5T1232 968t52-270q0-114-59.5-214T1052 321t-255-63q-105 0-196 29t-154.5 77t-109 110.5t-67 129.5T249 738q0 104 40 183t117 111q30 12 38-20q2-7 8-31t8-30q6-23-11-43q-51-61-51-151q0-151 104.5-259.5T776 389q151 0 235.5 82t84.5 213q0 170-68.5 289T852 1092q-61 0-98-43.5T731 944q8-35 26.5-93.5t30-103T799 672q0-50-27-83t-77-33q-62 0-105 57t-43 142q0 73 25 122l-99 418q-17 70-13 177q-206-91-333-281T0 768q0-209 103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinterestP(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 597q0-108 37.5-203.5T141 227t152-123t185-78T680 0q158 0 294 66.5T1195 260t85 287q0 96-19 188t-60 177t-100 149.5t-145 103t-189 38.5q-68 0-135-32t-96-88q-10 39-28 112.5t-23.5 95t-20.5 71t-26 71t-32 62.5t-46 77.5t-62 86.5l-14 5l-9-10q-15-157-15-188q0-92 21.5-206.5T348 972t52-203q-32-65-32-169q0-83 52-156t132-73q61 0 95 40.5T681 514q0 66-44 191t-44 187q0 63 45 104.5t109 41.5q55 0 102-25t78.5-68t56-95t38-110.5t20-111t6.5-99.5q0-173-109.5-269.5T653 163q-200 0-334 129.5T185 621q0 44 12.5 85t27 65t27 45.5T264 847q0 28-15 73t-37 45q-2 0-17-3q-51-15-90.5-56t-61-94.5t-32.5-108T0 597"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinterestSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1248 0q119 0 203.5 84.5T1536 288v960q0 119-84.5 203.5T1248 1536H523q85-122 108-210q9-34 53-209q21 39 73.5 67t112.5 28q181 0 295.5-147.5T1280 691q0-84-35-162.5t-96.5-139t-152.5-97T799 256q-104 0-194.5 28.5t-153 76.5T344 470.5t-66.5 128T256 731q0 102 39.5 180T412 1021q13 5 23.5 0t14.5-19q10-44 15-61q6-23-11-42q-50-62-50-150q0-150 103.5-256.5T778 386q149 0 232.5 81t83.5 210q0 168-67.5 286T853 1081q-60 0-97-43.5T733 934q8-34 26.5-92.5t29.5-102t11-74.5q0-49-26.5-81.5T698 551q-61 0-103.5 56.5T552 747q0 72 24 121l-98 414q-24 100-7 254H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plane(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1376 32q44 52 12 148t-108 172l-161 161l160 696q5 19-12 33l-128 96q-7 6-19 6q-4 0-7-1q-15-3-21-16L813 819l-259 259l53 194q5 17-8 31l-96 96q-9 9-23 9h-2q-15-2-24-13l-189-252L13 954q-11-7-13-23q-1-13 9-25l96-97q9-9 23-9q6 0 8 1l194 53l259-259L81 316q-14-8-17-24q-2-16 9-27l128-128q14-13 30-8l665 159l160-160q76-76 172-108t148 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1384 831L56 1569q-23 13-39.5 3T0 1536V64q0-26 16.5-36T56 31l1328 738q23 13 23 31t-23 31"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayCircle(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 0q209 0 385.5 103T1433 382.5T1536 768t-103 385.5t-279.5 279.5T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0m384 823q32-18 32-55t-32-55L608 393q-31-19-64-1q-32 19-32 56v640q0 37 32 56q16 8 32 8q17 0 32-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayCircleO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1184 768q0 37-32 55l-544 320q-15 9-32 9q-16 0-32-8q-32-19-32-56V448q0-37 32-56q33-18 64 1l544 320q32 18 32 55m128 0q0-148-73-273t-198-198t-273-73t-273 73t-198 198t-73 273t73 273t198 198t273 73t273-73t198-198t73-273m224 0q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plug(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1755 453q37 38 37 90.5t-37 90.5l-401 400l150 150l-160 160q-163 163-389.5 186.5T543 1430l-362 362H0v-181l362-362q-124-185-100.5-411.5T448 448l160-160l150 150l400-401q38-37 91-37t90 37t37 90.5t-37 90.5L939 619l234 234l401-400q38-37 91-37t90 37"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1408 608v192q0 40-28 68t-68 28H896v416q0 40-28 68t-68 28H608q-40 0-68-28t-28-68V896H96q-40 0-68-28T0 800V608q0-40 28-68t68-28h416V96q0-40 28-68t68-28h192q40 0 68 28t28 68v416h416q40 0 68 28t28 68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircle(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1216 832V704q0-26-19-45t-45-19H896V384q0-26-19-45t-45-19H704q-26 0-45 19t-19 45v256H384q-26 0-45 19t-19 45v128q0 26 19 45t45 19h256v256q0 26 19 45t45 19h128q26 0 45-19t19-45V896h256q26 0 45-19t19-45m320-64q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1280 832V704q0-26-19-45t-45-19H896V320q0-26-19-45t-45-19H704q-26 0-45 19t-19 45v320H320q-26 0-45 19t-19 45v128q0 26 19 45t45 19h320v320q0 26 19 45t45 19h128q26 0 45-19t19-45V896h320q26 0 45-19t19-45m256-544v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSquareO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1152 672v64q0 14-9 23t-23 9H768v352q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23V768H288q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h352V288q0-14 9-23t23-9h64q14 0 23 9t9 23v352h352q14 0 23 9t9 23m128 448V288q0-66-47-113t-113-47H288q-66 0-113 47t-47 113v832q0 66 47 113t113 47h832q66 0 113-47t47-113m128-832v832q0 119-84.5 203.5T1120 1408H288q-119 0-203.5-84.5T0 1120V288Q0 169 84.5 84.5T288 0h832q119 0 203.5 84.5T1408 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Podcast(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M994 1192q0 86-17 197q-31 215-55 313q-22 90-152 90t-152-90q-24-98-55-313q-17-110-17-197q0-168 224-168t224 168m542-424q0 240-134 434t-350 280q-8 3-15-3t-6-15q7-48 10-66q4-32 6-47q1-9 9-12q159-81 255.5-234t96.5-337q0-180-91-330.5T1070 203t-337-74q-124 7-237 61T302.5 330.5t-128 202T128 773q1 184 99 336.5T484 1341q7 3 9 12q3 21 6 45q1 9 5 32.5t6 35.5q1 9-6.5 15t-15.5 2q-148-58-261-169.5t-173.5-264T1 730q7-143 66-273.5t154.5-227T446.5 72T719 2q164-10 315.5 46.5t261 160.5t175 250.5T1536 768m-542-32q0 93-65.5 158.5T770 960t-158.5-65.5T546 736t65.5-158.5T770 512t158.5 65.5T994 736m288 32q0 122-53.5 228.5T1082 1174q-8 6-16 2t-10-14q-6-52-29-92q-7-10 3-20q58-54 91-127t33-155q0-111-58.5-204T938 422.5T726 386q-133 15-229 113T388 730q-10 92 23.5 176t98.5 144q10 10 3 20q-24 41-29 93q-2 9-10 13t-16-2q-95-74-148.5-183T258 757q3-131 69-244t177-181.5T745 257q144-7 268 60t196.5 187.5T1282 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PowerOff(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1536 896q0 156-61 298t-164 245t-245 164t-298 61t-298-61t-245-164t-164-245T0 896q0-182 80.5-343T307 283q43-32 95.5-25t83.5 50q32 42 24.5 94.5T461 487q-98 74-151.5 181T256 896q0 104 40.5 198.5T406 1258t163.5 109.5T768 1408t198.5-40.5T1130 1258t109.5-163.5T1280 896q0-121-53.5-228T1075 487q-42-32-49.5-84.5T1050 308q31-43 84-50t95 25q146 109 226.5 270t80.5 343M896 128v640q0 52-38 90t-90 38t-90-38t-38-90V128q0-52 38-90t90-38t90 38t38 90"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Print(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 1408h896v-256H384zm0-640h896V384h-160q-40 0-68-28t-28-68V128H384zm1152 64q0-26-19-45t-45-19t-45 19t-19 45t19 45t45 19t45-19t19-45m128 0v416q0 13-9.5 22.5t-22.5 9.5h-224v160q0 40-28 68t-68 28H352q-40 0-68-28t-28-68v-160H32q-13 0-22.5-9.5T0 1248V832q0-79 56.5-135.5T192 640h64V96q0-40 28-68t68-28h672q40 0 88 20t76 48l152 152q28 28 48 76t20 88v256h64q79 0 135.5 56.5T1664 832"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ProductHunt(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1150 762q0 56-39.5 95t-95.5 39H762V627h253q56 0 95.5 39.5T1150 762m179 0q0-130-91.5-222T1015 448H582v896h180v-269h253q130 0 222-91.5t92-221.5m463 134q0 182-71 348t-191 286t-286 191t-348 71t-348-71t-286-191t-191-286T0 896t71-348t191-286T548 71T896 0t348 71t286 191t191 286t71 348"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PuzzlePiece(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1664 1098q0 81-44.5 135t-123.5 54q-41 0-77.5-17.5t-59-38t-56.5-38t-71-17.5q-110 0-110 124q0 39 16 115t15 115v5q-22 0-33 1q-34 3-97.5 11.5T907 1561t-98 5q-61 0-103-26.5t-42-83.5q0-37 17.5-71t38-56.5t38-59T775 1192q0-79-54-123.5T586 1024q-84 0-143 45.5T384 1197q0 43 15 83t33.5 64.5t33.5 53t15 50.5q0 45-46 89q-37 35-117 35q-95 0-245-24q-9-2-27.5-4t-27.5-4l-13-2q-1 0-3-1q-2 0-2-1V512q2 1 17.5 3.5t34 5T73 524q150 24 245 24q80 0 117-35q46-44 46-89q0-22-15-50.5t-33.5-53T399 256t-15-83q0-82 59-127.5T587 0q80 0 134 44.5T775 168q0 41-17.5 77.5t-38 59t-38 56.5t-17.5 71q0 57 42 83.5T809 542q64 0 180-15t163-17v2q-1 2-3.5 17.5t-5 34t-3.5 21.5q-24 150-24 245q0 80 35 117q44 46 89 46q22 0 50.5-15t53-33.5T1408 911t83-15q82 0 127.5 59t45.5 143"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Qq(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M270 806q-8-19-8-52q0-20 11-49t24-45q-1-22 7.5-53t22.5-43q0-139 92.5-288.5T637 66Q776 0 961 0q133 0 266 55q49 21 90 48t71 56t55 68t42 74t32.5 84.5T1543 475t22 98l1 5q55 83 55 150q0 14-9 40t-9 38q0 1 1.5 3.5t3.5 5t2 3.5q77 114 120.5 214.5T1774 1241q0 43-19.5 100t-55.5 57q-9 0-19.5-7.5t-19-17.5t-19-26t-16-26.5t-13.5-26t-9-17.5q-1-1-3-1l-5 4q-59 154-132 223q20 20 61.5 38.5t69 41.5t35.5 65q-2 4-4 16t-7 18q-64 97-302 97q-53 0-110.5-9t-98-20t-104.5-30q-15-5-23-7q-14-4-46-4.5t-40-1.5q-41 45-127.5 65T598 1792q-35 0-69-1.5t-93-9t-101-20.5t-74.5-40t-32.5-64q0-40 10-59.5t41-48.5q11-2 40.5-13t49.5-12q4 0 14-2q2-2 2-4l-2-3q-48-11-108-105.5T202 1253l-5-3q-4 0-12 20q-18 41-54.5 74.5T53 1382h-1q-4 0-6-4.5t-5-5.5q-23-54-23-100q0-275 252-466"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Qrcode(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 1024v128H256v-128zm0-768v128H256V256zm768 0v128h-128V256zM128 1279h384V896H128zm0-767h384V128H128zm768 0h384V128H896zM640 768v640H0V768zm512 512v128h-128v-128zm256 0v128h-128v-128zm0-512v384h-384v-128H896v384H768V768h384v128h128V768zM640 0v640H0V0zm768 0v640H768V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Question(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M608 1000v240q0 16-12 28t-28 12H328q-16 0-28-12t-12-28v-240q0-16 12-28t28-12h240q16 0 28 12t12 28m316-600q0 54-15.5 101t-35 76.5t-55 59.5t-57.5 43.5t-61 35.5q-41 23-68.5 65T604 848q0 17-12 32.5T564 896H324q-15 0-25.5-18.5T288 840v-45q0-83 65-156.5T496 530q59-27 84-56t25-76q0-42-46.5-74T451 292q-65 0-108 29q-35 25-107 115q-13 16-31 16q-12 0-25-8L16 319Q3 309 .5 294T6 266Q166 0 470 0q80 0 161 31t146 83t106 127.5T924 400"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionCircle(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M896 1248v-192q0-14-9-23t-23-9H672q-14 0-23 9t-9 23v192q0 14 9 23t23 9h192q14 0 23-9t9-23m256-672q0-88-55.5-163T958 297t-170-41q-243 0-371 213q-15 24 8 42l132 100q7 6 19 6q16 0 25-12q53-68 86-92q34-24 86-24q48 0 85.5 26t37.5 59q0 38-20 61t-68 45q-63 28-115.5 86.5T640 892v36q0 14 9 23t23 9h192q14 0 23-9t9-23q0-19 21.5-49.5T972 829q32-18 49-28.5t46-35t44.5-48t28-60.5t12.5-81m384 192q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionCircleO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 1072v160q0 14-9 23t-23 9H688q-14 0-23-9t-9-23v-160q0-14 9-23t23-9h160q14 0 23 9t9 23m256-496q0 50-15 90t-45.5 69t-52 44t-59.5 36q-32 18-46.5 28t-26 24t-11.5 29v32q0 14-9 23t-23 9H688q-14 0-23-9t-9-23v-68q0-35 10.5-64.5t24-47.5t39-35.5t41-25.5t44.5-21q53-25 75-43t22-49q0-42-43.5-71.5T773 473q-56 0-95 27q-29 20-80 83q-9 12-25 12q-11 0-19-6l-108-82q-10-7-12-20t5-23q122-192 349-192q129 0 238.5 89.5T1136 576M768 128q-130 0-248.5 51t-204 136.5t-136.5 204T128 768t51 248.5t136.5 204t204 136.5t248.5 51t248.5-51t204-136.5t136.5-204t51-248.5t-51-248.5t-136.5-204t-204-136.5T768 128m768 640q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Quora(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1255 787q0-318-105-474.5T820 156q-222 0-326 157T390 787q0 316 104 471.5T820 1414q74 0 131-17q-22-43-39-73t-44-65t-53.5-56.5t-63-36T674 1152q-46 0-79 16l-49-97q105-91 276-91q132 0 215.5 54t150.5 155q67-149 67-402m390 632h117q3 27-2 67t-26.5 95t-58 100.5t-107 78T1406 1792q-71 0-130.5-19t-105.5-56t-79-78t-66-96q-97 27-205 27q-150 0-292.5-58t-253-158.5t-178-249T29 787q0-170 67.5-319.5T275 217T528.5 58T820 0q121 0 238.5 36t217 106t176 164.5t119.5 219t43 261.5q0 190-80.5 347.5T1315 1399q47 70 93.5 106.5T1513 1542q61 0 94-37.5t38-85.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuoteLeft(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 832v384q0 80-56 136t-136 56H192q-80 0-136-56T0 1216V512q0-104 40.5-198.5T150 150T313.5 40.5T512 0h64q26 0 45 19t19 45v128q0 26-19 45t-45 19h-64q-106 0-181 75t-75 181v32q0 40 28 68t68 28h224q80 0 136 56t56 136m896 0v384q0 80-56 136t-136 56h-384q-80 0-136-56t-56-136V512q0-104 40.5-198.5T1046 150t163.5-109.5T1408 0h64q26 0 45 19t19 45v128q0 26-19 45t-45 19h-64q-106 0-181 75t-75 181v32q0 40 28 68t68 28h224q80 0 136 56t56 136"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuoteRight(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 192v704q0 104-40.5 198.5T618 1258t-163.5 109.5T256 1408h-64q-26 0-45-19t-19-45v-128q0-26 19-45t45-19h64q106 0 181-75t75-181v-32q0-40-28-68t-68-28H192q-80 0-136-56T0 576V192q0-80 56-136T192 0h384q80 0 136 56t56 136m896 0v704q0 104-40.5 198.5T1514 1258t-163.5 109.5T1152 1408h-64q-26 0-45-19t-19-45v-128q0-26 19-45t45-19h64q106 0 181-75t75-181v-32q0-40-28-68t-68-28h-224q-80 0-136-56t-56-136V192q0-80 56-136t136-56h384q80 0 136 56t56 136"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ra(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 874q8-217 116-406t305-318h5q0 1-1 3q-8 8-28 33.5T364 263t-60 110.5T259.5 509t-14 150.5t39 157.5T393 971q50 50 102 69.5t90.5 11.5t69.5-23.5t47-32.5l16-16q39-51 53-116.5t6.5-122.5t-21-107t-26.5-80l-14-29q-10-25-30.5-49.5t-43-41T599 405t-35-19l-13-6l104-115q39 17 78 52t59 61l19 27q1-48-18.5-103.5T752 214l-20-31L893 0l160 181q-33 46-52.5 102.5T978 374l-4 33q22-37 61.5-72.5T1103 282l28-17l103 115q-44 14-85 50t-60 65l-19 29q-31 56-48 133.5t-7 170t57 156.5q33 45 77.5 60.5t85 5.5t76-26.5T1368 990l21-16q60-53 96.5-115t48.5-121.5t10-121.5t-18-118t-37-107.5t-45.5-93t-45-72T1364 178l-13-17q-14-13-7-13l10 3q40 29 62.5 46t62 50t64 58t58.5 65t55.5 77t45.5 88t38 103t23.5 117t10.5 136q3 259-108 465t-312 321t-456 115q-185 0-351-74t-283.5-198t-184-293T19 874"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Random(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M666 449q-60 92-137 273q-22-45-37-72.5T451.5 586t-51-56.5t-63-35T256 480H32q-14 0-23-9t-9-23V256q0-14 9-23t23-9h224q250 0 410 225m1126 799q0 14-9 23l-320 320q-9 9-23 9q-13 0-22.5-9.5t-9.5-22.5v-192q-32 0-85 .5t-81 1t-73-1t-71-5t-64-10.5t-63-18.5t-58-28.5t-59-40t-55-53.5t-56-69.5q59-93 136-273q22 45 37 72.5t40.5 63.5t51 56.5t63 35t81.5 14.5h256V928q0-14 9-23t23-9q12 0 24 10l319 319q9 9 9 23m0-896q0 14-9 23l-320 320q-9 9-23 9q-13 0-22.5-9.5T1408 672V480h-256q-48 0-87 15t-69 45t-51 61.5t-45 77.5q-32 62-78 171q-29 66-49.5 111t-54 105t-64 100t-74 83t-90 68.5t-106.5 42t-128 16.5H32q-14 0-23-9t-9-23v-192q0-14 9-23t23-9h224q48 0 87-15t69-45t51-61.5t45-77.5q32-62 78-171q29-66 49.5-111t54-105t64-100t74-83t90-68.5t106.5-42t128-16.5h256V32q0-14 9-23t23-9q12 0 24 10l319 319q9 9 9 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ravelry(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1111 1733q-6-1-11-4q-13-8-36-23t-86-65t-116.5-104.5t-112-140T660 1224q-17-3-175-37q66 213 235 362t391 184m-641-606l168 28q-25-76-41-167.5T578 842l-4-53q-84 82-121 224q5 65 17 114m110-609q-43 64-77 148q44-46 74-68zm1437 434q0-161-62-307t-167.5-252T1537 224.5T1233 162q-147 0-281 52.5T712 363q-30 58-45 160q60-51 143-83.5t158.5-43t143-13.5t108.5 1l40 3q33 1 53 15.5t24.5 33t6.5 37t-1 28.5q-126-11-227.5-.5t-183 43.5T790 615.5T659 714q4 36 11.5 92.5t35.5 178t62 179.5q123 6 247.5-14.5T1230 1096t162.5-67t109.5-59l37-24q22-16 39.5-20.5t30.5 5t17 34.5q14 97-39 121q-208 97-467 134q-135 20-317 16q41 96 110 176.5t137 127t130.5 79T1282 1662l39 12q143 23 263-15q195-99 314-289t119-418m74-37q-14 135-40 212q-70 208-181.5 346.5T1551 1727q-48 33-82 44q-72 26-163 16q-36 3-73 3q-283 0-504.5-173T433 1175q-1 0-4-.5t-5-.5q-6 50 2.5 112.5t26 115t36 98T520 1571l14 26q8 12 54 82q-71-38-124.5-106.5t-78.5-140t-39.5-137T328 1188l-2-42q-5-2-33.5-12.5t-48.5-18t-53-20.5t-57.5-25t-50-25.5t-42.5-27T16 992q19 10 50.5 25.5t113 45.5t145.5 38l2-32q11-149 94-290q41-202 176-365q28-115 81-214q15-28 32-45t49-32q158-74 303.5-104t302-11t306.5 97q220 115 333 336t87 474"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Recycle(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m836 1169l-15 368l-2 22l-420-29q-36-3-67-31.5t-47-65.5q-11-27-14.5-55t4-65t12-55t21.5-64t19-53q78 12 509 28M449 583l180 379l-147-92q-63 72-111.5 144.5t-72.5 125t-39.5 94.5t-18.5 63l-4 21L46 961q-17-26-18-56t6-47l8-18q35-63 114-188L16 566zm1231 517l-188 359q-12 29-36.5 46.5T1412 1526l-18 4q-71 7-219 12l8 164l-230-367l211-362l7 173q170 16 283 5t170-33zM895 176q-47 63-265 435L313 424l-19-12L519 56q20-31 60-45t80-10q24 2 48.5 12t42 21T791 67t36 34.5t36 39.5t32 35m655 307l212 363q18 37 12.5 76t-27.5 74q-13 20-33 37t-38 28t-48.5 22t-47 16t-51.5 14t-46 12q-34-72-265-436l313-195zm-143-226l142-83l-220 373l-419-20l151-86q-34-89-75-166t-75.5-123.5t-64.5-80T799 25l-17-13l405 1q31-3 58 10.5t39 28.5l11 15q39 61 112 190"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reddit(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1095 1167q16 16 0 31q-62 62-199 62t-199-62q-16-15 0-31q6-6 15-6t15 6q48 49 169 49q120 0 169-49q6-6 15-6t15 6M788 986q0 37-26 63t-63 26t-63.5-26t-26.5-63q0-38 26.5-64t63.5-26t63 26.5t26 63.5m395 0q0 37-26.5 63t-63.5 26t-63-26t-26-63t26-63.5t63-26.5t63.5 26t26.5 64m251-120q0-49-35-84t-85-35t-86 36q-130-90-311-96l63-283l200 45q0 37 26 63t63 26t63.5-26.5T1359 448t-26.5-63.5T1269 358q-54 0-80 50l-221-49q-19-5-25 16l-69 312q-180 7-309 97q-35-37-87-37q-50 0-85 35t-35 84q0 35 18.5 64t49.5 44q-6 27-6 56q0 142 140 243t337 101q198 0 338-101t140-243q0-32-7-57q30-15 48-43.5t18-63.5m358 30q0 182-71 348t-191 286t-286 191t-348 71t-348-71t-286-191t-191-286T0 896t71-348t191-286T548 71T896 0t348 71t286 191t191 286t71 348"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedditAlien(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 846q0 58-29.5 105.5T1683 1024q12 46 12 96q0 155-106.5 287T1298 1615.5T898 1692t-399.5-76.5t-290-208.5T102 1120q0-47 11-94q-51-25-82-73.5T0 846q0-82 58-140.5T199 647q85 0 145 63q218-152 515-162L975 27q3-13 15-21t26-5l369 81q18-37 54-59.5T1518 0q62 0 106 43.5t44 105.5t-44 106t-106 44t-105.5-43.5T1369 150l-334-74l-104 472q300 9 519 160q58-61 143-61q83 0 141 58.5t58 140.5M418 1045q0 62 43.5 106t105.5 44t106-44t44-106t-44-105.5T567 896q-61 0-105 44t-44 105m810 355q11-11 11-26t-11-26q-10-10-25-10t-26 10q-41 42-121 62t-160 20t-160-20t-121-62q-11-10-26-10t-25 10q-11 10-11 25.5t11 26.5q43 43 118.5 68t122.5 29.5t91 4.5t91-4.5t122.5-29.5t118.5-68m-3-205q62 0 105.5-44t43.5-106q0-61-44-105t-105-44q-62 0-106 43.5t-44 105.5t44 106t106 44"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedditSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M939 1001q13 13 0 26q-53 53-171 53t-171-53q-13-13 0-26q5-6 13-6t13 6q42 42 145 42t145-42q5-6 13-6t13 6M676 845q0 31-23 54t-54 23t-54-23t-23-54q0-32 22.5-54.5T599 768t54.5 22.5T676 845m338 0q0 31-23 54t-54 23t-54-23t-23-54q0-32 22.5-54.5T937 768t54.5 22.5T1014 845m215-103q0-42-30-72t-73-30q-42 0-73 31q-113-78-267-82l54-243l171 39q1 32 23.5 54t53.5 22q32 0 54.5-22.5T1165 384t-22.5-54.5T1088 307q-48 0-69 43l-189-42q-17-5-21 13l-60 268q-154 6-265 83q-30-32-74-32q-43 0-73 30t-30 72q0 30 16 55t42 38q-5 25-5 48q0 122 120 208.5t289 86.5q170 0 290-86.5T1179 883q0-25-6-49q25-13 40.5-37.5T1229 742m307-454v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Refresh(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1511 928q0 5-1 7q-64 268-268 434.5T764 1536q-146 0-282.5-55T238 1324l-129 129q-19 19-45 19t-45-19t-19-45V960q0-26 19-45t45-19h448q26 0 45 19t19 45t-19 45l-137 137q71 66 161 102t187 36q134 0 250-65t186-179q11-17 53-117q8-23 30-23h192q13 0 22.5 9.5t9.5 22.5m25-800v448q0 26-19 45t-45 19h-448q-26 0-45-19t-19-45t19-45l138-138Q969 256 768 256q-134 0-250 65T332 500q-11 17-53 117q-8 23-30 23H50q-13 0-22.5-9.5T18 608v-7q65-268 270-434.5T768 0q146 0 284 55.5T1297 212l130-129q19-19 45-19t45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Registered(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1042 703q0-88-60-121q-33-18-117-18H742v281h162q66 0 102-37t36-105m52 285l205 373q8 17-1 31q-8 16-27 16h-152q-20 0-28-17l-194-365H742v350q0 14-9 23t-23 9H576q-14 0-23-9t-9-23V416q0-14 9-23t23-9h294q128 0 190 24q85 31 134 109t49 180q0 92-42.5 165.5T1085 972q6 10 9 16M896 160q-150 0-286 58.5t-234.5 157t-157 234.5T160 896t58.5 286t157 234.5t234.5 157t286 58.5t286-58.5t234.5-157t157-234.5t58.5-286t-58.5-286t-157-234.5t-234.5-157T896 160m896 736q0 182-71 348t-191 286t-286 191t-348 71t-348-71t-286-191t-191-286T0 896t71-348t191-286T548 71T896 0t348 71t286 191t191 286t71 348"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Renren(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1133 1442q-171 94-368 94q-196 0-367-94q138-87 235.5-211T765 963q35 144 132.5 268t235.5 211M638 14v485q0 252-126.5 459.5T181 1265Q0 1050 0 770q0-187 83.5-349.5T313 151T638 14m898 756q0 280-181 495q-204-99-330.5-306.5T898 499V14q179 30 325 137t229.5 269.5T1536 770"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Repeat(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1536 128v448q0 26-19 45t-45 19h-448q-42 0-59-40q-17-39 14-69l138-138Q969 256 768 256q-104 0-198.5 40.5T406 406T296.5 569.5T256 768t40.5 198.5T406 1130t163.5 109.5T768 1280q119 0 225-52t179-147q7-10 23-12q15 0 25 9l137 138q9 8 9.5 20.5t-7.5 22.5q-109 132-264 204.5T768 1536q-156 0-298-61t-245-164t-164-245T0 768t61-298t164-245T470 61T768 0q147 0 284.5 55.5T1297 212l130-129q29-31 70-14q39 17 39 59"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Retweet(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1280 1248q0 13-9.5 22.5t-22.5 9.5H288q-8 0-13.5-2t-9-7t-5.5-8t-3-11.5t-1-11.5V640H64q-26 0-45-19T0 576q0-24 15-41l320-384q19-22 49-22t49 22l320 384q15 17 15 41q0 26-19 45t-45 19H512v384h576q16 0 25 11l160 192q7 10 7 21m640-416q0 24-15 41l-320 384q-20 23-49 23t-49-23l-320-384q-15-17-15-41q0-26 19-45t45-19h192V384H832q-16 0-25-12L647 180q-7-9-7-20q0-13 9.5-22.5T672 128h960q8 0 13.5 2t9 7t5.5 8t3 11.5t1 11.5v600h192q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Road(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1079 740v-4l-24-320q-1-13-11-22.5t-23-9.5H835q-13 0-23 9.5T801 416l-24 320v4q-1 12 8 20t21 8h244q12 0 21-8t8-20m759 467q0 73-46 73h-704q13 0 22-9.5t8-22.5l-20-256q-1-13-11-22.5t-23-9.5H792q-13 0-23 9.5T758 992l-20 256q-1 13 8 22.5t22 9.5H64q-46 0-46-73q0-54 26-116L461 47q8-19 26-33t38-14h339q-13 0-23 9.5T830 32l-15 192q-1 14 8 23t22 9h166q13 0 22-9t8-23l-15-192q-1-13-11-22.5T992 0h339q20 0 38 14t26 33l417 1044q26 62 26 116"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rocket(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1440 320q0-40-28-68t-68-28t-68 28t-28 68t28 68t68 28t68-28t28-68m224-288q0 249-75.5 430.5T1335 823q-81 80-195 176l-20 379q-2 16-16 26l-384 224q-7 4-16 4q-12 0-23-9l-64-64q-13-14-8-32l85-276l-281-281l-276 85q-3 1-9 1q-14 0-23-9l-64-64q-17-19-5-39l224-384q10-14 26-16l379-20q96-114 176-195q188-187 358-258t431-71q14 0 24 9.5t10 22.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RotateLeft(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1536 768q0 156-61 298t-164 245t-245 164t-298 61q-172 0-327-72.5T177 1259q-7-10-6.5-22.5t8.5-20.5l137-138q10-9 25-9q16 2 23 12q73 95 179 147t225 52q104 0 198.5-40.5T1130 1130t109.5-163.5T1280 768t-40.5-198.5T1130 406T966.5 296.5T768 256q-98 0-188 35.5T420 393l137 138q31 30 14 69q-17 40-59 40H64q-26 0-45-19T0 576V128q0-42 40-59q39-17 69 14l130 129Q346 111 483.5 55.5T768 0q156 0 298 61t245 164t164 245t61 298"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rouble(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1043 437q0-100-65-162t-171-62H487v448h320q106 0 171-62t65-162m237 0q0 193-126.5 315T827 874H487v118h505q14 0 23 9t9 23v128q0 14-9 23t-23 9H487v192q0 14-9.5 23t-22.5 9H288q-14 0-23-9t-9-23v-192H32q-14 0-23-9t-9-23v-128q0-14 9-23t23-9h224V874H32q-14 0-23-9t-9-23V693q0-13 9-22.5t23-9.5h224V32q0-14 9-23t23-9h539q200 0 326.5 122T1280 437"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RssSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 1152q0-53-37.5-90.5T384 1024t-90.5 37.5T256 1152t37.5 90.5T384 1280t90.5-37.5T512 1152m351 94q-13-233-176.5-396.5T290 673q-14-1-24 9t-10 23v128q0 13 8.5 22t21.5 10q154 11 264 121t121 264q1 13 10 21.5t22 8.5h128q13 0 23-10t9-24m384 1q-5-154-56-297.5t-139.5-260t-205-205t-260-139.5T289 289q-14-1-23 9q-10 10-10 23v128q0 13 9 22t22 10q204 7 378 111.5T943.5 871t111.5 378q1 13 10 22t22 9h128q13 0 23-10q11-9 9-23m289-959v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Safari(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M949 893q0 26-16.5 45T891 957q-26 0-45-16.5T827 899q0-26 17-45t42-19t44 16.5t19 41.5m15 58l350-581q-9 8-67.5 62.5T1121 549T984.5 676t-117 110.5T817 838l-349 580q7-7 67-62t126-116.5t136-127t117-111t50-50.5m647-55q0 201-104 371q-3-2-17-11t-26.5-16.5t-16.5-7.5q-13 0-13 13q0 10 59 44q-74 112-184.5 190.5T1067 1590l-16-67q-1-10-15-10q-5 0-8 5.5t-2 9.5l16 68q-72 15-146 15q-199 0-372-105q1-2 13-20.5t21.5-33.5t9.5-19q0-13-13-13q-6 0-17 14.5t-22.5 34.5t-13.5 23q-113-75-192-187.5T200 1060l69-15q10-3 10-15q0-5-5.5-8t-10.5-2l-68 15q-14-72-14-139q0-206 109-379q2 1 18.5 12t30 19t17.5 8q13 0 13-12q0-6-12.5-15.5T324 507l-20-12q77-112 189-189t244-107l15 67q2 10 15 10q5 0 8-5.5t2-10.5l-15-66q71-13 134-13q204 0 379 109q-39 56-39 65q0 13 12 13q11 0 48-64q111 75 187.5 186T1591 731l-56 12q-10 2-10 16q0 5 5.5 8t9.5 2l57-13q14 72 14 140m85 0q0-163-63.5-311T1462 330t-255-170.5T896 96t-311 63.5T330 330T159.5 585T96 896t63.5 311T330 1462t255 170.5t311 63.5t311-63.5t255-170.5t170.5-255t63.5-311m96 0q0 182-71 348t-191 286t-286 191t-348 71t-348-71t-286-191t-191-286T0 896t71-348t191-286T548 71T896 0t348 71t286 191t191 286t71 348"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scribd(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1468 1549q0 89-63 152.5t-153 63.5t-153.5-63.5T1035 1549q0-90 63.5-153.5T1252 1332t153 63.5t63 153.5m-233-281q-115 15-192.5 102.5T965 1576q0 74 33 138q-146 78-379 78q-109 0-201-21t-153.5-54.5T154 1640t-76-85t-44.5-83t-23.5-66.5t-6-39.5q0-19 4.5-42.5t18.5-56t36.5-58t64-43.5t94.5-18t94 17.5t63 41t35.5 53t17.5 49t4 33.5q0 34-23 81q28 27 82 42t93 17l40 1q115 0 190-51t75-133q0-26-9-48.5t-31.5-44.5t-49.5-41t-74-44t-93.5-47.5T516 1017q-28-13-43-20q-116-55-187-100T163.5 795t-72-125.5T71 507q0-78 20.5-150t66-137.5t112.5-114t166.5-77T658 0q120 0 220 26t164.5 67t109.5 94t64 105.5t19 103.5q0 46-15 82.5t-36.5 58t-48.5 36t-49 19.5t-39 5h-40l-39-5l-44-14l-41-28l-37-46l-24-70.5l-10-97.5q-15-16-59-25.5T672 300l-37-1q-68 0-117.5 31T447 400t-21 76q0 24 5 43t24 46t53 51t97 53.5T755 728q76 25 138.5 53.5t109 55.5t83 59t60.5 59.5t41 62.5t26.5 62t14.5 63.5t6 62t1 62.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1152 704q0-185-131.5-316.5T704 256T387.5 387.5T256 704t131.5 316.5T704 1152t316.5-131.5T1152 704m512 832q0 52-38 90t-90 38q-54 0-90-38l-343-342q-179 124-399 124q-143 0-273.5-55.5t-225-150t-150-225T0 704t55.5-273.5t150-225t225-150T704 0t273.5 55.5t225 150t150 225T1408 704q0 220-124 399l343 343q37 37 37 90"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchMinus(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1024 672v64q0 13-9.5 22.5T992 768H416q-13 0-22.5-9.5T384 736v-64q0-13 9.5-22.5T416 640h576q13 0 22.5 9.5t9.5 22.5m128 32q0-185-131.5-316.5T704 256T387.5 387.5T256 704t131.5 316.5T704 1152t316.5-131.5T1152 704m512 832q0 53-37.5 90.5T1536 1664q-54 0-90-38l-343-342q-179 124-399 124q-143 0-273.5-55.5t-225-150t-150-225T0 704t55.5-273.5t150-225t225-150T704 0t273.5 55.5t225 150t150 225T1408 704q0 220-124 399l343 343q37 37 37 90"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchPlus(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1024 672v64q0 13-9.5 22.5T992 768H768v224q0 13-9.5 22.5T736 1024h-64q-13 0-22.5-9.5T640 992V768H416q-13 0-22.5-9.5T384 736v-64q0-13 9.5-22.5T416 640h224V416q0-13 9.5-22.5T672 384h64q13 0 22.5 9.5T768 416v224h224q13 0 22.5 9.5t9.5 22.5m128 32q0-185-131.5-316.5T704 256T387.5 387.5T256 704t131.5 316.5T704 1152t316.5-131.5T1152 704m512 832q0 53-37.5 90.5T1536 1664q-54 0-90-38l-343-342q-179 124-399 124q-143 0-273.5-55.5t-225-150t-150-225T0 704t55.5-273.5t150-225t225-150T704 0t273.5 55.5t225 150t150 225T1408 704q0 220-124 399l343 343q37 37 37 90"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sellsy(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1500 1211V478q0-21-15-36t-35-15h-93q-20 0-35 15t-15 36v733q0 20 15 35t35 15h93q20 0 35-15t15-35m-284 0V680q0-20-15-35t-35-15h-101q-20 0-35 15t-15 35v531q0 20 15 35t35 15h101q20 0 35-15t15-35m-292 0V782q0-20-15-35t-35-15H773q-20 0-35 15t-15 35v429q0 20 15 35t35 15h101q20 0 35-15t15-35m-292 0V849q0-20-15-35t-35-15H481q-20 0-35 15t-15 35v362q0 20 15 35t35 15h101q20 0 35-15t15-35m1416-146q0 166-118 284t-284 118H402q-166 0-284-118T0 1065q0-116 63-214.5T231 702q-10-34-10-73q0-113 80.5-193.5T495 355q102 0 180 67q45-183 194-300T1207 5q149 0 275 73.5T1681.5 278t73.5 275q0 66-14 122q135 33 221 142.5t86 247.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Server(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 1280h1024v-128H128zm0-512h1024V640H128zm1568 448q0-40-28-68t-68-28t-68 28t-28 68t28 68t68 28t68-28t28-68M128 256h1024V128H128zm1568 448q0-40-28-68t-68-28t-68 28t-28 68t28 68t68 28t68-28t28-68m0-512q0-40-28-68t-68-28t-68 28t-28 68t28 68t68 28t68-28t28-68m96 832v384H0v-384zm0-512v384H0V512zm0-512v384H0V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareAlt(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1216 896q133 0 226.5 93.5T1536 1216t-93.5 226.5T1216 1536t-226.5-93.5T896 1216q0-12 2-34l-360-180q-92 86-218 86q-133 0-226.5-93.5T0 768t93.5-226.5T320 448q126 0 218 86l360-180q-2-22-2-34q0-133 93.5-226.5T1216 0t226.5 93.5T1536 320t-93.5 226.5T1216 640q-126 0-218-86L638 734q2 22 2 34t-2 34l360 180q92-86 218-86"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareAltSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1280 1067q0-88-62.5-151T1067 853q-84 0-145 58L681 791q2-16 2-23t-2-23l241-120q61 58 145 58q88 0 150.5-63t62.5-151t-62.5-150.5T1067 256t-151 62.5T853 469q0 7 2 23L614 612q-62-57-145-57q-88 0-150.5 62.5T256 768t62.5 150.5T469 981q83 0 145-57l241 120q-2 16-2 23q0 88 63 150.5t151 62.5t150.5-62.5T1280 1067m256-779v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1005 973l352-352q19-19 19-45t-19-45l-352-352q-30-31-69-14q-40 17-40 59v160q-119 0-216 19.5t-162.5 51t-114 79T327 629t-44.5 109T261 849.5T256 960q0 181 167 404q11 12 25 12q7 0 13-3q22-9 19-33q-44-354 62-473q46-52 130-75.5T896 768v160q0 42 40 59q12 5 24 5q26 0 45-19m531-685v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareSquareO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1408 989v259q0 119-84.5 203.5T1120 1536H288q-119 0-203.5-84.5T0 1248V416q0-119 84.5-203.5T288 128h255q13 0 22.5 9.5T575 160q0 27-26 32q-77 26-133 60q-10 4-16 4H288q-66 0-113 47t-47 113v832q0 66 47 113t113 47h832q66 0 113-47t47-113v-214q0-19 18-29q28-13 54-37q16-16 35-8q21 9 21 29m237-496l-384 384q-18 19-45 19q-12 0-25-5q-39-17-39-59V640H992q-323 0-438 131q-119 137-74 473q3 23-20 34q-8 2-12 2q-16 0-26-13q-10-14-21-31t-39.5-68.5T312 1068t-38.5-114T256 832q0-49 3.5-91t14-90t28-88t47-81.5t68.5-74t94.5-61.5T636 297.5T795.5 267T992 256h160V64q0-42 39-59q13-5 25-5q26 0 45 19l384 384q19 19 19 45t-19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shield(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1088 832V192H640v1137q119-63 213-137q235-184 235-360m192-768v768q0 86-33.5 170.5t-83 150t-118 127.5T919 1383t-121 77.5t-89.5 49.5t-42.5 20q-12 6-26 6t-26-6q-16-7-42.5-20t-89.5-49.5t-121-77.5t-126.5-103t-118-127.5t-83-150T0 832V64q0-26 19-45T64 0h1152q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ship(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1811 1555q19-19 45-19t45 19l128 128l-90 90l-83-83l-83 83q-18 19-45 19t-45-19l-83-83l-83 83q-19 19-45 19t-45-19l-83-83l-83 83q-19 19-45 19t-45-19l-83-83l-83 83q-19 19-45 19t-45-19l-83-83l-83 83q-19 19-45 19t-45-19l-83-83l-83 83q-19 19-45 19t-45-19l-83-83l-83 83q-19 19-45 19t-45-19L19 1645l90-90l83 83l83-83q19-19 45-19t45 19l83 83l83-83q19-19 45-19t45 19l83 83l83-83q19-19 45-19t45 19l83 83l83-83q19-19 45-19t45 19l83 83l83-83q19-19 45-19t45 19l83 83l83-83q19-19 45-19t45 19l83 83zm-1574-38q-19 19-45 19t-45-19L19 1389l90-90l83 82l83-82q19-19 45-19t45 19l83 82l64-64v-293L302 710q-17-26-7-56.5t40-40.5l177-58V256h128V128h256V0h256v128h256v128h128v299l177 58q30 10 40 40.5t-7 56.5l-210 314v293l19-18q19-19 45-19t45 19l83 82l83-82q19-19 45-19t45 19l128 128l-90 90l-83-83l-83 83q-18 19-45 19t-45-19l-83-83l-83 83q-19 19-45 19t-45-19l-83-83l-83 83q-19 19-45 19t-45-19l-83-83l-83 83q-19 19-45 19t-45-19l-83-83l-83 83q-19 19-45 19t-45-19l-83-83l-83 83q-19 19-45 19t-45-19l-83-83zM640 384v128l384-128l384 128V384h-128V256H768v128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shirtsinbulk(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h1536v1392l-776 338L0 1392zm1436 1327V401H100v926l661 294zm0-1026V100H100v201zM181 599v115h-37V599zm0 148v115h-37V747zm0 148v115h-37V895zm0 148v115h-37v-115zm0 148v115h-37v-115zm26 143l15-34l105 47l-15 33zm136 60l15-34l105 46l-15 34zm135 60l15-34l105 46l-15 34zm136 59l15-33l104 46l-15 34zm183 13l105-46l15 33l-105 47zm135-60l105-46l15 34l-105 46zm136-60l105-46l15 34l-105 46zm135-59l105-47l15 34l-105 46zM259 147v36H145v-36zm162 0v36H306v-36zm162 0v36H468v-36zm161 0v36H630v-36zm162 0v36H792v-36zm162 0v36H953v-36zm162 0v36h-115v-36zm161 0v36h-114v-36zM181 487v79h-37V451h115v36zm240-36v36H306v-36zm162 0v36H468v-36zm161 0v36H630v-36zm162 0v36H792v-36zm162 0v36H953v-36zm162 0v36h-115v-36zm125 115v-79h-78v-36h115v115zm0 148V599h37v115zm0 148V747h37v115zm0 148V895h37v115zm0 148v-115h37v115zm0 148v-115h37v115zm-595-35q-129 0-221-91.5T447 958q0-129 92-221t221-92q130 0 221.5 92t91.5 221q0 130-91.5 221.5T760 1271M595 890q0 36 19.5 56.5t49.5 25t64 7t64 2t49.5 9T861 1020q0 49-112 49q-97 0-123-51h-3l-31 63q67 42 162 42q29 0 56.5-5t55.5-16t45.5-33t17.5-53q0-46-27.5-69.5t-67.5-27t-79.5-3t-67-5T660 886q0-21 20.5-33t40.5-15t41-3q34 0 70.5 11t51.5 34h3l30-58q-3-1-21-8.5t-22.5-9t-19.5-7t-22-7t-20-4.5t-24-4t-23-1q-29 0-56.5 5t-54 16.5t-43 34T595 890"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBag(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1757 1408l35 313q3 28-16 50q-19 21-48 21H64q-29 0-48-21q-19-22-16-50l35-313zm-93-839l86 775H42l86-775q3-24 21-40.5t43-16.5h256v128q0 53 37.5 90.5T576 768t90.5-37.5T704 640V512h384v128q0 53 37.5 90.5T1216 768t90.5-37.5T1344 640V512h256q25 0 43 16.5t21 40.5m-384-185v256q0 26-19 45t-45 19t-45-19t-19-45V384q0-106-75-181t-181-75t-181 75t-75 181v256q0 26-19 45t-45 19t-45-19t-19-45V384q0-159 112.5-271.5T896 0t271.5 112.5T1280 384"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBasket(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1920 768q53 0 90.5 37.5T2048 896t-37.5 90.5t-90.5 37.5h-15l-115 662q-8 46-44 76t-82 30H384q-46 0-82-30t-44-76l-115-662h-15q-53 0-90.5-37.5T0 896t37.5-90.5T128 768zM485 1568q26-2 43.5-22.5T544 1499l-32-416q-2-26-22.5-43.5T443 1024t-43.5 22.5T384 1093l32 416q2 25 20.5 42t43.5 17zm411-64v-416q0-26-19-45t-45-19t-45 19t-19 45v416q0 26 19 45t45 19t45-19t19-45m384 0v-416q0-26-19-45t-45-19t-45 19t-19 45v416q0 26 19 45t45 19t45-19t19-45m352 5l32-416q2-26-15.5-46.5T1605 1024t-46.5 15.5t-22.5 43.5l-32 416q-2 26 15.5 46.5t43.5 22.5h5q25 0 43.5-17t20.5-42M476 292l-93 412H251l101-441q19-88 89-143.5T601 64h167q0-26 19-45t45-19h384q26 0 45 19t19 45h167q90 0 160 55.5t89 143.5l101 441h-132l-93-412q-11-44-45.5-72t-79.5-28h-167q0 26-19 45t-45 19H832q-26 0-45-19t-19-45H601q-45 0-79.5 28T476 292"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCart(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 1280q0 52-38 90t-90 38t-90-38t-38-90t38-90t90-38t90 38t38 90m896 0q0 52-38 90t-90 38t-90-38t-38-90t38-90t90-38t90 38t38 90m128-1088v512q0 24-16.5 42.5T1607 768L563 890q13 60 13 70q0 16-24 64h920q26 0 45 19t19 45t-19 45t-45 19H448q-26 0-45-19t-19-45q0-11 8-31.5t16-36t21.5-40T445 951L268 128H64q-26 0-45-19T0 64t19-45T64 0h256q16 0 28.5 6.5T368 22t13 24.5t8 26t5.5 29.5t4.5 26h1201q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shower(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1433 249q10 10 10 23t-10 23L807 921q-10 10-23 10t-23-10l-82-82q-10-10-10-23t10-23l44-44q-72-91-81.5-207T688 327q-74-71-176-71q-106 0-181 75t-75 181v1280H0V512q0-104 40.5-198.5T150 150T313.5 40.5T512 0q106 0 201 41t166 115q94-39 197-24.5t185 79.5l44-44q10-10 23-10t23 10zm-89 263q26 0 45 19t19 45t-19 45t-45 19t-45-19t-19-45t19-45t45-19m256 128q-26 0-45-19t-19-45t19-45t45-19t45 19t19 45t-19 45t-45 19m256-128q26 0 45 19t19 45t-19 45t-45 19t-45-19t-19-45t19-45t45-19m-640 128q26 0 45 19t19 45t-19 45t-45 19t-45-19t-19-45t19-45t45-19m192 64q0-26 19-45t45-19t45 19t19 45t-19 45t-45 19t-45-19t-19-45m320-64q26 0 45 19t19 45t-19 45t-45 19t-45-19t-19-45t19-45t45-19m-640 128q26 0 45 19t19 45t-19 45t-45 19t-45-19t-19-45t19-45t45-19m256 128q-26 0-45-19t-19-45t19-45t45-19t45 19t19 45t-19 45t-45 19m256-128q26 0 45 19t19 45t-19 45t-45 19t-45-19t-19-45t19-45t45-19m-384 256q-26 0-45-19t-19-45t19-45t45-19t45 19t19 45t-19 45t-45 19m256-128q26 0 45 19t19 45t-19 45t-45 19t-45-19t-19-45t19-45t45-19m-384 128q26 0 45 19t19 45t-19 45t-45 19t-45-19t-19-45t19-45t45-19m256 0q26 0 45 19t19 45t-19 45t-45 19t-45-19t-19-45t19-45t45-19m-128 128q26 0 45 19t19 45t-19 45t-45 19t-45-19t-19-45t19-45t45-19m-128 128q26 0 45 19t19 45t-19 45t-45 19t-45-19t-19-45t19-45t45-19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignIn(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1184 640q0 26-19 45l-544 544q-19 19-45 19t-45-19t-19-45V896H64q-26 0-45-19T0 832V448q0-26 19-45t45-19h448V96q0-26 19-45t45-19t45 19l544 544q19 19 19 45m352-352v704q0 119-84.5 203.5T1248 1280H928q-13 0-22.5-9.5T896 1248q0-4-1-20t-.5-26.5t3-23.5t10-19.5t20.5-6.5h320q66 0 113-47t47-113V288q0-66-47-113t-113-47H936l-11.5-1l-11.5-3l-8-5.5l-7-9l-2-13.5q0-4-1-20t-.5-26.5t3-23.5t10-19.5T928 0h320q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignLanguage(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M831 673q32 0 59 18l222 148q61 40 110 97l146 170q40 46 29 106l-72 413q-6 32-29.5 53.5T1240 1704l-527 56l-352 32h-9q-39 0-67.5-28t-28.5-68q0-37 27-64t65-32l260-32H160q-41 0-69.5-30T64 1467q2-39 32-65t69-26l442-1l-521-64q-41-5-66-37t-19-73q6-35 34.5-57.5T101 1121h10l481 60l-351-94q-38-10-62-41.5T161 977q6-36 33-58.5t62-22.5q6 0 20 2l448 96l217 37q1 0 3 .5t3 .5q23 0 30.5-23T965 973L779 848q-35-23-42-63.5t18-73.5q27-38 76-38m-70 202l186 125l-218-37l-5-2l-36-38l-238-262q-1-1-2.5-3.5T445 654q-24-31-18.5-70t37.5-64q31-23 68-17.5t64 33.5l142 147q-2 1-5 3.5t-4 4.5q-32 45-23 99t55 85m887-454l15 266q4 73-11 147l-48 219q-12 59-67 87l-106 54q2-62-39-109l-146-170q-53-61-117-103L907 664q-34-23-76-23q-51 0-88 37L508 366q-25-33-18-73.5t41-63.5q33-22 71.5-14t62.5 40l266 352l-262-455q-21-35-10.5-75T706 18q35-18 72.5-6T836 58l241 420l-136-337q-15-35-4.5-74T981 11q37-19 76-6t56 51l193 415l101 196q8 15 23 17.5t27-7.5t11-26l-12-224q-2-41 26-71t69-31q39 0 67 28.5t30 67.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignOut(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 1184q0 4 1 20t.5 26.5t-3 23.5t-10 19.5t-20.5 6.5H288q-119 0-203.5-84.5T0 992V288Q0 169 84.5 84.5T288 0h320q13 0 22.5 9.5T640 32q0 4 1 20t.5 26.5t-3 23.5t-10 19.5T608 128H288q-66 0-113 47t-47 113v704q0 66 47 113t113 47h312l11.5 1l11.5 3l8 5.5l7 9zm928-544q0 26-19 45l-544 544q-19 19-45 19t-45-19t-19-45V896H448q-26 0-45-19t-19-45V448q0-26 19-45t45-19h448V96q0-26 19-45t45-19t45 19l544 544q19 19 19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Signal(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 1312v192q0 14-9 23t-23 9H32q-14 0-23-9t-9-23v-192q0-14 9-23t23-9h192q14 0 23 9t9 23m384-128v320q0 14-9 23t-23 9H416q-14 0-23-9t-9-23v-320q0-14 9-23t23-9h192q14 0 23 9t9 23m384-256v576q0 14-9 23t-23 9H800q-14 0-23-9t-9-23V928q0-14 9-23t23-9h192q14 0 23 9t9 23m384-384v960q0 14-9 23t-23 9h-192q-14 0-23-9t-9-23V544q0-14 9-23t23-9h192q14 0 23 9t9 23m384-512v1472q0 14-9 23t-23 9h-192q-14 0-23-9t-9-23V32q0-14 9-23t23-9h192q14 0 23 9t9 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Simplybuilt(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M863 936q0-112-79.5-191.5T592 665t-191 79.5T322 936t79 191t191 79t191.5-79T863 936m863-1q0-112-79-191t-191-79t-191.5 79t-79.5 191q0 113 79.5 192t191.5 79t191-79.5t79-191.5m322-809v1348q0 44-31.5 75.5T1940 1581H108q-45 0-76.5-31.5T0 1474V126q0-44 31.5-75.5T108 19h431q44 0 76 31.5t32 75.5v161h754V126q0-44 32-75.5t76-31.5h431q45 0 76.5 31.5T2048 126"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sitemap(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 1120v320q0 40-28 68t-68 28h-320q-40 0-68-28t-28-68v-320q0-40 28-68t68-28h96V832H960v192h96q40 0 68 28t28 68v320q0 40-28 68t-68 28H736q-40 0-68-28t-28-68v-320q0-40 28-68t68-28h96V832H320v192h96q40 0 68 28t28 68v320q0 40-28 68t-68 28H96q-40 0-68-28t-28-68v-320q0-40 28-68t68-28h96V832q0-52 38-90t90-38h512V512h-96q-40 0-68-28t-28-68V96q0-40 28-68t68-28h320q40 0 68 28t28 68v320q0 40-28 68t-68 28h-96v192h512q52 0 90 38t38 90v192h96q40 0 68 28t28 68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Skyatlas(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1690 531q148 0 253 98.5T2048 874q0 157-109 261.5T1672 1240q-85 0-162-27.5t-138-73.5t-118-106t-109-126t-103.5-132.5T933 648T816 542t-136-73.5T521 441q-154 0-251.5 91.5T172 777q0 157 104 250t263 93q100 0 208-37.5T940 984q5-4 21-18.5t30-24t22-9.5q14 0 24.5 10.5T1048 967q0 24-60 77q-101 88-234.5 142T493 1240q-133 0-245.5-58t-180-165T0 776q0-205 141.5-341T489 299q120 0 226.5 43.5t185.5 113t151.5 153t139 167.5T1325 929.5t149.5 113T1647 1086q102 0 168.5-61.5T1882 862q0-95-64.5-159T1658 639q-30 0-81.5 18.5T1508 676q-20 0-35.5-15t-15.5-35q0-18 8.5-57t8.5-59q0-159-107.5-263T1100 143q-58 0-111.5 18.5t-84 40.5t-55.5 40.5t-33 18.5q-15 0-25.5-10.5T780 225q0-19 25-46q59-67 147-103.5T1134 39q191 0 318 125.5T1579 480q0 37-4 66q57-15 115-15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Skype(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1173 935q0-50-19.5-91.5T1105 775t-73-49t-82.5-34t-87.5-23l-104-24q-30-7-44-10.5T679 623t-30-16t-16.5-21t-7.5-30q0-77 144-77q43 0 77 12t54 28.5t38 33.5t40 29t48 12q47 0 75.5-32t28.5-77q0-55-56-99.5T932 318t-182-23q-68 0-132 15.5t-119.5 47t-89 87T376 573q0 61 19 106.5t56 75.5t80 48.5T634 836l146 36q90 22 112 36q32 20 32 60q0 39-40 64.5T779 1058q-51 0-91.5-16t-65-38.5t-45.5-45t-46-38.5t-54-16q-50 0-75.5 30t-25.5 75q0 92 122 157.5t291 65.5q73 0 140-18.5t122.5-53.5t88.5-93.5t33-131.5m363 217q0 159-112.5 271.5T1152 1536q-130 0-234-80q-77 16-150 16q-143 0-273.5-55.5t-225-150t-150-225T64 768q0-73 16-150Q0 514 0 384q0-159 112.5-271.5T384 0q130 0 234 80q77-16 150-16q143 0 273.5 55.5t225 150t150 225T1472 768q0 73-16 150q80 104 80 234"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Slack(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1519 776q62 0 103.5 40.5T1664 918q0 97-93 130l-172 59l56 167q7 21 7 47q0 59-42 102t-101 43q-47 0-85.5-27t-53.5-72l-55-165l-310 106l55 164q8 24 8 47q0 59-42 102t-102 43q-47 0-85-27t-53-72l-55-163l-153 53q-29 9-50 9q-61 0-101.5-40T196 1323q0-47 27.5-85t71.5-53l156-53l-105-313l-156 54q-26 8-48 8q-60 0-101-40.5T0 740q0-47 27.5-85T99 602l157-53l-53-159q-8-24-8-47q0-60 42-102.5T339 198q47 0 85 27t53 72l54 160l310-105l-54-160q-8-24-8-47q0-59 42.5-102T923 0q47 0 85.5 27.5T1062 99l53 161l162-55q21-6 43-6q60 0 102.5 39.5T1465 337q0 45-30 81.5t-74 51.5l-157 54l105 316l164-56q24-8 46-8m-794 262l310-105l-105-315l-310 107z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sliders(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M352 1152v128H0v-128zm352-128q26 0 45 19t19 45v256q0 26-19 45t-45 19H448q-26 0-45-19t-19-45v-256q0-26 19-45t45-19zm160-384v128H0V640zM224 128v128H0V128zm1312 1024v128H800v-128zM576 0q26 0 45 19t19 45v256q0 26-19 45t-45 19H320q-26 0-45-19t-19-45V64q0-26 19-45t45-19zm640 512q26 0 45 19t19 45v256q0 26-19 45t-45 19H960q-26 0-45-19t-19-45V576q0-26 19-45t45-19zm320 128v128h-224V640zm0-512v128H672V128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Slideshare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M873 740q0 83-63.5 142.5T657 942t-152.5-59.5T441 740q0-84 63.5-143T657 538t152.5 59T873 740m502 0q0 83-63 142.5T1159 942q-89 0-152.5-59.5T943 740q0-84 63.5-143t152.5-59q90 0 153 59t63 143m225 180V253q0-87-32-123.5T1457 93H345q-83 0-112.5 34T203 253v673q43 23 88.5 40t81 28t81 18.5t71 11t70 4t58.5.5t56.5-2t44.5-2q68-1 95 27q6 6 10 9q26 25 61 51q7-91 118-87q5 0 36.5 1.5t43 2t45.5 1t53-1t54.5-4.5t61-8.5t62-13.5t67-19.5t67.5-27t72-34.5m163-5q-121 149-372 252q84 285-23 465q-66 113-183 148q-104 32-182-15q-86-51-82-164l-1-326v-1q-8-2-24.5-6t-23.5-5l-1 338q4 114-83 164q-79 47-183 15q-117-36-182-150q-105-180-22-463q-251-103-372-252q-25-37-4-63t60 1q4 2 11.5 7t10.5 8V174q0-72 47-123T268 0h1257q67 0 114 51t47 123v694l21-15q39-27 60-1t-4 63"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmileO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1134 947q-37 121-138 195t-228 74t-228-74t-138-195q-8-25 4-48.5t38-31.5q25-8 48.5 4t31.5 38q25 80 92.5 129.5T768 1088t151.5-49.5T1012 909q8-26 32-38t49-4t37 31.5t4 48.5M640 512q0 53-37.5 90.5T512 640t-90.5-37.5T384 512t37.5-90.5T512 384t90.5 37.5T640 512m512 0q0 53-37.5 90.5T1024 640t-90.5-37.5T896 512t37.5-90.5T1024 384t90.5 37.5T1152 512m256 256q0-130-51-248.5t-136.5-204t-204-136.5T768 128t-248.5 51t-204 136.5t-136.5 204T128 768t51 248.5t136.5 204t204 136.5t248.5 51t248.5-51t204-136.5t136.5-204t51-248.5m128 0q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Snapchat(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1279 1020q0-22-22-27q-67-15-118-59t-80-108q-7-19-7-25q0-15 19.5-26t43-17t43-20.5T1177 701q0-19-18.5-31.5T1120 657q-12 0-32 8t-31 8q-4 0-12-2q5-95 5-114q0-79-17-114q-36-78-103-121.5T778 278q-199 0-275 165q-17 35-17 114q0 19 5 114q-4 2-14 2q-12 0-32-7.5t-30-7.5q-21 0-38.5 12T359 702q0 21 19.5 35.5t43 20.5t43 17t19.5 26q0 6-7 25q-64 138-198 167q-22 5-22 27q0 46 137 68q2 5 6 26t11.5 30.5t23.5 9.5q12 0 37.5-4.5t39.5-4.5q35 0 67 15t54 32.5t57.5 32.5t76.5 15q43 0 79-15t57.5-32.5T957 1160t67-15q14 0 39.5 4t38.5 4q16 0 23-10t11-30t6-25q137-22 137-68m257-252q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnapchatGhost(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M848 0q134-1 240.5 68.5T1252 261q27 58 27 179q0 47-9 191q14 7 28 7q18 0 51-13.5t51-13.5q29 0 56 18t27 46q0 32-31.5 54t-69 31.5t-69 29T1282 837q0 15 12 43q37 82 102.5 150t144.5 101q28 12 80 23q28 6 28 35q0 70-219 103q-7 11-11 39t-14 46.5t-33 18.5q-20 0-62-6.5t-64-6.5q-37 0-62 5q-32 5-63 22.5t-58 38t-58 40.5t-76 33.5t-99 13.5q-52 0-96.5-13.5t-75-33.5t-57.5-40.5t-58-38t-62-22.5q-26-5-63-5q-24 0-65.5 7.5T294 1398q-25 0-35-18.5t-14-47.5t-11-40q-219-33-219-103q0-29 28-35q52-11 80-23q78-32 144.5-101T370 880q12-28 12-43q0-28-31.5-47.5T281 760t-69.5-31.5T180 676q0-27 26-45.5t55-18.5q15 0 48 13t53 13q18 0 32-7q-9-142-9-190q0-122 27-180q64-137 172-198T848 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnapchatSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1280 1020q0-22-22-27q-67-14-118-58t-80-109q-7-14-7-25q0-15 19.5-26t42.5-17t42.5-20.5T1177 701q0-19-18.5-31.5T1120 657q-11 0-31 8t-32 8q-4 0-12-2q5-63 5-115q0-78-17-114q-36-78-102.5-121.5T778 277q-198 0-275 165q-18 38-18 115q0 38 6 114q-10 2-15 2q-11 0-31.5-8t-30.5-8q-20 0-37.5 12.5T359 702q0 21 19.5 35.5T421 758t42.5 17t19.5 26q0 11-7 25q-64 138-198 167q-22 5-22 27q0 47 138 69q2 5 6 26t11 30.5t23 9.5q13 0 38.5-5t38.5-5q35 0 67.5 15t54.5 32.5t57.5 32.5t76.5 15q43 0 79-15t57.5-32.5t54-32.5t67.5-15q13 0 39 4.5t39 4.5q15 0 22.5-9.5t11.5-31t5-24.5q138-22 138-69m256-732v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnowflakeO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1534 1117l-167 33l186 107q23 13 29.5 38.5t-6.5 48.5q-14 23-39 29.5t-48-6.5l-186-106l55 160q13 38-12 63.5t-60.5 20.5t-48.5-42l-102-300l-271-156v313l208 238q16 18 17 39t-11 36.5t-28.5 25t-37 5.5t-36.5-22l-112-128v214q0 26-19 45t-45 19t-45-19t-19-45v-214l-112 128q-16 18-36.5 22t-37-5.5t-28.5-25t-11-36.5t17-39l208-238v-313l-271 156l-102 300q-13 37-48.5 42t-60.5-20.5t-12-63.5l55-160l-186 106q-23 13-48 6.5T24 1344q-13-23-6.5-48.5T47 1257l186-107l-167-33q-29-6-42-29t-8.5-46.5t25.5-40T91 991l310 62l271-157l-271-157l-310 62q-4 1-13 1q-27 0-44-18t-19-40t11-43t40-26l167-33L47 535q-23-13-29.5-38.5T24 448t39-30t48 7l186 106l-55-160q-13-38 12-63.5t60.5-20.5t48.5 42l102 300l271 156V472L528 234q-16-18-17-39t11-36.5t28.5-25t37-5.5t36.5 22l112 128V64q0-26 19-45t45-19t45 19t19 45v214l112-128q16-18 36.5-22t37 5.5t28.5 25t11 36.5t-17 39L864 472v313l271-156l102-300q13-37 48.5-42t60.5 20.5t12 63.5l-55 160l186-106q23-13 48-6.5t39 29.5q13 23 6.5 48.5T1553 535l-186 107l167 33q27 5 40 26t11 43t-19 40t-44 18q-9 0-13-1l-310-62l-271 157l271 157l310-62q29-6 50 10.5t25.5 40t-8.5 46.5t-42 29"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sort(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1024 896q0 26-19 45l-448 448q-19 19-45 19t-45-19L19 941Q0 922 0 896t19-45t45-19h896q26 0 45 19t19 45m0-384q0 26-19 45t-45 19H64q-26 0-45-19T0 512t19-45L467 19q19-19 45-19t45 19l448 448q19 19 19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAlphaAsc(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1159 408h177l-72-218l-12-47q-2-16-2-20h-4l-3 20q0 1-3.5 18t-7.5 29zM704 1440q0 12-10 24l-319 319q-10 9-23 9q-12 0-23-9L9 1463q-15-16-7-35q8-20 30-20h192V32q0-14 9-23t23-9h192q14 0 23 9t9 23v1376h192q14 0 23 9t9 23m836 119v233H956v-90l369-529q12-18 21-27l11-9v-3q-2 0-6.5.5t-7.5.5q-12 3-30 3h-232v115H961v-229h567v89l-369 530q-6 8-21 26l-11 11v2l14-2q9-2 30-2h248v-119zm89-897v106h-288V662h75l-47-144h-243l-47 144h75v106H867V662h70L1167 0h162l230 662z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAlphaDesc(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1159 1432h177l-72-218l-12-47q-2-16-2-20h-4l-3 20q0 1-3.5 18t-7.5 29zm-455 8q0 12-10 24l-319 319q-10 9-23 9q-12 0-23-9L9 1463q-15-16-7-35q8-20 30-20h192V32q0-14 9-23t23-9h192q14 0 23 9t9 23v1376h192q14 0 23 9t9 23m925 246v106h-288v-106h75l-47-144h-243l-47 144h75v106H867v-106h70l230-662h162l230 662zm-89-1151v233H956v-90l369-529q12-18 21-27l11-9v-3q-2 0-6.5.5t-7.5.5q-12 3-30 3h-232v115H961V0h567v89l-369 530q-6 8-21 26l-11 10v3l14-3q9-1 30-1h248V535z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAmountAsc(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M704 1440q0 12-10 24l-319 319q-10 9-23 9q-12 0-23-9L9 1463q-15-16-7-35q8-20 30-20h192V32q0-14 9-23t23-9h192q14 0 23 9t9 23v1376h192q14 0 23 9t9 23m1056 128v192q0 14-9 23t-23 9H896q-14 0-23-9t-9-23v-192q0-14 9-23t23-9h832q14 0 23 9t9 23m-192-512v192q0 14-9 23t-23 9H896q-14 0-23-9t-9-23v-192q0-14 9-23t23-9h640q14 0 23 9t9 23m-192-512v192q0 14-9 23t-23 9H896q-14 0-23-9t-9-23V544q0-14 9-23t23-9h448q14 0 23 9t9 23M1184 32v192q0 14-9 23t-23 9H896q-14 0-23-9t-9-23V32q0-14 9-23t23-9h256q14 0 23 9t9 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAmountDesc(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1184 1568v192q0 14-9 23t-23 9H896q-14 0-23-9t-9-23v-192q0-14 9-23t23-9h256q14 0 23 9t9 23m-480-128q0 12-10 24l-319 319q-10 9-23 9q-12 0-23-9L9 1463q-15-16-7-35q8-20 30-20h192V32q0-14 9-23t23-9h192q14 0 23 9t9 23v1376h192q14 0 23 9t9 23m672-384v192q0 14-9 23t-23 9H896q-14 0-23-9t-9-23v-192q0-14 9-23t23-9h448q14 0 23 9t9 23m192-512v192q0 14-9 23t-23 9H896q-14 0-23-9t-9-23V544q0-14 9-23t23-9h640q14 0 23 9t9 23m192-512v192q0 14-9 23t-23 9H896q-14 0-23-9t-9-23V32q0-14 9-23t23-9h832q14 0 23 9t9 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAsc(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1024 512q0 26-19 45t-45 19H64q-26 0-45-19T0 512t19-45L467 19q19-19 45-19t45 19l448 448q19 19 19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortNumericAsc(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1314 1313q0-63-44-116t-103-53q-52 0-83 37t-31 94t36.5 95t104.5 38q50 0 85-27t35-68m-610 127q0 12-10 24l-319 319q-10 9-23 9q-12 0-23-9L9 1463q-15-16-7-35q8-20 30-20h192V32q0-14 9-23t23-9h192q14 0 23 9t9 23v1376h192q14 0 23 9t9 23m750-69q0 62-13 121.5t-41 114t-68 95.5t-98.5 65.5T1106 1792q-62 0-108-16q-24-8-42-15l39-113q15 7 31 11q37 13 75 13q84 0 134.5-58.5T1302 1468h-2q-21 23-61.5 37t-84.5 14q-106 0-173-71.5T914 1275q0-105 72-178t181-73q123 0 205 94.5t82 252.5m-30-717v114H955V654h167V222q0-7 .5-19t.5-17v-16h-2l-7 12q-8 13-26 31l-62 58l-82-86L1136 0h123v654z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortNumericDesc(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1314 289q0-63-44-116t-103-53q-52 0-83 37t-31 94t36.5 95t104.5 38q50 0 85-27t35-68M704 1440q0 12-10 24l-319 319q-10 9-23 9q-12 0-23-9L9 1463q-15-16-7-35q8-20 30-20h192V32q0-14 9-23t23-9h192q14 0 23 9t9 23v1376h192q14 0 23 9t9 23m720 238v114H955v-114h167v-432q0-7 .5-19t.5-17v-16h-2l-7 12q-8 13-26 31l-62 58l-82-86l192-185h123v654zm30-1331q0 62-13 121.5t-41 114t-68 95.5t-98.5 65.5T1106 768q-62 0-108-16q-24-8-42-15l39-113q15 7 31 11q37 13 75 13q84 0 134.5-58.5T1302 444h-2q-21 23-61.5 37t-84.5 14q-106 0-173-71.5T914 251q0-105 72-178t181-73q123 0 205 94.5t82 252.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Soundcloud(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m784 1116l16-241l-16-523q-1-10-7.5-17t-16.5-7q-9 0-16 7t-7 17l-14 523l14 241q1 10 7.5 16.5t15.5 6.5q22 0 24-23m296-29l11-211l-12-586q0-16-13-24q-8-5-16-5t-16 5q-13 8-13 24l-1 6l-10 579q0 1 11 236v1q0 10 6 17q9 11 23 11q11 0 20-9q9-7 9-20zM35 747l20 128l-20 126q-2 9-9 9t-9-9L0 875l17-128q2-9 9-9t9 9m86-79l26 207l-26 203q-2 9-10 9q-9 0-9-10L79 875l23-207q0-9 9-9q8 0 10 9m92-38l25 245l-25 237q0 11-11 11q-10 0-12-11l-21-237l21-245q2-12 12-12q11 0 11 12m94-7l23 252l-23 244q-2 13-14 13q-13 0-13-13l-21-244l21-252q0-13 13-13q12 0 14 13m94 18l21 234l-21 246q-2 16-16 16q-6 0-10.5-4.5T370 1121l-20-246l20-234q0-6 4.5-10.5T385 626q14 0 16 15m94-146l21 380l-21 246q0 7-5 12.5t-12 5.5q-16 0-18-18l-18-246l18-380q2-18 18-18q7 0 12 5.5t5 12.5m94-86l19 468l-19 244q0 8-5.5 13.5T570 1140q-18 0-20-19l-16-244l16-468q2-19 20-19q8 0 13.5 5.5T589 409m98-40l18 506l-18 242q-2 21-22 21q-19 0-21-21l-16-242l16-506q0-9 6.5-15.5T665 347q9 0 15 6.5t7 15.5m194-4l15 510l-15 239q0 10-7.5 17.5T856 1139t-17-7t-8-18l-14-239l14-510q0-11 7.5-18t17.5-7t17.5 7t7.5 18m99 19l14 492l-14 236q0 11-8 19t-19 8t-19-8t-9-19l-12-236l12-492q1-12 9-20t19-8t18.5 8t8.5 20m212 492l-14 231q0 13-9 22t-22 9t-22-9t-10-22l-6-114l-6-117l12-636v-3q2-15 12-24q9-7 20-7q8 0 15 5q14 8 16 26zm1112-19q0 117-83 199.5t-200 82.5h-786q-13-2-22-11t-9-22V207q0-23 28-33q85-34 181-34q195 0 338 131.5T1911 595q53-22 110-22q117 0 200 83t83 201"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpaceShuttle(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M620 864q-110 64-268 64H224v-64h-64q-13 0-22.5-23.5T128 784q0-24 7-49q-58-2-96.5-10.5T0 704t38.5-20.5T135 673q-7-25-7-49q0-33 9.5-56.5T160 544h64v-64h128q158 0 268 64h1113q42 7 106.5 18t80.5 14q89 15 150 40.5t83.5 47.5t22.5 40t-22.5 40t-83.5 47.5t-150 40.5q-16 3-80.5 14T1733 864zm1119-252q53 36 53 92t-53 92l81 30q68-48 68-122t-68-122zM625 880h1015q-217 38-456 80q-57 0-113 24t-83 48l-28 24l-288 288q-26 26-70.5 45t-89.5 19h-96l-93-464h29q157 0 273-64M352 464h-29L416 0h96q46 0 90 19t70 45l288 288q4 4 11 10.5t30.5 23t48.5 29t61.5 23T1184 448l456 80H625q-116-64-273-64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spinner(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M462 1394q0 53-37.5 90.5T334 1522q-52 0-90-38t-38-90q0-53 37.5-90.5T334 1266t90.5 37.5T462 1394m498 206q0 53-37.5 90.5T832 1728t-90.5-37.5T704 1600t37.5-90.5T832 1472t90.5 37.5T960 1600M256 896q0 53-37.5 90.5T128 1024t-90.5-37.5T0 896t37.5-90.5T128 768t90.5 37.5T256 896m1202 498q0 52-38 90t-90 38q-53 0-90.5-37.5T1202 1394t37.5-90.5t90.5-37.5t90.5 37.5t37.5 90.5M494 398q0 66-47 113t-113 47t-113-47t-47-113t47-113t113-47t113 47t47 113m1170 498q0 53-37.5 90.5T1536 1024t-90.5-37.5T1408 896t37.5-90.5T1536 768t90.5 37.5T1664 896m-640-704q0 80-56 136t-136 56t-136-56t-56-136t56-136T832 0t136 56t56 136m530 206q0 93-66 158.5T1330 622q-93 0-158.5-65.5T1106 398q0-92 65.5-158t158.5-66q92 0 158 66t66 158"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spoon(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 528q0 145-57 243.5T431 907l45 821q2 26-16 45t-44 19H224q-26 0-44-19t-16-45l45-821q-95-37-152-135.5T0 528q0-128 42.5-249.5T160 78.5T320 0t160 78.5t117.5 200T640 528"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spotify(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1127 1082q0-32-30-51q-193-115-447-115q-133 0-287 34q-42 9-42 52q0 20 13.5 34.5T370 1051q5 0 37-8q132-27 243-27q226 0 397 103q19 11 33 11q19 0 33-13.5t14-34.5m96-215q0-40-35-61q-237-141-548-141q-153 0-303 42q-48 13-48 64q0 25 17.5 42.5T349 831q7 0 37-8q122-33 251-33q279 0 488 124q24 13 38 13q25 0 42.5-17.5T1223 867m108-248q0-47-40-70q-126-73-293-110.5T655 401q-204 0-364 47q-23 7-38.5 25.5T237 522q0 31 20.5 52t51.5 21q11 0 40-8q133-37 307-37q159 0 309.5 34t253.5 95q21 12 40 12q29 0 50.5-20.5T1331 619m205 149q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Square(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1536 288v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1120 128H288q-66 0-113 47t-47 113v832q0 66 47 113t113 47h832q66 0 113-47t47-113V288q0-66-47-113t-113-47m288 160v832q0 119-84.5 203.5T1120 1408H288q-119 0-203.5-84.5T0 1120V288Q0 169 84.5 84.5T288 0h832q119 0 203.5 84.5T1408 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StackExchange(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1259 997v66q0 85-57.5 144.5T1063 1267h-57l-260 269v-269H217q-81 0-138.5-59.5T21 1063v-66zm0-326v255H21V671zm0-328v255H21V343zm0-140v67H21v-67q0-84 57.5-143.5T217 0h846q81 0 138.5 59.5T1259 203"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StackOverflow(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1289 1632H171v-480H11v640h1438v-640h-160zm-942-524l33-157l783 165l-33 156zm103-374l67-146l725 339l-67 145zm201-356l102-123l614 513l-102 123zM1048 0l477 641l-128 96L920 96zM330 1471v-159h800v159z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1664 615q0 22-26 48l-363 354l86 500q1 7 1 20q0 21-10.5 35.5T1321 1587q-19 0-40-12l-449-236l-449 236q-22 12-40 12q-21 0-31.5-14.5T301 1537q0-6 2-20l86-500L25 663Q0 636 0 615q0-37 56-46l502-73L783 41q19-41 49-41t49 41l225 455l502 73q56 9 56 46"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarHalf(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 0v1339l-449 236q-22 12-40 12q-21 0-31.5-14.5T301 1537q0-6 2-20l86-500L25 663Q0 636 0 615q0-37 56-46l502-73L783 41q19-41 49-41"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarHalfEmpty(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1186 925l257-250l-356-52l-66-10l-30-60l-159-322v963l59 31l318 168l-60-355l-12-66zm452-262l-363 354l86 500q5 33-6 51.5t-34 18.5q-17 0-40-12l-449-236l-449 236q-23 12-40 12q-23 0-34-18.5t-6-51.5l86-500L25 663q-32-32-23-59.5T56 569l502-73L783 41q20-41 49-41q28 0 49 41l225 455l502 73q45 7 54 34.5t-24 59.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1137 972l306-297l-422-62l-189-382l-189 382l-422 62l306 297l-73 421l378-199l377 199zm527-357q0 22-26 48l-363 354l86 500q1 7 1 20q0 50-41 50q-19 0-40-12l-449-236l-449 236q-22 12-40 12q-21 0-31.5-14.5T301 1537q0-6 2-20l86-500L25 663Q0 636 0 615q0-37 56-46l502-73L783 41q19-41 49-41t49 41l225 455l502 73q56 9 56 46"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Steam(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1582 454q0 101-71.5 172.5T1338 698t-172.5-71.5T1094 454t71.5-172.5T1338 210t172.5 71.5T1582 454m-770 742q0-104-73-177t-177-73q-27 0-54 6l104 42q77 31 109.5 106.5T723 1252q-31 77-107 109t-152 1q-21-8-62-24.5t-61-24.5q32 60 91 96.5t130 36.5q104 0 177-73t73-177m830-741q0-126-89.5-215.5T1337 150q-127 0-216.5 89.5T1031 455q0 127 89.5 216t216.5 89q126 0 215.5-89t89.5-216m150 0q0 189-133.5 322T1337 910l-437 319q-12 129-109 218t-229 89q-121 0-214-76t-118-192L0 1176V747l389 157q79-48 173-48q13 0 35 2l284-407q2-187 135.5-319T1337 0q188 0 321.5 133.5T1792 455"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SteamSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1242 519q0-80-57-136.5T1048 326t-136.5 57T855 519q0 80 56.5 136.5T1048 712t137-56.5t57-136.5m-610 588q0 83-58 140.5T434 1305q-56 0-103-29t-72-77q52 20 98 40q60 24 120-1.5t85-86.5q24-60-1.5-120T474 947l-82-33q22-5 42-5q82 0 140 57.5t58 140.5m904-819v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248v-153l172 69q20 92 93.5 152t168.5 60q104 0 181-70t87-173l345-252q150 0 255.5-105.5T1408 521q0-150-105.5-255.5T1047 160q-148 0-253 104.5T687 517L462 839q-9-1-28-1q-75 0-137 37L0 756V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288m-247 233q0 100-71 170.5T1047 762t-170.5-70.5T806 521t70.5-171t170.5-71q101 0 171.5 70.5T1289 521"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StepBackward(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M979 13q19-19 32-13t13 32v1472q0 26-13 32t-32-13L269 813q-9-9-13-19v678q0 26-19 45t-45 19H64q-26 0-45-19t-19-45V64q0-26 19-45T64 0h128q26 0 45 19t19 45v678q4-10 13-19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stethoscope(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1280 576q0-26-19-45t-45-19t-45 19t-19 45t19 45t45 19t45-19t19-45m128 0q0 62-35.5 111t-92.5 70v395q0 159-131.5 271.5T832 1536t-316.5-112.5T384 1152v-132q-164-20-274-128T0 640V128q0-26 19-45t45-19q6 0 16 2q17-30 47-48t65-18q53 0 90.5 37.5T320 128t-37.5 90.5T192 256q-33 0-64-18v402q0 106 94 181t226 75t226-75t94-181V238q-31 18-64 18q-53 0-90.5-37.5T576 128t37.5-90.5T704 0q35 0 65 18t47 48q10-2 16-2q26 0 45 19t19 45v512q0 144-110 252t-274 128v132q0 106 94 181t226 75t226-75t94-181V757q-57-21-92.5-70T1024 576q0-80 56-136t136-56t136 56t56 136"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StickyNote(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1024 1120v416H96q-40 0-68-28t-28-68V96q0-40 28-68T96 0h1344q40 0 68 28t28 68v928h-416q-40 0-68 28t-28 68m128 32h381q-15 82-65 132l-184 184q-50 50-132 65z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StickyNoteO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1400 1152h-248v248q29-10 41-22l185-185q12-12 22-41m-280-128h288V128H128v1280h896v-288q0-40 28-68t68-28m416-928v1024q0 40-20 88t-48 76l-184 184q-28 28-76 48t-88 20H96q-40 0-68-28t-28-68V96q0-40 28-68T96 0h1344q40 0 68 28t28 68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stop(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1536 64v1408q0 26-19 45t-45 19H64q-26 0-45-19t-19-45V64q0-26 19-45T64 0h1408q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopCircle(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1088 1056V480q0-14-9-23t-23-9H480q-14 0-23 9t-9 23v576q0 14 9 23t23 9h576q14 0 23-9t9-23m448-288q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopCircleO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 0q209 0 385.5 103T1433 382.5T1536 768t-103 385.5t-279.5 279.5T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0m0 1312q148 0 273-73t198-198t73-273t-73-273t-198-198t-273-73t-273 73t-198 198t-73 273t73 273t198 198t273 73m-288-224q-14 0-23-9t-9-23V480q0-14 9-23t23-9h576q14 0 23 9t9 23v576q0 14-9 23t-23 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StreetView(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1408 1504q0 63-61.5 113.5t-164 81t-225 46T704 1760t-253.5-15.5t-225-46t-164-81T0 1504q0-49 33-88.5t91-66.5t118-44.5t131-29.5q26-5 48 10.5t26 41.5q5 26-10.5 48t-41.5 26q-58 10-106 23.5t-76.5 25.5t-48.5 23.5t-27.5 19.5t-8.5 12q3 11 27 26.5t73 33t114 32.5t160.5 25t201.5 10t201.5-10t160.5-25t114-33t73-33.5t27-27.5q-1-4-8.5-11t-27.5-19t-48.5-23.5t-76.5-25t-106-23.5q-26-4-41.5-26t-10.5-48q4-26 26-41.5t48-10.5q71 12 131 29.5t118 44.5t91 66.5t33 88.5m-384-896v384q0 26-19 45t-45 19h-64v384q0 26-19 45t-45 19H576q-26 0-45-19t-19-45v-384h-64q-26 0-45-19t-19-45V608q0-53 37.5-90.5T512 480h384q53 0 90.5 37.5T1024 608m-96-384q0 93-65.5 158.5T704 448t-158.5-65.5T480 224t65.5-158.5T704 0t158.5 65.5T928 224"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Strikethrough(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1760 768q14 0 23 9t9 23v64q0 14-9 23t-23 9H32q-14 0-23-9t-9-23v-64q0-14 9-23t23-9zM483 704q-28-35-51-80q-48-98-48-188q0-181 134-309Q651 0 911 0q50 0 167 19q66 12 177 48q10 38 21 118q14 123 14 183q0 18-5 45l-12 3l-84-6l-14-2q-50-149-103-205q-88-91-210-91q-114 0-182 59q-67 58-67 146q0 73 66 140t279 129q69 20 173 66q58 28 95 52zm507 256h411q7 39 7 92q0 111-41 212q-23 56-71 104q-37 35-109 81q-80 48-153 66q-80 21-203 21q-114 0-195-23l-140-40q-57-16-72-28q-8-8-8-22v-13q0-108-2-156q-1-30 0-68l2-37v-44l102-2q15 34 30 71t22.5 56t12.5 27q35 57 80 94q43 36 105 57q59 22 132 22q64 0 139-27q77-26 122-86q47-61 47-129q0-84-81-157q-34-29-137-71"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stumbleupon(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1062 552V434q0-42-30-72t-72-30t-72 30t-30 72v612q0 175-126 299t-303 124q-178 0-303.5-125.5T0 1040V774h328v262q0 43 30 72.5t72 29.5t72-29.5t30-72.5V416q0-171 126.5-292T960 3q176 0 302 122t126 294v136l-195 58zm530 222h328v266q0 178-125.5 303.5T1491 1469q-177 0-303-124.5T1062 1044V776l131 61l195-58v270q0 42 30 71.5t72 29.5t72-29.5t30-71.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StumbleuponCircle(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m866 711l90-27v-62q0-79-58-135t-138-56t-138 55.5T564 621v283q0 20-14 33.5T517 951t-32.5-13.5T471 904V784H320v122q0 82 57.5 139t139.5 57q81 0 138.5-56.5T713 909V629q0-19 13.5-33t33.5-14q19 0 32.5 14t13.5 33v54zm333 195V784h-150v126q0 20-13.5 33.5T1002 957q-19 0-32.5-14T956 910V787l-90 26l-60-28v123q0 80 58 137t139 57t138.5-57t57.5-139m337-138q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Subscript(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M897 1113v167H649l-159-252l-24-42q-8-9-11-21h-3q-1 3-2.5 6.5t-3.5 8t-3 6.5q-10 20-25 44l-155 250H5v-167h128l197-291l-185-272H8V382h276l139 228q2 4 23 42q8 9 11 21h3q3-9 11-21l25-42l140-228h257v168H768L584 817l204 296zm639 217v206h-514l-4-27q-3-45-3-46q0-64 26-117t65-86.5t84-65t84-54.5t65-54t26-64q0-38-29.5-62.5T1265 935q-51 0-97 39q-14 11-36 38l-105-92q26-37 63-66q80-65 188-65q110 0 178 59.5t68 158.5q0 66-34.5 118.5t-84 86t-99.5 62.5t-87 63t-41 73h232v-80z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Subway(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1088 0q185 0 316.5 93.5T1536 320v896q0 130-125.5 222t-305.5 97l213 202q16 15 8 35t-30 20H240q-22 0-30-20t8-35l213-202q-180-5-305.5-97T0 1216V320Q0 187 131.5 93.5T448 0zM288 1312q66 0 113-47t47-113t-47-113t-113-47t-113 47t-47 113t47 113t113 47m416-544V256H160v512zm544 544q66 0 113-47t47-113t-47-113t-113-47t-113 47t-47 113t47 113t113 47m160-544V256H832v512z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Suitcase(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 256h512V128H640zm-352 0v1280h-64q-92 0-158-66T0 1312V480q0-92 66-158t158-66zm1120 0v1280H384V256h128V96q0-40 28-68t68-28h576q40 0 68 28t28 68v160zm384 224v832q0 92-66 158t-158 66h-64V256h64q92 0 158 66t66 158"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1440 896q0-117-45.5-223.5t-123-184t-184-123T864 320t-223.5 45.5t-184 123t-123 184T288 896t45.5 223.5t123 184t184 123T864 1472t223.5-45.5t184-123t123-184T1440 896m276 277q-4 15-20 20l-292 96v306q0 16-13 26q-15 10-29 4l-292-94l-180 248q-10 13-26 13t-26-13l-180-248l-292 94q-14 6-29-4q-13-10-13-26v-306l-292-96q-16-5-20-20q-5-17 4-29l180-248L16 648q-9-13-4-29q4-15 20-20l292-96V197q0-16 13-26q15-10 29-4l292 94L838 13q9-12 26-12t26 12l180 248l292-94q14-6 29 4q13 10 13 26v306l292 96q16 5 20 20q5 16-4 29l-180 248l180 248q9 12 4 29"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Superpowers(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1473 929q7-118-33-226.5t-113-189t-177-131T929 325q-116-7-225.5 32t-192 110.5t-135 175T317 863q-7 118 33 226.5t113 189t177.5 131T862 1467q155 9 293-59t224-195.5t94-283.5M1792 0l-349 348q120 117 180.5 272t50.5 321q-11 183-102 339t-241 255.5T999 1660L0 1792l347-347q-120-116-180.5-271.5T116 852q11-184 102-340t241.5-255.5T792 132q167-22 500-66t500-66"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Superscript(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M897 1241v167H649l-159-252l-24-42q-8-9-11-21h-3q-1 3-2.5 6.5t-3.5 8t-3 6.5q-10 20-25 44l-155 250H5v-167h128l197-291l-185-272H8V510h276l139 228q2 4 23 42q8 9 11 21h3q3-9 11-21l25-42l140-228h257v168H768L584 945l204 296zm637-679v206h-514l-3-27q-4-28-4-46q0-64 26-117t65-86.5t84-65t84-54.5t65-54t26-64q0-38-29.5-62.5T1263 167q-51 0-97 39q-14 11-36 38l-105-92q26-37 63-66q83-65 188-65q110 0 178 59.5t68 158.5q0 56-24.5 103t-62 76.5T1354 477t-82 50.5t-65.5 51.5t-30.5 63h232v-80z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Table(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 1248v-192q0-14-9-23t-23-9H160q-14 0-23 9t-9 23v192q0 14 9 23t23 9h320q14 0 23-9t9-23m0-384V672q0-14-9-23t-23-9H160q-14 0-23 9t-9 23v192q0 14 9 23t23 9h320q14 0 23-9t9-23m512 384v-192q0-14-9-23t-23-9H672q-14 0-23 9t-9 23v192q0 14 9 23t23 9h320q14 0 23-9t9-23M512 480V288q0-14-9-23t-23-9H160q-14 0-23 9t-9 23v192q0 14 9 23t23 9h320q14 0 23-9t9-23m512 384V672q0-14-9-23t-23-9H672q-14 0-23 9t-9 23v192q0 14 9 23t23 9h320q14 0 23-9t9-23m512 384v-192q0-14-9-23t-23-9h-320q-14 0-23 9t-9 23v192q0 14 9 23t23 9h320q14 0 23-9t9-23m-512-768V288q0-14-9-23t-23-9H672q-14 0-23 9t-9 23v192q0 14 9 23t23 9h320q14 0 23-9t9-23m512 384V672q0-14-9-23t-23-9h-320q-14 0-23 9t-9 23v192q0 14 9 23t23 9h320q14 0 23-9t9-23m0-384V288q0-14-9-23t-23-9h-320q-14 0-23 9t-9 23v192q0 14 9 23t23 9h320q14 0 23-9t9-23m128-320v1088q0 66-47 113t-113 47H160q-66 0-113-47T0 1248V160Q0 94 47 47T160 0h1344q66 0 113 47t47 113"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tablet(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 1280q0-26-19-45t-45-19t-45 19t-19 45t19 45t45 19t45-19t19-45m384-160V160q0-13-9.5-22.5T992 128H160q-13 0-22.5 9.5T128 160v960q0 13 9.5 22.5t22.5 9.5h832q13 0 22.5-9.5t9.5-22.5m128-960v1088q0 66-47 113t-113 47H160q-66 0-113-47T0 1248V160Q0 94 47 47T160 0h832q66 0 113 47t47 113"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 320q0-53-37.5-90.5T320 192t-90.5 37.5T192 320t37.5 90.5T320 448t90.5-37.5T448 320m1067 576q0 53-37 90l-491 492q-39 37-91 37q-53 0-90-37L91 762q-38-37-64.5-101T0 544V128q0-52 38-90t90-38h416q53 0 117 26.5T763 91l715 714q37 39 37 91"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tags(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 320q0-53-37.5-90.5T320 192t-90.5 37.5T192 320t37.5 90.5T320 448t90.5-37.5T448 320m1067 576q0 53-37 90l-491 492q-39 37-91 37q-53 0-90-37L91 762q-38-37-64.5-101T0 544V128q0-52 38-90t90-38h416q53 0 117 26.5T763 91l715 714q37 39 37 91m384 0q0 53-37 90l-491 492q-39 37-91 37q-36 0-59-14t-53-45l470-470q37-37 37-90q0-52-37-91L923 91q-38-38-102-64.5T704 0h224q53 0 117 26.5T1147 91l715 714q37 39 37 91"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tasks(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1024 1280h640v-128h-640zM640 768h1024V640H640zm640-512h384V128h-384zm512 832v256q0 26-19 45t-45 19H64q-26 0-45-19t-19-45v-256q0-26 19-45t45-19h1664q26 0 45 19t19 45m0-512v256q0 26-19 45t-45 19H64q-26 0-45-19T0 832V576q0-26 19-45t45-19h1664q26 0 45 19t19 45m0-512v256q0 26-19 45t-45 19H64q-26 0-45-19T0 320V64q0-26 19-45T64 0h1664q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Telegram(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1189 1307l147-693q9-44-10.5-63t-51.5-7L410 877q-29 11-39.5 25t-2.5 26.5t32 19.5l221 69l513-323q21-14 32-6q7 5-4 15l-415 375l-16 228q23 0 45-22l108-104l224 165q64 36 81-38m603-411q0 182-71 348t-191 286t-286 191t-348 71t-348-71t-286-191t-191-286T0 896t71-348t191-286T548 71T896 0t348 71t286 191t191 286t71 348"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Television(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 1120V160q0-13-9.5-22.5T1760 128H160q-13 0-22.5 9.5T128 160v960q0 13 9.5 22.5t22.5 9.5h1600q13 0 22.5-9.5t9.5-22.5m128-960v960q0 66-47 113t-113 47h-736v128h352q14 0 23 9t9 23v64q0 14-9 23t-23 9H544q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h352v-128H160q-66 0-113-47T0 1120V160Q0 94 47 47T160 0h1600q66 0 113 47t47 113"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TencentWeibo(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M810 572q0 80-57 136.5T617 765q-60 0-111-35q-62 67-115 146q-247 371-202 859q1 22-12.5 38.5T142 1792h-5q-20 0-35-13.5T85 1745q-14-126-3.5-247.5t29.5-217t54-186T234 939t74-125q61-90 132-165q-16-35-16-77q0-80 56.5-136.5T617 379t136.5 56.5T810 572m381 11q0 158-78 292t-212.5 212t-292.5 78q-64 0-131-14q-21-5-32.5-23.5T438 1088q5-20 23-31.5t39-7.5q51 13 108 13q97 0 186-38t153-102t102-153t38-186t-38-186t-102-153t-153-102t-186-38t-186 38t-153 102t-102 153t-38 186q0 114 52 218q10 20 3.5 40T159 871t-39.5 3T89 848Q25 725 25 583q0-119 46.5-227T196 170T382 46T608 0q158 0 292.5 78T1113 290.5t78 292.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Terminal(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m585 727l-466 466q-10 10-23 10t-23-10l-50-50q-10-10-10-23t10-23l393-393L23 311q-10-10-10-23t10-23l50-50q10-10 23-10t23 10l466 466q10 10 10 23t-10 23m1079 457v64q0 14-9 23t-23 9H672q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h960q14 0 23 9t9 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextHeight(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1744 1280q33 0 42 18.5t-11 44.5l-126 162q-20 26-49 26t-49-26l-126-162q-20-26-11-44.5t42-18.5h80V256h-80q-33 0-42-18.5t11-44.5l126-162q20-26 49-26t49 26l126 162q20 26 11 44.5t-42 18.5h-80v1024zM81 1l54 27q12 5 211 5q44 0 132-2t132-2q36 0 107.5.5T825 30h293q6 0 21 .5t20.5 0t16-3t17.5-9T1208 1l42-1q4 0 14 .5t14 .5q2 112 2 336q0 80-5 109q-39 14-68 18q-25-44-54-128q-3-9-11-48t-14.5-73.5t-7.5-35.5q-6-8-12-12.5t-15.5-6t-13-2.5t-18-.5t-16.5.5q-17 0-66.5-.5T904 157t-64 2t-71 6q-9 81-8 136q0 94 2 388t2 455q0 16-2.5 71.5t0 91.5t12.5 69q40 21 124 42.5t120 37.5q5 40 5 50q0 14-3 29l-34 1q-76 2-218-8t-207-10q-50 0-151 9t-152 9q-3-51-3-52v-9q17-27 61.5-43t98.5-29t78-27q19-42 19-383q0-101-3-303t-3-303V270q0-2 .5-15.5t.5-25t-1-25.5t-3-24t-5-14q-11-12-162-12q-33 0-93 12t-80 26q-19 13-34 72.5t-31.5 111T56 429q-42-26-56-44V2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextWidth(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m81 1l54 27q12 5 211 5q44 0 132-2t132-2q70 0 246.5-1t304.5-.5t247 4.5q33 1 56-31l42-1q4 0 14 .5t14 .5q2 112 2 336q0 80-5 109q-39 14-68 18q-25-44-54-128q-3-9-11-47.5t-15-73.5t-7-36q-10-13-27-19q-5-2-66-2q-30 0-93-1t-103-1t-94 2t-96 7q-9 81-8 136l1 152v-52q0 55 1 154t1.5 180t.5 153q0 16-2.5 71.5t0 91.5t12.5 69q40 21 124 42.5t120 37.5q5 40 5 50q0 14-3 29l-34 1q-76 2-218-8t-207-10q-50 0-151 9t-152 9q-3-51-3-52v-9q17-27 61.5-43t98.5-29t78-27q7-16 11.5-74t6-145.5t1.5-155t-.5-153.5t-.5-89q0-7-2.5-21.5T635 459q0-7 .5-44t1-73t0-76.5t-3-67.5t-6.5-32q-11-12-162-12q-41 0-163 13.5T164 192q-19 12-34 71.5T98.5 375T56 429q-42-26-56-44V2zm1229 1282q12 0 42 19.5t57.5 41.5t59.5 49t36 30q26 21 26 49t-26 49q-4 3-36 30t-59.5 49t-57.5 41.5t-42 19.5q-13 0-20.5-10.5t-10-28.5t-2.5-33.5t1.5-33t1.5-19.5H256q0 2 1.5 19.5t1.5 33t-2.5 33.5t-10 28.5T226 1661q-12 0-42-19.5t-57.5-41.5t-59.5-49t-36-30q-26-21-26-49t26-49q4-3 36-30t59.5-49t57.5-41.5t42-19.5q13 0 20.5 10.5t10 28.5t2.5 33.5t-1.5 33t-1.5 19.5h1024q0-2-1.5-19.5t-1.5-33t2.5-33.5t10-28.5t20.5-10.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Th(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 1120v192q0 40-28 68t-68 28H96q-40 0-68-28t-28-68v-192q0-40 28-68t68-28h320q40 0 68 28t28 68m0-512v192q0 40-28 68t-68 28H96q-40 0-68-28T0 800V608q0-40 28-68t68-28h320q40 0 68 28t28 68m640 512v192q0 40-28 68t-68 28H736q-40 0-68-28t-28-68v-192q0-40 28-68t68-28h320q40 0 68 28t28 68M512 96v192q0 40-28 68t-68 28H96q-40 0-68-28T0 288V96q0-40 28-68T96 0h320q40 0 68 28t28 68m640 512v192q0 40-28 68t-68 28H736q-40 0-68-28t-28-68V608q0-40 28-68t68-28h320q40 0 68 28t28 68m640 512v192q0 40-28 68t-68 28h-320q-40 0-68-28t-28-68v-192q0-40 28-68t68-28h320q40 0 68 28t28 68M1152 96v192q0 40-28 68t-68 28H736q-40 0-68-28t-28-68V96q0-40 28-68t68-28h320q40 0 68 28t28 68m640 512v192q0 40-28 68t-68 28h-320q-40 0-68-28t-28-68V608q0-40 28-68t68-28h320q40 0 68 28t28 68m0-512v192q0 40-28 68t-68 28h-320q-40 0-68-28t-28-68V96q0-40 28-68t68-28h320q40 0 68 28t28 68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThLarge(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 896v384q0 52-38 90t-90 38H128q-52 0-90-38t-38-90V896q0-52 38-90t90-38h512q52 0 90 38t38 90m0-768v384q0 52-38 90t-90 38H128q-52 0-90-38T0 512V128q0-52 38-90t90-38h512q52 0 90 38t38 90m896 768v384q0 52-38 90t-90 38h-512q-52 0-90-38t-38-90V896q0-52 38-90t90-38h512q52 0 90 38t38 90m0-768v384q0 52-38 90t-90 38h-512q-52 0-90-38t-38-90V128q0-52 38-90t90-38h512q52 0 90 38t38 90"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThList(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 1120v192q0 40-28 68t-68 28H96q-40 0-68-28t-28-68v-192q0-40 28-68t68-28h320q40 0 68 28t28 68m0-512v192q0 40-28 68t-68 28H96q-40 0-68-28T0 800V608q0-40 28-68t68-28h320q40 0 68 28t28 68m1280 512v192q0 40-28 68t-68 28H736q-40 0-68-28t-28-68v-192q0-40 28-68t68-28h960q40 0 68 28t28 68M512 96v192q0 40-28 68t-68 28H96q-40 0-68-28T0 288V96q0-40 28-68T96 0h320q40 0 68 28t28 68m1280 512v192q0 40-28 68t-68 28H736q-40 0-68-28t-28-68V608q0-40 28-68t68-28h960q40 0 68 28t28 68m0-512v192q0 40-28 68t-68 28H736q-40 0-68-28t-28-68V96q0-40 28-68t68-28h960q40 0 68 28t28 68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Themeisle(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M852 309q0 29-17 52.5T790 385t-45-23.5t-17-52.5t17-52.5t45-23.5t45 23.5t17 52.5M688 1685v-114q0-30-20.5-51.5T617 1498t-50 21.5t-20 51.5v114q0 30 20.5 52t49.5 22q30 0 50.5-22t20.5-52m172 0v-114q0-30-20-51.5t-50-21.5t-50.5 21.5T719 1571v114q0 30 20.5 52t50.5 22q29 0 49.5-22t20.5-52m174 0v-114q0-30-20.5-51.5T963 1498t-50.5 21.5T892 1571v114q0 30 20.5 52t50.5 22t50.5-22t20.5-52m174 0v-114q0-30-20.5-51.5T1137 1498t-50.5 21.5t-20.5 51.5v114q0 30 20.5 52t50.5 22t50.5-22t20.5-52m268-684q-84 160-232 259.5T921 1360q-123 0-229.5-51.5t-178.5-137T400 974t-41-232q0-88 21-174q-104 175-104 390q0 162 65 312t185 251q30-57 91-57q56 0 86 50q32-50 87-50q56 0 86 50q32-50 87-50t87 50q30-50 86-50q28 0 52.5 15.5t37.5 40.5q112-94 177-231.5t73-287.5m-150-29q0-75-72-75q-17 0-47 6q-95 19-149 19q-226 0-226-243q0-86 30-204q-83 127-83 275q0 150 89 260.5t235 110.5q111 0 210-70q13-48 13-79M884 313q0-50-32-89.5T771 184t-81 39.5t-32 89.5q0 51 31.5 90.5T771 443t81.5-39.5T884 313m629 339q0-96-37.5-179t-113-137t-173.5-54q-77 0-149 35t-127 94q-48 159-48 268q0 104 45.5 157t147.5 53q53 0 142-19q36-6 53-6q51 0 77.5 28t26.5 80q0 26-4 46q75-68 117.5-165.5T1513 652m279 217q0 111-33.5 249.5T1665 1323q-58 64-195 142.5T1242 1570l-4 1v114q0 43-29.5 75t-72.5 32q-56 0-86-50q-32 50-87 50t-87-50q-30 50-86 50q-55 0-87-50q-30 50-86 50q-47 0-75-33.5t-28-81.5q-90 68-198 68q-118 0-211-80q54-1 106-20q-113-31-182-127q32 7 71 7q89 0 164-46q-192-192-240-306q-24-56-24-160q0-57 9-125.5T40.5 741t55-141T182 495t120-42q59 0 81 52q19-29 42-54q2-3 12-13t13-16q10-15 23-38t25-42t28-39q87-111 211.5-177T998 60q35 0 62 4q59-64 146-64q83 0 140 57q5 5 5 12q0 5-6 13.5t-12.5 16t-16 17L1306 126q17 6 36 18t19 24q0 6-16 25q157 138 197 378q25-30 60-30q45 0 100 49q90 80 90 279"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Thermometer(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 1344q0 80-56 136t-136 56t-136-56t-56-136q0-60 35-110t93-71V256h128v907q58 21 93 71t35 110m128 0q0-77-34-144t-94-112V320q0-80-56-136t-136-56t-136 56t-56 136v768q-60 45-94 112t-34 144q0 133 93.5 226.5T448 1664t226.5-93.5T768 1344m128 0q0 185-131.5 316.5T448 1792t-316.5-131.5T0 1344q0-182 128-313V320q0-133 93.5-226.5T448 0t226.5 93.5T768 320v711q128 131 128 313m128-576v128H832V768zm0-256v128H832V512zm0-256v128H832V256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThermometerOne(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 1344q0 80-56 136t-136 56t-136-56t-56-136q0-60 35-110t93-71v-139h128v139q58 21 93 71t35 110m128 0q0-77-34-144t-94-112V320q0-80-56-136t-136-56t-136 56t-56 136v768q-60 45-94 112t-34 144q0 133 93.5 226.5T448 1664t226.5-93.5T768 1344m128 0q0 185-131.5 316.5T448 1792t-316.5-131.5T0 1344q0-182 128-313V320q0-133 93.5-226.5T448 0t226.5 93.5T768 320v711q128 131 128 313m128-576v128H832V768zm0-256v128H832V512zm0-256v128H832V256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThermometerThree(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 1344q0 80-56 136t-136 56t-136-56t-56-136q0-60 35-110t93-71V512h128v651q58 21 93 71t35 110m128 0q0-77-34-144t-94-112V320q0-80-56-136t-136-56t-136 56t-56 136v768q-60 45-94 112t-34 144q0 133 93.5 226.5T448 1664t226.5-93.5T768 1344m128 0q0 185-131.5 316.5T448 1792t-316.5-131.5T0 1344q0-182 128-313V320q0-133 93.5-226.5T448 0t226.5 93.5T768 320v711q128 131 128 313m128-576v128H832V768zm0-256v128H832V512zm0-256v128H832V256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThermometerTwo(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 1344q0 80-56 136t-136 56t-136-56t-56-136q0-60 35-110t93-71V768h128v395q58 21 93 71t35 110m128 0q0-77-34-144t-94-112V320q0-80-56-136t-136-56t-136 56t-56 136v768q-60 45-94 112t-34 144q0 133 93.5 226.5T448 1664t226.5-93.5T768 1344m128 0q0 185-131.5 316.5T448 1792t-316.5-131.5T0 1344q0-182 128-313V320q0-133 93.5-226.5T448 0t226.5 93.5T768 320v711q128 131 128 313m128-576v128H832V768zm0-256v128H832V512zm0-256v128H832V256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThermometerZero(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 1344q0 80-56 136t-136 56t-136-56t-56-136q0-79 56-135.5t136-56.5t136 56.5t56 135.5m128 0q0-77-34-144t-94-112V320q0-80-56-136t-136-56t-136 56t-56 136v768q-60 45-94 112t-34 144q0 133 93.5 226.5T448 1664t226.5-93.5T768 1344m128 0q0 185-131.5 316.5T448 1792t-316.5-131.5T0 1344q0-182 128-313V320q0-133 93.5-226.5T448 0t226.5 93.5T768 320v711q128 131 128 313m128-576v128H832V768zm0-256v128H832V512zm0-256v128H832V256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbTack(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M480 736V288q0-14-9-23t-23-9t-23 9t-9 23v448q0 14 9 23t23 9t23-9t9-23m672 352q0 26-19 45t-45 19H659l-51 483q-2 12-10.5 20.5T577 1664h-1q-27 0-32-27l-76-485H64q-26 0-45-19t-19-45q0-123 78.5-221.5T256 768V256q-52 0-90-38t-38-90t38-90t90-38h640q52 0 90 38t38 90t-38 90t-90 38v512q99 0 177.5 98.5T1152 1088"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsDown(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 320q0 26-19 45t-45 19q-27 0-45.5-19T128 320q0-27 18.5-45.5T192 256q26 0 45 18.5t19 45.5m160 512V192q0-26-19-45t-45-19H64q-26 0-45 19T0 192v640q0 26 19 45t45 19h288q26 0 45-19t19-45m1129-149q55 61 55 149q-1 78-57.5 135t-134.5 57h-277q4 14 8 24t11 22t10 18q18 37 27 57t19 58.5t10 76.5q0 24-.5 39t-5 45t-12 50t-24 45t-40 40.5t-60 26T992 1536q-26 0-45-19q-20-20-34-50t-19.5-52t-12.5-61q-9-42-13.5-60.5T850 1245t-31-48q-33-33-101-120q-49-64-101-121t-76-59q-25-2-43-20.5T480 833V192q0-26 19-44.5t45-19.5q35-1 158-44q77-26 120.5-39.5t121.5-29T1088 0h129q133 2 197 78q58 69 49 181q39 37 54 94q17 61 0 117q46 61 43 137q0 32-15 76"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsOup(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 1344q0-26-19-45t-45-19t-45 19t-19 45t19 45t45 19t45-19t19-45m1152-576q0-51-39-89.5t-89-38.5H928q0-58 48-159.5t48-160.5q0-98-32-145t-128-47q-26 26-38 85t-30.5 125.5T736 448q-22 23-77 91q-4 5-23 30t-31.5 41t-34.5 42.5t-40 44t-38.5 35.5t-40 27t-35.5 9h-32v640h32q13 0 31.5 3t33 6.5t38 11t35 11.5t35.5 12.5t29 10.5q211 73 342 73h121q192 0 192-167q0-26-5-56q30-16 47.5-52.5t17.5-73.5t-18-69q53-50 53-119q0-25-10-55.5t-25-47.5q32-1 53.5-47t21.5-81m128-1q0 89-49 163q9 33 9 69q0 77-38 144q3 21 3 43q0 101-60 178q1 139-85 219.5t-227 80.5H960q-96 0-189.5-22.5T554 1576q-116-40-138-40H128q-53 0-90.5-37.5T0 1408V768q0-53 37.5-90.5T128 640h274q36-24 137-155q58-75 107-128q24-25 35.5-85.5T712 145t62-108q39-37 90-37q84 0 151 32.5T1117 134t35 186q0 93-48 192h176q104 0 180 76t76 179"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsUp(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 1216q0-26-19-45t-45-19q-27 0-45.5 19t-18.5 45q0 27 18.5 45.5T192 1280q26 0 45-18.5t19-45.5m160-512v640q0 26-19 45t-45 19H64q-26 0-45-19t-19-45V704q0-26 19-45t45-19h288q26 0 45 19t19 45m1184 0q0 86-55 149q15 44 15 76q3 76-43 137q17 56 0 117q-15 57-54 94q9 112-49 181q-64 76-197 78h-129q-66 0-144-15.5t-121.5-29T702 1452q-123-43-158-44q-26-1-45-19.5t-19-44.5V703q0-25 18-43.5t43-20.5q24-2 76-59t101-121q68-87 101-120q18-18 31-48t17.5-48.5T881 182q7-39 12.5-61T913 69t34-50q19-19 45-19q46 0 82.5 10.5t60 26t40 40.5t24 45t12 50t5 45t.5 39q0 38-9.5 76t-19 60t-27.5 56q-3 6-10 18t-11 22t-8 24h277q78 0 135 57t57 135"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ticket(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m992 420l316 316l-572 572l-316-316zm-211 979l618-618q19-19 19-45t-19-45l-362-362q-18-18-45-18t-45 18L329 947q-19 19-19 45t19 45l362 362q18 18 45 18t45-18m889-637l-907 908q-37 37-90.5 37t-90.5-37l-126-126q56-56 56-136t-56-136t-136-56t-136 56L59 1146q-37-37-37-90.5T59 965L966 59q37-37 90.5-37t90.5 37l125 125q-56 56-56 136t56 136t136 56t136-56l126 125q37 37 37 90.5t-37 90.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimesCircle(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1149 994q0-26-19-45L949 768l181-181q19-19 19-45q0-27-19-46l-90-90q-19-19-46-19q-26 0-45 19L768 587L587 406q-19-19-45-19q-27 0-46 19l-90 90q-19 19-19 46q0 26 19 45l181 181l-181 181q-19 19-19 45q0 27 19 46l90 90q19 19 46 19q26 0 45-19l181-181l181 181q19 19 45 19q27 0 46-19l90-90q19-19 19-46m387-226q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimesCircleO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1097 951l-146 146q-10 10-23 10t-23-10L768 960l-137 137q-10 10-23 10t-23-10L439 951q-10-10-10-23t10-23l137-137l-137-137q-10-10-10-23t10-23l146-146q10-10 23-10t23 10l137 137l137-137q10-10 23-10t23 10l146 146q10 10 10 23t-10 23L960 768l137 137q10 10 10 23t-10 23m215-183q0-148-73-273t-198-198t-273-73t-273 73t-198 198t-73 273t73 273t198 198t273 73t273-73t198-198t73-273m224 0q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimesRectangle(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1175 1193l146-146q10-10 10-23t-10-23l-233-233l233-233q10-10 10-23t-10-23l-146-146q-10-10-23-10t-23 10L896 576L663 343q-10-10-23-10t-23 10L471 489q-10 10-10 23t10 23l233 233l-233 233q-10 10-10 23t10 23l146 146q10 10 23 10t23-10l233-233l233 233q10 10 23 10t23-10m617-1033v1216q0 66-47 113t-113 47H160q-66 0-113-47T0 1376V160Q0 94 47 47T160 0h1472q66 0 113 47t47 113"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimesRectangleO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1257 983l-146 146q-10 10-23 10t-23-10L896 960l-169 169q-10 10-23 10t-23-10L535 983q-10-10-10-23t10-23l169-169l-169-169q-10-10-10-23t10-23l146-146q10-10 23-10t23 10l169 169l169-169q10-10 23-10t23 10l146 146q10 10 10 23t-10 23l-169 169l169 169q10 10 10 23t-10 23M256 1280h1280V256H256zM1792 160v1216q0 66-47 113t-113 47H160q-66 0-113-47T0 1376V160Q0 94 47 47T160 0h1472q66 0 113 47t47 113"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tint(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 1088q0-36-20-69q-1-1-15.5-22.5t-25.5-38t-25-44t-21-50.5q-4-16-21-16t-21 16q-7 23-21 50.5t-25 44t-25.5 38T276 1019q-20 33-20 69q0 53 37.5 90.5T384 1216t90.5-37.5T512 1088m512-128q0 212-150 362t-362 150t-362-150T0 960q0-145 81-275q6-9 62.5-90.5t101-151t99.5-178T427 64q9-30 34-47t51-17t51.5 17T597 64q28 93 83 201.5t99.5 178t101 151T943 685q81 127 81 275"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleOff(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1152 640q0-104-40.5-198.5T1002 278T838.5 168.5T640 128t-198.5 40.5T278 278T168.5 441.5T128 640t40.5 198.5T278 1002t163.5 109.5T640 1152t198.5-40.5T1002 1002t109.5-163.5T1152 640m768 0q0-104-40.5-198.5T1770 278t-163.5-109.5T1408 128h-386q119 90 188.5 224t69.5 288t-69.5 288t-188.5 224h386q104 0 198.5-40.5T1770 1002t109.5-163.5T1920 640m128 0q0 130-51 248.5t-136.5 204t-204 136.5t-248.5 51H640q-130 0-248.5-51t-204-136.5T51 888.5T0 640t51-248.5t136.5-204T391.5 51T640 0h768q130 0 248.5 51t204 136.5t136.5 204t51 248.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleOn(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 640q0-130 51-248.5t136.5-204T391.5 51T640 0h768q130 0 248.5 51t204 136.5t136.5 204t51 248.5t-51 248.5t-136.5 204t-204 136.5t-248.5 51H640q-130 0-248.5-51t-204-136.5T51 888.5T0 640m1408 512q104 0 198.5-40.5T1770 1002t109.5-163.5T1920 640t-40.5-198.5T1770 278t-163.5-109.5T1408 128t-198.5 40.5T1046 278T936.5 441.5T896 640t40.5 198.5T1046 1002t163.5 109.5T1408 1152"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trademark(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M857 288v117q0 13-9.5 22t-22.5 9H527v812q0 13-9 22.5t-22 9.5H361q-13 0-22.5-9t-9.5-23V436H32q-13 0-22.5-9T0 405V288q0-14 9-23t23-9h793q13 0 22.5 9.5T857 288m1038-3l77 961q1 13-8 24q-10 10-23 10h-134q-12 0-21-8.5t-10-20.5l-46-588l-189 425q-8 19-29 19h-120q-20 0-29-19l-188-427l-45 590q-1 12-10 20.5t-21 8.5H964q-13 0-23-10q-9-10-9-24l78-961q1-12 10-20.5t21-8.5h142q20 0 29 19l220 520q10 24 20 51q3-7 9.5-24.5T1472 795l221-520q9-19 29-19h141q13 0 22 8.5t10 20.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Train(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1088 0q185 0 316.5 93.5T1536 320v896q0 130-125.5 222t-305.5 97l213 202q16 15 8 35t-30 20H240q-22 0-30-20t8-35l213-202q-180-5-305.5-97T0 1216V320Q0 187 131.5 93.5T448 0zM768 1344q80 0 136-56t56-136t-56-136t-136-56t-136 56t-56 136t56 136t136 56m576-576V256H192v512z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransgenderAlt(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1280 32q0-14 9-23t23-9h288q26 0 45 19t19 45v288q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23V218l-254 255q126 158 126 359q0 221-147.5 384.5T896 1404v132h96q14 0 23 9t9 23v64q0 14-9 23t-23 9h-96v96q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23v-96h-96q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h96v-132q-217-24-364.5-187.5T256 832q0-201 126-359l-52-53l-101 111q-9 10-22 10.5t-23-7.5l-48-44q-10-8-10.5-21.5T134 445l105-115l-111-112v134q0 14-9 23t-23 9H32q-14 0-23-9t-9-23V64q0-26 19-45T64 0h288q14 0 23 9t9 23v64q0 14-9 23t-23 9H219l106 107l86-94q9-10 22-10.5t23 7.5l48 44q10 8 10.5 21.5T506 227l-90 99l57 56q158-126 359-126t359 126l255-254h-134q-14 0-23-9t-9-23zM832 1280q185 0 316.5-131.5T1280 832t-131.5-316.5T832 384T515.5 515.5T384 832t131.5 316.5T832 1280"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 1248V544q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v704q0 14 9 23t23 9h64q14 0 23-9t9-23m256 0V544q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v704q0 14 9 23t23 9h64q14 0 23-9t9-23m256 0V544q0-14-9-23t-23-9h-64q-14 0-23 9t-9 23v704q0 14 9 23t23 9h64q14 0 23-9t9-23M480 256h448l-48-117q-7-9-17-11H546q-10 2-17 11zm928 32v64q0 14-9 23t-23 9h-96v948q0 83-47 143.5t-113 60.5H288q-66 0-113-58.5T128 1336V384H32q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h309l70-167q15-37 54-63t79-26h320q40 0 79 26t54 63l70 167h309q14 0 23 9t9 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrashO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 608v576q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23V608q0-14 9-23t23-9h64q14 0 23 9t9 23m256 0v576q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23V608q0-14 9-23t23-9h64q14 0 23 9t9 23m256 0v576q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23V608q0-14 9-23t23-9h64q14 0 23 9t9 23m128 724V384H256v948q0 22 7 40.5t14.5 27t10.5 8.5h832q3 0 10.5-8.5t14.5-27t7-40.5M480 256h448l-48-117q-7-9-17-11H546q-10 2-17 11zm928 32v64q0 14-9 23t-23 9h-96v948q0 83-47 143.5t-113 60.5H288q-66 0-113-58.5T128 1336V384H32q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h309l70-167q15-37 54-63t79-26h320q40 0 79 26t54 63l70 167h309q14 0 23 9t9 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tree(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1472 1472q0 26-19 45t-45 19H946q1 17 6 87.5t5 108.5q0 25-18 42.5t-43 17.5H576q-25 0-43-17.5t-18-42.5q0-38 5-108.5t6-87.5H64q-26 0-45-19t-19-45t19-45l402-403H192q-26 0-45-19t-19-45t19-45l402-403H352q-26 0-45-19t-19-45t19-45L691 19q19-19 45-19t45 19l384 384q19 19 19 45t-19 45t-45 19H923l402 403q19 19 19 45t-19 45t-45 19h-229l402 403q19 19 19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trello(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M704 1216V192q0-14-9-23t-23-9H192q-14 0-23 9t-9 23v1024q0 14 9 23t23 9h480q14 0 23-9t9-23m672-384V192q0-14-9-23t-23-9H864q-14 0-23 9t-9 23v640q0 14 9 23t23 9h480q14 0 23-9t9-23m160-768v1408q0 26-19 45t-45 19H64q-26 0-45-19t-19-45V64q0-26 19-45T64 0h1408q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tripadvisor(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M651 805q0 39-27.5 66.5T558 899q-39 0-66.5-27.5T464 805q0-38 27.5-65.5T558 712q38 0 65.5 27.5T651 805m1154-1q0 39-27.5 66.5T1711 898t-66.5-27.5T1617 804t27.5-66t66.5-27t66.5 27t27.5 66m-1040 1q0-79-56.5-136T572 612t-136.5 56.5T379 805t56.5 136.5T572 998t136.5-56.5T765 805m1153-1q0-80-56.5-136.5T1725 611q-79 0-136 56.5T1532 804t56.5 136.5T1725 997t136.5-56.5T1918 804m-1068 1q0 116-81.5 197.5T572 1084q-116 0-197.5-82T293 805t82-196.5T572 527t196.5 81.5T850 805m1154-1q0 115-81.5 196.5T1725 1082q-115 0-196.5-81.5T1447 804t81.5-196.5T1725 526q116 0 197.5 81.5T2004 804m-964 3q0-191-135.5-326.5T578 345q-125 0-231 62T179 575.5T117 807t62 231.5T347 1207t231 62q191 0 326.5-135.5T1040 807m668-573q-254-111-556-111q-319 0-573 110q117 0 223 45.5T984.5 401t122 183t45.5 223q0-115 43.5-219.5t118-180.5T1491 284t217-50m479 573q0-191-135-326.5T1726 345t-326.5 135.5T1264 807t135.5 326.5T1726 1269t326-135.5T2187 807m-266-566h383q-44 51-75 114.5T2189 470q110 151 110 337q0 156-77 288t-209 208.5t-287 76.5q-133 0-249-56t-196-155q-47 56-129 179q-11-22-53.5-82.5T1024 1168q-80 99-196.5 155.5T578 1380q-155 0-287-76.5T82 1095T5 807q0-186 110-337q-9-51-40-114.5T0 241h365Q514 141 720 84.5T1152 28q224 0 421 56t348 157"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trophy(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M458 755q-74-162-74-371H128v96q0 78 94.5 162T458 755m1078-275v-96h-256q0 209-74 371q141-29 235.5-113t94.5-162m128-128v128q0 71-41.5 143t-112 130t-173 97.5T1122 895q-42 54-95 95q-38 34-52.5 72.5T960 1152q0 54 30.5 91t97.5 37q75 0 133.5 45.5T1280 1440v64q0 14-9 23t-23 9H416q-14 0-23-9t-9-23v-64q0-69 58.5-114.5T576 1280q67 0 97.5-37t30.5-91q0-51-14.5-89.5T637 990q-53-41-95-95q-113-5-215.5-44.5t-173-97.5t-112-130T0 480V352q0-40 28-68t68-28h288v-96q0-66 47-113T544 0h576q66 0 113 47t47 113v96h288q40 0 68 28t28 68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Truck(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M576 1152q0-52-38-90t-90-38t-90 38t-38 90t38 90t90 38t90-38t38-90M192 640h384V384H418q-13 0-22 9L201 588q-9 9-9 22zm1280 512q0-52-38-90t-90-38t-90 38t-38 90t38 90t90 38t90-38t38-90M1728 64v1024q0 15-4 26.5t-13.5 18.5t-16.5 11.5t-23.5 6t-22.5 2t-25.5 0t-22.5-.5q0 106-75 181t-181 75t-181-75t-75-181H704q0 106-75 181t-181 75t-181-75t-75-181h-64q-3 0-22.5.5t-25.5 0t-22.5-2t-23.5-6t-16.5-11.5T4 1114.5T0 1088q0-26 19-45t45-19V704q0-8-.5-35t0-38t2.5-34.5t6.5-37t14-30.5t22.5-30l198-198q19-19 50.5-32t58.5-13h160V64q0-26 19-45t45-19h1024q26 0 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Try(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1152 704q0 191-94.5 353T801 1313.5T448 1408H288q-14 0-23-9t-9-23V765L41 831q-3 1-9 1q-10 0-19-6q-13-10-13-26V672q0-23 23-31l233-71v-93L41 543q-3 1-9 1q-10 0-19-6q-13-10-13-26V384q0-23 23-31l233-71V32q0-14 9-23t23-9h160q14 0 23 9t9 23v181L855 97q15-5 28 5t13 26v128q0 23-23 31L480 408v93l375-116q15-5 28 5t13 26v128q0 23-23 31L480 696v487q188-13 318-151t130-328q0-14 9-23t23-9h160q14 0 23 9t9 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tty(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 1184v192q0 14-9 23t-23 9H224q-14 0-23-9t-9-23v-192q0-14 9-23t23-9h192q14 0 23 9t9 23M256 800v192q0 14-9 23t-23 9H32q-14 0-23-9t-9-23V800q0-14 9-23t23-9h192q14 0 23 9t9 23m576 384v192q0 14-9 23t-23 9H608q-14 0-23-9t-9-23v-192q0-14 9-23t23-9h192q14 0 23 9t9 23M640 800v192q0 14-9 23t-23 9H416q-14 0-23-9t-9-23V800q0-14 9-23t23-9h192q14 0 23 9t9 23M66 640q-28 0-47-19T0 575V446h514v129q0 27-19 46t-46 19zm1150 544v192q0 14-9 23t-23 9H992q-14 0-23-9t-9-23v-192q0-14 9-23t23-9h192q14 0 23 9t9 23m-192-384v192q0 14-9 23t-23 9H800q-14 0-23-9t-9-23V800q0-14 9-23t23-9h192q14 0 23 9t9 23m576 384v192q0 14-9 23t-23 9h-192q-14 0-23-9t-9-23v-192q0-14 9-23t23-9h192q14 0 23 9t9 23m-192-384v192q0 14-9 23t-23 9h-192q-14 0-23-9t-9-23V800q0-14 9-23t23-9h192q14 0 23 9t9 23m384-408v13h-514v-10q0-104-382-102q-382 1-382 102v10H0v-13q0-17 8.5-43t34-64t65.5-75.5t110.5-76t160-67.5t224-47.5T896 0t293 18.5T1413 66t160.5 67.5t110.5 76t65.5 75.5t34 64t8.5 43m0 408v192q0 14-9 23t-23 9h-192q-14 0-23-9t-9-23V800q0-14 9-23t23-9h192q14 0 23 9t9 23m0-354v129q0 27-19 46t-46 19h-384q-27 0-46-19t-19-46V446z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tumblr(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m880 1329l80 237q-23 35-111 66t-177 32q-104 2-190.5-26T339 1564t-95-106t-55.5-120t-16.5-118V676H4V461q72-26 129-69.5t91-90t58-102t34-99T331 12q1-5 4.5-8.5T343 0h244v424h333v252H586v518q0 30 6.5 56t22.5 52.5t49.5 41.5t81.5 14q78-2 134-29"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TumblrSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1136 1333l-62-183q-44 22-103 22q-36 1-62-10.5t-38.5-31.5t-17.5-40.5t-5-43.5V648h257V454H849V128H661q-8 0-9 10q-5 44-17.5 87t-39 95t-77 95T400 483v165h130v418q0 57 21.5 115t65 111t121 85.5T914 1408q69-1 136.5-25t85.5-50m400-1045v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitch(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M800 434v434H655V434zm398 0v434h-145V434zm0 760l253-254V145H257v1049h326v217l217-217zM1596 0v1013l-434 434H836l-217 217H402v-217H4V289L113 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitter(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1588 152q-67 98-162 167q1 14 1 42q0 130-38 259.5T1273.5 869T1089 1079.5t-258 146t-323 54.5q-271 0-496-145q35 4 78 4q225 0 401-138q-105-2-188-64.5T189 777q33 5 61 5q43 0 85-11q-112-23-185.5-111.5T76 454v-4q68 38 146 41q-66-44-105-115T78 222q0-88 44-163q121 149 294.5 238.5T788 397q-8-38-8-74q0-134 94.5-228.5T1103 0q140 0 236 102q109-21 205-78q-37 115-142 178q93-10 186-50"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwitterSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1280 482q-56 25-121 34q68-40 93-117q-65 38-134 51q-61-66-153-66q-87 0-148.5 61.5T755 594q0 29 5 48q-129-7-242-65T326 422q-29 50-29 106q0 114 91 175q-47-1-100-26v2q0 75 50 133.5T461 885q-29 8-51 8q-13 0-39-4q21 63 74.5 104t121.5 42q-116 90-261 90q-26 0-50-3q148 94 322 94q112 0 210-35.5t168-95t120.5-137t75-162T1176 618q0-18-1-27q63-45 105-109m256-194v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Umbrella(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M896 764v580q0 104-76 180t-180 76t-180-76t-76-180q0-26 19-45t45-19t45 19t19 45q0 50 39 89t89 39t89-39t39-89V764q33-11 64-11t64 11m768 27q0 13-9.5 22.5T1632 823q-11 0-23-10q-49-46-93-69t-102-23q-68 0-128 37t-103 97q-7 10-17.5 28t-14.5 24q-11 17-28 17q-18 0-29-17q-4-6-14.5-24t-17.5-28q-43-60-102.5-97T832 721t-127.5 37T602 855q-7 10-17.5 28T570 907q-11 17-29 17q-17 0-28-17q-4-6-14.5-24T481 855q-43-60-103-97t-128-37q-58 0-102 23t-93 69q-12 10-23 10q-13 0-22.5-9.5T0 791q0-5 1-7q45-183 172.5-319.5t298-204.5T832 192q140 0 274.5 40T1353 345.5t194.5 187T1663 784q1 2 1 7M896 64v98q-42-2-64-2t-64 2V64q0-26 19-45t45-19t45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Underline(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M48 95q-37-2-45-4L0 3q13-1 40-1q60 0 112 4q132 7 166 7q86 0 168-3q116-4 146-5q56 0 86-2l-1 14l2 64v9q-60 9-124 9q-60 0-79 25q-13 14-13 132q0 13 .5 32.5t.5 25.5l1 229l14 280q6 124 51 202q35 59 96 92q88 47 177 47q104 0 191-28q56-18 99-51q48-36 65-64q36-56 53-114q21-73 21-229q0-79-3.5-128t-11-122.5T1244 268l-4-59q-5-67-24-88q-34-35-77-34l-100 2l-14-3l2-86h84l205 10q76 3 196-10l18 2q6 38 6 51q0 7-4 31q-45 12-84 13q-73 11-79 17q-15 15-15 41q0 7 1.5 27t1.5 31q8 19 22 396q6 195-15 304q-15 76-41 122q-38 65-112 123q-75 57-182 89q-109 33-255 33q-167 0-284-46q-119-47-179-122q-61-76-83-195q-16-80-16-237V347q0-188-17-213q-25-36-147-39m1488 1409v-64q0-14-9-23t-23-9H32q-14 0-23 9t-9 23v64q0 14 9 23t23 9h1472q14 0 23-9t9-23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UniversalAccess(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1374 657q-6-26-28.5-39.5T1297 610q-261 62-401 62t-401-62q-26-6-48.5 7.5T418 657t7.5 48.5T465 734q194 46 303 58q-2 158-15.5 269T726 1216.5T685 1332l-9 21q-10 25 1 49t36 34q9 4 23 4q44 0 60-41l8-20q54-139 71-259h42q17 120 71 259l8 20q16 41 60 41q14 0 23-4q25-10 36-34t1-49l-9-21q-28-71-41-115.5t-26.5-155.5t-15.5-269q109-12 303-58q26-6 39.5-28.5t7.5-48.5m-350-145q0-53-37.5-90.5T896 384t-90.5 37.5T768 512t37.5 90.5T896 640t90.5-37.5T1024 512m576 384q0 143-55.5 273.5t-150 225t-225 150T896 1600t-273.5-55.5t-225-150t-150-225T192 896t55.5-273.5t150-225t225-150T896 192t273.5 55.5t225 150t150 225T1600 896M896 128q-156 0-298 61T353 353T189 598t-61 298t61 298t164 245t245 164t298 61t298-61t245-164t164-245t61-298t-61-298t-164-245t-245-164t-298-61m896 768q0 182-71 348t-191 286t-286 191t-348 71t-348-71t-286-191t-191-286T0 896t71-348t191-286T548 71T896 0t348 71t286 191t191 286t71 348"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlock(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1664 448v256q0 26-19 45t-45 19h-64q-26 0-45-19t-19-45V448q0-106-75-181t-181-75t-181 75t-75 181v192h96q40 0 68 28t28 68v576q0 40-28 68t-68 28H96q-40 0-68-28t-28-68V736q0-40 28-68t68-28h672V448q0-185 131.5-316.5T1216 0t316.5 131.5T1664 448"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnlockAlt(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1056 768q40 0 68 28t28 68v576q0 40-28 68t-68 28H96q-40 0-68-28t-28-68V864q0-40 28-68t68-28h32V448q0-185 131.5-316.5T576 0t316.5 131.5T1024 448q0 26-19 45t-45 19h-64q-26 0-45-19t-19-45q0-106-75-181t-181-75t-181 75t-75 181v320z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Upload(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1280 1408q0-26-19-45t-45-19t-45 19t-19 45t19 45t45 19t45-19t19-45m256 0q0-26-19-45t-45-19t-45 19t-19 45t19 45t45 19t45-19t19-45m128-224v320q0 40-28 68t-68 28H96q-40 0-68-28t-28-68v-320q0-40 28-68t68-28h427q21 56 70.5 92t110.5 36h256q61 0 110.5-36t70.5-92h427q40 0 68 28t28 68m-325-648q-17 40-59 40h-256v448q0 26-19 45t-45 19H704q-26 0-45-19t-19-45V576H384q-42 0-59-40q-17-39 14-69L787 19q18-19 45-19t45 19l448 448q31 30 14 69"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Usb(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2288 677q16 8 16 27t-16 27l-320 192q-8 5-16 5q-9 0-16-4q-16-10-16-28V768h-858q37 58 83 165q16 37 24.5 55t24 49t27 47t27 34t31.5 26t33 8h96v-96q0-14 9-23t23-9h320q14 0 23 9t9 23v320q0 14-9 23t-23 9h-320q-14 0-23-9t-9-23v-96h-96q-32 0-61-10t-51-23.5t-45-40.5t-37-46t-33.5-57t-28.5-57.5t-28-60.5q-23-53-37-81.5t-36-65t-44.5-53.5t-46.5-17H504q-22 84-91 138t-157 54q-106 0-181-75T0 704t75-181t181-75q88 0 157 54t91 138h104q24 0 46.5-17t44.5-53.5t36-65t37-81.5q19-41 28-60.5t28.5-57.5t33.5-57t37-46t45-40.5t51-23.5t61-10h107q21-57 70-92.5T1344 0q80 0 136 56t56 136t-56 136t-136 56q-62 0-111-35.5t-70-92.5h-107q-17 0-33 8t-31.5 26t-27 34t-27 47t-24 49t-24.5 55q-46 107-83 165h1114V512q0-18 16-28t32 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func User(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1280 1271q0 109-62.5 187t-150.5 78H213q-88 0-150.5-78T0 1271q0-85 8.5-160.5t31.5-152t58.5-131t94-89T327 704q131 128 313 128t313-128q76 0 134.5 34.5t94 89t58.5 131t31.5 152t8.5 160.5m-256-887q0 159-112.5 271.5T640 768T368.5 655.5T256 384t112.5-271.5T640 0t271.5 112.5T1024 384"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCircle(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1523 1339q-22-155-87.5-257.5T1251 963q-67 74-159.5 115.5T896 1120t-195.5-41.5T541 963q-119 16-184.5 118.5T269 1339q106 150 271 237.5t356 87.5t356-87.5t271-237.5m-243-699q0-159-112.5-271.5T896 256T624.5 368.5T512 640t112.5 271.5T896 1024t271.5-112.5T1280 640m512 256q0 182-71 347.5t-190.5 286T1245 1721t-349 71q-182 0-348-71t-286-191t-191-286T0 896t71-348t191-286T548 71T896 0t348 71t286 191t191 286t71 348"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCircleO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M896 0q182 0 348 71t286 191t191 286t71 348q0 181-70.5 347T1531 1529t-286 191.5t-349 71.5t-349-71t-285.5-191.5t-190.5-286T0 896t71-348t191-286T548 71T896 0m619 1351q149-205 149-455q0-156-61-298t-164-245t-245-164t-298-61t-298 61t-245 164t-164 245t-61 298q0 250 149 455q66-327 306-327q131 128 313 128t313-128q240 0 306 327m-235-647q0-159-112.5-271.5T896 320T624.5 432.5T512 704t112.5 271.5T896 1088t271.5-112.5T1280 704"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserMd(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 1216q0 26-19 45t-45 19t-45-19t-19-45t19-45t45-19t45 19t19 45m1024 61q0 121-73 190t-194 69H267q-121 0-194-69T0 1277q0-68 5.5-131t24-138T77 875.5t81-103T278 712q-22 52-22 120v203q-58 20-93 70t-35 111q0 80 56 136t136 56t136-56t56-136q0-61-35.5-111t-92.5-70V832q0-62 25-93q132 104 295 104t295-104q25 31 25 93v64q-106 0-181 75t-75 181v89q-32 29-32 71q0 40 28 68t68 28t68-28t28-68q0-42-32-71v-89q0-52 38-90t90-38t90 38t38 90v89q-32 29-32 71q0 40 28 68t68 28t68-28t28-68q0-42-32-71v-89q0-68-34.5-127.5T1152 931q0-10 .5-42.5t0-48t-2.5-41.5t-7-47t-13-40q68 15 120 60.5t81 103t47.5 132.5t24 138t5.5 131m-320-893q0 159-112.5 271.5T704 768T432.5 655.5T320 384t112.5-271.5T704 0t271.5 112.5T1088 384"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserO(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1201 784q47 14 89.5 38t89 73t79.5 115.5t55 172t22 236.5q0 154-100 263.5T1195 1792H341q-141 0-241-109.5T0 1419q0-131 22-236.5t55-172T156.5 895t89-73t89.5-38q-79-125-79-272q0-104 40.5-198.5T406 150T569.5 40.5T768 0t198.5 40.5T1130 150t109.5 163.5T1280 512q0 147-79 272M768 128q-159 0-271.5 112.5T384 512t112.5 271.5T768 896t271.5-112.5T1152 512t-112.5-271.5T768 128m427 1536q88 0 150.5-71.5T1408 1419q0-239-78.5-377T1104 897q-145 127-336 127T432 897q-147 7-225.5 145T128 1419q0 102 62.5 173.5T341 1664z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserPlus(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M704 768q-159 0-271.5-112.5T320 384t112.5-271.5T704 0t271.5 112.5T1088 384T975.5 655.5T704 768m960 128h352q13 0 22.5 9.5t9.5 22.5v192q0 13-9.5 22.5t-22.5 9.5h-352v352q0 13-9.5 22.5t-22.5 9.5h-192q-13 0-22.5-9.5t-9.5-22.5v-352h-352q-13 0-22.5-9.5t-9.5-22.5V928q0-13 9.5-22.5t22.5-9.5h352V544q0-13 9.5-22.5t22.5-9.5h192q13 0 22.5 9.5t9.5 22.5zm-736 224q0 52 38 90t90 38h256v238q-68 50-171 50H267q-121 0-194-69T0 1277q0-53 3.5-103.5t14-109T44 956t43-97.5t62-81t85.5-53.5T346 704q19 0 39 17q79 61 154.5 91.5T704 843t164.5-30.5T1023 721q20-17 39-17q132 0 217 96h-223q-52 0-90 38t-38 90z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserSecret(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m576 1536l96-448l-96-128l-128-64zm256 0l128-640l-128 64l-96 128zM992 526q-2-4-4-6q-10-8-96-8q-70 0-167 19q-7 2-21 2t-21-2q-97-19-167-19q-86 0-96 8q-2 2-4 6q2 18 4 27q2 3 7.5 6.5T435 570q2 4 7.5 20.5t7 20.5t7.5 17t8.5 17t9 14t12 13.5t14 9.5t17.5 8t20.5 4t24.5 2q36 0 59-12.5t32.5-30T669 619t11.5-29.5T698 577h12q11 0 17.5 12.5T739 619t14.5 34.5t32.5 30t59 12.5q13 0 24.5-2t20.5-4t17.5-8t14-9.5t12-13.5t9-14t8.5-17t7.5-17t7-20.5T973 570q2-7 7.5-10.5t7.5-6.5q2-9 4-27m416 879q0 121-73 190t-194 69H267q-121 0-194-69T0 1405q0-61 4.5-118t19-125.5T61 1038t63.5-103.5T218 860l-90-220h214q-22-64-22-128q0-12 2-32q-194-40-194-96q0-57 210-99q17-62 51.5-134T460 37q32-37 76-37q30 0 84 31t84 31t84-31t84-31q44 0 76 37q36 42 70.5 114t51.5 134q210 42 210 99q0 56-194 96q7 81-20 160h214l-82 225q63 33 107.5 96.5T1371 1105t29 151.5t8 148.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserTimes(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M704 768q-159 0-271.5-112.5T320 384t112.5-271.5T704 0t271.5 112.5T1088 384T975.5 655.5T704 768m1077 320l249 249q9 9 9 23q0 13-9 22l-136 136q-9 9-22 9q-14 0-23-9l-249-249l-249 249q-9 9-23 9q-13 0-22-9l-136-136q-9-9-9-22q0-14 9-23l249-249l-249-249q-9-9-9-23q0-13 9-22l136-136q9-9 22-9q14 0 23 9l249 249l249-249q9-9 23-9q13 0 22 9l136 136q9 9 9 22q0 14-9 23zm-498 0l-181 181q-37 37-37 91q0 53 37 90l83 83q-21 3-44 3H267q-121 0-194-69T0 1277q0-53 3.5-103.5t14-109T44 956t43-97.5t62-81t85.5-53.5T346 704q19 0 39 17q154 122 319 122t319-122q20-17 39-17q28 0 57 6q-28 27-41 50t-13 56q0 54 37 91z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Venus(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1152 576q0 221-147.5 384.5T640 1148v260h224q14 0 23 9t9 23v64q0 14-9 23t-23 9H640v224q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23v-224H288q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h224v-260q-150-16-271.5-103t-186-224T2 529q11-134 80.5-249t182-188T510 4q170-19 319 54t236 212t87 306m-1024 0q0 185 131.5 316.5T576 1024t316.5-131.5T1024 576T892.5 259.5T576 128T259.5 259.5T128 576"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VenusDouble(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1790 529q12 155-52.5 292t-186 224t-271.5 103v260h224q14 0 23 9t9 23v64q0 14-9 23t-23 9h-224v224q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23v-224H640v224q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23v-224H288q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h224v-260q-150-16-271.5-103t-186-224T2 529q17-206 164.5-356.5T519 3q206-21 377 94q171-115 377-94q205 19 352.5 169.5T1790 529M896 889q128-131 128-313T896 263Q768 394 768 576t128 313m-320 135q115 0 218-57q-154-165-154-391q0-224 154-391q-103-57-218-57q-185 0-316.5 131.5T128 576t131.5 316.5T576 1024m576 384v-260q-137-15-256-94q-119 79-256 94v260zm64-384q185 0 316.5-131.5T1664 576t-131.5-316.5T1216 128q-115 0-218 57q154 167 154 391q0 226-154 391q103 57 218 57"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VenusMars(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1664 32q0-14 9-23t23-9h288q26 0 45 19t19 45v288q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23V218l-254 255q76 95 107.5 214t9.5 247q-32 180-164.5 310T1305 1401q-223 34-409-90q-117 78-256 93v132h96q14 0 23 9t9 23v64q0 14-9 23t-23 9h-96v96q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23v-96h-96q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h96v-132q-155-17-279.5-109.5T45.5 1057T6 750q25-187 159.5-322.5T486 263q224-34 410 90q146-97 320-97q201 0 359 126l255-254h-134q-14 0-23-9t-9-23zM896 1145q128-131 128-313T896 519Q768 650 768 832t128 313M128 832q0 185 131.5 316.5T576 1280q117 0 218-57q-154-167-154-391t154-391q-101-57-218-57q-185 0-316.5 131.5T128 832m1088 448q185 0 316.5-131.5T1664 832t-131.5-316.5T1216 384q-117 0-218 57q154 167 154 391t-154 391q101 57 218 57"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Viacoin(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1536 0l-192 448h192v192h-274l-55 128h329v192h-411l-357 832l-357-832H0V768h329l-55-128H0V448h192L0 0h256l323 768h378L1280 0zM768 1216l108-256H660z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Viadeo(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1050 978q0 215-147 374q-148 161-378 161q-232 0-378-161Q0 1193 0 978q0-147 68-270.5T257 511t268-73q96 0 182 31q-32 62-39 126q-66-28-143-28q-167 0-280.5 123T131 981q0 170 112.5 288.5T525 1388t281-118.5T918 981q0-89-32-166q66-13 123-49q41 98 41 212M846 789q0 192-79.5 345T528 1387l-14 1q-29 0-62-5q83-32 146.5-102.5T698 1126t58.5-189t30-192.5T794 566q0-69-3-103q55 160 55 326m-55-328v2Q718 249 585 23q88 59 142.5 186.5T791 461m244 203q-83 0-160-75q218-120 290-247q19-37 21-56q-42 94-139.5 166.5T842 550q-35-54-35-113q0-37 17-79t43-68q46-44 157-74q59-16 106-58.5T1204 57q74 105 74 253q0 109-24 170q-32 77-88.5 130.5T1035 664"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViadeoSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1050 913q0-78-28-147q-41 25-85 34q22 50 22 114q0 117-77 198.5T689 1194t-193.5-81.5T418 914q0-115 78-199.5T689 630q53 0 98 19q4-43 27-87q-60-21-125-21q-154 0-257.5 108.5T328 913t103.5 261T689 1280t257.5-106.5T1050 913M872 558q2 24 2 71q0 63-5 123t-20.5 132.5t-40.5 130t-68.5 106T639 1191q21 3 42 3h10q219-139 219-411q0-116-38-225m0 0q-4-80-44-171.5T730 256q92 156 142 302m335-105q0-102-51-174q-41 86-124 109q-69 19-109 53.5T883 541q0 40 24 77q74-17 140.5-67t95.5-115q-4 52-74.5 111.5T930 645q52 52 110 52q51 0 90-37t60-90q17-42 17-117m329-165v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoCamera(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 96v1088q0 42-39 59q-13 5-25 5q-27 0-45-19l-403-403v166q0 119-84.5 203.5T992 1280H288q-119 0-203.5-84.5T0 992V288Q0 169 84.5 84.5T288 0h704q119 0 203.5 84.5T1280 288v165l403-402q18-19 45-19q12 0 25 5q39 17 39 59"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vimeo(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1645 326q-10 236-332 651q-333 431-562 431q-142 0-240-263q-44-160-132-482q-72-262-157-262q-18 0-127 76l-77-98q24-21 108-96.5T256 167Q412 29 497 21q95-9 153 55.5T731 280q44 287 66 373q55 249 120 249q51 0 154-161q101-161 109-246q13-139-109-139q-57 0-121 26Q1070-11 1409 0q251 8 236 326"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VimeoSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1292 510q10-216-161-222q-231-8-312 261q44-19 82-19q85 0 74 96q-4 57-74 167T796 903q-43 0-82-169q-13-54-45-255q-30-189-160-177q-59 7-164 100l-81 72l-81 72l52 67q76-52 87-52q57 0 107 179q15 55 45 164.5t45 164.5q68 179 164 179q157 0 383-294q220-283 226-444m244-222v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vine(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1465 827v198q-101 23-198 23q-65 136-165.5 271T920 1534.5T792 1641q-80 45-162-3q-28-17-60.5-43.5t-85-83.5T382 1382.5t-107.5-184t-105.5-244T77.5 640T7 250h283q26 218 70 398.5t104.5 317T586 1201t140 195q169-169 287-406q-142-72-223-220t-81-333q0-192 104-314.5T1097 0q178 0 273 105.5t95 297.5q0 159-58 286q-7 1-19.5 3t-46 2t-63-6t-62-25.5T1166 611q31-103 31-184q0-87-29-132t-79-45q-53 0-85 49.5T972 440q0 186 105 293.5T1344 841q62 0 121-14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vk(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1949 264q23 64-150 294q-24 32-65 85q-40 51-55 72t-30.5 49.5t-12 42t13 34.5t32.5 43t57 53q4 2 5 4q141 131 191 221q3 5 6.5 12.5t7 26.5t-.5 34t-25 27.5t-59 12.5l-256 4q-24 5-56-5t-52-22l-20-12q-30-21-70-64t-68.5-77.5t-61-58t-56.5-15.5q-3 1-8 3.5t-17 14.5t-21.5 29.5t-17 52t-6.5 77.5q0 15-3.5 27.5t-7.5 18.5l-4 5q-18 19-53 22H971q-71 4-146-16.5t-131.5-53t-103-66T520 1082l-25-24q-10-10-27.5-30T396 937T290 786T167.5 575T37 303q-6-16-6-27t3-16l4-6q15-19 57-19l274-2q12 2 23 6.5t16 8.5l5 3q16 11 24 32q20 50 46 103.5t41 81.5l16 29q29 60 56 104t48.5 68.5T686 708t34 14t27-5q2-1 5-5t12-22t13.5-47t9.5-81t0-125q-2-40-9-73t-14-46l-6-12q-25-34-85-43q-13-2 5-24q16-19 38-30q53-26 239-24q82 1 135 13q20 5 33.5 13.5t20.5 24t10.5 32t3.5 45.5t-1 55t-2.5 70.5t-1.5 82.5q0 11-1 42t-.5 48t3.5 40.5t11.5 39T1189 715q8 2 17 4t26-11t38-34.5t52-67t68-107.5q60-104 107-225q4-10 10-17.5t11-10.5l4-3l5-2.5l13-3l20-.5l288-2q39-5 64 2.5t31 16.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeControlPhone(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M617 1689q0-11-13-58t-31-107t-20-69q-1-4-5-26.5t-8.5-36T526 1371q-15-14-51-14q-23 0-70 5.5t-71 5.5q-34 0-47-11q-6-5-11-15.5t-7.5-20t-6.5-24t-5-18.5q-37-128-37-255t37-255q1-4 5-18.5t6.5-24t7.5-20t11-15.5q13-11 47-11q24 0 71 5.5t70 5.5q36 0 51-14q9-8 13.5-21.5t8.5-36t5-26.5q2-9 20-69t31-107t13-58q0-22-43.5-52.5T498 264q-20-8-45-8q-34 0-98 18q-57 17-96.5 40.5t-71 66t-46 70T96 545q-6 12-9 19q-49 107-68 216T0 1024t19 244t68 216q56 122 83 161q63 91 179 127l6 2q64 18 98 18q25 0 45-8q32-12 75.5-42.5T617 1689m159-913q-26 0-45-19t-19-45.5t19-45.5q37-37 37-90q0-52-37-91q-19-19-19-45t19-45t45-19t45 19q75 75 75 181t-75 181q-21 19-45 19m181 181q-27 0-45-19q-19-19-19-45t19-45q112-114 112-272T912 304q-19-19-19-45t19-45t45-19t45 19q150 150 150 362t-150 362q-18 19-45 19m181 181q-27 0-45-19q-19-19-19-45t19-45q90-91 138.5-208t48.5-245t-48.5-245T1093 123q-19-19-19-45t19-45t45-19t45 19q109 109 167 249t58 294t-58 294t-167 249q-18 19-45 19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeDown(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 96v1088q0 26-19 45t-45 19t-45-19L326 896H64q-26 0-45-19T0 832V448q0-26 19-45t45-19h262L659 51q19-19 45-19t45 19t19 45m384 544q0 76-42.5 141.5T997 875q-10 5-25 5q-26 0-45-18.5T908 816q0-21 12-35.5t29-25t34-23t29-36t12-56.5t-12-56.5t-29-36t-34-23t-29-25t-12-35.5q0-27 19-45.5t45-18.5q15 0 25 5q70 27 112.5 93t42.5 142"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeOff(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 96v1088q0 26-19 45t-45 19t-45-19L326 896H64q-26 0-45-19T0 832V448q0-26 19-45t45-19h262L659 51q19-19 45-19t45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeUp(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 192v1088q0 26-19 45t-45 19t-45-19L326 992H64q-26 0-45-19T0 928V544q0-26 19-45t45-19h262l333-333q19-19 45-19t45 19t19 45m384 544q0 76-42.5 141.5T997 971q-10 5-25 5q-26 0-45-18.5T908 912q0-21 12-35.5t29-25t34-23t29-36t12-56.5t-12-56.5t-29-36t-34-23t-29-25t-12-35.5q0-27 19-45.5t45-18.5q15 0 25 5q70 27 112.5 93t42.5 142m256 0q0 153-85 282.5T1098 1207q-13 5-25 5q-27 0-46-19t-19-45q0-39 39-59q56-29 76-44q74-54 115.5-135.5T1280 736t-41.5-173.5T1123 427q-20-15-76-44q-39-20-39-59q0-26 19-45t45-19q13 0 26 5q140 59 225 188.5t85 282.5m256 0q0 230-127 422.5T1199 1442q-13 5-26 5q-26 0-45-19t-19-45q0-36 39-59q7-4 22.5-10.5t22.5-10.5q46-25 82-51q123-91 192-227t69-289t-69-289t-192-227q-36-26-82-51q-7-4-22.5-10.5T1148 148q-39-23-39-59q0-26 19-45t45-19q13 0 26 5q211 91 338 283.5T1664 736"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wechat(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M580 429q0-41-25-66t-66-25q-43 0-76 25.5T380 429q0 39 33 64.5t76 25.5q41 0 66-24.5t25-65.5m743 507q0-28-25.5-50t-65.5-22q-27 0-49.5 22.5T1160 936q0 28 22.5 50.5t49.5 22.5q40 0 65.5-22t25.5-51m-236-507q0-41-24.5-66T997 338q-43 0-76 25.5T888 429q0 39 33 64.5t76 25.5q41 0 65.5-24.5T1087 429m635 507q0-28-26-50t-65-22q-27 0-49.5 22.5T1559 936q0 28 22.5 50.5t49.5 22.5q39 0 65-22t26-51m-266-397q-31-4-70-4q-169 0-311 77T851.5 820.5T770 1108q0 78 23 152q-35 3-68 3q-26 0-50-1.5t-55-6.5t-44.5-7t-54.5-10.5t-50-10.5l-253 127l72-218Q0 933 0 646q0-169 97.5-311t264-223.5T725 30q176 0 332.5 66t262 182.5T1456 539m592 561q0 117-68.5 223.5T1794 1517l55 181l-199-109q-150 37-218 37q-169 0-311-70.5T897.5 1364T816 1100t81.5-264T1121 644.5t311-70.5q161 0 303 70.5t227.5 192T2048 1100"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Weibo(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M675 1124q21-34 11-69t-45-50q-34-14-73-1t-60 46q-22 34-13 68.5t43 50.5t74.5 2.5T675 1124m94-121q8-13 3.5-26.5T755 958q-14-5-28.5.5T705 977q-17 31 13 45q14 5 29-.5t22-18.5m174 107q-45 102-158 150t-224 12q-107-34-147.5-126.5T420 958q47-93 151.5-139T782 800q111 29 158.5 119.5T943 1110m312-160q-9-96-89-170T957.5 671T683 650q-223 23-369.5 141.5T181 1056q9 96 89 170t208.5 109t274.5 21q223-23 369.5-141.5T1255 950m308 4q0 68-37 139.5t-109 137t-168.5 117.5t-226 83t-270.5 31t-275-33.5t-240.5-93t-171.5-151T0 985q0-115 69.5-245T267 482q169-169 341.5-236t246.5 7q65 64 20 209q-4 14-1 20t10 7t14.5-.5T912 485l6-2q139-59 246-59t153 61q45 63 0 178q-2 13-4.5 20t4.5 12.5t12 7.5t17 6q57 18 103 47t80 81.5t34 116.5m-74-624q42 47 54.5 108.5T1537 556q-8 23-29.5 34t-44.5 4q-23-8-34-29.5t-4-44.5q20-63-24-111t-107-35q-24 5-45-8t-25-37q-5-24 8-44.5t37-25.5q60-13 119 5.5t101 65.5m181-163q87 96 112.5 222.5T1769 631q-9 27-34 40t-52 4t-40-34t-5-52q28-82 10-172t-80-158q-62-69-148-95.5t-173-8.5q-28 6-52-9.5t-30-43.5t9.5-51.5T1218 21q123-26 244 11.5T1670 167"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Whatsapp(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M985 878q13 0 97.5 44t89.5 53q2 5 2 15q0 33-17 76q-16 39-71 65.5T984 1158q-57 0-190-62q-98-45-170-118T476 793q-72-107-71-194v-8q3-91 74-158q24-22 52-22q6 0 18 1.5t19 1.5q19 0 26.5 6.5T610 448q8 20 33 88t25 75q0 21-34.5 57.5T599 715q0 7 5 15q34 73 102 137q56 53 151 101q12 7 22 7q15 0 54-48.5t52-48.5m-203 530q127 0 243.5-50t200.5-134t134-200.5t50-243.5t-50-243.5T1226 336t-200.5-134T782 152t-243.5 50T338 336T204 536.5T154 780q0 203 120 368l-79 233l242-77q158 104 345 104m0-1382q153 0 292.5 60T1315 247t161 240.5t60 292.5t-60 292.5t-161 240.5t-240.5 161t-292.5 60q-195 0-365-94L0 1574l136-405Q28 991 28 780q0-153 60-292.5T249 247T489.5 86T782 26"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wheelchair(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1023 1155l102 204q-58 179-210 290t-339 111q-156 0-288.5-77.5t-210-210T0 1184q0-181 104.5-330T379 643l17 131q-122 54-195 165.5T128 1184q0 185 131.5 316.5T576 1632q126 0 232.5-65t165-175.5T1023 1155m548 100l58 114l-256 128q-13 7-29 7q-40 0-57-35l-239-477H576q-24 0-42.5-16.5T512 935l-96-779q-2-17 6-42q14-51 57-82.5T576 0q66 0 113 47t47 113q0 69-52 117.5T564 319l37 289h423v128H617l16 128h455q40 0 57 35l228 455z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WheelchairAlt(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1374 813q34 35 29 82l-44 551q-4 42-34.5 70t-71.5 28q-6 0-9-1q-44-3-72.5-36.5T1146 1429l35-429l-143 8q55 113 55 240q0 216-148 372l-137-137q91-101 91-235q0-145-102.5-248T549 897q-134 0-236 92L176 851q120-114 284-141l264-300l-149-87l-181 161q-33 30-77 27.5T244 476t-26.5-77t34.5-73l239-213q26-23 60-26.5t64 14.5l488 283q36 21 48 68q17 67-26 117L920 801l371-20q49-3 83 32m-198-457q-74 0-126-52t-52-126t52-126t126-52t126.5 52t52.5 126t-52.5 126t-126.5 52M549 1598q106 0 196-61l139 139q-146 116-335 116q-148 0-273.5-73T77 1521T4 1248q0-188 116-336l139 139q-60 88-60 197q0 145 102.5 247.5T549 1598"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wifi(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M992 1395q-20 0-93-73.5t-73-93.5q0-32 62.5-54t103.5-22t103.5 22t62.5 54q0 20-73 93.5t-93 73.5m270-271q-2 0-40-25t-101.5-50t-128.5-25t-128.5 25t-101 50t-40.5 25q-18 0-93.5-75T553 956q0-13 10-23q78-77 196-121t233-44t233 44t196 121q10 10 10 23q0 18-75.5 93t-93.5 75m273-272q-11 0-23-8q-136-105-252-154.5T992 640q-85 0-170.5 22t-149 53T559 777t-79 53t-31 22q-17 0-92-75t-75-93q0-12 10-22q132-132 320-205t380-73t380 73t320 205q10 10 10 22q0 18-75 93t-92 75m271-271q-11 0-22-9q-179-157-371.5-236.5T992 256t-420.5 79.5T200 572q-11 9-22 9q-17 0-92.5-75T10 413q0-13 10-23q187-186 445-288T992 0t527 102t445 288q10 10 10 23q0 18-75.5 93t-92.5 75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WikipediaW(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1494 1511l-295-695q-25 49-158.5 305.5T842 1511q-1 1-27.5.5T788 1510q-82-193-255.5-587T273 327q-21-50-66.5-107.5T103 119T1 76q0-5-.5-24T0 25h583v50q-39 2-79.5 16T437 134t-10 64q26 59 216.5 499T879 1237q31-61 140-266.5T1150 723q-19-39-126-281T888 147q-38-69-201-71V26l513 1v47q-60 2-93.5 25t-12.5 69q33 70 87 189.5t86 187.5q110-214 173-363q24-55-10-79.5T1301 76q1-7 1-25V27q64 0 170.5-.5t180-1t92.5-.5v49q-62 2-119 33t-90 81l-213 442q13 33 127.5 290t121.5 274l441-1017q-14-38-49.5-62.5t-65-31.5t-55.5-8V25l460 4l1 2l-1 44q-139 4-201 145q-526 1216-559 1291z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowMaximize(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 1280h1280V512H256zM1792 160v1216q0 66-47 113t-113 47H160q-66 0-113-47T0 1376V160Q0 94 47 47T160 0h1472q66 0 113 47t47 113"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowMinimize(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 1056v192q0 66-47 113t-113 47H160q-66 0-113-47T0 1248v-192q0-66 47-113t113-47h1472q66 0 113 47t47 113"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowRestore(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 1536h768v-512H256zm1024-512h512V256h-768v256h96q66 0 113 47t47 113zm768-864v960q0 66-47 113t-113 47h-608v352q0 66-47 113t-113 47H160q-66 0-113-47T0 1632V672q0-66 47-113t113-47h608V160q0-66 47-113T928 0h960q66 0 113 47t47 113"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Windows(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M682 878v651L0 1435V878zm0-743v659H0V229zm982 743v786l-907-125V878zm0-878v794H757V125z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wordpress(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M127 896q0-163 67-313l367 1005q-196-95-315-281T127 896m1288-39q0 19-2.5 38.5t-10 49.5t-11.5 44t-17.5 59t-17.5 58l-76 256l-278-826q46-3 88-8q19-2 26-18.5t-2.5-31T1085 465l-205 10q-75-1-202-10q-12-1-20.5 5T646 485t-1.5 18.5t9 16.5t19.5 8l80 8l120 328l-168 504l-280-832q46-3 88-8q19-2 26-18.5t-2.5-31T508 465l-205 10q-7 0-23-.5t-26-.5q105-160 274.5-253.5T896 127q147 0 280.5 53T1415 329h-10q-55 0-92 40.5t-37 95.5q0 12 2 24t4 21.5t8 23t9 21t12 22.5t12.5 21t14.5 24t14 23q63 107 63 212M909 963l237 647q1 6 5 11q-126 44-255 44q-112 0-217-32zm661-436q95 174 95 369q0 209-104 385.5T1282 1560l235-678q59-169 59-276q0-42-6-79M896 0q182 0 348 71t286 191t191 286t71 348t-71 348t-191 286t-286 191t-348 71t-348-71t-286-191t-191-286T0 896t71-348t191-286T548 71T896 0m0 1751q173 0 331.5-68t273-182.5t182.5-273t68-331.5t-68-331.5t-182.5-273t-273-182.5T896 41t-331.5 68t-273 182.5t-182.5 273T41 896t68 331.5t182.5 273t273 182.5t331.5 68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wpbeginner(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 800h160V576H384zm837 332v-92q-104 36-243 38q-135 1-259.5-46.5T498 909l1 96q88 80 212 128.5t272 47.5q129 0 238-49M640 800h640V576H640zm1152-32q0 187-99 352q89 102 89 229q0 157-129.5 268T1339 1728q-122 0-225-52.5T953 1535q-19 1-57 1t-57-1q-58 88-161 140.5T453 1728q-184 0-313.5-111T10 1349q0-127 89-229Q0 955 0 768q0-209 120-385.5T446.5 103T896 0t449.5 103T1672 382.5T1792 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wpexplorer(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m948 1028l163 329h-51l-175-350l-171 350h-49l179-374l-78-33l21-49l240 102l-21 50zM563 436l304 130l-130 304l-304-130zm344 185l240 103l-103 239l-239-102zm281 150l191 81l-82 190l-190-81zm492 125q0-159-62-304t-167.5-250.5T1200 174t-304-62t-304 62t-250.5 167.5T174 592t-62 304t62 304t167.5 250.5T592 1618t304 62t304-62t250.5-167.5T1618 1200t62-304m112 0q0 182-71 348t-191 286t-286 191t-348 71t-348-71t-286-191t-191-286T0 896t71-348t191-286T548 71T896 0t348 71t286 191t191 286t71 348"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wpforms(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M515 783v128H263V783zm0-255v127H263V528zm758 511v128H932v-128zm0-256v128H601V783zm0-255v127H601V528zm135 860V148q0-8-6-14t-14-6h-32L978 384L768 213L558 384L180 128h-32q-8 0-14 6t-6 14v1240q0 8 6 14t14 6h1240q8 0 14-6t6-14M553 278l185-150H332zm430 0l221-150H798zm553-130v1240q0 62-43 105t-105 43H148q-62 0-105-43T0 1388V148Q0 86 43 43T148 0h1240q62 0 105 43t43 105"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wrench(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 1344q0-26-19-45t-45-19t-45 19t-19 45t19 45t45 19t45-19t19-45m644-420l-682 682q-37 37-90 37q-52 0-91-37L59 1498q-38-36-38-90q0-53 38-91l681-681q39 98 114.5 173.5T1028 924m634-435q0 39-23 106q-47 134-164.5 217.5T1216 896q-185 0-316.5-131.5T768 448t131.5-316.5T1216 0q58 0 121.5 16.5T1445 63q16 11 16 28t-16 28l-293 169v224l193 107q5-3 79-48.5t135.5-81T1630 454q15 0 23.5 10t8.5 25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xing(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M597 667q-10 18-257 456q-27 46-65 46H36q-21 0-31-17t0-36l253-448q1 0 0-1L97 388q-12-22-1-37q9-15 32-15h239q40 0 66 45zm806-642q11 16 0 37L875 996v1l336 615q11 20 1 37q-10 15-32 15H941q-42 0-66-45L536 997q18-32 531-942q25-45 64-45h241q22 0 31 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func XingSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M685 637q0-1-126-222q-21-34-52-34H323q-18 0-26 11q-7 12 1 29l125 216v1L227 984q-9 14 0 28q8 13 24 13h185q31 0 50-36zm624-497q-7-12-24-12h-187q-30 0-49 35L638 892q1 2 262 481q20 35 52 35h184q18 0 25-12q8-13-1-28L900 892v-1l409-723q8-16 0-28m227 148v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ycombinator(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m809 876l266-499H963L806 689q-24 48-44 92l-42-92l-155-312H445l263 493v324h101zM1536 0v1536H0V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yahoo(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m763 957l13 707q-62-11-105-11q-41 0-105 11l13-707q-40-69-168.5-295.5T194 287T13 0q58 15 108 15q44 0 111-15q63 111 133.5 229.5t167 276.5T671 733q37-61 109.5-177.5t117.5-190t105-176T1110 0q54 14 107 14q56 0 114-14q-28 39-60 88.5t-49.5 78.5t-56.5 96t-49 84Q970 595 763 957"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yelp(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M709 1319v127q-1 292-6 305q-12 32-51 40q-54 9-181.5-38T308 1664q-13-15-17-36q-1-12 4-26q4-10 34-47t181-216q1 0 60-70q15-19 39.5-24.5t49.5 3.5q24 10 37.5 29t12.5 42m-149-251q-3 55-52 70l-120 39q-275 88-292 88q-35-2-54-36q-12-25-17-75q-8-76 1-166.5T56 863t56-32q13 0 202 77q71 29 115 47l84 34q23 9 35.5 30.5T560 1068m826 297q-7 54-91.5 161T1159 1653q-37 14-63-7q-14-10-184-287l-47-77q-14-21-11.5-46t19.5-46q35-43 83-26q1 1 119 40q203 66 242 79.5t47 20.5q28 22 22 61M714 733q5 102-54 122q-58 17-114-71L168 186q-8-35 19-62q41-43 207.5-89.5T619 3q40 10 49 45q3 18 22 305.5T714 733m662 108q3 39-26 59q-15 10-329 86q-67 15-91 23l1-2q-23 6-46-4t-37-32q-30-47 0-87q1-1 75-102q125-171 150-204t34-39q28-19 65-2q48 23 123 133.5t81 167.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yoast(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M339 218h691l-26 72H339q-110 0-188.5 79T72 558v771q0 95 60.5 169.5T286 1592q23 5 98 5v72h-45q-140 0-239.5-100T0 1329V558q0-140 99.5-240T339 218M1190 0h247L955 1294q-23 61-40.5 103.5t-45 98t-54 93.5t-64.5 78.5t-79.5 65t-95.5 41t-116 18.5v-195q163-26 220-182q20-52 20-105q0-54-20-106L395 471h228l187 585zm474 558v1111H869q37-55 45-73h678V558q0-85-49.5-155T1413 304l25-67q101 34 163.5 123.5T1664 558"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Youtube(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M971 1244v211q0 67-39 67q-23 0-45-22v-301q22-22 45-22q39 0 39 67m338 1v46h-90v-46q0-68 45-68t45 68m-966-218h107v-94H138v94h105v569h100zm288 569h89v-494h-89v378q-30 42-57 42q-18 0-21-21q-1-3-1-35v-364h-89v391q0 49 8 73q12 37 58 37q48 0 102-61zm429-148v-197q0-73-9-99q-17-56-71-56q-50 0-93 54V933h-89v663h89v-48q45 55 93 55q54 0 71-55q9-27 9-100m338-10v-13h-91q0 51-2 61q-7 36-40 36q-46 0-46-69v-87h179v-103q0-79-27-116q-39-51-106-51q-68 0-107 51q-28 37-28 116v173q0 79 29 116q39 51 108 51q72 0 108-53q18-27 21-54q2-9 2-58M790 525V315q0-69-43-69t-43 69v210q0 70 43 70t43-70m719 751q0 234-26 350q-14 59-58 99t-102 46q-184 21-555 21t-555-21q-58-6-102.5-46T53 1626q-26-112-26-350q0-234 26-350q14-59 58-99t103-47q183-20 554-20t555 20q58 7 102.5 47t57.5 99q26 112 26 350M511 0h102L492 399v271H392V399q-14-74-61-212Q294 84 266 0h106l71 263zm370 333v175q0 81-28 118q-38 51-106 51q-67 0-105-51q-28-38-28-118V333q0-80 28-117q38-51 105-51q68 0 106 51q28 37 28 117m335-162v499h-91v-55q-53 62-103 62q-46 0-59-37q-8-24-8-75V171h91v367q0 33 1 35q3 22 21 22q27 0 57-43V171z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YoutubePlay(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m711 872l484-250l-484-253zM896 10q168 0 324.5 4.5T1450 24l73 4q1 0 17 1.5t23 3t23.5 4.5t28.5 8t28 13t31 19.5t29 26.5q6 6 15.5 18.5t29 58.5t26.5 101q8 64 12.5 136.5T1792 532v176q1 145-18 290q-7 55-25 99.5t-32 61.5l-14 17q-14 15-29 26.5t-31 19t-28 12.5t-28.5 8t-24 4.5t-23 3t-16.5 1.5q-251 19-627 19q-207-2-359.5-6.5T336 1256l-49-4l-36-4q-36-5-54.5-10t-51-21t-56.5-41q-6-6-15.5-18.5t-29-58.5T18 998q-8-64-12.5-136.5T0 748V572q-1-145 18-290q7-55 25-99.5T75 121l14-17q14-15 29-26.5T149 58t28-13t28.5-8t23.5-4.5t23-3t17-1.5q251-18 627-18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YoutubeSquare(children ...ElementRenderer) *FaIcon {
	return &FaIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M919 1175v-157q0-50-29-50q-17 0-33 16v224q16 16 33 16q29 0 29-49m184-122h66v-34q0-51-33-51t-33 51zM532 787v70h-80v423h-74V857h-78v-70zm201 126v367h-67v-40q-39 45-76 45q-33 0-42-28q-6-17-6-54V913h66v270q0 24 1 26q1 15 15 15q20 0 42-31V913zm252 111v146q0 52-7 73q-12 42-53 42q-35 0-68-41v36h-67V787h67v161q32-40 68-40q41 0 53 42q7 21 7 74m251 129v9q0 29-2 43q-3 22-15 40q-27 40-80 40q-52 0-81-38q-21-27-21-86v-129q0-59 20-86q29-38 80-38t78 38q21 29 21 86v76h-133v65q0 51 34 51q24 0 30-26q0-1 .5-7t.5-16.5V1153zM785 329v156q0 51-32 51t-32-51V329q0-52 32-52t32 52m533 713q0-177-19-260q-10-44-43-73.5t-76-34.5q-136-15-412-15q-275 0-411 15q-44 5-76.5 34.5T238 782q-20 87-20 260q0 176 20 260q10 43 42.5 73t75.5 35q137 15 412 15t412-15q43-5 75.5-35t42.5-73q20-84 20-260M563 391l90-296h-75l-51 195l-53-195h-78q7 23 23 69l24 69q35 103 46 158v201h74zm289 81V342q0-58-21-87q-29-38-78-38q-51 0-78 38q-21 29-21 87v130q0 58 21 87q27 38 78 38q49 0 78-38q21-27 21-87m181 120h67V222h-67v283q-22 31-42 31q-15 0-16-16q-1-2-1-26V222h-67v293q0 37 6 55q11 27 43 27q36 0 77-45zm503-304v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
