package line_md

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "1.0.5"
	hAttr          = "1em"
	viewbox        = "0 0 24 24"
	fill           = "currentColor"
)

type LineMdIcon struct {
	*SVGSVGElement
}

type LineMdIconFn func(children ...ElementRenderer) *LineMdIcon

var IconLookup = map[string]LineMdIconFn{
	"account":                                Account,
	"accountAdd":                             AccountAdd,
	"accountAlert":                           AccountAlert,
	"accountAlertLoop":                       AccountAlertLoop,
	"accountDelete":                          AccountDelete,
	"accountRemove":                          AccountRemove,
	"accountSmall":                           AccountSmall,
	"alert":                                  Alert,
	"alertCircle":                            AlertCircle,
	"alertCircleLoop":                        AlertCircleLoop,
	"alertCircleTwotone":                     AlertCircleTwotone,
	"alertCircleTwotoneLoop":                 AlertCircleTwotoneLoop,
	"alertLoop":                              AlertLoop,
	"alertSquare":                            AlertSquare,
	"alertSquareLoop":                        AlertSquareLoop,
	"alertSquareTwotone":                     AlertSquareTwotone,
	"alertSquareTwotoneLoop":                 AlertSquareTwotoneLoop,
	"alertTwotone":                           AlertTwotone,
	"alertTwotoneLoop":                       AlertTwotoneLoop,
	"alignCenter":                            AlignCenter,
	"alignJustify":                           AlignJustify,
	"alignLeft":                              AlignLeft,
	"arrowAlignCenter":                       ArrowAlignCenter,
	"arrowAlignLeft":                         ArrowAlignLeft,
	"arrowCloseLeft":                         ArrowCloseLeft,
	"arrowLeft":                              ArrowLeft,
	"arrowLeftCircle":                        ArrowLeftCircle,
	"arrowLeftCircleTwotone":                 ArrowLeftCircleTwotone,
	"arrowLeftSquare":                        ArrowLeftSquare,
	"arrowLeftSquareTwotone":                 ArrowLeftSquareTwotone,
	"arrowLongDiagonal":                      ArrowLongDiagonal,
	"arrowOpenLeft":                          ArrowOpenLeft,
	"arrowSmallLeft":                         ArrowSmallLeft,
	"arrowsDiagonal":                         ArrowsDiagonal,
	"arrowsHorizontal":                       ArrowsHorizontal,
	"arrowsHorizontalAlt":                    ArrowsHorizontalAlt,
	"backupRestore":                          BackupRestore,
	"beer":                                   Beer,
	"beerAltFilled":                          BeerAltFilled,
	"beerAltFilledLoop":                      BeerAltFilledLoop,
	"beerAltTwotone":                         BeerAltTwotone,
	"beerAltTwotoneLoop":                     BeerAltTwotoneLoop,
	"beerFilled":                             BeerFilled,
	"beerLoop":                               BeerLoop,
	"beerTwotone":                            BeerTwotone,
	"beerTwotoneLoop":                        BeerTwotoneLoop,
	"bell":                                   Bell,
	"bellAlert":                              BellAlert,
	"bellAlertLoop":                          BellAlertLoop,
	"bellLoop":                               BellLoop,
	"bellTwotone":                            BellTwotone,
	"bellTwotoneAlert":                       BellTwotoneAlert,
	"bellTwotoneAlertLoop":                   BellTwotoneAlertLoop,
	"bellTwotoneLoop":                        BellTwotoneLoop,
	"briefcase":                              Briefcase,
	"briefcaseFilled":                        BriefcaseFilled,
	"briefcaseTwotone":                       BriefcaseTwotone,
	"buyMeAcoffee":                           BuyMeAcoffee,
	"buyMeAcoffeeFilled":                     BuyMeAcoffeeFilled,
	"buyMeAcoffeeTwotone":                    BuyMeAcoffeeTwotone,
	"cake":                                   Cake,
	"cakeTwotone":                            CakeTwotone,
	"calendar":                               Calendar,
	"calendarOut":                            CalendarOut,
	"cancel":                                 Cancel,
	"cancelTwotone":                          CancelTwotone,
	"chat":                                   Chat,
	"chatBubble":                             ChatBubble,
	"chatBubbleFilled":                       ChatBubbleFilled,
	"chatBubbleTwotone":                      ChatBubbleTwotone,
	"chatTwotone":                            ChatTwotone,
	"checkAll":                               CheckAll,
	"checkListThree":                         CheckListThree,
	"checkListThreeFilled":                   CheckListThreeFilled,
	"checkListThreeTwotone":                  CheckListThreeTwotone,
	"chevronDoubleLeft":                      ChevronDoubleLeft,
	"chevronLeft":                            ChevronLeft,
	"chevronLeftCircle":                      ChevronLeftCircle,
	"chevronLeftCircleTwotone":               ChevronLeftCircleTwotone,
	"chevronLeftSquare":                      ChevronLeftSquare,
	"chevronLeftSquareTwotone":               ChevronLeftSquareTwotone,
	"chevronSmallDoubleLeft":                 ChevronSmallDoubleLeft,
	"chevronSmallLeft":                       ChevronSmallLeft,
	"chevronSmallTripleLeft":                 ChevronSmallTripleLeft,
	"chevronTripleLeft":                      ChevronTripleLeft,
	"circle":                                 Circle,
	"circleToConfirmCircleTransition":        CircleToConfirmCircleTransition,
	"circleToConfirmCircleTwotoneTransition": CircleToConfirmCircleTwotoneTransition,
	"circleTwotone":                          CircleTwotone,
	"circleTwotoneToConfirmCircleTwotoneTransition": CircleTwotoneToConfirmCircleTwotoneTransition,
	"clipboard":                           Clipboard,
	"clipboardArrow":                      ClipboardArrow,
	"clipboardArrowTwotone":               ClipboardArrowTwotone,
	"clipboardCheck":                      ClipboardCheck,
	"clipboardCheckToClipboardTransition": ClipboardCheckToClipboardTransition,
	"clipboardCheckTwotone":               ClipboardCheckTwotone,
	"clipboardCheckTwotoneToClipboardTwotoneTransition": ClipboardCheckTwotoneToClipboardTwotoneTransition,
	"clipboardList":                       ClipboardList,
	"clipboardListTwotone":                ClipboardListTwotone,
	"clipboardMinus":                      ClipboardMinus,
	"clipboardMinusTwotone":               ClipboardMinusTwotone,
	"clipboardPlus":                       ClipboardPlus,
	"clipboardPlusTwotone":                ClipboardPlusTwotone,
	"clipboardToClipboardCheckTransition": ClipboardToClipboardCheckTransition,
	"clipboardTwotone":                    ClipboardTwotone,
	"clipboardTwotoneToClipboardTwotoneCheckTransition": ClipboardTwotoneToClipboardTwotoneCheckTransition,
	"close":                                  Close,
	"closeCircle":                            CloseCircle,
	"closeCircleTwotone":                     CloseCircleTwotone,
	"closeSmall":                             CloseSmall,
	"closeToMenuAltTransition":               CloseToMenuAltTransition,
	"closeToMenuTransition":                  CloseToMenuTransition,
	"cloud":                                  Cloud,
	"cloudBracesLoop":                        CloudBracesLoop,
	"cloudDown":                              CloudDown,
	"cloudDownTwotone":                       CloudDownTwotone,
	"cloudDownloadLoop":                      CloudDownloadLoop,
	"cloudDownloadOutlineLoop":               CloudDownloadOutlineLoop,
	"cloudFilled":                            CloudFilled,
	"cloudLoop":                              CloudLoop,
	"cloudOffOutlineLoop":                    CloudOffOutlineLoop,
	"cloudOutlineLoop":                       CloudOutlineLoop,
	"cloudPrintLoop":                         CloudPrintLoop,
	"cloudPrintOutlineLoop":                  CloudPrintOutlineLoop,
	"cloudTagsLoop":                          CloudTagsLoop,
	"cloudTwotone":                           CloudTwotone,
	"cloudUp":                                CloudUp,
	"cloudUpTwotone":                         CloudUpTwotone,
	"cloudUploadLoop":                        CloudUploadLoop,
	"cloudUploadOutlineLoop":                 CloudUploadOutlineLoop,
	"coffee":                                 Coffee,
	"coffeeArrow":                            CoffeeArrow,
	"coffeeArrowFilled":                      CoffeeArrowFilled,
	"coffeeArrowTwotone":                     CoffeeArrowTwotone,
	"coffeeFilled":                           CoffeeFilled,
	"coffeeHalfEmptyTwotoneLoop":             CoffeeHalfEmptyTwotoneLoop,
	"coffeeLoop":                             CoffeeLoop,
	"coffeeTwotone":                          CoffeeTwotone,
	"coffeeTwotoneLoop":                      CoffeeTwotoneLoop,
	"cog":                                    Cog,
	"cogFilled":                              CogFilled,
	"cogFilledLoop":                          CogFilledLoop,
	"cogLoop":                                CogLoop,
	"cogOff":                                 CogOff,
	"cogOffFilled":                           CogOffFilled,
	"cogOffFilledLoop":                       CogOffFilledLoop,
	"cogOffLoop":                             CogOffLoop,
	"compass":                                Compass,
	"compassLoop":                            CompassLoop,
	"compassOff":                             CompassOff,
	"compassTwotone":                         CompassTwotone,
	"compassTwotoneLoop":                     CompassTwotoneLoop,
	"compassTwotoneOff":                      CompassTwotoneOff,
	"computer":                               Computer,
	"computerTwotone":                        ComputerTwotone,
	"confirm":                                Confirm,
	"confirmCircle":                          ConfirmCircle,
	"confirmCircleToCircleTransition":        ConfirmCircleToCircleTransition,
	"confirmCircleTwotone":                   ConfirmCircleTwotone,
	"confirmCircleTwotoneToCircleTransition": ConfirmCircleTwotoneToCircleTransition,
	"confirmCircleTwotoneToCircleTwotoneTransition": ConfirmCircleTwotoneToCircleTwotoneTransition,
	"confirmSquare":                                 ConfirmSquare,
	"confirmSquareToSquareTransition":               ConfirmSquareToSquareTransition,
	"confirmSquareTwotone":                          ConfirmSquareTwotone,
	"confirmSquareTwotoneToSquareTransition":        ConfirmSquareTwotoneToSquareTransition,
	"confirmSquareTwotoneToSquareTwotoneTransition": ConfirmSquareTwotoneToSquareTwotoneTransition,
	"construction":                                  Construction,
	"constructionTwotone":                           ConstructionTwotone,
	"discord":                                       Discord,
	"discordTwotone":                                DiscordTwotone,
	"document":                                      Document,
	"documentAdd":                                   DocumentAdd,
	"documentAddTwotone":                            DocumentAddTwotone,
	"documentCode":                                  DocumentCode,
	"documentCodeTwotone":                           DocumentCodeTwotone,
	"documentList":                                  DocumentList,
	"documentListTwotone":                           DocumentListTwotone,
	"documentRemove":                                DocumentRemove,
	"documentRemoveTwotone":                         DocumentRemoveTwotone,
	"documentReport":                                DocumentReport,
	"documentReportTwotone":                         DocumentReportTwotone,
	"documentTwotone":                               DocumentTwotone,
	"doubleArrowHorizontal":                         DoubleArrowHorizontal,
	"downloadLoop":                                  DownloadLoop,
	"downloadOffLoop":                               DownloadOffLoop,
	"downloadOffOutline":                            DownloadOffOutline,
	"downloadOffOutlineLoop":                        DownloadOffOutlineLoop,
	"downloadOutline":                               DownloadOutline,
	"downloadOutlineLoop":                           DownloadOutlineLoop,
	"downloadingLoop":                               DownloadingLoop,
	"edit":                                          Edit,
	"editTwotone":                                   EditTwotone,
	"editTwotoneFull":                               EditTwotoneFull,
	"email":                                         Email,
	"emailOpened":                                   EmailOpened,
	"emailOpenedTwotone":                            EmailOpenedTwotone,
	"emailOpenedTwotoneAlt":                         EmailOpenedTwotoneAlt,
	"emailTwotone":                                  EmailTwotone,
	"emailTwotoneAlt":                               EmailTwotoneAlt,
	"emojiAngry":                                    EmojiAngry,
	"emojiAngryTwotone":                             EmojiAngryTwotone,
	"emojiFrown":                                    EmojiFrown,
	"emojiFrownOpen":                                EmojiFrownOpen,
	"emojiFrownOpenTwotone":                         EmojiFrownOpenTwotone,
	"emojiFrownTwotone":                             EmojiFrownTwotone,
	"emojiGrin":                                     EmojiGrin,
	"emojiGrinTwotone":                              EmojiGrinTwotone,
	"emojiNeutral":                                  EmojiNeutral,
	"emojiNeutralTwotone":                           EmojiNeutralTwotone,
	"emojiSmile":                                    EmojiSmile,
	"emojiSmileTwotone":                             EmojiSmileTwotone,
	"emojiSmileWink":                                EmojiSmileWink,
	"emojiSmileWinkTwotone":                         EmojiSmileWinkTwotone,
	"externalLink":                                  ExternalLink,
	"externalLinkRounded":                           ExternalLinkRounded,
	"facebook":                                      Facebook,
	"filter":                                        Filter,
	"filterFilled":                                  FilterFilled,
	"filterTwotone":                                 FilterTwotone,
	"forkLeft":                                      ForkLeft,
	"gauge":                                         Gauge,
	"gaugeEmpty":                                    GaugeEmpty,
	"gaugeFull":                                     GaugeFull,
	"gaugeLoop":                                     GaugeLoop,
	"gaugeLow":                                      GaugeLow,
	"github":                                        Github,
	"githubLoop":                                    GithubLoop,
	"githubTwotone":                                 GithubTwotone,
	"gridThree":                                     GridThree,
	"gridThreeFilled":                               GridThreeFilled,
	"gridThreeTwotone":                              GridThreeTwotone,
	"hash":                                          Hash,
	"hashSmall":                                     HashSmall,
	"heart":                                         Heart,
	"heartFilled":                                   HeartFilled,
	"heartFilledHalf":                               HeartFilledHalf,
	"heartHalf":                                     HeartHalf,
	"heartHalfFilled":                               HeartHalfFilled,
	"heartHalfTwotone":                              HeartHalfTwotone,
	"heartTwotone":                                  HeartTwotone,
	"heartTwotoneHalf":                              HeartTwotoneHalf,
	"heartTwotoneHalfFilled":                        HeartTwotoneHalfFilled,
	"home":                                          Home,
	"homeMd":                                        HomeMd,
	"homeMdTwotone":                                 HomeMdTwotone,
	"homeMdTwotoneAlt":                              HomeMdTwotoneAlt,
	"homeSimple":                                    HomeSimple,
	"homeSimpleFilled":                              HomeSimpleFilled,
	"homeSimpleTwotone":                             HomeSimpleTwotone,
	"homeTwotone":                                   HomeTwotone,
	"homeTwotoneAlt":                                HomeTwotoneAlt,
	"iconifyOne":                                    IconifyOne,
	"iconifyTwo":                                    IconifyTwo,
	"image":                                         Image,
	"imageTwotone":                                  ImageTwotone,
	"instagram":                                     Instagram,
	"laptop":                                        Laptop,
	"laptopTwotone":                                 LaptopTwotone,
	"lightDark":                                     LightDark,
	"lightDarkLoop":                                 LightDarkLoop,
	"lightbulb":                                     Lightbulb,
	"lightbulbFilled":                               LightbulbFilled,
	"lightbulbOff":                                  LightbulbOff,
	"lightbulbOffFilled":                            LightbulbOffFilled,
	"lightbulbOffFilledLoop":                        LightbulbOffFilledLoop,
	"lightbulbOffLoop":                              LightbulbOffLoop,
	"lightbulbOffTwotone":                           LightbulbOffTwotone,
	"lightbulbOffTwotoneLoop":                       LightbulbOffTwotoneLoop,
	"lightbulbTwotone":                              LightbulbTwotone,
	"linkedin":                                      Linkedin,
	"list":                                          List,
	"listIndented":                                  ListIndented,
	"listThree":                                     ListThree,
	"listThreeFilled":                               ListThreeFilled,
	"listThreeTwotone":                              ListThreeTwotone,
	"loadingAltLoop":                                LoadingAltLoop,
	"loadingLoop":                                   LoadingLoop,
	"loadingTwotoneLoop":                            LoadingTwotoneLoop,
	"logIn":                                         LogIn,
	"logOut":                                        LogOut,
	"login":                                         Login,
	"logout":                                        Logout,
	"mapMarker":                                     MapMarker,
	"mapMarkerAlt":                                  MapMarkerAlt,
	"mapMarkerAltFilled":                            MapMarkerAltFilled,
	"mapMarkerAltTwotone":                           MapMarkerAltTwotone,
	"mapMarkerFilled":                               MapMarkerFilled,
	"mapMarkerMultipleAlt":                          MapMarkerMultipleAlt,
	"mapMarkerMultipleAltFilled":                    MapMarkerMultipleAltFilled,
	"mapMarkerMultipleAltTwotone":                   MapMarkerMultipleAltTwotone,
	"mapMarkerOff":                                  MapMarkerOff,
	"mapMarkerOffAlt":                               MapMarkerOffAlt,
	"mapMarkerOffAltFilled":                         MapMarkerOffAltFilled,
	"mapMarkerOffAltFilledLoop":                     MapMarkerOffAltFilledLoop,
	"mapMarkerOffAltLoop":                           MapMarkerOffAltLoop,
	"mapMarkerOffAltTwotone":                        MapMarkerOffAltTwotone,
	"mapMarkerOffAltTwotoneLoop":                    MapMarkerOffAltTwotoneLoop,
	"mapMarkerOffFilled":                            MapMarkerOffFilled,
	"mapMarkerOffFilledLoop":                        MapMarkerOffFilledLoop,
	"mapMarkerOffLoop":                              MapMarkerOffLoop,
	"mapMarkerOffTwotone":                           MapMarkerOffTwotone,
	"mapMarkerOffTwotoneLoop":                       MapMarkerOffTwotoneLoop,
	"mapMarkerTwotone":                              MapMarkerTwotone,
	"marker":                                        Marker,
	"markerFilled":                                  MarkerFilled,
	"markerTwotone":                                 MarkerTwotone,
	"mastodon":                                      Mastodon,
	"mastodonFilled":                                MastodonFilled,
	"mastodonTwotone":                               MastodonTwotone,
	"medicalServices":                               MedicalServices,
	"medicalServicesFilled":                         MedicalServicesFilled,
	"medicalServicesTwotone":                        MedicalServicesTwotone,
	"menu":                                          Menu,
	"menuFoldLeft":                                  MenuFoldLeft,
	"menuFoldRight":                                 MenuFoldRight,
	"menuToCloseAltTransition":                      MenuToCloseAltTransition,
	"menuToCloseTransition":                         MenuToCloseTransition,
	"menuUnfoldLeft":                                MenuUnfoldLeft,
	"menuUnfoldRight":                               MenuUnfoldRight,
	"minus":                                         Minus,
	"minusCircle":                                   MinusCircle,
	"minusCircleTwotone":                            MinusCircleTwotone,
	"minusSquare":                                   MinusSquare,
	"minusSquareTwotone":                            MinusSquareTwotone,
	"moon":                                          Moon,
	"moonAltLoop":                                   MoonAltLoop,
	"moonAltToSunnyOutlineLoopTransition":           MoonAltToSunnyOutlineLoopTransition,
	"moonFilled":                                    MoonFilled,
	"moonFilledAltLoop":                             MoonFilledAltLoop,
	"moonFilledAltToSunnyFilledLoopTransition": MoonFilledAltToSunnyFilledLoopTransition,
	"moonFilledLoop":                         MoonFilledLoop,
	"moonFilledToSunnyFilledLoopTransition":  MoonFilledToSunnyFilledLoopTransition,
	"moonFilledToSunnyFilledTransition":      MoonFilledToSunnyFilledTransition,
	"moonLoop":                               MoonLoop,
	"moonRisingAltLoop":                      MoonRisingAltLoop,
	"moonRisingFilledAltLoop":                MoonRisingFilledAltLoop,
	"moonRisingFilledLoop":                   MoonRisingFilledLoop,
	"moonRisingLoop":                         MoonRisingLoop,
	"moonRisingTwotoneAltLoop":               MoonRisingTwotoneAltLoop,
	"moonRisingTwotoneLoop":                  MoonRisingTwotoneLoop,
	"moonSimple":                             MoonSimple,
	"moonSimpleFilled":                       MoonSimpleFilled,
	"moonSimpleTwotone":                      MoonSimpleTwotone,
	"moonToSunnyOutlineLoopTransition":       MoonToSunnyOutlineLoopTransition,
	"moonToSunnyOutlineTransition":           MoonToSunnyOutlineTransition,
	"moonTwotone":                            MoonTwotone,
	"moonTwotoneAltLoop":                     MoonTwotoneAltLoop,
	"moonTwotoneLoop":                        MoonTwotoneLoop,
	"myLocation":                             MyLocation,
	"myLocationLoop":                         MyLocationLoop,
	"myLocationOff":                          MyLocationOff,
	"myLocationOffLoop":                      MyLocationOffLoop,
	"navigationLeftDown":                     NavigationLeftDown,
	"navigationLeftUp":                       NavigationLeftUp,
	"navigationRightDown":                    NavigationRightDown,
	"navigationRightUp":                      NavigationRightUp,
	"paintDrop":                              PaintDrop,
	"paintDropFilled":                        PaintDropFilled,
	"paintDropHalfFilled":                    PaintDropHalfFilled,
	"paintDropHalfFilledTwotone":             PaintDropHalfFilledTwotone,
	"paintDropHalfTwotone":                   PaintDropHalfTwotone,
	"paintDropTwotone":                       PaintDropTwotone,
	"patreon":                                Patreon,
	"pause":                                  Pause,
	"pauseToPlayFilledTransition":            PauseToPlayFilledTransition,
	"pauseToPlayTransition":                  PauseToPlayTransition,
	"peertube":                               Peertube,
	"peertubeAlt":                            PeertubeAlt,
	"pencil":                                 Pencil,
	"pencilTwotone":                          PencilTwotone,
	"pencilTwotoneAlt":                       PencilTwotoneAlt,
	"person":                                 Person,
	"personAdd":                              PersonAdd,
	"personAddFilled":                        PersonAddFilled,
	"personAddTwotone":                       PersonAddTwotone,
	"personFilled":                           PersonFilled,
	"personOff":                              PersonOff,
	"personOffFilled":                        PersonOffFilled,
	"personOffFilledLoop":                    PersonOffFilledLoop,
	"personOffLoop":                          PersonOffLoop,
	"personOffTwotone":                       PersonOffTwotone,
	"personOffTwotoneLoop":                   PersonOffTwotoneLoop,
	"personRemove":                           PersonRemove,
	"personRemoveFilled":                     PersonRemoveFilled,
	"personRemoveTwotone":                    PersonRemoveTwotone,
	"personSearch":                           PersonSearch,
	"personSearchFilled":                     PersonSearchFilled,
	"personSearchTwotone":                    PersonSearchTwotone,
	"personTwotone":                          PersonTwotone,
	"phone":                                  Phone,
	"phoneAdd":                               PhoneAdd,
	"phoneAddTwotone":                        PhoneAddTwotone,
	"phoneCall":                              PhoneCall,
	"phoneCallLoop":                          PhoneCallLoop,
	"phoneCallTwotone":                       PhoneCallTwotone,
	"phoneCallTwotoneLoop":                   PhoneCallTwotoneLoop,
	"phoneIncoming":                          PhoneIncoming,
	"phoneIncomingTwotone":                   PhoneIncomingTwotone,
	"phoneOff":                               PhoneOff,
	"phoneOffLoop":                           PhoneOffLoop,
	"phoneOffTwotone":                        PhoneOffTwotone,
	"phoneOffTwotoneLoop":                    PhoneOffTwotoneLoop,
	"phoneOutgoing":                          PhoneOutgoing,
	"phoneOutgoingTwotone":                   PhoneOutgoingTwotone,
	"phoneRemove":                            PhoneRemove,
	"phoneRemoveTwotone":                     PhoneRemoveTwotone,
	"phoneTwotone":                           PhoneTwotone,
	"pixelfed":                               Pixelfed,
	"pixelfedFilled":                         PixelfedFilled,
	"pixelfedTwotone":                        PixelfedTwotone,
	"play":                                   Play,
	"playFilled":                             PlayFilled,
	"playFilledToPauseTransition":            PlayFilledToPauseTransition,
	"playToPauseTransition":                  PlayToPauseTransition,
	"playTwotone":                            PlayTwotone,
	"pleroma":                                Pleroma,
	"plus":                                   Plus,
	"plusCircle":                             PlusCircle,
	"plusCircleTwotone":                      PlusCircleTwotone,
	"plusSquare":                             PlusSquare,
	"plusSquareTwotone":                      PlusSquareTwotone,
	"question":                               Question,
	"questionCircle":                         QuestionCircle,
	"questionCircleTwotone":                  QuestionCircleTwotone,
	"questionSquare":                         QuestionSquare,
	"questionSquareTwotone":                  QuestionSquareTwotone,
	"reddit":                                 Reddit,
	"redditCircle":                           RedditCircle,
	"redditCircleLoop":                       RedditCircleLoop,
	"redditLoop":                             RedditLoop,
	"remove":                                 Remove,
	"rotateNinety":                           RotateNinety,
	"rotateOneHundredEighty":                 RotateOneHundredEighty,
	"rotateTwoHundredSeventy":                RotateTwoHundredSeventy,
	"roundRampLeft":                          RoundRampLeft,
	"roundThreeHundredSixty":                 RoundThreeHundredSixty,
	"roundaboutLeft":                         RoundaboutLeft,
	"rss":                                    Rss,
	"search":                                 Search,
	"searchFilled":                           SearchFilled,
	"searchTwotone":                          SearchTwotone,
	"speed":                                  Speed,
	"speedLoop":                              SpeedLoop,
	"speedometer":                            Speedometer,
	"speedometerLoop":                        SpeedometerLoop,
	"square":                                 Square,
	"squareToConfirmSquareTransition":        SquareToConfirmSquareTransition,
	"squareToConfirmSquareTwotoneTransition": SquareToConfirmSquareTwotoneTransition,
	"squareTwotone":                          SquareTwotone,
	"squareTwotoneToConfirmSquareTwotoneTransition": SquareTwotoneToConfirmSquareTwotoneTransition,
	"star":                     Star,
	"starAlt":                  StarAlt,
	"starAltFilled":            StarAltFilled,
	"starAltTwotone":           StarAltTwotone,
	"starFilled":               StarFilled,
	"starFilledHalf":           StarFilledHalf,
	"starHalf":                 StarHalf,
	"starHalfFilled":           StarHalfFilled,
	"starHalfTwotone":          StarHalfTwotone,
	"starPulsatingFilledLoop":  StarPulsatingFilledLoop,
	"starPulsatingLoop":        StarPulsatingLoop,
	"starPulsatingTwotoneLoop": StarPulsatingTwotoneLoop,
	"starTwotone":              StarTwotone,
	"starTwotoneHalf":          StarTwotoneHalf,
	"sunRisingFilledLoop":      SunRisingFilledLoop,
	"sunRisingLoop":            SunRisingLoop,
	"sunRisingTwotoneLoop":     SunRisingTwotoneLoop,
	"sunnyFilled":              SunnyFilled,
	"sunnyFilledLoop":          SunnyFilledLoop,
	"sunnyFilledLoopToMoonAltFilledLoopTransition":     SunnyFilledLoopToMoonAltFilledLoopTransition,
	"sunnyFilledLoopToMoonFilledLoopTransition":        SunnyFilledLoopToMoonFilledLoopTransition,
	"sunnyFilledLoopToMoonFilledTransition":            SunnyFilledLoopToMoonFilledTransition,
	"sunnyOutline":                                     SunnyOutline,
	"sunnyOutlineLoop":                                 SunnyOutlineLoop,
	"sunnyOutlineToMoonAltLoopTransition":              SunnyOutlineToMoonAltLoopTransition,
	"sunnyOutlineToMoonLoopTransition":                 SunnyOutlineToMoonLoopTransition,
	"sunnyOutlineToMoonTransition":                     SunnyOutlineToMoonTransition,
	"sunnyOutlineTwotone":                              SunnyOutlineTwotone,
	"sunnyOutlineTwotoneLoop":                          SunnyOutlineTwotoneLoop,
	"switch":                                           Switch,
	"switchFilled":                                     SwitchFilled,
	"switchFilledToSwitchOffFilledTransition":          SwitchFilledToSwitchOffFilledTransition,
	"switchOff":                                        SwitchOff,
	"switchOffFilled":                                  SwitchOffFilled,
	"switchOffFilledToSwitchFilledTransition":          SwitchOffFilledToSwitchFilledTransition,
	"switchOffToSwitchTransition":                      SwitchOffToSwitchTransition,
	"switchToSwitchOffTransition":                      SwitchToSwitchOffTransition,
	"telegram":                                         Telegram,
	"textBox":                                          TextBox,
	"textBoxMultiple":                                  TextBoxMultiple,
	"textBoxMultipleToTextBoxTransition":               TextBoxMultipleToTextBoxTransition,
	"textBoxMultipleTwotone":                           TextBoxMultipleTwotone,
	"textBoxMultipleTwotoneToTextBoxTwotoneTransition": TextBoxMultipleTwotoneToTextBoxTwotoneTransition,
	"textBoxToTextBoxMultipleTransition":               TextBoxToTextBoxMultipleTransition,
	"textBoxTwotone":                                   TextBoxTwotone,
	"textBoxTwotoneToTextBoxMultipleTwotoneTransition": TextBoxTwotoneToTextBoxMultipleTwotoneTransition,
	"thumbsDown":                                       ThumbsDown,
	"thumbsDownTwotone":                                ThumbsDownTwotone,
	"thumbsUp":                                         ThumbsUp,
	"thumbsUpTwotone":                                  ThumbsUpTwotone,
	"tiktok":                                           Tiktok,
	"turnLeft":                                         TurnLeft,
	"turnSharpLeft":                                    TurnSharpLeft,
	"turnSlightLeft":                                   TurnSlightLeft,
	"twitter":                                          Twitter,
	"twitterTwotone":                                   TwitterTwotone,
	"twitterX":                                         TwitterX,
	"twitterXalt":                                      TwitterXalt,
	"uturnLeft":                                        UturnLeft,
	"uploadLoop":                                       UploadLoop,
	"uploadOffLoop":                                    UploadOffLoop,
	"uploadOffOutline":                                 UploadOffOutline,
	"uploadOffOutlineLoop":                             UploadOffOutlineLoop,
	"uploadOutline":                                    UploadOutline,
	"uploadOutlineLoop":                                UploadOutlineLoop,
	"uploadingLoop":                                    UploadingLoop,
	"valignBaseline":                                   ValignBaseline,
	"valignBaselineTwotone":                            ValignBaselineTwotone,
	"valignBottom":                                     ValignBottom,
	"valignBottomTwotone":                              ValignBottomTwotone,
	"valignMiddle":                                     ValignMiddle,
	"valignMiddleTwotone":                              ValignMiddleTwotone,
	"valignTop":                                        ValignTop,
	"valignTopTwotone":                                 ValignTopTwotone,
	"watch":                                            Watch,
	"watchLoop":                                        WatchLoop,
	"watchOff":                                         WatchOff,
	"watchOffLoop":                                     WatchOffLoop,
	"watchOffTwotone":                                  WatchOffTwotone,
	"watchOffTwotoneLoop":                              WatchOffTwotoneLoop,
	"watchTwotone":                                     WatchTwotone,
	"watchTwotoneLoop":                                 WatchTwotoneLoop,
	"weatherCloudyLoop":                                WeatherCloudyLoop,
	"youtube":                                          Youtube,
	"youtubeFilled":                                    YoutubeFilled,
	"youtubeTwotone":                                   YoutubeTwotone,
}

func Account(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="28" stroke-dashoffset="28" stroke-linecap="round" stroke-width="2"><path d="M4 21V20C4 16.6863 6.68629 14 10 14H14C17.3137 14 20 16.6863 20 20V21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="28;0"/></path><path d="M12 11C9.79086 11 8 9.20914 8 7C8 4.79086 9.79086 3 12 3C14.2091 3 16 4.79086 16 7C16 9.20914 14.2091 11 12 11Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="28;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccountAdd(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M3 21V20C3 17.7909 4.79086 16 7 16H11C13.2091 16 15 17.7909 15 20V21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="20" stroke-dashoffset="20" d="M9 13C7.34315 13 6 11.6569 6 10C6 8.34315 7.34315 7 9 7C10.6569 7 12 8.34315 12 10C12 11.6569 10.6569 13 9 13Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="20;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M15 6H21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M18 3V9"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccountAlert(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M5 21V20C5 17.7909 6.79086 16 9 16H13C15.2091 16 17 17.7909 17 20V21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="20" stroke-dashoffset="20" d="M11 13C9.34315 13 8 11.6569 8 10C8 8.34315 9.34315 7 11 7C12.6569 7 14 8.34315 14 10C14 11.6569 12.6569 13 11 13Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="20;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M20 3V7"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="6;0"/></path></g><circle cx="20" cy="11" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.4s" values="0;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccountAlertLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M5 21V20C5 17.7909 6.79086 16 9 16H13C15.2091 16 17 17.7909 17 20V21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="20" stroke-dashoffset="20" d="M11 13C9.34315 13 8 11.6569 8 10C8 8.34315 9.34315 7 11 7C12.6569 7 14 8.34315 14 10C14 11.6569 12.6569 13 11 13Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="20;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M20 3V7"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="6;0"/><animate attributeName="stroke-width" begin="1.5s" dur="3s" keyTimes="0;0.1;0.2;0.3;1" repeatCount="indefinite" values="2;3;3;2;2"/></path></g><circle cx="20" cy="11" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.4s" values="0;1"/><animate attributeName="r" begin="1.8s" dur="3s" keyTimes="0;0.1;0.2;0.3;1" repeatCount="indefinite" values="1;2;2;1;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccountDelete(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M3 21V20C3 17.7909 4.79086 16 7 16H11C13.2091 16 15 17.7909 15 20V21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="20" stroke-dashoffset="20" d="M9 13C7.34315 13 6 11.6569 6 10C6 8.34315 7.34315 7 9 7C10.6569 7 12 8.34315 12 10C12 11.6569 10.6569 13 9 13Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="20;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M15 3L21 9"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="10;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M21 3L15 9"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="10;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccountRemove(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M3 21V20C3 17.7909 4.79086 16 7 16H11C13.2091 16 15 17.7909 15 20V21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="20" stroke-dashoffset="20" d="M9 13C7.34315 13 6 11.6569 6 10C6 8.34315 7.34315 7 9 7C10.6569 7 12 8.34315 12 10C12 11.6569 10.6569 13 9 13Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="20;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M15 6H21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccountSmall(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="20" stroke-dashoffset="20" stroke-linecap="round" stroke-width="2"><path d="M6 19V18C6 15.7909 7.79086 14 10 14H14C16.2091 14 18 15.7909 18 18V19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path d="M12 11C10.3431 11 9 9.65685 9 8C9 6.34315 10.3431 5 12 5C13.6569 5 15 6.34315 15 8C15 9.65685 13.6569 11 12 11Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="20;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alert(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M12 3L21 20H3L12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 10V14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/></path></g><circle cx="12" cy="17" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.4s" values="0;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertCircle(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M12 7V13"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/></path></g><circle cx="12" cy="17" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.4s" values="0;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertCircleLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M12 7V13"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/><animate attributeName="stroke-width" begin="1s" dur="3s" keyTimes="0;0.1;0.2;0.3;1" repeatCount="indefinite" values="2;3;3;2;2"/></path></g><circle cx="12" cy="17" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.4s" values="0;1"/><animate attributeName="r" begin="1.3s" dur="3s" keyTimes="0;0.1;0.2;0.3;1" repeatCount="indefinite" values="1;2;2;1;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertCircleTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="8" stroke-dashoffset="8" d="M12 7V13"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/></path></g><circle cx="12" cy="17" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.4s" values="0;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertCircleTwotoneLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="8" stroke-dashoffset="8" d="M12 7V13"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/><animate attributeName="stroke-width" begin="1s" dur="3s" keyTimes="0;0.1;0.2;0.3;1" repeatCount="indefinite" values="2;3;3;2;2"/></path></g><circle cx="12" cy="17" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.4s" values="0;1"/><animate attributeName="r" begin="1.3s" dur="3s" keyTimes="0;0.1;0.2;0.3;1" repeatCount="indefinite" values="1;2;2;1;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M12 3L21 20H3L12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 10V14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/><animate attributeName="stroke-width" begin="1s" dur="3s" keyTimes="0;0.1;0.2;0.3;1" repeatCount="indefinite" values="2;3;3;2;2"/></path></g><circle cx="12" cy="17" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.4s" values="0;1"/><animate attributeName="r" begin="1.3s" dur="3s" keyTimes="0;0.1;0.2;0.3;1" repeatCount="indefinite" values="1;2;2;1;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertSquare(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M12 4H5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="64;128"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M12 7V13"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/></path></g><circle cx="12" cy="17" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.4s" values="0;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertSquareLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M12 4H5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="64;128"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M12 7V13"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/><animate attributeName="stroke-width" begin="1s" dur="3s" keyTimes="0;0.1;0.2;0.3;1" repeatCount="indefinite" values="2;3;3;2;2"/></path></g><circle cx="12" cy="17" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.4s" values="0;1"/><animate attributeName="r" begin="1.3s" dur="3s" keyTimes="0;0.1;0.2;0.3;1" repeatCount="indefinite" values="1;2;2;1;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertSquareTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="64" stroke-dashoffset="64" d="M12 4H5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="64;128"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="8" stroke-dashoffset="8" d="M12 7V13"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/></path></g><circle cx="12" cy="17" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.4s" values="0;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertSquareTwotoneLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="64" stroke-dashoffset="64" d="M12 4H5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="64;128"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="8" stroke-dashoffset="8" d="M12 7V13"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/><animate attributeName="stroke-width" begin="1s" dur="3s" keyTimes="0;0.1;0.2;0.3;1" repeatCount="indefinite" values="2;3;3;2;2"/></path></g><circle cx="12" cy="17" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.4s" values="0;1"/><animate attributeName="r" begin="1.3s" dur="3s" keyTimes="0;0.1;0.2;0.3;1" repeatCount="indefinite" values="1;2;2;1;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M12 3L21 20H3L12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="6" stroke-dashoffset="6" d="M12 10V14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/></path></g><circle cx="12" cy="17" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.4s" values="0;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertTwotoneLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M12 3L21 20H3L12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="6" stroke-dashoffset="6" d="M12 10V14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/><animate attributeName="stroke-width" begin="1s" dur="3s" keyTimes="0;0.1;0.2;0.3;1" repeatCount="indefinite" values="2;3;3;2;2"/></path></g><circle cx="12" cy="17" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.4s" values="0;1"/><animate attributeName="r" begin="1.3s" dur="3s" keyTimes="0;0.1;0.2;0.3;1" repeatCount="indefinite" values="1;2;2;1;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenter(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="9" stroke-dashoffset="9" d="M12 5H19M12 5H5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="9;0"/></path><path stroke-dasharray="7" stroke-dashoffset="7" d="M12 10H17M12 10H7"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="7;0"/></path><path stroke-dasharray="11" stroke-dashoffset="11" d="M12 15H21M12 15H3"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="11;0"/></path><path stroke-dasharray="9" stroke-dashoffset="9" d="M12 20H19M12 20H5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="9;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignJustify(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="8" stroke-dashoffset="8" stroke-linecap="round" stroke-width="2"><path d="M12 5H18M12 5H6"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="8;0"/></path><path d="M12 10H18M12 10H6"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="8;0"/></path><path d="M12 15H18M12 15H6"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="8;0"/></path><path d="M12 20H18M12 20H6"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeft(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="15" stroke-dashoffset="15" d="M4 5H17"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="15;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M4 10H14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M4 15H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="18;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M4 20H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="15;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowAlignCenter(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M12 3V21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="20;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M22 12H15.5M2 12H8.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M15 12L18 15M9 12L6 15M15 12L18 9M9 12L6 9"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowAlignLeft(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M3 3V21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="20;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M21 12H7.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="15;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M7 12L11 16M7 12L11 8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCloseLeft(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M3 3V21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="20;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M21 12H7.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="15;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M7 12L14 19M7 12L14 5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="12;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M21 12H3.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="20;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M3 12L10 19M3 12L10 5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.2s" values="12;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftCircle(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9" stroke-dasharray="60" stroke-dashoffset="60"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></circle><path stroke-dasharray="12" stroke-dashoffset="12" d="M17 12H7.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M7 12L11 16M7 12L11 8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftCircleTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9" fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.15s" values="0;0.3"/></circle><path stroke-dasharray="12" stroke-dashoffset="12" d="M17 12H7.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M7 12L11 16M7 12L11 8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftSquare(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M20 12v7a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="64;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M17 12H7.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M7 12L11 16M7 12L11 8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftSquareTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="64" stroke-dashoffset="64" d="M20 12v7a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="64;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.15s" values="0;0.3"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M17 12H7.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M7 12L11 16M7 12L11 8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongDiagonal(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M12 12L3.5 20.5M12 12L20.5 3.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="14;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M21 3H13M3 21V13M21 3V11M3 21H11"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowOpenLeft(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M21 3V21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="20;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M17 12H3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="15;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M3 12L10 19M3 12L10 5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="12;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSmallLeft(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M19 12H5.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="14;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M5 12L10 17M5 12L10 7"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsDiagonal(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="10" stroke-dashoffset="10" d="M10 14L3.5 20.5M14 10L20.5 3.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="10;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M21 3H15M3 21V15M21 3V9M3 21H9"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsHorizontal(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="12" stroke-dashoffset="12" d="M15 7H3.5M9 17H20.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="12;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M3 7L7 11M3 7L7 3M21 17L17 21M21 17L17 13"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsHorizontalAlt(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="12" stroke-dashoffset="12" d="M21 7H10.5M3 17H13.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="12;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M10 7L14 11M10 7L14 3M14 17L10 21M14 17L10 13"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BackupRestore(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="48" stroke-dashoffset="48" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.25 14C5.14 17.45 8.27 20 12 20C16.42 20 20 16.42 20 12C20 7.58 16.42 4 12 4C9.61 4 7.47 5.05 6 6.71L4 9"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="48;0"/></path><g fill="currentColor"><path fill-opacity="0" d="M3.25 10H3.89645C4.11917 10 4.23071 9.73071 4.07322 9.57322L3.42678 8.92678C3.26929 8.76929 3 8.88083 3 9.10355V9.75C3 9.88807 3.11193 10 3.25 10Z"><set attributeName="fill-opacity" begin="0.6s" to="1"/><animate fill="freeze" attributeName="d" begin="0.6s" dur="0.2s" values="M3.25 10H3.89645C4.11917 10 4.23071 9.73071 4.07322 9.57322L3.42678 8.92678C3.26929 8.76929 3 8.88083 3 9.10355V9.75C3 9.88807 3.11193 10 3.25 10Z;M3.5 10H7.79289C8.23835 10 8.46143 9.46143 8.14645 9.14645L3.85355 4.85355C3.53857 4.53857 3 4.76165 3 5.20711V9.5C3 9.77614 3.22386 10 3.5 10Z"/></path><circle cx="12" cy="12" r="0"><animate fill="freeze" attributeName="r" begin="0.8s" dur="0.2s" values="0;2"/></circle></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Beer(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M18 3L16 21H7L5 3z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="60;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M6 7.67C6.6 7.3 7.22 7 8 7C10 7 11 9 13 9C14.64 9 15.6 7.66 17 7.17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BeerAltFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdBeerAltFilled0"><g transform="translate(0 16)"><path fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 7.67h1C6.6 7.3 7.22 7 8 7C10 7 11 9 13 9C14.64 9 15.6 7.66 17 7.17h1"/><path fill="#fff" d="M5 7.67h1C6.6 7.3 7.22 7 8 7C10 7 11 9 13 9C14.64 9 15.6 7.66 17 7.17h1V24H6z"/><animateMotion fill="freeze" begin="0.6s" calcMode="linear" dur="0.6s" path="M0 0v-16"/></g></mask><path fill="none" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 3L16 21H7L5 3z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="60;0"/></path><path fill="currentColor" d="M18 3L16 21H7L5 3z" mask="url(#lineMdBeerAltFilled0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BeerAltFilledLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdBeerAltFilledLoop0"><g transform="translate(0 16)"><g><path stroke="#fff" stroke-width="2" d="M18 7C16 7 15 9 13 9C11 9 10 7 8 7C6 7 5 9 3 9C1 9 0 7 -2 7C-4 7 -5 9 -7 9"/><path fill="#fff" d="M18 7C16 7 15 9 13 9C11 9 10 7 8 7C6 7 5 9 3 9C1 9 0 7 -2 7C-4 7 -5 9 -7 9V24H18z"/><animateMotion calcMode="linear" dur="3s" path="M0 0h10" repeatCount="indefinite"/></g><animateMotion fill="freeze" begin="0.6s" calcMode="linear" dur="0.6s" path="M0 0v-16"/></g></mask><path fill="none" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 3L16 21H7L5 3z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="60;0"/></path><path fill="currentColor" d="M18 3L16 21H7L5 3z" mask="url(#lineMdBeerAltFilledLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BeerAltTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdBeerAltTwotone0"><g transform="translate(0 16)"><path fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 7.67h1C6.6 7.3 7.22 7 8 7C10 7 11 9 13 9C14.64 9 15.6 7.66 17 7.17h1"/><path fill="#fff" fill-opacity=".3" d="M5 7.67h1C6.6 7.3 7.22 7 8 7C10 7 11 9 13 9C14.64 9 15.6 7.66 17 7.17h1V24H6z"/><animateMotion fill="freeze" begin="0.6s" calcMode="linear" dur="0.6s" path="M0 0v-16"/></g></mask><path fill="none" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 3L16 21H7L5 3z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="60;0"/></path><path fill="currentColor" d="M18 3L16 21H7L5 3z" mask="url(#lineMdBeerAltTwotone0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BeerAltTwotoneLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdBeerAltTwotoneLoop0"><g transform="translate(0 16)"><g><path stroke="#fff" stroke-width="2" d="M18 7C16 7 15 9 13 9C11 9 10 7 8 7C6 7 5 9 3 9C1 9 0 7 -2 7C-4 7 -5 9 -7 9"/><path fill="#fff" fill-opacity=".3" d="M18 7C16 7 15 9 13 9C11 9 10 7 8 7C6 7 5 9 3 9C1 9 0 7 -2 7C-4 7 -5 9 -7 9V24H18z"/><animateMotion calcMode="linear" dur="3s" path="M0 0h10" repeatCount="indefinite"/></g><animateMotion fill="freeze" begin="0.6s" calcMode="linear" dur="0.6s" path="M0 0v-16"/></g></mask><path fill="none" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 3L16 21H7L5 3z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="60;0"/></path><path fill="currentColor" d="M18 3L16 21H7L5 3z" mask="url(#lineMdBeerAltTwotoneLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BeerFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M18 3L16 21H7L5 3z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="60;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M6 7.67C6.6 7.3 7.22 7 8 7C10 7 11 9 13 9C14.64 9 15.6 7.66 17 7.17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path></g><path fill="currentColor" fill-opacity="0" d="M17 8L16 21H7L6 8z"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.5s" values="0;1"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BeerLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdBeerLoop0"><path stroke="#fff" stroke-width="2" d="M18 7C16 7 15 9 13 9C11 9 10 7 8 7C6 7 5 9 3 9C1 9 0 7 -2 7C-4 7 -5 9 -7 9" opacity="0"><animateMotion calcMode="linear" dur="3s" path="M0 0h10" repeatCount="indefinite"/><animate fill="freeze" attributeName="opacity" begin="0.6s" dur="0.5s" values="0;1"/></path></mask><path fill="none" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 3L16 21H7L5 3z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="60;0"/></path><path fill="currentColor" d="M18 3L16 21H7L5 3z" mask="url(#lineMdBeerLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BeerTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M18 3L16 21H7L5 3z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="60;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M6 7.67C6.6 7.3 7.22 7 8 7C10 7 11 9 13 9C14.64 9 15.6 7.66 17 7.17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path></g><path fill="currentColor" fill-opacity="0" d="M17 8L16 21H7L6 8z"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BeerTwotoneLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdBeerTwotoneLoop0"><path stroke="#fff" stroke-width="2" d="M18 7C16 7 15 9 13 9C11 9 10 7 8 7C6 7 5 9 3 9C1 9 0 7 -2 7C-4 7 -5 9 -7 9" opacity="0"><animateMotion calcMode="linear" dur="3s" path="M0 0h10" repeatCount="indefinite"/><animate fill="freeze" attributeName="opacity" begin="0.6s" dur="0.5s" values="0;1"/></path></mask><path fill="currentColor" fill-opacity="0" d="M17 8L16 21H7L6 8z"><animate fill="freeze" attributeName="fill-opacity" begin="1.1s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 3L16 21H7L5 3z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="60;0"/></path><path fill="currentColor" d="M18 3L16 21H7L5 3z" mask="url(#lineMdBeerTwotoneLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="4" stroke-dashoffset="4" d="M12 3V5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="4;0"/></path><path stroke-dasharray="28" stroke-dashoffset="28" d="M12 5C8.68629 5 6 7.68629 6 11L6 17C5 17 4 18 4 19H12M12 5C15.3137 5 18 7.68629 18 11L18 17C19 17 20 18 20 19H12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.4s" values="28;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M10 20C10 21.1046 10.8954 22 12 22C13.1046 22 14 21.1046 14 20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellAlert(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="4" stroke-dashoffset="4" d="M12 3V5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="4;0"/></path><path stroke-dasharray="28" stroke-dashoffset="28" d="M12 5C8.68629 5 6 7.68629 6 11L6 17C5 17 4 18 4 19H12M12 5C15.3137 5 18 7.68629 18 11L18 17C19 17 20 18 20 19H12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.4s" values="28;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M10 20C10 21.1046 10.8954 22 12 22C13.1046 22 14 21.1046 14 20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M22 6v4" opacity="0"><set attributeName="opacity" begin="1s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path></g><circle cx="22" cy="14" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.4s" values="0;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellAlertLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><g><path stroke-dasharray="4" stroke-dashoffset="4" d="M12 3V5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="4;0"/></path><path stroke-dasharray="28" stroke-dashoffset="28" d="M12 5C8.68629 5 6 7.68629 6 11L6 17C5 17 4 18 4 19H12M12 5C15.3137 5 18 7.68629 18 11L18 17C19 17 20 18 20 19H12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.4s" values="28;0"/></path><animateTransform attributeName="transform" begin="0.8s" dur="6s" keyTimes="0;0.05;0.15;0.2;1" repeatCount="indefinite" type="rotate" values="0 12 3;3 12 3;-3 12 3;0 12 3;0 12 3"/></g><path stroke-dasharray="8" stroke-dashoffset="8" d="M10 20C10 21.1046 10.8954 22 12 22C13.1046 22 14 21.1046 14 20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/><animateTransform attributeName="transform" begin="1s" dur="6s" keyTimes="0;0.05;0.15;0.2;1" repeatCount="indefinite" type="rotate" values="0 12 8;6 12 8;-6 12 8;0 12 8;0 12 8"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M22 6v4" opacity="0"><set attributeName="opacity" begin="1s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/><animate attributeName="stroke-width" begin="1.7s" dur="3s" keyTimes="0;0.1;0.2;0.3;1" repeatCount="indefinite" values="2;3;3;2;2"/></path></g><circle cx="22" cy="14" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.4s" values="0;1"/><animate attributeName="r" begin="2s" dur="3s" keyTimes="0;0.1;0.2;0.3;1" repeatCount="indefinite" values="1;2;2;1;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><g><path stroke-dasharray="4" stroke-dashoffset="4" d="M12 3V5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="4;0"/></path><path stroke-dasharray="28" stroke-dashoffset="28" d="M12 5C8.68629 5 6 7.68629 6 11L6 17C5 17 4 18 4 19H12M12 5C15.3137 5 18 7.68629 18 11L18 17C19 17 20 18 20 19H12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.4s" values="28;0"/></path><animateTransform attributeName="transform" begin="0.8s" dur="6s" keyTimes="0;0.05;0.15;0.2;1" repeatCount="indefinite" type="rotate" values="0 12 3;3 12 3;-3 12 3;0 12 3;0 12 3"/></g><path stroke-dasharray="8" stroke-dashoffset="8" d="M10 20C10 21.1046 10.8954 22 12 22C13.1046 22 14 21.1046 14 20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/><animateTransform attributeName="transform" begin="1s" dur="6s" keyTimes="0;0.05;0.15;0.2;1" repeatCount="indefinite" type="rotate" values="0 12 8;6 12 8;-6 12 8;0 12 8;0 12 8"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="4" stroke-dashoffset="4" d="M12 3V5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="4;0"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="28" stroke-dashoffset="28" d="M12 5C8.68629 5 6 7.68629 6 11L6 17C5 17 4 18 4 19H12M12 5C15.3137 5 18 7.68629 18 11L18 17C19 17 20 18 20 19H12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.4s" values="28;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M10 20C10 21.1046 10.8954 22 12 22C13.1046 22 14 21.1046 14 20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellTwotoneAlert(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="4" stroke-dashoffset="4" d="M12 3V5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="4;0"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="28" stroke-dashoffset="28" d="M12 5C8.68629 5 6 7.68629 6 11L6 17C5 17 4 18 4 19H12M12 5C15.3137 5 18 7.68629 18 11L18 17C19 17 20 18 20 19H12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.4s" values="28;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M10 20C10 21.1046 10.8954 22 12 22C13.1046 22 14 21.1046 14 20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M22 6v4" opacity="0"><set attributeName="opacity" begin="1s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path></g><circle cx="22" cy="14" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.4s" values="0;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellTwotoneAlertLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><g><path stroke-dasharray="4" stroke-dashoffset="4" d="M12 3V5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="4;0"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="28" stroke-dashoffset="28" d="M12 5C8.68629 5 6 7.68629 6 11L6 17C5 17 4 18 4 19H12M12 5C15.3137 5 18 7.68629 18 11L18 17C19 17 20 18 20 19H12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.4s" values="28;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><animateTransform attributeName="transform" begin="0.8s" dur="6s" keyTimes="0;0.05;0.15;0.2;1" repeatCount="indefinite" type="rotate" values="0 12 3;3 12 3;-3 12 3;0 12 3;0 12 3"/></g><path stroke-dasharray="8" stroke-dashoffset="8" d="M10 20C10 21.1046 10.8954 22 12 22C13.1046 22 14 21.1046 14 20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/><animateTransform attributeName="transform" begin="1s" dur="6s" keyTimes="0;0.05;0.15;0.2;1" repeatCount="indefinite" type="rotate" values="0 12 8;6 12 8;-6 12 8;0 12 8;0 12 8"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M22 6v4" opacity="0"><set attributeName="opacity" begin="1s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/><animate attributeName="stroke-width" begin="1.7s" dur="3s" keyTimes="0;0.1;0.2;0.3;1" repeatCount="indefinite" values="2;3;3;2;2"/></path></g><circle cx="22" cy="14" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.4s" values="0;1"/><animate attributeName="r" begin="2s" dur="3s" keyTimes="0;0.1;0.2;0.3;1" repeatCount="indefinite" values="1;2;2;1;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellTwotoneLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><g><path stroke-dasharray="4" stroke-dashoffset="4" d="M12 3V5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="4;0"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="28" stroke-dashoffset="28" d="M12 5C8.68629 5 6 7.68629 6 11L6 17C5 17 4 18 4 19H12M12 5C15.3137 5 18 7.68629 18 11L18 17C19 17 20 18 20 19H12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.4s" values="28;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><animateTransform attributeName="transform" begin="0.8s" dur="6s" keyTimes="0;0.05;0.15;0.2;1" repeatCount="indefinite" type="rotate" values="0 12 3;3 12 3;-3 12 3;0 12 3;0 12 3"/></g><path stroke-dasharray="8" stroke-dashoffset="8" d="M10 20C10 21.1046 10.8954 22 12 22C13.1046 22 14 21.1046 14 20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/><animateTransform attributeName="transform" begin="1s" dur="6s" keyTimes="0;0.05;0.15;0.2;1" repeatCount="indefinite" type="rotate" values="0 12 8;6 12 8;-6 12 8;0 12 8;0 12 8"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Briefcase(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="16" stroke-dashoffset="16" d="M15 7V4C15 3.44772 14.5523 3 14 3H10C9.44772 3 9 3.44772 9 4V7" opacity="0"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="16;32"/><set attributeName="opacity" begin="0.7s" to="1"/></path><path stroke-dasharray="64" stroke-dashoffset="64" d="M9 7H20C20.5523 7 21 7.44772 21 8V19C21 19.5523 20.5523 20 20 20H4C3.44772 20 3 19.5523 3 19V8C3 7.44772 3.44772 7 4 7z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BriefcaseFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="none" stroke-dasharray="16" stroke-dashoffset="16" d="M15 7V4C15 3.44772 14.5523 3 14 3H10C9.44772 3 9 3.44772 9 4V7" opacity="0"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="16;32"/><set attributeName="opacity" begin="0.7s" to="1"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="64" stroke-dashoffset="64" d="M9 7H20C20.5523 7 21 7.44772 21 8V19C21 19.5523 20.5523 20 20 20H4C3.44772 20 3 19.5523 3 19V8C3 7.44772 3.44772 7 4 7z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.5s" values="0;1"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BriefcaseTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="none" stroke-dasharray="16" stroke-dashoffset="16" d="M15 7V4C15 3.44772 14.5523 3 14 3H10C9.44772 3 9 3.44772 9 4V7" opacity="0"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="16;32"/><set attributeName="opacity" begin="0.7s" to="1"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="64" stroke-dashoffset="64" d="M9 7H20C20.5523 7 21 7.44772 21 8V19C21 19.5523 20.5523 20 20 20H4C3.44772 20 3 19.5523 3 19V8C3 7.44772 3.44772 7 4 7z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.15s" values="0;0.3"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuyMeAcoffee(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdBuyMeACoffee0"><path fill="#fff" d="M5 6C5 4 7 6 11.5 6C16 6 19 4 19 6L19 7C19 8.5 17 9 12.5 9C8 9 5 9 5 7L5 6Z"/></mask><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="32" stroke-dashoffset="32" d="M7.5 10.5C7.5 10.5 8.33 17.43 8.5 19C8.67 20.57 10 21 11 21L13 21C14.5 21 15.875 19.86 16 19C16.125 18.14 17 7 17 7"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="32;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M16.5 6C16.5 3.5 14 3 12 3C10 3 9.1 3.43 8 4"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="12;24"/></path></g><rect width="16" height="5" x="20" y="4" fill="currentColor" mask="url(#lineMdBuyMeACoffee0)"><animate fill="freeze" attributeName="x" begin="0.4s" dur="0.4s" values="20;4"/></rect>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuyMeAcoffeeFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdBuyMeACoffeeFilled0"><path fill="#fff" d="M5 6C5 4 7 6 11.5 6C16 6 19 4 19 6L19 7C19 8.5 17 9 12.5 9C8 9 5 9 5 7L5 6Z"/></mask><mask id="lineMdBuyMeACoffeeFilled1"><path fill="#fff" d="M10.125 18.15C10.04 17.29 9.4 11.98 9.4 11.98C9.4 11.98 11.34 12.31 12.5 11.73C13.66 11.16 14.98 11 14.98 11C14.98 11 14.4 17.96 14.35 18.46C14.3 18.96 13.45 19.3 12.95 19.3L11.23 19.3C10.73 19.3 10.21 19 10.125 18.15Z"/></mask><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="32" stroke-dashoffset="32" d="M7.5 10.5C7.5 10.5 8.33 17.43 8.5 19C8.67 20.57 10 21 11 21L13 21C14.5 21 15.875 19.86 16 19C16.125 18.14 17 7 17 7"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="32;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M16.5 6C16.5 3.5 14 3 12 3C10 3 9.1 3.43 8 4"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="12;24"/></path></g><rect width="16" height="5" x="20" y="4" fill="currentColor" mask="url(#lineMdBuyMeACoffeeFilled0)"><animate fill="freeze" attributeName="x" begin="0.4s" dur="0.4s" values="20;4"/></rect><rect width="8" height="10" x="8" y="20" fill="currentColor" mask="url(#lineMdBuyMeACoffeeFilled1)"><animate fill="freeze" attributeName="y" begin="1s" dur="0.4s" values="20;10"/></rect>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuyMeAcoffeeTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdBuyMeACoffeeTwotone0"><path fill="#fff" d="M5 6C5 4 7 6 11.5 6C16 6 19 4 19 6L19 7C19 8.5 17 9 12.5 9C8 9 5 9 5 7L5 6Z"/></mask><mask id="lineMdBuyMeACoffeeTwotone1"><path fill="#fff" d="M10.125 18.15C10.04 17.29 9.4 11.98 9.4 11.98C9.4 11.98 11.34 12.31 12.5 11.73C13.66 11.16 14.98 11 14.98 11C14.98 11 14.4 17.96 14.35 18.46C14.3 18.96 13.45 19.3 12.95 19.3L11.23 19.3C10.73 19.3 10.21 19 10.125 18.15Z"/></mask><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="32" stroke-dashoffset="32" d="M7.5 10.5C7.5 10.5 8.33 17.43 8.5 19C8.67 20.57 10 21 11 21L13 21C14.5 21 15.875 19.86 16 19C16.125 18.14 17 7 17 7"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="32;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M16.5 6C16.5 3.5 14 3 12 3C10 3 9.1 3.43 8 4"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="12;24"/></path></g><rect width="16" height="5" x="20" y="4" fill="currentColor" mask="url(#lineMdBuyMeACoffeeTwotone0)"><animate fill="freeze" attributeName="x" begin="0.4s" dur="0.4s" values="20;4"/></rect><rect width="8" height="10" x="8" y="20" fill="currentColor" fill-opacity=".3" mask="url(#lineMdBuyMeACoffeeTwotone1)"><animate fill="freeze" attributeName="y" begin="1s" dur="0.4s" values="20;10"/></rect>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cake(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-dasharray="28" stroke-dashoffset="28" d="M12 10H18C19.1046 10 20 10.8954 20 12V21H12"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="28;56"/></path><path stroke-dasharray="28" stroke-dashoffset="28" d="M12 21H4V12C4 10.8954 4.89543 10 6 10H12"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="28;0"/></path><path stroke-dasharray="4" stroke-dashoffset="4" stroke-linecap="round" stroke-linejoin="round" d="M12 10V8" opacity="0"><set attributeName="opacity" begin="0.4s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="4;0"/></path><path stroke-dasharray="20" stroke-dashoffset="20" d="M4 16H5C7 16 8.5 14 8.5 14C8.5 14 10 16 12 16C14 16 15.5 14 15.5 14C15.5 14 17 16 19 16H20" opacity="0"><set attributeName="opacity" begin="0.4s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="20;0"/></path></g><path fill="currentColor" d="M14 4C14 5.10457 13.1046 6 12 6C10.8954 6 10 5.10457 10 4C10 2.89543 12 0 12 0C12 0 14 2.89543 14 4Z" opacity="0"><set attributeName="opacity" begin="0.6s" to="1"/><animate fill="freeze" attributeName="d" begin="0.6s" dur="0.2s" values="M13 5C13 5.5 12.5 6 12 6C11.5 6 11 5.5 11 5C11 4.5 12 4 12 4C12 4 13 4.5 13 5Z;M14 4C14 5.10457 13.1046 6 12 6C10.8954 6 10 5.10457 10 4C10 2.89543 12 0 12 0C12 0 14 2.89543 14 4Z"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CakeTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-dasharray="28" stroke-dashoffset="28" d="M12 10H18C19.1046 10 20 10.8954 20 12V21H12"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="28;56"/></path><path stroke-dasharray="28" stroke-dashoffset="28" d="M12 21H4V12C4 10.8954 4.89543 10 6 10H12"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="28;0"/></path><path stroke-dasharray="4" stroke-dashoffset="4" stroke-linecap="round" stroke-linejoin="round" d="M12 10V8" opacity="0"><set attributeName="opacity" begin="0.4s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="4;0"/></path><path stroke-dasharray="20" stroke-dashoffset="20" d="M4 16H5C7 16 8.5 14 8.5 14C8.5 14 10 16 12 16C14 16 15.5 14 15.5 14C15.5 14 17 16 19 16H20" opacity="0"><set attributeName="opacity" begin="0.4s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="20;0"/></path></g><g fill="currentColor"><path d="M14 4C14 5.10457 13.1046 6 12 6C10.8954 6 10 5.10457 10 4C10 2.89543 12 0 12 0C12 0 14 2.89543 14 4Z" opacity="0"><set attributeName="opacity" begin="0.6s" to="1"/><animate fill="freeze" attributeName="d" begin="0.6s" dur="0.2s" values="M13 5C13 5.5 12.5 6 12 6C11.5 6 11 5.5 11 5C11 4.5 12 4 12 4C12 4 13 4.5 13 5Z;M14 4C14 5.10457 13.1046 6 12 6C10.8954 6 10 5.10457 10 4C10 2.89543 12 0 12 0C12 0 14 2.89543 14 4Z"/></path><path d="M15.5 14C15.5 14 14 16 12 16C10 16 8.5 14 8.5 14C8.5 14 7 16 5 16H4V21H20V16H19C17 16 15.5 14 15.5 14Z" opacity="0"><animate fill="freeze" attributeName="opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="16" x="4" y="4" stroke-dasharray="64" stroke-dashoffset="64" rx="1"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="64;0"/></rect><path stroke-dasharray="6" stroke-dashoffset="6" d="M7 4V2M17 4V2"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="6;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M7 11H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="9" stroke-dashoffset="9" d="M7 15H14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="9;0"/></path></g><rect width="14" height="0" x="5" y="5" fill="currentColor"><animate fill="freeze" attributeName="height" begin="0.5s" dur="0.2s" values="0;3"/></rect>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarOut(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="16" x="4" y="4" stroke-dasharray="64" rx="1"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.5s" values="0;64"/></rect><path stroke-dasharray="6" d="M7 4V2M17 4V2"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="0;6"/></path><path stroke-dasharray="12" d="M7 11H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="0;12"/></path><path stroke-dasharray="9" d="M7 15H14"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="0;9"/></path></g><rect width="14" height="3" x="5" y="5" fill="currentColor"><animate fill="freeze" attributeName="height" begin="0.5s" dur="0.2s" values="3;0"/></rect>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cancel(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M5.63604 5.63603C9.15076 2.12131 14.8492 2.12131 18.364 5.63603C21.8787 9.15075 21.8787 14.8492 18.364 18.364C14.8492 21.8787 9.15076 21.8787 5.63604 18.364C2.12132 14.8492 2.12132 9.15075 5.63604 5.63603Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M6 6L18 18"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="18;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CancelTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M5.63604 5.63603C9.15076 2.12131 14.8492 2.12131 18.364 5.63603C21.8787 9.15075 21.8787 14.8492 18.364 18.364C14.8492 21.8787 9.15076 21.8787 5.63604 18.364C2.12132 14.8492 2.12132 9.15075 5.63604 5.63603Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="18" stroke-dashoffset="18" d="M6 6L18 18"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="18;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chat(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="68" stroke-dashoffset="68" d="M3 19.5V4C3 3.44772 3.44772 3 4 3H20C20.5523 3 21 3.44772 21 4V16C21 16.5523 20.5523 17 20 17H5.5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="68;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M8 7h8" opacity="0"><set attributeName="opacity" begin="0.7s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="10;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M8 10h8" opacity="0"><set attributeName="opacity" begin="0.8s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="10;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M8 13h4" opacity="0"><set attributeName="opacity" begin="0.9s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubble(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="68" stroke-dashoffset="68" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 19.5V4C3 3.44772 3.44772 3 4 3H20C20.5523 3 21 3.44772 21 4V16C21 16.5523 20.5523 17 20 17H5.5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="68;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill-opacity="0" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" stroke-dasharray="68" stroke-dashoffset="68" d="M3 19.5V4C3 3.44772 3.44772 3 4 3H20C20.5523 3 21 3.44772 21 4V16C21 16.5523 20.5523 17 20 17H5.5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="68;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.5s" values="0;1"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill-opacity="0" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" stroke-dasharray="68" stroke-dashoffset="68" d="M3 19.5V4C3 3.44772 3.44772 3 4 3H20C20.5523 3 21 3.44772 21 4V16C21 16.5523 20.5523 17 20 17H5.5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="68;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.15s" values="0;0.3"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="68" stroke-dashoffset="68" d="M3 19.5V4C3 3.44772 3.44772 3 4 3H20C20.5523 3 21 3.44772 21 4V16C21 16.5523 20.5523 17 20 17H5.5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="68;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.15s" values="0;0.3"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M8 7h8" opacity="0"><set attributeName="opacity" begin="0.8s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="10;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M8 10h8" opacity="0"><set attributeName="opacity" begin="0.9s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="10;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M8 13h4" opacity="0"><set attributeName="opacity" begin="1s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckAll(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><mask id="lineMdCheckAll0"><g fill="none" stroke="#fff" stroke-dasharray="22" stroke-dashoffset="22" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 13.5l4 4l10.75 -10.75"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="22;0"/></path><path stroke="#000" stroke-width="4" d="M7.5 13.5l4 4l10.75 -10.75" opacity="0"><set attributeName="opacity" begin="0.4s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.4s" values="22;0"/></path><path d="M7.5 13.5l4 4l10.75 -10.75" opacity="0"><set attributeName="opacity" begin="0.4s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.4s" values="22;0"/></path></g></mask></defs><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCheckAll0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckListThree(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-dasharray="10" stroke-dashoffset="10" stroke-width="2"><path d="M3 5L5 7L9 3"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/></path><path d="M3 12L5 14L9 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="10;0"/></path><path d="M3 19L5 21L9 17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="10;0"/></path></g><g stroke-dasharray="22" stroke-dashoffset="22"><rect width="9" height="3" x="11.5" y="3.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.5s" values="22;0"/></rect><rect width="9" height="3" x="11.5" y="10.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.5s" values="22;0"/></rect><rect width="9" height="3" x="11.5" y="17.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.5s" values="22;0"/></rect></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckListThreeFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g fill="none" stroke-dasharray="10" stroke-dashoffset="10" stroke-width="2"><path d="M3 5L5 7L9 3"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/></path><path d="M3 12L5 14L9 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="10;0"/></path><path d="M3 19L5 21L9 17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="10;0"/></path></g><g fill="currentColor" fill-opacity="0" stroke-dasharray="22" stroke-dashoffset="22"><rect width="9" height="3" x="11.5" y="3.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.5s" values="22;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.7s" dur="0.5s" values="0;1"/></rect><rect width="9" height="3" x="11.5" y="10.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.5s" values="22;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.9s" dur="0.5s" values="0;1"/></rect><rect width="9" height="3" x="11.5" y="17.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.5s" values="22;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.1s" dur="0.5s" values="0;1"/></rect></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckListThreeTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g fill="none" stroke-dasharray="10" stroke-dashoffset="10" stroke-width="2"><path d="M3 5L5 7L9 3"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/></path><path d="M3 12L5 14L9 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="10;0"/></path><path d="M3 19L5 21L9 17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="10;0"/></path></g><g fill="currentColor" fill-opacity="0" stroke-dasharray="22" stroke-dashoffset="22"><rect width="9" height="3" x="11.5" y="3.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.5s" values="22;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.7s" dur="0.15s" values="0;0.3"/></rect><rect width="9" height="3" x="11.5" y="10.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.5s" values="22;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.9s" dur="0.15s" values="0;0.3"/></rect><rect width="9" height="3" x="11.5" y="17.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.5s" values="22;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.1s" dur="0.15s" values="0;0.3"/></rect></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleLeft(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="10" stroke-dashoffset="10" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 12L12 5M5 12L12 19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="10;0"/></path><path d="M11 12L18 5M11 12L18 19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.3s" values="10;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeft(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="10" stroke-dashoffset="10" stroke-linecap="round" stroke-width="2" d="M8 12L15 5M8 12L15 19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="10;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeftCircle(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M10 12L13 9M10 12L13 15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeftCircleTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M10 12L13 9M10 12L13 15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeftSquare(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M4 12V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="64;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M10 12L13 9M10 12L13 15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeftSquareTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="64" stroke-dashoffset="64" d="M4 12V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="64;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M10 12L13 9M10 12L13 15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronSmallDoubleLeft(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="8" stroke-dashoffset="8" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 12L17 7M12 12L17 17"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="8;0"/></path><path d="M6 12L11 7M6 12L11 17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.3s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronSmallLeft(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path stroke="currentColor" stroke-dasharray="8" stroke-dashoffset="8" stroke-linecap="round" stroke-width="2" d="M9 12L14 7M9 12L14 17" fill="currentColor"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="8;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronSmallTripleLeft(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="8" stroke-dashoffset="8" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 12L19 7M14 12L19 17"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="8;0"/></path><path d="M9 12L14 7M9 12L14 17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.3s" values="8;0"/></path><path d="M4 12L9 7M4 12L9 17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.3s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronTripleLeft(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="10" stroke-dashoffset="10" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 12L9 5M2 12L9 19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="10;0"/></path><path d="M8 12L15 5M8 12L15 19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.3s" values="10;0"/></path><path d="M14 12L21 5M14 12L21 19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.3s" values="10;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circle(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleToConfirmCircleTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path stroke-dasharray="14" stroke-dashoffset="14" d="M8 12L11 15L16 10"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="14;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleToConfirmCircleTwotoneTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.2s" dur="0.15s" values="0;0.3"/></circle><path fill="none" stroke-dasharray="14" stroke-dashoffset="14" d="M8 12L11 15L16 10"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="14;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.15s" values="0;0.3"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleTwotoneToConfirmCircleTwotoneTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9" fill="currentColor" fill-opacity=".3"/><path fill="none" stroke-dasharray="14" stroke-dashoffset="14" d="M8 12L11 15L16 10"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="14;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clipboard(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="66" stroke-dashoffset="66" stroke-width="2" d="M12 3H19V21H5V3H12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="66;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M14.5 3.5V6.5H9.5V3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="12;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardArrow(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="16" stroke-dashoffset="16" d="M12 3H19V11"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="16;0"/></path><path stroke-dasharray="44" stroke-dashoffset="44" d="M19 17V21H5V3H12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.4s" values="44;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M21 14H12.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="10;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 14L15 17M12 14L15 11"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="6;0"/></path></g><path stroke-dasharray="12" stroke-dashoffset="12" d="M14.5 3.5V6.5H9.5V3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="12;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardArrowTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M6 4H10V6H14V4H18V10H20V18H18V20H6V4Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.4s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="16" stroke-dashoffset="16" d="M12 3H19V11"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="16;0"/></path><path stroke-dasharray="44" stroke-dashoffset="44" d="M19 17V21H5V3H12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.4s" values="44;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M21 14H12.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="10;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 14L15 17M12 14L15 11"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="6;0"/></path></g><path stroke-dasharray="12" stroke-dashoffset="12" d="M14.5 3.5V6.5H9.5V3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="12;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardCheck(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="66" stroke-dashoffset="66" d="M12 3H19V21H5V3H12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="66;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M9 13L11 15L15 11"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="10;0"/></path></g><path stroke-dasharray="12" stroke-dashoffset="12" d="M14.5 3.5V6.5H9.5V3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="12;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardCheckToClipboardTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path d="M12 3H19V21H5V3H12Z"/><path stroke-dasharray="10" d="M9 13L11 15L15 11"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;10"/></path></g><path d="M14.5 3.5V6.5H9.5V3.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardCheckTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M6 4H10V6H14V4H18V20H6V4Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="66" stroke-dashoffset="66" d="M12 3H19V21H5V3H12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="66;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M9 13L11 15L15 11"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="10;0"/></path></g><path stroke-dasharray="12" stroke-dashoffset="12" d="M14.5 3.5V6.5H9.5V3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="12;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardCheckTwotoneToClipboardTwotoneTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".3" d="M6 4H10V6H14V4H18V20H6V4Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path d="M12 3H19V21H5V3H12Z"/><path stroke-dasharray="10" d="M9 13L11 15L15 11"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;10"/></path></g><path d="M14.5 3.5V6.5H9.5V3.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardList(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="66" stroke-dashoffset="66" d="M12 3H19V21H5V3H12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="66;0"/></path><path stroke-dasharray="5" stroke-dashoffset="5" d="M9 10H12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="5;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M9 13H14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="6;0"/></path><path stroke-dasharray="7" stroke-dashoffset="7" d="M9 16H15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.4s" dur="0.2s" values="7;0"/></path></g><path stroke-dasharray="12" stroke-dashoffset="12" d="M14.5 3.5V6.5H9.5V3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="12;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardListTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M6 4H10V6H14V4H18V20H6V4Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.6s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="66" stroke-dashoffset="66" d="M12 3H19V21H5V3H12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="66;0"/></path><path stroke-dasharray="5" stroke-dashoffset="5" d="M9 10H12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="5;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M9 13H14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="6;0"/></path><path stroke-dasharray="7" stroke-dashoffset="7" d="M9 16H15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.4s" dur="0.2s" values="7;0"/></path></g><path stroke-dasharray="12" stroke-dashoffset="12" d="M14.5 3.5V6.5H9.5V3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="12;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardMinus(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="66" stroke-dashoffset="66" d="M12 3H19V21H5V3H12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="66;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M9 13H15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path></g><path stroke-dasharray="12" stroke-dashoffset="12" d="M14.5 3.5V6.5H9.5V3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="12;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardMinusTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M6 4H10V6H14V4H18V20H6V4Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="66" stroke-dashoffset="66" d="M12 3H19V21H5V3H12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="66;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M9 13H15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path></g><path stroke-dasharray="12" stroke-dashoffset="12" d="M14.5 3.5V6.5H9.5V3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="12;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardPlus(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="66" stroke-dashoffset="66" d="M12 3H19V21H5V3H12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="66;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M9 13H15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M12 10V16"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="8;0"/></path></g><path stroke-dasharray="12" stroke-dashoffset="12" d="M14.5 3.5V6.5H9.5V3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="12;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardPlusTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M6 4H10V6H14V4H18V20H6V4Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.4s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="66" stroke-dashoffset="66" d="M12 3H19V21H5V3H12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="66;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M9 13H15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M12 10V16"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="8;0"/></path></g><path stroke-dasharray="12" stroke-dashoffset="12" d="M14.5 3.5V6.5H9.5V3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="12;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardToClipboardCheckTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path d="M12 3H19V21H5V3H12Z"/><path stroke-dasharray="10" stroke-dashoffset="10" d="M9 13L11 15L15 11"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/></path></g><path d="M14.5 3.5V6.5H9.5V3.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M6 4H10V6H14V4H18V20H6V4Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="66" stroke-dashoffset="66" stroke-width="2" d="M12 3H19V21H5V3H12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="66;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M14.5 3.5V6.5H9.5V3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="12;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardTwotoneToClipboardTwotoneCheckTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".3" d="M6 4H10V6H14V4H18V20H6V4Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path d="M12 3H19V21H5V3H12Z"/><path stroke-dasharray="10" stroke-dashoffset="10" d="M9 13L11 15L15 11"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/></path></g><path d="M14.5 3.5V6.5H9.5V3.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Close(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="12" stroke-dashoffset="12" stroke-linecap="round" stroke-width="2" d="M12 12L19 19M12 12L5 5M12 12L5 19M12 12L19 5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="12;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseCircle(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M12 12L16 16M12 12L8 8M12 12L8 16M12 12L16 8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseCircleTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="8" stroke-dashoffset="8" d="M12 12L16 16M12 12L8 8M12 12L8 16M12 12L16 8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseSmall(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="16" stroke-dashoffset="16" stroke-linecap="round" stroke-width="2"><path d="M7 7L17 17"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="16;0"/></path><path d="M17 7L7 17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.4s" values="16;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseToMenuAltTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M5 5L19 19"><animate fill="freeze" attributeName="d" dur="0.4s" values="M5 5L19 19;M5 5L19 5"/></path><path d="M12 12H12" opacity="0"><animate fill="freeze" attributeName="d" begin="0.2s" dur="0.4s" values="M12 12H12;M5 12H19"/><set attributeName="opacity" begin="0.2s" to="1"/></path><path d="M5 19L19 5"><animate fill="freeze" attributeName="d" dur="0.4s" values="M5 19L19 5;M5 19L19 19"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseToMenuTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M5 5L12 12L19 5"><animate fill="freeze" attributeName="d" dur="0.4s" values="M5 5L12 12L19 5;M5 5L12 5L19 5"/></path><path d="M12 12H12"><animate fill="freeze" attributeName="d" dur="0.4s" values="M12 12H12;M5 12H19"/></path><path d="M5 19L12 12L19 19"><animate fill="freeze" attributeName="d" dur="0.4s" values="M5 19L12 12L19 19;M5 19L12 19L19 19"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M12 19C12 19 9.5 19 7 19C4.5 19 3 17 3 15C3 13 4.5 11 7 11C8 11 8.5 11.5 8.5 11.5M12 19H17C19.5 19 21 17 21 15C21 13 19.5 11 17 11C16 11 15.5 11.5 15.5 11.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="9" stroke-dashoffset="9" d="M17 11C17 11 17 10.5 17 10C17 7.5 15 5 12 5M7 11V10C7 7.5 9 5 12 5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.3s" values="9;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudBracesLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdCloudBracesLoop0"><g fill="#fff"><circle cx="12" cy="10" r="6"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;13;12"/></circle><rect width="9" height="7" x="8" y="13"/><rect width="15" height="12" x="1" y="8" rx="6"><animate attributeName="x" dur="15s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="13" height="10" x="10" y="10" rx="5"><animate attributeName="x" dur="15s" repeatCount="indefinite" values="10;11;10;9;10"/></rect></g><path d="M18.5 12H18a1 1 0 0 1-1-1v-1a2 2 0 0 0-2-2h-1.5v2H15v1a2 2 0 0 0 2 2 2 2 0 0 0-2 2v1h-1.5v2H15a2 2 0 0 0 2-2v-1a1 1 0 0 1 1-1h.5v-2Z"><animateMotion calcMode="linear" dur="6s" keyPoints="0;0.25;0.5;0.75;1" keyTimes="0;0.1;0.5;0.8;1" path="M0 0h-1h2z" repeatCount="indefinite"/></path><path d="M5.5 12v2H6a1 1 0 0 1 1 1v1a2 2 0 0 0 2 2h1.5v-2H9v-1a2 2 0 0 0-2-2 2 2 0 0 0 2-2v-1h1.5V8H9a2 2 0 0 0-2 2v1a1 1 0 0 1-1 1h-.5Z"><animateMotion calcMode="linear" dur="6s" keyPoints="0;0.25;0.5;0.75;1" keyTimes="0;0.1;0.5;0.8;1" path="M0 0h1h-2z" repeatCount="indefinite"/></path></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudBracesLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDown(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="16" stroke-dashoffset="16" d="M8 18H7C4.5 18 3 16 3 14C3 12 4.5 10 7 10C8 10 8.5 10.5 8.5 10.5M16 18H17C19.5 18 21 16 21 14C21 12 19.5 10 17 10C16 10 15.5 10.5 15.5 10.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="16;0"/></path><path stroke-dasharray="9" stroke-dashoffset="9" d="M17 10C17 10 17 9.5 17 9C17 6.5 15 4 12 4M7 10V9C7 6.5 9 4 12 4"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.3s" values="9;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M12 15V21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="3" stroke-dashoffset="3" d="M12 22L14 20M12 22L10 20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="3;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDownTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M9 6L12 4L15 6L21 14L17 19H7L3 14L9 6Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.3s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="16" stroke-dashoffset="16" d="M8 18H7C4.5 18 3 16 3 14C3 12 4.5 10 7 10C8 10 8.5 10.5 8.5 10.5M16 18H17C19.5 18 21 16 21 14C21 12 19.5 10 17 10C16 10 15.5 10.5 15.5 10.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="16;0"/></path><path stroke-dasharray="9" stroke-dashoffset="9" d="M17 10C17 10 17 9.5 17 9C17 6.5 15 4 12 4M7 10V9C7 6.5 9 4 12 4"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.3s" values="9;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M12 15V21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="3" stroke-dashoffset="3" d="M12 22L14 20M12 22L10 20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="3;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDownloadLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdCloudDownloadLoop0"><g fill="#fff"><circle cx="12" cy="10" r="6"/><rect width="9" height="7" x="8" y="13"/><rect width="15" height="12" x="1" y="8" rx="6"><animate attributeName="x" dur="15s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="13" height="10" x="10" y="10" rx="5"><animate attributeName="x" dur="19s" repeatCount="indefinite" values="10;9;10;11;10"/></rect></g><rect width="4" height="5" x="10" y="9"/><path d="M12 18L17 13H7L12 18Z"><animateMotion calcMode="linear" dur="1.5s" keyPoints="0;0.25;0.5;0.75;1" keyTimes="0;0.1;0.5;0.8;1" path="M0 0v1v-2z" repeatCount="indefinite"/></path></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudDownloadLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDownloadOutlineLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdCloudDownloadOutlineLoop0"><g fill="#fff"><circle cx="12" cy="10" r="6"/><rect width="9" height="8" x="8" y="12"/><rect width="17" height="12" x="1" y="8" rx="6"><animate attributeName="x" dur="19s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="17" height="10" x="6" y="10" rx="5"><animate attributeName="x" dur="23s" repeatCount="indefinite" values="6;5;6;7;6"/></rect></g><circle cx="12" cy="10" r="4"/><rect width="8" height="8" x="8" y="10"/><rect width="11" height="8" x="3" y="10" rx="4"><animate attributeName="x" dur="19s" repeatCount="indefinite" values="3;2;3;4;3"/></rect><rect width="13" height="6" x="8" y="12" rx="3"><animate attributeName="x" dur="23s" repeatCount="indefinite" values="8;7;8;9;8"/></rect><g fill="#fff"><rect width="3" height="4" x="10.5" y="10"/><path d="M12 17L16 13H8L12 17Z"><animateMotion calcMode="linear" dur="1.5s" keyPoints="0;0.25;0.5;0.75;1" keyTimes="0;0.1;0.5;0.8;1" path="M0 0v1v-2z" repeatCount="indefinite"/></path></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudDownloadOutlineLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M9 7L12 5L15 7L21 15L18 19H6L3 15L9 7Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.5s" values="0;1"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M12 19C12 19 9.5 19 7 19C4.5 19 3 17 3 15C3 13 4.5 11 7 11C8 11 8.5 11.5 8.5 11.5M12 19H17C19.5 19 21 17 21 15C21 13 19.5 11 17 11C16 11 15.5 11.5 15.5 11.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="9" stroke-dashoffset="9" d="M17 11C17 11 17 10.5 17 10C17 7.5 15 5 12 5M7 11V10C7 7.5 9 5 12 5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.3s" values="9;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdCloudLoop0"><g fill="#fff"><circle cx="12" cy="10" r="6"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;11;12;13;12"/></circle><rect width="9" height="7" x="8" y="13"/><rect width="15" height="12" x="1" y="8" rx="6"><animate attributeName="x" dur="19s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="13" height="10" x="10" y="10" rx="5"><animate attributeName="x" dur="23s" repeatCount="indefinite" values="10;9;10;11;10"/></rect></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudOffOutlineLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdCloudOffOutlineLoop0"><g fill="#fff"><circle cx="12" cy="10" r="6"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;11;12;13;12"/></circle><rect width="9" height="8" x="8" y="12"/><rect width="17" height="12" x="1" y="8" rx="6"><animate attributeName="x" dur="15s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="17" height="10" x="6" y="10" rx="5"><animate attributeName="x" dur="21s" repeatCount="indefinite" values="6;5;6;7;6"/></rect></g><circle cx="12" cy="10" r="4"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;11;12;13;12"/></circle><rect width="8" height="8" x="8" y="10"><animate attributeName="x" dur="30s" repeatCount="indefinite" values="8;7;8;9;8"/></rect><rect width="11" height="8" x="3" y="10" rx="4"><animate attributeName="x" dur="15s" repeatCount="indefinite" values="3;2;3;4;3"/></rect><rect width="13" height="6" x="8" y="12" rx="3"><animate attributeName="x" dur="21s" repeatCount="indefinite" values="8;7;8;9;8"/></rect><g fill="none" stroke-linecap="round" stroke-width="2"><path stroke="#000" d="M1 11h24" transform="rotate(45 13 12)"/><path stroke="#fff" d="M1 13h22" transform="rotate(45 13 12)"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M1 13h22;M3 13h22;M1 13h22"/></path></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudOffOutlineLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudOutlineLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdCloudOutlineLoop0"><g fill="#fff"><circle cx="12" cy="10" r="6"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;11;12;13;12"/></circle><rect width="9" height="8" x="8" y="12"/><rect width="17" height="12" x="1" y="8" rx="6"><animate attributeName="x" dur="21s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="17" height="10" x="6" y="10" rx="5"><animate attributeName="x" dur="17s" repeatCount="indefinite" values="6;5;6;7;6"/></rect></g><circle cx="12" cy="10" r="4"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;11;12;13;12"/></circle><rect width="8" height="8" x="8" y="10"><animate attributeName="x" dur="30s" repeatCount="indefinite" values="8;7;8;9;8"/></rect><rect width="11" height="8" x="3" y="10" rx="4"><animate attributeName="x" dur="21s" repeatCount="indefinite" values="3;2;3;4;3"/></rect><rect width="13" height="6" x="8" y="12" rx="3"><animate attributeName="x" dur="17s" repeatCount="indefinite" values="8;7;8;9;8"/></rect></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudOutlineLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudPrintLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdCloudPrintLoop0"><g fill="#fff"><circle cx="12" cy="8" r="6"/><rect width="15" height="12" x="1" y="6" rx="6"><animate attributeName="x" dur="24s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="13" height="10" x="10" y="8" rx="5"><animate attributeName="x" dur="17s" repeatCount="indefinite" values="10;9;10;11;10"/></rect></g><rect width="12" height="11" x="6" y="11" fill="#fff"/><rect width="8" height="7" x="8" y="13"/><path fill="#fff" d="M9 12h6v1H9zM9 14h6v1H9zM9 16h6v1H9zM9 18h6v1H9z"><animateMotion dur="1.5s" path="M0 0v2" repeatCount="indefinite"/></path></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudPrintLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudPrintOutlineLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdCloudPrintOutlineLoop0"><g fill="#fff"><circle cx="12" cy="8" r="6"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;11;12;13;12"/></circle><rect width="17" height="12" x="1" y="6" rx="6"><animate attributeName="x" dur="19s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="17" height="10" x="6" y="8" rx="5"><animate attributeName="x" dur="23s" repeatCount="indefinite" values="6;5;6;7;6"/></rect></g><circle cx="12" cy="8" r="4"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;11;12;13;12"/></circle><rect width="8" height="8" x="8" y="8"><animate attributeName="x" dur="30s" repeatCount="indefinite" values="8;7;8;9;8"/></rect><rect width="11" height="8" x="3" y="8" rx="4"><animate attributeName="x" dur="19s" repeatCount="indefinite" values="3;2;3;4;3"/></rect><rect width="13" height="6" x="8" y="10" rx="3"><animate attributeName="x" dur="23s" repeatCount="indefinite" values="8;7;8;9;8"/></rect><rect width="12" height="11" x="6" y="11" fill="#fff"/><rect width="8" height="7" x="8" y="13"/><path fill="#fff" d="M9 12h6v1H9zM9 14h6v1H9zM9 16h6v1H9zM9 18h6v1H9z"><animateMotion dur="1.5s" path="M0 0v2" repeatCount="indefinite"/></path></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudPrintOutlineLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudTagsLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdCloudTagsLoop0"><g fill="#fff"><circle cx="12" cy="10" r="6"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;13;12"/></circle><rect width="9" height="7" x="8" y="13"/><rect width="15" height="12" x="1" y="8" rx="6"><animate attributeName="x" dur="15s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="13" height="10" x="10" y="10" rx="5"><animate attributeName="x" dur="15s" repeatCount="indefinite" values="10;11;10;9;10"/></rect></g><path d="M6 10H8V14H12V16H6V10Z" transform="rotate(45 9 13)"><animateMotion calcMode="linear" dur="6s" keyPoints="0;0.25;0.5;0.75;1" keyTimes="0;0.1;0.5;0.8;1" path="M0 0h-1h2z" repeatCount="indefinite"/></path><path d="M12 10H14V14H18V16H12V10Z" transform="rotate(225 15 13)"><animateMotion calcMode="linear" dur="6s" keyPoints="0;0.25;0.5;0.75;1" keyTimes="0;0.1;0.5;0.8;1" path="M0 0h1h-2z" repeatCount="indefinite"/></path></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudTagsLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M9 7L12 5L15 7L21 15L18 19H6L3 15L9 7Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M12 19C12 19 9.5 19 7 19C4.5 19 3 17 3 15C3 13 4.5 11 7 11C8 11 8.5 11.5 8.5 11.5M12 19H17C19.5 19 21 17 21 15C21 13 19.5 11 17 11C16 11 15.5 11.5 15.5 11.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="9" stroke-dashoffset="9" d="M17 11C17 11 17 10.5 17 10C17 7.5 15 5 12 5M7 11V10C7 7.5 9 5 12 5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.3s" values="9;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUp(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="16" stroke-dashoffset="16" d="M8 18H7C4.5 18 3 16 3 14C3 12 4.5 10 7 10C8 10 8.5 10.5 8.5 10.5M16 18H17C19.5 18 21 16 21 14C21 12 19.5 10 17 10C16 10 15.5 10.5 15.5 10.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="16;0"/></path><path stroke-dasharray="9" stroke-dashoffset="9" d="M17 10C17 10 17 9.5 17 9C17 6.5 15 4 12 4M7 10V9C7 6.5 9 4 12 4"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.3s" values="9;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M12 20V14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="4" stroke-dashoffset="4" d="M12 13L14 15M12 13L10 15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="4;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUpTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M9 6L12 4L15 6L21 14L17 19H7L3 14L9 6Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.3s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="16" stroke-dashoffset="16" d="M8 18H7C4.5 18 3 16 3 14C3 12 4.5 10 7 10C8 10 8.5 10.5 8.5 10.5M16 18H17C19.5 18 21 16 21 14C21 12 19.5 10 17 10C16 10 15.5 10.5 15.5 10.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="16;0"/></path><path stroke-dasharray="9" stroke-dashoffset="9" d="M17 10C17 10 17 9.5 17 9C17 6.5 15 4 12 4M7 10V9C7 6.5 9 4 12 4"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.3s" values="9;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M12 20V14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="4" stroke-dashoffset="4" d="M12 13L14 15M12 13L10 15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="4;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUploadLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdCloudUploadLoop0"><g fill="#fff"><circle cx="12" cy="10" r="6"/><rect width="9" height="7" x="8" y="13"/><rect width="15" height="12" x="1" y="8" rx="6"><animate attributeName="x" dur="21s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="13" height="10" x="10" y="10" rx="5"><animate attributeName="x" dur="19s" repeatCount="indefinite" values="10;9;10;11;10"/></rect></g><rect width="4" height="5" x="10" y="12"/><path d="M12 8L17 13H7L12 8Z"><animateMotion calcMode="linear" dur="1.5s" keyPoints="0;0.25;0.5;0.75;1" keyTimes="0;0.1;0.5;0.8;1" path="M0 0v-1v2z" repeatCount="indefinite"/></path></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudUploadLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUploadOutlineLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdCloudUploadOutlineLoop0"><g fill="#fff"><circle cx="12" cy="10" r="6"/><rect width="9" height="8" x="8" y="12"/><rect width="17" height="12" x="1" y="8" rx="6"><animate attributeName="x" dur="24s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="17" height="10" x="6" y="10" rx="5"><animate attributeName="x" dur="15s" repeatCount="indefinite" values="6;5;6;7;6"/></rect></g><circle cx="12" cy="10" r="4"/><rect width="8" height="8" x="8" y="10"/><rect width="11" height="8" x="3" y="10" rx="4"><animate attributeName="x" dur="24s" repeatCount="indefinite" values="3;2;3;4;3"/></rect><rect width="13" height="6" x="8" y="12" rx="3"><animate attributeName="x" dur="15s" repeatCount="indefinite" values="8;7;8;9;8"/></rect><g fill="#fff"><rect width="3" height="4" x="10.5" y="12"/><path d="M12 9L16 13H8L12 9Z"><animateMotion calcMode="linear" dur="1.5s" keyPoints="0;0.25;0.5;0.75;1" keyTimes="0;0.1;0.5;0.8;1" path="M0 0v-1v2z" repeatCount="indefinite"/></path></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCloudUploadOutlineLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coffee(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="48" stroke-dashoffset="48" d="M17 4v9a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3V4z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.6s" values="48;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M17 9H20C20.55 9 21 8.55 21 8V5C21 4.45 20.55 4 20 4H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="14;28"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M11 20h8M11 20h-8"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="10;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoffeeArrow(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="40" stroke-dashoffset="40" d="M14 4V11C14 11.7956 13.6839 12.5587 13.1213 13.1213C12.5587 13.6839 11.7956 14 11 14H7C6.20435 14 5.44129 13.6839 4.87868 13.1213C4.31607 12.5587 4 11.7956 4 11V4z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="40;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M14 9H17C17.55 9 18 8.55 18 8V5C18 4.45 17.55 4 17 4H14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="14;28"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M4 18H19.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.3s" values="18;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M19.5 18l-3 -3M19.5 18l-3 3"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoffeeArrowFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="40" stroke-dashoffset="40" d="M14 4V11C14 11.7956 13.6839 12.5587 13.1213 13.1213C12.5587 13.6839 11.7956 14 11 14H7C6.20435 14 5.44129 13.6839 4.87868 13.1213C4.31607 12.5587 4 11.7956 4 11V4z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="40;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.4s" dur="0.5s" values="0;1"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M14 9H17C17.55 9 18 8.55 18 8V5C18 4.45 17.55 4 17 4H14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="14;28"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M4 18H19.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.3s" values="18;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M19.5 18l-3 -3M19.5 18l-3 3"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoffeeArrowTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="40" stroke-dashoffset="40" d="M14 4V11C14 11.7956 13.6839 12.5587 13.1213 13.1213C12.5587 13.6839 11.7956 14 11 14H7C6.20435 14 5.44129 13.6839 4.87868 13.1213C4.31607 12.5587 4 11.7956 4 11V4z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="40;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.4s" dur="0.15s" values="0;0.3"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M14 9H17C17.55 9 18 8.55 18 8V5C18 4.45 17.55 4 17 4H14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="14;28"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M4 18H19.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.3s" values="18;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M19.5 18l-3 -3M19.5 18l-3 3"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoffeeFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="48" stroke-dashoffset="48" d="M17 4v9a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3V4z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.6s" values="48;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.5s" values="0;1"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M17 9H20C20.55 9 21 8.55 21 8V5C21 4.45 20.55 4 20 4H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="14;28"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M11 20h8M11 20h-8"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="10;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoffeeHalfEmptyTwotoneLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M17 14V17a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3V14z"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="48" stroke-dashoffset="48" d="M17 9v9a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3V9z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="48;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M17 14H20C20.55 14 21 13.55 21 13V10C21 9.45 20.55 9 20 9H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="14;28"/></path></g><mask id="lineMdCoffeeHalfEmptyTwotoneLoop0"><path fill="none" stroke="#fff" stroke-width="2" d="M8 0c0 2-2 2-2 4s2 2 2 4-2 2-2 4 2 2 2 4M12 0c0 2-2 2-2 4s2 2 2 4-2 2-2 4 2 2 2 4M16 0c0 2-2 2-2 4s2 2 2 4-2 2-2 4 2 2 2 4"><animateMotion calcMode="linear" dur="3s" path="M0 0v-8" repeatCount="indefinite"/></path></mask><rect width="24" height="0" y="7" fill="currentColor" mask="url(#lineMdCoffeeHalfEmptyTwotoneLoop0)"><animate fill="freeze" attributeName="y" begin="0.8s" dur="0.6s" values="7;2"/><animate fill="freeze" attributeName="height" begin="0.8s" dur="0.6s" values="0;5"/></rect>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoffeeLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="48" stroke-dashoffset="48" d="M17 9v9a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3V9z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="48;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M17 14H20C20.55 14 21 13.55 21 13V10C21 9.45 20.55 9 20 9H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="14;28"/></path></g><mask id="lineMdCoffeeLoop0"><path fill="none" stroke="#fff" stroke-width="2" d="M8 0c0 2-2 2-2 4s2 2 2 4-2 2-2 4 2 2 2 4M12 0c0 2-2 2-2 4s2 2 2 4-2 2-2 4 2 2 2 4M16 0c0 2-2 2-2 4s2 2 2 4-2 2-2 4 2 2 2 4"><animateMotion calcMode="linear" dur="3s" path="M0 0v-8" repeatCount="indefinite"/></path></mask><rect width="24" height="0" y="7" fill="currentColor" mask="url(#lineMdCoffeeLoop0)"><animate fill="freeze" attributeName="y" begin="0.8s" dur="0.6s" values="7;2"/><animate fill="freeze" attributeName="height" begin="0.8s" dur="0.6s" values="0;5"/></rect>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoffeeTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="48" stroke-dashoffset="48" d="M17 4v9a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3V4z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.6s" values="48;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M17 9H20C20.55 9 21 8.55 21 8V5C21 4.45 20.55 4 20 4H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="14;28"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M11 20h8M11 20h-8"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="10;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoffeeTwotoneLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="48" stroke-dashoffset="48" d="M17 9v9a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3V9z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="48;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M17 14H20C20.55 14 21 13.55 21 13V10C21 9.45 20.55 9 20 9H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="14;28"/></path></g><mask id="lineMdCoffeeTwotoneLoop0"><path fill="none" stroke="#fff" stroke-width="2" d="M8 0c0 2-2 2-2 4s2 2 2 4-2 2-2 4 2 2 2 4M12 0c0 2-2 2-2 4s2 2 2 4-2 2-2 4 2 2 2 4M16 0c0 2-2 2-2 4s2 2 2 4-2 2-2 4 2 2 2 4"><animateMotion calcMode="linear" dur="3s" path="M0 0v-8" repeatCount="indefinite"/></path></mask><rect width="24" height="0" y="7" fill="currentColor" mask="url(#lineMdCoffeeTwotoneLoop0)"><animate fill="freeze" attributeName="y" begin="0.8s" dur="0.6s" values="7;2"/><animate fill="freeze" attributeName="height" begin="0.8s" dur="0.6s" values="0;5"/></rect>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cog(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><symbol id="lineMdCog0"><path fill="none" stroke-width="2" d="M15.24 6.37C15.65 6.6 16.04 6.88 16.38 7.2C16.6 7.4 16.8 7.61 16.99 7.83C17.46 8.4 17.85 9.05 18.11 9.77C18.2 10.03 18.28 10.31 18.35 10.59C18.45 11.04 18.5 11.52 18.5 12"><animate fill="freeze" attributeName="d" begin="0.8s" dur="0.2s" values="M15.24 6.37C15.65 6.6 16.04 6.88 16.38 7.2C16.6 7.4 16.8 7.61 16.99 7.83C17.46 8.4 17.85 9.05 18.11 9.77C18.2 10.03 18.28 10.31 18.35 10.59C18.45 11.04 18.5 11.52 18.5 12;M15.24 6.37C15.65 6.6 16.04 6.88 16.38 7.2C16.38 7.2 19 6.12 19.01 6.14C19.01 6.14 20.57 8.84 20.57 8.84C20.58 8.87 18.35 10.59 18.35 10.59C18.45 11.04 18.5 11.52 18.5 12"/></path></symbol></defs><g fill="none" stroke="currentColor" stroke-width="2"><g stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="42" stroke-dashoffset="42" d="M12 5.5C15.59 5.5 18.5 8.41 18.5 12C18.5 15.59 15.59 18.5 12 18.5C8.41 18.5 5.5 15.59 5.5 12C5.5 8.41 8.41 5.5 12 5.5z" opacity="0"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.5s" values="42;0"/><set attributeName="opacity" begin="0.2s" to="1"/><set attributeName="opacity" begin="0.7s" to="0"/></path><path stroke-dasharray="20" stroke-dashoffset="20" d="M12 9C13.66 9 15 10.34 15 12C15 13.66 13.66 15 12 15C10.34 15 9 13.66 9 12C9 10.34 10.34 9 12 9z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;0"/></path></g><g opacity="0"><use href="#lineMdCog0"/><use href="#lineMdCog0" transform="rotate(60 12 12)"/><use href="#lineMdCog0" transform="rotate(120 12 12)"/><use href="#lineMdCog0" transform="rotate(180 12 12)"/><use href="#lineMdCog0" transform="rotate(240 12 12)"/><use href="#lineMdCog0" transform="rotate(300 12 12)"/><set attributeName="opacity" begin="0.7s" to="1"/></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CogFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><symbol id="lineMdCogFilled0"><path fill="#fff" d="M11 13L15.74 5.5C16.03 5.67 16.31 5.85 16.57 6.05C16.57 6.05 16.57 6.05 16.57 6.05C16.64 6.1 16.71 6.16 16.77 6.22C18.14 7.34 19.09 8.94 19.4 10.75C19.41 10.84 19.42 10.92 19.43 11C19.43 11 19.43 11 19.43 11C19.48 11.33 19.5 11.66 19.5 12z"><animate fill="freeze" attributeName="d" begin="0.5s" dur="0.2s" values="M11 13L15.74 5.5C16.03 5.67 16.31 5.85 16.57 6.05C16.57 6.05 16.57 6.05 16.57 6.05C16.64 6.1 16.71 6.16 16.77 6.22C18.14 7.34 19.09 8.94 19.4 10.75C19.41 10.84 19.42 10.92 19.43 11C19.43 11 19.43 11 19.43 11C19.48 11.33 19.5 11.66 19.5 12z;M11 13L15.74 5.5C16.03 5.67 16.31 5.85 16.57 6.05C16.57 6.05 19.09 5.04 19.09 5.04C19.25 4.98 19.52 5.01 19.6 5.17C19.6 5.17 21.67 8.75 21.67 8.75C21.77 8.92 21.73 9.2 21.6 9.32C21.6 9.32 19.43 11 19.43 11C19.48 11.33 19.5 11.66 19.5 12z"/></path></symbol><mask id="lineMdCogFilled1"><path fill="none" stroke="#fff" stroke-dasharray="36" stroke-dashoffset="36" stroke-width="5" d="M12 7C14.76 7 17 9.24 17 12C17 14.76 14.76 17 12 17C9.24 17 7 14.76 7 12C7 9.24 9.24 7 12 7z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="36;0"/><set attributeName="opacity" begin="0.4s" to="0"/></path><g opacity="0"><use href="#lineMdCogFilled0"/><use href="#lineMdCogFilled0" transform="rotate(60 12 12)"/><use href="#lineMdCogFilled0" transform="rotate(120 12 12)"/><use href="#lineMdCogFilled0" transform="rotate(180 12 12)"/><use href="#lineMdCogFilled0" transform="rotate(240 12 12)"/><use href="#lineMdCogFilled0" transform="rotate(300 12 12)"/><set attributeName="opacity" begin="0.4s" to="1"/></g><circle cx="12" cy="12" r="3.5"/></mask></defs><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCogFilled1)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CogFilledLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><symbol id="lineMdCogFilledLoop0"><path fill="#fff" d="M11 13L15.74 5.5C16.03 5.67 16.31 5.85 16.57 6.05C16.57 6.05 16.57 6.05 16.57 6.05C16.64 6.1 16.71 6.16 16.77 6.22C18.14 7.34 19.09 8.94 19.4 10.75C19.41 10.84 19.42 10.92 19.43 11C19.43 11 19.43 11 19.43 11C19.48 11.33 19.5 11.66 19.5 12z"><animate fill="freeze" attributeName="d" begin="0.5s" dur="0.2s" values="M11 13L15.74 5.5C16.03 5.67 16.31 5.85 16.57 6.05C16.57 6.05 16.57 6.05 16.57 6.05C16.64 6.1 16.71 6.16 16.77 6.22C18.14 7.34 19.09 8.94 19.4 10.75C19.41 10.84 19.42 10.92 19.43 11C19.43 11 19.43 11 19.43 11C19.48 11.33 19.5 11.66 19.5 12z;M11 13L15.74 5.5C16.03 5.67 16.31 5.85 16.57 6.05C16.57 6.05 19.09 5.04 19.09 5.04C19.25 4.98 19.52 5.01 19.6 5.17C19.6 5.17 21.67 8.75 21.67 8.75C21.77 8.92 21.73 9.2 21.6 9.32C21.6 9.32 19.43 11 19.43 11C19.48 11.33 19.5 11.66 19.5 12z"/></path></symbol><mask id="lineMdCogFilledLoop1"><path fill="none" stroke="#fff" stroke-dasharray="36" stroke-dashoffset="36" stroke-width="5" d="M12 7C14.76 7 17 9.24 17 12C17 14.76 14.76 17 12 17C9.24 17 7 14.76 7 12C7 9.24 9.24 7 12 7z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="36;0"/><set attributeName="opacity" begin="0.4s" to="0"/></path><g opacity="0"><use href="#lineMdCogFilledLoop0"/><use href="#lineMdCogFilledLoop0" transform="rotate(60 12 12)"/><use href="#lineMdCogFilledLoop0" transform="rotate(120 12 12)"/><use href="#lineMdCogFilledLoop0" transform="rotate(180 12 12)"/><use href="#lineMdCogFilledLoop0" transform="rotate(240 12 12)"/><use href="#lineMdCogFilledLoop0" transform="rotate(300 12 12)"/><set attributeName="opacity" begin="0.4s" to="1"/><animateTransform attributeName="transform" dur="30s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></g><circle cx="12" cy="12" r="3.5"/></mask></defs><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCogFilledLoop1)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CogLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><symbol id="lineMdCogLoop0"><path fill="none" stroke-width="2" d="M15.24 6.37C15.65 6.6 16.04 6.88 16.38 7.2C16.6 7.4 16.8 7.61 16.99 7.83C17.46 8.4 17.85 9.05 18.11 9.77C18.2 10.03 18.28 10.31 18.35 10.59C18.45 11.04 18.5 11.52 18.5 12"><animate fill="freeze" attributeName="d" begin="0.8s" dur="0.2s" values="M15.24 6.37C15.65 6.6 16.04 6.88 16.38 7.2C16.6 7.4 16.8 7.61 16.99 7.83C17.46 8.4 17.85 9.05 18.11 9.77C18.2 10.03 18.28 10.31 18.35 10.59C18.45 11.04 18.5 11.52 18.5 12;M15.24 6.37C15.65 6.6 16.04 6.88 16.38 7.2C16.38 7.2 19 6.12 19.01 6.14C19.01 6.14 20.57 8.84 20.57 8.84C20.58 8.87 18.35 10.59 18.35 10.59C18.45 11.04 18.5 11.52 18.5 12"/></path></symbol></defs><g fill="none" stroke="currentColor" stroke-width="2"><g stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="42" stroke-dashoffset="42" d="M12 5.5C15.59 5.5 18.5 8.41 18.5 12C18.5 15.59 15.59 18.5 12 18.5C8.41 18.5 5.5 15.59 5.5 12C5.5 8.41 8.41 5.5 12 5.5z" opacity="0"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.5s" values="42;0"/><set attributeName="opacity" begin="0.2s" to="1"/><set attributeName="opacity" begin="0.7s" to="0"/></path><path stroke-dasharray="20" stroke-dashoffset="20" d="M12 9C13.66 9 15 10.34 15 12C15 13.66 13.66 15 12 15C10.34 15 9 13.66 9 12C9 10.34 10.34 9 12 9z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;0"/></path></g><g opacity="0"><use href="#lineMdCogLoop0"/><use href="#lineMdCogLoop0" transform="rotate(60 12 12)"/><use href="#lineMdCogLoop0" transform="rotate(120 12 12)"/><use href="#lineMdCogLoop0" transform="rotate(180 12 12)"/><use href="#lineMdCogLoop0" transform="rotate(240 12 12)"/><use href="#lineMdCogLoop0" transform="rotate(300 12 12)"/><set attributeName="opacity" begin="0.7s" to="1"/><animateTransform attributeName="transform" dur="30s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CogOff(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><symbol id="lineMdCogOff0"><path fill="none" stroke-width="2" d="M15.24 6.37C15.65 6.6 16.04 6.88 16.38 7.2C16.6 7.4 16.8 7.61 16.99 7.83C17.46 8.4 17.85 9.05 18.11 9.77C18.2 10.03 18.28 10.31 18.35 10.59C18.45 11.04 18.5 11.52 18.5 12"><animate fill="freeze" attributeName="d" begin="0.8s" dur="0.2s" values="M15.24 6.37C15.65 6.6 16.04 6.88 16.38 7.2C16.6 7.4 16.8 7.61 16.99 7.83C17.46 8.4 17.85 9.05 18.11 9.77C18.2 10.03 18.28 10.31 18.35 10.59C18.45 11.04 18.5 11.52 18.5 12;M15.24 6.37C15.65 6.6 16.04 6.88 16.38 7.2C16.38 7.2 19 6.12 19.01 6.14C19.01 6.14 20.57 8.84 20.57 8.84C20.58 8.87 18.35 10.59 18.35 10.59C18.45 11.04 18.5 11.52 18.5 12"/></path></symbol><mask id="lineMdCogOff1"><g fill="none" stroke="#fff" stroke-width="2"><g stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="42" stroke-dashoffset="42" d="M12 5.5C15.59 5.5 18.5 8.41 18.5 12C18.5 15.59 15.59 18.5 12 18.5C8.41 18.5 5.5 15.59 5.5 12C5.5 8.41 8.41 5.5 12 5.5z" opacity="0"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.5s" values="42;0"/><set attributeName="opacity" begin="0.2s" to="1"/><set attributeName="opacity" begin="0.7s" to="0"/></path><path stroke-dasharray="20" stroke-dashoffset="20" d="M12 9C13.66 9 15 10.34 15 12C15 13.66 13.66 15 12 15C10.34 15 9 13.66 9 12C9 10.34 10.34 9 12 9z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;0"/></path></g><g opacity="0"><use href="#lineMdCogOff0"/><use href="#lineMdCogOff0" transform="rotate(60 12 12)"/><use href="#lineMdCogOff0" transform="rotate(120 12 12)"/><use href="#lineMdCogOff0" transform="rotate(180 12 12)"/><use href="#lineMdCogOff0" transform="rotate(240 12 12)"/><use href="#lineMdCogOff0" transform="rotate(300 12 12)"/><set attributeName="opacity" begin="0.7s" to="1"/></g></g><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 12 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M0 13h24"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="26;0"/></g></mask></defs><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCogOff1)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CogOffFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><symbol id="lineMdCogOffFilled0"><path fill="#fff" d="M11 13L15.74 5.5C16.03 5.67 16.31 5.85 16.57 6.05C16.57 6.05 16.57 6.05 16.57 6.05C16.64 6.1 16.71 6.16 16.77 6.22C18.14 7.34 19.09 8.94 19.4 10.75C19.41 10.84 19.42 10.92 19.43 11C19.43 11 19.43 11 19.43 11C19.48 11.33 19.5 11.66 19.5 12z"><animate fill="freeze" attributeName="d" begin="0.5s" dur="0.2s" values="M11 13L15.74 5.5C16.03 5.67 16.31 5.85 16.57 6.05C16.57 6.05 16.57 6.05 16.57 6.05C16.64 6.1 16.71 6.16 16.77 6.22C18.14 7.34 19.09 8.94 19.4 10.75C19.41 10.84 19.42 10.92 19.43 11C19.43 11 19.43 11 19.43 11C19.48 11.33 19.5 11.66 19.5 12z;M11 13L15.74 5.5C16.03 5.67 16.31 5.85 16.57 6.05C16.57 6.05 19.09 5.04 19.09 5.04C19.25 4.98 19.52 5.01 19.6 5.17C19.6 5.17 21.67 8.75 21.67 8.75C21.77 8.92 21.73 9.2 21.6 9.32C21.6 9.32 19.43 11 19.43 11C19.48 11.33 19.5 11.66 19.5 12z"/></path></symbol><mask id="lineMdCogOffFilled1"><path fill="none" stroke="#fff" stroke-dasharray="36" stroke-dashoffset="36" stroke-width="5" d="M12 7C14.76 7 17 9.24 17 12C17 14.76 14.76 17 12 17C9.24 17 7 14.76 7 12C7 9.24 9.24 7 12 7z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="36;0"/><set attributeName="opacity" begin="0.4s" to="0"/></path><g opacity="0"><use href="#lineMdCogOffFilled0"/><use href="#lineMdCogOffFilled0" transform="rotate(60 12 12)"/><use href="#lineMdCogOffFilled0" transform="rotate(120 12 12)"/><use href="#lineMdCogOffFilled0" transform="rotate(180 12 12)"/><use href="#lineMdCogOffFilled0" transform="rotate(240 12 12)"/><use href="#lineMdCogOffFilled0" transform="rotate(300 12 12)"/><set attributeName="opacity" begin="0.4s" to="1"/></g><circle cx="12" cy="12" r="3.5"/><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 12 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M0 13h24"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="26;0"/></g></mask></defs><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCogOffFilled1)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CogOffFilledLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><symbol id="lineMdCogOffFilledLoop0"><path fill="#fff" d="M11 13L15.74 5.5C16.03 5.67 16.31 5.85 16.57 6.05C16.57 6.05 16.57 6.05 16.57 6.05C16.64 6.1 16.71 6.16 16.77 6.22C18.14 7.34 19.09 8.94 19.4 10.75C19.41 10.84 19.42 10.92 19.43 11C19.43 11 19.43 11 19.43 11C19.48 11.33 19.5 11.66 19.5 12z"><animate fill="freeze" attributeName="d" begin="0.5s" dur="0.2s" values="M11 13L15.74 5.5C16.03 5.67 16.31 5.85 16.57 6.05C16.57 6.05 16.57 6.05 16.57 6.05C16.64 6.1 16.71 6.16 16.77 6.22C18.14 7.34 19.09 8.94 19.4 10.75C19.41 10.84 19.42 10.92 19.43 11C19.43 11 19.43 11 19.43 11C19.48 11.33 19.5 11.66 19.5 12z;M11 13L15.74 5.5C16.03 5.67 16.31 5.85 16.57 6.05C16.57 6.05 19.09 5.04 19.09 5.04C19.25 4.98 19.52 5.01 19.6 5.17C19.6 5.17 21.67 8.75 21.67 8.75C21.77 8.92 21.73 9.2 21.6 9.32C21.6 9.32 19.43 11 19.43 11C19.48 11.33 19.5 11.66 19.5 12z"/></path></symbol><mask id="lineMdCogOffFilledLoop1"><path fill="none" stroke="#fff" stroke-dasharray="36" stroke-dashoffset="36" stroke-width="5" d="M12 7C14.76 7 17 9.24 17 12C17 14.76 14.76 17 12 17C9.24 17 7 14.76 7 12C7 9.24 9.24 7 12 7z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="36;0"/><set attributeName="opacity" begin="0.4s" to="0"/></path><g opacity="0"><use href="#lineMdCogOffFilledLoop0"/><use href="#lineMdCogOffFilledLoop0" transform="rotate(60 12 12)"/><use href="#lineMdCogOffFilledLoop0" transform="rotate(120 12 12)"/><use href="#lineMdCogOffFilledLoop0" transform="rotate(180 12 12)"/><use href="#lineMdCogOffFilledLoop0" transform="rotate(240 12 12)"/><use href="#lineMdCogOffFilledLoop0" transform="rotate(300 12 12)"/><set attributeName="opacity" begin="0.4s" to="1"/><animateTransform attributeName="transform" dur="30s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></g><circle cx="12" cy="12" r="3.5"/><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 12 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M0 13h24"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M-1 13h24;M1 13h24;M-1 13h24"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="26;0"/></g></mask></defs><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCogOffFilledLoop1)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CogOffLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><symbol id="lineMdCogOffLoop0"><path fill="none" stroke-width="2" d="M15.24 6.37C15.65 6.6 16.04 6.88 16.38 7.2C16.6 7.4 16.8 7.61 16.99 7.83C17.46 8.4 17.85 9.05 18.11 9.77C18.2 10.03 18.28 10.31 18.35 10.59C18.45 11.04 18.5 11.52 18.5 12"><animate fill="freeze" attributeName="d" begin="0.8s" dur="0.2s" values="M15.24 6.37C15.65 6.6 16.04 6.88 16.38 7.2C16.6 7.4 16.8 7.61 16.99 7.83C17.46 8.4 17.85 9.05 18.11 9.77C18.2 10.03 18.28 10.31 18.35 10.59C18.45 11.04 18.5 11.52 18.5 12;M15.24 6.37C15.65 6.6 16.04 6.88 16.38 7.2C16.38 7.2 19 6.12 19.01 6.14C19.01 6.14 20.57 8.84 20.57 8.84C20.58 8.87 18.35 10.59 18.35 10.59C18.45 11.04 18.5 11.52 18.5 12"/></path></symbol><mask id="lineMdCogOffLoop1"><g fill="none" stroke="#fff" stroke-width="2"><g stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="42" stroke-dashoffset="42" d="M12 5.5C15.59 5.5 18.5 8.41 18.5 12C18.5 15.59 15.59 18.5 12 18.5C8.41 18.5 5.5 15.59 5.5 12C5.5 8.41 8.41 5.5 12 5.5z" opacity="0"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.5s" values="42;0"/><set attributeName="opacity" begin="0.2s" to="1"/><set attributeName="opacity" begin="0.7s" to="0"/></path><path stroke-dasharray="20" stroke-dashoffset="20" d="M12 9C13.66 9 15 10.34 15 12C15 13.66 13.66 15 12 15C10.34 15 9 13.66 9 12C9 10.34 10.34 9 12 9z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;0"/></path></g><g opacity="0"><use href="#lineMdCogOffLoop0"/><use href="#lineMdCogOffLoop0" transform="rotate(60 12 12)"/><use href="#lineMdCogOffLoop0" transform="rotate(120 12 12)"/><use href="#lineMdCogOffLoop0" transform="rotate(180 12 12)"/><use href="#lineMdCogOffLoop0" transform="rotate(240 12 12)"/><use href="#lineMdCogOffLoop0" transform="rotate(300 12 12)"/><set attributeName="opacity" begin="0.7s" to="1"/><animateTransform attributeName="transform" dur="30s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></g></g><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 12 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M0 13h24"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M-1 13h24;M1 13h24;M-1 13h24"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="26;0"/></g></mask></defs><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCogOffLoop1)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compass(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdCompass0"><path fill="none" stroke="#fff" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path fill="#fff" d="M11 11L12 12L13 13L12 12z"><set attributeName="opacity" begin="0.6s" to="1"/><animate fill="freeze" attributeName="d" begin="0.6s" dur="0.3s" values="M11 11L12 12L13 13L12 12z;M10.2 10.2L17 7L13.8 13.8L7 17z"/><animateTransform attributeName="transform" begin="0.5s" dur="0.5s" type="rotate" values="-180 12 12;0 12 12"/></path><circle cx="12" cy="12" r="1" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.3s" values="0;1"/></circle></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCompass0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompassLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdCompassLoop0"><path fill="none" stroke="#fff" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path fill="#fff" d="M11 11L12 12L13 13L12 12z"><set attributeName="opacity" begin="0.6s" to="1"/><animate fill="freeze" attributeName="d" begin="0.6s" dur="0.3s" values="M11 11L12 12L13 13L12 12z;M10.2 10.2L17 7L13.8 13.8L7 17z"/><animateTransform attributeName="transform" begin="0.5s" dur="9s" repeatCount="indefinite" type="rotate" values="-180 12 12;0 12 12;0 12 12;0 12 12;0 12 12;270 12 12;-90 12 12;0 12 12;-180 12 12;-35 12 12;-40 12 12;-45 12 12;-45 12 12;-110 12 12;-135 12 12;-180 12 12"/></path><circle cx="12" cy="12" r="1" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.3s" values="0;1"/></circle></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCompassLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompassOff(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdCompassOff0"><path fill="none" stroke="#fff" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path fill="#fff" d="M11 11L12 12L13 13L12 12z"><set attributeName="opacity" begin="0.6s" to="1"/><animate fill="freeze" attributeName="d" begin="0.6s" dur="0.3s" values="M11 11L12 12L13 13L12 12z;M10.2 10.2L17 7L13.8 13.8L7 17z"/><animateTransform attributeName="transform" begin="0.5s" dur="0.5s" type="rotate" values="-180 12 12;0 12 12"/></path><circle cx="12" cy="12" r="1" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.3s" values="0;1"/></circle><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 12 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M0 13h24"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCompassOff0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompassTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdCompassTwotone0"><path fill="#fff" fill-opacity="0" stroke="#fff" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.15s" values="0;0.3"/></path><path fill="#fff" d="M11 11L12 12L13 13L12 12z"><set attributeName="opacity" begin="0.6s" to="1"/><animate fill="freeze" attributeName="d" begin="0.6s" dur="0.3s" values="M11 11L12 12L13 13L12 12z;M10.2 10.2L17 7L13.8 13.8L7 17z"/><animateTransform attributeName="transform" begin="0.5s" dur="0.5s" type="rotate" values="-180 12 12;0 12 12"/></path><circle cx="12" cy="12" r="1" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.3s" values="0;1"/></circle></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCompassTwotone0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompassTwotoneLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdCompassTwotoneLoop0"><path fill="#fff" fill-opacity="0" stroke="#fff" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.15s" values="0;0.3"/></path><path fill="#fff" d="M11 11L12 12L13 13L12 12z"><set attributeName="opacity" begin="0.6s" to="1"/><animate fill="freeze" attributeName="d" begin="0.6s" dur="0.3s" values="M11 11L12 12L13 13L12 12z;M10.2 10.2L17 7L13.8 13.8L7 17z"/><animateTransform attributeName="transform" begin="0.5s" dur="9s" repeatCount="indefinite" type="rotate" values="-180 12 12;0 12 12;0 12 12;0 12 12;0 12 12;270 12 12;-90 12 12;0 12 12;-180 12 12;-35 12 12;-40 12 12;-45 12 12;-45 12 12;-110 12 12;-135 12 12;-180 12 12"/></path><circle cx="12" cy="12" r="1" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.3s" values="0;1"/></circle></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCompassTwotoneLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompassTwotoneOff(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdCompassTwotoneOff0"><path fill="#fff" fill-opacity="0" stroke="#fff" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.15s" values="0;0.3"/></path><path fill="#fff" d="M11 11L12 12L13 13L12 12z"><set attributeName="opacity" begin="0.6s" to="1"/><animate fill="freeze" attributeName="d" begin="0.6s" dur="0.3s" values="M11 11L12 12L13 13L12 12z;M10.2 10.2L17 7L13.8 13.8L7 17z"/><animateTransform attributeName="transform" begin="0.5s" dur="0.5s" type="rotate" values="-180 12 12;0 12 12"/></path><circle cx="12" cy="12" r="1" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.3s" values="0;1"/></circle><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 12 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M0 13h24"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdCompassTwotoneOff0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Computer(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 21H17M12 21H7"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="6;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 21V17"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="6;0"/></path><path stroke-dasharray="64" stroke-dashoffset="64" d="M12 17H3V5H21V17Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.6s" values="64;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ComputerTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 21H17M12 21H7"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="6;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 21V17"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="6;0"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="64" stroke-dashoffset="64" d="M12 17H3V5H21V17Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.6s" values="64;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Confirm(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="24" stroke-dashoffset="24" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 11L11 17L21 7"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="24;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ConfirmCircle(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M8 12L11 15L16 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="14;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ConfirmCircleToCircleTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path stroke-dasharray="14" d="M8 12L11 15L16 10"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="28;14"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ConfirmCircleTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="14" stroke-dashoffset="14" d="M8 12L11 15L16 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="14;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ConfirmCircleTwotoneToCircleTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9" fill="currentColor" fill-opacity=".3"><animate fill="freeze" attributeName="fill-opacity" begin="0.2s" dur="0.15s" values="0.3;0"/></circle><path fill="none" stroke-dasharray="14" d="M8 12L11 15L16 10"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="28;14"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ConfirmCircleTwotoneToCircleTwotoneTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9" fill="currentColor" fill-opacity=".3"/><path fill="none" stroke-dasharray="14" d="M8 12L11 15L16 10"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="28;14"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ConfirmSquare(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M4 12V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="64;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M8 12L11 15L16 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="14;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ConfirmSquareToSquareTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 19V5a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1Z"/><path stroke-dasharray="14" d="M8 12L11 15L16 10"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="28;14"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ConfirmSquareTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="64" stroke-dashoffset="64" d="M4 12V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="64;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="14" stroke-dashoffset="14" d="M8 12L11 15L16 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="14;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ConfirmSquareTwotoneToSquareTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".3" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 19V5a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.2s" dur="0.15s" values="0.3;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ConfirmSquareTwotoneToSquareTwotoneTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity=".3" d="M20 19V5a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1Z"/><path fill="none" stroke-dasharray="14" d="M8 12L11 15L16 10"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="28;14"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Construction(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="44" stroke-dashoffset="44" d="M21 21H3V19C3 18 4 17 5 17H19C20 17 21 18 21 19V21Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="44;0"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M6 17L12 2M18 17L12 2"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.3s" values="18;0"/></path></g><path stroke-dasharray="8" stroke-dashoffset="8" d="M8 12L12.5 9.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M6 16L13.5 12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="10;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M9.5 17L14.5 14.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.3s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ConstructionTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M21 21H3V19C3 18 4 17 5 17H19C20 17 21 18 21 19V21Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.5s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="44" stroke-dashoffset="44" d="M21 21H3V19C3 18 4 17 5 17H19C20 17 21 18 21 19V21Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="44;0"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M6 17L12 2M18 17L12 2"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.3s" values="18;0"/></path></g><path stroke-dasharray="8" stroke-dashoffset="8" d="M8 12L12.5 9.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M6 16L13.5 12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="10;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M9.5 17L14.5 14.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.3s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Discord(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-opacity="0"><circle cx="9" cy="12" r="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.4s" values="0;1"/></circle><circle cx="15" cy="12" r="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="1.4s" dur="0.4s" values="0;1"/></circle></g><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="30" stroke-dashoffset="30" d="M15.5 17.5L16.5 19.5C16.5 19.5 20.671 18.172 22 16C22 15 22.53 7.853 19 5.5C17.5 4.5 15 4 15 4L14 6H12M8.52799 17.5L7.52799 19.5C7.52799 19.5 3.35699 18.172 2.02799 16C2.02799 15 1.49799 7.853 5.02799 5.5C6.52799 4.5 9.02799 4 9.02799 4L10.028 6H12.028"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="30;60"/></path><path stroke-dasharray="16" stroke-dashoffset="16" d="M5.5 16C10.5 18.5 13.5 18.5 18.5 16"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.4s" values="16;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiscordTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-opacity="0"><path d="M5 5L12 5.2L19 5L22 15L19 18.4H5L1.5 15L5 5Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.6s" dur="0.15s" values="0;0.3"/></path><circle cx="9" cy="12" r="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.4s" values="0;1"/></circle><circle cx="15" cy="12" r="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="1.4s" dur="0.4s" values="0;1"/></circle></g><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="30" stroke-dashoffset="30" d="M15.5 17.5L16.5 19.5C16.5 19.5 20.671 18.172 22 16C22 15 22.53 7.853 19 5.5C17.5 4.5 15 4 15 4L14 6H12M8.52799 17.5L7.52799 19.5C7.52799 19.5 3.35699 18.172 2.02799 16C2.02799 15 1.49799 7.853 5.02799 5.5C6.52799 4.5 9.02799 4 9.02799 4L10.028 6H12.028"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="30;60"/></path><path stroke-dasharray="16" stroke-dashoffset="16" d="M5.5 16C10.5 18.5 13.5 18.5 18.5 16"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.4s" values="16;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Document(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="64" stroke-dashoffset="64" stroke-width="2" d="M13 3L19 9V21H5V3H13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M12.5 3V8.5H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentAdd(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="64" stroke-dashoffset="64" stroke-width="2" d="M13 3L19 9V21H5V3H13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M12.5 3V8.5H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path><g stroke-dasharray="8" stroke-dashoffset="8" stroke-width="2"><path d="M9 14H15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path><path d="M12 11V17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="8;0"/></path></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentAddTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M5 3H12.5V8.5H19V21H5V3Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.4s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="64" stroke-dashoffset="64" stroke-width="2" d="M13 3L19 9V21H5V3H13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M12.5 3V8.5H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path><g stroke-dasharray="8" stroke-dashoffset="8" stroke-width="2"><path d="M9 14H15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path><path d="M12 11V17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="8;0"/></path></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentCode(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="64" stroke-dashoffset="64" stroke-width="2" d="M13 3L19 9V21H5V3H13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M12.5 3V8.5H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path><g stroke-dasharray="8" stroke-dashoffset="8" stroke-width="2"><path d="M10 13L8 15L10 17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path><path d="M14 13L16 15L14 17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="8;0"/></path></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentCodeTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M5 3H12.5V8.5H19V21H5V3Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.4s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="64" stroke-dashoffset="64" stroke-width="2" d="M13 3L19 9V21H5V3H13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M12.5 3V8.5H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path><g stroke-dasharray="8" stroke-dashoffset="8" stroke-width="2"><path d="M10 13L8 15L10 17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path><path d="M14 13L16 15L14 17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="8;0"/></path></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentList(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M13 3L19 9V21H5V3H13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M9 13H13"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="6;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M9 16H15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="8;0"/></path></g><path stroke-dasharray="14" stroke-dashoffset="14" d="M12.5 3V8.5H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentListTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M5 3H12.5V8.5H19V21H5V3Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.4s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><g stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M13 3L19 9V21H5V3H13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M9 13H13"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="6;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M9 16H15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="8;0"/></path></g><path stroke-dasharray="14" stroke-dashoffset="14" d="M12.5 3V8.5H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentRemove(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="64" stroke-dashoffset="64" stroke-width="2" d="M13 3L19 9V21H5V3H13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M12.5 3V8.5H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" stroke-width="2" d="M9 14H15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentRemoveTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M5 3H12.5V8.5H19V21H5V3Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="64" stroke-dashoffset="64" stroke-width="2" d="M13 3L19 9V21H5V3H13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M12.5 3V8.5H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" stroke-width="2" d="M9 14H15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentReport(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="64" stroke-dashoffset="64" stroke-width="2" d="M13 3L19 9V21H5V3H13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M12.5 3V8.5H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path><g stroke-width="2"><path stroke-dasharray="5" stroke-dashoffset="5" d="M9 17V14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="5;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 17V13"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="6;0"/></path><path stroke-dasharray="7" stroke-dashoffset="7" d="M15 17V12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.4s" dur="0.2s" values="7;0"/></path></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentReportTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M5 3H12.5V8.5H19V21H5V3Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.6s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="64" stroke-dashoffset="64" stroke-width="2" d="M13 3L19 9V21H5V3H13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M12.5 3V8.5H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path><g stroke-width="2"><path stroke-dasharray="5" stroke-dashoffset="5" d="M9 17V14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="5;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 17V13"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="6;0"/></path><path stroke-dasharray="7" stroke-dashoffset="7" d="M15 17V12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.4s" dur="0.2s" values="7;0"/></path></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M5 3H12.5V8.5H19V21H5V3Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="64" stroke-dashoffset="64" stroke-width="2" d="M13 3L19 9V21H5V3H13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M12.5 3V8.5H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleArrowHorizontal(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="10" stroke-dashoffset="10" d="M12 12H3.5M12 12H20.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="10;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M3 12L7 16M21 12L17 16M3 12L7 8M21 12L17 8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.3s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="none" stroke-dasharray="14" stroke-dashoffset="14" d="M6 19h12"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="14;0"/></path><path fill="currentColor" d="M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5"><animate attributeName="d" calcMode="linear" dur="1.5s" keyTimes="0;0.7;1" repeatCount="indefinite" values="M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5;M12 4 h2 v3 h2.5 L12 11.5M12 4 h-2 v3 h-2.5 L12 11.5;M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadOffLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdDownloadOffLoop0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M6 19h12"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="14;0"/></path><path fill="#fff" d="M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5"><animate attributeName="d" calcMode="linear" dur="1.5s" keyTimes="0;0.7;1" repeatCount="indefinite" values="M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5;M12 4 h2 v3 h2.5 L12 11.5M12 4 h-2 v3 h-2.5 L12 11.5;M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5"/></path><g stroke-dasharray="26" stroke-dashoffset="26" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path d="M0 13h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 13h22;M2 13h22;M0 13h22"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="26;0"/></g></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdDownloadOffLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadOffOutline(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdDownloadOffOutline0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M6 19h12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="14;0"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="18;0"/></path><g stroke-dasharray="26" stroke-dashoffset="26" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path d="M1 13h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="26;0"/></g></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdDownloadOffOutline0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadOffOutlineLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdDownloadOffOutlineLoop0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M6 19h12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="14;0"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="18;0"/><animate attributeName="d" calcMode="linear" dur="1.5s" keyTimes="0;0.7;1" repeatCount="indefinite" values="M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5;M12 4 h2 v3 h2.5 L12 11.5M12 4 h-2 v3 h-2.5 L12 11.5;M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5"/></path><g stroke-dasharray="26" stroke-dashoffset="26" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path d="M0 13h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 13h22;M2 13h22;M0 13h22"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="26;0"/></g></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdDownloadOffOutlineLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadOutline(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M6 19h12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="14;0"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="18;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadOutlineLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M6 19h12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="14;0"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="18;0"/><animate attributeName="d" calcMode="linear" dur="1.5s" keyTimes="0;0.7;1" repeatCount="indefinite" values="M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5;M12 4 h2 v3 h2.5 L12 11.5M12 4 h-2 v3 h-2.5 L12 11.5;M12 4 h2 v6 h2.5 L12 14.5M12 4 h-2 v6 h-2.5 L12 14.5"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadingLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="2 4" stroke-dashoffset="6" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21"><animate attributeName="stroke-dashoffset" dur="0.6s" repeatCount="indefinite" values="6;0"/></path><path stroke-dasharray="30" stroke-dashoffset="30" d="M12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.1s" dur="0.3s" values="30;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M12 8v7.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="10;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 15.5l3.5 -3.5M12 15.5l-3.5 -3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Edit(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M3 21H21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="20;0"/></path><path stroke-dasharray="44" stroke-dashoffset="44" d="M7 17V13L17 3L21 7L11 17H7"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.6s" values="44;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M14 6L18 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M20 7L17 4L15 6L18 9L20 7Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M3 21H21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="20;0"/></path><path stroke-dasharray="44" stroke-dashoffset="44" d="M7 17V13L17 3L21 7L11 17H7"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.6s" values="44;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M14 6L18 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditTwotoneFull(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M20 7L17 4L8 13V16H11L20 7Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M3 21H21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="20;0"/></path><path stroke-dasharray="44" stroke-dashoffset="44" d="M7 17V13L17 3L21 7L11 17H7"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.6s" values="44;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M14 6L18 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Email(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><rect width="18" height="14" x="3" y="5" stroke-dasharray="64" stroke-dashoffset="64" rx="1"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></rect><path stroke-dasharray="24" stroke-dashoffset="24" d="M3 6.5L12 12L21 6.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.4s" values="24;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmailOpened(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M3 8.06083C3 7.71247 3.1813 7.38921 3.47855 7.20755L12 2L20.5214 7.20755C20.8187 7.38921 21 7.71247 21 8.06083V18C21 18.5523 20.5523 19 20 19H4C3.44772 19 3 18.5523 3 18V8.06083Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="24" stroke-dashoffset="24" d="M3 8.5L12 14L21 8.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.4s" values="24;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmailOpenedTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M12 13L4 8L12 3L20 8L12 13Z"><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M3 8.06083C3 7.71247 3.1813 7.38921 3.47855 7.20755L12 2L20.5214 7.20755C20.8187 7.38921 21 7.71247 21 8.06083V18C21 18.5523 20.5523 19 20 19H4C3.44772 19 3 18.5523 3 18V8.06083Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="24" stroke-dashoffset="24" d="M3 8.5L12 14L21 8.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.4s" values="24;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmailOpenedTwotoneAlt(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M12 15L4 10V18H20V10L12 15Z"><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M3 8.06083C3 7.71247 3.1813 7.38921 3.47855 7.20755L12 2L20.5214 7.20755C20.8187 7.38921 21 7.71247 21 8.06083V18C21 18.5523 20.5523 19 20 19H4C3.44772 19 3 18.5523 3 18V8.06083Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="24" stroke-dashoffset="24" d="M3 8.5L12 14L21 8.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.4s" values="24;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmailTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M12 11L4 6H20L12 11Z"><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><rect width="18" height="14" x="3" y="5" stroke-dasharray="64" stroke-dashoffset="64" rx="1"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></rect><path stroke-dasharray="24" stroke-dashoffset="24" d="M3 6.5L12 12L21 6.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.4s" values="24;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmailTwotoneAlt(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M12 13L4 8V18H20V8L12 13Z"><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><rect width="18" height="14" x="3" y="5" stroke-dasharray="64" stroke-dashoffset="64" rx="1"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></rect><path stroke-dasharray="24" stroke-dashoffset="24" d="M3 6.5L12 12L21 6.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.4s" values="24;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiAngry(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M8 16C8.5 15 9.79086 14 12 14C14.2091 14 15.5 15 16 16"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="14;0"/></path><g stroke-dasharray="5" stroke-dashoffset="5" stroke-width="1"><path d="M9.5 8.5L7.5 7.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="5;0"/></path><path d="M14.5 8.5L16.5 7.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="5;0"/></path></g></g><g fill="currentColor" fill-opacity="0"><ellipse cx="9" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><ellipse cx="15" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></ellipse></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiAngryTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M8 16C8.5 15 9.79086 14 12 14C14.2091 14 15.5 15 16 16"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="14;0"/></path><g stroke-dasharray="5" stroke-dashoffset="5" stroke-width="1"><path d="M9.5 8.5L7.5 7.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="5;0"/></path><path d="M14.5 8.5L16.5 7.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="5;0"/></path></g></g><g fill="currentColor" fill-opacity="0"><ellipse cx="9" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><ellipse cx="15" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></ellipse></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiFrown(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M8 16C8.5 15 9.79086 14 12 14C14.2091 14 15.5 15 16 16"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="14;0"/></path></g><g fill="currentColor" fill-opacity="0"><ellipse cx="9" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><ellipse cx="15" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></ellipse></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiFrownOpen(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdEmojiFrownOpen0"><path fill="none" stroke="#fff" stroke-linecap="round" stroke-width="2" d="M8.5 16C9 15 10 14 12 14C14 14 15 15 15.5 16M8.5 16C9.5 15.5 11 15.5 12 15.5C13 15.5 14.5 15.5 15.5 16"/></mask><path fill="none" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><g fill="currentColor"><rect width="0" height="7" x="6" y="12" mask="url(#lineMdEmojiFrownOpen0)"><animate fill="freeze" attributeName="width" begin="1s" dur="0.2s" values="0;12"/></rect><ellipse cx="9" cy="9.5" fill-opacity="0" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><ellipse cx="15" cy="9.5" fill-opacity="0" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></ellipse></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiFrownOpenTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdEmojiFrownOpenTwotone0"><path fill="none" stroke="#fff" stroke-linecap="round" stroke-width="2" d="M8.5 16C9 15 10 14 12 14C14 14 15 15 15.5 16M8.5 16C9.5 15.5 11 15.5 12 15.5C13 15.5 14.5 15.5 15.5 16"/></mask><g fill="currentColor"><path fill-opacity="0" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><rect width="0" height="7" x="6" y="12" mask="url(#lineMdEmojiFrownOpenTwotone0)"><animate fill="freeze" attributeName="width" begin="1s" dur="0.2s" values="0;12"/></rect><ellipse cx="9" cy="9.5" fill-opacity="0" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><ellipse cx="15" cy="9.5" fill-opacity="0" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></ellipse></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiFrownTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="14" stroke-dashoffset="14" d="M8 16C8.5 15 9.79086 14 12 14C14.2091 14 15.5 15 16 16"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="14;0"/></path></g><g fill="currentColor" fill-opacity="0"><ellipse cx="9" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><ellipse cx="15" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></ellipse></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiGrin(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdEmojiGrin0"><path fill="none" stroke="#fff" stroke-linecap="round" stroke-width="2" d="M8 14C8.5 15.5 9.79086 17 12 17C14.2091 17 15.5 15.5 16 14M8 14C9 14.5 10 15 12 15C14 15 15 14.5 16 14"/></mask><path fill="none" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><g fill="currentColor"><rect width="0" height="7" x="6" y="12" mask="url(#lineMdEmojiGrin0)"><animate fill="freeze" attributeName="width" begin="1s" dur="0.2s" values="0;12"/></rect><ellipse cx="9" cy="9.5" fill-opacity="0" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><ellipse cx="15" cy="9.5" fill-opacity="0" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></ellipse></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiGrinTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdEmojiGrinTwotone0"><path fill="none" stroke="#fff" stroke-linecap="round" stroke-width="2" d="M8 14C8.5 15.5 9.79086 17 12 17C14.2091 17 15.5 15.5 16 14M8 14C9 14.5 10 15 12 15C14 15 15 14.5 16 14"/></mask><g fill="currentColor"><path fill-opacity="0" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><rect width="0" height="7" x="6" y="12" mask="url(#lineMdEmojiGrinTwotone0)"><animate fill="freeze" attributeName="width" begin="1s" dur="0.2s" values="0;12"/></rect><ellipse cx="9" cy="9.5" fill-opacity="0" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><ellipse cx="15" cy="9.5" fill-opacity="0" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></ellipse></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiNeutral(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M8 15H16"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="10;0"/></path></g><g fill="currentColor" fill-opacity="0"><ellipse cx="9" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><ellipse cx="15" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></ellipse></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiNeutralTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="10" stroke-dashoffset="10" d="M8 15H16"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="10;0"/></path></g><g fill="currentColor" fill-opacity="0"><ellipse cx="9" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><ellipse cx="15" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></ellipse></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiSmile(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M8 14C8.5 15.5 9.79086 17 12 17C14.2091 17 15.5 15.5 16 14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="14;0"/></path></g><g fill="currentColor" fill-opacity="0"><ellipse cx="9" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><ellipse cx="15" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></ellipse></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiSmileTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="14" stroke-dashoffset="14" d="M8 14C8.5 15.5 9.79086 17 12 17C14.2091 17 15.5 15.5 16 14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="14;0"/></path></g><g fill="currentColor" fill-opacity="0"><ellipse cx="9" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><ellipse cx="15" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></ellipse></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiSmileWink(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M8 14C8.5 15.5 9.79086 17 12 17C14.2091 17 15.5 15.5 16 14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="14;0"/></path></g><g fill="currentColor" fill-opacity="0"><ellipse cx="9" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><path d="M15 8.5C15.8284 8.5 16.5 8.94772 16.5 9.5C16.5 10.0523 15.8284 10.5 15 10.5C14.1716 10.5 13.5 10.0523 13.5 9.5C13.5 8.94772 14.1716 8.5 15 8.5Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiSmileWinkTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="14" stroke-dashoffset="14" d="M8 14C8.5 15.5 9.79086 17 12 17C14.2091 17 15.5 15.5 16 14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="14;0"/></path></g><g fill="currentColor" fill-opacity="0"><ellipse cx="9" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><path d="M15 8.5C15.8284 8.5 16.5 8.94772 16.5 9.5C16.5 10.0523 15.8284 10.5 15 10.5C14.1716 10.5 13.5 10.0523 13.5 9.5C13.5 8.94772 14.1716 8.5 15 8.5Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExternalLink(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="42" stroke-dashoffset="42" d="M11 5H5V19H19V13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="42;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M13 11L20 4"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.3s" values="12;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M21 3H15M21 3V9"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExternalLinkRounded(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="36" stroke-dashoffset="36" d="M12 5C8.13401 5 5 8.13401 5 12C5 15.866 8.13401 19 12 19C15.866 19 19 15.866 19 12"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="36;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M13 11L20 4"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.3s" values="12;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M21 3H15M21 3V9"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Facebook(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-dasharray="24" stroke-dashoffset="24" d="M17 4L15 4C12.5 4 11 5.5 11 8V20"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="24;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M8 12H15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="12;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filter(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="54" stroke-dashoffset="54" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 4h14l-5 6.5v9.5l-4 -4v-5.5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="54;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="54" stroke-dashoffset="54" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 4h14l-5 6.5v9.5l-4 -4v-5.5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="54;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.5s" values="0;1"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="54" stroke-dashoffset="54" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 4h14l-5 6.5v9.5l-4 -4v-5.5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="54;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.15s" values="0;0.3"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForkLeft(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="18" stroke-dashoffset="18" d="M5 13H9C11.7614 13 14 15.2386 14 18V20M14 4v16"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="18;36"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M5 13l3 -3M5 13l3 3M14 4l3 3M14 4l-3 3" opacity="0"><set attributeName="opacity" begin="0.4s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gauge(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path d="M12 12H12" opacity="0"><animate fill="freeze" attributeName="d" begin="1.2s" dur="0.2s" values="M12 12H12;M12 12H6.5"/><set attributeName="opacity" begin="1.2s" to="1"/><animateTransform fill="freeze" attributeName="transform" begin="1.4s" dur="0.6s" type="rotate" values="0 12 12;115 12 12"/></path></g><g fill="currentColor"><path fill-opacity="0" d="M12 21C9.41 21 7.15 20.79 5.94 19L12 21L18.06 19C16.85 20.79 14.59 21 12 21Z"><set attributeName="fill-opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="d" begin="0.5s" dur="0.4s" values="M12 21C9.41 21 7.15 20.79 5.94 19L12 21L18.06 19C16.85 20.79 14.59 21 12 21Z;M12 16C9.41 16 7.15 17.21 5.94 19L12 21L18.06 19C16.85 17.21 14.59 16 12 16Z"/></path><circle cx="7" cy="12" r="0" transform="rotate(15 12 12)"><animate fill="freeze" attributeName="r" begin="0.8s" dur="0.2s" values="0;1"/></circle><circle cx="7" cy="12" r="0" transform="rotate(65 12 12)"><animate fill="freeze" attributeName="r" begin="0.85s" dur="0.2s" values="0;1"/></circle><circle cx="7" cy="12" r="0" transform="rotate(115 12 12)"><animate fill="freeze" attributeName="r" begin="0.9s" dur="0.2s" values="0;1"/></circle><circle cx="7" cy="12" r="0" transform="rotate(165 12 12)"><animate fill="freeze" attributeName="r" begin="0.95s" dur="0.2s" values="0;1"/></circle><circle cx="12" cy="12" r="0"><animate fill="freeze" attributeName="r" begin="1.2s" dur="0.2s" values="0;2"/></circle></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GaugeEmpty(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path d="M12 12H12" opacity="0"><animate fill="freeze" attributeName="d" begin="1.2s" dur="0.2s" values="M12 12H12;M12 12H6.5"/><set attributeName="opacity" begin="1.2s" to="1"/><animateTransform fill="freeze" attributeName="transform" begin="1.4s" dur="0.2s" type="rotate" values="0 12 12;15 12 12"/></path></g><g fill="currentColor"><path fill-opacity="0" d="M12 21C9.41 21 7.15 20.79 5.94 19L12 21L18.06 19C16.85 20.79 14.59 21 12 21Z"><set attributeName="fill-opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="d" begin="0.5s" dur="0.4s" values="M12 21C9.41 21 7.15 20.79 5.94 19L12 21L18.06 19C16.85 20.79 14.59 21 12 21Z;M12 16C9.41 16 7.15 17.21 5.94 19L12 21L18.06 19C16.85 17.21 14.59 16 12 16Z"/></path><circle cx="7" cy="12" r="0" transform="rotate(15 12 12)"><animate fill="freeze" attributeName="r" begin="0.8s" dur="0.2s" values="0;1"/></circle><circle cx="7" cy="12" r="0" transform="rotate(65 12 12)"><animate fill="freeze" attributeName="r" begin="0.85s" dur="0.2s" values="0;1"/></circle><circle cx="7" cy="12" r="0" transform="rotate(115 12 12)"><animate fill="freeze" attributeName="r" begin="0.9s" dur="0.2s" values="0;1"/></circle><circle cx="7" cy="12" r="0" transform="rotate(165 12 12)"><animate fill="freeze" attributeName="r" begin="0.95s" dur="0.2s" values="0;1"/></circle><circle cx="12" cy="12" r="0"><animate fill="freeze" attributeName="r" begin="1.2s" dur="0.2s" values="0;2"/></circle></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GaugeFull(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path d="M12 12H12" opacity="0"><animate fill="freeze" attributeName="d" begin="1.2s" dur="0.2s" values="M12 12H12;M12 12H6.5"/><set attributeName="opacity" begin="1.2s" to="1"/><animateTransform fill="freeze" attributeName="transform" begin="1.4s" dur="0.8s" type="rotate" values="0 12 12;165 12 12"/></path></g><g fill="currentColor"><path fill-opacity="0" d="M12 21C9.41 21 7.15 20.79 5.94 19L12 21L18.06 19C16.85 20.79 14.59 21 12 21Z"><set attributeName="fill-opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="d" begin="0.5s" dur="0.4s" values="M12 21C9.41 21 7.15 20.79 5.94 19L12 21L18.06 19C16.85 20.79 14.59 21 12 21Z;M12 16C9.41 16 7.15 17.21 5.94 19L12 21L18.06 19C16.85 17.21 14.59 16 12 16Z"/></path><circle cx="7" cy="12" r="0" transform="rotate(15 12 12)"><animate fill="freeze" attributeName="r" begin="0.8s" dur="0.2s" values="0;1"/></circle><circle cx="7" cy="12" r="0" transform="rotate(65 12 12)"><animate fill="freeze" attributeName="r" begin="0.85s" dur="0.2s" values="0;1"/></circle><circle cx="7" cy="12" r="0" transform="rotate(115 12 12)"><animate fill="freeze" attributeName="r" begin="0.9s" dur="0.2s" values="0;1"/></circle><circle cx="7" cy="12" r="0" transform="rotate(165 12 12)"><animate fill="freeze" attributeName="r" begin="0.95s" dur="0.2s" values="0;1"/></circle><circle cx="12" cy="12" r="0"><animate fill="freeze" attributeName="r" begin="1.2s" dur="0.2s" values="0;2"/></circle></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GaugeLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path d="M12 12H12" opacity="0"><animate fill="freeze" attributeName="d" begin="1.2s" dur="0.2s" values="M12 12H12;M12 12H6.5"/><set attributeName="opacity" begin="1.2s" to="1"/><animateTransform attributeName="transform" begin="1.4s" dur="15s" repeatCount="indefinite" type="rotate" values="0 12 12;15 12 12;15 12 12;65 12 12;115 12 12;165 12 12;165 12 12;165 12 12;90 12 12;115 12 12;115 12 12;15 12 12;0 12 12"/></path></g><g fill="currentColor"><path fill-opacity="0" d="M12 21C9.41 21 7.15 20.79 5.94 19L12 21L18.06 19C16.85 20.79 14.59 21 12 21Z"><set attributeName="fill-opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="d" begin="0.5s" dur="0.4s" values="M12 21C9.41 21 7.15 20.79 5.94 19L12 21L18.06 19C16.85 20.79 14.59 21 12 21Z;M12 16C9.41 16 7.15 17.21 5.94 19L12 21L18.06 19C16.85 17.21 14.59 16 12 16Z"/></path><circle cx="7" cy="12" r="0" transform="rotate(15 12 12)"><animate fill="freeze" attributeName="r" begin="0.8s" dur="0.2s" values="0;1"/></circle><circle cx="7" cy="12" r="0" transform="rotate(65 12 12)"><animate fill="freeze" attributeName="r" begin="0.85s" dur="0.2s" values="0;1"/></circle><circle cx="7" cy="12" r="0" transform="rotate(115 12 12)"><animate fill="freeze" attributeName="r" begin="0.9s" dur="0.2s" values="0;1"/></circle><circle cx="7" cy="12" r="0" transform="rotate(165 12 12)"><animate fill="freeze" attributeName="r" begin="0.95s" dur="0.2s" values="0;1"/></circle><circle cx="12" cy="12" r="0"><animate fill="freeze" attributeName="r" begin="1.2s" dur="0.2s" values="0;2"/></circle></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GaugeLow(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path d="M12 12H12" opacity="0"><animate fill="freeze" attributeName="d" begin="1.2s" dur="0.2s" values="M12 12H12;M12 12H6.5"/><set attributeName="opacity" begin="1.2s" to="1"/><animateTransform fill="freeze" attributeName="transform" begin="1.4s" dur="0.4s" type="rotate" values="0 12 12;65 12 12"/></path></g><g fill="currentColor"><path fill-opacity="0" d="M12 21C9.41 21 7.15 20.79 5.94 19L12 21L18.06 19C16.85 20.79 14.59 21 12 21Z"><set attributeName="fill-opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="d" begin="0.5s" dur="0.4s" values="M12 21C9.41 21 7.15 20.79 5.94 19L12 21L18.06 19C16.85 20.79 14.59 21 12 21Z;M12 16C9.41 16 7.15 17.21 5.94 19L12 21L18.06 19C16.85 17.21 14.59 16 12 16Z"/></path><circle cx="7" cy="12" r="0" transform="rotate(15 12 12)"><animate fill="freeze" attributeName="r" begin="0.8s" dur="0.2s" values="0;1"/></circle><circle cx="7" cy="12" r="0" transform="rotate(65 12 12)"><animate fill="freeze" attributeName="r" begin="0.85s" dur="0.2s" values="0;1"/></circle><circle cx="7" cy="12" r="0" transform="rotate(115 12 12)"><animate fill="freeze" attributeName="r" begin="0.9s" dur="0.2s" values="0;1"/></circle><circle cx="7" cy="12" r="0" transform="rotate(165 12 12)"><animate fill="freeze" attributeName="r" begin="0.95s" dur="0.2s" values="0;1"/></circle><circle cx="12" cy="12" r="0"><animate fill="freeze" attributeName="r" begin="1.2s" dur="0.2s" values="0;2"/></circle></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Github(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="30" stroke-dashoffset="30" d="M12 4C13.6683 4 14.6122 4.39991 15 4.5C15.5255 4.07463 16.9375 3 18.5 3C18.8438 4 18.7863 5.21921 18.5 6C19.25 7 19.5 8 19.5 9.5C19.5 11.6875 19.017 13.0822 18 14C16.983 14.9178 15.8887 15.3749 14.5 15.5C15.1506 16.038 15 17.3743 15 18C15 18.7256 15 21 15 21M12 4C10.3317 4 9.38784 4.39991 9 4.5C8.47455 4.07463 7.0625 3 5.5 3C5.15625 4 5.21371 5.21921 5.5 6C4.75 7 4.5 8 4.5 9.5C4.5 11.6875 4.98301 13.0822 6 14C7.01699 14.9178 8.1113 15.3749 9.5 15.5C8.84944 16.038 9 17.3743 9 18C9 18.7256 9 21 9 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="30;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M9 19C7.59375 19 6.15625 18.4375 5.3125 17.8125C4.46875 17.1875 4.21875 16.1562 3 15.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="10;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GithubLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdGithubLoop0" width="24" height="24" x="0" y="0"><g fill="#fff"><ellipse cx="9.5" cy="9" rx="1.5" ry="1"/><ellipse cx="14.5" cy="9" rx="1.5" ry="1"/></g></mask><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="30" stroke-dashoffset="30" d="M12 4C13.6683 4 14.6122 4.39991 15 4.5C15.5255 4.07463 16.9375 3 18.5 3C18.8438 4 18.7863 5.21921 18.5 6C19.25 7 19.5 8 19.5 9.5C19.5 11.6875 19.017 13.0822 18 14C16.983 14.9178 15.8887 15.3749 14.5 15.5C15.1506 16.038 15 17.3743 15 18C15 18.7256 15 21 15 21M12 4C10.3317 4 9.38784 4.39991 9 4.5C8.47455 4.07463 7.0625 3 5.5 3C5.15625 4 5.21371 5.21921 5.5 6C4.75 7 4.5 8 4.5 9.5C4.5 11.6875 4.98301 13.0822 6 14C7.01699 14.9178 8.1113 15.3749 9.5 15.5C8.84944 16.038 9 17.3743 9 18C9 18.7256 9 21 9 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="30;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M9 19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="10;0"/><animate attributeName="d" dur="3s" repeatCount="indefinite" values="M9 19c-1.406 0-2.844-.563-3.688-1.188C4.47 17.188 4.22 16.157 3 15.5;M9 19c-1.406 0-3-.5-4-.5-.532 0-1 0-2-.5;M9 19c-1.406 0-2.844-.563-3.688-1.188C4.47 17.188 4.22 16.157 3 15.5"/></path></g><rect width="8" height="4" x="8" y="11" fill="currentColor" mask="url(#lineMdGithubLoop0)"><animate attributeName="y" dur="10s" keyTimes="0;0.45;0.46;0.54;0.55;1" repeatCount="indefinite" values="11;11;7;7;11;11"/></rect>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GithubTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M15 4.5C14.6122 4.39991 13.6683 4 12 4C10.3317 4 9.38784 4.39991 9 4.5C8.47455 4.07463 7.0625 3 5.5 3C5.15625 4 5.21371 5.21921 5.5 6C4.75 7 4.5 8 4.5 9.5C4.5 11.6875 4.98302 13.0822 6 14C7.01698 14.9178 8.1113 15.3749 9.5 15.5C8.84944 16.038 9 17.3743 9 18V22H15V18C15 17.3743 15.1506 16.038 14.5 15.5C15.8887 15.3749 16.983 14.9178 18 14C19.017 13.0822 19.5 11.6875 19.5 9.5C19.5 8 19.25 7 18.5 6C18.7863 5.21921 18.8438 4 18.5 3C16.9375 3 15.5255 4.07463 15 4.5Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="30" stroke-dashoffset="30" d="M12 4C13.6683 4 14.6122 4.39991 15 4.5C15.5255 4.07463 16.9375 3 18.5 3C18.8438 4 18.7863 5.21921 18.5 6C19.25 7 19.5 8 19.5 9.5C19.5 11.6875 19.017 13.0822 18 14C16.983 14.9178 15.8887 15.3749 14.5 15.5C15.1506 16.038 15 17.3743 15 18C15 18.7256 15 21 15 21M12 4C10.3317 4 9.38784 4.39991 9 4.5C8.47455 4.07463 7.0625 3 5.5 3C5.15625 4 5.21371 5.21921 5.5 6C4.75 7 4.5 8 4.5 9.5C4.5 11.6875 4.98301 13.0822 6 14C7.01699 14.9178 8.1113 15.3749 9.5 15.5C8.84944 16.038 9 17.3743 9 18C9 18.7256 9 21 9 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="30;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M9 19C7.59375 19 6.15625 18.4375 5.3125 17.8125C4.46875 17.1875 4.21875 16.1562 3 15.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="10;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridThree(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="10" stroke-dashoffset="10" stroke-linecap="round"><g><circle cx="5" cy="5" r="1.5"/><circle cx="12" cy="5" r="1.5"/><circle cx="19" cy="5" r="1.5"/><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/></g><g><circle cx="5" cy="12" r="1.5"/><circle cx="12" cy="12" r="1.5"/><circle cx="19" cy="12" r="1.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.2s" values="10;0"/></g><g><circle cx="5" cy="19" r="1.5"/><circle cx="12" cy="19" r="1.5"/><circle cx="19" cy="19" r="1.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="10;0"/></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridThreeFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="10" stroke-dashoffset="10" stroke-linecap="round"><g><circle cx="5" cy="5" r="1.5"/><circle cx="12" cy="5" r="1.5"/><circle cx="19" cy="5" r="1.5"/><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.5s" values="0;1"/></g><g><circle cx="5" cy="12" r="1.5"/><circle cx="12" cy="12" r="1.5"/><circle cx="19" cy="12" r="1.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="10;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.5s" values="0;1"/></g><g><circle cx="5" cy="19" r="1.5"/><circle cx="12" cy="19" r="1.5"/><circle cx="19" cy="19" r="1.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="10;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.5s" values="0;1"/></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridThreeTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="10" stroke-dashoffset="10" stroke-linecap="round"><g><circle cx="5" cy="5" r="1.5"/><circle cx="12" cy="5" r="1.5"/><circle cx="19" cy="5" r="1.5"/><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></g><g><circle cx="5" cy="12" r="1.5"/><circle cx="12" cy="12" r="1.5"/><circle cx="19" cy="12" r="1.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.2s" values="10;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.15s" values="0;0.3"/></g><g><circle cx="5" cy="19" r="1.5"/><circle cx="12" cy="19" r="1.5"/><circle cx="19" cy="19" r="1.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="10;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hash(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="20" stroke-dashoffset="20" stroke-linecap="round" stroke-width="2"><path d="M4 9H21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;0"/></path><path d="M3 15H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="20;0"/></path><path d="M10 3L8 21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="20;0"/></path><path d="M16 3L14 21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="20;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HashSmall(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="16" stroke-dashoffset="16" stroke-linecap="round" stroke-width="2"><path d="M6 9H19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="16;0"/></path><path d="M5 15H18"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="16;0"/></path><path d="M10 5L8 19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="16;0"/></path><path d="M16 5L14 19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="16;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="30" stroke-dashoffset="30" stroke-linecap="round" stroke-width="2" d="M12 8C12 8 12 8 12.7578 7C13.6343 5.84335 14.9398 5 16.5 5C18.9853 5 21 7.01472 21 9.5C21 10.4251 20.7209 11.285 20.2422 12C19.435 13.206 12 21 12 21M12 8C12 8 12 8 11.2422 7C10.3657 5.84335 9.06021 5 7.5 5C5.01472 5 3 7.01472 3 9.5C3 10.4251 3.27914 11.285 3.75777 12C4.56504 13.206 12 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="30;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M12 20L20.5 11V7L17 5.5L12 7L7 5.5L3.5 7V11L12 20Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.5s" values="0;1"/></path><path fill="none" stroke="currentColor" stroke-dasharray="30" stroke-dashoffset="30" stroke-linecap="round" stroke-width="2" d="M12 8C12 8 12 8 12.7578 7C13.6343 5.84335 14.9398 5 16.5 5C18.9853 5 21 7.01472 21 9.5C21 10.4251 20.7209 11.285 20.2422 12C19.435 13.206 12 21 12 21M12 8C12 8 12 8 11.2422 7C10.3657 5.84335 9.06021 5 7.5 5C5.01472 5 3 7.01472 3 9.5C3 10.4251 3.27914 11.285 3.75777 12C4.56504 13.206 12 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="30;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartFilledHalf(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M3.5 11L12 20V7L7 5.5L3.5 7V11Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.5s" values="0;1"/></path><path fill="none" stroke="currentColor" stroke-dasharray="30" stroke-dashoffset="30" stroke-linecap="round" stroke-width="2" d="M12 8C12 8 12 8 12.7578 7C13.6343 5.84335 14.9398 5 16.5 5C18.9853 5 21 7.01472 21 9.5C21 10.4251 20.7209 11.285 20.2422 12C19.435 13.206 12 21 12 21M12 8C12 8 12 8 11.2422 7C10.3657 5.84335 9.06021 5 7.5 5C5.01472 5 3 7.01472 3 9.5C3 10.4251 3.27914 11.285 3.75777 12C4.56504 13.206 12 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="30;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartHalf(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="30" stroke-dashoffset="30" d="M12 8C12 8 12 8 11.2422 7C10.3657 5.84335 9.06021 5 7.5 5C5.01472 5 3 7.01472 3 9.5C3 10.4251 3.27914 11.285 3.75777 12C4.56504 13.206 12 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="30;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M12 8V21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="15;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartHalfFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M3.5 11L12 20V7L7 5.5L3.5 7V11Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.5s" values="0;1"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="30" stroke-dashoffset="30" d="M12 8C12 8 12 8 11.2422 7C10.3657 5.84335 9.06021 5 7.5 5C5.01472 5 3 7.01472 3 9.5C3 10.4251 3.27914 11.285 3.75777 12C4.56504 13.206 12 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="30;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M12 8V21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="15;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartHalfTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M3.5 11L12 20V7L7 5.5L3.5 7V11Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="30" stroke-dashoffset="30" d="M12 8C12 8 12 8 11.2422 7C10.3657 5.84335 9.06021 5 7.5 5C5.01472 5 3 7.01472 3 9.5C3 10.4251 3.27914 11.285 3.75777 12C4.56504 13.206 12 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="30;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M12 8V21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="15;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M12 20L20.5 11V7L17 5.5L12 7L7 5.5L3.5 7V11L12 20Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke="currentColor" stroke-dasharray="30" stroke-dashoffset="30" stroke-linecap="round" stroke-width="2" d="M12 8C12 8 12 8 12.7578 7C13.6343 5.84335 14.9398 5 16.5 5C18.9853 5 21 7.01472 21 9.5C21 10.4251 20.7209 11.285 20.2422 12C19.435 13.206 12 21 12 21M12 8C12 8 12 8 11.2422 7C10.3657 5.84335 9.06021 5 7.5 5C5.01472 5 3 7.01472 3 9.5C3 10.4251 3.27914 11.285 3.75777 12C4.56504 13.206 12 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="30;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartTwotoneHalf(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M3.5 11L12 20V7L7 5.5L3.5 7V11Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke="currentColor" stroke-dasharray="30" stroke-dashoffset="30" stroke-linecap="round" stroke-width="2" d="M12 8C12 8 12 8 12.7578 7C13.6343 5.84335 14.9398 5 16.5 5C18.9853 5 21 7.01472 21 9.5C21 10.4251 20.7209 11.285 20.2422 12C19.435 13.206 12 21 12 21M12 8C12 8 12 8 11.2422 7C10.3657 5.84335 9.06021 5 7.5 5C5.01472 5 3 7.01472 3 9.5C3 10.4251 3.27914 11.285 3.75777 12C4.56504 13.206 12 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="30;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartTwotoneHalfFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-opacity="0"><path d="M3.5 11L12 20V7L7 5.5L3.5 7V11Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.5s" values="0;1"/></path><path d="M12 20L20.5 11V7L17 5.5L12 7V20Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.15s" values="0;0.3"/></path></g><path fill="none" stroke="currentColor" stroke-dasharray="30" stroke-dashoffset="30" stroke-linecap="round" stroke-width="2" d="M12 8C12 8 12 8 12.7578 7C13.6343 5.84335 14.9398 5 16.5 5C18.9853 5 21 7.01472 21 9.5C21 10.4251 20.7209 11.285 20.2422 12C19.435 13.206 12 21 12 21M12 8C12 8 12 8 11.2422 7C10.3657 5.84335 9.06021 5 7.5 5C5.01472 5 3 7.01472 3 9.5C3 10.4251 3.27914 11.285 3.75777 12C4.56504 13.206 12 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="30;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="15" stroke-dashoffset="15" d="M4.5 21.5h15"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="15;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M4.5 21.5V8M19.5 21.5V8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="15;0"/></path><path stroke-dasharray="24" stroke-dashoffset="24" d="M9.5 21.5V12.5H14.5V21.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.4s" values="24;0"/></path><path stroke-dasharray="30" stroke-dashoffset="30" stroke-width="2" d="M2 10L12 2L22 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="30;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeMd(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="21" stroke-dashoffset="21" d="M5 21H19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="21;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M5 21V8M19 21V8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="15;0"/></path><path stroke-dasharray="24" stroke-dashoffset="24" d="M9 21V13H15V21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.4s" values="24;0"/></path><path stroke-dasharray="26" stroke-dashoffset="26" d="M2 10L12 2L22 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="26;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeMdTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M6 8L12 3L18 8V20H16V13L15 12H9L8 13V20H6V8Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="21" stroke-dashoffset="21" d="M5 21H19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="21;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M5 21V8M19 21V8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="15;0"/></path><path stroke-dasharray="24" stroke-dashoffset="24" d="M9 21V13H15V21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.4s" values="24;0"/></path><path stroke-dasharray="26" stroke-dashoffset="26" d="M2 10L12 2L22 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="26;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeMdTwotoneAlt(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<rect width="4" height="6" x="10" y="14" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></rect><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="21" stroke-dashoffset="21" d="M5 21H19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="21;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M5 21V8M19 21V8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="15;0"/></path><path stroke-dasharray="24" stroke-dashoffset="24" d="M9 21V13H15V21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.4s" values="24;0"/></path><path stroke-dasharray="26" stroke-dashoffset="26" d="M2 10L12 2L22 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="26;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeSimple(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="21" stroke-dashoffset="21" d="M5 21H19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="21;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M5 21V8M19 21V8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="15;0"/></path><path stroke-dasharray="26" stroke-dashoffset="26" d="M2 10L12 2L22 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="26;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeSimpleFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M6 8L12 3L18 8V20H6V8Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.5s" values="0;1"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="21" stroke-dashoffset="21" d="M5 21H19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="21;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M5 21V8M19 21V8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="15;0"/></path><path stroke-dasharray="26" stroke-dashoffset="26" d="M2 10L12 2L22 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="26;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeSimpleTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M6 8L12 3L18 8V20H6V8Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="21" stroke-dashoffset="21" d="M5 21H19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="21;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M5 21V8M19 21V8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="15;0"/></path><path stroke-dasharray="26" stroke-dashoffset="26" d="M2 10L12 2L22 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="26;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M5 8.5L12 3L19 8.5V21H15V13L14 12H10L9 13V21H5V8.5Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="15" stroke-dashoffset="15" d="M4.5 21.5h15"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="15;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M4.5 21.5V8M19.5 21.5V8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="15;0"/></path><path stroke-dasharray="24" stroke-dashoffset="24" d="M9.5 21.5V12.5H14.5V21.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.4s" values="24;0"/></path><path stroke-dasharray="30" stroke-dashoffset="30" stroke-width="2" d="M2 10L12 2L22 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="30;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeTwotoneAlt(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<rect width="4" height="8" x="10" y="13" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></rect><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="15" stroke-dashoffset="15" d="M4.5 21.5h15"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="15;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M4.5 21.5V8M19.5 21.5V8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="15;0"/></path><path stroke-dasharray="24" stroke-dashoffset="24" d="M9.5 21.5V12.5H14.5V21.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.4s" values="24;0"/></path><path stroke-dasharray="30" stroke-dashoffset="30" stroke-width="2" d="M2 10L12 2L22 10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="30;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IconifyOne(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><g fill="currentColor" fill-opacity="0"><path fill-rule="evenodd" d="M12 18C15.125 18 17.3257 15.122 17 14.5C16.6728 13.875 15.5 16 12 16C8.5 16 7.3125 13.875 7 14.5C6.6875 15.125 8.875 18 12 18Z" clip-rule="evenodd"><animate fill="freeze" attributeName="fill-opacity" begin="1.0s" dur="0.2s" values="0;1"/></path><path d="M9.5 9C9.5 8.48223 9.01777 8 8.5 8C7.98223 8 7.5 8.48223 7.5 9V10.4375C7.5 10.9553 7.98223 11.5 8.5 11.5C9.01777 11.5 9.5 11.0178 9.5 10.5V9Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></path><path d="M16.5 9C16.5 8.48223 16.0178 8 15.5 8C14.9822 8 14.5 8.48223 14.5 9V10.4375C14.5 10.9553 14.9822 11.5 15.5 11.5C16.0178 11.5 16.5 11.0178 16.5 10.5V9Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IconifyTwo(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><path d="M4 7V21" class="il-md-length-15 il-md-duration-2 il-md-delay-0"/><path d="M4 3V5" class="il-md-length-15 il-md-duration-2 il-md-delay-0"/><path stroke-linecap="round" d="M18 4.25204C17.3608 4.08751 16.6906 4 16 4C11.5817 4 8 7.58172 8 12C8 16.4183 11.5817 20 16 20C16.6906 20 17.3608 19.9125 18 19.748" class="il-md-length-40 il-md-duration-3 il-md-delay-2"/><path stroke-linecap="round" d="M16 8C13.7909 8 12 9.79086 12 12C12 14.2091 13.7909 16 16 16C18.2091 16 20 14.2091 20 12C20 9.79086 18.2091 8 16 8Z" class="il-md-length-40 il-md-duration-5 il-md-delay-5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Image(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="66" stroke-dashoffset="66" stroke-width="2" d="M3 14V5H21V19H3V14"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="66;0"/></path><path stroke-dasharray="26" stroke-dashoffset="26" d="M3 16L7 13L10 15L16 10L21 14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.4s" values="26;0"/></path></g><circle cx="7.5" cy="9.5" r="1.5" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.4s" values="0;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path fill="none" stroke-dasharray="66" stroke-dashoffset="66" stroke-width="2" d="M3 14V5H21V19H3V14"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="66;0"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="52" stroke-dashoffset="52" d="M3 16L7 13L10 15L16 10L21 14V19H3Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.8s" values="52;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.15s" values="0;0.3"/></path></g><circle cx="7.5" cy="9.5" r="1.5" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.4s" values="0;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Instagram(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="66" stroke-dashoffset="66" d="M12 3H8C5.23858 3 3 5.23858 3 8V16C3 18.7614 5.23858 21 8 21H16C18.7614 21 21 18.7614 21 16V8C21 5.23858 18.7614 3 16 3z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="66;132"/></path><path stroke-dasharray="26" stroke-dashoffset="26" d="M12 8C14.20914 8 16 9.79086 16 12C16 14.20914 14.20914 16 12 16C9.79086 16 8 14.2091 8 12C8 9.79086 9.79086 8 12 8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.4s" values="26;0"/></path></g><circle cx="17" cy="7" r="1.5" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="1.1s" dur="0.4s" values="0;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laptop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="50" stroke-dashoffset="50" d="M12 17H5V7H19V17Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="50;0"/></path><path stroke-dasharray="20" stroke-dashoffset="20" d="M3 19H21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.3s" values="20;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LaptopTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="50" stroke-dashoffset="50" d="M12 17H5V7H19V17Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="50;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.0s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="20" stroke-dashoffset="20" d="M3 19H21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.3s" values="20;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightDark(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><mask id="lineMdLightDark0"><circle cx="7.5" cy="7.5" r="5.5" fill="#fff"/><circle cx="7.5" cy="7.5" r="5.5"><animate fill="freeze" attributeName="cx" dur="0.4s" values="7.5;11"/><animate fill="freeze" attributeName="r" dur="0.4s" values="5.5;6.5"/></circle></mask><mask id="lineMdLightDark1"><g fill="#fff"><circle cx="12" cy="9" r="5.5"><animate fill="freeze" attributeName="cy" begin="1s" dur="0.5s" values="9;15"/></circle><g fill-opacity="0"><use href="#lineMdLightDark2" transform="rotate(-75 12 15)"/><use href="#lineMdLightDark2" transform="rotate(-25 12 15)"/><use href="#lineMdLightDark2" transform="rotate(25 12 15)"/><use href="#lineMdLightDark2" transform="rotate(75 12 15)"/><set attributeName="fill-opacity" begin="1.5s" to="1"/></g></g><path d="M0 10h26v5h-26z"/><path fill="none" stroke="#fff" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" d="M1 12h22"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="26;52"/></path></mask><symbol id="lineMdLightDark2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="1.5s" dur="0.4s" values="M11 18h2L12 20z;M10.5 21.5h3L12 24z"/></path></symbol></defs><g fill="currentColor"><rect width="13" height="13" x="1" y="1" mask="url(#lineMdLightDark0)"/><path d="M-2 11h28v13h-28z" mask="url(#lineMdLightDark1)" transform="rotate(-45 12 12)"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightDarkLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><mask id="lineMdLightDarkLoop0"><circle cx="7.5" cy="7.5" r="5.5" fill="#fff"/><circle cx="7.5" cy="7.5" r="5.5"><animate fill="freeze" attributeName="cx" dur="0.4s" values="7.5;11"/><animate fill="freeze" attributeName="r" dur="0.4s" values="5.5;6.5"/></circle></mask><mask id="lineMdLightDarkLoop1"><g fill="#fff"><circle cx="12" cy="9" r="5.5"><animate fill="freeze" attributeName="cy" begin="1s" dur="0.5s" values="9;15"/></circle><g><g fill-opacity="0"><use href="#lineMdLightDarkLoop2" transform="rotate(-125 12 15)"/><use href="#lineMdLightDarkLoop2" transform="rotate(-75 12 15)"/><use href="#lineMdLightDarkLoop2" transform="rotate(-25 12 15)"/><use href="#lineMdLightDarkLoop2" transform="rotate(25 12 15)"/><use href="#lineMdLightDarkLoop2" transform="rotate(75 12 15)"/><set attributeName="fill-opacity" begin="1.5s" to="1"/></g><animateTransform attributeName="transform" dur="5s" repeatCount="indefinite" type="rotate" values="0 12 15;50 12 15"/></g></g><path d="M0 10h26v5h-26z"/><path fill="none" stroke="#fff" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" d="M1 12h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 12h22;M2 12h22;M0 12h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="26;52"/></path></mask><symbol id="lineMdLightDarkLoop2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="1.5s" dur="0.4s" values="M11 18h2L12 20z;M10.5 21.5h3L12 24z"/></path></symbol></defs><g fill="currentColor"><rect width="13" height="13" x="1" y="1" mask="url(#lineMdLightDarkLoop0)"/><path d="M-2 11h28v13h-28z" mask="url(#lineMdLightDarkLoop1)" transform="rotate(-45 12 12)"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lightbulb(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="46" stroke-dashoffset="46" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17H9V14.1973C7.2066 13.1599 6 11.2208 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 11.2208 16.7934 13.1599 15 14.1973V17z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="46;0"/></path><rect width="6" height="0" x="9" y="20" fill="currentColor" rx="1"><animate fill="freeze" attributeName="height" begin="0.5s" dur="0.2s" values="0;2"/></rect>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightbulbFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-opacity="0" stroke="currentColor" stroke-dasharray="46" stroke-dashoffset="46" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17H9V14.1973C7.2066 13.1599 6 11.2208 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 11.2208 16.7934 13.1599 15 14.1973V17z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="46;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.5s" values="0;1"/></path><rect width="6" height="0" x="9" y="20" rx="1"><animate fill="freeze" attributeName="height" begin="0.5s" dur="0.2s" values="0;2"/></rect></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightbulbOff(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdLightbulbOff0"><path fill="none" stroke="#fff" stroke-dasharray="46" stroke-dashoffset="46" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17H9V14.1973C7.2066 13.1599 6 11.2208 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 11.2208 16.7934 13.1599 15 14.1973V17z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="46;0"/></path><rect width="6" height="0" x="9" y="20" fill="#fff" rx="1"><animate fill="freeze" attributeName="height" begin="0.5s" dur="0.2s" values="0;2"/></rect><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M1 13h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdLightbulbOff0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightbulbOffFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdLightbulbOffFilled0"><g fill="#fff"><path fill-opacity="0" stroke="#fff" stroke-dasharray="46" stroke-dashoffset="46" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17H9V14.1973C7.2066 13.1599 6 11.2208 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 11.2208 16.7934 13.1599 15 14.1973V17z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="46;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.5s" values="0;1"/></path><rect width="6" height="0" x="9" y="20" rx="1"><animate fill="freeze" attributeName="height" begin="0.5s" dur="0.2s" values="0;2"/></rect></g><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M1 13h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdLightbulbOffFilled0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightbulbOffFilledLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdLightbulbOffFilledLoop0"><g fill="#fff"><path fill-opacity="0" stroke="#fff" stroke-dasharray="46" stroke-dashoffset="46" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17H9V14.1973C7.2066 13.1599 6 11.2208 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 11.2208 16.7934 13.1599 15 14.1973V17z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="46;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.5s" values="0;1"/></path><rect width="6" height="0" x="9" y="20" rx="1"><animate fill="freeze" attributeName="height" begin="0.5s" dur="0.2s" values="0;2"/></rect></g><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M0 13h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 13h22;M2 13h22;M0 13h22"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdLightbulbOffFilledLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightbulbOffLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdLightbulbOffLoop0"><path fill="none" stroke="#fff" stroke-dasharray="46" stroke-dashoffset="46" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17H9V14.1973C7.2066 13.1599 6 11.2208 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 11.2208 16.7934 13.1599 15 14.1973V17z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="46;0"/></path><rect width="6" height="0" x="9" y="20" fill="#fff" rx="1"><animate fill="freeze" attributeName="height" begin="0.5s" dur="0.2s" values="0;2"/></rect><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M0 13h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 13h22;M2 13h22;M0 13h22"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdLightbulbOffLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightbulbOffTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdLightbulbOffTwotone0"><g fill="#fff"><path fill-opacity="0" stroke="#fff" stroke-dasharray="46" stroke-dashoffset="46" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17H9V14.1973C7.2066 13.1599 6 11.2208 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 11.2208 16.7934 13.1599 15 14.1973V17z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="46;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.15s" values="0;0.3"/></path><rect width="6" height="0" x="9" y="20" rx="1"><animate fill="freeze" attributeName="height" begin="0.5s" dur="0.2s" values="0;2"/></rect></g><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M1 13h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdLightbulbOffTwotone0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightbulbOffTwotoneLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdLightbulbOffTwotoneLoop0"><g fill="#fff"><path fill-opacity="0" stroke="#fff" stroke-dasharray="46" stroke-dashoffset="46" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17H9V14.1973C7.2066 13.1599 6 11.2208 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 11.2208 16.7934 13.1599 15 14.1973V17z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="46;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.15s" values="0;0.3"/></path><rect width="6" height="0" x="9" y="20" rx="1"><animate fill="freeze" attributeName="height" begin="0.5s" dur="0.2s" values="0;2"/></rect></g><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M0 13h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 13h22;M2 13h22;M0 13h22"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdLightbulbOffTwotoneLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightbulbTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-opacity="0" stroke="currentColor" stroke-dasharray="46" stroke-dashoffset="46" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17H9V14.1973C7.2066 13.1599 6 11.2208 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 11.2208 16.7934 13.1599 15 14.1973V17z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="46;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.15s" values="0;0.3"/></path><rect width="6" height="0" x="9" y="20" rx="1"><animate fill="freeze" attributeName="height" begin="0.5s" dur="0.2s" values="0;2"/></rect></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linkedin(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="4" cy="4" r="2" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" dur="0.4s" values="0;1"/></circle><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-dasharray="12" stroke-dashoffset="12" d="M4 10V20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M10 10V20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="24" stroke-dashoffset="24" d="M10 15C10 12.2386 12.2386 10 15 10C17.7614 10 20 12.2386 20 15V20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.5s" values="24;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func List(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="14" stroke-dashoffset="14" stroke-linecap="round" stroke-width="2"><path d="M8 5H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.1s" dur="0.2s" values="14;0"/></path><path d="M8 10H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="14;0"/></path><path d="M8 15H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="14;0"/></path><path d="M8 20H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="14;0"/></path></g><g fill="currentColor" fill-opacity="0"><circle cx="4" cy="5" r="1"><animate fill="freeze" attributeName="fill-opacity" dur="0.2s" values="0;1"/></circle><circle cx="4" cy="10" r="1"><animate fill="freeze" attributeName="fill-opacity" begin="0.3s" dur="0.2s" values="0;1"/></circle><circle cx="4" cy="15" r="1"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></circle><circle cx="4" cy="20" r="1"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.2s" values="0;1"/></circle></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListIndented(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M8 5H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.1s" dur="0.2s" values="14;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M10 10H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M12 15H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="10;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M14 20H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path></g><g fill="currentColor" fill-opacity="0"><circle cx="4" cy="5" r="1"><animate fill="freeze" attributeName="fill-opacity" dur="0.2s" values="0;1"/></circle><circle cx="6" cy="10" r="1"><animate fill="freeze" attributeName="fill-opacity" begin="0.3s" dur="0.2s" values="0;1"/></circle><circle cx="8" cy="15" r="1"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></circle><circle cx="10" cy="20" r="1"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.2s" values="0;1"/></circle></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListThree(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round"><g stroke-dasharray="10" stroke-dashoffset="10"><circle cx="5" cy="5" r="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/></circle><circle cx="5" cy="12" r="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="10;0"/></circle><circle cx="5" cy="19" r="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.4s" dur="0.2s" values="10;0"/></circle></g><g stroke-dasharray="28" stroke-dashoffset="28"><rect width="11" height="3" x="9.5" y="3.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.1s" dur="0.5s" values="28;0"/></rect><rect width="11" height="3" x="9.5" y="10.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.5s" values="28;0"/></rect><rect width="11" height="3" x="9.5" y="17.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.5s" dur="0.5s" values="28;0"/></rect></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListThreeFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-linecap="round"><g stroke-dasharray="10" stroke-dashoffset="10"><circle cx="5" cy="5" r="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.0s" dur="0.5s" values="0;1"/></circle><circle cx="5" cy="12" r="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="10;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.2s" dur="0.5s" values="0;1"/></circle><circle cx="5" cy="19" r="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.4s" dur="0.2s" values="10;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.4s" dur="0.5s" values="0;1"/></circle></g><g stroke-dasharray="28" stroke-dashoffset="28"><rect width="11" height="3" x="9.5" y="3.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.1s" dur="0.5s" values="28;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.0s" dur="0.5s" values="0;1"/></rect><rect width="11" height="3" x="9.5" y="10.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.5s" values="28;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.2s" dur="0.5s" values="0;1"/></rect><rect width="11" height="3" x="9.5" y="17.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.5s" dur="0.5s" values="28;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.4s" dur="0.5s" values="0;1"/></rect></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListThreeTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-linecap="round"><g stroke-dasharray="10" stroke-dashoffset="10"><circle cx="5" cy="5" r="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.0s" dur="0.15s" values="0;0.3"/></circle><circle cx="5" cy="12" r="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="10;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.2s" dur="0.15s" values="0;0.3"/></circle><circle cx="5" cy="19" r="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.4s" dur="0.2s" values="10;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.4s" dur="0.15s" values="0;0.3"/></circle></g><g stroke-dasharray="28" stroke-dashoffset="28"><rect width="11" height="3" x="9.5" y="3.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.1s" dur="0.5s" values="28;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.0s" dur="0.15s" values="0;0.3"/></rect><rect width="11" height="3" x="9.5" y="10.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.5s" values="28;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.2s" dur="0.15s" values="0;0.3"/></rect><rect width="11" height="3" x="9.5" y="17.5" rx="1.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.5s" dur="0.5s" values="28;0"/><animate fill="freeze" attributeName="fill-opacity" begin="2.4s" dur="0.15s" values="0;0.3"/></rect></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LoadingAltLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="12" cy="3.5" r="1.5" fill="currentColor" opacity="0"><animateTransform attributeName="transform" calcMode="discrete" dur="2.4s" repeatCount="indefinite" type="rotate" values="0 12 12;90 12 12;180 12 12;270 12 12"/><animate attributeName="opacity" dur="0.6s" keyTimes="0;0.5;1" repeatCount="indefinite" values="1;1;0"/></circle><circle cx="12" cy="3.5" r="1.5" fill="currentColor" opacity="0"><animateTransform attributeName="transform" begin="0.2s" calcMode="discrete" dur="2.4s" repeatCount="indefinite" type="rotate" values="30 12 12;120 12 12;210 12 12;300 12 12"/><animate attributeName="opacity" begin="0.2s" dur="0.6s" keyTimes="0;0.5;1" repeatCount="indefinite" values="1;1;0"/></circle><circle cx="12" cy="3.5" r="1.5" fill="currentColor" opacity="0"><animateTransform attributeName="transform" begin="0.4s" calcMode="discrete" dur="2.4s" repeatCount="indefinite" type="rotate" values="60 12 12;150 12 12;240 12 12;330 12 12"/><animate attributeName="opacity" begin="0.4s" dur="0.6s" keyTimes="0;0.5;1" repeatCount="indefinite" values="1;1;0"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LoadingLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="15" stroke-dashoffset="15" stroke-linecap="round" stroke-width="2" d="M12 3C16.9706 3 21 7.02944 21 12"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="15;0"/><animateTransform attributeName="transform" dur="1.5s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LoadingTwotoneLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" stroke-opacity=".3" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="1.3s" values="60;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M12 3C16.9706 3 21 7.02944 21 12"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="15;0"/><animateTransform attributeName="transform" dur="1.5s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogIn(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="46" stroke-dashoffset="46" d="M8 5V4C8 3.44772 8.44772 3 9 3H18C18.5523 3 19 3.44772 19 4V20C19 20.5523 18.5523 21 18 21H9C8.44771 21 8 20.5523 8 20V19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="46;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M4 12h11" opacity="0"><set attributeName="opacity" begin="0.6s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M15 12l-3.5 -3.5M15 12l-3.5 3.5" opacity="0"><set attributeName="opacity" begin="0.8s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogOut(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="46" stroke-dashoffset="46" d="M16 5V4C16 3.44772 15.5523 3 15 3H6C5.44771 3 5 3.44772 5 4V20C5 20.5523 5.44772 21 6 21H15C15.5523 21 16 20.5523 16 20V19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="46;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M10 12h11" opacity="0"><set attributeName="opacity" begin="0.6s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M21 12l-3.5 -3.5M21 12l-3.5 3.5" opacity="0"><set attributeName="opacity" begin="0.8s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Login(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="32" stroke-dashoffset="32" d="M13 4L20 4C20.5523 4 21 4.44772 21 5V19C21 19.5523 20.5523 20 20 20H13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="32;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M3 12h11.5" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M14.5 12l-3.5 -3.5M14.5 12l-3.5 3.5" opacity="0"><set attributeName="opacity" begin="0.7s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Logout(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="32" stroke-dashoffset="32" d="M12 4H5C4.44772 4 4 4.44772 4 5V19C4 19.5523 4.44772 20 5 20H12"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="32;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M9 12h11.5" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M20.5 12l-3.5 -3.5M20.5 12l-3.5 3.5" opacity="0"><set attributeName="opacity" begin="0.7s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarker(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill="none" stroke="currentColor" stroke-dasharray="48" stroke-dashoffset="48" stroke-linecap="round" stroke-width="2" d="M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="48;0"/></path><circle cx="12" cy="9" r="2.5" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.4s" values="0;1"/></circle></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerAlt(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"><animate fill="freeze" attributeName="d" dur="0.4s" keyTimes="0;0.7;1" values="M12 20.5C12 20.5 11 19 11 18C11 17.5 11.5 17 12 17C12.5 17 13 17.5 13 18C13 19 12 20.5 12 20.5z;M12 20.5C12 20.5 5 13 5 8C5 4.5 8 1 12 1C16 1 19 4.5 19 8C19 13 12 20.5 12 20.5z;M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"/></path><circle cx="12" cy="9" r="2.5" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.4s" values="0;1"/></circle></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerAltFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdMapMarkerAltFilled0"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-width="2" d="M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"><animate fill="freeze" attributeName="d" dur="0.4s" keyTimes="0;0.7;1" values="M12 20.5C12 20.5 11 19 11 18C11 17.5 11.5 17 12 17C12.5 17 13 17.5 13 18C13 19 12 20.5 12 20.5z;M12 20.5C12 20.5 5 13 5 8C5 4.5 8 1 12 1C16 1 19 4.5 19 8C19 13 12 20.5 12 20.5z;M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"/></path><circle cx="12" cy="9" r="2.5" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.3s" dur="0.15s" values="0;1"/></circle></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdMapMarkerAltFilled0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerAltTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-opacity="0"><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"><animate fill="freeze" attributeName="d" dur="0.4s" keyTimes="0;0.7;1" values="M12 20.5C12 20.5 11 19 11 18C11 17.5 11.5 17 12 17C12.5 17 13 17.5 13 18C13 19 12 20.5 12 20.5z;M12 20.5C12 20.5 5 13 5 8C5 4.5 8 1 12 1C16 1 19 4.5 19 8C19 13 12 20.5 12 20.5z;M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"/><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.15s" values="0;0.3"/></path><circle cx="12" cy="9" r="2.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.4s" values="0;1"/></circle></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdMapMarkerFilled0"><path fill="#fff" fill-opacity="0" stroke="#fff" stroke-dasharray="48" stroke-dashoffset="48" stroke-linecap="round" stroke-width="2" d="M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="48;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.4s" values="0;1"/></path><circle cx="12" cy="9" r="2.5"/></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdMapMarkerFilled0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerMultipleAlt(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><symbol id="lineMdMapMarkerMultipleAlt0"><path stroke-linecap="round" d="M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"><animate fill="freeze" attributeName="d" dur="0.4s" keyTimes="0;0.7;1" values="M12 20.5C12 20.5 11 19 11 18C11 17.5 11.5 17 12 17C12.5 17 13 17.5 13 18C13 19 12 20.5 12 20.5z;M12 20.5C12 20.5 5 13 5 8C5 4.5 8 1 12 1C16 1 19 4.5 19 8C19 13 12 20.5 12 20.5z;M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"/></path></symbol><mask id="lineMdMapMarkerMultipleAlt1"><use x="-2" fill="none" stroke="#fff" stroke-width="2" href="#lineMdMapMarkerMultipleAlt0"/><use x="2" stroke="#000" stroke-width="4" href="#lineMdMapMarkerMultipleAlt0"/><use x="2" fill="none" stroke="#fff" stroke-width="2" href="#lineMdMapMarkerMultipleAlt0"/></mask></defs><rect width="24" height="24" fill="currentColor" mask="url(#lineMdMapMarkerMultipleAlt1)"/><circle cx="14" cy="9" r="2.5" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.4s" values="0;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerMultipleAltFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><symbol id="lineMdMapMarkerMultipleAltFilled0"><path stroke-linecap="round" d="M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"><animate fill="freeze" attributeName="d" dur="0.4s" keyTimes="0;0.7;1" values="M12 20.5C12 20.5 11 19 11 18C11 17.5 11.5 17 12 17C12.5 17 13 17.5 13 18C13 19 12 20.5 12 20.5z;M12 20.5C12 20.5 5 13 5 8C5 4.5 8 1 12 1C16 1 19 4.5 19 8C19 13 12 20.5 12 20.5z;M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"/></path></symbol><mask id="lineMdMapMarkerMultipleAltFilled1"><use x="-2" fill="none" stroke="#fff" stroke-width="2" href="#lineMdMapMarkerMultipleAltFilled0"/><use x="2" stroke="#000" stroke-width="4" href="#lineMdMapMarkerMultipleAltFilled0"/><use x="2" fill="#fff" stroke="#fff" stroke-width="2" href="#lineMdMapMarkerMultipleAltFilled0"/><circle cx="14" cy="9" r="2.5" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.3s" dur="0.15s" values="0;1"/></circle></mask></defs><rect width="24" height="24" fill="currentColor" mask="url(#lineMdMapMarkerMultipleAltFilled1)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerMultipleAltTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><symbol id="lineMdMapMarkerMultipleAltTwotone0"><path stroke-linecap="round" d="M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"><animate fill="freeze" attributeName="d" dur="0.4s" keyTimes="0;0.7;1" values="M12 20.5C12 20.5 11 19 11 18C11 17.5 11.5 17 12 17C12.5 17 13 17.5 13 18C13 19 12 20.5 12 20.5z;M12 20.5C12 20.5 5 13 5 8C5 4.5 8 1 12 1C16 1 19 4.5 19 8C19 13 12 20.5 12 20.5z;M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"/></path></symbol><mask id="lineMdMapMarkerMultipleAltTwotone1"><use x="-2" fill="none" stroke="#fff" stroke-width="2" href="#lineMdMapMarkerMultipleAltTwotone0"/><use x="2" stroke="#000" stroke-width="4" href="#lineMdMapMarkerMultipleAltTwotone0"/><use x="2" fill="#fff" fill-opacity="0" stroke="#fff" stroke-width="2" href="#lineMdMapMarkerMultipleAltTwotone0"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.15s" values="0;0.3"/></use></mask></defs><rect width="24" height="24" fill="currentColor" mask="url(#lineMdMapMarkerMultipleAltTwotone1)"/><circle cx="14" cy="9" r="2.5" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.4s" values="0;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerOff(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdMapMarkerOff0"><path fill="none" stroke="#fff" stroke-dasharray="48" stroke-dashoffset="48" stroke-linecap="round" stroke-width="2" d="M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="48;0"/></path><circle cx="12" cy="9" r="2.5" fill="#fff" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.4s" values="0;1"/></circle><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 12 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M1 13h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdMapMarkerOff0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerOffAlt(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdMapMarkerOffAlt0"><path fill="none" stroke="#fff" stroke-linecap="round" stroke-width="2" d="M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"><animate fill="freeze" attributeName="d" dur="0.4s" keyTimes="0;0.7;1" values="M12 20.5C12 20.5 11 19 11 18C11 17.5 11.5 17 12 17C12.5 17 13 17.5 13 18C13 19 12 20.5 12 20.5z;M12 20.5C12 20.5 5 13 5 8C5 4.5 8 1 12 1C16 1 19 4.5 19 8C19 13 12 20.5 12 20.5z;M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"/></path><circle cx="12" cy="9" r="2.5" fill="#fff" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.4s" values="0;1"/></circle><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 12 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M1 13h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdMapMarkerOffAlt0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerOffAltFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdMapMarkerOffAltFilled0"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-width="2" d="M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"><animate fill="freeze" attributeName="d" dur="0.4s" keyTimes="0;0.7;1" values="M12 20.5C12 20.5 11 19 11 18C11 17.5 11.5 17 12 17C12.5 17 13 17.5 13 18C13 19 12 20.5 12 20.5z;M12 20.5C12 20.5 5 13 5 8C5 4.5 8 1 12 1C16 1 19 4.5 19 8C19 13 12 20.5 12 20.5z;M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"/></path><circle cx="12" cy="9" r="2.5" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.3s" dur="0.15s" values="0;1"/></circle><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 12 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M1 13h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdMapMarkerOffAltFilled0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerOffAltFilledLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdMapMarkerOffAltFilledLoop0"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-width="2" d="M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"><animate fill="freeze" attributeName="d" dur="0.4s" keyTimes="0;0.7;1" values="M12 20.5C12 20.5 11 19 11 18C11 17.5 11.5 17 12 17C12.5 17 13 17.5 13 18C13 19 12 20.5 12 20.5z;M12 20.5C12 20.5 5 13 5 8C5 4.5 8 1 12 1C16 1 19 4.5 19 8C19 13 12 20.5 12 20.5z;M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"/></path><circle cx="12" cy="9" r="2.5" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.3s" dur="0.15s" values="0;1"/></circle><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 12 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M0 13h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 13h22;M2 13h22;M0 13h22"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdMapMarkerOffAltFilledLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerOffAltLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdMapMarkerOffAltLoop0"><path fill="none" stroke="#fff" stroke-linecap="round" stroke-width="2" d="M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"><animate fill="freeze" attributeName="d" dur="0.4s" keyTimes="0;0.7;1" values="M12 20.5C12 20.5 11 19 11 18C11 17.5 11.5 17 12 17C12.5 17 13 17.5 13 18C13 19 12 20.5 12 20.5z;M12 20.5C12 20.5 5 13 5 8C5 4.5 8 1 12 1C16 1 19 4.5 19 8C19 13 12 20.5 12 20.5z;M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"/></path><circle cx="12" cy="9" r="2.5" fill="#fff" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.4s" values="0;1"/></circle><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 12 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M0 13h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 13h22;M2 13h22;M0 13h22"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdMapMarkerOffAltLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerOffAltTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdMapMarkerOffAltTwotone0"><path fill="#fff" fill-opacity="0" stroke="#fff" stroke-linecap="round" stroke-width="2" d="M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"><animate fill="freeze" attributeName="d" dur="0.4s" keyTimes="0;0.7;1" values="M12 20.5C12 20.5 11 19 11 18C11 17.5 11.5 17 12 17C12.5 17 13 17.5 13 18C13 19 12 20.5 12 20.5z;M12 20.5C12 20.5 5 13 5 8C5 4.5 8 1 12 1C16 1 19 4.5 19 8C19 13 12 20.5 12 20.5z;M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"/><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.15s" values="0;0.3"/></path><circle cx="12" cy="9" r="2.5" fill="#fff" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.4s" values="0;1"/></circle><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 12 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M1 13h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdMapMarkerOffAltTwotone0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerOffAltTwotoneLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdMapMarkerOffAltTwotoneLoop0"><path fill="#fff" fill-opacity="0" stroke="#fff" stroke-linecap="round" stroke-width="2" d="M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"><animate fill="freeze" attributeName="d" dur="0.4s" keyTimes="0;0.7;1" values="M12 20.5C12 20.5 11 19 11 18C11 17.5 11.5 17 12 17C12.5 17 13 17.5 13 18C13 19 12 20.5 12 20.5z;M12 20.5C12 20.5 5 13 5 8C5 4.5 8 1 12 1C16 1 19 4.5 19 8C19 13 12 20.5 12 20.5z;M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"/><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.15s" values="0;0.3"/></path><circle cx="12" cy="9" r="2.5" fill="#fff" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.4s" values="0;1"/></circle><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 12 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M0 13h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 13h22;M2 13h22;M0 13h22"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdMapMarkerOffAltTwotoneLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerOffFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdMapMarkerOffFilled0"><path fill="#fff" fill-opacity="0" stroke="#fff" stroke-dasharray="48" stroke-dashoffset="48" stroke-linecap="round" stroke-width="2" d="M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="48;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.4s" values="0;1"/></path><circle cx="12" cy="9" r="2.5"/><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 12 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M1 13h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdMapMarkerOffFilled0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerOffFilledLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdMapMarkerOffFilledLoop0"><path fill="#fff" fill-opacity="0" stroke="#fff" stroke-dasharray="48" stroke-dashoffset="48" stroke-linecap="round" stroke-width="2" d="M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="48;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.4s" values="0;1"/></path><circle cx="12" cy="9" r="2.5"/><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 12 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M0 13h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 13h22;M2 13h22;M0 13h22"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdMapMarkerOffFilledLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerOffLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdMapMarkerOffLoop0"><path fill="none" stroke="#fff" stroke-dasharray="48" stroke-dashoffset="48" stroke-linecap="round" stroke-width="2" d="M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="48;0"/></path><circle cx="12" cy="9" r="2.5" fill="#fff" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.4s" values="0;1"/></circle><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 12 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M0 13h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 13h22;M2 13h22;M0 13h22"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdMapMarkerOffLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerOffTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdMapMarkerOffTwotone0"><path fill="#fff" fill-opacity="0" stroke="#fff" stroke-dasharray="48" stroke-dashoffset="48" stroke-linecap="round" stroke-width="2" d="M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="48;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.15s" values="0;0.3"/></path><circle cx="12" cy="9" r="2.5" fill="#fff" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.4s" values="0;1"/></circle><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 12 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M1 13h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdMapMarkerOffTwotone0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerOffTwotoneLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdMapMarkerOffTwotoneLoop0"><path fill="#fff" fill-opacity="0" stroke="#fff" stroke-dasharray="48" stroke-dashoffset="48" stroke-linecap="round" stroke-width="2" d="M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="48;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.15s" values="0;0.3"/></path><circle cx="12" cy="9" r="2.5" fill="#fff" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.4s" values="0;1"/></circle><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 12 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M0 13h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 13h22;M2 13h22;M0 13h22"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdMapMarkerOffTwotoneLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-opacity="0"><path stroke="currentColor" stroke-dasharray="48" stroke-dashoffset="48" stroke-linecap="round" stroke-width="2" d="M12 20.5C12 20.5 6 13.5 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 13.5 12 20.5 12 20.5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="48;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.15s" values="0;0.3"/></path><circle cx="12" cy="9" r="2.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.4s" values="0;1"/></circle></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Marker(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="44" stroke-dashoffset="44" stroke-width="2" d="M6 14L17 3L21 7L10 18L6 14Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="44;0"/></path><path stroke-dasharray="20" stroke-dashoffset="20" d="M9 17L6.5 19.5H2.5L7 15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="20;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarkerFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path fill="none" stroke-dasharray="44" stroke-dashoffset="44" stroke-width="2" d="M6 14L17 3L21 7L10 18L6 14Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="44;0"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="20" stroke-dashoffset="20" d="M9 17L6.5 19.5H2.5L7 15Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="16;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.5s" values="0;1"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarkerTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path fill="none" stroke-dasharray="44" stroke-dashoffset="44" stroke-width="2" d="M6 14L17 3L21 7L10 18L6 14Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="44;0"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="20" stroke-dashoffset="20" d="M9 17L6.5 19.5H2.5L7 15Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="16;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mastodon(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="80" stroke-dashoffset="80" d="M15.5 21.5C6 23.5 6.5 16.5 7.5 16.5C11 16.5 21 19 21 12.5V8.5C21 4.5 18.5 3 14 3H10C5.5 3 3 4 3 10V13C3 19 5 24 15.5 21.5Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="80;160"/></path><path stroke-dasharray="32" stroke-dashoffset="32" d="M7 13.5L7 8C7 8 7.5 6 9.5 6C11.5 6 12 8 12 8L12 10.5L12 8C12 8 12.5 6 14.5 6C16.5 6 17 8 17 8L17 13.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.4s" values="32;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MastodonFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdMastodonFilled0"><path fill="#fff" d="M17.25 2.315c2.366.346 4.361 2.131 4.67 4.396c.167 1.683.022 4.421.02 4.87c0 .133-.019 1.34-.027 1.468c-.207 3.236-2.247 4.514-4.391 4.922c-.03.008-.063.014-.097.02c-1.36.263-2.815.333-4.197.372c-.33.008-.66.008-.99.008c-1.373 0-2.742-.16-4.077-.479a.046.046 0 0 0-.059.047c.038.43.133.853.281 1.259c.185.47.832 1.597 3.234 1.597a17.89 17.89 0 0 0 4.146-.479a.048.048 0 0 1 .053.025a.047.047 0 0 1 .005.02v1.589a.05.05 0 0 1-.02.038c-.444.318-1.048.5-1.562.661c-.228.071-.459.133-.692.187c-2.12.478-4.331.362-6.388-.333c-1.921-.667-3.882-2.302-4.367-4.266a22.953 22.953 0 0 1-.545-3.233c-.151-1.64-.164-3.282-.229-4.928c-.045-1.148-.019-2.4.226-3.528c.51-2.292 2.61-3.896 4.91-4.233c.4-.058 1.15-.27 4.655-.27h.026c3.503 0 5.016.212 5.415.27Z"/><g fill="none" stroke="#000" stroke-dasharray="18" stroke-dashoffset="18" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12.2 12L12.2 9C12.2 8 13.1 6.5 14.85 6.5C16.6 6.5 17.5 8 17.5 9L17.5 14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.4s" values="18;0"/></path><path d="M6.9 14L6.9 9C6.9 8 7.8 6.45 9.54 6.5C11.3 6.47 12.2 8 12.2 9L12.2 12"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="18;0"/></path></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdMastodonFilled0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MastodonTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="80" stroke-dashoffset="80" d="M15.5 21.5C6 23.5 6.5 16.5 7.5 16.5C11 16.5 21 19 21 12.5V8.5C21 4.5 18.5 3 14 3H10C5.5 3 3 4 3 10V13C3 19 5 24 15.5 21.5Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="80;160"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="32" stroke-dashoffset="32" d="M7 13.5L7 8C7 8 7.5 6 9.5 6C11.5 6 12 8 12 8L12 10.5L12 8C12 8 12.5 6 14.5 6C16.5 6 17 8 17 8L17 13.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.4s" values="32;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MedicalServices(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="16" stroke-dashoffset="16" d="M15 7V4C15 3.44772 14.5523 3 14 3H10C9.44772 3 9 3.44772 9 4V7" opacity="0"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="16;32"/><set attributeName="opacity" begin="0.7s" to="1"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M12 11v6" opacity="0"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.0s" dur="0.2s" values="8;0"/><set attributeName="opacity" begin="1.0s" to="1"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M9 14h6" opacity="0"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="8;0"/><set attributeName="opacity" begin="1.2s" to="1"/></path><path stroke-dasharray="64" stroke-dashoffset="64" d="M9 7H20C20.5523 7 21 7.44772 21 8V20C21 20.5523 20.5523 21 20 21H4C3.44772 21 3 20.5523 3 20V8C3 7.44772 3.44772 7 4 7z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MedicalServicesFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdMedicalServicesFilled0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="16" stroke-dashoffset="16" d="M15 7V4C15 3.44772 14.5523 3 14 3H10C9.44772 3 9 3.44772 9 4V7" opacity="0"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="16;32"/><set attributeName="opacity" begin="0.7s" to="1"/></path><path fill="#fff" fill-opacity="0" stroke-dasharray="64" stroke-dashoffset="64" d="M9 7H20C20.5523 7 21 7.44772 21 8V20C21 20.5523 20.5523 21 20 21H4C3.44772 21 3 20.5523 3 20V8C3 7.44772 3.44772 7 4 7z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.5s" values="0;1"/></path><g stroke="#000" stroke-dasharray="8" stroke-dashoffset="8"><path d="M12 11v6" opacity="0"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.0s" dur="0.2s" values="8;0"/><set attributeName="opacity" begin="1.0s" to="1"/></path><path d="M9 14h6" opacity="0"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="8;0"/><set attributeName="opacity" begin="1.2s" to="1"/></path></g></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdMedicalServicesFilled0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MedicalServicesTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="16" stroke-dashoffset="16" d="M15 7V4C15 3.44772 14.5523 3 14 3H10C9.44772 3 9 3.44772 9 4V7" opacity="0"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="16;32"/><set attributeName="opacity" begin="0.7s" to="1"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M12 11v6" opacity="0"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.0s" dur="0.2s" values="8;0"/><set attributeName="opacity" begin="1.0s" to="1"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M9 14h6" opacity="0"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="8;0"/><set attributeName="opacity" begin="1.2s" to="1"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="64" stroke-dashoffset="64" d="M9 7H20C20.5523 7 21 7.44772 21 8V20C21 20.5523 20.5523 21 20 21H4C3.44772 21 3 20.5523 3 20V8C3 7.44772 3.44772 7 4 7z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.15s" values="0;0.3"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Menu(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="24" stroke-dashoffset="24" stroke-linecap="round" stroke-width="2"><path d="M5 5H19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="24;0"/></path><path d="M5 12H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="24;0"/></path><path d="M5 19H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="24;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuFoldLeft(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="10" stroke-dashoffset="10" d="M7 9L4 12L7 15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="10;0"/></path><path stroke-dasharray="16" stroke-dashoffset="16" d="M19 5H5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="16;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M19 12H10"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="16" stroke-dashoffset="16" d="M19 19H5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="16;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuFoldRight(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="10" stroke-dashoffset="10" d="M17 9L20 12L17 15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="10;0"/></path><path stroke-dasharray="16" stroke-dashoffset="16" d="M5 5H19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="16;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M5 12H14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="16" stroke-dashoffset="16" d="M5 19H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="16;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuToCloseAltTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M5 5L19 5"><animate fill="freeze" attributeName="d" begin="0.2s" dur="0.4s" values="M5 5L19 5;M5 5L19 19"/></path><path d="M5 12H19"><animate fill="freeze" attributeName="d" dur="0.4s" values="M5 12H19;M12 12H12"/><set attributeName="opacity" begin="0.4s" to="0"/></path><path d="M5 19L19 19"><animate fill="freeze" attributeName="d" begin="0.2s" dur="0.4s" values="M5 19L19 19;M5 19L19 5"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuToCloseTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M5 5L12 5L19 5"><animate fill="freeze" attributeName="d" dur="0.4s" values="M5 5L12 5L19 5;M5 5L12 12L19 5"/></path><path d="M5 12H19"><animate fill="freeze" attributeName="d" dur="0.4s" values="M5 12H19;M12 12H12"/></path><path d="M5 19L12 19L19 19"><animate fill="freeze" attributeName="d" dur="0.4s" values="M5 19L12 19L19 19;M5 19L12 12L19 19"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuUnfoldLeft(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="10" stroke-dashoffset="10" d="M21 9L18 12L21 15"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/></path><path stroke-dasharray="16" stroke-dashoffset="16" d="M19 5H5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="16;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M14 12H5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="16" stroke-dashoffset="16" d="M19 19H5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="16;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuUnfoldRight(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="10" stroke-dashoffset="10" d="M3 9L6 12L3 15"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/></path><path stroke-dasharray="16" stroke-dashoffset="16" d="M5 5H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="16;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M10 12H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="16" stroke-dashoffset="16" d="M5 19H19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="16;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path stroke="currentColor" stroke-dasharray="18" stroke-dashoffset="18" stroke-linecap="round" stroke-width="2" d="M5 12H19" fill="currentColor"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="18;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircle(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M7 12H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="12;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircleTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M7 12H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="12;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSquare(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M4 12V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="64;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M7 12H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="12;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSquareTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="64" stroke-dashoffset="64" d="M4 12V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="64;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M7 12H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="12;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Moon(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-opacity="0"><path d="M15.22 6.03L17.75 4.09L14.56 4L13.5 1L12.44 4L9.25 4.09L11.78 6.03L10.87 9.09L13.5 7.28L16.13 9.09L15.22 6.03Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.4s" values="0;1"/></path><path d="M19.61 12.25L21.25 11L19.19 10.95L18.5 9L17.81 10.95L15.75 11L17.39 12.25L16.8 14.23L18.5 13.06L20.2 14.23L19.61 12.25Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.1s" dur="0.4s" values="0;1"/></path></g><path fill="none" stroke="currentColor" stroke-dasharray="56" stroke-dashoffset="56" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="56;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="4" stroke-dashoffset="4" stroke-linecap="round" stroke-linejoin="round"><path d="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"><animate id="lineMdMoonAltLoop0" fill="freeze" attributeName="stroke-dashoffset" begin="0.7s;lineMdMoonAltLoop0.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonAltLoop0.begin+2s;lineMdMoonAltLoop0.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonAltLoop0.begin+1.2s;lineMdMoonAltLoop0.begin+3.2s;lineMdMoonAltLoop0.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonAltLoop0.begin+1.8s" to="M12 5h1.5M12 5h-1.5M12 5v1.5M12 5v-1.5"/><set attributeName="d" begin="lineMdMoonAltLoop0.begin+3.8s" to="M12 4h1.5M12 4h-1.5M12 4v1.5M12 4v-1.5"/><set attributeName="d" begin="lineMdMoonAltLoop0.begin+5.8s" to="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"/></path><path d="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"><animate id="lineMdMoonAltLoop1" fill="freeze" attributeName="stroke-dashoffset" begin="1.1s;lineMdMoonAltLoop1.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonAltLoop1.begin+2s;lineMdMoonAltLoop1.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonAltLoop1.begin+1.2s;lineMdMoonAltLoop1.begin+3.2s;lineMdMoonAltLoop1.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonAltLoop1.begin+1.8s" to="M17 11h1.5M17 11h-1.5M17 11v1.5M17 11v-1.5"/><set attributeName="d" begin="lineMdMoonAltLoop1.begin+3.8s" to="M18 12h1.5M18 12h-1.5M18 12v1.5M18 12v-1.5"/><set attributeName="d" begin="lineMdMoonAltLoop1.begin+5.8s" to="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"/></path><path d="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"><animate id="lineMdMoonAltLoop2" fill="freeze" attributeName="stroke-dashoffset" begin="2.9s;lineMdMoonAltLoop2.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonAltLoop2.begin+2s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonAltLoop2.begin+1.2s;lineMdMoonAltLoop2.begin+3.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonAltLoop2.begin+1.8s" to="M20 5h1.5M20 5h-1.5M20 5v1.5M20 5v-1.5"/><set attributeName="d" begin="lineMdMoonAltLoop2.begin+5.8s" to="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"/></path></g><path fill="none" stroke="currentColor" stroke-dasharray="56" stroke-dashoffset="56" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="56;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltToSunnyOutlineLoopTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="2" stroke-dashoffset="2" stroke-linecap="round" stroke-width="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.6s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.9s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="1.2s" values="2;0"/></path><animateTransform attributeName="transform" dur="30s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></g><mask id="lineMdMoonAltToSunnyOutlineLoopTransition0"><circle cx="12" cy="12" r="12" fill="#fff"/><circle cx="12" cy="12" r="8"><animate fill="freeze" attributeName="r" dur="0.4s" values="8;4"/></circle><circle cx="18" cy="6" r="12" fill="#fff"><animate fill="freeze" attributeName="cx" dur="0.4s" values="18;22"/><animate fill="freeze" attributeName="cy" dur="0.4s" values="6;2"/><animate fill="freeze" attributeName="r" dur="0.4s" values="12;3"/></circle><circle cx="18" cy="6" r="10"><animate fill="freeze" attributeName="cx" dur="0.4s" values="18;22"/><animate fill="freeze" attributeName="cy" dur="0.4s" values="6;2"/><animate fill="freeze" attributeName="r" dur="0.4s" values="10;1"/></circle></mask><circle cx="12" cy="12" r="10" fill="currentColor" mask="url(#lineMdMoonAltToSunnyOutlineLoopTransition0)"><animate fill="freeze" attributeName="r" dur="0.4s" values="10;6"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-opacity="0"><path d="M15.22 6.03L17.75 4.09L14.56 4L13.5 1L12.44 4L9.25 4.09L11.78 6.03L10.87 9.09L13.5 7.28L16.13 9.09L15.22 6.03Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.4s" values="0;1"/></path><path d="M19.61 12.25L21.25 11L19.19 10.95L18.5 9L17.81 10.95L15.75 11L17.39 12.25L16.8 14.23L18.5 13.06L20.2 14.23L19.61 12.25Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.1s" dur="0.4s" values="0;1"/></path></g><g fill-opacity="0" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" stroke-dasharray="56" stroke-dashoffset="56" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="56;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.5s" dur="0.5s" values="0;1"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonFilledAltLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="4" stroke-dashoffset="4" stroke-linecap="round" stroke-linejoin="round"><path d="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"><animate id="lineMdMoonFilledAltLoop0" fill="freeze" attributeName="stroke-dashoffset" begin="0.7s;lineMdMoonFilledAltLoop0.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonFilledAltLoop0.begin+2s;lineMdMoonFilledAltLoop0.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonFilledAltLoop0.begin+1.2s;lineMdMoonFilledAltLoop0.begin+3.2s;lineMdMoonFilledAltLoop0.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonFilledAltLoop0.begin+1.8s" to="M12 5h1.5M12 5h-1.5M12 5v1.5M12 5v-1.5"/><set attributeName="d" begin="lineMdMoonFilledAltLoop0.begin+3.8s" to="M12 4h1.5M12 4h-1.5M12 4v1.5M12 4v-1.5"/><set attributeName="d" begin="lineMdMoonFilledAltLoop0.begin+5.8s" to="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"/></path><path d="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"><animate id="lineMdMoonFilledAltLoop1" fill="freeze" attributeName="stroke-dashoffset" begin="1.1s;lineMdMoonFilledAltLoop1.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonFilledAltLoop1.begin+2s;lineMdMoonFilledAltLoop1.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonFilledAltLoop1.begin+1.2s;lineMdMoonFilledAltLoop1.begin+3.2s;lineMdMoonFilledAltLoop1.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonFilledAltLoop1.begin+1.8s" to="M17 11h1.5M17 11h-1.5M17 11v1.5M17 11v-1.5"/><set attributeName="d" begin="lineMdMoonFilledAltLoop1.begin+3.8s" to="M18 12h1.5M18 12h-1.5M18 12v1.5M18 12v-1.5"/><set attributeName="d" begin="lineMdMoonFilledAltLoop1.begin+5.8s" to="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"/></path><path d="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"><animate id="lineMdMoonFilledAltLoop2" fill="freeze" attributeName="stroke-dashoffset" begin="2.9s;lineMdMoonFilledAltLoop2.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonFilledAltLoop2.begin+2s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonFilledAltLoop2.begin+1.2s;lineMdMoonFilledAltLoop2.begin+3.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonFilledAltLoop2.begin+1.8s" to="M20 5h1.5M20 5h-1.5M20 5v1.5M20 5v-1.5"/><set attributeName="d" begin="lineMdMoonFilledAltLoop2.begin+5.8s" to="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"/></path></g><g fill-opacity="0" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" stroke-dasharray="56" stroke-dashoffset="56" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="56;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.5s" dur="0.5s" values="0;1"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonFilledAltToSunnyFilledLoopTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="2" stroke-dashoffset="2" stroke-linecap="round" stroke-width="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.6s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.9s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="1.2s" values="2;0"/></path><animateTransform attributeName="transform" dur="30s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></g><mask id="lineMdMoonFilledAltToSunnyFilledLoopTransition0"><circle cx="12" cy="12" r="12" fill="#fff"/><circle cx="18" cy="6" r="12" fill="#fff"><animate fill="freeze" attributeName="cx" dur="0.4s" values="18;22"/><animate fill="freeze" attributeName="cy" dur="0.4s" values="6;2"/><animate fill="freeze" attributeName="r" dur="0.4s" values="12;3"/></circle><circle cx="18" cy="6" r="10"><animate fill="freeze" attributeName="cx" dur="0.4s" values="18;22"/><animate fill="freeze" attributeName="cy" dur="0.4s" values="6;2"/><animate fill="freeze" attributeName="r" dur="0.4s" values="10;1"/></circle></mask><circle cx="12" cy="12" r="10" fill="currentColor" mask="url(#lineMdMoonFilledAltToSunnyFilledLoopTransition0)"><animate fill="freeze" attributeName="r" dur="0.4s" values="10;6"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonFilledLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-opacity="0"><path d="m15.22 6.03l2.53-1.94L14.56 4L13.5 1l-1.06 3l-3.19.09l2.53 1.94l-.91 3.06l2.63-1.81l2.63 1.81z"><animate id="lineMdMoonFilledLoop0" fill="freeze" attributeName="fill-opacity" begin="0.7s;lineMdMoonFilledLoop0.begin+6s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonFilledLoop0.begin+2.2s" dur="0.4s" values="1;0"/></path><path d="M13.61 5.25L15.25 4l-2.06-.05L12.5 2l-.69 1.95L9.75 4l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonFilledLoop0.begin+3s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonFilledLoop0.begin+5.2s" dur="0.4s" values="1;0"/></path><path d="M19.61 12.25L21.25 11l-2.06-.05L18.5 9l-.69 1.95l-2.06.05l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonFilledLoop0.begin+0.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonFilledLoop0.begin+2.8s" dur="0.4s" values="1;0"/></path><path d="m20.828 9.731l1.876-1.439l-2.366-.067L19.552 6l-.786 2.225l-2.366.067l1.876 1.439L17.601 12l1.951-1.342L21.503 12z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonFilledLoop0.begin+3.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonFilledLoop0.begin+5.6s" dur="0.4s" values="1;0"/></path></g><g fill-opacity="0" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" stroke-dasharray="56" stroke-dashoffset="56" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="56;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.5s" dur="0.5s" values="0;1"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonFilledToSunnyFilledLoopTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="2" stroke-dashoffset="2" stroke-linecap="round" stroke-width="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="1.2s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="1.5s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.5s" dur="1.2s" values="2;0"/></path><animateTransform attributeName="transform" dur="30s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></g><g fill="currentColor"><path d="M15.22 6.03L17.75 4.09L14.56 4L13.5 1L12.44 4L9.25 4.09L11.78 6.03L10.87 9.09L13.5 7.28L16.13 9.09L15.22 6.03Z"><animate fill="freeze" attributeName="fill-opacity" dur="0.4s" values="1;0"/></path><path d="M19.61 12.25L21.25 11L19.19 10.95L18.5 9L17.81 10.95L15.75 11L17.39 12.25L16.8 14.23L18.5 13.06L20.2 14.23L19.61 12.25Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.2s" dur="0.4s" values="1;0"/></path></g><g fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z"/><set attributeName="opacity" begin="0.6s" to="0"/></g><mask id="lineMdMoonFilledToSunnyFilledLoopTransition0"><circle cx="12" cy="12" r="12" fill="#fff"/><circle cx="18" cy="6" r="12" fill="#fff"><animate fill="freeze" attributeName="cx" begin="0.6s" dur="0.4s" values="18;22"/><animate fill="freeze" attributeName="cy" begin="0.6s" dur="0.4s" values="6;2"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="12;3"/></circle><circle cx="18" cy="6" r="10"><animate fill="freeze" attributeName="cx" begin="0.6s" dur="0.4s" values="18;22"/><animate fill="freeze" attributeName="cy" begin="0.6s" dur="0.4s" values="6;2"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="10;1"/></circle></mask><circle cx="12" cy="12" r="10" fill="currentColor" mask="url(#lineMdMoonFilledToSunnyFilledLoopTransition0)" opacity="0"><set attributeName="opacity" begin="0.6s" to="1"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="10;6"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonFilledToSunnyFilledTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="2" stroke-dashoffset="2" stroke-linecap="round" stroke-width="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="1.2s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="1.5s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.5s" dur="1.2s" values="2;0"/></path></g><g fill="currentColor"><path d="M15.22 6.03L17.75 4.09L14.56 4L13.5 1L12.44 4L9.25 4.09L11.78 6.03L10.87 9.09L13.5 7.28L16.13 9.09L15.22 6.03Z"><animate fill="freeze" attributeName="fill-opacity" dur="0.4s" values="1;0"/></path><path d="M19.61 12.25L21.25 11L19.19 10.95L18.5 9L17.81 10.95L15.75 11L17.39 12.25L16.8 14.23L18.5 13.06L20.2 14.23L19.61 12.25Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.2s" dur="0.4s" values="1;0"/></path></g><g fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z"/><set attributeName="opacity" begin="0.6s" to="0"/></g><mask id="lineMdMoonFilledToSunnyFilledTransition0"><circle cx="12" cy="12" r="12" fill="#fff"/><circle cx="18" cy="6" r="12" fill="#fff"><animate fill="freeze" attributeName="cx" begin="0.6s" dur="0.4s" values="18;22"/><animate fill="freeze" attributeName="cy" begin="0.6s" dur="0.4s" values="6;2"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="12;3"/></circle><circle cx="18" cy="6" r="10"><animate fill="freeze" attributeName="cx" begin="0.6s" dur="0.4s" values="18;22"/><animate fill="freeze" attributeName="cy" begin="0.6s" dur="0.4s" values="6;2"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="10;1"/></circle></mask><circle cx="12" cy="12" r="10" fill="currentColor" mask="url(#lineMdMoonFilledToSunnyFilledTransition0)" opacity="0"><set attributeName="opacity" begin="0.6s" to="1"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="10;6"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-opacity="0"><path d="m15.22 6.03l2.53-1.94L14.56 4L13.5 1l-1.06 3l-3.19.09l2.53 1.94l-.91 3.06l2.63-1.81l2.63 1.81z"><animate id="lineMdMoonLoop0" fill="freeze" attributeName="fill-opacity" begin="0.7s;lineMdMoonLoop0.begin+6s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonLoop0.begin+2.2s" dur="0.4s" values="1;0"/></path><path d="M13.61 5.25L15.25 4l-2.06-.05L12.5 2l-.69 1.95L9.75 4l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonLoop0.begin+3s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonLoop0.begin+5.2s" dur="0.4s" values="1;0"/></path><path d="M19.61 12.25L21.25 11l-2.06-.05L18.5 9l-.69 1.95l-2.06.05l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonLoop0.begin+0.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonLoop0.begin+2.8s" dur="0.4s" values="1;0"/></path><path d="m20.828 9.731l1.876-1.439l-2.366-.067L19.552 6l-.786 2.225l-2.366.067l1.876 1.439L17.601 12l1.951-1.342L21.503 12z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonLoop0.begin+3.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonLoop0.begin+5.6s" dur="0.4s" values="1;0"/></path></g><path fill="none" stroke="currentColor" stroke-dasharray="56" stroke-dashoffset="56" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="56;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonRisingAltLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="4" stroke-dashoffset="4" stroke-linecap="round" stroke-linejoin="round"><path d="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"><animate id="lineMdMoonRisingAltLoop0" fill="freeze" attributeName="stroke-dashoffset" begin="0.7s;lineMdMoonRisingAltLoop0.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingAltLoop0.begin+2s;lineMdMoonRisingAltLoop0.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingAltLoop0.begin+1.2s;lineMdMoonRisingAltLoop0.begin+3.2s;lineMdMoonRisingAltLoop0.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonRisingAltLoop0.begin+1.8s" to="M12 5h1.5M12 5h-1.5M12 5v1.5M12 5v-1.5"/><set attributeName="d" begin="lineMdMoonRisingAltLoop0.begin+3.8s" to="M12 4h1.5M12 4h-1.5M12 4v1.5M12 4v-1.5"/><set attributeName="d" begin="lineMdMoonRisingAltLoop0.begin+5.8s" to="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"/></path><path d="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"><animate id="lineMdMoonRisingAltLoop1" fill="freeze" attributeName="stroke-dashoffset" begin="1.1s;lineMdMoonRisingAltLoop1.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingAltLoop1.begin+2s;lineMdMoonRisingAltLoop1.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingAltLoop1.begin+1.2s;lineMdMoonRisingAltLoop1.begin+3.2s;lineMdMoonRisingAltLoop1.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonRisingAltLoop1.begin+1.8s" to="M17 11h1.5M17 11h-1.5M17 11v1.5M17 11v-1.5"/><set attributeName="d" begin="lineMdMoonRisingAltLoop1.begin+3.8s" to="M18 12h1.5M18 12h-1.5M18 12v1.5M18 12v-1.5"/><set attributeName="d" begin="lineMdMoonRisingAltLoop1.begin+5.8s" to="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"/></path><path d="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"><animate id="lineMdMoonRisingAltLoop2" fill="freeze" attributeName="stroke-dashoffset" begin="2.9s;lineMdMoonRisingAltLoop2.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingAltLoop2.begin+2s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingAltLoop2.begin+1.2s;lineMdMoonRisingAltLoop2.begin+3.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonRisingAltLoop2.begin+1.8s" to="M20 5h1.5M20 5h-1.5M20 5v1.5M20 5v-1.5"/><set attributeName="d" begin="lineMdMoonRisingAltLoop2.begin+5.8s" to="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"/></path></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z" transform="translate(0 22)"><animateMotion fill="freeze" calcMode="linear" dur="0.6s" path="M0 0v-22"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonRisingFilledAltLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="4" stroke-dashoffset="4" stroke-linecap="round" stroke-linejoin="round"><path d="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"><animate id="lineMdMoonRisingFilledAltLoop0" fill="freeze" attributeName="stroke-dashoffset" begin="0.7s;lineMdMoonRisingFilledAltLoop0.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingFilledAltLoop0.begin+2s;lineMdMoonRisingFilledAltLoop0.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingFilledAltLoop0.begin+1.2s;lineMdMoonRisingFilledAltLoop0.begin+3.2s;lineMdMoonRisingFilledAltLoop0.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonRisingFilledAltLoop0.begin+1.8s" to="M12 5h1.5M12 5h-1.5M12 5v1.5M12 5v-1.5"/><set attributeName="d" begin="lineMdMoonRisingFilledAltLoop0.begin+3.8s" to="M12 4h1.5M12 4h-1.5M12 4v1.5M12 4v-1.5"/><set attributeName="d" begin="lineMdMoonRisingFilledAltLoop0.begin+5.8s" to="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"/></path><path d="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"><animate id="lineMdMoonRisingFilledAltLoop1" fill="freeze" attributeName="stroke-dashoffset" begin="1.1s;lineMdMoonRisingFilledAltLoop1.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingFilledAltLoop1.begin+2s;lineMdMoonRisingFilledAltLoop1.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingFilledAltLoop1.begin+1.2s;lineMdMoonRisingFilledAltLoop1.begin+3.2s;lineMdMoonRisingFilledAltLoop1.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonRisingFilledAltLoop1.begin+1.8s" to="M17 11h1.5M17 11h-1.5M17 11v1.5M17 11v-1.5"/><set attributeName="d" begin="lineMdMoonRisingFilledAltLoop1.begin+3.8s" to="M18 12h1.5M18 12h-1.5M18 12v1.5M18 12v-1.5"/><set attributeName="d" begin="lineMdMoonRisingFilledAltLoop1.begin+5.8s" to="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"/></path><path d="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"><animate id="lineMdMoonRisingFilledAltLoop2" fill="freeze" attributeName="stroke-dashoffset" begin="2.9s;lineMdMoonRisingFilledAltLoop2.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingFilledAltLoop2.begin+2s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingFilledAltLoop2.begin+1.2s;lineMdMoonRisingFilledAltLoop2.begin+3.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonRisingFilledAltLoop2.begin+1.8s" to="M20 5h1.5M20 5h-1.5M20 5v1.5M20 5v-1.5"/><set attributeName="d" begin="lineMdMoonRisingFilledAltLoop2.begin+5.8s" to="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"/></path></g><path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z" transform="translate(0 22)"><animateMotion fill="freeze" calcMode="linear" dur="0.6s" path="M0 0v-22"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonRisingFilledLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-opacity="0"><path d="m15.22 6.03l2.53-1.94L14.56 4L13.5 1l-1.06 3l-3.19.09l2.53 1.94l-.91 3.06l2.63-1.81l2.63 1.81z"><animate id="lineMdMoonRisingFilledLoop0" fill="freeze" attributeName="fill-opacity" begin="0.7s;lineMdMoonRisingFilledLoop0.begin+6s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingFilledLoop0.begin+2.2s" dur="0.4s" values="1;0"/></path><path d="M13.61 5.25L15.25 4l-2.06-.05L12.5 2l-.69 1.95L9.75 4l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingFilledLoop0.begin+3s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingFilledLoop0.begin+5.2s" dur="0.4s" values="1;0"/></path><path d="M19.61 12.25L21.25 11l-2.06-.05L18.5 9l-.69 1.95l-2.06.05l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingFilledLoop0.begin+0.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingFilledLoop0.begin+2.8s" dur="0.4s" values="1;0"/></path><path d="m20.828 9.731l1.876-1.439l-2.366-.067L19.552 6l-.786 2.225l-2.366.067l1.876 1.439L17.601 12l1.951-1.342L21.503 12z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingFilledLoop0.begin+3.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingFilledLoop0.begin+5.6s" dur="0.4s" values="1;0"/></path></g><path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z" transform="translate(0 22)"><animateMotion fill="freeze" calcMode="linear" dur="0.6s" path="M0 0v-22"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonRisingLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-opacity="0"><path d="m15.22 6.03l2.53-1.94L14.56 4L13.5 1l-1.06 3l-3.19.09l2.53 1.94l-.91 3.06l2.63-1.81l2.63 1.81z"><animate id="lineMdMoonRisingLoop0" fill="freeze" attributeName="fill-opacity" begin="0.7s;lineMdMoonRisingLoop0.begin+6s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingLoop0.begin+2.2s" dur="0.4s" values="1;0"/></path><path d="M13.61 5.25L15.25 4l-2.06-.05L12.5 2l-.69 1.95L9.75 4l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingLoop0.begin+3s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingLoop0.begin+5.2s" dur="0.4s" values="1;0"/></path><path d="M19.61 12.25L21.25 11l-2.06-.05L18.5 9l-.69 1.95l-2.06.05l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingLoop0.begin+0.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingLoop0.begin+2.8s" dur="0.4s" values="1;0"/></path><path d="m20.828 9.731l1.876-1.439l-2.366-.067L19.552 6l-.786 2.225l-2.366.067l1.876 1.439L17.601 12l1.951-1.342L21.503 12z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingLoop0.begin+3.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingLoop0.begin+5.6s" dur="0.4s" values="1;0"/></path></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z" transform="translate(0 22)"><animateMotion fill="freeze" calcMode="linear" dur="0.6s" path="M0 0v-22"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonRisingTwotoneAltLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="4" stroke-dashoffset="4" stroke-linecap="round" stroke-linejoin="round"><path d="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"><animate id="lineMdMoonRisingTwotoneAltLoop0" fill="freeze" attributeName="stroke-dashoffset" begin="0.7s;lineMdMoonRisingTwotoneAltLoop0.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingTwotoneAltLoop0.begin+2s;lineMdMoonRisingTwotoneAltLoop0.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingTwotoneAltLoop0.begin+1.2s;lineMdMoonRisingTwotoneAltLoop0.begin+3.2s;lineMdMoonRisingTwotoneAltLoop0.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonRisingTwotoneAltLoop0.begin+1.8s" to="M12 5h1.5M12 5h-1.5M12 5v1.5M12 5v-1.5"/><set attributeName="d" begin="lineMdMoonRisingTwotoneAltLoop0.begin+3.8s" to="M12 4h1.5M12 4h-1.5M12 4v1.5M12 4v-1.5"/><set attributeName="d" begin="lineMdMoonRisingTwotoneAltLoop0.begin+5.8s" to="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"/></path><path d="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"><animate id="lineMdMoonRisingTwotoneAltLoop1" fill="freeze" attributeName="stroke-dashoffset" begin="1.1s;lineMdMoonRisingTwotoneAltLoop1.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingTwotoneAltLoop1.begin+2s;lineMdMoonRisingTwotoneAltLoop1.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingTwotoneAltLoop1.begin+1.2s;lineMdMoonRisingTwotoneAltLoop1.begin+3.2s;lineMdMoonRisingTwotoneAltLoop1.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonRisingTwotoneAltLoop1.begin+1.8s" to="M17 11h1.5M17 11h-1.5M17 11v1.5M17 11v-1.5"/><set attributeName="d" begin="lineMdMoonRisingTwotoneAltLoop1.begin+3.8s" to="M18 12h1.5M18 12h-1.5M18 12v1.5M18 12v-1.5"/><set attributeName="d" begin="lineMdMoonRisingTwotoneAltLoop1.begin+5.8s" to="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"/></path><path d="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"><animate id="lineMdMoonRisingTwotoneAltLoop2" fill="freeze" attributeName="stroke-dashoffset" begin="2.9s;lineMdMoonRisingTwotoneAltLoop2.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingTwotoneAltLoop2.begin+2s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonRisingTwotoneAltLoop2.begin+1.2s;lineMdMoonRisingTwotoneAltLoop2.begin+3.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonRisingTwotoneAltLoop2.begin+1.8s" to="M20 5h1.5M20 5h-1.5M20 5v1.5M20 5v-1.5"/><set attributeName="d" begin="lineMdMoonRisingTwotoneAltLoop2.begin+5.8s" to="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"/></path></g><path fill="currentColor" fill-opacity=".3" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z" transform="translate(0 22)"><animateMotion fill="freeze" calcMode="linear" dur="0.6s" path="M0 0v-22"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonRisingTwotoneLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-opacity="0"><path d="m15.22 6.03l2.53-1.94L14.56 4L13.5 1l-1.06 3l-3.19.09l2.53 1.94l-.91 3.06l2.63-1.81l2.63 1.81z"><animate id="lineMdMoonRisingTwotoneLoop0" fill="freeze" attributeName="fill-opacity" begin="0.7s;lineMdMoonRisingTwotoneLoop0.begin+6s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingTwotoneLoop0.begin+2.2s" dur="0.4s" values="1;0"/></path><path d="M13.61 5.25L15.25 4l-2.06-.05L12.5 2l-.69 1.95L9.75 4l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingTwotoneLoop0.begin+3s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingTwotoneLoop0.begin+5.2s" dur="0.4s" values="1;0"/></path><path d="M19.61 12.25L21.25 11l-2.06-.05L18.5 9l-.69 1.95l-2.06.05l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingTwotoneLoop0.begin+0.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingTwotoneLoop0.begin+2.8s" dur="0.4s" values="1;0"/></path><path d="m20.828 9.731l1.876-1.439l-2.366-.067L19.552 6l-.786 2.225l-2.366.067l1.876 1.439L17.601 12l1.951-1.342L21.503 12z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingTwotoneLoop0.begin+3.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonRisingTwotoneLoop0.begin+5.6s" dur="0.4s" values="1;0"/></path></g><path fill="currentColor" fill-opacity=".3" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z" transform="translate(0 22)"><animateMotion fill="freeze" calcMode="linear" dur="0.6s" path="M0 0v-22"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonSimple(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="64" stroke-dashoffset="64" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3C7.03 3 3 7.03 3 12C3 16.97 7.03 21 12 21C15.53 21 18.59 18.96 20.06 16C20.06 16 14 17.5 11 13C8 8.5 12 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonSimpleFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="64" stroke-dashoffset="64" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3C7.03 3 3 7.03 3 12C3 16.97 7.03 21 12 21C15.53 21 18.59 18.96 20.06 16C20.06 16 14 17.5 11 13C8 8.5 12 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.5s" values="0;1"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonSimpleTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="64" stroke-dashoffset="64" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3C7.03 3 3 7.03 3 12C3 16.97 7.03 21 12 21C15.53 21 18.59 18.96 20.06 16C20.06 16 14 17.5 11 13C8 8.5 12 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.15s" values="0;0.3"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonToSunnyOutlineLoopTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="2" stroke-dashoffset="2" stroke-linecap="round" stroke-width="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="1.2s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="1.5s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.5s" dur="1.2s" values="2;0"/></path><animateTransform attributeName="transform" dur="30s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></g><g fill="currentColor"><path d="M15.22 6.03L17.75 4.09L14.56 4L13.5 1L12.44 4L9.25 4.09L11.78 6.03L10.87 9.09L13.5 7.28L16.13 9.09L15.22 6.03Z"><animate fill="freeze" attributeName="fill-opacity" dur="0.4s" values="1;0"/></path><path d="M19.61 12.25L21.25 11L19.19 10.95L18.5 9L17.81 10.95L15.75 11L17.39 12.25L16.8 14.23L18.5 13.06L20.2 14.23L19.61 12.25Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.2s" dur="0.4s" values="1;0"/></path></g><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z"/><set attributeName="opacity" begin="0.6s" to="0"/></g><mask id="lineMdMoonToSunnyOutlineLoopTransition0"><circle cx="12" cy="12" r="12" fill="#fff"/><circle cx="12" cy="12" r="8"><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="8;4"/></circle><circle cx="18" cy="6" r="12" fill="#fff"><animate fill="freeze" attributeName="cx" begin="0.6s" dur="0.4s" values="18;22"/><animate fill="freeze" attributeName="cy" begin="0.6s" dur="0.4s" values="6;2"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="12;3"/></circle><circle cx="18" cy="6" r="10"><animate fill="freeze" attributeName="cx" begin="0.6s" dur="0.4s" values="18;22"/><animate fill="freeze" attributeName="cy" begin="0.6s" dur="0.4s" values="6;2"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="10;1"/></circle></mask><circle cx="12" cy="12" r="10" fill="currentColor" mask="url(#lineMdMoonToSunnyOutlineLoopTransition0)" opacity="0"><set attributeName="opacity" begin="0.6s" to="1"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="10;6"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonToSunnyOutlineTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="2" stroke-dashoffset="2" stroke-linecap="round" stroke-width="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="1.2s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="1.5s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.5s" dur="1.2s" values="2;0"/></path></g><g fill="currentColor"><path d="M15.22 6.03L17.75 4.09L14.56 4L13.5 1L12.44 4L9.25 4.09L11.78 6.03L10.87 9.09L13.5 7.28L16.13 9.09L15.22 6.03Z"><animate fill="freeze" attributeName="fill-opacity" dur="0.4s" values="1;0"/></path><path d="M19.61 12.25L21.25 11L19.19 10.95L18.5 9L17.81 10.95L15.75 11L17.39 12.25L16.8 14.23L18.5 13.06L20.2 14.23L19.61 12.25Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.2s" dur="0.4s" values="1;0"/></path></g><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z"/><set attributeName="opacity" begin="0.6s" to="0"/></g><mask id="lineMdMoonToSunnyOutlineTransition0"><circle cx="12" cy="12" r="12" fill="#fff"/><circle cx="12" cy="12" r="8"><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="8;4"/></circle><circle cx="18" cy="6" r="12" fill="#fff"><animate fill="freeze" attributeName="cx" begin="0.6s" dur="0.4s" values="18;22"/><animate fill="freeze" attributeName="cy" begin="0.6s" dur="0.4s" values="6;2"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="12;3"/></circle><circle cx="18" cy="6" r="10"><animate fill="freeze" attributeName="cx" begin="0.6s" dur="0.4s" values="18;22"/><animate fill="freeze" attributeName="cy" begin="0.6s" dur="0.4s" values="6;2"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="10;1"/></circle></mask><circle cx="12" cy="12" r="10" fill="currentColor" mask="url(#lineMdMoonToSunnyOutlineTransition0)" opacity="0"><set attributeName="opacity" begin="0.6s" to="1"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.4s" values="10;6"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-opacity="0"><path d="M15.22 6.03L17.75 4.09L14.56 4L13.5 1L12.44 4L9.25 4.09L11.78 6.03L10.87 9.09L13.5 7.28L16.13 9.09L15.22 6.03Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.4s" values="0;1"/></path><path d="M19.61 12.25L21.25 11L19.19 10.95L18.5 9L17.81 10.95L15.75 11L17.39 12.25L16.8 14.23L18.5 13.06L20.2 14.23L19.61 12.25Z"><animate fill="freeze" attributeName="fill-opacity" begin="1.1s" dur="0.4s" values="0;1"/></path></g><g fill-opacity="0" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" stroke-dasharray="56" stroke-dashoffset="56" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="56;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.5s" dur="0.15s" values="0;0.3"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonTwotoneAltLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="4" stroke-dashoffset="4" stroke-linecap="round" stroke-linejoin="round"><path d="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"><animate id="lineMdMoonTwotoneAltLoop0" fill="freeze" attributeName="stroke-dashoffset" begin="0.7s;lineMdMoonTwotoneAltLoop0.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonTwotoneAltLoop0.begin+2s;lineMdMoonTwotoneAltLoop0.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonTwotoneAltLoop0.begin+1.2s;lineMdMoonTwotoneAltLoop0.begin+3.2s;lineMdMoonTwotoneAltLoop0.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonTwotoneAltLoop0.begin+1.8s" to="M12 5h1.5M12 5h-1.5M12 5v1.5M12 5v-1.5"/><set attributeName="d" begin="lineMdMoonTwotoneAltLoop0.begin+3.8s" to="M12 4h1.5M12 4h-1.5M12 4v1.5M12 4v-1.5"/><set attributeName="d" begin="lineMdMoonTwotoneAltLoop0.begin+5.8s" to="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"/></path><path d="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"><animate id="lineMdMoonTwotoneAltLoop1" fill="freeze" attributeName="stroke-dashoffset" begin="1.1s;lineMdMoonTwotoneAltLoop1.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonTwotoneAltLoop1.begin+2s;lineMdMoonTwotoneAltLoop1.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonTwotoneAltLoop1.begin+1.2s;lineMdMoonTwotoneAltLoop1.begin+3.2s;lineMdMoonTwotoneAltLoop1.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonTwotoneAltLoop1.begin+1.8s" to="M17 11h1.5M17 11h-1.5M17 11v1.5M17 11v-1.5"/><set attributeName="d" begin="lineMdMoonTwotoneAltLoop1.begin+3.8s" to="M18 12h1.5M18 12h-1.5M18 12v1.5M18 12v-1.5"/><set attributeName="d" begin="lineMdMoonTwotoneAltLoop1.begin+5.8s" to="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"/></path><path d="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"><animate id="lineMdMoonTwotoneAltLoop2" fill="freeze" attributeName="stroke-dashoffset" begin="2.9s;lineMdMoonTwotoneAltLoop2.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonTwotoneAltLoop2.begin+2s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdMoonTwotoneAltLoop2.begin+1.2s;lineMdMoonTwotoneAltLoop2.begin+3.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdMoonTwotoneAltLoop2.begin+1.8s" to="M20 5h1.5M20 5h-1.5M20 5v1.5M20 5v-1.5"/><set attributeName="d" begin="lineMdMoonTwotoneAltLoop2.begin+5.8s" to="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"/></path></g><g fill-opacity="0" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" stroke-dasharray="56" stroke-dashoffset="56" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="56;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.5s" dur="0.15s" values="0;0.3"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonTwotoneLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-opacity="0"><path d="m15.22 6.03l2.53-1.94L14.56 4L13.5 1l-1.06 3l-3.19.09l2.53 1.94l-.91 3.06l2.63-1.81l2.63 1.81z"><animate id="lineMdMoonTwotoneLoop0" fill="freeze" attributeName="fill-opacity" begin="0.7s;lineMdMoonTwotoneLoop0.begin+6s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonTwotoneLoop0.begin+2.2s" dur="0.4s" values="1;0"/></path><path d="M13.61 5.25L15.25 4l-2.06-.05L12.5 2l-.69 1.95L9.75 4l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonTwotoneLoop0.begin+3s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonTwotoneLoop0.begin+5.2s" dur="0.4s" values="1;0"/></path><path d="M19.61 12.25L21.25 11l-2.06-.05L18.5 9l-.69 1.95l-2.06.05l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonTwotoneLoop0.begin+0.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonTwotoneLoop0.begin+2.8s" dur="0.4s" values="1;0"/></path><path d="m20.828 9.731l1.876-1.439l-2.366-.067L19.552 6l-.786 2.225l-2.366.067l1.876 1.439L17.601 12l1.951-1.342L21.503 12z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonTwotoneLoop0.begin+3.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdMoonTwotoneLoop0.begin+5.6s" dur="0.4s" values="1;0"/></path></g><g fill-opacity="0" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" stroke-dasharray="56" stroke-dashoffset="56" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="56;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.5s" dur="0.15s" values="0;0.3"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MyLocation(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="56" stroke-dashoffset="56" d="M12 4C16.4183 4 20 7.58172 20 12C20 16.4183 16.4183 20 12 20C7.58172 20 4 16.4183 4 12C4 7.58172 7.58172 4 12 4Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="56;0"/></path><path d="M12 4v0M20 12h0M12 20v0M4 12h0" opacity="0"><set attributeName="opacity" begin="0.9s" to="1"/><animate fill="freeze" attributeName="d" begin="0.9s" dur="0.2s" values="M12 4v0M20 12h0M12 20v0M4 12h0;M12 4v-2M20 12h2M12 20v2M4 12h-2"/></path></g><circle cx="12" cy="12" r="0" fill="currentColor" fill-opacity="0"><set attributeName="fill-opacity" begin="0.6s" to="1"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.2s" values="0;4"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MyLocationLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="56" stroke-dashoffset="56" d="M12 4C16.4183 4 20 7.58172 20 12C20 16.4183 16.4183 20 12 20C7.58172 20 4 16.4183 4 12C4 7.58172 7.58172 4 12 4Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="56;0"/></path><path d="M12 4v0M20 12h0M12 20v0M4 12h0" opacity="0"><set attributeName="opacity" begin="0.9s" to="1"/><animate fill="freeze" attributeName="d" begin="0.9s" dur="0.2s" values="M12 4v0M20 12h0M12 20v0M4 12h0;M12 4v-2M20 12h2M12 20v2M4 12h-2"/><animateTransform attributeName="transform" dur="30s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></path></g><circle cx="12" cy="12" r="0" fill="currentColor" fill-opacity="0"><set attributeName="fill-opacity" begin="0.6s" to="1"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.2s" values="0;4"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MyLocationOff(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdMyLocationOff0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="56" stroke-dashoffset="56" d="M12 4C16.4183 4 20 7.58172 20 12C20 16.4183 16.4183 20 12 20C7.58172 20 4 16.4183 4 12C4 7.58172 7.58172 4 12 4Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="56;0"/></path><path d="M12 4v0M20 12h0M12 20v0M4 12h0" opacity="0"><set attributeName="opacity" begin="0.9s" to="1"/><animate fill="freeze" attributeName="d" begin="0.9s" dur="0.2s" values="M12 4v0M20 12h0M12 20v0M4 12h0;M12 4v-2M20 12h2M12 20v2M4 12h-2"/></path></g><circle cx="12" cy="12" r="0" fill="#fff" fill-opacity="0"><set attributeName="fill-opacity" begin="0.6s" to="1"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.2s" values="0;4"/></circle><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 12 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M1 13h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdMyLocationOff0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MyLocationOffLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdMyLocationOffLoop0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="56" stroke-dashoffset="56" d="M12 4C16.4183 4 20 7.58172 20 12C20 16.4183 16.4183 20 12 20C7.58172 20 4 16.4183 4 12C4 7.58172 7.58172 4 12 4Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="56;0"/></path><path d="M12 4v0M20 12h0M12 20v0M4 12h0" opacity="0"><set attributeName="opacity" begin="0.9s" to="1"/><animate fill="freeze" attributeName="d" begin="0.9s" dur="0.2s" values="M12 4v0M20 12h0M12 20v0M4 12h0;M12 4v-2M20 12h2M12 20v2M4 12h-2"/><animateTransform attributeName="transform" dur="30s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></path></g><circle cx="12" cy="12" r="0" fill="#fff" fill-opacity="0"><set attributeName="fill-opacity" begin="0.6s" to="1"/><animate fill="freeze" attributeName="r" begin="0.6s" dur="0.2s" values="0;4"/></circle><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 12 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M0 13h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 13h22;M2 13h22;M0 13h22"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdMyLocationOffLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NavigationLeftDown(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="26" stroke-dashoffset="26" d="M21 9H12C9.23858 9 7 11.2386 7 14V20"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="26;0"/></path><g stroke-dasharray="8" stroke-dashoffset="8"><path d="M7 21L3 17"/><path d="M7 21L11 17"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="8;0"/></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NavigationLeftUp(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="26" stroke-dashoffset="26" d="M21 15H12C9.23858 15 7 12.7614 7 10V4"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="26;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M7 3L3 7M7 3L11 7"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NavigationRightDown(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="26" stroke-dashoffset="26" d="M3 9H12C14.76142 9 17 11.2386 17 14V20"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="26;0"/></path><g stroke-dasharray="8" stroke-dashoffset="8"><path d="M17 21L21 17"/><path d="M17 21L13 17"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="8;0"/></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NavigationRightUp(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="26" stroke-dashoffset="26" d="M3 15H12C14.76142 15 17 12.7614 17 10V4"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="26;0"/></path><g stroke-dasharray="8" stroke-dashoffset="8"><path d="M17 3L21 7"/><path d="M17 3L13 7"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="8;0"/></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaintDrop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="28" stroke-dashoffset="28" stroke-linecap="round" stroke-width="2" d="M12 3C12 3 19 9 19 15C19 17 18 21 12 21M12 3C12 3 5 9 5 15C5 17 6 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="28;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaintDropFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M12 4C12 4 18 9 18 15C18 19 15 20 12 20C9 20 6 19 6 15C6 9 12 4 12 4Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.4s" dur="0.5s" values="0;1"/></path><path fill="none" stroke="currentColor" stroke-dasharray="28" stroke-dashoffset="28" stroke-linecap="round" stroke-width="2" d="M12 3C12 3 19 9 19 15C19 17 18 21 12 21M12 3C12 3 5 9 5 15C5 17 6 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="28;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaintDropHalfFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M6 15C6 9 12 4 12 4C12 4 14.9522 6.46019 16.715 10L6.8347 18C6.31173 17.2671 6 16.2894 6 15Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.4s" dur="0.5s" values="0;1"/></path><path fill="none" stroke="currentColor" stroke-dasharray="28" stroke-dashoffset="28" stroke-linecap="round" stroke-width="2" d="M12 3C12 3 19 9 19 15C19 17 18 21 12 21M12 3C12 3 5 9 5 15C5 17 6 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="28;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaintDropHalfFilledTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M6 15C6 9 12 4 12 4C12 4 14.9522 6.46019 16.715 10L6.8347 18C6.31173 17.2671 6 16.2894 6 15Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.4s" dur="0.5s" values="0;1"/></path><path fill="currentColor" fill-opacity="0" d="M12 4C12 4 18 9 18 15C18 19 15 20 12 20C9 20 6 19 6 15C6 9 12 4 12 4Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.4s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke="currentColor" stroke-dasharray="28" stroke-dashoffset="28" stroke-linecap="round" stroke-width="2" d="M12 3C12 3 19 9 19 15C19 17 18 21 12 21M12 3C12 3 5 9 5 15C5 17 6 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="28;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaintDropHalfTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M6 15C6 9 12 4 12 4C12 4 14.9522 6.46019 16.715 10L6.8347 18C6.31173 17.2671 6 16.2894 6 15Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.4s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke="currentColor" stroke-dasharray="28" stroke-dashoffset="28" stroke-linecap="round" stroke-width="2" d="M12 3C12 3 19 9 19 15C19 17 18 21 12 21M12 3C12 3 5 9 5 15C5 17 6 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="28;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaintDropTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M12 4C12 4 18 9 18 15C18 19 15 20 12 20C9 20 6 19 6 15C6 9 12 4 12 4Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.4s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke="currentColor" stroke-dasharray="28" stroke-dashoffset="28" stroke-linecap="round" stroke-width="2" d="M12 3C12 3 19 9 19 15C19 17 18 21 12 21M12 3C12 3 5 9 5 15C5 17 6 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="28;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Patreon(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" stroke="currentColor"><path stroke-width="3" d="M4 2.5v0"><animate fill="freeze" attributeName="d" dur="0.4s" values="M4 2.5v0;M4 2.5v19"/></path><path fill-opacity="0" stroke-dasharray="40" stroke-dashoffset="40" stroke-linecap="round" stroke-width="2" d="M14.88 3.5C18.26 3.5 21 6.24 21 9.63C21 13.01 18.26 15.75 14.88 15.75C11.49 15.75 8.75 13 8.75 9.63C8.75 6.24 11.49 3.5 14.88 3.5z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="40;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.5s" values="0;1"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="32" stroke-dashoffset="32" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 6h2v12h-2z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="32;0"/></path><path d="M15 6h2v12h-2z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="32;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseToPlayFilledTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 6L9 18L7 18L7 6z"><animate fill="freeze" attributeName="d" dur="0.4s" values="M9 18L7 18L7 6L9 6L9 18;M13 15L8 18L8 6L13 9L13 15"/><set attributeName="opacity" begin="0.4s" to="0"/></path><path d="M15 6L17 6L17 18L15 18L15 6"><animate fill="freeze" attributeName="d" dur="0.4s" values="M15 6L17 6L17 18L15 18L15 6;M13 9L18 12L18 12L13 15L13 9"/><set attributeName="opacity" begin="0.4s" to="0"/></path><path d="M8 6L18 12L8 18z" opacity="0"><set attributeName="opacity" begin="0.4s" to="1"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseToPlayTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 6L9 18L7 18L7 6z"><animate fill="freeze" attributeName="d" dur="0.6s" keyTimes="0;0.66;1" values="M9 18L7 18L7 6L9 6L9 18;M13 15L8 18L8 6L13 9L13 15;M13 15L8 18L8 6L13 9L13 9"/><set attributeName="opacity" begin="0.6s" to="0"/></path><path d="M15 6L17 6L17 18L15 18L15 6"><animate fill="freeze" attributeName="d" dur="0.6s" keyTimes="0;0.66;1" values="M15 6L17 6L17 18L15 18L15 6;M13 9L18 12L18 12L13 15L13 9;M13 9L18 12L18 12L13 15L13 15"/><set attributeName="opacity" begin="0.6s" to="0"/></path><path d="M8 6L18 12L8 18z" opacity="0"><set attributeName="opacity" begin="0.6s" to="1"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Peertube(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdPeertube0"><path fill="#fff" d="M4 1.34L20 12L4 22.67z"><animate fill="freeze" attributeName="d" dur="0.4s" values="M12 12L12 12L12 12z;M4 1.34L20 12L4 22.67z"/></path><path d="M12 6.67v10.67L4 12z"/></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdPeertube0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PeertubeAlt(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8 9.34L12 6.67L4 12z"><animate fill="freeze" attributeName="d" begin="0.2s" dur="0.4s" values="M8 9.34L12 6.67L4 12z;M4 1.34L12 6.67L4 12z"/></path><path d="M12 6.67L12 12L12 17.34z"><animate fill="freeze" attributeName="d" dur="0.4s" values="M12 6.67L12 12L12 17.34z;M12 6.67L20 12L12 17.34z"/></path><path d="M4 12L12 17.34L8 14.67z"><animate fill="freeze" attributeName="d" begin="0.4s" dur="0.4s" values="M4 12L12 17.34L8 14.67z;M4 12L12 17.34L4 22.67z"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pencil(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="56" stroke-dashoffset="56" d="M3 21L4.99998 15L16 4C17 3 19 3 20 4C21 5 21 7 20 8L8.99998 19L3 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="56;0"/></path><g stroke-dasharray="6" stroke-dashoffset="6"><path d="M15 5L19 9"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/></path><path stroke-width="1" d="M6 15L9 18"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/></path></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PencilTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="56" stroke-dashoffset="56" d="M3 21L4.99998 15L16 4C17 3 19 3 20 4C21 5 21 7 20 8L8.99998 19L3 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="56;0"/></path><g stroke-dasharray="6" stroke-dashoffset="6"><path d="M15 5L19 9"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/></path><path stroke-width="1" d="M6 15L9 18"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/></path></g></g><path fill="currentColor" fill-opacity="0" d="M17 4H20V7L9 18L6 15L17 4Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PencilTwotoneAlt(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="56" stroke-dashoffset="56" d="M3 21L4.99998 15L16 4C17 3 19 3 20 4C21 5 21 7 20 8L8.99998 19L3 21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="56;0"/></path><g stroke-dasharray="6" stroke-dashoffset="6"><path d="M15 5L19 9"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/></path><path stroke-width="1" d="M6 15L9 18"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/></path></g></g><path fill="currentColor" fill-opacity="0" d="M9 18L18 9L15 6L6 15L9 18Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Person(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M12 5C13.66 5 15 6.34 15 8C15 9.65685 13.6569 11 12 11C10.3431 11 9 9.65685 9 8C9 6.34315 10.3431 5 12 5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="36" stroke-dashoffset="36" d="M12 14C16 14 19 16 19 17V19H5V17C5 16 8 14 12 14z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="36;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonAdd(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M10 5C11.66 5 13 6.34 13 8C13 9.65685 11.6569 11 10 11C8.3431 11 7 9.65685 7 8C7 6.34315 8.3431 5 10 5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="36" stroke-dashoffset="36" d="M10 15C14 15 17 17 17 18V19H3V18C3 17 6 15 10 15z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="36;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M18 11h4" opacity="0"><set attributeName="opacity" begin="0.9s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="6;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M20 9v4" opacity="0"><set attributeName="opacity" begin="1.1s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonAddFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><g fill-opacity="0"><path stroke-dasharray="20" stroke-dashoffset="20" d="M10 5C11.66 5 13 6.34 13 8C13 9.65685 11.6569 11 10 11C8.3431 11 7 9.65685 7 8C7 6.34315 8.3431 5 10 5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="36" stroke-dashoffset="36" d="M10 15C14 15 17 17 17 18V19H3V18C3 17 6 15 10 15z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="36;0"/></path><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.5s" values="0;1"/></g><path stroke-dasharray="6" stroke-dashoffset="6" d="M18 11h4" opacity="0"><set attributeName="opacity" begin="1.5s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.5s" dur="0.2s" values="6;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M20 9v4" opacity="0"><set attributeName="opacity" begin="1.7s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.7s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonAddTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><g fill-opacity="0"><path stroke-dasharray="20" stroke-dashoffset="20" d="M10 5C11.66 5 13 6.34 13 8C13 9.65685 11.6569 11 10 11C8.3431 11 7 9.65685 7 8C7 6.34315 8.3431 5 10 5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="36" stroke-dashoffset="36" d="M10 15C14 15 17 17 17 18V19H3V18C3 17 6 15 10 15z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="36;0"/></path><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></g><path stroke-dasharray="6" stroke-dashoffset="6" d="M18 11h4" opacity="0"><set attributeName="opacity" begin="1.1s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="6;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M20 9v4" opacity="0"><set attributeName="opacity" begin="1.3s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.3s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M12 5C13.66 5 15 6.34 15 8C15 9.65685 13.6569 11 12 11C10.3431 11 9 9.65685 9 8C9 6.34315 10.3431 5 12 5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="36" stroke-dashoffset="36" d="M12 14C16 14 19 16 19 17V19H5V17C5 16 8 14 12 14z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="36;0"/></path><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.5s" values="0;1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonOff(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdPersonOff0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M12 5C13.66 5 15 6.34 15 8C15 9.65685 13.6569 11 12 11C10.3431 11 9 9.65685 9 8C9 6.34315 10.3431 5 12 5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="36" stroke-dashoffset="36" d="M12 14C16 14 19 16 19 17V19H5V17C5 16 8 14 12 14z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="36;0"/></path></g><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 12 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M1 13h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.0s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdPersonOff0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonOffFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdPersonOffFilled0"><g fill="#fff" fill-opacity="0" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M12 5C13.66 5 15 6.34 15 8C15 9.65685 13.6569 11 12 11C10.3431 11 9 9.65685 9 8C9 6.34315 10.3431 5 12 5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="36" stroke-dashoffset="36" d="M12 14C16 14 19 16 19 17V19H5V17C5 16 8 14 12 14z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="36;0"/></path><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.5s" values="0;1"/></g><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 12 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M1 13h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.4s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdPersonOffFilled0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonOffFilledLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdPersonOffFilledLoop0"><g fill="#fff" fill-opacity="0" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M12 5C13.66 5 15 6.34 15 8C15 9.65685 13.6569 11 12 11C10.3431 11 9 9.65685 9 8C9 6.34315 10.3431 5 12 5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="36" stroke-dashoffset="36" d="M12 14C16 14 19 16 19 17V19H5V17C5 16 8 14 12 14z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="36;0"/></path><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.5s" values="0;1"/></g><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 12 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M0 13h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 13h22;M2 13h22;M0 13h22"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.4s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdPersonOffFilledLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonOffLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdPersonOffLoop0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M12 5C13.66 5 15 6.34 15 8C15 9.65685 13.6569 11 12 11C10.3431 11 9 9.65685 9 8C9 6.34315 10.3431 5 12 5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="36" stroke-dashoffset="36" d="M12 14C16 14 19 16 19 17V19H5V17C5 16 8 14 12 14z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="36;0"/></path></g><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 12 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M0 13h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 13h22;M2 13h22;M0 13h22"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.0s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdPersonOffLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonOffTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdPersonOffTwotone0"><g fill="#fff" fill-opacity="0" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M12 5C13.66 5 15 6.34 15 8C15 9.65685 13.6569 11 12 11C10.3431 11 9 9.65685 9 8C9 6.34315 10.3431 5 12 5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="36" stroke-dashoffset="36" d="M12 14C16 14 19 16 19 17V19H5V17C5 16 8 14 12 14z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="36;0"/></path><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></g><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 12 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M1 13h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdPersonOffTwotone0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonOffTwotoneLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdPersonOffTwotoneLoop0"><g fill="#fff" fill-opacity="0" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M12 5C13.66 5 15 6.34 15 8C15 9.65685 13.6569 11 12 11C10.3431 11 9 9.65685 9 8C9 6.34315 10.3431 5 12 5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="36" stroke-dashoffset="36" d="M12 14C16 14 19 16 19 17V19H5V17C5 16 8 14 12 14z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="36;0"/></path><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></g><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 12 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M0 13h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 13h22;M2 13h22;M0 13h22"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdPersonOffTwotoneLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonRemove(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M10 5C11.66 5 13 6.34 13 8C13 9.65685 11.6569 11 10 11C8.3431 11 7 9.65685 7 8C7 6.34315 8.3431 5 10 5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="36" stroke-dashoffset="36" d="M10 15C14 15 17 17 17 18V19H3V18C3 17 6 15 10 15z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="36;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M18 11h4" opacity="0"><set attributeName="opacity" begin="0.9s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonRemoveFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><g fill-opacity="0"><path stroke-dasharray="20" stroke-dashoffset="20" d="M10 5C11.66 5 13 6.34 13 8C13 9.65685 11.6569 11 10 11C8.3431 11 7 9.65685 7 8C7 6.34315 8.3431 5 10 5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="36" stroke-dashoffset="36" d="M10 15C14 15 17 17 17 18V19H3V18C3 17 6 15 10 15z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="36;0"/></path><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.5s" values="0;1"/></g><path stroke-dasharray="6" stroke-dashoffset="6" d="M18 11h4" opacity="0"><set attributeName="opacity" begin="1.5s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.5s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonRemoveTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><g fill-opacity="0"><path stroke-dasharray="20" stroke-dashoffset="20" d="M10 5C11.66 5 13 6.34 13 8C13 9.65685 11.6569 11 10 11C8.3431 11 7 9.65685 7 8C7 6.34315 8.3431 5 10 5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="36" stroke-dashoffset="36" d="M10 15C14 15 17 17 17 18V19H3V18C3 17 6 15 10 15z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="36;0"/></path><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></g><path stroke-dasharray="6" stroke-dashoffset="6" d="M18 11h4" opacity="0"><set attributeName="opacity" begin="1.1s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonSearch(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdPersonSearch0"><g stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><g fill="none" stroke="#fff"><path stroke-dasharray="20" stroke-dashoffset="20" d="M10 5C11.66 5 13 6.34 13 8C13 9.65685 11.6569 11 10 11C8.3431 11 7 9.65685 7 8C7 6.34315 8.3431 5 10 5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="36" stroke-dashoffset="36" d="M10 15C14 15 17 17 17 18V19H3V18C3 17 6 15 10 15z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="36;0"/></path></g><circle cx="16" cy="16" r="6" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;1"/></circle><path fill="none" stroke="#fff" stroke-dasharray="26" stroke-dashoffset="26" d="M16 19C14.34 19 13 17.66 13 16C13 14.34 14.34 13 16 13C 17.66 13 19 14.34 19 16C19 17.66 17.66 19 16 19v4" transform="rotate(-45 16 16)"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.4s" values="26;0"/></path></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdPersonSearch0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonSearchFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdPersonSearchFilled0"><g fill-opacity="0" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><g fill="#fff" stroke="#fff"><path stroke-dasharray="20" stroke-dashoffset="20" d="M10 5C11.66 5 13 6.34 13 8C13 9.65685 11.6569 11 10 11C8.3431 11 7 9.65685 7 8C7 6.34315 8.3431 5 10 5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="36" stroke-dashoffset="36" d="M10 15C14 15 17 17 17 18V19H3V18C3 17 6 15 10 15z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="36;0"/></path><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.5s" values="0;1"/></g><circle cx="16" cy="16" r="6"><animate fill="freeze" attributeName="fill-opacity" begin="1.5s" dur="0.15s" values="0;1"/></circle><path fill="#fff" stroke="#fff" stroke-dasharray="30" stroke-dashoffset="30" d="M16 19C14.34 19 13 17.66 13 16C13 14.34 14.34 13 16 13C 17.66 13 19 14.34 19 16C19 17.66 17.66 19 16 19v4z" transform="rotate(-45 16 16)"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.5s" dur="0.4s" values="30;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.9s" dur="0.4s" values="0;1"/></path></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdPersonSearchFilled0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonSearchTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdPersonSearchTwotone0"><g fill-opacity="0" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><g fill="#fff" stroke="#fff"><path stroke-dasharray="20" stroke-dashoffset="20" d="M10 5C11.66 5 13 6.34 13 8C13 9.65685 11.6569 11 10 11C8.3431 11 7 9.65685 7 8C7 6.34315 8.3431 5 10 5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="36" stroke-dashoffset="36" d="M10 15C14 15 17 17 17 18V19H3V18C3 17 6 15 10 15z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="36;0"/></path><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></g><circle cx="16" cy="16" r="6"><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;1"/></circle><path fill="#fff" stroke="#fff" stroke-dasharray="30" stroke-dashoffset="30" d="M16 19C14.34 19 13 17.66 13 16C13 14.34 14.34 13 16 13C 17.66 13 19 14.34 19 16C19 17.66 17.66 19 16 19v4z" transform="rotate(-45 16 16)"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.4s" values="30;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.6s" dur="0.15s" values="0;0.3"/></path></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdPersonSearchTwotone0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M12 5C13.66 5 15 6.34 15 8C15 9.65685 13.6569 11 12 11C10.3431 11 9 9.65685 9 8C9 6.34315 10.3431 5 12 5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="36" stroke-dashoffset="36" d="M12 14C16 14 19 16 19 17V19H5V17C5 16 8 14 12 14z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="36;0"/></path><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="64" stroke-dashoffset="64" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 3C8.5 3 10.5 7.5 10.5 8C10.5 9 9 10 8.5 11C8 12 9 13 10 14C10.3943 14.3943 12 16 13 15.5C14 15 15 13.5 16 13.5C16.5 13.5 21 15.5 21 16C21 18 19.5 19.5 18 20C16.5 20.5 15.5 20.5 13.5 20C11.5 19.5 10 19 7.5 16.5C5 14 4.5 12.5 4 10.5C3.5 8.5 3.5 7.5 4 6C4.5 4.5 6 3 8 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneAdd(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-opacity="0" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M8 3C8.5 3 10.5 7.5 10.5 8C10.5 9 9 10 8.5 11C8 12 9 13 10 14C10.3943 14.3943 12 16 13 15.5C14 15 15 13.5 16 13.5C16.5 13.5 21 15.5 21 16C21 18 19.5 19.5 18 20C16.5 20.5 15.5 20.5 13.5 20C11.5 19.5 10 19 7.5 16.5C5 14 4.5 12.5 4 10.5C3.5 8.5 3.5 7.5 4 6C4.5 4.5 6 3 8 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M15 6h6" opacity="0"><set attributeName="opacity" begin="0.7s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M18 3v6" opacity="0"><set attributeName="opacity" begin="0.9s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneAddTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M8 3C8.5 3 10.5 7.5 10.5 8C10.5 9 9 10 8.5 11C8 12 9 13 10 14C10.3943 14.3943 12 16 13 15.5C14 15 15 13.5 16 13.5C16.5 13.5 21 15.5 21 16C21 18 19.5 19.5 18 20C16.5 20.5 15.5 20.5 13.5 20C11.5 19.5 10 19 7.5 16.5C5 14 4.5 12.5 4 10.5C3.5 8.5 3.5 7.5 4 6C4.5 4.5 6 3 8 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.15s" values="0;0.3"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M15 6h6" opacity="0"><set attributeName="opacity" begin="0.8s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M18 3v6" opacity="0"><set attributeName="opacity" begin="1s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneCall(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M8 3C8.5 3 10.5 7.5 10.5 8C10.5 9 9 10 8.5 11C8 12 9 13 10 14C10.3943 14.3943 12 16 13 15.5C14 15 15 13.5 16 13.5C16.5 13.5 21 15.5 21 16C21 18 19.5 19.5 18 20C16.5 20.5 15.5 20.5 13.5 20C11.5 19.5 10 19 7.5 16.5C5 14 4.5 12.5 4 10.5C3.5 8.5 3.5 7.5 4 6C4.5 4.5 6 3 8 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="4" stroke-dashoffset="4" d="M14 7.04404C14.6608 7.34734 15.2571 7.76718 15.7624 8.27723M16.956 10C16.6606 9.35636 16.2546 8.77401 15.7624 8.27723" opacity="0"><set attributeName="opacity" begin="0.7s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="4;8"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M20.748 9C20.3874 7.59926 19.6571 6.347 18.6672 5.3535M15 3.25203C16.4105 3.61507 17.6704 4.3531 18.6672 5.3535" opacity="0"><set attributeName="opacity" begin="1s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="10;20"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneCallLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M8 3C8.5 3 10.5 7.5 10.5 8C10.5 9 9 10 8.5 11C8 12 9 13 10 14C10.3943 14.3943 12 16 13 15.5C14 15 15 13.5 16 13.5C16.5 13.5 21 15.5 21 16C21 18 19.5 19.5 18 20C16.5 20.5 15.5 20.5 13.5 20C11.5 19.5 10 19 7.5 16.5C5 14 4.5 12.5 4 10.5C3.5 8.5 3.5 7.5 4 6C4.5 4.5 6 3 8 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/><animateTransform attributeName="transform" begin="0.6s;lineMdPhoneCallLoop0.begin+2.6s" dur="0.5s" type="rotate" values="0 12 12;15 12 12;0 12 12;-12 12 12;0 12 12;12 12 12;0 12 12;-15 12 12;0 12 12"/></path><path stroke-dasharray="4" stroke-dashoffset="4" d="M14 7.04404C14.6608 7.34734 15.2571 7.76718 15.7624 8.27723M16.956 10C16.6606 9.35636 16.2546 8.77401 15.7624 8.27723" opacity="0"><set id="lineMdPhoneCallLoop0" attributeName="opacity" begin="0.7s;lineMdPhoneCallLoop0.begin+2.7s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s;lineMdPhoneCallLoop0.begin+2.7s" dur="0.2s" values="4;8"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.3s;lineMdPhoneCallLoop0.begin+3.3s" dur="0.3s" values="0;4"/><set attributeName="opacity" begin="1.6s;lineMdPhoneCallLoop0.begin+3.6s" to="0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M20.748 9C20.3874 7.59926 19.6571 6.347 18.6672 5.3535M15 3.25203C16.4105 3.61507 17.6704 4.3531 18.6672 5.3535" opacity="0"><set attributeName="opacity" begin="1s;lineMdPhoneCallLoop0.begin+3s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s;lineMdPhoneCallLoop0.begin+3s" dur="0.2s" values="10;20"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.5s;lineMdPhoneCallLoop0.begin+3.5s" dur="0.3s" values="0;10"/><set attributeName="opacity" begin="1.8s;lineMdPhoneCallLoop0.begin+3.8s" to="0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneCallTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="64" stroke-dashoffset="64" d="M8 3C8.5 3 10.5 7.5 10.5 8C10.5 9 9 10 8.5 11C8 12 9 13 10 14C10.3943 14.3943 12 16 13 15.5C14 15 15 13.5 16 13.5C16.5 13.5 21 15.5 21 16C21 18 19.5 19.5 18 20C16.5 20.5 15.5 20.5 13.5 20C11.5 19.5 10 19 7.5 16.5C5 14 4.5 12.5 4 10.5C3.5 8.5 3.5 7.5 4 6C4.5 4.5 6 3 8 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.15s" values="0;0.3"/></path><path stroke-dasharray="4" stroke-dashoffset="4" d="M14 7.04404C14.6608 7.34734 15.2571 7.76718 15.7624 8.27723M16.956 10C16.6606 9.35636 16.2546 8.77401 15.7624 8.27723" opacity="0"><set attributeName="opacity" begin="0.7s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="4;8"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M20.748 9C20.3874 7.59926 19.6571 6.347 18.6672 5.3535M15 3.25203C16.4105 3.61507 17.6704 4.3531 18.6672 5.3535" opacity="0"><set attributeName="opacity" begin="1s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="10;20"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneCallTwotoneLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="64" stroke-dashoffset="64" d="M8 3C8.5 3 10.5 7.5 10.5 8C10.5 9 9 10 8.5 11C8 12 9 13 10 14C10.3943 14.3943 12 16 13 15.5C14 15 15 13.5 16 13.5C16.5 13.5 21 15.5 21 16C21 18 19.5 19.5 18 20C16.5 20.5 15.5 20.5 13.5 20C11.5 19.5 10 19 7.5 16.5C5 14 4.5 12.5 4 10.5C3.5 8.5 3.5 7.5 4 6C4.5 4.5 6 3 8 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.15s" values="0;0.3"/><animateTransform attributeName="transform" begin="0.6s;lineMdPhoneCallTwotoneLoop0.begin+2.6s" dur="0.5s" type="rotate" values="0 12 12;15 12 12;0 12 12;-12 12 12;0 12 12;12 12 12;0 12 12;-15 12 12;0 12 12"/></path><path stroke-dasharray="4" stroke-dashoffset="4" d="M14 7.04404C14.6608 7.34734 15.2571 7.76718 15.7624 8.27723M16.956 10C16.6606 9.35636 16.2546 8.77401 15.7624 8.27723" opacity="0"><set id="lineMdPhoneCallTwotoneLoop0" attributeName="opacity" begin="0.7s;lineMdPhoneCallTwotoneLoop0.begin+2.7s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s;lineMdPhoneCallTwotoneLoop0.begin+2.7s" dur="0.2s" values="4;8"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.3s;lineMdPhoneCallTwotoneLoop0.begin+3.3s" dur="0.3s" values="0;4"/><set attributeName="opacity" begin="1.6s;lineMdPhoneCallTwotoneLoop0.begin+3.6s" to="0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M20.748 9C20.3874 7.59926 19.6571 6.347 18.6672 5.3535M15 3.25203C16.4105 3.61507 17.6704 4.3531 18.6672 5.3535" opacity="0"><set attributeName="opacity" begin="1s;lineMdPhoneCallTwotoneLoop0.begin+3s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s;lineMdPhoneCallTwotoneLoop0.begin+3s" dur="0.2s" values="10;20"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.5s;lineMdPhoneCallTwotoneLoop0.begin+3.5s" dur="0.3s" values="0;10"/><set attributeName="opacity" begin="1.8s;lineMdPhoneCallTwotoneLoop0.begin+3.8s" to="0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneIncoming(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M8 3C8.5 3 10.5 7.5 10.5 8C10.5 9 9 10 8.5 11C8 12 9 13 10 14C10.3943 14.3943 12 16 13 15.5C14 15 15 13.5 16 13.5C16.5 13.5 21 15.5 21 16C21 18 19.5 19.5 18 20C16.5 20.5 15.5 20.5 13.5 20C11.5 19.5 10 19 7.5 16.5C5 14 4.5 12.5 4 10.5C3.5 8.5 3.5 7.5 4 6C4.5 4.5 6 3 8 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M20 4l-4 4" opacity="0"><set attributeName="opacity" begin="0.7s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M16 8h4M16 8v-4" opacity="0"><set attributeName="opacity" begin="0.9s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneIncomingTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M8 3C8.5 3 10.5 7.5 10.5 8C10.5 9 9 10 8.5 11C8 12 9 13 10 14C10.3943 14.3943 12 16 13 15.5C14 15 15 13.5 16 13.5C16.5 13.5 21 15.5 21 16C21 18 19.5 19.5 18 20C16.5 20.5 15.5 20.5 13.5 20C11.5 19.5 10 19 7.5 16.5C5 14 4.5 12.5 4 10.5C3.5 8.5 3.5 7.5 4 6C4.5 4.5 6 3 8 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.15s" values="0;0.3"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M20 4l-4 4" opacity="0"><set attributeName="opacity" begin="0.8s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M16 8h4M16 8v-4" opacity="0"><set attributeName="opacity" begin="1s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneOff(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdPhoneOff0"><path fill="none" fill-opacity="0" stroke="#fff" stroke-dasharray="64" stroke-dashoffset="64" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 3C8.5 3 10.5 7.5 10.5 8C10.5 9 9 10 8.5 11C8 12 9 13 10 14C10.3943 14.3943 12 16 13 15.5C14 15 15 13.5 16 13.5C16.5 13.5 21 15.5 21 16C21 18 19.5 19.5 18 20C16.5 20.5 15.5 20.5 13.5 20C11.5 19.5 10 19 7.5 16.5C5 14 4.5 12.5 4 10.5C3.5 8.5 3.5 7.5 4 6C4.5 4.5 6 3 8 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(135 12 12)"><path stroke="#000" d="M1 11h24"/><path stroke="#fff" d="M1 13h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin=".7" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdPhoneOff0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneOffLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdPhoneOffLoop0"><path fill="none" fill-opacity="0" stroke="#fff" stroke-dasharray="64" stroke-dashoffset="64" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 3C8.5 3 10.5 7.5 10.5 8C10.5 9 9 10 8.5 11C8 12 9 13 10 14C10.3943 14.3943 12 16 13 15.5C14 15 15 13.5 16 13.5C16.5 13.5 21 15.5 21 16C21 18 19.5 19.5 18 20C16.5 20.5 15.5 20.5 13.5 20C11.5 19.5 10 19 7.5 16.5C5 14 4.5 12.5 4 10.5C3.5 8.5 3.5 7.5 4 6C4.5 4.5 6 3 8 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><g fill="none" stroke-linecap="round" stroke-width="2"><path stroke="#000" d="M1 11h24" transform="rotate(135 12 12)"/><path stroke="#fff" d="M1 13h22" transform="rotate(135 12 12)"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M1 13h22;M3 13h22;M1 13h22"/></path></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdPhoneOffLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneOffTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdPhoneOffTwotone0"><g fill-opacity="0" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="#fff" stroke-dasharray="64" stroke-dashoffset="64" d="M8 3C8.5 3 10.5 7.5 10.5 8C10.5 9 9 10 8.5 11C8 12 9 13 10 14C10.3943 14.3943 12 16 13 15.5C14 15 15 13.5 16 13.5C16.5 13.5 21 15.5 21 16C21 18 19.5 19.5 18 20C16.5 20.5 15.5 20.5 13.5 20C11.5 19.5 10 19 7.5 16.5C5 14 4.5 12.5 4 10.5C3.5 8.5 3.5 7.5 4 6C4.5 4.5 6 3 8 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.15s" values="0;0.3"/></path></g><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(135 12 12)"><path stroke="#000" d="M1 11h24"/><path stroke="#fff" d="M1 13h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin=".8" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdPhoneOffTwotone0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneOffTwotoneLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdPhoneOffTwotoneLoop0"><g fill-opacity="0" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="#fff" stroke-dasharray="64" stroke-dashoffset="64" d="M8 3C8.5 3 10.5 7.5 10.5 8C10.5 9 9 10 8.5 11C8 12 9 13 10 14C10.3943 14.3943 12 16 13 15.5C14 15 15 13.5 16 13.5C16.5 13.5 21 15.5 21 16C21 18 19.5 19.5 18 20C16.5 20.5 15.5 20.5 13.5 20C11.5 19.5 10 19 7.5 16.5C5 14 4.5 12.5 4 10.5C3.5 8.5 3.5 7.5 4 6C4.5 4.5 6 3 8 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.15s" values="0;0.3"/></path></g><g fill="none" stroke-linecap="round" stroke-width="2"><path stroke="#000" d="M1 11h24" transform="rotate(135 12 12)"/><path stroke="#fff" d="M1 13h22" transform="rotate(135 12 12)"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M1 13h22;M3 13h22;M1 13h22"/></path></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdPhoneOffTwotoneLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneOutgoing(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M8 3C8.5 3 10.5 7.5 10.5 8C10.5 9 9 10 8.5 11C8 12 9 13 10 14C10.3943 14.3943 12 16 13 15.5C14 15 15 13.5 16 13.5C16.5 13.5 21 15.5 21 16C21 18 19.5 19.5 18 20C16.5 20.5 15.5 20.5 13.5 20C11.5 19.5 10 19 7.5 16.5C5 14 4.5 12.5 4 10.5C3.5 8.5 3.5 7.5 4 6C4.5 4.5 6 3 8 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M16 8l4 -4" opacity="0"><set attributeName="opacity" begin="0.7s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M20 4h-4M20 4v4" opacity="0"><set attributeName="opacity" begin="0.9s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneOutgoingTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M8 3C8.5 3 10.5 7.5 10.5 8C10.5 9 9 10 8.5 11C8 12 9 13 10 14C10.3943 14.3943 12 16 13 15.5C14 15 15 13.5 16 13.5C16.5 13.5 21 15.5 21 16C21 18 19.5 19.5 18 20C16.5 20.5 15.5 20.5 13.5 20C11.5 19.5 10 19 7.5 16.5C5 14 4.5 12.5 4 10.5C3.5 8.5 3.5 7.5 4 6C4.5 4.5 6 3 8 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.15s" values="0;0.3"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M16 8l4 -4" opacity="0"><set attributeName="opacity" begin="0.8s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M20 4h-4M20 4v4" opacity="0"><set attributeName="opacity" begin="1s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneRemove(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-opacity="0" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M8 3C8.5 3 10.5 7.5 10.5 8C10.5 9 9 10 8.5 11C8 12 9 13 10 14C10.3943 14.3943 12 16 13 15.5C14 15 15 13.5 16 13.5C16.5 13.5 21 15.5 21 16C21 18 19.5 19.5 18 20C16.5 20.5 15.5 20.5 13.5 20C11.5 19.5 10 19 7.5 16.5C5 14 4.5 12.5 4 10.5C3.5 8.5 3.5 7.5 4 6C4.5 4.5 6 3 8 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M16 4l4 4" opacity="0"><set attributeName="opacity" begin="0.7s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M20 4l-4 4" opacity="0"><set attributeName="opacity" begin="0.9s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneRemoveTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M8 3C8.5 3 10.5 7.5 10.5 8C10.5 9 9 10 8.5 11C8 12 9 13 10 14C10.3943 14.3943 12 16 13 15.5C14 15 15 13.5 16 13.5C16.5 13.5 21 15.5 21 16C21 18 19.5 19.5 18 20C16.5 20.5 15.5 20.5 13.5 20C11.5 19.5 10 19 7.5 16.5C5 14 4.5 12.5 4 10.5C3.5 8.5 3.5 7.5 4 6C4.5 4.5 6 3 8 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.15s" values="0;0.3"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M16 4l4 4" opacity="0"><set attributeName="opacity" begin="0.8s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M20 4l-4 4" opacity="0"><set attributeName="opacity" begin="1s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="8;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill-opacity="0" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" stroke-dasharray="64" stroke-dashoffset="64" d="M8 3C8.5 3 10.5 7.5 10.5 8C10.5 9 9 10 8.5 11C8 12 9 13 10 14C10.3943 14.3943 12 16 13 15.5C14 15 15 13.5 16 13.5C16.5 13.5 21 15.5 21 16C21 18 19.5 19.5 18 20C16.5 20.5 15.5 20.5 13.5 20C11.5 19.5 10 19 7.5 16.5C5 14 4.5 12.5 4 10.5C3.5 8.5 3.5 7.5 4 6C4.5 4.5 6 3 8 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.15s" values="0;0.3"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pixelfed(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path fill="currentColor" fill-opacity="0" d="M11.2061 14.1832H13.0406C14.7687 14.1832 16.1697 12.8194 16.1697 11.1371C16.1697 9.45472 14.7687 8.09091 13.0406 8.09091H10.3929C9.39592 8.09091 8.58769 8.87772 8.58769 9.8483V16.6883L11.2061 14.1832Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.5s" values="0;1"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PixelfedFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdPixelfedFilled0"><path fill="#fff" fill-opacity="0" stroke="#fff" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.5s" values="0;1"/></path><path d="M11.2061 14.1832H13.0406C14.7687 14.1832 16.1697 12.8194 16.1697 11.1371C16.1697 9.45472 14.7687 8.09091 13.0406 8.09091H10.3929C9.39592 8.09091 8.58769 8.87772 8.58769 9.8483V16.6883L11.2061 14.1832Z"/></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdPixelfedFilled0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PixelfedTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.15s" values="0;0.3"/></path><path fill="currentColor" fill-opacity="0" d="M11.2061 14.1832H13.0406C14.7687 14.1832 16.1697 12.8194 16.1697 11.1371C16.1697 9.45472 14.7687 8.09091 13.0406 8.09091H10.3929C9.39592 8.09091 8.58769 8.87772 8.58769 9.8483V16.6883L11.2061 14.1832Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.5s" values="0;1"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="36" stroke-dashoffset="36" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 6L18 12L8 18z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="36;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="36" stroke-dashoffset="36" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 6L18 12L8 18z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="36;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.5s" values="0;1"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayFilledToPauseTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13 15L8 18L8 6L13 9L13 15"><animate fill="freeze" attributeName="d" dur="0.4s" values="M13 15L8 18L8 6L13 9L13 15;M9 18L7 18L7 6L9 6L9 18"/></path><path d="M13 9L18 12L18 12L13 15L13 9"><animate fill="freeze" attributeName="d" dur="0.4s" values="M13 9L18 12L18 12L13 15L13 9;M15 6L17 6L17 18L15 18L15 6"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayToPauseTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13 15L8 18L8 6L13 9L13 9"><animate fill="freeze" attributeName="d" dur="0.6s" keyTimes="0;0.33;1" values="M13 15L8 18L8 6L13 9L13 9;M13 15L8 18L8 6L13 9L13 15;M9 18L7 18L7 6L9 6L9 18"/></path><path d="M13 9L18 12L18 12L13 15L13 15"><animate fill="freeze" attributeName="d" dur="0.6s" keyTimes="0;0.33;1" values="M13 9L18 12L18 12L13 15L13 15;M13 9L18 12L18 12L13 15L13 9;M15 6L17 6L17 18L15 18L15 6"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="36" stroke-dashoffset="36" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 6L18 12L8 18z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="36;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.15s" values="0;0.3"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pleroma(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M6 3.5C6 2.75 6.75 2 7.5 2H11v20h-5z"><animate fill="freeze" attributeName="d" dur="0.4s" values="M6 2C6 2 6.75 2 7.5 2H11v0h-5z;M6 3.5C6 2.75 6.75 2 7.5 2H11v20h-5z"/></path><path d="M13.5 2h0v8.5c0 0.75 0 1.5 0 1.5h0z"><animate fill="freeze" attributeName="d" begin="0.5s" dur="0.2s" values="M13.5 2h0v8.5c0 0.75 0 1.5 0 1.5h0z;M13.5 2h5v8.5c0 0.75 -0.75 1.5 -1.5 1.5h-3.5z"/></path><path d="M13.5 17h5v0c0 0 -0.75 0 -1.5 0h-3.5z"><animate fill="freeze" attributeName="d" begin="0.8s" dur="0.2s" values="M13.5 17h0v0c0 0 0 0 0 0h0z;M13.5 17h5v3.5c0 0.75 -0.75 1.5 -1.5 1.5h-3.5z"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="18" stroke-dashoffset="18" stroke-linecap="round" stroke-width="2"><path d="M12 5V19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.3s" values="18;0"/></path><path d="M5 12H19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="18;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircle(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><g stroke-dasharray="12" stroke-dashoffset="12"><path d="M12 7V17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="12;0"/></path><path d="M7 12H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="12;0"/></path></g><path stroke-dasharray="60" stroke-dashoffset="60" d="M3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircleTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><g stroke-dasharray="12" stroke-dashoffset="12"><path d="M12 7V17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="12;0"/></path><path d="M7 12H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="12;0"/></path></g><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.15s" values="0;0.3"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSquare(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><g stroke-dasharray="12" stroke-dashoffset="12"><path d="M12 7V17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="12;0"/></path><path d="M7 12H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="12;0"/></path></g><path stroke-dasharray="64" stroke-dashoffset="64" d="M4 12V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="64;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSquareTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><g stroke-dasharray="12" stroke-dashoffset="12"><path d="M12 7V17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="12;0"/></path><path d="M7 12H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="12;0"/></path></g><path fill="currentColor" fill-opacity="0" stroke-dasharray="64" stroke-dashoffset="64" d="M4 12V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="64;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.15s" values="0;0.3"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Question(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="32" stroke-dashoffset="32" stroke-linecap="round" stroke-width="2" d="M7 8C7 5.23858 9.23857 3 12 3C14.7614 3 17 5.23858 17 8C17 9.6356 16.2147 11.0878 15.0005 12C14.1647 12.6279 12 14 12 17"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="32;0"/></path><circle cx="12" cy="21" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.2s" values="0;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionCircle(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="20" stroke-dashoffset="20" d="M8.99999 10C8.99999 8.34315 10.3431 7 12 7C13.6569 7 15 8.34315 15 10C15 10.9814 14.5288 11.8527 13.8003 12.4C13.0718 12.9473 12.5 13 12 14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.4s" values="20;0"/></path></g><circle cx="12" cy="17" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.2s" values="0;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionCircleTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="20" stroke-dashoffset="20" d="M8.99999 10C8.99999 8.34315 10.3431 7 12 7C13.6569 7 15 8.34315 15 10C15 10.9814 14.5288 11.8527 13.8003 12.4C13.0718 12.9473 12.5 13 12 14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.4s" values="20;0"/></path></g><circle cx="12" cy="17" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.2s" values="0;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionSquare(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M12 4H5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="64;128"/></path><path stroke-dasharray="20" stroke-dashoffset="20" d="M8.99999 10C8.99999 8.34315 10.3431 7 12 7C13.6569 7 15 8.34315 15 10C15 10.9814 14.5288 11.8527 13.8003 12.4C13.0718 12.9473 12.5 13 12 14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.4s" values="20;0"/></path></g><circle cx="12" cy="17" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.2s" values="0;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionSquareTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="64" stroke-dashoffset="64" d="M12 4H5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="64;128"/><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="20" stroke-dashoffset="20" d="M8.99999 10C8.99999 8.34315 10.3431 7 12 7C13.6569 7 15 8.34315 15 10C15 10.9814 14.5288 11.8527 13.8003 12.4C13.0718 12.9473 12.5 13 12 14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.4s" values="20;0"/></path></g><circle cx="12" cy="17" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.2s" values="0;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reddit(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdReddit0"><g fill="#fff" fill-opacity="0"><ellipse cx="12" cy="14.71" stroke="#fff" stroke-dasharray="48" stroke-dashoffset="48" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="8" ry="5.29"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="48;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.4s" values="0;1"/></ellipse><circle cx="7.24" cy="11.97" r="2.24"><set attributeName="fill-opacity" begin="1s" to="1"/><animate fill="freeze" attributeName="cx" begin="1s" dur="0.2s" values="7.24;3.94"/></circle><circle cx="16.76" cy="11.97" r="2.24"><set attributeName="fill-opacity" begin="1s" to="1"/><animate fill="freeze" attributeName="cx" begin="1s" dur="0.2s" values="16.76;20.06"/></circle><circle cx="18.45" cy="4.23" r="1.61"><set attributeName="fill-opacity" begin="2.6s" to="1"/></circle></g><path fill="none" stroke="#fff" stroke-dasharray="12" stroke-dashoffset="12" stroke-linecap="round" stroke-linejoin="round" stroke-width=".8" d="M12 8.75L13.18 3.11L18.21 4.18"><animate fill="freeze" attributeName="stroke-dashoffset" begin="2.4s" dur="0.2s" values="12;0"/></path><g fill-opacity="0"><circle cx="8.45" cy="13.59" r="1.61"><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.4s" values="0;1"/></circle><circle cx="15.55" cy="13.59" r="1.61"><animate fill="freeze" attributeName="fill-opacity" begin="1.6s" dur="0.4s" values="0;1"/></circle></g><path fill="none" stroke="#000" stroke-dasharray="8" stroke-dashoffset="8" stroke-linecap="round" stroke-width=".8" d="M8.47 17.52C8.47 17.52 9.41 18.58 12 18.58C14.58 18.58 15.53 17.52 15.53 17.52"><animate fill="freeze" attributeName="stroke-dashoffset" begin="2s" dur="0.2s" values="8;0"/></path></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdReddit0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedditCircle(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdRedditCircle0"><path fill="#fff" fill-opacity="0" stroke="#fff" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.5s" values="0;1"/></path><g fill-opacity="0"><ellipse cx="12" cy="13.77" rx="5.83" ry="4.06"><animate fill="freeze" attributeName="fill-opacity" begin="1.1s" dur="0.4s" values="0;1"/></ellipse><circle cx="8.99" cy="11.99" r="1.45"><set attributeName="fill-opacity" begin="1.5s" to="1"/><animate fill="freeze" attributeName="cx" begin="1.5s" dur="0.2s" values="8.99;6.79"/></circle><circle cx="15.01" cy="11.99" r="1.45"><set attributeName="fill-opacity" begin="1.5s" to="1"/><animate fill="freeze" attributeName="cx" begin="1.5s" dur="0.2s" values="15.01;17.21"/></circle><circle cx="16.18" cy="7.01" r="1.04"><set attributeName="fill-opacity" begin="3.1s" to="1"/></circle></g><path fill="none" stroke="#000" stroke-dasharray="8" stroke-dashoffset="8" stroke-linecap="round" stroke-linejoin="round" stroke-width=".54" d="M12 9.91L12.76 6.27L16 6.98"><animate fill="freeze" attributeName="stroke-dashoffset" begin="2.9s" dur="0.2s" values="8;0"/></path><g fill="#fff" fill-opacity="0"><circle cx="9.71" cy="13.04" r="1.04"><animate fill="freeze" attributeName="fill-opacity" begin="1.7s" dur="0.4s" values="0;1"/></circle><circle cx="14.29" cy="13.04" r="1.04"><animate fill="freeze" attributeName="fill-opacity" begin="2.1s" dur="0.4s" values="0;1"/></circle></g><path fill="none" stroke="#fff" stroke-dasharray="6" stroke-dashoffset="6" stroke-linecap="round" stroke-width=".54" d="M9.72 15.6C9.72 15.6 10.33 16.29 12 16.291C13.67 16.29 14.28 15.6 14.28 15.6"><animate fill="freeze" attributeName="stroke-dashoffset" begin="2.5s" dur="0.2s" values="6;0"/></path></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdRedditCircle0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedditCircleLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdRedditCircleLoop0"><path fill="#fff" fill-opacity="0" stroke="#fff" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.5s" values="0;1"/></path><g fill-opacity="0"><ellipse cx="12" cy="13.77" rx="5.83" ry="4.06"><animate fill="freeze" attributeName="fill-opacity" begin="1.1s" dur="0.4s" values="0;1"/></ellipse><circle cx="8.99" cy="11.99" r="1.45"><set attributeName="fill-opacity" begin="1.5s" to="1"/><animate fill="freeze" attributeName="cx" begin="1.5s" dur="0.2s" values="8.99;6.79"/></circle><circle cx="15.01" cy="11.99" r="1.45"><set attributeName="fill-opacity" begin="1.5s" to="1"/><animate fill="freeze" attributeName="cx" begin="1.5s" dur="0.2s" values="15.01;17.21"/></circle><circle cx="16.18" cy="7.01" r="1.04"><set attributeName="fill-opacity" begin="3.1s" to="1"/><animate attributeName="cx" begin="2.9s" dur="6s" repeatCount="indefinite" values="16.18;7.82;16.18"/></circle></g><path fill="none" stroke="#000" stroke-dasharray="8" stroke-dashoffset="8" stroke-linecap="round" stroke-linejoin="round" stroke-width=".54" d="M12 9.91L12.76 6.27L16 6.98"><animate fill="freeze" attributeName="stroke-dashoffset" begin="2.9s" dur="0.2s" values="8;0"/><animate attributeName="d" begin="2.9s" dur="6s" repeatCount="indefinite" values="M12 9.91L12.76 6.27L16 6.98;M12 9.91L12 5.2L12 6.98;M12 9.91L11.24 6.27L8 6.98;M12 9.91L12 5.2L12 6.98;M12 9.91L12.76 6.27L16 6.98"/></path><g fill="#fff" fill-opacity="0"><circle cx="9.71" cy="13.04" r="1.04"><animate fill="freeze" attributeName="fill-opacity" begin="1.7s" dur="0.4s" values="0;1"/></circle><circle cx="14.29" cy="13.04" r="1.04"><animate fill="freeze" attributeName="fill-opacity" begin="2.1s" dur="0.4s" values="0;1"/></circle></g><path fill="none" stroke="#fff" stroke-dasharray="6" stroke-dashoffset="6" stroke-linecap="round" stroke-width=".54" d="M9.72 15.6C9.72 15.6 10.33 16.29 12 16.291C13.67 16.29 14.28 15.6 14.28 15.6"><animate fill="freeze" attributeName="stroke-dashoffset" begin="2.5s" dur="0.2s" values="6;0"/></path></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdRedditCircleLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedditLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdRedditLoop0"><g fill="#fff" fill-opacity="0"><ellipse cx="12" cy="14.71" stroke="#fff" stroke-dasharray="48" stroke-dashoffset="48" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="8" ry="5.29"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="48;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.4s" values="0;1"/></ellipse><circle cx="7.24" cy="11.97" r="2.24"><set attributeName="fill-opacity" begin="1s" to="1"/><animate fill="freeze" attributeName="cx" begin="1s" dur="0.2s" values="7.24;3.94"/></circle><circle cx="16.76" cy="11.97" r="2.24"><set attributeName="fill-opacity" begin="1s" to="1"/><animate fill="freeze" attributeName="cx" begin="1s" dur="0.2s" values="16.76;20.06"/></circle><circle cx="18.45" cy="4.23" r="1.61"><set attributeName="fill-opacity" begin="2.6s" to="1"/><animate attributeName="cx" begin="2.4s" dur="6s" repeatCount="indefinite" values="18.45;5.75;18.45"/></circle></g><path fill="none" stroke="#fff" stroke-dasharray="12" stroke-dashoffset="12" stroke-linecap="round" stroke-linejoin="round" stroke-width=".8" d="M12 8.75L13.18 3.11L18.21 4.18"><animate fill="freeze" attributeName="stroke-dashoffset" begin="2.4s" dur="0.2s" values="12;0"/><animate attributeName="d" begin="2.4s" dur="6s" repeatCount="indefinite" values="M12 8.75L13.18 3.11L18.21 4.18;M12 8.75L12 2L12 4.18;M12 8.75L10.82 3.11L5.79 4.18;M12 8.75L12 2L12 4.18;M12 8.75L13.18 3.11L18.21 4.18"/></path><g fill-opacity="0"><circle cx="8.45" cy="13.59" r="1.61"><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.4s" values="0;1"/></circle><circle cx="15.55" cy="13.59" r="1.61"><animate fill="freeze" attributeName="fill-opacity" begin="1.6s" dur="0.4s" values="0;1"/></circle></g><path fill="none" stroke="#000" stroke-dasharray="8" stroke-dashoffset="8" stroke-linecap="round" stroke-width=".8" d="M8.47 17.52C8.47 17.52 9.41 18.58 12 18.58C14.58 18.58 15.53 17.52 15.53 17.52"><animate fill="freeze" attributeName="stroke-dashoffset" begin="2s" dur="0.2s" values="8;0"/></path></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdRedditLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Remove(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-dasharray="22" stroke-dashoffset="22" stroke-linecap="round" stroke-width="2"><path d="M19 5L5 19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.3s" values="22;0"/></path><path d="M5 5L19 19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="22;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RotateNinety(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M12 6C15.3137 6 18 8.68629 18 12V14.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="14;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M18 15L21 12M18 15L15 12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RotateOneHundredEighty(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="24" stroke-dashoffset="24" d="M12 6C15.3137 6 18 8.68629 18 12C18 15.3137 15.3137 18 12 18H9.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="24;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M9 18L12 21M9 18L12 15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RotateTwoHundredSeventy(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="32" stroke-dashoffset="32" d="M12 6C15.3137 6 18 8.68629 18 12C18 15.3137 15.3137 18 12 18C8.68629 18 6 15.3137 6 12V9.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="32;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M6 9L3 12M6 9L9 12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoundRampLeft(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="18" stroke-dashoffset="18" d="M17.5 17.5284C14.2121 15.6264 12 12.0716 12 8V4M12 20V4"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="18;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 4l-3 3M12 4l3 3" opacity="0"><set attributeName="opacity" begin="0.4s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoundThreeHundredSixty(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="40" stroke-dashoffset="40" stroke-linecap="round" stroke-width="2" d="M17 15.3264C19.412 14.6089 21 13.3869 21 12C21 9.79086 16.9706 8 12 8C7.02944 8 3 9.79086 3 12C3 14.0589 6.50005 15.7545 11 15.9756"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="40;0"/></path><path fill="currentColor" fill-opacity="0" d="M12.3414 16.6586L12.8586 16.1414C12.9367 16.0633 12.9367 15.9367 12.8586 15.8586L12.3414 15.3414C12.2154 15.2154 12 15.3047 12 15.4828V16.5172C12 16.6953 12.2154 16.7846 12.3414 16.6586Z"><set attributeName="fill-opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="d" begin="0.5s" dur="0.2s" values="M12.3414 16.6586L12.8586 16.1414C12.9367 16.0633 12.9367 15.9367 12.8586 15.8586L12.3414 15.3414C12.2154 15.2154 12 15.3047 12 15.4828V16.5172C12 16.6953 12.2154 16.7846 12.3414 16.6586Z;M9.85355 19.1464L12.6464 16.3536C12.8417 16.1583 12.8417 15.8417 12.6464 15.6464L9.85355 12.8536C9.53857 12.5386 9 12.7617 9 13.2071V18.7929C9 19.2383 9.53857 19.4614 9.85355 19.1464Z"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoundaboutLeft(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="40" stroke-dashoffset="40" d="M16 20V15C16 14.4477 16.4522 14.0104 16.9936 13.9013C19.279 13.4405 21 11.4212 21 9C21 6.23858 18.7614 4 16 4C13.5788 4 11.5595 5.72096 11.0987 8.00638C10.9896 8.54777 10.5523 9 10 9H3"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="40;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M3 9l3 -3M3 9l3 3" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rss(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="16" stroke-dashoffset="16" d="M4 11C6.39 11 8.68 11.95 10.36 13.64C12.05 15.32 13 17.61 13 20" opacity="0"><set attributeName="opacity" begin="0.3s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.3s" values="16;0"/></path><path stroke-dasharray="28" stroke-dashoffset="28" d="M4 4C8.24 4 12.31 5.69 15.31 8.69C18.31 11.69 20 15.76 20 20" opacity="0"><set attributeName="opacity" begin="0.7s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.4s" values="28;0"/></path></g><circle cx="5" cy="19" r="2" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" dur="0.2s" values="0;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="16" stroke-dashoffset="16" d="M10.5 13.5L3 21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="16;0"/></path><path stroke-dasharray="40" stroke-dashoffset="40" d="M10.7574 13.2426C8.41421 10.8995 8.41421 7.10051 10.7574 4.75736C13.1005 2.41421 16.8995 2.41421 19.2426 4.75736C21.5858 7.10051 21.5858 10.8995 19.2426 13.2426C16.8995 15.5858 13.1005 15.5858 10.7574 13.2426Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="40;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="none" stroke-dasharray="16" stroke-dashoffset="16" d="M10.5 13.5L3 21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="16;0"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="40" stroke-dashoffset="40" d="M10.7574 13.2426C8.41421 10.8995 8.41421 7.10051 10.7574 4.75736C13.1005 2.41421 16.8995 2.41421 19.2426 4.75736C21.5858 7.10051 21.5858 10.8995 19.2426 13.2426C16.8995 15.5858 13.1005 15.5858 10.7574 13.2426Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="40;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.5s" values="0;1"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="none" stroke-dasharray="16" stroke-dashoffset="16" d="M10.5 13.5L3 21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="16;0"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="40" stroke-dashoffset="40" d="M10.7574 13.2426C8.41421 10.8995 8.41421 7.10051 10.7574 4.75736C13.1005 2.41421 16.8995 2.41421 19.2426 4.75736C21.5858 7.10051 21.5858 10.8995 19.2426 13.2426C16.8995 15.5858 13.1005 15.5858 10.7574 13.2426Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="40;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.15s" values="0;0.3"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Speed(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdSpeed0"><path fill="none" stroke="#fff" stroke-dasharray="56" stroke-dashoffset="56" stroke-linecap="round" stroke-width="2" d="M5 19V19C4.69726 19 4.41165 18.8506 4.25702 18.5904C3.45852 17.2464 3 15.6767 3 14C3 9.02944 7.02944 5 12 5C16.9706 5 21 9.02944 21 14C21 15.6767 20.5415 17.2464 19.743 18.5904C19.5884 18.8506 19.3027 19 19 19z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="56;0"/></path><g fill-opacity="0" transform="rotate(-100 12 14)"><path d="M12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14Z"><animate fill="freeze" attributeName="d" begin="0.4s" dur="0.2s" values="M12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14Z;M16 14C16 16.21 14.21 18 12 18C9.79 18 8 16.21 8 14C8 11.79 12 0 12 0C12 0 16 11.79 16 14Z"/></path><path fill="#fff" d="M12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14Z"><animate fill="freeze" attributeName="d" begin="0.4s" dur="0.2s" values="M12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14Z;M14 14C14 15.1 13.1 16 12 16C10.9 16 10 15.1 10 14C10 12.9 12 4 12 4C12 4 14 12.9 14 14Z"/></path><set attributeName="fill-opacity" begin="0.4s" to="1"/><animateTransform fill="freeze" attributeName="transform" begin="0.6s" dur="0.3s" type="rotate" values="-100 12 14;45 12 14"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdSpeed0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpeedLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdSpeedLoop0"><path fill="none" stroke="#fff" stroke-dasharray="56" stroke-dashoffset="56" stroke-linecap="round" stroke-width="2" d="M5 19V19C4.69726 19 4.41165 18.8506 4.25702 18.5904C3.45852 17.2464 3 15.6767 3 14C3 9.02944 7.02944 5 12 5C16.9706 5 21 9.02944 21 14C21 15.6767 20.5415 17.2464 19.743 18.5904C19.5884 18.8506 19.3027 19 19 19z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="56;0"/></path><g fill-opacity="0" transform="rotate(-100 12 14)"><path d="M12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14Z"><animate fill="freeze" attributeName="d" begin="0.4s" dur="0.2s" values="M12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14Z;M16 14C16 16.21 14.21 18 12 18C9.79 18 8 16.21 8 14C8 11.79 12 0 12 0C12 0 16 11.79 16 14Z"/></path><path fill="#fff" d="M12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14Z"><animate fill="freeze" attributeName="d" begin="0.4s" dur="0.2s" values="M12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14Z;M14 14C14 15.1 13.1 16 12 16C10.9 16 10 15.1 10 14C10 12.9 12 4 12 4C12 4 14 12.9 14 14Z"/></path><set attributeName="fill-opacity" begin="0.4s" to="1"/><animateTransform attributeName="transform" begin="0.6s" dur="6s" repeatCount="indefinite" type="rotate" values="-100 12 14;45 12 14;45 12 14;45 12 14;20 12 14;10 12 14;0 12 14;35 12 14;45 12 14;55 12 14;50 12 14;15 12 14;-20 12 14;-100 12 14"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdSpeedLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Speedometer(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdSpeedometer0"><path fill="none" stroke="#fff" stroke-dasharray="46" stroke-dashoffset="46" stroke-linecap="round" stroke-width="2" d="M21 13C21 8.02944 16.9706 4 12 4C7.02944 4 3 8.02944 3 13C3 17.9706 7.02944 22 12 22" transform="rotate(45 12 13)"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="46;92"/></path><g fill-opacity="0" transform="rotate(-100 12 13)"><path d="M12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14Z"><animate fill="freeze" attributeName="d" begin="0.4s" dur="0.2s" values="M12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14Z;M17 13C17 15.7614 14.7614 18 12 18C9.23858 18 7 15.7614 7 13C7 10.2386 12 -2 12 -2C12 -2 17 10.2386 17 13Z"/></path><path fill="#fff" d="M12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14Z"><animate fill="freeze" attributeName="d" begin="0.4s" dur="0.2s" values="M12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14Z;M15 13C15 14.6568 13.6569 16 12 16C10.3431 16 9 14.6568 9 13C9 11.3431 12 2 12 2C12 2 15 11.3431 15 13Z"/></path><set attributeName="fill-opacity" begin="0.4s" to="1"/><animateTransform fill="freeze" attributeName="transform" begin="0.6s" dur="0.3s" type="rotate" values="-100 12 13;65 12 13"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdSpeedometer0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpeedometerLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdSpeedometerLoop0"><path fill="none" stroke="#fff" stroke-dasharray="46" stroke-dashoffset="46" stroke-linecap="round" stroke-width="2" d="M21 13C21 8.02944 16.9706 4 12 4C7.02944 4 3 8.02944 3 13C3 17.9706 7.02944 22 12 22" transform="rotate(45 12 13)"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="46;92"/></path><g fill-opacity="0" transform="rotate(-100 12 13)"><path d="M12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14Z"><animate fill="freeze" attributeName="d" begin="0.4s" dur="0.2s" values="M12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14Z;M17 13C17 15.7614 14.7614 18 12 18C9.23858 18 7 15.7614 7 13C7 10.2386 12 -2 12 -2C12 -2 17 10.2386 17 13Z"/></path><path fill="#fff" d="M12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14Z"><animate fill="freeze" attributeName="d" begin="0.4s" dur="0.2s" values="M12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14C12 14 12 14 12 14Z;M15 13C15 14.6568 13.6569 16 12 16C10.3431 16 9 14.6568 9 13C9 11.3431 12 2 12 2C12 2 15 11.3431 15 13Z"/></path><set attributeName="fill-opacity" begin="0.4s" to="1"/><animateTransform attributeName="transform" begin="0.6s" dur="6s" repeatCount="indefinite" type="rotate" values="-100 12 13;65 12 13;65 12 13;65 12 13;30 12 13;10 12 13;0 12 13;35 12 13;55 12 13;65 12 13;75 12 13;15 12 13;-20 12 13;-100 12 13"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdSpeedometerLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Square(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="64" stroke-dashoffset="64" stroke-linecap="round" stroke-width="2" d="M4 12V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="64;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareToConfirmSquareTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 19V5a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1Z"/><path stroke-dasharray="14" stroke-dashoffset="14" d="M8 12L11 15L16 10"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="14;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareToConfirmSquareTwotoneTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" d="M20 19V5a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.2s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke-dasharray="14" stroke-dashoffset="14" d="M8 12L11 15L16 10"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="14;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="64" stroke-dashoffset="64" stroke-linecap="round" stroke-width="2" d="M4 12V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="64;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.15s" values="0;0.3"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareTwotoneToConfirmSquareTwotoneTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity=".3" d="M20 19V5a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1Z"/><path fill="none" stroke-dasharray="14" stroke-dashoffset="14" d="M8 12L11 15L16 10"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="14;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="32" stroke-dashoffset="32" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3L9.65 8.76L3.44 9.22L8.2 13.24L6.71 19.28L12 16M12 3L14.35 8.76L20.56 9.22L15.8 13.24L17.29 19.28L12 16"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="32;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarAlt(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="64" stroke-dashoffset="64" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3L14.35 8.76L20.56 9.22L15.8 13.24L17.29 19.28L12 16L6.71 19.28L8.2 13.24L3.44 9.22L9.65 8.76Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarAltFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="64" stroke-dashoffset="64" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3L14.35 8.76L20.56 9.22L15.8 13.24L17.29 19.28L12 16L6.71 19.28L8.2 13.24L3.44 9.22L9.65 8.76Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.5s" values="0;1"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarAltTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="64" stroke-dashoffset="64" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3L14.35 8.76L20.56 9.22L15.8 13.24L17.29 19.28L12 16L6.71 19.28L8.2 13.24L3.44 9.22L9.65 8.76Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.15s" values="0;0.3"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="32" stroke-dashoffset="32" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3L9.65 8.76L3.44 9.22L8.2 13.24L6.71 19.28L12 16M12 3L14.35 8.76L20.56 9.22L15.8 13.24L17.29 19.28L12 16"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="32;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.5s" values="0;1"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarFilledHalf(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="32" stroke-dashoffset="32" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3L9.65 8.76L3.44 9.22L8.2 13.24L6.71 19.28L12 16M12 3L14.35 8.76L20.56 9.22L15.8 13.24L17.29 19.28L12 16"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="32;0"/></path><path fill="currentColor" fill-opacity="0" d="M12 3L9.65 8.76L3.44 9.22L8.2 13.24L6.71 19.28L12 16Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.5s" values="0;1"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarHalf(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="48" stroke-dashoffset="48" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3L9.65 8.76L3.44 9.22L8.2 13.24L6.71 19.28L12 16Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="48;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarHalfFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="48" stroke-dashoffset="48" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3L9.65 8.76L3.44 9.22L8.2 13.24L6.71 19.28L12 16Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="48;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.5s" values="0;1"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarHalfTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="48" stroke-dashoffset="48" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3L9.65 8.76L3.44 9.22L8.2 13.24L6.71 19.28L12 16Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="48;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.15s" values="0;0.3"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarPulsatingFilledLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="32" stroke-dashoffset="32" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3L9.65 8.76L3.44 9.22L8.2 13.24L6.71 19.28L12 16M12 3L14.35 8.76L20.56 9.22L15.8 13.24L17.29 19.28L12 16"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="32;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.5s" values="0;1"/><animate attributeName="d" dur="1.5s" repeatCount="indefinite" values="M12 3L9.65 8.76L3.44 9.22L8.2 13.24L6.71 19.28L12 16M12 3L14.35 8.76L20.56 9.22L15.8 13.24L17.29 19.28L12 16;M12 7L10.82 10.38L7.24 10.45L10.1 12.62L9.06 16.05L12 14M12 7L13.18 10.38L16.76 10.45L13.9 12.62L14.94 16.05L12 14;M12 3L9.65 8.76L3.44 9.22L8.2 13.24L6.71 19.28L12 16M12 3L14.35 8.76L20.56 9.22L15.8 13.24L17.29 19.28L12 16"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarPulsatingLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="32" stroke-dashoffset="32" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3L9.65 8.76L3.44 9.22L8.2 13.24L6.71 19.28L12 16M12 3L14.35 8.76L20.56 9.22L15.8 13.24L17.29 19.28L12 16"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="32;0"/><animate attributeName="d" dur="1.5s" repeatCount="indefinite" values="M12 3L9.65 8.76L3.44 9.22L8.2 13.24L6.71 19.28L12 16M12 3L14.35 8.76L20.56 9.22L15.8 13.24L17.29 19.28L12 16;M12 7L10.82 10.38L7.24 10.45L10.1 12.62L9.06 16.05L12 14M12 7L13.18 10.38L16.76 10.45L13.9 12.62L14.94 16.05L12 14;M12 3L9.65 8.76L3.44 9.22L8.2 13.24L6.71 19.28L12 16M12 3L14.35 8.76L20.56 9.22L15.8 13.24L17.29 19.28L12 16"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarPulsatingTwotoneLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="32" stroke-dashoffset="32" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3L9.65 8.76L3.44 9.22L8.2 13.24L6.71 19.28L12 16M12 3L14.35 8.76L20.56 9.22L15.8 13.24L17.29 19.28L12 16"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="32;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.15s" values="0;0.3"/><animate attributeName="d" dur="1.5s" repeatCount="indefinite" values="M12 3L9.65 8.76L3.44 9.22L8.2 13.24L6.71 19.28L12 16M12 3L14.35 8.76L20.56 9.22L15.8 13.24L17.29 19.28L12 16;M12 7L10.82 10.38L7.24 10.45L10.1 12.62L9.06 16.05L12 14M12 7L13.18 10.38L16.76 10.45L13.9 12.62L14.94 16.05L12 14;M12 3L9.65 8.76L3.44 9.22L8.2 13.24L6.71 19.28L12 16M12 3L14.35 8.76L20.56 9.22L15.8 13.24L17.29 19.28L12 16"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="32" stroke-dashoffset="32" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3L9.65 8.76L3.44 9.22L8.2 13.24L6.71 19.28L12 16M12 3L14.35 8.76L20.56 9.22L15.8 13.24L17.29 19.28L12 16"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="32;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.15s" values="0;0.3"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarTwotoneHalf(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="32" stroke-dashoffset="32" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3L9.65 8.76L3.44 9.22L8.2 13.24L6.71 19.28L12 16M12 3L14.35 8.76L20.56 9.22L15.8 13.24L17.29 19.28L12 16"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="32;0"/></path><path fill="currentColor" fill-opacity="0" d="M12 3L9.65 8.76L3.44 9.22L8.2 13.24L6.71 19.28L12 16Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.15s" values="0;0.3"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunRisingFilledLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="12" cy="32" r="6" fill="currentColor"><animate fill="freeze" attributeName="cy" dur="0.6s" values="32;12"/></circle><g fill="none" stroke="currentColor" stroke-dasharray="2" stroke-dashoffset="2" stroke-linecap="round" stroke-width="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.7s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.9s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="2;0"/></path><animateTransform attributeName="transform" dur="30s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunRisingLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><circle cx="12" cy="32" r="5"><animate fill="freeze" attributeName="cy" dur="0.6s" values="32;12"/></circle><g stroke-dasharray="2" stroke-dashoffset="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.7s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.9s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="2;0"/></path><animateTransform attributeName="transform" dur="30s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunRisingTwotoneLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><circle cx="12" cy="32" r="5" fill="currentColor" fill-opacity=".3"><animate fill="freeze" attributeName="cy" dur="0.6s" values="32;12"/></circle><g fill="none" stroke-dasharray="2" stroke-dashoffset="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.7s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.9s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="2;0"/></path><animateTransform attributeName="transform" dur="30s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunnyFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="34" stroke-dashoffset="34" d="M12 7C14.76 7 17 9.24 17 12C17 14.76 14.76 17 12 17C9.24 17 7 14.76 7 12C7 9.24 9.24 7 12 7"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="34;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.5s" values="0;1"/></path><g fill="none" stroke-dasharray="2" stroke-dashoffset="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.5s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.7s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="2;0"/></path></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunnyFilledLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="34" stroke-dashoffset="34" d="M12 7C14.76 7 17 9.24 17 12C17 14.76 14.76 17 12 17C9.24 17 7 14.76 7 12C7 9.24 9.24 7 12 7"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="34;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.5s" values="0;1"/></path><g fill="none" stroke-dasharray="2" stroke-dashoffset="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.5s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.7s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="2;0"/></path><animateTransform attributeName="transform" dur="30s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunnyFilledLoopToMoonAltFilledLoopTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><g stroke-dasharray="2"><path d="M12 21v1M21 12h1M12 3v-1M3 12h-1"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="4;2"/></path><path d="M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="4;2"/></path></g><path fill="currentColor" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/></path></g><g fill="none" stroke="currentColor" stroke-dasharray="4" stroke-dashoffset="4" stroke-linecap="round" stroke-linejoin="round"><path d="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"><animate id="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition0" fill="freeze" attributeName="stroke-dashoffset" begin="0.6s;lineMdSunnyFilledLoopToMoonAltFilledLoopTransition0.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition0.begin+2s;lineMdSunnyFilledLoopToMoonAltFilledLoopTransition0.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition0.begin+1.2s;lineMdSunnyFilledLoopToMoonAltFilledLoopTransition0.begin+3.2s;lineMdSunnyFilledLoopToMoonAltFilledLoopTransition0.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition0.begin+1.8s" to="M12 5h1.5M12 5h-1.5M12 5v1.5M12 5v-1.5"/><set attributeName="d" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition0.begin+3.8s" to="M12 4h1.5M12 4h-1.5M12 4v1.5M12 4v-1.5"/><set attributeName="d" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition0.begin+5.8s" to="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"/></path><path d="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"><animate id="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition1" fill="freeze" attributeName="stroke-dashoffset" begin="1s;lineMdSunnyFilledLoopToMoonAltFilledLoopTransition1.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition1.begin+2s;lineMdSunnyFilledLoopToMoonAltFilledLoopTransition1.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition1.begin+1.2s;lineMdSunnyFilledLoopToMoonAltFilledLoopTransition1.begin+3.2s;lineMdSunnyFilledLoopToMoonAltFilledLoopTransition1.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition1.begin+1.8s" to="M17 11h1.5M17 11h-1.5M17 11v1.5M17 11v-1.5"/><set attributeName="d" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition1.begin+3.8s" to="M18 12h1.5M18 12h-1.5M18 12v1.5M18 12v-1.5"/><set attributeName="d" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition1.begin+5.8s" to="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"/></path><path d="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"><animate id="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition2" fill="freeze" attributeName="stroke-dashoffset" begin="2.8s;lineMdSunnyFilledLoopToMoonAltFilledLoopTransition2.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition2.begin+2s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition2.begin+1.2s;lineMdSunnyFilledLoopToMoonAltFilledLoopTransition2.begin+3.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition2.begin+1.8s" to="M20 5h1.5M20 5h-1.5M20 5v1.5M20 5v-1.5"/><set attributeName="d" begin="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition2.begin+5.8s" to="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"/></path></g><mask id="lineMdSunnyFilledLoopToMoonAltFilledLoopTransition3"><circle cx="12" cy="12" r="12" fill="#fff"/><circle cx="22" cy="2" r="3" fill="#fff"><animate fill="freeze" attributeName="cx" begin="0.1s" dur="0.4s" values="22;18"/><animate fill="freeze" attributeName="cy" begin="0.1s" dur="0.4s" values="2;6"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="3;12"/></circle><circle cx="22" cy="2" r="1"><animate fill="freeze" attributeName="cx" begin="0.1s" dur="0.4s" values="22;18"/><animate fill="freeze" attributeName="cy" begin="0.1s" dur="0.4s" values="2;6"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="1;10"/></circle></mask><circle cx="12" cy="12" r="6" fill="currentColor" mask="url(#lineMdSunnyFilledLoopToMoonAltFilledLoopTransition3)"><set attributeName="opacity" begin="0.5s" to="0"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="6;10"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunnyFilledLoopToMoonFilledLoopTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><g stroke-dasharray="2"><path d="M12 21v1M21 12h1M12 3v-1M3 12h-1"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="4;2"/></path><path d="M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="4;2"/></path></g><path fill="currentColor" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/></path></g><g fill="currentColor" fill-opacity="0"><path d="m15.22 6.03l2.53-1.94L14.56 4L13.5 1l-1.06 3l-3.19.09l2.53 1.94l-.91 3.06l2.63-1.81l2.63 1.81z"><animate id="lineMdSunnyFilledLoopToMoonFilledLoopTransition0" fill="freeze" attributeName="fill-opacity" begin="0.6s;lineMdSunnyFilledLoopToMoonFilledLoopTransition0.begin+6s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyFilledLoopToMoonFilledLoopTransition0.begin+2.2s" dur="0.4s" values="1;0"/></path><path d="M13.61 5.25L15.25 4l-2.06-.05L12.5 2l-.69 1.95L9.75 4l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyFilledLoopToMoonFilledLoopTransition0.begin+3s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyFilledLoopToMoonFilledLoopTransition0.begin+5.2s" dur="0.4s" values="1;0"/></path><path d="M19.61 12.25L21.25 11l-2.06-.05L18.5 9l-.69 1.95l-2.06.05l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyFilledLoopToMoonFilledLoopTransition0.begin+0.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyFilledLoopToMoonFilledLoopTransition0.begin+2.8s" dur="0.4s" values="1;0"/></path><path d="m20.828 9.731l1.876-1.439l-2.366-.067L19.552 6l-.786 2.225l-2.366.067l1.876 1.439L17.601 12l1.951-1.342L21.503 12z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyFilledLoopToMoonFilledLoopTransition0.begin+3.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyFilledLoopToMoonFilledLoopTransition0.begin+5.6s" dur="0.4s" values="1;0"/></path></g><mask id="lineMdSunnyFilledLoopToMoonFilledLoopTransition1"><circle cx="12" cy="12" r="12" fill="#fff"/><circle cx="22" cy="2" r="3" fill="#fff"><animate fill="freeze" attributeName="cx" begin="0.1s" dur="0.4s" values="22;18"/><animate fill="freeze" attributeName="cy" begin="0.1s" dur="0.4s" values="2;6"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="3;12"/></circle><circle cx="22" cy="2" r="1"><animate fill="freeze" attributeName="cx" begin="0.1s" dur="0.4s" values="22;18"/><animate fill="freeze" attributeName="cy" begin="0.1s" dur="0.4s" values="2;6"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="1;10"/></circle></mask><circle cx="12" cy="12" r="6" fill="currentColor" mask="url(#lineMdSunnyFilledLoopToMoonFilledLoopTransition1)"><set attributeName="opacity" begin="0.5s" to="0"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="6;10"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunnyFilledLoopToMoonFilledTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><g stroke-dasharray="2"><path d="M12 21v1M21 12h1M12 3v-1M3 12h-1"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="4;2"/></path><path d="M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="4;2"/></path></g><path fill="currentColor" d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/></path></g><g fill="currentColor" fill-opacity="0"><path d="M15.22 6.03L17.75 4.09L14.56 4L13.5 1L12.44 4L9.25 4.09L11.78 6.03L10.87 9.09L13.5 7.28L16.13 9.09L15.22 6.03Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.4s" values="0;1"/></path><path d="M19.61 12.25L21.25 11L19.19 10.95L18.5 9L17.81 10.95L15.75 11L17.39 12.25L16.8 14.23L18.5 13.06L20.2 14.23L19.61 12.25Z"><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.4s" values="0;1"/></path></g><mask id="lineMdSunnyFilledLoopToMoonFilledTransition0"><circle cx="12" cy="12" r="12" fill="#fff"/><circle cx="22" cy="2" r="3" fill="#fff"><animate fill="freeze" attributeName="cx" begin="0.1s" dur="0.4s" values="22;18"/><animate fill="freeze" attributeName="cy" begin="0.1s" dur="0.4s" values="2;6"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="3;12"/></circle><circle cx="22" cy="2" r="1"><animate fill="freeze" attributeName="cx" begin="0.1s" dur="0.4s" values="22;18"/><animate fill="freeze" attributeName="cy" begin="0.1s" dur="0.4s" values="2;6"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="1;10"/></circle></mask><circle cx="12" cy="12" r="6" fill="currentColor" mask="url(#lineMdSunnyFilledLoopToMoonFilledTransition0)"><set attributeName="opacity" begin="0.5s" to="0"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="6;10"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunnyOutline(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="34" stroke-dashoffset="34" d="M12 7C14.76 7 17 9.24 17 12C17 14.76 14.76 17 12 17C9.24 17 7 14.76 7 12C7 9.24 9.24 7 12 7"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="34;0"/></path><g stroke-dasharray="2" stroke-dashoffset="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.5s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.7s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="2;0"/></path></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunnyOutlineLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="34" stroke-dashoffset="34" d="M12 7C14.76 7 17 9.24 17 12C17 14.76 14.76 17 12 17C9.24 17 7 14.76 7 12C7 9.24 9.24 7 12 7"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="34;0"/></path><g stroke-dasharray="2" stroke-dashoffset="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.5s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.7s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="2;0"/></path><animateTransform attributeName="transform" dur="30s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunnyOutlineToMoonAltLoopTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><g stroke-dasharray="2"><path d="M12 21v1M21 12h1M12 3v-1M3 12h-1"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="4;2"/></path><path d="M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="4;2"/></path></g><path d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/></path></g><g fill="none" stroke="currentColor" stroke-dasharray="4" stroke-dashoffset="4" stroke-linecap="round" stroke-linejoin="round"><path d="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"><animate id="lineMdSunnyOutlineToMoonAltLoopTransition0" fill="freeze" attributeName="stroke-dashoffset" begin="0.6s;lineMdSunnyOutlineToMoonAltLoopTransition0.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdSunnyOutlineToMoonAltLoopTransition0.begin+2s;lineMdSunnyOutlineToMoonAltLoopTransition0.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdSunnyOutlineToMoonAltLoopTransition0.begin+1.2s;lineMdSunnyOutlineToMoonAltLoopTransition0.begin+3.2s;lineMdSunnyOutlineToMoonAltLoopTransition0.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdSunnyOutlineToMoonAltLoopTransition0.begin+1.8s" to="M12 5h1.5M12 5h-1.5M12 5v1.5M12 5v-1.5"/><set attributeName="d" begin="lineMdSunnyOutlineToMoonAltLoopTransition0.begin+3.8s" to="M12 4h1.5M12 4h-1.5M12 4v1.5M12 4v-1.5"/><set attributeName="d" begin="lineMdSunnyOutlineToMoonAltLoopTransition0.begin+5.8s" to="M13 4h1.5M13 4h-1.5M13 4v1.5M13 4v-1.5"/></path><path d="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"><animate id="lineMdSunnyOutlineToMoonAltLoopTransition1" fill="freeze" attributeName="stroke-dashoffset" begin="1s;lineMdSunnyOutlineToMoonAltLoopTransition1.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdSunnyOutlineToMoonAltLoopTransition1.begin+2s;lineMdSunnyOutlineToMoonAltLoopTransition1.begin+4s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdSunnyOutlineToMoonAltLoopTransition1.begin+1.2s;lineMdSunnyOutlineToMoonAltLoopTransition1.begin+3.2s;lineMdSunnyOutlineToMoonAltLoopTransition1.begin+5.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdSunnyOutlineToMoonAltLoopTransition1.begin+1.8s" to="M17 11h1.5M17 11h-1.5M17 11v1.5M17 11v-1.5"/><set attributeName="d" begin="lineMdSunnyOutlineToMoonAltLoopTransition1.begin+3.8s" to="M18 12h1.5M18 12h-1.5M18 12v1.5M18 12v-1.5"/><set attributeName="d" begin="lineMdSunnyOutlineToMoonAltLoopTransition1.begin+5.8s" to="M19 11h1.5M19 11h-1.5M19 11v1.5M19 11v-1.5"/></path><path d="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"><animate id="lineMdSunnyOutlineToMoonAltLoopTransition2" fill="freeze" attributeName="stroke-dashoffset" begin="2.8s;lineMdSunnyOutlineToMoonAltLoopTransition2.begin+6s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdSunnyOutlineToMoonAltLoopTransition2.begin+2s" dur="0.4s" values="4;0"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="lineMdSunnyOutlineToMoonAltLoopTransition2.begin+1.2s;lineMdSunnyOutlineToMoonAltLoopTransition2.begin+3.2s" dur="0.4s" values="0;4"/><set attributeName="d" begin="lineMdSunnyOutlineToMoonAltLoopTransition2.begin+1.8s" to="M20 5h1.5M20 5h-1.5M20 5v1.5M20 5v-1.5"/><set attributeName="d" begin="lineMdSunnyOutlineToMoonAltLoopTransition2.begin+5.8s" to="M19 4h1.5M19 4h-1.5M19 4v1.5M19 4v-1.5"/></path></g><mask id="lineMdSunnyOutlineToMoonAltLoopTransition3"><circle cx="12" cy="12" r="12" fill="#fff"/><circle cx="12" cy="12" r="4"><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="4;8"/></circle><circle cx="22" cy="2" r="3" fill="#fff"><animate fill="freeze" attributeName="cx" begin="0.1s" dur="0.4s" values="22;18"/><animate fill="freeze" attributeName="cy" begin="0.1s" dur="0.4s" values="2;6"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="3;12"/></circle><circle cx="22" cy="2" r="1"><animate fill="freeze" attributeName="cx" begin="0.1s" dur="0.4s" values="22;18"/><animate fill="freeze" attributeName="cy" begin="0.1s" dur="0.4s" values="2;6"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="1;10"/></circle></mask><circle cx="12" cy="12" r="6" fill="currentColor" mask="url(#lineMdSunnyOutlineToMoonAltLoopTransition3)"><set attributeName="opacity" begin="0.5s" to="0"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="6;10"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunnyOutlineToMoonLoopTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><g stroke-dasharray="2"><path d="M12 21v1M21 12h1M12 3v-1M3 12h-1"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="4;2"/></path><path d="M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="4;2"/></path></g><path d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/></path></g><g fill="currentColor" fill-opacity="0"><path d="m15.22 6.03l2.53-1.94L14.56 4L13.5 1l-1.06 3l-3.19.09l2.53 1.94l-.91 3.06l2.63-1.81l2.63 1.81z"><animate id="lineMdSunnyOutlineToMoonLoopTransition0" fill="freeze" attributeName="fill-opacity" begin="0.6s;lineMdSunnyOutlineToMoonLoopTransition0.begin+6s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyOutlineToMoonLoopTransition0.begin+2.2s" dur="0.4s" values="1;0"/></path><path d="M13.61 5.25L15.25 4l-2.06-.05L12.5 2l-.69 1.95L9.75 4l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyOutlineToMoonLoopTransition0.begin+3s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyOutlineToMoonLoopTransition0.begin+5.2s" dur="0.4s" values="1;0"/></path><path d="M19.61 12.25L21.25 11l-2.06-.05L18.5 9l-.69 1.95l-2.06.05l1.64 1.25l-.59 1.98l1.7-1.17l1.7 1.17z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyOutlineToMoonLoopTransition0.begin+0.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyOutlineToMoonLoopTransition0.begin+2.8s" dur="0.4s" values="1;0"/></path><path d="m20.828 9.731l1.876-1.439l-2.366-.067L19.552 6l-.786 2.225l-2.366.067l1.876 1.439L17.601 12l1.951-1.342L21.503 12z"><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyOutlineToMoonLoopTransition0.begin+3.4s" dur="0.4s" values="0;1"/><animate fill="freeze" attributeName="fill-opacity" begin="lineMdSunnyOutlineToMoonLoopTransition0.begin+5.6s" dur="0.4s" values="1;0"/></path></g><mask id="lineMdSunnyOutlineToMoonLoopTransition1"><circle cx="12" cy="12" r="12" fill="#fff"/><circle cx="12" cy="12" r="4"><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="4;8"/></circle><circle cx="22" cy="2" r="3" fill="#fff"><animate fill="freeze" attributeName="cx" begin="0.1s" dur="0.4s" values="22;18"/><animate fill="freeze" attributeName="cy" begin="0.1s" dur="0.4s" values="2;6"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="3;12"/></circle><circle cx="22" cy="2" r="1"><animate fill="freeze" attributeName="cx" begin="0.1s" dur="0.4s" values="22;18"/><animate fill="freeze" attributeName="cy" begin="0.1s" dur="0.4s" values="2;6"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="1;10"/></circle></mask><circle cx="12" cy="12" r="6" fill="currentColor" mask="url(#lineMdSunnyOutlineToMoonLoopTransition1)"><set attributeName="opacity" begin="0.5s" to="0"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="6;10"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunnyOutlineToMoonTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><g stroke-dasharray="2"><path d="M12 21v1M21 12h1M12 3v-1M3 12h-1"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="4;2"/></path><path d="M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="4;2"/></path></g><path d="M7 6 C7 12.08 11.92 17 18 17 C18.53 17 19.05 16.96 19.56 16.89 C17.95 19.36 15.17 21 12 21 C7.03 21 3 16.97 3 12 C3 8.83 4.64 6.05 7.11 4.44 C7.04 4.95 7 5.47 7 6 Z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/></path></g><g fill="currentColor" fill-opacity="0"><path d="M15.22 6.03L17.75 4.09L14.56 4L13.5 1L12.44 4L9.25 4.09L11.78 6.03L10.87 9.09L13.5 7.28L16.13 9.09L15.22 6.03Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.4s" values="0;1"/></path><path d="M19.61 12.25L21.25 11L19.19 10.95L18.5 9L17.81 10.95L15.75 11L17.39 12.25L16.8 14.23L18.5 13.06L20.2 14.23L19.61 12.25Z"><animate fill="freeze" attributeName="fill-opacity" begin="1s" dur="0.4s" values="0;1"/></path></g><mask id="lineMdSunnyOutlineToMoonTransition0"><circle cx="12" cy="12" r="12" fill="#fff"/><circle cx="12" cy="12" r="4"><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="4;8"/></circle><circle cx="22" cy="2" r="3" fill="#fff"><animate fill="freeze" attributeName="cx" begin="0.1s" dur="0.4s" values="22;18"/><animate fill="freeze" attributeName="cy" begin="0.1s" dur="0.4s" values="2;6"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="3;12"/></circle><circle cx="22" cy="2" r="1"><animate fill="freeze" attributeName="cx" begin="0.1s" dur="0.4s" values="22;18"/><animate fill="freeze" attributeName="cy" begin="0.1s" dur="0.4s" values="2;6"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="1;10"/></circle></mask><circle cx="12" cy="12" r="6" fill="currentColor" mask="url(#lineMdSunnyOutlineToMoonTransition0)"><set attributeName="opacity" begin="0.5s" to="0"/><animate fill="freeze" attributeName="r" begin="0.1s" dur="0.4s" values="6;10"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunnyOutlineTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="34" stroke-dashoffset="34" d="M12 7C14.76 7 17 9.24 17 12C17 14.76 14.76 17 12 17C9.24 17 7 14.76 7 12C7 9.24 9.24 7 12 7"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="34;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke-dasharray="2" stroke-dashoffset="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.5s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.7s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="2;0"/></path></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunnyOutlineTwotoneLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="34" stroke-dashoffset="34" d="M12 7C14.76 7 17 9.24 17 12C17 14.76 14.76 17 12 17C9.24 17 7 14.76 7 12C7 9.24 9.24 7 12 7"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="34;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.15s" values="0;0.3"/></path><g fill="none" stroke-dasharray="2" stroke-dashoffset="2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.5s" dur="0.2s" values="M12 19v1M19 12h1M12 5v-1M5 12h-1;M12 21v1M21 12h1M12 3v-1M3 12h-1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="2;0"/></path><path d="M0 0"><animate fill="freeze" attributeName="d" begin="0.7s" dur="0.2s" values="M17 17l0.5 0.5M17 7l0.5 -0.5M7 7l-0.5 -0.5M7 17l-0.5 0.5;M18.5 18.5l0.5 0.5M18.5 5.5l0.5 -0.5M5.5 5.5l-0.5 -0.5M5.5 18.5l-0.5 0.5"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="2;0"/></path><animateTransform attributeName="transform" dur="30s" repeatCount="indefinite" type="rotate" values="0 12 12;360 12 12"/></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Switch(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<rect width="20" height="10" x="2" y="7" fill="none" stroke="currentColor" stroke-dasharray="54" stroke-dashoffset="54" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="54;0"/></rect><circle cx="17" cy="12" r="3" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.2s" values="0;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwitchFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdSwitchFilled0"><g fill="#fff" fill-opacity="0" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="8" x="3" y="8" stroke-dasharray="54" stroke-dashoffset="54" rx="4"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="54;0"/></rect><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.2s" values="0;1"/></g><circle cx="17" cy="12" r="3" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.2s" values="0;1"/></circle></mask><rect width="20" height="10" x="2" y="7" fill="currentColor" mask="url(#lineMdSwitchFilled0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwitchFilledToSwitchOffFilledTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 7a5 5 0 0 1 0 10H7A5 5 0 1 1 7 7h10Zm-10 2a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z"><animate fill="freeze" attributeName="d" dur="0.2s" values="M17 7a5 5 0 0 1 0 10H7A5 5 0 1 1 7 7h10Zm-10 2a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z;M17 7a5 5 0 0 1 0 10H7A5 5 0 1 1 7 7h10Zm0 2a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwitchOff(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<rect width="20" height="10" x="2" y="7" fill="none" stroke="currentColor" stroke-dasharray="54" stroke-dashoffset="54" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="54;0"/></rect><circle cx="7" cy="12" r="3" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.2s" values="0;1"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwitchOffFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdSwitchOffFilled0"><g fill="#fff" fill-opacity="0" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="8" x="3" y="8" stroke-dasharray="54" stroke-dashoffset="54" rx="4"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="54;0"/></rect><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.2s" values="0;1"/></g><circle cx="7" cy="12" r="3" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.9s" dur="0.2s" values="0;1"/></circle></mask><rect width="20" height="10" x="2" y="7" fill="currentColor" mask="url(#lineMdSwitchOffFilled0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwitchOffFilledToSwitchFilledTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 7a5 5 0 0 1 0 10H7A5 5 0 1 1 7 7h10Zm0 2a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z"><animate fill="freeze" attributeName="d" dur="0.2s" values="M17 7a5 5 0 0 1 0 10H7A5 5 0 1 1 7 7h10Zm0 2a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z;M17 7a5 5 0 0 1 0 10H7A5 5 0 1 1 7 7h10Zm-10 2a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwitchOffToSwitchTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<rect width="20" height="10" x="2" y="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="5"/><circle cx="7" cy="12" r="3" fill="currentColor"><animate fill="freeze" attributeName="cx" dur="0.2s" values="7;17"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwitchToSwitchOffTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<rect width="20" height="10" x="2" y="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="5"/><circle cx="17" cy="12" r="3" fill="currentColor"><animate fill="freeze" attributeName="cx" dur="0.2s" values="17;7"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Telegram(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="16" stroke-dashoffset="16" d="M21 5L18.5 20M21 5L9 13.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="16;0"/></path><path stroke-dasharray="22" stroke-dashoffset="22" d="M21 5L2 12.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="22;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M18.5 20L9 13.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.3s" values="12;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M2 12.5L9 13.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.3s" values="8;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 16L9 19M9 13.5L9 19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.3s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextBox(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="62" stroke-dashoffset="62" d="M20 6V5C20 4.45 19.55 4 19 4H5C4.45 4 4 4.45 4 5V19C4 19.55 4.45 20 5 20H19C19.55 20 20 19.55 20 19z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="62;124"/></path><g stroke-dasharray="10" stroke-dashoffset="10"><path d="M8 8h8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="10;0"/></path><path d="M8 12h8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="10;0"/></path></g><path stroke-dasharray="7" stroke-dashoffset="7" d="M8 16h5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="7;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextBoxMultiple(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="62" stroke-dashoffset="62" d="M22 4V3C22 2.45 21.55 2 21 2H7C6.45 2 6 2.45 6 3V17C6 17.55 6.45 18 7 18H21C21.55 18 22 17.55 22 17z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="62;124"/></path><g stroke-dasharray="10" stroke-dashoffset="10"><path d="M10 6h8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="10;0"/></path><path d="M10 10h8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="10;0"/></path></g><path stroke-dasharray="7" stroke-dashoffset="7" d="M10 14h5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="7;0"/></path><path stroke-dasharray="34" stroke-dashoffset="34" d="M2 6V21C2 21.55 2.45 22 3 22H18"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.4s" dur="0.4s" values="34;68"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextBoxMultipleToTextBoxTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 4V3C22 2.45 21.55 2 21 2H7C6.45 2 6 2.45 6 3V17C6 17.55 6.45 18 7 18H21C21.55 18 22 17.55 22 17zM10 6h8M10 10h8M10 14h5"><animateMotion fill="freeze" begin="0.4s" calcMode="linear" dur="0.4s" path="M0 0l-2 2"/></path><path stroke-dasharray="34" stroke-dashoffset="68" d="M2 6V21C2 21.55 2.45 22 3 22H18"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="68;34"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextBoxMultipleTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="62" stroke-dashoffset="62" d="M22 4V3C22 2.45 21.55 2 21 2H7C6.45 2 6 2.45 6 3V17C6 17.55 6.45 18 7 18H21C21.55 18 22 17.55 22 17z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="62;124"/><animate fill="freeze" attributeName="fill-opacity" begin="1.8s" dur="0.15s" values="0;0.3"/></path><g stroke-dasharray="10" stroke-dashoffset="10"><path d="M10 6h8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="10;0"/></path><path d="M10 10h8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="10;0"/></path></g><path stroke-dasharray="7" stroke-dashoffset="7" d="M10 14h5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="7;0"/></path><path stroke-dasharray="34" stroke-dashoffset="34" d="M2 6V21C2 21.55 2.45 22 3 22H18"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.4s" dur="0.4s" values="34;68"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextBoxMultipleTwotoneToTextBoxTwotoneTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity=".3" d="M22 4V3C22 2.45 21.55 2 21 2H7C6.45 2 6 2.45 6 3V17C6 17.55 6.45 18 7 18H21C21.55 18 22 17.55 22 17zM10 6h8M10 10h8M10 14h5"><animateMotion fill="freeze" begin="0.4s" calcMode="linear" dur="0.4s" path="M0 0l-2 2"/></path><path stroke-dasharray="34" stroke-dashoffset="68" d="M2 6V21C2 21.55 2.45 22 3 22H18"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="68;34"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextBoxToTextBoxMultipleTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 6V5C20 4.45 19.55 4 19 4H5C4.45 4 4 4.45 4 5V19C4 19.55 4.45 20 5 20H19C19.55 20 20 19.55 20 19zM8 8h8M8 12h8M8 16h5"><animateMotion fill="freeze" calcMode="linear" dur="0.4s" path="M0 0l2-2"/></path><path stroke-dasharray="34" stroke-dashoffset="34" d="M2 6V21C2 21.55 2.45 22 3 22H18"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.4s" values="34;68"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextBoxTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity="0" stroke-dasharray="62" stroke-dashoffset="62" d="M20 6V5C20 4.45 19.55 4 19 4H5C4.45 4 4 4.45 4 5V19C4 19.55 4.45 20 5 20H19C19.55 20 20 19.55 20 19z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="62;124"/><animate fill="freeze" attributeName="fill-opacity" begin="1.3s" dur="0.15s" values="0;0.3"/></path><g stroke-dasharray="10" stroke-dashoffset="10"><path d="M8 8h8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="10;0"/></path><path d="M8 12h8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="10;0"/></path></g><path stroke-dasharray="7" stroke-dashoffset="7" d="M8 16h5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="7;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextBoxTwotoneToTextBoxMultipleTwotoneTransition(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" fill-opacity=".3" d="M20 6V5C20 4.45 19.55 4 19 4H5C4.45 4 4 4.45 4 5V19C4 19.55 4.45 20 5 20H19C19.55 20 20 19.55 20 19zM8 8h8M8 12h8M8 16h5"><animateMotion fill="freeze" calcMode="linear" dur="0.4s" path="M0 0l2-2"/></path><path stroke-dasharray="34" stroke-dashoffset="34" d="M2 6V21C2 21.55 2.45 22 3 22H18"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.4s" values="34;68"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsDown(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="80" stroke-dashoffset="80" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 4H18L21 11V14H14L15 20L12 21L7 13H3V4H7V13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.8s" values="80;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsDownTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M3 4H7V13H3V4Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke="currentColor" stroke-dasharray="80" stroke-dashoffset="80" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 4H18L21 11V14H14L15 20L12 21L7 13H3V4H7V13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.8s" values="80;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsUp(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="80" stroke-dashoffset="80" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 11L12 3L15 4L14 10H21V13L18 20H7H3V11H7V20"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.8s" values="80;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsUpTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" d="M7 11V20H3V11H7Z"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.15s" values="0;0.3"/></path><path fill="none" stroke="currentColor" stroke-dasharray="80" stroke-dashoffset="80" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 11L12 3L15 4L14 10H21V13L18 20H7H3V11H7V20"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.8s" values="80;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tiktok(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdTiktok0"><path fill="#fff" d="M16.6 5.82C15.9164 5.03962 15.5397 4.03743 15.54 3H12.45V15.4C12.4262 16.071 12.1429 16.7066 11.6598 17.1729C11.1767 17.6393 10.5314 17.8999 9.86 17.9C8.44 17.9 7.26 16.74 7.26 15.3C7.26 13.58 8.92 12.29 10.63 12.82V9.66C7.18 9.2 4.16 11.88 4.16 15.3C4.16 18.63 6.92 21 9.85 21C12.99 21 15.54 18.45 15.54 15.3V9.01C16.793 9.90985 18.2974 10.3926 19.84 10.39V7.3C19.84 7.3 17.96 7.39 16.6 5.82Z"/><g fill="none" stroke="#000" stroke-width="4"><path stroke-dasharray="36" stroke-dashoffset="72" d="M11 11H10C7.79086 11 5.5 12.7909 5.5 15C5.5 17.2091 7 19.5 10 19.5C12.2091 19.5 14 17.2091 14 15V2.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="72;36"/></path><path stroke-dasharray="10" stroke-dashoffset="20" d="M18 2.5V10.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.1s" values="20;10"/></path></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdTiktok0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TurnLeft(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="24" stroke-dashoffset="24" d="M16 19V11C16 10.4477 15.5523 10 15 10H4"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="24;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M4 10l3 -3M4 10l3 3" opacity="0"><set attributeName="opacity" begin="0.4s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TurnSharpLeft(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="28" stroke-dashoffset="28" d="M17 20V15C17 14.4477 16.5523 14 16 14H8C7.44772 14 7 13.5523 7 13V4"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="28;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M7 4l-3 3M7 4l3 3" opacity="0"><set attributeName="opacity" begin="0.4s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TurnSlightLeft(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="20" stroke-dashoffset="20" d="M14 19V12L7 5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M7 5h4M7 5v4" opacity="0"><set attributeName="opacity" begin="0.4s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitter(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="62" stroke-dashoffset="62" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.8906 7.34375C19.7969 7.67188 19.4001 8.50548 18.7219 9.29669C18.2698 17.9717 9.84907 20.7974 4.08456 17.8869C3.29335 16.8414 6.93856 17.2653 8.26666 15.259C3.23684 12.6876 3.63244 5.82103 4.64971 6.1036C7.02333 9.29669 10.8381 9.57926 11.4597 9.29669C11.4597 8.562 11.1489 6.97958 12.8726 5.65148C13.8616 4.94505 15.9297 4.3125 17.8047 6.34375C18.125 6.55469 18.5859 6.64844 19.2734 6.49219C19.6797 6.28125 20.2262 6.427 19.9453 7.15625"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="62;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwitterTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="62" stroke-dashoffset="62" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.8906 7.34375C19.7969 7.67188 19.4001 8.50548 18.7219 9.29669C18.2698 17.9717 9.84907 20.7974 4.08456 17.8869C3.29335 16.8414 6.93856 17.2653 8.26666 15.259C3.23684 12.6876 3.63244 5.82103 4.64971 6.1036C7.02333 9.29669 10.8381 9.57926 11.4597 9.29669C11.4597 8.562 11.1489 6.97958 12.8726 5.65148C13.8616 4.94505 15.9297 4.3125 17.8047 6.34375C18.125 6.55469 18.5859 6.64844 19.2734 6.49219C19.6797 6.28125 20.2262 6.427 19.9453 7.15625"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="62;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.15s" values="0;0.3"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwitterX(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M1 2h2.5L3.5 2h-2.5z"><animate fill="freeze" attributeName="d" dur="0.4s" values="M1 2h2.5L3.5 2h-2.5z;M1 2h2.5L18.5 22h-2.5z"/></path><path d="M5.5 2h2.5L7.2 2h-2.5z"><animate fill="freeze" attributeName="d" dur="0.4s" values="M5.5 2h2.5L7.2 2h-2.5z;M5.5 2h2.5L23 22h-2.5z"/></path><path d="M3 2h5v0h-5z" opacity="0"><set attributeName="opacity" begin="0.4s" to="1"/><animate fill="freeze" attributeName="d" begin="0.4s" dur="0.4s" values="M3 2h5v0h-5z;M3 2h5v2h-5z"/></path><path d="M16 22h5v0h-5z" opacity="0"><set attributeName="opacity" begin="0.4s" to="1"/><animate fill="freeze" attributeName="d" begin="0.4s" dur="0.4s" values="M16 22h5v0h-5z;M16 22h5v-2h-5z"/></path><path d="M18.5 2h3.5L22 2h-3.5z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="d" begin="0.5s" dur="0.4s" values="M18.5 2h3.5L22 2h-3.5z;M18.5 2h3.5L5 22h-3.5z"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwitterXalt(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 2h2.5L11 2h-2.5zM13 2h2.5L15.5 2h-2.5zM10.5 2h5v0h-5zM8.5 2h5v0h-5zM10 2h3.5L13.5 2h-3.5z"><animate fill="freeze" attributeName="d" dur="0.8s" keyTimes="0;0.3;0.5;1" values="M8.5 2h2.5L11 2h-2.5zM13 2h2.5L15.5 2h-2.5zM10.5 2h5v0h-5zM8.5 2h5v0h-5zM10 2h3.5L13.5 2h-3.5z;M8.5 2h2.5L11 22h-2.5zM13 2h2.5L15.5 22h-2.5zM10.5 2h5v2h-5zM8.5 20h5v2h-5zM10 2h3.5L13.5 22h-3.5z;M8.5 2h2.5L11 22h-2.5zM13 2h2.5L15.5 22h-2.5zM10.5 2h5v2h-5zM8.5 20h5v2h-5zM10 2h3.5L13.5 22h-3.5z;M1 2h2.5L18.5 22h-2.5zM5.5 2h2.5L23 22h-2.5zM3 2h5v2h-5zM16 20h5v2h-5zM18.5 2h3.5L5 22h-3.5z"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UturnLeft(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="36" stroke-dashoffset="36" d="M7 16V9C7 6.23858 9.23858 4 12 4C14.7614 4 17 6.23858 17 9V20"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="36;72"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M7 16l-3 -3M7 16l3 -3" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="none" stroke-dasharray="14" stroke-dashoffset="14" d="M6 19h12"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="14;0"/></path><path fill="currentColor" d="M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5"><animate attributeName="d" calcMode="linear" dur="1.5s" keyTimes="0;0.7;1" repeatCount="indefinite" values="M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5;M12 15 h2 v-3 h2.5 L12 7.5M12 15 h-2 v-3 h-2.5 L12 7.5;M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadOffLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdUploadOffLoop0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M6 19h12"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="14;0"/></path><path fill="#fff" d="M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5"><animate attributeName="d" calcMode="linear" dur="1.5s" keyTimes="0;0.7;1" repeatCount="indefinite" values="M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5;M12 15 h2 v-3 h2.5 L12 7.5M12 15 h-2 v-3 h-2.5 L12 7.5;M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5"/></path><g stroke-dasharray="26" stroke-dashoffset="26" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path d="M0 13h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 13h22;M2 13h22;M0 13h22"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="26;0"/></g></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdUploadOffLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadOffOutline(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdUploadOffOutline0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M6 19h12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="14;0"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="18;0"/></path><g stroke-dasharray="26" stroke-dashoffset="26" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path d="M1 13h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="26;0"/></g></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdUploadOffOutline0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadOffOutlineLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdUploadOffOutlineLoop0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M6 19h12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="14;0"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="18;0"/><animate attributeName="d" calcMode="linear" dur="1.5s" keyTimes="0;0.7;1" repeatCount="indefinite" values="M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5;M12 15 h2 v-3 h2.5 L12 7.5M12 15 h-2 v-3 h-2.5 L12 7.5;M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5"/></path><g stroke-dasharray="26" stroke-dashoffset="26" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path d="M0 13h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 13h22;M2 13h22;M0 13h22"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="26;0"/></g></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdUploadOffOutlineLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadOutline(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M6 19h12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="14;0"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="18;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadOutlineLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M6 19h12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="14;0"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="18;0"/><animate attributeName="d" calcMode="linear" dur="1.5s" keyTimes="0;0.7;1" repeatCount="indefinite" values="M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5;M12 15 h2 v-3 h2.5 L12 7.5M12 15 h-2 v-3 h-2.5 L12 7.5;M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadingLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="2 4" stroke-dashoffset="6" d="M12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3"><animate attributeName="stroke-dashoffset" dur="0.6s" repeatCount="indefinite" values="6;0"/></path><path stroke-dasharray="30" stroke-dashoffset="30" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.1s" dur="0.3s" values="30;0"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M12 16v-7.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="10;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M12 8.5l3.5 3.5M12 8.5l-3.5 3.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="6;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ValignBaseline(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="20" stroke-dashoffset="20" d="M2.5 18.5H21.5M2.5 11.5H21.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;0"/></path><path stroke-dasharray="50" stroke-dashoffset="50" stroke-width="2" d="M12.75 6H18V18H6V6H12.75Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.4s" values="50;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ValignBaselineTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path fill="none" stroke-dasharray="20" stroke-dashoffset="20" d="M2.5 18.5H21.5M2.5 11.5H21.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;0"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="50" stroke-dashoffset="50" stroke-width="2" d="M12.75 6H18V18H6V6H12.75Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.4s" values="50;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.15s" values="0;0.3"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ValignBottom(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="20" stroke-dashoffset="20" d="M2.5 15.5H21.5M2.5 8.5H21.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;0"/></path><path stroke-dasharray="50" stroke-dashoffset="50" stroke-width="2" d="M12.75 9H18V21H6V9H12.75Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.4s" values="50;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ValignBottomTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path fill="none" stroke-dasharray="20" stroke-dashoffset="20" d="M2.5 15.5H21.5M2.5 8.5H21.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;0"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="50" stroke-dashoffset="50" stroke-width="2" d="M12.75 9H18V21H6V9H12.75Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.4s" values="50;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.15s" values="0;0.3"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ValignMiddle(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="20" stroke-dashoffset="20" d="M2.5 15.5H21.5M2.5 8.5H21.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;0"/></path><path stroke-dasharray="50" stroke-dashoffset="50" stroke-width="2" d="M12.75 6H18V18H6V6H12.75Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.4s" values="50;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ValignMiddleTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path fill="none" stroke-dasharray="20" stroke-dashoffset="20" d="M2.5 15.5H21.5M2.5 8.5H21.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;0"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="50" stroke-dashoffset="50" stroke-width="2" d="M12.75 6H18V18H6V6H12.75Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.4s" values="50;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.15s" values="0;0.3"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ValignTop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="20" stroke-dashoffset="20" d="M2.5 15.5H21.5M2.5 8.5H21.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;0"/></path><path stroke-dasharray="50" stroke-dashoffset="50" stroke-width="2" d="M12.75 3H18V15H6V3H12.75Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.4s" values="50;0"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ValignTopTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path fill="none" stroke-dasharray="20" stroke-dashoffset="20" d="M2.5 15.5H21.5M2.5 8.5H21.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;0"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="50" stroke-dashoffset="50" stroke-width="2" d="M12.75 3H18V15H6V3H12.75Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.4s" values="50;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.15s" values="0;0.3"/></path></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Watch(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><clipPath id="lineMdWatch0"><rect width="24" height="12"/></clipPath><symbol id="lineMdWatch1"><path fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M23 16.5C23 10.4249 18.0751 5.5 12 5.5C5.92487 5.5 1 10.4249 1 16.5z" clip-path="url(#lineMdWatch0)"><animate fill="freeze" attributeName="d" dur="0.5s" values="M23 16.5C23 11.5 18.0751 12 12 12C5.92487 12 1 11.5 1 16.5z;M23 16.5C23 10.4249 18.0751 5.5 12 5.5C5.92487 5.5 1 10.4249 1 16.5z"/></path></symbol><mask id="lineMdWatch2"><use href="#lineMdWatch1"/><use href="#lineMdWatch1" transform="rotate(180 12 12)"/><circle cx="12" cy="12" r="0" fill="#fff"><animate fill="freeze" attributeName="r" dur="0.2s" values="0;3"/></circle></mask></defs><rect width="24" height="24" fill="currentColor" mask="url(#lineMdWatch2)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WatchLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><clipPath id="lineMdWatchLoop0"><rect width="24" height="12"/></clipPath><symbol id="lineMdWatchLoop1"><path fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M23 16.5C23 10.4249 18.0751 5.5 12 5.5C5.92487 5.5 1 10.4249 1 16.5z" clip-path="url(#lineMdWatchLoop0)"><animate attributeName="d" dur="6s" keyTimes="0;0.07;0.93;1" repeatCount="indefinite" values="M23 16.5C23 11.5 18.0751 12 12 12C5.92487 12 1 11.5 1 16.5z;M23 16.5C23 10.4249 18.0751 5.5 12 5.5C5.92487 5.5 1 10.4249 1 16.5z;M23 16.5C23 10.4249 18.0751 5.5 12 5.5C5.92487 5.5 1 10.4249 1 16.5z;M23 16.5C23 11.5 18.0751 12 12 12C5.92487 12 1 11.5 1 16.5z"/></path></symbol><mask id="lineMdWatchLoop2"><use href="#lineMdWatchLoop1"/><use href="#lineMdWatchLoop1" transform="rotate(180 12 12)"/><circle cx="12" cy="12" r="0" fill="#fff"><animate attributeName="r" dur="6s" keyTimes="0;0.03;0.97;1" repeatCount="indefinite" values="0;3;3;0"/></circle></mask></defs><rect width="24" height="24" fill="currentColor" mask="url(#lineMdWatchLoop2)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WatchOff(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><clipPath id="lineMdWatchOff0"><rect width="24" height="12"/></clipPath><symbol id="lineMdWatchOff1"><path fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M23 16.5C23 10.4249 18.0751 5.5 12 5.5C5.92487 5.5 1 10.4249 1 16.5z" clip-path="url(#lineMdWatchOff0)"><animate fill="freeze" attributeName="d" dur="0.5s" values="M23 16.5C23 11.5 18.0751 12 12 12C5.92487 12 1 11.5 1 16.5z;M23 16.5C23 10.4249 18.0751 5.5 12 5.5C5.92487 5.5 1 10.4249 1 16.5z"/></path></symbol><mask id="lineMdWatchOff2"><use href="#lineMdWatchOff1"/><use href="#lineMdWatchOff1" transform="rotate(180 12 12)"/><circle cx="12" cy="12" r="0" fill="#fff"><animate fill="freeze" attributeName="r" dur="0.2s" values="0;3"/></circle><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M1 13h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="26;0"/></g></mask></defs><rect width="24" height="24" fill="currentColor" mask="url(#lineMdWatchOff2)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WatchOffLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><clipPath id="lineMdWatchOffLoop0"><rect width="24" height="12"/></clipPath><symbol id="lineMdWatchOffLoop1"><path fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M23 16.5C23 10.4249 18.0751 5.5 12 5.5C5.92487 5.5 1 10.4249 1 16.5z" clip-path="url(#lineMdWatchOffLoop0)"><animate attributeName="d" dur="6s" keyTimes="0;0.07;0.93;1" repeatCount="indefinite" values="M23 16.5C23 11.5 18.0751 12 12 12C5.92487 12 1 11.5 1 16.5z;M23 16.5C23 10.4249 18.0751 5.5 12 5.5C5.92487 5.5 1 10.4249 1 16.5z;M23 16.5C23 10.4249 18.0751 5.5 12 5.5C5.92487 5.5 1 10.4249 1 16.5z;M23 16.5C23 11.5 18.0751 12 12 12C5.92487 12 1 11.5 1 16.5z"/></path></symbol><mask id="lineMdWatchOffLoop2"><use href="#lineMdWatchOffLoop1"/><use href="#lineMdWatchOffLoop1" transform="rotate(180 12 12)"/><circle cx="12" cy="12" r="0" fill="#fff"><animate attributeName="r" dur="6s" keyTimes="0;0.03;0.97;1" repeatCount="indefinite" values="0;3;3;0"/></circle><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M0 13h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 13h22;M2 13h22;M0 13h22"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="26;0"/></g></mask></defs><rect width="24" height="24" fill="currentColor" mask="url(#lineMdWatchOffLoop2)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WatchOffTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><clipPath id="lineMdWatchOffTwotone0"><rect width="24" height="12"/></clipPath><symbol id="lineMdWatchOffTwotone1"><path fill="#fff" fill-opacity="0" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M23 16.5C23 10.4249 18.0751 5.5 12 5.5C5.92487 5.5 1 10.4249 1 16.5z" clip-path="url(#lineMdWatchOffTwotone0)"><animate fill="freeze" attributeName="d" dur="0.5s" values="M23 16.5C23 11.5 18.0751 12 12 12C5.92487 12 1 11.5 1 16.5z;M23 16.5C23 10.4249 18.0751 5.5 12 5.5C5.92487 5.5 1 10.4249 1 16.5z"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.15s" values="0;0.3"/></path></symbol><mask id="lineMdWatchOffTwotone2"><use href="#lineMdWatchOffTwotone1"/><use href="#lineMdWatchOffTwotone1" transform="rotate(180 12 12)"/><circle cx="12" cy="12" r="0" fill="#fff"><animate fill="freeze" attributeName="r" dur="0.2s" values="0;3"/></circle><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M1 13h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="26;0"/></g></mask></defs><rect width="24" height="24" fill="currentColor" mask="url(#lineMdWatchOffTwotone2)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WatchOffTwotoneLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><clipPath id="lineMdWatchOffTwotoneLoop0"><rect width="24" height="12"/></clipPath><symbol id="lineMdWatchOffTwotoneLoop1"><path fill="#fff" fill-opacity="0" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M23 16.5C23 10.4249 18.0751 5.5 12 5.5C5.92487 5.5 1 10.4249 1 16.5z" clip-path="url(#lineMdWatchOffTwotoneLoop0)"><animate attributeName="d" dur="6s" keyTimes="0;0.07;0.93;1" repeatCount="indefinite" values="M23 16.5C23 11.5 18.0751 12 12 12C5.92487 12 1 11.5 1 16.5z;M23 16.5C23 10.4249 18.0751 5.5 12 5.5C5.92487 5.5 1 10.4249 1 16.5z;M23 16.5C23 10.4249 18.0751 5.5 12 5.5C5.92487 5.5 1 10.4249 1 16.5z;M23 16.5C23 11.5 18.0751 12 12 12C5.92487 12 1 11.5 1 16.5z"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.15s" values="0;0.3"/></path></symbol><mask id="lineMdWatchOffTwotoneLoop2"><use href="#lineMdWatchOffTwotoneLoop1"/><use href="#lineMdWatchOffTwotoneLoop1" transform="rotate(180 12 12)"/><circle cx="12" cy="12" r="0" fill="#fff"><animate attributeName="r" dur="6s" keyTimes="0;0.03;0.97;1" repeatCount="indefinite" values="0;3;3;0"/></circle><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M0 13h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 13h22;M2 13h22;M0 13h22"/></path><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="26;0"/></g></mask></defs><rect width="24" height="24" fill="currentColor" mask="url(#lineMdWatchOffTwotoneLoop2)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WatchTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><clipPath id="lineMdWatchTwotone0"><rect width="24" height="12"/></clipPath><symbol id="lineMdWatchTwotone1"><path fill="#fff" fill-opacity="0" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M23 16.5C23 10.4249 18.0751 5.5 12 5.5C5.92487 5.5 1 10.4249 1 16.5z" clip-path="url(#lineMdWatchTwotone0)"><animate fill="freeze" attributeName="d" dur="0.5s" values="M23 16.5C23 11.5 18.0751 12 12 12C5.92487 12 1 11.5 1 16.5z;M23 16.5C23 10.4249 18.0751 5.5 12 5.5C5.92487 5.5 1 10.4249 1 16.5z"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.15s" values="0;0.3"/></path></symbol><mask id="lineMdWatchTwotone2"><use href="#lineMdWatchTwotone1"/><use href="#lineMdWatchTwotone1" transform="rotate(180 12 12)"/><circle cx="12" cy="12" r="0" fill="#fff"><animate fill="freeze" attributeName="r" dur="0.2s" values="0;3"/></circle></mask></defs><rect width="24" height="24" fill="currentColor" mask="url(#lineMdWatchTwotone2)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WatchTwotoneLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><clipPath id="lineMdWatchTwotoneLoop0"><rect width="24" height="12"/></clipPath><symbol id="lineMdWatchTwotoneLoop1"><path fill="#fff" fill-opacity="0" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M23 16.5C23 10.4249 18.0751 5.5 12 5.5C5.92487 5.5 1 10.4249 1 16.5z" clip-path="url(#lineMdWatchTwotoneLoop0)"><animate attributeName="d" dur="6s" keyTimes="0;0.07;0.93;1" repeatCount="indefinite" values="M23 16.5C23 11.5 18.0751 12 12 12C5.92487 12 1 11.5 1 16.5z;M23 16.5C23 10.4249 18.0751 5.5 12 5.5C5.92487 5.5 1 10.4249 1 16.5z;M23 16.5C23 10.4249 18.0751 5.5 12 5.5C5.92487 5.5 1 10.4249 1 16.5z;M23 16.5C23 11.5 18.0751 12 12 12C5.92487 12 1 11.5 1 16.5z"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.15s" values="0;0.3"/></path></symbol><mask id="lineMdWatchTwotoneLoop2"><use href="#lineMdWatchTwotoneLoop1"/><use href="#lineMdWatchTwotoneLoop1" transform="rotate(180 12 12)"/><circle cx="12" cy="12" r="0" fill="#fff"><animate attributeName="r" dur="6s" keyTimes="0;0.03;0.97;1" repeatCount="indefinite" values="0;3;3;0"/></circle></mask></defs><rect width="24" height="24" fill="currentColor" mask="url(#lineMdWatchTwotoneLoop2)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WeatherCloudyLoop(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdWeatherCloudyLoop0"><g fill="#fff"><circle cx="12" cy="11" r="6"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;11;12;13;12"/></circle><rect width="10" height="7" x="8" y="12"/><rect width="16" height="10" x="1" y="9" rx="5"><animate attributeName="x" dur="19s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="17" height="8" x="6" y="11" rx="4"><animate attributeName="x" dur="23s" repeatCount="indefinite" values="6;5;6;7;6"/></rect></g><circle cx="12" cy="11" r="4"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;11;12;13;12"/></circle><rect width="8" height="6" x="8" y="11"><animate attributeName="x" dur="30s" repeatCount="indefinite" values="8;7;8;9;8"/></rect><rect width="12" height="6" x="3" y="11" rx="3"><animate attributeName="x" dur="19s" repeatCount="indefinite" values="3;2;3;4;3"/></rect><rect width="13" height="4" x="8" y="13" rx="2"><animate attributeName="x" dur="23s" repeatCount="indefinite" values="8;7;8;9;8"/></rect></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdWeatherCloudyLoop0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Youtube(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 5C21 5 21 5 21 12C21 19 21 19 12 19C3 19 3 19 3 12C3 5 3 5 12 5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="60;0"/></path><path fill="currentColor" fill-opacity="0" d="M10 8.5L16 12L10 15.5z"><set attributeName="fill-opacity" begin="0.7s" to="1"/><animate fill="freeze" attributeName="d" begin="0.7s" dur="0.2s" values="M12 11L12 12L12 13z;M10 8.5L16 12L10 15.5z"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YoutubeFilled(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="lineMdYoutubeFilled0"><path fill="#fff" fill-opacity="0" stroke="#fff" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 5C21 5 21 5 21 12C21 19 21 19 12 19C3 19 3 19 3 12C3 5 3 5 12 5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.5s" values="0;1"/></path><path d="M10 8.5L16 12L10 15.5z"/></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdYoutubeFilled0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YoutubeTwotone(children ...ElementRenderer) *LineMdIcon {
	return &LineMdIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity="0" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 5C21 5 21 5 21 12C21 19 21 19 12 19C3 19 3 19 3 12C3 5 3 5 12 5z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.15s" values="0;0.3"/></path><path fill="currentColor" fill-opacity="0" d="M10 8.5L16 12L10 15.5z"><set attributeName="fill-opacity" begin="0.8s" to="1"/><animate fill="freeze" attributeName="d" begin="0.8s" dur="0.2s" values="M12 11L12 12L12 13z;M10 8.5L16 12L10 15.5z"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
