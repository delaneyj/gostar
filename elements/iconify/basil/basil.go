package basil

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = ""
	hAttr          = "1em"
	viewbox        = "0 0 24 24"
	fill           = "currentColor"
)

type BasilIcon struct {
	*SVGSVGElement
}

type BasilIconFn func(children ...ElementRenderer) *BasilIcon

var IconLookup = map[string]BasilIconFn{
	"addOutline":                     AddOutline,
	"addSolid":                       AddSolid,
	"adobeAfterEffectsOutline":       AdobeAfterEffectsOutline,
	"adobeAfterEffectsSolid":         AdobeAfterEffectsSolid,
	"adobeExperinceDesignOutline":    AdobeExperinceDesignOutline,
	"adobeExperinceDesignSolid":      AdobeExperinceDesignSolid,
	"adobeIllustratorOutline":        AdobeIllustratorOutline,
	"adobeIllustratorSolid":          AdobeIllustratorSolid,
	"adobeIndesignOutline":           AdobeIndesignOutline,
	"adobeIndesignSolid":             AdobeIndesignSolid,
	"adobeLightroomOutline":          AdobeLightroomOutline,
	"adobeLightroomSolid":            AdobeLightroomSolid,
	"adobePhotoshopOutline":          AdobePhotoshopOutline,
	"adobePhotoshopSolid":            AdobePhotoshopSolid,
	"adobePremiereOutline":           AdobePremiereOutline,
	"adobePremiereSolid":             AdobePremiereSolid,
	"alarmOutline":                   AlarmOutline,
	"alarmSolid":                     AlarmSolid,
	"androidOutline":                 AndroidOutline,
	"androidSolid":                   AndroidSolid,
	"appStoreOutline":                AppStoreOutline,
	"appStoreSolid":                  AppStoreSolid,
	"appleOutline":                   AppleOutline,
	"appleSolid":                     AppleSolid,
	"appsOutline":                    AppsOutline,
	"appsSolid":                      AppsSolid,
	"arrowDownOutline":               ArrowDownOutline,
	"arrowDownSolid":                 ArrowDownSolid,
	"arrowLeftOutline":               ArrowLeftOutline,
	"arrowLeftSolid":                 ArrowLeftSolid,
	"arrowRightOutline":              ArrowRightOutline,
	"arrowRightSolid":                ArrowRightSolid,
	"arrowUpOutline":                 ArrowUpOutline,
	"arrowUpSolid":                   ArrowUpSolid,
	"asanaOutline":                   AsanaOutline,
	"asanaSolid":                     AsanaSolid,
	"atSignOutline":                  AtSignOutline,
	"atSignSolid":                    AtSignSolid,
	"attachOutline":                  AttachOutline,
	"attachSolid":                    AttachSolid,
	"awardOutline":                   AwardOutline,
	"awardSolid":                     AwardSolid,
	"backspaceOutline":               BackspaceOutline,
	"backspaceSolid":                 BackspaceSolid,
	"bagOutline":                     BagOutline,
	"bagSolid":                       BagSolid,
	"bankOutline":                    BankOutline,
	"bankSolid":                      BankSolid,
	"batteryEmptyOutline":            BatteryEmptyOutline,
	"batteryFullOutline":             BatteryFullOutline,
	"batteryFullSolid":               BatteryFullSolid,
	"batteryLowOutline":              BatteryLowOutline,
	"batteryLowSolid":                BatteryLowSolid,
	"batteryMostOutline":             BatteryMostOutline,
	"batteryMostSolid":               BatteryMostSolid,
	"batteryQuarterOutline":          BatteryQuarterOutline,
	"batteryQuarterSolid":            BatteryQuarterSolid,
	"battryHalfOutline":              BattryHalfOutline,
	"battryHalfSolid":                BattryHalfSolid,
	"behanceOutline":                 BehanceOutline,
	"behanceSolid":                   BehanceSolid,
	"binocularOutline":               BinocularOutline,
	"binocularSolid":                 BinocularSolid,
	"bluetoothOutline":               BluetoothOutline,
	"bluetoothSolid":                 BluetoothSolid,
	"bookCheckOutline":               BookCheckOutline,
	"bookCheckSolid":                 BookCheckSolid,
	"bookMarkOutline":                BookMarkOutline,
	"bookMarkSolid":                  BookMarkSolid,
	"bookOpenOutline":                BookOpenOutline,
	"bookOpenSolid":                  BookOpenSolid,
	"bookOutline":                    BookOutline,
	"bookSolid":                      BookSolid,
	"bookmarkOutline":                BookmarkOutline,
	"bookmarkSolid":                  BookmarkSolid,
	"boxOutline":                     BoxOutline,
	"boxSolid":                       BoxSolid,
	"bullhornOutline":                BullhornOutline,
	"bullhornSolid":                  BullhornSolid,
	"calendarOutline":                CalendarOutline,
	"calendarSolid":                  CalendarSolid,
	"cameraOutline":                  CameraOutline,
	"cameraSolid":                    CameraSolid,
	"cancelOutline":                  CancelOutline,
	"cancelSolid":                    CancelSolid,
	"cardOutline":                    CardOutline,
	"cardSolid":                      CardSolid,
	"caretDownOutline":               CaretDownOutline,
	"caretDownSolid":                 CaretDownSolid,
	"caretLeftOutline":               CaretLeftOutline,
	"caretLeftSolid":                 CaretLeftSolid,
	"caretRightOutline":              CaretRightOutline,
	"caretRightSolid":                CaretRightSolid,
	"caretUpOutline":                 CaretUpOutline,
	"caretUpSolid":                   CaretUpSolid,
	"chartPieAltOutline":             ChartPieAltOutline,
	"chartPieAltSolid":               ChartPieAltSolid,
	"chartPieOutline":                ChartPieOutline,
	"chartPieSolid":                  ChartPieSolid,
	"chatOutline":                    ChatOutline,
	"chatSolid":                      ChatSolid,
	"checkOutline":                   CheckOutline,
	"checkSolid":                     CheckSolid,
	"checkedBoxOutline":              CheckedBoxOutline,
	"checkedBoxSolid":                CheckedBoxSolid,
	"chromeOutline":                  ChromeOutline,
	"chromeSolid":                    ChromeSolid,
	"clipboardAltOutline":            ClipboardAltOutline,
	"clipboardAltSolid":              ClipboardAltSolid,
	"clipboardOutline":               ClipboardOutline,
	"clipboardSolid":                 ClipboardSolid,
	"clockOutline":                   ClockOutline,
	"clockSolid":                     ClockSolid,
	"cloudCheckOutline":              CloudCheckOutline,
	"cloudCheckSolid":                CloudCheckSolid,
	"cloudDownloadOutline":           CloudDownloadOutline,
	"cloudDownloadSolid":             CloudDownloadSolid,
	"cloudOffOutline":                CloudOffOutline,
	"cloudOffSolid":                  CloudOffSolid,
	"cloudOutline":                   CloudOutline,
	"cloudSolid":                     CloudSolid,
	"cloudUploadOutline":             CloudUploadOutline,
	"cloudUploadSolid":               CloudUploadSolid,
	"collapseOutline":                CollapseOutline,
	"collapseSolid":                  CollapseSolid,
	"columnsOutline":                 ColumnsOutline,
	"columnsSolid":                   ColumnsSolid,
	"commentBlockOutline":            CommentBlockOutline,
	"commentBlockSolid":              CommentBlockSolid,
	"commentMinusOutline":            CommentMinusOutline,
	"commentMinusSolid":              CommentMinusSolid,
	"commentOutline":                 CommentOutline,
	"commentPlusOutline":             CommentPlusOutline,
	"commentPlusSolid":               CommentPlusSolid,
	"commentSolid":                   CommentSolid,
	"contactsOutline":                ContactsOutline,
	"contactsSolid":                  ContactsSolid,
	"copyOutline":                    CopyOutline,
	"copySolid":                      CopySolid,
	"crossOutline":                   CrossOutline,
	"crossSolid":                     CrossSolid,
	"currentLocationOutline":         CurrentLocationOutline,
	"currentLocationSolid":           CurrentLocationSolid,
	"cursorOutline":                  CursorOutline,
	"cursorSolid":                    CursorSolid,
	"desktopOutline":                 DesktopOutline,
	"desktopSolid":                   DesktopSolid,
	"dialpadOutline":                 DialpadOutline,
	"dialpadSolid":                   DialpadSolid,
	"diamondOutline":                 DiamondOutline,
	"diamondSolid":                   DiamondSolid,
	"dislikeOutline":                 DislikeOutline,
	"dislikeSolid":                   DislikeSolid,
	"documentOutline":                DocumentOutline,
	"documentSolid":                  DocumentSolid,
	"downloadOutline":                DownloadOutline,
	"downloadSolid":                  DownloadSolid,
	"dribbbleOutline":                DribbbleOutline,
	"dribbbleSolid":                  DribbbleSolid,
	"dropboxOutline":                 DropboxOutline,
	"dropboxSolid":                   DropboxSolid,
	"editAltOutline":                 EditAltOutline,
	"editAltSolid":                   EditAltSolid,
	"editOutline":                    EditOutline,
	"editSolid":                      EditSolid,
	"envelopeOpenOutline":            EnvelopeOpenOutline,
	"envelopeOpenSolid":              EnvelopeOpenSolid,
	"envelopeOutline":                EnvelopeOutline,
	"envelopeSolid":                  EnvelopeSolid,
	"exchangeOutline":                ExchangeOutline,
	"exchangeSolid":                  ExchangeSolid,
	"expandOutline":                  ExpandOutline,
	"expandSolid":                    ExpandSolid,
	"exploreOutline":                 ExploreOutline,
	"exploreSolid":                   ExploreSolid,
	"eyeClosedOutline":               EyeClosedOutline,
	"eyeClosedSolid":                 EyeClosedSolid,
	"eyeOutline":                     EyeOutline,
	"eyeSolid":                       EyeSolid,
	"facebookMessengerOutline":       FacebookMessengerOutline,
	"facebookMessengerSolid":         FacebookMessengerSolid,
	"facebookOutline":                FacebookOutline,
	"facebookSolid":                  FacebookSolid,
	"fastForwardOutline":             FastForwardOutline,
	"fastForwardSolid":               FastForwardSolid,
	"fastRewindOutline":              FastRewindOutline,
	"fastRewindSolid":                FastRewindSolid,
	"figmaOutline":                   FigmaOutline,
	"figmaSolid":                     FigmaSolid,
	"fileDownloadOutline":            FileDownloadOutline,
	"fileDownloadSolid":              FileDownloadSolid,
	"fileOutline":                    FileOutline,
	"fileSolid":                      FileSolid,
	"fileUploadOutline":              FileUploadOutline,
	"fileUploadSolid":                FileUploadSolid,
	"fileUserOutline":                FileUserOutline,
	"fileUserSolid":                  FileUserSolid,
	"filterOutline":                  FilterOutline,
	"filterSolid":                    FilterSolid,
	"fireOutline":                    FireOutline,
	"fireSolid":                      FireSolid,
	"flaskAltOutline":                FlaskAltOutline,
	"flaskAltSolid":                  FlaskAltSolid,
	"flaskOutline":                   FlaskOutline,
	"flaskSolid":                     FlaskSolid,
	"folderBlockOutline":             FolderBlockOutline,
	"folderBlockSolid":               FolderBlockSolid,
	"folderDeleteOutline":            FolderDeleteOutline,
	"folderDeleteSolid":              FolderDeleteSolid,
	"folderLockOutline":              FolderLockOutline,
	"folderLockSolid":                FolderLockSolid,
	"folderOpenOutline":              FolderOpenOutline,
	"folderOpenSolid":                FolderOpenSolid,
	"folderOutline":                  FolderOutline,
	"folderPlusOutline":              FolderPlusOutline,
	"folderPlusSolid":                FolderPlusSolid,
	"folderSolid":                    FolderSolid,
	"folderUserOutline":              FolderUserOutline,
	"folderUserSolid":                FolderUserSolid,
	"forwardOutline":                 ForwardOutline,
	"forwardSolid":                   ForwardSolid,
	"gamepadOutline":                 GamepadOutline,
	"gamepadSolid":                   GamepadSolid,
	"globeOutline":                   GlobeOutline,
	"globeSolid":                     GlobeSolid,
	"gmailOutline":                   GmailOutline,
	"gmailSolid":                     GmailSolid,
	"googleAltOutline":               GoogleAltOutline,
	"googleAltSolid":                 GoogleAltSolid,
	"googleDriveOutline":             GoogleDriveOutline,
	"googleDriveSolid":               GoogleDriveSolid,
	"googleOutline":                  GoogleOutline,
	"googlePlayOutline":              GooglePlayOutline,
	"googlePlaySolid":                GooglePlaySolid,
	"googleSolid":                    GoogleSolid,
	"groupOneHundredFiftyOneOutline": GroupOneHundredFiftyOneOutline,
	"groupTwentyThreeSolid":          GroupTwentyThreeSolid,
	"headphoneOutline":               HeadphoneOutline,
	"headphoneSolid":                 HeadphoneSolid,
	"headsetOutline":                 HeadsetOutline,
	"headsetSolid":                   HeadsetSolid,
	"heartOffOutline":                HeartOffOutline,
	"heartOffSolid":                  HeartOffSolid,
	"heartOutline":                   HeartOutline,
	"heartPlusOutline":               HeartPlusOutline,
	"heartSolid":                     HeartSolid,
	"heartbeatOutline":               HeartbeatOutline,
	"heartbeatSolid":                 HeartbeatSolid,
	"historyOutline":                 HistoryOutline,
	"historySolid":                   HistorySolid,
	"homeOutline":                    HomeOutline,
	"homeSolid":                      HomeSolid,
	"hotspotOutline":                 HotspotOutline,
	"hotspotSolid":                   HotspotSolid,
	"imageOutline":                   ImageOutline,
	"imageSolid":                     ImageSolid,
	"infoCircleOutline":              InfoCircleOutline,
	"infoCircleSolid":                InfoCircleSolid,
	"infoRectOutline":                InfoRectOutline,
	"infoRectSolid":                  InfoRectSolid,
	"infoTriangleOutline":            InfoTriangleOutline,
	"infoTriangleSolid":              InfoTriangleSolid,
	"instagramOutline":               InstagramOutline,
	"instagramSolid":                 InstagramSolid,
	"invoiceOutline":                 InvoiceOutline,
	"invoiceSolid":                   InvoiceSolid,
	"keyOutline":                     KeyOutline,
	"keySolid":                       KeySolid,
	"layoutOutline":                  LayoutOutline,
	"layoutSolid":                    LayoutSolid,
	"lightbulbAltOutline":            LightbulbAltOutline,
	"lightbulbAltSolid":              LightbulbAltSolid,
	"lightbulbOffOutline":            LightbulbOffOutline,
	"lightbulbOffSolid":              LightbulbOffSolid,
	"lightbulbOutline":               LightbulbOutline,
	"lightbulbSolid":                 LightbulbSolid,
	"lightningAltOutline":            LightningAltOutline,
	"lightningAltSolid":              LightningAltSolid,
	"lightningOutline":               LightningOutline,
	"lightningSolid":                 LightningSolid,
	"likeOutline":                    LikeOutline,
	"likeSolid":                      LikeSolid,
	"linkedinOutline":                LinkedinOutline,
	"linkedinSolid":                  LinkedinSolid,
	"locationCheckOutline":           LocationCheckOutline,
	"locationCheckSolid":             LocationCheckSolid,
	"locationOutline":                LocationOutline,
	"locationPlusOutline":            LocationPlusOutline,
	"locationPlusSolid":              LocationPlusSolid,
	"locationQuestionOutline":        LocationQuestionOutline,
	"locationQuestionSolid":          LocationQuestionSolid,
	"locationSolid":                  LocationSolid,
	"lockOutline":                    LockOutline,
	"lockSolid":                      LockSolid,
	"lockTimeOutline":                LockTimeOutline,
	"lockTimeSolid":                  LockTimeSolid,
	"loginOutline":                   LoginOutline,
	"loginSolid":                     LoginSolid,
	"logoutOutline":                  LogoutOutline,
	"logoutSolid":                    LogoutSolid,
	"mapLocationOutline":             MapLocationOutline,
	"mapLocationSolid":               MapLocationSolid,
	"mediumOutline":                  MediumOutline,
	"mediumSolid":                    MediumSolid,
	"medkitOutline":                  MedkitOutline,
	"medkitSolid":                    MedkitSolid,
	"menuOutline":                    MenuOutline,
	"menuSolid":                      MenuSolid,
	"microphoneOffOutline":           MicrophoneOffOutline,
	"microphoneOffSolid":             MicrophoneOffSolid,
	"microphoneOutline":              MicrophoneOutline,
	"microphoneSolid":                MicrophoneSolid,
	"mobilePhoneOutline":             MobilePhoneOutline,
	"mobilePhoneSolid":               MobilePhoneSolid,
	"moonOutline":                    MoonOutline,
	"moonSolid":                      MoonSolid,
	"mouseAltOutline":                MouseAltOutline,
	"mouseAltSolid":                  MouseAltSolid,
	"mouseOutline":                   MouseOutline,
	"mouseSolid":                     MouseSolid,
	"moveOutline":                    MoveOutline,
	"moveSolid":                      MoveSolid,
	"musicOutline":                   MusicOutline,
	"musicSolid":                     MusicSolid,
	"navigationOutline":              NavigationOutline,
	"navigationSolid":                NavigationSolid,
	"notificationOffOutline":         NotificationOffOutline,
	"notificationOffSolid":           NotificationOffSolid,
	"notificationOnOutline":          NotificationOnOutline,
	"notificationOnSolid":            NotificationOnSolid,
	"notificationOutline":            NotificationOutline,
	"notificationSolid":              NotificationSolid,
	"notionOutline":                  NotionOutline,
	"notionSolid":                    NotionSolid,
	"otherOneOutline":                OtherOneOutline,
	"otherOneSolid":                  OtherOneSolid,
	"otherTwoOutline":                OtherTwoOutline,
	"paletteOutline":                 PaletteOutline,
	"paletteSolid":                   PaletteSolid,
	"pauseOutline":                   PauseOutline,
	"pauseSolid":                     PauseSolid,
	"phoneInOutline":                 PhoneInOutline,
	"phoneInSolid":                   PhoneInSolid,
	"phoneMissOutline":               PhoneMissOutline,
	"phoneMissSolid":                 PhoneMissSolid,
	"phoneOffOutline":                PhoneOffOutline,
	"phoneOffSolid":                  PhoneOffSolid,
	"phoneOutOutline":                PhoneOutOutline,
	"phoneOutSolid":                  PhoneOutSolid,
	"phoneOutline":                   PhoneOutline,
	"phoneSolid":                     PhoneSolid,
	"pictureOutline":                 PictureOutline,
	"pictureSolid":                   PictureSolid,
	"pinOutline":                     PinOutline,
	"pinSolid":                       PinSolid,
	"pinterestOutline":               PinterestOutline,
	"pinterestSolid":                 PinterestSolid,
	"playOutline":                    PlayOutline,
	"playSolid":                      PlaySolid,
	"plusOutline":                    PlusOutline,
	"plusSolid":                      PlusSolid,
	"powerButtonOutline":             PowerButtonOutline,
	"powerButtonSolid":               PowerButtonSolid,
	"presentOutline":                 PresentOutline,
	"presentSolid":                   PresentSolid,
	"printerOutline":                 PrinterOutline,
	"printerSolid":                   PrinterSolid,
	"processorOutline":               ProcessorOutline,
	"processorSolid":                 ProcessorSolid,
	"pulseOutline":                   PulseOutline,
	"pulseSolid":                     PulseSolid,
	"qqOutline":                      QqOutline,
	"qqSolid":                        QqSolid,
	"redditOutline":                  RedditOutline,
	"redditSolid":                    RedditSolid,
	"refreshOutline":                 RefreshOutline,
	"refreshSolid":                   RefreshSolid,
	"replyOutline":                   ReplyOutline,
	"replySolid":                     ReplySolid,
	"rowsOutline":                    RowsOutline,
	"rowsSolid":                      RowsSolid,
	"sandWatchOutline":               SandWatchOutline,
	"sandWatchSolid":                 SandWatchSolid,
	"saveOutline":                    SaveOutline,
	"saveSolid":                      SaveSolid,
	"searchOutline":                  SearchOutline,
	"searchSolid":                    SearchSolid,
	"sendOutline":                    SendOutline,
	"sendSolid":                      SendSolid,
	"serverOutline":                  ServerOutline,
	"serverSolid":                    ServerSolid,
	"settingsAdjustOutline":          SettingsAdjustOutline,
	"settingsAdjustSolid":            SettingsAdjustSolid,
	"settingsAltOutline":             SettingsAltOutline,
	"settingsAltSolid":               SettingsAltSolid,
	"settingsOutline":                SettingsOutline,
	"settingsSolid":                  SettingsSolid,
	"shareBoxOutline":                ShareBoxOutline,
	"shareBoxSolid":                  ShareBoxSolid,
	"shareOutline":                   ShareOutline,
	"shareSolid":                     ShareSolid,
	"shieldOutline":                  ShieldOutline,
	"shieldSolid":                    ShieldSolid,
	"shoppingBagOutline":             ShoppingBagOutline,
	"shoppingBagSolid":               ShoppingBagSolid,
	"shoppingBasketOutline":          ShoppingBasketOutline,
	"shoppingBasketSolid":            ShoppingBasketSolid,
	"shoppingCartOutline":            ShoppingCartOutline,
	"shoppingCartSolid":              ShoppingCartSolid,
	"shuffleOutline":                 ShuffleOutline,
	"shuffleSolid":                   ShuffleSolid,
	"sketchOutline":                  SketchOutline,
	"sketchSolid":                    SketchSolid,
	"skipNextOutline":                SkipNextOutline,
	"skipNextSolid":                  SkipNextSolid,
	"skipPrevOutline":                SkipPrevOutline,
	"skipPrevSolid":                  SkipPrevSolid,
	"skypeOutline":                   SkypeOutline,
	"skypeSolid":                     SkypeSolid,
	"slackOutline":                   SlackOutline,
	"slackSolid":                     SlackSolid,
	"snapchatOutline":                SnapchatOutline,
	"snapchatSolid":                  SnapchatSolid,
	"sortOutline":                    SortOutline,
	"sortSolid":                      SortSolid,
	"stackOutline":                   StackOutline,
	"stackSolid":                     StackSolid,
	"starHalfOutline":                StarHalfOutline,
	"starHalfSolid":                  StarHalfSolid,
	"starOutline":                    StarOutline,
	"starSolid":                      StarSolid,
	"sunOutline":                     SunOutline,
	"sunSolid":                       SunSolid,
	"telegramOutline":                TelegramOutline,
	"telegramSolid":                  TelegramSolid,
	"timerOutline":                   TimerOutline,
	"timerSolid":                     TimerSolid,
	"toggleOffOutline":               ToggleOffOutline,
	"toggleOffSolid":                 ToggleOffSolid,
	"toggleOnOutline":                ToggleOnOutline,
	"toggleOnSolid":                  ToggleOnSolid,
	"trashAltOutline":                TrashAltOutline,
	"trashAltSolid":                  TrashAltSolid,
	"trashOutline":                   TrashOutline,
	"trashSolid":                     TrashSolid,
	"trelloOutline":                  TrelloOutline,
	"trelloSolid":                    TrelloSolid,
	"tumblrOutline":                  TumblrOutline,
	"tumblrSolid":                    TumblrSolid,
	"twitchOutline":                  TwitchOutline,
	"twitchSolid":                    TwitchSolid,
	"twitterOutline":                 TwitterOutline,
	"twitterSolid":                   TwitterSolid,
	"umbrellaOutline":                UmbrellaOutline,
	"umbrellaSolid":                  UmbrellaSolid,
	"universityOutline":              UniversityOutline,
	"universitySolid":                UniversitySolid,
	"unlockOutline":                  UnlockOutline,
	"unlockSolid":                    UnlockSolid,
	"uploadOutline":                  UploadOutline,
	"uploadSolid":                    UploadSolid,
	"userBlockOutline":               UserBlockOutline,
	"userBlockSolid":                 UserBlockSolid,
	"userClockOutline":               UserClockOutline,
	"userClockSolid":                 UserClockSolid,
	"userOutline":                    UserOutline,
	"userPlusOutline":                UserPlusOutline,
	"userPlusSolid":                  UserPlusSolid,
	"userSolid":                      UserSolid,
	"viberOutline":                   ViberOutline,
	"viberSolid":                     ViberSolid,
	"videoOutline":                   VideoOutline,
	"videoSolid":                     VideoSolid,
	"vkOutline":                      VkOutline,
	"vkSolid":                        VkSolid,
	"volumeDownOutline":              VolumeDownOutline,
	"volumeDownSolid":                VolumeDownSolid,
	"volumeOffOutline":               VolumeOffOutline,
	"volumeOffSolid":                 VolumeOffSolid,
	"volumeUpOutline":                VolumeUpOutline,
	"volumeUpSolid":                  VolumeUpSolid,
	"walletOutline":                  WalletOutline,
	"walletSolid":                    WalletSolid,
	"watchOutline":                   WatchOutline,
	"watchSolid":                     WatchSolid,
	"wechatSolid":                    WechatSolid,
	"whatsappOutline":                WhatsappOutline,
	"whatsappSolid":                  WhatsappSolid,
	"windowsOutline":                 WindowsOutline,
	"windowsSolid":                   WindowsSolid,
	"youtubeOutline":                 YoutubeOutline,
	"youtubeSolid":                   YoutubeSolid,
	"zoomInOutline":                  ZoomInOutline,
	"zoomInSolid":                    ZoomInSolid,
	"zoomOutOutline":                 ZoomOutOutline,
	"zoomOutSolid":                   ZoomOutSolid,
}

func AddOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.007 12a.75.75 0 0 1 .75-.75h3.493V7.757a.75.75 0 0 1 1.5 0v3.493h3.493a.75.75 0 1 1 0 1.5H12.75v3.493a.75.75 0 0 1-1.5 0V12.75H7.757a.75.75 0 0 1-.75-.75"/><path fill="currentColor" fill-rule="evenodd" d="M7.317 3.769a42.502 42.502 0 0 1 9.366 0c1.827.204 3.302 1.643 3.516 3.48c.37 3.157.37 6.346 0 9.503c-.215 1.837-1.69 3.275-3.516 3.48a42.5 42.5 0 0 1-9.366 0c-1.827-.205-3.302-1.643-3.516-3.48a40.903 40.903 0 0 1 0-9.503c.214-1.837 1.69-3.276 3.516-3.48m9.2 1.49a41.001 41.001 0 0 0-9.034 0A2.486 2.486 0 0 0 5.29 7.424a39.402 39.402 0 0 0 0 9.154a2.486 2.486 0 0 0 2.193 2.164c2.977.332 6.057.332 9.034 0a2.486 2.486 0 0 0 2.192-2.164a39.401 39.401 0 0 0 0-9.154a2.486 2.486 0 0 0-2.192-2.163" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.345 4.017a42.253 42.253 0 0 1 9.31 0c1.713.192 3.095 1.541 3.296 3.26a40.66 40.66 0 0 1 0 9.446c-.201 1.719-1.583 3.068-3.296 3.26a42.245 42.245 0 0 1-9.31 0c-1.713-.192-3.095-1.541-3.296-3.26a40.652 40.652 0 0 1 0-9.445a3.734 3.734 0 0 1 3.295-3.26M12 7.007a.75.75 0 0 1 .75.75v3.493h3.493a.75.75 0 1 1 0 1.5H12.75v3.493a.75.75 0 0 1-1.5 0V12.75H7.757a.75.75 0 0 1 0-1.5h3.493V7.757a.75.75 0 0 1 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobeAfterEffectsOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="m11.248 15.5l-.67-2.101H8.222L7.585 15.5H6.188l2.376-7.414h1.727l2.41 7.414zm-2.233-4.95l-.572 1.826h1.903l-.583-1.826a21.152 21.152 0 0 1-.31-1.133a73.53 73.53 0 0 0-.064-.253h-.022l-.114.465c-.074.31-.15.632-.238.921m8.669 2.079c0 .253-.022.451-.044.572h-3.586c.033.968.792 1.386 1.65 1.386c.627 0 1.078-.088 1.485-.242l.198.935c-.462.187-1.1.341-1.87.341c-1.738 0-2.761-1.078-2.761-2.717c0-1.485.902-2.882 2.618-2.882c1.749 0 2.31 1.43 2.31 2.607m-2.387-1.683c-.836 0-1.188.759-1.243 1.309h2.354c.01-.495-.21-1.309-1.111-1.309"/><path d="M17.258 2.833a47.721 47.721 0 0 0-10.516 0c-2.012.225-3.637 1.81-3.873 3.832a45.921 45.921 0 0 0 0 10.67c.236 2.022 1.86 3.607 3.873 3.832a47.77 47.77 0 0 0 10.516 0c2.012-.225 3.637-1.81 3.873-3.832a45.925 45.925 0 0 0 0-10.67c-.236-2.022-1.86-3.607-3.873-3.832m-10.35 1.49a46.22 46.22 0 0 1 10.184 0c1.33.15 2.395 1.199 2.55 2.517a44.421 44.421 0 0 1 0 10.32a2.89 2.89 0 0 1-2.55 2.516a46.216 46.216 0 0 1-10.184 0a2.89 2.89 0 0 1-2.55-2.516a44.421 44.421 0 0 1 0-10.32a2.89 2.89 0 0 1 2.55-2.516"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobeAfterEffectsSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.015 10.55l-.572 1.826h1.903l-.583-1.826a21.152 21.152 0 0 1-.31-1.133a73.53 73.53 0 0 0-.064-.253h-.022l-.114.465c-.074.31-.15.632-.238.921m6.282.396c-.836 0-1.188.759-1.243 1.309h2.354c.01-.495-.21-1.309-1.111-1.309"/><path fill="currentColor" fill-rule="evenodd" d="M6.77 3.082a47.472 47.472 0 0 1 10.46 0c1.899.212 3.43 1.707 3.653 3.613a45.67 45.67 0 0 1 0 10.61c-.223 1.906-1.754 3.401-3.652 3.614a47.468 47.468 0 0 1-10.461 0c-1.899-.213-3.43-1.708-3.653-3.613a45.672 45.672 0 0 1 0-10.611C3.34 4.789 4.871 3.294 6.77 3.082M11.248 15.5l-.67-2.101H8.222L7.585 15.5H6.188l2.376-7.414h1.727l2.41 7.414zm6.436-2.871c0 .253-.022.451-.044.572h-3.586c.033.968.792 1.386 1.65 1.386c.627 0 1.078-.088 1.485-.242l.198.935c-.462.187-1.1.341-1.87.341c-1.738 0-2.761-1.078-2.761-2.717c0-1.485.902-2.882 2.618-2.882c1.749 0 2.31 1.43 2.31 2.607" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobeExperinceDesignOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.076 10.726V7.69h1.353v6.303c0 .55.022 1.144.044 1.507h-1.21l-.055-.847h-.022c-.32.594-.968.968-1.75.968c-1.275 0-2.287-1.089-2.287-2.739c-.011-1.793 1.11-2.86 2.398-2.86c.737 0 1.265.308 1.507.704zm-.044 2.849c.033-.121.044-.253.044-.396v-.803c0-.11-.011-.242-.033-.352c-.121-.528-.55-.957-1.166-.957c-.87 0-1.353.77-1.353 1.771c0 .979.484 1.694 1.342 1.694c.55 0 1.034-.374 1.166-.957" clip-rule="evenodd"/><path fill="currentColor" d="m9.904 11.705l2.255 3.795h-1.562l-.803-1.474a27.018 27.018 0 0 1-.737-1.397h-.033c-.176.429-.374.825-.671 1.397L7.616 15.5H6.065l2.2-3.751l-2.112-3.663h1.55l.804 1.529c.242.462.418.825.616 1.243h.022c.198-.462.352-.792.583-1.243l.792-1.529h1.55z"/><path fill="currentColor" fill-rule="evenodd" d="M17.258 2.833a47.721 47.721 0 0 0-10.516 0c-2.012.225-3.637 1.81-3.873 3.832a45.921 45.921 0 0 0 0 10.67c.236 2.022 1.86 3.607 3.873 3.832a47.77 47.77 0 0 0 10.516 0c2.012-.225 3.637-1.81 3.873-3.832a45.925 45.925 0 0 0 0-10.67c-.236-2.022-1.86-3.607-3.873-3.832m-10.35 1.49a46.22 46.22 0 0 1 10.184 0c1.33.15 2.395 1.199 2.55 2.517a44.421 44.421 0 0 1 0 10.32a2.89 2.89 0 0 1-2.55 2.516a46.216 46.216 0 0 1-10.184 0a2.89 2.89 0 0 1-2.55-2.516a44.421 44.421 0 0 1 0-10.32a2.89 2.89 0 0 1 2.55-2.516" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobeExperinceDesignSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.032 13.575c.033-.121.044-.253.044-.396v-.803c0-.11-.011-.242-.033-.352c-.121-.528-.55-.957-1.166-.957c-.87 0-1.353.77-1.353 1.771c0 .979.484 1.694 1.342 1.694c.55 0 1.034-.374 1.166-.957"/><path fill="currentColor" fill-rule="evenodd" d="M6.77 3.082a47.472 47.472 0 0 1 10.46 0c1.899.212 3.43 1.707 3.653 3.613a45.67 45.67 0 0 1 0 10.61c-.223 1.906-1.754 3.401-3.652 3.614a47.468 47.468 0 0 1-10.461 0c-1.899-.213-3.43-1.708-3.653-3.613a45.672 45.672 0 0 1 0-10.611C3.34 4.789 4.871 3.294 6.77 3.082m9.306 7.644V7.69h1.353v6.303c0 .55.022 1.144.044 1.507h-1.21l-.055-.847h-.022c-.32.594-.968.968-1.75.968c-1.275 0-2.287-1.089-2.287-2.739c-.011-1.793 1.11-2.86 2.398-2.86c.737 0 1.265.308 1.507.704zm-6.172.979l2.255 3.795h-1.562l-.803-1.474a27.018 27.018 0 0 1-.737-1.397h-.033c-.176.429-.374.825-.671 1.397L7.616 15.5H6.065l2.2-3.751l-2.112-3.663h1.55l.804 1.529c.242.462.418.825.616 1.243h.022c.198-.462.352-.792.583-1.243l.792-1.529h1.55z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobeIllustratorOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.933 10.143V15.5h-1.364v-5.357zm-1.441-1.496c0-.418.307-.737.758-.737c.463 0 .749.319.76.737c0 .407-.298.726-.77.726c-.451 0-.748-.319-.748-.726"/><path fill="currentColor" fill-rule="evenodd" d="m12.634 15.5l-.671-2.101H9.609L8.971 15.5H7.574L9.95 8.086h1.727l2.409 7.414zm-2.233-4.95l-.572 1.826h1.903l-.583-1.826c-.116-.365-.217-.766-.31-1.133a60.256 60.256 0 0 0-.064-.253h-.022c-.038.147-.075.304-.114.465c-.074.31-.151.632-.238.921" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M17.258 2.833a47.721 47.721 0 0 0-10.516 0c-2.012.225-3.637 1.81-3.873 3.832a45.921 45.921 0 0 0 0 10.67c.236 2.022 1.86 3.607 3.873 3.832a47.77 47.77 0 0 0 10.516 0c2.012-.225 3.637-1.81 3.873-3.832a45.925 45.925 0 0 0 0-10.67c-.236-2.022-1.86-3.607-3.873-3.832m-10.35 1.49a46.22 46.22 0 0 1 10.184 0c1.33.15 2.395 1.199 2.55 2.517a44.421 44.421 0 0 1 0 10.32a2.89 2.89 0 0 1-2.55 2.516a46.216 46.216 0 0 1-10.184 0a2.89 2.89 0 0 1-2.55-2.516a44.421 44.421 0 0 1 0-10.32a2.89 2.89 0 0 1 2.55-2.516" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobeIllustratorSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.401 10.55l-.572 1.826h1.903l-.583-1.826c-.116-.365-.217-.766-.31-1.133a60.256 60.256 0 0 0-.064-.253h-.022c-.038.147-.075.304-.114.465c-.074.31-.151.632-.238.921"/><path fill="currentColor" fill-rule="evenodd" d="M6.77 3.082a47.472 47.472 0 0 1 10.46 0c1.899.212 3.43 1.707 3.653 3.613a45.67 45.67 0 0 1 0 10.61c-.223 1.906-1.754 3.401-3.652 3.614a47.468 47.468 0 0 1-10.461 0c-1.899-.213-3.43-1.708-3.653-3.613a45.672 45.672 0 0 1 0-10.611C3.34 4.789 4.871 3.294 6.77 3.082m9.162 7.061V15.5h-1.364v-5.357zm-1.44-1.496c0-.418.307-.737.758-.737c.463 0 .749.319.76.737c0 .407-.298.726-.77.726c-.451 0-.748-.319-.748-.726M12.634 15.5l-.671-2.101H9.609L8.971 15.5H7.574L9.95 8.086h1.727l2.409 7.414z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobeIndesignOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.346 10.726V7.69H15.7v6.303c0 .55.022 1.144.044 1.507h-1.21l-.055-.847h-.022c-.319.594-.968.968-1.749.968c-1.276 0-2.288-1.089-2.288-2.739c-.01-1.793 1.111-2.86 2.398-2.86c.737 0 1.265.308 1.507.704zm-.044 2.849c.033-.121.044-.253.044-.396v-.803c0-.11-.01-.242-.033-.352c-.12-.528-.55-.957-1.166-.957c-.869 0-1.353.77-1.353 1.771c0 .979.484 1.694 1.342 1.694c.55 0 1.034-.374 1.166-.957" clip-rule="evenodd"/><path fill="currentColor" d="M8.355 15.5V8.086h1.353V15.5z"/><path fill="currentColor" fill-rule="evenodd" d="M17.258 2.833a47.721 47.721 0 0 0-10.516 0c-2.012.225-3.637 1.81-3.873 3.832a45.921 45.921 0 0 0 0 10.67c.236 2.022 1.86 3.607 3.873 3.832a47.77 47.77 0 0 0 10.516 0c2.012-.225 3.637-1.81 3.873-3.832a45.925 45.925 0 0 0 0-10.67c-.236-2.022-1.86-3.607-3.873-3.832m-10.35 1.49a46.22 46.22 0 0 1 10.184 0c1.33.15 2.395 1.199 2.55 2.517a44.421 44.421 0 0 1 0 10.32a2.89 2.89 0 0 1-2.55 2.516a46.216 46.216 0 0 1-10.184 0a2.89 2.89 0 0 1-2.55-2.516a44.421 44.421 0 0 1 0-10.32a2.89 2.89 0 0 1 2.55-2.516" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobeIndesignSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.302 13.575c.033-.121.044-.253.044-.396v-.803c0-.11-.01-.242-.033-.352c-.12-.528-.55-.957-1.166-.957c-.869 0-1.353.77-1.353 1.771c0 .979.484 1.694 1.342 1.694c.55 0 1.034-.374 1.166-.957"/><path fill="currentColor" fill-rule="evenodd" d="M6.77 3.082a47.472 47.472 0 0 1 10.46 0c1.899.212 3.43 1.707 3.653 3.613a45.67 45.67 0 0 1 0 10.61c-.223 1.906-1.754 3.401-3.652 3.614a47.468 47.468 0 0 1-10.461 0c-1.899-.213-3.43-1.708-3.653-3.613a45.672 45.672 0 0 1 0-10.611C3.34 4.789 4.871 3.294 6.77 3.082m7.576 7.644V7.69H15.7v6.303c0 .55.022 1.144.044 1.507h-1.21l-.055-.847h-.022c-.319.594-.968.968-1.749.968c-1.276 0-2.288-1.089-2.288-2.739c-.01-1.793 1.111-2.86 2.398-2.86c.737 0 1.265.308 1.507.704zM8.355 15.5V8.086h1.353V15.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobeLightroomOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.728 15.5h-4.4V8.086h1.353v6.281h3.047zm1.882 0h-1.352v-3.63c0-.726-.011-1.243-.044-1.727h1.177l.044 1.023h.044c.264-.759.89-1.144 1.463-1.144c.132 0 .209.011.319.033v1.276a1.881 1.881 0 0 0-.396-.044c-.65 0-1.09.418-1.21 1.023a2.322 2.322 0 0 0-.044.418z"/><path fill="currentColor" fill-rule="evenodd" d="M17.258 2.833a47.721 47.721 0 0 0-10.516 0c-2.012.225-3.637 1.81-3.873 3.832a45.921 45.921 0 0 0 0 10.67c.236 2.022 1.86 3.607 3.873 3.832a47.77 47.77 0 0 0 10.516 0c2.012-.225 3.637-1.81 3.873-3.832a45.925 45.925 0 0 0 0-10.67c-.236-2.022-1.86-3.607-3.873-3.832m-10.35 1.49a46.22 46.22 0 0 1 10.184 0c1.33.15 2.395 1.199 2.55 2.517a44.421 44.421 0 0 1 0 10.32a2.89 2.89 0 0 1-2.55 2.516a46.216 46.216 0 0 1-10.184 0a2.89 2.89 0 0 1-2.55-2.516a44.421 44.421 0 0 1 0-10.32a2.89 2.89 0 0 1 2.55-2.516" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobeLightroomSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.77 3.082a47.472 47.472 0 0 1 10.46 0c1.899.212 3.43 1.707 3.653 3.613a45.67 45.67 0 0 1 0 10.61c-.223 1.906-1.754 3.401-3.652 3.614a47.468 47.468 0 0 1-10.461 0c-1.899-.213-3.43-1.708-3.653-3.613a45.672 45.672 0 0 1 0-10.611C3.34 4.789 4.871 3.294 6.77 3.082M12.728 15.5h-4.4V8.086h1.353v6.281h3.047zm1.883 0h-1.353v-3.63c0-.726-.011-1.243-.044-1.727h1.177l.044 1.023h.044c.264-.759.89-1.144 1.463-1.144c.132 0 .209.011.319.033v1.276a1.881 1.881 0 0 0-.396-.044c-.65 0-1.09.418-1.21 1.023a2.322 2.322 0 0 0-.044.418z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobePhotoshopOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.095 15.5H7.753V8.185c.484-.088 1.155-.154 2.068-.154c.99 0 1.716.209 2.19.605c.44.363.725.946.725 1.639c0 .704-.22 1.287-.638 1.683c-.539.539-1.386.792-2.343.792c-.253 0-.484-.011-.66-.044zm.781-6.435c-.374 0-.638.033-.78.066v2.508c.164.044.384.055.66.055c1.011 0 1.627-.495 1.627-1.364c0-.825-.572-1.265-1.507-1.265" clip-rule="evenodd"/><path fill="currentColor" d="M14.465 15.621c-.65 0-1.232-.165-1.628-.385l.264-.979c.308.187.89.385 1.375.385c.594 0 .858-.242.858-.594c0-.363-.22-.55-.88-.781c-1.045-.363-1.485-.935-1.474-1.562c0-.946.78-1.683 2.024-1.683c.594 0 1.11.154 1.419.319l-.264.957a2.406 2.406 0 0 0-1.133-.308c-.484 0-.748.231-.748.561c0 .341.253.506.935.748c.968.352 1.419.847 1.43 1.639c0 .968-.76 1.683-2.178 1.683"/><path fill="currentColor" fill-rule="evenodd" d="M17.258 2.833a47.721 47.721 0 0 0-10.516 0c-2.012.225-3.637 1.81-3.873 3.832a45.922 45.922 0 0 0 0 10.67c.236 2.022 1.86 3.607 3.873 3.832a47.77 47.77 0 0 0 10.516 0c2.012-.225 3.637-1.81 3.873-3.832a45.914 45.914 0 0 0 0-10.67c-.236-2.022-1.86-3.607-3.873-3.832m-10.35 1.49a46.22 46.22 0 0 1 10.184 0c1.33.15 2.395 1.199 2.55 2.517a44.421 44.421 0 0 1 0 10.32a2.89 2.89 0 0 1-2.55 2.516a46.217 46.217 0 0 1-10.184 0a2.89 2.89 0 0 1-2.55-2.516a44.421 44.421 0 0 1 0-10.32a2.89 2.89 0 0 1 2.55-2.516" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobePhotoshopSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.876 9.065c-.374 0-.638.033-.78.066v2.508c.164.044.384.055.66.055c1.011 0 1.627-.495 1.627-1.364c0-.825-.572-1.265-1.507-1.265"/><path fill="currentColor" fill-rule="evenodd" d="M6.77 3.082a47.472 47.472 0 0 1 10.46 0c1.899.212 3.43 1.707 3.653 3.613a45.67 45.67 0 0 1 0 10.61c-.223 1.906-1.754 3.401-3.652 3.614a47.468 47.468 0 0 1-10.461 0c-1.899-.213-3.43-1.708-3.653-3.613a45.672 45.672 0 0 1 0-10.611C3.34 4.789 4.871 3.294 6.77 3.082M9.095 15.5H7.753V8.185c.484-.088 1.155-.154 2.068-.154c.99 0 1.716.209 2.19.605c.44.363.725.946.725 1.639c0 .704-.22 1.287-.638 1.683c-.539.539-1.386.792-2.343.792c-.253 0-.484-.011-.66-.044zm5.37.121c-.65 0-1.232-.165-1.628-.385l.264-.979c.308.187.89.385 1.375.385c.594 0 .858-.242.858-.594c0-.363-.22-.55-.88-.781c-1.045-.363-1.485-.935-1.474-1.562c0-.946.78-1.683 2.024-1.683c.594 0 1.11.154 1.419.319l-.264.957a2.406 2.406 0 0 0-1.133-.308c-.484 0-.748.231-.748.561c0 .341.253.506.935.748c.968.352 1.419.847 1.43 1.639c0 .968-.76 1.683-2.178 1.683" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobePremiereOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.908 4.324a46.22 46.22 0 0 1 10.184 0a2.89 2.89 0 0 1 2.55 2.516a44.421 44.421 0 0 1 0 10.32a2.89 2.89 0 0 1-2.55 2.516a46.216 46.216 0 0 1-10.184 0a2.89 2.89 0 0 1-2.55-2.516a44.421 44.421 0 0 1 0-10.32a2.89 2.89 0 0 1 2.55-2.516m10.35-1.49a47.721 47.721 0 0 0-10.516 0c-2.012.224-3.637 1.809-3.873 3.831a45.921 45.921 0 0 0 0 10.67c.236 2.022 1.86 3.607 3.873 3.832a47.77 47.77 0 0 0 10.516 0c2.012-.225 3.637-1.81 3.873-3.832a45.925 45.925 0 0 0 0-10.67c-.236-2.022-1.86-3.607-3.873-3.832M9.375 15.5H8.033V8.185c.484-.088 1.155-.154 2.068-.154c.99 0 1.716.209 2.189.605c.44.363.726.946.726 1.639c0 .704-.22 1.287-.638 1.683c-.54.539-1.386.792-2.343.792c-.253 0-.484-.011-.66-.044zm.78-6.435c-.373 0-.637.033-.78.066v2.508c.165.044.385.055.66.055c1.012 0 1.628-.495 1.628-1.364c0-.825-.572-1.265-1.507-1.265m4.751 6.435h-1.353v-3.63c0-.726-.01-1.243-.044-1.727h1.177l.044 1.023h.044c.264-.759.891-1.144 1.463-1.144c.132 0 .21.011.32.033v1.276a1.881 1.881 0 0 0-.397-.044c-.649 0-1.089.418-1.21 1.023a2.322 2.322 0 0 0-.044.418z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobePremiereSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.156 9.065c-.374 0-.638.033-.781.066v2.508c.165.044.385.055.66.055c1.012 0 1.628-.495 1.628-1.364c0-.825-.572-1.265-1.507-1.265"/><path fill="currentColor" fill-rule="evenodd" d="M6.77 3.082a47.472 47.472 0 0 1 10.46 0c1.899.212 3.43 1.707 3.653 3.613a45.67 45.67 0 0 1 0 10.61c-.223 1.906-1.754 3.401-3.652 3.614a47.468 47.468 0 0 1-10.461 0c-1.899-.213-3.43-1.708-3.653-3.613a45.672 45.672 0 0 1 0-10.611C3.34 4.789 4.871 3.294 6.77 3.082M9.375 15.5H8.033V8.185c.484-.088 1.155-.154 2.068-.154c.99 0 1.716.209 2.189.605c.44.363.726.946.726 1.639c0 .704-.22 1.287-.638 1.683c-.54.539-1.386.792-2.343.792c-.253 0-.484-.011-.66-.044zm5.531 0h-1.353v-3.63c0-.726-.01-1.243-.044-1.727h1.177l.044 1.023h.044c.264-.759.891-1.144 1.463-1.144c.132 0 .21.011.32.033v1.276a1.881 1.881 0 0 0-.397-.044c-.649 0-1.089.418-1.21 1.023a2.322 2.322 0 0 0-.044.418z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlarmOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M12 5.75a7.25 7.25 0 1 0 0 14.5a7.25 7.25 0 0 0 0-14.5M3.25 13a8.75 8.75 0 1 1 17.5 0a8.75 8.75 0 0 1-17.5 0"/><path d="M12 7.25a.75.75 0 0 1 .75.75v4.584l2.648 1.655a.75.75 0 1 1-.796 1.272l-3-1.875A.75.75 0 0 1 11.25 13V8a.75.75 0 0 1 .75-.75M6.53 3.47a.75.75 0 0 1 0 1.06l-2.5 2.5a.75.75 0 0 1-1.06-1.06l2.5-2.5a.75.75 0 0 1 1.06 0m10.94 0a.75.75 0 0 0 0 1.06l2.5 2.5a.75.75 0 1 0 1.06-1.06l-2.5-2.5a.75.75 0 0 0-1.06 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlarmSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.53 3.47a.75.75 0 0 1 0 1.06l-2.5 2.5a.75.75 0 0 1-1.06-1.06l2.5-2.5a.75.75 0 0 1 1.06 0"/><path fill="currentColor" fill-rule="evenodd" d="M12 4.5a8.5 8.5 0 1 0 0 17a8.5 8.5 0 0 0 0-17m.75 3.5a.75.75 0 0 0-1.5 0v5a.75.75 0 0 0 .352.636l3 1.875a.75.75 0 1 0 .796-1.272l-2.648-1.655z" clip-rule="evenodd"/><path fill="currentColor" d="M17.47 4.53a.75.75 0 0 1 1.06-1.06l2.5 2.5a.75.75 0 0 1-1.06 1.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AndroidOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.115 4.667a.846.846 0 1 1-1.692 0a.846.846 0 0 1 1.692 0m2.539.846a.846.846 0 1 0 0-1.692a.846.846 0 0 0 0 1.692"/><path fill="currentColor" fill-rule="evenodd" d="M6.354.47a.75.75 0 0 1 1.061 0l1.228 1.227a6.259 6.259 0 0 1 3.215-.883h.207c1.175 0 2.274.322 3.215.883L16.508.47a.75.75 0 0 1 1.06 1.06l-1.091 1.092a6.269 6.269 0 0 1 1.853 3.943a2.16 2.16 0 0 1 2.843 2.05v5.077a2.16 2.16 0 0 1-2.82 2.058v1.865c0 .913-.445 1.721-1.129 2.222v1.47a2.442 2.442 0 1 1-4.884 0v-.942h-.757v.943a2.442 2.442 0 1 1-4.884 0v-1.471a2.746 2.746 0 0 1-1.128-2.222V15.75a2.16 2.16 0 0 1-2.82-2.058V8.615a2.16 2.16 0 0 1 2.842-2.05a6.27 6.27 0 0 1 1.853-3.943L6.354 1.53a.75.75 0 0 1 0-1.06m.965 7.485a.25.25 0 0 0-.248.25v9.41a1.25 1.25 0 0 0 1.245 1.25h7.291a1.25 1.25 0 0 0 1.245-1.25v-9.41a.25.25 0 0 0-.248-.25zm9.287-1.5h.203a4.789 4.789 0 0 0-4.744-4.14h-.207a4.789 4.789 0 0 0-4.744 4.14h9.488zM13.84 21.308v-.943h1.884v.943a.942.942 0 1 1-1.884 0m-5.526-.943h1.769v.943a.942.942 0 1 1-1.884 0v-.943zm10.699-6.012a.66.66 0 0 1-.66-.658V8.612a.66.66 0 0 1 1.32.003v5.077a.66.66 0 0 1-.66.66M5.57 8.61a.66.66 0 0 0-1.32.004v5.077a.66.66 0 0 0 1.32.005z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AndroidSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.016 1.646a.5.5 0 0 1 .708 0l1.105 1.106A5.621 5.621 0 0 1 11.647 2h.706a5.62 5.62 0 0 1 2.82.753l1.107-1.107a.5.5 0 1 1 .707.708l-.984.984A5.635 5.635 0 0 1 18 7.648a.353.353 0 0 1-.353.352H6.353A.353.353 0 0 1 6 7.647c0-1.728.776-3.275 2-4.31l-.984-.983a.5.5 0 0 1 0-.708M10.5 5.25a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0m3.724.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5" clip-rule="evenodd"/><path fill="currentColor" d="M4 8.947a1 1 0 0 0-1 1v6a1 1 0 0 0 2 0v-6a1 1 0 0 0-1-1m3 0a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1zm12.21 1a1 1 0 1 1 2 0v6a1 1 0 0 1-2 0zM7.793 21a.293.293 0 0 0-.293.294v.206a1.75 1.75 0 1 0 3.5 0v-.206a.293.293 0 0 0-.293-.294zm5.207.294c0-.163.131-.294.293-.294h2.913c.163 0 .294.131.294.294v.206a1.75 1.75 0 1 1-3.5 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppStoreOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.196 17.942a2.069 2.069 0 0 0-.073-.322h6.962l1.575 2.546c.785 1.295 2.46 1.697 3.753.979c1.285-.715 1.797-2.32 1.095-3.637c1.125-.328 1.992-1.341 1.992-2.617c0-1.56-1.298-2.729-2.777-2.729h-2.438l-2.519-3.966a1.747 1.747 0 0 0-.83-.69l1.134-1.89c.803-1.338.301-3.023-1.026-3.76a2.824 2.824 0 0 0-3.043.197a2.825 2.825 0 0 0-3.045-.2c-1.327.738-1.829 2.423-1.026 3.762l.847 1.411l-3.081 5.136h-2.42C1.799 12.162.5 13.33.5 14.892c0 1.275.868 2.289 1.993 2.616c-.703 1.318-.192 2.924 1.094 3.639c1.285.713 2.95.32 3.74-.958l.48-.728a2.338 2.338 0 0 0 .404-1.307c0-.009 0-.005-.001-.001a4.124 4.124 0 0 1-.004-.05a2.494 2.494 0 0 0-.01-.161M6.06 19.386l.502-.763c.163-.233.15-.466.15-.466c-.103-.709-1.52-.857-1.52-.857c-.602-.055-.903.237-1.025.405l-.318.452a1.204 1.204 0 0 0 .467 1.678c.61.34 1.391.138 1.744-.45m11.651-5.723h3.012c.705 0 1.277.55 1.277 1.228c0 .679-.572 1.229-1.277 1.229h-1.521c-.23 0-.192.123-.17.168l1.12 1.867c.352.588.143 1.34-.468 1.679c-.61.339-1.392.137-1.744-.45l-3.936-6.36c-.239-.457-.925-1.968-.324-3.297c0 0 .503-1.193.82-.727l2.879 4.534l.002.004c.035.04.125.111.33.124M13.01 16.12c.622 0 .46-.787.46-.787c-.244-1.614-2.036-1.67-2.036-1.67h-1.71c-.17-.015-.162-.098-.14-.15l5.2-8.668a1.204 1.204 0 0 0-.467-1.678c-.61-.34-1.392-.138-1.744.45L12 4.568l-.573-.954c-.352-.588-1.133-.79-1.744-.45c-.61.34-.82 1.09-.467 1.678L10.45 6.9a.264.264 0 0 1 0 .252l-3.737 6.23l-.003-.002s-.11.257-.572.282H3.277c-.705 0-1.277.55-1.277 1.23c0 .678.572 1.228 1.277 1.228z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppStoreSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.474 20.299l-.602.942a1.54 1.54 0 0 1-2.093.555a1.51 1.51 0 0 1-.56-2.073l.381-.558c.147-.208.508-.568 1.23-.5c0 0 1.7.183 1.823 1.058c0 0 .017.288-.179.576m16.994-6.128h-3.615c-.246-.016-.353-.103-.396-.154l-.002-.005L15 8.412c-.38-.576-.984.897-.984.897c-.721 1.642.102 3.509.39 4.073l4.722 7.857a1.54 1.54 0 0 0 2.093.555a1.51 1.51 0 0 0 .56-2.073l-1.343-2.306c-.026-.056-.072-.207.204-.208h1.826c.846 0 1.532-.68 1.532-1.518s-.686-1.518-1.532-1.518m-8.704 2.064s.193.972-.554.972H1.532c-.846 0-1.532-.68-1.532-1.518s.686-1.518 1.532-1.518h3.435c.555-.032.686-.349.686-.349l.003.002L10.14 6.13a.336.336 0 0 0 0-.312L8.66 3.277a1.51 1.51 0 0 1 .56-2.073a1.54 1.54 0 0 1 2.094.555L12 2.938l.685-1.177a1.54 1.54 0 0 1 2.093-.555a1.51 1.51 0 0 1 .561 2.073L9.1 13.986c-.027.065-.035.167.168.185h2.046s2.156.07 2.45 2.064"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppleOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.71 1.377a.75.75 0 0 0-.896-.703c-.5.1-2.132.506-3.42 1.794c-.658.658-1.022 1.488-1.22 2.244a7.03 7.03 0 0 0-.194 1.122a33.657 33.657 0 0 0-.243-.104c-.693-.296-1.59-.64-2.61-.64c-1.092 0-2.642.444-3.922 1.7c-1.298 1.274-2.255 3.318-2.255 6.4c0 1.732.866 4.25 2.033 6.295c.59 1.033 1.285 1.995 2.035 2.704c.729.688 1.617 1.238 2.585 1.213c.72-.003 1.45-.286 2.017-.506l.05-.02c.654-.253 1.12-.424 1.511-.424c.374 0 .844.171 1.508.426l.049.018c.576.221 1.319.506 2.032.506c.539 0 1.162-.288 1.708-.627a9.608 9.608 0 0 0 1.796-1.476c1.137-1.182 2.217-2.871 2.175-4.679a.75.75 0 0 0-.428-.66a4.626 4.626 0 0 1-1.39-1.084c-.473-.543-.822-1.212-.844-1.988a5.549 5.549 0 0 1 .683-2.806c.422-.77.9-1.21 1.122-1.327a.75.75 0 0 0 .262-1.1c-.342-.48-.992-1.112-1.758-1.621c-.757-.504-1.744-.965-2.754-.945c-.177.004-.35.017-.52.037c.37-.55.586-1.232.712-1.83a8.23 8.23 0 0 0 .177-1.919m-3.255 2.151c.541-.541 1.177-.883 1.704-1.094a7.056 7.056 0 0 1-.093.553c-.131.622-.337 1.142-.59 1.438c-.48.559-1.124 1.169-1.988 1.41c.028-.237.072-.49.138-.744c.155-.595.42-1.153.829-1.563m2.917 3.06c.565-.01 1.242.261 1.894.695c.344.228.648.48.89.712c-.357.358-.708.83-1.001 1.366a7.048 7.048 0 0 0-.867 3.57c.034 1.22.584 2.21 1.212 2.93c.46.527.98.934 1.43 1.21c-.127 1.095-.824 2.24-1.737 3.188a8.115 8.115 0 0 1-1.507 1.242c-.508.316-.824.401-.916.401c-.4 0-.888-.174-1.544-.425l-.08-.03c-.564-.217-1.286-.494-1.965-.494c-.7 0-1.418.279-1.98.497l-.072.028c-.648.25-1.127.424-1.537.424H8.57c-.376.012-.894-.212-1.52-.803c-.608-.575-1.217-1.403-1.763-2.358c-1.104-1.934-1.836-4.187-1.836-5.55c0-2.757.845-4.388 1.805-5.33c.978-.96 2.141-1.272 2.873-1.272c.695 0 1.353.235 2.02.52l.27.117c.238.103.485.21.709.295c.312.117.679.229 1.054.229s.746-.105 1.07-.22c.241-.085.505-.194.76-.299l.252-.104c.693-.282 1.382-.523 2.109-.538" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppleSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.703 5.616c.03-.905.278-2.237 1.145-3.112c.805-.813 1.785-1.209 2.395-1.39a.347.347 0 0 1 .45.353c-.04.866-.268 2.158-.856 2.85c-.554.652-1.444 1.516-2.749 1.662a.348.348 0 0 1-.385-.363"/><path fill="currentColor" d="M13.937 22.17c.627.24 1.26.482 1.833.482c1.227 0 4.798-2.83 4.926-5.768a.399.399 0 0 0-.215-.36c-.908-.508-2.392-1.774-2.443-3.615c-.06-2.146 1.017-3.85 1.832-4.555c.184-.159.247-.439.083-.617c-.742-.81-2.234-1.925-3.596-1.898c-.99.02-1.883.39-2.65.706c-.587.243-1.1.455-1.526.455c-.403 0-.887-.21-1.443-.453c-.748-.326-1.626-.708-2.61-.708c-1.824 0-5.428 1.513-5.428 7.352c0 3.095 3.195 9.543 5.892 9.461c.582 0 1.204-.24 1.822-.48c.608-.236 1.211-.47 1.767-.47c.537 0 1.143.233 1.756.468"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppsOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.067 3.874a9.689 9.689 0 0 0-2.134 0a.062.062 0 0 0-.036.018a.049.049 0 0 0-.016.029a9.289 9.289 0 0 0 0 2.158a.05.05 0 0 0 .016.03c.01.01.023.016.036.017a9.637 9.637 0 0 0 2.134 0a.062.062 0 0 0 .036-.018a.05.05 0 0 0 .016-.03a9.289 9.289 0 0 0 0-2.157a.049.049 0 0 0-.016-.03a.062.062 0 0 0-.036-.017m-2.3-1.49a11.051 11.051 0 0 1 2.466 0a1.56 1.56 0 0 1 1.376 1.363c.097.832.097 1.674 0 2.506a1.559 1.559 0 0 1-1.376 1.364c-.813.09-1.653.09-2.466 0A1.559 1.559 0 0 1 2.39 6.253a10.788 10.788 0 0 1 0-2.506a1.559 1.559 0 0 1 1.376-1.364m9.301 1.491a9.689 9.689 0 0 0-2.134 0a.062.062 0 0 0-.036.018a.05.05 0 0 0-.016.029a9.293 9.293 0 0 0 0 2.158a.05.05 0 0 0 .016.03c.01.01.023.016.036.017a9.637 9.637 0 0 0 2.134 0a.062.062 0 0 0 .036-.018a.05.05 0 0 0 .016-.03a9.293 9.293 0 0 0 0-2.157a.05.05 0 0 0-.016-.03a.062.062 0 0 0-.036-.017m-2.3-1.49a11.051 11.051 0 0 1 2.466 0c.713.079 1.29.64 1.375 1.363c.098.832.098 1.674 0 2.506a1.559 1.559 0 0 1-1.375 1.364c-.813.09-1.653.09-2.466 0A1.559 1.559 0 0 1 9.39 6.253a10.788 10.788 0 0 1 0-2.506a1.559 1.559 0 0 1 1.376-1.364m9.301 1.491a9.689 9.689 0 0 0-2.134 0a.062.062 0 0 0-.036.018a.05.05 0 0 0-.016.029a9.293 9.293 0 0 0 0 2.158a.05.05 0 0 0 .016.03c.01.01.023.016.036.017a9.637 9.637 0 0 0 2.134 0a.062.062 0 0 0 .036-.018a.05.05 0 0 0 .016-.03a9.293 9.293 0 0 0 0-2.157a.05.05 0 0 0-.016-.03a.062.062 0 0 0-.036-.017m-2.3-1.49a11.051 11.051 0 0 1 2.466 0c.713.079 1.29.64 1.375 1.363c.098.832.098 1.674 0 2.506a1.559 1.559 0 0 1-1.375 1.364c-.813.09-1.653.09-2.466 0a1.559 1.559 0 0 1-1.375-1.364a10.787 10.787 0 0 1 0-2.506a1.559 1.559 0 0 1 1.375-1.364m-11.7 8.491a9.692 9.692 0 0 0-2.134 0a.062.062 0 0 0-.036.018a.049.049 0 0 0-.016.029a9.288 9.288 0 0 0 0 2.158a.05.05 0 0 0 .016.03c.01.01.023.016.036.017a9.637 9.637 0 0 0 2.134 0a.062.062 0 0 0 .036-.018a.05.05 0 0 0 .016-.03a9.288 9.288 0 0 0 0-2.157a.049.049 0 0 0-.016-.03a.062.062 0 0 0-.036-.017m-2.3-1.49a11.051 11.051 0 0 1 2.466 0a1.56 1.56 0 0 1 1.376 1.363c.097.832.097 1.674 0 2.506a1.559 1.559 0 0 1-1.376 1.364c-.813.09-1.653.09-2.466 0a1.559 1.559 0 0 1-1.376-1.364a10.788 10.788 0 0 1 0-2.506a1.559 1.559 0 0 1 1.376-1.364m9.3 1.491a9.693 9.693 0 0 0-2.134 0a.062.062 0 0 0-.036.018a.049.049 0 0 0-.016.029a9.293 9.293 0 0 0 0 2.158a.05.05 0 0 0 .016.03c.01.01.023.016.036.017a9.637 9.637 0 0 0 2.134 0a.062.062 0 0 0 .036-.018a.05.05 0 0 0 .016-.03a9.293 9.293 0 0 0 0-2.157a.049.049 0 0 0-.016-.03a.062.062 0 0 0-.036-.017m-2.3-1.49a11.051 11.051 0 0 1 2.466 0c.713.079 1.29.64 1.375 1.363c.098.832.098 1.674 0 2.506a1.559 1.559 0 0 1-1.375 1.364c-.813.09-1.653.09-2.466 0a1.559 1.559 0 0 1-1.376-1.364a10.788 10.788 0 0 1 0-2.506a1.559 1.559 0 0 1 1.376-1.364m9.3 1.491a9.693 9.693 0 0 0-2.134 0a.062.062 0 0 0-.036.018a.049.049 0 0 0-.016.029a9.293 9.293 0 0 0 0 2.158a.05.05 0 0 0 .016.03c.01.01.023.016.036.017a9.637 9.637 0 0 0 2.134 0a.062.062 0 0 0 .036-.018a.05.05 0 0 0 .016-.03a9.293 9.293 0 0 0 0-2.157a.049.049 0 0 0-.016-.03a.062.062 0 0 0-.036-.017m-2.3-1.49a11.051 11.051 0 0 1 2.466 0c.713.079 1.29.64 1.375 1.363c.098.832.098 1.674 0 2.506a1.559 1.559 0 0 1-1.375 1.364c-.813.09-1.653.09-2.466 0a1.559 1.559 0 0 1-1.375-1.364a10.787 10.787 0 0 1 0-2.506a1.559 1.559 0 0 1 1.375-1.364m-11.7 8.491a9.692 9.692 0 0 0-2.134 0a.062.062 0 0 0-.036.018a.049.049 0 0 0-.016.029a9.288 9.288 0 0 0 0 2.158a.05.05 0 0 0 .016.03c.01.01.023.016.036.017a9.637 9.637 0 0 0 2.134 0a.062.062 0 0 0 .036-.018a.05.05 0 0 0 .016-.03a9.288 9.288 0 0 0 0-2.157a.049.049 0 0 0-.016-.03a.062.062 0 0 0-.036-.017m-2.3-1.49a11.19 11.19 0 0 1 2.466 0a1.56 1.56 0 0 1 1.376 1.363c.097.832.097 1.674 0 2.506a1.559 1.559 0 0 1-1.376 1.364c-.813.09-1.653.09-2.466 0a1.559 1.559 0 0 1-1.376-1.364a10.788 10.788 0 0 1 0-2.506a1.559 1.559 0 0 1 1.376-1.364m9.3 1.491a9.693 9.693 0 0 0-2.134 0a.062.062 0 0 0-.036.018a.049.049 0 0 0-.016.029a9.293 9.293 0 0 0 0 2.158a.05.05 0 0 0 .016.03c.01.01.023.016.036.017a9.637 9.637 0 0 0 2.134 0a.062.062 0 0 0 .036-.018a.05.05 0 0 0 .016-.03a9.293 9.293 0 0 0 0-2.157a.049.049 0 0 0-.016-.03a.062.062 0 0 0-.036-.017m-2.3-1.49a11.19 11.19 0 0 1 2.466 0c.713.079 1.29.64 1.375 1.363c.098.832.098 1.674 0 2.506a1.559 1.559 0 0 1-1.375 1.364c-.813.09-1.653.09-2.466 0a1.559 1.559 0 0 1-1.376-1.364a10.788 10.788 0 0 1 0-2.506a1.559 1.559 0 0 1 1.376-1.364m9.3 1.491a9.693 9.693 0 0 0-2.134 0a.062.062 0 0 0-.036.018a.049.049 0 0 0-.016.029a9.293 9.293 0 0 0 0 2.158a.05.05 0 0 0 .016.03c.01.01.023.016.036.017a9.637 9.637 0 0 0 2.134 0a.062.062 0 0 0 .036-.018a.05.05 0 0 0 .016-.03a9.293 9.293 0 0 0 0-2.157a.049.049 0 0 0-.016-.03a.062.062 0 0 0-.036-.017m-2.3-1.49a11.19 11.19 0 0 1 2.466 0c.713.079 1.29.64 1.375 1.363c.098.832.098 1.674 0 2.506a1.559 1.559 0 0 1-1.375 1.364c-.813.09-1.653.09-2.466 0a1.559 1.559 0 0 1-1.375-1.364a10.787 10.787 0 0 1 0-2.506a1.559 1.559 0 0 1 1.375-1.364" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppsSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.206 2.632a10.938 10.938 0 0 0-2.412 0A1.309 1.309 0 0 0 2.64 3.776a10.538 10.538 0 0 0 0 2.448c.07.606.556 1.077 1.154 1.144c.795.09 1.617.09 2.412 0A1.309 1.309 0 0 0 7.36 6.224a10.55 10.55 0 0 0 0-2.448a1.309 1.309 0 0 0-1.154-1.144m7 0a10.937 10.937 0 0 0-2.411 0A1.309 1.309 0 0 0 9.64 3.776a10.538 10.538 0 0 0 0 2.448c.07.606.556 1.077 1.154 1.144c.795.09 1.617.09 2.412 0a1.309 1.309 0 0 0 1.154-1.144a10.55 10.55 0 0 0 0-2.448a1.309 1.309 0 0 0-1.154-1.144m7 0a10.937 10.937 0 0 0-2.411 0a1.309 1.309 0 0 0-1.155 1.144a10.53 10.53 0 0 0 0 2.448c.07.606.556 1.077 1.154 1.144c.795.09 1.617.09 2.412 0a1.309 1.309 0 0 0 1.154-1.144a10.55 10.55 0 0 0 0-2.448a1.309 1.309 0 0 0-1.154-1.144m-14 7a10.938 10.938 0 0 0-2.412 0a1.309 1.309 0 0 0-1.154 1.144a10.537 10.537 0 0 0 0 2.448c.07.606.556 1.077 1.154 1.144c.795.09 1.617.09 2.412 0a1.309 1.309 0 0 0 1.154-1.144a10.55 10.55 0 0 0 0-2.448a1.309 1.309 0 0 0-1.154-1.144m7 0a10.937 10.937 0 0 0-2.411 0a1.309 1.309 0 0 0-1.155 1.144a10.537 10.537 0 0 0 0 2.448c.07.606.556 1.077 1.154 1.144c.795.09 1.617.09 2.412 0a1.309 1.309 0 0 0 1.154-1.144a10.54 10.54 0 0 0 0-2.448a1.309 1.309 0 0 0-1.154-1.144m7 0a10.937 10.937 0 0 0-2.411 0a1.309 1.309 0 0 0-1.155 1.144a10.53 10.53 0 0 0 0 2.448c.07.606.556 1.077 1.154 1.144c.795.09 1.617.09 2.412 0a1.309 1.309 0 0 0 1.154-1.144a10.54 10.54 0 0 0 0-2.448a1.309 1.309 0 0 0-1.154-1.144m-14 7a10.932 10.932 0 0 0-2.412 0a1.309 1.309 0 0 0-1.154 1.144a10.537 10.537 0 0 0 0 2.448c.07.606.556 1.077 1.154 1.144c.795.09 1.617.09 2.412 0a1.309 1.309 0 0 0 1.154-1.144a10.55 10.55 0 0 0 0-2.448a1.309 1.309 0 0 0-1.154-1.144m7 0a10.931 10.931 0 0 0-2.411 0a1.309 1.309 0 0 0-1.155 1.144a10.537 10.537 0 0 0 0 2.448c.07.606.556 1.077 1.154 1.144c.795.09 1.617.09 2.412 0a1.309 1.309 0 0 0 1.154-1.144a10.54 10.54 0 0 0 0-2.448a1.309 1.309 0 0 0-1.154-1.144m7 0a10.931 10.931 0 0 0-2.411 0a1.309 1.309 0 0 0-1.155 1.144a10.53 10.53 0 0 0 0 2.448c.07.606.556 1.077 1.154 1.144c.795.09 1.617.09 2.412 0a1.309 1.309 0 0 0 1.154-1.144a10.54 10.54 0 0 0 0-2.448a1.309 1.309 0 0 0-1.154-1.144"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.75 16.19l2.72-2.72a.75.75 0 1 1 1.06 1.06l-4 4a.75.75 0 0 1-1.06 0l-4-4a.75.75 0 1 1 1.06-1.06l2.72 2.72V6.5a.75.75 0 0 1 1.5 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.75 16.19l2.72-2.72a.75.75 0 1 1 1.06 1.06l-4 4a.75.75 0 0 1-1.06 0l-4-4a.75.75 0 1 1 1.06-1.06l2.72 2.72V6.5a.75.75 0 0 1 1.5 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.03 8.53a.75.75 0 1 0-1.06-1.06l-4 4a.748.748 0 0 0 0 1.06l4 4a.75.75 0 1 0 1.06-1.06l-2.72-2.72H18a.75.75 0 0 0 0-1.5H8.31z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.03 8.53a.75.75 0 1 0-1.06-1.06l-4 4a.748.748 0 0 0 0 1.06l4 4a.75.75 0 1 0 1.06-1.06l-2.72-2.72H18a.75.75 0 0 0 0-1.5H8.31z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.47 8.53a.75.75 0 0 1 1.06-1.06l4 4a.75.75 0 0 1 0 1.06l-4 4a.75.75 0 1 1-1.06-1.06l2.72-2.72H6.5a.75.75 0 0 1 0-1.5h9.69z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.47 8.53a.75.75 0 0 1 1.06-1.06l4 4a.75.75 0 0 1 0 1.06l-4 4a.75.75 0 1 1-1.06-1.06l2.72-2.72H6.5a.75.75 0 0 1 0-1.5h9.69z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.53 10.53a.75.75 0 1 1-1.06-1.06l4-4a.75.75 0 0 1 1.06 0l4 4a.75.75 0 1 1-1.06 1.06l-2.72-2.72v9.69a.75.75 0 0 1-1.5 0V7.81z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.53 10.53a.75.75 0 1 1-1.06-1.06l4-4a.75.75 0 0 1 1.06 0l4 4a.75.75 0 1 1-1.06 1.06l-2.72-2.72v9.69a.75.75 0 0 1-1.5 0V7.81z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AsanaOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.47 3.554a4 4 0 1 1 3.06 7.391a4 4 0 0 1-3.06-7.39M12 4.75a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5m-6.53 7.805a4 4 0 1 1 3.061 7.391a4 4 0 0 1-3.062-7.392M7 13.75a2.5 2.5 0 1 0 0 5.001a2.5 2.5 0 0 0 0-5.001m10-1.5a4 4 0 1 0 0 8a4 4 0 0 0 0-8m-.957 1.69a2.5 2.5 0 1 1 1.914 4.62a2.5 2.5 0 0 1-1.914-4.62" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AsanaSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3.5a3.75 3.75 0 1 0 0 7.5a3.75 3.75 0 0 0 0-7.5m-5 9a3.75 3.75 0 1 0 0 7.498A3.75 3.75 0 0 0 7 12.5m10 0a3.75 3.75 0 1 0 0 7.499a3.75 3.75 0 0 0 0-7.499"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AtSignOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.755 12a7.25 7.25 0 1 1 14.5 0c0 .034.002.067.006.1a.757.757 0 0 0-.004.095c.007.255-.026.65-.11 1.027c-.09.411-.21.647-.279.72a1.65 1.65 0 0 1-2.333.056c-.207-.197-.266-.487-.282-1.16l-.116-4.798a.75.75 0 1 0-1.5.036l.017.693a4.179 4.179 0 1 0 .532 5.94c.087.13.19.257.315.375a3.15 3.15 0 0 0 4.453-.107c.37-.388.557-.975.658-1.43a6.13 6.13 0 0 0 .145-1.388a.777.777 0 0 0-.006-.082a.763.763 0 0 0 .004-.077a8.75 8.75 0 1 0-3.337 6.875a.75.75 0 0 0-.928-1.178A7.25 7.25 0 0 1 4.755 12m7.25-2.679a2.679 2.679 0 1 0 0 5.358a2.679 2.679 0 0 0 0-5.358" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AtSignSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.755 12a7.25 7.25 0 1 1 14.5 0c0 .034.002.067.006.1a.757.757 0 0 0-.004.095c.007.255-.026.65-.11 1.027c-.09.411-.21.647-.279.72a1.65 1.65 0 0 1-2.333.056c-.207-.197-.266-.487-.282-1.16l-.116-4.798a.75.75 0 1 0-1.5.036l.017.693a4.179 4.179 0 1 0 .532 5.94c.087.13.19.257.315.375a3.15 3.15 0 0 0 4.453-.107c.37-.388.557-.975.658-1.43a6.13 6.13 0 0 0 .145-1.388a.777.777 0 0 0-.006-.082a.763.763 0 0 0 .004-.077a8.75 8.75 0 1 0-3.337 6.875a.75.75 0 0 0-.928-1.178A7.25 7.25 0 0 1 4.755 12m7.25-2.679a2.679 2.679 0 1 0 0 5.358a2.679 2.679 0 0 0 0-5.358" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AttachOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 12A6.75 6.75 0 0 1 9 5.25h8a4.75 4.75 0 1 1 0 9.5H8.857a2.75 2.75 0 1 1 0-5.5h7.857a.75.75 0 0 1 0 1.5H8.857a1.25 1.25 0 1 0 0 2.5H17a3.25 3.25 0 0 0 0-6.5H9a5.25 5.25 0 1 0 0 10.5h7.714a.75.75 0 0 1 0 1.5H9A6.75 6.75 0 0 1 2.25 12" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AttachSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 12A6.75 6.75 0 0 1 9 5.25h8a4.75 4.75 0 1 1 0 9.5H8.857a2.75 2.75 0 1 1 0-5.5h7.857a.75.75 0 0 1 0 1.5H8.857a1.25 1.25 0 1 0 0 2.5H17a3.25 3.25 0 0 0 0-6.5H9a5.25 5.25 0 1 0 0 10.5h7.714a.75.75 0 0 1 0 1.5H9A6.75 6.75 0 0 1 2.25 12" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AwardOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M12 6.25a3.25 3.25 0 1 0 0 6.5a3.25 3.25 0 0 0 0-6.5M10.25 9.5a1.75 1.75 0 1 1 3.5 0a1.75 1.75 0 0 1-3.5 0"/><path d="M12 2.25a7.25 7.25 0 0 0-6.063 11.226l-2.587 4.48a.75.75 0 0 0 .795 1.11l2.614-.514l.861 2.52a.75.75 0 0 0 1.36.133l2.58-4.468a7.313 7.313 0 0 0 .88 0l2.58 4.468a.75.75 0 0 0 1.36-.134l.858-2.526l2.616.52a.75.75 0 0 0 .796-1.11l-2.586-4.479A7.25 7.25 0 0 0 12 2.25M6.25 9.5a5.75 5.75 0 1 1 11.5 0a5.75 5.75 0 0 1-11.5 0m3.734 6.966a7.243 7.243 0 0 1-3.027-1.757l-1.482 2.567l1.637-.322a.75.75 0 0 1 .854.493l.54 1.579zm5.508 2.556l-1.476-2.556a7.244 7.244 0 0 0 3.027-1.757l1.48 2.563l-1.638-.326a.75.75 0 0 0-.856.495z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AwardSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.25 9.5a1.75 1.75 0 1 1 3.5 0a1.75 1.75 0 0 1-3.5 0"/><path fill="currentColor" fill-rule="evenodd" d="M5 9.5a7 7 0 1 1 12.923 3.733l2.727 4.722a.75.75 0 0 1-.796 1.11l-2.616-.52l-.858 2.526a.75.75 0 0 1-1.36.134l-2.72-4.711a7.119 7.119 0 0 1-.6 0l-2.72 4.711a.75.75 0 0 1-1.36-.132l-.861-2.52l-2.614.513a.75.75 0 0 1-.795-1.11l2.727-4.723A6.967 6.967 0 0 1 5 9.5m2.086 4.985a6.993 6.993 0 0 0 3.027 1.758l-1.607 2.783l-.54-1.579a.75.75 0 0 0-.854-.493l-1.637.322zm6.801 1.758l1.605 2.779l.537-1.581a.75.75 0 0 1 .856-.495l1.638.326l-1.609-2.787a6.994 6.994 0 0 1-3.027 1.758M12 6.25a3.25 3.25 0 1 0 0 6.5a3.25 3.25 0 0 0 0-6.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BackspaceOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.47 8.47a.75.75 0 0 1 1.06 0L14 10.94l2.47-2.47a.75.75 0 1 1 1.06 1.06L15.06 12l2.47 2.47a.75.75 0 0 1-1.06 1.06L14 13.06l-2.47 2.47a.75.75 0 1 1-1.06-1.06L12.94 12l-2.47-2.47a.75.75 0 0 1 0-1.06"/><path fill="currentColor" fill-rule="evenodd" d="M8.406 4.95a3.75 3.75 0 0 0-2.95 1.434l-3.56 4.535a1.75 1.75 0 0 0 0 2.162l3.56 4.535a3.75 3.75 0 0 0 2.95 1.434h10.647a2.75 2.75 0 0 0 2.75-2.75V7.7a2.75 2.75 0 0 0-2.75-2.75zm-1.77 2.36a2.25 2.25 0 0 1 1.77-.86h10.647c.69 0 1.25.56 1.25 1.25v8.6c0 .69-.56 1.25-1.25 1.25H8.406a2.25 2.25 0 0 1-1.77-.86l-3.561-4.536a.25.25 0 0 1 0-.309z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BackspaceSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.406 5.2a3.5 3.5 0 0 0-2.753 1.338l-3.561 4.535a1.5 1.5 0 0 0 0 1.853l3.561 4.536A3.5 3.5 0 0 0 8.406 18.8h10.647a2.5 2.5 0 0 0 2.5-2.5V7.7a2.5 2.5 0 0 0-2.5-2.5zm2.064 3.27a.75.75 0 0 1 1.06 0L14 10.94l2.47-2.47a.75.75 0 1 1 1.06 1.06L15.06 12l2.47 2.47a.75.75 0 0 1-1.06 1.06L14 13.06l-2.47 2.47a.75.75 0 1 1-1.06-1.06L12.94 12l-2.47-2.47a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BagOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.25 6.631v-1.17a1.75 1.75 0 0 1 1.49-1.73l1.22-.183a13.75 13.75 0 0 1 4.08 0l1.22.183a1.75 1.75 0 0 1 1.49 1.73v1.17l1.714.138a2.86 2.86 0 0 1 2.593 2.394a27.128 27.128 0 0 1 0 8.674a2.86 2.86 0 0 1-2.593 2.394l-1.872.15a57.078 57.078 0 0 1-9.184 0l-1.872-.15a2.86 2.86 0 0 1-2.593-2.394a27.129 27.129 0 0 1 0-8.674A2.86 2.86 0 0 1 5.536 6.77zm2.933-1.6a12.25 12.25 0 0 1 3.634 0l1.22.183a.25.25 0 0 1 .213.247v1.065a57.078 57.078 0 0 0-6.5 0V5.46a.25.25 0 0 1 .213-.247zM7.529 8.113c2.976-.24 5.966-.24 8.942 0l1.872.152a1.36 1.36 0 0 1 1.234 1.138c.062.385.115.77.16 1.158a17.517 17.517 0 0 1-15.474 0c.044-.387.098-.773.16-1.158a1.36 1.36 0 0 1 1.234-1.138zm-3.4 4.044a19.018 19.018 0 0 0 15.742 0a25.628 25.628 0 0 1-.294 5.44a1.36 1.36 0 0 1-1.234 1.139l-1.872.15c-2.976.24-5.966.24-8.942 0l-1.872-.15a1.36 1.36 0 0 1-1.234-1.139c-.291-1.8-.39-3.624-.294-5.44" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BagSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.25 5.461v1.42l-1.694.138a2.61 2.61 0 0 0-2.367 2.184c-.041.258-.08.516-.114.775a.298.298 0 0 0 .169.308l.077.036c5.429 2.57 11.93 2.57 17.358 0l.077-.036a.298.298 0 0 0 .168-.308a26.748 26.748 0 0 0-.113-.775a2.61 2.61 0 0 0-2.367-2.184l-1.694-.137v-1.42a1.75 1.75 0 0 0-1.49-1.731l-1.22-.183a13.75 13.75 0 0 0-4.08 0l-1.22.183a1.75 1.75 0 0 0-1.49 1.73m6.567-.43a12.25 12.25 0 0 0-3.634 0l-1.22.183a.25.25 0 0 0-.213.247v1.315a56.826 56.826 0 0 1 6.5 0V5.461a.25.25 0 0 0-.213-.247z" clip-rule="evenodd"/><path fill="currentColor" d="M21.118 12.07a.2.2 0 0 0-.282-.17c-5.571 2.467-12.101 2.467-17.672 0a.2.2 0 0 0-.282.17a26.88 26.88 0 0 0 .307 5.727a2.61 2.61 0 0 0 2.367 2.185l1.872.15c3.043.246 6.1.246 9.144 0l1.872-.15a2.61 2.61 0 0 0 2.367-2.185c.306-1.895.41-3.815.307-5.726"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BankOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.415 2.395a1.99 1.99 0 0 1 1.17 0l2.986.918a16.722 16.722 0 0 1 4.39 2.089c1.054.705.555 2.348-.713 2.348H4.752c-1.268 0-1.767-1.643-.714-2.348a16.721 16.721 0 0 1 4.391-2.09zm.73 1.434a.49.49 0 0 0-.29 0l-2.985.918A15.22 15.22 0 0 0 5.5 6.25h13a15.22 15.22 0 0 0-3.37-1.503z" clip-rule="evenodd"/><path fill="currentColor" d="M4.25 21a.75.75 0 0 1 .75-.75h14a.75.75 0 0 1 0 1.5H5a.75.75 0 0 1-.75-.75m2-4a.75.75 0 0 0 1.5 0v-6a.75.75 0 0 0-1.5 0zm5.75.75a.75.75 0 0 1-.75-.75v-6a.75.75 0 0 1 1.5 0v6a.75.75 0 0 1-.75.75m4.25-.75a.75.75 0 0 0 1.5 0v-6a.75.75 0 0 0-1.5 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BankSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.512 2.634a1.74 1.74 0 0 0-1.023 0l-2.986.918A16.471 16.471 0 0 0 4.178 5.61c-.848.567-.446 1.89.574 1.89h14.496c1.02 0 1.422-1.323.575-1.89a16.472 16.472 0 0 0-4.326-2.058zM4.25 21a.75.75 0 0 1 .75-.75h14a.75.75 0 0 1 0 1.5H5a.75.75 0 0 1-.75-.75m2-4a.75.75 0 0 0 1.5 0v-6a.75.75 0 0 0-1.5 0zm5.75.75a.75.75 0 0 1-.75-.75v-6a.75.75 0 0 1 1.5 0v6a.75.75 0 0 1-.75.75m4.25-.75a.75.75 0 0 0 1.5 0v-6a.75.75 0 0 0-1.5 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryEmptyOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M20.75 14v-4a.25.25 0 0 0-.25-.25h-1.4a.75.75 0 0 1-.75-.75v-.81a.934.934 0 0 0-.862-.932L16.22 7.16a67.146 67.146 0 0 0-10.44 0l-1.347.105a1.143 1.143 0 0 0-1.038.95a22.899 22.899 0 0 0 0 7.572c.087.517.515.909 1.038.95l1.347.105c3.475.27 6.965.27 10.44 0l1.268-.1a.934.934 0 0 0 .862-.93V15a.75.75 0 0 1 .75-.75h1.4a.25.25 0 0 0 .25-.25m1.5-4v4a1.75 1.75 0 0 1-1.75 1.75h-.65v.06a2.434 2.434 0 0 1-2.245 2.427l-1.269.1c-3.552.276-7.12.276-10.672 0l-1.348-.106a2.64 2.64 0 0 1-2.4-2.197a24.398 24.398 0 0 1 0-8.068a2.642 2.642 0 0 1 2.4-2.197l1.348-.105a68.646 68.646 0 0 1 10.672 0l1.269.099A2.434 2.434 0 0 1 19.85 8.19v.06h.65c.966 0 1.75.783 1.75 1.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFullOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.75 14.208V9.79c0-.604-.464-1.107-1.066-1.154a61.932 61.932 0 0 0-9.789 0l-.188.015a1.004 1.004 0 0 0-.9.78a11.686 11.686 0 0 0 0 5.12c.095.429.459.745.896.781l.175.015c3.269.27 6.554.27 9.823 0a1.143 1.143 0 0 0 1.049-1.14"/><path fill="currentColor" fill-rule="evenodd" d="M22.25 14v-4a1.75 1.75 0 0 0-1.75-1.75h-.65v-.06a2.434 2.434 0 0 0-2.245-2.427l-1.269-.1a68.646 68.646 0 0 0-10.672 0l-1.348.106c-1.21.094-2.2 1-2.4 2.197a24.398 24.398 0 0 0 0 8.068a2.642 2.642 0 0 0 2.4 2.197l1.348.105a68.63 68.63 0 0 0 10.672 0l1.269-.099a2.434 2.434 0 0 0 2.245-2.427v-.06h.65A1.75 1.75 0 0 0 22.25 14m-1.5-4v4a.25.25 0 0 1-.25.25h-1.4a.75.75 0 0 0-.75.75v.81a.934.934 0 0 1-.862.932l-1.268.099c-3.475.27-6.965.27-10.44 0l-1.347-.105a1.142 1.142 0 0 1-1.038-.95a22.898 22.898 0 0 1 0-7.572a1.142 1.142 0 0 1 1.038-.95L5.78 7.16c3.475-.27 6.965-.27 10.44 0l1.268.1a.934.934 0 0 1 .862.93V9c0 .414.336.75.75.75h1.4a.25.25 0 0 1 .25.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFullSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M22 10v4a1.5 1.5 0 0 1-1.5 1.5h-.9v.31c0 1.14-.878 2.09-2.015 2.178l-1.268.099a68.39 68.39 0 0 1-10.634 0l-1.348-.105a2.392 2.392 0 0 1-2.173-1.99a24.149 24.149 0 0 1 0-7.985a2.392 2.392 0 0 1 2.173-1.989l1.348-.105a68.397 68.397 0 0 1 10.634 0l1.268.099A2.184 2.184 0 0 1 19.6 8.19v.31h.9A1.5 1.5 0 0 1 22 10m-5.25 4.208V9.79c0-.604-.464-1.107-1.066-1.154a61.932 61.932 0 0 0-9.789 0l-.188.015a1.004 1.004 0 0 0-.9.78a11.686 11.686 0 0 0 0 5.12c.095.429.459.745.896.781l.175.015c3.269.27 6.554.27 9.823 0a1.143 1.143 0 0 0 1.049-1.14" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryLowOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M6.51 8.829a.75.75 0 0 1 .24.55v5.225a.75.75 0 0 1-.81.748l-.236-.02a1.004 1.004 0 0 1-.898-.78a11.686 11.686 0 0 1 0-5.12c.097-.43.462-.747.9-.781m.803.178c-.153-.142-.594-.195-.802-.178z"/><path d="M20.75 14v-4a.25.25 0 0 0-.25-.25h-1.4a.75.75 0 0 1-.75-.75v-.81a.934.934 0 0 0-.862-.932L16.22 7.16a67.146 67.146 0 0 0-10.44 0l-1.347.105a1.143 1.143 0 0 0-1.038.95a22.899 22.899 0 0 0 0 7.572c.087.517.515.909 1.038.95l1.347.105c3.475.27 6.965.27 10.44 0l1.268-.1a.934.934 0 0 0 .862-.93V15a.75.75 0 0 1 .75-.75h1.4a.25.25 0 0 0 .25-.25m1.5-4v4a1.75 1.75 0 0 1-1.75 1.75h-.65v.06a2.434 2.434 0 0 1-2.245 2.427l-1.268.1a68.892 68.892 0 0 1-10.673 0l-1.348-.106a2.64 2.64 0 0 1-2.4-2.197a24.398 24.398 0 0 1 0-8.068a2.642 2.642 0 0 1 2.4-2.197l1.348-.105a68.647 68.647 0 0 1 10.673 0l1.268.099A2.434 2.434 0 0 1 19.85 8.19v.06h.65c.967 0 1.75.783 1.75 1.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryLowSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M22 10v4a1.5 1.5 0 0 1-1.5 1.5h-.9v.31c0 1.14-.878 2.09-2.015 2.178l-1.268.099a68.39 68.39 0 0 1-10.634 0l-1.348-.105a2.392 2.392 0 0 1-2.173-1.99a24.149 24.149 0 0 1 0-7.985a2.392 2.392 0 0 1 2.173-1.989l1.348-.105a68.397 68.397 0 0 1 10.634 0l1.268.099A2.184 2.184 0 0 1 19.6 8.19v.31h.9A1.5 1.5 0 0 1 22 10M6.75 9.38a.75.75 0 0 0-.809-.748l-.234.019a1.004 1.004 0 0 0-.9.78a11.686 11.686 0 0 0 0 5.12c.095.429.459.745.896.781l.236.02a.75.75 0 0 0 .811-.748z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryMostOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.067 8.484a61.932 61.932 0 0 0-7.172.152l-.188.015a1.004 1.004 0 0 0-.9.78a11.686 11.686 0 0 0 0 5.12c.095.429.459.745.896.781l.175.015c2.392.197 4.793.25 7.19.159a.75.75 0 0 0 .722-.75V9.233a.75.75 0 0 0-.723-.75"/><path fill="currentColor" fill-rule="evenodd" d="M22.25 14v-4a1.75 1.75 0 0 0-1.75-1.75h-.65v-.06a2.434 2.434 0 0 0-2.245-2.427l-1.269-.1a68.646 68.646 0 0 0-10.672 0l-1.348.106c-1.21.094-2.2 1-2.4 2.197a24.398 24.398 0 0 0 0 8.068a2.642 2.642 0 0 0 2.4 2.197l1.348.105a68.63 68.63 0 0 0 10.672 0l1.269-.099a2.434 2.434 0 0 0 2.245-2.427v-.06h.65A1.75 1.75 0 0 0 22.25 14m-1.5-4v4a.25.25 0 0 1-.25.25h-1.4a.75.75 0 0 0-.75.75v.81a.934.934 0 0 1-.862.932l-1.268.099c-3.475.27-6.965.27-10.44 0l-1.347-.105a1.142 1.142 0 0 1-1.038-.95a22.898 22.898 0 0 1 0-7.572a1.142 1.142 0 0 1 1.038-.95L5.78 7.16c3.475-.27 6.965-.27 10.44 0l1.268.1a.934.934 0 0 1 .862.93V9c0 .414.336.75.75.75h1.4a.25.25 0 0 1 .25.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryMostSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M22 10v4a1.5 1.5 0 0 1-1.5 1.5h-.9v.31c0 1.14-.878 2.09-2.015 2.178l-1.268.099a68.39 68.39 0 0 1-10.634 0l-1.348-.105a2.392 2.392 0 0 1-2.173-1.99a24.149 24.149 0 0 1 0-7.985a2.392 2.392 0 0 1 2.173-1.989l1.348-.105a68.397 68.397 0 0 1 10.634 0l1.268.099A2.184 2.184 0 0 1 19.6 8.19v.31h.9A1.5 1.5 0 0 1 22 10m-8.933-1.516a61.933 61.933 0 0 0-7.172.152l-.188.015a1.004 1.004 0 0 0-.9.78a11.686 11.686 0 0 0 0 5.12c.095.429.459.745.896.781l.175.015c2.392.197 4.793.25 7.19.159a.75.75 0 0 0 .722-.75V9.233a.75.75 0 0 0-.723-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryQuarterOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.75 9.256a.75.75 0 0 0-.784-.75c-.69.032-1.381.075-2.07.13l-.19.015a1.004 1.004 0 0 0-.9.78a11.686 11.686 0 0 0 0 5.12c.096.429.46.745.898.781l.174.015c.695.057 1.39.103 2.087.136a.75.75 0 0 0 .785-.75z"/><path fill="currentColor" fill-rule="evenodd" d="M22.25 14v-4a1.75 1.75 0 0 0-1.75-1.75h-.65v-.06a2.434 2.434 0 0 0-2.245-2.427l-1.268-.1a68.647 68.647 0 0 0-10.673 0l-1.348.106c-1.21.094-2.2 1-2.4 2.197a24.398 24.398 0 0 0 0 8.068a2.642 2.642 0 0 0 2.4 2.197l1.348.105c3.552.277 7.12.277 10.673 0l1.268-.099a2.434 2.434 0 0 0 2.245-2.427v-.06h.65A1.75 1.75 0 0 0 22.25 14m-1.5-4v4a.25.25 0 0 1-.25.25h-1.4a.75.75 0 0 0-.75.75v.81a.934.934 0 0 1-.862.932l-1.268.099c-3.475.27-6.965.27-10.44 0l-1.347-.105a1.142 1.142 0 0 1-1.038-.95a22.899 22.899 0 0 1 0-7.572a1.142 1.142 0 0 1 1.038-.95L5.78 7.16c3.475-.27 6.965-.27 10.44 0l1.268.1a.934.934 0 0 1 .862.93V9c0 .414.336.75.75.75h1.4a.25.25 0 0 1 .25.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryQuarterSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M22 10v4a1.5 1.5 0 0 1-1.5 1.5h-.9v.31c0 1.14-.878 2.09-2.015 2.178l-1.268.099a68.39 68.39 0 0 1-10.634 0l-1.348-.105a2.392 2.392 0 0 1-2.173-1.99a24.149 24.149 0 0 1 0-7.985a2.392 2.392 0 0 1 2.173-1.989l1.348-.105a68.397 68.397 0 0 1 10.634 0l1.268.099A2.184 2.184 0 0 1 19.6 8.19v.31h.9A1.5 1.5 0 0 1 22 10M8.75 9.256a.75.75 0 0 0-.784-.75c-.69.032-1.381.075-2.07.13l-.19.015a1.004 1.004 0 0 0-.9.78a11.686 11.686 0 0 0 0 5.12c.096.429.46.745.898.781l.174.015c.695.057 1.39.103 2.087.136a.75.75 0 0 0 .785-.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BattryHalfOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.79 9.197a.75.75 0 0 0-.76-.75c-1.38.016-2.758.08-4.135.189l-.188.015a1.004 1.004 0 0 0-.9.78a11.686 11.686 0 0 0 0 5.12c.095.429.459.745.896.781l.175.015c1.382.114 2.767.18 4.152.197a.75.75 0 0 0 .76-.75z"/><path fill="currentColor" fill-rule="evenodd" d="M22.25 14v-4a1.75 1.75 0 0 0-1.75-1.75h-.65v-.06a2.434 2.434 0 0 0-2.245-2.427l-1.269-.1a68.646 68.646 0 0 0-10.672 0l-1.348.106c-1.21.094-2.2 1-2.4 2.197a24.398 24.398 0 0 0 0 8.068a2.642 2.642 0 0 0 2.4 2.197l1.348.105a68.63 68.63 0 0 0 10.672 0l1.269-.099a2.434 2.434 0 0 0 2.245-2.427v-.06h.65A1.75 1.75 0 0 0 22.25 14m-1.5-4v4a.25.25 0 0 1-.25.25h-1.4a.75.75 0 0 0-.75.75v.81a.934.934 0 0 1-.862.932l-1.268.099c-3.475.27-6.965.27-10.44 0l-1.347-.105a1.142 1.142 0 0 1-1.038-.95a22.898 22.898 0 0 1 0-7.572a1.142 1.142 0 0 1 1.038-.95L5.78 7.16c3.475-.27 6.965-.27 10.44 0l1.268.1a.934.934 0 0 1 .862.93V9c0 .414.336.75.75.75h1.4a.25.25 0 0 1 .25.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BattryHalfSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M22 10v4a1.5 1.5 0 0 1-1.5 1.5h-.9v.31c0 1.14-.878 2.09-2.015 2.178l-1.268.099a68.39 68.39 0 0 1-10.634 0l-1.348-.105a2.392 2.392 0 0 1-2.173-1.99a24.149 24.149 0 0 1 0-7.985a2.392 2.392 0 0 1 2.173-1.989l1.348-.105a68.397 68.397 0 0 1 10.634 0l1.268.099A2.184 2.184 0 0 1 19.6 8.19v.31h.9A1.5 1.5 0 0 1 22 10m-11.21-.803a.75.75 0 0 0-.76-.75c-1.38.016-2.758.08-4.135.189l-.188.015a1.004 1.004 0 0 0-.9.78a11.686 11.686 0 0 0 0 5.12c.095.429.459.745.896.781l.175.015c1.382.114 2.767.18 4.152.197a.75.75 0 0 0 .76-.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BehanceOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.762 8.39a.2.2 0 0 0-.2.2v2.375c0 .11.09.2.2.2h2.754a1.388 1.388 0 0 0 0-2.776zm0 4.555a.2.2 0 0 0-.2.2v2.874c0 .11.09.2.2.2h2.932a1.637 1.637 0 0 0 0-3.274zm11.375.142a.088.088 0 0 1-.088-.088c0-.974.789-1.763 1.762-1.763h.39c.973 0 1.762.79 1.762 1.763c0 .048-.04.088-.088.088z"/><path fill="currentColor" fill-rule="evenodd" d="M12.57 11.993a4.352 4.352 0 0 0-.2-.234c.433-.683.653-1.48.632-2.287a4.616 4.616 0 0 0-.407-2.019a3.782 3.782 0 0 0-1.289-1.567l-.011-.008a4.584 4.584 0 0 0-1.771-.73a8.676 8.676 0 0 0-1.855-.18H2a1.5 1.5 0 0 0-1.5 1.5v11.957a1.5 1.5 0 0 0 1.5 1.5h5.8c.658 0 1.312-.083 1.948-.248a5.829 5.829 0 0 0 1.83-.831l.007-.005a4.476 4.476 0 0 0 1.35-1.48a5.729 5.729 0 0 0 .808 1.169c.538.58 1.2 1.03 1.936 1.322l.018.007a6.46 6.46 0 0 0 2.367.42a6.047 6.047 0 0 0 3.303-.883a5.49 5.49 0 0 0 2.218-3.12a1.5 1.5 0 0 0 .008-.741c.112-.193.183-.413.2-.649a7.612 7.612 0 0 0-.201-2.36l-.002-.006a6.307 6.307 0 0 0-.998-2.151l-.015-.022a5.463 5.463 0 0 0-1.891-1.596a5.667 5.667 0 0 0-2.57-.58a6.153 6.153 0 0 0-2.372.464l-.01.004a5.79 5.79 0 0 0-1.886 1.28l-.012.011a5.668 5.668 0 0 0-1.228 1.96l-.001.002a6.105 6.105 0 0 0-.037.101m-2.742-.07a2.854 2.854 0 0 1 1.664 1.132a3.296 3.296 0 0 1 .542 1.9a3.24 3.24 0 0 1-.352 1.57c-.223.433-.55.805-.95 1.083a4.331 4.331 0 0 1-1.361.617a6.256 6.256 0 0 1-1.57.2H2V6.468h5.678a7.176 7.176 0 0 1 1.56.152c.433.077.845.246 1.208.495c.342.24.612.568.78.95c.2.441.294.923.275 1.407a2.584 2.584 0 0 1-.409 1.502c-.312.42-.78.749-1.264.95m6.174 3.07a2.318 2.318 0 0 1 0-.214h6.295a6.052 6.052 0 0 0-.162-1.895a4.807 4.807 0 0 0-.76-1.64a3.961 3.961 0 0 0-1.351-1.146a4.166 4.166 0 0 0-1.902-.427a4.655 4.655 0 0 0-1.806.35c-.525.221-1 .543-1.398.948a4.168 4.168 0 0 0-.904 1.441a5.105 5.105 0 0 0-.313 1.81c-.01.626.093 1.249.304 1.839c.18.534.464 1.027.837 1.45c.386.415.86.739 1.388.947a4.96 4.96 0 0 0 1.845.323a4.547 4.547 0 0 0 2.51-.664a3.994 3.994 0 0 0 1.55-2.227h-2.102a1.457 1.457 0 0 1-.589.72a2.125 2.125 0 0 1-1.236.351a2.128 2.128 0 0 1-1.617-.512a2.353 2.353 0 0 1-.59-1.453" clip-rule="evenodd"/><path fill="currentColor" d="M20.447 6.396a.4.4 0 0 0-.4-.396h-4.025a.4.4 0 0 0-.4.4v1.07a.4.4 0 0 0 .4.4h4.036a.4.4 0 0 0 .4-.404z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BehanceSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.187 11.362a3.243 3.243 0 0 0 1.343-1.108a3.15 3.15 0 0 0 .451-1.751a3.813 3.813 0 0 0-.304-1.64a2.634 2.634 0 0 0-.86-1.109c-.4-.29-.855-.487-1.333-.576A7.504 7.504 0 0 0 7.764 5H1.8a.3.3 0 0 0-.3.3v13.341a.3.3 0 0 0 .3.3h6.1c.584 0 1.166-.078 1.73-.233a4.658 4.658 0 0 0 1.501-.72a3.42 3.42 0 0 0 1.05-1.263a3.956 3.956 0 0 0 .388-1.829a4.006 4.006 0 0 0-.599-2.216a3.148 3.148 0 0 0-1.836-1.32zM4.55 7.25a.2.2 0 0 0-.2.2v3.1c0 .11.09.2.2.2h2.8a1.75 1.75 0 0 0 0-3.5zm-.2 5.45c0-.11.09-.2.2-.2h2.925a2.125 2.125 0 0 1 0 4.25H4.55a.2.2 0 0 1-.2-.2z" clip-rule="evenodd"/><path fill="currentColor" d="M15.439 5.71h4.723a.3.3 0 0 1 .3.296l.013.962a.3.3 0 0 1-.3.304h-4.736a.3.3 0 0 1-.3-.3V6.01a.3.3 0 0 1 .3-.3"/><path fill="currentColor" fill-rule="evenodd" d="M17.972 16.87a2.25 2.25 0 0 0 1.363-.41c.26-.173.467-.42.598-.712a.215.215 0 0 1 .194-.13h1.921c.13 0 .226.122.187.246a4.658 4.658 0 0 1-1.64 2.358a4.802 4.802 0 0 1-2.77.775a5.197 5.197 0 0 1-2.036-.376a4.196 4.196 0 0 1-1.531-1.109a5.016 5.016 0 0 1-.924-1.695a6.43 6.43 0 0 1-.335-2.15a6.286 6.286 0 0 1 .346-2.117a4.93 4.93 0 0 1 .996-1.684a4.73 4.73 0 0 1 1.543-1.109a4.883 4.883 0 0 1 1.993-.41a4.383 4.383 0 0 1 2.099.5a4.45 4.45 0 0 1 1.49 1.34c.391.578.676 1.228.839 1.917a7.56 7.56 0 0 1 .178 2.217h-6.945a2.853 2.853 0 0 0 .65 1.95c.242.224.526.393.833.496c.307.103.63.138.95.103M15.6 12.399a.1.1 0 0 0 .101.101h4.048a.101.101 0 0 0 .101-.101a2.024 2.024 0 0 0-2.024-2.024h-.202a2.024 2.024 0 0 0-2.024 2.024" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BinocularOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m1.623 12.342l2.273-6.437a.751.751 0 0 1 .04-.093c.38-.738 1.062-1.303 1.79-1.68c.732-.379 1.583-.604 2.385-.604c1.148 0 2.29.46 3.147 1.192c.275.234.527.502.742.798a4.61 4.61 0 0 1 .742-.798c.857-.731 1.999-1.192 3.147-1.192c.802 0 1.653.225 2.386.604c.727.377 1.409.942 1.789 1.68c.015.03.029.061.04.093l2.273 6.436a5.194 5.194 0 1 1-9.962 2.686h-.83A5.195 5.195 0 0 1 1.25 14.28M4.182 9.6l1.11-3.142c.201-.363.594-.72 1.123-.993c.542-.281 1.158-.437 1.696-.437c.761 0 1.564.313 2.174.833c.608.519.965 1.189.965 1.874v4.567a5.196 5.196 0 0 0-7.068-2.701m5.882 3.933a3.694 3.694 0 1 1-7.239 1.489a3.694 3.694 0 0 1 7.239-1.489m2.686-1.23c.194-.473.456-.91.774-1.3a5.184 5.184 0 0 1 4.1-1.92a5.191 5.191 0 0 1 2.194.518l-1.11-3.143c-.201-.363-.594-.72-1.122-.993c-.543-.281-1.16-.437-1.697-.437c-.761 0-1.564.313-2.174.833c-.608.519-.965 1.189-.965 1.874zm1.111 1.975a3.695 3.695 0 0 0 7.389.002v-.002a4 4 0 0 0-.005-.202a3.696 3.696 0 0 0-3.69-3.493a3.69 3.69 0 0 0-3.694 3.695" clip-rule="evenodd"/><path fill="currentColor" d="M1.623 12.342a5.18 5.18 0 0 0-.373 1.936"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BinocularSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.096 4.91c.354.302.665.66.904 1.061c.239-.401.55-.759.904-1.061c.816-.696 1.901-1.132 2.985-1.132c.758 0 1.57.214 2.27.576c.698.361 1.334.895 1.682 1.572a.525.525 0 0 1 .023.051l2.13 5.6a5.194 5.194 0 1 1-9.609 3.2h-.77a5.195 5.195 0 1 1-9.609-3.2l2.13-5.6a.504.504 0 0 1 .023-.05c.349-.678.984-1.212 1.681-1.573c.701-.362 1.513-.576 2.271-.576c1.084 0 2.169.436 2.985 1.132M2.75 14.278a3.694 3.694 0 1 1 7.389 0a3.694 3.694 0 0 1-7.389 0m11.111 0a3.694 3.694 0 1 1 7.389 0a3.694 3.694 0 0 1-7.389 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BluetoothOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.02 6.48a.75.75 0 0 1 1.06-1.06l4.67 4.669V2.604c0-.891 1.077-1.338 1.707-.708l4.346 4.347a1 1 0 0 1 0 1.414L13.51 11.95l4.293 4.293a1 1 0 0 1 0 1.414l-4.346 4.346c-.63.63-1.707.184-1.707-.707v-7.485l-4.67 4.67a.75.75 0 1 1-1.06-1.062l5.47-5.47zm7.23-2.67v6.279l3.14-3.14zm3.14 13.14l-3.14-3.14v6.279z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BluetoothSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.02 6.48a.75.75 0 0 1 1.06-1.06l4.67 4.669V2.604c0-.891 1.077-1.338 1.707-.708l4.346 4.347a1 1 0 0 1 0 1.414L13.51 11.95l4.293 4.293a1 1 0 0 1 0 1.414l-4.346 4.346c-.63.63-1.707.184-1.707-.707v-7.485l-4.67 4.67a.75.75 0 1 1-1.06-1.062l5.47-5.47zm7.23-2.67v6.279l3.14-3.14zm3.14 13.14l-3.14-3.14v6.279z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookCheckOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.541 7.48a.75.75 0 0 1-.022 1.061l-3.125 3a.75.75 0 0 1-1.038 0l-1.875-1.8a.75.75 0 1 1 1.038-1.082l1.356 1.301l2.606-2.501a.75.75 0 0 1 1.06.022"/><path fill="currentColor" fill-rule="evenodd" d="M3.75 8A4.75 4.75 0 0 1 8.5 3.25h10c.966 0 1.75.784 1.75 1.75v15a1.75 1.75 0 0 1-1.75 1.75h-11A3.75 3.75 0 0 1 3.75 18zm1.5 7a3.734 3.734 0 0 1 2.25-.75h11.25V5a.25.25 0 0 0-.25-.25h-10A3.25 3.25 0 0 0 5.25 8zm0 3a2.25 2.25 0 0 0 2.25 2.25h11a.25.25 0 0 0 .25-.25v-4.25H7.5A2.25 2.25 0 0 0 5.25 18" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookCheckSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.5 3.5A4.5 4.5 0 0 0 4 8v10h.035A3.5 3.5 0 0 0 7.5 21H19a1 1 0 0 0 1-1V5a1.5 1.5 0 0 0-1.5-1.5zm-1 12h11v4h-11a2 2 0 1 1 0-4m8.02-6.959a.75.75 0 1 0-1.04-1.082L11.876 9.96l-1.357-1.3a.75.75 0 1 0-1.038 1.082l1.875 1.8a.75.75 0 0 0 1.038 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookMarkOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.75 8A4.75 4.75 0 0 1 8.5 3.25h10c.966 0 1.75.784 1.75 1.75v15a1.75 1.75 0 0 1-1.75 1.75h-11A3.75 3.75 0 0 1 3.75 18zm15-3v9.25H7.5c-.844 0-1.623.279-2.25.75V8A3.25 3.25 0 0 1 8.5 4.75h3.208a20.793 20.793 0 0 0 .071 5.97l.063.397a.75.75 0 0 0 1.292.392l1.366-1.48l1.366 1.48a.75.75 0 0 0 1.292-.392l.063-.397a20.79 20.79 0 0 0 .071-5.97H18.5a.25.25 0 0 1 .25.25m-2.972-.25h-2.556a19.293 19.293 0 0 0-.108 4.57l.651-.706a1 1 0 0 1 1.47 0l.651.706a19.293 19.293 0 0 0-.108-4.57m-8.278 11h11.25V20a.25.25 0 0 1-.25.25h-11a2.25 2.25 0 0 1 0-4.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookMarkSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.02 3.5c.122 0 .215.108.199.229a20.944 20.944 0 0 0 .103 6.225l.014.087a.25.25 0 0 0 .456.095l1.624-2.501a.1.1 0 0 1 .168 0l1.624 2.501a.25.25 0 0 0 .456-.095l.014-.087c.345-2.06.379-4.158.103-6.225a.202.202 0 0 1 .2-.229H18.5A1.5 1.5 0 0 1 20 5v15a1 1 0 0 1-1 1H7.5a3.5 3.5 0 0 1-3.465-3H4V8a4.5 4.5 0 0 1 4.5-4.5zm-4.52 12h11v4h-11a2 2 0 1 1 0-4" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookOpenOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.602 18.636a.75.75 0 0 0 .398.11a.75.75 0 0 0 .398-.11l1.135-.681a8.308 8.308 0 0 1 7.36-.59c.89.356 1.857-.3 1.857-1.257V4.45c0-.578-.352-1.097-.889-1.312a10.701 10.701 0 0 0-9.48.76L12 4.124l-.382-.229a10.701 10.701 0 0 0-9.48-.76A1.413 1.413 0 0 0 1.25 4.45v11.66c0 .957.967 1.612 1.857 1.256a8.308 8.308 0 0 1 7.36.59zM2.75 4.508v11.387a9.809 9.809 0 0 1 8.489.774l.011.006V5.425l-.403-.242a9.201 9.201 0 0 0-8.097-.675m10.011 12.16l-.011.007V5.425l.403-.242a9.201 9.201 0 0 1 8.097-.675v11.387a9.809 9.809 0 0 0-8.489.774" clip-rule="evenodd"/><path fill="currentColor" d="M9.275 19.042a6.5 6.5 0 0 0-6.55 0l-.103.06a.75.75 0 1 0 .756 1.296l.103-.06a5 5 0 0 1 5.038 0l1.088.634a4.75 4.75 0 0 0 4.786 0l1.088-.634a5 5 0 0 1 5.038 0l.103.06a.75.75 0 0 0 .756-1.296l-.103-.06a6.5 6.5 0 0 0-6.55 0l-1.087.634a3.25 3.25 0 0 1-3.276 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookOpenSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.49 4.11a10.451 10.451 0 0 0-9.26-.74a1.163 1.163 0 0 0-.731 1.08v11.66c0 .78.789 1.314 1.514 1.024a8.558 8.558 0 0 1 7.582.608l1.135.68c.087.053.18.075.269.074a.503.503 0 0 0 .27-.073l1.134-.681a8.558 8.558 0 0 1 7.582-.608a1.104 1.104 0 0 0 1.514-1.025V4.45c0-.476-.29-.903-.731-1.08a10.451 10.451 0 0 0-9.259.742l-.51.306zm1.26 2.39a.75.75 0 0 0-1.5 0V16a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/><path fill="currentColor" d="M2.725 19.042a6.5 6.5 0 0 1 6.55 0l1.087.634a3.25 3.25 0 0 0 3.276 0l1.087-.634a6.5 6.5 0 0 1 6.55 0l.103.06a.75.75 0 1 1-.756 1.296l-.103-.06a5 5 0 0 0-5.038 0l-1.088.634a4.75 4.75 0 0 1-4.786 0l-1.088-.634a5 5 0 0 0-5.038 0l-.103.06a.75.75 0 0 1-.756-1.296z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.25 8A.75.75 0 0 1 9 7.25h7a.75.75 0 0 1 0 1.5H9A.75.75 0 0 1 8.25 8M9 10.25a.75.75 0 0 0 0 1.5h5a.75.75 0 0 0 0-1.5z"/><path fill="currentColor" fill-rule="evenodd" d="M8.5 3.25A4.75 4.75 0 0 0 3.75 8v10a3.75 3.75 0 0 0 3.75 3.75h11A1.75 1.75 0 0 0 20.25 20V5a1.75 1.75 0 0 0-1.75-1.75zm10.25 11V5a.25.25 0 0 0-.25-.25h-10A3.25 3.25 0 0 0 5.25 8v7a3.734 3.734 0 0 1 2.25-.75zm0 1.5H7.5a2.25 2.25 0 0 0 0 4.5h11a.25.25 0 0 0 .25-.25z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 8a4.5 4.5 0 0 1 4.5-4.5h10A1.5 1.5 0 0 1 20 5v15a1 1 0 0 1-1 1H7.5a3.5 3.5 0 0 1-3.465-3H4zm14.5 7.5h-11a2 2 0 1 0 0 4h11zM8.25 8A.75.75 0 0 1 9 7.25h7a.75.75 0 0 1 0 1.5H9A.75.75 0 0 1 8.25 8M9 10.25a.75.75 0 0 0 0 1.5h5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.924 4.418a23.778 23.778 0 0 0-7.848 0a1.194 1.194 0 0 0-.973.94a35.64 35.64 0 0 0-.142 13.327l.163.913l3.67-3.487a1.75 1.75 0 0 1 2.411 0l3.67 3.487l.164-.913a35.64 35.64 0 0 0-.142-13.327a1.194 1.194 0 0 0-.973-.94M7.828 2.94a25.278 25.278 0 0 1 8.344 0a2.694 2.694 0 0 1 2.195 2.123c.923 4.579.973 9.29.149 13.888l-.222 1.235a1.323 1.323 0 0 1-2.213.726l-3.909-3.713a.25.25 0 0 0-.344 0l-3.909 3.713a1.323 1.323 0 0 1-2.213-.726l-.222-1.235a37.14 37.14 0 0 1 .148-13.888A2.694 2.694 0 0 1 7.828 2.94" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.13 3.186a25.028 25.028 0 0 0-8.26 0A2.444 2.444 0 0 0 5.877 5.11a36.89 36.89 0 0 0-.148 13.795l.334 1.86a.732.732 0 0 0 1.224.4l4.023-3.822a1 1 0 0 1 1.378 0l4.023 3.822a.732.732 0 0 0 1.224-.4l.334-1.86a36.89 36.89 0 0 0-.148-13.795a2.444 2.444 0 0 0-1.991-1.925"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.75 11a.75.75 0 0 1-.75.75H8a.75.75 0 0 1 0-1.5h8a.75.75 0 0 1 .75.75"/><path fill="currentColor" fill-rule="evenodd" d="m4.386 8.147l1.188-3.248a2.358 2.358 0 0 1 1.89-1.525a32.912 32.912 0 0 1 9.072 0c.864.12 1.59.707 1.89 1.525l1.189 3.248c.222.347.375.743.44 1.17l.101.679a24.67 24.67 0 0 1-.11 8.011a2.664 2.664 0 0 1-2.36 2.182l-1.23.122a45.116 45.116 0 0 1-8.912 0l-1.23-.122a2.664 2.664 0 0 1-2.36-2.182a24.675 24.675 0 0 1-.11-8.011l.102-.678a2.99 2.99 0 0 1 .44-1.171M7.67 4.86c2.873-.4 5.787-.4 8.66 0a.858.858 0 0 1 .687.555l.506 1.383a3.045 3.045 0 0 0-.137-.017l-.978-.097a44.631 44.631 0 0 0-8.816 0l-.978.097a2.05 2.05 0 0 0-.137.017l.506-1.383a.858.858 0 0 1 .687-.555M5.686 8.9l.018.007l.023-.063c.246-.316.615-.53 1.035-.571l.978-.097a43.132 43.132 0 0 1 8.52 0l.978.097c.42.041.789.255 1.035.57l.023.064l.018-.007a1.5 1.5 0 0 1 .257.64l.102.678c.375 2.497.34 5.038-.104 7.524a1.164 1.164 0 0 1-1.03.954l-1.231.122c-2.865.284-5.75.284-8.616 0l-1.23-.122a1.164 1.164 0 0 1-1.031-.954a23.176 23.176 0 0 1-.104-7.524l.102-.678a1.5 1.5 0 0 1 .257-.64" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.594 4.99L4.537 8.378a2.742 2.742 0 0 0-.344.977l-.102.678a24.426 24.426 0 0 0 .11 7.93a2.414 2.414 0 0 0 2.137 1.977l1.23.123c2.948.292 5.916.292 8.863 0l1.23-.122a2.414 2.414 0 0 0 2.138-1.978c.469-2.62.506-5.298.11-7.93l-.102-.678a2.74 2.74 0 0 0-.365-1.013l-1.037-3.347a2.342 2.342 0 0 0-1.914-1.627a32.594 32.594 0 0 0-8.983 0A2.343 2.343 0 0 0 5.594 4.99m10.69-.137a31.094 31.094 0 0 0-8.57 0a.843.843 0 0 0-.688.584l-.501 1.606c.037-.005.075-.01.113-.013l.978-.097c2.916-.29 5.852-.29 8.768 0l.978.097c.036.003.071.008.107.012l-.497-1.604a.841.841 0 0 0-.688-.585M16 11.75a.75.75 0 0 0 0-1.5H8a.75.75 0 0 0 0 1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BullhornOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M20.085 16.932a62.152 62.152 0 0 0 0-12.864c-.159-1.525-1.917-2.298-3.148-1.383l-4.103 3.05a7.25 7.25 0 0 1-4.326 1.432H4.68c-.69 0-1.25.56-1.25 1.25v4.166c0 .69.56 1.25 1.25 1.25h.624l-.845 3.155a.75.75 0 0 0 .408.874l3.68 1.717a.75.75 0 0 0 1.042-.486l1.263-4.712a.78.78 0 0 0 .024-.15a7.25 7.25 0 0 1 1.959 1.034l4.103 3.05c1.23.915 2.99.142 3.148-1.383m-1.492-12.71a60.696 60.696 0 0 1 0 12.555a.478.478 0 0 1-.76.334l-4.105-3.05a8.75 8.75 0 0 0-5.22-1.728H4.93V8.667h3.58a8.75 8.75 0 0 0 5.22-1.728l4.103-3.05a.478.478 0 0 1 .76.334m-9.157 9.67a7.256 7.256 0 0 0-.928-.059H6.856L6.07 16.77l2.3 1.073l1.032-3.85a.759.759 0 0 1 .034-.1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BullhornSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.836 16.906c.443-4.26.443-8.553 0-12.813c-.138-1.332-1.675-2.007-2.75-1.208l-4.103 3.05a7.5 7.5 0 0 1-4.475 1.482H4.68a1 1 0 0 0-1 1v4.166a1 1 0 0 0 1 1h3.83a7.5 7.5 0 0 1 4.474 1.481l4.103 3.05c1.075.8 2.612.125 2.75-1.208"/><path fill="currentColor" d="M9.093 15.802a.75.75 0 0 0-.727-.936h-2a.75.75 0 0 0-.67.415l-1 2a.75.75 0 0 0 .262.964l2.134 1.384a.75.75 0 0 0 1.135-.443z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 3.25a.75.75 0 0 1 .75.75v1.668a47.911 47.911 0 0 1 8.5 0V4a.75.75 0 0 1 1.5 0v1.816a3.375 3.375 0 0 1 2.872 2.899l.087.653c.364 2.746.332 5.53-.094 8.268a3.01 3.01 0 0 1-2.678 2.532l-1.193.118a48.345 48.345 0 0 1-9.488 0l-1.193-.118a3.01 3.01 0 0 1-2.678-2.532a28.995 28.995 0 0 1-.094-8.268l.087-.653A3.375 3.375 0 0 1 6.25 5.816V4A.75.75 0 0 1 7 3.25m.445 3.953c3.03-.299 6.08-.299 9.11 0l.905.09c.867.085 1.56.756 1.675 1.619l.087.653c.03.228.057.456.082.685H4.696c.025-.229.052-.457.082-.685l.087-.653a1.875 1.875 0 0 1 1.675-1.62zM4.577 11.75a27.495 27.495 0 0 0 .29 5.655a1.51 1.51 0 0 0 1.343 1.27l1.193.118c3.057.302 6.137.302 9.194 0l1.193-.118a1.51 1.51 0 0 0 1.343-1.27c.292-1.872.388-3.767.29-5.655z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.75 4a.75.75 0 0 0-1.5 0v1.816a3.375 3.375 0 0 0-2.872 2.899l-.087.653c-.015.11-.029.221-.042.332a.493.493 0 0 0 .492.55H20.26a.493.493 0 0 0 .492-.55c-.014-.11-.027-.222-.042-.332l-.087-.653a3.375 3.375 0 0 0-2.872-2.899V4a.75.75 0 0 0-1.5 0v1.668a47.911 47.911 0 0 0-8.5 0zm13.195 8.226a.494.494 0 0 0-.496-.476H3.551a.494.494 0 0 0-.496.476a28.92 28.92 0 0 0 .33 5.41a3.01 3.01 0 0 0 2.678 2.532l1.193.118c3.155.31 6.333.31 9.488 0l1.193-.118a3.01 3.01 0 0 0 2.678-2.532a28.99 28.99 0 0 0 .33-5.41"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8.25 13a3.75 3.75 0 1 1 7.5 0a3.75 3.75 0 0 1-7.5 0M12 10.75a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5"/><path d="M10.616 5.75a1.892 1.892 0 0 0-1.892 1.892a1.28 1.28 0 0 1-1.175 1.275l-2.23.18a.975.975 0 0 0-.88.792a20.82 20.82 0 0 0-.098 7.181l.097.604c.092.57.561 1.005 1.137 1.052l1.943.158a55.43 55.43 0 0 0 8.964 0l1.942-.158a1.255 1.255 0 0 0 1.138-1.052l.097-.604a20.82 20.82 0 0 0-.099-7.181a.975.975 0 0 0-.88-.791l-2.229-.181a1.278 1.278 0 0 1-1.175-1.275a1.892 1.892 0 0 0-1.892-1.892zM7.23 7.438a3.393 3.393 0 0 1 3.386-3.188h2.768c1.805 0 3.28 1.41 3.386 3.188l2.032.165a2.474 2.474 0 0 1 2.232 2.007c.481 2.54.517 5.145.106 7.698l-.097.605a2.755 2.755 0 0 1-2.497 2.308l-1.942.158a56.882 56.882 0 0 1-9.208 0l-1.942-.158a2.755 2.755 0 0 1-2.497-2.308l-.097-.605a22.319 22.319 0 0 1 .106-7.698a2.475 2.475 0 0 1 2.232-2.007z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.75 13a2.25 2.25 0 1 1 4.5 0a2.25 2.25 0 0 1-4.5 0"/><path fill="currentColor" fill-rule="evenodd" d="M7.474 7.642A3.142 3.142 0 0 1 10.616 4.5h2.768a3.142 3.142 0 0 1 3.142 3.142a.03.03 0 0 0 .026.029l2.23.18c.999.082 1.82.82 2.007 1.805a22.07 22.07 0 0 1 .104 7.613l-.097.604a2.505 2.505 0 0 1-2.27 2.099l-1.943.157a56.61 56.61 0 0 1-9.166 0l-1.943-.157a2.505 2.505 0 0 1-2.27-2.1l-.097-.603c-.407-2.525-.371-5.1.104-7.613a2.226 2.226 0 0 1 2.007-1.804l2.23-.181a.028.028 0 0 0 .026-.029M12 9.25a3.75 3.75 0 1 0 0 7.5a3.75 3.75 0 0 0 0-7.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CancelOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.47 8.47a.75.75 0 0 1 1.06 0L12 10.94l2.47-2.47a.75.75 0 1 1 1.06 1.06L13.06 12l2.47 2.47a.75.75 0 0 1-1.06 1.06L12 13.06l-2.47 2.47a.75.75 0 1 1-1.06-1.06L10.94 12L8.47 9.53a.75.75 0 0 1 0-1.06"/><path fill="currentColor" fill-rule="evenodd" d="M7.317 3.769a42.502 42.502 0 0 1 9.366 0c1.827.204 3.302 1.643 3.516 3.48c.37 3.157.37 6.346 0 9.503c-.215 1.837-1.69 3.275-3.516 3.48a42.5 42.5 0 0 1-9.366 0c-1.827-.205-3.302-1.643-3.516-3.48a40.903 40.903 0 0 1 0-9.503c.214-1.837 1.69-3.276 3.516-3.48m9.2 1.49a41.001 41.001 0 0 0-9.034 0A2.486 2.486 0 0 0 5.29 7.424a39.402 39.402 0 0 0 0 9.154a2.486 2.486 0 0 0 2.193 2.164c2.977.332 6.057.332 9.034 0a2.486 2.486 0 0 0 2.192-2.164a39.401 39.401 0 0 0 0-9.154a2.486 2.486 0 0 0-2.192-2.163" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CancelSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.345 4.017a42.253 42.253 0 0 1 9.31 0c1.713.192 3.095 1.541 3.296 3.26a40.66 40.66 0 0 1 0 9.446c-.201 1.719-1.583 3.068-3.296 3.26a42.245 42.245 0 0 1-9.31 0c-1.713-.192-3.095-1.541-3.296-3.26a40.652 40.652 0 0 1 0-9.445a3.734 3.734 0 0 1 3.295-3.26M8.47 8.47a.75.75 0 0 1 1.06 0L12 10.94l2.47-2.47a.75.75 0 1 1 1.06 1.06L13.06 12l2.47 2.47a.75.75 0 0 1-1.06 1.06L12 13.06l-2.47 2.47a.75.75 0 1 1-1.06-1.06L10.94 12L8.47 9.53a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M19.184 4.912L12 4.75l-7.184.162a2.809 2.809 0 0 0-2.699 2.3a26.484 26.484 0 0 0 0 9.575a2.809 2.809 0 0 0 2.7 2.3L12 19.25l7.184-.162a2.809 2.809 0 0 0 2.699-2.3c.581-3.166.581-6.41 0-9.575a2.809 2.809 0 0 0-2.7-2.3M4.85 6.412L12 6.25l7.15.162c.62.014 1.145.461 1.257 1.072c.154.834.264 1.674.332 2.516H3.26a24.89 24.89 0 0 1 .332-2.516A1.309 1.309 0 0 1 4.85 6.412M3.181 12c0 1.512.137 3.023.412 4.516c.112.61.637 1.058 1.257 1.072l7.15.162l7.15-.162a1.309 1.309 0 0 0 1.257-1.072c.275-1.493.412-3.004.412-4.516z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M19.184 4.666L12 4.5l-7.184.166A3.057 3.057 0 0 0 1.87 7.17a26.718 26.718 0 0 0 0 9.66a3.058 3.058 0 0 0 2.945 2.504L12 19.5l7.184-.166a3.058 3.058 0 0 0 2.945-2.504a26.717 26.717 0 0 0 0-9.66a3.057 3.057 0 0 0-2.945-2.504M21 11a1 1 0 0 1-1 1H4a1 1 0 1 1 0-2h16a1 1 0 0 1 1 1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretDownOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.53 8.97a.75.75 0 0 1 0 1.06l-4 4a.75.75 0 0 1-1.06 0l-4-4a.75.75 0 1 1 1.06-1.06L12 12.44l3.47-3.47a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretDownSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.53 8.97a.75.75 0 0 1 0 1.06l-4 4a.75.75 0 0 1-1.06 0l-4-4a.75.75 0 1 1 1.06-1.06L12 12.44l3.47-3.47a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretLeftOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.03 7.47a.75.75 0 0 1 0 1.06L10.56 12l3.47 3.47a.75.75 0 1 1-1.06 1.06l-4-4a.75.75 0 0 1 0-1.06l4-4a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretLeftSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.03 7.47a.75.75 0 0 1 0 1.06L10.56 12l3.47 3.47a.75.75 0 1 1-1.06 1.06l-4-4a.75.75 0 0 1 0-1.06l4-4a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretRightOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.97 7.47a.75.75 0 0 1 1.06 0l4 4a.75.75 0 0 1 0 1.06l-4 4a.75.75 0 1 1-1.06-1.06L13.44 12L9.97 8.53a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretRightSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.97 7.47a.75.75 0 0 1 1.06 0l4 4a.75.75 0 0 1 0 1.06l-4 4a.75.75 0 1 1-1.06-1.06L13.44 12L9.97 8.53a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretUpOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.53 14.03a.75.75 0 0 1-1.06 0L12 10.56l-3.47 3.47a.75.75 0 0 1-1.06-1.06l4-4a.75.75 0 0 1 1.06 0l4 4a.75.75 0 0 1 0 1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretUpSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.53 14.03a.75.75 0 0 1-1.06 0L12 10.56l-3.47 3.47a.75.75 0 0 1-1.06-1.06l4-4a.75.75 0 0 1 1.06 0l4 4a.75.75 0 0 1 0 1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartPieAltOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M11.25 5.788a7.251 7.251 0 1 0 7.962 7.962H12a.75.75 0 0 1-.75-.75zM3.25 13A8.75 8.75 0 0 1 12 4.25a.75.75 0 0 1 .75.75v7.25H20a.75.75 0 0 1 .75.75a8.75 8.75 0 1 1-17.5 0"/><path d="M15.5 4.674V9.5h4.826A6.512 6.512 0 0 0 15.5 4.674m-.502-1.612c3.62.45 6.49 3.32 6.94 6.94c.069.548-.386.998-.938.998h-6.5a.5.5 0 0 1-.5-.5V4c0-.552.45-1.007.998-.938"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartPieAltSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.938 10.002a8.004 8.004 0 0 0-6.94-6.94C14.45 2.993 14 3.448 14 4v6.5a.5.5 0 0 0 .5.5H21c.552 0 1.007-.45.938-.998"/><path fill="currentColor" d="M12 4.5a8.5 8.5 0 1 0 8.5 8.5a.5.5 0 0 0-.5-.5h-7a.5.5 0 0 1-.5-.5V5a.5.5 0 0 0-.5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartPieOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.25 3.784a8.25 8.25 0 1 0 6.669 13.964l-6.39-5.165A.75.75 0 0 1 11.25 12zm1.5 0v7.466h7.466a8.252 8.252 0 0 0-7.466-7.466m7.466 8.966h-6.095l4.741 3.831a8.198 8.198 0 0 0 1.354-3.831M2.25 12c0-5.385 4.365-9.75 9.75-9.75s9.75 4.365 9.75 9.75a9.712 9.712 0 0 1-2.167 6.13A9.733 9.733 0 0 1 12 21.75A9.75 9.75 0 0 1 2.25 12" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartPieSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.25 2.745c0-.116-.1-.208-.216-.196A9.5 9.5 0 0 0 2.5 12a9.5 9.5 0 0 0 16.243 6.692a.197.197 0 0 0-.017-.292l-7.197-5.817A.75.75 0 0 1 11.25 12z"/><path fill="currentColor" d="M19.67 17.234c.09.073.224.054.288-.044a9.446 9.446 0 0 0 1.494-4.225a.197.197 0 0 0-.197-.215h-6.568a.2.2 0 0 0-.126.355zm1.585-5.984c.116 0 .208-.1.197-.216a9.503 9.503 0 0 0-8.486-8.486a.197.197 0 0 0-.216.197v8.205a.3.3 0 0 0 .3.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 9.5A1.25 1.25 0 1 0 8 12a1.25 1.25 0 0 0 0-2.5m4 0a1.25 1.25 0 1 0 0 2.5a1.25 1.25 0 0 0 0-2.5m2.75 1.25a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0"/><path fill="currentColor" fill-rule="evenodd" d="M16.1 4.593a50.577 50.577 0 0 0-8.098-.04l-.193.015A4.93 4.93 0 0 0 3.25 9.483V18a.75.75 0 0 0 1.105.66l3.91-2.101a1.25 1.25 0 0 1 .593-.149h8.976c1.132 0 2.102-.81 2.305-1.923c.412-2.257.444-4.567.096-6.835l-.102-.669a2.666 2.666 0 0 0-2.408-2.252zM8.116 6.049a49.078 49.078 0 0 1 7.858.038l1.624.139c.536.046.972.453 1.053.985l.103.668a19.165 19.165 0 0 1-.09 6.339a.843.843 0 0 1-.829.692H8.858a2.75 2.75 0 0 0-1.302.328L4.75 16.746V9.483a3.43 3.43 0 0 1 3.171-3.42z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.002 4.553a50.577 50.577 0 0 1 8.099.04l1.624.138a2.666 2.666 0 0 1 2.408 2.252l.102.669a20.665 20.665 0 0 1-.096 6.835a2.343 2.343 0 0 1-2.305 1.923H8.858c-.207 0-.41.051-.592.149l-3.911 2.102A.75.75 0 0 1 3.25 18V9.483a4.93 4.93 0 0 1 4.559-4.915zM8 9.5A1.25 1.25 0 1 0 8 12a1.25 1.25 0 0 0 0-2.5m4 0a1.25 1.25 0 1 0 0 2.5a1.25 1.25 0 0 0 0-2.5m2.75 1.25a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M18.03 7.97a.75.75 0 0 1 0 1.06l-7 7a.75.75 0 0 1-1.06 0l-4-4a.75.75 0 1 1 1.06-1.06l3.47 3.47l6.47-6.47a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M18.03 7.97a.75.75 0 0 1 0 1.06l-7 7a.75.75 0 0 1-1.06 0l-4-4a.75.75 0 1 1 1.06-1.06l3.47 3.47l6.47-6.47a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckedBoxOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.483 5.26A2.486 2.486 0 0 0 5.29 7.422a39.402 39.402 0 0 0 0 9.154a2.486 2.486 0 0 0 2.193 2.163c2.977.333 6.057.333 9.034 0a2.486 2.486 0 0 0 2.192-2.163c.256-2.185.328-4.386.216-6.58a.2.2 0 0 1 .059-.152l1.038-1.04a.198.198 0 0 1 .339.125a40.903 40.903 0 0 1-.162 7.822c-.215 1.836-1.69 3.275-3.516 3.48a42.5 42.5 0 0 1-9.366 0c-1.827-.205-3.302-1.644-3.516-3.48a40.903 40.903 0 0 1 0-9.504c.214-1.837 1.69-3.275 3.516-3.48a42.502 42.502 0 0 1 9.366 0a3.989 3.989 0 0 1 1.76.64a.19.19 0 0 1 .025.295l-.803.803a.211.211 0 0 1-.25.033a2.488 2.488 0 0 0-.898-.28a41.001 41.001 0 0 0-9.034 0"/><path fill="currentColor" d="M21.03 6.03a.75.75 0 1 0-1.06-1.06l-8.47 8.47l-2.47-2.47a.75.75 0 1 0-1.06 1.06l3 3a.75.75 0 0 0 1.06 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckedBoxSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.317 3.769a42.502 42.502 0 0 1 9.366 0a4 4 0 0 1 2.82 1.667L11.5 13.439l-2.47-2.47a.75.75 0 1 0-1.06 1.061l3 3a.75.75 0 0 0 1.06 0l8.116-8.115c.022.11.04.22.053.333c.37 3.157.37 6.347 0 9.504c-.215 1.836-1.69 3.275-3.516 3.48a42.5 42.5 0 0 1-9.366 0c-1.827-.205-3.302-1.644-3.516-3.48a40.903 40.903 0 0 1 0-9.504c.214-1.837 1.69-3.275 3.516-3.48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChromeOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.11 6.27A9.706 9.706 0 0 0 2.25 12c0 4.997 3.759 9.116 8.604 9.683A9.844 9.844 0 0 0 12 21.75A9.75 9.75 0 0 0 21.75 12a9.72 9.72 0 0 0-.832-3.948A9.752 9.752 0 0 0 12 2.25a9.735 9.735 0 0 0-7.8 3.9a.746.746 0 0 0-.09.12m.616 1.833a8.252 8.252 0 0 0 5.87 12.028l2.226-3.859a4.35 4.35 0 0 1-4.68-2.26zm.934-1.382l2.228 3.855A4.352 4.352 0 0 1 12 7.65h7.011A8.245 8.245 0 0 0 12 3.75a8.233 8.233 0 0 0-6.34 2.97m14.084 2.43h-4.458A4.333 4.333 0 0 1 16.35 12a4.33 4.33 0 0 1-.675 2.329l-3.413 5.917A8.25 8.25 0 0 0 19.744 9.15M12 9.15a2.85 2.85 0 1 0 0 5.7a2.85 2.85 0 0 0 0-5.7" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChromeSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.139 7.197a.195.195 0 0 0-.336-.002A9.456 9.456 0 0 0 2.5 12c0 4.462 3.076 8.206 7.224 9.226a.198.198 0 0 0 .216-.092l2.466-4.14l.198-.343c.085-.148-.044-.333-.214-.318a4.35 4.35 0 0 1-4.247-2.32l-1.97-3.406z"/><path fill="currentColor" d="M9.566 13.483a.78.78 0 0 0-.03-.058l-.074-.127a2.85 2.85 0 1 1 5.039.07a.763.763 0 0 0-.037.057l-.07.12a2.848 2.848 0 0 1-4.83-.063"/><path fill="currentColor" d="m15.675 14.329l-1.978 3.428l-2.051 3.445a.194.194 0 0 0 .161.296L12 21.5a9.5 9.5 0 0 0 9.5-9.5a9.5 9.5 0 0 0-.392-2.71a.197.197 0 0 0-.19-.14h-5.193c-.172 0-.267.205-.168.345c.5.708.793 1.572.793 2.505a4.33 4.33 0 0 1-.675 2.329M12 7.65h8.12c.15 0 .245-.158.172-.29A9.497 9.497 0 0 0 12 2.5a9.474 9.474 0 0 0-6.949 3.022a.197.197 0 0 0-.024.236L7.47 9.852l.2.345c.085.148.31.129.382-.027A4.35 4.35 0 0 1 12 7.65"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardAltOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m15.65 4.263l.813.101a3.75 3.75 0 0 1 3.287 3.721v10.49a3.641 3.641 0 0 1-3.191 3.613c-3.028.377-6.09.377-9.118 0a3.641 3.641 0 0 1-3.191-3.613V8.085a3.75 3.75 0 0 1 3.287-3.72l.813-.102A2.751 2.751 0 0 1 11 2.25h2a2.75 2.75 0 0 1 2.65 2.013m-7.4 1.524l-.528.066A2.25 2.25 0 0 0 5.75 8.085v10.49c0 1.08.805 1.991 1.876 2.125a35.39 35.39 0 0 0 8.747 0a2.141 2.141 0 0 0 1.877-2.125V8.085a2.25 2.25 0 0 0-1.972-2.232l-.528-.066V7a.75.75 0 0 1-.75.75H9A.75.75 0 0 1 8.25 7zM9.75 5c0-.69.56-1.25 1.25-1.25h2c.69 0 1.25.56 1.25 1.25v1.25h-4.5z" clip-rule="evenodd"/><path fill="currentColor" d="M15.75 11.75A.75.75 0 0 0 15 11H9a.75.75 0 0 0 0 1.5h6a.75.75 0 0 0 .75-.75m-1 3A.75.75 0 0 0 14 14H9a.75.75 0 0 0 0 1.5h5a.75.75 0 0 0 .75-.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardAltSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8.25 5A2.75 2.75 0 0 1 11 2.25h2A2.75 2.75 0 0 1 15.75 5v2a.75.75 0 0 1-.75.75H9A.75.75 0 0 1 8.25 7zM11 3.75c-.69 0-1.25.56-1.25 1.25v1.25h4.5V5c0-.69-.56-1.25-1.25-1.25z"/><path d="M6.487 4.929c.126-.06.267.036.266.176L6.75 7A2.25 2.25 0 0 0 9 9.25h6A2.25 2.25 0 0 0 17.25 7V5.104c0-.14.14-.236.267-.175A3.498 3.498 0 0 1 19.5 8.085v10.49a3.39 3.39 0 0 1-2.972 3.365a36.639 36.639 0 0 1-9.056 0A3.391 3.391 0 0 1 4.5 18.575V8.085a3.5 3.5 0 0 1 1.987-3.156M15 12a.75.75 0 0 1 0 1.5H9A.75.75 0 0 1 9 12zm-1 3a.75.75 0 0 1 0 1.5H9A.75.75 0 0 1 9 15z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m15.65 4.263l.813.101a3.75 3.75 0 0 1 3.287 3.721v10.49a3.641 3.641 0 0 1-3.191 3.613c-3.028.377-6.09.377-9.118 0a3.641 3.641 0 0 1-3.191-3.613V8.085a3.75 3.75 0 0 1 3.287-3.72l.813-.102A2.751 2.751 0 0 1 11 2.25h2a2.75 2.75 0 0 1 2.65 2.013m-7.4 1.524l-.528.066A2.25 2.25 0 0 0 5.75 8.085v10.49c0 1.08.805 1.991 1.876 2.125a35.39 35.39 0 0 0 8.747 0a2.141 2.141 0 0 0 1.877-2.125V8.085a2.25 2.25 0 0 0-1.972-2.232l-.528-.066V7a.75.75 0 0 1-.75.75H9A.75.75 0 0 1 8.25 7zM9.75 5c0-.69.56-1.25 1.25-1.25h2c.69 0 1.25.56 1.25 1.25v1.25h-4.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 2.25A2.75 2.75 0 0 0 8.25 5v2c0 .414.336.75.75.75h6a.75.75 0 0 0 .75-.75V5A2.75 2.75 0 0 0 13 2.25zM9.75 5c0-.69.56-1.25 1.25-1.25h2c.69 0 1.25.56 1.25 1.25v1.25h-4.5z" clip-rule="evenodd"/><path fill="currentColor" d="M6.753 5.105c0-.14-.14-.237-.266-.176A3.501 3.501 0 0 0 4.5 8.085v10.49a3.391 3.391 0 0 0 2.972 3.365c3.007.374 6.049.374 9.056 0a3.39 3.39 0 0 0 2.972-3.365V8.085a3.498 3.498 0 0 0-1.983-3.156c-.127-.061-.267.035-.267.175V7A2.25 2.25 0 0 1 15 9.25H9A2.25 2.25 0 0 1 6.75 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.75 7a.75.75 0 0 0-1.5 0v5a.75.75 0 0 0 .352.636l3 1.875a.75.75 0 1 0 .796-1.272l-2.648-1.655z"/><path fill="currentColor" fill-rule="evenodd" d="M12 3.25a8.75 8.75 0 1 0 0 17.5a8.75 8.75 0 0 0 0-17.5M4.75 12a7.25 7.25 0 1 1 14.5 0a7.25 7.25 0 0 1-14.5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 12a8.5 8.5 0 1 1 17 0a8.5 8.5 0 0 1-17 0m9.25-5a.75.75 0 0 0-1.5 0v5a.75.75 0 0 0 .352.636l3 1.875a.75.75 0 1 0 .796-1.272l-2.648-1.655z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudCheckOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M15.03 10.47a.75.75 0 0 1 0 1.06l-3.5 3.5a.75.75 0 0 1-1.06 0l-2-2a.75.75 0 0 1 1.06-1.06L11 13.44l2.97-2.97a.75.75 0 0 1 1.06 0"/><path d="M12.932 6.208a3.91 3.91 0 0 0-3.524 2.214a.75.75 0 0 1-.947.373a4.375 4.375 0 1 0-1.586 8.455h11.648a2.978 2.978 0 1 0-.77-5.855a.75.75 0 0 1-.939-.812a3.91 3.91 0 0 0-3.882-4.374m-4.552.986A5.41 5.41 0 0 1 18.332 9.8a4.478 4.478 0 1 1 .191 8.951H6.875A5.875 5.875 0 1 1 8.38 7.194"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudCheckSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.38 7.194a5.41 5.41 0 0 1 9.952 2.605a4.478 4.478 0 1 1 .191 8.951H6.875A5.875 5.875 0 1 1 8.38 7.194m6.65 4.336a.75.75 0 0 0-1.06-1.06L11 13.44l-1.47-1.47a.75.75 0 1 0-1.06 1.06l2 2a.75.75 0 0 0 1.06 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDownloadOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.086 12.532a.75.75 0 0 0-1.055-.118L12.75 13.44V10a.75.75 0 1 0-1.5 0v3.44l-1.282-1.026a.75.75 0 0 0-.937 1.172l2.498 1.998a.747.747 0 0 0 .946-.003l2.494-1.995a.75.75 0 0 0 .117-1.055"/><path fill="currentColor" fill-rule="evenodd" d="M8.46 7.284a5.296 5.296 0 0 1 9.734 2.543a4.386 4.386 0 1 1 .17 8.77H7A5.75 5.75 0 1 1 8.46 7.284m4.45-.922a3.796 3.796 0 0 0-3.422 2.15a.75.75 0 0 1-.947.372A4.25 4.25 0 1 0 7 17.096h11.362a2.887 2.887 0 1 0-.747-5.675a.75.75 0 0 1-.938-.812a3.795 3.795 0 0 0-3.769-4.247" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDownloadSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.38 7.194a5.41 5.41 0 0 1 9.952 2.605a4.478 4.478 0 1 1 .191 8.951H6.875A5.875 5.875 0 1 1 8.38 7.194M12 15.75c.18 0 .345-.063.475-.17l2.494-1.994a.75.75 0 0 0-.938-1.172L12.75 13.44V10a.75.75 0 0 0-1.5 0v3.44l-1.282-1.025a.75.75 0 1 0-.937 1.172l2.498 1.998a.747.747 0 0 0 .465.166z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudOffOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M21.03 4.83a.75.75 0 0 0-1.06-1.06l-16 16a.75.75 0 0 0 1.06 1.06l2.08-2.08h11.413a4.478 4.478 0 1 0-.19-8.951a5.38 5.38 0 0 0-.437-1.834zm-4.31 4.312L8.61 17.25h9.912a2.978 2.978 0 1 0-.77-5.854a.75.75 0 0 1-.939-.813a3.957 3.957 0 0 0-.095-1.44" clip-rule="evenodd"/><path fill="currentColor" d="M12.932 4.708c1.107 0 2.136.333 2.993.903a.24.24 0 0 1 .032.371l-.728.728a.261.261 0 0 1-.317.036a3.91 3.91 0 0 0-5.504 1.676a.75.75 0 0 1-.947.373a4.375 4.375 0 0 0-3.708 7.906c.152.086.19.295.067.419l-.724.723a.243.243 0 0 1-.299.038A5.875 5.875 0 0 1 8.38 7.195a5.405 5.405 0 0 1 4.552-2.487"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudOffSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.03 4.83a.75.75 0 0 0-1.06-1.06l-16 16a.75.75 0 0 0 1.06 1.06l2.08-2.08h11.413a4.478 4.478 0 1 0-.19-8.951a5.38 5.38 0 0 0-.437-1.834zm-8.098-.122c1.107 0 2.136.333 2.993.903a.24.24 0 0 1 .032.371L4.097 17.843a.243.243 0 0 1-.3.038A5.875 5.875 0 0 1 8.38 7.195a5.405 5.405 0 0 1 4.552-2.487"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.932 6.208a3.91 3.91 0 0 0-3.524 2.214a.75.75 0 0 1-.947.373a4.375 4.375 0 1 0-1.586 8.455h11.648a2.978 2.978 0 1 0-.77-5.855a.75.75 0 0 1-.939-.812a3.91 3.91 0 0 0-3.882-4.374m-4.552.986A5.41 5.41 0 0 1 18.332 9.8a4.478 4.478 0 1 1 .191 8.951H6.875A5.875 5.875 0 1 1 8.38 7.194" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.932 4.708c-1.912 0-3.59.992-4.552 2.486A5.875 5.875 0 1 0 6.875 18.75h11.648a4.478 4.478 0 1 0-.19-8.951a5.41 5.41 0 0 0-5.401-5.09"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUploadOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.086 12.469a.75.75 0 0 1-1.055.117l-1.28-1.026V15a.75.75 0 1 1-1.5 0v-3.44l-1.282 1.026a.75.75 0 0 1-.937-1.172l2.497-1.998A.747.747 0 0 1 12 9.25h.002c.18 0 .344.064.473.17l2.494 1.994a.75.75 0 0 1 .117 1.055"/><path fill="currentColor" fill-rule="evenodd" d="M8.38 7.194a5.41 5.41 0 0 1 9.952 2.605a4.478 4.478 0 1 1 .191 8.951H6.875A5.875 5.875 0 1 1 8.38 7.194m4.552-.986a3.91 3.91 0 0 0-3.524 2.214a.75.75 0 0 1-.947.373a4.375 4.375 0 1 0-1.586 8.455h11.648a2.978 2.978 0 1 0-.77-5.855a.75.75 0 0 1-.939-.812a3.91 3.91 0 0 0-3.882-4.374" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUploadSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.46 7.284a5.296 5.296 0 0 1 9.734 2.543a4.386 4.386 0 1 1 .17 8.77H7A5.75 5.75 0 1 1 8.46 7.284m6.626 5.185a.75.75 0 0 1-1.055.117L12.75 11.56V15a.75.75 0 0 1-1.5 0v-3.44l-1.282 1.025a.75.75 0 1 1-.937-1.172l2.498-1.997a.747.747 0 0 1 .465-.167h.008c.18 0 .344.064.473.17l2.494 1.994a.75.75 0 0 1 .117 1.055" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollapseOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M18.028 9.964a.75.75 0 0 0-.75-.75h-2.492V6.722a.75.75 0 0 0-1.5 0v3.242c0 .415.335.75.75.75h3.242a.75.75 0 0 0 .75-.75m-3.992 8.064a.75.75 0 0 0 .75-.75v-2.493h2.492a.75.75 0 0 0 0-1.5h-3.242a.75.75 0 0 0-.75.75v3.243c0 .414.335.75.75.75m-4.072 0a.75.75 0 0 1-.75-.75v-2.493H6.722a.75.75 0 0 1 0-1.5h3.242a.75.75 0 0 1 .75.75v3.243a.75.75 0 0 1-.75.75M5.972 9.964a.75.75 0 0 1 .75-.75h2.492V6.722a.75.75 0 0 1 1.5 0v3.242a.75.75 0 0 1-.75.75H6.722a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollapseSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M18.028 9.964a.75.75 0 0 0-.75-.75h-2.492V6.722a.75.75 0 0 0-1.5 0v3.242c0 .415.335.75.75.75h3.242a.75.75 0 0 0 .75-.75m-3.992 8.064a.75.75 0 0 0 .75-.75v-2.493h2.492a.75.75 0 0 0 0-1.5h-3.242a.75.75 0 0 0-.75.75v3.243c0 .414.335.75.75.75m-4.072 0a.75.75 0 0 1-.75-.75v-2.493H6.722a.75.75 0 0 1 0-1.5h3.242a.75.75 0 0 1 .75.75v3.243a.75.75 0 0 1-.75.75M5.972 9.964a.75.75 0 0 1 .75-.75h2.492V6.722a.75.75 0 0 1 1.5 0v3.242a.75.75 0 0 1-.75.75H6.722a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ColumnsOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M18.808 4.957a7.012 7.012 0 0 0-4.116 0a73.408 73.408 0 0 0 0 14.086a7.012 7.012 0 0 0 4.116 0a73.424 73.424 0 0 0 0-14.086Zm1.483-.248a74.944 74.944 0 0 1 0 14.582a1.347 1.347 0 0 1-.925 1.148c-1.693.55-3.54.55-5.232 0a1.347 1.347 0 0 1-.925-1.148a74.918 74.918 0 0 1 0-14.582c.053-.535.419-.984.925-1.148a8.508 8.508 0 0 1 5.232 0c.506.164.872.613.925 1.148Zm-10.983.248a7.012 7.012 0 0 0-4.116 0a73.413 73.413 0 0 0 0 14.086a7.012 7.012 0 0 0 4.116 0a73.414 73.414 0 0 0 0-14.086Zm1.483-.248a74.944 74.944 0 0 1 0 14.582a1.347 1.347 0 0 1-.925 1.148c-1.693.55-3.54.55-5.232 0a1.347 1.347 0 0 1-.925-1.148a74.913 74.913 0 0 1 0-14.582c.053-.535.419-.984.925-1.148a8.508 8.508 0 0 1 5.232 0c.506.164.872.613.925 1.148Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ColumnsSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.542 19.267a74.67 74.67 0 0 0 0-14.534c-.043-.435-.34-.8-.753-.935a8.258 8.258 0 0 0-5.078 0c-.412.134-.71.5-.753.935a74.662 74.662 0 0 0 0 14.534c.043.434.341.8.753.934a8.259 8.259 0 0 0 5.078 0c.412-.133.71-.5.753-.934Zm9.5 0a74.67 74.67 0 0 0 0-14.534c-.043-.435-.34-.8-.753-.935a8.258 8.258 0 0 0-5.078 0c-.412.134-.71.5-.753.935a74.67 74.67 0 0 0 0 14.534c.043.434.341.8.753.934a8.259 8.259 0 0 0 5.078 0c.412-.133.71-.5.753-.934Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentBlockOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.534 4.131c.098-.143-.01-.343-.182-.322c-5.395.634-8.879 6.295-6.76 11.495l.932 2.289a.25.25 0 0 1-.075.289l-1.971 1.583a.75.75 0 0 0 .47 1.335h7.82c4.798 0 8.718-3.84 8.97-8.58c.009-.179-.207-.271-.343-.156c-.317.271-.663.51-1.032.71a.43.43 0 0 0-.219.306a7.484 7.484 0 0 1-7.376 6.22H6.079l.31-.248a1.75 1.75 0 0 0 .524-2.025l-.932-2.289c-1.623-3.983.757-8.296 4.68-9.28a.433.433 0 0 0 .294-.254c.157-.379.352-.738.58-1.073"/><path fill="currentColor" fill-rule="evenodd" d="M12.832 10.107a4.5 4.5 0 1 1 7.336-5.215a4.5 4.5 0 0 1-7.336 5.215m2.144-.022a3 3 0 0 0 4.109-4.109zm3.048-5.17l-4.109 4.109a3 3 0 0 1 4.109-4.109" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentBlockSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 7.5c0 .972.308 1.872.832 2.607A4.48 4.48 0 0 0 16.5 12A4.5 4.5 0 1 0 12 7.5m4.5 3a2.985 2.985 0 0 1-1.524-.415l4.109-4.109A3 3 0 0 1 16.5 10.5m-2.585-1.476l4.109-4.109a3 3 0 0 0-4.109 4.109" clip-rule="evenodd"/><path fill="currentColor" d="M4.823 15.21C2.796 10.233 6.057 4.82 11.17 4.085c.17-.025.282.166.194.312a6 6 0 0 0 8.782 7.869c.139-.107.345-.01.333.165a8.733 8.733 0 0 1-8.711 8.119h-7.82a.5.5 0 0 1-.314-.89l1.972-1.583a.5.5 0 0 0 .15-.579z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentMinusOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.1 11.45a.75.75 0 0 1 0 1.5H9.6a.75.75 0 0 1 0-1.5z"/><path fill="currentColor" fill-rule="evenodd" d="M4.592 15.304C2.344 9.787 6.403 3.75 12.36 3.75h.321a8.068 8.068 0 0 1 8.068 8.068a8.982 8.982 0 0 1-8.982 8.982h-7.82a.75.75 0 0 1-.47-1.335l1.971-1.583a.25.25 0 0 0 .075-.29zM12.36 5.25c-4.893 0-8.226 4.957-6.38 9.488l.932 2.289a1.75 1.75 0 0 1-.525 2.024l-.309.249h5.689a7.482 7.482 0 0 0 7.482-7.482a6.568 6.568 0 0 0-6.568-6.568z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentMinusSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.823 15.21C2.643 9.857 6.581 4 12.361 4h.321C17 4 20.5 7.5 20.5 11.818a8.732 8.732 0 0 1-8.732 8.732h-7.82a.5.5 0 0 1-.314-.89l1.972-1.583a.5.5 0 0 0 .15-.579zM15.1 11.45H9.6a.75.75 0 0 0 0 1.5h5.5a.75.75 0 1 0 0-1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.592 15.304C2.344 9.787 6.403 3.75 12.36 3.75h.321a8.068 8.068 0 0 1 8.068 8.068a8.982 8.982 0 0 1-8.982 8.982h-7.82a.75.75 0 0 1-.47-1.335l1.971-1.583a.25.25 0 0 0 .075-.29zM12.36 5.25c-4.893 0-8.226 4.957-6.38 9.488l.932 2.289a1.75 1.75 0 0 1-.525 2.024l-.309.249h5.689a7.482 7.482 0 0 0 7.482-7.482a6.568 6.568 0 0 0-6.568-6.568z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentPlusOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.35 8.7a.75.75 0 0 1 .75.75v2h2a.75.75 0 0 1 0 1.5h-2v2a.75.75 0 0 1-1.5 0v-2h-2a.75.75 0 0 1 0-1.5h2v-2a.75.75 0 0 1 .75-.75"/><path fill="currentColor" fill-rule="evenodd" d="M4.592 15.304C2.344 9.787 6.403 3.75 12.36 3.75h.321a8.068 8.068 0 0 1 8.068 8.068a8.982 8.982 0 0 1-8.982 8.982h-7.82a.75.75 0 0 1-.47-1.335l1.971-1.583a.25.25 0 0 0 .075-.29zM12.36 5.25c-4.893 0-8.226 4.957-6.38 9.488l.932 2.289a1.75 1.75 0 0 1-.525 2.024l-.309.249h5.689a7.482 7.482 0 0 0 7.482-7.482a6.568 6.568 0 0 0-6.568-6.568z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentPlusSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.823 15.21C2.643 9.857 6.581 4 12.361 4h.321C17 4 20.5 7.5 20.5 11.818a8.732 8.732 0 0 1-8.732 8.732h-7.82a.5.5 0 0 1-.314-.89l1.972-1.583a.5.5 0 0 0 .15-.579zM13.1 9.45a.75.75 0 0 0-1.5 0v2h-2a.75.75 0 0 0 0 1.5h2v2a.75.75 0 0 0 1.5 0v-2h2a.75.75 0 0 0 0-1.5h-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.36 4C6.58 4 2.644 9.857 4.824 15.21l.933 2.288a.5.5 0 0 1-.15.579L3.634 19.66a.5.5 0 0 0 .313.89h7.82a8.732 8.732 0 0 0 8.733-8.732C20.5 7.5 17 4 12.682 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContactsOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 9a2 2 0 1 1 4 0a2 2 0 0 1-4 0m1 3.5a3 3 0 0 0-3 3a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1a3 3 0 0 0-3-3z"/><path fill="currentColor" fill-rule="evenodd" d="M7.543 2.883a31.331 31.331 0 0 1 8.913 0a3.196 3.196 0 0 1 2.73 2.874l.126 1.396c.293 3.225.293 6.47 0 9.694l-.127 1.396a3.197 3.197 0 0 1-2.729 2.874a31.334 31.334 0 0 1-8.913 0a3.197 3.197 0 0 1-2.728-2.874l-.127-1.396a53.504 53.504 0 0 1 0-9.694l.127-1.396a3.197 3.197 0 0 1 2.728-2.874m8.7 1.484a29.832 29.832 0 0 0-8.486 0a1.697 1.697 0 0 0-1.448 1.526l-.127 1.396a52.003 52.003 0 0 0 0 9.422l.127 1.396c.07.783.67 1.414 1.448 1.526a29.86 29.86 0 0 0 8.486 0a1.696 1.696 0 0 0 1.448-1.526l.127-1.396a52.009 52.009 0 0 0 0-9.422l-.127-1.396a1.697 1.697 0 0 0-1.448-1.526" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContactsSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.457 2.883a31.332 31.332 0 0 0-8.913 0a3.197 3.197 0 0 0-2.73 2.874l-.126 1.396a53.504 53.504 0 0 0 0 9.694l.127 1.396a3.197 3.197 0 0 0 2.729 2.874c2.955.425 5.957.425 8.912 0a3.197 3.197 0 0 0 2.73-2.874l.126-1.396c.293-3.225.293-6.47 0-9.694l-.127-1.396a3.196 3.196 0 0 0-2.729-2.874M10 9a2 2 0 1 1 4 0a2 2 0 0 1-4 0m-2 6.5a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 3.25A5.75 5.75 0 0 0 3.25 9v7.107a.75.75 0 0 0 1.5 0V9A4.25 4.25 0 0 1 9 4.75h7.013a.75.75 0 0 0 0-1.5z"/><path fill="currentColor" fill-rule="evenodd" d="M18.403 6.793a44.372 44.372 0 0 0-9.806 0a2.011 2.011 0 0 0-1.774 1.76a42.581 42.581 0 0 0 0 9.894a2.01 2.01 0 0 0 1.774 1.76c3.241.362 6.565.362 9.806 0a2.01 2.01 0 0 0 1.774-1.76a42.579 42.579 0 0 0 0-9.894a2.011 2.011 0 0 0-1.774-1.76M8.764 8.284c3.13-.35 6.342-.35 9.472 0a.51.51 0 0 1 .45.444a40.95 40.95 0 0 1 0 9.544a.51.51 0 0 1-.45.444c-3.13.35-6.342.35-9.472 0a.511.511 0 0 1-.45-.444a40.95 40.95 0 0 1 0-9.544a.511.511 0 0 1 .45-.444" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopySolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.25 9A5.75 5.75 0 0 1 9 3.25h7.013a.75.75 0 0 1 0 1.5H9A4.25 4.25 0 0 0 4.75 9v7.107a.75.75 0 0 1-1.5 0z"/><path fill="currentColor" d="M18.403 6.793a44.372 44.372 0 0 0-9.806 0a2.011 2.011 0 0 0-1.774 1.76a42.581 42.581 0 0 0 0 9.894a2.01 2.01 0 0 0 1.774 1.76c3.241.362 6.565.362 9.806 0a2.01 2.01 0 0 0 1.774-1.76a42.579 42.579 0 0 0 0-9.894a2.011 2.011 0 0 0-1.774-1.76"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CrossOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="m8.464 15.535l7.072-7.07m-7.072 0l7.072 7.07"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CrossSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.066 8.995a.75.75 0 1 0-1.06-1.061L12 10.939L8.995 7.934a.75.75 0 1 0-1.06 1.06L10.938 12l-3.005 3.005a.75.75 0 0 0 1.06 1.06L12 13.06l3.005 3.006a.75.75 0 0 0 1.06-1.06L13.062 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrentLocationOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 8.25a3.75 3.75 0 1 0 0 7.5a3.75 3.75 0 0 0 0-7.5"/><path fill="currentColor" fill-rule="evenodd" d="M12 1.25a.75.75 0 0 1 .75.75v1.282a8.752 8.752 0 0 1 7.968 7.968H22a.75.75 0 0 1 0 1.5h-1.282a8.752 8.752 0 0 1-7.968 7.968V22a.75.75 0 0 1-1.5 0v-1.282a8.752 8.752 0 0 1-7.968-7.968H2a.75.75 0 0 1 0-1.5h1.282a8.752 8.752 0 0 1 7.968-7.968V2a.75.75 0 0 1 .75-.75M4.75 12a7.25 7.25 0 1 0 14.5 0a7.25 7.25 0 0 0-14.5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrentLocationSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.75 2a.75.75 0 0 0-1.5 0v1.784a8.252 8.252 0 0 0-7.466 7.466H2a.75.75 0 0 0 0 1.5h1.784a8.252 8.252 0 0 0 7.466 7.466V22a.75.75 0 0 0 1.5 0v-1.784a8.252 8.252 0 0 0 7.466-7.466H22a.75.75 0 0 0 0-1.5h-1.784a8.252 8.252 0 0 0-7.466-7.466zm-4.5 10a3.75 3.75 0 1 1 7.5 0a3.75 3.75 0 0 1-7.5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CursorOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.255 3.724c-.1-1.192 1.235-1.85 2.152-1.238l.296.198a90.863 90.863 0 0 1 12.638 10.18c.97.936.115 2.412-1.094 2.282l-5.266-.564a1.208 1.208 0 0 0-1.234.68l-2.129 4.697c-.514 1.135-2.207 1.03-2.544-.195a86.016 86.016 0 0 1-2.79-15.693zm1.517.142l.007.08a84.508 84.508 0 0 0 2.636 15.035l1.966-4.338c.48-1.06 1.604-1.676 2.76-1.552l4.806.515A89.38 89.38 0 0 0 6.87 3.93z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CursorSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.407 2.486c-.917-.612-2.251.046-2.152 1.238l.029.347a86.016 86.016 0 0 0 2.79 15.693c.337 1.224 2.03 1.33 2.544.195l2.129-4.697c.203-.449.697-.737 1.234-.68l5.266.564c1.209.13 2.063-1.346 1.094-2.281A90.863 90.863 0 0 0 7.703 2.684z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DesktopOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m12 3.75l7.274.205a2.584 2.584 0 0 1 2.494 2.29a37.425 37.425 0 0 1 0 8.51a2.584 2.584 0 0 1-2.494 2.29l-5.524.156v2.05H17a.75.75 0 1 1 0 1.5H7a.75.75 0 0 1 0-1.5h3.25V17.2l-5.524-.156a2.584 2.584 0 0 1-2.494-2.29a37.434 37.434 0 0 1 0-8.51a2.584 2.584 0 0 1 2.494-2.29zm0 1.5l-7.231.205c-.54.015-.985.424-1.047.96a35.934 35.934 0 0 0 0 8.17c.062.536.507.945 1.047.96l7.23.205l7.232-.205c.54-.015.985-.424 1.047-.96a35.94 35.94 0 0 0 0-8.17a1.084 1.084 0 0 0-1.047-.96z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DesktopSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m12 3.75l7.274.205a2.584 2.584 0 0 1 2.494 2.29a37.436 37.436 0 0 1 0 8.51a2.584 2.584 0 0 1-2.494 2.29l-6.024.17v2.035H17a.75.75 0 1 1 0 1.5H7a.75.75 0 0 1 0-1.5h3.75v-2.035l-6.024-.17a2.584 2.584 0 0 1-2.494-2.29a37.434 37.434 0 0 1 0-8.51a2.584 2.584 0 0 1 2.494-2.29zm-7 10a.75.75 0 1 0 0 1.5h14a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DialpadOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.946 1.351a8.578 8.578 0 0 0-1.892 0c-.62.07-1.123.557-1.197 1.187a8.279 8.279 0 0 0 0 1.924c.074.63.578 1.118 1.197 1.187a8.55 8.55 0 0 0 1.892 0a1.357 1.357 0 0 0 1.197-1.187a8.279 8.279 0 0 0 0-1.924a1.357 1.357 0 0 0-1.197-1.187m-1.611 2.82a6.78 6.78 0 0 1 0-1.34c.44-.043.89-.043 1.33 0c.045.445.045.894 0 1.34c-.44.04-.89.04-1.33 0m7.611-2.82a8.579 8.579 0 0 0-1.892 0c-.62.07-1.123.557-1.197 1.187a8.279 8.279 0 0 0 0 1.924c.074.63.578 1.118 1.197 1.187a8.55 8.55 0 0 0 1.892 0a1.357 1.357 0 0 0 1.197-1.187a8.271 8.271 0 0 0 0-1.924a1.357 1.357 0 0 0-1.197-1.187m-1.611 2.82a6.787 6.787 0 0 1 0-1.34c.44-.043.89-.043 1.33 0c.045.445.045.894 0 1.34c-.44.04-.89.04-1.33 0m5.719-2.82a8.579 8.579 0 0 1 1.892 0c.62.07 1.123.557 1.197 1.187a8.271 8.271 0 0 1 0 1.924a1.357 1.357 0 0 1-1.197 1.187a8.55 8.55 0 0 1-1.892 0a1.357 1.357 0 0 1-1.197-1.187a8.271 8.271 0 0 1 0-1.924a1.357 1.357 0 0 1 1.197-1.187m.28 1.48a6.787 6.787 0 0 0 0 1.34c.44.04.892.04 1.332 0a6.772 6.772 0 0 0 0-1.34a7.096 7.096 0 0 0-1.332 0M6.946 6.851a8.578 8.578 0 0 0-1.892 0c-.62.07-1.123.557-1.197 1.187a8.279 8.279 0 0 0 0 1.924c.074.63.578 1.118 1.197 1.187a8.55 8.55 0 0 0 1.892 0a1.357 1.357 0 0 0 1.197-1.187a8.279 8.279 0 0 0 0-1.924a1.357 1.357 0 0 0-1.197-1.187m-1.611 2.82a6.78 6.78 0 0 1 0-1.34c.44-.043.89-.043 1.33 0c.045.445.045.894 0 1.34c-.44.04-.89.04-1.33 0m5.719-2.82a8.579 8.579 0 0 1 1.892 0c.62.07 1.123.557 1.197 1.187a8.271 8.271 0 0 1 0 1.924a1.357 1.357 0 0 1-1.197 1.187a8.55 8.55 0 0 1-1.892 0a1.357 1.357 0 0 1-1.197-1.187a8.279 8.279 0 0 1 0-1.924a1.357 1.357 0 0 1 1.197-1.187m.28 1.48a6.787 6.787 0 0 0 0 1.34c.44.04.892.04 1.332 0a6.772 6.772 0 0 0 0-1.34a7.096 7.096 0 0 0-1.332 0m7.612-1.48a8.579 8.579 0 0 0-1.892 0c-.62.07-1.123.557-1.197 1.187a8.271 8.271 0 0 0 0 1.924c.074.63.578 1.118 1.197 1.187a8.55 8.55 0 0 0 1.892 0a1.357 1.357 0 0 0 1.197-1.187a8.271 8.271 0 0 0 0-1.924a1.357 1.357 0 0 0-1.197-1.187m-1.612 2.82a6.787 6.787 0 0 1 0-1.34c.44-.043.892-.043 1.332 0c.044.445.044.894 0 1.34c-.44.04-.892.04-1.332 0m-12.28 2.68a8.582 8.582 0 0 1 1.892 0c.62.07 1.123.557 1.197 1.187a8.279 8.279 0 0 1 0 1.924a1.357 1.357 0 0 1-1.197 1.187a8.55 8.55 0 0 1-1.892 0a1.357 1.357 0 0 1-1.197-1.187a8.279 8.279 0 0 1 0-1.924a1.357 1.357 0 0 1 1.197-1.187m.28 1.48a6.78 6.78 0 0 0 0 1.34c.44.04.892.04 1.332 0a6.778 6.778 0 0 0 0-1.34a7.096 7.096 0 0 0-1.331 0m7.611-1.48a8.582 8.582 0 0 0-1.892 0c-.62.07-1.123.557-1.197 1.187a8.279 8.279 0 0 0 0 1.924c.074.63.578 1.118 1.197 1.187a8.55 8.55 0 0 0 1.892 0a1.357 1.357 0 0 0 1.197-1.187a8.271 8.271 0 0 0 0-1.924a1.357 1.357 0 0 0-1.197-1.187m-1.611 2.82a6.787 6.787 0 0 1 0-1.34c.44-.043.89-.043 1.33 0c.045.445.045.894 0 1.34c-.44.04-.89.04-1.33 0m-.281 2.68a8.582 8.582 0 0 1 1.892 0c.62.07 1.123.557 1.197 1.187a8.271 8.271 0 0 1 0 1.924a1.357 1.357 0 0 1-1.197 1.187a8.55 8.55 0 0 1-1.892 0a1.357 1.357 0 0 1-1.197-1.187a8.279 8.279 0 0 1 0-1.924a1.357 1.357 0 0 1 1.197-1.187m.28 1.48a6.787 6.787 0 0 0 0 1.34c.44.04.892.04 1.332 0a6.772 6.772 0 0 0 0-1.34a7.096 7.096 0 0 0-1.332 0m7.612-6.98a8.582 8.582 0 0 0-1.892 0c-.62.07-1.123.557-1.197 1.187a8.271 8.271 0 0 0 0 1.924c.074.63.578 1.118 1.197 1.187a8.55 8.55 0 0 0 1.892 0a1.357 1.357 0 0 0 1.197-1.187a8.271 8.271 0 0 0 0-1.924a1.357 1.357 0 0 0-1.197-1.187m-1.612 2.82a6.787 6.787 0 0 1 0-1.34c.44-.043.892-.043 1.332 0c.044.445.044.894 0 1.34c-.44.04-.892.04-1.332 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DialpadSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.918 1.6a8.328 8.328 0 0 0-1.836 0a1.107 1.107 0 0 0-.976.967a8.028 8.028 0 0 0 0 1.866c.06.513.47.911.976.968a8.41 8.41 0 0 0 1.836 0c.506-.057.916-.455.976-.968a8.03 8.03 0 0 0 0-1.866a1.107 1.107 0 0 0-.976-.967m6 0a8.328 8.328 0 0 0-1.836 0a1.107 1.107 0 0 0-.976.967a8.026 8.026 0 0 0 0 1.866c.06.513.47.911.976.968a8.41 8.41 0 0 0 1.836 0c.506-.057.916-.455.976-.968a8.026 8.026 0 0 0 0-1.866a1.107 1.107 0 0 0-.976-.967m6 0a8.328 8.328 0 0 0-1.836 0a1.107 1.107 0 0 0-.976.967a8.026 8.026 0 0 0 0 1.866c.06.513.47.911.976.968a8.41 8.41 0 0 0 1.836 0c.506-.057.916-.455.976-.968a8.026 8.026 0 0 0 0-1.866a1.107 1.107 0 0 0-.976-.967m-12 5.5a8.328 8.328 0 0 0-1.836 0a1.107 1.107 0 0 0-.976.967a8.028 8.028 0 0 0 0 1.866c.06.513.47.911.976.968a8.41 8.41 0 0 0 1.836 0c.506-.057.916-.455.976-.968a8.03 8.03 0 0 0 0-1.866a1.107 1.107 0 0 0-.976-.967m6 0a8.328 8.328 0 0 0-1.836 0a1.107 1.107 0 0 0-.976.967a8.026 8.026 0 0 0 0 1.866c.06.513.47.911.976.968a8.41 8.41 0 0 0 1.836 0c.506-.057.916-.455.976-.968a8.026 8.026 0 0 0 0-1.866a1.107 1.107 0 0 0-.976-.967m6 0a8.328 8.328 0 0 0-1.836 0a1.107 1.107 0 0 0-.976.967a8.026 8.026 0 0 0 0 1.866c.06.513.47.911.976.968a8.41 8.41 0 0 0 1.836 0c.506-.057.916-.455.976-.968a8.026 8.026 0 0 0 0-1.866a1.107 1.107 0 0 0-.976-.967m-12 5.5a8.324 8.324 0 0 0-1.836 0a1.107 1.107 0 0 0-.976.967a8.028 8.028 0 0 0 0 1.866c.06.513.47.911.976.968a8.41 8.41 0 0 0 1.836 0c.506-.057.916-.455.976-.968a8.029 8.029 0 0 0 0-1.866a1.107 1.107 0 0 0-.976-.967m6 0a8.324 8.324 0 0 0-1.836 0a1.107 1.107 0 0 0-.976.967a8.026 8.026 0 0 0 0 1.866c.06.513.47.911.976.968a8.41 8.41 0 0 0 1.836 0c.506-.057.916-.455.976-.968a8.026 8.026 0 0 0 0-1.866a1.107 1.107 0 0 0-.976-.967m0 5.5a8.324 8.324 0 0 0-1.836 0a1.107 1.107 0 0 0-.976.967a8.026 8.026 0 0 0 0 1.866c.06.513.47.911.976.968a8.41 8.41 0 0 0 1.836 0c.506-.057.916-.455.976-.968a8.026 8.026 0 0 0 0-1.866a1.107 1.107 0 0 0-.976-.967m6-5.5a8.324 8.324 0 0 0-1.836 0a1.107 1.107 0 0 0-.976.967a8.026 8.026 0 0 0 0 1.866c.06.513.47.911.976.968a8.41 8.41 0 0 0 1.836 0c.506-.057.916-.455.976-.968a8.026 8.026 0 0 0 0-1.866a1.107 1.107 0 0 0-.976-.967"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiamondOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.535 3.25a1.75 1.75 0 0 0-1.456.78L1.49 9.412c-.219.33-.245.752-.068 1.106A32.952 32.952 0 0 0 9.45 20.8l1.75 1.501a1.23 1.23 0 0 0 1.601 0l1.75-1.5a32.951 32.951 0 0 0 8.029-10.283c.177-.354.15-.776-.069-1.106L18.92 4.03a1.75 1.75 0 0 0-1.455-.779zm-.208 1.611a.25.25 0 0 1 .208-.111h2.34L6.96 9.346a1.75 1.75 0 0 0-.1.326a69.284 69.284 0 0 1-2.11-.19l-1.405-.147zM3.28 10.836a31.45 31.45 0 0 0 6.994 8.695l-3.125-8.334a70.78 70.78 0 0 1-2.554-.222zm5.506.454L12 19.864l3.215-8.574a70.793 70.793 0 0 1-6.43 0m8.067-.093l-3.125 8.334a31.45 31.45 0 0 0 6.994-8.695l-1.315.139c-.85.089-1.701.163-2.554.222m3.803-1.862l-1.406.148a69.27 69.27 0 0 1-2.11.19a1.747 1.747 0 0 0-.099-.327L15.125 4.75h2.34a.25.25 0 0 1 .208.111zm-5.063.435c-2.393.124-4.79.124-7.184 0L10.5 4.75h3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiamondSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.375 20.4a.196.196 0 0 0 .318-.203l-2.52-8.638a.401.401 0 0 0-.356-.288c-.408-.03-.815-.063-1.222-.1l-3.37-.307a.198.198 0 0 0-.194.29c1.823 3.41 4.448 6.66 7.344 9.246m2.472 1.826c.02.07.08.122.153.122a.163.163 0 0 0 .153-.122l3.084-10.573a.2.2 0 0 0-.202-.257a70.797 70.797 0 0 1-6.07 0a.2.2 0 0 0-.202.256zm2.46-2.029c-.056.19.17.335.318.203c2.896-2.587 5.521-5.836 7.344-9.246a.198.198 0 0 0-.193-.29l-3.37.307c-.408.037-.815.07-1.223.1a.401.401 0 0 0-.356.288zm7.494-10.841a.2.2 0 0 0 .147-.313L18.71 4.348a1.5 1.5 0 0 0-1.235-.648h-2.49a.2.2 0 0 0-.185.276l2.257 5.523a.4.4 0 0 0 .4.247c.27-.021.542-.044.813-.07zm-8.636-5.41a.4.4 0 0 0-.37-.246h-1.59a.4.4 0 0 0-.37.246L8.477 9.604a.2.2 0 0 0 .175.276c2.23.108 4.466.108 6.697 0a.2.2 0 0 0 .174-.276zm-3.965.03a.2.2 0 0 0-.185-.276h-2.49a1.5 1.5 0 0 0-1.235.648L2.052 9.043a.2.2 0 0 0 .147.313l3.531.32c.271.026.542.049.813.07a.399.399 0 0 0 .4-.247z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DislikeOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.298 18.75c.166 0 .32-.09.402-.234l3.346-5.873a4.25 4.25 0 0 0 .531-2.568l-.353-3.214a1.236 1.236 0 0 0-1.083-1.092l-2.117-.252a9.568 9.568 0 0 0-4.15.422c-.66.22-1.188.72-1.442 1.367l-1.704 4.33a1.421 1.421 0 0 0 1.565 1.92l3.416-.593a1.562 1.562 0 0 1 1.8 1.84l-.665 3.395a.463.463 0 0 0 .454.552m1.705.509a1.962 1.962 0 0 1-3.631-1.35l.665-3.395a.058.058 0 0 0 0-.03a.067.067 0 0 0-.017-.025a.066.066 0 0 0-.025-.017a.057.057 0 0 0-.03-.001l-3.415.593a2.92 2.92 0 0 1-3.218-3.947l1.704-4.33a3.845 3.845 0 0 1 2.365-2.241a11.068 11.068 0 0 1 4.8-.489l2.117.252a2.736 2.736 0 0 1 2.397 2.419l.353 3.213a5.75 5.75 0 0 1-.72 3.475z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DislikeSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.663 20.25a1.4 1.4 0 0 0 1.149-.6l4.232-6.077a5.75 5.75 0 0 0 .997-3.914l-.326-2.961a2.736 2.736 0 0 0-2.396-2.419L14.2 4.027a11.068 11.068 0 0 0-4.8.489a3.845 3.845 0 0 0-2.365 2.24L5.327 11.1a2.818 2.818 0 0 0 3.243 3.78l3.77-.85a.074.074 0 0 1 .041 0c.011.004.023.01.035.023a.083.083 0 0 1 .02.034a.074.074 0 0 1 0 .04l-1.128 4.374a1.4 1.4 0 0 0 1.356 1.749"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.75 13a.75.75 0 0 0-.75-.75H9a.75.75 0 0 0 0 1.5h6a.75.75 0 0 0 .75-.75m0 4a.75.75 0 0 0-.75-.75H9a.75.75 0 0 0 0 1.5h6a.75.75 0 0 0 .75-.75"/><path fill="currentColor" fill-rule="evenodd" d="M7 2.25A2.75 2.75 0 0 0 4.25 5v14A2.75 2.75 0 0 0 7 21.75h10A2.75 2.75 0 0 0 19.75 19V7.968c0-.381-.124-.751-.354-1.055l-2.998-3.968a1.75 1.75 0 0 0-1.396-.695zM5.75 5c0-.69.56-1.25 1.25-1.25h7.25v4.397c0 .414.336.75.75.75h3.25V19c0 .69-.56 1.25-1.25 1.25H7c-.69 0-1.25-.56-1.25-1.25z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.25 2.5a.25.25 0 0 0-.25-.25H7A2.75 2.75 0 0 0 4.25 5v14A2.75 2.75 0 0 0 7 21.75h10A2.75 2.75 0 0 0 19.75 19V9.147a.25.25 0 0 0-.25-.25H15a.75.75 0 0 1-.75-.75zm.75 9.75a.75.75 0 0 1 0 1.5H9a.75.75 0 0 1 0-1.5zm0 4a.75.75 0 0 1 0 1.5H9a.75.75 0 0 1 0-1.5z" clip-rule="evenodd"/><path fill="currentColor" d="M15.75 2.824c0-.184.193-.301.336-.186c.121.098.23.212.323.342l3.013 4.197c.068.096-.006.22-.124.22H16a.25.25 0 0 1-.25-.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.738 3.25c-.774 0-1.419.591-1.487 1.362a37.118 37.118 0 0 0-.107 4.845l-.253.018l-1.49.108a1.26 1.26 0 0 0-.97 1.935a16.055 16.055 0 0 0 4.163 4.395l.596.429a1.388 1.388 0 0 0 1.62 0l.596-.429a16.054 16.054 0 0 0 4.163-4.395a1.26 1.26 0 0 0-.97-1.935l-1.49-.108a42.986 42.986 0 0 0-.253-.018c.07-1.615.034-3.234-.107-4.845a1.492 1.492 0 0 0-1.487-1.362zm-.056 6.865a35.624 35.624 0 0 1 .063-5.365h2.51a35.61 35.61 0 0 1 .064 5.365a.75.75 0 0 0 .711.796c.324.016.647.036.97.06l1.081.079a14.555 14.555 0 0 1-3.55 3.645l-.53.381l-.532-.381a14.554 14.554 0 0 1-3.55-3.645L9 10.97c.323-.023.647-.043.97-.059a.75.75 0 0 0 .712-.796" clip-rule="evenodd"/><path fill="currentColor" d="M5.75 17a.75.75 0 0 0-1.5 0v2c0 .966.784 1.75 1.75 1.75h12A1.75 1.75 0 0 0 19.75 19v-2a.75.75 0 0 0-1.5 0v2a.25.25 0 0 1-.25.25H6a.25.25 0 0 1-.25-.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 16.25a.75.75 0 0 1 .75.75v2c0 .138.112.25.25.25h12a.25.25 0 0 0 .25-.25v-2a.75.75 0 0 1 1.5 0v2A1.75 1.75 0 0 1 18 20.75H6A1.75 1.75 0 0 1 4.25 19v-2a.75.75 0 0 1 .75-.75" clip-rule="evenodd"/><path fill="currentColor" d="M10.738 3.75a.992.992 0 0 0-.988.906a36.618 36.618 0 0 0-.082 5.27c-.247.013-.493.03-.74.047l-1.49.109a.76.76 0 0 0-.585 1.167a15.555 15.555 0 0 0 4.032 4.258l.597.429a.888.888 0 0 0 1.036 0l.597-.429a15.556 15.556 0 0 0 4.032-4.258a.76.76 0 0 0-.585-1.167l-1.49-.109a42.274 42.274 0 0 0-.74-.047a36.62 36.62 0 0 0-.081-5.27a.992.992 0 0 0-.989-.906z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DribbbleOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.96 4.004a34.766 34.766 0 0 1 3.26 4.885a22.338 22.338 0 0 0 4.82-2.508A8.226 8.226 0 0 0 12 3.75c-.704 0-1.387.088-2.04.254m9 3.565a23.848 23.848 0 0 1-5.04 2.659c.285.576.555 1.162.808 1.756a13.89 13.89 0 0 1 5.517.297a8.208 8.208 0 0 0-1.285-4.713m1.097 6.215a12.288 12.288 0 0 0-4.752-.364a34.449 34.449 0 0 1 1.513 5.278a8.256 8.256 0 0 0 3.239-4.914m-4.612 5.715a32.94 32.94 0 0 0-1.642-5.806a12.285 12.285 0 0 0-6.992 4.721A8.213 8.213 0 0 0 12 20.25a8.24 8.24 0 0 0 3.445-.751m-9.719-2.141a13.79 13.79 0 0 1 7.501-5.066a33.177 33.177 0 0 0-.737-1.569a23.9 23.9 0 0 1-8.736 1.021a8.215 8.215 0 0 0 1.972 5.613m-1.79-7.105a22.414 22.414 0 0 0 7.837-.88a33.27 33.27 0 0 0-3.306-4.83a8.266 8.266 0 0 0-4.532 5.71m4.47-7.32A9.726 9.726 0 0 1 12 2.25a9.735 9.735 0 0 1 7.733 3.812A9.71 9.71 0 0 1 21.75 12a9.757 9.757 0 0 1-5.112 8.578A9.71 9.71 0 0 1 12 21.75a9.72 9.72 0 0 1-6.7-2.666a9.725 9.725 0 0 1-2.982-8.237a9.76 9.76 0 0 1 6.088-7.913" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DribbbleSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.025 5.573a.195.195 0 0 0 .035-.298A9.722 9.722 0 0 0 12 2.25c-.737 0-1.455.082-2.145.237a.196.196 0 0 0-.11.313a42.31 42.31 0 0 1 3.585 5.453a.2.2 0 0 0 .245.09a27.455 27.455 0 0 0 5.45-2.77M2.677 10.06a.195.195 0 0 1-.183-.238a9.77 9.77 0 0 1 5.422-6.678a.197.197 0 0 1 .236.056a40.823 40.823 0 0 1 3.696 5.474c.062.11.005.25-.117.287a27.563 27.563 0 0 1-9.054 1.1M4.8 18.216a.198.198 0 0 1-.307.006a9.71 9.71 0 0 1-2.24-6.479a.198.198 0 0 1 .21-.19c.478.022.96.034 1.443.034c3 0 5.893-.454 8.616-1.298a.2.2 0 0 1 .238.1c.308.613.6 1.235.877 1.864a.201.201 0 0 1-.134.275A16.682 16.682 0 0 0 4.8 18.216m11.472 2.386a.197.197 0 0 1-.108.217A9.713 9.713 0 0 1 12 21.75a9.71 9.71 0 0 1-6.125-2.164a.195.195 0 0 1-.03-.272a15.176 15.176 0 0 1 8.293-5.4a.197.197 0 0 1 .232.12a40.461 40.461 0 0 1 1.902 6.568m5.102-6.716a.195.195 0 0 1 .145.231a9.753 9.753 0 0 1-3.606 5.636a.196.196 0 0 1-.31-.117a41.926 41.926 0 0 0-1.706-5.803a.198.198 0 0 1 .162-.267a15.34 15.34 0 0 1 5.315.32m-1.401-7.141a.198.198 0 0 1 .277.057a9.704 9.704 0 0 1 1.498 5.386a.199.199 0 0 1-.244.188a16.801 16.801 0 0 0-6.118-.235a.202.202 0 0 1-.215-.12a41.62 41.62 0 0 0-.958-2.079a.202.202 0 0 1 .109-.277c2-.769 3.893-1.751 5.65-2.92"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DropboxOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.097 2.492a.75.75 0 0 1 .805 0L12 5.74l5.097-3.248a.75.75 0 0 1 .806 0l5.5 3.505a.75.75 0 0 1 0 1.265l-4.507 2.871l4.507 2.872a.75.75 0 0 1 0 1.265l-4.606 2.935l.22.146a.75.75 0 0 1 0 1.248l-6.6 4.4a.75.75 0 0 1-.832 0l-6.601-4.4a.75.75 0 0 1 0-1.248l.22-.146L.597 14.27a.75.75 0 0 1 0-1.266l4.507-2.871L.597 7.26a.75.75 0 0 1 0-1.265zm4.482 11.13l-1.068.711L6.5 16.251l-4.104-2.614L6.5 11.023zm3.947.736l-1.104-.737l4.078-2.598l4.104 2.614l-4.104 2.615zM10.33 15.59l-3.578 2.385l5.249 3.499l5.247-3.5l-3.54-2.36L12 14.526zM12 7.518l4.104 2.615L12 12.748l-4.105-2.615zM6.5 4.014L2.396 6.63L6.5 9.243l4.104-2.614zm6.896 2.615L17.5 9.244l4.104-2.615L17.5 4.014z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DropboxSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.66 3.253a.3.3 0 0 0-.322 0L1.398 6.4a.3.3 0 0 0 0 .506l4.94 3.148a.3.3 0 0 0 .323 0l4.942-3.148a.3.3 0 0 0 0-.506zm11.001 0a.3.3 0 0 0-.322 0L12.397 6.4a.3.3 0 0 0 0 .506l4.942 3.149a.3.3 0 0 0 .322 0l4.942-3.149a.3.3 0 0 0 0-.506zM1.397 13.915a.3.3 0 0 1 0-.506l4.941-3.148a.3.3 0 0 1 .323 0l4.942 3.148a.3.3 0 0 1 0 .506L6.66 17.063a.3.3 0 0 1-.323 0zm16.264-3.654a.3.3 0 0 0-.322 0l-4.942 3.148a.3.3 0 0 0 0 .506l4.942 3.149a.3.3 0 0 0 .322 0l4.942-3.149a.3.3 0 0 0 0-.506zM6.897 18.594a.3.3 0 0 1 0-.506l4.942-3.148a.3.3 0 0 1 .323 0l4.94 3.148a.3.3 0 0 1 0 .506l-4.94 3.148a.3.3 0 0 1-.323 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditAltOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.137 3.47a.75.75 0 0 0-1.06 0l-9.193 9.192a.75.75 0 0 0-.195.34l-1 3.83a.75.75 0 0 0 .915.915l3.828-1a.75.75 0 0 0 .341-.196l9.192-9.192a.75.75 0 0 0 0-1.06zM6.088 13.579l8.519-8.518l1.767 1.767l-8.518 8.519l-2.393.625z" clip-rule="evenodd"/><path fill="currentColor" d="M4 19.25a.75.75 0 0 0 0 1.5h15a.75.75 0 0 0 0-1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditAltSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.607 3.5a.5.5 0 0 1 .353.146l2.829 2.829a.5.5 0 0 1 0 .707l-9.193 9.192a.5.5 0 0 1-.227.13l-3.828 1a.5.5 0 0 1-.61-.61l1-3.828a.5.5 0 0 1 .13-.227l9.192-9.193a.5.5 0 0 1 .354-.146M4 19.25a.75.75 0 0 0 0 1.5h15a.75.75 0 0 0 0-1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M21.455 5.416a.75.75 0 0 1-.096.943l-9.193 9.192a.75.75 0 0 1-.34.195l-3.829 1a.75.75 0 0 1-.915-.915l1-3.828a.778.778 0 0 1 .161-.312L17.47 2.47a.75.75 0 0 1 1.06 0l2.829 2.828a.756.756 0 0 1 .096.118m-1.687.412L18 4.061l-8.518 8.518l-.625 2.393l2.393-.625z" clip-rule="evenodd"/><path fill="currentColor" d="M19.641 17.16a44.4 44.4 0 0 0 .261-7.04a.403.403 0 0 1 .117-.3l.984-.984a.198.198 0 0 1 .338.127a45.91 45.91 0 0 1-.21 8.372c-.236 2.022-1.86 3.607-3.873 3.832a47.77 47.77 0 0 1-10.516 0c-2.012-.225-3.637-1.81-3.873-3.832a45.922 45.922 0 0 1 0-10.67c.236-2.022 1.86-3.607 3.873-3.832a47.75 47.75 0 0 1 7.989-.213a.2.2 0 0 1 .128.34l-.993.992a.402.402 0 0 1-.297.117a46.164 46.164 0 0 0-6.66.255a2.89 2.89 0 0 0-2.55 2.516a44.421 44.421 0 0 0 0 10.32a2.89 2.89 0 0 0 2.55 2.516c3.355.375 6.827.375 10.183 0a2.89 2.89 0 0 0 2.55-2.516"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.477 3.004c.167.015.24.219.12.338l-8.32 8.32a.75.75 0 0 0-.195.34l-1 3.83a.75.75 0 0 0 .915.915l3.829-1a.751.751 0 0 0 .34-.196l8.438-8.438a.198.198 0 0 1 .339.12a45.723 45.723 0 0 1-.06 10.073c-.223 1.905-1.754 3.4-3.652 3.613a47.468 47.468 0 0 1-10.461 0c-1.899-.213-3.43-1.708-3.653-3.613a45.672 45.672 0 0 1 0-10.611C3.34 4.789 4.871 3.294 6.77 3.082a47.512 47.512 0 0 1 9.707-.078"/><path fill="currentColor" d="M17.823 4.237a.25.25 0 0 1 .354 0l1.414 1.415a.25.25 0 0 1 0 .353L11.298 14.3a.253.253 0 0 1-.114.065l-1.914.5a.25.25 0 0 1-.305-.305l.5-1.914a.25.25 0 0 1 .065-.114z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeOpenOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M19.806 5.37a3.136 3.136 0 0 1 1.268 2.08c.045.3.086.602.122.903c.28 2.603.268 5.605-.122 8.198a3.138 3.138 0 0 1-2.831 2.66l-1.51.13c-3.15.274-6.317.274-9.465 0l-1.51-.13a3.138 3.138 0 0 1-2.832-2.66c-.39-2.593-.402-5.595-.122-8.198a31.1 31.1 0 0 1 .122-.904A3.136 3.136 0 0 1 4.19 5.373l-.001-.007l.047-.026c.173-.12.358-.223.553-.307l4.416-2.454a5.75 5.75 0 0 1 5.585 0l4.406 2.448c.204.087.398.195.578.321l.032.018zm-.84 1.244l-4.904-2.723a4.25 4.25 0 0 0-4.128 0L5.043 6.608a1.637 1.637 0 0 0-.634 1.065a29.28 29.28 0 0 0-.034.233l5.561 3.09a4.25 4.25 0 0 0 4.128 0l5.56-3.09a26.386 26.386 0 0 0-.033-.233a1.636 1.636 0 0 0-.626-1.059m.74 8.866c.238-1.979.272-3.976.102-5.96l-5.016 2.787a5.75 5.75 0 0 1-5.585 0L4.191 9.52a29.101 29.101 0 0 0 .218 6.807a1.637 1.637 0 0 0 1.478 1.389l1.51.131c3.063.266 6.143.266 9.206 0l1.51-.131a1.638 1.638 0 0 0 1.478-1.389c.042-.281.08-.564.114-.847" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeOpenSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.807 5.687c.298.245.546.55.727.897a.236.236 0 0 1-.091.307l-6.266 3.88a4.25 4.25 0 0 1-4.4.045L3.47 7.088a.236.236 0 0 1-.103-.293a2.89 2.89 0 0 1 .823-1.106v-.003l.012-.007a2.9 2.9 0 0 1 .894-.496l4.11-2.284a5.75 5.75 0 0 1 5.585 0l4.105 2.28c.334.114.641.286.908.505z"/><path fill="currentColor" d="M2.989 8.954a.248.248 0 0 1 .373-.187l5.653 3.34a5.75 5.75 0 0 0 5.951-.061l5.645-3.495a.248.248 0 0 1 .377.183a30.35 30.35 0 0 1-.161 7.78a2.888 2.888 0 0 1-2.606 2.447l-1.51.131a54.394 54.394 0 0 1-9.422 0l-1.51-.131a2.888 2.888 0 0 1-2.606-2.448a30.352 30.352 0 0 1-.184-7.559"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.804 8.353c-.28 2.603-.268 5.605.122 8.198a3.138 3.138 0 0 0 2.831 2.66l1.51.13c3.15.274 6.316.274 9.466 0l1.51-.13a3.138 3.138 0 0 0 2.831-2.66c.39-2.593.402-5.595.122-8.198a30.348 30.348 0 0 0-.122-.904a3.138 3.138 0 0 0-2.831-2.66l-1.51-.13a54.647 54.647 0 0 0-9.465 0l-1.51.13a3.138 3.138 0 0 0-2.832 2.66a31.1 31.1 0 0 0-.122.904m4.593-2.2a53.145 53.145 0 0 1 9.205 0l1.51.131a1.64 1.64 0 0 1 1.479 1.389l.034.233l-5.561 3.09a4.25 4.25 0 0 1-4.128 0l-5.561-3.09l.034-.233a1.638 1.638 0 0 1 1.478-1.389zM19.808 9.52a29.099 29.099 0 0 1-.217 6.807a1.638 1.638 0 0 1-1.478 1.389l-1.51.131a53.152 53.152 0 0 1-9.206 0l-1.51-.131a1.638 1.638 0 0 1-1.478-1.389a29.1 29.1 0 0 1-.218-6.807l5.016 2.787a5.75 5.75 0 0 0 5.585 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.29 4.908a54.397 54.397 0 0 1 9.42 0l1.511.13a2.889 2.889 0 0 1 2.313 1.546a.236.236 0 0 1-.091.307l-6.266 3.88a4.25 4.25 0 0 1-4.4.045L3.47 7.088a.236.236 0 0 1-.103-.293A2.889 2.889 0 0 1 5.78 5.039z"/><path fill="currentColor" d="M3.362 8.767a.248.248 0 0 0-.373.187a30.351 30.351 0 0 0 .184 7.56A2.888 2.888 0 0 0 5.78 18.96l1.51.131c3.135.273 6.287.273 9.422 0l1.51-.13a2.888 2.888 0 0 0 2.606-2.449a30.35 30.35 0 0 0 .161-7.779a.248.248 0 0 0-.377-.182l-5.645 3.494a5.75 5.75 0 0 1-5.951.061z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExchangeOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.53 6.53a.75.75 0 0 0-1.06-1.06l-2 2a.75.75 0 0 0 0 1.06l2 2a.75.75 0 1 0 1.06-1.06l-.72-.72H17a.75.75 0 0 0 0-1.5H6.81zm8.94 6.94a.75.75 0 0 0 0 1.06l.72.72H7a.75.75 0 0 0 0 1.5h10.19l-.72.72a.75.75 0 1 0 1.06 1.06l2-2a.75.75 0 0 0 0-1.06l-2-2a.75.75 0 0 0-1.06 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExchangeSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.53 6.53a.75.75 0 0 0-1.06-1.06l-2 2a.75.75 0 0 0 0 1.06l2 2a.75.75 0 1 0 1.06-1.06l-.72-.72H17a.75.75 0 0 0 0-1.5H6.81zm8.94 6.94a.75.75 0 0 0 0 1.06l.72.72H7a.75.75 0 0 0 0 1.5h10.19l-.72.72a.75.75 0 1 0 1.06 1.06l2-2a.75.75 0 0 0 0-1.06l-2-2a.75.75 0 0 0-1.06 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExpandOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.664 6.343c0 .414.336.75.75.75h2.493v2.493a.75.75 0 0 0 1.5 0V6.343a.75.75 0 0 0-.75-.75h-3.243a.75.75 0 0 0-.75.75m3.993 7.321a.75.75 0 0 0-.75.75v2.493h-2.493a.75.75 0 1 0 0 1.5h3.243a.75.75 0 0 0 .75-.75v-3.243a.75.75 0 0 0-.75-.75m-11.314 0a.75.75 0 0 1 .75.75v2.493h2.493a.75.75 0 0 1 0 1.5H6.343a.75.75 0 0 1-.75-.75v-3.243a.75.75 0 0 1 .75-.75m3.993-7.321a.75.75 0 0 1-.75.75H7.093v2.493a.75.75 0 1 1-1.5 0V6.343a.75.75 0 0 1 .75-.75h3.243a.75.75 0 0 1 .75.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExpandSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.664 6.343c0 .414.336.75.75.75h2.493v2.493a.75.75 0 1 0 1.5 0V6.343a.75.75 0 0 0-.75-.75h-3.243a.75.75 0 0 0-.75.75m3.993 7.321a.75.75 0 0 0-.75.75v2.493h-2.493a.75.75 0 0 0 0 1.5h3.243a.75.75 0 0 0 .75-.75v-3.243a.75.75 0 0 0-.75-.75m-11.314 0a.75.75 0 0 1 .75.75v2.493h2.493a.75.75 0 0 1 0 1.5H6.343a.75.75 0 0 1-.75-.75v-3.243a.75.75 0 0 1 .75-.75m3.993-7.321a.75.75 0 0 1-.75.75H7.093v2.493a.75.75 0 1 1-1.5 0V6.343a.75.75 0 0 1 .75-.75h3.243a.75.75 0 0 1 .75.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExploreOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M12 4.75a7.25 7.25 0 1 0 0 14.5a7.25 7.25 0 0 0 0-14.5M3.25 12a8.75 8.75 0 1 1 17.5 0a8.75 8.75 0 0 1-17.5 0"/><path d="M10.85 9.309a5 5 0 0 0-1.756 3.04l-.438 2.633c-.227 1.368 1.363 2.285 2.433 1.405l2.061-1.696a5 5 0 0 0 1.756-3.04l.438-2.633c.227-1.368-1.363-2.285-2.433-1.405zM12 10.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExploreSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.25 12a8.75 8.75 0 1 1 17.5 0a8.75 8.75 0 0 1-17.5 0m7.6-2.691a5 5 0 0 0-1.756 3.04l-.657 3.949c-.151.911.909 1.523 1.622.936l3.091-2.543a5 5 0 0 0 1.756-3.04l.656-3.949c.152-.911-.908-1.523-1.621-.936z" clip-rule="evenodd"/><path fill="currentColor" d="M10.5 12a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeClosedOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M20.53 4.53a.75.75 0 0 0-1.06-1.06l-16 16a.75.75 0 1 0 1.06 1.06l2.847-2.847c1.367.644 2.94 1.067 4.623 1.067c2.684 0 5.09-1.077 6.82-2.405c.867-.665 1.583-1.407 2.089-2.136c.492-.709.841-1.486.841-2.209c0-.723-.35-1.5-.841-2.209c-.506-.729-1.222-1.47-2.088-2.136c-.263-.201-.54-.397-.832-.583zM16.9 8.161l-1.771 1.771a3.75 3.75 0 0 1-5.197 5.197l-1.417 1.416A9.25 9.25 0 0 0 12 17.25c2.287 0 4.38-.923 5.907-2.095c.762-.585 1.364-1.218 1.77-1.801c.419-.604.573-1.077.573-1.354c0-.277-.154-.75-.573-1.354c-.406-.583-1.008-1.216-1.77-1.801c-.313-.24-.65-.47-1.008-.684m-5.87 5.87a2.25 2.25 0 0 0 3-3z" clip-rule="evenodd"/><path fill="currentColor" d="M12 5.25c1.032 0 2.024.16 2.951.431a.243.243 0 0 1 .1.407l-.824.825a.254.254 0 0 1-.237.067A8.872 8.872 0 0 0 12 6.75c-2.287 0-4.38.923-5.907 2.095c-.762.585-1.364 1.218-1.77 1.801c-.419.604-.573 1.077-.573 1.354c0 .277.154.75.573 1.354c.354.51.858 1.057 1.488 1.577c.116.095.127.27.02.377l-.708.709a.246.246 0 0 1-.333.016a9.52 9.52 0 0 1-1.699-1.824C2.6 13.5 2.25 12.723 2.25 12c0-.723.35-1.5.841-2.209c.506-.729 1.222-1.47 2.088-2.136C6.91 6.327 9.316 5.25 12 5.25"/><path fill="currentColor" d="M12 8.25c.118 0 .236.005.351.016c.197.019.268.254.129.394l-1.213 1.212a2.256 2.256 0 0 0-1.395 1.395L8.66 12.48c-.14.14-.375.068-.394-.129A3.75 3.75 0 0 1 12 8.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeClosedSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M20.53 4.53a.75.75 0 0 0-1.06-1.06l-16 16a.75.75 0 1 0 1.06 1.06l3.035-3.035C8.883 18.103 10.392 18.5 12 18.5c2.618 0 4.972-1.051 6.668-2.353c.85-.652 1.547-1.376 2.035-2.08c.48-.692.797-1.418.797-2.067c0-.649-.317-1.375-.797-2.066c-.488-.705-1.185-1.429-2.035-2.08c-.27-.208-.558-.41-.86-.601zm-5.4 5.402l-1.1 1.098a2.25 2.25 0 0 1-3 3l-1.1 1.1a3.75 3.75 0 0 0 5.197-5.197" clip-rule="evenodd"/><path fill="currentColor" d="M12.67 8.31a.26.26 0 0 0 .23-.07l1.95-1.95a.243.243 0 0 0-.104-.407A10.214 10.214 0 0 0 12 5.5c-2.618 0-4.972 1.051-6.668 2.353c-.85.652-1.547 1.376-2.036 2.08c-.48.692-.796 1.418-.796 2.067c0 .649.317 1.375.796 2.066a9.287 9.287 0 0 0 1.672 1.79a.246.246 0 0 0 .332-.017l2.94-2.94a.26.26 0 0 0 .07-.23a3.75 3.75 0 0 1 4.36-4.36"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8.25 12a3.75 3.75 0 1 1 7.5 0a3.75 3.75 0 0 1-7.5 0M12 9.75a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5"/><path d="M4.323 10.646c-.419.604-.573 1.077-.573 1.354c0 .277.154.75.573 1.354c.406.583 1.008 1.216 1.77 1.801C7.62 16.327 9.713 17.25 12 17.25s4.38-.923 5.907-2.095c.762-.585 1.364-1.218 1.77-1.801c.419-.604.573-1.077.573-1.354c0-.277-.154-.75-.573-1.354c-.406-.583-1.008-1.216-1.77-1.801C16.38 7.673 14.287 6.75 12 6.75s-4.38.923-5.907 2.095c-.762.585-1.364 1.218-1.77 1.801m.856-2.991C6.91 6.327 9.316 5.25 12 5.25s5.09 1.077 6.82 2.405c.867.665 1.583 1.407 2.089 2.136c.492.709.841 1.486.841 2.209c0 .723-.35 1.5-.841 2.209c-.506.729-1.222 1.47-2.088 2.136c-1.73 1.328-4.137 2.405-6.821 2.405s-5.09-1.077-6.82-2.405c-.867-.665-1.583-1.407-2.089-2.136C2.6 13.5 2.25 12.723 2.25 12c0-.723.35-1.5.841-2.209c.506-.729 1.222-1.47 2.088-2.136"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 9.75a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5"/><path fill="currentColor" fill-rule="evenodd" d="M12 5.5c-2.618 0-4.972 1.051-6.668 2.353c-.85.652-1.547 1.376-2.036 2.08c-.48.692-.796 1.418-.796 2.067c0 .649.317 1.375.796 2.066c.49.705 1.186 1.429 2.036 2.08C7.028 17.45 9.382 18.5 12 18.5c2.618 0 4.972-1.051 6.668-2.353c.85-.652 1.547-1.376 2.035-2.08c.48-.692.797-1.418.797-2.067c0-.649-.317-1.375-.797-2.066c-.488-.705-1.185-1.429-2.035-2.08C16.972 6.55 14.618 5.5 12 5.5M8.25 12a3.75 3.75 0 1 1 7.5 0a3.75 3.75 0 0 1-7.5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacebookMessengerOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.777 14.038l2.65-3.92c.262-.386-.235-.805-.615-.529l-2.858 2.015a.571.571 0 0 1-.652 0l-2.116-1.477c-.633-.437-1.538-.277-1.963.335l-2.65 3.92c-.262.386.235.806.615.529l2.858-2.015a.571.571 0 0 1 .652 0l2.116 1.452c.633.462 1.538.302 1.963-.31"/><path fill="currentColor" fill-rule="evenodd" d="M12 2.25A9.75 9.75 0 0 0 2.25 12a9.724 9.724 0 0 0 3 7.036V21.5a.75.75 0 0 0 .987.712l2.78-.927A9.745 9.745 0 0 0 12 21.75c5.385 0 9.75-4.365 9.75-9.75S17.385 2.25 12 2.25M3.75 12a8.25 8.25 0 1 1 5.516 7.787a.75.75 0 0 0-.486-.004l-2.03.676v-1.75a.75.75 0 0 0-.25-.56A8.228 8.228 0 0 1 3.75 12" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacebookMessengerSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 12a9.5 9.5 0 1 1 6.516 9.022l-2.858.952A.5.5 0 0 1 5.5 21.5v-2.572A9.475 9.475 0 0 1 2.5 12m12.277 2.038l2.65-3.92c.262-.386-.235-.805-.615-.529l-2.858 2.015a.571.571 0 0 1-.652 0l-2.116-1.477c-.633-.437-1.538-.277-1.963.335l-2.65 3.92c-.262.386.235.806.615.529l2.858-2.015a.571.571 0 0 1 .652 0l2.116 1.452c.633.462 1.538.302 1.963-.31" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacebookOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.488 3.788A5.25 5.25 0 0 1 14.2 2.25h2.7a.75.75 0 0 1 .75.75v3.6a.75.75 0 0 1-.75.75h-2.7a.15.15 0 0 0-.15.15v1.95h2.85a.75.75 0 0 1 .728.932l-.9 3.6a.75.75 0 0 1-.728.568h-1.95V21a.75.75 0 0 1-.75.75H9.7a.75.75 0 0 1-.75-.75v-6.45H7a.75.75 0 0 1-.75-.75v-3.6A.75.75 0 0 1 7 9.45h1.95V7.5a5.25 5.25 0 0 1 1.538-3.712M14.2 3.75a3.75 3.75 0 0 0-3.75 3.75v2.7a.75.75 0 0 1-.75.75H7.75v2.1H9.7a.75.75 0 0 1 .75.75v6.45h2.1V13.8a.75.75 0 0 1 .75-.75h2.114l.525-2.1H13.3a.75.75 0 0 1-.75-.75V7.5a1.65 1.65 0 0 1 1.65-1.65h1.95v-2.1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacebookSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.2 2.875A4.625 4.625 0 0 0 9.575 7.5v2.575H7.1c-.124 0-.225.1-.225.225v3.4c0 .124.1.225.225.225h2.475V20.9c0 .124.1.225.225.225h3.4c.124 0 .225-.1.225-.225v-6.975h2.497c.103 0 .193-.07.218-.17l.85-3.4a.225.225 0 0 0-.218-.28h-3.347V7.5a.775.775 0 0 1 .775-.775h2.6c.124 0 .225-.1.225-.225V3.1c0-.124-.1-.225-.225-.225z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastForwardOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.533 13.31a25.762 25.762 0 0 1-7.069 3.844l-.466.165c-1.023.364-2.1-.329-2.238-1.38a30.543 30.543 0 0 1 0-7.877c.138-1.052 1.215-1.745 2.238-1.381l.466.166a25.76 25.76 0 0 1 7.07 3.844c.037-.883.112-1.762.226-2.629c.138-1.052 1.215-1.745 2.238-1.381l.466.166a25.76 25.76 0 0 1 7.143 3.9c.82.635.82 1.871 0 2.505a25.761 25.761 0 0 1-7.143 3.902l-.466.165c-1.023.364-2.1-.329-2.238-1.38a30.082 30.082 0 0 1-.227-2.63m-.844-1.376a.084.084 0 0 1 0 .132a24.26 24.26 0 0 1-6.727 3.674l-.466.166a.194.194 0 0 1-.249-.163a29.063 29.063 0 0 1 0-7.486a.194.194 0 0 1 .25-.162l.465.165a24.26 24.26 0 0 1 6.727 3.674m10 0a.084.084 0 0 1 0 .132a24.26 24.26 0 0 1-6.727 3.674l-.466.166a.194.194 0 0 1-.249-.163a29.064 29.064 0 0 1 0-7.486a.194.194 0 0 1 .25-.162l.465.165a24.26 24.26 0 0 1 6.727 3.674" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastForwardSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.764 11.284a1.333 1.333 0 0 0-.31-.339A25.511 25.511 0 0 0 5.38 7.082l-.466-.165c-.87-.31-1.79.279-1.907 1.177a30.314 30.314 0 0 0 0 7.812c.118.898 1.037 1.486 1.907 1.178l.466-.166a25.51 25.51 0 0 0 7.073-3.863a1.34 1.34 0 0 0 .31-.339a29.97 29.97 0 0 0 .244 3.19c.118.898 1.037 1.486 1.907 1.178l.466-.166a25.511 25.511 0 0 0 7.073-3.863c.69-.534.69-1.576 0-2.11a25.512 25.512 0 0 0-7.073-3.863l-.466-.165c-.87-.31-1.79.279-1.907 1.177a29.969 29.969 0 0 0-.244 3.19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastRewindOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.393 13.253a1.584 1.584 0 0 1 0-2.505a25.759 25.759 0 0 1 7.143-3.901l.466-.166c1.023-.364 2.1.329 2.238 1.381c.114.867.19 1.746.227 2.629a25.76 25.76 0 0 1 7.069-3.844l.466-.166c1.023-.364 2.1.329 2.238 1.381c.34 2.59.34 5.286 0 7.876c-.138 1.052-1.215 1.745-2.238 1.381l-.466-.165a25.762 25.762 0 0 1-7.07-3.845a30.082 30.082 0 0 1-.226 2.63c-.138 1.051-1.215 1.744-2.238 1.38l-.466-.165a25.76 25.76 0 0 1-7.143-3.902m.918-1.32a.084.084 0 0 0 0 .133a24.261 24.261 0 0 0 6.727 3.674l.466.166c.1.035.232-.033.249-.163c.322-2.46.322-5.025 0-7.486a.194.194 0 0 0-.25-.162l-.465.165a24.26 24.26 0 0 0-6.727 3.674m10 0a.084.084 0 0 0 0 .133a24.26 24.26 0 0 0 6.727 3.674l.466.166c.1.035.232-.033.249-.163c.322-2.46.322-5.025 0-7.486a.194.194 0 0 0-.25-.162l-.465.165c-2.423.86-4.694 2.1-6.727 3.674" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastRewindSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.546 13.055a1.334 1.334 0 0 1 0-2.11A25.51 25.51 0 0 1 8.62 7.082l.466-.165c.87-.31 1.79.279 1.907 1.177c.138 1.05.22 2.119.244 3.19a1.34 1.34 0 0 1 .31-.339a25.51 25.51 0 0 1 7.073-3.863l.466-.165c.87-.31 1.79.279 1.907 1.177a30.31 30.31 0 0 1 0 7.812c-.118.898-1.037 1.486-1.907 1.178l-.466-.166a25.51 25.51 0 0 1-7.073-3.863a1.331 1.331 0 0 1-.31-.339a29.944 29.944 0 0 1-.244 3.19c-.118.898-1.037 1.486-1.907 1.178l-.466-.166a25.51 25.51 0 0 1-7.073-3.863"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FigmaOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.75 15.373a3.999 3.999 0 0 0 6.195-4.654A3.999 3.999 0 0 0 17.583 9a3.996 3.996 0 0 0 1.668-3.25a4 4 0 0 0-4-4h-6.5A4 4 0 0 0 6.418 9a4 4 0 0 0 0 6.5a4 4 0 1 0 6.332 3.25zm-4-12.123a2.5 2.5 0 1 0 0 5h2.5v-5zm2.5 13h-2.5a2.5 2.5 0 1 0 2.5 2.5zm-2.5-6.5a2.5 2.5 0 0 0 0 5h2.5v-5zm4 2.5a2.5 2.5 0 1 0 5.001 0a2.5 2.5 0 0 0-5.001 0m2.5-4a2.5 2.5 0 1 0 0-5h-2.5v5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FigmaSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.555 3.305A2.75 2.75 0 0 1 8.5 2.5h2.75V8H8.5a2.75 2.75 0 0 1-1.945-4.695M15.5 8h-2.75V2.5h2.75a2.75 2.75 0 0 1 0 5.5m0 1.5a2.75 2.75 0 1 0 0 5.5a2.75 2.75 0 0 0 0-5.5m-8.945 7.806A2.75 2.75 0 0 1 8.5 16.5h2.75v2.75a2.75 2.75 0 1 1-4.695-1.945M8.5 9.5a2.75 2.75 0 1 0 0 5.5h2.75V9.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileDownloadOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.032 13.164a.75.75 0 0 1 .937 1.172l-2.494 1.995a.747.747 0 0 1-.473.169h-.008a.747.747 0 0 1-.465-.166l-2.497-1.998a.75.75 0 0 1 .937-1.172l1.281 1.026v-3.44a.75.75 0 1 1 1.5 0v3.44z"/><path fill="currentColor" fill-rule="evenodd" d="M7 2.25A2.75 2.75 0 0 0 4.25 5v14A2.75 2.75 0 0 0 7 21.75h10A2.75 2.75 0 0 0 19.75 19V8.198a1.75 1.75 0 0 0-.328-1.02L16.408 2.98a1.75 1.75 0 0 0-1.421-.73zM5.75 5c0-.69.56-1.25 1.25-1.25h7.25v4.397c0 .414.336.75.75.75h3.25V19c0 .69-.56 1.25-1.25 1.25H7c-.69 0-1.25-.56-1.25-1.25z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileDownloadSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.25 2.5a.25.25 0 0 0-.25-.25H7A2.75 2.75 0 0 0 4.25 5v14A2.75 2.75 0 0 0 7 21.75h10A2.75 2.75 0 0 0 19.75 19V9.147a.25.25 0 0 0-.25-.25H15a.75.75 0 0 1-.75-.75zm-.219 10.664a.75.75 0 0 1 .938 1.172l-2.494 1.995a.747.747 0 0 1-.473.169h-.008a.747.747 0 0 1-.465-.166l-2.497-1.998a.75.75 0 0 1 .937-1.172l1.281 1.026v-3.44a.75.75 0 1 1 1.5 0v3.44z" clip-rule="evenodd"/><path fill="currentColor" d="M15.75 2.824c0-.184.193-.301.336-.186c.121.098.23.212.323.342l3.013 4.197c.068.096-.006.22-.124.22H16a.25.25 0 0 1-.25-.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.25 5A2.75 2.75 0 0 1 7 2.25h7.987a1.75 1.75 0 0 1 1.422.73l3.013 4.197c.213.298.328.655.328 1.02V19A2.75 2.75 0 0 1 17 21.75H7A2.75 2.75 0 0 1 4.25 19zM7 3.75c-.69 0-1.25.56-1.25 1.25v14c0 .69.56 1.25 1.25 1.25h10c.69 0 1.25-.56 1.25-1.25V8.897H15a.75.75 0 0 1-.75-.75V3.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 2.25a.25.25 0 0 1 .25.25v5.647c0 .414.336.75.75.75h4.5a.25.25 0 0 1 .25.25V19A2.75 2.75 0 0 1 17 21.75H7A2.75 2.75 0 0 1 4.25 19V5A2.75 2.75 0 0 1 7 2.25z"/><path fill="currentColor" d="M16.086 2.638c-.143-.115-.336.002-.336.186v4.323c0 .138.112.25.25.25h3.298c.118 0 .192-.124.124-.22L16.408 2.98a1.748 1.748 0 0 0-.322-.342"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileUploadOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.25 5A2.75 2.75 0 0 1 7 2.25h7.987a1.75 1.75 0 0 1 1.422.73l3.013 4.197c.213.298.328.655.328 1.02V19A2.75 2.75 0 0 1 17 21.75H7A2.75 2.75 0 0 1 4.25 19zM7 3.75c-.69 0-1.25.56-1.25 1.25v14c0 .69.56 1.25 1.25 1.25h10c.69 0 1.25-.56 1.25-1.25V8.897H15a.75.75 0 0 1-.75-.75V3.75z" clip-rule="evenodd"/><path fill="currentColor" d="M15.086 13.219a.75.75 0 0 1-1.055.117l-1.28-1.026v3.44a.75.75 0 0 1-1.5 0v-3.44l-1.282 1.026a.75.75 0 0 1-.937-1.172l2.497-1.998a.747.747 0 0 1 .465-.166h.008c.18 0 .344.064.473.17l2.494 1.994a.75.75 0 0 1 .117 1.055"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileUploadSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.25 2.5a.25.25 0 0 0-.25-.25H7A2.75 2.75 0 0 0 4.25 5v14A2.75 2.75 0 0 0 7 21.75h10A2.75 2.75 0 0 0 19.75 19V9.147a.25.25 0 0 0-.25-.25H15a.75.75 0 0 1-.75-.75zm-.219 10.836a.75.75 0 0 0 .938-1.172l-2.494-1.995a.747.747 0 0 0-.473-.169h-.008a.747.747 0 0 0-.465.166l-2.497 1.998a.75.75 0 0 0 .937 1.172l1.281-1.026v3.44a.75.75 0 1 0 1.5 0v-3.44z" clip-rule="evenodd"/><path fill="currentColor" d="M15.75 2.824c0-.184.193-.301.336-.186c.121.098.23.212.323.342l3.013 4.197c.068.096-.006.22-.124.22H16a.25.25 0 0 1-.25-.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileUserOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 9a2 2 0 1 0 0 4a2 2 0 0 0 0-4m-4 8.5a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1"/><path fill="currentColor" fill-rule="evenodd" d="M7 2.25A2.75 2.75 0 0 0 4.25 5v14A2.75 2.75 0 0 0 7 21.75h10A2.75 2.75 0 0 0 19.75 19V8.198a1.75 1.75 0 0 0-.328-1.02L16.408 2.98a1.75 1.75 0 0 0-1.421-.73zM5.75 5c0-.69.56-1.25 1.25-1.25h7.25v4.397c0 .414.336.75.75.75h3.25V19c0 .69-.56 1.25-1.25 1.25H7c-.69 0-1.25-.56-1.25-1.25z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileUserSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.25 2.5a.25.25 0 0 0-.25-.25H7A2.75 2.75 0 0 0 4.25 5v14A2.75 2.75 0 0 0 7 21.75h10A2.75 2.75 0 0 0 19.75 19V9.147a.25.25 0 0 0-.25-.25H15a.75.75 0 0 1-.75-.75zM12 10a2 2 0 1 0 0 4a2 2 0 0 0 0-4m-4 8.5a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1" clip-rule="evenodd"/><path fill="currentColor" d="M15.75 2.824c0-.184.193-.301.336-.186c.121.098.23.212.323.342l3.013 4.197c.068.096-.006.22-.124.22H16a.25.25 0 0 1-.25-.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M17.986 5.424a53.892 53.892 0 0 0-11.972 0a.343.343 0 0 0-.228.556l3.517 4.348a8.75 8.75 0 0 1 1.947 5.503v2.889l1.5 1.1v-3.99a8.75 8.75 0 0 1 1.947-5.502l3.517-4.348a.343.343 0 0 0-.228-.556M5.848 3.933a55.39 55.39 0 0 1 12.305 0c1.446.162 2.143 1.858 1.228 2.99l-3.517 4.348a7.25 7.25 0 0 0-1.614 4.56V21.3a.75.75 0 0 1-1.194.605l-3-2.2a.75.75 0 0 1-.306-.605v-3.27c0-1.66-.57-3.269-1.613-4.56L4.62 6.924c-.916-1.132-.22-2.828 1.228-2.99" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.523 4.226a58.727 58.727 0 0 0-13.046 0a1.373 1.373 0 0 0-.915 2.229l3.769 4.659A7.5 7.5 0 0 1 10 15.83v3.142a.75.75 0 0 0 .306.605l2.771 2.032a.58.58 0 0 0 .923-.468V15.83a7.5 7.5 0 0 1 1.67-4.717l3.768-4.66a1.373 1.373 0 0 0-.915-2.228"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FireOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.507 1.508a.75.75 0 0 1 .734-.239c3.608.829 5.433 4.783 5.003 8.321c-.166 1.376-.578 2.454-1.185 3.312l-.011.016c.138-.068.265-.141.384-.218c.592-.385 1.026-.892 1.592-1.552l.211-.246a.75.75 0 0 1 1.238.151A7.25 7.25 0 1 1 8.025 8.26l.067-.068a.75.75 0 0 1 .114-.094c.865-.583 1.487-1.06 1.906-1.62c.395-.529.638-1.175.638-2.154a5.73 5.73 0 0 0-.378-2.055a.75.75 0 0 1 .135-.76m1.664 1.743c.052.35.079.708.079 1.072c0 1.268-.328 2.237-.937 3.052c-.571.766-1.363 1.353-2.208 1.925l-.073.073a.75.75 0 0 1-.128.103a5.75 5.75 0 1 0 8.648 3.346c-.38.411-.8.809-1.303 1.136c-.924.599-2.08.941-3.749.941a.75.75 0 0 1-.362-1.407c.679-.374 1.254-.831 1.697-1.456c.44-.624.779-1.457.92-2.626c.307-2.531-.74-5.029-2.584-6.16" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FireSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.063 2.065a.667.667 0 0 0-.759.149a.805.805 0 0 0-.205.738c.099.44.151.9.151 1.37c0 1.076-.271 1.83-.738 2.455c-.474.635-1.16 1.152-2.027 1.735a.25.25 0 0 0-.038.031l-.105.106a6.75 6.75 0 1 0 9.685 2.63a.25.25 0 0 0-.413-.05l-.208.241c-.878 1.026-1.587 1.855-3.04 2.225c-.062.015-.1.004-.127-.013a.244.244 0 0 1-.091-.124a.411.411 0 0 1 .074-.414c.777-.843 1.329-1.987 1.526-3.614c.37-3.048-1.015-6.294-3.685-7.465"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlaskAltOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 15a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-3 2a1 1 0 1 0 0 2a1 1 0 0 0 0-2"/><path fill="currentColor" fill-rule="evenodd" d="M14.495 3.25H8.5a.75.75 0 0 0 0 1.5h.26v5.061c-4.452 2.105-5.095 8.241-1.094 11.209l.295.218a6.78 6.78 0 0 0 8.078 0l.295-.218c4.001-2.968 3.358-9.104-1.094-11.209V4.75h.26a.75.75 0 0 0 0-1.5zm-5.64 16.784a5.28 5.28 0 0 0 6.29 0l.295-.22c2.016-1.494 2.504-4.03 1.644-6.064H6.915c-.86 2.033-.37 4.57 1.645 6.065zm-.958-7.784h8.206a5.07 5.07 0 0 0-1.676-1.16a1.138 1.138 0 0 1-.687-1.045V4.75h-3.48v5.295c0 .454-.27.865-.688 1.045a5.07 5.07 0 0 0-1.675 1.16" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlaskAltSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.495 3.25H8.5a.75.75 0 0 0 0 1.5h.26v5.061c-4.452 2.105-5.095 8.241-1.094 11.209l.295.218a6.78 6.78 0 0 0 8.078 0l.295-.218c4.001-2.968 3.358-9.104-1.094-11.209V4.75h.26a.75.75 0 0 0 0-1.5zm-6.598 9h8.206a5.07 5.07 0 0 0-1.676-1.16a1.138 1.138 0 0 1-.687-1.045V4.75h-3.48v5.295c0 .454-.27.865-.688 1.045a5.07 5.07 0 0 0-1.675 1.16M10 17a1 1 0 1 0 0 2a1 1 0 0 0 0-2m2-1a1 1 0 1 1 2 0a1 1 0 0 1-2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlaskOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 15a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-3 2a1 1 0 1 0 0 2a1 1 0 0 0 0-2"/><path fill="currentColor" fill-rule="evenodd" d="M14.493 3.25H8.5a.75.75 0 1 0 0 1.5h.26v5.087a7.25 7.25 0 0 1-1.256 4.078l-3.093 4.548a1.855 1.855 0 0 0 1.326 2.887c4.162.468 8.364.468 12.526 0a1.855 1.855 0 0 0 1.326-2.887l-3.093-4.548a7.25 7.25 0 0 1-1.256-4.078V4.75h.26a.75.75 0 0 0 0-1.5zM5.905 19.86c4.05.455 8.14.455 12.19 0a.355.355 0 0 0 .254-.553l-3.094-4.549a8.73 8.73 0 0 1-.591-1.008H9.336a8.745 8.745 0 0 1-.591 1.008L5.65 19.307a.355.355 0 0 0 .254.553m4.015-7.61h4.16a8.748 8.748 0 0 1-.34-2.413V4.75h-3.48v5.087a8.75 8.75 0 0 1-.34 2.413" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlaskSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.495 3.25H8.5a.75.75 0 0 0 0 1.5h.26v5.087a7.25 7.25 0 0 1-1.256 4.078l-3.093 4.548a1.855 1.855 0 0 0 1.326 2.887l.087.01l.017.002c4.093.46 8.225.46 12.318 0l.018-.002l.086-.01a1.855 1.855 0 0 0 1.326-2.887l-3.093-4.548a7.25 7.25 0 0 1-1.256-4.078V4.75h.26a.75.75 0 0 0 0-1.5zm-4.666 9.3h4.342a8.75 8.75 0 0 1-.43-2.713V4.75H10.26v5.087a8.75 8.75 0 0 1-.431 2.713M10 17a1 1 0 1 0 0 2a1 1 0 0 0 0-2m2-1a1 1 0 1 1 2 0a1 1 0 0 1-2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderBlockOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.68 9.983c.04.248.075.496.105.745a.421.421 0 0 0 .243.329c.355.165.69.364 1.004.592c.142.104.35-.002.336-.178a22.231 22.231 0 0 0-.207-1.724l-.064-.402a2.658 2.658 0 0 0-2.625-2.239H9.158A2.298 2.298 0 0 0 6.903 5.25H4.612a2.68 2.68 0 0 0-2.66 2.36l-.273 2.27a24.23 24.23 0 0 0 .222 7.243a2.629 2.629 0 0 0 2.398 2.15l1.514.108c1.724.123 3.45.185 5.178.185c.155 0 .251-.17.18-.306a5.96 5.96 0 0 1-.378-.902a.42.42 0 0 0-.392-.294a71.17 71.17 0 0 1-4.481-.18l-1.514-.108a1.128 1.128 0 0 1-1.03-.922a22.73 22.73 0 0 1-.208-6.796l.273-2.27A1.18 1.18 0 0 1 4.61 6.75h2.292c.44 0 .797.357.797.797c0 .585.474 1.06 1.06 1.06h8.712c.57 0 1.054.413 1.144.975z"/><path fill="currentColor" fill-rule="evenodd" d="M12 16.5c0 .972.308 1.872.832 2.607A4.48 4.48 0 0 0 16.5 21a4.5 4.5 0 1 0-4.5-4.5m4.5 3a2.985 2.985 0 0 1-1.524-.415l4.109-4.109A3 3 0 0 1 16.5 19.5m-2.585-1.476l4.109-4.109a3 3 0 0 0-4.109 4.109" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderBlockSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 16.5c0 .806.159 1.575.447 2.277c.056.136-.041.29-.188.289a72.194 72.194 0 0 1-4.91-.184l-1.515-.108a2.128 2.128 0 0 1-1.942-1.74a23.73 23.73 0 0 1-.217-7.095l.273-2.27A2.18 2.18 0 0 1 4.612 5.75h2.291c.993 0 1.797.804 1.797 1.797c0 .033.027.06.06.06h8.712c1.06 0 1.964.77 2.131 1.817l.064.402c.07.43.125.862.167 1.295c.017.167-.17.277-.315.193A6 6 0 0 0 10.5 16.5"/><path fill="currentColor" fill-rule="evenodd" d="M12 16.5c0 .972.308 1.872.832 2.607A4.48 4.48 0 0 0 16.5 21a4.5 4.5 0 1 0-4.5-4.5m4.5 3a2.985 2.985 0 0 1-1.524-.415l4.109-4.109A3 3 0 0 1 16.5 19.5m-2.585-1.476l4.109-4.109a3 3 0 0 0-4.109 4.109" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderDeleteOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.702 11.202a.75.75 0 0 1 1.06 0L12 12.439l1.238-1.237a.75.75 0 1 1 1.06 1.06L13.061 13.5l1.237 1.237a.75.75 0 0 1-1.06 1.061L12 14.561l-1.237 1.237a.75.75 0 1 1-1.061-1.06l1.237-1.238l-1.237-1.237a.75.75 0 0 1 0-1.061"/><path fill="currentColor" fill-rule="evenodd" d="M21.161 9.747a22.229 22.229 0 0 1-.084 7.498a2.468 2.468 0 0 1-2.252 2.019l-1.638.117a72.71 72.71 0 0 1-10.374 0l-1.514-.108a2.629 2.629 0 0 1-2.398-2.15a24.23 24.23 0 0 1-.222-7.244L2.95 7.61a2.68 2.68 0 0 1 2.66-2.36h2.292c1.118 0 2.05.798 2.255 1.856h8.314c1.307 0 2.42.95 2.625 2.24zm-1.56 7.229c.423-2.31.45-4.674.08-6.993l-.065-.401a1.158 1.158 0 0 0-1.144-.976H9.76a1.06 1.06 0 0 1-1.06-1.06a.797.797 0 0 0-.797-.796H5.612c-.597 0-1.1.446-1.171 1.039l-.273 2.269a22.73 22.73 0 0 0 .208 6.796c.093.506.516.886 1.03.922l1.514.109c3.382.242 6.778.242 10.16 0l1.638-.118a.968.968 0 0 0 .884-.791" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderDeleteSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M20.668 9.826c.387 2.43.36 4.909-.083 7.33a1.968 1.968 0 0 1-1.795 1.609l-1.638.117c-3.43.245-6.874.245-10.304 0l-1.514-.108a2.128 2.128 0 0 1-1.942-1.74a23.73 23.73 0 0 1-.217-7.095l.273-2.27A2.18 2.18 0 0 1 5.612 5.75h2.291c.993 0 1.797.804 1.797 1.797c0 .033.027.06.06.06h8.712c1.06 0 1.964.77 2.131 1.817zM9.702 11.202a.75.75 0 0 1 1.06 0L12 12.439l1.237-1.237a.75.75 0 0 1 1.061 1.06L13.061 13.5l1.237 1.237a.75.75 0 0 1-1.06 1.061L12 14.561l-1.237 1.237a.75.75 0 1 1-1.061-1.06l1.237-1.238l-1.237-1.237a.75.75 0 0 1 0-1.061" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderLockOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M18.5 11.25c-1.272 0-2.372.89-2.636 2.135l-.022.104a2.84 2.84 0 0 0-.036.973l.153 1.112l-.17.022a1.765 1.765 0 0 0-1.539 1.75v2.307c0 .888.658 1.637 1.538 1.751c1.8.234 3.624.234 5.424 0a1.765 1.765 0 0 0 1.538-1.75v-2.307c0-.888-.658-1.637-1.538-1.751a19.559 19.559 0 0 0-.171-.022l.153-1.112a2.828 2.828 0 0 0-.036-.973l-.022-.104A2.695 2.695 0 0 0 18.5 11.25m1.044 4.196l.164-1.189c.02-.152.015-.306-.017-.456l-.022-.104a1.195 1.195 0 0 0-2.338 0l-.022.104c-.032.15-.037.304-.017.456l.164 1.19a21.077 21.077 0 0 1 2.088 0m-3.563 1.637a19.564 19.564 0 0 1 5.038 0c.132.018.231.13.231.264v2.306a.265.265 0 0 1-.231.264a19.555 19.555 0 0 1-5.038 0a.265.265 0 0 1-.231-.264v-2.306c0-.134.099-.246.231-.264" clip-rule="evenodd"/><path fill="currentColor" d="M12.75 18.365a.301.301 0 0 0-.303-.3a71.146 71.146 0 0 1-5.527-.18l-1.514-.109a1.128 1.128 0 0 1-1.03-.922a22.73 22.73 0 0 1-.208-6.796l.273-2.27A1.18 1.18 0 0 1 5.61 6.75h2.292c.44 0 .797.357.797.797c0 .585.474 1.06 1.06 1.06h8.712c.57 0 1.054.413 1.144.975l.013.083a.425.425 0 0 0 .282.329c.351.125.682.297.985.508c.15.105.372-.013.347-.195a22.126 22.126 0 0 0-.082-.56l-.064-.402a2.658 2.658 0 0 0-2.625-2.239h-8.314A2.298 2.298 0 0 0 7.903 5.25H5.612a2.68 2.68 0 0 0-2.66 2.36l-.273 2.27a24.23 24.23 0 0 0 .222 7.243a2.629 2.629 0 0 0 2.398 2.15l1.514.108c1.877.134 3.759.195 5.64.184a.299.299 0 0 0 .297-.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderLockSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M18.5 11.25c-1.273 0-2.372.89-2.636 2.135l-.022.104a2.84 2.84 0 0 0-.036.973l.153 1.112a20.05 20.05 0 0 0-.171.022a1.765 1.765 0 0 0-1.538 1.75v2.307c0 .888.658 1.637 1.538 1.751c1.8.234 3.624.234 5.424 0a1.765 1.765 0 0 0 1.538-1.75v-2.307c0-.888-.658-1.637-1.538-1.751a20.05 20.05 0 0 0-.171-.022l.153-1.112a2.828 2.828 0 0 0-.036-.973l-.022-.104A2.695 2.695 0 0 0 18.5 11.25m1.044 4.196l.164-1.189c.02-.152.015-.306-.017-.456l-.022-.104a1.195 1.195 0 0 0-2.338 0l-.022.104c-.032.15-.038.304-.017.456l.164 1.19a21.079 21.079 0 0 1 2.088 0" clip-rule="evenodd"/><path fill="currentColor" d="M20.692 9.98c.025.169-.161.29-.313.214a4.195 4.195 0 0 0-5.982 2.879l-.022.104c-.08.373-.11.753-.089 1.133a.46.46 0 0 1-.19.394a3.264 3.264 0 0 0-1.346 2.643v1.418a.299.299 0 0 1-.297.3a72.8 72.8 0 0 1-5.605-.183l-1.514-.108a2.128 2.128 0 0 1-1.942-1.74a23.73 23.73 0 0 1-.217-7.095l.273-2.27A2.18 2.18 0 0 1 5.612 5.75h2.291c.993 0 1.797.804 1.797 1.797c0 .033.027.06.06.06h8.712c1.06 0 1.964.77 2.131 1.817l.064.402z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpenOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M20.361 18.58c-.405.39-.943.641-1.536.684l-1.638.117a72.71 72.71 0 0 1-10.374 0l-1.514-.108a2.629 2.629 0 0 1-2.398-2.15a24.23 24.23 0 0 1-.222-7.244L2.95 7.61a2.68 2.68 0 0 1 2.66-2.36h2.292c1.118 0 2.05.798 2.255 1.856h8.314c1.307 0 2.42.95 2.625 2.24l.064.4l.04.254h.335a2.093 2.093 0 0 1 1.951 2.852l-1.25 3.213a5.878 5.878 0 0 1-1.876 2.514m-.745-8.998l.064.401c0 .006.002.011.003.017H10.37a2.75 2.75 0 0 0-2.565 1.757L5.473 17.78l-.068-.005a1.128 1.128 0 0 1-1.03-.922a22.73 22.73 0 0 1-.208-6.796l.273-2.27A1.18 1.18 0 0 1 5.61 6.75h2.292c.44 0 .797.357.797.797c0 .585.474 1.06 1.06 1.06h8.712c.57 0 1.054.413 1.144.975M7.039 17.893a71.29 71.29 0 0 0 10.041-.008l1.638-.118l.195-.018l-.002-.002a4.376 4.376 0 0 0 1.929-2.226l1.25-3.213a.593.593 0 0 0-.554-.808H10.37c-.516 0-.979.317-1.165.799z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpenSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M21.488 10.25a22.174 22.174 0 0 0-.08-.543l-.064-.401a2.908 2.908 0 0 0-2.872-2.45h-8.117A2.548 2.548 0 0 0 7.903 5H5.612a2.93 2.93 0 0 0-2.909 2.58l-.272 2.27a24.48 24.48 0 0 0 .224 7.318a2.878 2.878 0 0 0 2.626 2.354l1.514.108c3.466.248 6.945.248 10.41 0l1.638-.117a2.718 2.718 0 0 0 1.986-1.087a5.876 5.876 0 0 0 1.409-2.111l1.25-3.213a2.093 2.093 0 0 0-1.952-2.852zM5.612 6.5a1.43 1.43 0 0 0-1.42 1.26l-.272 2.268a22.98 22.98 0 0 0 .21 6.87c.113.62.63 1.083 1.258 1.128l.086.006l2.332-6.025a2.75 2.75 0 0 1 2.564-1.757h9.603c-.014-.102-.03-.204-.046-.306l-.064-.402a1.408 1.408 0 0 0-1.39-1.186H9.758a.81.81 0 0 1-.809-.81c0-.577-.469-1.046-1.047-1.046z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M19.602 16.976c.422-2.31.448-4.674.078-6.993l-.064-.401a1.158 1.158 0 0 0-1.144-.976H9.76a1.06 1.06 0 0 1-1.06-1.06a.797.797 0 0 0-.797-.796H5.612c-.597 0-1.1.446-1.171 1.039l-.273 2.269a22.73 22.73 0 0 0 .208 6.796c.093.506.516.886 1.03.922l1.514.109c3.382.242 6.778.242 10.16 0l1.638-.118a.968.968 0 0 0 .884-.791m1.56-7.23a22.229 22.229 0 0 1-.085 7.5a2.468 2.468 0 0 1-2.252 2.018l-1.638.117a72.71 72.71 0 0 1-10.374 0l-1.514-.108a2.629 2.629 0 0 1-2.398-2.15a24.23 24.23 0 0 1-.222-7.244L2.95 7.61a2.68 2.68 0 0 1 2.66-2.36h2.292c1.118 0 2.05.798 2.255 1.856h8.314c1.307 0 2.42.95 2.625 2.24z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderPlusOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.75 13.5a.75.75 0 0 1 .75-.75h1.75V11a.75.75 0 0 1 1.5 0v1.75h1.75a.75.75 0 0 1 0 1.5h-1.75V16a.75.75 0 0 1-1.5 0v-1.75H9.5a.75.75 0 0 1-.75-.75"/><path fill="currentColor" fill-rule="evenodd" d="M21.161 9.747a22.229 22.229 0 0 1-.084 7.498a2.468 2.468 0 0 1-2.252 2.019l-1.638.117a72.71 72.71 0 0 1-10.374 0l-1.514-.108a2.629 2.629 0 0 1-2.398-2.15a24.23 24.23 0 0 1-.222-7.244L2.95 7.61a2.68 2.68 0 0 1 2.66-2.36h2.292c1.118 0 2.05.798 2.255 1.856h8.314c1.307 0 2.42.95 2.625 2.24zm-1.56 7.229c.423-2.31.45-4.674.08-6.993l-.065-.401a1.158 1.158 0 0 0-1.144-.976H9.76a1.06 1.06 0 0 1-1.06-1.06a.797.797 0 0 0-.797-.796H5.612c-.597 0-1.1.446-1.171 1.039l-.273 2.269a22.73 22.73 0 0 0 .208 6.796c.093.506.516.886 1.03.922l1.514.109c3.382.242 6.778.242 10.16 0l1.638-.118a.968.968 0 0 0 .884-.791" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderPlusSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M20.668 9.826c.387 2.43.36 4.909-.083 7.33a1.968 1.968 0 0 1-1.795 1.609l-1.638.117c-3.43.245-6.874.245-10.304 0l-1.514-.108a2.128 2.128 0 0 1-1.942-1.74a23.73 23.73 0 0 1-.217-7.095l.273-2.27A2.18 2.18 0 0 1 5.612 5.75h2.291c.993 0 1.797.804 1.797 1.797c0 .033.027.06.06.06h8.712c1.06 0 1.964.77 2.131 1.817zM12.75 11a.75.75 0 0 0-1.5 0v1.75H9.5a.75.75 0 0 0 0 1.5h1.75V16a.75.75 0 0 0 1.5 0v-1.75h1.75a.75.75 0 0 0 0-1.5h-1.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.585 17.156c.443-2.421.47-4.9.082-7.33l-.064-.402a2.158 2.158 0 0 0-2.13-1.818H9.758a.06.06 0 0 1-.059-.06c0-.992-.804-1.796-1.797-1.796h-2.29a2.18 2.18 0 0 0-2.164 1.92l-.273 2.269a23.73 23.73 0 0 0 .217 7.094a2.128 2.128 0 0 0 1.942 1.74l1.514.11c3.43.245 6.874.245 10.304 0l1.638-.118a1.968 1.968 0 0 0 1.795-1.61"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderUserOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.68 9.983l.059.393a.32.32 0 0 0 .232.26c.357.103.692.262.992.466c.15.102.372-.01.353-.192a22.202 22.202 0 0 0-.155-1.163l-.064-.402a2.658 2.658 0 0 0-2.625-2.239H9.158A2.298 2.298 0 0 0 6.903 5.25H4.612a2.68 2.68 0 0 0-2.66 2.36l-.273 2.27a24.23 24.23 0 0 0 .222 7.243a2.629 2.629 0 0 0 2.398 2.15l1.514.108c2.18.156 4.366.213 6.551.172a.313.313 0 0 0 .293-.235a4.43 4.43 0 0 1 .383-.957c.077-.142-.026-.323-.188-.319c-2.311.06-4.625.008-6.932-.157l-1.514-.109a1.128 1.128 0 0 1-1.03-.922a22.73 22.73 0 0 1-.208-6.796l.273-2.27A1.18 1.18 0 0 1 4.61 6.75h2.292c.44 0 .797.357.797.797c0 .585.474 1.06 1.06 1.06h8.712c.57 0 1.054.413 1.144.975z"/><path fill="currentColor" d="M18 12a2 2 0 1 0 0 4a2 2 0 0 0 0-4m-4 8.5a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3a1 1 0 0 1-1 1h-6a1 1 0 0 1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderUserSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.188 16.084a.224.224 0 0 1-.084.334a4.516 4.516 0 0 0-2.265 2.366a.422.422 0 0 1-.377.267a72.174 72.174 0 0 1-6.614-.169l-1.514-.108a2.128 2.128 0 0 1-1.942-1.74a23.73 23.73 0 0 1-.217-7.095l.273-2.27A2.18 2.18 0 0 1 4.612 5.75h2.291c.993 0 1.797.804 1.797 1.797c0 .033.027.06.06.06h8.712c1.06 0 1.964.77 2.131 1.817l.064.402c.042.26.079.522.111.784c.02.165-.16.28-.31.211a3.5 3.5 0 0 0-4.28 5.262"/><path fill="currentColor" d="M18 12a2 2 0 1 0 0 4a2 2 0 0 0 0-4m-4 8.5a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3a1 1 0 0 1-1 1h-6a1 1 0 0 1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.46 16.82a.75.75 0 0 1-.13-.86l.702-1.386a9.75 9.75 0 0 1 8.699-5.345h.345a59.47 59.47 0 0 1 .097-1.638l.068-.931a1.002 1.002 0 0 1 1.539-.771a19.632 19.632 0 0 1 5.373 5.089l.456.635a.75.75 0 0 1 0 .875l-.456.635a19.63 19.63 0 0 1-5.373 5.089a1.002 1.002 0 0 1-1.539-.771l-.068-.93a61.5 61.5 0 0 1-.111-1.957a13.998 13.998 0 0 0-6.27 1.282L3.314 16.98a.75.75 0 0 1-.854-.16m2.218-2.122l.485-.224a15.497 15.497 0 0 1 7.682-1.38a.75.75 0 0 1 .692.725c.025.861.07 1.722.132 2.582l.006.075a18.13 18.13 0 0 0 4.26-4.228l.142-.198l-.142-.197a18.13 18.13 0 0 0-4.26-4.228l-.006.075a59.758 59.758 0 0 0-.123 2.304a.75.75 0 0 1-.75.725h-1.065a8.25 8.25 0 0 0-7.053 3.969" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.554 16.06a.5.5 0 0 0 .656.68l2.499-1.153c2.1-.97 4.393-1.413 6.681-1.309c.027.735.068 1.47.121 2.204l.069.938a.754.754 0 0 0 1.158.581a19.55 19.55 0 0 0 5.351-5.068l.46-.64a.5.5 0 0 0 0-.584l-.46-.64A19.55 19.55 0 0 0 13.738 6a.754.754 0 0 0-1.158.58l-.069.94a61.91 61.91 0 0 0-.108 1.89h-.644a9.5 9.5 0 0 0-8.475 5.209z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GamepadOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 7.25a.75.75 0 0 1 .75.75v1.25H10a.75.75 0 0 1 0 1.5H8.75V12a.75.75 0 0 1-1.5 0v-1.25H6a.75.75 0 0 1 0-1.5h1.25V8A.75.75 0 0 1 8 7.25m9 4.5a1 1 0 1 0-2 0V12a1 1 0 1 0 2 0zM16 7a1 1 0 0 1 1 1v.25a1 1 0 1 1-2 0V8a1 1 0 0 1 1-1m3 3a1 1 0 0 1-1 1h-.25a1 1 0 1 1 0-2H18a1 1 0 0 1 1 1m-4.75 1a1 1 0 1 0 0-2H14a1 1 0 1 0 0 2z"/><path fill="currentColor" fill-rule="evenodd" d="M12.75 3a.75.75 0 0 0-1.5 0v1.487c-1.44.02-2.878.096-4.312.229l-.332.03a4.71 4.71 0 0 0-4.274 4.61l-.012.713a42.94 42.94 0 0 0 .672 8.328a2.86 2.86 0 0 0 2.814 2.353h.221a3.124 3.124 0 0 0 2.856-1.857l1.06-2.392c.791-1.783 3.323-1.783 4.114 0l1.06 2.392a3.124 3.124 0 0 0 2.856 1.857h.221a2.86 2.86 0 0 0 2.814-2.353a42.9 42.9 0 0 0 .672-8.328l-.012-.713a4.71 4.71 0 0 0-4.274-4.61l-.332-.03a54.878 54.878 0 0 0-4.312-.23zM7.076 6.21a53.385 53.385 0 0 1 9.848 0l.332.03a3.209 3.209 0 0 1 2.912 3.141l.012.713a41.437 41.437 0 0 1-.648 8.037a1.36 1.36 0 0 1-1.338 1.119h-.221a1.624 1.624 0 0 1-1.484-.965l-1.061-2.392c-1.319-2.972-5.537-2.972-6.856 0l-1.06 2.392c-.26.587-.843.965-1.485.965h-.22a1.36 1.36 0 0 1-1.339-1.119a41.44 41.44 0 0 1-.648-8.036l.012-.714A3.21 3.21 0 0 1 6.745 6.24z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GamepadSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25a.75.75 0 0 1 .75.75v1.737c1.431.02 2.862.095 4.29.228l.33.03a4.46 4.46 0 0 1 4.048 4.365l.012.713a42.69 42.69 0 0 1-.668 8.28a2.61 2.61 0 0 1-2.568 2.147h-.221a2.874 2.874 0 0 1-2.627-1.709l-1.06-2.39c-.88-1.983-3.692-1.983-4.571 0l-1.06 2.39A2.874 2.874 0 0 1 6.026 20.5h-.22a2.61 2.61 0 0 1-2.57-2.148a42.689 42.689 0 0 1-.667-8.279l.012-.713A4.46 4.46 0 0 1 6.63 4.995l.332-.03a54.636 54.636 0 0 1 4.289-.228V3a.75.75 0 0 1 .75-.75M17 8a1 1 0 1 0-2 0v.25a1 1 0 1 0 2 0zm-9-.75a.75.75 0 0 1 .75.75v1.25H10a.75.75 0 0 1 0 1.5H8.75V12a.75.75 0 0 1-1.5 0v-1.25H6a.75.75 0 0 1 0-1.5h1.25V8A.75.75 0 0 1 8 7.25m8 3.5a1 1 0 0 1 1 1V12a1 1 0 1 1-2 0v-.25a1 1 0 0 1 1-1m2 .25a1 1 0 1 0 0-2h-.25a1 1 0 1 0 0 2zm-2.75-1a1 1 0 0 1-1 1H14a1 1 0 1 1 0-2h.25a1 1 0 0 1 1 1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.955 2.045a.75.75 0 0 1 0 1.06l-.55.551c3.288 3.831 3.118 9.61-.51 13.238a9.712 9.712 0 0 1-6.145 2.827v1.53h1.75a.75.75 0 1 1 0 1.5h-5a.75.75 0 0 1 0-1.5h1.75v-1.53a9.709 9.709 0 0 1-5.594-2.316l-.55.55a.75.75 0 0 1-1.061-1.06l.707-.708A1.272 1.272 0 0 1 5.5 16.15a8.222 8.222 0 0 0 5.497 2.101h.006A8.25 8.25 0 0 0 17.15 4.5a1.272 1.272 0 0 1 .038-1.748l.707-.707a.75.75 0 0 1 1.061 0"/><path fill="currentColor" fill-rule="evenodd" d="M4.25 10a6.75 6.75 0 1 1 13.5 0a6.75 6.75 0 0 1-13.5 0M11 4.75a5.25 5.25 0 0 0-5.159 6.23a2.25 2.25 0 1 1 2.56 3.583a5.252 5.252 0 0 0 7.746-3.522a3.75 3.75 0 1 1-2.89-5.782A5.23 5.23 0 0 0 11 4.75M10.75 9a2.25 2.25 0 1 1 4.5 0a2.25 2.25 0 0 1-4.5 0M7.5 11.75a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.955 3.106a.75.75 0 0 0-1.06-1.06l-.708.706A1.272 1.272 0 0 0 17.15 4.5a8.25 8.25 0 0 1-6.146 13.75h-.006c-1.968 0-3.935-.7-5.497-2.1a1.272 1.272 0 0 0-1.748.037l-.707.707a.75.75 0 1 0 1.06 1.061l.551-.55a9.709 9.709 0 0 0 5.594 2.316v1.53H8.5a.75.75 0 1 0 0 1.5h5a.75.75 0 0 0 0-1.5h-1.75v-1.53a9.712 9.712 0 0 0 6.144-2.827c3.63-3.629 3.8-9.407.51-13.238z"/><path fill="currentColor" fill-rule="evenodd" d="M4.5 10a6.5 6.5 0 1 1 13 0a6.5 6.5 0 0 1-13 0M13 6.25a2.75 2.75 0 1 0 0 5.5a2.75 2.75 0 0 0 0-5.5M6.25 12.5a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GmailOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.856 6.84a.75.75 0 0 0-1.106.66V17a.75.75 0 0 0 1.5 0V8.756l5.394 2.904c.222.12.49.12.712 0l5.394-2.904V17a.75.75 0 0 0 1.5 0V7.5a.75.75 0 0 0-1.106-.66L12 10.148z"/><path fill="currentColor" fill-rule="evenodd" d="M17.31 3.722a59.632 59.632 0 0 0-10.62 0l-1.518.135a3.53 3.53 0 0 0-3.179 3.006a35.508 35.508 0 0 0 0 10.274a3.53 3.53 0 0 0 3.18 3.005l1.516.136c3.534.316 7.088.316 10.622 0l1.517-.136a3.53 3.53 0 0 0 3.179-3.005a35.508 35.508 0 0 0 0-10.274a3.53 3.53 0 0 0-3.18-3.006zM6.824 5.216a58.133 58.133 0 0 1 10.354 0l1.517.136a2.03 2.03 0 0 1 1.829 1.728a34.005 34.005 0 0 1 0 9.84a2.03 2.03 0 0 1-1.829 1.728l-1.517.136c-3.444.308-6.91.308-10.354 0l-1.517-.136a2.03 2.03 0 0 1-1.829-1.728a34.008 34.008 0 0 1 0-9.84a2.03 2.03 0 0 1 1.829-1.728z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GmailSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.712 3.97a59.38 59.38 0 0 1 10.576 0l1.518.136A3.28 3.28 0 0 1 21.76 6.9a35.257 35.257 0 0 1 0 10.2a3.28 3.28 0 0 1-2.954 2.793l-1.518.136a59.38 59.38 0 0 1-10.576 0l-1.518-.136A3.28 3.28 0 0 1 2.24 17.1a35.257 35.257 0 0 1 0-10.2a3.28 3.28 0 0 1 2.954-2.794zm-.856 2.87a.75.75 0 0 0-1.106.66V17a.75.75 0 0 0 1.5 0V8.756l5.394 2.904c.222.12.49.12.712 0l5.394-2.904V17a.75.75 0 0 0 1.5 0V7.5a.75.75 0 0 0-1.106-.66L12 10.148z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleAltOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 12c0-5.384 4.366-9.75 9.75-9.75a9.71 9.71 0 0 1 6.64 2.623a.75.75 0 0 1 .018 1.08l-2.545 2.545a.75.75 0 0 1-1.029.03A4.621 4.621 0 0 0 12 7.35a4.65 4.65 0 0 0 0 9.3a4.637 4.637 0 0 0 3.883-2.1H12a.75.75 0 0 1-.75-.75v-3.6a.75.75 0 0 1 .75-.75h8.825a.75.75 0 0 1 .736.604A10.2 10.2 0 0 1 21.75 12c0 5.384-4.366 9.75-9.75 9.75S2.25 17.384 2.25 12M12 3.75a8.25 8.25 0 1 0 8.185 7.2H12.75v2.1h4.336a.75.75 0 0 1 .707 1A6.148 6.148 0 0 1 5.85 12A6.15 6.15 0 0 1 12 5.85c1.313 0 2.527.415 3.524 1.116l1.499-1.5A8.187 8.187 0 0 0 12 3.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleAltSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.855 10.361a.197.197 0 0 0-.194-.161H12.2a.2.2 0 0 0-.2.2v3.2c0 .11.09.2.2.2h4.886A5.398 5.398 0 0 1 6.6 12A5.4 5.4 0 0 1 12 6.6a5.37 5.37 0 0 1 3.44 1.245a.205.205 0 0 0 .276-.01l2.266-2.267a.197.197 0 0 0-.007-.286A8.953 8.953 0 0 0 12 3a9 9 0 1 0 9 9c0-.547-.051-1.113-.145-1.639"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleDriveOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.994 3.124a.75.75 0 0 1 .65-.374h6.713a.75.75 0 0 1 .649.374l6.643 11.482a.75.75 0 0 1-.01.77l-3.41 5.518a.75.75 0 0 1-.638.356H5.41a.75.75 0 0 1-.638-.356l-3.41-5.518a.75.75 0 0 1-.01-.77zm.651 1.87L2.874 14.97l2.518 4.077l2.603-4.441l.001-.003l3.119-5.361zm3.338 5.739l-2.036 3.499h4.07zm3.77 3.499H20.7L14.924 4.25H9.947zm4.902 1.5H9.073L6.718 19.75h11.455z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleDriveSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.243 21.32a.3.3 0 0 1-.26-.45l3.194-5.507a.3.3 0 0 1 .259-.149l11.62-.001a.3.3 0 0 1 .26.45l-3.178 5.507a.3.3 0 0 1-.26.15zm.094-16.348a.3.3 0 0 1 .52 0l3.185 5.493a.3.3 0 0 1 0 .3L5.21 20.87a.3.3 0 0 1-.52 0l-3.178-5.506a.3.3 0 0 1 0-.3zM21.47 13.3a.3.3 0 0 1-.26.45h-6.355a.3.3 0 0 1-.26-.15L8.736 3.45a.3.3 0 0 1 .26-.45h6.355a.3.3 0 0 1 .26.15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25c-5.384 0-9.75 4.366-9.75 9.75s4.366 9.75 9.75 9.75s9.75-4.366 9.75-9.75a10.2 10.2 0 0 0-.19-1.946a.75.75 0 0 0-.735-.604H12a.75.75 0 0 0-.75.75v3.6c0 .414.336.75.75.75h3.883a4.68 4.68 0 0 1-1.05 1.132h-.002l.002.002A4.616 4.616 0 0 1 12 16.65a4.65 4.65 0 0 1 0-9.3c1.184 0 2.26.446 3.084 1.178a.75.75 0 0 0 1.029-.03l2.545-2.546a.75.75 0 0 0-.019-1.079A9.708 9.708 0 0 0 12 2.25m3.202 15a6.148 6.148 0 0 1-8.573-2.252L4.952 16.29A8.246 8.246 0 0 0 12 20.25c1.821 0 3.505-.59 4.87-1.59zm2.803.406A8.221 8.221 0 0 0 20.25 12c0-.348-.023-.704-.065-1.05H12.75v2.1h4.336a.75.75 0 0 1 .707 1a6.166 6.166 0 0 1-1.397 2.245zM5.02 7.601A8.244 8.244 0 0 1 12 3.75c1.892 0 3.63.642 5.023 1.716l-1.5 1.5A6.105 6.105 0 0 0 12 5.85a6.147 6.147 0 0 0-5.281 2.997zm-.69 1.354A8.227 8.227 0 0 0 3.75 12a8.23 8.23 0 0 0 .53 2.914l1.767-1.362A6.16 6.16 0 0 1 5.85 12c0-.606.088-1.192.251-1.746z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePlayOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.576 5.56a2515.203 2515.203 0 0 0-5.79-3.118a1.858 1.858 0 0 0-1.28-.214a1.54 1.54 0 0 0-.915.61c-.235.328-.341.735-.341 1.18V19.84c0 .388.076.797.288 1.14c.228.369.6.636 1.083.68a.747.747 0 0 0 .17-.003c.35.001.682-.111.99-.279l.11-.059l.615-.333a5749.164 5749.164 0 0 0 12.265-6.648l.85-.463a447.065 447.065 0 0 1 1.085-.587l.073-.038l.026-.014l.018-.01c.321-.168.765-.528.893-1.074a1.256 1.256 0 0 0-.149-.942c-.17-.286-.437-.5-.741-.661l-.008-.004a497.726 497.726 0 0 1-3.917-2.125a.752.752 0 0 0-.128-.055l-1.523-.823zm2.131 5.29L8.311 5.505l2.553 1.377l3.673 1.982l.753.406zM4.75 19.784V4.086l7.896 7.824zm10.026-7.877l1.907-1.896c1.787.975 3.152 1.716 3.441 1.866l.022.012a.408.408 0 0 1-.018.01c-.155.08-.974.525-2.147 1.163l-1.309.712zm-6.367 6.346l5.304-5.287l1.581 1.557a5153.142 5153.142 0 0 1-6.885 3.73" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePlaySolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.974 19.974c-.082.081.016.214.117.159c2.589-1.4 6.083-3.291 9.036-4.895a.1.1 0 0 0 .022-.16l-2.08-2.048a.1.1 0 0 0-.142 0zm6.947-9.198a.1.1 0 0 0 .141 0l2.082-2.08a.1.1 0 0 0-.023-.159l-1.02-.55a33413.245 33413.245 0 0 0-8.159-4.402c-.101-.055-.2.078-.118.16zM3 3.196c0-.232.035-.427.096-.587c.063-.165.272-.178.398-.053l9.365 9.279a.1.1 0 0 1 0 .142l-9.399 9.37c-.138.139-.367.105-.412-.085A2.022 2.022 0 0 1 3 20.805zm14.688 11.193a.1.1 0 0 1-.118-.017l-2.438-2.4a.1.1 0 0 1 0-.142l2.451-2.437a.1.1 0 0 1 .119-.016a629.14 629.14 0 0 0 4.21 2.284a.09.09 0 0 1 .041.043a.55.55 0 0 1 .04.308a.09.09 0 0 1-.049.067c-.285.15-1.24.67-2.572 1.394l-.008.004z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.66 10.2c.096 0 .179.067.195.161c.094.526.145 1.092.145 1.639a8.967 8.967 0 0 1-2.293 6.001a.197.197 0 0 1-.274.018l-2.445-2.07a.206.206 0 0 1-.016-.297a5.398 5.398 0 0 0 1.114-1.852H12.2a.2.2 0 0 1-.2-.2v-3.2c0-.11.09-.2.2-.2zm-6.187 6.6a.205.205 0 0 1 .226.024l2.568 2.173a.196.196 0 0 1-.01.309A8.959 8.959 0 0 1 12 21a8.993 8.993 0 0 1-7.548-4.097a.197.197 0 0 1 .046-.263l2.545-1.962a.207.207 0 0 1 .303.062a5.398 5.398 0 0 0 7.127 2.06M6.68 12.926a.205.205 0 0 1-.076.197L3.869 15.23a.196.196 0 0 1-.304-.084A8.98 8.98 0 0 1 3 12c0-1.152.217-2.254.612-3.267a.196.196 0 0 1 .299-.085l2.732 2.004c.065.047.095.13.078.208a5.419 5.419 0 0 0-.042 2.066m.468-3.765c.096.07.231.042.295-.058A5.396 5.396 0 0 1 12 6.6a5.37 5.37 0 0 1 3.44 1.245a.205.205 0 0 0 .276-.01l2.266-2.267a.197.197 0 0 0-.007-.286A8.953 8.953 0 0 0 12 3a8.992 8.992 0 0 0-7.484 4a.197.197 0 0 0 .049.267z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GroupOneHundredFiftyOneOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.25 10.215c0-3.969 3.805-6.965 8.202-6.965c3.498 0 6.6 1.884 7.74 4.65a.198.198 0 0 1-.172.269c-.372.03-.742.084-1.106.158a.209.209 0 0 1-.229-.114c-.963-1.979-3.33-3.463-6.233-3.463c-3.835 0-6.702 2.568-6.702 5.464c0 1.565.814 3.01 2.183 4.031a.75.75 0 0 1 .3.602v.94l1.549-.43a.75.75 0 0 1 .411.003c.61.179 1.264.287 1.945.313a.208.208 0 0 1 .2.2c.018.427.075.84.167 1.238a.056.056 0 0 1-.053.068a9.566 9.566 0 0 1-2.47-.32l-2.297.638a.75.75 0 0 1-.951-.723v-1.563C1.224 13.965.25 12.201.25 10.214"/><path fill="currentColor" d="M6.855 8.258a1.065 1.065 0 1 1-2.13 0a1.065 1.065 0 0 1 2.13 0m4.258 1.065a1.065 1.065 0 1 0 0-2.13a1.065 1.065 0 0 0 0 2.13m3.433 5.854a.931.931 0 1 0 0-1.862a.931.931 0 0 0 0 1.862m5.456-.931a.931.931 0 1 1-1.863 0a.931.931 0 0 1 1.863 0"/><path fill="currentColor" fill-rule="evenodd" d="M16.786 9.859c3.71 0 6.96 2.532 6.96 5.928c0 1.68-.81 2.945-2.066 4.001v1.244a.75.75 0 0 1-.95.723l-1.882-.523a8.101 8.101 0 0 1-2.058.262c-3.71 0-6.96-2.532-6.96-5.928s3.245-5.707 6.956-5.707m5.464 5.707c0-2.324-2.311-4.429-5.46-4.429c-3.148 0-5.46 2.105-5.46 4.429c0 2.324 2.312 4.428 5.46 4.428a6.59 6.59 0 0 0 1.847-.26a.75.75 0 0 1 .412-.003l1.131.314v-.62a.75.75 0 0 1 .302-.6c1.114-.832 1.768-2.001 1.768-3.26" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GroupTwentyThreeSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.4 5.25c-2.78 0-5.15 2.08-5.15 4.78c0 1.863.872 3.431 2.028 4.73c1.153 1.295 2.64 2.382 3.983 3.292l2.319 1.57a.75.75 0 0 0 .84 0l2.319-1.57c1.344-.91 2.83-1.997 3.982-3.292c1.157-1.299 2.029-2.867 2.029-4.73c0-2.7-2.37-4.78-5.15-4.78c-1.434 0-2.695.672-3.6 1.542c-.905-.87-2.167-1.542-3.6-1.542m.35 7.5A.75.75 0 0 1 9.5 12h1.75v-1.75a.75.75 0 0 1 1.5 0V12h1.75a.75.75 0 0 1 0 1.5h-1.75v1.75a.75.75 0 0 1-1.5 0V13.5H9.5a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadphoneOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 3.75C8.436 3.75 5.75 6.205 5.75 9v1.261c.082-.007.166-.011.25-.011h3a.75.75 0 0 1 .75.75v7a.75.75 0 0 1-.75.75H6A2.75 2.75 0 0 1 3.25 16v-3c0-.854.39-1.617 1-2.121V9c0-3.832 3.582-6.75 7.75-6.75S19.75 5.168 19.75 9v1.879c.61.504 1 1.267 1 2.121v3A2.75 2.75 0 0 1 18 18.75h-3a.75.75 0 0 1-.75-.75v-7a.75.75 0 0 1 .75-.75h3c.084 0 .168.004.25.011V9c0-2.795-2.686-5.25-6.25-5.25m-6 8c-.69 0-1.25.56-1.25 1.25v3c0 .69.56 1.25 1.25 1.25h2.25v-5.5zM19.25 13c0-.69-.56-1.25-1.25-1.25h-2.25v5.5H18c.69 0 1.25-.56 1.25-1.25z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadphoneSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3.75C8.436 3.75 5.75 6.205 5.75 9v1.512A2.55 2.55 0 0 1 6 10.5h2A1.5 1.5 0 0 1 9.5 12v5A1.5 1.5 0 0 1 8 18.5H6A2.5 2.5 0 0 1 3.5 16v-3c0-.7.287-1.332.75-1.785V9c0-3.832 3.582-6.75 7.75-6.75S19.75 5.168 19.75 9v2.215A2.49 2.49 0 0 1 20.5 13v3a2.5 2.5 0 0 1-2.5 2.5h-2a1.5 1.5 0 0 1-1.5-1.5v-5a1.5 1.5 0 0 1 1.5-1.5h2c.084 0 .168.004.25.012V9c0-2.795-2.686-5.25-6.25-5.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadsetOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 3.75c-3.241 0-5.756 2.03-6.185 4.5H8a.75.75 0 0 1 .75.75v7a.75.75 0 0 1-.75.75H5A2.75 2.75 0 0 1 2.25 14v-3a2.75 2.75 0 0 1 2.035-2.656C4.667 4.84 8.074 2.25 12 2.25c3.926 0 7.333 2.59 7.715 6.094A2.751 2.751 0 0 1 21.75 11v3a2.751 2.751 0 0 1-2.045 2.659A4.751 4.751 0 0 1 15 20.75h-1.145a2 2 0 1 1 0-1.5H15a3.251 3.251 0 0 0 3.163-2.5H16a.75.75 0 0 1-.75-.75V9a.75.75 0 0 1 .75-.75h2.185c-.429-2.47-2.944-4.5-6.185-4.5m-7 6c-.69 0-1.25.56-1.25 1.25v3c0 .69.56 1.25 1.25 1.25h2.25v-5.5zM20.25 11c0-.69-.56-1.25-1.25-1.25h-2.25v5.5H19c.69 0 1.25-.56 1.25-1.25z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadsetSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.285 8.344A2.751 2.751 0 0 0 2.25 11v3A2.75 2.75 0 0 0 5 16.75h2.5a.75.75 0 0 0 .75-.75V9a.75.75 0 0 0-.75-.75H5.815c.429-2.47 2.944-4.5 6.185-4.5c3.241 0 5.756 2.03 6.185 4.5H16.5a.75.75 0 0 0-.75.75v7c0 .414.336.75.75.75h1.663A3.251 3.251 0 0 1 15 19.25h-1.145a2 2 0 1 0 0 1.5H15c2.4 0 4.384-1.78 4.705-4.091A2.751 2.751 0 0 0 21.75 14v-3a2.751 2.751 0 0 0-2.035-2.656C19.333 4.84 15.926 2.25 12 2.25c-3.926 0-7.333 2.59-7.715 6.094"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartOffOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.22 5.599c.116-.116.052-.31-.11-.326a5.55 5.55 0 0 0-.51-.023c-1.434 0-2.695.672-3.6 1.542c-.905-.87-2.167-1.542-3.6-1.542c-2.78 0-5.15 2.08-5.15 4.78c0 1.85.789 3.476 1.882 4.852c.201.253.414.5.636.738a.295.295 0 0 0 .425.006l.637-.637a.305.305 0 0 0 .006-.422a11.263 11.263 0 0 1-.53-.619c-.96-1.208-1.556-2.515-1.556-3.918c0-1.75 1.57-3.28 3.65-3.28c1.194 0 2.31.713 3.005 1.619a.75.75 0 0 0 1.19 0c.556-.724 1.38-1.325 2.297-1.537a.428.428 0 0 0 .207-.112z"/><path fill="currentColor" fill-rule="evenodd" d="m4.503 19.436l-.033.034a.75.75 0 1 0 1.06 1.06l2.718-2.717c.256.192.513.376.769.553c.521.362 1.05.695 1.533.941c.454.231.975.443 1.45.443s.996-.212 1.45-.443a13.9 13.9 0 0 0 1.534-.941c1.366-.947 2.793-2.112 3.884-3.484c1.093-1.376 1.882-3.002 1.882-4.852c0-1.296-.546-2.45-1.421-3.298L21.53 4.53a.75.75 0 0 0-1.033-1.086L4.53 19.41a.717.717 0 0 1-.027.026M18.268 7.793L9.32 16.74c.183.134.367.264.55.392c.496.343.962.634 1.36.838c.429.218.677.279.77.279c.093 0 .341-.061.77-.28a12.35 12.35 0 0 0 1.36-.837c1.307-.906 2.602-1.974 3.564-3.185c.96-1.208 1.556-2.515 1.556-3.918c0-.847-.367-1.642-.982-2.237" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartOffSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.22 5.599c.116-.116.052-.31-.11-.326a5.553 5.553 0 0 0-.51-.023c-1.434 0-2.695.672-3.6 1.542c-.905-.87-2.167-1.542-3.6-1.542c-2.78 0-5.15 2.08-5.15 4.78c0 1.863.872 3.431 2.028 4.73c.214.24.44.473.673.7c.116.112.3.108.414-.006zM4.503 19.436l-.033.034a.75.75 0 1 0 1.06 1.06l2.99-2.99c.25.177.498.347.741.512l2.319 1.57a.75.75 0 0 0 .84 0l2.319-1.57c1.344-.91 2.83-1.997 3.982-3.292c1.157-1.299 2.029-2.867 2.029-4.73c0-1.296-.546-2.45-1.421-3.298L21.53 4.53a.75.75 0 0 0-1.033-1.086L4.53 19.41a.72.72 0 0 1-.027.026"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.25 10.03c0-2.7 2.37-4.78 5.15-4.78c1.433 0 2.695.672 3.6 1.542c.905-.87 2.166-1.542 3.6-1.542c2.78 0 5.15 2.08 5.15 4.78c0 1.85-.789 3.476-1.882 4.852c-1.09 1.372-2.518 2.537-3.884 3.484c-.523.362-1.05.695-1.534.941c-.453.231-.975.443-1.45.443s-.996-.212-1.45-.443a13.795 13.795 0 0 1-1.533-.941c-1.367-.947-2.794-2.112-3.885-3.484C4.039 13.506 3.25 11.88 3.25 10.03M8.4 6.75c-2.08 0-3.65 1.53-3.65 3.28c0 1.403.596 2.71 1.556 3.918c.962 1.21 2.257 2.279 3.565 3.185c.495.343.96.634 1.36.838c.428.218.676.279.769.279c.093 0 .341-.061.77-.28a12.35 12.35 0 0 0 1.36-.837c1.307-.906 2.602-1.974 3.564-3.185c.96-1.208 1.556-2.515 1.556-3.918c0-1.75-1.57-3.28-3.65-3.28c-1.194 0-2.31.713-3.005 1.619a.75.75 0 0 1-1.19 0C10.71 7.463 9.595 6.75 8.4 6.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartPlusOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 9.75a.75.75 0 0 1 .75.75v1.75h1.75a.75.75 0 0 1 0 1.5h-1.75v1.75a.75.75 0 0 1-1.5 0v-1.75H9.5a.75.75 0 0 1 0-1.5h1.75V10.5a.75.75 0 0 1 .75-.75"/><path fill="currentColor" fill-rule="evenodd" d="M8.4 5.25c-2.78 0-5.15 2.08-5.15 4.78c0 1.85.789 3.476 1.882 4.852c1.09 1.372 2.518 2.537 3.885 3.484c.521.362 1.05.695 1.533.941c.454.231.975.443 1.45.443s.996-.212 1.45-.443a13.9 13.9 0 0 0 1.534-.941c1.366-.947 2.793-2.112 3.884-3.484c1.093-1.376 1.882-3.002 1.882-4.852c0-2.7-2.37-4.78-5.15-4.78c-1.434 0-2.695.672-3.6 1.542c-.905-.87-2.167-1.542-3.6-1.542m-3.65 4.78c0-1.75 1.57-3.28 3.65-3.28c1.194 0 2.31.713 3.005 1.619a.75.75 0 0 0 1.19 0c.696-.906 1.81-1.619 3.005-1.619c2.08 0 3.65 1.53 3.65 3.28c0 1.403-.596 2.71-1.556 3.918c-.962 1.21-2.257 2.279-3.565 3.185a12.35 12.35 0 0 1-1.36.838c-.428.218-.676.279-.769.279c-.093 0-.341-.061-.77-.28a12.347 12.347 0 0 1-1.36-.837c-1.307-.906-2.602-1.974-3.564-3.185c-.96-1.208-1.556-2.515-1.556-3.918" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.4 5.25c-2.78 0-5.15 2.08-5.15 4.78c0 1.863.872 3.431 2.028 4.73c1.153 1.295 2.64 2.382 3.983 3.292l2.319 1.57a.75.75 0 0 0 .84 0l2.319-1.57c1.344-.91 2.83-1.997 3.982-3.292c1.157-1.299 2.029-2.867 2.029-4.73c0-2.7-2.37-4.78-5.15-4.78c-1.434 0-2.695.672-3.6 1.542c-.905-.87-2.167-1.542-3.6-1.542"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartbeatOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.25 10.03c0-2.7 2.37-4.78 5.15-4.78c1.433 0 2.695.672 3.6 1.542c.905-.87 2.166-1.542 3.6-1.542c2.78 0 5.15 2.08 5.15 4.78c0 .162-.006.323-.018.482a.29.29 0 0 1-.293.266h-.897a.316.316 0 0 1-.308-.347c.01-.133.016-.266.016-.401c0-1.75-1.57-3.28-3.65-3.28c-1.194 0-2.31.713-3.005 1.619a.75.75 0 0 1-1.19 0C9.71 7.463 8.595 6.75 7.4 6.75c-2.08 0-3.65 1.53-3.65 3.28c0 1.403.596 2.71 1.556 3.918c.962 1.21 2.257 2.279 3.565 3.185c.495.343.96.634 1.36.838c.428.218.676.279.769.279c.093 0 .341-.061.77-.28a12.35 12.35 0 0 0 1.36-.837c.425-.295.85-.607 1.262-.936a.304.304 0 0 1 .454.084l.451.789a.297.297 0 0 1-.072.38c-.412.326-.83.63-1.242.916a13.74 13.74 0 0 1-1.533.941c-.453.231-.975.443-1.45.443s-.996-.212-1.45-.443a13.79 13.79 0 0 1-1.533-.941c-1.367-.947-2.794-2.112-3.885-3.484C3.039 13.506 2.25 11.88 2.25 10.03"/><path fill="currentColor" d="M14.807 9.688a.75.75 0 0 0-1.26-.165l-2.067 2.505H10a.75.75 0 0 0 0 1.5h1.833a.75.75 0 0 0 .579-.273l1.53-1.855l1.793 3.912a.75.75 0 0 0 1.26.165l1.608-1.95H21a.75.75 0 0 0 0-1.5h-2.75a.75.75 0 0 0-.578.274L16.6 13.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartbeatSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.4 5.25c-2.78 0-5.15 2.08-5.15 4.78c0 1.863.872 3.431 2.028 4.73c1.153 1.295 2.64 2.382 3.983 3.292l2.319 1.57a.75.75 0 0 0 .84 0l2.319-1.57c.43-.291.874-.6 1.317-.928a.298.298 0 0 0 .09-.372L13.781 14a.2.2 0 0 0-.333-.039l-.073.088a2 2 0 0 1-1.543.728H10a2 2 0 0 1 0-4h.702a.4.4 0 0 0 .309-.146l1.571-1.905a2 2 0 0 1 3.361.44l.803 1.751c.084.184.324.233.498.131a1.994 1.994 0 0 1 1.006-.271h1.188a.29.29 0 0 0 .292-.263c.013-.16.02-.321.02-.485c0-2.7-2.37-4.78-5.15-4.78c-1.434 0-2.695.672-3.6 1.542c-.905-.87-2.167-1.542-3.6-1.542"/><path fill="currentColor" d="M14.807 9.688a.75.75 0 0 0-1.26-.165l-2.067 2.505H10a.75.75 0 0 0 0 1.5h1.833a.75.75 0 0 0 .579-.273l1.53-1.855l1.793 3.912a.75.75 0 0 0 1.26.165l1.608-1.95H21a.75.75 0 0 0 0-1.5h-2.75a.75.75 0 0 0-.578.274L16.6 13.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HistoryOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.865 6.882A7.25 7.25 0 1 1 4.75 12a.75.75 0 0 0-1.5 0a8.75 8.75 0 1 0 2.552-6.176a.756.756 0 0 0-.07.08L4.475 4.646a.5.5 0 0 0-.852.309L3.27 8.844a.5.5 0 0 0 .543.543l3.89-.354a.5.5 0 0 0 .307-.851L6.782 6.954a.757.757 0 0 0 .083-.072"/><path fill="currentColor" d="M12.75 7a.75.75 0 0 0-1.5 0v5a.75.75 0 0 0 .352.636l3 1.875a.75.75 0 1 0 .796-1.272l-2.648-1.655z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HistorySolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.865 6.882A7.25 7.25 0 1 1 4.75 12a.75.75 0 0 0-1.5 0a8.75 8.75 0 1 0 2.552-6.176a.756.756 0 0 0-.07.08L4.475 4.646a.5.5 0 0 0-.852.309L3.27 8.844a.5.5 0 0 0 .543.543l3.89-.354a.5.5 0 0 0 .307-.851L6.782 6.954a.757.757 0 0 0 .083-.072"/><path fill="currentColor" d="M12.75 7a.75.75 0 0 0-1.5 0v5a.75.75 0 0 0 .352.636l3 1.875a.75.75 0 1 0 .796-1.272l-2.648-1.655z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.558 5.534a2.25 2.25 0 0 0-3.116 0l-4.626 4.44a.748.748 0 0 0-.218.405a27.344 27.344 0 0 0-.121 9.15l.112.721h2.977v-6.211a.75.75 0 0 1 .75-.75h5.368a.75.75 0 0 1 .75.75v6.211h2.977l.112-.72a27.34 27.34 0 0 0-.12-9.151a.748.748 0 0 0-.219-.405zM9.404 4.452a3.75 3.75 0 0 1 5.192 0l4.627 4.44c.34.326.57.752.655 1.216a28.86 28.86 0 0 1 .127 9.653l-.18 1.157a.983.983 0 0 1-.972.832h-4.169a.75.75 0 0 1-.75-.75v-6.211h-3.868V21a.75.75 0 0 1-.75.75H5.147a.983.983 0 0 1-.972-.832l-.18-1.157a28.86 28.86 0 0 1 .127-9.653c.085-.464.315-.89.655-1.217z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.796 4.136a2.5 2.5 0 0 0-3.592 0L5.405 9.092c-.275.284-.46.644-.532 1.034a28.756 28.756 0 0 0-.127 9.624l.176 1.13c.056.357.364.62.725.62H9a.5.5 0 0 0 .5-.5v-7h5v7a.5.5 0 0 0 .5.5h3.353a.733.733 0 0 0 .724-.62l.177-1.13a28.759 28.759 0 0 0-.127-9.624a2.007 2.007 0 0 0-.533-1.034z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HotspotOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M12 11.65a1.25 1.25 0 1 0 0 2.5a1.25 1.25 0 0 0 0-2.5M9.25 12.9a2.75 2.75 0 1 1 5.5 0a2.75 2.75 0 0 1-5.5 0"/><path d="M12 7.65a5.25 5.25 0 0 0-3.712 8.962a.75.75 0 1 1-1.061 1.06a6.75 6.75 0 1 1 9.546 0a.75.75 0 0 1-1.06-1.06A5.25 5.25 0 0 0 12 7.649"/><path d="M12 3.75a9.15 9.15 0 0 0-6.47 15.62a.75.75 0 1 1-1.06 1.06a10.65 10.65 0 1 1 15.06 0a.75.75 0 0 1-1.06-1.06A9.15 9.15 0 0 0 12 3.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HotspotSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M12 11.65a1.25 1.25 0 1 0 0 2.5a1.25 1.25 0 0 0 0-2.5M9.25 12.9a2.75 2.75 0 1 1 5.5 0a2.75 2.75 0 0 1-5.5 0"/><path d="M12 7.65a5.25 5.25 0 0 0-3.712 8.962a.75.75 0 1 1-1.061 1.06a6.75 6.75 0 1 1 9.546 0a.75.75 0 0 1-1.06-1.06A5.25 5.25 0 0 0 12 7.649"/><path d="M12 3.75a9.15 9.15 0 0 0-6.47 15.62a.75.75 0 1 1-1.06 1.06a10.65 10.65 0 1 1 15.06 0a.75.75 0 0 1-1.06-1.06A9.15 9.15 0 0 0 12 3.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 9a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0"/><path fill="currentColor" fill-rule="evenodd" d="M7.268 4.658a54.647 54.647 0 0 1 9.465 0l1.51.132a3.138 3.138 0 0 1 2.831 2.66a30.604 30.604 0 0 1 0 9.1a3.138 3.138 0 0 1-2.831 2.66l-1.51.131c-3.15.274-6.316.274-9.465 0l-1.51-.131a3.138 3.138 0 0 1-2.832-2.66a30.601 30.601 0 0 1 0-9.1a3.138 3.138 0 0 1 2.831-2.66zm9.335 1.495a53.147 53.147 0 0 0-9.206 0l-1.51.131A1.638 1.638 0 0 0 4.41 7.672a29.101 29.101 0 0 0-.311 5.17L7.97 8.97a.75.75 0 0 1 1.09.032l3.672 4.13l2.53-.844a.75.75 0 0 1 .796.21l3.519 3.91a29.101 29.101 0 0 0 .014-8.736a1.638 1.638 0 0 0-1.478-1.388zm2.017 11.435l-3.349-3.721l-2.534.844a.75.75 0 0 1-.798-.213l-3.471-3.905l-4.244 4.243c.049.498.11.996.185 1.491a1.638 1.638 0 0 0 1.478 1.389l1.51.131c3.063.266 6.143.266 9.206 0l1.51-.131c.178-.016.35-.06.507-.128" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.268 4.658a54.647 54.647 0 0 1 9.465 0l1.51.132a3.138 3.138 0 0 1 2.831 2.66a30.604 30.604 0 0 1 0 9.1c-.04.264-.112.517-.212.754c-.066.157-.27.181-.386.055l-4.421-4.864a.75.75 0 0 0-.792-.207l-2.531.844l-3.671-4.13A.75.75 0 0 0 7.97 8.97l-4.914 4.914a.246.246 0 0 1-.422-.159a30.601 30.601 0 0 1 .292-6.276a3.138 3.138 0 0 1 2.831-2.66zM14 9a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0" clip-rule="evenodd"/><path fill="currentColor" d="M2.961 16.1a.249.249 0 0 0-.07.21l.035.24a3.138 3.138 0 0 0 2.831 2.66l1.51.131c3.15.274 6.316.274 9.466 0l1.51-.131a3.13 3.13 0 0 0 1.185-.347c.137-.071.16-.252.056-.366l-4.1-4.51a.25.25 0 0 0-.265-.07l-2.382.794a.75.75 0 0 1-.798-.213l-3.295-3.707a.25.25 0 0 0-.364-.01z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoCircleOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 10.75a.75.75 0 0 1 .75.75v5a.75.75 0 0 1-1.5 0v-5a.75.75 0 0 1 .75-.75M12 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path fill="currentColor" fill-rule="evenodd" d="M3.25 12a8.75 8.75 0 1 1 17.5 0a8.75 8.75 0 0 1-17.5 0M12 4.75a7.25 7.25 0 1 0 0 14.5a7.25 7.25 0 0 0 0-14.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoCircleSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.25 12a8.75 8.75 0 1 1 17.5 0a8.75 8.75 0 0 1-17.5 0M13 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-1 2.75a.75.75 0 0 1 .75.75v5a.75.75 0 0 1-1.5 0v-5a.75.75 0 0 1 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoRectOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 10.75a.75.75 0 0 1 .75.75v5a.75.75 0 1 1-1.5 0v-5a.75.75 0 0 1 .75-.75M12 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path fill="currentColor" fill-rule="evenodd" d="M7.317 3.769a42.502 42.502 0 0 1 9.366 0c1.827.204 3.302 1.642 3.516 3.48c.37 3.156.37 6.346 0 9.503c-.215 1.836-1.69 3.275-3.516 3.48a42.5 42.5 0 0 1-9.366 0c-1.827-.205-3.302-1.644-3.516-3.48a40.903 40.903 0 0 1 0-9.504c.214-1.837 1.69-3.275 3.516-3.48m9.2 1.49a41.001 41.001 0 0 0-9.034 0A2.486 2.486 0 0 0 5.29 7.423a39.402 39.402 0 0 0 0 9.154a2.486 2.486 0 0 0 2.193 2.163c2.977.333 6.057.333 9.034 0a2.486 2.486 0 0 0 2.192-2.163a39.401 39.401 0 0 0 0-9.154a2.486 2.486 0 0 0-2.192-2.164" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoRectSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.317 3.769a42.502 42.502 0 0 1 9.366 0c1.827.204 3.302 1.642 3.516 3.48c.37 3.156.37 6.346 0 9.503c-.215 1.836-1.69 3.275-3.516 3.48a42.5 42.5 0 0 1-9.366 0c-1.827-.205-3.302-1.644-3.516-3.48a40.903 40.903 0 0 1 0-9.504c.214-1.837 1.69-3.275 3.516-3.48M13 7.999A1 1 0 1 1 11 8a1 1 0 0 1 2 0m-1 2.75a.75.75 0 0 1 .75.75v5a.75.75 0 1 1-1.5 0v-5a.75.75 0 0 1 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoTriangleOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 9a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-1 2.75a.75.75 0 0 1 .75.75v5a.75.75 0 1 1-1.5 0v-5a.75.75 0 0 1 .75-.75"/><path fill="currentColor" fill-rule="evenodd" d="M14.27 3.993a2.749 2.749 0 0 0-4.54 0l-.432.632a75.95 75.95 0 0 0-6.944 12.563l-.09.208a2.511 2.511 0 0 0 2.024 3.497a69.43 69.43 0 0 0 15.424 0a2.511 2.511 0 0 0 2.024-3.497l-.09-.208a75.951 75.951 0 0 0-6.944-12.563zm-3.302.846a1.25 1.25 0 0 1 2.064 0l.432.632a74.444 74.444 0 0 1 6.806 12.315l.09.208a1.011 1.011 0 0 1-.814 1.408c-5.015.56-10.077.56-15.092 0a1.011 1.011 0 0 1-.815-1.408l.09-.208a74.45 74.45 0 0 1 6.807-12.315z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoTriangleSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.73 3.993a2.749 2.749 0 0 1 4.54 0l.432.632a75.95 75.95 0 0 1 6.944 12.563l.09.208a2.511 2.511 0 0 1-2.024 3.497a69.412 69.412 0 0 1-15.424 0a2.511 2.511 0 0 1-2.024-3.497l.09-.208A75.95 75.95 0 0 1 9.298 4.625zM13 9a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-1 2.75a.75.75 0 0 1 .75.75v5a.75.75 0 1 1-1.5 0v-5a.75.75 0 0 1 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InstagramOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 7a1 1 0 1 1 2 0a1 1 0 0 1-2 0"/><path fill="currentColor" fill-rule="evenodd" d="M12 7.25a4.75 4.75 0 1 0 0 9.5a4.75 4.75 0 0 0 0-9.5M8.75 12a3.25 3.25 0 1 1 6.5 0a3.25 3.25 0 0 1-6.5 0" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M17.258 2.833a47.721 47.721 0 0 0-10.516 0c-2.012.225-3.637 1.81-3.873 3.832a45.922 45.922 0 0 0 0 10.67c.236 2.022 1.86 3.607 3.873 3.832a47.77 47.77 0 0 0 10.516 0c2.012-.225 3.637-1.81 3.873-3.832a45.914 45.914 0 0 0 0-10.67c-.236-2.022-1.86-3.607-3.873-3.832m-10.35 1.49a46.22 46.22 0 0 1 10.184 0c1.33.15 2.395 1.199 2.55 2.517a44.421 44.421 0 0 1 0 10.32a2.89 2.89 0 0 1-2.55 2.516a46.217 46.217 0 0 1-10.184 0a2.89 2.89 0 0 1-2.55-2.516a44.421 44.421 0 0 1 0-10.32a2.89 2.89 0 0 1 2.55-2.516" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InstagramSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 8.75a3.25 3.25 0 1 0 0 6.5a3.25 3.25 0 0 0 0-6.5"/><path fill="currentColor" fill-rule="evenodd" d="M6.77 3.082a47.472 47.472 0 0 1 10.46 0c1.899.212 3.43 1.707 3.653 3.613a45.67 45.67 0 0 1 0 10.61c-.223 1.906-1.754 3.401-3.652 3.614a47.468 47.468 0 0 1-10.461 0c-1.899-.213-3.43-1.708-3.653-3.613a45.672 45.672 0 0 1 0-10.611C3.34 4.789 4.871 3.294 6.77 3.082M17 6a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-9.75 6a4.75 4.75 0 1 1 9.5 0a4.75 4.75 0 0 1-9.5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InvoiceOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 9.75a.75.75 0 0 0-.75-.75h-6a.75.75 0 0 0 0 1.5h6a.75.75 0 0 0 .75-.75m-1 3a.75.75 0 0 0-.75-.75h-5a.75.75 0 1 0 0 1.5h5a.75.75 0 0 0 .75-.75m.25 2.25a.75.75 0 1 1 0 1.5h-6a.75.75 0 0 1 0-1.5z"/><path fill="currentColor" fill-rule="evenodd" d="M6 21.75h13A2.75 2.75 0 0 0 21.75 19v-5.5a.75.75 0 0 0-.75-.75h-3.25V4.943c0-1.423-1.609-2.251-2.767-1.424l-.175.125a2.26 2.26 0 0 1-2.622-.004a3.77 3.77 0 0 0-4.372 0a2.26 2.26 0 0 1-2.622.004l-.175-.125c-1.158-.827-2.767 0-2.767 1.424V18A3.75 3.75 0 0 0 6 21.75M8.686 4.86a2.27 2.27 0 0 1 2.628 0a3.76 3.76 0 0 0 4.366.005l.175-.125a.25.25 0 0 1 .395.203V19c0 .45.108.875.3 1.25H6A2.25 2.25 0 0 1 3.75 18V4.943a.25.25 0 0 1 .395-.203l.175.125a3.76 3.76 0 0 0 4.366-.005M17.75 19v-4.75h2.5V19a1.25 1.25 0 0 1-2.5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InvoiceSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M19 21.5H6A3.5 3.5 0 0 1 2.5 18V4.943c0-1.067 1.056-1.744 1.985-1.422c.133.046.263.113.387.202l.175.125a2.51 2.51 0 0 0 2.912-.005a3.52 3.52 0 0 1 4.082 0a2.51 2.51 0 0 0 2.912.005l.175-.125c.993-.71 2.372 0 2.372 1.22V12.5H21a.75.75 0 0 1 .75.75v5.5A2.75 2.75 0 0 1 19 21.5M17.75 14v4.75a1.25 1.25 0 0 0 2.5 0V14zM13.5 9.75a.75.75 0 0 0-.75-.75h-6a.75.75 0 0 0 0 1.5h6a.75.75 0 0 0 .75-.75m-1 3a.75.75 0 0 0-.75-.75h-5a.75.75 0 1 0 0 1.5h5a.75.75 0 0 0 .75-.75m.25 2.25a.75.75 0 1 1 0 1.5h-6a.75.75 0 0 1 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 10a2 2 0 1 0 0 4a2 2 0 0 0 0-4"/><path fill="currentColor" fill-rule="evenodd" d="M7 6.25a5.75 5.75 0 1 0 5.05 8.5h3.2V17c0 .414.336.75.75.75h3.5a.75.75 0 0 0 .75-.75v-2.25H22a.75.75 0 0 0 .75-.75v-3A1.75 1.75 0 0 0 21 9.25h-8.95a5.749 5.749 0 0 0-5.05-3M2.75 12a4.25 4.25 0 0 1 8.147-1.7a.75.75 0 0 0 .687.45H21a.25.25 0 0 1 .25.25v2.25H19.5a.75.75 0 0 0-.75.75v2.25h-2V14a.75.75 0 0 0-.75-.75h-4.416a.75.75 0 0 0-.687.45A4.251 4.251 0 0 1 2.75 12" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeySolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.25 12a5.75 5.75 0 0 1 10.8-2.75H21c.966 0 1.75.784 1.75 1.75v2.5a.75.75 0 0 1-.75.75h-2.25V16a.75.75 0 0 1-.75.75h-2.5a.75.75 0 0 1-.75-.75v-1.75h-3.457A5.751 5.751 0 0 1 1.25 12M7 10a2 2 0 1 0 0 4a2 2 0 0 0 0-4" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.93 4.47a17.516 17.516 0 0 0-3.86 0a.666.666 0 0 0-.586.577a16.816 16.816 0 0 0 0 3.907a.666.666 0 0 0 .587.576c1.271.142 2.587.142 3.858 0a.666.666 0 0 0 .587-.576c.152-1.298.152-2.61 0-3.908a.666.666 0 0 0-.587-.576M4.903 2.98a19.017 19.017 0 0 1 4.192 0c.99.11 1.793.89 1.91 1.892a18.322 18.322 0 0 1 0 4.256a2.166 2.166 0 0 1-1.91 1.893c-1.382.154-2.81.154-4.192 0a2.166 2.166 0 0 1-1.91-1.893a18.316 18.316 0 0 1 0-4.256a2.166 2.166 0 0 1 1.91-1.892M8.93 14.47a17.514 17.514 0 0 0-3.86 0a.666.666 0 0 0-.586.576a16.816 16.816 0 0 0 0 3.908a.666.666 0 0 0 .587.576c1.271.142 2.587.142 3.858 0a.666.666 0 0 0 .587-.576c.152-1.298.152-2.61 0-3.908a.666.666 0 0 0-.587-.576m-4.026-1.49a19.023 19.023 0 0 1 4.192 0c.99.11 1.793.89 1.91 1.892a18.322 18.322 0 0 1 0 4.256a2.166 2.166 0 0 1-1.91 1.892c-1.382.155-2.81.155-4.192 0a2.166 2.166 0 0 1-1.91-1.892a18.316 18.316 0 0 1 0-4.256a2.166 2.166 0 0 1 1.91-1.892M18.93 4.47a17.517 17.517 0 0 0-3.86 0a.666.666 0 0 0-.586.577a16.817 16.817 0 0 0 0 3.907a.666.666 0 0 0 .587.576c1.271.142 2.587.142 3.858 0a.666.666 0 0 0 .587-.576c.152-1.298.152-2.61 0-3.908a.666.666 0 0 0-.587-.576m-4.026-1.49a19.017 19.017 0 0 1 4.192 0c.99.11 1.793.89 1.91 1.892a18.322 18.322 0 0 1 0 4.256a2.166 2.166 0 0 1-1.91 1.893c-1.382.154-2.81.154-4.192 0a2.166 2.166 0 0 1-1.91-1.893a18.31 18.31 0 0 1 0-4.256a2.166 2.166 0 0 1 1.91-1.892m4.027 11.49a17.514 17.514 0 0 0-3.86 0a.666.666 0 0 0-.586.576a16.817 16.817 0 0 0 0 3.908a.666.666 0 0 0 .587.576c1.271.142 2.587.142 3.858 0a.666.666 0 0 0 .587-.576c.152-1.298.152-2.61 0-3.908a.666.666 0 0 0-.587-.576m-4.026-1.49a19.023 19.023 0 0 1 4.192 0c.99.11 1.793.89 1.91 1.892a18.322 18.322 0 0 1 0 4.256a2.166 2.166 0 0 1-1.91 1.892c-1.382.155-2.81.155-4.192 0a2.166 2.166 0 0 1-1.91-1.892a18.31 18.31 0 0 1 0-4.256a2.166 2.166 0 0 1 1.91-1.892" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.068 3.228a18.766 18.766 0 0 0-4.136 0a1.916 1.916 0 0 0-1.69 1.673a18.067 18.067 0 0 0 0 4.198a1.916 1.916 0 0 0 1.69 1.673c1.363.152 2.773.152 4.136 0a1.916 1.916 0 0 0 1.69-1.673a18.066 18.066 0 0 0 0-4.198a1.916 1.916 0 0 0-1.69-1.673m0 10a18.765 18.765 0 0 0-4.136 0a1.916 1.916 0 0 0-1.69 1.673a18.067 18.067 0 0 0 0 4.198a1.916 1.916 0 0 0 1.69 1.673c1.363.152 2.773.152 4.136 0a1.916 1.916 0 0 0 1.69-1.673a18.065 18.065 0 0 0 0-4.198a1.916 1.916 0 0 0-1.69-1.673m10-10a18.766 18.766 0 0 0-4.136 0a1.916 1.916 0 0 0-1.69 1.673a18.066 18.066 0 0 0 0 4.198a1.916 1.916 0 0 0 1.69 1.673c1.364.152 2.772.152 4.136 0a1.916 1.916 0 0 0 1.69-1.673a18.066 18.066 0 0 0 0-4.198a1.916 1.916 0 0 0-1.69-1.673m0 10a18.765 18.765 0 0 0-4.136 0a1.916 1.916 0 0 0-1.69 1.673a18.065 18.065 0 0 0 0 4.198a1.916 1.916 0 0 0 1.69 1.673c1.364.152 2.772.152 4.136 0a1.916 1.916 0 0 0 1.69-1.673a18.065 18.065 0 0 0 0-4.198a1.916 1.916 0 0 0-1.69-1.673"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightbulbAltOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 3.75c-3.476 0-6.25 2.717-6.25 6.016c0 2.005.82 3.733 2.285 4.81c.323.237.6.591.705 1.04c.087.369.186.818.284 1.294h2.226v-1.083a.75.75 0 0 1 .117-.403l.802-1.26l-1.799-1.412a.75.75 0 0 1-.17-.993l1.167-1.832a.75.75 0 0 1 1.266.806l-.802 1.259l1.799 1.413a.75.75 0 0 1 .17.992l-1.05 1.649v.864h2.226c.098-.476.197-.925.284-1.294c.106-.449.382-.803.705-1.04c1.464-1.077 2.285-2.806 2.285-4.81c0-3.299-2.774-6.016-6.25-6.016m2.689 14.66H9.31c.11.637.197 1.24.224 1.674c.027.457.368.866.871.974l.196.043c.92.199 1.875.199 2.796 0l.196-.043c.503-.108.844-.517.872-.974c.026-.433.112-1.037.223-1.674M4.25 9.766C4.25 5.59 7.744 2.25 12 2.25s7.75 3.341 7.75 7.516c0 2.424-1.004 4.627-2.897 6.018a.324.324 0 0 0-.133.176a50.74 50.74 0 0 0-.394 1.843c-.183.938-.332 1.848-.363 2.372c-.07 1.158-.922 2.105-2.052 2.35l-.196.042c-1.13.244-2.3.244-3.43 0l-.196-.042c-1.13-.244-1.982-1.192-2.052-2.35c-.031-.524-.18-1.434-.363-2.372a50.745 50.745 0 0 0-.394-1.843a.324.324 0 0 0-.133-.176C5.254 14.394 4.25 12.19 4.25 9.767" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightbulbAltSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 9.766C4.5 5.737 7.874 2.5 12 2.5c4.126 0 7.5 3.237 7.5 7.266c0 2.354-.973 4.478-2.795 5.817a.571.571 0 0 0-.228.32c-.059.25-.124.538-.19.848a.2.2 0 0 1-.197.159h-3.14a.2.2 0 0 1-.2-.2v-.664l1.05-1.648a.75.75 0 0 0-.17-.993l-1.8-1.413l.803-1.26a.75.75 0 0 0-1.266-.805l-1.166 1.832a.75.75 0 0 0 .17.993l1.798 1.413l-.802 1.26a.75.75 0 0 0-.117.402v.882a.2.2 0 0 1-.2.2H7.91a.2.2 0 0 1-.196-.158c-.067-.31-.132-.599-.19-.848a.57.57 0 0 0-.23-.32C5.475 14.243 4.5 12.12 4.5 9.766m3.781 8.644a.201.201 0 0 0-.198.235c.105.604.182 1.147.204 1.515c.063 1.041.83 1.899 1.855 2.12l.196.043a7.877 7.877 0 0 0 3.324 0l.196-.043c1.025-.221 1.792-1.079 1.855-2.12c.022-.368.1-.91.204-1.515a.201.201 0 0 0-.198-.235z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightbulbOffOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3.75c-3.476 0-6.25 2.717-6.25 6.016c0 1.679.576 3.164 1.623 4.235c.12.123.125.32.004.441l-.635.635a.29.29 0 0 1-.412.003c-1.364-1.376-2.08-3.266-2.08-5.314C4.25 5.59 7.744 2.25 12 2.25c1.992 0 3.817.732 5.195 1.939a.29.29 0 0 1 .01.425l-.637.638a.311.311 0 0 1-.423.011A6.356 6.356 0 0 0 12 3.75"/><path fill="currentColor" fill-rule="evenodd" d="M19.997 3.944L4.03 19.91a.714.714 0 0 1-.027.026l-.033.034a.75.75 0 1 0 1.06 1.06l2.737-2.736c.139.756.244 1.45.27 1.881c.07 1.158.922 2.105 2.052 2.35l.196.042c1.13.244 2.3.244 3.43 0l.196-.042c1.13-.244 1.982-1.192 2.052-2.35c.031-.524.18-1.434.363-2.372c.131-.675.276-1.34.394-1.843a.324.324 0 0 1 .133-.176c1.893-1.391 2.897-3.594 2.897-6.018a7.3 7.3 0 0 0-.584-2.87L21.03 5.03a.75.75 0 0 0-1.033-1.086M18.25 9.766c0-.59-.089-1.16-.254-1.701L9.15 16.91h5.825c.098-.476.197-.925.284-1.294c.106-.449.382-.803.705-1.04c1.464-1.077 2.285-2.806 2.285-4.81M9.311 18.41c.11.637.197 1.24.224 1.674c.027.457.368.866.871.974l.196.043c.92.199 1.875.199 2.796 0l.196-.043c.503-.108.844-.517.872-.974c.026-.433.112-1.037.223-1.674z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightbulbOffSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 9.766C4.5 5.737 7.874 2.5 12 2.5c1.925 0 3.687.705 5.018 1.866a.29.29 0 0 1 .01.425L6.92 14.901a.29.29 0 0 1-.412.003C5.194 13.578 4.5 11.754 4.5 9.766m-.497 10.17l-.033.034a.75.75 0 1 0 1.06 1.06l4.121-4.12h6.858c.142 0 .264-.1.294-.238c.061-.28.12-.541.174-.77a.571.571 0 0 1 .228-.319c1.822-1.34 2.795-3.463 2.795-5.817a7.05 7.05 0 0 0-.526-2.68L21.03 5.03a.75.75 0 0 0-1.033-1.086L4.03 19.91a.714.714 0 0 1-.027.026"/><path fill="currentColor" d="M8.281 18.41a.201.201 0 0 0-.198.235c.105.604.182 1.147.204 1.515c.063 1.041.83 1.899 1.855 2.12l.196.043a7.877 7.877 0 0 0 3.324 0l.196-.043c1.025-.221 1.792-1.079 1.855-2.12c.022-.368.1-.91.204-1.515a.201.201 0 0 0-.198-.235z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightbulbOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 3.75c-3.476 0-6.25 2.717-6.25 6.016c0 2.005.82 3.733 2.285 4.81c.323.237.6.591.705 1.04c.087.369.186.818.284 1.294h5.952c.098-.476.197-.925.284-1.294c.106-.449.382-.803.705-1.04c1.464-1.077 2.285-2.806 2.285-4.81c0-3.299-2.774-6.016-6.25-6.016m2.689 14.66H9.31c.11.637.197 1.24.224 1.674c.027.457.368.866.871.974l.196.043c.92.199 1.875.199 2.796 0l.196-.043c.503-.108.844-.517.872-.974c.026-.433.112-1.037.223-1.674M4.25 9.766C4.25 5.59 7.744 2.25 12 2.25s7.75 3.341 7.75 7.516c0 2.424-1.004 4.627-2.897 6.018a.324.324 0 0 0-.133.176a50.74 50.74 0 0 0-.394 1.843c-.183.938-.332 1.848-.363 2.372c-.07 1.158-.922 2.105-2.052 2.35l-.196.042c-1.13.244-2.3.244-3.43 0l-.196-.042c-1.13-.244-1.982-1.192-2.052-2.35c-.031-.524-.18-1.434-.363-2.372a50.745 50.745 0 0 0-.394-1.843a.324.324 0 0 0-.133-.176C5.254 14.394 4.25 12.19 4.25 9.767" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightbulbSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 9.766C4.5 5.737 7.874 2.5 12 2.5c4.126 0 7.5 3.237 7.5 7.266c0 2.354-.973 4.478-2.795 5.817a.571.571 0 0 0-.228.32c-.054.228-.113.489-.174.769a.301.301 0 0 1-.294.238H7.991a.301.301 0 0 1-.294-.238c-.061-.28-.12-.541-.174-.77a.57.57 0 0 0-.228-.319C5.473 14.243 4.5 12.12 4.5 9.766m3.9 8.644a.302.302 0 0 0-.296.353c.093.556.162 1.053.183 1.397c.063 1.041.83 1.899 1.855 2.12l.196.043a7.877 7.877 0 0 0 3.324 0l.196-.043c1.025-.221 1.792-1.079 1.855-2.12c.021-.343.09-.84.184-1.397a.302.302 0 0 0-.297-.354z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightningAltOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.614 2.31a.75.75 0 0 1 .456.69v6.998H18a.75.75 0 0 1 .653 1.12l-.492.87a35.748 35.748 0 0 1-7.05 8.842l-.796.725A.75.75 0 0 1 9.06 21v-6.939H5a.75.75 0 0 1-.653-1.119a35.801 35.801 0 0 1 6.675-8.773l1.778-1.71a.75.75 0 0 1 .814-.149m-7.33 10.251H9.81a.75.75 0 0 1 .75.75v5.983a34.247 34.247 0 0 0 6.153-7.796H13.32a.75.75 0 0 1-.75-.75V4.762l-.508.488a34.3 34.3 0 0 0-5.777 7.311" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightningAltSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.82 3a.5.5 0 0 0-.847-.36l-1.778 1.71a35.551 35.551 0 0 0-6.63 8.715a.5.5 0 0 0 .435.746h4.31V21a.5.5 0 0 0 .837.37l.795-.725a35.498 35.498 0 0 0 7.001-8.78l.492-.87a.5.5 0 0 0-.435-.747h-4.18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightningOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.94 2.512a20.75 20.75 0 0 1 4.692 0l1.596.182a.75.75 0 0 1 .62 1.003l-2.203 6.016H17a.75.75 0 0 1 .702 1.014l-.19.507a35.749 35.749 0 0 1-5.535 9.745l-.391.49A.75.75 0 0 1 10.25 21v-7.152H8.429a.75.75 0 0 1-.748-.684l-.122-1.382a39.75 39.75 0 0 1 0-7.027l.122-1.382a.75.75 0 0 1 .663-.68zm4.523 1.49a19.25 19.25 0 0 0-4.354 0l-.987.113l-.069.772c-.2 2.25-.2 4.513 0 6.763l.062.698H11a.75.75 0 0 1 .75.75v5.704a34.251 34.251 0 0 0 4.163-7.589H13.57a.75.75 0 0 1-.704-1.007l2.244-6.13z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightningSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.604 2.76a20.5 20.5 0 0 0-4.637 0l-1.595.182a.5.5 0 0 0-.441.453l-.123 1.382a39.5 39.5 0 0 0 0 6.983l.123 1.382a.5.5 0 0 0 .498.456H10.5V21a.5.5 0 0 0 .89.312l.391-.49a35.5 35.5 0 0 0 5.497-9.676l.19-.507A.5.5 0 0 0 17 9.963h-2.713l2.325-6.352a.5.5 0 0 0-.413-.669z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LikeOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.702 5.25c-.166 0-.32.09-.402.234l-3.346 5.873a4.25 4.25 0 0 0-.531 2.568l.353 3.214a1.236 1.236 0 0 0 1.083 1.092l2.117.252a9.567 9.567 0 0 0 4.15-.422c.66-.22 1.188-.72 1.442-1.367l1.704-4.33a1.42 1.42 0 0 0-1.565-1.92l-3.416.593a1.562 1.562 0 0 1-1.8-1.84l.665-3.395a.463.463 0 0 0-.454-.552m-1.705-.509a1.963 1.963 0 0 1 3.631 1.349l-.665 3.396a.056.056 0 0 0 0 .03a.065.065 0 0 0 .017.025a.066.066 0 0 0 .025.017a.057.057 0 0 0 .03.001l3.415-.593a2.92 2.92 0 0 1 3.218 3.947l-1.704 4.33a3.846 3.846 0 0 1-2.365 2.241c-1.545.514-3.184.68-4.8.488l-2.117-.251a2.736 2.736 0 0 1-2.397-2.419l-.353-3.213a5.75 5.75 0 0 1 .72-3.475z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LikeSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.337 3.75a1.4 1.4 0 0 0-1.149.6l-4.232 6.077a5.75 5.75 0 0 0-.997 3.914l.326 2.961a2.736 2.736 0 0 0 2.397 2.419l2.117.252c1.616.192 3.255.025 4.8-.489a3.846 3.846 0 0 0 2.365-2.24l1.709-4.343A2.818 2.818 0 0 0 15.43 9.12l-3.898.878l1.16-4.499a1.4 1.4 0 0 0-1.356-1.749"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkedinOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 1.25a2.75 2.75 0 1 0 0 5.5a2.75 2.75 0 0 0 0-5.5M3.75 4a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0m-1.5 4A.75.75 0 0 1 3 7.25h4a.75.75 0 0 1 .75.75v13a.75.75 0 0 1-.75.75H3a.75.75 0 0 1-.75-.75zm1.5.75v11.5h2.5V8.75zM9.25 8a.75.75 0 0 1 .75-.75h4a.75.75 0 0 1 .75.75v.434l.435-.187a7.792 7.792 0 0 1 2.358-.595C20.318 7.4 22.75 9.58 22.75 12.38V21a.75.75 0 0 1-.75.75h-4a.75.75 0 0 1-.75-.75v-7a1.25 1.25 0 0 0-2.5 0v7a.75.75 0 0 1-.75.75h-4a.75.75 0 0 1-.75-.75zm1.5.75v11.5h2.5V14a2.75 2.75 0 1 1 5.5 0v6.25h2.5v-7.87c0-1.904-1.661-3.408-3.57-3.234a6.31 6.31 0 0 0-1.904.48l-1.48.635a.75.75 0 0 1-1.046-.69V8.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkedinSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.75 1.875a2.125 2.125 0 1 0 0 4.25a2.125 2.125 0 0 0 0-4.25m-2 6A.125.125 0 0 0 2.625 8v13c0 .069.056.125.125.125h4A.125.125 0 0 0 6.875 21V8a.125.125 0 0 0-.125-.125zm6.5 0A.125.125 0 0 0 9.125 8v13c0 .069.056.125.125.125h4a.125.125 0 0 0 .125-.125v-7a1.875 1.875 0 1 1 3.75 0v7c0 .069.056.125.125.125h4a.125.125 0 0 0 .125-.125v-8.62c0-2.427-2.11-4.325-4.525-4.106a7.168 7.168 0 0 0-2.169.548l-1.306.56V8a.125.125 0 0 0-.125-.125z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationCheckOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M15.53 7.47a.75.75 0 0 1 0 1.06l-3.5 3.5a.75.75 0 0 1-1.06 0l-2-2a.75.75 0 1 1 1.06-1.06l1.47 1.47l2.97-2.97a.75.75 0 0 1 1.06 0"/><path d="M3.524 8.857a8.288 8.288 0 0 1 8.26-7.607h.432a8.288 8.288 0 0 1 8.26 7.607a8.944 8.944 0 0 1-1.99 6.396l-4.793 5.861a2.187 2.187 0 0 1-3.386 0l-4.793-5.861a8.943 8.943 0 0 1-1.99-6.396m8.26-6.107A6.788 6.788 0 0 0 5.02 8.98a7.443 7.443 0 0 0 1.656 5.323l4.793 5.862a.687.687 0 0 0 1.064 0l4.793-5.862A7.443 7.443 0 0 0 18.98 8.98a6.788 6.788 0 0 0-6.765-6.23z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationCheckSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.774 8.877a8.038 8.038 0 0 1 8.01-7.377h.432a8.038 8.038 0 0 1 8.01 7.377a8.693 8.693 0 0 1-1.933 6.217L13.5 20.956a1.937 1.937 0 0 1-3 0l-4.792-5.862a8.693 8.693 0 0 1-1.934-6.217M15.53 8.53a.75.75 0 0 0-1.06-1.06l-2.97 2.97l-1.47-1.47a.75.75 0 1 0-1.06 1.06l2 2a.75.75 0 0 0 1.06 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M7.25 10a4.75 4.75 0 1 1 9.5 0a4.75 4.75 0 0 1-9.5 0M12 6.75a3.25 3.25 0 1 0 0 6.5a3.25 3.25 0 0 0 0-6.5"/><path d="M3.524 8.857a8.288 8.288 0 0 1 8.26-7.607h.432a8.288 8.288 0 0 1 8.26 7.607a8.944 8.944 0 0 1-1.99 6.396l-4.793 5.861a2.187 2.187 0 0 1-3.386 0l-4.793-5.861a8.943 8.943 0 0 1-1.99-6.396m8.26-6.107A6.788 6.788 0 0 0 5.02 8.98a7.443 7.443 0 0 0 1.656 5.323l4.793 5.862a.687.687 0 0 0 1.064 0l4.793-5.862A7.443 7.443 0 0 0 18.98 8.98a6.788 6.788 0 0 0-6.765-6.23z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationPlusOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 6.25a.75.75 0 0 1 .75.75v2h2a.75.75 0 0 1 0 1.5h-2v2a.75.75 0 1 1-1.5 0v-2h-2a.75.75 0 0 1 0-1.5h2V7a.75.75 0 0 1 .75-.75"/><path fill="currentColor" fill-rule="evenodd" d="M11.784 1.25a8.288 8.288 0 0 0-8.26 7.607a8.943 8.943 0 0 0 1.99 6.396l4.793 5.861a2.187 2.187 0 0 0 3.386 0l4.793-5.861a8.944 8.944 0 0 0 1.99-6.396a8.288 8.288 0 0 0-8.26-7.607zM5.02 8.98a6.788 6.788 0 0 1 6.765-6.23h.432a6.788 6.788 0 0 1 6.765 6.23a7.443 7.443 0 0 1-1.656 5.323l-4.793 5.862a.687.687 0 0 1-1.064 0l-4.793-5.862A7.443 7.443 0 0 1 5.02 8.98" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationPlusSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.774 8.877a8.038 8.038 0 0 1 8.01-7.377h.432a8.038 8.038 0 0 1 8.01 7.377a8.693 8.693 0 0 1-1.933 6.217L13.5 20.956a1.937 1.937 0 0 1-3 0l-4.792-5.862a8.693 8.693 0 0 1-1.934-6.217M12 6.25a.75.75 0 0 1 .75.75v2h2a.75.75 0 0 1 0 1.5h-2v2a.75.75 0 0 1-1.5 0v-2h-2a.75.75 0 0 1 0-1.5h2V7a.75.75 0 0 1 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationQuestionOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.524 8.857a8.288 8.288 0 0 1 8.26-7.607h.432a8.288 8.288 0 0 1 8.26 7.607a8.944 8.944 0 0 1-1.99 6.396l-4.793 5.861a2.187 2.187 0 0 1-3.386 0l-4.793-5.861a8.943 8.943 0 0 1-1.99-6.396m8.26-6.107A6.788 6.788 0 0 0 5.02 8.98a7.443 7.443 0 0 0 1.656 5.323l4.793 5.862a.687.687 0 0 0 1.064 0l4.793-5.862A7.443 7.443 0 0 0 18.98 8.98a6.788 6.788 0 0 0-6.765-6.23z" clip-rule="evenodd"/><path fill="currentColor" d="M13 16a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/><path fill="currentColor" fill-rule="evenodd" d="M12 7.75c-.947 0-1.75.795-1.75 1.821a.75.75 0 1 1-1.5 0c0-1.814 1.435-3.321 3.25-3.321s3.25 1.507 3.25 3.321c0 1.431-.888 2.664-2.152 3.127a.735.735 0 0 0-.287.184c-.057.063-.061.102-.061.118a.75.75 0 0 1-1.5 0c0-.924.743-1.494 1.332-1.71a1.82 1.82 0 0 0 1.168-1.719c0-1.026-.803-1.821-1.75-1.821" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationQuestionSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.774 8.877a8.038 8.038 0 0 1 8.01-7.377h.432a8.038 8.038 0 0 1 8.01 7.377a8.693 8.693 0 0 1-1.933 6.217L13.5 20.956a1.937 1.937 0 0 1-3 0l-4.792-5.862a8.693 8.693 0 0 1-1.934-6.217m6.476.694c0-1.026.803-1.821 1.75-1.821s1.75.795 1.75 1.821a1.82 1.82 0 0 1-1.168 1.719c-.589.216-1.332.786-1.332 1.71a.75.75 0 0 0 1.5 0c0-.016.004-.055.061-.118a.735.735 0 0 1 .287-.184c1.264-.463 2.152-1.696 2.152-3.127c0-1.814-1.435-3.321-3.25-3.321S8.75 7.757 8.75 9.571a.75.75 0 0 0 1.5 0M12 17a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.75 10a3.25 3.25 0 1 1 6.5 0a3.25 3.25 0 0 1-6.5 0"/><path fill="currentColor" fill-rule="evenodd" d="M3.774 8.877a8.038 8.038 0 0 1 8.01-7.377h.432a8.038 8.038 0 0 1 8.01 7.377a8.693 8.693 0 0 1-1.933 6.217L13.5 20.956a1.937 1.937 0 0 1-3 0l-4.792-5.862a8.693 8.693 0 0 1-1.934-6.217M12 5.25a4.75 4.75 0 1 0 0 9.5a4.75 4.75 0 0 0 0-9.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 16a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0"/><path fill="currentColor" fill-rule="evenodd" d="m7.622 10.597l-.316-2.839a4.96 4.96 0 0 1 0-1.095l.023-.205a4.7 4.7 0 0 1 9.342 0l.023.205c.04.364.04.731 0 1.095l-.316 2.84l.687.054a2.36 2.36 0 0 1 2.142 1.972a20.89 20.89 0 0 1 0 6.752a2.361 2.361 0 0 1-2.142 1.972l-1.496.12c-2.376.19-4.762.19-7.138 0l-1.496-.12a2.361 2.361 0 0 1-2.142-1.972a20.891 20.891 0 0 1 0-6.752a2.361 2.361 0 0 1 2.142-1.972zM11.626 3.8a3.2 3.2 0 0 1 3.554 2.825l.023.205c.028.253.028.51 0 .764l-.321 2.89a44.84 44.84 0 0 0-5.764 0l-.32-2.89a3.46 3.46 0 0 1 0-.764l.022-.205A3.2 3.2 0 0 1 11.626 3.8m3.824 8.229a43.367 43.367 0 0 0-6.9 0l-1.495.12a.861.861 0 0 0-.782.719a19.39 19.39 0 0 0 0 6.266a.861.861 0 0 0 .782.72l1.496.12c2.296.183 4.602.183 6.899 0l1.496-.12a.861.861 0 0 0 .781-.72a19.39 19.39 0 0 0 0-6.266a.861.861 0 0 0-.781-.72z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7.306 7.758l.343 3.088l-.694.055a2.111 2.111 0 0 0-1.915 1.764a20.641 20.641 0 0 0 0 6.67A2.111 2.111 0 0 0 6.955 21.1l1.496.12c2.362.188 4.736.188 7.098 0l1.496-.12a2.111 2.111 0 0 0 1.915-1.764a20.64 20.64 0 0 0 0-6.67a2.111 2.111 0 0 0-1.915-1.764l-.694-.055l.343-3.088a5.01 5.01 0 0 0 0-1.095l-.023-.205a4.7 4.7 0 0 0-9.342 0l-.023.205a4.96 4.96 0 0 0 0 1.095M12.374 3.8A3.2 3.2 0 0 0 8.82 6.624l-.023.205a3.46 3.46 0 0 0 0 .764l.349 3.139c1.9-.122 3.807-.122 5.708 0l.349-3.14a3.46 3.46 0 0 0 0-.763l-.023-.205a3.2 3.2 0 0 0-2.806-2.825M12 14.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockTimeOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m5.625 9.631l-.319-2.873a4.96 4.96 0 0 1 0-1.095l.023-.205a4.7 4.7 0 0 1 9.342 0l.023.205c.04.364.04.731 0 1.095l-.281 2.529a.416.416 0 0 1-.333.357a6.966 6.966 0 0 0-2.706 1.2a.414.414 0 0 1-.256.082a40.922 40.922 0 0 0-4.697.141l-1.135.1a1.14 1.14 0 0 0-1.027.967a22.863 22.863 0 0 0 0 6.732a1.14 1.14 0 0 0 1.027.967l1.135.1c.946.083 1.894.133 2.843.15a.416.416 0 0 1 .341.194c.219.34.466.66.74.958c.12.132.032.353-.147.354a42.4 42.4 0 0 1-3.908-.162l-1.135-.1a2.64 2.64 0 0 1-2.38-2.24a24.364 24.364 0 0 1 0-7.174a2.64 2.64 0 0 1 2.38-2.24zM9.626 2.8a3.2 3.2 0 0 1 3.554 2.825l.023.205c.028.253.028.51 0 .764l-.324 2.915c-1.917-.13-3.84-.13-5.758 0l-.324-2.915a3.46 3.46 0 0 1 0-.764l.023-.205A3.2 3.2 0 0 1 9.626 2.8" clip-rule="evenodd"/><path fill="currentColor" d="M16.25 15a.75.75 0 0 0-1.5 0v1.773c0 .24.115.465.309.606l1 .727a.75.75 0 1 0 .882-1.213l-.691-.502z"/><path fill="currentColor" fill-rule="evenodd" d="M15.5 22a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11m0-1.5a4 4 0 1 0 0-8a4 4 0 0 0 0 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockTimeSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m5.306 6.758l.347 3.122l-.476.042a2.389 2.389 0 0 0-2.154 2.028a24.113 24.113 0 0 0 0 7.1a2.39 2.39 0 0 0 2.154 2.028l1.134.1c1.22.107 2.444.161 3.668.162c.175 0 .266-.212.154-.346a7.002 7.002 0 0 1 3.947-11.35a.416.416 0 0 0 .333-.357l.28-2.529a4.89 4.89 0 0 0 0-1.095l-.022-.205a4.7 4.7 0 0 0-9.342 0l-.023.205a4.96 4.96 0 0 0 0 1.095M10.374 2.8A3.2 3.2 0 0 0 6.82 5.624l-.023.205a3.46 3.46 0 0 0 0 .764l.352 3.164a42.166 42.166 0 0 1 5.702 0l.352-3.164a3.46 3.46 0 0 0 0-.764l-.023-.205a3.2 3.2 0 0 0-2.806-2.825" clip-rule="evenodd"/><path fill="currentColor" d="M16.25 15a.75.75 0 0 0-1.5 0v1.773c0 .24.115.465.309.606l1 .727a.75.75 0 1 0 .882-1.213l-.691-.502z"/><path fill="currentColor" fill-rule="evenodd" d="M15.5 22a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11m0-1.5a4 4 0 1 0 0-8a4 4 0 0 0 0 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LoginOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 18.25a.75.75 0 0 0 0 1.5h6A1.75 1.75 0 0 0 19.75 18V6A1.75 1.75 0 0 0 18 4.25h-6a.75.75 0 0 0 0 1.5h6a.25.25 0 0 1 .25.25v12a.25.25 0 0 1-.25.25z"/><path fill="currentColor" fill-rule="evenodd" d="M3.25 13.115c0 .69.56 1.25 1.25 1.25h4.613l.02.22l.054.556a1.227 1.227 0 0 0 1.752.988c1.634-.783 3.212-1.958 4.466-3.267a1.253 1.253 0 0 0 0-1.734l-.1-.104a15.06 15.06 0 0 0-4.366-3.163a1.227 1.227 0 0 0-1.752.988l-.054.555l-.02.22H4.5c-.69 0-1.25.56-1.25 1.25zm7.303.416a.75.75 0 0 0-.745-.666H4.75v-1.74h5.058a.75.75 0 0 0 .749-.704c.018-.29.04-.581.069-.871l.016-.162a13.556 13.556 0 0 1 3.516 2.607a13.568 13.568 0 0 1-3.516 2.607l-.016-.162a25.457 25.457 0 0 1-.073-.91" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LoginSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.25 19a.75.75 0 0 1 .75-.75h6a.25.25 0 0 0 .25-.25V6a.25.25 0 0 0-.25-.25h-6a.75.75 0 0 1 0-1.5h6c.966 0 1.75.784 1.75 1.75v12A1.75 1.75 0 0 1 18 19.75h-6a.75.75 0 0 1-.75-.75"/><path fill="currentColor" d="M3.5 13.115a1 1 0 0 0 1 1h4.856c.023.356.052.71.086 1.066l.03.305a.718.718 0 0 0 1.025.578a16.844 16.844 0 0 0 4.884-3.539l.03-.031a.721.721 0 0 0 0-.998l-.03-.031a16.842 16.842 0 0 0-4.884-3.539a.718.718 0 0 0-1.025.578l-.03.305c-.034.355-.063.71-.086 1.066H4.5a1 1 0 0 0-1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoutOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 18.25a.75.75 0 0 0 0 1.5h6A1.75 1.75 0 0 0 19.75 18V6A1.75 1.75 0 0 0 18 4.25h-6a.75.75 0 0 0 0 1.5h6a.25.25 0 0 1 .25.25v12a.25.25 0 0 1-.25.25z"/><path fill="currentColor" fill-rule="evenodd" d="M14.503 14.365c.69 0 1.25-.56 1.25-1.25v-2.24c0-.69-.56-1.25-1.25-1.25H9.89a26.723 26.723 0 0 0-.02-.22l-.054-.556a1.227 1.227 0 0 0-1.751-.988a15.059 15.059 0 0 0-4.368 3.164l-.099.103a1.253 1.253 0 0 0 0 1.734l.1.103a15.06 15.06 0 0 0 4.367 3.164a1.227 1.227 0 0 0 1.751-.988l.054-.556l.02-.22zm-5.308-1.5a.75.75 0 0 0-.748.704c-.019.29-.042.581-.07.871l-.016.162a13.562 13.562 0 0 1-3.516-2.607a13.558 13.558 0 0 1 3.516-2.607l.016.162c.028.29.051.58.07.871a.75.75 0 0 0 .748.704h5.058v1.74z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoutSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.25 19a.75.75 0 0 1 .75-.75h6a.25.25 0 0 0 .25-.25V6a.25.25 0 0 0-.25-.25h-6a.75.75 0 0 1 0-1.5h6c.966 0 1.75.784 1.75 1.75v12A1.75 1.75 0 0 1 18 19.75h-6a.75.75 0 0 1-.75-.75"/><path fill="currentColor" d="M15.612 13.115a1 1 0 0 1-1 1H9.756c-.023.356-.052.71-.086 1.066l-.03.305a.718.718 0 0 1-1.025.578a16.844 16.844 0 0 1-4.885-3.539l-.03-.031a.721.721 0 0 1 0-.998l.03-.031a16.843 16.843 0 0 1 4.885-3.539a.718.718 0 0 1 1.025.578l.03.305c.034.355.063.71.086 1.066h4.856a1 1 0 0 1 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapLocationOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.25 9a3.75 3.75 0 1 1 7.5 0a3.75 3.75 0 0 1-7.5 0M12 6.75a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M5.456 8.127a6.403 6.403 0 0 1 6.382-5.877h.324a6.403 6.403 0 0 1 6.382 5.877a6.895 6.895 0 0 1-1.534 4.931l-3.595 4.397a1.828 1.828 0 0 1-2.83 0L6.99 13.058a6.895 6.895 0 0 1-1.534-4.93m6.382-4.377a4.903 4.903 0 0 0-4.887 4.5a5.395 5.395 0 0 0 1.2 3.859l3.595 4.396c.131.16.377.16.508 0l3.595-4.396a5.395 5.395 0 0 0 1.2-3.859a4.903 4.903 0 0 0-4.887-4.5z" clip-rule="evenodd"/><path fill="currentColor" d="M7.67 16.335a.75.75 0 1 0-1.34-.67l-2 4A.75.75 0 0 0 5 20.75h14a.75.75 0 0 0 .67-1.085l-2-4a.75.75 0 1 0-1.34.67l1.457 2.915H6.214z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapLocationSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.75 9a2.25 2.25 0 1 1 4.5 0a2.25 2.25 0 0 1-4.5 0"/><path fill="currentColor" fill-rule="evenodd" d="M11.838 2.5a6.153 6.153 0 0 0-6.132 5.648A6.645 6.645 0 0 0 7.184 12.9l3.595 4.396a1.578 1.578 0 0 0 2.442 0l3.595-4.396a6.644 6.644 0 0 0 1.478-4.752A6.153 6.153 0 0 0 12.162 2.5zM12 5.25a3.75 3.75 0 1 0 0 7.5a3.75 3.75 0 0 0 0-7.5" clip-rule="evenodd"/><path fill="currentColor" d="M7.335 15.33a.75.75 0 0 1 .336 1.005L6.214 19.25h11.573l-1.458-2.915a.75.75 0 1 1 1.342-.67l2 4A.75.75 0 0 1 19 20.75H5a.75.75 0 0 1-.67-1.085l2-4a.75.75 0 0 1 1.005-.336"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediumOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16.41 8.98l.84-.803V8h-2.902l-2.069 5.16L9.927 8H6.883v.177l.98 1.18a.406.406 0 0 1 .133.343v4.641a.534.534 0 0 1-.142.46l-1.104 1.34v.175h3.127v-.176l-1.103-1.338a.553.553 0 0 1-.152-.46v-4.014l2.745 5.989h.32l2.357-5.989v4.773c0 .127 0 .151-.084.235l-.848.823v.176h4.117v-.176l-.818-.804a.245.245 0 0 1-.094-.235V9.215a.245.245 0 0 1 .094-.235"/><path fill="currentColor" fill-rule="evenodd" d="M17.258 2.833a47.721 47.721 0 0 0-10.516 0c-2.012.225-3.637 1.81-3.873 3.832a45.921 45.921 0 0 0 0 10.67c.236 2.022 1.86 3.607 3.873 3.832a47.77 47.77 0 0 0 10.516 0c2.012-.225 3.637-1.81 3.873-3.832a45.925 45.925 0 0 0 0-10.67c-.236-2.022-1.86-3.607-3.873-3.832m-10.35 1.49a46.22 46.22 0 0 1 10.184 0c1.33.15 2.395 1.199 2.55 2.517a44.421 44.421 0 0 1 0 10.32a2.89 2.89 0 0 1-2.55 2.516a46.216 46.216 0 0 1-10.184 0a2.89 2.89 0 0 1-2.55-2.516a44.421 44.421 0 0 1 0-10.32a2.89 2.89 0 0 1 2.55-2.516" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediumSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.77 3.082a47.472 47.472 0 0 1 10.46 0c1.899.212 3.43 1.707 3.653 3.613a45.67 45.67 0 0 1 0 10.61c-.223 1.906-1.754 3.401-3.652 3.614a47.468 47.468 0 0 1-10.461 0c-1.899-.213-3.43-1.708-3.653-3.613a45.672 45.672 0 0 1 0-10.611C3.34 4.789 4.871 3.294 6.77 3.082m9.64 5.898l.84-.803V8h-2.902l-2.069 5.16L9.927 8H6.883v.177l.98 1.18a.406.406 0 0 1 .133.343v4.641a.534.534 0 0 1-.142.46l-1.104 1.34v.175h3.127v-.176l-1.103-1.338a.553.553 0 0 1-.152-.46v-4.014l2.745 5.989h.32l2.357-5.989v4.773c0 .127 0 .151-.084.235l-.848.823v.176h4.117v-.176l-.818-.804a.245.245 0 0 1-.094-.235V9.215a.245.245 0 0 1 .094-.235" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MedkitOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.5 15.25a.75.75 0 0 0 0 1.5h1.75v1.75a.75.75 0 1 0 1.5 0v-1.75h1.75a.75.75 0 1 0 0-1.5h-1.75V13.5a.75.75 0 0 0-1.5 0v1.75z"/><path fill="currentColor" fill-rule="evenodd" d="M7.25 5.461V6.7c-.37.028-.738.059-1.107.093a3.246 3.246 0 0 0-2.946 3.233v7.95a3.247 3.247 0 0 0 2.946 3.233c3.896.363 7.818.363 11.714 0a3.247 3.247 0 0 0 2.946-3.233v-7.95a3.246 3.246 0 0 0-2.946-3.233A63.222 63.222 0 0 0 16.75 6.7V5.46a1.75 1.75 0 0 0-1.49-1.73l-1.22-.183a13.75 13.75 0 0 0-4.08 0l-1.22.183a1.75 1.75 0 0 0-1.49 1.73m6.567-.43a12.25 12.25 0 0 0-3.634 0l-1.22.183a.25.25 0 0 0-.213.247v1.143a63.161 63.161 0 0 1 6.5 0V5.46a.25.25 0 0 0-.213-.247zm3.901 3.255a61.661 61.661 0 0 0-11.436 0a1.746 1.746 0 0 0-1.585 1.739v.225h14.606v-.225c0-.902-.687-1.656-1.585-1.74m1.585 3.464H4.697v6.225c0 .902.687 1.656 1.585 1.74a61.88 61.88 0 0 0 11.436 0a1.746 1.746 0 0 0 1.585-1.74z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MedkitSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.25 5.461V6.95a62.92 62.92 0 0 0-1.084.091c-1.441.134-2.565 1.471-2.704 2.91c-.016.164.12.299.285.299h16.506a.279.279 0 0 0 .285-.3c-.14-1.438-1.263-2.775-2.704-2.909a62.813 62.813 0 0 0-1.084-.091V5.46a1.75 1.75 0 0 0-1.49-1.73l-1.22-.183a13.75 13.75 0 0 0-4.08 0l-1.22.183a1.75 1.75 0 0 0-1.49 1.73m6.567-.43a12.25 12.25 0 0 0-3.634 0l-1.22.183a.25.25 0 0 0-.213.247v1.393a62.913 62.913 0 0 1 6.5 0V5.461a.25.25 0 0 0-.213-.247zm6.736 7.02a.3.3 0 0 0-.3-.3H3.747a.3.3 0 0 0-.3.3v5.925a2.996 2.996 0 0 0 2.719 2.984a62.99 62.99 0 0 0 11.668 0a2.996 2.996 0 0 0 2.719-2.984zM9.5 15.25a.75.75 0 1 0 0 1.5h1.75v1.75a.75.75 0 0 0 1.5 0v-1.75h1.75a.75.75 0 0 0 0-1.5h-1.75V13.5a.75.75 0 1 0-1.5 0v1.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M19.75 12a.75.75 0 0 0-.75-.75H5a.75.75 0 0 0 0 1.5h14a.75.75 0 0 0 .75-.75m0-5a.75.75 0 0 0-.75-.75H5a.75.75 0 0 0 0 1.5h14a.75.75 0 0 0 .75-.75m0 10a.75.75 0 0 0-.75-.75H5a.75.75 0 0 0 0 1.5h14a.75.75 0 0 0 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M19.75 12a.75.75 0 0 0-.75-.75H5a.75.75 0 0 0 0 1.5h14a.75.75 0 0 0 .75-.75m0-5a.75.75 0 0 0-.75-.75H5a.75.75 0 0 0 0 1.5h14a.75.75 0 0 0 .75-.75m0 10a.75.75 0 0 0-.75-.75H5a.75.75 0 0 0 0 1.5h14a.75.75 0 0 0 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneOffOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.734 2.699a11.29 11.29 0 0 1 4.532 0c.647.132 1.145.65 1.252 1.302l.055.329c.092.56.166 1.123.22 1.688a.249.249 0 0 1-.071.2l-.923.922a.255.255 0 0 1-.433-.166a24.13 24.13 0 0 0-.273-2.4l-.054-.33a.093.093 0 0 0-.074-.076a9.79 9.79 0 0 0-3.93 0a.093.093 0 0 0-.074.076l-.054.33a24.147 24.147 0 0 0-.072 7.387a.252.252 0 0 1-.07.214l-.919.918c-.125.125-.336.081-.365-.094l-.054-.328a25.647 25.647 0 0 1 0-8.341L8.481 4a1.593 1.593 0 0 1 1.253-1.3"/><path fill="currentColor" fill-rule="evenodd" d="m15.87 9.99l5.16-5.16a.75.75 0 0 0-1.06-1.06l-16 16a.75.75 0 1 0 1.06 1.06l2.629-2.628h.007c1.19.153 2.387.242 3.584.268v4.03a.75.75 0 0 0 1.5 0v-4.03a34.194 34.194 0 0 0 3.584-.267a3.426 3.426 0 0 0 2.958-2.918l.45-3.18a.75.75 0 0 0-1.485-.21l-.45 3.18a1.926 1.926 0 0 1-1.663 1.64a32.704 32.704 0 0 1-7.125.127l2.33-2.33a11.29 11.29 0 0 0 2.917-.21a1.593 1.593 0 0 0 1.252-1.303l.055-.328c.146-.89.246-1.784.298-2.681m-1.663 1.663l-1.339 1.34a9.78 9.78 0 0 0 1.097-.161a.093.093 0 0 0 .074-.076l.054-.33a24.8 24.8 0 0 0 .114-.773" clip-rule="evenodd"/><path fill="currentColor" d="M6.193 15.074c.018.125.047.245.087.361a.272.272 0 0 1-.058.282l-.762.763c-.112.111-.299.093-.374-.046a3.42 3.42 0 0 1-.378-1.15l-.45-3.179a.75.75 0 1 1 1.485-.21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneOffSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.08 4.41c.113.689.197 1.38.253 2.073a.249.249 0 0 1-.072.195l-6.015 6.015a.19.19 0 0 1-.325-.104a25.147 25.147 0 0 1 0-8.178l.054-.33c.073-.446.415-.801.859-.892a10.79 10.79 0 0 1 4.332 0c.444.09.785.446.86.893zm-8.887 9.664c.049.346.188.66.394.921c.086.11.088.269-.01.368l-.706.706a.236.236 0 0 1-.346-.005a3.42 3.42 0 0 1-.817-1.78l-.45-3.179a.75.75 0 1 1 1.485-.21zm2.363 3.23L5.03 20.83a.75.75 0 0 1-1.06-1.06l16-16a.75.75 0 1 1 1.06 1.06l-5.698 5.698a25.09 25.09 0 0 1-.253 2.061l-.054.33a1.093 1.093 0 0 1-.859.892c-.77.158-1.553.23-2.335.219l-1.884 1.884c2.066.13 4.14.063 6.197-.2a1.926 1.926 0 0 0 1.663-1.64l.45-3.18a.75.75 0 0 1 1.486.211l-.451 3.18a3.426 3.426 0 0 1-2.958 2.918c-1.19.152-2.387.24-3.584.267v4.03a.75.75 0 0 1-1.5 0v-4.03c-.9-.02-1.798-.075-2.694-.166"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.734 2.699a11.29 11.29 0 0 1 4.532 0c.647.132 1.145.65 1.252 1.302l.055.329a25.65 25.65 0 0 1 0 8.34l-.055.33a1.593 1.593 0 0 1-1.252 1.301a11.29 11.29 0 0 1-4.532 0A1.593 1.593 0 0 1 8.48 13l-.054-.328a25.647 25.647 0 0 1 0-8.341L8.481 4a1.593 1.593 0 0 1 1.253-1.3m4.231 1.47a9.79 9.79 0 0 0-3.93 0a.093.093 0 0 0-.074.075l-.054.33a24.147 24.147 0 0 0 0 7.853l.054.329a.093.093 0 0 0 .074.076a9.79 9.79 0 0 0 3.93 0a.093.093 0 0 0 .074-.076l.054-.33c.428-2.6.428-5.252 0-7.853l-.054-.329a.093.093 0 0 0-.074-.076" clip-rule="evenodd"/><path fill="currentColor" d="M4.895 11.257a.75.75 0 0 1 .848.638l.45 3.18c.122.858.803 1.53 1.663 1.64c2.752.351 5.536.351 8.288 0a1.926 1.926 0 0 0 1.663-1.64l.45-3.18a.75.75 0 0 1 1.486.21l-.451 3.18a3.426 3.426 0 0 1-2.958 2.918c-1.19.152-2.387.24-3.584.267v4.03a.75.75 0 0 1-1.5 0v-4.03a34.197 34.197 0 0 1-3.584-.267a3.426 3.426 0 0 1-2.958-2.918l-.45-3.18a.75.75 0 0 1 .637-.848"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.166 3.189a10.79 10.79 0 0 0-4.332 0c-.444.09-.786.446-.86.893l-.053.329a25.147 25.147 0 0 0 0 8.178l.054.33c.073.446.415.801.859.892c1.43.293 2.903.293 4.332 0c.444-.09.785-.446.86-.893l.053-.329c.447-2.708.447-5.47 0-8.178l-.054-.33a1.093 1.093 0 0 0-.859-.892"/><path fill="currentColor" d="M4.895 10.257a.75.75 0 0 1 .848.638l.45 3.18c.122.858.803 1.53 1.663 1.64c2.752.351 5.536.351 8.288 0a1.926 1.926 0 0 0 1.663-1.64l.45-3.18a.75.75 0 0 1 1.486.21l-.451 3.18a3.426 3.426 0 0 1-2.958 2.918c-1.19.152-2.387.24-3.584.267v4.03a.75.75 0 0 1-1.5 0v-4.03a34.197 34.197 0 0 1-3.584-.267a3.426 3.426 0 0 1-2.958-2.918l-.45-3.18a.75.75 0 0 1 .637-.848"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobilePhoneOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m18.75 10.5l-.006.243l.006.257l-.006.243l.006.257l-.006.245l.006.255l-.162 7.103a3.359 3.359 0 0 1-2.996 3.263a33.364 33.364 0 0 1-7.184 0a3.359 3.359 0 0 1-2.996-3.263L5.25 12l.006-.257l-.006-.243l.006-.256L5.25 11l.006-.256l-.006-.244l.164-6.098c.002-.085.008-.169.018-.252l.003-.035l.003-.029a2.663 2.663 0 0 1 1.355-2.068a2.871 2.871 0 0 1 1.235-.405a43.333 43.333 0 0 1 7.944 0a2.871 2.871 0 0 1 1.233.404a2.663 2.663 0 0 1 1.356 2.063l.008.074c.01.082.015.165.017.248zm-1.676-6.18l.157 6.868l-.127 4.73a2.1 2.1 0 0 1-1.907 2.035a34.875 34.875 0 0 1-6.394 0a2.1 2.1 0 0 1-1.907-2.035l-.127-4.73l.157-6.867c0-.015 0-.03.002-.044c.06-.403.292-.75.623-.96a1.16 1.16 0 0 1 .412-.125c2.683-.29 5.39-.29 8.074 0c.148.016.287.059.412.124c.33.211.564.559.624.961zM12 21.7a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobilePhoneSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m18.75 11l-.006.243l.006.257l-.006.245l.006.255l-.189 8.253a2.047 2.047 0 0 1-1.826 1.989c-3.147.34-6.323.34-9.47 0a2.047 2.047 0 0 1-1.827-1.989L5.25 12l.006-.257l-.006-.243l.006-.256L5.25 11l.168-6.641l.007-.131l.01-.462A2.1 2.1 0 0 1 7.32 1.728c.087-.01.175-.018.262-.027a2.82 2.82 0 0 1 .365-.06a41.778 41.778 0 0 1 8.106 0c.124.013.246.033.364.06l.263.027a2.1 2.1 0 0 1 1.884 2.038l.011.463c.003.043.005.086.006.13zM16 17.75a.75.75 0 0 0 0-1.5H8a.75.75 0 0 0 0 1.5zm-4 3.95a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.486 4.768a7.25 7.25 0 1 0 7.399 9.51a6.25 6.25 0 0 1-7.398-9.51M3.25 12a8.75 8.75 0 0 1 10.074-8.65a.75.75 0 0 1 .336 1.342a4.75 4.75 0 1 0 5.83 7.499a.75.75 0 0 1 1.22.654A8.751 8.751 0 0 1 3.25 12" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3.5a8.5 8.5 0 1 0 8.46 9.32a.5.5 0 0 0-.812-.435a5 5 0 1 1-6.137-7.893a.5.5 0 0 0-.225-.895A8.563 8.563 0 0 0 12 3.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MouseAltOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.25 10a.75.75 0 0 0 1.5 0V7a.75.75 0 0 0-1.5 0z"/><path fill="currentColor" fill-rule="evenodd" d="M18.75 9.074a6.75 6.75 0 0 0-13.5 0v5.852a6.75 6.75 0 0 0 13.5 0zm-5.931-5.186a5.25 5.25 0 0 1 4.431 5.186v5.852a5.25 5.25 0 0 1-10.5 0V9.074a5.25 5.25 0 0 1 6.069-5.186" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MouseAltSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.014 2.654a6.5 6.5 0 0 1 5.486 6.42v5.852a6.5 6.5 0 0 1-13 0V9.074a6.5 6.5 0 0 1 7.514-6.42M11.25 10a.75.75 0 0 0 1.5 0V7a.75.75 0 0 0-1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MouseOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.053 2.407a6.75 6.75 0 0 1 5.697 6.667v5.852a6.75 6.75 0 0 1-13.5 0V9.074a6.75 6.75 0 0 1 7.803-6.667m4.197 6.667a5.25 5.25 0 0 0-4.5-5.196V9.25h4.5zm-6 .176h-4.5v-.176a5.25 5.25 0 0 1 4.5-5.196zm-4.5 5.676V10.75h10.5v4.176a5.25 5.25 0 0 1-10.5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MouseSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.973 2.647a.195.195 0 0 0-.223.195V8.95a.3.3 0 0 0 .3.3h5.15a.29.29 0 0 0 .293-.3c-.14-3.078-2.419-5.813-5.48-6.296a7.865 7.865 0 0 0-.04-.007m-1.946 0a.195.195 0 0 1 .223.195V8.95a.3.3 0 0 1-.3.3H5.8a.29.29 0 0 1-.293-.3c.14-3.078 2.419-5.813 5.48-6.296zM5.8 10.75a.3.3 0 0 0-.3.3v3.876a6.5 6.5 0 1 0 13 0V11.05a.3.3 0 0 0-.3-.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoveOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.47 6.53a.75.75 0 0 0 1.06 0l.72-.72v5.44H5.81l.72-.72a.75.75 0 1 0-1.06-1.06l-2 2a.75.75 0 0 0 0 1.06l2 2a.75.75 0 0 0 1.06-1.06l-.72-.72h5.44v5.44l-.72-.72a.75.75 0 1 0-1.06 1.06l2 2a.75.75 0 0 0 1.06 0l2-2a.75.75 0 1 0-1.06-1.06l-.72.72v-5.44h5.44l-.72.72a.75.75 0 1 0 1.06 1.06l2-2a.75.75 0 0 0 0-1.06l-2-2a.75.75 0 1 0-1.06 1.06l.72.72h-5.44V5.81l.72.72a.75.75 0 1 0 1.06-1.06l-2-2a.75.75 0 0 0-1.06 0l-2 2a.75.75 0 0 0 0 1.06"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoveSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.47 6.53a.75.75 0 0 0 1.06 0l.72-.72v5.44H5.81l.72-.72a.75.75 0 1 0-1.06-1.06l-2 2a.75.75 0 0 0 0 1.06l2 2a.75.75 0 0 0 1.06-1.06l-.72-.72h5.44v5.44l-.72-.72a.75.75 0 1 0-1.06 1.06l2 2a.75.75 0 0 0 1.06 0l2-2a.75.75 0 1 0-1.06-1.06l-.72.72v-5.44h5.44l-.72.72a.75.75 0 1 0 1.06 1.06l2-2a.75.75 0 0 0 0-1.06l-2-2a.75.75 0 1 0-1.06 1.06l.72.72h-5.44V5.81l.72.72a.75.75 0 1 0 1.06-1.06l-2-2a.75.75 0 0 0-1.06 0l-2 2a.75.75 0 0 0 0 1.06"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25a.75.75 0 0 0-.75.75v11.26a4.25 4.25 0 1 0 1.486 2.888a.76.76 0 0 0 .014-.148V7.75H18a2.75 2.75 0 1 0 0-5.5zm.75 4H18a1.25 1.25 0 1 0 0-2.5h-5.25zm-4.25 8.5a2.75 2.75 0 1 0 0 5.5a2.75 2.75 0 0 0 0-5.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2.25a.75.75 0 0 0-.75.75v11.26a4.25 4.25 0 1 0 1.486 2.888a.76.76 0 0 0 .014-.148V7.75H18a2.75 2.75 0 1 0 0-5.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NavigationOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.762 3.25c.46.006.99.223 1.276.734l.174.312a82.518 82.518 0 0 1 6.474 14.753c.218.67-.14 1.237-.589 1.5a1.484 1.484 0 0 1-1.552-.037l-4.728-3.04c-.522-.336-1.231-.337-1.72-.028L6.464 20.37a1.483 1.483 0 0 1-1.553.01c-.427-.255-.809-.808-.61-1.482A77.671 77.671 0 0 1 10.307 4.29l.164-.308c.284-.531.832-.737 1.29-.73m-.007 1.513l-.125.233A76.203 76.203 0 0 0 5.84 18.99l4.455-2.813c1-.63 2.337-.607 3.334.034l4.496 2.892a81.088 81.088 0 0 0-6.222-14.074z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NavigationSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.82 4.106c-.234-.417-.671-.601-1.06-.606c-.39-.005-.838.17-1.068.6l-.164.307a77.416 77.416 0 0 0-5.986 14.56c-.158.537.141.984.498 1.199c.356.213.86.264 1.291-.008l4.632-2.925c.574-.363 1.388-.358 1.99.029l4.727 3.04c.425.273.927.244 1.29.031c.37-.217.652-.673.479-1.207a82.268 82.268 0 0 0-6.455-14.708z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationOffOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a1 1 0 0 1 1 1v.75h.557c1.182 0 2.256.488 3.024 1.279a.24.24 0 0 1-.007.336l-.703.703c-.105.105-.278.095-.382-.01a2.708 2.708 0 0 0-1.932-.808h-3.114a2.714 2.714 0 0 0-2.709 2.544l-.22 3.534a8.877 8.877 0 0 1-1.574 4.516c-.05.073-.05.195-.111.258l-.764.77c-.098.1-.26.1-.34-.013a1.611 1.611 0 0 1-.017-1.871a7.377 7.377 0 0 0 1.308-3.754l.221-3.533a4.214 4.214 0 0 1 4.206-3.951H11V3a1 1 0 0 1 1-1"/><path fill="currentColor" fill-rule="evenodd" d="M17.786 8.074L21.03 4.83a.75.75 0 0 0-1.06-1.06l-16 16a.75.75 0 1 0 1.06 1.06l3.046-3.045l1.174.14V19a2.75 2.75 0 1 0 5.5 0v-1.075l3.407-.409a1.617 1.617 0 0 0 1.135-2.528a7.376 7.376 0 0 1-1.308-3.754zM16.372 9.49l-6.947 6.947l.334.04c1.489.178 2.993.178 4.482 0l3.737-.449a.117.117 0 0 0 .082-.183a8.877 8.877 0 0 1-1.573-4.516zM12 20.25c-.69 0-1.25-.56-1.25-1.25v-.75h2.5V19c0 .69-.56 1.25-1.25 1.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationOffSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 3a1 1 0 1 0-2 0v.75h-.557A4.214 4.214 0 0 0 6.237 7.7l-.221 3.534a7.377 7.377 0 0 1-1.308 3.754a1.61 1.61 0 0 0 .017 1.872c.081.112.242.112.34.014L16.574 5.365a.24.24 0 0 0 .007-.336a4.205 4.205 0 0 0-3.024-1.279H13z"/><path fill="currentColor" fill-rule="evenodd" d="M17.786 8.074L21.03 4.83a.75.75 0 0 0-1.06-1.06l-16 16a.75.75 0 1 0 1.06 1.06l3.046-3.045l1.174.14V19a2.75 2.75 0 1 0 5.5 0v-1.075l3.407-.409a1.617 1.617 0 0 0 1.135-2.528a7.376 7.376 0 0 1-1.308-3.754zM10.75 19a1.25 1.25 0 1 0 2.5 0v-.75h-2.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationOnOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.429 2.413a.75.75 0 0 0-1.13-.986l-1.292 1.48a4.75 4.75 0 0 0-1.17 3.024L2.78 8.65a.75.75 0 1 0 1.5.031l.056-2.718a3.25 3.25 0 0 1 .801-2.069z"/><path fill="currentColor" fill-rule="evenodd" d="M6.237 7.7a4.214 4.214 0 0 1 4.206-3.95H11V3a1 1 0 1 1 2 0v.75h.557a4.214 4.214 0 0 1 4.206 3.95l.221 3.534a7.376 7.376 0 0 0 1.308 3.754a1.617 1.617 0 0 1-1.135 2.529l-3.407.408V19a2.75 2.75 0 1 1-5.5 0v-1.075l-3.407-.409a1.617 1.617 0 0 1-1.135-2.528a7.377 7.377 0 0 0 1.308-3.754zm4.206-2.45a2.714 2.714 0 0 0-2.709 2.544l-.22 3.534a8.877 8.877 0 0 1-1.574 4.516a.117.117 0 0 0 .082.183l3.737.449c1.489.178 2.993.178 4.482 0l3.737-.449a.117.117 0 0 0 .082-.183a8.876 8.876 0 0 1-1.573-4.516l-.221-3.534a2.714 2.714 0 0 0-2.709-2.544zm1.557 15c-.69 0-1.25-.56-1.25-1.25v-.75h2.5V19c0 .69-.56 1.25-1.25 1.25" clip-rule="evenodd"/><path fill="currentColor" d="M17.643 1.355a.75.75 0 0 0-.072 1.058l1.292 1.48a3.25 3.25 0 0 1 .8 2.07l.057 2.717a.75.75 0 0 0 1.5-.031l-.057-2.718a4.75 4.75 0 0 0-1.17-3.024l-1.292-1.48a.75.75 0 0 0-1.058-.072"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationOnSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.429 2.413a.75.75 0 0 0-1.13-.986l-1.292 1.48a4.75 4.75 0 0 0-1.17 3.024L2.78 8.65a.75.75 0 0 0 1.5.031l.056-2.718a3.25 3.25 0 0 1 .801-2.069z"/><path fill="currentColor" fill-rule="evenodd" d="M6.237 7.7a4.214 4.214 0 0 1 4.206-3.95H11V3a1 1 0 1 1 2 0v.75h.557a4.214 4.214 0 0 1 4.206 3.95l.221 3.534a7.376 7.376 0 0 0 1.308 3.754a1.617 1.617 0 0 1-1.135 2.529l-3.407.408V19a2.75 2.75 0 1 1-5.5 0v-1.075l-3.407-.409a1.617 1.617 0 0 1-1.135-2.528a7.377 7.377 0 0 0 1.308-3.754zM10.75 19a1.25 1.25 0 0 0 2.5 0v-.75h-2.5z" clip-rule="evenodd"/><path fill="currentColor" d="M17.643 1.355a.75.75 0 0 0-.072 1.058l1.292 1.48a3.25 3.25 0 0 1 .8 2.07l.057 2.717a.75.75 0 1 0 1.5-.031l-.057-2.718a4.75 4.75 0 0 0-1.17-3.024l-1.292-1.48a.75.75 0 0 0-1.058-.072"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13 3a1 1 0 1 0-2 0v.75h-.557A4.214 4.214 0 0 0 6.237 7.7l-.221 3.534a7.377 7.377 0 0 1-1.308 3.754a1.617 1.617 0 0 0 1.135 2.529l3.407.408V19a2.75 2.75 0 1 0 5.5 0v-1.075l3.407-.409a1.617 1.617 0 0 0 1.135-2.528a7.376 7.376 0 0 1-1.308-3.754l-.221-3.533a4.214 4.214 0 0 0-4.206-3.951H13zm-2.557 2.25a2.714 2.714 0 0 0-2.709 2.544l-.22 3.534a8.877 8.877 0 0 1-1.574 4.516a.117.117 0 0 0 .082.183l3.737.449c1.489.178 2.993.178 4.482 0l3.737-.449a.117.117 0 0 0 .082-.183a8.877 8.877 0 0 1-1.573-4.516l-.221-3.534a2.714 2.714 0 0 0-2.709-2.544zm1.557 15c-.69 0-1.25-.56-1.25-1.25v-.75h2.5V19c0 .69-.56 1.25-1.25 1.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13 3a1 1 0 1 0-2 0v.75h-.557A4.214 4.214 0 0 0 6.237 7.7l-.221 3.534a7.377 7.377 0 0 1-1.308 3.754a1.617 1.617 0 0 0 1.135 2.529l3.407.408V19a2.75 2.75 0 1 0 5.5 0v-1.075l3.407-.409a1.617 1.617 0 0 0 1.135-2.528a7.376 7.376 0 0 1-1.308-3.754l-.221-3.533a4.214 4.214 0 0 0-4.206-3.951H13zm-2.25 16a1.25 1.25 0 1 0 2.5 0v-.75h-2.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotionOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.575 7.25c.062.345-.018.68-.338.711l-.532.096l-.22 8.317c-.468.25-.894.388-1.245.375c-.567-.02-.704-.215-1.11-.795l-3.306-5.947l-.15 5.636l1.088.304s-.018.681-.9.65l-2.44.068c-.066-.155.015-.532.265-.596l.64-.169l.197-7.45l-.88-.106c-.062-.345.129-.829.622-.85l2.617-.099l3.445 6.027l.139-5.221l-.916-.147a.62.62 0 0 1 .584-.734z"/><path fill="currentColor" fill-rule="evenodd" d="M17.258 2.833a47.721 47.721 0 0 0-10.516 0c-2.012.225-3.637 1.81-3.873 3.832a45.921 45.921 0 0 0 0 10.67c.236 2.022 1.86 3.607 3.873 3.832a47.77 47.77 0 0 0 10.516 0c2.012-.225 3.637-1.81 3.873-3.832a45.925 45.925 0 0 0 0-10.67c-.236-2.022-1.86-3.607-3.873-3.832m-10.35 1.49a46.22 46.22 0 0 1 10.184 0c1.33.15 2.395 1.199 2.55 2.517a44.421 44.421 0 0 1 0 10.32a2.89 2.89 0 0 1-2.55 2.516a46.216 46.216 0 0 1-10.184 0a2.89 2.89 0 0 1-2.55-2.516a44.421 44.421 0 0 1 0-10.32a2.89 2.89 0 0 1 2.55-2.516" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotionSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.77 3.082a47.472 47.472 0 0 1 10.46 0c1.899.212 3.43 1.707 3.653 3.613a45.67 45.67 0 0 1 0 10.61c-.223 1.906-1.754 3.401-3.652 3.614a47.468 47.468 0 0 1-10.461 0c-1.899-.213-3.43-1.708-3.653-3.613a45.672 45.672 0 0 1 0-10.611C3.34 4.789 4.871 3.294 6.77 3.082m9.805 4.168c.062.345-.018.68-.338.711l-.532.096l-.22 8.317c-.468.25-.894.388-1.245.375c-.567-.02-.704-.215-1.11-.795l-3.306-5.947l-.15 5.636l1.088.304s-.018.681-.9.65l-2.44.068c-.066-.155.015-.532.265-.596l.64-.169l.197-7.45l-.88-.106c-.062-.345.129-.829.622-.85l2.617-.099l3.445 6.027l.139-5.221l-.916-.147a.62.62 0 0 1 .584-.734z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OtherOneOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 10.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m4.5 1.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0m6 0a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OtherOneSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 10.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m4.5 1.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0m6 0a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OtherTwoOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" stroke="currentColor" stroke-width="1.5" d="M12 5.25a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5Zm0 6a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5Zm0 6a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaletteOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.89 15.403a1 1 0 1 0 .243 1.985a1 1 0 0 0-.243-1.985m-3.342-2.614a1 1 0 1 1 1.985-.243a1 1 0 0 1-1.985.243m2.371-5.327a1 1 0 1 0 .243 1.985a1 1 0 0 0-.243-1.985m7.778 1.819a1 1 0 1 1 1.986-.243a1 1 0 0 1-1.986.243m-3.704 1.598a1 1 0 1 0 .242 1.985a1 1 0 0 0-.242-1.985m-.6-3.842a1 1 0 1 1 1.985-.243a1 1 0 0 1-1.985.243"/><path fill="currentColor" fill-rule="evenodd" d="M2.253 11.764c.13-5.384 4.6-9.642 9.983-9.511c5.383.13 9.642 4.6 9.511 9.983a1.465 1.465 0 0 1-.334.91a1.834 1.834 0 0 1-.721.513c-.502.21-1.131.275-1.754.28c-2.689.018-5.024 2.33-5.177 5.058c-.035.636-.144 1.283-.418 1.792a1.872 1.872 0 0 1-.62.695a1.662 1.662 0 0 1-.96.263c-5.382-.13-9.64-4.6-9.51-9.983M12.2 3.752a8.25 8.25 0 0 0-.4 16.496c.068.001.093-.014.106-.021a.397.397 0 0 0 .117-.15c.117-.217.21-.601.24-1.164c.195-3.475 3.142-6.45 6.665-6.475c.562-.003.956-.068 1.186-.163a.48.48 0 0 0 .134-.076A8.25 8.25 0 0 0 12.2 3.752" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaletteSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.23 2.503a9.5 9.5 0 0 0-.46 18.994c.338.008.643-.086.898-.28c.248-.187.42-.444.539-.72c.235-.54.314-1.249.318-1.98c.015-2.641 2.332-4.935 4.948-4.836c.72.027 1.413-.012 1.948-.188c.27-.09.534-.225.736-.436c.212-.221.332-.503.34-.827a9.5 9.5 0 0 0-9.267-9.727m-3.34 12.9a1 1 0 1 0 .243 1.985a1 1 0 0 0-.243-1.985m-3.342-2.614a1 1 0 1 1 1.985-.243a1 1 0 0 1-1.985.243M7.92 7.462a1 1 0 1 0 .243 1.985a1 1 0 0 0-.243-1.985m7.778 1.819a1 1 0 1 1 1.986-.243a1 1 0 0 1-1.986.243m-3.704 1.598a1 1 0 1 0 .242 1.985a1 1 0 0 0-.242-1.985m-.6-3.842a1 1 0 1 1 1.985-.243a1 1 0 0 1-1.985.243" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M18.535 4.766c.73.27 1.215.965 1.215 1.743V17.49c0 .778-.485 1.474-1.215 1.743a4.44 4.44 0 0 1-3.07 0a1.857 1.857 0 0 1-1.215-1.743V6.51c0-.778.485-1.474 1.215-1.743a4.44 4.44 0 0 1 3.07 0M18.25 6.51a.357.357 0 0 0-.234-.335a2.94 2.94 0 0 0-2.032 0a.357.357 0 0 0-.234.335v10.98c0 .15.093.284.234.335a2.94 2.94 0 0 0 2.032 0a.357.357 0 0 0 .234-.335zM8.535 4.766c.73.27 1.215.965 1.215 1.743V17.49c0 .778-.485 1.474-1.215 1.743a4.44 4.44 0 0 1-3.07 0A1.857 1.857 0 0 1 4.25 17.49V6.51c0-.778.485-1.474 1.215-1.743a4.44 4.44 0 0 1 3.07 0M8.25 6.51a.357.357 0 0 0-.234-.335a2.94 2.94 0 0 0-2.032 0a.357.357 0 0 0-.234.335v10.98c0 .15.093.284.234.335a2.94 2.94 0 0 0 2.032 0a.357.357 0 0 0 .234-.335z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.276 5.47c.435.16.724.575.724 1.039V17.49c0 .464-.29.879-.724 1.039a3.69 3.69 0 0 1-2.552 0A1.107 1.107 0 0 1 14 17.491V6.51c0-.464.29-.879.724-1.04a3.69 3.69 0 0 1 2.552 0m-8 0c.435.16.724.575.724 1.039V17.49c0 .464-.29.879-.724 1.039a3.69 3.69 0 0 1-2.552 0A1.107 1.107 0 0 1 6 17.491V6.51c0-.464.29-.879.724-1.04a3.69 3.69 0 0 1 2.552 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneInOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.477 10.481a18.716 18.716 0 0 0 9.693 9.416l.758.338a3.75 3.75 0 0 0 4.566-1.23l1.274-1.763a1.75 1.75 0 0 0-.32-2.387l-2.225-1.795a1.75 1.75 0 0 0-2.504.32l-.497.67a11.892 11.892 0 0 1-5.117-5.118l.67-.496c.798-.592.944-1.73.32-2.505L9.298 3.706a1.75 1.75 0 0 0-2.386-.32L5.14 4.666A3.75 3.75 0 0 0 3.92 9.26l.555 1.22zm10.285 8.038A17.216 17.216 0 0 1 5.84 9.856l-.554-1.218a2.25 2.25 0 0 1 .732-2.756l1.773-1.28a.25.25 0 0 1 .34.046l1.796 2.225a.25.25 0 0 1-.046.358l-1.168.866a.75.75 0 0 0-.237.912a13.387 13.387 0 0 0 6.67 6.67a.75.75 0 0 0 .912-.237l.866-1.168a.25.25 0 0 1 .358-.046l2.224 1.795a.25.25 0 0 1 .046.34l-1.274 1.764a2.25 2.25 0 0 1-2.74.738l-.764-.34z" clip-rule="evenodd"/><path fill="currentColor" d="M13.75 9a.75.75 0 0 0 .75.75h3.828a.75.75 0 1 0 0-1.5h-2.017l3.159-3.159a.75.75 0 1 0-1.061-1.06L15.25 7.188V5.172a.75.75 0 0 0-1.5 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneInSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.566 19.152A18.466 18.466 0 0 1 5 9.86v-.002l-.519-1.14a3.5 3.5 0 0 1 1.237-4.355l1.333-.894a1 1 0 0 1 1.335.203l2.43 3.012a1 1 0 0 1-.183 1.431l-1.257.932a12.14 12.14 0 0 0 5.511 5.51l.932-1.256a1 1 0 0 1 1.431-.183l3.012 2.43a1 1 0 0 1 .203 1.335l-.888 1.324a3.5 3.5 0 0 1-4.331 1.247z"/><path fill="currentColor" d="M13.75 9a.75.75 0 0 0 .75.75h3.828a.75.75 0 0 0 0-1.5h-2.017l3.159-3.159a.75.75 0 1 0-1.061-1.06L15.25 7.188V5.172a.75.75 0 0 0-1.5 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneMissOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.664 4.414a.75.75 0 0 1 .75-.75h2.829a.75.75 0 0 1 0 1.5H9.225l2.598 2.599a.25.25 0 0 0 .354 0l3.878-3.88a.75.75 0 0 1 1.061 1.062l-3.879 3.878a1.75 1.75 0 0 1-2.475 0L8.165 6.225v1.018a.75.75 0 1 1-1.5 0zm11.308 9.158a17.217 17.217 0 0 0-12.434.183l-.013.005l-.78.3a2.25 2.25 0 0 0-1.416 2.458l.346 2.148a.25.25 0 0 0 .273.21l2.843-.305a.25.25 0 0 0 .22-.285l-.214-1.439a.75.75 0 0 1 .478-.812a13.388 13.388 0 0 1 9.433 0a.75.75 0 0 1 .477.813l-.214 1.438c-.02.14.08.27.22.285l2.843.304a.25.25 0 0 0 .274-.208l.348-2.16a2.25 2.25 0 0 0-1.43-2.465l-1.253-.47zm-12.99-1.21a18.716 18.716 0 0 1 13.513-.196l.002.001l1.255.47a3.75 3.75 0 0 1 2.385 4.11l-.349 2.159a1.75 1.75 0 0 1-1.913 1.46l-2.843-.303a1.75 1.75 0 0 1-1.545-1.998l.123-.824a11.891 11.891 0 0 0-7.238 0l.123.824a1.75 1.75 0 0 1-1.545 1.998l-2.842.304a1.75 1.75 0 0 1-1.914-1.462l-.346-2.148a3.75 3.75 0 0 1 2.36-4.098z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneMissSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.67 4.75A.75.75 0 0 1 7.42 4h2.828a.75.75 0 1 1 0 1.5H9.231l2.598 2.598a.25.25 0 0 0 .354 0L16.06 4.22a.75.75 0 1 1 1.06 1.06l-3.876 3.88a1.75 1.75 0 0 1-2.475 0L8.17 6.561v1.017a.75.75 0 1 1-1.5 0z" clip-rule="evenodd"/><path fill="currentColor" d="M18.416 12.65a18.466 18.466 0 0 0-13.334.195l-.695.266a3.5 3.5 0 0 0-2.18 3.944l.307 1.564a1 1 0 0 0 1.088.801l3.848-.411a1 1 0 0 0 .882-1.142l-.23-1.547a12.14 12.14 0 0 1 7.794 0l-.23 1.547a1 1 0 0 0 .883 1.142l3.848.411a1 1 0 0 0 1.087-.8l.31-1.575a3.5 3.5 0 0 0-2.204-3.954l-1.172-.44z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneOffOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.994 10.825a.306.306 0 0 0 .03-.394c-.342-.474-.65-.975-.92-1.499l.67-.496c.8-.592.945-1.73.32-2.505L9.3 3.706a1.75 1.75 0 0 0-2.386-.32L5.14 4.666A3.75 3.75 0 0 0 3.92 9.26l.555 1.22l.002.002a18.705 18.705 0 0 0 2.45 3.95c.11.138.315.147.44.022l.64-.64a.304.304 0 0 0 .021-.405A17.206 17.206 0 0 1 5.84 9.856l-.554-1.218a2.25 2.25 0 0 1 .732-2.756l1.773-1.28a.25.25 0 0 1 .34.046l1.796 2.225a.25.25 0 0 1-.046.358l-1.168.866a.75.75 0 0 0-.237.912c.391.864.87 1.678 1.427 2.431a.294.294 0 0 0 .446.03z"/><path fill="currentColor" fill-rule="evenodd" d="M19.997 3.944L4.03 19.91a.936.936 0 0 1-.027.026l-.033.034a.75.75 0 0 0 1.06 1.06l4.208-4.207a18.709 18.709 0 0 0 4.932 3.074l.758.338a3.75 3.75 0 0 0 4.566-1.23l1.274-1.763a1.75 1.75 0 0 0-.32-2.387l-2.225-1.795a1.75 1.75 0 0 0-2.504.32l-.497.67a11.866 11.866 0 0 1-1.95-1.261L21.03 5.03a.75.75 0 0 0-1.033-1.086m-4.85 11.735a13.36 13.36 0 0 1-2.942-1.824l-1.903 1.904a17.205 17.205 0 0 0 4.46 2.76l.012.005l.764.34a2.25 2.25 0 0 0 2.74-.737l1.274-1.763a.25.25 0 0 0-.046-.341l-2.224-1.795a.25.25 0 0 0-.358.046l-.866 1.168a.75.75 0 0 1-.912.237" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneOffSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.498 11.322a.305.305 0 0 0 .026-.398a12.112 12.112 0 0 1-1.148-1.876l1.257-.932a1 1 0 0 0 .183-1.431l-2.43-3.012a1 1 0 0 0-1.335-.203l-1.333.894a3.5 3.5 0 0 0-1.237 4.355L5 9.86a18.454 18.454 0 0 0 2.5 4a.296.296 0 0 0 .439.02zm9.499-7.378L4.03 19.91a.911.911 0 0 1-.027.026l-.033.034a.75.75 0 0 0 1.06 1.06l4.8-4.8a18.454 18.454 0 0 0 4.736 2.922l.68.303a3.5 3.5 0 0 0 4.33-1.247l.889-1.324a1 1 0 0 0-.203-1.335l-3.012-2.43a1 1 0 0 0-1.431.183l-.932 1.257a12.125 12.125 0 0 1-2.08-1.304L21.03 5.03a.75.75 0 0 0-1.033-1.086"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneOutOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M5.84 9.856a17.216 17.216 0 0 0 8.922 8.663l.012.005l.764.34a2.25 2.25 0 0 0 2.74-.737l1.274-1.763a.25.25 0 0 0-.046-.341l-2.224-1.795a.25.25 0 0 0-.358.046l-.866 1.168a.75.75 0 0 1-.912.237a13.387 13.387 0 0 1-6.67-6.67a.75.75 0 0 1 .237-.912L9.88 7.23a.25.25 0 0 0 .046-.358L8.132 4.648a.25.25 0 0 0-.341-.046l-1.773 1.28a2.25 2.25 0 0 0-.732 2.756zm8.33 10.041a18.716 18.716 0 0 1-9.693-9.416l-.002-.002l-.554-1.22A3.75 3.75 0 0 1 5.14 4.666l1.773-1.28a1.75 1.75 0 0 1 2.386.32l1.795 2.225a1.75 1.75 0 0 1-.32 2.505l-.67.496a11.892 11.892 0 0 0 5.118 5.118l.497-.67a1.75 1.75 0 0 1 2.504-.32l2.225 1.795a1.75 1.75 0 0 1 .32 2.387l-1.274 1.764a3.75 3.75 0 0 1-4.566 1.229z"/><path d="M19.47 4.03c.14.141.22.332.22.53v3.83a.75.75 0 0 1-1.5 0V6.37l-3.16 3.16a.75.75 0 1 1-1.06-1.061l3.159-3.16H15.11a.75.75 0 0 1 0-1.5h3.828a.75.75 0 0 1 .53.22"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneOutSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M19.47 4.03c.14.141.22.332.22.53v3.83a.75.75 0 1 1-1.5 0V6.37l-3.16 3.16a.75.75 0 1 1-1.06-1.061l3.159-3.16H15.11a.75.75 0 0 1 0-1.5h3.828a.75.75 0 0 1 .53.22" clip-rule="evenodd"/><path fill="currentColor" d="M5 9.86a18.466 18.466 0 0 0 9.566 9.292l.68.303a3.5 3.5 0 0 0 4.33-1.247l.889-1.324a1 1 0 0 0-.203-1.335l-3.012-2.43a1 1 0 0 0-1.431.183l-.932 1.257a12.14 12.14 0 0 1-5.51-5.511l1.256-.932a1 1 0 0 0 .183-1.431l-2.43-3.012a1 1 0 0 0-1.335-.203l-1.333.894a3.5 3.5 0 0 0-1.237 4.355z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.84 9.856a17.216 17.216 0 0 0 8.922 8.663l.012.005l.764.34a2.25 2.25 0 0 0 2.74-.737l1.274-1.763a.25.25 0 0 0-.046-.341l-2.224-1.795a.25.25 0 0 0-.358.046l-.866 1.168a.75.75 0 0 1-.912.237a13.387 13.387 0 0 1-6.67-6.67a.75.75 0 0 1 .237-.912L9.88 7.23a.25.25 0 0 0 .046-.358L8.132 4.648a.25.25 0 0 0-.341-.046l-1.773 1.28a2.25 2.25 0 0 0-.732 2.756zm8.33 10.041a18.716 18.716 0 0 1-9.693-9.416l-.002-.002l-.554-1.22A3.75 3.75 0 0 1 5.14 4.666l1.773-1.28a1.75 1.75 0 0 1 2.386.32l1.795 2.225a1.75 1.75 0 0 1-.32 2.505l-.67.496a11.892 11.892 0 0 0 5.118 5.118l.497-.67a1.75 1.75 0 0 1 2.504-.32l2.225 1.795a1.75 1.75 0 0 1 .32 2.387l-1.274 1.764a3.75 3.75 0 0 1-4.566 1.229z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 9.86a18.466 18.466 0 0 0 9.566 9.292l.68.303a3.5 3.5 0 0 0 4.33-1.247l.889-1.324a1 1 0 0 0-.203-1.335l-3.012-2.43a1 1 0 0 0-1.431.183l-.932 1.257a12.14 12.14 0 0 1-5.51-5.511l1.256-.932a1 1 0 0 0 .183-1.431l-2.43-3.012a1 1 0 0 0-1.335-.203l-1.333.894a3.5 3.5 0 0 0-1.237 4.355z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PictureOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 11.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0m-1.864 1.603a.75.75 0 0 0-1.272 0l-2.5 4A.75.75 0 0 0 8 18.25h8a.75.75 0 0 0 .6-1.2l-1.5-2a.75.75 0 0 0-1.2 0l-.844 1.125z"/><path fill="currentColor" fill-rule="evenodd" d="M7 2.25A2.75 2.75 0 0 0 4.25 5v14A2.75 2.75 0 0 0 7 21.75h10A2.75 2.75 0 0 0 19.75 19V7.968c0-.381-.124-.751-.354-1.055l-2.998-3.968a1.75 1.75 0 0 0-1.396-.695zM5.75 5c0-.69.56-1.25 1.25-1.25h7.25v4.397c0 .414.336.75.75.75h3.25V19c0 .69-.56 1.25-1.25 1.25H7c-.69 0-1.25-.56-1.25-1.25z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PictureSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.25 2.5a.25.25 0 0 0-.25-.25H7A2.75 2.75 0 0 0 4.25 5v14A2.75 2.75 0 0 0 7 21.75h10A2.75 2.75 0 0 0 19.75 19V9.147a.25.25 0 0 0-.25-.25H15a.75.75 0 0 1-.75-.75zm-1.25 9a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0m-1.864 1.602a.75.75 0 0 0-1.272 0l-2.5 4A.75.75 0 0 0 8 18.25h8a.75.75 0 0 0 .6-1.2l-1.5-2a.75.75 0 0 0-1.2 0l-.844 1.125z" clip-rule="evenodd"/><path fill="currentColor" d="M15.75 2.824c0-.184.193-.301.336-.186c.121.098.23.212.323.342l3.013 4.197c.068.096-.006.22-.124.22H16a.25.25 0 0 1-.25-.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.786 3.725a1.75 1.75 0 0 0-2.846.548L12.347 7.99A4.745 4.745 0 0 0 8.07 9.291l-1.71 1.71a.75.75 0 0 0 0 1.06l2.495 2.496l-5.385 5.386a.75.75 0 1 0 1.06 1.06l5.386-5.385l2.495 2.495a.75.75 0 0 0 1.061 0l1.71-1.71a4.745 4.745 0 0 0 1.302-4.277l3.716-1.592a1.75 1.75 0 0 0 .548-2.846zm-1.468 1.139a.25.25 0 0 1 .407-.078l3.963 3.962a.25.25 0 0 1-.079.407l-4.315 1.85a.75.75 0 0 0-.41.941a3.25 3.25 0 0 1-.763 3.396l-1.18 1.18l-4.99-4.99l1.18-1.18a3.25 3.25 0 0 1 3.396-.762a.75.75 0 0 0 .942-.41z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.432 4.079a1.25 1.25 0 0 0-2.033.391l-1.76 4.105a4.25 4.25 0 0 0-4.215 1.07L7.067 11a.75.75 0 0 0 0 1.06l2.142 2.143l-5.74 5.739a.75.75 0 1 0 1.061 1.06l5.74-5.739l2.141 2.142a.75.75 0 0 0 1.06 0l1.358-1.356a4.25 4.25 0 0 0 1.069-4.217l4.105-1.76a1.25 1.25 0 0 0 .392-2.032z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinterestOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75s9.75-4.365 9.75-9.75S17.385 2.25 12 2.25M3.75 12a8.25 8.25 0 1 1 5.789 7.877a9.236 9.236 0 0 0 1-2.128l.516-1.977a2.232 2.232 0 0 0 1.906.95c2.535 0 4.317-2.304 4.317-5.166c0-2.749-2.245-4.803-5.124-4.803c-3.587 0-5.492 2.41-5.492 5.029c0 1.223.647 2.737 1.686 3.224c.154.071.255 0 .279-.113c.024-.113.166-.682.231-.944a.263.263 0 0 0-.059-.237a3.218 3.218 0 0 1-.594-1.894a3.598 3.598 0 0 1 3.753-3.61a3.283 3.283 0 0 1 3.473 3.384c0 2.244-1.14 3.806-2.612 3.806a1.188 1.188 0 0 1-1.235-1.502c.084-.35.196-.711.304-1.063c.197-.638.385-1.245.385-1.704a1.046 1.046 0 0 0-1.051-1.188c-.832 0-1.503.867-1.503 2.019c-.006.407.07.811.226 1.188l-.98 4.156a7.86 7.86 0 0 0 0 2.37a8.231 8.231 0 0 1-1.818-1.002a8.312 8.312 0 0 1-2.964-4.027A8.24 8.24 0 0 1 3.75 12" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinterestSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.935 2.375a9.625 9.625 0 0 0-3.507 18.563a9.102 9.102 0 0 1 0-2.75l1.135-4.813A3.472 3.472 0 0 1 9.3 12c0-1.334.777-2.338 1.74-2.338a1.21 1.21 0 0 1 1.217 1.375c0 .532-.218 1.236-.446 1.974c-.125.407-.254.825-.352 1.23a1.375 1.375 0 0 0 1.43 1.74c1.705 0 3.025-1.809 3.025-4.407a3.804 3.804 0 0 0-4.022-3.919a4.166 4.166 0 0 0-4.345 4.18a3.726 3.726 0 0 0 .688 2.193a.303.303 0 0 1 .068.275c-.075.303-.24.963-.268 1.093c-.027.13-.144.213-.323.13c-1.203-.563-1.952-2.316-1.952-3.732c0-3.032 2.207-5.823 6.36-5.823c3.333 0 5.932 2.378 5.932 5.562c0 3.313-2.062 5.98-4.998 5.98a2.586 2.586 0 0 1-2.207-1.1l-.598 2.29a10.69 10.69 0 0 1-1.196 2.523a9.74 9.74 0 0 0 2.88.399a9.625 9.625 0 0 0 0-19.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M19.266 10.484a1.917 1.917 0 0 1 0 3.032a35.764 35.764 0 0 1-9.916 5.416l-.653.232c-1.248.443-2.566-.401-2.735-1.69a42.49 42.49 0 0 1 0-10.948c.169-1.289 1.487-2.133 2.735-1.69l.653.232a35.765 35.765 0 0 1 9.916 5.416m-.918 1.846a.417.417 0 0 0 0-.66a34.262 34.262 0 0 0-9.5-5.189l-.653-.232a.572.572 0 0 0-.746.472a40.99 40.99 0 0 0 0 10.558a.572.572 0 0 0 .746.471l.653-.231a34.263 34.263 0 0 0 9.5-5.19" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaySolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.266 13.516a1.917 1.917 0 0 0 0-3.032A35.762 35.762 0 0 0 9.35 5.068l-.653-.232c-1.248-.443-2.567.401-2.736 1.69a42.49 42.49 0 0 0 0 10.948c.17 1.289 1.488 2.133 2.736 1.69l.653-.232a35.762 35.762 0 0 0 9.916-5.416"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12 17V7m-5 5h10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.75 7a.75.75 0 0 0-1.5 0v4.25H7a.75.75 0 0 0 0 1.5h4.25V17a.75.75 0 0 0 1.5 0v-4.25H17a.75.75 0 0 0 0-1.5h-4.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PowerButtonOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 8.082c0-.183-.19-.302-.348-.212a4.75 4.75 0 1 0 4.696 0c-.159-.09-.348.03-.348.212v1.234c0 .077.036.15.095.199a3.25 3.25 0 1 1-4.19 0A.261.261 0 0 0 10 9.316z"/><path fill="currentColor" d="M12.75 7a.75.75 0 0 0-1.5 0v5a.75.75 0 0 0 1.5 0z"/><path fill="currentColor" fill-rule="evenodd" d="M3.25 12a8.75 8.75 0 1 1 17.5 0a8.75 8.75 0 0 1-17.5 0M12 4.75a7.25 7.25 0 1 0 0 14.5a7.25 7.25 0 0 0 0-14.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PowerButtonSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.25 12a8.75 8.75 0 1 1 17.5 0a8.75 8.75 0 0 1-17.5 0M12 6.25a.75.75 0 0 1 .75.75v5a.75.75 0 0 1-1.5 0V7a.75.75 0 0 1 .75-.75m-2 1.832c0-.183-.19-.302-.348-.212a4.75 4.75 0 1 0 4.696 0c-.159-.09-.348.03-.348.212v1.234c0 .077.036.15.095.199a3.25 3.25 0 1 1-4.19 0A.261.261 0 0 0 10 9.316z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PresentOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.25 5.5A3.25 3.25 0 0 1 12 3.423a3.25 3.25 0 0 1 5.24 3.827H18A2.75 2.75 0 0 1 20.75 10v2a1.75 1.75 0 0 1-1.281 1.687c.144 1.826.06 3.665-.25 5.473a2.459 2.459 0 0 1-2.15 2.028l-.915.102a37.402 37.402 0 0 1-8.309 0l-.914-.102a2.459 2.459 0 0 1-2.15-2.028a22.111 22.111 0 0 1-.25-5.473A1.75 1.75 0 0 1 3.25 12v-2A2.75 2.75 0 0 1 6 7.25h.76a3.235 3.235 0 0 1-.51-1.75m5 0a1.75 1.75 0 1 0-3.5 0a1.75 1.75 0 0 0 3.5 0m3.25 1.75a1.75 1.75 0 1 0 0-3.5a1.75 1.75 0 0 0 0 3.5M4.75 10c0-.69.56-1.25 1.25-1.25h5.25v3.5H5a.25.25 0 0 1-.25-.25zm8 3.75h5.219c.14 1.72.064 3.453-.228 5.156a.959.959 0 0 1-.839.791l-.914.103c-1.076.12-2.157.191-3.238.214zm0-1.5H19a.25.25 0 0 0 .25-.25v-2c0-.69-.56-1.25-1.25-1.25h-5.25zm-1.5 1.5v6.264a35.905 35.905 0 0 1-3.238-.214l-.914-.103a.959.959 0 0 1-.839-.79a20.612 20.612 0 0 1-.228-5.157z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PresentSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.25 5.5A3.25 3.25 0 0 1 12 3.423A3.25 3.25 0 0 1 17.062 7.5H18a2.5 2.5 0 0 1 2.5 2.5v1.25a.75.75 0 0 1-.75.75h-6.7a.3.3 0 0 1-.3-.3V8.24a3.267 3.267 0 0 1-.75-.663a3.267 3.267 0 0 1-.75.662V11.7a.3.3 0 0 1-.3.3h-6.7a.75.75 0 0 1-.75-.75V10A2.5 2.5 0 0 1 6 7.5h.938a3.236 3.236 0 0 1-.688-2m5 0a1.75 1.75 0 1 0-3.5 0a1.75 1.75 0 0 0 3.5 0m1.5 0a1.75 1.75 0 1 0 3.5 0a1.75 1.75 0 0 0-3.5 0" clip-rule="evenodd"/><path fill="currentColor" d="M11.25 13.65a.3.3 0 0 0-.3-.3H5.649a.833.833 0 0 0-.82.692a11.592 11.592 0 0 0 0 3.916l.224 1.309a2.008 2.008 0 0 0 1.755 1.656l1.065.119a37.15 37.15 0 0 0 3.071.215a.298.298 0 0 0 .306-.299zm1.806 7.607a.298.298 0 0 1-.306-.299V13.65a.3.3 0 0 1 .3-.3h5.301c.406 0 .752.292.82.692c.223 1.296.223 2.62 0 3.916l-.223 1.309a2.008 2.008 0 0 1-1.756 1.656l-1.065.119a37.177 37.177 0 0 1-3.071.215"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PrinterOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.75 11a1.25 1.25 0 1 0 0 2.5a1.25 1.25 0 0 0 0-2.5"/><path fill="currentColor" fill-rule="evenodd" d="m4.173 8.115l2.077-.23V3.912c0-.643.49-1.18 1.13-1.24a50.255 50.255 0 0 1 9.24 0c.64.06 1.13.597 1.13 1.24v3.971l2.077.231a1.842 1.842 0 0 1 1.584 1.387a14.4 14.4 0 0 1 0 6.94l-.183.739a.75.75 0 0 1-.728.569h-2.75v2.333a1.21 1.21 0 0 1-1.09 1.205c-3.099.31-6.22.31-9.32 0a1.21 1.21 0 0 1-1.09-1.205V17.75H3.5a.75.75 0 0 1-.728-.57l-.183-.738a14.4 14.4 0 0 1 0-6.94a1.842 1.842 0 0 1 1.584-1.387M16.25 7.74a52.005 52.005 0 0 0-8.5 0V4.146a48.756 48.756 0 0 1 8.5 0zm0 8.762a.37.37 0 0 0-.247-.345a45.342 45.342 0 0 0-8.006 0a.37.37 0 0 0-.247.349v3.314a45.2 45.2 0 0 0 8.5 0zm3.663-.252H17.75v-.333a1.21 1.21 0 0 0-1.09-1.205l-.407-.039a1.858 1.858 0 0 0-.173-.027a38.091 38.091 0 0 0-8.16 0a1.887 1.887 0 0 0-.173.027l-.407.04a1.21 1.21 0 0 0-1.09 1.204v.333H4.087l-.042-.169a12.9 12.9 0 0 1 0-6.217a.342.342 0 0 1 .294-.258l2.082-.231a50.501 50.501 0 0 1 11.158 0l2.082.231c.143.016.26.119.294.258a12.902 12.902 0 0 1 0 6.217z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PrinterSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m4.2 8.364l2.05-.228V3.913c0-.643.49-1.18 1.13-1.24a50.256 50.256 0 0 1 9.24 0c.64.06 1.13.597 1.13 1.24v4.223l2.05.228a1.592 1.592 0 0 1 1.369 1.198a14.15 14.15 0 0 1 0 6.82l-.153.615a.662.662 0 0 1-.643.503H17.75v2.583a1.21 1.21 0 0 1-1.09 1.205c-3.099.31-6.221.31-9.32 0a1.21 1.21 0 0 1-1.09-1.205V17.5H3.627a.662.662 0 0 1-.643-.503l-.153-.615a14.15 14.15 0 0 1 0-6.82a1.592 1.592 0 0 1 1.37-1.198m3.55-.374a51.753 51.753 0 0 1 8.5 0V4.146a48.756 48.756 0 0 0-8.5 0zm0 8.19v3.64a45.2 45.2 0 0 0 8.5 0v-3.64a45.318 45.318 0 0 0-8.5 0m10-5.18a1.25 1.25 0 1 0 0 2.5a1.25 1.25 0 0 0 0-2.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ProcessorOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M13.808 8.447a16.407 16.407 0 0 0-3.616 0c-.898.1-1.626.808-1.732 1.717a15.808 15.808 0 0 0 0 3.672c.106.91.834 1.616 1.732 1.717c1.192.133 2.424.133 3.616 0a1.963 1.963 0 0 0 1.732-1.717c.143-1.22.143-2.452 0-3.672a1.963 1.963 0 0 0-1.732-1.717m-3.45 1.491c1.082-.12 2.202-.12 3.284 0a.464.464 0 0 1 .409.4c.129 1.104.129 2.22 0 3.324a.464.464 0 0 1-.41.4a14.92 14.92 0 0 1-3.283 0a.464.464 0 0 1-.409-.4a14.324 14.324 0 0 1 0-3.324a.464.464 0 0 1 .41-.4"/><path d="M15.75 3a.75.75 0 0 0-1.5 0v2.524a32.03 32.03 0 0 0-1.5-.07V4a.75.75 0 0 0-1.5 0v1.454a32.03 32.03 0 0 0-1.5.07V3a.75.75 0 0 0-1.5 0v2.672A3.178 3.178 0 0 0 5.688 8.25H3a.75.75 0 0 0 0 1.5h2.537c-.036.5-.06 1-.073 1.5H4a.75.75 0 0 0 0 1.5h1.464c.013.5.037 1 .073 1.5H3a.75.75 0 0 0 0 1.5h2.688a3.178 3.178 0 0 0 2.562 2.578V21a.75.75 0 0 0 1.5 0v-2.524c.498.035.999.059 1.5.07V20a.75.75 0 0 0 1.5 0v-1.454a31.975 31.975 0 0 0 1.5-.07V21a.75.75 0 0 0 1.5 0v-2.672a3.178 3.178 0 0 0 2.562-2.578H21a.75.75 0 0 0 0-1.5h-2.537c.036-.5.06-1 .073-1.5H20a.75.75 0 0 0 0-1.5h-1.464c-.013-.5-.037-1-.073-1.5H21a.75.75 0 0 0 0-1.5h-2.688a3.178 3.178 0 0 0-2.562-2.578zM8.633 7.13a30.565 30.565 0 0 1 6.734 0c.773.087 1.39.698 1.479 1.459a29.362 29.362 0 0 1 0 6.822a1.677 1.677 0 0 1-1.48 1.458a30.574 30.574 0 0 1-6.733 0a1.677 1.677 0 0 1-1.479-1.458a29.365 29.365 0 0 1 0-6.822a1.677 1.677 0 0 1 1.48-1.458"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ProcessorSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.358 9.938c1.082-.12 2.202-.12 3.284 0a.464.464 0 0 1 .409.4c.129 1.104.129 2.22 0 3.324a.464.464 0 0 1-.41.4a14.92 14.92 0 0 1-3.283 0a.464.464 0 0 1-.409-.4a14.324 14.324 0 0 1 0-3.324a.464.464 0 0 1 .41-.4"/><path fill="currentColor" fill-rule="evenodd" d="M15 2.25a.75.75 0 0 1 .75.75v2.927a2.929 2.929 0 0 1 2.308 2.323H21a.75.75 0 0 1 0 1.5h-2.788c.037.5.061 1 .073 1.5H20a.75.75 0 0 1 0 1.5h-1.715c-.012.5-.036 1-.073 1.5H21a.75.75 0 0 1 0 1.5h-2.942a2.929 2.929 0 0 1-2.308 2.323V21a.75.75 0 0 1-1.5 0v-2.774c-.498.035-.999.059-1.5.07V20a.75.75 0 0 1-1.5 0v-1.704a31.963 31.963 0 0 1-1.5-.07V21a.75.75 0 0 1-1.5 0v-2.927a2.929 2.929 0 0 1-2.308-2.323H3a.75.75 0 0 1 0-1.5h2.788c-.037-.5-.061-1-.074-1.5H4a.75.75 0 0 1 0-1.5h1.714c.013-.5.037-1 .074-1.5H3a.75.75 0 0 1 0-1.5h2.942A2.929 2.929 0 0 1 8.25 5.927V3a.75.75 0 0 1 1.5 0v2.774c.498-.035.999-.059 1.5-.07V4a.75.75 0 0 1 1.5 0v1.704c.501.011 1.002.035 1.5.07V3a.75.75 0 0 1 .75-.75m-1.192 6.197a16.407 16.407 0 0 0-3.616 0c-.898.1-1.626.808-1.732 1.717a15.808 15.808 0 0 0 0 3.672c.106.91.834 1.616 1.732 1.717c1.192.133 2.424.133 3.616 0a1.963 1.963 0 0 0 1.732-1.717c.143-1.22.143-2.452 0-3.672a1.963 1.963 0 0 0-1.732-1.717" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PulseOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.24 4.876a.75.75 0 0 0-1.454-.106l-2.09 6.48H2a.75.75 0 0 0 0 1.5h3.242a.75.75 0 0 0 .714-.52L7.27 8.16l1.84 10.965a.75.75 0 0 0 1.46.083l2.43-8.48l1.28 4.48a.75.75 0 0 0 1.368.172l1.54-2.629h1.166a2.751 2.751 0 1 0 0-1.5h-1.595a.75.75 0 0 0-.648.37l-.87 1.488l-1.519-5.314a.75.75 0 0 0-1.442 0l-2.239 7.814zM19.75 12a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PulseSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.555 4.252a.75.75 0 0 1 .685.624l1.8 10.732l2.239-7.815a.75.75 0 0 1 1.442.001l1.518 5.314l.871-1.487a.75.75 0 0 1 .648-.371h1.857a2.5 2.5 0 1 1 0 1.5h-1.428l-1.54 2.63a.75.75 0 0 1-1.368-.174l-1.28-4.48l-2.43 8.48a.75.75 0 0 1-1.46-.082L7.269 8.16l-1.313 4.07a.75.75 0 0 1-.714.52H2a.75.75 0 0 1 0-1.5h2.696l2.09-6.48a.75.75 0 0 1 .769-.518"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QqOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.01 22.597c-.737.076-2.183.15-3.5.153a20.016 20.016 0 0 1-1.893-.07a4.86 4.86 0 0 1-.728-.12a1.74 1.74 0 0 1-.34-.123a.984.984 0 0 1-.41-.37a1.098 1.098 0 0 1-.152-.684c.026-.235.123-.426.226-.569c.192-.268.476-.462.707-.597c.102-.059.21-.114.319-.167a6.5 6.5 0 0 1-1.051-1.818a5.818 5.818 0 0 1-.379.431a2.38 2.38 0 0 1-.373.317c-.108.071-.369.229-.702.188a.868.868 0 0 1-.611-.39a1.342 1.342 0 0 1-.171-.39a3.28 3.28 0 0 1-.091-.908c.013-.72.174-1.76.613-3.173c.34-1.092.719-2.032 1.229-3.296l.007-.017l.296-.734c-.039-2.318.42-4.526 1.529-6.19c1.156-1.734 2.98-2.815 5.466-2.82h.018c2.486.005 4.31 1.086 5.466 2.82c1.108 1.664 1.568 3.872 1.53 6.19c.103.259.201.502.295.734l.007.018c.51 1.263.89 2.203 1.229 3.295c.439 1.412.6 2.452.613 3.173c.007.356-.023.663-.09.908c-.034.12-.085.26-.172.39a.868.868 0 0 1-.611.39c-.333.04-.594-.117-.702-.188a2.376 2.376 0 0 1-.373-.317a5.809 5.809 0 0 1-.379-.431a6.498 6.498 0 0 1-1.05 1.818c.108.052.215.108.317.167c.232.135.515.33.708.598c.103.142.2.333.226.567c.027.246-.032.481-.152.685a.983.983 0 0 1-.41.37a1.74 1.74 0 0 1-.34.124a4.862 4.862 0 0 1-.728.119c-.531.052-1.206.071-1.893.07c-1.317-.002-2.763-.077-3.5-.153m-5.502-12.22c-.057-2.19.372-4.118 1.275-5.475c.877-1.316 2.235-2.15 4.227-2.152c1.992.003 3.35.836 4.227 2.152c.903 1.357 1.332 3.285 1.275 5.476a.75.75 0 0 0 .053.298l.354.88c.518 1.281.875 2.168 1.194 3.196c.256.822.401 1.48.477 1.987c-.274-.389-.49-.733-.51-.765a.75.75 0 0 0-.738-.347h-.002a.75.75 0 0 0-.647.742c-.002.996-.531 2.384-1.727 3.398a.75.75 0 0 0-.068 1.08v.001a.76.76 0 0 0 .332.21c.141.043.311.1.486.167c-.36.018-.77.026-1.204.025c-1.325-.002-2.86-.087-3.502-.157c-.643.07-2.177.155-3.502.157a24.8 24.8 0 0 1-1.205-.025a7.76 7.76 0 0 1 .487-.167a.75.75 0 0 0 .331-.21h.001a.75.75 0 0 0 .186-.636v-.001a.75.75 0 0 0-.254-.444c-1.196-1.014-1.725-2.402-1.727-3.398a.75.75 0 0 0-.647-.742h-.002a.75.75 0 0 0-.739.347c-.02.032-.235.376-.509.765a13.48 13.48 0 0 1 .477-1.987c.319-1.028.676-1.914 1.194-3.196l.354-.88a.75.75 0 0 0 .053-.298" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QqSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.194 18.424c-.419.05-1.576-1.777-1.735-2.031c-.005-.009-.018-.004-.018.006c-.012 1.215-.639 2.787-1.98 3.93a.009.009 0 0 0 .003.016c.666.207 2.122.749 1.772 1.34c-.285.482-4.902.308-6.235.157c-1.333.15-5.95.324-6.236-.157c-.353-.595 1.126-1.14 1.786-1.344v-.002c-1.36-1.153-1.991-2.748-1.992-3.971v-.001h-.002c-.05.08-1.309 2.11-1.75 2.057c-.21-.026-.486-1.158.364-3.894c.4-1.29.86-2.363 1.568-4.132C5.62 5.83 7.507 2 12 2c4.443 0 6.373 3.755 6.261 8.397c.707 1.766 1.168 2.845 1.568 4.132c.85 2.736.576 3.869.366 3.894"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedditOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.985 13.5a1.447 1.447 0 1 1-2.895 0a1.447 1.447 0 0 1 2.895 0m-8.528 1.447a1.447 1.447 0 1 0 0-2.894a1.447 1.447 0 0 0 0 2.894m.393 1.682a.75.75 0 0 0-.8 1.268l.345.218a6.75 6.75 0 0 0 7.202 0l.345-.217a.75.75 0 1 0-.8-1.27l-.345.218a5.25 5.25 0 0 1-5.602 0z"/><path fill="currentColor" fill-rule="evenodd" d="M19.96 1.75c-.873 0-1.636.479-2.038 1.188l-3.6-.739a.75.75 0 0 0-.841.442c-.411.966-1.027 2.431-1.542 3.685c-.182.444-.352.863-.495 1.225c-2.047.082-3.943.617-5.494 1.484a3.066 3.066 0 1 0-3.762 4.421c-.132.48-.201.977-.201 1.491c0 2.142 1.209 4.012 3.025 5.317c1.818 1.307 4.29 2.091 6.988 2.091c2.697 0 5.17-.784 6.988-2.09c1.816-1.306 3.025-3.176 3.025-5.318c0-.514-.07-1.012-.2-1.49a3.066 3.066 0 1 0-3.762-4.422c-1.427-.797-3.145-1.314-5.003-1.456c.088-.216.181-.446.28-.684c.418-1.021.906-2.185 1.294-3.103l3.018.62a2.343 2.343 0 1 0 2.32-2.662m-.842 2.342a.842.842 0 1 1 1.685 0a.842.842 0 0 1-1.685 0M12 9.04c.154 0 .307.003.46.009a.621.621 0 0 1-.012.094l-.002.005l.003-.01l.029-.088c2.224.086 4.201.77 5.634 1.8c1.537 1.104 2.401 2.563 2.401 4.098c0 1.536-.865 2.995-2.4 4.1c-1.536 1.102-3.694 1.808-6.113 1.808c-2.419 0-4.577-.706-6.112-1.809c-1.537-1.104-2.401-2.563-2.401-4.099c0-1.535.864-2.994 2.4-4.098C7.424 9.744 9.582 9.039 12 9.039M1.75 10.605a1.566 1.566 0 0 1 2.947-.738c-.799.63-1.463 1.373-1.934 2.204a1.567 1.567 0 0 1-1.013-1.466m19.487 1.466c-.471-.83-1.135-1.575-1.934-2.204a1.566 1.566 0 1 1 1.934 2.204" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedditSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M19.96 1.75c-.873 0-1.636.479-2.038 1.188l-3.6-.739a.75.75 0 0 0-.841.442c-.411.966-1.027 2.431-1.542 3.685c-.226.552-.434 1.066-.596 1.48c-1.918.092-3.692.59-5.157 1.384a3.066 3.066 0 1 0-3.817 4.57a5.371 5.371 0 0 0-.132 1.187c0 2.04 1.151 3.842 2.921 5.114c1.77 1.273 4.191 2.044 6.842 2.044c2.65 0 5.071-.771 6.842-2.044c1.77-1.272 2.921-3.073 2.921-5.114c0-.405-.045-.802-.132-1.186a3.065 3.065 0 1 0-3.817-4.57c-1.39-.755-3.059-1.241-4.864-1.368l.377-.928c.419-1.021.907-2.185 1.295-3.103l3.018.62a2.343 2.343 0 1 0 2.32-2.662m-.842 2.342a.842.842 0 1 1 1.685 0a.842.842 0 0 1-1.685 0m1.983 8.25c-.47-.883-1.166-1.673-2.024-2.334a1.566 1.566 0 1 1 2.024 2.334M4.923 10.008A1.566 1.566 0 1 0 2.9 12.342c.47-.883 1.166-1.673 2.024-2.334m10.614 4.94a1.447 1.447 0 1 0 0-2.895a1.447 1.447 0 0 0 0 2.894M9.905 13.5a1.447 1.447 0 1 1-2.895 0a1.447 1.447 0 0 1 2.895 0M8.85 16.629a.75.75 0 0 0-.8 1.268l.345.218a6.75 6.75 0 0 0 7.202 0l.345-.217a.75.75 0 1 0-.8-1.27l-.345.218a5.25 5.25 0 0 1-5.602 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RefreshOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.545 8.163a.75.75 0 0 1-.487-1.044l1.66-3.535a.75.75 0 0 1 1.36.002l.732 1.569a.755.755 0 0 1 .08-.027a8.15 8.15 0 1 1-5.8 5.903a.75.75 0 1 1 1.456.364a6.65 6.65 0 1 0 4.907-4.862l.74 1.583a.75.75 0 0 1-.872 1.043z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RefreshSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.545 8.163a.75.75 0 0 1-.487-1.044l1.66-3.535a.75.75 0 0 1 1.36.002l.732 1.569a.755.755 0 0 1 .08-.027a8.15 8.15 0 1 1-5.8 5.903a.75.75 0 1 1 1.456.364a6.65 6.65 0 1 0 4.907-4.862l.74 1.583a.75.75 0 0 1-.872 1.043z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReplyOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M20.54 16.807a.75.75 0 0 0 .13-.86l-.731-1.441a9.75 9.75 0 0 0-8.698-5.346h-.404a61.744 61.744 0 0 0-.099-1.659l-.068-.939a1.004 1.004 0 0 0-1.542-.773a19.8 19.8 0 0 0-5.42 5.133l-.46.641a.75.75 0 0 0 0 .875l.46.64a19.8 19.8 0 0 0 5.42 5.134c.64.41 1.486-.015 1.542-.773l.068-.94c.048-.659.086-1.319.113-1.98a14.131 14.131 0 0 1 6.335 1.295l2.5 1.154a.75.75 0 0 0 .854-.16m-2.21-2.118l-.515-.237a15.63 15.63 0 0 0-7.747-1.393a.75.75 0 0 0-.692.726c-.026.87-.07 1.738-.134 2.606l-.006.087a18.3 18.3 0 0 1-4.31-4.274L4.78 12l.146-.203a18.3 18.3 0 0 1 4.31-4.275l.006.087c.057.775.098 1.55.125 2.326a.75.75 0 0 0 .75.725h1.124a8.25 8.25 0 0 1 7.088 4.028" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReplySolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.446 16.06a.5.5 0 0 1-.655.68l-2.5-1.153a14.381 14.381 0 0 0-6.681-1.309a61.43 61.43 0 0 1-.121 2.204l-.069.938a.754.754 0 0 1-1.158.581a19.55 19.55 0 0 1-5.351-5.068l-.46-.64a.5.5 0 0 1 0-.584l.46-.64A19.55 19.55 0 0 1 9.262 6a.754.754 0 0 1 1.158.58l.069.94c.046.63.082 1.26.108 1.89h.644a9.5 9.5 0 0 1 8.475 5.209z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RowsOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.957 5.192a7.012 7.012 0 0 0 0 4.116c4.685.451 9.401.451 14.086 0a7.012 7.012 0 0 0 0-4.116a73.412 73.412 0 0 0-14.086 0M4.71 3.71a74.913 74.913 0 0 1 14.582 0c.535.053.984.419 1.148.925c.55 1.693.55 3.54 0 5.232a1.347 1.347 0 0 1-1.148.925a74.944 74.944 0 0 1-14.582 0a1.347 1.347 0 0 1-1.148-.925a8.508 8.508 0 0 1 0-5.232A1.347 1.347 0 0 1 4.71 3.71m.247 10.982a7.012 7.012 0 0 0 0 4.116c4.685.451 9.401.451 14.086 0a7.012 7.012 0 0 0 0-4.116a73.408 73.408 0 0 0-14.086 0M4.71 13.21a74.918 74.918 0 0 1 14.582 0c.535.053.984.419 1.148.925c.55 1.693.55 3.54 0 5.232a1.347 1.347 0 0 1-1.148.925a74.944 74.944 0 0 1-14.582 0a1.347 1.347 0 0 1-1.148-.925a8.508 8.508 0 0 1 0-5.232a1.347 1.347 0 0 1 1.148-.925" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RowsSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.267 3.958a74.662 74.662 0 0 0-14.534 0c-.435.043-.8.34-.934.753a8.258 8.258 0 0 0 0 5.078c.133.412.5.71.934.753c4.833.472 9.7.472 14.534 0c.435-.043.8-.341.934-.753a8.259 8.259 0 0 0 0-5.078c-.133-.412-.5-.71-.934-.753m0 9.5a74.67 74.67 0 0 0-14.534 0c-.435.043-.8.34-.934.753a8.258 8.258 0 0 0 0 5.078c.133.412.5.71.934.753c4.833.472 9.7.472 14.534 0c.435-.043.8-.341.934-.753a8.259 8.259 0 0 0 0-5.078c-.133-.412-.5-.71-.934-.753"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SandWatchOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.47 8.53a.75.75 0 0 0 1.06 0l1-1A.75.75 0 0 0 13 6.25h-2a.75.75 0 0 0-.53 1.28zm0 6.94a.75.75 0 0 1 1.06 0l1 1a.75.75 0 0 1-.53 1.28h-2a.75.75 0 0 1-.53-1.28z"/><path fill="currentColor" fill-rule="evenodd" d="m13.149 12l.187-.161a21.075 21.075 0 0 0 4.844-6.057a1.428 1.428 0 0 0-1.135-2.094l-1.174-.102a44.447 44.447 0 0 0-7.742 0l-1.174.102A1.428 1.428 0 0 0 5.82 5.782a21.076 21.076 0 0 0 4.844 6.057l.187.161l-.187.161A21.076 21.076 0 0 0 5.82 18.22a1.428 1.428 0 0 0 1.135 2.093l1.174.102c2.576.226 5.166.226 7.742 0l1.174-.102a1.428 1.428 0 0 0 1.135-2.093a21.076 21.076 0 0 0-4.844-6.058zm-1.145-.994l.002-.002l.351-.302a19.575 19.575 0 0 0 4.447-5.529L15.74 5.08a42.945 42.945 0 0 0-7.48 0l-1.065.093a19.576 19.576 0 0 0 4.448 5.53l.35.302a.01.01 0 0 0 .003.001H12a.01.01 0 0 0 .004 0m.002 1.99l-.002-.002H12a.01.01 0 0 0-.004 0l-.003.002l-.35.302a19.577 19.577 0 0 0-4.448 5.53l1.064.092c2.49.218 4.993.218 7.481 0l1.064-.093a19.576 19.576 0 0 0-4.447-5.53z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SandWatchSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m10.664 11.839l.187.161l-.187.161A21.075 21.075 0 0 0 5.82 18.22a1.428 1.428 0 0 0 1.135 2.093l1.174.102c2.576.226 5.166.226 7.742 0l1.174-.102a1.428 1.428 0 0 0 1.135-2.093a21.074 21.074 0 0 0-4.844-6.058l-.187-.16l.187-.162a21.074 21.074 0 0 0 4.844-6.057a1.428 1.428 0 0 0-1.135-2.094l-1.174-.102a44.446 44.446 0 0 0-7.742 0l-1.174.102A1.428 1.428 0 0 0 5.82 5.782a21.075 21.075 0 0 0 4.844 6.057M12 9.75a.75.75 0 0 1-.53-.22l-2-2A.75.75 0 0 1 10 6.25h4a.75.75 0 0 1 .53 1.28l-2 2a.75.75 0 0 1-.53.22m0 4.5a.75.75 0 0 0-.53.22l-2 2a.75.75 0 0 0 .53 1.28h4a.75.75 0 0 0 .53-1.28l-2-2a.75.75 0 0 0-.53-.22" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SaveOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.49 6.237A2.814 2.814 0 0 1 6.283 3.75h10.742a.75.75 0 0 1 .495.186l2.76 2.424c.621.545.964 1.34.934 2.165l-.344 9.574a2.75 2.75 0 0 1-2.748 2.651H6.113a2.642 2.642 0 0 1-2.621-2.307a48.06 48.06 0 0 1-.064-11.686zm2.794-.987c-.667 0-1.228.5-1.305 1.161l-.06.52a46.56 46.56 0 0 0 .06 11.322c.074.57.56.997 1.134.997h.637V15c0-.966.784-1.75 1.75-1.75h7c.966 0 1.75.784 1.75 1.75v4.25h.873a1.25 1.25 0 0 0 1.25-1.205l.343-9.573a1.25 1.25 0 0 0-.424-.985L16.75 5.255V7.6A1.75 1.75 0 0 1 15 9.35H9A1.75 1.75 0 0 1 7.25 7.6V5.25zm2.466 0V7.6c0 .138.112.25.25.25h6a.25.25 0 0 0 .25-.25V5.25zm7 14h-7.5V15a.25.25 0 0 1 .25-.25h7a.25.25 0 0 1 .25.25z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SaveSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.25 7.6c0 .966.784 1.75 1.75 1.75h6a1.75 1.75 0 0 0 1.75-1.75V4.276c0-.152.124-.276.276-.276a.5.5 0 0 1 .38.174l2.963 3.455a2.5 2.5 0 0 1 .6 1.725l-.342 8.744A2.5 2.5 0 0 1 18.13 20.5h-.379a.5.5 0 0 1-.5-.5v-5a1.75 1.75 0 0 0-1.75-1.75h-7A1.75 1.75 0 0 0 6.75 15v5a.5.5 0 0 1-.5.5h-.137a2.392 2.392 0 0 1-2.373-2.089a47.81 47.81 0 0 1-.063-11.625l.06-.52A2.564 2.564 0 0 1 6.284 4h.466a.5.5 0 0 1 .5.5z"/><path fill="currentColor" d="M8.25 20a.5.5 0 0 0 .5.5h6.5a.5.5 0 0 0 .5-.5v-5a.25.25 0 0 0-.25-.25h-7a.25.25 0 0 0-.25.25zm7-15.5a.5.5 0 0 0-.5-.5h-5.5a.5.5 0 0 0-.5.5v3.1c0 .138.112.25.25.25h6a.25.25 0 0 0 .25-.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.385 15.446a6.75 6.75 0 1 1 1.06-1.06l5.156 5.155a.75.75 0 1 1-1.06 1.06zm-7.926-1.562a5.25 5.25 0 1 1 7.43-.005l-.005.005l-.005.004a5.25 5.25 0 0 1-7.42-.004" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.385 15.446a6.751 6.751 0 1 1 1.06-1.06l5.156 5.155a.75.75 0 1 1-1.06 1.06zM6.46 13.884a5.25 5.25 0 1 1 7.43-.005l-.005.005l-.005.004a5.25 5.25 0 0 1-7.42-.004" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SendOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M21.75 12a.75.75 0 0 1-.386.656l-2.282 1.268a75.748 75.748 0 0 1-14.193 6.084l-.665.208a.75.75 0 0 1-.974-.716v-5.75c0-.391.3-.716.69-.748l.228-.018A44.247 44.247 0 0 0 10.555 12a44.24 44.24 0 0 0-6.478-.992l-.135-.01a.75.75 0 0 1-.692-.748V4.5a.75.75 0 0 1 .974-.716l.665.208a75.75 75.75 0 0 1 14.193 6.084l2.282 1.268a.75.75 0 0 1 .386.656m-2.294 0l-1.103-.612A74.247 74.247 0 0 0 4.75 5.52v4.04c2.93.264 5.828.81 8.654 1.631l.305.089a.75.75 0 0 1-.001 1.44l-.39.113A45.749 45.749 0 0 1 4.75 14.44v4.04a74.25 74.25 0 0 0 13.603-5.868z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SendSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.243 12.437a.5.5 0 0 0 0-.874l-2.282-1.268A75.497 75.497 0 0 0 4.813 4.231l-.665-.208A.5.5 0 0 0 3.5 4.5v5.75a.5.5 0 0 0 .474.5l1.01.053a44.41 44.41 0 0 1 7.314.998l.238.053c.053.011.076.033.089.05a.163.163 0 0 1 .029.096c0 .04-.013.074-.029.096c-.013.017-.036.039-.089.05l-.238.053a44.509 44.509 0 0 1-7.315.999l-1.01.053a.5.5 0 0 0-.473.499v5.75a.5.5 0 0 0 .65.477l.664-.208a75.499 75.499 0 0 0 14.147-6.064z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServerOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 7.5a1 1 0 1 1 2 0a1 1 0 0 1-2 0m3 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0m-2 6a1 1 0 1 0 0 2a1 1 0 0 0 0-2m3 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2"/><path fill="currentColor" fill-rule="evenodd" d="M5.947 3.25A2.75 2.75 0 0 0 3.197 6v3c0 .788.332 1.499.863 2a2.742 2.742 0 0 0-.863 2v3a2.75 2.75 0 0 0 2.75 2.75h5.303v1.5H5a.75.75 0 0 0 0 1.5h14a.75.75 0 0 0 0-1.5h-6.25v-1.5h5.197a2.75 2.75 0 0 0 2.75-2.75v-3c0-.788-.331-1.499-.862-2a2.743 2.743 0 0 0 .862-2V6a2.75 2.75 0 0 0-2.75-2.75zm0 7h12c.69 0 1.25-.56 1.25-1.25V6c0-.69-.56-1.25-1.25-1.25h-12c-.69 0-1.25.56-1.25 1.25v3c0 .69.56 1.25 1.25 1.25m0 1.5c-.69 0-1.25.56-1.25 1.25v3c0 .69.56 1.25 1.25 1.25h12c.69 0 1.25-.56 1.25-1.25v-3c0-.69-.56-1.25-1.25-1.25z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServerSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.947 5a1.5 1.5 0 0 1 1.5-1.5h13a1.5 1.5 0 0 1 1.5 1.5v4a1.5 1.5 0 0 1-1.5 1.5h-13a1.5 1.5 0 0 1-1.5-1.5zM6 7a1 1 0 1 1 2 0a1 1 0 0 1-2 0m4-1a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-6.053 7a1.5 1.5 0 0 1 1.5-1.5h13a1.5 1.5 0 0 1 1.5 1.5v4a1.5 1.5 0 0 1-1.5 1.5H12.75v1.75H19a.75.75 0 0 1 0 1.5H5a.75.75 0 0 1 0-1.5h6.25V18.5H5.447a1.5 1.5 0 0 1-1.5-1.5zM6 15a1 1 0 1 1 2 0a1 1 0 0 1-2 0m3 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SettingsAdjustOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.354 8.75H4a.75.75 0 0 1 0-1.5h9.354a2.751 2.751 0 0 1 5.293 0H20a.75.75 0 0 1 0 1.5h-1.354a2.751 2.751 0 0 1-5.292 0M14.75 8a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0m-4.103 8.75H20a.75.75 0 0 0 0-1.5h-9.353a2.751 2.751 0 0 0-5.293 0H4a.75.75 0 0 0 0 1.5h1.354a2.751 2.751 0 0 0 5.292 0M6.75 16a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SettingsAdjustSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.878 8.75H4a.75.75 0 0 1 0-1.5h9.878a2.251 2.251 0 0 1 4.244 0H20a.75.75 0 0 1 0 1.5h-1.878a2.251 2.251 0 0 1-4.244 0m6.122 8a.75.75 0 0 0 0-1.5h-9.878a2.251 2.251 0 0 0-4.244 0H4a.75.75 0 0 0 0 1.5h1.878a2.25 2.25 0 0 0 4.244 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SettingsAltOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M12 3.08L9.787 5.368a1.25 1.25 0 0 1-.899.381H5.75v3.138c0 .34-.138.664-.381.9L3.079 12l2.29 2.213c.243.235.381.56.381.899v3.138h3.138c.34 0 .664.138.9.381L12 20.921l2.213-2.29a1.25 1.25 0 0 1 .899-.381h3.138v-3.138c0-.34.138-.664.381-.9L20.921 12l-2.29-2.213a1.25 1.25 0 0 1-.381-.899V5.75h-3.138a1.25 1.25 0 0 1-.9-.381zm-.899-1.23a1.25 1.25 0 0 1 1.798 0l2.319 2.4H18.5c.69 0 1.25.56 1.25 1.25v3.282l2.4 2.32a1.25 1.25 0 0 1 0 1.797l-2.4 2.319V18.5c0 .69-.56 1.25-1.25 1.25h-3.282l-2.32 2.4a1.25 1.25 0 0 1-1.797 0l-2.319-2.4H5.5c-.69 0-1.25-.56-1.25-1.25v-3.282l-2.4-2.32a1.25 1.25 0 0 1 0-1.797l2.4-2.319V5.5c0-.69.56-1.25 1.25-1.25h3.282z"/><path d="M7.25 12a4.75 4.75 0 1 1 9.5 0a4.75 4.75 0 0 1-9.5 0M12 8.75a3.25 3.25 0 1 0 0 6.5a3.25 3.25 0 0 0 0-6.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SettingsAltSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.75 12a3.25 3.25 0 1 1 6.5 0a3.25 3.25 0 0 1-6.5 0"/><path fill="currentColor" fill-rule="evenodd" d="M11.46 1.838a.75.75 0 0 1 1.08 0L15.111 4.5h3.638a.75.75 0 0 1 .75.75v3.638l2.662 2.573a.75.75 0 0 1 0 1.078L19.5 15.112v3.638a.75.75 0 0 1-.75.75h-3.638l-2.573 2.661a.75.75 0 0 1-1.078 0L8.888 19.5H5.25a.75.75 0 0 1-.75-.75v-3.64l-2.661-2.57a.75.75 0 0 1 0-1.078L4.5 8.888V5.25a.75.75 0 0 1 .75-.75h3.638zM12 7.25a4.75 4.75 0 1 0 0 9.5a4.75 4.75 0 0 0 0-9.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SettingsOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M11.199 2.587a1.65 1.65 0 0 1 1.602 0l7.2 4c.524.291.849.843.849 1.443v7.94c0 .6-.325 1.152-.849 1.443l-7.2 4a1.65 1.65 0 0 1-1.602 0l-7.2-4a1.65 1.65 0 0 1-.849-1.443V8.03c0-.6.325-1.152.849-1.443zm.874 1.312a.15.15 0 0 0-.146 0l-7.2 4a.15.15 0 0 0-.077.13v7.942c0 .054.03.104.077.13l7.2 4a.15.15 0 0 0 .146 0l7.2-4a.15.15 0 0 0 .077-.13V8.03a.15.15 0 0 0-.077-.131z"/><path d="M7.25 12a4.75 4.75 0 1 1 9.5 0a4.75 4.75 0 0 1-9.5 0M12 8.75a3.25 3.25 0 1 0 0 6.5a3.25 3.25 0 0 0 0-6.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SettingsSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 8.75a3.25 3.25 0 1 0 0 6.5a3.25 3.25 0 0 0 0-6.5"/><path fill="currentColor" fill-rule="evenodd" d="M12.68 2.806a1.4 1.4 0 0 0-1.36 0l-7.2 4A1.4 1.4 0 0 0 3.4 8.03v7.94c0 .509.276.977.72 1.224l7.2 4a1.4 1.4 0 0 0 1.36 0l7.2-4a1.4 1.4 0 0 0 .72-1.223V8.03a1.4 1.4 0 0 0-.72-1.224zM7.25 12a4.75 4.75 0 1 1 9.5 0a4.75 4.75 0 0 1-9.5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareBoxOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.25 15.05a.75.75 0 0 0 1.403.387l.903-1.6a5.25 5.25 0 0 1 5.354-2.611l.149.022c.027.678.065 1.356.114 2.032l.068.931c.055.757.899 1.18 1.538.771a19.631 19.631 0 0 0 5.374-5.089l.456-.635a.75.75 0 0 0 0-.875l-.456-.635a19.631 19.631 0 0 0-5.373-5.09a1.002 1.002 0 0 0-1.539.772l-.068.93c-.04.547-.072 1.093-.097 1.64h-1.019a6.75 6.75 0 0 0-6.748 6.587zm7.883-5.308a6.75 6.75 0 0 0-6.3 2.489a5.25 5.25 0 0 1 5.224-4.732h1.74a.75.75 0 0 0 .749-.724c.026-.77.067-1.538.123-2.305l.006-.075a18.13 18.13 0 0 1 4.26 4.228l.142.198l-.142.197a18.132 18.132 0 0 1-4.26 4.228l-.006-.075a59.785 59.785 0 0 1-.132-2.594a.75.75 0 0 0-.638-.72z" clip-rule="evenodd"/><path fill="currentColor" d="M19.641 17.411a44.03 44.03 0 0 0 .25-3.032a.407.407 0 0 1 .137-.285c.364-.32.716-.654 1.056-1c.128-.13.351-.038.348.144a45.91 45.91 0 0 1-.3 4.348c-.237 2.022-1.862 3.607-3.874 3.832a47.77 47.77 0 0 1-10.516 0c-2.012-.225-3.637-1.81-3.873-3.832a45.92 45.92 0 0 1 0-10.67c.236-2.022 1.86-3.607 3.873-3.832a47.672 47.672 0 0 1 6.958-.26c.196.006.33.198.307.393a2.33 2.33 0 0 0-.013.122l-.051.704a.301.301 0 0 1-.312.279a46.171 46.171 0 0 0-6.723.253a2.89 2.89 0 0 0-2.55 2.516a44.421 44.421 0 0 0 0 10.32a2.89 2.89 0 0 0 2.55 2.516c3.356.375 6.828.375 10.184 0a2.89 2.89 0 0 0 2.55-2.516"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareBoxSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.028 14.094a.407.407 0 0 0-.137.285a44.464 44.464 0 0 1-.25 3.032c-.154 1.318-1.22 2.367-2.55 2.516a46.216 46.216 0 0 1-10.183 0a2.892 2.892 0 0 1-2.55-2.516a44.421 44.421 0 0 1 0-10.32a2.893 2.893 0 0 1 2.55-2.517a46.21 46.21 0 0 1 6.725-.252c.163.005.3-.118.312-.28l.05-.703a1.28 1.28 0 0 1 .012-.123c.024-.194-.11-.386-.307-.393a47.672 47.672 0 0 0-6.958.26c-2.012.226-3.637 1.81-3.873 3.833a45.921 45.921 0 0 0 0 10.67c.236 2.022 1.86 3.607 3.873 3.832a47.77 47.77 0 0 0 10.516 0c2.012-.225 3.637-1.81 3.873-3.833a45.91 45.91 0 0 0 .3-4.347c.004-.183-.22-.273-.347-.144c-.34.346-.692.679-1.056 1"/><path fill="currentColor" d="M7.867 15.55a.5.5 0 0 1-.367-.494l.06-2.463a6.5 6.5 0 0 1 6.498-6.344h1.257a61.47 61.47 0 0 1 .107-1.87l.068-.93a.752.752 0 0 1 1.155-.58a19.38 19.38 0 0 1 5.305 5.025l.456.635a.5.5 0 0 1 0 .583l-.456.635a19.38 19.38 0 0 1-5.305 5.024a.754.754 0 0 1-1.083-.303a.771.771 0 0 1-.072-.275l-.068-.931a61.112 61.112 0 0 1-.121-2.23l-.354-.054a5.5 5.5 0 0 0-5.61 2.736l-.902 1.6a.5.5 0 0 1-.568.236"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.25 5.5a3.25 3.25 0 1 1 .833 2.173l-2.717 1.482l-3.04 1.737a3.254 3.254 0 0 1 .31 2.464l5.447 2.971a3.25 3.25 0 1 1-.719 1.316l-5.447-2.97a3.25 3.25 0 1 1-.652-4.902l3.37-1.926l2.729-1.489a3.254 3.254 0 0 1-.114-.856m3.25-1.75a1.75 1.75 0 1 0 0 3.5a1.75 1.75 0 0 0 0-3.5m-11 7a1.75 1.75 0 1 0 0 3.5a1.75 1.75 0 0 0 0-3.5m9.25 7.75a1.75 1.75 0 1 1 3.5 0a1.75 1.75 0 0 1-3.5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.5 2.5a3 3 0 0 0-2.902 3.765a.75.75 0 0 0-.207.077l-2.757 1.503L8.128 9.85a.755.755 0 0 0-.1.068a3 3 0 1 0 .682 4.612l2.926 1.627l2.954 1.611a3 3 0 1 0 .719-1.316l-2.947-1.608l-2.946-1.636a3.007 3.007 0 0 0-.308-2.19l3.258-1.862l2.743-1.497a.75.75 0 0 0 .177-.133A3 3 0 1 0 17.5 2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.645 5.264a1.25 1.25 0 0 0-1.29 0l-.516.31a10.75 10.75 0 0 1-4.537 1.494l-.325.03a.25.25 0 0 0-.227.25V8.99a9.25 9.25 0 0 0 2.82 6.65l3.256 3.148a.25.25 0 0 0 .348 0l3.255-3.147a9.25 9.25 0 0 0 2.821-6.65V7.346a.25.25 0 0 0-.227-.248l-.325-.031a10.75 10.75 0 0 1-4.537-1.493zM10.58 3.979a2.75 2.75 0 0 1 2.838 0l.516.31a9.25 9.25 0 0 0 3.904 1.286l.325.03a1.75 1.75 0 0 1 1.586 1.742v1.644a10.75 10.75 0 0 1-3.278 7.73l-3.256 3.146a1.75 1.75 0 0 1-2.432 0L7.528 16.72A10.75 10.75 0 0 1 4.25 8.991V7.347a1.75 1.75 0 0 1 1.586-1.742l.325-.03a9.25 9.25 0 0 0 3.904-1.285z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.16 4.407a2.25 2.25 0 0 0-2.32 0l-.517.311a9.75 9.75 0 0 1-4.115 1.354l-.325.031A1.25 1.25 0 0 0 4.75 7.347v1.644a10.25 10.25 0 0 0 3.126 7.37l3.255 3.147a1.25 1.25 0 0 0 1.738 0l3.255-3.147a10.25 10.25 0 0 0 3.126-7.37V7.347a1.25 1.25 0 0 0-1.133-1.244l-.325-.03a9.75 9.75 0 0 1-4.115-1.355z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBagOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.25 7.25v-.12a4.75 4.75 0 1 1 9.5 0v.12h1.501c.571 0 1.056.419 1.14.984l.218 1.493c.43 2.938.43 5.924 0 8.862a3.135 3.135 0 0 1-2.779 2.664l-.629.065a40.68 40.68 0 0 1-8.402 0l-.629-.065a3.135 3.135 0 0 1-2.779-2.664a30.565 30.565 0 0 1 0-8.862l.219-1.493a1.151 1.151 0 0 1 1.139-.984zm3.94-3.267a3.25 3.25 0 0 1 4.06 3.147v.12h-6.5v-.12a3.25 3.25 0 0 1 2.44-3.147M7.25 8.75V11a.75.75 0 0 0 1.5 0V8.75h6.5V11a.75.75 0 0 0 1.5 0V8.75h1.2l.175 1.194a29.098 29.098 0 0 1 0 8.428a1.635 1.635 0 0 1-1.45 1.39l-.629.064c-2.69.28-5.402.28-8.092 0l-.63-.065a1.635 1.635 0 0 1-1.449-1.39a29.067 29.067 0 0 1 0-8.427L6.05 8.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBagSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.25 7.13v.37H5.749a.901.901 0 0 0-.892.77L4.64 9.763a30.316 30.316 0 0 0 0 8.79a2.885 2.885 0 0 0 2.557 2.451l.629.066c2.776.288 5.574.288 8.35 0l.63-.066a2.885 2.885 0 0 0 2.556-2.451a30.318 30.318 0 0 0 0-8.79l-.218-1.493a.901.901 0 0 0-.892-.77H16.75v-.37a4.75 4.75 0 1 0-9.5 0m5.56-3.147A3.25 3.25 0 0 0 8.75 7.13v.37h6.5v-.37a3.25 3.25 0 0 0-2.44-3.147M8.75 9a.75.75 0 0 0-1.5 0v2a.75.75 0 0 0 1.5 0zm8 0a.75.75 0 0 0-1.5 0v2a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBasketOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 13.25a.75.75 0 0 1 .75.75v2a.75.75 0 0 1-1.5 0v-2a.75.75 0 0 1 .75-.75m4.75.75a.75.75 0 0 0-1.5 0v2a.75.75 0 0 0 1.5 0z"/><path fill="currentColor" fill-rule="evenodd" d="M9.65 3.375a.75.75 0 0 0-1.3-.75l-2 3.464a.752.752 0 0 0-.069.161H6a2.75 2.75 0 0 0-1.739 4.88l.667 4.585l.447 2.093a3.049 3.049 0 0 0 2.561 2.384c2.697.375 5.432.375 8.128 0a3.049 3.049 0 0 0 2.561-2.384l.447-2.093l.667-4.584A2.75 2.75 0 0 0 18 6.25h-.281a.754.754 0 0 0-.07-.162l-2-3.464a.75.75 0 1 0-1.298.75l1.66 2.875H7.99zm8.484 8.372a2.819 2.819 0 0 1-.134.003H6c-.045 0-.09-.001-.133-.003l.538 3.703l.437 2.045a1.549 1.549 0 0 0 1.301 1.211c2.559.356 5.155.356 7.714 0a1.549 1.549 0 0 0 1.301-1.21l.437-2.046zM4.75 9c0-.69.56-1.25 1.25-1.25h12a1.25 1.25 0 1 1 0 2.5H6c-.69 0-1.25-.56-1.25-1.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBasketSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.65 4.375a.75.75 0 0 0-1.3-.75L6.69 6.5H5.5a2 2 0 0 0-1.058 3.697a.362.362 0 0 0 .193.053h14.73a.362.362 0 0 0 .193-.053A2 2 0 0 0 18.5 6.5h-1.19l-1.66-2.875a.75.75 0 1 0-1.3.75L15.578 6.5H8.423z"/><path fill="currentColor" fill-rule="evenodd" d="M19.335 11.985a.2.2 0 0 0-.197-.235H4.862a.2.2 0 0 0-.197.235l1 5.698a2.773 2.773 0 0 0 2.35 2.267c2.644.368 5.326.368 7.97 0a2.773 2.773 0 0 0 2.35-2.267zM10.75 14a.75.75 0 0 0-1.5 0v2a.75.75 0 0 0 1.5 0zm3.25-.75a.75.75 0 0 1 .75.75v2a.75.75 0 0 1-1.5 0v-2a.75.75 0 0 1 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCartOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M19.148 5.25H5.335l-1.18-2.115A.75.75 0 0 0 3.5 2.75H2a.75.75 0 0 0 0 1.5h1.06l1.164 2.088L6.91 12.28l.003.006l.237.523l-2.697 2.877a.75.75 0 0 0 .462 1.258l2.458.281a40.68 40.68 0 0 0 9.254 0l2.458-.28a.75.75 0 0 0-.17-1.491l-2.458.28a39.256 39.256 0 0 1-8.914 0l-.975-.11l1.98-2.112a.768.768 0 0 0 .053-.064l.752.098c1.055.138 2.122.165 3.182.08a9.29 9.29 0 0 0 6.366-3.268l.579-.685a.734.734 0 0 0 .053-.072l1.078-1.642c.764-1.164-.071-2.71-1.463-2.71M8.656 11.944a.484.484 0 0 1-.377-.278l-.002-.003l-2.22-4.913h13.09a.25.25 0 0 1 .21.387l-1.053 1.604l-.549.65a7.79 7.79 0 0 1-5.338 2.741c-.957.076-1.919.052-2.87-.072z" clip-rule="evenodd"/><path fill="currentColor" d="M6.5 18.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3M16 20a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCartSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.148 5.25H5.335l-1.18-2.115A.75.75 0 0 0 3.5 2.75H2a.75.75 0 0 0 0 1.5h1.06l1.164 2.088L6.91 12.28l.003.006l.237.523l-2.697 2.877a.75.75 0 0 0 .462 1.258l2.458.281a40.68 40.68 0 0 0 9.254 0l2.458-.28a.75.75 0 0 0-.17-1.491l-2.458.28a39.256 39.256 0 0 1-8.914 0l-.975-.11l1.98-2.112a.768.768 0 0 0 .053-.064l.752.098c1.055.138 2.122.165 3.182.08a9.29 9.29 0 0 0 6.366-3.268l.579-.685a.734.734 0 0 0 .053-.072l1.078-1.642c.764-1.164-.071-2.71-1.463-2.71M6.5 18.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3M16 20a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShuffleOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.47 5.47a.75.75 0 0 0 0 1.06l.72.72h-3.813a1.75 1.75 0 0 0-1.575.987l-.21.434a.4.4 0 0 0 0 .35l.568 1.174a.2.2 0 0 0 .36 0l.632-1.304a.25.25 0 0 1 .225-.141h3.812l-.72.72a.75.75 0 1 0 1.061 1.06l2-2a.75.75 0 0 0 0-1.06l-2-2a.75.75 0 0 0-1.06 0m-6.436 9.859a.4.4 0 0 0 0-.35l-.57-1.174a.2.2 0 0 0-.36 0l-.63 1.304a.25.25 0 0 1-.226.141H5a.75.75 0 0 0 0 1.5h3.248a1.75 1.75 0 0 0 1.575-.987z"/><path fill="currentColor" d="M16.47 18.53a.75.75 0 0 1 0-1.06l.72-.72h-3.813a1.75 1.75 0 0 1-1.575-.987L8.473 8.89a.25.25 0 0 0-.225-.141H5a.75.75 0 0 1 0-1.5h3.248c.671 0 1.283.383 1.575.987l3.329 6.872a.25.25 0 0 0 .225.141h3.812l-.72-.72a.75.75 0 1 1 1.061-1.06l2 2a.75.75 0 0 1 0 1.06l-2 2a.75.75 0 0 1-1.06 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShuffleSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.47 5.47a.75.75 0 0 0 0 1.06l.72.72h-3.813a1.75 1.75 0 0 0-1.575.987l-.21.434a.4.4 0 0 0 0 .35l.568 1.174a.2.2 0 0 0 .36 0l.632-1.304a.25.25 0 0 1 .225-.141h3.812l-.72.72a.75.75 0 1 0 1.061 1.06l2-2a.75.75 0 0 0 0-1.06l-2-2a.75.75 0 0 0-1.06 0m-6.436 9.859a.4.4 0 0 0 0-.35l-.57-1.174a.2.2 0 0 0-.36 0l-.63 1.304a.25.25 0 0 1-.226.141H5a.75.75 0 0 0 0 1.5h3.248a1.75 1.75 0 0 0 1.575-.987z"/><path fill="currentColor" d="M16.47 18.53a.75.75 0 0 1 0-1.06l.72-.72h-3.813a1.75 1.75 0 0 1-1.575-.987L8.473 8.89a.25.25 0 0 0-.225-.141H5a.75.75 0 0 1 0-1.5h3.248c.671 0 1.283.383 1.575.987l3.329 6.872a.25.25 0 0 0 .225.141h3.812l-.72-.72a.75.75 0 1 1 1.061-1.06l2 2a.75.75 0 0 1 0 1.06l-2 2a.75.75 0 0 1-1.06 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SketchOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.998 2.25h.009a.65.65 0 0 1 .077.005l5.904.624a.748.748 0 0 1 .529.299l4.575 6.142a.749.749 0 0 1-.03.948L12.574 22.48a.75.75 0 0 1-1.154 0L.936 10.27a.753.753 0 0 1-.034-.947L5.477 3.18a.744.744 0 0 1 .532-.302l5.903-.624a.751.751 0 0 1 .077-.005h.001zm9.004 6.775l-2.548-3.421l-.374 3.421zm-7.335-5.094L16.7 7.847l.388-3.555zm-6.759.361l3.42-.361l-3.032 3.916zm5.09-.067l3.717 4.8H8.28zm-8.865 6.3h3.125L9.379 17.8zm2.41-4.921l-2.55 3.421h2.923zm15.32 4.921L14.617 17.8l3.121-7.274zm-12.973 0h8.216l-4.108 9.573z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SketchSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.877 3.153a.15.15 0 0 0-.134-.241l-3.503.37a.15.15 0 0 0-.133.166l.397 3.64a.15.15 0 0 0 .268.075zm7.347 4.01a.15.15 0 0 0 .268-.075l.397-3.64a.15.15 0 0 0-.133-.166l-3.503-.37a.15.15 0 0 0-.134.24zM7.836 8.635a.15.15 0 0 1-.119-.242l4.162-5.376a.15.15 0 0 1 .238 0l4.162 5.376a.15.15 0 0 1-.119.242zm1.157 9.625c.068.159-.14.288-.251.158l-6.9-8.035a.15.15 0 0 1 .113-.248h3.452a.15.15 0 0 1 .138.09zm9.596-8.125a.15.15 0 0 0-.138.09l-3.449 8.036c-.067.158.14.287.252.157l6.9-8.035a.15.15 0 0 0-.113-.248zm.261-1.667a.15.15 0 0 0 .149.167h2.78a.15.15 0 0 0 .12-.24l-2.425-3.254a.15.15 0 0 0-.269.073zM4.791 5.214a.15.15 0 0 0-.27-.073L2.098 8.395a.15.15 0 0 0 .12.24h2.78a.15.15 0 0 0 .15-.167zm2.437 5.13a.15.15 0 0 1 .138-.21h9.264c.107 0 .18.11.137.21l-4.631 10.794a.15.15 0 0 1-.276 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkipNextOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.607 10.748c.82.634.82 1.87 0 2.505a25.758 25.758 0 0 1-7.143 3.9l-.466.166c-1.023.364-2.1-.329-2.238-1.381c-.34-2.59-.34-5.286 0-7.876c.138-1.052 1.215-1.745 2.238-1.381l.466.165a25.76 25.76 0 0 1 7.143 3.902m-.918 1.318a.084.084 0 0 0 0-.132A24.257 24.257 0 0 0 6.962 8.26l-.466-.166a.194.194 0 0 0-.249.163a29.063 29.063 0 0 0 0 7.486c.017.13.15.198.25.163l.465-.166c2.423-.86 4.694-2.1 6.727-3.674M18 6.25a.75.75 0 0 1 .75.75v10a.75.75 0 0 1-1.5 0V7a.75.75 0 0 1 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkipNextSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.75 7a.75.75 0 0 0-1.5 0v10a.75.75 0 0 0 1.5 0zm-4.296 3.945c.69.534.69 1.576 0 2.11a25.51 25.51 0 0 1-7.073 3.863l-.466.166c-.87.308-1.79-.28-1.907-1.178a30.314 30.314 0 0 1 0-7.812c.118-.898 1.037-1.486 1.907-1.177l.466.165a25.511 25.511 0 0 1 7.073 3.863"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkipPrevOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.75 7a.75.75 0 0 0-1.5 0v10a.75.75 0 0 0 1.5 0z"/><path fill="currentColor" fill-rule="evenodd" d="M9.393 13.253a1.584 1.584 0 0 1 0-2.505a25.76 25.76 0 0 1 7.143-3.902l.466-.165c1.023-.364 2.1.329 2.238 1.381c.34 2.59.34 5.286 0 7.876c-.138 1.052-1.215 1.745-2.238 1.381l-.466-.165a25.758 25.758 0 0 1-7.143-3.902m.918-1.32a.084.084 0 0 0 0 .133a24.257 24.257 0 0 0 6.727 3.674l.466.166c.1.035.232-.033.249-.163c.322-2.46.322-5.025 0-7.486a.194.194 0 0 0-.25-.163l-.465.166c-2.423.86-4.694 2.1-6.727 3.674" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkipPrevSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.75 7a.75.75 0 0 0-1.5 0v10a.75.75 0 0 0 1.5 0zm3.102 5.66a.834.834 0 0 1 0-1.32a25.009 25.009 0 0 1 6.935-3.787l.466-.165a.944.944 0 0 1 1.243.772a29.813 29.813 0 0 1 0 7.68a.944.944 0 0 1-1.243.772l-.466-.165a25.01 25.01 0 0 1-6.935-3.788"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkypeOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.204 11.273c.374.11.73.274 1.057.485c.335.22.615.516.816.862c.209.384.315.815.307 1.252a2.574 2.574 0 0 1-.478 1.576a3.035 3.035 0 0 1-1.362 1.042a4.67 4.67 0 0 1-1.845.331a4.6 4.6 0 0 1-2.146-.464a3.088 3.088 0 0 1-1.098-.971A2.406 2.406 0 0 1 8 14.044a.818.818 0 1 1 1.637 0c.016.15.073.292.165.412c.13.187.303.34.504.447c.435.204.913.3 1.393.282a3.1 3.1 0 0 0 1.217-.207c.26-.093.488-.26.656-.478a.942.942 0 0 0 .176-.6a.983.983 0 0 0-.097-.478a.923.923 0 0 0-.297-.301a2.474 2.474 0 0 0-.624-.282c-.18-.054-.375-.107-.578-.16l-.015-.002a.467.467 0 0 0-.044-.011a11.627 11.627 0 0 0-.407-.098c-.361-.078-.612-.14-.769-.188l-.01-.006a.038.038 0 0 0-.008-.004c-.363-.088-.72-.2-1.07-.334a2.865 2.865 0 0 1-1.176-.877a2.323 2.323 0 0 1-.475-1.467c0-.53.174-1.047.495-1.47c.334-.426.781-.75 1.29-.934A4.88 4.88 0 0 1 11.718 7c.47-.005.939.06 1.39.193c.39.114.756.3 1.079.548c.28.212.512.478.685.784a.817.817 0 1 1-1.441.775a1 1 0 0 0-.266-.28a1.704 1.704 0 0 0-.543-.264a3.099 3.099 0 0 0-.905-.12a3.283 3.283 0 0 0-1.163.177c-.223.076-.42.21-.573.39a.791.791 0 0 0-.168.488a.69.69 0 0 0 .131.463c.134.163.307.29.504.366a7.63 7.63 0 0 0 1.175.346c.336.074.607.14.837.203c.261.065.515.133.744.203"/><path fill="currentColor" fill-rule="evenodd" d="M8 2.75a5.25 5.25 0 0 0-4.696 7.6A8.85 8.85 0 0 0 13.65 20.696a5.25 5.25 0 0 0 7.045-7.045A8.85 8.85 0 0 0 10.35 3.304A5.231 5.231 0 0 0 8 2.749M4.25 8a3.75 3.75 0 0 1 5.611-3.256a.75.75 0 0 0 .536.081a7.35 7.35 0 0 1 8.778 8.778a.75.75 0 0 0 .081.536a3.75 3.75 0 0 1-5.118 5.118a.75.75 0 0 0-.535-.082a7.35 7.35 0 0 1-8.778-8.778a.75.75 0 0 0-.081-.536A3.73 3.73 0 0 1 4.25 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkypeSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.977 3.742a8.5 8.5 0 0 1 10.281 10.281a4.5 4.5 0 0 1-6.234 6.234A8.5 8.5 0 0 1 3.742 9.977a4.5 4.5 0 0 1 6.234-6.234m3.527 7.53c.374.112.73.275 1.057.486c.335.22.615.516.816.862c.209.384.315.815.307 1.252a2.574 2.574 0 0 1-.478 1.576a3.035 3.035 0 0 1-1.362 1.042a4.67 4.67 0 0 1-1.845.331a4.6 4.6 0 0 1-2.146-.464a3.088 3.088 0 0 1-1.098-.971a2.406 2.406 0 0 1-.455-1.342a.818.818 0 1 1 1.637 0c.016.15.074.292.165.412c.13.187.303.34.504.447c.435.204.913.3 1.393.282a3.1 3.1 0 0 0 1.217-.207c.26-.093.488-.26.656-.478a.942.942 0 0 0 .175-.6a.983.983 0 0 0-.096-.478a.923.923 0 0 0-.297-.301a2.472 2.472 0 0 0-.624-.282c-.18-.054-.375-.107-.578-.16l-.015-.002a.467.467 0 0 0-.044-.011a11.627 11.627 0 0 0-.407-.098c-.361-.078-.612-.14-.769-.188a.027.027 0 0 1-.01-.006a.038.038 0 0 0-.008-.004c-.363-.088-.72-.2-1.07-.334a2.865 2.865 0 0 1-1.176-.877a2.323 2.323 0 0 1-.475-1.467c0-.53.174-1.047.495-1.47c.334-.426.781-.75 1.29-.934A4.88 4.88 0 0 1 12.018 7c.47-.005.939.06 1.39.193c.39.114.756.3 1.079.548c.28.212.512.478.685.784a.817.817 0 1 1-1.441.775a1 1 0 0 0-.266-.28a1.704 1.704 0 0 0-.543-.264a3.099 3.099 0 0 0-.905-.12a3.283 3.283 0 0 0-1.163.177c-.223.076-.42.21-.573.39a.791.791 0 0 0-.168.488a.69.69 0 0 0 .131.463c.134.163.308.29.504.366c.382.146.775.262 1.175.346c.336.074.607.14.837.203c.261.065.515.133.744.203" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlackOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 1.25a2.248 2.248 0 0 0-2.25 2.25v5a2.248 2.248 0 0 0 2.25 2.25a2.248 2.248 0 0 0 2.25-2.25v-5a2.248 2.248 0 0 0-2.25-2.25m-.75 2.25c0-.416.334-.75.75-.75s.75.334.75.75v5c0 .416-.334.75-.75.75a.748.748 0 0 1-.75-.75zm6.75 2.75a2.248 2.248 0 0 0-2.25 2.25V10c0 .414.336.75.75.75h1.5a2.248 2.248 0 0 0 2.25-2.25a2.248 2.248 0 0 0-2.25-2.25m-.75 2.25c0-.416.334-.75.75-.75s.75.334.75.75s-.334.75-.75.75h-.75zm-12.5 7a2.248 2.248 0 0 1 2.25-2.25a2.248 2.248 0 0 1 2.25 2.25v5a2.248 2.248 0 0 1-2.25 2.25a2.248 2.248 0 0 1-2.25-2.25zm2.25-.75a.748.748 0 0 0-.75.75v5c0 .416.334.75.75.75s.75-.334.75-.75v-5a.748.748 0 0 0-.75-.75m-6-1.5a2.248 2.248 0 0 0-2.25 2.25a2.248 2.248 0 0 0 2.25 2.25a2.248 2.248 0 0 0 2.25-2.25V14a.75.75 0 0 0-.75-.75zm-.75 2.25c0-.416.334-.75.75-.75h.75v.75c0 .416-.334.75-.75.75a.748.748 0 0 1-.75-.75m10.5-1a2.248 2.248 0 0 1 2.25-2.25h5a2.248 2.248 0 0 1 2.25 2.25a2.248 2.248 0 0 1-2.25 2.25h-5a2.248 2.248 0 0 1-2.25-2.25m2.25-.75a.748.748 0 0 0-.75.75c0 .416.334.75.75.75h5c.416 0 .75-.334.75-.75a.748.748 0 0 0-.75-.75zm-1.5 4.5a.75.75 0 0 0-.75.75v1.5a2.248 2.248 0 0 0 2.25 2.25a2.248 2.248 0 0 0 2.25-2.25a2.248 2.248 0 0 0-2.25-2.25zm.75 2.25v-.75h.75c.416 0 .75.334.75.75s-.334.75-.75.75a.748.748 0 0 1-.75-.75m-13.5-11A2.248 2.248 0 0 1 3.5 7.25h5a2.248 2.248 0 0 1 2.25 2.25a2.248 2.248 0 0 1-2.25 2.25h-5A2.248 2.248 0 0 1 1.25 9.5m2.25-.75a.748.748 0 0 0-.75.75c0 .416.334.75.75.75h5c.416 0 .75-.334.75-.75a.748.748 0 0 0-.75-.75zm5-7.5A2.248 2.248 0 0 0 6.25 3.5A2.248 2.248 0 0 0 8.5 5.75H10a.75.75 0 0 0 .75-.75V3.5A2.248 2.248 0 0 0 8.5 1.25M7.75 3.5c0-.416.334-.75.75-.75s.75.334.75.75v.75H8.5a.748.748 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlackSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 1.5c-1.106 0-2 .894-2 2v5c0 1.106.894 2 2 2c1.106 0 2-.894 2-2v-5c0-1.106-.894-2-2-2m6 5c-1.106 0-2 .894-2 2V10a.5.5 0 0 0 .5.5h1.5c1.106 0 2-.894 2-2c0-1.106-.894-2-2-2m-11 7c-1.106 0-2 .894-2 2v5c0 1.106.894 2 2 2c1.106 0 2-.894 2-2v-5c0-1.106-.894-2-2-2m-6 0c-1.106 0-2 .894-2 2c0 1.106.894 2 2 2c1.106 0 2-.894 2-2V14a.5.5 0 0 0-.5-.5zm12-1c-1.106 0-2 .894-2 2c0 1.106.894 2 2 2h5c1.106 0 2-.894 2-2c0-1.106-.894-2-2-2zm-1.5 6a.5.5 0 0 0-.5.5v1.5c0 1.106.894 2 2 2c1.106 0 2-.894 2-2c0-1.106-.894-2-2-2zM3.5 7.5c-1.106 0-2 .894-2 2c0 1.106.894 2 2 2h5c1.106 0 2-.894 2-2c0-1.106-.894-2-2-2zm5-6c-1.106 0-2 .894-2 2c0 1.106.894 2 2 2H10a.5.5 0 0 0 .5-.5V3.5c0-1.106-.894-2-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnapchatOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M22.364 17.707c-.326.806-1.092 1.183-1.528 1.36a7.102 7.102 0 0 1-1.035.323l-.023.088a1.767 1.767 0 0 1-.618.976a1.74 1.74 0 0 1-1.142.378h-.015l-.031-.001a3.526 3.526 0 0 1-.652-.089a3.239 3.239 0 0 0-.679-.075c-.13 0-.26.01-.393.03a1.912 1.912 0 0 0-.71.403a5.37 5.37 0 0 1-3.353 1.398l-.035.002h-.036c-.038 0-.082 0-.128-.002a2.5 2.5 0 0 1-.102.002h-.034l-.033-.002A5.365 5.365 0 0 1 8.463 21.1a1.926 1.926 0 0 0-.709-.402c-.13-.02-.26-.03-.39-.03c-.224.002-.447.03-.667.08c-.2.049-.404.08-.61.094a1.75 1.75 0 0 1-1.21-.352a1.78 1.78 0 0 1-.654-1.002a8.005 8.005 0 0 1-.024-.097a7.06 7.06 0 0 1-1.035-.322c-.436-.179-1.2-.556-1.525-1.358a1.885 1.885 0 0 1-.136-.603v-.012a1.74 1.74 0 0 1 .347-1.135c.249-.33.619-.572 1.052-.653a3.15 3.15 0 0 0 1.904-1.217l.016-.021l.016-.02c.217-.268.41-.56.573-.871a4.327 4.327 0 0 1-.315-.117a3.745 3.745 0 0 1-.943-.56c-.305-.256-.946-.926-.765-1.939l.023-.126l.044-.12a2.29 2.29 0 0 1 .837-1.098a2.217 2.217 0 0 1 1.341-.402c-.027-1.067.09-2.135.35-3.17l.026-.105l.041-.1a6.453 6.453 0 0 1 2.302-2.873a6.007 6.007 0 0 1 3.467-1.066l.344-.003a6.014 6.014 0 0 1 3.482 1.065a6.46 6.46 0 0 1 2.306 2.873l.041.1l.027.106a11.79 11.79 0 0 1 .349 3.16h.008c.281 0 .557.053.815.151c.368.12.69.338.935.622c.273.316.443.7.501 1.101l.014.093l.001.094c.021 1.13-.85 1.823-1.74 2.198l-.011.005a5.106 5.106 0 0 1-.307.113c.25.486.58.914.971 1.264c.476.425 1.03.723 1.616.878c.404.094.748.33.982.643c.249.332.365.734.345 1.132v.006a1.89 1.89 0 0 1-.134.603m-1.392-.562a.388.388 0 0 0 .028-.124a.24.24 0 0 0-.047-.156a.214.214 0 0 0-.133-.083a5.276 5.276 0 0 1-.342-.101l-.004-.002a5.368 5.368 0 0 1-.842-.361a5.544 5.544 0 0 1-1.076-.754a5.904 5.904 0 0 1-1.52-2.163l-.005-.012a1.07 1.07 0 0 1-.091-.86a.925.925 0 0 1 .352-.411l.05-.033c.232-.147.508-.24.724-.313l.04-.013l.01-.003l.011-.004a4.13 4.13 0 0 0 .193-.071l.026-.011l.018-.008h.002a2.854 2.854 0 0 0 .328-.17c.367-.225.451-.443.448-.6a.653.653 0 0 0-.15-.333a.59.59 0 0 0-.3-.187l-.006-.003a.82.82 0 0 0-.428-.058a.65.65 0 0 0-.172.053a2.23 2.23 0 0 1-.259.118l-.01.004a1.953 1.953 0 0 1-.16.054l-.006.001a2.15 2.15 0 0 1-.426.078a.704.704 0 0 1-.477-.17l.028-.51l.004-.061a10.3 10.3 0 0 0-.216-3.869a4.96 4.96 0 0 0-1.768-2.208a4.514 4.514 0 0 0-2.618-.8l-.362.002a4.507 4.507 0 0 0-2.615.802a4.953 4.953 0 0 0-1.764 2.207a10.3 10.3 0 0 0-.217 3.863c.012.191.023.383.033.574a.748.748 0 0 1-.526.171a2.075 2.075 0 0 1-.384-.054l-.007-.001a2.122 2.122 0 0 1-.521-.201a.492.492 0 0 0-.221-.048a.716.716 0 0 0-.442.131a.79.79 0 0 0-.287.38c-.044.248.137.455.363.611l.06.04a2.429 2.429 0 0 0 .393.203a2.616 2.616 0 0 0 .255.091a3.59 3.59 0 0 1 .724.314c.019.011.036.023.054.035a.92.92 0 0 1 .347.408a1.091 1.091 0 0 1-.092.862l-.003.01a6.825 6.825 0 0 1-.963 1.594a4.76 4.76 0 0 1-1.682 1.402c-.36.178-.743.308-1.14.385a.214.214 0 0 0-.133.083a.24.24 0 0 0-.048.156a.385.385 0 0 0 .029.125a.732.732 0 0 0 .283.316c.2.138.494.264.877.376c.114.033.236.065.366.096c.232.055.49.106.77.152c.126.02.18.24.253.597c.029.15.063.298.102.445a.282.282 0 0 0 .104.168c.053.04.119.059.183.05c.13-.007.257-.026.383-.057a4.685 4.685 0 0 1 1.008-.122c.243 0 .487.022.727.066a3.4 3.4 0 0 1 1.38.751A3.87 3.87 0 0 0 11.887 21c.03 0 .06-.001.09-.004c.037.003.086.004.137.004a3.87 3.87 0 0 0 2.422-1.016a3.41 3.41 0 0 1 1.38-.751c.24-.044.483-.066.728-.066a4.7 4.7 0 0 1 1.009.113c.125.03.253.047.38.052h.022a.24.24 0 0 0 .17-.052a.27.27 0 0 0 .096-.16a6.33 6.33 0 0 0 .102-.441c.073-.356.127-.574.252-.595c.306-.05.584-.106.832-.166c.107-.026.208-.054.304-.081c.35-.103.627-.217.825-.342a1.538 1.538 0 0 0 .173-.129a.617.617 0 0 0 .163-.219z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnapchatSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.967 18.289c-.194.48-1.165.877-2.808 1.147c-.153.026-.22.292-.308.727c-.036.18-.075.355-.125.54a.33.33 0 0 1-.117.195a.293.293 0 0 1-.209.063h-.025a2.477 2.477 0 0 1-.466-.064a5.79 5.79 0 0 0-1.233-.137c-.3 0-.595.027-.89.08a4.17 4.17 0 0 0-1.687.918A4.73 4.73 0 0 1 12.14 23c-.062 0-.122-.002-.167-.005a1.182 1.182 0 0 1-.11.005a4.725 4.725 0 0 1-2.96-1.242a4.179 4.179 0 0 0-1.687-.918a5.135 5.135 0 0 0-.89-.08a5.717 5.717 0 0 0-1.23.148c-.154.038-.31.062-.468.07a.307.307 0 0 1-.224-.06a.345.345 0 0 1-.128-.206a8.014 8.014 0 0 1-.124-.544c-.09-.437-.155-.704-.308-.73c-1.644-.269-2.614-.667-2.808-1.149A.47.47 0 0 1 1 18.137a.294.294 0 0 1 .058-.191a.261.261 0 0 1 .163-.102c1.36-.262 2.584-1.037 3.45-2.184a8.343 8.343 0 0 0 1.176-1.947l.004-.013a1.334 1.334 0 0 0 .113-1.053c-.21-.531-.912-.768-1.376-.925a4.182 4.182 0 0 1-.311-.112c-.412-.173-1.088-.54-.998-1.043a.967.967 0 0 1 .352-.465a.875.875 0 0 1 .54-.16a.598.598 0 0 1 .27.058a2.57 2.57 0 0 0 1.114.314a.895.895 0 0 0 .643-.21a81.188 81.188 0 0 0-.04-.7a12.59 12.59 0 0 1 .266-4.723a6.054 6.054 0 0 1 2.155-2.697a5.508 5.508 0 0 1 3.196-.98L12.218 1a5.516 5.516 0 0 1 3.2.98a6.062 6.062 0 0 1 2.16 2.698a12.59 12.59 0 0 1 .264 4.728l-.004.076l-.035.622c.165.14.371.213.583.208c.367-.03.724-.135 1.052-.312a.786.786 0 0 1 .342-.074c.136 0 .268.028.392.08l.006.004a.72.72 0 0 1 .367.229a.798.798 0 0 1 .185.407c.004.245-.166.61-1.005.964a5.07 5.07 0 0 1-.311.112c-.466.157-1.166.394-1.376.925a1.332 1.332 0 0 0 .11 1.051l.007.014a7.216 7.216 0 0 0 1.858 2.644a6.647 6.647 0 0 0 2.767 1.488a.262.262 0 0 1 .162.102a.3.3 0 0 1 .058.191a.474.474 0 0 1-.035.154z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.47 7.53a.75.75 0 0 0 1.06 0l.72-.72V17a.75.75 0 0 0 1.5 0V6.81l.72.72a.75.75 0 1 0 1.06-1.06l-2-2a.75.75 0 0 0-1.06 0l-2 2a.75.75 0 0 0 0 1.06m-4.72 9.66l.72-.72a.75.75 0 1 1 1.06 1.06l-2 2a.75.75 0 0 1-1.06 0l-2-2a.75.75 0 1 1 1.06-1.06l.72.72V7a.75.75 0 0 1 1.5 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.47 7.53a.75.75 0 0 0 1.06 0l.72-.72V17a.75.75 0 0 0 1.5 0V6.81l.72.72a.75.75 0 1 0 1.06-1.06l-2-2a.75.75 0 0 0-1.06 0l-2 2a.75.75 0 0 0 0 1.06m-4.72 9.66l.72-.72a.75.75 0 1 1 1.06 1.06l-2 2a.75.75 0 0 1-1.06 0l-2-2a.75.75 0 1 1 1.06-1.06l.72.72V7a.75.75 0 0 1 1.5 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StackOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.155 3.36c-.671-.35-1.58-.355-2.257-.006c-2.593 1.34-4.945 2.875-7.006 4.578c-.319.263-.545.646-.541 1.096c.004.449.234.826.551 1.085c2.044 1.671 4.398 3.208 6.943 4.537c.672.35 1.58.355 2.257.006c2.593-1.34 4.945-2.875 7.006-4.578c.319-.263.545-.646.541-1.096c-.004-.449-.235-.826-.551-1.085c-2.044-1.671-4.398-3.208-6.943-4.537m-1.57 1.327a1.07 1.07 0 0 1 .876.003c2.429 1.268 4.667 2.727 6.606 4.302c-1.95 1.601-4.182 3.055-6.653 4.331c-.24.124-.633.123-.875-.003c-2.429-1.268-4.667-2.727-6.607-4.302c1.951-1.601 4.183-3.055 6.654-4.331" clip-rule="evenodd"/><path fill="currentColor" d="M21.197 12.698a.75.75 0 0 1-.239 1.033l-6.107 3.813c-.829.518-1.857.757-2.851.757s-2.022-.24-2.851-.756l-6.04-3.77a.75.75 0 1 1 .794-1.273l6.04 3.77c.544.34 1.283.529 2.057.529c.773 0 1.512-.19 2.056-.53l6.108-3.812a.75.75 0 0 1 1.033.239"/><path fill="currentColor" d="M21.197 16.453a.75.75 0 0 1-.239 1.033L15.67 20.79c-1.058.66-2.38.971-3.67.971c-1.288 0-2.61-.31-3.669-.97l-5.222-3.26a.75.75 0 0 1 .795-1.273l5.222 3.26c.773.483 1.807.743 2.875.743s2.102-.26 2.875-.743l5.288-3.303a.75.75 0 0 1 1.034.239"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StackSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.039 3.582c-.6-.313-1.423-.318-2.027-.006c-2.578 1.332-4.915 2.858-6.96 4.548c-.274.226-.454.541-.451.902c.003.36.188.671.46.894c2.029 1.66 4.368 3.187 6.9 4.508c.6.313 1.423.318 2.027.006c2.578-1.332 4.915-2.858 6.96-4.548c.274-.226.454-.541.451-.902c-.003-.36-.188-.671-.46-.894c-2.029-1.66-4.368-3.187-6.9-4.508"/><path fill="currentColor" d="M20.958 13.731a.75.75 0 1 0-.794-1.272l-6.108 3.813c-.544.34-1.283.529-2.056.529c-.774 0-1.513-.19-2.057-.529l-6.04-3.77a.75.75 0 0 0-.794 1.273l6.04 3.77c.829.517 1.857.756 2.85.756c.995 0 2.023-.24 2.852-.757z"/><path fill="currentColor" d="M20.958 17.486a.75.75 0 1 0-.794-1.272l-5.29 3.303c-.772.482-1.806.743-2.874.743s-2.102-.26-2.875-.743l-5.222-3.26a.75.75 0 1 0-.795 1.273l5.222 3.26c1.058.66 2.381.97 3.67.97c1.289 0 2.611-.31 3.67-.971z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarHalfOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.419 4.23c.641-1.104 2.331-.65 2.331.628v11.68a2.75 2.75 0 0 1-1.469 2.434l-3.275 1.725c-1.297.683-2.809-.436-2.535-1.876l.7-3.676a1.25 1.25 0 0 0-.43-1.197l-2.5-2.07c-1.163-.964-.64-2.852.856-3.078l3.43-.518a1.25 1.25 0 0 0 .894-.608zm.831 1.557l-1.532 2.64a2.75 2.75 0 0 1-1.967 1.338l-3.43.518a.25.25 0 0 0-.122.44l2.499 2.07a2.75 2.75 0 0 1 .947 2.633l-.7 3.675a.25.25 0 0 0 .362.268l3.275-1.724a1.25 1.25 0 0 0 .668-1.106z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarHalfSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 2.462c0-.26-.208-.488-.45-.395a.987.987 0 0 0-.504.431L8.833 7.175a1.25 1.25 0 0 1-.894.61l-5.086.767c-.855.13-1.154 1.208-.489 1.76l3.79 3.138c.35.29.515.75.43 1.197l-.992 5.205a1 1 0 0 0 1.449 1.072l4.79-2.522c.046-.025.094-.046.143-.065c.271-.101.527-.322.527-.613z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.486 4.114c.675-1.162 2.353-1.162 3.028 0l2.065 3.56c.19.328.52.551.895.608l3.43.518c1.494.226 2.018 2.114.854 3.078l-2.499 2.07a1.25 1.25 0 0 0-.43 1.197l.7 3.676c.274 1.44-1.238 2.558-2.535 1.876L12.582 18.9a1.25 1.25 0 0 0-1.164 0l-3.412 1.797c-1.297.683-2.809-.436-2.535-1.876l.7-3.676a1.25 1.25 0 0 0-.43-1.197l-2.5-2.07c-1.163-.964-.64-2.852.856-3.078l3.43-.518a1.25 1.25 0 0 0 .894-.609zm1.73.753a.25.25 0 0 0-.432 0l-2.066 3.56a2.75 2.75 0 0 1-1.967 1.338l-3.43.518a.25.25 0 0 0-.122.44l2.499 2.07a2.75 2.75 0 0 1 .947 2.632l-.7 3.676a.25.25 0 0 0 .362.268l3.412-1.796a2.75 2.75 0 0 1 2.562 0l3.412 1.796a.25.25 0 0 0 .362-.268l-.7-3.676a2.75 2.75 0 0 1 .947-2.632l2.5-2.07a.25.25 0 0 0-.123-.44l-3.43-.518a2.75 2.75 0 0 1-1.967-1.339z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.865 2.996a1 1 0 0 0-1.73 0L8.421 7.674a1.25 1.25 0 0 1-.894.608L2.44 9.05c-.854.13-1.154 1.208-.488 1.76l3.789 3.138c.35.291.515.75.43 1.197L5.18 20.35a1 1 0 0 0 1.448 1.072l4.79-2.522a1.25 1.25 0 0 1 1.164 0l4.79 2.522a1 1 0 0 0 1.448-1.072l-.991-5.205a1.25 1.25 0 0 1 .43-1.197l3.79-3.139c.665-.55.365-1.63-.49-1.759l-5.085-.768a1.25 1.25 0 0 1-.895-.608z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.25a.75.75 0 0 1 .75.75v1a.75.75 0 0 1-1.5 0V2a.75.75 0 0 1 .75-.75"/><path fill="currentColor" fill-rule="evenodd" d="M6.25 12a5.75 5.75 0 1 1 11.5 0a5.75 5.75 0 0 1-11.5 0M12 7.75a4.25 4.25 0 1 0 0 8.5a4.25 4.25 0 0 0 0-8.5" clip-rule="evenodd"/><path fill="currentColor" d="M5.46 4.399a.75.75 0 0 0-1.061 1.06l.707.707a.75.75 0 1 0 1.06-1.06zM22.75 12a.75.75 0 0 1-.75.75h-1a.75.75 0 0 1 0-1.5h1a.75.75 0 0 1 .75.75m-3.149-6.54a.75.75 0 1 0-1.06-1.061l-.707.707a.75.75 0 1 0 1.06 1.06zM12 20.25a.75.75 0 0 1 .75.75v1a.75.75 0 0 1-1.5 0v-1a.75.75 0 0 1 .75-.75m6.894-2.416a.75.75 0 1 0-1.06 1.06l.707.707a.75.75 0 1 0 1.06-1.06zM3.75 12a.75.75 0 0 1-.75.75H2a.75.75 0 0 1 0-1.5h1a.75.75 0 0 1 .75.75m2.416 6.894a.75.75 0 0 0-1.06-1.06l-.707.707a.75.75 0 0 0 1.06 1.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.25a.75.75 0 0 1 .75.75v1a.75.75 0 0 1-1.5 0V2a.75.75 0 0 1 .75-.75m0 5a5.75 5.75 0 1 0 0 11.5a5.75 5.75 0 0 0 0-11.5M5.46 4.399a.75.75 0 0 0-1.061 1.06l.707.707a.75.75 0 1 0 1.06-1.06zM22.75 12a.75.75 0 0 1-.75.75h-1a.75.75 0 0 1 0-1.5h1a.75.75 0 0 1 .75.75m-3.149-6.54a.75.75 0 1 0-1.06-1.061l-.707.707a.75.75 0 1 0 1.06 1.06zM12 20.25a.75.75 0 0 1 .75.75v1a.75.75 0 0 1-1.5 0v-1a.75.75 0 0 1 .75-.75m6.894-2.416a.75.75 0 1 0-1.06 1.06l.707.707a.75.75 0 1 0 1.06-1.06zM3.75 12a.75.75 0 0 1-.75.75H2a.75.75 0 0 1 0-1.5h1a.75.75 0 0 1 .75.75m2.416 6.894a.75.75 0 0 0-1.06-1.06l-.707.707a.75.75 0 0 0 1.06 1.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelegramOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.788 14.02a.746.746 0 0 0 .132.031a456.056 456.056 0 0 1 .844 2.002c.503 1.202 1.01 2.44 1.121 2.796c.139.438.285.736.445.94c.083.104.178.196.29.266a.88.88 0 0 0 .186.088c.32.12.612.07.795.009a1.313 1.313 0 0 0 .304-.15L9.91 20l2.826-1.762l3.265 2.502c.048.037.1.068.156.093c.392.17.772.23 1.13.182c.356-.05.639-.199.85-.368a1.992 1.992 0 0 0 .564-.728l.009-.022l.003-.008l.002-.004v-.002l.001-.001a.756.756 0 0 0 .04-.133l2.98-15.025a.752.752 0 0 0 .014-.146c0-.44-.166-.859-.555-1.112c-.334-.217-.705-.227-.94-.209c-.252.02-.486.082-.643.132a3.458 3.458 0 0 0-.26.094l-.011.005l-16.714 6.556l-.002.001a2.296 2.296 0 0 0-.167.069a2.522 2.522 0 0 0-.38.212c-.227.155-.75.581-.661 1.285c.07.56.454.905.689 1.071c.128.091.25.156.34.199c.04.02.126.054.163.07l.01.003zm14.138-9.152h-.002a.785.785 0 0 1-.026.011L3.164 11.444a.818.818 0 0 1-.026.01l-.01.003a1.126 1.126 0 0 0-.09.04a.851.851 0 0 0 .086.043l3.142 1.058a.75.75 0 0 1 .16.076l10.377-6.075l.01-.005a1.59 1.59 0 0 1 .124-.068c.072-.037.187-.091.317-.131c.09-.028.357-.107.645-.014a.854.854 0 0 1 .588.689a.84.84 0 0 1 .003.424c-.07.275-.262.489-.437.653c-.15.14-2.096 2.016-4.015 3.868l-2.613 2.52l-.465.45l5.872 4.502a.536.536 0 0 0 .251.04a.229.229 0 0 0 .117-.052a.495.495 0 0 0 .103-.12l.002-.001l2.89-14.573a1.858 1.858 0 0 0-.267.086zm-8.461 12.394l-1.172-.898l-.284 1.805zm-2.247-2.68l1.165-1.125l2.613-2.522l.973-.938l-6.52 3.817l.035.082a339.2 339.2 0 0 1 1.22 2.92l.283-1.8a.747.747 0 0 1 .231-.435" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelegramSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M18.483 19.79v-.002l.018-.043L21.5 4.625v-.048c0-.377-.14-.706-.442-.903c-.265-.173-.57-.185-.784-.169a2.681 2.681 0 0 0-.586.12a3.23 3.23 0 0 0-.24.088l-.013.005l-16.72 6.559l-.005.002a1.353 1.353 0 0 0-.149.061a2.27 2.27 0 0 0-.341.19c-.215.148-.624.496-.555 1.048c.057.458.372.748.585.899a2.062 2.062 0 0 0 .403.22l.032.014l.01.003l.007.003l2.926.985c-.01.183.008.37.057.555l1.465 5.559a1.5 1.5 0 0 0 2.834.196l2.288-2.446l3.929 3.012l.056.024c.357.156.69.205.995.164c.305-.042.547-.17.729-.315a1.742 1.742 0 0 0 .49-.635l.008-.017l.003-.006zM7.135 13.875a.3.3 0 0 1 .13-.33l9.921-6.3s.584-.355.563 0c0 0 .104.062-.209.353c-.296.277-7.071 6.818-7.757 7.48a.278.278 0 0 0-.077.136L8.6 19.434z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimerOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 3.018a.75.75 0 1 1 0-1.5h3.536a.75.75 0 0 1 0 1.5zm-3.47.452a.75.75 0 0 1 0 1.06l-2.5 2.5a.75.75 0 0 1-1.06-1.06l2.5-2.5a.75.75 0 0 1 1.06 0M12 5.75A7.25 7.25 0 1 0 19.25 13a.75.75 0 1 1 1.5 0A8.75 8.75 0 1 1 12 4.25a.75.75 0 0 1 0 1.5"/><path fill="currentColor" d="M17.188 8.364a.75.75 0 0 0-1.052-1.052l-3.17 2.465l-2.072 1.48a1.684 1.684 0 1 0 2.35 2.349l1.479-2.072z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimerSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 3.578a.75.75 0 0 1 0-1.5h3.536a.75.75 0 0 1 0 1.5zm-3.47.452a.75.75 0 0 1 0 1.06l-2.5 2.5a.75.75 0 0 1-1.06-1.06l2.5-2.5a.75.75 0 0 1 1.06 0"/><path fill="currentColor" fill-rule="evenodd" d="M12 5.06a8.5 8.5 0 1 0 0 17a8.5 8.5 0 0 0 0-17m4.99 3.711a.5.5 0 0 0-.7-.701l-3.175 2.468l-2.075 1.483a1.434 1.434 0 1 0 2 2l1.482-2.076z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleOffOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 15.5a3.496 3.496 0 0 1-3.464-3.868A3.496 3.496 0 0 1 8 8.5a3.496 3.496 0 0 1 3.464 3.868A3.496 3.496 0 0 1 8 15.5"/><path fill="currentColor" fill-rule="evenodd" d="M8.074 18.75h7.852a6.75 6.75 0 0 0 0-13.5H8.074a6.75 6.75 0 0 0 0 13.5m0-1.5a5.25 5.25 0 1 1 0-10.5h7.852a5.25 5.25 0 1 1 0 10.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleOffSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M22.346 13.014a6.5 6.5 0 0 1-6.42 5.486H8.074a6.5 6.5 0 1 1 0-13h7.852a6.5 6.5 0 0 1 6.42 7.514M8 15.5a3.496 3.496 0 0 1-3.464-3.868A3.496 3.496 0 0 1 8 8.5a3.496 3.496 0 0 1 3.464 3.868A3.496 3.496 0 0 1 8 15.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleOnOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 15.5a3.496 3.496 0 0 0 3.464-3.868A3.496 3.496 0 0 0 16 8.5a3.496 3.496 0 0 0-3.464 3.868A3.496 3.496 0 0 0 16 15.5"/><path fill="currentColor" fill-rule="evenodd" d="M15.926 18.75H8.074a6.75 6.75 0 0 1 0-13.5h7.852a6.75 6.75 0 0 1 0 13.5m0-1.5a5.25 5.25 0 1 0 0-10.5H8.074a5.25 5.25 0 0 0 0 10.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleOnSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.654 13.014a6.5 6.5 0 0 0 6.42 5.486h7.852a6.5 6.5 0 1 0 0-13H8.074a6.5 6.5 0 0 0-6.42 7.514M16 15.5a3.496 3.496 0 0 0 3.464-3.868A3.496 3.496 0 0 0 16 8.5a3.496 3.496 0 0 0-3.464 3.868A3.496 3.496 0 0 0 16 15.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrashAltOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 2.25a.75.75 0 0 0-.75.75v.75H5a.75.75 0 0 0 0 1.5h14a.75.75 0 0 0 0-1.5h-4.25V3a.75.75 0 0 0-.75-.75zM13.06 15l1.47 1.47a.75.75 0 1 1-1.06 1.06L12 16.06l-1.47 1.47a.75.75 0 1 1-1.06-1.06L10.94 15l-1.47-1.47a.75.75 0 1 1 1.06-1.06L12 13.94l1.47-1.47a.75.75 0 1 1 1.06 1.06z"/><path fill="currentColor" fill-rule="evenodd" d="M5.991 7.917a.75.75 0 0 1 .746-.667h10.526a.75.75 0 0 1 .746.667l.2 1.802c.363 3.265.363 6.56 0 9.826l-.02.177a2.853 2.853 0 0 1-2.44 2.51a27.04 27.04 0 0 1-7.498 0a2.853 2.853 0 0 1-2.44-2.51l-.02-.177a44.489 44.489 0 0 1 0-9.826zm1.417.833l-.126 1.134a42.99 42.99 0 0 0 0 9.495l.02.177a1.353 1.353 0 0 0 1.157 1.191c2.35.329 4.733.329 7.082 0a1.353 1.353 0 0 0 1.157-1.19l.02-.178c.35-3.155.35-6.34 0-9.495l-.126-1.134z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrashAltSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.25 3a.75.75 0 0 1 .75-.75h4a.75.75 0 0 1 .75.75v.75H19a.75.75 0 0 1 0 1.5H5a.75.75 0 0 1 0-1.5h4.25z"/><path fill="currentColor" fill-rule="evenodd" d="M6.24 7.945a.5.5 0 0 1 .497-.445h10.526a.5.5 0 0 1 .497.445l.2 1.801a44.213 44.213 0 0 1 0 9.771l-.02.177a2.603 2.603 0 0 1-2.226 2.29a26.788 26.788 0 0 1-7.428 0a2.603 2.603 0 0 1-2.227-2.29l-.02-.177a44.239 44.239 0 0 1 0-9.77zm8.29 4.525a.75.75 0 0 1 0 1.06L13.06 15l1.47 1.47a.75.75 0 1 1-1.06 1.06L12 16.06l-1.47 1.47a.75.75 0 1 1-1.06-1.06L10.94 15l-1.47-1.47a.75.75 0 1 1 1.06-1.06L12 13.94l1.47-1.47a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrashOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 2.25a.75.75 0 0 0-.75.75v.75H5a.75.75 0 0 0 0 1.5h14a.75.75 0 0 0 0-1.5h-4.25V3a.75.75 0 0 0-.75-.75zm0 8.4a.75.75 0 0 1 .75.75v7a.75.75 0 0 1-1.5 0v-7a.75.75 0 0 1 .75-.75m4.75.75a.75.75 0 0 0-1.5 0v7a.75.75 0 0 0 1.5 0z"/><path fill="currentColor" fill-rule="evenodd" d="M5.991 7.917a.75.75 0 0 1 .746-.667h10.526a.75.75 0 0 1 .746.667l.2 1.802c.363 3.265.363 6.56 0 9.826l-.02.177a2.853 2.853 0 0 1-2.44 2.51a27.04 27.04 0 0 1-7.498 0a2.853 2.853 0 0 1-2.44-2.51l-.02-.177a44.489 44.489 0 0 1 0-9.826zm1.417.833l-.126 1.134a42.99 42.99 0 0 0 0 9.495l.02.177a1.353 1.353 0 0 0 1.157 1.191c2.35.329 4.733.329 7.082 0a1.353 1.353 0 0 0 1.157-1.19l.02-.178c.35-3.155.35-6.34 0-9.495l-.126-1.134z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrashSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.25 3a.75.75 0 0 1 .75-.75h4a.75.75 0 0 1 .75.75v.75H19a.75.75 0 0 1 0 1.5H5a.75.75 0 0 1 0-1.5h4.25z"/><path fill="currentColor" fill-rule="evenodd" d="M6.24 7.945a.5.5 0 0 1 .497-.445h10.526a.5.5 0 0 1 .497.445l.2 1.801a44.213 44.213 0 0 1 0 9.771l-.02.177a2.603 2.603 0 0 1-2.226 2.29a26.788 26.788 0 0 1-7.428 0a2.603 2.603 0 0 1-2.227-2.29l-.02-.177a44.239 44.239 0 0 1 0-9.77zm4.51 3.455a.75.75 0 0 0-1.5 0v7a.75.75 0 0 0 1.5 0zm4 0a.75.75 0 0 0-1.5 0v7a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrelloOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.873 16.407a.949.949 0 0 0 .074-.368v-9.08A.966.966 0 0 0 9.973 6H6.972A.966.966 0 0 0 6 6.959v9.08c0 .254.103.498.285.678c.182.18.429.282.687.283h3a.982.982 0 0 0 .69-.281a.96.96 0 0 0 .211-.312m6.843-3.689a.955.955 0 0 0 .284-.679v-5.08A.966.966 0 0 0 17.029 6h-3.002a.968.968 0 0 0-.974.959v5.08c0 .255.103.5.286.68c.182.18.43.28.688.281h3.002a.981.981 0 0 0 .686-.283"/><path fill="currentColor" fill-rule="evenodd" d="M17.258 2.833a47.721 47.721 0 0 0-10.516 0c-2.012.225-3.637 1.81-3.873 3.832a45.921 45.921 0 0 0 0 10.67c.236 2.022 1.86 3.607 3.873 3.832a47.77 47.77 0 0 0 10.516 0c2.012-.225 3.637-1.81 3.873-3.832a45.925 45.925 0 0 0 0-10.67c-.236-2.022-1.86-3.607-3.873-3.832m-10.35 1.49a46.22 46.22 0 0 1 10.184 0c1.33.15 2.395 1.199 2.55 2.517a44.421 44.421 0 0 1 0 10.32a2.89 2.89 0 0 1-2.55 2.516a46.216 46.216 0 0 1-10.184 0a2.89 2.89 0 0 1-2.55-2.516a44.421 44.421 0 0 1 0-10.32a2.89 2.89 0 0 1 2.55-2.516" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrelloSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.77 3.082a47.472 47.472 0 0 1 10.46 0c1.899.212 3.43 1.707 3.653 3.613a45.67 45.67 0 0 1 0 10.61c-.223 1.906-1.754 3.401-3.652 3.614a47.468 47.468 0 0 1-10.461 0c-1.899-.213-3.43-1.708-3.653-3.613a45.672 45.672 0 0 1 0-10.611C3.34 4.789 4.871 3.294 6.77 3.082m4.103 13.325a.949.949 0 0 0 .074-.368v-9.08A.966.966 0 0 0 9.973 6H6.972A.966.966 0 0 0 6 6.959v9.08c0 .254.103.498.285.678c.182.18.429.282.687.283h3a.982.982 0 0 0 .69-.281a.96.96 0 0 0 .211-.312m6.842-3.69A.955.955 0 0 0 18 12.04V6.96a.966.966 0 0 0-.971-.96h-3.002a.968.968 0 0 0-.974.959v5.08c0 .255.103.5.286.68c.182.18.43.28.688.281h3.002a.981.981 0 0 0 .686-.283" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TumblrOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.507 2.252a.752.752 0 0 1 .05-.002h2.546a.75.75 0 0 1 .75.75v3.714h2.737a.75.75 0 0 1 .75.75v2.615a.75.75 0 0 1-.75.75h-2.737v4.682a.746.746 0 0 1-.008.112a.776.776 0 0 0 .85.886a.749.749 0 0 1 .057-.004a3.303 3.303 0 0 0 1.15-.249a.75.75 0 0 1 1 .46l.81 2.467a.75.75 0 0 1-.092.655c-.17.252-.431.441-.663.578c-.249.148-.543.28-.856.391a6.582 6.582 0 0 1-2.12.38c-2.097.032-3.615-.726-4.602-1.825a5.468 5.468 0 0 1-1.385-3.58v-4.953H7a.75.75 0 0 1-.75-.75V7.83a.75.75 0 0 1 .485-.702a4.577 4.577 0 0 0 2.95-4.026a.88.88 0 0 1 .822-.85m.618 1.498A6.076 6.076 0 0 1 7.75 8.324v1.005h.994a.75.75 0 0 1 .75.75v5.702c0 .805.302 1.8 1.001 2.579c.68.757 1.775 1.354 3.468 1.327h.012a5.08 5.08 0 0 0 1.625-.294c.209-.074.388-.154.527-.23l-.42-1.276a4.804 4.804 0 0 1-.879.116a2.277 2.277 0 0 1-2.475-2.543v-5.38a.75.75 0 0 1 .75-.75h2.737V8.213h-2.737a.75.75 0 0 1-.75-.75V3.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TumblrSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.256 17.154a.209.209 0 0 0-.268-.128c-.39.138-.798.215-1.211.229a1.525 1.525 0 0 1-1.676-1.73a.212.212 0 0 0 .002-.03V10.08h3.287a.2.2 0 0 0 .2-.2V7.664a.2.2 0 0 0-.2-.2h-3.287V3.2a.2.2 0 0 0-.2-.2h-2.346a.131.131 0 0 0-.122.13a5.326 5.326 0 0 1-3.306 4.65a.205.205 0 0 0-.129.188V9.88c0 .11.09.2.2.2h1.544v5.702c0 1.944 1.438 4.717 5.23 4.656c1.212 0 2.557-.499 2.965-.944a.17.17 0 0 0 .03-.17z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwitchOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.75 3.4c0-.635.515-1.15 1.15-1.15h16.2c.635 0 1.15.515 1.15 1.15v11.472c0 .305-.121.597-.337.813l-2.722 2.722a1.15 1.15 0 0 1-.813.337h-4.54l-3.45 3.451c-.6.599-1.622.175-1.622-.672v-2.78H3.9a1.15 1.15 0 0 1-1.15-1.15zm1.5.35v13.494h3.866c.635 0 1.15.515 1.15 1.15v1.802l2.614-2.615a1.15 1.15 0 0 1 .814-.337h4.539l2.517-2.517V3.75zm6.219 3.2a.75.75 0 0 1 .75.75v4.272a.75.75 0 1 1-1.5 0V7.7a.75.75 0 0 1 .75-.75m5.016 0a.75.75 0 0 1 .75.75v4.272a.75.75 0 1 1-1.5 0V7.7a.75.75 0 0 1 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwitchSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.9 2.5a.9.9 0 0 0-.9.9v14.194a.9.9 0 0 0 .9.9h4.116v3.03a.7.7 0 0 0 1.194.494l3.525-3.524h4.643a.9.9 0 0 0 .636-.264l2.722-2.722a.9.9 0 0 0 .264-.636V3.4a.9.9 0 0 0-.9-.9zm7.319 5.2a.75.75 0 0 0-1.5 0v4.272a.75.75 0 1 0 1.5 0zm5.016 0a.75.75 0 0 0-1.5 0v4.272a.75.75 0 1 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwitterOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.022 3.343c.508-.087 1.077-.116 1.613-.025a4.85 4.85 0 0 1 2.54 1.273c.456.01.905-.08 1.302-.208a5.36 5.36 0 0 0 1.098-.501l.009-.006a.75.75 0 0 1 1.042 1.037c-.207.315-.496.877-.819 1.507l-.155.301c-.185.36-.375.724-.552 1.036c-.111.196-.23.395-.35.567v.274A12.34 12.34 0 0 1 8.287 21.03a12.32 12.32 0 0 1-6.694-1.97a.75.75 0 0 1 .5-1.375a7.471 7.471 0 0 0 4.033-.642a4.858 4.858 0 0 1-2.61-2.922a.75.75 0 0 1 .147-.722l.01-.01A4.848 4.848 0 0 1 2.05 9.793v-.052a.75.75 0 0 1 .553-.724A4.84 4.84 0 0 1 2.09 6.84a4.9 4.9 0 0 1 .65-2.442a.75.75 0 0 1 1.232-.1a10.89 10.89 0 0 0 7.006 3.93a4.85 4.85 0 0 1 2.562-4.406c.402-.214.934-.385 1.482-.479m-11.28 7.548a3.35 3.35 0 0 0 2.503 2.164a.75.75 0 0 1 .072 1.453c-.272.083-.551.14-.834.173a3.358 3.358 0 0 0 2.59 1.3a.75.75 0 0 1 .45 1.339a8.97 8.97 0 0 1-3.548 1.695a10.82 10.82 0 0 0 3.313.515h.009A10.838 10.838 0 0 0 19.25 8.607v-.535a.75.75 0 0 1 .186-.495c.07-.079.19-.261.36-.56c.16-.282.338-.622.523-.981l.033-.066a4.992 4.992 0 0 1-1.593.097a.75.75 0 0 1-.47-.237a3.35 3.35 0 0 0-1.904-1.032a3.42 3.42 0 0 0-1.11.025a3.605 3.605 0 0 0-1.028.323a3.35 3.35 0 0 0-1.678 3.74a.75.75 0 0 1-.767.925a12.39 12.39 0 0 1-8.149-3.627a3.41 3.41 0 0 0-.063.657v.002a3.34 3.34 0 0 0 1.486 2.785A.75.75 0 0 1 4.64 11a4.798 4.798 0 0 1-.897-.11" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwitterSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.81 6.227c.058-.08-.028-.185-.12-.149a8.489 8.489 0 0 1-2.05.513a4.13 4.13 0 0 0 1.72-2.02c.034-.086-.06-.163-.14-.118c-.727.401-1.51.69-2.325.857a.1.1 0 0 1-.093-.03a4.1 4.1 0 0 0-6.991 3.65a.102.102 0 0 1-.104.123a11.64 11.64 0 0 1-8.224-4.17a.098.098 0 0 0-.163.015a4.16 4.16 0 0 0-.48 1.943a4.09 4.09 0 0 0 1.82 3.41a4.05 4.05 0 0 1-1.709-.43c-.068-.035-.15.014-.147.09a4.23 4.23 0 0 0 .933 2.468A4.1 4.1 0 0 0 6.1 13.79a3.93 3.93 0 0 1-1.1.17a4.901 4.901 0 0 1-.606-.045c-.075-.01-.136.06-.11.13A4.11 4.11 0 0 0 8.06 16.73a8.22 8.22 0 0 1-5.625 1.741c-.106-.007-.155.134-.064.188a11.57 11.57 0 0 0 5.919 1.62A11.59 11.59 0 0 0 20 8.6v-.48a.1.1 0 0 1 .04-.08a8.433 8.433 0 0 0 1.77-1.813"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UmbrellaOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.75 2a.75.75 0 0 0-1.5 0v.278c-4.984.38-8.92 4.505-8.999 9.567a.917.917 0 0 0 .766.918c2.727.455 5.479.7 8.233.739V19a.75.75 0 0 1-1.5 0v-.5a.75.75 0 0 0-1.5 0v.5a2.25 2.25 0 0 0 4.5 0v-5.498a54.635 54.635 0 0 0 8.233-.739a.917.917 0 0 0 .766-.918a9.753 9.753 0 0 0-8.999-9.567zm7.476 9.366a8.25 8.25 0 0 0-16.452 0c5.45.854 11.001.854 16.452 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UmbrellaSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.75 1.5a.75.75 0 0 0-1.5 0v1.03a9.5 9.5 0 0 0-8.749 9.319a.667.667 0 0 0 .558.668a54.39 54.39 0 0 0 8.191.735V19a.75.75 0 0 1-1.5 0v-.5a.75.75 0 0 0-1.5 0v.5a2.25 2.25 0 0 0 4.5 0v-5.748a54.39 54.39 0 0 0 8.192-.735a.667.667 0 0 0 .557-.668c-.077-4.926-3.902-8.941-8.749-9.32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UniversityOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.612 3.302c.243-.07.5-.07.743 0c.518.147 1.04.283 1.564.42c2.461.641 4.96 1.293 7.184 3.104l1.024.834c.415.338.623.84.623 1.34v7a.75.75 0 0 1-1.5 0v-4.943l-.163.133a11.946 11.946 0 0 1-2.398 1.513c.04.091.061.191.061.297v4.294a2.75 2.75 0 0 1-1.751 2.562l-4 1.56a2.75 2.75 0 0 1-1.998 0l-4-1.56a2.75 2.75 0 0 1-1.751-2.562V13c0-.108.023-.211.064-.304c-.83-.399-1.64-.89-2.417-1.522l-1.024-.834c-.83-.677-.83-2.003 0-2.68l1.04-.85c2.207-1.8 4.689-2.449 7.132-3.087a74.375 74.375 0 0 0 1.567-.421m9.638 5.699c0-.09-.036-.15-.07-.178l-1.024-.834C18 6.5 16.078 5.843 13.64 5.202a90.449 90.449 0 0 1-1.656-.446c-.57.161-1.124.307-1.662.449c-2.42.636-4.529 1.191-6.46 2.768l-1.041.849c-.035.028-.071.087-.071.177s.036.15.07.178l1.025.834c1.948 1.587 4.076 2.146 6.515 2.787c.537.14 1.088.286 1.656.446c.57-.161 1.124-.307 1.662-.449c2.42-.636 4.529-1.191 6.46-2.767l1.041-.85c.035-.028.071-.087.071-.177m-7.294 5.276c1.1-.287 2.207-.577 3.294-.972v3.989c0 .515-.316.977-.796 1.165l-4 1.559a1.25 1.25 0 0 1-.908 0l-4-1.56a1.25 1.25 0 0 1-.796-1.164v-3.998c1.099.4 2.219.692 3.33.982c.525.137 1.047.273 1.565.42c.243.07.5.07.743 0c.519-.148 1.042-.284 1.568-.421" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UniversitySolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.355 3.302a1.34 1.34 0 0 0-.743 0c-.519.148-1.042.284-1.568.421C7.701 4.335 5.12 5.01 2.913 6.81l-1.04.85c-.83.677-.83 2.003 0 2.68l1.024.834a12.01 12.01 0 0 0 2.353 1.491v4.629a2.75 2.75 0 0 0 1.751 2.562l4 1.56a2.75 2.75 0 0 0 1.998 0l4-1.56a2.75 2.75 0 0 0 1.751-2.562v-4.62a11.948 11.948 0 0 0 2.337-1.484l.163-.133V16a.75.75 0 0 0 1.5 0V9c0-.5-.208-1.002-.623-1.34l-1.024-.834c-2.224-1.81-4.913-2.512-7.184-3.104c-.524-.137-1.046-.273-1.564-.42M8.3 12.812a.75.75 0 0 0-.598 1.376c1.278.555 2.6 1.01 3.956 1.358c.236.06.484.06.72 0a26.176 26.176 0 0 0 3.945-1.359a.75.75 0 0 0-.6-1.374c-1.197.522-2.435.95-3.705 1.277A24.828 24.828 0 0 1 8.3 12.812" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnlockOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 16a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0"/><path fill="currentColor" fill-rule="evenodd" d="M9.81 4.005a3.2 3.2 0 0 1 4.164 1.808l.075.192c.093.238.16.486.198.738l.217 1.423l1.483-.226l-.217-1.423a4.96 4.96 0 0 0-.283-1.057l-.075-.193a4.7 4.7 0 0 0-9.024 2.418l.03.204c.056.362.151.717.284 1.058l.655 1.675l-.382.03a2.361 2.361 0 0 0-2.142 1.972a20.891 20.891 0 0 0 0 6.752a2.361 2.361 0 0 0 2.142 1.972l1.496.12c2.376.19 4.762.19 7.138 0l1.496-.12a2.361 2.361 0 0 0 2.142-1.972a20.89 20.89 0 0 0 0-6.752a2.361 2.361 0 0 0-2.142-1.972l-1.496-.12a44.845 44.845 0 0 0-6.69-.033l-.82-2.098a3.46 3.46 0 0 1-.197-.738L7.83 7.46a3.2 3.2 0 0 1 1.98-3.455m5.64 8.023a43.366 43.366 0 0 0-6.9 0l-1.496.12a.861.861 0 0 0-.781.719a19.39 19.39 0 0 0 0 6.266a.861.861 0 0 0 .781.72l1.497.12c2.296.183 4.602.183 6.898 0l1.496-.12a.861.861 0 0 0 .782-.72a19.39 19.39 0 0 0 0-6.266a.861.861 0 0 0-.782-.72z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnlockSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.533 3.811A3.2 3.2 0 0 0 7.83 7.46l.03.203c.04.253.106.5.199.738l.915 2.342a44.612 44.612 0 0 1 6.574.039l1.496.12a2.111 2.111 0 0 1 1.915 1.763a20.64 20.64 0 0 1 0 6.67a2.111 2.111 0 0 1-1.915 1.764l-1.496.12a44.613 44.613 0 0 1-7.098 0l-1.496-.12a2.111 2.111 0 0 1-1.915-1.764a20.641 20.641 0 0 1 0-6.67A2.111 2.111 0 0 1 6.955 10.9l.457-.036l-.75-1.918A4.96 4.96 0 0 1 6.38 7.89l-.031-.204a4.7 4.7 0 0 1 9.024-2.418l.075.193c.133.34.228.695.283 1.057l.141.928a.5.5 0 0 1-.419.57l-.494.075a.5.5 0 0 1-.57-.419l-.14-.928a3.459 3.459 0 0 0-.198-.738l-.075-.192a3.2 3.2 0 0 0-3.442-2.002M12 14.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.738 16.627a1.492 1.492 0 0 1-1.487-1.362a37.117 37.117 0 0 1-.107-4.845a35.68 35.68 0 0 1-.253-.018l-1.49-.108a1.26 1.26 0 0 1-.97-1.935c1.097-1.716 3.106-3.636 4.76-4.824a1.388 1.388 0 0 1 1.619 0c1.653 1.188 3.662 3.108 4.759 4.824a1.26 1.26 0 0 1-.97 1.935l-1.49.108l-.253.018c.07 1.616.034 3.234-.107 4.845a1.492 1.492 0 0 1-1.487 1.362zm-.056-6.865a35.624 35.624 0 0 0 .063 5.365h2.51c.156-1.784.177-3.577.064-5.365a.75.75 0 0 1 .711-.796c.324-.016.647-.036.97-.06l1.081-.078a14.556 14.556 0 0 0-3.55-3.646L12 4.801l-.531.381a14.555 14.555 0 0 0-3.55 3.646L9 8.907c.323.023.647.043.97.059a.75.75 0 0 1 .711.796" clip-rule="evenodd"/><path fill="currentColor" d="M5.75 17a.75.75 0 0 0-1.5 0v2c0 .966.784 1.75 1.75 1.75h12A1.75 1.75 0 0 0 19.75 19v-2a.75.75 0 0 0-1.5 0v2a.25.25 0 0 1-.25.25H6a.25.25 0 0 1-.25-.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 16.25a.75.75 0 0 1 .75.75v2c0 .138.112.25.25.25h12a.25.25 0 0 0 .25-.25v-2a.75.75 0 0 1 1.5 0v2A1.75 1.75 0 0 1 18 20.75H6A1.75 1.75 0 0 1 4.25 19v-2a.75.75 0 0 1 .75-.75m5.738-.123a.992.992 0 0 1-.989-.905a36.618 36.618 0 0 1-.08-5.27a42.17 42.17 0 0 1-.741-.048l-1.49-.109a.76.76 0 0 1-.585-1.167a15.555 15.555 0 0 1 4.032-4.258l.597-.429a.888.888 0 0 1 1.036 0l.597.43a15.556 15.556 0 0 1 4.032 4.257a.76.76 0 0 1-.585 1.167l-1.49.109c-.246.018-.493.034-.74.047c.1 1.757.072 3.518-.082 5.27a.992.992 0 0 1-.988.906z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserBlockOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 3.25a4.25 4.25 0 1 0 0 8.5a4.25 4.25 0 0 0 0-8.5M7.25 7.5a2.75 2.75 0 1 1 5.5 0a2.75 2.75 0 0 1-5.5 0" clip-rule="evenodd"/><path fill="currentColor" d="M3.75 17A2.25 2.25 0 0 1 6 14.75h.34c.027 0 .053.004.078.012l.866.283a8.75 8.75 0 0 0 3.071.425a.303.303 0 0 0 .28-.244c.074-.34.176-.669.305-.985a.22.22 0 0 0-.223-.3a7.251 7.251 0 0 1-2.967-.322l-.866-.283a1.752 1.752 0 0 0-.543-.086H6A3.75 3.75 0 0 0 2.25 17v1.188c0 .754.546 1.396 1.29 1.517c2.595.424 5.221.59 7.84.5c.163-.005.251-.194.16-.328a5.998 5.998 0 0 1-.508-.903a.42.42 0 0 0-.387-.25a38.6 38.6 0 0 1-6.864-.5a.037.037 0 0 1-.031-.036z"/><path fill="currentColor" fill-rule="evenodd" d="M12 16.5c0 .972.308 1.872.832 2.607A4.48 4.48 0 0 0 16.5 21a4.5 4.5 0 1 0-4.5-4.5m4.5 3a2.985 2.985 0 0 1-1.524-.415l4.109-4.109A3 3 0 0 1 16.5 19.5m-2.585-1.476l4.109-4.109a3 3 0 0 0-4.109 4.109" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserBlockSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.25 7.5a3.75 3.75 0 1 1 7.5 0a3.75 3.75 0 0 1-7.5 0m-4 9.5A3.75 3.75 0 0 1 6 13.25h.34c.185 0 .369.03.544.086l.866.283a7.251 7.251 0 0 0 2.967.323a.22.22 0 0 1 .223.3a5.983 5.983 0 0 0-.44 2.258c0 1.252.384 2.415 1.04 3.377c.091.134.003.323-.16.328a40.077 40.077 0 0 1-7.84-.5a1.537 1.537 0 0 1-1.29-1.517z"/><path fill="currentColor" fill-rule="evenodd" d="M12 16.5c0 .972.308 1.872.832 2.607A4.48 4.48 0 0 0 16.5 21a4.5 4.5 0 1 0-4.5-4.5m4.5 3a2.985 2.985 0 0 1-1.524-.415l4.109-4.109A3 3 0 0 1 16.5 19.5m-2.585-1.476l4.109-4.109a3 3 0 0 0-4.109 4.109" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserClockOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 3.25a4.25 4.25 0 1 0 0 8.5a4.25 4.25 0 0 0 0-8.5M7.25 7.5a2.75 2.75 0 1 1 5.5 0a2.75 2.75 0 0 1-5.5 0" clip-rule="evenodd"/><path fill="currentColor" d="M3.75 17A2.25 2.25 0 0 1 6 14.75h.34c.027 0 .053.004.078.012l.866.283c.68.222 1.38.358 2.084.41a.302.302 0 0 0 .312-.239c.076-.327.174-.645.294-.952a.214.214 0 0 0-.192-.29a7.251 7.251 0 0 1-2.032-.355l-.866-.283a1.752 1.752 0 0 0-.543-.086H6A3.75 3.75 0 0 0 2.25 17v1.188c0 .754.546 1.396 1.29 1.517c2.157.353 4.337.527 6.517.524c.152 0 .247-.165.18-.3a6.958 6.958 0 0 1-.38-.914a.418.418 0 0 0-.388-.29a38.603 38.603 0 0 1-5.688-.5a.037.037 0 0 1-.031-.037zm13.5-1.7a.75.75 0 1 0-1.5 0v1.773c0 .24.115.465.309.606l1 .728a.75.75 0 1 0 .882-1.214l-.691-.502z"/><path fill="currentColor" fill-rule="evenodd" d="M16.5 22.3a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11m0-1.5a4 4 0 1 0 0-8a4 4 0 0 0 0 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserClockSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 3.75a3.75 3.75 0 1 0 0 7.5a3.75 3.75 0 0 0 0-7.5m-4 9.5A3.75 3.75 0 0 0 2.25 17v1.188c0 .754.546 1.396 1.29 1.517c2.21.361 4.443.536 6.676.524c.156-.001.25-.174.174-.31A6.968 6.968 0 0 1 9.5 16.5c0-.787.13-1.544.37-2.25a.213.213 0 0 0-.192-.28a7.247 7.247 0 0 1-1.928-.35l-.866-.284a1.752 1.752 0 0 0-.543-.086zM17.25 15a.75.75 0 0 0-1.5 0v1.773c0 .24.115.465.309.606l1 .728a.75.75 0 0 0 .882-1.214l-.691-.502z"/><path fill="currentColor" fill-rule="evenodd" d="M16.5 22a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11m0-1.5a4 4 0 1 0 0-8a4 4 0 0 0 0 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.75 7.5a4.25 4.25 0 1 1 8.5 0a4.25 4.25 0 0 1-8.5 0M12 4.75a2.75 2.75 0 1 0 0 5.5a2.75 2.75 0 0 0 0-5.5m-4 10A2.25 2.25 0 0 0 5.75 17v1.188c0 .018.013.034.031.037c4.119.672 8.32.672 12.438 0a.037.037 0 0 0 .031-.037V17A2.25 2.25 0 0 0 16 14.75h-.34a.253.253 0 0 0-.079.012l-.865.283a8.751 8.751 0 0 1-5.432 0l-.866-.283a.252.252 0 0 0-.077-.012zM4.25 17A3.75 3.75 0 0 1 8 13.25h.34c.185 0 .369.03.544.086l.866.283a7.251 7.251 0 0 0 4.5 0l.866-.283c.175-.057.359-.086.543-.086H16A3.75 3.75 0 0 1 19.75 17v1.188c0 .754-.546 1.396-1.29 1.517a40.095 40.095 0 0 1-12.92 0a1.537 1.537 0 0 1-1.29-1.517z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserPlusOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.75 7.5a4.25 4.25 0 1 1 8.5 0a4.25 4.25 0 0 1-8.5 0M11 4.75a2.75 2.75 0 1 0 0 5.5a2.75 2.75 0 0 0 0-5.5M3.25 17A3.75 3.75 0 0 1 7 13.25h.34c.185 0 .369.03.544.086l.866.283a7.251 7.251 0 0 0 4.5 0l.866-.283c.175-.057.359-.086.543-.086H15A3.75 3.75 0 0 1 18.75 17v1.188c0 .754-.546 1.396-1.29 1.517a40.095 40.095 0 0 1-12.92 0a1.537 1.537 0 0 1-1.29-1.517zM7 14.75A2.25 2.25 0 0 0 4.75 17v1.188c0 .018.013.034.031.037c4.119.672 8.32.672 12.438 0a.037.037 0 0 0 .031-.037V17A2.25 2.25 0 0 0 15 14.75h-.34a.253.253 0 0 0-.079.012l-.865.283a8.751 8.751 0 0 1-5.432 0l-.866-.283a.252.252 0 0 0-.077-.012z" clip-rule="evenodd"/><path fill="currentColor" d="M19.5 6.25a.75.75 0 0 1 .75.75v1.75H22a.75.75 0 0 1 0 1.5h-1.75V12a.75.75 0 0 1-1.5 0v-1.75H17a.75.75 0 0 1 0-1.5h1.75V7a.75.75 0 0 1 .75-.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserPlusSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 3.75a3.75 3.75 0 1 0 0 7.5a3.75 3.75 0 0 0 0-7.5m-4 9.5A3.75 3.75 0 0 0 3.25 17v1.188c0 .754.546 1.396 1.29 1.517c4.278.699 8.642.699 12.92 0a1.537 1.537 0 0 0 1.29-1.517V17A3.75 3.75 0 0 0 15 13.25h-.34c-.185 0-.369.03-.544.086l-.866.283a7.251 7.251 0 0 1-4.5 0l-.866-.283a1.752 1.752 0 0 0-.543-.086zm12.5-7a.75.75 0 0 1 .75.75v1.75H22a.75.75 0 0 1 0 1.5h-1.75V12a.75.75 0 0 1-1.5 0v-1.75H17a.75.75 0 0 1 0-1.5h1.75V7a.75.75 0 0 1 .75-.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3.75a3.75 3.75 0 1 0 0 7.5a3.75 3.75 0 0 0 0-7.5m-4 9.5A3.75 3.75 0 0 0 4.25 17v1.188c0 .754.546 1.396 1.29 1.517c4.278.699 8.642.699 12.92 0a1.537 1.537 0 0 0 1.29-1.517V17A3.75 3.75 0 0 0 16 13.25h-.34c-.185 0-.369.03-.544.086l-.866.283a7.251 7.251 0 0 1-4.5 0l-.866-.283a1.752 1.752 0 0 0-.543-.086z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViberOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.965 6.202a.822.822 0 0 0-.537.106h-.014c-.375.22-.713.497-1.001.823c-.24.277-.37.557-.404.827c-.02.16-.006.322.041.475l.018.01c.27.793.622 1.556 1.052 2.274a13.374 13.374 0 0 0 2.03 2.775l.024.034l.038.028l.023.027l.028.024a13.564 13.564 0 0 0 2.782 2.04c1.155.629 1.856.926 2.277 1.05v.006c.123.038.235.055.348.055a1.61 1.61 0 0 0 .964-.414c.325-.288.6-.627.814-1.004v-.007c.201-.38.133-.738-.157-.981A12.126 12.126 0 0 0 14.41 13c-.448-.243-.903-.096-1.087.15l-.393.496c-.202.246-.568.212-.568.212l-.01.006c-2.731-.697-3.46-3.462-3.46-3.462s-.034-.376.219-.568l.492-.396c.236-.192.4-.646.147-1.094a11.807 11.807 0 0 0-1.347-1.88a.748.748 0 0 0-.44-.263zM12.58 5a.5.5 0 0 0 0 1c1.264 0 2.314.413 3.145 1.205c.427.433.76.946.978 1.508c.219.563.319 1.164.293 1.766a.5.5 0 0 0 1 .042a5.359 5.359 0 0 0-.361-2.17a5.442 5.442 0 0 0-1.204-1.854l-.01-.01C15.39 5.502 14.085 5 12.579 5"/><path fill="currentColor" d="M12.545 6.644a.5.5 0 0 0 0 1h.017c.912.065 1.576.369 2.041.868c.477.514.724 1.153.705 1.943a.5.5 0 0 0 1 .023c.024-1.037-.31-1.932-.972-2.646V7.83c-.677-.726-1.606-1.11-2.724-1.185l-.017-.002z"/><path fill="currentColor" d="M12.526 8.319a.5.5 0 1 0-.052.998c.418.022.685.148.853.317c.169.17.295.443.318.87a.5.5 0 1 0 .998-.053c-.032-.6-.22-1.13-.605-1.52c-.387-.39-.914-.58-1.512-.612"/><path fill="currentColor" fill-rule="evenodd" d="M7.067 2.384a22.15 22.15 0 0 1 9.664 0l.339.075a5.155 5.155 0 0 1 3.872 3.763a19.718 19.718 0 0 1 0 9.7a5.155 5.155 0 0 1-3.872 3.763l-.34.075a22.15 22.15 0 0 1-6.077.499L8 22.633a.75.75 0 0 1-1.24-.435l-.439-2.622a5.155 5.155 0 0 1-3.465-3.654a19.717 19.717 0 0 1 0-9.7a5.155 5.155 0 0 1 3.872-3.763zm9.337 1.463a20.65 20.65 0 0 0-9.01 0l-.34.076A3.655 3.655 0 0 0 4.31 6.591a18.217 18.217 0 0 0 0 8.962a3.655 3.655 0 0 0 2.745 2.668l.09.02a.75.75 0 0 1 .576.608l.294 1.758l1.872-1.675a.75.75 0 0 1 .553-.19a20.653 20.653 0 0 0 5.964-.445l.339-.076a3.655 3.655 0 0 0 2.745-2.668c.746-2.94.746-6.021 0-8.962a3.655 3.655 0 0 0-2.745-2.668z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViberSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.676 2.628a21.9 21.9 0 0 0-9.555 0l-.339.075a4.905 4.905 0 0 0-3.684 3.58a19.467 19.467 0 0 0 0 9.577a4.905 4.905 0 0 0 3.444 3.52l.465 2.776a.5.5 0 0 0 .826.29l2.731-2.443a21.898 21.898 0 0 0 6.112-.487l.34-.075a4.905 4.905 0 0 0 3.684-3.58a19.469 19.469 0 0 0 0-9.577a4.905 4.905 0 0 0-3.685-3.58zM7.965 6.202a.822.822 0 0 0-.537.106h-.014c-.375.22-.713.497-1.001.823c-.24.277-.37.557-.404.827c-.02.16-.006.322.041.475l.018.01c.27.793.622 1.556 1.052 2.274a13.374 13.374 0 0 0 2.03 2.775l.024.034l.038.028l.023.027l.028.024a13.564 13.564 0 0 0 2.782 2.04c1.155.629 1.856.926 2.277 1.05v.006c.123.038.235.055.348.055a1.61 1.61 0 0 0 .964-.414c.325-.288.6-.627.814-1.004v-.007c.201-.38.133-.738-.157-.981A12.126 12.126 0 0 0 14.41 13c-.448-.243-.903-.096-1.087.15l-.393.496c-.202.246-.568.212-.568.212l-.01.006c-2.731-.697-3.46-3.462-3.46-3.462s-.034-.376.219-.568l.492-.396c.236-.192.4-.646.147-1.094a11.807 11.807 0 0 0-1.347-1.88a.748.748 0 0 0-.44-.263zM12.579 5a.5.5 0 0 0 0 1c1.265 0 2.315.413 3.146 1.205c.427.433.76.946.978 1.508c.219.563.319 1.164.293 1.766a.5.5 0 0 0 1 .042a5.359 5.359 0 0 0-.361-2.17a5.442 5.442 0 0 0-1.204-1.854l-.01-.01C15.39 5.502 14.085 5 12.579 5m-.034 1.644a.5.5 0 0 0 0 1h.017c.912.065 1.576.369 2.041.868c.477.514.724 1.153.705 1.943a.5.5 0 0 0 1 .023c.024-1.037-.31-1.932-.972-2.646V7.83c-.677-.726-1.606-1.11-2.724-1.185l-.017-.002zm-.019 1.675a.5.5 0 1 0-.052.998c.418.022.685.148.853.317c.169.17.295.443.318.87a.5.5 0 1 0 .998-.053c-.032-.6-.22-1.13-.605-1.52c-.387-.39-.914-.58-1.512-.612" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.062 5.62a36.776 36.776 0 0 0-7.118-.055l-1.582.141A2.835 2.835 0 0 0 2.81 8.12a26.819 26.819 0 0 0 0 7.76a2.836 2.836 0 0 0 2.553 2.414l1.582.14a36.77 36.77 0 0 0 7.118-.054l.608-.064a2.845 2.845 0 0 0 2.493-2.272l3.047 1.618a.75.75 0 0 0 1.099-.597l.025-.284a54.757 54.757 0 0 0 0-9.563l-.025-.284a.75.75 0 0 0-1.1-.596l-3.046 1.618a2.845 2.845 0 0 0-2.493-2.272zM7.078 7.059a35.275 35.275 0 0 1 6.827.053l.608.064a1.346 1.346 0 0 1 1.19 1.143c.358 2.441.358 4.92 0 7.362a1.346 1.346 0 0 1-1.19 1.143l-.608.064a35.272 35.272 0 0 1-6.827.053l-1.582-.142a1.336 1.336 0 0 1-1.203-1.136a25.319 25.319 0 0 1 0-7.326A1.335 1.335 0 0 1 5.496 7.2zM17.36 9.55c.149 1.63.149 3.27 0 4.9l2.547 1.353a53.245 53.245 0 0 0 0-7.606z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.966 5.814c2.353-.21 4.72-.192 7.07.055l.608.064A2.595 2.595 0 0 1 16.93 8.08l3.279-1.742a.75.75 0 0 1 1.099.597l.025.283a54.758 54.758 0 0 1 0 9.564l-.025.284a.75.75 0 0 1-1.1.596l-3.278-1.741a2.596 2.596 0 0 1-2.287 2.146l-.608.064c-2.35.247-4.717.265-7.07.055l-1.581-.141a2.585 2.585 0 0 1-2.328-2.201a26.568 26.568 0 0 1 0-7.687a2.585 2.585 0 0 1 2.328-2.201zm10.15 8.553a.75.75 0 0 1 .236.078l2.555 1.358a53.246 53.246 0 0 0 0-7.606l-2.555 1.358a.749.749 0 0 1-.236.079a26.72 26.72 0 0 1 0 4.733" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VkOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M23.405 16.865a11.394 11.394 0 0 0-2.123-2.964a15.527 15.527 0 0 0-.905-.94l-.038-.037a18.573 18.573 0 0 1-.346-.344a29.72 29.72 0 0 0 3.01-5.238l.033-.074l.023-.078c.109-.363.233-1.053-.207-1.677c-.456-.644-1.185-.76-1.674-.76h-2.247a2.232 2.232 0 0 0-2.19 1.442a17.506 17.506 0 0 1-1.806 3.271V6.833c0-.34-.032-.91-.397-1.39c-.436-.576-1.067-.69-1.505-.69H9.467a1.813 1.813 0 0 0-1.85 1.682l-.003.045v.045c0 .485.192.843.346 1.068c.069.101.142.193.189.251l.01.013c.05.062.083.103.116.149c.088.118.213.302.249.776v1.473a19.785 19.785 0 0 1-1.751-3.792l-.008-.022l-.008-.021c-.122-.319-.317-.783-.708-1.137c-.456-.415-.996-.53-1.487-.53h-2.28c-.497 0-1.096.116-1.543.587C.3 5.804.25 6.36.25 6.654v.134l.028.13a19.43 19.43 0 0 0 3.801 8.02a10.151 10.151 0 0 0 7.893 4.692l.041.003h.042c.726 0 1.483-.063 2.052-.442c.767-.512.828-1.297.828-1.689v-1.138c.197.16.441.374.74.662c.362.362.65.676.897.95l.132.146c.192.214.381.425.553.598c.216.217.483.456.817.633c.363.191.744.278 1.148.278h2.281c.481 0 1.17-.114 1.655-.676c.528-.612.488-1.363.322-1.902l-.03-.097zm-5.72.106a26.137 26.137 0 0 0-.957-1.014l-.003-.003c-1.357-1.308-1.99-1.535-2.438-1.535c-.239 0-.502.026-.673.24a.774.774 0 0 0-.147.348a2.452 2.452 0 0 0-.032.444v2.051c0 .255-.042.362-.16.44c-.157.105-.492.19-1.211.19a8.65 8.65 0 0 1-6.752-4.052l-.008-.013l-.01-.012A17.93 17.93 0 0 1 1.75 6.63c.004-.13.032-.209.078-.257c.047-.05.162-.12.454-.12h2.28c.253 0 .385.056.48.141c.106.098.2.263.312.557c.56 1.646 1.316 3.186 2.033 4.318c.358.565.71 1.036 1.028 1.369c.159.166.314.304.463.402c.143.094.306.169.474.169c.088 0 .191-.01.29-.053a.523.523 0 0 0 .25-.232c.103-.188.132-.465.132-.828V8.723c-.053-.818-.3-1.279-.54-1.606a5.66 5.66 0 0 0-.15-.193l-.013-.016a2.694 2.694 0 0 1-.122-.16a.39.39 0 0 1-.085-.214a.313.313 0 0 1 .324-.282h3.595c.206 0 .275.05.31.097c.05.065.092.2.092.484v4.528c0 .538.248.902.608.902c.414 0 .713-.251 1.235-.773l.009-.01l.008-.009a19.01 19.01 0 0 0 2.84-4.722l.004-.012a.732.732 0 0 1 .74-.485h2.3c.312 0 .414.08.448.127c.035.05.06.157 0 .367a28.216 28.216 0 0 1-3.029 5.2l-.008.012c-.115.177-.242.373-.26.597c-.02.242.084.461.267.697c.133.196.408.465.687.738l.026.026c.292.286.609.596.863.896l.007.007l.007.008a9.892 9.892 0 0 1 1.865 2.586c.076.26.03.394-.03.463c-.069.08-.224.155-.518.155h-2.282a.91.91 0 0 1-.447-.105a1.96 1.96 0 0 1-.454-.364c-.143-.143-.298-.316-.488-.527z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VkSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.504 18.381h-2.282c-.863 0-1.123-.699-2.67-2.247c-1.352-1.303-1.923-1.465-2.265-1.465c-.473 0-.602.13-.602.781v2.052c0 .555-.18.88-1.63.88a8.9 8.9 0 0 1-6.955-4.17a18.18 18.18 0 0 1-3.6-7.558c0-.342.13-.652.782-.652h2.28c.586 0 .797.261 1.027.864c1.107 3.258 2.996 6.092 3.763 6.092c.294 0 .422-.13.422-.863V8.739c-.097-1.531-.91-1.66-.91-2.214a.563.563 0 0 1 .585-.523h3.584c.49 0 .652.244.652.83v4.53c0 .489.21.651.358.651c.294 0 .52-.162 1.059-.7a18.76 18.76 0 0 0 2.802-4.66a.982.982 0 0 1 .993-.65h2.281c.684 0 .829.342.684.83a28.466 28.466 0 0 1-3.062 5.262c-.246.375-.344.57 0 1.01c.226.342 1.026 1.01 1.563 1.645c.782.78 1.431 1.682 1.922 2.67c.196.636-.131.961-.782.961"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeDownOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.967 5.69c-.17-1.359-1.776-1.992-2.828-1.115L8.798 8.192a.25.25 0 0 1-.16.058H5A1.75 1.75 0 0 0 3.25 10v4c0 .966.784 1.75 1.75 1.75h3.638a.25.25 0 0 1 .16.058l4.34 3.617c1.053.877 2.66.244 2.83-1.116c.523-4.19.523-8.428 0-12.618m-1.868.037a.233.233 0 0 1 .38.15c.508 4.066.508 8.18 0 12.246a.233.233 0 0 1-.38.15l-4.34-3.617a1.75 1.75 0 0 0-1.121-.406H5a.25.25 0 0 1-.25-.25v-4A.25.25 0 0 1 5 9.75h3.638c.41 0 .806-.144 1.12-.406z" clip-rule="evenodd"/><path fill="currentColor" d="M19.032 7.948a.75.75 0 1 0-1.408.517c.405 1.101.626 2.291.626 3.535a10.25 10.25 0 0 1-.626 3.535a.75.75 0 0 0 1.408.517A11.667 11.667 0 0 0 19.75 12c0-1.423-.253-2.788-.718-4.052"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeDownSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.977 4.201a1 1 0 0 1 1.633.645l.095.766c.53 4.242.53 8.534 0 12.776l-.095.766a1 1 0 0 1-1.633.644l-5.019-4.182a.5.5 0 0 0-.32-.116H4.75c-.69 0-1.25-.56-1.25-1.25v-4.5c0-.69.56-1.25 1.25-1.25h3.888a.5.5 0 0 0 .32-.116zm5.055 3.747a.75.75 0 0 0-1.408.517c.405 1.1.626 2.291.626 3.535a10.25 10.25 0 0 1-.626 3.535a.75.75 0 0 0 1.408.517A11.667 11.667 0 0 0 19.75 12c0-1.423-.253-2.788-.718-4.052"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeOffOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.139 4.575c1.052-.877 2.658-.244 2.828 1.116c.524 4.19.524 8.428 0 12.618c-.17 1.36-1.776 1.993-2.828 1.116l-4.341-3.617a.25.25 0 0 0-.16-.058H4A1.75 1.75 0 0 1 2.25 14v-4c0-.966.784-1.75 1.75-1.75h3.638a.25.25 0 0 0 .16-.058zm1.34 1.302a.233.233 0 0 0-.38-.15l-4.34 3.617a1.75 1.75 0 0 1-1.121.406H4a.25.25 0 0 0-.25.25v4c0 .138.112.25.25.25h3.638c.41 0 .806.143 1.12.406l4.341 3.617a.233.233 0 0 0 .38-.15c.508-4.066.508-8.18 0-12.246" clip-rule="evenodd"/><path fill="currentColor" d="M17.202 9.702a.75.75 0 0 1 1.06 0l1.238 1.237l1.237-1.237a.75.75 0 0 1 1.061 1.06L20.561 12l1.237 1.237a.75.75 0 0 1-1.06 1.061L19.5 13.061l-1.237 1.237a.75.75 0 0 1-1.061-1.06L18.439 12l-1.237-1.237a.75.75 0 0 1 0-1.061"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeOffSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.61 4.846a1 1 0 0 0-1.633-.645L7.958 8.384a.5.5 0 0 1-.32.116H3.75c-.69 0-1.25.56-1.25 1.25v4.5c0 .69.56 1.25 1.25 1.25h3.888a.5.5 0 0 1 .32.116l5.02 4.182a1 1 0 0 0 1.632-.644l.095-.766c.53-4.242.53-8.534 0-12.776zm2.627 4.856a.75.75 0 0 1 1.061 0l1.238 1.237l1.237-1.237a.75.75 0 1 1 1.06 1.06L20.597 12l1.238 1.237a.75.75 0 1 1-1.061 1.061l-1.238-1.237l-1.237 1.237a.75.75 0 0 1-1.06-1.06L18.474 12l-1.238-1.238a.75.75 0 0 1 0-1.06"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeUpOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.139 4.575c1.052-.877 2.658-.244 2.828 1.116c.524 4.19.524 8.428 0 12.618c-.17 1.36-1.776 1.993-2.828 1.116l-4.341-3.617a.25.25 0 0 0-.16-.058H4A1.75 1.75 0 0 1 2.25 14v-4c0-.966.784-1.75 1.75-1.75h3.638a.25.25 0 0 0 .16-.058zm1.34 1.302a.233.233 0 0 0-.38-.15l-4.34 3.617a1.75 1.75 0 0 1-1.121.406H4a.25.25 0 0 0-.25.25v4c0 .138.112.25.25.25h3.638c.41 0 .806.143 1.12.406l4.341 3.617a.233.233 0 0 0 .38-.15c.508-4.066.508-8.18 0-12.246" clip-rule="evenodd"/><path fill="currentColor" d="M19.908 5.946a.75.75 0 0 1 .948.476c.58 1.755.894 3.63.894 5.578a17.75 17.75 0 0 1-.894 5.578a.75.75 0 1 1-1.424-.471A16.26 16.26 0 0 0 20.25 12a16.26 16.26 0 0 0-.818-5.107a.75.75 0 0 1 .476-.947m-1.876 2.002a.75.75 0 0 0-1.408.517c.405 1.101.626 2.291.626 3.535a10.25 10.25 0 0 1-.626 3.535a.75.75 0 0 0 1.408.517A11.667 11.667 0 0 0 18.75 12c0-1.423-.253-2.788-.718-4.052"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeUpSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.61 4.846a1 1 0 0 0-1.633-.645L7.958 8.384a.5.5 0 0 1-.32.116H3.75c-.69 0-1.25.56-1.25 1.25v4.5c0 .69.56 1.25 1.25 1.25h3.888a.5.5 0 0 1 .32.116l5.02 4.182a1 1 0 0 0 1.632-.644l.095-.766c.53-4.242.53-8.534 0-12.776zm5.298 1.1a.75.75 0 0 1 .948.476c.58 1.755.894 3.63.894 5.578a17.75 17.75 0 0 1-.894 5.578a.75.75 0 1 1-1.424-.471A16.26 16.26 0 0 0 20.25 12a16.26 16.26 0 0 0-.818-5.107a.75.75 0 0 1 .476-.947m-1.876 2.002a.75.75 0 0 0-1.408.517c.405 1.1.626 2.291.626 3.535a10.25 10.25 0 0 1-.626 3.535a.75.75 0 0 0 1.408.517A11.671 11.671 0 0 0 18.75 12c0-1.423-.253-2.788-.718-4.052"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WalletOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 12a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0"/><path fill="currentColor" fill-rule="evenodd" d="M20.441 6.674a4.71 4.71 0 0 0-3.84-2.839l-.652-.068a44.4 44.4 0 0 0-9.9.068l-.432.051a3.681 3.681 0 0 0-3.214 3.169a37.391 37.391 0 0 0 0 9.89a3.681 3.681 0 0 0 3.214 3.169l.432.051c3.287.392 6.607.415 9.9.068l.652-.068a4.71 4.71 0 0 0 3.84-2.839a2.77 2.77 0 0 0 1.963-2.324c.233-1.994.233-4.01 0-6.004a2.77 2.77 0 0 0-1.963-2.324m-4.65-1.416a42.9 42.9 0 0 0-9.565.066l-.432.052A2.181 2.181 0 0 0 3.89 7.253a35.89 35.89 0 0 0 0 9.494a2.18 2.18 0 0 0 1.904 1.877l.432.052c3.176.378 6.385.4 9.566.066l.652-.069a3.207 3.207 0 0 0 2.124-1.131a26.874 26.874 0 0 1-4.526-.118a2.772 2.772 0 0 1-2.446-2.422a25.846 25.846 0 0 1 0-6.004a2.772 2.772 0 0 1 2.446-2.422a26.864 26.864 0 0 1 4.526-.118a3.207 3.207 0 0 0-2.124-1.131zm3.486 2.757l.002.011l.006.04l.199-.032c.103.01.205.021.308.033c.587.065 1.055.53 1.122 1.105c.22 1.879.22 3.777 0 5.656a1.272 1.272 0 0 1-1.122 1.105c-.103.012-.205.023-.308.033l-.199-.031l-.006.039l-.002.011c-1.678.152-3.4.135-5.069-.052a1.272 1.272 0 0 1-1.122-1.105a24.341 24.341 0 0 1 0-5.656a1.272 1.272 0 0 1 1.122-1.105a25.371 25.371 0 0 1 5.07-.052" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WalletSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m21.01 10.171l.003 3.623c-.026.345-.059.69-.099 1.034a1.272 1.272 0 0 1-1.122 1.105c-1.84.206-3.744.206-5.584 0a1.272 1.272 0 0 1-1.122-1.105a24.341 24.341 0 0 1 0-5.656a1.272 1.272 0 0 1 1.122-1.105a25.345 25.345 0 0 1 5.584 0c.587.065 1.055.53 1.122 1.105c.039.332.071.666.096 1M17 10.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3" clip-rule="evenodd"/><path fill="currentColor" d="M20.404 6.04c.155.269-.137.57-.446.536a26.846 26.846 0 0 0-5.916 0a2.772 2.772 0 0 0-2.446 2.422a25.846 25.846 0 0 0 0 6.004a2.772 2.772 0 0 0 2.446 2.422a26.84 26.84 0 0 0 5.916 0c.311-.035.606.269.449.54a4.971 4.971 0 0 1-3.78 2.45l-.652.068a44.67 44.67 0 0 1-9.956-.069l-.432-.051a3.931 3.931 0 0 1-3.432-3.384a37.64 37.64 0 0 1 0-9.956a3.931 3.931 0 0 1 3.432-3.384l.432-.051a44.65 44.65 0 0 1 9.956-.069l.652.069a4.96 4.96 0 0 1 3.777 2.453"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WatchOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.566 6.552a.747.747 0 0 0 .057.31A6.735 6.735 0 0 0 5.25 12c0 2.058.92 3.9 2.373 5.138c-.04.095-.06.2-.057.31l.088 2.887c.02.672.466 1.257 1.109 1.455a11.02 11.02 0 0 0 6.474 0a1.571 1.571 0 0 0 1.11-1.455l.087-2.887a.748.748 0 0 0-.056-.31a6.757 6.757 0 0 0 2.069-3.132a.75.75 0 0 0 1.303-.506v-3a.75.75 0 0 0-1.303-.506a6.757 6.757 0 0 0-2.07-3.132c.04-.095.06-.2.057-.31l-.088-2.887a1.571 1.571 0 0 0-1.109-1.455a11.02 11.02 0 0 0-6.474 0a1.572 1.572 0 0 0-1.11 1.455zm7.23-2.908a9.52 9.52 0 0 0-5.593 0a.072.072 0 0 0-.05.066l-.067 2.2A6.723 6.723 0 0 1 12 5.25c1.043 0 2.032.237 2.914.66l-.067-2.2a.071.071 0 0 0-.05-.066M9.154 20.29l-.067-2.2a6.72 6.72 0 0 0 2.914.66a6.723 6.723 0 0 0 2.914-.66l-.067 2.2a.07.07 0 0 1-.05.066a9.52 9.52 0 0 1-5.594 0a.072.072 0 0 1-.05-.066M12 6.75a5.25 5.25 0 1 0 0 10.5a5.25 5.25 0 0 0 0-10.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WatchSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7.573 7.24l.084-3.566a1.57 1.57 0 0 1 1.11-1.465a11.008 11.008 0 0 1 6.467 0a1.57 1.57 0 0 1 1.109 1.465l.084 3.566a6.498 6.498 0 0 1 1.85 3.063a.75.75 0 0 1 1.473.197v3a.75.75 0 0 1-1.474.197a6.498 6.498 0 0 1-1.85 3.063l-.083 3.566a1.57 1.57 0 0 1-1.11 1.465c-2.106.647-4.36.647-6.467 0a1.57 1.57 0 0 1-1.109-1.465l-.084-3.566A6.482 6.482 0 0 1 5.5 12c0-1.88.798-3.573 2.073-4.76m1.634-3.597a9.508 9.508 0 0 1 5.586 0a.07.07 0 0 1 .05.066l.058 2.473A6.473 6.473 0 0 0 12 5.5a6.473 6.473 0 0 0-2.901.682l.058-2.473c0-.03.02-.057.05-.066m-.108 14.175l.058 2.473c0 .03.02.057.05.066c1.82.56 3.766.56 5.586 0a.07.07 0 0 0 .05-.066l.058-2.473A6.473 6.473 0 0 1 12 18.5a6.473 6.473 0 0 1-2.901-.682m.603-3.52a.75.75 0 0 0-1.06 1.06a4.75 4.75 0 0 0 6.717 0a.75.75 0 1 0-1.06-1.06a3.25 3.25 0 0 1-4.597 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WechatSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8.452 3.5C4.149 3.5.5 6.425.5 10.214c0 1.94.97 3.668 2.484 4.878v1.682a.5.5 0 0 0 .634.482l2.363-.657a9.307 9.307 0 0 0 2.63.329a.202.202 0 0 0 .192-.24a6.457 6.457 0 0 1-.097-1.122c0-3.898 3.405-6.637 7.197-7.01a.198.198 0 0 0 .172-.26C15.088 5.48 12 3.5 8.452 3.5M6.855 7.46a1.065 1.065 0 1 1-2.13 0a1.065 1.065 0 0 1 2.13 0m4.258 1.064a1.065 1.065 0 1 0 0-2.129a1.065 1.065 0 0 0 0 2.13"/><path d="M16.79 9.887c3.617 0 6.71 2.461 6.71 5.679c0 1.632-.81 3.084-2.07 4.104v1.362a.5.5 0 0 1-.634.482l-1.947-.541a7.845 7.845 0 0 1-2.059.271c-3.616 0-6.71-2.46-6.71-5.678s3.094-5.679 6.71-5.679m-2.244 4.758a.932.932 0 1 0 0-1.863a.932.932 0 0 0 0 1.863m5.456-.931a.931.931 0 1 1-1.863 0a.931.931 0 0 1 1.863 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WhatsappOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.886 7.17c.183.005.386.015.579.443c.128.285.343.81.519 1.238c.137.333.249.607.277.663c.064.128.104.275.02.448l-.028.058a1.43 1.43 0 0 1-.23.37a9.386 9.386 0 0 0-.143.17c-.085.104-.17.206-.242.278c-.129.128-.262.266-.114.522c.149.256.668 1.098 1.435 1.777a6.634 6.634 0 0 0 1.903 1.2c.07.03.127.055.17.076c.257.128.41.108.558-.064c.149-.173.643-.749.817-1.005c.168-.256.34-.216.578-.128c.238.089 1.504.71 1.761.837l.143.07c.179.085.3.144.352.23c.064.109.064.62-.148 1.222c-.218.6-1.267 1.176-1.742 1.22l-.135.016c-.436.052-.988.12-2.956-.655c-2.426-.954-4.027-3.32-4.35-3.799a2.768 2.768 0 0 0-.053-.076l-.006-.008c-.147-.197-1.048-1.402-1.048-2.646c0-1.19.587-1.81.854-2.092l.047-.05a.95.95 0 0 1 .687-.32c.173 0 .347 0 .495.005"/><path fill="currentColor" fill-rule="evenodd" d="M2.184 21.331a.4.4 0 0 0 .487.494l4.607-1.204a10 10 0 0 0 4.76 1.207h.004c5.486 0 9.958-4.446 9.958-9.912a9.828 9.828 0 0 0-2.914-7.011A9.917 9.917 0 0 0 12.042 2c-5.486 0-9.958 4.446-9.958 9.911c0 1.739.458 3.447 1.33 4.954zm2.677-4.068a1.5 1.5 0 0 0-.148-1.15a8.377 8.377 0 0 1-1.129-4.202c0-4.63 3.793-8.411 8.458-8.411c2.27 0 4.388.877 5.986 2.468a8.328 8.328 0 0 1 2.472 5.948c0 4.63-3.793 8.412-8.458 8.412h-.005a8.5 8.5 0 0 1-4.044-1.026a1.5 1.5 0 0 0-1.094-.132l-2.762.721z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WhatsappSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" stroke="currentColor" stroke-linejoin="round" d="M9.92 7.408a53 53 0 0 1 .523 1.245l.003.008a105.648 105.648 0 0 0 .239.575l.022.052l.001.001c.1.2.202.523.022.892zm0 0c-.12-.268-.277-.48-.498-.608m.499.608L9.422 6.8m3.832 7.165a4.013 4.013 0 0 0-.185-.083l-.009-.004a6.139 6.139 0 0 1-1.769-1.114c-.647-.574-1.11-1.28-1.29-1.581c.086-.088.18-.201.255-.293l.01-.01c.049-.06.093-.114.135-.162c.158-.18.23-.33.298-.473l.004-.008l.027-.056l-.776.93l-.028.016l-.402.233m3.73 2.605l-.01.018l.015-.016m-.005-.002l-.01.019l-.033.038c-.012.015-.012.011.002.002l.002-.001a.224.224 0 0 1 .014-.008l-.008.016l-.19.382m.223-.448l.005.002m0 0a.412.412 0 0 0 .072.03s-.014-.003-.038 0a.203.203 0 0 0-.064.018l-.008.016l-.19.382m.228-.446c.189-.22.591-.693.732-.9m-.96 1.346c.258.128.41.108.56-.064c.148-.173.642-.749.816-1.005m-1.376 1.069a3.446 3.446 0 0 0-.169-.076a6.634 6.634 0 0 1-1.903-1.2c-.767-.68-1.286-1.521-1.435-1.777m4.467 1.706l-.002.004l.418.274m-.416-.278l.002-.002l.414.28m-.416-.278c.138-.208.338-.378.62-.412c.227-.026.436.051.545.092l.004.002c.146.054.55.245.915.421c.38.183.763.371.894.436c.047.024.092.045.136.066l-.216.451m-2.482-.778c.168-.256.34-.216.578-.128c.238.089 1.504.71 1.761.837l.143.07M9.524 11.36l.401-.233l.03-.017l.069-.72a9.386 9.386 0 0 0-.144.17c-.085.104-.17.206-.242.278c-.129.128-.262.266-.114.522Zm7.365 2.762l.216-.45l.012.005c.08.038.175.084.256.134a.86.86 0 0 1 .298.287m-.782.024c.179.086.3.145.352.232m.43-.256l-.578 1.477c.212-.601.212-1.113.148-1.221m.43-.256c.08.134.096.289.102.368c.009.106.007.227-.004.356a3.766 3.766 0 0 1-.205.92l-.001.003c-.167.46-.602.83-.974 1.07c-.377.242-.838.445-1.19.478m2.272-3.195l-.43.255m0 0l-1.843 2.94m0 0a4.686 4.686 0 0 0-.122.014l-.06-.496l.06.496c-.24.03-.54.064-1.046-.02c-.487-.082-1.154-.273-2.154-.666c-2.57-1.011-4.241-3.482-4.575-3.975l-.006-.009l.415-.28l-.415.28l-.03-.045l-.009-.012v-.002l-.005-.005a6.744 6.744 0 0 1-.625-1.005c-.26-.513-.523-1.207-.523-1.941c0-1.385.697-2.125.99-2.436l.044-.046zM9.422 6.8c-.21-.122-.418-.127-.512-.13m.512.13l-.512-.13m0 0H8.9zM7.516 20.182a.5.5 0 0 0-.364-.044l-4.44 1.16l1.185-4.3a.5.5 0 0 0-.05-.384a9.375 9.375 0 0 1-1.263-4.703c0-5.186 4.246-9.411 9.458-9.411c2.534 0 4.906.982 6.692 2.76a9.327 9.327 0 0 1 2.766 6.656c0 5.187-4.246 9.412-9.458 9.412h-.005a9.499 9.499 0 0 1-4.52-1.146Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowsOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M21.094 12.756a.2.2 0 0 0-.064-.01h-8.554a.2.2 0 0 0-.2.2v6.421a.2.2 0 0 0 .168.198l8.553 1.392a.2.2 0 0 0 .233-.197v-7.815a.2.2 0 0 0-.136-.19M19.73 19.23v-4.986h-5.953v4.017zm-7.318-7.985c.02.007.042.011.064.011h8.554a.2.2 0 0 0 .2-.2V3.238a.2.2 0 0 0-.233-.197l-8.553 1.417a.2.2 0 0 0-.168.197v6.402a.2.2 0 0 0 .136.19m7.318-6.475l-5.953.987v3.999h5.953zm-9.49 6.475a.2.2 0 0 0 .143-.191V4.968a.2.2 0 0 0-.233-.197L2.982 5.958a.2.2 0 0 0-.167.197v4.9a.2.2 0 0 0 .2.2h7.168a.2.2 0 0 0 .057-.009M8.883 6.501l-4.568.757v2.497h4.568zm1.357 6.253a.2.2 0 0 0-.057-.009H3.015a.2.2 0 0 0-.2.2v4.946a.2.2 0 0 0 .168.197l7.168 1.168a.2.2 0 0 0 .232-.198v-6.113a.2.2 0 0 0-.143-.191M8.883 17.53v-3.285H4.315v2.54z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowsSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M13.027 10.507V5.122l7.453-1.235v6.62zm7.453 9.606L13.027 18.9v-5.405h7.453zM9.633 10.505H3.565V6.622l6.068-1.005zm0 7.907l-6.068-.989v-3.928h6.068z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YoutubeOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M10.386 8.357A.75.75 0 0 0 9.25 9v6a.75.75 0 0 0 1.136.643l5-3a.75.75 0 0 0 0-1.286zM13.542 12l-2.792 1.675v-3.35z"/><path d="M17.03 4.641a64.499 64.499 0 0 0-10.06 0l-2.241.176a2.975 2.975 0 0 0-2.703 2.475a28.566 28.566 0 0 0 0 9.416a2.975 2.975 0 0 0 2.703 2.475l2.24.176c3.349.262 6.713.262 10.062 0l2.24-.176a2.975 2.975 0 0 0 2.703-2.475c.52-3.117.52-6.299 0-9.416a2.975 2.975 0 0 0-2.703-2.475zM7.087 6.137a62.998 62.998 0 0 1 9.828 0l2.24.175c.676.053 1.229.56 1.34 1.228a27.066 27.066 0 0 1 0 8.92a1.475 1.475 0 0 1-1.34 1.228l-2.24.175a62.98 62.98 0 0 1-9.828 0l-2.24-.175a1.475 1.475 0 0 1-1.34-1.228a27.066 27.066 0 0 1 0-8.92a1.475 1.475 0 0 1 1.34-1.228z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YoutubeSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.989 4.89a64.248 64.248 0 0 1 10.022 0l2.24.176a2.725 2.725 0 0 1 2.476 2.268c.517 3.09.517 6.243 0 9.332a2.725 2.725 0 0 1-2.475 2.268l-2.24.175a64.24 64.24 0 0 1-10.023 0l-2.24-.175a2.725 2.725 0 0 1-2.476-2.268a28.315 28.315 0 0 1 0-9.332a2.725 2.725 0 0 1 2.475-2.268zM10 14.47V9.53a.3.3 0 0 1 .454-.257l4.117 2.47a.3.3 0 0 1 0 .514l-4.117 2.47A.3.3 0 0 1 10 14.47" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomInOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.2 6.95a.75.75 0 0 1 .75.75v1.75h1.75a.75.75 0 0 1 0 1.5h-1.75v1.75a.75.75 0 0 1-1.5 0v-1.75H7.7a.75.75 0 0 1 0-1.5h1.75V7.7a.75.75 0 0 1 .75-.75"/><path fill="currentColor" fill-rule="evenodd" d="M5.399 14.945a6.75 6.75 0 0 0 8.986.5l5.156 5.156a.75.75 0 1 0 1.06-1.06l-5.155-5.156a6.75 6.75 0 1 0-10.047.56m1.06-8.486a5.25 5.25 0 0 0 7.42 7.43l.005-.005l.005-.005a5.25 5.25 0 0 0-7.43-7.42" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomInSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.385 15.446a6.751 6.751 0 1 1 1.06-1.06l5.156 5.155a.75.75 0 1 1-1.06 1.06zM6.95 10.2a.75.75 0 0 1 .75-.75h1.75V7.7a.75.75 0 0 1 1.5 0v1.75h1.75a.75.75 0 0 1 0 1.5h-1.75v1.75a.75.75 0 0 1-1.5 0v-1.75H7.7a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomOutOutline(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.7 9.45a.75.75 0 0 0 0 1.5h5a.75.75 0 0 0 0-1.5z"/><path fill="currentColor" fill-rule="evenodd" d="M5.399 14.945a6.75 6.75 0 0 0 8.986.5l5.156 5.156a.75.75 0 1 0 1.06-1.06l-5.155-5.156a6.751 6.751 0 1 0-10.047.56m1.06-8.486a5.25 5.25 0 0 0 7.42 7.43l.005-.005l.004-.005A5.25 5.25 0 0 0 6.46 6.46" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomOutSolid(children ...ElementRenderer) *BasilIcon {
	return &BasilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.399 14.945a6.751 6.751 0 0 0 8.986.5l5.156 5.156a.75.75 0 0 0 1.06-1.06l-5.155-5.156a6.75 6.75 0 1 0-10.047.56M7.7 9.45a.75.75 0 0 0 0 1.5h5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
