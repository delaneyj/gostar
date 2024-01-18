package ant_design

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "4.3.1"
	hAttr          = "1em"
	viewbox        = "0 0 1024 1024"
	fill           = "currentColor"
)

type AntDesignIcon struct {
	*SVGSVGElement
}

type AntDesignIconFn func(children ...ElementRenderer) *AntDesignIcon

var IconLookup = map[string]AntDesignIconFn{
	"accountBookFilled":            AccountBookFilled,
	"accountBookOutlined":          AccountBookOutlined,
	"accountBookTwotone":           AccountBookTwotone,
	"aimOutlined":                  AimOutlined,
	"alertFilled":                  AlertFilled,
	"alertOutlined":                AlertOutlined,
	"alertTwotone":                 AlertTwotone,
	"alibabaOutlined":              AlibabaOutlined,
	"alignCenterOutlined":          AlignCenterOutlined,
	"alignLeftOutlined":            AlignLeftOutlined,
	"alignRightOutlined":           AlignRightOutlined,
	"alipayCircleFilled":           AlipayCircleFilled,
	"alipayCircleOutlined":         AlipayCircleOutlined,
	"alipayOutlined":               AlipayOutlined,
	"alipaySquareFilled":           AlipaySquareFilled,
	"aliwangwangFilled":            AliwangwangFilled,
	"aliwangwangOutlined":          AliwangwangOutlined,
	"aliyunOutlined":               AliyunOutlined,
	"amazonCircleFilled":           AmazonCircleFilled,
	"amazonOutlined":               AmazonOutlined,
	"amazonSquareFilled":           AmazonSquareFilled,
	"androidFilled":                AndroidFilled,
	"androidOutlined":              AndroidOutlined,
	"antCloudOutlined":             AntCloudOutlined,
	"antDesignOutlined":            AntDesignOutlined,
	"apartmentOutlined":            ApartmentOutlined,
	"apiFilled":                    ApiFilled,
	"apiOutlined":                  ApiOutlined,
	"apiTwotone":                   ApiTwotone,
	"appleFilled":                  AppleFilled,
	"appleOutlined":                AppleOutlined,
	"appstoreAddOutlined":          AppstoreAddOutlined,
	"appstoreFilled":               AppstoreFilled,
	"appstoreOutlined":             AppstoreOutlined,
	"appstoreTwotone":              AppstoreTwotone,
	"areaChartOutlined":            AreaChartOutlined,
	"arrowDownOutlined":            ArrowDownOutlined,
	"arrowLeftOutlined":            ArrowLeftOutlined,
	"arrowRightOutlined":           ArrowRightOutlined,
	"arrowUpOutlined":              ArrowUpOutlined,
	"arrowsAltOutlined":            ArrowsAltOutlined,
	"audioFilled":                  AudioFilled,
	"audioMutedOutlined":           AudioMutedOutlined,
	"audioOutlined":                AudioOutlined,
	"audioTwotone":                 AudioTwotone,
	"auditOutlined":                AuditOutlined,
	"backwardFilled":               BackwardFilled,
	"backwardOutlined":             BackwardOutlined,
	"bankFilled":                   BankFilled,
	"bankOutlined":                 BankOutlined,
	"bankTwotone":                  BankTwotone,
	"barChartOutlined":             BarChartOutlined,
	"barcodeOutlined":              BarcodeOutlined,
	"barsOutlined":                 BarsOutlined,
	"behanceCircleFilled":          BehanceCircleFilled,
	"behanceOutlined":              BehanceOutlined,
	"behanceSquareFilled":          BehanceSquareFilled,
	"behanceSquareOutlined":        BehanceSquareOutlined,
	"bellFilled":                   BellFilled,
	"bellOutlined":                 BellOutlined,
	"bellTwotone":                  BellTwotone,
	"bgColorsOutlined":             BgColorsOutlined,
	"blockOutlined":                BlockOutlined,
	"boldOutlined":                 BoldOutlined,
	"bookFilled":                   BookFilled,
	"bookOutlined":                 BookOutlined,
	"bookTwotone":                  BookTwotone,
	"borderBottomOutlined":         BorderBottomOutlined,
	"borderHorizontalOutlined":     BorderHorizontalOutlined,
	"borderInnerOutlined":          BorderInnerOutlined,
	"borderLeftOutlined":           BorderLeftOutlined,
	"borderOuterOutlined":          BorderOuterOutlined,
	"borderOutlined":               BorderOutlined,
	"borderRightOutlined":          BorderRightOutlined,
	"borderTopOutlined":            BorderTopOutlined,
	"borderVerticleOutlined":       BorderVerticleOutlined,
	"borderlessTableOutlined":      BorderlessTableOutlined,
	"boxPlotFilled":                BoxPlotFilled,
	"boxPlotOutlined":              BoxPlotOutlined,
	"boxPlotTwotone":               BoxPlotTwotone,
	"branchesOutlined":             BranchesOutlined,
	"bugFilled":                    BugFilled,
	"bugOutlined":                  BugOutlined,
	"bugTwotone":                   BugTwotone,
	"buildFilled":                  BuildFilled,
	"buildOutlined":                BuildOutlined,
	"buildTwotone":                 BuildTwotone,
	"bulbFilled":                   BulbFilled,
	"bulbOutlined":                 BulbOutlined,
	"bulbTwotone":                  BulbTwotone,
	"calculatorFilled":             CalculatorFilled,
	"calculatorOutlined":           CalculatorOutlined,
	"calculatorTwotone":            CalculatorTwotone,
	"calendarFilled":               CalendarFilled,
	"calendarOutlined":             CalendarOutlined,
	"calendarTwotone":              CalendarTwotone,
	"cameraFilled":                 CameraFilled,
	"cameraOutlined":               CameraOutlined,
	"cameraTwotone":                CameraTwotone,
	"carFilled":                    CarFilled,
	"carOutlined":                  CarOutlined,
	"carTwotone":                   CarTwotone,
	"caretDownFilled":              CaretDownFilled,
	"caretDownOutlined":            CaretDownOutlined,
	"caretLeftFilled":              CaretLeftFilled,
	"caretLeftOutlined":            CaretLeftOutlined,
	"caretRightFilled":             CaretRightFilled,
	"caretRightOutlined":           CaretRightOutlined,
	"caretUpFilled":                CaretUpFilled,
	"caretUpOutlined":              CaretUpOutlined,
	"carryOutFilled":               CarryOutFilled,
	"carryOutOutlined":             CarryOutOutlined,
	"carryOutTwotone":              CarryOutTwotone,
	"checkCircleFilled":            CheckCircleFilled,
	"checkCircleOutlined":          CheckCircleOutlined,
	"checkCircleTwotone":           CheckCircleTwotone,
	"checkOutlined":                CheckOutlined,
	"checkSquareFilled":            CheckSquareFilled,
	"checkSquareOutlined":          CheckSquareOutlined,
	"checkSquareTwotone":           CheckSquareTwotone,
	"chromeFilled":                 ChromeFilled,
	"chromeOutlined":               ChromeOutlined,
	"ciCircleFilled":               CiCircleFilled,
	"ciCircleOutlined":             CiCircleOutlined,
	"ciCircleTwotone":              CiCircleTwotone,
	"ciOutlined":                   CiOutlined,
	"ciTwotone":                    CiTwotone,
	"clearOutlined":                ClearOutlined,
	"clockCircleFilled":            ClockCircleFilled,
	"clockCircleOutlined":          ClockCircleOutlined,
	"clockCircleTwotone":           ClockCircleTwotone,
	"closeCircleFilled":            CloseCircleFilled,
	"closeCircleOutlined":          CloseCircleOutlined,
	"closeCircleTwotone":           CloseCircleTwotone,
	"closeOutlined":                CloseOutlined,
	"closeSquareFilled":            CloseSquareFilled,
	"closeSquareOutlined":          CloseSquareOutlined,
	"closeSquareTwotone":           CloseSquareTwotone,
	"cloudDownloadOutlined":        CloudDownloadOutlined,
	"cloudFilled":                  CloudFilled,
	"cloudOutlined":                CloudOutlined,
	"cloudServerOutlined":          CloudServerOutlined,
	"cloudSyncOutlined":            CloudSyncOutlined,
	"cloudTwotone":                 CloudTwotone,
	"cloudUploadOutlined":          CloudUploadOutlined,
	"clusterOutlined":              ClusterOutlined,
	"codeFilled":                   CodeFilled,
	"codeOutlined":                 CodeOutlined,
	"codeSandboxCircleFilled":      CodeSandboxCircleFilled,
	"codeSandboxOutlined":          CodeSandboxOutlined,
	"codeSandboxSquareFilled":      CodeSandboxSquareFilled,
	"codeTwotone":                  CodeTwotone,
	"codepenCircleFilled":          CodepenCircleFilled,
	"codepenCircleOutlined":        CodepenCircleOutlined,
	"codepenOutlined":              CodepenOutlined,
	"codepenSquareFilled":          CodepenSquareFilled,
	"coffeeOutlined":               CoffeeOutlined,
	"columnHeightOutlined":         ColumnHeightOutlined,
	"columnWidthOutlined":          ColumnWidthOutlined,
	"commentOutlined":              CommentOutlined,
	"compassFilled":                CompassFilled,
	"compassOutlined":              CompassOutlined,
	"compassTwotone":               CompassTwotone,
	"compressOutlined":             CompressOutlined,
	"consoleSqlOutlined":           ConsoleSqlOutlined,
	"contactsFilled":               ContactsFilled,
	"contactsOutlined":             ContactsOutlined,
	"contactsTwotone":              ContactsTwotone,
	"containerFilled":              ContainerFilled,
	"containerOutlined":            ContainerOutlined,
	"containerTwotone":             ContainerTwotone,
	"controlFilled":                ControlFilled,
	"controlOutlined":              ControlOutlined,
	"controlTwotone":               ControlTwotone,
	"copyFilled":                   CopyFilled,
	"copyOutlined":                 CopyOutlined,
	"copyTwotone":                  CopyTwotone,
	"copyrightCircleFilled":        CopyrightCircleFilled,
	"copyrightCircleOutlined":      CopyrightCircleOutlined,
	"copyrightCircleTwotone":       CopyrightCircleTwotone,
	"copyrightOutlined":            CopyrightOutlined,
	"copyrightTwotone":             CopyrightTwotone,
	"creditCardFilled":             CreditCardFilled,
	"creditCardOutlined":           CreditCardOutlined,
	"creditCardTwotone":            CreditCardTwotone,
	"crownFilled":                  CrownFilled,
	"crownOutlined":                CrownOutlined,
	"crownTwotone":                 CrownTwotone,
	"customerServiceFilled":        CustomerServiceFilled,
	"customerServiceOutlined":      CustomerServiceOutlined,
	"customerServiceTwotone":       CustomerServiceTwotone,
	"dashOutlined":                 DashOutlined,
	"dashboardFilled":              DashboardFilled,
	"dashboardOutlined":            DashboardOutlined,
	"dashboardTwotone":             DashboardTwotone,
	"databaseFilled":               DatabaseFilled,
	"databaseOutlined":             DatabaseOutlined,
	"databaseTwotone":              DatabaseTwotone,
	"deleteColumnOutlined":         DeleteColumnOutlined,
	"deleteFilled":                 DeleteFilled,
	"deleteOutlined":               DeleteOutlined,
	"deleteRowOutlined":            DeleteRowOutlined,
	"deleteTwotone":                DeleteTwotone,
	"deliveredProcedureOutlined":   DeliveredProcedureOutlined,
	"deploymentUnitOutlined":       DeploymentUnitOutlined,
	"desktopOutlined":              DesktopOutlined,
	"diffFilled":                   DiffFilled,
	"diffOutlined":                 DiffOutlined,
	"diffTwotone":                  DiffTwotone,
	"dingdingOutlined":             DingdingOutlined,
	"dingtalkCircleFilled":         DingtalkCircleFilled,
	"dingtalkOutlined":             DingtalkOutlined,
	"dingtalkSquareFilled":         DingtalkSquareFilled,
	"disconnectOutlined":           DisconnectOutlined,
	"dislikeFilled":                DislikeFilled,
	"dislikeOutlined":              DislikeOutlined,
	"dislikeTwotone":               DislikeTwotone,
	"dollarCircleFilled":           DollarCircleFilled,
	"dollarCircleOutlined":         DollarCircleOutlined,
	"dollarCircleTwotone":          DollarCircleTwotone,
	"dollarOutlined":               DollarOutlined,
	"dollarTwotone":                DollarTwotone,
	"dotChartOutlined":             DotChartOutlined,
	"doubleLeftOutlined":           DoubleLeftOutlined,
	"doubleRightOutlined":          DoubleRightOutlined,
	"downCircleFilled":             DownCircleFilled,
	"downCircleOutlined":           DownCircleOutlined,
	"downCircleTwotone":            DownCircleTwotone,
	"downOutlined":                 DownOutlined,
	"downSquareFilled":             DownSquareFilled,
	"downSquareOutlined":           DownSquareOutlined,
	"downSquareTwotone":            DownSquareTwotone,
	"downloadOutlined":             DownloadOutlined,
	"dragOutlined":                 DragOutlined,
	"dribbbleCircleFilled":         DribbbleCircleFilled,
	"dribbbleOutlined":             DribbbleOutlined,
	"dribbbleSquareFilled":         DribbbleSquareFilled,
	"dribbbleSquareOutlined":       DribbbleSquareOutlined,
	"dropboxCircleFilled":          DropboxCircleFilled,
	"dropboxOutlined":              DropboxOutlined,
	"dropboxSquareFilled":          DropboxSquareFilled,
	"editFilled":                   EditFilled,
	"editOutlined":                 EditOutlined,
	"editTwotone":                  EditTwotone,
	"ellipsisOutlined":             EllipsisOutlined,
	"enterOutlined":                EnterOutlined,
	"environmentFilled":            EnvironmentFilled,
	"environmentOutlined":          EnvironmentOutlined,
	"environmentTwotone":           EnvironmentTwotone,
	"euroCircleFilled":             EuroCircleFilled,
	"euroCircleOutlined":           EuroCircleOutlined,
	"euroCircleTwotone":            EuroCircleTwotone,
	"euroOutlined":                 EuroOutlined,
	"euroTwotone":                  EuroTwotone,
	"exceptionOutlined":            ExceptionOutlined,
	"exclamationCircleFilled":      ExclamationCircleFilled,
	"exclamationCircleOutlined":    ExclamationCircleOutlined,
	"exclamationCircleTwotone":     ExclamationCircleTwotone,
	"exclamationOutlined":          ExclamationOutlined,
	"expandAltOutlined":            ExpandAltOutlined,
	"expandOutlined":               ExpandOutlined,
	"experimentFilled":             ExperimentFilled,
	"experimentOutlined":           ExperimentOutlined,
	"experimentTwotone":            ExperimentTwotone,
	"exportOutlined":               ExportOutlined,
	"eyeFilled":                    EyeFilled,
	"eyeInvisibleFilled":           EyeInvisibleFilled,
	"eyeInvisibleOutlined":         EyeInvisibleOutlined,
	"eyeInvisibleTwotone":          EyeInvisibleTwotone,
	"eyeOutlined":                  EyeOutlined,
	"eyeTwotone":                   EyeTwotone,
	"facebookFilled":               FacebookFilled,
	"facebookOutlined":             FacebookOutlined,
	"fallOutlined":                 FallOutlined,
	"fastBackwardFilled":           FastBackwardFilled,
	"fastBackwardOutlined":         FastBackwardOutlined,
	"fastForwardFilled":            FastForwardFilled,
	"fastForwardOutlined":          FastForwardOutlined,
	"fieldBinaryOutlined":          FieldBinaryOutlined,
	"fieldNumberOutlined":          FieldNumberOutlined,
	"fieldStringOutlined":          FieldStringOutlined,
	"fieldTimeOutlined":            FieldTimeOutlined,
	"fileAddFilled":                FileAddFilled,
	"fileAddOutlined":              FileAddOutlined,
	"fileAddTwotone":               FileAddTwotone,
	"fileDoneOutlined":             FileDoneOutlined,
	"fileExcelFilled":              FileExcelFilled,
	"fileExcelOutlined":            FileExcelOutlined,
	"fileExcelTwotone":             FileExcelTwotone,
	"fileExclamationFilled":        FileExclamationFilled,
	"fileExclamationOutlined":      FileExclamationOutlined,
	"fileExclamationTwotone":       FileExclamationTwotone,
	"fileFilled":                   FileFilled,
	"fileGifOutlined":              FileGifOutlined,
	"fileImageFilled":              FileImageFilled,
	"fileImageOutlined":            FileImageOutlined,
	"fileImageTwotone":             FileImageTwotone,
	"fileJpgOutlined":              FileJpgOutlined,
	"fileMarkdownFilled":           FileMarkdownFilled,
	"fileMarkdownOutlined":         FileMarkdownOutlined,
	"fileMarkdownTwotone":          FileMarkdownTwotone,
	"fileOutlined":                 FileOutlined,
	"filePdfFilled":                FilePdfFilled,
	"filePdfOutlined":              FilePdfOutlined,
	"filePdfTwotone":               FilePdfTwotone,
	"filePptFilled":                FilePptFilled,
	"filePptOutlined":              FilePptOutlined,
	"filePptTwotone":               FilePptTwotone,
	"fileProtectOutlined":          FileProtectOutlined,
	"fileSearchOutlined":           FileSearchOutlined,
	"fileSyncOutlined":             FileSyncOutlined,
	"fileTextFilled":               FileTextFilled,
	"fileTextOutlined":             FileTextOutlined,
	"fileTextTwotone":              FileTextTwotone,
	"fileTwotone":                  FileTwotone,
	"fileUnknownFilled":            FileUnknownFilled,
	"fileUnknownOutlined":          FileUnknownOutlined,
	"fileUnknownTwotone":           FileUnknownTwotone,
	"fileWordFilled":               FileWordFilled,
	"fileWordOutlined":             FileWordOutlined,
	"fileWordTwotone":              FileWordTwotone,
	"fileZipFilled":                FileZipFilled,
	"fileZipOutlined":              FileZipOutlined,
	"fileZipTwotone":               FileZipTwotone,
	"filterFilled":                 FilterFilled,
	"filterOutlined":               FilterOutlined,
	"filterTwotone":                FilterTwotone,
	"fireFilled":                   FireFilled,
	"fireOutlined":                 FireOutlined,
	"fireTwotone":                  FireTwotone,
	"flagFilled":                   FlagFilled,
	"flagOutlined":                 FlagOutlined,
	"flagTwotone":                  FlagTwotone,
	"folderAddFilled":              FolderAddFilled,
	"folderAddOutlined":            FolderAddOutlined,
	"folderAddTwotone":             FolderAddTwotone,
	"folderFilled":                 FolderFilled,
	"folderOpenFilled":             FolderOpenFilled,
	"folderOpenOutlined":           FolderOpenOutlined,
	"folderOpenTwotone":            FolderOpenTwotone,
	"folderOutlined":               FolderOutlined,
	"folderTwotone":                FolderTwotone,
	"folderViewOutlined":           FolderViewOutlined,
	"fontColorsOutlined":           FontColorsOutlined,
	"fontSizeOutlined":             FontSizeOutlined,
	"forkOutlined":                 ForkOutlined,
	"formOutlined":                 FormOutlined,
	"formatPainterFilled":          FormatPainterFilled,
	"formatPainterOutlined":        FormatPainterOutlined,
	"forwardFilled":                ForwardFilled,
	"forwardOutlined":              ForwardOutlined,
	"frownFilled":                  FrownFilled,
	"frownOutlined":                FrownOutlined,
	"frownTwotone":                 FrownTwotone,
	"fullscreenExitOutlined":       FullscreenExitOutlined,
	"fullscreenOutlined":           FullscreenOutlined,
	"functionOutlined":             FunctionOutlined,
	"fundFilled":                   FundFilled,
	"fundOutlined":                 FundOutlined,
	"fundProjectionScreenOutlined": FundProjectionScreenOutlined,
	"fundTwotone":                  FundTwotone,
	"fundViewOutlined":             FundViewOutlined,
	"funnelPlotFilled":             FunnelPlotFilled,
	"funnelPlotOutlined":           FunnelPlotOutlined,
	"funnelPlotTwotone":            FunnelPlotTwotone,
	"gatewayOutlined":              GatewayOutlined,
	"gifOutlined":                  GifOutlined,
	"giftFilled":                   GiftFilled,
	"giftOutlined":                 GiftOutlined,
	"giftTwotone":                  GiftTwotone,
	"githubFilled":                 GithubFilled,
	"githubOutlined":               GithubOutlined,
	"gitlabFilled":                 GitlabFilled,
	"gitlabOutlined":               GitlabOutlined,
	"globalOutlined":               GlobalOutlined,
	"goldFilled":                   GoldFilled,
	"goldOutlined":                 GoldOutlined,
	"goldTwotone":                  GoldTwotone,
	"goldenFilled":                 GoldenFilled,
	"googleCircleFilled":           GoogleCircleFilled,
	"googleOutlined":               GoogleOutlined,
	"googlePlusCircleFilled":       GooglePlusCircleFilled,
	"googlePlusOutlined":           GooglePlusOutlined,
	"googlePlusSquareFilled":       GooglePlusSquareFilled,
	"googleSquareFilled":           GoogleSquareFilled,
	"groupOutlined":                GroupOutlined,
	"hddFilled":                    HddFilled,
	"hddOutlined":                  HddOutlined,
	"hddTwotone":                   HddTwotone,
	"heartFilled":                  HeartFilled,
	"heartOutlined":                HeartOutlined,
	"heartTwotone":                 HeartTwotone,
	"heatMapOutlined":              HeatMapOutlined,
	"highlightFilled":              HighlightFilled,
	"highlightOutlined":            HighlightOutlined,
	"highlightTwotone":             HighlightTwotone,
	"historyOutlined":              HistoryOutlined,
	"holderOutlined":               HolderOutlined,
	"homeFilled":                   HomeFilled,
	"homeOutlined":                 HomeOutlined,
	"homeTwotone":                  HomeTwotone,
	"hourglassFilled":              HourglassFilled,
	"hourglassOutlined":            HourglassOutlined,
	"hourglassTwotone":             HourglassTwotone,
	"htmlFiveFilled":               HtmlFiveFilled,
	"htmlFiveOutlined":             HtmlFiveOutlined,
	"htmlFiveTwotone":              HtmlFiveTwotone,
	"idcardFilled":                 IdcardFilled,
	"idcardOutlined":               IdcardOutlined,
	"idcardTwotone":                IdcardTwotone,
	"ieCircleFilled":               IeCircleFilled,
	"ieOutlined":                   IeOutlined,
	"ieSquareFilled":               IeSquareFilled,
	"importOutlined":               ImportOutlined,
	"inboxOutlined":                InboxOutlined,
	"infoCircleFilled":             InfoCircleFilled,
	"infoCircleOutlined":           InfoCircleOutlined,
	"infoCircleTwotone":            InfoCircleTwotone,
	"infoOutlined":                 InfoOutlined,
	"insertRowAboveOutlined":       InsertRowAboveOutlined,
	"insertRowBelowOutlined":       InsertRowBelowOutlined,
	"insertRowLeftOutlined":        InsertRowLeftOutlined,
	"insertRowRightOutlined":       InsertRowRightOutlined,
	"instagramFilled":              InstagramFilled,
	"instagramOutlined":            InstagramOutlined,
	"insuranceFilled":              InsuranceFilled,
	"insuranceOutlined":            InsuranceOutlined,
	"insuranceTwotone":             InsuranceTwotone,
	"interactionFilled":            InteractionFilled,
	"interactionOutlined":          InteractionOutlined,
	"interactionTwotone":           InteractionTwotone,
	"issuesCloseOutlined":          IssuesCloseOutlined,
	"italicOutlined":               ItalicOutlined,
	"keyOutlined":                  KeyOutlined,
	"laptopOutlined":               LaptopOutlined,
	"layoutFilled":                 LayoutFilled,
	"layoutOutlined":               LayoutOutlined,
	"layoutTwotone":                LayoutTwotone,
	"leftCircleFilled":             LeftCircleFilled,
	"leftCircleOutlined":           LeftCircleOutlined,
	"leftCircleTwotone":            LeftCircleTwotone,
	"leftOutlined":                 LeftOutlined,
	"leftSquareFilled":             LeftSquareFilled,
	"leftSquareOutlined":           LeftSquareOutlined,
	"leftSquareTwotone":            LeftSquareTwotone,
	"likeFilled":                   LikeFilled,
	"likeOutlined":                 LikeOutlined,
	"likeTwotone":                  LikeTwotone,
	"lineChartOutlined":            LineChartOutlined,
	"lineHeightOutlined":           LineHeightOutlined,
	"lineOutlined":                 LineOutlined,
	"linkOutlined":                 LinkOutlined,
	"linkedinFilled":               LinkedinFilled,
	"linkedinOutlined":             LinkedinOutlined,
	"loadingOutlined":              LoadingOutlined,
	"loadingThreeQuartersOutlined": LoadingThreeQuartersOutlined,
	"lockFilled":                   LockFilled,
	"lockOutlined":                 LockOutlined,
	"lockTwotone":                  LockTwotone,
	"loginOutlined":                LoginOutlined,
	"logoutOutlined":               LogoutOutlined,
	"macCommandFilled":             MacCommandFilled,
	"macCommandOutlined":           MacCommandOutlined,
	"mailFilled":                   MailFilled,
	"mailOutlined":                 MailOutlined,
	"mailTwotone":                  MailTwotone,
	"manOutlined":                  ManOutlined,
	"medicineBoxFilled":            MedicineBoxFilled,
	"medicineBoxOutlined":          MedicineBoxOutlined,
	"medicineBoxTwotone":           MedicineBoxTwotone,
	"mediumCircleFilled":           MediumCircleFilled,
	"mediumOutlined":               MediumOutlined,
	"mediumSquareFilled":           MediumSquareFilled,
	"mediumWorkmarkOutlined":       MediumWorkmarkOutlined,
	"mehFilled":                    MehFilled,
	"mehOutlined":                  MehOutlined,
	"mehTwotone":                   MehTwotone,
	"menuFoldOutlined":             MenuFoldOutlined,
	"menuOutlined":                 MenuOutlined,
	"menuUnfoldOutlined":           MenuUnfoldOutlined,
	"mergeCellsOutlined":           MergeCellsOutlined,
	"messageFilled":                MessageFilled,
	"messageOutlined":              MessageOutlined,
	"messageTwotone":               MessageTwotone,
	"minusCircleFilled":            MinusCircleFilled,
	"minusCircleOutlined":          MinusCircleOutlined,
	"minusCircleTwotone":           MinusCircleTwotone,
	"minusOutlined":                MinusOutlined,
	"minusSquareFilled":            MinusSquareFilled,
	"minusSquareOutlined":          MinusSquareOutlined,
	"minusSquareTwotone":           MinusSquareTwotone,
	"mobileFilled":                 MobileFilled,
	"mobileOutlined":               MobileOutlined,
	"mobileTwotone":                MobileTwotone,
	"moneyCollectFilled":           MoneyCollectFilled,
	"moneyCollectOutlined":         MoneyCollectOutlined,
	"moneyCollectTwotone":          MoneyCollectTwotone,
	"monitorOutlined":              MonitorOutlined,
	"moreOutlined":                 MoreOutlined,
	"nodeCollapseOutlined":         NodeCollapseOutlined,
	"nodeExpandOutlined":           NodeExpandOutlined,
	"nodeIndexOutlined":            NodeIndexOutlined,
	"notificationFilled":           NotificationFilled,
	"notificationOutlined":         NotificationOutlined,
	"notificationTwotone":          NotificationTwotone,
	"numberOutlined":               NumberOutlined,
	"oneToOneOutlined":             OneToOneOutlined,
	"orderedListOutlined":          OrderedListOutlined,
	"paperClipOutlined":            PaperClipOutlined,
	"partitionOutlined":            PartitionOutlined,
	"pauseCircleFilled":            PauseCircleFilled,
	"pauseCircleOutlined":          PauseCircleOutlined,
	"pauseCircleTwotone":           PauseCircleTwotone,
	"pauseOutlined":                PauseOutlined,
	"payCircleFilled":              PayCircleFilled,
	"payCircleOutlined":            PayCircleOutlined,
	"percentageOutlined":           PercentageOutlined,
	"phoneFilled":                  PhoneFilled,
	"phoneOutlined":                PhoneOutlined,
	"phoneTwotone":                 PhoneTwotone,
	"picCenterOutlined":            PicCenterOutlined,
	"picLeftOutlined":              PicLeftOutlined,
	"picRightOutlined":             PicRightOutlined,
	"pictureFilled":                PictureFilled,
	"pictureOutlined":              PictureOutlined,
	"pictureTwotone":               PictureTwotone,
	"pieChartFilled":               PieChartFilled,
	"pieChartOutlined":             PieChartOutlined,
	"pieChartTwotone":              PieChartTwotone,
	"playCircleFilled":             PlayCircleFilled,
	"playCircleOutlined":           PlayCircleOutlined,
	"playCircleTwotone":            PlayCircleTwotone,
	"playSquareFilled":             PlaySquareFilled,
	"playSquareOutlined":           PlaySquareOutlined,
	"playSquareTwotone":            PlaySquareTwotone,
	"plusCircleFilled":             PlusCircleFilled,
	"plusCircleOutlined":           PlusCircleOutlined,
	"plusCircleTwotone":            PlusCircleTwotone,
	"plusOutlined":                 PlusOutlined,
	"plusSquareFilled":             PlusSquareFilled,
	"plusSquareOutlined":           PlusSquareOutlined,
	"plusSquareTwotone":            PlusSquareTwotone,
	"poundCircleFilled":            PoundCircleFilled,
	"poundCircleOutlined":          PoundCircleOutlined,
	"poundCircleTwotone":           PoundCircleTwotone,
	"poundOutlined":                PoundOutlined,
	"poweroffOutlined":             PoweroffOutlined,
	"printerFilled":                PrinterFilled,
	"printerOutlined":              PrinterOutlined,
	"printerTwotone":               PrinterTwotone,
	"profileFilled":                ProfileFilled,
	"profileOutlined":              ProfileOutlined,
	"profileTwotone":               ProfileTwotone,
	"projectFilled":                ProjectFilled,
	"projectOutlined":              ProjectOutlined,
	"projectTwotone":               ProjectTwotone,
	"propertySafetyFilled":         PropertySafetyFilled,
	"propertySafetyOutlined":       PropertySafetyOutlined,
	"propertySafetyTwotone":        PropertySafetyTwotone,
	"pullRequestOutlined":          PullRequestOutlined,
	"pushpinFilled":                PushpinFilled,
	"pushpinOutlined":              PushpinOutlined,
	"pushpinTwotone":               PushpinTwotone,
	"qqCircleFilled":               QqCircleFilled,
	"qqOutlined":                   QqOutlined,
	"qqSquareFilled":               QqSquareFilled,
	"qrcodeOutlined":               QrcodeOutlined,
	"questionCircleFilled":         QuestionCircleFilled,
	"questionCircleOutlined":       QuestionCircleOutlined,
	"questionCircleTwotone":        QuestionCircleTwotone,
	"questionOutlined":             QuestionOutlined,
	"radarChartOutlined":           RadarChartOutlined,
	"radiusBottomleftOutlined":     RadiusBottomleftOutlined,
	"radiusBottomrightOutlined":    RadiusBottomrightOutlined,
	"radiusSettingOutlined":        RadiusSettingOutlined,
	"radiusUpleftOutlined":         RadiusUpleftOutlined,
	"radiusUprightOutlined":        RadiusUprightOutlined,
	"readFilled":                   ReadFilled,
	"readOutlined":                 ReadOutlined,
	"reconciliationFilled":         ReconciliationFilled,
	"reconciliationOutlined":       ReconciliationOutlined,
	"reconciliationTwotone":        ReconciliationTwotone,
	"redEnvelopeFilled":            RedEnvelopeFilled,
	"redEnvelopeOutlined":          RedEnvelopeOutlined,
	"redEnvelopeTwotone":           RedEnvelopeTwotone,
	"redditCircleFilled":           RedditCircleFilled,
	"redditOutlined":               RedditOutlined,
	"redditSquareFilled":           RedditSquareFilled,
	"redoOutlined":                 RedoOutlined,
	"reloadOutlined":               ReloadOutlined,
	"restFilled":                   RestFilled,
	"restOutlined":                 RestOutlined,
	"restTwotone":                  RestTwotone,
	"retweetOutlined":              RetweetOutlined,
	"rightCircleFilled":            RightCircleFilled,
	"rightCircleOutlined":          RightCircleOutlined,
	"rightCircleTwotone":           RightCircleTwotone,
	"rightOutlined":                RightOutlined,
	"rightSquareFilled":            RightSquareFilled,
	"rightSquareOutlined":          RightSquareOutlined,
	"rightSquareTwotone":           RightSquareTwotone,
	"riseOutlined":                 RiseOutlined,
	"robotFilled":                  RobotFilled,
	"robotOutlined":                RobotOutlined,
	"rocketFilled":                 RocketFilled,
	"rocketOutlined":               RocketOutlined,
	"rocketTwotone":                RocketTwotone,
	"rollbackOutlined":             RollbackOutlined,
	"rotateLeftOutlined":           RotateLeftOutlined,
	"rotateRightOutlined":          RotateRightOutlined,
	"safetyCertificateFilled":      SafetyCertificateFilled,
	"safetyCertificateOutlined":    SafetyCertificateOutlined,
	"safetyCertificateTwotone":     SafetyCertificateTwotone,
	"safetyOutlined":               SafetyOutlined,
	"saveFilled":                   SaveFilled,
	"saveOutlined":                 SaveOutlined,
	"saveTwotone":                  SaveTwotone,
	"scanOutlined":                 ScanOutlined,
	"scheduleFilled":               ScheduleFilled,
	"scheduleOutlined":             ScheduleOutlined,
	"scheduleTwotone":              ScheduleTwotone,
	"scissorOutlined":              ScissorOutlined,
	"searchOutlined":               SearchOutlined,
	"securityScanFilled":           SecurityScanFilled,
	"securityScanOutlined":         SecurityScanOutlined,
	"securityScanTwotone":          SecurityScanTwotone,
	"selectOutlined":               SelectOutlined,
	"sendOutlined":                 SendOutlined,
	"settingFilled":                SettingFilled,
	"settingOutlined":              SettingOutlined,
	"settingTwotone":               SettingTwotone,
	"shakeOutlined":                ShakeOutlined,
	"shareAltOutlined":             ShareAltOutlined,
	"shopFilled":                   ShopFilled,
	"shopOutlined":                 ShopOutlined,
	"shopTwotone":                  ShopTwotone,
	"shoppingCartOutlined":         ShoppingCartOutlined,
	"shoppingFilled":               ShoppingFilled,
	"shoppingOutlined":             ShoppingOutlined,
	"shoppingTwotone":              ShoppingTwotone,
	"shrinkOutlined":               ShrinkOutlined,
	"signalFilled":                 SignalFilled,
	"sisternodeOutlined":           SisternodeOutlined,
	"sketchCircleFilled":           SketchCircleFilled,
	"sketchOutlined":               SketchOutlined,
	"sketchSquareFilled":           SketchSquareFilled,
	"skinFilled":                   SkinFilled,
	"skinOutlined":                 SkinOutlined,
	"skinTwotone":                  SkinTwotone,
	"skypeFilled":                  SkypeFilled,
	"skypeOutlined":                SkypeOutlined,
	"slackCircleFilled":            SlackCircleFilled,
	"slackOutlined":                SlackOutlined,
	"slackSquareFilled":            SlackSquareFilled,
	"slackSquareOutlined":          SlackSquareOutlined,
	"slidersFilled":                SlidersFilled,
	"slidersOutlined":              SlidersOutlined,
	"slidersTwotone":               SlidersTwotone,
	"smallDashOutlined":            SmallDashOutlined,
	"smileFilled":                  SmileFilled,
	"smileOutlined":                SmileOutlined,
	"smileTwotone":                 SmileTwotone,
	"snippetsFilled":               SnippetsFilled,
	"snippetsOutlined":             SnippetsOutlined,
	"snippetsTwotone":              SnippetsTwotone,
	"solutionOutlined":             SolutionOutlined,
	"sortAscendingOutlined":        SortAscendingOutlined,
	"sortDescendingOutlined":       SortDescendingOutlined,
	"soundFilled":                  SoundFilled,
	"soundOutlined":                SoundOutlined,
	"soundTwotone":                 SoundTwotone,
	"splitCellsOutlined":           SplitCellsOutlined,
	"starFilled":                   StarFilled,
	"starOutlined":                 StarOutlined,
	"starTwotone":                  StarTwotone,
	"stepBackwardFilled":           StepBackwardFilled,
	"stepBackwardOutlined":         StepBackwardOutlined,
	"stepForwardFilled":            StepForwardFilled,
	"stepForwardOutlined":          StepForwardOutlined,
	"stockOutlined":                StockOutlined,
	"stopFilled":                   StopFilled,
	"stopOutlined":                 StopOutlined,
	"stopTwotone":                  StopTwotone,
	"strikethroughOutlined":        StrikethroughOutlined,
	"subnodeOutlined":              SubnodeOutlined,
	"swapLeftOutlined":             SwapLeftOutlined,
	"swapOutlined":                 SwapOutlined,
	"swapRightOutlined":            SwapRightOutlined,
	"switcherFilled":               SwitcherFilled,
	"switcherOutlined":             SwitcherOutlined,
	"switcherTwotone":              SwitcherTwotone,
	"syncOutlined":                 SyncOutlined,
	"tableOutlined":                TableOutlined,
	"tabletFilled":                 TabletFilled,
	"tabletOutlined":               TabletOutlined,
	"tabletTwotone":                TabletTwotone,
	"tagFilled":                    TagFilled,
	"tagOutlined":                  TagOutlined,
	"tagTwotone":                   TagTwotone,
	"tagsFilled":                   TagsFilled,
	"tagsOutlined":                 TagsOutlined,
	"tagsTwotone":                  TagsTwotone,
	"taobaoCircleFilled":           TaobaoCircleFilled,
	"taobaoCircleOutlined":         TaobaoCircleOutlined,
	"taobaoOutlined":               TaobaoOutlined,
	"taobaoSquareFilled":           TaobaoSquareFilled,
	"teamOutlined":                 TeamOutlined,
	"thunderboltFilled":            ThunderboltFilled,
	"thunderboltOutlined":          ThunderboltOutlined,
	"thunderboltTwotone":           ThunderboltTwotone,
	"toTopOutlined":                ToTopOutlined,
	"toolFilled":                   ToolFilled,
	"toolOutlined":                 ToolOutlined,
	"toolTwotone":                  ToolTwotone,
	"trademarkCircleFilled":        TrademarkCircleFilled,
	"trademarkCircleOutlined":      TrademarkCircleOutlined,
	"trademarkCircleTwotone":       TrademarkCircleTwotone,
	"trademarkOutlined":            TrademarkOutlined,
	"transactionOutlined":          TransactionOutlined,
	"translationOutlined":          TranslationOutlined,
	"trophyFilled":                 TrophyFilled,
	"trophyOutlined":               TrophyOutlined,
	"trophyTwotone":                TrophyTwotone,
	"twitterCircleFilled":          TwitterCircleFilled,
	"twitterOutlined":              TwitterOutlined,
	"twitterSquareFilled":          TwitterSquareFilled,
	"underlineOutlined":            UnderlineOutlined,
	"undoOutlined":                 UndoOutlined,
	"ungroupOutlined":              UngroupOutlined,
	"unlockFilled":                 UnlockFilled,
	"unlockOutlined":               UnlockOutlined,
	"unlockTwotone":                UnlockTwotone,
	"unorderedListOutlined":        UnorderedListOutlined,
	"upCircleFilled":               UpCircleFilled,
	"upCircleOutlined":             UpCircleOutlined,
	"upCircleTwotone":              UpCircleTwotone,
	"upOutlined":                   UpOutlined,
	"upSquareFilled":               UpSquareFilled,
	"upSquareOutlined":             UpSquareOutlined,
	"upSquareTwotone":              UpSquareTwotone,
	"uploadOutlined":               UploadOutlined,
	"usbFilled":                    UsbFilled,
	"usbOutlined":                  UsbOutlined,
	"usbTwotone":                   UsbTwotone,
	"userAddOutlined":              UserAddOutlined,
	"userDeleteOutlined":           UserDeleteOutlined,
	"userOutlined":                 UserOutlined,
	"userSwitchOutlined":           UserSwitchOutlined,
	"usergroupAddOutlined":         UsergroupAddOutlined,
	"usergroupDeleteOutlined":      UsergroupDeleteOutlined,
	"verifiedOutlined":             VerifiedOutlined,
	"verticalAlignBottomOutlined":  VerticalAlignBottomOutlined,
	"verticalAlignMiddleOutlined":  VerticalAlignMiddleOutlined,
	"verticalAlignTopOutlined":     VerticalAlignTopOutlined,
	"verticalLeftOutlined":         VerticalLeftOutlined,
	"verticalRightOutlined":        VerticalRightOutlined,
	"videoCameraAddOutlined":       VideoCameraAddOutlined,
	"videoCameraFilled":            VideoCameraFilled,
	"videoCameraOutlined":          VideoCameraOutlined,
	"videoCameraTwotone":           VideoCameraTwotone,
	"walletFilled":                 WalletFilled,
	"walletOutlined":               WalletOutlined,
	"walletTwotone":                WalletTwotone,
	"warningFilled":                WarningFilled,
	"warningOutlined":              WarningOutlined,
	"warningTwotone":               WarningTwotone,
	"wechatFilled":                 WechatFilled,
	"wechatOutlined":               WechatOutlined,
	"weiboCircleFilled":            WeiboCircleFilled,
	"weiboCircleOutlined":          WeiboCircleOutlined,
	"weiboOutlined":                WeiboOutlined,
	"weiboSquareFilled":            WeiboSquareFilled,
	"weiboSquareOutlined":          WeiboSquareOutlined,
	"whatsAppOutlined":             WhatsAppOutlined,
	"wifiOutlined":                 WifiOutlined,
	"windowsFilled":                WindowsFilled,
	"windowsOutlined":              WindowsOutlined,
	"womanOutlined":                WomanOutlined,
	"yahooFilled":                  YahooFilled,
	"yahooOutlined":                YahooOutlined,
	"youtubeFilled":                YoutubeFilled,
	"youtubeOutlined":              YoutubeOutlined,
	"yuqueFilled":                  YuqueFilled,
	"yuqueOutlined":                YuqueOutlined,
	"zhihuCircleFilled":            ZhihuCircleFilled,
	"zhihuOutlined":                ZhihuOutlined,
	"zhihuSquareFilled":            ZhihuSquareFilled,
	"zoomInOutlined":               ZoomInOutlined,
	"zoomOutOutlined":              ZoomOutOutlined,
}

func AccountBookFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 184H712v-64c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v64H384v-64c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v64H144c-17.7 0-32 14.3-32 32v664c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V216c0-17.7-14.3-32-32-32M648.3 426.8l-87.7 161.1h45.7c5.5 0 10 4.5 10 10v21.3c0 5.5-4.5 10-10 10h-63.4v29.7h63.4c5.5 0 10 4.5 10 10v21.3c0 5.5-4.5 10-10 10h-63.4V752c0 5.5-4.5 10-10 10h-41.3c-5.5 0-10-4.5-10-10v-51.8h-63.1c-5.5 0-10-4.5-10-10v-21.3c0-5.5 4.5-10 10-10h63.1v-29.7h-63.1c-5.5 0-10-4.5-10-10v-21.3c0-5.5 4.5-10 10-10h45.2l-88-161.1c-2.6-4.8-.9-10.9 4-13.6c1.5-.8 3.1-1.2 4.8-1.2h46c3.8 0 7.2 2.1 8.9 5.5l72.9 144.3l73.2-144.3a10 10 0 0 1 8.9-5.5h45c5.5 0 10 4.5 10 10c.1 1.7-.3 3.3-1.1 4.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccountBookOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 184H712v-64c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v64H384v-64c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v64H144c-17.7 0-32 14.3-32 32v664c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V216c0-17.7-14.3-32-32-32m-40 656H184V256h128v48c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-48h256v48c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-48h128zM639.5 414h-45c-3 0-5.8 1.7-7.1 4.4L514 563.8h-2.8l-73.4-145.4a8 8 0 0 0-7.1-4.4h-46c-1.3 0-2.7.3-3.8 1c-3.9 2.1-5.3 7-3.2 10.9l89.3 164h-48.6c-4.4 0-8 3.6-8 8v21.3c0 4.4 3.6 8 8 8h65.1v33.7h-65.1c-4.4 0-8 3.6-8 8v21.3c0 4.4 3.6 8 8 8h65.1V752c0 4.4 3.6 8 8 8h41.3c4.4 0 8-3.6 8-8v-53.8h65.4c4.4 0 8-3.6 8-8v-21.3c0-4.4-3.6-8-8-8h-65.4v-33.7h65.4c4.4 0 8-3.6 8-8v-21.3c0-4.4-3.6-8-8-8h-49.1l89.3-164.1c.6-1.2 1-2.5 1-3.8c.1-4.4-3.4-8-7.9-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccountBookTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M712 304c0 4.4-3.6 8-8 8h-56c-4.4 0-8-3.6-8-8v-48H384v48c0 4.4-3.6 8-8 8h-56c-4.4 0-8-3.6-8-8v-48H184v584h656V256H712zm-65.6 121.8l-89.3 164.1h49.1c4.4 0 8 3.6 8 8v21.3c0 4.4-3.6 8-8 8h-65.4v33.7h65.4c4.4 0 8 3.6 8 8v21.3c0 4.4-3.6 8-8 8h-65.4V752c0 4.4-3.6 8-8 8h-41.3c-4.4 0-8-3.6-8-8v-53.8h-65.1c-4.4 0-8-3.6-8-8v-21.3c0-4.4 3.6-8 8-8h65.1v-33.7h-65.1c-4.4 0-8-3.6-8-8v-21.3c0-4.4 3.6-8 8-8H467l-89.3-164c-2.1-3.9-.7-8.8 3.2-10.9c1.1-.7 2.5-1 3.8-1h46a8 8 0 0 1 7.1 4.4l73.4 145.4h2.8l73.4-145.4c1.3-2.7 4.1-4.4 7.1-4.4h45c4.5 0 8 3.6 7.9 8c0 1.3-.4 2.6-1 3.8"/><path fill="currentColor" d="M639.5 414h-45c-3 0-5.8 1.7-7.1 4.4L514 563.8h-2.8l-73.4-145.4a8 8 0 0 0-7.1-4.4h-46c-1.3 0-2.7.3-3.8 1c-3.9 2.1-5.3 7-3.2 10.9l89.3 164h-48.6c-4.4 0-8 3.6-8 8v21.3c0 4.4 3.6 8 8 8h65.1v33.7h-65.1c-4.4 0-8 3.6-8 8v21.3c0 4.4 3.6 8 8 8h65.1V752c0 4.4 3.6 8 8 8h41.3c4.4 0 8-3.6 8-8v-53.8h65.4c4.4 0 8-3.6 8-8v-21.3c0-4.4-3.6-8-8-8h-65.4v-33.7h65.4c4.4 0 8-3.6 8-8v-21.3c0-4.4-3.6-8-8-8h-49.1l89.3-164.1c.6-1.2 1-2.5 1-3.8c.1-4.4-3.4-8-7.9-8"/><path fill="currentColor" d="M880 184H712v-64c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v64H384v-64c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v64H144c-17.7 0-32 14.3-32 32v664c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V216c0-17.7-14.3-32-32-32m-40 656H184V256h128v48c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-48h256v48c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-48h128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AimOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M952 474H829.8C812.5 327.6 696.4 211.5 550 194.2V72c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8v122.2C327.6 211.5 211.5 327.6 194.2 474H72c-4.4 0-8 3.6-8 8v60c0 4.4 3.6 8 8 8h122.2C211.5 696.4 327.6 812.5 474 829.8V952c0 4.4 3.6 8 8 8h60c4.4 0 8-3.6 8-8V829.8C696.4 812.5 812.5 696.4 829.8 550H952c4.4 0 8-3.6 8-8v-60c0-4.4-3.6-8-8-8M512 756c-134.8 0-244-109.2-244-244s109.2-244 244-244s244 109.2 244 244s-109.2 244-244 244"/><path fill="currentColor" d="M512 392c-32.1 0-62.1 12.4-84.8 35.2c-22.7 22.7-35.2 52.7-35.2 84.8s12.5 62.1 35.2 84.8C449.9 619.4 480 632 512 632s62.1-12.5 84.8-35.2C619.4 574.1 632 544 632 512s-12.5-62.1-35.2-84.8C574.1 404.4 544.1 392 512 392"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 244c176.18 0 319 142.82 319 319v233a32 32 0 0 1-32 32H225a32 32 0 0 1-32-32V563c0-176.18 142.82-319 319-319M484 68h56a8 8 0 0 1 8 8v96a8 8 0 0 1-8 8h-56a8 8 0 0 1-8-8V76a8 8 0 0 1 8-8M177.25 191.66a8 8 0 0 1 11.32 0l67.88 67.88a8 8 0 0 1 0 11.31l-39.6 39.6a8 8 0 0 1-11.31 0l-67.88-67.88a8 8 0 0 1 0-11.31l39.6-39.6zm669.6 0l39.6 39.6a8 8 0 0 1 0 11.3l-67.88 67.9a8 8 0 0 1-11.32 0l-39.6-39.6a8 8 0 0 1 0-11.32l67.89-67.88a8 8 0 0 1 11.31 0M192 892h640a32 32 0 0 1 32 32v24a8 8 0 0 1-8 8H168a8 8 0 0 1-8-8v-24a32 32 0 0 1 32-32m148-317v253h64V575z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M193 796c0 17.7 14.3 32 32 32h574c17.7 0 32-14.3 32-32V563c0-176.2-142.8-319-319-319S193 386.8 193 563zm72-233c0-136.4 110.6-247 247-247s247 110.6 247 247v193H404V585c0-5.5-4.5-10-10-10h-44c-5.5 0-10 4.5-10 10v171h-75zm-48.1-252.5l39.6-39.6c3.1-3.1 3.1-8.2 0-11.3l-67.9-67.9a8.03 8.03 0 0 0-11.3 0l-39.6 39.6a8.03 8.03 0 0 0 0 11.3l67.9 67.9c3.1 3.1 8.1 3.1 11.3 0m669.6-79.2l-39.6-39.6a8.03 8.03 0 0 0-11.3 0l-67.9 67.9a8.03 8.03 0 0 0 0 11.3l39.6 39.6c3.1 3.1 8.2 3.1 11.3 0l67.9-67.9c3.1-3.2 3.1-8.2 0-11.3M832 892H192c-17.7 0-32 14.3-32 32v24c0 4.4 3.6 8 8 8h688c4.4 0 8-3.6 8-8v-24c0-17.7-14.3-32-32-32M484 180h56c4.4 0 8-3.6 8-8V76c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v96c0 4.4 3.6 8 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M340 585c0-5.5 4.5-10 10-10h44c5.5 0 10 4.5 10 10v171h355V563c0-136.4-110.6-247-247-247S265 426.6 265 563v193h75z"/><path fill="currentColor" d="m216.9 310.5l39.6-39.6c3.1-3.1 3.1-8.2 0-11.3l-67.9-67.9a8.03 8.03 0 0 0-11.3 0l-39.6 39.6a8.03 8.03 0 0 0 0 11.3l67.9 67.9c3.1 3.1 8.1 3.1 11.3 0m669.6-79.2l-39.6-39.6a8.03 8.03 0 0 0-11.3 0l-67.9 67.9a8.03 8.03 0 0 0 0 11.3l39.6 39.6c3.1 3.1 8.2 3.1 11.3 0l67.9-67.9c3.1-3.2 3.1-8.2 0-11.3M484 180h56c4.4 0 8-3.6 8-8V76c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v96c0 4.4 3.6 8 8 8m348 712H192c-17.7 0-32 14.3-32 32v24c0 4.4 3.6 8 8 8h688c4.4 0 8-3.6 8-8v-24c0-17.7-14.3-32-32-32m-639-96c0 17.7 14.3 32 32 32h574c17.7 0 32-14.3 32-32V563c0-176.2-142.8-319-319-319S193 386.8 193 563zm72-233c0-136.4 110.6-247 247-247s247 110.6 247 247v193H404V585c0-5.5-4.5-10-10-10h-44c-5.5 0-10 4.5-10 10v171h-75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlibabaOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M602.9 669.8c-37.2 2.6-33.6-17.3-11.5-46.2c50.4-67.2 143.7-158.5 147.9-225.2c5.8-86.6-81.3-113.4-171-113.4c-62.4 1.6-127 18.9-171 34.6c-151.6 53.5-246.6 137.5-306.9 232c-62.4 93.4-43 183.2 91.8 185.8c101.8-4.2 170.5-32.5 239.7-68.2c.5 0-192.5 55.1-263.9 14.7c-7.9-4.2-15.7-10-17.8-26.2c0-33.1 54.6-67.7 86.6-78.7v-56.7c64.5 22.6 140.6 16.3 205.7-32c2.1 5.8 4.2 13.1 3.7 21h11c2.6-22.6-12.6-44.6-37.8-46.2c7.3 5.8 12.6 10.5 15.2 14.7l-1 1l-.5.5c-83.9 58.8-165.3 31.5-173.1 29.9l46.7-45.7l-13.1-33.1c92.9-32.5 169.5-56.2 296.9-78.7l-28.5-23l14.7-8.9c75.5 21 126.4 36.7 123.8 76.6c-1 6.8-3.7 14.7-7.9 23.1C660.1 466.1 594 538 567.2 569c-17.3 20.5-34.6 39.4-46.7 58.3c-13.6 19.4-20.5 37.3-21 53.5c2.6 131.8 391.4-61.9 468-112.9c-111.7 47.8-232.9 93.5-364.6 101.9m85-302.9c2.8 5.2 4.1 11.6 4.1 19.1c-.1-6.8-1.4-13.3-4.1-19.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenterOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M264 230h496c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8H264c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m496 424c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8H264c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8zm144 140H120c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-424H120c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeftOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M120 230h496c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8H120c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m0 424h496c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8H120c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m784 140H120c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-424H120c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRightOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M904 158H408c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h496c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 424H408c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h496c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 212H120c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-424H120c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlipayCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M308.6 545.7c-19.8 2-57.1 10.7-77.4 28.6c-61 53-24.5 150 99 150c71.8 0 143.5-45.7 199.8-119c-80.2-38.9-148.1-66.8-221.4-59.6m460.5 67c100.1 33.4 154.7 43 166.7 44.8A445.9 445.9 0 0 0 960 512c0-247.4-200.6-448-448-448S64 264.6 64 512s200.6 448 448 448c155.9 0 293.2-79.7 373.5-200.5c-75.6-29.8-213.6-85-286.8-120.1c-69.9 85.7-160.1 137.8-253.7 137.8c-158.4 0-212.1-138.1-137.2-229c16.3-19.8 44.2-38.7 87.3-49.4c67.5-16.5 175 10.3 275.7 43.4c18.1-33.3 33.4-69.9 44.7-108.9H305.1V402h160v-56.2H271.3v-31.3h193.8v-80.1s0-13.5 13.7-13.5H557v93.6h191.7v31.3H557.1V402h156.4c-15 61.1-37.7 117.4-66.2 166.8c47.5 17.1 90.1 33.3 121.8 43.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlipayCircleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M308.6 545.7c-19.8 2-57.1 10.7-77.4 28.6c-61 53-24.5 150 99 150c71.8 0 143.5-45.7 199.8-119c-80.2-38.9-148.1-66.8-221.4-59.6m460.5 67c100.1 33.4 154.7 43 166.7 44.8A445.9 445.9 0 0 0 960 512c0-247.4-200.6-448-448-448S64 264.6 64 512s200.6 448 448 448c155.9 0 293.2-79.7 373.5-200.5c-75.6-29.8-213.6-85-286.8-120.1c-69.9 85.7-160.1 137.8-253.7 137.8c-158.4 0-212.1-138.1-137.2-229c16.3-19.8 44.2-38.7 87.3-49.4c67.5-16.5 175 10.3 275.7 43.4c18.1-33.3 33.4-69.9 44.7-108.9H305.1V402h160v-56.2H271.3v-31.3h193.8v-80.1s0-13.5 13.7-13.5H557v93.6h191.7v31.3H557.1V402h156.4c-15 61.1-37.7 117.4-66.2 166.8c47.5 17.1 90.1 33.3 121.8 43.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlipayOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M789 610.3c-38.7-12.9-90.7-32.7-148.5-53.6c34.8-60.3 62.5-129 80.7-203.6H530.5v-68.6h233.6v-38.3H530.5V132h-95.4c-16.7 0-16.7 16.5-16.7 16.5v97.8H182.2v38.3h236.3v68.6H223.4v38.3h378.4a667.18 667.18 0 0 1-54.5 132.9c-122.8-40.4-253.8-73.2-336.1-53c-52.6 13-86.5 36.1-106.5 60.3c-91.4 111-25.9 279.6 167.2 279.6C386 811.2 496 747.6 581.2 643C708.3 704 960 808.7 960 808.7V659.4s-31.6-2.5-171-49.1M253.9 746.6c-150.5 0-195-118.3-120.6-183.1c24.8-21.9 70.2-32.6 94.4-35c89.4-8.8 172.2 25.2 269.9 72.8c-68.8 89.5-156.3 145.3-243.7 145.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlipaySquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M308.6 545.7c-19.8 2-57.1 10.7-77.4 28.6c-61 53-24.5 150 99 150c71.8 0 143.5-45.7 199.8-119c-80.2-38.9-148.1-66.8-221.4-59.6M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m29.4 663.2S703 689.4 598.7 639.5C528.8 725.2 438.6 777.3 345 777.3c-158.4 0-212.1-138.1-137.2-229c16.3-19.8 44.2-38.7 87.3-49.4c67.5-16.5 175 10.3 275.7 43.4c18.1-33.3 33.4-69.9 44.7-108.9H305.1V402h160v-56.2H271.3v-31.3h193.8v-80.1s0-13.5 13.7-13.5H557v93.6h191.7v31.3H557.1V402h156.4c-15 61.1-37.7 117.4-66.2 166.8c47.5 17.1 90.1 33.3 121.8 43.9c114.3 38.2 140.2 40.2 140.2 40.2v122.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AliwangwangFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M868.2 377.4c-18.9-45.1-46.3-85.6-81.2-120.6a377.26 377.26 0 0 0-120.5-81.2A375.65 375.65 0 0 0 519 145.8c-41.9 0-82.9 6.7-121.9 20C306 123.3 200.8 120 170.6 120c-2.2 0-7.4 0-9.4.2c-11.9.4-22.8 6.5-29.2 16.4c-6.5 9.9-7.7 22.4-3.4 33.5l64.3 161.6a378.59 378.59 0 0 0-52.8 193.2c0 51.4 10 101 29.8 147.6c18.9 45 46.2 85.6 81.2 120.5c34.7 34.8 75.4 62.1 120.5 81.2C418.3 894 467.9 904 519 904c51.3 0 100.9-10 147.7-29.8c44.9-18.9 85.5-46.3 120.4-81.2c34.7-34.8 62.1-75.4 81.2-120.6a376.5 376.5 0 0 0 29.8-147.6c-.2-51.2-10.1-100.8-29.9-147.4m-325.2 79c0 20.4-16.6 37.1-37.1 37.1c-20.4 0-37.1-16.7-37.1-37.1v-55.1c0-20.4 16.6-37.1 37.1-37.1c20.4 0 37.1 16.6 37.1 37.1zm175.2 0c0 20.4-16.6 37.1-37.1 37.1S644 476.8 644 456.4v-55.1c0-20.4 16.7-37.1 37.1-37.1c20.4 0 37.1 16.6 37.1 37.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AliwangwangOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M868.2 377.4c-18.9-45.1-46.3-85.6-81.2-120.6a377.26 377.26 0 0 0-120.5-81.2A375.65 375.65 0 0 0 519 145.8c-41.9 0-82.9 6.7-121.9 20C306 123.3 200.8 120 170.6 120c-2.2 0-7.4 0-9.4.2c-11.9.4-22.8 6.5-29.2 16.4c-6.5 9.9-7.7 22.4-3.4 33.5l64.3 161.6a378.59 378.59 0 0 0-52.8 193.2c0 51.4 10 101 29.8 147.6c18.9 45 46.2 85.6 81.2 120.5c34.7 34.8 75.4 62.1 120.5 81.2C418.3 894 467.9 904 519 904c51.3 0 100.9-10.1 147.7-29.8c44.9-18.9 85.5-46.3 120.4-81.2c34.7-34.8 62.1-75.4 81.2-120.6a376.5 376.5 0 0 0 29.8-147.6c-.2-51.2-10.1-100.8-29.9-147.4m-66.4 266.5a307.08 307.08 0 0 1-65.9 98c-28.4 28.5-61.3 50.7-97.7 65.9h-.1c-38 16-78.3 24.2-119.9 24.2a306.51 306.51 0 0 1-217.5-90.2c-28.4-28.5-50.6-61.4-65.8-97.8v-.1c-16-37.8-24.1-78.2-24.1-119.9c0-55.4 14.8-109.7 42.8-157l13.2-22.1l-9.5-23.9L206 192c14.9.6 35.9 2.1 59.7 5.6c43.8 6.5 82.5 17.5 114.9 32.6l19 8.9l19.9-6.8c31.5-10.8 64.8-16.2 98.9-16.2a306.51 306.51 0 0 1 217.5 90.2c28.4 28.5 50.6 61.4 65.8 97.8l.1.1l.1.1c16 37.6 24.1 78 24.2 119.8c-.1 41.7-8.3 82-24.3 119.8M681.1 364.2c-20.4 0-37.1 16.7-37.1 37.1v55.1c0 20.4 16.6 37.1 37.1 37.1s37.1-16.7 37.1-37.1v-55.1c0-20.5-16.7-37.1-37.1-37.1m-175.2 0c-20.5 0-37.1 16.7-37.1 37.1v55.1c0 20.4 16.7 37.1 37.1 37.1c20.5 0 37.1-16.7 37.1-37.1v-55.1c0-20.5-16.7-37.1-37.1-37.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AliyunOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M959.2 383.9c-.3-82.1-66.9-148.6-149.1-148.6H575.9l21.6 85.2l201 43.7a42.58 42.58 0 0 1 32.9 39.7c.1.5.1 216.1 0 216.6a42.58 42.58 0 0 1-32.9 39.7l-201 43.7l-21.6 85.3h234.2c82.1 0 148.8-66.5 149.1-148.6zM225.5 660.4a42.58 42.58 0 0 1-32.9-39.7c-.1-.6-.1-216.1 0-216.6c.8-19.4 14.6-35.5 32.9-39.7l201-43.7l21.6-85.2H213.8c-82.1 0-148.8 66.4-149.1 148.6V641c.3 82.1 67 148.6 149.1 148.6H448l-21.6-85.3zm200.9-158.8h171v21.3h-171z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AmazonCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M485 467.5c-11.6 4.9-20.9 12.2-27.8 22c-6.9 9.8-10.4 21.6-10.4 35.5c0 17.8 7.5 31.5 22.4 41.2c14.1 9.1 28.9 11.4 44.4 6.8c17.9-5.2 30-17.9 36.4-38.1c3-9.3 4.5-19.7 4.5-31.3v-50.2c-12.6.4-24.4 1.6-35.5 3.7c-11.1 2.1-22.4 5.6-34 10.4M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m35.8 262.7c-7.2-10.9-20.1-16.4-38.7-16.4c-1.3 0-3 .1-5.3.3c-2.2.2-6.6 1.5-12.9 3.7a79.4 79.4 0 0 0-17.9 9.1c-5.5 3.8-11.5 10-18 18.4c-6.4 8.5-11.5 18.4-15.3 29.8l-94-8.4c0-12.4 2.4-24.7 7-36.9c4.7-12.2 11.8-23.9 21.4-35c9.6-11.2 21.1-21 34.5-29.4c13.4-8.5 29.6-15.2 48.4-20.3c18.9-5.1 39.1-7.6 60.9-7.6c21.3 0 40.6 2.6 57.8 7.7c17.2 5.2 31.1 11.5 41.4 19.1a117 117 0 0 1 25.9 25.7c6.9 9.6 11.7 18.5 14.4 26.7c2.7 8.2 4 15.7 4 22.8v182.5c0 6.4 1.4 13 4.3 19.8c2.9 6.8 6.3 12.8 10.2 18c3.9 5.2 7.9 9.9 12 14.3c4.1 4.3 7.6 7.7 10.6 9.9l4.1 3.4l-72.5 69.4c-8.5-7.7-16.9-15.4-25.2-23.4c-8.3-8-14.5-14-18.5-18.1l-6.1-6.2c-2.4-2.3-5-5.7-8-10.2c-8.1 12.2-18.5 22.8-31.1 31.8c-12.7 9-26.3 15.6-40.7 19.7c-14.5 4.1-29.4 6.5-44.7 7.1c-15.3.6-30-1.5-43.9-6.5c-13.9-5-26.5-11.7-37.6-20.3c-11.1-8.6-19.9-20.2-26.5-35c-6.6-14.8-9.9-31.5-9.9-50.4c0-17.4 3-33.3 8.9-47.7c6-14.5 13.6-26.5 23-36.1c9.4-9.6 20.7-18.2 34-25.7s26.4-13.4 39.2-17.7c12.8-4.2 26.6-7.8 41.5-10.7c14.9-2.9 27.6-4.8 38.2-5.7c10.6-.9 21.2-1.6 31.8-2v-39.4c0-13.5-2.3-23.5-6.7-30.1m180.5 379.6c-2.8 3.3-7.5 7.8-14.1 13.5s-16.8 12.7-30.5 21.1c-13.7 8.4-28.8 16-45 22.9c-16.3 6.9-36.3 12.9-60.1 18c-23.7 5.1-48.2 7.6-73.3 7.6c-25.4 0-50.7-3.2-76.1-9.6c-25.4-6.4-47.6-14.3-66.8-23.7c-19.1-9.4-37.6-20.2-55.1-32.2c-17.6-12.1-31.7-22.9-42.4-32.5c-10.6-9.6-19.6-18.7-26.8-27.1c-1.7-1.9-2.8-3.6-3.2-5.1c-.4-1.5-.3-2.8.3-3.7c.6-.9 1.5-1.6 2.6-2.2a7.42 7.42 0 0 1 7.4.8c40.9 24.2 72.9 41.3 95.9 51.4c82.9 36.4 168 45.7 255.3 27.9c40.5-8.3 82.1-22.2 124.9-41.8c3.2-1.2 6-1.5 8.3-.9c2.3.6 3.5 2.4 3.5 5.4c0 2.8-1.6 6.3-4.8 10.2m59.9-29c-1.8 11.1-4.9 21.6-9.1 31.8c-7.2 17.1-16.3 30-27.1 38.4c-3.6 2.9-6.4 3.8-8.3 2.8c-1.9-1-1.9-3.5 0-7.4c4.5-9.3 9.2-21.8 14.2-37.7c5-15.8 5.7-26 2.1-30.5c-1.1-1.5-2.7-2.6-5-3.6c-2.2-.9-5.1-1.5-8.6-1.9s-6.7-.6-9.4-.8c-2.8-.2-6.5-.2-11.2 0c-4.7.2-8 .4-10.1.6a874.4 874.4 0 0 1-17.1 1.5c-1.3.2-2.7.4-4.1.5c-1.5.1-2.7.2-3.5.3l-2.7.3c-1 .1-1.7.2-2.2.2h-3.2l-1-.2l-.6-.5l-.5-.9c-1.3-3.3 3.7-7.4 15-12.4s22.3-8.1 32.9-9.3c9.8-1.5 21.3-1.5 34.5-.3s21.3 3.7 24.3 7.4c2.3 3.5 2.5 10.7.7 21.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AmazonOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M825 768.9c-3.3-.9-7.3-.4-11.9 1.3c-61.6 28.2-121.5 48.3-179.7 60.2C507.7 856 385.2 842.6 266 790.3c-33.1-14.6-79.1-39.2-138-74a9.36 9.36 0 0 0-5.3-2c-2-.1-3.7.1-5.3.9c-1.6.8-2.8 1.8-3.7 3.1c-.9 1.3-1.1 3.1-.4 5.4c.6 2.2 2.1 4.7 4.6 7.4c10.4 12.2 23.3 25.2 38.6 39s35.6 29.4 60.9 46.8c25.3 17.4 51.8 32.9 79.3 46.4c27.6 13.5 59.6 24.9 96.1 34.1s73 13.8 109.4 13.8c36.2 0 71.4-3.7 105.5-10.9c34.2-7.3 63-15.9 86.5-25.9c23.4-9.9 45-21 64.8-33c19.8-12 34.4-22.2 43.9-30.3c9.5-8.2 16.3-14.6 20.2-19.4c4.6-5.7 6.9-10.6 6.9-14.9c.1-4.5-1.7-7.1-5-7.9M527.4 348.1c-15.2 1.3-33.5 4.1-55 8.3c-21.5 4.1-41.4 9.3-59.8 15.4s-37.2 14.6-56.3 25.4c-19.2 10.8-35.5 23.2-49 37s-24.5 31.1-33.1 52c-8.6 20.8-12.9 43.7-12.9 68.7c0 27.1 4.7 51.2 14.3 72.5c9.5 21.3 22.2 38 38.2 50.4c15.9 12.4 34 22.1 54 29.2c20 7.1 41.2 10.3 63.2 9.4c22-.9 43.5-4.3 64.4-10.3c20.8-5.9 40.4-15.4 58.6-28.3c18.2-12.9 33.1-28.2 44.8-45.7c4.3 6.6 8.1 11.5 11.5 14.7l8.7 8.9c5.8 5.9 14.7 14.6 26.7 26.1c11.9 11.5 24.1 22.7 36.3 33.7l104.4-99.9l-6-4.9c-4.3-3.3-9.4-8-15.2-14.3c-5.8-6.2-11.6-13.1-17.2-20.5c-5.7-7.4-10.6-16.1-14.7-25.9c-4.1-9.8-6.2-19.3-6.2-28.5V258.7c0-10.1-1.9-21-5.7-32.8c-3.9-11.7-10.7-24.5-20.7-38.3c-10-13.8-22.4-26.2-37.2-37c-14.9-10.8-34.7-20-59.6-27.4c-24.8-7.4-52.6-11.1-83.2-11.1c-31.3 0-60.4 3.7-87.6 10.9c-27.1 7.3-50.3 17-69.7 29.2c-19.3 12.2-35.9 26.3-49.7 42.4c-13.8 16.1-24.1 32.9-30.8 50.4c-6.7 17.5-10.1 35.2-10.1 53.1L408 310c5.5-16.4 12.9-30.6 22-42.8c9.2-12.2 17.9-21 25.8-26.5c8-5.5 16.6-9.9 25.7-13.2c9.2-3.3 15.4-5 18.6-5.4c3.2-.3 5.7-.4 7.6-.4c26.7 0 45.2 7.9 55.6 23.6c6.5 9.5 9.7 23.9 9.7 43.3v56.6c-15.2.6-30.4 1.6-45.6 2.9M573.1 500c0 16.6-2.2 31.7-6.5 45c-9.2 29.1-26.7 47.4-52.4 54.8c-22.4 6.6-43.7 3.3-63.9-9.8c-21.5-14-32.2-33.8-32.2-59.3c0-19.9 5-36.9 15-51.1c10-14.1 23.3-24.7 40-31.7s33-12 49-14.9c15.9-3 33-4.8 51-5.4zm335.2 218.9c-4.3-5.4-15.9-8.9-34.9-10.7c-19-1.8-35.5-1.7-49.7.4c-15.3 1.8-31.1 6.2-47.3 13.4c-16.3 7.1-23.4 13.1-21.6 17.8l.7 1.3l.9.7l1.4.2h4.6c.8 0 1.8-.1 3.2-.2c1.4-.1 2.7-.3 3.9-.4c1.2-.1 2.9-.3 5.1-.4c2.1-.1 4.1-.4 6-.7c.3 0 3.7-.3 10.3-.9c6.6-.6 11.4-1 14.3-1.3c2.9-.3 7.8-.6 14.5-.9c6.7-.3 12.1-.3 16.1 0c4 .3 8.5.7 13.6 1.1c5.1.4 9.2 1.3 12.4 2.7c3.2 1.3 5.6 3 7.1 5.1c5.2 6.6 4.2 21.2-3 43.9s-14 40.8-20.4 54.2c-2.8 5.7-2.8 9.2 0 10.7s6.7.1 11.9-4c15.6-12.2 28.6-30.6 39.1-55.3c6.1-14.6 10.5-29.8 13.1-45.7c2.4-15.9 2-26.2-1.3-31"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AmazonSquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M547.8 326.7c-7.2-10.9-20.1-16.4-38.7-16.4c-1.3 0-3 .1-5.3.3c-2.2.2-6.6 1.5-12.9 3.7a79.4 79.4 0 0 0-17.9 9.1c-5.5 3.8-11.5 10-18 18.4c-6.4 8.5-11.5 18.4-15.3 29.8l-94-8.4c0-12.4 2.4-24.7 7-36.9s11.8-23.9 21.4-35c9.6-11.2 21.1-21 34.5-29.4c13.4-8.5 29.6-15.2 48.4-20.3c18.9-5.1 39.1-7.6 60.9-7.6c21.3 0 40.6 2.6 57.8 7.7c17.2 5.2 31.1 11.5 41.4 19.1a117 117 0 0 1 25.9 25.7c6.9 9.6 11.7 18.5 14.4 26.7c2.7 8.2 4 15.7 4 22.8v182.5c0 6.4 1.4 13 4.3 19.8c2.9 6.8 6.3 12.8 10.2 18c3.9 5.2 7.9 9.9 12 14.3c4.1 4.3 7.6 7.7 10.6 9.9l4.1 3.4l-72.5 69.4c-8.5-7.7-16.9-15.4-25.2-23.4c-8.3-8-14.5-14-18.5-18.1l-6.1-6.2c-2.4-2.3-5-5.7-8-10.2c-8.1 12.2-18.5 22.8-31.1 31.8c-12.7 9-26.3 15.6-40.7 19.7c-14.5 4.1-29.4 6.5-44.7 7.1c-15.3.6-30-1.5-43.9-6.5c-13.9-5-26.5-11.7-37.6-20.3c-11.1-8.6-19.9-20.2-26.5-35c-6.6-14.8-9.9-31.5-9.9-50.4c0-17.4 3-33.3 8.9-47.7c6-14.5 13.6-26.5 23-36.1c9.4-9.6 20.7-18.2 34-25.7s26.4-13.4 39.2-17.7c12.8-4.2 26.6-7.8 41.5-10.7c14.9-2.9 27.6-4.8 38.2-5.7c10.6-.9 21.2-1.6 31.8-2v-39.4c0-13.5-2.3-23.5-6.7-30.1m180.5 379.6c-2.8 3.3-7.5 7.8-14.1 13.5s-16.8 12.7-30.5 21.1c-13.7 8.4-28.8 16-45 22.9c-16.3 6.9-36.3 12.9-60.1 18c-23.7 5.1-48.2 7.6-73.3 7.6c-25.4 0-50.7-3.2-76.1-9.6c-25.4-6.4-47.6-14.3-66.8-23.7c-19.1-9.4-37.6-20.2-55.1-32.2c-17.6-12.1-31.7-22.9-42.4-32.5c-10.6-9.6-19.6-18.7-26.8-27.1c-1.7-1.9-2.8-3.6-3.2-5.1c-.4-1.5-.3-2.8.3-3.7c.6-.9 1.5-1.6 2.6-2.2a7.42 7.42 0 0 1 7.4.8c40.9 24.2 72.9 41.3 95.9 51.4c82.9 36.4 168 45.7 255.3 27.9c40.5-8.3 82.1-22.2 124.9-41.8c3.2-1.2 6-1.5 8.3-.9c2.3.6 3.5 2.4 3.5 5.4c0 2.8-1.6 6.3-4.8 10.2m59.9-29c-1.8 11.1-4.9 21.6-9.1 31.8c-7.2 17.1-16.3 30-27.1 38.4c-3.6 2.9-6.4 3.8-8.3 2.8c-1.9-1-1.9-3.5 0-7.4c4.5-9.3 9.2-21.8 14.2-37.7c5-15.8 5.7-26 2.1-30.5c-1.1-1.5-2.7-2.6-5-3.6c-2.2-.9-5.1-1.5-8.6-1.9s-6.7-.6-9.4-.8c-2.8-.2-6.5-.2-11.2 0c-4.7.2-8 .4-10.1.6a874.4 874.4 0 0 1-17.1 1.5c-1.3.2-2.7.4-4.1.5c-1.5.1-2.7.2-3.5.3l-2.7.3c-1 .1-1.7.2-2.2.2h-3.2l-1-.2l-.6-.5l-.5-.9c-1.3-3.3 3.7-7.4 15-12.4s22.3-8.1 32.9-9.3c9.8-1.5 21.3-1.5 34.5-.3s21.3 3.7 24.3 7.4c2.3 3.5 2.5 10.7.7 21.7M485 467.5c-11.6 4.9-20.9 12.2-27.8 22c-6.9 9.8-10.4 21.6-10.4 35.5c0 17.8 7.5 31.5 22.4 41.2c14.1 9.1 28.9 11.4 44.4 6.8c17.9-5.2 30-17.9 36.4-38.1c3-9.3 4.5-19.7 4.5-31.3v-50.2c-12.6.4-24.4 1.6-35.5 3.7c-11.1 2.1-22.4 5.6-34 10.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AndroidFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M270.1 741.7c0 23.4 19.1 42.5 42.6 42.5h48.7v120.4c0 30.5 24.5 55.4 54.6 55.4c30.2 0 54.6-24.8 54.6-55.4V784.1h85v120.4c0 30.5 24.5 55.4 54.6 55.4c30.2 0 54.6-24.8 54.6-55.4V784.1h48.7c23.5 0 42.6-19.1 42.6-42.5V346.4h-486zm357.1-600.1l44.9-65c2.6-3.8 2-8.9-1.5-11.4c-3.5-2.4-8.5-1.2-11.1 2.6l-46.6 67.6c-30.7-12.1-64.9-18.8-100.8-18.8c-35.9 0-70.1 6.7-100.8 18.8l-46.6-67.5c-2.6-3.8-7.6-5.1-11.1-2.6c-3.5 2.4-4.1 7.4-1.5 11.4l44.9 65c-71.4 33.2-121.4 96.1-127.8 169.6h486c-6.6-73.6-56.7-136.5-128-169.7M409.5 244.1a26.9 26.9 0 1 1 26.9-26.9a26.97 26.97 0 0 1-26.9 26.9m208.4 0a26.9 26.9 0 1 1 26.9-26.9a26.97 26.97 0 0 1-26.9 26.9m223.4 100.7c-30.2 0-54.6 24.8-54.6 55.4v216.4c0 30.5 24.5 55.4 54.6 55.4c30.2 0 54.6-24.8 54.6-55.4V400.1c.1-30.6-24.3-55.3-54.6-55.3m-658.6 0c-30.2 0-54.6 24.8-54.6 55.4v216.4c0 30.5 24.5 55.4 54.6 55.4c30.2 0 54.6-24.8 54.6-55.4V400.1c0-30.6-24.5-55.3-54.6-55.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AndroidOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448.3 225.2c-18.6 0-32 13.4-32 31.9s13.5 31.9 32 31.9c18.6 0 32-13.4 32-31.9c.1-18.4-13.4-31.9-32-31.9m393.9 96.4c-13.8-13.8-32.7-21.5-53.2-21.5c-3.9 0-7.4.4-10.7 1v-1h-3.6c-5.5-30.6-18.6-60.5-38.1-87.4c-18.7-25.7-43-47.9-70.8-64.9l25.1-35.8v-3.3c0-.8.4-2.3.7-3.8c.6-2.4 1.4-5.5 1.4-8.9c0-18.5-13.5-31.9-32-31.9c-9.8 0-19.5 5.7-25.9 15.4l-29.3 42.1c-30-9.8-62.4-15-93.8-15c-31.3 0-63.7 5.2-93.8 15L389 79.4c-6.6-9.6-16.1-15.4-26-15.4c-18.6 0-32 13.4-32 31.9c0 6.2 2.5 12.8 6.7 17.4l22.6 32.3c-28.7 17-53.5 39.4-72.2 65.1c-19.4 26.9-32 56.8-36.7 87.4h-5.5v1c-3.2-.6-6.7-1-10.7-1c-20.3 0-39.2 7.5-53.1 21.3c-13.8 13.8-21.5 32.6-21.5 53v235c0 20.3 7.5 39.1 21.4 52.9c13.8 13.8 32.8 21.5 53.2 21.5c3.9 0 7.4-.4 10.7-1v93.5c0 29.2 23.9 53.1 53.2 53.1H331v58.3c0 20.3 7.5 39.1 21.4 52.9c13.8 13.8 32.8 21.5 53.2 21.5c20.3 0 39.2-7.5 53.1-21.3c13.8-13.8 21.5-32.6 21.5-53v-58.2H544v58.1c0 20.3 7.5 39.1 21.4 52.9c13.8 13.8 32.8 21.5 53.2 21.5c20.4 0 39.2-7.5 53.1-21.6c13.8-13.8 21.5-32.6 21.5-53v-58.2h31.9c29.3 0 53.2-23.8 53.2-53.1v-91.4c3.2.6 6.7 1 10.7 1c20.3 0 39.2-7.5 53.1-21.3c13.8-13.8 21.5-32.6 21.5-53v-235c-.1-20.3-7.6-39-21.4-52.9M246 609.6c0 6.8-3.9 10.6-10.7 10.6c-6.8 0-10.7-3.8-10.7-10.6V374.5c0-6.8 3.9-10.6 10.7-10.6c6.8 0 10.7 3.8 10.7 10.6zm131.1-396.8c37.5-27.3 85.3-42.3 135-42.3s97.5 15.1 135 42.5c32.4 23.7 54.2 54.2 62.7 87.5H314.4c8.5-33.4 30.5-64 62.7-87.7m39.3 674.7c-.6 5.6-4.4 8.7-10.5 8.7c-6.8 0-10.7-3.8-10.7-10.6v-58.2h21.2zm202.3 8.7c-6.8 0-10.7-3.8-10.7-10.6v-58.2h21.2v60.1c-.6 5.6-4.3 8.7-10.5 8.7m95.8-132.6H309.9V364h404.6zm85.2-154c0 6.8-3.9 10.6-10.7 10.6c-6.8 0-10.7-3.8-10.7-10.6V374.5c0-6.8 3.9-10.6 10.7-10.6c6.8 0 10.7 3.8 10.7 10.6zM576.1 225.2c-18.6 0-32 13.4-32 31.9s13.5 31.9 32 31.9c18.6 0 32.1-13.4 32.1-32c-.1-18.6-13.4-31.8-32.1-31.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AntCloudOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M378.9 738c-3.1 0-6.1-.5-8.8-1.5l4.4 30.7h26.3l-15.5-29.9c-2.1.5-4.2.7-6.4.7m421-291.2c-12.6 0-24.8 1.5-36.5 4.2c-21.4-38.4-62.3-64.3-109.3-64.3c-6.9 0-13.6.6-20.2 1.6c-35.4-77.4-113.4-131.1-203.9-131.1c-112.3 0-205.3 82.6-221.6 190.4C127.3 455.5 64 523.8 64 607c0 88.4 71.6 160.1 160 160.2h50l13.2-27.6c-26.2-8.3-43.3-29-39.1-48.8c4.6-21.6 32.8-33.9 63.1-27.5c22.9 4.9 40.4 19.1 45.5 35.1a26.1 26.1 0 0 1 22.1-12.4h.2c-.8-3.2-1.2-6.5-1.2-9.9c0-20.1 14.8-36.7 34.1-39.6v-25.4c0-4.4 3.6-8 8-8s8 3.6 8 8v26.3c4.6 1.2 8.8 3.2 12.6 5.8l19.5-21.4c3-3.3 8-3.5 11.3-.5c3.3 3 3.5 8 .5 11.3l-20 22l-.2.2a40 40 0 0 1-46.9 59.2c-.4 5.6-2.6 10.7-6 14.8l20 38.4H804v-.1c86.5-2.2 156-73 156-160.1c0-88.5-71.7-160.2-160.1-160.2M338.2 737.2l-4.3 30h24.4l-5.9-41.5c-3.5 4.6-8.3 8.5-14.2 11.5M797.5 305a48 48 0 1 0 96 0a48 48 0 1 0-96 0m-65.7 61.3a24 24 0 1 0 48 0a24 24 0 1 0-48 0M303.4 742.9l-11.6 24.3h26l3.5-24.7c-5.7.8-11.7 1-17.9.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AntDesignOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M716.3 313.8c19-18.9 19-49.7 0-68.6l-69.9-69.9l.1.1c-18.5-18.5-50.3-50.3-95.3-95.2c-21.2-20.7-55.5-20.5-76.5.5L80.9 474.2a53.84 53.84 0 0 0 0 76.4L474.6 944a54.14 54.14 0 0 0 76.5 0l165.1-165c19-18.9 19-49.7 0-68.6a48.7 48.7 0 0 0-68.7 0l-125 125.2c-5.2 5.2-13.3 5.2-18.5 0L189.5 521.4c-5.2-5.2-5.2-13.3 0-18.5l314.4-314.2c.4-.4.9-.7 1.3-1.1c5.2-4.1 12.4-3.7 17.2 1.1l125.2 125.1c19 19 49.8 19 68.7 0M408.6 514.4a106.3 106.2 0 1 0 212.6 0a106.3 106.2 0 1 0-212.6 0m536.2-38.6L821.9 353.5c-19-18.9-49.8-18.9-68.7.1a48.4 48.4 0 0 0 0 68.6l83 82.9c5.2 5.2 5.2 13.3 0 18.5l-81.8 81.7a48.4 48.4 0 0 0 0 68.6a48.7 48.7 0 0 0 68.7 0l121.8-121.7a53.93 53.93 0 0 0-.1-76.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ApartmentOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M908 640H804V488c0-4.4-3.6-8-8-8H548v-96h108c8.8 0 16-7.2 16-16V80c0-8.8-7.2-16-16-16H368c-8.8 0-16 7.2-16 16v288c0 8.8 7.2 16 16 16h108v96H228c-4.4 0-8 3.6-8 8v152H116c-8.8 0-16 7.2-16 16v288c0 8.8 7.2 16 16 16h288c8.8 0 16-7.2 16-16V656c0-8.8-7.2-16-16-16H292v-88h440v88H620c-8.8 0-16 7.2-16 16v288c0 8.8 7.2 16 16 16h288c8.8 0 16-7.2 16-16V656c0-8.8-7.2-16-16-16m-564 76v168H176V716zm84-408V140h168v168zm420 576H680V716h168z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ApiFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m917.7 148.8l-42.4-42.4c-1.6-1.6-3.6-2.3-5.7-2.3s-4.1.8-5.7 2.3l-76.1 76.1a199.27 199.27 0 0 0-112.1-34.3c-51.2 0-102.4 19.5-141.5 58.6L432.3 308.7a8.03 8.03 0 0 0 0 11.3L704 591.7c1.6 1.6 3.6 2.3 5.7 2.3c2 0 4.1-.8 5.7-2.3l101.9-101.9c68.9-69 77-175.7 24.3-253.5l76.1-76.1c3.1-3.2 3.1-8.3 0-11.4M578.9 546.7a8.03 8.03 0 0 0-11.3 0L501 613.3L410.7 523l66.7-66.7c3.1-3.1 3.1-8.2 0-11.3L441 408.6a8.03 8.03 0 0 0-11.3 0L363 475.3l-43-43a7.85 7.85 0 0 0-5.7-2.3c-2 0-4.1.8-5.7 2.3L206.8 534.2c-68.9 68.9-77 175.7-24.3 253.5l-76.1 76.1a8.03 8.03 0 0 0 0 11.3l42.4 42.4c1.6 1.6 3.6 2.3 5.7 2.3s4.1-.8 5.7-2.3l76.1-76.1c33.7 22.9 72.9 34.3 112.1 34.3c51.2 0 102.4-19.5 141.5-58.6l101.9-101.9c3.1-3.1 3.1-8.2 0-11.3l-43-43l66.7-66.7c3.1-3.1 3.1-8.2 0-11.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ApiOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m917.7 148.8l-42.4-42.4c-1.6-1.6-3.6-2.3-5.7-2.3s-4.1.8-5.7 2.3l-76.1 76.1a199.27 199.27 0 0 0-112.1-34.3c-51.2 0-102.4 19.5-141.5 58.6L432.3 308.7a8.03 8.03 0 0 0 0 11.3L704 591.7c1.6 1.6 3.6 2.3 5.7 2.3c2 0 4.1-.8 5.7-2.3l101.9-101.9c68.9-69 77-175.7 24.3-253.5l76.1-76.1c3.1-3.2 3.1-8.3 0-11.4M769.1 441.7l-59.4 59.4l-186.8-186.8l59.4-59.4c24.9-24.9 58.1-38.7 93.4-38.7c35.3 0 68.4 13.7 93.4 38.7c24.9 24.9 38.7 58.1 38.7 93.4c0 35.3-13.8 68.4-38.7 93.4m-190.2 105a8.03 8.03 0 0 0-11.3 0L501 613.3L410.7 523l66.7-66.7c3.1-3.1 3.1-8.2 0-11.3L441 408.6a8.03 8.03 0 0 0-11.3 0L363 475.3l-43-43a7.85 7.85 0 0 0-5.7-2.3c-2 0-4.1.8-5.7 2.3L206.8 534.2c-68.9 69-77 175.7-24.3 253.5l-76.1 76.1a8.03 8.03 0 0 0 0 11.3l42.4 42.4c1.6 1.6 3.6 2.3 5.7 2.3s4.1-.8 5.7-2.3l76.1-76.1c33.7 22.9 72.9 34.3 112.1 34.3c51.2 0 102.4-19.5 141.5-58.6l101.9-101.9c3.1-3.1 3.1-8.2 0-11.3l-43-43l66.7-66.7c3.1-3.1 3.1-8.2 0-11.3zM441.7 769.1a131.32 131.32 0 0 1-93.4 38.7c-35.3 0-68.4-13.7-93.4-38.7a131.32 131.32 0 0 1-38.7-93.4c0-35.3 13.7-68.4 38.7-93.4l59.4-59.4l186.8 186.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ApiTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M254.9 582.3c-25 25-38.7 58.1-38.7 93.4s13.8 68.5 38.7 93.4c25 25 58.1 38.7 93.4 38.7c35.3 0 68.5-13.8 93.4-38.7l59.4-59.4l-186.8-186.8zm420.8-366.1c-35.3 0-68.5 13.8-93.4 38.7l-59.4 59.4l186.8 186.8l59.4-59.4c24.9-25 38.7-58.1 38.7-93.4s-13.8-68.5-38.7-93.4c-25-25-58.1-38.7-93.4-38.7"/><path fill="currentColor" d="M578.9 546.7a8.03 8.03 0 0 0-11.3 0L501 613.3L410.7 523l66.7-66.7c3.1-3.1 3.1-8.2 0-11.3L441 408.6a8.03 8.03 0 0 0-11.3 0L363 475.3l-43-43a7.85 7.85 0 0 0-5.7-2.3c-2 0-4.1.8-5.7 2.3L206.8 534.2a199.45 199.45 0 0 0-58.6 140.4c-.2 39.5 11.2 79.1 34.3 113.1l-76.1 76.1a8.03 8.03 0 0 0 0 11.3l42.4 42.4c1.6 1.6 3.6 2.3 5.7 2.3s4.1-.8 5.7-2.3l76.1-76.1c33.7 22.9 72.9 34.3 112.1 34.3c51.2 0 102.4-19.5 141.5-58.6l101.9-101.9c3.1-3.1 3.1-8.2 0-11.3l-43-43l66.7-66.7c3.1-3.1 3.1-8.2 0-11.3zM441.7 769.1a131.32 131.32 0 0 1-93.4 38.7c-35.3 0-68.4-13.7-93.4-38.7c-24.9-24.9-38.7-58.1-38.7-93.4s13.7-68.4 38.7-93.4l59.4-59.4l186.8 186.8zm476-620.3l-42.4-42.4c-1.6-1.6-3.6-2.3-5.7-2.3s-4.1.8-5.7 2.3l-76.1 76.1a199.27 199.27 0 0 0-112.1-34.3c-51.2 0-102.4 19.5-141.5 58.6L432.3 308.7a8.03 8.03 0 0 0 0 11.3L704 591.7c1.6 1.6 3.6 2.3 5.7 2.3c2 0 4.1-.8 5.7-2.3l101.9-101.9c68.9-69 77-175.7 24.3-253.5l76.1-76.1c3.1-3.2 3.1-8.3 0-11.4M769.1 441.7l-59.4 59.4l-186.8-186.8l59.4-59.4c24.9-24.9 58.1-38.7 93.4-38.7s68.4 13.7 93.4 38.7c24.9 24.9 38.7 58.1 38.7 93.4s-13.8 68.4-38.7 93.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M747.4 535.7c-.4-68.2 30.5-119.6 92.9-157.5c-34.9-50-87.7-77.5-157.3-82.8c-65.9-5.2-138 38.4-164.4 38.4c-27.9 0-91.7-36.6-141.9-36.6C273.1 298.8 163 379.8 163 544.6c0 48.7 8.9 99 26.7 150.8c23.8 68.2 109.6 235.3 199.1 232.6c46.8-1.1 79.9-33.2 140.8-33.2c59.1 0 89.7 33.2 141.9 33.2c90.3-1.3 167.9-153.2 190.5-221.6c-121.1-57.1-114.6-167.2-114.6-170.7m-105.1-305c50.7-60.2 46.1-115 44.6-134.7c-44.8 2.6-96.6 30.5-126.1 64.8c-32.5 36.8-51.6 82.3-47.5 133.6c48.4 3.7 92.6-21.2 129-63.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M747.4 535.7c-.4-68.2 30.5-119.6 92.9-157.5c-34.9-50-87.7-77.5-157.3-82.8c-65.9-5.2-138 38.4-164.4 38.4c-27.9 0-91.7-36.6-141.9-36.6C273.1 298.8 163 379.8 163 544.6c0 48.7 8.9 99 26.7 150.8c23.8 68.2 109.6 235.3 199.1 232.6c46.8-1.1 79.9-33.2 140.8-33.2c59.1 0 89.7 33.2 141.9 33.2c90.3-1.3 167.9-153.2 190.5-221.6c-121.1-57.1-114.6-167.2-114.6-170.7m-10.6 267c-14.3 19.9-28.7 35.6-41.9 45.7c-10.5 8-18.6 11.4-24 11.6c-9-.1-17.7-2.3-34.7-8.8c-1.2-.5-2.5-1-4.2-1.6l-4.4-1.7c-17.4-6.7-27.8-10.3-41.1-13.8c-18.6-4.8-37.1-7.4-56.9-7.4c-20.2 0-39.2 2.5-58.1 7.2c-13.9 3.5-25.6 7.4-42.7 13.8c-.7.3-8.1 3.1-10.2 3.9c-3.5 1.3-6.2 2.3-8.7 3.2c-10.4 3.6-17 5.1-22.9 5.2c-.7 0-1.3-.1-1.8-.2c-1.1-.2-2.5-.6-4.1-1.3c-4.5-1.8-9.9-5.1-16-9.8c-14-10.9-29.4-28-45.1-49.9c-27.5-38.6-53.5-89.8-66-125.7c-15.4-44.8-23-87.7-23-128.6c0-60.2 17.8-106 48.4-137.1c26.3-26.6 61.7-41.5 97.8-42.3c5.9.1 14.5 1.5 25.4 4.5c8.6 2.3 18 5.4 30.7 9.9c3.8 1.4 16.9 6.1 18.5 6.7c7.7 2.8 13.5 4.8 19.2 6.6c18.2 5.8 32.3 9 47.6 9c15.5 0 28.8-3.3 47.7-9.8c7.1-2.4 32.9-12 37.5-13.6c25.6-9.1 44.5-14 60.8-15.2c4.8-.4 9.1-.4 13.2-.1c22.7 1.8 42.1 6.3 58.6 13.8c-37.6 43.4-57 96.5-56.9 158.4c-.3 14.7.9 31.7 5.1 51.8c6.4 30.5 18.6 60.7 37.9 89c14.7 21.5 32.9 40.9 54.7 57.8c-11.5 23.7-25.6 48.2-40.4 68.8m-94.5-572c50.7-60.2 46.1-115 44.6-134.7c-44.8 2.6-96.6 30.5-126.1 64.8c-32.5 36.8-51.6 82.3-47.5 133.6c48.4 3.7 92.6-21.2 129-63.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppstoreAddOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 144H160c-8.8 0-16 7.2-16 16v304c0 8.8 7.2 16 16 16h304c8.8 0 16-7.2 16-16V160c0-8.8-7.2-16-16-16m-52 268H212V212h200zm452-268H560c-8.8 0-16 7.2-16 16v304c0 8.8 7.2 16 16 16h304c8.8 0 16-7.2 16-16V160c0-8.8-7.2-16-16-16m-52 268H612V212h200zm52 132H560c-8.8 0-16 7.2-16 16v304c0 8.8 7.2 16 16 16h304c8.8 0 16-7.2 16-16V560c0-8.8-7.2-16-16-16m-52 268H612V612h200zM424 712H296V584c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v128H104c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h128v128c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V776h128c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppstoreFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M864 144H560c-8.8 0-16 7.2-16 16v304c0 8.8 7.2 16 16 16h304c8.8 0 16-7.2 16-16V160c0-8.8-7.2-16-16-16m0 400H560c-8.8 0-16 7.2-16 16v304c0 8.8 7.2 16 16 16h304c8.8 0 16-7.2 16-16V560c0-8.8-7.2-16-16-16M464 144H160c-8.8 0-16 7.2-16 16v304c0 8.8 7.2 16 16 16h304c8.8 0 16-7.2 16-16V160c0-8.8-7.2-16-16-16m0 400H160c-8.8 0-16 7.2-16 16v304c0 8.8 7.2 16 16 16h304c8.8 0 16-7.2 16-16V560c0-8.8-7.2-16-16-16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppstoreOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 144H160c-8.8 0-16 7.2-16 16v304c0 8.8 7.2 16 16 16h304c8.8 0 16-7.2 16-16V160c0-8.8-7.2-16-16-16m-52 268H212V212h200zm452-268H560c-8.8 0-16 7.2-16 16v304c0 8.8 7.2 16 16 16h304c8.8 0 16-7.2 16-16V160c0-8.8-7.2-16-16-16m-52 268H612V212h200zM464 544H160c-8.8 0-16 7.2-16 16v304c0 8.8 7.2 16 16 16h304c8.8 0 16-7.2 16-16V560c0-8.8-7.2-16-16-16m-52 268H212V612h200zm452-268H560c-8.8 0-16 7.2-16 16v304c0 8.8 7.2 16 16 16h304c8.8 0 16-7.2 16-16V560c0-8.8-7.2-16-16-16m-52 268H612V612h200z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppstoreTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M864 144H560c-8.8 0-16 7.2-16 16v304c0 8.8 7.2 16 16 16h304c8.8 0 16-7.2 16-16V160c0-8.8-7.2-16-16-16m-52 268H612V212h200zM464 544H160c-8.8 0-16 7.2-16 16v304c0 8.8 7.2 16 16 16h304c8.8 0 16-7.2 16-16V560c0-8.8-7.2-16-16-16m-52 268H212V612h200zm52-668H160c-8.8 0-16 7.2-16 16v304c0 8.8 7.2 16 16 16h304c8.8 0 16-7.2 16-16V160c0-8.8-7.2-16-16-16m-52 268H212V212h200zm452 132H560c-8.8 0-16 7.2-16 16v304c0 8.8 7.2 16 16 16h304c8.8 0 16-7.2 16-16V560c0-8.8-7.2-16-16-16m-52 268H612V612h200z"/><path fill="currentColor" fill-opacity=".15" d="M212 212h200v200H212zm400 0h200v200H612zM212 612h200v200H212zm400 0h200v200H612z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AreaChartOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M888 792H200V168c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v688c0 4.4 3.6 8 8 8h752c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-616-64h536c4.4 0 8-3.6 8-8V284c0-7.2-8.7-10.7-13.7-5.7L592 488.6l-125.4-124a8.03 8.03 0 0 0-11.3 0l-189 189.6a7.87 7.87 0 0 0-2.3 5.6V720c0 4.4 3.6 8 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M862 465.3h-81c-4.6 0-9 2-12.1 5.5L550 723.1V160c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8v563.1L255.1 470.8c-3-3.5-7.4-5.5-12.1-5.5h-81c-6.8 0-10.5 8.1-6 13.2L487.9 861a31.96 31.96 0 0 0 48.3 0L868 478.5c4.5-5.2.8-13.2-6-13.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M872 474H286.9l350.2-304c5.6-4.9 2.2-14-5.2-14h-88.5c-3.9 0-7.6 1.4-10.5 3.9L155 487.8a31.96 31.96 0 0 0 0 48.3L535.1 866c1.5 1.3 3.3 2 5.2 2h91.5c7.4 0 10.8-9.2 5.2-14L286.9 550H872c4.4 0 8-3.6 8-8v-60c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M869 487.8L491.2 159.9c-2.9-2.5-6.6-3.9-10.5-3.9h-88.5c-7.4 0-10.8 9.2-5.2 14l350.2 304H152c-4.4 0-8 3.6-8 8v60c0 4.4 3.6 8 8 8h585.1L386.9 854c-5.6 4.9-2.2 14 5.2 14h91.5c1.9 0 3.8-.7 5.2-2L869 536.2a32.07 32.07 0 0 0 0-48.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M868 545.5L536.1 163a31.96 31.96 0 0 0-48.3 0L156 545.5a7.97 7.97 0 0 0 6 13.2h81c4.6 0 9-2 12.1-5.5L474 300.9V864c0 4.4 3.6 8 8 8h60c4.4 0 8-3.6 8-8V300.9l218.9 252.3c3 3.5 7.4 5.5 12.1 5.5h81c6.8 0 10.5-8 6-13.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsAltOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m855 160.1l-189.2 23.5c-6.6.8-9.3 8.8-4.7 13.5l54.7 54.7l-153.5 153.5a8.03 8.03 0 0 0 0 11.3l45.1 45.1c3.1 3.1 8.2 3.1 11.3 0l153.6-153.6l54.7 54.7a7.94 7.94 0 0 0 13.5-4.7L863.9 169a7.9 7.9 0 0 0-8.9-8.9M416.6 562.3a8.03 8.03 0 0 0-11.3 0L251.8 715.9l-54.7-54.7a7.94 7.94 0 0 0-13.5 4.7L160.1 855c-.6 5.2 3.7 9.5 8.9 8.9l189.2-23.5c6.6-.8 9.3-8.8 4.7-13.5l-54.7-54.7l153.6-153.6c3.1-3.1 3.1-8.2 0-11.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AudioFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 624c93.9 0 170-75.2 170-168V232c0-92.8-76.1-168-170-168s-170 75.2-170 168v224c0 92.8 76.1 168 170 168m330-170c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8c0 140.3-113.7 254-254 254S258 594.3 258 454c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8c0 168.7 126.6 307.9 290 327.6V884H326.7c-13.7 0-24.7 14.3-24.7 32v36c0 4.4 2.8 8 6.2 8h407.6c3.4 0 6.2-3.6 6.2-8v-36c0-17.7-11-32-24.7-32H548V782.1c165.3-18 294-158 294-328.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AudioMutedOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M682 455V311l-76 76v68c-.1 50.7-42 92.1-94 92c-19.1.1-36.8-5.4-52-15l-54 55c29.1 22.4 65.9 36 106 36c93.8 0 170-75.1 170-168"/><path fill="currentColor" d="M833 446h-60c-4.4 0-8 3.6-8 8c0 140.3-113.7 254-254 254c-63 0-120.7-23-165-61l-54 54c48.9 43.2 110.8 72.3 179 81v102H326c-13.9 0-24.9 14.3-25 32v36c.1 4.4 2.9 8 6 8h408c3.2 0 6-3.6 6-8v-36c0-17.7-11-32-25-32H547V782c165.3-17.9 294-157.9 294-328c0-4.4-3.6-8-8-8m13.1-377.7l-43.5-41.9c-3.1-3-8.1-3-11.2.1l-129 129C634.3 101.2 577 64 511 64c-93.9 0-170 75.3-170 168v224c0 6.7.4 13.3 1.2 19.8l-68 68c-10.5-27.9-16.3-58.2-16.2-89.8c-.2-4.4-3.8-8-8-8h-60c-4.4 0-8 3.6-8 8c0 53 12.5 103 34.6 147.4l-137 137c-3.1 3.1-3.1 8.2 0 11.3l42.7 42.7c3.1 3.1 8.2 3.1 11.3 0L846.2 79.8l.1-.1c3.1-3.2 3-8.3-.2-11.4M417 401V232c0-50.6 41.9-92 94-92c46 0 84.1 32.3 92.3 74.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AudioOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M842 454c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8c0 140.3-113.7 254-254 254S258 594.3 258 454c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8c0 168.7 126.6 307.9 290 327.6V884H326.7c-13.7 0-24.7 14.3-24.7 32v36c0 4.4 2.8 8 6.2 8h407.6c3.4 0 6.2-3.6 6.2-8v-36c0-17.7-11-32-24.7-32H548V782.1c165.3-18 294-158 294-328.1M512 624c93.9 0 170-75.2 170-168V232c0-92.8-76.1-168-170-168s-170 75.2-170 168v224c0 92.8 76.1 168 170 168m-94-392c0-50.6 41.9-92 94-92s94 41.4 94 92v224c0 50.6-41.9 92-94 92s-94-41.4-94-92z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AudioTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M512 552c54.3 0 98-43.2 98-96V232c0-52.8-43.7-96-98-96s-98 43.2-98 96v224c0 52.8 43.7 96 98 96"/><path fill="currentColor" d="M842 454c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8c0 140.3-113.7 254-254 254S258 594.3 258 454c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8c0 168.7 126.6 307.9 290 327.6V884H326.7c-13.7 0-24.7 14.3-24.7 32v36c0 4.4 2.8 8 6.2 8h407.6c3.4 0 6.2-3.6 6.2-8v-36c0-17.7-11-32-24.7-32H548V782.1c165.3-18 294-158 294-328.1"/><path fill="currentColor" d="M512 624c93.9 0 170-75.2 170-168V232c0-92.8-76.1-168-170-168s-170 75.2-170 168v224c0 92.8 76.1 168 170 168m-98-392c0-52.8 43.7-96 98-96s98 43.2 98 96v224c0 52.8-43.7 96-98 96s-98-43.2-98-96z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AuditOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M296 250c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h384c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8zm184 144H296c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h184c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8m-48 458H208V148h560v320c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V108c0-17.7-14.3-32-32-32H168c-17.7 0-32 14.3-32 32v784c0 17.7 14.3 32 32 32h264c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m440-88H728v-36.6c46.3-13.8 80-56.6 80-107.4c0-61.9-50.1-112-112-112s-112 50.1-112 112c0 50.7 33.7 93.6 80 107.4V764H520c-8.8 0-16 7.2-16 16v152c0 8.8 7.2 16 16 16h352c8.8 0 16-7.2 16-16V780c0-8.8-7.2-16-16-16M646 620c0-27.6 22.4-50 50-50s50 22.4 50 50s-22.4 50-50 50s-50-22.4-50-50m180 266H566v-60h260z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BackwardFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M485.6 249.9L198.2 498c-8.3 7.1-8.3 20.8 0 27.9l287.4 248.2c10.7 9.2 26.4.9 26.4-14V263.8c0-14.8-15.7-23.2-26.4-13.9m320 0L518.2 498a18.6 18.6 0 0 0-6.2 14c0 5.2 2.1 10.4 6.2 14l287.4 248.2c10.7 9.2 26.4.9 26.4-14V263.8c0-14.8-15.7-23.2-26.4-13.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BackwardOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M485.6 249.9L198.2 498c-8.3 7.1-8.3 20.8 0 27.9l287.4 248.2c10.7 9.2 26.4.9 26.4-14V263.8c0-14.8-15.7-23.2-26.4-13.9m320 0L518.2 498a18.6 18.6 0 0 0-6.2 14c0 5.2 2.1 10.4 6.2 14l287.4 248.2c10.7 9.2 26.4.9 26.4-14V263.8c0-14.8-15.7-23.2-26.4-13.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BankFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M894 462c30.9 0 43.8-39.7 18.7-58L530.8 126.2a31.81 31.81 0 0 0-37.6 0L111.3 404c-25.1 18.2-12.2 58 18.8 58H192v374h-72c-4.4 0-8 3.6-8 8v52c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-52c0-4.4-3.6-8-8-8h-72V462zM381 836H264V462h117zm189 0H453V462h117zm190 0H642V462h118z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BankOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M894 462c30.9 0 43.8-39.7 18.7-58L530.8 126.2a31.81 31.81 0 0 0-37.6 0L111.3 404c-25.1 18.2-12.2 58 18.8 58H192v374h-72c-4.4 0-8 3.6-8 8v52c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-52c0-4.4-3.6-8-8-8h-72V462zM512 196.7l271.1 197.2H240.9zM264 462h117v374H264zm189 0h117v374H453zm307 374H642V462h118z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BankTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M240.9 393.9h542.2L512 196.7z"/><path fill="currentColor" d="M894 462c30.9 0 43.8-39.7 18.7-58L530.8 126.2a31.81 31.81 0 0 0-37.6 0L111.3 404c-25.1 18.2-12.2 58 18.8 58H192v374h-72c-4.4 0-8 3.6-8 8v52c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-52c0-4.4-3.6-8-8-8h-72V462zM381 836H264V462h117zm189 0H453V462h117zm190 0H642V462h118zM240.9 393.9L512 196.7l271.1 197.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarChartOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M888 792H200V168c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v688c0 4.4 3.6 8 8 8h752c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-600-80h56c4.4 0 8-3.6 8-8V560c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v144c0 4.4 3.6 8 8 8m152 0h56c4.4 0 8-3.6 8-8V384c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v320c0 4.4 3.6 8 8 8m152 0h56c4.4 0 8-3.6 8-8V462c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v242c0 4.4 3.6 8 8 8m152 0h56c4.4 0 8-3.6 8-8V304c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v400c0 4.4 3.6 8 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarcodeOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M120 160H72c-4.4 0-8 3.6-8 8v688c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V168c0-4.4-3.6-8-8-8m833 0h-48c-4.4 0-8 3.6-8 8v688c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V168c0-4.4-3.6-8-8-8M200 736h112c4.4 0 8-3.6 8-8V168c0-4.4-3.6-8-8-8H200c-4.4 0-8 3.6-8 8v560c0 4.4 3.6 8 8 8m321 0h48c4.4 0 8-3.6 8-8V168c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v560c0 4.4 3.6 8 8 8m126 0h178c4.4 0 8-3.6 8-8V168c0-4.4-3.6-8-8-8H647c-4.4 0-8 3.6-8 8v560c0 4.4 3.6 8 8 8m-255 0h48c4.4 0 8-3.6 8-8V168c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v560c0 4.4 3.6 8 8 8m-79 64H201c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h112c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8m257 0h-48c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8m256 0H648c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h178c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8m-385 0h-48c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M912 192H328c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h584c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 284H328c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h584c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 284H328c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h584c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8M104 228a56 56 0 1 0 112 0a56 56 0 1 0-112 0m0 284a56 56 0 1 0 112 0a56 56 0 1 0-112 0m0 284a56 56 0 1 0 112 0a56 56 0 1 0-112 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BehanceCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M420.3 470.3c8.7-6.3 12.9-16.7 12.9-31c.3-6.8-1.1-13.5-4.1-19.6c-2.7-4.9-6.7-9-11.6-11.9a44.8 44.8 0 0 0-16.6-6c-6.4-1.2-12.9-1.8-19.3-1.7h-70.3v79.7h76.1c13.1.1 24.2-3.1 32.9-9.5m11.8 72c-9.8-7.5-22.9-11.2-39.2-11.2h-81.8v94h80.2c7.5 0 14.4-.7 21.1-2.1a50.5 50.5 0 0 0 17.8-7.2c5.1-3.3 9.2-7.8 12.3-13.6c3-5.8 4.5-13.2 4.5-22.1c0-17.7-5-30.2-14.9-37.8M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m86.5 286.9h138.4v33.7H598.5zM512 628.8a89.52 89.52 0 0 1-27 31c-11.8 8.2-24.9 14.2-38.8 17.7a167.4 167.4 0 0 1-44.6 5.7H236V342.1h161c16.3 0 31.1 1.5 44.6 4.3c13.4 2.8 24.8 7.6 34.4 14.1c9.5 6.5 17 15.2 22.3 26c5.2 10.7 7.9 24.1 7.9 40c0 17.2-3.9 31.4-11.7 42.9c-7.9 11.5-19.3 20.8-34.8 28.1c21.1 6 36.6 16.7 46.8 31.7c10.4 15.2 15.5 33.4 15.5 54.8c0 17.4-3.3 32.3-10 44.8M790.8 576H612.4c0 19.4 6.7 38 16.8 48c10.2 9.9 24.8 14.9 43.9 14.9c13.8 0 25.5-3.5 35.5-10.4c9.9-6.9 15.9-14.2 18.1-21.8h59.8c-9.6 29.7-24.2 50.9-44 63.7c-19.6 12.8-43.6 19.2-71.5 19.2c-19.5 0-37-3.2-52.7-9.3c-15.1-5.9-28.7-14.9-39.9-26.5a121.2 121.2 0 0 1-25.1-41.2c-6.1-16.9-9.1-34.7-8.9-52.6c0-18.5 3.1-35.7 9.1-51.7c11.5-31.1 35.4-56 65.9-68.9c16.3-6.8 33.8-10.2 51.5-10c21 0 39.2 4 55 12.2a111.6 111.6 0 0 1 38.6 32.8c10.1 13.7 17.2 29.3 21.7 46.9c4.3 17.3 5.8 35.5 4.6 54.7m-122-95.6c-10.8 0-19.9 1.9-26.9 5.6c-7 3.7-12.8 8.3-17.2 13.6a48.4 48.4 0 0 0-9.1 17.4c-1.6 5.3-2.7 10.7-3.1 16.2H723c-1.6-17.3-7.6-30.1-15.6-39.1c-8.4-8.9-21.9-13.7-38.6-13.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BehanceOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M634 294.3h199.5v48.4H634zM434.1 485.8c44.1-21.1 67.2-53.2 67.2-102.8c0-98.1-73-121.9-157.3-121.9H112v492.4h238.5c89.4 0 173.3-43 173.3-143c0-61.8-29.2-107.5-89.7-124.7M220.2 345.1h101.5c39.1 0 74.2 10.9 74.2 56.3c0 41.8-27.3 58.6-66 58.6H220.2zm115.5 324.8H220.1V534.3H338c47.6 0 77.7 19.9 77.7 70.3c0 49.6-35.9 65.3-80 65.3m575.8-89.5c0-105.5-61.7-193.4-173.3-193.4c-108.5 0-182.3 81.7-182.3 188.8c0 111 69.9 187.2 182.3 187.2c85.1 0 140.2-38.3 166.7-120h-86.3c-9.4 30.5-47.6 46.5-77.3 46.5c-57.4 0-87.4-33.6-87.4-90.7h256.9c.3-5.9.7-12.1.7-18.4M653.9 537c3.1-46.9 34.4-76.2 81.2-76.2c49.2 0 73.8 28.9 78.1 76.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BehanceSquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M598.5 350.9h138.4v33.7H598.5zM512 628.8a89.52 89.52 0 0 1-27 31c-11.8 8.2-24.9 14.2-38.8 17.7a167.4 167.4 0 0 1-44.6 5.7H236V342.1h161c16.3 0 31.1 1.5 44.6 4.3c13.4 2.8 24.8 7.6 34.4 14.1c9.5 6.5 17 15.2 22.3 26c5.2 10.7 7.9 24.1 7.9 40c0 17.2-3.9 31.4-11.7 42.9c-7.9 11.5-19.3 20.8-34.8 28.1c21.1 6 36.6 16.7 46.8 31.7c10.4 15.2 15.5 33.4 15.5 54.8c0 17.4-3.3 32.3-10 44.8M790.8 576H612.4c0 19.4 6.7 38 16.8 48c10.2 9.9 24.8 14.9 43.9 14.9c13.8 0 25.5-3.5 35.5-10.4c9.9-6.9 15.9-14.2 18.1-21.8h59.8c-9.6 29.7-24.2 50.9-44 63.7c-19.6 12.8-43.6 19.2-71.5 19.2c-19.5 0-37-3.2-52.7-9.3c-15.1-5.9-28.7-14.9-39.9-26.5a121.2 121.2 0 0 1-25.1-41.2c-6.1-16.9-9.1-34.7-8.9-52.6c0-18.5 3.1-35.7 9.1-51.7c11.5-31.1 35.4-56 65.9-68.9c16.3-6.8 33.8-10.2 51.5-10c21 0 39.2 4 55 12.2a111.6 111.6 0 0 1 38.6 32.8c10.1 13.7 17.2 29.3 21.7 46.9c4.3 17.3 5.8 35.5 4.6 54.7m-122-95.6c-10.8 0-19.9 1.9-26.9 5.6c-7 3.7-12.8 8.3-17.2 13.6a48.4 48.4 0 0 0-9.1 17.4c-1.6 5.3-2.7 10.7-3.1 16.2H723c-1.6-17.3-7.6-30.1-15.6-39.1c-8.4-8.9-21.9-13.7-38.6-13.7m-248.5-10.1c8.7-6.3 12.9-16.7 12.9-31c.3-6.8-1.1-13.5-4.1-19.6c-2.7-4.9-6.7-9-11.6-11.9a44.8 44.8 0 0 0-16.6-6c-6.4-1.2-12.9-1.8-19.3-1.7h-70.3v79.7h76.1c13.1.1 24.2-3.1 32.9-9.5m11.8 72c-9.8-7.5-22.9-11.2-39.2-11.2h-81.8v94h80.2c7.5 0 14.4-.7 21.1-2.1s12.7-3.8 17.8-7.2c5.1-3.3 9.2-7.8 12.3-13.6c3-5.8 4.5-13.2 4.5-22.1c0-17.7-5-30.2-14.9-37.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BehanceSquareOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M598.5 350.9h138.4v33.7H598.5zM512 628.8a89.52 89.52 0 0 1-27 31c-11.8 8.2-24.9 14.2-38.8 17.7a167.4 167.4 0 0 1-44.6 5.7H236V342.1h161c16.3 0 31.1 1.5 44.6 4.3c13.4 2.8 24.8 7.6 34.4 14.1c9.5 6.5 17 15.2 22.3 26c5.2 10.7 7.9 24.1 7.9 40c0 17.2-3.9 31.4-11.7 42.9c-7.9 11.5-19.3 20.8-34.8 28.1c21.1 6 36.6 16.7 46.8 31.7c10.4 15.2 15.5 33.4 15.5 54.8c0 17.4-3.3 32.3-10 44.8M790.8 576H612.4c0 19.4 6.7 38 16.8 48c10.2 9.9 24.8 14.9 43.9 14.9c13.8 0 25.5-3.5 35.5-10.4c9.9-6.9 15.9-14.2 18.1-21.8h59.8c-9.6 29.7-24.2 50.9-44 63.7c-19.6 12.8-43.6 19.2-71.5 19.2c-19.5 0-37-3.2-52.7-9.3c-15.1-5.9-28.7-14.9-39.9-26.5a121.2 121.2 0 0 1-25.1-41.2c-6.1-16.9-9.1-34.7-8.9-52.6c0-18.5 3.1-35.7 9.1-51.7c11.5-31.1 35.4-56 65.9-68.9c16.3-6.8 33.8-10.2 51.5-10c21 0 39.2 4 55 12.2a111.6 111.6 0 0 1 38.6 32.8c10.1 13.7 17.2 29.3 21.7 46.9c4.3 17.3 5.8 35.5 4.6 54.7m-122-95.6c-10.8 0-19.9 1.9-26.9 5.6c-7 3.7-12.8 8.3-17.2 13.6a48.4 48.4 0 0 0-9.1 17.4c-1.6 5.3-2.7 10.7-3.1 16.2H723c-1.6-17.3-7.6-30.1-15.6-39.1c-8.4-8.9-21.9-13.7-38.6-13.7m-248.5-10.1c8.7-6.3 12.9-16.7 12.9-31c.3-6.8-1.1-13.5-4.1-19.6c-2.7-4.9-6.7-9-11.6-11.9a44.8 44.8 0 0 0-16.6-6c-6.4-1.2-12.9-1.8-19.3-1.7h-70.3v79.7h76.1c13.1.1 24.2-3.1 32.9-9.5m11.8 72c-9.8-7.5-22.9-11.2-39.2-11.2h-81.8v94h80.2c7.5 0 14.4-.7 21.1-2.1s12.7-3.8 17.8-7.2c5.1-3.3 9.2-7.8 12.3-13.6c3-5.8 4.5-13.2 4.5-22.1c0-17.7-5-30.2-14.9-37.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M816 768h-24V428c0-141.1-104.3-257.8-240-277.2V112c0-22.1-17.9-40-40-40s-40 17.9-40 40v38.8C336.3 170.2 232 286.9 232 428v340h-24c-17.7 0-32 14.3-32 32v32c0 4.4 3.6 8 8 8h216c0 61.8 50.2 112 112 112s112-50.2 112-112h216c4.4 0 8-3.6 8-8v-32c0-17.7-14.3-32-32-32M512 888c-26.5 0-48-21.5-48-48h96c0 26.5-21.5 48-48 48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M816 768h-24V428c0-141.1-104.3-257.7-240-277.1V112c0-22.1-17.9-40-40-40s-40 17.9-40 40v38.9c-135.7 19.4-240 136-240 277.1v340h-24c-17.7 0-32 14.3-32 32v32c0 4.4 3.6 8 8 8h216c0 61.8 50.2 112 112 112s112-50.2 112-112h216c4.4 0 8-3.6 8-8v-32c0-17.7-14.3-32-32-32M512 888c-26.5 0-48-21.5-48-48h96c0 26.5-21.5 48-48 48M304 768V428c0-55.6 21.6-107.8 60.9-147.1S456.4 220 512 220c55.6 0 107.8 21.6 147.1 60.9S720 372.4 720 428v340z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M512 220c-55.6 0-107.8 21.6-147.1 60.9S304 372.4 304 428v340h416V428c0-55.6-21.6-107.8-60.9-147.1S567.6 220 512 220m280 208c0-141.1-104.3-257.8-240-277.2v.1c135.7 19.4 240 136 240 277.1M472 150.9v-.1C336.3 170.2 232 286.9 232 428c0-141.1 104.3-257.7 240-277.1"/><path fill="currentColor" d="M816 768h-24V428c0-141.1-104.3-257.7-240-277.1V112c0-22.1-17.9-40-40-40s-40 17.9-40 40v38.9c-135.7 19.4-240 136-240 277.1v340h-24c-17.7 0-32 14.3-32 32v32c0 4.4 3.6 8 8 8h216c0 61.8 50.2 112 112 112s112-50.2 112-112h216c4.4 0 8-3.6 8-8v-32c0-17.7-14.3-32-32-32M512 888c-26.5 0-48-21.5-48-48h96c0 26.5-21.5 48-48 48m208-120H304V428c0-55.6 21.6-107.8 60.9-147.1S456.4 220 512 220c55.6 0 107.8 21.6 147.1 60.9S720 372.4 720 428z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BgColorsOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M766.4 744.3c43.7 0 79.4-36.2 79.4-80.5c0-53.5-79.4-140.8-79.4-140.8S687 610.3 687 663.8c0 44.3 35.7 80.5 79.4 80.5m-377.1-44.1c7.1 7.1 18.6 7.1 25.6 0l256.1-256c7.1-7.1 7.1-18.6 0-25.6l-256-256c-.6-.6-1.3-1.2-2-1.7l-78.2-78.2a9.11 9.11 0 0 0-12.8 0l-48 48a9.11 9.11 0 0 0 0 12.8l67.2 67.2l-207.8 207.9c-7.1 7.1-7.1 18.6 0 25.6zm12.9-448.6l178.9 178.9H223.4zM904 816H120c-4.4 0-8 3.6-8 8v80c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-80c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BlockOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M856 376H648V168c0-8.8-7.2-16-16-16H168c-8.8 0-16 7.2-16 16v464c0 8.8 7.2 16 16 16h208v208c0 8.8 7.2 16 16 16h464c8.8 0 16-7.2 16-16V392c0-8.8-7.2-16-16-16m-480 16v188H220V220h360v156H392c-8.8 0-16 7.2-16 16m204 52v136H444V444zm224 360H444V648h188c8.8 0 16-7.2 16-16V444h156z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoldOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M697.8 481.4c33.6-35 54.2-82.3 54.2-134.3v-10.2C752 229.3 663.9 142 555.3 142H259.4c-15.1 0-27.4 12.3-27.4 27.4v679.1c0 16.3 13.2 29.5 29.5 29.5h318.7c117 0 211.8-94.2 211.8-210.5v-11c0-73-37.4-137.3-94.2-175.1M328 238h224.7c57.1 0 103.3 44.4 103.3 99.3v9.5c0 54.8-46.3 99.3-103.3 99.3H328zm366.6 429.4c0 62.9-51.7 113.9-115.5 113.9H328V542.7h251.1c63.8 0 115.5 51 115.5 113.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 64H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V96c0-17.7-14.3-32-32-32M668 345.9L621.5 312L572 347.4V124h96z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 64H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V96c0-17.7-14.3-32-32-32m-260 72h96v209.9L621.5 312L572 347.4zm220 752H232V136h280v296.9c0 3.3 1 6.6 3 9.3a15.9 15.9 0 0 0 22.3 3.7l83.8-59.9l81.4 59.4c2.7 2 6 3.1 9.4 3.1c8.8 0 16-7.2 16-16V136h64v752z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 64H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V96c0-17.7-14.3-32-32-32m-260 72h96v209.9L621.5 312L572 347.4zM232 888V136h280v296.9c0 3.3 1 6.6 3 9.3a15.9 15.9 0 0 0 22.3 3.7l83.8-59.9l81.4 59.4c2.7 2 6 3.1 9.4 3.1c8.8 0 16-7.2 16-16V136h64v752z"/><path fill="currentColor" fill-opacity=".15" d="M668 345.9V136h-96v211.4l49.5-35.4z"/><path fill="currentColor" fill-opacity=".15" d="M727.9 136v296.5c0 8.8-7.2 16-16 16c-3.4 0-6.7-1.1-9.4-3.1L621.1 386l-83.8 59.9a15.9 15.9 0 0 1-22.3-3.7c-2-2.7-3-6-3-9.3V136H232v752h559.9V136z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderBottomOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M872 808H152c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h720c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-720-94h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m0-498h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m0 332h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m0-166h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m166 166h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m0-332h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m332 0h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m0 332h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m222-72h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-388 72h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m388-404h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-388 72h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m388 426h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-388 72h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m388-404h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-388 72h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderHorizontalOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M540 144h-56c-4.4 0-8 3.6-8 8v720c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V152c0-4.4-3.6-8-8-8m-166 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m498 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-664 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m498 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8M208 310h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m664 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 332h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-664 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 166h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-332h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m664 332h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-332h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8M374 808h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-332h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m332 332h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-332h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderInnerOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M872 476H548V144h-72v332H152c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h324v332h72V548h324c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-166h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 498h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-664h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 498h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8M650 216h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m56 592h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-332 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-56-592h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m-166 0h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m56 592h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-56-426h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m56 260h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderLeftOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M208 144h-56c-4.4 0-8 3.6-8 8v720c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V152c0-4.4-3.6-8-8-8m166 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m498 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-332 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m166 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8M540 310h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m332 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 332h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-332 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 166h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-332h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m332 332h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-332h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8M374 808h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-332h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m332 332h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-332h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderOuterOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656zM484 366h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8M302 548h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m364 0h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m-182 0h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m0 182h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderRightOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M872 144h-56c-4.4 0-8 3.6-8 8v720c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V152c0-4.4-3.6-8-8-8m-166 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-498 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m332 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-166 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m166 166h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-332 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 332h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m332 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 166h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-332h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8M208 808h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-332h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m498 332h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-332h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8M374 808h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-332h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderTopOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M872 144H152c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h720c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8M208 310h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 498h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-332h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 166h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m166-166h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 332h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m332 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-332h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m166 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-332 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m332 332h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-332 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m332-498h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-332 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m332 332h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-332 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderVerticleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M872 476H152c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h720c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-166h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 498h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-664h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 498h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8M650 216h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m56 592h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-332 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-56-592h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m-166 0h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m332 0h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8M208 808h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m332 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8M152 382h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m332 0h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8M208 642h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m332 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderlessTableOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M117 368h231v64H117zm559 0h241v64H676zm-264 0h200v64H412zm0 224h200v64H412zm264 0h241v64H676zm-559 0h231v64H117zm295-160V179h-64v666h64V592zm264-64V179h-64v666h64V432z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxPlotFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M952 224h-52c-4.4 0-8 3.6-8 8v248h-92V304c0-4.4-3.6-8-8-8H448v432h344c4.4 0 8-3.6 8-8V548h92v244c0 4.4 3.6 8 8 8h52c4.4 0 8-3.6 8-8V232c0-4.4-3.6-8-8-8m-728 80v176h-92V232c0-4.4-3.6-8-8-8H72c-4.4 0-8 3.6-8 8v560c0 4.4 3.6 8 8 8h52c4.4 0 8-3.6 8-8V548h92v172c0 4.4 3.6 8 8 8h152V296H232c-4.4 0-8 3.6-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxPlotOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M952 224h-52c-4.4 0-8 3.6-8 8v248h-92V304c0-4.4-3.6-8-8-8H232c-4.4 0-8 3.6-8 8v176h-92V232c0-4.4-3.6-8-8-8H72c-4.4 0-8 3.6-8 8v560c0 4.4 3.6 8 8 8h52c4.4 0 8-3.6 8-8V548h92v172c0 4.4 3.6 8 8 8h560c4.4 0 8-3.6 8-8V548h92v244c0 4.4 3.6 8 8 8h52c4.4 0 8-3.6 8-8V232c0-4.4-3.6-8-8-8M296 368h88v288h-88zm432 288H448V368h280z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxPlotTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M296 368h88v288h-88zm152 0h280v288H448z"/><path fill="currentColor" d="M952 224h-52c-4.4 0-8 3.6-8 8v248h-92V304c0-4.4-3.6-8-8-8H232c-4.4 0-8 3.6-8 8v176h-92V232c0-4.4-3.6-8-8-8H72c-4.4 0-8 3.6-8 8v560c0 4.4 3.6 8 8 8h52c4.4 0 8-3.6 8-8V548h92v172c0 4.4 3.6 8 8 8h560c4.4 0 8-3.6 8-8V548h92v244c0 4.4 3.6 8 8 8h52c4.4 0 8-3.6 8-8V232c0-4.4-3.6-8-8-8M384 656h-88V368h88zm344 0H448V368h280z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BranchesOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M740 161c-61.8 0-112 50.2-112 112c0 50.1 33.1 92.6 78.5 106.9v95.9L320 602.4V318.1c44.2-15 76-56.9 76-106.1c0-61.8-50.2-112-112-112s-112 50.2-112 112c0 49.2 31.8 91 76 106.1V706c-44.2 15-76 56.9-76 106.1c0 61.8 50.2 112 112 112s112-50.2 112-112c0-49.2-31.8-91-76-106.1v-27.8l423.5-138.7a50.52 50.52 0 0 0 34.9-48.2V378.2c42.9-15.8 73.6-57 73.6-105.2c0-61.8-50.2-112-112-112m-504 51a48.01 48.01 0 0 1 96 0a48.01 48.01 0 0 1-96 0m96 600a48.01 48.01 0 0 1-96 0a48.01 48.01 0 0 1 96 0m408-491a48.01 48.01 0 0 1 0-96a48.01 48.01 0 0 1 0 96"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BugFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M304 280h416c4.4 0 8-3.6 8-8c0-40-8.8-76.7-25.9-108.1c-17.2-31.5-42.5-56.8-74-74C596.7 72.8 560 64 520 64h-16c-40 0-76.7 8.8-108.1 25.9c-31.5 17.2-56.8 42.5-74 74C304.8 195.3 296 232 296 272c0 4.4 3.6 8 8 8"/><path fill="currentColor" d="M940 512H792V412c76.8 0 139-62.2 139-139c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8c0 34.8-28.2 63-63 63H232c-34.8 0-63-28.2-63-63c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8c0 76.8 62.2 139 139 139v100H84c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h148v96c0 6.5.2 13 .7 19.3C164.1 728.6 116 796.7 116 876c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8c0-44.2 23.9-82.9 59.6-103.7c6 17.2 13.6 33.6 22.7 49c24.3 41.5 59 76.2 100.5 100.5c28.9 16.9 61 28.8 95.3 34.5c4.4 0 8-3.6 8-8V484c0-4.4 3.6-8 8-8h60c4.4 0 8 3.6 8 8v464.2c0 4.4 3.6 8 8 8c34.3-5.7 66.4-17.6 95.3-34.5c41.5-24.3 76.2-59 100.5-100.5c9.1-15.5 16.7-31.9 22.7-49C812.1 793.1 836 831.8 836 876c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8c0-79.3-48.1-147.4-116.7-176.7c.4-6.4.7-12.8.7-19.3v-96h148c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BugOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M304 280h56c4.4 0 8-3.6 8-8c0-28.3 5.9-53.2 17.1-73.5c10.6-19.4 26-34.8 45.4-45.4C450.9 142 475.7 136 504 136h16c28.3 0 53.2 5.9 73.5 17.1c19.4 10.6 34.8 26 45.4 45.4C650 218.9 656 243.7 656 272c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8c0-40-8.8-76.7-25.9-108.1c-17.2-31.5-42.5-56.8-74-74C596.7 72.8 560 64 520 64h-16c-40 0-76.7 8.8-108.1 25.9c-31.5 17.2-56.8 42.5-74 74C304.8 195.3 296 232 296 272c0 4.4 3.6 8 8 8"/><path fill="currentColor" d="M940 512H792V412c76.8 0 139-62.2 139-139c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8c0 34.8-28.2 63-63 63H232c-34.8 0-63-28.2-63-63c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8c0 76.8 62.2 139 139 139v100H84c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h148v96c0 6.5.2 13 .7 19.3C164.1 728.6 116 796.7 116 876c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8c0-44.2 23.9-82.9 59.6-103.7c6 17.2 13.6 33.6 22.7 49c24.3 41.5 59 76.2 100.5 100.5S460.5 960 512 960s99.8-13.9 141.3-38.2c41.5-24.3 76.2-59 100.5-100.5c9.1-15.5 16.7-31.9 22.7-49C812.1 793.1 836 831.8 836 876c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8c0-79.3-48.1-147.4-116.7-176.7c.4-6.4.7-12.8.7-19.3v-96h148c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8M716 680c0 36.8-9.7 72-27.8 102.9c-17.7 30.3-43 55.6-73.3 73.3C584 874.3 548.8 884 512 884s-72-9.7-102.9-27.8c-30.3-17.7-55.6-43-73.3-73.3C317.7 752 308 716.8 308 680V412h408z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BugTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M308 412v268c0 36.779 9.678 71.962 27.805 102.907a205.39 205.39 0 0 0 73.288 73.288C440.038 874.322 475.22 884 512 884c36.779 0 71.962-9.678 102.907-27.805a205.39 205.39 0 0 0 73.288-73.288C706.322 751.962 716 716.78 716 680V412zm484 172v96c0 6.503-.222 12.954-.658 19.346C859.931 728.636 908 796.705 908 876a8 8 0 0 1-8 8h-56a8 8 0 0 1-8-8c0-44.24-23.94-82.888-59.57-103.697a278.635 278.635 0 0 1-22.658 49.018a281.39 281.39 0 0 1-100.45 100.45C611.84 946.072 563.546 960 512 960s-99.84-13.929-141.321-38.228a281.39 281.39 0 0 1-100.45-100.45a278.635 278.635 0 0 1-22.658-49.019C211.94 793.113 188 831.76 188 876a8 8 0 0 1-8 8h-56a8 8 0 0 1-8-8c0-79.295 48.07-147.363 116.658-176.654A284.122 284.122 0 0 1 232 680v-96H84a8 8 0 0 1-8-8v-56a8 8 0 0 1 8-8h148V412c-76.768 0-139-62.232-139-139a8 8 0 0 1 8-8h60a8 8 0 0 1 8 8c0 34.794 28.206 63 63 63h560c34.794 0 63-28.206 63-63a8 8 0 0 1 8-8h60a8 8 0 0 1 8 8c0 76.768-62.232 139-139 139v100h148a8 8 0 0 1 8 8v56a8 8 0 0 1-8 8zM368 272a8 8 0 0 1-8 8h-56a8 8 0 0 1-8-8c0-40.039 8.779-76.746 25.9-108.068c17.235-31.526 42.506-56.797 74.032-74.031C427.254 72.779 463.962 64 504 64h16c40.039 0 76.746 8.779 108.068 25.9c31.526 17.235 56.797 42.506 74.031 74.032C719.221 195.254 728 231.962 728 272a8 8 0 0 1-8 8h-56a8 8 0 0 1-8-8c0-28.326-5.938-53.154-17.077-73.531c-10.625-19.437-25.955-34.767-45.392-45.392C573.154 141.937 548.326 136 520 136h-16c-28.326 0-53.154 5.938-73.531 17.077c-19.437 10.625-34.767 25.955-45.392 45.392C373.937 218.846 368 243.674 368 272"/><path fill="currentColor" fill-opacity=".15" d="M308 412v268c0 36.779 9.678 71.962 27.805 102.907a205.39 205.39 0 0 0 73.288 73.288C440.038 874.322 475.22 884 512 884c36.779 0 71.962-9.678 102.907-27.805a205.39 205.39 0 0 0 73.288-73.288C706.322 751.962 716 716.78 716 680V412z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M916 210H376c-17.7 0-32 14.3-32 32v236H108c-17.7 0-32 14.3-32 32v272c0 17.7 14.3 32 32 32h540c17.7 0 32-14.3 32-32V546h236c17.7 0 32-14.3 32-32V242c0-17.7-14.3-32-32-32M612 746H412V546h200zm268-268H680V278h200z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M916 210H376c-17.7 0-32 14.3-32 32v236H108c-17.7 0-32 14.3-32 32v272c0 17.7 14.3 32 32 32h540c17.7 0 32-14.3 32-32V546h236c17.7 0 32-14.3 32-32V242c0-17.7-14.3-32-32-32m-504 68h200v200H412zm-68 468H144V546h200zm268 0H412V546h200zm268-268H680V278h200z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M144 546h200v200H144zm268-268h200v200H412z"/><path fill="currentColor" d="M916 210H376c-17.7 0-32 14.3-32 32v236H108c-17.7 0-32 14.3-32 32v272c0 17.7 14.3 32 32 32h540c17.7 0 32-14.3 32-32V546h236c17.7 0 32-14.3 32-32V242c0-17.7-14.3-32-32-32M344 746H144V546h200zm268 0H412V546h200zm0-268H412V278h200zm268 0H680V278h200z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BulbFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M348 676.1C250 619.4 184 513.4 184 392c0-181.1 146.9-328 328-328s328 146.9 328 328c0 121.4-66 227.4-164 284.1V792c0 17.7-14.3 32-32 32H380c-17.7 0-32-14.3-32-32zM392 888h240c4.4 0 8 3.6 8 8v32c0 17.7-14.3 32-32 32H416c-17.7 0-32-14.3-32-32v-32c0-4.4 3.6-8 8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BulbOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M632 888H392c-4.4 0-8 3.6-8 8v32c0 17.7 14.3 32 32 32h192c17.7 0 32-14.3 32-32v-32c0-4.4-3.6-8-8-8M512 64c-181.1 0-328 146.9-328 328c0 121.4 66 227.4 164 284.1V792c0 17.7 14.3 32 32 32h264c17.7 0 32-14.3 32-32V676.1c98-56.7 164-162.7 164-284.1c0-181.1-146.9-328-328-328m127.9 549.8L604 634.6V752H420V634.6l-35.9-20.8C305.4 568.3 256 484.5 256 392c0-141.4 114.6-256 256-256s256 114.6 256 256c0 92.5-49.4 176.3-128.1 221.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BulbTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M512 136c-141.4 0-256 114.6-256 256c0 92.5 49.4 176.3 128.1 221.8l35.9 20.8V752h184V634.6l35.9-20.8C718.6 568.3 768 484.5 768 392c0-141.4-114.6-256-256-256"/><path fill="currentColor" d="M632 888H392c-4.4 0-8 3.6-8 8v32c0 17.7 14.3 32 32 32h192c17.7 0 32-14.3 32-32v-32c0-4.4-3.6-8-8-8M512 64c-181.1 0-328 146.9-328 328c0 121.4 66 227.4 164 284.1V792c0 17.7 14.3 32 32 32h264c17.7 0 32-14.3 32-32V676.1c98-56.7 164-162.7 164-284.1c0-181.1-146.9-328-328-328m127.9 549.8L604 634.6V752H420V634.6l-35.9-20.8C305.4 568.3 256 484.5 256 392c0-141.4 114.6-256 256-256s256 114.6 256 256c0 92.5-49.4 176.3-128.1 221.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalculatorFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M440.2 765h-50.8c-2.2 0-4.5-1.1-5.9-2.9L348 718.6l-35.5 43.5a7.38 7.38 0 0 1-5.9 2.9h-50.8c-6.6 0-10.2-7.9-5.8-13.1l62.7-76.8l-61.2-74.9c-4.3-5.2-.7-13.1 5.9-13.1h50.9c2.2 0 4.5 1.1 5.9 2.9l34 41.6l34-41.6c1.5-1.9 3.6-2.9 5.9-2.9h50.8c6.6 0 10.2 7.9 5.9 13.1L383.5 675l62.7 76.8c4.2 5.3.6 13.2-6 13.2m7.8-382c0 2.2-1.4 4-3.2 4H376v68.7c0 1.9-1.8 3.3-4 3.3h-48c-2.2 0-4-1.4-4-3.2V387h-68.8c-1.8 0-3.2-1.8-3.2-4v-48c0-2.2 1.4-4 3.2-4H320v-68.8c0-1.8 1.8-3.2 4-3.2h48c2.2 0 4 1.4 4 3.2V331h68.7c1.9 0 3.3 1.8 3.3 4zm328 369c0 2.2-1.4 4-3.2 4H579.2c-1.8 0-3.2-1.8-3.2-4v-48c0-2.2 1.4-4 3.2-4h193.5c1.9 0 3.3 1.8 3.3 4zm0-104c0 2.2-1.4 4-3.2 4H579.2c-1.8 0-3.2-1.8-3.2-4v-48c0-2.2 1.4-4 3.2-4h193.5c1.9 0 3.3 1.8 3.3 4zm0-265c0 2.2-1.4 4-3.2 4H579.2c-1.8 0-3.2-1.8-3.2-4v-48c0-2.2 1.4-4 3.2-4h193.5c1.9 0 3.3 1.8 3.3 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalculatorOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M251.2 387H320v68.8c0 1.8 1.8 3.2 4 3.2h48c2.2 0 4-1.4 4-3.3V387h68.8c1.8 0 3.2-1.8 3.2-4v-48c0-2.2-1.4-4-3.3-4H376v-68.8c0-1.8-1.8-3.2-4-3.2h-48c-2.2 0-4 1.4-4 3.2V331h-68.8c-1.8 0-3.2 1.8-3.2 4v48c0 2.2 1.4 4 3.2 4m328 0h193.6c1.8 0 3.2-1.8 3.2-4v-48c0-2.2-1.4-4-3.3-4H579.2c-1.8 0-3.2 1.8-3.2 4v48c0 2.2 1.4 4 3.2 4m0 265h193.6c1.8 0 3.2-1.8 3.2-4v-48c0-2.2-1.4-4-3.3-4H579.2c-1.8 0-3.2 1.8-3.2 4v48c0 2.2 1.4 4 3.2 4m0 104h193.6c1.8 0 3.2-1.8 3.2-4v-48c0-2.2-1.4-4-3.3-4H579.2c-1.8 0-3.2 1.8-3.2 4v48c0 2.2 1.4 4 3.2 4m-195.7-81l61.2-74.9c4.3-5.2.7-13.1-5.9-13.1H388c-2.3 0-4.5 1-5.9 2.9l-34 41.6l-34-41.6a7.85 7.85 0 0 0-5.9-2.9h-50.9c-6.6 0-10.2 7.9-5.9 13.1l61.2 74.9l-62.7 76.8c-4.4 5.2-.8 13.1 5.8 13.1h50.8c2.3 0 4.5-1 5.9-2.9l35.5-43.5l35.5 43.5c1.5 1.8 3.7 2.9 5.9 2.9h50.8c6.6 0 10.2-7.9 5.9-13.1zM880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-36 732H180V180h664z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalculatorTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/><path fill="currentColor" fill-opacity=".15" d="M184 840h656V184H184zm256.2-75h-50.8c-2.2 0-4.5-1.1-5.9-2.9L348 718.6l-35.5 43.5a7.38 7.38 0 0 1-5.9 2.9h-50.8c-6.6 0-10.2-7.9-5.8-13.1l62.7-76.8l-61.2-74.9c-4.3-5.2-.7-13.1 5.9-13.1h50.9c2.2 0 4.5 1.1 5.9 2.9l34 41.6l34-41.6c1.5-1.9 3.6-2.9 5.9-2.9h50.8c6.6 0 10.2 7.9 5.9 13.1L383.5 675l62.7 76.8c4.2 5.3.6 13.2-6 13.2M576 335c0-2.2 1.4-4 3.2-4h193.5c1.9 0 3.3 1.8 3.3 4v48c0 2.2-1.4 4-3.2 4H579.2c-1.8 0-3.2-1.8-3.2-4zm0 265c0-2.2 1.4-4 3.2-4h193.5c1.9 0 3.3 1.8 3.3 4v48c0 2.2-1.4 4-3.2 4H579.2c-1.8 0-3.2-1.8-3.2-4zm0 104c0-2.2 1.4-4 3.2-4h193.5c1.9 0 3.3 1.8 3.3 4v48c0 2.2-1.4 4-3.2 4H579.2c-1.8 0-3.2-1.8-3.2-4zM248 335c0-2.2 1.4-4 3.2-4H320v-68.8c0-1.8 1.8-3.2 4-3.2h48c2.2 0 4 1.4 4 3.2V331h68.7c1.9 0 3.3 1.8 3.3 4v48c0 2.2-1.4 4-3.2 4H376v68.7c0 1.9-1.8 3.3-4 3.3h-48c-2.2 0-4-1.4-4-3.2V387h-68.8c-1.8 0-3.2-1.8-3.2-4z"/><path fill="currentColor" d="m383.5 675l61.3-74.8c4.3-5.2.7-13.1-5.9-13.1h-50.8c-2.3 0-4.4 1-5.9 2.9l-34 41.6l-34-41.6a7.69 7.69 0 0 0-5.9-2.9h-50.9c-6.6 0-10.2 7.9-5.9 13.1l61.2 74.9l-62.7 76.8c-4.4 5.2-.8 13.1 5.8 13.1h50.8c2.3 0 4.4-1 5.9-2.9l35.5-43.5l35.5 43.5c1.4 1.8 3.7 2.9 5.9 2.9h50.8c6.6 0 10.2-7.9 6-13.2zM251.2 387H320v68.8c0 1.8 1.8 3.2 4 3.2h48c2.2 0 4-1.4 4-3.3V387h68.8c1.8 0 3.2-1.8 3.2-4v-48c0-2.2-1.4-4-3.3-4H376v-68.8c0-1.8-1.8-3.2-4-3.2h-48c-2.2 0-4 1.4-4 3.2V331h-68.8c-1.8 0-3.2 1.8-3.2 4v48c0 2.2 1.4 4 3.2 4m328 369h193.6c1.8 0 3.2-1.8 3.2-4v-48c0-2.2-1.4-4-3.3-4H579.2c-1.8 0-3.2 1.8-3.2 4v48c0 2.2 1.4 4 3.2 4m0-104h193.6c1.8 0 3.2-1.8 3.2-4v-48c0-2.2-1.4-4-3.3-4H579.2c-1.8 0-3.2 1.8-3.2 4v48c0 2.2 1.4 4 3.2 4m0-265h193.6c1.8 0 3.2-1.8 3.2-4v-48c0-2.2-1.4-4-3.3-4H579.2c-1.8 0-3.2 1.8-3.2 4v48c0 2.2 1.4 4 3.2 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M112 880c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V460H112zm768-696H712v-64c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v64H384v-64c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v64H144c-17.7 0-32 14.3-32 32v176h800V216c0-17.7-14.3-32-32-32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 184H712v-64c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v64H384v-64c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v64H144c-17.7 0-32 14.3-32 32v664c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V216c0-17.7-14.3-32-32-32m-40 656H184V460h656zM184 392V256h128v48c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-48h256v48c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-48h128v136z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M712 304c0 4.4-3.6 8-8 8h-56c-4.4 0-8-3.6-8-8v-48H384v48c0 4.4-3.6 8-8 8h-56c-4.4 0-8-3.6-8-8v-48H184v136h656V256H712z"/><path fill="currentColor" d="M880 184H712v-64c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v64H384v-64c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v64H144c-17.7 0-32 14.3-32 32v664c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V216c0-17.7-14.3-32-32-32m-40 656H184V460h656zm0-448H184V256h128v48c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-48h256v48c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-48h128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M864 260H728l-32.4-90.8a32.07 32.07 0 0 0-30.2-21.2H358.6c-13.5 0-25.6 8.5-30.1 21.2L296 260H160c-44.2 0-80 35.8-80 80v456c0 44.2 35.8 80 80 80h704c44.2 0 80-35.8 80-80V340c0-44.2-35.8-80-80-80M512 716c-88.4 0-160-71.6-160-160s71.6-160 160-160s160 71.6 160 160s-71.6 160-160 160m-96-160a96 96 0 1 0 192 0a96 96 0 1 0-192 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M864 248H728l-32.4-90.8a32.07 32.07 0 0 0-30.2-21.2H358.6c-13.5 0-25.6 8.5-30.1 21.2L296 248H160c-44.2 0-80 35.8-80 80v456c0 44.2 35.8 80 80 80h704c44.2 0 80-35.8 80-80V328c0-44.2-35.8-80-80-80m8 536c0 4.4-3.6 8-8 8H160c-4.4 0-8-3.6-8-8V328c0-4.4 3.6-8 8-8h186.7l17.1-47.8l22.9-64.2h250.5l22.9 64.2l17.1 47.8H864c4.4 0 8 3.6 8 8zM512 384c-88.4 0-160 71.6-160 160s71.6 160 160 160s160-71.6 160-160s-71.6-160-160-160m0 256c-53 0-96-43-96-96s43-96 96-96s96 43 96 96s-43 96-96 96"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M864 320H677.2l-17.1-47.8l-22.9-64.2H386.7l-22.9 64.2l-17.1 47.8H160c-4.4 0-8 3.6-8 8v456c0 4.4 3.6 8 8 8h704c4.4 0 8-3.6 8-8V328c0-4.4-3.6-8-8-8M512 704c-88.4 0-160-71.6-160-160s71.6-160 160-160s160 71.6 160 160s-71.6 160-160 160"/><path fill="currentColor" d="M512 384c-88.4 0-160 71.6-160 160s71.6 160 160 160s160-71.6 160-160s-71.6-160-160-160m0 256c-53 0-96-43-96-96s43-96 96-96s96 43 96 96s-43 96-96 96"/><path fill="currentColor" d="M864 248H728l-32.4-90.8a32.07 32.07 0 0 0-30.2-21.2H358.6c-13.5 0-25.6 8.5-30.1 21.2L296 248H160c-44.2 0-80 35.8-80 80v456c0 44.2 35.8 80 80 80h704c44.2 0 80-35.8 80-80V328c0-44.2-35.8-80-80-80m8 536c0 4.4-3.6 8-8 8H160c-4.4 0-8-3.6-8-8V328c0-4.4 3.6-8 8-8h186.7l17.1-47.8l22.9-64.2h250.5l22.9 64.2l17.1 47.8H864c4.4 0 8 3.6 8 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M959 413.4L935.3 372a8 8 0 0 0-10.9-2.9l-50.7 29.6l-78.3-216.2a63.9 63.9 0 0 0-60.9-44.4H301.2c-34.7 0-65.5 22.4-76.2 55.5l-74.6 205.2l-50.8-29.6a8 8 0 0 0-10.9 2.9L65 413.4c-2.2 3.8-.9 8.6 2.9 10.8l60.4 35.2l-14.5 40c-1.2 3.2-1.8 6.6-1.8 10v348.2c0 15.7 11.8 28.4 26.3 28.4h67.6c12.3 0 23-9.3 25.6-22.3l7.7-37.7h545.6l7.7 37.7c2.7 13 13.3 22.3 25.6 22.3h67.6c14.5 0 26.3-12.7 26.3-28.4V509.4c0-3.4-.6-6.8-1.8-10l-14.5-40l60.3-35.2a8 8 0 0 0 3-10.8M264 621c-22.1 0-40-17.9-40-40s17.9-40 40-40s40 17.9 40 40s-17.9 40-40 40m388 75c0 4.4-3.6 8-8 8H380c-4.4 0-8-3.6-8-8v-84c0-4.4 3.6-8 8-8h40c4.4 0 8 3.6 8 8v36h168v-36c0-4.4 3.6-8 8-8h40c4.4 0 8 3.6 8 8zm108-75c-22.1 0-40-17.9-40-40s17.9-40 40-40s40 17.9 40 40s-17.9 40-40 40M220 418l72.7-199.9l.5-1.3l.4-1.3c1.1-3.3 4.1-5.5 7.6-5.5h427.6l75.4 208z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M380 704h264c4.4 0 8-3.6 8-8v-84c0-4.4-3.6-8-8-8h-40c-4.4 0-8 3.6-8 8v36H428v-36c0-4.4-3.6-8-8-8h-40c-4.4 0-8 3.6-8 8v84c0 4.4 3.6 8 8 8m340-123a40 40 0 1 0 80 0a40 40 0 1 0-80 0m239-167.6L935.3 372a8 8 0 0 0-10.9-2.9l-50.7 29.6l-78.3-216.2a63.9 63.9 0 0 0-60.9-44.4H301.2c-34.7 0-65.5 22.4-76.2 55.5l-74.6 205.2l-50.8-29.6a8 8 0 0 0-10.9 2.9L65 413.4c-2.2 3.8-.9 8.6 2.9 10.8l60.4 35.2l-14.5 40c-1.2 3.2-1.8 6.6-1.8 10v348.2c0 15.7 11.8 28.4 26.3 28.4h67.6c12.3 0 23-9.3 25.6-22.3l7.7-37.7h545.6l7.7 37.7c2.7 13 13.3 22.3 25.6 22.3h67.6c14.5 0 26.3-12.7 26.3-28.4V509.4c0-3.4-.6-6.8-1.8-10l-14.5-40l60.3-35.2a8 8 0 0 0 3-10.8M840 517v237H184V517l15.6-43h624.8zM292.7 218.1l.5-1.3l.4-1.3c1.1-3.3 4.1-5.5 7.6-5.5h427.6l75.4 208H220zM224 581a40 40 0 1 0 80 0a40 40 0 1 0-80 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M199.6 474L184 517v237h656V517l-15.6-43zM264 621c-22.1 0-40-17.9-40-40s17.9-40 40-40s40 17.9 40 40s-17.9 40-40 40m388 75c0 4.4-3.6 8-8 8H380c-4.4 0-8-3.6-8-8v-84c0-4.4 3.6-8 8-8h40c4.4 0 8 3.6 8 8v36h168v-36c0-4.4 3.6-8 8-8h40c4.4 0 8 3.6 8 8zm108-75c-22.1 0-40-17.9-40-40s17.9-40 40-40s40 17.9 40 40s-17.9 40-40 40"/><path fill="currentColor" d="M720 581a40 40 0 1 0 80 0a40 40 0 1 0-80 0"/><path fill="currentColor" d="M959 413.4L935.3 372a8 8 0 0 0-10.9-2.9l-50.7 29.6l-78.3-216.2a63.9 63.9 0 0 0-60.9-44.4H301.2c-34.7 0-65.5 22.4-76.2 55.5l-74.6 205.2l-50.8-29.6a8 8 0 0 0-10.9 2.9L65 413.4c-2.2 3.8-.9 8.6 2.9 10.8l60.4 35.2l-14.5 40c-1.2 3.2-1.8 6.6-1.8 10v348.2c0 15.7 11.8 28.4 26.3 28.4h67.6c12.3 0 23-9.3 25.6-22.3l7.7-37.7h545.6l7.7 37.7c2.7 13 13.3 22.3 25.6 22.3h67.6c14.5 0 26.3-12.7 26.3-28.4V509.4c0-3.4-.6-6.8-1.8-10l-14.5-40l60.3-35.2a8 8 0 0 0 3-10.8M292.7 218.1l.5-1.3l.4-1.3c1.1-3.3 4.1-5.5 7.6-5.5h427.6l75.4 208H220zM840 754H184V517l15.6-43h624.8l15.6 43z"/><path fill="currentColor" d="M224 581a40 40 0 1 0 80 0a40 40 0 1 0-80 0m420 23h-40c-4.4 0-8 3.6-8 8v36H428v-36c0-4.4-3.6-8-8-8h-40c-4.4 0-8 3.6-8 8v84c0 4.4 3.6 8 8 8h264c4.4 0 8-3.6 8-8v-84c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretDownFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M840.4 300H183.6c-19.7 0-30.7 20.8-18.5 35l328.4 380.8c9.4 10.9 27.5 10.9 37 0L858.9 335c12.2-14.2 1.2-35-18.5-35"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretDownOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M840.4 300H183.6c-19.7 0-30.7 20.8-18.5 35l328.4 380.8c9.4 10.9 27.5 10.9 37 0L858.9 335c12.2-14.2 1.2-35-18.5-35"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretLeftFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M689 165.1L308.2 493.5c-10.9 9.4-10.9 27.5 0 37L689 858.9c14.2 12.2 35 1.2 35-18.5V183.6c0-19.7-20.8-30.7-35-18.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretLeftOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M689 165.1L308.2 493.5c-10.9 9.4-10.9 27.5 0 37L689 858.9c14.2 12.2 35 1.2 35-18.5V183.6c0-19.7-20.8-30.7-35-18.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretRightFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M715.8 493.5L335 165.1c-14.2-12.2-35-1.2-35 18.5v656.8c0 19.7 20.8 30.7 35 18.5l380.8-328.4c10.9-9.4 10.9-27.6 0-37"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretRightOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M715.8 493.5L335 165.1c-14.2-12.2-35-1.2-35 18.5v656.8c0 19.7 20.8 30.7 35 18.5l380.8-328.4c10.9-9.4 10.9-27.6 0-37"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretUpFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M858.9 689L530.5 308.2c-9.4-10.9-27.5-10.9-37 0L165.1 689c-12.2 14.2-1.2 35 18.5 35h656.8c19.7 0 30.7-20.8 18.5-35"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretUpOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M858.9 689L530.5 308.2c-9.4-10.9-27.5-10.9-37 0L165.1 689c-12.2 14.2-1.2 35 18.5 35h656.8c19.7 0 30.7-20.8 18.5-35"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarryOutFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 184H712v-64c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v64H384v-64c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v64H144c-17.7 0-32 14.3-32 32v664c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V216c0-17.7-14.3-32-32-32M694.5 432.7L481.9 725.4a16.1 16.1 0 0 1-26 0l-126.4-174c-3.8-5.3 0-12.7 6.5-12.7h55.2c5.1 0 10 2.5 13 6.6l64.7 89l150.9-207.8c3-4.1 7.8-6.6 13-6.6H688c6.5.1 10.3 7.5 6.5 12.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarryOutOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 184H712v-64c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v64H384v-64c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v64H144c-17.7 0-32 14.3-32 32v664c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V216c0-17.7-14.3-32-32-32m-40 656H184V256h128v48c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-48h256v48c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-48h128zM688 420h-55.2c-5.1 0-10 2.5-13 6.6L468.9 634.4l-64.7-89c-3-4.1-7.8-6.6-13-6.6H336c-6.5 0-10.3 7.4-6.5 12.7l126.4 174a16.1 16.1 0 0 0 26 0l212.6-292.7c3.8-5.4 0-12.8-6.5-12.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarryOutTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 184H712v-64c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v64H384v-64c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v64H144c-17.7 0-32 14.3-32 32v664c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V216c0-17.7-14.3-32-32-32m-40 656H184V256h128v48c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-48h256v48c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-48h128z"/><path fill="currentColor" fill-opacity=".15" d="M712 304c0 4.4-3.6 8-8 8h-56c-4.4 0-8-3.6-8-8v-48H384v48c0 4.4-3.6 8-8 8h-56c-4.4 0-8-3.6-8-8v-48H184v584h656V256H712zm-17.5 128.8L481.9 725.5a16.1 16.1 0 0 1-26 0l-126.4-174c-3.8-5.3 0-12.7 6.5-12.7h55.2c5.2 0 10 2.5 13 6.6l64.7 89l150.9-207.8c3-4.1 7.9-6.6 13-6.6H688c6.5 0 10.3 7.4 6.5 12.8"/><path fill="currentColor" d="M688 420h-55.2c-5.1 0-10 2.5-13 6.6L468.9 634.4l-64.7-89c-3-4.1-7.8-6.6-13-6.6H336c-6.5 0-10.3 7.4-6.5 12.7l126.4 174a16.1 16.1 0 0 0 26 0l212.6-292.7c3.8-5.4 0-12.8-6.5-12.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m193.5 301.7l-210.6 292a31.8 31.8 0 0 1-51.7 0L318.5 484.9c-3.8-5.3 0-12.7 6.5-12.7h46.9c10.2 0 19.9 4.9 25.9 13.3l71.2 98.8l157.2-218c6-8.3 15.6-13.3 25.9-13.3H699c6.5 0 10.3 7.4 6.5 12.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M699 353h-46.9c-10.2 0-19.9 4.9-25.9 13.3L469 584.3l-71.2-98.8c-6-8.3-15.6-13.3-25.9-13.3H325c-6.5 0-10.3 7.4-6.5 12.7l124.6 172.8a31.8 31.8 0 0 0 51.7 0l210.6-292c3.9-5.3.1-12.7-6.4-12.7"/><path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircleTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372m193.4 225.7l-210.6 292a31.8 31.8 0 0 1-51.7 0L318.5 484.9c-3.8-5.3 0-12.7 6.5-12.7h46.9c10.3 0 19.9 5 25.9 13.3l71.2 98.8l157.2-218c6-8.4 15.7-13.3 25.9-13.3H699c6.5 0 10.3 7.4 6.4 12.7"/><path fill="currentColor" d="M699 353h-46.9c-10.2 0-19.9 4.9-25.9 13.3L469 584.3l-71.2-98.8c-6-8.3-15.6-13.3-25.9-13.3H325c-6.5 0-10.3 7.4-6.5 12.7l124.6 172.8a31.8 31.8 0 0 0 51.7 0l210.6-292c3.9-5.3.1-12.7-6.4-12.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M912 190h-69.9c-9.8 0-19.1 4.5-25.1 12.2L404.7 724.5L207 474a32 32 0 0 0-25.1-12.2H112c-6.7 0-10.4 7.7-6.3 12.9l273.9 347c12.8 16.2 37.4 16.2 50.3 0l488.4-618.9c4.1-5.1.4-12.8-6.3-12.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckSquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M695.5 365.7l-210.6 292a31.8 31.8 0 0 1-51.7 0L308.5 484.9c-3.8-5.3 0-12.7 6.5-12.7h46.9c10.2 0 19.9 4.9 25.9 13.3l71.2 98.8l157.2-218c6-8.3 15.6-13.3 25.9-13.3H689c6.5 0 10.3 7.4 6.5 12.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckSquareOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M433.1 657.7a31.8 31.8 0 0 0 51.7 0l210.6-292c3.8-5.3 0-12.7-6.5-12.7H642c-10.2 0-19.9 4.9-25.9 13.3L459 584.3l-71.2-98.8c-6-8.3-15.6-13.3-25.9-13.3H315c-6.5 0-10.3 7.4-6.5 12.7z"/><path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckSquareTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/><path fill="currentColor" fill-opacity=".15" d="M184 840h656V184H184zm130-367.8h46.9c10.2 0 19.9 4.9 25.9 13.3l71.2 98.8l157.2-218c6-8.3 15.6-13.3 25.9-13.3H688c6.5 0 10.3 7.4 6.5 12.7l-210.6 292a31.8 31.8 0 0 1-51.7 0L307.5 484.9c-3.8-5.3 0-12.7 6.5-12.7"/><path fill="currentColor" d="M432.2 657.7a31.8 31.8 0 0 0 51.7 0l210.6-292c3.8-5.3 0-12.7-6.5-12.7h-46.9c-10.3 0-19.9 5-25.9 13.3L458 584.3l-71.2-98.8c-6-8.4-15.7-13.3-25.9-13.3H314c-6.5 0-10.3 7.4-6.5 12.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChromeFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M371.8 512c0 77.5 62.7 140.2 140.2 140.2S652.2 589.5 652.2 512S589.5 371.8 512 371.8S371.8 434.4 371.8 512M900 362.4l-234.3 12.1c63.6 74.3 64.6 181.5 11.1 263.7l-188 289.2c78 4.2 158.4-12.9 231.2-55.2c180-104 253-322.1 180-509.8M320.3 591.9L163.8 284.1A415.35 415.35 0 0 0 96 512c0 208 152.3 380.3 351.4 410.8l106.9-209.4c-96.6 18.2-189.9-34.8-234-121.5m218.5-285.5l344.4 18.1C848 254.7 792.6 194 719.8 151.7C653.9 113.6 581.5 95.5 510.5 96c-122.5.5-242.2 55.2-322.1 154.5l128.2 196.9c32-91.9 124.8-146.7 222.2-141"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChromeOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M928 512.3v-.3c0-229.8-186.2-416-416-416S96 282.2 96 512v.4c0 229.8 186.2 416 416 416s416-186.2 416-416v-.3zm-6.7-74.6l.6 3.3zM676.7 638.2c53.5-82.2 52.5-189.4-11.1-263.7l162.4-8.4c20.5 44.4 32 93.8 32 145.9c0 185.2-144.6 336.6-327.1 347.4zM512 652.3c-77.5 0-140.2-62.7-140.2-140.2c0-77.7 62.7-140.2 140.2-140.2S652.2 434.5 652.2 512S589.5 652.3 512 652.3m369.2-331.7l-3-5.7zM512 164c121.3 0 228.2 62.1 290.4 156.2l-263.6-13.9c-97.5-5.7-190.2 49.2-222.3 141.1L227.8 311c63.1-88.9 166.9-147 284.2-147M102.5 585.8c26 145 127.1 264 261.6 315.1C229.6 850 128.5 731 102.5 585.8M164 512c0-55.9 13.2-108.7 36.6-155.5l119.7 235.4c44.1 86.7 137.4 139.7 234 121.6l-74 145.1C302.9 842.5 164 693.5 164 512m324.7 415.4c4 .2 8 .4 12 .5c-4-.2-8-.3-12-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CiCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m-63.6 656c-103 0-162.4-68.6-162.4-182.6v-49C286 373.5 345.4 304 448.3 304c88.3 0 152.3 56.9 152.3 138.1c0 2.4-2 4.4-4.4 4.4h-52.6c-4.2 0-7.6-3.2-8-7.4c-4-46.1-37.6-77.6-87-77.6c-61.1 0-95.6 45.4-95.6 126.9v49.3c0 80.3 34.5 125.1 95.6 125.1c49.3 0 82.8-29.5 87-72.4c.4-4.1 3.8-7.3 8-7.3h52.7c2.4 0 4.4 2 4.4 4.4c0 77.4-64.3 132.5-152.3 132.5M738 704.1c0 4.4-3.6 8-8 8h-50.4c-4.4 0-8-3.6-8-8V319.9c0-4.4 3.6-8 8-8H730c4.4 0 8 3.6 8 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CiCircleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372m218-572.1h-50.4c-4.4 0-8 3.6-8 8v384.2c0 4.4 3.6 8 8 8H730c4.4 0 8-3.6 8-8V319.9c0-4.4-3.6-8-8-8m-281.4 49.6c49.5 0 83.1 31.5 87 77.6c.4 4.2 3.8 7.4 8 7.4h52.6c2.4 0 4.4-2 4.4-4.4c0-81.2-64-138.1-152.3-138.1C345.4 304 286 373.5 286 488.4v49c0 114 59.4 182.6 162.3 182.6c88 0 152.3-55.1 152.3-132.5c0-2.4-2-4.4-4.4-4.4h-52.7c-4.2 0-7.6 3.2-8 7.3c-4.2 43-37.7 72.4-87 72.4c-61.1 0-95.6-44.9-95.6-125.2v-49.3c.1-81.4 34.6-126.8 95.7-126.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CiCircleTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372m-63.5 522.8c49.3 0 82.8-29.4 87-72.4c.4-4.1 3.8-7.3 8-7.3h52.7c2.4 0 4.4 2 4.4 4.4c0 77.4-64.3 132.5-152.3 132.5C345.4 720 286 651.4 286 537.4v-49C286 373.5 345.4 304 448.3 304c88.3 0 152.3 56.9 152.3 138.1c0 2.4-2 4.4-4.4 4.4h-52.6c-4.2 0-7.6-3.2-8-7.4c-3.9-46.1-37.5-77.6-87-77.6c-61.1 0-95.6 45.4-95.7 126.8v49.3c0 80.3 34.5 125.2 95.6 125.2M738 704.1c0 4.4-3.6 8-8 8h-50.4c-4.4 0-8-3.6-8-8V319.9c0-4.4 3.6-8 8-8H730c4.4 0 8 3.6 8 8z"/><path fill="currentColor" d="M730 311.9h-50.4c-4.4 0-8 3.6-8 8v384.2c0 4.4 3.6 8 8 8H730c4.4 0 8-3.6 8-8V319.9c0-4.4-3.6-8-8-8m-281.4 49.6c49.5 0 83.1 31.5 87 77.6c.4 4.2 3.8 7.4 8 7.4h52.6c2.4 0 4.4-2 4.4-4.4c0-81.2-64-138.1-152.3-138.1C345.4 304 286 373.5 286 488.4v49c0 114 59.4 182.6 162.3 182.6c88 0 152.3-55.1 152.3-132.5c0-2.4-2-4.4-4.4-4.4h-52.7c-4.2 0-7.6 3.2-8 7.3c-4.2 43-37.7 72.4-87 72.4c-61.1 0-95.6-44.9-95.6-125.2v-49.3c.1-81.4 34.6-126.8 95.7-126.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CiOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372m218-572.1h-50.4c-4.4 0-8 3.6-8 8v384.2c0 4.4 3.6 8 8 8H730c4.4 0 8-3.6 8-8V319.9c0-4.4-3.6-8-8-8m-281.4 49.6c49.5 0 83.1 31.5 87 77.6c.4 4.2 3.8 7.4 8 7.4h52.6c2.4 0 4.4-2 4.4-4.4c0-81.2-64-138.1-152.3-138.1C345.4 304 286 373.5 286 488.4v49c0 114 59.4 182.6 162.3 182.6c88 0 152.3-55.1 152.3-132.5c0-2.4-2-4.4-4.4-4.4h-52.7c-4.2 0-7.6 3.2-8 7.3c-4.2 43-37.7 72.4-87 72.4c-61.1 0-95.6-44.9-95.6-125.2v-49.3c.1-81.4 34.6-126.8 95.7-126.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CiTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372m-63.5 522.8c49.3 0 82.8-29.4 87-72.4c.4-4.1 3.8-7.3 8-7.3h52.7c2.4 0 4.4 2 4.4 4.4c0 77.4-64.3 132.5-152.3 132.5C345.4 720 286 651.4 286 537.4v-49C286 373.5 345.4 304 448.3 304c88.3 0 152.3 56.9 152.3 138.1c0 2.4-2 4.4-4.4 4.4h-52.6c-4.2 0-7.6-3.2-8-7.4c-3.9-46.1-37.5-77.6-87-77.6c-61.1 0-95.6 45.4-95.7 126.8v49.3c0 80.3 34.5 125.2 95.6 125.2M738 704.1c0 4.4-3.6 8-8 8h-50.4c-4.4 0-8-3.6-8-8V319.9c0-4.4 3.6-8 8-8H730c4.4 0 8 3.6 8 8z"/><path fill="currentColor" d="M730 311.9h-50.4c-4.4 0-8 3.6-8 8v384.2c0 4.4 3.6 8 8 8H730c4.4 0 8-3.6 8-8V319.9c0-4.4-3.6-8-8-8m-281.4 49.6c49.5 0 83.1 31.5 87 77.6c.4 4.2 3.8 7.4 8 7.4h52.6c2.4 0 4.4-2 4.4-4.4c0-81.2-64-138.1-152.3-138.1C345.4 304 286 373.5 286 488.4v49c0 114 59.4 182.6 162.3 182.6c88 0 152.3-55.1 152.3-132.5c0-2.4-2-4.4-4.4-4.4h-52.7c-4.2 0-7.6 3.2-8 7.3c-4.2 43-37.7 72.4-87 72.4c-61.1 0-95.6-44.9-95.6-125.2v-49.3c.1-81.4 34.6-126.8 95.7-126.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClearOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m899.1 869.6l-53-305.6H864c14.4 0 26-11.6 26-26V346c0-14.4-11.6-26-26-26H618V138c0-14.4-11.6-26-26-26H432c-14.4 0-26 11.6-26 26v182H160c-14.4 0-26 11.6-26 26v192c0 14.4 11.6 26 26 26h17.9l-53 305.6c-.3 1.5-.4 3-.4 4.4c0 14.4 11.6 26 26 26h723c1.5 0 3-.1 4.4-.4c14.2-2.4 23.7-15.9 21.2-30M204 390h272V182h72v208h272v104H204zm468 440V674c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v156H416V674c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v156H202.8l45.1-260H776l45.1 260z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m176.5 585.7l-28.6 39a7.99 7.99 0 0 1-11.2 1.7L483.3 569.8a7.92 7.92 0 0 1-3.3-6.5V288c0-4.4 3.6-8 8-8h48.1c4.4 0 8 3.6 8 8v247.5l142.6 103.1c3.6 2.5 4.4 7.5 1.8 11.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockCircleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" d="M686.7 638.6L544.1 535.5V288c0-4.4-3.6-8-8-8H488c-4.4 0-8 3.6-8 8v275.4c0 2.6 1.2 5 3.3 6.5l165.4 120.6c3.6 2.6 8.6 1.8 11.2-1.7l28.6-39c2.6-3.7 1.8-8.7-1.8-11.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockCircleTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372m176.5 509.7l-28.6 39a7.99 7.99 0 0 1-11.2 1.7L483.3 569.8a7.92 7.92 0 0 1-3.3-6.5V288c0-4.4 3.6-8 8-8h48.1c4.4 0 8 3.6 8 8v247.5l142.6 103.1c3.6 2.5 4.4 7.5 1.8 11.1"/><path fill="currentColor" d="M686.7 638.6L544.1 535.5V288c0-4.4-3.6-8-8-8H488c-4.4 0-8 3.6-8 8v275.3c0 2.6 1.2 5 3.3 6.5l165.4 120.6c3.6 2.6 8.6 1.9 11.2-1.7l28.6-39c2.6-3.6 1.8-8.6-1.8-11.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M512 64c247.4 0 448 200.6 448 448S759.4 960 512 960S64 759.4 64 512S264.6 64 512 64m127.978 274.82l-.034.006c-.023.007-.042.018-.083.059L512 466.745l-127.86-127.86c-.042-.041-.06-.052-.084-.059a.118.118 0 0 0-.07 0c-.022.007-.041.018-.082.059l-45.02 45.019c-.04.04-.05.06-.058.083a.118.118 0 0 0 0 .07l.01.022a.268.268 0 0 0 .049.06L466.745 512l-127.86 127.862c-.041.04-.052.06-.059.083a.118.118 0 0 0 0 .07c.007.022.018.041.059.082l45.019 45.02c.04.04.06.05.083.058a.118.118 0 0 0 .07 0c.022-.007.041-.018.082-.059L512 557.254l127.862 127.861c.04.041.06.052.083.059a.118.118 0 0 0 .07 0c.022-.007.041-.018.082-.059l45.02-45.019c.04-.04.05-.06.058-.083a.118.118 0 0 0 0-.07l-.01-.022a.268.268 0 0 0-.049-.06L557.254 512l127.861-127.86c.041-.042.052-.06.059-.084a.118.118 0 0 0 0-.07c-.007-.022-.018-.041-.059-.082l-45.019-45.02a.199.199 0 0 0-.083-.058a.118.118 0 0 0-.07 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseCircleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M512 64c247.4 0 448 200.6 448 448S759.4 960 512 960S64 759.4 64 512S264.6 64 512 64m0 76c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372m128.013 198.826c.023.007.042.018.083.059l45.02 45.019c.04.04.05.06.058.083a.118.118 0 0 1 0 .07c-.007.022-.018.041-.059.082L557.254 512l127.861 127.862a.268.268 0 0 1 .05.06l.009.023a.118.118 0 0 1 0 .07c-.007.022-.018.041-.059.082l-45.019 45.02c-.04.04-.06.05-.083.058a.118.118 0 0 1-.07 0c-.022-.007-.041-.018-.082-.059L512 557.254L384.14 685.115c-.042.041-.06.052-.084.059a.118.118 0 0 1-.07 0c-.022-.007-.041-.018-.082-.059l-45.02-45.019a.199.199 0 0 1-.058-.083a.118.118 0 0 1 0-.07c.007-.022.018-.041.059-.082L466.745 512l-127.86-127.86a.268.268 0 0 1-.05-.061l-.009-.023a.118.118 0 0 1 0-.07c.007-.022.018-.041.059-.082l45.019-45.02c.04-.04.06-.05.083-.058a.118.118 0 0 1 .07 0c.022.007.041.018.082.059L512 466.745l127.862-127.86c.04-.041.06-.052.083-.059a.118.118 0 0 1 .07 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseCircleTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372m171.8 527.1c1.2 1.5 1.9 3.3 1.9 5.2c0 4.5-3.6 8-8 8l-66-.3l-99.3-118.4l-99.3 118.5l-66.1.3c-4.4 0-8-3.6-8-8c0-1.9.7-3.7 1.9-5.2L471 512.3l-130.1-155a8.32 8.32 0 0 1-1.9-5.2c0-4.5 3.6-8 8-8l66.1.3l99.3 118.4l99.4-118.5l66-.3c4.4 0 8 3.6 8 8c0 1.9-.6 3.8-1.8 5.2l-130.1 155z"/><path fill="currentColor" d="M685.8 352c0-4.4-3.6-8-8-8l-66 .3l-99.4 118.5l-99.3-118.4l-66.1-.3c-4.4 0-8 3.5-8 8c0 1.9.7 3.7 1.9 5.2l130.1 155l-130.1 154.9a8.32 8.32 0 0 0-1.9 5.2c0 4.4 3.6 8 8 8l66.1-.3l99.3-118.5L611.7 680l66 .3c4.4 0 8-3.5 8-8c0-1.9-.7-3.7-1.9-5.2L553.9 512.2l130.1-155c1.2-1.4 1.8-3.3 1.8-5.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M799.855 166.312c.023.007.043.018.084.059l57.69 57.69c.041.041.052.06.059.084a.118.118 0 0 1 0 .069c-.007.023-.018.042-.059.083L569.926 512l287.703 287.703c.041.04.052.06.059.083a.118.118 0 0 1 0 .07c-.007.022-.018.042-.059.083l-57.69 57.69c-.041.041-.06.052-.084.059a.118.118 0 0 1-.069 0c-.023-.007-.042-.018-.083-.059L512 569.926L224.297 857.629c-.04.041-.06.052-.083.059a.118.118 0 0 1-.07 0c-.022-.007-.042-.018-.083-.059l-57.69-57.69c-.041-.041-.052-.06-.059-.084a.118.118 0 0 1 0-.069c.007-.023.018-.042.059-.083L454.073 512L166.371 224.297c-.041-.04-.052-.06-.059-.083a.118.118 0 0 1 0-.07c.007-.022.018-.042.059-.083l57.69-57.69c.041-.041.06-.052.084-.059a.118.118 0 0 1 .069 0c.023.007.042.018.083.059L512 454.073l287.703-287.702c.04-.041.06-.052.083-.059a.118.118 0 0 1 .07 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseSquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M880 112c17.7 0 32 14.3 32 32v736c0 17.7-14.3 32-32 32H144c-17.7 0-32-14.3-32-32V144c0-17.7 14.3-32 32-32ZM639.978 338.82l-.034.006c-.023.007-.042.018-.083.059L512 466.745l-127.86-127.86c-.042-.041-.06-.052-.084-.059a.118.118 0 0 0-.07 0c-.022.007-.041.018-.082.059l-45.02 45.019c-.04.04-.05.06-.058.083a.118.118 0 0 0 0 .07l.01.022a.268.268 0 0 0 .049.06L466.745 512l-127.86 127.862c-.041.04-.052.06-.059.083a.118.118 0 0 0 0 .07c.007.022.018.041.059.082l45.019 45.02c.04.04.06.05.083.058a.118.118 0 0 0 .07 0c.022-.007.041-.018.082-.059L512 557.254l127.862 127.861c.04.041.06.052.083.059a.118.118 0 0 0 .07 0c.022-.007.041-.018.082-.059l45.02-45.019c.04-.04.05-.06.058-.083a.118.118 0 0 0 0-.07l-.01-.022a.268.268 0 0 0-.049-.06L557.254 512l127.861-127.86c.041-.042.052-.06.059-.084a.118.118 0 0 0 0-.07c-.007-.022-.018-.041-.059-.082l-45.019-45.02a.199.199 0 0 0-.083-.058a.118.118 0 0 0-.07 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseSquareOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M880 112c17.7 0 32 14.3 32 32v736c0 17.7-14.3 32-32 32H144c-17.7 0-32-14.3-32-32V144c0-17.7 14.3-32 32-32Zm-40 72H184v656h656zM640.013 338.826c.023.007.042.018.083.059l45.02 45.019c.04.04.05.06.058.083a.118.118 0 0 1 0 .07c-.007.022-.018.041-.059.082L557.254 512l127.861 127.862a.268.268 0 0 1 .05.06l.009.023a.118.118 0 0 1 0 .07c-.007.022-.018.041-.059.082l-45.019 45.02c-.04.04-.06.05-.083.058a.118.118 0 0 1-.07 0c-.022-.007-.041-.018-.082-.059L512 557.254L384.14 685.115c-.042.041-.06.052-.084.059a.118.118 0 0 1-.07 0c-.022-.007-.041-.018-.082-.059l-45.02-45.019a.199.199 0 0 1-.058-.083a.118.118 0 0 1 0-.07c.007-.022.018-.041.059-.082L466.745 512l-127.86-127.86a.268.268 0 0 1-.05-.061l-.009-.023a.118.118 0 0 1 0-.07c.007-.022.018-.041.059-.082l45.019-45.02c.04-.04.06-.05.083-.058a.118.118 0 0 1 .07 0c.022.007.041.018.082.059L512 466.745l127.862-127.86c.04-.041.06-.052.083-.059a.118.118 0 0 1 .07 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseSquareTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/><path fill="currentColor" fill-opacity=".15" d="M184 840h656V184H184zm163.9-473.9A7.95 7.95 0 0 1 354 353h58.9c4.7 0 9.2 2.1 12.3 5.7L512 462.2l86.8-103.5c3-3.6 7.5-5.7 12.3-5.7H670c6.8 0 10.5 7.9 6.1 13.1L553.8 512l122.3 145.9c4.4 5.2.7 13.1-6.1 13.1h-58.9c-4.7 0-9.2-2.1-12.3-5.7L512 561.8l-86.8 103.5c-3 3.6-7.5 5.7-12.3 5.7H354c-6.8 0-10.5-7.9-6.1-13.1L470.2 512z"/><path fill="currentColor" d="M354 671h58.9c4.8 0 9.3-2.1 12.3-5.7L512 561.8l86.8 103.5c3.1 3.6 7.6 5.7 12.3 5.7H670c6.8 0 10.5-7.9 6.1-13.1L553.8 512l122.3-145.9c4.4-5.2.7-13.1-6.1-13.1h-58.9c-4.8 0-9.3 2.1-12.3 5.7L512 462.2l-86.8-103.5c-3.1-3.6-7.6-5.7-12.3-5.7H354c-6.8 0-10.5 7.9-6.1 13.1L470.2 512L347.9 657.9A7.95 7.95 0 0 0 354 671"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDownloadOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M624 706.3h-74.1V464c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8v242.3H400c-6.7 0-10.4 7.7-6.3 12.9l112 141.7a8 8 0 0 0 12.6 0l112-141.7c4.1-5.2.4-12.9-6.3-12.9"/><path fill="currentColor" d="M811.4 366.7C765.6 245.9 648.9 160 512.2 160S258.8 245.8 213 366.6C127.3 389.1 64 467.2 64 560c0 110.5 89.5 200 199.9 200H304c4.4 0 8-3.6 8-8v-60c0-4.4-3.6-8-8-8h-40.1c-33.7 0-65.4-13.4-89-37.7c-23.5-24.2-36-56.8-34.9-90.6c.9-26.4 9.9-51.2 26.2-72.1c16.7-21.3 40.1-36.8 66.1-43.7l37.9-9.9l13.9-36.6c8.6-22.8 20.6-44.1 35.7-63.4a245.6 245.6 0 0 1 52.4-49.9c41.1-28.9 89.5-44.2 140-44.2s98.9 15.3 140 44.2c19.9 14 37.5 30.8 52.4 49.9c15.1 19.3 27.1 40.7 35.7 63.4l13.8 36.5l37.8 10C846.1 454.5 884 503.8 884 560c0 33.1-12.9 64.3-36.3 87.7a123.07 123.07 0 0 1-87.6 36.3H720c-4.4 0-8 3.6-8 8v60c0 4.4 3.6 8 8 8h40.1C870.5 760 960 670.5 960 560c0-92.7-63.1-170.7-148.6-193.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M811.4 418.7C765.6 297.9 648.9 212 512.2 212S258.8 297.8 213 418.6C127.3 441.1 64 519.1 64 612c0 110.5 89.5 200 199.9 200h496.2C870.5 812 960 722.5 960 612c0-92.7-63.1-170.7-148.6-193.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M811.4 418.7C765.6 297.9 648.9 212 512.2 212S258.8 297.8 213 418.6C127.3 441.1 64 519.1 64 612c0 110.5 89.5 200 199.9 200h496.2C870.5 812 960 722.5 960 612c0-92.7-63.1-170.7-148.6-193.3m36.3 281a123.07 123.07 0 0 1-87.6 36.3H263.9c-33.1 0-64.2-12.9-87.6-36.3A123.3 123.3 0 0 1 140 612c0-28 9.1-54.3 26.2-76.3a125.7 125.7 0 0 1 66.1-43.7l37.9-9.9l13.9-36.6c8.6-22.8 20.6-44.1 35.7-63.4a245.6 245.6 0 0 1 52.4-49.9c41.1-28.9 89.5-44.2 140-44.2s98.9 15.3 140 44.2c19.9 14 37.5 30.8 52.4 49.9c15.1 19.3 27.1 40.7 35.7 63.4l13.8 36.5l37.8 10c54.3 14.5 92.1 63.8 92.1 120c0 33.1-12.9 64.3-36.3 87.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudServerOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M704 446H320c-4.4 0-8 3.6-8 8v402c0 4.4 3.6 8 8 8h384c4.4 0 8-3.6 8-8V454c0-4.4-3.6-8-8-8m-328 64h272v117H376zm272 290H376V683h272z"/><path fill="currentColor" d="M424 748a32 32 0 1 0 64 0a32 32 0 1 0-64 0m0-178a32 32 0 1 0 64 0a32 32 0 1 0-64 0"/><path fill="currentColor" d="M811.4 368.9C765.6 248 648.9 162 512.2 162S258.8 247.9 213 368.8C126.9 391.5 63.5 470.2 64 563.6C64.6 668 145.6 752.9 247.6 762c4.7.4 8.7-3.3 8.7-8v-60.4c0-4-3-7.4-7-7.9c-27-3.4-52.5-15.2-72.1-34.5c-24-23.5-37.2-55.1-37.2-88.6c0-28 9.1-54.4 26.2-76.4c16.7-21.4 40.2-36.9 66.1-43.7l37.9-10l13.9-36.7c8.6-22.8 20.6-44.2 35.7-63.5c14.9-19.2 32.6-36 52.4-50c41.1-28.9 89.5-44.2 140-44.2s98.9 15.3 140 44.3c19.9 14 37.5 30.8 52.4 50c15.1 19.3 27.1 40.7 35.7 63.5l13.8 36.6l37.8 10c54.2 14.4 92.1 63.7 92.1 120c0 33.6-13.2 65.1-37.2 88.6c-19.5 19.2-44.9 31.1-71.9 34.5c-4 .5-6.9 3.9-6.9 7.9V754c0 4.7 4.1 8.4 8.8 8c101.7-9.2 182.5-94 183.2-198.2c.6-93.4-62.7-172.1-148.6-194.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudSyncOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M811.4 368.9C765.6 248 648.9 162 512.2 162S258.8 247.9 213 368.8C126.9 391.5 63.5 470.2 64 563.6C64.6 668 145.6 752.9 247.6 762c4.7.4 8.7-3.3 8.7-8v-60.4c0-4-3-7.4-7-7.9c-27-3.4-52.5-15.2-72.1-34.5c-24-23.5-37.2-55.1-37.2-88.6c0-28 9.1-54.4 26.2-76.4c16.7-21.4 40.2-36.9 66.1-43.7l37.9-10l13.9-36.7c8.6-22.8 20.6-44.2 35.7-63.5c14.9-19.2 32.6-36 52.4-50c41.1-28.9 89.5-44.2 140-44.2s98.9 15.3 140 44.3c19.9 14 37.5 30.8 52.4 50c15.1 19.3 27.1 40.7 35.7 63.5l13.8 36.6l37.8 10c54.2 14.4 92.1 63.7 92.1 120c0 33.6-13.2 65.1-37.2 88.6c-19.5 19.2-44.9 31.1-71.9 34.5c-4 .5-6.9 3.9-6.9 7.9V754c0 4.7 4.1 8.4 8.8 8c101.7-9.2 182.5-94 183.2-198.2c.6-93.4-62.7-172.1-148.6-194.9"/><path fill="currentColor" d="M376.9 656.4c1.8-33.5 15.7-64.7 39.5-88.6c25.4-25.5 60-39.8 96-39.8c36.2 0 70.3 14.1 96 39.8c1.4 1.4 2.7 2.8 4.1 4.3l-25 19.6a8 8 0 0 0 3 14.1l98.2 24c5 1.2 9.9-2.6 9.9-7.7l.5-101.3c0-6.7-7.6-10.5-12.9-6.3L663 532.7c-36.6-42-90.4-68.6-150.5-68.6c-107.4 0-195 85.1-199.4 191.7c-.2 4.5 3.4 8.3 8 8.3H369c4.2-.1 7.7-3.4 7.9-7.7M703 664h-47.9c-4.2 0-7.7 3.3-8 7.6c-1.8 33.5-15.7 64.7-39.5 88.6c-25.4 25.5-60 39.8-96 39.8c-36.2 0-70.3-14.1-96-39.8c-1.4-1.4-2.7-2.8-4.1-4.3l25-19.6a8 8 0 0 0-3-14.1l-98.2-24c-5-1.2-9.9 2.6-9.9 7.7l-.4 101.4c0 6.7 7.6 10.5 12.9 6.3l23.2-18.2c36.6 42 90.4 68.6 150.5 68.6c107.4 0 195-85.1 199.4-191.7c.2-4.5-3.4-8.3-8-8.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="m791.9 492l-37.8-10l-13.8-36.5c-8.6-22.7-20.6-44.1-35.7-63.4a245.73 245.73 0 0 0-52.4-49.9c-41.1-28.9-89.5-44.2-140-44.2s-98.9 15.3-140 44.2a245.6 245.6 0 0 0-52.4 49.9a240.47 240.47 0 0 0-35.7 63.4l-13.9 36.6l-37.9 9.9a125.7 125.7 0 0 0-66.1 43.7A123.1 123.1 0 0 0 140 612c0 33.1 12.9 64.3 36.3 87.7c23.4 23.4 54.5 36.3 87.6 36.3h496.2c33.1 0 64.2-12.9 87.6-36.3A123.3 123.3 0 0 0 884 612c0-56.2-37.8-105.5-92.1-120"/><path fill="currentColor" d="M811.4 418.7C765.6 297.9 648.9 212 512.2 212S258.8 297.8 213 418.6C127.3 441.1 64 519.1 64 612c0 110.5 89.5 200 199.9 200h496.2C870.5 812 960 722.5 960 612c0-92.7-63.1-170.7-148.6-193.3m36.3 281a123.07 123.07 0 0 1-87.6 36.3H263.9c-33.1 0-64.2-12.9-87.6-36.3A123.3 123.3 0 0 1 140 612c0-28 9.1-54.3 26.2-76.3a125.7 125.7 0 0 1 66.1-43.7l37.9-9.9l13.9-36.6c8.6-22.8 20.6-44.1 35.7-63.4a245.6 245.6 0 0 1 52.4-49.9c41.1-28.9 89.5-44.2 140-44.2s98.9 15.3 140 44.2c19.9 14 37.5 30.8 52.4 49.9c15.1 19.3 27.1 40.7 35.7 63.4l13.8 36.5l37.8 10c54.3 14.5 92.1 63.8 92.1 120c0 33.1-12.9 64.3-36.3 87.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUploadOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M518.3 459a8 8 0 0 0-12.6 0l-112 141.7a7.98 7.98 0 0 0 6.3 12.9h73.9V856c0 4.4 3.6 8 8 8h60c4.4 0 8-3.6 8-8V613.7H624c6.7 0 10.4-7.7 6.3-12.9z"/><path fill="currentColor" d="M811.4 366.7C765.6 245.9 648.9 160 512.2 160S258.8 245.8 213 366.6C127.3 389.1 64 467.2 64 560c0 110.5 89.5 200 199.9 200H304c4.4 0 8-3.6 8-8v-60c0-4.4-3.6-8-8-8h-40.1c-33.7 0-65.4-13.4-89-37.7c-23.5-24.2-36-56.8-34.9-90.6c.9-26.4 9.9-51.2 26.2-72.1c16.7-21.3 40.1-36.8 66.1-43.7l37.9-9.9l13.9-36.6c8.6-22.8 20.6-44.1 35.7-63.4a245.6 245.6 0 0 1 52.4-49.9c41.1-28.9 89.5-44.2 140-44.2s98.9 15.3 140 44.2c19.9 14 37.5 30.8 52.4 49.9c15.1 19.3 27.1 40.7 35.7 63.4l13.8 36.5l37.8 10C846.1 454.5 884 503.8 884 560c0 33.1-12.9 64.3-36.3 87.7a123.07 123.07 0 0 1-87.6 36.3H720c-4.4 0-8 3.6-8 8v60c0 4.4 3.6 8 8 8h40.1C870.5 760 960 670.5 960 560c0-92.7-63.1-170.7-148.6-193.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClusterOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M888 680h-54V540H546v-92h238c8.8 0 16-7.2 16-16V168c0-8.8-7.2-16-16-16H240c-8.8 0-16 7.2-16 16v264c0 8.8 7.2 16 16 16h238v92H190v140h-54c-4.4 0-8 3.6-8 8v176c0 4.4 3.6 8 8 8h176c4.4 0 8-3.6 8-8V688c0-4.4-3.6-8-8-8h-54v-72h220v72h-54c-4.4 0-8 3.6-8 8v176c0 4.4 3.6 8 8 8h176c4.4 0 8-3.6 8-8V688c0-4.4-3.6-8-8-8h-54v-72h220v72h-54c-4.4 0-8 3.6-8 8v176c0 4.4 3.6 8 8 8h176c4.4 0 8-3.6 8-8V688c0-4.4-3.6-8-8-8M256 805.3c0 1.5-1.2 2.7-2.7 2.7h-58.7c-1.5 0-2.7-1.2-2.7-2.7v-58.7c0-1.5 1.2-2.7 2.7-2.7h58.7c1.5 0 2.7 1.2 2.7 2.7zm288 0c0 1.5-1.2 2.7-2.7 2.7h-58.7c-1.5 0-2.7-1.2-2.7-2.7v-58.7c0-1.5 1.2-2.7 2.7-2.7h58.7c1.5 0 2.7 1.2 2.7 2.7zM288 384V216h448v168zm544 421.3c0 1.5-1.2 2.7-2.7 2.7h-58.7c-1.5 0-2.7-1.2-2.7-2.7v-58.7c0-1.5 1.2-2.7 2.7-2.7h58.7c1.5 0 2.7 1.2 2.7 2.7zM360 300a40 40 0 1 0 80 0a40 40 0 1 0-80 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M513.1 518.1l-192 161c-5.2 4.4-13.1.7-13.1-6.1v-62.7c0-2.3 1.1-4.6 2.9-6.1L420.7 512l-109.8-92.2a7.63 7.63 0 0 1-2.9-6.1V351c0-6.8 7.9-10.5 13.1-6.1l192 160.9c3.9 3.2 3.9 9.1 0 12.3M716 673c0 4.4-3.4 8-7.5 8h-185c-4.1 0-7.5-3.6-7.5-8v-48c0-4.4 3.4-8 7.5-8h185c4.1 0 7.5 3.6 7.5 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M516 673c0 4.4 3.4 8 7.5 8h185c4.1 0 7.5-3.6 7.5-8v-48c0-4.4-3.4-8-7.5-8h-185c-4.1 0-7.5 3.6-7.5 8zm-194.9 6.1l192-161c3.8-3.2 3.8-9.1 0-12.3l-192-160.9A7.95 7.95 0 0 0 308 351v62.7c0 2.4 1 4.6 2.9 6.1L420.7 512l-109.8 92.2a8.1 8.1 0 0 0-2.9 6.1V673c0 6.8 7.9 10.5 13.1 6.1M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeSandboxCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m243.7 589.2L512 794L268.3 653.2V371.8l110-63.6l-.4-.2h.2L512 231l134 77h-.2l-.3.2l110.1 63.6v281.4zM307.9 536.7l87.6 49.9V681l96.7 55.9V524.8L307.9 418.4zm203.9-151.8L418 331l-91.1 52.6l185.2 107l185.2-106.9l-91.4-52.8zm20 352l97.3-56.2v-94.1l87-49.5V418.5L531.8 525z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeSandboxOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m709.6 210l.4-.2h.2L512 96L313.9 209.8h-.2l.7.3L151.5 304v416L512 928l360.5-208V304zM482.7 843.6L339.6 761V621.4L210 547.8V372.9l272.7 157.3zM238.2 321.5l134.7-77.8l138.9 79.7l139.1-79.9l135.2 78l-273.9 158zM814 548.3l-128.8 73.1v139.1l-143.9 83V530.4L814 373.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeSandboxSquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m307.9 536.7l87.6 49.9V681l96.7 55.9V524.8L307.9 418.4zM880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M755.7 653.2L512 794L268.3 653.2V371.8l110-63.6l-.4-.2h.2L512 231l134 77h-.2l-.3.2l110.1 63.6v281.4zm-223.9 83.7l97.3-56.2v-94.1l87-49.5V418.5L531.8 525zm-20-352L418 331l-91.1 52.6l185.2 107l185.2-106.9l-91.4-52.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/><path fill="currentColor" fill-opacity=".15" d="M184 840h656V184H184zm339.5-223h185c4.1 0 7.5 3.6 7.5 8v48c0 4.4-3.4 8-7.5 8h-185c-4.1 0-7.5-3.6-7.5-8v-48c0-4.4 3.4-8 7.5-8M308 610.3c0-2.3 1.1-4.6 2.9-6.1L420.7 512l-109.8-92.2a7.63 7.63 0 0 1-2.9-6.1V351c0-6.8 7.9-10.5 13.1-6.1l192 160.9c3.9 3.2 3.9 9.1 0 12.3l-192 161c-5.2 4.4-13.1.7-13.1-6.1z"/><path fill="currentColor" d="m321.1 679.1l192-161c3.9-3.2 3.9-9.1 0-12.3l-192-160.9A7.95 7.95 0 0 0 308 351v62.7c0 2.4 1 4.6 2.9 6.1L420.7 512l-109.8 92.2a8.1 8.1 0 0 0-2.9 6.1V673c0 6.8 7.9 10.5 13.1 6.1M516 673c0 4.4 3.4 8 7.5 8h185c4.1 0 7.5-3.6 7.5-8v-48c0-4.4-3.4-8-7.5-8h-185c-4.1 0-7.5 3.6-7.5 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodepenCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M488.1 414.7V303.4L300.9 428l83.6 55.8zm254.1 137.7v-79.8l-59.8 39.9zM512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m278 533c0 1.1-.1 2.1-.2 3.1c0 .4-.1.7-.2 1a14.16 14.16 0 0 1-.8 3.2c-.2.6-.4 1.2-.6 1.7c-.2.4-.4.8-.5 1.2c-.3.5-.5 1.1-.8 1.6c-.2.4-.4.7-.7 1.1c-.3.5-.7 1-1 1.5c-.3.4-.5.7-.8 1c-.4.4-.8.9-1.2 1.3c-.3.3-.6.6-1 .9c-.4.4-.9.8-1.4 1.1c-.4.3-.7.6-1.1.8c-.1.1-.3.2-.4.3L525.2 786c-4 2.7-8.6 4-13.2 4c-4.7 0-9.3-1.4-13.3-4L244.6 616.9c-.1-.1-.3-.2-.4-.3l-1.1-.8c-.5-.4-.9-.7-1.3-1.1c-.3-.3-.6-.6-1-.9c-.4-.4-.8-.8-1.2-1.3a7 7 0 0 1-.8-1c-.4-.5-.7-1-1-1.5c-.2-.4-.5-.7-.7-1.1c-.3-.5-.6-1.1-.8-1.6c-.2-.4-.4-.8-.5-1.2c-.2-.6-.4-1.2-.6-1.7c-.1-.4-.3-.8-.4-1.2c-.2-.7-.3-1.3-.4-2c-.1-.3-.1-.7-.2-1c-.1-1-.2-2.1-.2-3.1V427.9c0-1 .1-2.1.2-3.1c.1-.3.1-.7.2-1a14.16 14.16 0 0 1 .8-3.2c.2-.6.4-1.2.6-1.7c.2-.4.4-.8.5-1.2c.2-.5.5-1.1.8-1.6c.2-.4.4-.7.7-1.1c.6-.9 1.2-1.7 1.8-2.5c.4-.4.8-.9 1.2-1.3c.3-.3.6-.6 1-.9c.4-.4.9-.8 1.3-1.1c.4-.3.7-.6 1.1-.8c.1-.1.3-.2.4-.3L498.7 239c8-5.3 18.5-5.3 26.5 0l254.1 169.1c.1.1.3.2.4.3l1.1.8l1.4 1.1c.3.3.6.6 1 .9c.4.4.8.8 1.2 1.3c.7.8 1.3 1.6 1.8 2.5c.2.4.5.7.7 1.1c.3.5.6 1 .8 1.6c.2.4.4.8.5 1.2c.2.6.4 1.2.6 1.7c.1.4.3.8.4 1.2c.2.7.3 1.3.4 2c.1.3.1.7.2 1c.1 1 .2 2.1.2 3.1zm-254.1 13.3v111.3L723.1 597l-83.6-55.8zM281.8 472.6v79.8l59.8-39.9zM512 456.1l-84.5 56.4l84.5 56.4l84.5-56.4zM723.1 428L535.9 303.4v111.3l103.6 69.1zM384.5 541.2L300.9 597l187.2 124.6V610.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodepenCircleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M488.1 414.7V303.4L300.9 428l83.6 55.8zm254.1 137.7v-79.8l-59.8 39.9zM512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m278 533c0 1.1-.1 2.1-.2 3.1c0 .4-.1.7-.2 1a14.16 14.16 0 0 1-.8 3.2c-.2.6-.4 1.2-.6 1.7c-.2.4-.4.8-.5 1.2c-.3.5-.5 1.1-.8 1.6c-.2.4-.4.7-.7 1.1c-.3.5-.7 1-1 1.5c-.3.4-.5.7-.8 1c-.4.4-.8.9-1.2 1.3c-.3.3-.6.6-1 .9c-.4.4-.9.8-1.4 1.1c-.4.3-.7.6-1.1.8c-.1.1-.3.2-.4.3L525.2 786c-4 2.7-8.6 4-13.2 4c-4.7 0-9.3-1.4-13.3-4L244.6 616.9c-.1-.1-.3-.2-.4-.3l-1.1-.8c-.5-.4-.9-.7-1.3-1.1c-.3-.3-.6-.6-1-.9c-.4-.4-.8-.8-1.2-1.3a7 7 0 0 1-.8-1c-.4-.5-.7-1-1-1.5c-.2-.4-.5-.7-.7-1.1c-.3-.5-.6-1.1-.8-1.6c-.2-.4-.4-.8-.5-1.2c-.2-.6-.4-1.2-.6-1.7c-.1-.4-.3-.8-.4-1.2c-.2-.7-.3-1.3-.4-2c-.1-.3-.1-.7-.2-1c-.1-1-.2-2.1-.2-3.1V427.9c0-1 .1-2.1.2-3.1c.1-.3.1-.7.2-1a14.16 14.16 0 0 1 .8-3.2c.2-.6.4-1.2.6-1.7c.2-.4.4-.8.5-1.2c.2-.5.5-1.1.8-1.6c.2-.4.4-.7.7-1.1c.6-.9 1.2-1.7 1.8-2.5c.4-.4.8-.9 1.2-1.3c.3-.3.6-.6 1-.9c.4-.4.9-.8 1.3-1.1c.4-.3.7-.6 1.1-.8c.1-.1.3-.2.4-.3L498.7 239c8-5.3 18.5-5.3 26.5 0l254.1 169.1c.1.1.3.2.4.3l1.1.8l1.4 1.1c.3.3.6.6 1 .9c.4.4.8.8 1.2 1.3c.7.8 1.3 1.6 1.8 2.5c.2.4.5.7.7 1.1c.3.5.6 1 .8 1.6c.2.4.4.8.5 1.2c.2.6.4 1.2.6 1.7c.1.4.3.8.4 1.2c.2.7.3 1.3.4 2c.1.3.1.7.2 1c.1 1 .2 2.1.2 3.1zm-254.1 13.3v111.3L723.1 597l-83.6-55.8zM281.8 472.6v79.8l59.8-39.9zM512 456.1l-84.5 56.4l84.5 56.4l84.5-56.4zM723.1 428L535.9 303.4v111.3l103.6 69.1zM384.5 541.2L300.9 597l187.2 124.6V610.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodepenOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m911.7 385.3l-.3-1.5c-.2-1-.3-1.9-.6-2.9c-.2-.6-.4-1.1-.5-1.7c-.3-.8-.5-1.7-.9-2.5c-.2-.6-.5-1.1-.8-1.7c-.4-.8-.8-1.5-1.2-2.3c-.3-.5-.6-1.1-1-1.6c-.8-1.2-1.7-2.4-2.6-3.6c-.5-.6-1.1-1.3-1.7-1.9c-.4-.5-.9-.9-1.4-1.3c-.6-.6-1.3-1.1-1.9-1.6c-.5-.4-1-.8-1.6-1.2c-.2-.1-.4-.3-.6-.4L531.1 117.8a34.3 34.3 0 0 0-38.1 0L127.3 361.3c-.2.1-.4.3-.6.4c-.5.4-1 .8-1.6 1.2c-.7.5-1.3 1.1-1.9 1.6c-.5.4-.9.9-1.4 1.3c-.6.6-1.2 1.2-1.7 1.9c-1 1.1-1.8 2.3-2.6 3.6c-.3.5-.7 1-1 1.6c-.4.7-.8 1.5-1.2 2.3c-.3.5-.5 1.1-.8 1.7c-.3.8-.6 1.7-.9 2.5c-.2.6-.4 1.1-.5 1.7c-.2.9-.4 1.9-.6 2.9l-.3 1.5c-.2 1.5-.3 3-.3 4.5v243.5c0 1.5.1 3 .3 4.5l.3 1.5l.6 2.9c.2.6.3 1.1.5 1.7c.3.9.6 1.7.9 2.5c.2.6.5 1.1.8 1.7c.4.8.7 1.5 1.2 2.3c.3.5.6 1.1 1 1.6c.5.7.9 1.4 1.5 2.1l1.2 1.5c.5.6 1.1 1.3 1.7 1.9c.4.5.9.9 1.4 1.3c.6.6 1.3 1.1 1.9 1.6c.5.4 1 .8 1.6 1.2c.2.1.4.3.6.4L493 905.7c5.6 3.8 12.3 5.8 19.1 5.8c6.6 0 13.3-1.9 19.1-5.8l365.6-243.5c.2-.1.4-.3.6-.4c.5-.4 1-.8 1.6-1.2c.7-.5 1.3-1.1 1.9-1.6c.5-.4.9-.9 1.4-1.3c.6-.6 1.2-1.2 1.7-1.9l1.2-1.5l1.5-2.1c.3-.5.7-1 1-1.6c.4-.8.8-1.5 1.2-2.3c.3-.5.5-1.1.8-1.7c.3-.8.6-1.7.9-2.5c.2-.5.4-1.1.5-1.7c.3-.9.4-1.9.6-2.9l.3-1.5c.2-1.5.3-3 .3-4.5V389.8c-.3-1.5-.4-3-.6-4.5M546.4 210.5l269.4 179.4l-120.3 80.4l-149-99.6V210.5zm-68.8 0v160.2l-149 99.6l-120.3-80.4zM180.7 454.1l86 57.5l-86 57.5zm296.9 358.5L208.3 633.2l120.3-80.4l149 99.6zM512 592.8l-121.6-81.2L512 430.3l121.6 81.2zm34.4 219.8V652.4l149-99.6l120.3 80.4zM843.3 569l-86-57.5l86-57.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodepenSquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M723.1 428L535.9 303.4v111.3l103.6 69.1zM512 456.1l-84.5 56.4l84.5 56.4l84.5-56.4zm23.9 154.2v111.3L723.1 597l-83.6-55.8zm-151.4-69.1L300.9 597l187.2 124.6V610.3zM880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-90 485c0 1.1-.1 2.1-.2 3.1c0 .4-.1.7-.2 1a14.16 14.16 0 0 1-.8 3.2c-.2.6-.4 1.2-.6 1.7c-.2.4-.4.8-.5 1.2c-.3.5-.5 1.1-.8 1.6c-.2.4-.4.7-.7 1.1c-.3.5-.7 1-1 1.5c-.3.4-.5.7-.8 1c-.4.4-.8.9-1.2 1.3c-.3.3-.6.6-1 .9c-.4.4-.9.8-1.4 1.1c-.4.3-.7.6-1.1.8c-.1.1-.3.2-.4.3L525.2 786c-4 2.7-8.6 4-13.2 4c-4.7 0-9.3-1.4-13.3-4L244.6 616.9c-.1-.1-.3-.2-.4-.3l-1.1-.8c-.5-.4-.9-.7-1.3-1.1c-.3-.3-.6-.6-1-.9c-.4-.4-.8-.8-1.2-1.3a7 7 0 0 1-.8-1c-.4-.5-.7-1-1-1.5c-.2-.4-.5-.7-.7-1.1c-.3-.5-.6-1.1-.8-1.6c-.2-.4-.4-.8-.5-1.2c-.2-.6-.4-1.2-.6-1.7c-.1-.4-.3-.8-.4-1.2c-.2-.7-.3-1.3-.4-2c-.1-.3-.1-.7-.2-1c-.1-1-.2-2.1-.2-3.1V427.9c0-1 .1-2.1.2-3.1c.1-.3.1-.7.2-1a14.16 14.16 0 0 1 .8-3.2c.2-.6.4-1.2.6-1.7c.2-.4.4-.8.5-1.2c.2-.5.5-1.1.8-1.6c.2-.4.4-.7.7-1.1c.6-.9 1.2-1.7 1.8-2.5c.4-.4.8-.9 1.2-1.3c.3-.3.6-.6 1-.9c.4-.4.9-.8 1.3-1.1c.4-.3.7-.6 1.1-.8c.1-.1.3-.2.4-.3L498.7 239c8-5.3 18.5-5.3 26.5 0l254.1 169.1c.1.1.3.2.4.3l1.1.8l1.4 1.1c.3.3.6.6 1 .9c.4.4.8.8 1.2 1.3c.7.8 1.3 1.6 1.8 2.5c.2.4.5.7.7 1.1c.3.5.6 1 .8 1.6c.2.4.4.8.5 1.2c.2.6.4 1.2.6 1.7c.1.4.3.8.4 1.2c.2.7.3 1.3.4 2c.1.3.1.7.2 1c.1 1 .2 2.1.2 3.1zm-47.8-44.6v-79.8l-59.8 39.9zm-460.4-79.8v79.8l59.8-39.9zm206.3-57.9V303.4L300.9 428l83.6 55.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoffeeOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M275 281c19.9 0 36-16.1 36-36V36c0-19.9-16.1-36-36-36s-36 16.1-36 36v209c0 19.9 16.1 36 36 36m613 144H768c0-39.8-32.2-72-72-72H200c-39.8 0-72 32.2-72 72v248c0 3.4.2 6.7.7 9.9c-.5 7-.7 14-.7 21.1c0 176.7 143.3 320 320 320c160.1 0 292.7-117.5 316.3-271H888c39.8 0 72-32.2 72-72V497c0-39.8-32.2-72-72-72M696 681h-1.1c.7 7.6 1.1 15.2 1.1 23c0 137-111 248-248 248S200 841 200 704c0-7.8.4-15.4 1.1-23H200V425h496zm192-8H776V497h112zM613 281c19.9 0 36-16.1 36-36V36c0-19.9-16.1-36-36-36s-36 16.1-36 36v209c0 19.9 16.1 36 36 36m-170 0c19.9 0 36-16.1 36-36V36c0-19.9-16.1-36-36-36s-36 16.1-36 36v209c0 19.9 16.1 36 36 36"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ColumnHeightOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M840 836H184c-4.4 0-8 3.6-8 8v60c0 4.4 3.6 8 8 8h656c4.4 0 8-3.6 8-8v-60c0-4.4-3.6-8-8-8m0-724H184c-4.4 0-8 3.6-8 8v60c0 4.4 3.6 8 8 8h656c4.4 0 8-3.6 8-8v-60c0-4.4-3.6-8-8-8M610.8 378c6 0 9.4-7 5.7-11.7L515.7 238.7a7.14 7.14 0 0 0-11.3 0L403.6 366.3a7.23 7.23 0 0 0 5.7 11.7H476v268h-62.8c-6 0-9.4 7-5.7 11.7l100.8 127.5c2.9 3.7 8.5 3.7 11.3 0l100.8-127.5c3.7-4.7.4-11.7-5.7-11.7H548V378z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ColumnWidthOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M180 176h-60c-4.4 0-8 3.6-8 8v656c0 4.4 3.6 8 8 8h60c4.4 0 8-3.6 8-8V184c0-4.4-3.6-8-8-8m724 0h-60c-4.4 0-8 3.6-8 8v656c0 4.4 3.6 8 8 8h60c4.4 0 8-3.6 8-8V184c0-4.4-3.6-8-8-8M785.3 504.3L657.7 403.6a7.23 7.23 0 0 0-11.7 5.7V476H378v-62.8c0-6-7-9.4-11.7-5.7L238.7 508.3a7.14 7.14 0 0 0 0 11.3l127.5 100.8c4.7 3.7 11.7.4 11.7-5.7V548h268v62.8c0 6 7 9.4 11.7 5.7l127.5-100.8c3.8-2.9 3.8-8.5.2-11.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M573 421c-23.1 0-41 17.9-41 40s17.9 40 41 40c21.1 0 39-17.9 39-40s-17.9-40-39-40m-280 0c-23.1 0-41 17.9-41 40s17.9 40 41 40c21.1 0 39-17.9 39-40s-17.9-40-39-40"/><path fill="currentColor" d="M894 345c-48.1-66-115.3-110.1-189-130v.1c-17.1-19-36.4-36.5-58-52.1c-163.7-119-393.5-82.7-513 81c-96.3 133-92.2 311.9 6 439l.8 132.6c0 3.2.5 6.4 1.5 9.4c5.3 16.9 23.3 26.2 40.1 20.9L309 806c33.5 11.9 68.1 18.7 102.5 20.6l-.5.4c89.1 64.9 205.9 84.4 313 49l127.1 41.4c3.2 1 6.5 1.6 9.9 1.6c17.7 0 32-14.3 32-32V753c88.1-119.6 90.4-284.9 1-408M323 735l-12-5l-99 31l-1-104l-8-9c-84.6-103.2-90.2-251.9-11-361c96.4-132.2 281.2-161.4 413-66c132.2 96.1 161.5 280.6 66 412c-80.1 109.9-223.5 150.5-348 102m505-17l-8 10l1 104l-98-33l-12 5c-56 20.8-115.7 22.5-171 7l-.2-.1C613.7 788.2 680.7 742.2 729 676c76.4-105.3 88.8-237.6 44.4-350.4l.6.4c23 16.5 44.1 37.1 62 62c72.6 99.6 68.5 235.2-8 330"/><path fill="currentColor" d="M433 421c-23.1 0-41 17.9-41 40s17.9 40 41 40c21.1 0 39-17.9 39-40s-17.9-40-39-40"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompassFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64M327.3 702.4c-2 .9-4.4 0-5.3-2.1c-.4-1-.4-2.2 0-3.2l98.7-225.5l132.1 132.1zm375.1-375.1l-98.7 225.5l-132.1-132.1L697.1 322c2-.9 4.4 0 5.3 2.1c.4 1 .4 2.1 0 3.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompassOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372m198.4-588.1a32 32 0 0 0-24.5.5L414.9 415L296.4 686c-3.6 8.2-3.6 17.5 0 25.7c3.4 7.8 9.7 13.9 17.7 17c3.8 1.5 7.7 2.2 11.7 2.2c4.4 0 8.7-.9 12.8-2.7l271-118.6l118.5-271a32.06 32.06 0 0 0-17.7-42.7M576.8 534.4l26.2 26.2l-42.4 42.4l-26.2-26.2L380 644.4L447.5 490L422 464.4l42.4-42.4l25.5 25.5L644.4 380zM464.4 422L422 464.4l25.5 25.6l86.9 86.8l26.2 26.2l42.4-42.4l-26.2-26.2l-86.8-86.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompassTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372M327.6 701.7c-2 .9-4.4 0-5.3-2.1c-.4-1-.4-2.2 0-3.2L421 470.9L553.1 603zm375.1-375.1L604 552.1L471.9 420l225.5-98.7c2-.9 4.4 0 5.3 2.1c.4 1 .4 2.1 0 3.2"/><path fill="currentColor" d="M322.3 696.4c-.4 1-.4 2.2 0 3.2c.9 2.1 3.3 3 5.3 2.1L553.1 603L421 470.9zm375.1-375.1L471.9 420L604 552.1l98.7-225.5c.4-1.1.4-2.2 0-3.2c-.9-2.1-3.3-3-5.3-2.1"/><path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompressOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M326 664H104c-8.8 0-16 7.2-16 16v48c0 8.8 7.2 16 16 16h174v176c0 8.8 7.2 16 16 16h48c8.8 0 16-7.2 16-16V696c0-17.7-14.3-32-32-32m16-576h-48c-8.8 0-16 7.2-16 16v176H104c-8.8 0-16 7.2-16 16v48c0 8.8 7.2 16 16 16h222c17.7 0 32-14.3 32-32V104c0-8.8-7.2-16-16-16m578 576H698c-17.7 0-32 14.3-32 32v224c0 8.8 7.2 16 16 16h48c8.8 0 16-7.2 16-16V744h174c8.8 0 16-7.2 16-16v-48c0-8.8-7.2-16-16-16m0-384H746V104c0-8.8-7.2-16-16-16h-48c-8.8 0-16 7.2-16 16v224c0 17.7 14.3 32 32 32h222c8.8 0 16-7.2 16-16v-48c0-8.8-7.2-16-16-16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ConsoleSqlOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M301.3 496.7c-23.8 0-40.2-10.5-41.6-26.9H205c.9 43.4 36.9 70.3 93.9 70.3c59.1 0 95-28.4 95-75.5c0-35.8-20-55.9-64.5-64.5l-29.1-5.6c-23.8-4.7-33.8-11.9-33.8-24.2c0-15 13.3-24.5 33.4-24.5c20.1 0 35.3 11.1 36.6 27h53c-.9-41.7-37.5-70.3-90.3-70.3c-54.4 0-89.7 28.9-89.7 73c0 35.5 21.2 58 62.5 65.8l29.7 5.9c25.8 5.2 35.6 11.9 35.6 24.4c.1 14.7-14.5 25.1-36 25.1"/><path fill="currentColor" d="M928 140H96c-17.7 0-32 14.3-32 32v496c0 17.7 14.3 32 32 32h380v112H304c-8.8 0-16 7.2-16 16v48c0 4.4 3.6 8 8 8h432c4.4 0 8-3.6 8-8v-48c0-8.8-7.2-16-16-16H548V700h380c17.7 0 32-14.3 32-32V172c0-17.7-14.3-32-32-32m-40 488H136V212h752z"/><path fill="currentColor" d="M828.5 486.7h-95.8V308.5h-57.4V534h153.2zm-298.6 53.4c14.1 0 27.2-2 39.1-5.8l13.3 20.3h53.3L607.9 511c21.1-20 33-51.1 33-89.8c0-73.3-43.3-118.8-110.9-118.8s-111.2 45.3-111.2 118.8c-.1 73.7 43 118.9 111.1 118.9m0-190c31.6 0 52.7 27.7 52.7 71.1c0 16.7-3.6 30.6-10 40.5l-5.2-6.9h-48.8L542 491c-3.9.9-8 1.4-12.2 1.4c-31.7 0-52.8-27.5-52.8-71.2c.1-43.6 21.2-71.1 52.9-71.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContactsFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M928 224H768v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56H548v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56H328v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56H96c-17.7 0-32 14.3-32 32v576c0 17.7 14.3 32 32 32h832c17.7 0 32-14.3 32-32V256c0-17.7-14.3-32-32-32M661 736h-43.9c-4.2 0-7.6-3.3-7.9-7.5c-3.8-50.6-46-90.5-97.2-90.5s-93.4 40-97.2 90.5c-.3 4.2-3.7 7.5-7.9 7.5H363a8 8 0 0 1-8-8.4c2.8-53.3 32-99.7 74.6-126.1a111.8 111.8 0 0 1-29.1-75.5c0-61.9 49.9-112 111.4-112c61.5 0 111.4 50.1 111.4 112c0 29.1-11 55.5-29.1 75.5c42.7 26.5 71.8 72.8 74.6 126.1c.4 4.6-3.2 8.4-7.8 8.4M512 474c-28.5 0-51.7 23.3-51.7 52s23.2 52 51.7 52c28.5 0 51.7-23.3 51.7-52s-23.2-52-51.7-52"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContactsOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M594.3 601.5a111.8 111.8 0 0 0 29.1-75.5c0-61.9-49.9-112-111.4-112s-111.4 50.1-111.4 112c0 29.1 11 55.5 29.1 75.5a158.09 158.09 0 0 0-74.6 126.1a8 8 0 0 0 8 8.4H407c4.2 0 7.6-3.3 7.9-7.5c3.8-50.6 46-90.5 97.2-90.5s93.4 40 97.2 90.5c.3 4.2 3.7 7.5 7.9 7.5H661a8 8 0 0 0 8-8.4c-2.8-53.3-32-99.7-74.7-126.1M512 578c-28.5 0-51.7-23.3-51.7-52s23.2-52 51.7-52s51.7 23.3 51.7 52s-23.2 52-51.7 52m416-354H768v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56H548v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56H328v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56H96c-17.7 0-32 14.3-32 32v576c0 17.7 14.3 32 32 32h832c17.7 0 32-14.3 32-32V256c0-17.7-14.3-32-32-32m-40 568H136V296h120v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56h148v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56h148v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56h120z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContactsTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M460.3 526a51.7 52 0 1 0 103.4 0a51.7 52 0 1 0-103.4 0"/><path fill="currentColor" fill-opacity=".15" d="M768 352c0 4.4-3.6 8-8 8h-56c-4.4 0-8-3.6-8-8v-56H548v56c0 4.4-3.6 8-8 8h-56c-4.4 0-8-3.6-8-8v-56H328v56c0 4.4-3.6 8-8 8h-56c-4.4 0-8-3.6-8-8v-56H136v496h752V296H768zM661 736h-43.8c-4.2 0-7.6-3.3-7.9-7.5c-3.8-50.5-46-90.5-97.2-90.5s-93.4 39.9-97.2 90.5c-.3 4.2-3.7 7.5-7.9 7.5h-43.9a8 8 0 0 1-8-8.4c2.8-53.3 31.9-99.6 74.6-126.1c-18.1-20-29.1-46.4-29.1-75.5c0-61.9 49.9-112 111.4-112s111.4 50.1 111.4 112c0 29.1-11 55.6-29.1 75.5c42.7 26.4 71.9 72.8 74.7 126.1a8 8 0 0 1-8 8.4"/><path fill="currentColor" d="M594.3 601.5a111.8 111.8 0 0 0 29.1-75.5c0-61.9-49.9-112-111.4-112s-111.4 50.1-111.4 112c0 29.1 11 55.5 29.1 75.5a158.09 158.09 0 0 0-74.6 126.1a8 8 0 0 0 8 8.4H407c4.2 0 7.6-3.3 7.9-7.5c3.8-50.6 46-90.5 97.2-90.5s93.4 40 97.2 90.5c.3 4.2 3.7 7.5 7.9 7.5H661a8 8 0 0 0 8-8.4c-2.8-53.3-32-99.7-74.7-126.1M512 578c-28.5 0-51.7-23.3-51.7-52s23.2-52 51.7-52s51.7 23.3 51.7 52s-23.2 52-51.7 52"/><path fill="currentColor" d="M928 224H768v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56H548v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56H328v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56H96c-17.7 0-32 14.3-32 32v576c0 17.7 14.3 32 32 32h832c17.7 0 32-14.3 32-32V256c0-17.7-14.3-32-32-32m-40 568H136V296h120v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56h148v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56h148v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56h120z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContainerFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 64H192c-17.7 0-32 14.3-32 32v529c0-.6.4-1 1-1h219.3l5.2 24.7C397.6 708.5 450.8 752 512 752s114.4-43.5 126.4-103.3l5.2-24.7H863c.6 0 1 .4 1 1V96c0-17.7-14.3-32-32-32M712 493c0 4.4-3.6 8-8 8H320c-4.4 0-8-3.6-8-8v-48c0-4.4 3.6-8 8-8h384c4.4 0 8 3.6 8 8zm0-160c0 4.4-3.6 8-8 8H320c-4.4 0-8-3.6-8-8v-48c0-4.4 3.6-8 8-8h384c4.4 0 8 3.6 8 8zm151 354H694.1c-11.6 32.8-32 62.3-59.1 84.7c-34.5 28.6-78.2 44.3-123 44.3s-88.5-15.8-123-44.3a194.02 194.02 0 0 1-59.1-84.7H161c-.6 0-1-.4-1-1v242c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V686c0 .6-.4 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContainerOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 64H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V96c0-17.7-14.3-32-32-32m-40 824H232V687h97.9c11.6 32.8 32 62.3 59.1 84.7c34.5 28.5 78.2 44.3 123 44.3s88.5-15.7 123-44.3c27.1-22.4 47.5-51.9 59.1-84.7H792v-63H643.6l-5.2 24.7C626.4 708.5 573.2 752 512 752s-114.4-43.5-126.5-103.3l-5.2-24.7H232V136h560zM320 341h384c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8H320c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8m0 160h384c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8H320c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContainerTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M635 771.7c-34.5 28.6-78.2 44.3-123 44.3s-88.5-15.8-123-44.3a194.02 194.02 0 0 1-59.1-84.7H232v201h560V687h-97.9c-11.6 32.8-32 62.3-59.1 84.7"/><path fill="currentColor" d="M320 501h384c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8H320c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8"/><path fill="currentColor" d="M832 64H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V96c0-17.7-14.3-32-32-32m-40 824H232V687h97.9c11.6 32.8 32 62.3 59.1 84.7c34.5 28.5 78.2 44.3 123 44.3s88.5-15.7 123-44.3c27.1-22.4 47.5-51.9 59.1-84.7H792zm0-264H643.6l-5.2 24.7C626.4 708.5 573.2 752 512 752s-114.4-43.5-126.5-103.3l-5.2-24.7H232V136h560z"/><path fill="currentColor" d="M320 341h384c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8H320c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ControlFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M404 683v77c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8v-77c-41.7-13.6-72-52.8-72-99s30.3-85.5 72-99V264c0-4.4 3.6-8 8-8h48c4.4 0 8 3.6 8 8v221c41.7 13.6 72 52.8 72 99s-30.3 85.5-72 99m279.6-143.9c.2 0 .3-.1.4-.1v221c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8V539c.2 0 .3.1.4.1c-42-13.4-72.4-52.7-72.4-99.1c0-46.4 30.4-85.7 72.4-99.1c-.2 0-.3.1-.4.1v-77c0-4.4 3.6-8 8-8h48c4.4 0 8 3.6 8 8v77c-.2 0-.3-.1-.4-.1c42 13.4 72.4 52.7 72.4 99.1c0 46.4-30.4 85.7-72.4 99.1M616 440a36 36 0 1 0 72 0a36 36 0 1 0-72 0M403.4 566.5l-1.5-2.4c0-.1-.1-.1-.1-.2l-.9-1.2c-.1-.1-.2-.2-.2-.3c-1-1.3-2-2.5-3.2-3.6l-.2-.2c-.4-.4-.8-.8-1.2-1.1c-.8-.8-1.7-1.5-2.6-2.1h-.1l-1.2-.9c-.1-.1-.3-.2-.4-.3c-1.2-.8-2.5-1.6-3.9-2.2c-.2-.1-.5-.2-.7-.4c-.4-.2-.7-.3-1.1-.5c-.3-.1-.7-.3-1-.4c-.5-.2-1-.4-1.5-.5c-.4-.1-.9-.3-1.3-.4l-.9-.3l-1.4-.3c-.2-.1-.5-.1-.7-.2c-.7-.1-1.4-.3-2.1-.4c-.2 0-.4 0-.6-.1c-.6-.1-1.1-.1-1.7-.2c-.2 0-.4 0-.7-.1c-.8 0-1.5-.1-2.3-.1s-1.5 0-2.3.1c-.2 0-.4 0-.7.1c-.6 0-1.2.1-1.7.2c-.2 0-.4 0-.6.1c-.7.1-1.4.2-2.1.4c-.2.1-.5.1-.7.2l-1.4.3l-.9.3c-.4.1-.9.3-1.3.4c-.5.2-1 .4-1.5.5c-.3.1-.7.3-1 .4c-.4.2-.7.3-1.1.5c-.2.1-.5.2-.7.4c-1.3.7-2.6 1.4-3.9 2.2c-.1.1-.3.2-.4.3l-1.2.9h-.1c-.9.7-1.8 1.4-2.6 2.1c-.4.4-.8.7-1.2 1.1l-.2.2a54.8 54.8 0 0 0-3.2 3.6c-.1.1-.2.2-.2.3l-.9 1.2c0 .1-.1.1-.1.2l-1.5 2.4c-.1.2-.2.3-.3.5c-2.7 5.1-4.3 10.9-4.3 17s1.6 12 4.3 17c.1.2.2.3.3.5l1.5 2.4c0 .1.1.1.1.2l.9 1.2c.1.1.2.2.2.3c1 1.3 2 2.5 3.2 3.6l.2.2c.4.4.8.8 1.2 1.1c.8.8 1.7 1.5 2.6 2.1h.1l1.2.9c.1.1.3.2.4.3c1.2.8 2.5 1.6 3.9 2.2c.2.1.5.2.7.4c.4.2.7.3 1.1.5c.3.1.7.3 1 .4c.5.2 1 .4 1.5.5c.4.1.9.3 1.3.4l.9.3l1.4.3c.2.1.5.1.7.2c.7.1 1.4.3 2.1.4c.2 0 .4 0 .6.1c.6.1 1.1.1 1.7.2c.2 0 .4 0 .7.1c.8 0 1.5.1 2.3.1s1.5 0 2.3-.1c.2 0 .4 0 .7-.1c.6 0 1.2-.1 1.7-.2c.2 0 .4 0 .6-.1c.7-.1 1.4-.2 2.1-.4c.2-.1.5-.1.7-.2l1.4-.3l.9-.3c.4-.1.9-.3 1.3-.4c.5-.2 1-.4 1.5-.5c.3-.1.7-.3 1-.4c.4-.2.7-.3 1.1-.5c.2-.1.5-.2.7-.4c1.3-.7 2.6-1.4 3.9-2.2c.1-.1.3-.2.4-.3l1.2-.9h.1c.9-.7 1.8-1.4 2.6-2.1c.4-.4.8-.7 1.2-1.1l.2-.2c1.1-1.1 2.2-2.4 3.2-3.6c.1-.1.2-.2.2-.3l.9-1.2c0-.1.1-.1.1-.2l1.5-2.4c.1-.2.2-.3.3-.5c2.7-5.1 4.3-10.9 4.3-17s-1.6-12-4.3-17c-.1-.2-.2-.4-.3-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ControlOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656zM340 683v77c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8v-77c-10.1 3.3-20.8 5-32 5s-21.9-1.8-32-5m64-198V264c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v221c10.1-3.3 20.8-5 32-5s21.9 1.8 32 5m-64 198c10.1 3.3 20.8 5 32 5s21.9-1.8 32-5c41.8-13.5 72-52.7 72-99s-30.2-85.5-72-99c-10.1-3.3-20.8-5-32-5s-21.9 1.8-32 5c-41.8 13.5-72 52.7-72 99s30.2 85.5 72 99m.1-115.7c.3-.6.7-1.2 1-1.8v-.1l1.2-1.8c.1-.2.2-.3.3-.5c.3-.5.7-.9 1-1.4c.1-.1.2-.3.3-.4c.5-.6.9-1.1 1.4-1.6l.3-.3l1.2-1.2l.4-.4c.5-.5 1-.9 1.6-1.4c.6-.5 1.1-.9 1.7-1.3c.2-.1.3-.2.5-.3c.5-.3.9-.7 1.4-1c.1-.1.3-.2.4-.3c.6-.4 1.2-.7 1.9-1.1c.1-.1.3-.1.4-.2c.5-.3 1-.5 1.6-.8l.6-.3c.7-.3 1.3-.6 2-.8c.7-.3 1.4-.5 2.1-.7c.2-.1.4-.1.6-.2c.6-.2 1.1-.3 1.7-.4c.2 0 .3-.1.5-.1c.7-.2 1.5-.3 2.2-.4c.2 0 .3 0 .5-.1c.6-.1 1.2-.1 1.8-.2h.6c.8 0 1.5-.1 2.3-.1s1.5 0 2.3.1h.6c.6 0 1.2.1 1.8.2c.2 0 .3 0 .5.1c.7.1 1.5.2 2.2.4c.2 0 .3.1.5.1c.6.1 1.2.3 1.7.4c.2.1.4.1.6.2c.7.2 1.4.4 2.1.7c.7.2 1.3.5 2 .8l.6.3c.5.2 1.1.5 1.6.8c.1.1.3.1.4.2c.6.3 1.3.7 1.9 1.1c.1.1.3.2.4.3c.5.3 1 .6 1.4 1c.2.1.3.2.5.3c.6.4 1.2.9 1.7 1.3s1.1.9 1.6 1.4l.4.4l1.2 1.2l.3.3c.5.5 1 1.1 1.4 1.6c.1.1.2.3.3.4c.4.4.7.9 1 1.4c.1.2.2.3.3.5l1.2 1.8s0 .1.1.1a36.18 36.18 0 0 1 5.1 18.5c0 6-1.5 11.7-4.1 16.7c-.3.6-.7 1.2-1 1.8c0 0 0 .1-.1.1l-1.2 1.8c-.1.2-.2.3-.3.5c-.3.5-.7.9-1 1.4c-.1.1-.2.3-.3.4c-.5.6-.9 1.1-1.4 1.6l-.3.3l-1.2 1.2l-.4.4c-.5.5-1 .9-1.6 1.4c-.6.5-1.1.9-1.7 1.3c-.2.1-.3.2-.5.3c-.5.3-.9.7-1.4 1c-.1.1-.3.2-.4.3c-.6.4-1.2.7-1.9 1.1c-.1.1-.3.1-.4.2c-.5.3-1 .5-1.6.8l-.6.3c-.7.3-1.3.6-2 .8c-.7.3-1.4.5-2.1.7c-.2.1-.4.1-.6.2c-.6.2-1.1.3-1.7.4c-.2 0-.3.1-.5.1c-.7.2-1.5.3-2.2.4c-.2 0-.3 0-.5.1c-.6.1-1.2.1-1.8.2h-.6c-.8 0-1.5.1-2.3.1s-1.5 0-2.3-.1h-.6c-.6 0-1.2-.1-1.8-.2c-.2 0-.3 0-.5-.1c-.7-.1-1.5-.2-2.2-.4c-.2 0-.3-.1-.5-.1c-.6-.1-1.2-.3-1.7-.4c-.2-.1-.4-.1-.6-.2c-.7-.2-1.4-.4-2.1-.7c-.7-.2-1.3-.5-2-.8l-.6-.3c-.5-.2-1.1-.5-1.6-.8c-.1-.1-.3-.1-.4-.2c-.6-.3-1.3-.7-1.9-1.1c-.1-.1-.3-.2-.4-.3c-.5-.3-1-.6-1.4-1c-.2-.1-.3-.2-.5-.3c-.6-.4-1.2-.9-1.7-1.3s-1.1-.9-1.6-1.4l-.4-.4l-1.2-1.2l-.3-.3c-.5-.5-1-1.1-1.4-1.6c-.1-.1-.2-.3-.3-.4c-.4-.4-.7-.9-1-1.4c-.1-.2-.2-.3-.3-.5l-1.2-1.8v-.1c-.4-.6-.7-1.2-1-1.8c-2.6-5-4.1-10.7-4.1-16.7s1.5-11.7 4.1-16.7M620 539v221c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V539c-10.1 3.3-20.8 5-32 5s-21.9-1.8-32-5m64-198v-77c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v77c10.1-3.3 20.8-5 32-5s21.9 1.8 32 5m-64 198c10.1 3.3 20.8 5 32 5s21.9-1.8 32-5c41.8-13.5 72-52.7 72-99s-30.2-85.5-72-99c-10.1-3.3-20.8-5-32-5s-21.9 1.8-32 5c-41.8 13.5-72 52.7-72 99s30.2 85.5 72 99m.1-115.7c.3-.6.7-1.2 1-1.8v-.1l1.2-1.8c.1-.2.2-.3.3-.5c.3-.5.7-.9 1-1.4c.1-.1.2-.3.3-.4c.5-.6.9-1.1 1.4-1.6l.3-.3l1.2-1.2l.4-.4c.5-.5 1-.9 1.6-1.4c.6-.5 1.1-.9 1.7-1.3c.2-.1.3-.2.5-.3c.5-.3.9-.7 1.4-1c.1-.1.3-.2.4-.3c.6-.4 1.2-.7 1.9-1.1c.1-.1.3-.1.4-.2c.5-.3 1-.5 1.6-.8l.6-.3c.7-.3 1.3-.6 2-.8c.7-.3 1.4-.5 2.1-.7c.2-.1.4-.1.6-.2c.6-.2 1.1-.3 1.7-.4c.2 0 .3-.1.5-.1c.7-.2 1.5-.3 2.2-.4c.2 0 .3 0 .5-.1c.6-.1 1.2-.1 1.8-.2h.6c.8 0 1.5-.1 2.3-.1s1.5 0 2.3.1h.6c.6 0 1.2.1 1.8.2c.2 0 .3 0 .5.1c.7.1 1.5.2 2.2.4c.2 0 .3.1.5.1c.6.1 1.2.3 1.7.4c.2.1.4.1.6.2c.7.2 1.4.4 2.1.7c.7.2 1.3.5 2 .8l.6.3c.5.2 1.1.5 1.6.8c.1.1.3.1.4.2c.6.3 1.3.7 1.9 1.1c.1.1.3.2.4.3c.5.3 1 .6 1.4 1c.2.1.3.2.5.3c.6.4 1.2.9 1.7 1.3s1.1.9 1.6 1.4l.4.4l1.2 1.2l.3.3c.5.5 1 1.1 1.4 1.6c.1.1.2.3.3.4c.4.4.7.9 1 1.4c.1.2.2.3.3.5l1.2 1.8v.1a36.18 36.18 0 0 1 5.1 18.5c0 6-1.5 11.7-4.1 16.7c-.3.6-.7 1.2-1 1.8v.1l-1.2 1.8c-.1.2-.2.3-.3.5c-.3.5-.7.9-1 1.4c-.1.1-.2.3-.3.4c-.5.6-.9 1.1-1.4 1.6l-.3.3l-1.2 1.2l-.4.4c-.5.5-1 .9-1.6 1.4c-.6.5-1.1.9-1.7 1.3c-.2.1-.3.2-.5.3c-.5.3-.9.7-1.4 1c-.1.1-.3.2-.4.3c-.6.4-1.2.7-1.9 1.1c-.1.1-.3.1-.4.2c-.5.3-1 .5-1.6.8l-.6.3c-.7.3-1.3.6-2 .8c-.7.3-1.4.5-2.1.7c-.2.1-.4.1-.6.2c-.6.2-1.1.3-1.7.4c-.2 0-.3.1-.5.1c-.7.2-1.5.3-2.2.4c-.2 0-.3 0-.5.1c-.6.1-1.2.1-1.8.2h-.6c-.8 0-1.5.1-2.3.1s-1.5 0-2.3-.1h-.6c-.6 0-1.2-.1-1.8-.2c-.2 0-.3 0-.5-.1c-.7-.1-1.5-.2-2.2-.4c-.2 0-.3-.1-.5-.1c-.6-.1-1.2-.3-1.7-.4c-.2-.1-.4-.1-.6-.2c-.7-.2-1.4-.4-2.1-.7c-.7-.2-1.3-.5-2-.8l-.6-.3c-.5-.2-1.1-.5-1.6-.8c-.1-.1-.3-.1-.4-.2c-.6-.3-1.3-.7-1.9-1.1c-.1-.1-.3-.2-.4-.3c-.5-.3-1-.6-1.4-1c-.2-.1-.3-.2-.5-.3c-.6-.4-1.2-.9-1.7-1.3s-1.1-.9-1.6-1.4l-.4-.4l-1.2-1.2l-.3-.3c-.5-.5-1-1.1-1.4-1.6c-.1-.1-.2-.3-.3-.4c-.4-.4-.7-.9-1-1.4c-.1-.2-.2-.3-.3-.5l-1.2-1.8v-.1c-.4-.6-.7-1.2-1-1.8c-2.6-5-4.1-10.7-4.1-16.7s1.5-11.7 4.1-16.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ControlTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/><path fill="currentColor" fill-opacity=".15" d="M616 440a36 36 0 1 0 72 0a36 36 0 1 0-72 0M340.4 601.5l1.5 2.4c0 .1.1.1.1.2l.9 1.2c.1.1.2.2.2.3c1 1.3 2 2.5 3.2 3.6l.2.2c.4.4.8.8 1.2 1.1c.8.8 1.7 1.5 2.6 2.1h.1l1.2.9c.1.1.3.2.4.3c1.2.8 2.5 1.6 3.9 2.2c.2.1.5.2.7.4c.4.2.7.3 1.1.5c.3.1.7.3 1 .4c.5.2 1 .4 1.5.5c.4.1.9.3 1.3.4l.9.3l1.4.3c.2.1.5.1.7.2c.7.1 1.4.3 2.1.4c.2 0 .4 0 .6.1c.6.1 1.1.1 1.7.2c.2 0 .4 0 .7.1c.8 0 1.5.1 2.3.1s1.5 0 2.3-.1c.2 0 .4 0 .7-.1c.6 0 1.2-.1 1.7-.2c.2 0 .4 0 .6-.1c.7-.1 1.4-.2 2.1-.4c.2-.1.5-.1.7-.2l1.4-.3l.9-.3c.4-.1.9-.3 1.3-.4c.5-.2 1-.4 1.5-.5c.3-.1.7-.3 1-.4c.4-.2.7-.3 1.1-.5c.2-.1.5-.2.7-.4c1.3-.7 2.6-1.4 3.9-2.2c.1-.1.3-.2.4-.3l1.2-.9h.1c.9-.7 1.8-1.4 2.6-2.1c.4-.4.8-.7 1.2-1.1l.2-.2c1.1-1.1 2.2-2.4 3.2-3.6c.1-.1.2-.2.2-.3l.9-1.2c0-.1.1-.1.1-.2l1.5-2.4c.1-.2.2-.3.3-.5c2.7-5.1 4.3-10.9 4.3-17s-1.6-12-4.3-17c-.1-.2-.2-.4-.3-.5l-1.5-2.4c0-.1-.1-.1-.1-.2l-.9-1.2c-.1-.1-.2-.2-.2-.3c-1-1.3-2-2.5-3.2-3.6l-.2-.2c-.4-.4-.8-.8-1.2-1.1c-.8-.8-1.7-1.5-2.6-2.1h-.1l-1.2-.9c-.1-.1-.3-.2-.4-.3c-1.2-.8-2.5-1.6-3.9-2.2c-.2-.1-.5-.2-.7-.4c-.4-.2-.7-.3-1.1-.5c-.3-.1-.7-.3-1-.4c-.5-.2-1-.4-1.5-.5c-.4-.1-.9-.3-1.3-.4l-.9-.3l-1.4-.3c-.2-.1-.5-.1-.7-.2c-.7-.1-1.4-.3-2.1-.4c-.2 0-.4 0-.6-.1c-.6-.1-1.1-.1-1.7-.2c-.2 0-.4 0-.7-.1c-.8 0-1.5-.1-2.3-.1s-1.5 0-2.3.1c-.2 0-.4 0-.7.1c-.6 0-1.2.1-1.7.2c-.2 0-.4 0-.6.1c-.7.1-1.4.2-2.1.4c-.2.1-.5.1-.7.2l-1.4.3l-.9.3c-.4.1-.9.3-1.3.4c-.5.2-1 .4-1.5.5c-.3.1-.7.3-1 .4c-.4.2-.7.3-1.1.5c-.2.1-.5.2-.7.4c-1.3.7-2.6 1.4-3.9 2.2c-.1.1-.3.2-.4.3l-1.2.9h-.1c-.9.7-1.8 1.4-2.6 2.1c-.4.4-.8.7-1.2 1.1l-.2.2a54.8 54.8 0 0 0-3.2 3.6c-.1.1-.2.2-.2.3l-.9 1.2c0 .1-.1.1-.1.2l-1.5 2.4c-.1.2-.2.3-.3.5c-2.7 5.1-4.3 10.9-4.3 17s1.6 12 4.3 17c.1.2.2.3.3.5"/><path fill="currentColor" fill-opacity=".15" d="M184 840h656V184H184zm436.4-499.1c-.2 0-.3.1-.4.1v-77c0-4.4 3.6-8 8-8h48c4.4 0 8 3.6 8 8v77c-.2 0-.3-.1-.4-.1c42 13.4 72.4 52.7 72.4 99.1c0 46.4-30.4 85.7-72.4 99.1c.2 0 .3-.1.4-.1v221c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8V539c.2 0 .3.1.4.1c-42-13.4-72.4-52.7-72.4-99.1c0-46.4 30.4-85.7 72.4-99.1M340 485V264c0-4.4 3.6-8 8-8h48c4.4 0 8 3.6 8 8v221c41.7 13.6 72 52.8 72 99s-30.3 85.5-72 99v77c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8v-77c-41.7-13.6-72-52.8-72-99s30.3-85.5 72-99"/><path fill="currentColor" d="M340 683v77c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8v-77c41.7-13.5 72-52.8 72-99s-30.3-85.4-72-99V264c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v221c-41.7 13.5-72 52.8-72 99s30.3 85.4 72 99m.1-116c.1-.2.2-.3.3-.5l1.5-2.4c0-.1.1-.1.1-.2l.9-1.2c0-.1.1-.2.2-.3c1-1.2 2.1-2.5 3.2-3.6l.2-.2c.4-.4.8-.7 1.2-1.1c.8-.7 1.7-1.4 2.6-2.1h.1l1.2-.9c.1-.1.3-.2.4-.3c1.3-.8 2.6-1.5 3.9-2.2c.2-.2.5-.3.7-.4c.4-.2.7-.3 1.1-.5c.3-.1.7-.3 1-.4c.5-.1 1-.3 1.5-.5c.4-.1.9-.3 1.3-.4l.9-.3l1.4-.3c.2-.1.5-.1.7-.2c.7-.2 1.4-.3 2.1-.4c.2-.1.4-.1.6-.1c.5-.1 1.1-.2 1.7-.2c.3-.1.5-.1.7-.1c.8-.1 1.5-.1 2.3-.1s1.5.1 2.3.1c.3.1.5.1.7.1c.6.1 1.1.1 1.7.2c.2.1.4.1.6.1c.7.1 1.4.3 2.1.4c.2.1.5.1.7.2l1.4.3l.9.3c.4.1.9.3 1.3.4c.5.1 1 .3 1.5.5c.3.1.7.3 1 .4c.4.2.7.3 1.1.5c.2.2.5.3.7.4c1.4.6 2.7 1.4 3.9 2.2c.1.1.3.2.4.3l1.2.9h.1c.9.6 1.8 1.3 2.6 2.1c.4.3.8.7 1.2 1.1l.2.2c1.2 1.1 2.2 2.3 3.2 3.6c0 .1.1.2.2.3l.9 1.2c0 .1.1.1.1.2l1.5 2.4A36.03 36.03 0 0 1 408 584c0 6.1-1.6 11.9-4.3 17c-.1.2-.2.3-.3.5l-1.5 2.4c0 .1-.1.1-.1.2l-.9 1.2c0 .1-.1.2-.2.3c-1 1.2-2.1 2.5-3.2 3.6l-.2.2c-.4.4-.8.7-1.2 1.1c-.8.7-1.7 1.4-2.6 2.1h-.1l-1.2.9c-.1.1-.3.2-.4.3c-1.3.8-2.6 1.5-3.9 2.2c-.2.2-.5.3-.7.4c-.4.2-.7.3-1.1.5c-.3.1-.7.3-1 .4c-.5.1-1 .3-1.5.5c-.4.1-.9.3-1.3.4l-.9.3l-1.4.3c-.2.1-.5.1-.7.2c-.7.2-1.4.3-2.1.4c-.2.1-.4.1-.6.1c-.5.1-1.1.2-1.7.2c-.3.1-.5.1-.7.1c-.8.1-1.5.1-2.3.1s-1.5-.1-2.3-.1c-.3-.1-.5-.1-.7-.1c-.6-.1-1.1-.1-1.7-.2c-.2-.1-.4-.1-.6-.1c-.7-.1-1.4-.3-2.1-.4c-.2-.1-.5-.1-.7-.2l-1.4-.3l-.9-.3c-.4-.1-.9-.3-1.3-.4c-.5-.1-1-.3-1.5-.5c-.3-.1-.7-.3-1-.4c-.4-.2-.7-.3-1.1-.5c-.2-.2-.5-.3-.7-.4c-1.4-.6-2.7-1.4-3.9-2.2c-.1-.1-.3-.2-.4-.3l-1.2-.9h-.1c-.9-.6-1.8-1.3-2.6-2.1c-.4-.3-.8-.7-1.2-1.1l-.2-.2c-1.2-1.1-2.2-2.3-3.2-3.6c0-.1-.1-.2-.2-.3l-.9-1.2c0-.1-.1-.1-.1-.2l-1.5-2.4c-.1-.2-.2-.3-.3-.5c-2.7-5-4.3-10.9-4.3-17s1.6-11.9 4.3-17m280.3-27.9c-.1 0-.2-.1-.4-.1v221c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V539c-.1 0-.2.1-.4.1c42-13.4 72.4-52.7 72.4-99.1c0-46.4-30.4-85.7-72.4-99.1c.1 0 .2.1.4.1v-77c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v77c.1 0 .2-.1.4-.1c-42 13.4-72.4 52.7-72.4 99.1c0 46.4 30.4 85.7 72.4 99.1M652 404c19.9 0 36 16.1 36 36s-16.1 36-36 36s-36-16.1-36-36s16.1-36 36-36"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 64H296c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h496v688c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V96c0-17.7-14.3-32-32-32M704 192H192c-17.7 0-32 14.3-32 32v530.7c0 8.5 3.4 16.6 9.4 22.6l173.3 173.3c2.2 2.2 4.7 4 7.4 5.5v1.9h4.2c3.5 1.3 7.2 2 11 2H704c17.7 0 32-14.3 32-32V224c0-17.7-14.3-32-32-32M382 896h-.2L232 746.2v-.2h150z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 64H296c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h496v688c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V96c0-17.7-14.3-32-32-32M704 192H192c-17.7 0-32 14.3-32 32v530.7c0 8.5 3.4 16.6 9.4 22.6l173.3 173.3c2.2 2.2 4.7 4 7.4 5.5v1.9h4.2c3.5 1.3 7.2 2 11 2H704c17.7 0 32-14.3 32-32V224c0-17.7-14.3-32-32-32M350 856.2L263.9 770H350zM664 888H414V746c0-22.1-17.9-40-40-40H232V264h432z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M232 706h142c22.1 0 40 17.9 40 40v142h250V264H232z"/><path fill="currentColor" d="M832 64H296c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h496v688c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V96c0-17.7-14.3-32-32-32"/><path fill="currentColor" d="M704 192H192c-17.7 0-32 14.3-32 32v530.7c0 8.5 3.4 16.6 9.4 22.6l173.3 173.3c2.2 2.2 4.7 4 7.4 5.5v1.9h4.2c3.5 1.3 7.2 2 11 2H704c17.7 0 32-14.3 32-32V224c0-17.7-14.3-32-32-32M350 856.2L263.9 770H350zM664 888H414V746c0-22.1-17.9-40-40-40H232V264h432z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyrightCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m5.4 670c-110 0-173.4-73.2-173.4-194.9v-52.3C344 364.2 407.4 290 517.3 290c94.3 0 162.7 60.7 162.7 147.4c0 2.6-2.1 4.7-4.7 4.7h-56.7c-4.2 0-7.6-3.2-8-7.4c-4-49.5-40-83.4-93-83.4c-65.3 0-102.1 48.5-102.1 135.5v52.6c0 85.7 36.9 133.6 102.1 133.6c52.8 0 88.7-31.7 93-77.8c.4-4.1 3.8-7.3 8-7.3h56.8c2.6 0 4.7 2.1 4.7 4.7c0 82.6-68.7 141.4-162.7 141.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyrightCircleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372m5.6-532.7c53 0 89 33.8 93 83.4c.3 4.2 3.8 7.4 8 7.4h56.7c2.6 0 4.7-2.1 4.7-4.7c0-86.7-68.4-147.4-162.7-147.4C407.4 290 344 364.2 344 486.8v52.3C344 660.8 407.4 734 517.3 734c94 0 162.7-58.8 162.7-141.4c0-2.6-2.1-4.7-4.7-4.7h-56.8c-4.2 0-7.6 3.2-8 7.3c-4.2 46.1-40.1 77.8-93 77.8c-65.3 0-102.1-47.9-102.1-133.6v-52.6c.1-87 37-135.5 102.2-135.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyrightCircleTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372m5.5 533c52.9 0 88.8-31.7 93-77.8c.4-4.1 3.8-7.3 8-7.3h56.8c2.6 0 4.7 2.1 4.7 4.7c0 82.6-68.7 141.4-162.7 141.4C407.4 734 344 660.8 344 539.1v-52.3C344 364.2 407.4 290 517.3 290c94.3 0 162.7 60.7 162.7 147.4c0 2.6-2.1 4.7-4.7 4.7h-56.7c-4.2 0-7.7-3.2-8-7.4c-4-49.6-40-83.4-93-83.4c-65.2 0-102.1 48.5-102.2 135.5v52.6c0 85.7 36.8 133.6 102.1 133.6"/><path fill="currentColor" d="M517.6 351.3c53 0 89 33.8 93 83.4c.3 4.2 3.8 7.4 8 7.4h56.7c2.6 0 4.7-2.1 4.7-4.7c0-86.7-68.4-147.4-162.7-147.4C407.4 290 344 364.2 344 486.8v52.3C344 660.8 407.4 734 517.3 734c94 0 162.7-58.8 162.7-141.4c0-2.6-2.1-4.7-4.7-4.7h-56.8c-4.2 0-7.6 3.2-8 7.3c-4.2 46.1-40.1 77.8-93 77.8c-65.3 0-102.1-47.9-102.1-133.6v-52.6c.1-87 37-135.5 102.2-135.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyrightOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372m5.6-532.7c53 0 89 33.8 93 83.4c.3 4.2 3.8 7.4 8 7.4h56.7c2.6 0 4.7-2.1 4.7-4.7c0-86.7-68.4-147.4-162.7-147.4C407.4 290 344 364.2 344 486.8v52.3C344 660.8 407.4 734 517.3 734c94 0 162.7-58.8 162.7-141.4c0-2.6-2.1-4.7-4.7-4.7h-56.8c-4.2 0-7.6 3.2-8 7.3c-4.2 46.1-40.1 77.8-93 77.8c-65.3 0-102.1-47.9-102.1-133.6v-52.6c.1-87 37-135.5 102.2-135.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyrightTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372m5.5 533c52.9 0 88.8-31.7 93-77.8c.4-4.1 3.8-7.3 8-7.3h56.8c2.6 0 4.7 2.1 4.7 4.7c0 82.6-68.7 141.4-162.7 141.4C407.4 734 344 660.8 344 539.1v-52.3C344 364.2 407.4 290 517.3 290c94.3 0 162.7 60.7 162.7 147.4c0 2.6-2.1 4.7-4.7 4.7h-56.7c-4.2 0-7.7-3.2-8-7.4c-4-49.6-40-83.4-93-83.4c-65.2 0-102.1 48.5-102.2 135.5v52.6c0 85.7 36.8 133.6 102.1 133.6"/><path fill="currentColor" d="M517.6 351.3c53 0 89 33.8 93 83.4c.3 4.2 3.8 7.4 8 7.4h56.7c2.6 0 4.7-2.1 4.7-4.7c0-86.7-68.4-147.4-162.7-147.4C407.4 290 344 364.2 344 486.8v52.3C344 660.8 407.4 734 517.3 734c94 0 162.7-58.8 162.7-141.4c0-2.6-2.1-4.7-4.7-4.7h-56.8c-4.2 0-7.6 3.2-8 7.3c-4.2 46.1-40.1 77.8-93 77.8c-65.3 0-102.1-47.9-102.1-133.6v-52.6c.1-87 37-135.5 102.2-135.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M928 160H96c-17.7 0-32 14.3-32 32v160h896V192c0-17.7-14.3-32-32-32M64 832c0 17.7 14.3 32 32 32h832c17.7 0 32-14.3 32-32V440H64zm579-184c0-4.4 3.6-8 8-8h165c4.4 0 8 3.6 8 8v72c0 4.4-3.6 8-8 8H651c-4.4 0-8-3.6-8-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M928 160H96c-17.7 0-32 14.3-32 32v640c0 17.7 14.3 32 32 32h832c17.7 0 32-14.3 32-32V192c0-17.7-14.3-32-32-32m-792 72h752v120H136zm752 560H136V440h752zm-237-64h165c4.4 0 8-3.6 8-8v-72c0-4.4-3.6-8-8-8H651c-4.4 0-8 3.6-8 8v72c0 4.4 3.6 8 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M136 792h752V440H136zm507-144c0-4.4 3.6-8 8-8h165c4.4 0 8 3.6 8 8v72c0 4.4-3.6 8-8 8H651c-4.4 0-8-3.6-8-8zM136 232h752v120H136z"/><path fill="currentColor" d="M651 728h165c4.4 0 8-3.6 8-8v-72c0-4.4-3.6-8-8-8H651c-4.4 0-8 3.6-8 8v72c0 4.4 3.6 8 8 8"/><path fill="currentColor" d="M928 160H96c-17.7 0-32 14.3-32 32v640c0 17.7 14.3 32 32 32h832c17.7 0 32-14.3 32-32V192c0-17.7-14.3-32-32-32m-40 632H136V440h752zm0-440H136V232h752z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CrownFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M899.6 276.5L705 396.4L518.4 147.5a8.06 8.06 0 0 0-12.9 0L319 396.4L124.3 276.5c-5.7-3.5-13.1 1.2-12.2 7.9L188.5 865c1.1 7.9 7.9 14 16 14h615.1c8 0 14.9-6 15.9-14l76.4-580.6c.8-6.7-6.5-11.4-12.3-7.9M512 734.2c-62.1 0-112.6-50.5-112.6-112.6S449.9 509 512 509s112.6 50.5 112.6 112.6S574.1 734.2 512 734.2m0-160.9c-26.6 0-48.2 21.6-48.2 48.3c0 26.6 21.6 48.3 48.2 48.3s48.2-21.6 48.2-48.3c0-26.6-21.6-48.3-48.2-48.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CrownOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M899.6 276.5L705 396.4L518.4 147.5a8.06 8.06 0 0 0-12.9 0L319 396.4L124.3 276.5c-5.7-3.5-13.1 1.2-12.2 7.9L188.5 865c1.1 7.9 7.9 14 16 14h615.1c8 0 14.9-6 15.9-14l76.4-580.6c.8-6.7-6.5-11.4-12.3-7.9m-126 534.1H250.3l-53.8-409.4l139.8 86.1L512 252.9l175.7 234.4l139.8-86.1zM512 509c-62.1 0-112.6 50.5-112.6 112.6S449.9 734.2 512 734.2s112.6-50.5 112.6-112.6S574.1 509 512 509m0 160.9c-26.6 0-48.2-21.6-48.2-48.3c0-26.6 21.6-48.3 48.2-48.3s48.2 21.6 48.2 48.3c0 26.6-21.6 48.3-48.2 48.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CrownTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M911.9 283.9v.5L835.5 865c-1 8-7.9 14-15.9 14H204.5c-8.1 0-14.9-6.1-16-14l-76.4-580.6v-.6v1.6L188.5 866c1.1 7.9 7.9 14 16 14h615.1c8 0 14.9-6 15.9-14l76.4-580.6c.1-.5.1-1 0-1.5"/><path fill="currentColor" fill-opacity=".15" d="m773.6 810.6l53.9-409.4l-139.8 86.1L512 252.9L336.3 487.3l-139.8-86.1l53.8 409.4zm-374.2-189c0-62.1 50.5-112.6 112.6-112.6s112.6 50.5 112.6 112.6v1c0 62.1-50.5 112.6-112.6 112.6s-112.6-50.5-112.6-112.6z"/><path fill="currentColor" d="M512 734.2c61.9 0 112.3-50.2 112.6-112.1v-.5c0-62.1-50.5-112.6-112.6-112.6s-112.6 50.5-112.6 112.6v.5c.3 61.9 50.7 112.1 112.6 112.1m0-160.9c26.6 0 48.2 21.6 48.2 48.3c0 26.6-21.6 48.3-48.2 48.3s-48.2-21.6-48.2-48.3c0-26.6 21.6-48.3 48.2-48.3"/><path fill="currentColor" d="M188.5 865c1.1 7.9 7.9 14 16 14h615.1c8 0 14.9-6 15.9-14l76.4-580.6v-.5c.3-6.4-6.7-10.8-12.3-7.4L705 396.4L518.4 147.5a8.06 8.06 0 0 0-12.9 0L319 396.4L124.3 276.5c-5.5-3.4-12.6.9-12.2 7.3v.6zm147.8-377.7L512 252.9l175.7 234.4l139.8-86.1l-53.9 409.4H250.3l-53.8-409.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CustomerServiceFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 128c-212.1 0-384 171.9-384 384v360c0 13.3 10.7 24 24 24h184c35.3 0 64-28.7 64-64V624c0-35.3-28.7-64-64-64H200v-48c0-172.3 139.7-312 312-312s312 139.7 312 312v48H688c-35.3 0-64 28.7-64 64v208c0 35.3 28.7 64 64 64h184c13.3 0 24-10.7 24-24V512c0-212.1-171.9-384-384-384"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CustomerServiceOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 128c-212.1 0-384 171.9-384 384v360c0 13.3 10.7 24 24 24h184c35.3 0 64-28.7 64-64V624c0-35.3-28.7-64-64-64H200v-48c0-172.3 139.7-312 312-312s312 139.7 312 312v48H688c-35.3 0-64 28.7-64 64v208c0 35.3 28.7 64 64 64h184c13.3 0 24-10.7 24-24V512c0-212.1-171.9-384-384-384M328 632v192H200V632zm496 192H696V632h128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CustomerServiceTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M696 632h128v192H696zm-496 0h128v192H200z"/><path fill="currentColor" d="M512 128c-212.1 0-384 171.9-384 384v360c0 13.3 10.7 24 24 24h184c35.3 0 64-28.7 64-64V624c0-35.3-28.7-64-64-64H200v-48c0-172.3 139.7-312 312-312s312 139.7 312 312v48H688c-35.3 0-64 28.7-64 64v208c0 35.3 28.7 64 64 64h184c13.3 0 24-10.7 24-24V512c0-212.1-171.9-384-384-384M328 632v192H200V632zm496 192H696V632h128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M112 476h160v72H112zm320 0h160v72H432zm320 0h160v72H752z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M924.8 385.6a446.7 446.7 0 0 0-96-142.4a446.7 446.7 0 0 0-142.4-96C631.1 123.8 572.5 112 512 112s-119.1 11.8-174.4 35.2a446.7 446.7 0 0 0-142.4 96a446.7 446.7 0 0 0-96 142.4C75.8 440.9 64 499.5 64 560c0 132.7 58.3 257.7 159.9 343.1l1.7 1.4c5.8 4.8 13.1 7.5 20.6 7.5h531.7c7.5 0 14.8-2.7 20.6-7.5l1.7-1.4C901.7 817.7 960 692.7 960 560c0-60.5-11.9-119.1-35.2-174.4M482 232c0-4.4 3.6-8 8-8h44c4.4 0 8 3.6 8 8v80c0 4.4-3.6 8-8 8h-44c-4.4 0-8-3.6-8-8zM270 582c0 4.4-3.6 8-8 8h-80c-4.4 0-8-3.6-8-8v-44c0-4.4 3.6-8 8-8h80c4.4 0 8 3.6 8 8zm90.7-204.5l-31.1 31.1a8.03 8.03 0 0 1-11.3 0L261.7 352a8.03 8.03 0 0 1 0-11.3l31.1-31.1c3.1-3.1 8.2-3.1 11.3 0l56.6 56.6c3.1 3.1 3.1 8.2 0 11.3m291.1 83.6l-84.5 84.5c5 18.7.2 39.4-14.5 54.1a55.95 55.95 0 0 1-79.2 0a55.95 55.95 0 0 1 0-79.2a55.87 55.87 0 0 1 54.1-14.5l84.5-84.5c3.1-3.1 8.2-3.1 11.3 0l28.3 28.3c3.1 3.1 3.1 8.1 0 11.3m43-52.4l-31.1-31.1a8.03 8.03 0 0 1 0-11.3l56.6-56.6c3.1-3.1 8.2-3.1 11.3 0l31.1 31.1c3.1 3.1 3.1 8.2 0 11.3l-56.6 56.6a8.03 8.03 0 0 1-11.3 0M846 582c0 4.4-3.6 8-8 8h-80c-4.4 0-8-3.6-8-8v-44c0-4.4 3.6-8 8-8h80c4.4 0 8 3.6 8 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M924.8 385.6a446.7 446.7 0 0 0-96-142.4a446.7 446.7 0 0 0-142.4-96C631.1 123.8 572.5 112 512 112s-119.1 11.8-174.4 35.2a446.7 446.7 0 0 0-142.4 96a446.7 446.7 0 0 0-96 142.4C75.8 440.9 64 499.5 64 560c0 132.7 58.3 257.7 159.9 343.1l1.7 1.4c5.8 4.8 13.1 7.5 20.6 7.5h531.7c7.5 0 14.8-2.7 20.6-7.5l1.7-1.4C901.7 817.7 960 692.7 960 560c0-60.5-11.9-119.1-35.2-174.4M761.4 836H262.6A371.12 371.12 0 0 1 140 560c0-99.4 38.7-192.8 109-263c70.3-70.3 163.7-109 263-109c99.4 0 192.8 38.7 263 109c70.3 70.3 109 163.7 109 263c0 105.6-44.5 205.5-122.6 276M623.5 421.5a8.03 8.03 0 0 0-11.3 0L527.7 506c-18.7-5-39.4-.2-54.1 14.5a55.95 55.95 0 0 0 0 79.2a55.95 55.95 0 0 0 79.2 0a55.87 55.87 0 0 0 14.5-54.1l84.5-84.5c3.1-3.1 3.1-8.2 0-11.3zM490 320h44c4.4 0 8-3.6 8-8v-80c0-4.4-3.6-8-8-8h-44c-4.4 0-8 3.6-8 8v80c0 4.4 3.6 8 8 8m260 218v44c0 4.4 3.6 8 8 8h80c4.4 0 8-3.6 8-8v-44c0-4.4-3.6-8-8-8h-80c-4.4 0-8 3.6-8 8m12.7-197.2l-31.1-31.1a8.03 8.03 0 0 0-11.3 0l-56.6 56.6a8.03 8.03 0 0 0 0 11.3l31.1 31.1c3.1 3.1 8.2 3.1 11.3 0l56.6-56.6c3.1-3.1 3.1-8.2 0-11.3m-458.6-31.1a8.03 8.03 0 0 0-11.3 0l-31.1 31.1a8.03 8.03 0 0 0 0 11.3l56.6 56.6c3.1 3.1 8.2 3.1 11.3 0l31.1-31.1c3.1-3.1 3.1-8.2 0-11.3zM262 530h-80c-4.4 0-8 3.6-8 8v44c0 4.4 3.6 8 8 8h80c4.4 0 8-3.6 8-8v-44c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M512 188c-99.3 0-192.7 38.7-263 109c-70.3 70.2-109 163.6-109 263c0 105.6 44.5 205.5 122.6 276h498.8A371.12 371.12 0 0 0 884 560c0-99.3-38.7-192.7-109-263c-70.2-70.3-163.6-109-263-109m-30 44c0-4.4 3.6-8 8-8h44c4.4 0 8 3.6 8 8v80c0 4.4-3.6 8-8 8h-44c-4.4 0-8-3.6-8-8zM270 582c0 4.4-3.6 8-8 8h-80c-4.4 0-8-3.6-8-8v-44c0-4.4 3.6-8 8-8h80c4.4 0 8 3.6 8 8zm90.7-204.4l-31.1 31.1a8.03 8.03 0 0 1-11.3 0l-56.6-56.6a8.03 8.03 0 0 1 0-11.3l31.1-31.1c3.1-3.1 8.2-3.1 11.3 0l56.6 56.6c3.1 3.1 3.1 8.2 0 11.3m291.1 83.5l-84.5 84.5c5 18.7.2 39.4-14.5 54.1a55.95 55.95 0 0 1-79.2 0a55.95 55.95 0 0 1 0-79.2a55.87 55.87 0 0 1 54.1-14.5l84.5-84.5c3.1-3.1 8.2-3.1 11.3 0l28.3 28.3c3.1 3.1 3.1 8.2 0 11.3m43-52.4l-31.1-31.1a8.03 8.03 0 0 1 0-11.3l56.6-56.6c3.1-3.1 8.2-3.1 11.3 0l31.1 31.1c3.1 3.1 3.1 8.2 0 11.3l-56.6 56.6a8.03 8.03 0 0 1-11.3 0M846 538v44c0 4.4-3.6 8-8 8h-80c-4.4 0-8-3.6-8-8v-44c0-4.4 3.6-8 8-8h80c4.4 0 8 3.6 8 8"/><path fill="currentColor" d="M623.5 421.5a8.03 8.03 0 0 0-11.3 0L527.7 506c-18.7-5-39.4-.2-54.1 14.5a55.95 55.95 0 0 0 0 79.2a55.95 55.95 0 0 0 79.2 0a55.87 55.87 0 0 0 14.5-54.1l84.5-84.5c3.1-3.1 3.1-8.2 0-11.3zM490 320h44c4.4 0 8-3.6 8-8v-80c0-4.4-3.6-8-8-8h-44c-4.4 0-8 3.6-8 8v80c0 4.4 3.6 8 8 8"/><path fill="currentColor" d="M924.8 385.6a446.7 446.7 0 0 0-96-142.4a446.7 446.7 0 0 0-142.4-96C631.1 123.8 572.5 112 512 112s-119.1 11.8-174.4 35.2a446.7 446.7 0 0 0-142.4 96a446.7 446.7 0 0 0-96 142.4C75.8 440.9 64 499.5 64 560c0 132.7 58.3 257.7 159.9 343.1l1.7 1.4c5.8 4.8 13.1 7.5 20.6 7.5h531.7c7.5 0 14.8-2.7 20.6-7.5l1.7-1.4C901.7 817.7 960 692.7 960 560c0-60.5-11.9-119.1-35.2-174.4M761.4 836H262.6A371.12 371.12 0 0 1 140 560c0-99.4 38.7-192.8 109-263c70.3-70.3 163.7-109 263-109c99.4 0 192.8 38.7 263 109c70.3 70.3 109 163.7 109 263c0 105.6-44.5 205.5-122.6 276"/><path fill="currentColor" d="m762.7 340.8l-31.1-31.1a8.03 8.03 0 0 0-11.3 0l-56.6 56.6a8.03 8.03 0 0 0 0 11.3l31.1 31.1c3.1 3.1 8.2 3.1 11.3 0l56.6-56.6c3.1-3.1 3.1-8.2 0-11.3M750 538v44c0 4.4 3.6 8 8 8h80c4.4 0 8-3.6 8-8v-44c0-4.4-3.6-8-8-8h-80c-4.4 0-8 3.6-8 8M304.1 309.7a8.03 8.03 0 0 0-11.3 0l-31.1 31.1a8.03 8.03 0 0 0 0 11.3l56.6 56.6c3.1 3.1 8.2 3.1 11.3 0l31.1-31.1c3.1-3.1 3.1-8.2 0-11.3zM262 530h-80c-4.4 0-8 3.6-8 8v44c0 4.4 3.6 8 8 8h80c4.4 0 8-3.6 8-8v-44c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 64H192c-17.7 0-32 14.3-32 32v224h704V96c0-17.7-14.3-32-32-32M288 232c-22.1 0-40-17.9-40-40s17.9-40 40-40s40 17.9 40 40s-17.9 40-40 40M160 928c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V704H160zm128-136c22.1 0 40 17.9 40 40s-17.9 40-40 40s-40-17.9-40-40s17.9-40 40-40M160 640h704V384H160zm128-168c22.1 0 40 17.9 40 40s-17.9 40-40 40s-40-17.9-40-40s17.9-40 40-40"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 64H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V96c0-17.7-14.3-32-32-32m-600 72h560v208H232zm560 480H232V408h560zm0 272H232V680h560zM304 240a40 40 0 1 0 80 0a40 40 0 1 0-80 0m0 272a40 40 0 1 0 80 0a40 40 0 1 0-80 0m0 272a40 40 0 1 0 80 0a40 40 0 1 0-80 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M232 616h560V408H232zm112-144c22.1 0 40 17.9 40 40s-17.9 40-40 40s-40-17.9-40-40s17.9-40 40-40M232 888h560V680H232zm112-144c22.1 0 40 17.9 40 40s-17.9 40-40 40s-40-17.9-40-40s17.9-40 40-40M232 344h560V136H232zm112-144c22.1 0 40 17.9 40 40s-17.9 40-40 40s-40-17.9-40-40s17.9-40 40-40"/><path fill="currentColor" d="M304 512a40 40 0 1 0 80 0a40 40 0 1 0-80 0m0 272a40 40 0 1 0 80 0a40 40 0 1 0-80 0m0-544a40 40 0 1 0 80 0a40 40 0 1 0-80 0"/><path fill="currentColor" d="M832 64H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V96c0-17.7-14.3-32-32-32m-40 824H232V680h560zm0-272H232V408h560zm0-272H232V136h560z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeleteColumnOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M651.1 641.9c-1.4-1.2-3.2-1.9-5.1-1.9h-54.7c-2.4 0-4.6 1.1-6.1 2.9L512 730.7l-73.1-87.8c-1.5-1.8-3.8-2.9-6.1-2.9H378c-1.9 0-3.7.7-5.1 1.9c-3.4 2.8-3.9 7.9-1 11.3L474.2 776L371.8 898.9c-2.8 3.4-2.4 8.4 1 11.3c1.4 1.2 3.2 1.9 5.1 1.9h54.7c2.4 0 4.6-1.1 6.1-2.9l73.1-87.8l73.1 87.8c1.5 1.8 3.8 2.9 6.1 2.9h55c1.9 0 3.7-.7 5.1-1.9c3.4-2.8 3.9-7.9 1-11.3L549.8 776l102.4-122.9c2.8-3.4 2.3-8.4-1.1-11.2M472 544h80c4.4 0 8-3.6 8-8V120c0-4.4-3.6-8-8-8h-80c-4.4 0-8 3.6-8 8v416c0 4.4 3.6 8 8 8M350 386H184V136c0-3.3-2.7-6-6-6h-60c-3.3 0-6 2.7-6 6v292c0 16.6 13.4 30 30 30h208c3.3 0 6-2.7 6-6v-60c0-3.3-2.7-6-6-6m556-256h-60c-3.3 0-6 2.7-6 6v250H674c-3.3 0-6 2.7-6 6v60c0 3.3 2.7 6 6 6h208c16.6 0 30-13.4 30-30V136c0-3.3-2.7-6-6-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeleteFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M864 256H736v-80c0-35.3-28.7-64-64-64H352c-35.3 0-64 28.7-64 64v80H160c-17.7 0-32 14.3-32 32v32c0 4.4 3.6 8 8 8h60.4l24.7 523c1.6 34.1 29.8 61 63.9 61h454c34.2 0 62.3-26.8 63.9-61l24.7-523H888c4.4 0 8-3.6 8-8v-32c0-17.7-14.3-32-32-32m-200 0H360v-72h304z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeleteOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M360 184h-8c4.4 0 8-3.6 8-8zh304v-8c0 4.4 3.6 8 8 8h-8v72h72v-80c0-35.3-28.7-64-64-64H352c-35.3 0-64 28.7-64 64v80h72zm504 72H160c-17.7 0-32 14.3-32 32v32c0 4.4 3.6 8 8 8h60.4l24.7 523c1.6 34.1 29.8 61 63.9 61h454c34.2 0 62.3-26.8 63.9-61l24.7-523H888c4.4 0 8-3.6 8-8v-32c0-17.7-14.3-32-32-32M731.3 840H292.7l-24.2-512h487z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeleteRowOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m819.8 512l102.4-122.9c2.8-3.4 2.4-8.4-1-11.3c-1.4-1.2-3.2-1.9-5.1-1.9h-54.7c-2.4 0-4.6 1.1-6.1 2.9L782 466.7l-73.1-87.8c-1.5-1.8-3.8-2.9-6.1-2.9H648c-1.9 0-3.7.7-5.1 1.9c-3.4 2.8-3.9 7.9-1 11.3L744.2 512L641.8 634.9c-2.8 3.4-2.4 8.4 1 11.3c1.4 1.2 3.2 1.9 5.1 1.9h54.7c2.4 0 4.6-1.1 6.1-2.9l73.1-87.8l73.1 87.8c1.5 1.8 3.8 2.9 6.1 2.9h55c1.9 0 3.7-.7 5.1-1.9c3.4-2.8 3.9-7.9 1-11.3zM536 464H120c-4.4 0-8 3.6-8 8v80c0 4.4 3.6 8 8 8h416c4.4 0 8-3.6 8-8v-80c0-4.4-3.6-8-8-8m-84 204h-60c-3.3 0-6 2.7-6 6v166H136c-3.3 0-6 2.7-6 6v60c0 3.3 2.7 6 6 6h292c16.6 0 30-13.4 30-30V674c0-3.3-2.7-6-6-6M136 184h250v166c0 3.3 2.7 6 6 6h60c3.3 0 6-2.7 6-6V142c0-16.6-13.4-30-30-30H136c-3.3 0-6 2.7-6 6v60c0 3.3 2.7 6 6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeleteTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M292.7 840h438.6l24.2-512h-487z"/><path fill="currentColor" d="M864 256H736v-80c0-35.3-28.7-64-64-64H352c-35.3 0-64 28.7-64 64v80H160c-17.7 0-32 14.3-32 32v32c0 4.4 3.6 8 8 8h60.4l24.7 523c1.6 34.1 29.8 61 63.9 61h454c34.2 0 62.3-26.8 63.9-61l24.7-523H888c4.4 0 8-3.6 8-8v-32c0-17.7-14.3-32-32-32m-504-72h304v72H360zm371.3 656H292.7l-24.2-512h487z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeliveredProcedureOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m632 698.3l141.9-112c4.1-3.2 4.1-9.4 0-12.6L632 461.7c-5.3-4.2-13-.4-13 6.3v76H295c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h324v76c0 6.7 7.8 10.4 13 6.3m261.3-405L730.7 130.7c-7.5-7.5-16.7-13-26.7-16V112H144c-17.7 0-32 14.3-32 32v278c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V184h136v136c0 17.7 14.3 32 32 32h320c17.7 0 32-14.3 32-32V205.8l136 136V422c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-83.5c0-17-6.7-33.2-18.7-45.2M640 288H384V184h256zm264 436h-56c-4.4 0-8 3.6-8 8v108H184V732c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v148c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V732c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeploymentUnitOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M888.3 693.2c-42.5-24.6-94.3-18-129.2 12.8l-53-30.7V523.6c0-15.7-8.4-30.3-22-38.1l-136-78.3v-67.1c44.2-15 76-56.8 76-106.1c0-61.9-50.1-112-112-112s-112 50.1-112 112c0 49.3 31.8 91.1 76 106.1v67.1l-136 78.3c-13.6 7.8-22 22.4-22 38.1v151.6l-53 30.7c-34.9-30.8-86.8-37.4-129.2-12.8c-53.5 31-71.7 99.4-41 152.9c30.8 53.5 98.9 71.9 152.2 41c42.5-24.6 62.7-73 53.6-118.8l48.7-28.3l140.6 81c6.8 3.9 14.4 5.9 22 5.9s15.2-2 22-5.9L674.5 740l48.7 28.3c-9.1 45.7 11.2 94.2 53.6 118.8c53.3 30.9 121.5 12.6 152.2-41c30.8-53.6 12.6-122-40.7-152.9m-673 138.4a47.6 47.6 0 0 1-65.2-17.6c-13.2-22.9-5.4-52.3 17.5-65.5a47.6 47.6 0 0 1 65.2 17.6c13.2 22.9 5.4 52.3-17.5 65.5M464 234a48.01 48.01 0 0 1 96 0a48.01 48.01 0 0 1-96 0m170 446.2l-122 70.3l-122-70.3V539.8l122-70.3l122 70.3zm239.9 133.9c-13.2 22.9-42.4 30.8-65.2 17.6c-22.8-13.2-30.7-42.6-17.5-65.5s42.4-30.8 65.2-17.6c22.9 13.2 30.7 42.5 17.5 65.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DesktopOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M928 140H96c-17.7 0-32 14.3-32 32v496c0 17.7 14.3 32 32 32h380v112H304c-8.8 0-16 7.2-16 16v48c0 4.4 3.6 8 8 8h432c4.4 0 8-3.6 8-8v-48c0-8.8-7.2-16-16-16H548V700h380c17.7 0 32-14.3 32-32V172c0-17.7-14.3-32-32-32m-40 488H136V212h752z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiffFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M854.2 306.6L611.3 72.9c-6-5.7-13.9-8.9-22.2-8.9H296c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h277l219 210.6V824c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V329.6c0-8.7-3.5-17-9.8-23M553.4 201.4c-6-6-14.1-9.4-22.6-9.4H192c-17.7 0-32 14.3-32 32v704c0 17.7 14.3 32 32 32h512c17.7 0 32-14.3 32-32V397.3c0-8.5-3.4-16.6-9.4-22.6zM568 753c0 3.8-3.4 7-7.5 7h-225c-4.1 0-7.5-3.2-7.5-7v-42c0-3.8 3.4-7 7.5-7h225c4.1 0 7.5 3.2 7.5 7zm0-220c0 3.8-3.4 7-7.5 7H476v84.9c0 3.9-3.1 7.1-7 7.1h-42c-3.8 0-7-3.2-7-7.1V540h-84.5c-4.1 0-7.5-3.2-7.5-7v-42c0-3.9 3.4-7 7.5-7H420v-84.9c0-3.9 3.2-7.1 7-7.1h42c3.9 0 7 3.2 7 7.1V484h84.5c4.1 0 7.5 3.1 7.5 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiffOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M476 399.1c0-3.9-3.1-7.1-7-7.1h-42c-3.8 0-7 3.2-7 7.1V484h-84.5c-4.1 0-7.5 3.1-7.5 7v42c0 3.8 3.4 7 7.5 7H420v84.9c0 3.9 3.2 7.1 7 7.1h42c3.9 0 7-3.2 7-7.1V540h84.5c4.1 0 7.5-3.2 7.5-7v-42c0-3.9-3.4-7-7.5-7H476zM560.5 704h-225c-4.1 0-7.5 3.2-7.5 7v42c0 3.8 3.4 7 7.5 7h225c4.1 0 7.5-3.2 7.5-7v-42c0-3.8-3.4-7-7.5-7m-7.1-502.6c-6-6-14.1-9.4-22.6-9.4H192c-17.7 0-32 14.3-32 32v704c0 17.7 14.3 32 32 32h512c17.7 0 32-14.3 32-32V397.3c0-8.5-3.4-16.6-9.4-22.6zM664 888H232V264h282.2L664 413.8zm190.2-581.4L611.3 72.9c-6-5.7-13.9-8.9-22.2-8.9H296c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h277l219 210.6V824c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V329.6c0-8.7-3.5-17-9.8-23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiffTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M232 264v624h432V413.8L514.2 264zm336 489c0 3.8-3.4 7-7.5 7h-225c-4.1 0-7.5-3.2-7.5-7v-42c0-3.8 3.4-7 7.5-7h225c4.1 0 7.5 3.2 7.5 7zm0-262v42c0 3.8-3.4 7-7.5 7H476v84.9c0 3.9-3.1 7.1-7 7.1h-42c-3.8 0-7-3.2-7-7.1V540h-84.5c-4.1 0-7.5-3.2-7.5-7v-42c0-3.9 3.4-7 7.5-7H420v-84.9c0-3.9 3.2-7.1 7-7.1h42c3.9 0 7 3.2 7 7.1V484h84.5c4.1 0 7.5 3.1 7.5 7"/><path fill="currentColor" d="M854.2 306.6L611.3 72.9c-6-5.7-13.9-8.9-22.2-8.9H296c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h277l219 210.6V824c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V329.6c0-8.7-3.5-17-9.8-23"/><path fill="currentColor" d="M553.4 201.4c-6-6-14.1-9.4-22.6-9.4H192c-17.7 0-32 14.3-32 32v704c0 17.7 14.3 32 32 32h512c17.7 0 32-14.3 32-32V397.3c0-8.5-3.4-16.6-9.4-22.6zM664 888H232V264h282.2L664 413.8z"/><path fill="currentColor" d="M476 399.1c0-3.9-3.1-7.1-7-7.1h-42c-3.8 0-7 3.2-7 7.1V484h-84.5c-4.1 0-7.5 3.1-7.5 7v42c0 3.8 3.4 7 7.5 7H420v84.9c0 3.9 3.2 7.1 7 7.1h42c3.9 0 7-3.2 7-7.1V540h84.5c4.1 0 7.5-3.2 7.5-7v-42c0-3.9-3.4-7-7.5-7H476zM560.5 704h-225c-4.1 0-7.5 3.2-7.5 7v42c0 3.8 3.4 7 7.5 7h225c4.1 0 7.5-3.2 7.5-7v-42c0-3.8-3.4-7-7.5-7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DingdingOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M573.7 252.5C422.5 197.4 201.3 96.7 201.3 96.7c-15.7-4.1-17.9 11.1-17.9 11.1c-5 61.1 33.6 160.5 53.6 182.8c19.9 22.3 319.1 113.7 319.1 113.7S326 357.9 270.5 341.9c-55.6-16-37.9 17.8-37.9 17.8c11.4 61.7 64.9 131.8 107.2 138.4c42.2 6.6 220.1 4 220.1 4s-35.5 4.1-93.2 11.9c-42.7 5.8-97 12.5-111.1 17.8c-33.1 12.5 24 62.6 24 62.6c84.7 76.8 129.7 50.5 129.7 50.5c33.3-10.7 61.4-18.5 85.2-24.2L565 743.1h84.6L603 928l205.3-271.9H700.8l22.3-38.7c.3.5.4.8.4.8S799.8 496.1 829 433.8l.6-1h-.1c5-10.8 8.6-19.7 10-25.8c17-71.3-114.5-99.4-265.8-154.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DingtalkCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m227 385.3c-1 4.2-3.5 10.4-7 17.8h.1l-.4.7c-20.3 43.1-73.1 127.7-73.1 127.7s-.1-.2-.3-.5l-15.5 26.8h74.5L575.1 810l32.3-128h-58.6l20.4-84.7c-16.5 3.9-35.9 9.4-59 16.8c0 0-31.2 18.2-89.9-35c0 0-39.6-34.7-16.6-43.4c9.8-3.7 47.4-8.4 77-12.3c40-5.4 64.6-8.2 64.6-8.2S422 517 392.7 512.5c-29.3-4.6-66.4-53.1-74.3-95.8c0 0-12.2-23.4 26.3-12.3c38.5 11.1 197.9 43.2 197.9 43.2s-207.4-63.3-221.2-78.7c-13.8-15.4-40.6-84.2-37.1-126.5c0 0 1.5-10.5 12.4-7.7c0 0 153.3 69.7 258.1 107.9c104.8 37.9 195.9 57.3 184.2 106.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DingtalkOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M573.7 252.5C422.5 197.4 201.3 96.7 201.3 96.7c-15.7-4.1-17.9 11.1-17.9 11.1c-5 61.1 33.6 160.5 53.6 182.8c19.9 22.3 319.1 113.7 319.1 113.7S326 357.9 270.5 341.9c-55.6-16-37.9 17.8-37.9 17.8c11.4 61.7 64.9 131.8 107.2 138.4c42.2 6.6 220.1 4 220.1 4s-35.5 4.1-93.2 11.9c-42.7 5.8-97 12.5-111.1 17.8c-33.1 12.5 24 62.6 24 62.6c84.7 76.8 129.7 50.5 129.7 50.5c33.3-10.7 61.4-18.5 85.2-24.2L565 743.1h84.6L603 928l205.3-271.9H700.8l22.3-38.7c.3.5.4.8.4.8S799.8 496.1 829 433.8l.6-1h-.1c5-10.8 8.6-19.7 10-25.8c17-71.3-114.5-99.4-265.8-154.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DingtalkSquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M739 449.3c-1 4.2-3.5 10.4-7 17.8h.1l-.4.7c-20.3 43.1-73.1 127.7-73.1 127.7s-.1-.2-.3-.5l-15.5 26.8h74.5L575.1 810l32.3-128h-58.6l20.4-84.7c-16.5 3.9-35.9 9.4-59 16.8c0 0-31.2 18.2-89.9-35c0 0-39.6-34.7-16.6-43.4c9.8-3.7 47.4-8.4 77-12.3c40-5.4 64.6-8.2 64.6-8.2S422 517 392.7 512.5c-29.3-4.6-66.4-53.1-74.3-95.8c0 0-12.2-23.4 26.3-12.3c38.5 11.1 197.9 43.2 197.9 43.2s-207.4-63.3-221.2-78.7c-13.8-15.4-40.6-84.2-37.1-126.5c0 0 1.5-10.5 12.4-7.7c0 0 153.3 69.7 258.1 107.9c104.8 37.9 195.9 57.3 184.2 106.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DisconnectOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832.6 191.4c-84.6-84.6-221.5-84.6-306 0l-96.9 96.9l51 51l96.9-96.9c53.8-53.8 144.6-59.5 204 0c59.5 59.5 53.8 150.2 0 204l-96.9 96.9l51.1 51.1l96.9-96.9c84.4-84.6 84.4-221.5-.1-306.1M446.5 781.6c-53.8 53.8-144.6 59.5-204 0c-59.5-59.5-53.8-150.2 0-204l96.9-96.9l-51.1-51.1l-96.9 96.9c-84.6 84.6-84.6 221.5 0 306s221.5 84.6 306 0l96.9-96.9l-51-51zM260.3 209.4a8.03 8.03 0 0 0-11.3 0L209.4 249a8.03 8.03 0 0 0 0 11.3l554.4 554.4c3.1 3.1 8.2 3.1 11.3 0l39.6-39.6c3.1-3.1 3.1-8.2 0-11.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DislikeFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M885.9 490.3c3.6-12 5.4-24.4 5.4-37c0-28.3-9.3-55.5-26.1-77.7c3.6-12 5.4-24.4 5.4-37c0-28.3-9.3-55.5-26.1-77.7c3.6-12 5.4-24.4 5.4-37c0-51.6-30.7-98.1-78.3-118.4a66.1 66.1 0 0 0-26.5-5.4H273v428h.3l85.8 310.8C372.9 889 418.9 924 470.9 924c29.7 0 57.4-11.8 77.9-33.4c20.5-21.5 31-49.7 29.5-79.4l-6-122.9h239.9c12.1 0 23.9-3.2 34.3-9.3c40.4-23.5 65.5-66.1 65.5-111c0-28.3-9.3-55.5-26.1-77.7M112 132v364c0 17.7 14.3 32 32 32h65V100h-65c-17.7 0-32 14.3-32 32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DislikeOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M885.9 490.3c3.6-12 5.4-24.4 5.4-37c0-28.3-9.3-55.5-26.1-77.7c3.6-12 5.4-24.4 5.4-37c0-28.3-9.3-55.5-26.1-77.7c3.6-12 5.4-24.4 5.4-37c0-51.6-30.7-98.1-78.3-118.4a66.1 66.1 0 0 0-26.5-5.4H144c-17.7 0-32 14.3-32 32v364c0 17.7 14.3 32 32 32h129.3l85.8 310.8C372.9 889 418.9 924 470.9 924c29.7 0 57.4-11.8 77.9-33.4c20.5-21.5 31-49.7 29.5-79.4l-6-122.9h239.9c12.1 0 23.9-3.2 34.3-9.3c40.4-23.5 65.5-66.1 65.5-111c0-28.3-9.3-55.5-26.1-77.7M184 456V172h81v284zm627.2 160.4H496.8l9.6 198.4c.6 11.9-4.7 23.1-14.6 30.5c-6.1 4.5-13.6 6.8-21.1 6.7a44.28 44.28 0 0 1-42.2-32.3L329 459.2V172h415.4a56.85 56.85 0 0 1 33.6 51.8c0 9.7-2.3 18.9-6.9 27.3l-13.9 25.4l21.9 19a56.76 56.76 0 0 1 19.6 43c0 9.7-2.3 18.9-6.9 27.3l-13.9 25.4l21.9 19a56.76 56.76 0 0 1 19.6 43c0 9.7-2.3 18.9-6.9 27.3l-14 25.5l21.9 19a56.76 56.76 0 0 1 19.6 43c0 19.1-11 37.5-28.8 48.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DislikeTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M273 100.1v428h.3zM820.4 525l-21.9-19l14-25.5a56.2 56.2 0 0 0 6.9-27.3c0-16.5-7.1-32.2-19.6-43l-21.9-19l13.9-25.4a56.2 56.2 0 0 0 6.9-27.3c0-16.5-7.1-32.2-19.6-43l-21.9-19l13.9-25.4a56.2 56.2 0 0 0 6.9-27.3c0-22.4-13.2-42.6-33.6-51.8H345v345.2c18.6 67.2 46.4 168 83.5 302.5a44.28 44.28 0 0 0 42.2 32.3c7.5.1 15-2.2 21.1-6.7c9.9-7.4 15.2-18.6 14.6-30.5l-9.6-198.4h314.4C829 605.5 840 587.1 840 568c0-16.5-7.1-32.2-19.6-43"/><path fill="currentColor" d="M112 132v364c0 17.7 14.3 32 32 32h65V100h-65c-17.7 0-32 14.3-32 32m773.9 358.3c3.6-12 5.4-24.4 5.4-37c0-28.3-9.3-55.5-26.1-77.7c3.6-12 5.4-24.4 5.4-37c0-28.3-9.3-55.5-26.1-77.7c3.6-12 5.4-24.4 5.4-37c0-51.6-30.7-98.1-78.3-118.4a66.1 66.1 0 0 0-26.5-5.4H273l.3 428l85.8 310.8C372.9 889 418.9 924 470.9 924c29.7 0 57.4-11.8 77.9-33.4c20.5-21.5 31-49.7 29.5-79.4l-6-122.9h239.9c12.1 0 23.9-3.2 34.3-9.3c40.4-23.5 65.5-66.1 65.5-111c0-28.3-9.3-55.5-26.1-77.7m-74.7 126.1H496.8l9.6 198.4c.6 11.9-4.7 23.1-14.6 30.5c-6.1 4.5-13.6 6.8-21.1 6.7a44.28 44.28 0 0 1-42.2-32.3c-37.1-134.4-64.9-235.2-83.5-302.5V172h399.4a56.85 56.85 0 0 1 33.6 51.8c0 9.7-2.3 18.9-6.9 27.3l-13.9 25.4l21.9 19a56.76 56.76 0 0 1 19.6 43c0 9.7-2.3 18.9-6.9 27.3l-13.9 25.4l21.9 19a56.76 56.76 0 0 1 19.6 43c0 9.7-2.3 18.9-6.9 27.3l-14 25.5l21.9 19a56.76 56.76 0 0 1 19.6 43c0 19.1-11 37.5-28.8 48.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DollarCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m22.3 665.2l.2 31.7c0 4.4-3.6 8.1-8 8.1h-28.4c-4.4 0-8-3.6-8-8v-31.4C401.3 723 359.5 672.4 355 617.4c-.4-4.7 3.3-8.7 8-8.7h46.2c3.9 0 7.3 2.8 7.9 6.6c5.1 31.7 29.8 55.4 74.1 61.3V533.9l-24.7-6.3c-52.3-12.5-102.1-45.1-102.1-112.7c0-72.9 55.4-112.1 126.2-119v-33c0-4.4 3.6-8 8-8h28.1c4.4 0 8 3.6 8 8v32.7c68.5 6.9 119.9 46.9 125.9 109.2c.5 4.7-3.2 8.8-8 8.8h-44.9c-4 0-7.4-3-7.9-6.9c-4-29.2-27.4-53-65.5-58.2v134.3l25.4 5.9c64.8 16 108.9 47 108.9 116.4c0 75.3-56 117.3-134.3 124.1M426.6 410.3c0 25.4 15.7 45.1 49.5 57.3c4.7 1.9 9.4 3.4 15 5v-124c-36.9 4.7-64.5 25.4-64.5 61.7m116.5 135.2c-2.8-.6-5.6-1.3-8.8-2.2V677c42.6-3.8 72-27.2 72-66.4c0-30.7-15.9-50.7-63.2-65.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DollarCircleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372m47.7-395.2l-25.4-5.9V348.6c38 5.2 61.5 29 65.5 58.2c.5 4 3.9 6.9 7.9 6.9h44.9c4.7 0 8.4-4.1 8-8.8c-6.1-62.3-57.4-102.3-125.9-109.2V263c0-4.4-3.6-8-8-8h-28.1c-4.4 0-8 3.6-8 8v33c-70.8 6.9-126.2 46-126.2 119c0 67.6 49.8 100.2 102.1 112.7l24.7 6.3v142.7c-44.2-5.9-69-29.5-74.1-61.3c-.6-3.8-4-6.6-7.9-6.6H363c-4.7 0-8.4 4-8 8.7c4.5 55 46.2 105.6 135.2 112.1V761c0 4.4 3.6 8 8 8h28.4c4.4 0 8-3.6 8-8.1l-.2-31.7c78.3-6.9 134.3-48.8 134.3-124c-.1-69.4-44.2-100.4-109-116.4m-68.6-16.2c-5.6-1.6-10.3-3.1-15-5c-33.8-12.2-49.5-31.9-49.5-57.3c0-36.3 27.5-57 64.5-61.7zM534.3 677V543.3c3.1.9 5.9 1.6 8.8 2.2c47.3 14.4 63.2 34.4 63.2 65.1c0 39.1-29.4 62.6-72 66.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DollarCircleTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" fill-opacity=".15" d="M426.6 410.3c0 25.4 15.7 45.1 49.5 57.3c4.7 1.9 9.4 3.4 15 5v-124c-37 4.7-64.5 25.4-64.5 61.7m116.5 135.2c-2.9-.6-5.7-1.3-8.8-2.2V677c42.6-3.8 72-27.3 72-66.4c0-30.7-15.9-50.7-63.2-65.1"/><path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372m22.4 589.2l.2 31.7c0 4.5-3.6 8.1-8 8.1h-28.4c-4.4 0-8-3.6-8-8v-31.4c-89-6.5-130.7-57.1-135.2-112.1c-.4-4.7 3.3-8.7 8-8.7h46.2c3.9 0 7.3 2.8 7.9 6.6c5.1 31.8 29.9 55.4 74.1 61.3V534l-24.7-6.3c-52.3-12.5-102.1-45.1-102.1-112.7c0-73 55.4-112.1 126.2-119v-33c0-4.4 3.6-8 8-8h28.1c4.4 0 8 3.6 8 8v32.7c68.5 6.9 119.8 46.9 125.9 109.2a8.1 8.1 0 0 1-8 8.8h-44.9c-4 0-7.4-2.9-7.9-6.9c-4-29.2-27.5-53-65.5-58.2v134.3l25.4 5.9c64.8 16 108.9 47 109 116.4c0 75.2-56 117.1-134.3 124"/><path fill="currentColor" d="m559.7 488.8l-25.4-5.9V348.6c38 5.2 61.5 29 65.5 58.2c.5 4 3.9 6.9 7.9 6.9h44.9c4.7 0 8.4-4.1 8-8.8c-6.1-62.3-57.4-102.3-125.9-109.2V263c0-4.4-3.6-8-8-8h-28.1c-4.4 0-8 3.6-8 8v33c-70.8 6.9-126.2 46-126.2 119c0 67.6 49.8 100.2 102.1 112.7l24.7 6.3v142.7c-44.2-5.9-69-29.5-74.1-61.3c-.6-3.8-4-6.6-7.9-6.6H363c-4.7 0-8.4 4-8 8.7c4.5 55 46.2 105.6 135.2 112.1V761c0 4.4 3.6 8 8 8h28.4c4.4 0 8-3.6 8-8.1l-.2-31.7c78.3-6.9 134.3-48.8 134.3-124c-.1-69.4-44.2-100.4-109-116.4m-68.6-16.2c-5.6-1.6-10.3-3.1-15-5c-33.8-12.2-49.5-31.9-49.5-57.3c0-36.3 27.5-57 64.5-61.7zM534.3 677V543.3c3.1.9 5.9 1.6 8.8 2.2c47.3 14.4 63.2 34.4 63.2 65.1c0 39.1-29.4 62.6-72 66.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DollarOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372m47.7-395.2l-25.4-5.9V348.6c38 5.2 61.5 29 65.5 58.2c.5 4 3.9 6.9 7.9 6.9h44.9c4.7 0 8.4-4.1 8-8.8c-6.1-62.3-57.4-102.3-125.9-109.2V263c0-4.4-3.6-8-8-8h-28.1c-4.4 0-8 3.6-8 8v33c-70.8 6.9-126.2 46-126.2 119c0 67.6 49.8 100.2 102.1 112.7l24.7 6.3v142.7c-44.2-5.9-69-29.5-74.1-61.3c-.6-3.8-4-6.6-7.9-6.6H363c-4.7 0-8.4 4-8 8.7c4.5 55 46.2 105.6 135.2 112.1V761c0 4.4 3.6 8 8 8h28.4c4.4 0 8-3.6 8-8.1l-.2-31.7c78.3-6.9 134.3-48.8 134.3-124c-.1-69.4-44.2-100.4-109-116.4m-68.6-16.2c-5.6-1.6-10.3-3.1-15-5c-33.8-12.2-49.5-31.9-49.5-57.3c0-36.3 27.5-57 64.5-61.7zM534.3 677V543.3c3.1.9 5.9 1.6 8.8 2.2c47.3 14.4 63.2 34.4 63.2 65.1c0 39.1-29.4 62.6-72 66.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DollarTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" fill-opacity=".15" d="M426.6 410.3c0 25.4 15.7 45.1 49.5 57.3c4.7 1.9 9.4 3.4 15 5v-124c-37 4.7-64.5 25.4-64.5 61.7m116.5 135.2c-2.9-.6-5.7-1.3-8.8-2.2V677c42.6-3.8 72-27.3 72-66.4c0-30.7-15.9-50.7-63.2-65.1"/><path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372m22.4 589.2l.2 31.7c0 4.5-3.6 8.1-8 8.1h-28.4c-4.4 0-8-3.6-8-8v-31.4c-89-6.5-130.7-57.1-135.2-112.1c-.4-4.7 3.3-8.7 8-8.7h46.2c3.9 0 7.3 2.8 7.9 6.6c5.1 31.8 29.9 55.4 74.1 61.3V534l-24.7-6.3c-52.3-12.5-102.1-45.1-102.1-112.7c0-73 55.4-112.1 126.2-119v-33c0-4.4 3.6-8 8-8h28.1c4.4 0 8 3.6 8 8v32.7c68.5 6.9 119.8 46.9 125.9 109.2a8.1 8.1 0 0 1-8 8.8h-44.9c-4 0-7.4-2.9-7.9-6.9c-4-29.2-27.5-53-65.5-58.2v134.3l25.4 5.9c64.8 16 108.9 47 109 116.4c0 75.2-56 117.1-134.3 124"/><path fill="currentColor" d="m559.7 488.8l-25.4-5.9V348.6c38 5.2 61.5 29 65.5 58.2c.5 4 3.9 6.9 7.9 6.9h44.9c4.7 0 8.4-4.1 8-8.8c-6.1-62.3-57.4-102.3-125.9-109.2V263c0-4.4-3.6-8-8-8h-28.1c-4.4 0-8 3.6-8 8v33c-70.8 6.9-126.2 46-126.2 119c0 67.6 49.8 100.2 102.1 112.7l24.7 6.3v142.7c-44.2-5.9-69-29.5-74.1-61.3c-.6-3.8-4-6.6-7.9-6.6H363c-4.7 0-8.4 4-8 8.7c4.5 55 46.2 105.6 135.2 112.1V761c0 4.4 3.6 8 8 8h28.4c4.4 0 8-3.6 8-8.1l-.2-31.7c78.3-6.9 134.3-48.8 134.3-124c-.1-69.4-44.2-100.4-109-116.4m-68.6-16.2c-5.6-1.6-10.3-3.1-15-5c-33.8-12.2-49.5-31.9-49.5-57.3c0-36.3 27.5-57 64.5-61.7zM534.3 677V543.3c3.1.9 5.9 1.6 8.8 2.2c47.3 14.4 63.2 34.4 63.2 65.1c0 39.1-29.4 62.6-72 66.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotChartOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M888 792H200V168c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v688c0 4.4 3.6 8 8 8h752c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8M288 604a64 64 0 1 0 128 0a64 64 0 1 0-128 0m118-224a48 48 0 1 0 96 0a48 48 0 1 0-96 0m158 228a96 96 0 1 0 192 0a96 96 0 1 0-192 0m148-314a56 56 0 1 0 112 0a56 56 0 1 0-112 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleLeftOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m272.9 512l265.4-339.1c4.1-5.2.4-12.9-6.3-12.9h-77.3c-4.9 0-9.6 2.3-12.6 6.1L186.8 492.3a31.99 31.99 0 0 0 0 39.5l255.3 326.1c3 3.9 7.7 6.1 12.6 6.1H532c6.7 0 10.4-7.7 6.3-12.9zm304 0l265.4-339.1c4.1-5.2.4-12.9-6.3-12.9h-77.3c-4.9 0-9.6 2.3-12.6 6.1L490.8 492.3a31.99 31.99 0 0 0 0 39.5l255.3 326.1c3 3.9 7.7 6.1 12.6 6.1H836c6.7 0 10.4-7.7 6.3-12.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleRightOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M533.2 492.3L277.9 166.1c-3-3.9-7.7-6.1-12.6-6.1H188c-6.7 0-10.4 7.7-6.3 12.9L447.1 512L181.7 851.1A7.98 7.98 0 0 0 188 864h77.3c4.9 0 9.6-2.3 12.6-6.1l255.3-326.1c9.1-11.7 9.1-27.9 0-39.5m304 0L581.9 166.1c-3-3.9-7.7-6.1-12.6-6.1H492c-6.7 0-10.4 7.7-6.3 12.9L751.1 512L485.7 851.1A7.98 7.98 0 0 0 492 864h77.3c4.9 0 9.6-2.3 12.6-6.1l255.3-326.1c9.1-11.7 9.1-27.9 0-39.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m184.5 353.7l-178 246a7.95 7.95 0 0 1-12.9 0l-178-246c-3.8-5.3 0-12.7 6.5-12.7H381c10.2 0 19.9 4.9 25.9 13.2L512 563.6l105.2-145.4c6-8.3 15.6-13.2 25.9-13.2H690c6.5 0 10.3 7.4 6.5 12.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownCircleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M690 405h-46.9c-10.2 0-19.9 4.9-25.9 13.2L512 563.6L406.8 418.2c-6-8.3-15.6-13.2-25.9-13.2H334c-6.5 0-10.3 7.4-6.5 12.7l178 246c3.2 4.4 9.7 4.4 12.9 0l178-246c3.9-5.3.1-12.7-6.4-12.7"/><path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownCircleTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372m184.4 277.7l-178 246a7.95 7.95 0 0 1-12.9 0l-178-246c-3.8-5.3 0-12.7 6.5-12.7h46.9c10.3 0 19.9 4.9 25.9 13.2L512 563.6l105.2-145.4c6-8.3 15.7-13.2 25.9-13.2H690c6.5 0 10.3 7.4 6.4 12.7"/><path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" d="M690 405h-46.9c-10.2 0-19.9 4.9-25.9 13.2L512 563.6L406.8 418.2c-6-8.3-15.6-13.2-25.9-13.2H334c-6.5 0-10.3 7.4-6.5 12.7l178 246c3.2 4.4 9.7 4.4 12.9 0l178-246c3.9-5.3.1-12.7-6.4-12.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M884 256h-75c-5.1 0-9.9 2.5-12.9 6.6L512 654.2L227.9 262.6c-3-4.1-7.8-6.6-12.9-6.6h-75c-6.5 0-10.3 7.4-6.5 12.7l352.6 486.1c12.8 17.6 39 17.6 51.7 0l352.6-486.1c3.9-5.3.1-12.7-6.4-12.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownSquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M696.5 412.7l-178 246a7.95 7.95 0 0 1-12.9 0l-178-246c-3.8-5.3 0-12.7 6.5-12.7H381c10.2 0 19.9 4.9 25.9 13.2L512 558.6l105.2-145.4c6-8.3 15.6-13.2 25.9-13.2H690c6.5 0 10.3 7.4 6.5 12.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownSquareOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M505.5 658.7c3.2 4.4 9.7 4.4 12.9 0l178-246c3.8-5.3 0-12.7-6.5-12.7H643c-10.2 0-19.9 4.9-25.9 13.2L512 558.6L406.8 413.2c-6-8.3-15.6-13.2-25.9-13.2H334c-6.5 0-10.3 7.4-6.5 12.7z"/><path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownSquareTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/><path fill="currentColor" fill-opacity=".15" d="M184 840h656V184H184zm150-440h46.9c10.3 0 19.9 4.9 25.9 13.2L512 558.6l105.2-145.4c6-8.3 15.7-13.2 25.9-13.2H690c6.5 0 10.3 7.4 6.4 12.7l-178 246a7.95 7.95 0 0 1-12.9 0l-178-246c-3.8-5.3 0-12.7 6.5-12.7"/><path fill="currentColor" d="M505.5 658.7c3.2 4.4 9.7 4.4 12.9 0l178-246c3.9-5.3.1-12.7-6.4-12.7h-46.9c-10.2 0-19.9 4.9-25.9 13.2L512 558.6L406.8 413.2c-6-8.3-15.6-13.2-25.9-13.2H334c-6.5 0-10.3 7.4-6.5 12.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M505.7 661a8 8 0 0 0 12.6 0l112-141.7c4.1-5.2.4-12.9-6.3-12.9h-74.1V168c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8v338.3H400c-6.7 0-10.4 7.7-6.3 12.9zM878 626h-60c-4.4 0-8 3.6-8 8v154H214V634c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8v198c0 17.7 14.3 32 32 32h684c17.7 0 32-14.3 32-32V634c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DragOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M909.3 506.3L781.7 405.6a7.23 7.23 0 0 0-11.7 5.7V476H548V254h64.8c6 0 9.4-7 5.7-11.7L517.7 114.7a7.14 7.14 0 0 0-11.3 0L405.6 242.3a7.23 7.23 0 0 0 5.7 11.7H476v222H254v-64.8c0-6-7-9.4-11.7-5.7L114.7 506.3a7.14 7.14 0 0 0 0 11.3l127.5 100.8c4.7 3.7 11.7.4 11.7-5.7V548h222v222h-64.8c-6 0-9.4 7-5.7 11.7l100.8 127.5c2.9 3.7 8.5 3.7 11.3 0l100.8-127.5c3.7-4.7.4-11.7-5.7-11.7H548V548h222v64.8c0 6 7 9.4 11.7 5.7l127.5-100.8a7.3 7.3 0 0 0 .1-11.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DribbbleCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M675.1 328.3a245.2 245.2 0 0 0-220.8-55.1c6.8 9.1 51.5 69.9 91.8 144c87.5-32.8 124.5-82.6 129-88.9M554 552.8c-138.7 48.3-188.6 144.6-193 153.6c41.7 32.5 94.1 51.9 151 51.9c34.1 0 66.6-6.9 96.1-19.5c-3.7-21.6-17.9-96.8-52.5-186.6zm47.7-11.9c32.2 88.4 45.3 160.4 47.8 175.4c55.2-37.3 94.5-96.4 105.4-164.9c-8.4-2.6-76.1-22.8-153.2-10.5M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 736c-158.8 0-288-129.2-288-288s129.2-288 288-288s288 129.2 288 288s-129.2 288-288 288m53.1-346.2c5.7 11.7 11.2 23.6 16.3 35.6c1.8 4.2 3.6 8.4 5.3 12.7c81.8-10.3 163.2 6.2 171.3 7.9c-.5-58.1-21.3-111.4-55.5-153.3c-5.3 7.1-46.5 60-137.4 97.1M498.6 432c-40.8-72.5-84.7-133.4-91.2-142.3c-68.8 32.5-120.3 95.9-136.2 172.2c11 .2 112.4.7 227.4-29.9m30.6 82.5c3.2-1 6.4-2 9.7-2.9c-6.2-14-12.9-28-19.9-41.7c-122.8 36.8-242.1 35.2-252.8 35c-.1 2.5-.1 5-.1 7.5c0 63.2 23.9 120.9 63.2 164.5c5.5-9.6 73-121.4 199.9-162.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DribbbleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 96C282.6 96 96 282.6 96 512s186.6 416 416 416s416-186.6 416-416S741.4 96 512 96m275.1 191.8c49.5 60.5 79.5 137.5 80.2 221.4c-11.7-2.5-129.2-26.3-247.4-11.4c-2.5-6.1-5-12.2-7.6-18.3c-7.4-17.3-15.3-34.6-23.6-51.5C720 374.3 779.6 298 787.1 287.8M512 157.2c90.3 0 172.8 33.9 235.5 89.5c-6.4 9.1-59.9 81-186.2 128.4c-58.2-107-122.7-194.8-132.6-208c27.3-6.6 55.2-9.9 83.3-9.9M360.9 191c9.4 12.8 72.9 100.9 131.7 205.5C326.4 440.6 180 440 164.1 439.8c23.1-110.3 97.4-201.9 196.8-248.8M156.7 512.5c0-3.6.1-7.3.2-10.9c15.5.3 187.7 2.5 365.2-50.6c10.2 19.9 19.9 40.1 28.8 60.3c-4.7 1.3-9.4 2.7-14 4.2C353.6 574.9 256.1 736.4 248 750.1c-56.7-63-91.3-146.3-91.3-237.6M512 867.8c-82.2 0-157.9-28-218.1-75c6.4-13.1 78.3-152 278.7-221.9l2.3-.8c49.9 129.6 70.5 238.3 75.8 269.5A350.46 350.46 0 0 1 512 867.8m198.5-60.7c-3.6-21.6-22.5-125.6-69-253.3C752.9 536 850.7 565.2 862.8 569c-15.8 98.8-72.5 184.2-152.3 238.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DribbbleSquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M498.6 432c-40.8-72.5-84.7-133.4-91.2-142.3c-68.8 32.5-120.3 95.9-136.2 172.2c11 .2 112.4.7 227.4-29.9m66.5 21.8c5.7 11.7 11.2 23.6 16.3 35.6c1.8 4.2 3.6 8.4 5.3 12.7c81.8-10.3 163.2 6.2 171.3 7.9c-.5-58.1-21.3-111.4-55.5-153.3c-5.3 7.1-46.5 60-137.4 97.1M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M512 800c-158.8 0-288-129.2-288-288s129.2-288 288-288s288 129.2 288 288s-129.2 288-288 288m89.7-259.1c32.2 88.4 45.3 160.4 47.8 175.4c55.2-37.3 94.5-96.4 105.4-164.9c-8.4-2.6-76.1-22.8-153.2-10.5m-72.5-26.4c3.2-1 6.4-2 9.7-2.9c-6.2-14-12.9-28-19.9-41.7c-122.8 36.8-242.1 35.2-252.8 35c-.1 2.5-.1 5-.1 7.5c0 63.2 23.9 120.9 63.2 164.5c5.5-9.6 73-121.4 199.9-162.4m145.9-186.2a245.2 245.2 0 0 0-220.8-55.1c6.8 9.1 51.5 69.9 91.8 144c87.5-32.8 124.5-82.6 129-88.9M554 552.8c-138.7 48.3-188.6 144.6-193 153.6c41.7 32.5 94.1 51.9 151 51.9c34.1 0 66.6-6.9 96.1-19.5c-3.7-21.6-17.9-96.8-52.5-186.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DribbbleSquareOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M498.6 432c-40.8-72.5-84.7-133.4-91.2-142.3c-68.8 32.5-120.3 95.9-136.2 172.2c11 .2 112.4.7 227.4-29.9m66.5 21.8c5.7 11.7 11.2 23.6 16.3 35.6c1.8 4.2 3.6 8.4 5.3 12.7c81.8-10.3 163.2 6.2 171.3 7.9c-.5-58.1-21.3-111.4-55.5-153.3c-5.3 7.1-46.5 60-137.4 97.1M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M512 800c-158.8 0-288-129.2-288-288s129.2-288 288-288s288 129.2 288 288s-129.2 288-288 288m89.7-259.1c32.2 88.4 45.3 160.4 47.8 175.4c55.2-37.3 94.5-96.4 105.4-164.9c-8.4-2.6-76.1-22.8-153.2-10.5m-72.5-26.4c3.2-1 6.4-2 9.7-2.9c-6.2-14-12.9-28-19.9-41.7c-122.8 36.8-242.1 35.2-252.8 35c-.1 2.5-.1 5-.1 7.5c0 63.2 23.9 120.9 63.2 164.5c5.5-9.6 73-121.4 199.9-162.4m145.9-186.2a245.2 245.2 0 0 0-220.8-55.1c6.8 9.1 51.5 69.9 91.8 144c87.5-32.8 124.5-82.6 129-88.9M554 552.8c-138.7 48.3-188.6 144.6-193 153.6c41.7 32.5 94.1 51.9 151 51.9c34.1 0 66.6-6.9 96.1-19.5c-3.7-21.6-17.9-96.8-52.5-186.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DropboxCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m512.3 361.7l-151.8 93.8l151.8 93.9l151.5-93.9zM512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m151.2 595.5L512.6 750l-151-90.5v-33.1l45.4 29.4l105.6-87.7l105.6 87.7l45.1-29.4v33.1zm-45.6-22.4l-105.3-87.7L407 637.1l-151-99.2l104.5-82.4L256 371.2L407 274l105.3 87.7L617.6 274L768 372.1l-104.2 83.5L768 539z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DropboxOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m64 556.9l264.2 173.5L512.5 577L246.8 412.7zm896-290.3L696.8 95L512.5 248.5l265.2 164.2L512.5 577l184.3 153.4L960 558.8L777.7 412.7zM513 609.8L328.2 763.3l-79.4-51.5v57.8L513 928l263.7-158.4v-57.8l-78.9 51.5zM328.2 95L64 265.1l182.8 147.6l265.7-164.2zM64 556.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DropboxSquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M663.2 659.5L512.6 750l-151-90.5v-33.1l45.4 29.4l105.6-87.7l105.6 87.7l45.1-29.4v33.1zm-45.6-22.4l-105.3-87.7L407 637.1l-151-99.2l104.5-82.4L256 371.2L407 274l105.3 87.7L617.6 274L768 372.1l-104.2 83.5L768 539zM512.3 361.7l-151.8 93.8l151.8 93.9l151.5-93.9zm151.5 93.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 836H144c-17.7 0-32 14.3-32 32v36c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-36c0-17.7-14.3-32-32-32m-622.3-84c2 0 4-.2 6-.5L431.9 722c2-.4 3.9-1.3 5.3-2.8l423.9-423.9a9.96 9.96 0 0 0 0-14.1L694.9 114.9c-1.9-1.9-4.4-2.9-7.1-2.9s-5.2 1-7.1 2.9L256.8 538.8c-1.5 1.5-2.4 3.3-2.8 5.3l-29.5 168.2a33.5 33.5 0 0 0 9.4 29.8c6.6 6.4 14.9 9.9 23.8 9.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M257.7 752c2 0 4-.2 6-.5L431.9 722c2-.4 3.9-1.3 5.3-2.8l423.9-423.9a9.96 9.96 0 0 0 0-14.1L694.9 114.9c-1.9-1.9-4.4-2.9-7.1-2.9s-5.2 1-7.1 2.9L256.8 538.8c-1.5 1.5-2.4 3.3-2.8 5.3l-29.5 168.2a33.5 33.5 0 0 0 9.4 29.8c6.6 6.4 14.9 9.9 23.8 9.9m67.4-174.4L687.8 215l73.3 73.3l-362.7 362.6l-88.9 15.7zM880 836H144c-17.7 0-32 14.3-32 32v36c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-36c0-17.7-14.3-32-32-32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M761.1 288.3L687.8 215L325.1 577.6l-15.6 89l88.9-15.7z"/><path fill="currentColor" d="M880 836H144c-17.7 0-32 14.3-32 32v36c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-36c0-17.7-14.3-32-32-32m-622.3-84c2 0 4-.2 6-.5L431.9 722c2-.4 3.9-1.3 5.3-2.8l423.9-423.9a9.96 9.96 0 0 0 0-14.1L694.9 114.9c-1.9-1.9-4.4-2.9-7.1-2.9s-5.2 1-7.1 2.9L256.8 538.8c-1.5 1.5-2.4 3.3-2.8 5.3l-29.5 168.2a33.5 33.5 0 0 0 9.4 29.8c6.6 6.4 14.9 9.9 23.8 9.9m67.4-174.4L687.8 215l73.3 73.3l-362.7 362.6l-88.9 15.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M176 511a56 56 0 1 0 112 0a56 56 0 1 0-112 0m280 0a56 56 0 1 0 112 0a56 56 0 1 0-112 0m280 0a56 56 0 1 0 112 0a56 56 0 1 0-112 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnterOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M864 170h-60c-4.4 0-8 3.6-8 8v518H310v-73c0-6.7-7.8-10.5-13-6.3l-141.9 112a8 8 0 0 0 0 12.6l141.9 112c5.3 4.2 13 .4 13-6.3v-75h498c35.3 0 64-28.7 64-64V178c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvironmentFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 327c-29.9 0-58 11.6-79.2 32.8A111.6 111.6 0 0 0 400 439c0 29.9 11.7 58 32.8 79.2A111.6 111.6 0 0 0 512 551c29.9 0 58-11.7 79.2-32.8C612.4 497 624 468.9 624 439c0-29.9-11.6-58-32.8-79.2S541.9 327 512 327m342.6-37.9a362.49 362.49 0 0 0-79.9-115.7a370.83 370.83 0 0 0-118.2-77.8C610.7 76.6 562.1 67 512 67c-50.1 0-98.7 9.6-144.5 28.5c-44.3 18.3-84 44.5-118.2 77.8A363.6 363.6 0 0 0 169.4 289c-19.5 45-29.4 92.8-29.4 142c0 70.6 16.9 140.9 50.1 208.7c26.7 54.5 64 107.6 111 158.1c80.3 86.2 164.5 138.9 188.4 153a43.9 43.9 0 0 0 22.4 6.1c7.8 0 15.5-2 22.4-6.1c23.9-14.1 108.1-66.8 188.4-153c47-50.4 84.3-103.6 111-158.1C867.1 572 884 501.8 884 431.1c0-49.2-9.9-97-29.4-142M512 615c-97.2 0-176-78.8-176-176s78.8-176 176-176s176 78.8 176 176s-78.8 176-176 176"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvironmentOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M854.6 289.1a362.49 362.49 0 0 0-79.9-115.7a370.83 370.83 0 0 0-118.2-77.8C610.7 76.6 562.1 67 512 67c-50.1 0-98.7 9.6-144.5 28.5c-44.3 18.3-84 44.5-118.2 77.8A363.6 363.6 0 0 0 169.4 289c-19.5 45-29.4 92.8-29.4 142c0 70.6 16.9 140.9 50.1 208.7c26.7 54.5 64 107.6 111 158.1c80.3 86.2 164.5 138.9 188.4 153a43.9 43.9 0 0 0 22.4 6.1c7.8 0 15.5-2 22.4-6.1c23.9-14.1 108.1-66.8 188.4-153c47-50.4 84.3-103.6 111-158.1C867.1 572 884 501.8 884 431.1c0-49.2-9.9-97-29.4-142M512 880.2c-65.9-41.9-300-207.8-300-449.1c0-77.9 31.1-151.1 87.6-206.3C356.3 169.5 431.7 139 512 139s155.7 30.5 212.4 85.9C780.9 280 812 353.2 812 431.1c0 241.3-234.1 407.2-300 449.1m0-617.2c-97.2 0-176 78.8-176 176s78.8 176 176 176s176-78.8 176-176s-78.8-176-176-176m79.2 255.2A111.6 111.6 0 0 1 512 551c-29.9 0-58-11.7-79.2-32.8A111.6 111.6 0 0 1 400 439c0-29.9 11.7-58 32.8-79.2C454 338.6 482.1 327 512 327c29.9 0 58 11.6 79.2 32.8C612.4 381 624 409.1 624 439c0 29.9-11.6 58-32.8 79.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvironmentTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M724.4 224.9C667.7 169.5 592.3 139 512 139s-155.7 30.5-212.4 85.8C243.1 280 212 353.2 212 431.1c0 241.3 234.1 407.2 300 449.1c65.9-41.9 300-207.8 300-449.1c0-77.9-31.1-151.1-87.6-206.2M512 615c-97.2 0-176-78.8-176-176s78.8-176 176-176s176 78.8 176 176s-78.8 176-176 176"/><path fill="currentColor" d="M512 263c-97.2 0-176 78.8-176 176s78.8 176 176 176s176-78.8 176-176s-78.8-176-176-176m79.2 255.2A111.6 111.6 0 0 1 512 551c-29.9 0-58-11.7-79.2-32.8A111.6 111.6 0 0 1 400 439c0-29.9 11.7-58 32.8-79.2C454 338.6 482.1 327 512 327c29.9 0 58 11.6 79.2 32.8S624 409.1 624 439c0 29.9-11.6 58-32.8 79.2"/><path fill="currentColor" d="M854.6 289.1a362.49 362.49 0 0 0-79.9-115.7a370.83 370.83 0 0 0-118.2-77.8C610.7 76.6 562.1 67 512 67c-50.1 0-98.7 9.6-144.5 28.5c-44.3 18.3-84 44.5-118.2 77.8A363.6 363.6 0 0 0 169.4 289c-19.5 45-29.4 92.8-29.4 142c0 70.6 16.9 140.9 50.1 208.7c26.7 54.5 64 107.6 111 158.1c80.3 86.2 164.5 138.9 188.4 153a43.9 43.9 0 0 0 22.4 6.1c7.8 0 15.5-2 22.4-6.1c23.9-14.1 108.1-66.8 188.4-153c47-50.4 84.3-103.6 111-158.1C867.1 572 884 501.8 884 431.1c0-49.2-9.9-97-29.4-142M512 880.2c-65.9-41.9-300-207.8-300-449.1c0-77.9 31.1-151.1 87.6-206.3C356.3 169.5 431.7 139 512 139s155.7 30.5 212.4 85.9C780.9 280 812 353.2 812 431.1c0 241.3-234.1 407.2-300 449.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EuroCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m63.5 375.8c4.4 0 8 3.6 8 8V475c0 4.4-3.6 8-8 8h-136c-.3 4.4-.3 9.1-.3 13.8v36h136.2c4.4 0 8 3.6 8 8V568c0 4.4-3.6 8-8 8H444.9c15.3 62 61.3 98.6 129.8 98.6c19.9 0 37.1-1.2 51.8-4.1c4.9-1 9.5 2.8 9.5 7.8v42.8c0 3.8-2.7 7-6.4 7.8c-15.9 3.4-34.3 5.1-55.3 5.1c-109.8 0-183-58.8-200.2-158H344c-4.4 0-8-3.6-8-8v-27.2c0-4.4 3.6-8 8-8h26.1v-36.9c0-4.4 0-8.8.3-12.8H344c-4.4 0-8-3.6-8-8v-27.2c0-4.4 3.6-8 8-8h31.7c19.7-94.2 92-149.9 198.6-149.9c20.9 0 39.4 1.9 55.3 5.4c3.7.8 6.3 4 6.3 7.8V346h.1c0 5.1-4.6 8.8-9.6 7.8c-14.7-2.9-31.8-4.4-51.7-4.4c-65.4 0-110.4 33.5-127.6 90.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EuroCircleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372m117.7-588.6c-15.9-3.5-34.4-5.4-55.3-5.4c-106.7 0-178.9 55.7-198.6 149.9H344c-4.4 0-8 3.6-8 8v27.2c0 4.4 3.6 8 8 8h26.4c-.3 4.1-.3 8.4-.3 12.8v36.9H344c-4.4 0-8 3.6-8 8V568c0 4.4 3.6 8 8 8h30.2c17.2 99.2 90.4 158 200.2 158c20.9 0 39.4-1.7 55.3-5.1c3.7-.8 6.4-4 6.4-7.8v-42.8c0-5-4.6-8.8-9.5-7.8c-14.7 2.8-31.9 4.1-51.8 4.1c-68.5 0-114.5-36.6-129.8-98.6h130.6c4.4 0 8-3.6 8-8v-27.2c0-4.4-3.6-8-8-8H439.2v-36c0-4.7 0-9.4.3-13.8h135.9c4.4 0 8-3.6 8-8v-27.2c0-4.4-3.6-8-8-8H447.1c17.2-56.9 62.3-90.4 127.6-90.4c19.9 0 37.1 1.5 51.7 4.4a8 8 0 0 0 9.6-7.8v-42.8c0-3.8-2.6-7-6.3-7.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EuroCircleTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372m117.1 581.1c0 3.8-2.7 7-6.4 7.8c-15.9 3.4-34.4 5.1-55.3 5.1c-109.8 0-183-58.8-200.2-158H337c-4.4 0-8-3.6-8-8v-27.2c0-4.4 3.6-8 8-8h26.1v-36.9c0-4.4 0-8.7.3-12.8H337c-4.4 0-8-3.6-8-8v-27.2c0-4.4 3.6-8 8-8h31.8C388.5 345.7 460.7 290 567.4 290c20.9 0 39.4 1.9 55.3 5.4c3.7.8 6.3 4 6.3 7.8V346a8 8 0 0 1-9.6 7.8c-14.6-2.9-31.8-4.4-51.7-4.4c-65.3 0-110.4 33.5-127.6 90.4h128.3c4.4 0 8 3.6 8 8V475c0 4.4-3.6 8-8 8H432.5c-.3 4.4-.3 9.1-.3 13.8v36h136.4c4.4 0 8 3.6 8 8V568c0 4.4-3.6 8-8 8H438c15.3 62 61.3 98.6 129.8 98.6c19.9 0 37.1-1.3 51.8-4.1c4.9-1 9.5 2.8 9.5 7.8z"/><path fill="currentColor" d="M619.6 670.5c-14.7 2.8-31.9 4.1-51.8 4.1c-68.5 0-114.5-36.6-129.8-98.6h130.6c4.4 0 8-3.6 8-8v-27.2c0-4.4-3.6-8-8-8H432.2v-36c0-4.7 0-9.4.3-13.8h135.9c4.4 0 8-3.6 8-8v-27.2c0-4.4-3.6-8-8-8H440.1c17.2-56.9 62.3-90.4 127.6-90.4c19.9 0 37.1 1.5 51.7 4.4a8 8 0 0 0 9.6-7.8v-42.8c0-3.8-2.6-7-6.3-7.8c-15.9-3.5-34.4-5.4-55.3-5.4c-106.7 0-178.9 55.7-198.6 149.9H337c-4.4 0-8 3.6-8 8v27.2c0 4.4 3.6 8 8 8h26.4c-.3 4.1-.3 8.4-.3 12.8v36.9H337c-4.4 0-8 3.6-8 8V568c0 4.4 3.6 8 8 8h30.2c17.2 99.2 90.4 158 200.2 158c20.9 0 39.4-1.7 55.3-5.1c3.7-.8 6.4-4 6.4-7.8v-42.8c0-5-4.6-8.8-9.5-7.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EuroOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372m117.7-588.6c-15.9-3.5-34.4-5.4-55.3-5.4c-106.7 0-178.9 55.7-198.6 149.9H344c-4.4 0-8 3.6-8 8v27.2c0 4.4 3.6 8 8 8h26.4c-.3 4.1-.3 8.4-.3 12.8v36.9H344c-4.4 0-8 3.6-8 8V568c0 4.4 3.6 8 8 8h30.2c17.2 99.2 90.4 158 200.2 158c20.9 0 39.4-1.7 55.3-5.1c3.7-.8 6.4-4 6.4-7.8v-42.8c0-5-4.6-8.8-9.5-7.8c-14.7 2.8-31.9 4.1-51.8 4.1c-68.5 0-114.5-36.6-129.8-98.6h130.6c4.4 0 8-3.6 8-8v-27.2c0-4.4-3.6-8-8-8H439.2v-36c0-4.7 0-9.4.3-13.8h135.9c4.4 0 8-3.6 8-8v-27.2c0-4.4-3.6-8-8-8H447.1c17.2-56.9 62.3-90.4 127.6-90.4c19.9 0 37.1 1.5 51.7 4.4a8 8 0 0 0 9.6-7.8v-42.8c0-3.8-2.6-7-6.3-7.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EuroTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372m117.1 581.1c0 3.8-2.7 7-6.4 7.8c-15.9 3.4-34.4 5.1-55.3 5.1c-109.8 0-183-58.8-200.2-158H337c-4.4 0-8-3.6-8-8v-27.2c0-4.4 3.6-8 8-8h26.1v-36.9c0-4.4 0-8.7.3-12.8H337c-4.4 0-8-3.6-8-8v-27.2c0-4.4 3.6-8 8-8h31.8C388.5 345.7 460.7 290 567.4 290c20.9 0 39.4 1.9 55.3 5.4c3.7.8 6.3 4 6.3 7.8V346a8 8 0 0 1-9.6 7.8c-14.6-2.9-31.8-4.4-51.7-4.4c-65.3 0-110.4 33.5-127.6 90.4h128.3c4.4 0 8 3.6 8 8V475c0 4.4-3.6 8-8 8H432.5c-.3 4.4-.3 9.1-.3 13.8v36h136.4c4.4 0 8 3.6 8 8V568c0 4.4-3.6 8-8 8H438c15.3 62 61.3 98.6 129.8 98.6c19.9 0 37.1-1.3 51.8-4.1c4.9-1 9.5 2.8 9.5 7.8z"/><path fill="currentColor" d="M619.6 670.5c-14.7 2.8-31.9 4.1-51.8 4.1c-68.5 0-114.5-36.6-129.8-98.6h130.6c4.4 0 8-3.6 8-8v-27.2c0-4.4-3.6-8-8-8H432.2v-36c0-4.7 0-9.4.3-13.8h135.9c4.4 0 8-3.6 8-8v-27.2c0-4.4-3.6-8-8-8H440.1c17.2-56.9 62.3-90.4 127.6-90.4c19.9 0 37.1 1.5 51.7 4.4a8 8 0 0 0 9.6-7.8v-42.8c0-3.8-2.6-7-6.3-7.8c-15.9-3.5-34.4-5.4-55.3-5.4c-106.7 0-178.9 55.7-198.6 149.9H337c-4.4 0-8 3.6-8 8v27.2c0 4.4 3.6 8 8 8h26.4c-.3 4.1-.3 8.4-.3 12.8v36.9H337c-4.4 0-8 3.6-8 8V568c0 4.4 3.6 8 8 8h30.2c17.2 99.2 90.4 158 200.2 158c20.9 0 39.4-1.7 55.3-5.1c3.7-.8 6.4-4 6.4-7.8v-42.8c0-5-4.6-8.8-9.5-7.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExceptionOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M688 312v-48c0-4.4-3.6-8-8-8H296c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h384c4.4 0 8-3.6 8-8m-392 88c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h184c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8zm376 116c-119.3 0-216 96.7-216 216s96.7 216 216 216s216-96.7 216-216s-96.7-216-216-216m107.5 323.5C750.8 868.2 712.6 884 672 884s-78.8-15.8-107.5-44.5C535.8 810.8 520 772.6 520 732s15.8-78.8 44.5-107.5C593.2 595.8 631.4 580 672 580s78.8 15.8 107.5 44.5C808.2 653.2 824 691.4 824 732s-15.8 78.8-44.5 107.5M640 812a32 32 0 1 0 64 0a32 32 0 1 0-64 0m12-64h40c4.4 0 8-3.6 8-8V628c0-4.4-3.6-8-8-8h-40c-4.4 0-8 3.6-8 8v112c0 4.4 3.6 8 8 8M440 852H208V148h560v344c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V108c0-17.7-14.3-32-32-32H168c-17.7 0-32 14.3-32 32v784c0 17.7 14.3 32 32 32h272c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m-32 232c0-4.4 3.6-8 8-8h48c4.4 0 8 3.6 8 8v272c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8zm32 440a48.01 48.01 0 0 1 0-96a48.01 48.01 0 0 1 0 96"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationCircleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" d="M464 688a48 48 0 1 0 96 0a48 48 0 1 0-96 0m24-112h48c4.4 0 8-3.6 8-8V296c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v272c0 4.4 3.6 8 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationCircleTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372m-32 156c0-4.4 3.6-8 8-8h48c4.4 0 8 3.6 8 8v272c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8zm32 440a48.01 48.01 0 0 1 0-96a48.01 48.01 0 0 1 0 96"/><path fill="currentColor" d="M488 576h48c4.4 0 8-3.6 8-8V296c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v272c0 4.4 3.6 8 8 8m-24 112a48 48 0 1 0 96 0a48 48 0 1 0-96 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 804a64 64 0 1 0 128 0a64 64 0 1 0-128 0m32-168h64c4.4 0 8-3.6 8-8V164c0-4.4-3.6-8-8-8h-64c-4.4 0-8 3.6-8 8v464c0 4.4 3.6 8 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExpandAltOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m855 160.1l-189.2 23.5c-6.6.8-9.3 8.8-4.7 13.5l54.7 54.7l-153.5 153.5a8.03 8.03 0 0 0 0 11.3l45.1 45.1c3.1 3.1 8.2 3.1 11.3 0l153.6-153.6l54.7 54.7a7.94 7.94 0 0 0 13.5-4.7L863.9 169a7.9 7.9 0 0 0-8.9-8.9M416.6 562.3a8.03 8.03 0 0 0-11.3 0L251.8 715.9l-54.7-54.7a7.94 7.94 0 0 0-13.5 4.7L160.1 855c-.6 5.2 3.7 9.5 8.9 8.9l189.2-23.5c6.6-.8 9.3-8.8 4.7-13.5l-54.7-54.7l153.6-153.6c3.1-3.1 3.1-8.2 0-11.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExpandOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M342 88H120c-17.7 0-32 14.3-32 32v224c0 8.8 7.2 16 16 16h48c8.8 0 16-7.2 16-16V168h174c8.8 0 16-7.2 16-16v-48c0-8.8-7.2-16-16-16m578 576h-48c-8.8 0-16 7.2-16 16v176H682c-8.8 0-16 7.2-16 16v48c0 8.8 7.2 16 16 16h222c17.7 0 32-14.3 32-32V680c0-8.8-7.2-16-16-16M342 856H168V680c0-8.8-7.2-16-16-16h-48c-8.8 0-16 7.2-16 16v224c0 17.7 14.3 32 32 32h222c8.8 0 16-7.2 16-16v-48c0-8.8-7.2-16-16-16M904 88H682c-8.8 0-16 7.2-16 16v48c0 8.8 7.2 16 16 16h174v176c0 8.8 7.2 16 16 16h48c8.8 0 16-7.2 16-16V120c0-17.7-14.3-32-32-32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExperimentFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m218.9 636.3l42.6 26.6c.1.1.3.2.4.3l12.7 8l.3.3a186.9 186.9 0 0 0 94.1 25.1c44.9 0 87.2-15.7 121-43.8a256.27 256.27 0 0 1 164.9-59.9c52.3 0 102.2 15.7 144.6 44.5l7.9 5l-111.6-289V179.8h63.5c4.4 0 8-3.6 8-8V120c0-4.4-3.6-8-8-8H264.7c-4.4 0-8 3.6-8 8v51.9c0 4.4 3.6 8 8 8h63.5v173.6zm333-203.1c22 0 39.9 17.9 39.9 39.9S573.9 513 551.9 513S512 495.1 512 473.1s17.9-39.9 39.9-39.9M878 825.1l-29.9-77.4l-85.7-53.5l-.1.1c-.7-.5-1.5-1-2.2-1.5l-8.1-5l-.3-.3c-29-17.5-62.3-26.8-97-26.8c-44.9 0-87.2 15.7-121 43.8a256.27 256.27 0 0 1-164.9 59.9c-53 0-103.5-16.1-146.2-45.6l-28.9-18.1L146 825.1c-2.8 7.4-4.3 15.2-4.3 23c0 35.2 28.6 63.8 63.8 63.8h612.9c7.9 0 15.7-1.5 23-4.3a63.6 63.6 0 0 0 36.6-82.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExperimentOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 472a40 40 0 1 0 80 0a40 40 0 1 0-80 0m367 352.9L696.3 352V178H768v-68H256v68h71.7v174L145 824.9c-2.8 7.4-4.3 15.2-4.3 23.1c0 35.3 28.7 64 64 64h614.6c7.9 0 15.7-1.5 23.1-4.3c33-12.7 49.4-49.8 36.6-82.8M395.7 364.7V180h232.6v184.7L719.2 600c-20.7-5.3-42.1-8-63.9-8c-61.2 0-119.2 21.5-165.3 60a188.78 188.78 0 0 1-121.3 43.9c-32.7 0-64.1-8.3-91.8-23.7zM210.5 844l41.7-107.8c35.7 18.1 75.4 27.8 116.6 27.8c61.2 0 119.2-21.5 165.3-60c33.9-28.2 76.3-43.9 121.3-43.9c35 0 68.4 9.5 97.6 27.1L813.5 844z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExperimentTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M551.9 513c19.6 0 35.9-14.2 39.3-32.8A40.02 40.02 0 0 1 552 512a40 40 0 0 1-40-39.4v.5c0 22 17.9 39.9 39.9 39.9M752 687.8l-.3-.3c-29-17.5-62.3-26.8-97-26.8c-44.9 0-87.2 15.7-121 43.8a256.27 256.27 0 0 1-164.9 59.9c-41.2 0-81-9.8-116.7-28L210.5 844h603l-59.9-155.2z"/><path fill="currentColor" d="M879 824.9L696.3 352V178H768v-68H256v68h71.7v174L145 824.9c-2.8 7.4-4.3 15.2-4.3 23.1c0 35.3 28.7 64 64 64h614.6c7.9 0 15.7-1.5 23.1-4.3c33-12.7 49.4-49.8 36.6-82.8M395.7 364.7V180h232.6v184.7L719.2 600c-20.7-5.3-42.1-8-63.9-8c-61.2 0-119.2 21.5-165.3 60a188.78 188.78 0 0 1-121.3 43.9c-32.7 0-64.1-8.3-91.8-23.7zM210.5 844l41.6-107.6l.1-.2c35.7 18.1 75.4 27.8 116.6 27.8c61.2 0 119.2-21.5 165.3-60c33.9-28.2 76.3-43.9 121.3-43.9c35 0 68.4 9.5 97.6 27.1l.6 1.6L813.5 844z"/><path fill="currentColor" d="M552 512c19.3 0 35.4-13.6 39.2-31.8c.6-2.7.8-5.4.8-8.2c0-22.1-17.9-40-40-40s-40 17.9-40 40v.6a40 40 0 0 0 40 39.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExportOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M880 912H144c-17.7 0-32-14.3-32-32V144c0-17.7 14.3-32 32-32h360c4.4 0 8 3.6 8 8v56c0 4.4-3.6 8-8 8H184v656h656V520c0-4.4 3.6-8 8-8h56c4.4 0 8 3.6 8 8v360c0 17.7-14.3 32-32 32M770.87 199.131l-52.2-52.2c-4.7-4.7-1.9-12.8 4.7-13.6l179.4-21c5.1-.6 9.5 3.7 8.9 8.9l-21 179.4c-.8 6.6-8.9 9.4-13.6 4.7l-52.4-52.4l-256.2 256.2c-3.1 3.1-8.2 3.1-11.3 0l-42.4-42.4c-3.1-3.1-3.1-8.2 0-11.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M396 512a112 112 0 1 0 224 0a112 112 0 1 0-224 0m546.2-25.8C847.4 286.5 704.1 186 512 186c-192.2 0-335.4 100.5-430.2 300.3a60.3 60.3 0 0 0 0 51.5C176.6 737.5 319.9 838 512 838c192.2 0 335.4-100.5 430.2-300.3c7.7-16.2 7.7-35 0-51.5M508 688c-97.2 0-176-78.8-176-176s78.8-176 176-176s176 78.8 176 176s-78.8 176-176 176"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeInvisibleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M508 624a112 112 0 0 0 112-112c0-3.28-.15-6.53-.43-9.74L498.26 623.57c3.21.28 6.45.43 9.74.43m370.72-458.44L836 122.88a8 8 0 0 0-11.31 0L715.37 232.23Q624.91 186 512 186q-288.3 0-430.2 300.3a60.3 60.3 0 0 0 0 51.5q56.7 119.43 136.55 191.45L112.56 835a8 8 0 0 0 0 11.31L155.25 889a8 8 0 0 0 11.31 0l712.16-712.12a8 8 0 0 0 0-11.32M332 512a176 176 0 0 1 258.88-155.28l-48.62 48.62a112.08 112.08 0 0 0-140.92 140.92l-48.62 48.62A175.09 175.09 0 0 1 332 512"/><path fill="currentColor" d="M942.2 486.2Q889.4 375 816.51 304.85L672.37 449A176.08 176.08 0 0 1 445 676.37L322.74 798.63Q407.82 838 512 838q288.3 0 430.2-300.3a60.29 60.29 0 0 0 0-51.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeInvisibleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M942.2 486.2Q889.47 375.11 816.7 305l-50.88 50.88C807.31 395.53 843.45 447.4 874.7 512C791.5 684.2 673.4 766 512 766q-72.67 0-133.87-22.38L323 798.75Q408 838 512 838q288.3 0 430.2-300.3a60.29 60.29 0 0 0 0-51.5m-63.57-320.64L836 122.88a8 8 0 0 0-11.32 0L715.31 232.2Q624.86 186 512 186q-288.3 0-430.2 300.3a60.3 60.3 0 0 0 0 51.5q56.69 119.4 136.5 191.41L112.48 835a8 8 0 0 0 0 11.31L155.17 889a8 8 0 0 0 11.31 0l712.15-712.12a8 8 0 0 0 0-11.32M149.3 512C232.6 339.8 350.7 258 512 258c54.54 0 104.13 9.36 149.12 28.39l-70.3 70.3a176 176 0 0 0-238.13 238.13l-83.42 83.42C223.1 637.49 183.3 582.28 149.3 512m246.7 0a112.11 112.11 0 0 1 146.2-106.69L401.31 546.2A112 112 0 0 1 396 512"/><path fill="currentColor" d="M508 624c-3.46 0-6.87-.16-10.25-.47l-52.82 52.82a176.09 176.09 0 0 0 227.42-227.42l-52.82 52.82c.31 3.38.47 6.79.47 10.25a111.94 111.94 0 0 1-112 112"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeInvisibleTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="m254.89 758.85l125.57-125.57a176 176 0 0 1 248.82-248.82L757 256.72Q651.69 186.07 512 186q-288.3 0-430.2 300.3a60.3 60.3 0 0 0 0 51.5q69.27 145.91 173.09 221.05M942.2 486.2Q889.46 375.11 816.7 305L672.48 449.27a176.09 176.09 0 0 1-227.22 227.21L323 798.75Q408 838 512 838q288.3 0 430.2-300.3a60.29 60.29 0 0 0 0-51.5"/><path fill="currentColor" d="M942.2 486.2Q889.47 375.11 816.7 305l-50.88 50.88C807.31 395.53 843.45 447.4 874.7 512C791.5 684.2 673.4 766 512 766q-72.67 0-133.87-22.38L323 798.75Q408 838 512 838q288.3 0 430.2-300.3a60.29 60.29 0 0 0 0-51.5m-63.57-320.64L836 122.88a8 8 0 0 0-11.32 0L715.31 232.2Q624.86 186 512 186q-288.3 0-430.2 300.3a60.3 60.3 0 0 0 0 51.5q56.69 119.4 136.5 191.41L112.48 835a8 8 0 0 0 0 11.31L155.17 889a8 8 0 0 0 11.31 0l712.15-712.12a8 8 0 0 0 0-11.32M149.3 512C232.6 339.8 350.7 258 512 258c54.54 0 104.13 9.36 149.12 28.39l-70.3 70.3a176 176 0 0 0-238.13 238.13l-83.42 83.42C223.1 637.49 183.3 582.28 149.3 512m246.7 0a112.11 112.11 0 0 1 146.2-106.69L401.31 546.2A112 112 0 0 1 396 512"/><path fill="currentColor" d="M508 624c-3.46 0-6.87-.16-10.25-.47l-52.82 52.82a176.09 176.09 0 0 0 227.42-227.42l-52.82 52.82c.31 3.38.47 6.79.47 10.25a111.94 111.94 0 0 1-112 112"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M942.2 486.2C847.4 286.5 704.1 186 512 186c-192.2 0-335.4 100.5-430.2 300.3a60.3 60.3 0 0 0 0 51.5C176.6 737.5 319.9 838 512 838c192.2 0 335.4-100.5 430.2-300.3c7.7-16.2 7.7-35 0-51.5M512 766c-161.3 0-279.4-81.8-362.7-254C232.6 339.8 350.7 258 512 258c161.3 0 279.4 81.8 362.7 254C791.5 684.2 673.4 766 512 766m-4-430c-97.2 0-176 78.8-176 176s78.8 176 176 176s176-78.8 176-176s-78.8-176-176-176m0 288c-61.9 0-112-50.1-112-112s50.1-112 112-112s112 50.1 112 112s-50.1 112-112 112"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M81.8 537.8a60.3 60.3 0 0 1 0-51.5C176.6 286.5 319.8 186 512 186c-192.2 0-335.4 100.5-430.2 300.3a60.3 60.3 0 0 0 0 51.5C176.6 737.5 319.9 838 512 838c-192.1 0-335.4-100.5-430.2-300.2"/><path fill="currentColor" fill-opacity=".15" d="M512 258c-161.3 0-279.4 81.8-362.7 254C232.6 684.2 350.7 766 512 766c161.4 0 279.5-81.8 362.7-254C791.4 339.8 673.3 258 512 258m-4 430c-97.2 0-176-78.8-176-176s78.8-176 176-176s176 78.8 176 176s-78.8 176-176 176"/><path fill="currentColor" d="M942.2 486.2C847.4 286.5 704.1 186 512 186c-192.2 0-335.4 100.5-430.2 300.3a60.3 60.3 0 0 0 0 51.5C176.6 737.5 319.9 838 512 838c192.2 0 335.4-100.5 430.2-300.3c7.7-16.2 7.7-35 0-51.5M512 766c-161.3 0-279.4-81.8-362.7-254C232.6 339.8 350.7 258 512 258s279.4 81.8 362.7 254C791.5 684.2 673.4 766 512 766"/><path fill="currentColor" d="M508 336c-97.2 0-176 78.8-176 176s78.8 176 176 176s176-78.8 176-176s-78.8-176-176-176m0 288c-61.9 0-112-50.1-112-112s50.1-112 112-112s112 50.1 112 112s-50.1 112-112 112"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacebookFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-92.4 233.5h-63.9c-50.1 0-59.8 23.8-59.8 58.8v77.1h119.6l-15.6 120.7h-104V912H539.2V602.2H434.9V481.4h104.3v-89c0-103.3 63.1-159.6 155.3-159.6c44.2 0 82.1 3.3 93.2 4.8v107.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacebookOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-32 736H663.9V602.2h104l15.6-120.7H663.9v-77.1c0-35 9.7-58.8 59.8-58.8h63.9v-108c-11.1-1.5-49-4.8-93.2-4.8c-92.2 0-155.3 56.3-155.3 159.6v89H434.9v120.7h104.3V848H176V176h672z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FallOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m925.9 804l-24-199.2c-.8-6.6-8.9-9.4-13.6-4.7L829 659.5L557.7 388.3c-6.3-6.2-16.4-6.2-22.6 0L433.3 490L156.6 213.3a8.03 8.03 0 0 0-11.3 0l-45 45.2a8.03 8.03 0 0 0 0 11.3L422 591.7c6.2 6.3 16.4 6.3 22.6 0L546.4 490l226.1 226l-59.3 59.3a8.01 8.01 0 0 0 4.7 13.6l199.2 24c5.1.7 9.5-3.7 8.8-8.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastBackwardFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M517.6 273.5L230.2 499.3a16.14 16.14 0 0 0 0 25.4l287.4 225.8c10.7 8.4 26.4.8 26.4-12.7V286.2c0-13.5-15.7-21.1-26.4-12.7m320 0L550.2 499.3a16.14 16.14 0 0 0 0 25.4l287.4 225.8c10.7 8.4 26.4.8 26.4-12.7V286.2c0-13.5-15.7-21.1-26.4-12.7m-620-25.5h-51.2c-3.5 0-6.4 2.7-6.4 6v516c0 3.3 2.9 6 6.4 6h51.2c3.5 0 6.4-2.7 6.4-6V254c0-3.3-2.9-6-6.4-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastBackwardOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M517.6 273.5L230.2 499.3a16.14 16.14 0 0 0 0 25.4l287.4 225.8c10.7 8.4 26.4.8 26.4-12.7V286.2c0-13.5-15.7-21.1-26.4-12.7m320 0L550.2 499.3a16.14 16.14 0 0 0 0 25.4l287.4 225.8c10.7 8.4 26.4.8 26.4-12.7V286.2c0-13.5-15.7-21.1-26.4-12.7m-620-25.5h-51.2c-3.5 0-6.4 2.7-6.4 6v516c0 3.3 2.9 6 6.4 6h51.2c3.5 0 6.4-2.7 6.4-6V254c0-3.3-2.9-6-6.4-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastForwardFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M793.8 499.3L506.4 273.5c-10.7-8.4-26.4-.8-26.4 12.7v451.6c0 13.5 15.7 21.1 26.4 12.7l287.4-225.8a16.14 16.14 0 0 0 0-25.4m-320 0L186.4 273.5c-10.7-8.4-26.4-.8-26.4 12.7v451.5c0 13.5 15.7 21.1 26.4 12.7l287.4-225.8c4.1-3.2 6.2-8 6.2-12.7c0-4.6-2.1-9.4-6.2-12.6M857.6 248h-51.2c-3.5 0-6.4 2.7-6.4 6v516c0 3.3 2.9 6 6.4 6h51.2c3.5 0 6.4-2.7 6.4-6V254c0-3.3-2.9-6-6.4-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastForwardOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M793.8 499.3L506.4 273.5c-10.7-8.4-26.4-.8-26.4 12.7v451.6c0 13.5 15.7 21.1 26.4 12.7l287.4-225.8a16.14 16.14 0 0 0 0-25.4m-320 0L186.4 273.5c-10.7-8.4-26.4-.8-26.4 12.7v451.5c0 13.5 15.7 21.1 26.4 12.7l287.4-225.8c4.1-3.2 6.2-8 6.2-12.7c0-4.6-2.1-9.4-6.2-12.6M857.6 248h-51.2c-3.5 0-6.4 2.7-6.4 6v516c0 3.3 2.9 6 6.4 6h51.2c3.5 0 6.4-2.7 6.4-6V254c0-3.3-2.9-6-6.4-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FieldBinaryOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 395.4h91V649h79V267c0-4.4-3.6-8-8-8h-48.2c-3.7 0-7 2.6-7.7 6.3c-2.6 12.1-6.9 22.3-12.9 30.9c-7.2 10.1-15.9 18.2-26.3 24.4c-10.3 6.2-22 10.5-35 12.9c-10.4 1.9-21 3-32 3.1c-4.4.1-7.9 3.6-7.9 8v42.8c0 4.4 3.6 8 8 8M871 702H567c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h304c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8M443.9 312.7c-16.1-19-34.4-32.4-55.2-40.4c-21.3-8.2-44.1-12.3-68.4-12.3c-23.9 0-46.4 4.1-67.7 12.3c-20.8 8-39 21.4-54.8 40.3c-15.9 19.1-28.7 44.7-38.3 77c-9.6 32.5-14.5 73-14.5 121.5c0 49.9 4.9 91.4 14.5 124.4c9.6 32.8 22.4 58.7 38.3 77.7c15.8 18.9 34 32.3 54.8 40.3c21.3 8.2 43.8 12.3 67.7 12.3c24.4 0 47.2-4.1 68.4-12.3c20.8-8 39.2-21.4 55.2-40.4c16.1-19 29-44.9 38.6-77.7c9.6-33 14.5-74.5 14.5-124.4c0-48.4-4.9-88.9-14.5-121.5c-9.5-32.1-22.4-57.7-38.6-76.8m-29.5 251.7c-1 21.4-4.2 42-9.5 61.9c-5.5 20.7-14.5 38.5-27 53.4c-13.6 16.3-33.2 24.3-57.6 24.3c-24 0-43.2-8.1-56.7-24.4c-12.2-14.8-21.1-32.6-26.6-53.3c-5.3-19.9-8.5-40.6-9.5-61.9c-1-20.8-1.5-38.5-1.5-53.2c0-8.8.1-19.4.4-31.8c.2-12.7 1.1-25.8 2.6-39.2c1.5-13.6 4-27.1 7.6-40.5c3.7-13.8 8.8-26.3 15.4-37.4c6.9-11.6 15.8-21.1 26.7-28.3c11.4-7.6 25.3-11.3 41.5-11.3c16.1 0 30.1 3.7 41.7 11.2c11.1 7.2 20.3 16.6 27.4 28.2c6.9 11.2 12.1 23.8 15.6 37.7c3.3 13.2 5.8 26.6 7.5 40.1c1.8 13.5 2.8 26.6 3 39.4c.2 12.4.4 23 .4 31.8c.1 14.8-.4 32.5-1.4 53.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FieldNumberOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M508 280h-63.3c-3.3 0-6 2.7-6 6v340.2H433L197.4 282.6c-1.1-1.6-3-2.6-4.9-2.6H126c-3.3 0-6 2.7-6 6v464c0 3.3 2.7 6 6 6h62.7c3.3 0 6-2.7 6-6V405.1h5.7l238.2 348.3c1.1 1.6 3 2.6 5 2.6H508c3.3 0 6-2.7 6-6V286c0-3.3-2.7-6-6-6m378 413H582c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h304c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8m-152.2-63c52.9 0 95.2-17.2 126.2-51.7c29.4-32.9 44-75.8 44-128.8c0-53.1-14.6-96.5-44-129.3c-30.9-34.8-73.2-52.2-126.2-52.2c-53.7 0-95.9 17.5-126.3 52.8c-29.2 33.1-43.4 75.9-43.4 128.7c0 52.4 14.3 95.2 43.5 128.3c30.6 34.7 73 52.2 126.2 52.2m-71.5-263.7c16.9-20.6 40.3-30.9 71.4-30.9c31.5 0 54.8 9.6 71 29.1c16.4 20.3 24.9 48.6 24.9 84.9c0 36.3-8.4 64.1-24.8 83.9c-16.5 19.4-40 29.2-71.1 29.2c-31.2 0-55-10.3-71.4-30.4c-16.3-20.1-24.5-47.3-24.5-82.6c.1-35.8 8.2-63 24.5-83.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FieldStringOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M875.6 515.9c2.1.8 4.4-.3 5.2-2.4c.2-.4.2-.9.2-1.4v-58.3c0-1.8-1.1-3.3-2.8-3.8c-6-1.8-17.2-3-27.2-3c-32.9 0-61.7 16.7-73.5 41.2v-28.6c0-4.4-3.6-8-8-8H717c-4.4 0-8 3.6-8 8V729c0 4.4 3.6 8 8 8h54.8c4.4 0 8-3.6 8-8V572.7c0-36.2 26.1-60.2 65.1-60.2c10.4.1 26.6 1.8 30.7 3.4m-537-40.5l-54.7-12.6c-61.2-14.2-87.7-34.8-87.7-70.7c0-44.6 39.1-73.5 96.9-73.5c52.8 0 91.4 26.5 99.9 68.9h70C455.9 311.6 387.6 259 293.4 259c-103.3 0-171 55.5-171 139c0 68.6 38.6 109.5 122.2 128.5l61.6 14.3c63.6 14.9 91.6 37.1 91.6 75.1c0 44.1-43.5 75.2-102.5 75.2c-60.6 0-104.5-27.2-112.8-70.5H111c7.2 79.9 75.6 130.4 179.1 130.4C402.3 751 471 695.2 471 605.3c0-70.2-38.6-108.5-132.4-129.9M841 729a36 36 0 1 0 72 0a36 36 0 1 0-72 0M653 457.8h-51.4V396c0-4.4-3.6-8-8-8h-54.7c-4.4 0-8 3.6-8 8v61.8H495c-4.4 0-8 3.6-8 8v42.3c0 4.4 3.6 8 8 8h35.9v147.5c0 56.2 27.4 79.4 93.1 79.4c11.7 0 23.6-1.2 33.8-3.1c1.9-.3 3.2-2 3.2-3.9v-49.3c0-2.2-1.8-4-4-4h-.4c-4.9.5-6.2.6-8.3.8c-4.1.3-7.8.5-12.6.5c-24.1 0-34.1-10.3-34.1-35.6V516.1H653c4.4 0 8-3.6 8-8v-42.3c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FieldTimeOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M945 412H689c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h256c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8M811 548H689c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h122c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8M477.3 322.5H434c-6.2 0-11.2 5-11.2 11.2v248c0 3.6 1.7 6.9 4.6 9l148.9 108.6c5 3.6 12 2.6 15.6-2.4l25.7-35.1v-.1c3.6-5 2.5-12-2.5-15.6l-126.7-91.6V333.7c.1-6.2-5-11.2-11.1-11.2"/><path fill="currentColor" d="M804.8 673.9H747c-5.6 0-10.9 2.9-13.9 7.7c-12.7 20.1-27.5 38.7-44.5 55.7c-29.3 29.3-63.4 52.3-101.3 68.3c-39.3 16.6-81 25-124 25c-43.1 0-84.8-8.4-124-25c-37.9-16-72-39-101.3-68.3s-52.3-63.4-68.3-101.3c-16.6-39.2-25-80.9-25-124c0-43.1 8.4-84.7 25-124c16-37.9 39-72 68.3-101.3c29.3-29.3 63.4-52.3 101.3-68.3c39.2-16.6 81-25 124-25c43.1 0 84.8 8.4 124 25c37.9 16 72 39 101.3 68.3c17 17 31.8 35.6 44.5 55.7c3 4.8 8.3 7.7 13.9 7.7h57.8c6.9 0 11.3-7.2 8.2-13.3c-65.2-129.7-197.4-214-345-215.7c-216.1-2.7-395.6 174.2-396 390.1C71.6 727.5 246.9 903 463.2 903c149.5 0 283.9-84.6 349.8-215.8c3.1-6.1-1.4-13.3-8.2-13.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileAddFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M480 580H372a8 8 0 0 0-8 8v48a8 8 0 0 0 8 8h108v108a8 8 0 0 0 8 8h48a8 8 0 0 0 8-8V644h108a8 8 0 0 0 8-8v-48a8 8 0 0 0-8-8H544V472a8 8 0 0 0-8-8h-48a8 8 0 0 0-8 8zm374.6-291.3c6 6 9.4 14.1 9.4 22.6V928c0 17.7-14.3 32-32 32H192c-17.7 0-32-14.3-32-32V96c0-17.7 14.3-32 32-32h424.7c8.5 0 16.7 3.4 22.7 9.4zM790.2 326L602 137.8V326z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileAddOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M854.6 288.6L639.4 73.4c-6-6-14.1-9.4-22.6-9.4H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V311.3c0-8.5-3.4-16.7-9.4-22.7M790.2 326H602V137.8zm1.8 562H232V136h302v216a42 42 0 0 0 42 42h216zM544 472c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v108H372c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h108v108c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V644h108c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8H544z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileAddTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M534 352V136H232v752h560V394H576a42 42 0 0 1-42-42m126 236v48c0 4.4-3.6 8-8 8H544v108c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8V644H372c-4.4 0-8-3.6-8-8v-48c0-4.4 3.6-8 8-8h108V472c0-4.4 3.6-8 8-8h48c4.4 0 8 3.6 8 8v108h108c4.4 0 8 3.6 8 8"/><path fill="currentColor" d="M854.6 288.6L639.4 73.4c-6-6-14.1-9.4-22.6-9.4H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V311.3c0-8.5-3.4-16.7-9.4-22.7M602 137.8L790.2 326H602zM792 888H232V136h302v216a42 42 0 0 0 42 42h216z"/><path fill="currentColor" d="M544 472c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v108H372c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h108v108c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V644h108c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8H544z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileDoneOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M688 312v-48c0-4.4-3.6-8-8-8H296c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h384c4.4 0 8-3.6 8-8m-392 88c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h184c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8zm376 116c-119.3 0-216 96.7-216 216s96.7 216 216 216s216-96.7 216-216s-96.7-216-216-216m107.5 323.5C750.8 868.2 712.6 884 672 884s-78.8-15.8-107.5-44.5C535.8 810.8 520 772.6 520 732s15.8-78.8 44.5-107.5C593.2 595.8 631.4 580 672 580s78.8 15.8 107.5 44.5C808.2 653.2 824 691.4 824 732s-15.8 78.8-44.5 107.5M761 656h-44.3c-2.6 0-5 1.2-6.5 3.3l-63.5 87.8l-23.1-31.9a7.92 7.92 0 0 0-6.5-3.3H573c-6.5 0-10.3 7.4-6.5 12.7l73.8 102.1c3.2 4.4 9.7 4.4 12.9 0l114.2-158c3.9-5.3.1-12.7-6.4-12.7M440 852H208V148h560v344c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V108c0-17.7-14.3-32-32-32H168c-17.7 0-32 14.3-32 32v784c0 17.7 14.3 32 32 32h272c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileExcelFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M854.6 288.7c6 6 9.4 14.1 9.4 22.6V928c0 17.7-14.3 32-32 32H192c-17.7 0-32-14.3-32-32V96c0-17.7 14.3-32 32-32h424.7c8.5 0 16.7 3.4 22.7 9.4zM790.2 326L602 137.8V326zM575.34 477.84l-61.22 102.3L452.3 477.8a12 12 0 0 0-10.27-5.79h-38.44a12 12 0 0 0-6.4 1.85a12 12 0 0 0-3.75 16.56l82.34 130.42l-83.45 132.78a12 12 0 0 0-1.84 6.39a12 12 0 0 0 12 12h34.46a12 12 0 0 0 10.21-5.7l62.7-101.47l62.3 101.45a12 12 0 0 0 10.23 5.72h37.48a12 12 0 0 0 6.48-1.9a12 12 0 0 0 3.62-16.58l-83.83-130.55l85.3-132.47a12 12 0 0 0 1.9-6.5a12 12 0 0 0-12-12h-35.7a12 12 0 0 0-10.29 5.84z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileExcelOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M854.6 288.6L639.4 73.4c-6-6-14.1-9.4-22.6-9.4H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V311.3c0-8.5-3.4-16.7-9.4-22.7M790.2 326H602V137.8zm1.8 562H232V136h302v216a42 42 0 0 0 42 42h216zM514.1 580.1l-61.8-102.4c-2.2-3.6-6.1-5.8-10.3-5.8h-38.4c-2.3 0-4.5.6-6.4 1.9c-5.6 3.5-7.3 10.9-3.7 16.6l82.3 130.4l-83.4 132.8a12.04 12.04 0 0 0 10.2 18.4h34.5c4.2 0 8-2.2 10.2-5.7L510 664.8l62.3 101.4c2.2 3.6 6.1 5.7 10.2 5.7H620c2.3 0 4.5-.7 6.5-1.9c5.6-3.6 7.2-11 3.6-16.6l-84-130.4l85.3-132.5a12.04 12.04 0 0 0-10.1-18.5h-35.7c-4.2 0-8.1 2.2-10.3 5.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileExcelTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M534 352V136H232v752h560V394H576a42 42 0 0 1-42-42m51.6 120h35.7a12.04 12.04 0 0 1 10.1 18.5L546.1 623l84 130.4c3.6 5.6 2 13-3.6 16.6c-2 1.2-4.2 1.9-6.5 1.9h-37.5c-4.1 0-8-2.1-10.2-5.7L510 664.8l-62.7 101.5c-2.2 3.5-6 5.7-10.2 5.7h-34.5a12.04 12.04 0 0 1-10.2-18.4l83.4-132.8l-82.3-130.4c-3.6-5.7-1.9-13.1 3.7-16.6c1.9-1.3 4.1-1.9 6.4-1.9H442c4.2 0 8.1 2.2 10.3 5.8l61.8 102.4l61.2-102.3c2.2-3.6 6.1-5.8 10.3-5.8"/><path fill="currentColor" d="M854.6 288.6L639.4 73.4c-6-6-14.1-9.4-22.6-9.4H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V311.3c0-8.5-3.4-16.7-9.4-22.7M602 137.8L790.2 326H602zM792 888H232V136h302v216a42 42 0 0 0 42 42h216z"/><path fill="currentColor" d="m514.1 580.1l-61.8-102.4c-2.2-3.6-6.1-5.8-10.3-5.8h-38.4c-2.3 0-4.5.6-6.4 1.9c-5.6 3.5-7.3 10.9-3.7 16.6l82.3 130.4l-83.4 132.8a12.04 12.04 0 0 0 10.2 18.4h34.5c4.2 0 8-2.2 10.2-5.7L510 664.8l62.3 101.4c2.2 3.6 6.1 5.7 10.2 5.7H620c2.3 0 4.5-.7 6.5-1.9c5.6-3.6 7.2-11 3.6-16.6l-84-130.4l85.3-132.5a12.04 12.04 0 0 0-10.1-18.5h-35.7c-4.2 0-8.1 2.2-10.3 5.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileExclamationFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M854.6 288.7c6 6 9.4 14.1 9.4 22.6V928c0 17.7-14.3 32-32 32H192c-17.7 0-32-14.3-32-32V96c0-17.7 14.3-32 32-32h424.7c8.5 0 16.7 3.4 22.7 9.4zM790.2 326L602 137.8V326zM512 784a40 40 0 1 0 0-80a40 40 0 0 0 0 80m32-152V448a8 8 0 0 0-8-8h-48a8 8 0 0 0-8 8v184a8 8 0 0 0 8 8h48a8 8 0 0 0 8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileExclamationOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M854.6 288.6L639.4 73.4c-6-6-14.1-9.4-22.6-9.4H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V311.3c0-8.5-3.4-16.7-9.4-22.7M790.2 326H602V137.8zm1.8 562H232V136h302v216a42 42 0 0 0 42 42h216zM472 744a40 40 0 1 0 80 0a40 40 0 1 0-80 0m16-104h48c4.4 0 8-3.6 8-8V448c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v184c0 4.4 3.6 8 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileExclamationTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M534 352V136H232v752h560V394H576a42 42 0 0 1-42-42m-54 96c0-4.4 3.6-8 8-8h48c4.4 0 8 3.6 8 8v184c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8zm32 336c-22.1 0-40-17.9-40-40s17.9-40 40-40s40 17.9 40 40s-17.9 40-40 40"/><path fill="currentColor" d="M854.6 288.6L639.4 73.4c-6-6-14.1-9.4-22.6-9.4H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V311.3c0-8.5-3.4-16.7-9.4-22.7M602 137.8L790.2 326H602zM792 888H232V136h302v216a42 42 0 0 0 42 42h216z"/><path fill="currentColor" d="M488 640h48c4.4 0 8-3.6 8-8V448c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v184c0 4.4 3.6 8 8 8m-16 104a40 40 0 1 0 80 0a40 40 0 1 0-80 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M854.6 288.7c6 6 9.4 14.1 9.4 22.6V928c0 17.7-14.3 32-32 32H192c-17.7 0-32-14.3-32-32V96c0-17.7 14.3-32 32-32h424.7c8.5 0 16.7 3.4 22.7 9.4zM790.2 326L602 137.8V326z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileGifOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M551.5 490.5H521c-4.6 0-8.4 3.7-8.4 8.4V720c0 4.6 3.7 8.4 8.4 8.4h30.5c4.6 0 8.4-3.7 8.4-8.4V498.9c-.1-4.6-3.8-8.4-8.4-8.4M477.3 600h-88.1c-4.6 0-8.4 3.7-8.4 8.4v23.8c0 4.6 3.7 8.4 8.4 8.4h47.6v.7c-.6 29.9-23 49.8-56.5 49.8c-39.2 0-63.6-30.7-63.6-81.4c0-50.1 23.9-80.6 62.3-80.6c28.1 0 47.5 13.5 55.4 38.3l.9 2.8h49.2l-.7-4.6C475.9 515.9 434.7 484 379 484c-68.8 0-113 49.4-113 125.9c0 77.5 43.7 126.1 113.6 126.1c64.4 0 106-40.3 106-102.9v-24.8c0-4.6-3.7-8.3-8.3-8.3"/><path fill="currentColor" d="M854.6 288.6L639.4 73.4c-6-6-14.1-9.4-22.6-9.4H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V311.3c0-8.5-3.4-16.7-9.4-22.7M602 137.8L790.2 326H602zM792 888H232V136h302v216c0 23.2 18.8 42 42 42h216z"/><path fill="currentColor" d="M608.2 727.8h32.3c4.6 0 8.4-3.7 8.4-8.4v-84.8h87.8c4.6 0 8.4-3.7 8.4-8.4v-25.5c0-4.6-3.7-8.4-8.4-8.4h-87.8v-58.9h96.8c4.6 0 8.4-3.7 8.4-8.4v-26.8c0-4.6-3.7-8.4-8.4-8.4H608.2c-4.6 0-8.4 3.7-8.4 8.4v221.1c0 4.8 3.8 8.5 8.4 8.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileImageFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M854.6 288.7L639.4 73.4c-6-6-14.2-9.4-22.7-9.4H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V311.3c0-8.5-3.4-16.6-9.4-22.6M400 402c22.1 0 40 17.9 40 40s-17.9 40-40 40s-40-17.9-40-40s17.9-40 40-40m296 294H328c-6.7 0-10.4-7.7-6.3-12.9l99.8-127.2a8 8 0 0 1 12.6 0l41.1 52.4l77.8-99.2a8 8 0 0 1 12.6 0l136.5 174c4.3 5.2.5 12.9-6.1 12.9m-94-370V137.8L790.2 326z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileImageOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m553.1 509.1l-77.8 99.2l-41.1-52.4a8 8 0 0 0-12.6 0l-99.8 127.2a7.98 7.98 0 0 0 6.3 12.9H696c6.7 0 10.4-7.7 6.3-12.9l-136.5-174a8.1 8.1 0 0 0-12.7 0M360 442a40 40 0 1 0 80 0a40 40 0 1 0-80 0m494.6-153.4L639.4 73.4c-6-6-14.1-9.4-22.6-9.4H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V311.3c0-8.5-3.4-16.7-9.4-22.7M790.2 326H602V137.8zm1.8 562H232V136h302v216a42 42 0 0 0 42 42h216z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileImageTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M534 352V136H232v752h560V394H576a42 42 0 0 1-42-42m-134 50c22.1 0 40 17.9 40 40s-17.9 40-40 40s-40-17.9-40-40s17.9-40 40-40m296 294H328.1c-6.7 0-10.4-7.7-6.3-12.9l99.8-127.2a8 8 0 0 1 12.6 0l41.1 52.4l77.8-99.2a8.1 8.1 0 0 1 12.7 0l136.5 174c4.1 5.2.4 12.9-6.3 12.9"/><path fill="currentColor" d="M854.6 288.6L639.4 73.4c-6-6-14.1-9.4-22.6-9.4H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V311.3c0-8.5-3.4-16.7-9.4-22.7M602 137.8L790.2 326H602zM792 888H232V136h302v216a42 42 0 0 0 42 42h216z"/><path fill="currentColor" d="m553.1 509.1l-77.8 99.2l-41.1-52.4a8 8 0 0 0-12.6 0l-99.8 127.2a7.98 7.98 0 0 0 6.3 12.9H696c6.7 0 10.4-7.7 6.3-12.9l-136.5-174a8.1 8.1 0 0 0-12.7 0M360 442a40 40 0 1 0 80 0a40 40 0 1 0-80 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileJpgOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M874.6 301.8L596.8 21.3c-4.5-4.5-9.4-8.3-14.7-11.5c-1.4-.8-2.8-1.6-4.3-2.3c-.9-.5-1.9-.9-2.8-1.3c-9-4-18.9-6.2-29-6.2H201c-39.8 0-73 32.2-73 72v880c0 39.8 33.2 72 73 72h623c39.8 0 71-32.2 71-72V352.5c0-19-7-37.2-20.4-50.7M583 110.4L783.8 312H583zM823 952H200V72h311v240c0 39.8 33.2 72 73 72h239zM350 696.5c0 24.2-7.5 31.4-21.9 31.4c-9 0-18.4-5.8-24.8-18.5L272.9 732c13.4 22.9 32.3 34.2 61.3 34.2c41.6 0 60.8-29.9 60.8-66.2V577h-45zM501.3 577H437v186h44v-62h21.6c39.1 0 73.1-19.6 73.1-63.6c0-45.8-33.5-60.4-74.4-60.4m-.8 89H481v-53h18.2c21.5 0 33.4 6.2 33.4 24.9c0 18.1-10.5 28.1-32.1 28.1m182.5-9v36h30v30.1c-4 2.9-11 4.7-17.7 4.7c-34.3 0-50.7-21.4-50.7-58.2c0-36.1 19.7-57.4 47.1-57.4c15.3 0 25 6.2 34 14.4l23.7-28.3c-12.7-12.8-32.1-24.2-59.2-24.2c-49.6 0-91.1 35.3-91.1 97c0 62.7 40 95.1 91.5 95.1c25.9 0 49.2-10.2 61.5-22.6V657z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileMarkdownFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M854.6 288.7c6 6 9.4 14.1 9.4 22.6V928c0 17.7-14.3 32-32 32H192c-17.7 0-32-14.3-32-32V96c0-17.7 14.3-32 32-32h424.7c8.5 0 16.7 3.4 22.7 9.4zM790.2 326L602 137.8V326zM426.13 600.93l59.11 132.97a16 16 0 0 0 14.62 9.5h24.06a16 16 0 0 0 14.63-9.51l59.1-133.35V758a16 16 0 0 0 16.01 16H641a16 16 0 0 0 16-16V486a16 16 0 0 0-16-16h-34.75a16 16 0 0 0-14.67 9.62L512.1 662.2l-79.48-182.59a16 16 0 0 0-14.67-9.61H383a16 16 0 0 0-16 16v272a16 16 0 0 0 16 16h27.13a16 16 0 0 0 16-16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileMarkdownOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M854.6 288.6L639.4 73.4c-6-6-14.1-9.4-22.6-9.4H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V311.3c0-8.5-3.4-16.7-9.4-22.7M790.2 326H602V137.8zm1.8 562H232V136h302v216a42 42 0 0 0 42 42h216zM429 481.2c-1.9-4.4-6.2-7.2-11-7.2h-35c-6.6 0-12 5.4-12 12v272c0 6.6 5.4 12 12 12h27.1c6.6 0 12-5.4 12-12V582.1l66.8 150.2a12 12 0 0 0 11 7.1H524c4.7 0 9-2.8 11-7.1l66.8-150.6V758c0 6.6 5.4 12 12 12H641c6.6 0 12-5.4 12-12V486c0-6.6-5.4-12-12-12h-34.7c-4.8 0-9.1 2.8-11 7.2l-83.1 191z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileMarkdownTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M534 352V136H232v752h560V394H576a42 42 0 0 1-42-42m72.3 122H641c6.6 0 12 5.4 12 12v272c0 6.6-5.4 12-12 12h-27.2c-6.6 0-12-5.4-12-12V581.7L535 732.3c-2 4.3-6.3 7.1-11 7.1h-24.1a12 12 0 0 1-11-7.1l-66.8-150.2V758c0 6.6-5.4 12-12 12H383c-6.6 0-12-5.4-12-12V486c0-6.6 5.4-12 12-12h35c4.8 0 9.1 2.8 11 7.2l83.2 191l83.1-191c1.9-4.4 6.2-7.2 11-7.2"/><path fill="currentColor" d="M854.6 288.6L639.4 73.4c-6-6-14.1-9.4-22.6-9.4H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V311.3c0-8.5-3.4-16.7-9.4-22.7M602 137.8L790.2 326H602zM792 888H232V136h302v216a42 42 0 0 0 42 42h216z"/><path fill="currentColor" d="M429 481.2c-1.9-4.4-6.2-7.2-11-7.2h-35c-6.6 0-12 5.4-12 12v272c0 6.6 5.4 12 12 12h27.1c6.6 0 12-5.4 12-12V582.1l66.8 150.2a12 12 0 0 0 11 7.1H524c4.7 0 9-2.8 11-7.1l66.8-150.6V758c0 6.6 5.4 12 12 12H641c6.6 0 12-5.4 12-12V486c0-6.6-5.4-12-12-12h-34.7c-4.8 0-9.1 2.8-11 7.2l-83.1 191z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M854.6 288.6L639.4 73.4c-6-6-14.1-9.4-22.6-9.4H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V311.3c0-8.5-3.4-16.7-9.4-22.7M790.2 326H602V137.8zm1.8 562H232V136h302v216a42 42 0 0 0 42 42h216z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePdfFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M854.6 288.7c6 6 9.4 14.1 9.4 22.6V928c0 17.7-14.3 32-32 32H192c-17.7 0-32-14.3-32-32V96c0-17.7 14.3-32 32-32h424.7c8.5 0 16.7 3.4 22.7 9.4zM790.2 326L602 137.8V326zM633.22 637.26c-15.18-.5-31.32.67-49.65 2.96c-24.3-14.99-40.66-35.58-52.28-65.83l1.07-4.38l1.24-5.18c4.3-18.13 6.61-31.36 7.3-44.7c.52-10.07-.04-19.36-1.83-27.97c-3.3-18.59-16.45-29.46-33.02-30.13c-15.45-.63-29.65 8-33.28 21.37c-5.91 21.62-2.45 50.07 10.08 98.59c-15.96 38.05-37.05 82.66-51.2 107.54c-18.89 9.74-33.6 18.6-45.96 28.42c-16.3 12.97-26.48 26.3-29.28 40.3c-1.36 6.49.69 14.97 5.36 21.92c5.3 7.88 13.28 13 22.85 13.74c24.15 1.87 53.83-23.03 86.6-79.26c3.29-1.1 6.77-2.26 11.02-3.7l11.9-4.02c7.53-2.54 12.99-4.36 18.39-6.11c23.4-7.62 41.1-12.43 57.2-15.17c27.98 14.98 60.32 24.8 82.1 24.8c17.98 0 30.13-9.32 34.52-23.99c3.85-12.88.8-27.82-7.48-36.08c-8.56-8.41-24.3-12.43-45.65-13.12M385.23 765.68v-.36l.13-.34a54.86 54.86 0 0 1 5.6-10.76c4.28-6.58 10.17-13.5 17.47-20.87c3.92-3.95 8-7.8 12.79-12.12c1.07-.96 7.91-7.05 9.19-8.25l11.17-10.4l-8.12 12.93c-12.32 19.64-23.46 33.78-33 43c-3.51 3.4-6.6 5.9-9.1 7.51a16.43 16.43 0 0 1-2.61 1.42c-.41.17-.77.27-1.13.3a2.2 2.2 0 0 1-1.12-.15a2.07 2.07 0 0 1-1.27-1.91M511.17 547.4l-2.26 4l-1.4-4.38c-3.1-9.83-5.38-24.64-6.01-38c-.72-15.2.49-24.32 5.29-24.32c6.74 0 9.83 10.8 10.07 27.05c.22 14.28-2.03 29.14-5.7 35.65zm-5.81 58.46l1.53-4.05l2.09 3.8c11.69 21.24 26.86 38.96 43.54 51.31l3.6 2.66l-4.39.9c-16.33 3.38-31.54 8.46-52.34 16.85c2.17-.88-21.62 8.86-27.64 11.17l-5.25 2.01l2.8-4.88c12.35-21.5 23.76-47.32 36.05-79.77zm157.62 76.26c-7.86 3.1-24.78.33-54.57-12.39l-7.56-3.22l8.2-.6c23.3-1.73 39.8-.45 49.42 3.07c4.1 1.5 6.83 3.39 8.04 5.55a4.64 4.64 0 0 1-1.36 6.31a6.7 6.7 0 0 1-2.17 1.28"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePdfOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m531.3 574.4l.3-1.4c5.8-23.9 13.1-53.7 7.4-80.7c-3.8-21.3-19.5-29.6-32.9-30.2c-15.8-.7-29.9 8.3-33.4 21.4c-6.6 24-.7 56.8 10.1 98.6c-13.6 32.4-35.3 79.5-51.2 107.5c-29.6 15.3-69.3 38.9-75.2 68.7c-1.2 5.5.2 12.5 3.5 18.8c3.7 7 9.6 12.4 16.5 15c3 1.1 6.6 2 10.8 2c17.6 0 46.1-14.2 84.1-79.4c5.8-1.9 11.8-3.9 17.6-5.9c27.2-9.2 55.4-18.8 80.9-23.1c28.2 15.1 60.3 24.8 82.1 24.8c21.6 0 30.1-12.8 33.3-20.5c5.6-13.5 2.9-30.5-6.2-39.6c-13.2-13-45.3-16.4-95.3-10.2c-24.6-15-40.7-35.4-52.4-65.8M421.6 726.3c-13.9 20.2-24.4 30.3-30.1 34.7c6.7-12.3 19.8-25.3 30.1-34.7m87.6-235.5c5.2 8.9 4.5 35.8.5 49.4c-4.9-19.9-5.6-48.1-2.7-51.4c.8.1 1.5.7 2.2 2m-1.6 120.5c10.7 18.5 24.2 34.4 39.1 46.2c-21.6 4.9-41.3 13-58.9 20.2c-4.2 1.7-8.3 3.4-12.3 5c13.3-24.1 24.4-51.4 32.1-71.4m155.6 65.5c.1.2.2.5-.4.9h-.2l-.2.3c-.8.5-9 5.3-44.3-8.6c40.6-1.9 45 7.3 45.1 7.4m191.4-388.2L639.4 73.4c-6-6-14.1-9.4-22.6-9.4H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V311.3c0-8.5-3.4-16.7-9.4-22.7M790.2 326H602V137.8zm1.8 562H232V136h302v216a42 42 0 0 0 42 42h216z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePdfTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M509.2 490.8c-.7-1.3-1.4-1.9-2.2-2c-2.9 3.3-2.2 31.5 2.7 51.4c4-13.6 4.7-40.5-.5-49.4m-1.6 120.5c-7.7 20-18.8 47.3-32.1 71.4c4-1.6 8.1-3.3 12.3-5c17.6-7.2 37.3-15.3 58.9-20.2c-14.9-11.8-28.4-27.7-39.1-46.2"/><path fill="currentColor" fill-opacity=".15" d="M534 352V136H232v752h560V394H576a42 42 0 0 1-42-42m55 287.6c16.1-1.9 30.6-2.8 44.3-2.3c12.8.4 23.6 2 32 5.1c.2.1.3.1.5.2c.4.2.8.3 1.2.5c.5.2 1.1.4 1.6.7c.1.1.3.1.4.2c4.1 1.8 7.5 4 10.1 6.6c9.1 9.1 11.8 26.1 6.2 39.6c-3.2 7.7-11.7 20.5-33.3 20.5c-21.8 0-53.9-9.7-82.1-24.8c-25.5 4.3-53.7 13.9-80.9 23.1c-5.8 2-11.8 4-17.6 5.9c-38 65.2-66.5 79.4-84.1 79.4c-4.2 0-7.8-.9-10.8-2c-6.9-2.6-12.8-8-16.5-15c-.9-1.7-1.6-3.4-2.2-5.2c-1.6-4.8-2.1-9.6-1.3-13.6l.6-2.7c.1-.2.1-.4.2-.6c.2-.7.4-1.4.7-2.1c0-.1.1-.2.1-.3c4.1-11.9 13.6-23.4 27.7-34.6c12.3-9.8 27.1-18.7 45.9-28.4c15.9-28 37.6-75.1 51.2-107.4c-10.8-41.8-16.7-74.6-10.1-98.6c.9-3.3 2.5-6.4 4.6-9.1c.2-.2.3-.4.5-.6c.1-.1.1-.2.2-.2c6.3-7.5 16.9-11.9 28.1-11.5c16.6.7 29.7 11.5 33 30.1c1.7 8 2.2 16.5 1.9 25.7v.7c0 .5 0 1-.1 1.5c-.7 13.3-3 26.6-7.3 44.7c-.4 1.6-.8 3.2-1.2 5.2l-1 4.1l-.1.3c.1.2.1.3.2.5l1.8 4.5c.1.3.3.7.4 1c.7 1.6 1.4 3.3 2.1 4.8v.1c8.7 18.8 19.7 33.4 33.9 45.1c4.3 3.5 8.9 6.7 13.9 9.8c1.8-.5 3.5-.7 5.3-.9"/><path fill="currentColor" fill-opacity=".15" d="M391.5 761c5.7-4.4 16.2-14.5 30.1-34.7c-10.3 9.4-23.4 22.4-30.1 34.7m270.9-83l.2-.3h.2c.6-.4.5-.7.4-.9c-.1-.1-4.5-9.3-45.1-7.4c35.3 13.9 43.5 9.1 44.3 8.6"/><path fill="currentColor" d="M854.6 288.6L639.4 73.4c-6-6-14.1-9.4-22.6-9.4H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V311.3c0-8.5-3.4-16.7-9.4-22.7M602 137.8L790.2 326H602zM792 888H232V136h302v216a42 42 0 0 0 42 42h216z"/><path fill="currentColor" d="M535.9 585.3c-.8-1.7-1.5-3.3-2.2-4.9c-.1-.3-.3-.7-.4-1l-1.8-4.5c-.1-.2-.1-.3-.2-.5l.1-.3l.2-1.1c4-16.3 8.6-35.3 9.4-54.4v-.7c.3-8.6-.2-17.2-2-25.6c-3.8-21.3-19.5-29.6-32.9-30.2c-11.3-.5-21.8 4-28.1 11.4c-.1.1-.1.2-.2.2c-.2.2-.4.4-.5.6c-2.1 2.7-3.7 5.8-4.6 9.1c-6.6 24-.7 56.8 10.1 98.6c-13.6 32.4-35.3 79.4-51.2 107.4v.1c-27.7 14.3-64.1 35.8-73.6 62.9c0 .1-.1.2-.1.3c-.2.7-.5 1.4-.7 2.1c-.1.2-.1.4-.2.6c-.2.9-.5 1.8-.6 2.7c-.9 4-.4 8.8 1.3 13.6c.6 1.8 1.3 3.5 2.2 5.2c3.7 7 9.6 12.4 16.5 15c3 1.1 6.6 2 10.8 2c17.6 0 46.1-14.2 84.1-79.4c5.8-1.9 11.8-3.9 17.6-5.9c27.2-9.2 55.4-18.8 80.9-23.1c28.2 15.1 60.3 24.8 82.1 24.8c21.6 0 30.1-12.8 33.3-20.5c5.6-13.5 2.9-30.5-6.2-39.6c-2.6-2.6-6-4.8-10.1-6.6c-.1-.1-.3-.1-.4-.2c-.5-.2-1.1-.4-1.6-.7c-.4-.2-.8-.3-1.2-.5c-.2-.1-.3-.1-.5-.2c-16.2-5.8-41.7-6.7-76.3-2.8l-5.3.6c-5-3-9.6-6.3-13.9-9.8c-14.2-11.3-25.1-25.8-33.8-44.7M391.5 761c6.7-12.3 19.8-25.3 30.1-34.7c-13.9 20.2-24.4 30.3-30.1 34.7M507 488.8c.8.1 1.5.7 2.2 2c5.2 8.9 4.5 35.8.5 49.4c-4.9-19.9-5.6-48.1-2.7-51.4m-19.2 188.9c-4.2 1.7-8.3 3.4-12.3 5c13.3-24.1 24.4-51.4 32.1-71.4c10.7 18.5 24.2 34.4 39.1 46.2c-21.6 4.9-41.3 13-58.9 20.2m175.4-.9c.1.2.2.5-.4.9h-.2l-.2.3c-.8.5-9 5.3-44.3-8.6c40.6-1.9 45 7.3 45.1 7.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePptFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M854.6 288.7c6 6 9.4 14.1 9.4 22.6V928c0 17.7-14.3 32-32 32H192c-17.7 0-32-14.3-32-32V96c0-17.7 14.3-32 32-32h424.7c8.5 0 16.7 3.4 22.7 9.4zM790.2 326L602 137.8V326zM468.53 760v-91.54h59.27c60.57 0 100.2-39.65 100.2-98.12c0-58.22-39.58-98.34-99.98-98.34H424a12 12 0 0 0-12 12v276a12 12 0 0 0 12 12h32.53a12 12 0 0 0 12-12m0-139.33h34.9c47.82 0 67.19-12.93 67.19-50.33c0-32.05-18.12-50.12-49.87-50.12h-52.22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePptOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M424 476c-4.4 0-8 3.6-8 8v276c0 4.4 3.6 8 8 8h32.5c4.4 0 8-3.6 8-8v-95.5h63.3c59.4 0 96.2-38.9 96.2-94.1c0-54.5-36.3-94.3-96-94.3H424zm150.6 94.3c0 43.4-26.5 54.3-71.2 54.3h-38.9V516.2h56.2c33.8 0 53.9 19.7 53.9 54.1m280-281.7L639.4 73.4c-6-6-14.1-9.4-22.6-9.4H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V311.3c0-8.5-3.4-16.7-9.4-22.7M790.2 326H602V137.8zm1.8 562H232V136h302v216a42 42 0 0 0 42 42h216z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePptTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M464.5 516.2v108.4h38.9c44.7 0 71.2-10.9 71.2-54.3c0-34.4-20.1-54.1-53.9-54.1z"/><path fill="currentColor" fill-opacity=".15" d="M534 352V136H232v752h560V394H576a42 42 0 0 1-42-42m90 218.4c0 55.2-36.8 94.1-96.2 94.1h-63.3V760c0 4.4-3.6 8-8 8H424c-4.4 0-8-3.6-8-8V484c0-4.4 3.6-8 8-8v.1h104c59.7 0 96 39.8 96 94.3"/><path fill="currentColor" d="M854.6 288.6L639.4 73.4c-6-6-14.1-9.4-22.6-9.4H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V311.3c0-8.5-3.4-16.7-9.4-22.7M602 137.8L790.2 326H602zM792 888H232V136h302v216a42 42 0 0 0 42 42h216z"/><path fill="currentColor" d="M424 476.1c-4.4-.1-8 3.5-8 7.9v276c0 4.4 3.6 8 8 8h32.5c4.4 0 8-3.6 8-8v-95.5h63.3c59.4 0 96.2-38.9 96.2-94.1c0-54.5-36.3-94.3-96-94.3zm150.6 94.2c0 43.4-26.5 54.3-71.2 54.3h-38.9V516.2h56.2c33.8 0 53.9 19.7 53.9 54.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileProtectOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M644.7 669.2a7.92 7.92 0 0 0-6.5-3.3H594c-6.5 0-10.3 7.4-6.5 12.7l73.8 102.1c3.2 4.4 9.7 4.4 12.9 0l114.2-158c3.8-5.3 0-12.7-6.5-12.7h-44.3c-2.6 0-5 1.2-6.5 3.3l-63.5 87.8zM688 306v-48c0-4.4-3.6-8-8-8H296c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h384c4.4 0 8-3.6 8-8m-392 88c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h184c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8zm184 458H208V148h560v296c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V108c0-17.7-14.3-32-32-32H168c-17.7 0-32 14.3-32 32v784c0 17.7 14.3 32 32 32h312c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m402.6-320.8l-192-66.7c-.9-.3-1.7-.4-2.6-.4s-1.8.1-2.6.4l-192 66.7a7.96 7.96 0 0 0-5.4 7.5v251.1c0 2.5 1.1 4.8 3.1 6.3l192 150.2c1.4 1.1 3.2 1.7 4.9 1.7s3.5-.6 4.9-1.7l192-150.2c1.9-1.5 3.1-3.8 3.1-6.3V538.7c0-3.4-2.2-6.4-5.4-7.5M826 763.7L688 871.6L550 763.7V577l138-48l138 48z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileSearchOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M688 312v-48c0-4.4-3.6-8-8-8H296c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h384c4.4 0 8-3.6 8-8m-392 88c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h184c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8zm144 452H208V148h560v344c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V108c0-17.7-14.3-32-32-32H168c-17.7 0-32 14.3-32 32v784c0 17.7 14.3 32 32 32h272c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m445.7 51.5l-93.3-93.3C814.7 780.7 828 743.9 828 704c0-97.2-78.8-176-176-176s-176 78.8-176 176s78.8 176 176 176c35.8 0 69-10.7 96.8-29l94.7 94.7c1.6 1.6 3.6 2.3 5.6 2.3s4.1-.8 5.6-2.3l31-31a7.9 7.9 0 0 0 0-11.2M652 816c-61.9 0-112-50.1-112-112s50.1-112 112-112s112 50.1 112 112s-50.1 112-112 112"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileSyncOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M296 256c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h384c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8zm192 200v-48c0-4.4-3.6-8-8-8H296c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h184c4.4 0 8-3.6 8-8m-48 396H208V148h560v344c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V108c0-17.7-14.3-32-32-32H168c-17.7 0-32 14.3-32 32v784c0 17.7 14.3 32 32 32h272c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m104.1-115.6c1.8-34.5 16.2-66.8 40.8-91.4c26.2-26.2 62-41 99.1-41c37.4 0 72.6 14.6 99.1 41c3.2 3.2 6.3 6.6 9.2 10.1L769.2 673a8 8 0 0 0 3 14.1l93.3 22.5c5 1.2 9.8-2.6 9.9-7.7l.6-95.4a8 8 0 0 0-12.9-6.4l-20.3 15.8C805.4 569.6 748.1 540 684 540c-109.9 0-199.6 86.9-204 195.7c-.2 4.5 3.5 8.3 8 8.3h48.1c4.3 0 7.8-3.3 8-7.6M880 744h-48.1c-4.3 0-7.8 3.3-8 7.6c-1.8 34.5-16.2 66.8-40.8 91.4c-26.2 26.2-62 41-99.1 41c-37.4 0-72.6-14.6-99.1-41c-3.2-3.2-6.3-6.6-9.2-10.1l23.1-17.9a8 8 0 0 0-3-14.1l-93.3-22.5c-5-1.2-9.8 2.6-9.9 7.7l-.6 95.4a8 8 0 0 0 12.9 6.4l20.3-15.8C562.6 918.4 619.9 948 684 948c109.9 0 199.6-86.9 204-195.7c.2-4.5-3.5-8.3-8-8.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileTextFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M854.6 288.7c6 6 9.4 14.1 9.4 22.6V928c0 17.7-14.3 32-32 32H192c-17.7 0-32-14.3-32-32V96c0-17.7 14.3-32 32-32h424.7c8.5 0 16.7 3.4 22.7 9.4zM790.2 326L602 137.8V326zM320 482a8 8 0 0 0-8 8v48a8 8 0 0 0 8 8h384a8 8 0 0 0 8-8v-48a8 8 0 0 0-8-8zm0 136a8 8 0 0 0-8 8v48a8 8 0 0 0 8 8h184a8 8 0 0 0 8-8v-48a8 8 0 0 0-8-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileTextOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M854.6 288.6L639.4 73.4c-6-6-14.1-9.4-22.6-9.4H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V311.3c0-8.5-3.4-16.7-9.4-22.7M790.2 326H602V137.8zm1.8 562H232V136h302v216a42 42 0 0 0 42 42h216zM504 618H320c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h184c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8M312 490v48c0 4.4 3.6 8 8 8h384c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8H320c-4.4 0-8 3.6-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileTextTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M534 352V136H232v752h560V394H576a42 42 0 0 1-42-42m-22 322c0 4.4-3.6 8-8 8H320c-4.4 0-8-3.6-8-8v-48c0-4.4 3.6-8 8-8h184c4.4 0 8 3.6 8 8zm200-184v48c0 4.4-3.6 8-8 8H320c-4.4 0-8-3.6-8-8v-48c0-4.4 3.6-8 8-8h384c4.4 0 8 3.6 8 8"/><path fill="currentColor" d="M854.6 288.6L639.4 73.4c-6-6-14.1-9.4-22.6-9.4H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V311.3c0-8.5-3.4-16.7-9.4-22.7M602 137.8L790.2 326H602zM792 888H232V136h302v216a42 42 0 0 0 42 42h216z"/><path fill="currentColor" d="M312 490v48c0 4.4 3.6 8 8 8h384c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8H320c-4.4 0-8 3.6-8 8m192 128H320c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h184c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M534 352V136H232v752h560V394H576a42 42 0 0 1-42-42"/><path fill="currentColor" d="M854.6 288.6L639.4 73.4c-6-6-14.1-9.4-22.6-9.4H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V311.3c0-8.5-3.4-16.7-9.4-22.7M602 137.8L790.2 326H602zM792 888H232V136h302v216a42 42 0 0 0 42 42h216z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileUnknownFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M854.6 288.7c6 6 9.4 14.1 9.4 22.6V928c0 17.7-14.3 32-32 32H192c-17.7 0-32-14.3-32-32V96c0-17.7 14.3-32 32-32h424.7c8.5 0 16.7 3.4 22.7 9.4zM790.2 326L602 137.8V326zM402 549c0 5.4 4.4 9.5 9.8 9.5h32.4c5.4 0 9.8-4.2 9.8-9.4c0-28.2 25.8-51.6 58-51.6s58 23.4 58 51.5c0 25.3-21 47.2-49.3 50.9c-19.3 2.8-34.5 20.3-34.7 40.1v32c0 5.5 4.5 10 10 10h32c5.5 0 10-4.5 10-10v-12.2c0-6 4-11.5 9.7-13.3c44.6-14.4 75-54 74.3-98.9c-.8-55.5-49.2-100.8-108.5-101.6c-61.4-.7-111.5 45.6-111.5 103m110 227a32 32 0 1 0 0-64a32 32 0 0 0 0 64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileUnknownOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M854.6 288.7L639.4 73.4c-6-6-14.2-9.4-22.7-9.4H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V311.3c0-8.5-3.4-16.6-9.4-22.6M790.2 326H602V137.8zm1.8 562H232V136h302v216a42 42 0 0 0 42 42h216zM402 549c0 5.4 4.4 9.5 9.8 9.5h32.4c5.4 0 9.8-4.2 9.8-9.4c0-28.2 25.8-51.6 58-51.6s58 23.4 58 51.5c0 25.3-21 47.2-49.3 50.9c-19.3 2.8-34.5 20.3-34.7 40.1v32c0 5.5 4.5 10 10 10h32c5.5 0 10-4.5 10-10v-12.2c0-6 4-11.5 9.7-13.3c44.6-14.4 75-54 74.3-98.9c-.8-55.5-49.2-100.8-108.5-101.6c-61.4-.7-111.5 45.6-111.5 103m78 195a32 32 0 1 0 64 0a32 32 0 1 0-64 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileUnknownTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M534 352V136H232v752h560V394H576a42 42 0 0 1-42-42m-22 424c-17.7 0-32-14.3-32-32s14.3-32 32-32s32 14.3 32 32s-14.3 32-32 32m110-228.4c.7 44.9-29.7 84.5-74.3 98.9c-5.7 1.8-9.7 7.3-9.7 13.3V672c0 5.5-4.5 10-10 10h-32c-5.5 0-10-4.5-10-10v-32c.2-19.8 15.4-37.3 34.7-40.1C549 596.2 570 574.3 570 549c0-28.1-25.8-51.5-58-51.5s-58 23.4-58 51.6c0 5.2-4.4 9.4-9.8 9.4h-32.4c-5.4 0-9.8-4.1-9.8-9.5c0-57.4 50.1-103.7 111.5-103c59.3.8 107.7 46.1 108.5 101.6"/><path fill="currentColor" d="M854.6 288.7L639.4 73.4c-6-6-14.2-9.4-22.7-9.4H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V311.3c0-8.5-3.4-16.6-9.4-22.6M602 137.8L790.2 326H602zM792 888H232V136h302v216a42 42 0 0 0 42 42h216z"/><path fill="currentColor" d="M480 744a32 32 0 1 0 64 0a32 32 0 1 0-64 0m-78-195c0 5.4 4.4 9.5 9.8 9.5h32.4c5.4 0 9.8-4.2 9.8-9.4c0-28.2 25.8-51.6 58-51.6s58 23.4 58 51.5c0 25.3-21 47.2-49.3 50.9c-19.3 2.8-34.5 20.3-34.7 40.1v32c0 5.5 4.5 10 10 10h32c5.5 0 10-4.5 10-10v-12.2c0-6 4-11.5 9.7-13.3c44.6-14.4 75-54 74.3-98.9c-.8-55.5-49.2-100.8-108.5-101.6c-61.4-.7-111.5 45.6-111.5 103"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileWordFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M854.6 288.7c6 6 9.4 14.1 9.4 22.6V928c0 17.7-14.3 32-32 32H192c-17.7 0-32-14.3-32-32V96c0-17.7 14.3-32 32-32h424.7c8.5 0 16.7 3.4 22.7 9.4zM790.2 326L602 137.8V326zM512 566.1l52.81 197a12 12 0 0 0 11.6 8.9h31.77a12 12 0 0 0 11.6-8.88l74.37-276a12 12 0 0 0 .4-3.12a12 12 0 0 0-12-12h-35.57a12 12 0 0 0-11.7 9.31l-45.78 199.1l-49.76-199.32A12 12 0 0 0 528.1 472h-32.2a12 12 0 0 0-11.64 9.1L434.6 680.01L388.5 481.3a12 12 0 0 0-11.68-9.29h-35.39a12 12 0 0 0-3.11.41a12 12 0 0 0-8.47 14.7l74.17 276A12 12 0 0 0 415.6 772h31.99a12 12 0 0 0 11.59-8.9l52.81-197z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileWordOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M854.6 288.6L639.4 73.4c-6-6-14.1-9.4-22.6-9.4H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V311.3c0-8.5-3.4-16.7-9.4-22.7M790.2 326H602V137.8zm1.8 562H232V136h302v216a42 42 0 0 0 42 42h216zM528.1 472h-32.2c-5.5 0-10.3 3.7-11.6 9.1L434.6 680l-46.1-198.7c-1.3-5.4-6.1-9.3-11.7-9.3h-35.4a12.02 12.02 0 0 0-11.6 15.1l74.2 276c1.4 5.2 6.2 8.9 11.6 8.9h32c5.4 0 10.2-3.6 11.6-8.9l52.8-197l52.8 197c1.4 5.2 6.2 8.9 11.6 8.9h31.8c5.4 0 10.2-3.6 11.6-8.9l74.4-276a12.04 12.04 0 0 0-11.6-15.1H647c-5.6 0-10.4 3.9-11.7 9.3l-45.8 199.1l-49.8-199.3c-1.3-5.4-6.1-9.1-11.6-9.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileWordTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M534 352V136H232v752h560V394H576a42 42 0 0 1-42-42m101.3 129.3c1.3-5.4 6.1-9.3 11.7-9.3h35.6a12.04 12.04 0 0 1 11.6 15.1l-74.4 276c-1.4 5.3-6.2 8.9-11.6 8.9h-31.8c-5.4 0-10.2-3.7-11.6-8.9l-52.8-197l-52.8 197c-1.4 5.3-6.2 8.9-11.6 8.9h-32c-5.4 0-10.2-3.7-11.6-8.9l-74.2-276a12.02 12.02 0 0 1 11.6-15.1h35.4c5.6 0 10.4 3.9 11.7 9.3L434.6 680l49.7-198.9c1.3-5.4 6.1-9.1 11.6-9.1h32.2c5.5 0 10.3 3.7 11.6 9.1l49.8 199.3z"/><path fill="currentColor" d="M854.6 288.6L639.4 73.4c-6-6-14.1-9.4-22.6-9.4H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V311.3c0-8.5-3.4-16.7-9.4-22.7M602 137.8L790.2 326H602zM792 888H232V136h302v216a42 42 0 0 0 42 42h216z"/><path fill="currentColor" d="M528.1 472h-32.2c-5.5 0-10.3 3.7-11.6 9.1L434.6 680l-46.1-198.7c-1.3-5.4-6.1-9.3-11.7-9.3h-35.4a12.02 12.02 0 0 0-11.6 15.1l74.2 276c1.4 5.2 6.2 8.9 11.6 8.9h32c5.4 0 10.2-3.6 11.6-8.9l52.8-197l52.8 197c1.4 5.2 6.2 8.9 11.6 8.9h31.8c5.4 0 10.2-3.6 11.6-8.9l74.4-276a12.04 12.04 0 0 0-11.6-15.1H647c-5.6 0-10.4 3.9-11.7 9.3l-45.8 199.1l-49.8-199.3c-1.3-5.4-6.1-9.1-11.6-9.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileZipFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M854.6 288.7c6 6 9.4 14.1 9.4 22.6V928c0 17.7-14.3 32-32 32H192c-17.7 0-32-14.3-32-32V96c0-17.7 14.3-32 32-32h424.7c8.5 0 16.7 3.4 22.7 9.4zM790.2 326L602 137.8V326zM296 136v64h64v-64zm64 64v64h64v-64zm-64 64v64h64v-64zm64 64v64h64v-64zm-64 64v64h64v-64zm64 64v64h64v-64zm-64 64v64h64v-64zm0 64v160h128V584zm48 48h32v64h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileZipOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M296 392h64v64h-64zm0 190v160h128V582h-64v-62h-64zm80 48v64h-32v-64zm-16-302h64v64h-64zm-64-64h64v64h-64zm64 192h64v64h-64zm0-256h64v64h-64zm494.6 88.6L639.4 73.4c-6-6-14.1-9.4-22.6-9.4H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V311.3c0-8.5-3.4-16.7-9.4-22.7M790.2 326H602V137.8zm1.8 562H232V136h64v64h64v-64h174v216a42 42 0 0 0 42 42h216z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileZipTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M344 630h32v2h-32z"/><path fill="currentColor" fill-opacity=".15" d="M534 352V136H360v64h64v64h-64v64h64v64h-64v64h64v64h-64v62h64v160H296V520h64v-64h-64v-64h64v-64h-64v-64h64v-64h-64v-64h-64v752h560V394H576a42 42 0 0 1-42-42"/><path fill="currentColor" d="M854.6 288.6L639.4 73.4c-6-6-14.1-9.4-22.6-9.4H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V311.3c0-8.5-3.4-16.7-9.4-22.7M602 137.8L790.2 326H602zM792 888H232V136h64v64h64v-64h174v216a42 42 0 0 0 42 42h216z"/><path fill="currentColor" d="M296 392h64v64h-64zm0-128h64v64h-64zm0 318v160h128V582h-64v-62h-64zm48 50v-2h32v64h-32zm16-432h64v64h-64zm0 256h64v64h-64zm0-128h64v64h-64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M349 838c0 17.7 14.2 32 31.8 32h262.4c17.6 0 31.8-14.3 31.8-32V642H349zm531.1-684H143.9c-24.5 0-39.8 26.7-27.5 48l221.3 376h348.8l221.3-376c12.1-21.3-3.2-48-27.7-48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880.1 154H143.9c-24.5 0-39.8 26.7-27.5 48L349 597.4V838c0 17.7 14.2 32 31.8 32h262.4c17.6 0 31.8-14.3 31.8-32V597.4L907.7 202c12.2-21.3-3.1-48-27.6-48M603.4 798H420.6V642h182.9v156zm9.6-236.6l-9.5 16.6h-183l-9.5-16.6L212.7 226h598.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M420.6 798h182.9V642H420.6zM411 561.4l9.5 16.6h183l9.5-16.6L811.3 226H212.7z"/><path fill="currentColor" d="M880.1 154H143.9c-24.5 0-39.8 26.7-27.5 48L349 597.4V838c0 17.7 14.2 32 31.8 32h262.4c17.6 0 31.8-14.3 31.8-32V597.4L907.7 202c12.2-21.3-3.1-48-27.6-48M603.5 798H420.6V642h182.9zm9.5-236.6l-9.5 16.6h-183l-9.5-16.6L212.7 226h598.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FireFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M834.1 469.2A347.49 347.49 0 0 0 751.2 354l-29.1-26.7a8.09 8.09 0 0 0-13 3.3l-13 37.3c-8.1 23.4-23 47.3-44.1 70.8c-1.4 1.5-3 1.9-4.1 2c-1.1.1-2.8-.1-4.3-1.5c-1.4-1.2-2.1-3-2-4.8c3.7-60.2-14.3-128.1-53.7-202C555.3 171 510 123.1 453.4 89.7l-41.3-24.3c-5.4-3.2-12.3 1-12 7.3l2.2 48c1.5 32.8-2.3 61.8-11.3 85.9c-11 29.5-26.8 56.9-47 81.5a295.64 295.64 0 0 1-47.5 46.1a352.6 352.6 0 0 0-100.3 121.5A347.75 347.75 0 0 0 160 610c0 47.2 9.3 92.9 27.7 136a349.4 349.4 0 0 0 75.5 110.9c32.4 32 70 57.2 111.9 74.7C418.5 949.8 464.5 959 512 959s93.5-9.2 136.9-27.3A348.6 348.6 0 0 0 760.8 857c32.4-32 57.8-69.4 75.5-110.9a344.2 344.2 0 0 0 27.7-136c0-48.8-10-96.2-29.9-140.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FireOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M834.1 469.2A347.49 347.49 0 0 0 751.2 354l-29.1-26.7a8.09 8.09 0 0 0-13 3.3l-13 37.3c-8.1 23.4-23 47.3-44.1 70.8c-1.4 1.5-3 1.9-4.1 2c-1.1.1-2.8-.1-4.3-1.5c-1.4-1.2-2.1-3-2-4.8c3.7-60.2-14.3-128.1-53.7-202C555.3 171 510 123.1 453.4 89.7l-41.3-24.3c-5.4-3.2-12.3 1-12 7.3l2.2 48c1.5 32.8-2.3 61.8-11.3 85.9c-11 29.5-26.8 56.9-47 81.5a295.64 295.64 0 0 1-47.5 46.1a352.6 352.6 0 0 0-100.3 121.5A347.75 347.75 0 0 0 160 610c0 47.2 9.3 92.9 27.7 136a349.4 349.4 0 0 0 75.5 110.9c32.4 32 70 57.2 111.9 74.7C418.5 949.8 464.5 959 512 959s93.5-9.2 136.9-27.3A348.6 348.6 0 0 0 760.8 857c32.4-32 57.8-69.4 75.5-110.9a344.2 344.2 0 0 0 27.7-136c0-48.8-10-96.2-29.9-140.9M713 808.5c-53.7 53.2-125 82.4-201 82.4s-147.3-29.2-201-82.4c-53.5-53.1-83-123.5-83-198.4c0-43.5 9.8-85.2 29.1-124c18.8-37.9 46.8-71.8 80.8-97.9a349.6 349.6 0 0 0 58.6-56.8c25-30.5 44.6-64.5 58.2-101a240 240 0 0 0 12.1-46.5c24.1 22.2 44.3 49 61.2 80.4c33.4 62.6 48.8 118.3 45.8 165.7a74.01 74.01 0 0 0 24.4 59.8a73.36 73.36 0 0 0 53.4 18.8c19.7-1 37.8-9.7 51-24.4c13.3-14.9 24.8-30.1 34.4-45.6c14 17.9 25.7 37.4 35 58.4c15.9 35.8 24 73.9 24 113.1c0 74.9-29.5 145.4-83 198.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FireTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M737 438.6c-9.6 15.5-21.1 30.7-34.4 45.6a73.1 73.1 0 0 1-51 24.4a73.36 73.36 0 0 1-53.4-18.8a74.01 74.01 0 0 1-24.4-59.8c3-47.4-12.4-103.1-45.8-165.7c-16.9-31.4-37.1-58.2-61.2-80.4a240 240 0 0 1-12.1 46.5a354.26 354.26 0 0 1-58.2 101a349.6 349.6 0 0 1-58.6 56.8c-34 26.1-62 60-80.8 97.9a275.96 275.96 0 0 0-29.1 124c0 74.9 29.5 145.3 83 198.4c53.7 53.2 125 82.4 201 82.4s147.3-29.2 201-82.4c53.5-53 83-123.5 83-198.4c0-39.2-8.1-77.3-24-113.1c-9.3-21-21-40.5-35-58.4"/><path fill="currentColor" d="M834.1 469.2A347.49 347.49 0 0 0 751.2 354l-29.1-26.7a8.09 8.09 0 0 0-13 3.3l-13 37.3c-8.1 23.4-23 47.3-44.1 70.8c-1.4 1.5-3 1.9-4.1 2c-1.1.1-2.8-.1-4.3-1.5c-1.4-1.2-2.1-3-2-4.8c3.7-60.2-14.3-128.1-53.7-202C555.3 171 510 123.1 453.4 89.7l-41.3-24.3c-5.4-3.2-12.3 1-12 7.3l2.2 48c1.5 32.8-2.3 61.8-11.3 85.9c-11 29.5-26.8 56.9-47 81.5a295.64 295.64 0 0 1-47.5 46.1a352.6 352.6 0 0 0-100.3 121.5A347.75 347.75 0 0 0 160 610c0 47.2 9.3 92.9 27.7 136a349.4 349.4 0 0 0 75.5 110.9c32.4 32 70 57.2 111.9 74.7C418.5 949.8 464.5 959 512 959s93.5-9.2 136.9-27.3A348.6 348.6 0 0 0 760.8 857c32.4-32 57.8-69.4 75.5-110.9a344.2 344.2 0 0 0 27.7-136c0-48.8-10-96.2-29.9-140.9M713 808.5c-53.7 53.2-125 82.4-201 82.4s-147.3-29.2-201-82.4c-53.5-53.1-83-123.5-83-198.4c0-43.5 9.8-85.2 29.1-124c18.8-37.9 46.8-71.8 80.8-97.9a349.6 349.6 0 0 0 58.6-56.8c25-30.5 44.6-64.5 58.2-101a240 240 0 0 0 12.1-46.5c24.1 22.2 44.3 49 61.2 80.4c33.4 62.6 48.8 118.3 45.8 165.7a74.01 74.01 0 0 0 24.4 59.8a73.36 73.36 0 0 0 53.4 18.8c19.7-1 37.8-9.7 51-24.4c13.3-14.9 24.8-30.1 34.4-45.6c14 17.9 25.7 37.4 35 58.4c15.9 35.8 24 73.9 24 113.1c0 74.9-29.5 145.4-83 198.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 305H624V192c0-17.7-14.3-32-32-32H184v-40c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v784c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V640h248v113c0 17.7 14.3 32 32 32h416c17.7 0 32-14.3 32-32V337c0-17.7-14.3-32-32-32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 305H624V192c0-17.7-14.3-32-32-32H184v-40c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v784c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V640h248v113c0 17.7 14.3 32 32 32h416c17.7 0 32-14.3 32-32V337c0-17.7-14.3-32-32-32M184 568V232h368v336zm656 145H504v-73h112c4.4 0 8-3.6 8-8V377h216z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M184 232h368v336H184z"/><path fill="currentColor" fill-opacity=".15" d="M624 632c0 4.4-3.6 8-8 8H504v73h336V377H624z"/><path fill="currentColor" d="M880 305H624V192c0-17.7-14.3-32-32-32H184v-40c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v784c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V640h248v113c0 17.7 14.3 32 32 32h416c17.7 0 32-14.3 32-32V337c0-17.7-14.3-32-32-32M184 568V232h368v336zm656 145H504v-73h112c4.4 0 8-3.6 8-8V377h216z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderAddFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 298.4H521L403.7 186.2a8.15 8.15 0 0 0-5.5-2.2H144c-17.7 0-32 14.3-32 32v592c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V330.4c0-17.7-14.3-32-32-32M632 577c0 3.8-3.4 7-7.5 7H540v84.9c0 3.9-3.2 7.1-7 7.1h-42c-3.8 0-7-3.2-7-7.1V584h-84.5c-4.1 0-7.5-3.2-7.5-7v-42c0-3.8 3.4-7 7.5-7H484v-84.9c0-3.9 3.2-7.1 7-7.1h42c3.8 0 7 3.2 7 7.1V528h84.5c4.1 0 7.5 3.2 7.5 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderAddOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M484 443.1V528h-84.5c-4.1 0-7.5 3.1-7.5 7v42c0 3.8 3.4 7 7.5 7H484v84.9c0 3.9 3.2 7.1 7 7.1h42c3.9 0 7-3.2 7-7.1V584h84.5c4.1 0 7.5-3.2 7.5-7v-42c0-3.9-3.4-7-7.5-7H540v-84.9c0-3.9-3.1-7.1-7-7.1h-42c-3.8 0-7 3.2-7 7.1m396-144.7H521L403.7 186.2a8.15 8.15 0 0 0-5.5-2.2H144c-17.7 0-32 14.3-32 32v592c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V330.4c0-17.7-14.3-32-32-32M840 768H184V256h188.5l119.6 114.4H840z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderAddTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M372.5 256H184v512h656V370.4H492.1zM540 443.1V528h84.5c4.1 0 7.5 3.1 7.5 7v42c0 3.8-3.4 7-7.5 7H540v84.9c0 3.9-3.1 7.1-7 7.1h-42c-3.8 0-7-3.2-7-7.1V584h-84.5c-4.1 0-7.5-3.2-7.5-7v-42c0-3.9 3.4-7 7.5-7H484v-84.9c0-3.9 3.2-7.1 7-7.1h42c3.9 0 7 3.2 7 7.1"/><path fill="currentColor" d="M880 298.4H521L403.7 186.2a8.15 8.15 0 0 0-5.5-2.2H144c-17.7 0-32 14.3-32 32v592c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V330.4c0-17.7-14.3-32-32-32M840 768H184V256h188.5l119.6 114.4H840z"/><path fill="currentColor" d="M484 443.1V528h-84.5c-4.1 0-7.5 3.1-7.5 7v42c0 3.8 3.4 7 7.5 7H484v84.9c0 3.9 3.2 7.1 7 7.1h42c3.9 0 7-3.2 7-7.1V584h84.5c4.1 0 7.5-3.2 7.5-7v-42c0-3.9-3.4-7-7.5-7H540v-84.9c0-3.9-3.1-7.1-7-7.1h-42c-3.8 0-7 3.2-7 7.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 298.4H521L403.7 186.2a8.15 8.15 0 0 0-5.5-2.2H144c-17.7 0-32 14.3-32 32v592c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V330.4c0-17.7-14.3-32-32-32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpenFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M928 444H820V330.4c0-17.7-14.3-32-32-32H473L355.7 186.2a8.15 8.15 0 0 0-5.5-2.2H96c-17.7 0-32 14.3-32 32v592c0 17.7 14.3 32 32 32h698c13 0 24.8-7.9 29.7-20l134-332c1.5-3.8 2.3-7.9 2.3-12c0-17.7-14.3-32-32-32m-180 0H238c-13 0-24.8 7.9-29.7 20L136 643.2V256h188.5l119.6 114.4H748z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpenOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M928 444H820V330.4c0-17.7-14.3-32-32-32H473L355.7 186.2a8.15 8.15 0 0 0-5.5-2.2H96c-17.7 0-32 14.3-32 32v592c0 17.7 14.3 32 32 32h698c13 0 24.8-7.9 29.7-20l134-332c1.5-3.8 2.3-7.9 2.3-12c0-17.7-14.3-32-32-32M136 256h188.5l119.6 114.4H748V444H238c-13 0-24.8 7.9-29.7 20L136 643.2zm635.3 512H159l103.3-256h612.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpenTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M159 768h612.3l103.4-256H262.3z"/><path fill="currentColor" d="M928 444H820V330.4c0-17.7-14.3-32-32-32H473L355.7 186.2a8.15 8.15 0 0 0-5.5-2.2H96c-17.7 0-32 14.3-32 32v592c0 17.7 14.3 32 32 32h698c13 0 24.8-7.9 29.7-20l134-332c1.5-3.8 2.3-7.9 2.3-12c0-17.7-14.3-32-32-32M136 256h188.5l119.6 114.4H748V444H238c-13 0-24.8 7.9-29.7 20L136 643.2zm635.3 512H159l103.3-256h612.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 298.4H521L403.7 186.2a8.15 8.15 0 0 0-5.5-2.2H144c-17.7 0-32 14.3-32 32v592c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V330.4c0-17.7-14.3-32-32-32M840 768H184V256h188.5l119.6 114.4H840z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 298.4H521L403.7 186.2a8.15 8.15 0 0 0-5.5-2.2H144c-17.7 0-32 14.3-32 32v592c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V330.4c0-17.7-14.3-32-32-32M840 768H184V256h188.5l119.6 114.4H840z"/><path fill="currentColor" fill-opacity=".15" d="M372.5 256H184v512h656V370.4H492.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderViewOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M309.1 554.3c-5.4 11.6-5.4 24.9 0 36.4C353.3 684 421.6 732 512.5 732s159.2-48.1 203.4-141.3c5.4-11.5 5.4-24.8.1-36.3l-.1-.1l-.1-.1C671.7 461 603.4 413 512.5 413s-159.2 48.1-203.4 141.3M512.5 477c62.1 0 107.4 30 141.1 95.5C620 638 574.6 668 512.5 668s-107.4-30-141.1-95.5c33.7-65.5 79-95.5 141.1-95.5"/><path fill="currentColor" d="M457 573a56 56 0 1 0 112 0a56 56 0 1 0-112 0"/><path fill="currentColor" d="M880 298.4H521L403.7 186.2c-1.5-1.4-3.5-2.2-5.5-2.2H144c-17.7 0-32 14.3-32 32v592c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V330.4c0-17.7-14.3-32-32-32M840 768H184V256h188.5l119.6 114.4H840z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FontColorsOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M904 816H120c-4.4 0-8 3.6-8 8v80c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-80c0-4.4-3.6-8-8-8m-650.3-80h85c4.2 0 8-2.7 9.3-6.8l53.7-166h219.2l53.2 166c1.3 4 5 6.8 9.3 6.8h89.1c1.1 0 2.2-.2 3.2-.5a9.7 9.7 0 0 0 6-12.4L573.6 118.6a9.9 9.9 0 0 0-9.2-6.6H462.1c-4.2 0-7.9 2.6-9.2 6.6L244.5 723.1c-.4 1-.5 2.1-.5 3.2c-.1 5.3 4.3 9.7 9.7 9.7m255.9-516.1h4.1l83.8 263.8H424.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FontSizeOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M920 416H616c-4.4 0-8 3.6-8 8v112c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8v-56h60v320h-46c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h164c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8h-46V480h60v56c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V424c0-4.4-3.6-8-8-8M656 296V168c0-4.4-3.6-8-8-8H104c-4.4 0-8 3.6-8 8v128c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-64h168v560h-92c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h264c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-92V232h168v64c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForkOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M752 100c-61.8 0-112 50.2-112 112c0 47.7 29.9 88.5 72 104.6v27.6L512 601.4L312 344.2v-27.6c42.1-16.1 72-56.9 72-104.6c0-61.8-50.2-112-112-112s-112 50.2-112 112c0 50.6 33.8 93.5 80 107.3v34.4c0 9.7 3.3 19.3 9.3 27L476 672.3v33.6c-44.2 15-76 56.9-76 106.1c0 61.8 50.2 112 112 112s112-50.2 112-112c0-49.2-31.8-91-76-106.1v-33.6l226.7-291.6c6-7.7 9.3-17.3 9.3-27v-34.4c46.2-13.8 80-56.7 80-107.3c0-61.8-50.2-112-112-112M224 212a48.01 48.01 0 0 1 96 0a48.01 48.01 0 0 1-96 0m336 600a48.01 48.01 0 0 1-96 0a48.01 48.01 0 0 1 96 0m192-552a48.01 48.01 0 0 1 0-96a48.01 48.01 0 0 1 0 96"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M904 512h-56c-4.4 0-8 3.6-8 8v320H184V184h320c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V520c0-4.4-3.6-8-8-8"/><path fill="currentColor" d="M355.9 534.9L354 653.8c-.1 8.9 7.1 16.2 16 16.2h.4l118-2.9c2-.1 4-.9 5.4-2.3l415.9-415c3.1-3.1 3.1-8.2 0-11.3L785.4 114.3c-1.6-1.6-3.6-2.3-5.7-2.3s-4.1.8-5.7 2.3l-415.8 415a8.3 8.3 0 0 0-2.3 5.6m63.5 23.6L779.7 199l45.2 45.1l-360.5 359.7l-45.7 1.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatPainterFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M840 192h-56v-72c0-13.3-10.7-24-24-24H168c-13.3 0-24 10.7-24 24v272c0 13.3 10.7 24 24 24h592c13.3 0 24-10.7 24-24V256h32v200H465c-22.1 0-40 17.9-40 40v136h-44c-4.4 0-8 3.6-8 8v228c0 1.1.2 2.2.6 3.1c-.4 1.6-.6 3.2-.6 4.9c0 46.4 37.6 84 84 84s84-37.6 84-84c0-1.7-.2-3.3-.6-4.9c.4-1 .6-2 .6-3.1V640c0-4.4-3.6-8-8-8h-44V520h351c22.1 0 40-17.9 40-40V232c0-22.1-17.9-40-40-40"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatPainterOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M840 192h-56v-72c0-13.3-10.7-24-24-24H168c-13.3 0-24 10.7-24 24v272c0 13.3 10.7 24 24 24h592c13.3 0 24-10.7 24-24V256h32v200H465c-22.1 0-40 17.9-40 40v136h-44c-4.4 0-8 3.6-8 8v228c0 .6.1 1.3.2 1.9c-.1 2-.2 4.1-.2 6.1c0 46.4 37.6 84 84 84s84-37.6 84-84c0-2.1-.1-4.1-.2-6.1c.1-.6.2-1.2.2-1.9V640c0-4.4-3.6-8-8-8h-44V520h351c22.1 0 40-17.9 40-40V232c0-22.1-17.9-40-40-40M720 352H208V160h512zM477 876c0 11-9 20-20 20s-20-9-20-20V696h40z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M825.8 498L538.4 249.9c-10.7-9.2-26.4-.9-26.4 14v496.3c0 14.9 15.7 23.2 26.4 14L825.8 526c8.3-7.2 8.3-20.8 0-28m-320 0L218.4 249.9c-10.7-9.2-26.4-.9-26.4 14v496.3c0 14.9 15.7 23.2 26.4 14L505.8 526c4.1-3.6 6.2-8.8 6.2-14c0-5.2-2.1-10.4-6.2-14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M825.8 498L538.4 249.9c-10.7-9.2-26.4-.9-26.4 14v496.3c0 14.9 15.7 23.2 26.4 14L825.8 526c8.3-7.2 8.3-20.8 0-28m-320 0L218.4 249.9c-10.7-9.2-26.4-.9-26.4 14v496.3c0 14.9 15.7 23.2 26.4 14L505.8 526c4.1-3.6 6.2-8.8 6.2-14c0-5.2-2.1-10.4-6.2-14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FrownFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64M288 421a48.01 48.01 0 0 1 96 0a48.01 48.01 0 0 1-96 0m376 272h-48.1c-4.2 0-7.8-3.2-8.1-7.4C604 636.1 562.5 597 512 597s-92.1 39.1-95.8 88.6c-.3 4.2-3.9 7.4-8.1 7.4H360a8 8 0 0 1-8-8.4c4.4-84.3 74.5-151.6 160-151.6s155.6 67.3 160 151.6a8 8 0 0 1-8 8.4m24-224a48.01 48.01 0 0 1 0-96a48.01 48.01 0 0 1 0 96"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FrownOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M288 421a48 48 0 1 0 96 0a48 48 0 1 0-96 0m352 0a48 48 0 1 0 96 0a48 48 0 1 0-96 0M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m263 711c-34.2 34.2-74 61-118.3 79.8C611 874.2 562.3 884 512 884c-50.3 0-99-9.8-144.8-29.2A370.4 370.4 0 0 1 248.9 775c-34.2-34.2-61-74-79.8-118.3C149.8 611 140 562.3 140 512s9.8-99 29.2-144.8A370.4 370.4 0 0 1 249 248.9c34.2-34.2 74-61 118.3-79.8C413 149.8 461.7 140 512 140c50.3 0 99 9.8 144.8 29.2A370.4 370.4 0 0 1 775.1 249c34.2 34.2 61 74 79.8 118.3C874.2 413 884 461.7 884 512s-9.8 99-29.2 144.8A368.89 368.89 0 0 1 775 775M512 533c-85.5 0-155.6 67.3-160 151.6a8 8 0 0 0 8 8.4h48.1c4.2 0 7.8-3.2 8.1-7.4C420 636.1 461.5 597 512 597s92.1 39.1 95.8 88.6c.3 4.2 3.9 7.4 8.1 7.4H664a8 8 0 0 0 8-8.4C667.6 600.3 597.5 533 512 533"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FrownTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372M288 421a48.01 48.01 0 0 1 96 0a48.01 48.01 0 0 1-96 0m376 272h-48.1c-4.2 0-7.8-3.2-8.1-7.4C604 636.1 562.5 597 512 597s-92.1 39.1-95.8 88.6c-.3 4.2-3.9 7.4-8.1 7.4H360a8 8 0 0 1-8-8.4c4.4-84.3 74.5-151.6 160-151.6s155.6 67.3 160 151.6a8 8 0 0 1-8 8.4m24-224a48.01 48.01 0 0 1 0-96a48.01 48.01 0 0 1 0 96"/><path fill="currentColor" d="M288 421a48 48 0 1 0 96 0a48 48 0 1 0-96 0m224 112c-85.5 0-155.6 67.3-160 151.6a8 8 0 0 0 8 8.4h48.1c4.2 0 7.8-3.2 8.1-7.4c3.7-49.5 45.3-88.6 95.8-88.6s92 39.1 95.8 88.6c.3 4.2 3.9 7.4 8.1 7.4H664a8 8 0 0 0 8-8.4C667.6 600.3 597.5 533 512 533m128-112a48 48 0 1 0 96 0a48 48 0 1 0-96 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FullscreenExitOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M391 240.9c-.8-6.6-8.9-9.4-13.6-4.7l-43.7 43.7L200 146.3a8.03 8.03 0 0 0-11.3 0l-42.4 42.3a8.03 8.03 0 0 0 0 11.3L280 333.6l-43.9 43.9a8.01 8.01 0 0 0 4.7 13.6L401 410c5.1.6 9.5-3.7 8.9-8.9zm10.1 373.2L240.8 633c-6.6.8-9.4 8.9-4.7 13.6l43.9 43.9L146.3 824a8.03 8.03 0 0 0 0 11.3l42.4 42.3c3.1 3.1 8.2 3.1 11.3 0L333.7 744l43.7 43.7A8.01 8.01 0 0 0 391 783l18.9-160.1c.6-5.1-3.7-9.4-8.8-8.8m221.8-204.2L783.2 391c6.6-.8 9.4-8.9 4.7-13.6L744 333.6L877.7 200c3.1-3.1 3.1-8.2 0-11.3l-42.4-42.3a8.03 8.03 0 0 0-11.3 0L690.3 279.9l-43.7-43.7a8.01 8.01 0 0 0-13.6 4.7L614.1 401c-.6 5.2 3.7 9.5 8.8 8.9M744 690.4l43.9-43.9a8.01 8.01 0 0 0-4.7-13.6L623 614c-5.1-.6-9.5 3.7-8.9 8.9L633 783.1c.8 6.6 8.9 9.4 13.6 4.7l43.7-43.7L824 877.7c3.1 3.1 8.2 3.1 11.3 0l42.4-42.3c3.1-3.1 3.1-8.2 0-11.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FullscreenOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m290 236.4l43.9-43.9a8.01 8.01 0 0 0-4.7-13.6L169 160c-5.1-.6-9.5 3.7-8.9 8.9L179 329.1c.8 6.6 8.9 9.4 13.6 4.7l43.7-43.7L370 423.7c3.1 3.1 8.2 3.1 11.3 0l42.4-42.3c3.1-3.1 3.1-8.2 0-11.3zm352.7 187.3c3.1 3.1 8.2 3.1 11.3 0l133.7-133.6l43.7 43.7a8.01 8.01 0 0 0 13.6-4.7L863.9 169c.6-5.1-3.7-9.5-8.9-8.9L694.8 179c-6.6.8-9.4 8.9-4.7 13.6l43.9 43.9L600.3 370a8.03 8.03 0 0 0 0 11.3zM845 694.9c-.8-6.6-8.9-9.4-13.6-4.7l-43.7 43.7L654 600.3a8.03 8.03 0 0 0-11.3 0l-42.4 42.3a8.03 8.03 0 0 0 0 11.3L734 787.6l-43.9 43.9a8.01 8.01 0 0 0 4.7 13.6L855 864c5.1.6 9.5-3.7 8.9-8.9zm-463.7-94.6a8.03 8.03 0 0 0-11.3 0L236.3 733.9l-43.7-43.7a8.01 8.01 0 0 0-13.6 4.7L160.1 855c-.6 5.1 3.7 9.5 8.9 8.9L329.2 845c6.6-.8 9.4-8.9 4.7-13.6L290 787.6L423.7 654c3.1-3.1 3.1-8.2 0-11.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FunctionOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M841 370c3-3.3 2.7-8.3-.6-11.3c-1.5-1.3-3.4-2.1-5.3-2.1h-72.6c-2.4 0-4.6 1-6.1 2.8L633.5 504.6c-2.9 3.4-7.9 3.8-11.3.9c-.9-.8-1.6-1.7-2.1-2.8l-63.5-141.3c-1.3-2.9-4.1-4.7-7.3-4.7H380.7l.9-4.7l8-42.3c10.5-55.4 38-81.4 85.8-81.4c18.6 0 35.5 1.7 48.8 4.7l14.1-66.8c-22.6-4.7-35.2-6.1-54.9-6.1c-103.3 0-156.4 44.3-175.9 147.3l-9.4 49.4h-97.6c-3.8 0-7.1 2.7-7.8 6.4L181.9 415c-.9 4.3 1.9 8.6 6.2 9.5c.5.1 1.1.2 1.6.2H284l-89 429.9c-.9 4.3 1.9 8.6 6.2 9.5c.5.1 1.1.2 1.6.2H269c3.8 0 7.1-2.7 7.8-6.4l89.7-433.1h135.8l68.2 139.1c1.4 2.9 1 6.4-1.2 8.8l-180.6 203c-2.9 3.3-2.6 8.4.7 11.3c1.5 1.3 3.4 2 5.3 2h72.7c2.4 0 4.6-1 6.1-2.8l123.7-146.7c2.8-3.4 7.9-3.8 11.3-1c.9.8 1.6 1.7 2.1 2.8L676.4 784c1.3 2.8 4.1 4.7 7.3 4.7h64.6c4.4 0 8-3.6 8-8c0-1.2-.3-2.4-.8-3.5l-95.2-198.9c-1.4-2.9-.9-6.4 1.3-8.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FundFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M926 164H94c-17.7 0-32 14.3-32 32v640c0 17.7 14.3 32 32 32h832c17.7 0 32-14.3 32-32V196c0-17.7-14.3-32-32-32m-92.3 194.4l-297 297.2a8.03 8.03 0 0 1-11.3 0L410.9 541.1L238.4 713.7a8.03 8.03 0 0 1-11.3 0l-36.8-36.8a8.03 8.03 0 0 1 0-11.3l214.9-215c3.1-3.1 8.2-3.1 11.3 0L531 565l254.5-254.6c3.1-3.1 8.2-3.1 11.3 0l36.8 36.8c3.2 3 3.2 8.1.1 11.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FundOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M926 164H94c-17.7 0-32 14.3-32 32v640c0 17.7 14.3 32 32 32h832c17.7 0 32-14.3 32-32V196c0-17.7-14.3-32-32-32m-40 632H134V236h752zm-658.9-82.3c3.1 3.1 8.2 3.1 11.3 0l172.5-172.5l114.4 114.5c3.1 3.1 8.2 3.1 11.3 0l297-297.2c3.1-3.1 3.1-8.2 0-11.3l-36.8-36.8a8.03 8.03 0 0 0-11.3 0L531 565L416.6 450.5a8.03 8.03 0 0 0-11.3 0l-214.9 215a8.03 8.03 0 0 0 0 11.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FundProjectionScreenOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M312.1 591.5c3.1 3.1 8.2 3.1 11.3 0l101.8-101.8l86.1 86.2c3.1 3.1 8.2 3.1 11.3 0l226.3-226.5c3.1-3.1 3.1-8.2 0-11.3l-36.8-36.8c-3.1-3.1-8.2-3.1-11.3 0L517 485.3l-86.1-86.2c-3.1-3.1-8.2-3.1-11.3 0L275.3 543.4c-3.1 3.1-3.1 8.2 0 11.3z"/><path fill="currentColor" d="M904 160H548V96c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v64H120c-17.7 0-32 14.3-32 32v520c0 17.7 14.3 32 32 32h356.4v32L311.6 884.1c-3.7 2.4-4.7 7.3-2.3 11l30.3 47.2v.1c2.4 3.7 7.4 4.7 11.1 2.3L512 838.9l161.3 105.8c3.7 2.4 8.7 1.4 11.1-2.3v-.1l30.3-47.2c2.4-3.7 1.3-8.6-2.3-11L548 776.3V744h356c17.7 0 32-14.3 32-32V192c0-17.7-14.3-32-32-32m-40 512H160V232h704z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FundTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M928 160H96c-17.7 0-32 14.3-32 32v640c0 17.7 14.3 32 32 32h832c17.7 0 32-14.3 32-32V192c0-17.7-14.3-32-32-32m-40 632H136V232h752z"/><path fill="currentColor" fill-opacity=".15" d="M136 792h752V232H136zm56.4-130.5l214.9-215c3.1-3.1 8.2-3.1 11.3 0L533 561l254.5-254.6c3.1-3.1 8.2-3.1 11.3 0l36.8 36.8c3.1 3.1 3.1 8.2 0 11.3l-297 297.2a8.03 8.03 0 0 1-11.3 0L412.9 537.2L240.4 709.7a8.03 8.03 0 0 1-11.3 0l-36.7-36.9a8.03 8.03 0 0 1 0-11.3"/><path fill="currentColor" d="M229.1 709.7c3.1 3.1 8.2 3.1 11.3 0l172.5-172.5l114.4 114.5c3.1 3.1 8.2 3.1 11.3 0l297-297.2c3.1-3.1 3.1-8.2 0-11.3l-36.8-36.8a8.03 8.03 0 0 0-11.3 0L533 561L418.6 446.5a8.03 8.03 0 0 0-11.3 0l-214.9 215a8.03 8.03 0 0 0 0 11.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FundViewOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m956 686.5l-.1-.1l-.1-.1C911.7 593 843.4 545 752.5 545s-159.2 48.1-203.4 141.3v.1c-5.4 11.5-5.4 24.9 0 36.4C593.3 816 661.6 864 752.5 864s159.2-48.1 203.4-141.3c5.4-11.5 5.4-24.8.1-36.2M752.5 800c-62.1 0-107.4-30-141.1-95.5C645 639 690.4 609 752.5 609c62.1 0 107.4 30 141.1 95.5C860 770 814.6 800 752.5 800"/><path fill="currentColor" d="M697 705a56 56 0 1 0 112 0a56 56 0 1 0-112 0M136 232h704v253h72V192c0-17.7-14.3-32-32-32H96c-17.7 0-32 14.3-32 32v520c0 17.7 14.3 32 32 32h352v-72H136z"/><path fill="currentColor" d="m724.9 338.1l-36.8-36.8c-3.1-3.1-8.2-3.1-11.3 0L493 485.3l-86.1-86.2c-3.1-3.1-8.2-3.1-11.3 0L251.3 543.4c-3.1 3.1-3.1 8.2 0 11.3l36.8 36.8c3.1 3.1 8.2 3.1 11.3 0l101.8-101.8l86.1 86.2c3.1 3.1 8.2 3.1 11.3 0l226.3-226.5c3.2-3.1 3.2-8.2 0-11.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FunnelPlotFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336.7 586h350.6l84.9-148H251.8zm543.4-432H143.9c-24.5 0-39.8 26.7-27.5 48L215 374h594l98.7-172c12.2-21.3-3.1-48-27.6-48M349 838c0 17.7 14.2 32 31.8 32h262.4c17.6 0 31.8-14.3 31.8-32V650H349z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FunnelPlotOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880.1 154H143.9c-24.5 0-39.8 26.7-27.5 48L349 607.4V838c0 17.7 14.2 32 31.8 32h262.4c17.6 0 31.8-14.3 31.8-32V607.4L907.7 202c12.2-21.3-3.1-48-27.6-48M603.4 798H420.6V650h182.9v148zm9.6-226.6l-8.4 14.6H419.3l-8.4-14.6L334.4 438h355.2zM726.3 374H297.7l-85-148h598.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FunnelPlotTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M420.6 798h182.9V650H420.6zM297.7 374h428.6l85-148H212.7zm113.2 197.4l8.4 14.6h185.3l8.4-14.6L689.6 438H334.4z"/><path fill="currentColor" d="M880.1 154H143.9c-24.5 0-39.8 26.7-27.5 48L349 607.4V838c0 17.7 14.2 32 31.8 32h262.4c17.6 0 31.8-14.3 31.8-32V607.4L907.7 202c12.2-21.3-3.1-48-27.6-48M603.5 798H420.6V650h182.9zm9.5-226.6l-8.4 14.6H419.3l-8.4-14.6L334.4 438h355.2zM726.3 374H297.7l-85-148h598.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GatewayOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M928 392c8.8 0 16-7.2 16-16V192c0-8.8-7.2-16-16-16H744c-8.8 0-16 7.2-16 16v56H296v-56c0-8.8-7.2-16-16-16H96c-8.8 0-16 7.2-16 16v184c0 8.8 7.2 16 16 16h56v240H96c-8.8 0-16 7.2-16 16v184c0 8.8 7.2 16 16 16h184c8.8 0 16-7.2 16-16v-56h432v56c0 8.8 7.2 16 16 16h184c8.8 0 16-7.2 16-16V648c0-8.8-7.2-16-16-16h-56V392zM792 240h88v88h-88zm-648 88v-88h88v88zm88 456h-88v-88h88zm648-88v88h-88v-88zm-80-64h-56c-8.8 0-16 7.2-16 16v56H296v-56c0-8.8-7.2-16-16-16h-56V392h56c8.8 0 16-7.2 16-16v-56h432v56c0 8.8 7.2 16 16 16h56z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GifOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M944 299H692c-4.4 0-8 3.6-8 8v406c0 4.4 3.6 8 8 8h59.2c4.4 0 8-3.6 8-8V549.9h168.2c4.4 0 8-3.6 8-8V495c0-4.4-3.6-8-8-8H759.2V364.2H944c4.4 0 8-3.6 8-8V307c0-4.4-3.6-8-8-8m-356 1h-56c-4.4 0-8 3.6-8 8v406c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V308c0-4.4-3.6-8-8-8M452 500.9H290.5c-4.4 0-8 3.6-8 8v43.7c0 4.4 3.6 8 8 8h94.9l-.3 8.9c-1.2 58.8-45.6 98.5-110.9 98.5c-76.2 0-123.9-59.7-123.9-156.7c0-95.8 46.8-155.2 121.5-155.2c54.8 0 93.1 26.9 108.5 75.4h76.2c-13.6-87.2-86-143.4-184.7-143.4C150 288 72 375.2 72 511.9C72 650.2 149.1 736 273 736c114.1 0 187-70.7 187-181.6v-45.5c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GiftFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M160 894c0 17.7 14.3 32 32 32h286V550H160zm386 32h286c17.7 0 32-14.3 32-32V550H546zm334-616H732.4c13.6-21.4 21.6-46.8 21.6-74c0-76.1-61.9-138-138-138c-41.4 0-78.7 18.4-104 47.4c-25.3-29-62.6-47.4-104-47.4c-76.1 0-138 61.9-138 138c0 27.2 7.9 52.6 21.6 74H144c-17.7 0-32 14.3-32 32v140h366V310h68v172h366V342c0-17.7-14.3-32-32-32m-402-4h-70c-38.6 0-70-31.4-70-70s31.4-70 70-70s70 31.4 70 70zm138 0h-70v-70c0-38.6 31.4-70 70-70s70 31.4 70 70s-31.4 70-70 70"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GiftOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 310H732.4c13.6-21.4 21.6-46.8 21.6-74c0-76.1-61.9-138-138-138c-41.4 0-78.7 18.4-104 47.4c-25.3-29-62.6-47.4-104-47.4c-76.1 0-138 61.9-138 138c0 27.2 7.9 52.6 21.6 74H144c-17.7 0-32 14.3-32 32v200c0 4.4 3.6 8 8 8h40v344c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V550h40c4.4 0 8-3.6 8-8V342c0-17.7-14.3-32-32-32m-334-74c0-38.6 31.4-70 70-70s70 31.4 70 70s-31.4 70-70 70h-70zm-138-70c38.6 0 70 31.4 70 70v70h-70c-38.6 0-70-31.4-70-70s31.4-70 70-70M180 482V378h298v104zm48 68h250v308H228zm568 308H546V550h250zm48-376H546V378h298z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GiftTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M546 378h298v104H546zM228 550h250v308H228zm-48-172h298v104H180zm366 172h250v308H546z"/><path fill="currentColor" d="M880 310H732.4c13.6-21.4 21.6-46.8 21.6-74c0-76.1-61.9-138-138-138c-41.4 0-78.7 18.4-104 47.4c-25.3-29-62.6-47.4-104-47.4c-76.1 0-138 61.9-138 138c0 27.2 7.9 52.6 21.6 74H144c-17.7 0-32 14.3-32 32v200c0 4.4 3.6 8 8 8h40v344c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V550h40c4.4 0 8-3.6 8-8V342c0-17.7-14.3-32-32-32M478 858H228V550h250zm0-376H180V378h298zm0-176h-70c-38.6 0-70-31.4-70-70s31.4-70 70-70s70 31.4 70 70zm68-70c0-38.6 31.4-70 70-70s70 31.4 70 70s-31.4 70-70 70h-70zm250 622H546V550h250zm48-376H546V378h298z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GithubFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M511.6 76.3C264.3 76.2 64 276.4 64 523.5C64 718.9 189.3 885 363.8 946c23.5 5.9 19.9-10.8 19.9-22.2v-77.5c-135.7 15.9-141.2-73.9-150.3-88.9C215 726 171.5 718 184.5 703c30.9-15.9 62.4 4 98.9 57.9c26.4 39.1 77.9 32.5 104 26c5.7-23.5 17.9-44.5 34.7-60.8c-140.6-25.2-199.2-111-199.2-213c0-49.5 16.3-95 48.3-131.7c-20.4-60.5 1.9-112.3 4.9-120c58.1-5.2 118.5 41.6 123.2 45.3c33-8.9 70.7-13.6 112.9-13.6c42.4 0 80.2 4.9 113.5 13.9c11.3-8.6 67.3-48.8 121.3-43.9c2.9 7.7 24.7 58.3 5.5 118c32.4 36.8 48.9 82.7 48.9 132.3c0 102.2-59 188.1-200 212.9a127.5 127.5 0 0 1 38.1 91v112.5c.8 9 0 17.9 15 17.9c177.1-59.7 304.6-227 304.6-424.1c0-247.2-200.4-447.3-447.5-447.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GithubOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M511.6 76.3C264.3 76.2 64 276.4 64 523.5C64 718.9 189.3 885 363.8 946c23.5 5.9 19.9-10.8 19.9-22.2v-77.5c-135.7 15.9-141.2-73.9-150.3-88.9C215 726 171.5 718 184.5 703c30.9-15.9 62.4 4 98.9 57.9c26.4 39.1 77.9 32.5 104 26c5.7-23.5 17.9-44.5 34.7-60.8c-140.6-25.2-199.2-111-199.2-213c0-49.5 16.3-95 48.3-131.7c-20.4-60.5 1.9-112.3 4.9-120c58.1-5.2 118.5 41.6 123.2 45.3c33-8.9 70.7-13.6 112.9-13.6c42.4 0 80.2 4.9 113.5 13.9c11.3-8.6 67.3-48.8 121.3-43.9c2.9 7.7 24.7 58.3 5.5 118c32.4 36.8 48.9 82.7 48.9 132.3c0 102.2-59 188.1-200 212.9a127.5 127.5 0 0 1 38.1 91v112.5c.8 9 0 17.9 15 17.9c177.1-59.7 304.6-227 304.6-424.1c0-247.2-200.4-447.3-447.5-447.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitlabFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m910.5 553.2l-109-370.8c-6.8-20.4-23.1-34.1-44.9-34.1s-39.5 12.3-46.3 32.7l-72.2 215.4H386.2L314 181.1c-6.8-20.4-24.5-32.7-46.3-32.7s-39.5 13.6-44.9 34.1L113.9 553.2c-4.1 13.6 1.4 28.6 12.3 36.8l385.4 289l386.7-289c10.8-8.1 16.3-23.1 12.2-36.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitlabOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M913.9 552.2L805 181.4v-.1c-7.6-22.9-25.7-36.5-48.3-36.5c-23.4 0-42.5 13.5-49.7 35.2l-71.4 213H388.8l-71.4-213c-7.2-21.7-26.3-35.2-49.7-35.2c-23.1 0-42.5 14.8-48.4 36.6L110.5 552.2c-4.4 14.7 1.2 31.4 13.5 40.7l368.5 276.4c2.6 3.6 6.2 6.3 10.4 7.8l8.6 6.4l8.5-6.4c4.9-1.7 9-4.7 11.9-8.9l368.4-275.4c12.4-9.2 18-25.9 13.6-40.6M751.7 193.4c1-1.8 2.9-1.9 3.5-1.9c1.1 0 2.5.3 3.4 3L818 394.3H684.5zm-487.4 1c.9-2.6 2.3-2.9 3.4-2.9c2.7 0 2.9.1 3.4 1.7l67.3 201.2H206.5zM158.8 558.7l28.2-97.3l202.4 270.2zm73.9-116.4h122.1l90.8 284.3zM512.9 776L405.7 442.3H620zm157.9-333.7h119.5L580 723.1zm-40.7 293.9l207.3-276.7l29.5 99.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobalOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M854.4 800.9c.2-.3.5-.6.7-.9C920.6 722.1 960 621.7 960 512s-39.4-210.1-104.8-288c-.2-.3-.5-.5-.7-.8c-1.1-1.3-2.1-2.5-3.2-3.7c-.4-.5-.8-.9-1.2-1.4l-4.1-4.7l-.1-.1c-1.5-1.7-3.1-3.4-4.6-5.1l-.1-.1c-3.2-3.4-6.4-6.8-9.7-10.1l-.1-.1l-4.8-4.8l-.3-.3c-1.5-1.5-3-2.9-4.5-4.3c-.5-.5-1-1-1.6-1.5c-1-1-2-1.9-3-2.8c-.3-.3-.7-.6-1-1C736.4 109.2 629.5 64 512 64s-224.4 45.2-304.3 119.2c-.3.3-.7.6-1 1c-1 .9-2 1.9-3 2.9c-.5.5-1 1-1.6 1.5c-1.5 1.4-3 2.9-4.5 4.3l-.3.3l-4.8 4.8l-.1.1c-3.3 3.3-6.5 6.7-9.7 10.1l-.1.1c-1.6 1.7-3.1 3.4-4.6 5.1l-.1.1c-1.4 1.5-2.8 3.1-4.1 4.7c-.4.5-.8.9-1.2 1.4c-1.1 1.2-2.1 2.5-3.2 3.7c-.2.3-.5.5-.7.8C103.4 301.9 64 402.3 64 512s39.4 210.1 104.8 288c.2.3.5.6.7.9l3.1 3.7c.4.5.8.9 1.2 1.4l4.1 4.7c0 .1.1.1.1.2c1.5 1.7 3 3.4 4.6 5l.1.1c3.2 3.4 6.4 6.8 9.6 10.1l.1.1c1.6 1.6 3.1 3.2 4.7 4.7l.3.3c3.3 3.3 6.7 6.5 10.1 9.6c80.1 74 187 119.2 304.5 119.2s224.4-45.2 304.3-119.2a300 300 0 0 0 10-9.6l.3-.3c1.6-1.6 3.2-3.1 4.7-4.7l.1-.1c3.3-3.3 6.5-6.7 9.6-10.1l.1-.1c1.5-1.7 3.1-3.3 4.6-5c0-.1.1-.1.1-.2c1.4-1.5 2.8-3.1 4.1-4.7c.4-.5.8-.9 1.2-1.4a99 99 0 0 0 3.3-3.7m4.1-142.6c-13.8 32.6-32 62.8-54.2 90.2a444.07 444.07 0 0 0-81.5-55.9c11.6-46.9 18.8-98.4 20.7-152.6H887c-3 40.9-12.6 80.6-28.5 118.3M887 484H743.5c-1.9-54.2-9.1-105.7-20.7-152.6c29.3-15.6 56.6-34.4 81.5-55.9A373.86 373.86 0 0 1 887 484M658.3 165.5c39.7 16.8 75.8 40 107.6 69.2a394.72 394.72 0 0 1-59.4 41.8c-15.7-45-35.8-84.1-59.2-115.4c3.7 1.4 7.4 2.9 11 4.4m-90.6 700.6c-9.2 7.2-18.4 12.7-27.7 16.4V697a389.1 389.1 0 0 1 115.7 26.2c-8.3 24.6-17.9 47.3-29 67.8c-17.4 32.4-37.8 58.3-59 75.1m59-633.1c11 20.6 20.7 43.3 29 67.8A389.1 389.1 0 0 1 540 327V141.6c9.2 3.7 18.5 9.1 27.7 16.4c21.2 16.7 41.6 42.6 59 75M540 640.9V540h147.5c-1.6 44.2-7.1 87.1-16.3 127.8l-.3 1.2A445.02 445.02 0 0 0 540 640.9m0-156.9V383.1c45.8-2.8 89.8-12.5 130.9-28.1l.3 1.2c9.2 40.7 14.7 83.5 16.3 127.8zm-56 56v100.9c-45.8 2.8-89.8 12.5-130.9 28.1l-.3-1.2c-9.2-40.7-14.7-83.5-16.3-127.8zm-147.5-56c1.6-44.2 7.1-87.1 16.3-127.8l.3-1.2c41.1 15.6 85 25.3 130.9 28.1V484zM484 697v185.4c-9.2-3.7-18.5-9.1-27.7-16.4c-21.2-16.7-41.7-42.7-59.1-75.1c-11-20.6-20.7-43.3-29-67.8c37.2-14.6 75.9-23.3 115.8-26.1m0-370a389.1 389.1 0 0 1-115.7-26.2c8.3-24.6 17.9-47.3 29-67.8c17.4-32.4 37.8-58.4 59.1-75.1c9.2-7.2 18.4-12.7 27.7-16.4V327zM365.7 165.5c3.7-1.5 7.3-3 11-4.4c-23.4 31.3-43.5 70.4-59.2 115.4c-21-12-40.9-26-59.4-41.8c31.8-29.2 67.9-52.4 107.6-69.2M165.5 365.7c13.8-32.6 32-62.8 54.2-90.2c24.9 21.5 52.2 40.3 81.5 55.9c-11.6 46.9-18.8 98.4-20.7 152.6H137c3-40.9 12.6-80.6 28.5-118.3M137 540h143.5c1.9 54.2 9.1 105.7 20.7 152.6a444.07 444.07 0 0 0-81.5 55.9A373.86 373.86 0 0 1 137 540m228.7 318.5c-39.7-16.8-75.8-40-107.6-69.2c18.5-15.8 38.4-29.7 59.4-41.8c15.7 45 35.8 84.1 59.2 115.4c-3.7-1.4-7.4-2.9-11-4.4m292.6 0c-3.7 1.5-7.3 3-11 4.4c23.4-31.3 43.5-70.4 59.2-115.4c21 12 40.9 26 59.4 41.8a373.81 373.81 0 0 1-107.6 69.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoldFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m905.9 806.7l-40.2-248c-.6-3.9-4-6.7-7.9-6.7H596.2c-3.9 0-7.3 2.8-7.9 6.7l-40.2 248c-.1.4-.1.9-.1 1.3c0 4.4 3.6 8 8 8h342c.4 0 .9 0 1.3-.1c4.3-.7 7.3-4.8 6.6-9.2m-470.2-248c-.6-3.9-4-6.7-7.9-6.7H166.2c-3.9 0-7.3 2.8-7.9 6.7l-40.2 248c-.1.4-.1.9-.1 1.3c0 4.4 3.6 8 8 8h342c.4 0 .9 0 1.3-.1c4.4-.7 7.3-4.8 6.6-9.2zM342 472h342c.4 0 .9 0 1.3-.1c4.4-.7 7.3-4.8 6.6-9.2l-40.2-248c-.6-3.9-4-6.7-7.9-6.7H382.2c-3.9 0-7.3 2.8-7.9 6.7l-40.2 248c-.1.4-.1.9-.1 1.3c0 4.4 3.6 8 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoldOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M342 472h342c.4 0 .9 0 1.3-.1c4.4-.7 7.3-4.8 6.6-9.2l-40.2-248c-.6-3.9-4-6.7-7.9-6.7H382.2c-3.9 0-7.3 2.8-7.9 6.7l-40.2 248c-.1.4-.1.9-.1 1.3c0 4.4 3.6 8 8 8m91.2-196h159.5l20.7 128h-201zm2.5 282.7c-.6-3.9-4-6.7-7.9-6.7H166.2c-3.9 0-7.3 2.8-7.9 6.7l-40.2 248c-.1.4-.1.9-.1 1.3c0 4.4 3.6 8 8 8h342c.4 0 .9 0 1.3-.1c4.4-.7 7.3-4.8 6.6-9.2zM196.5 748l20.7-128h159.5l20.7 128zm709.4 58.7l-40.2-248c-.6-3.9-4-6.7-7.9-6.7H596.2c-3.9 0-7.3 2.8-7.9 6.7l-40.2 248c-.1.4-.1.9-.1 1.3c0 4.4 3.6 8 8 8h342c.4 0 .9 0 1.3-.1c4.3-.7 7.3-4.8 6.6-9.2M626.5 748l20.7-128h159.5l20.7 128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoldTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M435.7 558.7c-.6-3.9-4-6.7-7.9-6.7H166.2c-3.9 0-7.3 2.8-7.9 6.7l-40.2 248c-.1.4-.1.9-.1 1.3c0 4.4 3.6 8 8 8h342c.4 0 .9 0 1.3-.1c4.4-.7 7.3-4.8 6.6-9.2zM196.5 748l20.7-128h159.5l20.7 128zm709.4 58.7l-40.2-248c-.6-3.9-4-6.7-7.9-6.7H596.2c-3.9 0-7.3 2.8-7.9 6.7l-40.2 248c-.1.4-.1.9-.1 1.3c0 4.4 3.6 8 8 8h342c.4 0 .9 0 1.3-.1c4.3-.7 7.3-4.8 6.6-9.2M626.5 748l20.7-128h159.5l20.7 128zM342 472h342c.4 0 .9 0 1.3-.1c4.4-.7 7.3-4.8 6.6-9.2l-40.2-248c-.6-3.9-4-6.7-7.9-6.7H382.2c-3.9 0-7.3 2.8-7.9 6.7l-40.2 248c-.1.4-.1.9-.1 1.3c0 4.4 3.6 8 8 8m91.2-196h159.5l20.7 128h-201z"/><path fill="currentColor" fill-opacity=".15" d="M592.7 276H433.2l-20.8 128h201zM217.2 620l-20.7 128h200.9l-20.7-128zm430 0l-20.7 128h200.9l-20.7-128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoldenFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m905.9 806.7l-40.2-248c-.6-3.9-4-6.7-7.9-6.7H596.2c-3.9 0-7.3 2.8-7.9 6.7l-40.2 248c-.1.4-.1.9-.1 1.3c0 4.4 3.6 8 8 8h342c.4 0 .9 0 1.3-.1c4.3-.7 7.3-4.8 6.6-9.2m-470.2-248c-.6-3.9-4-6.7-7.9-6.7H166.2c-3.9 0-7.3 2.8-7.9 6.7l-40.2 248c-.1.4-.1.9-.1 1.3c0 4.4 3.6 8 8 8h342c.4 0 .9 0 1.3-.1c4.4-.7 7.3-4.8 6.6-9.2zM342 472h342c.4 0 .9 0 1.3-.1c4.4-.7 7.3-4.8 6.6-9.2l-40.2-248c-.6-3.9-4-6.7-7.9-6.7H382.2c-3.9 0-7.3 2.8-7.9 6.7l-40.2 248c-.1.4-.1.9-.1 1.3c0 4.4 3.6 8 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m167 633.6C638.4 735 583 757 516.9 757c-95.7 0-178.5-54.9-218.8-134.9C281.5 589 272 551.6 272 512s9.5-77 26.1-110.1c40.3-80.1 123.1-135 218.8-135c66 0 121.4 24.3 163.9 63.8L610.6 401c-25.4-24.3-57.7-36.6-93.6-36.6c-63.8 0-117.8 43.1-137.1 101c-4.9 14.7-7.7 30.4-7.7 46.6s2.8 31.9 7.7 46.6c19.3 57.9 73.3 101 137 101c33 0 61-8.7 82.9-23.4c26-17.4 43.2-43.3 48.9-74H516.9v-94.8h230.7c2.9 16.1 4.4 32.8 4.4 50.1c0 74.7-26.7 137.4-73 180.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M881 442.4H519.7v148.5h206.4c-8.9 48-35.9 88.6-76.6 115.8c-34.4 23-78.3 36.6-129.9 36.6c-99.9 0-184.4-67.5-214.6-158.2c-7.6-23-12-47.6-12-72.9s4.4-49.9 12-72.9c30.3-90.6 114.8-158.1 214.7-158.1c56.3 0 106.8 19.4 146.6 57.4l110-110.1c-66.5-62-153.2-100-256.6-100c-149.9 0-279.6 86-342.7 211.4c-26 51.8-40.8 110.4-40.8 172.4S151 632.8 177 684.6C240.1 810 369.8 896 519.7 896c103.6 0 190.4-34.4 253.8-93c72.5-66.8 114.4-165.2 114.4-282.1c0-27.2-2.4-53.3-6.9-78.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePlusCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m36.5 558.8c-43.9 61.8-132.1 79.8-200.9 53.3c-69-26.3-118-99.2-112.1-173.5c1.5-90.9 85.2-170.6 176.1-167.5c43.6-2 84.6 16.9 118 43.6c-14.3 16.2-29 31.8-44.8 46.3c-40.1-27.7-97.2-35.6-137.3-3.6c-57.4 39.7-60 133.4-4.8 176.1c53.7 48.7 155.2 24.5 170.1-50.1c-33.6-.5-67.4 0-101-1.1c-.1-20.1-.2-40.1-.1-60.2c56.2-.2 112.5-.3 168.8.2c3.3 47.3-3 97.5-32 136.5M791 536.5c-16.8.2-33.6.3-50.4.4c-.2 16.8-.3 33.6-.3 50.4H690c-.2-16.8-.2-33.5-.3-50.3c-16.8-.2-33.6-.3-50.4-.5v-50.1c16.8-.2 33.6-.3 50.4-.3c.1-16.8.3-33.6.4-50.4h50.2l.3 50.4c16.8.2 33.6.2 50.4.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePlusOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M879.5 470.4c-.3-27-.4-54.2-.5-81.3h-80.8c-.3 27-.5 54.1-.7 81.3c-27.2.1-54.2.3-81.2.6v80.9c27 .3 54.2.5 81.2.8c.3 27 .3 54.1.5 81.1h80.9c.1-27 .3-54.1.5-81.3c27.2-.3 54.2-.4 81.2-.7v-80.9c-26.9-.2-54.1-.2-81.1-.5m-530 .4c-.1 32.3 0 64.7.1 97c54.2 1.8 108.5 1 162.7 1.8c-23.9 120.3-187.4 159.3-273.9 80.7c-89-68.9-84.8-220 7.7-284c64.7-51.6 156.6-38.9 221.3 5.8c25.4-23.5 49.2-48.7 72.1-74.7c-53.8-42.9-119.8-73.5-190-70.3c-146.6-4.9-281.3 123.5-283.7 270.2c-9.4 119.9 69.4 237.4 180.6 279.8c110.8 42.7 252.9 13.6 323.7-86c46.7-62.9 56.8-143.9 51.3-220c-90.7-.7-181.3-.6-271.9-.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePlusSquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M548.5 622.8c-43.9 61.8-132.1 79.8-200.9 53.3c-69-26.3-118-99.2-112.1-173.5c1.5-90.9 85.2-170.6 176.1-167.5c43.6-2 84.6 16.9 118 43.6c-14.3 16.2-29 31.8-44.8 46.3c-40.1-27.7-97.2-35.6-137.3-3.6c-57.4 39.7-60 133.4-4.8 176.1c53.7 48.7 155.2 24.5 170.1-50.1c-33.6-.5-67.4 0-101-1.1c-.1-20.1-.2-40.1-.1-60.2c56.2-.2 112.5-.3 168.8.2c3.3 47.3-3 97.5-32 136.5M791 536.5c-16.8.2-33.6.3-50.4.4c-.2 16.8-.3 33.6-.3 50.4H690c-.2-16.8-.2-33.5-.3-50.3c-16.8-.2-33.6-.3-50.4-.5v-50.1c16.8-.2 33.6-.3 50.4-.3c.1-16.8.3-33.6.4-50.4h50.2l.3 50.4c16.8.2 33.6.2 50.4.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleSquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M679 697.6C638.4 735 583 757 516.9 757c-95.7 0-178.5-54.9-218.8-134.9A245.02 245.02 0 0 1 272 512c0-39.6 9.5-77 26.1-110.1c40.3-80.1 123.1-135 218.8-135c66 0 121.4 24.3 163.9 63.8L610.6 401c-25.4-24.3-57.7-36.6-93.6-36.6c-63.8 0-117.8 43.1-137.1 101c-4.9 14.7-7.7 30.4-7.7 46.6s2.8 31.9 7.7 46.6c19.3 57.9 73.3 101 137 101c33 0 61-8.7 82.9-23.4c26-17.4 43.2-43.3 48.9-74H516.9v-94.8h230.7c2.9 16.1 4.4 32.8 4.4 50.1c0 74.7-26.7 137.4-73 180.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GroupOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M912 820.1V203.9c28-9.9 48-36.6 48-67.9c0-39.8-32.2-72-72-72c-31.3 0-58 20-67.9 48H203.9C194 84 167.3 64 136 64c-39.8 0-72 32.2-72 72c0 31.3 20 58 48 67.9v616.2C84 830 64 856.7 64 888c0 39.8 32.2 72 72 72c31.3 0 58-20 67.9-48h616.2c9.9 28 36.6 48 67.9 48c39.8 0 72-32.2 72-72c0-31.3-20-58-48-67.9M888 112c13.3 0 24 10.7 24 24s-10.7 24-24 24s-24-10.7-24-24s10.7-24 24-24M136 912c-13.3 0-24-10.7-24-24s10.7-24 24-24s24 10.7 24 24s-10.7 24-24 24m0-752c-13.3 0-24-10.7-24-24s10.7-24 24-24s24 10.7 24 24s-10.7 24-24 24m704 680H184V184h656zm48 72c-13.3 0-24-10.7-24-24s10.7-24 24-24s24 10.7 24 24s-10.7 24-24 24"/><path fill="currentColor" d="M288 474h448c8.8 0 16-7.2 16-16V282c0-8.8-7.2-16-16-16H288c-8.8 0-16 7.2-16 16v176c0 8.8 7.2 16 16 16m56-136h336v64H344zm-56 420h448c8.8 0 16-7.2 16-16V566c0-8.8-7.2-16-16-16H288c-8.8 0-16 7.2-16 16v176c0 8.8 7.2 16 16 16m56-136h336v64H344z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HddFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 64H192c-17.7 0-32 14.3-32 32v224h704V96c0-17.7-14.3-32-32-32M456 216c0 4.4-3.6 8-8 8H264c-4.4 0-8-3.6-8-8v-48c0-4.4 3.6-8 8-8h184c4.4 0 8 3.6 8 8zM160 928c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V704H160zm576-136c22.1 0 40 17.9 40 40s-17.9 40-40 40s-40-17.9-40-40s17.9-40 40-40M160 640h704V384H160zm96-152c0-4.4 3.6-8 8-8h184c4.4 0 8 3.6 8 8v48c0 4.4-3.6 8-8 8H264c-4.4 0-8-3.6-8-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HddOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 64H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V96c0-17.7-14.3-32-32-32m-600 72h560v208H232zm560 480H232V408h560zm0 272H232V680h560zM496 208H312c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h184c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8M312 544h184c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8H312c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8m328 244a40 40 0 1 0 80 0a40 40 0 1 0-80 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HddTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M232 888h560V680H232zm448-140c22.1 0 40 17.9 40 40s-17.9 40-40 40s-40-17.9-40-40s17.9-40 40-40M232 616h560V408H232zm72-128c0-4.4 3.6-8 8-8h184c4.4 0 8 3.6 8 8v48c0 4.4-3.6 8-8 8H312c-4.4 0-8-3.6-8-8zm-72-144h560V136H232zm72-128c0-4.4 3.6-8 8-8h184c4.4 0 8 3.6 8 8v48c0 4.4-3.6 8-8 8H312c-4.4 0-8-3.6-8-8z"/><path fill="currentColor" d="M832 64H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V96c0-17.7-14.3-32-32-32m-40 824H232V680h560zm0-272H232V408h560zm0-272H232V136h560z"/><path fill="currentColor" d="M312 544h184c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8H312c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8m0-272h184c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8H312c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8m328 516a40 40 0 1 0 80 0a40 40 0 1 0-80 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M923 283.6a260.04 260.04 0 0 0-56.9-82.8a264.4 264.4 0 0 0-84-55.5A265.34 265.34 0 0 0 679.7 125c-49.3 0-97.4 13.5-139.2 39c-10 6.1-19.5 12.8-28.5 20.1c-9-7.3-18.5-14-28.5-20.1c-41.8-25.5-89.9-39-139.2-39c-35.5 0-69.9 6.8-102.4 20.3c-31.4 13-59.7 31.7-84 55.5a258.44 258.44 0 0 0-56.9 82.8c-13.9 32.3-21 66.6-21 101.9c0 33.3 6.8 68 20.3 103.3c11.3 29.5 27.5 60.1 48.2 91c32.8 48.9 77.9 99.9 133.9 151.6c92.8 85.7 184.7 144.9 188.6 147.3l23.7 15.2c10.5 6.7 24 6.7 34.5 0l23.7-15.2c3.9-2.5 95.7-61.6 188.6-147.3c56-51.7 101.1-102.7 133.9-151.6c20.7-30.9 37-61.5 48.2-91c13.5-35.3 20.3-70 20.3-103.3c.1-35.3-7-69.6-20.9-101.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M923 283.6a260.04 260.04 0 0 0-56.9-82.8a264.4 264.4 0 0 0-84-55.5A265.34 265.34 0 0 0 679.7 125c-49.3 0-97.4 13.5-139.2 39c-10 6.1-19.5 12.8-28.5 20.1c-9-7.3-18.5-14-28.5-20.1c-41.8-25.5-89.9-39-139.2-39c-35.5 0-69.9 6.8-102.4 20.3c-31.4 13-59.7 31.7-84 55.5a258.44 258.44 0 0 0-56.9 82.8c-13.9 32.3-21 66.6-21 101.9c0 33.3 6.8 68 20.3 103.3c11.3 29.5 27.5 60.1 48.2 91c32.8 48.9 77.9 99.9 133.9 151.6c92.8 85.7 184.7 144.9 188.6 147.3l23.7 15.2c10.5 6.7 24 6.7 34.5 0l23.7-15.2c3.9-2.5 95.7-61.6 188.6-147.3c56-51.7 101.1-102.7 133.9-151.6c20.7-30.9 37-61.5 48.2-91c13.5-35.3 20.3-70 20.3-103.3c.1-35.3-7-69.6-20.9-101.9M512 814.8S156 586.7 156 385.5C156 283.6 240.3 201 344.3 201c73.1 0 136.5 40.8 167.7 100.4C543.2 241.8 606.6 201 679.7 201c104 0 188.3 82.6 188.3 184.5c0 201.2-356 429.3-356 429.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M923 283.6a260.04 260.04 0 0 0-56.9-82.8a264.4 264.4 0 0 0-84-55.5A265.34 265.34 0 0 0 679.7 125c-49.3 0-97.4 13.5-139.2 39c-10 6.1-19.5 12.8-28.5 20.1c-9-7.3-18.5-14-28.5-20.1c-41.8-25.5-89.9-39-139.2-39c-35.5 0-69.9 6.8-102.4 20.3c-31.4 13-59.7 31.7-84 55.5a258.44 258.44 0 0 0-56.9 82.8c-13.9 32.3-21 66.6-21 101.9c0 33.3 6.8 68 20.3 103.3c11.3 29.5 27.5 60.1 48.2 91c32.8 48.9 77.9 99.9 133.9 151.6c92.8 85.7 184.7 144.9 188.6 147.3l23.7 15.2c10.5 6.7 24 6.7 34.5 0l23.7-15.2c3.9-2.5 95.7-61.6 188.6-147.3c56-51.7 101.1-102.7 133.9-151.6c20.7-30.9 37-61.5 48.2-91c13.5-35.3 20.3-70 20.3-103.3c.1-35.3-7-69.6-20.9-101.9M512 814.8S156 586.7 156 385.5C156 283.6 240.3 201 344.3 201c73.1 0 136.5 40.8 167.7 100.4C543.2 241.8 606.6 201 679.7 201c104 0 188.3 82.6 188.3 184.5c0 201.2-356 429.3-356 429.3"/><path fill="currentColor" fill-opacity=".15" d="M679.7 201c-73.1 0-136.5 40.8-167.7 100.4C480.8 241.8 417.4 201 344.3 201c-104 0-188.3 82.6-188.3 184.5c0 201.2 356 429.3 356 429.3s356-228.1 356-429.3C868 283.6 783.7 201 679.7 201"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeatMapOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m955.7 856l-416-720c-6.2-10.7-16.9-16-27.7-16s-21.6 5.3-27.7 16l-416 720C56 877.4 71.4 904 96 904h832c24.6 0 40-26.6 27.7-48m-790.4-23.9L512 231.9L858.7 832H165.3zm319-474.1l-228 394c-12.3 21.3 3.1 48 27.7 48h455.8c24.7 0 40.1-26.7 27.7-48L539.7 358c-6.2-10.7-17-16-27.7-16c-10.8 0-21.6 5.3-27.7 16m214 386H325.7L512 422zm-214-194.1l-57 98.4C415 669.5 430.4 696 455 696h114c24.6 0 39.9-26.5 27.7-47.7l-57-98.4c-6.1-10.6-16.9-15.9-27.7-15.9s-21.5 5.3-27.7 15.9m57.1 98.4h-58.7l29.4-50.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HighlightFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M957.6 507.4L603.2 158.2a7.9 7.9 0 0 0-11.2 0L353.3 393.4a8.03 8.03 0 0 0-.1 11.3l.1.1l40 39.4l-117.2 115.3a8.03 8.03 0 0 0-.1 11.3l.1.1l39.5 38.9l-189.1 187H72.1c-4.4 0-8.1 3.6-8.1 8V860c0 4.4 3.6 8 8 8h344.9c2.1 0 4.1-.8 5.6-2.3l76.1-75.6l40.4 39.8a7.9 7.9 0 0 0 11.2 0l117.1-115.6l40.1 39.5a7.9 7.9 0 0 0 11.2 0l238.7-235.2c3.4-3 3.4-8 .3-11.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HighlightOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M957.6 507.4L603.2 158.2a7.9 7.9 0 0 0-11.2 0L353.3 393.4a8.03 8.03 0 0 0-.1 11.3l.1.1l40 39.4l-117.2 115.3a8.03 8.03 0 0 0-.1 11.3l.1.1l39.5 38.9l-189.1 187H72.1c-4.4 0-8.1 3.6-8.1 8V860c0 4.4 3.6 8 8 8h344.9c2.1 0 4.1-.8 5.6-2.3l76.1-75.6l40.4 39.8a7.9 7.9 0 0 0 11.2 0l117.1-115.6l40.1 39.5a7.9 7.9 0 0 0 11.2 0l238.7-235.2c3.4-3 3.4-8 .3-11.2M389.8 796.2H229.6l134.4-133l80.1 78.9zm154.8-62.1L373.2 565.2l68.6-67.6l171.4 168.9zM713.1 658L450.3 399.1L597.6 254l262.8 259z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HighlightTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M229.6 796.3h160.2l54.3-54.1l-80.1-78.9zm220.7-397.1l262.8 258.9l147.3-145l-262.8-259zm-77.1 166.1l171.4 168.9l68.6-67.6l-171.4-168.9z"/><path fill="currentColor" d="M957.6 507.5L603.2 158.3a7.9 7.9 0 0 0-11.2 0L353.3 393.5a8.03 8.03 0 0 0-.1 11.3l.1.1l40 39.4l-117.2 115.3a8.03 8.03 0 0 0-.1 11.3l.1.1l39.5 38.9l-189.1 187H72.1c-4.4 0-8.1 3.6-8.1 8v55.2c0 4.4 3.6 8 8 8h344.9c2.1 0 4.1-.8 5.6-2.3l76.1-75.6L539 830a7.9 7.9 0 0 0 11.2 0l117.1-115.6l40.1 39.5a7.9 7.9 0 0 0 11.2 0l238.7-235.2c3.4-3 3.4-8 .3-11.2M389.8 796.3H229.6l134.4-133l80.1 78.9zm154.8-62.1L373.2 565.3l68.6-67.6l171.4 168.9zm168.5-76.1L450.3 399.2l147.3-145.1l262.8 259z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HistoryOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M536.1 273H488c-4.4 0-8 3.6-8 8v275.3c0 2.6 1.2 5 3.3 6.5l165.3 120.7c3.6 2.6 8.6 1.9 11.2-1.7l28.6-39c2.7-3.7 1.9-8.7-1.7-11.2L544.1 528.5V281c0-4.4-3.6-8-8-8m219.8 75.2l156.8 38.3c5 1.2 9.9-2.6 9.9-7.7l.8-161.5c0-6.7-7.7-10.5-12.9-6.3L752.9 334.1a8 8 0 0 0 3 14.1m167.7 301.1l-56.7-19.5a8 8 0 0 0-10.1 4.8c-1.9 5.1-3.9 10.1-6 15.1c-17.8 42.1-43.3 80-75.9 112.5a353 353 0 0 1-112.5 75.9a352.18 352.18 0 0 1-137.7 27.8c-47.8 0-94.1-9.3-137.7-27.8a353 353 0 0 1-112.5-75.9c-32.5-32.5-58-70.4-75.9-112.5A353.44 353.44 0 0 1 171 512c0-47.8 9.3-94.2 27.8-137.8c17.8-42.1 43.3-80 75.9-112.5a353 353 0 0 1 112.5-75.9C430.6 167.3 477 158 524.8 158s94.1 9.3 137.7 27.8A353 353 0 0 1 775 261.7c10.2 10.3 19.8 21 28.6 32.3l59.8-46.8C784.7 146.6 662.2 81.9 524.6 82C285 82.1 92.6 276.7 95 516.4C97.4 751.9 288.9 942 524.8 942c185.5 0 343.5-117.6 403.7-282.3c1.5-4.2-.7-8.9-4.9-10.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HolderOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M300 276.497a56 56 0 1 0 56-96.994a56 56 0 0 0-56 96.994m0 284a56 56 0 1 0 56-96.994a56 56 0 0 0-56 96.994M640 228a56 56 0 1 0 112 0a56 56 0 0 0-112 0m0 284a56 56 0 1 0 112 0a56 56 0 0 0-112 0M300 844.497a56 56 0 1 0 56-96.994a56 56 0 0 0-56 96.994M640 796a56 56 0 1 0 112 0a56 56 0 0 0-112 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M946.5 505L534.6 93.4a31.93 31.93 0 0 0-45.2 0L77.5 505c-12 12-18.8 28.3-18.8 45.3c0 35.3 28.7 64 64 64h43.4V908c0 17.7 14.3 32 32 32H448V716h112v224h265.9c17.7 0 32-14.3 32-32V614.3h43.4c17 0 33.3-6.7 45.3-18.8c24.9-25 24.9-65.5-.1-90.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M946.5 505L560.1 118.8l-25.9-25.9a31.5 31.5 0 0 0-44.4 0L77.5 505a63.9 63.9 0 0 0-18.8 46c.4 35.2 29.7 63.3 64.9 63.3h42.5V940h691.8V614.3h43.4c17.1 0 33.2-6.7 45.3-18.8a63.6 63.6 0 0 0 18.7-45.3c0-17-6.7-33.1-18.8-45.2M568 868H456V664h112zm217.9-325.7V868H632V640c0-22.1-17.9-40-40-40H432c-22.1 0-40 17.9-40 40v228H238.1V542.3h-96l370-369.7l23.1 23.1L882 542.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="m512.1 172.6l-370 369.7h96V868H392V640c0-22.1 17.9-40 40-40h160c22.1 0 40 17.9 40 40v228h153.9V542.3H882L535.2 195.7zm434.5 422.9c-6 6-13.1 10.8-20.8 13.9c7.7-3.2 14.8-7.9 20.8-13.9m-887-34.7c5 30.3 31.4 53.5 63.1 53.5h.9c-31.9 0-58.9-23-64-53.5m-.9-10.5v-1.9zm.1-2.6c.1-3.1.5-6.1 1-9.1c-.6 2.9-.9 6-1 9.1"/><path fill="currentColor" d="M951 510c0-.1-.1-.1-.1-.2l-1.8-2.1c-.1-.1-.2-.3-.4-.4c-.7-.8-1.5-1.6-2.2-2.4L560.1 118.8l-25.9-25.9a31.5 31.5 0 0 0-44.4 0L77.5 505a63.6 63.6 0 0 0-16 26.6l-.6 2.1l-.3 1.1l-.3 1.2c-.2.7-.3 1.4-.4 2.1c0 .1 0 .3-.1.4c-.6 3-.9 6-1 9.1v3.3c0 .5 0 1 .1 1.5c0 .5 0 .9.1 1.4c0 .5.1 1 .1 1.5c0 .6.1 1.2.2 1.8c0 .3.1.6.1.9l.3 2.5v.1c5.1 30.5 32.2 53.5 64 53.5h42.5V940h691.7V614.3h43.4c8.6 0 16.9-1.7 24.5-4.9s14.7-7.9 20.8-13.9a63.6 63.6 0 0 0 18.7-45.3c0-14.7-5-28.8-14.3-40.2M568 868H456V664h112zm217.9-325.7V868H632V640c0-22.1-17.9-40-40-40H432c-22.1 0-40 17.9-40 40v228H238.1V542.3h-96l370-369.7l23.1 23.1L882 542.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HourglassFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M742 318V184h86c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8H196c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h86v134c0 81.5 42.4 153.2 106.4 194c-64 40.8-106.4 112.5-106.4 194v134h-86c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h632c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-86V706c0-81.5-42.4-153.2-106.4-194c64-40.8 106.4-112.5 106.4-194"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HourglassOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M742 318V184h86c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8H196c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h86v134c0 81.5 42.4 153.2 106.4 194c-64 40.8-106.4 112.5-106.4 194v134h-86c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h632c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-86V706c0-81.5-42.4-153.2-106.4-194c64-40.8 106.4-112.5 106.4-194m-72 388v134H354V706c0-42.2 16.4-81.9 46.3-111.7C430.1 564.4 469.8 548 512 548s81.9 16.4 111.7 46.3C653.6 624.1 670 663.8 670 706m0-388c0 42.2-16.4 81.9-46.3 111.7C593.9 459.6 554.2 476 512 476s-81.9-16.4-111.7-46.3A156.63 156.63 0 0 1 354 318V184h316z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HourglassTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M512 548c-42.2 0-81.9 16.4-111.7 46.3A156.63 156.63 0 0 0 354 706v134h316V706c0-42.2-16.4-81.9-46.3-111.7A156.63 156.63 0 0 0 512 548M354 318c0 42.2 16.4 81.9 46.3 111.7C430.1 459.6 469.8 476 512 476s81.9-16.4 111.7-46.3C653.6 399.9 670 360.2 670 318V184H354z"/><path fill="currentColor" d="M742 318V184h86c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8H196c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h86v134c0 81.5 42.4 153.2 106.4 194c-64 40.8-106.4 112.5-106.4 194v134h-86c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h632c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-86V706c0-81.5-42.4-153.2-106.4-194c64-40.8 106.4-112.5 106.4-194m-72 388v134H354V706c0-42.2 16.4-81.9 46.3-111.7C430.1 564.4 469.8 548 512 548s81.9 16.4 111.7 46.3C653.6 624.1 670 663.8 670 706m0-388c0 42.2-16.4 81.9-46.3 111.7C593.9 459.6 554.2 476 512 476s-81.9-16.4-111.7-46.3A156.63 156.63 0 0 1 354 318V184h316z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HtmlFiveFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m145.2 96l66 746.6L512 928l299.6-85.4L878.9 96zm595 177.1l-4.8 47.2l-1.7 19.5H382.3l8.2 94.2h335.1l-3.3 24.3l-21.2 242.2l-1.7 16.2l-187 51.6v.3h-1.2l-.3.1v-.1h-.1l-188.6-52L310.8 572h91.1l6.5 73.2l102.4 27.7h.4l102-27.6l11.4-118.6H510.9v-.1H306l-22.8-253.5l-1.7-24.3h460.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HtmlFiveOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m145 96l66 746.6L511.8 928l299.6-85.4L878.7 96zm610.9 700.6l-244.1 69.6l-245.2-69.6l-56.7-641.2h603.8zM281 249l1.7 24.3l22.7 253.5h206.5v-.1h112.9l-11.4 118.5L511 672.9v.2h-.8l-102.4-27.7l-6.5-73.2h-91l11.3 144.7l188.6 52h1.7v-.4l187.7-51.7l1.7-16.3l21.2-242.2l3.2-24.3H511v.2H389.9l-8.2-94.2h352.1l1.7-19.5l4.8-47.2L742 249H511z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HtmlFiveTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m145 96l66 746.6L511.8 928l299.6-85.4L878.7 96zm610.9 700.6l-244.1 69.6l-245.2-69.6l-56.7-641.2h603.8z"/><path fill="currentColor" fill-opacity=".15" d="m209.9 155.4l56.7 641.2l245.2 69.6l244.1-69.6l57.8-641.2zm530.4 117.9l-4.8 47.2l-1.7 19.5H381.7l8.2 94.2H511v-.2h214.7l-3.2 24.3l-21.2 242.2l-1.7 16.3l-187.7 51.7v.4h-1.7l-188.6-52l-11.3-144.7h91l6.5 73.2l102.4 27.7h.8v-.2l102.4-27.7l11.4-118.5H511.9v.1H305.4l-22.7-253.5L281 249h461z"/><path fill="currentColor" d="m281 249l1.7 24.3l22.7 253.5h206.5v-.1h112.9l-11.4 118.5L511 672.9v.2h-.8l-102.4-27.7l-6.5-73.2h-91l11.3 144.7l188.6 52h1.7v-.4l187.7-51.7l1.7-16.3l21.2-242.2l3.2-24.3H511v.2H389.9l-8.2-94.2h352.1l1.7-19.5l4.8-47.2L742 249H511z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IdcardFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M373 411c-28.5 0-51.7 23.3-51.7 52s23.2 52 51.7 52s51.7-23.3 51.7-52s-23.2-52-51.7-52m555-251H96c-17.7 0-32 14.3-32 32v640c0 17.7 14.3 32 32 32h832c17.7 0 32-14.3 32-32V192c0-17.7-14.3-32-32-32M608 420c0-4.4 1-8 2.3-8h123.4c1.3 0 2.3 3.6 2.3 8v48c0 4.4-1 8-2.3 8H610.3c-1.3 0-2.3-3.6-2.3-8zm-86 253h-43.9c-4.2 0-7.6-3.3-7.9-7.5c-3.8-50.5-46-90.5-97.2-90.5s-93.4 40-97.2 90.5c-.3 4.2-3.7 7.5-7.9 7.5H224a8 8 0 0 1-8-8.4c2.8-53.3 32-99.7 74.6-126.1a111.8 111.8 0 0 1-29.1-75.5c0-61.9 49.9-112 111.4-112s111.4 50.1 111.4 112c0 29.1-11 55.5-29.1 75.5c42.7 26.5 71.8 72.8 74.6 126.1c.4 4.6-3.2 8.4-7.8 8.4m278.9-53H615.1c-3.9 0-7.1-3.6-7.1-8v-48c0-4.4 3.2-8 7.1-8h185.7c3.9 0 7.1 3.6 7.1 8v48h.1c0 4.4-3.2 8-7.1 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IdcardOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M928 160H96c-17.7 0-32 14.3-32 32v640c0 17.7 14.3 32 32 32h832c17.7 0 32-14.3 32-32V192c0-17.7-14.3-32-32-32m-40 632H136V232h752zM610.3 476h123.4c1.3 0 2.3-3.6 2.3-8v-48c0-4.4-1-8-2.3-8H610.3c-1.3 0-2.3 3.6-2.3 8v48c0 4.4 1 8 2.3 8m4.8 144h185.7c3.9 0 7.1-3.6 7.1-8v-48c0-4.4-3.2-8-7.1-8H615.1c-3.9 0-7.1 3.6-7.1 8v48c0 4.4 3.2 8 7.1 8M224 673h43.9c4.2 0 7.6-3.3 7.9-7.5c3.8-50.5 46-90.5 97.2-90.5s93.4 40 97.2 90.5c.3 4.2 3.7 7.5 7.9 7.5H522a8 8 0 0 0 8-8.4c-2.8-53.3-32-99.7-74.6-126.1a111.8 111.8 0 0 0 29.1-75.5c0-61.9-49.9-112-111.4-112s-111.4 50.1-111.4 112c0 29.1 11 55.5 29.1 75.5a158.09 158.09 0 0 0-74.6 126.1c-.4 4.6 3.2 8.4 7.8 8.4m149-262c28.5 0 51.7 23.3 51.7 52s-23.2 52-51.7 52s-51.7-23.3-51.7-52s23.2-52 51.7-52"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IdcardTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M928 160H96c-17.7 0-32 14.3-32 32v640c0 17.7 14.3 32 32 32h832c17.7 0 32-14.3 32-32V192c0-17.7-14.3-32-32-32m-40 632H136V232h752z"/><path fill="currentColor" fill-opacity=".15" d="M136 792h752V232H136zm472-372c0-4.4 1-8 2.3-8h123.4c1.3 0 2.3 3.6 2.3 8v48c0 4.4-1 8-2.3 8H610.3c-1.3 0-2.3-3.6-2.3-8zm0 144c0-4.4 3.2-8 7.1-8h185.7c3.9 0 7.1 3.6 7.1 8v48c0 4.4-3.2 8-7.1 8H615.1c-3.9 0-7.1-3.6-7.1-8zM216.2 664.6c2.8-53.3 31.9-99.6 74.6-126.1c-18.1-20-29.1-46.4-29.1-75.5c0-61.9 49.9-112 111.4-112s111.4 50.1 111.4 112c0 29.1-11 55.6-29.1 75.5c42.6 26.4 71.8 72.8 74.6 126.1a8 8 0 0 1-8 8.4h-43.9c-4.2 0-7.6-3.3-7.9-7.5c-3.8-50.5-46-90.5-97.2-90.5s-93.4 40-97.2 90.5c-.3 4.2-3.7 7.5-7.9 7.5H224c-4.6 0-8.2-3.8-7.8-8.4"/><path fill="currentColor" fill-opacity=".15" d="M321.3 463a51.7 52 0 1 0 103.4 0a51.7 52 0 1 0-103.4 0"/><path fill="currentColor" d="M610.3 476h123.4c1.3 0 2.3-3.6 2.3-8v-48c0-4.4-1-8-2.3-8H610.3c-1.3 0-2.3 3.6-2.3 8v48c0 4.4 1 8 2.3 8m4.8 144h185.7c3.9 0 7.1-3.6 7.1-8v-48c0-4.4-3.2-8-7.1-8H615.1c-3.9 0-7.1 3.6-7.1 8v48c0 4.4 3.2 8 7.1 8M224 673h43.9c4.2 0 7.6-3.3 7.9-7.5c3.8-50.5 46-90.5 97.2-90.5s93.4 40 97.2 90.5c.3 4.2 3.7 7.5 7.9 7.5H522a8 8 0 0 0 8-8.4c-2.8-53.3-32-99.7-74.6-126.1a111.8 111.8 0 0 0 29.1-75.5c0-61.9-49.9-112-111.4-112s-111.4 50.1-111.4 112c0 29.1 11 55.5 29.1 75.5a158.09 158.09 0 0 0-74.6 126.1c-.4 4.6 3.2 8.4 7.8 8.4m149-262c28.5 0 51.7 23.3 51.7 52s-23.2 52-51.7 52s-51.7-23.3-51.7-52s23.2-52 51.7-52"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IeCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M693.6 284.4c-24 0-51.1 11.7-72.6 22c46.3 18 86 57.3 112.3 99.6c7.1-18.9 14.6-47.9 14.6-67.9c0-32-22.8-53.7-54.3-53.7M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m253.9 492.9H437.1c0 100.4 144.3 136 196.8 47.4h120.8c-32.6 91.7-119.7 146-216.8 146c-35.1 0-70.3-.1-101.7-15.6c-87.4 44.5-180.3 56.6-180.3-42c0-45.8 23.2-107.1 44-145C335 484 381.3 422.8 435.6 374.5c-43.7 18.9-91.1 66.3-122 101.2c25.9-112.8 129.5-193.6 237.1-186.5c130-59.8 209.7-34.1 209.7 38.6c0 27.4-10.6 63.3-21.4 87.9c25.2 45.5 33.3 97.6 26.9 141.2M540.5 399.1c-53.7 0-102 39.7-104 94.9h208c-2-55.1-50.6-94.9-104-94.9M320.6 602.9c-73 152.4 11.5 172.2 100.3 123.3c-46.6-27.5-82.6-72.2-100.3-123.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IeOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M852.6 367.6c16.3-36.9 32.1-90.7 32.1-131.8c0-109.1-119.5-147.6-314.5-57.9c-161.4-10.8-316.8 110.5-355.6 279.7c46.3-52.3 117.4-123.4 183-151.7C316.1 378.3 246.7 470 194 565.6c-31.1 56.9-66 148.8-66 217.5c0 147.9 139.3 129.8 270.4 63c47.1 23.1 99.8 23.4 152.5 23.4c145.7 0 276.4-81.4 325.2-219H694.9c-78.8 132.9-295.2 79.5-295.2-71.2h493.2c9.6-65.4-2.5-143.6-40.3-211.7M224.8 648.3c26.6 76.7 80.6 143.8 150.4 185c-133.1 73.4-259.9 43.6-150.4-185m174-163.3c3-82.7 75.4-142.3 156-142.3c80.1 0 153 59.6 156 142.3zm276.8-281.4c32.1-15.4 72.8-33 108.8-33c47.1 0 81.4 32.6 81.4 80.6c0 30-11.1 73.5-21.9 101.8c-39.3-63.5-98.9-122.4-168.3-149.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IeSquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M765.9 556.9H437.1c0 100.4 144.3 136 196.8 47.4h120.8c-32.6 91.7-119.7 146-216.8 146c-35.1 0-70.3-.1-101.7-15.6c-87.4 44.5-180.3 56.6-180.3-42c0-45.8 23.2-107.1 44-145C335 484 381.3 422.8 435.6 374.5c-43.7 18.9-91.1 66.3-122 101.2c25.9-112.8 129.5-193.6 237.1-186.5c130-59.8 209.7-34.1 209.7 38.6c0 27.4-10.6 63.3-21.4 87.9c25.2 45.5 33.3 97.6 26.9 141.2m-72.3-272.5c-24 0-51.1 11.7-72.6 22c46.3 18 86 57.3 112.3 99.6c7.1-18.9 14.6-47.9 14.6-67.9c0-32-22.8-53.7-54.3-53.7M540.5 399.1c-53.7 0-102 39.7-104 94.9h208c-2-55.1-50.6-94.9-104-94.9M320.6 602.9c-73 152.4 11.5 172.2 100.3 123.3c-46.6-27.5-82.6-72.2-100.3-123.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImportOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M880 912H144c-17.7 0-32-14.3-32-32V144c0-17.7 14.3-32 32-32h360c4.4 0 8 3.6 8 8v56c0 4.4-3.6 8-8 8H184v656h656V520c0-4.4 3.6-8 8-8h56c4.4 0 8 3.6 8 8v360c0 17.7-14.3 32-32 32M653.3 424.6l52.2 52.2c4.7 4.7 1.9 12.8-4.7 13.6l-179.4 21c-5.1.6-9.5-3.7-8.9-8.9l21-179.4c.8-6.6 8.9-9.4 13.6-4.7l52.4 52.4l256.2-256.2c3.1-3.1 8.2-3.1 11.3 0l42.4 42.4c3.1 3.1 3.1 8.2 0 11.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m885.2 446.3l-.2-.8l-112.2-285.1c-5-16.1-19.9-27.2-36.8-27.2H281.2c-17 0-32.1 11.3-36.9 27.6L139.4 443l-.3.7l-.2.8c-1.3 4.9-1.7 9.9-1 14.8c-.1 1.6-.2 3.2-.2 4.8V830a60.9 60.9 0 0 0 60.8 60.8h627.2c33.5 0 60.8-27.3 60.9-60.8V464.1c0-1.3 0-2.6-.1-3.7c.4-4.9 0-9.6-1.3-14.1m-295.8-43l-.3 15.7c-.8 44.9-31.8 75.1-77.1 75.1c-22.1 0-41.1-7.1-54.8-20.6S436 441.2 435.6 419l-.3-15.7H229.5L309 210h399.2l81.7 193.3zm-375 76.8h157.3c24.3 57.1 76 90.8 140.4 90.8c33.7 0 65-9.4 90.3-27.2c22.2-15.6 39.5-37.4 50.7-63.6h156.5V814H214.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m32 664c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8V456c0-4.4 3.6-8 8-8h48c4.4 0 8 3.6 8 8zm-32-344a48.01 48.01 0 0 1 0-96a48.01 48.01 0 0 1 0 96"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoCircleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" d="M464 336a48 48 0 1 0 96 0a48 48 0 1 0-96 0m72 112h-48c-4.4 0-8 3.6-8 8v272c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V456c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoCircleTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372m32 588c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8V456c0-4.4 3.6-8 8-8h48c4.4 0 8 3.6 8 8zm-32-344a48.01 48.01 0 0 1 0-96a48.01 48.01 0 0 1 0 96"/><path fill="currentColor" d="M464 336a48 48 0 1 0 96 0a48 48 0 1 0-96 0m72 112h-48c-4.4 0-8 3.6-8 8v272c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V456c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 224a64 64 0 1 0 128 0a64 64 0 1 0-128 0m96 168h-64c-4.4 0-8 3.6-8 8v464c0 4.4 3.6 8 8 8h64c4.4 0 8-3.6 8-8V400c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InsertRowAboveOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M878.7 336H145.3c-18.4 0-33.3 14.3-33.3 32v464c0 17.7 14.9 32 33.3 32h733.3c18.4 0 33.3-14.3 33.3-32V368c.1-17.7-14.8-32-33.2-32M360 792H184V632h176zm0-224H184V408h176zm240 224H424V632h176zm0-224H424V408h176zm240 224H664V632h176zm0-224H664V408h176zm64-408H120c-4.4 0-8 3.6-8 8v80c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-80c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InsertRowBelowOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M904 768H120c-4.4 0-8 3.6-8 8v80c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-80c0-4.4-3.6-8-8-8m-25.3-608H145.3c-18.4 0-33.3 14.3-33.3 32v464c0 17.7 14.9 32 33.3 32h733.3c18.4 0 33.3-14.3 33.3-32V192c.1-17.7-14.8-32-33.2-32M360 616H184V456h176zm0-224H184V232h176zm240 224H424V456h176zm0-224H424V232h176zm240 224H664V456h176zm0-224H664V232h176z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InsertRowLeftOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M248 112h-80c-4.4 0-8 3.6-8 8v784c0 4.4 3.6 8 8 8h80c4.4 0 8-3.6 8-8V120c0-4.4-3.6-8-8-8m584 0H368c-17.7 0-32 14.9-32 33.3v733.3c0 18.4 14.3 33.3 32 33.3h464c17.7 0 32-14.9 32-33.3V145.3c0-18.4-14.3-33.3-32-33.3M568 840H408V664h160zm0-240H408V424h160zm0-240H408V184h160zm224 480H632V664h160zm0-240H632V424h160zm0-240H632V184h160z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InsertRowRightOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M856 112h-80c-4.4 0-8 3.6-8 8v784c0 4.4 3.6 8 8 8h80c4.4 0 8-3.6 8-8V120c0-4.4-3.6-8-8-8m-200 0H192c-17.7 0-32 14.9-32 33.3v733.3c0 18.4 14.3 33.3 32 33.3h464c17.7 0 32-14.9 32-33.3V145.3c0-18.4-14.3-33.3-32-33.3M392 840H232V664h160zm0-240H232V424h160zm0-240H232V184h160zm224 480H456V664h160zm0-240H456V424h160zm0-240H456V184h160z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InstagramFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 378.7c-73.4 0-133.3 59.9-133.3 133.3S438.6 645.3 512 645.3S645.3 585.4 645.3 512S585.4 378.7 512 378.7M911.8 512c0-55.2.5-109.9-2.6-165c-3.1-64-17.7-120.8-64.5-167.6c-46.9-46.9-103.6-61.4-167.6-64.5c-55.2-3.1-109.9-2.6-165-2.6c-55.2 0-109.9-.5-165 2.6c-64 3.1-120.8 17.7-167.6 64.5C132.6 226.3 118.1 283 115 347c-3.1 55.2-2.6 109.9-2.6 165s-.5 109.9 2.6 165c3.1 64 17.7 120.8 64.5 167.6c46.9 46.9 103.6 61.4 167.6 64.5c55.2 3.1 109.9 2.6 165 2.6c55.2 0 109.9.5 165-2.6c64-3.1 120.8-17.7 167.6-64.5c46.9-46.9 61.4-103.6 64.5-167.6c3.2-55.1 2.6-109.8 2.6-165M512 717.1c-113.5 0-205.1-91.6-205.1-205.1S398.5 306.9 512 306.9S717.1 398.5 717.1 512S625.5 717.1 512 717.1m213.5-370.7c-26.5 0-47.9-21.4-47.9-47.9s21.4-47.9 47.9-47.9s47.9 21.4 47.9 47.9a47.84 47.84 0 0 1-47.9 47.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InstagramOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 306.9c-113.5 0-205.1 91.6-205.1 205.1S398.5 717.1 512 717.1S717.1 625.5 717.1 512S625.5 306.9 512 306.9m0 338.4c-73.4 0-133.3-59.9-133.3-133.3S438.6 378.7 512 378.7S645.3 438.6 645.3 512S585.4 645.3 512 645.3m213.5-394.6c-26.5 0-47.9 21.4-47.9 47.9s21.4 47.9 47.9 47.9s47.9-21.3 47.9-47.9a47.84 47.84 0 0 0-47.9-47.9M911.8 512c0-55.2.5-109.9-2.6-165c-3.1-64-17.7-120.8-64.5-167.6c-46.9-46.9-103.6-61.4-167.6-64.5c-55.2-3.1-109.9-2.6-165-2.6c-55.2 0-109.9-.5-165 2.6c-64 3.1-120.8 17.7-167.6 64.5C132.6 226.3 118.1 283 115 347c-3.1 55.2-2.6 109.9-2.6 165s-.5 109.9 2.6 165c3.1 64 17.7 120.8 64.5 167.6c46.9 46.9 103.6 61.4 167.6 64.5c55.2 3.1 109.9 2.6 165 2.6c55.2 0 109.9.5 165-2.6c64-3.1 120.8-17.7 167.6-64.5c46.9-46.9 61.4-103.6 64.5-167.6c3.2-55.1 2.6-109.8 2.6-165m-88 235.8c-7.3 18.2-16.1 31.8-30.2 45.8c-14.1 14.1-27.6 22.9-45.8 30.2C695.2 844.7 570.3 840 512 840c-58.3 0-183.3 4.7-235.9-16.1c-18.2-7.3-31.8-16.1-45.8-30.2c-14.1-14.1-22.9-27.6-30.2-45.8C179.3 695.2 184 570.3 184 512c0-58.3-4.7-183.3 16.1-235.9c7.3-18.2 16.1-31.8 30.2-45.8s27.6-22.9 45.8-30.2C328.7 179.3 453.7 184 512 184s183.3-4.7 235.9 16.1c18.2 7.3 31.8 16.1 45.8 30.2c14.1 14.1 22.9 27.6 30.2 45.8C844.7 328.7 840 453.7 840 512c0 58.3 4.7 183.2-16.2 235.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InsuranceFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M519.9 358.8h97.9v41.6h-97.9zm347-188.9L527.1 54.1C523 52.7 517.5 52 512 52s-11 .7-15.1 2.1L157.1 169.9c-8.3 2.8-15.1 12.4-15.1 21.2v482.4c0 8.8 5.7 20.4 12.6 25.9L499.3 968c3.5 2.7 8 4.1 12.6 4.1s9.2-1.4 12.6-4.1l344.7-268.6c6.9-5.4 12.6-17 12.6-25.9V191.1c.2-8.8-6.6-18.3-14.9-21.2M411.3 656h-.2c0 4.4-3.6 8-8 8h-37.3c-4.4 0-8-3.6-8-8V471.4c-7.7 9.2-15.4 17.9-23.1 26a6.04 6.04 0 0 1-10.2-2.4l-13.2-43.5c-.6-2-.2-4.1 1.2-5.6c37-43.4 64.7-95.1 82.2-153.6c1.1-3.5 5-5.3 8.4-3.7l38.6 18.3c2.7 1.3 4.1 4.4 3.2 7.2a429.2 429.2 0 0 1-33.6 79zm296.5-49.2l-26.3 35.3a5.92 5.92 0 0 1-8.9.7c-30.6-29.3-56.8-65.2-78.1-106.9V656c0 4.4-3.6 8-8 8h-36.2c-4.4 0-8-3.6-8-8V536c-22 44.7-49 80.8-80.6 107.6a5.9 5.9 0 0 1-8.9-1.4L430 605.7a6 6 0 0 1 1.6-8.1c28.6-20.3 51.9-45.2 71-76h-55.1c-4.4 0-8-3.6-8-8V478c0-4.4 3.6-8 8-8h94.9v-18.6h-65.9c-4.4 0-8-3.6-8-8V316c0-4.4 3.6-8 8-8h184.7c4.4 0 8 3.6 8 8v127.2c0 4.4-3.6 8-8 8h-66.7v18.6h98.8c4.4 0 8 3.6 8 8v35.6c0 4.4-3.6 8-8 8h-59c18.1 29.1 41.8 54.3 72.3 76.9c2.6 2.1 3.2 5.9 1.2 8.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InsuranceOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M441.6 306.8L403 288.6a6.1 6.1 0 0 0-8.4 3.7c-17.5 58.5-45.2 110.1-82.2 153.6a6.05 6.05 0 0 0-1.2 5.6l13.2 43.5c1.3 4.4 7 5.7 10.2 2.4c7.7-8.1 15.4-16.9 23.1-26V656c0 4.4 3.6 8 8 8H403c4.4 0 8-3.6 8-8V393.1a429.2 429.2 0 0 0 33.6-79c1-2.9-.3-6-3-7.3m26.8 9.2v127.2c0 4.4 3.6 8 8 8h65.9v18.6h-94.9c-4.4 0-8 3.6-8 8v35.6c0 4.4 3.6 8 8 8h55.1c-19.1 30.8-42.4 55.7-71 76a6 6 0 0 0-1.6 8.1l22.8 36.5c1.9 3.1 6.2 3.8 8.9 1.4c31.6-26.8 58.7-62.9 80.6-107.6v120c0 4.4 3.6 8 8 8h36.2c4.4 0 8-3.6 8-8V536c21.3 41.7 47.5 77.5 78.1 106.9c2.6 2.5 6.8 2.1 8.9-.7l26.3-35.3c2-2.7 1.4-6.5-1.2-8.4c-30.5-22.6-54.2-47.8-72.3-76.9h59c4.4 0 8-3.6 8-8V478c0-4.4-3.6-8-8-8h-98.8v-18.6h66.7c4.4 0 8-3.6 8-8V316c0-4.4-3.6-8-8-8H476.4c-4.4 0-8 3.6-8 8m51.5 42.8h97.9v41.6h-97.9zm347-188.9L527.1 54.1C523 52.7 517.5 52 512 52s-11 .7-15.1 2.1L157.1 169.9c-8.3 2.8-15.1 12.4-15.1 21.2v482.4c0 8.8 5.7 20.4 12.6 25.9L499.3 968c3.5 2.7 8 4.1 12.6 4.1s9.2-1.4 12.6-4.1l344.7-268.6c6.9-5.4 12.6-17 12.6-25.9V191.1c.2-8.8-6.6-18.3-14.9-21.2M810 654.3L512 886.5L214 654.3V226.7l298-101.6l298 101.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InsuranceTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M866.9 169.9L527.1 54.1C523 52.7 517.5 52 512 52s-11 .7-15.1 2.1L157.1 169.9c-8.3 2.8-15.1 12.4-15.1 21.2v482.4c0 8.8 5.7 20.4 12.6 25.9L499.3 968c3.5 2.7 8 4.1 12.6 4.1s9.2-1.4 12.6-4.1l344.7-268.6c6.9-5.4 12.6-17 12.6-25.9V191.1c.2-8.8-6.6-18.3-14.9-21.2M810 654.3L512 886.5L214 654.3V226.7l298-101.6l298 101.6z"/><path fill="currentColor" fill-opacity=".15" d="M521.9 358.8h97.9v41.6h-97.9z"/><path fill="currentColor" fill-opacity=".15" d="M214 226.7v427.6l298 232.2l298-232.2V226.7L512 125.1zM413.3 656h-.2c0 4.4-3.6 8-8 8h-37.3c-4.4 0-8-3.6-8-8V471.4c-7.7 9.2-15.4 17.9-23.1 26a6.04 6.04 0 0 1-10.2-2.4l-13.2-43.5c-.6-2-.2-4.1 1.2-5.6c37-43.4 64.7-95.1 82.2-153.6c1.1-3.5 5-5.3 8.4-3.7l38.6 18.3c2.7 1.3 4.1 4.4 3.2 7.2a429.2 429.2 0 0 1-33.6 79zm257.9-340v127.2c0 4.4-3.6 8-8 8h-66.7v18.6h98.8c4.4 0 8 3.6 8 8v35.6c0 4.4-3.6 8-8 8h-59c18.1 29.1 41.8 54.3 72.3 76.9c2.6 2.1 3.2 5.9 1.2 8.5l-26.3 35.3a5.92 5.92 0 0 1-8.9.7c-30.6-29.3-56.8-65.2-78.1-106.9V656c0 4.4-3.6 8-8 8h-36.2c-4.4 0-8-3.6-8-8V536c-22 44.7-49 80.8-80.6 107.6a6.38 6.38 0 0 1-4.8 1.4c-1.7-.3-3.2-1.3-4.1-2.8L432 605.7a6 6 0 0 1 1.6-8.1c28.6-20.3 51.9-45.2 71-76h-55.1c-4.4 0-8-3.6-8-8V478c0-4.4 3.6-8 8-8h94.9v-18.6h-65.9c-4.4 0-8-3.6-8-8V316c0-4.4 3.6-8 8-8h184.7c4.4 0 8 3.6 8 8"/><path fill="currentColor" d="m443.7 306.9l-38.6-18.3c-3.4-1.6-7.3.2-8.4 3.7c-17.5 58.5-45.2 110.2-82.2 153.6a5.7 5.7 0 0 0-1.2 5.6l13.2 43.5c1.4 4.5 7 5.8 10.2 2.4c7.7-8.1 15.4-16.8 23.1-26V656c0 4.4 3.6 8 8 8h37.3c4.4 0 8-3.6 8-8h.2V393.1a429.2 429.2 0 0 0 33.6-79c.9-2.8-.5-5.9-3.2-7.2m26.8 9.1v127.4c0 4.4 3.6 8 8 8h65.9V470h-94.9c-4.4 0-8 3.6-8 8v35.6c0 4.4 3.6 8 8 8h55.1c-19.1 30.8-42.4 55.7-71 76a6 6 0 0 0-1.6 8.1l22.8 36.5c.9 1.5 2.4 2.5 4.1 2.8c1.7.3 3.5-.2 4.8-1.4c31.6-26.8 58.6-62.9 80.6-107.6v120c0 4.4 3.6 8 8 8h36.2c4.4 0 8-3.6 8-8V535.9c21.3 41.7 47.5 77.6 78.1 106.9c2.6 2.5 6.7 2.2 8.9-.7l26.3-35.3c2-2.6 1.4-6.4-1.2-8.5c-30.5-22.6-54.2-47.8-72.3-76.9h59c4.4 0 8-3.6 8-8v-35.6c0-4.4-3.6-8-8-8h-98.8v-18.6h66.7c4.4 0 8-3.6 8-8V316c0-4.4-3.6-8-8-8H478.5c-4.4 0-8 3.6-8 8m51.4 42.8h97.9v41.6h-97.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InteractionFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M726 585.7c0 55.3-44.7 100.1-99.7 100.1H420.6v53.4c0 5.7-6.5 8.8-10.9 5.3l-109.1-85.7c-3.5-2.7-3.5-8 0-10.7l109.1-85.7c4.4-3.5 10.9-.3 10.9 5.3v53.4h205.7c19.6 0 35.5-16 35.5-35.6v-78.9c0-3.7 3-6.8 6.8-6.8h50.7c3.7 0 6.8 3 6.8 6.8v79.1zm-2.6-209.9l-109.1 85.7c-4.4 3.5-10.9.3-10.9-5.3v-53.4H397.7c-19.6 0-35.5 16-35.5 35.6v78.9c0 3.7-3 6.8-6.8 6.8h-50.7c-3.7 0-6.8-3-6.8-6.8v-78.9c0-55.3 44.7-100.1 99.7-100.1h205.7v-53.4c0-5.7 6.5-8.8 10.9-5.3l109.1 85.7c3.6 2.5 3.6 7.8.1 10.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InteractionOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656zM304.8 524h50.7c3.7 0 6.8-3 6.8-6.8v-78.9c0-19.7 15.9-35.6 35.5-35.6h205.7v53.4c0 5.7 6.5 8.8 10.9 5.3l109.1-85.7c3.5-2.7 3.5-8 0-10.7l-109.1-85.7c-4.4-3.5-10.9-.3-10.9 5.3V338H397.7c-55.1 0-99.7 44.8-99.7 100.1V517c0 4 3 7 6.8 7m-4.2 134.9l109.1 85.7c4.4 3.5 10.9.3 10.9-5.3v-53.4h205.7c55.1 0 99.7-44.8 99.7-100.1v-78.9c0-3.7-3-6.8-6.8-6.8h-50.7c-3.7 0-6.8 3-6.8 6.8v78.9c0 19.7-15.9 35.6-35.5 35.6H420.6V568c0-5.7-6.5-8.8-10.9-5.3l-109.1 85.7c-3.5 2.5-3.5 7.8 0 10.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InteractionTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/><path fill="currentColor" fill-opacity=".15" d="M184 840h656V184H184zm114-401.9c0-55.3 44.6-100.1 99.7-100.1h205.8v-53.4c0-5.6 6.5-8.8 10.9-5.3L723.5 365c3.5 2.7 3.5 8 0 10.7l-109.1 85.7c-4.4 3.5-10.9.4-10.9-5.3v-53.4H397.8c-19.6 0-35.5 15.9-35.5 35.6v78.9c0 3.8-3.1 6.8-6.8 6.8h-50.7c-3.8 0-6.8-3-6.8-7zm2.6 210.3l109.1-85.7c4.4-3.5 10.9-.4 10.9 5.3v53.4h205.6c19.6 0 35.5-15.9 35.5-35.6v-78.9c0-3.8 3.1-6.8 6.8-6.8h50.7c3.8 0 6.8 3.1 6.8 6.8v78.9c0 55.3-44.6 100.1-99.7 100.1H420.6v53.4c0 5.6-6.5 8.8-10.9 5.3l-109.1-85.7c-3.5-2.7-3.5-8 0-10.5"/><path fill="currentColor" d="M304.8 524h50.7c3.7 0 6.8-3 6.8-6.8v-78.9c0-19.7 15.9-35.6 35.5-35.6h205.7v53.4c0 5.7 6.5 8.8 10.9 5.3l109.1-85.7c3.5-2.7 3.5-8 0-10.7l-109.1-85.7c-4.4-3.5-10.9-.3-10.9 5.3V338H397.7c-55.1 0-99.7 44.8-99.7 100.1V517c0 4 3 7 6.8 7m-4.2 134.9l109.1 85.7c4.4 3.5 10.9.3 10.9-5.3v-53.4h205.7c55.1 0 99.7-44.8 99.7-100.1v-78.9c0-3.7-3-6.8-6.8-6.8h-50.7c-3.7 0-6.8 3-6.8 6.8v78.9c0 19.7-15.9 35.6-35.5 35.6H420.6V568c0-5.7-6.5-8.8-10.9-5.3l-109.1 85.7c-3.5 2.5-3.5 7.8 0 10.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IssuesCloseOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 688a48 48 0 1 0 96 0a48 48 0 1 0-96 0m72-112c4.4 0 8-3.6 8-8V296c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v272c0 4.4 3.6 8 8 8zm400-188h-59.3c-2.6 0-5 1.2-6.5 3.3L763.7 538.1l-49.9-68.8a7.92 7.92 0 0 0-6.5-3.3H648c-6.5 0-10.3 7.4-6.5 12.7l109.2 150.7a16.1 16.1 0 0 0 26 0l165.8-228.7c3.8-5.3 0-12.7-6.5-12.7m-44 306h-64.2c-5.5 0-10.6 2.9-13.6 7.5a352.2 352.2 0 0 1-49.8 62.2A355.92 355.92 0 0 1 651.1 840a355 355 0 0 1-138.7 27.9c-48.1 0-94.8-9.4-138.7-27.9a355.92 355.92 0 0 1-113.3-76.3A353.06 353.06 0 0 1 184 650.5c-18.6-43.8-28-90.5-28-138.5s9.4-94.7 28-138.5c17.9-42.4 43.6-80.5 76.4-113.2c32.8-32.7 70.9-58.4 113.3-76.3a355 355 0 0 1 138.7-27.9c48.1 0 94.8 9.4 138.7 27.9c42.4 17.9 80.5 43.6 113.3 76.3c19 19 35.6 39.8 49.8 62.2c2.9 4.7 8.1 7.5 13.6 7.5H892c6 0 9.8-6.3 7.2-11.6C828.8 178.5 684.7 82 517.7 80C278.9 77.2 80.5 272.5 80 511.2C79.5 750.1 273.3 944 512.4 944c169.2 0 315.6-97 386.7-238.4A8 8 0 0 0 892 694"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ItalicOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M798 160H366c-4.4 0-8 3.6-8 8v64c0 4.4 3.6 8 8 8h181.2l-156 544H229c-4.4 0-8 3.6-8 8v64c0 4.4 3.6 8 8 8h432c4.4 0 8-3.6 8-8v-64c0-4.4-3.6-8-8-8H474.4l156-544H798c4.4 0 8-3.6 8-8v-64c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M608 112c-167.9 0-304 136.1-304 304c0 70.3 23.9 135 63.9 186.5l-41.1 41.1l-62.3-62.3a8.15 8.15 0 0 0-11.4 0l-39.8 39.8a8.15 8.15 0 0 0 0 11.4l62.3 62.3l-44.9 44.9l-62.3-62.3a8.15 8.15 0 0 0-11.4 0l-39.8 39.8a8.15 8.15 0 0 0 0 11.4l62.3 62.3l-65.3 65.3a8.03 8.03 0 0 0 0 11.3l42.3 42.3c3.1 3.1 8.2 3.1 11.3 0l253.6-253.6A304.06 304.06 0 0 0 608 720c167.9 0 304-136.1 304-304S775.9 112 608 112m161.2 465.2C726.2 620.3 668.9 644 608 644c-60.9 0-118.2-23.7-161.2-66.8c-43.1-43-66.8-100.3-66.8-161.2c0-60.9 23.7-118.2 66.8-161.2c43-43.1 100.3-66.8 161.2-66.8c60.9 0 118.2 23.7 161.2 66.8c43.1 43 66.8 100.3 66.8 161.2c0 60.9-23.7 118.2-66.8 161.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LaptopOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M956.9 845.1L896.4 632V168c0-17.7-14.3-32-32-32h-704c-17.7 0-32 14.3-32 32v464L67.9 845.1C60.4 866 75.8 888 98 888h828.8c22.2 0 37.6-22 30.1-42.9M200.4 208h624v395h-624zm228.3 608l8.1-37h150.3l8.1 37zm224 0l-19.1-86.7c-.8-3.7-4.1-6.3-7.8-6.3H398.2c-3.8 0-7 2.6-7.8 6.3L371.3 816H151l42.3-149h638.2l42.3 149z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 912h496c17.7 0 32-14.3 32-32V340H384zm496-800H384v164h528V144c0-17.7-14.3-32-32-32m-768 32v736c0 17.7 14.3 32 32 32h176V112H144c-17.7 0-32 14.3-32 32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-696 72h136v656H184zm656 656H384V384h456zM384 320V184h456v136z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M384 185h456v136H384zm-200 0h136v656H184zm696-73H144c-17.7 0-32 14.3-32 32v1c0-17.7 14.3-32 32-32h736c17.7 0 32 14.3 32 32v-1c0-17.7-14.3-32-32-32M384 385h456v456H384z"/><path fill="currentColor" d="M880 113H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V145c0-17.7-14.3-32-32-32M320 841H184V185h136zm520 0H384V385h456zm0-520H384V185h456z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m104 316.9c0 10.2-4.9 19.9-13.2 25.9L457.4 512l145.4 105.2c8.3 6 13.2 15.6 13.2 25.9V690c0 6.5-7.4 10.3-12.7 6.5l-246-178a7.95 7.95 0 0 1 0-12.9l246-178a8 8 0 0 1 12.7 6.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftCircleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m603.3 327.5l-246 178a7.95 7.95 0 0 0 0 12.9l246 178c5.3 3.8 12.7 0 12.7-6.5V643c0-10.2-4.9-19.9-13.2-25.9L457.4 512l145.4-105.2c8.3-6 13.2-15.6 13.2-25.9V334c0-6.5-7.4-10.3-12.7-6.5"/><path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftCircleTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372m104 240.9c0 10.3-4.9 19.9-13.2 25.9L457.4 512l145.4 105.1c8.3 6 13.2 15.7 13.2 25.9v46.9c0 6.5-7.4 10.3-12.7 6.5l-246-178a7.95 7.95 0 0 1 0-12.9l246-178c5.3-3.8 12.7 0 12.7 6.5z"/><path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" d="m603.3 327.5l-246 178a7.95 7.95 0 0 0 0 12.9l246 178c5.3 3.8 12.7 0 12.7-6.5V643c0-10.2-4.9-19.9-13.2-25.9L457.4 512l145.4-105.2c8.3-6 13.2-15.6 13.2-25.9V334c0-6.5-7.4-10.3-12.7-6.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M724 218.3V141c0-6.7-7.7-10.4-12.9-6.3L260.3 486.8a31.86 31.86 0 0 0 0 50.3l450.8 352.1c5.3 4.1 12.9.4 12.9-6.3v-77.3c0-4.9-2.3-9.6-6.1-12.6l-360-281l360-281.1c3.8-3 6.1-7.7 6.1-12.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftSquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M624 380.9c0 10.2-4.9 19.9-13.2 25.9L465.4 512l145.4 105.2c8.3 6 13.2 15.6 13.2 25.9V690c0 6.5-7.4 10.3-12.7 6.5l-246-178a7.95 7.95 0 0 1 0-12.9l246-178c5.3-3.8 12.7 0 12.7 6.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftSquareOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m365.3 518.5l246 178c5.3 3.8 12.7 0 12.7-6.5v-46.9c0-10.2-4.9-19.9-13.2-25.9L465.4 512l145.4-105.2c8.3-6 13.2-15.6 13.2-25.9V334c0-6.5-7.4-10.3-12.7-6.5l-246 178a8.05 8.05 0 0 0 0 13"/><path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftSquareTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/><path fill="currentColor" fill-opacity=".15" d="M184 840h656V184H184zm181.3-334.5l246-178c5.3-3.8 12.7 0 12.7 6.5v46.9c0 10.3-4.9 19.9-13.2 25.9L465.4 512l145.4 105.2c8.3 6 13.2 15.7 13.2 25.9V690c0 6.5-7.4 10.3-12.7 6.4l-246-178a7.95 7.95 0 0 1 0-12.9"/><path fill="currentColor" d="m365.3 518.4l246 178c5.3 3.9 12.7.1 12.7-6.4v-46.9c0-10.2-4.9-19.9-13.2-25.9L465.4 512l145.4-105.2c8.3-6 13.2-15.6 13.2-25.9V334c0-6.5-7.4-10.3-12.7-6.5l-246 178a7.95 7.95 0 0 0 0 12.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LikeFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M885.9 533.7c16.8-22.2 26.1-49.4 26.1-77.7c0-44.9-25.1-87.4-65.5-111.1a67.67 67.67 0 0 0-34.3-9.3H572.4l6-122.9c1.4-29.7-9.1-57.9-29.5-79.4A106.62 106.62 0 0 0 471 99.9c-52 0-98 35-111.8 85.1l-85.9 311h-.3v428h472.3c9.2 0 18.2-1.8 26.5-5.4c47.6-20.3 78.3-66.8 78.3-118.4c0-12.6-1.8-25-5.4-37c16.8-22.2 26.1-49.4 26.1-77.7c0-12.6-1.8-25-5.4-37c16.8-22.2 26.1-49.4 26.1-77.7c-.2-12.6-2-25.1-5.6-37.1M112 528v364c0 17.7 14.3 32 32 32h65V496h-65c-17.7 0-32 14.3-32 32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LikeOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M885.9 533.7c16.8-22.2 26.1-49.4 26.1-77.7c0-44.9-25.1-87.4-65.5-111.1a67.67 67.67 0 0 0-34.3-9.3H572.4l6-122.9c1.4-29.7-9.1-57.9-29.5-79.4A106.62 106.62 0 0 0 471 99.9c-52 0-98 35-111.8 85.1l-85.9 311H144c-17.7 0-32 14.3-32 32v364c0 17.7 14.3 32 32 32h601.3c9.2 0 18.2-1.8 26.5-5.4c47.6-20.3 78.3-66.8 78.3-118.4c0-12.6-1.8-25-5.4-37c16.8-22.2 26.1-49.4 26.1-77.7c0-12.6-1.8-25-5.4-37c16.8-22.2 26.1-49.4 26.1-77.7c-.2-12.6-2-25.1-5.6-37.1M184 852V568h81v284zm636.4-353l-21.9 19l13.9 25.4a56.2 56.2 0 0 1 6.9 27.3c0 16.5-7.2 32.2-19.6 43l-21.9 19l13.9 25.4a56.2 56.2 0 0 1 6.9 27.3c0 16.5-7.2 32.2-19.6 43l-21.9 19l13.9 25.4a56.2 56.2 0 0 1 6.9 27.3c0 22.4-13.2 42.6-33.6 51.8H329V564.8l99.5-360.5a44.1 44.1 0 0 1 42.2-32.3c7.6 0 15.1 2.2 21.1 6.7c9.9 7.4 15.2 18.6 14.6 30.5l-9.6 198.4h314.4C829 418.5 840 436.9 840 456c0 16.5-7.2 32.1-19.6 43"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LikeTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M273 495.9v428l.3-428zm538.2-88.3H496.8l9.6-198.4c.6-11.9-4.7-23.1-14.6-30.5c-6.1-4.5-13.6-6.8-21.1-6.7c-19.6.1-36.9 13.4-42.2 32.3c-37.1 134.4-64.9 235.2-83.5 302.5V852h399.4a56.85 56.85 0 0 0 33.6-51.8c0-9.7-2.3-18.9-6.9-27.3l-13.9-25.4l21.9-19a56.76 56.76 0 0 0 19.6-43c0-9.7-2.3-18.9-6.9-27.3l-13.9-25.4l21.9-19a56.76 56.76 0 0 0 19.6-43c0-9.7-2.3-18.9-6.9-27.3l-14-25.5l21.9-19a56.76 56.76 0 0 0 19.6-43c0-19.1-11-37.5-28.8-48.4"/><path fill="currentColor" d="M112 528v364c0 17.7 14.3 32 32 32h65V496h-65c-17.7 0-32 14.3-32 32m773.9 5.7c16.8-22.2 26.1-49.4 26.1-77.7c0-44.9-25.1-87.5-65.5-111a67.67 67.67 0 0 0-34.3-9.3H572.3l6-122.9c1.5-29.7-9-57.9-29.5-79.4a106.4 106.4 0 0 0-77.9-33.4c-52 0-98 35-111.8 85.1l-85.8 310.8l-.3 428h472.1c9.3 0 18.2-1.8 26.5-5.4c47.6-20.3 78.3-66.8 78.3-118.4c0-12.6-1.8-25-5.4-37c16.8-22.2 26.1-49.4 26.1-77.7c0-12.6-1.8-25-5.4-37c16.8-22.2 26.1-49.4 26.1-77.7c0-12.6-1.8-25-5.4-37M820.4 499l-21.9 19l14 25.5a56.2 56.2 0 0 1 6.9 27.3c0 16.5-7.1 32.2-19.6 43l-21.9 19l13.9 25.4a56.2 56.2 0 0 1 6.9 27.3c0 16.5-7.1 32.2-19.6 43l-21.9 19l13.9 25.4a56.2 56.2 0 0 1 6.9 27.3c0 22.4-13.2 42.6-33.6 51.8H345V506.8c18.6-67.2 46.4-168 83.5-302.5a44.28 44.28 0 0 1 42.2-32.3c7.5-.1 15 2.2 21.1 6.7c9.9 7.4 15.2 18.6 14.6 30.5l-9.6 198.4h314.4C829 418.5 840 436.9 840 456c0 16.5-7.1 32.2-19.6 43"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineChartOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M888 792H200V168c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v688c0 4.4 3.6 8 8 8h752c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8M305.8 637.7c3.1 3.1 8.1 3.1 11.3 0l138.3-137.6L583 628.5c3.1 3.1 8.2 3.1 11.3 0l275.4-275.3c3.1-3.1 3.1-8.2 0-11.3l-39.6-39.6a8.03 8.03 0 0 0-11.3 0l-230 229.9L461.4 404a8.03 8.03 0 0 0-11.3 0L266.3 586.7a8.03 8.03 0 0 0 0 11.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineHeightOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M648 160H104c-4.4 0-8 3.6-8 8v128c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-64h168v560h-92c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h264c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-92V232h168v64c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V168c0-4.4-3.6-8-8-8m272.8 546H856V318h64.8c6 0 9.4-7 5.7-11.7L825.7 178.7a7.14 7.14 0 0 0-11.3 0L713.6 306.3a7.23 7.23 0 0 0 5.7 11.7H784v388h-64.8c-6 0-9.4 7-5.7 11.7l100.8 127.5c2.9 3.7 8.5 3.7 11.3 0l100.8-127.5a7.2 7.2 0 0 0-5.6-11.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M904 476H120c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M574 665.4a8.03 8.03 0 0 0-11.3 0L446.5 781.6c-53.8 53.8-144.6 59.5-204 0c-59.5-59.5-53.8-150.2 0-204l116.2-116.2c3.1-3.1 3.1-8.2 0-11.3l-39.8-39.8a8.03 8.03 0 0 0-11.3 0L191.4 526.5c-84.6 84.6-84.6 221.5 0 306s221.5 84.6 306 0l116.2-116.2c3.1-3.1 3.1-8.2 0-11.3zm258.6-474c-84.6-84.6-221.5-84.6-306 0L410.3 307.6a8.03 8.03 0 0 0 0 11.3l39.7 39.7c3.1 3.1 8.2 3.1 11.3 0l116.2-116.2c53.8-53.8 144.6-59.5 204 0c59.5 59.5 53.8 150.2 0 204L665.3 562.6a8.03 8.03 0 0 0 0 11.3l39.8 39.8c3.1 3.1 8.2 3.1 11.3 0l116.2-116.2c84.5-84.6 84.5-221.5 0-306.1M610.1 372.3a8.03 8.03 0 0 0-11.3 0L372.3 598.7a8.03 8.03 0 0 0 0 11.3l39.6 39.6c3.1 3.1 8.2 3.1 11.3 0l226.4-226.4c3.1-3.1 3.1-8.2 0-11.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkedinFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M349.3 793.7H230.6V411.9h118.7zm-59.3-434a68.8 68.8 0 1 1 68.8-68.8c-.1 38-30.9 68.8-68.8 68.8m503.7 434H675.1V608c0-44.3-.8-101.2-61.7-101.2c-61.7 0-71.2 48.2-71.2 98v188.9H423.7V411.9h113.8v52.2h1.6c15.8-30 54.5-61.7 112.3-61.7c120.2 0 142.3 79.1 142.3 181.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkedinOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M847.7 112H176.3c-35.5 0-64.3 28.8-64.3 64.3v671.4c0 35.5 28.8 64.3 64.3 64.3h671.4c35.5 0 64.3-28.8 64.3-64.3V176.3c0-35.5-28.8-64.3-64.3-64.3m0 736c-447.8-.1-671.7-.2-671.7-.3c.1-447.8.2-671.7.3-671.7c447.8.1 671.7.2 671.7.3c-.1 447.8-.2 671.7-.3 671.7M230.6 411.9h118.7v381.8H230.6zm59.4-52.2c37.9 0 68.8-30.8 68.8-68.8a68.8 68.8 0 1 0-137.6 0c-.1 38 30.7 68.8 68.8 68.8m252.3 245.1c0-49.8 9.5-98 71.2-98c60.8 0 61.7 56.9 61.7 101.2v185.7h118.6V584.3c0-102.8-22.2-181.9-142.3-181.9c-57.7 0-96.4 31.7-112.3 61.7h-1.6v-52.2H423.7v381.8h118.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LoadingOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M988 548c-19.9 0-36-16.1-36-36c0-59.4-11.6-117-34.6-171.3a440.45 440.45 0 0 0-94.3-139.9a437.71 437.71 0 0 0-139.9-94.3C629 83.6 571.4 72 512 72c-19.9 0-36-16.1-36-36s16.1-36 36-36c69.1 0 136.2 13.5 199.3 40.3C772.3 66 827 103 874 150c47 47 83.9 101.8 109.7 162.7c26.7 63.1 40.2 130.2 40.2 199.3c.1 19.9-16 36-35.9 36"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LoadingThreeQuartersOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 1024c-69.1 0-136.2-13.5-199.3-40.2C251.7 958 197 921 150 874c-47-47-84-101.7-109.8-162.7C13.5 648.2 0 581.1 0 512c0-19.9 16.1-36 36-36s36 16.1 36 36c0 59.4 11.6 117 34.6 171.3c22.2 52.4 53.9 99.5 94.3 139.9c40.4 40.4 87.5 72.2 139.9 94.3C395 940.4 452.6 952 512 952c59.4 0 117-11.6 171.3-34.6c52.4-22.2 99.5-53.9 139.9-94.3c40.4-40.4 72.2-87.5 94.3-139.9C940.4 629 952 571.4 952 512c0-59.4-11.6-117-34.6-171.3a440.45 440.45 0 0 0-94.3-139.9a437.71 437.71 0 0 0-139.9-94.3C629 83.6 571.4 72 512 72c-19.9 0-36-16.1-36-36s16.1-36 36-36c69.1 0 136.2 13.5 199.3 40.2C772.3 66 827 103 874 150c47 47 83.9 101.8 109.7 162.7c26.7 63.1 40.2 130.2 40.2 199.3s-13.5 136.2-40.2 199.3C958 772.3 921 827 874 874c-47 47-101.8 83.9-162.7 109.7c-63.1 26.8-130.2 40.3-199.3 40.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 464h-68V240c0-70.7-57.3-128-128-128H388c-70.7 0-128 57.3-128 128v224h-68c-17.7 0-32 14.3-32 32v384c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V496c0-17.7-14.3-32-32-32M540 701v53c0 4.4-3.6 8-8 8h-40c-4.4 0-8-3.6-8-8v-53a48.01 48.01 0 1 1 56 0m152-237H332V240c0-30.9 25.1-56 56-56h248c30.9 0 56 25.1 56 56z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 464h-68V240c0-70.7-57.3-128-128-128H388c-70.7 0-128 57.3-128 128v224h-68c-17.7 0-32 14.3-32 32v384c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V496c0-17.7-14.3-32-32-32M332 240c0-30.9 25.1-56 56-56h248c30.9 0 56 25.1 56 56v224H332zm460 600H232V536h560zM484 701v53c0 4.4 3.6 8 8 8h40c4.4 0 8-3.6 8-8v-53a48.01 48.01 0 1 0-56 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 464h-68V240c0-70.7-57.3-128-128-128H388c-70.7 0-128 57.3-128 128v224h-68c-17.7 0-32 14.3-32 32v384c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V496c0-17.7-14.3-32-32-32M332 240c0-30.9 25.1-56 56-56h248c30.9 0 56 25.1 56 56v224H332zm460 600H232V536h560z"/><path fill="currentColor" fill-opacity=".15" d="M232 840h560V536H232zm280-226a48.01 48.01 0 0 1 28 87v53c0 4.4-3.6 8-8 8h-40c-4.4 0-8-3.6-8-8v-53a48.01 48.01 0 0 1 28-87"/><path fill="currentColor" d="M484 701v53c0 4.4 3.6 8 8 8h40c4.4 0 8-3.6 8-8v-53a48.01 48.01 0 1 0-56 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LoginOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M521.7 82c-152.5-.4-286.7 78.5-363.4 197.7c-3.4 5.3.4 12.3 6.7 12.3h70.3c4.8 0 9.3-2.1 12.3-5.8c7-8.5 14.5-16.7 22.4-24.5c32.6-32.5 70.5-58.1 112.7-75.9c43.6-18.4 90-27.8 137.9-27.8c47.9 0 94.3 9.3 137.9 27.8c42.2 17.8 80.1 43.4 112.7 75.9c32.6 32.5 58.1 70.4 76 112.5C865.7 417.8 875 464.1 875 512c0 47.9-9.4 94.2-27.8 137.8c-17.8 42.1-43.4 80-76 112.5s-70.5 58.1-112.7 75.9A352.8 352.8 0 0 1 520.6 866c-47.9 0-94.3-9.4-137.9-27.8A353.84 353.84 0 0 1 270 762.3c-7.9-7.9-15.3-16.1-22.4-24.5c-3-3.7-7.6-5.8-12.3-5.8H165c-6.3 0-10.2 7-6.7 12.3C234.9 863.2 368.5 942 520.6 942c236.2 0 428-190.1 430.4-425.6C953.4 277.1 761.3 82.6 521.7 82M395.02 624v-76h-314c-4.4 0-8-3.6-8-8v-56c0-4.4 3.6-8 8-8h314v-76c0-6.7 7.8-10.5 13-6.3l141.9 112a8 8 0 0 1 0 12.6l-141.9 112c-5.2 4.1-13 .4-13-6.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoutOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M868 732h-70.3c-4.8 0-9.3 2.1-12.3 5.8c-7 8.5-14.5 16.7-22.4 24.5a353.84 353.84 0 0 1-112.7 75.9A352.8 352.8 0 0 1 512.4 866c-47.9 0-94.3-9.4-137.9-27.8a353.84 353.84 0 0 1-112.7-75.9a353.28 353.28 0 0 1-76-112.5C167.3 606.2 158 559.9 158 512s9.4-94.2 27.8-137.8c17.8-42.1 43.4-80 76-112.5s70.5-58.1 112.7-75.9c43.6-18.4 90-27.8 137.9-27.8c47.9 0 94.3 9.3 137.9 27.8c42.2 17.8 80.1 43.4 112.7 75.9c7.9 7.9 15.3 16.1 22.4 24.5c3 3.7 7.6 5.8 12.3 5.8H868c6.3 0 10.2-7 6.7-12.3C798 160.5 663.8 81.6 511.3 82C271.7 82.6 79.6 277.1 82 516.4C84.4 751.9 276.2 942 512.4 942c152.1 0 285.7-78.8 362.3-197.7c3.4-5.3-.4-12.3-6.7-12.3m88.9-226.3L815 393.7c-5.3-4.2-13-.4-13 6.3v76H488c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h314v76c0 6.7 7.8 10.5 13 6.3l141.9-112a8 8 0 0 0 0-12.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MacCommandFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M624 672c0 26.5 21.5 48 48 48s48-21.5 48-48s-21.5-48-48-48h-48zm96-320c0-26.5-21.5-48-48-48s-48 21.5-48 48v48h48c26.5 0 48-21.5 48-48"/><path fill="currentColor" d="M928 64H96c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h832c17.7 0 32-14.3 32-32V96c0-17.7-14.3-32-32-32M672 560c61.9 0 112 50.1 112 112s-50.1 112-112 112s-112-50.1-112-112v-48h-96v48c0 61.9-50.1 112-112 112s-112-50.1-112-112s50.1-112 112-112h48v-96h-48c-61.9 0-112-50.1-112-112s50.1-112 112-112s112 50.1 112 112v48h96v-48c0-61.9 50.1-112 112-112s112 50.1 112 112s-50.1 112-112 112h-48v96z"/><path fill="currentColor" d="M464 464h96v96h-96zM352 304c-26.5 0-48 21.5-48 48s21.5 48 48 48h48v-48c0-26.5-21.5-48-48-48m-48 368c0 26.5 21.5 48 48 48s48-21.5 48-48v-48h-48c-26.5 0-48 21.5-48 48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MacCommandOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/><path fill="currentColor" d="M370.8 554.4c-54.6 0-98.8 44.2-98.8 98.8s44.2 98.8 98.8 98.8s98.8-44.2 98.8-98.8v-42.4h84.7v42.4c0 54.6 44.2 98.8 98.8 98.8s98.8-44.2 98.8-98.8s-44.2-98.8-98.8-98.8h-42.4v-84.7h42.4c54.6 0 98.8-44.2 98.8-98.8c0-54.6-44.2-98.8-98.8-98.8s-98.8 44.2-98.8 98.8v42.4h-84.7v-42.4c0-54.6-44.2-98.8-98.8-98.8S272 316.2 272 370.8s44.2 98.8 98.8 98.8h42.4v84.7h-42.4zm42.4 98.8c0 23.4-19 42.4-42.4 42.4s-42.4-19-42.4-42.4s19-42.4 42.4-42.4h42.4zm197.6-282.4c0-23.4 19-42.4 42.4-42.4s42.4 19 42.4 42.4s-19 42.4-42.4 42.4h-42.4zm0 240h42.4c23.4 0 42.4 19 42.4 42.4s-19 42.4-42.4 42.4s-42.4-19-42.4-42.4zM469.6 469.6h84.7v84.7h-84.7zm-98.8-56.4c-23.4 0-42.4-19-42.4-42.4s19-42.4 42.4-42.4s42.4 19 42.4 42.4v42.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M928 160H96c-17.7 0-32 14.3-32 32v640c0 17.7 14.3 32 32 32h832c17.7 0 32-14.3 32-32V192c0-17.7-14.3-32-32-32m-80.8 108.9L531.7 514.4c-7.8 6.1-18.7 6.1-26.5 0L189.6 268.9A7.2 7.2 0 0 1 194 256h648.8a7.2 7.2 0 0 1 4.4 12.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M928 160H96c-17.7 0-32 14.3-32 32v640c0 17.7 14.3 32 32 32h832c17.7 0 32-14.3 32-32V192c0-17.7-14.3-32-32-32m-40 110.8V792H136V270.8l-27.6-21.5l39.3-50.5l42.8 33.3h643.1l42.8-33.3l39.3 50.5zM833.6 232L512 482L190.4 232l-42.8-33.3l-39.3 50.5l27.6 21.5l341.6 265.6a55.99 55.99 0 0 0 68.7 0L888 270.8l27.6-21.5l-39.3-50.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M477.5 536.3L135.9 270.7l-27.5-21.4l27.6 21.5V792h752V270.8L546.2 536.3a55.99 55.99 0 0 1-68.7 0"/><path fill="currentColor" fill-opacity=".15" d="m876.3 198.8l39.3 50.5l-27.6 21.5l27.7-21.5l-39.3-50.5z"/><path fill="currentColor" d="M928 160H96c-17.7 0-32 14.3-32 32v640c0 17.7 14.3 32 32 32h832c17.7 0 32-14.3 32-32V192c0-17.7-14.3-32-32-32m-94.5 72.1L512 482L190.5 232.1zm54.5 38.7V792H136V270.8l-27.6-21.5l27.5 21.4l341.6 265.6a55.99 55.99 0 0 0 68.7 0zl27.6-21.5l-39.3-50.5h.1l39.3 50.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ManOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M874 120H622c-3.3 0-6 2.7-6 6v56c0 3.3 2.7 6 6 6h160.4L583.1 387.3c-50-38.5-111-59.3-175.1-59.3c-76.9 0-149.3 30-203.6 84.4S120 539.1 120 616s30 149.3 84.4 203.6C258.7 874 331.1 904 408 904s149.3-30 203.6-84.4C666 765.3 696 692.9 696 616c0-64.1-20.8-124.9-59.2-174.9L836 241.9V402c0 3.3 2.7 6 6 6h56c3.3 0 6-2.7 6-6V150c0-16.5-13.5-30-30-30M408 828c-116.9 0-212-95.1-212-212s95.1-212 212-212s212 95.1 212 212s-95.1 212-212 212"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MedicineBoxFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M839.2 278.1a32 32 0 0 0-30.4-22.1H736V144c0-17.7-14.3-32-32-32H320c-17.7 0-32 14.3-32 32v112h-72.8a31.9 31.9 0 0 0-30.4 22.1L112 502v378c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V502zM660 628c0 4.4-3.6 8-8 8H544v108c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8V636H372c-4.4 0-8-3.6-8-8v-48c0-4.4 3.6-8 8-8h108V464c0-4.4 3.6-8 8-8h48c4.4 0 8 3.6 8 8v108h108c4.4 0 8 3.6 8 8zm4-372H360v-72h304z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MedicineBoxOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M839.2 278.1a32 32 0 0 0-30.4-22.1H736V144c0-17.7-14.3-32-32-32H320c-17.7 0-32 14.3-32 32v112h-72.8a31.9 31.9 0 0 0-30.4 22.1L112 502v378c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V502zM360 184h304v72H360zm480 656H184V513.4L244.3 328h535.4L840 513.4zM652 572H544V464c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v108H372c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h108v108c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V636h108c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MedicineBoxTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M244.3 328L184 513.4V840h656V513.4L779.7 328zM660 628c0 4.4-3.6 8-8 8H544v108c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8V636H372c-4.4 0-8-3.6-8-8v-48c0-4.4 3.6-8 8-8h108V464c0-4.4 3.6-8 8-8h48c4.4 0 8 3.6 8 8v108h108c4.4 0 8 3.6 8 8z"/><path fill="currentColor" d="M652 572H544V464c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v108H372c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h108v108c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V636h108c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8"/><path fill="currentColor" d="M839.2 278.1a32 32 0 0 0-30.4-22.1H736V144c0-17.7-14.3-32-32-32H320c-17.7 0-32 14.3-32 32v112h-72.8a31.9 31.9 0 0 0-30.4 22.1L112 502v378c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V502zM360 184h304v72H360zm480 656H184V513.4L244.3 328h535.4L840 513.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediumCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m256 253.7l-40.8 39.1c-3.6 2.7-5.3 7.1-4.6 11.4v287.7c-.7 4.4 1 8.8 4.6 11.4l40 39.1v8.7H566.4v-8.3l41.3-40.1c4.1-4.1 4.1-5.3 4.1-11.4V422.5l-115 291.6h-15.5L347.5 422.5V618c-1.2 8.2 1.7 16.5 7.5 22.4l53.8 65.1v8.7H256v-8.7l53.8-65.1a26.1 26.1 0 0 0 7-22.4V392c.7-6.3-1.7-12.4-6.5-16.7l-47.8-57.6V309H411l114.6 251.5l100.9-251.3H768z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediumOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m834.7 279.8l61.3-58.9V208H683.7L532.4 586.4L360.3 208H137.7v12.9l71.6 86.6c7 6.4 10.6 15.8 9.7 25.2V673c2.2 12.3-1.7 24.8-10.3 33.7L128 805v12.7h228.6v-12.9l-80.6-98a39.99 39.99 0 0 1-11.1-33.7V378.7l200.7 439.2h23.3l172.6-439.2v349.9c0 9.2 0 11.1-6 17.2l-62.1 60.3V819h301.2v-12.9l-59.9-58.9c-5.2-4-7.9-10.7-6.8-17.2V297a18.1 18.1 0 0 1 6.8-17.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediumSquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M768 317.7l-40.8 39.1c-3.6 2.7-5.3 7.1-4.6 11.4v287.7c-.7 4.4 1 8.8 4.6 11.4l40 39.1v8.7H566.4v-8.3l41.3-40.1c4.1-4.1 4.1-5.3 4.1-11.4V422.5l-115 291.6h-15.5L347.5 422.5V618c-1.2 8.2 1.7 16.5 7.5 22.4l53.8 65.1v8.7H256v-8.7l53.8-65.1a26.1 26.1 0 0 0 7-22.4V392c.7-6.3-1.7-12.4-6.5-16.7l-47.8-57.6V309H411l114.6 251.5l100.9-251.3H768z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediumWorkmarkOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M517.2 590.55c0 3.55 0 4.36 2.4 6.55l13.43 13.25v.57h-59.57v-25.47a41.44 41.44 0 0 1-39.5 27.65c-30.61 0-52.84-24.25-52.84-68.87c0-41.8 23.99-69.69 57.65-69.69a35.15 35.15 0 0 1 34.61 21.67v-56.19a6.99 6.99 0 0 0-2.71-6.79l-12.8-12.45v-.56l59.33-7.04zm-43.74-8.09v-83.83a22.2 22.2 0 0 0-17.74-8.4c-14.48 0-28.47 13.25-28.47 52.62c0 36.86 12.07 49.88 27.1 49.88a23.91 23.91 0 0 0 19.11-10.27m83.23 28.46V497.74a7.65 7.65 0 0 0-2.4-6.79l-13.19-13.74v-.57h59.56v114.8c0 3.55 0 4.36 2.4 6.54l13.12 12.45v.57zm-2.16-175.67c0-13.4 10.74-24.25 23.99-24.25c13.25 0 23.98 10.86 23.98 24.25c0 13.4-10.73 24.25-23.98 24.25s-23.99-10.85-23.99-24.25m206.83 155.06c0 3.55 0 4.6 2.4 6.79l13.43 13.25v.57h-59.88V581.9a43.4 43.4 0 0 1-41.01 31.2c-26.55 0-40.78-19.56-40.78-56.59c0-17.86 0-37.43.56-59.41a6.91 6.91 0 0 0-2.4-6.55L620.5 477.2v-.57h59.09v73.81c0 24.25 3.51 40.42 18.54 40.42a23.96 23.96 0 0 0 19.35-12.2v-80.85a7.65 7.65 0 0 0-2.4-6.79l-13.27-13.82v-.57h59.56V590.3zm202.76 20.6c0-4.36.8-59.97.8-72.75c0-24.25-3.76-40.98-20.63-40.98a26.7 26.7 0 0 0-21.19 11.64a99.68 99.68 0 0 1 2.4 23.04c0 16.81-.56 38.23-.8 59.66a6.91 6.91 0 0 0 2.4 6.55l13.43 12.45v.56h-60.12c0-4.04.8-59.98.8-72.76c0-24.65-3.76-40.98-20.39-40.98c-8.2.3-15.68 4.8-19.83 11.96v82.46c0 3.56 0 4.37 2.4 6.55l13.11 12.45v.56h-59.48V498.15a7.65 7.65 0 0 0-2.4-6.8l-13.19-14.14v-.57H841v28.78c5.53-19 23.13-31.76 42.7-30.96c19.82 0 33.26 11.16 38.93 32.34a46.41 46.41 0 0 1 44.77-32.34c26.55 0 41.58 19.8 41.58 57.23c0 17.87-.56 38.24-.8 59.66a6.5 6.5 0 0 0 2.72 6.55l13.11 12.45v.57h-59.88zM215.87 593.3l17.66 17.05v.57h-89.62v-.57l17.99-17.05a6.91 6.91 0 0 0 2.4-6.55V477.69c0-4.6 0-10.83.8-16.16L104.66 613.1h-.72l-62.6-139.45c-1.37-3.47-1.77-3.72-2.65-6.06v91.43a32.08 32.08 0 0 0 2.96 17.87l25.19 33.46v.57H0v-.57l25.18-33.55a32.16 32.16 0 0 0 2.96-17.78V457.97A19.71 19.71 0 0 0 24 444.15L6.16 420.78v-.56h63.96l53.56 118.1l47.17-118.1h62.6v.56l-17.58 19.8a6.99 6.99 0 0 0-2.72 6.8v139.37a6.5 6.5 0 0 0 2.72 6.55m70.11-54.65v.56c0 34.6 17.67 48.5 38.38 48.5a43.5 43.5 0 0 0 40.77-24.97h.56c-7.2 34.2-28.14 50.36-59.48 50.36c-33.82 0-65.72-20.61-65.72-68.39c0-50.2 31.98-70.25 67.32-70.25c28.46 0 58.76 13.58 58.76 57.24v6.95zm0-6.95h39.42v-7.04c0-35.57-7.28-45.03-18.23-45.03c-13.27 0-21.35 14.15-21.35 52.07z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MehFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64M288 421a48.01 48.01 0 0 1 96 0a48.01 48.01 0 0 1-96 0m384 200c0 4.4-3.6 8-8 8H360c-4.4 0-8-3.6-8-8v-48c0-4.4 3.6-8 8-8h304c4.4 0 8 3.6 8 8zm16-152a48.01 48.01 0 0 1 0-96a48.01 48.01 0 0 1 0 96"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MehOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M288 421a48 48 0 1 0 96 0a48 48 0 1 0-96 0m352 0a48 48 0 1 0 96 0a48 48 0 1 0-96 0M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m263 711c-34.2 34.2-74 61-118.3 79.8C611 874.2 562.3 884 512 884c-50.3 0-99-9.8-144.8-29.2A370.4 370.4 0 0 1 248.9 775c-34.2-34.2-61-74-79.8-118.3C149.8 611 140 562.3 140 512s9.8-99 29.2-144.8A370.4 370.4 0 0 1 249 248.9c34.2-34.2 74-61 118.3-79.8C413 149.8 461.7 140 512 140c50.3 0 99 9.8 144.8 29.2A370.4 370.4 0 0 1 775.1 249c34.2 34.2 61 74 79.8 118.3C874.2 413 884 461.7 884 512s-9.8 99-29.2 144.8A368.89 368.89 0 0 1 775 775M664 565H360c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h304c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MehTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372M288 421a48.01 48.01 0 0 1 96 0a48.01 48.01 0 0 1-96 0m384 200c0 4.4-3.6 8-8 8H360c-4.4 0-8-3.6-8-8v-48c0-4.4 3.6-8 8-8h304c4.4 0 8 3.6 8 8zm16-152a48.01 48.01 0 0 1 0-96a48.01 48.01 0 0 1 0 96"/><path fill="currentColor" d="M288 421a48 48 0 1 0 96 0a48 48 0 1 0-96 0m376 144H360c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h304c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8m-24-144a48 48 0 1 0 96 0a48 48 0 1 0-96 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuFoldOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M408 442h480c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8H408c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m-8 204c0 4.4 3.6 8 8 8h480c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8H408c-4.4 0-8 3.6-8 8zm504-486H120c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 632H120c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8M115.4 518.9L271.7 642c5.8 4.6 14.4.5 14.4-6.9V388.9c0-7.4-8.5-11.5-14.4-6.9L115.4 505.1a8.74 8.74 0 0 0 0 13.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M904 160H120c-4.4 0-8 3.6-8 8v64c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-64c0-4.4-3.6-8-8-8m0 624H120c-4.4 0-8 3.6-8 8v64c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-64c0-4.4-3.6-8-8-8m0-312H120c-4.4 0-8 3.6-8 8v64c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-64c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuUnfoldOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M408 442h480c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8H408c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m-8 204c0 4.4 3.6 8 8 8h480c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8H408c-4.4 0-8 3.6-8 8zm504-486H120c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 632H120c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8M142.4 642.1L298.7 519a8.84 8.84 0 0 0 0-13.9L142.4 381.9c-5.8-4.6-14.4-.5-14.4 6.9v246.3a8.9 8.9 0 0 0 14.4 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MergeCellsOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M482.2 508.4L331.3 389c-3-2.4-7.3-.2-7.3 3.6V478H184V184h204v128c0 2.2 1.8 4 4 4h60c2.2 0 4-1.8 4-4V144c0-15.5-12.5-28-28-28H144c-15.5 0-28 12.5-28 28v736c0 15.5 12.5 28 28 28h284c15.5 0 28-12.5 28-28V712c0-2.2-1.8-4-4-4h-60c-2.2 0-4 1.8-4 4v128H184V546h140v85.4c0 3.8 4.4 6 7.3 3.6l150.9-119.4c2.4-1.8 2.4-5.4 0-7.2M880 116H596c-15.5 0-28 12.5-28 28v168c0 2.2 1.8 4 4 4h60c2.2 0 4-1.8 4-4V184h204v294H700v-85.4c0-3.8-4.3-6-7.3-3.6l-151 119.4c-2.3 1.8-2.3 5.3 0 7.1l151 119.5c2.9 2.3 7.3.2 7.3-3.6V546h140v294H636V712c0-2.2-1.8-4-4-4h-60c-2.2 0-4 1.8-4 4v168c0 15.5 12.5 28 28 28h284c15.5 0 28-12.5 28-28V144c0-15.5-12.5-28-28-28"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M924.3 338.4a447.57 447.57 0 0 0-96.1-143.3a443.09 443.09 0 0 0-143-96.3A443.91 443.91 0 0 0 512 64h-2c-60.5.3-119 12.3-174.1 35.9a444.08 444.08 0 0 0-141.7 96.5a445 445 0 0 0-95 142.8A449.89 449.89 0 0 0 65 514.1c.3 69.4 16.9 138.3 47.9 199.9v152c0 25.4 20.6 46 45.9 46h151.8a447.72 447.72 0 0 0 199.5 48h2.1c59.8 0 117.7-11.6 172.3-34.3A443.2 443.2 0 0 0 827 830.5c41.2-40.9 73.6-88.7 96.3-142c23.5-55.2 35.5-113.9 35.8-174.5c.2-60.9-11.6-120-34.8-175.6M312.4 560c-26.4 0-47.9-21.5-47.9-48s21.5-48 47.9-48s47.9 21.5 47.9 48s-21.4 48-47.9 48m199.6 0c-26.4 0-47.9-21.5-47.9-48s21.5-48 47.9-48s47.9 21.5 47.9 48s-21.5 48-47.9 48m199.6 0c-26.4 0-47.9-21.5-47.9-48s21.5-48 47.9-48s47.9 21.5 47.9 48s-21.5 48-47.9 48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 512a48 48 0 1 0 96 0a48 48 0 1 0-96 0m200 0a48 48 0 1 0 96 0a48 48 0 1 0-96 0m-400 0a48 48 0 1 0 96 0a48 48 0 1 0-96 0m661.2-173.6c-22.6-53.7-55-101.9-96.3-143.3a444.35 444.35 0 0 0-143.3-96.3C630.6 75.7 572.2 64 512 64h-2c-60.6.3-119.3 12.3-174.5 35.9a445.35 445.35 0 0 0-142 96.5c-40.9 41.3-73 89.3-95.2 142.8c-23 55.4-34.6 114.3-34.3 174.9A449.4 449.4 0 0 0 112 714v152a46 46 0 0 0 46 46h152.1A449.4 449.4 0 0 0 510 960h2.1c59.9 0 118-11.6 172.7-34.3a444.48 444.48 0 0 0 142.8-95.2c41.3-40.9 73.8-88.7 96.5-142c23.6-55.2 35.6-113.9 35.9-174.5c.3-60.9-11.5-120-34.8-175.6m-151.1 438C704 845.8 611 884 512 884h-1.7c-60.3-.3-120.2-15.3-173.1-43.5l-8.4-4.5H188V695.2l-4.5-8.4C155.3 633.9 140.3 574 140 513.7c-.4-99.7 37.7-193.3 107.6-263.8c69.8-70.5 163.1-109.5 262.8-109.9h1.7c50 0 98.5 9.7 144.2 28.9c44.6 18.7 84.6 45.6 119 80c34.3 34.3 61.3 74.4 80 119c19.4 46.2 29.1 95.2 28.9 145.8c-.6 99.6-39.7 192.9-110.1 262.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M775.3 248.9a369.62 369.62 0 0 0-119-80A370.2 370.2 0 0 0 512.1 140h-1.7c-99.7.4-193 39.4-262.8 109.9c-69.9 70.5-108 164.1-107.6 263.8c.3 60.3 15.3 120.2 43.5 173.1l4.5 8.4V836h140.8l8.4 4.5c52.9 28.2 112.8 43.2 173.1 43.5h1.7c99 0 192-38.2 262.1-107.6c70.4-69.8 109.5-163.1 110.1-262.7c.2-50.6-9.5-99.6-28.9-145.8a370.15 370.15 0 0 0-80-119M312 560a48.01 48.01 0 0 1 0-96a48.01 48.01 0 0 1 0 96m200 0a48.01 48.01 0 0 1 0-96a48.01 48.01 0 0 1 0 96m200 0a48.01 48.01 0 0 1 0-96a48.01 48.01 0 0 1 0 96"/><path fill="currentColor" d="M664 512a48 48 0 1 0 96 0a48 48 0 1 0-96 0m-400 0a48 48 0 1 0 96 0a48 48 0 1 0-96 0"/><path fill="currentColor" d="M925.2 338.4c-22.6-53.7-55-101.9-96.3-143.3a444.35 444.35 0 0 0-143.3-96.3C630.6 75.7 572.2 64 512 64h-2c-60.6.3-119.3 12.3-174.5 35.9a445.35 445.35 0 0 0-142 96.5c-40.9 41.3-73 89.3-95.2 142.8c-23 55.4-34.6 114.3-34.3 174.9A449.4 449.4 0 0 0 112 714v152a46 46 0 0 0 46 46h152.1A449.4 449.4 0 0 0 510 960h2.1c59.9 0 118-11.6 172.7-34.3a444.48 444.48 0 0 0 142.8-95.2c41.3-40.9 73.8-88.7 96.5-142c23.6-55.2 35.6-113.9 35.9-174.5c.3-60.9-11.5-120-34.8-175.6m-151.1 438C704 845.8 611 884 512 884h-1.7c-60.3-.3-120.2-15.3-173.1-43.5l-8.4-4.5H188V695.2l-4.5-8.4C155.3 633.9 140.3 574 140 513.7c-.4-99.7 37.7-193.3 107.6-263.8c69.8-70.5 163.1-109.5 262.8-109.9h1.7c50 0 98.5 9.7 144.2 28.9c44.6 18.7 84.6 45.6 119 80c34.3 34.3 61.3 74.4 80 119c19.4 46.2 29.1 95.2 28.9 145.8c-.6 99.6-39.7 192.9-110.1 262.7"/><path fill="currentColor" d="M464 512a48 48 0 1 0 96 0a48 48 0 1 0-96 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m192 472c0 4.4-3.6 8-8 8H328c-4.4 0-8-3.6-8-8v-48c0-4.4 3.6-8 8-8h368c4.4 0 8 3.6 8 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M696 480H328c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h368c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8"/><path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircleTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372m192 396c0 4.4-3.6 8-8 8H328c-4.4 0-8-3.6-8-8v-48c0-4.4 3.6-8 8-8h368c4.4 0 8 3.6 8 8z"/><path fill="currentColor" d="M696 480H328c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h368c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M872 474H152c-4.4 0-8 3.6-8 8v60c0 4.4 3.6 8 8 8h720c4.4 0 8-3.6 8-8v-60c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M704 536c0 4.4-3.6 8-8 8H328c-4.4 0-8-3.6-8-8v-48c0-4.4 3.6-8 8-8h368c4.4 0 8 3.6 8 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSquareOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M328 544h368c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8H328c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8"/><path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSquareTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/><path fill="currentColor" fill-opacity=".15" d="M184 840h656V184H184zm136-352c0-4.4 3.6-8 8-8h368c4.4 0 8 3.6 8 8v48c0 4.4-3.6 8-8 8H328c-4.4 0-8-3.6-8-8z"/><path fill="currentColor" d="M328 544h368c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8H328c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M744 62H280c-35.3 0-64 28.7-64 64v768c0 35.3 28.7 64 64 64h464c35.3 0 64-28.7 64-64V126c0-35.3-28.7-64-64-64M512 824c-22.1 0-40-17.9-40-40s17.9-40 40-40s40 17.9 40 40s-17.9 40-40 40"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M744 62H280c-35.3 0-64 28.7-64 64v768c0 35.3 28.7 64 64 64h464c35.3 0 64-28.7 64-64V126c0-35.3-28.7-64-64-64m-8 824H288V134h448zM472 784a40 40 0 1 0 80 0a40 40 0 1 0-80 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M744 64H280c-35.3 0-64 28.7-64 64v768c0 35.3 28.7 64 64 64h464c35.3 0 64-28.7 64-64V128c0-35.3-28.7-64-64-64m-8 824H288V136h448z"/><path fill="currentColor" fill-opacity=".15" d="M288 888h448V136H288zm224-142c22.1 0 40 17.9 40 40s-17.9 40-40 40s-40-17.9-40-40s17.9-40 40-40"/><path fill="currentColor" d="M472 786a40 40 0 1 0 80 0a40 40 0 1 0-80 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneyCollectFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M911.5 699.7a8 8 0 0 0-10.3-4.8L840 717.2V179c0-37.6-30.4-68-68-68H252c-37.6 0-68 30.4-68 68v538.2l-61.3-22.3c-.9-.3-1.8-.5-2.7-.5c-4.4 0-8 3.6-8 8V762c0 3.3 2.1 6.3 5.3 7.5L501 909.1c7.1 2.6 14.8 2.6 21.9 0l383.8-139.5c3.2-1.2 5.3-4.2 5.3-7.5v-59.6c0-1-.2-1.9-.5-2.8m-243.8-377L564 514.3h57.6c4.4 0 8 3.6 8 8v27.1c0 4.4-3.6 8-8 8h-76.3v39h76.3c4.4 0 8 3.6 8 8v27.1c0 4.4-3.6 8-8 8h-76.3V703c0 4.4-3.6 8-8 8h-49.9c-4.4 0-8-3.6-8-8v-63.4h-76c-4.4 0-8-3.6-8-8v-27.1c0-4.4 3.6-8 8-8h76v-39h-76c-4.4 0-8-3.6-8-8v-27.1c0-4.4 3.6-8 8-8h57L356.5 322.8c-2.1-3.8-.7-8.7 3.2-10.8c1.2-.7 2.5-1 3.8-1h55.7a8 8 0 0 1 7.1 4.4L511 484.2h3.3L599 315.4c1.3-2.7 4.1-4.4 7.1-4.4h54.5c4.4 0 8 3.6 8.1 7.9c0 1.3-.4 2.6-1 3.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneyCollectOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M911.5 700.7a8 8 0 0 0-10.3-4.8L840 718.2V180c0-37.6-30.4-68-68-68H252c-37.6 0-68 30.4-68 68v538.2l-61.3-22.3c-.9-.3-1.8-.5-2.7-.5c-4.4 0-8 3.6-8 8V763c0 3.3 2.1 6.3 5.3 7.5L501 910.1c7.1 2.6 14.8 2.6 21.9 0l383.8-139.5c3.2-1.2 5.3-4.2 5.3-7.5v-59.6c0-1-.2-1.9-.5-2.8M512 837.5l-256-93.1V184h512v560.4zM660.6 312h-54.5c-3 0-5.8 1.7-7.1 4.4l-84.7 168.8H511l-84.7-168.8a8 8 0 0 0-7.1-4.4h-55.7c-1.3 0-2.6.3-3.8 1c-3.9 2.1-5.3 7-3.2 10.8l103.9 191.6h-57c-4.4 0-8 3.6-8 8v27.1c0 4.4 3.6 8 8 8h76v39h-76c-4.4 0-8 3.6-8 8v27.1c0 4.4 3.6 8 8 8h76V704c0 4.4 3.6 8 8 8h49.9c4.4 0 8-3.6 8-8v-63.5h76.3c4.4 0 8-3.6 8-8v-27.1c0-4.4-3.6-8-8-8h-76.3v-39h76.3c4.4 0 8-3.6 8-8v-27.1c0-4.4-3.6-8-8-8H564l103.7-191.6c.6-1.2 1-2.5 1-3.8c-.1-4.3-3.7-7.9-8.1-7.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneyCollectTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="m256 744.4l256 93.1l256-93.1V184H256zM359.7 313c1.2-.7 2.5-1 3.8-1h55.7a8 8 0 0 1 7.1 4.4L511 485.2h3.3L599 316.4c1.3-2.7 4.1-4.4 7.1-4.4h54.5c4.4 0 8 3.6 8.1 7.9c0 1.3-.4 2.6-1 3.8L564 515.3h57.6c4.4 0 8 3.6 8 8v27.1c0 4.4-3.6 8-8 8h-76.3v39h76.3c4.4 0 8 3.6 8 8v27.1c0 4.4-3.6 8-8 8h-76.3V704c0 4.4-3.6 8-8 8h-49.9c-4.4 0-8-3.6-8-8v-63.4h-76c-4.4 0-8-3.6-8-8v-27.1c0-4.4 3.6-8 8-8h76v-39h-76c-4.4 0-8-3.6-8-8v-27.1c0-4.4 3.6-8 8-8h57L356.5 323.8c-2.1-3.8-.7-8.7 3.2-10.8"/><path fill="currentColor" d="M911.5 700.7a8 8 0 0 0-10.3-4.8L840 718.2V180c0-37.6-30.4-68-68-68H252c-37.6 0-68 30.4-68 68v538.2l-61.3-22.3c-.9-.3-1.8-.5-2.7-.5c-4.4 0-8 3.6-8 8V763c0 3.3 2.1 6.3 5.3 7.5L501 910.1c7.1 2.6 14.8 2.6 21.9 0l383.8-139.5c3.2-1.2 5.3-4.2 5.3-7.5v-59.6c0-1-.2-1.9-.5-2.8M768 744.4l-256 93.1l-256-93.1V184h512z"/><path fill="currentColor" d="M460.4 515.4h-57c-4.4 0-8 3.6-8 8v27.1c0 4.4 3.6 8 8 8h76v39h-76c-4.4 0-8 3.6-8 8v27.1c0 4.4 3.6 8 8 8h76V704c0 4.4 3.6 8 8 8h49.9c4.4 0 8-3.6 8-8v-63.5h76.3c4.4 0 8-3.6 8-8v-27.1c0-4.4-3.6-8-8-8h-76.3v-39h76.3c4.4 0 8-3.6 8-8v-27.1c0-4.4-3.6-8-8-8H564l103.7-191.6c.6-1.2 1-2.5 1-3.8c-.1-4.3-3.7-7.9-8.1-7.9h-54.5c-3 0-5.8 1.7-7.1 4.4l-84.7 168.8H511l-84.7-168.8a8 8 0 0 0-7.1-4.4h-55.7c-1.3 0-2.6.3-3.8 1c-3.9 2.1-5.3 7-3.2 10.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MonitorOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m692.8 412.7l.2-.2l-34.6-44.3a7.97 7.97 0 0 0-11.2-1.4l-50.4 39.3l-70.5-90.1a7.97 7.97 0 0 0-11.2-1.4l-37.9 29.7a7.97 7.97 0 0 0-1.4 11.2l70.5 90.2l-.2.1l34.6 44.3c2.7 3.5 7.7 4.1 11.2 1.4l50.4-39.3l64.1 82c2.7 3.5 7.7 4.1 11.2 1.4l37.9-29.6c3.5-2.7 4.1-7.7 1.4-11.2zM608 112c-167.9 0-304 136.1-304 304c0 70.3 23.9 135 63.9 186.5L114.3 856.1a8.03 8.03 0 0 0 0 11.3l42.3 42.3c3.1 3.1 8.2 3.1 11.3 0l253.6-253.6C473 696.1 537.7 720 608 720c167.9 0 304-136.1 304-304S775.9 112 608 112m161.2 465.2C726.2 620.3 668.9 644 608 644s-118.2-23.7-161.2-66.8C403.7 534.2 380 476.9 380 416s23.7-118.2 66.8-161.2c43-43.1 100.3-66.8 161.2-66.8s118.2 23.7 161.2 66.8c43.1 43 66.8 100.3 66.8 161.2s-23.7 118.2-66.8 161.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M456 231a56 56 0 1 0 112 0a56 56 0 1 0-112 0m0 280a56 56 0 1 0 112 0a56 56 0 1 0-112 0m0 280a56 56 0 1 0 112 0a56 56 0 1 0-112 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NodeCollapseOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M952 612c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8H298c-14.2-35.2-48.7-60-89-60c-53 0-96 43-96 96s43 96 96 96c40.3 0 74.8-24.8 89-60h150.3v152c0 55.2 44.8 100 100 100H952c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8H548.3c-15.5 0-28-12.5-28-28V612zM451.7 313.7l172.5 136.2c6.3 5.1 15.8.5 15.8-7.7V344h264c4.4 0 8-3.6 8-8v-60c0-4.4-3.6-8-8-8H640v-98.2c0-8.1-9.4-12.8-15.8-7.7L451.7 298.3c-4.9 3.9-4.9 11.5 0 15.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NodeExpandOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M952 612c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8H298c-14.2-35.2-48.7-60-89-60c-53 0-96 43-96 96s43 96 96 96c40.3 0 74.8-24.8 89-60h150.3v152c0 55.2 44.8 100 100 100H952c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8H548.3c-15.5 0-28-12.5-28-28V612zM456 344h264v98.2c0 8.1 9.5 12.8 15.8 7.7l172.5-136.2c5-3.9 5-11.4 0-15.3L735.8 162.1c-6.4-5.1-15.8-.5-15.8 7.7V268H456c-4.4 0-8 3.6-8 8v60c0 4.4 3.6 8 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NodeIndexOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M843.5 737.4c-12.4-75.2-79.2-129.1-155.3-125.4S550.9 676 546 752c-153.5-4.8-208-40.7-199.1-113.7c3.3-27.3 19.8-41.9 50.1-49c18.4-4.3 38.8-4.9 57.3-3.2c1.7.2 3.5.3 5.2.5c11.3 2.7 22.8 5 34.3 6.8c34.1 5.6 68.8 8.4 101.8 6.6c92.8-5 156-45.9 159.2-132.7c3.1-84.1-54.7-143.7-147.9-183.6c-29.9-12.8-61.6-22.7-93.3-30.2c-14.3-3.4-26.3-5.7-35.2-7.2c-7.9-75.9-71.5-133.8-147.8-134.4c-76.3-.6-140.9 56.1-150.1 131.9s40 146.3 114.2 163.9c74.2 17.6 149.9-23.3 175.7-95.1c9.4 1.7 18.7 3.6 28 5.8c28.2 6.6 56.4 15.4 82.4 26.6c70.7 30.2 109.3 70.1 107.5 119.9c-1.6 44.6-33.6 65.2-96.2 68.6c-27.5 1.5-57.6-.9-87.3-5.8c-8.3-1.4-15.9-2.8-22.6-4.3c-3.9-.8-6.6-1.5-7.8-1.8l-3.1-.6c-2.2-.3-5.9-.8-10.7-1.3c-25-2.3-52.1-1.5-78.5 4.6c-55.2 12.9-93.9 47.2-101.1 105.8c-15.7 126.2 78.6 184.7 276 188.9c29.1 70.4 106.4 107.9 179.6 87c73.3-20.9 119.3-93.4 106.9-168.6M329.1 345.2c-46 0-83.3-37.3-83.3-83.3s37.3-83.3 83.3-83.3s83.3 37.3 83.3 83.3s-37.3 83.3-83.3 83.3M695.6 845c-46 0-83.3-37.3-83.3-83.3s37.3-83.3 83.3-83.3s83.3 37.3 83.3 83.3s-37.3 83.3-83.3 83.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112c-3.8 0-7.7.7-11.6 2.3L292 345.9H128c-8.8 0-16 7.4-16 16.6v299c0 9.2 7.2 16.6 16 16.6h101.6c-3.7 11.6-5.6 23.9-5.6 36.4c0 65.9 53.8 119.5 120 119.5c55.4 0 102.1-37.6 115.9-88.4l408.6 164.2c3.9 1.5 7.8 2.3 11.6 2.3c16.9 0 32-14.2 32-33.2V145.2C912 126.2 897 112 880 112M344 762.3c-26.5 0-48-21.4-48-47.8c0-11.2 3.9-21.9 11-30.4l84.9 34.1c-2 24.6-22.7 44.1-47.9 44.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112c-3.8 0-7.7.7-11.6 2.3L292 345.9H128c-8.8 0-16 7.4-16 16.6v299c0 9.2 7.2 16.6 16 16.6h101.7c-3.7 11.6-5.7 23.9-5.7 36.4c0 65.9 53.8 119.5 120 119.5c55.4 0 102.1-37.6 115.9-88.4l408.6 164.2c3.9 1.5 7.8 2.3 11.6 2.3c16.9 0 32-14.2 32-33.2V145.2C912 126.2 897 112 880 112M344 762.3c-26.5 0-48-21.4-48-47.8c0-11.2 3.9-21.9 11-30.4l84.9 34.1c-2 24.6-22.7 44.1-47.9 44.1m496 58.4L318.8 611.3l-12.9-5.2H184V417.9h121.9l12.9-5.2L840 203.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M229.6 678.1c-3.7 11.6-5.6 23.9-5.6 36.4c0-12.5 2-24.8 5.7-36.4zm76.3-260.2H184v188.2h121.9l12.9 5.2L840 820.7V203.3L318.8 412.7z"/><path fill="currentColor" d="M880 112c-3.8 0-7.7.7-11.6 2.3L292 345.9H128c-8.8 0-16 7.4-16 16.6v299c0 9.2 7.2 16.6 16 16.6h101.7c-3.7 11.6-5.7 23.9-5.7 36.4c0 65.9 53.8 119.5 120 119.5c55.4 0 102.1-37.6 115.9-88.4l408.6 164.2c3.9 1.5 7.8 2.3 11.6 2.3c16.9 0 32-14.2 32-33.2V145.2C912 126.2 897 112 880 112M344 762.3c-26.5 0-48-21.4-48-47.8c0-11.2 3.9-21.9 11-30.4l84.9 34.1c-2 24.6-22.7 44.1-47.9 44.1m496 58.4L318.8 611.3l-12.9-5.2H184V417.9h121.9l12.9-5.2L840 203.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumberOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M872 394c4.4 0 8-3.6 8-8v-60c0-4.4-3.6-8-8-8H708V152c0-4.4-3.6-8-8-8h-64c-4.4 0-8 3.6-8 8v166H400V152c0-4.4-3.6-8-8-8h-64c-4.4 0-8 3.6-8 8v166H152c-4.4 0-8 3.6-8 8v60c0 4.4 3.6 8 8 8h168v236H152c-4.4 0-8 3.6-8 8v60c0 4.4 3.6 8 8 8h168v166c0 4.4 3.6 8 8 8h64c4.4 0 8-3.6 8-8V706h228v166c0 4.4 3.6 8 8 8h64c4.4 0 8-3.6 8-8V706h164c4.4 0 8-3.6 8-8v-60c0-4.4-3.6-8-8-8H708V394zM628 630H400V394h228z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OneToOneOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M316 672h60c4.4 0 8-3.6 8-8V360c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8v304c0 4.4 3.6 8 8 8m196-50c22.1 0 40-17.9 40-39c0-23.1-17.9-41-40-41s-40 17.9-40 41c0 21.1 17.9 39 40 39m0-140c22.1 0 40-17.9 40-39c0-23.1-17.9-41-40-41s-40 17.9-40 41c0 21.1 17.9 39 40 39"/><path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/><path fill="currentColor" d="M648 672h60c4.4 0 8-3.6 8-8V360c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8v304c0 4.4 3.6 8 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OrderedListOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M920 760H336c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h584c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-568H336c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h584c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 284H336c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h584c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8M216 712H100c-2.2 0-4 1.8-4 4v34c0 2.2 1.8 4 4 4h72.4v20.5h-35.7c-2.2 0-4 1.8-4 4v34c0 2.2 1.8 4 4 4h35.7V838H100c-2.2 0-4 1.8-4 4v34c0 2.2 1.8 4 4 4h116c2.2 0 4-1.8 4-4V716c0-2.2-1.8-4-4-4M100 188h38v120c0 2.2 1.8 4 4 4h40c2.2 0 4-1.8 4-4V152c0-4.4-3.6-8-8-8h-78c-2.2 0-4 1.8-4 4v36c0 2.2 1.8 4 4 4m116 240H100c-2.2 0-4 1.8-4 4v36c0 2.2 1.8 4 4 4h68.4l-70.3 77.7a8.3 8.3 0 0 0-2.1 5.4V592c0 2.2 1.8 4 4 4h116c2.2 0 4-1.8 4-4v-36c0-2.2-1.8-4-4-4h-68.4l70.3-77.7a8.3 8.3 0 0 0 2.1-5.4V432c0-2.2-1.8-4-4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperClipOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M779.3 196.6c-94.2-94.2-247.6-94.2-341.7 0l-261 260.8c-1.7 1.7-2.6 4-2.6 6.4s.9 4.7 2.6 6.4l36.9 36.9a9 9 0 0 0 12.7 0l261-260.8c32.4-32.4 75.5-50.2 121.3-50.2s88.9 17.8 121.2 50.2c32.4 32.4 50.2 75.5 50.2 121.2c0 45.8-17.8 88.8-50.2 121.2l-266 265.9l-43.1 43.1c-40.3 40.3-105.8 40.3-146.1 0c-19.5-19.5-30.2-45.4-30.2-73s10.7-53.5 30.2-73l263.9-263.8c6.7-6.6 15.5-10.3 24.9-10.3h.1c9.4 0 18.1 3.7 24.7 10.3c6.7 6.7 10.3 15.5 10.3 24.9c0 9.3-3.7 18.1-10.3 24.7L372.4 653c-1.7 1.7-2.6 4-2.6 6.4s.9 4.7 2.6 6.4l36.9 36.9a9 9 0 0 0 12.7 0l215.6-215.6c19.9-19.9 30.8-46.3 30.8-74.4s-11-54.6-30.8-74.4c-41.1-41.1-107.9-41-149 0L463 364L224.8 602.1A172.22 172.22 0 0 0 174 724.8c0 46.3 18.1 89.8 50.8 122.5c33.9 33.8 78.3 50.7 122.7 50.7c44.4 0 88.8-16.9 122.6-50.7l309.2-309C824.8 492.7 850 432 850 367.5c.1-64.6-25.1-125.3-70.7-170.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PartitionOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640.6 429.8h257.1c7.9 0 14.3-6.4 14.3-14.3V158.3c0-7.9-6.4-14.3-14.3-14.3H640.6c-7.9 0-14.3 6.4-14.3 14.3v92.9H490.6c-3.9 0-7.1 3.2-7.1 7.1v221.5h-85.7v-96.5c0-7.9-6.4-14.3-14.3-14.3H126.3c-7.9 0-14.3 6.4-14.3 14.3v257.2c0 7.9 6.4 14.3 14.3 14.3h257.1c7.9 0 14.3-6.4 14.3-14.3V544h85.7v221.5c0 3.9 3.2 7.1 7.1 7.1h135.7v92.9c0 7.9 6.4 14.3 14.3 14.3h257.1c7.9 0 14.3-6.4 14.3-14.3v-257c0-7.9-6.4-14.3-14.3-14.3h-257c-7.9 0-14.3 6.4-14.3 14.3v100h-78.6v-393h78.6v100c0 7.9 6.4 14.3 14.3 14.3m53.5-217.9h150V362h-150zM329.9 587h-150V437h150zm364.2 75.1h150v150.1h-150z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m-80 600c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8V360c0-4.4 3.6-8 8-8h48c4.4 0 8 3.6 8 8zm224 0c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8V360c0-4.4 3.6-8 8-8h48c4.4 0 8 3.6 8 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseCircleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372m-88-532h-48c-4.4 0-8 3.6-8 8v304c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V360c0-4.4-3.6-8-8-8m224 0h-48c-4.4 0-8 3.6-8 8v304c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V360c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseCircleTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372m-80 524c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8V360c0-4.4 3.6-8 8-8h48c4.4 0 8 3.6 8 8zm224 0c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8V360c0-4.4 3.6-8 8-8h48c4.4 0 8 3.6 8 8z"/><path fill="currentColor" d="M424 352h-48c-4.4 0-8 3.6-8 8v304c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V360c0-4.4-3.6-8-8-8m224 0h-48c-4.4 0-8 3.6-8 8v304c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V360c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M304 176h80v672h-80zm408 0h-64c-4.4 0-8 3.6-8 8v656c0 4.4 3.6 8 8 8h64c4.4 0 8-3.6 8-8V184c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PayCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m166.6 246.8L567.5 515.6h62c4.4 0 8 3.6 8 8v29.9c0 4.4-3.6 8-8 8h-82V603h82c4.4 0 8 3.6 8 8v29.9c0 4.4-3.6 8-8 8h-82V717c0 4.4-3.6 8-8 8h-54.3c-4.4 0-8-3.6-8-8v-68.1h-81.7c-4.4 0-8-3.6-8-8V611c0-4.4 3.6-8 8-8h81.7v-41.5h-81.7c-4.4 0-8-3.6-8-8v-29.9c0-4.4 3.6-8 8-8h61.4L345.4 310.8a8.07 8.07 0 0 1 7-11.9h60.7c3 0 5.8 1.7 7.1 4.4l90.6 180h3.4l90.6-180a8 8 0 0 1 7.1-4.4h59.5c4.4 0 8 3.6 8 8c.2 1.4-.2 2.7-.8 3.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PayCircleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372m159.6-585h-59.5c-3 0-5.8 1.7-7.1 4.4l-90.6 180H511l-90.6-180a8 8 0 0 0-7.1-4.4h-60.7c-1.3 0-2.6.3-3.8 1c-3.9 2.1-5.3 7-3.2 10.9L457 515.7h-61.4c-4.4 0-8 3.6-8 8v29.9c0 4.4 3.6 8 8 8h81.7V603h-81.7c-4.4 0-8 3.6-8 8v29.9c0 4.4 3.6 8 8 8h81.7V717c0 4.4 3.6 8 8 8h54.3c4.4 0 8-3.6 8-8v-68.1h82c4.4 0 8-3.6 8-8V611c0-4.4-3.6-8-8-8h-82v-41.5h82c4.4 0 8-3.6 8-8v-29.9c0-4.4-3.6-8-8-8h-62l111.1-204.8c.6-1.2 1-2.5 1-3.8c-.1-4.4-3.7-8-8.1-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PercentageOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m855.7 210.8l-42.4-42.4a8.03 8.03 0 0 0-11.3 0L168.3 801.9a8.03 8.03 0 0 0 0 11.3l42.4 42.4c3.1 3.1 8.2 3.1 11.3 0L855.6 222c3.2-3 3.2-8.1.1-11.2M304 448c79.4 0 144-64.6 144-144s-64.6-144-144-144s-144 64.6-144 144s64.6 144 144 144m0-216c39.7 0 72 32.3 72 72s-32.3 72-72 72s-72-32.3-72-72s32.3-72 72-72m416 344c-79.4 0-144 64.6-144 144s64.6 144 144 144s144-64.6 144-144s-64.6-144-144-144m0 216c-39.7 0-72-32.3-72-72s32.3-72 72-72s72 32.3 72 72s-32.3 72-72 72"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M885.6 230.2L779.1 123.8a80.83 80.83 0 0 0-57.3-23.8c-21.7 0-42.1 8.5-57.4 23.8L549.8 238.4a80.83 80.83 0 0 0-23.8 57.3c0 21.7 8.5 42.1 23.8 57.4l83.8 83.8A393.82 393.82 0 0 1 553.1 553A395.34 395.34 0 0 1 437 633.8L353.2 550a80.83 80.83 0 0 0-57.3-23.8c-21.7 0-42.1 8.5-57.4 23.8L123.8 664.5a80.89 80.89 0 0 0-23.8 57.4c0 21.7 8.5 42.1 23.8 57.4l106.3 106.3c24.4 24.5 58.1 38.4 92.7 38.4c7.3 0 14.3-.6 21.2-1.8c134.8-22.2 268.5-93.9 376.4-201.7C828.2 612.8 899.8 479.2 922.3 344c6.8-41.3-6.9-83.8-36.7-113.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M877.1 238.7L770.6 132.3c-13-13-30.4-20.3-48.8-20.3s-35.8 7.2-48.8 20.3L558.3 246.8c-13 13-20.3 30.5-20.3 48.9c0 18.5 7.2 35.8 20.3 48.9l89.6 89.7a405.46 405.46 0 0 1-86.4 127.3c-36.7 36.9-79.6 66-127.2 86.6l-89.6-89.7c-13-13-30.4-20.3-48.8-20.3a68.2 68.2 0 0 0-48.8 20.3L132.3 673c-13 13-20.3 30.5-20.3 48.9c0 18.5 7.2 35.8 20.3 48.9l106.4 106.4c22.2 22.2 52.8 34.9 84.2 34.9c6.5 0 12.8-.5 19.2-1.6c132.4-21.8 263.8-92.3 369.9-198.3C818 606 888.4 474.6 910.4 342.1c6.3-37.6-6.3-76.3-33.3-103.4m-37.6 91.5c-19.5 117.9-82.9 235.5-178.4 331s-213 158.9-330.9 178.4c-14.8 2.5-30-2.5-40.8-13.2L184.9 721.9L295.7 611l119.8 120l.9.9l21.6-8a481.29 481.29 0 0 0 285.7-285.8l8-21.6l-120.8-120.7l110.8-110.9l104.5 104.5c10.8 10.8 15.8 26 13.3 40.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M721.7 184.9L610.9 295.8l120.8 120.7l-8 21.6A481.29 481.29 0 0 1 438 723.9l-21.6 8l-.9-.9l-119.8-120l-110.8 110.9l104.5 104.5c10.8 10.7 26 15.7 40.8 13.2c117.9-19.5 235.4-82.9 330.9-178.4s158.9-213.1 178.4-331c2.5-14.8-2.5-30-13.3-40.8z"/><path fill="currentColor" d="M877.1 238.7L770.6 132.3c-13-13-30.4-20.3-48.8-20.3s-35.8 7.2-48.8 20.3L558.3 246.8c-13 13-20.3 30.5-20.3 48.9c0 18.5 7.2 35.8 20.3 48.9l89.6 89.7a405.46 405.46 0 0 1-86.4 127.3c-36.7 36.9-79.6 66-127.2 86.6l-89.6-89.7c-13-13-30.4-20.3-48.8-20.3a68.2 68.2 0 0 0-48.8 20.3L132.3 673c-13 13-20.3 30.5-20.3 48.9c0 18.5 7.2 35.8 20.3 48.9l106.4 106.4c22.2 22.2 52.8 34.9 84.2 34.9c6.5 0 12.8-.5 19.2-1.6c132.4-21.8 263.8-92.3 369.9-198.3C818 606 888.4 474.6 910.4 342.1c6.3-37.6-6.3-76.3-33.3-103.4m-37.6 91.5c-19.5 117.9-82.9 235.5-178.4 331s-213 158.9-330.9 178.4c-14.8 2.5-30-2.5-40.8-13.2L184.9 721.9L295.7 611l119.8 120l.9.9l21.6-8a481.29 481.29 0 0 0 285.7-285.8l8-21.6l-120.8-120.7l110.8-110.9l104.5 104.5c10.8 10.8 15.8 26 13.3 40.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PicCenterOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M952 792H72c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h880c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-632H72c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h880c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8M848 660c8.8 0 16-7.2 16-16V380c0-8.8-7.2-16-16-16H176c-8.8 0-16 7.2-16 16v264c0 8.8 7.2 16 16 16zM232 436h560v152H232z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PicLeftOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M952 792H72c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h880c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-632H72c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h880c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8M608 660c8.8 0 16-7.2 16-16V380c0-8.8-7.2-16-16-16H96c-8.8 0-16 7.2-16 16v264c0 8.8 7.2 16 16 16zM152 436h400v152H152zm552 210c0 4.4 3.6 8 8 8h224c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8H712c-4.4 0-8 3.6-8 8zm8-204h224c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8H712c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PicRightOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M952 792H72c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h880c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-632H72c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h880c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-24 500c8.8 0 16-7.2 16-16V380c0-8.8-7.2-16-16-16H416c-8.8 0-16 7.2-16 16v264c0 8.8 7.2 16 16 16zM472 436h400v152H472zM80 646c0 4.4 3.6 8 8 8h224c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8H88c-4.4 0-8 3.6-8 8zm8-204h224c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8H88c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PictureFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M928 160H96c-17.7 0-32 14.3-32 32v640c0 17.7 14.3 32 32 32h832c17.7 0 32-14.3 32-32V192c0-17.7-14.3-32-32-32M338 304c35.3 0 64 28.7 64 64s-28.7 64-64 64s-64-28.7-64-64s28.7-64 64-64m513.9 437.1a8.11 8.11 0 0 1-5.2 1.9H177.2c-4.4 0-8-3.6-8-8c0-1.9.7-3.7 1.9-5.2l170.3-202c2.8-3.4 7.9-3.8 11.3-1c.3.3.7.6 1 1l99.4 118l158.1-187.5c2.8-3.4 7.9-3.8 11.3-1c.3.3.7.6 1 1l229.6 271.6c2.6 3.3 2.2 8.4-1.2 11.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PictureOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M928 160H96c-17.7 0-32 14.3-32 32v640c0 17.7 14.3 32 32 32h832c17.7 0 32-14.3 32-32V192c0-17.7-14.3-32-32-32m-40 632H136v-39.9l138.5-164.3l150.1 178L658.1 489L888 761.6zm0-129.8L664.2 396.8c-3.2-3.8-9-3.8-12.2 0L424.6 666.4l-144-170.7c-3.2-3.8-9-3.8-12.2 0L136 652.7V232h752zM304 456a88 88 0 1 0 0-176a88 88 0 0 0 0 176m0-116c15.5 0 28 12.5 28 28s-12.5 28-28 28s-28-12.5-28-28s12.5-28 28-28"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PictureTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M928 160H96c-17.7 0-32 14.3-32 32v640c0 17.7 14.3 32 32 32h832c17.7 0 32-14.3 32-32V192c0-17.7-14.3-32-32-32m-40 632H136v-39.9l138.5-164.3l150.1 178L658.1 489L888 761.6zm0-129.8L664.2 396.8c-3.2-3.8-9-3.8-12.2 0L424.6 666.4l-144-170.7c-3.2-3.8-9-3.8-12.2 0L136 652.7V232h752z"/><path fill="currentColor" fill-opacity=".15" d="m424.6 765.8l-150.1-178L136 752.1V792h752v-30.4L658.1 489z"/><path fill="currentColor" fill-opacity=".15" d="m136 652.7l132.4-157c3.2-3.8 9-3.8 12.2 0l144 170.7L652 396.8c3.2-3.8 9-3.8 12.2 0L888 662.2V232H136zM304 280a88 88 0 1 1 0 176a88 88 0 0 1 0-176"/><path fill="currentColor" fill-opacity=".15" d="M276 368a28 28 0 1 0 56 0a28 28 0 1 0-56 0"/><path fill="currentColor" d="M304 456a88 88 0 1 0 0-176a88 88 0 0 0 0 176m0-116c15.5 0 28 12.5 28 28s-12.5 28-28 28s-28-12.5-28-28s12.5-28 28-28"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PieChartFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M863.1 518.5H505.5V160.9c0-4.4-3.6-8-8-8h-26a398.57 398.57 0 0 0-282.5 117a397.47 397.47 0 0 0-85.6 127C82.6 446.2 72 498.5 72 552.5S82.6 658.7 103.4 708c20.1 47.5 48.9 90.3 85.6 127c36.7 36.7 79.4 65.5 127 85.6a396.64 396.64 0 0 0 155.6 31.5a398.57 398.57 0 0 0 282.5-117c36.7-36.7 65.5-79.4 85.6-127a396.64 396.64 0 0 0 31.5-155.6v-26c-.1-4.4-3.7-8-8.1-8M951 463l-2.6-28.2c-8.5-92-49.3-178.8-115.1-244.3A398.5 398.5 0 0 0 588.4 75.6L560.1 73c-4.7-.4-8.7 3.2-8.7 7.9v383.7c0 4.4 3.6 8 8 8l383.6-1c4.7-.1 8.4-4 8-8.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PieChartOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M864 518H506V160c0-4.4-3.6-8-8-8h-26a398.46 398.46 0 0 0-282.8 117.1a398.19 398.19 0 0 0-85.7 127.1A397.61 397.61 0 0 0 72 552a398.46 398.46 0 0 0 117.1 282.8c36.7 36.7 79.5 65.6 127.1 85.7A397.61 397.61 0 0 0 472 952a398.46 398.46 0 0 0 282.8-117.1c36.7-36.7 65.6-79.5 85.7-127.1A397.61 397.61 0 0 0 872 552v-26c0-4.4-3.6-8-8-8M705.7 787.8A331.59 331.59 0 0 1 470.4 884c-88.1-.4-170.9-34.9-233.2-97.2C174.5 724.1 140 640.7 140 552c0-88.7 34.5-172.1 97.2-234.8c54.6-54.6 124.9-87.9 200.8-95.5V586h364.3c-7.7 76.3-41.3 147-96.6 201.8M952 462.4l-2.6-28.2c-8.5-92.1-49.4-179-115.2-244.6A399.4 399.4 0 0 0 589 74.6L560.7 72c-4.7-.4-8.7 3.2-8.7 7.9V464c0 4.4 3.6 8 8 8l384-1c4.7 0 8.4-4 8-8.6m-332.2-58.2V147.6a332.24 332.24 0 0 1 166.4 89.8c45.7 45.6 77 103.6 90 166.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PieChartTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M316.2 920.5c-47.6-20.1-90.4-49-127.1-85.7a398.19 398.19 0 0 1-85.7-127.1A397.12 397.12 0 0 1 72 552.2v.2a398.57 398.57 0 0 0 117 282.5c36.7 36.7 79.4 65.5 127 85.6A396.64 396.64 0 0 0 471.6 952c27 0 53.6-2.7 79.7-7.9c-25.9 5.2-52.4 7.8-79.3 7.8c-54 .1-106.4-10.5-155.8-31.4M560 472c-4.4 0-8-3.6-8-8V79.9c0-1.3.3-2.5.9-3.6c-.9 1.3-1.5 2.9-1.5 4.6v383.7c0 4.4 3.6 8 8 8l383.6-1c1.6 0 3.1-.5 4.4-1.3c-1 .5-2.2.7-3.4.7z"/><path fill="currentColor" fill-opacity=".15" d="M619.8 147.6v256.6l256.4-.7c-13-62.5-44.3-120.5-90-166.1a332.24 332.24 0 0 0-166.4-89.8"/><path fill="currentColor" fill-opacity=".15" d="M438 221.7c-75.9 7.6-146.2 40.9-200.8 95.5C174.5 379.9 140 463.3 140 552s34.5 172.1 97.2 234.8c62.3 62.3 145.1 96.8 233.2 97.2c88.2.4 172.7-34.1 235.3-96.2C761 733 794.6 662.3 802.3 586H438z"/><path fill="currentColor" d="M864 518H506V160c0-4.4-3.6-8-8-8h-26a398.46 398.46 0 0 0-282.8 117.1a398.19 398.19 0 0 0-85.7 127.1A397.61 397.61 0 0 0 72 552v.2c0 53.9 10.6 106.2 31.4 155.5c20.1 47.6 49 90.4 85.7 127.1c36.7 36.7 79.5 65.6 127.1 85.7A397.61 397.61 0 0 0 472 952c26.9 0 53.4-2.6 79.3-7.8c26.1-5.3 51.7-13.1 76.4-23.6c47.6-20.1 90.4-49 127.1-85.7c36.7-36.7 65.6-79.5 85.7-127.1A397.61 397.61 0 0 0 872 552v-26c0-4.4-3.6-8-8-8M705.7 787.8A331.59 331.59 0 0 1 470.4 884c-88.1-.4-170.9-34.9-233.2-97.2C174.5 724.1 140 640.7 140 552s34.5-172.1 97.2-234.8c54.6-54.6 124.9-87.9 200.8-95.5V586h364.3c-7.7 76.3-41.3 147-96.6 201.8"/><path fill="currentColor" d="m952 462.4l-2.6-28.2c-8.5-92.1-49.4-179-115.2-244.6A399.4 399.4 0 0 0 589 74.6L560.7 72c-3.4-.3-6.4 1.5-7.8 4.3a8.7 8.7 0 0 0-.9 3.6V464c0 4.4 3.6 8 8 8l384-1c1.2 0 2.3-.3 3.4-.7a8.1 8.1 0 0 0 4.6-7.9m-332.2-58.2V147.6a332.24 332.24 0 0 1 166.4 89.8c45.7 45.6 77 103.6 90 166.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m144.1 454.9L437.7 677.8a8.02 8.02 0 0 1-12.7-6.5V353.7a8 8 0 0 1 12.7-6.5L656.1 506a7.9 7.9 0 0 1 0 12.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayCircleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" d="m719.4 499.1l-296.1-215A15.9 15.9 0 0 0 398 297v430c0 13.1 14.8 20.5 25.3 12.9l296.1-215a15.9 15.9 0 0 0 0-25.8m-257.6 134V390.9L628.5 512z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayCircleTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372m164.1 378.2L457.7 677.1a8.02 8.02 0 0 1-12.7-6.5V353a8 8 0 0 1 12.7-6.5l218.4 158.8a7.9 7.9 0 0 1 0 12.9"/><path fill="currentColor" d="M676.1 505.3L457.7 346.5A8 8 0 0 0 445 353v317.6a8.02 8.02 0 0 0 12.7 6.5l218.4-158.9a7.9 7.9 0 0 0 0-12.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaySquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M641.7 520.8L442.3 677.6c-7.4 5.8-18.3.6-18.3-8.8V355.3c0-9.4 10.9-14.7 18.3-8.8l199.4 156.7a11.2 11.2 0 0 1 0 17.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaySquareOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m442.3 677.6l199.4-156.7a11.3 11.3 0 0 0 0-17.7L442.3 346.4c-7.4-5.8-18.3-.6-18.3 8.8v313.5c0 9.4 10.9 14.7 18.3 8.9"/><path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaySquareTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/><path fill="currentColor" fill-opacity=".15" d="M184 840h656V184H184zm240-484.7c0-9.4 10.9-14.7 18.3-8.8l199.4 156.7a11.2 11.2 0 0 1 0 17.6L442.3 677.6c-7.4 5.8-18.3.6-18.3-8.8z"/><path fill="currentColor" d="m442.3 677.6l199.4-156.8a11.2 11.2 0 0 0 0-17.6L442.3 346.5c-7.4-5.9-18.3-.6-18.3 8.8v313.5c0 9.4 10.9 14.6 18.3 8.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m192 472c0 4.4-3.6 8-8 8H544v152c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8V544H328c-4.4 0-8-3.6-8-8v-48c0-4.4 3.6-8 8-8h152V328c0-4.4 3.6-8 8-8h48c4.4 0 8 3.6 8 8v152h152c4.4 0 8 3.6 8 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M696 480H544V328c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v152H328c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h152v152c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V544h152c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8"/><path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircleTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372m192 396c0 4.4-3.6 8-8 8H544v152c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8V544H328c-4.4 0-8-3.6-8-8v-48c0-4.4 3.6-8 8-8h152V328c0-4.4 3.6-8 8-8h48c4.4 0 8 3.6 8 8v152h152c4.4 0 8 3.6 8 8z"/><path fill="currentColor" d="M696 480H544V328c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v152H328c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h152v152c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V544h152c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M482 152h60q8 0 8 8v704q0 8-8 8h-60q-8 0-8-8V160q0-8 8-8"/><path fill="currentColor" d="M192 474h672q8 0 8 8v60q0 8-8 8H160q-8 0-8-8v-60q0-8 8-8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M704 536c0 4.4-3.6 8-8 8H544v152c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8V544H328c-4.4 0-8-3.6-8-8v-48c0-4.4 3.6-8 8-8h152V328c0-4.4 3.6-8 8-8h48c4.4 0 8 3.6 8 8v152h152c4.4 0 8 3.6 8 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSquareOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M328 544h152v152c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V544h152c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8H544V328c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v152H328c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8"/><path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSquareTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/><path fill="currentColor" fill-opacity=".15" d="M184 840h656V184H184zm136-352c0-4.4 3.6-8 8-8h152V328c0-4.4 3.6-8 8-8h48c4.4 0 8 3.6 8 8v152h152c4.4 0 8 3.6 8 8v48c0 4.4-3.6 8-8 8H544v152c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8V544H328c-4.4 0-8-3.6-8-8z"/><path fill="currentColor" d="M328 544h152v152c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V544h152c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8H544V328c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v152H328c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PoundCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m146 658c0 4.4-3.6 8-8 8H376.2c-4.4 0-8-3.6-8-8v-38.5c0-3.7 2.5-6.9 6.1-7.8c44-10.9 72.8-49 72.8-94.2c0-14.7-2.5-29.4-5.9-44.2H374c-4.4 0-8-3.6-8-8v-30c0-4.4 3.6-8 8-8h53.7c-7.8-25.1-14.6-50.7-14.6-77.1c0-75.8 58.6-120.3 151.5-120.3c26.5 0 51.4 5.5 70.3 12.7c3.1 1.2 5.2 4.2 5.2 7.5v39.5a8 8 0 0 1-10.6 7.6c-17.9-6.4-39-10.5-60.4-10.5c-53.3 0-87.3 26.6-87.3 70.2c0 24.7 6.2 47.9 13.4 70.5h112c4.4 0 8 3.6 8 8v30c0 4.4-3.6 8-8 8h-98.6c3.1 13.2 5.3 26.9 5.3 41c0 40.7-16.5 73.9-43.9 91.1v4.7h180c4.4 0 8 3.6 8 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PoundCircleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372m138-209.8H469.8v-4.7c27.4-17.2 43.9-50.4 43.9-91.1c0-14.1-2.2-27.9-5.3-41H607c4.4 0 8-3.6 8-8v-30c0-4.4-3.6-8-8-8H495c-7.2-22.6-13.4-45.7-13.4-70.5c0-43.5 34-70.2 87.3-70.2c21.5 0 42.5 4.1 60.4 10.5c5.2 1.9 10.6-2 10.6-7.6v-39.5c0-3.3-2.1-6.3-5.2-7.5c-18.8-7.2-43.8-12.7-70.3-12.7c-92.9 0-151.5 44.5-151.5 120.3c0 26.3 6.9 52 14.6 77.1H374c-4.4 0-8 3.6-8 8v30c0 4.4 3.6 8 8 8h67.1c3.4 14.7 5.9 29.4 5.9 44.2c0 45.2-28.8 83.3-72.8 94.2c-3.6.9-6.1 4.1-6.1 7.8V722c0 4.4 3.6 8 8 8H650c4.4 0 8-3.6 8-8v-39.8c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PoundCircleTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372m146 582.1c0 4.4-3.6 8-8 8H376.2c-4.4 0-8-3.6-8-8v-38.5c0-3.7 2.5-6.9 6.1-7.8c44-10.9 72.8-49 72.8-94.2c0-14.7-2.5-29.4-5.9-44.2H374c-4.4 0-8-3.6-8-8v-30c0-4.4 3.6-8 8-8h53.7c-7.8-25.1-14.6-50.7-14.6-77.1c0-75.8 58.6-120.3 151.5-120.3c26.5 0 51.4 5.5 70.3 12.7c3.1 1.2 5.2 4.2 5.2 7.5v39.5a8 8 0 0 1-10.6 7.6c-17.9-6.4-39-10.5-60.4-10.5c-53.3 0-87.3 26.6-87.3 70.2c0 24.7 6.2 47.9 13.4 70.5h112c4.4 0 8 3.6 8 8v30c0 4.4-3.6 8-8 8h-98.6c3.1 13.2 5.3 26.9 5.3 41c0 40.7-16.5 73.9-43.9 91.1v4.7h180c4.4 0 8 3.6 8 8z"/><path fill="currentColor" d="M650 674.3H470v-4.7c27.4-17.2 43.9-50.4 43.9-91.1c0-14.1-2.2-27.8-5.3-41h98.6c4.4 0 8-3.6 8-8v-30c0-4.4-3.6-8-8-8h-112c-7.2-22.6-13.4-45.8-13.4-70.5c0-43.6 34-70.2 87.3-70.2c21.4 0 42.5 4.1 60.4 10.5a8 8 0 0 0 10.6-7.6v-39.5c0-3.3-2.1-6.3-5.2-7.5c-18.9-7.2-43.8-12.7-70.3-12.7c-92.9 0-151.5 44.5-151.5 120.3c0 26.4 6.8 52 14.6 77.1H374c-4.4 0-8 3.6-8 8v30c0 4.4 3.6 8 8 8h67.2c3.4 14.8 5.9 29.5 5.9 44.2c0 45.2-28.8 83.3-72.8 94.2c-3.6.9-6.1 4.1-6.1 7.8v38.5c0 4.4 3.6 8 8 8H650c4.4 0 8-3.6 8-8v-39.8c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PoundOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372m138-209.8H469.8v-4.7c27.4-17.2 43.9-50.4 43.9-91.1c0-14.1-2.2-27.9-5.3-41H607c4.4 0 8-3.6 8-8v-30c0-4.4-3.6-8-8-8H495c-7.2-22.6-13.4-45.7-13.4-70.5c0-43.5 34-70.2 87.3-70.2c21.5 0 42.5 4.1 60.4 10.5c5.2 1.9 10.6-2 10.6-7.6v-39.5c0-3.3-2.1-6.3-5.2-7.5c-18.8-7.2-43.8-12.7-70.3-12.7c-92.9 0-151.5 44.5-151.5 120.3c0 26.3 6.9 52 14.6 77.1H374c-4.4 0-8 3.6-8 8v30c0 4.4 3.6 8 8 8h67.1c3.4 14.7 5.9 29.4 5.9 44.2c0 45.2-28.8 83.3-72.8 94.2c-3.6.9-6.1 4.1-6.1 7.8V722c0 4.4 3.6 8 8 8H650c4.4 0 8-3.6 8-8v-39.8c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PoweroffOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M705.6 124.9a8 8 0 0 0-11.6 7.2v64.2c0 5.5 2.9 10.6 7.5 13.6a352.2 352.2 0 0 1 62.2 49.8c32.7 32.8 58.4 70.9 76.3 113.3a355 355 0 0 1 27.9 138.7c0 48.1-9.4 94.8-27.9 138.7a355.92 355.92 0 0 1-76.3 113.3a353.06 353.06 0 0 1-113.2 76.4c-43.8 18.6-90.5 28-138.5 28s-94.7-9.4-138.5-28a353.06 353.06 0 0 1-113.2-76.4A355.92 355.92 0 0 1 184 650.4a355 355 0 0 1-27.9-138.7c0-48.1 9.4-94.8 27.9-138.7c17.9-42.4 43.6-80.5 76.3-113.3c19-19 39.8-35.6 62.2-49.8c4.7-2.9 7.5-8.1 7.5-13.6V132c0-6-6.3-9.8-11.6-7.2C178.5 195.2 82 339.3 80 506.3C77.2 745.1 272.5 943.5 511.2 944c239 .5 432.8-193.3 432.8-432.4c0-169.2-97-315.7-238.4-386.7M480 560h64c4.4 0 8-3.6 8-8V88c0-4.4-3.6-8-8-8h-64c-4.4 0-8 3.6-8 8v464c0 4.4 3.6 8 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PrinterFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M732 120c0-4.4-3.6-8-8-8H300c-4.4 0-8 3.6-8 8v148h440zm120 212H172c-44.2 0-80 35.8-80 80v328c0 17.7 14.3 32 32 32h168v132c0 4.4 3.6 8 8 8h424c4.4 0 8-3.6 8-8V772h168c17.7 0 32-14.3 32-32V412c0-44.2-35.8-80-80-80M664 844H360V568h304zm164-360c0 4.4-3.6 8-8 8h-40c-4.4 0-8-3.6-8-8v-40c0-4.4 3.6-8 8-8h40c4.4 0 8 3.6 8 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PrinterOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M820 436h-40c-4.4 0-8 3.6-8 8v40c0 4.4 3.6 8 8 8h40c4.4 0 8-3.6 8-8v-40c0-4.4-3.6-8-8-8m32-104H732V120c0-4.4-3.6-8-8-8H300c-4.4 0-8 3.6-8 8v212H172c-44.2 0-80 35.8-80 80v328c0 17.7 14.3 32 32 32h168v132c0 4.4 3.6 8 8 8h424c4.4 0 8-3.6 8-8V772h168c17.7 0 32-14.3 32-32V412c0-44.2-35.8-80-80-80M360 180h304v152H360zm304 664H360V568h304zm200-140H732V500H292v204H160V412c0-6.6 5.4-12 12-12h680c6.6 0 12 5.4 12 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PrinterTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M360 180h304v152H360zm492 220H172c-6.6 0-12 5.4-12 12v292h132V500h440v204h132V412c0-6.6-5.4-12-12-12m-24 84c0 4.4-3.6 8-8 8h-40c-4.4 0-8-3.6-8-8v-40c0-4.4 3.6-8 8-8h40c4.4 0 8 3.6 8 8z"/><path fill="currentColor" d="M852 332H732V120c0-4.4-3.6-8-8-8H300c-4.4 0-8 3.6-8 8v212H172c-44.2 0-80 35.8-80 80v328c0 17.7 14.3 32 32 32h168v132c0 4.4 3.6 8 8 8h424c4.4 0 8-3.6 8-8V772h168c17.7 0 32-14.3 32-32V412c0-44.2-35.8-80-80-80M360 180h304v152H360zm304 664H360V568h304zm200-140H732V500H292v204H160V412c0-6.6 5.4-12 12-12h680c6.6 0 12 5.4 12 12z"/><path fill="currentColor" d="M820 436h-40c-4.4 0-8 3.6-8 8v40c0 4.4 3.6 8 8 8h40c4.4 0 8-3.6 8-8v-40c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ProfileFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M380 696c-22.1 0-40-17.9-40-40s17.9-40 40-40s40 17.9 40 40s-17.9 40-40 40m0-144c-22.1 0-40-17.9-40-40s17.9-40 40-40s40 17.9 40 40s-17.9 40-40 40m0-144c-22.1 0-40-17.9-40-40s17.9-40 40-40s40 17.9 40 40s-17.9 40-40 40m304 272c0 4.4-3.6 8-8 8H492c-4.4 0-8-3.6-8-8v-48c0-4.4 3.6-8 8-8h184c4.4 0 8 3.6 8 8zm0-144c0 4.4-3.6 8-8 8H492c-4.4 0-8-3.6-8-8v-48c0-4.4 3.6-8 8-8h184c4.4 0 8 3.6 8 8zm0-144c0 4.4-3.6 8-8 8H492c-4.4 0-8-3.6-8-8v-48c0-4.4 3.6-8 8-8h184c4.4 0 8 3.6 8 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ProfileOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656zM492 400h184c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8H492c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8m0 144h184c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8H492c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8m0 144h184c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8H492c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8M340 368a40 40 0 1 0 80 0a40 40 0 1 0-80 0m0 144a40 40 0 1 0 80 0a40 40 0 1 0-80 0m0 144a40 40 0 1 0 80 0a40 40 0 1 0-80 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ProfileTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/><path fill="currentColor" fill-opacity=".15" d="M184 840h656V184H184zm300-496c0-4.4 3.6-8 8-8h184c4.4 0 8 3.6 8 8v48c0 4.4-3.6 8-8 8H492c-4.4 0-8-3.6-8-8zm0 144c0-4.4 3.6-8 8-8h184c4.4 0 8 3.6 8 8v48c0 4.4-3.6 8-8 8H492c-4.4 0-8-3.6-8-8zm0 144c0-4.4 3.6-8 8-8h184c4.4 0 8 3.6 8 8v48c0 4.4-3.6 8-8 8H492c-4.4 0-8-3.6-8-8zM380 328c22.1 0 40 17.9 40 40s-17.9 40-40 40s-40-17.9-40-40s17.9-40 40-40m0 144c22.1 0 40 17.9 40 40s-17.9 40-40 40s-40-17.9-40-40s17.9-40 40-40m0 144c22.1 0 40 17.9 40 40s-17.9 40-40 40s-40-17.9-40-40s17.9-40 40-40"/><path fill="currentColor" d="M340 656a40 40 0 1 0 80 0a40 40 0 1 0-80 0m0-144a40 40 0 1 0 80 0a40 40 0 1 0-80 0m0-144a40 40 0 1 0 80 0a40 40 0 1 0-80 0m152 320h184c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8H492c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8m0-144h184c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8H492c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8m0-144h184c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8H492c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ProjectFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M368 744c0 4.4-3.6 8-8 8h-80c-4.4 0-8-3.6-8-8V280c0-4.4 3.6-8 8-8h80c4.4 0 8 3.6 8 8zm192-280c0 4.4-3.6 8-8 8h-80c-4.4 0-8-3.6-8-8V280c0-4.4 3.6-8 8-8h80c4.4 0 8 3.6 8 8zm192 72c0 4.4-3.6 8-8 8h-80c-4.4 0-8-3.6-8-8V280c0-4.4 3.6-8 8-8h80c4.4 0 8 3.6 8 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ProjectOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M280 752h80c4.4 0 8-3.6 8-8V280c0-4.4-3.6-8-8-8h-80c-4.4 0-8 3.6-8 8v464c0 4.4 3.6 8 8 8m192-280h80c4.4 0 8-3.6 8-8V280c0-4.4-3.6-8-8-8h-80c-4.4 0-8 3.6-8 8v184c0 4.4 3.6 8 8 8m192 72h80c4.4 0 8-3.6 8-8V280c0-4.4-3.6-8-8-8h-80c-4.4 0-8 3.6-8 8v256c0 4.4 3.6 8 8 8m216-432H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ProjectTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/><path fill="currentColor" fill-opacity=".15" d="M184 840h656V184H184zm472-560c0-4.4 3.6-8 8-8h80c4.4 0 8 3.6 8 8v256c0 4.4-3.6 8-8 8h-80c-4.4 0-8-3.6-8-8zm-192 0c0-4.4 3.6-8 8-8h80c4.4 0 8 3.6 8 8v184c0 4.4-3.6 8-8 8h-80c-4.4 0-8-3.6-8-8zm-192 0c0-4.4 3.6-8 8-8h80c4.4 0 8 3.6 8 8v464c0 4.4-3.6 8-8 8h-80c-4.4 0-8-3.6-8-8z"/><path fill="currentColor" d="M280 752h80c4.4 0 8-3.6 8-8V280c0-4.4-3.6-8-8-8h-80c-4.4 0-8 3.6-8 8v464c0 4.4 3.6 8 8 8m192-280h80c4.4 0 8-3.6 8-8V280c0-4.4-3.6-8-8-8h-80c-4.4 0-8 3.6-8 8v184c0 4.4 3.6 8 8 8m192 72h80c4.4 0 8-3.6 8-8V280c0-4.4-3.6-8-8-8h-80c-4.4 0-8 3.6-8 8v256c0 4.4 3.6 8 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PropertySafetyFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M866.9 169.9L527.1 54.1C523 52.7 517.5 52 512 52s-11 .7-15.1 2.1L157.1 169.9c-8.3 2.8-15.1 12.4-15.1 21.2v482.4c0 8.8 5.7 20.4 12.6 25.9L499.3 968c3.5 2.7 8 4.1 12.6 4.1s9.2-1.4 12.6-4.1l344.7-268.6c6.9-5.4 12.6-17 12.6-25.9V191.1c.2-8.8-6.6-18.3-14.9-21.2M648.3 332.8l-87.7 161.1h45.7c5.5 0 10 4.5 10 10v21.3c0 5.5-4.5 10-10 10h-63.4v29.7h63.4c5.5 0 10 4.5 10 10v21.3c0 5.5-4.5 10-10 10h-63.4V658c0 5.5-4.5 10-10 10h-41.3c-5.5 0-10-4.5-10-10v-51.8h-63.1c-5.5 0-10-4.5-10-10v-21.3c0-5.5 4.5-10 10-10h63.1v-29.7h-63.1c-5.5 0-10-4.5-10-10v-21.3c0-5.5 4.5-10 10-10h45.2l-88-161.1c-2.6-4.8-.9-10.9 4-13.6c1.5-.8 3.1-1.2 4.8-1.2h46c3.8 0 7.2 2.1 8.9 5.5l72.9 144.3l73.2-144.3a10 10 0 0 1 8.9-5.5h45c5.5 0 10 4.5 10 10c.1 1.7-.3 3.3-1.1 4.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PropertySafetyOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M866.9 169.9L527.1 54.1C523 52.7 517.5 52 512 52s-11 .7-15.1 2.1L157.1 169.9c-8.3 2.8-15.1 12.4-15.1 21.2v482.4c0 8.8 5.7 20.4 12.6 25.9L499.3 968c3.5 2.7 8 4.1 12.6 4.1s9.2-1.4 12.6-4.1l344.7-268.6c6.9-5.4 12.6-17 12.6-25.9V191.1c.2-8.8-6.6-18.3-14.9-21.2M810 654.3L512 886.5L214 654.3V226.7l298-101.6l298 101.6zM430.5 318h-46c-1.7 0-3.3.4-4.8 1.2a10.1 10.1 0 0 0-4 13.6l88 161.1h-45.2c-5.5 0-10 4.5-10 10v21.3c0 5.5 4.5 10 10 10h63.1v29.7h-63.1c-5.5 0-10 4.5-10 10v21.3c0 5.5 4.5 10 10 10h63.1V658c0 5.5 4.5 10 10 10h41.3c5.5 0 10-4.5 10-10v-51.8h63.4c5.5 0 10-4.5 10-10v-21.3c0-5.5-4.5-10-10-10h-63.4v-29.7h63.4c5.5 0 10-4.5 10-10v-21.3c0-5.5-4.5-10-10-10h-45.7l87.7-161.1a10.05 10.05 0 0 0-8.8-14.8h-45c-3.8 0-7.2 2.1-8.9 5.5l-73.2 144.3l-72.9-144.3c-1.7-3.4-5.2-5.5-9-5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PropertySafetyTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M866.9 169.9L527.1 54.1C523 52.7 517.5 52 512 52s-11 .7-15.1 2.1L157.1 169.9c-8.3 2.8-15.1 12.4-15.1 21.2v482.4c0 8.8 5.7 20.4 12.6 25.9L499.3 968c3.5 2.7 8 4.1 12.6 4.1s9.2-1.4 12.6-4.1l344.7-268.6c6.9-5.4 12.6-17 12.6-25.9V191.1c.2-8.8-6.6-18.3-14.9-21.2M810 654.3L512 886.5L214 654.3V226.7l298-101.6l298 101.6z"/><path fill="currentColor" fill-opacity=".15" d="M214 226.7v427.6l298 232.2l298-232.2V226.7L512 125.1zM593.9 318h45c5.5 0 10 4.5 10 10c.1 1.7-.3 3.3-1.1 4.8l-87.7 161.1h45.7c5.5 0 10 4.5 10 10v21.3c0 5.5-4.5 10-10 10h-63.4v29.7h63.4c5.5 0 10 4.5 10 10v21.3c0 5.5-4.5 10-10 10h-63.4V658c0 5.5-4.5 10-10 10h-41.3c-5.5 0-10-4.5-10-10v-51.8H418c-5.5 0-10-4.5-10-10v-21.3c0-5.5 4.5-10 10-10h63.1v-29.7H418c-5.5 0-10-4.5-10-10v-21.3c0-5.5 4.5-10 10-10h45.2l-88-161.1c-2.6-4.8-.9-10.9 4-13.6c1.5-.8 3.1-1.2 4.8-1.2h46c3.8 0 7.2 2.1 8.9 5.5l72.9 144.3L585 323.5a10 10 0 0 1 8.9-5.5"/><path fill="currentColor" d="M438.9 323.5a9.88 9.88 0 0 0-8.9-5.5h-46c-1.7 0-3.3.4-4.8 1.2c-4.9 2.7-6.6 8.8-4 13.6l88 161.1H418c-5.5 0-10 4.5-10 10v21.3c0 5.5 4.5 10 10 10h63.1v29.7H418c-5.5 0-10 4.5-10 10v21.3c0 5.5 4.5 10 10 10h63.1V658c0 5.5 4.5 10 10 10h41.3c5.5 0 10-4.5 10-10v-51.8h63.4c5.5 0 10-4.5 10-10v-21.3c0-5.5-4.5-10-10-10h-63.4v-29.7h63.4c5.5 0 10-4.5 10-10v-21.3c0-5.5-4.5-10-10-10h-45.7l87.7-161.1c.8-1.5 1.2-3.1 1.1-4.8c0-5.5-4.5-10-10-10h-45a10 10 0 0 0-8.9 5.5l-73.2 144.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PullRequestOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M788 705.9V192c0-8.8-7.2-16-16-16H602v-68.8c0-6-7-9.4-11.7-5.7L462.7 202.3a7.14 7.14 0 0 0 0 11.3l127.5 100.8c4.7 3.7 11.7.4 11.7-5.7V240h114v465.9c-44.2 15-76 56.9-76 106.1c0 61.8 50.2 112 112 112s112-50.2 112-112c.1-49.2-31.7-91-75.9-106.1M752 860a48.01 48.01 0 0 1 0-96a48.01 48.01 0 0 1 0 96M384 212c0-61.8-50.2-112-112-112s-112 50.2-112 112c0 49.2 31.8 91 76 106.1V706c-44.2 15-76 56.9-76 106.1c0 61.8 50.2 112 112 112s112-50.2 112-112c0-49.2-31.8-91-76-106.1V318.1c44.2-15.1 76-56.9 76-106.1m-160 0a48.01 48.01 0 0 1 96 0a48.01 48.01 0 0 1-96 0m96 600a48.01 48.01 0 0 1-96 0a48.01 48.01 0 0 1 96 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PushpinFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M878.3 392.1L631.9 145.7c-6.5-6.5-15-9.7-23.5-9.7s-17 3.2-23.5 9.7L423.8 306.9c-12.2-1.4-24.5-2-36.8-2c-73.2 0-146.4 24.1-206.5 72.3c-15.4 12.3-16.6 35.4-2.7 49.4l181.7 181.7l-215.4 215.2a15.8 15.8 0 0 0-4.6 9.8l-3.4 37.2c-.9 9.4 6.6 17.4 15.9 17.4c.5 0 1 0 1.5-.1l37.2-3.4c3.7-.3 7.2-2 9.8-4.6l215.4-215.4l181.7 181.7c6.5 6.5 15 9.7 23.5 9.7c9.7 0 19.3-4.2 25.9-12.4c56.3-70.3 79.7-158.3 70.2-243.4l161.1-161.1c12.9-12.8 12.9-33.8 0-46.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PushpinOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M878.3 392.1L631.9 145.7c-6.5-6.5-15-9.7-23.5-9.7s-17 3.2-23.5 9.7L423.8 306.9c-12.2-1.4-24.5-2-36.8-2c-73.2 0-146.4 24.1-206.5 72.3a33.23 33.23 0 0 0-2.7 49.4l181.7 181.7l-215.4 215.2a15.8 15.8 0 0 0-4.6 9.8l-3.4 37.2c-.9 9.4 6.6 17.4 15.9 17.4c.5 0 1 0 1.5-.1l37.2-3.4c3.7-.3 7.2-2 9.8-4.6l215.4-215.4l181.7 181.7c6.5 6.5 15 9.7 23.5 9.7c9.7 0 19.3-4.2 25.9-12.4c56.3-70.3 79.7-158.3 70.2-243.4l161.1-161.1c12.9-12.8 12.9-33.8 0-46.8M666.2 549.3l-24.5 24.5l3.8 34.4a259.92 259.92 0 0 1-30.4 153.9L262 408.8c12.9-7.1 26.3-13.1 40.3-17.9c27.2-9.4 55.7-14.1 84.7-14.1c9.6 0 19.3.5 28.9 1.6l34.4 3.8l24.5-24.5L608.5 224L800 415.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PushpinTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="m474.8 357.7l-24.5 24.5l-34.4-3.8c-9.6-1.1-19.3-1.6-28.9-1.6c-29 0-57.5 4.7-84.7 14.1c-14 4.8-27.4 10.8-40.3 17.9l353.1 353.3a259.92 259.92 0 0 0 30.4-153.9l-3.8-34.4l24.5-24.5L800 415.5L608.5 224z"/><path fill="currentColor" d="M878.3 392.1L631.9 145.7c-6.5-6.5-15-9.7-23.5-9.7s-17 3.2-23.5 9.7L423.8 306.9c-12.2-1.4-24.5-2-36.8-2c-73.2 0-146.4 24.1-206.5 72.3a33.23 33.23 0 0 0-2.7 49.4l181.7 181.7l-215.4 215.2a15.8 15.8 0 0 0-4.6 9.8l-3.4 37.2c-.9 9.4 6.6 17.4 15.9 17.4c.5 0 1 0 1.5-.1l37.2-3.4c3.7-.3 7.2-2 9.8-4.6l215.4-215.4l181.7 181.7c6.5 6.5 15 9.7 23.5 9.7c9.7 0 19.3-4.2 25.9-12.4c56.3-70.3 79.7-158.3 70.2-243.4l161.1-161.1c12.9-12.8 12.9-33.8 0-46.8M666.2 549.3l-24.5 24.5l3.8 34.4a259.92 259.92 0 0 1-30.4 153.9L262 408.8c12.9-7.1 26.3-13.1 40.3-17.9c27.2-9.4 55.7-14.1 84.7-14.1c9.6 0 19.3.5 28.9 1.6l34.4 3.8l24.5-24.5L608.5 224L800 415.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QqCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m210.5 612.4c-11.5 1.4-44.9-52.7-44.9-52.7c0 31.3-16.2 72.2-51.1 101.8c16.9 5.2 54.9 19.2 45.9 34.4c-7.3 12.3-125.6 7.9-159.8 4c-34.2 3.8-152.5 8.3-159.8-4c-9.1-15.2 28.9-29.2 45.8-34.4c-35-29.5-51.1-70.4-51.1-101.8c0 0-33.4 54.1-44.9 52.7c-5.4-.7-12.4-29.6 9.4-99.7c10.3-33 22-60.5 40.2-105.8c-3.1-116.9 45.3-215 160.4-215c113.9 0 163.3 96.1 160.4 215c18.1 45.2 29.9 72.8 40.2 105.8c21.7 70.1 14.6 99.1 9.3 99.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QqOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M824.8 613.2c-16-51.4-34.4-94.6-62.7-165.3C766.5 262.2 689.3 112 511.5 112C331.7 112 256.2 265.2 261 447.9c-28.4 70.8-46.7 113.7-62.7 165.3c-34 109.5-23 154.8-14.6 155.8c18 2.2 70.1-82.4 70.1-82.4c0 49 25.2 112.9 79.8 159c-26.4 8.1-85.7 29.9-71.6 53.8c11.4 19.3 196.2 12.3 249.5 6.3c53.3 6 238.1 13 249.5-6.3c14.1-23.8-45.3-45.7-71.6-53.8c54.6-46.2 79.8-110.1 79.8-159c0 0 52.1 84.6 70.1 82.4c8.5-1.1 19.5-46.4-14.5-155.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QqSquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M722.5 676.4c-11.5 1.4-44.9-52.7-44.9-52.7c0 31.3-16.2 72.2-51.1 101.8c16.9 5.2 54.9 19.2 45.9 34.4c-7.3 12.3-125.6 7.9-159.8 4c-34.2 3.8-152.5 8.3-159.8-4c-9.1-15.2 28.9-29.2 45.8-34.4c-35-29.5-51.1-70.4-51.1-101.8c0 0-33.4 54.1-44.9 52.7c-5.4-.7-12.4-29.6 9.4-99.7c10.3-33 22-60.5 40.2-105.8c-3.1-116.9 45.3-215 160.4-215c113.9 0 163.3 96.1 160.4 215c18.1 45.2 29.9 72.8 40.2 105.8c21.7 70.1 14.6 99.1 9.3 99.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QrcodeOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M468 128H160c-17.7 0-32 14.3-32 32v308c0 4.4 3.6 8 8 8h332c4.4 0 8-3.6 8-8V136c0-4.4-3.6-8-8-8m-56 284H192V192h220zm-138-74h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m194 210H136c-4.4 0-8 3.6-8 8v308c0 17.7 14.3 32 32 32h308c4.4 0 8-3.6 8-8V556c0-4.4-3.6-8-8-8m-56 284H192V612h220zm-138-74h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m590-630H556c-4.4 0-8 3.6-8 8v332c0 4.4 3.6 8 8 8h332c4.4 0 8-3.6 8-8V160c0-17.7-14.3-32-32-32m-32 284H612V192h220zm-138-74h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m194 210h-48c-4.4 0-8 3.6-8 8v134h-78V556c0-4.4-3.6-8-8-8H556c-4.4 0-8 3.6-8 8v332c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V644h78v102c0 4.4 3.6 8 8 8h190c4.4 0 8-3.6 8-8V556c0-4.4-3.6-8-8-8M746 832h-48c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8m142 0h-48c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 708c-22.1 0-40-17.9-40-40s17.9-40 40-40s40 17.9 40 40s-17.9 40-40 40m62.9-219.5a48.3 48.3 0 0 0-30.9 44.8V620c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8v-21.5c0-23.1 6.7-45.9 19.9-64.9c12.9-18.6 30.9-32.8 52.1-40.9c34-13.1 56-41.6 56-72.7c0-44.1-43.1-80-96-80s-96 35.9-96 80v7.6c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8V420c0-39.3 17.2-76 48.4-103.3C430.4 290.4 470 276 512 276s81.6 14.5 111.6 40.7C654.8 344 672 380.7 672 420c0 57.8-38.1 109.8-97.1 132.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionCircleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" d="M623.6 316.7C593.6 290.4 554 276 512 276s-81.6 14.5-111.6 40.7C369.2 344 352 380.7 352 420v7.6c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V420c0-44.1 43.1-80 96-80s96 35.9 96 80c0 31.1-22 59.6-56.1 72.7c-21.2 8.1-39.2 22.3-52.1 40.9c-13.1 19-19.9 41.8-19.9 64.9V620c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8v-22.7a48.3 48.3 0 0 1 30.9-44.8c59-22.7 97.1-74.7 97.1-132.5c.1-39.3-17.1-76-48.3-103.3M472 732a40 40 0 1 0 80 0a40 40 0 1 0-80 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionCircleTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372m0 632c-22.1 0-40-17.9-40-40s17.9-40 40-40s40 17.9 40 40s-17.9 40-40 40m62.9-219.5a48.3 48.3 0 0 0-30.9 44.8V620c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8v-21.5c0-23.1 6.7-45.9 19.9-64.9c12.9-18.6 30.9-32.8 52.1-40.9c34-13.1 56-41.6 56-72.7c0-44.1-43.1-80-96-80s-96 35.9-96 80v7.6c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8V420c0-39.3 17.2-76 48.4-103.3C430.4 290.4 470 276 512 276s81.6 14.5 111.6 40.7C654.8 344 672 380.7 672 420c0 57.8-38.1 109.8-97.1 132.5"/><path fill="currentColor" d="M472 732a40 40 0 1 0 80 0a40 40 0 1 0-80 0m151.6-415.3C593.6 290.5 554 276 512 276s-81.6 14.4-111.6 40.7C369.2 344 352 380.7 352 420v7.6c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V420c0-44.1 43.1-80 96-80s96 35.9 96 80c0 31.1-22 59.6-56 72.7c-21.2 8.1-39.2 22.3-52.1 40.9c-13.2 19-19.9 41.8-19.9 64.9V620c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8v-22.7a48.3 48.3 0 0 1 30.9-44.8c59-22.7 97.1-74.7 97.1-132.5c0-39.3-17.2-76-48.4-103.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M764 280.9c-14-30.6-33.9-58.1-59.3-81.6C653.1 151.4 584.6 125 512 125s-141.1 26.4-192.7 74.2c-25.4 23.6-45.3 51-59.3 81.7c-14.6 32-22 65.9-22 100.9v27c0 6.2 5 11.2 11.2 11.2h54c6.2 0 11.2-5 11.2-11.2v-27c0-99.5 88.6-180.4 197.6-180.4s197.6 80.9 197.6 180.4c0 40.8-14.5 79.2-42 111.2c-27.2 31.7-65.6 54.4-108.1 64c-24.3 5.5-46.2 19.2-61.7 38.8a110.85 110.85 0 0 0-23.9 68.6v31.4c0 6.2 5 11.2 11.2 11.2h54c6.2 0 11.2-5 11.2-11.2v-31.4c0-15.7 10.9-29.5 26-32.9c58.4-13.2 111.4-44.7 149.3-88.7c19.1-22.3 34-47.1 44.3-74c10.7-27.9 16.1-57.2 16.1-87c0-35-7.4-69-22-100.9M512 787c-30.9 0-56 25.1-56 56s25.1 56 56 56s56-25.1 56-56s-25.1-56-56-56"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadarChartOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m926.8 397.1l-396-288a31.81 31.81 0 0 0-37.6 0l-396 288a31.99 31.99 0 0 0-11.6 35.8l151.3 466a32 32 0 0 0 30.4 22.1h489.5c13.9 0 26.1-8.9 30.4-22.1l151.3-466c4.2-13.2-.5-27.6-11.7-35.8M838.6 417l-98.5 32l-200-144.7V199.9zM466 567.2l-89.1 122.3l-55.2-169.2zm-116.3-96.8L484 373.3v140.8zM512 599.2l93.9 128.9H418.1zm28.1-225.9l134.2 97.1L540.1 514zM558 567.2l144.3-46.9l-55.2 169.2zm-74-367.3v104.4L283.9 449l-98.5-32zM169.3 470.8l86.5 28.1l80.4 246.4l-53.8 73.9zM327.1 853l50.3-69h269.3l50.3 69zm414.5-33.8l-53.8-73.9l80.4-246.4l86.5-28.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadiusBottomleftOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M712 824h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m2-696h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8M136 374h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m0-174h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m752 624h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-174h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-174h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-174h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-174h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-348 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-230 72h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m230 624H358c-87.3 0-158-70.7-158-158V484c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v182c0 127 103 230 230 230h182c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadiusBottomrightOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M368 824h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-58-624h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m578 102h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-174h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8M192 824h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-174h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-174h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-174h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0-174h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m292 72h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m174 0h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m230 276h-56c-4.4 0-8 3.6-8 8v182c0 87.3-70.7 158-158 158H484c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h182c127 0 230-103 230-230V484c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadiusSettingOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M396 140h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-44 684h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m524-204h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8M192 344h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 160h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 160h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 160h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m320 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m160 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m140-284c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V370c0-127-103-230-230-230H484c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h170c87.3 0 158 70.7 158 158zM236 96H92c-4.4 0-8 3.6-8 8v144c0 4.4 3.6 8 8 8h144c4.4 0 8-3.6 8-8V104c0-4.4-3.6-8-8-8m-48 101.6c0 1.3-1.1 2.4-2.4 2.4h-43.2c-1.3 0-2.4-1.1-2.4-2.4v-43.2c0-1.3 1.1-2.4 2.4-2.4h43.2c1.3 0 2.4 1.1 2.4 2.4zM920 780H776c-4.4 0-8 3.6-8 8v144c0 4.4 3.6 8 8 8h144c4.4 0 8-3.6 8-8V788c0-4.4-3.6-8-8-8m-48 101.6c0 1.3-1.1 2.4-2.4 2.4h-43.2c-1.3 0-2.4-1.1-2.4-2.4v-43.2c0-1.3 1.1-2.4 2.4-2.4h43.2c1.3 0 2.4 1.1 2.4 2.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadiusUpleftOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M656 200h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m58 624h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8M192 650h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 174h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m696-696h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 174h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 174h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 174h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 174h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-348 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-174 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m174-696H358c-127 0-230 103-230 230v182c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V358c0-87.3 70.7-158 158-158h182c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadiusUprightOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M368 128h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-2 696h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m522-174h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 174h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8M192 128h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 174h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 174h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 174h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 174h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m348 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m174 0h-56c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m-48-696H484c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h182c87.3 0 158 70.7 158 158v182c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V358c0-127-103-230-230-230"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReadFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M928 161H699.2c-49.1 0-97.1 14.1-138.4 40.7L512 233l-48.8-31.3A255.2 255.2 0 0 0 324.8 161H96c-17.7 0-32 14.3-32 32v568c0 17.7 14.3 32 32 32h228.8c49.1 0 97.1 14.1 138.4 40.7l44.4 28.6c1.3.8 2.8 1.3 4.3 1.3s3-.4 4.3-1.3l44.4-28.6C602 807.1 650.1 793 699.2 793H928c17.7 0 32-14.3 32-32V193c0-17.7-14.3-32-32-32M404 553.5c0 4.1-3.2 7.5-7.1 7.5H211.1c-3.9 0-7.1-3.4-7.1-7.5v-45c0-4.1 3.2-7.5 7.1-7.5h185.7c3.9 0 7.1 3.4 7.1 7.5v45zm0-140c0 4.1-3.2 7.5-7.1 7.5H211.1c-3.9 0-7.1-3.4-7.1-7.5v-45c0-4.1 3.2-7.5 7.1-7.5h185.7c3.9 0 7.1 3.4 7.1 7.5v45zm416 140c0 4.1-3.2 7.5-7.1 7.5H627.1c-3.9 0-7.1-3.4-7.1-7.5v-45c0-4.1 3.2-7.5 7.1-7.5h185.7c3.9 0 7.1 3.4 7.1 7.5v45zm0-140c0 4.1-3.2 7.5-7.1 7.5H627.1c-3.9 0-7.1-3.4-7.1-7.5v-45c0-4.1 3.2-7.5 7.1-7.5h185.7c3.9 0 7.1 3.4 7.1 7.5v45z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReadOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M928 161H699.2c-49.1 0-97.1 14.1-138.4 40.7L512 233l-48.8-31.3A255.2 255.2 0 0 0 324.8 161H96c-17.7 0-32 14.3-32 32v568c0 17.7 14.3 32 32 32h228.8c49.1 0 97.1 14.1 138.4 40.7l44.4 28.6c1.3.8 2.8 1.3 4.3 1.3s3-.4 4.3-1.3l44.4-28.6C602 807.1 650.1 793 699.2 793H928c17.7 0 32-14.3 32-32V193c0-17.7-14.3-32-32-32M324.8 721H136V233h188.8c35.4 0 69.8 10.1 99.5 29.2l48.8 31.3l6.9 4.5v462c-47.6-25.6-100.8-39-155.2-39m563.2 0H699.2c-54.4 0-107.6 13.4-155.2 39V298l6.9-4.5l48.8-31.3c29.7-19.1 64.1-29.2 99.5-29.2H888zM396.9 361H211.1c-3.9 0-7.1 3.4-7.1 7.5v45c0 4.1 3.2 7.5 7.1 7.5h185.7c3.9 0 7.1-3.4 7.1-7.5v-45c.1-4.1-3.1-7.5-7-7.5m223.1 7.5v45c0 4.1 3.2 7.5 7.1 7.5h185.7c3.9 0 7.1-3.4 7.1-7.5v-45c0-4.1-3.2-7.5-7.1-7.5H627.1c-3.9 0-7.1 3.4-7.1 7.5M396.9 501H211.1c-3.9 0-7.1 3.4-7.1 7.5v45c0 4.1 3.2 7.5 7.1 7.5h185.7c3.9 0 7.1-3.4 7.1-7.5v-45c.1-4.1-3.1-7.5-7-7.5m416 0H627.1c-3.9 0-7.1 3.4-7.1 7.5v45c0 4.1 3.2 7.5 7.1 7.5h185.7c3.9 0 7.1-3.4 7.1-7.5v-45c.1-4.1-3.1-7.5-7-7.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReconciliationFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M676 623c-18.8 0-34 15.2-34 34s15.2 34 34 34s34-15.2 34-34s-15.2-34-34-34m204-455H668c0-30.9-25.1-56-56-56h-80c-30.9 0-56 25.1-56 56H264c-17.7 0-32 14.3-32 32v200h-88c-17.7 0-32 14.3-32 32v448c0 17.7 14.3 32 32 32h336c17.7 0 32-14.3 32-32v-16h368c17.7 0 32-14.3 32-32V200c0-17.7-14.3-32-32-32M448 848H176V616h272zm0-296H176v-88h272zm20-272v-48h72v-56h64v56h72v48zm180 168v56c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8v-56c0-4.4 3.6-8 8-8h48c4.4 0 8 3.6 8 8m28 301c-50.8 0-92-41.2-92-92s41.2-92 92-92s92 41.2 92 92s-41.2 92-92 92m92-245c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8v-96c0-4.4 3.6-8 8-8h48c4.4 0 8 3.6 8 8zm-92 61c-50.8 0-92 41.2-92 92s41.2 92 92 92s92-41.2 92-92s-41.2-92-92-92m0 126c-18.8 0-34-15.2-34-34s15.2-34 34-34s34 15.2 34 34s-15.2 34-34 34"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReconciliationOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M676 565c-50.8 0-92 41.2-92 92s41.2 92 92 92s92-41.2 92-92s-41.2-92-92-92m0 126c-18.8 0-34-15.2-34-34s15.2-34 34-34s34 15.2 34 34s-15.2 34-34 34m204-523H668c0-30.9-25.1-56-56-56h-80c-30.9 0-56 25.1-56 56H264c-17.7 0-32 14.3-32 32v200h-88c-17.7 0-32 14.3-32 32v448c0 17.7 14.3 32 32 32h336c17.7 0 32-14.3 32-32v-16h368c17.7 0 32-14.3 32-32V200c0-17.7-14.3-32-32-32m-412 64h72v-56h64v56h72v48H468zm-20 616H176V616h272zm0-296H176v-88h272zm392 240H512V432c0-17.7-14.3-32-32-32H304V240h100v104h336V240h100zM704 408v96c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8v-96c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8M592 512h48c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReconciliationTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M740 344H404V240H304v160h176c17.7 0 32 14.3 32 32v360h328V240H740zM584 448c0-4.4 3.6-8 8-8h48c4.4 0 8 3.6 8 8v56c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8zm92 301c-50.8 0-92-41.2-92-92s41.2-92 92-92s92 41.2 92 92s-41.2 92-92 92m92-341v96c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8v-96c0-4.4 3.6-8 8-8h48c4.4 0 8 3.6 8 8"/><path fill="currentColor" fill-opacity=".15" d="M642 657a34 34 0 1 0 68 0a34 34 0 1 0-68 0"/><path fill="currentColor" d="M592 512h48c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8m112-104v96c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8v-96c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8"/><path fill="currentColor" d="M880 168H668c0-30.9-25.1-56-56-56h-80c-30.9 0-56 25.1-56 56H264c-17.7 0-32 14.3-32 32v200h-88c-17.7 0-32 14.3-32 32v448c0 17.7 14.3 32 32 32h336c17.7 0 32-14.3 32-32v-16h368c17.7 0 32-14.3 32-32V200c0-17.7-14.3-32-32-32m-412 64h72v-56h64v56h72v48H468zm-20 616H176V616h272zm0-296H176v-88h272zm392 240H512V432c0-17.7-14.3-32-32-32H304V240h100v104h336V240h100z"/><path fill="currentColor" d="M676 565c-50.8 0-92 41.2-92 92s41.2 92 92 92s92-41.2 92-92s-41.2-92-92-92m0 126c-18.8 0-34-15.2-34-34s15.2-34 34-34s34 15.2 34 34s-15.2 34-34 34"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedEnvelopeFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 64H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V96c0-17.7-14.3-32-32-32M647 470.4l-87.2 161h45.9c4.6 0 8.4 3.8 8.4 8.4v25.1c0 4.6-3.8 8.4-8.4 8.4h-63.3v28.6h63.3c4.6 0 8.4 3.8 8.4 8.4v25c.2 4.6-3.6 8.5-8.2 8.5h-63.3v49.9c0 4.6-3.8 8.4-8.4 8.4h-43.7c-4.6 0-8.4-3.8-8.4-8.4v-49.9h-63c-4.6 0-8.4-3.8-8.4-8.4v-25.1c0-4.6 3.8-8.4 8.4-8.4h63v-28.6h-63c-4.6 0-8.4-3.8-8.4-8.4v-25.1c0-4.6 3.8-8.4 8.4-8.4h45.4l-87.5-161c-2.2-4.1-.7-9.1 3.4-11.4c1.3-.6 2.6-1 3.9-1h48.8c3.2 0 6.1 1.8 7.5 4.6l71.9 141.8l71.9-141.9a8.5 8.5 0 0 1 7.5-4.6h47.8c4.6 0 8.4 3.8 8.4 8.4c-.1 1.5-.5 2.9-1.1 4.1M512.6 323L289 148h446z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedEnvelopeOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M440.6 462.6a8.38 8.38 0 0 0-7.5-4.6h-48.8c-1.3 0-2.6.4-3.9 1a8.4 8.4 0 0 0-3.4 11.4l87.4 161.1H419c-4.6 0-8.4 3.8-8.4 8.4V665c0 4.6 3.8 8.4 8.4 8.4h63V702h-63c-4.6 0-8.4 3.8-8.4 8.4v25.1c0 4.6 3.8 8.4 8.4 8.4h63v49.9c0 4.6 3.8 8.4 8.4 8.4h43.7c4.6 0 8.4-3.8 8.4-8.4v-49.9h63.3c4.7 0 8.4-3.8 8.2-8.5v-25c0-4.6-3.8-8.4-8.4-8.4h-63.3v-28.6h63.3c4.6 0 8.4-3.8 8.4-8.4v-25.1c0-4.6-3.8-8.4-8.4-8.4h-45.9l87.2-161a8.45 8.45 0 0 0-7.4-12.4h-47.8c-3.1 0-6 1.8-7.5 4.6l-71.9 141.9zM832 64H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V96c0-17.7-14.3-32-32-32m-40 824H232V193.1l260.3 204.1c11.6 9.1 27.9 9.1 39.5 0L792 193.1zm0-751.3h-31.7L512 331.3L263.7 136.7H232v-.7h560z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedEnvelopeTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 64H192c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V96c0-17.7-14.3-32-32-32m-40 824H232V193.1l260.3 204.1c11.6 9.1 27.9 9.1 39.5 0L792 193.1zm0-751.3h-31.7L512 331.3L263.7 136.7H232v-.7h560z"/><path fill="currentColor" fill-opacity=".15" d="M492.3 397.2L232 193.1V888h560V193.1L531.8 397.2a31.99 31.99 0 0 1-39.5 0m99.4 60.9h47.8a8.45 8.45 0 0 1 7.4 12.4l-87.2 161h45.9c4.6 0 8.4 3.8 8.4 8.4V665c0 4.6-3.8 8.4-8.4 8.4h-63.3V702h63.3c4.6 0 8.4 3.8 8.4 8.4v25c.2 4.7-3.5 8.5-8.2 8.5h-63.3v49.9c0 4.6-3.8 8.4-8.4 8.4h-43.7c-4.6 0-8.4-3.8-8.4-8.4v-49.9h-63c-4.6 0-8.4-3.8-8.4-8.4v-25.1c0-4.6 3.8-8.4 8.4-8.4h63v-28.6h-63c-4.6 0-8.4-3.8-8.4-8.4v-25.1c0-4.6 3.8-8.4 8.4-8.4h45.4L377 470.4a8.4 8.4 0 0 1 3.4-11.4c1.3-.6 2.6-1 3.9-1h48.8c3.2 0 6.1 1.8 7.5 4.6l71.7 142l71.9-141.9a8.6 8.6 0 0 1 7.5-4.6"/><path fill="currentColor" fill-opacity=".15" d="M232 136.7h31.7L512 331.3l248.3-194.6H792v-.7H232z"/><path fill="currentColor" d="M440.6 462.6a8.38 8.38 0 0 0-7.5-4.6h-48.8c-1.3 0-2.6.4-3.9 1a8.4 8.4 0 0 0-3.4 11.4l87.4 161.1H419c-4.6 0-8.4 3.8-8.4 8.4V665c0 4.6 3.8 8.4 8.4 8.4h63V702h-63c-4.6 0-8.4 3.8-8.4 8.4v25.1c0 4.6 3.8 8.4 8.4 8.4h63v49.9c0 4.6 3.8 8.4 8.4 8.4h43.7c4.6 0 8.4-3.8 8.4-8.4v-49.9h63.3c4.7 0 8.4-3.8 8.2-8.5v-25c0-4.6-3.8-8.4-8.4-8.4h-63.3v-28.6h63.3c4.6 0 8.4-3.8 8.4-8.4v-25.1c0-4.6-3.8-8.4-8.4-8.4h-45.9l87.2-161a8.45 8.45 0 0 0-7.4-12.4h-47.8c-3.1 0-6 1.8-7.5 4.6l-71.9 141.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedditCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M584 548a36 36 0 1 0 72 0a36 36 0 1 0-72 0m144-108a35.9 35.9 0 0 0-32.5 20.6c18.8 14.3 34.4 30.7 45.9 48.8A35.98 35.98 0 0 0 728 440M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m245 477.9c4.6 13.5 7 27.6 7 42.1c0 99.4-112.8 180-252 180s-252-80.6-252-180c0-14.5 2.4-28.6 7-42.1A72.01 72.01 0 0 1 296 404c27.1 0 50.6 14.9 62.9 37c36.2-19.8 80.2-32.8 128.1-36.1l58.4-131.1c4.3-9.8 15.2-14.8 25.5-11.8l91.6 26.5a54.03 54.03 0 0 1 101.6 25.6c0 29.8-24.2 54-54 54c-23.5 0-43.5-15.1-50.9-36.1L577 308.3l-43 96.5c49.1 3 94.2 16.1 131.2 36.3c12.3-22.1 35.8-37 62.9-37c39.8 0 72 32.2 72 72c-.1 29.3-17.8 54.6-43.1 65.8m-171.3 83c-14.9 11.7-44.3 24.3-73.7 24.3s-58.9-12.6-73.7-24.3c-9.3-7.3-22.7-5.7-30 3.6c-7.3 9.3-5.7 22.7 3.6 30c25.7 20.4 65 33.5 100.1 33.5c35.1 0 74.4-13.1 100.2-33.5c9.3-7.3 10.9-20.8 3.6-30a21.46 21.46 0 0 0-30.1-3.6M296 440a35.98 35.98 0 0 0-13.4 69.4c11.5-18.1 27.1-34.5 45.9-48.8A35.9 35.9 0 0 0 296 440m72 108a36 36 0 1 0 72 0a36 36 0 1 0-72 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedditOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M288 568a56 56 0 1 0 112 0a56 56 0 1 0-112 0m338.7 119.7c-23.1 18.2-68.9 37.8-114.7 37.8s-91.6-19.6-114.7-37.8c-14.4-11.3-35.3-8.9-46.7 5.5s-8.9 35.3 5.5 46.7C396.3 771.6 457.5 792 512 792s115.7-20.4 155.9-52.1a33.25 33.25 0 1 0-41.2-52.2M960 456c0-61.9-50.1-112-112-112c-42.1 0-78.7 23.2-97.9 57.6c-57.6-31.5-127.7-51.8-204.1-56.5L612.9 195l127.9 36.9c11.5 32.6 42.6 56.1 79.2 56.1c46.4 0 84-37.6 84-84s-37.6-84-84-84c-32 0-59.8 17.9-74 44.2L603.5 123a33.2 33.2 0 0 0-39.6 18.4l-90.8 203.9c-74.5 5.2-142.9 25.4-199.2 56.2A111.94 111.94 0 0 0 176 344c-61.9 0-112 50.1-112 112c0 45.8 27.5 85.1 66.8 102.5c-7.1 21-10.8 43-10.8 65.5c0 154.6 175.5 280 392 280s392-125.4 392-280c0-22.6-3.8-44.5-10.8-65.5C932.5 541.1 960 501.8 960 456M820 172.5a31.5 31.5 0 1 1 0 63a31.5 31.5 0 0 1 0-63M120 456c0-30.9 25.1-56 56-56a56 56 0 0 1 50.6 32.1c-29.3 22.2-53.5 47.8-71.5 75.9a56.23 56.23 0 0 1-35.1-52m392 381.5c-179.8 0-325.5-95.6-325.5-213.5S332.2 410.5 512 410.5S837.5 506.1 837.5 624S691.8 837.5 512 837.5M868.8 508c-17.9-28.1-42.2-53.7-71.5-75.9c9-18.9 28.3-32.1 50.6-32.1c30.9 0 56 25.1 56 56c.1 23.5-14.5 43.7-35.1 52M624 568a56 56 0 1 0 112 0a56 56 0 1 0-112 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedditSquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M296 440a35.98 35.98 0 0 0-13.4 69.4c11.5-18.1 27.1-34.5 45.9-48.8A35.9 35.9 0 0 0 296 440m289.7 184.9c-14.9 11.7-44.3 24.3-73.7 24.3s-58.9-12.6-73.7-24.3c-9.3-7.3-22.7-5.7-30 3.6c-7.3 9.3-5.7 22.7 3.6 30c25.7 20.4 65 33.5 100.1 33.5c35.1 0 74.4-13.1 100.2-33.5c9.3-7.3 10.9-20.8 3.6-30a21.46 21.46 0 0 0-30.1-3.6M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M757 541.9c4.6 13.5 7 27.6 7 42.1c0 99.4-112.8 180-252 180s-252-80.6-252-180c0-14.5 2.4-28.6 7-42.1A72.01 72.01 0 0 1 296 404c27.1 0 50.6 14.9 62.9 37c36.2-19.8 80.2-32.8 128.1-36.1l58.4-131.1c4.3-9.8 15.2-14.8 25.5-11.8l91.6 26.5a54.03 54.03 0 0 1 101.6 25.6c0 29.8-24.2 54-54 54c-23.5 0-43.5-15.1-50.9-36.1L577 308.3l-43 96.5c49.1 3 94.2 16.1 131.2 36.3c12.3-22.1 35.8-37 62.9-37c39.8 0 72 32.2 72 72c-.1 29.3-17.8 54.6-43.1 65.8M584 548a36 36 0 1 0 72 0a36 36 0 1 0-72 0m144-108a35.9 35.9 0 0 0-32.5 20.6c18.8 14.3 34.4 30.7 45.9 48.8A35.98 35.98 0 0 0 728 440M368 548a36 36 0 1 0 72 0a36 36 0 1 0-72 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedoOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M758.2 839.1C851.8 765.9 912 651.9 912 523.9C912 303 733.5 124.3 512.6 124C291.4 123.7 112 302.8 112 523.9c0 125.2 57.5 236.9 147.6 310.2c3.5 2.8 8.6 2.2 11.4-1.3l39.4-50.5c2.7-3.4 2.1-8.3-1.2-11.1c-8.1-6.6-15.9-13.7-23.4-21.2a318.64 318.64 0 0 1-68.6-101.7C200.4 609 192 567.1 192 523.9s8.4-85.1 25.1-124.5c16.1-38.1 39.2-72.3 68.6-101.7c29.4-29.4 63.6-52.5 101.7-68.6C426.9 212.4 468.8 204 512 204s85.1 8.4 124.5 25.1c38.1 16.1 72.3 39.2 101.7 68.6c29.4 29.4 52.5 63.6 68.6 101.7c16.7 39.4 25.1 81.3 25.1 124.5s-8.4 85.1-25.1 124.5a318.64 318.64 0 0 1-68.6 101.7c-9.3 9.3-19.1 18-29.3 26L668.2 724a8 8 0 0 0-14.1 3l-39.6 162.2c-1.2 5 2.6 9.9 7.7 9.9l167 .8c6.7 0 10.5-7.7 6.3-12.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReloadOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m909.1 209.3l-56.4 44.1C775.8 155.1 656.2 92 521.9 92C290 92 102.3 279.5 102 511.5C101.7 743.7 289.8 932 521.9 932c181.3 0 335.8-115 394.6-276.1c1.5-4.2-.7-8.9-4.9-10.3l-56.7-19.5a8 8 0 0 0-10.1 4.8c-1.8 5-3.8 10-5.9 14.9c-17.3 41-42.1 77.8-73.7 109.4A344.77 344.77 0 0 1 655.9 829c-42.3 17.9-87.4 27-133.8 27c-46.5 0-91.5-9.1-133.8-27A341.5 341.5 0 0 1 279 755.2a342.16 342.16 0 0 1-73.7-109.4c-17.9-42.4-27-87.4-27-133.9s9.1-91.5 27-133.9c17.3-41 42.1-77.8 73.7-109.4c31.6-31.6 68.4-56.4 109.3-73.8c42.3-17.9 87.4-27 133.8-27c46.5 0 91.5 9.1 133.8 27a341.5 341.5 0 0 1 109.3 73.8c9.9 9.9 19.2 20.4 27.8 31.4l-60.2 47a8 8 0 0 0 3 14.1l175.6 43c5 1.2 9.9-2.6 9.9-7.7l.8-180.9c-.1-6.6-7.8-10.3-13-6.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RestFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 256h-28.1l-35.7-120.9c-4-13.7-16.5-23.1-30.7-23.1h-451c-14.3 0-26.8 9.4-30.7 23.1L220.1 256H192c-17.7 0-32 14.3-32 32v28c0 4.4 3.6 8 8 8h45.8l47.7 558.7a32 32 0 0 0 31.9 29.3h429.2a32 32 0 0 0 31.9-29.3L802.2 324H856c4.4 0 8-3.6 8-8v-28c0-17.7-14.3-32-32-32M508 704c-79.5 0-144-64.5-144-144s64.5-144 144-144s144 64.5 144 144s-64.5 144-144 144M291 256l22.4-76h397.2l22.4 76zm137 304a80 80 0 1 0 160 0a80 80 0 1 0-160 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RestOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M508 704c79.5 0 144-64.5 144-144s-64.5-144-144-144s-144 64.5-144 144s64.5 144 144 144m0-224c44.2 0 80 35.8 80 80s-35.8 80-80 80s-80-35.8-80-80s35.8-80 80-80"/><path fill="currentColor" d="M832 256h-28.1l-35.7-120.9c-4-13.7-16.5-23.1-30.7-23.1h-451c-14.3 0-26.8 9.4-30.7 23.1L220.1 256H192c-17.7 0-32 14.3-32 32v28c0 4.4 3.6 8 8 8h45.8l47.7 558.7a32 32 0 0 0 31.9 29.3h429.2a32 32 0 0 0 31.9-29.3L802.2 324H856c4.4 0 8-3.6 8-8v-28c0-17.7-14.3-32-32-32m-518.6-76h397.2l22.4 76H291zm376.2 664H326.4L282 324h451.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RestTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M326.4 844h363.2l44.3-520H282zM508 416c79.5 0 144 64.5 144 144s-64.5 144-144 144s-144-64.5-144-144s64.5-144 144-144"/><path fill="currentColor" d="M508 704c79.5 0 144-64.5 144-144s-64.5-144-144-144s-144 64.5-144 144s64.5 144 144 144m0-224c44.2 0 80 35.8 80 80s-35.8 80-80 80s-80-35.8-80-80s35.8-80 80-80"/><path fill="currentColor" d="M832 256h-28.1l-35.7-120.9c-4-13.7-16.5-23.1-30.7-23.1h-451c-14.3 0-26.8 9.4-30.7 23.1L220.1 256H192c-17.7 0-32 14.3-32 32v28c0 4.4 3.6 8 8 8h45.8l47.7 558.7a32 32 0 0 0 31.9 29.3h429.2a32 32 0 0 0 31.9-29.3L802.2 324H856c4.4 0 8-3.6 8-8v-28c0-17.7-14.3-32-32-32m-518.6-76h397.2l22.4 76H291zm376.2 664H326.4L282 324h451.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RetweetOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M136 552h63.6c4.4 0 8-3.6 8-8V288.7h528.6v72.6c0 1.9.6 3.7 1.8 5.2a8.3 8.3 0 0 0 11.7 1.4L893 255.4c4.3-5 3.6-10.3 0-13.2L749.7 129.8a8.22 8.22 0 0 0-5.2-1.8c-4.6 0-8.4 3.8-8.4 8.4V209H199.7c-39.5 0-71.7 32.2-71.7 71.8V544c0 4.4 3.6 8 8 8m752-80h-63.6c-4.4 0-8 3.6-8 8v255.3H287.8v-72.6c0-1.9-.6-3.7-1.8-5.2a8.3 8.3 0 0 0-11.7-1.4L131 768.6c-4.3 5-3.6 10.3 0 13.2l143.3 112.4c1.5 1.2 3.3 1.8 5.2 1.8c4.6 0 8.4-3.8 8.4-8.4V815h536.6c39.5 0 71.7-32.2 71.7-71.8V480c-.2-4.4-3.8-8-8.2-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m154.7 454.5l-246 178c-5.3 3.8-12.7 0-12.7-6.5v-46.9c0-10.2 4.9-19.9 13.2-25.9L566.6 512L421.2 406.8c-8.3-6-13.2-15.6-13.2-25.9V334c0-6.5 7.4-10.3 12.7-6.5l246 178c4.4 3.2 4.4 9.8 0 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightCircleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m666.7 505.5l-246-178A8 8 0 0 0 408 334v46.9c0 10.2 4.9 19.9 13.2 25.9L566.6 512L421.2 617.2c-8.3 6-13.2 15.6-13.2 25.9V690c0 6.5 7.4 10.3 12.7 6.5l246-178c4.4-3.2 4.4-9.8 0-13"/><path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightCircleTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372m154.7 378.4l-246 178c-5.3 3.8-12.7 0-12.7-6.5V643c0-10.2 4.9-19.9 13.2-25.9L566.6 512L421.2 406.8c-8.3-6-13.2-15.6-13.2-25.9V334c0-6.5 7.4-10.3 12.7-6.5l246 178c4.4 3.2 4.4 9.7 0 12.9"/><path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" d="m666.7 505.5l-246-178c-5.3-3.8-12.7 0-12.7 6.5v46.9c0 10.3 4.9 19.9 13.2 25.9L566.6 512L421.2 617.1c-8.3 6-13.2 15.7-13.2 25.9v46.9c0 6.5 7.4 10.3 12.7 6.5l246-178c4.4-3.2 4.4-9.7 0-12.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M765.7 486.8L314.9 134.7A7.97 7.97 0 0 0 302 141v77.3c0 4.9 2.3 9.6 6.1 12.6l360 281.1l-360 281.1c-3.9 3-6.1 7.7-6.1 12.6V883c0 6.7 7.7 10.4 12.9 6.3l450.8-352.1a31.96 31.96 0 0 0 0-50.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightSquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M658.7 518.5l-246 178c-5.3 3.8-12.7 0-12.7-6.5v-46.9c0-10.2 4.9-19.9 13.2-25.9L558.6 512L413.2 406.8c-8.3-6-13.2-15.6-13.2-25.9V334c0-6.5 7.4-10.3 12.7-6.5l246 178c4.4 3.2 4.4 9.8 0 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightSquareOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m412.7 696.5l246-178c4.4-3.2 4.4-9.7 0-12.9l-246-178c-5.3-3.8-12.7 0-12.7 6.5V381c0 10.2 4.9 19.9 13.2 25.9L558.6 512L413.2 617.2c-8.3 6-13.2 15.6-13.2 25.9V690c0 6.5 7.4 10.3 12.7 6.5"/><path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightSquareTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/><path fill="currentColor" fill-opacity=".15" d="M184 840h656V184H184zm216-196.9c0-10.2 4.9-19.9 13.2-25.9L558.6 512L413.2 406.8c-8.3-6-13.2-15.6-13.2-25.9V334c0-6.5 7.4-10.3 12.7-6.5l246 178c4.4 3.2 4.4 9.7 0 12.9l-246 178c-5.3 3.9-12.7.1-12.7-6.4z"/><path fill="currentColor" d="m412.7 696.4l246-178c4.4-3.2 4.4-9.7 0-12.9l-246-178c-5.3-3.8-12.7 0-12.7 6.5v46.9c0 10.3 4.9 19.9 13.2 25.9L558.6 512L413.2 617.2c-8.3 6-13.2 15.7-13.2 25.9V690c0 6.5 7.4 10.3 12.7 6.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RiseOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m917 211.1l-199.2 24c-6.6.8-9.4 8.9-4.7 13.6l59.3 59.3l-226 226l-101.8-101.7c-6.3-6.3-16.4-6.2-22.6 0L100.3 754.1a8.03 8.03 0 0 0 0 11.3l45 45.2c3.1 3.1 8.2 3.1 11.3 0L433.3 534L535 635.7c6.3 6.2 16.4 6.2 22.6 0L829 364.5l59.3 59.3a8.01 8.01 0 0 0 13.6-4.7l24-199.2c.7-5.1-3.7-9.5-8.9-8.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RobotFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M852 64H172c-17.7 0-32 14.3-32 32v660c0 17.7 14.3 32 32 32h680c17.7 0 32-14.3 32-32V96c0-17.7-14.3-32-32-32M300 328c0-33.1 26.9-60 60-60s60 26.9 60 60s-26.9 60-60 60s-60-26.9-60-60m372 248c0 4.4-3.6 8-8 8H360c-4.4 0-8-3.6-8-8v-60c0-4.4 3.6-8 8-8h304c4.4 0 8 3.6 8 8zm-8-188c-33.1 0-60-26.9-60-60s26.9-60 60-60s60 26.9 60 60s-26.9 60-60 60m135 476H225c-13.8 0-25 14.3-25 32v56c0 4.4 2.8 8 6.2 8h611.5c3.4 0 6.2-3.6 6.2-8v-56c.1-17.7-11.1-32-24.9-32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RobotOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M300 328a60 60 0 1 0 120 0a60 60 0 1 0-120 0M852 64H172c-17.7 0-32 14.3-32 32v660c0 17.7 14.3 32 32 32h680c17.7 0 32-14.3 32-32V96c0-17.7-14.3-32-32-32m-32 660H204V128h616zM604 328a60 60 0 1 0 120 0a60 60 0 1 0-120 0m250.2 556H169.8c-16.5 0-29.8 14.3-29.8 32v36c0 4.4 3.3 8 7.4 8h729.1c4.1 0 7.4-3.6 7.4-8v-36c.1-17.7-13.2-32-29.7-32M664 508H360c-4.4 0-8 3.6-8 8v60c0 4.4 3.6 8 8 8h304c4.4 0 8-3.6 8-8v-60c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RocketFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M864 736c0-111.6-65.4-208-160-252.9V317.3c0-15.1-5.3-29.7-15.1-41.2L536.5 95.4C530.1 87.8 521 84 512 84s-18.1 3.8-24.5 11.4L335.1 276.1a63.97 63.97 0 0 0-15.1 41.2v165.8C225.4 528 160 624.4 160 736h156.5c-2.3 7.2-3.5 15-3.5 23.8c0 22.1 7.6 43.7 21.4 60.8a97.2 97.2 0 0 0 43.1 30.6c23.1 54 75.6 88.8 134.5 88.8c29.1 0 57.3-8.6 81.4-24.8c23.6-15.8 41.9-37.9 53-64a97 97 0 0 0 43.1-30.5a97.52 97.52 0 0 0 21.4-60.8c0-8.4-1.1-16.4-3.1-23.8zM512 352a48.01 48.01 0 0 1 0 96a48.01 48.01 0 0 1 0-96m116.1 432.2c-5.2 3-11.2 4.2-17.1 3.4l-19.5-2.4l-2.8 19.4c-5.4 37.9-38.4 66.5-76.7 66.5s-71.3-28.6-76.7-66.5l-2.8-19.5l-19.5 2.5a27.7 27.7 0 0 1-17.1-3.5c-8.7-5-14.1-14.3-14.1-24.4c0-10.6 5.9-19.4 14.6-23.8h231.3c8.8 4.5 14.6 13.3 14.6 23.8c-.1 10.2-5.5 19.6-14.2 24.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RocketOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M864 736c0-111.6-65.4-208-160-252.9V317.3c0-15.1-5.3-29.7-15.1-41.2L536.5 95.4C530.1 87.8 521 84 512 84s-18.1 3.8-24.5 11.4L335.1 276.1a63.97 63.97 0 0 0-15.1 41.2v165.8C225.4 528 160 624.4 160 736h156.5c-2.3 7.2-3.5 15-3.5 23.8c0 22.1 7.6 43.7 21.4 60.8a97.2 97.2 0 0 0 43.1 30.6c23.1 54 75.6 88.8 134.5 88.8c29.1 0 57.3-8.6 81.4-24.8c23.6-15.8 41.9-37.9 53-64a97 97 0 0 0 43.1-30.5a97.52 97.52 0 0 0 21.4-60.8c0-8.4-1.1-16.4-3.1-23.8H864zM762.3 621.4c9.4 14.6 17 30.3 22.5 46.6H700V558.7a211.6 211.6 0 0 1 62.3 62.7M388 483.1V318.8l124-147l124 147V668H388zM239.2 668c5.5-16.3 13.1-32 22.5-46.6c16.3-25.2 37.5-46.5 62.3-62.7V668zm388.9 116.2c-5.2 3-11.2 4.2-17.1 3.4l-19.5-2.4l-2.8 19.4c-5.4 37.9-38.4 66.5-76.7 66.5c-38.3 0-71.3-28.6-76.7-66.5l-2.8-19.5l-19.5 2.5a27.7 27.7 0 0 1-17.1-3.5c-8.7-5-14.1-14.3-14.1-24.4c0-10.6 5.9-19.4 14.6-23.8h231.3c8.8 4.5 14.6 13.3 14.6 23.8c-.1 10.2-5.5 19.6-14.2 24.5M464 400a48 48 0 1 0 96 0a48 48 0 1 0-96 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RocketTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M261.7 621.4c-9.4 14.6-17 30.3-22.5 46.6H324V558.7c-24.8 16.2-46 37.5-62.3 62.7M700 558.7V668h84.8c-5.5-16.3-13.1-32-22.5-46.6a211.6 211.6 0 0 0-62.3-62.7m-64-239.9l-124-147l-124 147V668h248zM512 448a48.01 48.01 0 0 1 0-96a48.01 48.01 0 0 1 0 96"/><path fill="currentColor" d="M864 736c0-111.6-65.4-208-160-252.9V317.3c0-15.1-5.3-29.7-15.1-41.2L536.5 95.4C530.1 87.8 521 84 512 84s-18.1 3.8-24.5 11.4L335.1 276.1a63.97 63.97 0 0 0-15.1 41.2v165.8C225.4 528 160 624.4 160 736h156.5c-2.3 7.2-3.5 15-3.5 23.8c0 22.1 7.6 43.7 21.4 60.8a97.2 97.2 0 0 0 43.1 30.6c23.1 54 75.6 88.8 134.5 88.8c29.1 0 57.3-8.6 81.4-24.8c23.6-15.8 41.9-37.9 53-64a97 97 0 0 0 43.1-30.5a97.52 97.52 0 0 0 21.4-60.8c0-8.4-1.1-16.4-3.1-23.8zm-540-68h-84.8c5.5-16.3 13.1-32 22.5-46.6c16.3-25.2 37.5-46.5 62.3-62.7zm64-184.9V318.8l124-147l124 147V668H388zm240.1 301.1c-5.2 3-11.2 4.2-17.1 3.4l-19.5-2.4l-2.8 19.4c-5.4 37.9-38.4 66.5-76.7 66.5s-71.3-28.6-76.7-66.5l-2.8-19.5l-19.5 2.5a27.7 27.7 0 0 1-17.1-3.5c-8.7-5-14.1-14.3-14.1-24.4c0-10.6 5.9-19.4 14.6-23.8h231.3c8.8 4.5 14.6 13.3 14.6 23.8c-.1 10.2-5.5 19.6-14.2 24.5M700 668V558.7a211.6 211.6 0 0 1 62.3 62.7c9.4 14.6 17 30.3 22.5 46.6z"/><path fill="currentColor" d="M464 400a48 48 0 1 0 96 0a48 48 0 1 0-96 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RollbackOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M793 242H366v-74c0-6.7-7.7-10.4-12.9-6.3l-142 112a8 8 0 0 0 0 12.6l142 112c5.2 4.1 12.9.4 12.9-6.3v-74h415v470H175c-4.4 0-8 3.6-8 8v60c0 4.4 3.6 8 8 8h618c35.3 0 64-28.7 64-64V306c0-35.3-28.7-64-64-64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RotateLeftOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M672 418H144c-17.7 0-32 14.3-32 32v414c0 17.7 14.3 32 32 32h528c17.7 0 32-14.3 32-32V450c0-17.7-14.3-32-32-32m-44 402H188V494h440z"/><path fill="currentColor" d="M819.3 328.5c-78.8-100.7-196-153.6-314.6-154.2l-.2-64c0-6.5-7.6-10.1-12.6-6.1l-128 101c-4 3.1-3.9 9.1 0 12.3L492 318.6c5.1 4 12.7.4 12.6-6.1v-63.9c12.9.1 25.9.9 38.8 2.5c42.1 5.2 82.1 18.2 119 38.7c38.1 21.2 71.2 49.7 98.4 84.3c27.1 34.7 46.7 73.7 58.1 115.8c11 40.7 14 82.7 8.9 124.8c-.7 5.4-1.4 10.8-2.4 16.1h74.9c14.8-103.6-11.3-213-81-302.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RotateRightOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M480.5 251.2c13-1.6 25.9-2.4 38.8-2.5v63.9c0 6.5 7.5 10.1 12.6 6.1L660 217.6c4-3.2 4-9.2 0-12.3l-128-101c-5.1-4-12.6-.4-12.6 6.1l-.2 64c-118.6.5-235.8 53.4-314.6 154.2c-69.6 89.2-95.7 198.6-81.1 302.4h74.9c-.9-5.3-1.7-10.7-2.4-16.1c-5.1-42.1-2.1-84.1 8.9-124.8c11.4-42.2 31-81.1 58.1-115.8c27.2-34.7 60.3-63.2 98.4-84.3c37-20.6 76.9-33.6 119.1-38.8"/><path fill="currentColor" d="M880 418H352c-17.7 0-32 14.3-32 32v414c0 17.7 14.3 32 32 32h528c17.7 0 32-14.3 32-32V450c0-17.7-14.3-32-32-32m-44 402H396V494h440z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SafetyCertificateFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M866.9 169.9L527.1 54.1C523 52.7 517.5 52 512 52s-11 .7-15.1 2.1L157.1 169.9c-8.3 2.8-15.1 12.4-15.1 21.2v482.4c0 8.8 5.7 20.4 12.6 25.9L499.3 968c3.5 2.7 8 4.1 12.6 4.1s9.2-1.4 12.6-4.1l344.7-268.6c6.9-5.4 12.6-17 12.6-25.9V191.1c.2-8.8-6.6-18.3-14.9-21.2M694.5 340.7L481.9 633.4a16.1 16.1 0 0 1-26 0l-126.4-174c-3.8-5.3 0-12.7 6.5-12.7h55.2c5.1 0 10 2.5 13 6.6l64.7 89l150.9-207.8c3-4.1 7.8-6.6 13-6.6H688c6.5.1 10.3 7.5 6.5 12.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SafetyCertificateOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M866.9 169.9L527.1 54.1C523 52.7 517.5 52 512 52s-11 .7-15.1 2.1L157.1 169.9c-8.3 2.8-15.1 12.4-15.1 21.2v482.4c0 8.8 5.7 20.4 12.6 25.9L499.3 968c3.5 2.7 8 4.1 12.6 4.1s9.2-1.4 12.6-4.1l344.7-268.6c6.9-5.4 12.6-17 12.6-25.9V191.1c.2-8.8-6.6-18.3-14.9-21.2M810 654.3L512 886.5L214 654.3V226.7l298-101.6l298 101.6zm-405.8-201c-3-4.1-7.8-6.6-13-6.6H336c-6.5 0-10.3 7.4-6.5 12.7l126.4 174a16.1 16.1 0 0 0 26 0l212.6-292.7c3.8-5.3 0-12.7-6.5-12.7h-55.2c-5.1 0-10 2.5-13 6.6L468.9 542.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SafetyCertificateTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M866.9 169.9L527.1 54.1C523 52.7 517.5 52 512 52s-11 .7-15.1 2.1L157.1 169.9c-8.3 2.8-15.1 12.4-15.1 21.2v482.4c0 8.8 5.7 20.4 12.6 25.9L499.3 968c3.5 2.7 8 4.1 12.6 4.1s9.2-1.4 12.6-4.1l344.7-268.6c6.9-5.4 12.6-17 12.6-25.9V191.1c.2-8.8-6.6-18.3-14.9-21.2M810 654.3L512 886.5L214 654.3V226.7l298-101.6l298 101.6z"/><path fill="currentColor" fill-opacity=".15" d="M214 226.7v427.6l298 232.2l298-232.2V226.7L512 125.1zM632.8 328H688c6.5 0 10.3 7.4 6.5 12.7L481.9 633.4a16.1 16.1 0 0 1-26 0l-126.4-174c-3.8-5.3 0-12.7 6.5-12.7h55.2c5.2 0 10 2.5 13 6.6l64.7 89.1l150.9-207.8c3-4.1 7.9-6.6 13-6.6"/><path fill="currentColor" d="M404.2 453.3c-3-4.1-7.8-6.6-13-6.6H336c-6.5 0-10.3 7.4-6.5 12.7l126.4 174a16.1 16.1 0 0 0 26 0l212.6-292.7c3.8-5.3 0-12.7-6.5-12.7h-55.2c-5.1 0-10 2.5-13 6.6L468.9 542.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SafetyOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64L128 192v384c0 212.1 171.9 384 384 384s384-171.9 384-384V192zm312 512c0 172.3-139.7 312-312 312S200 748.3 200 576V246l312-110l312 110z"/><path fill="currentColor" d="M378.4 475.1a35.91 35.91 0 0 0-50.9 0a35.91 35.91 0 0 0 0 50.9l129.4 129.4l2.1 2.1a33.98 33.98 0 0 0 48.1 0L730.6 434a33.98 33.98 0 0 0 0-48.1l-2.8-2.8a33.98 33.98 0 0 0-48.1 0L483 579.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SaveFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M893.3 293.3L730.7 130.7c-12-12-28.3-18.7-45.3-18.7H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V338.5c0-17-6.7-33.2-18.7-45.2M384 176h256v112H384zm128 554c-79.5 0-144-64.5-144-144s64.5-144 144-144s144 64.5 144 144s-64.5 144-144 144m0-224c-44.2 0-80 35.8-80 80s35.8 80 80 80s80-35.8 80-80s-35.8-80-80-80"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SaveOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M893.3 293.3L730.7 130.7c-7.5-7.5-16.7-13-26.7-16V112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V338.5c0-17-6.7-33.2-18.7-45.2M384 184h256v104H384zm456 656H184V184h136v136c0 17.7 14.3 32 32 32h320c17.7 0 32-14.3 32-32V205.8l136 136zM512 442c-79.5 0-144 64.5-144 144s64.5 144 144 144s144-64.5 144-144s-64.5-144-144-144m0 224c-44.2 0-80-35.8-80-80s35.8-80 80-80s80 35.8 80 80s-35.8 80-80 80"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SaveTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M704 320c0 17.7-14.3 32-32 32H352c-17.7 0-32-14.3-32-32V184H184v656h656V341.8l-136-136zM512 730c-79.5 0-144-64.5-144-144s64.5-144 144-144s144 64.5 144 144s-64.5 144-144 144"/><path fill="currentColor" d="M512 442c-79.5 0-144 64.5-144 144s64.5 144 144 144s144-64.5 144-144s-64.5-144-144-144m0 224c-44.2 0-80-35.8-80-80s35.8-80 80-80s80 35.8 80 80s-35.8 80-80 80"/><path fill="currentColor" d="M893.3 293.3L730.7 130.7c-.7-.7-1.4-1.3-2.1-2c-.1-.1-.3-.2-.4-.3c-.7-.7-1.5-1.3-2.2-1.9a64 64 0 0 0-22-11.7V112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V338.5c0-17-6.7-33.2-18.7-45.2M384 184h256v104H384zm456 656H184V184h136v136c0 17.7 14.3 32 32 32h320c17.7 0 32-14.3 32-32V205.8l136 136z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScanOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M136 384h56c4.4 0 8-3.6 8-8V200h176c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8H196c-37.6 0-68 30.4-68 68v180c0 4.4 3.6 8 8 8m512-184h176v176c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V196c0-37.6-30.4-68-68-68H648c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8M376 824H200V648c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v180c0 37.6 30.4 68 68 68h180c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m512-184h-56c-4.4 0-8 3.6-8 8v176H648c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h180c37.6 0 68-30.4 68-68V648c0-4.4-3.6-8-8-8m16-164H120c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScheduleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M928 224H768v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56H548v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56H328v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56H96c-17.7 0-32 14.3-32 32v576c0 17.7 14.3 32 32 32h832c17.7 0 32-14.3 32-32V256c0-17.7-14.3-32-32-32M424 688c0 4.4-3.6 8-8 8H232c-4.4 0-8-3.6-8-8v-48c0-4.4 3.6-8 8-8h184c4.4 0 8 3.6 8 8zm0-136c0 4.4-3.6 8-8 8H232c-4.4 0-8-3.6-8-8v-48c0-4.4 3.6-8 8-8h184c4.4 0 8 3.6 8 8zm374.5-91.3l-165 228.7a15.9 15.9 0 0 1-25.8 0L493.5 531.2c-3.8-5.3 0-12.7 6.5-12.7h54.9c5.1 0 9.9 2.5 12.9 6.6l52.8 73.1l103.7-143.7c3-4.2 7.8-6.6 12.9-6.6H792c6.5.1 10.3 7.5 6.5 12.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScheduleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M928 224H768v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56H548v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56H328v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56H96c-17.7 0-32 14.3-32 32v576c0 17.7 14.3 32 32 32h832c17.7 0 32-14.3 32-32V256c0-17.7-14.3-32-32-32m-40 568H136V296h120v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56h148v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56h148v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56h120zM416 496H232c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h184c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8m0 136H232c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h184c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8m308.2-177.4L620.6 598.3l-52.8-73.1c-3-4.2-7.8-6.6-12.9-6.6H500c-6.5 0-10.3 7.4-6.5 12.7l114.1 158.2a15.9 15.9 0 0 0 25.8 0l165-228.7c3.8-5.3 0-12.7-6.5-12.7H737c-5-.1-9.8 2.4-12.8 6.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScheduleTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M768 352c0 4.4-3.6 8-8 8h-56c-4.4 0-8-3.6-8-8v-56H548v56c0 4.4-3.6 8-8 8h-56c-4.4 0-8-3.6-8-8v-56H328v56c0 4.4-3.6 8-8 8h-56c-4.4 0-8-3.6-8-8v-56H136v496h752V296H768zM424 688c0 4.4-3.6 8-8 8H232c-4.4 0-8-3.6-8-8v-48c0-4.4 3.6-8 8-8h184c4.4 0 8 3.6 8 8zm0-136c0 4.4-3.6 8-8 8H232c-4.4 0-8-3.6-8-8v-48c0-4.4 3.6-8 8-8h184c4.4 0 8 3.6 8 8zm374.4-91.2l-165 228.7a15.9 15.9 0 0 1-25.8 0L493.5 531.3c-3.8-5.3 0-12.7 6.5-12.7h54.9c5.1 0 9.9 2.4 12.9 6.6l52.8 73.1l103.6-143.7c3-4.1 7.8-6.6 12.8-6.5h54.9c6.5 0 10.3 7.4 6.5 12.7"/><path fill="currentColor" d="M724.2 454.6L620.6 598.3l-52.8-73.1c-3-4.2-7.8-6.6-12.9-6.6H500c-6.5 0-10.3 7.4-6.5 12.7l114.1 158.2a15.9 15.9 0 0 0 25.8 0l165-228.7c3.8-5.3 0-12.7-6.5-12.7H737c-5-.1-9.8 2.4-12.8 6.5M416 496H232c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h184c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8"/><path fill="currentColor" d="M928 224H768v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56H548v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56H328v-56c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v56H96c-17.7 0-32 14.3-32 32v576c0 17.7 14.3 32 32 32h832c17.7 0 32-14.3 32-32V256c0-17.7-14.3-32-32-32m-40 568H136V296h120v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56h148v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56h148v56c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-56h120z"/><path fill="currentColor" d="M416 632H232c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h184c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScissorOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m567.1 512l318.5-319.3c5-5 1.5-13.7-5.6-13.7h-90.5c-2.1 0-4.2.8-5.6 2.3l-273.3 274l-90.2-90.5c12.5-22.1 19.7-47.6 19.7-74.8c0-83.9-68.1-152-152-152s-152 68.1-152 152s68.1 152 152 152c27.7 0 53.6-7.4 75.9-20.3l90 90.3l-90.1 90.3A151.04 151.04 0 0 0 288 582c-83.9 0-152 68.1-152 152s68.1 152 152 152s152-68.1 152-152c0-27.2-7.2-52.7-19.7-74.8l90.2-90.5l273.3 274c1.5 1.5 3.5 2.3 5.6 2.3H880c7.1 0 10.7-8.6 5.6-13.7zM288 370c-44.1 0-80-35.9-80-80s35.9-80 80-80s80 35.9 80 80s-35.9 80-80 80m0 444c-44.1 0-80-35.9-80-80s35.9-80 80-80s80 35.9 80 80s-35.9 80-80 80"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M909.6 854.5L649.9 594.8C690.2 542.7 712 479 712 412c0-80.2-31.3-155.4-87.9-212.1c-56.6-56.7-132-87.9-212.1-87.9s-155.5 31.3-212.1 87.9C143.2 256.5 112 331.8 112 412c0 80.1 31.3 155.5 87.9 212.1C256.5 680.8 331.8 712 412 712c67 0 130.6-21.8 182.7-62l259.7 259.6a8.2 8.2 0 0 0 11.6 0l43.6-43.5a8.2 8.2 0 0 0 0-11.6M570.4 570.4C528 612.7 471.8 636 412 636s-116-23.3-158.4-65.6C211.3 528 188 471.8 188 412s23.3-116.1 65.6-158.4C296 211.3 352.2 188 412 188s116.1 23.2 158.4 65.6S636 352.2 636 412s-23.3 116.1-65.6 158.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SecurityScanFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M866.9 169.9L527.1 54.1C523 52.7 517.5 52 512 52s-11 .7-15.1 2.1L157.1 169.9c-8.3 2.8-15.1 12.4-15.1 21.2v482.4c0 8.8 5.7 20.4 12.6 25.9L499.3 968c3.5 2.7 8 4.1 12.6 4.1s9.2-1.4 12.6-4.1l344.7-268.6c6.9-5.4 12.6-17 12.6-25.9V191.1c.2-8.8-6.6-18.3-14.9-21.2M626.8 554c-48.5 48.5-123 55.2-178.6 20.1l-77.5 77.5a8.03 8.03 0 0 1-11.3 0l-34-34a8.03 8.03 0 0 1 0-11.3l77.5-77.5c-35.1-55.7-28.4-130.1 20.1-178.6c56.3-56.3 147.5-56.3 203.8 0c56.3 56.3 56.3 147.5 0 203.8m-158.54-45.27a80.1 80.1 0 1 0 113.27-113.28a80.1 80.1 0 1 0-113.27 113.28"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SecurityScanOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M866.9 169.9L527.1 54.1C523 52.7 517.5 52 512 52s-11 .7-15.1 2.1L157.1 169.9c-8.3 2.8-15.1 12.4-15.1 21.2v482.4c0 8.8 5.7 20.4 12.6 25.9L499.3 968c3.5 2.7 8 4.1 12.6 4.1s9.2-1.4 12.6-4.1l344.7-268.6c6.9-5.4 12.6-17 12.6-25.9V191.1c.2-8.8-6.6-18.3-14.9-21.2M810 654.3L512 886.5L214 654.3V226.7l298-101.6l298 101.6zM402.9 528.8l-77.5 77.5a8.03 8.03 0 0 0 0 11.3l34 34c3.1 3.1 8.2 3.1 11.3 0l77.5-77.5c55.7 35.1 130.1 28.4 178.6-20.1c56.3-56.3 56.3-147.5 0-203.8c-56.3-56.3-147.5-56.3-203.8 0c-48.5 48.5-55.2 123-20.1 178.6m65.4-133.3c31.3-31.3 82-31.3 113.2 0c31.3 31.3 31.3 82 0 113.2c-31.3 31.3-82 31.3-113.2 0s-31.3-81.9 0-113.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SecurityScanTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M866.9 169.9L527.1 54.1C523 52.7 517.5 52 512 52s-11 .7-15.1 2.1L157.1 169.9c-8.3 2.8-15.1 12.4-15.1 21.2v482.4c0 8.8 5.7 20.4 12.6 25.9L499.3 968c3.5 2.7 8 4.1 12.6 4.1s9.2-1.4 12.6-4.1l344.7-268.6c6.9-5.4 12.6-17 12.6-25.9V191.1c.2-8.8-6.6-18.3-14.9-21.2M810 654.3L512 886.5L214 654.3V226.7l298-101.6l298 101.6z"/><path fill="currentColor" fill-opacity=".15" d="M460.7 451.1a80.1 80.1 0 1 0 160.2 0a80.1 80.1 0 1 0-160.2 0"/><path fill="currentColor" fill-opacity=".15" d="M214 226.7v427.6l298 232.2l298-232.2V226.7L512 125.1zm428.7 122.5c56.3 56.3 56.3 147.5 0 203.8c-48.5 48.5-123 55.2-178.6 20.1l-77.5 77.5a8.03 8.03 0 0 1-11.3 0l-34-34a8.03 8.03 0 0 1 0-11.3l77.5-77.5c-35.1-55.7-28.4-130.1 20.1-178.6c56.3-56.3 147.5-56.3 203.8 0"/><path fill="currentColor" d="m418.8 527.8l-77.5 77.5a8.03 8.03 0 0 0 0 11.3l34 34c3.1 3.1 8.2 3.1 11.3 0l77.5-77.5c55.6 35.1 130.1 28.4 178.6-20.1c56.3-56.3 56.3-147.5 0-203.8c-56.3-56.3-147.5-56.3-203.8 0c-48.5 48.5-55.2 122.9-20.1 178.6m65.4-133.3a80.1 80.1 0 0 1 113.3 0a80.1 80.1 0 0 1 0 113.3c-31.3 31.3-82 31.3-113.3 0s-31.3-82 0-113.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SelectOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h360c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8H184V184h656v320c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V144c0-17.7-14.3-32-32-32M653.3 599.4l52.2-52.2a8.01 8.01 0 0 0-4.7-13.6l-179.4-21c-5.1-.6-9.5 3.7-8.9 8.9l21 179.4c.8 6.6 8.9 9.4 13.6 4.7l52.4-52.4l256.2 256.2c3.1 3.1 8.2 3.1 11.3 0l42.4-42.4c3.1-3.1 3.1-8.2 0-11.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SendOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M931.4 498.9L94.9 79.5c-3.4-1.7-7.3-2.1-11-1.2c-8.5 2.1-13.8 10.7-11.7 19.3l86.2 352.2c1.3 5.3 5.2 9.6 10.4 11.3l147.7 50.7l-147.6 50.7c-5.2 1.8-9.1 6-10.3 11.3L72.2 926.5c-.9 3.7-.5 7.6 1.2 10.9c3.9 7.9 13.5 11.1 21.5 7.2l836.5-417c3.1-1.5 5.6-4.1 7.2-7.1c3.9-8 .7-17.6-7.2-21.6M170.8 826.3l50.3-205.6l295.2-101.3c2.3-.8 4.2-2.6 5-5c1.4-4.2-.8-8.7-5-10.2L221.1 403L171 198.2l628 314.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SettingFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512.5 390.6c-29.9 0-57.9 11.6-79.1 32.8c-21.1 21.2-32.8 49.2-32.8 79.1c0 29.9 11.7 57.9 32.8 79.1c21.2 21.1 49.2 32.8 79.1 32.8c29.9 0 57.9-11.7 79.1-32.8c21.1-21.2 32.8-49.2 32.8-79.1c0-29.9-11.7-57.9-32.8-79.1a110.96 110.96 0 0 0-79.1-32.8m412.3 235.5l-65.4-55.9c3.1-19 4.7-38.4 4.7-57.7s-1.6-38.8-4.7-57.7l65.4-55.9a32.03 32.03 0 0 0 9.3-35.2l-.9-2.6a442.5 442.5 0 0 0-79.6-137.7l-1.8-2.1a32.12 32.12 0 0 0-35.1-9.5l-81.2 28.9c-30-24.6-63.4-44-99.6-57.5l-15.7-84.9a32.05 32.05 0 0 0-25.8-25.7l-2.7-.5c-52-9.4-106.8-9.4-158.8 0l-2.7.5a32.05 32.05 0 0 0-25.8 25.7l-15.8 85.3a353.44 353.44 0 0 0-98.9 57.3l-81.8-29.1a32 32 0 0 0-35.1 9.5l-1.8 2.1a445.93 445.93 0 0 0-79.6 137.7l-.9 2.6c-4.5 12.5-.8 26.5 9.3 35.2l66.2 56.5c-3.1 18.8-4.6 38-4.6 57c0 19.2 1.5 38.4 4.6 57l-66 56.5a32.03 32.03 0 0 0-9.3 35.2l.9 2.6c18.1 50.3 44.8 96.8 79.6 137.7l1.8 2.1a32.12 32.12 0 0 0 35.1 9.5l81.8-29.1c29.8 24.5 63 43.9 98.9 57.3l15.8 85.3a32.05 32.05 0 0 0 25.8 25.7l2.7.5a448.27 448.27 0 0 0 158.8 0l2.7-.5a32.05 32.05 0 0 0 25.8-25.7l15.7-84.9c36.2-13.6 69.6-32.9 99.6-57.5l81.2 28.9a32 32 0 0 0 35.1-9.5l1.8-2.1c34.8-41.1 61.5-87.4 79.6-137.7l.9-2.6c4.3-12.4.6-26.3-9.5-35m-412.3 52.2c-97.1 0-175.8-78.7-175.8-175.8s78.7-175.8 175.8-175.8s175.8 78.7 175.8 175.8s-78.7 175.8-175.8 175.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SettingOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m924.8 625.7l-65.5-56c3.1-19 4.7-38.4 4.7-57.8s-1.6-38.8-4.7-57.8l65.5-56a32.03 32.03 0 0 0 9.3-35.2l-.9-2.6a443.74 443.74 0 0 0-79.7-137.9l-1.8-2.1a32.12 32.12 0 0 0-35.1-9.5l-81.3 28.9c-30-24.6-63.5-44-99.7-57.6l-15.7-85a32.05 32.05 0 0 0-25.8-25.7l-2.7-.5c-52.1-9.4-106.9-9.4-159 0l-2.7.5a32.05 32.05 0 0 0-25.8 25.7l-15.8 85.4a351.86 351.86 0 0 0-99 57.4l-81.9-29.1a32 32 0 0 0-35.1 9.5l-1.8 2.1a446.02 446.02 0 0 0-79.7 137.9l-.9 2.6c-4.5 12.5-.8 26.5 9.3 35.2l66.3 56.6c-3.1 18.8-4.6 38-4.6 57.1c0 19.2 1.5 38.4 4.6 57.1L99 625.5a32.03 32.03 0 0 0-9.3 35.2l.9 2.6c18.1 50.4 44.9 96.9 79.7 137.9l1.8 2.1a32.12 32.12 0 0 0 35.1 9.5l81.9-29.1c29.8 24.5 63.1 43.9 99 57.4l15.8 85.4a32.05 32.05 0 0 0 25.8 25.7l2.7.5a449.4 449.4 0 0 0 159 0l2.7-.5a32.05 32.05 0 0 0 25.8-25.7l15.7-85a350 350 0 0 0 99.7-57.6l81.3 28.9a32 32 0 0 0 35.1-9.5l1.8-2.1c34.8-41.1 61.6-87.5 79.7-137.9l.9-2.6c4.5-12.3.8-26.3-9.3-35M788.3 465.9c2.5 15.1 3.8 30.6 3.8 46.1s-1.3 31-3.8 46.1l-6.6 40.1l74.7 63.9a370.03 370.03 0 0 1-42.6 73.6L721 702.8l-31.4 25.8c-23.9 19.6-50.5 35-79.3 45.8l-38.1 14.3l-17.9 97a377.5 377.5 0 0 1-85 0l-17.9-97.2l-37.8-14.5c-28.5-10.8-55-26.2-78.7-45.7l-31.4-25.9l-93.4 33.2c-17-22.9-31.2-47.6-42.6-73.6l75.5-64.5l-6.5-40c-2.4-14.9-3.7-30.3-3.7-45.5c0-15.3 1.2-30.6 3.7-45.5l6.5-40l-75.5-64.5c11.3-26.1 25.6-50.7 42.6-73.6l93.4 33.2l31.4-25.9c23.7-19.5 50.2-34.9 78.7-45.7l37.9-14.3l17.9-97.2c28.1-3.2 56.8-3.2 85 0l17.9 97l38.1 14.3c28.7 10.8 55.4 26.2 79.3 45.8l31.4 25.8l92.8-32.9c17 22.9 31.2 47.6 42.6 73.6L781.8 426zM512 326c-97.2 0-176 78.8-176 176s78.8 176 176 176s176-78.8 176-176s-78.8-176-176-176m79.2 255.2A111.6 111.6 0 0 1 512 614c-29.9 0-58-11.7-79.2-32.8A111.6 111.6 0 0 1 400 502c0-29.9 11.7-58 32.8-79.2C454 401.6 482.1 390 512 390c29.9 0 58 11.6 79.2 32.8A111.6 111.6 0 0 1 624 502c0 29.9-11.7 58-32.8 79.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SettingTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="m859.3 569.7l.2.1c3.1-18.9 4.6-38.2 4.6-57.3c0-17.1-1.3-34.3-3.7-51.1c2.4 16.7 3.6 33.6 3.6 50.5c0 19.4-1.6 38.8-4.7 57.8M99 398.1c-.5-.4-.9-.8-1.4-1.3c.7.7 1.4 1.4 2.2 2.1l65.5 55.9v-.1zm536.6-216h.1l-15.5-83.8c-.2-1-.4-1.9-.7-2.8c.1.5.3 1.1.4 1.6zm54 546.5l31.4-25.8l92.8 32.9c17-22.9 31.3-47.5 42.6-73.6l-74.7-63.9l6.6-40.1c2.5-15.1 3.8-30.6 3.8-46.1s-1.3-31-3.8-46.1l-6.5-39.9l74.7-63.9c-11.4-26-25.6-50.7-42.6-73.6l-92.8 32.9l-31.4-25.8c-23.9-19.6-50.6-35-79.3-45.8l-38.1-14.3l-17.9-97a377.5 377.5 0 0 0-85 0l-17.9 97.2l-37.9 14.3c-28.5 10.8-55 26.2-78.7 45.7l-31.4 25.9l-93.4-33.2c-17 22.9-31.3 47.5-42.6 73.6l75.5 64.5l-6.5 40c-2.5 14.9-3.7 30.2-3.7 45.5c0 15.2 1.3 30.6 3.7 45.5l6.5 40l-75.5 64.5c11.4 26 25.6 50.7 42.6 73.6l93.4-33.2l31.4 25.9c23.7 19.5 50.2 34.9 78.7 45.7l37.8 14.5l17.9 97.2c28.2 3.2 56.9 3.2 85 0l17.9-97l38.1-14.3c28.8-10.8 55.4-26.2 79.3-45.8m-177.1-50.3c-30.5 0-59.2-7.8-84.3-21.5C373.3 627 336 568.9 336 502c0-97.2 78.8-176 176-176c66.9 0 125 37.3 154.8 92.2c13.7 25 21.5 53.7 21.5 84.3c0 97.1-78.7 175.8-175.8 175.8M207.2 812.8c-5.5 1.9-11.2 2.3-16.6 1.2c5.7 1.2 11.7 1 17.5-1l81.4-29c-.1-.1-.3-.2-.4-.3zm717.6-414.7l-65.5 56c0 .2.1.5.1.7l65.4-55.9c7.1-6.1 11.1-14.9 11.2-24c-.3 8.8-4.3 17.3-11.2 23.2"/><path fill="currentColor" fill-opacity=".15" d="M935.8 646.6c.5 4.7 0 9.5-1.7 14.1l-.9 2.6a446.02 446.02 0 0 1-79.7 137.9l-1.8 2.1a32 32 0 0 1-35.1 9.5l-81.3-28.9a350 350 0 0 1-99.7 57.6l-15.7 85a32.05 32.05 0 0 1-25.8 25.7l-2.7.5a445.2 445.2 0 0 1-79.2 7.1h.3c26.7 0 53.4-2.4 79.4-7.1l2.7-.5a32.05 32.05 0 0 0 25.8-25.7l15.7-84.9c36.2-13.6 69.6-32.9 99.6-57.5l81.2 28.9a32 32 0 0 0 35.1-9.5l1.8-2.1c34.8-41.1 61.5-87.4 79.6-137.7l.9-2.6c1.6-4.7 2.1-9.7 1.5-14.5"/><path fill="currentColor" d="M688 502c0-30.3-7.7-58.9-21.2-83.8C637 363.3 578.9 326 512 326c-97.2 0-176 78.8-176 176c0 66.9 37.3 125 92.2 154.8c24.9 13.5 53.4 21.2 83.8 21.2c97.2 0 176-78.8 176-176m-288 0c0-29.9 11.7-58 32.8-79.2C454 401.6 482.1 390 512 390c29.9 0 58 11.6 79.2 32.8A111.6 111.6 0 0 1 624 502c0 29.9-11.7 58-32.8 79.2A111.6 111.6 0 0 1 512 614c-29.9 0-58-11.7-79.2-32.8A111.6 111.6 0 0 1 400 502"/><path fill="currentColor" d="M594.1 952.2a32.05 32.05 0 0 0 25.8-25.7l15.7-85a350 350 0 0 0 99.7-57.6l81.3 28.9a32 32 0 0 0 35.1-9.5l1.8-2.1c34.8-41.1 61.6-87.5 79.7-137.9l.9-2.6c1.7-4.6 2.2-9.4 1.7-14.1c-.9-7.9-4.7-15.4-11-20.9l-65.3-55.9l-.2-.1c3.1-19 4.7-38.4 4.7-57.8c0-16.9-1.2-33.9-3.6-50.5c-.3-2.2-.7-4.4-1-6.6c0-.2-.1-.5-.1-.7l65.5-56c6.9-5.9 10.9-14.4 11.2-23.2c.1-4-.5-8.1-1.9-12l-.9-2.6a443.74 443.74 0 0 0-79.7-137.9l-1.8-2.1a32.12 32.12 0 0 0-35.1-9.5l-81.3 28.9c-30-24.6-63.4-44-99.6-57.6h-.1l-15.7-85c-.1-.5-.2-1.1-.4-1.6a32.08 32.08 0 0 0-25.4-24.1l-2.7-.5c-52.1-9.4-106.9-9.4-159 0l-2.7.5a32.05 32.05 0 0 0-25.8 25.7l-15.8 85.4a351.86 351.86 0 0 0-99 57.4l-81.9-29.1a32 32 0 0 0-35.1 9.5l-1.8 2.1a446.02 446.02 0 0 0-79.7 137.9l-.9 2.6a32.09 32.09 0 0 0 7.9 33.9c.5.4.9.9 1.4 1.3l66.3 56.6v.1c-3.1 18.8-4.6 37.9-4.6 57c0 19.2 1.5 38.4 4.6 57.1L99 625.5a32.03 32.03 0 0 0-9.3 35.2l.9 2.6c18.1 50.4 44.9 96.9 79.7 137.9l1.8 2.1c4.9 5.7 11.4 9.4 18.5 10.7c5.4 1 11.1.7 16.6-1.2l81.9-29.1c.1.1.3.2.4.3c29.7 24.3 62.8 43.6 98.6 57.1l15.8 85.4a32.05 32.05 0 0 0 25.8 25.7l2.7.5c26.1 4.7 52.8 7.1 79.5 7.1h.3c26.6 0 53.3-2.4 79.2-7.1zm-39.8-66.5a377.5 377.5 0 0 1-85 0l-17.9-97.2l-37.8-14.5c-28.5-10.8-55-26.2-78.7-45.7l-31.4-25.9l-93.4 33.2c-17-22.9-31.2-47.6-42.6-73.6l75.5-64.5l-6.5-40c-2.4-14.9-3.7-30.3-3.7-45.5c0-15.3 1.2-30.6 3.7-45.5l6.5-40l-75.5-64.5c11.3-26.1 25.6-50.7 42.6-73.6l93.4 33.2l31.4-25.9c23.7-19.5 50.2-34.9 78.7-45.7l37.9-14.3l17.9-97.2c28.1-3.2 56.8-3.2 85 0l17.9 97l38.1 14.3c28.7 10.8 55.4 26.2 79.3 45.8l31.4 25.8l92.8-32.9c17 22.9 31.2 47.6 42.6 73.6L781.8 426l6.5 39.9c2.5 15.1 3.8 30.6 3.8 46.1s-1.3 31-3.8 46.1l-6.6 40.1l74.7 63.9a370.03 370.03 0 0 1-42.6 73.6L721 702.8l-31.4 25.8c-23.9 19.6-50.5 35-79.3 45.8l-38.1 14.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShakeOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M324 666a48 48 0 1 0 96 0a48 48 0 1 0-96 0m616.7-309.6L667.6 83.2C655.2 70.9 638.7 64 621.1 64s-34.1 6.8-46.5 19.2L83.3 574.5a65.85 65.85 0 0 0 0 93.1l273.2 273.2c12.3 12.3 28.9 19.2 46.5 19.2s34.1-6.8 46.5-19.2l491.3-491.3c25.6-25.7 25.6-67.5-.1-93.1M403 880.1L143.9 621l477.2-477.2l259 259.2zM152.8 373.7a7.9 7.9 0 0 0 11.2 0L373.7 164a7.9 7.9 0 0 0 0-11.2l-38.4-38.4a7.9 7.9 0 0 0-11.2 0L114.3 323.9a7.9 7.9 0 0 0 0 11.2zm718.6 276.6a7.9 7.9 0 0 0-11.2 0L650.3 860.1a7.9 7.9 0 0 0 0 11.2l38.4 38.4a7.9 7.9 0 0 0 11.2 0L909.7 700a7.9 7.9 0 0 0 0-11.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareAltOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M752 664c-28.5 0-54.8 10-75.4 26.7L469.4 540.8a160.68 160.68 0 0 0 0-57.6l207.2-149.9C697.2 350 723.5 360 752 360c66.2 0 120-53.8 120-120s-53.8-120-120-120s-120 53.8-120 120c0 11.6 1.6 22.7 4.7 33.3L439.9 415.8C410.7 377.1 364.3 352 312 352c-88.4 0-160 71.6-160 160s71.6 160 160 160c52.3 0 98.7-25.1 127.9-63.8l196.8 142.5c-3.1 10.6-4.7 21.8-4.7 33.3c0 66.2 53.8 120 120 120s120-53.8 120-120s-53.8-120-120-120m0-476c28.7 0 52 23.3 52 52s-23.3 52-52 52s-52-23.3-52-52s23.3-52 52-52M312 600c-48.5 0-88-39.5-88-88s39.5-88 88-88s88 39.5 88 88s-39.5 88-88 88m440 236c-28.7 0-52-23.3-52-52s23.3-52 52-52s52 23.3 52 52s-23.3 52-52 52"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShopFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M882 272.1V144c0-17.7-14.3-32-32-32H174c-17.7 0-32 14.3-32 32v128.1c-16.7 1-30 14.9-30 31.9v131.7a177 177 0 0 0 14.4 70.4c4.3 10.2 9.6 19.8 15.6 28.9v345c0 17.6 14.3 32 32 32h274V736h128v176h274c17.7 0 32-14.3 32-32V535a175 175 0 0 0 15.6-28.9c9.5-22.3 14.4-46 14.4-70.4V304c0-17-13.3-30.9-30-31.9m-72 568H640V704c0-17.7-14.3-32-32-32H416c-17.7 0-32 14.3-32 32v136.1H214V597.9c2.9 1.4 5.9 2.8 9 4c22.3 9.4 46 14.1 70.4 14.1s48-4.7 70.4-14.1c13.8-5.8 26.8-13.2 38.7-22.1c.2-.1.4-.1.6 0a180.4 180.4 0 0 0 38.7 22.1c22.3 9.4 46 14.1 70.4 14.1c24.4 0 48-4.7 70.4-14.1c13.8-5.8 26.8-13.2 38.7-22.1c.2-.1.4-.1.6 0a180.4 180.4 0 0 0 38.7 22.1c22.3 9.4 46 14.1 70.4 14.1c24.4 0 48-4.7 70.4-14.1c3-1.3 6-2.6 9-4v242.2zm0-568.1H214v-88h596z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShopOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M882 272.1V144c0-17.7-14.3-32-32-32H174c-17.7 0-32 14.3-32 32v128.1c-16.7 1-30 14.9-30 31.9v131.7a177 177 0 0 0 14.4 70.4c4.3 10.2 9.6 19.8 15.6 28.9v345c0 17.6 14.3 32 32 32h676c17.7 0 32-14.3 32-32V535a175 175 0 0 0 15.6-28.9c9.5-22.3 14.4-46 14.4-70.4V304c0-17-13.3-30.9-30-31.9M214 184h596v88H214zm362 656.1H448V736h128zm234 0H640V704c0-17.7-14.3-32-32-32H416c-17.7 0-32 14.3-32 32v136.1H214V597.9c2.9 1.4 5.9 2.8 9 4c22.3 9.4 46 14.1 70.4 14.1s48-4.7 70.4-14.1c13.8-5.8 26.8-13.2 38.7-22.1c.2-.1.4-.1.6 0a180.4 180.4 0 0 0 38.7 22.1c22.3 9.4 46 14.1 70.4 14.1c24.4 0 48-4.7 70.4-14.1c13.8-5.8 26.8-13.2 38.7-22.1c.2-.1.4-.1.6 0a180.4 180.4 0 0 0 38.7 22.1c22.3 9.4 46 14.1 70.4 14.1c24.4 0 48-4.7 70.4-14.1c3-1.3 6-2.6 9-4v242.2zm30-404.4c0 59.8-49 108.3-109.3 108.3c-40.8 0-76.4-22.1-95.2-54.9c-2.9-5-8.1-8.1-13.9-8.1h-.6c-5.7 0-11 3.1-13.9 8.1A109.24 109.24 0 0 1 512 544c-40.7 0-76.2-22-95-54.7c-3-5.1-8.4-8.3-14.3-8.3s-11.4 3.2-14.3 8.3a109.63 109.63 0 0 1-95.1 54.7C233 544 184 495.5 184 435.7v-91.2c0-.3.2-.5.5-.5h655c.3 0 .5.2.5.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShopTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M839.5 344h-655c-.3 0-.5.2-.5.5v91.2c0 59.8 49 108.3 109.3 108.3c40.7 0 76.2-22 95.1-54.7c2.9-5.1 8.4-8.3 14.3-8.3s11.3 3.2 14.3 8.3c18.8 32.7 54.3 54.7 95 54.7c40.8 0 76.4-22.1 95.1-54.9c2.9-5 8.2-8.1 13.9-8.1h.6c5.8 0 11 3.1 13.9 8.1c18.8 32.8 54.4 54.9 95.2 54.9C791 544 840 495.5 840 435.7v-91.2c0-.3-.2-.5-.5-.5"/><path fill="currentColor" d="M882 272.1V144c0-17.7-14.3-32-32-32H174c-17.7 0-32 14.3-32 32v128.1c-16.7 1-30 14.9-30 31.9v131.7a177 177 0 0 0 14.4 70.4c4.3 10.2 9.6 19.8 15.6 28.9v345c0 17.6 14.3 32 32 32h676c17.7 0 32-14.3 32-32V535a175 175 0 0 0 15.6-28.9c9.5-22.3 14.4-46 14.4-70.4V304c0-17-13.3-30.9-30-31.9M214 184h596v88H214zm362 656.1H448V736h128zm234.4 0H640V704c0-17.7-14.3-32-32-32H416c-17.7 0-32 14.3-32 32v136.1H214V597.9c2.9 1.4 5.9 2.8 9 4c22.3 9.4 46 14.1 70.4 14.1c24.4 0 48-4.7 70.4-14.1c13.8-5.8 26.8-13.2 38.7-22.1c.2-.1.4-.1.6 0a180.4 180.4 0 0 0 38.7 22.1c22.3 9.4 46 14.1 70.4 14.1s48-4.7 70.4-14.1c13.8-5.8 26.8-13.2 38.7-22.1c.2-.1.4-.1.6 0a180.4 180.4 0 0 0 38.7 22.1c22.3 9.4 46 14.1 70.4 14.1s48-4.7 70.4-14.1c3-1.3 6-2.6 9-4zM840 435.7c0 59.8-49 108.3-109.3 108.3c-40.8 0-76.4-22.1-95.2-54.9c-2.9-5-8.1-8.1-13.9-8.1h-.6c-5.7 0-11 3.1-13.9 8.1A109.24 109.24 0 0 1 512 544c-40.7 0-76.2-22-95-54.7c-3-5.1-8.4-8.3-14.3-8.3s-11.4 3.2-14.3 8.3a109.63 109.63 0 0 1-95.1 54.7C233 544 184 495.5 184 435.7v-91.2c0-.3.2-.5.5-.5h655c.3 0 .5.2.5.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCartOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M922.9 701.9H327.4l29.9-60.9l496.8-.9c16.8 0 31.2-12 34.2-28.6l68.8-385.1c1.8-10.1-.9-20.5-7.5-28.4a34.99 34.99 0 0 0-26.6-12.5l-632-2.1l-5.4-25.4c-3.4-16.2-18-28-34.6-28H96.5a35.3 35.3 0 1 0 0 70.6h125.9L246 312.8l58.1 281.3l-74.8 122.1a34.96 34.96 0 0 0-3 36.8c6 11.9 18.1 19.4 31.5 19.4h62.8a102.43 102.43 0 0 0-20.6 61.7c0 56.6 46 102.6 102.6 102.6s102.6-46 102.6-102.6c0-22.3-7.4-44-20.6-61.7h161.1a102.43 102.43 0 0 0-20.6 61.7c0 56.6 46 102.6 102.6 102.6s102.6-46 102.6-102.6c0-22.3-7.4-44-20.6-61.7H923c19.4 0 35.3-15.8 35.3-35.3a35.42 35.42 0 0 0-35.4-35.2M305.7 253l575.8 1.9l-56.4 315.8l-452.3.8zm96.9 612.7c-17.4 0-31.6-14.2-31.6-31.6c0-17.4 14.2-31.6 31.6-31.6s31.6 14.2 31.6 31.6a31.6 31.6 0 0 1-31.6 31.6m325.1 0c-17.4 0-31.6-14.2-31.6-31.6c0-17.4 14.2-31.6 31.6-31.6s31.6 14.2 31.6 31.6a31.6 31.6 0 0 1-31.6 31.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 312H696v-16c0-101.6-82.4-184-184-184s-184 82.4-184 184v16H192c-17.7 0-32 14.3-32 32v536c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V344c0-17.7-14.3-32-32-32m-208 0H400v-16c0-61.9 50.1-112 112-112s112 50.1 112 112z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 312H696v-16c0-101.6-82.4-184-184-184s-184 82.4-184 184v16H192c-17.7 0-32 14.3-32 32v536c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V344c0-17.7-14.3-32-32-32m-432-16c0-61.9 50.1-112 112-112s112 50.1 112 112v16H400zm392 544H232V384h96v88c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-88h224v88c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-88h96z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M696 472c0 4.4-3.6 8-8 8h-56c-4.4 0-8-3.6-8-8v-88H400v88c0 4.4-3.6 8-8 8h-56c-4.4 0-8-3.6-8-8v-88h-96v456h560V384h-96z"/><path fill="currentColor" d="M832 312H696v-16c0-101.6-82.4-184-184-184s-184 82.4-184 184v16H192c-17.7 0-32 14.3-32 32v536c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V344c0-17.7-14.3-32-32-32m-432-16c0-61.9 50.1-112 112-112s112 50.1 112 112v16H400zm392 544H232V384h96v88c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-88h224v88c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-88h96z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShrinkOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m881.7 187.4l-45.1-45.1a8.03 8.03 0 0 0-11.3 0L667.8 299.9l-54.7-54.7a7.94 7.94 0 0 0-13.5 4.7L576.1 439c-.6 5.2 3.7 9.5 8.9 8.9l189.2-23.5c6.6-.8 9.3-8.8 4.7-13.5l-54.7-54.7l157.6-157.6c3-3 3-8.1-.1-11.2M439 576.1l-189.2 23.5c-6.6.8-9.3 8.9-4.7 13.5l54.7 54.7l-157.5 157.5a8.03 8.03 0 0 0 0 11.3l45.1 45.1c3.1 3.1 8.2 3.1 11.3 0l157.6-157.6l54.7 54.7a7.94 7.94 0 0 0 13.5-4.7L447.9 585a7.9 7.9 0 0 0-8.9-8.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignalFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M584 352H440c-17.7 0-32 14.3-32 32v544c0 17.7 14.3 32 32 32h144c17.7 0 32-14.3 32-32V384c0-17.7-14.3-32-32-32M892 64H748c-17.7 0-32 14.3-32 32v832c0 17.7 14.3 32 32 32h144c17.7 0 32-14.3 32-32V96c0-17.7-14.3-32-32-32M276 640H132c-17.7 0-32 14.3-32 32v256c0 17.7 14.3 32 32 32h144c17.7 0 32-14.3 32-32V672c0-17.7-14.3-32-32-32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SisternodeOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M672 432c-120.3 0-219.9 88.5-237.3 204H320c-15.5 0-28-12.5-28-28V244h291c14.2 35.2 48.7 60 89 60c53 0 96-43 96-96s-43-96-96-96c-40.3 0-74.8 24.8-89 60H112v72h108v364c0 55.2 44.8 100 100 100h114.7c17.4 115.5 117 204 237.3 204c132.5 0 240-107.5 240-240S804.5 432 672 432m128 266c0 4.4-3.6 8-8 8h-86v86c0 4.4-3.6 8-8 8h-52c-4.4 0-8-3.6-8-8v-86h-86c-4.4 0-8-3.6-8-8v-52c0-4.4 3.6-8 8-8h86v-86c0-4.4 3.6-8 8-8h52c4.4 0 8 3.6 8 8v86h86c4.4 0 8 3.6 8 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SketchCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m582.3 625.6l147.9-166.3h-63.4zm90-202.3h62.5l-92.1-115.1zm-274.7 36L512 684.5l114.4-225.2zM512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m286.7 380.2L515.8 762.3c-1 1.1-2.4 1.7-3.8 1.7s-2.8-.6-3.8-1.7L225.3 444.2a5.14 5.14 0 0 1-.2-6.6L365.6 262c1-1.2 2.4-1.9 4-1.9h284.6c1.6 0 3 .7 4 1.9l140.5 175.6a4.9 4.9 0 0 1 0 6.6m-190.5-20.9L512 326.1l-96.2 97.2zM420.3 301.1l-23.1 89.8l88.8-89.8zm183.4 0H538l88.8 89.8zm-222.4 7.1l-92.1 115.1h62.5zm-87.5 151.1l147.9 166.3l-84.5-166.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SketchOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m925.6 405.1l-203-253.7a6.5 6.5 0 0 0-5-2.4H306.4c-1.9 0-3.8.9-5 2.4l-203 253.7a6.5 6.5 0 0 0 .2 8.3l408.6 459.5c1.2 1.4 3 2.1 4.8 2.1c1.8 0 3.5-.8 4.8-2.1l408.6-459.5a6.5 6.5 0 0 0 .2-8.3M645.2 206.4l34.4 133.9l-132.5-133.9zm8.2 178.5H370.6L512 242zM378.8 206.4h98.1L344.3 340.3zm-53.4 7l-44.1 171.5h-93.1zM194.6 434.9H289l125.8 247.7zM512 763.4L345.1 434.9h333.7zm97.1-80.8L735 434.9h94.4zm133.6-297.7l-44.1-171.5l137.2 171.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SketchSquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M608.2 423.3L512 326.1l-96.2 97.2zm-25.9 202.3l147.9-166.3h-63.4zm90-202.3h62.5l-92.1-115.1zM880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-81.3 332.2L515.8 762.3c-1 1.1-2.4 1.7-3.8 1.7s-2.8-.6-3.8-1.7L225.3 444.2a5.14 5.14 0 0 1-.2-6.6L365.6 262c1-1.2 2.4-1.9 4-1.9h284.6c1.6 0 3 .7 4 1.9l140.5 175.6a4.9 4.9 0 0 1 0 6.6m-401.1 15.1L512 684.5l114.4-225.2zm-16.3-151.1l-92.1 115.1h62.5zm-87.5 151.1l147.9 166.3l-84.5-166.3zm126.5-158.2l-23.1 89.8l88.8-89.8zm183.4 0H538l88.8 89.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkinFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M870 126H663.8c-17.4 0-32.9 11.9-37 29.3C614.3 208.1 567 246 512 246s-102.3-37.9-114.8-90.7a37.93 37.93 0 0 0-37-29.3H154a44 44 0 0 0-44 44v252a44 44 0 0 0 44 44h75v388a44 44 0 0 0 44 44h478a44 44 0 0 0 44-44V466h75a44 44 0 0 0 44-44V170a44 44 0 0 0-44-44"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkinOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M870 126H663.8c-17.4 0-32.9 11.9-37 29.3C614.3 208.1 567 246 512 246s-102.3-37.9-114.8-90.7a37.93 37.93 0 0 0-37-29.3H154a44 44 0 0 0-44 44v252a44 44 0 0 0 44 44h75v388a44 44 0 0 0 44 44h478a44 44 0 0 0 44-44V466h75a44 44 0 0 0 44-44V170a44 44 0 0 0-44-44m-28 268H723v432H301V394H182V198h153.3c28.2 71.2 97.5 120 176.7 120s148.5-48.8 176.7-120H842z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkinTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M512 318c-79.2 0-148.5-48.8-176.7-120H182v196h119v432h422V394h119V198H688.7c-28.2 71.2-97.5 120-176.7 120"/><path fill="currentColor" d="M870 126H663.8c-17.4 0-32.9 11.9-37 29.3C614.3 208.1 567 246 512 246s-102.3-37.9-114.8-90.7a37.93 37.93 0 0 0-37-29.3H154a44 44 0 0 0-44 44v252a44 44 0 0 0 44 44h75v388a44 44 0 0 0 44 44h478a44 44 0 0 0 44-44V466h75a44 44 0 0 0 44-44V170a44 44 0 0 0-44-44m-28 268H723v432H301V394H182V198h153.3c28.2 71.2 97.5 120 176.7 120s148.5-48.8 176.7-120H842z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkypeFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M883.7 578.6c4.1-22.5 6.3-45.5 6.3-68.5c0-51-10-100.5-29.7-147c-19-45-46.3-85.4-81-120.1a375.79 375.79 0 0 0-120.1-80.9c-46.6-19.7-96-29.7-147-29.7c-24 0-48.1 2.3-71.5 6.8A225.1 225.1 0 0 0 335.6 113c-59.7 0-115.9 23.3-158.1 65.5A222.25 222.25 0 0 0 112 336.6c0 38 9.8 75.4 28.1 108.4c-3.7 21.4-5.7 43.3-5.7 65.1c0 51 10 100.5 29.7 147c19 45 46.2 85.4 80.9 120.1c34.7 34.7 75.1 61.9 120.1 80.9c46.6 19.7 96 29.7 147 29.7c22.2 0 44.4-2 66.2-5.9c33.5 18.9 71.3 29 110 29c59.7 0 115.9-23.2 158.1-65.5c42.3-42.2 65.5-98.4 65.5-158.1c.1-38-9.7-75.5-28.2-108.7m-370 162.9c-134.2 0-194.2-66-194.2-115.4c0-25.4 18.7-43.1 44.5-43.1c57.4 0 42.6 82.5 149.7 82.5c54.9 0 85.2-29.8 85.2-60.3c0-18.3-9-38.7-45.2-47.6l-119.4-29.8c-96.1-24.1-113.6-76.1-113.6-124.9c0-101.4 95.5-139.5 185.2-139.5c82.6 0 180 45.7 180 106.5c0 26.1-22.6 41.2-48.4 41.2c-49 0-40-67.8-138.7-67.8c-49 0-76.1 22.2-76.1 53.9s38.7 41.8 72.3 49.5l88.4 19.6c96.8 21.6 121.3 78.1 121.3 131.3c0 82.3-63.3 143.9-191 143.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkypeOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M883.7 578.6c4.1-22.5 6.3-45.5 6.3-68.5c0-51-10-100.5-29.7-147c-19-45-46.3-85.4-81-120.1a375.79 375.79 0 0 0-120.1-80.9c-46.6-19.7-96-29.7-147-29.7c-24 0-48.1 2.3-71.5 6.8A225.1 225.1 0 0 0 335.6 113c-59.7 0-115.9 23.3-158.1 65.5A222.25 222.25 0 0 0 112 336.6c0 38 9.8 75.4 28.1 108.4c-3.7 21.4-5.7 43.3-5.7 65.1c0 51 10 100.5 29.7 147c19 45 46.2 85.4 80.9 120.1c34.7 34.7 75.1 61.9 120.1 80.9c46.6 19.7 96 29.7 147 29.7c22.2 0 44.4-2 66.2-5.9c33.5 18.9 71.3 29 110 29c59.7 0 115.9-23.2 158.1-65.5c42.3-42.2 65.5-98.4 65.5-158.1c.1-38-9.7-75.5-28.2-108.7m-88.1 216C766.9 823.4 729 839 688.4 839c-26.1 0-51.8-6.8-74.6-19.7l-22.5-12.7l-25.5 4.5c-17.8 3.2-35.8 4.8-53.6 4.8c-41.4 0-81.3-8.1-119.1-24.1c-36.3-15.3-69-37.3-97.2-65.5a304.29 304.29 0 0 1-65.5-97.1c-16-37.7-24-77.6-24-119c0-17.4 1.6-35.2 4.6-52.8l4.4-25.1L203 410a151.02 151.02 0 0 1-19.1-73.4c0-40.6 15.7-78.5 44.4-107.2C257.1 200.7 295 185 335.6 185a153 153 0 0 1 71.4 17.9l22.4 11.8l24.8-4.8c18.9-3.6 38.4-5.5 58-5.5c41.4 0 81.3 8.1 119 24c36.5 15.4 69.1 37.4 97.2 65.5c28.2 28.1 50.2 60.8 65.6 97.2c16 37.7 24 77.6 24 119c0 18.4-1.7 37-5.1 55.5l-4.7 25.5l12.6 22.6c12.6 22.5 19.2 48 19.2 73.7c0 40.7-15.7 78.5-44.4 107.2M583.4 466.2L495 446.6c-33.6-7.7-72.3-17.8-72.3-49.5s27.1-53.9 76.1-53.9c98.7 0 89.7 67.8 138.7 67.8c25.8 0 48.4-15.2 48.4-41.2c0-60.8-97.4-106.5-180-106.5c-89.7 0-185.2 38.1-185.2 139.5c0 48.8 17.4 100.8 113.6 124.9l119.4 29.8c36.1 8.9 45.2 29.2 45.2 47.6c0 30.5-30.3 60.3-85.2 60.3c-107.2 0-92.3-82.5-149.7-82.5c-25.8 0-44.5 17.8-44.5 43.1c0 49.4 60 115.4 194.2 115.4c127.7 0 191-61.5 191-144c0-53.1-24.5-109.6-121.3-131.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlackCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64M361.5 580.2c0 27.8-22.5 50.4-50.3 50.4c-13.3 0-26.1-5.3-35.6-14.8c-9.4-9.5-14.7-22.3-14.7-35.6c0-27.8 22.5-50.4 50.3-50.4h50.3zm134 134.4c0 27.8-22.5 50.4-50.3 50.4c-27.8 0-50.3-22.6-50.3-50.4V580.2c0-27.8 22.5-50.4 50.3-50.4c13.3 0 26.1 5.3 35.6 14.8s14.7 22.3 14.7 35.6zm-50.2-218.4h-134c-27.8 0-50.3-22.6-50.3-50.4c0-27.8 22.5-50.4 50.3-50.4h134c27.8 0 50.3 22.6 50.3 50.4c-.1 27.9-22.6 50.4-50.3 50.4m0-134.4c-13.3 0-26.1-5.3-35.6-14.8S395 324.8 395 311.4c0-27.8 22.5-50.4 50.3-50.4c27.8 0 50.3 22.6 50.3 50.4v50.4zm83.7-50.4c0-27.8 22.5-50.4 50.3-50.4c27.8 0 50.3 22.6 50.3 50.4v134.4c0 27.8-22.5 50.4-50.3 50.4c-27.8 0-50.3-22.6-50.3-50.4zM579.3 765c-27.8 0-50.3-22.6-50.3-50.4v-50.4h50.3c27.8 0 50.3 22.6 50.3 50.4c0 27.8-22.5 50.4-50.3 50.4m134-134.4h-134c-13.3 0-26.1-5.3-35.6-14.8S529 593.6 529 580.2c0-27.8 22.5-50.4 50.3-50.4h134c27.8 0 50.3 22.6 50.3 50.4c0 27.8-22.5 50.4-50.3 50.4m0-134.4H663v-50.4c0-27.8 22.5-50.4 50.3-50.4s50.3 22.6 50.3 50.4c0 27.8-22.5 50.4-50.3 50.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlackOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M409.4 128c-42.4 0-76.7 34.4-76.7 76.8c0 20.3 8.1 39.9 22.4 54.3c14.4 14.4 33.9 22.5 54.3 22.5h76.7v-76.8c0-42.3-34.3-76.7-76.7-76.8m0 204.8H204.7c-42.4 0-76.7 34.4-76.7 76.8s34.4 76.8 76.7 76.8h204.6c42.4 0 76.7-34.4 76.7-76.8c.1-42.4-34.3-76.8-76.6-76.8M614 486.4c42.4 0 76.8-34.4 76.7-76.8V204.8c0-42.4-34.3-76.8-76.7-76.8c-42.4 0-76.7 34.4-76.7 76.8v204.8c0 42.5 34.3 76.8 76.7 76.8m281.4-76.8c0-42.4-34.4-76.8-76.7-76.8S742 367.2 742 409.6v76.8h76.7c42.3 0 76.7-34.4 76.7-76.8m-76.8 128H614c-42.4 0-76.7 34.4-76.7 76.8c0 20.3 8.1 39.9 22.4 54.3c14.4 14.4 33.9 22.5 54.3 22.5h204.6c42.4 0 76.7-34.4 76.7-76.8c.1-42.4-34.3-76.7-76.7-76.8M614 742.4h-76.7v76.8c0 42.4 34.4 76.8 76.7 76.8c42.4 0 76.8-34.4 76.7-76.8c.1-42.4-34.3-76.7-76.7-76.8M409.4 537.6c-42.4 0-76.7 34.4-76.7 76.8v204.8c0 42.4 34.4 76.8 76.7 76.8c42.4 0 76.8-34.4 76.7-76.8V614.4c0-20.3-8.1-39.9-22.4-54.3c-14.4-14.4-34-22.5-54.3-22.5M128 614.4c0 20.3 8.1 39.9 22.4 54.3c14.4 14.4 33.9 22.5 54.3 22.5c42.4 0 76.8-34.4 76.7-76.8v-76.8h-76.7c-42.3 0-76.7 34.4-76.7 76.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlackSquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M529 311.4c0-27.8 22.5-50.4 50.3-50.4c27.8 0 50.3 22.6 50.3 50.4v134.4c0 27.8-22.5 50.4-50.3 50.4c-27.8 0-50.3-22.6-50.3-50.4zM361.5 580.2c0 27.8-22.5 50.4-50.3 50.4c-13.3 0-26.1-5.3-35.6-14.8c-9.4-9.5-14.7-22.3-14.7-35.6c0-27.8 22.5-50.4 50.3-50.4h50.3zm134 134.4c0 27.8-22.5 50.4-50.3 50.4c-27.8 0-50.3-22.6-50.3-50.4V580.2c0-27.8 22.5-50.4 50.3-50.4c13.3 0 26.1 5.3 35.6 14.8s14.7 22.3 14.7 35.6zm-50.2-218.4h-134c-27.8 0-50.3-22.6-50.3-50.4c0-27.8 22.5-50.4 50.3-50.4h134c27.8 0 50.3 22.6 50.3 50.4c-.1 27.9-22.6 50.4-50.3 50.4m0-134.4c-13.3 0-26.1-5.3-35.6-14.8S395 324.8 395 311.4c0-27.8 22.5-50.4 50.3-50.4c27.8 0 50.3 22.6 50.3 50.4v50.4zm134 403.2c-27.8 0-50.3-22.6-50.3-50.4v-50.4h50.3c27.8 0 50.3 22.6 50.3 50.4c0 27.8-22.5 50.4-50.3 50.4m134-134.4h-134c-13.3 0-26.1-5.3-35.6-14.8c-9.4-9.5-14.7-22.3-14.7-35.6c0-27.8 22.5-50.4 50.3-50.4h134c27.8 0 50.3 22.6 50.3 50.4c0 27.8-22.5 50.4-50.3 50.4m0-134.4H663v-50.4c0-27.8 22.5-50.4 50.3-50.4s50.3 22.6 50.3 50.4c0 27.8-22.5 50.4-50.3 50.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlackSquareOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M529 311.4c0-27.8 22.5-50.4 50.3-50.4c27.8 0 50.3 22.6 50.3 50.4v134.4c0 27.8-22.5 50.4-50.3 50.4c-27.8 0-50.3-22.6-50.3-50.4zM361.5 580.2c0 27.8-22.5 50.4-50.3 50.4c-13.3 0-26.1-5.3-35.6-14.8c-9.4-9.5-14.7-22.3-14.7-35.6c0-27.8 22.5-50.4 50.3-50.4h50.3zm134 134.4c0 27.8-22.5 50.4-50.3 50.4c-27.8 0-50.3-22.6-50.3-50.4V580.2c0-27.8 22.5-50.4 50.3-50.4c13.3 0 26.1 5.3 35.6 14.8s14.7 22.3 14.7 35.6zm-50.2-218.4h-134c-27.8 0-50.3-22.6-50.3-50.4c0-27.8 22.5-50.4 50.3-50.4h134c27.8 0 50.3 22.6 50.3 50.4c-.1 27.9-22.6 50.4-50.3 50.4m0-134.4c-13.3 0-26.1-5.3-35.6-14.8S395 324.8 395 311.4c0-27.8 22.5-50.4 50.3-50.4c27.8 0 50.3 22.6 50.3 50.4v50.4zm134 403.2c-27.8 0-50.3-22.6-50.3-50.4v-50.4h50.3c27.8 0 50.3 22.6 50.3 50.4c0 27.8-22.5 50.4-50.3 50.4m134-134.4h-134c-13.3 0-26.1-5.3-35.6-14.8c-9.4-9.5-14.7-22.3-14.7-35.6c0-27.8 22.5-50.4 50.3-50.4h134c27.8 0 50.3 22.6 50.3 50.4c0 27.8-22.5 50.4-50.3 50.4m0-134.4H663v-50.4c0-27.8 22.5-50.4 50.3-50.4s50.3 22.6 50.3 50.4c0 27.8-22.5 50.4-50.3 50.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlidersFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M904 296h-66v-96c0-4.4-3.6-8-8-8h-52c-4.4 0-8 3.6-8 8v96h-66c-4.4 0-8 3.6-8 8v416c0 4.4 3.6 8 8 8h66v96c0 4.4 3.6 8 8 8h52c4.4 0 8-3.6 8-8v-96h66c4.4 0 8-3.6 8-8V304c0-4.4-3.6-8-8-8m-584-72h-66v-56c0-4.4-3.6-8-8-8h-52c-4.4 0-8 3.6-8 8v56h-66c-4.4 0-8 3.6-8 8v560c0 4.4 3.6 8 8 8h66v56c0 4.4 3.6 8 8 8h52c4.4 0 8-3.6 8-8v-56h66c4.4 0 8-3.6 8-8V232c0-4.4-3.6-8-8-8m292 180h-66V232c0-4.4-3.6-8-8-8h-52c-4.4 0-8 3.6-8 8v172h-66c-4.4 0-8 3.6-8 8v200c0 4.4 3.6 8 8 8h66v172c0 4.4 3.6 8 8 8h52c4.4 0 8-3.6 8-8V620h66c4.4 0 8-3.6 8-8V412c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlidersOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M320 224h-66v-56c0-4.4-3.6-8-8-8h-52c-4.4 0-8 3.6-8 8v56h-66c-4.4 0-8 3.6-8 8v560c0 4.4 3.6 8 8 8h66v56c0 4.4 3.6 8 8 8h52c4.4 0 8-3.6 8-8v-56h66c4.4 0 8-3.6 8-8V232c0-4.4-3.6-8-8-8m-60 508h-80V292h80zm644-436h-66v-96c0-4.4-3.6-8-8-8h-52c-4.4 0-8 3.6-8 8v96h-66c-4.4 0-8 3.6-8 8v416c0 4.4 3.6 8 8 8h66v96c0 4.4 3.6 8 8 8h52c4.4 0 8-3.6 8-8v-96h66c4.4 0 8-3.6 8-8V304c0-4.4-3.6-8-8-8m-60 364h-80V364h80zM612 404h-66V232c0-4.4-3.6-8-8-8h-52c-4.4 0-8 3.6-8 8v172h-66c-4.4 0-8 3.6-8 8v200c0 4.4 3.6 8 8 8h66v172c0 4.4 3.6 8 8 8h52c4.4 0 8-3.6 8-8V620h66c4.4 0 8-3.6 8-8V412c0-4.4-3.6-8-8-8m-60 145a3 3 0 0 1-3 3h-74a3 3 0 0 1-3-3v-74a3 3 0 0 1 3-3h74a3 3 0 0 1 3 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlidersTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M180 292h80v440h-80zm369 180h-74a3 3 0 0 0-3 3v74a3 3 0 0 0 3 3h74a3 3 0 0 0 3-3v-74a3 3 0 0 0-3-3m215-108h80v296h-80z"/><path fill="currentColor" d="M904 296h-66v-96c0-4.4-3.6-8-8-8h-52c-4.4 0-8 3.6-8 8v96h-66c-4.4 0-8 3.6-8 8v416c0 4.4 3.6 8 8 8h66v96c0 4.4 3.6 8 8 8h52c4.4 0 8-3.6 8-8v-96h66c4.4 0 8-3.6 8-8V304c0-4.4-3.6-8-8-8m-60 364h-80V364h80zM612 404h-66V232c0-4.4-3.6-8-8-8h-52c-4.4 0-8 3.6-8 8v172h-66c-4.4 0-8 3.6-8 8v200c0 4.4 3.6 8 8 8h66v172c0 4.4 3.6 8 8 8h52c4.4 0 8-3.6 8-8V620h66c4.4 0 8-3.6 8-8V412c0-4.4-3.6-8-8-8m-60 145a3 3 0 0 1-3 3h-74a3 3 0 0 1-3-3v-74a3 3 0 0 1 3-3h74a3 3 0 0 1 3 3zM320 224h-66v-56c0-4.4-3.6-8-8-8h-52c-4.4 0-8 3.6-8 8v56h-66c-4.4 0-8 3.6-8 8v560c0 4.4 3.6 8 8 8h66v56c0 4.4 3.6 8 8 8h52c4.4 0 8-3.6 8-8v-56h66c4.4 0 8-3.6 8-8V232c0-4.4-3.6-8-8-8m-60 508h-80V292h80z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmallDashOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M112 476h72v72h-72zm182 0h72v72h-72zm364 0h72v72h-72zm182 0h72v72h-72zm-364 0h72v72h-72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmileFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64M288 421a48.01 48.01 0 0 1 96 0a48.01 48.01 0 0 1-96 0m224 272c-85.5 0-155.6-67.3-160-151.6a8 8 0 0 1 8-8.4h48.1c4.2 0 7.8 3.2 8.1 7.4C420 589.9 461.5 629 512 629s92.1-39.1 95.8-88.6c.3-4.2 3.9-7.4 8.1-7.4H664a8 8 0 0 1 8 8.4C667.6 625.7 597.5 693 512 693m176-224a48.01 48.01 0 0 1 0-96a48.01 48.01 0 0 1 0 96"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmileOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M288 421a48 48 0 1 0 96 0a48 48 0 1 0-96 0m352 0a48 48 0 1 0 96 0a48 48 0 1 0-96 0M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m263 711c-34.2 34.2-74 61-118.3 79.8C611 874.2 562.3 884 512 884c-50.3 0-99-9.8-144.8-29.2A370.4 370.4 0 0 1 248.9 775c-34.2-34.2-61-74-79.8-118.3C149.8 611 140 562.3 140 512s9.8-99 29.2-144.8A370.4 370.4 0 0 1 249 248.9c34.2-34.2 74-61 118.3-79.8C413 149.8 461.7 140 512 140c50.3 0 99 9.8 144.8 29.2A370.4 370.4 0 0 1 775.1 249c34.2 34.2 61 74 79.8 118.3C874.2 413 884 461.7 884 512s-9.8 99-29.2 144.8A368.89 368.89 0 0 1 775 775M664 533h-48.1c-4.2 0-7.8 3.2-8.1 7.4C604 589.9 562.5 629 512 629s-92.1-39.1-95.8-88.6c-.3-4.2-3.9-7.4-8.1-7.4H360a8 8 0 0 0-8 8.4c4.4 84.3 74.5 151.6 160 151.6s155.6-67.3 160-151.6a8 8 0 0 0-8-8.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmileTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372M288 421a48.01 48.01 0 0 1 96 0a48.01 48.01 0 0 1-96 0m224 272c-85.5 0-155.6-67.3-160-151.6a8 8 0 0 1 8-8.4h48.1c4.2 0 7.8 3.2 8.1 7.4C420 589.9 461.5 629 512 629s92.1-39.1 95.8-88.6c.3-4.2 3.9-7.4 8.1-7.4H664a8 8 0 0 1 8 8.4C667.6 625.7 597.5 693 512 693m176-224a48.01 48.01 0 0 1 0-96a48.01 48.01 0 0 1 0 96"/><path fill="currentColor" d="M288 421a48 48 0 1 0 96 0a48 48 0 1 0-96 0m376 112h-48.1c-4.2 0-7.8 3.2-8.1 7.4c-3.7 49.5-45.3 88.6-95.8 88.6s-92-39.1-95.8-88.6c-.3-4.2-3.9-7.4-8.1-7.4H360a8 8 0 0 0-8 8.4c4.4 84.3 74.5 151.6 160 151.6s155.6-67.3 160-151.6a8 8 0 0 0-8-8.4m-24-112a48 48 0 1 0 96 0a48 48 0 1 0-96 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnippetsFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 112H724V72c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v40H500V72c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v40H320c-17.7 0-32 14.3-32 32v120h-96c-17.7 0-32 14.3-32 32v632c0 17.7 14.3 32 32 32h512c17.7 0 32-14.3 32-32v-96h96c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M664 486H514V336h.2L664 485.8zm128 274h-56V456L544 264H360v-80h68v32c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-32h152v32c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-32h68z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnippetsOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 112H724V72c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v40H500V72c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v40H320c-17.7 0-32 14.3-32 32v120h-96c-17.7 0-32 14.3-32 32v632c0 17.7 14.3 32 32 32h512c17.7 0 32-14.3 32-32v-96h96c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M664 888H232V336h218v174c0 22.1 17.9 40 40 40h174zm0-402H514V336h.2L664 485.8zm128 274h-56V456L544 264H360v-80h68v32c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-32h152v32c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-32h68z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnippetsTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M450 510V336H232v552h432V550H490c-22.1 0-40-17.9-40-40"/><path fill="currentColor" d="M832 112H724V72c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v40H500V72c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v40H320c-17.7 0-32 14.3-32 32v120h-96c-17.7 0-32 14.3-32 32v632c0 17.7 14.3 32 32 32h512c17.7 0 32-14.3 32-32v-96h96c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M664 888H232V336h218v174c0 22.1 17.9 40 40 40h174zm0-402H514V336h.2L664 485.8zm128 274h-56V456L544 264H360v-80h68v32c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-32h152v32c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-32h68z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SolutionOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M688 264c0-4.4-3.6-8-8-8H296c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h384c4.4 0 8-3.6 8-8zm-8 136H296c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h384c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8M480 544H296c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h184c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8m-48 308H208V148h560v344c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V108c0-17.7-14.3-32-32-32H168c-17.7 0-32 14.3-32 32v784c0 17.7 14.3 32 32 32h264c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m356.8-74.4c29-26.3 47.2-64.3 47.2-106.6c0-79.5-64.5-144-144-144s-144 64.5-144 144c0 42.3 18.2 80.3 47.2 106.6c-57 32.5-96.2 92.7-99.2 162.1c-.2 4.5 3.5 8.3 8 8.3h48.1c4.2 0 7.7-3.3 8-7.6C564 871.2 621.7 816 692 816s128 55.2 131.9 124.4c.2 4.2 3.7 7.6 8 7.6H880c4.6 0 8.2-3.8 8-8.3c-2.9-69.5-42.2-129.6-99.2-162.1M692 591c44.2 0 80 35.8 80 80s-35.8 80-80 80s-80-35.8-80-80s35.8-80 80-80"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAscendingOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M839.6 433.8L749 150.5a9.24 9.24 0 0 0-8.9-6.5h-77.4c-4.1 0-7.6 2.6-8.9 6.5l-91.3 283.3c-.3.9-.5 1.9-.5 2.9c0 5.1 4.2 9.3 9.3 9.3h56.4c4.2 0 7.8-2.8 9-6.8l17.5-61.6h89l17.3 61.5c1.1 4 4.8 6.8 9 6.8h61.2c1 0 1.9-.1 2.8-.4c2.4-.8 4.3-2.4 5.5-4.6c1.1-2.2 1.3-4.7.6-7.1M663.3 325.5l32.8-116.9h6.3l32.1 116.9zm143.5 492.9H677.2v-.4l132.6-188.9c1.1-1.6 1.7-3.4 1.7-5.4v-36.4c0-5.1-4.2-9.3-9.3-9.3h-204c-5.1 0-9.3 4.2-9.3 9.3v43c0 5.1 4.2 9.3 9.3 9.3h122.6v.4L587.7 828.9a9.35 9.35 0 0 0-1.7 5.4v36.4c0 5.1 4.2 9.3 9.3 9.3h211.4c5.1 0 9.3-4.2 9.3-9.3v-43a9.2 9.2 0 0 0-9.2-9.3M416 702h-76V172c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v530h-76c-6.7 0-10.5 7.8-6.3 13l112 141.9a8 8 0 0 0 12.6 0l112-141.9c4.1-5.2.4-13-6.3-13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortDescendingOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M839.6 433.8L749 150.5a9.24 9.24 0 0 0-8.9-6.5h-77.4c-4.1 0-7.6 2.6-8.9 6.5l-91.3 283.3c-.3.9-.5 1.9-.5 2.9c0 5.1 4.2 9.3 9.3 9.3h56.4c4.2 0 7.8-2.8 9-6.8l17.5-61.6h89l17.3 61.5c1.1 4 4.8 6.8 9 6.8h61.2c1 0 1.9-.1 2.8-.4c2.4-.8 4.3-2.4 5.5-4.6c1.1-2.2 1.3-4.7.6-7.1M663.3 325.5l32.8-116.9h6.3l32.1 116.9zm143.5 492.9H677.2v-.4l132.6-188.9c1.1-1.6 1.7-3.4 1.7-5.4v-36.4c0-5.1-4.2-9.3-9.3-9.3h-204c-5.1 0-9.3 4.2-9.3 9.3v43c0 5.1 4.2 9.3 9.3 9.3h122.6v.4L587.7 828.9a9.35 9.35 0 0 0-1.7 5.4v36.4c0 5.1 4.2 9.3 9.3 9.3h211.4c5.1 0 9.3-4.2 9.3-9.3v-43a9.2 9.2 0 0 0-9.2-9.3M310.3 167.1a8 8 0 0 0-12.6 0L185.7 309c-4.2 5.3-.4 13 6.3 13h76v530c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V322h76c6.7 0 10.5-7.8 6.3-13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoundFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m892.1 737.8l-110.3-63.7a15.9 15.9 0 0 0-21.7 5.9l-19.9 34.5c-4.4 7.6-1.8 17.4 5.8 21.8L856.3 800a15.9 15.9 0 0 0 21.7-5.9l19.9-34.5c4.4-7.6 1.7-17.4-5.8-21.8M760 344a15.9 15.9 0 0 0 21.7 5.9L892 286.2c7.6-4.4 10.2-14.2 5.8-21.8L878 230a15.9 15.9 0 0 0-21.7-5.9L746 287.8a15.99 15.99 0 0 0-5.8 21.8zm174 132H806c-8.8 0-16 7.2-16 16v40c0 8.8 7.2 16 16 16h128c8.8 0 16-7.2 16-16v-40c0-8.8-7.2-16-16-16M625.9 115c-5.9 0-11.9 1.6-17.4 5.3L254 352H90c-8.8 0-16 7.2-16 16v288c0 8.8 7.2 16 16 16h164l354.5 231.7c5.5 3.6 11.6 5.3 17.4 5.3c16.7 0 32.1-13.3 32.1-32.1V147.1c0-18.8-15.4-32.1-32.1-32.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoundOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M625.9 115c-5.9 0-11.9 1.6-17.4 5.3L254 352H90c-8.8 0-16 7.2-16 16v288c0 8.8 7.2 16 16 16h164l354.5 231.7c5.5 3.6 11.6 5.3 17.4 5.3c16.7 0 32.1-13.3 32.1-32.1V147.1c0-18.8-15.4-32.1-32.1-32.1M586 803L293.4 611.7l-18-11.7H146V424h129.4l17.9-11.7L586 221zm348-327H806c-8.8 0-16 7.2-16 16v40c0 8.8 7.2 16 16 16h128c8.8 0 16-7.2 16-16v-40c0-8.8-7.2-16-16-16m-41.9 261.8l-110.3-63.7a15.9 15.9 0 0 0-21.7 5.9l-19.9 34.5c-4.4 7.6-1.8 17.4 5.8 21.8L856.3 800a15.9 15.9 0 0 0 21.7-5.9l19.9-34.5c4.4-7.6 1.7-17.4-5.8-21.8M760 344a15.9 15.9 0 0 0 21.7 5.9L892 286.2c7.6-4.4 10.2-14.2 5.8-21.8L878 230a15.9 15.9 0 0 0-21.7-5.9L746 287.8a15.99 15.99 0 0 0-5.8 21.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoundTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M275.4 424H146v176h129.4l18 11.7L586 803V221L293.3 412.3z"/><path fill="currentColor" d="m892.1 737.8l-110.3-63.7a15.9 15.9 0 0 0-21.7 5.9l-19.9 34.5c-4.4 7.6-1.8 17.4 5.8 21.8L856.3 800a15.9 15.9 0 0 0 21.7-5.9l19.9-34.5c4.4-7.6 1.7-17.4-5.8-21.8M934 476H806c-8.8 0-16 7.2-16 16v40c0 8.8 7.2 16 16 16h128c8.8 0 16-7.2 16-16v-40c0-8.8-7.2-16-16-16M760 344a15.9 15.9 0 0 0 21.7 5.9L892 286.2c7.6-4.4 10.2-14.2 5.8-21.8L878 230a15.9 15.9 0 0 0-21.7-5.9L746 287.8a15.99 15.99 0 0 0-5.8 21.8zM625.9 115c-5.9 0-11.9 1.6-17.4 5.3L254 352H90c-8.8 0-16 7.2-16 16v288c0 8.8 7.2 16 16 16h164l354.5 231.7c5.5 3.6 11.6 5.3 17.4 5.3c16.7 0 32.1-13.3 32.1-32.1V147.1c0-18.8-15.4-32.1-32.1-32.1M586 803L293.4 611.7l-18-11.7H146V424h129.4l17.9-11.7L586 221z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SplitCellsOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M938.2 508.4L787.3 389c-3-2.4-7.3-.2-7.3 3.6V478H636V184h204v128c0 2.2 1.8 4 4 4h60c2.2 0 4-1.8 4-4V144c0-15.5-12.5-28-28-28H596c-15.5 0-28 12.5-28 28v736c0 15.5 12.5 28 28 28h284c15.5 0 28-12.5 28-28V712c0-2.2-1.8-4-4-4h-60c-2.2 0-4 1.8-4 4v128H636V546h144v85.4c0 3.8 4.4 6 7.3 3.6l150.9-119.4c2.4-1.8 2.4-5.4 0-7.2M428 116H144c-15.5 0-28 12.5-28 28v168c0 2.2 1.8 4 4 4h60c2.2 0 4-1.8 4-4V184h204v294H244v-85.4c0-3.8-4.3-6-7.3-3.6l-151 119.4c-2.3 1.8-2.3 5.3 0 7.1l151 119.5c2.9 2.3 7.3.2 7.3-3.6V546h144v294H184V712c0-2.2-1.8-4-4-4h-60c-2.2 0-4 1.8-4 4v168c0 15.5 12.5 28 28 28h284c15.5 0 28-12.5 28-28V144c0-15.5-12.5-28-28-28"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m908.1 353.1l-253.9-36.9L540.7 86.1c-3.1-6.3-8.2-11.4-14.5-14.5c-15.8-7.8-35-1.3-42.9 14.5L369.8 316.2l-253.9 36.9c-7 1-13.4 4.3-18.3 9.3a32.05 32.05 0 0 0 .6 45.3l183.7 179.1l-43.4 252.9a31.95 31.95 0 0 0 46.4 33.7L512 754l227.1 119.4c6.2 3.3 13.4 4.4 20.3 3.2c17.4-3 29.1-19.5 26.1-36.9l-43.4-252.9l183.7-179.1c5-4.9 8.3-11.3 9.3-18.3c2.7-17.5-9.5-33.7-27-36.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m908.1 353.1l-253.9-36.9L540.7 86.1c-3.1-6.3-8.2-11.4-14.5-14.5c-15.8-7.8-35-1.3-42.9 14.5L369.8 316.2l-253.9 36.9c-7 1-13.4 4.3-18.3 9.3a32.05 32.05 0 0 0 .6 45.3l183.7 179.1l-43.4 252.9a31.95 31.95 0 0 0 46.4 33.7L512 754l227.1 119.4c6.2 3.3 13.4 4.4 20.3 3.2c17.4-3 29.1-19.5 26.1-36.9l-43.4-252.9l183.7-179.1c5-4.9 8.3-11.3 9.3-18.3c2.7-17.5-9.5-33.7-27-36.3M664.8 561.6l36.1 210.3L512 672.7L323.1 772l36.1-210.3l-152.8-149L417.6 382L512 190.7L606.4 382l211.2 30.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="m512.5 190.4l-94.4 191.3l-211.2 30.7l152.8 149l-36.1 210.3l188.9-99.3l188.9 99.2l-36.1-210.3l152.8-148.9l-211.2-30.7z"/><path fill="currentColor" d="m908.6 352.8l-253.9-36.9L541.2 85.8c-3.1-6.3-8.2-11.4-14.5-14.5c-15.8-7.8-35-1.3-42.9 14.5L370.3 315.9l-253.9 36.9c-7 1-13.4 4.3-18.3 9.3a32.05 32.05 0 0 0 .6 45.3l183.7 179.1L239 839.4a31.95 31.95 0 0 0 46.4 33.7l227.1-119.4l227.1 119.4c6.2 3.3 13.4 4.4 20.3 3.2c17.4-3 29.1-19.5 26.1-36.9l-43.4-252.9l183.7-179.1c5-4.9 8.3-11.3 9.3-18.3c2.7-17.5-9.5-33.7-27-36.3M665.3 561.3l36.1 210.3l-188.9-99.2l-188.9 99.3l36.1-210.3l-152.8-149l211.2-30.7l94.4-191.3l94.4 191.3l211.2 30.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StepBackwardFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m347.6 528.95l383.2 301.02c14.25 11.2 35.2 1.1 35.2-16.95V210.97c0-18.05-20.95-28.14-35.2-16.94L347.6 495.05a21.53 21.53 0 0 0 0 33.9M330 864h-64a8 8 0 0 1-8-8V168a8 8 0 0 1 8-8h64a8 8 0 0 1 8 8v688a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StepBackwardOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m347.6 528.95l383.2 301.02c14.25 11.2 35.2 1.1 35.2-16.95V210.97c0-18.05-20.95-28.14-35.2-16.94L347.6 495.05a21.53 21.53 0 0 0 0 33.9M330 864h-64a8 8 0 0 1-8-8V168a8 8 0 0 1 8-8h64a8 8 0 0 1 8 8v688a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StepForwardFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M676.4 528.95L293.2 829.97c-14.25 11.2-35.2 1.1-35.2-16.95V210.97c0-18.05 20.95-28.14 35.2-16.94l383.2 301.02a21.53 21.53 0 0 1 0 33.9M694 864h64a8 8 0 0 0 8-8V168a8 8 0 0 0-8-8h-64a8 8 0 0 0-8 8v688a8 8 0 0 0 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StepForwardOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M676.4 528.95L293.2 829.97c-14.25 11.2-35.2 1.1-35.2-16.95V210.97c0-18.05 20.95-28.14 35.2-16.94l383.2 301.02a21.53 21.53 0 0 1 0 33.9M694 864h64a8 8 0 0 0 8-8V168a8 8 0 0 0-8-8h-64a8 8 0 0 0-8 8v688a8 8 0 0 0 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StockOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M904 747H120c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8M165.7 621.8l39.7 39.5c3.1 3.1 8.2 3.1 11.3 0l234.7-233.9l97.6 97.3a32.11 32.11 0 0 0 45.2 0l264.2-263.2c3.1-3.1 3.1-8.2 0-11.3l-39.7-39.6a8.03 8.03 0 0 0-11.3 0l-235.7 235l-97.7-97.3a32.11 32.11 0 0 0-45.2 0L165.7 610.5a7.94 7.94 0 0 0 0 11.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m234.8 736.5L223.5 277.2c16-19.7 34-37.7 53.7-53.7l523.3 523.3c-16 19.6-34 37.7-53.7 53.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372c0-89 31.3-170.8 83.5-234.8l523.3 523.3C682.8 852.7 601 884 512 884m288.5-137.2L277.2 223.5C341.2 171.3 423 140 512 140c205.4 0 372 166.6 372 372c0 89-31.3 170.8-83.5 234.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m288.5 682.8L277.7 224C258 240 240 258 224 277.7l522.8 522.8C682.8 852.7 601 884 512 884c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372c0 89-31.3 170.8-83.5 234.8"/><path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372c89 0 170.8-31.3 234.8-83.5L224 277.7c16-19.7 34-37.7 53.7-53.7l522.8 522.8C852.7 682.8 884 601 884 512c0-205.4-166.6-372-372-372"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StrikethroughOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M952 474H569.9c-10-2-20.5-4-31.6-6c-15.9-2.9-22.2-4.1-30.8-5.8c-51.3-10-82.2-20-106.8-34.2c-35.1-20.5-52.2-48.3-52.2-85.1c0-37 15.2-67.7 44-89c28.4-21 68.8-32.1 116.8-32.1c54.8 0 97.1 14.4 125.8 42.8c14.6 14.4 25.3 32.1 31.8 52.6c1.3 4.1 2.8 10 4.3 17.8c.9 4.8 5.2 8.2 9.9 8.2h72.8c5.6 0 10.1-4.6 10.1-10.1v-1c-.7-6.8-1.3-12.1-2-16c-7.3-43.5-28-81.7-59.7-110.3c-44.4-40.5-109.7-61.8-188.7-61.8c-72.3 0-137.4 18.1-183.3 50.9c-25.6 18.4-45.4 41.2-58.6 67.7c-13.5 27.1-20.3 58.4-20.3 92.9c0 29.5 5.7 54.5 17.3 76.5c8.3 15.7 19.6 29.5 34.1 42H72c-4.4 0-8 3.6-8 8v60c0 4.4 3.6 8 8 8h433.2c2.1.4 3.9.8 5.9 1.2c30.9 6.2 49.5 10.4 66.6 15.2c23 6.5 40.6 13.3 55.2 21.5c35.8 20.2 53.3 49.2 53.3 89c0 35.3-15.5 66.8-43.6 88.8c-30.5 23.9-75.6 36.4-130.5 36.4c-43.7 0-80.7-8.5-110.2-25c-29.1-16.3-49.1-39.8-59.7-69.5c-.8-2.2-1.7-5.2-2.7-9c-1.2-4.4-5.3-7.5-9.7-7.5h-79.7c-5.6 0-10.1 4.6-10.1 10.1v1c.2 2.3.4 4.2.6 5.7c6.5 48.8 30.3 88.8 70.7 118.8c47.1 34.8 113.4 53.2 191.8 53.2c84.2 0 154.8-19.8 204.2-57.3c25-18.9 44.2-42.2 57.1-69c13-27.1 19.7-57.9 19.7-91.5c0-31.8-5.8-58.4-17.8-81.4c-5.8-11.2-13.1-21.5-21.8-30.8H952c4.4 0 8-3.6 8-8v-60a8 8 0 0 0-8-7.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SubnodeOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M688 240c-138 0-252 102.8-269.6 236H249c-14.2-35.2-48.7-60-89-60c-53 0-96 43-96 96s43 96 96 96c40.3 0 74.8-24.8 89-60h169.3C436 681.2 550 784 688 784c150.2 0 272-121.8 272-272S838.2 240 688 240m128 298c0 4.4-3.6 8-8 8h-86v86c0 4.4-3.6 8-8 8h-52c-4.4 0-8-3.6-8-8v-86h-86c-4.4 0-8-3.6-8-8v-52c0-4.4 3.6-8 8-8h86v-86c0-4.4 3.6-8 8-8h52c4.4 0 8 3.6 8 8v86h86c4.4 0 8 3.6 8 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwapLeftOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M872 572H266.8l144.3-183c4.1-5.2.4-13-6.3-13H340c-9.8 0-19.1 4.5-25.1 12.2l-164 208c-16.5 21-1.6 51.8 25.1 51.8h696c4.4 0 8-3.6 8-8v-60c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwapOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M847.9 592H152c-4.4 0-8 3.6-8 8v60c0 4.4 3.6 8 8 8h605.2L612.9 851c-4.1 5.2-.4 13 6.3 13h72.5c4.9 0 9.5-2.2 12.6-6.1l168.8-214.1c16.5-21 1.6-51.8-25.2-51.8M872 356H266.8l144.3-183c4.1-5.2.4-13-6.3-13h-72.5c-4.9 0-9.5 2.2-12.6 6.1L150.9 380.2c-16.5 21-1.6 51.8 25.1 51.8h696c4.4 0 8-3.6 8-8v-60c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwapRightOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m873.1 596.2l-164-208A32 32 0 0 0 684 376h-64.8c-6.7 0-10.4 7.7-6.3 13l144.3 183H152c-4.4 0-8 3.6-8 8v60c0 4.4 3.6 8 8 8h695.9c26.8 0 41.7-30.8 25.2-51.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwitcherFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M752 240H144c-17.7 0-32 14.3-32 32v608c0 17.7 14.3 32 32 32h608c17.7 0 32-14.3 32-32V272c0-17.7-14.3-32-32-32M596 606c0 4.4-3.6 8-8 8H308c-4.4 0-8-3.6-8-8v-48c0-4.4 3.6-8 8-8h280c4.4 0 8 3.6 8 8zm284-494H264c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h576v576c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V144c0-17.7-14.3-32-32-32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwitcherOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M752 240H144c-17.7 0-32 14.3-32 32v608c0 17.7 14.3 32 32 32h608c17.7 0 32-14.3 32-32V272c0-17.7-14.3-32-32-32m-40 600H184V312h528zm168-728H264c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h576v576c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V144c0-17.7-14.3-32-32-32M300 550h296v64H300z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwitcherTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M184 840h528V312H184zm116-290h296v64H300z"/><path fill="currentColor" d="M880 112H264c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h576v576c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V144c0-17.7-14.3-32-32-32"/><path fill="currentColor" d="M752 240H144c-17.7 0-32 14.3-32 32v608c0 17.7 14.3 32 32 32h608c17.7 0 32-14.3 32-32V272c0-17.7-14.3-32-32-32m-40 600H184V312h528z"/><path fill="currentColor" d="M300 550h296v64H300z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SyncOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M168 504.2c1-43.7 10-86.1 26.9-126c17.3-41 42.1-77.7 73.7-109.4S337 212.3 378 195c42.4-17.9 87.4-27 133.9-27s91.5 9.1 133.8 27A341.5 341.5 0 0 1 755 268.8c9.9 9.9 19.2 20.4 27.8 31.4l-60.2 47a8 8 0 0 0 3 14.1l175.7 43c5 1.2 9.9-2.6 9.9-7.7l.8-180.9c0-6.7-7.7-10.5-12.9-6.3l-56.4 44.1C765.8 155.1 646.2 92 511.8 92C282.7 92 96.3 275.6 92 503.8a8 8 0 0 0 8 8.2h60c4.4 0 7.9-3.5 8-7.8m756 7.8h-60c-4.4 0-7.9 3.5-8 7.8c-1 43.7-10 86.1-26.9 126c-17.3 41-42.1 77.8-73.7 109.4A342.45 342.45 0 0 1 512.1 856a342.24 342.24 0 0 1-243.2-100.8c-9.9-9.9-19.2-20.4-27.8-31.4l60.2-47a8 8 0 0 0-3-14.1l-175.7-43c-5-1.2-9.9 2.6-9.9 7.7l-.7 181c0 6.7 7.7 10.5 12.9 6.3l56.4-44.1C258.2 868.9 377.8 932 512.2 932c229.2 0 415.5-183.7 419.8-411.8a8 8 0 0 0-8-8.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M928 160H96c-17.7 0-32 14.3-32 32v640c0 17.7 14.3 32 32 32h832c17.7 0 32-14.3 32-32V192c0-17.7-14.3-32-32-32m-40 208H676V232h212zm0 224H676V432h212zM412 432h200v160H412zm200-64H412V232h200zm-476 64h212v160H136zm0-200h212v136H136zm0 424h212v136H136zm276 0h200v136H412zm476 136H676V656h212z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TabletFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M800 64H224c-35.3 0-64 28.7-64 64v768c0 35.3 28.7 64 64 64h576c35.3 0 64-28.7 64-64V128c0-35.3-28.7-64-64-64M512 824c-22.1 0-40-17.9-40-40s17.9-40 40-40s40 17.9 40 40s-17.9 40-40 40"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TabletOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M800 64H224c-35.3 0-64 28.7-64 64v768c0 35.3 28.7 64 64 64h576c35.3 0 64-28.7 64-64V128c0-35.3-28.7-64-64-64m-8 824H232V136h560zM472 784a40 40 0 1 0 80 0a40 40 0 1 0-80 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TabletTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M800 64H224c-35.3 0-64 28.7-64 64v768c0 35.3 28.7 64 64 64h576c35.3 0 64-28.7 64-64V128c0-35.3-28.7-64-64-64m-8 824H232V136h560z"/><path fill="currentColor" fill-opacity=".15" d="M232 888h560V136H232zm280-144c22.1 0 40 17.9 40 40s-17.9 40-40 40s-40-17.9-40-40s17.9-40 40-40"/><path fill="currentColor" d="M472 784a40 40 0 1 0 80 0a40 40 0 1 0-80 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m938 458.8l-29.6-312.6c-1.5-16.2-14.4-29-30.6-30.6L565.2 86h-.4c-3.2 0-5.7 1-7.6 2.9L88.9 557.2a9.96 9.96 0 0 0 0 14.1l363.8 363.8c1.9 1.9 4.4 2.9 7.1 2.9s5.2-1 7.1-2.9l468.3-468.3c2-2.1 3-5 2.8-8M699 387c-35.3 0-64-28.7-64-64s28.7-64 64-64s64 28.7 64 64s-28.7 64-64 64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m938 458.8l-29.6-312.6c-1.5-16.2-14.4-29-30.6-30.6L565.2 86h-.4c-3.2 0-5.7 1-7.6 2.9L88.9 557.2a9.96 9.96 0 0 0 0 14.1l363.8 363.8c1.9 1.9 4.4 2.9 7.1 2.9s5.2-1 7.1-2.9l468.3-468.3c2-2.1 3-5 2.8-8M459.7 834.7L189.3 564.3L589 164.6L836 188l23.4 247zM680 256c-48.5 0-88 39.5-88 88s39.5 88 88 88s88-39.5 88-88s-39.5-88-88-88m0 120c-17.7 0-32-14.3-32-32s14.3-32 32-32s32 14.3 32 32s-14.3 32-32 32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M589 164.6L189.3 564.3l270.4 270.4L859.4 435L836 188zM680 432c-48.5 0-88-39.5-88-88s39.5-88 88-88s88 39.5 88 88s-39.5 88-88 88"/><path fill="currentColor" d="M680 256c-48.5 0-88 39.5-88 88s39.5 88 88 88s88-39.5 88-88s-39.5-88-88-88m0 120c-17.7 0-32-14.3-32-32s14.3-32 32-32s32 14.3 32 32s-14.3 32-32 32"/><path fill="currentColor" d="m938 458.8l-29.6-312.6c-1.5-16.2-14.4-29-30.6-30.6L565.2 86h-.4c-3.2 0-5.7 1-7.6 2.9L88.9 557.2a9.96 9.96 0 0 0 0 14.1l363.8 363.8a9.9 9.9 0 0 0 7.1 2.9c2.7 0 5.2-1 7.1-2.9l468.3-468.3c2-2.1 3-5 2.8-8M459.7 834.7L189.3 564.3L589 164.6L836 188l23.4 247z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagsFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M483.2 790.3L861.4 412c1.7-1.7 2.5-4 2.3-6.3l-25.5-301.4c-.7-7.8-6.8-13.9-14.6-14.6L522.2 64.3c-2.3-.2-4.7.6-6.3 2.3L137.7 444.8a8.03 8.03 0 0 0 0 11.3l334.2 334.2c3.1 3.2 8.2 3.2 11.3 0m122.7-533.4c18.7-18.7 49.1-18.7 67.9 0c18.7 18.7 18.7 49.1 0 67.9c-18.7 18.7-49.1 18.7-67.9 0c-18.7-18.7-18.7-49.1 0-67.9m283.8 282.9l-39.6-39.5a8.03 8.03 0 0 0-11.3 0l-362 361.3l-237.6-237a8.03 8.03 0 0 0-11.3 0l-39.6 39.5a8.03 8.03 0 0 0 0 11.3l243.2 242.8l39.6 39.5c3.1 3.1 8.2 3.1 11.3 0l407.3-406.6c3.1-3.1 3.1-8.2 0-11.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagsOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M483.2 790.3L861.4 412c1.7-1.7 2.5-4 2.3-6.3l-25.5-301.4c-.7-7.8-6.8-13.9-14.6-14.6L522.2 64.3c-2.3-.2-4.7.6-6.3 2.3L137.7 444.8a8.03 8.03 0 0 0 0 11.3l334.2 334.2c3.1 3.2 8.2 3.2 11.3 0m62.6-651.7l224.6 19l19 224.6L477.5 694L233.9 450.5zm60.16 186.23a48 48 0 1 0 67.88-67.89a48 48 0 1 0-67.88 67.89M889.7 539.8l-39.6-39.5a8.03 8.03 0 0 0-11.3 0l-362 361.3l-237.6-237a8.03 8.03 0 0 0-11.3 0l-39.6 39.5a8.03 8.03 0 0 0 0 11.3l243.2 242.8l39.6 39.5c3.1 3.1 8.2 3.1 11.3 0l407.3-406.6c3.1-3.1 3.1-8.2 0-11.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagsTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="m477.5 694l311.9-311.8l-19-224.6l-224.6-19l-311.9 311.9zm116-415.5a47.81 47.81 0 0 1 33.9-33.9c16.6-4.4 34.2.3 46.4 12.4a47.93 47.93 0 0 1 12.4 46.4a47.81 47.81 0 0 1-33.9 33.9c-16.6 4.4-34.2-.3-46.4-12.4a48.3 48.3 0 0 1-12.4-46.4"/><path fill="currentColor" fill-opacity=".15" d="M476.6 792.6c-1.7-.2-3.4-1-4.7-2.3L137.7 456.1a8.03 8.03 0 0 1 0-11.3L515.9 66.6c1.2-1.3 2.9-2.1 4.7-2.3h-.4c-2.3-.2-4.7.6-6.3 2.3L135.7 444.8a8.03 8.03 0 0 0 0 11.3l334.2 334.2c1.8 1.9 4.3 2.6 6.7 2.3"/><path fill="currentColor" d="m889.7 539.8l-39.6-39.5a8.03 8.03 0 0 0-11.3 0l-362 361.3l-237.6-237a8.03 8.03 0 0 0-11.3 0l-39.6 39.5a8.03 8.03 0 0 0 0 11.3l243.2 242.8l39.6 39.5c3.1 3.1 8.2 3.1 11.3 0l407.3-406.6c3.1-3.1 3.1-8.2 0-11.3M652.3 337.3a47.81 47.81 0 0 0 33.9-33.9c4.4-16.6-.3-34.2-12.4-46.4a47.93 47.93 0 0 0-46.4-12.4a47.81 47.81 0 0 0-33.9 33.9c-4.4 16.6.3 34.2 12.4 46.4a48.3 48.3 0 0 0 46.4 12.4"/><path fill="currentColor" d="M137.7 444.8a8.03 8.03 0 0 0 0 11.3l334.2 334.2c1.3 1.3 2.9 2.1 4.7 2.3c2.4.3 4.8-.5 6.6-2.3L861.4 412c1.7-1.7 2.5-4 2.3-6.3l-25.5-301.4c-.7-7.8-6.8-13.9-14.6-14.6L522.2 64.3h-1.6c-1.8.2-3.4 1-4.7 2.3zm408.1-306.2l224.6 19l19 224.6L477.5 694L233.9 450.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TaobaoCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64M315.7 291.5c27.3 0 49.5 22.1 49.5 49.4s-22.1 49.4-49.5 49.4a49.4 49.4 0 1 1 0-98.8M366.9 578c-13.6 42.3-10.2 26.7-64.4 144.5l-78.5-49s87.7-79.8 105.6-116.2c19.2-38.4-21.1-58.9-21.1-58.9l-60.2-37.5l32.7-50.2c45.4 33.7 48.7 36.6 79.2 67.2c23.8 23.9 20.7 56.8 6.7 100.1m427.2 55c-15.3 143.8-202.4 90.3-202.4 90.3l10.2-41.1l43.3 9.3c80 5 72.3-64.9 72.3-64.9V423c.6-77.3-72.6-85.4-204.2-38.3l30.6 8.3c-2.5 9-12.5 23.2-25.2 38.6h176v35.6h-99.1v44.5h98.7v35.7h-98.7V622c14.9-4.8 28.6-11.5 40.5-20.5l-8.7-32.5l46.5-14.4l38.8 94.9l-57.3 23.9l-10.2-37.8c-25.6 19.5-78.8 48-171.8 45.4c-99.2 2.6-73.7-112-73.7-112l2.5-1.3H472c-.5 14.7-6.6 38.7 1.7 51.8c6.8 10.8 24.2 12.6 35.3 13.1c1.3.1 2.6.1 3.9.1v-85.3h-101v-35.7h101v-44.5H487c-22.7 24.1-43.5 44.1-43.5 44.1l-30.6-26.7c21.7-22.9 43.3-59.1 56.8-83.2c-10.9 4.4-22 9.2-33.6 14.2c-11.2 14.3-24.2 29-38.7 43.5c.5.8-50-28.4-50-28.4c52.2-44.4 81.4-139.9 81.4-139.9l72.5 20.4s-5.9 14-18.4 35.6c290.3-82.3 307.4 50.5 307.4 50.5s19.1 91.8 3.8 235.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TaobaoCircleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64M315.7 291.5c27.3 0 49.5 22.1 49.5 49.4s-22.1 49.4-49.5 49.4a49.4 49.4 0 1 1 0-98.8M366.9 578c-13.6 42.3-10.2 26.7-64.4 144.5l-78.5-49s87.7-79.8 105.6-116.2c19.2-38.4-21.1-58.9-21.1-58.9l-60.2-37.5l32.7-50.2c45.4 33.7 48.7 36.6 79.2 67.2c23.8 23.9 20.7 56.8 6.7 100.1m427.2 55c-15.3 143.8-202.4 90.3-202.4 90.3l10.2-41.1l43.3 9.3c80 5 72.3-64.9 72.3-64.9V423c.6-77.3-72.6-85.4-204.2-38.3l30.6 8.3c-2.5 9-12.5 23.2-25.2 38.6h176v35.6h-99.1v44.5h98.7v35.7h-98.7V622c14.9-4.8 28.6-11.5 40.5-20.5l-8.7-32.5l46.5-14.4l38.8 94.9l-57.3 23.9l-10.2-37.8c-25.6 19.5-78.8 48-171.8 45.4c-99.2 2.6-73.7-112-73.7-112l2.5-1.3H472c-.5 14.7-6.6 38.7 1.7 51.8c6.8 10.8 24.2 12.6 35.3 13.1c1.3.1 2.6.1 3.9.1v-85.3h-101v-35.7h101v-44.5H487c-22.7 24.1-43.5 44.1-43.5 44.1l-30.6-26.7c21.7-22.9 43.3-59.1 56.8-83.2c-10.9 4.4-22 9.2-33.6 14.2c-11.2 14.3-24.2 29-38.7 43.5c.5.8-50-28.4-50-28.4c52.2-44.4 81.4-139.9 81.4-139.9l72.5 20.4s-5.9 14-18.4 35.6c290.3-82.3 307.4 50.5 307.4 50.5s19.1 91.8 3.8 235.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TaobaoOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M168.5 273.7a68.7 68.7 0 1 0 137.4 0a68.7 68.7 0 1 0-137.4 0m730 79.2s-23.7-184.4-426.9-70.1c17.3-30 25.6-49.5 25.6-49.5L396.4 205s-40.6 132.6-113 194.4c0 0 70.1 40.6 69.4 39.4c20.1-20.1 38.2-40.6 53.7-60.4c16.1-7 31.5-13.6 46.7-19.8c-18.6 33.5-48.7 83.8-78.8 115.6l42.4 37s28.8-27.7 60.4-61.2h36v61.8H372.9v49.5h140.3v118.5c-1.7 0-3.6 0-5.4-.2c-15.4-.7-39.5-3.3-49-18.2c-11.5-18.1-3-51.5-2.4-71.9h-97l-3.4 1.8s-35.5 159.1 102.3 155.5c129.1 3.6 203-36 238.6-63.1l14.2 52.6l79.6-33.2l-53.9-131.9l-64.6 20.1l12.1 45.2c-16.6 12.4-35.6 21.7-56.2 28.4V561.3h137.1v-49.5H628.1V450h137.6v-49.5H521.3c17.6-21.4 31.5-41.1 35-53.6l-42.5-11.6c182.8-65.5 284.5-54.2 283.6 53.2v282.8s10.8 97.1-100.4 90.1l-60.2-12.9l-14.2 57.1S882.5 880 903.7 680.2c21.3-200-5.2-327.3-5.2-327.3m-707.4 18.3l-45.4 69.7l83.6 52.1s56 28.5 29.4 81.9C233.8 625.5 112 736.3 112 736.3l109 68.1c75.4-163.7 70.5-142 89.5-200.7c19.5-60.1 23.7-105.9-9.4-139.1c-42.4-42.6-47-46.6-110-93.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TaobaoSquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M315.7 291.5c27.3 0 49.5 22.1 49.5 49.4s-22.1 49.4-49.5 49.4a49.4 49.4 0 1 1 0-98.8M366.9 578c-13.6 42.3-10.2 26.7-64.4 144.5l-78.5-49s87.7-79.8 105.6-116.2c19.2-38.4-21.1-58.9-21.1-58.9l-60.2-37.5l32.7-50.2c45.4 33.7 48.7 36.6 79.2 67.2c23.8 23.9 20.7 56.8 6.7 100.1m427.2 55c-15.3 143.8-202.4 90.3-202.4 90.3l10.2-41.1l43.3 9.3c80 5 72.3-64.9 72.3-64.9V423c.6-77.3-72.6-85.4-204.2-38.3l30.6 8.3c-2.5 9-12.5 23.2-25.2 38.6h176v35.6h-99.1v44.5h98.7v35.7h-98.7V622c14.9-4.8 28.6-11.5 40.5-20.5l-8.7-32.5l46.5-14.4l38.8 94.9l-57.3 23.9l-10.2-37.8c-25.6 19.5-78.8 48-171.8 45.4c-99.2 2.6-73.7-112-73.7-112l2.5-1.3H472c-.5 14.7-6.6 38.7 1.7 51.8c6.8 10.8 24.2 12.6 35.3 13.1c1.3.1 2.6.1 3.9.1v-85.3h-101v-35.7h101v-44.5H487c-22.7 24.1-43.5 44.1-43.5 44.1l-30.6-26.7c21.7-22.9 43.3-59.1 56.8-83.2c-10.9 4.4-22 9.2-33.6 14.2c-11.2 14.3-24.2 29-38.7 43.5c.5.8-50-28.4-50-28.4c52.2-44.4 81.4-139.9 81.4-139.9l72.5 20.4s-5.9 14-18.4 35.6c290.3-82.3 307.4 50.5 307.4 50.5s19.1 91.8 3.8 235.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TeamOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M824.2 699.9a301.55 301.55 0 0 0-86.4-60.4C783.1 602.8 812 546.8 812 484c0-110.8-92.4-201.7-203.2-200c-109.1 1.7-197 90.6-197 200c0 62.8 29 118.8 74.2 155.5a300.95 300.95 0 0 0-86.4 60.4C345 754.6 314 826.8 312 903.8a8 8 0 0 0 8 8.2h56c4.3 0 7.9-3.4 8-7.7c1.9-58 25.4-112.3 66.7-153.5A226.62 226.62 0 0 1 612 684c60.9 0 118.2 23.7 161.3 66.8C814.5 792 838 846.3 840 904.3c.1 4.3 3.7 7.7 8 7.7h56a8 8 0 0 0 8-8.2c-2-77-33-149.2-87.8-203.9M612 612c-34.2 0-66.4-13.3-90.5-37.5a126.86 126.86 0 0 1-37.5-91.8c.3-32.8 13.4-64.5 36.3-88c24-24.6 56.1-38.3 90.4-38.7c33.9-.3 66.8 12.9 91 36.6c24.8 24.3 38.4 56.8 38.4 91.4c0 34.2-13.3 66.3-37.5 90.5A127.3 127.3 0 0 1 612 612M361.5 510.4c-.9-8.7-1.4-17.5-1.4-26.4c0-15.9 1.5-31.4 4.3-46.5c.7-3.6-1.2-7.3-4.5-8.8c-13.6-6.1-26.1-14.5-36.9-25.1a127.54 127.54 0 0 1-38.7-95.4c.9-32.1 13.8-62.6 36.3-85.6c24.7-25.3 57.9-39.1 93.2-38.7c31.9.3 62.7 12.6 86 34.4c7.9 7.4 14.7 15.6 20.4 24.4c2 3.1 5.9 4.4 9.3 3.2c17.6-6.1 36.2-10.4 55.3-12.4c5.6-.6 8.8-6.6 6.3-11.6c-32.5-64.3-98.9-108.7-175.7-109.9c-110.9-1.7-203.3 89.2-203.3 199.9c0 62.8 28.9 118.8 74.2 155.5c-31.8 14.7-61.1 35-86.5 60.4c-54.8 54.7-85.8 126.9-87.8 204a8 8 0 0 0 8 8.2h56.1c4.3 0 7.9-3.4 8-7.7c1.9-58 25.4-112.3 66.7-153.5c29.4-29.4 65.4-49.8 104.7-59.7c3.9-1 6.5-4.7 6-8.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThunderboltFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M848 359.3H627.7L825.8 109c4.1-5.3.4-13-6.3-13H436c-2.8 0-5.5 1.5-6.9 4L170 547.5c-3.1 5.3.7 12 6.9 12h174.4l-89.4 357.6c-1.9 7.8 7.5 13.3 13.3 7.7L853.5 373c5.2-4.9 1.7-13.7-5.5-13.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThunderboltOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M848 359.3H627.7L825.8 109c4.1-5.3.4-13-6.3-13H436c-2.8 0-5.5 1.5-6.9 4L170 547.5c-3.1 5.3.7 12 6.9 12h174.4l-89.4 357.6c-1.9 7.8 7.5 13.3 13.3 7.7L853.5 373c5.2-4.9 1.7-13.7-5.5-13.7M378.2 732.5l60.3-241H281.1l189.6-327.4h224.6L487 427.4h211z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThunderboltTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M695.4 164.1H470.8L281.2 491.5h157.4l-60.3 241l319.8-305.1h-211z"/><path fill="currentColor" d="M848.1 359.3H627.8L825.9 109c4.1-5.3.4-13-6.3-13H436.1c-2.8 0-5.5 1.5-6.9 4L170.1 547.5c-3.1 5.3.7 12 6.9 12h174.4L262 917.1c-1.9 7.8 7.5 13.3 13.3 7.7L853.6 373c5.2-4.9 1.7-13.7-5.5-13.7M378.3 732.5l60.3-241H281.2l189.6-327.4h224.6L487.1 427.4h211z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToTopOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M885 780H165c-4.4 0-8 3.6-8 8v60c0 4.4 3.6 8 8 8h720c4.4 0 8-3.6 8-8v-60c0-4.4-3.6-8-8-8M400 325.7h73.9V664c0 4.4 3.6 8 8 8h60c4.4 0 8-3.6 8-8V325.7H624c6.7 0 10.4-7.7 6.3-12.9L518.3 171a8 8 0 0 0-12.6 0l-112 141.7c-4.1 5.3-.4 13 6.3 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToolFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M865.3 244.7c-.3-.3-61.1 59.8-182.1 180.6l-84.9-84.9l180.9-180.9c-95.2-57.3-217.5-42.6-296.8 36.7A244.42 244.42 0 0 0 419 432l1.8 6.7l-283.5 283.4c-6.2 6.2-6.2 16.4 0 22.6l141.4 141.4c6.2 6.2 16.4 6.2 22.6 0l283.3-283.3l6.7 1.8c83.7 22.3 173.6-.9 236-63.3c79.4-79.3 94.1-201.6 38-296.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToolOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M876.6 239.5c-.5-.9-1.2-1.8-2-2.5c-5-5-13.1-5-18.1 0L684.2 409.3l-67.9-67.9L788.7 169c.8-.8 1.4-1.6 2-2.5c3.6-6.1 1.6-13.9-4.5-17.5c-98.2-58-226.8-44.7-311.3 39.7c-67 67-89.2 162-66.5 247.4l-293 293c-3 3-2.8 7.9.3 11l169.7 169.7c3.1 3.1 8.1 3.3 11 .3l292.9-292.9c85.5 22.8 180.5.7 247.6-66.4c84.4-84.5 97.7-213.1 39.7-311.3M786 499.8c-58.1 58.1-145.3 69.3-214.6 33.6l-8.8 8.8l-.1-.1l-274 274.1l-79.2-79.2l230.1-230.1s0 .1.1.1l52.8-52.8c-35.7-69.3-24.5-156.5 33.6-214.6a184.2 184.2 0 0 1 144-53.5L537 318.9a32.05 32.05 0 0 0 0 45.3l124.5 124.5a32.05 32.05 0 0 0 45.3 0l132.8-132.8c3.7 51.8-14.4 104.8-53.6 143.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToolTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M706.8 488.7a32.05 32.05 0 0 1-45.3 0L537 364.2a32.05 32.05 0 0 1 0-45.3l132.9-132.8a184.2 184.2 0 0 0-144 53.5c-58.1 58.1-69.3 145.3-33.6 214.6L439.5 507c-.1 0-.1-.1-.1-.1L209.3 737l79.2 79.2l274-274.1l.1.1l8.8-8.8c69.3 35.7 156.5 24.5 214.6-33.6c39.2-39.1 57.3-92.1 53.6-143.9z"/><path fill="currentColor" d="M876.6 239.5c-.5-.9-1.2-1.8-2-2.5c-5-5-13.1-5-18.1 0L684.2 409.3l-67.9-67.9L788.7 169c.8-.8 1.4-1.6 2-2.5c3.6-6.1 1.6-13.9-4.5-17.5c-98.2-58-226.8-44.7-311.3 39.7c-67 67-89.2 162-66.5 247.4l-293 293c-3 3-2.8 7.9.3 11l169.7 169.7c3.1 3.1 8.1 3.3 11 .3l292.9-292.9c85.5 22.8 180.5.7 247.6-66.4c84.4-84.5 97.7-213.1 39.7-311.3M786 499.8c-58.1 58.1-145.3 69.3-214.6 33.6l-8.8 8.8l-.1-.1l-274 274.1l-79.2-79.2l230.1-230.1s0 .1.1.1l52.8-52.8c-35.7-69.3-24.5-156.5 33.6-214.6a184.2 184.2 0 0 1 144-53.5L537 318.9a32.05 32.05 0 0 0 0 45.3l124.5 124.5a32.05 32.05 0 0 0 45.3 0l132.8-132.8c3.7 51.8-14.4 104.8-53.6 143.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrademarkCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m164.7 660.2c-1.1.5-2.3.8-3.5.8h-62c-3.1 0-5.9-1.8-7.2-4.6l-74.6-159.2h-88.7V717c0 4.4-3.6 8-8 8H378c-4.4 0-8-3.6-8-8V307c0-4.4 3.6-8 8-8h155.6c98.8 0 144.2 59.9 144.2 131.1c0 70.2-43.6 106.4-78.4 119.2l80.8 164.2c2.1 3.9.4 8.7-3.5 10.7M523.9 357h-83.4v148H522c53 0 82.8-25.6 82.8-72.4c0-50.3-32.9-75.6-80.9-75.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrademarkCircleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372m87.5-334.7c34.8-12.8 78.4-49 78.4-119.2c0-71.2-45.5-131.1-144.2-131.1H378c-4.4 0-8 3.6-8 8v410c0 4.4 3.6 8 8 8h54.5c4.4 0 8-3.6 8-8V561.2h88.7l74.6 159.2c1.3 2.8 4.1 4.6 7.2 4.6h62a7.9 7.9 0 0 0 7.1-11.5zM522 505h-81.5V357h83.4c48 0 80.9 25.3 80.9 75.5c0 46.9-29.8 72.5-82.8 72.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrademarkCircleTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372m170.7 584.2c-1.1.5-2.3.8-3.5.8h-62c-3.1 0-5.9-1.8-7.2-4.6l-74.6-159.2h-88.7V717c0 4.4-3.6 8-8 8H384c-4.4 0-8-3.6-8-8V307c0-4.4 3.6-8 8-8h155.6c98.8 0 144.2 59.9 144.2 131.1c0 70.2-43.6 106.4-78.4 119.2l80.8 164.2c2.1 3.9.4 8.7-3.5 10.7"/><path fill="currentColor" fill-opacity=".15" d="M529.9 357h-83.4v148H528c53 0 82.8-25.6 82.8-72.4c0-50.3-32.9-75.6-80.9-75.6"/><path fill="currentColor" d="M605.4 549.3c34.8-12.8 78.4-49 78.4-119.2c0-71.2-45.4-131.1-144.2-131.1H384c-4.4 0-8 3.6-8 8v410c0 4.4 3.6 8 8 8h54.7c4.4 0 8-3.6 8-8V561.2h88.7L610 720.4c1.3 2.8 4.1 4.6 7.2 4.6h62c1.2 0 2.4-.3 3.5-.8c3.9-2 5.6-6.8 3.5-10.7zM528 505h-81.5V357h83.4c48 0 80.9 25.3 80.9 75.6c0 46.8-29.8 72.4-82.8 72.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrademarkOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372m87.5-334.7c34.8-12.8 78.4-49 78.4-119.2c0-71.2-45.5-131.1-144.2-131.1H378c-4.4 0-8 3.6-8 8v410c0 4.4 3.6 8 8 8h54.5c4.4 0 8-3.6 8-8V561.2h88.7l74.6 159.2c1.3 2.8 4.1 4.6 7.2 4.6h62a7.9 7.9 0 0 0 7.1-11.5zM522 505h-81.5V357h83.4c48 0 80.9 25.3 80.9 75.5c0 46.9-29.8 72.5-82.8 72.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransactionOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M668.6 320c0-4.4-3.6-8-8-8h-54.5c-3 0-5.8 1.7-7.1 4.4l-84.7 168.8H511l-84.7-168.8a8 8 0 0 0-7.1-4.4h-55.7c-1.3 0-2.6.3-3.8 1c-3.9 2.1-5.3 7-3.2 10.8l103.9 191.6h-57c-4.4 0-8 3.6-8 8v27.1c0 4.4 3.6 8 8 8h76v39h-76c-4.4 0-8 3.6-8 8v27.1c0 4.4 3.6 8 8 8h76V704c0 4.4 3.6 8 8 8h49.9c4.4 0 8-3.6 8-8v-63.5h76.3c4.4 0 8-3.6 8-8v-27.1c0-4.4-3.6-8-8-8h-76.3v-39h76.3c4.4 0 8-3.6 8-8v-27.1c0-4.4-3.6-8-8-8H564l103.7-191.6c.5-1.1.9-2.4.9-3.7M157.9 504.2a352.7 352.7 0 0 1 103.5-242.4c32.5-32.5 70.3-58.1 112.4-75.9c43.6-18.4 89.9-27.8 137.6-27.8c47.8 0 94.1 9.3 137.6 27.8c42.1 17.8 79.9 43.4 112.4 75.9c10 10 19.3 20.5 27.9 31.4l-50 39.1a8 8 0 0 0 3 14.1l156.8 38.3c5 1.2 9.9-2.6 9.9-7.7l.8-161.5c0-6.7-7.7-10.5-12.9-6.3l-47.8 37.4C770.7 146.3 648.6 82 511.5 82C277 82 86.3 270.1 82 503.8a8 8 0 0 0 8 8.2h60c4.3 0 7.8-3.5 7.9-7.8M934 512h-60c-4.3 0-7.9 3.5-8 7.8a352.7 352.7 0 0 1-103.5 242.4a352.57 352.57 0 0 1-112.4 75.9c-43.6 18.4-89.9 27.8-137.6 27.8s-94.1-9.3-137.6-27.8a352.57 352.57 0 0 1-112.4-75.9c-10-10-19.3-20.5-27.9-31.4l49.9-39.1a8 8 0 0 0-3-14.1l-156.8-38.3c-5-1.2-9.9 2.6-9.9 7.7l-.8 161.7c0 6.7 7.7 10.5 12.9 6.3l47.8-37.4C253.3 877.7 375.4 942 512.5 942C747 942 937.7 753.9 942 520.2a8 8 0 0 0-8-8.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TranslationOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M140 188h584v164h76V144c0-17.7-14.3-32-32-32H96c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h544v-76H140z"/><path fill="currentColor" d="M414.3 256h-60.6c-3.4 0-6.4 2.2-7.6 5.4L219 629.4c-.3.8-.4 1.7-.4 2.6c0 4.4 3.6 8 8 8h55.1c3.4 0 6.4-2.2 7.6-5.4L322 540h196.2L422 261.4c-1.3-3.2-4.3-5.4-7.7-5.4m12.4 228h-85.5L384 360.2zM936 528H800v-93c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v93H592c-13.3 0-24 10.7-24 24v176c0 13.3 10.7 24 24 24h136v152c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V752h136c13.3 0 24-10.7 24-24V552c0-13.3-10.7-24-24-24M728 680h-88v-80h88zm160 0h-88v-80h88z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrophyFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M868 160h-92v-40c0-4.4-3.6-8-8-8H256c-4.4 0-8 3.6-8 8v40h-92a44 44 0 0 0-44 44v148c0 81.7 60 149.6 138.2 162C265.6 630.2 359 721.8 476 734.5v105.2H280c-17.7 0-32 14.3-32 32V904c0 4.4 3.6 8 8 8h512c4.4 0 8-3.6 8-8v-32.3c0-17.7-14.3-32-32-32H548V734.5C665 721.8 758.4 630.2 773.8 514C852 501.6 912 433.7 912 352V204a44 44 0 0 0-44-44M248 439.6c-37.1-11.9-64-46.7-64-87.6V232h64zM840 352c0 41-26.9 75.8-64 87.6V232h64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrophyOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M868 160h-92v-40c0-4.4-3.6-8-8-8H256c-4.4 0-8 3.6-8 8v40h-92a44 44 0 0 0-44 44v148c0 81.7 60 149.6 138.2 162C265.7 630.2 359 721.7 476 734.5v105.2H280c-17.7 0-32 14.3-32 32V904c0 4.4 3.6 8 8 8h512c4.4 0 8-3.6 8-8v-32.3c0-17.7-14.3-32-32-32H548V734.5C665 721.7 758.3 630.2 773.8 514C852 501.6 912 433.7 912 352V204a44 44 0 0 0-44-44M184 352V232h64v207.6a91.99 91.99 0 0 1-64-87.6m520 128c0 49.1-19.1 95.4-53.9 130.1c-34.8 34.8-81 53.9-130.1 53.9h-16c-49.1 0-95.4-19.1-130.1-53.9c-34.8-34.8-53.9-81-53.9-130.1V184h384zm136-128c0 41-26.9 75.8-64 87.6V232h64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrophyTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M320 480c0 49.1 19.1 95.3 53.9 130.1c34.7 34.8 81 53.9 130.1 53.9h16c49.1 0 95.3-19.1 130.1-53.9c34.8-34.7 53.9-81 53.9-130.1V184H320zM184 352c0 41 26.9 75.8 64 87.6c-37.1-11.9-64-46.7-64-87.6m364 382.5C665 721.8 758.4 630.2 773.8 514C758.3 630.2 665 721.7 548 734.5M250.2 514C265.6 630.2 359 721.8 476 734.5C359 721.7 265.7 630.2 250.2 514"/><path fill="currentColor" d="M868 160h-92v-40c0-4.4-3.6-8-8-8H256c-4.4 0-8 3.6-8 8v40h-92a44 44 0 0 0-44 44v148c0 81.7 60 149.6 138.2 162C265.7 630.2 359 721.7 476 734.5v105.2H280c-17.7 0-32 14.3-32 32V904c0 4.4 3.6 8 8 8h512c4.4 0 8-3.6 8-8v-32.3c0-17.7-14.3-32-32-32H548V734.5C665 721.7 758.3 630.2 773.8 514C852 501.6 912 433.7 912 352V204a44 44 0 0 0-44-44M248 439.6a91.99 91.99 0 0 1-64-87.6V232h64zM704 480c0 49.1-19.1 95.4-53.9 130.1c-34.8 34.8-81 53.9-130.1 53.9h-16c-49.1 0-95.4-19.1-130.1-53.9c-34.8-34.8-53.9-81-53.9-130.1V184h384zm136-128c0 41-26.9 75.8-64 87.6V232h64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwitterCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m215.3 337.7c.3 4.7.3 9.6.3 14.4c0 146.8-111.8 315.9-316.1 315.9c-63 0-121.4-18.3-170.6-49.8c9 1 17.6 1.4 26.8 1.4c52 0 99.8-17.6 137.9-47.4c-48.8-1-89.8-33-103.8-77c17.1 2.5 32.5 2.5 50.1-2a111 111 0 0 1-88.9-109v-1.4c14.7 8.3 32 13.4 50.1 14.1a111.13 111.13 0 0 1-49.5-92.4c0-20.7 5.4-39.6 15.1-56a315.28 315.28 0 0 0 229 116.1C492 353.1 548.4 292 616.2 292c32 0 60.8 13.4 81.1 35c25.1-4.7 49.1-14.1 70.5-26.7c-8.3 25.7-25.7 47.4-48.8 61.1c22.4-2.4 44-8.6 64-17.3c-15.1 22.2-34 41.9-55.7 57.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwitterOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M928 254.3c-30.6 13.2-63.9 22.7-98.2 26.4a170.1 170.1 0 0 0 75-94a336.64 336.64 0 0 1-108.2 41.2A170.1 170.1 0 0 0 672 174c-94.5 0-170.5 76.6-170.5 170.6c0 13.2 1.6 26.4 4.2 39.1c-141.5-7.4-267.7-75-351.6-178.5a169.32 169.32 0 0 0-23.2 86.1c0 59.2 30.1 111.4 76 142.1a172 172 0 0 1-77.1-21.7v2.1c0 82.9 58.6 151.6 136.7 167.4a180.6 180.6 0 0 1-44.9 5.8c-11.1 0-21.6-1.1-32.2-2.6C211 652 273.9 701.1 348.8 702.7c-58.6 45.9-132 72.9-211.7 72.9c-14.3 0-27.5-.5-41.2-2.1C171.5 822 261.2 850 357.8 850C671.4 850 843 590.2 843 364.7c0-7.4 0-14.8-.5-22.2c33.2-24.3 62.3-54.4 85.5-88.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwitterSquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M727.3 401.7c.3 4.7.3 9.6.3 14.4c0 146.8-111.8 315.9-316.1 315.9c-63 0-121.4-18.3-170.6-49.8c9 1 17.6 1.4 26.8 1.4c52 0 99.8-17.6 137.9-47.4c-48.8-1-89.8-33-103.8-77c17.1 2.5 32.5 2.5 50.1-2a111 111 0 0 1-88.9-109v-1.4c14.7 8.3 32 13.4 50.1 14.1a111.13 111.13 0 0 1-49.5-92.4c0-20.7 5.4-39.6 15.1-56a315.28 315.28 0 0 0 229 116.1C492 353.1 548.4 292 616.2 292c32 0 60.8 13.4 81.1 35c25.1-4.7 49.1-14.1 70.5-26.7c-8.3 25.7-25.7 47.4-48.8 61.1c22.4-2.4 44-8.6 64-17.3c-15.1 22.2-34 41.9-55.7 57.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnderlineOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M824 804H200c-4.4 0-8 3.4-8 7.6v60.8c0 4.2 3.6 7.6 8 7.6h624c4.4 0 8-3.4 8-7.6v-60.8c0-4.2-3.6-7.6-8-7.6m-312-76c69.4 0 134.6-27.1 183.8-76.2C745 602.7 772 537.4 772 468V156c0-6.6-5.4-12-12-12h-60c-6.6 0-12 5.4-12 12v312c0 97-79 176-176 176s-176-79-176-176V156c0-6.6-5.4-12-12-12h-60c-6.6 0-12 5.4-12 12v312c0 69.4 27.1 134.6 76.2 183.8C377.3 701 442.6 728 512 728"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UndoOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M511.4 124C290.5 124.3 112 303 112 523.9c0 128 60.2 242 153.8 315.2l-37.5 48c-4.1 5.3-.3 13 6.3 12.9l167-.8c5.2 0 9-4.9 7.7-9.9L369.8 727a8 8 0 0 0-14.1-3L315 776.1c-10.2-8-20-16.7-29.3-26a318.64 318.64 0 0 1-68.6-101.7C200.4 609 192 567.1 192 523.9s8.4-85.1 25.1-124.5c16.1-38.1 39.2-72.3 68.6-101.7c29.4-29.4 63.6-52.5 101.7-68.6C426.9 212.4 468.8 204 512 204s85.1 8.4 124.5 25.1c38.1 16.1 72.3 39.2 101.7 68.6c29.4 29.4 52.5 63.6 68.6 101.7c16.7 39.4 25.1 81.3 25.1 124.5s-8.4 85.1-25.1 124.5a318.64 318.64 0 0 1-68.6 101.7c-7.5 7.5-15.3 14.5-23.4 21.2a7.93 7.93 0 0 0-1.2 11.1l39.4 50.5c2.8 3.5 7.9 4.1 11.4 1.3C854.5 760.8 912 649.1 912 523.9c0-221.1-179.4-400.2-400.6-399.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UngroupOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M736 550H288c-8.8 0-16 7.2-16 16v176c0 8.8 7.2 16 16 16h448c8.8 0 16-7.2 16-16V566c0-8.8-7.2-16-16-16m-56 136H344v-64h336zm208 130c-39.8 0-72 32.2-72 72s32.2 72 72 72s72-32.2 72-72s-32.2-72-72-72m0 96c-13.3 0-24-10.7-24-24s10.7-24 24-24s24 10.7 24 24s-10.7 24-24 24M736 266H288c-8.8 0-16 7.2-16 16v176c0 8.8 7.2 16 16 16h448c8.8 0 16-7.2 16-16V282c0-8.8-7.2-16-16-16m-56 136H344v-64h336zm208-194c39.8 0 72-32.2 72-72s-32.2-72-72-72s-72 32.2-72 72s32.2 72 72 72m0-96c13.3 0 24 10.7 24 24s-10.7 24-24 24s-24-10.7-24-24s10.7-24 24-24M136 64c-39.8 0-72 32.2-72 72s32.2 72 72 72s72-32.2 72-72s-32.2-72-72-72m0 96c-13.3 0-24-10.7-24-24s10.7-24 24-24s24 10.7 24 24s-10.7 24-24 24m0 656c-39.8 0-72 32.2-72 72s32.2 72 72 72s72-32.2 72-72s-32.2-72-72-72m0 96c-13.3 0-24-10.7-24-24s10.7-24 24-24s24 10.7 24 24s-10.7 24-24 24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnlockFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 464H332V240c0-30.9 25.1-56 56-56h248c30.9 0 56 25.1 56 56v68c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-68c0-70.7-57.3-128-128-128H388c-70.7 0-128 57.3-128 128v224h-68c-17.7 0-32 14.3-32 32v384c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V496c0-17.7-14.3-32-32-32M540 701v53c0 4.4-3.6 8-8 8h-40c-4.4 0-8-3.6-8-8v-53a48.01 48.01 0 1 1 56 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnlockOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 464H332V240c0-30.9 25.1-56 56-56h248c30.9 0 56 25.1 56 56v68c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-68c0-70.7-57.3-128-128-128H388c-70.7 0-128 57.3-128 128v224h-68c-17.7 0-32 14.3-32 32v384c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V496c0-17.7-14.3-32-32-32m-40 376H232V536h560zM484 701v53c0 4.4 3.6 8 8 8h40c4.4 0 8-3.6 8-8v-53a48.01 48.01 0 1 0-56 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnlockTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M232 840h560V536H232zm280-226a48.01 48.01 0 0 1 28 87v53c0 4.4-3.6 8-8 8h-40c-4.4 0-8-3.6-8-8v-53a48.01 48.01 0 0 1 28-87"/><path fill="currentColor" d="M484 701v53c0 4.4 3.6 8 8 8h40c4.4 0 8-3.6 8-8v-53a48.01 48.01 0 1 0-56 0"/><path fill="currentColor" d="M832 464H332V240c0-30.9 25.1-56 56-56h248c30.9 0 56 25.1 56 56v68c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-68c0-70.7-57.3-128-128-128H388c-70.7 0-128 57.3-128 128v224h-68c-17.7 0-32 14.3-32 32v384c0 17.7 14.3 32 32 32h640c17.7 0 32-14.3 32-32V496c0-17.7-14.3-32-32-32m-40 376H232V536h560z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnorderedListOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M912 192H328c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h584c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 284H328c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h584c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8m0 284H328c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h584c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8M104 228a56 56 0 1 0 112 0a56 56 0 1 0-112 0m0 284a56 56 0 1 0 112 0a56 56 0 1 0-112 0m0 284a56 56 0 1 0 112 0a56 56 0 1 0-112 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m178 555h-46.9c-10.2 0-19.9-4.9-25.9-13.2L512 460.4L406.8 605.8c-6 8.3-15.6 13.2-25.9 13.2H334c-6.5 0-10.3-7.4-6.5-12.7l178-246c3.2-4.4 9.7-4.4 12.9 0l178 246c3.9 5.3.1 12.7-6.4 12.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpCircleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M518.5 360.3a7.95 7.95 0 0 0-12.9 0l-178 246c-3.8 5.3 0 12.7 6.5 12.7H381c10.2 0 19.9-4.9 25.9-13.2L512 460.4l105.2 145.4c6 8.3 15.6 13.2 25.9 13.2H690c6.5 0 10.3-7.4 6.5-12.7z"/><path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpCircleTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M512 140c-205.4 0-372 166.6-372 372s166.6 372 372 372s372-166.6 372-372s-166.6-372-372-372m178 479h-46.9c-10.2 0-19.9-4.9-25.9-13.2L512 460.4L406.8 605.8c-6 8.3-15.6 13.2-25.9 13.2H334c-6.5 0-10.3-7.4-6.5-12.7l178-246c3.2-4.4 9.7-4.4 12.9 0l178 246c3.9 5.3.1 12.7-6.4 12.7"/><path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372s372 166.6 372 372s-166.6 372-372 372"/><path fill="currentColor" d="M518.4 360.3a7.95 7.95 0 0 0-12.9 0l-178 246c-3.8 5.3 0 12.7 6.5 12.7h46.9c10.3 0 19.9-4.9 25.9-13.2L512 460.4l105.2 145.4c6 8.3 15.7 13.2 25.9 13.2H690c6.5 0 10.3-7.4 6.4-12.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M890.5 755.3L537.9 269.2c-12.8-17.6-39-17.6-51.7 0L133.5 755.3A8 8 0 0 0 140 768h75c5.1 0 9.9-2.5 12.9-6.6L512 369.8l284.1 391.6c3 4.1 7.8 6.6 12.9 6.6h75c6.5 0 10.3-7.4 6.5-12.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpSquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M690 624h-46.9c-10.2 0-19.9-4.9-25.9-13.2L512 465.4L406.8 610.8c-6 8.3-15.6 13.2-25.9 13.2H334c-6.5 0-10.3-7.4-6.5-12.7l178-246c3.2-4.4 9.7-4.4 12.9 0l178 246c3.9 5.3.1 12.7-6.4 12.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpSquareOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M334 624h46.9c10.2 0 19.9-4.9 25.9-13.2L512 465.4l105.2 145.4c6 8.3 15.6 13.2 25.9 13.2H690c6.5 0 10.3-7.4 6.5-12.7l-178-246a7.95 7.95 0 0 0-12.9 0l-178 246A7.96 7.96 0 0 0 334 624"/><path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpSquareTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 728H184V184h656z"/><path fill="currentColor" fill-opacity=".15" d="M184 840h656V184H184zm143.5-228.7l178-246c3.2-4.4 9.7-4.4 12.9 0l178 246c3.9 5.3.1 12.7-6.4 12.7h-46.9c-10.2 0-19.9-4.9-25.9-13.2L512 465.4L406.8 610.8c-6 8.3-15.6 13.2-25.9 13.2H334c-6.5 0-10.3-7.4-6.5-12.7"/><path fill="currentColor" d="M334 624h46.9c10.3 0 19.9-4.9 25.9-13.2L512 465.4l105.2 145.4c6 8.3 15.7 13.2 25.9 13.2H690c6.5 0 10.3-7.4 6.4-12.7l-178-246a7.95 7.95 0 0 0-12.9 0l-178 246c-3.8 5.3 0 12.7 6.5 12.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 317.7h73.9V656c0 4.4 3.6 8 8 8h60c4.4 0 8-3.6 8-8V317.7H624c6.7 0 10.4-7.7 6.3-12.9L518.3 163a8 8 0 0 0-12.6 0l-112 141.7c-4.1 5.3-.4 13 6.3 13M878 626h-60c-4.4 0-8 3.6-8 8v154H214V634c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8v198c0 17.7 14.3 32 32 32h684c17.7 0 32-14.3 32-32V634c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UsbFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M408 312h48c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8m352 120V144c0-17.7-14.3-32-32-32H296c-17.7 0-32 14.3-32 32v288c-66.2 0-120 52.1-120 116v356c0 4.4 3.6 8 8 8h720c4.4 0 8-3.6 8-8V548c0-63.9-53.8-116-120-116m-72 0H336V184h352zM568 312h48c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UsbOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M760 432V144c0-17.7-14.3-32-32-32H296c-17.7 0-32 14.3-32 32v288c-66.2 0-120 52.1-120 116v356c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V548c0-24.3 21.6-44 48.1-44h495.8c26.5 0 48.1 19.7 48.1 44v356c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V548c0-63.9-53.8-116-120-116m-424 0V184h352v248zm120-184h-48c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8m160 0h-48c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UsbTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M759.9 504H264.1c-26.5 0-48.1 19.7-48.1 44v292h592V548c0-24.3-21.6-44-48.1-44"/><path fill="currentColor" d="M456 248h-48c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8m160 0h-48c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8"/><path fill="currentColor" d="M760 432V144c0-17.7-14.3-32-32-32H296c-17.7 0-32 14.3-32 32v288c-66.2 0-120 52.1-120 116v356c0 4.4 3.6 8 8 8h720c4.4 0 8-3.6 8-8V548c0-63.9-53.8-116-120-116M336 184h352v248H336zm472 656H216V548c0-24.3 21.6-44 48.1-44h495.8c26.5 0 48.1 19.7 48.1 44z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserAddOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M678.3 642.4c24.2-13 51.9-20.4 81.4-20.4h.1c3 0 4.4-3.6 2.2-5.6a371.67 371.67 0 0 0-103.7-65.8c-.4-.2-.8-.3-1.2-.5C719.2 505 759.6 431.7 759.6 349c0-137-110.8-248-247.5-248S264.7 212 264.7 349c0 82.7 40.4 156 102.6 201.1c-.4.2-.8.3-1.2.5c-44.7 18.9-84.8 46-119.3 80.6a373.42 373.42 0 0 0-80.4 119.5A373.6 373.6 0 0 0 137 888.8a8 8 0 0 0 8 8.2h59.9c4.3 0 7.9-3.5 8-7.8c2-77.2 32.9-149.5 87.6-204.3C357 628.2 432.2 597 512.2 597c56.7 0 111.1 15.7 158 45.1a8.1 8.1 0 0 0 8.1.3M512.2 521c-45.8 0-88.9-17.9-121.4-50.4A171.2 171.2 0 0 1 340.5 349c0-45.9 17.9-89.1 50.3-121.6S466.3 177 512.2 177s88.9 17.9 121.4 50.4A171.2 171.2 0 0 1 683.9 349c0 45.9-17.9 89.1-50.3 121.6C601.1 503.1 558 521 512.2 521M880 759h-84v-84c0-4.4-3.6-8-8-8h-56c-4.4 0-8 3.6-8 8v84h-84c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h84v84c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8v-84h84c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserDeleteOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M678.3 655.4c24.2-13 51.9-20.4 81.4-20.4h.1c3 0 4.4-3.6 2.2-5.6a371.67 371.67 0 0 0-103.7-65.8c-.4-.2-.8-.3-1.2-.5C719.2 518 759.6 444.7 759.6 362c0-137-110.8-248-247.5-248S264.7 225 264.7 362c0 82.7 40.4 156 102.6 201.1c-.4.2-.8.3-1.2.5c-44.7 18.9-84.8 46-119.3 80.6a373.42 373.42 0 0 0-80.4 119.5A373.6 373.6 0 0 0 137 901.8a8 8 0 0 0 8 8.2h59.9c4.3 0 7.9-3.5 8-7.8c2-77.2 32.9-149.5 87.6-204.3C357 641.2 432.2 610 512.2 610c56.7 0 111.1 15.7 158 45.1a8.1 8.1 0 0 0 8.1.3M512.2 534c-45.8 0-88.9-17.9-121.4-50.4A171.2 171.2 0 0 1 340.5 362c0-45.9 17.9-89.1 50.3-121.6S466.3 190 512.2 190s88.9 17.9 121.4 50.4A171.2 171.2 0 0 1 683.9 362c0 45.9-17.9 89.1-50.3 121.6C601.1 516.1 558 534 512.2 534M880 772H640c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h240c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M858.5 763.6a374 374 0 0 0-80.6-119.5a375.63 375.63 0 0 0-119.5-80.6c-.4-.2-.8-.3-1.2-.5C719.5 518 760 444.7 760 362c0-137-111-248-248-248S264 225 264 362c0 82.7 40.5 156 102.8 201.1c-.4.2-.8.3-1.2.5c-44.8 18.9-85 46-119.5 80.6a375.63 375.63 0 0 0-80.6 119.5A371.7 371.7 0 0 0 136 901.8a8 8 0 0 0 8 8.2h60c4.4 0 7.9-3.5 8-7.8c2-77.2 33-149.5 87.8-204.3c56.7-56.7 132-87.9 212.2-87.9s155.5 31.2 212.2 87.9C779 752.7 810 825 812 902.2c.1 4.4 3.6 7.8 8 7.8h60a8 8 0 0 0 8-8.2c-1-47.8-10.9-94.3-29.5-138.2M512 534c-45.9 0-89.1-17.9-121.6-50.4S340 407.9 340 362c0-45.9 17.9-89.1 50.4-121.6S466.1 190 512 190s89.1 17.9 121.6 50.4S684 316.1 684 362c0 45.9-17.9 89.1-50.4 121.6S557.9 534 512 534"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserSwitchOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M759 335c0-137-111-248-248-248S263 198 263 335c0 82.8 40.6 156.2 103 201.2c-.4.2-.7.3-.9.4c-44.7 18.9-84.8 46-119.3 80.6c-34.5 34.5-61.5 74.7-80.4 119.5C146.9 780.5 137 827 136 874.8c-.1 4.5 3.5 8.2 8 8.2h59.9c4.3 0 7.9-3.5 8-7.8c2-77.2 32.9-149.5 87.6-204.3C356 614.2 431 583 511 583c137 0 248-111 248-248M511 507c-95 0-172-77-172-172s77-172 172-172s172 77 172 172s-77 172-172 172m105 221h264c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8H703.5l47.2-60.1c1.1-1.4 1.7-3.2 1.7-4.9c0-4.4-3.6-8-8-8h-72.6c-4.9 0-9.5 2.3-12.6 6.1l-68.5 87.1c-4.4 5.6-6.8 12.6-6.8 19.8c.1 17.7 14.4 32 32.1 32m240 64H592c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h176.5l-47.2 60.1c-1.1 1.4-1.7 3.2-1.7 4.9c0 4.4 3.6 8 8 8h72.6c4.9 0 9.5-2.3 12.6-6.1l68.5-87.1c4.4-5.6 6.8-12.6 6.8-19.8c-.1-17.7-14.4-32-32.1-32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UsergroupAddOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M892 772h-80v-80c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v80h-80c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h80v80c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8v-80h80c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8M373.5 498.4c-.9-8.7-1.4-17.5-1.4-26.4c0-15.9 1.5-31.4 4.3-46.5c.7-3.6-1.2-7.3-4.5-8.8c-13.6-6.1-26.1-14.5-36.9-25.1a127.54 127.54 0 0 1-38.7-95.4c.9-32.1 13.8-62.6 36.3-85.6c24.7-25.3 57.9-39.1 93.2-38.7c31.9.3 62.7 12.6 86 34.4c7.9 7.4 14.7 15.6 20.4 24.4c2 3.1 5.9 4.4 9.3 3.2c17.6-6.1 36.2-10.4 55.3-12.4c5.6-.6 8.8-6.6 6.3-11.6c-32.5-64.3-98.9-108.7-175.7-109.9c-110.8-1.7-203.2 89.2-203.2 200c0 62.8 28.9 118.8 74.2 155.5c-31.8 14.7-61.1 35-86.5 60.4c-54.8 54.7-85.8 126.9-87.8 204a8 8 0 0 0 8 8.2h56.1c4.3 0 7.9-3.4 8-7.7c1.9-58 25.4-112.3 66.7-153.5c29.4-29.4 65.4-49.8 104.7-59.7c3.8-1.1 6.4-4.8 5.9-8.8M824 472c0-109.4-87.9-198.3-196.9-200C516.3 270.3 424 361.2 424 472c0 62.8 29 118.8 74.2 155.5a300.95 300.95 0 0 0-86.4 60.4C357 742.6 326 814.8 324 891.8a8 8 0 0 0 8 8.2h56c4.3 0 7.9-3.4 8-7.7c1.9-58 25.4-112.3 66.7-153.5C505.8 695.7 563 672 624 672c110.4 0 200-89.5 200-200m-109.5 90.5C690.3 586.7 658.2 600 624 600s-66.3-13.3-90.5-37.5a127.26 127.26 0 0 1-37.5-91.8c.3-32.8 13.4-64.5 36.3-88c24-24.6 56.1-38.3 90.4-38.7c33.9-.3 66.8 12.9 91 36.6c24.8 24.3 38.4 56.8 38.4 91.4c-.1 34.2-13.4 66.3-37.6 90.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UsergroupDeleteOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M888 784H664c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h224c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8M373.5 510.4c-.9-8.7-1.4-17.5-1.4-26.4c0-15.9 1.5-31.4 4.3-46.5c.7-3.6-1.2-7.3-4.5-8.8c-13.6-6.1-26.1-14.5-36.9-25.1a127.54 127.54 0 0 1-38.7-95.4c.9-32.1 13.8-62.6 36.3-85.6c24.7-25.3 57.9-39.1 93.2-38.7c31.9.3 62.7 12.6 86 34.4c7.9 7.4 14.7 15.6 20.4 24.4c2 3.1 5.9 4.4 9.3 3.2c17.6-6.1 36.2-10.4 55.3-12.4c5.6-.6 8.8-6.6 6.3-11.6c-32.5-64.3-98.9-108.7-175.7-109.9c-110.9-1.7-203.3 89.2-203.3 199.9c0 62.8 28.9 118.8 74.2 155.5c-31.8 14.7-61.1 35-86.5 60.4c-54.8 54.7-85.8 126.9-87.8 204a8 8 0 0 0 8 8.2h56.1c4.3 0 7.9-3.4 8-7.7c1.9-58 25.4-112.3 66.7-153.5c29.4-29.4 65.4-49.8 104.7-59.7c3.9-1 6.5-4.7 6-8.7M824 484c0-109.4-87.9-198.3-196.9-200C516.3 282.3 424 373.2 424 484c0 62.8 29 118.8 74.2 155.5a300.95 300.95 0 0 0-86.4 60.4C357 754.6 326 826.8 324 903.8a8 8 0 0 0 8 8.2h56c4.3 0 7.9-3.4 8-7.7c1.9-58 25.4-112.3 66.7-153.5C505.8 707.7 563 684 624 684c110.4 0 200-89.5 200-200m-109.5 90.5C690.3 598.7 658.2 612 624 612s-66.3-13.3-90.5-37.5a127.26 127.26 0 0 1-37.5-91.8c.3-32.8 13.4-64.5 36.3-88c24-24.6 56.1-38.3 90.4-38.7c33.9-.3 66.8 12.9 91 36.6c24.8 24.3 38.4 56.8 38.4 91.4c-.1 34.2-13.4 66.3-37.6 90.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerifiedOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m447.8 588.8l-7.3-32.5c-.2-1-.6-1.9-1.1-2.7c-2.5-3.7-7.4-4.7-11.1-2.2L405 567V411c0-4.4-3.6-8-8-8h-81c-4.4 0-8 3.6-8 8v36c0 4.4 3.6 8 8 8h37v192.4c0 1.7.5 3.3 1.5 4.7c2.6 3.6 7.6 4.4 11.2 1.8l79-56.8c2.6-1.9 3.8-5.1 3.1-8.3m-56.7-216.6l.2.2c3.2 3 8.3 2.8 11.3-.5l24.1-26.2c2.9-3.2 2.8-8.1-.3-11.2l-53.7-52.1c-3.1-3-8.1-3-11.2.1l-24.7 24.7c-3.1 3.1-3.1 8.2.1 11.3z"/><path fill="currentColor" d="M866.9 169.9L527.1 54.1C523 52.7 517.5 52 512 52s-11 .7-15.1 2.1L157.1 169.9c-8.3 2.8-15.1 12.4-15.1 21.2v482.4c0 8.8 5.7 20.4 12.6 25.9L499.3 968c3.5 2.7 8 4.1 12.6 4.1s9.2-1.4 12.6-4.1l344.7-268.6c6.9-5.4 12.6-17 12.6-25.9V191.1c.2-8.8-6.6-18.3-14.9-21.2M810 654.3L512 886.5L214 654.3V226.7l298-101.6l298 101.6z"/><path fill="currentColor" d="M452 297v36c0 4.4 3.6 8 8 8h108v274h-38V405c0-4.4-3.6-8-8-8h-35c-4.4 0-8 3.6-8 8v210h-31c-4.4 0-8 3.6-8 8v37c0 4.4 3.6 8 8 8h244c4.4 0 8-3.6 8-8v-37c0-4.4-3.6-8-8-8h-72V493h58c4.4 0 8-3.6 8-8v-35c0-4.4-3.6-8-8-8h-58V341h63c4.4 0 8-3.6 8-8v-36c0-4.4-3.6-8-8-8H460c-4.4 0-8 3.6-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerticalAlignBottomOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M859.9 780H164.1c-4.5 0-8.1 3.6-8.1 8v60c0 4.4 3.6 8 8.1 8h695.8c4.5 0 8.1-3.6 8.1-8v-60c0-4.4-3.6-8-8.1-8M505.7 669a8 8 0 0 0 12.6 0l112-141.7c4.1-5.2.4-12.9-6.3-12.9h-74.1V176c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8v338.3H400c-6.7 0-10.4 7.7-6.3 12.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerticalAlignMiddleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M859.9 474H164.1c-4.5 0-8.1 3.6-8.1 8v60c0 4.4 3.6 8 8.1 8h695.8c4.5 0 8.1-3.6 8.1-8v-60c0-4.4-3.6-8-8.1-8m-353.6-74.7c2.9 3.7 8.5 3.7 11.3 0l100.8-127.5c3.7-4.7.4-11.7-5.7-11.7H550V104c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8v156h-62.8c-6 0-9.4 7-5.7 11.7zm11.4 225.4a7.14 7.14 0 0 0-11.3 0L405.6 752.3a7.23 7.23 0 0 0 5.7 11.7H474v156c0 4.4 3.6 8 8 8h60c4.4 0 8-3.6 8-8V764h62.8c6 0 9.4-7 5.7-11.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerticalAlignTopOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M859.9 168H164.1c-4.5 0-8.1 3.6-8.1 8v60c0 4.4 3.6 8 8.1 8h695.8c4.5 0 8.1-3.6 8.1-8v-60c0-4.4-3.6-8-8.1-8M518.3 355a8 8 0 0 0-12.6 0l-112 141.7a7.98 7.98 0 0 0 6.3 12.9h73.9V848c0 4.4 3.6 8 8 8h60c4.4 0 8-3.6 8-8V509.7H624c6.7 0 10.4-7.7 6.3-12.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerticalLeftOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M762 164h-64c-4.4 0-8 3.6-8 8v688c0 4.4 3.6 8 8 8h64c4.4 0 8-3.6 8-8V172c0-4.4-3.6-8-8-8m-508 0v72.4c0 9.5 4.2 18.4 11.4 24.5L564.6 512L265.4 763.1c-7.2 6.1-11.4 15-11.4 24.5V860c0 6.8 7.9 10.5 13.1 6.1L689 512L267.1 157.9A7.95 7.95 0 0 0 254 164"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerticalRightOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M326 164h-64c-4.4 0-8 3.6-8 8v688c0 4.4 3.6 8 8 8h64c4.4 0 8-3.6 8-8V172c0-4.4-3.6-8-8-8m444 72.4V164c0-6.8-7.9-10.5-13.1-6.1L335 512l421.9 354.1c5.2 4.4 13.1.7 13.1-6.1v-72.4c0-9.4-4.2-18.4-11.4-24.5L459.4 512l299.2-251.1c7.2-6.1 11.4-15.1 11.4-24.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoCameraAddOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M368 724H252V608c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v116H72c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8h116v116c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V788h116c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8"/><path fill="currentColor" d="M912 302.3L784 376V224c0-35.3-28.7-64-64-64H128c-35.3 0-64 28.7-64 64v352h72V232h576v560H448v72h272c35.3 0 64-28.7 64-64V648l128 73.7c21.3 12.3 48-3.1 48-27.6V330c0-24.6-26.7-40-48-27.7M888 625l-104-59.8V458.9L888 399z"/><path fill="currentColor" d="M320 360c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8H208c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoCameraFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M912 302.3L784 376V224c0-35.3-28.7-64-64-64H128c-35.3 0-64 28.7-64 64v576c0 35.3 28.7 64 64 64h592c35.3 0 64-28.7 64-64V648l128 73.7c21.3 12.3 48-3.1 48-27.6V330c0-24.6-26.7-40-48-27.7M328 352c0 4.4-3.6 8-8 8H208c-4.4 0-8-3.6-8-8v-48c0-4.4 3.6-8 8-8h112c4.4 0 8 3.6 8 8zm560 273l-104-59.8V458.9L888 399z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoCameraOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M912 302.3L784 376V224c0-35.3-28.7-64-64-64H128c-35.3 0-64 28.7-64 64v576c0 35.3 28.7 64 64 64h592c35.3 0 64-28.7 64-64V648l128 73.7c21.3 12.3 48-3.1 48-27.6V330c0-24.6-26.7-40-48-27.7M712 792H136V232h576zm176-167l-104-59.8V458.9L888 399zM208 360h112c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8H208c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoCameraTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".15" d="M136 792h576V232H136zm64-488c0-4.4 3.6-8 8-8h112c4.4 0 8 3.6 8 8v48c0 4.4-3.6 8-8 8H208c-4.4 0-8-3.6-8-8z"/><path fill="currentColor" d="M912 302.3L784 376V224c0-35.3-28.7-64-64-64H128c-35.3 0-64 28.7-64 64v576c0 35.3 28.7 64 64 64h592c35.3 0 64-28.7 64-64V648l128 73.7c21.3 12.3 48-3.1 48-27.6V330c0-24.6-26.7-40-48-27.7M712 792H136V232h576zm176-167l-104-59.8V458.9L888 399z"/><path fill="currentColor" d="M208 360h112c4.4 0 8-3.6 8-8v-48c0-4.4-3.6-8-8-8H208c-4.4 0-8 3.6-8 8v48c0 4.4 3.6 8 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WalletFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-32 464H528V448h320zm-268-64a40 40 0 1 0 80 0a40 40 0 1 0-80 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WalletOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 464H528V448h312zm0 264H184V184h656v200H496c-17.7 0-32 14.3-32 32v192c0 17.7 14.3 32 32 32h344zM580 512a40 40 0 1 0 80 0a40 40 0 1 0-80 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WalletTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32m-40 464H528V448h312zm0-192H496c-17.7 0-32 14.3-32 32v192c0 17.7 14.3 32 32 32h344v200H184V184h656z"/><path fill="currentColor" fill-opacity=".15" d="M528 576h312V448H528zm92-104c22.1 0 40 17.9 40 40s-17.9 40-40 40s-40-17.9-40-40s17.9-40 40-40"/><path fill="currentColor" d="M580 512a40 40 0 1 0 80 0a40 40 0 1 0-80 0"/><path fill="currentColor" fill-opacity=".15" d="M184 840h656V640H496c-17.7 0-32-14.3-32-32V416c0-17.7 14.3-32 32-32h344V184H184z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WarningFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m955.7 856l-416-720c-6.2-10.7-16.9-16-27.7-16s-21.6 5.3-27.7 16l-416 720C56 877.4 71.4 904 96 904h832c24.6 0 40-26.6 27.7-48M480 416c0-4.4 3.6-8 8-8h48c4.4 0 8 3.6 8 8v184c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8zm32 352a48.01 48.01 0 0 1 0-96a48.01 48.01 0 0 1 0 96"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WarningOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 720a48 48 0 1 0 96 0a48 48 0 1 0-96 0m16-304v184c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V416c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8m475.7 440l-416-720c-6.2-10.7-16.9-16-27.7-16s-21.6 5.3-27.7 16l-416 720C56 877.4 71.4 904 96 904h832c24.6 0 40-26.6 27.7-48m-783.5-27.9L512 239.9l339.8 588.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WarningTwotone(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m955.7 856l-416-720c-6.2-10.7-16.9-16-27.7-16s-21.6 5.3-27.7 16l-416 720C56 877.4 71.4 904 96 904h832c24.6 0 40-26.6 27.7-48m-783.5-27.9L512 239.9l339.8 588.2z"/><path fill="currentColor" fill-opacity=".15" d="M172.2 828.1h679.6L512 239.9zM560 720a48.01 48.01 0 0 1-96 0a48.01 48.01 0 0 1 96 0m-16-304v184c0 4.4-3.6 8-8 8h-48c-4.4 0-8-3.6-8-8V416c0-4.4 3.6-8 8-8h48c4.4 0 8 3.6 8 8"/><path fill="currentColor" d="M464 720a48 48 0 1 0 96 0a48 48 0 1 0-96 0m16-304v184c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V416c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WechatFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M690.1 377.4c5.9 0 11.8.2 17.6.5c-24.4-128.7-158.3-227.1-319.9-227.1C209 150.8 64 271.4 64 420.2c0 81.1 43.6 154.2 111.9 203.6a21.5 21.5 0 0 1 9.1 17.6c0 2.4-.5 4.6-1.1 6.9c-5.5 20.3-14.2 52.8-14.6 54.3c-.7 2.6-1.7 5.2-1.7 7.9c0 5.9 4.8 10.8 10.8 10.8c2.3 0 4.2-.9 6.2-2l70.9-40.9c5.3-3.1 11-5 17.2-5c3.2 0 6.4.5 9.5 1.4c33.1 9.5 68.8 14.8 105.7 14.8c6 0 11.9-.1 17.8-.4c-7.1-21-10.9-43.1-10.9-66c0-135.8 132.2-245.8 295.3-245.8m-194.3-86.5c23.8 0 43.2 19.3 43.2 43.1s-19.3 43.1-43.2 43.1c-23.8 0-43.2-19.3-43.2-43.1s19.4-43.1 43.2-43.1m-215.9 86.2c-23.8 0-43.2-19.3-43.2-43.1s19.3-43.1 43.2-43.1s43.2 19.3 43.2 43.1s-19.4 43.1-43.2 43.1m586.8 415.6c56.9-41.2 93.2-102 93.2-169.7c0-124-120.8-224.5-269.9-224.5c-149 0-269.9 100.5-269.9 224.5S540.9 847.5 690 847.5c30.8 0 60.6-4.4 88.1-12.3c2.6-.8 5.2-1.2 7.9-1.2c5.2 0 9.9 1.6 14.3 4.1l59.1 34c1.7 1 3.3 1.7 5.2 1.7a9 9 0 0 0 6.4-2.6a9 9 0 0 0 2.6-6.4c0-2.2-.9-4.4-1.4-6.6c-.3-1.2-7.6-28.3-12.2-45.3c-.5-1.9-.9-3.8-.9-5.7c.1-5.9 3.1-11.2 7.6-14.5M600.2 587.2c-19.9 0-36-16.1-36-35.9c0-19.8 16.1-35.9 36-35.9s36 16.1 36 35.9c0 19.8-16.2 35.9-36 35.9m179.9 0c-19.9 0-36-16.1-36-35.9c0-19.8 16.1-35.9 36-35.9s36 16.1 36 35.9a36.08 36.08 0 0 1-36 35.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WechatOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M690.1 377.4c5.9 0 11.8.2 17.6.5c-24.4-128.7-158.3-227.1-319.9-227.1C209 150.8 64 271.4 64 420.2c0 81.1 43.6 154.2 111.9 203.6a21.5 21.5 0 0 1 9.1 17.6c0 2.4-.5 4.6-1.1 6.9c-5.5 20.3-14.2 52.8-14.6 54.3c-.7 2.6-1.7 5.2-1.7 7.9c0 5.9 4.8 10.8 10.8 10.8c2.3 0 4.2-.9 6.2-2l70.9-40.9c5.3-3.1 11-5 17.2-5c3.2 0 6.4.5 9.5 1.4c33.1 9.5 68.8 14.8 105.7 14.8c6 0 11.9-.1 17.8-.4c-7.1-21-10.9-43.1-10.9-66c0-135.8 132.2-245.8 295.3-245.8m-194.3-86.5c23.8 0 43.2 19.3 43.2 43.1s-19.3 43.1-43.2 43.1c-23.8 0-43.2-19.3-43.2-43.1s19.4-43.1 43.2-43.1m-215.9 86.2c-23.8 0-43.2-19.3-43.2-43.1s19.3-43.1 43.2-43.1s43.2 19.3 43.2 43.1s-19.4 43.1-43.2 43.1m586.8 415.6c56.9-41.2 93.2-102 93.2-169.7c0-124-120.8-224.5-269.9-224.5c-149 0-269.9 100.5-269.9 224.5S540.9 847.5 690 847.5c30.8 0 60.6-4.4 88.1-12.3c2.6-.8 5.2-1.2 7.9-1.2c5.2 0 9.9 1.6 14.3 4.1l59.1 34c1.7 1 3.3 1.7 5.2 1.7a9 9 0 0 0 6.4-2.6a9 9 0 0 0 2.6-6.4c0-2.2-.9-4.4-1.4-6.6c-.3-1.2-7.6-28.3-12.2-45.3c-.5-1.9-.9-3.8-.9-5.7c.1-5.9 3.1-11.2 7.6-14.5M600.2 587.2c-19.9 0-36-16.1-36-35.9c0-19.8 16.1-35.9 36-35.9s36 16.1 36 35.9c0 19.8-16.2 35.9-36 35.9m179.9 0c-19.9 0-36-16.1-36-35.9c0-19.8 16.1-35.9 36-35.9s36 16.1 36 35.9a36.08 36.08 0 0 1-36 35.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WeiboCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m-44.4 672C353.1 736 236 680.4 236 588.9c0-47.8 30.2-103.1 82.3-155.3c69.5-69.6 150.6-101.4 181.1-70.8c13.5 13.5 14.8 36.8 6.1 64.6c-4.5 14 13.1 6.3 13.1 6.3c56.2-23.6 105.2-25 123.1.7c9.6 13.7 8.6 32.8-.2 55.1c-4.1 10.2 1.3 11.8 9 14.1c31.7 9.8 66.9 33.6 66.9 75.5c.2 69.5-99.7 156.9-249.8 156.9m207.3-290.8a34.9 34.9 0 0 0-7.2-34.1a34.68 34.68 0 0 0-33.1-10.7a18.24 18.24 0 0 1-7.6-35.7c24.1-5.1 50.1 2.3 67.7 21.9c17.7 19.6 22.4 46.3 14.9 69.8a18.13 18.13 0 0 1-22.9 11.7a18.18 18.18 0 0 1-11.8-22.9m106 34.3s0 .1 0 0a21.1 21.1 0 0 1-26.6 13.7a21.19 21.19 0 0 1-13.6-26.7c11-34.2 4-73.2-21.7-101.8a104.04 104.04 0 0 0-98.9-32.1a21.14 21.14 0 0 1-25.1-16.3a21.07 21.07 0 0 1 16.2-25.1c49.4-10.5 102.8 4.8 139.1 45.1c36.3 40.2 46.1 95.1 30.6 143.2m-334.5 6.1c-91.4 9-160.7 65.1-154.7 125.2c5.9 60.1 84.8 101.5 176.2 92.5c91.4-9.1 160.7-65.1 154.7-125.3c-5.9-60.1-84.8-101.5-176.2-92.4m80.2 141.7c-18.7 42.3-72.3 64.8-117.8 50.1c-43.9-14.2-62.5-57.7-43.3-96.8c18.9-38.4 68-60.1 111.5-48.8c45 11.7 68 54.2 49.6 95.5m-93-32.2c-14.2-5.9-32.4.2-41.2 13.9c-8.8 13.8-4.7 30.2 9.3 36.6c14.3 6.5 33.2.3 42-13.8c8.8-14.3 4.2-30.6-10.1-36.7m34.9-14.5c-5.4-2.2-12.2.5-15.4 5.8c-3.1 5.4-1.4 11.5 4.1 13.8c5.5 2.3 12.6-.3 15.8-5.8c3-5.6 1-11.8-4.5-13.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WeiboCircleOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m-44.4 672C353.1 736 236 680.4 236 588.9c0-47.8 30.2-103.1 82.3-155.3c69.5-69.6 150.6-101.4 181.1-70.8c13.5 13.5 14.8 36.8 6.1 64.6c-4.5 14 13.1 6.3 13.1 6.3c56.2-23.6 105.2-25 123.1.7c9.6 13.7 8.6 32.8-.2 55.1c-4.1 10.2 1.3 11.8 9 14.1c31.7 9.8 66.9 33.6 66.9 75.5c.2 69.5-99.7 156.9-249.8 156.9m207.3-290.8a34.9 34.9 0 0 0-7.2-34.1a34.68 34.68 0 0 0-33.1-10.7a18.24 18.24 0 0 1-7.6-35.7c24.1-5.1 50.1 2.3 67.7 21.9c17.7 19.6 22.4 46.3 14.9 69.8a18.13 18.13 0 0 1-22.9 11.7a18.18 18.18 0 0 1-11.8-22.9m106 34.3s0 .1 0 0a21.1 21.1 0 0 1-26.6 13.7a21.19 21.19 0 0 1-13.6-26.7c11-34.2 4-73.2-21.7-101.8a104.04 104.04 0 0 0-98.9-32.1a21.14 21.14 0 0 1-25.1-16.3a21.07 21.07 0 0 1 16.2-25.1c49.4-10.5 102.8 4.8 139.1 45.1c36.3 40.2 46.1 95.1 30.6 143.2m-334.5 6.1c-91.4 9-160.7 65.1-154.7 125.2c5.9 60.1 84.8 101.5 176.2 92.5c91.4-9.1 160.7-65.1 154.7-125.3c-5.9-60.1-84.8-101.5-176.2-92.4m80.2 141.7c-18.7 42.3-72.3 64.8-117.8 50.1c-43.9-14.2-62.5-57.7-43.3-96.8c18.9-38.4 68-60.1 111.5-48.8c45 11.7 68 54.2 49.6 95.5m-93-32.2c-14.2-5.9-32.4.2-41.2 13.9c-8.8 13.8-4.7 30.2 9.3 36.6c14.3 6.5 33.2.3 42-13.8c8.8-14.3 4.2-30.6-10.1-36.7m34.9-14.5c-5.4-2.2-12.2.5-15.4 5.8c-3.1 5.4-1.4 11.5 4.1 13.8c5.5 2.3 12.6-.3 15.8-5.8c3-5.6 1-11.8-4.5-13.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WeiboOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M457.3 543c-68.1-17.7-145 16.2-174.6 76.2c-30.1 61.2-1 129.1 67.8 151.3c71.2 23 155.2-12.2 184.4-78.3c28.7-64.6-7.2-131-77.6-149.2m-52 156.2c-13.8 22.1-43.5 31.7-65.8 21.6c-22-10-28.5-35.7-14.6-57.2c13.7-21.4 42.3-31 64.4-21.7c22.4 9.5 29.6 35 16 57.3m45.5-58.5c-5 8.6-16.1 12.7-24.7 9.1c-8.5-3.5-11.2-13.1-6.4-21.5c5-8.4 15.6-12.4 24.1-9.1c8.7 3.2 11.8 12.9 7 21.5m334.5-197.2c15 4.8 31-3.4 35.9-18.3c11.8-36.6 4.4-78.4-23.2-109a111.39 111.39 0 0 0-106-34.3a28.45 28.45 0 0 0-21.9 33.8a28.39 28.39 0 0 0 33.8 21.8c18.4-3.9 38.3 1.8 51.9 16.7a54.2 54.2 0 0 1 11.3 53.3a28.45 28.45 0 0 0 18.2 36m99.8-206c-56.7-62.9-140.4-86.9-217.7-70.5a32.98 32.98 0 0 0-25.4 39.3a33.12 33.12 0 0 0 39.3 25.5c55-11.7 114.4 5.4 154.8 50.1c40.3 44.7 51.2 105.7 34 159.1c-5.6 17.4 3.9 36 21.3 41.7c17.4 5.6 36-3.9 41.6-21.2v-.1c24.1-75.4 8.9-161.1-47.9-223.9M729 499c-12.2-3.6-20.5-6.1-14.1-22.1c13.8-34.7 15.2-64.7.3-86c-28-40.1-104.8-37.9-192.8-1.1c0 0-27.6 12.1-20.6-9.8c13.5-43.5 11.5-79.9-9.6-101c-47.7-47.8-174.6 1.8-283.5 110.6C127.3 471.1 80 557.5 80 632.2C80 775.1 263.2 862 442.5 862c235 0 391.3-136.5 391.3-245c0-65.5-55.2-102.6-104.8-118M443 810.8c-143 14.1-266.5-50.5-275.8-144.5c-9.3-93.9 99.2-181.5 242.2-195.6c143-14.2 266.5 50.5 275.8 144.4C694.4 709 586 796.6 443 810.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WeiboSquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M433.6 595.1c-14.2-5.9-32.4.2-41.2 13.9c-8.8 13.8-4.7 30.2 9.3 36.6c14.3 6.5 33.2.3 42-13.8c8.8-14.3 4.2-30.6-10.1-36.7M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M467.6 736C353.1 736 236 680.4 236 588.9c0-47.8 30.2-103.1 82.3-155.3c69.5-69.6 150.6-101.4 181.1-70.8c13.5 13.5 14.8 36.8 6.1 64.6c-4.5 14 13.1 6.3 13.1 6.3c56.2-23.6 105.2-25 123.1.7c9.6 13.7 8.6 32.8-.2 55.1c-4.1 10.2 1.3 11.8 9 14.1c31.7 9.8 66.9 33.6 66.9 75.5c.2 69.5-99.7 156.9-249.8 156.9m207.3-290.8a34.9 34.9 0 0 0-7.2-34.1a34.68 34.68 0 0 0-33.1-10.7a18.24 18.24 0 0 1-7.6-35.7c24.1-5.1 50.1 2.3 67.7 21.9c17.7 19.6 22.4 46.3 14.9 69.8a18.13 18.13 0 0 1-22.9 11.7a18.18 18.18 0 0 1-11.8-22.9m106 34.3s0 .1 0 0a21.1 21.1 0 0 1-26.6 13.7a21.19 21.19 0 0 1-13.6-26.7c11-34.2 4-73.2-21.7-101.8a104.04 104.04 0 0 0-98.9-32.1a21.14 21.14 0 0 1-25.1-16.3a21.07 21.07 0 0 1 16.2-25.1c49.4-10.5 102.8 4.8 139.1 45.1c36.3 40.2 46.1 95.1 30.6 143.2m-334.5 6.1c-91.4 9-160.7 65.1-154.7 125.2c5.9 60.1 84.8 101.5 176.2 92.5c91.4-9.1 160.7-65.1 154.7-125.3c-5.9-60.1-84.8-101.5-176.2-92.4m80.2 141.7c-18.7 42.3-72.3 64.8-117.8 50.1c-43.9-14.2-62.5-57.7-43.3-96.8c18.9-38.4 68-60.1 111.5-48.8c45 11.7 68 54.2 49.6 95.5m-58.1-46.7c-5.4-2.2-12.2.5-15.4 5.8c-3.1 5.4-1.4 11.5 4.1 13.8c5.5 2.3 12.6-.3 15.8-5.8c3-5.6 1-11.8-4.5-13.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WeiboSquareOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M433.6 595.1c-14.2-5.9-32.4.2-41.2 13.9c-8.8 13.8-4.7 30.2 9.3 36.6c14.3 6.5 33.2.3 42-13.8c8.8-14.3 4.2-30.6-10.1-36.7M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M467.6 736C353.1 736 236 680.4 236 588.9c0-47.8 30.2-103.1 82.3-155.3c69.5-69.6 150.6-101.4 181.1-70.8c13.5 13.5 14.8 36.8 6.1 64.6c-4.5 14 13.1 6.3 13.1 6.3c56.2-23.6 105.2-25 123.1.7c9.6 13.7 8.6 32.8-.2 55.1c-4.1 10.2 1.3 11.8 9 14.1c31.7 9.8 66.9 33.6 66.9 75.5c.2 69.5-99.7 156.9-249.8 156.9m207.3-290.8a34.9 34.9 0 0 0-7.2-34.1a34.68 34.68 0 0 0-33.1-10.7a18.24 18.24 0 0 1-7.6-35.7c24.1-5.1 50.1 2.3 67.7 21.9c17.7 19.6 22.4 46.3 14.9 69.8a18.13 18.13 0 0 1-22.9 11.7a18.18 18.18 0 0 1-11.8-22.9m106 34.3s0 .1 0 0a21.1 21.1 0 0 1-26.6 13.7a21.19 21.19 0 0 1-13.6-26.7c11-34.2 4-73.2-21.7-101.8a104.04 104.04 0 0 0-98.9-32.1a21.14 21.14 0 0 1-25.1-16.3a21.07 21.07 0 0 1 16.2-25.1c49.4-10.5 102.8 4.8 139.1 45.1c36.3 40.2 46.1 95.1 30.6 143.2m-334.5 6.1c-91.4 9-160.7 65.1-154.7 125.2c5.9 60.1 84.8 101.5 176.2 92.5c91.4-9.1 160.7-65.1 154.7-125.3c-5.9-60.1-84.8-101.5-176.2-92.4m80.2 141.7c-18.7 42.3-72.3 64.8-117.8 50.1c-43.9-14.2-62.5-57.7-43.3-96.8c18.9-38.4 68-60.1 111.5-48.8c45 11.7 68 54.2 49.6 95.5m-58.1-46.7c-5.4-2.2-12.2.5-15.4 5.8c-3.1 5.4-1.4 11.5 4.1 13.8c5.5 2.3 12.6-.3 15.8-5.8c3-5.6 1-11.8-4.5-13.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WhatsAppOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M713.5 599.9c-10.9-5.6-65.2-32.2-75.3-35.8c-10.1-3.8-17.5-5.6-24.8 5.6c-7.4 11.1-28.4 35.8-35 43.3c-6.4 7.4-12.9 8.3-23.8 2.8c-64.8-32.4-107.3-57.8-150-131.1c-11.3-19.5 11.3-18.1 32.4-60.2c3.6-7.4 1.8-13.7-1-19.3c-2.8-5.6-24.8-59.8-34-81.9c-8.9-21.5-18.1-18.5-24.8-18.9c-6.4-.4-13.7-.4-21.1-.4c-7.4 0-19.3 2.8-29.4 13.7c-10.1 11.1-38.6 37.8-38.6 92s39.5 106.7 44.9 114.1c5.6 7.4 77.7 118.6 188.4 166.5c70 30.2 97.4 32.8 132.4 27.6c21.3-3.2 65.2-26.6 74.3-52.5c9.1-25.8 9.1-47.9 6.4-52.5c-2.7-4.9-10.1-7.7-21-13"/><path fill="currentColor" d="M925.2 338.4c-22.6-53.7-55-101.9-96.3-143.3c-41.3-41.3-89.5-73.8-143.3-96.3C630.6 75.7 572.2 64 512 64h-2c-60.6.3-119.3 12.3-174.5 35.9c-53.3 22.8-101.1 55.2-142 96.5c-40.9 41.3-73 89.3-95.2 142.8c-23 55.4-34.6 114.3-34.3 174.9c.3 69.4 16.9 138.3 48 199.9v152c0 25.4 20.6 46 46 46h152.1c61.6 31.1 130.5 47.7 199.9 48h2.1c59.9 0 118-11.6 172.7-34.3c53.5-22.3 101.6-54.3 142.8-95.2c41.3-40.9 73.8-88.7 96.5-142c23.6-55.2 35.6-113.9 35.9-174.5c.3-60.9-11.5-120-34.8-175.6m-151.1 438C704 845.8 611 884 512 884h-1.7c-60.3-.3-120.2-15.3-173.1-43.5l-8.4-4.5H188V695.2l-4.5-8.4C155.3 633.9 140.3 574 140 513.7c-.4-99.7 37.7-193.3 107.6-263.8c69.8-70.5 163.1-109.5 262.8-109.9h1.7c50 0 98.5 9.7 144.2 28.9c44.6 18.7 84.6 45.6 119 80c34.3 34.3 61.3 74.4 80 119c19.4 46.2 29.1 95.2 28.9 145.8c-.6 99.6-39.7 192.9-110.1 262.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M723 620.5C666.8 571.6 593.4 542 513 542s-153.8 29.6-210.1 78.6a8.1 8.1 0 0 0-.8 11.2l36 42.9c2.9 3.4 8 3.8 11.4.9C393.1 637.2 450.3 614 513 614s119.9 23.2 163.5 61.5c3.4 2.9 8.5 2.5 11.4-.9l36-42.9c2.8-3.3 2.4-8.3-.9-11.2m117.4-140.1C751.7 406.5 637.6 362 513 362s-238.7 44.5-327.5 118.4a8.05 8.05 0 0 0-1 11.3l36 42.9c2.8 3.4 7.9 3.8 11.2 1C308 472.2 406.1 434 513 434s205 38.2 281.2 101.6c3.4 2.8 8.4 2.4 11.2-1l36-42.9c2.8-3.4 2.4-8.5-1-11.3m116.7-139C835.7 241.8 680.3 182 511 182c-168.2 0-322.6 59-443.7 157.4a8 8 0 0 0-1.1 11.4l36 42.9c2.8 3.3 7.8 3.8 11.1 1.1C222 306.7 360.3 254 511 254c151.8 0 291 53.5 400 142.7c3.4 2.8 8.4 2.3 11.2-1.1l36-42.9c2.9-3.4 2.4-8.5-1.1-11.3M448 778a64 64 0 1 0 128 0a64 64 0 1 0-128 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowsFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M523.8 191.4v288.9h382V128.1zm0 642.2l382 62.2v-352h-382zM120.1 480.2H443V201.9l-322.9 53.5zm0 290.4L443 823.2V543.8H120.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowsOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M120.1 770.6L443 823.2V543.8H120.1zm63.4-163.5h196.2v141.6l-196.2-31.9zm340.3 226.5l382 62.2v-352h-382zm63.4-226.5h255.3v214.4l-255.3-41.6zm-63.4-415.7v288.8h382V128.1zm318.7 225.5H587.3V245l255.3-42.3v214.2zm-722.4 63.3H443V201.9l-322.9 53.5zM183.5 309l196.2-32.5v140.4H183.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WomanOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M712.8 548.8c53.6-53.6 83.2-125 83.2-200.8c0-75.9-29.5-147.2-83.2-200.8C659.2 93.6 587.8 64 512 64s-147.2 29.5-200.8 83.2C257.6 200.9 228 272.1 228 348c0 63.8 20.9 124.4 59.4 173.9c7.3 9.4 15.2 18.3 23.7 26.9c8.5 8.5 17.5 16.4 26.8 23.7c39.6 30.8 86.3 50.4 136.1 57V736H360c-4.4 0-8 3.6-8 8v60c0 4.4 3.6 8 8 8h114v140c0 4.4 3.6 8 8 8h60c4.4 0 8-3.6 8-8V812h114c4.4 0 8-3.6 8-8v-60c0-4.4-3.6-8-8-8H550V629.5c61.5-8.2 118.2-36.1 162.8-80.7M512 556c-55.6 0-107.7-21.6-147.1-60.9C325.6 455.8 304 403.6 304 348s21.6-107.7 60.9-147.1C404.2 161.5 456.4 140 512 140s107.7 21.6 147.1 60.9C698.4 240.2 720 292.4 720 348s-21.6 107.7-60.9 147.1C619.7 534.4 567.6 556 512 556"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YahooFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M937.3 231H824.7c-15.5 0-27.7 12.6-27.1 28.1l13.1 366h84.4l65.4-366.4c2.7-15.2-7.8-27.7-23.2-27.7m-77.4 450.4h-14.1c-27.1 0-49.2 22.2-49.2 49.3v14.1c0 27.1 22.2 49.3 49.2 49.3h14.1c27.1 0 49.2-22.2 49.2-49.3v-14.1c0-27.1-22.2-49.3-49.2-49.3M402.6 231C216.2 231 65 357 65 512.5S216.2 794 402.6 794s337.6-126 337.6-281.5S589.1 231 402.6 231m225.2 225.2h-65.3L458.9 559.8v65.3h84.4v56.3H318.2v-56.3h84.4v-65.3L242.9 399.9h-37v-56.3h168.5v56.3h-37l93.4 93.5l28.1-28.1V400h168.8v56.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YahooOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M859.9 681.4h-14.1c-27.1 0-49.2 22.2-49.2 49.3v14.1c0 27.1 22.2 49.3 49.2 49.3h14.1c27.1 0 49.2-22.2 49.2-49.3v-14.1c0-27.1-22.2-49.3-49.2-49.3M402.6 231C216.2 231 65 357 65 512.5S216.2 794 402.6 794s337.6-126 337.6-281.5S589.1 231 402.6 231m0 507C245.1 738 121 634.6 121 512.5c0-62.3 32.3-119.7 84.9-161v48.4h37l159.8 159.9v65.3h-84.4v56.3h225.1v-56.3H459v-65.3l103.5-103.6h65.3v-56.3H459v65.3l-28.1 28.1l-93.4-93.5h37v-56.3H216.4c49.4-35 114.3-56.6 186.2-56.6c157.6 0 281.6 103.4 281.6 225.5S560.2 738 402.6 738m534.7-507H824.7c-15.5 0-27.7 12.6-27.1 28.1l13.1 366h84.4l65.4-366.4c2.7-15.2-7.8-27.7-23.2-27.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YoutubeFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M941.3 296.1a112.3 112.3 0 0 0-79.2-79.3C792.2 198 512 198 512 198s-280.2 0-350.1 18.7A112.12 112.12 0 0 0 82.7 296C64 366 64 512 64 512s0 146 18.7 215.9c10.3 38.6 40.7 69 79.2 79.3C231.8 826 512 826 512 826s280.2 0 350.1-18.8c38.6-10.3 68.9-40.7 79.2-79.3C960 658 960 512 960 512s0-146-18.7-215.9M423 646V378l232 133z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YoutubeOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M960 509.2c0-2.2 0-4.7-.1-7.6c-.1-8.1-.3-17.2-.5-26.9c-.8-27.9-2.2-55.7-4.4-81.9c-3-36.1-7.4-66.2-13.4-88.8a139.52 139.52 0 0 0-98.3-98.5c-28.3-7.6-83.7-12.3-161.7-15.2c-37.1-1.4-76.8-2.3-116.5-2.8c-13.9-.2-26.8-.3-38.4-.4h-29.4c-11.6.1-24.5.2-38.4.4c-39.7.5-79.4 1.4-116.5 2.8c-78 3-133.5 7.7-161.7 15.2A139.35 139.35 0 0 0 82.4 304C76.3 326.6 72 356.7 69 392.8c-2.2 26.2-3.6 54-4.4 81.9c-.3 9.7-.4 18.8-.5 26.9c0 2.9-.1 5.4-.1 7.6v5.6c0 2.2 0 4.7.1 7.6c.1 8.1.3 17.2.5 26.9c.8 27.9 2.2 55.7 4.4 81.9c3 36.1 7.4 66.2 13.4 88.8c12.8 47.9 50.4 85.7 98.3 98.5c28.2 7.6 83.7 12.3 161.7 15.2c37.1 1.4 76.8 2.3 116.5 2.8c13.9.2 26.8.3 38.4.4h29.4c11.6-.1 24.5-.2 38.4-.4c39.7-.5 79.4-1.4 116.5-2.8c78-3 133.5-7.7 161.7-15.2c47.9-12.8 85.5-50.5 98.3-98.5c6.1-22.6 10.4-52.7 13.4-88.8c2.2-26.2 3.6-54 4.4-81.9c.3-9.7.4-18.8.5-26.9c0-2.9.1-5.4.1-7.6zm-72 5.2c0 2.1 0 4.4-.1 7.1c-.1 7.8-.3 16.4-.5 25.7c-.7 26.6-2.1 53.2-4.2 77.9c-2.7 32.2-6.5 58.6-11.2 76.3c-6.2 23.1-24.4 41.4-47.4 47.5c-21 5.6-73.9 10.1-145.8 12.8c-36.4 1.4-75.6 2.3-114.7 2.8c-13.7.2-26.4.3-37.8.3h-28.6l-37.8-.3c-39.1-.5-78.2-1.4-114.7-2.8c-71.9-2.8-124.9-7.2-145.8-12.8c-23-6.2-41.2-24.4-47.4-47.5c-4.7-17.7-8.5-44.1-11.2-76.3c-2.1-24.7-3.4-51.3-4.2-77.9c-.3-9.3-.4-18-.5-25.7c0-2.7-.1-5.1-.1-7.1v-4.8c0-2.1 0-4.4.1-7.1c.1-7.8.3-16.4.5-25.7c.7-26.6 2.1-53.2 4.2-77.9c2.7-32.2 6.5-58.6 11.2-76.3c6.2-23.1 24.4-41.4 47.4-47.5c21-5.6 73.9-10.1 145.8-12.8c36.4-1.4 75.6-2.3 114.7-2.8c13.7-.2 26.4-.3 37.8-.3h28.6l37.8.3c39.1.5 78.2 1.4 114.7 2.8c71.9 2.8 124.9 7.2 145.8 12.8c23 6.2 41.2 24.4 47.4 47.5c4.7 17.7 8.5 44.1 11.2 76.3c2.1 24.7 3.4 51.3 4.2 77.9c.3 9.3.4 18 .5 25.7c0 2.7.1 5.1.1 7.1zM423 646l232-135l-232-133z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YuqueFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M854.6 370.6c-9.9-39.4 9.9-102.2 73.4-124.4l-67.9-3.6s-25.7-90-143.6-98c-117.9-8.1-195-3-195-3s87.4 55.6 52.4 154.7c-25.6 52.5-65.8 95.6-108.8 144.7c-1.3 1.3-2.5 2.6-3.5 3.7C319.4 605 96 860 96 860c245.9 64.4 410.7-6.3 508.2-91.1c20.5-.2 35.9-.3 46.3-.3c135.8 0 250.6-117.6 245.9-248.4c-3.2-89.9-31.9-110.2-41.8-149.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YuqueOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M854.6 370.6c-9.9-39.4 9.9-102.2 73.4-124.4l-67.9-3.6s-25.7-90-143.6-98c-117.8-8.1-194.9-3-195-3c.1 0 87.4 55.6 52.4 154.7c-25.6 52.5-65.8 95.6-108.8 144.7c-1.3 1.3-2.5 2.6-3.5 3.7C319.4 605 96 860 96 860c245.9 64.4 410.7-6.3 508.2-91.1c20.5-.2 35.9-.3 46.3-.3c135.8 0 250.6-117.6 245.9-248.4c-3.2-89.9-31.9-110.2-41.8-149.6m-204.1 334c-10.6 0-26.2.1-46.8.3l-23.6.2l-17.8 15.5c-47.1 41-104.4 71.5-171.4 87.6c-52.5 12.6-110 16.2-172.7 9.6c18-20.5 36.5-41.6 55.4-63.1c92-104.6 173.8-197.5 236.9-268.5l1.4-1.4l1.3-1.5c4.1-4.6 20.6-23.3 24.7-28.1c9.7-11.1 17.3-19.9 24.5-28.6c30.7-36.7 52.2-67.8 69-102.2l1.6-3.3l1.2-3.4c13.7-38.8 15.4-76.9 6.2-112.8c22.5.7 46.5 1.9 71.7 3.6c33.3 2.3 55.5 12.9 71.1 29.2c5.8 6 10.2 12.5 13.4 18.7c1 2 1.7 3.6 2.3 5l5 17.7c-15.7 34.5-19.9 73.3-11.4 107.2c3 11.8 6.9 22.4 12.3 34.4c2.1 4.7 9.5 20.1 11 23.3c10.3 22.7 15.4 43 16.7 78.7c3.3 94.6-82.7 181.9-182 181.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZhihuCircleFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64m-90.7 477.8l-.1 1.5c-1.5 20.4-6.3 43.9-12.9 67.6l24-18.1l71 80.7c9.2 33-3.3 63.1-3.3 63.1l-95.7-111.9v-.1c-8.9 29-20.1 57.3-33.3 84.7c-22.6 45.7-55.2 54.7-89.5 57.7c-34.4 3-23.3-5.3-23.3-5.3c68-55.5 78-87.8 96.8-123.1c11.9-22.3 20.4-64.3 25.3-96.8H264.1s4.8-31.2 19.2-41.7h101.6c.6-15.3-1.3-102.8-2-131.4h-49.4c-9.2 45-41 56.7-48.1 60.1c-7 3.4-23.6 7.1-21.1 0c2.6-7.1 27-46.2 43.2-110.7c16.3-64.6 63.9-62 63.9-62c-12.8 22.5-22.4 73.6-22.4 73.6h159.7c10.1 0 10.6 39 10.6 39h-90.8c-.7 22.7-2.8 83.8-5 131.4H519s12.2 15.4 12.2 41.7zm346.5 167h-87.6l-69.5 46.6l-16.4-46.6h-40.1V321.5h213.6zM408.2 611s0-.1 0 0m216 94.3l56.8-38.1h45.6h-.1V364.7H596.7v302.5h14.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZhihuOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M564.7 230.1V803h60l25.2 71.4L756.3 803h131.5V230.1zm247.7 497h-59.9l-75.1 50.4l-17.8-50.4h-18V308.3h170.7v418.8zM526.1 486.9H393.3c2.1-44.9 4.3-104.3 6.6-172.9h130.9l-.1-8.1c0-.6-.2-14.7-2.3-29.1c-2.1-15-6.6-34.9-21-34.9H287.8c4.4-20.6 15.7-69.7 29.4-93.8l6.4-11.2l-12.9-.7c-.8 0-19.6-.9-41.4 10.6c-35.7 19-51.7 56.4-58.7 84.4c-18.4 73.1-44.6 123.9-55.7 145.6c-3.3 6.4-5.3 10.2-6.2 12.8c-1.8 4.9-.8 9.8 2.8 13c10.5 9.5 38.2-2.9 38.5-3c.6-.3 1.3-.6 2.2-1c13.9-6.3 55.1-25 69.8-84.5h56.7c.7 32.2 3.1 138.4 2.9 172.9h-141l-2.1 1.5c-23.1 16.9-30.5 63.2-30.8 65.2l-1.4 9.2h167c-12.3 78.3-26.5 113.4-34 127.4c-3.7 7-7.3 14-10.7 20.8c-21.3 42.2-43.4 85.8-126.3 153.6c-3.6 2.8-7 8-4.8 13.7c2.4 6.3 9.3 9.1 24.6 9.1c5.4 0 11.8-.3 19.4-1c49.9-4.4 100.8-18 135.1-87.6c17-35.1 31.7-71.7 43.9-108.9L497 850l5-12c.8-1.9 19-46.3 5.1-95.9l-.5-1.8l-108.1-123l-22 16.6c6.4-26.1 10.6-49.9 12.5-71.1h158.7v-8c0-40.1-18.5-63.9-19.2-64.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZhihuSquareFilled(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M880 112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V144c0-17.7-14.3-32-32-32M432.3 592.8l71 80.7c9.2 33-3.3 63.1-3.3 63.1l-95.7-111.9v-.1c-8.9 29-20.1 57.3-33.3 84.7c-22.6 45.7-55.2 54.7-89.5 57.7c-34.4 3-23.3-5.3-23.3-5.3c68-55.5 78-87.8 96.8-123.1c11.9-22.3 20.4-64.3 25.3-96.8H264.1s4.8-31.2 19.2-41.7h101.6c.6-15.3-1.3-102.8-2-131.4h-49.4c-9.2 45-41 56.7-48.1 60.1c-7 3.4-23.6 7.1-21.1 0c2.6-7.1 27-46.2 43.2-110.7c16.3-64.6 63.9-62 63.9-62c-12.8 22.5-22.4 73.6-22.4 73.6h159.7c10.1 0 10.6 39 10.6 39h-90.8c-.7 22.7-2.8 83.8-5 131.4H519s12.2 15.4 12.2 41.7h-110l-.1 1.5c-1.5 20.4-6.3 43.9-12.9 67.6zm335.5 116h-87.6l-69.5 46.6l-16.4-46.6h-40.1V321.5h213.6zM408.2 611s0-.1 0 0m216 94.3l56.8-38.1h45.6h-.1V364.7H596.7v302.5h14.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomInOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M637 443H519V309c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8v134H325c-4.4 0-8 3.6-8 8v60c0 4.4 3.6 8 8 8h118v134c0 4.4 3.6 8 8 8h60c4.4 0 8-3.6 8-8V519h118c4.4 0 8-3.6 8-8v-60c0-4.4-3.6-8-8-8m284 424L775 721c122.1-148.9 113.6-369.5-26-509c-148-148.1-388.4-148.1-537 0c-148.1 148.6-148.1 389 0 537c139.5 139.6 360.1 148.1 509 26l146 146c3.2 2.8 8.3 2.8 11 0l43-43c2.8-2.7 2.8-7.8 0-11M696 696c-118.8 118.7-311.2 118.7-430 0c-118.7-118.8-118.7-311.2 0-430c118.8-118.7 311.2-118.7 430 0c118.7 118.8 118.7 311.2 0 430"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomOutOutlined(children ...ElementRenderer) *AntDesignIcon {
	return &AntDesignIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M637 443H325c-4.4 0-8 3.6-8 8v60c0 4.4 3.6 8 8 8h312c4.4 0 8-3.6 8-8v-60c0-4.4-3.6-8-8-8m284 424L775 721c122.1-148.9 113.6-369.5-26-509c-148-148.1-388.4-148.1-537 0c-148.1 148.6-148.1 389 0 537c139.5 139.6 360.1 148.1 509 26l146 146c3.2 2.8 8.3 2.8 11 0l43-43c2.8-2.7 2.8-7.8 0-11M696 696c-118.8 118.7-311.2 118.7-430 0c-118.7-118.8-118.7-311.2 0-430c118.8-118.7 311.2-118.7 430 0c118.7 118.8 118.7 311.2 0 430"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
