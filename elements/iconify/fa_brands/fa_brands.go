package fa_brands

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "5.15.4"
	hAttr          = "1em"
	viewbox        = "0 0 448 512"
	fill           = "currentColor"
)

type FaBrandsIcon struct {
	*SVGSVGElement
}

type FaBrandsIconFn func(children ...ElementRenderer) *FaBrandsIcon

var IconLookup = map[string]FaBrandsIconFn{
	"accessibleIcon":              AccessibleIcon,
	"accusoft":                    Accusoft,
	"acquisitionsIncorporated":    AcquisitionsIncorporated,
	"adn":                         Adn,
	"adobe":                       Adobe,
	"adversal":                    Adversal,
	"affiliatetheme":              Affiliatetheme,
	"airbnb":                      Airbnb,
	"algolia":                     Algolia,
	"alipay":                      Alipay,
	"amazon":                      Amazon,
	"amazonPay":                   AmazonPay,
	"amilia":                      Amilia,
	"android":                     Android,
	"angellist":                   Angellist,
	"angrycreative":               Angrycreative,
	"angular":                     Angular,
	"appStore":                    AppStore,
	"appStoreIos":                 AppStoreIos,
	"apper":                       Apper,
	"apple":                       Apple,
	"applePay":                    ApplePay,
	"artstation":                  Artstation,
	"asymmetrik":                  Asymmetrik,
	"atlassian":                   Atlassian,
	"audible":                     Audible,
	"autoprefixer":                Autoprefixer,
	"avianex":                     Avianex,
	"aviato":                      Aviato,
	"aws":                         Aws,
	"bandcamp":                    Bandcamp,
	"battleNet":                   BattleNet,
	"behance":                     Behance,
	"behanceSquare":               BehanceSquare,
	"bimobject":                   Bimobject,
	"bitbucket":                   Bitbucket,
	"bitcoin":                     Bitcoin,
	"bity":                        Bity,
	"blackTie":                    BlackTie,
	"blackberry":                  Blackberry,
	"blogger":                     Blogger,
	"bloggerB":                    BloggerB,
	"bluetooth":                   Bluetooth,
	"bluetoothB":                  BluetoothB,
	"bootstrap":                   Bootstrap,
	"btc":                         Btc,
	"buffer":                      Buffer,
	"buromobelexperte":            Buromobelexperte,
	"buyNlarge":                   BuyNlarge,
	"buysellads":                  Buysellads,
	"canadianMapleLeaf":           CanadianMapleLeaf,
	"ccAmazonPay":                 CcAmazonPay,
	"ccAmex":                      CcAmex,
	"ccApplePay":                  CcApplePay,
	"ccDinersClub":                CcDinersClub,
	"ccDiscover":                  CcDiscover,
	"ccJcb":                       CcJcb,
	"ccMastercard":                CcMastercard,
	"ccPaypal":                    CcPaypal,
	"ccStripe":                    CcStripe,
	"ccVisa":                      CcVisa,
	"centercode":                  Centercode,
	"centos":                      Centos,
	"chrome":                      Chrome,
	"chromecast":                  Chromecast,
	"cloudflare":                  Cloudflare,
	"cloudscale":                  Cloudscale,
	"cloudsmith":                  Cloudsmith,
	"cloudversify":                Cloudversify,
	"codepen":                     Codepen,
	"codiepie":                    Codiepie,
	"confluence":                  Confluence,
	"connectdevelop":              Connectdevelop,
	"contao":                      Contao,
	"cottonBureau":                CottonBureau,
	"cpanel":                      Cpanel,
	"creativeCommons":             CreativeCommons,
	"creativeCommonsBy":           CreativeCommonsBy,
	"creativeCommonsNc":           CreativeCommonsNc,
	"creativeCommonsNcEu":         CreativeCommonsNcEu,
	"creativeCommonsNcJp":         CreativeCommonsNcJp,
	"creativeCommonsNd":           CreativeCommonsNd,
	"creativeCommonsPd":           CreativeCommonsPd,
	"creativeCommonsPdAlt":        CreativeCommonsPdAlt,
	"creativeCommonsRemix":        CreativeCommonsRemix,
	"creativeCommonsSa":           CreativeCommonsSa,
	"creativeCommonsSampling":     CreativeCommonsSampling,
	"creativeCommonsSamplingPlus": CreativeCommonsSamplingPlus,
	"creativeCommonsShare":        CreativeCommonsShare,
	"creativeCommonsZero":         CreativeCommonsZero,
	"criticalRole":                CriticalRole,
	"cssThree":                    CssThree,
	"cssThreeAlt":                 CssThreeAlt,
	"cuttlefish":                  Cuttlefish,
	"dandD":                       DandD,
	"dandDbeyond":                 DandDbeyond,
	"dailymotion":                 Dailymotion,
	"dashcube":                    Dashcube,
	"deezer":                      Deezer,
	"delicious":                   Delicious,
	"deploydog":                   Deploydog,
	"deskpro":                     Deskpro,
	"dev":                         Dev,
	"deviantart":                  Deviantart,
	"dhl":                         Dhl,
	"diaspora":                    Diaspora,
	"digg":                        Digg,
	"digitalOcean":                DigitalOcean,
	"discord":                     Discord,
	"discourse":                   Discourse,
	"dochub":                      Dochub,
	"docker":                      Docker,
	"draftTwoDigital":             DraftTwoDigital,
	"dribbble":                    Dribbble,
	"dribbbleSquare":              DribbbleSquare,
	"dropbox":                     Dropbox,
	"drupal":                      Drupal,
	"dyalog":                      Dyalog,
	"earlybirds":                  Earlybirds,
	"ebay":                        Ebay,
	"edge":                        Edge,
	"edgeLegacy":                  EdgeLegacy,
	"elementor":                   Elementor,
	"ello":                        Ello,
	"ember":                       Ember,
	"empire":                      Empire,
	"envira":                      Envira,
	"erlang":                      Erlang,
	"ethereum":                    Ethereum,
	"etsy":                        Etsy,
	"evernote":                    Evernote,
	"expeditedssl":                Expeditedssl,
	"facebook":                    Facebook,
	"facebookF":                   FacebookF,
	"facebookMessenger":           FacebookMessenger,
	"facebookSquare":              FacebookSquare,
	"fantasyFlightGames":          FantasyFlightGames,
	"fedex":                       Fedex,
	"fedora":                      Fedora,
	"figma":                       Figma,
	"firefox":                     Firefox,
	"firefoxBrowser":              FirefoxBrowser,
	"firstOrder":                  FirstOrder,
	"firstOrderAlt":               FirstOrderAlt,
	"firstdraft":                  Firstdraft,
	"fiveHundredPx":               FiveHundredPx,
	"flickr":                      Flickr,
	"flipboard":                   Flipboard,
	"fly":                         Fly,
	"fontAwesome":                 FontAwesome,
	"fontAwesomeAlt":              FontAwesomeAlt,
	"fontAwesomeFlag":             FontAwesomeFlag,
	"fonticons":                   Fonticons,
	"fonticonsFi":                 FonticonsFi,
	"fortAwesome":                 FortAwesome,
	"fortAwesomeAlt":              FortAwesomeAlt,
	"forumbee":                    Forumbee,
	"foursquare":                  Foursquare,
	"freeCodeCamp":                FreeCodeCamp,
	"freebsd":                     Freebsd,
	"fulcrum":                     Fulcrum,
	"galacticRepublic":            GalacticRepublic,
	"galacticSenate":              GalacticSenate,
	"getPocket":                   GetPocket,
	"gg":                          Gg,
	"ggCircle":                    GgCircle,
	"git":                         Git,
	"gitAlt":                      GitAlt,
	"gitSquare":                   GitSquare,
	"github":                      Github,
	"githubAlt":                   GithubAlt,
	"githubSquare":                GithubSquare,
	"gitkraken":                   Gitkraken,
	"gitlab":                      Gitlab,
	"gitter":                      Gitter,
	"glide":                       Glide,
	"glideG":                      GlideG,
	"gofore":                      Gofore,
	"goodreads":                   Goodreads,
	"goodreadsG":                  GoodreadsG,
	"google":                      Google,
	"googleDrive":                 GoogleDrive,
	"googlePay":                   GooglePay,
	"googlePlay":                  GooglePlay,
	"googlePlus":                  GooglePlus,
	"googlePlusG":                 GooglePlusG,
	"googlePlusSquare":            GooglePlusSquare,
	"googleWallet":                GoogleWallet,
	"gratipay":                    Gratipay,
	"grav":                        Grav,
	"gripfire":                    Gripfire,
	"grunt":                       Grunt,
	"guilded":                     Guilded,
	"gulp":                        Gulp,
	"hackerNews":                  HackerNews,
	"hackerNewsSquare":            HackerNewsSquare,
	"hackerrank":                  Hackerrank,
	"hips":                        Hips,
	"hireAhelper":                 HireAhelper,
	"hive":                        Hive,
	"hooli":                       Hooli,
	"hornbill":                    Hornbill,
	"hotjar":                      Hotjar,
	"houzz":                       Houzz,
	"htmlFive":                    HtmlFive,
	"hubspot":                     Hubspot,
	"ideal":                       Ideal,
	"imdb":                        Imdb,
	"innosoft":                    Innosoft,
	"instagram":                   Instagram,
	"instagramSquare":             InstagramSquare,
	"instalod":                    Instalod,
	"intercom":                    Intercom,
	"internetExplorer":            InternetExplorer,
	"invision":                    Invision,
	"ioxhost":                     Ioxhost,
	"itchIo":                      ItchIo,
	"itunes":                      Itunes,
	"itunesNote":                  ItunesNote,
	"java":                        Java,
	"jediOrder":                   JediOrder,
	"jenkins":                     Jenkins,
	"jira":                        Jira,
	"joget":                       Joget,
	"joomla":                      Joomla,
	"js":                          Js,
	"jsSquare":                    JsSquare,
	"jsfiddle":                    Jsfiddle,
	"kaggle":                      Kaggle,
	"keybase":                     Keybase,
	"keycdn":                      Keycdn,
	"kickstarter":                 Kickstarter,
	"kickstarterK":                KickstarterK,
	"korvue":                      Korvue,
	"laravel":                     Laravel,
	"lastfm":                      Lastfm,
	"lastfmSquare":                LastfmSquare,
	"leanpub":                     Leanpub,
	"less":                        Less,
	"line":                        Line,
	"linkedin":                    Linkedin,
	"linkedinIn":                  LinkedinIn,
	"linode":                      Linode,
	"linux":                       Linux,
	"lyft":                        Lyft,
	"magento":                     Magento,
	"mailchimp":                   Mailchimp,
	"mandalorian":                 Mandalorian,
	"markdown":                    Markdown,
	"mastodon":                    Mastodon,
	"maxcdn":                      Maxcdn,
	"mdb":                         Mdb,
	"medapps":                     Medapps,
	"medium":                      Medium,
	"mediumM":                     MediumM,
	"medrt":                       Medrt,
	"meetup":                      Meetup,
	"megaport":                    Megaport,
	"mendeley":                    Mendeley,
	"microblog":                   Microblog,
	"microsoft":                   Microsoft,
	"mix":                         Mix,
	"mixcloud":                    Mixcloud,
	"mixer":                       Mixer,
	"mizuni":                      Mizuni,
	"modx":                        Modx,
	"monero":                      Monero,
	"napster":                     Napster,
	"neos":                        Neos,
	"nimblr":                      Nimblr,
	"nintendoSwitch":              NintendoSwitch,
	"node":                        Node,
	"nodeJs":                      NodeJs,
	"npm":                         Npm,
	"nsEight":                     NsEight,
	"nutritionix":                 Nutritionix,
	"octopusDeploy":               OctopusDeploy,
	"odnoklassniki":               Odnoklassniki,
	"odnoklassnikiSquare":         OdnoklassnikiSquare,
	"oldRepublic":                 OldRepublic,
	"opencart":                    Opencart,
	"openid":                      Openid,
	"opera":                       Opera,
	"optinMonster":                OptinMonster,
	"orcid":                       Orcid,
	"osi":                         Osi,
	"pageFour":                    PageFour,
	"pagelines":                   Pagelines,
	"palfed":                      Palfed,
	"patreon":                     Patreon,
	"paypal":                      Paypal,
	"pennyArcade":                 PennyArcade,
	"perbyte":                     Perbyte,
	"periscope":                   Periscope,
	"phabricator":                 Phabricator,
	"phoenixFramework":            PhoenixFramework,
	"phoenixSquadron":             PhoenixSquadron,
	"php":                         Php,
	"piedPiper":                   PiedPiper,
	"piedPiperAlt":                PiedPiperAlt,
	"piedPiperHat":                PiedPiperHat,
	"piedPiperPp":                 PiedPiperPp,
	"piedPiperSquare":             PiedPiperSquare,
	"pinterest":                   Pinterest,
	"pinterestP":                  PinterestP,
	"pinterestSquare":             PinterestSquare,
	"playstation":                 Playstation,
	"productHunt":                 ProductHunt,
	"pushed":                      Pushed,
	"python":                      Python,
	"qq":                          Qq,
	"quinscape":                   Quinscape,
	"quora":                       Quora,
	"rproject":                    Rproject,
	"raspberryPi":                 RaspberryPi,
	"ravelry":                     Ravelry,
	"react":                       React,
	"reacteurope":                 Reacteurope,
	"readme":                      Readme,
	"rebel":                       Rebel,
	"redRiver":                    RedRiver,
	"reddit":                      Reddit,
	"redditAlien":                 RedditAlien,
	"redditSquare":                RedditSquare,
	"redhat":                      Redhat,
	"rendact":                     Rendact,
	"renren":                      Renren,
	"replyd":                      Replyd,
	"researchgate":                Researchgate,
	"resolving":                   Resolving,
	"rev":                         Rev,
	"rocketchat":                  Rocketchat,
	"rockrms":                     Rockrms,
	"rust":                        Rust,
	"safari":                      Safari,
	"salesforce":                  Salesforce,
	"sass":                        Sass,
	"schlix":                      Schlix,
	"scribd":                      Scribd,
	"searchengin":                 Searchengin,
	"sellcast":                    Sellcast,
	"sellsy":                      Sellsy,
	"servicestack":                Servicestack,
	"shirtsinbulk":                Shirtsinbulk,
	"shopify":                     Shopify,
	"shopware":                    Shopware,
	"simplybuilt":                 Simplybuilt,
	"sistrix":                     Sistrix,
	"sith":                        Sith,
	"sketch":                      Sketch,
	"skyatlas":                    Skyatlas,
	"skype":                       Skype,
	"slack":                       Slack,
	"slackHash":                   SlackHash,
	"slideshare":                  Slideshare,
	"snapchat":                    Snapchat,
	"snapchatGhost":               SnapchatGhost,
	"snapchatSquare":              SnapchatSquare,
	"soundcloud":                  Soundcloud,
	"sourcetree":                  Sourcetree,
	"speakap":                     Speakap,
	"speakerDeck":                 SpeakerDeck,
	"spotify":                     Spotify,
	"squarespace":                 Squarespace,
	"stackExchange":               StackExchange,
	"stackOverflow":               StackOverflow,
	"stackpath":                   Stackpath,
	"staylinked":                  Staylinked,
	"steam":                       Steam,
	"steamSquare":                 SteamSquare,
	"steamSymbol":                 SteamSymbol,
	"stickerMule":                 StickerMule,
	"strava":                      Strava,
	"stripe":                      Stripe,
	"stripeS":                     StripeS,
	"studiovinari":                Studiovinari,
	"stumbleupon":                 Stumbleupon,
	"stumbleuponCircle":           StumbleuponCircle,
	"superpowers":                 Superpowers,
	"supple":                      Supple,
	"suse":                        Suse,
	"swift":                       Swift,
	"symfony":                     Symfony,
	"teamspeak":                   Teamspeak,
	"telegram":                    Telegram,
	"telegramPlane":               TelegramPlane,
	"tencentWeibo":                TencentWeibo,
	"theRedYeti":                  TheRedYeti,
	"themeco":                     Themeco,
	"themeisle":                   Themeisle,
	"thinkPeaks":                  ThinkPeaks,
	"tiktok":                      Tiktok,
	"tradeFederation":             TradeFederation,
	"trello":                      Trello,
	"tripadvisor":                 Tripadvisor,
	"tumblr":                      Tumblr,
	"tumblrSquare":                TumblrSquare,
	"twitch":                      Twitch,
	"twitter":                     Twitter,
	"twitterSquare":               TwitterSquare,
	"typoThree":                   TypoThree,
	"uber":                        Uber,
	"ubuntu":                      Ubuntu,
	"uikit":                       Uikit,
	"umbraco":                     Umbraco,
	"uncharted":                   Uncharted,
	"uniregistry":                 Uniregistry,
	"unity":                       Unity,
	"unsplash":                    Unsplash,
	"untappd":                     Untappd,
	"ups":                         Ups,
	"usb":                         Usb,
	"usps":                        Usps,
	"ussunnah":                    Ussunnah,
	"vaadin":                      Vaadin,
	"viacoin":                     Viacoin,
	"viadeo":                      Viadeo,
	"viadeoSquare":                ViadeoSquare,
	"viber":                       Viber,
	"vimeo":                       Vimeo,
	"vimeoSquare":                 VimeoSquare,
	"vimeoV":                      VimeoV,
	"vine":                        Vine,
	"vk":                          Vk,
	"vnv":                         Vnv,
	"vuejs":                       Vuejs,
	"watchmanMonitoring":          WatchmanMonitoring,
	"waze":                        Waze,
	"weebly":                      Weebly,
	"weibo":                       Weibo,
	"weixin":                      Weixin,
	"whatsapp":                    Whatsapp,
	"whatsappSquare":              WhatsappSquare,
	"whmcs":                       Whmcs,
	"wikipediaW":                  WikipediaW,
	"windows":                     Windows,
	"wix":                         Wix,
	"wizardsOfTheCoast":           WizardsOfTheCoast,
	"wodu":                        Wodu,
	"wolfPackBattalion":           WolfPackBattalion,
	"wordpress":                   Wordpress,
	"wordpressSimple":             WordpressSimple,
	"wpbeginner":                  Wpbeginner,
	"wpexplorer":                  Wpexplorer,
	"wpforms":                     Wpforms,
	"wpressr":                     Wpressr,
	"xbox":                        Xbox,
	"xing":                        Xing,
	"xingSquare":                  XingSquare,
	"ycombinator":                 Ycombinator,
	"yahoo":                       Yahoo,
	"yammer":                      Yammer,
	"yandex":                      Yandex,
	"yandexInternational":         YandexInternational,
	"yarn":                        Yarn,
	"yelp":                        Yelp,
	"yoast":                       Yoast,
	"youtube":                     Youtube,
	"youtubeSquare":               YoutubeSquare,
	"zhihu":                       Zhihu,
}

func AccessibleIcon(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M423.9 255.8L411 413.1c-3.3 40.7-63.9 35.1-60.6-4.9l10-122.5l-41.1 2.3c10.1 20.7 15.8 43.9 15.8 68.5c0 41.2-16.1 78.7-42.3 106.5l-39.3-39.3c57.9-63.7 13.1-167.2-74-167.2c-25.9 0-49.5 9.9-67.2 26L73 243.2c22-20.7 50.1-35.1 81.4-40.2l75.3-85.7l-42.6-24.8l-51.6 46c-30 26.8-70.6-18.5-40.5-45.4l68-60.7c9.8-8.8 24.1-10.2 35.5-3.6c0 0 139.3 80.9 139.5 81.1c16.2 10.1 20.7 36 6.1 52.6L285.7 229l106.1-5.9c18.5-1.1 33.6 14.4 32.1 32.7m-64.9-154c28.1 0 50.9-22.8 50.9-50.9C409.9 22.8 387.1 0 359 0c-28.1 0-50.9 22.8-50.9 50.9c0 28.1 22.8 50.9 50.9 50.9M179.6 456.5c-80.6 0-127.4-90.6-82.7-156.1l-39.7-39.7C36.4 287 24 320.3 24 356.4c0 130.7 150.7 201.4 251.4 122.5l-39.7-39.7c-16 10.9-35.3 17.3-56.1 17.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Accusoft(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M322.1 252v-1l-51.2-65.8s-12 1.6-25 15.1c-9 9.3-242.1 239.1-243.4 240.9c-7 10 1.6 6.8 15.7 1.7c.8 0 114.5-36.6 114.5-36.6c.5-.6-.1-.1.6-.6c-.4-5.1-.8-26.2-1-27.7c-.6-5.2 2.2-6.9 7-8.9l92.6-33.8c.6-.8 88.5-81.7 90.2-83.3m160.1 120.1c13.3 16.1 20.7 13.3 30.8 9.3c3.2-1.2 115.4-47.6 117.8-48.9c8-4.3-1.7-16.7-7.2-23.4c-2.1-2.5-205.1-245.6-207.2-248.3c-9.7-12.2-14.3-12.9-38.4-12.8c-10.2 0-106.8.5-116.5.6c-19.2.1-32.9-.3-19.2 16.9C250 75 476.5 365.2 482.2 372.1m152.7 1.6c-2.3-.3-24.6-4.7-38-7.2c0 0-115 50.4-117.5 51.6c-16 7.3-26.9-3.2-36.7-14.6l-57.1-74c-5.4-.9-60.4-9.6-65.3-9.3c-3.1.2-9.6.8-14.4 2.9c-4.9 2.1-145.2 52.8-150.2 54.7c-5.1 2-11.4 3.6-11.1 7.6c.2 2.5 2 2.6 4.6 3.5c2.7.8 300.9 67.6 308 69.1c15.6 3.3 38.5 10.5 53.6 1.7c2.1-1.2 123.8-76.4 125.8-77.8c5.4-4 4.3-6.8-1.7-8.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AcquisitionsIncorporated(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M357.45 468.2c-1.2-7.7-1.3-7.6-9.6-7.6c-99.8.2-111.8-2.4-112.7-2.6c-12.3-1.7-20.6-10.5-21-23.1c-.1-1.6-.2-71.6-1-129.1c-.1-4.7 1.6-6.4 5.9-7.5c12.5-3 24.9-6.1 37.3-9.7c4.3-1.3 6.8-.2 8.4 3.5c4.5 10.3 8.8 20.6 13.2 30.9c1.6 3.7.1 4.4-3.4 4.4c-10-.2-20-.1-30.4-.1v27h116c-1.4-9.5-2.7-18.1-4-27.5c-7 0-13.8.4-20.4-.1c-22.6-1.6-18.3-4.4-84-158.6c-8.8-20.1-27.9-62.1-36.5-89.2c-4.4-14 5.5-25.4 18.9-26.6c18.6-1.7 37.5-1.6 56.2-2c20.6-.4 41.2-.4 61.8-.5c3.1 0 4-1.4 4.3-4.3c1.2-9.8 2.7-19.5 4-29.2c.8-5.3 1.6-10.7 2.4-16.1L23.75 0c-3.6 0-5.3 1.1-4.6 5.3c2.2 13.2-.8.8 6.4 45.3c63.4 0 71.8.9 101.8.5c12.3-.2 37 3.5 37.7 22.1c.4 11.4-1.1 11.3-32.6 87.4c-53.8 129.8-50.7 120.3-67.3 161c-1.7 4.1-3.6 5.2-7.6 5.2c-8.5-.2-17-.3-25.4.1c-1.9.1-5.2 1.8-5.5 3.2c-1.5 8.1-2.2 16.3-3.2 24.9h114.3v-27.6c-6.9 0-33.5.4-35.3-2.9c5.3-12.3 10.4-24.4 15.7-36.7c16.3 4 31.9 7.8 47.6 11.7c3.4.9 4.6 3 4.6 6.8c-.1 42.9.1 85.9.2 128.8c0 10.2-5.5 19.1-14.9 23.1c-6.5 2.7-3.3 3.4-121.4 2.4c-5.3 0-7.1 2-7.6 6.8c-1.5 12.9-2.9 25.9-5 38.8c-.8 5 1.3 5.7 5.3 5.7c183.2.6-30.7 0 337.1 0c-2.5-15-4.4-29.4-6.6-43.7m-174.9-205.7c-13.3-4.2-26.6-8.2-39.9-12.5a44.53 44.53 0 0 1-5.8-2.9c17.2-44.3 34.2-88.1 51.3-132.1c7.5 2.4 7.9-.8 9.4 0c9.3 22.5 18.1 60.1 27 82.8c6.6 16.7 13 33.5 19.7 50.9a35.78 35.78 0 0 1-3.9 2.1c-13.1 3.9-26.4 7.5-39.4 11.7a27.66 27.66 0 0 1-18.4 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Adn(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m248 167.5l64.9 98.8H183.1zM496 256c0 136.9-111.1 248-248 248S0 392.9 0 256S111.1 8 248 8s248 111.1 248 248m-99.8 82.7L248 115.5L99.8 338.7h30.4l33.6-51.7h168.6l33.6 51.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Adobe(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M315.5 64h170.9v384L315.5 64zm-119 0H25.6v384L196.5 64zM256 206.1L363.5 448h-73l-30.7-76.8h-78.7L256 206.1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Adversal(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M482.1 32H28.7C5.8 32 0 37.9 0 60.9v390.2C0 474.4 5.8 480 28.7 480h453.4c24.4 0 29.9-5.2 29.9-29.7V62.2c0-24.6-5.4-30.2-29.9-30.2M178.4 220.3c-27.5-20.2-72.1-8.7-84.2 23.4c-4.3 11.1-9.3 9.5-17.5 8.3c-9.7-1.5-17.2-3.2-22.5-5.5c-28.8-11.4 8.6-55.3 24.9-64.3c41.1-21.4 83.4-22.2 125.3-4.8c40.9 16.8 34.5 59.2 34.5 128.5c2.7 25.8-4.3 58.3 9.3 88.8c1.9 4.4.4 7.9-2.7 10.7c-8.4 6.7-39.3 2.2-46.6-7.4c-1.9-2.2-1.8-3.6-3.9-6.2c-3.6-3.9-7.3-2.2-11.9 1c-57.4 36.4-140.3 21.4-147-43.3c-3.1-29.3 12.4-57.1 39.6-71c38.2-19.5 112.2-11.8 114-30.9c1.1-10.2-1.9-20.1-11.3-27.3m286.7 222c0 15.1-11.1 9.9-17.8 9.9H52.4c-7.4 0-18.2 4.8-17.8-10.7c.4-13.9 10.5-9.1 17.1-9.1c132.3-.4 264.5-.4 396.8 0c6.8 0 16.6-4.4 16.6 9.9m3.8-340.5v291c0 5.7-.7 13.9-8.1 13.9c-12.4-.4-27.5 7.1-36.1-5.6c-5.8-8.7-7.8-4-12.4-1.2c-53.4 29.7-128.1 7.1-144.4-85.2c-6.1-33.4-.7-67.1 15.7-100c11.8-23.9 56.9-76.1 136.1-30.5v-71c0-26.2-.1-26.2 26-26.2c3.1 0 6.6.4 9.7 0c10.1-.8 13.6 4.4 13.6 14.3c-.1.2-.1.3-.1.5m-51.5 232.3c-19.5 47.6-72.9 43.3-90 5.2c-15.1-33.3-15.5-68.2.4-101.5c16.3-34.1 59.7-35.7 81.5-4.8c20.6 28.8 14.9 84.6 8.1 101.1m-294.8 35.3c-7.5-1.3-33-3.3-33.7-27.8c-.4-13.9 7.8-23 19.8-25.8c24.4-5.9 49.3-9.9 73.7-14.7c8.9-2 7.4 4.4 7.8 9.5c1.4 33-26.1 59.2-67.6 58.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Affiliatetheme(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M159.7 237.4C108.4 308.3 43.1 348.2 14 326.6C-15.2 304.9 2.8 230 54.2 159.1c51.3-70.9 116.6-110.8 145.7-89.2c29.1 21.6 11.1 96.6-40.2 167.5m351.2-57.3C437.1 303.5 319 367.8 246.4 323.7c-25-15.2-41.3-41.2-49-73.8c-33.6 64.8-92.8 113.8-164.1 133.2c49.8 59.3 124.1 96.9 207 96.9c150 0 271.6-123.1 271.6-274.9c.1-8.5-.3-16.8-1-25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Airbnb(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M224 373.12c-25.24-31.67-40.08-59.43-45-83.18c-22.55-88 112.61-88 90.06 0c-5.45 24.25-20.29 52-45 83.18zm138.15 73.23c-42.06 18.31-83.67-10.88-119.3-50.47c103.9-130.07 46.11-200-18.85-200c-54.92 0-85.16 46.51-73.28 100.5c6.93 29.19 25.23 62.39 54.43 99.5c-32.53 36.05-60.55 52.69-85.15 54.92c-50 7.43-89.11-41.06-71.3-91.09c15.1-39.16 111.72-231.18 115.87-241.56c15.75-30.07 25.56-57.4 59.38-57.4c32.34 0 43.4 25.94 60.37 59.87c36 70.62 89.35 177.48 114.84 239.09c13.17 33.07-1.37 71.29-37.01 86.64m47-136.12C280.27 35.93 273.13 32 224 32c-45.52 0-64.87 31.67-84.66 72.79C33.18 317.1 22.89 347.19 22 349.81C-3.22 419.14 48.74 480 111.63 480c21.71 0 60.61-6.06 112.37-62.4c58.68 63.78 101.26 62.4 112.37 62.4c62.89.05 114.85-60.86 89.61-130.19c.02-3.89-16.82-38.9-16.82-39.58z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Algolia(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M229.3 182.6c-49.3 0-89.2 39.9-89.2 89.2c0 49.3 39.9 89.2 89.2 89.2s89.2-39.9 89.2-89.2c0-49.3-40-89.2-89.2-89.2m62.7 56.6l-58.9 30.6c-1.8.9-3.8-.4-3.8-2.3V201c0-1.5 1.3-2.7 2.7-2.6c26.2 1 48.9 15.7 61.1 37.1c.7 1.3.2 3-1.1 3.7M389.1 32H58.9C26.4 32 0 58.4 0 90.9V421c0 32.6 26.4 59 58.9 59H389c32.6 0 58.9-26.4 58.9-58.9V90.9C448 58.4 421.6 32 389.1 32m-202.6 84.7c0-10.8 8.7-19.5 19.5-19.5h45.3c10.8 0 19.5 8.7 19.5 19.5v15.4c0 1.8-1.7 3-3.3 2.5c-12.3-3.4-25.1-5.1-38.1-5.1c-13.5 0-26.7 1.8-39.4 5.5c-1.7.5-3.4-.8-3.4-2.5v-15.8zm-84.4 37l9.2-9.2c7.6-7.6 19.9-7.6 27.5 0l7.7 7.7c1.1 1.1 1 3-.3 4c-6.2 4.5-12.1 9.4-17.6 14.9c-5.4 5.4-10.4 11.3-14.8 17.4c-1 1.3-2.9 1.5-4 .3l-7.7-7.7c-7.6-7.5-7.6-19.8 0-27.4m127.2 244.8c-70 0-126.6-56.7-126.6-126.6s56.7-126.6 126.6-126.6c70 0 126.6 56.6 126.6 126.6c0 69.8-56.7 126.6-126.6 126.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alipay(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M377.74 32H70.26C31.41 32 0 63.41 0 102.26v307.48C0 448.59 31.41 480 70.26 480h307.48c38.52 0 69.76-31.08 70.26-69.6c-45.96-25.62-110.59-60.34-171.6-88.44c-32.07 43.97-84.14 81-148.62 81c-70.59 0-93.73-45.3-97.04-76.37c-3.97-39.01 14.88-81.5 99.52-81.5c35.38 0 79.35 10.25 127.13 24.96c16.53-30.09 26.45-60.34 26.45-60.34h-178.2v-16.7h92.08v-31.24H88.28v-19.01h109.44V92.34h50.92v50.42h109.44v19.01H248.63v31.24h88.77s-15.21 46.62-38.35 90.92c48.93 16.7 100.01 36.04 148.62 52.74V102.26C447.83 63.57 416.43 32 377.74 32M47.28 322.95c.99 20.17 10.25 53.73 69.93 53.73c52.07 0 92.58-39.68 117.87-72.9c-44.63-18.68-84.48-31.41-109.44-31.41c-67.45 0-79.35 33.06-78.36 50.58"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Amazon(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M257.2 162.7c-48.7 1.8-169.5 15.5-169.5 117.5c0 109.5 138.3 114 183.5 43.2c6.5 10.2 35.4 37.5 45.3 46.8l56.8-56S341 288.9 341 261.4V114.3C341 89 316.5 32 228.7 32C140.7 32 94 87 94 136.3l73.5 6.8c16.3-49.5 54.2-49.5 54.2-49.5c40.7-.1 35.5 29.8 35.5 69.1m0 86.8c0 80-84.2 68-84.2 17.2c0-47.2 50.5-56.7 84.2-57.8zm136 163.5c-7.7 10-70 67-174.5 67S34.2 408.5 9.7 379c-6.8-7.7 1-11.3 5.5-8.3C88.5 415.2 203 488.5 387.7 401c7.5-3.7 13.3 2 5.5 12m39.8 2.2c-6.5 15.8-16 26.8-21.2 31c-5.5 4.5-9.5 2.7-6.5-3.8s19.3-46.5 12.7-55c-6.5-8.3-37-4.3-48-3.2c-10.8 1-13 2-14-.3c-2.3-5.7 21.7-15.5 37.5-17.5c15.7-1.8 41-.8 46 5.7c3.7 5.1 0 27.1-6.5 43.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AmazonPay(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 325.3c2.3-4.2 5.2-4.9 9.7-2.5c10.4 5.6 20.6 11.4 31.2 16.7a595.88 595.88 0 0 0 127.4 46.3a616.61 616.61 0 0 0 63.2 11.8a603.33 603.33 0 0 0 95 5.2c17.4-.4 34.8-1.8 52.1-3.8a603.66 603.66 0 0 0 163.3-42.8c2.9-1.2 5.9-2 9.1-1.2c6.7 1.8 9 9 4.1 13.9a70 70 0 0 1-9.6 7.4c-30.7 21.1-64.2 36.4-99.6 47.9a473.31 473.31 0 0 1-75.1 17.6a431 431 0 0 1-53.2 4.8a21.3 21.3 0 0 0-2.5.3H308a21.3 21.3 0 0 0-2.5-.3c-3.6-.2-7.2-.3-10.7-.4a426.3 426.3 0 0 1-50.4-5.3A448.4 448.4 0 0 1 164 420a443.33 443.33 0 0 1-145.6-87c-1.8-1.6-3-3.8-4.4-5.7zM172 65.1l-4.3.6a80.92 80.92 0 0 0-38 15.1c-2.4 1.7-4.6 3.5-7.1 5.4a4.29 4.29 0 0 1-.4-1.4c-.4-2.7-.8-5.5-1.3-8.2c-.7-4.6-3-6.6-7.6-6.6h-11.5c-6.9 0-8.2 1.3-8.2 8.2v209.3c0 1 0 2 .1 3c.2 3 2 4.9 4.9 5c7 .1 14.1.1 21.1 0c2.9 0 4.7-2 5-5c.1-1 .1-2 .1-3v-72.4c1.1.9 1.7 1.4 2.2 1.9c17.9 14.9 38.5 19.8 61 15.4c20.4-4 34.6-16.5 43.8-34.9c7-13.9 9.9-28.7 10.3-44.1c.5-17.1-1.2-33.9-8.1-49.8c-8.5-19.6-22.6-32.5-43.9-36.9c-3.2-.7-6.5-1-9.8-1.5c-2.8-.1-5.5-.1-8.3-.1M124.6 107a3.48 3.48 0 0 1 1.7-3.3c13.7-9.5 28.8-14.5 45.6-13.2c14.9 1.1 27.1 8.4 33.5 25.9c3.9 10.7 4.9 21.8 4.9 33c0 10.4-.8 20.6-4 30.6c-6.8 21.3-22.4 29.4-42.6 28.5c-14-.6-26.2-6-37.4-13.9a3.57 3.57 0 0 1-1.7-3.3c.1-14.1 0-28.1 0-42.2s.1-28 0-42.1m205.7-41.9c-1 .1-2 .3-2.9.4a148 148 0 0 0-28.9 4.1c-6.1 1.6-12 3.8-17.9 5.8c-3.6 1.2-5.4 3.8-5.3 7.7c.1 3.3-.1 6.6 0 9.9c.1 4.8 2.1 6.1 6.8 4.9c7.8-2 15.6-4.2 23.5-5.7c12.3-2.3 24.7-3.3 37.2-1.4c6.5 1 12.6 2.9 16.8 8.4c3.7 4.8 5.1 10.5 5.3 16.4c.3 8.3.2 16.6.3 24.9a7.84 7.84 0 0 1-.2 1.4c-.5-.1-.9 0-1.3-.1a180.56 180.56 0 0 0-32-4.9c-11.3-.6-22.5.1-33.3 3.9c-12.9 4.5-23.3 12.3-29.4 24.9c-4.7 9.8-5.4 20.2-3.9 30.7c2 14 9 24.8 21.4 31.7c11.9 6.6 24.8 7.4 37.9 5.4c15.1-2.3 28.5-8.7 40.3-18.4a7.36 7.36 0 0 1 1.6-1.1c.6 3.8 1.1 7.4 1.8 11c.6 3.1 2.5 5.1 5.4 5.2c5.4.1 10.9.1 16.3 0a4.84 4.84 0 0 0 4.8-4.7a26.2 26.2 0 0 0 .1-2.8v-106a80 80 0 0 0-.9-12.9c-1.9-12.9-7.4-23.5-19-30.4c-6.7-4-14.1-6-21.8-7.1c-3.6-.5-7.2-.8-10.8-1.3c-3.9.1-7.9.1-11.9.1m35 127.7a3.33 3.33 0 0 1-1.5 3c-11.2 8.1-23.5 13.5-37.4 14.9c-5.7.6-11.4.4-16.8-1.8a20.08 20.08 0 0 1-12.4-13.3a32.9 32.9 0 0 1-.1-19.4c2.5-8.3 8.4-13 16.4-15.6a61.33 61.33 0 0 1 24.8-2.2c8.4.7 16.6 2.3 25 3.4c1.6.2 2.1 1 2.1 2.6c-.1 4.8 0 9.5 0 14.3s-.2 9.4-.1 14.1m259.9 129.4c-1-5-4.8-6.9-9.1-8.3a88.42 88.42 0 0 0-21-3.9a147.32 147.32 0 0 0-39.2 1.9c-14.3 2.7-27.9 7.3-40 15.6a13.75 13.75 0 0 0-3.7 3.5a5.11 5.11 0 0 0-.5 4c.4 1.5 2.1 1.9 3.6 1.8a16.2 16.2 0 0 0 2.2-.1c7.8-.8 15.5-1.7 23.3-2.5c11.4-1.1 22.9-1.8 34.3-.9a71.64 71.64 0 0 1 14.4 2.7c5.1 1.4 7.4 5.2 7.6 10.4c.4 8-1.4 15.7-3.5 23.3c-4.1 15.4-10 30.3-15.8 45.1a17.6 17.6 0 0 0-1 3c-.5 2.9 1.2 4.8 4.1 4.1a10.56 10.56 0 0 0 4.8-2.5a145.91 145.91 0 0 0 12.7-13.4c12.8-16.4 20.3-35.3 24.7-55.6c.8-3.6 1.4-7.3 2.1-10.9zM493.1 199q-19.35-53.55-38.7-107.2c-2-5.7-4.2-11.3-6.3-16.9c-1.1-2.9-3.2-4.8-6.4-4.8c-7.6-.1-15.2-.2-22.9-.1c-2.5 0-3.7 2-3.2 4.5a43.1 43.1 0 0 0 1.9 6.1q29.4 72.75 59.1 145.5c1.7 4.1 2.1 7.6.2 11.8c-3.3 7.3-5.9 15-9.3 22.3c-3 6.5-8 11.4-15.2 13.3a42.13 42.13 0 0 1-15.4 1.1c-2.5-.2-5-.8-7.5-1c-3.4-.2-5.1 1.3-5.2 4.8q-.15 5 0 9.9c.1 5.5 2 8 7.4 8.9a108.18 108.18 0 0 0 16.9 2c17.1.4 30.7-6.5 39.5-21.4a131.63 131.63 0 0 0 9.2-18.4q35.55-89.7 70.6-179.6a26.62 26.62 0 0 0 1.6-5.5c.4-2.8-.9-4.4-3.7-4.4c-6.6-.1-13.3 0-19.9 0a7.54 7.54 0 0 0-7.7 5.2c-.5 1.4-1.1 2.7-1.6 4.1l-34.8 100c-2.5 7.2-5.1 14.5-7.7 22.2c-.4-1.1-.6-1.7-.9-2.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Amilia(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M240.1 32c-61.9 0-131.5 16.9-184.2 55.4c-5.1 3.1-9.1 9.2-7.2 19.4c1.1 5.1 5.1 27.4 10.2 39.6c4.1 10.2 14.2 10.2 20.3 6.1c32.5-22.3 96.5-47.7 152.3-47.7c57.9 0 58.9 28.4 58.9 73.1v38.5C203 227.7 78.2 251 46.7 264.2C11.2 280.5 16.3 357.7 16.3 376s15.2 104 124.9 104c47.8 0 113.7-20.7 153.3-42.1v25.4c0 3 2.1 8.2 6.1 9.1c3.1 1 50.7 2 59.9 2s62.5.3 66.5-.7c4.1-1 5.1-6.1 5.1-9.1V168c-.1-80.3-57.9-136-192-136m50.2 348c-21.4 13.2-48.7 24.4-79.1 24.4c-52.8 0-58.9-33.5-59-44.7c0-12.2-3-42.7 18.3-52.9c24.3-13.2 75.1-29.4 119.8-33.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Android(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M420.55 301.93a24 24 0 1 1 24-24a24 24 0 0 1-24 24m-265.1 0a24 24 0 1 1 24-24a24 24 0 0 1-24 24m273.7-144.48l47.94-83a10 10 0 1 0-17.27-10l-48.54 84.07a301.25 301.25 0 0 0-246.56 0l-48.54-84.07a10 10 0 1 0-17.27 10l47.94 83C64.53 202.22 8.24 285.55 0 384h576c-8.24-98.45-64.54-181.78-146.85-226.55"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Angellist(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M347.1 215.4c11.7-32.6 45.4-126.9 45.4-157.1c0-26.6-15.7-48.9-43.7-48.9c-44.6 0-84.6 131.7-97.1 163.1C242 144 196.6 0 156.6 0c-31.1 0-45.7 22.9-45.7 51.7c0 35.3 34.2 126.8 46.6 162c-6.3-2.3-13.1-4.3-20-4.3c-23.4 0-48.3 29.1-48.3 52.6c0 8.9 4.9 21.4 8 29.7c-36.9 10-51.1 34.6-51.1 71.7C46 435.6 114.4 512 210.6 512c118 0 191.4-88.6 191.4-202.9c0-43.1-6.9-82-54.9-93.7M311.7 108c4-12.3 21.1-64.3 37.1-64.3c8.6 0 10.9 8.9 10.9 16c0 19.1-38.6 124.6-47.1 148l-34-6zM142.3 48.3c0-11.9 14.5-45.7 46.3 47.1l34.6 100.3c-15.6-1.3-27.7-3-35.4 1.4c-10.9-28.8-45.5-119.7-45.5-148.8M140 244c29.3 0 67.1 94.6 67.1 107.4c0 5.1-4.9 11.4-10.6 11.4c-20.9 0-76.9-76.9-76.9-97.7c.1-7.7 12.7-21.1 20.4-21.1m184.3 186.3c-29.1 32-66.3 48.6-109.7 48.6c-59.4 0-106.3-32.6-128.9-88.3c-17.1-43.4 3.8-68.3 20.6-68.3c11.4 0 54.3 60.3 54.3 73.1c0 4.9-7.7 8.3-11.7 8.3c-16.1 0-22.4-15.5-51.1-51.4c-29.7 29.7 20.5 86.9 58.3 86.9c26.1 0 43.1-24.2 38-42c3.7 0 8.3.3 11.7-.6c1.1 27.1 9.1 59.4 41.7 61.7c0-.9 2-7.1 2-7.4c0-17.4-10.6-32.6-10.6-50.3c0-28.3 21.7-55.7 43.7-71.7c8-6 17.7-9.7 27.1-13.1c9.7-3.7 20-8 27.4-15.4c-1.1-11.2-5.7-21.1-16.9-21.1c-27.7 0-120.6 4-120.6-39.7c0-6.7.1-13.1 17.4-13.1c32.3 0 114.3 8 138.3 29.1c18.1 16.1 24.3 113.2-31 174.7m-98.6-126c9.7 3.1 19.7 4 29.7 6c-7.4 5.4-14 12-20.3 19.1c-2.8-8.5-6.2-16.8-9.4-25.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Angrycreative(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m640 238.2l-3.2 28.2l-34.5 2.3l-2 18.1l34.5-2.3l-3.2 28.2l-34.4 2.2l-2.3 20.1l34.4-2.2l-3 26.1l-64.7 4.1l12.7-113.2L527 365.2l-31.9 2l-23.8-117.8l30.3-2l13.6 79.4l31.7-82.4zM426.8 371.5l28.3-1.8L468 249.6l-28.4 1.9zM162 388.1l-19.4-36l-3.5 37.4l-28.2 1.7l2.7-29.1c-11 18-32 34.3-56.9 35.8C23.9 399.9-3 377 .3 339.7c2.6-29.3 26.7-62.8 67.5-65.4c37.7-2.4 47.6 23.2 51.3 28.8l2.8-30.8l38.9-2.5c20.1-1.3 38.7 3.7 42.5 23.7l2.6-26.6l64.8-4.2l-2.7 27.9l-36.4 2.4l-1.7 17.9l36.4-2.3l-2.7 27.9l-36.4 2.3l-1.9 19.9l36.3-2.3l-2.1 20.8l55-117.2l23.8-1.6L370.4 369l8.9-85.6l-22.3 1.4l2.9-27.9l75-4.9l-3 28l-24.3 1.6l-9.7 91.9l-58 3.7l-4.3-15.6l-39.4 2.5l-8 16.3zm-44.3-70.2l-26.4 1.7C84.6 307.2 76.9 303 65 303.8c-19 1.2-33.3 17.5-34.6 33.3c-1.4 16 7.3 32.5 28.7 31.2c12.8-.8 21.3-8.6 28.9-18.9l27-1.7zm56.1-7.7c1.2-12.9-7.6-13.6-26.1-12.4l-2.7 28.5c14.2-.9 27.5-2.1 28.8-16.1m21.1 70.8l5.8-60c-5 13.5-14.7 21.1-27.9 26.6zm135.4-45l-7.9-37.8l-15.8 39.3zm-170.1-74.6l-4.3-17.5l-39.6 2.6l-8.1 18.2l-31.9 2.1l57-121.9l23.9-1.6l30.7 102l9.9-104.7l27-1.8l37.8 63.6l6.5-66.6l28.5-1.9l-4 41.2c7.4-13.5 22.9-44.7 63.6-47.5c40.5-2.8 52.4 29.3 53.4 30.3l3.3-32l39.3-2.7c12.7-.9 27.8.3 36.3 9.7l-4.4-11.9l32.2-2.2l12.9 43.2l23-45.7l31-2.2l-43.6 78.4l-4.8 44.3l-28.4 1.9l4.8-44.3l-15.8-43c1 22.3-9.2 40.1-32 49.6l25.2 38.8l-36.4 2.4l-19.2-36.8l-4 38.3l-28.4 1.9l3.3-31.5c-6.7 9.3-19.7 35.4-59.6 38c-26.2 1.7-45.6-10.3-55.4-39.2l-4 40.3l-25 1.6l-37.6-63.3l-6.3 66.2zm276.6-82.1c10.2-.7 17.5-2.1 21.6-4.3c4.5-2.4 7-6.4 7.6-12.1c.6-5.3-.6-8.8-3.4-10.4c-3.6-2.1-10.6-2.8-22.9-2zM327.7 214c5.6 5.9 12.7 8.5 21.3 7.9c4.7-.3 9.1-1.8 13.3-4.1c5.5-3 10.6-8 15.1-14.3l-34.2 2.3l2.4-23.9l63.1-4.3l1.2-12l-31.2 2.1c-4.1-3.7-7.8-6.6-11.1-8.1c-4-1.7-8.1-2.8-12.2-2.5c-8 .5-15.3 3.6-22 9.2c-7.7 6.4-12 14.5-12.9 24.4c-1.1 9.6 1.4 17.3 7.2 23.3m-201.3 8.2l23.8-1.6l-8.3-37.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Angular(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M185.7 268.1h76.2l-38.1-91.6zM223.8 32L16 106.4l31.8 275.7l176 97.9l176-97.9l31.8-275.7zM354 373.8h-48.6l-26.2-65.4H168.6l-26.2 65.4H93.7L223.8 81.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppStore(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m255.9 120.9l9.1-15.7c5.6-9.8 18.1-13.1 27.9-7.5c9.8 5.6 13.1 18.1 7.5 27.9l-87.5 151.5h63.3c20.5 0 32 24.1 23.1 40.8H113.8c-11.3 0-20.4-9.1-20.4-20.4c0-11.3 9.1-20.4 20.4-20.4h52l66.6-115.4l-20.8-36.1c-5.6-9.8-2.3-22.2 7.5-27.9c9.8-5.6 22.2-2.3 27.9 7.5zm-78.7 218l-19.6 34c-5.6 9.8-18.1 13.1-27.9 7.5c-9.8-5.6-13.1-18.1-7.5-27.9l14.6-25.2c16.4-5.1 29.8-1.2 40.4 11.6m168.9-61.7h53.1c11.3 0 20.4 9.1 20.4 20.4c0 11.3-9.1 20.4-20.4 20.4h-29.5l19.9 34.5c5.6 9.8 2.3 22.2-7.5 27.9c-9.8 5.6-22.2 2.3-27.9-7.5c-33.5-58.1-58.7-101.6-75.4-130.6c-17.1-29.5-4.9-59.1 7.2-69.1c13.4 23 33.4 57.7 60.1 104M256 8C119 8 8 119 8 256s111 248 248 248s248-111 248-248S393 8 256 8m216 248c0 118.7-96.1 216-216 216c-118.7 0-216-96.1-216-216c0-118.7 96.1-216 216-216c118.7 0 216 96.1 216 216"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppStoreIos(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 32H48C21.5 32 0 53.5 0 80v352c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V80c0-26.5-21.5-48-48-48M127 384.5c-5.5 9.6-17.8 12.8-27.3 7.3c-9.6-5.5-12.8-17.8-7.3-27.3l14.3-24.7c16.1-4.9 29.3-1.1 39.6 11.4zm138.9-53.9H84c-11 0-20-9-20-20s9-20 20-20h51l65.4-113.2l-20.5-35.4c-5.5-9.6-2.2-21.8 7.3-27.3c9.6-5.5 21.8-2.2 27.3 7.3l8.9 15.4l8.9-15.4c5.5-9.6 17.8-12.8 27.3-7.3c9.6 5.5 12.8 17.8 7.3 27.3l-85.8 148.6h62.1c20.2 0 31.5 23.7 22.7 40m98.1 0h-29l19.6 33.9c5.5 9.6 2.2 21.8-7.3 27.3c-9.6 5.5-21.8 2.2-27.3-7.3c-32.9-56.9-57.5-99.7-74-128.1c-16.7-29-4.8-58 7.1-67.8c13.1 22.7 32.7 56.7 58.9 102h52c11 0 20 9 20 20c0 11.1-9 20-20 20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Apper(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M42.1 239.1c22.2 0 29 2.8 33.5 14.6h.8v-22.9c0-11.3-4.8-15.4-17.9-15.4c-11.3 0-14.4 2.5-15.1 12.8H4.8c.3-13.9 1.5-19.1 5.8-24.4C17.9 195 29.5 192 56.7 192c33 0 47.1 5 53.9 18.9c2 4.3 4 15.6 4 23.7v76.3H76.3l1.3-19.1h-1c-5.3 15.6-13.6 20.4-35.5 20.4c-30.3 0-41.1-10.1-41.1-37.3c0-25.2 12.3-35.8 42.1-35.8m17.1 48.1c13.1 0 16.9-3 16.9-13.4c0-9.1-4.3-11.6-19.6-11.6c-13.1 0-17.9 3-17.9 12.1c-.1 10.4 3.7 12.9 20.6 12.9m77.8-94.9h38.3l-1.5 20.6h.8c9.1-17.1 15.9-20.9 37.5-20.9c14.4 0 24.7 3 31.5 9.1c9.8 8.6 12.8 20.4 12.8 48.1c0 30-3 43.1-12.1 52.9c-6.8 7.3-16.4 10.1-33.2 10.1c-20.4 0-29.2-5.5-33.8-21.2h-.8v70.3H137zm80.9 60.7c0-27.5-3.3-32.5-20.7-32.5c-16.9 0-20.7 5-20.7 28.7c0 28 3.5 33.5 21.2 33.5c16.4 0 20.2-5.6 20.2-29.7m57.9-60.7h38.3l-1.5 20.6h.8c9.1-17.1 15.9-20.9 37.5-20.9c14.4 0 24.7 3 31.5 9.1c9.8 8.6 12.8 20.4 12.8 48.1c0 30-3 43.1-12.1 52.9c-6.8 7.3-16.4 10.1-33.3 10.1c-20.4 0-29.2-5.5-33.8-21.2h-.8v70.3h-39.5v-169zm80.9 60.7c0-27.5-3.3-32.5-20.7-32.5c-16.9 0-20.7 5-20.7 28.7c0 28 3.5 33.5 21.2 33.5c16.4 0 20.2-5.6 20.2-29.7m53.8-3.8c0-25.4 3.3-37.8 12.3-45.8c8.8-8.1 22.2-11.3 45.1-11.3c42.8 0 55.7 12.8 55.7 55.7v11.1h-75.3c-.3 2-.3 4-.3 4.8c0 16.9 4.5 21.9 20.1 21.9c13.9 0 17.9-3 17.9-13.9h37.5v2.3c0 9.8-2.5 18.9-6.8 24.7c-7.3 9.8-19.6 13.6-44.3 13.6c-27.5 0-41.6-3.3-50.6-12.3c-8.5-8.5-11.3-21.3-11.3-50.8m76.4-11.6c-.3-1.8-.3-3.3-.3-3.8c0-12.3-3.3-14.6-19.6-14.6c-14.4 0-17.1 3-18.1 15.1l-.3 3.3zm55.6-45.3h38.3l-1.8 19.9h.7c6.8-14.9 14.4-20.2 29.7-20.2c10.8 0 19.1 3.3 23.4 9.3c5.3 7.3 6.8 14.4 6.8 34c0 1.5 0 5 .2 9.3h-35c.3-1.8.3-3.3.3-4c0-15.4-2-19.4-10.3-19.4c-6.3 0-10.8 3.3-13.1 9.3c-1 3-1 4.3-1 12.3v68h-38.3V192.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Apple(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M318.7 268.7c-.2-36.7 16.4-64.4 50-84.8c-18.8-26.9-47.2-41.7-84.7-44.6c-35.5-2.8-74.3 20.7-88.5 20.7c-15 0-49.4-19.7-76.4-19.7C63.3 141.2 4 184.8 4 273.5q0 39.3 14.4 81.2c12.8 36.7 59 126.7 107.2 125.2c25.2-.6 43-17.9 75.8-17.9c31.8 0 48.3 17.9 76.4 17.9c48.6-.7 90.4-82.5 102.6-119.3c-65.2-30.7-61.7-90-61.7-91.9m-56.6-164.2c27.3-32.4 24.8-61.9 24-72.5c-24.1 1.4-52 16.4-67.9 34.9c-17.5 19.8-27.8 44.3-25.6 71.9c26.1 2 49.9-11.4 69.5-34.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ApplePay(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M116.9 158.5c-7.5 8.9-19.5 15.9-31.5 14.9c-1.5-12 4.4-24.8 11.3-32.6c7.5-9.1 20.6-15.6 31.3-16.1c1.2 12.4-3.7 24.7-11.1 33.8m10.9 17.2c-17.4-1-32.3 9.9-40.5 9.9c-8.4 0-21-9.4-34.8-9.1c-17.9.3-34.5 10.4-43.6 26.5c-18.8 32.3-4.9 80 13.3 106.3c8.9 13 19.5 27.3 33.5 26.8c13.3-.5 18.5-8.6 34.5-8.6c16.1 0 20.8 8.6 34.8 8.4c14.5-.3 23.6-13 32.5-26c10.1-14.8 14.3-29.1 14.5-29.9c-.3-.3-28-10.9-28.3-42.9c-.3-26.8 21.9-39.5 22.9-40.3c-12.5-18.6-32-20.6-38.8-21.1m100.4-36.2v194.9h30.3v-66.6h41.9c38.3 0 65.1-26.3 65.1-64.3s-26.4-64-64.1-64zm30.3 25.5h34.9c26.3 0 41.3 14 41.3 38.6s-15 38.8-41.4 38.8h-34.8zm162.2 170.9c19 0 36.6-9.6 44.6-24.9h.6v23.4h28v-97c0-28.1-22.5-46.3-57.1-46.3c-32.1 0-55.9 18.4-56.8 43.6h27.3c2.3-12 13.4-19.9 28.6-19.9c18.5 0 28.9 8.6 28.9 24.5v10.8l-37.8 2.3c-35.1 2.1-54.1 16.5-54.1 41.5c.1 25.2 19.7 42 47.8 42m8.2-23.1c-16.1 0-26.4-7.8-26.4-19.6c0-12.3 9.9-19.4 28.8-20.5l33.6-2.1v11c0 18.2-15.5 31.2-36 31.2m102.5 74.6c29.5 0 43.4-11.3 55.5-45.4L640 193h-30.8l-35.6 115.1h-.6L537.4 193h-31.6L557 334.9l-2.8 8.6c-4.6 14.6-12.1 20.3-25.5 20.3c-2.4 0-7-.3-8.9-.5v23.4c1.8.4 9.3.7 11.6.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Artstation(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2 377.4l43 74.3A51.35 51.35 0 0 0 90.9 480h285.4l-59.2-102.6zM501.8 350L335.6 59.3A51.38 51.38 0 0 0 290.2 32h-88.4l257.3 447.6l40.7-70.5c1.9-3.2 21-29.7 2-59.1M275 304.5l-115.5-200L44 304.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Asymmetrik(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M517.5 309.2c38.8-40 58.1-80 58.5-116.1c.8-65.5-59.4-118.2-169.4-135C277.9 38.4 118.1 73.6 0 140.5C52 114 110.6 92.3 170.7 82.3c74.5-20.5 153-25.4 221.3-14.8C544.5 91.3 588.8 195 490.8 299.2c-10.2 10.8-22 21.1-35 30.6L304.9 103.4L114.7 388.9c-65.6-29.4-76.5-90.2-19.1-151.2c20.8-22.2 48.3-41.9 79.5-58.1c20-12.2 39.7-22.6 62-30.7c-65.1 20.3-122.7 52.9-161.6 92.9c-27.7 28.6-41.4 57.1-41.7 82.9c-.5 35.1 23.4 65.1 68.4 83l-34.5 51.7h101.6l22-34.4c22.2 1 45.3 0 68.6-2.7l-22.8 37.1h135.5L340 406.3c18.6-5.3 36.9-11.5 54.5-18.7l45.9 71.8H542L468.6 349c18.5-12.1 35-25.5 48.9-39.8m-187.6 80.5l-25-40.6l-32.7 53.3c-23.4 3.5-46.7 5.1-69.2 4.4l101.9-159.3l78.7 123c-17.2 7.4-35.3 13.9-53.7 19.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Atlassian(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M152.2 236.4c-7.7-8.2-19.7-7.7-24.8 2.8L1.6 490.2c-5 10 2.4 21.7 13.4 21.7h175c5.8.1 11-3.2 13.4-8.4c37.9-77.8 15.1-196.3-51.2-267.1M244.4 8.1c-122.3 193.4-8.5 348.6 65 495.5c2.5 5.1 7.7 8.4 13.4 8.4H497c11.2 0 18.4-11.8 13.4-21.7c0 0-234.5-470.6-240.4-482.3c-5.3-10.6-18.8-10.8-25.6.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Audible(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 199.9v54l-320 200L0 254v-54l320 200zm-194.5 72l47.1-29.4c-37.2-55.8-100.7-92.6-172.7-92.6c-72 0-135.5 36.7-172.6 92.4h.3c2.5-2.3 5.1-4.5 7.7-6.7c89.7-74.4 219.4-58.1 290.2 36.3m-220.1 18.8c16.9-11.9 36.5-18.7 57.4-18.7c34.4 0 65.2 18.4 86.4 47.6l45.4-28.4c-20.9-29.9-55.6-49.5-94.8-49.5c-38.9 0-73.4 19.4-94.4 49M103.6 161.1c131.8-104.3 318.2-76.4 417.5 62.1l.7 1l48.8-30.4C517.1 112.1 424.8 58.1 319.9 58.1c-103.5 0-196.6 53.5-250.5 135.6c9.9-10.5 22.7-23.5 34.2-32.6m467 32.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Autoprefixer(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m318.4 16l-161 480h77.5l25.4-81.4h119.5L405 496h77.5zm-40.3 341.9l41.2-130.4h1.5l40.9 130.4zM640 405l-10-31.4L462.1 358l19.4 56.5zm-462.1-47L10 373.7L0 405l158.5 9.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Avianex(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M453.1 32h-312c-38.9 0-76.2 31.2-83.3 69.7L1.2 410.3C-5.9 448.8 19.9 480 58.9 480h312c38.9 0 76.2-31.2 83.3-69.7l56.7-308.5c7-38.6-18.8-69.8-57.8-69.8m-58.2 347.3l-32 13.5l-115.4-110c-14.7 10-29.2 19.5-41.7 27.1l22.1 64.2l-17.9 12.7l-40.6-61l-52.4-48.1l15.7-15.4l58 31.1c9.3-10.5 20.8-22.6 32.8-34.9L203 228.9l-68.8-99.8l18.8-28.9l8.9-4.8L265 207.8l4.9 4.5c19.4-18.8 33.8-32.4 33.8-32.4c7.7-6.5 21.5-2.9 30.7 7.9c9 10.5 10.6 24.7 2.7 31.3c-1.8 1.3-15.5 11.4-35.3 25.6l4.5 7.3l94.9 119.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Aviato(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m107.2 283.5l-19-41.8H36.1l-19 41.8H0l62.2-131.4l62.2 131.4zm-45-98.1l-19.6 42.5h39.2zm112.7 102.4l-62.2-131.4h17.1l45.1 96l45.1-96h17zm80.6-4.3V156.4H271v127.1zm209.1-115.6v115.6h-17.3V167.9h-41.2v-11.5h99.6v11.5zM640 218.8c0 9.2-1.7 17.8-5.1 25.8c-3.4 8-8.2 15.1-14.2 21.1c-6 6-13.1 10.8-21.1 14.2c-8 3.4-16.6 5.1-25.8 5.1s-17.8-1.7-25.8-5.1c-8-3.4-15.1-8.2-21.1-14.2c-6-6-10.8-13-14.2-21.1c-3.4-8-5.1-16.6-5.1-25.8s1.7-17.8 5.1-25.8c3.4-8 8.2-15.1 14.2-21.1c6-6 13-8.4 21.1-11.9c8-3.4 16.6-5.1 25.8-5.1s17.8 1.7 25.8 5.1c8 3.4 15.1 5.8 21.1 11.9c6 6 10.7 13.1 14.2 21.1c3.4 8 5.1 16.6 5.1 25.8m-15.5 0c0-7.3-1.3-14-3.9-20.3c-2.6-6.3-6.2-11.7-10.8-16.3c-4.6-4.6-10-8.2-16.2-10.9c-6.2-2.7-12.8-4-19.8-4s-13.6 1.3-19.8 4c-6.2 2.7-11.6 6.3-16.2 10.9c-4.6 4.6-8.2 10-10.8 16.3c-2.6 6.3-3.9 13.1-3.9 20.3c0 7.3 1.3 14 3.9 20.3c2.6 6.3 6.2 11.7 10.8 16.3c4.6 4.6 10 8.2 16.2 10.9c6.2 2.7 12.8 4 19.8 4s13.6-1.3 19.8-4c6.2-2.7 11.6-6.3 16.2-10.9c4.6-4.6 8.2-10 10.8-16.3c2.6-6.3 3.9-13.1 3.9-20.3m-94.8 96.7v-6.3l88.9-10l-242.9 13.4c.6-2.2 1.1-4.6 1.4-7.2c.3-2 .5-4.2.6-6.5l64.8-8.1l-64.9 1.9c0-.4-.1-.7-.1-1.1c-2.8-17.2-25.5-23.7-25.5-23.7l-1.1-26.3h23.8l19 41.8h17.1L348.6 152l-62.2 131.4h17.1l19-41.8h23.6L345 268s-22.7 6.5-25.5 23.7c-.1.3-.1.7-.1 1.1l-64.9-1.9l64.8 8.1c.1 2.3.3 4.4.6 6.5c.3 2.6.8 5 1.4 7.2L78.4 299.2l88.9 10v6.3c-5.9.9-10.5 6-10.5 12.2c0 6.8 5.6 12.4 12.4 12.4c6.8 0 12.4-5.6 12.4-12.4c0-6.2-4.6-11.3-10.5-12.2v-5.8l80.3 9v5.4c-5.7 1.1-9.9 6.2-9.9 12.1c0 6.8 5.6 10.2 12.4 10.2c6.8 0 12.4-3.4 12.4-10.2c0-6-4.3-11-9.9-12.1v-4.9l28.4 3.2v23.7h-5.9V360h5.9v-6.6h5v6.6h5.9v-13.8h-5.9V323l38.3 4.3c8.1 11.4 19 13.6 19 13.6l-.1 6.7l-5.1.2l-.1 12.1h4.1l.1-5h5.2l.1 5h4.1l-.1-12.1l-5.1-.2l-.1-6.7s10.9-2.1 19-13.6l38.3-4.3v23.2h-5.9V360h5.9v-6.6h5v6.6h5.9v-13.8h-5.9v-23.7l28.4-3.2v4.9c-5.7 1.1-9.9 6.2-9.9 12.1c0 6.8 5.6 10.2 12.4 10.2c6.8 0 12.4-3.4 12.4-10.2c0-6-4.3-11-9.9-12.1v-5.4l80.3-9v5.8c-5.9.9-10.5 6-10.5 12.2c0 6.8 5.6 12.4 12.4 12.4c6.8 0 12.4-5.6 12.4-12.4c-.2-6.3-4.7-11.4-10.7-12.3m-200.8-87.6l19.6-42.5l19.6 42.5h-17.9l-1.7-40.3l-1.7 40.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Aws(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M180.41 203.01c-.72 22.65 10.6 32.68 10.88 39.05a8.164 8.164 0 0 1-4.1 6.27l-12.8 8.96a10.66 10.66 0 0 1-5.63 1.92c-.43-.02-8.19 1.83-20.48-25.61a78.608 78.608 0 0 1-62.61 29.45c-16.28.89-60.4-9.24-58.13-56.21c-1.59-38.28 34.06-62.06 70.93-60.05c7.1.02 21.6.37 46.99 6.27v-15.62c2.69-26.46-14.7-46.99-44.81-43.91c-2.4.01-19.4-.5-45.84 10.11c-7.36 3.38-8.3 2.82-10.75 2.82c-7.41 0-4.36-21.48-2.94-24.2c5.21-6.4 35.86-18.35 65.94-18.18a76.857 76.857 0 0 1 55.69 17.28a70.285 70.285 0 0 1 17.67 52.36zM93.99 235.4c32.43-.47 46.16-19.97 49.29-30.47c2.46-10.05 2.05-16.41 2.05-27.4c-9.67-2.32-23.59-4.85-39.56-4.87c-15.15-1.14-42.82 5.63-41.74 32.26c-1.24 16.79 11.12 31.4 29.96 30.48m170.92 23.05c-7.86.72-11.52-4.86-12.68-10.37l-49.8-164.65c-.97-2.78-1.61-5.65-1.92-8.58a4.61 4.61 0 0 1 3.86-5.25c.24-.04-2.13 0 22.25 0c8.78-.88 11.64 6.03 12.55 10.37l35.72 140.83l33.16-140.83c.53-3.22 2.94-11.07 12.8-10.24h17.16c2.17-.18 11.11-.5 12.68 10.37l33.42 142.63L420.98 80.1c.48-2.18 2.72-11.37 12.68-10.37h19.72c.85-.13 6.15-.81 5.25 8.58c-.43 1.85 3.41-10.66-52.75 169.9c-1.15 5.51-4.82 11.09-12.68 10.37h-18.69c-10.94 1.15-12.51-9.66-12.68-10.75L328.67 110.7l-32.78 136.99c-.16 1.09-1.73 11.9-12.68 10.75h-18.3zm273.48 5.63c-5.88.01-33.92-.3-57.36-12.29a12.802 12.802 0 0 1-7.81-11.91v-10.75c0-8.45 6.2-6.9 8.83-5.89c10.04 4.06 16.48 7.14 28.81 9.6c36.65 7.53 52.77-2.3 56.72-4.48c13.15-7.81 14.19-25.68 5.25-34.95c-10.48-8.79-15.48-9.12-53.13-21c-4.64-1.29-43.7-13.61-43.79-52.36c-.61-28.24 25.05-56.18 69.52-55.95c12.67-.01 46.43 4.13 55.57 15.62c1.35 2.09 2.02 4.55 1.92 7.04v10.11c0 4.44-1.62 6.66-4.87 6.66c-7.71-.86-21.39-11.17-49.16-10.75c-6.89-.36-39.89.91-38.41 24.97c-.43 18.96 26.61 26.07 29.7 26.89c36.46 10.97 48.65 12.79 63.12 29.58c17.14 22.25 7.9 48.3 4.35 55.44c-19.08 37.49-68.42 34.44-69.26 34.42m40.2 104.86c-70.03 51.72-171.69 79.25-258.49 79.25A469.127 469.127 0 0 1 2.83 327.46c-6.53-5.89-.77-13.96 7.17-9.47a637.37 637.37 0 0 0 316.88 84.12a630.22 630.22 0 0 0 241.59-49.55c11.78-5 21.77 7.8 10.12 16.38m29.19-33.29c-8.96-11.52-59.28-5.38-81.81-2.69c-6.79.77-7.94-5.12-1.79-9.47c40.07-28.17 105.88-20.1 113.44-10.63c7.55 9.47-2.05 75.41-39.56 106.91c-5.76 4.87-11.27 2.3-8.71-4.1c8.44-21.25 27.39-68.49 18.43-80.02"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bandcamp(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 8C119 8 8 119 8 256s111 248 248 248s248-111 248-248S393 8 256 8m48.2 326.1h-181L207.9 178h181Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BattleNet(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448.61 225.62c26.87.18 35.57-7.43 38.92-12.37c12.47-16.32-7.06-47.6-52.85-71.33c17.76-33.58 30.11-63.68 36.34-85.3c3.38-11.83 1.09-19 .45-20.25c-1.72 10.52-15.85 48.46-48.2 100.05c-25-11.22-56.52-20.1-93.77-23.8c-8.94-16.94-34.88-63.86-60.48-88.93C252.18 7.14 238.7 1.07 228.18.22h-.05c-13.83-1.55-22.67 5.85-27.4 11c-17.2 18.53-24.33 48.87-25 84.07c-7.24-12.35-17.17-24.63-28.5-25.93h-.18c-20.66-3.48-38.39 29.22-36 81.29c-38.36 1.38-71 5.75-93 11.23c-9.9 2.45-16.22 7.27-17.76 9.72c1-.38 22.4-9.22 111.56-9.22c5.22 53 29.75 101.82 26 93.19c-9.73 15.4-38.24 62.36-47.31 97.7c-5.87 22.88-4.37 37.61.15 47.14c5.57 12.75 16.41 16.72 23.2 18.26c25 5.71 55.38-3.63 86.7-21.14c-7.53 12.84-13.9 28.51-9.06 39.34c7.31 19.65 44.49 18.66 88.44-9.45c20.18 32.18 40.07 57.94 55.7 74.12a39.79 39.79 0 0 0 8.75 7.09c5.14 3.21 8.58 3.37 8.58 3.37c-8.24-6.75-34-38-62.54-91.78c22.22-16 45.65-38.87 67.47-69.27c122.82 4.6 143.29-24.76 148-31.64c14.67-19.88 3.43-57.44-57.32-93.69m-77.85 106.22c23.81-37.71 30.34-67.77 29.45-92.33c27.86 17.57 47.18 37.58 49.06 58.83c1.14 12.93-8.1 29.12-78.51 33.5M216.9 387.69c9.76-6.23 19.53-13.12 29.2-20.49c6.68 13.33 13.6 26.1 20.6 38.19c-40.6 21.86-68.84 12.76-49.8-17.7m215-171.35c-10.29-5.34-21.16-10.34-32.38-15.05a722.459 722.459 0 0 0 22.74-36.9c39.06 24.1 45.9 53.18 9.64 51.95M279.18 398c-5.51-11.35-11-23.5-16.5-36.44c43.25 1.27 62.42-18.73 63.28-20.41c0 .07-25 15.64-62.53 12.25a718.78 718.78 0 0 0 85.06-84q13.06-15.31 24.93-31.11c-.36-.29-1.54-3-16.51-12c-51.7 60.27-102.34 98-132.75 115.92c-20.59-11.18-40.84-31.78-55.71-61.49c-20-39.92-30-82.39-31.57-116.07c12.3.91 25.27 2.17 38.85 3.88c-22.29 36.8-14.39 63-13.47 64.23c0-.07-.95-29.17 20.14-59.57a695.23 695.23 0 0 0 44.67 152.84c.93-.38 1.84.88 18.67-8.25c-26.33-74.47-33.76-138.17-34-173.43c20-12.42 48.18-19.8 81.63-17.81c44.57 2.67 86.36 15.25 116.32 30.71q-10.69 15.66-23.33 32.47C365.63 152 339.1 145.84 337.5 146c.11 0 25.9 14.07 41.52 47.22a717.63 717.63 0 0 0-115.34-31.71a646.608 646.608 0 0 0-39.39-6.05c-.07.45-1.81 1.85-2.16 20.33C300 190.28 358.78 215.68 389.36 233c.74 23.55-6.95 51.61-25.41 79.57c-24.6 37.31-56.39 67.23-84.77 85.43m27.4-287c-44.56-1.66-73.58 7.43-94.69 20.67c2-52.3 21.31-76.38 38.21-75.28C267 52.15 305 108.55 306.58 111m-130.65 3.1c.48 12.11 1.59 24.62 3.21 37.28c-14.55-.85-28.74-1.25-42.4-1.26c-.08 3.24-.12-51 24.67-49.59h.09c5.76 1.09 10.63 6.88 14.43 13.57m-28.06 162c20.76 39.7 43.3 60.57 65.25 72.31c-46.79 24.76-77.53 20-84.92 4.51c-.2-.21-11.13-15.3 19.67-76.81zm210.06 74.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Behance(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M232 237.2c31.8-15.2 48.4-38.2 48.4-74c0-70.6-52.6-87.8-113.3-87.8H0v354.4h171.8c64.4 0 124.9-30.9 124.9-102.9c0-44.5-21.1-77.4-64.7-89.7M77.9 135.9H151c28.1 0 53.4 7.9 53.4 40.5c0 30.1-19.7 42.2-47.5 42.2h-79zm83.3 233.7H77.9V272h84.9c34.3 0 56 14.3 56 50.6c0 35.8-25.9 47-57.6 47m358.5-240.7H376V94h143.7zM576 305.2c0-75.9-44.4-139.2-124.9-139.2c-78.2 0-131.3 58.8-131.3 135.8c0 79.9 50.3 134.7 131.3 134.7c61.3 0 101-27.6 120.1-86.3H509c-6.7 21.9-34.3 33.5-55.7 33.5c-41.3 0-63-24.2-63-65.3h185.1c.3-4.2.6-8.7.6-13.2M390.4 274c2.3-33.7 24.7-54.8 58.5-54.8c35.4 0 53.2 20.8 56.2 54.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BehanceSquare(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M186.5 293c0 19.3-14 25.4-31.2 25.4h-45.1v-52.9h46c18.6.1 30.3 7.8 30.3 27.5m-7.7-82.3c0-17.7-13.7-21.9-28.9-21.9h-39.6v44.8H153c15.1 0 25.8-6.6 25.8-22.9m132.3 23.2c-18.3 0-30.5 11.4-31.7 29.7h62.2c-1.7-18.5-11.3-29.7-30.5-29.7M448 80v352c0 26.5-21.5 48-48 48H48c-26.5 0-48-21.5-48-48V80c0-26.5 21.5-48 48-48h352c26.5 0 48 21.5 48 48M271.7 185h77.8v-18.9h-77.8zm-43 110.3c0-24.1-11.4-44.9-35-51.6c17.2-8.2 26.2-17.7 26.2-37c0-38.2-28.5-47.5-61.4-47.5H68v192h93.1c34.9-.2 67.6-16.9 67.6-55.9M380 280.5c0-41.1-24.1-75.4-67.6-75.4c-42.4 0-71.1 31.8-71.1 73.6c0 43.3 27.3 73 71.1 73c33.2 0 54.7-14.9 65.1-46.8h-33.7c-3.7 11.9-18.6 18.1-30.2 18.1c-22.4 0-34.1-13.1-34.1-35.3h100.2c.1-2.3.3-4.8.3-7.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bimobject(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M416 32H32C14.4 32 0 46.4 0 64v384c0 17.6 14.4 32 32 32h384c17.6 0 32-14.4 32-32V64c0-17.6-14.4-32-32-32m-64 257.4c0 49.4-11.4 82.6-103.8 82.6h-16.9c-44.1 0-62.4-14.9-70.4-38.8h-.9V368H96V136h64v74.7h1.1c4.6-30.5 39.7-38.8 69.7-38.8h17.3c92.4 0 103.8 33.1 103.8 82.5v35zm-64-28.9v22.9c0 21.7-3.4 33.8-38.4 33.8h-45.3c-28.9 0-44.1-6.5-44.1-35.7v-19c0-29.3 15.2-35.7 44.1-35.7h45.3c35-.2 38.4 12 38.4 33.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bitbucket(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.2 32A16 16 0 0 0 6 47.8a26.35 26.35 0 0 0 .2 2.8l67.9 412.1a21.77 21.77 0 0 0 21.3 18.2h325.7a16 16 0 0 0 16-13.4L505 50.7a16 16 0 0 0-13.2-18.3a24.58 24.58 0 0 0-2.8-.2zm285.9 297.8h-104l-28.1-147h157.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bitcoin(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M504 256c0 136.967-111.033 248-248 248S8 392.967 8 256S119.033 8 256 8s248 111.033 248 248m-141.651-35.33c4.937-32.999-20.191-50.739-54.55-62.573l11.146-44.702l-27.213-6.781l-10.851 43.524c-7.154-1.783-14.502-3.464-21.803-5.13l10.929-43.81l-27.198-6.781l-11.153 44.686c-5.922-1.349-11.735-2.682-17.377-4.084l.031-.14l-37.53-9.37l-7.239 29.062s20.191 4.627 19.765 4.913c11.022 2.751 13.014 10.044 12.68 15.825l-12.696 50.925c.76.194 1.744.473 2.829.907c-.907-.225-1.876-.473-2.876-.713l-17.796 71.338c-1.349 3.348-4.767 8.37-12.471 6.464c.271.395-19.78-4.937-19.78-4.937l-13.51 31.147l35.414 8.827c6.588 1.651 13.045 3.379 19.4 5.006l-11.262 45.213l27.182 6.781l11.153-44.733a1038.209 1038.209 0 0 0 21.687 5.627l-11.115 44.523l27.213 6.781l11.262-45.128c46.404 8.781 81.299 5.239 95.986-36.727c11.836-33.79-.589-53.281-25.004-65.991c17.78-4.098 31.174-15.792 34.747-39.949m-62.177 87.179c-8.41 33.79-65.308 15.523-83.755 10.943l14.944-59.899c18.446 4.603 77.6 13.717 68.811 48.956m8.417-87.667c-7.673 30.736-55.031 15.12-70.393 11.292l13.548-54.327c15.363 3.828 64.836 10.973 56.845 43.035"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bity(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M78.4 67.2C173.8-22 324.5-24 421.5 71c14.3 14.1-6.4 37.1-22.4 21.5c-84.8-82.4-215.8-80.3-298.9-3.2c-16.3 15.1-36.5-8.3-21.8-22.1m98.9 418.6c19.3 5.7 29.3-23.6 7.9-30C73 421.9 9.4 306.1 37.7 194.8c5-19.6-24.9-28.1-30.2-7.1c-32.1 127.4 41.1 259.8 169.8 298.1m148.1-2c121.9-40.2 192.9-166.9 164.4-291c-4.5-19.7-34.9-13.8-30 7.9c24.2 107.7-37.1 217.9-143.2 253.4c-21.2 7-10.4 36 8.8 29.7m-62.9-79l.2-71.8c0-8.2-6.6-14.8-14.8-14.8c-8.2 0-14.8 6.7-14.8 14.8l-.2 71.8c0 8.2 6.6 14.8 14.8 14.8s14.8-6.6 14.8-14.8m71-269c2.1 90.9 4.7 131.9-85.5 132.5c-92.5-.7-86.9-44.3-85.5-132.5c0-21.8-32.5-19.6-32.5 0v71.6c0 69.3 60.7 90.9 118 90.1c57.3.8 118-20.8 118-90.1v-71.6c0-19.6-32.5-21.8-32.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BlackTie(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 32v448h448V32zm316.5 325.2L224 445.9l-92.5-88.7l64.5-184l-64.5-86.6h184.9L252 173.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blackberry(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M166 116.9c0 23.4-16.4 49.1-72.5 49.1H23.4l21-88.8h67.8c42.1 0 53.8 23.3 53.8 39.7m126.2-39.7h-67.8L205.7 166h70.1c53.8 0 70.1-25.7 70.1-49.1c.1-16.4-11.6-39.7-53.7-39.7M88.8 208.1H21L0 296.9h70.1c56.1 0 72.5-23.4 72.5-49.1c0-16.3-11.7-39.7-53.8-39.7m180.1 0h-67.8l-18.7 88.8h70.1c53.8 0 70.1-23.4 70.1-49.1c0-16.3-11.7-39.7-53.7-39.7m189.3-53.8h-67.8l-18.7 88.8h70.1c53.8 0 70.1-23.4 70.1-49.1c.1-16.3-11.6-39.7-53.7-39.7m-28 137.9h-67.8L343.7 381h70.1c56.1 0 70.1-23.4 70.1-49.1c0-16.3-11.6-39.7-53.7-39.7M240.8 346H173l-18.7 88.8h70.1c56.1 0 70.1-25.7 70.1-49.1c.1-16.3-11.6-39.7-53.7-39.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blogger(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M162.4 196c4.8-4.9 6.2-5.1 36.4-5.1c27.2 0 28.1.1 32.1 2.1c5.8 2.9 8.3 7 8.3 13.6c0 5.9-2.4 10-7.6 13.4c-2.8 1.8-4.5 1.9-31.1 2.1c-16.4.1-29.5-.2-31.5-.8c-10.3-2.9-14.1-17.7-6.6-25.3m61.4 94.5c-53.9 0-55.8.2-60.2 4.1c-3.5 3.1-5.7 9.4-5.1 13.9c.7 4.7 4.8 10.1 9.2 12c2.2 1 14.1 1.7 56.3 1.2l47.9-.6l9.2-1.5c9-5.1 10.5-17.4 3.1-24.4c-5.3-4.7-5-4.7-60.4-4.7m223.4 130.1c-3.5 28.4-23 50.4-51.1 57.5c-7.2 1.8-9.7 1.9-172.9 1.8c-157.8 0-165.9-.1-172-1.8c-8.4-2.2-15.6-5.5-22.3-10c-5.6-3.8-13.9-11.8-17-16.4c-3.8-5.6-8.2-15.3-10-22C.1 423 0 420.3 0 256.3C0 93.2 0 89.7 1.8 82.6C8.1 57.9 27.7 39 53 33.4c7.3-1.6 332.1-1.9 340-.3c21.2 4.3 37.9 17.1 47.6 36.4c7.7 15.3 7-1.5 7.3 180.6c.2 115.8 0 164.5-.7 170.5m-85.4-185.2c-1.1-5-4.2-9.6-7.7-11.5c-1.1-.6-8-1.3-15.5-1.7c-12.4-.6-13.8-.8-17.8-3.1c-6.2-3.6-7.9-7.6-8-18.3c0-20.4-8.5-39.4-25.3-56.5c-12-12.2-25.3-20.5-40.6-25.1c-3.6-1.1-11.8-1.5-39.2-1.8c-42.9-.5-52.5.4-67.1 6.2c-27 10.7-46.3 33.4-53.4 62.4c-1.3 5.4-1.6 14.2-1.9 64.3c-.4 62.8 0 72.1 4 84.5c9.7 30.7 37.1 53.4 64.6 58.4c9.2 1.7 122.2 2.1 133.7.5c20.1-2.7 35.9-10.8 50.7-25.9c10.7-10.9 17.4-22.8 21.8-38.5c3.2-10.9 2.9-88.4 1.7-93.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BloggerB(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M446.6 222.7c-1.8-8-6.8-15.4-12.5-18.5c-1.8-1-13-2.2-25-2.7c-20.1-.9-22.3-1.3-28.7-5c-10.1-5.9-12.8-12.3-12.9-29.5c-.1-33-13.8-63.7-40.9-91.3c-19.3-19.7-40.9-33-65.5-40.5c-5.9-1.8-19.1-2.4-63.3-2.9c-69.4-.8-84.8.6-108.4 10C45.9 59.5 14.7 96.1 3.3 142.9C1.2 151.7.7 165.8.2 246.8c-.6 101.5.1 116.4 6.4 136.5c15.6 49.6 59.9 86.3 104.4 94.3c14.8 2.7 197.3 3.3 216 .8c32.5-4.4 58-17.5 81.9-41.9c17.3-17.7 28.1-36.8 35.2-62.1c4.9-17.6 4.5-142.8 2.5-151.7m-322.1-63.6c7.8-7.9 10-8.2 58.8-8.2c43.9 0 45.4.1 51.8 3.4c9.3 4.7 13.4 11.3 13.4 21.9c0 9.5-3.8 16.2-12.3 21.6c-4.6 2.9-7.3 3.1-50.3 3.3c-26.5.2-47.7-.4-50.8-1.2c-16.6-4.7-22.8-28.5-10.6-40.8m191.8 199.8l-14.9 2.4l-77.5.9c-68.1.8-87.3-.4-90.9-2c-7.1-3.1-13.8-11.7-14.9-19.4c-1.1-7.3 2.6-17.3 8.2-22.4c7.1-6.4 10.2-6.6 97.3-6.7c89.6-.1 89.1-.1 97.6 7.8c12.1 11.3 9.5 31.2-4.9 39.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bluetooth(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M292.6 171.1L249.7 214l-.3-86zm-43.2 219.8l43.1-43.1l-42.9-42.9zM416 259.4C416 465 344.1 512 230.9 512S32 465 32 259.4S115.4 0 228.6 0S416 53.9 416 259.4m-158.5 0l79.4-88.6L211.8 36.5v176.9L138 139.6l-27 26.9l92.7 93l-92.7 93l26.9 26.9l73.8-73.8l2.3 170l127.4-127.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BluetoothB(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m196.48 260.023l92.626-103.333L143.125 0v206.33l-86.111-86.111l-31.406 31.405l108.061 108.399L25.608 368.422l31.406 31.405l86.111-86.111L145.84 512l148.552-148.644zm40.86-102.996l-49.977 49.978l-.338-100.295zM187.363 313.04l49.977 49.978l-50.315 50.316z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bootstrap(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M333.5 201.4c0-22.1-15.6-34.3-43-34.3h-50.4v71.2h42.5c32.8-.1 50.9-13.3 50.9-36.9M517 188.6c-9.5-30.9-10.9-68.8-9.8-98.1C508.3 60 484.5 32 452.5 32H123.7C91.6 32 67.9 60.1 69 90.5c1 29.3-.3 67.2-9.8 98.1c-9.6 31-25.7 50.6-52.2 53.1v28.5c26.4 2.5 42.6 22.1 52.2 53.1c9.5 30.9 10.9 68.8 9.8 98.1c-1.1 30.5 22.7 58.5 54.7 58.5h328.7c32.1 0 55.8-28.1 54.7-58.5c-1-29.3.3-67.2 9.8-98.1c9.6-31 25.7-50.6 52.1-53.1v-28.5c-26.3-2.5-42.5-22.1-52-53.1M300.2 375.1h-97.9V136.8h97.4c43.3 0 71.7 23.4 71.7 59.4c0 25.3-19.1 47.9-43.5 51.8v1.3c33.2 3.6 55.5 26.6 55.5 58.3c0 42.1-31.3 67.5-83.2 67.5m-10-108.7h-50.1v78.4h52.3c34.2 0 52.3-13.7 52.3-39.5c0-25.7-18.6-38.9-54.5-38.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Btc(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M310.204 242.638c27.73-14.18 45.377-39.39 41.28-81.3c-5.358-57.351-52.458-76.573-114.85-81.929V0h-48.528v77.203c-12.605 0-25.525.315-38.444.63V0h-48.528v79.409c-17.842.539-38.622.276-97.37 0v51.678c38.314-.678 58.417-3.14 63.023 21.427v217.429c-2.925 19.492-18.524 16.685-53.255 16.071L3.765 443.68c88.481 0 97.37.315 97.37.315V512h48.528v-67.06c13.234.315 26.154.315 38.444.315V512h48.528v-68.005c81.299-4.412 135.647-24.894 142.895-101.467c5.671-61.446-23.32-88.862-69.326-99.89M150.608 134.553c27.415 0 113.126-8.507 113.126 48.528c0 54.515-85.71 48.212-113.126 48.212zm0 251.776V279.821c32.772 0 133.127-9.138 133.127 53.255c-.001 60.186-100.355 53.253-133.127 53.253"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Buffer(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m427.84 380.67l-196.5 97.82a18.6 18.6 0 0 1-14.67 0L20.16 380.67c-4-2-4-5.28 0-7.29L67.22 350a18.65 18.65 0 0 1 14.69 0l134.76 67a18.51 18.51 0 0 0 14.67 0l134.76-67a18.62 18.62 0 0 1 14.68 0l47.06 23.43c4.05 1.96 4.05 5.24 0 7.24m0-136.53l-47.06-23.43a18.62 18.62 0 0 0-14.68 0l-134.76 67.08a18.68 18.68 0 0 1-14.67 0L81.91 220.71a18.65 18.65 0 0 0-14.69 0l-47.06 23.43c-4 2-4 5.29 0 7.31l196.51 97.8a18.6 18.6 0 0 0 14.67 0l196.5-97.8c4.05-2.02 4.05-5.3 0-7.31M20.16 130.42l196.5 90.29a20.08 20.08 0 0 0 14.67 0l196.51-90.29c4-1.86 4-4.89 0-6.74L231.33 33.4a19.88 19.88 0 0 0-14.67 0l-196.5 90.28c-4.05 1.85-4.05 4.88 0 6.74" class="a"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Buromobelexperte(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 32v128h128V32zm120 120H8V40h112zm40-120v128h128V32zm120 120H168V40h112zm40-120v128h128V32zm120 120H328V40h112zM0 192v128h128V192zm120 120H8V200h112zm40-120v128h128V192zm120 120H168V200h112zm40-120v128h128V192zm120 120H328V200h112zM0 352v128h128V352zm120 120H8V360h112zm40-120v128h128V352zm120 120H168V360h112zm40-120v128h128V352z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuyNlarge(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M288 32C133.27 32 7.79 132.32 7.79 256S133.27 480 288 480s280.21-100.32 280.21-224S442.73 32 288 32m-85.39 357.19L64.1 390.55l77.25-290.74h133.44c63.15 0 84.93 28.65 78 72.84a60.24 60.24 0 0 1-1.5 6.85a77.39 77.39 0 0 0-17.21-1.93c-42.35 0-76.69 33.88-76.69 75.65c0 37.14 27.14 68 62.93 74.45c-18.24 37.16-56.16 60.92-117.71 61.52M358 207.11h32l-22.16 90.31h-35.41l-11.19-35.63l-7.83 35.63h-37.83l26.63-90.31h31.34l15 36.75zm145.86 182.08H306.79L322.63 328a78.8 78.8 0 0 0 11.47.83c42.34 0 76.69-33.87 76.69-75.65c0-32.65-21-60.46-50.38-71.06l21.33-82.35h92.5l-53.05 205.36h103.87zM211.7 269.39H187l-13.8 56.47h24.7c16.14 0 32.11-3.18 37.94-26.65c5.56-22.31-7.99-29.82-24.14-29.82M233 170h-21.34L200 217.71h21.37c18 0 35.38-14.64 39.21-30.14C265.23 168.71 251.07 170 233 170"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Buysellads(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m224 150.7l42.9 160.7h-85.8zM448 80v352c0 26.5-21.5 48-48 48H48c-26.5 0-48-21.5-48-48V80c0-26.5 21.5-48 48-48h352c26.5 0 48 21.5 48 48m-65.3 325.3l-94.5-298.7H159.8L65.3 405.3H156l111.7-91.6l24.2 91.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CanadianMapleLeaf(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M383.8 351.7c2.5-2.5 105.2-92.4 105.2-92.4l-17.5-7.5c-10-4.9-7.4-11.5-5-17.4c2.4-7.6 20.1-67.3 20.1-67.3s-47.7 10-57.7 12.5c-7.5 2.4-10-2.5-12.5-7.5s-15-32.4-15-32.4s-52.6 59.9-55.1 62.3c-10 7.5-20.1 0-17.6-10c0-10 27.6-129.6 27.6-129.6s-30.1 17.4-40.1 22.4c-7.5 5-12.6 5-17.6-5C293.5 72.3 255.9 0 255.9 0s-37.5 72.3-42.5 79.8c-5 10-10 10-17.6 5c-10-5-40.1-22.4-40.1-22.4S183.3 182 183.3 192c2.5 10-7.5 17.5-17.6 10c-2.5-2.5-55.1-62.3-55.1-62.3S98.1 167 95.6 172s-5 9.9-12.5 7.5C73 177 25.4 167 25.4 167s17.6 59.7 20.1 67.3c2.4 6 5 12.5-5 17.4L23 259.3s102.6 89.9 105.2 92.4c5.1 5 10 7.5 5.1 22.5c-5.1 15-10.1 35.1-10.1 35.1s95.2-20.1 105.3-22.6c8.7-.9 18.3 2.5 18.3 12.5S241 512 241 512h30s-5.8-102.7-5.8-112.8s9.5-13.4 18.4-12.5c10 2.5 105.2 22.6 105.2 22.6s-5-20.1-10-35.1s0-17.5 5-22.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CcAmazonPay(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M124.7 201.8c.1-11.8 0-23.5 0-35.3v-35.3c0-1.3.4-2 1.4-2.7c11.5-8 24.1-12.1 38.2-11.1c12.5.9 22.7 7 28.1 21.7c3.3 8.9 4.1 18.2 4.1 27.7c0 8.7-.7 17.3-3.4 25.6c-5.7 17.8-18.7 24.7-35.7 23.9c-11.7-.5-21.9-5-31.4-11.7c-.9-.8-1.4-1.6-1.3-2.8m154.9 14.6c4.6 1.8 9.3 2 14.1 1.5c11.6-1.2 21.9-5.7 31.3-12.5c.9-.6 1.3-1.3 1.3-2.5c-.1-3.9 0-7.9 0-11.8c0-4-.1-8 0-12c0-1.4-.4-2-1.8-2.2c-7-.9-13.9-2.2-20.9-2.9c-7-.6-14-.3-20.8 1.9c-6.7 2.2-11.7 6.2-13.7 13.1c-1.6 5.4-1.6 10.8.1 16.2c1.6 5.5 5.2 9.2 10.4 11.2M576 80v352c0 26.5-21.5 48-48 48H48c-26.5 0-48-21.5-48-48V80c0-26.5 21.5-48 48-48h480c26.5 0 48 21.5 48 48m-207.5 23.9c.4 1.7.9 3.4 1.6 5.1c16.5 40.6 32.9 81.3 49.5 121.9c1.4 3.5 1.7 6.4.2 9.9c-2.8 6.2-4.9 12.6-7.8 18.7c-2.6 5.5-6.7 9.5-12.7 11.2c-4.2 1.1-8.5 1.3-12.9.9c-2.1-.2-4.2-.7-6.3-.8c-2.8-.2-4.2 1.1-4.3 4c-.1 2.8-.1 5.6 0 8.3c.1 4.6 1.6 6.7 6.2 7.5c4.7.8 9.4 1.6 14.2 1.7c14.3.3 25.7-5.4 33.1-17.9c2.9-4.9 5.6-10.1 7.7-15.4c19.8-50.1 39.5-100.3 59.2-150.5c.6-1.5 1.1-3 1.3-4.6c.4-2.4-.7-3.6-3.1-3.7c-5.6-.1-11.1 0-16.7 0c-3.1 0-5.3 1.4-6.4 4.3c-.4 1.1-.9 2.3-1.3 3.4l-29.1 83.7c-2.1 6.1-4.2 12.1-6.5 18.6c-.4-.9-.6-1.4-.8-1.9c-10.8-29.9-21.6-59.9-32.4-89.8c-1.7-4.7-3.5-9.5-5.3-14.2c-.9-2.5-2.7-4-5.4-4c-6.4-.1-12.8-.2-19.2-.1c-2.2 0-3.3 1.6-2.8 3.7M242.4 206c1.7 11.7 7.6 20.8 18 26.6c9.9 5.5 20.7 6.2 31.7 4.6c12.7-1.9 23.9-7.3 33.8-15.5c.4-.3.8-.6 1.4-1c.5 3.2.9 6.2 1.5 9.2c.5 2.6 2.1 4.3 4.5 4.4c4.6.1 9.1.1 13.7 0c2.3-.1 3.8-1.6 4-3.9c.1-.8.1-1.6.1-2.3v-88.8c0-3.6-.2-7.2-.7-10.8c-1.6-10.8-6.2-19.7-15.9-25.4c-5.6-3.3-11.8-5-18.2-5.9c-3-.4-6-.7-9.1-1.1h-10c-.8.1-1.6.3-2.5.3c-8.2.4-16.3 1.4-24.2 3.5c-5.1 1.3-10 3.2-15 4.9c-3 1-4.5 3.2-4.4 6.5c.1 2.8-.1 5.6 0 8.3c.1 4.1 1.8 5.2 5.7 4.1c6.5-1.7 13.1-3.5 19.7-4.8c10.3-1.9 20.7-2.7 31.1-1.2c5.4.8 10.5 2.4 14.1 7c3.1 4 4.2 8.8 4.4 13.7c.3 6.9.2 13.9.3 20.8c0 .4-.1.7-.2 1.2c-.4 0-.8 0-1.1-.1c-8.8-2.1-17.7-3.6-26.8-4.1c-9.5-.5-18.9.1-27.9 3.2c-10.8 3.8-19.5 10.3-24.6 20.8c-4.1 8.3-4.6 17-3.4 25.8M98.7 106.9v175.3c0 .8 0 1.7.1 2.5c.2 2.5 1.7 4.1 4.1 4.2c5.9.1 11.8.1 17.7 0c2.5 0 4-1.7 4.1-4.1c.1-.8.1-1.7.1-2.5v-60.7c.9.7 1.4 1.2 1.9 1.6c15 12.5 32.2 16.6 51.1 12.9c17.1-3.4 28.9-13.9 36.7-29.2c5.8-11.6 8.3-24.1 8.7-37c.5-14.3-1-28.4-6.8-41.7c-7.1-16.4-18.9-27.3-36.7-30.9c-2.7-.6-5.5-.8-8.2-1.2h-7c-1.2.2-2.4.3-3.6.5c-11.7 1.4-22.3 5.8-31.8 12.7c-2 1.4-3.9 3-5.9 4.5c-.1-.5-.3-.8-.4-1.2c-.4-2.3-.7-4.6-1.1-6.9c-.6-3.9-2.5-5.5-6.4-5.6h-9.7c-5.9-.1-6.9 1-6.9 6.8M493.6 339c-2.7-.7-5.1 0-7.6 1c-43.9 18.4-89.5 30.2-136.8 35.8c-14.5 1.7-29.1 2.8-43.7 3.2c-26.6.7-53.2-.8-79.6-4.3c-17.8-2.4-35.5-5.7-53-9.9c-37-8.9-72.7-21.7-106.7-38.8c-8.8-4.4-17.4-9.3-26.1-14c-3.8-2.1-6.2-1.5-8.2 2.1v1.7c1.2 1.6 2.2 3.4 3.7 4.8c36 32.2 76.6 56.5 122 72.9c21.9 7.9 44.4 13.7 67.3 17.5c14 2.3 28 3.8 42.2 4.5c3 .1 6 .2 9 .4c.7 0 1.4.2 2.1.3h17.7c.7-.1 1.4-.3 2.1-.3c14.9-.4 29.8-1.8 44.6-4c21.4-3.2 42.4-8.1 62.9-14.7c29.6-9.6 57.7-22.4 83.4-40.1c2.8-1.9 5.7-3.8 8-6.2c4.3-4.4 2.3-10.4-3.3-11.9m50.4-27.7c-.8-4.2-4-5.8-7.6-7c-5.7-1.9-11.6-2.8-17.6-3.3c-11-.9-22-.4-32.8 1.6c-12 2.2-23.4 6.1-33.5 13.1c-1.2.8-2.4 1.8-3.1 3c-.6.9-.7 2.3-.5 3.4c.3 1.3 1.7 1.6 3 1.5c.6 0 1.2 0 1.8-.1l19.5-2.1c9.6-.9 19.2-1.5 28.8-.8c4.1.3 8.1 1.2 12 2.2c4.3 1.1 6.2 4.4 6.4 8.7c.3 6.7-1.2 13.1-2.9 19.5c-3.5 12.9-8.3 25.4-13.3 37.8c-.3.8-.7 1.7-.8 2.5c-.4 2.5 1 4 3.4 3.5c1.4-.3 3-1.1 4-2.1c3.7-3.6 7.5-7.2 10.6-11.2c10.7-13.8 17-29.6 20.7-46.6c.7-3 1.2-6.1 1.7-9.1c.2-4.7.2-9.6.2-14.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CcAmex(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M325.1 167.8c0-16.4-14.1-18.4-27.4-18.4l-39.1-.3v69.3H275v-25.1h18c18.4 0 14.5 10.3 14.8 25.1h16.6v-13.5c0-9.2-1.5-15.1-11-18.4c7.4-3 11.8-10.7 11.7-18.7m-29.4 11.3H275v-15.3h21c5.1 0 10.7 1 10.7 7.4c0 6.6-5.3 7.9-11 7.9M279 268.6h-52.7l-21 22.8l-20.5-22.8h-66.5l-.1 69.3h65.4l21.3-23l20.4 23h32.2l.1-23.3c18.9 0 49.3 4.6 49.3-23.3c0-17.3-12.3-22.7-27.9-22.7m-103.8 54.7h-40.6v-13.8h36.3v-14.1h-36.3v-12.5h41.7l17.9 20.2zm65.8 8.2l-25.3-28.1L241 276zm37.8-31h-21.2v-17.6h21.5c5.6 0 10.2 2.3 10.2 8.4c0 6.4-4.6 9.2-10.5 9.2m-31.6-136.7v-14.6h-55.5v69.3h55.5v-14.3h-38.9v-13.8h37.8v-14.1h-37.8v-12.5zM576 255.4h-.2zm-194.6 31.9c0-16.4-14.1-18.7-27.1-18.7h-39.4l-.1 69.3h16.6l.1-25.3h17.6c11 0 14.8 2 14.8 13.8l-.1 11.5h16.6l.1-13.8c0-8.9-1.8-15.1-11-18.4c7.7-3.1 11.8-10.8 11.9-18.4m-29.2 11.2h-20.7v-15.6h21c5.1 0 10.7 1 10.7 7.4c0 6.9-5.4 8.2-11 8.2m-172.8-80v-69.3h-27.6l-19.7 47l-21.7-47H83.3v65.7l-28.1-65.7H30.7L1 218.5h17.9l6.4-15.3h34.5l6.4 15.3H100v-54.2l24 54.2h14.6l24-54.2v54.2zM31.2 188.8l11.2-27.6l11.5 27.6zm477.4 158.9v-4.5c-10.8 5.6-3.9 4.5-156.7 4.5c0-25.2.1-23.9 0-25.2c-1.7-.1-3.2-.1-9.4-.1c0 17.9-.1 6.8-.1 25.3h-39.6c0-12.1.1-15.3.1-29.2c-10 6-22.8 6.4-34.3 6.2c0 14.7-.1 8.3-.1 23h-48.9c-5.1-5.7-2.7-3.1-15.4-17.4c-3.2 3.5-12.8 13.9-16.1 17.4h-82v-92.3h83.1c5 5.6 2.8 3.1 15.5 17.2c3.2-3.5 12.2-13.4 15.7-17.2h58c9.8 0 18 1.9 24.3 5.6v-5.6c54.3 0 64.3-1.4 75.7 5.1v-5.1h78.2v5.2c11.4-6.9 19.6-5.2 64.9-5.2v5c10.3-5.9 16.6-5.2 54.3-5V80c0-26.5-21.5-48-48-48h-480c-26.5 0-48 21.5-48 48v109.8c9.4-21.9 19.7-46 23.1-53.9h39.7c4.3 10.1 1.6 3.7 9 21.1v-21.1h46c2.9 6.2 11.1 24 13.9 30c5.8-13.6 10.1-23.9 12.6-30h103c0-.1 11.5 0 11.6 0c43.7.2 53.6-.8 64.4 5.3v-5.3H363v9.3c7.6-6.1 17.9-9.3 30.7-9.3h27.6c0 .5 1.9.3 2.3.3H456c4.2 9.8 2.6 6 8.8 20.6v-20.6h43.3c4.9 8-1-1.8 11.2 18.4v-18.4h39.9v92h-41.6c-5.4-9-1.4-2.2-13.2-21.9v21.9h-52.8c-6.4-14.8-.1-.3-6.6-15.3h-19c-4.2 10-2.2 5.2-6.4 15.3h-26.8c-12.3 0-22.3-3-29.7-8.9v8.9h-66.5c-.3-13.9-.1-24.8-.1-24.8c-1.8-.3-3.4-.2-9.8-.2v25.1H151.2v-11.4c-2.5 5.6-2.7 5.9-5.1 11.4h-29.5c-4-8.9-2.9-6.4-5.1-11.4v11.4H58.6c-4.2-10.1-2.2-5.3-6.4-15.3H33c-4.2 10-2.2 5.2-6.4 15.3H0V432c0 26.5 21.5 48 48 48h480.1c26.5 0 48-21.5 48-48v-90.4c-12.7 8.3-32.7 6.1-67.5 6.1m36.3-64.5H575v-14.6h-32.9c-12.8 0-23.8 6.6-23.8 20.7c0 33 42.7 12.8 42.7 27.4c0 5.1-4.3 6.4-8.4 6.4h-32l-.1 14.8h32c8.4 0 17.6-1.8 22.5-8.9v-25.8c-10.5-13.8-39.3-1.3-39.3-13.5c0-5.8 4.6-6.5 9.2-6.5m-57 39.8h-32.2l-.1 14.8h32.2c14.8 0 26.2-5.6 26.2-22c0-33.2-42.9-11.2-42.9-26.3c0-5.6 4.9-6.4 9.2-6.4h30.4v-14.6h-33.2c-12.8 0-23.5 6.6-23.5 20.7c0 33 42.7 12.5 42.7 27.4c-.1 5.4-4.7 6.4-8.8 6.4m-42.2-40.1v-14.3h-55.2l-.1 69.3h55.2l.1-14.3l-38.6-.3v-13.8H445v-14.1h-37.8v-12.5zm-56.3-108.1c-.3.2-1.4 2.2-1.4 7.6c0 6 .9 7.7 1.1 7.9c.2.1 1.1.5 3.4.5l7.3-16.9c-1.1 0-2.1-.1-3.1-.1c-5.6 0-7 .7-7.3 1m20.4-10.5h-.1zm-16.2-15.2c-23.5 0-34 12-34 35.3c0 22.2 10.2 34 33 34h19.2l6.4-15.3h34.3l6.6 15.3h33.7v-51.9l31.2 51.9h23.6v-69h-16.9v48.1l-29.1-48.1h-25.3v65.4l-27.9-65.4h-24.8l-23.5 54.5h-7.4c-13.3 0-16.1-8.1-16.1-19.9c0-23.8 15.7-20 33.1-19.7v-15.2zm42.1 12.1l11.2 27.6h-22.8zm-101.1-12v69.3h16.9v-69.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CcApplePay(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M302.2 218.4c0 17.2-10.5 27.1-29 27.1h-24.3v-54.2h24.4c18.4 0 28.9 9.8 28.9 27.1m47.5 62.6c0 8.3 7.2 13.7 18.5 13.7c14.4 0 25.2-9.1 25.2-21.9v-7.7l-23.5 1.5c-13.3.9-20.2 5.8-20.2 14.4M576 79v352c0 26.5-21.5 48-48 48H48c-26.5 0-48-21.5-48-48V79c0-26.5 21.5-48 48-48h480c26.5 0 48 21.5 48 48M127.8 197.2c8.4.7 16.8-4.2 22.1-10.4c5.2-6.4 8.6-15 7.7-23.7c-7.4.3-16.6 4.9-21.9 11.3c-4.8 5.5-8.9 14.4-7.9 22.8m60.6 74.5c-.2-.2-19.6-7.6-19.8-30c-.2-18.7 15.3-27.7 16-28.2c-8.8-13-22.4-14.4-27.1-14.7c-12.2-.7-22.6 6.9-28.4 6.9c-5.9 0-14.7-6.6-24.3-6.4c-12.5.2-24.2 7.3-30.5 18.6c-13.1 22.6-3.4 56 9.3 74.4c6.2 9.1 13.7 19.1 23.5 18.7c9.3-.4 13-6 24.2-6c11.3 0 14.5 6 24.3 5.9c10.2-.2 16.5-9.1 22.8-18.2c6.9-10.4 9.8-20.4 10-21m135.4-53.4c0-26.6-18.5-44.8-44.9-44.8h-51.2v136.4h21.2v-46.6h29.3c26.8 0 45.6-18.4 45.6-45m90 23.7c0-19.7-15.8-32.4-40-32.4c-22.5 0-39.1 12.9-39.7 30.5h19.1c1.6-8.4 9.4-13.9 20-13.9c13 0 20.2 6 20.2 17.2v7.5l-26.4 1.6c-24.6 1.5-37.9 11.6-37.9 29.1c0 17.7 13.7 29.4 33.4 29.4c13.3 0 25.6-6.7 31.2-17.4h.4V310h19.6v-68zM516 210.9h-21.5l-24.9 80.6h-.4l-24.9-80.6H422l35.9 99.3l-1.9 6c-3.2 10.2-8.5 14.2-17.9 14.2c-1.7 0-4.9-.2-6.2-.3v16.4c1.2.4 6.5.5 8.1.5c20.7 0 30.4-7.9 38.9-31.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CcDinersClub(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M239.7 79.9c-96.9 0-175.8 78.6-175.8 175.8c0 96.9 78.9 175.8 175.8 175.8c97.2 0 175.8-78.9 175.8-175.8c0-97.2-78.6-175.8-175.8-175.8m-39.9 279.6c-41.7-15.9-71.4-56.4-71.4-103.8s29.7-87.9 71.4-104.1zm79.8.3V151.6c41.7 16.2 71.4 56.7 71.4 104.1s-29.7 87.9-71.4 104.1M528 32H48C21.5 32 0 53.5 0 80v352c0 26.5 21.5 48 48 48h480c26.5 0 48-21.5 48-48V80c0-26.5-21.5-48-48-48M329.7 448h-90.3c-106.2 0-193.8-85.5-193.8-190.2C45.6 143.2 133.2 64 239.4 64h90.3c105 0 200.7 79.2 200.7 193.8c0 104.7-95.7 190.2-200.7 190.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CcDiscover(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M520.4 196.1c0-7.9-5.5-12.1-15.6-12.1h-4.9v24.9h4.7c10.3 0 15.8-4.4 15.8-12.8M528 32H48C21.5 32 0 53.5 0 80v352c0 26.5 21.5 48 48 48h480c26.5 0 48-21.5 48-48V80c0-26.5-21.5-48-48-48m-44.1 138.9c22.6 0 52.9-4.1 52.9 24.4c0 12.6-6.6 20.7-18.7 23.2l25.8 34.4h-19.6l-22.2-32.8h-2.2v32.8h-16zm-55.9.1h45.3v14H444v18.2h28.3V217H444v22.2h29.3V253H428zm-68.7 0l21.9 55.2l22.2-55.2h17.5l-35.5 84.2h-8.6l-35-84.2zm-55.9-3c24.7 0 44.6 20 44.6 44.6c0 24.7-20 44.6-44.6 44.6c-24.7 0-44.6-20-44.6-44.6c0-24.7 20-44.6 44.6-44.6m-49.3 6.1v19c-20.1-20.1-46.8-4.7-46.8 19c0 25 27.5 38.5 46.8 19.2v19c-29.7 14.3-63.3-5.7-63.3-38.2c0-31.2 33.1-53 63.3-38m-97.2 66.3c11.4 0 22.4-15.3-3.3-24.4c-15-5.5-20.2-11.4-20.2-22.7c0-23.2 30.6-31.4 49.7-14.3l-8.4 10.8c-10.4-11.6-24.9-6.2-24.9 2.5c0 4.4 2.7 6.9 12.3 10.3c18.2 6.6 23.6 12.5 23.6 25.6c0 29.5-38.8 37.4-56.6 11.3l10.3-9.9c3.7 7.1 9.9 10.8 17.5 10.8M55.4 253H32v-82h23.4c26.1 0 44.1 17 44.1 41.1c0 18.5-13.2 40.9-44.1 40.9m67.5 0h-16v-82h16zM544 433c0 8.2-6.8 15-15 15H128c189.6-35.6 382.7-139.2 416-160zM74.1 191.6c-5.2-4.9-11.6-6.6-21.9-6.6H48v54.2h4.2c10.3 0 17-2 21.9-6.4c5.7-5.2 8.9-12.8 8.9-20.7s-3.2-15.5-8.9-20.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CcJcb(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M431.5 244.3V212c41.2 0 38.5.2 38.5.2c7.3 1.3 13.3 7.3 13.3 16c0 8.8-6 14.5-13.3 15.8c-1.2.4-3.3.3-38.5.3m42.8 20.2c-2.8-.7-3.3-.5-42.8-.5v35c39.6 0 40 .2 42.8-.5c7.5-1.5 13.5-8 13.5-17c0-8.7-6-15.5-13.5-17M576 80v352c0 26.5-21.5 48-48 48H48c-26.5 0-48-21.5-48-48V80c0-26.5 21.5-48 48-48h480c26.5 0 48 21.5 48 48M182 192.3h-57c0 67.1 10.7 109.7-35.8 109.7c-19.5 0-38.8-5.7-57.2-14.8v28c30 8.3 68 8.3 68 8.3c97.9 0 82-47.7 82-131.2m178.5 4.5c-63.4-16-165-14.9-165 59.3c0 77.1 108.2 73.6 165 59.2V287C312.9 311.7 253 309 253 256s59.8-55.6 107.5-31.2zM544 286.5c0-18.5-16.5-30.5-38-32v-.8c19.5-2.7 30.3-15.5 30.3-30.2c0-19-15.7-30-37-31c0 0 6.3-.3-120.3-.3v127.5h122.7c24.3.1 42.3-12.9 42.3-33.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CcMastercard(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M482.9 410.3c0 6.8-4.6 11.7-11.2 11.7c-6.8 0-11.2-5.2-11.2-11.7c0-6.5 4.4-11.7 11.2-11.7c6.6 0 11.2 5.2 11.2 11.7m-310.8-11.7c-7.1 0-11.2 5.2-11.2 11.7c0 6.5 4.1 11.7 11.2 11.7c6.5 0 10.9-4.9 10.9-11.7c-.1-6.5-4.4-11.7-10.9-11.7m117.5-.3c-5.4 0-8.7 3.5-9.5 8.7h19.1c-.9-5.7-4.4-8.7-9.6-8.7m107.8.3c-6.8 0-10.9 5.2-10.9 11.7c0 6.5 4.1 11.7 10.9 11.7c6.8 0 11.2-4.9 11.2-11.7c0-6.5-4.4-11.7-11.2-11.7m105.9 26.1c0 .3.3.5.3 1.1c0 .3-.3.5-.3 1.1c-.3.3-.3.5-.5.8c-.3.3-.5.5-1.1.5c-.3.3-.5.3-1.1.3c-.3 0-.5 0-1.1-.3c-.3 0-.5-.3-.8-.5c-.3-.3-.5-.5-.5-.8c-.3-.5-.3-.8-.3-1.1c0-.5 0-.8.3-1.1c0-.5.3-.8.5-1.1c.3-.3.5-.3.8-.5c.5-.3.8-.3 1.1-.3c.5 0 .8 0 1.1.3c.5.3.8.3 1.1.5s.2.6.5 1.1m-2.2 1.4c.5 0 .5-.3.8-.3c.3-.3.3-.5.3-.8c0-.3 0-.5-.3-.8c-.3 0-.5-.3-1.1-.3h-1.6v3.5h.8V426h.3l1.1 1.4h.8zM576 81v352c0 26.5-21.5 48-48 48H48c-26.5 0-48-21.5-48-48V81c0-26.5 21.5-48 48-48h480c26.5 0 48 21.5 48 48M64 220.6c0 76.5 62.1 138.5 138.5 138.5c27.2 0 53.9-8.2 76.5-23.1c-72.9-59.3-72.4-171.2 0-230.5c-22.6-15-49.3-23.1-76.5-23.1c-76.4-.1-138.5 62-138.5 138.2m224 108.8c70.5-55 70.2-162.2 0-217.5c-70.2 55.3-70.5 162.6 0 217.5m-142.3 76.3c0-8.7-5.7-14.4-14.7-14.7c-4.6 0-9.5 1.4-12.8 6.5c-2.4-4.1-6.5-6.5-12.2-6.5c-3.8 0-7.6 1.4-10.6 5.4V392h-8.2v36.7h8.2c0-18.9-2.5-30.2 9-30.2c10.2 0 8.2 10.2 8.2 30.2h7.9c0-18.3-2.5-30.2 9-30.2c10.2 0 8.2 10 8.2 30.2h8.2v-23zm44.9-13.7h-7.9v4.4c-2.7-3.3-6.5-5.4-11.7-5.4c-10.3 0-18.2 8.2-18.2 19.3c0 11.2 7.9 19.3 18.2 19.3c5.2 0 9-1.9 11.7-5.4v4.6h7.9zm40.5 25.6c0-15-22.9-8.2-22.9-15.2c0-5.7 11.9-4.8 18.5-1.1l3.3-6.5c-9.4-6.1-30.2-6-30.2 8.2c0 14.3 22.9 8.3 22.9 15c0 6.3-13.5 5.8-20.7.8l-3.5 6.3c11.2 7.6 32.6 6 32.6-7.5m35.4 9.3l-2.2-6.8c-3.8 2.1-12.2 4.4-12.2-4.1v-16.6h13.1V392h-13.1v-11.2h-8.2V392h-7.6v7.3h7.6V416c0 17.6 17.3 14.4 22.6 10.9m13.3-13.4h27.5c0-16.2-7.4-22.6-17.4-22.6c-10.6 0-18.2 7.9-18.2 19.3c0 20.5 22.6 23.9 33.8 14.2l-3.8-6c-7.8 6.4-19.6 5.8-21.9-4.9m59.1-21.5c-4.6-2-11.6-1.8-15.2 4.4V392h-8.2v36.7h8.2V408c0-11.6 9.5-10.1 12.8-8.4zm10.6 18.3c0-11.4 11.6-15.1 20.7-8.4l3.8-6.5c-11.6-9.1-32.7-4.1-32.7 15c0 19.8 22.4 23.8 32.7 15l-3.8-6.5c-9.2 6.5-20.7 2.6-20.7-8.6m66.7-18.3H408v4.4c-8.3-11-29.9-4.8-29.9 13.9c0 19.2 22.4 24.7 29.9 13.9v4.6h8.2zm33.7 0c-2.4-1.2-11-2.9-15.2 4.4V392h-7.9v36.7h7.9V408c0-11 9-10.3 12.8-8.4zm40.3-14.9h-7.9v19.3c-8.2-10.9-29.9-5.1-29.9 13.9c0 19.4 22.5 24.6 29.9 13.9v4.6h7.9zm7.6-75.1v4.6h.8V302h1.9v-.8h-4.6v.8zm6.6 123.8c0-.5 0-1.1-.3-1.6c-.3-.3-.5-.8-.8-1.1c-.3-.3-.8-.5-1.1-.8c-.5 0-1.1-.3-1.6-.3c-.3 0-.8.3-1.4.3c-.5.3-.8.5-1.1.8c-.5.3-.8.8-.8 1.1c-.3.5-.3 1.1-.3 1.6c0 .3 0 .8.3 1.4c0 .3.3.8.8 1.1c.3.3.5.5 1.1.8c.5.3 1.1.3 1.4.3c.5 0 1.1 0 1.6-.3c.3-.3.8-.5 1.1-.8c.3-.3.5-.8.8-1.1c.3-.6.3-1.1.3-1.4m3.2-124.7h-1.4l-1.6 3.5l-1.6-3.5h-1.4v5.4h.8v-4.1l1.6 3.5h1.1l1.4-3.5v4.1h1.1zm4.4-80.5c0-76.2-62.1-138.3-138.5-138.3c-27.2 0-53.9 8.2-76.5 23.1c72.1 59.3 73.2 171.5 0 230.5c22.6 15 49.5 23.1 76.5 23.1c76.4.1 138.5-61.9 138.5-138.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CcPaypal(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M186.3 258.2c0 12.2-9.7 21.5-22 21.5c-9.2 0-16-5.2-16-15c0-12.2 9.5-22 21.7-22c9.3 0 16.3 5.7 16.3 15.5M80.5 209.7h-4.7c-1.5 0-3 1-3.2 2.7l-4.3 26.7l8.2-.3c11 0 19.5-1.5 21.5-14.2c2.3-13.4-6.2-14.9-17.5-14.9m284 0H360c-1.8 0-3 1-3.2 2.7l-4.2 26.7l8-.3c13 0 22-3 22-18c-.1-10.6-9.6-11.1-18.1-11.1M576 80v352c0 26.5-21.5 48-48 48H48c-26.5 0-48-21.5-48-48V80c0-26.5 21.5-48 48-48h480c26.5 0 48 21.5 48 48M128.3 215.4c0-21-16.2-28-34.7-28h-40c-2.5 0-5 2-5.2 4.7L32 294.2c-.3 2 1.2 4 3.2 4h19c2.7 0 5.2-2.9 5.5-5.7l4.5-26.6c1-7.2 13.2-4.7 18-4.7c28.6 0 46.1-17 46.1-45.8m84.2 8.8h-19c-3.8 0-4 5.5-4.2 8.2c-5.8-8.5-14.2-10-23.7-10c-24.5 0-43.2 21.5-43.2 45.2c0 19.5 12.2 32.2 31.7 32.2c9 0 20.2-4.9 26.5-11.9c-.5 1.5-1 4.7-1 6.2c0 2.3 1 4 3.2 4H200c2.7 0 5-2.9 5.5-5.7l10.2-64.3c.3-1.9-1.2-3.9-3.2-3.9m40.5 97.9l63.7-92.6c.5-.5.5-1 .5-1.7c0-1.7-1.5-3.5-3.2-3.5h-19.2c-1.7 0-3.5 1-4.5 2.5l-26.5 39l-11-37.5c-.8-2.2-3-4-5.5-4h-18.7c-1.7 0-3.2 1.8-3.2 3.5c0 1.2 19.5 56.8 21.2 62.1c-2.7 3.8-20.5 28.6-20.5 31.6c0 1.8 1.5 3.2 3.2 3.2h19.2c1.8-.1 3.5-1.1 4.5-2.6m159.3-106.7c0-21-16.2-28-34.7-28h-39.7c-2.7 0-5.2 2-5.5 4.7l-16.2 102c-.2 2 1.3 4 3.2 4h20.5c2 0 3.5-1.5 4-3.2l4.5-29c1-7.2 13.2-4.7 18-4.7c28.4 0 45.9-17 45.9-45.8m84.2 8.8h-19c-3.8 0-4 5.5-4.3 8.2c-5.5-8.5-14-10-23.7-10c-24.5 0-43.2 21.5-43.2 45.2c0 19.5 12.2 32.2 31.7 32.2c9.3 0 20.5-4.9 26.5-11.9c-.3 1.5-1 4.7-1 6.2c0 2.3 1 4 3.2 4H484c2.7 0 5-2.9 5.5-5.7l10.2-64.3c.3-1.9-1.2-3.9-3.2-3.9m47.5-33.3c0-2-1.5-3.5-3.2-3.5h-18.5c-1.5 0-3 1.2-3.2 2.7l-16.2 104l-.3.5c0 1.8 1.5 3.5 3.5 3.5h16.5c2.5 0 5-2.9 5.2-5.7L544 191.2zm-90 51.8c-12.2 0-21.7 9.7-21.7 22c0 9.7 7 15 16.2 15c12 0 21.7-9.2 21.7-21.5c.1-9.8-6.9-15.5-16.2-15.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CcStripe(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M492.4 220.8c-8.9 0-18.7 6.7-18.7 22.7h36.7c0-16-9.3-22.7-18-22.7M375 223.4c-8.2 0-13.3 2.9-17 7l.2 52.8c3.5 3.7 8.5 6.7 16.8 6.7c13.1 0 21.9-14.3 21.9-33.4c0-18.6-9-33.2-21.9-33.1M528 32H48C21.5 32 0 53.5 0 80v352c0 26.5 21.5 48 48 48h480c26.5 0 48-21.5 48-48V80c0-26.5-21.5-48-48-48M122.2 281.1c0 25.6-20.3 40.1-49.9 40.3c-12.2 0-25.6-2.4-38.8-8.1v-33.9c12 6.4 27.1 11.3 38.9 11.3c7.9 0 13.6-2.1 13.6-8.7c0-17-54-10.6-54-49.9c0-25.2 19.2-40.2 48-40.2c11.8 0 23.5 1.8 35.3 6.5v33.4c-10.8-5.8-24.5-9.1-35.3-9.1c-7.5 0-12.1 2.2-12.1 7.7c0 16 54.3 8.4 54.3 50.7m68.8-56.6h-27V275c0 20.9 22.5 14.4 27 12.6v28.9c-4.7 2.6-13.3 4.7-24.9 4.7c-21.1 0-36.9-15.5-36.9-36.5l.2-113.9l34.7-7.4v30.8H191zm74 2.4c-4.5-1.5-18.7-3.6-27.1 7.4v84.4h-35.5V194.2h30.7l2.2 10.5c8.3-15.3 24.9-12.2 29.6-10.5h.1zm44.1 91.8h-35.7V194.2h35.7zm0-142.9l-35.7 7.6v-28.9l35.7-7.6zm74.1 145.5c-12.4 0-20-5.3-25.1-9l-.1 40.2l-35.5 7.5V194.2h31.3l1.8 8.8c4.9-4.5 13.9-11.1 27.8-11.1c24.9 0 48.4 22.5 48.4 63.8c0 45.1-23.2 65.5-48.6 65.6m160.4-51.5h-69.5c1.6 16.6 13.8 21.5 27.6 21.5c14.1 0 25.2-3 34.9-7.9V312c-9.7 5.3-22.4 9.2-39.4 9.2c-34.6 0-58.8-21.7-58.8-64.5c0-36.2 20.5-64.9 54.3-64.9c33.7 0 51.3 28.7 51.3 65.1c0 3.5-.3 10.9-.4 12.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CcVisa(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M470.1 231.3s7.6 37.2 9.3 45H446c3.3-8.9 16-43.5 16-43.5c-.2.3 3.3-9.1 5.3-14.9zM576 80v352c0 26.5-21.5 48-48 48H48c-26.5 0-48-21.5-48-48V80c0-26.5 21.5-48 48-48h480c26.5 0 48 21.5 48 48M152.5 331.2L215.7 176h-42.5l-39.3 106l-4.3-21.5l-14-71.4c-2.3-9.9-9.4-12.7-18.2-13.1H32.7l-.7 3.1c15.8 4 29.9 9.8 42.2 17.1l35.8 135zm94.4.2L272.1 176h-40.2l-25.1 155.4zm139.9-50.8c.2-17.7-10.6-31.2-33.7-42.3c-14.1-7.1-22.7-11.9-22.7-19.2c.2-6.6 7.3-13.4 23.1-13.4c13.1-.3 22.7 2.8 29.9 5.9l3.6 1.7l5.5-33.6c-7.9-3.1-20.5-6.6-36-6.6c-39.7 0-67.6 21.2-67.8 51.4c-.3 22.3 20 34.7 35.2 42.2c15.5 7.6 20.8 12.6 20.8 19.3c-.2 10.4-12.6 15.2-24.1 15.2c-16 0-24.6-2.5-37.7-8.3l-5.3-2.5l-5.6 34.9c9.4 4.3 26.8 8.1 44.8 8.3c42.2.1 69.7-20.8 70-53M528 331.4L495.6 176h-31.1c-9.6 0-16.9 2.8-21 12.9l-59.7 142.5H426s6.9-19.2 8.4-23.3H486c1.2 5.5 4.8 23.3 4.8 23.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Centercode(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M329.2 268.6c-3.8 35.2-35.4 60.6-70.6 56.8c-35.2-3.8-60.6-35.4-56.8-70.6c3.8-35.2 35.4-60.6 70.6-56.8c35.1 3.8 60.6 35.4 56.8 70.6m-85.8 235.1C96.7 496-8.2 365.5 10.1 224.3c11.2-86.6 65.8-156.9 139.1-192c161-77.1 349.7 37.4 354.7 216.6c4.1 147-118.4 262.2-260.5 254.8m179.9-180c27.9-118-160.5-205.9-237.2-234.2c-57.5 56.3-69.1 188.6-33.8 344.4c68.8 15.8 169.1-26.4 271-110.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Centos(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m289.6 97.5l31.6 31.7l-76.3 76.5V97.5zm-162.4 31.7l76.3 76.5V97.5h-44.7zm41.5-41.6h44.7v127.9l10.8 10.8l10.8-10.8V87.6h44.7L224.2 32zm26.2 168.1l-10.8-10.8H55.5v-44.8L0 255.7l55.5 55.6v-44.8h128.6zm79.3-20.7h107.9v-44.8l-31.6-31.7zm173.3 20.7L392 200.1v44.8H264.3l-10.8 10.8l10.8 10.8H392v44.8zM65.4 176.2l32.5-31.7l90.3 90.5h15.3v-15.3l-90.3-90.5l31.6-31.7H65.4zm316.7-78.7h-78.5l31.6 31.7l-90.3 90.5V235h15.3l90.3-90.5l31.6 31.7zM203.5 413.9V305.8l-76.3 76.5l31.6 31.7h44.7zM65.4 235h108.8l-76.3-76.5l-32.5 31.7zm316.7 100.2l-31.6 31.7l-90.3-90.5h-15.3v15.3l90.3 90.5l-31.6 31.7h78.5zm0-58.8H274.2l76.3 76.5l31.6-31.7zm-60.9 105.8l-76.3-76.5v108.1h44.7zM97.9 352.9l76.3-76.5H65.4v44.8zm181.8 70.9H235V295.9l-10.8-10.8l-10.8 10.8v127.9h-44.7l55.5 55.6zm-166.5-41.6l90.3-90.5v-15.3h-15.3l-90.3 90.5l-32.5-31.7v78.7h79.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chrome(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M131.5 217.5L55.1 100.1c47.6-59.2 119-91.8 192-92.1c42.3-.3 85.5 10.5 124.8 33.2c43.4 25.2 76.4 61.4 97.4 103L264 133.4c-58.1-3.4-113.4 29.3-132.5 84.1m32.9 38.5c0 46.2 37.4 83.6 83.6 83.6s83.6-37.4 83.6-83.6s-37.4-83.6-83.6-83.6s-83.6 37.3-83.6 83.6m314.9-89.2L339.6 174c37.9 44.3 38.5 108.2 6.6 157.2L234.1 503.6c46.5 2.5 94.4-7.7 137.8-32.9c107.4-62 150.9-192 107.4-303.9M133.7 303.6L40.4 120.1C14.9 159.1 0 205.9 0 256c0 124 90.8 226.7 209.5 244.9l63.7-124.8c-57.6 10.8-113.2-20.8-139.5-72.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chromecast(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M447.8 64H64c-23.6 0-42.7 19.1-42.7 42.7v63.9H64v-63.9h383.8v298.6H298.6V448H448c23.6 0 42.7-19.1 42.7-42.7V106.7c0-23.6-19.3-42.7-42.9-42.7M21.3 383.6v63.9h63.9c0-35.3-28.6-63.9-63.9-63.9m0-85V341c58.9 0 106.6 48.1 106.6 107h42.7c.1-82.4-66.9-149.3-149.3-149.4M213.4 448h42.7c-.5-129.5-105.3-234.3-234.8-234.6v42.4c106-.2 192 86.2 192.1 192.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloudflare(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m407.906 319.913l-230.8-2.928a4.58 4.58 0 0 1-3.632-1.926a4.648 4.648 0 0 1-.494-4.147a6.143 6.143 0 0 1 5.361-4.076l232.94-2.936c27.631-1.26 57.546-23.574 68.022-50.784l13.286-34.542a7.944 7.944 0 0 0 .524-2.936a7.735 7.735 0 0 0-.164-1.631A151.91 151.91 0 0 0 201.257 198.4A68.12 68.12 0 0 0 94.2 269.59C41.924 271.106 0 313.728 0 366.12a96.054 96.054 0 0 0 1.029 13.958a4.508 4.508 0 0 0 4.445 3.871l426.1.051c.043 0 .08-.019.122-.02a5.606 5.606 0 0 0 5.271-4l3.273-11.265c3.9-13.4 2.448-25.8-4.1-34.9c-6.016-8.392-16.05-13.328-28.234-13.902m105.95-98.813c-2.141 0-4.271.062-6.391.164a3.771 3.771 0 0 0-3.324 2.653l-9.077 31.193c-3.9 13.4-2.449 25.786 4.1 34.89c6.02 8.4 16.054 13.323 28.238 13.9l49.2 2.939a4.491 4.491 0 0 1 3.51 1.894a4.64 4.64 0 0 1 .514 4.169a6.153 6.153 0 0 1-5.351 4.075l-51.125 2.939c-27.754 1.27-57.669 23.574-68.145 50.784l-3.695 9.606a2.716 2.716 0 0 0 2.427 3.68c.046 0 .088.017.136.017h175.91a4.69 4.69 0 0 0 4.539-3.37a124.807 124.807 0 0 0 4.682-34C640 277.3 583.524 221.1 513.856 221.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloudscale(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m318.1 154l-9.4 7.6c-22.5-19.3-51.5-33.6-83.3-33.6C153.8 128 96 188.8 96 260.3c0 6.6.4 13.1 1.4 19.4c-2-56 41.8-97.4 92.6-97.4c24.2 0 46.2 9.4 62.6 24.7l-25.2 20.4c-8.3-.9-16.8 1.8-23.1 8.1c-11.1 11-11.1 28.9 0 40c11.1 11 28.9 11 40 0c6.3-6.3 9-14.9 8.1-23.1l75.2-88.8c6.3-6.5-3.3-15.9-9.5-9.6m-83.8 111.5c-5.6 5.5-14.6 5.5-20.2 0c-5.6-5.6-5.6-14.6 0-20.2s14.6-5.6 20.2 0s5.6 14.7 0 20.2M224 32C100.5 32 0 132.5 0 256s100.5 224 224 224s224-100.5 224-224S347.5 32 224 32m0 384c-88.2 0-160-71.8-160-160S135.8 96 224 96s160 71.8 160 160s-71.8 160-160 160"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloudsmith(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M332.5 419.9c0 46.4-37.6 84.1-84 84.1s-84-37.7-84-84.1s37.6-84 84-84s84 37.6 84 84m-84-243.9c46.4 0 80-37.6 80-84s-33.6-84-80-84s-88 37.6-88 84s-29.6 76-76 76s-84 41.6-84 88s37.6 80 84 80s84-33.6 84-80s33.6-80 80-80"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloudversify(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M148.6 304c8.2 68.5 67.4 115.5 146 111.3c51.2 43.3 136.8 45.8 186.4-5.6c69.2 1.1 118.5-44.6 131.5-99.5c14.8-62.5-18.2-132.5-92.1-155.1c-33-88.1-131.4-101.5-186.5-85c-57.3 17.3-84.3 53.2-99.3 109.7c-7.8 2.7-26.5 8.9-45 24.1c11.7 0 15.2 8.9 15.2 19.5v20.4c0 10.7-8.7 19.5-19.5 19.5h-20.2c-10.7 0-19.5-6-19.5-16.7V240H98.8C95 240 88 244.3 88 251.9v40.4c0 6.4 5.3 11.8 11.7 11.8h48.9zm227.4 8c-10.7 46.3 21.7 72.4 55.3 86.8C324.1 432.6 259.7 348 296 288c-33.2 21.6-33.7 71.2-29.2 92.9c-17.9-12.4-53.8-32.4-57.4-79.8c-3-39.9 21.5-75.7 57-93.9C297 191.4 369.9 198.7 400 248c-14.1-48-53.8-70.1-101.8-74.8c30.9-30.7 64.4-50.3 114.2-43.7c69.8 9.3 133.2 82.8 67.7 150.5c35-16.3 48.7-54.4 47.5-76.9l10.5 19.6c11.8 22 15.2 47.6 9.4 72c-9.2 39-40.6 68.8-79.7 76.5c-32.1 6.3-83.1-5.1-91.8-59.2M128 208H88.2c-8.9 0-16.2-7.3-16.2-16.2v-39.6c0-8.9 7.3-16.2 16.2-16.2H128c8.9 0 16.2 7.3 16.2 16.2v39.6c0 8.9-7.3 16.2-16.2 16.2M10.1 168C4.5 168 0 163.5 0 157.9v-27.8c0-5.6 4.5-10.1 10.1-10.1h27.7c5.5 0 10.1 4.5 10.1 10.1v27.8c0 5.6-4.5 10.1-10.1 10.1zM168 142.7v-21.4c0-5.1 4.2-9.3 9.3-9.3h21.4c5.1 0 9.3 4.2 9.3 9.3v21.4c0 5.1-4.2 9.3-9.3 9.3h-21.4c-5.1 0-9.3-4.2-9.3-9.3M56 235.5v25c0 6.3-5.1 11.5-11.4 11.5H19.4C13.1 272 8 266.8 8 260.5v-25c0-6.3 5.1-11.5 11.4-11.5h25.1c6.4 0 11.5 5.2 11.5 11.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Codepen(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m502.285 159.704l-234-156c-7.987-4.915-16.511-4.96-24.571 0l-234 156C3.714 163.703 0 170.847 0 177.989v155.999c0 7.143 3.714 14.286 9.715 18.286l234 156.022c7.987 4.915 16.511 4.96 24.571 0l234-156.022c6-3.999 9.715-11.143 9.715-18.286V177.989c-.001-7.142-3.715-14.286-9.716-18.285M278 63.131l172.286 114.858l-76.857 51.429L278 165.703zm-44 0v102.572l-95.429 63.715l-76.857-51.429zM44 219.132l55.143 36.857L44 292.846zm190 229.715L61.714 333.989l76.857-51.429L234 346.275zm22-140.858l-77.715-52l77.715-52l77.715 52zm22 140.858V346.275l95.429-63.715l76.857 51.429zm190-156.001l-55.143-36.857L468 219.132z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Codiepie(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M422.5 202.9c30.7 0 33.5 53.1-.3 53.1h-10.8v44.3h-26.6v-97.4zM472 352.6C429.9 444.5 350.4 504 248 504C111 504 0 393 0 256S111 8 248 8c97.4 0 172.8 53.7 218.2 138.4l-186 108.8zm-38.5 12.5l-60.3-30.7c-27.1 44.3-70.4 71.4-122.4 71.4c-82.5 0-149.2-66.7-149.2-148.9c0-82.5 66.7-149.2 149.2-149.2c48.4 0 88.9 23.5 116.9 63.4l59.5-34.6c-40.7-62.6-104.7-100-179.2-100c-121.2 0-219.5 98.3-219.5 219.5S126.8 475.5 248 475.5c78.6 0 146.5-42.1 185.5-110.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Confluence(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.3 412.2c-4.5 7.6-2.1 17.5 5.5 22.2l105.9 65.2c7.7 4.7 17.7 2.4 22.4-5.3c0-.1.1-.2.1-.2c67.1-112.2 80.5-95.9 280.9-.7c8.1 3.9 17.8.4 21.7-7.7c.1-.1.1-.3.2-.4l50.4-114.1c3.6-8.1-.1-17.6-8.1-21.3c-22.2-10.4-66.2-31.2-105.9-50.3C127.5 179 44.6 345.3 2.3 412.2m507.4-312.1c4.5-7.6 2.1-17.5-5.5-22.2L398.4 12.8c-7.5-5-17.6-3.1-22.6 4.4c-.2.3-.4.6-.6 1c-67.3 112.6-81.1 95.6-280.6.9c-8.1-3.9-17.8-.4-21.7 7.7c-.1.1-.1.3-.2.4L22.2 141.3c-3.6 8.1.1 17.6 8.1 21.3c22.2 10.4 66.3 31.2 106 50.4c248 120 330.8-45.4 373.4-112.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Connectdevelop(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m550.5 241l-50.089-86.786c1.071-2.142 1.875-4.553 1.875-7.232c0-8.036-6.696-14.733-14.732-15.001l-55.447-95.893c.536-1.607 1.071-3.214 1.071-4.821c0-8.571-6.964-15.268-15.268-15.268c-4.821 0-8.839 2.143-11.786 5.625H299.518C296.839 18.143 292.821 16 288 16s-8.839 2.143-11.518 5.625H170.411C167.464 18.143 163.447 16 158.625 16c-8.303 0-15.268 6.696-15.268 15.268c0 1.607.536 3.482 1.072 4.821l-55.983 97.233c-5.356 2.41-9.107 7.5-9.107 13.661c0 .535.268 1.071.268 1.607l-53.304 92.143c-7.232 1.339-12.59 7.5-12.59 15c0 7.232 5.089 13.393 12.054 15l55.179 95.358c-.536 1.607-.804 2.946-.804 4.821c0 7.232 5.089 13.393 12.054 14.732l51.697 89.732c-.536 1.607-1.071 3.482-1.071 5.357c0 8.571 6.964 15.268 15.268 15.268c4.821 0 8.839-2.143 11.518-5.357h106.875C279.161 493.857 283.447 496 288 496s8.839-2.143 11.518-5.357h107.143c2.678 2.946 6.696 4.821 10.982 4.821c8.571 0 15.268-6.964 15.268-15.268c0-1.607-.267-2.946-.803-4.285l51.697-90.268c6.964-1.339 12.054-7.5 12.054-14.732c0-1.607-.268-3.214-.804-4.821l54.911-95.358c6.964-1.339 12.322-7.5 12.322-15c-.002-7.232-5.092-13.393-11.788-14.732M153.535 450.732l-43.66-75.803h43.66zm0-83.839h-43.66c-.268-1.071-.804-2.142-1.339-3.214l44.999-47.41zm0-62.411l-50.357 53.304c-1.339-.536-2.679-1.34-4.018-1.607L43.447 259.75c.535-1.339.535-2.679.535-4.018s0-2.41-.268-3.482l51.965-90c2.679-.268 5.357-1.072 7.768-2.679l50.089 51.965v92.946zm0-102.322l-45.803-47.41c1.339-2.143 2.143-4.821 2.143-7.767c0-.268-.268-.804-.268-1.072l43.928-15.804zm0-80.625l-43.66 15.804l43.66-75.536zm326.519 39.108l.804 1.339L445.5 329.125l-63.75-67.232l98.036-101.518zM291.75 355.107l11.518 11.786H280.5zm-.268-11.25l-83.303-85.446l79.553-84.375l83.036 87.589zm5.357 5.893l79.286-82.232l67.5 71.25l-5.892 28.125H313.714zM410.411 44.393c1.071.536 2.142 1.072 3.482 1.34l57.857 100.714v.536c0 2.946.803 5.624 2.143 7.767L376.393 256l-83.035-87.589zm-9.107-2.143L287.732 162.518l-57.054-60.268l166.339-60zm-123.483 0c2.678 2.678 6.16 4.285 10.179 4.285s7.5-1.607 10.179-4.285h75L224.786 95.821L173.893 42.25zm-116.249 5.625l1.071-2.142a33.834 33.834 0 0 0 2.679-.804l51.161 53.84l-54.911 19.821zm0 79.286l60.803-21.964l59.732 63.214l-79.553 84.107l-40.982-42.053zm0 92.678L198 257.607l-36.428 38.304zm0 87.858l42.053-44.464l82.768 85.982l-17.143 17.678H161.572zm6.964 162.053c-1.607-1.607-3.482-2.678-5.893-3.482l-1.071-1.607v-89.732h99.91l-91.607 94.821zm129.911 0c-2.679-2.41-6.428-4.285-10.447-4.285s-7.767 1.875-10.447 4.285h-96.429l91.607-94.821h38.304l91.607 94.821zm120-11.786l-4.286 7.5c-1.339.268-2.41.803-3.482 1.339l-89.196-91.875h114.376zm12.856-22.232l12.858-60.803h21.964zm34.822-68.839h-20.357l4.553-21.16l17.143 18.214c-.535.803-1.071 1.874-1.339 2.946m66.161-107.411l-55.447 96.697c-1.339.535-2.679 1.071-4.018 1.874l-20.625-21.964l34.554-163.928l45.803 79.286c-.267 1.339-.803 2.678-.803 4.285c0 1.339.268 2.411.536 3.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Contao(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M45.4 305c14.4 67.1 26.4 129 68.2 175H34c-18.7 0-34-15.2-34-34V66c0-18.7 15.2-34 34-34h57.7C77.9 44.6 65.6 59.2 54.8 75.6c-45.4 70-27 146.8-9.4 229.4M478 32h-90.2c21.4 21.4 39.2 49.5 52.7 84.1l-137.1 29.3c-14.9-29-37.8-53.3-82.6-43.9c-24.6 5.3-41 19.3-48.3 34.6c-8.8 18.7-13.2 39.8 8.2 140.3c21.1 100.2 33.7 117.7 49.5 131.2c12.9 11.1 33.4 17 58.3 11.7c44.5-9.4 55.7-40.7 57.4-73.2l137.4-29.6c3.2 71.5-18.7 125.2-57.4 163.6H478c18.7 0 34-15.2 34-34V66c0-18.8-15.2-34-34-34"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CottonBureau(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M474.31 330.41c-23.66 91.85-94.23 144.59-201.9 148.35V429.6c0-48 26.41-74.39 74.39-74.39c62 0 99.2-37.2 99.2-99.21c0-61.37-36.53-98.28-97.38-99.06c-33-69.32-146.5-64.65-177.24 0C110.52 157.72 74 194.63 74 256c0 62.13 37.27 99.41 99.4 99.41c48 0 74.55 26.23 74.55 74.39V479c-134.43-5-211.1-85.07-211.1-223c0-141.82 81.35-223.2 223.2-223.2c114.77 0 189.84 53.2 214.69 148.81H500C473.88 71.51 388.22 8 259.82 8C105 8 12 101.19 12 255.82C12 411.14 105.19 504.34 259.82 504c128.27 0 213.87-63.81 239.67-173.59zM357 182.33c41.37 3.45 64.2 29 64.2 73.67c0 48-26.43 74.41-74.4 74.41c-28.61 0-49.33-9.59-61.59-27.33c83.06-16.55 75.59-99.67 71.79-120.75m-81.68 97.36c-2.46-10.34-16.33-87 56.23-97c2.27 10.09 16.52 87.11-56.26 97zM260 132c28.61 0 49 9.67 61.44 27.61c-28.36 5.48-49.36 20.59-61.59 43.45c-12.23-22.86-33.23-38-61.6-43.45c12.41-17.69 33.27-27.35 61.57-27.35zm-71.52 50.72c73.17 10.57 58.91 86.81 56.49 97c-72.41-9.84-59-86.95-56.25-97zM173.2 330.41c-48 0-74.4-26.4-74.4-74.41c0-44.36 22.86-70 64.22-73.67c-6.75 37.2-1.38 106.53 71.65 120.75c-12.14 17.63-32.84 27.3-61.14 27.3zm53.21 12.39A80.8 80.8 0 0 0 260 309.25c7.77 14.49 19.33 25.54 33.82 33.55a80.28 80.28 0 0 0-33.58 33.83c-8-14.5-19.07-26.23-33.56-33.83z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cpanel(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M210.3 220.2c-5.6-24.8-26.9-41.2-51-41.2h-37c-7.1 0-12.5 4.5-14.3 10.9L73.1 320l24.7-.1c6.8 0 12.3-4.5 14.2-10.7l25.8-95.7h19.8c8.4 0 16.2 5.6 18.3 14.8c2.5 10.9-5.9 22.6-18.3 22.6h-10.3c-7 0-12.5 4.6-14.3 10.8l-6.4 23.8h32c37.2 0 58.3-36.2 51.7-65.3m-156.5 28h18.6c6.9 0 12.4-4.4 14.3-10.9l6.2-23.6h-40C30 213.7 9 227.8 1.7 254.8C-7 288.6 18.5 320 52 320h12.4l7.1-26.1c1.2-4.4-2.2-8.3-6.4-8.3H53.8c-24.7 0-24.9-37.4 0-37.4m247.5-34.8h-77.9l-3.5 13.4c-2.4 9.6 4.5 18.5 14.2 18.5h57.5c4 0 2.4 4.3 2.1 5.3l-8.6 31.8c-.4 1.4-.9 5.3-5.5 5.3h-34.9c-5.3 0-5.3-7.9 0-7.9h21.6c6.8 0 12.3-4.6 14.2-10.8l3.5-13.2h-48.4c-39.2 0-43.6 63.8-.7 63.8l57.5.2c11.2 0 20.6-7.2 23.4-17.8l14-51.8c4.8-19.2-9.7-36.8-28.5-36.8M633.1 179h-18.9c-4.9 0-9.2 3.2-10.4 7.9L568.2 320c20.7 0 39.8-13.8 44.9-34.5l26.5-98.2c1.2-4.3-2-8.3-6.5-8.3m-236.3 34.7v.1h-48.3l-26.2 98c-1.2 4.4 2.2 8.3 6.4 8.3h18.9c4.8 0 9.2-3 10.4-7.8l17.2-64H395c12.5 0 21.4 11.8 18.1 23.4l-10.6 40c-1.2 4.3 1.9 8.3 6.4 8.3H428c4.6 0 9.1-2.9 10.3-7.8l8.8-33.1c9-33.1-15.9-65.4-50.3-65.4m98.3 74.6c-3.6 0-6-3.4-5.1-6.7l8-30c.9-3.9 3.7-6 7.8-6h32.9c2.6 0 4.6 2.4 3.9 5.1l-.7 2.6c-.6 2-1.9 3-3.9 3h-21.6c-7 0-12.6 4.6-14.2 10.8l-3.5 13h53.4c10.5 0 20.3-6.6 23.2-17.6l3.2-12c4.9-19.1-9.3-36.8-28.3-36.8h-47.3c-17.9 0-33.8 12-38.6 29.6l-10.8 40c-5 17.7 8.3 36.7 28.3 36.7h66.7c6.8 0 12.3-4.5 14.2-10.7l5.7-21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreativeCommons(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m245.83 214.87l-33.22 17.28c-9.43-19.58-25.24-19.93-27.46-19.93c-22.13 0-33.22 14.61-33.22 43.84c0 23.57 9.21 43.84 33.22 43.84c14.47 0 24.65-7.09 30.57-21.26l30.55 15.5c-6.17 11.51-25.69 38.98-65.1 38.98c-22.6 0-73.96-10.32-73.96-77.05c0-58.69 43-77.06 72.63-77.06c30.72-.01 52.7 11.95 65.99 35.86m143.05 0l-32.78 17.28c-9.5-19.77-25.72-19.93-27.9-19.93c-22.14 0-33.22 14.61-33.22 43.84c0 23.55 9.23 43.84 33.22 43.84c14.45 0 24.65-7.09 30.54-21.26l31 15.5c-2.1 3.75-21.39 38.98-65.09 38.98c-22.69 0-73.96-9.87-73.96-77.05c0-58.67 42.97-77.06 72.63-77.06c30.71-.01 52.58 11.95 65.56 35.86M247.56 8.05C104.74 8.05 0 123.11 0 256.05c0 138.49 113.6 248 247.56 248c129.93 0 248.44-100.87 248.44-248c0-137.87-106.62-248-248.44-248m.87 450.81c-112.54 0-203.7-93.04-203.7-202.81c0-105.42 85.43-203.27 203.72-203.27c112.53 0 202.82 89.46 202.82 203.26c-.01 121.69-99.68 202.82-202.84 202.82"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreativeCommonsBy(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M314.9 194.4v101.4h-28.3v120.5h-77.1V295.9h-28.3V194.4c0-4.4 1.6-8.2 4.6-11.3c3.1-3.1 6.9-4.7 11.3-4.7H299c4.1 0 7.8 1.6 11.1 4.7c3.1 3.2 4.8 6.9 4.8 11.3m-101.5-63.7c0-23.3 11.5-35 34.5-35s34.5 11.7 34.5 35c0 23-11.5 34.5-34.5 34.5s-34.5-11.5-34.5-34.5M247.6 8C389.4 8 496 118.1 496 256c0 147.1-118.5 248-248.4 248C113.6 504 0 394.5 0 256C0 123.1 104.7 8 247.6 8m.8 44.7C130.2 52.7 44.7 150.6 44.7 256c0 109.8 91.2 202.8 203.7 202.8c103.2 0 202.8-81.1 202.8-202.8c.1-113.8-90.2-203.3-202.8-203.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreativeCommonsNc(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M247.6 8C387.4 8 496 115.9 496 256c0 147.2-118.5 248-248.4 248C113.1 504 0 393.2 0 256C0 123.1 104.7 8 247.6 8M55.8 189.1c-7.4 20.4-11.1 42.7-11.1 66.9c0 110.9 92.1 202.4 203.7 202.4c122.4 0 177.2-101.8 178.5-104.1l-93.4-41.6c-7.7 37.1-41.2 53-68.2 55.4v38.1h-28.8V368c-27.5-.3-52.6-10.2-75.3-29.7l34.1-34.5c31.7 29.4 86.4 31.8 86.4-2.2c0-6.2-2.2-11.2-6.6-15.1c-14.2-6-1.8-.1-219.3-97.4M248.4 52.3c-38.4 0-112.4 8.7-170.5 93l94.8 42.5c10-31.3 40.4-42.9 63.8-44.3v-38.1h28.8v38.1c22.7 1.2 43.4 8.9 62 23L295 199.7c-42.7-29.9-83.5-8-70 11.1c53.4 24.1 43.8 19.8 93 41.6l127.1 56.7c4.1-17.4 6.2-35.1 6.2-53.1c0-57-19.8-105-59.3-143.9c-39.3-39.9-87.2-59.8-143.6-59.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreativeCommonsNcEu(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M247.7 8C103.6 8 0 124.8 0 256c0 136.3 111.7 248 247.7 248C377.9 504 496 403.1 496 256C496 117 388.4 8 247.7 8m.6 450.7c-112 0-203.6-92.5-203.6-202.7c0-23.2 3.7-45.2 10.9-66l65.7 29.1h-4.7v29.5h23.3c0 6.2-.4 3.2-.4 19.5h-22.8v29.5h27c11.4 67 67.2 101.3 124.6 101.3c26.6 0 50.6-7.9 64.8-15.8l-10-46.1c-8.7 4.6-28.2 10.8-47.3 10.8c-28.2 0-58.1-10.9-67.3-50.2h90.3l128.3 56.8c-1.5 2.1-56.2 104.3-178.8 104.3m-16.7-190.6l-.5-.4l.9.4zm77.2-19.5h3.7v-29.5h-70.3l-28.6-12.6c2.5-5.5 5.4-10.5 8.8-14.3c12.9-15.8 31.1-22.4 51.1-22.4c18.3 0 35.3 5.4 46.1 10l11.6-47.3c-15-6.6-37-12.4-62.3-12.4c-39 0-72.2 15.8-95.9 42.3c-5.3 6.1-9.8 12.9-13.9 20.1l-81.6-36.1c64.6-96.8 157.7-93.6 170.7-93.6c113 0 203 90.2 203 203.4c0 18.7-2.1 36.3-6.3 52.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreativeCommonsNcJp(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M247.7 8C103.6 8 0 124.8 0 256c0 136.4 111.8 248 247.7 248C377.9 504 496 403.2 496 256C496 117.2 388.5 8 247.7 8m.6 450.7c-112 0-203.6-92.5-203.6-202.7c0-21.1 3-41.2 9-60.3l127 56.5h-27.9v38.6h58.1l5.7 11.8v18.7h-63.8V360h63.8v56h61.7v-56h64.2v-35.7l81 36.1c-1.5 2.2-57.1 98.3-175.2 98.3m87.6-137.3h-57.6v-18.7l2.9-5.6zm6.5-51.4v-17.8h-38.6l63-116H301l-43.4 96l-23-10.2l-39.6-85.7h-65.8l27.3 51l-81.9-36.5c27.8-44.1 82.6-98.1 173.7-98.1c112.8 0 203 90 203 203.4c0 21-2.7 40.6-7.9 59z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreativeCommonsNd(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M247.6 8C389.4 8 496 118.1 496 256c0 147.1-118.5 248-248.4 248C113.6 504 0 394.5 0 256C0 123.1 104.7 8 247.6 8m.8 44.7C130.2 52.7 44.7 150.6 44.7 256c0 109.8 91.2 202.8 203.7 202.8c103.2 0 202.8-81.1 202.8-202.8c.1-113.8-90.2-203.3-202.8-203.3m94 144.3v42.5H162.1V197zm0 79.8v42.5H162.1v-42.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreativeCommonsPd(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M248 8C111 8 0 119.1 0 256c0 137 111 248 248 248s248-111 248-248C496 119.1 385 8 248 8m0 449.5c-139.2 0-235.8-138-190.2-267.9l78.8 35.1c-2.1 10.5-3.3 21.5-3.3 32.9c0 99 73.9 126.9 120.4 126.9c22.9 0 53.5-6.7 79.4-29.5L297 311.1c-5.5 6.3-17.6 16.7-36.3 16.7c-37.8 0-53.7-39.9-53.9-71.9c230.4 102.6 216.5 96.5 217.9 96.8c-34.3 62.4-100.6 104.8-176.7 104.8m194.2-150l-224-100c18.8-34 54.9-30.7 74.7-11l40.4-41.6c-27.1-23.3-58-27.5-78.1-27.5c-47.4 0-80.9 20.5-100.7 51.6l-74.9-33.4c36.1-54.9 98.1-91.2 168.5-91.2c111.1 0 201.5 90.4 201.5 201.5c0 18-2.4 35.4-6.8 52c-.3-.1-.4-.2-.6-.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreativeCommonsPdAlt(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M247.6 8C104.7 8 0 123.1 0 256c0 138.5 113.6 248 247.6 248C377.5 504 496 403.1 496 256C496 118.1 389.4 8 247.6 8m.8 450.8c-112.5 0-203.7-93-203.7-202.8c0-105.4 85.5-203.3 203.7-203.3c112.6 0 202.9 89.5 202.8 203.3c0 121.7-99.6 202.8-202.8 202.8M316.7 186h-53.2v137.2h53.2c21.4 0 70-5.1 70-68.6c0-63.4-48.6-68.6-70-68.6m.8 108.5h-19.9v-79.7l19.4-.1c3.8 0 35-2.1 35 39.9c0 24.6-10.5 39.9-34.5 39.9M203.7 186h-68.2v137.3h34.6V279h27c54.1 0 57.1-37.5 57.1-46.5c0-31-16.8-46.5-50.5-46.5m-4.9 67.3h-29.2v-41.6h28.3c30.9 0 28.8 41.6.9 41.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreativeCommonsRemix(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M247.6 8C389.4 8 496 118.1 496 256c0 147.1-118.5 248-248.4 248C113.6 504 0 394.5 0 256C0 123.1 104.7 8 247.6 8m.8 44.7C130.2 52.7 44.7 150.6 44.7 256c0 109.8 91.2 202.8 203.7 202.8c103.2 0 202.8-81.1 202.8-202.8c.1-113.8-90.2-203.3-202.8-203.3m161.7 207.7l4.9 2.2v70c-7.2 3.6-63.4 27.5-67.3 28.8c-6.5-1.8-113.7-46.8-137.3-56.2l-64.2 26.6l-63.3-27.5v-63.8l59.3-24.8c-.7-.7-.4 5-.4-70.4l67.3-29.7L361 178.5v61.6zm-70.4 81.5v-43.8h-.4v-1.8l-113.8-46.5V295l113.8 46.9v-.4zm7.5-57.6l39.9-16.4l-36.8-15.5l-39 16.4zm52.3 38.1v-43L355.2 298v43.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreativeCommonsSa(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M247.6 8C389.4 8 496 118.1 496 256c0 147.1-118.5 248-248.4 248C113.6 504 0 394.5 0 256C0 123.1 104.7 8 247.6 8m.8 44.7C130.2 52.7 44.7 150.6 44.7 256c0 109.8 91.2 202.8 203.7 202.8c103.2 0 202.8-81.1 202.8-202.8c.1-113.8-90.2-203.3-202.8-203.3M137.7 221c13-83.9 80.5-95.7 108.9-95.7c99.8 0 127.5 82.5 127.5 134.2c0 63.6-41 132.9-128.9 132.9c-38.9 0-99.1-20-109.4-97h62.5c1.5 30.1 19.6 45.2 54.5 45.2c23.3 0 58-18.2 58-82.8c0-82.5-49.1-80.6-56.7-80.6c-33.1 0-51.7 14.6-55.8 43.8h18.2l-49.2 49.2l-49-49.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreativeCommonsSampling(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M247.6 8C389.4 8 496 118.1 496 256c0 147.1-118.5 248-248.4 248C113.6 504 0 394.5 0 256C0 123.1 104.7 8 247.6 8m.8 44.7C130.2 52.7 44.7 150.6 44.7 256c0 109.8 91.2 202.8 203.7 202.8c103.2 0 202.8-81.1 202.8-202.8c.1-113.8-90.2-203.3-202.8-203.3m3.6 53.2c2.8-.3 11.5 1 11.5 11.5l6.6 107.2l4.9-59.3c0-6 4.7-10.6 10.6-10.6c5.9 0 10.6 4.7 10.6 10.6c0 2.5-.5-5.7 5.7 81.5l5.8-64.2c.3-2.9 2.9-9.3 10.2-9.3c3.8 0 9.9 2.3 10.6 8.9l11.5 96.5l5.3-12.8c1.8-4.4 5.2-6.6 10.2-6.6h58v21.3h-50.9l-18.2 44.3c-3.9 9.9-19.5 9.1-20.8-3.1l-4-31.9l-7.5 92.6c-.3 3-3 9.3-10.2 9.3c-3 0-9.8-2.1-10.6-9.3c0-1.9.6 5.8-6.2-77.9l-5.3 72.2c-1.1 4.8-4.8 9.3-10.6 9.3c-2.9 0-9.8-2-10.6-9.3c0-1.9.5 6.7-5.8-87.7l-5.8 94.8c0 6.3-3.6 12.4-10.6 12.4c-5.2 0-10.6-4.1-10.6-12l-5.8-87.7c-5.8 92.5-5.3 84-5.3 85.9c-1.1 4.8-4.8 9.3-10.6 9.3c-3 0-9.8-2.1-10.6-9.3c0-.7-.4-1.1-.4-2.6l-6.2-88.6L182 348c-.7 6.5-6.7 9.3-10.6 9.3c-5.8 0-9.6-4.1-10.6-8.9L149.7 272c-2 4-3.5 8.4-11.1 8.4H87.2v-21.3H132l13.7-27.9c4.4-9.9 18.2-7.2 19.9 2.7l3.1 20.4l8.4-97.9c0-6 4.8-10.6 10.6-10.6c.5 0 10.6-.2 10.6 12.4l4.9 69.1l6.6-92.6c0-10.1 9.5-10.6 10.2-10.6c.6 0 10.6.7 10.6 10.6l5.3 80.6l6.2-97.9c.1-1.1-.6-10.3 9.9-11.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreativeCommonsSamplingPlus(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M247.6 8C389.4 8 496 118.1 496 256c0 147.1-118.5 248-248.4 248C113.6 504 0 394.5 0 256C0 123.1 104.7 8 247.6 8m.8 44.7C130.2 52.7 44.7 150.6 44.7 256c0 109.8 91.2 202.8 203.7 202.8c103.2 0 202.8-81.1 202.8-202.8c.1-113.8-90.2-203.3-202.8-203.3m107 205.6c-4.7 0-9 2.8-10.7 7.2l-4 9.5l-11-92.8c-1.7-13.9-22-13.4-23.1.4l-4.3 51.4l-5.2-68.8c-1.1-14.3-22.1-14.2-23.2 0l-3.5 44.9l-5.9-94.3c-.9-14.5-22.3-14.4-23.2 0l-5.1 83.7l-4.3-66.3c-.9-14.4-22.2-14.4-23.2 0l-5.3 80.2l-4.1-57c-1.1-14.3-22-14.3-23.2-.2l-7.7 89.8l-1.8-12.2c-1.7-11.4-17.1-13.6-22-3.3l-13.2 27.7H87.5v23.2h51.3c4.4 0 8.4-2.5 10.4-6.4l10.7 73.1c2 13.5 21.9 13 23.1-.7l3.8-43.6l5.7 78.3c1.1 14.4 22.3 14.2 23.2-.1l4.6-70.4l4.8 73.3c.9 14.4 22.3 14.4 23.2-.1l4.9-80.5l4.5 71.8c.9 14.3 22.1 14.5 23.2.2l4.6-58.6l4.9 64.4c1.1 14.3 22 14.2 23.1.1l6.8-83l2.7 22.3c1.4 11.8 17.7 14.1 22.3 3.1l18-43.4h50.5V258zm-78 5.2h-21.9v21.9c0 4.1-3.3 7.5-7.5 7.5c-4.1 0-7.5-3.3-7.5-7.5v-21.9h-21.9c-4.1 0-7.5-3.3-7.5-7.5c0-4.1 3.4-7.5 7.5-7.5h21.9v-21.9c0-4.1 3.4-7.5 7.5-7.5s7.5 3.3 7.5 7.5v21.9h21.9c4.1 0 7.5 3.3 7.5 7.5c0 4.1-3.4 7.5-7.5 7.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreativeCommonsShare(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M247.6 8C389.4 8 496 118.1 496 256c0 147.1-118.5 248-248.4 248C113.6 504 0 394.5 0 256C0 123.1 104.7 8 247.6 8m.8 44.7C130.2 52.7 44.7 150.6 44.7 256c0 109.8 91.2 202.8 203.7 202.8c103.2 0 202.8-81.1 202.8-202.8c.1-113.8-90.2-203.3-202.8-203.3m101 132.4c7.8 0 13.7 6.1 13.7 13.7v182.5c0 7.7-6.1 13.7-13.7 13.7H214.3c-7.7 0-13.7-6-13.7-13.7v-54h-54c-7.8 0-13.7-6-13.7-13.7V131.1c0-8.2 6.6-12.7 12.4-13.7h136.4c7.7 0 13.7 6 13.7 13.7v54zM159.9 300.3h40.7V198.9c0-7.4 5.8-12.6 12-13.7h55.8v-40.3H159.9zm176.2-88.1H227.6v155.4h108.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreativeCommonsZero(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M247.6 8C389.4 8 496 118.1 496 256c0 147.1-118.5 248-248.4 248C113.6 504 0 394.5 0 256C0 123.1 104.7 8 247.6 8m.8 44.7C130.2 52.7 44.7 150.6 44.7 256c0 109.8 91.2 202.8 203.7 202.8c103.2 0 202.8-81.1 202.8-202.8c.1-113.8-90.2-203.3-202.8-203.3m-.4 60.5c-81.9 0-102.5 77.3-102.5 142.8c0 65.5 20.6 142.8 102.5 142.8S350.5 321.5 350.5 256c0-65.5-20.6-142.8-102.5-142.8m0 53.9c3.3 0 6.4.5 9.2 1.2c5.9 5.1 8.8 12.1 3.1 21.9l-54.5 100.2c-1.7-12.7-1.9-25.1-1.9-34.4c0-28.8 2-88.9 44.1-88.9m40.8 46.2c2.9 15.4 3.3 31.4 3.3 42.7c0 28.9-2 88.9-44.1 88.9c-13.5 0-32.6-7.7-20.1-26.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CriticalRole(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M225.82 0c.26.15 216.57 124.51 217.12 124.72c3 1.18 3.7 3.46 3.7 6.56q-.11 125.17 0 250.36a5.88 5.88 0 0 1-3.38 5.78c-21.37 12-207.86 118.29-218.93 124.58h-3C142 466.34 3.08 386.56 2.93 386.48a3.29 3.29 0 0 1-1.88-3.24c0-.87 0-225.94-.05-253.1a5 5 0 0 1 2.93-4.93C27.19 112.11 213.2 6 224.07 0zM215.4 20.42l-.22-.16Q118.06 75.55 21 130.87c0 .12.08.23.13.35l30.86 11.64c-7.71 6-8.32 6-10.65 5.13c-.1 0-24.17-9.28-26.8-10v230.43c.88-1.41 64.07-110.91 64.13-111c1.62-2.82 3-1.92 9.12-1.52c1.4.09 1.48.22.78 1.42c-41.19 71.33-36.4 63-67.48 116.94c-.81 1.4-.61 1.13 1.25 1.13h186.5c1.44 0 1.69-.23 1.7-1.64v-8.88c0-1.34 2.36-.81-18.37-1c-7.46-.07-14.14-3.22-21.38-12.7c-7.38-9.66-14.62-19.43-21.85-29.21c-2.28-3.08-3.45-2.38-16.76-2.38c-1.75 0-1.78 0-1.76 1.82c.29 26.21.15 25.27 1 32.66c.52 4.37 2.16 4.2 9.69 4.81c3.14.26 3.88 4.08.52 4.92c-1.57.39-31.6.51-33.67-.1a2.42 2.42 0 0 1 .3-4.73c3.29-.76 6.16.81 6.66-4.44c1.3-13.66 1.17-9 1.1-79.42c0-10.82-.35-12.58-5.36-13.55c-1.22-.24-3.54-.16-4.69-.55c-2.88-1-2-4.84 1.77-4.85c33.67 0 46.08-1.07 56.06 4.86c7.74 4.61 12 11.48 12.51 20.4c.88 14.59-6.51 22.35-15 32.59a1.46 1.46 0 0 0 0 2.22c2.6 3.25 5 6.63 7.71 9.83c27.56 33.23 24.11 30.54 41.28 33.06c.89.13 1-.42 1-1.15v-11c0-1 .32-1.43 1.41-1.26a72.37 72.37 0 0 0 23.58-.3c1.08-.15 1.5.2 1.48 1.33c0 .11.88 26.69.87 26.8c-.05 1.52.67 1.62 1.89 1.62h186.71Q386.51 304.6 346 234.33c2.26-.66-.4 0 6.69-1.39c2-.39 2.05-.41 3.11 1.44c7.31 12.64 77.31 134 77.37 134.06V138c-1.72.5-103.3 38.72-105.76 39.68c-1.08.42-1.55.2-1.91-.88c-.63-1.9-1.34-3.76-2.09-5.62c-.32-.79-.09-1.13.65-1.39c.1 0 95.53-35.85 103-38.77c-65.42-37.57-130.56-75-196-112.6l86.82 150.39l-.28.33c-9.57-.9-10.46-1.6-11.8-3.94c-1-1.69-73.5-127.71-82-142.16c-9.1 14.67-83.56 146.21-85.37 146.32c-2.93.17-5.88.08-9.25.08q43.25-74.74 86.18-149zm51.93 129.92a37.68 37.68 0 0 0 5.54-.85c1.69-.3 2.53.2 2.6 1.92c0 .11.07 19.06-.86 20.45s-1.88 1.22-2.6-.19c-5-9.69 6.22-9.66-39.12-12c-.7 0-1 .23-1 .93c0 .13 3.72 122 3.73 122.11c0 .89.52 1.2 1.21 1.51a83.92 83.92 0 0 1 8.7 4.05c7.31 4.33 11.38 10.84 12.41 19.31c1.44 11.8-2.77 35.77-32.21 37.14c-2.75.13-28.26 1.08-34.14-23.25c-4.66-19.26 8.26-32.7 19.89-36.4a2.45 2.45 0 0 0 2-2.66c.1-5.63 3-107.1 3.71-121.35c.05-1.08-.62-1.16-1.35-1.15c-32.35.52-36.75-.34-40.22 8.52c-2.42 6.18-4.14 1.32-3.95.23q1.59-9 3.31-18c.4-2.11 1.43-2.61 3.43-1.86c5.59 2.11 6.72 1.7 37.25 1.92c1.73 0 1.78-.08 1.82-1.85c.68-27.49.58-22.59 1-29.55a2.69 2.69 0 0 0-1.63-2.8c-5.6-2.91-8.75-7.55-8.9-13.87c-.35-14.81 17.72-21.67 27.38-11.51c6.84 7.19 5.8 18.91-2.45 24.15a4.35 4.35 0 0 0-2.22 4.34c0 .59-.11-4.31 1 30.05c0 .9.43 1.12 1.24 1.11c.1 0 23-.09 34.47-.37zM68.27 141.7c19.84-4.51 32.68-.56 52.49 1.69c2.76.31 3.74 1.22 3.62 4c-.21 5-1.16 22.33-1.24 23.15a2.65 2.65 0 0 1-1.63 2.34c-4.06 1.7-3.61-4.45-4-7.29c-3.13-22.43-73.87-32.7-74.63 25.4c-.31 23.92 17 53.63 54.08 50.88c27.24-2 19-20.19 24.84-20.47a2.72 2.72 0 0 1 3 3.36c-1.83 10.85-3.42 18.95-3.45 19.15c-1.54 9.17-86.7 22.09-93.35-42.06c-2.71-25.85 10.44-53.37 40.27-60.15m80 87.67h-19.49a2.57 2.57 0 0 1-2.66-1.79c2.38-3.75 5.89.92 5.86-6.14c-.08-25.75.21-38 .23-40.1c0-3.42-.53-4.65-3.32-4.94c-7-.72-3.11-3.37-1.11-3.38c11.84-.1 22.62-.18 30.05.72c8.77 1.07 16.71 12.63 7.93 22.62c-2 2.25-4 4.42-6.14 6.73c.95 1.15 6.9 8.82 17.28 19.68c2.66 2.78 6.15 3.51 9.88 3.13a2.21 2.21 0 0 0 2.23-2.12c.3-3.42.26 4.73.45-40.58c0-5.65-.34-6.58-3.23-6.83c-3.95-.35-4-2.26-.69-3.37l19.09-.09c.32 0 4.49.53 1 3.38c0 .05-.16 0-.24 0c-3.61.26-3.94 1-4 4.62c-.27 43.93.07 40.23.41 42.82c.11.84.27 2.23 5.1 2.14c2.49 0 3.86 3.37 0 3.4c-10.37.08-20.74 0-31.11.07c-10.67 0-13.47-6.2-24.21-20.82c-1.6-2.18-8.31-2.36-8.2-.37c.88 16.47 0 17.78 4 17.67c4.75-.1 4.73 3.57.83 3.55zm275-10.15c-1.21 7.13.17 10.38-5.3 10.34c-61.55-.42-47.82-.22-50.72-.31a18.4 18.4 0 0 1-3.63-.73c-2.53-.6 1.48-1.23-.38-5.6c-1.43-3.37-2.78-6.78-4.11-10.19a1.94 1.94 0 0 0-2-1.44a138 138 0 0 0-14.58.07a2.23 2.23 0 0 0-1.62 1.06c-1.58 3.62-3.07 7.29-4.51 11c-1.27 3.23 7.86 1.32 12.19 2.16c3 .57 4.53 3.72.66 3.73H322.9c-2.92 0-3.09-3.15-.74-3.21a6.3 6.3 0 0 0 5.92-3.47c1.5-3 2.8-6 4.11-9.09c18.18-42.14 17.06-40.17 18.42-41.61a1.83 1.83 0 0 1 3 0c2.93 3.34 18.4 44.71 23.62 51.92c2 2.7 5.74 2 6.36 2c3.61.13 4-1.11 4.13-4.29c.09-1.87.08 1.17.07-41.24c0-4.46-2.36-3.74-5.55-4.27c-.26 0-2.56-.63-.08-3.06c.21-.2-.89-.24 21.7-.15c2.32 0 5.32 2.75-1.21 3.45a2.56 2.56 0 0 0-2.66 2.83c-.07 1.63-.19 38.89.29 41.21a3.06 3.06 0 0 0 3.23 2.43c13.25.43 14.92.44 16-3.41c1.67-5.78 4.13-2.52 3.73-.19zm-104.72 64.37c-4.24 0-4.42-3.39-.61-3.41c35.91-.16 28.11.38 37.19-.65c1.68-.19 2.38.24 2.25 1.89c-.26 3.39-.64 6.78-1 10.16c-.25 2.16-3.2 2.61-3.4-.15c-.38-5.31-2.15-4.45-15.63-5.08c-1.58-.07-1.64 0-1.64 1.52V304c0 1.65 0 1.6 1.62 1.47c3.12-.25 10.31.34 15.69-1.52c.47-.16 3.3-1.79 3.07 1.76c0 .21-.76 10.35-1.18 11.39c-.53 1.29-1.88 1.51-2.58.32c-1.17-2 0-5.08-3.71-5.3c-15.42-.9-12.91-2.55-12.91 6c0 12.25-.76 16.11 3.89 16.24c16.64.48 14.4 0 16.43-5.71c.84-2.37 3.5-1.77 3.18.58c-.44 3.21-.85 6.43-1.23 9.64c0 .36-.16 2.4-4.66 2.39c-37.16-.08-34.54-.19-35.21-.31c-2.72-.51-2.2-3 .22-3.45c1.1-.19 4 .54 4.16-2.56c2.44-56.22-.07-51.34-3.91-51.33zm-.41-109.52c2.46.61 3.13 1.76 2.95 4.65c-.33 5.3-.34 9-.55 9.69c-.66 2.23-3.15 2.12-3.34-.27c-.38-4.81-3.05-7.82-7.57-9.15c-26.28-7.73-32.81 15.46-27.17 30.22c5.88 15.41 22 15.92 28.86 13.78c5.92-1.85 5.88-6.5 6.91-7.58c1.23-1.3 2.25-1.84 3.12 1.1c0 .1.57 11.89-6 12.75c-1.6.21-19.38 3.69-32.68-3.39c-21-11.19-16.74-35.47-6.88-45.33c14-14.06 39.91-7.06 42.32-6.47zM289.8 280.14c3.28 0 3.66 3 .16 3.43c-2.61.32-5-.42-5 5.46c0 2-.19 29.05.4 41.45c.11 2.29 1.15 3.52 3.44 3.65c22 1.21 14.95-1.65 18.79-6.34c1.83-2.24 2.76.84 2.76 1.08c.35 13.62-4 12.39-5.19 12.4l-38.16-.19c-1.93-.23-2.06-3-.42-3.38c2-.48 4.94.4 5.13-2.8c1-15.87.57-44.65.34-47.81c-.27-3.77-2.8-3.27-5.68-3.71c-2.47-.38-2-3.22.34-3.22c1.45-.02 17.97-.03 23.09-.02m-31.63-57.79c.07 4.08 2.86 3.46 6 3.58c2.61.1 2.53 3.41-.07 3.43c-6.48 0-13.7 0-21.61-.06c-3.84 0-3.38-3.35 0-3.37c4.49 0 3.24 1.61 3.41-45.54c0-5.08-3.27-3.54-4.72-4.23c-2.58-1.23-1.36-3.09.41-3.15c1.29 0 20.19-.41 21.17.21s1.87 1.65-.42 2.86c-1 .52-3.86-.28-4.15 2.47c0 .21-.82 1.63-.07 43.8zm-36.91 274.27a2.93 2.93 0 0 0 3.26 0c17-9.79 182-103.57 197.42-112.51c-.14-.43 11.26-.18-181.52-.27c-1.22 0-1.57.37-1.53 1.56c0 .1 1.25 44.51 1.22 50.38a28.33 28.33 0 0 1-1.36 7.71c-.55 1.83.38-.5-13.5 32.23c-.73 1.72-1 2.21-2-.08c-4.19-10.34-8.28-20.72-12.57-31a23.6 23.6 0 0 1-2-10.79c.16-2.46.8-16.12 1.51-48c0-1.95 0-2-2-2h-183c2.58 1.63 178.32 102.57 196 112.76zm-90.9-188.75c0 2.4.36 2.79 2.76 3c11.54 1.17 21 3.74 25.64-7.32c6-14.46 2.66-34.41-12.48-38.84c-2-.59-16-2.76-15.94 1.51c.05 8.04.01 11.61.02 41.65m105.75-15.05c0 2.13 1.07 38.68 1.09 39.13c.34 9.94-25.58 5.77-25.23-2.59c.08-2 1.37-37.42 1.1-39.43c-14.1 7.44-14.42 40.21 6.44 48.8a17.9 17.9 0 0 0 22.39-7.07c4.91-7.76 6.84-29.47-5.43-39a2.53 2.53 0 0 1-.36.12zm-12.28-198c-9.83 0-9.73 14.75-.07 14.87s10.1-14.88.07-14.91zm-80.15 103.83c0 1.8.41 2.4 2.17 2.58c13.62 1.39 12.51-11 12.16-13.36c-1.69-11.22-14.38-10.2-14.35-7.81c.05 4.5-.03 13.68.02 18.59m212.32 6.4l-6.1-15.84c-2.16 5.48-4.16 10.57-6.23 15.84z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CssThree(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m480 32l-64 368l-223.3 80L0 400l19.6-94.8h82l-8 40.6L210 390.2l134.1-44.4l18.8-97.1H29.5l16-82h333.7l10.5-52.7H56.3l16.3-82z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CssThreeAlt(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 32l34.9 395.8L192 480l157.1-52.2L384 32zm313.1 80l-4.8 47.3L193 208.6l-.3.1h111.5l-12.8 146.6l-98.2 28.7l-98.8-29.2l-6.4-73.9h48.9l3.2 38.3l52.6 13.3l54.7-15.4l3.7-61.6l-166.3-.5v-.1l-.2.1l-3.6-46.3L193.1 162l6.5-2.7H76.7L70.9 112z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cuttlefish(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M344 305.5c-17.5 31.6-57.4 54.5-96 54.5c-56.6 0-104-47.4-104-104s47.4-104 104-104c38.6 0 78.5 22.9 96 54.5c13.7-50.9 41.7-93.3 87-117.8C385.7 39.1 320.5 8 248 8C111 8 0 119 0 256s111 248 248 248c72.5 0 137.7-31.1 183-80.7c-45.3-24.5-73.3-66.9-87-117.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DandD(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M82.5 98.9c-.6-17.2 2-33.8 12.7-48.2c.3 7.4 1.2 14.5 4.2 21.6c5.9-27.5 19.7-49.3 42.3-65.5c-1.9 5.9-3.5 11.8-3 17.7c8.7-7.4 18.8-17.8 44.4-22.7c14.7-2.8 29.7-2 42.1 1c38.5 9.3 61 34.3 69.7 72.3c5.3 23.1.7 45-8.3 66.4c-5.2 12.4-12 24.4-20.7 35.1c-2-1.9-3.9-3.8-5.8-5.6c-42.8-40.8-26.8-25.2-37.4-37.4c-1.1-1.2-1-2.2-.1-3.6c8.3-13.5 11.8-28.2 10-44c-1.1-9.8-4.3-18.9-11.3-26.2c-14.5-15.3-39.2-15-53.5.6c-11.4 12.5-14.1 27.4-10.9 43.6c.2 1.3.4 2.7 0 3.9c-3.4 13.7-4.6 27.6-2.5 41.6c.1.5.1 1.1.1 1.6c0 .3-.1.5-.2 1.1c-21.8-11-36-28.3-43.2-52.2c-8.3 17.8-11.1 35.5-6.6 54.1c-15.6-15.2-21.3-34.3-22-55.2m469.6 123.2c-11.6-11.6-25-20.4-40.1-26.6c-12.8-5.2-26-7.9-39.9-7.1c-10 .6-19.6 3.1-29 6.4c-2.5.9-5.1 1.6-7.7 2.2c-4.9 1.2-7.3-3.1-4.7-6.8c3.2-4.6 3.4-4.2 15-12c.6-.4 1.2-.8 2.2-1.5h-2.5c-.6 0-1.2.2-1.9.3c-19.3 3.3-30.7 15.5-48.9 29.6c-10.4 8.1-13.8 3.8-12-.5c1.4-3.5 3.3-6.7 5.1-10c1-1.8 2.3-3.4 3.5-5.1c-.2-.2-.5-.3-.7-.5c-27 18.3-46.7 42.4-57.7 73.3c.3.3.7.6 1 .9c.3-.6.5-1.2.9-1.7c10.4-12.1 22.8-21.8 36.6-29.8c18.2-10.6 37.5-18.3 58.7-20.2c4.3-.4 8.7-.1 13.1-.1c-1.8.7-3.5.9-5.3 1.1c-18.5 2.4-35.5 9-51.5 18.5c-30.2 17.9-54.5 42.2-75.1 70.4c-.3.4-.4.9-.7 1.3c14.5 5.3 24 17.3 36.1 25.6c.2-.1.3-.2.4-.4l1.2-2.7c12.2-26.9 27-52.3 46.7-74.5c16.7-18.8 38-25.3 62.5-20c5.9 1.3 11.4 4.4 17.2 6.8c2.3-1.4 5.1-3.2 8-4.7c8.4-4.3 17.4-7 26.7-9c14.7-3.1 29.5-4.9 44.5-1.3v-.5c-.5-.4-1.2-.8-1.7-1.4M316.7 397.6c-39.4-33-22.8-19.5-42.7-35.6c-.8.9 0-.2-1.9 3c-11.2 19.1-25.5 35.3-44 47.6c-10.3 6.8-21.5 11.8-34.1 11.8c-21.6 0-38.2-9.5-49.4-27.8c-12-19.5-13.3-40.7-8.2-62.6c7.8-33.8 30.1-55.2 38.6-64.3c-18.7-6.2-33 1.7-46.4 13.9c.8-13.9 4.3-26.2 11.8-37.3c-24.3 10.6-45.9 25-64.8 43.9c-.3-5.8 5.4-43.7 5.6-44.7c.3-2.7-.6-5.3-3-7.4c-24.2 24.7-44.5 51.8-56.1 84.6c7.4-5.9 14.9-11.4 23.6-16.2c-8.3 22.3-19.6 52.8-7.8 101.1c4.6 19 11.9 36.8 24.1 52.3c2.9 3.7 6.3 6.9 9.5 10.3c.2-.2.4-.3.6-.5c-1.4-7-2.2-14.1-1.5-21.9c2.2 3.2 3.9 6 5.9 8.6c12.6 16 28.7 27.4 47.2 35.6c25 11.3 51.1 13.3 77.9 8.6c54.9-9.7 90.7-48.6 116-98.8c1-1.8.6-2.9-.9-4.2m172-46.4c-9.5-3.1-22.2-4.2-28.7-2.9c9.9 4 14.1 6.6 18.8 12c12.6 14.4 10.4 34.7-5.4 45.6c-11.7 8.1-24.9 10.5-38.9 9.1c-1.2-.1-2.3-.4-3-.6c2.8-3.7 6-7 8.1-10.8c9.4-16.8 5.4-42.1-8.7-56.1c-2.1-2.1-4.6-3.9-7-5.9c-.3 1.3-.1 2.1.1 2.8c4.2 16.6-8.1 32.4-24.8 31.8c-7.6-.3-13.9-3.8-19.6-8.5c-19.5-16.1-39.1-32.1-58.5-48.3c-5.9-4.9-12.5-8.1-20.1-8.7c-4.6-.4-9.3-.6-13.9-.9c-5.9-.4-8.8-2.8-10.4-8.4c-.9-3.4-1.5-6.8-2.2-10.2c-1.5-8.1-6.2-13-14.3-14.2c-4.4-.7-8.9-1-13.3-1.5c-13-1.4-19.8-7.4-22.6-20.3c-5 11-1.6 22.4 7.3 29.9c4.5 3.8 9.3 7.3 13.8 11.2c4.6 3.8 7.4 8.7 7.9 14.8c.4 4.7.8 9.5 1.8 14.1c2.2 10.6 8.9 18.4 17 25.1c16.5 13.7 33 27.3 49.5 41.1c17.9 15 13.9 32.8 13 56c-.9 22.9 12.2 42.9 33.5 51.2c1 .4 2 .6 3.6 1.1c-15.7-18.2-10.1-44.1.7-52.3c.3 2.2.4 4.3.9 6.4c9.4 44.1 45.4 64.2 85 56.9c16-2.9 30.6-8.9 42.9-19.8c2-1.8 3.7-4.1 5.9-6.5c-19.3 4.6-35.8.1-50.9-10.6c.7-.3 1.3-.3 1.9-.3c21.3 1.8 40.6-3.4 57-17.4c19.5-16.6 26.6-42.9 17.4-66c-8.3-20.1-23.6-32.3-43.8-38.9M99.4 179.3c-5.3-9.2-13.2-15.6-22.1-21.3c13.7-.5 26.6.2 39.6 3.7c-7-12.2-8.5-24.7-5-38.7c5.3 11.9 13.7 20.1 23.6 26.8c19.7 13.2 35.7 19.6 46.7 30.2c3.4 3.3 6.3 7.1 9.6 10.9c-.8-2.1-1.4-4.1-2.2-6c-5-10.6-13-18.6-22.6-25c-1.8-1.2-2.8-2.5-3.4-4.5c-3.3-12.5-3-25.1-.7-37.6c1-5.5 2.8-10.9 4.5-16.3c.8-2.4 2.3-4.6 4-6.6c.6 6.9 0 25.5 19.6 46c10.8 11.3 22.4 21.9 33.9 32.7c9 8.5 18.3 16.7 25.5 26.8c1.1 1.6 2.2 3.3 3.8 4.7c-5-13-14.2-24.1-24.2-33.8c-9.6-9.3-19.4-18.4-29.2-27.4c-3.3-3-4.6-6.7-5.1-10.9c-1.2-10.4 0-20.6 4.3-30.2c.5-1 1.1-2 1.9-3.3c.5 4.2.6 7.9 1.4 11.6c4.8 23.1 20.4 36.3 49.3 63.5c10 9.4 19.3 19.2 25.6 31.6c4.8 9.3 7.3 19 5.7 29.6c-.1.6.5 1.7 1.1 2c6.2 2.6 10 6.9 9.7 14.3c7.7-2.6 12.5-8 16.4-14.5c4.2 20.2-9.1 50.3-27.2 58.7c.4-4.5 5-23.4-16.5-27.7c-6.8-1.3-12.8-1.3-22.9-2.1c4.7-9 10.4-20.6.5-22.4c-24.9-4.6-52.8 1.9-57.8 4.6c8.2.4 16.3 1 23.5 3.3c-2 6.5-4 12.7-5.8 18.9c-1.9 6.5 2.1 14.6 9.3 9.6c1.2-.9 2.3-1.9 3.3-2.7c-3.1 17.9-2.9 15.9-2.8 18.3c.3 10.2 9.5 7.8 15.7 7.3c-2.5 11.8-29.5 27.3-45.4 25.8c7-4.7 12.7-10.3 15.9-17.9c-6.5.8-12.9 1.6-19.2 2.4l-.3-.9c4.7-3.4 8-7.8 10.2-13.1c8.7-21.1-3.6-38-25-39.9c-9.1-.8-17.8.8-25.9 5.5c6.2-15.6 17.2-26.6 32.6-34.5c-15.2-4.3-8.9-2.7-24.6-6.3c14.6-9.3 30.2-13.2 46.5-14.6c-5.2-3.2-48.1-3.6-70.2 20.9c7.9 1.4 15.5 2.8 23.2 4.2c-23.8 7-44 19.7-62.4 35.6c1.1-4.8 2.7-9.5 3.3-14.3c.6-4.5.8-9.2.1-13.6c-1.5-9.4-8.9-15.1-19.7-16.3c-7.9-.9-15.6.1-23.3 1.3c-.9.1-1.7.3-2.9 0c15.8-14.8 36-21.7 53.1-33.5c6-4.5 6.8-8.2 3-14.9m128.4 26.8c3.3 16 12.6 25.5 23.8 24.3c-4.6-11.3-12.1-19.5-23.8-24.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DandDbeyond(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M313.8 241.5c13.8 0 21-10.1 24.8-17.9c-1-1.1-5-4.2-7.4-6.6c-2.4 4.3-8.2 10.7-13.9 10.7c-10.2 0-15.4-14.7-3.2-26.6c-.5-.2-4.3-1.8-8 2.4c0-3 1-5.1 2.1-6.6c-3.5 1.3-9.8 5.6-11.4 7.9c.2-5.8 1.6-7.5.6-9l-.2-.2s-8.5 5.6-9.3 14.7c0 0 1.1-1.6 2.1-1.9c.6-.3 1.3 0 .6 1.9c-.2.6-5.8 15.7 5.1 26c-.6-1.6-1.9-7.6 2.4-1.9c-.3.1 5.8 7.1 15.7 7.1m52.4-21.1c0-4-4.9-4.4-5.6-4.5c2 3.9.9 7.5.2 9c2.5-.4 5.4-1.6 5.4-4.5m10.3 5.2c0-6.4-6.2-11.4-13.5-10.7c8 1.3 5.6 13.8-5 11.4c3.7-2.6 3.2-9.9-1.3-12.5c1.4 4.2-3 8.2-7.4 4.6c-2.4-1.9-8-6.6-10.6-8.6c-2.4-2.1-5.5-1-6.6-1.8c-1.3-1.1-.5-3.8-2.2-5c-1.6-.8-3-.3-4.8-1c-1.6-.6-2.7-1.9-2.6-3.5c-2.5 4.4 3.4 6.3 4.5 8.5c1 1.9-.8 4.8 4 8.5c14.8 11.6 9.1 8 10.4 18.1c.6 4.3 4.2 6.7 6.4 7.4c-2.1-1.9-2.9-6.4 0-9.3c0 13.9 19.2 13.3 23.1 6.4c-2.4 1.1-7-.2-9-1.9c7.7 1 14.2-4.1 14.6-10.6m-39.4-18.4c2 .8 1.6.7 6.4 4.5c10.2-24.5 21.7-15.7 22-15.5c2.2-1.9 9.8-3.8 13.8-2.7c-2.4-2.7-7.5-6.2-13.3-6.2c-4.7 0-7.4 2.2-8 1.3c-.8-1.4 3.2-3.4 3.2-3.4c-5.4.2-9.6 6.7-11.2 5.9c-1.1-.5 1.4-3.7 1.4-3.7c-5.1 2.9-9.3 9.1-10.2 13c4.6-5.8 13.8-9.8 19.7-9c-10.5.5-19.5 9.7-23.8 15.8m242.5 51.9c-20.7 0-40 1.3-50.3 2.1l7.4 8.2v77.2l-7.4 8.2c10.4.8 30.9 2.1 51.6 2.1c42.1 0 59.1-20.7 59.1-48.9c0-29.3-23.2-48.9-60.4-48.9m-15.1 75.6v-53.3c30.1-3.3 46.8 3.8 46.8 26.3c0 25.6-21.4 30.2-46.8 27M301.6 181c-1-3.4-.2-6.9 1.1-9.4c1 3 2.6 6.4 7.5 9c-.5-2.4-.2-5.6.5-8c-1.4-5.4 2.1-9.9 6.4-9.9c6.9 0 8.5 8.8 4.7 14.4c2.1 3.2 5.5 5.6 7.7 7.8c3.2-3.7 5.5-9.5 5.5-13.8c0-8.2-5.5-15.9-16.7-16.5c-20-.9-20.2 16.6-20 18.9c.5 5.2 3.4 7.8 3.3 7.5m-.4 6c-.5 1.8-7 3.7-10.2 6.9c4.8-1 7-.2 7.8 1.8c.5 1.4-.2 3.4-.5 5.6c1.6-1.8 7-5.5 11-6.2c-1-.3-3.4-.8-4.3-.8c2.9-3.4 9.3-4.5 12.8-3.7c-2.2-.2-6.7 1.1-8.5 2.6c1.6.3 3 .6 4.3 1.1c-2.1.8-4.8 3.4-5.8 6.1c7-5 13.1 5.2 7 8.2c.8.2 2.7 0 3.5-.5c-.3 1.1-1.9 3-3 3.4c2.9 0 7-1.9 8.2-4.6c0 0-1.8.6-2.6-.2s.3-4.3.3-4.3c-2.3 2.9-3.4-1.3-1.3-4.2c-1-.3-3.5-.6-4.6-.5c3.2-1.1 10.4-1.8 11.2-.3c.6 1.1-1 3.4-1 3.4c4-.5 8.3 1.1 6.7 5.1c2.9-1.4 5.5-5.9 4.8-10.4c-.3 1-1.6 2.4-2.9 2.7c.2-1.4-1-2.2-1.9-2.6c1.7-9.6-14.6-14.2-14.1-23.9c-1 1.3-1.8 5-.8 7.1c2.7 3.2 8.7 6.7 10.1 12.2c-2.6-6.4-15.1-11.4-14.6-20.2c-1.6 1.6-2.6 7.8-1.3 11c2.4 1.4 4.5 3.8 4.8 6.1c-2.2-5.1-11.4-6.1-13.9-12.2c-.6 2.2-.3 5 1 6.7c0 0-2.2-.8-7-.6c1.7.6 5.1 3.5 4.8 5.2m25.9 7.4c-2.7 0-3.5-2.1-4.2-4.3c3.3 1.3 4.2 4.3 4.2 4.3m38.9 3.7l-1-.6c-1.1-1-2.9-1.4-4.7-1.4c-2.9 0-5.8 1.3-7.5 3.4c-.8.8-1.4 1.8-2.1 2.6v15.7c3.5 2.6 7.1-2.9 3-7.2c1.5.3 4.6 2.7 5.1 3.2c0 0 2.6-.5 5-.5c2.1 0 3.9.3 5.6 1.1V196c-1.1.5-2.2 1-2.7 1.4zM79.9 305.9c17.2-4.6 16.2-18 16.2-19.9c0-20.6-24.1-25-37-25H3l8.3 8.6v29.5H0l11.4 14.6V346L3 354.6c61.7 0 73.8 1.5 86.4-5.9c6.7-4 9.9-9.8 9.9-17.6c0-5.1 2.6-18.8-19.4-25.2m-41.3-27.5c20 0 29.6-.8 29.6 9.1v3c0 12.1-19 8.8-29.6 8.8zm0 59.2V315c12.2 0 32.7-2.3 32.7 8.8v4.5h.2c0 11.2-12.5 9.3-32.9 9.3m101.2-19.3l23.1.2v-.2l14.1-21.2h-37.2v-14.9h52.4l-14.1-21v-.2l-73.5.2l7.4 8.2v77.1l-7.4 8.2h81.2l14.1-21.2l-60.1.2zm214.7-60.1c-73.9 0-77.5 99.3-.3 99.3c77.9 0 74.1-99.3.3-99.3m-.3 77.5c-37.4 0-36.9-55.3.2-55.3c36.8.1 38.8 55.3-.2 55.3m-91.3-8.3l44.1-66.2h-41.7l6.1 7.2l-20.5 37.2h-.3l-21-37.2l6.4-7.2h-44.9l44.1 65.8l.2 19.4l-7.7 8.2h42.6l-7.2-8.2zm-28.4-151.3c1.6 1.3 2.9 2.4 2.9 6.6v38.8c0 4.2-.8 5.3-2.7 6.4c-.1.1-7.5 4.5-7.9 4.6h35.1c10 0 17.4-1.5 26-8.6c-.6-5 .2-9.5.8-12c0-.2-1.8 1.4-2.7 3.5c0-5.7 1.6-15.4 9.6-20.5c-.1 0-3.7-.8-9 1.1c2-3.1 10-7.9 10.4-7.9c-8.2-26-38-22.9-32.2-22.9c-30.9 0-32.6.3-39.9-4c.1.8.5 8.2 9.6 14.9m21.5 5.5c4.6 0 23.1-3.3 23.1 17.3c0 20.7-18.4 17.3-23.1 17.3zm228.9 79.6l7 8.3V312h-.3c-5.4-14.4-42.3-41.5-45.2-50.9h-31.6l7.4 8.5v76.9l-7.2 8.3h39l-7.4-8.2v-47.4h.3c3.7 10.6 44.5 42.9 48.5 55.6h21.3v-85.2l7.4-8.3zm-106.7-96.1c-32.2 0-32.8.2-39.9-4c.1.7.5 8.3 9.6 14.9c3.1 2 2.9 4.3 2.9 9.5c1.8-1.1 3.8-2.2 6.1-3c-1.1 1.1-2.7 2.7-3.5 4.5c1-1.1 7.5-5.1 14.6-3.5c-1.6.3-4 1.1-6.1 2.9c.1 0 2.1-1.1 7.5-.3v-4.3c4.7 0 23.1-3.4 23.1 17.3c0 20.5-18.5 17.3-19.7 17.3c5.7 4.4 5.8 12 2.2 16.3h.3c33.4 0 36.7-27.3 36.7-34c0-3.8-1.1-32-33.8-33.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dailymotion(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M298.93 267a48.4 48.4 0 0 0-24.36-6.21q-19.83 0-33.44 13.27t-13.61 33.42q0 21.16 13.28 34.6t33.43 13.44q20.5 0 34.11-13.78T322 307.47a47.13 47.13 0 0 0-6.1-23.47a44.13 44.13 0 0 0-16.97-17M0 32v448h448V32Zm374.71 373.26h-53.1v-23.89h-.67q-15.79 26.2-55.78 26.2q-27.56 0-48.89-13.1a88.29 88.29 0 0 1-32.94-35.77q-11.6-22.68-11.59-50.89q0-27.56 11.76-50.22a89.9 89.9 0 0 1 32.93-35.78q21.18-13.09 47.72-13.1a80.87 80.87 0 0 1 29.74 5.21q13.28 5.21 25 17V153l55.79-12.09Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dashcube(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M326.6 104H110.4c-51.1 0-91.2 43.3-91.2 93.5V427c0 50.5 40.1 85 91.2 85h227.2c51.1 0 91.2-34.5 91.2-85V0zM153.9 416.5c-17.7 0-32.4-15.1-32.4-32.8V240.8c0-17.7 14.7-32.5 32.4-32.5h140.7c17.7 0 32 14.8 32 32.5v123.5l51.1 52.3H153.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Deezer(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M451.46 244.71H576V172H451.46Zm0-173.89v72.67H576V70.82Zm0 275.06H576V273.2H451.46ZM0 447.09h124.54v-72.67H0Zm150.47 0H275v-72.67H150.47Zm150.52 0h124.54v-72.67H301Zm150.47 0H576v-72.67H451.46ZM301 345.88h124.53V273.2H301Zm-150.52 0H275V273.2H150.47Zm0-101.17H275V172H150.47Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Delicious(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M446.5 68c-.4-1.5-.9-3-1.4-4.5c-.9-2.5-2-4.8-3.3-7.1c-1.4-2.4-3-4.8-4.7-6.9c-2.1-2.5-4.4-4.8-6.9-6.8c-1.1-.9-2.2-1.7-3.3-2.5c-1.3-.9-2.6-1.7-4-2.4c-1.8-1-3.6-1.8-5.5-2.5c-1.7-.7-3.5-1.3-5.4-1.7c-3.8-1-7.9-1.5-12-1.5H48C21.5 32 0 53.5 0 80v352c0 4.1.5 8.2 1.5 12c2 7.7 5.8 14.6 11 20.3c1 1.1 2.1 2.2 3.3 3.3c5.7 5.2 12.6 9 20.3 11c3.8 1 7.9 1.5 12 1.5h352c26.5 0 48-21.5 48-48V80c-.1-4.1-.6-8.2-1.6-12M416 432c0 8.8-7.2 16-16 16H224V256H32V80c0-8.8 7.2-16 16-16h176v192h192z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Deploydog(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M382.2 136h51.7v239.6h-51.7v-20.7c-19.8 24.8-52.8 24.1-73.8 14.7c-26.2-11.7-44.3-38.1-44.3-71.8c0-29.8 14.8-57.9 43.3-70.8c20.2-9.1 52.7-10.6 74.8 12.9zm-64.7 161.8c0 18.2 13.6 33.5 33.2 33.5c19.8 0 33.2-16.4 33.2-32.9c0-17.1-13.7-33.2-33.2-33.2c-19.6 0-33.2 16.4-33.2 32.6M188.5 136h51.7v239.6h-51.7v-20.7c-19.8 24.8-52.8 24.1-73.8 14.7c-26.2-11.7-44.3-38.1-44.3-71.8c0-29.8 14.8-57.9 43.3-70.8c20.2-9.1 52.7-10.6 74.8 12.9zm-64.7 161.8c0 18.2 13.6 33.5 33.2 33.5c19.8 0 33.2-16.4 33.2-32.9c0-17.1-13.7-33.2-33.2-33.2c-19.7 0-33.2 16.4-33.2 32.6M448 96c17.5 0 32 14.4 32 32v256c0 17.5-14.4 32-32 32H64c-17.5 0-32-14.4-32-32V128c0-17.5 14.4-32 32-32zm0-32H64C28.8 64 0 92.8 0 128v256c0 35.2 28.8 64 64 64h384c35.2 0 64-28.8 64-64V128c0-35.2-28.8-64-64-64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Deskpro(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m205.9 512l31.1-38.4c12.3-.2 25.6-1.4 36.5-6.6c38.9-18.6 38.4-61.9 38.3-63.8c-.1-5-.8-4.4-28.9-37.4H362c-.2 50.1-7.3 68.5-10.2 75.7c-9.4 23.7-43.9 62.8-95.2 69.4c-8.7 1.1-32.8 1.2-50.7 1.1m200.4-167.7c38.6 0 58.5-13.6 73.7-30.9l-175.5-.3l-17.4 31.3zm-43.6-223.9v168.3h-73.5l-32.7 55.5H250c-52.3 0-58.1-56.5-58.3-58.9c-1.2-13.2-21.3-11.6-20.1 1.8c1.4 15.8 8.8 40 26.4 57.1h-91c-25.5 0-110.8-26.8-107-114V16.9C0 .9 9.7.3 15 .1h82c.2 0 .3.1.5.1c4.3-.4 50.1-2.1 50.1 43.7c0 13.3 20.2 13.4 20.2 0c0-18.2-5.5-32.8-15.8-43.7h84.2c108.7-.4 126.5 79.4 126.5 120.2m-132.5 56l64 29.3c13.3-45.5-42.2-71.7-64-29.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dev(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M120.12 208.29c-3.88-2.9-7.77-4.35-11.65-4.35H91.03v104.47h17.45c3.88 0 7.77-1.45 11.65-4.35c3.88-2.9 5.82-7.25 5.82-13.06v-69.65c-.01-5.8-1.96-10.16-5.83-13.06M404.1 32H43.9C19.7 32 .06 51.59 0 75.8v360.4C.06 460.41 19.7 480 43.9 480h360.2c24.21 0 43.84-19.59 43.9-43.8V75.8c-.06-24.21-19.7-43.8-43.9-43.8M154.2 291.19c0 18.81-11.61 47.31-48.36 47.25h-46.4V172.98h47.38c35.44 0 47.36 28.46 47.37 47.28zm100.68-88.66H201.6v38.42h32.57v29.57H201.6v38.41h53.29v29.57h-62.18c-11.16.29-20.44-8.53-20.72-19.69V193.7c-.27-11.15 8.56-20.41 19.71-20.69h63.19zm103.64 115.29c-13.2 30.75-36.85 24.63-47.44 0l-38.53-144.8h32.57l29.71 113.72l29.57-113.72h32.58z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Deviantart(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m320 93.2l-98.2 179.1l7.4 9.5H320v127.7H159.1l-13.5 9.2l-43.7 84c-.3 0-8.6 8.6-9.2 9.2H0v-93.2l93.2-179.4l-7.4-9.2H0V102.5h156l13.5-9.2l43.7-84c.3 0 8.6-8.6 9.2-9.2H320z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dhl(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M238 301.2h58.7L319 271h-58.7zM0 282.9v6.4h81.8l4.7-6.4zM172.9 271c-8.7 0-6-3.6-4.6-5.5c2.8-3.8 7.6-10.4 10.4-14.1c2.8-3.7 2.8-5.9-2.8-5.9h-51l-41.1 55.8h100.1c33.1 0 51.5-22.5 57.2-30.3zm317.5-6.9l39.3-53.4h-62.2l-39.3 53.4zM95.3 271H0v6.4h90.6zm111-26.6c-2.8 3.8-7.5 10.4-10.3 14.2c-1.4 2-4.1 5.5 4.6 5.5h45.6s7.3-10 13.5-18.4c8.4-11.4.7-35-29.2-35H112.6l-20.4 27.8h111.4c5.6 0 5.5 2.2 2.7 5.9M0 301.2h73.1l4.7-6.4H0zm323 0h58.7L404 271h-58.7c-.1 0-22.3 30.2-22.3 30.2m222 .1h95v-6.4h-90.3zm22.3-30.3l-4.7 6.4H640V271zm-13.5 18.3H640v-6.4h-81.5zm-164.2-78.6l-22.5 30.6h-26.2l22.5-30.6h-58.7l-39.3 53.4H409l39.3-53.4zm33.5 60.3s-4.3 5.9-6.4 8.7c-7.4 10-.9 21.6 23.2 21.6h94.3l22.3-30.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Diaspora(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M251.64 354.55c-1.4 0-88 119.9-88.7 119.9S76.34 414 76 413.25s86.6-125.7 86.6-127.4c0-2.2-129.6-44-137.6-47.1c-1.3-.5 31.4-101.8 31.7-102.1c.6-.7 144.4 47 145.5 47c.4 0 .9-.6 1-1.3c.4-2 1-148.6 1.7-149.6c.8-1.2 104.5-.7 105.1-.3c1.5 1 3.5 156.1 6.1 156.1c1.4 0 138.7-47 139.3-46.3c.8.9 31.9 102.2 31.5 102.6c-.9.9-140.2 47.1-140.6 48.8c-.3 1.4 82.8 122.1 82.5 122.9s-85.5 63.5-86.3 63.5c-1-.2-89-125.5-90.9-125.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Digg(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M81.7 172.3H0v174.4h132.7V96h-51zm0 133.4H50.9v-92.3h30.8zm297.2-133.4v174.4h81.8v28.5h-81.8V416H512V172.3zm81.8 133.4h-30.8v-92.3h30.8zm-235.6 41h82.1v28.5h-82.1V416h133.3V172.3H225.1zm51.2-133.3h30.8v92.3h-30.8zM153.3 96h51.3v51h-51.3zm0 76.3h51.3v174.4h-51.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DigitalOcean(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M87 481.8h73.7v-73.6H87zM25.4 346.6v61.6H87v-61.6zm466.2-169.7c-23-74.2-82.4-133.3-156.6-156.6C164.9-32.8 8 93.7 8 255.9h95.8c0-101.8 101-180.5 208.1-141.7c39.7 14.3 71.5 46.1 85.8 85.7c39.1 107-39.7 207.8-141.4 208v.3h-.3V504c162.6 0 288.8-156.8 235.6-327.1m-235.3 231v-95.3h-95.6v95.6H256v-.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Discord(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M524.531 69.836a1.5 1.5 0 0 0-.764-.7A485.065 485.065 0 0 0 404.081 32.03a1.816 1.816 0 0 0-1.923.91a337.461 337.461 0 0 0-14.9 30.6a447.848 447.848 0 0 0-134.426 0a309.541 309.541 0 0 0-15.135-30.6a1.89 1.89 0 0 0-1.924-.91a483.689 483.689 0 0 0-119.688 37.107a1.712 1.712 0 0 0-.788.676C39.068 183.651 18.186 294.69 28.43 404.354a2.016 2.016 0 0 0 .765 1.375a487.666 487.666 0 0 0 146.825 74.189a1.9 1.9 0 0 0 2.063-.676A348.2 348.2 0 0 0 208.12 430.4a1.86 1.86 0 0 0-1.019-2.588a321.173 321.173 0 0 1-45.868-21.853a1.885 1.885 0 0 1-.185-3.126a251.047 251.047 0 0 0 9.109-7.137a1.819 1.819 0 0 1 1.9-.256c96.229 43.917 200.41 43.917 295.5 0a1.812 1.812 0 0 1 1.924.233a234.533 234.533 0 0 0 9.132 7.16a1.884 1.884 0 0 1-.162 3.126a301.407 301.407 0 0 1-45.89 21.83a1.875 1.875 0 0 0-1 2.611a391.055 391.055 0 0 0 30.014 48.815a1.864 1.864 0 0 0 2.063.7A486.048 486.048 0 0 0 610.7 405.729a1.882 1.882 0 0 0 .765-1.352c12.264-126.783-20.532-236.912-86.934-334.541M222.491 337.58c-28.972 0-52.844-26.587-52.844-59.239s23.409-59.241 52.844-59.241c29.665 0 53.306 26.82 52.843 59.239c0 32.654-23.41 59.241-52.843 59.241m195.38 0c-28.971 0-52.843-26.587-52.843-59.239s23.409-59.241 52.843-59.241c29.667 0 53.307 26.82 52.844 59.239c0 32.654-23.177 59.241-52.844 59.241"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Discourse(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M225.9 32C103.3 32 0 130.5 0 252.1C0 256 .1 480 .1 480l225.8-.2c122.7 0 222.1-102.3 222.1-223.9C448 134.3 348.6 32 225.9 32M224 384c-19.4 0-37.9-4.3-54.4-12.1L88.5 392l22.9-75c-9.8-18.1-15.4-38.9-15.4-61c0-70.7 57.3-128 128-128s128 57.3 128 128s-57.3 128-128 128"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dochub(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M397.9 160H256V19.6zM304 192v130c0 66.8-36.5 100.1-113.3 100.1H96V84.8h94.7c12 0 23.1.8 33.1 2.5v-84C212.9 1.1 201.4 0 189.2 0H0v512h189.2C329.7 512 400 447.4 400 318.1V192z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Docker(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M349.9 236.3h-66.1v-59.4h66.1zm0-204.3h-66.1v60.7h66.1zm78.2 144.8H362v59.4h66.1zm-156.3-72.1h-66.1v60.1h66.1zm78.1 0h-66.1v60.1h66.1zm276.8 100c-14.4-9.7-47.6-13.2-73.1-8.4c-3.3-24-16.7-44.9-41.1-63.7l-14-9.3l-9.3 14c-18.4 27.8-23.4 73.6-3.7 103.8c-8.7 4.7-25.8 11.1-48.4 10.7H2.4c-8.7 50.8 5.8 116.8 44 162.1c37.1 43.9 92.7 66.2 165.4 66.2c157.4 0 273.9-72.5 328.4-204.2c21.4.4 67.6.1 91.3-45.2c1.5-2.5 6.6-13.2 8.5-17.1zm-511.1-27.9h-66v59.4h66.1v-59.4zm78.1 0h-66.1v59.4h66.1zm78.1 0h-66.1v59.4h66.1zm-78.1-72.1h-66.1v60.1h66.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DraftTwoDigital(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m480 398.1l-144-82.2v64.7h-91.3c30.8-35 81.8-95.9 111.8-149.3c35.2-62.6 16.1-123.4-12.8-153.3c-4.4-4.6-62.2-62.9-166-41.2c-59.1 12.4-89.4 43.4-104.3 67.3c-13.1 20.9-17 39.8-18.2 47.7c-5.5 33 19.4 67.1 56.7 67.1c31.7 0 57.3-25.7 57.3-57.4c0-27.1-19.7-52.1-48-56.8c1.8-7.3 17.7-21.1 26.3-24.7c41.1-17.3 78 5.2 83.3 33.5c8.3 44.3-37.1 90.4-69.7 127.6C84.5 328.1 18.3 396.8 0 415.9l336-.1V480zM369.9 371l47.1 27.2l-47.1 27.2zM134.2 161.4c0 12.4-10 22.4-22.4 22.4s-22.4-10-22.4-22.4s10-22.4 22.4-22.4s22.4 10.1 22.4 22.4M82.5 380.5c25.6-27.4 97.7-104.7 150.8-169.9c35.1-43.1 40.3-82.4 28.4-112.7c-7.4-18.8-17.5-30.2-24.3-35.7c45.3 2.1 68 23.4 82.2 38.3c0 0 42.4 48.2 5.8 113.3c-37 65.9-110.9 147.5-128.5 166.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dribbble(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 8C119.252 8 8 119.252 8 256s111.252 248 248 248s248-111.252 248-248S392.748 8 256 8m163.97 114.366c29.503 36.046 47.369 81.957 47.835 131.955c-6.984-1.477-77.018-15.682-147.502-6.818c-5.752-14.041-11.181-26.393-18.617-41.614c78.321-31.977 113.818-77.482 118.284-83.523M396.421 97.87c-3.81 5.427-35.697 48.286-111.021 76.519c-34.712-63.776-73.185-116.168-79.04-124.008c67.176-16.193 137.966 1.27 190.061 47.489m-230.48-33.25c5.585 7.659 43.438 60.116 78.537 122.509c-99.087 26.313-186.36 25.934-195.834 25.809C62.38 147.205 106.678 92.573 165.941 64.62M44.17 256.323c0-2.166.043-4.322.108-6.473c9.268.19 111.92 1.513 217.706-30.146c6.064 11.868 11.857 23.915 17.174 35.949c-76.599 21.575-146.194 83.527-180.531 142.306C64.794 360.405 44.17 310.73 44.17 256.323m81.807 167.113c22.127-45.233 82.178-103.622 167.579-132.756c29.74 77.283 42.039 142.053 45.189 160.638c-68.112 29.013-150.015 21.053-212.768-27.882m248.38 8.489c-2.171-12.886-13.446-74.897-41.152-151.033c66.38-10.626 124.7 6.768 131.947 9.055c-9.442 58.941-43.273 109.844-90.795 141.978"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DribbbleSquare(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M90.2 228.2c8.9-42.4 37.4-77.7 75.7-95.7c3.6 4.9 28 38.8 50.7 79c-64 17-120.3 16.8-126.4 16.7M314.6 154c-33.6-29.8-79.3-41.1-122.6-30.6c3.8 5.1 28.6 38.9 51 80c48.6-18.3 69.1-45.9 71.6-49.4M140.1 364c40.5 31.6 93.3 36.7 137.3 18c-2-12-10-53.8-29.2-103.6c-55.1 18.8-93.8 56.4-108.1 85.6m98.8-108.2c-3.4-7.8-7.2-15.5-11.1-23.2C159.6 253 93.4 252.2 87.4 252c0 1.4-.1 2.8-.1 4.2c0 35.1 13.3 67.1 35.1 91.4c22.2-37.9 67.1-77.9 116.5-91.8m34.9 16.3c17.9 49.1 25.1 89.1 26.5 97.4c30.7-20.7 52.5-53.6 58.6-91.6c-4.6-1.5-42.3-12.7-85.1-5.8m-20.3-48.4c4.8 9.8 8.3 17.8 12 26.8c45.5-5.7 90.7 3.4 95.2 4.4c-.3-32.3-11.8-61.9-30.9-85.1c-2.9 3.9-25.8 33.2-76.3 53.9M448 80v352c0 26.5-21.5 48-48 48H48c-26.5 0-48-21.5-48-48V80c0-26.5 21.5-48 48-48h352c26.5 0 48 21.5 48 48m-64 176c0-88.2-71.8-160-160-160S64 167.8 64 256s71.8 160 160 160s160-71.8 160-160"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dropbox(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m264.4 116.3l-132 84.3l132 84.3l-132 84.3L0 284.1l132.3-84.3L0 116.3L132.3 32zM131.6 395.7l132-84.3l132 84.3l-132 84.3zm132.8-111.6l132-84.3l-132-83.6L395.7 32L528 116.3l-132.3 84.3L528 284.8l-132.3 84.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drupal(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M319.5 114.7c-22.2-14-43.5-19.5-64.7-33.5c-13-8.8-31.3-30-46.5-48.3c-2.7 29.3-11.5 41.2-22 49.5c-21.3 17-34.8 22.2-53.5 32.3C117 123 32 181.5 32 290.5C32 399.7 123.8 480 225.8 480C327.5 480 416 406 416 294c0-112.3-83-171-96.5-179.3m2.5 325.6c-20.1 20.1-90.1 28.7-116.7 4.2c-4.8-4.8.3-12 6.5-12c0 0 17 13.3 51.5 13.3c27 0 46-7.7 54.5-14c6.1-4.6 8.4 4.3 4.2 8.5m-54.5-52.6c8.7-3.6 29-3.8 36.8 1.3c4.1 2.8 16.1 18.8 6.2 23.7c-8.4 4.2-1.2-15.7-26.5-15.7c-14.7 0-19.5 5.2-26.7 11c-7 6-9.8 8-12.2 4.7c-6-8.2 15.9-22.3 22.4-25M360 405c-15.2-1-45.5-48.8-65-49.5c-30.9-.9-104.1 80.7-161.3 42c-38.8-26.6-14.6-104.8 51.8-105.2c49.5-.5 83.8 49 108.5 48.5c21.3-.3 61.8-41.8 81.8-41.8c48.7 0 23.3 109.3-15.8 106"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dyalog(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 32v119.2h64V96h107.2C284.6 96 352 176.2 352 255.9C352 332 293.4 416 171.2 416H0v64h171.2C331.9 480 416 367.3 416 255.9c0-58.7-22.1-113.4-62.3-154.3C308.9 56 245.7 32 171.2 32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Earlybirds(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M313.2 47.5c1.2-13 21.3-14 36.6-8.7c.9.3 26.2 9.7 19 15.2c-27.9-7.4-56.4 18.2-55.6-6.5m-201 6.9c30.7-8.1 62 20 61.1-7.1c-1.3-14.2-23.4-15.3-40.2-9.6c-1 .3-28.7 10.5-20.9 16.7M319.4 160c-8.8 0-16 7.2-16 16s7.2 16 16 16s16-7.2 16-16s-7.2-16-16-16m-159.7 0c-8.8 0-16 7.2-16 16s7.2 16 16 16s16-7.2 16-16s-7.2-16-16-16m318.5 163.2c-9.9 24-40.7 11-63.9-1.2c-13.5 69.1-58.1 111.4-126.3 124.2c.3.9-2-.1 24 1c33.6 1.4 63.8-3.1 97.4-8c-19.8-13.8-11.4-37.1-9.8-38.1c1.4-.9 14.7 1.7 21.6 11.5c8.6-12.5 28.4-14.8 30.2-13.6c1.6 1.1 6.6 20.9-6.9 34.6c4.7-.9 8.2-1.6 9.8-2.1c2.6-.8 17.7 11.3 3.1 13.3c-14.3 2.3-22.6 5.1-47.1 10.8c-45.9 10.7-85.9 11.8-117.7 12.8l1 11.6c3.8 18.1-23.4 24.3-27.6 6.2c.8 17.9-27.1 21.8-28.4-1l-.5 5.3c-.7 18.4-28.4 17.9-28.3-.6c-7.5 13.5-28.1 6.8-26.4-8.5l1.2-12.4c-36.7.9-59.7 3.1-61.8 3.1c-20.9 0-20.9-31.6 0-31.6c2.4 0 27.7 1.3 63.2 2.8c-61.1-15.5-103.7-55-114.9-118.2c-25 12.8-57.5 26.8-68.2.8c-10.5-25.4 21.5-42.6 66.8-73.4c.7-6.6 1.6-13.3 2.7-19.8c-14.4-19.6-11.6-36.3-16.1-60.4c-16.8 2.4-23.2-9.1-23.6-23.1c.3-7.3 2.1-14.9 2.4-15.4c1.1-1.8 10.1-2 12.7-2.6c6-31.7 50.6-33.2 90.9-34.5c19.7-21.8 45.2-41.5 80.9-48.3C203.3 29 215.2 8.5 216.2 8c1.7-.8 21.2 4.3 26.3 23.2c5.2-8.8 18.3-11.4 19.6-10.7c1.1.6 6.4 15-4.9 25.9c40.3 3.5 72.2 24.7 96 50.7c36.1 1.5 71.8 5.9 77.1 34c2.7.6 11.6.8 12.7 2.6c.3.5 2.1 8.1 2.4 15.4c-.5 13.9-6.8 25.4-23.6 23.1c-3.2 17.3-2.7 32.9-8.7 47.7c2.4 11.7 4 23.8 4.8 36.4c37 25.4 70.3 42.5 60.3 66.9M207.4 159.9c.9-44-37.9-42.2-78.6-40.3c-21.7 1-38.9 1.9-45.5 13.9c-11.4 20.9 5.9 92.9 23.2 101.2c9.8 4.7 73.4 7.9 86.3-7.1c8.2-9.4 15-49.4 14.6-67.7m52 58.3c-4.3-12.4-6-30.1-15.3-32.7c-2-.5-9-.5-11 0c-10 2.8-10.8 22.1-17 37.2c15.4 0 19.3 9.7 23.7 9.7c4.3 0 6.3-11.3 19.6-14.2m135.7-84.7c-6.6-12.1-24.8-12.9-46.5-13.9c-40.2-1.9-78.2-3.8-77.3 40.3c-.5 18.3 5 58.3 13.2 67.8c13 14.9 76.6 11.8 86.3 7.1c15.8-7.6 36.5-78.9 24.3-101.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ebay(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m606 189.5l-54.8 109.9l-54.9-109.9h-37.5l10.9 20.6c-11.5-19-35.9-26-63.3-26c-31.8 0-67.9 8.7-71.5 43.1h33.7c1.4-13.8 15.7-21.8 35-21.8c26 0 41 9.6 41 33v3.4c-12.7 0-28 .1-41.7.4c-42.4.9-69.6 10-76.7 34.4c1-5.2 1.5-10.6 1.5-16.2c0-52.1-39.7-76.2-75.4-76.2c-21.3 0-43 5.5-58.7 24.2v-80.6h-32.1v169.5c0 10.3-.6 22.9-1.1 33.1h31.5c.7-6.3 1.1-12.9 1.1-19.5c13.6 16.6 35.4 24.9 58.7 24.9c36.9 0 64.9-21.9 73.3-54.2c-.5 2.8-.7 5.8-.7 9c0 24.1 21.1 45 60.6 45c26.6 0 45.8-5.7 61.9-25.5c0 6.6.3 13.3 1.1 20.2h29.8c-.7-8.2-1-17.5-1-26.8v-65.6c0-9.3-1.7-17.2-4.8-23.8l61.5 116.1l-28.5 54.1h35.9L640 189.5zM243.7 313.8c-29.6 0-50.2-21.5-50.2-53.8c0-32.4 20.6-53.8 50.2-53.8c29.8 0 50.2 21.4 50.2 53.8c0 32.3-20.4 53.8-50.2 53.8m200.9-47.3c0 30-17.9 48.4-51.6 48.4c-25.1 0-35-13.4-35-25.8c0-19.1 18.1-24.4 47.2-25.3c13.1-.5 27.6-.6 39.4-.6zm-411.9 1.6h128.8v-8.5c0-51.7-33.1-75.4-78.4-75.4c-56.8 0-83 30.8-83 77.6c0 42.5 25.3 74 82.5 74c31.4 0 68-11.7 74.4-46.1h-33.1c-12 35.8-87.7 36.7-91.2-21.6m95-21.4H33.3c6.9-56.6 92.1-54.7 94.4 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Edge(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M481.92 134.48C440.87 54.18 352.26 8 255.91 8C137.05 8 37.51 91.68 13.47 203.66c26-46.49 86.22-79.14 149.46-79.14c79.27 0 121.09 48.93 122.25 50.18c22 23.8 33 50.39 33 83.1c0 10.4-5.31 25.82-15.11 38.57c-1.57 2-6.39 4.84-6.39 11c0 5.06 3.29 9.92 9.14 14c27.86 19.37 80.37 16.81 80.51 16.81A115.39 115.39 0 0 0 444.94 322a118.92 118.92 0 0 0 58.95-102.44c.5-43.43-15.5-72.3-21.97-85.08M212.77 475.67a154.88 154.88 0 0 1-46.64-45c-32.94-47.42-34.24-95.6-20.1-136A155.5 155.5 0 0 1 203 215.75c59-45.2 94.84-5.65 99.06-1a80 80 0 0 0-4.89-10.14c-9.24-15.93-24-36.41-56.56-53.51c-33.72-17.69-70.59-18.59-77.64-18.59c-38.71 0-77.9 13-107.53 35.69C35.68 183.3 12.77 208.72 8.6 243c-1.08 12.31-2.75 62.8 23 118.27a248 248 0 0 0 248.3 141.61c-38.12-6.62-65.85-26.64-67.13-27.21m250.72-98.33a7.76 7.76 0 0 0-7.92-.23a181.66 181.66 0 0 1-20.41 9.12a197.54 197.54 0 0 1-69.55 12.52c-91.67 0-171.52-63.06-171.52-144a61.12 61.12 0 0 1 6.52-26.75a168.72 168.72 0 0 0-38.76 50c-14.92 29.37-33 88.13 13.33 151.66c6.51 8.91 23 30 56 47.67c23.57 12.65 49 19.61 71.7 19.61c35.14 0 115.43-33.44 163-108.87a7.75 7.75 0 0 0-2.39-10.73"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EdgeLegacy(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m25.71 228.16l.35-.48c0 .16 0 .32-.07.48Zm460.58 15.51c0-44-7.76-84.46-28.81-122.4C416.5 47.88 343.91 8 258.89 8C119 7.72 40.62 113.21 26.06 227.68c42.42-61.31 117.07-121.38 220.37-125c0 0 109.67 0 99.42 105H170c6.37-37.39 18.55-59 34.34-78.93c-75.05 34.9-121.85 96.1-120.75 188.32c.83 71.45 50.13 144.84 120.75 172c83.35 31.84 192.77 7.2 240.13-21.33V363.31c-80.87 56.49-270.87 60.92-272.26-67.57h314.08z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Elementor(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M425.6 32H22.4C10 32 0 42 0 54.4v403.2C0 470 10 480 22.4 480h403.2c12.4 0 22.4-10 22.4-22.4V54.4C448 42 438 32 425.6 32M164.3 355.5h-39.8v-199h39.8zm159.3 0H204.1v-39.8h119.5zm0-79.6H204.1v-39.8h119.5zm0-79.7H204.1v-39.8h119.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ello(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M248 8C111.03 8 0 119.03 0 256s111.03 248 248 248s248-111.03 248-248S384.97 8 248 8m143.84 285.2C375.31 358.51 315.79 404.8 248 404.8s-127.31-46.29-143.84-111.6c-1.65-7.44 2.48-15.71 9.92-17.36c7.44-1.65 15.71 2.48 17.36 9.92c14.05 52.91 62 90.11 116.56 90.11s102.51-37.2 116.56-90.11c1.65-7.44 9.92-12.4 17.36-9.92c7.44 1.65 12.4 9.92 9.92 17.36"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ember(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M639.9 254.6c-1.1-10.7-10.7-6.8-10.7-6.8s-15.6 12.1-29.3 10.7c-13.7-1.3-9.4-32-9.4-32s3-28.1-5.1-30.4c-8.1-2.4-18 7.3-18 7.3s-12.4 13.7-18.3 31.2l-1.6.5s1.9-30.6-.3-37.6c-1.6-3.5-16.4-3.2-18.8 3s-14.2 49.2-15 67.2c0 0-23.1 19.6-43.3 22.8s-25-9.4-25-9.4s54.8-15.3 52.9-59.1s-44.2-27.6-49-24c-4.6 3.5-29.4 18.4-36.6 59.7c-.2 1.4-.7 7.5-.7 7.5s-21.2 14.2-33 18c0 0 33-55.6-7.3-80.9c-11.4-6.8-21.3-.5-27.2 5.3c13.6-17.3 46.4-64.2 36.9-105.2c-5.8-24.4-18-27.1-29.2-23.1c-17 6.7-23.5 16.7-23.5 16.7s-22 32-27.1 79.5s-12.6 105.1-12.6 105.1s-10.5 10.2-20.2 10.7s-5.4-28.7-5.4-28.7s7.5-44.6 7-52.1s-1.1-11.6-9.9-14.2c-8.9-2.7-18.5 8.6-18.5 8.6s-25.5 38.7-27.7 44.6l-1.3 2.4l-1.3-1.6s18-52.7.8-53.5s-28.5 18.8-28.5 18.8s-19.6 32.8-20.4 36.5l-1.3-1.6s8.1-38.2 6.4-47.6c-1.6-9.4-10.5-7.5-10.5-7.5s-11.3-1.3-14.2 5.9s-13.7 55.3-15 70.7c0 0-28.2 20.2-46.8 20.4c-18.5.3-16.7-11.8-16.7-11.8s68-23.3 49.4-69.2c-8.3-11.8-18-15.5-31.7-15.3c-13.7.3-30.3 8.6-41.3 33.3c-5.3 11.8-6.8 23-7.8 31.5c0 0-12.3 2.4-18.8-2.9s-10 0-10 0s-11.2 14-.1 18.3s28.1 6.1 28.1 6.1c1.6 7.5 6.2 19.5 19.6 29.7c20.2 15.3 58.8-1.3 58.8-1.3l15.9-8.8s.5 14.6 12.1 16.7s16.4 1 36.5-47.9c11.8-25 12.6-23.6 12.6-23.6l1.3-.3s-9.1 46.8-5.6 59.7C187.7 319.4 203 318 203 318s8.3 2.4 15-21.2s19.6-49.9 19.6-49.9h1.6s-5.6 48.1 3 63.7s30.9 5.3 30.9 5.3s15.6-7.8 18-10.2c0 0 18.5 15.8 44.6 12.9c58.3-11.5 79.1-25.9 79.1-25.9s10 24.4 41.1 26.7c35.5 2.7 54.8-18.6 54.8-18.6s-.3 13.5 12.1 18.6s20.7-22.8 20.7-22.8l20.7-57.2h1.9s1.1 37.3 21.5 43.2s47-13.7 47-13.7s6.4-3.5 5.3-14.3m-578 5.3c.8-32 21.8-45.9 29-39c7.3 7 4.6 22-9.1 31.4c-13.7 9.5-19.9 7.6-19.9 7.6m272.8-123.8s19.1-49.7 23.6-25.5s-40 96.2-40 96.2c.5-16.2 16.4-70.7 16.4-70.7m22.8 138.4c-12.6 33-43.3 19.6-43.3 19.6s-3.5-11.8 6.4-44.9s33.3-20.2 33.3-20.2s16.2 12.4 3.6 45.5m84.6-14.6s-3-10.5 8.1-30.6c11-20.2 19.6-9.1 19.6-9.1s9.4 10.2-1.3 25.5s-26.4 14.2-26.4 14.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Empire(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M287.6 54.2c-10.8-2.2-22.1-3.3-33.5-3.6V32.4c78.1 2.2 146.1 44 184.6 106.6l-15.8 9.1c-6.1-9.7-12.7-18.8-20.2-27.1l-18 15.5c-26-29.6-61.4-50.7-101.9-58.4zM53.4 322.4l23-7.7c-6.4-18.3-10-38.2-10-58.7s3.3-40.4 9.7-58.7l-22.7-7.7c3.6-10.8 8.3-21.3 13.6-31l-15.8-9.1C34 181 24.1 217.5 24.1 256s10 75 27.1 106.6l15.8-9.1c-5.3-10-9.7-20.3-13.6-31.1M213.1 434c-40.4-8-75.8-29.1-101.9-58.7l-18 15.8c-7.5-8.6-14.4-17.7-20.2-27.4l-16 9.4c38.5 62.3 106.8 104.3 184.9 106.6v-18.3c-11.3-.3-22.7-1.7-33.5-3.6zM93.3 120.9l18 15.5c26-29.6 61.4-50.7 101.9-58.4l-4.7-23.8c10.8-2.2 22.1-3.3 33.5-3.6V32.4C163.9 34.6 95.9 76.4 57.4 139l15.8 9.1c6-9.7 12.6-18.9 20.1-27.2m309.4 270.2l-18-15.8c-26 29.6-61.4 50.7-101.9 58.7l4.7 23.8c-10.8 1.9-22.1 3.3-33.5 3.6v18.3c78.1-2.2 146.4-44.3 184.9-106.6l-16.1-9.4c-5.7 9.7-12.6 18.8-20.1 27.4M496 256c0 137-111 248-248 248S0 393 0 256S111 8 248 8s248 111 248 248m-12.2 0c0-130.1-105.7-235.8-235.8-235.8S12.2 125.9 12.2 256S117.9 491.8 248 491.8S483.8 386.1 483.8 256m-39-106.6l-15.8 9.1c5.3 9.7 10 20.2 13.6 31l-22.7 7.7c6.4 18.3 9.7 38.2 9.7 58.7s-3.6 40.4-10 58.7l23 7.7c-3.9 10.8-8.3 21-13.6 31l15.8 9.1C462 331 471.9 294.5 471.9 256s-9.9-75-27.1-106.6m-183 177.7c16.3-3.3 30.4-11.6 40.7-23.5l51.2 44.8c11.9-13.6 21.3-29.3 27.1-46.8l-64.2-22.1c2.5-7.5 3.9-15.2 3.9-23.5s-1.4-16.1-3.9-23.5l64.5-22.1c-6.1-17.4-15.5-33.2-27.4-46.8l-51.2 44.8c-10.2-11.9-24.4-20.5-40.7-23.8l13.3-66.4c-8.6-1.9-17.7-2.8-27.1-2.8c-9.4 0-18.5.8-27.1 2.8l13.3 66.4c-16.3 3.3-30.4 11.9-40.7 23.8l-51.2-44.8c-11.9 13.6-21.3 29.3-27.4 46.8l64.5 22.1c-2.5 7.5-3.9 15.2-3.9 23.5s1.4 16.1 3.9 23.5l-64.2 22.1c5.8 17.4 15.2 33.2 27.1 46.8l51.2-44.8c10.2 11.9 24.4 20.2 40.7 23.5l-13.3 66.7c8.6 1.7 17.7 2.8 27.1 2.8c9.4 0 18.5-1.1 27.1-2.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Envira(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 32c477.6 0 366.6 317.3 367.1 366.3L448 480h-26l-70.4-71.2c-39 4.2-124.4 34.5-214.4-37C47 300.3 52 214.7 0 32m79.7 46c-49.7-23.5-5.2 9.2-5.2 9.2c45.2 31.2 66 73.7 90.2 119.9c31.5 60.2 79 139.7 144.2 167.7c65 28 34.2 12.5 6-8.5c-28.2-21.2-68.2-87-91-130.2c-31.7-60-61-118.6-144.2-158.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Erlang(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M87.2 53.5H0v405h100.4c-49.7-52.6-78.8-125.3-78.7-212.1c-.1-76.7 24-142.7 65.5-192.9m238.2 9.7c-45.9.1-85.1 33.5-89.2 83.2h169.9c-1.1-49.7-34.5-83.1-80.7-83.2m230.7-9.6h.3l-.1-.1zm.3 0c31.4 42.7 48.7 97.5 46.2 162.7c.5 6 .5 11.7 0 24.1H230.2c-.2 109.7 38.9 194.9 138.6 195.3c68.5-.3 118-51 151.9-106.1l96.4 48.2c-17.4 30.9-36.5 57.8-57.9 80.8H640v-405z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ethereum(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M311.9 260.8L160 353.6L8 260.8L160 0zM160 383.4L8 290.6L160 512l152-221.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Etsy(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 348c-1.75 10.75-13.75 110-15.5 132c-117.879-4.299-219.895-4.743-368.5 0v-25.5c45.457-8.948 60.627-8.019 61-35.25c1.793-72.322 3.524-244.143 0-322c-1.029-28.46-12.13-26.765-61-36v-25.5c73.886 2.358 255.933 8.551 362.999-3.75c-3.5 38.25-7.75 126.5-7.75 126.5H332C320.947 115.665 313.241 68 277.25 68h-137c-10.25 0-10.75 3.5-10.75 9.75V241.5c58 .5 88.5-2.5 88.5-2.5c29.77-.951 27.56-8.502 40.75-65.251h25.75c-4.407 101.351-3.91 61.829-1.75 160.25H257c-9.155-40.086-9.065-61.045-39.501-61.5c0 0-21.5-2-88-2v139c0 26 14.25 38.25 44.25 38.25H263c63.636 0 66.564-24.996 98.751-99.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Evernote(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M120.82 132.21c1.6 22.31-17.55 21.59-21.61 21.59c-68.93 0-73.64-1-83.58 3.34c-.56.22-.74 0-.37-.37L123.79 46.45c.38-.37.6-.22.38.37c-4.35 9.99-3.35 15.09-3.35 85.39m79 308c-14.68-37.08 13-76.93 52.52-76.62c17.49 0 22.6 23.21 7.95 31.42c-6.19 3.3-24.95 1.74-25.14 19.2c-.05 17.09 19.67 25 31.2 24.89A45.64 45.64 0 0 0 312 393.45v-.08c0-11.63-7.79-47.22-47.54-55.34c-7.72-1.54-65-6.35-68.35-50.52c-3.74 16.93-17.4 63.49-43.11 69.09c-8.74 1.94-69.68 7.64-112.92-36.77c0 0-18.57-15.23-28.23-57.95c-3.38-15.75-9.28-39.7-11.14-62c0-18 11.14-30.45 25.07-32.2c81 0 90 2.32 101-7.8c9.82-9.24 7.8-15.5 7.8-102.78c1-8.3 7.79-30.81 53.41-24.14c6 .86 31.91 4.18 37.48 30.64l64.26 11.15c20.43 3.71 70.94 7 80.6 57.94c22.66 121.09 8.91 238.46 7.8 238.46C362.15 485.53 267.06 480 267.06 480c-18.95-.23-54.25-9.4-67.27-39.83zm80.94-204.84c-1 1.92-2.2 6 .85 7c14.09 4.93 39.75 6.84 45.88 5.53c3.11-.25 3.05-4.43 2.48-6.65c-3.53-21.85-40.83-26.5-49.24-5.92z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Expeditedssl(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M248 43.4C130.6 43.4 35.4 138.6 35.4 256S130.6 468.6 248 468.6S460.6 373.4 460.6 256S365.4 43.4 248 43.4m-97.4 132.9c0-53.7 43.7-97.4 97.4-97.4s97.4 43.7 97.4 97.4v26.6c0 5-3.9 8.9-8.9 8.9h-17.7c-5 0-8.9-3.9-8.9-8.9v-26.6c0-82.1-124-82.1-124 0v26.6c0 5-3.9 8.9-8.9 8.9h-17.7c-5 0-8.9-3.9-8.9-8.9v-26.6zM389.7 380c0 9.7-8 17.7-17.7 17.7H124c-9.7 0-17.7-8-17.7-17.7V238.3c0-9.7 8-17.7 17.7-17.7h248c9.7 0 17.7 8 17.7 17.7zm-248-137.3v132.9c0 2.5-1.9 4.4-4.4 4.4h-8.9c-2.5 0-4.4-1.9-4.4-4.4V242.7c0-2.5 1.9-4.4 4.4-4.4h8.9c2.5 0 4.4 1.9 4.4 4.4m141.7 48.7c0 13-7.2 24.4-17.7 30.4v31.6c0 5-3.9 8.9-8.9 8.9h-17.7c-5 0-8.9-3.9-8.9-8.9v-31.6c-10.5-6.1-17.7-17.4-17.7-30.4c0-19.7 15.8-35.4 35.4-35.4s35.5 15.8 35.5 35.4M248 8C111 8 0 119 0 256s111 248 248 248s248-111 248-248S385 8 248 8m0 478.3C121 486.3 17.7 383 17.7 256S121 25.7 248 25.7S478.3 129 478.3 256S375 486.3 248 486.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Facebook(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M504 256C504 119 393 8 256 8S8 119 8 256c0 123.78 90.69 226.38 209.25 245V327.69h-63V256h63v-54.64c0-62.15 37-96.48 93.67-96.48c27.14 0 55.52 4.84 55.52 4.84v61h-31.28c-30.8 0-40.41 19.12-40.41 38.73V256h68.78l-11 71.69h-57.78V501C413.31 482.38 504 379.78 504 256"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacebookF(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m279.14 288l14.22-92.66h-88.91v-60.13c0-25.35 12.42-50.06 52.24-50.06h40.42V6.26S260.43 0 225.36 0c-73.22 0-121.08 44.38-121.08 124.72v70.62H22.89V288h81.39v224h100.17V288z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacebookMessenger(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256.55 8C116.52 8 8 110.34 8 248.57c0 72.3 29.71 134.78 78.07 177.94c8.35 7.51 6.63 11.86 8.05 58.23A19.92 19.92 0 0 0 122 502.31c52.91-23.3 53.59-25.14 62.56-22.7C337.85 521.8 504 423.7 504 248.57C504 110.34 396.59 8 256.55 8m149.24 185.13l-73 115.57a37.37 37.37 0 0 1-53.91 9.93l-58.08-43.47a15 15 0 0 0-18 0l-78.37 59.44c-10.46 7.93-24.16-4.6-17.11-15.67l73-115.57a37.36 37.36 0 0 1 53.91-9.93l58.06 43.46a15 15 0 0 0 18 0l78.41-59.38c10.44-7.98 24.14 4.54 17.09 15.62"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacebookSquare(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 32H48A48 48 0 0 0 0 80v352a48 48 0 0 0 48 48h137.25V327.69h-63V256h63v-54.64c0-62.15 37-96.48 93.67-96.48c27.14 0 55.52 4.84 55.52 4.84v61h-31.27c-30.81 0-40.42 19.12-40.42 38.73V256h68.78l-11 71.69h-57.78V480H400a48 48 0 0 0 48-48V80a48 48 0 0 0-48-48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FantasyFlightGames(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 32.86L32.86 256L256 479.14L479.14 256zM88.34 255.83c1.96-2 11.92-12.3 96.49-97.48c41.45-41.75 86.19-43.77 119.77-18.69c24.63 18.4 62.06 58.9 62.15 59c.68.74 1.07 2.86.58 3.38c-11.27 11.84-22.68 23.54-33.5 34.69c-34.21-32.31-40.52-38.24-48.51-43.95c-17.77-12.69-41.4-10.13-56.98 5.1c-2.17 2.13-1.79 3.43.12 5.35c2.94 2.95 28.1 28.33 35.09 35.78c-11.95 11.6-23.66 22.97-35.69 34.66c-12.02-12.54-24.48-25.53-36.54-38.11c-21.39 21.09-41.69 41.11-61.85 60.99a42569.01 42569.01 0 0 1-41.13-40.72m234.82 101.6c-35.49 35.43-78.09 38.14-106.99 20.47c-22.08-13.5-39.38-32.08-72.93-66.84c12.05-12.37 23.79-24.42 35.37-36.31c33.02 31.91 37.06 36.01 44.68 42.09c18.48 14.74 42.52 13.67 59.32-1.8c3.68-3.39 3.69-3.64.14-7.24c-10.59-10.73-21.19-21.44-31.77-32.18c-1.32-1.34-3.03-2.48-.8-4.69c10.79-10.71 21.48-21.52 32.21-32.29c.26-.26.65-.38 1.91-1.07c12.37 12.87 24.92 25.92 37.25 38.75c21.01-20.73 41.24-40.68 61.25-60.42c13.68 13.4 27.13 26.58 40.86 40.03c-20.17 20.86-81.68 82.71-100.5 101.5M256 0L0 256l256 256l256-256zM16 256L256 16l240 240l-240 240z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fedex(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m586 284.5l53.3-59.9h-62.4l-21.7 24.8l-22.5-24.8H414v-16h56.1v-48.1H318.9V236h-.5c-9.6-11-21.5-14.8-35.4-14.8c-28.4 0-49.8 19.4-57.3 44.9c-18-59.4-97.4-57.6-121.9-14v-24.2H49v-26.2h60v-41.1H0V345h49v-77.5h48.9c-1.5 5.7-2.3 11.8-2.3 18.2c0 73.1 102.6 91.4 130.2 23.7h-42c-14.7 20.9-45.8 8.9-45.8-14.6h85.5c3.7 30.5 27.4 56.9 60.1 56.9c14.1 0 27-6.9 34.9-18.6h.5V345h212.2l22.1-25l22.3 25H640zm-446.7-16.6c6.1-26.3 41.7-25.6 46.5 0zm153.4 48.9c-34.6 0-34-62.8 0-62.8c32.6 0 34.5 62.8 0 62.8m167.8 19.1h-94.4V169.4h95v30.2H405v33.9h55.5v28.1h-56.1v44.7h56.1zm-45.9-39.8v-24.4h56.1v-44l50.7 57l-50.7 57v-45.6zm138.6 10.3l-26.1 29.5H489l45.6-51.2l-45.6-51.2h39.7l26.6 29.3l25.6-29.3h38.5l-45.4 51l46 51.4h-40.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fedora(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M225 32C101.3 31.7.8 131.7.4 255.4L0 425.7a53.6 53.6 0 0 0 53.6 53.9l170.2.4c123.7.3 224.3-99.7 224.6-223.4S348.7 32.3 225 32m169.8 157.2L333 126.6c2.3-4.7 3.8-9.2 3.8-14.3v-1.6l55.2 56.1a101 101 0 0 1 2.8 22.4M331 94.3a106.06 106.06 0 0 1 58.5 63.8l-54.3-54.6a26.48 26.48 0 0 0-4.2-9.2M118.1 247.2a49.66 49.66 0 0 0-7.7 11.4l-8.5-8.5a85.78 85.78 0 0 1 16.2-2.9M97 251.4l11.8 11.9l-.9 8a34.74 34.74 0 0 0 2.4 12.5l-27-27.2a80.6 80.6 0 0 1 13.7-5.2m-18.2 7.4l38.2 38.4a53.17 53.17 0 0 0-14.1 4.7L67.6 266a107 107 0 0 1 11.2-7.2m-15.2 9.8l35.3 35.5a67.25 67.25 0 0 0-10.5 8.5L53.5 278a64.33 64.33 0 0 1 10.1-9.4m-13.3 12.3l34.9 35a56.84 56.84 0 0 0-7.7 11.4l-35.8-35.9c2.8-3.8 5.7-7.2 8.6-10.5m-11 14.3l36.4 36.6a48.29 48.29 0 0 0-3.6 15.2l-39.5-39.8a99.81 99.81 0 0 1 6.7-12m-8.8 16.3l41.3 41.8a63.47 63.47 0 0 0 6.7 26.2L25.8 326c1.4-4.9 2.9-9.6 4.7-14.5m-7.9 43l61.9 62.2a31.24 31.24 0 0 0-3.6 14.3v1.1l-55.4-55.7a88.27 88.27 0 0 1-2.9-21.9m5.3 30.7l54.3 54.6a28.44 28.44 0 0 0 4.2 9.2a106.32 106.32 0 0 1-58.5-63.8m-5.3-37a80.69 80.69 0 0 1 2.1-17l72.2 72.5a37.59 37.59 0 0 0-9.9 8.7zm253.3-51.8l-42.6-.1l-.1 56c-.2 69.3-64.4 115.8-125.7 102.9c-5.7 0-19.9-8.7-19.9-24.2a24.89 24.89 0 0 1 24.5-24.6c6.3 0 6.3 1.6 15.7 1.6a55.91 55.91 0 0 0 56.1-55.9l.1-47c0-4.5-4.5-9-8.9-9l-33.6-.1c-32.6-.1-32.5-49.4.1-49.3l42.6.1l.1-56a105.18 105.18 0 0 1 105.6-105a86.35 86.35 0 0 1 20.2 2.3c11.2 1.8 19.9 11.9 19.9 24c0 15.5-14.9 27.8-30.3 23.9c-27.4-5.9-65.9 14.4-66 54.9l-.1 47a8.94 8.94 0 0 0 8.9 9l33.6.1c32.5.2 32.4 49.5-.2 49.4m23.5-.3a35.58 35.58 0 0 0 7.6-11.4l8.5 8.5a102 102 0 0 1-16.1 2.9m21-4.2L308.6 280l.9-8.1a34.74 34.74 0 0 0-2.4-12.5l27 27.2a74.89 74.89 0 0 1-13.7 5.3m18-7.4l-38-38.4c4.9-1.1 9.6-2.4 13.7-4.7l36.2 35.9c-3.8 2.5-7.9 5-11.9 7.2m15.5-9.8l-35.3-35.5a61.06 61.06 0 0 0 10.5-8.5l34.9 35a124.56 124.56 0 0 1-10.1 9m13.2-12.3l-34.9-35a63.18 63.18 0 0 0 7.7-11.4l35.8 35.9a130.28 130.28 0 0 1-8.6 10.5m11-14.3l-36.4-36.6a48.29 48.29 0 0 0 3.6-15.2l39.5 39.8a87.72 87.72 0 0 1-6.7 12m13.5-30.9a140.63 140.63 0 0 1-4.7 14.3L345.6 190a58.19 58.19 0 0 0-7.1-26.2zm1-5.6l-71.9-72.1a32 32 0 0 0 9.9-9.2l64.3 64.7a90.93 90.93 0 0 1-2.3 16.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Figma(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 95.792C14 42.888 56.888 0 109.793 0h164.368c52.905 0 95.793 42.888 95.793 95.792c0 33.5-17.196 62.984-43.243 80.105c26.047 17.122 43.243 46.605 43.243 80.105c0 52.905-42.888 95.793-95.793 95.793h-2.08c-24.802 0-47.403-9.426-64.415-24.891v88.263c0 53.61-44.009 96.833-97.357 96.833C57.536 512 14 469.243 14 416.207c0-33.498 17.195-62.98 43.24-80.102C31.194 318.983 14 289.5 14 256.002c0-33.5 17.196-62.983 43.243-80.105C31.196 158.776 14 129.292 14 95.792m162.288 95.795h-66.495c-35.576 0-64.415 28.84-64.415 64.415c0 35.438 28.617 64.192 64.003 64.414l.412-.001h66.495zm31.378 64.415c0 35.575 28.839 64.415 64.415 64.415h2.08c35.576 0 64.415-28.84 64.415-64.415s-28.839-64.415-64.415-64.415h-2.08c-35.576 0-64.415 28.84-64.415 64.415m-97.873 95.793l-.412-.001c-35.386.221-64.003 28.975-64.003 64.413c0 35.445 29.225 64.415 64.931 64.415c36.282 0 65.979-29.436 65.979-65.455v-63.372zm0-320.417c-35.576 0-64.415 28.84-64.415 64.414c0 35.576 28.84 64.415 64.415 64.415h66.495V31.377zm97.873 128.829h66.495c35.576 0 64.415-28.839 64.415-64.415c0-35.575-28.839-64.414-64.415-64.414h-66.495z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Firefox(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M503.52 241.48c-.12-1.56-.24-3.12-.24-4.68v-.12l-.36-4.68v-.12a245.86 245.86 0 0 0-7.32-41.15c0-.12 0-.12-.12-.24l-1.08-4c-.12-.24-.12-.48-.24-.6c-.36-1.2-.72-2.52-1.08-3.72c-.12-.24-.12-.6-.24-.84c-.36-1.2-.72-2.4-1.08-3.48c-.12-.36-.24-.6-.36-1c-.36-1.2-.72-2.28-1.2-3.48l-.36-1.08c-.36-1.08-.84-2.28-1.2-3.36a8.27 8.27 0 0 0-.36-1c-.48-1.08-.84-2.28-1.32-3.36c-.12-.24-.24-.6-.36-.84c-.48-1.2-1-2.28-1.44-3.48c0-.12-.12-.24-.12-.36c-1.56-3.84-3.24-7.68-5-11.4l-.36-.72c-.48-1-.84-1.8-1.32-2.64c-.24-.48-.48-1.08-.72-1.56c-.36-.84-.84-1.56-1.2-2.4c-.36-.6-.6-1.2-1-1.8s-.84-1.44-1.2-2.28c-.36-.6-.72-1.32-1.08-1.92s-.84-1.44-1.2-2.16a18.07 18.07 0 0 0-1.2-2c-.36-.72-.84-1.32-1.2-2s-.84-1.32-1.2-2s-.84-1.32-1.2-1.92s-.84-1.44-1.32-2.16a15.63 15.63 0 0 0-1.2-1.8L463.2 119a15.63 15.63 0 0 0-1.2-1.8c-.48-.72-1.08-1.56-1.56-2.28c-.36-.48-.72-1.08-1.08-1.56l-1.8-2.52c-.36-.48-.6-.84-1-1.32c-1-1.32-1.8-2.52-2.76-3.72a248.76 248.76 0 0 0-23.51-26.64A186.82 186.82 0 0 0 412 62.46c-4-3.48-8.16-6.72-12.48-9.84a162.49 162.49 0 0 0-24.6-15.12c-2.4-1.32-4.8-2.52-7.2-3.72a254 254 0 0 0-55.43-19.56c-1.92-.36-3.84-.84-5.64-1.2h-.12c-1-.12-1.8-.36-2.76-.48a236.35 236.35 0 0 0-38-4h-10.63a234.62 234.62 0 0 0-45.48 5c-33.59 7.08-63.23 21.24-82.91 39c-1.08 1-1.92 1.68-2.4 2.16l-.48.48h.13l-.12.12l.12-.12a.12.12 0 0 0 .12-.12l-.12.12a.42.42 0 0 1 .24-.12c14.64-8.76 34.92-16 49.44-19.56l5.88-1.44c.36-.12.84-.12 1.2-.24c1.68-.36 3.36-.72 5.16-1.08c.24 0 .6-.12.84-.12C250.94 20.94 319.34 40.14 367 85.61a171.49 171.49 0 0 1 26.88 32.76c30.36 49.2 27.48 111.11 3.84 147.59c-34.44 53-111.35 71.27-159 24.84a84.19 84.19 0 0 1-25.56-59a74.05 74.05 0 0 1 6.24-31c1.68-3.84 13.08-25.67 18.24-24.59c-13.08-2.76-37.55 2.64-54.71 28.19c-15.36 22.92-14.52 58.2-5 83.28a132.85 132.85 0 0 1-12.12-39.24c-12.24-82.55 43.31-153 94.31-170.51c-27.48-24-96.47-22.31-147.71 15.36c-29.88 22-51.23 53.16-62.51 90.36c1.68-20.88 9.6-52.08 25.8-83.88c-17.16 8.88-39 37-49.8 62.88c-15.6 37.43-21 82.19-16.08 124.79c.36 3.24.72 6.36 1.08 9.6c19.92 117.11 122 206.38 244.78 206.38C392.77 503.42 504 392.19 504 255c-.12-4.52-.24-9.08-.48-13.52"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FirefoxBrowser(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M130.63 123.49c.16.01.08.01 0 0m351.42 45.35c-10.61-25.5-32.08-53-48.94-61.73c13.72 26.89 21.67 53.88 24.7 74c0 0 0 .14.05.41c-27.58-68.75-74.35-96.47-112.55-156.83c-1.93-3.05-3.86-6.11-5.74-9.33c-1-1.65-1.86-3.34-2.69-5.05a44.88 44.88 0 0 1-3.64-9.62a.63.63 0 0 0-.55-.66a.9.9 0 0 0-.46 0l-.12.07l-.18.1l.1-.14c-54.23 31.77-76.72 87.38-82.5 122.78a130 130 0 0 0-48.33 12.33a6.25 6.25 0 0 0-3.09 7.75a6.13 6.13 0 0 0 7.79 3.79l.52-.21a117.84 117.84 0 0 1 42.11-11l1.42-.1c2-.12 4-.2 6-.22A122.61 122.61 0 0 1 291 140c.67.2 1.32.42 2 .63c1.89.57 3.76 1.2 5.62 1.87c1.36.5 2.71 1 4.05 1.58c1.09.44 2.18.88 3.25 1.35q2.52 1.13 5 2.35c.75.37 1.5.74 2.25 1.13q2.4 1.26 4.74 2.63q1.51.87 3 1.8a124.89 124.89 0 0 1 42.66 44.13c-13-9.15-36.35-18.19-58.82-14.28c87.74 43.86 64.18 194.9-57.39 189.2a108.43 108.43 0 0 1-31.74-6.12a139.5 139.5 0 0 1-7.16-2.93c-1.38-.63-2.76-1.27-4.12-2c-29.84-15.34-54.44-44.42-57.51-79.75c0 0 11.25-41.95 80.62-41.95c7.5 0 28.93-20.92 29.33-27c-.09-2-42.54-18.87-59.09-35.18c-8.85-8.71-13.05-12.91-16.77-16.06a69.58 69.58 0 0 0-6.31-4.77a113.05 113.05 0 0 1-.69-59.63c-25.06 11.41-44.55 29.45-58.71 45.37h-.12c-9.67-12.25-9-52.65-8.43-61.08c-.12-.53-7.22 3.68-8.15 4.31a178.54 178.54 0 0 0-23.84 20.43a214 214 0 0 0-22.77 27.33a205.84 205.84 0 0 0-32.73 73.9c-.06.27-2.33 10.21-4 22.48q-.42 2.87-.78 5.74c-.57 3.69-1 7.71-1.44 14c0 .24 0 .48-.05.72c-.18 2.71-.34 5.41-.49 8.12v1.24c0 134.7 109.21 243.89 243.92 243.89c120.64 0 220.82-87.58 240.43-202.62c.41-3.12.74-6.26 1.11-9.41c4.85-41.83-.54-85.79-15.82-122.55Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FirstOrder(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.9 229.2c.1-.1.2-.3.3-.4c0 .1 0 .3-.1.4zM224 96.6c-7.1 0-14.6.6-21.4 1.7l3.7 67.4l-22-64c-14.3 3.7-27.7 9.4-40 16.6l29.4 61.4l-45.1-50.9c-11.4 8.9-21.7 19.1-30.6 30.9l50.6 45.4l-61.1-29.7c-7.1 12.3-12.9 25.7-16.6 40l64.3 22.6l-68-4c-.9 7.1-1.4 14.6-1.4 22s.6 14.6 1.4 21.7l67.7-4l-64 22.6c3.7 14.3 9.4 27.7 16.6 40.3l61.1-29.7L97.7 352c8.9 11.7 19.1 22.3 30.9 30.9l44.9-50.9l-29.5 61.4c12.3 7.4 25.7 13.1 40 16.9l22.3-64.6l-4 68c7.1 1.1 14.6 1.7 21.7 1.7c7.4 0 14.6-.6 21.7-1.7l-4-68.6l22.6 65.1c14.3-4 27.7-9.4 40-16.9L274.9 332l44.9 50.9c11.7-8.9 22-19.1 30.6-30.9l-50.6-45.1l61.1 29.4c7.1-12.3 12.9-25.7 16.6-40.3l-64-22.3l67.4 4c1.1-7.1 1.4-14.3 1.4-21.7s-.3-14.9-1.4-22l-67.7 4l64-22.3c-3.7-14.3-9.1-28-16.6-40.3l-60.9 29.7l50.6-45.4c-8.9-11.7-19.1-22-30.6-30.9l-45.1 50.9l29.4-61.1c-12.3-7.4-25.7-13.1-40-16.9L241.7 166l4-67.7c-7.1-1.2-14.3-1.7-21.7-1.7M443.4 128v256L224 512L4.6 384V128L224 0zm-17.1 10.3L224 20.9L21.7 138.3v235.1L224 491.1l202.3-117.7zM224 37.1l187.7 109.4v218.9L224 474.9L36.3 365.4V146.6zm0 50.9c-92.3 0-166.9 75.1-166.9 168c0 92.6 74.6 167.7 166.9 167.7c92 0 166.9-75.1 166.9-167.7c0-92.9-74.9-168-166.9-168"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FirstOrderAlt(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M248 8C111.03 8 0 119.03 0 256s111.03 248 248 248s248-111.03 248-248S384.97 8 248 8m0 488.21C115.34 496.21 7.79 388.66 7.79 256S115.34 15.79 248 15.79S488.21 123.34 488.21 256S380.66 496.21 248 496.21m0-459.92C126.66 36.29 28.29 134.66 28.29 256S126.66 475.71 248 475.71S467.71 377.34 467.71 256S369.34 36.29 248 36.29m0 431.22c-116.81 0-211.51-94.69-211.51-211.51S131.19 44.49 248 44.49S459.51 139.19 459.51 256S364.81 467.51 248 467.51m186.23-162.98a191.613 191.613 0 0 1-20.13 48.69l-74.13-35.88l61.48 54.82a193.515 193.515 0 0 1-37.2 37.29l-54.8-61.57l35.88 74.27a190.944 190.944 0 0 1-48.63 20.23l-27.29-78.47l4.79 82.93c-8.61 1.18-17.4 1.8-26.33 1.8s-17.72-.62-26.33-1.8l4.76-82.46l-27.15 78.03a191.365 191.365 0 0 1-48.65-20.2l35.93-74.34l-54.87 61.64a193.85 193.85 0 0 1-37.22-37.28l61.59-54.9l-74.26 35.93a191.638 191.638 0 0 1-20.14-48.69l77.84-27.11l-82.23 4.76c-1.16-8.57-1.78-17.32-1.78-26.21c0-9 .63-17.84 1.82-26.51l82.38 4.77l-77.94-27.16a191.726 191.726 0 0 1 20.23-48.67l74.22 35.92l-61.52-54.86a193.85 193.85 0 0 1 37.28-37.22l54.76 61.53l-35.83-74.17a191.49 191.49 0 0 1 48.65-20.13l26.87 77.25l-4.71-81.61c8.61-1.18 17.39-1.8 26.32-1.8s17.71.62 26.32 1.8l-4.74 82.16l27.05-77.76c17.27 4.5 33.6 11.35 48.63 20.17l-35.82 74.12l54.72-61.47a193.13 193.13 0 0 1 37.24 37.23l-61.45 54.77l74.12-35.86a191.515 191.515 0 0 1 20.2 48.65l-77.81 27.1l82.24-4.75c1.19 8.66 1.82 17.5 1.82 26.49c0 8.88-.61 17.63-1.78 26.19l-82.12-4.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Firstdraft(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 192h-64v128H192v128H0v-25.6h166.4v-128h128v-128H384zm-25.6 38.4v128h-128v128H64V512h192V384h128V230.4zm25.6 192h-89.6V512H320v-64h64zM0 0v384h128V256h128V128h128V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FiveHundredPx(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M103.3 344.3c-6.5-14.2-6.9-18.3 7.4-23.1c25.6-8 8 9.2 43.2 49.2h.3v-93.9c1.2-50.2 44-92.2 97.7-92.2c53.9 0 97.7 43.5 97.7 96.8c0 63.4-60.8 113.2-128.5 93.3c-10.5-4.2-2.1-31.7 8.5-28.6c53 0 89.4-10.1 89.4-64.4c0-61-77.1-89.6-116.9-44.6c-23.5 26.4-17.6 42.1-17.6 157.6c50.7 31 118.3 22 160.4-20.1c24.8-24.8 38.5-58 38.5-93c0-35.2-13.8-68.2-38.8-93.3c-24.8-24.8-57.8-38.5-93.3-38.5s-68.8 13.8-93.5 38.5c-.3.3-16 16.5-21.2 23.9l-.5.6c-3.3 4.7-6.3 9.1-20.1 6.1c-6.9-1.7-14.3-5.8-14.3-11.8V20c0-5 3.9-10.5 10.5-10.5h241.3c8.3 0 8.3 11.6 8.3 15.1c0 3.9 0 15.1-8.3 15.1H130.3v132.9h.3c104.2-109.8 282.8-36 282.8 108.9c0 178.1-244.8 220.3-310.1 62.8m63.3-260.8c-.5 4.2 4.6 24.5 14.6 20.6C306 56.6 384 144.5 390.6 144.5c4.8 0 22.8-15.3 14.3-22.8c-93.2-89-234.5-57-238.3-38.2M393 414.7C283 524.6 94 475.5 61 310.5c0-12.2-30.4-7.4-28.9 3.3c24 173.4 246 256.9 381.6 121.3c6.9-7.8-12.6-28.4-20.7-20.4M213.6 306.6c0 4 4.3 7.3 5.5 8.5c3 3 6.1 4.4 8.5 4.4c3.8 0 2.6.2 22.3-19.5c19.6 19.3 19.1 19.5 22.3 19.5c5.4 0 18.5-10.4 10.7-18.2L265.6 284l18.2-18.2c6.3-6.8-10.1-21.8-16.2-15.7L249.7 268c-18.6-18.8-18.4-19.5-21.5-19.5c-5 0-18 11.7-12.4 17.3L234 284c-18.1 17.9-20.4 19.2-20.4 22.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flickr(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 32H48C21.5 32 0 53.5 0 80v352c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V80c0-26.5-21.5-48-48-48M144.5 319c-35.1 0-63.5-28.4-63.5-63.5s28.4-63.5 63.5-63.5s63.5 28.4 63.5 63.5s-28.4 63.5-63.5 63.5m159 0c-35.1 0-63.5-28.4-63.5-63.5s28.4-63.5 63.5-63.5s63.5 28.4 63.5 63.5s-28.4 63.5-63.5 63.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flipboard(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 32v448h448V32zm358.4 179.2h-89.6v89.6h-89.6v89.6H89.6V121.6h268.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fly(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M197.8 427.8c12.9 11.7 33.7 33.3 33.2 50.7c0 .8-.1 1.6-.1 2.5c-1.8 19.8-18.8 31.1-39.1 31c-25-.1-39.9-16.8-38.7-35.8c1-16.2 20.5-36.7 32.4-47.6c2.3-2.1 2.7-2.7 5.6-3.6c3.4 0 3.9.3 6.7 2.8M331.9 67.3c-16.3-25.7-38.6-40.6-63.3-52.1C243.1 4.5 214-.2 192 0c-44.1 0-71.2 13.2-81.1 17.3C57.3 45.2 26.5 87.2 28 158.6c7.1 82.2 97 176 155.8 233.8c1.7 1.6 4.5 4.5 6.2 5.1l3.3.1c2.1-.7 1.8-.5 3.5-2.1c52.3-49.2 140.7-145.8 155.9-215.7c7-39.2 3.1-72.5-20.8-112.5M186.8 351.9c-28-51.1-65.2-130.7-69.3-189c-3.4-47.5 11.4-131.2 69.3-136.7zM328.7 180c-16.4 56.8-77.3 128-118.9 170.3C237.6 298.4 275 217 277 158.4c1.6-45.9-9.8-105.8-48-131.4c88.8 18.3 115.5 98.1 99.7 153"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FontAwesome(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 32H48A48 48 0 0 0 0 80v352a48 48 0 0 0 48 48h352a48 48 0 0 0 48-48V80a48 48 0 0 0-48-48m-64 280c-31.6 11.2-41.2 16-59.8 16c-31.4 0-43.2-16-74.6-16a80 80 0 0 0-25.6 4v-32a85.9 85.9 0 0 1 25.6-4c31.2 0 43.2 16 74.6 16c10.2 0 17.8-1.4 27.8-4.6v-96c-10 3.2-17.6 4.6-27.8 4.6c-31.4 0-43.2-16-74.6-16c-25.4 0-37.4 10.4-57.6 14.4V352a16 16 0 0 1-32 0V160a16 16 0 0 1 32 0v6.4c20.2-4 32.2-14.4 57.6-14.4c31.2 0 43.2 16 74.6 16c18.6 0 28.2-4.8 59.8-16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FontAwesomeAlt(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 32H48A48 48 0 0 0 0 80v352a48 48 0 0 0 48 48h352a48 48 0 0 0 48-48V80a48 48 0 0 0-48-48m16 400a16 16 0 0 1-16 16H48a16 16 0 0 1-16-16V80a16 16 0 0 1 16-16h352a16 16 0 0 1 16 16ZM201.6 152c-25.4 0-37.4 10.4-57.6 14.4V160a16 16 0 0 0-32 0v192a16 16 0 0 0 32 0V198.4c20.2-4 32.2-14.4 57.6-14.4c31.4 0 43.2 16 74.6 16c10.2 0 17.8-1.4 27.8-4.6v96c-10 3.2-17.6 4.6-27.8 4.6c-31.4 0-43.4-16-74.6-16a85.9 85.9 0 0 0-25.6 4v32a80 80 0 0 1 25.6-4c31.4 0 43.2 16 74.6 16c18.6 0 28.2-4.8 59.8-16V152c-31.6 11.2-41.2 16-59.8 16c-31.4 0-43.4-16-74.6-16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FontAwesomeFlag(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 48v336c-63 23-82 32-119 32c-63 0-87-32-150-32c-20 0-36 4-51 8v-64c15-4 31-8 51-8c63 0 87 32 150 32c20 0 35-3 55-9V135c-20 6-35 9-55 9c-63 0-87-32-150-32c-51 0-75 21-115 29v307a31.6 31.6 0 0 1-32 32a31.6 31.6 0 0 1-32-32V64a31.6 31.6 0 0 1 32-32a31.6 31.6 0 0 1 32 32v13c40-8 64-29 115-29c63 0 87 32 150 32c37 0 56-9 119-32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fonticons(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 32v448h448V32zm187 140.9c-18.4 0-19 9.9-19 27.4v23.3c0 2.4-3.5 4.4-.6 4.4h67.4l-11.1 37.3H168v112.9c0 5.8-2 6.7 3.2 7.3l43.5 4.1v25.1H84V389l21.3-2c5.2-.6 6.7-2.3 6.7-7.9V267.7c0-2.3-2.9-2.3-5.8-2.3H84V228h28v-21c0-49.6 26.5-70 77.3-70c34.1 0 64.7 8.2 64.7 52.8l-50.7 6.1c.3-18.7-4.4-23-16.3-23m74.3 241.8v-25.1l20.4-2.6c5.2-.6 7.6-1.7 7.6-7.3V271.8c0-4.1-2.9-6.7-6.7-7.9l-24.2-6.4l6.7-29.5h80.2v151.7c0 5.8-2.6 6.4 2.9 7.3l15.7 2.6v25.1zm80.8-255.5l9 33.2l-7.3 7.3l-31.2-16.6l-31.2 16.6l-7.3-7.3l9-33.2l-21.8-24.2l3.5-9.6h27.7l15.5-28h9.3l15.5 28h27.7l3.5 9.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FonticonsFi(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M114.4 224h92.4l-15.2 51.2h-76.4V433c0 8-2.8 9.2 4.4 10l59.6 5.6V483H0v-35.2l29.2-2.8c7.2-.8 9.2-3.2 9.2-10.8V278.4c0-3.2-4-3.2-8-3.2H0V224h38.4v-28.8c0-68 36.4-96 106-96c46.8 0 88.8 11.2 88.8 72.4l-69.6 8.4c.4-25.6-6-31.6-22.4-31.6c-25.2 0-26 13.6-26 37.6v32c0 3.2-4.8 6-.8 6M384 483H243.2v-34.4l28-3.6c7.2-.8 10.4-2.4 10.4-10V287c0-5.6-4-9.2-9.2-10.8l-33.2-8.8l9.2-40.4h110v208c0 8-3.6 8.8 4 10l21.6 3.6zm-30-347.2l12.4 45.6l-10 10l-42.8-22.8l-42.8 22.8l-10-10l12.4-45.6l-30-36.4l4.8-10h38L307.2 51H320l21.2 38.4h38l4.8 13.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FortAwesome(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M489.2 287.9h-27.4c-2.6 0-4.6 2-4.6 4.6v32h-36.6V146.2c0-2.6-2-4.6-4.6-4.6h-27.4c-2.6 0-4.6 2-4.6 4.6v32h-36.6v-32c0-2.6-2-4.6-4.6-4.6h-27.4c-2.6 0-4.6 2-4.6 4.6v32h-36.6v-32c0-6-8-4.6-11.7-4.6v-38c8.3-2 17.1-3.4 25.7-3.4c10.9 0 20.9 4.3 31.4 4.3c4.6 0 27.7-1.1 27.7-8v-60c0-2.6-2-4.6-4.6-4.6c-5.1 0-15.1 4.3-24 4.3c-9.7 0-20.9-4.3-32.6-4.3c-8 0-16 1.1-23.7 2.9v-4.9c5.4-2.6 9.1-8.3 9.1-14.3c0-20.7-31.4-20.8-31.4 0c0 6 3.7 11.7 9.1 14.3v111.7c-3.7 0-11.7-1.4-11.7 4.6v32h-36.6v-32c0-2.6-2-4.6-4.6-4.6h-27.4c-2.6 0-4.6 2-4.6 4.6v32H128v-32c0-2.6-2-4.6-4.6-4.6H96c-2.6 0-4.6 2-4.6 4.6v178.3H54.8v-32c0-2.6-2-4.6-4.6-4.6H22.8c-2.6 0-4.6 2-4.6 4.6V512h182.9v-96c0-72.6 109.7-72.6 109.7 0v96h182.9V292.5c.1-2.6-1.9-4.6-4.5-4.6m-288.1-4.5c0 2.6-2 4.6-4.6 4.6h-27.4c-2.6 0-4.6-2-4.6-4.6v-64c0-2.6 2-4.6 4.6-4.6h27.4c2.6 0 4.6 2 4.6 4.6zm146.4 0c0 2.6-2 4.6-4.6 4.6h-27.4c-2.6 0-4.6-2-4.6-4.6v-64c0-2.6 2-4.6 4.6-4.6h27.4c2.6 0 4.6 2 4.6 4.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FortAwesomeAlt(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M208 237.4h-22.2c-2.1 0-3.7 1.6-3.7 3.7v51.7c0 2.1 1.6 3.7 3.7 3.7H208c2.1 0 3.7-1.6 3.7-3.7v-51.7c0-2.1-1.6-3.7-3.7-3.7m118.2 0H304c-2.1 0-3.7 1.6-3.7 3.7v51.7c0 2.1 1.6 3.7 3.7 3.7h22.2c2.1 0 3.7-1.6 3.7-3.7v-51.7c-.1-2.1-1.7-3.7-3.7-3.7m132-125.1c-2.3-3.2-4.6-6.4-7.1-9.5c-9.8-12.5-20.8-24-32.8-34.4c-4.5-3.9-9.1-7.6-13.9-11.2c-1.6-1.2-3.2-2.3-4.8-3.5C372 34.1 340.3 20 306 13c-16.2-3.3-32.9-5-50-5s-33.9 1.7-50 5c-34.3 7.1-66 21.2-93.3 40.8c-1.6 1.1-3.2 2.3-4.8 3.5c-4.8 3.6-9.4 7.3-13.9 11.2c-3 2.6-5.9 5.3-8.8 8s-5.7 5.5-8.4 8.4c-5.5 5.7-10.7 11.8-15.6 18c-2.4 3.1-4.8 6.3-7.1 9.5C25.2 153 8.3 202.5 8.3 256c0 2 .1 4 .1 6c.1.7.1 1.3.1 2c.1 1.3.1 2.7.2 4c0 .8.1 1.5.1 2.3c0 1.3.1 2.5.2 3.7c.1.8.1 1.6.2 2.4c.1 1.1.2 2.3.3 3.5c0 .8.1 1.6.2 2.4c.1 1.2.3 2.4.4 3.6c.1.8.2 1.5.3 2.3c.1 1.3.3 2.6.5 3.9c.1.6.2 1.3.3 1.9l.9 5.7c.1.6.2 1.1.3 1.7c.3 1.3.5 2.7.8 4c.2.8.3 1.6.5 2.4c.2 1 .5 2.1.7 3.2c.2.9.4 1.7.6 2.6c.2 1 .4 2 .7 3c.2.9.5 1.8.7 2.7c.3 1 .5 1.9.8 2.9c.3.9.5 1.8.8 2.7c.2.9.5 1.9.8 2.8s.5 1.8.8 2.7c.3 1 .6 1.9.9 2.8c.6 1.6 1.1 3.3 1.7 4.9c.4 1 .7 1.9 1 2.8c.3 1 .7 2 1.1 3c.3.8.6 1.5.9 2.3l1.2 3c.3.7.6 1.5.9 2.2c.4 1 .9 2 1.3 3l.9 2.1c.5 1 .9 2 1.4 3c.3.7.6 1.3.9 2c.5 1 1 2.1 1.5 3.1c.2.6.5 1.1.8 1.7c.6 1.1 1.1 2.2 1.7 3.3c.1.2.2.3.3.5c2.2 4.1 4.4 8.2 6.8 12.2c.2.4.5.8.7 1.2c.7 1.1 1.3 2.2 2 3.3c.3.5.6.9.9 1.4c.6 1.1 1.3 2.1 2 3.2c.3.5.6.9.9 1.4c.7 1.1 1.4 2.1 2.1 3.2c.2.4.5.8.8 1.2c.7 1.1 1.5 2.2 2.3 3.3c.2.2.3.5.5.7c37.5 51.7 94.4 88.5 160 99.4c.9.1 1.7.3 2.6.4c1 .2 2.1.4 3.1.5s1.9.3 2.8.4c1 .2 2 .3 3 .4c.9.1 1.9.2 2.9.3s1.9.2 2.9.3s2.1.2 3.1.3c.9.1 1.8.1 2.7.2c1.1.1 2.3.1 3.4.2c.8 0 1.7.1 2.5.1c1.3 0 2.6.1 3.9.1c.7.1 1.4.1 2.1.1c2 .1 4 .1 6 .1s4-.1 6-.1c.7 0 1.4-.1 2.1-.1c1.3 0 2.6 0 3.9-.1c.8 0 1.7-.1 2.5-.1c1.1-.1 2.3-.1 3.4-.2c.9 0 1.8-.1 2.7-.2c1-.1 2.1-.2 3.1-.3s1.9-.2 2.9-.3c.9-.1 1.9-.2 2.9-.3s2-.3 3-.4s1.9-.3 2.8-.4c1-.2 2.1-.3 3.1-.5c.9-.1 1.7-.3 2.6-.4c65.6-11 122.5-47.7 160.1-102.4c.2-.2.3-.5.5-.7c.8-1.1 1.5-2.2 2.3-3.3c.2-.4.5-.8.8-1.2c.7-1.1 1.4-2.1 2.1-3.2c.3-.5.6-.9.9-1.4c.6-1.1 1.3-2.1 2-3.2c.3-.5.6-.9.9-1.4c.7-1.1 1.3-2.2 2-3.3c.2-.4.5-.8.7-1.2c2.4-4 4.6-8.1 6.8-12.2c.1-.2.2-.3.3-.5c.6-1.1 1.1-2.2 1.7-3.3c.2-.6.5-1.1.8-1.7c.5-1 1-2.1 1.5-3.1c.3-.7.6-1.3.9-2c.5-1 1-2 1.4-3l.9-2.1c.5-1 .9-2 1.3-3c.3-.7.6-1.5.9-2.2l1.2-3c.3-.8.6-1.5.9-2.3c.4-1 .7-2 1.1-3s.7-1.9 1-2.8c.6-1.6 1.2-3.3 1.7-4.9c.3-1 .6-1.9.9-2.8s.5-1.8.8-2.7c.2-.9.5-1.9.8-2.8s.6-1.8.8-2.7c.3-1 .5-1.9.8-2.9c.2-.9.5-1.8.7-2.7c.2-1 .5-2 .7-3c.2-.9.4-1.7.6-2.6c.2-1 .5-2.1.7-3.2c.2-.8.3-1.6.5-2.4c.3-1.3.6-2.7.8-4c.1-.6.2-1.1.3-1.7l.9-5.7c.1-.6.2-1.3.3-1.9c.1-1.3.3-2.6.5-3.9c.1-.8.2-1.5.3-2.3c.1-1.2.3-2.4.4-3.6c0-.8.1-1.6.2-2.4c.1-1.1.2-2.3.3-3.5c.1-.8.1-1.6.2-2.4c.1 1.7.1.5.2-.7c0-.8.1-1.5.1-2.3c.1-1.3.2-2.7.2-4c.1-.7.1-1.3.1-2c.1-2 .1-4 .1-6c0-53.5-16.9-103-45.8-143.7M448 371.5c-9.4 15.5-20.6 29.9-33.6 42.9c-20.6 20.6-44.5 36.7-71.2 48c-13.9 5.8-28.2 10.3-42.9 13.2v-75.8c0-58.6-88.6-58.6-88.6 0v75.8c-14.7-2.9-29-7.3-42.9-13.2c-26.7-11.3-50.6-27.4-71.2-48c-13-13-24.2-27.4-33.6-42.9v-71.3c0-2.1 1.6-3.7 3.7-3.7h22.1c2.1 0 3.7 1.6 3.7 3.7V326h29.6V182c0-2.1 1.6-3.7 3.7-3.7h22.1c2.1 0 3.7 1.6 3.7 3.7v25.9h29.5V182c0-2.1 1.6-3.7 3.7-3.7H208c2.1 0 3.7 1.6 3.7 3.7v25.9h29.5V182c0-4.8 6.5-3.7 9.5-3.7V88.1c-4.4-2-7.4-6.7-7.4-11.5c0-16.8 25.4-16.8 25.4 0c0 4.8-3 9.4-7.4 11.5V92c6.3-1.4 12.7-2.3 19.2-2.3c9.4 0 18.4 3.5 26.3 3.5c7.2 0 15.2-3.5 19.4-3.5c2.1 0 3.7 1.6 3.7 3.7v48.4c0 5.6-18.7 6.5-22.4 6.5c-8.6 0-16.6-3.5-25.4-3.5c-7 0-14.1 1.2-20.8 2.8v30.7c3 0 9.5-1.1 9.5 3.7v25.9h29.5V182c0-2.1 1.6-3.7 3.7-3.7h22.2c2.1 0 3.7 1.6 3.7 3.7v25.9h29.5V182c0-2.1 1.6-3.7 3.7-3.7h22.1c2.1 0 3.7 1.6 3.7 3.7v144h29.5v-25.8c0-2.1 1.6-3.7 3.7-3.7h22.2c2.1 0 3.7 1.6 3.7 3.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Forumbee(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.8 309.7C2 292.7 0 275.5 0 258.3C0 135 99.8 35 223.1 35c16.6 0 33.3 2 49.3 5.5C149 87.5 51.9 186 5.8 309.7m392.9-189.2C385 103 369 87.8 350.9 75.2c-149.6 44.3-266.3 162.1-309.7 312c12.5 18.1 28 35.6 45.2 49c43.1-151.3 161.2-271.7 312.3-315.7m15.8 252.7c15.2-25.1 25.4-53.7 29.5-82.8c-79.4 42.9-145 110.6-187.6 190.3c30-4.4 58.9-15.3 84.6-31.3c35 13.1 70.9 24.3 107 33.6c-9.3-36.5-20.4-74.5-33.5-109.8m29.7-145.5c-2.6-19.5-7.9-38.7-15.8-56.8C290.5 216.7 182 327.5 137.1 466c18.1 7.6 37 12.5 56.6 15.2C240 367.1 330.5 274.4 444.2 227.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Foursquare(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M323.1 3H49.9C12.4 3 0 31.3 0 49.1v433.8c0 20.3 12.1 27.7 18.2 30.1c6.2 2.5 22.8 4.6 32.9-7.1C180 356.5 182.2 354 182.2 354c3.1-3.4 3.4-3.1 6.8-3.1h83.4c35.1 0 40.6-25.2 44.3-39.7l48.6-243C373.8 25.8 363.1 3 323.1 3m-16.3 73.8l-11.4 59.7c-1.2 6.5-9.5 13.2-16.9 13.2H172.1c-12 0-20.6 8.3-20.6 20.3v13c0 12 8.6 20.6 20.6 20.6h90.4c8.3 0 16.6 9.2 14.8 18.2c-1.8 8.9-10.5 53.8-11.4 58.8c-.9 4.9-6.8 13.5-16.9 13.5h-73.5c-13.5 0-17.2 1.8-26.5 12.6c0 0-8.9 11.4-89.5 108.3c-.9.9-1.8.6-1.8-.3V75.9c0-7.7 6.8-16.6 16.6-16.6h219c8.2 0 15.6 7.7 13.5 17.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FreeCodeCamp(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M97.22 96.21c10.36-10.65 16-17.12 16-21.9c0-2.76-1.92-5.51-3.83-7.42a14.81 14.81 0 0 0-8.39-2.84c-8.48 0-20.92 8.79-35.84 25.69C23.68 137 2.51 182.81 3.37 250.34s17.47 117 54.06 161.87C76.22 435.86 90.62 448 100.9 448a13.55 13.55 0 0 0 8.37-3.84c1.91-2.76 3.81-5.63 3.81-8.38c0-5.63-3.86-12.2-13.2-20.55c-44.45-42.33-67.32-97-67.48-165c-.15-61.43 21.6-112.4 64.82-154.02m142.25 323.86c.58.37.91.55.91.55Zm93.79.55l.17-.13c-.19.13-.26.18-.17.13m3.13-158.18c-16.24-4.15 50.41-82.89-68.05-177.17c0 0 15.54 49.38-62.83 159.57c-74.27 104.35 23.46 168.73 34 175.23c-6.73-4.35-47.4-35.7 9.55-128.64c11-18.3 25.53-34.87 43.5-72.16c0 0 15.91 22.45 7.6 71.13C287.7 364 354 342.91 355 343.94c22.75 26.78-17.72 73.51-21.58 76.55c5.49-3.65 117.71-78 33-188.1c-5.99 6.01-13.8 34.2-30.03 30.05M510.88 89.69C496 72.79 483.52 64 475 64a14.81 14.81 0 0 0-8.39 2.84c-1.91 1.91-3.83 4.66-3.83 7.42c0 4.78 5.6 11.26 16 21.9c43.23 41.61 65 92.59 64.82 154.06c-.16 68-23 122.63-67.48 165c-9.34 8.35-13.18 14.92-13.2 20.55c0 2.75 1.9 5.62 3.81 8.38a13.61 13.61 0 0 0 8.37 3.85c10.28 0 24.68-12.13 43.47-35.79c36.59-44.85 53.14-94.38 54.06-161.87S552.32 137 510.88 89.69"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Freebsd(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M303.7 96.2c11.1-11.1 115.5-77 139.2-53.2c23.7 23.7-42.1 128.1-53.2 139.2c-11.1 11.1-39.4.9-63.1-22.9c-23.8-23.7-34.1-52-22.9-63.1M109.9 68.1C73.6 47.5 22 24.6 5.6 41.1c-16.6 16.6 7.1 69.4 27.9 105.7c18.5-32.2 44.8-59.3 76.4-78.7M406.7 174c3.3 11.3 2.7 20.7-2.7 26.1c-20.3 20.3-87.5-27-109.3-70.1c-18-32.3-11.1-53.4 14.9-48.7c5.7-3.6 12.3-7.6 19.6-11.6c-29.8-15.5-63.6-24.3-99.5-24.3c-119.1 0-215.6 96.5-215.6 215.6c0 119 96.5 215.6 215.6 215.6S445.3 380.1 445.3 261c0-38.4-10.1-74.5-27.7-105.8c-3.9 7-7.6 13.3-10.9 18.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fulcrum(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m95.75 164.14l-35.38 43.55L25 164.14l35.38-43.55zM144.23 0l-20.54 198.18L72.72 256l51 57.82L144.23 512V300.89L103.15 256l41.08-44.89zm79.67 164.14l35.38 43.55l35.38-43.55l-35.38-43.55zm-48.48 47L216.5 256l-41.08 44.89V512L196 313.82L247 256l-51-57.82L175.42 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GalacticRepublic(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M248 504C111.25 504 0 392.75 0 256S111.25 8 248 8s248 111.25 248 248s-111.25 248-248 248m0-479.47C120.37 24.53 16.53 128.37 16.53 256S120.37 487.47 248 487.47S479.47 383.63 479.47 256S375.63 24.53 248 24.53m27.62 21.81v24.62a185.933 185.933 0 0 1 83.57 34.54l17.39-17.36c-28.75-22.06-63.3-36.89-100.96-41.8m-55.37.07c-37.64 4.94-72.16 19.8-100.88 41.85l17.28 17.36h.08c24.07-17.84 52.55-30.06 83.52-34.67zm12.25 50.17v82.87c-10.04 2.03-19.42 5.94-27.67 11.42l-58.62-58.59l-21.93 21.93l58.67 58.67c-5.47 8.23-9.45 17.59-11.47 27.62h-82.9v31h82.9c2.02 10.02 6.01 19.31 11.47 27.54l-58.67 58.69l21.93 21.93l58.62-58.62a77.873 77.873 0 0 0 27.67 11.47v82.9h31v-82.9c10.05-2.03 19.37-6.06 27.62-11.55l58.67 58.69l21.93-21.93l-58.67-58.69c5.46-8.23 9.47-17.52 11.5-27.54h82.87v-31h-82.87c-2.02-10.02-6.03-19.38-11.5-27.62l58.67-58.67l-21.93-21.93l-58.67 58.67c-8.25-5.49-17.57-9.47-27.62-11.5V96.58zm183.24 30.72l-17.36 17.36a186.337 186.337 0 0 1 34.67 83.67h24.62c-4.95-37.69-19.83-72.29-41.93-101.03m-335.55.13c-22.06 28.72-36.91 63.26-41.85 100.91h24.65c4.6-30.96 16.76-59.45 34.59-83.52zM38.34 283.67c4.92 37.64 19.75 72.18 41.8 100.9l17.36-17.39c-17.81-24.07-29.92-52.57-34.51-83.52H38.34zm394.7 0c-4.61 30.99-16.8 59.5-34.67 83.6l17.36 17.36c22.08-28.74 36.98-63.29 41.93-100.96zM136.66 406.38l-17.36 17.36c28.73 22.09 63.3 36.98 100.96 41.93v-24.64c-30.99-4.63-59.53-16.79-83.6-34.65m222.53.05c-24.09 17.84-52.58 30.08-83.57 34.67v24.57c37.67-4.92 72.21-19.79 100.96-41.85l-17.31-17.39z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GalacticSenate(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M249.86 33.48v26.07C236.28 80.17 226 168.14 225.39 274.9c11.74-15.62 19.13-33.33 19.13-48.24v-16.88c-.03-5.32.75-10.53 2.19-15.65c.65-2.14 1.39-4.08 2.62-5.82c1.23-1.75 3.43-3.79 6.68-3.79c3.24 0 5.45 2.05 6.68 3.79c1.23 1.75 1.97 3.68 2.62 5.82c1.44 5.12 2.22 10.33 2.19 15.65v16.88c0 14.91 7.39 32.62 19.13 48.24c-.63-106.76-10.91-194.73-24.49-215.35V33.48zm-26.34 147.77c-9.52 2.15-18.7 5.19-27.46 9.08c8.9 16.12 9.76 32.64 1.71 37.29c-8 4.62-21.85-4.23-31.36-19.82c-11.58 8.79-21.88 19.32-30.56 31.09c14.73 9.62 22.89 22.92 18.32 30.66c-4.54 7.7-20.03 7.14-35.47-.96c-5.78 13.25-9.75 27.51-11.65 42.42c9.68.18 18.67 2.38 26.18 6.04c17.78-.3 32.77-1.96 40.49-4.22c5.55-26.35 23.02-48.23 46.32-59.51c.73-25.55 1.88-49.67 3.48-72.07m64.96 0c1.59 22.4 2.75 46.52 3.47 72.07c23.29 11.28 40.77 33.16 46.32 59.51c7.72 2.26 22.71 3.92 40.49 4.22c7.51-3.66 16.5-5.85 26.18-6.04c-1.9-14.91-5.86-29.17-11.65-42.42c-15.44 8.1-30.93 8.66-35.47.96c-4.57-7.74 3.6-21.05 18.32-30.66c-8.68-11.77-18.98-22.3-30.56-31.09c-9.51 15.59-23.36 24.44-31.36 19.82c-8.05-4.65-7.19-21.16 1.71-37.29a147.49 147.49 0 0 0-27.45-9.08m-32.48 8.6c-3.23 0-5.86 8.81-6.09 19.93h-.05v16.88c0 41.42-49.01 95.04-93.49 95.04c-52 0-122.75-1.45-156.37 29.17v2.51c9.42 17.12 20.58 33.17 33.18 47.97C45.7 380.26 84.77 360.4 141.2 360c45.68 1.02 79.03 20.33 90.76 40.87c.01.01-.01.04 0 .05c7.67 2.14 15.85 3.23 24.04 3.21c8.19.02 16.37-1.07 24.04-3.21c.01-.01-.01-.04 0-.05c11.74-20.54 45.08-39.85 90.76-40.87c56.43.39 95.49 20.26 108.02 41.35c12.6-14.8 23.76-30.86 33.18-47.97v-2.51c-33.61-30.62-104.37-29.17-156.37-29.17c-44.48 0-93.49-53.62-93.49-95.04v-16.88h-.05c-.23-11.12-2.86-19.93-6.09-19.93m0 96.59c22.42 0 40.6 18.18 40.6 40.6s-18.18 40.65-40.6 40.65s-40.6-18.23-40.6-40.65c0-22.42 18.18-40.6 40.6-40.6m0 7.64c-18.19 0-32.96 14.77-32.96 32.96S237.81 360 256 360s32.96-14.77 32.96-32.96s-14.77-32.96-32.96-32.96m0 6.14c14.81 0 26.82 12.01 26.82 26.82s-12.01 26.82-26.82 26.82s-26.82-12.01-26.82-26.82s12.01-26.82 26.82-26.82m-114.8 66.67c-10.19.07-21.6.36-30.5 1.66c.43 4.42 1.51 18.63 7.11 29.76c9.11-2.56 18.36-3.9 27.62-3.9c41.28.94 71.48 34.35 78.26 74.47l.11 4.7c10.4 1.91 21.19 2.94 32.21 2.94c11.03 0 21.81-1.02 32.21-2.94l.11-4.7c6.78-40.12 36.98-73.53 78.26-74.47c9.26 0 18.51 1.34 27.62 3.9c5.6-11.13 6.68-25.34 7.11-29.76c-8.9-1.3-20.32-1.58-30.5-1.66c-18.76.42-35.19 4.17-48.61 9.67c-12.54 16.03-29.16 30.03-49.58 33.07c-.09.02-.17.04-.27.05c-.05.01-.11.04-.16.05c-5.24 1.07-10.63 1.6-16.19 1.6c-5.55 0-10.95-.53-16.19-1.6c-.05-.01-.11-.04-.16-.05c-.1-.02-.17-.04-.27-.05c-20.42-3.03-37.03-17.04-49.58-33.07c-13.42-5.49-29.86-9.25-48.61-9.67"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GetPocket(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M407.6 64h-367C18.5 64 0 82.5 0 104.6v135.2C0 364.5 99.7 464 224.2 464c124 0 223.8-99.5 223.8-224.2V104.6c0-22.4-17.7-40.6-40.4-40.6m-162 268.5c-12.4 11.8-31.4 11.1-42.4 0C89.5 223.6 88.3 227.4 88.3 209.3c0-16.9 13.8-30.7 30.7-30.7c17 0 16.1 3.8 105.2 89.3c90.6-86.9 88.6-89.3 105.5-89.3c16.9 0 30.7 13.8 30.7 30.7c0 17.8-2.9 15.7-114.8 123.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gg(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m179.2 230.4l102.4 102.4l-102.4 102.4L0 256L179.2 76.8l44.8 44.8l-25.6 25.6l-19.2-19.2l-128 128l128 128l51.5-51.5l-77.1-76.5zM332.8 76.8L230.4 179.2l102.4 102.4l25.6-25.6l-77.1-76.5l51.5-51.5l128 128l-128 128l-19.2-19.2l-25.6 25.6l44.8 44.8L512 256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GgCircle(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M257 8C120 8 9 119 9 256s111 248 248 248s248-111 248-248S394 8 257 8m-49.5 374.8L81.8 257.1l125.7-125.7l35.2 35.4l-24.2 24.2l-11.1-11.1l-77.2 77.2l77.2 77.2l26.6-26.6l-53.1-52.9l24.4-24.4l77.2 77.2zm99-2.2l-35.2-35.2l24.1-24.4l11.1 11.1l77.2-77.2l-77.2-77.2l-26.5 26.5l53.1 52.9l-24.4 24.4l-77.2-77.2l75-75L432.2 255z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Git(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M216.29 158.39H137C97 147.9 6.51 150.63 6.51 233.18c0 30.09 15 51.23 35 61c-25.1 23-37 33.85-37 49.21c0 11 4.47 21.14 17.89 26.81C8.13 383.61 0 393.35 0 411.65c0 32.11 28.05 50.82 101.63 50.82c70.75 0 111.79-26.42 111.79-73.18c0-58.66-45.16-56.5-151.63-63l13.43-21.55c27.27 7.58 118.7 10 118.7-67.89c0-18.7-7.73-31.71-15-41.07l37.41-2.84zm-63.42 241.9c0 32.06-104.89 32.1-104.89 2.43c0-8.14 5.27-15 10.57-21.54c77.71 5.3 94.32 3.37 94.32 19.11m-50.81-134.58c-52.8 0-50.46-71.16 1.2-71.16c49.54 0 50.82 71.16-1.2 71.16m133.3 100.51v-32.1c26.75-3.66 27.24-2 27.24-11V203.61c0-8.5-2.05-7.38-27.24-16.26l4.47-32.92H324v168.71c0 6.51.4 7.32 6.51 8.14l20.73 2.84v32.1zm52.45-244.31c-23.17 0-36.59-13.43-36.59-36.61s13.42-35.77 36.59-35.77c23.58 0 37 12.62 37 35.77s-13.42 36.61-37 36.61M512 350.46c-17.49 8.53-43.1 16.26-66.28 16.26c-48.38 0-66.67-19.5-66.67-65.46V194.75c0-5.42 1.05-4.06-31.71-4.06V154.5c35.78-4.07 50-22 54.47-66.27h38.63c0 65.83-1.34 61.81 3.26 61.81H501v40.65h-60.56v97.15c0 6.92-4.92 51.41 60.57 26.84z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitAlt(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M439.55 236.05L244 40.45a28.87 28.87 0 0 0-40.81 0l-40.66 40.63l51.52 51.52c27.06-9.14 52.68 16.77 43.39 43.68l49.66 49.66c34.23-11.8 61.18 31 35.47 56.69c-26.49 26.49-70.21-2.87-56-37.34L240.22 199v121.85c25.3 12.54 22.26 41.85 9.08 55a34.34 34.34 0 0 1-48.55 0c-17.57-17.6-11.07-46.91 11.25-56v-123c-20.8-8.51-24.6-30.74-18.64-45L142.57 101L8.45 235.14a28.86 28.86 0 0 0 0 40.81l195.61 195.6a28.86 28.86 0 0 0 40.8 0l194.69-194.69a28.86 28.86 0 0 0 0-40.81"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitSquare(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M100.59 334.24c48.57 3.31 58.95 2.11 58.95 11.94c0 20-65.55 20.06-65.55 1.52c.01-5.09 3.29-9.4 6.6-13.46m27.95-116.64c-32.29 0-33.75 44.47-.75 44.47c32.51 0 31.71-44.47.75-44.47M448 80v352a48 48 0 0 1-48 48H48a48 48 0 0 1-48-48V80a48 48 0 0 1 48-48h352a48 48 0 0 1 48 48m-227 69.31c0 14.49 8.38 22.88 22.86 22.88c14.74 0 23.13-8.39 23.13-22.88S258.62 127 243.88 127c-14.48 0-22.88 7.84-22.88 22.31M199.18 195h-49.55c-25-6.55-81.56-4.85-81.56 46.75c0 18.8 9.4 32 21.85 38.11C74.23 294.23 66.8 301 66.8 310.6c0 6.87 2.79 13.22 11.18 16.76c-8.9 8.4-14 14.48-14 25.92C64 373.35 81.53 385 127.52 385c44.22 0 69.87-16.51 69.87-45.73c0-36.67-28.23-35.32-94.77-39.38l8.38-13.43c17 4.74 74.19 6.23 74.19-42.43c0-11.69-4.83-19.82-9.4-25.67l23.38-1.78zm84.34 109.84l-13-1.78c-3.82-.51-4.07-1-4.07-5.09V192.52h-52.6l-2.79 20.57c15.75 5.55 17 4.86 17 10.17V298c0 5.62-.31 4.58-17 6.87v20.06h72.42zM384 315l-6.87-22.37c-40.93 15.37-37.85-12.41-37.85-16.73v-60.72h37.85v-25.41h-35.82c-2.87 0-2 2.52-2-38.63h-24.18c-2.79 27.7-11.68 38.88-34 41.42v22.62c20.47 0 19.82-.85 19.82 2.54v66.57c0 28.72 11.43 40.91 41.67 40.91c14.45 0 30.45-4.83 41.38-10.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Github(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M165.9 397.4c0 2-2.3 3.6-5.2 3.6c-3.3.3-5.6-1.3-5.6-3.6c0-2 2.3-3.6 5.2-3.6c3-.3 5.6 1.3 5.6 3.6m-31.1-4.5c-.7 2 1.3 4.3 4.3 4.9c2.6 1 5.6 0 6.2-2s-1.3-4.3-4.3-5.2c-2.6-.7-5.5.3-6.2 2.3m44.2-1.7c-2.9.7-4.9 2.6-4.6 4.9c.3 2 2.9 3.3 5.9 2.6c2.9-.7 4.9-2.6 4.6-4.6c-.3-1.9-3-3.2-5.9-2.9M244.8 8C106.1 8 0 113.3 0 252c0 110.9 69.8 205.8 169.5 239.2c12.8 2.3 17.3-5.6 17.3-12.1c0-6.2-.3-40.4-.3-61.4c0 0-70 15-84.7-29.8c0 0-11.4-29.1-27.8-36.6c0 0-22.9-15.7 1.6-15.4c0 0 24.9 2 38.6 25.8c21.9 38.6 58.6 27.5 72.9 20.9c2.3-16 8.8-27.1 16-33.7c-55.9-6.2-112.3-14.3-112.3-110.5c0-27.5 7.6-41.3 23.6-58.9c-2.6-6.5-11.1-33.3 2.6-67.9c20.9-6.5 69 27 69 27c20-5.6 41.5-8.5 62.8-8.5s42.8 2.9 62.8 8.5c0 0 48.1-33.6 69-27c13.7 34.7 5.2 61.4 2.6 67.9c16 17.7 25.8 31.5 25.8 58.9c0 96.5-58.9 104.2-114.8 110.5c9.2 7.9 17 22.9 17 46.4c0 33.7-.3 75.4-.3 83.6c0 6.5 4.6 14.4 17.3 12.1C428.2 457.8 496 362.9 496 252C496 113.3 383.5 8 244.8 8M97.2 352.9c-1.3 1-1 3.3.7 5.2c1.6 1.6 3.9 2.3 5.2 1c1.3-1 1-3.3-.7-5.2c-1.6-1.6-3.9-2.3-5.2-1m-10.8-8.1c-.7 1.3.3 2.9 2.3 3.9c1.6 1 3.6.7 4.3-.7c.7-1.3-.3-2.9-2.3-3.9c-2-.6-3.6-.3-4.3.7m32.4 35.6c-1.6 1.3-1 4.3 1.3 6.2c2.3 2.3 5.2 2.6 6.5 1c1.3-1.3.7-4.3-1.3-6.2c-2.2-2.3-5.2-2.6-6.5-1m-11.4-14.7c-1.6 1-1.6 3.6 0 5.9c1.6 2.3 4.3 3.3 5.6 2.3c1.6-1.3 1.6-3.9 0-6.2c-1.4-2.3-4-3.3-5.6-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GithubAlt(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M186.1 328.7c0 20.9-10.9 55.1-36.7 55.1s-36.7-34.2-36.7-55.1s10.9-55.1 36.7-55.1s36.7 34.2 36.7 55.1M480 278.2c0 31.9-3.2 65.7-17.5 95c-37.9 76.6-142.1 74.8-216.7 74.8c-75.8 0-186.2 2.7-225.6-74.8c-14.6-29-20.2-63.1-20.2-95c0-41.9 13.9-81.5 41.5-113.6c-5.2-15.8-7.7-32.4-7.7-48.8c0-21.5 4.9-32.3 14.6-51.8c45.3 0 74.3 9 108.8 36c29-6.9 58.8-10 88.7-10c27 0 54.2 2.9 80.4 9.2c34-26.7 63-35.2 107.8-35.2c9.8 19.5 14.6 30.3 14.6 51.8c0 16.4-2.6 32.7-7.7 48.2c27.5 32.4 39 72.3 39 114.2m-64.3 50.5c0-43.9-26.7-82.6-73.5-82.6c-18.9 0-37 3.4-56 6c-14.9 2.3-29.8 3.2-45.1 3.2c-15.2 0-30.1-.9-45.1-3.2c-18.7-2.6-37-6-56-6c-46.8 0-73.5 38.7-73.5 82.6c0 87.8 80.4 101.3 150.4 101.3h48.2c70.3 0 150.6-13.4 150.6-101.3m-82.6-55.1c-25.8 0-36.7 34.2-36.7 55.1s10.9 55.1 36.7 55.1s36.7-34.2 36.7-55.1s-10.9-55.1-36.7-55.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GithubSquare(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 32H48C21.5 32 0 53.5 0 80v352c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V80c0-26.5-21.5-48-48-48M277.3 415.7c-8.4 1.5-11.5-3.7-11.5-8c0-5.4.2-33 .2-55.3c0-15.6-5.2-25.5-11.3-30.7c37-4.1 76-9.2 76-73.1c0-18.2-6.5-27.3-17.1-39c1.7-4.3 7.4-22-1.7-45c-13.9-4.3-45.7 17.9-45.7 17.9c-13.2-3.7-27.5-5.6-41.6-5.6c-14.1 0-28.4 1.9-41.6 5.6c0 0-31.8-22.2-45.7-17.9c-9.1 22.9-3.5 40.6-1.7 45c-10.6 11.7-15.6 20.8-15.6 39c0 63.6 37.3 69 74.3 73.1c-4.8 4.3-9.1 11.7-10.6 22.3c-9.5 4.3-33.8 11.7-48.3-13.9c-9.1-15.8-25.5-17.1-25.5-17.1c-16.2-.2-1.1 10.2-1.1 10.2c10.8 5 18.4 24.2 18.4 24.2c9.7 29.7 56.1 19.7 56.1 19.7c0 13.9.2 36.5.2 40.6c0 4.3-3 9.5-11.5 8c-66-22.1-112.2-84.9-112.2-158.3c0-91.8 70.2-161.5 162-161.5S388 165.6 388 257.4c.1 73.4-44.7 136.3-110.7 158.3m-98.1-61.1c-1.9.4-3.7-.4-3.9-1.7c-.2-1.5 1.1-2.8 3-3.2c1.9-.2 3.7.6 3.9 1.9c.3 1.3-1 2.6-3 3m-9.5-.9c0 1.3-1.5 2.4-3.5 2.4c-2.2.2-3.7-.9-3.7-2.4c0-1.3 1.5-2.4 3.5-2.4c1.9-.2 3.7.9 3.7 2.4m-13.7-1.1c-.4 1.3-2.4 1.9-4.1 1.3c-1.9-.4-3.2-1.9-2.8-3.2c.4-1.3 2.4-1.9 4.1-1.5c2 .6 3.3 2.1 2.8 3.4m-12.3-5.4c-.9 1.1-2.8.9-4.3-.6c-1.5-1.3-1.9-3.2-.9-4.1c.9-1.1 2.8-.9 4.3.6c1.3 1.3 1.8 3.3.9 4.1m-9.1-9.1c-.9.6-2.6 0-3.7-1.5s-1.1-3.2 0-3.9c1.1-.9 2.8-.2 3.7 1.3c1.1 1.5 1.1 3.3 0 4.1m-6.5-9.7c-.9.9-2.4.4-3.5-.6c-1.1-1.3-1.3-2.8-.4-3.5c.9-.9 2.4-.4 3.5.6c1.1 1.3 1.3 2.8.4 3.5m-6.7-7.4c-.4.9-1.7 1.1-2.8.4c-1.3-.6-1.9-1.7-1.5-2.6c.4-.6 1.5-.9 2.8-.4c1.3.7 1.9 1.8 1.5 2.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gitkraken(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M565.7 118.1c-2.3-6.1-9.3-9.2-15.3-6.6c-5.7 2.4-8.5 8.9-6.3 14.6c10.9 29 16.9 60.5 16.9 93.3c0 134.6-100.3 245.7-230.2 262.7V358.4c7.9-1.5 15.5-3.6 23-6.2v104c106.7-25.9 185.9-122.1 185.9-236.8c0-91.8-50.8-171.8-125.8-213.3c-5.7-3.2-13-.9-15.9 5c-2.7 5.5-.6 12.2 4.7 15.1c67.9 37.6 113.9 110 113.9 193.2c0 93.3-57.9 173.1-139.8 205.4v-92.2c14.2-4.5 24.9-17.7 24.9-33.5c0-13.1-6.8-24.4-17.3-30.5c8.3-79.5 44.5-58.6 44.5-83.9V170c0-38-87.9-161.8-129-164.7c-2.5-.2-5-.2-7.6 0C251.1 8.3 163.2 132 163.2 170v14.8c0 25.3 36.3 4.3 44.5 83.9c-10.6 6.1-17.3 17.4-17.3 30.5c0 15.8 10.6 29 24.8 33.5v92.2c-81.9-32.2-139.8-112-139.8-205.4c0-83.1 46-155.5 113.9-193.2c5.4-3 7.4-9.6 4.7-15.1c-2.9-5.9-10.1-8.2-15.9-5c-75 41.5-125.8 121.5-125.8 213.3c0 114.7 79.2 210.8 185.9 236.8v-104c7.6 2.5 15.1 4.6 23 6.2v123.7C131.4 465.2 31 354.1 31 219.5c0-32.8 6-64.3 16.9-93.3c2.2-5.8-.6-12.2-6.3-14.6c-6-2.6-13 .4-15.3 6.6C14.5 149.7 8 183.8 8 219.5c0 155.1 122.6 281.6 276.3 287.8V361.4c6.8.4 15 .5 23.4 0v145.8C461.4 501.1 584 374.6 584 219.5c0-35.7-6.5-69.8-18.3-101.4M365.9 275.5c13 0 23.7 10.5 23.7 23.7c0 13.1-10.6 23.7-23.7 23.7c-13 0-23.7-10.5-23.7-23.7c0-13.1 10.6-23.7 23.7-23.7m-139.8 47.3c-13.2 0-23.7-10.7-23.7-23.7s10.5-23.7 23.7-23.7c13.1 0 23.7 10.6 23.7 23.7c0 13-10.5 23.7-23.7 23.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gitlab(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M105.2 24.9c-3.1-8.9-15.7-8.9-18.9 0L29.8 199.7h132c-.1 0-56.6-174.8-56.6-174.8M.9 287.7c-2.6 8 .3 16.9 7.1 22l247.9 184l-226.2-294zm160.8-88l94.3 294l94.3-294zm349.4 88l-28.8-88l-226.3 294l247.9-184c6.9-5.1 9.7-14 7.2-22M425.7 24.9c-3.1-8.9-15.7-8.9-18.9 0l-56.6 174.8h132z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gitter(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M66.4 322.5H16V0h50.4zM166.9 76.1h-50.4V512h50.4zm100.6 0h-50.4V512h50.4zM368 76h-50.4v247H368z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Glide(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M252.8 148.6c0 8.8-1.6 17.7-3.4 26.4c-5.8 27.8-11.6 55.8-17.3 83.6c-1.4 6.3-8.3 4.9-13.7 4.9c-23.8 0-30.5-26-30.5-45.5c0-29.3 11.2-68.1 38.5-83.1c4.3-2.5 9.2-4.2 14.1-4.2c11.4 0 12.3 8.3 12.3 17.9M448 80v352c0 26.5-21.5 48-48 48H48c-26.5 0-48-21.5-48-48V80c0-26.5 21.5-48 48-48h352c26.5 0 48 21.5 48 48m-64 187c0-5.1-20.8-37.7-25.5-39.5c-2.2-.9-7.2-2.3-9.6-2.3c-23.1 0-38.7 10.5-58.2 21.5l-.5-.5c4.3-29.4 14.6-57.2 14.6-87.4c0-44.6-23.8-62.7-67.5-62.7c-71.7 0-108 70.8-108 123.5c0 54.7 32 85 86.3 85c7.5 0 6.9-.6 6.9 2.3c-10.5 80.3-56.5 82.9-56.5 58.9c0-24.4 28-36.5 28.3-38c-.2-7.6-29.3-17.2-36.7-17.2c-21.1 0-32.7 33-32.7 50.6c0 32.3 20.4 54.7 53.3 54.7c48.2 0 83.4-49.7 94.3-91.7c9.4-37.7 7-39.4 12.3-42.1c20-10.1 35.8-16.8 58.4-16.8c11.1 0 19 2.3 36.7 5.2c1.8.1 4.1-1.7 4.1-3.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlideG(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M407.1 211.2c-3.5-1.4-11.6-3.8-15.4-3.8c-37.1 0-62.2 16.8-93.5 34.5l-.9-.9c7-47.3 23.5-91.9 23.5-140.4C320.8 29.1 282.6 0 212.4 0C97.3 0 39 113.7 39 198.4C39 286.3 90.3 335 177.6 335c12 0 11-1 11 3.8c-16.9 128.9-90.8 133.1-90.8 94.6c0-39.2 45-58.6 45.5-61c-.3-12.2-47-27.6-58.9-27.6c-33.9.1-52.4 51.2-52.4 79.3C32 476 64.8 512 117.5 512c77.4 0 134-77.8 151.4-145.4c15.1-60.5 11.2-63.3 19.7-67.6c32.2-16.2 57.5-27 93.8-27c17.8 0 30.5 3.7 58.9 8.4c2.9 0 6.7-2.9 6.7-5.8c0-8-33.4-60.5-40.9-63.4m-175.3-84.4c-9.3 44.7-18.6 89.6-27.8 134.3c-2.3 10.2-13.3 7.8-22 7.8c-38.3 0-49-41.8-49-73.1c0-47 18-109.3 61.8-133.4c7-4.1 14.8-6.7 22.6-6.7c18.6 0 20 13.3 20 28.7c-.1 14.3-2.7 28.5-5.6 42.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gofore(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M324 319.8h-13.2v34.7c-24.5 23.1-56.3 35.8-89.9 35.8c-73.2 0-132.4-60.2-132.4-134.4c0-74.1 59.2-134.4 132.4-134.4c35.3 0 68.6 14 93.6 39.4l62.3-63.3C335 55.3 279.7 32 220.7 32C98 32 0 132.6 0 256c0 122.5 97 224 220.7 224c63.2 0 124.5-26.2 171-82.5c-2-27.6-13.4-77.7-67.7-77.7m-12.1-112.5H205.6v89H324c33.5 0 60.5 15.1 76 41.8v-30.6c0-65.2-40.4-100.2-88.1-100.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Goodreads(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299.9 191.2c5.1 37.3-4.7 79-35.9 100.7c-22.3 15.5-52.8 14.1-70.8 5.7c-37.1-17.3-49.5-58.6-46.8-97.2c4.3-60.9 40.9-87.9 75.3-87.5c46.9-.2 71.8 31.8 78.2 78.3M448 88v336c0 30.9-25.1 56-56 56H56c-30.9 0-56-25.1-56-56V88c0-30.9 25.1-56 56-56h336c30.9 0 56 25.1 56 56M330 313.2s-.1-34-.1-217.3h-29v40.3c-.8.3-1.2-.5-1.6-1.2c-9.6-20.7-35.9-46.3-76-46c-51.9.4-87.2 31.2-100.6 77.8c-4.3 14.9-5.8 30.1-5.5 45.6c1.7 77.9 45.1 117.8 112.4 115.2c28.9-1.1 54.5-17 69-45.2c.5-1 1.1-1.9 1.7-2.9c.2.1.4.1.6.2c.3 3.8.2 30.7.1 34.5c-.2 14.8-2 29.5-7.2 43.5c-7.8 21-22.3 34.7-44.5 39.5c-17.8 3.9-35.6 3.8-53.2-1.2c-21.5-6.1-36.5-19-41.1-41.8c-.3-1.6-1.3-1.3-2.3-1.3h-26.8c.8 10.6 3.2 20.3 8.5 29.2c24.2 40.5 82.7 48.5 128.2 37.4c49.9-12.3 67.3-54.9 67.4-106.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoodreadsG(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M42.6 403.3h2.8c12.7 0 25.5 0 38.2.1c1.6 0 3.1-.4 3.6 2.1c7.1 34.9 30 54.6 62.9 63.9c26.9 7.6 54.1 7.8 81.3 1.8c33.8-7.4 56-28.3 68-60.4c8-21.5 10.7-43.8 11-66.5c.1-5.8.3-47-.2-52.8l-.9-.3c-.8 1.5-1.7 2.9-2.5 4.4c-22.1 43.1-61.3 67.4-105.4 69.1c-103 4-169.4-57-172-176.2c-.5-23.7 1.8-46.9 8.3-69.7C58.3 47.7 112.3.6 191.6 0c61.3-.4 101.5 38.7 116.2 70.3c.5 1.1 1.3 2.3 2.4 1.9V10.6h44.3c0 280.3.1 332.2.1 332.2c-.1 78.5-26.7 143.7-103 162.2c-69.5 16.9-159 4.8-196-57.2c-8-13.5-11.8-28.3-13-44.5M188.9 36.5c-52.5-.5-108.5 40.7-115 133.8c-4.1 59 14.8 122.2 71.5 148.6c27.6 12.9 74.3 15 108.3-8.7c47.6-33.2 62.7-97 54.8-154c-9.7-71.1-47.8-120-119.6-119.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Google(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M488 261.8C488 403.3 391.1 504 248 504C110.8 504 0 393.2 0 256S110.8 8 248 8c66.8 0 123 24.5 166.3 64.9l-67.5 64.9C258.5 52.6 94.3 116.6 94.3 256c0 86.5 69.1 156.6 153.7 156.6c98.2 0 135-70.4 140.8-106.9H248v-85.3h236.1c2.3 12.7 3.9 24.9 3.9 41.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleDrive(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M339 314.9L175.4 32h161.2l163.6 282.9zm-137.5 23.6L120.9 480h310.5L512 338.5zM154.1 67.4L0 338.5L80.6 480L237 208.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePay(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M105.72 215v41.25h57.1a49.66 49.66 0 0 1-21.14 32.6c-9.54 6.55-21.72 10.28-36 10.28c-27.6 0-50.93-18.91-59.3-44.22a65.61 65.61 0 0 1 0-41c8.37-25.46 31.7-44.37 59.3-44.37a56.43 56.43 0 0 1 40.51 16.08L176.47 155a101.24 101.24 0 0 0-70.75-27.84a105.55 105.55 0 0 0-94.38 59.11a107.64 107.64 0 0 0 0 96.18v.15a105.41 105.41 0 0 0 94.38 59c28.47 0 52.55-9.53 70-25.91c20-18.61 31.41-46.15 31.41-78.91a133.76 133.76 0 0 0-1.75-21.78Zm389.41-4c-10.13-9.38-23.93-14.14-41.39-14.14c-22.46 0-39.34 8.34-50.5 24.86l20.85 13.26q11.45-17 31.26-17a34.05 34.05 0 0 1 22.75 8.79a28.14 28.14 0 0 1 9.69 21.23v5.51c-9.1-5.07-20.55-7.75-34.64-7.75c-16.44 0-29.65 3.88-39.49 11.77s-14.82 18.31-14.82 31.56a39.74 39.74 0 0 0 13.94 31.27c9.25 8.34 21 12.51 34.79 12.51c16.29 0 29.21-7.3 39-21.89h1v17.72h22.61V250c.07-16.55-4.92-29.66-15.05-39m-19.23 89.3a37.32 37.32 0 0 1-26.57 11.16a28.61 28.61 0 0 1-18.33-6.25a19.41 19.41 0 0 1-7.77-15.63c0-7 3.22-12.81 9.54-17.42s14.53-7 24.07-7c13.16-.16 23.46 2.84 30.8 8.78c0 10.13-3.96 18.91-11.74 26.36m-93.65-142a55.71 55.71 0 0 0-40.51-16.3h-62.67v186.74h23.63V253.1h39c16 0 29.5-5.36 40.51-15.93c.88-.89 1.76-1.79 2.65-2.68a54.45 54.45 0 0 0-2.61-76.23Zm-16.58 62.23a30.65 30.65 0 0 1-23.34 9.68H302.7V165h39.63a32 32 0 0 1 22.6 9.23a33.18 33.18 0 0 1 .74 46.26ZM614.31 201l-36.54 91.7h-.45L539.9 201h-25.69L566 320.55l-29.35 64.32H561L640 201Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePlay(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M325.3 234.3L104.6 13l280.8 161.2zM47 0C34 6.8 25.3 19.2 25.3 35.3v441.3c0 16.1 8.7 28.5 21.7 35.3l256.6-256zm425.2 225.6l-58.9-34.1l-65.7 64.5l65.7 64.5l60.1-34.1c18-14.3 18-46.5-1.2-60.8M104.6 499l280.8-161.2l-60.1-60.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePlus(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 8C119.1 8 8 119.1 8 256s111.1 248 248 248s248-111.1 248-248S392.9 8 256 8m-70.7 372a124 124 0 0 1 0-248c31.3 0 60.1 11 83 32.3l-33.6 32.6c-13.2-12.9-31.3-19.1-49.4-19.1c-42.9 0-77.2 35.5-77.2 78.1s34.2 78.1 77.2 78.1c32.6 0 64.9-19.1 70.1-53.3h-70.1v-42.6h116.9a109.2 109.2 0 0 1 1.9 20.7c0 70.8-47.5 121.2-118.8 121.2m230.2-106.2v35.5H380v-35.5h-35.5v-35.5H380v-35.5h35.5v35.5h35.2v35.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePlusG(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M386.061 228.496c1.834 9.692 3.143 19.384 3.143 31.956C389.204 370.205 315.599 448 204.8 448c-106.084 0-192-85.915-192-192s85.916-192 192-192c51.864 0 95.083 18.859 128.611 50.292l-52.126 50.03c-14.145-13.621-39.028-29.599-76.485-29.599c-65.484 0-118.92 54.221-118.92 121.277c0 67.056 53.436 121.277 118.92 121.277c75.961 0 104.513-54.745 108.965-82.773H204.8v-66.009h181.261zm185.406 6.437V179.2h-56.001v55.733h-55.733v56.001h55.733v55.733h56.001v-55.733H627.2v-56.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePlusSquare(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 32H48C21.5 32 0 53.5 0 80v352c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V80c0-26.5-21.5-48-48-48M164 356c-55.3 0-100-44.7-100-100s44.7-100 100-100c27 0 49.5 9.8 67 26.2l-27.1 26.1c-7.4-7.1-20.3-15.4-39.8-15.4c-34.1 0-61.9 28.2-61.9 63.2c0 34.9 27.8 63.2 61.9 63.2c39.6 0 54.4-28.5 56.8-43.1H164v-34.4h94.4c1 5 1.6 10.1 1.6 16.6c0 57.1-38.3 97.6-96 97.6m220-81.8h-29v29h-29.2v-29h-29V245h29v-29H355v29h29z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleWallet(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M156.8 126.8c37.6 60.6 64.2 113.1 84.3 162.5c-8.3 33.8-18.8 66.5-31.3 98.3c-13.2-52.3-26.5-101.3-56-148.5c6.5-36.4 2.3-73.6 3-112.3M109.3 200H16.1c-6.5 0-10.5 7.5-6.5 12.7C51.8 267 81.3 330.5 101.3 400h103.5c-16.2-69.7-38.7-133.7-82.5-193.5c-3-4-8-6.5-13-6.5m47.8-88c68.5 108 130 234.5 138.2 368H409c-12-138-68.4-265-143.2-368zm251.8-68.5c-1.8-6.8-8.2-11.5-15.2-11.5h-88.3c-5.3 0-9 5-7.8 10.3c13.2 46.5 22.3 95.5 26.5 146c48.2 86.2 79.7 178.3 90.6 270.8c15.8-60.5 25.3-133.5 25.3-203c0-73.6-12.1-145.1-31.1-212.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gratipay(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M248 8C111.1 8 0 119.1 0 256s111.1 248 248 248s248-111.1 248-248S384.9 8 248 8m114.6 226.4l-113 152.7l-112.7-152.7c-8.7-11.9-19.1-50.4 13.6-72c28.1-18.1 54.6-4.2 68.5 11.9c15.9 17.9 46.6 16.9 61.7 0c13.9-16.1 40.4-30 68.1-11.9c32.9 21.6 22.6 60 13.8 72"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grav(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M301.1 212c4.4 4.4 4.4 11.9 0 16.3l-9.7 9.7c-4.4 4.7-11.9 4.7-16.6 0l-10.5-10.5c-4.4-4.7-4.4-11.9 0-16.6l9.7-9.7c4.4-4.4 11.9-4.4 16.6 0zm-30.2-19.7c3-3 3-7.8 0-10.5c-2.8-3-7.5-3-10.5 0c-2.8 2.8-2.8 7.5 0 10.5c3.1 2.8 7.8 2.8 10.5 0m-26 5.3c-3 2.8-3 7.5 0 10.2c2.8 3 7.5 3 10.5 0c2.8-2.8 2.8-7.5 0-10.2c-3-3-7.7-3-10.5 0m72.5-13.3c-19.9-14.4-33.8-43.2-11.9-68.1c21.6-24.9 40.7-17.2 59.8.8c11.9 11.3 29.3 24.9 17.2 48.2c-12.5 23.5-45.1 33.2-65.1 19.1m47.7-44.5c-8.9-10-23.3 6.9-15.5 16.1c7.4 9 32.1 2.4 15.5-16.1M504 256c0 137-111 248-248 248S8 393 8 256S119 8 256 8s248 111 248 248m-66.2 42.6c2.5-16.1-20.2-16.6-25.2-25.7c-13.6-24.1-27.7-36.8-54.5-30.4c11.6-8 23.5-6.1 23.5-6.1c.3-6.4 0-13-9.4-24.9c3.9-12.5.3-22.4.3-22.4c15.5-8.6 26.8-24.4 29.1-43.2c3.6-31-18.8-59.2-49.8-62.8c-22.1-2.5-43.7 7.7-54.3 25.7c-23.2 40.1 1.4 70.9 22.4 81.4c-14.4-1.4-34.3-11.9-40.1-34.3c-6.6-25.7 2.8-49.8 8.9-61.4c0 0-4.4-5.8-8-8.9c0 0-13.8 0-24.6 5.3c11.9-15.2 25.2-14.4 25.2-14.4c0-6.4-.6-14.9-3.6-21.6c-5.4-11-23.8-12.9-31.7 2.8c.1-.2.3-.4.4-.5c-5 11.9-1.1 55.9 16.9 87.2c-2.5 1.4-9.1 6.1-13 10c-21.6 9.7-56.2 60.3-56.2 60.3c-28.2 10.8-77.2 50.9-70.6 79.7c.3 3 1.4 5.5 3 7.5c-2.8 2.2-5.5 5-8.3 8.3c-11.9 13.8-5.3 35.2 17.7 24.4c15.8-7.2 29.6-20.2 36.3-30.4c0 0-5.5-5-16.3-4.4c27.7-6.6 34.3-9.4 46.2-9.1c8 3.9 8-34.3 8-34.3c0-14.7-2.2-31-11.1-41.5c12.5 12.2 29.1 32.7 28 60.6c-.8 18.3-15.2 23-15.2 23c-9.1 16.6-43.2 65.9-30.4 106c0 0-9.7-14.9-10.2-22.1c-17.4 19.4-46.5 52.3-24.6 64.5c26.6 14.7 108.8-88.6 126.2-142.3c34.6-20.8 55.4-47.3 63.9-65c22 43.5 95.3 94.5 101.1 59"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gripfire(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M112.5 301.4c0-73.8 105.1-122.5 105.1-203c0-47.1-34-88-39.1-90.4c.4 3.3.6 6.7.6 10C179.1 110.1 32 171.9 32 286.6c0 49.8 32.2 79.2 66.5 108.3c65.1 46.7 78.1 71.4 78.1 86.6c0 10.1-4.8 17-4.8 22.3c13.1-16.7 17.4-31.9 17.5-46.4c0-29.6-21.7-56.3-44.2-86.5c-16-22.3-32.6-42.6-32.6-69.5m205.3-39c-12.1-66.8-78-124.4-94.7-130.9l4 7.2c2.4 5.1 3.4 10.9 3.4 17.1c0 44.7-54.2 111.2-56.6 116.7c-2.2 5.1-3.2 10.5-3.2 15.8c0 20.1 15.2 42.1 17.9 42.1c2.4 0 56.6-55.4 58.1-87.7c6.4 11.7 9.1 22.6 9.1 33.4c0 41.2-41.8 96.9-41.8 96.9c0 11.6 31.9 53.2 35.5 53.2c1 0 2.2-1.4 3.2-2.4c37.9-39.3 67.3-85 67.3-136.8c0-8-.7-16.2-2.2-24.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grunt(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M61.3 189.3c-1.1 10 5.2 19.1 5.2 19.1c.7-7.5 2.2-12.8 4-16.6c.4 10.3 3.2 23.5 12.8 34.1c6.9 7.6 35.6 23.3 54.9 6.1c1 2.4 2.1 5.3 3 8.5c2.9 10.3-2.7 25.3-2.7 25.3s15.1-17.1 13.9-32.5c10.8-.5 21.4-8.4 21.1-19.5c0 0-18.9 10.4-35.5-8.8c-9.7-11.2-40.9-42-83.1-31.8c4.3 1 8.9 2.4 13.5 4.1h-.1c-4.2 2-6.5 7.1-7 12m28.3-1.8c19.5 11 37.4 25.7 44.9 37c-5.7 3.3-21.7 10.4-38-1.7c-10.3-7.6-9.8-26.2-6.9-35.3m142.1 45.8c-1.2 15.5 13.9 32.5 13.9 32.5s-5.6-15-2.7-25.3c.9-3.2 2-6 3-8.5c19.3 17.3 48 1.5 54.8-6.1c9.6-10.6 12.3-23.8 12.8-34.1c1.8 3.8 3.4 9.1 4 16.6c0 0 6.4-9.1 5.2-19.1c-.6-5-2.9-10-7-11.8h-.1c4.6-1.8 9.2-3.2 13.5-4.1c-42.3-10.2-73.4 20.6-83.1 31.8c-16.7 19.2-35.5 8.8-35.5 8.8c-.2 10.9 10.4 18.9 21.2 19.3m62.7-45.8c3 9.1 3.4 27.7-7 35.4c-16.3 12.1-32.2 5-37.9 1.6c7.5-11.4 25.4-26 44.9-37M160 418.5h-29.4c-5.5 0-8.2 1.6-9.5 2.9c-1.9 2-2.2 4.7-.9 8.1c3.5 9.1 11.4 16.5 13.7 18.6c3.1 2.7 7.5 4.3 11.8 4.3c4.4 0 8.3-1.7 11-4.6c7.5-8.2 11.9-17.1 13-19.8c.6-1.5 1.3-4.5-.9-6.8c-1.8-1.8-4.7-2.7-8.8-2.7m189.2-101.2c-2.4 17.9-13 33.8-24.6 43.7c-3.1-22.7-3.7-55.5-3.7-62.4c0-14.7 9.5-24.5 12.2-26.1c2.5-1.5 5.4-3 8.3-4.6c18-9.6 40.4-21.6 40.4-43.7c0-16.2-9.3-23.2-15.4-27.8c-.8-.6-1.5-1.1-2.2-1.7c-2.1-1.7-3.7-3-4.3-4.4c-4.4-9.8-3.6-34.2-1.7-37.6c.6-.6 16.7-20.9 11.8-39.2c-2-7.4-6.9-13.3-14.1-17c-5.3-2.7-11.9-4.2-19.5-4.5c-.1-2-.5-3.9-.9-5.9c-.6-2.6-1.1-5.3-.9-8.1c.4-4.7.8-9 2.2-11.3c8.4-13.3 28.8-17.6 29-17.6l12.3-2.4l-8.1-9.5c-.1-.2-17.3-17.5-46.3-17.5c-7.9 0-16 1.3-24.1 3.9c-24.2 7.8-42.9 30.5-49.4 39.3c-3.1-1-6.3-1.9-9.6-2.7c-4.2-15.8 9-38.5 9-38.5s-13.6-3-33.7 15.2c-2.6-6.5-8.1-20.5-1.8-37.2C184.6 10.1 177.2 26 175 40.4c-7.6-5.4-6.7-23.1-7.2-27.6c-7.5.9-29.2 21.9-28.2 48.3c-2 .5-3.9 1.1-5.9 1.7c-6.5-8.8-25.1-31.5-49.4-39.3c-7.9-2.2-16-3.5-23.9-3.5c-29 0-46.1 17.3-46.3 17.5L6 46.9l12.3 2.4c.2 0 20.6 4.3 29 17.6c1.4 2.2 1.8 6.6 2.2 11.3c.2 2.8-.4 5.5-.9 8.1c-.4 1.9-.8 3.9-.9 5.9c-7.7.3-14.2 1.8-19.5 4.5c-7.2 3.7-12.1 9.6-14.1 17c-5 18.2 11.2 38.5 11.8 39.2c1.9 3.4 2.7 27.8-1.7 37.6c-.6 1.4-2.2 2.7-4.3 4.4c-.7.5-1.4 1.1-2.2 1.7c-6.1 4.6-15.4 11.7-15.4 27.8c0 22.1 22.4 34.1 40.4 43.7c3 1.6 5.8 3.1 8.3 4.6c2.7 1.6 12.2 11.4 12.2 26.1c0 6.9-.6 39.7-3.7 62.4c-11.6-9.9-22.2-25.9-24.6-43.8c0 0-29.2 22.6-20.6 70.8c5.2 29.5 23.2 46.1 47 54.7c8.8 19.1 29.4 45.7 67.3 49.6C143 504.3 163 512 192.2 512h.2c29.1 0 49.1-7.7 63.6-19.5c37.9-3.9 58.5-30.5 67.3-49.6c23.8-8.7 41.7-25.2 47-54.7c8.2-48.4-21.1-70.9-21.1-70.9M305.7 37.7c5.6-1.8 11.6-2.7 17.7-2.7c11 0 19.9 3 24.7 5c-3.1 1.4-6.4 3.2-9.7 5.3c-2.4-.4-5.6-.8-9.2-.8c-10.5 0-20.5 3.1-28.7 8.9c-12.3 8.7-18 16.9-20.7 22.4c-2.2-1.3-4.5-2.5-7.1-3.7c-1.6-.8-3.1-1.5-4.7-2.2c6.1-9.1 19.9-26.5 37.7-32.2m21 18.2c-.8 1-1.6 2.1-2.3 3.2c-3.3 5.2-3.9 11.6-4.4 17.8c-.5 6.4-1.1 12.5-4.4 17c-4.2.8-8.1 1.7-11.5 2.7c-2.3-3.1-5.6-7-10.5-11.2c1.4-4.8 5.5-16.1 13.5-22.5c5.6-4.3 12.2-6.7 19.6-7M45.6 45.3c-3.3-2.2-6.6-4-9.7-5.3c4.8-2 13.7-5 24.7-5c6.1 0 12 .9 17.7 2.7c17.8 5.8 31.6 23.2 37.7 32.1c-1.6.7-3.2 1.4-4.8 2.2c-2.5 1.2-4.9 2.5-7.1 3.7c-2.6-5.4-8.3-13.7-20.7-22.4c-8.3-5.8-18.2-8.9-28.8-8.9c-3.4.1-6.6.5-9 .9m44.7 40.1c-4.9 4.2-8.3 8-10.5 11.2c-3.4-.9-7.3-1.9-11.5-2.7C65 89.5 64.5 83.4 64 77c-.5-6.2-1.1-12.6-4.4-17.8c-.7-1.1-1.5-2.2-2.3-3.2c7.4.3 14 2.6 19.5 7c8 6.3 12.1 17.6 13.5 22.4M58.1 259.9c-2.7-1.6-5.6-3.1-8.4-4.6c-14.9-8-30.2-16.3-30.2-30.5c0-11.1 4.3-14.6 8.9-18.2l.5-.4c.7-.6 1.4-1.2 2.2-1.8c-.9 7.2-1.9 13.3-2.7 14.9c0 0 12.1-15 15.7-44.3c1.4-11.5-1.1-34.3-5.1-43c.2 4.9 0 9.8-.3 14.4c-.4-.8-.8-1.6-1.3-2.2c-3.2-4-11.8-17.5-9.4-26.6c.9-3.5 3.1-6 6.7-7.8c3.8-1.9 8.8-2.9 15.1-2.9c12.3 0 25.9 3.7 32.9 6c25.1 8 55.4 30.9 64.1 37.7c.2.2.4.3.4.3l5.6 3.9l-3.5-5.8c-.2-.3-19.1-31.4-53.2-46.5c2-2.9 7.4-8.1 21.6-15.1c21.4-10.5 46.5-15.8 74.3-15.8c27.9 0 52.9 5.3 74.3 15.8c14.2 6.9 19.6 12.2 21.6 15.1c-34 15.1-52.9 46.2-53.1 46.5l-3.5 5.8l5.6-3.9s.2-.1.4-.3c8.7-6.8 39-29.8 64.1-37.7c7-2.2 20.6-6 32.9-6c6.3 0 11.3 1 15.1 2.9c3.5 1.8 5.7 4.4 6.7 7.8c2.5 9.1-6.1 22.6-9.4 26.6c-.5.6-.9 1.3-1.3 2.2c-.3-4.6-.5-9.5-.3-14.4c-4 8.8-6.5 31.5-5.1 43c3.6 29.3 15.7 44.3 15.7 44.3c-.8-1.6-1.8-7.7-2.7-14.9c.7.6 1.5 1.2 2.2 1.8l.5.4c4.6 3.7 8.9 7.1 8.9 18.2c0 14.2-15.4 22.5-30.2 30.5c-2.9 1.5-5.7 3.1-8.4 4.6c-8.7 5-18 16.7-19.1 34.2c-.9 14.6.9 49.9 3.4 75.9c-12.4 4.8-26.7 6.4-39.7 6.8c-2-4.1-3.9-8.5-5.5-13.1c-.7-2-19.6-51.1-26.4-62.2c5.5 39 17.5 73.7 23.5 89.6c-3.5-.5-7.3-.7-11.7-.7h-117c-4.4 0-8.3.3-11.7.7c6-15.9 18.1-50.6 23.5-89.6c-6.8 11.2-25.7 60.3-26.4 62.2c-1.6 4.6-3.5 9-5.5 13.1c-13-.4-27.2-2-39.7-6.8c2.5-26 4.3-61.2 3.4-75.9c-.9-17.4-10.3-29.2-19-34.2M34.8 404.6c-12.1-20-8.7-54.1-3.7-59.1c10.9 34.4 47.2 44.3 74.4 45.4c-2.7 4.2-5.2 7.6-7 10l-1.4 1.4c-7.2 7.8-8.6 18.5-4.1 31.8c-22.7-.1-46.3-9.8-58.2-29.5m45.7 43.5c6 1.1 12.2 1.9 18.6 2.4c3.5 8 7.4 15.9 12.3 23.1c-14.4-5.9-24.4-16-30.9-25.5M192 498.2c-60.6-.1-78.3-45.8-84.9-64.7c-3.7-10.5-3.4-18.2.9-23.1c2.9-3.3 9.5-7.2 24.6-7.2h118.8c15.1 0 21.8 3.9 24.6 7.2c4.2 4.8 4.5 12.6.9 23.1c-6.6 18.8-24.3 64.6-84.9 64.7m80.6-24.6c4.9-7.2 8.8-15.1 12.3-23.1c6.4-.5 12.6-1.3 18.6-2.4c-6.5 9.5-16.5 19.6-30.9 25.5m76.6-69c-12 19.7-35.6 29.3-58.1 29.7c4.5-13.3 3.1-24.1-4.1-31.8c-.4-.5-.9-1-1.4-1.5c-1.8-2.4-4.3-5.8-7-10c27.2-1.2 63.5-11 74.4-45.4c5 5 8.4 39.1-3.8 59M191.9 187.7h.2c12.7-.1 27.2-17.8 27.2-17.8c-9.9 6-18.8 8.1-27.3 8.3c-8.5-.2-17.4-2.3-27.3-8.3c0 0 14.5 17.6 27.2 17.8m61.7 230.7h-29.4c-4.2 0-7.2.9-8.9 2.7c-2.2 2.3-1.5 5.2-.9 6.7c1 2.6 5.5 11.3 13 19.3c2.7 2.9 6.6 4.5 11 4.5s8.7-1.6 11.8-4.2c2.3-2 10.2-9.2 13.7-18.1c1.3-3.3 1-6-.9-7.9c-1.3-1.3-4-2.9-9.4-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Guilded(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M443.427 64H4.571c0 103.26 22.192 180.06 43.418 222.358C112.046 414.135 224 448 225.256 448a312.824 312.824 0 0 0 140.55-103.477c25.907-33.923 53.1-87.19 65.916-145.761H171.833c4.14 36.429 22.177 67.946 45.1 86.944h88.589c-17.012 28.213-48.186 54.4-80.456 69.482c-31.232-13.259-69.09-46.544-96.548-98.362c-26.726-53.833-27.092-105.883-27.092-105.883h336.147A625.91 625.91 0 0 0 443.427 64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gulp(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m209.8 391.1l-14.1 24.6l-4.6 80.2c0 8.9-28.3 16.1-63.1 16.1s-63.1-7.2-63.1-16.1l-5.8-79.4l-14.9-25.4c41.2 17.3 126 16.7 165.6 0m-196-253.3l13.6 125.5c5.9-20 20.8-47 40-55.2c6.3-2.7 12.7-2.7 18.7.9c5.2 3 9.6 9.3 10.1 11.8c1.2 6.5-2 9.1-4.5 9.1c-3 0-5.3-4.6-6.8-7.3c-4.1-7.3-10.3-7.6-16.9-2.8c-6.9 5-12.9 13.4-17.1 20.7c-5.1 8.8-9.4 18.5-12 28.2c-1.5 5.6-2.9 14.6-.6 19.9c1 2.2 2.5 3.6 4.9 3.6c5 0 12.3-6.6 15.8-10.1c4.5-4.5 10.3-11.5 12.5-16l5.2-15.5c2.6-6.8 9.9-5.6 9.9 0c0 10.2-3.7 13.6-10 34.7c-5.8 19.5-7.6 25.8-7.6 25.8c-.7 2.8-3.4 7.5-6.3 7.5c-1.2 0-2.1-.4-2.6-1.2c-1-1.4-.9-5.3-.8-6.3c.2-3.2 6.3-22.2 7.3-25.2c-2 2.2-4.1 4.4-6.4 6.6c-5.4 5.1-14.1 11.8-21.5 11.8c-3.4 0-5.6-.9-7.7-2.4l7.6 79.6c2 5 39.2 17.1 88.2 17.1c49.1 0 86.3-12.2 88.2-17.1l10.9-94.6c-5.7 5.2-12.3 11.6-19.6 14.8c-5.4 2.3-17.4 3.8-17.4-5.7c0-5.2 9.1-14.8 14.4-21.5c1.4-1.7 4.7-5.9 4.7-8.1c0-2.9-6-2.2-11.7 2.5c-3.2 2.7-6.2 6.3-8.7 9.7c-4.3 6-6.6 11.2-8.5 15.5c-6.2 14.2-4.1 8.6-9.1 22c-5 13.3-4.2 11.8-5.2 14c-.9 1.9-2.2 3.5-4 4.5c-1.9 1-4.5.9-6.1-.3c-.9-.6-1.3-1.9-1.3-3.7c0-.9.1-1.8.3-2.7c1.5-6.1 7.8-18.1 15-34.3c1.6-3.7 1-2.6.8-2.3c-6.2 6-10.9 8.9-14.4 10.5c-5.8 2.6-13 2.6-14.5-4.1c-.1-.4-.1-.8-.2-1.2c-11.8 9.2-24.3 11.7-20-8.1c-4.6 8.2-12.6 14.9-22.4 14.9c-4.1 0-7.1-1.4-8.6-5.1c-2.3-5.5 1.3-14.9 4.6-23.8c1.7-4.5 4-9.9 7.1-16.2c1.6-3.4 4.2-5.4 7.6-4.5c.6.2 1.1.4 1.6.7c2.6 1.8 1.6 4.5.3 7.2c-3.8 7.5-7.1 13-9.3 20.8c-.9 3.3-2 9 1.5 9c2.4 0 4.7-.8 6.9-2.4c4.6-3.4 8.3-8.5 11.1-13.5c2-3.6 4.4-8.3 5.6-12.3c.5-1.7 1.1-3.3 1.8-4.8c1.1-2.5 2.6-5.1 5.2-5.1c1.3 0 2.4.5 3.2 1.5c1.7 2.2 1.3 4.5.4 6.9c-2 5.6-4.7 10.6-6.9 16.7c-1.3 3.5-2.7 8-2.7 11.7c0 3.4 3.7 2.6 6.8 1.2c2.4-1.1 4.8-2.8 6.8-4.5c1.2-4.9.9-3.8 26.4-68.2c1.3-3.3 3.7-4.7 6.1-4.7c1.2 0 2.2.4 3.2 1.1c1.7 1.3 1.7 4.1 1 6.2c-.7 1.9-.6 1.3-4.5 10.5c-5.2 12.1-8.6 20.8-13.2 31.9c-1.9 4.6-7.7 18.9-8.7 22.3c-.6 2.2-1.3 5.8 1 5.8c5.4 0 19.3-13.1 23.1-17c.2-.3.5-.4.9-.6c.6-1.9 1.2-3.7 1.7-5.5c1.4-3.8 2.7-8.2 5.3-11.3c.8-1 1.7-1.6 2.7-1.6c2.8 0 4.2 1.2 4.2 4c0 1.1-.7 5.1-1.1 6.2c1.4-1.5 2.9-3 4.5-4.5c15-13.9 25.7-6.8 25.7.2c0 7.4-8.9 17.7-13.8 23.4c-1.6 1.9-4.9 5.4-5 6.4c0 1.3.9 1.8 2.2 1.8c2 0 6.4-3.5 8-4.7c5-3.9 11.8-9.9 16.6-14.1l14.8-136.8c-30.5 17.1-197.6 17.2-228.3.2m229.7-8.5c0 21-231.2 21-231.2 0c0-8.8 51.8-15.9 115.6-15.9c9 0 17.8.1 26.3.4l12.6-48.7L228.1.6c1.4-1.4 5.8-.2 9.9 3.5s6.6 7.9 5.3 9.3l-.1.1L185.9 74l-10 40.7c39.9 2.6 67.6 8.1 67.6 14.6m-69.4 4.6c0-.8-.9-1.5-2.5-2.1l-.2.8c0 1.3-5 2.4-11.1 2.4s-11.1-1.1-11.1-2.4c0-.1 0-.2.1-.3l.2-.7c-1.8.6-3 1.4-3 2.3c0 2.1 6.2 3.7 13.7 3.7c7.7.1 13.9-1.6 13.9-3.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HackerNews(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 32v448h448V32zm21.2 197.2H21c.1-.1.2-.3.3-.4c0 .1 0 .3-.1.4m218 53.9V384h-31.4V281.3L128 128h37.3c52.5 98.3 49.2 101.2 59.3 125.6c12.3-27 5.8-24.4 60.6-125.6H320z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HackerNewsSquare(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 32H48C21.5 32 0 53.5 0 80v352c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V80c0-26.5-21.5-48-48-48M21.2 229.2H21c.1-.1.2-.3.3-.4c0 .1 0 .3-.1.4m218 53.9V384h-31.4V281.3L128 128h37.3c52.5 98.3 49.2 101.2 59.3 125.6c12.3-27 5.8-24.4 60.6-125.6H320z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hackerrank(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M477.5 128C463 103.05 285.13 0 256.16 0S49.25 102.79 34.84 128s-14.49 230.8 0 256s192.38 128 221.32 128S463 409.08 477.49 384s14.51-231 .01-256M316.13 414.22c-4 0-40.91-35.77-38-38.69c.87-.87 6.26-1.48 17.55-1.83c0-26.23.59-68.59.94-86.32c0-2-.44-3.43-.44-5.85h-79.93c0 7.1-.46 36.2 1.37 72.88c.23 4.54-1.58 6-5.74 5.94c-10.13 0-20.27-.11-30.41-.08c-4.1 0-5.87-1.53-5.74-6.11c.92-33.44 3-84-.15-212.67v-3.17c-9.67-.35-16.38-1-17.26-1.84c-2.92-2.92 34.54-38.69 38.49-38.69s41.17 35.78 38.27 38.69c-.87.87-7.9 1.49-16.77 1.84v3.16c-2.42 25.75-2 79.59-2.63 105.39h80.26c0-4.55.39-34.74-1.2-83.64c-.1-3.39.95-5.17 4.21-5.2c11.07-.08 22.15-.13 33.23-.06c3.46 0 4.57 1.72 4.5 5.38C333 354.64 336 341.29 336 373.69c8.87.35 16.82 1 17.69 1.84c2.88 2.91-33.62 38.69-37.58 38.69z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hips(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M251.6 157.6c0-1.9-.9-2.8-2.8-2.8h-40.9c-1.6 0-2.7 1.4-2.7 2.8v201.8c0 1.4 1.1 2.8 2.7 2.8h40.9c1.9 0 2.8-.9 2.8-2.8zM156.5 168c-16.1-11.8-36.3-17.9-60.3-18c-18.1-.1-34.6 3.7-49.8 11.4V80.2c0-1.8-.9-2.7-2.8-2.7H2.7c-1.8 0-2.7.9-2.7 2.7v279.2c0 1.9.9 2.8 2.7 2.8h41c1.9 0 2.8-.9 2.8-2.8V223.3c0-.8-2.8-27 45.8-27c48.5 0 45.8 26.1 45.8 27v122.6c0 9 7.3 16.3 16.4 16.3h27.3c1.8 0 2.7-.9 2.7-2.8V223.3c0-23.4-9.3-41.8-28-55.3m478.4 110.1c-6.8-15.7-18.4-27-34.9-34.1l-57.6-25.3c-8.6-3.6-9.2-11.2-2.6-16.1c7.4-5.5 44.3-13.9 84 6.8c1.7 1 4-.3 4-2.4v-44.7c0-1.3-.6-2.1-1.9-2.6c-17.7-6.6-36.1-9.9-55.1-9.9c-26.5 0-45.3 5.8-58.5 15.4c-.5.4-28.4 20-22.7 53.7c3.4 19.6 15.8 34.2 37.2 43.6l53.6 23.5c11.6 5.1 15.2 13.3 12.2 21.2c-3.7 9.1-13.2 13.6-36.5 13.6c-24.3 0-44.7-8.9-58.4-19.1c-2.1-1.4-4.4.2-4.4 2.3v34.4c0 10.4 4.9 17.3 14.6 20.7c15.6 5.5 31.6 8.2 48.2 8.2c12.7 0 25.8-1.2 36.3-4.3c.7-.3 36-8.9 45.6-45.8c3.5-13.5 2.4-26.5-3.1-39.1M376.2 149.8c-31.7 0-104.2 20.1-104.2 103.5v183.5c0 .8.6 2.7 2.7 2.7h40.9c1.9 0 2.8-.9 2.8-2.7V348c16.5 12.7 35.8 19.1 57.7 19.1c60.5 0 108.7-48.5 108.7-108.7c.1-60.3-48.2-108.6-108.6-108.6m0 170.9c-17.2 0-31.9-6.1-44-18.2c-12.2-12.2-18.2-26.8-18.2-44c0-34.5 27.6-62.2 62.2-62.2c34.5 0 62.2 27.6 62.2 62.2c.1 34.3-27.3 62.2-62.2 62.2M228.3 72.5c-15.9 0-28.8 12.9-28.9 28.9c0 15.6 12.7 28.9 28.9 28.9s28.9-13.1 28.9-28.9c0-16.2-13-28.9-28.9-28.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HireAhelper(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M443.1 0H71.9C67.9 37.3 37.4 67.8 0 71.7v371.5c37.4 4.9 66 32.4 71.9 68.8h372.2c3-36.4 32.5-65.8 67.9-69.8V71.7c-36.4-5.9-65-35.3-68.9-71.7m-37 404.9c-36.3 0-18.8-2-55.1-2c-35.8 0-21 2-56.1 2c-5.9 0-4.9-8.2 0-9.8c22.8-7.6 22.9-10.2 24.6-12.8c10.4-15.6 5.9-83 5.9-113c0-5.3-6.4-12.8-13.8-12.8H200.4c-7.4 0-13.8 7.5-13.8 12.8c0 30-4.5 97.4 5.9 113c1.7 2.5 1.8 5.2 24.6 12.8c4.9 1.6 6 9.8 0 9.8c-35.1 0-20.3-2-56.1-2c-36.3 0-18.8 2-55.1 2c-7.9 0-5.8-10.8 0-10.8c10.2-3.4 13.5-3.5 21.7-13.8c7.7-12.9 7.9-44.4 7.9-127.8V151.3c0-22.2-12.2-28.3-28.6-32.4c-8.8-2.2-4-11.8 1-11.8c36.5 0 20.6 2 57.1 2c32.7 0 16.5-2 49.2-2c3.3 0 8.5 8.3 1 10.8c-4.9 1.6-27.6 3.7-27.6 39.3c0 45.6-.2 55.8 1 68.8c0 1.3 2.3 12.8 12.8 12.8h109.2c10.5 0 12.8-11.5 12.8-12.8c1.2-13 1-23.2 1-68.8c0-35.6-22.7-37.7-27.6-39.3c-7.5-2.5-2.3-10.8 1-10.8c32.7 0 16.5 2 49.2 2c36.5 0 20.6-2 57.1-2c4.9 0 9.9 9.6 1 11.8c-16.4 4.1-28.6 10.3-28.6 32.4v101.2c0 83.4.1 114.9 7.9 127.8c8.2 10.2 11.4 10.4 21.7 13.8c5.8 0 7.8 10.8 0 10.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hive(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M260.353 254.878L131.538 33.1a2.208 2.208 0 0 0-3.829.009L.3 254.887a2.234 2.234 0 0 0 0 2.235L129.116 478.9a2.208 2.208 0 0 0 3.83-.009l127.412-221.778a2.239 2.239 0 0 0-.005-2.235m39.078-25.713a2.19 2.19 0 0 0 1.9 1.111h66.509a2.226 2.226 0 0 0 1.9-3.341L259.115 33.111a2.187 2.187 0 0 0-1.9-1.111h-66.508a2.226 2.226 0 0 0-1.9 3.341ZM511.7 254.886L384.9 33.112A2.2 2.2 0 0 0 382.99 32h-66.6a2.226 2.226 0 0 0-1.906 3.34L440.652 256L314.481 476.66a2.226 2.226 0 0 0 1.906 3.34h66.6a2.2 2.2 0 0 0 1.906-1.112L511.7 257.114a2.243 2.243 0 0 0 0-2.228m-145.684 30.031h-66.508a2.187 2.187 0 0 0-1.9 1.111l-108.8 190.631a2.226 2.226 0 0 0 1.9 3.341h66.509a2.187 2.187 0 0 0 1.9-1.111l108.8-190.631a2.226 2.226 0 0 0-1.901-3.341"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hooli(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m144.5 352l38.3.8c-13.2-4.6-26-10.2-38.3-16.8zm57.7-5.3v5.3l-19.4.8c36.5 12.5 69.9 14.2 94.7 7.2c-19.9.2-45.8-2.6-75.3-13.3m408.9-115.2c15.9 0 28.9-12.9 28.9-28.9s-12.9-24.5-28.9-24.5c-15.9 0-28.9 8.6-28.9 24.5s12.9 28.9 28.9 28.9m-29 120.5H640V241.5h-57.9zm-73.7 0h57.9V156.7L508.4 184zm-31-119.4c-18.2-18.2-50.4-17.1-50.4-17.1s-32.3-1.1-50.4 17.1c-18.2 18.2-16.8 33.9-16.8 52.6s-1.4 34.3 16.8 52.5s50.4 17.1 50.4 17.1s32.3 1.1 50.4-17.1c18.2-18.2 16.8-33.8 16.8-52.5c-.1-18.8 1.3-34.5-16.8-52.6m-39.8 71.9c0 3.6-1.8 12.5-10.7 12.5s-10.7-8.9-10.7-12.5v-40.4c0-8.7 7.3-10.9 10.7-10.9s10.7 2.1 10.7 10.9zm-106.2-71.9c-18.2-18.2-50.4-17.1-50.4-17.1s-32.2-1.1-50.4 17.1c-1.9 1.9-3.7 3.9-5.3 6c-38.2-29.6-72.5-46.5-102.1-61.1v-20.7l-22.5 10.6c-54.4-22.1-89-18.2-97.3.1c0 0-24.9 32.8 61.8 110.8V352h57.9v-28.6c-6.5-4.2-13-8.7-19.4-13.6c-14.8-11.2-27.4-21.6-38.4-31.4v-31c13.1 14.7 30.5 31.4 53.4 50.3l4.5 3.6v-29.8c0-6.9 1.7-18.2 10.8-18.2s10.6 6.9 10.6 15V317c18 12.2 37.3 22.1 57.7 29.6v-93.9c0-18.7-13.4-37.4-40.6-37.4c-15.8-.1-30.5 8.2-38.5 21.9v-54.3c41.9 20.9 83.9 46.5 99.9 58.3c-10.2 14.6-9.3 28.1-9.3 43.7c0 18.7-1.4 34.3 16.8 52.5s50.4 17.1 50.4 17.1s32.3 1.1 50.4-17.1c18.2-18.2 16.7-33.8 16.7-52.5c0-18.5 1.5-34.2-16.7-52.3M65.2 184v63.3c-48.7-54.5-38.9-76-35.2-79.1c13.5-11.4 37.5-8 64.4 2.1zm226.5 120.5c0 3.6-1.8 12.5-10.7 12.5s-10.7-8.9-10.7-12.5v-40.4c0-8.7 7.3-10.9 10.7-10.9s10.7 2.1 10.7 10.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hornbill(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M76.38 370.3a37.8 37.8 0 1 1-32.07-32.42c-78.28-111.35 52-190.53 52-190.53c-5.86 43-8.24 91.16-8.24 91.16c-67.31 41.49.93 64.06 39.81 72.87a140.38 140.38 0 0 0 131.66 91.94c1.92 0 3.77-.21 5.67-.28l.11 18.86c-99.22 1.39-158.7-29.14-188.94-51.6m108-327.7A37.57 37.57 0 0 0 181 21.45a37.95 37.95 0 1 0-31.17 54.22c-22.55 29.91-53.83 89.57-52.42 190l21.84-.15c0-.9-.14-1.77-.14-2.68A140.42 140.42 0 0 1 207 132.71c8-37.71 30.7-114.3 73.8-44.29c0 0 48.14 2.38 91.18 8.24c0 0-77.84-128-187.59-54.06zm304.19 134.17a37.94 37.94 0 1 0-53.84-28.7C403 126.13 344.89 99 251.28 100.33l.14 22.5c2.7-.15 5.39-.41 8.14-.41a140.37 140.37 0 0 1 130.49 88.76c39.1 9 105.06 31.58 38.46 72.54c0 0-2.34 48.13-8.21 91.16c0 0 133.45-81.16 49-194.61a37.45 37.45 0 0 0 19.31-3.5zM374.06 436.24c21.43-32.46 46.42-89.69 45.14-179.66l-19.52.14c.08 2.06.3 4.07.3 6.15a140.34 140.34 0 0 1-91.39 131.45c-8.85 38.95-31.44 106.66-72.77 39.49c0 0-48.12-2.34-91.19-8.22c0 0 79.92 131.34 191.9 51a37.5 37.5 0 0 0 3.64 14a37.93 37.93 0 1 0 33.89-54.29z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hotjar(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M414.9 161.5C340.2 29 121.1 0 121.1 0S222.2 110.4 93 197.7C11.3 252.8-21 324.4 14 402.6c26.8 59.9 83.5 84.3 144.6 93.4c-29.2-55.1-6.6-122.4-4.1-129.6c57.1 86.4 165 0 110.8-93.9c71 15.4 81.6 138.6 27.1 215.5c80.5-25.3 134.1-88.9 148.8-145.6c15.5-59.3 3.7-127.9-26.3-180.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Houzz(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M275.9 330.7H171.3V480H17V32h109.5v104.5l305.1 85.6V480H275.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HtmlFive(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 32l34.9 395.8L191.5 480l157.6-52.2L384 32zm308.2 127.9H124.4l4.1 49.4h175.6l-13.6 148.4l-97.9 27v.3h-1.1l-98.7-27.3l-6-75.8h47.7L138 320l53.5 14.5l53.7-14.5l6-62.2H84.3L71.5 112.2h241.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hubspot(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M267.4 211.6c-25.1 23.7-40.8 57.3-40.8 94.6c0 29.3 9.7 56.3 26 78L203.1 434c-4.4-1.6-9.1-2.5-14-2.5c-10.8 0-20.9 4.2-28.5 11.8c-7.6 7.6-11.8 17.8-11.8 28.6s4.2 20.9 11.8 28.5c7.6 7.6 17.8 11.6 28.5 11.6c10.8 0 20.9-3.9 28.6-11.6c7.6-7.6 11.8-17.8 11.8-28.5c0-4.2-.6-8.2-1.9-12.1l50-50.2c22 16.9 49.4 26.9 79.3 26.9c71.9 0 130-58.3 130-130.2c0-65.2-47.7-119.2-110.2-128.7V116c17.5-7.4 28.2-23.8 28.2-42.9c0-26.1-20.9-47.9-47-47.9S311.2 47 311.2 73.1c0 19.1 10.7 35.5 28.2 42.9v61.2c-15.2 2.1-29.6 6.7-42.7 13.6c-27.6-20.9-117.5-85.7-168.9-124.8c1.2-4.4 2-9 2-13.8C129.8 23.4 106.3 0 77.4 0C48.6 0 25.2 23.4 25.2 52.2c0 28.9 23.4 52.3 52.2 52.3c9.8 0 18.9-2.9 26.8-7.6zm89.5 163.6c-38.1 0-69-30.9-69-69s30.9-69 69-69s69 30.9 69 69s-30.9 69-69 69"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ideal(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M125.61 165.48a49.07 49.07 0 1 0 49.06 49.06a49.08 49.08 0 0 0-49.06-49.06M86.15 425.84h78.94V285.32H86.15Zm151.46-211.6c0-20-10-22.53-18.74-22.53h-14.05v45.79h14.05c9.75 0 18.74-2.81 18.74-23.26m201.69 46v-91.31h22.75v68.57h33.69C486.5 113.08 388.61 86.19 299.67 86.19h-94.83V169h14c25.6 0 41.5 17.35 41.5 45.26c0 28.81-15.52 46-41.5 46h-14v165.62h94.83c144.61 0 194.94-67.16 196.72-165.64Zm-109.75 0H273.3V169h54.43v22.73H296v10.58h30V225h-30v12.5h33.51Zm74.66 0l-5.16-17.67h-29.74l-5.18 17.67h-23.66L368 168.92h32.35l27.53 91.34ZM299.65 32H32v448h267.65c161.85 0 251-79.73 251-224.52C550.62 172 518 32 299.65 32m0 426.92H53.07V53.07h246.58c142.1 0 229.9 64.61 229.9 202.41c0 134.09-81 203.44-229.9 203.44m83.86-264.85L376 219.88h16.4l-7.52-25.81Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Imdb(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 32H48C21.5 32 0 53.5 0 80v352c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V80c0-26.5-21.5-48-48-48M21.3 229.2H21c.1-.1.2-.3.3-.4zM97 319.8H64V192h33zm113.2 0h-28.7v-86.4l-11.6 86.4h-20.6l-12.2-84.5v84.5h-29V192h42.8c3.3 19.8 6 39.9 8.7 59.9l7.6-59.9h43zm11.4 0V192h24.6c17.6 0 44.7-1.6 49 20.9c1.7 7.6 1.4 16.3 1.4 24.4c0 88.5 11.1 82.6-75 82.5m160.9-29.2c0 15.7-2.4 30.9-22.2 30.9c-9 0-15.2-3-20.9-9.8l-1.9 8.1h-29.8V192h31.7v41.7c6-6.5 12-9.2 20.9-9.2c21.4 0 22.2 12.8 22.2 30.1zM265 229.9c0-9.7 1.6-16-10.3-16v83.7c12.2.3 10.3-8.7 10.3-18.4zm85.5 26.1c0-5.4 1.1-12.7-6.2-12.7c-6 0-4.9 8.9-4.9 12.7c0 .6-1.1 39.6 1.1 44.7c.8 1.6 2.2 2.4 3.8 2.4c7.8 0 6.2-9 6.2-14.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Innosoft(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M422.559 159.71a27.379 27.379 0 0 0-13.866-23.337a26.42 26.42 0 0 0-26.211.133L73.943 314.647V176.261a11.955 11.955 0 0 1 6.047-10.34l138.076-79.713a12.153 12.153 0 0 1 11.922.025l32.656 18.853l-150.063 86.637v56l247.061-142.637l-118.513-68.407c-10.992-6.129-22.3-6.255-33.8-.27l-164.6 95.026c-10.634 6.12-16.771 16.39-17.29 29.124v191.5c.17 10.135 5.08 18.672 13.474 23.428a27.037 27.037 0 0 0 26.736-.045l308.408-178.066v138.281a11.976 11.976 0 0 1-5.92 10.368L230.025 425.77a12.175 12.175 0 0 1-11.937.062l-32.723-18.9l150.051-86.627v-56L88.367 406.932l118.794 68.583a33.88 33.88 0 0 0 34.25-.327l164.527-94.995c10.746-6.631 16.649-17.118 16.624-29.528Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Instagram(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M224.1 141c-63.6 0-114.9 51.3-114.9 114.9s51.3 114.9 114.9 114.9S339 319.5 339 255.9S287.7 141 224.1 141m0 189.6c-41.1 0-74.7-33.5-74.7-74.7s33.5-74.7 74.7-74.7s74.7 33.5 74.7 74.7s-33.6 74.7-74.7 74.7m146.4-194.3c0 14.9-12 26.8-26.8 26.8c-14.9 0-26.8-12-26.8-26.8s12-26.8 26.8-26.8s26.8 12 26.8 26.8m76.1 27.2c-1.7-35.9-9.9-67.7-36.2-93.9c-26.2-26.2-58-34.4-93.9-36.2c-37-2.1-147.9-2.1-184.9 0c-35.8 1.7-67.6 9.9-93.9 36.1s-34.4 58-36.2 93.9c-2.1 37-2.1 147.9 0 184.9c1.7 35.9 9.9 67.7 36.2 93.9s58 34.4 93.9 36.2c37 2.1 147.9 2.1 184.9 0c35.9-1.7 67.7-9.9 93.9-36.2c26.2-26.2 34.4-58 36.2-93.9c2.1-37 2.1-147.8 0-184.8M398.8 388c-7.8 19.6-22.9 34.7-42.6 42.6c-29.5 11.7-99.5 9-132.1 9s-102.7 2.6-132.1-9c-19.6-7.8-34.7-22.9-42.6-42.6c-11.7-29.5-9-99.5-9-132.1s-2.6-102.7 9-132.1c7.8-19.6 22.9-34.7 42.6-42.6c29.5-11.7 99.5-9 132.1-9s102.7-2.6 132.1 9c19.6 7.8 34.7 22.9 42.6 42.6c11.7 29.5 9 99.5 9 132.1s2.7 102.7-9 132.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InstagramSquare(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M224 202.66A53.34 53.34 0 1 0 277.36 256A53.38 53.38 0 0 0 224 202.66m124.71-41a54 54 0 0 0-30.41-30.41c-21-8.29-71-6.43-94.3-6.43s-73.25-1.93-94.31 6.43a54 54 0 0 0-30.41 30.41c-8.28 21-6.43 71.05-6.43 94.33s-1.85 73.27 6.47 94.34a54 54 0 0 0 30.41 30.41c21 8.29 71 6.43 94.31 6.43s73.24 1.93 94.3-6.43a54 54 0 0 0 30.41-30.41c8.35-21 6.43-71.05 6.43-94.33s1.92-73.26-6.43-94.33ZM224 338a82 82 0 1 1 82-82a81.9 81.9 0 0 1-82 82m85.38-148.3a19.14 19.14 0 1 1 19.13-19.14a19.1 19.1 0 0 1-19.09 19.18ZM400 32H48A48 48 0 0 0 0 80v352a48 48 0 0 0 48 48h352a48 48 0 0 0 48-48V80a48 48 0 0 0-48-48m-17.12 290c-1.29 25.63-7.14 48.34-25.85 67s-41.4 24.63-67 25.85c-26.41 1.49-105.59 1.49-132 0c-25.63-1.29-48.26-7.15-67-25.85s-24.63-41.42-25.85-67c-1.49-26.42-1.49-105.61 0-132c1.29-25.63 7.07-48.34 25.85-67s41.47-24.56 67-25.78c26.41-1.49 105.59-1.49 132 0c25.63 1.29 48.33 7.15 67 25.85s24.63 41.42 25.85 67.05c1.49 26.32 1.49 105.44 0 131.88"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Instalod(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M153.384 480h233.729l115.441-204.235l-298.325 57.446Zm351.342-239.922L387.113 32H155.669L360.23 267.9ZM124.386 48.809L7.274 256l115.962 205.154l102.391-295.593Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Intercom(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M392 32H56C25.1 32 0 57.1 0 88v336c0 30.9 25.1 56 56 56h336c30.9 0 56-25.1 56-56V88c0-30.9-25.1-56-56-56m-108.3 82.1c0-19.8 29.9-19.8 29.9 0v199.5c0 19.8-29.9 19.8-29.9 0zm-74.6-7.5c0-19.8 29.9-19.8 29.9 0v216.5c0 19.8-29.9 19.8-29.9 0zm-74.7 7.5c0-19.8 29.9-19.8 29.9 0v199.5c0 19.8-29.9 19.8-29.9 0zM59.7 144c0-19.8 29.9-19.8 29.9 0v134.3c0 19.8-29.9 19.8-29.9 0zm323.4 227.8c-72.8 63-241.7 65.4-318.1 0c-15-12.8 4.4-35.5 19.4-22.7c65.9 55.3 216.1 53.9 279.3 0c14.9-12.9 34.3 9.8 19.4 22.7m5.2-93.5c0 19.8-29.9 19.8-29.9 0V144c0-19.8 29.9-19.8 29.9 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InternetExplorer(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M483.049 159.706c10.855-24.575 21.424-60.438 21.424-87.871c0-72.722-79.641-98.371-209.673-38.577c-107.632-7.181-211.221 73.67-237.098 186.457c30.852-34.862 78.271-82.298 121.977-101.158C125.404 166.85 79.128 228.002 43.992 291.725C23.246 329.651 0 390.94 0 436.747c0 98.575 92.854 86.5 180.251 42.006c31.423 15.43 66.559 15.573 101.695 15.573c97.124 0 184.249-54.294 216.814-146.022H377.927c-52.509 88.593-196.819 52.996-196.819-47.436H509.9c6.407-43.581-1.655-95.715-26.851-141.162M64.559 346.877c17.711 51.15 53.703 95.871 100.266 123.304c-88.741 48.94-173.267 29.096-100.266-123.304m115.977-108.873c2-55.151 50.276-94.871 103.98-94.871c53.418 0 101.981 39.72 103.981 94.871zm184.536-187.6c21.425-10.287 48.563-22.003 72.558-22.003c31.422 0 54.274 21.717 54.274 53.722c0 20.003-7.427 49.007-14.569 67.867c-26.28-42.292-65.986-81.584-112.263-99.586"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Invision(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M407.4 32H40.6C18.2 32 0 50.2 0 72.6v366.8C0 461.8 18.2 480 40.6 480h366.8c22.4 0 40.6-18.2 40.6-40.6V72.6c0-22.4-18.2-40.6-40.6-40.6M176.1 145.6c.4 23.4-22.4 27.3-26.6 27.4c-14.9 0-27.1-12-27.1-27c.1-35.2 53.1-35.5 53.7-.4M332.8 377c-65.6 0-34.1-74-25-106.6c14.1-46.4-45.2-59-59.9.7l-25.8 103.3H177l8.1-32.5c-31.5 51.8-94.6 44.4-94.6-4.3c.1-14.3.9-14 23-104.1H81.7l9.7-35.6h76.4c-33.6 133.7-32.6 126.9-32.9 138.2c0 20.9 40.9 13.5 57.4-23.2l19.8-79.4h-32.3l9.7-35.6h68.8l-8.9 40.5c40.5-75.5 127.9-47.8 101.8 38c-14.2 51.1-14.6 50.7-14.9 58.8c0 15.5 17.5 22.6 31.8-16.9L386 325c-10.5 36.7-29.4 52-53.2 52"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ioxhost(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M616 160h-67.3C511.2 70.7 422.9 8 320 8C183 8 72 119 72 256c0 16.4 1.6 32.5 4.7 48H24c-13.3 0-24 10.8-24 24c0 13.3 10.7 24 24 24h67.3c37.5 89.3 125.8 152 228.7 152c137 0 248-111 248-248c0-16.4-1.6-32.5-4.7-48H616c13.3 0 24-10.8 24-24c0-13.3-10.7-24-24-24m-96 96c0 110.5-89.5 200-200 200c-75.7 0-141.6-42-175.5-104H424c13.3 0 24-10.8 24-24c0-13.3-10.7-24-24-24H125.8c-3.8-15.4-5.8-31.4-5.8-48c0-110.5 89.5-200 200-200c75.7 0 141.6 42 175.5 104H216c-13.3 0-24 10.8-24 24c0 13.3 10.7 24 24 24h298.2c3.8 15.4 5.8 31.4 5.8 48m-304-24h208c13.3 0 24 10.7 24 24c0 13.2-10.7 24-24 24H216c-13.3 0-24-10.7-24-24c0-13.2 10.7-24 24-24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ItchIo(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M71.92 34.77C50.2 47.67 7.4 96.84 7 109.73v21.34c0 27.06 25.29 50.84 48.25 50.84c27.57 0 50.54-22.85 50.54-50c0 27.12 22.18 50 49.76 50s49-22.85 49-50c0 27.12 23.59 50 51.16 50h.5c27.57 0 51.16-22.85 51.16-50c0 27.12 21.47 50 49 50s49.76-22.85 49.76-50c0 27.12 23 50 50.54 50c23 0 48.25-23.78 48.25-50.84v-21.34c-.4-12.9-43.2-62.07-64.92-75C372.56 32.4 325.76 32 256 32S91.14 33.1 71.92 34.77m132.32 134.39c-22 38.4-77.9 38.71-99.85.25c-13.17 23.14-43.17 32.07-56 27.66c-3.87 40.15-13.67 237.13 17.73 269.15c80 18.67 302.08 18.12 379.76 0c31.65-32.27 21.32-232 17.75-269.15c-12.92 4.44-42.88-4.6-56-27.66c-22 38.52-77.85 38.1-99.85-.24c-7.1 12.49-23.05 28.94-51.76 28.94a57.54 57.54 0 0 1-51.75-28.94zm-41.58 53.77c16.47 0 31.09 0 49.22 19.78a436.91 436.91 0 0 1 88.18 0C318.22 223 332.85 223 349.31 223c52.33 0 65.22 77.53 83.87 144.45c17.26 62.15-5.52 63.67-33.95 63.73c-42.15-1.57-65.49-32.18-65.49-62.79c-39.25 6.43-101.93 8.79-155.55 0c0 30.61-23.34 61.22-65.49 62.79c-28.42-.06-51.2-1.58-33.94-63.73c18.67-67 31.56-144.45 83.88-144.45zM256 270.79s-44.38 40.77-52.35 55.21l29-1.17v25.32c0 1.55 21.34.16 23.33.16c11.65.54 23.31 1 23.31-.16v-25.28l29 1.17c-8-14.48-52.35-55.24-52.35-55.24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Itunes(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M223.6 80.3C129 80.3 52.5 157 52.5 251.5S129 422.8 223.6 422.8s171.2-76.7 171.2-171.2c0-94.6-76.7-171.3-171.2-171.3m79.4 240c-3.2 13.6-13.5 21.2-27.3 23.8c-12.1 2.2-22.2 2.8-31.9-5c-11.8-10-12-26.4-1.4-36.8c8.4-8 20.3-9.6 38-12.8c3-.5 5.6-1.2 7.7-3.7c3.2-3.6 2.2-2 2.2-80.8c0-5.6-2.7-7.1-8.4-6.1c-4 .7-91.9 17.1-91.9 17.1c-5 1.1-6.7 2.6-6.7 8.3c0 116.1.5 110.8-1.2 118.5c-2.1 9-7.6 15.8-14.9 19.6c-8.3 4.6-23.4 6.6-31.4 5.2c-21.4-4-28.9-28.7-14.4-42.9c8.4-8 20.3-9.6 38-12.8c3-.5 5.6-1.2 7.7-3.7c5-5.7.9-127 2.6-133.7c.4-2.6 1.5-4.8 3.5-6.4c2.1-1.7 5.8-2.7 6.7-2.7c101-19 113.3-21.4 115.1-21.4c5.7-.4 9 3 9 8.7c-.1 170.6.4 161.4-1 167.6M345.2 32H102.8C45.9 32 0 77.9 0 134.8v242.4C0 434.1 45.9 480 102.8 480h242.4c57 0 102.8-45.9 102.8-102.8V134.8C448 77.9 402.1 32 345.2 32M223.6 444c-106.3 0-192.5-86.2-192.5-192.5S117.3 59 223.6 59s192.5 86.2 192.5 192.5S329.9 444 223.6 444"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ItunesNote(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M381.9 388.2c-6.4 27.4-27.2 42.8-55.1 48c-24.5 4.5-44.9 5.6-64.5-10.2c-23.9-20.1-24.2-53.4-2.7-74.4c17-16.2 40.9-19.5 76.8-25.8c6-1.1 11.2-2.5 15.6-7.4c6.4-7.2 4.4-4.1 4.4-163.2c0-11.2-5.5-14.3-17-12.3c-8.2 1.4-185.7 34.6-185.7 34.6c-10.2 2.2-13.4 5.2-13.4 16.7c0 234.7 1.1 223.9-2.5 239.5c-4.2 18.2-15.4 31.9-30.2 39.5c-16.8 9.3-47.2 13.4-63.4 10.4c-43.2-8.1-58.4-58-29.1-86.6c17-16.2 40.9-19.5 76.8-25.8c6-1.1 11.2-2.5 15.6-7.4c10.1-11.5 1.8-256.6 5.2-270.2c.8-5.2 3-9.6 7.1-12.9c4.2-3.5 11.8-5.5 13.4-5.5c204-38.2 228.9-43.1 232.4-43.1c11.5-.8 18.1 6 18.1 17.6c.2 344.5 1.1 326-1.8 338.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Java(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M277.74 312.9c9.8-6.7 23.4-12.5 23.4-12.5s-38.7 7-77.2 10.2c-47.1 3.9-97.7 4.7-123.1 1.3c-60.1-8 33-30.1 33-30.1s-36.1-2.4-80.6 19c-52.5 25.4 130 37 224.5 12.1m-85.4-32.1c-19-42.7-83.1-80.2 0-145.8C296 53.2 242.84 0 242.84 0c21.5 84.5-75.6 110.1-110.7 162.6c-23.9 35.9 11.7 74.4 60.2 118.2m114.6-176.2c.1 0-175.2 43.8-91.5 140.2c24.7 28.4-6.5 54-6.5 54s62.7-32.4 33.9-72.9c-26.9-37.8-47.5-56.6 64.1-121.3m-6.1 270.5a12.19 12.19 0 0 1-2 2.6c128.3-33.7 81.1-118.9 19.8-97.3a17.33 17.33 0 0 0-8.2 6.3a70.45 70.45 0 0 1 11-3c31-6.5 75.5 41.5-20.6 91.4M348 437.4s14.5 11.9-15.9 21.2c-57.9 17.5-240.8 22.8-291.6.7c-18.3-7.9 16-19 26.8-21.3c11.2-2.4 17.7-2 17.7-2c-20.3-14.3-131.3 28.1-56.4 40.2C232.84 509.4 401 461.3 348 437.4M124.44 396c-78.7 22 47.9 67.4 148.1 24.5a185.89 185.89 0 0 1-28.2-13.8c-44.7 8.5-65.4 9.1-106 4.5c-33.5-3.8-13.9-15.2-13.9-15.2m179.8 97.2c-78.7 14.8-175.8 13.1-233.3 3.6c0-.1 11.8 9.7 72.4 13.6c92.2 5.9 233.8-3.3 237.1-46.9c0 0-6.4 16.5-76.2 29.7M260.64 353c-59.2 11.4-93.5 11.1-136.8 6.6c-33.5-3.5-11.6-19.7-11.6-19.7c-86.8 28.8 48.2 61.4 169.5 25.9a60.37 60.37 0 0 1-21.1-12.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JediOrder(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M398.5 373.6c95.9-122.1 17.2-233.1 17.2-233.1c45.4 85.8-41.4 170.5-41.4 170.5c105-171.5-60.5-271.5-60.5-271.5c96.9 72.7-10.1 190.7-10.1 190.7c85.8 158.4-68.6 230.1-68.6 230.1s-.4-16.9-2.2-85.7c4.3 4.5 34.5 36.2 34.5 36.2l-24.2-47.4l62.6-9.1l-62.6-9.1l20.2-55.5l-31.4 45.9c-2.2-87.7-7.8-305.1-7.9-306.9v-2.4v1v-1v2.4c0 1-5.6 219-7.9 306.9l-31.4-45.9l20.2 55.5l-62.6 9.1l62.6 9.1l-24.2 47.4l34.5-36.2c-1.8 68.8-2.2 85.7-2.2 85.7s-154.4-71.7-68.6-230.1c0 0-107-118.1-10.1-190.7c0 0-165.5 99.9-60.5 271.5c0 0-86.8-84.8-41.4-170.5c0 0-78.7 111 17.2 233.1c0 0-26.2-16.1-49.4-77.7c0 0 16.9 183.3 222 185.7h4.1c205-2.4 222-185.7 222-185.7c-23.6 61.5-49.9 77.7-49.9 77.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Jenkins(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M487.1 425c-1.4-11.2-19-23.1-28.2-31.9c-5.1-5-29-23.1-30.4-29.9c-1.4-6.6 9.7-21.5 13.3-28.9c5.1-10.7 8.8-23.7 11.3-32.6c18.8-66.1 20.7-156.9-6.2-211.2c-10.2-20.6-38.6-49-56.4-62.5c-42-31.7-119.6-35.3-170.1-16.6c-14.1 5.2-27.8 9.8-40.1 17.1c-33.1 19.4-68.3 32.5-78.1 71.6c-24.2 10.8-31.5 41.8-30.3 77.8c.2 7 4.1 15.8 2.7 22.4c-.7 3.3-5.2 7.6-6.1 9.8c-11.6 27.7-2.3 64 11.1 83.7c8.1 11.9 21.5 22.4 39.2 25.2c.7 10.6 3.3 19.7 8.2 30.4c3.1 6.8 14.7 19 10.4 27.7c-2.2 4.4-21 13.8-27.3 17.6C89 407.2 73.7 415 54.2 429c-12.6 9-32.3 10.2-29.2 31.1c2.1 14.1 10.1 31.6 14.7 45.8c.7 2 1.4 4.1 2.1 6h422c4.9-15.3 9.7-30.9 14.6-47.2c3.4-11.4 10.2-27.8 8.7-39.7M205.9 33.7c1.8-.5 3.4.7 4.9 2.4c-.2 5.2-5.4 5.1-8.9 6.8c-5.4 6.7-13.4 9.8-20 17.2c-6.8 7.5-14.4 27.7-23.4 30c-4.5 1.1-9.7-.8-13.6-.5c-10.4.7-17.7 6-28.3 7.5c13.6-29.9 56.1-54 89.3-63.4m-104.8 93.6c13.5-14.9 32.1-24.1 54.8-25.9c11.7 29.7-8.4 65-.9 97.6c2.3 9.9 10.2 25.4-2.4 25.7c.3-28.3-34.8-46.3-61.3-29.6c-1.8-21.5-4.9-51.7 9.8-67.8m36.7 200.2c-1-4.1-2.7-12.9-2.3-15.1c1.6-8.7 17.1-12.5 11-24.7c-11.3-.1-13.8 10.2-24.1 11.3c-26.7 2.6-45.6-35.4-44.4-58.4c1-19.5 17.6-38.2 40.1-35.8c16 1.8 21.4 19.2 24.5 34.7c9.2.5 22.5-.4 26.9-7.6c-.6-17.5-8.8-31.6-8.2-47.7c1-30.3 17.5-57.6 4.8-87.4c13.6-30.9 53.5-55.3 83.1-70c36.6-18.3 94.9-3.7 129.3 15.8c19.7 11.1 34.4 32.7 48.3 50.7c-19.5-5.8-36.1 4.2-33.1 20.3c16.3-14.9 44.2-.2 52.5 16.4c7.9 15.8 7.8 39.3 9 62.8c2.9 57-10.4 115.9-39.1 157.1c-7.7 11-14.1 23-24.9 30.6c-26 18.2-65.4 34.7-99.2 23.4c-44.7-15-65-44.8-89.5-78.8c.7 18.7 13.8 34.1 26.8 48.4c11.3 12.5 25 26.6 39.7 32.4c-12.3-2.9-31.1-3.8-36.2 7.2c-28.6-1.9-55.1-4.8-68.7-24.2c-10.6-15.4-21.4-41.4-26.3-61.4m222 124.1c4.1-3 11.1-2.9 17.4-3.6c-5.4-2.7-13-3.7-19.3-2.2c-.1-4.2-2-6.8-3.2-10.2c10.6-3.8 35.5-28.5 49.6-20.3c6.7 3.9 9.5 26.2 10.1 37c.4 9-.8 18-4.5 22.8c-18.8-.6-35.8-2.8-50.7-7c.9-6.1-1-12.1.6-16.5m-17.2-20c-16.8.8-26-1.2-38.3-10.8c.2-.8 1.4-.5 1.5-1.4c18 8 40.8-3.3 59-4.9c-7.9 5.1-14.6 11.6-22.2 17.1m-12.1 33.2c-1.6-9.4-3.5-12-2.8-20.2c25-16.6 29.7 28.6 2.8 20.2M226 438.6c-11.6-.7-48.1-14-38.5-23.7c9.4 6.5 27.5 4.9 41.3 7.3c.8 4.4-2.8 10.2-2.8 16.4M57.7 497.1c-4.3-12.7-9.2-25.1-14.8-36.9c30.8-23.8 65.3-48.9 102.2-63.5c2.8-1.1 23.2 25.4 26.2 27.6c16.5 11.7 37 21 56.2 30.2c1.2 8.8 3.9 20.2 8.7 35.5c.7 2.3 1.4 4.7 2.2 7.2H57.7zm240.6 5.7h-.8c.3-.2.5-.4.8-.5zm7.5-5.7c2.1-1.4 4.3-2.8 6.4-4.3c1.1 1.4 2.2 2.8 3.2 4.3zm15.1-24.7c-10.8 7.3-20.6 18.3-33.3 25.2c-6 3.3-27 11.7-33.4 10.2c-3.6-.8-3.9-5.3-5.4-9.5c-3.1-9-10.1-23.4-10.8-37c-.8-17.2-2.5-46 16-42.4c14.9 2.9 32.3 9.7 43.9 16.1c7.1 3.9 11.1 8.6 21.9 9.5c-.1 1.4-.1 2.8-.2 4.3c-5.9 3.9-15.3 3.8-21.8 7.1c9.5.4 17 2.7 23.5 5.9c-.1 3.4-.3 7-.4 10.6m53.4 24.7h-14c-.1-3.2-2.8-5.8-6.1-5.8s-5.9 2.6-6.1 5.8h-17.4c-2.8-4.4-5.7-8.6-8.9-12.5c2.1-2.2 4-4.7 6-6.9c9 3.7 14.8-4.9 21.7-4.2c7.9.8 14.2 11.7 25.4 11zm8.7 0c.2-4 .4-7.8.6-11.5c15.6-7.3 29 1.3 35.7 11.5zm83.4-37c-2.3 11.2-5.8 24-9.9 37.1c-.2-.1-.4-.1-.6-.1H428c.6-1.1 1.2-2.2 1.9-3.3c-2.6-6.1-9-8.7-10.9-15.5c12.1-22.7 6.5-93.4-24.2-78.5c4.3-6.3 15.6-11.5 20.8-19.3c13 10.4 20.8 20.3 33.2 31.4c6.8 6 20 13.3 21.4 23.1c.8 5.5-2.6 18.9-3.8 25.1M222.2 130.5c5.4-14.9 27.2-34.7 45-32c7.7 1.2 18 8.2 12.2 17.7c-30.2-7-45.2 12.6-54.4 33.1c-8.1-2-4.9-13.1-2.8-18.8m184.1 63.1c8.2-3.6 22.4-.7 29.6-5.3c-4.2-11.5-10.3-21.4-9.3-37.7c.5 0 1 0 1.4.1c6.8 14.2 12.7 29.2 21.4 41.7c-5.7 13.5-43.6 25.4-43.1 1.2m-96.8 2.7c-6.8-10.9-19-32.5-14.5-45.3c6.5 11.9 8.6 24.4 17.8 33.3c4.1 4 12.2 9 8.2 20.2c-.9 2.7-7.8 8.6-11.7 9.7c-14.4 4.3-47.9.9-36.6-17.1c11.9.7 27.9 7.8 36.8-.8m27.3 70c3.8 6.6 1.4 18.7 12.1 20.6c20.2 3.4 43.6-12.3 58.1-17.8c9-15.2-.8-20.7-8.9-30.5c-16.6-20-38.8-44.8-38-74.7c6.7-4.9 7.3 7.4 8.2 9.7c8.7 20.3 30.4 46.2 46.3 63.5c3.9 4.3 10.3 8.4 11 11.2c2.1 8.2-5.4 18-4.5 23.5c-21.7 13.9-45.8 29.1-81.4 25.6c-7.4-6.7-10.3-21.4-2.9-31.1m-201.3-9.2c-6.8-3.9-8.4-21-16.4-21.4c-11.4-.7-9.3 22.2-9.3 35.5c-7.8-7.1-9.2-29.1-3.5-40.3c-6.6-3.2-9.5 3.6-13.1 5.9c4.7-34.1 49.8-15.8 42.3 20.3m299.6 28.8c-10.1 19.2-24.4 40.4-54 41c-.6-6.2-1.1-15.6 0-19.4c22.7-2.2 36.6-13.7 54-21.6m-141.9 12.4c18.9 9.9 53.6 11 79.3 10.2c1.4 5.6 1.3 12.6 1.4 19.4c-33 1.8-72-6.4-80.7-29.6m92.2 46.7c-1.7 4.3-5.3 9.3-9.8 11.1c-12.1 4.9-45.6 8.7-62.4-.3c-10.7-5.7-17.5-18.5-23.4-26c-2.8-3.6-16.9-12.9-.2-12.9c13.1 32.7 58 29 95.8 28.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Jira(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M490 241.7C417.1 169 320.6 71.8 248.5 0C83 164.9 6 241.7 6 241.7c-7.9 7.9-7.9 20.7 0 28.7C138.8 402.7 67.8 331.9 248.5 512c379.4-378 15.7-16.7 241.5-241.7c8-7.9 8-20.7 0-28.6m-241.5 90l-76-75.7l76-75.7l76 75.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Joget(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M378.1 45C337.6 19.9 292.6 8 248.2 8C165 8 83.8 49.9 36.9 125.9c-71.9 116.6-35.6 269.3 81 341.2s269.3 35.6 341.2-80.9c71.9-116.6 35.6-269.4-81-341.2m51.8 323.2c-40.4 65.5-110.4 101.5-182 101.5c-6.8 0-13.6-.4-20.4-1c-9-13.6-19.9-33.3-23.7-42.4c-5.7-13.7-27.2-45.6 31.2-67.1c51.7-19.1 176.7-16.5 208.8-17.6c-4 9-8.6 17.9-13.9 26.6m-200.8-86.3c-55.5-1.4-81.7-20.8-58.5-48.2s51.1-40.7 68.9-51.2c17.9-10.5 27.3-33.7-23.6-29.7C87.3 161.5 48.6 252.1 37.6 293c-8.8-49.7-.1-102.7 28.5-149.1C128 43.4 259.6 12.2 360.1 74.1c74.8 46.1 111.2 130.9 99.3 212.7c-24.9-.5-179.3-3.6-230.3-4.9m183.8-54.8c-22.7-6-57 11.3-86.7 27.2c-29.7 15.8-31.1 8.2-31.1 8.2s40.2-28.1 50.7-34.5s31.9-14 13.4-24.6c-3.2-1.8-6.7-2.7-10.4-2.7c-17.8 0-41.5 18.7-67.5 35.6c-31.5 20.5-65.3 31.3-65.3 31.3l169.5-1.6l46.5-23.4s3.6-9.5-19.1-15.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Joomla(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.6 92.1C.6 58.8 27.4 32 60.4 32c30 0 54.5 21.9 59.2 50.2c32.6-7.6 67.1.6 96.5 30l-44.3 44.3c-20.5-20.5-42.6-16.3-55.4-3.5c-14.3 14.3-14.3 37.9 0 52.2l99.5 99.5l-44 44.3c-87.7-87.2-49.7-49.7-99.8-99.7c-26.8-26.5-35-64.8-24.8-98.9C20.4 144.6.6 120.7.6 92.1m129.5 116.4l44.3 44.3c10-10 89.7-89.7 99.7-99.8c14.3-14.3 37.6-14.3 51.9 0c12.8 12.8 17 35-3.5 55.4l44 44.3c31.2-31.2 38.5-67.6 28.9-101.2c29.2-4.1 51.9-29.2 51.9-59.5c0-33.2-26.8-60.1-59.8-60.1c-30.3 0-55.4 22.5-59.5 51.6c-33.8-9.9-71.7-1.5-98.3 25.1c-18.3 19.1-71.1 71.5-99.6 99.9m266.3 152.2c8.2-32.7-.9-68.5-26.3-93.9c-11.8-12.2 5 4.7-99.5-99.7l-44.3 44.3l99.7 99.7c14.3 14.3 14.3 37.6 0 51.9c-12.8 12.8-35 17-55.4-3.5l-44 44.3c27.6 30.2 68 38.8 102.7 28c5.5 27.4 29.7 48.1 58.9 48.1c33 0 59.8-26.8 59.8-60.1c0-30.2-22.5-55-51.6-59.1m-84.3-53.1l-44-44.3c-87 86.4-50.4 50.4-99.7 99.8c-14.3 14.3-37.6 14.3-51.9 0c-13.1-13.4-16.9-35.3 3.2-55.4l-44-44.3c-30.2 30.2-38 65.2-29.5 98.3c-26.7 6-46.2 29.9-46.2 58.2C0 453.2 26.8 480 59.8 480c28.6 0 52.5-19.8 58.6-46.7c32.7 8.2 68.5-.6 94.2-26c32.1-32 12.2-12.4 99.5-99.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Js(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 32v448h448V32zm243.8 349.4c0 43.6-25.6 63.5-62.9 63.5c-33.7 0-53.2-17.4-63.2-38.5l34.3-20.7c6.6 11.7 12.6 21.6 27.1 21.6c13.8 0 22.6-5.4 22.6-26.5V237.7h42.1zm99.6 63.5c-39.1 0-64.4-18.6-76.7-43l34.3-19.8c9 14.7 20.8 25.6 41.5 25.6c17.4 0 28.6-8.7 28.6-20.8c0-14.4-11.4-19.5-30.7-28l-10.5-4.5c-30.4-12.9-50.5-29.2-50.5-63.5c0-31.6 24.1-55.6 61.6-55.6c26.8 0 46 9.3 59.8 33.7L368 290c-7.2-12.9-15-18-27.1-18c-12.3 0-20.1 7.8-20.1 18c0 12.6 7.8 17.7 25.9 25.6l10.5 4.5c35.8 15.3 55.9 31 55.9 66.2c0 37.8-29.8 58.6-69.7 58.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JsSquare(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 32H48C21.5 32 0 53.5 0 80v352c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V80c0-26.5-21.5-48-48-48M243.8 381.4c0 43.6-25.6 63.5-62.9 63.5c-33.7 0-53.2-17.4-63.2-38.5l34.3-20.7c6.6 11.7 12.6 21.6 27.1 21.6c13.8 0 22.6-5.4 22.6-26.5V237.7h42.1zm99.6 63.5c-39.1 0-64.4-18.6-76.7-43l34.3-19.8c9 14.7 20.8 25.6 41.5 25.6c17.4 0 28.6-8.7 28.6-20.8c0-14.4-11.4-19.5-30.7-28l-10.5-4.5c-30.4-12.9-50.5-29.2-50.5-63.5c0-31.6 24.1-55.6 61.6-55.6c26.8 0 46 9.3 59.8 33.7L368 290c-7.2-12.9-15-18-27.1-18c-12.3 0-20.1 7.8-20.1 18c0 12.6 7.8 17.7 25.9 25.6l10.5 4.5c35.8 15.3 55.9 31 55.9 66.2c0 37.8-29.8 58.6-69.7 58.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Jsfiddle(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M510.634 237.462c-4.727-2.621-5.664-5.748-6.381-10.776c-2.352-16.488-3.539-33.619-9.097-49.095c-35.895-99.957-153.99-143.386-246.849-91.646c-27.37 15.25-48.971 36.369-65.493 63.903c-3.184-1.508-5.458-2.71-7.824-3.686c-30.102-12.421-59.049-10.121-85.331 9.167c-25.531 18.737-36.422 44.548-32.676 76.408c.355 3.025-1.967 7.621-4.514 9.545c-39.712 29.992-56.031 78.065-41.902 124.615c13.831 45.569 57.514 79.796 105.608 81.433c30.291 1.031 60.637.546 90.959.539c84.041-.021 168.09.531 252.12-.48c52.664-.634 96.108-36.873 108.212-87.293c11.54-48.074-11.144-97.3-56.832-122.634m21.107 156.88c-18.23 22.432-42.343 35.253-71.28 35.65c-56.874.781-113.767.23-170.652.23c0 .7-163.028.159-163.728.154c-43.861-.332-76.739-19.766-95.175-59.995c-18.902-41.245-4.004-90.848 34.186-116.106c9.182-6.073 12.505-11.566 10.096-23.136c-5.49-26.361 4.453-47.956 26.42-62.981c22.987-15.723 47.422-16.146 72.034-3.083c10.269 5.45 14.607 11.564 22.198-2.527c14.222-26.399 34.557-46.727 60.671-61.294c97.46-54.366 228.37 7.568 230.24 132.697c.122 8.15 2.412 12.428 9.848 15.894c57.56 26.829 74.456 96.122 35.142 144.497m-87.789-80.499c-5.848 31.157-34.622 55.096-66.666 55.095c-16.953-.001-32.058-6.545-44.079-17.705c-27.697-25.713-71.141-74.98-95.937-93.387c-20.056-14.888-41.99-12.333-60.272 3.782c-49.996 44.071 15.859 121.775 67.063 77.188c4.548-3.96 7.84-9.543 12.744-12.844c8.184-5.509 20.766-.884 13.168 10.622c-17.358 26.284-49.33 38.197-78.863 29.301c-28.897-8.704-48.84-35.968-48.626-70.179c1.225-22.485 12.364-43.06 35.414-55.965c22.575-12.638 46.369-13.146 66.991 2.474C295.68 280.7 320.467 323.97 352.185 343.47c24.558 15.099 54.254 7.363 68.823-17.506c28.83-49.209-34.592-105.016-78.868-63.46c-3.989 3.744-6.917 8.932-11.41 11.72c-10.975 6.811-17.333-4.113-12.809-10.353c20.703-28.554 50.464-40.44 83.271-28.214c31.429 11.714 49.108 44.366 42.76 78.186"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Kaggle(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M304.2 501.5L158.4 320.3L298.2 185c2.6-2.7 1.7-10.5-5.3-10.5h-69.2c-3.5 0-7 1.8-10.5 5.3L80.9 313.5V7.5q0-7.5-7.5-7.5H21.5Q14 0 14 7.5v497q0 7.5 7.5 7.5h51.9q7.5 0 7.5-7.5v-109l30.8-29.3l110.5 140.6c3 3.5 6.5 5.3 10.5 5.3h66.9q5.25 0 6-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keybase(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M286.17 419a18 18 0 1 0 18 18a18 18 0 0 0-18-18m111.92-147.6c-9.5-14.62-39.37-52.45-87.26-73.71q-9.1-4.06-18.38-7.27a78.43 78.43 0 0 0-47.88-104.13c-12.41-4.1-23.33-6-32.41-5.77c-.6-2-1.89-11 9.4-35L198.66 32l-5.48 7.56c-8.69 12.06-16.92 23.55-24.34 34.89a51 51 0 0 0-8.29-1.25c-41.53-2.45-39-2.33-41.06-2.33c-50.61 0-50.75 52.12-50.75 45.88l-2.36 36.68c-1.61 27 19.75 50.21 47.63 51.85l8.93.54a214 214 0 0 0-46.29 35.54C14 304.66 14 374 14 429.77v33.64l23.32-29.8a148.6 148.6 0 0 0 14.56 37.56c5.78 10.13 14.87 9.45 19.64 7.33c4.21-1.87 10-6.92 3.75-20.11a178.29 178.29 0 0 1-15.76-53.13l46.82-59.83l-24.66 74.11c58.23-42.4 157.38-61.76 236.25-38.59c34.2 10.05 67.45.69 84.74-23.84c.72-1 1.2-2.16 1.85-3.22a156.09 156.09 0 0 1 2.8 28.43c0 23.3-3.69 52.93-14.88 81.64c-2.52 6.46 1.76 14.5 8.6 15.74c7.42 1.57 15.33-3.1 18.37-11.15C429 443 434 414 434 382.32c0-38.58-13-77.46-35.91-110.92M142.37 128.58l-15.7-.93l-1.39 21.79l13.13.78a93 93 0 0 0 .32 19.57l-22.38-1.34a12.28 12.28 0 0 1-11.76-12.79L107 119c1-12.17 13.87-11.27 13.26-11.32l29.11 1.73a144.35 144.35 0 0 0-7 19.17m148.42 172.18a10.51 10.51 0 0 1-14.35-1.39l-9.68-11.49l-34.42 27a8.09 8.09 0 0 1-11.13-1.08l-15.78-18.64a7.38 7.38 0 0 1 1.34-10.34l34.57-27.18l-14.14-16.74l-17.09 13.45a7.75 7.75 0 0 1-10.59-1s-3.72-4.42-3.8-4.53a7.38 7.38 0 0 1 1.37-10.34L214 225.19s-18.51-22-18.6-22.14a9.56 9.56 0 0 1 1.74-13.42a10.38 10.38 0 0 1 14.3 1.37l81.09 96.32a9.58 9.58 0 0 1-1.74 13.44M187.44 419a18 18 0 1 0 18 18a18 18 0 0 0-18-18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keycdn(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m63.8 409.3l60.5-59c32.1 42.8 71.1 66 126.6 67.4c30.5.7 60.3-7 86.4-22.4c5.1 5.3 18.5 19.5 20.9 22c-32.2 20.7-69.6 31.1-108.1 30.2c-43.3-1.1-84.6-16.7-117.7-44.4c.3-.6-38.2 37.5-38.6 37.9c9.5 29.8-13.1 62.4-46.3 62.4C20.7 503.3 0 481.7 0 454.9c0-34.3 33.1-56.6 63.8-45.6m354.9-252.4c19.1 31.3 29.6 67.4 28.7 104c-1.1 44.8-19 87.5-48.6 121c.3.3 23.8 25.2 24.1 25.5c9.6-1.3 19.2 2 25.9 9.1c11.3 12 10.9 30.9-1.1 42.4c-12 11.3-30.9 10.9-42.4-1.1c-6.7-7-9.4-16.8-7.6-26.3c-24.9-26.6-44.4-47.2-44.4-47.2c42.7-34.1 63.3-79.6 64.4-124.2c.7-28.9-7.2-57.2-21.1-82.2zM104 53.1c6.7 7 9.4 16.8 7.6 26.3l45.9 48.1c-4.7 3.8-13.3 10.4-22.8 21.3c-25.4 28.5-39.6 64.8-40.7 102.9c-.7 28.9 6.1 57.2 20 82.4l-22 21.5C72.7 324 63.1 287.9 64.2 250.9c1-44.6 18.3-87.6 47.5-121.1l-25.3-26.4c-9.6 1.3-19.2-2-25.9-9.1c-11.3-12-10.9-30.9 1.1-42.4C73.5 40.7 92.2 41 104 53.1M464.9 8c26 0 47.1 22.4 47.1 48.3S490.9 104 464.9 104c-6.3.1-14-1.1-15.9-1.8l-62.9 59.7c-32.7-43.6-76.7-65.9-126.9-67.2c-30.5-.7-60.3 6.8-86.2 22.4l-21.1-22C184.1 74.3 221.5 64 260 64.9c43.3 1.1 84.6 16.7 117.7 44.6l41.1-38.6c-1.5-4.7-2.2-9.6-2.2-14.5C416.5 29.7 438.9 8 464.9 8M256.7 113.4c5.5 0 10.9.4 16.4 1.1c78.1 9.8 133.4 81.1 123.8 159.1c-9.8 78.1-81.1 133.4-159.1 123.8c-78.1-9.8-133.4-81.1-123.8-159.2c9.3-72.4 70.1-124.6 142.7-124.8m-59 119.4c.6 22.7 12.2 41.8 32.4 52.2l-11 51.7h73.7l-11-51.7c20.1-10.9 32.1-29 32.4-52.2c-.4-32.8-25.8-57.5-58.3-58.3c-32.1.8-57.3 24.8-58.2 58.3M256 160"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Kickstarter(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 480H48c-26.4 0-48-21.6-48-48V80c0-26.4 21.6-48 48-48h352c26.4 0 48 21.6 48 48v352c0 26.4-21.6 48-48 48M199.6 178.5c0-30.7-17.6-45.1-39.7-45.1c-25.8 0-40 19.8-40 44.5v154.8c0 25.8 13.7 45.6 40.5 45.6c21.5 0 39.2-14 39.2-45.6v-41.8l60.6 75.7c12.3 14.9 39 16.8 55.8 0c14.6-15.1 14.8-36.8 4-50.4l-49.1-62.8l40.5-58.7c9.4-13.5 9.5-34.5-5.6-49.1c-16.4-15.9-44.6-17.3-61.4 7l-44.8 64.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KickstarterK(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M147.3 114.4c0-56.2-32.5-82.4-73.4-82.4C26.2 32 0 68.2 0 113.4v283c0 47.3 25.3 83.4 74.9 83.4c39.8 0 72.4-25.6 72.4-83.4v-76.5l112.1 138.3c22.7 27.2 72.1 30.7 103.2 0c27-27.6 27.3-67.4 7.4-92.2l-90.8-114.8l74.9-107.4c17.4-24.7 17.5-63.1-10.4-89.8c-30.3-29-82.4-31.6-113.6 12.8L147.3 185z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Korvue(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M386.5 34h-327C26.8 34 0 60.8 0 93.5v327.1C0 453.2 26.8 480 59.5 480h327.1c33 0 59.5-26.8 59.5-59.5v-327C446 60.8 419.2 34 386.5 34M87.1 120.8h96v116l61.8-116h110.9l-81.2 132H87.1zm161.8 272.1l-65.7-113.6v113.6h-96V262.1h191.5l88.6 130.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laravel(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M504.4 115.83a5.72 5.72 0 0 0-.28-.68a8.52 8.52 0 0 0-.53-1.25a6 6 0 0 0-.54-.71a9.36 9.36 0 0 0-.72-.94c-.23-.22-.52-.4-.77-.6a8.84 8.84 0 0 0-.9-.68L404.4 55.55a8 8 0 0 0-8 0L300.12 111a8.07 8.07 0 0 0-.88.69a7.68 7.68 0 0 0-.78.6a8.23 8.23 0 0 0-.72.93c-.17.24-.39.45-.54.71a9.7 9.7 0 0 0-.52 1.25c-.08.23-.21.44-.28.68a8.08 8.08 0 0 0-.28 2.08v105.24l-80.22 46.19V63.44a7.8 7.8 0 0 0-.28-2.09c-.06-.24-.2-.45-.28-.68a8.35 8.35 0 0 0-.52-1.24c-.14-.26-.37-.47-.54-.72a9.36 9.36 0 0 0-.72-.94a9.46 9.46 0 0 0-.78-.6a9.8 9.8 0 0 0-.88-.68L115.61 1.07a8 8 0 0 0-8 0L11.34 56.49a6.52 6.52 0 0 0-.88.69a7.81 7.81 0 0 0-.79.6a8.15 8.15 0 0 0-.71.93c-.18.25-.4.46-.55.72a7.88 7.88 0 0 0-.51 1.24a6.46 6.46 0 0 0-.29.67a8.18 8.18 0 0 0-.28 2.1v329.7a8 8 0 0 0 4 6.95l192.5 110.84a8.83 8.83 0 0 0 1.33.54c.21.08.41.2.63.26a7.92 7.92 0 0 0 4.1 0c.2-.05.37-.16.55-.22a8.6 8.6 0 0 0 1.4-.58L404.4 400.09a8 8 0 0 0 4-6.95V287.88l92.24-53.11a8 8 0 0 0 4-7V117.92a8.63 8.63 0 0 0-.24-2.09M111.6 17.28l80.19 46.15l-80.2 46.18l-80.18-46.17Zm88.25 60V278.6l-46.53 26.79l-33.69 19.4V123.5l46.53-26.79Zm0 412.78L23.37 388.5V77.32L57.06 96.7l46.52 26.8v215.18a6.94 6.94 0 0 0 .12.9a8 8 0 0 0 .16 1.18a5.92 5.92 0 0 0 .38.9a6.38 6.38 0 0 0 .42 1a8.54 8.54 0 0 0 .6.78a7.62 7.62 0 0 0 .66.84c.23.22.52.38.77.58a8.93 8.93 0 0 0 .86.66l92.19 52.18Zm8-106.17l-80.06-45.32l84.09-48.41l92.26-53.11l80.13 46.13l-58.8 33.56Zm184.52 4.57L215.88 490.11V397.8l130.72-74.6l45.77-26.15Zm0-119.13L358.68 250l-46.53-26.79v-91.42l33.69 19.4L392.37 178Zm8-105.28l-80.2-46.17l80.2-46.16l80.18 46.15Zm8 105.28V178L455 151.19l33.68-19.4v91.39Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lastfm(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m225.8 367.1l-18.8-51s-30.5 34-76.2 34c-40.5 0-69.2-35.2-69.2-91.5c0-72.1 36.4-97.9 72.1-97.9c66.5 0 74.8 53.3 100.9 134.9c18.8 56.9 54 102.6 155.4 102.6c72.7 0 122-22.3 122-80.9c0-72.9-62.7-80.6-115-92.1c-25.8-5.9-33.4-16.4-33.4-34c0-19.9 15.8-31.7 41.6-31.7c28.2 0 43.4 10.6 45.7 35.8l58.6-7c-4.7-52.8-41.1-74.5-100.9-74.5c-52.8 0-104.4 19.9-104.4 83.9c0 39.9 19.4 65.1 68 76.8c44.9 10.6 79.8 13.8 79.8 45.7c0 21.7-21.1 30.5-61 30.5c-59.2 0-83.9-31.1-97.9-73.9c-32-96.8-43.6-163-161.3-163C45.7 113.8 0 168.3 0 261c0 89.1 45.7 137.2 127.9 137.2c66.2 0 97.9-31.1 97.9-31.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LastfmSquare(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 32H48C21.5 32 0 53.5 0 80v352c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V80c0-26.5-21.5-48-48-48m-92.2 312.9c-63.4 0-85.4-28.6-97.1-64.1c-16.3-51-21.5-84.3-63-84.3c-22.4 0-45.1 16.1-45.1 61.2c0 35.2 18 57.2 43.3 57.2c28.6 0 47.6-21.3 47.6-21.3l11.7 31.9s-19.8 19.4-61.2 19.4c-51.3 0-79.9-30.1-79.9-85.8c0-57.9 28.6-92 82.5-92c73.5 0 80.8 41.4 100.8 101.9c8.8 26.8 24.2 46.2 61.2 46.2c24.9 0 38.1-5.5 38.1-19.1c0-19.9-21.8-22-49.9-28.6c-30.4-7.3-42.5-23.1-42.5-48c0-40 32.3-52.4 65.2-52.4c37.4 0 60.1 13.6 63 46.6l-36.7 4.4c-1.5-15.8-11-22.4-28.6-22.4c-16.1 0-26 7.3-26 19.8c0 11 4.8 17.6 20.9 21.3c32.7 7.1 71.8 12 71.8 57.5c.1 36.7-30.7 50.6-76.1 50.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Leanpub(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m386.539 111.485l15.096 248.955l-10.979-.275c-36.232-.824-71.64 8.783-102.657 27.997c-31.016-19.214-66.424-27.997-102.657-27.997c-45.564 0-82.07 10.705-123.516 27.723L93.117 129.6c28.546-11.803 61.484-18.115 92.226-18.115c41.173 0 73.836 13.175 102.657 42.544c27.723-28.271 59.013-41.721 98.539-42.544M569.07 448c-25.526 0-47.485-5.215-70.542-15.645c-34.31-15.645-69.993-24.978-107.871-24.978c-38.977 0-74.934 12.901-102.657 40.623c-27.723-27.723-63.68-40.623-102.657-40.623c-37.878 0-73.561 9.333-107.871 24.978C55.239 442.236 32.731 448 8.303 448H6.93L49.475 98.859C88.726 76.626 136.486 64 181.775 64C218.83 64 256.984 71.685 288 93.095C319.016 71.685 357.17 64 394.225 64c45.289 0 93.049 12.626 132.3 34.859zm-43.368-44.741l-34.036-280.246c-30.742-13.999-67.248-21.41-101.009-21.41c-38.428 0-74.385 12.077-102.657 38.702c-28.272-26.625-64.228-38.702-102.657-38.702c-33.761 0-70.267 7.411-101.009 21.41L50.298 403.259c47.211-19.487 82.894-33.486 135.045-33.486c37.604 0 70.817 9.606 102.657 29.644c31.84-20.038 65.052-29.644 102.657-29.644c52.151 0 87.834 13.999 135.045 33.486"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Less(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M612.7 219c0-20.5 3.2-32.6 3.2-54.6c0-34.2-12.6-45.2-40.5-45.2h-20.5v24.2h6.3c14.2 0 17.3 4.7 17.3 22.1c0 16.3-1.6 32.6-1.6 51.5c0 24.2 7.9 33.6 23.6 37.3v1.6c-15.8 3.7-23.6 13.1-23.6 37.3c0 18.9 1.6 34.2 1.6 51.5c0 17.9-3.7 22.6-17.3 22.6v.5h-6.3V393h20.5c27.8 0 40.5-11 40.5-45.2c0-22.6-3.2-34.2-3.2-54.6c0-11 6.8-22.6 27.3-23.6v-27.3c-20.5-.7-27.3-12.3-27.3-23.3m-105.6 32c-15.8-6.3-30.5-10-30.5-20.5c0-7.9 6.3-12.6 17.9-12.6s22.1 4.7 33.6 13.1l21-27.8c-13.1-10-31-20.5-55.2-20.5c-35.7 0-59.9 20.5-59.9 49.4c0 25.7 22.6 38.9 41.5 46.2c16.3 6.3 32.1 11.6 32.1 22.1c0 7.9-6.3 13.1-20.5 13.1c-13.1 0-26.3-5.3-40.5-16.3l-21 30.5c15.8 13.1 39.9 22.1 59.9 22.1c42 0 64.6-22.1 64.6-51s-22.5-41-43-47.8m-358.9 59.4c-3.7 0-8.4-3.2-8.4-13.1V119.1H65.2c-28.4 0-41 11-41 45.2c0 22.6 3.2 35.2 3.2 54.6c0 11-6.8 22.6-27.3 23.6v27.3c20.5.5 27.3 12.1 27.3 23.1c0 19.4-3.2 31-3.2 53.6c0 34.2 12.6 45.2 40.5 45.2h20.5v-24.2h-6.3c-13.1 0-17.3-5.3-17.3-22.6s1.6-32.1 1.6-51.5c0-24.2-7.9-33.6-23.6-37.3v-1.6c15.8-3.7 23.6-13.1 23.6-37.3c0-18.9-1.6-34.2-1.6-51.5s3.7-22.1 17.3-22.1H93v150.8c0 32.1 11 53.1 43.1 53.1c10 0 17.9-1.6 23.6-3.7l-5.3-34.2c-3.1.8-4.6.8-6.2.8M379.9 251c-16.3-6.3-31-10-31-20.5c0-7.9 6.3-12.6 17.9-12.6c11.6 0 22.1 4.7 33.6 13.1l21-27.8c-13.1-10-31-20.5-55.2-20.5c-35.7 0-59.9 20.5-59.9 49.4c0 25.7 22.6 38.9 41.5 46.2c16.3 6.3 32.1 11.6 32.1 22.1c0 7.9-6.3 13.1-20.5 13.1c-13.1 0-26.3-5.3-40.5-16.3l-20.5 30.5c15.8 13.1 39.9 22.1 59.9 22.1c42 0 64.6-22.1 64.6-51c.1-28.9-22.5-41-43-47.8m-155-68.8c-38.4 0-75.1 32.1-74.1 82.5c0 52 34.2 82.5 79.3 82.5c18.9 0 39.9-6.8 56.2-17.9l-15.8-27.8c-11.6 6.8-22.6 10-34.2 10c-21 0-37.3-10-41.5-34.2H290c.5-3.7 1.6-11 1.6-19.4c.6-42.6-22.6-75.7-66.7-75.7m-30 66.2c3.2-21 15.8-31 30.5-31c18.9 0 26.3 13.1 26.3 31z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Line(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M272.1 204.2v71.1c0 1.8-1.4 3.2-3.2 3.2h-11.4c-1.1 0-2.1-.6-2.6-1.3l-32.6-44v42.2c0 1.8-1.4 3.2-3.2 3.2h-11.4c-1.8 0-3.2-1.4-3.2-3.2v-71.1c0-1.8 1.4-3.2 3.2-3.2H219c1 0 2.1.5 2.6 1.4l32.6 44v-42.2c0-1.8 1.4-3.2 3.2-3.2h11.4c1.8-.1 3.3 1.4 3.3 3.1m-82-3.2h-11.4c-1.8 0-3.2 1.4-3.2 3.2v71.1c0 1.8 1.4 3.2 3.2 3.2h11.4c1.8 0 3.2-1.4 3.2-3.2v-71.1c0-1.7-1.4-3.2-3.2-3.2m-27.5 59.6h-31.1v-56.4c0-1.8-1.4-3.2-3.2-3.2h-11.4c-1.8 0-3.2 1.4-3.2 3.2v71.1c0 .9.3 1.6.9 2.2c.6.5 1.3.9 2.2.9h45.7c1.8 0 3.2-1.4 3.2-3.2v-11.4c0-1.7-1.4-3.2-3.1-3.2M332.1 201h-45.7c-1.7 0-3.2 1.4-3.2 3.2v71.1c0 1.7 1.4 3.2 3.2 3.2h45.7c1.8 0 3.2-1.4 3.2-3.2v-11.4c0-1.8-1.4-3.2-3.2-3.2H301v-12h31.1c1.8 0 3.2-1.4 3.2-3.2V234c0-1.8-1.4-3.2-3.2-3.2H301v-12h31.1c1.8 0 3.2-1.4 3.2-3.2v-11.4c-.1-1.7-1.5-3.2-3.2-3.2M448 113.7V399c-.1 44.8-36.8 81.1-81.7 81H81c-44.8-.1-81.1-36.9-81-81.7V113c.1-44.8 36.9-81.1 81.7-81H367c44.8.1 81.1 36.8 81 81.7m-61.6 122.6c0-73-73.2-132.4-163.1-132.4c-89.9 0-163.1 59.4-163.1 132.4c0 65.4 58 120.2 136.4 130.6c19.1 4.1 16.9 11.1 12.6 36.8c-.7 4.1-3.3 16.1 14.1 8.8c17.4-7.3 93.9-55.3 128.2-94.7c23.6-26 34.9-52.3 34.9-81.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linkedin(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M416 32H31.9C14.3 32 0 46.5 0 64.3v383.4C0 465.5 14.3 480 31.9 480H416c17.6 0 32-14.5 32-32.3V64.3c0-17.8-14.4-32.3-32-32.3M135.4 416H69V202.2h66.5V416zm-33.2-243c-21.3 0-38.5-17.3-38.5-38.5S80.9 96 102.2 96c21.2 0 38.5 17.3 38.5 38.5c0 21.3-17.2 38.5-38.5 38.5m282.1 243h-66.4V312c0-24.8-.5-56.7-34.5-56.7c-34.6 0-39.9 27-39.9 54.9V416h-66.4V202.2h63.7v29.2h.9c8.9-16.8 30.6-34.5 62.9-34.5c67.2 0 79.7 44.3 79.7 101.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkedinIn(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M100.28 448H7.4V148.9h92.88zM53.79 108.1C24.09 108.1 0 83.5 0 53.8a53.79 53.79 0 0 1 107.58 0c0 29.7-24.1 54.3-53.79 54.3M447.9 448h-92.68V302.4c0-34.7-.7-79.2-48.29-79.2c-48.29 0-55.69 37.7-55.69 76.7V448h-92.78V148.9h89.08v40.8h1.3c12.4-23.5 42.69-48.3 87.88-48.3c94 0 111.28 61.9 111.28 142.3V448z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linode(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M437.4 226.3c-.3-.9-.9-1.4-1.4-2l-70-38.6c-.9-.6-2-.6-3.1 0l-58.9 36c-.9.6-1.4 1.7-1.4 2.6l-.9 31.4l-24-16c-.9-.6-2.3-.6-3.1 0L240 260.9l-1.4-35.1c0-.9-.6-2-1.4-2.3l-36-24.3l33.7-17.4c1.1-.6 1.7-1.7 1.7-2.9l-5.7-132.3c0-.9-.9-2-1.7-2.6L138.6.3c-.9-.3-1.7-.3-2.3-.3L12.6 38.6c-1.4.6-2.3 2-2 3.7L38 175.4c.9 3.4 34 27.4 38.6 30.9l-26.9 12.9c-1.4.9-2 2.3-1.7 3.4l20.6 100.3c.6 2.9 23.7 23.1 27.1 26.3l-17.4 10.6c-.9.6-1.7 2-1.4 3.1c1.4 7.1 15.4 77.7 16.9 79.1l65.1 69.1c.6.6 1.4.6 2.3.9c.6 0 1.1-.3 1.7-.6l83.7-66.9c.9-.6 1.1-1.4 1.1-2.3l-2-46l28 23.7c1.1.9 2.9.9 4 0l66.9-53.4c.9-.6 1.1-1.4 1.1-2.3l2.3-33.4l20.3 14c1.1.9 2.6.9 3.7 0l54.6-43.7c.6-.3 1.1-1.1 1.1-2c.9-6.5 10.3-70.8 9.7-72.8m-204.8 4.8l4 92.6l-90.6 61.2l-14-96.6zm-7.7-180l5.4 126l-106.6 55.4L104 97.7zM44 173.1L18 48l79.7 49.4l19.4 132.9zm30.6 147.8L55.7 230l70 58.3l13.7 93.4zm24.3 117.7l-13.7-67.1l61.7 60.9l9.7 67.4zm64.5 64.5l-10.6-70.9l85.7-61.4l3.1 70zm82-115.1c0-3.4.9-22.9-2-25.1l-24.3-20l22.3-14.9c2.3-1.7 1.1-5.7 1.1-8l29.4 22.6l.6 68.3zm94.3-25.4l-60.9 48.6l-.6-68.6l65.7-46.9zm27.7-25.7l-19.1-13.4l2-34c.3-.9-.3-2-1.1-2.6L308 259.7l.6-30l64.6 40.6zm54.6-39.8l-48.3 38.3l5.7-65.1l51.1-36.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linux(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M220.8 123.3c1 .5 1.8 1.7 3 1.7c1.1 0 2.8-.4 2.9-1.5c.2-1.4-1.9-2.3-3.2-2.9c-1.7-.7-3.9-1-5.5-.1c-.4.2-.8.7-.6 1.1c.3 1.3 2.3 1.1 3.4 1.7m-21.9 1.7c1.2 0 2-1.2 3-1.7c1.1-.6 3.1-.4 3.5-1.6c.2-.4-.2-.9-.6-1.1c-1.6-.9-3.8-.6-5.5.1c-1.3.6-3.4 1.5-3.2 2.9c.1 1 1.8 1.5 2.8 1.4M420 403.8c-3.6-4-5.3-11.6-7.2-19.7c-1.8-8.1-3.9-16.8-10.5-22.4c-1.3-1.1-2.6-2.1-4-2.9c-1.3-.8-2.7-1.5-4.1-2c9.2-27.3 5.6-54.5-3.7-79.1c-11.4-30.1-31.3-56.4-46.5-74.4c-17.1-21.5-33.7-41.9-33.4-72C311.1 85.4 315.7.1 234.8 0C132.4-.2 158 103.4 156.9 135.2c-1.7 23.4-6.4 41.8-22.5 64.7c-18.9 22.5-45.5 58.8-58.1 96.7c-6 17.9-8.8 36.1-6.2 53.3c-6.5 5.8-11.4 14.7-16.6 20.2c-4.2 4.3-10.3 5.9-17 8.3s-14 6-18.5 14.5c-2.1 3.9-2.8 8.1-2.8 12.4c0 3.9.6 7.9 1.2 11.8c1.2 8.1 2.5 15.7.8 20.8c-5.2 14.4-5.9 24.4-2.2 31.7c3.8 7.3 11.4 10.5 20.1 12.3c17.3 3.6 40.8 2.7 59.3 12.5c19.8 10.4 39.9 14.1 55.9 10.4c11.6-2.6 21.1-9.6 25.9-20.2c12.5-.1 26.3-5.4 48.3-6.6c14.9-1.2 33.6 5.3 55.1 4.1c.6 2.3 1.4 4.6 2.5 6.7v.1c8.3 16.7 23.8 24.3 40.3 23c16.6-1.3 34.1-11 48.3-27.9c13.6-16.4 36-23.2 50.9-32.2c7.4-4.5 13.4-10.1 13.9-18.3c.4-8.2-4.4-17.3-15.5-29.7M223.7 87.3c9.8-22.2 34.2-21.8 44-.4c6.5 14.2 3.6 30.9-4.3 40.4c-1.6-.8-5.9-2.6-12.6-4.9c1.1-1.2 3.1-2.7 3.9-4.6c4.8-11.8-.2-27-9.1-27.3c-7.3-.5-13.9 10.8-11.8 23c-4.1-2-9.4-3.5-13-4.4c-1-6.9-.3-14.6 2.9-21.8M183 75.8c10.1 0 20.8 14.2 19.1 33.5c-3.5 1-7.1 2.5-10.2 4.6c1.2-8.9-3.3-20.1-9.6-19.6c-8.4.7-9.8 21.2-1.8 28.1c1 .8 1.9-.2-5.9 5.5c-15.6-14.6-10.5-52.1 8.4-52.1m-13.6 60.7c6.2-4.6 13.6-10 14.1-10.5c4.7-4.4 13.5-14.2 27.9-14.2c7.1 0 15.6 2.3 25.9 8.9c6.3 4.1 11.3 4.4 22.6 9.3c8.4 3.5 13.7 9.7 10.5 18.2c-2.6 7.1-11 14.4-22.7 18.1c-11.1 3.6-19.8 16-38.2 14.9c-3.9-.2-7-1-9.6-2.1c-8-3.5-12.2-10.4-20-15c-8.6-4.8-13.2-10.4-14.7-15.3c-1.4-4.9 0-9 4.2-12.3m3.3 334c-2.7 35.1-43.9 34.4-75.3 18c-29.9-15.8-68.6-6.5-76.5-21.9c-2.4-4.7-2.4-12.7 2.6-26.4v-.2c2.4-7.6.6-16-.6-23.9c-1.2-7.8-1.8-15 .9-20c3.5-6.7 8.5-9.1 14.8-11.3c10.3-3.7 11.8-3.4 19.6-9.9c5.5-5.7 9.5-12.9 14.3-18c5.1-5.5 10-8.1 17.7-6.9c8.1 1.2 15.1 6.8 21.9 16l19.6 35.6c9.5 19.9 43.1 48.4 41 68.9m-1.4-25.9c-4.1-6.6-9.6-13.6-14.4-19.6c7.1 0 14.2-2.2 16.7-8.9c2.3-6.2 0-14.9-7.4-24.9c-13.5-18.2-38.3-32.5-38.3-32.5c-13.5-8.4-21.1-18.7-24.6-29.9s-3-23.3-.3-35.2c5.2-22.9 18.6-45.2 27.2-59.2c2.3-1.7.8 3.2-8.7 20.8c-8.5 16.1-24.4 53.3-2.6 82.4c.6-20.7 5.5-41.8 13.8-61.5c12-27.4 37.3-74.9 39.3-112.7c1.1.8 4.6 3.2 6.2 4.1c4.6 2.7 8.1 6.7 12.6 10.3c12.4 10 28.5 9.2 42.4 1.2c6.2-3.5 11.2-7.5 15.9-9c9.9-3.1 17.8-8.6 22.3-15c7.7 30.4 25.7 74.3 37.2 95.7c6.1 11.4 18.3 35.5 23.6 64.6c3.3-.1 7 .4 10.9 1.4c13.8-35.7-11.7-74.2-23.3-84.9c-4.7-4.6-4.9-6.6-2.6-6.5c12.6 11.2 29.2 33.7 35.2 59c2.8 11.6 3.3 23.7.4 35.7c16.4 6.8 35.9 17.9 30.7 34.8c-2.2-.1-3.2 0-4.2 0c3.2-10.1-3.9-17.6-22.8-26.1c-19.6-8.6-36-8.6-38.3 12.5c-12.1 4.2-18.3 14.7-21.4 27.3c-2.8 11.2-3.6 24.7-4.4 39.9c-.5 7.7-3.6 18-6.8 29c-32.1 22.9-76.7 32.9-114.3 7.2m257.4-11.5c-.9 16.8-41.2 19.9-63.2 46.5c-13.2 15.7-29.4 24.4-43.6 25.5s-26.5-4.8-33.7-19.3c-4.7-11.1-2.4-23.1 1.1-36.3c3.7-14.2 9.2-28.8 9.9-40.6c.8-15.2 1.7-28.5 4.2-38.7c2.6-10.3 6.6-17.2 13.7-21.1c.3-.2.7-.3 1-.5c.8 13.2 7.3 26.6 18.8 29.5c12.6 3.3 30.7-7.5 38.4-16.3c9-.3 15.7-.9 22.6 5.1c9.9 8.5 7.1 30.3 17.1 41.6c10.6 11.6 14 19.5 13.7 24.6M173.3 148.7c2 1.9 4.7 4.5 8 7.1c6.6 5.2 15.8 10.6 27.3 10.6c11.6 0 22.5-5.9 31.8-10.8c4.9-2.6 10.9-7 14.8-10.4s5.9-6.3 3.1-6.6s-2.6 2.6-6 5.1c-4.4 3.2-9.7 7.4-13.9 9.8c-7.4 4.2-19.5 10.2-29.9 10.2s-18.7-4.8-24.9-9.7c-3.1-2.5-5.7-5-7.7-6.9c-1.5-1.4-1.9-4.6-4.3-4.9c-1.4-.1-1.8 3.7 1.7 6.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lyft(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 81.1h77.8v208.7c0 33.1 15 52.8 27.2 61c-12.7 11.1-51.2 20.9-80.2-2.8C7.8 334 0 310.7 0 289zm485.9 173.5v-22h23.8v-76.8h-26.1c-10.1-46.3-51.2-80.7-100.3-80.7c-56.6 0-102.7 46-102.7 102.7V357c16 2.3 35.4-.3 51.7-14c17.1-14 24.8-37.2 24.8-59v-6.7h38.8v-76.8h-38.8v-23.3c0-34.6 52.2-34.6 52.2 0v77.1c0 56.6 46 102.7 102.7 102.7v-76.5c-14.5 0-26.1-11.7-26.1-25.9m-294.3-99v113c0 15.4-23.8 15.4-23.8 0v-113H91v132.7c0 23.8 8 54 45 63.9c37 9.8 58.2-10.6 58.2-10.6c-2.1 13.4-14.5 23.3-34.9 25.3c-15.5 1.6-35.2-3.6-45-7.8v70.3c25.1 7.5 51.5 9.8 77.6 4.7c47.1-9.1 76.8-48.4 76.8-100.8V155.1h-77.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Magento(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M445.7 127.9V384l-63.4 36.5V164.7L223.8 73.1L65.2 164.7l.4 255.9L2.3 384V128.1L224.2 0zM255.6 420.5L224 438.9l-31.8-18.2v-256l-63.3 36.6l.1 255.9l94.9 54.9l95.1-54.9v-256l-63.4-36.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mailchimp(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M330.61 243.52a36.15 36.15 0 0 1 9.3 0c1.66-3.83 1.95-10.43.45-17.61c-2.23-10.67-5.25-17.14-11.48-16.13s-6.47 8.74-4.24 19.42c1.26 6 3.49 11.14 6 14.32zM277.05 252c4.47 2 7.2 3.26 8.28 2.13c1.89-1.94-3.48-9.39-12.12-13.09a31.44 31.44 0 0 0-30.61 3.68c-3 2.18-5.81 5.22-5.41 7.06c.85 3.74 10-2.71 22.6-3.48c7-.44 12.8 1.75 17.26 3.71zm-9 5.13c-9.07 1.42-15 6.53-13.47 10.1c.9.34 1.17.81 5.21-.81a37 37 0 0 1 18.72-1.95c2.92.34 4.31.52 4.94-.49c1.46-2.22-5.71-8-15.39-6.85zm54.17 17.1c3.38-6.87-10.9-13.93-14.3-7s10.92 13.88 14.32 6.97zm15.66-20.47c-7.66-.13-7.95 15.8-.26 15.93s7.98-15.81.28-15.96zm-218.79 78.9c-1.32.31-6 1.45-8.47-2.35c-5.2-8 11.11-20.38 3-35.77c-9.1-17.47-27.82-13.54-35.05-5.54c-8.71 9.6-8.72 23.54-5 24.08c4.27.57 4.08-6.47 7.38-11.63a12.83 12.83 0 0 1 17.85-3.72c11.59 7.59 1.37 17.76 2.28 28.62c1.39 16.68 18.42 16.37 21.58 9a2.08 2.08 0 0 0-.2-2.33c.03.89.68-1.3-3.35-.39zm299.72-17.07c-3.35-11.73-2.57-9.22-6.78-20.52c2.45-3.67 15.29-24-3.07-43.25c-10.4-10.92-33.9-16.54-41.1-18.54c-1.5-11.39 4.65-58.7-21.52-83c20.79-21.55 33.76-45.29 33.73-65.65c-.06-39.16-48.15-51-107.42-26.47l-12.55 5.33c-.06-.05-22.71-22.27-23.05-22.57C169.5-18-41.77 216.81 25.78 273.85l14.76 12.51a72.49 72.49 0 0 0-4.1 33.5c3.36 33.4 36 60.42 67.53 60.38c57.73 133.06 267.9 133.28 322.29 3c1.74-4.47 9.11-24.61 9.11-42.38s-10.09-25.27-16.53-25.27zm-316 48.16c-22.82-.61-47.46-21.15-49.91-45.51c-6.17-61.31 74.26-75.27 84-12.33c4.54 29.64-4.67 58.49-34.12 57.81zM84.3 249.55C69.14 252.5 55.78 261.09 47.6 273c-4.88-4.07-14-12-15.59-15c-13.01-24.85 14.24-73 33.3-100.21C112.42 90.56 186.19 39.68 220.36 48.91c5.55 1.57 23.94 22.89 23.94 22.89s-34.15 18.94-65.8 45.35c-42.66 32.85-74.89 80.59-94.2 132.4M323.18 350.7s-35.74 5.3-69.51-7.07c6.21-20.16 27 6.1 96.4-13.81c15.29-4.38 35.37-13 51-25.35a102.85 102.85 0 0 1 7.12 24.28c3.66-.66 14.25-.52 11.44 18.1c-3.29 19.87-11.73 36-25.93 50.84A106.86 106.86 0 0 1 362.55 421a132.45 132.45 0 0 1-20.34 8.58c-53.51 17.48-108.3-1.74-126-43a66.33 66.33 0 0 1-3.55-9.74c-7.53-27.2-1.14-59.83 18.84-80.37c1.23-1.31 2.48-2.85 2.48-4.79a8.45 8.45 0 0 0-1.92-4.54c-7-10.13-31.19-27.4-26.33-60.83c3.5-24 24.49-40.91 44.07-39.91l5 .29c8.48.5 15.89 1.59 22.88 1.88c11.69.5 22.2-1.19 34.64-11.56c4.2-3.5 7.57-6.54 13.26-7.51a17.45 17.45 0 0 1 13.6 2.24c10 6.64 11.4 22.73 11.92 34.49c.29 6.72 1.1 23 1.38 27.63c.63 10.67 3.43 12.17 9.11 14c3.19 1.05 6.15 1.83 10.51 3.06c13.21 3.71 21 7.48 26 12.31a16.38 16.38 0 0 1 4.74 9.29c1.56 11.37-8.82 25.4-36.31 38.16c-46.71 21.68-93.68 14.45-100.48 13.68c-20.15-2.71-31.63 23.32-19.55 41.15c22.64 33.41 122.4 20 151.37-21.35c.69-1 .12-1.59-.73-1c-41.77 28.58-97.06 38.21-128.46 26c-4.77-1.85-14.73-6.44-15.94-16.67c43.6 13.49 71 .74 71 .74s2.03-2.79-.56-2.53M171.31 157.5c16.74-19.35 37.36-36.18 55.83-45.63a.73.73 0 0 1 1 1c-1.46 2.66-4.29 8.34-5.19 12.65a.75.75 0 0 0 1.16.79c11.49-7.83 31.48-16.22 49-17.3a.77.77 0 0 1 .52 1.38a41.86 41.86 0 0 0-7.71 7.74a.75.75 0 0 0 .59 1.19c12.31.09 29.66 4.4 41 10.74c.76.43.22 1.91-.64 1.72c-69.55-15.94-123.08 18.53-134.5 26.83a.76.76 0 0 1-1-1.12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mandalorian(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M232.27 511.89c-1-3.26-1.69-15.83-1.39-24.58c.55-15.89 1-24.72 1.4-28.76c.64-6.2 2.87-20.72 3.28-21.38c.6-1 .4-27.87-.24-33.13c-.31-2.58-.63-11.9-.69-20.73c-.13-16.47-.53-20.12-2.73-24.76c-1.1-2.32-1.23-3.84-1-11.43a92.38 92.38 0 0 0-.34-12.71c-2-13-3.46-27.7-3.25-33.9s.43-7.15 2.06-9.67c3.05-4.71 6.51-14 8.62-23.27c2.26-9.86 3.88-17.18 4.59-20.74a109.54 109.54 0 0 1 4.42-15.05c2.27-6.25 2.49-15.39.37-15.39c-.3 0-1.38 1.22-2.41 2.71s-4.76 4.8-8.29 7.36c-8.37 6.08-11.7 9.39-12.66 12.58s-1 7.23-.16 7.76c.34.21 1.29 2.4 2.11 4.88a28.83 28.83 0 0 1 .72 15.36c-.39 1.77-1 5.47-1.46 8.23s-1 6.46-1.25 8.22a9.85 9.85 0 0 1-1.55 4.26c-1 1-1.14.91-2.05-.53a14.87 14.87 0 0 1-1.44-4.75c-.25-1.74-1.63-7.11-3.08-11.93c-3.28-10.9-3.52-16.15-1-21a14.24 14.24 0 0 0 1.67-4.61c0-2.39-2.2-5.32-7.41-9.89c-7-6.18-8.63-7.92-10.23-11.3c-1.71-3.6-3.06-4.06-4.54-1.54c-1.78 3-2.6 9.11-3 22l-.34 12.19l2 2.25c3.21 3.7 12.07 16.45 13.78 19.83c3.41 6.74 4.34 11.69 4.41 23.56s.95 22.75 2 24.71c.36.66.51 1.35.34 1.52s.41 2.09 1.29 4.27a38.14 38.14 0 0 1 2.06 9a91 91 0 0 0 1.71 10.37c2.23 9.56 2.77 14.08 2.39 20.14c-.2 3.27-.53 11.07-.73 17.32c-1.31 41.76-1.85 58-2 61.21c-.12 2-.39 11.51-.6 21.07c-.36 16.3-1.3 27.37-2.42 28.65c-.64.73-8.07-4.91-12.52-9.49c-3.75-3.87-4-4.79-2.83-9.95c.7-3 2.26-18.29 3.33-32.62c.36-4.78.81-10.5 1-12.71c.83-9.37 1.66-20.35 2.61-34.78c.56-8.46 1.33-16.44 1.72-17.73s.89-9.89 1.13-19.11l.43-16.77l-2.26-4.3c-1.72-3.28-4.87-6.94-13.22-15.34c-6-6.07-11.84-12.3-12.91-13.85l-1.95-2.81l.75-10.9c1.09-15.71 1.1-48.57 0-59.06l-.89-8.7l-3.28-4.52c-5.86-8.08-5.8-7.75-6.22-33.27c-.1-6.07-.38-11.5-.63-12.06c-.83-1.87-3.05-2.66-8.54-3.05c-8.86-.62-11-1.9-23.85-14.55c-6.15-6-12.34-12-13.75-13.19c-2.81-2.42-2.79-2-.56-9.63l1.35-4.65l-1.69-3a32.22 32.22 0 0 0-2.59-4.07c-1.33-1.51-5.5-10.89-6-13.49a4.24 4.24 0 0 1 .87-3.9c2.23-2.86 3.4-5.68 4.45-10.73c2.33-11.19 7.74-26.09 10.6-29.22c3.18-3.47 7.7-1 9.41 5c1.34 4.79 1.37 9.79.1 18.55a101.2 101.2 0 0 0-1 11.11c0 4 .19 4.69 2.25 7.39c3.33 4.37 7.73 7.41 15.2 10.52a18.67 18.67 0 0 1 4.72 2.85c11.17 10.72 18.62 16.18 22.95 16.85c5.18.8 8 4.54 10 13.39c1.31 5.65 4 11.14 5.46 11.14a9.38 9.38 0 0 0 3.33-1.39c2-1.22 2.25-1.73 2.25-4.18a132.88 132.88 0 0 0-2-17.84c-.37-1.66-.78-4.06-.93-5.35s-.61-3.85-1-5.69c-2.55-11.16-3.65-15.46-4.1-16c-1.55-2-4.08-10.2-4.93-15.92c-1.64-11.11-4-14.23-12.91-17.39A43.15 43.15 0 0 1 165.24 78c-1.15-1-4-3.22-6.35-5.06s-4.41-3.53-4.6-3.76a22.7 22.7 0 0 0-2.69-2c-6.24-4.22-8.84-7-11.26-12l-2.44-5l-.22-13l-.22-13l6.91-6.55c3.95-3.75 8.48-7.35 10.59-8.43c3.31-1.69 4.45-1.89 11.37-2c8.53-.19 10.12 0 11.66 1.56s1.36 6.4-.29 8.5a6.66 6.66 0 0 0-1.34 2.32c0 .58-2.61 4.91-5.42 9a30.39 30.39 0 0 0-2.37 6.82c20.44 13.39 21.55 3.77 14.07 29L194 66.92c3.11-8.66 6.47-17.26 8.61-26.22c.29-7.63-12-4.19-15.4-8.68c-2.33-5.93 3.13-14.18 6.06-19.2c1.6-2.34 6.62-4.7 8.82-4.15c.88.22 4.16-.35 7.37-1.28a45.3 45.3 0 0 1 7.55-1.68a29.57 29.57 0 0 0 6-1.29c3.65-1.11 4.5-1.17 6.35-.4a29.54 29.54 0 0 0 5.82 1.36a18.18 18.18 0 0 1 6 1.91a22.67 22.67 0 0 0 5 2.17c2.51.68 3 .57 7.05-1.67l4.35-2.4L268.32 5c10.44-.4 10.81-.47 15.26-2.68L288.16 0l2.46 1.43c1.76 1 3.14 2.73 4.85 6c2.36 4.51 2.38 4.58 1.37 7.37c-.88 2.44-.89 3.3-.1 6.39a35.76 35.76 0 0 0 2.1 5.91a13.55 13.55 0 0 1 1.31 4c.31 4.33 0 5.3-2.41 6.92c-2.17 1.47-7 7.91-7 9.34a14.77 14.77 0 0 1-1.07 3c-5 11.51-6.76 13.56-14.26 17c-9.2 4.2-12.3 5.19-16.21 5.19c-3.1 0-4 .25-4.54 1.26a18.33 18.33 0 0 1-4.09 3.71a13.62 13.62 0 0 0-4.38 4.78a5.89 5.89 0 0 1-2.49 2.91a6.88 6.88 0 0 0-2.45 1.71a67.62 67.62 0 0 1-7 5.38c-3.33 2.34-6.87 5-7.87 6A7.27 7.27 0 0 1 224 100a5.76 5.76 0 0 0-2.13 1.65c-1.31 1.39-1.49 2.11-1.14 4.6a36.45 36.45 0 0 0 1.42 5.88c1.32 3.8 1.31 7.86 0 10.57s-.89 6.65 1.35 9.59c2 2.63 2.16 4.56.71 8.84a33.45 33.45 0 0 0-1.06 8.91c0 4.88.22 6.28 1.46 8.38s1.82 2.48 3.24 2.32c2-.23 2.3-1.05 4.71-12.12c2.18-10 3.71-11.92 13.76-17.08c2.94-1.51 7.46-4 10-5.44s6.79-3.69 9.37-4.91a40.09 40.09 0 0 0 15.22-11.67c7.11-8.79 10-16.22 12.85-33.3a18.37 18.37 0 0 1 2.86-7.73a20.39 20.39 0 0 0 2.89-7.31c1-5.3 2.85-9.08 5.58-11.51c4.7-4.18 6-1.09 4.59 10.87c-.46 3.86-1.1 10.33-1.44 14.38l-.61 7.36l4.45 4.09l4.45 4.09l.11 8.42c.06 4.63.47 9.53.92 10.89l.82 2.47l-6.43 6.28c-8.54 8.33-12.88 13.93-16.76 21.61c-1.77 3.49-3.74 7.11-4.38 8c-2.18 3.11-6.46 13-8.76 20.26l-2.29 7.22l-7 6.49c-3.83 3.57-8 7.25-9.17 8.17c-3.05 2.32-4.26 5.15-4.26 10a14.62 14.62 0 0 0 1.59 7.26a42 42 0 0 1 2.09 4.83a9.28 9.28 0 0 0 1.57 2.89c1.4 1.59 1.92 16.12.83 23.22c-.68 4.48-3.63 12-4.7 12c-1.79 0-4.06 9.27-5.07 20.74c-.18 2-.62 5.94-1 8.7s-1 10-1.35 16.05c-.77 12.22-.19 18.77 2 23.15c3.41 6.69.52 12.69-11 22.84l-4 3.49l.07 5.19a40.81 40.81 0 0 0 1.14 8.87c4.61 16 4.73 16.92 4.38 37.13c-.46 26.4-.26 40.27.63 44.15a61.31 61.31 0 0 1 1.08 7c.17 2 .66 5.33 1.08 7.36c.47 2.26.78 11 .79 22.74v19.06l-1.81 2.63c-2.71 3.91-15.11 13.54-15.49 12.29zm29.53-45.11c-.18-.3-.33-6.87-.33-14.59c0-14.06-.89-27.54-2.26-34.45c-.4-2-.81-9.7-.9-17.06c-.15-11.93-1.4-24.37-2.64-26.38c-.66-1.07-3-17.66-3-21.3c0-4.23 1-6 5.28-9.13s4.86-3.14 5.48-.72c.28 1.1 1.45 5.62 2.6 10c3.93 15.12 4.14 16.27 4.05 21.74c-.1 5.78-.13 6.13-1.74 17.73c-1 7.07-1.17 12.39-1 28.43c.17 19.4-.64 35.73-2 41.27c-.71 2.78-2.8 5.48-3.43 4.43zm-71-37.58a101 101 0 0 1-1.73-10.79a100.5 100.5 0 0 0-1.73-10.79a37.53 37.53 0 0 1-1-6.49c-.31-3.19-.91-7.46-1.33-9.48c-1-4.79-3.35-19.35-3.42-21.07c0-.74-.34-4.05-.7-7.36c-.67-6.21-.84-27.67-.22-28.29c1-1 6.63 2.76 11.33 7.43l5.28 5.25l-.45 6.47c-.25 3.56-.6 10.23-.78 14.83s-.49 9.87-.67 11.71s-.61 9.36-.94 16.72c-.79 17.41-1.94 31.29-2.65 32a.62.62 0 0 1-1-.14zm-87.18-266.59c21.07 12.79 17.84 14.15 28.49 17.66c13 4.29 18.87 7.13 23.15 16.87C111.6 233.28 86.25 255 78.55 268c-31 52-6 101.59 62.75 87.21c-14.18 29.23-78 28.63-98.68-4.9c-24.68-39.95-22.09-118.3 61-187.66zm210.79 179c56.66 6.88 82.32-37.74 46.54-89.23c0 0-26.87-29.34-64.28-68c3-15.45 9.49-32.12 30.57-53.82c89.2 63.51 92 141.61 92.46 149.36c4.3 70.64-78.7 91.18-105.29 61.71z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Markdown(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M593.8 59.1H46.2C20.7 59.1 0 79.8 0 105.2v301.5c0 25.5 20.7 46.2 46.2 46.2h547.7c25.5 0 46.2-20.7 46.1-46.1V105.2c0-25.4-20.7-46.1-46.2-46.1M338.5 360.6H277v-120l-61.5 76.9l-61.5-76.9v120H92.3V151.4h61.5l61.5 76.9l61.5-76.9h61.5v209.2zm135.3 3.1L381.5 256H443V151.4h61.5V256H566z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mastodon(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M433 179.11c0-97.2-63.71-125.7-63.71-125.7c-62.52-28.7-228.56-28.4-290.48 0c0 0-63.72 28.5-63.72 125.7c0 115.7-6.6 259.4 105.63 289.1c40.51 10.7 75.32 13 103.33 11.4c50.81-2.8 79.32-18.1 79.32-18.1l-1.7-36.9s-36.31 11.4-77.12 10.1c-40.41-1.4-83-4.4-89.63-54a102.54 102.54 0 0 1-.9-13.9c85.63 20.9 158.65 9.1 178.75 6.7c56.12-6.7 105-41.3 111.23-72.9c9.8-49.8 9-121.5 9-121.5m-75.12 125.2h-46.63v-114.2c0-49.7-64-51.6-64 6.9v62.5h-46.33V197c0-58.5-64-56.6-64-6.9v114.2H90.19c0-122.1-5.2-147.9 18.41-175c25.9-28.9 79.82-30.8 103.83 6.1l11.6 19.5l11.6-19.5c24.11-37.1 78.12-34.8 103.83-6.1c23.71 27.3 18.4 53 18.4 175z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Maxcdn(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M461.1 442.7h-97.4L415.6 200c2.3-10.2.9-19.5-4.4-25.7c-5-6.1-13.7-9.6-24.2-9.6h-49.3l-59.5 278h-97.4l59.5-278h-83.4l-59.5 278H0l59.5-278l-44.6-95.4H387c39.4 0 75.3 16.3 98.3 44.9c23.3 28.6 31.8 67.4 23.6 105.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mdb(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.37 160.41L7 352h43.91l5.59-79.83L84.43 352h44.71l25.54-77.43l4.79 77.43H205l-12.79-191.59H146.7L106 277.74L63.67 160.41zm281 0h-47.9V352h47.9s95 .8 94.2-95.79c-.78-94.21-94.18-95.78-94.18-95.78zm-1.2 146.46V204.78s46 4.27 46.8 50.57s-46.78 51.54-46.78 51.54zm238.29-74.24a56.16 56.16 0 0 0 8-38.31c-5.34-35.76-55.08-34.32-55.08-34.32h-51.9v191.58H482s87 4.79 87-63.85c0-43.14-33.52-55.08-33.52-55.08zm-51.9-31.94s13.57-1.59 16 9.59c1.43 6.66-4 12-4 12h-12v-21.57zm-.1 109.46l.1-24.92V267h.08s41.58-4.73 41.19 22.43c-.33 25.65-41.35 20.74-41.35 20.74z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Medapps(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M118.3 238.4c3.5-12.5 6.9-33.6 13.2-33.6c8.3 1.8 9.6 23.4 18.6 36.6c4.6-23.5 5.3-85.1 14.1-86.7c9-.7 19.7 66.5 22 77.5c9.9 4.1 48.9 6.6 48.9 6.6c1.9 7.3-24 7.6-40 7.8c-4.6 14.8-5.4 27.7-11.4 28c-4.7.2-8.2-28.8-17.5-49.6l-9.4 65.5c-4.4 13-15.5-22.5-21.9-39.3c-3.3-.1-62.4-1.6-47.6-7.8zM228 448c21.2 0 21.2-32 0-32H92c-21.2 0-21.2 32 0 32zm-24 64c21.2 0 21.2-32 0-32h-88c-21.2 0-21.2 32 0 32zm34.2-141.5c3.2-18.9 5.2-36.4 11.9-48.8c7.9-14.7 16.1-28.1 24-41c24.6-40.4 45.9-75.2 45.9-125.5C320 69.6 248.2 0 160 0S0 69.6 0 155.2c0 50.2 21.3 85.1 45.9 125.5c7.9 12.9 16 26.3 24 41c6.7 12.5 8.7 29.8 11.9 48.9c3.5 21 36.1 15.7 32.6-5.1c-3.6-21.7-5.6-40.7-15.3-58.6C66.5 246.5 33 211.3 33 155.2C33 87.3 90 32 160 32s127 55.3 127 123.2c0 56.1-33.5 91.3-66.1 151.6c-9.7 18-11.7 37.4-15.3 58.6c-3.4 20.6 29 26.4 32.6 5.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Medium(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 32v448h448V32zm372.2 106.1l-24 23c-2.1 1.6-3.1 4.2-2.7 6.7v169.3c-.4 2.6.6 5.2 2.7 6.7l23.5 23v5.1h-118V367l24.3-23.6c2.4-2.4 2.4-3.1 2.4-6.7V199.8l-67.6 171.6h-9.1L125 199.8v115c-.7 4.8 1 9.7 4.4 13.2l31.6 38.3v5.1H71.2v-5.1l31.6-38.3c3.4-3.5 4.9-8.4 4.1-13.2v-133c.4-3.7-1-7.3-3.8-9.8L75 138.1V133h87.3l67.4 148L289 133.1h83.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediumM(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M71.5 142.3c.6-5.9-1.7-11.8-6.1-15.8L20.3 72.1V64h140.2l108.4 237.7L364.2 64h133.7v8.1l-38.6 37c-3.3 2.5-5 6.7-4.3 10.8v272c-.7 4.1 1 8.3 4.3 10.8l37.7 37v8.1H307.3v-8.1l39.1-37.9c3.8-3.8 3.8-5 3.8-10.8V171.2L241.5 447.1h-14.7L100.4 171.2v184.9c-1.1 7.8 1.5 15.6 7 21.2l50.8 61.6v8.1h-144v-8L65 377.3c5.4-5.6 7.9-13.5 6.5-21.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Medrt(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M113.7 256c0 121.8 83.9 222.8 193.5 241.1c-18.7 4.5-38.2 6.9-58.2 6.9C111.4 504 0 393 0 256S111.4 8 248.9 8c20.1 0 39.6 2.4 58.2 6.9C197.5 33.2 113.7 134.2 113.7 256m297.4 100.3c-77.7 55.4-179.6 47.5-240.4-14.6c5.5 14.1 12.7 27.7 21.7 40.5c61.6 88.2 182.4 109.3 269.7 47c87.3-62.3 108.1-184.3 46.5-272.6c-9-12.9-19.3-24.3-30.5-34.2c37.4 78.8 10.7 178.5-67 233.9m-218.8-244c-1.4 1-2.7 2.1-4 3.1c64.3-17.8 135.9 4 178.9 60.5c35.7 47 42.9 106.6 24.4 158c56.7-56.2 67.6-142.1 22.3-201.8c-50-65.5-149.1-74.4-221.6-19.8M296 224c-4.4 0-8-3.6-8-8v-40c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v40c0 4.4-3.6 8-8 8h-40c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h40c4.4 0 8 3.6 8 8v40c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8v-40c0-4.4 3.6-8 8-8h40c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Meetup(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M99 414.3c1.1 5.7-2.3 11.1-8 12.3c-5.4 1.1-10.9-2.3-12-8c-1.1-5.4 2.3-11.1 7.7-12.3c5.4-1.2 11.1 2.3 12.3 8m143.1 71.4c-6.3 4.6-8 13.4-3.7 20c4.6 6.6 13.4 8.3 20 3.7c6.3-4.6 8-13.4 3.4-20c-4.2-6.5-13.1-8.3-19.7-3.7m-86-462.3c6.3-1.4 10.3-7.7 8.9-14c-1.1-6.6-7.4-10.6-13.7-9.1c-6.3 1.4-10.3 7.7-9.1 14c1.4 6.6 7.6 10.6 13.9 9.1M34.4 226.3c-10-6.9-23.7-4.3-30.6 6c-6.9 10-4.3 24 5.7 30.9c10 7.1 23.7 4.6 30.6-5.7c6.9-10.4 4.3-24.1-5.7-31.2m272-170.9c10.6-6.3 13.7-20 7.7-30.3c-6.3-10.6-19.7-14-30-7.7s-13.7 20-7.4 30.6c6 10.3 19.4 13.7 29.7 7.4m-191.1 58c7.7-5.4 9.4-16 4.3-23.7s-15.7-9.4-23.1-4.3c-7.7 5.4-9.4 16-4.3 23.7c5.1 7.8 15.6 9.5 23.1 4.3m372.3 156c-7.4 1.7-12.3 9.1-10.6 16.9c1.4 7.4 8.9 12.3 16.3 10.6c7.4-1.4 12.3-8.9 10.6-16.6c-1.5-7.4-8.9-12.3-16.3-10.9m39.7-56.8c-1.1-5.7-6.6-9.1-12-8c-5.7 1.1-9.1 6.9-8 12.6c1.1 5.4 6.6 9.1 12.3 8c5.4-1.5 9.1-6.9 7.7-12.6M447 138.9c-8.6 6-10.6 17.7-4.9 26.3c5.7 8.6 17.4 10.6 26 4.9c8.3-6 10.3-17.7 4.6-26.3c-5.7-8.7-17.4-10.9-25.7-4.9m-6.3 139.4c26.3 43.1 15.1 100-26.3 129.1c-17.4 12.3-37.1 17.7-56.9 17.1c-12 47.1-69.4 64.6-105.1 32.6c-1.1.9-2.6 1.7-3.7 2.9c-39.1 27.1-92.3 17.4-119.4-22.3c-9.7-14.3-14.6-30.6-15.1-46.9c-65.4-10.9-90-94-41.1-139.7c-28.3-46.9.6-107.4 53.4-114.9C151.6 70 234.1 38.6 290.1 82c67.4-22.3 136.3 29.4 130.9 101.1c41.1 12.6 52.8 66.9 19.7 95.2m-70 74.3c-3.1-20.6-40.9-4.6-43.1-27.1c-3.1-32 43.7-101.1 40-128c-3.4-24-19.4-29.1-33.4-29.4c-13.4-.3-16.9 2-21.4 4.6c-2.9 1.7-6.6 4.9-11.7-.3c-6.3-6-11.1-11.7-19.4-12.9c-12.3-2-17.7 2-26.6 9.7c-3.4 2.9-12 12.9-20 9.1c-3.4-1.7-15.4-7.7-24-11.4c-16.3-7.1-40 4.6-48.6 20c-12.9 22.9-38 113.1-41.7 125.1c-8.6 26.6 10.9 48.6 36.9 47.1c11.1-.6 18.3-4.6 25.4-17.4c4-7.4 41.7-107.7 44.6-112.6c2-3.4 8.9-8 14.6-5.1c5.7 3.1 6.9 9.4 6 15.1c-1.1 9.7-28 70.9-28.9 77.7c-3.4 22.9 26.9 26.6 38.6 4c3.7-7.1 45.7-92.6 49.4-98.3c4.3-6.3 7.4-8.3 11.7-8c3.1 0 8.3.9 7.1 10.9c-1.4 9.4-35.1 72.3-38.9 87.7c-4.6 20.6 6.6 41.4 24.9 50.6c11.4 5.7 62.5 15.7 58.5-11.1m5.7 92.3c-10.3 7.4-12.9 22-5.7 32.6c7.1 10.6 21.4 13.1 32 6c10.6-7.4 13.1-22 6-32.6c-7.4-10.6-21.7-13.5-32.3-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Megaport(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M214.5 209.6v66.2l33.5 33.5l33.3-33.3v-66.4l-33.4-33.4zM248 8C111 8 0 119 0 256s111 248 248 248s248-111 248-248S385 8 248 8m145.1 414.4L367 441.6l-26-19.2v-65.5l-33.4-33.4l-33.4 33.4v65.5L248 441.6l-26.1-19.2v-65.5l-33.4-33.4l-33.5 33.4v65.5l-26.1 19.2l-26.1-19.2v-87l59.5-59.5V188l59.5-59.5V52.9l26.1-19.2L274 52.9v75.6l59.5 59.5v87.6l59.7 59.7v87.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mendeley(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M624.6 325.2c-12.3-12.4-29.7-19.2-48.4-17.2c-43.3-1-49.7-34.9-37.5-98.8c22.8-57.5-14.9-131.5-87.4-130.8c-77.4.7-81.7 82-130.9 82c-48.1 0-54-81.3-130.9-82c-72.9-.8-110.1 73.3-87.4 130.8c12.2 63.9 5.8 97.8-37.5 98.8c-21.2-2.3-37 6.5-53 22.5c-19.9 19.7-19.3 94.8 42.6 102.6c47.1 5.9 81.6-42.9 61.2-87.8c-47.3-103.7 185.9-106.1 146.5-8.2c-.1.1-.2.2-.3.4c-26.8 42.8 6.8 97.4 58.8 95.2c52.1 2.1 85.4-52.6 58.8-95.2c-.1-.2-.2-.3-.3-.4c-39.4-97.9 193.8-95.5 146.5 8.2c-4.6 10-6.7 21.3-5.7 33c4.9 53.4 68.7 74.1 104.9 35.2c17.8-14.8 23.1-65.6 0-88.3m-303.9-19.1h-.6c-43.4 0-62.8-37.5-62.8-62.8c0-34.7 28.2-62.8 62.8-62.8h.6c34.7 0 62.8 28.1 62.8 62.8c0 25-19.2 62.8-62.8 62.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microblog(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M399.36 362.23c29.49-34.69 47.1-78.34 47.1-125.79C446.46 123.49 346.86 32 224 32S1.54 123.49 1.54 236.44S101.14 440.87 224 440.87a239.28 239.28 0 0 0 79.44-13.44a7.18 7.18 0 0 1 8.12 2.56c18.58 25.09 47.61 42.74 79.89 49.92a4.42 4.42 0 0 0 5.22-3.43a4.37 4.37 0 0 0-.85-3.62a87 87 0 0 1 3.69-110.69ZM329.52 212.4l-57.3 43.49L293 324.75a6.5 6.5 0 0 1-9.94 7.22L224 290.92L164.94 332a6.51 6.51 0 0 1-9.95-7.22l20.79-68.86l-57.3-43.49a6.5 6.5 0 0 1 3.8-11.68l71.88-1.51l23.66-67.92a6.5 6.5 0 0 1 12.28 0l23.66 67.92l71.88 1.51a6.5 6.5 0 0 1 3.88 11.68Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microsoft(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 32h214.6v214.6H0zm233.4 0H448v214.6H233.4zM0 265.4h214.6V480H0zm233.4 0H448V480H233.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mix(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 64v348.9c0 56.2 88 58.1 88 0V174.3c7.9-52.9 88-50.4 88 6.5v175.3c0 57.9 96 58 96 0V240c5.3-54.7 88-52.5 88 4.3v23.8c0 59.9 88 56.6 88 0V64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mixcloud(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M424.43 219.729C416.124 134.727 344.135 68 256.919 68c-72.266 0-136.224 46.516-159.205 114.074c-54.545 8.029-96.63 54.822-96.63 111.582c0 62.298 50.668 112.966 113.243 112.966h289.614c52.329 0 94.969-42.362 94.969-94.693c0-45.131-32.118-83.063-74.48-92.2m-20.489 144.53H114.327c-39.04 0-70.881-31.564-70.881-70.604s31.841-70.604 70.881-70.604c18.827 0 36.548 7.475 49.838 20.766c19.963 19.963 50.133-10.227 30.18-30.18c-14.675-14.398-32.672-24.365-52.053-29.349c19.935-44.3 64.79-73.926 114.628-73.926c69.496 0 125.979 56.483 125.979 125.702c0 13.568-2.215 26.857-6.369 39.594c-8.943 27.517 32.133 38.939 40.147 13.29c2.769-8.306 4.984-16.889 6.369-25.472c19.381 7.476 33.502 26.303 33.502 48.453c0 28.795-23.535 52.33-52.607 52.33m235.069-52.33c0 44.024-12.737 86.386-37.102 122.657c-4.153 6.092-10.798 9.414-17.72 9.414c-16.317 0-27.127-18.826-17.443-32.949c19.381-29.349 29.903-63.682 29.903-99.122s-10.521-69.773-29.903-98.845c-15.655-22.831 19.361-47.24 35.163-23.534c24.366 35.993 37.102 78.356 37.102 122.379m-70.88 0c0 31.565-9.137 62.021-26.857 88.325c-4.153 6.091-10.798 9.136-17.72 9.136c-17.201 0-27.022-18.979-17.443-32.948c13.013-19.104 19.658-41.255 19.658-64.513c0-22.981-6.645-45.408-19.658-64.512c-15.761-22.986 19.008-47.095 35.163-23.535c17.719 26.026 26.857 56.483 26.857 88.047"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mixer(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M114.57 76.07a45.71 45.71 0 0 0-67.51-6.41c-17.58 16.18-19 43.52-4.75 62.77l91.78 123l-92.33 124.15c-14.23 19.25-13.11 46.59 4.74 62.77a45.71 45.71 0 0 0 67.5-6.41L242.89 262.7a12.14 12.14 0 0 0 0-14.23Zm355.67 303.51l-92.33-124.13l91.78-123c14.22-19.25 12.83-46.59-4.75-62.77a45.71 45.71 0 0 0-67.51 6.41l-128 172.12a12.14 12.14 0 0 0 0 14.23L398 435.94a45.71 45.71 0 0 0 67.51 6.41c17.84-16.18 18.96-43.52 4.73-62.77"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mizuni(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M248 8C111 8 0 119.1 0 256c0 137 111 248 248 248s248-111 248-248C496 119.1 385 8 248 8m-80 351.9c-31.4 10.6-58.8 27.3-80 48.2V136c0-22.1 17.9-40 40-40s40 17.9 40 40zm120-9.9c-12.9-2-26.2-3.1-39.8-3.1c-13.8 0-27.2 1.1-40.2 3.1V136c0-22.1 17.9-40 40-40s40 17.9 40 40zm120 57.7c-21.2-20.8-48.6-37.4-80-48V136c0-22.1 17.9-40 40-40s40 17.9 40 40z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Modx(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m356 241.8l36.7 23.7V480l-133-83.8zM440 75H226.3l-23 37.8l153.5 96.5zm-89 142.8L55.2 32v214.5l46 29zM97 294.2L8 437h213.7l125-200.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Monero(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M352 384h108.4C417 455.9 338.1 504 248 504S79 455.9 35.6 384H144V256.2L248 361l104-105zM88 336V128l159.4 159.4L408 128v208h74.8c8.5-25.1 13.2-52 13.2-80C496 119 385 8 248 8S0 119 0 256c0 28 4.6 54.9 13.2 80z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Napster(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M298.3 373.6c-14.2 13.6-31.3 24.1-50.4 30.5c-19-6.4-36.2-16.9-50.3-30.5zm44-199.6c20-16.9 43.6-29.2 69.6-36.2V299c0 219.4-328 217.6-328 .3V137.7c25.9 6.9 49.6 19.6 69.5 36.4c56.8-40 132.5-39.9 188.9-.1m-208.8-58.5c64.4-60 164.3-60.1 228.9-.2c-7.1 3.5-13.9 7.3-20.6 11.5c-58.7-30.5-129.2-30.4-187.9.1c-6.3-4-13.9-8.2-20.4-11.4M43.8 93.2v69.3c-58.4 36.5-58.4 121.1.1 158.3c26.4 245.1 381.7 240.3 407.6 1.5l.3-1.7c58.7-36.3 58.9-121.7.2-158.2V93.2c-17.3.5-34 3-50.1 7.4c-82-91.5-225.5-91.5-307.5.1c-16.3-4.4-33.1-7-50.6-7.5M259.2 352s36-.3 61.3-1.5c10.2-.5 21.1-4 25.5-6.5c26.3-15.1 25.4-39.2 26.2-47.4c-79.5-.6-99.9-3.9-113 55.4m-135.5-55.3c.8 8.2-.1 32.3 26.2 47.4c4.4 2.5 15.2 6 25.5 6.5c25.3 1.1 61.3 1.5 61.3 1.5c-13.2-59.4-33.7-56.1-113-55.4m169.1 123.4c-3.2-5.3-6.9-7.3-6.9-7.3c-24.8 7.3-52.2 6.9-75.9 0c0 0-2.9 1.5-6.4 6.6c-2.8 4.1-3.7 9.6-3.7 9.6c29.1 17.6 67.1 17.6 96.2 0c-.1-.1-.3-4-3.3-8.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Neos(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M415.44 512h-95.11L212.12 357.46v91.1L125.69 512H28V29.82L68.47 0h108.05l123.74 176.13V63.45L386.69 0h97.69v461.5zM38.77 35.27V496l72-52.88V194l215.5 307.64h84.79l52.35-38.17h-78.27L69 13zm82.54 466.61l80-58.78v-101l-79.76-114.4v220.94L49 501.89h72.34zM80.63 10.77l310.6 442.57h82.37V10.77h-79.75v317.56L170.91 10.77zM311 191.65l72 102.81V15.93l-72 53z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Nimblr(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M246.6 299.29c15.57 0 27.15 11.46 27.15 27s-11.62 27-27.15 27c-15.7 0-27.15-11.57-27.15-27s11.55-27 27.15-27M113 326.25c0-15.61 11.68-27 27.15-27s27.15 11.46 27.15 27s-11.47 27-27.15 27c-15.44 0-27.15-11.31-27.15-27M191.76 159C157 159 89.45 178.77 59.25 227L14 0v335.48C14 433.13 93.61 512 191.76 512s177.76-78.95 177.76-176.52S290.13 159 191.76 159m0 308.12c-73.27 0-132.51-58.9-132.51-131.59s59.24-131.59 132.51-131.59s132.51 58.86 132.51 131.54S265 467.07 191.76 467.07z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NintendoSwitch(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M95.9 33.5c-44.6 8-80.5 41-91.8 84.4C0 133.6-.3 142.8.2 264.4C.4 376 .5 378.6 2.4 387.3c10.3 46.5 43.3 79.6 90.3 90.5c6.1 1.4 13.9 1.7 64.1 1.9c51.9.4 57.3.3 58.7-1.1c1.4-1.4 1.5-19.3 1.5-222.2c0-150.5-.3-221.3-.9-222.6c-.9-1.7-2.5-1.8-56.9-1.7c-44.2.1-57.5.4-63.3 1.4zm83.9 222.6V444l-37.8-.5c-34.8-.4-38.5-.6-45.5-2.3c-29.9-7.7-52-30.7-58.3-60.7c-2-9.4-2-240.1-.1-249.3c5.6-26.1 23.7-47.7 48-57.4c12.2-4.9 17.9-5.5 57.6-5.6l35.9-.1v188zm-75.9-131.2c-5.8 1.1-14.7 5.6-19.5 9.7c-9.7 8.4-14.6 20.4-13.8 34.5c.4 7.3.8 9.3 3.8 15.2c4.4 9 10.9 15.6 19.9 20c6.2 3.1 7.8 3.4 15.9 3.7c7.3.3 9.9 0 14.8-1.7c20.1-6.8 32.3-26.3 28.8-46.4c-3.9-23.7-26.6-39.7-49.9-35zm158.2-92.3c-.4.3-.6 100.8-.6 223.5c0 202.3.1 222.8 1.5 223.4c2.5.9 74.5.6 83.4-.4c37.7-4.3 71-27.2 89-61.2c2.3-4.4 5.4-11.7 7-16.2c5.8-17.4 5.7-12.8 5.7-146.1c0-106.4-.2-122.3-1.5-129c-9.2-48.3-46.1-84.8-94.5-93.1c-6.5-1.1-16.5-1.4-48.8-1.4c-22.4-.1-40.9.2-41.2.5zm99.1 202.1c14.5 3.8 26.3 14.8 31.2 28.9c3.1 8.7 3 21.5-.1 29.5c-5.7 14.7-16.8 25-31.1 28.8c-23.2 6-47.9-8-54.6-31c-2-7-1.9-18.9.4-26.2c6.9-22.7 31-36.1 54.2-30z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Node(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M316.3 452c-2.1 0-4.2-.6-6.1-1.6L291 439c-2.9-1.6-1.5-2.2-.5-2.5c3.8-1.3 4.6-1.6 8.7-4c.4-.2 1-.1 1.4.1l14.8 8.8c.5.3 1.3.3 1.8 0L375 408c.5-.3.9-.9.9-1.6v-66.7c0-.7-.3-1.3-.9-1.6l-57.8-33.3c-.5-.3-1.2-.3-1.8 0l-57.8 33.3c-.6.3-.9 1-.9 1.6v66.7c0 .6.4 1.2.9 1.5l15.8 9.1c8.6 4.3 13.9-.8 13.9-5.8v-65.9c0-.9.7-1.7 1.7-1.7h7.3c.9 0 1.7.7 1.7 1.7v65.9c0 11.5-6.2 18-17.1 18c-3.3 0-6 0-13.3-3.6l-15.2-8.7c-3.7-2.2-6.1-6.2-6.1-10.5v-66.7c0-4.3 2.3-8.4 6.1-10.5l57.8-33.4c3.7-2.1 8.5-2.1 12.1 0l57.8 33.4c3.7 2.2 6.1 6.2 6.1 10.5v66.7c0 4.3-2.3 8.4-6.1 10.5l-57.8 33.4c-1.7 1.1-3.8 1.7-6 1.7m46.7-65.8c0-12.5-8.4-15.8-26.2-18.2c-18-2.4-19.8-3.6-19.8-7.8c0-3.5 1.5-8.1 14.8-8.1c11.9 0 16.3 2.6 18.1 10.6c.2.8.8 1.3 1.6 1.3h7.5c.5 0 .9-.2 1.2-.5c.3-.4.5-.8.4-1.3c-1.2-13.8-10.3-20.2-28.8-20.2c-16.5 0-26.3 7-26.3 18.6c0 12.7 9.8 16.1 25.6 17.7c18.9 1.9 20.4 4.6 20.4 8.3c0 6.5-5.2 9.2-17.4 9.2c-15.3 0-18.7-3.8-19.8-11.4c-.1-.8-.8-1.4-1.7-1.4h-7.5c-.9 0-1.7.7-1.7 1.7c0 9.7 5.3 21.3 30.6 21.3c18.5 0 29-7.2 29-19.8m54.5-50.1c0 6.1-5 11.1-11.1 11.1s-11.1-5-11.1-11.1c0-6.3 5.2-11.1 11.1-11.1c6-.1 11.1 4.8 11.1 11.1m-1.8 0c0-5.2-4.2-9.3-9.4-9.3c-5.1 0-9.3 4.1-9.3 9.3c0 5.2 4.2 9.4 9.3 9.4c5.2-.1 9.4-4.3 9.4-9.4m-4.5 6.2h-2.6c-.1-.6-.5-3.8-.5-3.9c-.2-.7-.4-1.1-1.3-1.1h-2.2v5h-2.4v-12.5h4.3c1.5 0 4.4 0 4.4 3.3c0 2.3-1.5 2.8-2.4 3.1c1.7.1 1.8 1.2 2.1 2.8c.1 1 .3 2.7.6 3.3m-2.8-8.8c0-1.7-1.2-1.7-1.8-1.7h-2v3.5h1.9c1.6 0 1.9-1.1 1.9-1.8M137.3 191c0-2.7-1.4-5.1-3.7-6.4l-61.3-35.3c-1-.6-2.2-.9-3.4-1h-.6c-1.2 0-2.3.4-3.4 1L3.7 184.6C1.4 185.9 0 188.4 0 191l.1 95c0 1.3.7 2.5 1.8 3.2c1.1.7 2.5.7 3.7 0L42 268.3c2.3-1.4 3.7-3.8 3.7-6.4v-44.4c0-2.6 1.4-5.1 3.7-6.4l15.5-8.9c1.2-.7 2.4-1 3.7-1c1.3 0 2.6.3 3.7 1l15.5 8.9c2.3 1.3 3.7 3.8 3.7 6.4v44.4c0 2.6 1.4 5.1 3.7 6.4l36.4 20.9c1.1.7 2.6.7 3.7 0c1.1-.6 1.8-1.9 1.8-3.2zM472.5 87.3v176.4c0 2.6-1.4 5.1-3.7 6.4l-61.3 35.4c-2.3 1.3-5.1 1.3-7.4 0l-61.3-35.4c-2.3-1.3-3.7-3.8-3.7-6.4v-70.8c0-2.6 1.4-5.1 3.7-6.4l61.3-35.4c2.3-1.3 5.1-1.3 7.4 0l15.3 8.8c1.7 1 3.9-.3 3.9-2.2v-94c0-2.8 3-4.6 5.5-3.2l36.5 20.4c2.3 1.2 3.8 3.7 3.8 6.4m-46 128.9c0-.7-.4-1.3-.9-1.6l-21-12.2c-.6-.3-1.3-.3-1.9 0l-21 12.2c-.6.3-.9.9-.9 1.6v24.3c0 .7.4 1.3.9 1.6l21 12.1c.6.3 1.3.3 1.8 0l21-12.1c.6-.3.9-.9.9-1.6v-24.3zm209.8-.7c2.3-1.3 3.7-3.8 3.7-6.4V192c0-2.6-1.4-5.1-3.7-6.4l-60.9-35.4c-2.3-1.3-5.1-1.3-7.4 0l-61.3 35.4c-2.3 1.3-3.7 3.8-3.7 6.4v70.8c0 2.7 1.4 5.1 3.7 6.4l60.9 34.7c2.2 1.3 5 1.3 7.3 0l36.8-20.5c2.5-1.4 2.5-5 0-6.4L550 241.6c-1.2-.7-1.9-1.9-1.9-3.2v-22.2c0-1.3.7-2.5 1.9-3.2l19.2-11.1c1.1-.7 2.6-.7 3.7 0l19.2 11.1c1.1.7 1.9 1.9 1.9 3.2v17.4c0 2.8 3.1 4.6 5.6 3.2zM559 219c-.4.3-.7.7-.7 1.2v13.6c0 .5.3 1 .7 1.2l11.8 6.8c.4.3 1 .3 1.4 0L584 235c.4-.3.7-.7.7-1.2v-13.6c0-.5-.3-1-.7-1.2l-11.8-6.8c-.4-.3-1-.3-1.4 0zm-254.2 43.5v-70.4c0-2.6-1.6-5.1-3.9-6.4l-61.1-35.2c-2.1-1.2-5-1.4-7.4 0l-61.1 35.2c-2.3 1.3-3.9 3.7-3.9 6.4v70.4c0 2.8 1.9 5.2 4 6.4l61.2 35.2c2.4 1.4 5.2 1.3 7.4 0l61-35.2c1.8-1 3.1-2.7 3.6-4.7c.1-.5.2-1.1.2-1.7m-74.3-124.9l-.8.5h1.1zm76.2 130.2l-.4-.7v.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NodeJs(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M224 508c-6.7 0-13.5-1.8-19.4-5.2l-61.7-36.5c-9.2-5.2-4.7-7-1.7-8c12.3-4.3 14.8-5.2 27.9-12.7c1.4-.8 3.2-.5 4.6.4l47.4 28.1c1.7 1 4.1 1 5.7 0l184.7-106.6c1.7-1 2.8-3 2.8-5V149.3c0-2.1-1.1-4-2.9-5.1L226.8 37.7c-1.7-1-4-1-5.7 0L36.6 144.3c-1.8 1-2.9 3-2.9 5.1v213.1c0 2 1.1 4 2.9 4.9l50.6 29.2c27.5 13.7 44.3-2.4 44.3-18.7V167.5c0-3 2.4-5.3 5.4-5.3h23.4c2.9 0 5.4 2.3 5.4 5.3V378c0 36.6-20 57.6-54.7 57.6c-10.7 0-19.1 0-42.5-11.6l-48.4-27.9C8.1 389.2.7 376.3.7 362.4V149.3c0-13.8 7.4-26.8 19.4-33.7L204.6 9c11.7-6.6 27.2-6.6 38.8 0l184.7 106.7c12 6.9 19.4 19.8 19.4 33.7v213.1c0 13.8-7.4 26.7-19.4 33.7L243.4 502.8c-5.9 3.4-12.6 5.2-19.4 5.2m149.1-210.1c0-39.9-27-50.5-83.7-58c-57.4-7.6-63.2-11.5-63.2-24.9c0-11.1 4.9-25.9 47.4-25.9c37.9 0 51.9 8.2 57.7 33.8c.5 2.4 2.7 4.2 5.2 4.2h24c1.5 0 2.9-.6 3.9-1.7s1.5-2.6 1.4-4.1c-3.7-44.1-33-64.6-92.2-64.6c-52.7 0-84.1 22.2-84.1 59.5c0 40.4 31.3 51.6 81.8 56.6c60.5 5.9 65.2 14.8 65.2 26.7c0 20.6-16.6 29.4-55.5 29.4c-48.9 0-59.6-12.3-63.2-36.6c-.4-2.6-2.6-4.5-5.3-4.5h-23.9c-3 0-5.3 2.4-5.3 5.3c0 31.1 16.9 68.2 97.8 68.2c58.4-.1 92-23.2 92-63.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Npm(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M288 288h-32v-64h32zm288-128v192H288v32H160v-32H0V160zm-416 32H32v128h64v-96h32v96h32zm160 0H192v160h64v-32h64zm224 0H352v128h64v-96h32v96h32v-96h32v96h32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NsEight(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M104.324 269.172h26.067v-26.178h-26.067Zm52.466-26.178l-.055-26.178v-.941a39.325 39.325 0 0 0-78.644.941v.166h26.4v-.166a12.98 12.98 0 0 1 25.956 0v26.178Zm52.356 25.846a91.1 91.1 0 0 1-91.1 91.1h-.609a91.1 91.1 0 0 1-91.1-91.1H0v.166A117.33 117.33 0 0 0 117.44 386.28h.775A117.331 117.331 0 0 0 235.49 268.84v-26.012h-26.344Zm-157.233 0a65.362 65.362 0 0 0 130.723 0h-26.344a39.023 39.023 0 0 1-78.035 0v-25.957H51.968v-26.62A65.42 65.42 0 0 1 182.8 217.48v25.293h26.344V217.48a91.761 91.761 0 0 0-183.522 0v25.4h26.291Zm418.4-71.173c13.67 0 24.573 6.642 30.052 18.264l.719 1.549l23.245-11.511l-.609-1.439c-8.025-19.26-28.5-31.27-53.407-31.27c-23.134 0-43.611 11.4-50.972 28.447c-.123 26.876-.158 23.9 0 24.85c4.7 11.013 14.555 19.37 28.668 24.241a102.033 102.033 0 0 0 19.813 3.984c5.479.72 10.626 1.384 15.829 3.1c6.364 2.1 10.46 5.257 12.84 9.851v9.851c-3.708 7.527-13.781 12.342-25.791 12.342c-14.334 0-25.956-6.918-31.933-19.039l-.72-1.494l-23.021 11.507l.553 1.439c7.915 19.426 29.609 32.044 55.289 32.044c23.632 0 44.608-11.4 52.3-28.447l.166-25.9l-.166-.664c-4.87-11.014-15.219-19.647-28.944-24.241c-7.693-2.712-14.335-3.6-20.7-4.427a83.777 83.777 0 0 1-14.832-2.878c-6.31-1.937-10.4-5.092-12.619-9.63v-8.412c3.377-7.357 12.896-12.117 24.242-12.117ZM287.568 311.344h26.067v-68.4h-26.067Zm352.266-53.3c-2.933-6.254-8.3-12.01-15.441-16.714A37.99 37.99 0 0 0 637.4 226l.166-25.347l-.166-.664c-7.362-15.989-26.733-26.729-48.15-26.729S548.461 184 541.1 199.992l-.166 25.347l.166.664a39.643 39.643 0 0 0 13.006 15.331c-7.2 4.7-12.508 10.46-15.441 16.714l-.166 28.889l.166.72c7.582 15.994 27.893 26.731 50.585 26.731s43.057-10.737 50.584-26.731l.166-28.89Zm-73.22-50.806c3.6-6.31 12.563-10.516 22.58-10.516s19.038 4.206 22.636 10.516v13.725c-3.542 6.2-12.563 10.349-22.636 10.349s-19.094-4.15-22.58-10.349Zm47.319 72.169c-3.764 6.641-13.338 10.9-24.683 10.9c-11.125 0-20.976-4.372-24.684-10.9V263.25c3.708-6.309 13.5-10.515 24.684-10.515c11.345 0 20.919 4.15 24.683 10.515ZM376.4 265.962l-59.827-89.713h-29v40.623h26.51v.387l62.539 94.085H402.3V176.249h-25.9Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Nutritionix(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M88 8.1S221.4-.1 209 112.5c0 0 19.1-74.9 103-40.6c0 0-17.7 74-88 56c0 0 14.6-54.6 66.1-56.6c0 0-39.9-10.3-82.1 48.8c0 0-19.8-94.5-93.6-99.7c0 0 75.2 19.4 77.6 107.5c0 .1-106.4 7-104-119.8m312 315.6c0 48.5-9.7 95.3-32 132.3c-42.2 30.9-105 48-168 48c-62.9 0-125.8-17.1-168-48C9.7 419 0 372.2 0 323.7C0 275.3 17.7 229 40 192c42.2-30.9 97.1-48.6 160-48.6c63 0 117.8 17.6 160 48.6c22.3 37 40 83.3 40 131.7M120 428c0-15.5-12.5-28-28-28s-28 12.5-28 28s12.5 28 28 28s28-12.5 28-28m0-66.2c0-15.5-12.5-28-28-28s-28 12.5-28 28s12.5 28 28 28s28-12.5 28-28m0-66.2c0-15.5-12.5-28-28-28s-28 12.5-28 28s12.5 28 28 28s28-12.5 28-28M192 428c0-15.5-12.5-28-28-28s-28 12.5-28 28s12.5 28 28 28s28-12.5 28-28m0-66.2c0-15.5-12.5-28-28-28s-28 12.5-28 28s12.5 28 28 28s28-12.5 28-28m0-66.2c0-15.5-12.5-28-28-28s-28 12.5-28 28s12.5 28 28 28s28-12.5 28-28M264 428c0-15.5-12.5-28-28-28s-28 12.5-28 28s12.5 28 28 28s28-12.5 28-28m0-66.2c0-15.5-12.5-28-28-28s-28 12.5-28 28s12.5 28 28 28s28-12.5 28-28m0-66.2c0-15.5-12.5-28-28-28s-28 12.5-28 28s12.5 28 28 28s28-12.5 28-28M336 428c0-15.5-12.5-28-28-28s-28 12.5-28 28s12.5 28 28 28s28-12.5 28-28m0-66.2c0-15.5-12.5-28-28-28s-28 12.5-28 28s12.5 28 28 28s28-12.5 28-28m0-66.2c0-15.5-12.5-28-28-28s-28 12.5-28 28s12.5 28 28 28s28-12.5 28-28m24-39.6c-4.8-22.3-7.4-36.9-16-56c-38.8-19.9-90.5-32-144-32S94.8 180.1 56 200c-8.8 19.5-11.2 33.9-16 56c42.2-7.9 98.7-14.8 160-14.8s117.8 6.9 160 14.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OctopusDeploy(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M455.6 349.2c-45.891-39.09-36.67-77.877-16.095-128.11C475.16 134.04 415.967 34.14 329.93 8.3C237.04-19.6 134.252 24.341 99.677 117.147a180.862 180.862 0 0 0-10.988 73.544c1.733 29.543 14.717 52.97 24.09 80.3c17.2 50.161-28.1 92.743-66.662 117.582c-46.806 30.2-36.319 39.857-8.428 41.858c23.378 1.68 44.478-4.548 65.265-15.045c9.2-4.647 40.687-18.931 45.13-28.588c-12.184 26.59-36.962 72.702-21.463 102.102c19.1 36.229 67.112-31.77 76.709-45.812c8.591-12.572 42.963-81.279 63.627-46.926c18.865 31.361 8.6 76.391 35.738 104.622c32.854 34.2 51.155-18.312 51.412-44.221c.163-16.411-6.1-95.852 29.9-59.944c21.421 21.381 52.905 71.181 88.561 67.023c38.736-4.516-22.123-67.967-28.262-78.695c5.393 4.279 53.665 34.128 53.818 9.52c.11-18.789-30.085-34.667-42.524-45.267"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Odnoklassniki(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M275.1 334c-27.4 17.4-65.1 24.3-90 26.9l20.9 20.6l76.3 76.3c27.9 28.6-17.5 73.3-45.7 45.7c-19.1-19.4-47.1-47.4-76.3-76.6L84 503.4c-28.2 27.5-73.6-17.6-45.4-45.7c19.4-19.4 47.1-47.4 76.3-76.3l20.6-20.6c-24.6-2.6-62.9-9.1-90.6-26.9c-32.6-21-46.9-33.3-34.3-59c7.4-14.6 27.7-26.9 54.6-5.7c0 0 36.3 28.9 94.9 28.9s94.9-28.9 94.9-28.9c26.9-21.1 47.1-8.9 54.6 5.7c12.4 25.7-1.9 38-34.5 59.1M30.3 129.7C30.3 58 88.6 0 160 0s129.7 58 129.7 129.7c0 71.4-58.3 129.4-129.7 129.4s-129.7-58-129.7-129.4m66 0c0 35.1 28.6 63.7 63.7 63.7s63.7-28.6 63.7-63.7c0-35.4-28.6-64-63.7-64s-63.7 28.6-63.7 64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OdnoklassnikiSquare(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M184.2 177.1c0-22.1 17.9-40 39.8-40s39.8 17.9 39.8 40c0 22-17.9 39.8-39.8 39.8s-39.8-17.9-39.8-39.8M448 80v352c0 26.5-21.5 48-48 48H48c-26.5 0-48-21.5-48-48V80c0-26.5 21.5-48 48-48h352c26.5 0 48 21.5 48 48m-305.1 97.1c0 44.6 36.4 80.9 81.1 80.9s81.1-36.2 81.1-80.9c0-44.8-36.4-81.1-81.1-81.1s-81.1 36.2-81.1 81.1m174.5 90.7c-4.6-9.1-17.3-16.8-34.1-3.6c0 0-22.7 18-59.3 18s-59.3-18-59.3-18c-16.8-13.2-29.5-5.5-34.1 3.6c-7.9 16.1 1.1 23.7 21.4 37c17.3 11.1 41.2 15.2 56.6 16.8l-12.9 12.9c-18.2 18-35.5 35.5-47.7 47.7c-17.6 17.6 10.7 45.8 28.4 28.6l47.7-47.9c18.2 18.2 35.7 35.7 47.7 47.9c17.6 17.2 46-10.7 28.6-28.6l-47.7-47.7l-13-12.9c15.5-1.6 39.1-5.9 56.2-16.8c20.4-13.3 29.3-21 21.5-37"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OldRepublic(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M235.76 10.23c7.5-.31 15-.28 22.5-.09c3.61.14 7.2.4 10.79.73c4.92.27 9.79 1.03 14.67 1.62c2.93.43 5.83.98 8.75 1.46c7.9 1.33 15.67 3.28 23.39 5.4c12.24 3.47 24.19 7.92 35.76 13.21c26.56 12.24 50.94 29.21 71.63 49.88c20.03 20.09 36.72 43.55 48.89 69.19c1.13 2.59 2.44 5.1 3.47 7.74c2.81 6.43 5.39 12.97 7.58 19.63c4.14 12.33 7.34 24.99 9.42 37.83c.57 3.14 1.04 6.3 1.4 9.47c.55 3.83.94 7.69 1.18 11.56c.83 8.34.84 16.73.77 25.1c-.07 4.97-.26 9.94-.75 14.89c-.24 3.38-.51 6.76-.98 10.12c-.39 2.72-.63 5.46-1.11 8.17c-.9 5.15-1.7 10.31-2.87 15.41c-4.1 18.5-10.3 36.55-18.51 53.63c-15.77 32.83-38.83 62.17-67.12 85.12a246.503 246.503 0 0 1-56.91 34.86c-6.21 2.68-12.46 5.25-18.87 7.41c-3.51 1.16-7.01 2.38-10.57 3.39c-6.62 1.88-13.29 3.64-20.04 5c-4.66.91-9.34 1.73-14.03 2.48c-5.25.66-10.5 1.44-15.79 1.74c-6.69.66-13.41.84-20.12.81c-6.82.03-13.65-.12-20.45-.79c-3.29-.23-6.57-.5-9.83-.95c-2.72-.39-5.46-.63-8.17-1.11c-4.12-.72-8.25-1.37-12.35-2.22c-4.25-.94-8.49-1.89-12.69-3.02c-8.63-2.17-17.08-5.01-25.41-8.13c-10.49-4.12-20.79-8.75-30.64-14.25c-2.14-1.15-4.28-2.29-6.35-3.57c-11.22-6.58-21.86-14.1-31.92-22.34c-34.68-28.41-61.41-66.43-76.35-108.7c-3.09-8.74-5.71-17.65-7.8-26.68c-1.48-6.16-2.52-12.42-3.58-18.66c-.4-2.35-.61-4.73-.95-7.09c-.6-3.96-.75-7.96-1.17-11.94c-.8-9.47-.71-18.99-.51-28.49c.14-3.51.34-7.01.7-10.51c.31-3.17.46-6.37.92-9.52c.41-2.81.65-5.65 1.16-8.44c.7-3.94 1.3-7.9 2.12-11.82c3.43-16.52 8.47-32.73 15.26-48.18c1.15-2.92 2.59-5.72 3.86-8.59c8.05-16.71 17.9-32.56 29.49-47.06c20-25.38 45.1-46.68 73.27-62.47c7.5-4.15 15.16-8.05 23.07-11.37c15.82-6.88 32.41-11.95 49.31-15.38c3.51-.67 7.04-1.24 10.56-1.85c2.62-.47 5.28-.7 7.91-1.08c3.53-.53 7.1-.68 10.65-1.04c2.46-.24 4.91-.36 7.36-.51m8.64 24.41c-9.23.1-18.43.99-27.57 2.23c-7.3 1.08-14.53 2.6-21.71 4.3c-13.91 3.5-27.48 8.34-40.46 14.42c-10.46 4.99-20.59 10.7-30.18 17.22c-4.18 2.92-8.4 5.8-12.34 9.03c-5.08 3.97-9.98 8.17-14.68 12.59c-2.51 2.24-4.81 4.7-7.22 7.06c-28.22 28.79-48.44 65.39-57.5 104.69c-2.04 8.44-3.54 17.02-4.44 25.65c-1.1 8.89-1.44 17.85-1.41 26.8c.11 7.14.38 14.28 1.22 21.37c.62 7.12 1.87 14.16 3.2 21.18c1.07 4.65 2.03 9.32 3.33 13.91c6.29 23.38 16.5 45.7 30.07 65.75c8.64 12.98 18.78 24.93 29.98 35.77c16.28 15.82 35.05 29.04 55.34 39.22c7.28 3.52 14.66 6.87 22.27 9.63c5.04 1.76 10.06 3.57 15.22 4.98c11.26 3.23 22.77 5.6 34.39 7.06c2.91.29 5.81.61 8.72.9c13.82 1.08 27.74 1 41.54-.43c4.45-.6 8.92-.99 13.35-1.78c3.63-.67 7.28-1.25 10.87-2.1c4.13-.98 8.28-1.91 12.36-3.07c26.5-7.34 51.58-19.71 73.58-36.2c15.78-11.82 29.96-25.76 42.12-41.28c3.26-4.02 6.17-8.31 9.13-12.55c3.39-5.06 6.58-10.25 9.6-15.54c2.4-4.44 4.74-8.91 6.95-13.45c5.69-12.05 10.28-24.62 13.75-37.49c2.59-10.01 4.75-20.16 5.9-30.45c1.77-13.47 1.94-27.1 1.29-40.65c-.29-3.89-.67-7.77-1-11.66c-2.23-19.08-6.79-37.91-13.82-55.8c-5.95-15.13-13.53-29.63-22.61-43.13c-12.69-18.8-28.24-35.68-45.97-49.83c-25.05-20-54.47-34.55-85.65-42.08c-7.78-1.93-15.69-3.34-23.63-4.45c-3.91-.59-7.85-.82-11.77-1.24c-7.39-.57-14.81-.72-22.22-.58M139.26 83.53c13.3-8.89 28.08-15.38 43.3-20.18c-3.17 1.77-6.44 3.38-9.53 5.29c-11.21 6.68-21.52 14.9-30.38 24.49c-6.8 7.43-12.76 15.73-17.01 24.89c-3.29 6.86-5.64 14.19-6.86 21.71c-.93 4.85-1.3 9.81-1.17 14.75c.13 13.66 4.44 27.08 11.29 38.82c5.92 10.22 13.63 19.33 22.36 27.26c4.85 4.36 10.24 8.09 14.95 12.6c2.26 2.19 4.49 4.42 6.43 6.91c2.62 3.31 4.89 6.99 5.99 11.1c.9 3.02.66 6.2.69 9.31c.02 4.1-.04 8.2.03 12.3c.14 3.54-.02 7.09.11 10.63c.08 2.38.02 4.76.05 7.14c.16 5.77.06 11.53.15 17.3c.11 2.91.02 5.82.13 8.74c.03 1.63.13 3.28-.03 4.91c-.91.12-1.82.18-2.73.16c-10.99 0-21.88-2.63-31.95-6.93c-6-2.7-11.81-5.89-17.09-9.83c-5.75-4.19-11.09-8.96-15.79-14.31c-6.53-7.24-11.98-15.39-16.62-23.95c-1.07-2.03-2.24-4.02-3.18-6.12c-1.16-2.64-2.62-5.14-3.67-7.82c-4.05-9.68-6.57-19.94-8.08-30.31c-.49-4.44-1.09-8.88-1.2-13.35c-.7-15.73.84-31.55 4.67-46.82c2.12-8.15 4.77-16.18 8.31-23.83c6.32-14.2 15.34-27.18 26.3-38.19c6.28-6.2 13.13-11.84 20.53-16.67m175.37-20.12c2.74.74 5.41 1.74 8.09 2.68c6.36 2.33 12.68 4.84 18.71 7.96c13.11 6.44 25.31 14.81 35.82 24.97c10.2 9.95 18.74 21.6 25.14 34.34c1.28 2.75 2.64 5.46 3.81 8.26c6.31 15.1 10 31.26 11.23 47.57c.41 4.54.44 9.09.45 13.64c.07 11.64-1.49 23.25-4.3 34.53c-1.97 7.27-4.35 14.49-7.86 21.18c-3.18 6.64-6.68 13.16-10.84 19.24c-6.94 10.47-15.6 19.87-25.82 27.22c-10.48 7.64-22.64 13.02-35.4 15.38c-3.51.69-7.08 1.08-10.66 1.21c-1.85.06-3.72.16-5.56-.1c-.28-2.15 0-4.31-.01-6.46c-.03-3.73.14-7.45.1-11.17c.19-7.02.02-14.05.21-21.07c.03-2.38-.03-4.76.03-7.14c.17-5.07-.04-10.14.14-15.21c.1-2.99-.24-6.04.51-8.96c.66-2.5 1.78-4.86 3.09-7.08c4.46-7.31 11.06-12.96 17.68-18.26c5.38-4.18 10.47-8.77 15.02-13.84c7.68-8.37 14.17-17.88 18.78-28.27c2.5-5.93 4.52-12.1 5.55-18.46c.86-4.37 1.06-8.83 1.01-13.27c-.02-7.85-1.4-15.65-3.64-23.17c-1.75-5.73-4.27-11.18-7.09-16.45c-3.87-6.93-8.65-13.31-13.96-19.2c-9.94-10.85-21.75-19.94-34.6-27.1c-1.85-1.02-3.84-1.82-5.63-2.97m-100.8 58.45c.98-1.18 1.99-2.33 3.12-3.38c-.61.93-1.27 1.81-1.95 2.68c-3.1 3.88-5.54 8.31-7.03 13.06c-.87 3.27-1.68 6.6-1.73 10c-.07 2.52-.08 5.07.32 7.57c1.13 7.63 4.33 14.85 8.77 21.12c2 2.7 4.25 5.27 6.92 7.33c1.62 1.27 3.53 2.09 5.34 3.05c3.11 1.68 6.32 3.23 9.07 5.48c2.67 2.09 4.55 5.33 4.4 8.79c-.01 73.67 0 147.34-.01 221.02c0 1.35-.08 2.7.04 4.04c.13 1.48.82 2.83 1.47 4.15c.86 1.66 1.78 3.34 3.18 4.62c.85.77 1.97 1.4 3.15 1.24c1.5-.2 2.66-1.35 3.45-2.57c.96-1.51 1.68-3.16 2.28-4.85c.76-2.13.44-4.42.54-6.63c.14-4.03-.02-8.06.14-12.09c.03-5.89.03-11.77.06-17.66c.14-3.62.03-7.24.11-10.86c.15-4.03-.02-8.06.14-12.09c.03-5.99.03-11.98.07-17.97c.14-3.62.02-7.24.11-10.86c.14-3.93-.02-7.86.14-11.78c.03-5.99.03-11.98.06-17.97c.16-3.94-.01-7.88.19-11.82c.29 1.44.13 2.92.22 4.38c.19 3.61.42 7.23.76 10.84c.32 3.44.44 6.89.86 10.32c.37 3.1.51 6.22.95 9.31c.57 4.09.87 8.21 1.54 12.29c1.46 9.04 2.83 18.11 5.09 26.99c1.13 4.82 2.4 9.61 4 14.3c2.54 7.9 5.72 15.67 10.31 22.62c1.73 2.64 3.87 4.98 6.1 7.21c.27.25.55.51.88.71c.6.25 1.31-.07 1.7-.57c.71-.88 1.17-1.94 1.7-2.93c4.05-7.8 8.18-15.56 12.34-23.31c.7-1.31 1.44-2.62 2.56-3.61c1.75-1.57 3.84-2.69 5.98-3.63c2.88-1.22 5.9-2.19 9.03-2.42c6.58-.62 13.11.75 19.56 1.85c3.69.58 7.4 1.17 11.13 1.41c3.74.1 7.48.05 11.21-.28c8.55-.92 16.99-2.96 24.94-6.25c5.3-2.24 10.46-4.83 15.31-7.93c11.46-7.21 21.46-16.57 30.04-27.01c1.17-1.42 2.25-2.9 3.46-4.28c-1.2 3.24-2.67 6.37-4.16 9.48c-1.25 2.9-2.84 5.61-4.27 8.42c-5.16 9.63-11.02 18.91-17.75 27.52c-4.03 5.21-8.53 10.05-13.33 14.57c-6.64 6.05-14.07 11.37-22.43 14.76c-8.21 3.37-17.31 4.63-26.09 3.29c-3.56-.58-7.01-1.69-10.41-2.88c-2.79-.97-5.39-2.38-8.03-3.69c-3.43-1.71-6.64-3.81-9.71-6.08c2.71 3.06 5.69 5.86 8.7 8.61c4.27 3.76 8.74 7.31 13.63 10.23c3.98 2.45 8.29 4.4 12.84 5.51c1.46.37 2.96.46 4.45.6c-1.25 1.1-2.63 2.04-3.99 2.98c-9.61 6.54-20.01 11.86-30.69 16.43c-20.86 8.7-43.17 13.97-65.74 15.34c-4.66.24-9.32.36-13.98.36c-4.98-.11-9.97-.13-14.92-.65c-11.2-.76-22.29-2.73-33.17-5.43c-10.35-2.71-20.55-6.12-30.3-10.55c-8.71-3.86-17.12-8.42-24.99-13.79c-1.83-1.31-3.74-2.53-5.37-4.08c6.6-1.19 13.03-3.39 18.99-6.48c5.74-2.86 10.99-6.66 15.63-11.07c2.24-2.19 4.29-4.59 6.19-7.09c-3.43 2.13-6.93 4.15-10.62 5.78c-4.41 2.16-9.07 3.77-13.81 5.02c-5.73 1.52-11.74 1.73-17.61 1.14c-8.13-.95-15.86-4.27-22.51-8.98c-4.32-2.94-8.22-6.43-11.96-10.06c-9.93-10.16-18.2-21.81-25.66-33.86c-3.94-6.27-7.53-12.75-11.12-19.22c-1.05-2.04-2.15-4.05-3.18-6.1c2.85 2.92 5.57 5.97 8.43 8.88c8.99 8.97 18.56 17.44 29.16 24.48c7.55 4.9 15.67 9.23 24.56 11.03c3.11.73 6.32.47 9.47.81c2.77.28 5.56.2 8.34.3c5.05.06 10.11.04 15.16-.16c3.65-.16 7.27-.66 10.89-1.09c2.07-.25 4.11-.71 6.14-1.2c3.88-.95 8.11-.96 11.83.61c4.76 1.85 8.44 5.64 11.38 9.71c2.16 3.02 4.06 6.22 5.66 9.58c1.16 2.43 2.46 4.79 3.55 7.26c1 2.24 2.15 4.42 3.42 6.52c.67 1.02 1.4 2.15 2.62 2.55c1.06-.75 1.71-1.91 2.28-3.03c2.1-4.16 3.42-8.65 4.89-13.05c2.02-6.59 3.78-13.27 5.19-20.02c2.21-9.25 3.25-18.72 4.54-28.13c.56-3.98.83-7.99 1.31-11.97c.87-10.64 1.9-21.27 2.24-31.94c.08-1.86.24-3.71.25-5.57c.01-4.35.25-8.69.22-13.03c-.01-2.38-.01-4.76 0-7.13c.05-5.07-.2-10.14-.22-15.21c-.2-6.61-.71-13.2-1.29-19.78c-.73-5.88-1.55-11.78-3.12-17.51c-2.05-7.75-5.59-15.03-9.8-21.82c-3.16-5.07-6.79-9.88-11.09-14.03c-3.88-3.86-8.58-7.08-13.94-8.45c-1.5-.41-3.06-.45-4.59-.64c.07-2.99.7-5.93 1.26-8.85c1.59-7.71 3.8-15.3 6.76-22.6c1.52-4.03 3.41-7.9 5.39-11.72c3.45-6.56 7.62-12.79 12.46-18.46m31.27 1.7c.35-.06.71-.12 1.07-.19c.19 1.79.09 3.58.1 5.37v38.13c-.01 1.74.13 3.49-.15 5.22c-.36-.03-.71-.05-1.06-.05c-.95-3.75-1.72-7.55-2.62-11.31c-.38-1.53-.58-3.09-1.07-4.59c-1.7-.24-3.43-.17-5.15-.2c-5.06-.01-10.13 0-15.19-.01c-1.66-.01-3.32.09-4.98-.03c-.03-.39-.26-.91.16-1.18c1.28-.65 2.72-.88 4.06-1.35c3.43-1.14 6.88-2.16 10.31-3.31c1.39-.48 2.9-.72 4.16-1.54c.04-.56.02-1.13-.05-1.68c-1.23-.55-2.53-.87-3.81-1.28c-3.13-1.03-6.29-1.96-9.41-3.02c-1.79-.62-3.67-1-5.41-1.79c-.03-.37-.07-.73-.11-1.09c5.09-.19 10.2.06 15.3-.12c3.36-.13 6.73.08 10.09-.07c.12-.39.26-.77.37-1.16c1.08-4.94 2.33-9.83 3.39-14.75m5.97-.2c.36.05.72.12 1.08.2c.98 3.85 1.73 7.76 2.71 11.61c.36 1.42.56 2.88 1.03 4.27c2.53.18 5.07-.01 7.61.05c5.16.12 10.33.12 15.49.07c.76-.01 1.52.03 2.28.08c-.04.36-.07.72-.1 1.08c-1.82.83-3.78 1.25-5.67 1.89c-3.73 1.23-7.48 2.39-11.22 3.57c-.57.17-1.12.42-1.67.64c-.15.55-.18 1.12-.12 1.69c.87.48 1.82.81 2.77 1.09c4.88 1.52 9.73 3.14 14.63 4.6c.38.13.78.27 1.13.49c.4.27.23.79.15 1.18c-1.66.13-3.31.03-4.97.04c-5.17.01-10.33-.01-15.5.01c-1.61.03-3.22-.02-4.82.21c-.52 1.67-.72 3.42-1.17 5.11c-.94 3.57-1.52 7.24-2.54 10.78c-.36.01-.71.02-1.06.06c-.29-1.73-.15-3.48-.15-5.22v-38.13c.02-1.78-.08-3.58.11-5.37M65.05 168.33c1.12-2.15 2.08-4.4 3.37-6.46c-1.82 7.56-2.91 15.27-3.62 23c-.8 7.71-.85 15.49-.54 23.23c1.05 19.94 5.54 39.83 14.23 57.88c2.99 5.99 6.35 11.83 10.5 17.11c6.12 7.47 12.53 14.76 19.84 21.09c4.8 4.1 9.99 7.78 15.54 10.8c3.27 1.65 6.51 3.39 9.94 4.68c5.01 2.03 10.19 3.61 15.42 4.94c3.83.96 7.78 1.41 11.52 2.71c5 1.57 9.47 4.61 13.03 8.43c4.93 5.23 8.09 11.87 10.2 18.67c.99 2.9 1.59 5.91 2.17 8.92c.15.75.22 1.52.16 2.29c-6.5 2.78-13.26 5.06-20.26 6.18c-4.11.78-8.29.99-12.46 1.08c-10.25.24-20.47-1.76-30.12-5.12c-3.74-1.42-7.49-2.85-11.03-4.72c-8.06-3.84-15.64-8.7-22.46-14.46c-2.92-2.55-5.83-5.13-8.4-8.03c-9.16-9.83-16.3-21.41-21.79-33.65c-2.39-5.55-4.61-11.18-6.37-16.96c-1.17-3.94-2.36-7.89-3.26-11.91c-.75-2.94-1.22-5.95-1.87-8.92c-.46-2.14-.69-4.32-1.03-6.48c-.85-5.43-1.28-10.93-1.33-16.43c.11-6.18.25-12.37 1.07-18.5c.4-2.86.67-5.74 1.15-8.6c.98-5.7 2.14-11.37 3.71-16.93c3.09-11.65 7.48-22.95 12.69-33.84m363.73-6.44c1.1 1.66 1.91 3.48 2.78 5.26c2.1 4.45 4.24 8.9 6.02 13.49c7.61 18.76 12.3 38.79 13.04 59.05c.02 1.76.07 3.52.11 5.29c.13 9.57-1.27 19.09-3.18 28.45c-.73 3.59-1.54 7.17-2.58 10.69c-4.04 14.72-10 29-18.41 41.78c-8.21 12.57-19.01 23.55-31.84 31.41c-5.73 3.59-11.79 6.64-18.05 9.19c-5.78 2.19-11.71 4.03-17.8 5.11c-6.4 1.05-12.91 1.52-19.4 1.23c-7.92-.48-15.78-2.07-23.21-4.85c-1.94-.8-3.94-1.46-5.84-2.33c-.21-1.51.25-2.99.53-4.46c1.16-5.74 3.03-11.36 5.7-16.58c2.37-4.51 5.52-8.65 9.46-11.9c2.43-2.05 5.24-3.61 8.16-4.83c3.58-1.5 7.47-1.97 11.24-2.83c7.23-1.71 14.37-3.93 21.15-7c10.35-4.65 19.71-11.38 27.65-19.46c1.59-1.61 3.23-3.18 4.74-4.87c3.37-3.76 6.71-7.57 9.85-11.53c7.48-10.07 12.82-21.59 16.71-33.48c1.58-5.3 3.21-10.6 4.21-16.05c.63-2.87 1.04-5.78 1.52-8.68c.87-6.09 1.59-12.22 1.68-18.38c.12-6.65.14-13.32-.53-19.94c-.73-7.99-1.87-15.96-3.71-23.78"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Opencart(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M423.3 440.7c0 25.3-20.3 45.6-45.6 45.6s-45.8-20.3-45.8-45.6s20.6-45.8 45.8-45.8c25.4 0 45.6 20.5 45.6 45.8m-253.9-45.8c-25.3 0-45.6 20.6-45.6 45.8s20.3 45.6 45.6 45.6s45.8-20.3 45.8-45.6s-20.5-45.8-45.8-45.8m291.7-270C158.9 124.9 81.9 112.1 0 25.7c34.4 51.7 53.3 148.9 373.1 144.2c333.3-5 130 86.1 70.8 188.9c186.7-166.7 319.4-233.9 17.2-233.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Openid(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m271.5 432l-68 32C88.5 453.7 0 392.5 0 318.2c0-71.5 82.5-131 191.7-144.3v43c-71.5 12.5-124 53-124 101.3c0 51 58.5 93.3 135.7 103v-340l68-33.2v384zM448 291l-131.3-28.5l36.8-20.7c-19.5-11.5-43.5-20-70-24.8v-43c46.2 5.5 87.7 19.5 120.3 39.3l35-19.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Opera(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M313.9 32.7c-170.2 0-252.6 223.8-147.5 355.1c36.5 45.4 88.6 75.6 147.5 75.6c36.3 0 70.3-11.1 99.4-30.4c-43.8 39.2-101.9 63-165.3 63c-3.9 0-8 0-11.9-.3C104.6 489.6 0 381.1 0 248C0 111 111 0 248 0h.8c63.1.3 120.7 24.1 164.4 63.1c-29-19.4-63.1-30.4-99.3-30.4m101.8 397.7c-40.9 24.7-90.7 23.6-132-5.8c56.2-20.5 97.7-91.6 97.7-176.6c0-84.7-41.2-155.8-97.4-176.6c41.8-29.2 91.2-30.3 132.9-5c105.9 98.7 105.5 265.7-1.2 364"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OptinMonster(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M572.6 421.4c5.6-9.5 4.7-15.2-5.4-11.6c-3-4.9-7-9.5-11.1-13.8c2.9-9.7-.7-14.2-10.8-9.2c-4.6-3.2-10.3-6.5-15.9-9.2c0-15.1-11.6-11.6-17.6-5.7c-10.4-1.5-18.7-.3-26.8 5.7c.3-6.5.3-13 .3-19.7c12.6 0 40.2-11 45.9-36.2c1.4-6.8 1.6-13.8-.3-21.9c-3-13.5-14.3-21.3-25.1-25.7c-.8-5.9-7.6-14.3-14.9-15.9s-12.4 4.9-14.1 10.3c-8.5 0-19.2 2.8-21.1 8.4c-5.4-.5-11.1-1.4-16.8-1.9c2.7-1.9 5.4-3.5 8.4-4.6c5.4-9.2 14.6-11.4 25.7-11.6V256c19.5-.5 43-5.9 53.8-18.1c12.7-13.8 14.6-37.3 12.4-55.1c-2.4-17.3-9.7-37.6-24.6-48.1c-8.4-5.9-21.6-.8-22.7 9.5c-2.2 19.6 1.2 30-38.6 25.1c-10.3-23.8-24.6-44.6-42.7-60C341 49.6 242.9 55.5 166.4 71.7c19.7 4.6 41.1 8.6 59.7 16.5c-26.2 2.4-52.7 11.3-76.2 23.2c-32.8 17-44 29.9-56.7 42.4c14.9-2.2 28.9-5.1 43.8-3.8c-9.7 5.4-18.4 12.2-26.5 20c-25.8.9-23.8-5.3-26.2-25.9c-1.1-10.5-14.3-15.4-22.7-9.7c-28.1 19.9-33.5 79.9-12.2 103.5c10.8 12.2 35.1 17.3 54.9 17.8c-.3 1.1-.3 1.9-.3 2.7c10.8.5 19.5 2.7 24.6 11.6c3 1.1 5.7 2.7 8.1 4.6c-5.4.5-11.1 1.4-16.5 1.9c-3.3-6.6-13.7-8.1-21.1-8.1c-1.6-5.7-6.5-12.2-14.1-10.3c-6.8 1.9-14.1 10-14.9 15.9c-22.5 9.5-30.1 26.8-25.1 47.6c5.3 24.8 33 36.2 45.9 36.2v19.7c-6.6-5-14.3-7.5-26.8-5.7c-5.5-5.5-17.3-10.1-17.3 5.7c-5.9 2.7-11.4 5.9-15.9 9.2c-9.8-4.9-13.6-1.7-11.1 9.2c-4.1 4.3-7.8 8.6-11.1 13.8c-10.2-3.7-11 2.2-5.4 11.6c-1.1 3.5-1.6 7-1.9 10.8c-.5 31.6 44.6 64 73.5 65.1c17.3.5 34.6-8.4 43-23.5c113.2 4.9 226.7 4.1 340.2 0c8.1 15.1 25.4 24.3 42.7 23.5c29.2-1.1 74.3-33.5 73.5-65.1c.2-3.7-.7-7.2-1.7-10.7m-73.8-254c1.1-3 2.4-8.4 2.4-14.6c0-5.9 6.8-8.1 14.1-.8c11.1 11.6 14.9 40.5 13.8 51.1c-4.1-13.6-13-29-30.3-35.7m-4.6 6.7c19.5 6.2 28.6 27.6 29.7 48.9c-1.1 2.7-3 5.4-4.9 7.6c-5.7 5.9-15.4 10-26.2 12.2c4.3-21.3.3-47.3-12.7-63c4.9-.8 10.9-2.4 14.1-5.7m-24.1 6.8c13.8 11.9 20 39.2 14.1 63.5c-4.1.5-8.1.8-11.6.8c-1.9-21.9-6.8-44-14.3-64.6c3.7.3 8.1.3 11.8.3M47.5 203c-1.1-10.5 2.4-39.5 13.8-51.1c7-7.3 14.1-5.1 14.1.8c0 6.2 1.4 11.6 2.4 14.6c-17.3 6.8-26.2 22.2-30.3 35.7m9.7 27.6c-1.9-2.2-3.5-4.9-4.9-7.6c1.4-21.3 10.3-42.7 29.7-48.9c3.2 3.2 9.2 4.9 14.1 5.7c-13 15.7-17 41.6-12.7 63c-10.8-2.2-20.5-6-26.2-12.2m47.9 14.6c-4.1 0-8.1-.3-12.7-.8c-4.6-18.6-1.9-38.9 5.4-53v.3l12.2-5.1c4.9-1.9 9.7-3.8 14.9-4.9c-10.7 19.7-17.4 41.3-19.8 63.5m184-162.7c41.9 0 76.2 34 76.2 75.9c0 42.2-34.3 76.2-76.2 76.2s-76.2-34-76.2-76.2c0-41.8 34.3-75.9 76.2-75.9m115.6 174.3c-.3 17.8-7 48.9-23 57c-13.2 6.6-6.5-7.5-16.5-58.1c13.3.3 26.6.3 39.5 1.1m-54-1.6c.8 4.9 3.8 40.3-1.6 41.9c-11.6 3.5-40 4.3-51.1-1.1c-4.1-3-4.6-35.9-4.3-41.1v.3c18.9-.3 38.1-.3 57 0M278.3 309c-13 3.5-41.6 4.1-54.6-1.6c-6.5-2.7-3.8-42.4-1.9-51.6c19.2-.5 38.4-.5 57.8-.8v.3c1.1 8.3 3.3 51.2-1.3 53.7m-106.5-51.1c12.2-.8 24.6-1.4 36.8-1.6c-2.4 15.4-3 43.5-4.9 52.2c-1.1 6.8-4.3 6.8-9.7 4.3c-21.9-9.8-27.6-35.2-22.2-54.9m-35.4 31.3c7.8-1.1 15.7-1.9 23.5-2.7c1.6 6.2 3.8 11.9 7 17.6c10 17 44 35.7 45.1 7c6.2 14.9 40.8 12.2 54.9 10.8c15.7-1.4 23.8-1.4 26.8-14.3c12.4 4.3 30.8 4.1 44 3c11.3-.8 20.8-.5 24.6-8.9c1.1 5.1 1.9 11.6 4.6 16.8c10.8 21.3 37.3 1.4 46.8-31.6c8.6.8 17.6 1.9 26.5 2.7c-.4 1.3-3.8 7.3 7.3 11.6c-47.6 47-95.7 87.8-163.2 107c-63.2-20.8-112.1-59.5-155.9-106.5c9.6-3.4 10.4-8.8 8-12.5m-21.6 172.5c-3.8 17.8-21.9 29.7-39.7 28.9c-19.2-.8-46.5-17-59.2-36.5c-2.7-31.1 43.8-61.3 66.2-54.6c14.9 4.3 27.8 30.8 33.5 54c0 3-.3 5.7-.8 8.2m-8.7-66c-.5-13.5-.5-27-.3-40.5h.3c2.7-1.6 5.7-3.8 7.8-6.5c6.5-1.6 13-5.1 15.1-9.2c3.3-7.1-7-7.5-5.4-12.4c2.7-1.1 5.7-2.2 7.8-3.5c29.2 29.2 58.6 56.5 97.3 77c-36.8 11.3-72.4 27.6-105.9 47c-1.2-18.6-7.7-35.9-16.7-51.9m337.6 64.6c-103 3.5-206.2 4.1-309.4 0c0 .3 0 .3-.3.3v-.3h.3c35.1-21.6 72.2-39.2 112.4-50.8c11.6 5.1 23 9.5 34.9 13.2c2.2.8 2.2.8 4.3 0c14.3-4.1 28.4-9.2 42.2-15.4c41.5 11.7 78.8 31.7 115.6 53m10.5-12.4c-35.9-19.5-73-35.9-111.9-47.6c38.1-20 71.9-47.3 103.5-76.7c2.2 1.4 4.6 2.4 7.6 3.2c0 .8.3 1.9.5 2.4c-4.6 2.7-7.8 6.2-5.9 10.3c2.2 3.8 8.6 7.6 15.1 8.9c2.4 2.7 5.1 5.1 8.1 6.8c0 13.8-.3 27.6-.8 41.3l.3-.3c-9.3 15.9-15.5 37-16.5 51.7m105.9 6.2c-12.7 19.5-40 35.7-59.2 36.5c-19.3.9-40.5-13.2-40.5-37c5.7-23.2 18.9-49.7 33.5-54c22.7-6.9 69.2 23.4 66.2 54.5M372.9 75.2c-3.8-72.1-100.8-79.7-126-23.5c44.6-24.3 90.3-15.7 126 23.5M74.8 407.1c-15.7 1.6-49.5 25.4-49.5 43.2c0 11.6 15.7 19.5 32.2 14.9c12.2-3.2 31.1-17.6 35.9-27.3c6-11.6-3.7-32.7-18.6-30.8m215.9-176.2c28.6 0 51.9-21.6 51.9-48.4c0-36.1-40.5-58.1-72.2-44.3c9.5 3 16.5 11.6 16.5 21.6c0 23.3-33.3 32-46.5 11.3c-7.3 34.1 19.4 59.8 50.3 59.8M68 474.1c.5 6.5 12.2 12.7 21.6 9.5c6.8-2.7 14.6-10.5 17.3-16.2c3-7-1.1-20-9.7-18.4c-8.9 1.6-29.7 16.7-29.2 25.1m433.2-67c-14.9-1.9-24.6 19.2-18.9 30.8c4.9 9.7 24.1 24.1 36.2 27.3c16.5 4.6 32.2-3.2 32.2-14.9c0-17.8-33.8-41.6-49.5-43.2M478.8 449c-8.4-1.6-12.4 11.3-9.5 18.4c2.4 5.7 10.3 13.5 17.3 16.2c9.2 3.2 21.1-3 21.3-9.5c.9-8.4-20.2-23.5-29.1-25.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Orcid(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M294.75 188.19h-45.92V342h47.47c67.62 0 83.12-51.34 83.12-76.91c0-41.64-26.54-76.9-84.67-76.9M256 8C119 8 8 119 8 256s111 248 248 248s248-111 248-248S393 8 256 8m-80.79 360.76h-29.84v-207.5h29.84zm-14.92-231.14a19.57 19.57 0 1 1 19.57-19.57a19.64 19.64 0 0 1-19.57 19.57M300 369h-81V161.26h80.6c76.73 0 110.44 54.83 110.44 103.85C410 318.39 368.38 369 300 369"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Osi(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 266.44C10.3 130.64 105.4 34 221.8 18.34c138.8-18.6 255.6 75.8 278 201.1c21.3 118.8-44 230-151.6 274c-9.3 3.8-14.4 1.7-18-7.7q-26.7-69.45-53.4-139c-3.1-8.1-1-13.2 7-16.8c24.2-11 39.3-29.4 43.3-55.8a71.47 71.47 0 0 0-64.5-82.2c-39-3.4-71.8 23.7-77.5 59.7c-5.2 33 11.1 63.7 41.9 77.7c9.6 4.4 11.5 8.6 7.8 18.4q-26.85 69.9-53.7 139.9c-2.6 6.9-8.3 9.3-15.5 6.5c-52.6-20.3-101.4-61-130.8-119c-24.9-49.2-25.2-87.7-26.8-108.7m20.9-1.9c.4 6.6.6 14.3 1.3 22.1c6.3 71.9 49.6 143.5 131 183.1c3.2 1.5 4.4.8 5.6-2.3q22.35-58.65 45-117.3c1.3-3.3.6-4.8-2.4-6.7c-31.6-19.9-47.3-48.5-45.6-86c1-21.6 9.3-40.5 23.8-56.3c30-32.7 77-39.8 115.5-17.6a91.64 91.64 0 0 1 45.2 90.4c-3.6 30.6-19.3 53.9-45.7 69.8c-2.7 1.6-3.5 2.9-2.3 6q22.8 58.8 45.2 117.7c1.2 3.1 2.4 3.8 5.6 2.3c35.5-16.6 65.2-40.3 88.1-72c34.8-48.2 49.1-101.9 42.3-161c-13.7-117.5-119.4-214.8-255.5-198c-106.1 13-195.3 102.5-197.1 225.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageFour(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M248 504C111 504 0 393 0 256S111 8 248 8c20.9 0 41.3 2.6 60.7 7.5L42.3 392H248zm0-143.6V146.8L98.6 360.4zm96 31.6v92.7c45.7-19.2 84.5-51.7 111.4-92.7zm57.4-138.2l-21.2 8.4l21.2 8.3zm-20.3 54.5c-6.7 0-8 6.3-8 12.9v7.7h16.2v-10c0-5.9-2.3-10.6-8.2-10.6M496 256c0 37.3-8.2 72.7-23 104.4H344V27.3C433.3 64.8 496 153.1 496 256M360.4 143.6h68.2V96h-13.9v32.6h-13.9V99h-13.9v29.6h-12.7V96h-13.9v47.6zm68.1 185.3H402v-11c0-15.4-5.6-25.2-20.9-25.2c-15.4 0-20.7 10.6-20.7 25.9v25.3h68.2v-15zm0-103l-68.2 29.7V268l68.2 29.5v-16.6l-14.4-5.7v-26.5l14.4-5.9zm-4.8-68.5h-35.6V184H402v-12.2h11c8.6 15.8 1.3 35.3-18.6 35.3c-22.5 0-28.3-25.3-15.5-37.7l-11.6-10.6c-16.2 17.5-12.2 63.9 27.1 63.9c34 0 44.7-35.9 29.3-65.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pagelines(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 312.7c-55.1 136.7-187.1 54-187.1 54c-40.5 81.8-107.4 134.4-184.6 134.7c-16.1 0-16.6-24.4 0-24.4c64.4-.3 120.5-42.7 157.2-110.1c-41.1 15.9-118.6 27.9-161.6-82.2c109-44.9 159.1 11.2 178.3 45.5c9.9-24.4 17-50.9 21.6-79.7c0 0-139.7 21.9-149.5-98.1c119.1-47.9 152.6 76.7 152.6 76.7c1.6-16.7 3.3-52.6 3.3-53.4c0 0-106.3-73.7-38.1-165.2c124.6 43 61.4 162.4 61.4 162.4c.5 1.6.5 23.8 0 33.4c0 0 45.2-89 136.4-57.5c-4.2 134-141.9 106.4-141.9 106.4c-4.4 27.4-11.2 53.4-20 77.5c0 0 83-91.8 172-20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Palfed(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384.9 193.9c0-47.4-55.2-44.2-95.4-29.8c-1.3 39.4-2.5 80.7-3 119.8c.7 2.8 2.6 6.2 15.1 6.2c36.8 0 83.4-42.8 83.3-96.2m-194.5 72.2c.2 0 6.5-2.7 11.2-2.7c26.6 0 20.7 44.1-14.4 44.1c-21.5 0-37.1-18.1-37.1-43c0-42 42.9-95.6 100.7-126.5c1-12.4 3-22 10.5-28.2c11.2-9 26.6-3.5 29.5 11.1c72.2-22.2 135.2 1 135.2 72c0 77.9-79.3 152.6-140.1 138.2c-.1 39.4.9 74.4 2.7 100v.2c.2 3.4.6 12.5-5.3 19.1c-9.6 10.6-33.4 10-36.4-22.3c-4.1-44.4.2-206.1 1.4-242.5c-21.5 15-58.5 50.3-58.5 75.9c.2 2.5.4 4 .6 4.6M8 181.1s-.1 37.4 38.4 37.4h30l22.4 217.2s0 44.3 44.7 44.3h288.9s44.7-.4 44.7-44.3l22.4-217.2h30s38.4 1.2 38.4-37.4c0 0 .1-37.4-38.4-37.4h-30.1c-7.3-25.6-30.2-74.3-119.4-74.3h-28V50.3s-2.7-18.4-21.1-18.4h-85.8s-21.1 0-21.1 18.4v19.1h-28.1s-105 4.2-120.5 74.3h-29S8 142.5 8 181.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Patreon(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 194.8c0 101.3-82.4 183.8-183.8 183.8c-101.7 0-184.4-82.4-184.4-183.8c0-101.6 82.7-184.3 184.4-184.3C429.6 10.5 512 93.2 512 194.8M0 501.5h90v-491H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paypal(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M111.4 295.9c-3.5 19.2-17.4 108.7-21.5 134c-.3 1.8-1 2.5-3 2.5H12.3c-7.6 0-13.1-6.6-12.1-13.9L58.8 46.6c1.5-9.6 10.1-16.9 20-16.9c152.3 0 165.1-3.7 204 11.4c60.1 23.3 65.6 79.5 44 140.3c-21.5 62.6-72.5 89.5-140.1 90.3c-43.4.7-69.5-7-75.3 24.2M357.1 152c-1.8-1.3-2.5-1.8-3 1.3c-2 11.4-5.1 22.5-8.8 33.6c-39.9 113.8-150.5 103.9-204.5 103.9c-6.1 0-10.1 3.3-10.9 9.4c-22.6 140.4-27.1 169.7-27.1 169.7c-1 7.1 3.5 12.9 10.6 12.9h63.5c8.6 0 15.7-6.3 17.4-14.9c.7-5.4-1.1 6.1 14.4-91.3c4.6-22 14.3-19.7 29.3-19.7c71 0 126.4-28.8 142.9-112.3c6.5-34.8 4.6-71.4-23.8-92.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PennyArcade(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M421.91 164.27c-4.49 19.45-1.4 6.06-15.1 65.29l39.73-10.61c-22.34-49.61-17.29-38.41-24.63-54.68m-206.09 51.11c-20.19 5.4-11.31 3.03-39.63 10.58l4.46 46.19c28.17-7.59 20.62-5.57 34.82-9.34c42.3-9.79 32.85-56.42.35-47.43m326.16-26.19l-45.47-99.2c-5.69-12.37-19.46-18.84-32.62-15.33c-70.27 18.75-38.72 10.32-135.59 36.23a27.618 27.618 0 0 0-18.89 17.41C144.26 113.27 0 153.75 0 226.67c0 33.5 30.67 67.11 80.9 95.37l1.74 17.88a27.891 27.891 0 0 0-17.77 28.67l4.3 44.48c1.39 14.31 13.43 25.21 27.8 25.2c5.18-.01-3.01 1.78 122.53-31.76c12.57-3.37 21.12-15.02 20.58-28.02c216.59 45.5 401.99-5.98 399.89-84.83c.01-28.15-22.19-66.56-97.99-104.47M255.14 298.3l-21.91 5.88l-48.44 12.91l2.46 23.55l20.53-5.51l4.51 44.51l-115.31 30.78l-4.3-44.52l20.02-5.35l-11.11-114.64l-20.12 5.39l-4.35-44.5c178.15-47.54 170.18-46.42 186.22-46.65c56.66-1.13 64.15 71.84 42.55 104.43a86.7 86.7 0 0 1-50.75 33.72m199.18 16.62l-3.89-39.49l14.9-3.98l-6.61-14.68l-57.76 15.42l-4.1 17.54l19.2-5.12l4.05 39.54l-112.85 30.07l-4.46-44.43l20.99-5.59l33.08-126.47l-17.15 4.56l-4.2-44.48c93.36-24.99 65.01-17.41 135.59-36.24l66.67 145.47l20.79-5.56l4.3 44.48z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Perbyte(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M305.314 284.578H246.6V383.3h58.711q24.423 0 38.193-13.77t13.77-36.11q0-21.826-14.032-35.335t-37.928-13.507M149.435 128.7H90.724v98.723h58.711q24.42 0 38.19-13.773t13.77-36.107q0-21.826-14.029-35.338T149.435 128.7M366.647 32H81.353A81.445 81.445 0 0 0 0 113.352v285.295A81.445 81.445 0 0 0 81.353 480h285.294A81.445 81.445 0 0 0 448 398.647V113.352A81.445 81.445 0 0 0 366.647 32m63.635 366.647a63.706 63.706 0 0 1-63.635 63.635H81.353a63.706 63.706 0 0 1-63.635-63.635V113.352a63.706 63.706 0 0 1 63.635-63.634h285.294a63.706 63.706 0 0 1 63.635 63.634ZM305.314 128.7H246.6v98.723h58.711q24.423 0 38.193-13.773t13.77-36.107q0-21.826-14.032-35.338T305.314 128.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Periscope(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M370 63.6C331.4 22.6 280.5 0 226.6 0C111.9 0 18.5 96.2 18.5 214.4c0 75.1 57.8 159.8 82.7 192.7C137.8 455.5 192.6 512 226.6 512c41.6 0 112.9-94.2 120.9-105c24.6-33.1 82-118.3 82-192.6c0-56.5-21.1-110.1-59.5-150.8M226.6 493.9c-42.5 0-190-167.3-190-279.4c0-107.4 83.9-196.3 190-196.3c100.8 0 184.7 89 184.7 196.3c.1 112.1-147.4 279.4-184.7 279.4M338 206.8c0 59.1-51.1 109.7-110.8 109.7c-100.6 0-150.7-108.2-92.9-181.8v.4c0 24.5 20.1 44.4 44.8 44.4c24.7 0 44.8-19.9 44.8-44.4c0-18.2-11.1-33.8-26.9-40.7c76.6-19.2 141 39.3 141 112.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phabricator(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m323 262.1l-.1-13s21.7-19.8 21.1-21.2l-9.5-20c-.6-1.4-29.5-.5-29.5-.5l-9.4-9.3s.2-28.5-1.2-29.1l-20.1-9.2c-1.4-.6-20.7 21-20.7 21l-13.1-.2s-20.5-21.4-21.9-20.8l-20 8.3c-1.4.5.2 28.9.2 28.9l-9.1 9.1s-29.2-.9-29.7.4l-8.1 19.8c-.6 1.4 21 21 21 21l.1 12.9s-21.7 19.8-21.1 21.2l9.5 20c.6 1.4 29.5.5 29.5.5l9.4 9.3s-.2 31.8 1.2 32.3l20.1 8.3c1.4.6 20.7-23.5 20.7-23.5l13.1.2s20.5 23.8 21.8 23.3l20-7.5c1.4-.6-.2-32.1-.2-32.1l9.1-9.1s29.2.9 29.7-.5l8.1-19.8c.7-1.1-20.9-20.7-20.9-20.7m-44.9-8.7c.7 17.1-12.8 31.6-30.1 32.4c-17.3.8-32.1-12.5-32.8-29.6c-.7-17.1 12.8-31.6 30.1-32.3c17.3-.8 32.1 12.5 32.8 29.5m201.2-37.9l-97-97l-.1.1c-75.1-73.3-195.4-72.8-269.8 1.6c-50.9 51-27.8 27.9-95.7 95.3c-22.3 22.3-22.3 58.7 0 81c69.9 69.4 46.4 46 97.4 97l.1-.1c75.1 73.3 195.4 72.9 269.8-1.6c51-50.9 27.9-27.9 95.3-95.3c22.3-22.3 22.3-58.7 0-81M140.4 363.8c-59.6-59.5-59.6-156 0-215.5c59.5-59.6 156-59.5 215.6 0c59.5 59.5 59.6 156 0 215.6c-59.6 59.5-156 59.4-215.6-.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoenixFramework(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M212.9 344.3c3.8-.1 22.8-1.4 25.6-2.2c-2.4-2.6-43.6-1-68-49.6c-4.3-8.6-7.5-17.6-6.4-27.6c2.9-25.5 32.9-30 52-18.5c36 21.6 63.3 91.3 113.7 97.5c37 4.5 84.6-17 108.2-45.4c-.6-.1-.8-.2-1-.1c-.4.1-.8.2-1.1.3c-33.3 12.1-94.3 9.7-134.7-14.8c-37.6-22.8-53.1-58.7-51.8-74.6c1.8-21.3 22.9-23.2 35.9-19.6c14.4 3.9 24.4 17.6 38.9 27.4c15.6 10.4 32.9 13.7 51.3 10.3c14.9-2.7 34.4-12.3 36.5-14.5c-1.1-.1-1.8-.1-2.5-.2c-6.2-.6-12.4-.8-18.5-1.7C279.8 194.5 262.1 47.4 138.5 37.9C94.2 34.5 39.1 46 2.2 72.9c-.8.6-1.5 1.2-2.2 1.8c.1.2.1.3.2.5c.8 0 1.6-.1 2.4-.2c6.3-1 12.5-.8 18.7.3c23.8 4.3 47.7 23.1 55.9 76.5c5.3 34.3-.7 50.8 8 86.1c19 77.1 91 107.6 127.7 106.4M75.3 64.9c-.9-1-.9-1.2-1.3-2c12.1-2.6 24.2-4.1 36.6-4.8c-1.1 14.7-22.2 21.3-35.3 6.8m196.9 350.5c-42.8 1.2-92-26.7-123.5-61.4c-4.6-5-16.8-20.2-18.6-23.4l.4-.4c6.6 4.1 25.7 18.6 54.8 27c24.2 7 48.1 6.3 71.6-3.3c22.7-9.3 41-.5 43.1 2.9c-18.5 3.8-20.1 4.4-24 7.9c-5.1 4.4-4.6 11.7 7 17.2c26.2 12.4 63-2.8 97.2 25.4c2.4 2 8.1 7.8 10.1 10.7c-.1.2-.3.3-.4.5c-4.8-1.5-16.4-7.5-40.2-9.3c-24.7-2-46.3 5.3-77.5 6.2m174.8-252c16.4-5.2 41.3-13.4 66.5-3.3c16.1 6.5 26.2 18.7 32.1 34.6c3.5 9.4 5.1 19.7 5.1 28.7c-.2 0-.4 0-.6.1c-.2-.4-.4-.9-.5-1.3c-5-22-29.9-43.8-67.6-29.9c-50.2 18.6-130.4 9.7-176.9-48c-.7-.9-2.4-1.7-1.3-3.2c.1-.2 2.1.6 3 1.3c18.1 13.4 38.3 21.9 60.3 26.2c30.5 6.1 54.6 2.9 79.9-5.2m102.7 117.5c-32.4.2-33.8 50.1-103.6 64.4c-18.2 3.7-38.7 4.6-44.9 4.2v-.4c2.8-1.5 14.7-2.6 29.7-16.6c7.9-7.3 15.3-15.1 22.8-22.9c19.5-20.2 41.4-42.2 81.9-39c23.1 1.8 29.3 8.2 36.1 12.7c.3.2.4.5.7.9c-.5 0-.7.1-.9 0c-7-2.7-14.3-3.3-21.8-3.3m-12.3-24.1c-.1.2-.1.4-.2.6c-28.9-4.4-48-7.9-68.5 4c-17 9.9-31.4 20.5-62 24.4c-27.1 3.4-45.1 2.4-66.1-8c-.3-.2-.6-.4-1-.6c0-.2.1-.3.1-.5c24.9 3.8 36.4 5.1 55.5-5.8c22.3-12.9 40.1-26.6 71.3-31c29.6-4.1 51.3 2.5 70.9 16.9M268.6 97.3c-.6-.6-1.1-1.2-2.1-2.3c7.6 0 29.7-1.2 53.4 8.4c19.7 8 32.2 21 50.2 32.9c11.1 7.3 23.4 9.3 36.4 8.1c4.3-.4 8.5-1.2 12.8-1.7c.4-.1.9 0 1.5.3c-.6.4-1.2.9-1.8 1.2c-8.1 4-16.7 6.3-25.6 7.1c-26.1 2.6-50.3-3.7-73.4-15.4c-19.3-9.9-36.4-22.9-51.4-38.6M640 335.7c-3.5 3.1-22.7 11.6-42.7 5.3c-12.3-3.9-19.5-14.9-31.6-24.1c-10-7.6-20.9-7.9-28.1-8.4c.6-.8.9-1.2 1.2-1.4c14.8-9.2 30.5-12.2 47.3-6.5c12.5 4.2 19.2 13.5 30.4 24.2c10.8 10.4 21 9.9 23.1 10.5c.1-.1.2 0 .4.4m-212.5 137c2.2 1.2 1.6 1.5 1.5 2c-18.5-1.4-33.9-7.6-46.8-22.2c-21.8-24.7-41.7-27.9-48.6-29.7c.5-.2.8-.4 1.1-.4c13.1.1 26.1.7 38.9 3.9c25.3 6.4 35 25.4 41.6 35.3c3.2 4.8 7.3 8.3 12.3 11.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoenixSquadron(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M96 63.38C142.49 27.25 201.55 7.31 260.51 8.81c29.58-.38 59.11 5.37 86.91 15.33c-24.13-4.63-49-6.34-73.38-2.45C231.17 27 191 48.84 162.21 80.87c5.67-1 10.78-3.67 16-5.86c18.14-7.87 37.49-13.26 57.23-14.83c19.74-2.13 39.64-.43 59.28 1.92c-14.42 2.79-29.12 4.57-43 9.59c-34.43 11.07-65.27 33.16-86.3 62.63c-13.8 19.71-23.63 42.86-24.67 67.13c-.35 16.49 5.22 34.81 19.83 44a53.27 53.27 0 0 0 37.52 6.74c15.45-2.46 30.07-8.64 43.6-16.33c11.52-6.82 22.67-14.55 32-24.25c3.79-3.22 2.53-8.45 2.62-12.79c-2.12-.34-4.38-1.11-6.3.3a203 203 0 0 1-35.82 15.37c-20 6.17-42.16 8.46-62.1.78c12.79 1.73 26.06.31 37.74-5.44c20.23-9.72 36.81-25.2 54.44-38.77a526.57 526.57 0 0 1 88.9-55.31c25.71-12 52.94-22.78 81.57-24.12c-15.63 13.72-32.15 26.52-46.78 41.38c-14.51 14-27.46 29.5-40.11 45.18c-3.52 4.6-8.95 6.94-13.58 10.16a150.7 150.7 0 0 0-51.89 60.1c-9.33 19.68-14.5 41.85-11.77 63.65c1.94 13.69 8.71 27.59 20.9 34.91c12.9 8 29.05 8.07 43.48 5.1c32.8-7.45 61.43-28.89 81-55.84c20.44-27.52 30.52-62.2 29.16-96.35c-.52-7.5-1.57-15-1.66-22.49c8 19.48 14.82 39.71 16.65 60.83c2 14.28.75 28.76-1.62 42.9c-1.91 11-5.67 21.51-7.78 32.43a165 165 0 0 0 39.34-81.07a183.64 183.64 0 0 0-14.21-104.64c20.78 32 32.34 69.58 35.71 107.48c.49 12.73.49 25.51 0 38.23A243.21 243.21 0 0 1 482 371.34c-26.12 47.34-68 85.63-117.19 108c-78.29 36.23-174.68 31.32-248-14.68A248.34 248.34 0 0 1 25.36 366A238.34 238.34 0 0 1 0 273.08v-31.34C3.93 172 40.87 105.82 96 63.38m222 80.33a79.13 79.13 0 0 0 16-4.48c5-1.77 9.24-5.94 10.32-11.22c-8.96 4.99-17.98 9.92-26.32 15.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Php(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M320 104.5c171.4 0 303.2 72.2 303.2 151.5S491.3 407.5 320 407.5c-171.4 0-303.2-72.2-303.2-151.5S148.7 104.5 320 104.5m0-16.8C143.3 87.7 0 163 0 256s143.3 168.3 320 168.3S640 349 640 256S496.7 87.7 320 87.7M218.2 242.5c-7.9 40.5-35.8 36.3-70.1 36.3l13.7-70.6c38 0 63.8-4.1 56.4 34.3M97.4 350.3h36.7l8.7-44.8c41.1 0 66.6 3 90.2-19.1c26.1-24 32.9-66.7 14.3-88.1c-9.7-11.2-25.3-16.7-46.5-16.7h-70.7zm185.7-213.6h36.5l-8.7 44.8c31.5 0 60.7-2.3 74.8 10.7c14.8 13.6 7.7 31-8.3 113.1h-37c15.4-79.4 18.3-86 12.7-92c-5.4-5.8-17.7-4.6-47.4-4.6l-18.8 96.6h-36.5zM505 242.5c-8 41.1-36.7 36.3-70.1 36.3l13.7-70.6c38.2 0 63.8-4.1 56.4 34.3M384.2 350.3H421l8.7-44.8c43.2 0 67.1 2.5 90.2-19.1c26.1-24 32.9-66.7 14.3-88.1c-9.7-11.2-25.3-16.7-46.5-16.7H417z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PiedPiper(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M455.93 23.2c-26.7 6.8-68.14 28.49-114.58 67.46A206 206 0 0 0 240 64C125.13 64 32 157.12 32 272s93.13 208 208 208s208-93.13 208-208a207.25 207.25 0 0 0-58.75-144.81a155.35 155.35 0 0 0-17 27.4A176.16 176.16 0 0 1 417.1 272c0 97.66-79.44 177.11-177.09 177.11a175.81 175.81 0 0 1-87.63-23.4c82.94-107.33 150.79-37.77 184.31-226.65c5.79-32.62 28-94.26 126.23-160.18c8.08-5.43 2.43-18.08-6.99-15.68M125 406.4A176.66 176.66 0 0 1 62.9 272c0-97.66 79.45-177.1 177.1-177.1a174 174 0 0 1 76.63 17.75C250.64 174.76 189.77 265.52 125 406.4" class="cls-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PiedPiperAlt(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M244 246c-3.2-2-6.3-2.9-10.1-2.9c-6.6 0-12.6 3.2-19.3 3.7l1.7 4.9zm135.9 197.9c-19 0-64.1 9.5-79.9 19.8l6.9 45.1c35.7 6.1 70.1 3.6 106-9.8c-4.8-10-23.5-55.1-33-55.1M340.8 177c6.6 2.8 11.5 9.2 22.7 22.1c2-1.4 7.5-5.2 7.5-8.6c0-4.9-11.8-13.2-13.2-23c11.2-5.7 25.2-6 37.6-8.9c68.1-16.4 116.3-52.9 146.8-116.7C548.3 29.3 554 16.1 554.6 2l-2 2.6c-28.4 50-33 63.2-81.3 100c-31.9 24.4-69.2 40.2-106.6 54.6l-6.3-.3v-21.8c-19.6 1.6-19.7-14.6-31.6-23c-18.7 20.6-31.6 40.8-58.9 51.1c-12.7 4.8-19.6 10-25.9 21.8c34.9-16.4 91.2-13.5 98.8-10M555.5 0l-.6 1.1l-.3.9l.6-.6zm-59.2 382.1c-33.9-56.9-75.3-118.4-150-115.5l-.3-6c-1.1-13.5 32.8 3.2 35.1-31l-14.4 7.2c-19.8-45.7-8.6-54.3-65.5-54.3c-14.7 0-26.7 1.7-41.4 4.6c2.9 18.6 2.2 36.7-10.9 50.3l19.5 5.5c-1.7 3.2-2.9 6.3-2.9 9.8c0 21 42.8 2.9 42.8 33.6c0 18.4-36.8 60.1-54.9 60.1c-8 0-53.7-50-53.4-60.1l.3-4.6l52.3-11.5c13-2.6 12.3-22.7-2.9-22.7c-3.7 0-43.1 9.2-49.4 10.6c-2-5.2-7.5-14.1-13.8-14.1c-3.2 0-6.3 3.2-9.5 4c-9.2 2.6-31 2.9-21.5 20.1L15.9 298.5c-5.5 1.1-8.9 6.3-8.9 11.8c0 6 5.5 10.9 11.5 10.9c8 0 131.3-28.4 147.4-32.2c2.6 3.2 4.6 6.3 7.8 8.6c20.1 14.4 59.8 85.9 76.4 85.9c24.1 0 58-22.4 71.3-41.9c3.2-4.3 6.9-7.5 12.4-6.9c.6 13.8-31.6 34.2-33 43.7c-1.4 10.2-1 35.2-.3 41.1c26.7 8.1 52-3.6 77.9-2.9c4.3-21 10.6-41.9 9.8-63.5l-.3-9.5c-1.4-34.2-10.9-38.5-34.8-58.6c-1.1-1.1-2.6-2.6-3.7-4c2.2-1.4 1.1-1 4.6-1.7c88.5 0 56.3 183.6 111.5 229.9c33.1-15 72.5-27.9 103.5-47.2c-29-25.6-52.6-45.7-72.7-79.9m-196.2 46.1v27.2l11.8-3.4l-2.9-23.8zm-68.7-150.4l24.1 61.2l21-13.8l-31.3-50.9zm84.4 154.9l2 12.4c9-1.5 58.4-6.6 58.4-14.1c0-1.4-.6-3.2-.9-4.6c-26.8 0-36.9 3.8-59.5 6.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PiedPiperHat(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 24.9c-80.8 53.6-89.4 92.5-96.4 104.4c-6.7 12.2-11.7 60.3-23.3 83.6c-11.7 23.6-54.2 42.2-66.1 50c-11.7 7.8-28.3 38.1-41.9 64.2c-108.1-4.4-167.4 38.8-259.2 93.6c29.4-9.7 43.3-16.7 43.3-16.7c94.2-36 139.3-68.3 281.1-49.2c1.1 0 1.9.6 2.8.8c3.9 2.2 5.3 6.9 3.1 10.8l-53.9 95.8c-2.5 4.7-7.8 7.2-13.1 6.1c-126.8-23.8-226.9 17.3-318.9 18.6C24.1 488 0 453.4 0 451.8c0-1.1.6-1.7 1.7-1.7c0 0 38.3 0 103.1-15.3C178.4 294.5 244 245.4 315.4 245.4c0 0 71.7 0 90.6 61.9c22.8-39.7 28.3-49.2 28.3-49.2c5.3-9.4 35-77.2 86.4-141.4c51.5-64 90.4-79.9 119.3-91.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PiedPiperPp(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M205.3 174.6c0 21.1-14.2 38.1-31.7 38.1c-7.1 0-12.8-1.2-17.2-3.7v-68c4.4-2.7 10.1-4.2 17.2-4.2c17.5 0 31.7 16.9 31.7 37.8m52.6 67c-7.1 0-12.8 1.5-17.2 4.2v68c4.4 2.5 10.1 3.7 17.2 3.7c17.4 0 31.7-16.9 31.7-37.8c0-21.1-14.3-38.1-31.7-38.1M448 80v352c0 26.5-21.5 48-48 48H48c-26.5 0-48-21.5-48-48V80c0-26.5 21.5-48 48-48h352c26.5 0 48 21.5 48 48M185 255.1c41 0 74.2-35.6 74.2-79.6c0-44-33.2-79.6-74.2-79.6c-12 0-24.1 3.2-34.6 8.8h-45.7V311l51.8-10.1v-50.6c8.6 3.1 18.1 4.8 28.5 4.8m158.4 25.3c0-44-33.2-79.6-73.9-79.6c-3.2 0-6.4.2-9.6.7c-3.7 12.5-10.1 23.8-19.2 33.4c-13.8 15-32.2 23.8-51.8 24.8V416l51.8-10.1v-50.6c8.6 3.2 18.2 4.7 28.7 4.7c40.8 0 74-35.6 74-79.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PiedPiperSquare(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M32 419L0 479.2l.8-328C.8 85.3 54 32 120 32h327.2c-93 28.9-189.9 94.2-253.9 168.6C122.7 282 82.6 338 32 419M448 32S305.2 98.8 261.6 199.1c-23.2 53.6-28.9 118.1-71 158.6c-28.9 27.8-69.8 38.2-105.3 56.3c-23.2 12-66.4 40.5-84.9 66h328.4c66 0 119.3-53.3 119.3-119.2c-.1 0-.1-328.8-.1-328.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pinterest(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M496 256c0 137-111 248-248 248c-25.6 0-50.2-3.9-73.4-11.1c10.1-16.5 25.2-43.5 30.8-65c3-11.6 15.4-59 15.4-59c8.1 15.4 31.7 28.5 56.8 28.5c74.8 0 128.7-68.8 128.7-154.3c0-81.9-66.9-143.2-152.9-143.2c-107 0-163.9 71.8-163.9 150.1c0 36.4 19.4 81.7 50.3 96.1c4.7 2.2 7.2 1.2 8.3-3.3c.8-3.4 5-20.3 6.9-28.1c.6-2.5.3-4.7-1.7-7.1c-10.1-12.5-18.3-35.3-18.3-56.6c0-54.7 41.4-107.6 112-107.6c60.9 0 103.6 41.5 103.6 100.9c0 67.1-33.9 113.6-78 113.6c-24.3 0-42.6-20.1-36.7-44.8c7-29.5 20.5-61.3 20.5-82.6c0-19-10.2-34.9-31.4-34.9c-24.9 0-44.9 25.7-44.9 60.2c0 22 7.4 36.8 7.4 36.8s-24.5 103.8-29 123.2c-5 21.4-3 51.6-.9 71.2C65.4 450.9 0 361.1 0 256C0 119 111 8 248 8s248 111 248 248"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinterestP(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M204 6.5C101.4 6.5 0 74.9 0 185.6C0 256 39.6 296 63.6 296c9.9 0 15.6-27.6 15.6-35.4c0-9.3-23.7-29.1-23.7-67.8c0-80.4 61.2-137.4 140.4-137.4c68.1 0 118.5 38.7 118.5 109.8c0 53.1-21.3 152.7-90.3 152.7c-24.9 0-46.2-18-46.2-43.8c0-37.8 26.4-74.4 26.4-113.4c0-66.2-93.9-54.2-93.9 25.8c0 16.8 2.1 35.4 9.6 50.7c-13.8 59.4-42 147.9-42 209.1c0 18.9 2.7 37.5 4.5 56.4c3.4 3.8 1.7 3.4 6.9 1.5c50.4-69 48.6-82.5 71.4-172.8c12.3 23.4 44.1 36 69.3 36c106.2 0 153.9-103.5 153.9-196.8C384 71.3 298.2 6.5 204 6.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinterestSquare(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 80v352c0 26.5-21.5 48-48 48H154.4c9.8-16.4 22.4-40 27.4-59.3c3-11.5 15.3-58.4 15.3-58.4c8 15.3 31.4 28.2 56.3 28.2c74.1 0 127.4-68.1 127.4-152.7c0-81.1-66.2-141.8-151.4-141.8c-106 0-162.2 71.1-162.2 148.6c0 36 19.2 80.8 49.8 95.1c4.7 2.2 7.1 1.2 8.2-3.3c.8-3.4 5-20.1 6.8-27.8c.6-2.5.3-4.6-1.7-7c-10.1-12.3-18.3-34.9-18.3-56c0-54.2 41-106.6 110.9-106.6c60.3 0 102.6 41.1 102.6 99.9c0 66.4-33.5 112.4-77.2 112.4c-24.1 0-42.1-19.9-36.4-44.4c6.9-29.2 20.3-60.7 20.3-81.8c0-53-75.5-45.7-75.5 25c0 21.7 7.3 36.5 7.3 36.5c-31.4 132.8-36.1 134.5-29.6 192.6l2.2.8H48c-26.5 0-48-21.5-48-48V80c0-26.5 21.5-48 48-48h352c26.5 0 48 21.5 48 48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Playstation(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M570.9 372.3c-11.3 14.2-38.8 24.3-38.8 24.3L327 470.2v-54.3l150.9-53.8c17.1-6.1 19.8-14.8 5.8-19.4c-13.9-4.6-39.1-3.3-56.2 2.9L327 381.1v-56.4c23.2-7.8 47.1-13.6 75.7-16.8c40.9-4.5 90.9.6 130.2 15.5c44.2 14 49.2 34.7 38 48.9m-224.4-92.5v-139c0-16.3-3-31.3-18.3-35.6c-11.7-3.8-19 7.1-19 23.4v347.9l-93.8-29.8V32c39.9 7.4 98 24.9 129.2 35.4C424.1 94.7 451 128.7 451 205.2c0 74.5-46 102.8-104.5 74.6M43.2 410.2c-45.4-12.8-53-39.5-32.3-54.8c19.1-14.2 51.7-24.9 51.7-24.9l134.5-47.8v54.5l-96.8 34.6c-17.1 6.1-19.7 14.8-5.8 19.4c13.9 4.6 39.1 3.3 56.2-2.9l46.4-16.9v48.8c-51.6 9.3-101.4 7.3-153.9-10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ProductHunt(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M326.3 218.8c0 20.5-16.7 37.2-37.2 37.2h-70.3v-74.4h70.3c20.5 0 37.2 16.7 37.2 37.2M504 256c0 137-111 248-248 248S8 393 8 256S119 8 256 8s248 111 248 248m-128.1-37.2c0-47.9-38.9-86.8-86.8-86.8H169.2v248h49.6v-74.4h70.3c47.9 0 86.8-38.9 86.8-86.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pushed(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m407 111.9l-98.5-9l14-33.4c10.4-23.5-10.8-40.4-28.7-37L22.5 76.9c-15.1 2.7-26 18.3-21.4 36.6l105.1 348.3c6.5 21.3 36.7 24.2 47.7 7l35.3-80.8l235.2-231.3c16.4-16.8 4.3-42.9-17.4-44.8M297.6 53.6c5.1-.7 7.5 2.5 5.2 7.4L286 100.9L108.6 84.6zM22.7 107.9c-3.1-5.1 1-10 6.1-9.1l248.7 22.7l-96.9 230.7zM136 456.4c-2.6 4-7.9 3.1-9.4-1.2L43.5 179.7l127.7 197.6c-7 15-35.2 79.1-35.2 79.1m272.8-314.5L210.1 337.3l89.7-213.7l106.4 9.7c4 1.1 5.7 5.3 2.6 8.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Python(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M439.8 200.5c-7.7-30.9-22.3-54.2-53.4-54.2h-40.1v47.4c0 36.8-31.2 67.8-66.8 67.8H172.7c-29.2 0-53.4 25-53.4 54.3v101.8c0 29 25.2 46 53.4 54.3c33.8 9.9 66.3 11.7 106.8 0c26.9-7.8 53.4-23.5 53.4-54.3v-40.7H226.2v-13.6h160.2c31.1 0 42.6-21.7 53.4-54.2c11.2-33.5 10.7-65.7 0-108.6M286.2 404c11.1 0 20.1 9.1 20.1 20.3c0 11.3-9 20.4-20.1 20.4c-11 0-20.1-9.2-20.1-20.4c.1-11.3 9.1-20.3 20.1-20.3M167.8 248.1h106.8c29.7 0 53.4-24.5 53.4-54.3V91.9c0-29-24.4-50.7-53.4-55.6c-35.8-5.9-74.7-5.6-106.8.1c-45.2 8-53.4 24.7-53.4 55.6v40.7h106.9v13.6h-147c-31.1 0-58.3 18.7-66.8 54.2c-9.8 40.7-10.2 66.1 0 108.6c7.6 31.6 25.7 54.2 56.8 54.2H101v-48.8c0-35.3 30.5-66.4 66.8-66.4m-6.7-142.6c-11.1 0-20.1-9.1-20.1-20.3c.1-11.3 9-20.4 20.1-20.4c11 0 20.1 9.2 20.1 20.4s-9 20.3-20.1 20.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Qq(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M433.754 420.445c-11.526 1.393-44.86-52.741-44.86-52.741c0 31.345-16.136 72.247-51.051 101.786c16.842 5.192 54.843 19.167 45.803 34.421c-7.316 12.343-125.51 7.881-159.632 4.037c-34.122 3.844-152.316 8.306-159.632-4.037c-9.045-15.25 28.918-29.214 45.783-34.415c-34.92-29.539-51.059-70.445-51.059-101.792c0 0-33.334 54.134-44.859 52.741c-5.37-.65-12.424-29.644 9.347-99.704c10.261-33.024 21.995-60.478 40.144-105.779C60.683 98.063 108.982.006 224 0c113.737.006 163.156 96.133 160.264 214.963c18.118 45.223 29.912 72.85 40.144 105.778c21.768 70.06 14.716 99.053 9.346 99.704"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Quinscape(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M313.6 474.6h-1a158.1 158.1 0 0 1 0-316.2c94.9 0 168.2 83.1 157 176.6c4 5.1 8.2 9.6 11.2 15.3c13.4-30.3 20.3-62.4 20.3-97.7C501.1 117.5 391.6 8 256.5 8S12 117.5 12 252.6s109.5 244.6 244.5 244.6a237.36 237.36 0 0 0 70.4-10.1c-5.2-3.5-8.9-8.1-13.3-12.5m-.1-.1l.4.1zm78.4-168.9a99.2 99.2 0 1 0 99.2 99.2a99.18 99.18 0 0 0-99.2-99.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Quora(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M440.5 386.7h-29.3c-1.5 13.5-10.5 30.8-33 30.8c-20.5 0-35.3-14.2-49.5-35.8c44.2-34.2 74.7-87.5 74.7-153C403.5 111.2 306.8 32 205 32C105.3 32 7.3 111.7 7.3 228.7c0 134.1 131.3 221.6 249 189C276 451.3 302 480 351.5 480c81.8 0 90.8-75.3 89-93.3M297 329.2C277.5 300 253.3 277 205.5 277c-30.5 0-54.3 10-69 22.8l12.2 24.3c6.2-3 13-4 19.8-4c35.5 0 53.7 30.8 69.2 61.3c-10 3-20.7 4.2-32.7 4.2c-75 0-107.5-53-107.5-156.7C97.5 124.5 130 71 205 71c76.2 0 108.7 53.5 108.7 157.7c.1 41.8-5.4 75.6-16.7 100.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rproject(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M581 226.6C581 119.1 450.9 32 290.5 32S0 119.1 0 226.6C0 322.4 103.3 402 239.4 418.1V480h99.1v-61.5c24.3-2.7 47.6-7.4 69.4-13.9L448 480h112l-67.4-113.7c54.5-35.4 88.4-84.9 88.4-139.7m-466.8 14.5c0-73.5 98.9-133 220.8-133s211.9 40.7 211.9 133c0 50.1-26.5 85-70.3 106.4c-2.4-1.6-4.7-2.9-6.4-3.7c-10.2-5.2-27.8-10.5-27.8-10.5s86.6-6.4 86.6-92.7s-90.6-87.9-90.6-87.9h-199V361c-74.1-21.5-125.2-67.1-125.2-119.9m225.1 38.3v-55.6c57.8 0 87.8-6.8 87.8 27.3c0 36.5-38.2 28.3-87.8 28.3m-.9 72.5H365c10.8 0 18.9 11.7 24 19.2c-16.1 1.9-33 2.8-50.6 2.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RaspberryPi(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m372 232.5l-3.7-6.5c.1-46.4-21.4-65.3-46.5-79.7c7.6-2 15.4-3.6 17.6-13.2c13.1-3.3 15.8-9.4 17.1-15.8c3.4-2.3 14.8-8.7 13.6-19.7c6.4-4.4 10-10.1 8.1-18.1c6.9-7.5 8.7-13.7 5.8-19.4c8.3-10.3 4.6-15.6 1.1-20.9c6.2-11.2.7-23.2-16.6-21.2c-6.9-10.1-21.9-7.8-24.2-7.8c-2.6-3.2-6-6-16.5-4.7c-6.8-6.1-14.4-5-22.3-2.1c-9.3-7.3-15.5-1.4-22.6.8C271.6.6 269 5.5 263.5 7.6c-12.3-2.6-16.1 3-22 8.9l-6.9-.1c-18.6 10.8-27.8 32.8-31.1 44.1c-3.3-11.3-12.5-33.3-31.1-44.1l-6.9.1c-5.9-5.9-9.7-11.5-22-8.9c-5.6-2-8.1-7-19.4-3.4c-4.6-1.4-8.9-4.4-13.9-4.3c-2.6.1-5.5 1-8.7 3.5c-7.9-3-15.5-4-22.3 2.1c-10.5-1.3-14 1.4-16.5 4.7c-2.3 0-17.3-2.3-24.2 7.8C21.2 16 15.8 28 22 39.2c-3.5 5.4-7.2 10.7 1.1 20.9c-2.9 5.7-1.1 11.9 5.8 19.4c-1.8 8 1.8 13.7 8.1 18.1c-1.2 11 10.2 17.4 13.6 19.7c1.3 6.4 4 12.4 17.1 15.8c2.2 9.5 10 11.2 17.6 13.2c-25.1 14.4-46.6 33.3-46.5 79.7l-3.7 6.5c-28.8 17.2-54.7 72.7-14.2 117.7c2.6 14.1 7.1 24.2 11 35.4c5.9 45.2 44.5 66.3 54.6 68.8c14.9 11.2 30.8 21.8 52.2 29.2C159 504.2 181 512 203 512h1c22.1 0 44-7.8 64.2-28.4c21.5-7.4 37.3-18 52.2-29.2c10.2-2.5 48.7-23.6 54.6-68.8c3.9-11.2 8.4-21.3 11-35.4c40.6-45.1 14.7-100.5-14-117.7m-22.2-8c-1.5 18.7-98.9-65.1-82.1-67.9c45.7-7.5 83.6 19.2 82.1 67.9m-43 93.1c-24.5 15.8-59.8 5.6-78.8-22.8s-14.6-64.2 9.9-80c24.5-15.8 59.8-5.6 78.8 22.8s14.6 64.2-9.9 80M238.9 29.3c.8 4.2 1.8 6.8 2.9 7.6c5.4-5.8 9.8-11.7 16.8-17.3c0 3.3-1.7 6.8 2.5 9.4c3.7-5 8.8-9.5 15.5-13.3c-3.2 5.6-.6 7.3 1.2 9.6c5.1-4.4 10-8.8 19.4-12.3c-2.6 3.1-6.2 6.2-2.4 9.8c5.3-3.3 10.6-6.6 23.1-8.9c-2.8 3.1-8.7 6.3-5.1 9.4c6.6-2.5 14-4.4 22.1-5.4c-3.9 3.2-7.1 6.3-3.9 8.8c7.1-2.2 16.9-5.1 26.4-2.6l-6 6.1c-.7.8 14.1.6 23.9.8c-3.6 5-7.2 9.7-9.3 18.2c1 1 5.8.4 10.4 0c-4.7 9.9-12.8 12.3-14.7 16.6c2.9 2.2 6.8 1.6 11.2.1c-3.4 6.9-10.4 11.7-16 17.3c1.4 1 3.9 1.6 9.7.9c-5.2 5.5-11.4 10.5-18.8 15c1.3 1.5 5.8 1.5 10 1.6c-6.7 6.5-15.3 9.9-23.4 14.2c4 2.7 6.9 2.1 10 2.1c-5.7 4.7-15.4 7.1-24.4 10c1.7 2.7 3.4 3.4 7.1 4.1c-9.5 5.3-23.2 2.9-27 5.6c.9 2.7 3.6 4.4 6.7 5.8c-15.4.9-57.3-.6-65.4-32.3c15.7-17.3 44.4-37.5 93.7-62.6c-38.4 12.8-73 30-102 53.5c-34.3-15.9-10.8-55.9 5.8-71.8m-34.4 114.6c24.2-.3 54.1 17.8 54 34.7c-.1 15-21 27.1-53.8 26.9c-32.1-.4-53.7-15.2-53.6-29.8c0-11.9 26.2-32.5 53.4-31.8m-123-12.8c3.7-.7 5.4-1.5 7.1-4.1c-9-2.8-18.7-5.3-24.4-10c3.1 0 6 .7 10-2.1c-8.1-4.3-16.7-7.7-23.4-14.2c4.2-.1 8.7 0 10-1.6c-7.4-4.5-13.6-9.5-18.8-15c5.8.7 8.3.1 9.7-.9c-5.6-5.6-12.7-10.4-16-17.3c4.3 1.5 8.3 2 11.2-.1c-1.9-4.2-10-6.7-14.7-16.6c4.6.4 9.4 1 10.4 0c-2.1-8.5-5.8-13.3-9.3-18.2c9.8-.1 24.6 0 23.9-.8l-6-6.1c9.5-2.5 19.3.4 26.4 2.6c3.2-2.5-.1-5.6-3.9-8.8c8.1 1.1 15.4 2.9 22.1 5.4c3.5-3.1-2.3-6.3-5.1-9.4c12.5 2.3 17.8 5.6 23.1 8.9c3.8-3.6.2-6.7-2.4-9.8c9.4 3.4 14.3 7.9 19.4 12.3c1.7-2.3 4.4-4 1.2-9.6c6.7 3.8 11.8 8.3 15.5 13.3c4.1-2.6 2.5-6.2 2.5-9.4c7 5.6 11.4 11.5 16.8 17.3c1.1-.8 2-3.4 2.9-7.6c16.6 15.9 40.1 55.9 6 71.8c-29-23.5-63.6-40.7-102-53.5c49.3 25 78 45.3 93.7 62.6c-8 31.8-50 33.2-65.4 32.3c3.1-1.4 5.8-3.2 6.7-5.8c-4-2.8-17.6-.4-27.2-5.6m60.1 24.1c16.8 2.8-80.6 86.5-82.1 67.9c-1.5-48.7 36.5-75.5 82.1-67.9M38.2 342c-23.7-18.8-31.3-73.7 12.6-98.3c26.5-7 9 107.8-12.6 98.3m91 98.2c-13.3 7.9-45.8 4.7-68.8-27.9c-15.5-27.4-13.5-55.2-2.6-63.4c16.3-9.8 41.5 3.4 60.9 25.6c16.9 20 24.6 55.3 10.5 65.7m-26.4-119.7c-24.5-15.8-28.9-51.6-9.9-80s54.3-38.6 78.8-22.8s28.9 51.6 9.9 80c-19.1 28.4-54.4 38.6-78.8 22.8M205 496c-29.4 1.2-58.2-23.7-57.8-32.3c-.4-12.7 35.8-22.6 59.3-22c23.7-1 55.6 7.5 55.7 18.9c.5 11-28.8 35.9-57.2 35.4m58.9-124.9c.2 29.7-26.2 53.8-58.8 54c-32.6.2-59.2-23.8-59.4-53.4v-.6c-.2-29.7 26.2-53.8 58.8-54c32.6-.2 59.2 23.8 59.4 53.4zm82.2 42.7c-25.3 34.6-59.6 35.9-72.3 26.3c-13.3-12.4-3.2-50.9 15.1-72c20.9-23.3 43.3-38.5 58.9-26.6c10.5 10.3 16.7 49.1-1.7 72.3m22.9-73.2c-21.5 9.4-39-105.3-12.6-98.3c43.9 24.7 36.3 79.6 12.6 98.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ravelry(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M498.252 234.223c-1.208-10.34-1.7-20.826-3.746-31a310.306 310.306 0 0 0-9.622-36.6a184.068 184.068 0 0 0-30.874-57.5a251.154 251.154 0 0 0-18.818-21.689a237.362 237.362 0 0 0-47.113-36.116a240.8 240.8 0 0 0-56.723-24.668c-11.018-3.1-22.272-5.431-33.515-7.615c-6.78-1.314-13.749-1.667-20.627-2.482c-.316-.036-.6-.358-.9-.553q-16.143.009-32.288.006c-2.41.389-4.808.925-7.236 1.15a179.331 179.331 0 0 0-34.256 7.1a221.5 221.5 0 0 0-39.768 16.355a281.385 281.385 0 0 0-38.08 24.158c-6.167 4.61-12.268 9.36-17.974 14.518c-10.173 9.207-20.372 18.433-29.927 28.268a243.878 243.878 0 0 0-33.648 43.95a206.488 206.488 0 0 0-20.494 44.6a198.2 198.2 0 0 0-7.691 34.759a201.13 201.13 0 0 0-1.552 35.521a299.716 299.716 0 0 0 4.425 40.24a226.865 226.865 0 0 0 16.73 53.3a210.543 210.543 0 0 0 24 39.528a213.589 213.589 0 0 0 26.358 28.416a251.313 251.313 0 0 0 41.787 30.586a287.831 287.831 0 0 0 55.9 25.277a269.5 269.5 0 0 0 40.641 9.835c6.071 1.01 12.275 1.253 18.412 1.873a4.149 4.149 0 0 1 1.19.56h32.289c2.507-.389 5-.937 7.527-1.143c16.336-1.332 32.107-5.335 47.489-10.717a219.992 219.992 0 0 0 48.952-23.818c9.749-6.447 19.395-13.077 28.737-20.1c5.785-4.348 10.988-9.5 16.3-14.457c3.964-3.7 7.764-7.578 11.51-11.5a232.162 232.162 0 0 0 31.427-41.639c9.542-16.045 17.355-32.905 22.3-50.926c2.859-10.413 4.947-21.045 7.017-31.652c1.032-5.279 1.251-10.723 1.87-16.087c.036-.317.358-.6.552-.9v-37.056a9.757 9.757 0 0 1-.561-1.782m-161.117-1.15s-16.572-2.98-28.47-2.98c-27.2 0-33.57 14.9-33.57 37.04V360.8h-73.513V170.062H275.1v31.931c8.924-26.822 26.771-36.189 62.04-36.189Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func React(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M418.2 177.2c-5.4-1.8-10.8-3.5-16.2-5.1c.9-3.7 1.7-7.4 2.5-11.1c12.3-59.6 4.2-107.5-23.1-123.3c-26.3-15.1-69.2.6-112.6 38.4c-4.3 3.7-8.5 7.6-12.5 11.5c-2.7-2.6-5.5-5.2-8.3-7.7c-45.5-40.4-91.1-57.4-118.4-41.5c-26.2 15.2-34 60.3-23 116.7c1.1 5.6 2.3 11.1 3.7 16.7c-6.4 1.8-12.7 3.8-18.6 5.9C38.3 196.2 0 225.4 0 255.6c0 31.2 40.8 62.5 96.3 81.5c4.5 1.5 9 3 13.6 4.3c-1.5 6-2.8 11.9-4 18c-10.5 55.5-2.3 99.5 23.9 114.6c27 15.6 72.4-.4 116.6-39.1c3.5-3.1 7-6.3 10.5-9.7c4.4 4.3 9 8.4 13.6 12.4c42.8 36.8 85.1 51.7 111.2 36.6c27-15.6 35.8-62.9 24.4-120.5c-.9-4.4-1.9-8.9-3-13.5c3.2-.9 6.3-1.9 9.4-2.9c57.7-19.1 99.5-50 99.5-81.7c0-30.3-39.4-59.7-93.8-78.4M282.9 92.3c37.2-32.4 71.9-45.1 87.7-36c16.9 9.7 23.4 48.9 12.8 100.4c-.7 3.4-1.4 6.7-2.3 10c-22.2-5-44.7-8.6-67.3-10.6c-13-18.6-27.2-36.4-42.6-53.1c3.9-3.7 7.7-7.2 11.7-10.7M167.2 307.5c5.1 8.7 10.3 17.4 15.8 25.9c-15.6-1.7-31.1-4.2-46.4-7.5c4.4-14.4 9.9-29.3 16.3-44.5c4.6 8.8 9.3 17.5 14.3 26.1m-30.3-120.3c14.4-3.2 29.7-5.8 45.6-7.8c-5.3 8.3-10.5 16.8-15.4 25.4c-4.9 8.5-9.7 17.2-14.2 26c-6.3-14.9-11.6-29.5-16-43.6m27.4 68.9c6.6-13.8 13.8-27.3 21.4-40.6s15.8-26.2 24.4-38.9c15-1.1 30.3-1.7 45.9-1.7s31 .6 45.9 1.7c8.5 12.6 16.6 25.5 24.3 38.7s14.9 26.7 21.7 40.4c-6.7 13.8-13.9 27.4-21.6 40.8c-7.6 13.3-15.7 26.2-24.2 39c-14.9 1.1-30.4 1.6-46.1 1.6s-30.9-.5-45.6-1.4c-8.7-12.7-16.9-25.7-24.6-39s-14.8-26.8-21.5-40.6m180.6 51.2c5.1-8.8 9.9-17.7 14.6-26.7c6.4 14.5 12 29.2 16.9 44.3c-15.5 3.5-31.2 6.2-47 8c5.4-8.4 10.5-17 15.5-25.6m14.4-76.5c-4.7-8.8-9.5-17.6-14.5-26.2c-4.9-8.5-10-16.9-15.3-25.2c16.1 2 31.5 4.7 45.9 8c-4.6 14.8-10 29.2-16.1 43.4M256.2 118.3c10.5 11.4 20.4 23.4 29.6 35.8c-19.8-.9-39.7-.9-59.5 0c9.8-12.9 19.9-24.9 29.9-35.8M140.2 57c16.8-9.8 54.1 4.2 93.4 39c2.5 2.2 5 4.6 7.6 7c-15.5 16.7-29.8 34.5-42.9 53.1c-22.6 2-45 5.5-67.2 10.4c-1.3-5.1-2.4-10.3-3.5-15.5c-9.4-48.4-3.2-84.9 12.6-94m-24.5 263.6c-4.2-1.2-8.3-2.5-12.4-3.9c-21.3-6.7-45.5-17.3-63-31.2c-10.1-7-16.9-17.8-18.8-29.9c0-18.3 31.6-41.7 77.2-57.6c5.7-2 11.5-3.8 17.3-5.5c6.8 21.7 15 43 24.5 63.6c-9.6 20.9-17.9 42.5-24.8 64.5m116.6 98c-16.5 15.1-35.6 27.1-56.4 35.3c-11.1 5.3-23.9 5.8-35.3 1.3c-15.9-9.2-22.5-44.5-13.5-92c1.1-5.6 2.3-11.2 3.7-16.7c22.4 4.8 45 8.1 67.9 9.8c13.2 18.7 27.7 36.6 43.2 53.4c-3.2 3.1-6.4 6.1-9.6 8.9m24.5-24.3c-10.2-11-20.4-23.2-30.3-36.3c9.6.4 19.5.6 29.5.6c10.3 0 20.4-.2 30.4-.7c-9.2 12.7-19.1 24.8-29.6 36.4m130.7 30c-.9 12.2-6.9 23.6-16.5 31.3c-15.9 9.2-49.8-2.8-86.4-34.2c-4.2-3.6-8.4-7.5-12.7-11.5c15.3-16.9 29.4-34.8 42.2-53.6c22.9-1.9 45.7-5.4 68.2-10.5c1 4.1 1.9 8.2 2.7 12.2c4.9 21.6 5.7 44.1 2.5 66.3m18.2-107.5c-2.8.9-5.6 1.8-8.5 2.6c-7-21.8-15.6-43.1-25.5-63.8c9.6-20.4 17.7-41.4 24.5-62.9c5.2 1.5 10.2 3.1 15 4.7c46.6 16 79.3 39.8 79.3 58c0 19.6-34.9 44.9-84.8 61.4m-149.7-15c25.3 0 45.8-20.5 45.8-45.8s-20.5-45.8-45.8-45.8c-25.3 0-45.8 20.5-45.8 45.8s20.5 45.8 45.8 45.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reacteurope(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m250.6 211.74l5.8-4.1l5.8 4.1l-2.1-6.8l5.7-4.3l-7.1-.1l-2.3-6.8l-2.3 6.8l-7.2.1l5.7 4.3zm63.7 0l5.8-4.1l5.8 4.1l-2.1-6.8l5.7-4.3l-7.2-.1l-2.3-6.8l-2.3 6.8l-7.2.1l5.7 4.3zm-91.3 50.5h-3.4c-4.8 0-3.8 4-3.8 12.1c0 4.7-2.3 6.1-5.8 6.1s-5.8-1.4-5.8-6.1v-36.6c0-4.7 2.3-6.1 5.8-6.1s5.8 1.4 5.8 6.1c0 7.2-.7 10.5 3.8 10.5h3.4c4.7-.1 3.8-3.9 3.8-12.3c0-9.9-6.7-14.1-16.8-14.1h-.2c-10.1 0-16.8 4.2-16.8 14.1V276c0 10.4 6.7 14.1 16.8 14.1h.2c10.1 0 16.8-3.8 16.8-14.1c0-9.86 1.1-13.76-3.8-13.76m-80.7 17.4h-14.7v-19.3H139c2.5 0 3.8-1.3 3.8-3.8v-2.1c0-2.5-1.3-3.8-3.8-3.8h-11.4v-18.3H142c2.5 0 3.8-1.3 3.8-3.8v-2.1c0-2.5-1.3-3.8-3.8-3.8h-21.7c-2.4-.1-3.7 1.3-3.7 3.8v59.1c0 2.5 1.3 3.8 3.8 3.8h21.9c2.5 0 3.8-1.3 3.8-3.8v-2.1c0-2.5-1.3-3.8-3.8-3.8m-42-18.5c4.6-2 7.3-6 7.3-12.4v-11.9c0-10.1-6.7-14.1-16.8-14.1H77.4c-2.5 0-3.8 1.3-3.8 3.8v59.1c0 2.5 1.3 3.8 3.8 3.8h3.4c2.5 0 3.8-1.3 3.8-3.8v-22.9h5.6l7.4 23.5a4.1 4.1 0 0 0 4.3 3.2h3.3c2.8 0 4-1.8 3.2-4.4zm-3.8-14c0 4.8-2.5 6.1-6.1 6.1h-5.8v-20.9h5.8c3.6 0 6.1 1.3 6.1 6.1zM176 226a3.82 3.82 0 0 0-4.2-3.4h-6.9a3.68 3.68 0 0 0-4 3.4l-11 59.2c-.5 2.7.9 4.1 3.4 4.1h3a3.74 3.74 0 0 0 4.1-3.5l1.8-11.3h12.2l1.8 11.3a3.74 3.74 0 0 0 4.1 3.5h3.5c2.6 0 3.9-1.4 3.4-4.1zm-12.3 39.3l4.7-29.7l4.7 29.7zm89.3 20.2v-53.2h7.5c2.5 0 3.8-1.3 3.8-3.8v-2.1c0-2.5-1.3-3.8-3.8-3.8h-25.8c-2.5 0-3.8 1.3-3.8 3.8v2.1c0 2.5 1.3 3.8 3.8 3.8h7.3v53.2c0 2.5 1.3 3.8 3.8 3.8h3.4c2.5.04 3.8-1.3 3.8-3.76zm248-.8h-19.4V258h16.1a1.89 1.89 0 0 0 2-2v-.8a1.89 1.89 0 0 0-2-2h-16.1v-25.8h19.1a1.89 1.89 0 0 0 2-2v-.8a1.77 1.77 0 0 0-2-1.9h-22.2a1.62 1.62 0 0 0-2 1.8v63a1.81 1.81 0 0 0 2 1.9H501a1.81 1.81 0 0 0 2-1.9v-.8a1.84 1.84 0 0 0-2-1.96zm-93.1-62.9h-.8c-10.1 0-15.3 4.7-15.3 14.1V276c0 9.3 5.2 14.1 15.3 14.1h.8c10.1 0 15.3-4.8 15.3-14.1v-40.1c0-9.36-5.2-14.06-15.3-14.06zm10.2 52.4c-.1 8-3 11.1-10.5 11.1s-10.5-3.1-10.5-11.1v-36.6c0-7.9 3-11.1 10.5-11.1s10.5 3.2 10.5 11.1zm-46.5-14.5c6.1-1.6 9.2-6.1 9.2-13.3v-9.7c0-9.4-5.2-14.1-15.3-14.1h-13.7a1.81 1.81 0 0 0-2 1.9v63a1.81 1.81 0 0 0 2 1.9h1.2a1.74 1.74 0 0 0 1.9-1.9v-26.9h11.6l10.4 27.2a2.32 2.32 0 0 0 2.3 1.5h1.5c1.4 0 2-1 1.5-2.3zm-6.4-3.9H355v-28.5h10.2c7.5 0 10.5 3.1 10.5 11.1v6.4c0 7.84-3 11.04-10.5 11.04zm85.9-33.1h-13.7a1.62 1.62 0 0 0-2 1.8v63a1.81 1.81 0 0 0 2 1.9h1.2a1.74 1.74 0 0 0 1.9-1.9v-26.1h10.6c10.1 0 15.3-4.8 15.3-14.1v-10.5c0-9.4-5.2-14.1-15.3-14.1m10.2 22.8c0 7.9-3 11.1-10.5 11.1h-10.2v-29.2h10.2c7.5-.1 10.5 3.1 10.5 11zM259.5 308l-2.3-6.8l-2.3 6.8l-7.1.1l5.7 4.3l-2.1 6.8l5.8-4.1l5.8 4.1l-2.1-6.8l5.7-4.3zm227.6-136.1a364.42 364.42 0 0 0-35.6-11.3c19.6-78 11.6-134.7-22.3-153.9C394.7-12.66 343.3 11 291 61.94q5.1 4.95 10.2 10.2c82.5-80 119.6-53.5 120.9-52.8c22.4 12.7 36 55.8 15.5 137.8a587.83 587.83 0 0 0-84.6-13C281.1 43.64 212.4 2 170.8 2C140 2 127 23 123.2 29.74c-18.1 32-13.3 84.2.1 133.8c-70.5 20.3-120.7 54.1-120.3 95c.5 59.6 103.2 87.8 122.1 92.8c-20.5 81.9-10.1 135.6 22.3 153.9c28 15.8 75.1 6 138.2-55.2q-5.1-4.95-10.2-10.2c-82.5 80-119.7 53.5-120.9 52.8c-22.3-12.6-36-55.6-15.5-137.9c12.4 2.9 41.8 9.5 84.6 13c71.9 100.4 140.6 142 182.1 142c30.8 0 43.8-21 47.6-27.7c18-31.9 13.3-84.1-.1-133.8c152.3-43.8 156.2-130.2 33.9-176.3zM135.9 36.84c2.9-5.1 11.9-20.3 34.9-20.3c36.8 0 98.8 39.6 163.3 126.2a714 714 0 0 0-93.9.9a547.76 547.76 0 0 1 42.2-52.4Q277.3 86 272.2 81a598.25 598.25 0 0 0-50.7 64.2a569.69 569.69 0 0 0-84.4 14.6c-.2-1.4-24.3-82.2-1.2-123zm304.8 438.3c-2.9 5.1-11.8 20.3-34.9 20.3c-36.7 0-98.7-39.4-163.3-126.2a695.38 695.38 0 0 0 93.9-.9a547.76 547.76 0 0 1-42.2 52.4q5.1 5.25 10.2 10.2a588.47 588.47 0 0 0 50.7-64.2c47.3-4.7 80.3-13.5 84.4-14.6c22.7 84.4 4.5 117 1.2 123m9.1-138.6c-3.6-11.9-7.7-24.1-12.4-36.4a12.67 12.67 0 0 1-10.7-5.7l-.1.1a19.61 19.61 0 0 1-5.4 3.6c5.7 14.3 10.6 28.4 14.7 42.2a535.3 535.3 0 0 1-72 13c3.5-5.3 17.2-26.2 32.2-54.2a24.6 24.6 0 0 1-6-3.2c-1.1 1.2-3.6 4.2-10.9 4.2c-6.2 11.2-17.4 30.9-33.9 55.2a711.91 711.91 0 0 1-112.4 1c-7.9-11.2-21.5-31.1-36.8-57.8a21 21 0 0 1-3-1.5c-1.9 1.6-3.9 3.2-12.6 3.2c6.3 11.2 17.5 30.7 33.8 54.6a548.81 548.81 0 0 1-72.2-11.7q5.85-21 14.1-42.9c-3.2 0-5.4.2-8.4-1a17.58 17.58 0 0 1-6.9 1c-4.9 13.4-9.1 26.5-12.7 39.4C-31.7 297-12.1 216 126.7 175.64c3.6 11.9 7.7 24.1 12.4 36.4c10.4 0 12.9 3.4 14.4 5.3a12 12 0 0 1 2.3-2.2c-5.8-14.7-10.9-29.2-15.2-43.3c7-1.8 32.4-8.4 72-13c-15.9 24.3-26.7 43.9-32.8 55.3a14.22 14.22 0 0 1 6.4 8a23.42 23.42 0 0 1 10.2-8.4c6.5-11.7 17.9-31.9 34.8-56.9a711.72 711.72 0 0 1 112.4-1c31.5 44.6 28.9 48.1 42.5 64.5a21.42 21.42 0 0 1 10.4-7.4c-6.4-11.4-17.6-31-34.3-55.5c40.4 4.1 65 10 72.2 11.7c-4 14.4-8.9 29.2-14.6 44.2a20.74 20.74 0 0 1 6.8 4.3l.1.1a12.72 12.72 0 0 1 8.9-5.6c4.9-13.4 9.2-26.6 12.8-39.5a359.71 359.71 0 0 1 34.5 11c106.1 39.9 74 87.9 72.6 90.4c-19.8 35.1-80.1 55.2-105.7 62.5m-114.4-114h-1.2a1.74 1.74 0 0 0-1.9 1.9v49.8c0 7.9-2.6 11.1-10.1 11.1s-10.1-3.1-10.1-11.1v-49.8a1.69 1.69 0 0 0-1.9-1.9H309a1.81 1.81 0 0 0-2 1.9v51.5c0 9.6 5 14.1 15.1 14.1h.4c10.1 0 15.1-4.6 15.1-14.1v-51.5a2 2 0 0 0-2.2-1.9M321.7 308l-2.3-6.8l-2.3 6.8l-7.1.1l5.7 4.3l-2.1 6.8l5.8-4.1l5.8 4.1l-2.1-6.8l5.7-4.3zm-31.1 7.4l-2.3-6.8l-2.3 6.8l-7.1.1l5.7 4.3l-2.1 6.8l5.8-4.1l5.8 4.1l-2.1-6.8l5.7-4.3zm5.1-30.8h-19.4v-26.7h16.1a1.89 1.89 0 0 0 2-2v-.8a1.89 1.89 0 0 0-2-2h-16.1v-25.8h19.1a1.89 1.89 0 0 0 2-2v-.8a1.77 1.77 0 0 0-2-1.9h-22.2a1.81 1.81 0 0 0-2 1.9v63a1.81 1.81 0 0 0 2 1.9h22.5a1.77 1.77 0 0 0 2-1.9v-.8a1.83 1.83 0 0 0-2-2.06zm-7.4-99.4L286 192l-7.1.1l5.7 4.3l-2.1 6.8l5.8-4.1l5.8 4.1l-2.1-6.8l5.7-4.3l-7.1-.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Readme(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M528.3 46.5H388.5c-48.1 0-89.9 33.3-100.4 80.3c-10.6-47-52.3-80.3-100.4-80.3H48c-26.5 0-48 21.5-48 48v245.8c0 26.5 21.5 48 48 48h89.7c102.2 0 132.7 24.4 147.3 75c.7 2.8 5.2 2.8 6 0c14.7-50.6 45.2-75 147.3-75H528c26.5 0 48-21.5 48-48V94.6c0-26.4-21.3-47.9-47.7-48.1M242 311.9c0 1.9-1.5 3.5-3.5 3.5H78.2c-1.9 0-3.5-1.5-3.5-3.5V289c0-1.9 1.5-3.5 3.5-3.5h160.4c1.9 0 3.5 1.5 3.5 3.5v22.9zm0-60.9c0 1.9-1.5 3.5-3.5 3.5H78.2c-1.9 0-3.5-1.5-3.5-3.5v-22.9c0-1.9 1.5-3.5 3.5-3.5h160.4c1.9 0 3.5 1.5 3.5 3.5V251zm0-60.9c0 1.9-1.5 3.5-3.5 3.5H78.2c-1.9 0-3.5-1.5-3.5-3.5v-22.9c0-1.9 1.5-3.5 3.5-3.5h160.4c1.9 0 3.5 1.5 3.5 3.5v22.9zm259.3 121.7c0 1.9-1.5 3.5-3.5 3.5H337.5c-1.9 0-3.5-1.5-3.5-3.5v-22.9c0-1.9 1.5-3.5 3.5-3.5h160.4c1.9 0 3.5 1.5 3.5 3.5v22.9zm0-60.9c0 1.9-1.5 3.5-3.5 3.5H337.5c-1.9 0-3.5-1.5-3.5-3.5V228c0-1.9 1.5-3.5 3.5-3.5h160.4c1.9 0 3.5 1.5 3.5 3.5v22.9zm0-60.9c0 1.9-1.5 3.5-3.5 3.5H337.5c-1.9 0-3.5-1.5-3.5-3.5v-22.8c0-1.9 1.5-3.5 3.5-3.5h160.4c1.9 0 3.5 1.5 3.5 3.5V190z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rebel(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256.5 504C117.2 504 9 387.8 13.2 249.9C16 170.7 56.4 97.7 129.7 49.5c.3 0 1.9-.6 1.1.8c-5.8 5.5-111.3 129.8-14.1 226.4c49.8 49.5 90 2.5 90 2.5c38.5-50.1-.6-125.9-.6-125.9c-10-24.9-45.7-40.1-45.7-40.1l28.8-31.8c24.4 10.5 43.2 38.7 43.2 38.7c.8-29.6-21.9-61.4-21.9-61.4L255.1 8l44.3 50.1c-20.5 28.8-21.9 62.6-21.9 62.6c13.8-23 43.5-39.3 43.5-39.3l28.5 31.8c-27.4 8.9-45.4 39.9-45.4 39.9c-15.8 28.5-27.1 89.4.6 127.3c32.4 44.6 87.7-2.8 87.7-2.8c102.7-91.9-10.5-225-10.5-225c-6.1-5.5.8-2.8.8-2.8c50.1 36.5 114.6 84.4 116.2 204.8C500.9 400.2 399 504 256.5 504"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedRiver(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M353.2 32H94.8C42.4 32 0 74.4 0 126.8v258.4C0 437.6 42.4 480 94.8 480h258.4c52.4 0 94.8-42.4 94.8-94.8V126.8c0-52.4-42.4-94.8-94.8-94.8M144.9 200.9v56.3c0 27-21.9 48.9-48.9 48.9V151.9c0-13.2 10.7-23.9 23.9-23.9h154.2c0 27-21.9 48.9-48.9 48.9h-56.3c-12.3-.6-24.6 11.6-24 24m176.3 72h-56.3c-12.3-.6-24.6 11.6-24 24v56.3c0 27-21.9 48.9-48.9 48.9V247.9c0-13.2 10.7-23.9 23.9-23.9h154.2c0 27-21.9 48.9-48.9 48.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reddit(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M201.5 305.5c-13.8 0-24.9-11.1-24.9-24.6c0-13.8 11.1-24.9 24.9-24.9c13.6 0 24.6 11.1 24.6 24.9c0 13.6-11.1 24.6-24.6 24.6M504 256c0 137-111 248-248 248S8 393 8 256S119 8 256 8s248 111 248 248m-132.3-41.2c-9.4 0-17.7 3.9-23.8 10c-22.4-15.5-52.6-25.5-86.1-26.6l17.4-78.3l55.4 12.5c0 13.6 11.1 24.6 24.6 24.6c13.8 0 24.9-11.3 24.9-24.9s-11.1-24.9-24.9-24.9c-9.7 0-18 5.8-22.1 13.8l-61.2-13.6c-3-.8-6.1 1.4-6.9 4.4l-19.1 86.4c-33.2 1.4-63.1 11.3-85.5 26.8c-6.1-6.4-14.7-10.2-24.1-10.2c-34.9 0-46.3 46.9-14.4 62.8c-1.1 5-1.7 10.2-1.7 15.5c0 52.6 59.2 95.2 132 95.2c73.1 0 132.3-42.6 132.3-95.2c0-5.3-.6-10.8-1.9-15.8c31.3-16 19.8-62.5-14.9-62.5M302.8 331c-18.2 18.2-76.1 17.9-93.6 0c-2.2-2.2-6.1-2.2-8.3 0c-2.5 2.5-2.5 6.4 0 8.6c22.8 22.8 87.3 22.8 110.2 0c2.5-2.2 2.5-6.1 0-8.6c-2.2-2.2-6.1-2.2-8.3 0m7.7-75c-13.6 0-24.6 11.1-24.6 24.9c0 13.6 11.1 24.6 24.6 24.6c13.8 0 24.9-11.1 24.9-24.6c0-13.8-11-24.9-24.9-24.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedditAlien(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M440.3 203.5c-15 0-28.2 6.2-37.9 15.9c-35.7-24.7-83.8-40.6-137.1-42.3L293 52.3l88.2 19.8c0 21.6 17.6 39.2 39.2 39.2c22 0 39.7-18.1 39.7-39.7s-17.6-39.7-39.7-39.7c-15.4 0-28.7 9.3-35.3 22l-97.4-21.6c-4.9-1.3-9.7 2.2-11 7.1L246.3 177c-52.9 2.2-100.5 18.1-136.3 42.8c-9.7-10.1-23.4-16.3-38.4-16.3c-55.6 0-73.8 74.6-22.9 100.1c-1.8 7.9-2.6 16.3-2.6 24.7c0 83.8 94.4 151.7 210.3 151.7c116.4 0 210.8-67.9 210.8-151.7c0-8.4-.9-17.2-3.1-25.1c49.9-25.6 31.5-99.7-23.8-99.7M129.4 308.9c0-22 17.6-39.7 39.7-39.7c21.6 0 39.2 17.6 39.2 39.7c0 21.6-17.6 39.2-39.2 39.2c-22 .1-39.7-17.6-39.7-39.2m214.3 93.5c-36.4 36.4-139.1 36.4-175.5 0c-4-3.5-4-9.7 0-13.7c3.5-3.5 9.7-3.5 13.2 0c27.8 28.5 120 29 149 0c3.5-3.5 9.7-3.5 13.2 0c4.1 4 4.1 10.2.1 13.7m-.8-54.2c-21.6 0-39.2-17.6-39.2-39.2c0-22 17.6-39.7 39.2-39.7c22 0 39.7 17.6 39.7 39.7c-.1 21.5-17.7 39.2-39.7 39.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedditSquare(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M283.2 345.5c2.7 2.7 2.7 6.8 0 9.2c-24.5 24.5-93.8 24.6-118.4 0c-2.7-2.4-2.7-6.5 0-9.2c2.4-2.4 6.5-2.4 8.9 0c18.7 19.2 81 19.6 100.5 0c2.4-2.3 6.6-2.3 9 0m-91.3-53.8c0-14.9-11.9-26.8-26.5-26.8c-14.9 0-26.8 11.9-26.8 26.8c0 14.6 11.9 26.5 26.8 26.5c14.6 0 26.5-11.9 26.5-26.5m90.7-26.8c-14.6 0-26.5 11.9-26.5 26.8c0 14.6 11.9 26.5 26.5 26.5c14.9 0 26.8-11.9 26.8-26.5c0-14.9-11.9-26.8-26.8-26.8M448 80v352c0 26.5-21.5 48-48 48H48c-26.5 0-48-21.5-48-48V80c0-26.5 21.5-48 48-48h352c26.5 0 48 21.5 48 48m-99.7 140.6c-10.1 0-19 4.2-25.6 10.7c-24.1-16.7-56.5-27.4-92.5-28.6l18.7-84.2l59.5 13.4c0 14.6 11.9 26.5 26.5 26.5c14.9 0 26.8-12.2 26.8-26.8c0-14.6-11.9-26.8-26.8-26.8c-10.4 0-19.3 6.2-23.8 14.9l-65.7-14.6c-3.3-.9-6.5 1.5-7.4 4.8l-20.5 92.8c-35.7 1.5-67.8 12.2-91.9 28.9c-6.5-6.8-15.8-11-25.9-11c-37.5 0-49.8 50.4-15.5 67.5c-1.2 5.4-1.8 11-1.8 16.7c0 56.5 63.7 102.3 141.9 102.3c78.5 0 142.2-45.8 142.2-102.3c0-5.7-.6-11.6-2.1-17c33.6-17.2 21.2-67.2-16.1-67.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Redhat(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341.52 285.56c33.65 0 82.34-6.94 82.34-47c.22-6.74.86-1.82-20.88-96.24c-4.62-19.15-8.68-27.84-42.31-44.65c-26.09-13.34-82.92-35.37-99.73-35.37c-15.66 0-20.2 20.17-38.87 20.17c-18 0-31.31-15.06-48.12-15.06c-16.14 0-26.66 11-34.78 33.62c-27.5 77.55-26.28 74.27-26.12 78.27c0 24.8 97.64 106.11 228.47 106.11M429 254.84c4.65 22 4.65 24.35 4.65 27.25c0 37.66-42.33 58.56-98 58.56c-125.74.08-235.91-73.65-235.91-122.33a49.55 49.55 0 0 1 4.06-19.72C58.56 200.86 0 208.93 0 260.63c0 84.67 200.63 189 359.49 189c121.79 0 152.51-55.08 152.51-98.58c0-34.21-29.59-73.05-82.93-96.24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rendact(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M248 8C111 8 0 119 0 256s111 248 248 248c18.6 0 36.7-2.1 54.1-5.9-5.6-7.4-10.8-14.4-15.9-21.3-12.4 2.1-25.2 3.3-38.3 3.3C124.3 480 24 379.7 24 256S124.3 32 248 32s224 100.3 224 224c0 71-33 134.2-84.5 175.3-25.9 18.8-39.1 21.4-83.5-44.2-78.7-112.9-48-71.1-73.7-108.3 72.8 8.9 228.5-72 168.6-168.6C314-26.8 15 93.8 59.7 226.4c3.2 9.8 14.4 38.6 45.6 38.6 2 0 2.6-.6 2-1.7-4.4-8.7-20.1-9.8-20.1-37.4 0-40.5 40.5-89.6 100.3-120 66.1-32.3 131.9-30.2 158.2 5.4 27.2 38.3-20.9 119.2-120.4 136.9 7.5-9.4 57-75.2 62.8-84 22.7-34.6 23.6-49 14-59.2-15.5-16.9-29.5-10.3-50.7-11.7-10.8-.9-113.7 181.2-136.4 216.9-5.9 9-21.2 34.1-21.2 50.9 0 21.3 2.8 51.4 20.6 51.4 10.6 0 8-18.7 8-26.6 0-12.9 27.4-49.4 74.8-104.6 20.4 36.1 57.7 114.3 130.2 209.7 98-33.1 168.5-125.8 168.5-235C496 119 385 8 248 8z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Renren(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M214 169.1c0 110.4-61 205.4-147.6 247.4C30 373.2 8 317.7 8 256.6C8 133.9 97.1 32.2 214 12.5zM255 504c-42.9 0-83.3-11-118.5-30.4C193.7 437.5 239.9 382.9 255 319c15.5 63.9 61.7 118.5 118.8 154.7C338.7 493 298.3 504 255 504m190.6-87.5C359 374.5 298 279.6 298 169.1V12.5c116.9 19.7 206 121.4 206 244.1c0 61.1-22 116.6-58.4 159.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Replyd(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M320 480H128C57.6 480 0 422.4 0 352V160C0 89.6 57.6 32 128 32h192c70.4 0 128 57.6 128 128v192c0 70.4-57.6 128-128 128M193.4 273.2c-6.1-2-11.6-3.1-16.4-3.1c-7.2 0-13.5 1.9-18.9 5.6c-5.4 3.7-9.6 9-12.8 15.8h-1.1l-4.2-18.3h-28v138.9h36.1v-89.7c1.5-5.4 4.4-9.8 8.7-13.2c4.3-3.4 9.8-5.1 16.2-5.1c4.6 0 9.8 1 15.6 3.1zm115.2 103.4c-3.2 2.4-7.7 4.8-13.7 7.1c-6 2.3-12.8 3.5-20.4 3.5c-12.2 0-21.1-3-26.5-8.9c-5.5-5.9-8.5-14.7-9-26.4h83.3c.9-4.8 1.6-9.4 2.1-13.9c.5-4.4.7-8.6.7-12.5c0-10.7-1.6-19.7-4.7-26.9c-3.2-7.2-7.3-13-12.5-17.2c-5.2-4.3-11.1-7.3-17.8-9.2c-6.7-1.8-13.5-2.8-20.6-2.8c-21.1 0-37.5 6.1-49.2 18.3s-17.5 30.5-17.5 55c0 22.8 5.2 40.7 15.6 53.7c10.4 13.1 26.8 19.6 49.2 19.6c10.7 0 20.9-1.5 30.4-4.6c9.5-3.1 17.1-6.8 22.6-11.2zm-21.8-70.3c3.8 5.4 5.3 13.1 4.6 23.1h-51.7c.9-9.4 3.7-17 8.2-22.6c4.5-5.6 11.5-8.5 21-8.5c8.2-.1 14.1 2.6 17.9 8m79.9 2.5c4.1 3.9 9.4 5.8 16.1 5.8c7 0 12.6-1.9 16.7-5.8s6.1-9.1 6.1-15.6s-2-11.6-6.1-15.4c-4.1-3.8-9.6-5.7-16.7-5.7c-6.7 0-12 1.9-16.1 5.7c-4.1 3.8-6.1 8.9-6.1 15.4s2 11.7 6.1 15.6m0 100.5c4.1 3.9 9.4 5.8 16.1 5.8c7 0 12.6-1.9 16.7-5.8s6.1-9.1 6.1-15.6s-2-11.6-6.1-15.4c-4.1-3.8-9.6-5.7-16.7-5.7c-6.7 0-12 1.9-16.1 5.7c-4.1 3.8-6.1 8.9-6.1 15.4c0 6.6 2 11.7 6.1 15.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Researchgate(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 32v448h448V32zm262.2 334.4c-6.6 3-33.2 6-50-14.2c-9.2-10.6-25.3-33.3-42.2-63.6c-8.9 0-14.7 0-21.4-.6v46.4c0 23.5 6 21.2 25.8 23.9v8.1c-6.9-.3-23.1-.8-35.6-.8c-13.1 0-26.1.6-33.6.8v-8.1c15.5-2.9 22-1.3 22-23.9V225c0-22.6-6.4-21-22-23.9V193c25.8 1 53.1-.6 70.9-.6c31.7 0 55.9 14.4 55.9 45.6c0 21.1-16.7 42.2-39.2 47.5c13.6 24.2 30 45.6 42.2 58.9c7.2 7.8 17.2 14.7 27.2 14.7zm22.9-135c-23.3 0-32.2-15.7-32.2-32.2V167c0-12.2 8.8-30.4 34-30.4s30.4 17.9 30.4 17.9l-10.7 7.2s-5.5-12.5-19.7-12.5c-7.9 0-19.7 7.3-19.7 19.7v26.8c0 13.4 6.6 23.3 17.9 23.3c14.1 0 21.5-10.9 21.5-26.8h-17.9v-10.7h30.4c0 20.5 4.7 49.9-34 49.9m-116.5 44.7c-9.4 0-13.6-.3-20-.8v-69.7c6.4-.6 15-.6 22.5-.6c23.3 0 37.2 12.2 37.2 34.5c0 21.9-15 36.6-39.7 36.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Resolving(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M281.2 278.2c46-13.3 49.6-23.5 44-43.4L314 195.5c-6.1-20.9-18.4-28.1-71.1-12.8L54.7 236.8l28.6 98.6zM248.5 8C131.4 8 33.2 88.7 7.2 197.5l221.9-63.9c34.8-10.2 54.2-11.7 79.3-8.2c36.3 6.1 52.7 25 61.4 55.2l10.7 37.8c8.2 28.1 1 50.6-23.5 73.6c-19.4 17.4-31.2 24.5-61.4 33.2L203 351.8l220.4 27.1l9.7 34.2l-48.1 13.3l-286.8-37.3l23 80.2c36.8 22 80.3 34.7 126.3 34.7c137 0 248.5-111.4 248.5-248.3C497 119.4 385.5 8 248.5 8M38.3 388.6L0 256.8c0 48.5 14.3 93.4 38.3 131.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rev(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M289.67 274.89a65.57 65.57 0 1 1-65.56-65.56a65.64 65.64 0 0 1 65.56 65.56m139.55-5.05h-.13a204.69 204.69 0 0 0-74.32-153l-45.38 26.2a157.07 157.07 0 0 1 71.81 131.84C381.2 361.5 310.73 432 224.11 432S67 361.5 67 274.88c0-81.88 63-149.27 143-156.43v39.12l108.77-62.79L210 32v38.32c-106.7 7.25-191 96-191 204.57c0 111.59 89.12 202.29 200.06 205v.11h210.16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rocketchat(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M284.046 224.8a34.114 34.114 0 1 0 34.317 34.113a34.217 34.217 0 0 0-34.317-34.113m-110.45 0a34.114 34.114 0 1 0 34.317 34.113A34.217 34.217 0 0 0 173.6 224.8Zm220.923 0a34.114 34.114 0 1 0 34.317 34.113a34.215 34.215 0 0 0-34.317-34.113m153.807-55.319c-15.535-24.172-37.31-45.57-64.681-63.618c-52.886-34.817-122.374-54-195.666-54a405.975 405.975 0 0 0-72.032 6.357a238.524 238.524 0 0 0-49.51-36.588C99.684-11.7 40.859.711 11.135 11.421A14.291 14.291 0 0 0 5.58 34.782C26.542 56.458 61.222 99.3 52.7 138.252c-33.142 33.9-51.112 74.776-51.112 117.337c0 43.372 17.97 84.248 51.112 118.148c8.526 38.956-26.154 81.816-47.116 103.491a14.284 14.284 0 0 0 5.555 23.34c29.724 10.709 88.549 23.147 155.324-10.2a238.679 238.679 0 0 0 49.51-36.589A405.972 405.972 0 0 0 288 460.14c73.313 0 142.8-19.159 195.667-53.975c27.371-18.049 49.145-39.426 64.679-63.619c17.309-26.923 26.07-55.916 26.07-86.125c-.022-31.021-8.782-59.991-26.09-86.936ZM284.987 409.9a345.65 345.65 0 0 1-89.446-11.5l-20.129 19.393a184.366 184.366 0 0 1-37.138 27.585a145.767 145.767 0 0 1-52.522 14.87c.983-1.771 1.881-3.563 2.842-5.356q30.258-55.68 16.325-100.078c-32.992-25.962-52.778-59.2-52.778-95.4c0-83.1 104.254-150.469 232.846-150.469s232.867 67.373 232.867 150.469c0 83.111-104.254 150.486-232.867 150.486"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rockrms(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M248 8C111 8 0 119 0 256s111 248 248 248s248-111 248-248S385 8 248 8m157.4 419.5h-90l-112-131.3c-17.9-20.4-3.9-56.1 26.6-56.1h75.3l-84.6-99.3l-84.3 98.9h-90L193.5 67.2c14.4-18.4 41.3-17.3 54.5 0l157.7 185.1c19 22.8 2 57.2-27.6 56.1c-.6 0-74.2.2-74.2.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rust(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m508.52 249.75l-21.82-13.51c-.17-2-.34-3.93-.55-5.88l18.72-17.5a7.35 7.35 0 0 0-2.44-12.25l-24-9c-.54-1.88-1.08-3.78-1.67-5.64l15-20.83a7.35 7.35 0 0 0-4.79-11.54l-25.42-4.15c-.9-1.73-1.79-3.45-2.73-5.15l10.68-23.42a7.35 7.35 0 0 0-6.95-10.39l-25.82.91q-1.79-2.22-3.61-4.4L439 81.84a7.36 7.36 0 0 0-8.84-8.84L405 78.93q-2.17-1.83-4.4-3.61l.91-25.82a7.35 7.35 0 0 0-10.39-7L367.7 53.23c-1.7-.94-3.43-1.84-5.15-2.73l-4.15-25.42a7.35 7.35 0 0 0-11.54-4.79L326 35.26c-1.86-.59-3.75-1.13-5.64-1.67l-9-24a7.35 7.35 0 0 0-12.25-2.44l-17.5 18.72c-1.95-.21-3.91-.38-5.88-.55L262.25 3.48a7.35 7.35 0 0 0-12.5 0L236.24 25.3c-2 .17-3.93.34-5.88.55l-17.5-18.72a7.35 7.35 0 0 0-12.25 2.44l-9 24c-1.89.55-3.79 1.08-5.66 1.68l-20.82-15a7.35 7.35 0 0 0-11.54 4.79l-4.15 25.41c-1.73.9-3.45 1.79-5.16 2.73l-23.4-10.63a7.35 7.35 0 0 0-10.39 7l.92 25.81c-1.49 1.19-3 2.39-4.42 3.61L81.84 73A7.36 7.36 0 0 0 73 81.84L78.93 107c-1.23 1.45-2.43 2.93-3.62 4.41l-25.81-.91a7.42 7.42 0 0 0-6.37 3.26a7.35 7.35 0 0 0-.57 7.13l10.66 23.41c-.94 1.7-1.83 3.43-2.73 5.16l-25.41 4.14a7.35 7.35 0 0 0-4.79 11.54l15 20.82c-.59 1.87-1.13 3.77-1.68 5.66l-24 9a7.35 7.35 0 0 0-2.44 12.25l18.72 17.5c-.21 1.95-.38 3.91-.55 5.88l-21.86 13.5a7.35 7.35 0 0 0 0 12.5l21.82 13.51c.17 2 .34 3.92.55 5.87l-18.72 17.5a7.35 7.35 0 0 0 2.44 12.25l24 9c.55 1.89 1.08 3.78 1.68 5.65l-15 20.83a7.35 7.35 0 0 0 4.79 11.54l25.42 4.15c.9 1.72 1.79 3.45 2.73 5.14l-10.63 23.43a7.35 7.35 0 0 0 .57 7.13a7.13 7.13 0 0 0 6.37 3.26l25.83-.91q1.77 2.22 3.6 4.4L73 430.16a7.36 7.36 0 0 0 8.84 8.84l25.16-5.93q2.18 1.83 4.41 3.61l-.92 25.82a7.35 7.35 0 0 0 10.39 6.95l23.43-10.68c1.69.94 3.42 1.83 5.14 2.73l4.15 25.42a7.34 7.34 0 0 0 11.54 4.78l20.83-15c1.86.6 3.76 1.13 5.65 1.68l9 24a7.36 7.36 0 0 0 12.25 2.44l17.5-18.72c1.95.21 3.92.38 5.88.55l13.51 21.82a7.35 7.35 0 0 0 12.5 0l13.51-21.82c2-.17 3.93-.34 5.88-.56l17.5 18.73a7.36 7.36 0 0 0 12.25-2.44l9-24c1.89-.55 3.78-1.08 5.65-1.68l20.82 15a7.34 7.34 0 0 0 11.54-4.78l4.15-25.42c1.72-.9 3.45-1.79 5.15-2.73l23.42 10.68a7.35 7.35 0 0 0 10.39-6.95l-.91-25.82q2.22-1.79 4.4-3.61l25.15 5.93a7.36 7.36 0 0 0 8.84-8.84L433.07 405q1.83-2.17 3.61-4.4l25.82.91a7.23 7.23 0 0 0 6.37-3.26a7.35 7.35 0 0 0 .58-7.13l-10.68-23.42c.94-1.7 1.83-3.43 2.73-5.15l25.42-4.15a7.35 7.35 0 0 0 4.79-11.54l-15-20.83c.59-1.87 1.13-3.76 1.67-5.65l24-9a7.35 7.35 0 0 0 2.44-12.25l-18.72-17.5c.21-1.95.38-3.91.55-5.87l21.82-13.51a7.35 7.35 0 0 0 0-12.5Zm-151 129.08A13.91 13.91 0 0 0 341 389.51l-7.64 35.67a187.51 187.51 0 0 1-156.36-.74l-7.64-35.66a13.87 13.87 0 0 0-16.46-10.68l-31.51 6.76a187.38 187.38 0 0 1-16.26-19.21H258.3c1.72 0 2.89-.29 2.89-1.91v-54.19c0-1.57-1.17-1.91-2.89-1.91h-44.83l.05-34.35H262c4.41 0 23.66 1.28 29.79 25.87c1.91 7.55 6.17 32.14 9.06 40c2.89 8.82 14.6 26.46 27.1 26.46H407a187.3 187.3 0 0 1-17.34 20.09Zm25.77 34.49A15.24 15.24 0 1 1 368 398.08h.44a15.23 15.23 0 0 1 14.8 15.24Zm-225.62-.68a15.24 15.24 0 1 1-15.25-15.25h.45a15.25 15.25 0 0 1 14.75 15.25Zm-88.1-178.49l32.83-14.6a13.88 13.88 0 0 0 7.06-18.33L102.69 186h26.56v119.73h-53.6a187.65 187.65 0 0 1-6.08-71.58m-11.26-36.06a15.24 15.24 0 0 1 15.23-15.25H74a15.24 15.24 0 1 1-15.67 15.24Zm155.16 24.49l.05-35.32h63.26c3.28 0 23.07 3.77 23.07 18.62c0 12.29-15.19 16.7-27.68 16.7ZM399 306.71c-9.8 1.13-20.63-4.12-22-10.09c-5.78-32.49-15.39-39.4-30.57-51.4c18.86-11.95 38.46-29.64 38.46-53.26c0-25.52-17.49-41.59-29.4-49.48c-16.76-11-35.28-13.23-40.27-13.23h-198.9a187.49 187.49 0 0 1 104.89-59.19l23.47 24.6a13.82 13.82 0 0 0 19.6.44l26.26-25a187.51 187.51 0 0 1 128.37 91.43l-18 40.57a14 14 0 0 0 7.09 18.33l34.59 15.33a187.12 187.12 0 0 1 .4 32.54h-19.28c-1.91 0-2.69 1.27-2.69 3.13v8.82C421 301 409.31 305.58 399 306.71M240 60.21A15.24 15.24 0 0 1 255.21 45h.45A15.24 15.24 0 1 1 240 60.21M436.84 214a15.24 15.24 0 1 1 0-30.48h.44a15.24 15.24 0 0 1-.44 30.48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Safari(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m274.69 274.69l-37.38-37.38L166 346ZM256 8C119 8 8 119 8 256s111 248 248 248s248-111 248-248S393 8 256 8m155.85 174.79l14.78-6.13a8 8 0 0 1 10.45 4.34a8 8 0 0 1-4.33 10.46L418 197.57a8 8 0 0 1-10.45-4.33a8 8 0 0 1 4.3-10.45M314.43 94l6.12-14.78a8 8 0 0 1 10.45-4.3a8 8 0 0 1 4.33 10.45l-6.13 14.78a8 8 0 0 1-10.45 4.33A8 8 0 0 1 314.43 94M256 60a8 8 0 0 1 8 8v16a8 8 0 0 1-8 8a8 8 0 0 1-8-8V68a8 8 0 0 1 8-8m-75 14.92a8 8 0 0 1 10.46 4.33L197.57 94a8 8 0 1 1-14.78 6.12l-6.13-14.78A8 8 0 0 1 181 74.92m-63.58 42.49a8 8 0 0 1 11.31 0L140 128.72a8 8 0 0 1 0 11.28a8 8 0 0 1-11.31 0l-11.31-11.31a8 8 0 0 1 .03-11.28ZM60 256a8 8 0 0 1 8-8h16a8 8 0 0 1 8 8a8 8 0 0 1-8 8H68a8 8 0 0 1-8-8m40.15 73.21l-14.78 6.13A8 8 0 0 1 74.92 331a8 8 0 0 1 4.33-10.46L94 314.43a8 8 0 0 1 10.45 4.33a8 8 0 0 1-4.3 10.45m4.33-136A8 8 0 0 1 94 197.57l-14.78-6.12a8 8 0 0 1-4.3-10.45a8 8 0 0 1 10.45-4.33l14.78 6.13a8 8 0 0 1 4.33 10.44ZM197.57 418l-6.12 14.78a8 8 0 0 1-14.79-6.12l6.13-14.78a8 8 0 1 1 14.78 6.12M264 444a8 8 0 0 1-8 8a8 8 0 0 1-8-8v-16a8 8 0 0 1 8-8a8 8 0 0 1 8 8Zm67-6.92a8 8 0 0 1-10.46-4.33L314.43 418a8 8 0 0 1 4.33-10.45a8 8 0 0 1 10.45 4.33l6.13 14.78a8 8 0 0 1-4.34 10.42m63.58-42.49a8 8 0 0 1-11.31 0L372 383.28a8 8 0 0 1 0-11.28a8 8 0 0 1 11.31 0l11.31 11.31a8 8 0 0 1-.03 11.28ZM286.25 286.25L110.34 401.66l115.41-175.91l175.91-115.41ZM437.08 331a8 8 0 0 1-10.45 4.33l-14.78-6.13a8 8 0 0 1-4.33-10.45a8 8 0 0 1 10.48-4.32l14.78 6.12a8 8 0 0 1 4.3 10.45m6.92-67h-16a8 8 0 0 1-8-8a8 8 0 0 1 8-8h16a8 8 0 0 1 8 8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Salesforce(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M248.89 245.64h-26.35c.69-5.16 3.32-14.12 13.64-14.12c6.75 0 11.97 3.82 12.71 14.12m136.66-13.88c-.47 0-14.11-1.77-14.11 20s13.63 20 14.11 20c13 0 14.11-13.54 14.11-20c0-21.76-13.66-20-14.11-20m-243.22 23.76a8.63 8.63 0 0 0-3.29 7.29c0 4.78 2.08 6.05 3.29 7.05c4.7 3.7 15.07 2.12 20.93.95v-16.94c-5.32-1.07-16.73-1.96-20.93 1.65M640 232c0 87.58-80 154.39-165.36 136.43c-18.37 33-70.73 70.75-132.2 41.63c-41.16 96.05-177.89 92.18-213.81-5.17C8.91 428.78-50.19 266.52 53.36 205.61C18.61 126.18 76 32 167.67 32a124.24 124.24 0 0 1 98.56 48.7c20.7-21.4 49.4-34.81 81.15-34.81c42.34 0 79 23.52 98.8 58.57C539 63.78 640 132.69 640 232m-519.55 31.8c0-11.76-11.69-15.17-17.87-17.17c-5.27-2.11-13.41-3.51-13.41-8.94c0-9.46 17-6.66 25.17-2.12c0 0 1.17.71 1.64-.47c.24-.7 2.36-6.58 2.59-7.29a1.13 1.13 0 0 0-.7-1.41c-12.33-7.63-40.7-8.51-40.7 12.7c0 12.46 11.49 15.44 17.88 17.17c4.72 1.58 13.17 3 13.17 8.7c0 4-3.53 7.06-9.17 7.06a31.76 31.76 0 0 1-19-6.35c-.47-.23-1.42-.71-1.65.71l-2.4 7.47c-.47.94.23 1.18.23 1.41c1.75 1.4 10.3 6.59 22.82 6.59c13.17 0 21.4-7.06 21.4-18.11zm32-42.58c-10.13 0-18.66 3.17-21.4 5.18a1 1 0 0 0-.24 1.41l2.59 7.06a1 1 0 0 0 1.18.7c.65 0 6.8-4 16.93-4c4 0 7.06.71 9.18 2.36c3.6 2.8 3.06 8.29 3.06 10.58c-4.79-.3-19.11-3.44-29.41 3.76a16.92 16.92 0 0 0-7.34 14.54c0 5.9 1.51 10.4 6.59 14.35c12.24 8.16 36.28 2 38.1 1.41c1.58-.32 3.53-.66 3.53-1.88v-33.88c.04-4.61.32-21.64-22.78-21.64zM199 200.24a1.11 1.11 0 0 0-1.18-1.18H188a1.11 1.11 0 0 0-1.17 1.18v79a1.11 1.11 0 0 0 1.17 1.18h9.88a1.11 1.11 0 0 0 1.18-1.18zm55.75 28.93c-2.1-2.31-6.79-7.53-17.65-7.53c-3.51 0-14.16.23-20.7 8.94c-6.35 7.63-6.58 18.11-6.58 21.41c0 3.12.15 14.26 7.06 21.17c2.64 2.91 9.06 8.23 22.81 8.23c10.82 0 16.47-2.35 18.58-3.76c.47-.24.71-.71.24-1.88l-2.35-6.83a1.26 1.26 0 0 0-1.41-.7c-2.59.94-6.35 2.82-15.29 2.82c-17.42 0-16.85-14.74-16.94-16.7h37.17a1.23 1.23 0 0 0 1.17-.94c-.29 0 2.07-14.7-6.09-24.23zm36.69 52.69c13.17 0 21.41-7.06 21.41-18.11c0-11.76-11.7-15.17-17.88-17.17c-4.14-1.66-13.41-3.38-13.41-8.94c0-3.76 3.29-6.35 8.47-6.35a38.11 38.11 0 0 1 16.7 4.23s1.18.71 1.65-.47c.23-.7 2.35-6.58 2.58-7.29a1.13 1.13 0 0 0-.7-1.41c-7.91-4.9-16.74-4.94-20.23-4.94c-12 0-20.46 7.29-20.46 17.64c0 12.46 11.48 15.44 17.87 17.17c6.11 2 13.17 3.26 13.17 8.7c0 4-3.52 7.06-9.17 7.06a31.8 31.8 0 0 1-19-6.35a1 1 0 0 0-1.65.71l-2.35 7.52c-.47.94.23 1.18.23 1.41c1.72 1.4 10.33 6.59 22.79 6.59zM357.09 224c0-.71-.24-1.18-1.18-1.18h-11.76c0-.14.94-8.94 4.47-12.47c4.16-4.15 11.76-1.64 12-1.64c1.17.47 1.41 0 1.64-.47l2.83-7.77c.7-.94 0-1.17-.24-1.41c-5.09-2-17.35-2.87-24.46 4.24c-5.48 5.48-7 13.92-8 19.52h-8.47a1.28 1.28 0 0 0-1.17 1.18l-1.42 7.76c0 .7.24 1.17 1.18 1.17h8.23c-8.51 47.9-8.75 50.21-10.35 55.52c-1.08 3.62-3.29 6.9-5.88 7.76c-.09 0-3.88 1.68-9.64-.24c0 0-.94-.47-1.41.71c-.24.71-2.59 6.82-2.83 7.53s0 1.41.47 1.41c5.11 2 13 1.77 17.88 0c6.28-2.28 9.72-7.89 11.53-12.94c2.75-7.71 2.81-9.79 11.76-59.74h12.23a1.29 1.29 0 0 0 1.18-1.18zm53.39 16c-.56-1.68-5.1-18.11-25.17-18.11c-15.25 0-23 10-25.16 18.11c-1 3-3.18 14 0 23.52c.09.3 4.41 18.12 25.16 18.12c14.95 0 22.9-9.61 25.17-18.12c3.21-9.61 1.01-20.52 0-23.52m45.4-16.7c-5-1.65-16.62-1.9-22.11 5.41v-4.47a1.11 1.11 0 0 0-1.18-1.17h-9.4a1.11 1.11 0 0 0-1.18 1.17v55.28a1.12 1.12 0 0 0 1.18 1.18h9.64a1.12 1.12 0 0 0 1.18-1.18v-27.77c0-2.91.05-11.37 4.46-15.05c4.9-4.9 12-3.36 13.41-3.06a1.57 1.57 0 0 0 1.41-.94a74 74 0 0 0 3.06-8a1.16 1.16 0 0 0-.47-1.41zm46.81 54.1l-2.12-7.29c-.47-1.18-1.41-.71-1.41-.71c-4.23 1.82-10.15 1.89-11.29 1.89c-4.64 0-17.17-1.13-17.17-19.76c0-6.23 1.85-19.76 16.47-19.76a34.85 34.85 0 0 1 11.52 1.65s.94.47 1.18-.71c.94-2.59 1.64-4.47 2.59-7.53c.23-.94-.47-1.17-.71-1.17c-11.59-3.87-22.34-2.53-27.76 0c-1.59.74-16.23 6.49-16.23 27.52c0 2.9-.58 30.11 28.94 30.11a44.45 44.45 0 0 0 15.52-2.83a1.3 1.3 0 0 0 .47-1.42zm53.87-39.52c-.8-3-5.37-16.23-22.35-16.23c-16 0-23.52 10.11-25.64 18.59a38.58 38.58 0 0 0-1.65 11.76c0 25.87 18.84 29.4 29.88 29.4c10.82 0 16.46-2.35 18.58-3.76c.47-.24.71-.71.24-1.88l-2.36-6.83a1.26 1.26 0 0 0-1.41-.7c-2.59.94-6.35 2.82-15.29 2.82c-17.42 0-16.85-14.74-16.93-16.7h37.16a1.25 1.25 0 0 0 1.18-.94c-.24-.01.94-7.07-1.41-15.54zm-23.29-6.35c-10.33 0-13 9-13.64 14.12H546c-.88-11.92-7.62-14.13-12.73-14.13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sass(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M301.84 378.92c-.3.6-.6 1.08 0 0m249.13-87a131.16 131.16 0 0 0-58 13.5c-5.9-11.9-12-22.3-13-30.1c-1.2-9.1-2.5-14.5-1.1-25.3s7.7-26.1 7.6-27.2s-1.4-6.6-14.3-6.7s-24 2.5-25.29 5.9a122.83 122.83 0 0 0-5.3 19.1c-2.3 11.7-25.79 53.5-39.09 75.3c-4.4-8.5-8.1-16-8.9-22c-1.2-9.1-2.5-14.5-1.1-25.3s7.7-26.1 7.6-27.2s-1.4-6.6-14.29-6.7s-24 2.5-25.3 5.9s-2.7 11.4-5.3 19.1s-33.89 77.3-42.08 95.4c-4.2 9.2-7.8 16.6-10.4 21.6c-.4.8-.7 1.3-.9 1.7c.3-.5.5-1 .5-.8c-2.2 4.3-3.5 6.7-3.5 6.7v.1c-1.7 3.2-3.6 6.1-4.5 6.1c-.6 0-1.9-8.4.3-19.9c4.7-24.2 15.8-61.8 15.7-63.1c-.1-.7 2.1-7.2-7.3-10.7c-9.1-3.3-12.4 2.2-13.2 2.2s-1.4 2-1.4 2s10.1-42.4-19.39-42.4c-18.4 0-44 20.2-56.58 38.5c-7.9 4.3-25 13.6-43 23.5c-6.9 3.8-14 7.7-20.7 11.4c-.5-.5-.9-1-1.4-1.5c-35.79-38.2-101.87-65.2-99.07-116.5c1-18.7 7.5-67.8 127.07-127.4c98-48.8 176.35-35.4 189.84-5.6c19.4 42.5-41.89 121.6-143.66 133c-38.79 4.3-59.18-10.7-64.28-16.3c-5.3-5.9-6.1-6.2-8.1-5.1c-3.3 1.8-1.2 7 0 10.1c3 7.9 15.5 21.9 36.79 28.9c18.7 6.1 64.18 9.5 119.17-11.8c61.78-23.8 109.87-90.1 95.77-145.6C386.52 18.32 293-.18 204.57 31.22c-52.69 18.7-109.67 48.1-150.66 86.4c-48.69 45.6-56.48 85.3-53.28 101.9c11.39 58.9 92.57 97.3 125.06 125.7c-1.6.9-3.1 1.7-4.5 2.5c-16.29 8.1-78.18 40.5-93.67 74.7c-17.5 38.8 2.9 66.6 16.29 70.4c41.79 11.6 84.58-9.3 107.57-43.6s20.2-79.1 9.6-99.5c-.1-.3-.3-.5-.4-.8c4.2-2.5 8.5-5 12.8-7.5c8.29-4.9 16.39-9.4 23.49-13.3c-4 10.8-6.9 23.8-8.4 42.6c-1.8 22 7.3 50.5 19.1 61.7c5.2 4.9 11.49 5 15.39 5c13.8 0 20-11.4 26.89-25c8.5-16.6 16-35.9 16-35.9s-9.4 52.2 16.3 52.2c9.39 0 18.79-12.1 23-18.3v.1s.2-.4.7-1.2c1-1.5 1.5-2.4 1.5-2.4v-.3c3.8-6.5 12.1-21.4 24.59-46c16.2-31.8 31.69-71.5 31.69-71.5a201.24 201.24 0 0 0 6.2 25.8c2.8 9.5 8.7 19.9 13.4 30c-3.8 5.2-6.1 8.2-6.1 8.2a.31.31 0 0 0 .1.2c-3 4-6.4 8.3-9.9 12.5c-12.79 15.2-28 32.6-30 37.6c-2.4 5.9-1.8 10.3 2.8 13.7c3.4 2.6 9.4 3 15.69 2.5c11.5-.8 19.6-3.6 23.5-5.4a82.2 82.2 0 0 0 20.19-10.6c12.5-9.2 20.1-22.4 19.4-39.8c-.4-9.6-3.5-19.2-7.3-28.2c1.1-1.6 2.3-3.3 3.4-5C434.8 301.72 450.1 270 450.1 270a201.24 201.24 0 0 0 6.2 25.8c2.4 8.1 7.09 17 11.39 25.7c-18.59 15.1-30.09 32.6-34.09 44.1c-7.4 21.3-1.6 30.9 9.3 33.1c4.9 1 11.9-1.3 17.1-3.5a79.46 79.46 0 0 0 21.59-11.1c12.5-9.2 24.59-22.1 23.79-39.6c-.3-7.9-2.5-15.8-5.4-23.4c15.7-6.6 36.09-10.2 62.09-7.2c55.68 6.5 66.58 41.3 64.48 55.8s-13.8 22.6-17.7 25s-5.1 3.3-4.8 5.1c.5 2.6 2.3 2.5 5.6 1.9c4.6-.8 29.19-11.8 30.29-38.7c1.6-34-31.09-71.4-89-71.1zm-429.18 144.7c-18.39 20.1-44.19 27.7-55.28 21.3C54.61 451 59.31 421.42 82 400c13.8-13 31.59-25 43.39-32.4c2.7-1.6 6.6-4 11.4-6.9c.8-.5 1.2-.7 1.2-.7c.9-.6 1.9-1.1 2.9-1.7c8.29 30.4.3 57.2-19.1 78.3zm134.36-91.4c-6.4 15.7-19.89 55.7-28.09 53.6c-7-1.8-11.3-32.3-1.4-62.3c5-15.1 15.6-33.1 21.9-40.1c10.09-11.3 21.19-14.9 23.79-10.4c3.5 5.9-12.2 49.4-16.2 59.2m111 53c-2.7 1.4-5.2 2.3-6.4 1.6c-.9-.5 1.1-2.4 1.1-2.4s13.9-14.9 19.4-21.7c3.2-4 6.9-8.7 10.89-13.9c0 .5.1 1 .1 1.6c-.13 17.9-17.32 30-25.12 34.8zm85.58-19.5c-2-1.4-1.7-6.1 5-20.7c2.6-5.7 8.59-15.3 19-24.5a36.18 36.18 0 0 1 1.9 10.8c-.1 22.5-16.2 30.9-25.89 34.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Schlix(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m350.5 157.7l-54.2-46.1l73.4-39l78.3 44.2zM192 122.1l45.7-28.2l34.7 34.6l-55.4 29zm-65.1 6.6l31.9-22.1L176 135l-36.7 22.5zm-23.3 88.2l-8.8-34.8l29.6-18.3l13.1 35.3zm-21.2-83.7l23.9-18.1l8.9 24l-26.7 18.3zM59 206.5l-3.6-28.4l22.3-15.5l6.1 28.7zm-30.6 16.6l20.8-12.8l3.3 33.4l-22.9 12zM1.4 268l19.2-10.2l.4 38.2l-21 8.8zm59.1 59.3l-28.3 8.3l-1.6-46.8l25.1-10.7zM99 263.2l-31.1 13l-5.2-40.8L90.1 221zM123.2 377l-41.6 5.9l-8.1-63.5l35.2-10.8zm28.5-139.9l21.2 57.1l-46.2 13.6l-13.7-54.1zm85.7 230.5l-70.9-3.3l-24.3-95.8l55.2-8.6zm-84.9-279.7l42.2-22.4l28 45.9l-50.8 21.3zm41 94.9l61.3-18.7l52.8 86.6l-79.8 11.3zm51.4-85.6l67.3-28.8l65.5 65.4l-88.6 26.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scribd(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M42.3 252.7c-16.1-19-24.7-45.9-24.8-79.9c0-100.4 75.2-153.1 167.2-153.1c98.6-1.6 156.8 49 184.3 70.6l-50.5 72.1l-37.3-24.6l26.9-38.6c-36.5-24-79.4-36.5-123-35.8c-50.7-.8-111.7 27.2-111.7 76.2c0 18.7 11.2 20.7 28.6 15.6c23.3-5.3 41.9.6 55.8 14c26.4 24.3 23.2 67.6-.7 91.9c-29.2 29.5-85.2 27.3-114.8-8.4m317.7 5.9c-15.5-18.8-38.9-29.4-63.2-28.6c-38.1-2-71.1 28-70.5 67.2c-.7 16.8 6 33 18.4 44.3c14.1 13.9 33 19.7 56.3 14.4c17.4-5.1 28.6-3.1 28.6 15.6c0 4.3-.5 8.5-1.4 12.7c-16.7 40.9-59.5 64.4-121.4 64.4c-51.9.2-102.4-16.4-144.1-47.3l33.7-39.4l-35.6-27.4L0 406.3l15.4 13.8c52.5 46.8 120.4 72.5 190.7 72.2c51.4 0 94.4-10.5 133.6-44.1c57.1-51.4 54.2-149.2 20.3-189.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Searchengin(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m220.6 130.3l-67.2 28.2V43.2L98.7 233.5l54.7-24.2v130.3zm-83.2-96.7l-1.3 4.7l-15.2 52.9C80.6 106.7 52 145.8 52 191.5c0 52.3 34.3 95.9 83.4 105.5v53.6C57.5 340.1 0 272.4 0 191.6c0-80.5 59.8-147.2 137.4-158m311.4 447.2c-11.2 11.2-23.1 12.3-28.6 10.5c-5.4-1.8-27.1-19.9-60.4-44.4c-33.3-24.6-33.6-35.7-43-56.7c-9.4-20.9-30.4-42.6-57.5-52.4l-9.7-14.7c-24.7 16.9-53 26.9-81.3 28.7l2.1-6.6l15.9-49.5c46.5-11.9 80.9-54 80.9-104.2c0-54.5-38.4-102.1-96-107.1V32.3C254.4 37.4 320 106.8 320 191.6c0 33.6-11.2 64.7-29 90.4l14.6 9.6c9.8 27.1 31.5 48 52.4 57.4s32.2 9.7 56.8 43c24.6 33.2 42.7 54.9 44.5 60.3s.7 17.3-10.5 28.5m-9.9-17.9c0-4.4-3.6-8-8-8s-8 3.6-8 8s3.6 8 8 8s8-3.6 8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sellcast(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M353.4 32H94.7C42.6 32 0 74.6 0 126.6v258.7C0 437.4 42.6 480 94.7 480h258.7c52.1 0 94.7-42.6 94.7-94.6V126.6c0-52-42.6-94.6-94.7-94.6m-50 316.4c-27.9 48.2-89.9 64.9-138.2 37.2c-22.9 39.8-54.9 8.6-42.3-13.2l15.7-27.2c5.9-10.3 19.2-13.9 29.5-7.9c18.6 10.8-.1-.1 18.5 10.7c27.6 15.9 63.4 6.3 79.4-21.3c15.9-27.6 6.3-63.4-21.3-79.4c-17.8-10.2-.6-.4-18.6-10.6c-24.6-14.2-3.4-51.9 21.6-37.5c18.6 10.8-.1-.1 18.5 10.7c48.4 28 65.1 90.3 37.2 138.5m21.8-208.8c-17 29.5-16.3 28.8-19 31.5c-6.5 6.5-16.3 8.7-26.5 3.6c-18.6-10.8.1.1-18.5-10.7c-27.6-15.9-63.4-6.3-79.4 21.3s-6.3 63.4 21.3 79.4c0 0 18.5 10.6 18.6 10.6c24.6 14.2 3.4 51.9-21.6 37.5c-18.6-10.8.1.1-18.5-10.7c-48.2-27.8-64.9-90.1-37.1-138.4c27.9-48.2 89.9-64.9 138.2-37.2l4.8-8.4c14.3-24.9 52-3.3 37.7 21.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sellsy(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M539.71 237.308c3.064-12.257 4.29-24.821 4.29-37.384C544 107.382 468.618 32 376.076 32c-77.22 0-144.634 53.012-163.02 127.781c-15.322-13.176-34.934-20.53-55.157-20.53c-46.271 0-83.962 37.69-83.962 83.961c0 7.354.92 15.015 3.065 22.369c-42.9 20.225-70.785 63.738-70.785 111.234C6.216 424.843 61.68 480 129.401 480h381.198c67.72 0 123.184-55.157 123.184-123.184c.001-56.384-38.916-106.025-94.073-119.508M199.88 401.554c0 8.274-7.048 15.321-15.321 15.321H153.61c-8.274 0-15.321-7.048-15.321-15.321V290.626c0-8.273 7.048-15.321 15.321-15.321h30.949c8.274 0 15.321 7.048 15.321 15.321zm89.477 0c0 8.274-7.048 15.321-15.322 15.321h-30.949c-8.274 0-15.321-7.048-15.321-15.321V270.096c0-8.274 7.048-15.321 15.321-15.321h30.949c8.274 0 15.322 7.048 15.322 15.321zm89.477 0c0 8.274-7.047 15.321-15.321 15.321h-30.949c-8.274 0-15.322-7.048-15.322-15.321V238.84c0-8.274 7.048-15.321 15.322-15.321h30.949c8.274 0 15.321 7.048 15.321 15.321zm87.027 0c0 8.274-7.048 15.321-15.322 15.321h-28.497c-8.274 0-15.321-7.048-15.321-15.321V176.941c0-8.579 7.047-15.628 15.321-15.628h28.497c8.274 0 15.322 7.048 15.322 15.628z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Servicestack(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M88 216c81.7 10.2 273.7 102.3 304 232H0c99.5-8.1 184.5-137 88-232m32-152c32.3 35.6 47.7 83.9 46.4 133.6C249.3 231.3 373.7 321.3 400 448h96C455.3 231.9 222.8 79.5 120 64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shirtsinbulk(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m100 410.3l30.6 13.4l4.4-9.9l-30.6-13.4zm39.4 17.5l30.6 13.4l4.4-9.9l-30.6-13.4zm172.1-14l4.4 9.9l30.6-13.4l-4.4-9.9zM179.1 445l30.3 13.7l4.4-9.9l-30.3-13.4zM60.4 392.8L91 406.2l4.4-9.6l-30.6-13.7zm211.4 38.5l4.4 9.9l30.6-13.4l-4.4-9.9zm-39.3 17.5l4.4 9.9l30.6-13.7l-4.4-9.6zm118.4-52.2l4.4 9.6l30.6-13.4l-4.4-9.9zM170 46.6h-33.5v10.5H170zm-47.2 0H89.2v10.5h33.5zm-47.3 0H42.3v10.5h33.3zm141.5 0h-33.2v10.5H217zm94.5 0H278v10.5h33.5zm47.3 0h-33.5v10.5h33.5zm-94.6 0H231v10.5h33.2zm141.5 0h-33.3v10.5h33.3zM52.8 351.1H42v33.5h10.8zm70-215.9H89.2v10.5h33.5zm-70 10.6h22.8v-10.5H42v33.5h10.8zm168.9 228.6c50.5 0 91.3-40.8 91.3-91.3c0-50.2-40.8-91.3-91.3-91.3c-50.2 0-91.3 41.1-91.3 91.3c0 50.5 41.1 91.3 91.3 91.3m-48.2-111.1c0-25.4 29.5-31.8 49.6-31.8c16.9 0 29.2 5.8 44.3 12l-8.8 16.9h-.9c-6.4-9.9-24.8-13.1-35.6-13.1c-9 0-29.8 1.8-29.8 14.9c0 21.6 78.5-10.2 78.5 37.9c0 25.4-31.5 31.2-51 31.2c-18.1 0-32.4-2.9-47.2-12.2l9-18.4h.9c6.1 12.2 23.6 14.9 35.9 14.9c8.7 0 32.7-1.2 32.7-14.3c0-26.1-77.6 6.3-77.6-38M52.8 178.4H42V212h10.8zm342.4 206.2H406v-33.5h-10.8zM52.8 307.9H42v33.5h10.8zM0 3.7v406l221.7 98.6L448 409.7V3.7zm418.8 387.1L222 476.5L29.2 390.8V120.7h389.7v270.1zm0-299.3H29.2V32.9h389.7v58.6zm-366 130.1H42v33.5h10.8zm0 43.2H42v33.5h10.8zM170 135.2h-33.5v10.5H170zm225.2 163.1H406v-33.5h-10.8zm0-43.2H406v-33.5h-10.8zM217 135.2h-33.2v10.5H217zM395.2 212H406v-33.5h-10.8zm0 129.5H406V308h-10.8zm-131-206.3H231v10.5h33.2zm47.3 0H278v10.5h33.5zm83.7 33.6H406v-33.5h-33.5v10.5h22.8zm-36.4-33.6h-33.5v10.5h33.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shopify(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M388.32 104.1a4.66 4.66 0 0 0-4.4-4c-2 0-37.23-.8-37.23-.8s-21.61-20.82-29.62-28.83V503.2L442.76 472s-54.04-365.5-54.44-367.9m-99.67-33.63a116.67 116.67 0 0 0-7.21-17.61C271 32.85 255.42 22 237 22a15 15 0 0 0-4 .4c-.4-.8-1.2-1.2-1.6-2c-8-8.77-18.4-12.77-30.82-12.4c-24 .8-48 18-67.25 48.83c-13.61 21.62-24 48.84-26.82 70.06c-27.62 8.4-46.83 14.41-47.23 14.81c-14 4.4-14.41 4.8-16 18c-1.2 10-38 291.82-38 291.82L307.86 504V65.67a41.66 41.66 0 0 0-4.4.4s-5.6 1.6-14.81 4.4m-55.24 17.22c-16 4.8-33.63 10.4-50.84 15.61c4.8-18.82 14.41-37.63 25.62-50c4.4-4.4 10.41-9.61 17.21-12.81c6.81 14.37 8.41 33.99 8.01 47.2m-32.83-63.25A27.49 27.49 0 0 1 215 28c-6.4 3.2-12.81 8.41-18.81 14.41c-15.21 16.42-26.82 42-31.62 66.45c-14.42 4.41-28.83 8.81-42 12.81c8.76-38.39 41.18-96.43 78.01-97.23m-46.43 220.17c1.6 25.61 69.25 31.22 73.25 91.66c2.8 47.64-25.22 80.06-65.65 82.47c-48.83 3.2-75.65-25.62-75.65-25.62l10.4-44s26.82 20.42 48.44 18.82c14-.8 19.22-12.41 18.81-20.42c-2-33.62-57.24-31.62-60.84-86.86c-3.2-46.44 27.22-93.27 94.47-97.68c26-1.6 39.23 4.81 39.23 4.81l-15.21 57.6s-17.21-8-37.63-6.4c-29.62 2.01-30.02 20.81-29.62 25.62m95.27-161.73c0-12-1.6-29.22-7.21-43.63c18.42 3.6 27.22 24 31.23 36.43q-10.81 3-24.02 7.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shopware(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M403.5 455.41A246.17 246.17 0 0 1 256 504C118.81 504 8 393 8 256C8 118.81 119 8 256 8a247.39 247.39 0 0 1 165.7 63.5a3.57 3.57 0 0 1-2.86 6.18A418.62 418.62 0 0 0 362.13 74c-129.36 0-222.4 53.47-222.4 155.35c0 109 92.13 145.88 176.83 178.73c33.64 13 65.4 25.36 87 41.59a3.58 3.58 0 0 1 0 5.72zM503 233.09a3.64 3.64 0 0 0-1.27-2.44c-51.76-43-93.62-60.48-144.48-60.48c-84.13 0-80.25 52.17-80.25 53.63c0 42.6 52.06 62 112.34 84.49c31.07 11.59 63.19 23.57 92.68 39.93a3.57 3.57 0 0 0 5-1.82A249 249 0 0 0 503 233.09"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Simplybuilt(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M481.2 64h-106c-14.5 0-26.6 11.8-26.6 26.3v39.6H163.3V90.3c0-14.5-12-26.3-26.6-26.3h-106C16.1 64 4.3 75.8 4.3 90.3v331.4c0 14.5 11.8 26.3 26.6 26.3h450.4c14.8 0 26.6-11.8 26.6-26.3V90.3c-.2-14.5-12-26.3-26.7-26.3M149.8 355.8c-36.6 0-66.4-29.7-66.4-66.4c0-36.9 29.7-66.6 66.4-66.6c36.9 0 66.6 29.7 66.6 66.6c0 36.7-29.7 66.4-66.6 66.4m212.4 0c-36.9 0-66.6-29.7-66.6-66.6c0-36.6 29.7-66.4 66.6-66.4c36.6 0 66.4 29.7 66.4 66.4c0 36.9-29.8 66.6-66.4 66.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sistrix(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 449L301.2 300.2c20-27.9 31.9-62.2 31.9-99.2c0-93.1-74.7-168.9-166.5-168.9C74.7 32 0 107.8 0 200.9s74.7 168.9 166.5 168.9c39.8 0 76.3-14.2 105-37.9l146 148.1zM166.5 330.8c-70.6 0-128.1-58.3-128.1-129.9S95.9 71 166.5 71s128.1 58.3 128.1 129.9s-57.4 129.9-128.1 129.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sith(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 32l69.71 118.75l-58.86-11.52l69.84 91.03a146.741 146.741 0 0 0 0 51.45l-69.84 91.03l58.86-11.52L0 480l118.75-69.71l-11.52 58.86l91.03-69.84c17.02 3.04 34.47 3.04 51.48 0l91.03 69.84l-11.52-58.86L448 480l-69.71-118.78l58.86 11.52l-69.84-91.03c3.03-17.01 3.04-34.44 0-51.45l69.84-91.03l-58.86 11.52L448 32l-118.75 69.71l11.52-58.9l-91.06 69.87c-8.5-1.52-17.1-2.29-25.71-2.29s-17.21.78-25.71 2.29l-91.06-69.87l11.52 58.9zm224 99.78c31.8 0 63.6 12.12 87.85 36.37c48.5 48.5 48.49 127.21 0 175.7s-127.2 48.46-175.7-.03c-48.5-48.5-48.49-127.21 0-175.7c24.24-24.25 56.05-36.34 87.85-36.34m0 36.66c-22.42 0-44.83 8.52-61.92 25.61c-34.18 34.18-34.19 89.68 0 123.87s89.65 34.18 123.84 0c34.18-34.18 34.19-89.68 0-123.87c-17.09-17.09-39.5-25.61-61.92-25.61"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sketch(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27.5 162.2L9 187.1h90.5l6.9-130.7zM396.3 45.7L267.7 32l135.7 147.2zM112.2 218.3l-11.2-22H9.9L234.8 458zm2-31.2h284l-81.5-88.5L256.3 33zm297.3 9.1L277.6 458l224.8-261.7h-90.9zM415.4 69L406 56.4l.9 17.3l6.1 113.4h90.3zM113.5 93.5l-4.6 85.6L244.7 32L116.1 45.7zm287.7 102.7h-290l42.4 82.9L256.3 480z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Skyatlas(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 329.3c0 65.9-52.5 114.4-117.5 114.4c-165.9 0-196.6-249.7-359.7-249.7c-146.9 0-147.1 212.2 5.6 212.2c42.5 0 90.9-17.8 125.3-42.5c5.6-4.1 16.9-16.3 22.8-16.3s10.9 5 10.9 10.9c0 7.8-13.1 19.1-18.7 24.1c-40.9 35.6-100.3 61.2-154.7 61.2c-83.4.1-154-59-154-144.9s67.5-149.1 152.8-149.1c185.3 0 222.5 245.9 361.9 245.9c99.9 0 94.8-139.7 3.4-139.7c-17.5 0-35 11.6-46.9 11.6c-8.4 0-15.9-7.2-15.9-15.6c0-11.6 5.3-23.7 5.3-36.3c0-66.6-50.9-114.7-116.9-114.7c-53.1 0-80 36.9-88.8 36.9c-6.2 0-11.2-5-11.2-11.2c0-5.6 4.1-10.3 7.8-14.4c25.3-28.8 64.7-43.7 102.8-43.7c79.4 0 139.1 58.4 139.1 137.8c0 6.9-.3 13.7-1.2 20.6c11.9-3.1 24.1-4.7 35.9-4.7c60.7 0 111.9 45.3 111.9 107.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Skype(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M424.7 299.8c2.9-14 4.7-28.9 4.7-43.8c0-113.5-91.9-205.3-205.3-205.3c-14.9 0-29.7 1.7-43.8 4.7C161.3 40.7 137.7 32 112 32C50.2 32 0 82.2 0 144c0 25.7 8.7 49.3 23.3 68.2c-2.9 14-4.7 28.9-4.7 43.8c0 113.5 91.9 205.3 205.3 205.3c14.9 0 29.7-1.7 43.8-4.7c19 14.6 42.6 23.3 68.2 23.3c61.8 0 112-50.2 112-112c.1-25.6-8.6-49.2-23.2-68.1m-194.6 91.5c-65.6 0-120.5-29.2-120.5-65c0-16 9-30.6 29.5-30.6c31.2 0 34.1 44.9 88.1 44.9c25.7 0 42.3-11.4 42.3-26.3c0-18.7-16-21.6-42-28c-62.5-15.4-117.8-22-117.8-87.2c0-59.2 58.6-81.1 109.1-81.1c55.1 0 110.8 21.9 110.8 55.4c0 16.9-11.4 31.8-30.3 31.8c-28.3 0-29.2-33.5-75-33.5c-25.7 0-42 7-42 22.5c0 19.8 20.8 21.8 69.1 33c41.4 9.3 90.7 26.8 90.7 77.6c0 59.1-57.1 86.5-112 86.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Slack(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M94.12 315.1c0 25.9-21.16 47.06-47.06 47.06S0 341 0 315.1c0-25.9 21.16-47.06 47.06-47.06h47.06zm23.72 0c0-25.9 21.16-47.06 47.06-47.06s47.06 21.16 47.06 47.06v117.84c0 25.9-21.16 47.06-47.06 47.06s-47.06-21.16-47.06-47.06zm47.06-188.98c-25.9 0-47.06-21.16-47.06-47.06S139 32 164.9 32s47.06 21.16 47.06 47.06v47.06zm0 23.72c25.9 0 47.06 21.16 47.06 47.06s-21.16 47.06-47.06 47.06H47.06C21.16 243.96 0 222.8 0 196.9s21.16-47.06 47.06-47.06zm188.98 47.06c0-25.9 21.16-47.06 47.06-47.06c25.9 0 47.06 21.16 47.06 47.06s-21.16 47.06-47.06 47.06h-47.06zm-23.72 0c0 25.9-21.16 47.06-47.06 47.06c-25.9 0-47.06-21.16-47.06-47.06V79.06c0-25.9 21.16-47.06 47.06-47.06c25.9 0 47.06 21.16 47.06 47.06zM283.1 385.88c25.9 0 47.06 21.16 47.06 47.06c0 25.9-21.16 47.06-47.06 47.06c-25.9 0-47.06-21.16-47.06-47.06v-47.06zm0-23.72c-25.9 0-47.06-21.16-47.06-47.06c0-25.9 21.16-47.06 47.06-47.06h117.84c25.9 0 47.06 21.16 47.06 47.06c0 25.9-21.16 47.06-47.06 47.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlackHash(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M446.2 270.4c-6.2-19-26.9-29.1-46-22.9l-45.4 15.1l-30.3-90l45.4-15.1c19.1-6.2 29.1-26.8 23-45.9c-6.2-19-26.9-29.1-46-22.9l-45.4 15.1l-15.7-47c-6.2-19-26.9-29.1-46-22.9c-19.1 6.2-29.1 26.8-23 45.9l15.7 47l-93.4 31.2l-15.7-47c-6.2-19-26.9-29.1-46-22.9c-19.1 6.2-29.1 26.8-23 45.9l15.7 47l-45.3 15c-19.1 6.2-29.1 26.8-23 45.9c5 14.5 19.1 24 33.6 24.6c6.8 1 12-1.6 57.7-16.8l30.3 90L78 354.8c-19 6.2-29.1 26.9-23 45.9c5 14.5 19.1 24 33.6 24.6c6.8 1 12-1.6 57.7-16.8l15.7 47c5.9 16.9 24.7 29 46 22.9c19.1-6.2 29.1-26.8 23-45.9l-15.7-47l93.6-31.3l15.7 47c5.9 16.9 24.7 29 46 22.9c19.1-6.2 29.1-26.8 23-45.9l-15.7-47l45.4-15.1c19-6 29.1-26.7 22.9-45.7m-254.1 47.2l-30.3-90.2l93.5-31.3l30.3 90.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Slideshare(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M187.7 153.7c-34 0-61.7 25.7-61.7 57.7c0 31.7 27.7 57.7 61.7 57.7s61.7-26 61.7-57.7c0-32-27.7-57.7-61.7-57.7m143.4 0c-34 0-61.7 25.7-61.7 57.7c0 31.7 27.7 57.7 61.7 57.7c34.3 0 61.7-26 61.7-57.7c.1-32-27.4-57.7-61.7-57.7m156.6 90l-6 4.3V49.7c0-27.4-20.6-49.7-46-49.7H76.6c-25.4 0-46 22.3-46 49.7V248c-2-1.4-4.3-2.9-6.3-4.3c-15.1-10.6-25.1 4-16 17.7c18.3 22.6 53.1 50.3 106.3 72C58.3 525.1 252 555.7 248.9 457.5c0-.7.3-56.6.3-96.6c5.1 1.1 9.4 2.3 13.7 3.1c0 39.7.3 92.8.3 93.5c-3.1 98.3 190.6 67.7 134.3-124c53.1-21.7 88-49.4 106.3-72c9.1-13.8-.9-28.3-16.1-17.8m-30.5 19.2c-68.9 37.4-128.3 31.1-160.6 29.7c-23.7-.9-32.6 9.1-33.7 24.9c-10.3-7.7-18.6-15.5-20.3-17.1c-5.1-5.4-13.7-8-27.1-7.7c-31.7 1.1-89.7 7.4-157.4-28V72.3c0-34.9 8.9-45.7 40.6-45.7h317.7c30.3 0 40.9 12.9 40.9 45.7v190.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Snapchat(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M248 8C111 8 0 119 0 256s111 248 248 248s248-111 248-248S385 8 248 8m169.5 338.9c-3.5 8.1-18.1 14-44.8 18.2c-1.4 1.9-2.5 9.8-4.3 15.9c-1.1 3.7-3.7 5.9-8.1 5.9h-.2c-6.2 0-12.8-2.9-25.8-2.9c-17.6 0-23.7 4-37.4 13.7c-14.5 10.3-28.4 19.1-49.2 18.2c-21 1.6-38.6-11.2-48.5-18.2c-13.8-9.7-19.8-13.7-37.4-13.7c-12.5 0-20.4 3.1-25.8 3.1c-5.4 0-7.5-3.3-8.3-6c-1.8-6.1-2.9-14.1-4.3-16c-13.8-2.1-44.8-7.5-45.5-21.4c-.2-3.6 2.3-6.8 5.9-7.4c46.3-7.6 67.1-55.1 68-57.1c0-.1.1-.2.2-.3c2.5-5 3-9.2 1.6-12.5c-3.4-7.9-17.9-10.7-24-13.2c-15.8-6.2-18-13.4-17-18.3c1.6-8.5 14.4-13.8 21.9-10.3c5.9 2.8 11.2 4.2 15.7 4.2c3.3 0 5.5-.8 6.6-1.4c-1.4-23.9-4.7-58 3.8-77.1C183.1 100 230.7 96 244.7 96c.6 0 6.1-.1 6.7-.1c34.7 0 68 17.8 84.3 54.3c8.5 19.1 5.2 53.1 3.8 77.1c1.1.6 2.9 1.3 5.7 1.4c4.3-.2 9.2-1.6 14.7-4.2c4-1.9 9.6-1.6 13.6 0c6.3 2.3 10.3 6.8 10.4 11.9c.1 6.5-5.7 12.1-17.2 16.6c-1.4.6-3.1 1.1-4.9 1.7c-6.5 2.1-16.4 5.2-19 11.5c-1.4 3.3-.8 7.5 1.6 12.5c.1.1.1.2.2.3c.9 2 21.7 49.5 68 57.1c4 1 7.1 5.5 4.9 10.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnapchatGhost(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M510.846 392.673c-5.211 12.157-27.239 21.089-67.36 27.318c-2.064 2.786-3.775 14.686-6.507 23.956c-1.625 5.566-5.623 8.869-12.128 8.869l-.297-.005c-9.395 0-19.203-4.323-38.852-4.323c-26.521 0-35.662 6.043-56.254 20.588c-21.832 15.438-42.771 28.764-74.027 27.399c-31.646 2.334-58.025-16.908-72.871-27.404c-20.714-14.643-29.828-20.582-56.241-20.582c-18.864 0-30.736 4.72-38.852 4.72c-8.073 0-11.213-4.922-12.422-9.04c-2.703-9.189-4.404-21.263-6.523-24.13c-20.679-3.209-67.31-11.344-68.498-32.15a10.627 10.627 0 0 1 8.877-11.069c69.583-11.455 100.924-82.901 102.227-85.934c.074-.176.155-.344.237-.515c3.713-7.537 4.544-13.849 2.463-18.753c-5.05-11.896-26.872-16.164-36.053-19.796c-23.715-9.366-27.015-20.128-25.612-27.504c2.437-12.836 21.725-20.735 33.002-15.453c8.919 4.181 16.843 6.297 23.547 6.297c5.022 0 8.212-1.204 9.96-2.171c-2.043-35.936-7.101-87.29 5.687-115.969C158.122 21.304 229.705 15.42 250.826 15.42c.944 0 9.141-.089 10.11-.089c52.148 0 102.254 26.78 126.723 81.643c12.777 28.65 7.749 79.792 5.695 116.009c1.582.872 4.357 1.942 8.599 2.139c6.397-.286 13.815-2.389 22.069-6.257c6.085-2.846 14.406-2.461 20.48.058l.029.01c9.476 3.385 15.439 10.215 15.589 17.87c.184 9.747-8.522 18.165-25.878 25.018c-2.118.835-4.694 1.655-7.434 2.525c-9.797 3.106-24.6 7.805-28.616 17.271c-2.079 4.904-1.256 11.211 2.46 18.748c.087.168.166.342.239.515c1.301 3.03 32.615 74.46 102.23 85.934c6.427 1.058 11.163 7.877 7.725 15.859"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnapchatSquare(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 32H48C21.5 32 0 53.5 0 80v352c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V80c0-26.5-21.5-48-48-48m-6.5 314.9c-3.5 8.1-18.1 14-44.8 18.2c-1.4 1.9-2.5 9.8-4.3 15.9c-1.1 3.7-3.7 5.9-8.1 5.9h-.2c-6.2 0-12.8-2.9-25.8-2.9c-17.6 0-23.7 4-37.4 13.7c-14.5 10.3-28.4 19.1-49.2 18.2c-21 1.6-38.6-11.2-48.5-18.2c-13.8-9.7-19.8-13.7-37.4-13.7c-12.5 0-20.4 3.1-25.8 3.1c-5.4 0-7.5-3.3-8.3-6c-1.8-6.1-2.9-14.1-4.3-16c-13.8-2.1-44.8-7.5-45.5-21.4c-.2-3.6 2.3-6.8 5.9-7.4c46.3-7.6 67.1-55.1 68-57.1c0-.1.1-.2.2-.3c2.5-5 3-9.2 1.6-12.5c-3.4-7.9-17.9-10.7-24-13.2c-15.8-6.2-18-13.4-17-18.3c1.6-8.5 14.4-13.8 21.9-10.3c5.9 2.8 11.2 4.2 15.7 4.2c3.3 0 5.5-.8 6.6-1.4c-1.4-23.9-4.7-58 3.8-77.1C159.1 100 206.7 96 220.7 96c.6 0 6.1-.1 6.7-.1c34.7 0 68 17.8 84.3 54.3c8.5 19.1 5.2 53.1 3.8 77.1c1.1.6 2.9 1.3 5.7 1.4c4.3-.2 9.2-1.6 14.7-4.2c4-1.9 9.6-1.6 13.6 0c6.3 2.3 10.3 6.8 10.4 11.9c.1 6.5-5.7 12.1-17.2 16.6c-1.4.6-3.1 1.1-4.9 1.7c-6.5 2.1-16.4 5.2-19 11.5c-1.4 3.3-.8 7.5 1.6 12.5c.1.1.1.2.2.3c.9 2 21.7 49.5 68 57.1c4 1 7.1 5.5 4.9 10.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Soundcloud(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m111.4 256.3l5.8 65l-5.8 68.3c-.3 2.5-2.2 4.4-4.4 4.4s-4.2-1.9-4.2-4.4l-5.6-68.3l5.6-65c0-2.2 1.9-4.2 4.2-4.2c2.2 0 4.1 2 4.4 4.2m21.4-45.6c-2.8 0-4.7 2.2-5 5l-5 105.6l5 68.3c.3 2.8 2.2 5 5 5c2.5 0 4.7-2.2 4.7-5l5.8-68.3l-5.8-105.6c0-2.8-2.2-5-4.7-5m25.5-24.1c-3.1 0-5.3 2.2-5.6 5.3l-4.4 130l4.4 67.8c.3 3.1 2.5 5.3 5.6 5.3c2.8 0 5.3-2.2 5.3-5.3l5.3-67.8l-5.3-130c0-3.1-2.5-5.3-5.3-5.3M7.2 283.2c-1.4 0-2.2 1.1-2.5 2.5L0 321.3l4.7 35c.3 1.4 1.1 2.5 2.5 2.5s2.2-1.1 2.5-2.5l5.6-35l-5.6-35.6c-.3-1.4-1.1-2.5-2.5-2.5m23.6-21.9c-1.4 0-2.5 1.1-2.5 2.5l-6.4 57.5l6.4 56.1c0 1.7 1.1 2.8 2.5 2.8s2.5-1.1 2.8-2.5l7.2-56.4l-7.2-57.5c-.3-1.4-1.4-2.5-2.8-2.5m25.3-11.4c-1.7 0-3.1 1.4-3.3 3.3L47 321.3l5.8 65.8c.3 1.7 1.7 3.1 3.3 3.1c1.7 0 3.1-1.4 3.1-3.1l6.9-65.8l-6.9-68.1c0-1.9-1.4-3.3-3.1-3.3m25.3-2.2c-1.9 0-3.6 1.4-3.6 3.6l-5.8 70l5.8 67.8c0 2.2 1.7 3.6 3.6 3.6s3.6-1.4 3.9-3.6l6.4-67.8l-6.4-70c-.3-2.2-2-3.6-3.9-3.6m241.4-110.9c-1.1-.8-2.8-1.4-4.2-1.4c-2.2 0-4.2.8-5.6 1.9c-1.9 1.7-3.1 4.2-3.3 6.7v.8l-3.3 176.7l1.7 32.5l1.7 31.7c.3 4.7 4.2 8.6 8.9 8.6s8.6-3.9 8.6-8.6l3.9-64.2l-3.9-177.5c-.4-3-2-5.8-4.5-7.2m-26.7 15.3c-1.4-.8-2.8-1.4-4.4-1.4s-3.1.6-4.4 1.4c-2.2 1.4-3.6 3.9-3.6 6.7l-.3 1.7l-2.8 160.8s0 .3 3.1 65.6v.3c0 1.7.6 3.3 1.7 4.7c1.7 1.9 3.9 3.1 6.4 3.1c2.2 0 4.2-1.1 5.6-2.5c1.7-1.4 2.5-3.3 2.5-5.6l.3-6.7l3.1-58.6l-3.3-162.8c-.3-2.8-1.7-5.3-3.9-6.7m-111.4 22.5c-3.1 0-5.8 2.8-5.8 6.1l-4.4 140.6l4.4 67.2c.3 3.3 2.8 5.8 5.8 5.8c3.3 0 5.8-2.5 6.1-5.8l5-67.2l-5-140.6c-.2-3.3-2.7-6.1-6.1-6.1m376.7 62.8c-10.8 0-21.1 2.2-30.6 6.1c-6.4-70.8-65.8-126.4-138.3-126.4c-17.8 0-35 3.3-50.3 9.4c-6.1 2.2-7.8 4.4-7.8 9.2v249.7c0 5 3.9 8.6 8.6 9.2h218.3c43.3 0 78.6-35 78.6-78.3c.1-43.6-35.2-78.9-78.5-78.9m-296.7-60.3c-4.2 0-7.5 3.3-7.8 7.8l-3.3 136.7l3.3 65.6c.3 4.2 3.6 7.5 7.8 7.5c4.2 0 7.5-3.3 7.5-7.5l3.9-65.6l-3.9-136.7c-.3-4.5-3.3-7.8-7.5-7.8m-53.6-7.8c-3.3 0-6.4 3.1-6.4 6.7l-3.9 145.3l3.9 66.9c.3 3.6 3.1 6.4 6.4 6.4c3.6 0 6.4-2.8 6.7-6.4l4.4-66.9l-4.4-145.3c-.3-3.6-3.1-6.7-6.7-6.7m26.7 3.4c-3.9 0-6.9 3.1-6.9 6.9L227 321.3l3.9 66.4c.3 3.9 3.1 6.9 6.9 6.9s6.9-3.1 6.9-6.9l4.2-66.4l-4.2-141.7c0-3.9-3-6.9-6.9-6.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sourcetree(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427.2 203c0-112.1-90.9-203-203-203C112.1-.2 21.2 90.6 21 202.6A202.86 202.86 0 0 0 161.5 396v101.7a14.3 14.3 0 0 0 14.3 14.3h96.4a14.3 14.3 0 0 0 14.3-14.3V396.1A203.18 203.18 0 0 0 427.2 203m-271.6 0c0-90.8 137.3-90.8 137.3 0c-.1 89.9-137.3 91-137.3 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Speakap(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 391.78C-15.41 303.59-8 167.42 80.64 87.64s224.8-73 304.21 15.24s72 224.36-16.64 304.14c-18.74 16.87 64 43.09 42 52.26c-82.06 34.21-253.91 35-346.23-67.5zm213.31-211.6l38.5-40.86c-9.61-8.89-32-26.83-76.17-27.6c-52.33-.91-95.86 28.3-96.77 80c-.2 11.33.29 36.72 29.42 54.83c34.46 21.42 86.52 21.51 86 52.26c-.37 21.28-26.42 25.81-38.59 25.6c-3-.05-30.23-.46-47.61-24.62l-40 42.61c28.16 27 59 32.62 83.49 33.05c10.23.18 96.42.33 97.84-81c.28-15.81-2.07-39.72-28.86-56.59c-34.36-21.64-85-19.45-84.43-49.75c.41-23.25 31-25.37 37.53-25.26c.43 0 26.62.26 39.62 17.37z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpeakerDeck(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.86 296H100a100 100 0 0 1 0-200h132.84a40 40 0 0 1 0 80H98c-26.47 0-26.45 40 0 40h113.82a100 100 0 0 1 0 200H40a40 40 0 0 1 0-80h173.86c26.48 0 26.46-40 0-40M298 416a120.21 120.21 0 0 0 51.11-80h64.55a19.83 19.83 0 0 0 19.66-20V196a19.83 19.83 0 0 0-19.66-20H296.42a60.77 60.77 0 0 0 0-80h136.93c43.44 0 78.65 35.82 78.65 80v160c0 44.18-35.21 80-78.65 80z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spotify(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M248 8C111.1 8 0 119.1 0 256s111.1 248 248 248s248-111.1 248-248S384.9 8 248 8m100.7 364.9c-4.2 0-6.8-1.3-10.7-3.6c-62.4-37.6-135-39.2-206.7-24.5c-3.9 1-9 2.6-11.9 2.6c-9.7 0-15.8-7.7-15.8-15.8c0-10.3 6.1-15.2 13.6-16.8c81.9-18.1 165.6-16.5 237 26.2c6.1 3.9 9.7 7.4 9.7 16.5s-7.1 15.4-15.2 15.4m26.9-65.6c-5.2 0-8.7-2.3-12.3-4.2c-62.5-37-155.7-51.9-238.6-29.4c-4.8 1.3-7.4 2.6-11.9 2.6c-10.7 0-19.4-8.7-19.4-19.4s5.2-17.8 15.5-20.7c27.8-7.8 56.2-13.6 97.8-13.6c64.9 0 127.6 16.1 177 45.5c8.1 4.8 11.3 11 11.3 19.7c-.1 10.8-8.5 19.5-19.4 19.5m31-76.2c-5.2 0-8.4-1.3-12.9-3.9c-71.2-42.5-198.5-52.7-280.9-29.7c-3.6 1-8.1 2.6-12.9 2.6c-13.2 0-23.3-10.3-23.3-23.6c0-13.6 8.4-21.3 17.4-23.9c35.2-10.3 74.6-15.2 117.5-15.2c73 0 149.5 15.2 205.4 47.8c7.8 4.5 12.9 10.7 12.9 22.6c0 13.6-11 23.3-23.2 23.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Squarespace(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M186.12 343.34c-9.65 9.65-9.65 25.29 0 34.94c9.65 9.65 25.29 9.65 34.94 0L378.24 221.1c19.29-19.29 50.57-19.29 69.86 0s19.29 50.57 0 69.86L293.95 445.1c19.27 19.29 50.53 19.31 69.82.04l.04-.04l119.25-119.24c38.59-38.59 38.59-101.14 0-139.72c-38.59-38.59-101.15-38.59-139.72 0zm244.53-104.8c-9.65-9.65-25.29-9.65-34.93 0l-157.2 157.18c-19.27 19.29-50.53 19.31-69.82.05l-.05-.05c-9.64-9.64-25.27-9.65-34.92-.01l-.01.01c-9.65 9.64-9.66 25.28-.02 34.93l.02.02c38.58 38.57 101.14 38.57 139.72 0l157.2-157.2c9.65-9.65 9.65-25.29.01-34.93m-261.99 87.33l157.18-157.18c9.64-9.65 9.64-25.29 0-34.94c-9.64-9.64-25.27-9.64-34.91 0L133.72 290.93c-19.28 19.29-50.56 19.3-69.85.01l-.01-.01c-19.29-19.28-19.31-50.54-.03-69.84l.03-.03L218.03 66.89c-19.28-19.29-50.55-19.3-69.85-.02l-.02.02L28.93 186.14c-38.58 38.59-38.58 101.14 0 139.72c38.6 38.59 101.13 38.59 139.73.01m-87.33-52.4c9.64 9.64 25.27 9.64 34.91 0l157.21-157.19c19.28-19.29 50.55-19.3 69.84-.02l.02.02c9.65 9.65 25.29 9.65 34.93 0c9.65-9.65 9.65-25.29 0-34.93c-38.59-38.59-101.13-38.59-139.72 0L81.33 238.54c-9.65 9.64-9.65 25.28-.01 34.93z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StackExchange(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.7 332.3h412.7v22c0 37.7-29.3 68-65.3 68h-19L259.3 512v-89.7H83c-36 0-65.3-30.3-65.3-68zm0-23.6h412.7v-85H17.7zm0-109.4h412.7v-85H17.7zM365 0H83C47 0 17.7 30.3 17.7 67.7V90h412.7V67.7C430.3 30.3 401 0 365 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StackOverflow(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M290.7 311L95 269.7L86.8 309l195.7 41zm51-87L188.2 95.7l-25.5 30.8l153.5 128.3zm-31.2 39.7L129.2 179l-16.7 36.5L293.7 300zM262 32l-32 24l119.3 160.3l32-24zm20.5 328h-200v39.7h200zm39.7 80H42.7V320h-40v160h359.5V320h-40z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stackpath(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M244.6 232.4c0 8.5-4.26 20.49-21.34 20.49h-19.61v-41.47h19.61c17.13 0 21.34 12.36 21.34 20.98M448 32v448H0V32zM151.3 287.84c0-21.24-12.12-34.54-46.72-44.85c-20.57-7.41-26-10.91-26-18.63s7-14.61 20.41-14.61c14.09 0 20.79 8.45 20.79 18.35h30.7l.19-.57c.5-19.57-15.06-41.65-51.12-41.65c-23.37 0-52.55 10.75-52.55 38.29c0 19.4 9.25 31.29 50.74 44.37c17.26 6.15 21.91 10.4 21.91 19.48c0 15.2-19.13 14.23-19.47 14.23c-20.4 0-25.65-9.1-25.65-21.9h-30.8l-.18.56c-.68 31.32 28.38 45.22 56.63 45.22c29.98 0 51.12-13.55 51.12-38.29m125.38-55.63c0-25.3-18.43-45.46-53.42-45.46h-51.78v138.18h32.17v-47.36h19.61c30.25 0 53.42-15.95 53.42-45.36M297.94 325L347 186.78h-31.09L268 325zm106.52-138.22h-31.09L325.46 325h29.94z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Staylinked(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m382.7 292.5l2.7 2.7l-170-167.3c-3.5-3.5-9.7-3.7-13.8-.5L144.3 171c-4.2 3.2-4.6 8.7-1.1 12.2l68.1 64.3c3.6 3.5 9.9 3.7 14 .5l.1-.1c4.1-3.2 10.4-3 14 .5l84 81.3c3.6 3.5 3.2 9-.9 12.2l-93.2 74c-4.2 3.3-10.5 3.1-14.2-.4L63.2 268c-3.5-3.5-9.7-3.7-13.9-.5L3.5 302.4c-4.2 3.2-4.7 8.7-1.2 12.2L211 510.7s7.4 6.8 17.3-.8l198-163.9c4-3.2 4.4-8.7.7-12.2zm54.5-83.4L226.7 2.5c-1.5-1.2-8-5.5-16.3 1.1L3.6 165.7c-4.2 3.2-4.8 8.7-1.2 12.2l42.3 41.7l171.7 165.1c3.7 3.5 10.1 3.7 14.3.4l50.2-38.8l-.3-.3l7.7-6c4.2-3.2 4.6-8.7.9-12.2l-57.1-54.4c-3.6-3.5-10-3.7-14.2-.5l-.1.1c-4.2 3.2-10.5 3.1-14.2-.4L109 180.8c-3.6-3.5-3.1-8.9 1.1-12.2l92.2-71.5c4.1-3.2 10.3-3 13.9.5l160.4 159c3.7 3.5 10 3.7 14.1.5l45.8-35.8c4.1-3.2 4.4-8.7.7-12.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Steam(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M496 256c0 137-111.2 248-248.4 248c-113.8 0-209.6-76.3-239-180.4l95.2 39.3c6.4 32.1 34.9 56.4 68.9 56.4c39.2 0 71.9-32.4 70.2-73.5l84.5-60.2c52.1 1.3 95.8-40.9 95.8-93.5c0-51.6-42-93.5-93.7-93.5s-93.7 42-93.7 93.5v1.2L176.6 279c-15.5-.9-30.7 3.4-43.5 12.1L0 236.1C10.2 108.4 117.1 8 247.6 8C384.8 8 496 119 496 256M155.7 384.3l-30.5-12.6a52.79 52.79 0 0 0 27.2 25.8c26.9 11.2 57.8-1.6 69-28.4c5.4-13 5.5-27.3.1-40.3c-5.4-13-15.5-23.2-28.5-28.6c-12.9-5.4-26.7-5.2-38.9-.6l31.5 13c19.8 8.2 29.2 30.9 20.9 50.7c-8.3 19.9-31 29.2-50.8 21m173.8-129.9c-34.4 0-62.4-28-62.4-62.3s28-62.3 62.4-62.3s62.4 28 62.4 62.3s-27.9 62.3-62.4 62.3m.1-15.6c25.9 0 46.9-21 46.9-46.8c0-25.9-21-46.8-46.9-46.8s-46.9 21-46.9 46.8c.1 25.8 21.1 46.8 46.9 46.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SteamSquare(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M185.2 356.5c7.7-18.5-1-39.7-19.6-47.4l-29.5-12.2c11.4-4.3 24.3-4.5 36.4.5c12.2 5.1 21.6 14.6 26.7 26.7c5 12.2 5 25.6-.1 37.7c-10.5 25.1-39.4 37-64.6 26.5c-11.6-4.8-20.4-13.6-25.4-24.2l28.5 11.8c18.6 7.8 39.9-.9 47.6-19.4M400 32H48C21.5 32 0 53.5 0 80v160.7l116.6 48.1c12-8.2 26.2-12.1 40.7-11.3l55.4-80.2v-1.1c0-48.2 39.3-87.5 87.6-87.5s87.6 39.3 87.6 87.5c0 49.2-40.9 88.7-89.6 87.5l-79 56.3c1.6 38.5-29.1 68.8-65.7 68.8c-31.8 0-58.5-22.7-64.5-52.7L0 319.2V432c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V80c0-26.5-21.5-48-48-48m-99.7 222.5c-32.2 0-58.4-26.1-58.4-58.3s26.2-58.3 58.4-58.3s58.4 26.2 58.4 58.3s-26.2 58.3-58.4 58.3m.1-14.6c24.2 0 43.9-19.6 43.9-43.8c0-24.2-19.6-43.8-43.9-43.8c-24.2 0-43.9 19.6-43.9 43.8c0 24.2 19.7 43.8 43.9 43.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SteamSymbol(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M395.5 177.5c0 33.8-27.5 61-61 61c-33.8 0-61-27.3-61-61s27.3-61 61-61c33.5 0 61 27.2 61 61m52.5.2c0 63-51 113.8-113.7 113.8L225 371.3c-4 43-40.5 76.8-84.5 76.8c-40.5 0-74.7-28.8-83-67L0 358V250.7L97.2 290c15.1-9.2 32.2-13.3 52-11.5l71-101.7c.5-62.3 51.5-112.8 114-112.8C397 64 448 115 448 177.7M203 363c0-34.7-27.8-62.5-62.5-62.5c-4.5 0-9 .5-13.5 1.5l26 10.5c25.5 10.2 38 39 27.7 64.5c-10.2 25.5-39.2 38-64.7 27.5c-10.2-4-20.5-8.3-30.7-12.2c10.5 19.7 31.2 33.2 55.2 33.2c34.7 0 62.5-27.8 62.5-62.5m207.5-185.3c0-42-34.3-76.2-76.2-76.2c-42.3 0-76.5 34.2-76.5 76.2c0 42.2 34.3 76.2 76.5 76.2c41.9.1 76.2-33.9 76.2-76.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StickerMule(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M561.7 199.6c-1.3.3.3 0 0 0m-6.2-77.4c-7.7-22.3-5.1-7.2-13.4-36.9c-1.6-6.5-3.6-14.5-6.2-20c-4.4-8.7-4.6-7.5-4.6-9.5c0-5.3 30.7-45.3 19-46.9c-5.7-.6-12.2 11.6-20.6 17c-8.6 4.2-8 5-10.3 5c-2.6 0-5.7-3-6.2-5c-2-5.7 1.9-25.9-3.6-25.9c-3.6 0-12.3 24.8-17 25.8c-5.2 1.3-27.9-11.4-75.1 18c-25.3 13.2-86.9 65.2-87 65.3c-6.7 4.7-20 4.7-35.5 16c-44.4 30.1-109.6 9.4-110.7 9c-110.6-26.8-128-15.2-159 11.5c-20.8 17.9-23.7 36.5-24.2 38.9c-4.2 20.4 5.2 48.3 6.7 64.3c1.8 19.3-2.7 17.7 7.7 98.3c.5 1 4.1 0 5.1 1.5c0 8.4-3.8 12.1-4.1 13c-1.5 4.5-1.5 10.5 0 16c2.3 8.2 8.2 37.2 8.2 46.9c0 41.8.4 44 2.6 49.4c3.9 10 12.5 9.1 17 12c3.1 3.5-.5 8.5 1 12.5c.5 2 3.6 4 6.2 5c9.2 3.6 27 .3 29.9-2.5c1.6-1.5.5-4.5 3.1-5c5.1 0 10.8-.5 14.4-2.5c5.1-2.5 4.1-6 1.5-10.5c-.4-.8-7-13.3-9.8-16c-2.1-2-5.1-3-7.2-4.5c-5.8-4.9-10.3-19.4-10.3-19.5c-4.6-19.4-10.3-46.3-4.1-66.8c4.6-17.2 39.5-87.7 39.6-87.8c4.1-6.5 17-11.5 27.3-7c6 1.9 19.3 22 65.4 30.9c47.9 8.7 97.4-2 112.2-2c2.8 2-1.9 13-.5 38.9c0 26.4-.4 13.7-4.1 29.9c-2.2 9.7 3.4 23.2-1.5 46.9c-1.4 9.8-9.9 32.7-8.2 43.4c.5 1 1 2 1.5 3.5c.5 4.5 1.5 8.5 4.6 10c7.3 3.6 12-3.5 9.8 11.5c-.7 3.1-2.6 12 1.5 15c4.4 3.7 30.6 3.4 36.5.5c2.6-1.5 1.6-4.5 6.4-7.4c1.9-.9 11.3-.4 11.3-6.5c.3-1.8-9.2-19.9-9.3-20c-2.6-3.5-9.2-4.5-11.3-8c-6.9-10.1-1.7-52.6.5-59.4c3-11 5.6-22.4 8.7-32.4c11-42.5 10.3-50.6 16.5-68.3c.8-1.8 6.4-23.1 10.3-29.9c9.3-17 21.7-32.4 33.5-47.4c18-22.9 34-46.9 52-69.8c6.1-7 8.2-13.7 18-8c10.8 5.7 21.6 7 31.9 17c14.6 12.8 10.2 18.2 11.8 22.9c1.5 5 7.7 10.5 14.9 9.5c10.4-2 13-2.5 13.4-2.5c2.6-.5 5.7-5 7.2-8c3.1-5.5 7.2-9 7.2-16.5c0-7.7-.4-2.8-20.6-52.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Strava(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M158.4 0L7 292h89.2l62.2-116.1L220.1 292h88.5zm150.2 292l-43.9 88.2l-44.6-88.2h-67.6l112.2 220l111.5-220z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stripe(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m165 144.7l-43.3 9.2l-.2 142.4c0 26.3 19.8 43.3 46.1 43.3c14.6 0 25.3-2.7 31.2-5.9v-33.8c-5.7 2.3-33.7 10.5-33.7-15.7V221h33.7v-37.8h-33.7zm89.1 51.6l-2.7-13.1H213v153.2h44.3V233.3c10.5-13.8 28.2-11.1 33.9-9.3v-40.8c-6-2.1-26.7-6-37.1 13.1m92.3-72.3l-44.6 9.5v36.2l44.6-9.5zM44.9 228.3c0-6.9 5.8-9.6 15.1-9.7c13.5 0 30.7 4.1 44.2 11.4v-41.8c-14.7-5.8-29.4-8.1-44.1-8.1c-36 0-60 18.8-60 50.2c0 49.2 67.5 41.2 67.5 62.4c0 8.2-7.1 10.9-17 10.9c-14.7 0-33.7-6.1-48.6-14.2v40c16.5 7.1 33.2 10.1 48.5 10.1c36.9 0 62.3-15.8 62.3-47.8c0-52.9-67.9-43.4-67.9-63.4M640 261.6c0-45.5-22-81.4-64.2-81.4s-67.9 35.9-67.9 81.1c0 53.5 30.3 78.2 73.5 78.2c21.2 0 37.1-4.8 49.2-11.5v-33.4c-12.1 6.1-26 9.8-43.6 9.8c-17.3 0-32.5-6.1-34.5-26.9h86.9c.2-2.3.6-11.6.6-15.9m-87.9-16.8c0-20 12.3-28.4 23.4-28.4c10.9 0 22.5 8.4 22.5 28.4zm-112.9-64.6c-17.4 0-28.6 8.2-34.8 13.9l-2.3-11H363v204.8l44.4-9.4l.1-50.2c6.4 4.7 15.9 11.2 31.4 11.2c31.8 0 60.8-23.2 60.8-79.6c.1-51.6-29.3-79.7-60.5-79.7m-10.6 122.5c-10.4 0-16.6-3.8-20.9-8.4l-.3-66c4.6-5.1 11-8.8 21.2-8.8c16.2 0 27.4 18.2 27.4 41.4c.1 23.9-10.9 41.8-27.4 41.8m-126.7 33.7h44.6V183.2h-44.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StripeS(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M155.3 154.6c0-22.3 18.6-30.9 48.4-30.9c43.4 0 98.5 13.3 141.9 36.7V26.1C298.3 7.2 251.1 0 203.8 0C88.1 0 11 60.4 11 161.4c0 157.9 216.8 132.3 216.8 200.4c0 26.4-22.9 34.9-54.7 34.9c-47.2 0-108.2-19.5-156.1-45.5v128.5a396.09 396.09 0 0 0 156 32.4c118.6 0 200.3-51 200.3-153.6c0-170.2-218-139.7-218-203.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Studiovinari(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m480.3 187.7l4.2 28v28l-25.1 44.1l-39.8 78.4l-56.1 67.5l-79.1 37.8l-17.7 24.5l-7.7 12l-9.6 4s17.3-63.6 19.4-63.6c2.1 0 20.3.7 20.3.7l66.7-38.6l-92.5 26.1l-55.9 36.8l-22.8 28l-6.6 1.4l20.8-73.6l6.9-5.5l20.7 12.9l88.3-45.2l56.8-51.5l14.8-68.4l-125.4 23.3l15.2-18.2l-173.4-53.3l81.9-10.5l-166-122.9L133.5 108L32.2 0l252.9 126.6l-31.5-38L378 163L234.7 64l18.7 38.4l-49.6-18.1L158.3 0l194.6 122L310 66.2l108 96.4l12-8.9l-21-16.4l4.2-37.8L451 89.1l29.2 24.7l11.5 4.2l-7 6.2l8.5 12l-13.1 7.4l-10.3 20.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stumbleupon(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M502.9 266v69.7c0 62.1-50.3 112.4-112.4 112.4c-61.8 0-112.4-49.8-112.4-111.3v-70.2l34.3 16l51.1-15.2V338c0 14.7 12 26.5 26.7 26.5S417 352.7 417 338v-72zm-224.7-58.2l34.3 16l51.1-15.2V173c0-60.5-51.1-109-112.1-109c-60.8 0-112.1 48.2-112.1 108.2v162.4c0 14.9-12 26.7-26.7 26.7S86 349.5 86 334.6V266H0v69.7C0 397.7 50.3 448 112.4 448c61.6 0 112.4-49.5 112.4-110.8V176.9c0-14.7 12-26.7 26.7-26.7s26.7 12 26.7 26.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StumbleuponCircle(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 8C119 8 8 119 8 256s111 248 248 248s248-111 248-248S393 8 256 8m0 177.5c-9.8 0-17.8 8-17.8 17.8v106.9c0 40.9-33.9 73.9-74.9 73.9c-41.4 0-74.9-33.5-74.9-74.9v-46.5h57.3v45.8c0 10 8 17.8 17.8 17.8s17.8-7.9 17.8-17.8V200.1c0-40 34.2-72.1 74.7-72.1c40.7 0 74.7 32.3 74.7 72.6v23.7l-34.1 10.1l-22.9-10.7v-20.6c.1-9.6-7.9-17.6-17.7-17.6m167.6 123.6c0 41.4-33.5 74.9-74.9 74.9c-41.2 0-74.9-33.2-74.9-74.2V263l22.9 10.7l34.1-10.1v47.1c0 9.8 8 17.6 17.8 17.6s17.8-7.9 17.8-17.6v-48h57.3c-.1 45.9-.1 46.4-.1 46.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Superpowers(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 32c-83.3 11-166.8 22-250 33c-92 12.5-163.3 86.7-169 180c-3.3 55.5 18 109.5 57.8 148.2L0 480c83.3-11 166.5-22 249.8-33c91.8-12.5 163.3-86.8 168.7-179.8c3.5-55.5-18-109.5-57.7-148.2zm-79.7 232.3c-4.2 79.5-74 139.2-152.8 134.5c-79.5-4.7-140.7-71-136.3-151c4.5-79.2 74.3-139.3 153-134.5c79.3 4.7 140.5 71 136.1 151"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Supple(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 262.5c0 64.1-109 116.1-243.5 116.1c-24.8 0-48.6-1.8-71.1-5c7.7.4 15.5.6 23.4.6c134.5 0 243.5-56.9 243.5-127.1c0-29.4-19.1-56.4-51.2-78c60 21.1 98.9 55.1 98.9 93.4M47.7 227.9c-.1-70.2 108.8-127.3 243.3-127.6c7.9 0 15.6.2 23.3.5c-22.5-3.2-46.3-4.9-71-4.9C108.8 96.3-.1 148.5 0 212.6c.1 38.3 39.1 72.3 99.3 93.3c-32.3-21.5-51.5-48.6-51.6-78m60.2 39.9s10.5 13.2 29.3 13.2c17.9 0 28.4-11.5 28.4-25.1c0-28-40.2-25.1-40.2-39.7c0-5.4 5.3-9.1 12.5-9.1c5.7 0 11.3 2.6 11.3 6.6v3.9h14.2v-7.9c0-12.1-15.4-16.8-25.4-16.8c-16.5 0-28.5 10.2-28.5 24.1c0 26.6 40.2 25.4 40.2 39.9c0 6.6-5.8 10.1-12.3 10.1c-11.9 0-20.7-10.1-20.7-10.1zm120.8-73.6v54.4c0 11.3-7.1 17.8-17.8 17.8c-10.7 0-17.8-6.5-17.8-17.7v-54.5h-15.8v55c0 18.9 13.4 31.9 33.7 31.9c20.1 0 33.4-13 33.4-31.9v-55zm34.4 85.4h15.8v-29.5h15.5c16 0 27.2-11.5 27.2-28.1s-11.2-27.8-27.2-27.8h-39.1v13.4h7.8zm15.8-43v-29.1h12.9c8.7 0 13.7 5.7 13.7 14.4c0 8.9-5.1 14.7-14 14.7zm57 43h15.8v-29.5h15.5c16 0 27.2-11.5 27.2-28.1s-11.2-27.8-27.2-27.8h-39.1v13.4h7.8zm15.7-43v-29.1h12.9c8.7 0 13.7 5.7 13.7 14.4c0 8.9-5 14.7-14 14.7zm57.1 34.8c0 5.8 2.4 8.2 8.2 8.2h37.6c5.8 0 8.2-2.4 8.2-8.2v-13h-14.3v5.2c0 1.7-1 2.6-2.6 2.6h-18.6c-1.7 0-2.6-1-2.6-2.6v-61.2c0-5.7-2.4-8.2-8.2-8.2H401v13.4h5.2c1.7 0 2.6 1 2.6 2.6v61.2zm63.4 0c0 5.8 2.4 8.2 8.2 8.2H519c5.7 0 8.2-2.4 8.2-8.2v-13h-14.3v5.2c0 1.7-1 2.6-2.6 2.6h-19.7c-1.7 0-2.6-1-2.6-2.6v-20.3h27.7v-13.4H488v-22.4h19.2c1.7 0 2.6 1 2.6 2.6v5.2H524v-13c0-5.7-2.5-8.2-8.2-8.2h-51.6v13.4h7.8v63.9zm58.9-76v5.9h1.6v-5.9h2.7v-1.2h-7v1.2zm5.7-1.2v7.1h1.5v-5.7l2.3 5.7h1.3l2.3-5.7v5.7h1.5v-7.1h-2.3l-2.1 5.1l-2.1-5.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Suse(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M471.08 102.66s-.3 18.3-.3 20.3c-9.1-3-74.4-24.1-135.7-26.3c-51.9-1.8-122.8-4.3-223 57.3c-19.4 12.4-73.9 46.1-99.6 109.7C7 277-.12 307 7 335.06a111 111 0 0 0 16.5 35.7c17.4 25 46.6 41.6 78.1 44.4c44.4 3.9 78.1-16 90-53.3c8.2-25.8 0-63.6-31.5-82.9c-25.6-15.7-53.3-12.1-69.2-1.6c-13.9 9.2-21.8 23.5-21.6 39.2c.3 27.8 24.3 42.6 41.5 42.6a49 49 0 0 0 15.8-2.7c6.5-1.8 13.3-6.5 13.3-14.9c0-12.1-11.6-14.8-16.8-13.9c-2.9.5-4.5 2-11.8 2.4c-2-.2-12-3.1-12-14V316c.2-12.3 13.2-18 25.5-16.9c32.3 2.8 47.7 40.7 28.5 65.7c-18.3 23.7-76.6 23.2-99.7-20.4c-26-49.2 12.7-111.2 87-98.4c33.2 5.7 83.6 35.5 102.4 104.3h45.9c-5.7-17.6-8.9-68.3 42.7-68.3c56.7 0 63.9 39.9 79.8 68.3H460c-12.8-18.3-21.7-38.7-18.9-55.8c5.6-33.8 39.7-18.4 82.4-17.4c66.5.4 102.1-27 103.1-28c3.7-3.1 6.5-15.8 7-17.7c1.3-5.1-3.2-2.4-3.2-2.4c-8.7 5.2-30.5 15.2-50.9 15.6c-25.3.5-76.2-25.4-81.6-28.2c-.3-.4.1 1.2-11-25.5c88.4 58.3 118.3 40.5 145.2 21.7c.8-.6 4.3-2.9 3.6-5.7c-13.8-48.1-22.4-62.7-34.5-69.6c-37-21.6-125-34.7-129.2-35.3c.1-.1-.9-.3-.9.7zm60.4 72.8a37.54 37.54 0 0 1 38.9-36.3c33.4 1.2 48.8 42.3 24.4 65.2c-24.2 22.7-64.4 4.6-63.3-28.9m38.6-25.3a26.27 26.27 0 1 0 25.4 27.2a26.19 26.19 0 0 0-25.4-27.2m4.3 28.8c-15.4 0-15.4-15.6 0-15.6s15.4 15.64 0 15.64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Swift(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 156.09c0-4.51-.08-9-.2-13.52a196.31 196.31 0 0 0-2.58-29.42a99.62 99.62 0 0 0-9.22-28A94.08 94.08 0 0 0 394.84 44a99.17 99.17 0 0 0-28-9.22a195 195 0 0 0-29.43-2.59c-4.51-.12-9-.17-13.52-.2H124.14c-4.51 0-9 .08-13.52.2c-2.45.07-4.91.15-7.37.27a171.68 171.68 0 0 0-22.06 2.32a103.06 103.06 0 0 0-21.21 6.1q-3.46 1.45-6.81 3.12a94.66 94.66 0 0 0-18.39 12.32c-1.88 1.61-3.69 3.28-5.43 5A93.86 93.86 0 0 0 12 85.17a99.45 99.45 0 0 0-9.22 28a196.31 196.31 0 0 0-2.54 29.4c-.13 4.51-.18 9-.21 13.52v199.83c0 4.51.08 9 .21 13.51a196.08 196.08 0 0 0 2.58 29.42a99.3 99.3 0 0 0 9.22 28A94.31 94.31 0 0 0 53.17 468a99.47 99.47 0 0 0 28 9.21a195 195 0 0 0 29.43 2.59c4.5.12 9 .17 13.52.2h199.79c4.51 0 9-.08 13.52-.2a196.59 196.59 0 0 0 29.44-2.59a99.57 99.57 0 0 0 28-9.21A94.22 94.22 0 0 0 436 426.84a99.3 99.3 0 0 0 9.22-28a194.79 194.79 0 0 0 2.59-29.42c.12-4.5.17-9 .2-13.51V172.14c-.01-5.35-.01-10.7-.01-16.05m-69.88 241c-20-38.93-57.23-29.27-76.31-19.47c-1.72 1-3.48 2-5.25 3l-.42.25c-39.5 21-92.53 22.54-145.85-.38A234.64 234.64 0 0 1 45 290.12a230.63 230.63 0 0 0 39.17 23.37c56.36 26.4 113 24.49 153 0c-57-43.85-104.6-101-141.09-147.22a197.09 197.09 0 0 1-18.78-25.9c43.7 40 112.7 90.22 137.48 104.12c-52.57-55.49-98.89-123.94-96.72-121.74c82.79 83.42 159.18 130.59 159.18 130.59c2.88 1.58 5 2.85 6.73 4a127.44 127.44 0 0 0 4.16-12.47c13.22-48.33-1.66-103.58-35.31-149.2C329.61 141.75 375 229.34 356.4 303.42c-.44 1.73-.95 3.4-1.44 5.09c38.52 47.4 28.04 98.17 23.13 88.59z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Symfony(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 8C119 8 8 119 8 256s111 248 248 248s248-111 248-248S393 8 256 8m133.74 143.54c-11.47.41-19.4-6.45-19.77-16.87c-.27-9.18 6.68-13.44 6.53-18.85c-.23-6.55-10.16-6.82-12.87-6.67c-39.78 1.29-48.59 57-58.89 113.85c21.43 3.15 36.65-.72 45.14-6.22c12-7.75-3.34-15.72-1.42-24.56c4-18.16 32.55-19 32 5.3c-.36 17.86-25.92 41.81-77.6 35.7c-10.76 59.52-18.35 115-58.2 161.72c-29 34.46-58.4 39.82-71.58 40.26c-24.65.85-41-12.31-41.58-29.84c-.56-17 14.45-26.26 24.31-26.59c21.89-.75 30.12 25.67 14.88 34c-12.09 9.71.11 12.61 2.05 12.55c10.42-.36 17.34-5.51 22.18-9c24-20 33.24-54.86 45.35-118.35c8.19-49.66 17-78 18.23-82c-16.93-12.75-27.08-28.55-49.85-34.72c-15.61-4.23-25.12-.63-31.81 7.83c-7.92 10-5.29 23 2.37 30.7l12.63 14c15.51 17.93 24 31.87 20.8 50.62c-5.06 29.93-40.72 52.9-82.88 39.94c-36-11.11-42.7-36.56-38.38-50.62c7.51-24.15 42.36-11.72 34.62 13.6c-2.79 8.6-4.92 8.68-6.28 13.07c-4.56 14.77 41.85 28.4 51-1.39c4.47-14.52-5.3-21.71-22.25-39.85c-28.47-31.75-16-65.49 2.95-79.67C204.23 140.13 251.94 197 262 205.29c37.17-109 100.53-105.46 102.43-105.53c25.16-.81 44.19 10.59 44.83 28.65c.25 7.69-4.17 22.59-19.52 23.13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Teamspeak(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M244.2 346.79c2.4-12.3-12-30-32.4-48.7c-20.9-19.2-48.2-39.1-63.4-46.6c-21.7-12-41.7-1.8-46.3 22.7c-5 26.2 0 51.4 14.5 73.9c10.2 15.5 25.4 22.7 43.4 24c11.6.6 52.5 2.2 61.7-1c11.9-4.3 20.1-11.8 22.5-24.3m205 20.8a5.22 5.22 0 0 0-8.3 2.4c-8 25.4-44.7 112.5-172.1 121.5c-149.7 10.5 80.3 43.6 145.4-6.4c22.7-17.4 47.6-35 46.6-85.4c-.4-10.1-4.9-26.69-11.6-32.1m62-122.4c-.3-18.9-8.6-33.4-26-42.2c-2.9-1.3-5-2.7-5.9-6.4A222.64 222.64 0 0 0 438.9 103c-1.1-1.5-3.5-3.2-2.2-5c8.5-11.5-.3-18-7-24.4Q321.4-31.11 177.4 13.09c-40.1 12.3-73.9 35.6-102 67.4c-4 4.3-6.7 9.1-3 14.5c3 4 1.3 6.2-1 9.3C51.6 132 38.2 162.59 32.1 196c-.7 4.3-2.9 6-6.4 7.8c-14.2 7-22.5 18.5-24.9 34L0 264.29v20.9c0 30.8 21 50.4 51.8 49c7.7-.3 11.7-4.3 12-11.5c2-77.5-2.4-95.4 3.7-125.8C92.1 72.39 234.3 5 345.3 65.39C411.4 102 445.7 159 447.6 234.79c.8 28.2 0 56.5 0 84.6c0 7 2.2 12.5 9.4 14.2c24.1 5 49.2-12 53.2-36.7c2.9-17.1 1-34.5 1-51.7m-159.6 131.5c36.5 2.8 59.3-28.5 58.4-60.5c-2.1-45.2-66.2-16.5-87.8-8c-73.2 28.1-45 54.9-22.2 60.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Telegram(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M248 8C111 8 0 119 0 256s111 248 248 248s248-111 248-248S385 8 248 8m121.8 169.9l-40.7 191.8c-3 13.6-11.1 16.9-22.4 10.5l-62-45.7l-29.9 28.8c-3.3 3.3-6.1 6.1-12.5 6.1l4.4-63.1l114.9-103.8c5-4.4-1.1-6.9-7.7-2.5l-142 89.4l-61.2-19.1c-13.3-4.2-13.6-13.3 2.8-19.7l239.1-92.2c11.1-4 20.8 2.7 17.2 19.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelegramPlane(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m446.7 98.6l-67.6 318.8c-5.1 22.5-18.4 28.1-37.3 17.5l-103-75.9l-49.7 47.8c-5.5 5.5-10.1 10.1-20.7 10.1l7.4-104.9l190.9-172.5c8.3-7.4-1.8-11.5-12.9-4.1L117.8 284L16.2 252.2c-22.1-6.9-22.5-22.1 4.6-32.7L418.2 66.4c18.4-6.9 34.5 4.1 28.5 32.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TencentWeibo(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M72.3 495.8c1.4 19.9-27.6 22.2-29.7 2.9C31 368.8 73.7 259.2 144 185.5c-15.6-34 9.2-77.1 50.6-77.1c30.3 0 55.1 24.6 55.1 55.1c0 44-49.5 70.8-86.9 45.1c-65.7 71.3-101.4 169.8-90.5 287.2M192 .1C66.1.1-12.3 134.3 43.7 242.4C52.4 259.8 79 246.9 70 229C23.7 136.4 91 29.8 192 29.8c75.4 0 136.9 61.4 136.9 136.9c0 90.8-86.9 153.9-167.7 133.1c-19.1-4.1-25.6 24.4-6.6 29.1c110.7 23.2 204-60 204-162.3C358.6 74.7 284 .1 192 .1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TheRedYeti(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m488.23 241.7l20.7 7.1c-9.6-23.9-23.9-37-31.7-44.8l7.1-18.2c.2 0 12.3-27.8-2.5-30.7c-.6-11.3-6.6-27-18.4-27c-7.6-10.6-17.7-12.3-30.7-5.9a122.2 122.2 0 0 0-25.3 16.5c-5.3-6.4-3 .4-3-29.8c-37.1-24.3-45.4-11.7-74.8 3l.5.5a239.36 239.36 0 0 0-68.4-13.3c-5.5-8.7-18.6-19.1-25.1-25.1l24.8 7.1c-5.5-5.5-26.8-12.9-34.2-15.2c18.2-4.1 29.8-20.8 42.5-33c-34.9-10.1-67.9-5.9-97.9 11.8l12-44.2L182 0c-31.6 24.2-33 41.9-33.7 45.5c-.9-2.4-6.3-19.6-15.2-27a35.12 35.12 0 0 0-.5 25.3c3 8.4 5.9 14.8 8.4 18.9c-16-3.3-28.3-4.9-49.2 0h-3.7l33 14.3a194.26 194.26 0 0 0-46.7 67.4l-1.7 8.4l1.7 1.7l7.6-4.7c-3.3 11.6-5.3 19.4-6.6 25.8a200.18 200.18 0 0 0-27.8 40.3c-15 1-31.8 10.8-40.3 14.3l3 3.4l28.8 1c-.5 1-.7 2.2-1.2 3.2c-7.3 6.4-39.8 37.7-33 80.7l20.2-22.4c.5 1.7.7 3.4 1.2 5.2c0 25.5.4 89.6 64.9 150.5c43.6 40 96 60.2 157.5 60.2c121.7 0 223-87.3 223-211.5c6.8-9.7-1.2 3 16.7-25.1l13 14.3l2.5-.5A181.84 181.84 0 0 0 495 255a44.74 44.74 0 0 0-6.8-13.3zM398 111.2l-.5 21.9c5.5 18.1 16.9 17.2 22.4 17.2l-3.4-4.7l22.4-5.4a242.44 242.44 0 0 1-27 0c12.8-2.1 33.3-29 43-11.3c3.4 7.6 6.4 17.2 9.3 27.8l1.7-5.9a56.38 56.38 0 0 1-1.7-15.2c5.4.5 8.8 3.4 9.3 10.1c.5 6.4 1.7 14.8 3.4 25.3l4.7-11.3c4.6 0 4.5-3.6-2.5 20.7c-20.9-8.7-35.1-8.4-46.5-8.4l18.2-16c-25.3 8.2-33 10.8-54.8 20.9c-1.1-5.4-5-13.5-16-19.9c-3.2 3.8-2.8.9-.7 14.8h-2.5a62.32 62.32 0 0 0-8.4-23.1l4.2-3.4c8.4-7.1 11.8-14.3 10.6-21.9c-.5-6.4-5.4-13.5-13.5-20.7c5.6-3.4 15.2-.4 28.3 8.5m-39.6-10.1c2.7 1.9 11.4 5.4 18.9 17.2c4.2 8.4 4 9.8 3.4 11.1c-.5 2.4-.5 4.3-3 7.1c-1.7 2.5-5.4 4.7-11.8 7.6c-7.6-13-16.5-23.6-27.8-31.2zM91 143.1l1.2-1.7c1.2-2.9 4.2-7.6 9.3-15.2l2.5-3.4l-13 12.3l5.4-4.7l-10.1 9.3l-4.2 1.2c12.3-24.1 23.1-41.3 32.5-50.2c9.3-9.3 16-16 20.2-19.4l-6.4 1.2c-11.3-4.2-19.4-7.1-24.8-8.4c2.5-.5 3.7-.5 3.2-.5c10.3 0 17.5.5 20.9 1.2a52.35 52.35 0 0 0 16 2.5l.5-1.7l-8.4-35.8l13.5 29a42.89 42.89 0 0 0 5.9-14.3c1.7-6.4 5.4-13 10.1-19.4s7.6-10.6 9.3-11.3a234.68 234.68 0 0 0-6.4 25.3l-1.7 7.1l-.5 4.7l2.5 2.5C190.4 39.9 214 34 239.8 34.5l21.1.5c-11.8 13.5-27.8 21.9-48.5 24.8a201.26 201.26 0 0 1-23.4 2.9l-.2-.5l-2.5-1.2a20.75 20.75 0 0 0-14 2c-2.5-.2-4.9-.5-7.1-.7l-2.5 1.7l.5 1.2c2 .2 3.9.5 6.2.7l-2 3.4l3.4-.5l-10.6 11.3c-4.2 3-5.4 6.4-4.2 9.3l5.4-3.4h1.2a39.4 39.4 0 0 1 25.3-15.2v-3c6.4.5 13 1 19.4 1.2c6.4 0 8.4.5 5.4 1.2a189.6 189.6 0 0 1 20.7 13.5c13.5 10.1 23.6 21.9 30 35.4c8.8 18.2 13.5 37.1 13.5 56.6a141.13 141.13 0 0 1-3 28.3a209.91 209.91 0 0 1-16 46l2.5.5c18.2-19.7 41.9-16 49.2-16l-6.4 5.9l22.4 17.7l-1.7 30.7c-5.4-12.3-16.5-21.1-33-27.8c16.5 14.8 23.6 21.1 21.9 20.2c-4.8-2.8-3.5-1.9-10.8-3.7c4.1 4.1 17.5 18.8 18.2 20.7l.2.2l-.2.2c0 1.8 1.6-1.2-14 22.9c-75.2-15.3-106.27-42.7-141.2-63.2l11.8 1.2c-11.8-18.5-15.6-17.7-38.4-26.1L149 225c-8.8-3-18.2-3-28.3.5l7.6-10.6l-1.2-1.7c-14.9 4.3-19.8 9.2-22.6 11.3c-1.1-5.5-2.8-12.4-12.3-28.8l-1.2 27l-13.2-5c1.5-25.2 5.4-50.5 13.2-74.6m276.5 330c-49.9 25-56.1 22.4-59 23.9c-29.8-11.8-50.9-31.7-63.5-58.8l30 16.5c-9.8-9.3-18.3-16.5-38.4-44.3l11.8 23.1l-17.7-7.6c14.2 21.1 23.5 51.7 66.6 73.5c-120.8 24.2-199-72.1-200.9-74.3a262.57 262.57 0 0 0 35.4 24.8c3.4 1.7 7.1 2.5 10.1 1.2l-16-20.7c9.2 4.2 9.5 4.5 69.1 29c-42.5-20.7-73.8-40.8-93.2-60.2c-.5 6.4-1.2 10.1-1.2 10.1a80.25 80.25 0 0 1 20.7 26.6c-39-18.9-57.6-47.6-71.3-82.6c49.9 55.1 118.9 37.5 120.5 37.1c34.8 16.4 69.9 23.6 113.9 10.6c3.3 0 20.3 17 25.3 39.1l4.2-3l-2.5-23.6c9 9 24.9 22.6 34.4 13c-15.6-5.3-23.5-9.5-29.5-31.7c4.6 4.2 7.6 9 27.8 15l1.2-1.2l-10.5-14.2c11.7-4.8-3.5 1 32-10.8c4.3 34.3 9 49.2.7 89.5m115.3-214.4l-2.5.5l3 9.3c-3.5 5.9-23.7 44.3-71.6 79.7c-39.5 29.8-76.6 39.1-80.9 40.3l-7.6-7.1l-1.2 3l14.3 16l-7.1-4.7l3.4 4.2h-1.2l-21.9-13.5l9.3 26.6l-19-27.9l-1.2 2.5l7.6 29c-6.1-8.2-21-32.6-56.8-39.6l32.5 21.2a214.82 214.82 0 0 1-93.2-6.4c-4.2-1.2-8.9-2.5-13.5-4.2l1.2-3l-44.8-22.4l26.1 22.4c-57.7 9.1-113-25.4-126.4-83.4l-2.5-16.4l-22.27 22.3c19.5-57.5 25.6-57.9 51.4-70.1c-9.1-5.3-1.6-3.3-38.4-9.3c15.8-5.8 33-15.4 73 5.2a18.5 18.5 0 0 1 3.7-1.7c.6-3.2.4-.8 1-11.8c3.9 10 3.6 8.7 3 9.3l1.7.5c12.7-6.5 8.9-4.5 17-8.9l-5.4 13.5l22.3-5.8l-8.4 8.4l2.5 2.5c4.5-1.8 30.3 3.4 40.8 16l-23.6-2.5c39.4 23 51.5 54 55.8 69.6l1.7-1.2c-2.8-22.3-12.4-33.9-16-40.1c4.2 5 39.2 34.6 110.4 46c-11.3-.5-23.1 5.4-34.9 18.9l46.7-20.2l-9.3 21.9c7.6-10.1 14.8-23.6 21.2-39.6v-.5l1.2-3l-1.2 16c13.5-41.8 25.3-78.5 35.4-109.7l13.5-27.8v-2l-5.4-4.2h10.1l5.9 4.2l2.5-1.2l-3.4-16l12.3 18.9l41.8-20.2l-14.8 13l.5 2.9l17.7-.5a184 184 0 0 1 33 4.2l-23.6 2.5l-1.2 3l26.6 23.1a254.21 254.21 0 0 1 27 32c-11.2-3.3-10.3-3.4-21.2-3.4l12.3 32.5zm-6.1-71.3l-3.9 13l-14.3-11.8zm-254.8 7.1c1.7 10.6 4.7 17.7 8.8 21.9c-9.3 6.6-27.5 13.9-46.5 16l.5 1.2a50.22 50.22 0 0 0 24.8-2.5l-7.1 13c4.2-1.7 10.1-7.1 17.7-14.8c11.9-5.5 12.7-5.1 20.2-16c-12.7-6.4-15.7-13.7-18.4-18.8m3.7-102.3c-6.4-3.4-10.6 3-12.3 18.9s2.5 29.5 11.8 39.6s18.2 10.6 26.1 3s3.4-23.6-11.3-47.7a39.57 39.57 0 0 0-14.27-13.8zm-4.7 46.3c5.4 2.2 10.5 1.9 12.3-10.6v-4.7l-1.2.5c-4.3-3.1-2.5-4.5-1.7-6.2l.5-.5c-.9-1.2-5-8.1-12.5 4.7c-.5-13.5.5-21.9 3-24.8c1.2-2.5 4.7-1.2 11.3 4.2c6.4 5.4 11.3 16 15.2 32.5c6.5 28-19.8 26.2-26.9 4.9m-45-5.5c1.6.3 9.3-1.1 9.3-14.8h-.5c-5.4-1.1-2.2-5.5-.7-5.9c-1.7-3-3.4-4.2-5.4-4.7c-8.1 0-11.6 12.7-8.1 21.2a7.51 7.51 0 0 0 5.43 4.2zM216 82.9l-2.5.5l.5 3a48.94 48.94 0 0 1 26.1 5.9c-2.5-5.5-10-14.3-28.3-14.3l.5 2.5zm-71.8 49.4c21.7 16.8 16.5 21.4 46.5 23.6l-2.9-4.7a42.67 42.67 0 0 0 14.8-28.3c1.7-16-1.2-29.5-8.8-41.3l13-7.6a2.26 2.26 0 0 0-.5-1.7a14.21 14.21 0 0 0-13.5 1.7c-12.7 6.7-28 20.9-29 22.4c-1.7 1.7-3.4 5.9-5.4 13.5a99.61 99.61 0 0 0-2.9 23.6c-4.7-8-10.5-6.4-19.9-5.9l7.1 7.6c-16.5 0-23.3 15.4-23.6 16c6.8 0 4.6-7.6 30-12.3c-4.3-6.3-3.3-5-4.9-6.6m18.7-18.7c1.2-7.6 3.4-13 6.4-17.2c5.4-6.4 10.6-10.1 16-11.8c4.2-1.7 7.1 1.2 10.1 9.3a72.14 72.14 0 0 1 3 25.3c-.5 9.3-3.4 17.2-8.4 23.1c-2.9 3.4-5.4 5.9-6.4 7.6a39.21 39.21 0 0 1-11.3-.5l-7.1-3.4l-5.4-6.4c.8-10 1.3-18.8 3.1-26m42 56.1c-34.8 14.4-34.7 14-36.1 14.3c-20.8 4.7-19-24.4-18.9-24.8l5.9-1.2l-.5-2.5c-20.2-2.6-31 4.2-32.5 4.9c.5.5 3 3.4 5.9 9.3c4.2-6.4 8.8-10.1 15.2-10.6a83.47 83.47 0 0 0 1.7 33.7c.1.5 2.6 17.4 27.5 24.1c11.3 3 27 1.2 48.9-5.4l-9.2.5c-4.2-14.8-6.4-24.8-5.9-29.5c11.3-8.8 21.9-11.3 30.7-7.6h2.5l-11.8-7.6l-7.1.5c-5.9 1.2-12.3 4.2-19.4 8.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Themeco(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M202.9 8.43c9.9-5.73 26-5.82 35.95-.21L430 115.85c10 5.6 18 19.44 18 30.86V364c0 11.44-8.06 25.29-18 31L238.81 503.74c-9.93 5.66-26 5.57-35.85-.21L17.86 395.12C8 389.34 0 375.38 0 364V146.71c0-11.44 8-25.36 17.91-31.08zm-77.4 199.83c-15.94 0-31.89.14-47.83.14v101.45H96.8V280h28.7c49.71 0 49.56-71.74 0-71.74m140.14 100.29l-30.73-34.64c37-7.51 34.8-65.23-10.87-65.51c-16.09 0-32.17-.14-48.26-.14v101.59h19.13v-33.91h18.41l29.56 33.91h22.76zm-41.59-82.32c23.34 0 23.26 32.46 0 32.46h-29.13v-32.46zm-95.56-1.6c21.18 0 21.11 38.85 0 38.85H96.18v-38.84zm192.65-18.25c-68.46 0-71 105.8 0 105.8c69.48-.01 69.41-105.8 0-105.8m0 17.39c44.12 0 44.8 70.86 0 70.86s-44.43-70.86 0-70.86"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Themeisle(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M208 88.286c0-10 6.286-21.714 17.715-21.714c11.142 0 17.714 11.714 17.714 21.714c0 10.285-6.572 21.714-17.714 21.714C214.286 110 208 98.571 208 88.286m304 160c0 36.001-11.429 102.286-36.286 129.714c-22.858 24.858-87.428 61.143-120.857 70.572l-1.143.286v32.571c0 16.286-12.572 30.571-29.143 30.571c-10 0-19.429-5.714-24.572-14.286c-5.427 8.572-14.856 14.286-24.856 14.286c-10 0-19.429-5.714-24.858-14.286c-5.142 8.572-14.571 14.286-24.57 14.286c-10.286 0-19.429-5.714-24.858-14.286c-5.143 8.572-14.571 14.286-24.571 14.286c-18.857 0-29.429-15.714-29.429-32.857c-16.286 12.285-35.715 19.428-56.571 19.428c-22 0-43.429-8.285-60.286-22.857c10.285-.286 20.571-2.286 30.285-5.714c-20.857-5.714-39.428-18.857-52-36.286c21.37 4.645 46.209 1.673 67.143-11.143c-22-22-56.571-58.857-68.572-87.428C1.143 321.714 0 303.714 0 289.429c0-49.714 20.286-160 86.286-160c10.571 0 18.857 4.858 23.143 14.857a158.792 158.792 0 0 1 12-15.428c2-2.572 5.714-5.429 7.143-8.286c7.999-12.571 11.714-21.142 21.714-34C182.571 45.428 232 17.143 285.143 17.143c6 0 12 .285 17.714 1.143C313.714 6.571 328.857 0 344.572 0c14.571 0 29.714 6 40 16.286c.857.858 1.428 2.286 1.428 3.428c0 3.714-10.285 13.429-12.857 16.286c4.286 1.429 15.714 6.858 15.714 12c0 2.857-2.857 5.143-4.571 7.143c31.429 27.714 49.429 67.143 56.286 108c4.286-5.143 10.285-8.572 17.143-8.572c10.571 0 20.857 7.144 28.571 14.001C507.143 187.143 512 221.714 512 248.286M188 89.428c0 18.286 12.571 37.143 32.286 37.143c19.714 0 32.285-18.857 32.285-37.143c0-18-12.571-36.857-32.285-36.857c-19.715 0-32.286 18.858-32.286 36.857M237.714 194c0-19.714 3.714-39.143 8.571-58.286c-52.039 79.534-13.531 184.571 68.858 184.571c21.428 0 42.571-7.714 60-20c2-7.429 3.714-14.857 3.714-22.572c0-14.286-6.286-21.428-20.572-21.428c-4.571 0-9.143.857-13.429 1.714c-63.343 12.668-107.142 3.669-107.142-63.999m-41.142 254.858c0-11.143-8.858-20.857-20.286-20.857c-11.429 0-20 9.715-20 20.857v32.571c0 11.143 8.571 21.142 20 21.142c11.428 0 20.286-9.715 20.286-21.142zm49.143 0c0-11.143-8.572-20.857-20-20.857c-11.429 0-20.286 9.715-20.286 20.857v32.571c0 11.143 8.857 21.142 20.286 21.142c11.428 0 20-10 20-21.142zm49.713 0c0-11.143-8.857-20.857-20.285-20.857c-11.429 0-20.286 9.715-20.286 20.857v32.571c0 11.143 8.857 21.142 20.286 21.142c11.428 0 20.285-9.715 20.285-21.142zm49.715 0c0-11.143-8.857-20.857-20.286-20.857c-11.428 0-20.286 9.715-20.286 20.857v32.571c0 11.143 8.858 21.142 20.286 21.142c11.429 0 20.286-10 20.286-21.142zM421.714 286c-30.857 59.142-90.285 102.572-158.571 102.572c-96.571 0-160.571-84.572-160.571-176.572c0-16.857 2-33.429 6-49.714c-20 33.715-29.714 72.572-29.714 111.429c0 60.286 24.857 121.715 71.429 160.857c5.143-9.714 14.857-16.286 26-16.286c10 0 19.428 5.714 24.571 14.286c5.429-8.571 14.571-14.286 24.858-14.286c10 0 19.428 5.714 24.571 14.286c5.429-8.571 14.857-14.286 24.858-14.286c10 0 19.428 5.714 24.857 14.286c5.143-8.571 14.571-14.286 24.572-14.286c10.857 0 20.857 6.572 25.714 16c43.427-36.286 68.569-92 71.426-148.286m10.572-99.714c0-53.714-34.571-105.714-92.572-105.714c-30.285 0-58.571 15.143-78.857 36.857C240.862 183.812 233.41 254 302.286 254c28.805 0 97.357-28.538 84.286 36.857c28.857-26 45.714-65.714 45.714-104.571"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThinkPeaks(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m465.4 409.4l87.1-150.2l-32-.3l-55.1 95L259.2 0L23 407.4l32 .3L259.2 55.6zm-355.3-44.1h32.1l117.4-202.5L463 511.9l32.5.1l-235.8-404.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tiktok(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 209.91a210.06 210.06 0 0 1-122.77-39.25v178.72A162.55 162.55 0 1 1 185 188.31v89.89a74.62 74.62 0 1 0 52.23 71.18V0h88a121.18 121.18 0 0 0 1.86 22.17A122.18 122.18 0 0 0 381 102.39a121.43 121.43 0 0 0 67 20.14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TradeFederation(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M248 8.8c-137 0-248 111-248 248s111 248 248 248s248-111 248-248s-111-248-248-248m0 482.8c-129.7 0-234.8-105.1-234.8-234.8S118.3 22 248 22s234.8 105.1 234.8 234.8S377.7 491.6 248 491.6m155.1-328.5v-46.8H209.3V198H54.2l36.7 46h117.7v196.8h48.8V245h83.3v-47h-83.3v-34.8h145.7zm-73.3 45.1v23.9h-82.9v197.4h-26.8V232.1H96.3l-20.1-23.9h143.9v-80.6h171.8V152h-145v56.2zm-161.3-69l-12.4-20.7l2.1 23.8l-23.5 5.4l23.3 5.4l-2.1 24l12.3-20.5l22.2 9.5l-15.7-18.1l15.8-18.1zm-29.6-19.7l9.3-11.5l-12.7 5.9l-8-12.4l1.7 13.9l-14.3 3.8l13.7 2.7l-.8 14.7l6.8-12.2l13.8 5.3zm165.4 145.2l-13.1 5.6l-7.3-12.2l1.3 14.2l-13.9 3.2l13.9 3.2l-1.2 14.2l7.3-12.2l13.1 5.5l-9.4-10.7zm106.9-77.2l-20.9 9.1l-12-19.6l2.2 22.7l-22.3 5.4l22.2 4.9l-1.8 22.9l11.5-19.6l21.2 8.8l-15.1-17zM248 29.9c-125.3 0-226.9 101.6-226.9 226.9S122.7 483.7 248 483.7s226.9-101.6 226.9-226.9S373.3 29.9 248 29.9M342.6 196v51h-83.3v195.7h-52.7V245.9H89.9l-40-49.9h157.4v-81.6h197.8v50.7H259.4V196zM248 43.2c60.3 0 114.8 25 153.6 65.2H202.5V190H45.1C73.1 104.8 153.4 43.2 248 43.2m0 427.1c-117.9 0-213.6-95.6-213.6-213.5c0-21.2 3.1-41.8 8.9-61.1L87.1 252h114.7v196.8h64.6V253h83.3v-62.7h-83.2v-19.2h145.6v-50.8c30.8 37 49.3 84.6 49.3 136.5c.1 117.9-95.5 213.5-213.4 213.5M178.8 275l-11-21.4l1.7 24.5l-23.7 3.9l23.8 5.9l-3.7 23.8l13-20.9l21.5 10.8l-15.8-18.8l16.9-17.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trello(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M392.3 32H56.1C25.1 32 0 57.1 0 88c-.1 0 0-4 0 336c0 30.9 25.1 56 56 56h336.2c30.8-.2 55.7-25.2 55.7-56V88c.1-30.8-24.8-55.8-55.6-56M197 371.3c-.2 14.7-12.1 26.6-26.9 26.6H87.4c-14.8.1-26.9-11.8-27-26.6V117.1c0-14.8 12-26.9 26.9-26.9h82.9c14.8 0 26.9 12 26.9 26.9v254.2zm193.1-112c0 14.8-12 26.9-26.9 26.9h-81c-14.8 0-26.9-12-26.9-26.9V117.2c0-14.8 12-26.9 26.8-26.9h81.1c14.8 0 26.9 12 26.9 26.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tripadvisor(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M528.91 178.82L576 127.58H471.66a326.11 326.11 0 0 0-367 0H0l47.09 51.24a143.911 143.911 0 0 0 194.77 211.91l46.14 50.2l46.11-50.17a143.94 143.94 0 0 0 241.77-105.58h-.03a143.56 143.56 0 0 0-46.94-106.36M144.06 382.57a97.39 97.39 0 1 1 97.39-97.39a97.39 97.39 0 0 1-97.39 97.39M288 282.37c0-64.09-46.62-119.08-108.09-142.59a281 281 0 0 1 216.17 0C334.61 163.3 288 218.29 288 282.37m143.88 100.2h-.01a97.405 97.405 0 1 1 .01 0M144.06 234.12h-.01a51.06 51.06 0 1 0 51.06 51.06v-.11a51 51 0 0 0-51.05-50.95m287.82 0a51.06 51.06 0 1 0 51.06 51.06a51.06 51.06 0 0 0-51.06-51.06"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tumblr(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M309.8 480.3c-13.6 14.5-50 31.7-97.4 31.7c-120.8 0-147-88.8-147-140.6v-144H17.9c-5.5 0-10-4.5-10-10v-68c0-7.2 4.5-13.6 11.3-16c62-21.8 81.5-76 84.3-117.1c.8-11 6.5-16.3 16.1-16.3h70.9c5.5 0 10 4.5 10 10v115.2h83c5.5 0 10 4.4 10 9.9v81.7c0 5.5-4.5 10-10 10h-83.4V360c0 34.2 23.7 53.6 68 35.8c4.8-1.9 9-3.2 12.7-2.2c3.5.9 5.8 3.4 7.4 7.9l22 64.3c1.8 5 3.3 10.6-.4 14.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TumblrSquare(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 32H48C21.5 32 0 53.5 0 80v352c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V80c0-26.5-21.5-48-48-48m-82.3 364.2c-8.5 9.1-31.2 19.8-60.9 19.8c-75.5 0-91.9-55.5-91.9-87.9v-90h-29.7c-3.4 0-6.2-2.8-6.2-6.2v-42.5c0-4.5 2.8-8.5 7.1-10c38.8-13.7 50.9-47.5 52.7-73.2c.5-6.9 4.1-10.2 10-10.2h44.3c3.4 0 6.2 2.8 6.2 6.2v72h51.9c3.4 0 6.2 2.8 6.2 6.2v51.1c0 3.4-2.8 6.2-6.2 6.2h-52.1V321c0 21.4 14.8 33.5 42.5 22.4c3-1.2 5.6-2 8-1.4c2.2.5 3.6 2.1 4.6 4.9l13.8 40.2c1 3.2 2 6.7-.3 9.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitch(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M391.17 103.47h-38.63v109.7h38.63ZM285 103h-38.63v109.75H285ZM120.83 0L24.31 91.42v329.16h115.83V512l96.53-91.42h77.25L487.69 256V0Zm328.24 237.75l-77.22 73.12h-77.24l-67.6 64v-64h-86.87V36.58h308.93Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitter(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M459.37 151.716c.325 4.548.325 9.097.325 13.645c0 138.72-105.583 298.558-298.558 298.558c-59.452 0-114.68-17.219-161.137-47.106c8.447.974 16.568 1.299 25.34 1.299c49.055 0 94.213-16.568 130.274-44.832c-46.132-.975-84.792-31.188-98.112-72.772c6.498.974 12.995 1.624 19.818 1.624c9.421 0 18.843-1.3 27.614-3.573c-48.081-9.747-84.143-51.98-84.143-102.985v-1.299c13.969 7.797 30.214 12.67 47.431 13.319c-28.264-18.843-46.781-51.005-46.781-87.391c0-19.492 5.197-37.36 14.294-52.954c51.655 63.675 129.3 105.258 216.365 109.807c-1.624-7.797-2.599-15.918-2.599-24.04c0-57.828 46.782-104.934 104.934-104.934c30.213 0 57.502 12.67 76.67 33.137c23.715-4.548 46.456-13.32 66.599-25.34c-7.798 24.366-24.366 44.833-46.132 57.827c21.117-2.273 41.584-8.122 60.426-16.243c-14.292 20.791-32.161 39.308-52.628 54.253"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwitterSquare(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 32H48C21.5 32 0 53.5 0 80v352c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V80c0-26.5-21.5-48-48-48m-48.9 158.8c.2 2.8.2 5.7.2 8.5c0 86.7-66 186.6-186.6 186.6c-37.2 0-71.7-10.8-100.7-29.4c5.3.6 10.4.8 15.8.8c30.7 0 58.9-10.4 81.4-28c-28.8-.6-53-19.5-61.3-45.5c10.1 1.5 19.2 1.5 29.6-1.2c-30-6.1-52.5-32.5-52.5-64.4v-.8c8.7 4.9 18.9 7.9 29.6 8.3a65.447 65.447 0 0 1-29.2-54.6c0-12.2 3.2-23.4 8.9-33.1c32.3 39.8 80.8 65.8 135.2 68.6c-9.3-44.5 24-80.6 64-80.6c18.9 0 35.9 7.9 47.9 20.7c14.8-2.8 29-8.3 41.6-15.8c-4.9 15.2-15.2 28-28.8 36.1c13.2-1.4 26-5.1 37.8-10.2c-8.9 13.1-20.1 24.7-32.9 34"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TypoThree(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M178.7 78.4c0-24.7 5.4-32.4 13.9-39.4c-69.5 8.5-149.3 34-176.3 66.4c-5.4 7.7-9.3 20.8-9.3 37.1C7 246 113.8 480 191.1 480c36.3 0 97.3-59.5 146.7-139c-7 2.3-11.6 2.3-18.5 2.3c-57.2 0-140.6-198.5-140.6-264.9M301.5 32c-30.1 0-41.7 5.4-41.7 36.3c0 66.4 53.8 198.5 101.7 198.5c26.3 0 78.8-99.7 78.8-182.3c0-40.9-67-52.5-138.8-52.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Uber(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M414.1 32H33.9C15.2 32 0 47.2 0 65.9V446c0 18.8 15.2 34 33.9 34H414c18.7 0 33.9-15.2 33.9-33.9V65.9C448 47.2 432.8 32 414.1 32M237.6 391.1C163 398.6 96.4 344.2 88.9 269.6h94.4V290c0 3.7 3 6.8 6.8 6.8H258c3.7 0 6.8-3 6.8-6.8v-67.9c0-3.7-3-6.8-6.8-6.8h-67.9c-3.7 0-6.8 3-6.8 6.8v20.4H88.9c7-69.4 65.4-122.2 135.1-122.2c69.7 0 128.1 52.8 135.1 122.2c7.5 74.5-46.9 141.1-121.5 148.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ubuntu(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M248 8C111 8 0 119 0 256s111 248 248 248s248-111 248-248S385 8 248 8m52.7 93c8.8-15.2 28.3-20.5 43.5-11.7c15.3 8.8 20.5 28.3 11.7 43.6c-8.8 15.2-28.3 20.5-43.5 11.7c-15.3-8.9-20.5-28.4-11.7-43.6M87.4 287.9c-17.6 0-31.9-14.3-31.9-31.9c0-17.6 14.3-31.9 31.9-31.9c17.6 0 31.9 14.3 31.9 31.9c0 17.6-14.3 31.9-31.9 31.9m28.1 3.1c22.3-17.9 22.4-51.9 0-69.9c8.6-32.8 29.1-60.7 56.5-79.1l23.7 39.6c-51.5 36.3-51.5 112.5 0 148.8L172 370c-27.4-18.3-47.8-46.3-56.5-79m228.7 131.7c-15.3 8.8-34.7 3.6-43.5-11.7c-8.8-15.3-3.6-34.8 11.7-43.6c15.2-8.8 34.7-3.6 43.5 11.7c8.8 15.3 3.6 34.8-11.7 43.6m.3-69.5c-26.7-10.3-56.1 6.6-60.5 35c-5.2 1.4-48.9 14.3-96.7-9.4l22.5-40.3c57 26.5 123.4-11.7 128.9-74.4l46.1.7c-2.3 34.5-17.3 65.5-40.3 88.4m-5.9-105.3c-5.4-62-71.3-101.2-128.9-74.4l-22.5-40.3c47.9-23.7 91.5-10.8 96.7-9.4c4.4 28.3 33.8 45.3 60.5 35c23.1 22.9 38 53.9 40.2 88.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Uikit(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M443.9 128v256L218 512L0 384V169.7l87.6 45.1v117l133.5 75.5l135.8-75.5v-151l-101.1-57.6l87.6-53.1zM308.6 49.1L223.8 0l-88.6 54.8l86 47.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Umbraco(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M255.35 8C118.36 7.83 7.14 118.72 7 255.68c-.07 137 111 248.2 248 248.27c136.85 0 247.82-110.7 248-247.67S392.34 8.17 255.35 8m145 266q-1.14 40.68-14 65t-43.51 35q-30.61 10.7-85.45 10.47h-4.6q-54.78.22-85.44-10.47t-43.52-35q-12.85-24.36-14-65a224.81 224.81 0 0 1 0-30.71a418.37 418.37 0 0 1 3.6-43.88c1.88-13.39 3.57-22.58 5.4-32c1-4.88 1.28-6.42 1.82-8.45a5.09 5.09 0 0 1 4.9-3.89h.69l32 5a5.07 5.07 0 0 1 4.16 5a5 5 0 0 1 0 .77l-1.7 8.78q-2.41 13.25-4.84 33.68a380.62 380.62 0 0 0-2.64 42.15q-.28 40.43 8.13 59.83a43.87 43.87 0 0 0 31.31 25.18A243 243 0 0 0 250 340.6h10.25a242.64 242.64 0 0 0 57.27-5.16a43.86 43.86 0 0 0 31.15-25.23q8.53-19.42 8.13-59.78a388 388 0 0 0-2.6-42.15q-2.48-20.38-4.89-33.68l-1.69-8.78a5 5 0 0 1 0-.77a5 5 0 0 1 4.2-5l32-5h.82a5 5 0 0 1 4.9 3.89c.55 2.05.81 3.57 1.83 8.45c1.82 9.62 3.52 18.78 5.39 32a415.71 415.71 0 0 1 3.61 43.88a228.06 228.06 0 0 1-.04 30.73z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Uncharted(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M171.73 232.813a5.381 5.381 0 0 0 4.97-3.313a48.081 48.081 0 0 1 14.9-25.256c1.243-.828 1.657-2.484 1.657-4.141a4.22 4.22 0 0 0-2.071-3.312L74.429 128.473L148.958 85a9.941 9.941 0 0 0 4.968-8.281a9.108 9.108 0 0 0-4.968-8.281L126.6 55.6a9.748 9.748 0 0 0-9.523 0l-100.2 57.966a9.943 9.943 0 0 0-4.969 8.281v115.107a9.109 9.109 0 0 0 4.969 8.281l22.358 12.835a8.829 8.829 0 0 0 4.968 1.242a9.4 9.4 0 0 0 6.625-2.484a10.8 10.8 0 0 0 2.9-7.039V164.5l115.932 67.9a4.5 4.5 0 0 0 2.07.413M323.272 377.73a12.478 12.478 0 0 0-4.969 1.242l-74.528 43.062V287.882c0-2.9-2.9-5.8-6.211-4.555a53.036 53.036 0 0 1-28.984.414a4.86 4.86 0 0 0-6.21 4.555v133.323l-74.529-43.061a8.83 8.83 0 0 0-4.969-1.242a9.631 9.631 0 0 0-9.523 9.523v26.085a9.107 9.107 0 0 0 4.969 8.281l100.2 57.553a8.829 8.829 0 0 0 4.968 1.242a11.027 11.027 0 0 0 4.969-1.242l100.2-57.553a9.941 9.941 0 0 0 4.968-8.281v-26.085c-.823-4.554-5.383-9.109-10.351-9.109M286.007 78a23 23 0 1 0-23-23a23 23 0 0 0 23 23m63.627-10.086a23 23 0 1 0 23 23a23 23 0 0 0-23-23m63.182 83.686a23 23 0 1 0-23-23a23 23 0 0 0 23 23m-63.182-9.2a23 23 0 1 0 23 23a23 23 0 0 0-23-23m-63.627 83.244a23 23 0 1 0-23-23a23 23 0 0 0 23 23.004Zm-62.074 36.358a23 23 0 1 0-23-23a23 23 0 0 0 23 23.004Zm188.883-82.358a23 23 0 1 0 23 23a23 23 0 0 0-23-22.996Zm0 72.272a23 23 0 1 0 23 23a23 23 0 0 0-23-22.996Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Uniregistry(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 480c39.5 0 76.2-11.8 106.8-32.2H85.3C115.8 468.2 152.5 480 192 480m-89.1-193.1v-12.4H0v12.4c0 2.5 0 5 .1 7.4h103.1c-.2-2.4-.3-4.9-.3-7.4m20.5 57H8.5c2.6 8.5 5.8 16.8 9.6 24.8h138.3c-12.9-5.7-24.1-14.2-33-24.8m-17.7-34.7H1.3c.9 7.6 2.2 15 3.9 22.3h109.7c-4-6.9-7.2-14.4-9.2-22.3m-2.8-69.3H0v17.3h102.9zm0-173.2H0v4.9h102.9zm0-34.7H0v2.5h102.9zm0 69.3H0v7.4h102.9zm0 104H0v14.8h102.9zm0-69.3H0v9.9h102.9zm0 34.6H0V183h102.9zm166.2 160.9h109.7c1.8-7.3 3.1-14.7 3.9-22.3H278.3c-2.1 7.9-5.2 15.4-9.2 22.3m12-185.7H384V136H281.1zm0 37.2H384v-12.4H281.1zm0-74.3H384v-7.4H281.1zm0-76.7v2.5H384V32zm-203 410.9h227.7c11.8-8.7 22.7-18.6 32.2-29.7H44.9c9.6 11 21.4 21 33.2 29.7m203-371.3H384v-4.9H281.1zm0 148.5H384v-14.8H281.1zM38.8 405.7h305.3c6.7-8.5 12.6-17.6 17.8-27.2H23c5.2 9.6 9.2 18.7 15.8 27.2m188.8-37.1H367c3.7-8 5.8-16.2 8.5-24.8h-115c-8.8 10.7-20.1 19.2-32.9 24.8m53.5-81.7c0 2.5-.1 5-.4 7.4h103.1c.1-2.5.2-4.9.2-7.4v-12.4H281.1zm0-29.7H384v-17.3H281.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unity(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m498.11 206.4l-52.8-191.68L248.2 66.08L219 116.14l-59.2-.43L15.54 256l144.28 140.32l59.17-.43l29.24 50l197.08 51.36l52.8-191.62l-30-49.63Zm-274.34-82.2l150.78-37.69L288 232.33H114.87Zm0 263.63l-108.9-108.12H288l86.55 145.81Zm193 14L330.17 256l86.58-145.84L458.56 256Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unsplash(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 230.17V480H0V230.17h141.13v124.92h165.74V230.17ZM306.87 32H141.13v124.91h165.74Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Untappd(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M401.3 49.9c-79.8 160.1-84.6 152.5-87.9 173.2l-5.2 32.8c-1.9 12-6.6 23.5-13.7 33.4L145.6 497.1c-7.6 10.6-20.4 16.2-33.4 14.6c-40.3-5-77.8-32.2-95.3-68.5c-5.7-11.8-4.5-25.8 3.1-36.4l148.9-207.9c7.1-9.9 16.4-18 27.2-23.7l29.3-15.5c18.5-9.8 9.7-11.9 135.6-138.9c1-4.8 1-7.3 3.6-8c3-.7 6.6-1 6.3-4.6l-.4-4.6c-.2-1.9 1.3-3.6 3.2-3.6c4.5-.1 13.2 1.2 25.6 10c12.3 8.9 16.4 16.8 17.7 21.1c.6 1.8-.6 3.7-2.4 4.2l-4.5 1.1c-3.4.9-2.5 4.4-2.3 7.4c.1 2.8-2.3 3.6-6.5 6.1M230.1 36.4c3.4.9 2.5 4.4 2.3 7.4c-.2 2.7 2.1 3.5 6.4 6c7.9 15.9 15.3 30.5 22.2 44c.7 1.3 2.3 1.5 3.3.5c11.2-12 24.6-26.2 40.5-42.6c1.3-1.4 1.4-3.5.1-4.9c-8-8.2-16.5-16.9-25.6-26.1c-1-4.7-1-7.3-3.6-8c-3-.8-6.6-1-6.3-4.6c.3-3.3 1.4-8.1-2.8-8.2c-4.5-.1-13.2 1.1-25.6 10c-12.3 8.9-16.4 16.8-17.7 21.1c-1.4 4.2 3.6 4.6 6.8 5.4M620 406.7L471.2 198.8c-13.2-18.5-26.6-23.4-56.4-39.1c-11.2-5.9-14.2-10.9-30.5-28.9c-1-1.1-2.9-.9-3.6.5c-46.3 88.8-47.1 82.8-49 94.8c-1.7 10.7-1.3 20 .3 29.8c1.9 12 6.6 23.5 13.7 33.4l148.9 207.9c7.6 10.6 20.2 16.2 33.1 14.7c40.3-4.9 78-32 95.7-68.6c5.4-11.9 4.3-25.9-3.4-36.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ups(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M103.2 303c-5.2 3.6-32.6 13.1-32.6-19V180H37.9v102.6c0 74.9 80.2 51.1 97.9 39V180h-32.6zM4 74.82v220.9c0 103.7 74.9 135.2 187.7 184.1c112.4-48.9 187.7-80.2 187.7-184.1V74.82c-116.3-61.6-281.8-49.6-375.4 0m358.1 220.9c0 86.6-53.2 113.6-170.4 165.3c-117.5-51.8-170.5-78.7-170.5-165.3v-126.4c102.3-93.8 231.6-100 340.9-89.8zm-209.6-107.4v212.8h32.7v-68.7c24.4 7.3 71.7-2.6 71.7-78.5c0-97.4-80.7-80.92-104.4-65.6m32.7 117.3v-100.3c8.4-4.2 38.4-12.7 38.4 49.3c0 67.9-36.4 51.8-38.4 51m79.1-86.4c.1 47.3 51.6 42.5 52.2 70.4c.6 23.5-30.4 23-50.8 4.9v30.1c36.2 21.5 81.9 8.1 83.2-33.5c1.7-51.5-54.1-46.6-53.4-73.2c.6-20.3 30.6-20.5 48.5-2.2v-28.4c-28.5-22-79.9-9.2-79.7 31.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Usb(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M641.5 256c0 3.1-1.7 6.1-4.5 7.5L547.9 317c-1.4.8-2.8 1.4-4.5 1.4c-1.4 0-3.1-.3-4.5-1.1c-2.8-1.7-4.5-4.5-4.5-7.8v-35.6H295.7c25.3 39.6 40.5 106.9 69.6 106.9H392V354c0-5 3.9-8.9 8.9-8.9H490c5 0 8.9 3.9 8.9 8.9v89.1c0 5-3.9 8.9-8.9 8.9h-89.1c-5 0-8.9-3.9-8.9-8.9v-26.7h-26.7c-75.4 0-81.1-142.5-124.7-142.5H140.3c-8.1 30.6-35.9 53.5-69 53.5C32 327.3 0 295.3 0 256s32-71.3 71.3-71.3c33.1 0 61 22.8 69 53.5c39.1 0 43.9 9.5 74.6-60.4C255 88.7 273 95.7 323.8 95.7c7.5-20.9 27-35.6 50.4-35.6c29.5 0 53.5 23.9 53.5 53.5s-23.9 53.5-53.5 53.5c-23.4 0-42.9-14.8-50.4-35.6H294c-29.1 0-44.3 67.4-69.6 106.9h310.1v-35.6c0-3.3 1.7-6.1 4.5-7.8c2.8-1.7 6.4-1.4 8.9.3l89.1 53.5c2.8 1.1 4.5 4.1 4.5 7.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Usps(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M460.3 241.7c25.8-41.3 15.2-48.8-11.7-48.8h-27c-.1 0-1.5-1.4-10.9 8c-11.2 5.6-37.9 6.3-37.9 8.7c0 4.5 70.3-3.1 88.1 0c9.5 1.5-1.5 20.4-4.4 32c-.5 4.5 2.4 2.3 3.8.1m-112.1 22.6c64-21.3 97.3-23.9 102-26.2c4.4-2.9-4.4-6.6-26.2-5.8c-51.7 2.2-137.6 37.1-172.6 53.9l-30.7-93.3h196.6c-2.7-28.2-152.9-22.6-337.9-22.6L27 415.8c196.4-97.3 258.9-130.3 321.2-151.5M94.7 96c253.3 53.7 330 65.7 332.1 85.2c36.4 0 45.9 0 52.4 6.6c21.1 19.7-14.6 67.7-14.6 67.7c-4.4 2.9-406.4 160.2-406.4 160.2h423.1L549 96z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ussunnah(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m156.8 285.1l5.7 14.4h-8.2c-1.3-3.2-3.1-7.7-3.8-9.5c-2.5-6.3-1.1-8.4 0-10c1.9-2.7 3.2-4.4 3.6-5.2c0 2.2.8 5.7 2.7 10.3m297.3 18.8c-2.1 13.8-5.7 27.1-10.5 39.7l43 23.4l-44.8-18.8c-5.3 13.2-12 25.6-19.9 37.2l34.2 30.2l-36.8-26.4c-8.4 11.8-18 22.6-28.7 32.3l24.9 34.7l-28.1-31.8c-11 9.6-23.1 18-36.1 25.1l15.7 37.2l-19.3-35.3c-13.1 6.8-27 12.1-41.6 15.9l6.7 38.4l-10.5-37.4c-14.3 3.4-29.2 5.3-44.5 5.4L256 512l-1.9-38.4c-15.3-.1-30.2-2-44.5-5.3L199 505.6l6.7-38.2c-14.6-3.7-28.6-9.1-41.7-15.8l-19.2 35.1l15.6-37c-13-7-25.2-15.4-36.2-25.1l-27.9 31.6l24.7-34.4c-10.7-9.7-20.4-20.5-28.8-32.3l-36.5 26.2l33.9-29.9c-7.9-11.6-14.6-24.1-20-37.3l-44.4 18.7L67.8 344c-4.8-12.7-8.4-26.1-10.5-39.9l-51 9l50.3-14.2c-1.1-8.5-1.7-17.1-1.7-25.9c0-4.7.2-9.4.5-14.1L0 256l56-2.8c1.3-13.1 3.8-25.8 7.5-38.1L6.4 199l58.9 10.4c4-12 9.1-23.5 15.2-34.4l-55.1-30l58.3 24.6C90 159 97.2 149.2 105.3 140L55.8 96.4l53.9 38.7c8.1-8.6 17-16.5 26.6-23.6l-40-55.6l45.6 51.6c9.5-6.6 19.7-12.3 30.3-17.2l-27.3-64.9l33.8 62.1c10.5-4.4 21.4-7.9 32.7-10.4L199 6.4l19.5 69.2c11-2.1 22.3-3.2 33.8-3.4L256 0l3.6 72.2c11.5.2 22.8 1.4 33.8 3.5L313 6.4l-12.4 70.7c11.3 2.6 22.2 6.1 32.6 10.5l33.9-62.2l-27.4 65.1c10.6 4.9 20.7 10.7 30.2 17.2l45.8-51.8l-40.1 55.9c9.5 7.1 18.4 15 26.5 23.6l54.2-38.9l-49.7 43.9c8 9.1 15.2 18.9 21.5 29.4l58.7-24.7l-55.5 30.2c6.1 10.9 11.1 22.3 15.1 34.3l59.3-10.4l-57.5 16.2c3.7 12.2 6.2 24.9 7.5 37.9L512 256l-56 2.8c.3 4.6.5 9.3.5 14.1c0 8.7-.6 17.3-1.6 25.8l50.7 14.3zm-21.8-31c0-97.5-79-176.5-176.5-176.5s-176.5 79-176.5 176.5s79 176.5 176.5 176.5s176.5-79 176.5-176.5m-24 0c0 84.3-68.3 152.6-152.6 152.6s-152.6-68.3-152.6-152.6s68.3-152.6 152.6-152.6s152.6 68.3 152.6 152.6M195 241c0 2.1 1.3 3.8 3.6 5.1c3.3 1.9 6.2 4.6 8.2 8.2c2.8-5.7 4.3-9.5 4.3-11.2c0-2.2-1.1-4.4-3.2-7c-2.1-2.5-3.2-5.2-3.3-7.7c-6.5 6.8-9.6 10.9-9.6 12.6m-40.7-19c0 2.1 1.3 3.8 3.6 5.1c3.5 1.9 6.2 4.6 8.2 8.2c2.8-5.7 4.3-9.5 4.3-11.2c0-2.2-1.1-4.4-3.2-7c-2.1-2.5-3.2-5.2-3.3-7.7c-6.5 6.8-9.6 10.9-9.6 12.6m-19 0c0 2.1 1.3 3.8 3.6 5.1c3.3 1.9 6.2 4.6 8.2 8.2c2.8-5.7 4.3-9.5 4.3-11.2c0-2.2-1.1-4.4-3.2-7c-2.1-2.5-3.2-5.2-3.3-7.7c-6.4 6.8-9.6 10.9-9.6 12.6m204.9 87.9c-8.4-3-8.7-6.8-8.7-15.6V182c-8.2 12.5-14.2 18.6-18 18.6c6.3 14.4 9.5 23.9 9.5 28.3v64.3c0 2.2-2.2 6.5-4.7 6.5h-18c-2.8-7.5-10.2-26.9-15.3-40.3c-2 2.5-7.2 9.2-10.7 13.7c2.4 1.6 4.1 3.6 5.2 6.3c2.6 6.7 6.4 16.5 7.9 20.2h-9.2c-3.9-10.4-9.6-25.4-11.8-31.1c-2 2.5-7.2 9.2-10.7 13.7c2.4 1.6 4.1 3.6 5.2 6.3c.8 2 2.8 7.3 4.3 10.9H256c-1.5-4.1-5.6-14.6-8.4-22c-2 2.5-7.2 9.2-10.7 13.7c2.5 1.6 4.3 3.6 5.2 6.3c.2.6.5 1.4.6 1.7H225c-4.6-13.9-11.4-27.7-11.4-34.1c0-2.2.3-5.1 1.1-8.2c-8.8 10.8-14 15.9-14 25c0 7.5 10.4 28.3 10.4 33.3c0 1.7-.5 3.3-1.4 4.9c-9.6-12.7-15.5-20.7-18.8-20.7h-12l-11.2-28c-3.8-9.6-5.7-16-5.7-18.8c0-3.8.5-7.7 1.7-12.2c-1 1.3-3.7 4.7-5.5 7.1c-.8-2.1-3.1-7.7-4.6-11.5c-2.1 2.5-7.5 9.1-11.2 13.6c.9 2.3 3.3 8.1 4.9 12.2c-2.5 3.3-9.1 11.8-13.6 17.7c-4 5.3-5.8 13.3-2.7 21.8c2.5 6.7 2 7.9-1.7 14.1H191c5.5 0 14.3 14 15.5 22c13.2-16 15.4-19.6 16.8-21.6h107c3.9 0 7.2-1.9 9.9-5.8m20.1-26.6V181.7c-9 12.5-15.9 18.6-20.7 18.6c7.1 14.4 10.7 23.9 10.7 28.3v66.3c0 17.5 8.6 20.4 24 20.4c8.1 0 12.5-.8 13.7-2.7c-4.3-1.6-7.6-2.5-9.9-3.3c-8.1-3.2-17.8-7.4-17.8-26"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vaadin(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M224.5 140.7c1.5-17.6 4.9-52.7 49.8-52.7h98.6c20.7 0 32.1-7.8 32.1-21.6V54.1c0-12.2 9.3-22.1 21.5-22.1S448 41.9 448 54.1v36.5c0 42.9-21.5 62-66.8 62H280.7c-30.1 0-33 14.7-33 27.1c0 1.3-.1 2.5-.2 3.7c-.7 12.3-10.9 22.2-23.4 22.2s-22.7-9.8-23.4-22.2c-.1-1.2-.2-2.4-.2-3.7c0-12.3-3-27.1-33-27.1H66.8c-45.3 0-66.8-19.1-66.8-62V54.1C0 41.9 9.4 32 21.6 32s21.5 9.9 21.5 22.1v12.3C43.1 80.2 54.5 88 75.2 88h98.6c44.8 0 48.3 35.1 49.8 52.7zM224 456c11.5 0 21.4-7 25.7-16.3c1.1-1.8 97.1-169.6 98.2-171.4c11.9-19.6-3.2-44.3-27.2-44.3c-13.9 0-23.3 6.4-29.8 20.3L224 362l-66.9-117.7c-6.4-13.9-15.9-20.3-29.8-20.3c-24 0-39.1 24.6-27.2 44.3c1.1 1.9 97.1 169.6 98.2 171.4c4.3 9.3 14.2 16.3 25.7 16.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Viacoin(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 32h-64l-80.7 192h-94.5L64 32H0l48 112H0v48h68.5l13.8 32H0v48h102.8L192 480l89.2-208H384v-48h-82.3l13.8-32H384v-48h-48zM192 336l-27-64h54z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Viadeo(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M276.2 150.5v.7C258.3 98.6 233.6 47.8 205.4 0c43.3 29.2 67 100 70.8 150.5m32.7 121.7c7.6 18.2 11 37.5 11 57c0 77.7-57.8 141-137.8 139.4l3.8-.3c74.2-46.7 109.3-118.6 109.3-205.1c0-38.1-6.5-75.9-18.9-112c1 11.7 1 23.7 1 35.4c0 91.8-18.1 241.6-116.6 280C95 455.2 49.4 398 49.4 329.2c0-75.6 57.4-142.3 135.4-142.3c16.8 0 33.7 3.1 49.1 9.6c1.7-15.1 6.5-29.9 13.4-43.3c-19.9-7.2-41.2-10.7-62.5-10.7c-161.5 0-238.7 195.9-129.9 313.7c67.9 74.6 192 73.9 259.8 0c56.6-61.3 60.9-142.4 36.4-201c-12.7 8-27.1 13.9-42.2 17M418.1 11.7c-31 66.5-81.3 47.2-115.8 80.1c-12.4 12-20.6 34-20.6 50.5c0 14.1 4.5 27.1 12 38.8c47.4-11 98.3-46 118.2-90.7c-.7 5.5-4.8 14.4-7.2 19.2c-20.3 35.7-64.6 65.6-99.7 84.9c14.8 14.4 33.7 25.8 55 25.8c79 0 110.1-134.6 58.1-208.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViadeoSquare(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 32H48C21.5 32 0 53.5 0 80v352c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V80c0-26.5-21.5-48-48-48M280.7 381.2c-42.4 46.2-120 46.6-162.4 0c-68-73.6-19.8-196.1 81.2-196.1c13.3 0 26.6 2.1 39.1 6.7c-4.3 8.4-7.3 17.6-8.4 27.1c-9.7-4.1-20.2-6-30.7-6c-48.8 0-84.6 41.7-84.6 88.9c0 43 28.5 78.7 69.5 85.9c61.5-24 72.9-117.6 72.9-175c0-7.3 0-14.8-.6-22.1c-11.2-32.9-26.6-64.6-44.2-94.5c27.1 18.3 41.9 62.5 44.2 94.1v.4c7.7 22.5 11.8 46.2 11.8 70c0 54.1-21.9 99-68.3 128.2l-2.4.2c50 1 86.2-38.6 86.2-87.2c0-12.2-2.1-24.3-6.9-35.7c9.5-1.9 18.5-5.6 26.4-10.5c15.3 36.6 12.6 87.3-22.8 125.6M309 233.7c-13.3 0-25.1-7.1-34.4-16.1c21.9-12 49.6-30.7 62.3-53c1.5-3 4.1-8.6 4.5-12c-12.5 27.9-44.2 49.8-73.9 56.7c-4.7-7.3-7.5-15.5-7.5-24.3c0-10.3 5.2-24.1 12.9-31.6c21.6-20.5 53-8.5 72.4-50c32.5 46.2 13.1 130.3-36.3 130.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Viber(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M444 49.9C431.3 38.2 379.9.9 265.3.4c0 0-135.1-8.1-200.9 52.3C27.8 89.3 14.9 143 13.5 209.5c-1.4 66.5-3.1 191.1 117 224.9h.1l-.1 51.6s-.8 20.9 13 25.1c16.6 5.2 26.4-10.7 42.3-27.8c8.7-9.4 20.7-23.2 29.8-33.7c82.2 6.9 145.3-8.9 152.5-11.2c16.6-5.4 110.5-17.4 125.7-142c15.8-128.6-7.6-209.8-49.8-246.5M457.9 287c-12.9 104-89 110.6-103 115.1c-6 1.9-61.5 15.7-131.2 11.2c0 0-52 62.7-68.2 79c-5.3 5.3-11.1 4.8-11-5.7c0-6.9.4-85.7.4-85.7c-.1 0-.1 0 0 0c-101.8-28.2-95.8-134.3-94.7-189.8c1.1-55.5 11.6-101 42.6-131.6c55.7-50.5 170.4-43 170.4-43c96.9.4 143.3 29.6 154.1 39.4c35.7 30.6 53.9 103.8 40.6 211.1m-139-80.8c.4 8.6-12.5 9.2-12.9.6c-1.1-22-11.4-32.7-32.6-33.9c-8.6-.5-7.8-13.4.7-12.9c27.9 1.5 43.4 17.5 44.8 46.2m20.3 11.3c1-42.4-25.5-75.6-75.8-79.3c-8.5-.6-7.6-13.5.9-12.9c58 4.2 88.9 44.1 87.8 92.5c-.1 8.6-13.1 8.2-12.9-.3m47 13.4c.1 8.6-12.9 8.7-12.9.1c-.6-81.5-54.9-125.9-120.8-126.4c-8.5-.1-8.5-12.9 0-12.9c73.7.5 133 51.4 133.7 139.2M374.9 329v.2c-10.8 19-31 40-51.8 33.3l-.2-.3c-21.1-5.9-70.8-31.5-102.2-56.5c-16.2-12.8-31-27.9-42.4-42.4c-10.3-12.9-20.7-28.2-30.8-46.6c-21.3-38.5-26-55.7-26-55.7c-6.7-20.8 14.2-41 33.3-51.8h.2c9.2-4.8 18-3.2 23.9 3.9c0 0 12.4 14.8 17.7 22.1c5 6.8 11.7 17.7 15.2 23.8c6.1 10.9 2.3 22-3.7 26.6l-12 9.6c-6.1 4.9-5.3 14-5.3 14s17.8 67.3 84.3 84.3c0 0 9.1.8 14-5.3l9.6-12c4.6-6 15.7-9.8 26.6-3.7c14.7 8.3 33.4 21.2 45.8 32.9c7 5.7 8.6 14.4 3.8 23.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vimeo(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M403.2 32H44.8C20.1 32 0 52.1 0 76.8v358.4C0 459.9 20.1 480 44.8 480h358.4c24.7 0 44.8-20.1 44.8-44.8V76.8c0-24.7-20.1-44.8-44.8-44.8M377 180.8c-1.4 31.5-23.4 74.7-66 129.4c-44 57.2-81.3 85.8-111.7 85.8c-18.9 0-34.8-17.4-47.9-52.3c-25.5-93.3-36.4-148-57.4-148c-2.4 0-10.9 5.1-25.4 15.2l-15.2-19.6c37.3-32.8 72.9-69.2 95.2-71.2c25.2-2.4 40.7 14.8 46.5 51.7c20.7 131.2 29.9 151 67.6 91.6c13.5-21.4 20.8-37.7 21.8-48.9c3.5-33.2-25.9-30.9-45.8-22.4c15.9-52.1 46.3-77.4 91.2-76c33.3.9 49 22.5 47.1 64.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VimeoSquare(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 32H48C21.5 32 0 53.5 0 80v352c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V80c0-26.5-21.5-48-48-48m-16.2 149.6c-1.4 31.1-23.2 73.8-65.3 127.9c-43.5 56.5-80.3 84.8-110.4 84.8c-18.7 0-34.4-17.2-47.3-51.6c-25.2-92.3-35.9-146.4-56.7-146.4c-2.4 0-10.8 5-25.1 15.1L64 192c36.9-32.4 72.1-68.4 94.1-70.4c24.9-2.4 40.2 14.6 46 51.1c20.5 129.6 29.6 149.2 66.8 90.5c13.4-21.2 20.6-37.2 21.5-48.3c3.4-32.8-25.6-30.6-45.2-22.2c15.7-51.5 45.8-76.5 90.1-75.1c32.9 1 48.4 22.4 46.5 64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VimeoV(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M447.8 153.6c-2 43.6-32.4 103.3-91.4 179.1c-60.9 79.2-112.4 118.8-154.6 118.8c-26.1 0-48.2-24.1-66.3-72.3C100.3 250 85.3 174.3 56.2 174.3c-3.4 0-15.1 7.1-35.2 21.1L0 168.2c51.6-45.3 100.9-95.7 131.8-98.5c34.9-3.4 56.3 20.5 64.4 71.5c28.7 181.5 41.4 208.9 93.6 126.7c18.7-29.6 28.8-52.1 30.2-67.6c4.8-45.9-35.8-42.8-63.3-31c22-72.1 64.1-107.1 126.2-105.1c45.8 1.2 67.5 31.1 64.9 89.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vine(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 254.7v52.1c-18.4 4.2-36.9 6.1-52.1 6.1c-36.9 77.4-103 143.8-125.1 156.2c-14 7.9-27.1 8.4-42.7-.8C137 452 34.2 367.7 0 102.7h74.5C93.2 261.8 139 343.4 189.3 404.5c27.9-27.9 54.8-65.1 75.6-106.9c-49.8-25.3-80.1-80.9-80.1-145.6c0-65.6 37.7-115.1 102.2-115.1c114.9 0 106.2 127.9 81.6 181.5c0 0-46.4 9.2-63.5-20.5c3.4-11.3 8.2-30.8 8.2-48.5c0-31.3-11.3-46.6-28.4-46.6c-18.2 0-30.8 17.1-30.8 50c.1 79.2 59.4 118.7 129.9 101.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vk(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M545 117.7c3.7-12.5 0-21.7-17.8-21.7h-58.9c-15 0-21.9 7.9-25.6 16.7c0 0-30 73.1-72.4 120.5c-13.7 13.7-20 18.1-27.5 18.1c-3.7 0-9.4-4.4-9.4-16.9V117.7c0-15-4.2-21.7-16.6-21.7h-92.6c-9.4 0-15 7-15 13.5c0 14.2 21.2 17.5 23.4 57.5v86.8c0 19-3.4 22.5-10.9 22.5c-20 0-68.6-73.4-97.4-157.4c-5.8-16.3-11.5-22.9-26.6-22.9H38.8c-16.8 0-20.2 7.9-20.2 16.7c0 15.6 20 93.1 93.1 195.5C160.4 378.1 229 416 291.4 416c37.5 0 42.1-8.4 42.1-22.9c0-66.8-3.4-73.1 15.4-73.1c8.7 0 23.7 4.4 58.7 38.1c40 40 46.6 57.9 69 57.9h58.9c16.8 0 25.3-8.4 20.4-25c-11.2-34.9-86.9-106.7-90.3-111.5c-8.7-11.2-6.2-16.2 0-26.2c.1-.1 72-101.3 79.4-135.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vnv(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M104.9 352c-34.1 0-46.4-30.4-46.4-30.4L2.6 210.1S-7.8 192 13 192h32.8c10.4 0 13.2 8.7 18.8 18.1l36.7 74.5s5.2 13.1 21.1 13.1s21.1-13.1 21.1-13.1l36.7-74.5c5.6-9.5 8.4-18.1 18.8-18.1h32.8c20.8 0 10.4 18.1 10.4 18.1l-55.8 111.5S174.2 352 140 352zm395 0c-34.1 0-46.4-30.4-46.4-30.4l-55.9-111.5S387.2 192 408 192h32.8c10.4 0 13.2 8.7 18.8 18.1l36.7 74.5s5.2 13.1 21.1 13.1s21.1-13.1 21.1-13.1l36.8-74.5c5.6-9.5 8.4-18.1 18.8-18.1H627c20.8 0 10.4 18.1 10.4 18.1l-55.9 111.5S569.3 352 535.1 352zM337.6 192c34.1 0 46.4 30.4 46.4 30.4l55.9 111.5s10.4 18.1-10.4 18.1h-32.8c-10.4 0-13.2-8.7-18.8-18.1l-36.7-74.5s-5.2-13.1-21.1-13.1c-15.9 0-21.1 13.1-21.1 13.1l-36.7 74.5c-5.6 9.4-8.4 18.1-18.8 18.1h-32.9c-20.8 0-10.4-18.1-10.4-18.1l55.9-111.5s12.2-30.4 46.4-30.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vuejs(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M356.9 64.3H280l-56 88.6l-48-88.6H0L224 448L448 64.3zm-301.2 32h53.8L224 294.5L338.4 96.3h53.8L224 384.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WatchmanMonitoring(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 16C123.452 16 16 123.452 16 256s107.452 240 240 240s240-107.452 240-240S388.548 16 256 16M121.69 429.122C70.056 388.972 36.741 326.322 36.741 256a218.519 218.519 0 0 1 9.587-64.122l102.9-17.895l-.121 10.967l-13.943 2.013s-.144 12.5-.144 19.549a12.778 12.778 0 0 0 4.887 10.349l9.468 7.4Zm105.692-283.27l8.48-7.618s6.934-5.38-.143-9.344c-7.188-4.024-39.53-34.5-39.53-34.5c-5.348-5.477-8.257-7.347-15.46 0c0 0-32.342 30.474-39.529 34.5c-7.078 3.964-.144 9.344-.144 9.344l8.481 7.618l-.048 4.369l-73.507-19.176c39.644-56.938 105.532-94.3 180.018-94.3a218.754 218.754 0 0 1 164.934 75.025l-193.512 37.7Zm34.063 329.269l-33.9-250.857l9.467-7.4a12.778 12.778 0 0 0 4.888-10.349c0-7.044-.144-19.549-.144-19.549l-13.943-2.013l-.116-10.474l241.711 31.391a218.872 218.872 0 0 1 5.851 50.13c0 119.074-95.428 216.212-213.814 219.121"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Waze(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M502.17 201.67C516.69 287.53 471.23 369.59 389 409.8c13 34.1-12.4 70.2-48.32 70.2a51.68 51.68 0 0 1-51.57-49c-6.44.19-64.2 0-76.33-.64A51.69 51.69 0 0 1 159 479.92c-33.86-1.36-57.95-34.84-47-67.92c-37.21-13.11-72.54-34.87-99.62-70.8c-13-17.28-.48-41.8 20.84-41.8c46.31 0 32.22-54.17 43.15-110.26C94.8 95.2 193.12 32 288.09 32c102.48 0 197.15 70.67 214.08 169.67M373.51 388.28c42-19.18 81.33-56.71 96.29-102.14c40.48-123.09-64.15-228-181.71-228c-83.45 0-170.32 55.42-186.07 136c-9.53 48.91 5 131.35-68.75 131.35C58.21 358.6 91.6 378.11 127 389.54c24.66-21.8 63.87-15.47 79.83 14.34c14.22 1 79.19 1.18 87.9.82a51.69 51.69 0 0 1 78.78-16.42M205.12 187.13c0-34.74 50.84-34.75 50.84 0s-50.84 34.74-50.84 0m116.57 0c0-34.74 50.86-34.75 50.86 0s-50.86 34.75-50.86 0m-122.61 70.69c-3.44-16.94 22.18-22.18 25.62-5.21l.06.28c4.14 21.42 29.85 44 64.12 43.07c35.68-.94 59.25-22.21 64.11-42.77c4.46-16.05 28.6-10.36 25.47 6c-5.23 22.18-31.21 62-91.46 62.9c-42.55 0-80.88-27.84-87.9-64.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Weebly(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M425.09 65.83c-39.88 0-73.28 25.73-83.66 64.33c-18.16-58.06-65.5-64.33-84.95-64.33c-19.78 0-66.8 6.28-85.28 64.33c-10.38-38.6-43.45-64.33-83.66-64.33C38.59 65.83 0 99.72 0 143.03c0 28.96 4.18 33.27 77.17 233.48c22.37 60.57 67.77 69.35 92.74 69.35c39.23 0 70.04-19.46 85.93-53.98c15.89 34.83 46.69 54.29 85.93 54.29c24.97 0 70.36-9.1 92.74-69.67c76.55-208.65 77.5-205.58 77.5-227.2c.63-48.32-36.01-83.47-86.92-83.47m26.34 114.81l-65.57 176.44c-7.92 21.49-21.22 37.22-46.24 37.22c-23.44 0-37.38-12.41-44.03-33.9l-39.28-117.42h-.95L216.08 360.4c-6.96 21.5-20.9 33.6-44.02 33.6c-25.02 0-38.33-15.74-46.24-37.22L60.88 181.55c-5.38-14.83-7.92-23.91-7.92-34.5c0-16.34 15.84-29.36 38.33-29.36c18.69 0 31.99 11.8 36.11 29.05l44.03 139.82h.95l44.66-136.79c6.02-19.67 16.47-32.08 38.96-32.08s32.94 12.11 38.96 32.08l44.66 136.79h.95l44.03-139.82c4.12-17.25 17.42-29.05 36.11-29.05c22.17 0 38.33 13.32 38.33 35.71c-.32 7.87-4.12 16.04-7.61 27.24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Weibo(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M407 177.6c7.6-24-13.4-46.8-37.4-41.7c-22 4.8-28.8-28.1-7.1-32.8c50.1-10.9 92.3 37.1 76.5 84.8c-6.8 21.2-38.8 10.8-32-10.3M214.8 446.7C108.5 446.7 0 395.3 0 310.4c0-44.3 28-95.4 76.3-143.7C176 67 279.5 65.8 249.9 161c-4 13.1 12.3 5.7 12.3 6c79.5-33.6 140.5-16.8 114 51.4c-3.7 9.4 1.1 10.9 8.3 13.1c135.7 42.3 34.8 215.2-169.7 215.2m143.7-146.3c-5.4-55.7-78.5-94-163.4-85.7c-84.8 8.6-148.8 60.3-143.4 116s78.5 94 163.4 85.7c84.8-8.6 148.8-60.3 143.4-116M347.9 35.1c-25.9 5.6-16.8 43.7 8.3 38.3c72.3-15.2 134.8 52.8 111.7 124c-7.4 24.2 29.1 37 37.4 12c31.9-99.8-55.1-195.9-157.4-174.3m-78.5 311c-17.1 38.8-66.8 60-109.1 46.3c-40.8-13.1-58-53.4-40.3-89.7c17.7-35.4 63.1-55.4 103.4-45.1c42 10.8 63.1 50.2 46 88.5m-86.3-30c-12.9-5.4-30 .3-38 12.9c-8.3 12.9-4.3 28 8.6 34c13.1 6 30.8.3 39.1-12.9c8-13.1 3.7-28.3-9.7-34m32.6-13.4c-5.1-1.7-11.4.6-14.3 5.4c-2.9 5.1-1.4 10.6 3.7 12.9c5.1 2 11.7-.3 14.6-5.4c2.8-5.2 1.1-10.9-4-12.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Weixin(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M385.2 167.6c6.4 0 12.6.3 18.8 1.1C387.4 90.3 303.3 32 207.7 32C100.5 32 13 104.8 13 197.4c0 53.4 29.3 97.5 77.9 131.6l-19.3 58.6l68-34.1c24.4 4.8 43.8 9.7 68.2 9.7c6.2 0 12.1-.3 18.3-.8c-4-12.9-6.2-26.6-6.2-40.8c-.1-84.9 72.9-154 165.3-154m-104.5-52.9c14.5 0 24.2 9.7 24.2 24.4c0 14.5-9.7 24.2-24.2 24.2c-14.8 0-29.3-9.7-29.3-24.2c.1-14.7 14.6-24.4 29.3-24.4m-136.4 48.6c-14.5 0-29.3-9.7-29.3-24.2c0-14.8 14.8-24.4 29.3-24.4c14.8 0 24.4 9.7 24.4 24.4c0 14.6-9.6 24.2-24.4 24.2M563 319.4c0-77.9-77.9-141.3-165.4-141.3c-92.7 0-165.4 63.4-165.4 141.3S305 460.7 397.6 460.7c19.3 0 38.9-5.1 58.6-9.9l53.4 29.3l-14.8-48.6C534 402.1 563 363.2 563 319.4m-219.1-24.5c-9.7 0-19.3-9.7-19.3-19.6c0-9.7 9.7-19.3 19.3-19.3c14.8 0 24.4 9.7 24.4 19.3c0 10-9.7 19.6-24.4 19.6m107.1 0c-9.7 0-19.3-9.7-19.3-19.6c0-9.7 9.7-19.3 19.3-19.3c14.5 0 24.4 9.7 24.4 19.3c.1 10-9.9 19.6-24.4 19.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Whatsapp(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M380.9 97.1C339 55.1 283.2 32 223.9 32c-122.4 0-222 99.6-222 222c0 39.1 10.2 77.3 29.6 111L0 480l117.7-30.9c32.4 17.7 68.9 27 106.1 27h.1c122.3 0 224.1-99.6 224.1-222c0-59.3-25.2-115-67.1-157m-157 341.6c-33.2 0-65.7-8.9-94-25.7l-6.7-4l-69.8 18.3L72 359.2l-4.4-7c-18.5-29.4-28.2-63.3-28.2-98.2c0-101.7 82.8-184.5 184.6-184.5c49.3 0 95.6 19.2 130.4 54.1c34.8 34.9 56.2 81.2 56.1 130.5c0 101.8-84.9 184.6-186.6 184.6m101.2-138.2c-5.5-2.8-32.8-16.2-37.9-18c-5.1-1.9-8.8-2.8-12.5 2.8c-3.7 5.6-14.3 18-17.6 21.8c-3.2 3.7-6.5 4.2-12 1.4c-32.6-16.3-54-29.1-75.5-66c-5.7-9.8 5.7-9.1 16.3-30.3c1.8-3.7.9-6.9-.5-9.7c-1.4-2.8-12.5-30.1-17.1-41.2c-4.5-10.8-9.1-9.3-12.5-9.5c-3.2-.2-6.9-.2-10.6-.2c-3.7 0-9.7 1.4-14.8 6.9c-5.1 5.6-19.4 19-19.4 46.3c0 27.3 19.9 53.7 22.6 57.4c2.8 3.7 39.1 59.7 94.8 83.8c35.2 15.2 49 16.5 66.6 13.9c10.7-1.6 32.8-13.4 37.4-26.4c4.6-13 4.6-24.1 3.2-26.4c-1.3-2.5-5-3.9-10.5-6.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WhatsappSquare(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M224 122.8c-72.7 0-131.8 59.1-131.9 131.8c0 24.9 7 49.2 20.2 70.1l3.1 5l-13.3 48.6l49.9-13.1l4.8 2.9c20.2 12 43.4 18.4 67.1 18.4h.1c72.6 0 133.3-59.1 133.3-131.8c0-35.2-15.2-68.3-40.1-93.2c-25-25-58-38.7-93.2-38.7m77.5 188.4c-3.3 9.3-19.1 17.7-26.7 18.8c-12.6 1.9-22.4.9-47.5-9.9c-39.7-17.2-65.7-57.2-67.7-59.8c-2-2.6-16.2-21.5-16.2-41s10.2-29.1 13.9-33.1c3.6-4 7.9-5 10.6-5c2.6 0 5.3 0 7.6.1c2.4.1 5.7-.9 8.9 6.8c3.3 7.9 11.2 27.4 12.2 29.4s1.7 4.3.3 6.9c-7.6 15.2-15.7 14.6-11.6 21.6c15.3 26.3 30.6 35.4 53.9 47.1c4 2 6.3 1.7 8.6-1c2.3-2.6 9.9-11.6 12.5-15.5c2.6-4 5.3-3.3 8.9-2c3.6 1.3 23.1 10.9 27.1 12.9s6.6 3 7.6 4.6c.9 1.9.9 9.9-2.4 19.1M400 32H48C21.5 32 0 53.5 0 80v352c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V80c0-26.5-21.5-48-48-48M223.9 413.2c-26.6 0-52.7-6.7-75.8-19.3L64 416l22.5-82.2c-13.9-24-21.2-51.3-21.2-79.3C65.4 167.1 136.5 96 223.9 96c42.4 0 82.2 16.5 112.2 46.5c29.9 30 47.9 69.8 47.9 112.2c0 87.4-72.7 158.5-160.1 158.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Whmcs(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 161v-21.3l-28.5-8.8l-2.2-10.4l20.1-20.7L427 80.4l-29 7.5l-7.2-7.5l7.5-28.2l-19.1-11.6l-21.3 21l-10.7-3.2l-7-26.4h-22.6l-6.2 26.4l-12.1 3.2l-19.7-21l-19.4 11l8.1 27.7l-8.1 8.4l-28.5-7.5l-11 19.1l20.7 21l-2.9 10.4l-28.5 7.8l-.3 21.7l28.8 7.5l2.4 12.1l-20.1 19.9l10.4 18.5l29.6-7.5l7.2 8.6l-8.1 26.9l19.9 11.6l19.4-20.4l11.6 2.9l6.7 28.5l22.6.3l6.7-28.8l11.6-3.5l20.7 21.6l20.4-12.1l-8.8-28l7.8-8.1l28.8 8.8l10.3-20.1l-20.9-18.8l2.2-12.1zm-119.2 45.2c-31.3 0-56.8-25.4-56.8-56.8s25.4-56.8 56.8-56.8s56.8 25.4 56.8 56.8c0 31.5-25.4 56.8-56.8 56.8m72.3 16.4l46.9 14.5V277l-55.1 13.4l-4.1 22.7l38.9 35.3l-19.2 37.9l-54-16.7l-14.6 15.2l16.7 52.5l-38.3 22.7l-38.9-40.5l-21.7 6.6l-12.6 54l-42.4-.5l-12.6-53.6l-21.7-5.6l-36.4 38.4l-37.4-21.7l15.2-50.5l-13.7-16.1l-55.5 14.1l-19.7-34.8l37.9-37.4l-4.8-22.8l-54-14.1l.5-40.9L54 219.9l5.7-19.7l-38.9-39.4L41.5 125l53.6 14.1l15.2-15.7l-15.2-52l36.4-20.7l36.8 39.4L191 84l11.6-52H245l11.6 45.9L234 72l-6.3-1.7l-3.3 5.7l-11 19.1l-3.3 5.6l4.6 4.6l17.2 17.4l-.3 1l-23.8 6.5l-6.2 1.7l-.1 6.4l-.2 12.9C153.8 161.6 118 204 118 254.7c0 58.3 47.3 105.7 105.7 105.7c50.5 0 92.7-35.4 103.2-82.8l13.2.2l6.9.1l1.6-6.7l5.6-24l1.9-.6l17.1 17.8l4.7 4.9l5.8-3.4l20.4-12.1l5.8-3.5l-2-6.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WikipediaW(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m640 51.2l-.3 12.2c-28.1.8-45 15.8-55.8 40.3c-25 57.8-103.3 240-155.3 358.6H415l-81.9-193.1c-32.5 63.6-68.3 130-99.2 193.1c-.3.3-15 0-15-.3C172 352.3 122.8 243.4 75.8 133.4C64.4 106.7 26.4 63.4.2 63.7c0-3.1-.3-10-.3-14.2h161.9v13.9c-19.2 1.1-52.8 13.3-43.3 34.2c21.9 49.7 103.6 240.3 125.6 288.6c15-29.7 57.8-109.2 75.3-142.8c-13.9-28.3-58.6-133.9-72.8-160c-9.7-17.8-36.1-19.4-55.8-19.7V49.8l142.5.3v13.1c-19.4.6-38.1 7.8-29.4 26.1c18.9 40 30.6 68.1 48.1 104.7c5.6-10.8 34.7-69.4 48.1-100.8c8.9-20.6-3.9-28.6-38.6-29.4c.3-3.6 0-10.3.3-13.6c44.4-.3 111.1-.3 123.1-.6v13.6c-22.5.8-45.8 12.8-58.1 31.7l-59.2 122.8c6.4 16.1 63.3 142.8 69.2 156.7L559.2 91.8c-8.6-23.1-36.4-28.1-47.2-28.3V49.6l127.8 1.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Windows(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 93.7l183.6-25.3v177.4H0zm0 324.6l183.6 25.3V268.4H0zm203.8 28L448 480V268.4H203.8zm0-380.6v180.1H448V32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wix(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M393.38 131.69c0 13.03 2.08 32.69-28.68 43.83c-9.52 3.45-15.95 9.66-15.95 9.66c0-31 4.72-42.22 17.4-48.86c9.75-5.11 27.23-4.63 27.23-4.63m-115.8 35.54l-34.24 132.66l-28.48-108.57c-7.69-31.99-20.81-48.53-48.43-48.53c-27.37 0-40.66 16.18-48.43 48.53L89.52 299.89L55.28 167.23C49.73 140.51 23.86 128.96 0 131.96l65.57 247.93s21.63 1.56 32.46-3.96c14.22-7.25 20.98-12.84 29.59-46.57c7.67-30.07 29.11-118.41 31.12-124.7c4.76-14.94 11.09-13.81 15.4 0c1.97 6.3 23.45 94.63 31.12 124.7c8.6 33.73 15.37 39.32 29.59 46.57c10.82 5.52 32.46 3.96 32.46 3.96l65.57-247.93c-24.42-3.07-49.82 8.93-55.3 35.27m115.78 5.21s-4.1 6.34-13.46 11.57c-6.01 3.36-11.78 5.64-17.97 8.61c-15.14 7.26-13.18 13.95-13.18 35.2v152.07s16.55 2.09 27.37-3.43c13.93-7.1 17.13-13.95 17.26-44.78V181.41l-.02.01zm163.44 84.08L640 132.78s-35.11-5.98-52.5 9.85c-13.3 12.1-24.41 29.55-54.18 72.47c-.47.73-6.25 10.54-13.07 0c-29.29-42.23-40.8-60.29-54.18-72.47c-17.39-15.83-52.5-9.85-52.5-9.85l83.2 123.74l-82.97 123.36s36.57 4.62 53.95-11.21c11.49-10.46 17.58-20.37 52.51-70.72c6.81-10.52 12.57-.77 13.07 0c29.4 42.38 39.23 58.06 53.14 70.72c17.39 15.83 53.32 11.21 53.32 11.21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WizardsOfTheCoast(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M219.19 345.69c-1.9 1.38-11.07 8.44-.26 23.57c4.64 6.42 14.11 12.79 21.73 6.55c6.5-4.88 7.35-12.92.26-23.04c-5.47-7.76-14.28-12.88-21.73-7.08m336.75 75.94c-.34 1.7-.55 1.67.79 0c2.09-4.19 4.19-10.21 4.98-19.9c3.14-38.49-40.33-71.49-101.34-78.03c-54.73-6.02-124.38 9.17-188.8 60.49l-.26 1.57c2.62 4.98 4.98 10.74 3.4 21.21l.79.26c63.89-58.4 131.19-77.25 184.35-73.85c58.4 3.67 100.03 34.04 100.03 68.08c-.01 9.96-2.63 15.72-3.94 20.17M392.28 240.42c.79 7.07 4.19 10.21 9.17 10.47c5.5.26 9.43-2.62 10.47-6.55c.79-3.4 2.09-29.85 2.09-29.85s-11.26 6.55-14.93 10.47c-3.66 3.68-7.33 8.39-6.8 15.46m-50.02-151.1C137.75 89.32 13.1 226.8.79 241.2c-1.05.52-1.31.79.79 1.31c60.49 16.5 155.81 81.18 196.13 202.16l1.05.26c55.25-69.92 140.88-128.05 236.99-128.05c80.92 0 130.15 42.16 130.15 80.39c0 18.33-6.55 33.52-22.26 46.35c0 .96-.2.79.79.79c14.66-10.74 27.5-28.8 27.5-48.18c0-22.78-12.05-38.23-12.05-38.23c7.07 7.07 10.74 16.24 10.74 16.24c5.76-40.85 26.97-62.32 26.97-62.32c-2.36-9.69-6.81-17.81-6.81-17.81c7.59 8.12 14.4 27.5 14.4 41.37c0 10.47-3.4 22.78-12.57 31.95l.26.52c8.12-4.98 16.5-16.76 16.5-37.97c0-15.71-4.71-25.92-4.71-25.92c5.76-5.24 11.26-9.17 15.97-11.78c.79 3.4 2.09 9.69 2.36 14.93c0 1.05.79 1.83 1.05 0c.79-5.76-.26-16.24-.26-16.5c6.02-3.14 9.69-4.45 9.69-4.45C617.74 176 489.43 89.32 342.26 89.32m-99.24 289.62c-11.06 8.99-24.2 4.08-30.64-4.19c-7.45-9.58-6.76-24.09 4.19-32.47c14.85-11.35 27.08-.49 31.16 5.5c.28.39 12.13 16.57-4.71 31.16m2.09-136.43l9.43-17.81l11.78 70.96l-12.57 6.02l-24.62-28.8l14.14-26.71l3.67 4.45zm18.59 117.58l-.26-.26c2.05-4.1-2.5-6.61-17.54-31.69c-1.31-2.36-3.14-2.88-4.45-2.62l-.26-.52c7.86-5.76 15.45-10.21 25.4-15.71l.52.26c1.31 1.83 2.09 2.88 3.4 4.71l-.26.52c-1.05-.26-2.36-.79-5.24.26c-2.09.79-7.86 3.67-12.31 7.59v1.31c1.57 2.36 3.93 6.55 5.76 9.69h.26c10.05-6.28 7.56-4.55 11.52-7.86h.26c.52 1.83.52 1.83 1.83 5.5l-.26.26c-3.06.61-4.65.34-11.52 5.5v.26c9.46 17.02 11.01 16.75 12.57 15.97l.26.26c-2.34 1.59-6.27 4.21-9.68 6.57m55.26-32.47c-3.14 1.57-6.02 2.88-9.95 4.98l-.26-.26c1.29-2.59 1.16-2.71-11.78-32.47l-.26-.26c-.15 0-8.9 3.65-9.95 7.33h-.52l-1.05-5.76l.26-.52c7.29-4.56 25.53-11.64 27.76-12.57l.52.26l3.14 4.98l-.26.52c-3.53-1.76-7.35.76-12.31 2.62v.26c12.31 32.01 12.67 30.64 14.66 30.64zm44.77-16.5c-4.19 1.05-5.24 1.31-9.69 2.88l-.26-.26l.52-4.45c-1.05-3.4-3.14-11.52-3.67-13.62l-.26-.26c-3.4.79-8.9 2.62-12.83 3.93l-.26.26c.79 2.62 3.14 9.95 4.19 13.88c.79 2.36 1.83 2.88 2.88 3.14v.52c-3.67 1.05-7.07 2.62-10.21 3.93l-.26-.26c1.05-1.31 1.05-2.88.26-4.98c-1.05-3.14-8.12-23.83-9.17-27.23c-.52-1.83-1.57-3.14-2.62-3.14v-.52c3.14-1.05 6.02-2.09 10.74-3.4l.26.26l-.26 4.71c1.31 3.93 2.36 7.59 3.14 9.69h.26c3.93-1.31 9.43-2.88 12.83-3.93l.26-.26l-2.62-9.43c-.52-1.83-1.05-3.4-2.62-3.93v-.26c4.45-1.05 7.33-1.83 10.74-2.36l.26.26c-1.05 1.31-1.05 2.88-.52 4.45c1.57 6.28 4.71 20.43 6.28 26.45c.54 2.62 1.85 3.41 2.63 3.93m32.21-6.81l-.26.26c-4.71.52-14.14 2.36-22.52 4.19l-.26-.26l.79-4.19c-1.57-7.86-3.4-18.59-4.98-26.19c-.26-1.83-.79-2.88-2.62-3.67l.79-.52c9.17-1.57 20.16-2.36 24.88-2.62l.26.26c.52 2.36.79 3.14 1.57 5.5l-.26.26c-1.14-1.14-3.34-3.2-16.24-.79l-.26.26c.26 1.57 1.05 6.55 1.57 9.95l.26.26c9.52-1.68 4.76-.06 10.74-2.36h.26c0 1.57-.26 1.83-.26 5.24h-.26c-4.81-1.03-2.15-.9-10.21 0l-.26.26c.26 2.09 1.57 9.43 2.09 12.57l.26.26c1.15.38 14.21-.65 16.24-4.71h.26c-.53 2.38-1.05 4.21-1.58 6.04m10.74-44.51c-4.45 2.36-8.12 2.88-11 2.88c-.25.02-11.41 1.09-17.54-9.95c-6.74-10.79-.98-25.2 5.5-31.69c8.8-8.12 23.35-10.1 28.54-17.02c8.03-10.33-13.04-22.31-29.59-5.76l-2.62-2.88l5.24-16.24c25.59-1.57 45.2-3.04 50.02 16.24c.79 3.14 0 9.43-.26 12.05c0 2.62-1.83 18.85-2.09 23.04c-.52 4.19-.79 18.33-.79 20.69c.26 2.36.52 4.19 1.57 5.5c1.57 1.83 5.76 1.83 5.76 1.83l-.79 4.71c-11.82-1.07-10.28-.59-20.43-1.05c-3.22-5.15-2.23-3.28-4.19-7.86c0 .01-4.19 3.94-7.33 5.51m37.18 21.21c-6.35-10.58-19.82-7.16-21.73 5.5c-2.63 17.08 14.3 19.79 20.69 10.21l.26.26c-.52 1.83-1.83 6.02-1.83 6.28l-.52.52c-10.3 6.87-28.5-2.5-25.66-18.59c1.94-10.87 14.44-18.93 28.8-9.95l.26.52c0 1.06-.27 3.41-.27 5.25m5.77-87.73v-6.55c.69 0 19.65 3.28 27.76 7.33l-1.57 17.54s10.21-9.43 15.45-10.74c5.24-1.57 14.93 7.33 14.93 7.33l-11.26 11.26c-12.07-6.35-19.59-.08-20.69.79c-5.29 38.72-8.6 42.17 4.45 46.09l-.52 4.71c-17.55-4.29-18.53-4.5-36.92-7.33l.79-4.71c7.25 0 7.48-5.32 7.59-6.81c0 0 4.98-53.16 4.98-55.25c-.02-2.87-4.99-3.66-4.99-3.66m10.99 114.44c-8.12-2.09-14.14-11-10.74-20.69c3.14-9.43 12.31-12.31 18.85-10.21c9.17 2.62 12.83 11.78 10.74 19.38c-2.61 8.9-9.42 13.87-18.85 11.52m42.16 9.69c-2.36-.52-7.07-2.36-8.64-2.88v-.26l1.57-1.83c.59-8.24.59-7.27.26-7.59c-4.82-1.81-6.66-2.36-7.07-2.36c-1.31 1.83-2.88 4.45-3.67 5.5l-.79 3.4v.26c-1.31-.26-3.93-1.31-6.02-1.57v-.26l2.62-1.83c3.4-4.71 9.95-14.14 13.88-20.16v-2.09l.52-.26c2.09.79 5.5 2.09 7.59 2.88c.48.48.18-1.87-1.05 25.14c-.24 1.81.02 2.6.8 3.91m-4.71-89.82c11.25-18.27 30.76-16.19 34.04-3.4L539.7 198c2.34-6.25-2.82-9.9-4.45-11.26l1.83-3.67c12.22 10.37 16.38 13.97 22.52 20.43c-25.91 73.07-30.76 80.81-24.62 84.32l-1.83 4.45c-6.37-3.35-8.9-4.42-17.81-8.64l2.09-6.81c-.26-.26-3.93 3.93-9.69 3.67c-19.06-1.3-22.89-31.75-9.67-52.9m29.33 79.34c0-5.71-6.34-7.89-7.86-5.24c-1.31 2.09 1.05 4.98 2.88 8.38c1.57 2.62 2.62 6.28 1.05 9.43c-2.64 6.34-12.4 5.31-15.45-.79c0-.7-.27.09 1.83-4.71l.79-.26c-.57 5.66 6.06 9.61 8.38 4.98c1.05-2.09-.52-5.5-2.09-8.38c-1.57-2.62-3.67-6.28-1.83-9.69c2.72-5.06 11.25-4.47 14.66 2.36v.52zm21.21 13.36c-1.96-3.27-.91-2.14-4.45-4.71h-.26c-2.36 4.19-5.76 10.47-8.64 16.24c-1.31 2.36-1.05 3.4-.79 3.93l-.26.26l-5.76-4.45l.26-.26l2.09-1.31c3.14-5.76 6.55-12.05 9.17-17.02v-.26c-2.64-1.98-1.22-1.51-6.02-1.83v-.26l3.14-3.4h.26c3.67 2.36 9.95 6.81 12.31 8.9l.26.26zm27.23-44.26l-2.88-2.88c.79-2.36 1.83-4.98 2.09-7.59c.75-9.74-11.52-11.84-11.52-4.98c0 4.98 7.86 19.38 7.86 27.76c0 10.21-5.76 15.71-13.88 16.5c-8.38.79-20.16-10.47-20.16-10.47l4.98-14.4l2.88 2.09c-2.97 17.8 17.68 20.37 13.35 5.24c-1.06-4.02-18.75-34.2 2.09-38.23c13.62-2.36 23.04 16.5 23.04 16.5zm35.62-10.21c-11-30.38-60.49-127.53-191.95-129.62c-53.42-1.05-94.27 15.45-132.76 37.97l85.63-9.17l-91.39 20.69l25.14 19.64l-3.93-16.5c7.5-1.71 39.15-8.45 66.77-8.9l-22.26 80.39c13.61-.7 18.97-8.98 19.64-22.78l4.98-1.05l.26 26.71c-22.46 3.21-37.3 6.69-49.49 9.95l13.09-43.21l-61.54-36.66l2.36 8.12l10.21 4.98c6.28 18.59 19.38 56.56 20.43 58.66c1.95 4.28 3.16 5.78 12.05 4.45l1.05 4.98c-16.08 4.86-23.66 7.61-39.02 14.4l-2.36-4.71c4.4-2.94 8.73-3.94 5.5-12.83c-23.7-62.5-21.48-58.14-22.78-59.44l2.36-4.45l33.52 67.3c-3.84-11.87 1.68 1.69-32.99-78.82l-41.9 88.51l4.71-13.88l-35.88-42.16l27.76 93.48l-11.78 8.38C95 228.58 101.05 231.87 93.23 231.52c-5.5-.26-13.62 5.5-13.62 5.5L74.63 231c30.56-23.53 31.62-24.33 58.4-42.68l4.19 7.07s-5.76 4.19-7.86 7.07c-5.9 9.28 1.67 13.28 61.8 75.68l-18.85-58.92l39.8-10.21l25.66 30.64l4.45-12.31l-4.98-24.62l13.09-3.4l.52 3.14l3.67-10.47l-94.27 29.33l11.26-4.98l-13.62-42.42l17.28-9.17l30.11 36.14l28.54-13.09c-1.41-7.47-2.47-14.5-4.71-19.64l17.28 13.88l4.71-2.09l-59.18-42.68l23.08 11.5c18.98-6.07 25.23-7.47 32.21-9.69l2.62 11c-12.55 12.55 1.43 16.82 6.55 19.38l-13.62-61.01l12.05 28.28c4.19-1.31 7.33-2.09 7.33-2.09l2.62 8.64s-3.14 1.05-6.28 2.09l8.9 20.95l33.78-65.73l-20.69 61.01c42.42-24.09 81.44-36.66 131.98-35.88c67.04 1.05 167.33 40.85 199.8 139.83c.78 2.1-.01 2.63-.79.27M203.48 152.43s1.83-.52 4.19-1.31l9.43 7.59c-.4 0-3.44-.25-11.26 2.36zm143.76 38.5c-1.57-.6-26.46-4.81-33.26 20.69l21.73 17.02zM318.43 67.07c-58.4 0-106.05 12.05-114.96 14.4v.79c8.38 2.09 14.4 4.19 21.21 11.78l1.57.26c6.55-1.83 48.97-13.88 110.24-13.88c180.16 0 301.67 116.79 301.67 223.37v9.95c0 1.31.79 2.62 1.05.52c.52-2.09.79-8.64.79-19.64c.26-83.79-96.63-227.55-321.57-227.55m211.06 169.68c1.31-5.76 0-12.31-7.33-13.09c-9.62-1.13-16.14 23.79-17.02 33.52c-.79 5.5-1.31 14.93 6.02 14.93c4.68-.01 9.72-.91 18.33-35.36m-61.53 42.95c-2.62-.79-9.43-.79-12.57 10.47c-1.83 6.81.52 13.35 6.02 14.66c3.67 1.05 8.9.52 11.78-10.74c2.62-9.94-1.83-13.61-5.23-14.39M491 300.65c1.83.52 3.14 1.05 5.76 1.83c0-1.83.52-8.38.79-12.05c-1.05 1.31-5.5 8.12-6.55 9.95z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wodu(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M178.414 339.706H141.1l-28.934-116.231h-.478l-28.46 116.231H45.2L0 168.946h37.548l27.026 116.231h.478l29.655-116.231h35.157l29.178 117.667h.479L187.5 168.946h36.831zM271.4 212.713c38.984 0 64.1 25.828 64.1 65.291c0 39.222-25.111 65.05-64.1 65.05c-38.743 0-63.855-25.828-63.855-65.05c.002-39.463 25.114-65.291 63.855-65.291m0 104.753c23.2 0 30.133-19.852 30.133-39.462c0-19.852-6.934-39.7-30.133-39.7c-27.7 0-29.894 19.85-29.894 39.7c.002 19.61 6.937 39.462 29.894 39.462m163.684 6.456h-.478c-7.893 13.392-21.765 19.132-37.548 19.132c-37.31 0-55.485-32.045-55.485-66.246c0-33.243 18.415-64.095 54.767-64.095c14.589 0 28.938 6.218 36.831 18.416h.24v-62.183h33.96v170.76h-32.287zM405.428 238.3c-22.24 0-29.894 19.134-29.894 39.463c0 19.371 8.848 39.7 29.894 39.7c22.482 0 29.178-19.613 29.178-39.94c0-20.087-7.174-39.223-29.178-39.223M592.96 339.706h-32.287v-17.219h-.718c-8.609 13.87-23.436 20.567-37.786 20.567c-36.113 0-45.2-20.328-45.2-50.941v-76.052h33.959V285.9c0 20.329 5.979 30.372 21.765 30.372c18.415 0 26.306-10.283 26.306-35.393v-64.818h33.961zm9.493-36.83H640v36.83h-37.547z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WolfPackBattalion(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m267.73 471.53l10.56 15.84l5.28-12.32l5.28 7V512c21.06-7.92 21.11-66.86 25.51-97.21c4.62-31.89-.88-92.81 81.37-149.11c-8.88-23.61-12-49.43-2.64-80.05C421 189 447 196.21 456.43 239.73l-30.35 8.36c11.15 23 17 46.76 13.2 72.14L412 313.18l-6.16 33.43l-18.47-7l-8.8 33.39l-19.35-7l26.39 21.11l8.8-28.15L419 364.2l7-35.63l26.39 14.52c.25-20 7-58.06-8.8-84.45l26.39 5.28c4-22.07-2.38-39.21-7.92-56.74l22.43 9.68c-.44-25.07-29.94-56.79-61.58-58.5c-20.22-1.09-56.74-25.17-54.1-51.9c2-19.87 17.45-42.62 43.11-49.7c-44 36.51-9.68 67.3 5.28 73.46c4.4-11.44 17.54-69.08 0-130.2c-40.39 22.87-89.65 65.1-93.2 147.79l-58 38.71l-3.52 93.25L369.78 220l7 7l-17.59 3.52l-44 38.71l-15.84-5.28l-28.1 49.25l-3.52 119.64l21.11 15.84l-32.55 15.84l-32.55-15.84l21.11-15.84l-3.52-119.64l-28.15-49.26l-15.84 5.28l-44-38.71l-17.58-3.51l7-7l107.33 59.82l-3.52-93.25l-58.06-38.71C185 65.1 135.77 22.87 95.3 0c-17.54 61.12-4.4 118.76 0 130.2c15-6.16 49.26-36.95 5.28-73.46c25.66 7.08 41.15 29.83 43.11 49.7c2.63 26.74-33.88 50.81-54.1 51.9c-31.65 1.72-61.15 33.44-61.59 58.51l22.43-9.68c-5.54 17.53-11.91 34.67-7.92 56.74l26.39-5.28c-15.76 26.39-9.05 64.43-8.8 84.45l26.39-14.52l7 35.63l24.63-5.28l8.8 28.15L153.35 366L134 373l-8.8-33.43l-18.47 7l-6.16-33.43l-27.27 7c-3.82-25.38 2-49.1 13.2-72.14l-30.35-8.36c9.4-43.52 35.47-50.77 63.34-54.1c9.36 30.62 6.24 56.45-2.64 80.05c82.25 56.3 76.75 117.23 81.37 149.11c4.4 30.35 4.45 89.29 25.51 97.21v-29.83l5.28-7l5.28 12.32l10.56-15.84l11.44 21.11l11.43-21.1zm79.17-95L331.06 366c7.47-4.36 13.76-8.42 19.35-12.32c-.6 7.22-.27 13.84-3.51 22.84zm28.15-49.26c-.4 10.94-.9 21.66-1.76 31.67c-7.85-1.86-15.57-3.8-21.11-7c8.24-7.94 15.55-16.32 22.87-24.68zm24.63 5.28c0-13.43-2.05-24.21-5.28-33.43a235 235 0 0 1-18.47 27.27zm3.52-80.94c19.44 12.81 27.8 33.66 29.91 56.3c-12.32-4.53-24.63-9.31-36.95-10.56c5.06-12 6.65-28.14 7-45.74zm-1.76-45.74c.81 14.3 1.84 28.82 1.76 42.23c19.22-8.11 29.78-9.72 44-14.08c-10.61-18.96-27.2-25.53-45.76-28.16zM165.68 376.52L181.52 366c-7.47-4.36-13.76-8.42-19.35-12.32c.6 7.26.27 13.88 3.51 22.88zm-28.15-49.26c.4 10.94.9 21.66 1.76 31.67c7.85-1.86 15.57-3.8 21.11-7c-8.24-7.93-15.55-16.31-22.87-24.67m-24.64 5.28c0-13.43 2-24.21 5.28-33.43a235 235 0 0 0 18.47 27.27zm-3.52-80.94c-19.44 12.81-27.8 33.66-29.91 56.3c12.32-4.53 24.63-9.31 37-10.56c-5-12-6.65-28.14-7-45.74zm1.76-45.74c-.81 14.3-1.84 28.82-1.76 42.23c-19.22-8.11-29.78-9.72-44-14.08c10.63-18.95 27.23-25.52 45.76-28.15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wordpress(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m61.7 169.4l101.5 278C92.2 413 43.3 340.2 43.3 256c0-30.9 6.6-60.1 18.4-86.6m337.9 75.9c0-26.3-9.4-44.5-17.5-58.7c-10.8-17.5-20.9-32.4-20.9-49.9c0-19.6 14.8-37.8 35.7-37.8c.9 0 1.8.1 2.8.2c-37.9-34.7-88.3-55.9-143.7-55.9c-74.3 0-139.7 38.1-177.8 95.9c5 .2 9.7.3 13.7.3c22.2 0 56.7-2.7 56.7-2.7c11.5-.7 12.8 16.2 1.4 17.5c0 0-11.5 1.3-24.3 2l77.5 230.4L249.8 247l-33.1-90.8c-11.5-.7-22.3-2-22.3-2c-11.5-.7-10.1-18.2 1.3-17.5c0 0 35.1 2.7 56 2.7c22.2 0 56.7-2.7 56.7-2.7c11.5-.7 12.8 16.2 1.4 17.5c0 0-11.5 1.3-24.3 2l76.9 228.7l21.2-70.9c9-29.4 16-50.5 16-68.7m-139.9 29.3l-63.8 185.5c19.1 5.6 39.2 8.7 60.1 8.7c24.8 0 48.5-4.3 70.6-12.1c-.6-.9-1.1-1.9-1.5-2.9zm183-120.7c.9 6.8 1.4 14 1.4 21.9c0 21.6-4 45.8-16.2 76.2l-65 187.9C426.2 403 468.7 334.5 468.7 256c0-37-9.4-71.8-26-102.1M504 256c0 136.8-111.3 248-248 248C119.2 504 8 392.7 8 256C8 119.2 119.2 8 256 8c136.7 0 248 111.2 248 248m-11.4 0c0-130.5-106.2-236.6-236.6-236.6C125.5 19.4 19.4 125.5 19.4 256S125.6 492.6 256 492.6c130.5 0 236.6-106.1 236.6-236.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WordpressSimple(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 8C119.3 8 8 119.2 8 256c0 136.7 111.3 248 248 248s248-111.3 248-248C504 119.2 392.7 8 256 8M33 256c0-32.3 6.9-63 19.3-90.7l106.4 291.4C84.3 420.5 33 344.2 33 256m223 223c-21.9 0-43-3.2-63-9.1l66.9-194.4l68.5 187.8c.5 1.1 1 2.1 1.6 3.1c-23.1 8.1-48 12.6-74 12.6m30.7-327.5c13.4-.7 25.5-2.1 25.5-2.1c12-1.4 10.6-19.1-1.4-18.4c0 0-36.1 2.8-59.4 2.8c-21.9 0-58.7-2.8-58.7-2.8c-12-.7-13.4 17.7-1.4 18.4c0 0 11.4 1.4 23.4 2.1l34.7 95.2L200.6 393l-81.2-241.5c13.4-.7 25.5-2.1 25.5-2.1c12-1.4 10.6-19.1-1.4-18.4c0 0-36.1 2.8-59.4 2.8c-4.2 0-9.1-.1-14.4-.3C109.6 73 178.1 33 256 33c58 0 110.9 22.2 150.6 58.5c-1-.1-1.9-.2-2.9-.2c-21.9 0-37.4 19.1-37.4 39.6c0 18.4 10.6 33.9 21.9 52.3c8.5 14.8 18.4 33.9 18.4 61.5c0 19.1-7.3 41.2-17 72.1l-22.2 74.3zm81.4 297.2l68.1-196.9c12.7-31.8 17-57.2 17-79.9c0-8.2-.5-15.8-1.5-22.9c17.4 31.8 27.3 68.2 27.3 107c0 82.3-44.6 154.1-110.9 192.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wpbeginner(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M462.799 322.374C519.01 386.682 466.961 480 370.944 480c-39.602 0-78.824-17.687-100.142-50.04c-6.887.356-22.702.356-29.59 0C219.848 462.381 180.588 480 141.069 480c-95.49 0-148.348-92.996-91.855-157.626C-29.925 190.523 80.479 32 256.006 32c175.632 0 285.87 158.626 206.793 290.374m-339.647-82.972h41.529v-58.075h-41.529zm217.18 86.072v-23.839c-60.506 20.915-132.355 9.198-187.589-33.971l.246 24.897c51.101 46.367 131.746 57.875 187.343 32.913m-150.753-86.072h166.058v-58.075H189.579z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wpexplorer(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 256c0 141.2-114.7 256-256 256C114.8 512 0 397.3 0 256S114.7 0 256 0s256 114.7 256 256m-32 0c0-123.2-100.3-224-224-224C132.5 32 32 132.5 32 256s100.5 224 224 224s224-100.5 224-224M160.9 124.6l86.9 37.1l-37.1 86.9l-86.9-37.1zm110 169.1l46.6 94h-14.6l-50-100l-48.9 100h-14l51.1-106.9l-22.3-9.4l6-14l68.6 29.1l-6 14.3zm-11.8-116.3l68.6 29.4l-29.4 68.3L230 246zm80.3 42.9l54.6 23.1l-23.4 54.3l-54.3-23.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wpforms(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 75.2v361.7c0 24.3-19 43.2-43.2 43.2H43.2C19.3 480 0 461.4 0 436.8V75.2C0 51.1 18.8 32 43.2 32h361.7c24 0 43.1 18.8 43.1 43.2m-37.3 361.6V75.2c0-3-2.6-5.8-5.8-5.8h-9.3L285.3 144L224 94.1L162.8 144L52.5 69.3h-9.3c-3.2 0-5.8 2.8-5.8 5.8v361.7c0 3 2.6 5.8 5.8 5.8h361.7c3.2.1 5.8-2.7 5.8-5.8M150.2 186v37H76.7v-37zm0 74.4v37.3H76.7v-37.3zm11.1-147.3l54-43.7H96.8zm210 72.9v37h-196v-37zm0 74.4v37.3h-196v-37.3zm-84.6-147.3l64.5-43.7H232.8zM371.3 335v37.3h-99.4V335z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wpressr(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M248 8C111.03 8 0 119.03 0 256s111.03 248 248 248s248-111.03 248-248S384.97 8 248 8m171.33 158.6c-15.18 34.51-30.37 69.02-45.63 103.5c-2.44 5.51-6.89 8.24-12.97 8.24c-23.02-.01-46.03.06-69.05-.05c-5.12-.03-8.25 1.89-10.34 6.72c-10.19 23.56-20.63 47-30.95 70.5c-1.54 3.51-4.06 5.29-7.92 5.29c-45.94-.01-91.87-.02-137.81 0c-3.13 0-5.63-1.15-7.72-3.45c-11.21-12.33-22.46-24.63-33.68-36.94c-2.69-2.95-2.79-6.18-1.21-9.73c8.66-19.54 17.27-39.1 25.89-58.66c12.93-29.35 25.89-58.69 38.75-88.08c1.7-3.88 4.28-5.68 8.54-5.65c14.24.1 28.48.02 42.72.05c6.24.01 9.2 4.84 6.66 10.59c-13.6 30.77-27.17 61.55-40.74 92.33c-5.72 12.99-11.42 25.99-17.09 39c-3.91 8.95 7.08 11.97 10.95 5.6c.23-.37-1.42 4.18 30.01-67.69c1.36-3.1 3.41-4.4 6.77-4.39c15.21.08 30.43.02 45.64.04c5.56.01 7.91 3.64 5.66 8.75c-8.33 18.96-16.71 37.9-24.98 56.89c-4.98 11.43 8.08 12.49 11.28 5.33c.04-.08 27.89-63.33 32.19-73.16c2.02-4.61 5.44-6.51 10.35-6.5c26.43.05 52.86 0 79.29.05c12.44.02 13.93-13.65 3.9-13.64c-25.26.03-50.52.02-75.78.02c-6.27 0-7.84-2.47-5.27-8.27c5.78-13.06 11.59-26.11 17.3-39.21c1.73-3.96 4.52-5.79 8.84-5.78c23.09.06 25.98.02 130.78.03c6.08-.01 8.03 2.79 5.62 8.27"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xbox(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M369.9 318.2c44.3 54.3 64.7 98.8 54.4 118.7c-7.9 15.1-56.7 44.6-92.6 55.9c-29.6 9.3-68.4 13.3-100.4 10.2c-38.2-3.7-76.9-17.4-110.1-39C93.3 445.8 87 438.3 87 423.4c0-29.9 32.9-82.3 89.2-142.1c32-33.9 76.5-73.7 81.4-72.6c9.4 2.1 84.3 75.1 112.3 109.5M188.6 143.8c-29.7-26.9-58.1-53.9-86.4-63.4c-15.2-5.1-16.3-4.8-28.7 8.1c-29.2 30.4-53.5 79.7-60.3 122.4c-5.4 34.2-6.1 43.8-4.2 60.5c5.6 50.5 17.3 85.4 40.5 120.9c9.5 14.6 12.1 17.3 9.3 9.9c-4.2-11-.3-37.5 9.5-64c14.3-39 53.9-112.9 120.3-194.4m311.6 63.5C483.3 127.3 432.7 77 425.6 77c-7.3 0-24.2 6.5-36 13.9c-23.3 14.5-41 31.4-64.3 52.8C367.7 197 427.5 283.1 448.2 346c6.8 20.7 9.7 41.1 7.4 52.3c-1.7 8.5-1.7 8.5 1.4 4.6c6.1-7.7 19.9-31.3 25.4-43.5c7.4-16.2 15-40.2 18.6-58.7c4.3-22.5 3.9-70.8-.8-93.4M141.3 43C189 40.5 251 77.5 255.6 78.4c.7.1 10.4-4.2 21.6-9.7c63.9-31.1 94-25.8 107.4-25.2c-63.9-39.3-152.7-50-233.9-11.7c-23.4 11.1-24 11.9-9.4 11.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xing(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M162.7 210c-1.8 3.3-25.2 44.4-70.1 123.5c-4.9 8.3-10.8 12.5-17.7 12.5H9.8c-7.7 0-12.1-7.5-8.5-14.4l69-121.3c.2 0 .2-.1 0-.3l-43.9-75.6c-4.3-7.8.3-14.1 8.5-14.1H100c7.3 0 13.3 4.1 18 12.2zM382.6 46.1l-144 253v.3L330.2 466c3.9 7.1.2 14.1-8.5 14.1h-65.2c-7.6 0-13.6-4-18-12.2l-92.4-168.5c3.3-5.8 51.5-90.8 144.8-255.2c4.6-8.1 10.4-12.2 17.5-12.2h65.7c8 0 12.3 6.7 8.5 14.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func XingSquare(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 32H48C21.5 32 0 53.5 0 80v352c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V80c0-26.5-21.5-48-48-48M140.4 320.2H93.8c-5.5 0-8.7-5.3-6-10.3l49.3-86.7c.1 0 .1-.1 0-.2l-31.4-54c-3-5.6.2-10.1 6-10.1h46.6c5.2 0 9.5 2.9 12.9 8.7l31.9 55.3c-1.3 2.3-18 31.7-50.1 88.2c-3.5 6.2-7.7 9.1-12.6 9.1m219.7-214.1L257.3 286.8v.2l65.5 119c2.8 5.1.1 10.1-6 10.1h-46.6c-5.5 0-9.7-2.9-12.9-8.7l-66-120.3c2.3-4.1 36.8-64.9 103.4-182.3c3.3-5.8 7.4-8.7 12.5-8.7h46.9c5.7-.1 8.8 4.7 6 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ycombinator(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 32v448H0V32zM236 287.5L313.5 142h-32.7L235 233c-4.7 9.3-9 18.3-12.8 26.8L210 233l-45.2-91h-35l76.7 143.8v94.5H236z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yahoo(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M223.69 141.06L167 284.23l-56-143.17H14.93l105.83 249.13L82.19 480h94.17l140.91-338.94Zm105.4 135.79a58.22 58.22 0 1 0 58.22 58.22a58.22 58.22 0 0 0-58.22-58.22M394.65 32l-93 223.47h104.79L499.07 32Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yammer(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M421.78 152.17A23.06 23.06 0 0 0 400.9 112c-.83.43-1.71.9-2.63 1.4c-15.25 8.4-118.33 80.62-106.69 88.77s82.04-23.61 130.2-50m0 217.17c-48.16-26.38-118.64-58.1-130.2-50s91.42 80.35 106.69 88.74c.92.51 1.8 1 2.63 1.41a23.07 23.07 0 0 0 20.88-40.15M464.21 237c-.95 0-1.95-.06-3-.06c-17.4 0-142.52 13.76-136.24 26.51s83.3 18.74 138.21 18.76a23 23 0 0 0 1-45.21zM31 96.65a24.88 24.88 0 0 1 46.14-18.4l81 205.06h1.21l77-203.53a23.52 23.52 0 0 1 44.45 15.27L171.2 368.44C152.65 415.66 134.08 448 77.91 448a139.67 139.67 0 0 1-23.81-1.95a21.31 21.31 0 0 1 6.9-41.77c.66.06 10.91.66 13.86.66c30.47 0 43.74-18.94 58.07-59.41z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yandex(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M153.1 315.8L65.7 512H2l96-209.8c-45.1-22.9-75.2-64.4-75.2-141.1C22.7 53.7 90.8 0 171.7 0H254v512h-55.1V315.8zm45.8-269.3h-29.4c-44.4 0-87.4 29.4-87.4 114.6c0 82.3 39.4 108.8 87.4 108.8h29.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YandexInternational(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M129.5 512V345.9L18.5 48h55.8l81.8 229.7L250.2 0h51.3L180.8 347.8V512z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yarn(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M393.9 345.2c-39 9.3-48.4 32.1-104 47.4c0 0-2.7 4-10.4 5.8c-13.4 3.3-63.9 6-68.5 6.1c-12.4.1-19.9-3.2-22-8.2c-6.4-15.3 9.2-22 9.2-22c-8.1-5-9-9.9-9.8-8.1c-2.4 5.8-3.6 20.1-10.1 26.5c-8.8 8.9-25.5 5.9-35.3.8c-10.8-5.7.8-19.2.8-19.2s-5.8 3.4-10.5-3.6c-6-9.3-17.1-37.3 11.5-62c-1.3-10.1-4.6-53.7 40.6-85.6c0 0-20.6-22.8-12.9-43.3c5-13.4 7-13.3 8.6-13.9c5.7-2.2 11.3-4.6 15.4-9.1c20.6-22.2 46.8-18 46.8-18s12.4-37.8 23.9-30.4c3.5 2.3 16.3 30.6 16.3 30.6s13.6-7.9 15.1-5c8.2 16 9.2 46.5 5.6 65.1c-6.1 30.6-21.4 47.1-27.6 57.5c-1.4 2.4 16.5 10 27.8 41.3c10.4 28.6 1.1 52.7 2.8 55.3c.8 1.4 13.7.8 36.4-13.2c12.8-7.9 28.1-16.9 45.4-17c16.7-.5 17.6 19.2 4.9 22.2M496 256c0 136.9-111.1 248-248 248S0 392.9 0 256S111.1 8 248 8s248 111.1 248 248m-79.3 75.2c-1.7-13.6-13.2-23-28-22.8c-22 .3-40.5 11.7-52.8 19.2c-4.8 3-8.9 5.2-12.4 6.8c3.1-44.5-22.5-73.1-28.7-79.4c7.8-11.3 18.4-27.8 23.4-53.2c4.3-21.7 3-55.5-6.9-74.5c-1.6-3.1-7.4-11.2-21-7.4c-9.7-20-13-22.1-15.6-23.8c-1.1-.7-23.6-16.4-41.4 28c-12.2.9-31.3 5.3-47.5 22.8c-2 2.2-5.9 3.8-10.1 5.4h.1c-8.4 3-12.3 9.9-16.9 22.3c-6.5 17.4.2 34.6 6.8 45.7c-17.8 15.9-37 39.8-35.7 82.5c-34 36-11.8 73-5.6 79.6c-1.6 11.1 3.7 19.4 12 23.8c12.6 6.7 30.3 9.6 43.9 2.8c4.9 5.2 13.8 10.1 30 10.1c6.8 0 58-2.9 72.6-6.5c6.8-1.6 11.5-4.5 14.6-7.1c9.8-3.1 36.8-12.3 62.2-28.7c18-11.7 24.2-14.2 37.6-17.4c12.9-3.2 21-15.1 19.4-28.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yelp(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m42.9 240.32l99.62 48.61c19.2 9.4 16.2 37.51-4.5 42.71L30.5 358.45a22.79 22.79 0 0 1-28.21-19.6a197.16 197.16 0 0 1 9-85.32a22.8 22.8 0 0 1 31.61-13.21m44 239.25a199.45 199.45 0 0 0 79.42 32.11A22.78 22.78 0 0 0 192.94 490l3.9-110.82c.7-21.3-25.5-31.91-39.81-16.1l-74.21 82.4a22.82 22.82 0 0 0 4.09 34.09zm145.34-109.92l58.81 94a22.93 22.93 0 0 0 34 5.5a198.36 198.36 0 0 0 52.71-67.61A23 23 0 0 0 364.17 370l-105.42-34.26c-20.31-6.5-37.81 15.8-26.51 33.91m148.33-132.23a197.44 197.44 0 0 0-50.41-69.31a22.85 22.85 0 0 0-34 4.4l-62 91.92c-11.9 17.7 4.7 40.61 25.2 34.71L366 268.63a23 23 0 0 0 14.61-31.21zM62.11 30.18a22.86 22.86 0 0 0-9.9 32l104.12 180.44c11.7 20.2 42.61 11.9 42.61-11.4V22.88a22.67 22.67 0 0 0-24.5-22.8a320.37 320.37 0 0 0-112.33 30.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yoast(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M91.3 76h186l-7 18.9h-179c-39.7 0-71.9 31.6-71.9 70.3v205.4c0 35.4 24.9 70.3 84 70.3V460H91.3C41.2 460 0 419.8 0 370.5V165.2C0 115.9 40.7 76 91.3 76m229.1-56h66.5C243.1 398.1 241.2 418.9 202.2 459.3c-20.8 21.6-49.3 31.7-78.3 32.7v-51.1c49.2-7.7 64.6-49.9 64.6-75.3c0-20.1.6-12.6-82.1-223.2h61.4L218.2 299zM448 161.5V460H234c6.6-9.6 10.7-16.3 12.1-19.4h182.5V161.5c0-32.5-17.1-51.9-48.2-62.9l6.7-17.6c41.7 13.6 60.9 43.1 60.9 80.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Youtube(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M549.655 124.083c-6.281-23.65-24.787-42.276-48.284-48.597C458.781 64 288 64 288 64S117.22 64 74.629 75.486c-23.497 6.322-42.003 24.947-48.284 48.597c-11.412 42.867-11.412 132.305-11.412 132.305s0 89.438 11.412 132.305c6.281 23.65 24.787 41.5 48.284 47.821C117.22 448 288 448 288 448s170.78 0 213.371-11.486c23.497-6.321 42.003-24.171 48.284-47.821c11.412-42.867 11.412-132.305 11.412-132.305s0-89.438-11.412-132.305m-317.51 213.508V175.185l142.739 81.205z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YoutubeSquare(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m186.8 202.1l95.2 54.1l-95.2 54.1zM448 80v352c0 26.5-21.5 48-48 48H48c-26.5 0-48-21.5-48-48V80c0-26.5 21.5-48 48-48h352c26.5 0 48 21.5 48 48m-42 176.3s0-59.6-7.6-88.2c-4.2-15.8-16.5-28.2-32.2-32.4C337.9 128 224 128 224 128s-113.9 0-142.2 7.7c-15.7 4.2-28 16.6-32.2 32.4c-7.6 28.5-7.6 88.2-7.6 88.2s0 59.6 7.6 88.2c4.2 15.8 16.5 27.7 32.2 31.9C110.1 384 224 384 224 384s113.9 0 142.2-7.7c15.7-4.2 28-16.1 32.2-31.9c7.6-28.5 7.6-88.1 7.6-88.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Zhihu(children ...ElementRenderer) *FaBrandsIcon {
	return &FaBrandsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M170.54 148.13v217.54l23.43.01l7.71 26.37l42.01-26.37h49.53V148.13zm97.75 193.93h-27.94l-27.9 17.51l-5.08-17.47l-11.9-.04V171.75h72.82zm-118.46-94.39H97.5c1.74-27.1 2.2-51.59 2.2-73.46h51.16s1.97-22.56-8.58-22.31h-88.5c3.49-13.12 7.87-26.66 13.12-40.67c0 0-24.07 0-32.27 21.57c-3.39 8.9-13.21 43.14-30.7 78.12c5.89-.64 25.37-1.18 36.84-22.21c2.11-5.89 2.51-6.66 5.14-14.53h28.87c0 10.5-1.2 66.88-1.68 73.44H20.83c-11.74 0-15.56 23.62-15.56 23.62h65.58C66.45 321.1 42.83 363.12 0 396.34c20.49 5.85 40.91-.93 51-9.9c0 0 22.98-20.9 35.59-69.25l53.96 64.94s7.91-26.89-1.24-39.99c-7.58-8.92-28.06-33.06-36.79-41.81L87.9 311.95c4.36-13.98 6.99-27.55 7.87-40.67h61.65s-.09-23.62-7.59-23.62zm412.02-1.6c20.83-25.64 44.98-58.57 44.98-58.57s-18.65-14.8-27.38-4.06c-6 8.15-36.83 48.2-36.83 48.2zm-150.09-59.09c-9.01-8.25-25.91 2.13-25.91 2.13s39.52 55.04 41.12 57.45l19.46-13.73s-25.67-37.61-34.66-45.86h-.01zM640 258.35c-19.78 0-130.91.93-131.06.93v-101c4.81 0 12.42-.4 22.85-1.2c40.88-2.41 70.13-4 87.77-4.81c0 0 12.22-27.19-.59-33.44c-3.07-1.18-23.17 4.58-23.17 4.58s-165.22 16.49-232.36 18.05c1.6 8.82 7.62 17.08 15.78 19.55c13.31 3.48 22.69 1.7 49.15.89c24.83-1.6 43.68-2.43 56.51-2.43v99.81H351.41s2.82 22.31 25.51 22.85h107.94v70.92c0 13.97-11.19 21.99-24.48 21.12c-14.08.11-26.08-1.15-41.69-1.81c1.99 3.97 6.33 14.39 19.31 21.84c9.88 4.81 16.17 6.57 26.02 6.57c29.56 0 45.67-17.28 44.89-45.31v-73.32h122.36c9.68 0 8.7-23.78 8.7-23.78z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
