package oui

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "2.0.0"
	hAttr          = "1em"
	viewbox        = "0 0 0 0"
	fill           = "currentColor"
)

type OuiIcon struct {
	*SVGSVGElement
}

type OuiIconFn func(children ...ElementRenderer) *OuiIcon

var IconLookup = map[string]OuiIconFn{
	"accessibility":              Accessibility,
	"aggregate":                  Aggregate,
	"alert":                      Alert,
	"analyzeEvent":               AnalyzeEvent,
	"annotation":                 Annotation,
	"anomalyDetection":           AnomalyDetection,
	"apmTrace":                   ApmTrace,
	"appAddData":                 AppAddData,
	"appAdvancedSettings":        AppAdvancedSettings,
	"appApm":                     AppApm,
	"appAppSearch":               AppAppSearch,
	"appAuditbeat":               AppAuditbeat,
	"appCanvas":                  AppCanvas,
	"appCode":                    AppCode,
	"appConsole":                 AppConsole,
	"appCrossClusterReplication": AppCrossClusterReplication,
	"appDashboard":               AppDashboard,
	"appDevtools":                AppDevtools,
	"appDiscover":                AppDiscover,
	"appEms":                     AppEms,
	"appFilebeat":                AppFilebeat,
	"appGis":                     AppGis,
	"appGraph":                   AppGraph,
	"appGrok":                    AppGrok,
	"appHeartbeat":               AppHeartbeat,
	"appIndexManagement":         AppIndexManagement,
	"appIndexPattern":            AppIndexPattern,
	"appIndexRollup":             AppIndexRollup,
	"appLens":                    AppLens,
	"appLogs":                    AppLogs,
	"appManagement":              AppManagement,
	"appMetricbeat":              AppMetricbeat,
	"appMetrics":                 AppMetrics,
	"appMl":                      AppMl,
	"appMonitoring":              AppMonitoring,
	"appNotebook":                AppNotebook,
	"appPacketbeat":              AppPacketbeat,
	"appPipeline":                AppPipeline,
	"appRecentlyViewed":          AppRecentlyViewed,
	"appReporting":               AppReporting,
	"appSavedObjects":            AppSavedObjects,
	"appSearchProfiler":          AppSearchProfiler,
	"appSecurity":                AppSecurity,
	"appSecurityAnalytics":       AppSecurityAnalytics,
	"appSpaces":                  AppSpaces,
	"appSql":                     AppSql,
	"appTimelion":                AppTimelion,
	"appUpgradeAssistant":        AppUpgradeAssistant,
	"appUptime":                  AppUptime,
	"appUsersRoles":              AppUsersRoles,
	"appVisualize":               AppVisualize,
	"appWatches":                 AppWatches,
	"appWorkplaceSearch":         AppWorkplaceSearch,
	"apps":                       Apps,
	"arrowDown":                  ArrowDown,
	"arrowLeft":                  ArrowLeft,
	"arrowRight":                 ArrowRight,
	"arrowUp":                    ArrowUp,
	"asterisk":                   Asterisk,
	"beaker":                     Beaker,
	"bell":                       Bell,
	"bellSlash":                  BellSlash,
	"bolt":                       Bolt,
	"boxesHorizontal":            BoxesHorizontal,
	"boxesVertical":              BoxesVertical,
	"branch":                     Branch,
	"broom":                      Broom,
	"brush":                      Brush,
	"bug":                        Bug,
	"bullseye":                   Bullseye,
	"calendar":                   Calendar,
	"chatLeft":                   ChatLeft,
	"chatRight":                  ChatRight,
	"check":                      Check,
	"checkInCircleEmpty":         CheckInCircleEmpty,
	"checkInCircleFilled":        CheckInCircleFilled,
	"cheer":                      Cheer,
	"clock":                      Clock,
	"cloudDrizzle":               CloudDrizzle,
	"cloudStormy":                CloudStormy,
	"cloudSunny":                 CloudSunny,
	"color":                      Color,
	"compute":                    Compute,
	"console":                    Console,
	"continuityAbove":            ContinuityAbove,
	"continuityAboveBelow":       ContinuityAboveBelow,
	"continuityBelow":            ContinuityBelow,
	"continuityWithin":           ContinuityWithin,
	"controlsHorizontal":         ControlsHorizontal,
	"controlsVertical":           ControlsVertical,
	"copy":                       Copy,
	"copyClipboard":              CopyClipboard,
	"cross":                      Cross,
	"crossInCircleEmpty":         CrossInCircleEmpty,
	"crossInCircleFilled":        CrossInCircleFilled,
	"crosshairs":                 Crosshairs,
	"currency":                   Currency,
	"cut":                        Cut,
	"database":                   Database,
	"dockedBottom":               DockedBottom,
	"dockedDetached":             DockedDetached,
	"dockedLeft":                 DockedLeft,
	"dockedRight":                DockedRight,
	"dockedTakeover":             DockedTakeover,
	"dockedTop":                  DockedTop,
	"document":                   Document,
	"documentEdit":               DocumentEdit,
	"documentation":              Documentation,
	"documents":                  Documents,
	"dot":                        Dot,
	"download":                   Download,
	"editorAlignCenter":          EditorAlignCenter,
	"editorAlignLeft":            EditorAlignLeft,
	"editorAlignRight":           EditorAlignRight,
	"editorBold":                 EditorBold,
	"editorCodeBlock":            EditorCodeBlock,
	"editorComment":              EditorComment,
	"editorDistributeHorizontal": EditorDistributeHorizontal,
	"editorDistributeVertical":   EditorDistributeVertical,
	"editorHeading":              EditorHeading,
	"editorItalic":               EditorItalic,
	"editorItemAlignBottom":      EditorItemAlignBottom,
	"editorItemAlignCenter":      EditorItemAlignCenter,
	"editorItemAlignLeft":        EditorItemAlignLeft,
	"editorItemAlignMiddle":      EditorItemAlignMiddle,
	"editorItemAlignRight":       EditorItemAlignRight,
	"editorItemAlignTop":         EditorItemAlignTop,
	"editorLink":                 EditorLink,
	"editorOrderedList":          EditorOrderedList,
	"editorPositionBottomLeft":   EditorPositionBottomLeft,
	"editorPositionBottomRight":  EditorPositionBottomRight,
	"editorPositionTopLeft":      EditorPositionTopLeft,
	"editorPositionTopRight":     EditorPositionTopRight,
	"editorRedo":                 EditorRedo,
	"editorStrike":               EditorStrike,
	"editorTable":                EditorTable,
	"editorUnderline":            EditorUnderline,
	"editorUndo":                 EditorUndo,
	"editorUnorderedList":        EditorUnorderedList,
	"email":                      Email,
	"empty":                      Empty,
	"eql":                        Eql,
	"eraser":                     Eraser,
	"exit":                       Exit,
	"expand":                     Expand,
	"expandMini":                 ExpandMini,
	"export":                     Export,
	"eye":                        Eye,
	"eyeClosed":                  EyeClosed,
	"faceHappy":                  FaceHappy,
	"faceNeutral":                FaceNeutral,
	"faceSad":                    FaceSad,
	"filter":                     Filter,
	"flag":                       Flag,
	"fold":                       Fold,
	"folderCheck":                FolderCheck,
	"folderClosed":               FolderClosed,
	"folderExclamation":          FolderExclamation,
	"folderOpen":                 FolderOpen,
	"fullScreen":                 FullScreen,
	"fullScreenExit":             FullScreenExit,
	"function":                   Function,
	"gear":                       Gear,
	"glasses":                    Glasses,
	"globe":                      Globe,
	"grab":                       Grab,
	"grabHorizontal":             GrabHorizontal,
	"grid":                       Grid,
	"heart":                      Heart,
	"heatmap":                    Heatmap,
	"help":                       Help,
	"home":                       Home,
	"iinCircle":                  IinCircle,
	"image":                      Image,
	"import":                     Import,
	"indexClose":                 IndexClose,
	"indexEdit":                  IndexEdit,
	"indexFlush":                 IndexFlush,
	"indexMapping":               IndexMapping,
	"indexOpen":                  IndexOpen,
	"indexRuntime":               IndexRuntime,
	"indexSettings":              IndexSettings,
	"inputOutput":                InputOutput,
	"inspect":                    Inspect,
	"integrationGeneral":         IntegrationGeneral,
	"integrationObservability":   IntegrationObservability,
	"integrationSearch":          IntegrationSearch,
	"integrationSecurity":        IntegrationSecurity,
	"invert":                     Invert,
	"ip":                         Ip,
	"keyboardShortcut":           KeyboardShortcut,
	"kqlField":                   KqlField,
	"kqlFunction":                KqlFunction,
	"kqlOperand":                 KqlOperand,
	"kqlSelector":                KqlSelector,
	"kqlValue":                   KqlValue,
	"layers":                     Layers,
	"link":                       Link,
	"list":                       List,
	"listAdd":                    ListAdd,
	"lock":                       Lock,
	"lockOpen":                   LockOpen,
	"logstashFilter":             LogstashFilter,
	"logstashIf":                 LogstashIf,
	"logstashInput":              LogstashInput,
	"logstashOutput":             LogstashOutput,
	"logstashQueue":              LogstashQueue,
	"magnet":                     Magnet,
	"magnifyWithMinus":           MagnifyWithMinus,
	"magnifyWithPlus":            MagnifyWithPlus,
	"mapMarker":                  MapMarker,
	"memory":                     Memory,
	"menu":                       Menu,
	"menuDown":                   MenuDown,
	"menuLeft":                   MenuLeft,
	"menuRight":                  MenuRight,
	"menuUp":                     MenuUp,
	"merge":                      Merge,
	"minimize":                   Minimize,
	"minus":                      Minus,
	"minusInCircle":              MinusInCircle,
	"minusInCircleFilled":        MinusInCircleFilled,
	"mlClassificationJob":        MlClassificationJob,
	"mlCreateAdvancedJob":        MlCreateAdvancedJob,
	"mlCreateMultiMetricJob":     MlCreateMultiMetricJob,
	"mlCreatePopulationJob":      MlCreatePopulationJob,
	"mlCreateSingleMetricJob":    MlCreateSingleMetricJob,
	"mlDataVisualizer":           MlDataVisualizer,
	"mlOutlierDetectionJob":      MlOutlierDetectionJob,
	"mlRegressionJob":            MlRegressionJob,
	"mobile":                     Mobile,
	"moon":                       Moon,
	"nested":                     Nested,
	"node":                       Node,
	"number":                     Number,
	"offline":                    Offline,
	"online":                     Online,
	"package":                    Package,
	"pageSelect":                 PageSelect,
	"pagesSelect":                PagesSelect,
	"paint":                      Paint,
	"paperClip":                  PaperClip,
	"partial":                    Partial,
	"pause":                      Pause,
	"pencil":                     Pencil,
	"percent":                    Percent,
	"pin":                        Pin,
	"pinFilled":                  PinFilled,
	"play":                       Play,
	"playFilled":                 PlayFilled,
	"plus":                       Plus,
	"plusInCircle":               PlusInCircle,
	"plusInCircleFilled":         PlusInCircleFilled,
	"polygon":                    Polygon,
	"popout":                     Popout,
	"power":                      Power,
	"push":                       Push,
	"questionInCircle":           QuestionInCircle,
	"quote":                      Quote,
	"radius":                     Radius,
	"redeploy":                   Redeploy,
	"refresh":                    Refresh,
	"reporter":                   Reporter,
	"returnKey":                  ReturnKey,
	"save":                       Save,
	"scale":                      Scale,
	"search":                     Search,
	"securitySignal":             SecuritySignal,
	"securitySignalDetected":     SecuritySignalDetected,
	"securitySignalResolved":     SecuritySignalResolved,
	"shard":                      Shard,
	"share":                      Share,
	"snowflake":                  Snowflake,
	"sortDown":                   SortDown,
	"sortLeft":                   SortLeft,
	"sortRight":                  SortRight,
	"sortUp":                     SortUp,
	"sortable":                   Sortable,
	"starEmpty":                  StarEmpty,
	"starEmptySpace":             StarEmptySpace,
	"starFilled":                 StarFilled,
	"starFilledSpace":            StarFilledSpace,
	"starMinusEmpty":             StarMinusEmpty,
	"starMinusFilled":            StarMinusFilled,
	"starPlusEmpty":              StarPlusEmpty,
	"starPlusFilled":             StarPlusFilled,
	"stats":                      Stats,
	"stop":                       Stop,
	"stopFilled":                 StopFilled,
	"stopSlash":                  StopSlash,
	"storage":                    Storage,
	"string":                     String,
	"submodule":                  Submodule,
	"swatchInput":                SwatchInput,
	"symlink":                    Symlink,
	"tableDensityCompact":        TableDensityCompact,
	"tableDensityExpanded":       TableDensityExpanded,
	"tableDensityNormal":         TableDensityNormal,
	"tableOfContents":            TableOfContents,
	"tag":                        Tag,
	"tear":                       Tear,
	"temperature":                Temperature,
	"thumbsDown":                 ThumbsDown,
	"thumbsUp":                   ThumbsUp,
	"timeline":                   Timeline,
	"timeslider":                 Timeslider,
	"tokenAlias":                 TokenAlias,
	"tokenAnnotation":            TokenAnnotation,
	"tokenArray":                 TokenArray,
	"tokenBinary":                TokenBinary,
	"tokenBoolean":               TokenBoolean,
	"tokenClass":                 TokenClass,
	"tokenCompletionSuggester":   TokenCompletionSuggester,
	"tokenConstant":              TokenConstant,
	"tokenDate":                  TokenDate,
	"tokenDenseVector":           TokenDenseVector,
	"tokenElement":               TokenElement,
	"tokenEnum":                  TokenEnum,
	"tokenEnumMember":            TokenEnumMember,
	"tokenEvent":                 TokenEvent,
	"tokenException":             TokenException,
	"tokenField":                 TokenField,
	"tokenFile":                  TokenFile,
	"tokenFlattened":             TokenFlattened,
	"tokenFunction":              TokenFunction,
	"tokenGeo":                   TokenGeo,
	"tokenHistogram":             TokenHistogram,
	"tokenInterface":             TokenInterface,
	"tokenIp":                    TokenIp,
	"tokenJoin":                  TokenJoin,
	"tokenKey":                   TokenKey,
	"tokenKeyword":               TokenKeyword,
	"tokenMethod":                TokenMethod,
	"tokenModule":                TokenModule,
	"tokenNamespace":             TokenNamespace,
	"tokenNested":                TokenNested,
	"tokenNull":                  TokenNull,
	"tokenNumber":                TokenNumber,
	"tokenObject":                TokenObject,
	"tokenOperator":              TokenOperator,
	"tokenPackage":               TokenPackage,
	"tokenParameter":             TokenParameter,
	"tokenPercolator":            TokenPercolator,
	"tokenProperty":              TokenProperty,
	"tokenRange":                 TokenRange,
	"tokenRankFeature":           TokenRankFeature,
	"tokenRankFeatures":          TokenRankFeatures,
	"tokenRepo":                  TokenRepo,
	"tokenSearchType":            TokenSearchType,
	"tokenShape":                 TokenShape,
	"tokenString":                TokenString,
	"tokenStruct":                TokenStruct,
	"tokenSymbol":                TokenSymbol,
	"tokenText":                  TokenText,
	"tokenTokenCount":            TokenTokenCount,
	"tokenVariable":              TokenVariable,
	"training":                   Training,
	"trash":                      Trash,
	"undeploy":                   Undeploy,
	"unfold":                     Unfold,
	"unlink":                     Unlink,
	"user":                       User,
	"users":                      Users,
	"vector":                     Vector,
	"videoPlayer":                VideoPlayer,
	"visArea":                    VisArea,
	"visAreaStacked":             VisAreaStacked,
	"visBarHorizontal":           VisBarHorizontal,
	"visBarHorizontalStacked":    VisBarHorizontalStacked,
	"visBarVertical":             VisBarVertical,
	"visBarVerticalStacked":      VisBarVerticalStacked,
	"visBuilder":                 VisBuilder,
	"visBuilderSavedObject":      VisBuilderSavedObject,
	"visGauge":                   VisGauge,
	"visGoal":                    VisGoal,
	"visLine":                    VisLine,
	"visMapCoordinate":           VisMapCoordinate,
	"visMapRegion":               VisMapRegion,
	"visMetric":                  VisMetric,
	"visPie":                     VisPie,
	"visQueryDql":                VisQueryDql,
	"visQueryPpl":                VisQueryPpl,
	"visQueryPromql":             VisQueryPromql,
	"visQuerySql":                VisQuerySql,
	"visTable":                   VisTable,
	"visTagCloud":                VisTagCloud,
	"visText":                    VisText,
	"visTimelion":                VisTimelion,
	"visVega":                    VisVega,
	"visVisualBuilder":           VisVisualBuilder,
	"wordWrap":                   WordWrap,
	"wordWrapDisabled":           WordWrapDisabled,
	"wrench":                     Wrench,
}

func Accessibility(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 1 0 16A8 8 0 0 1 8 0m0 1a7 7 0 1 0 0 14A7 7 0 0 0 8 1m3.974 4.342a.5.5 0 0 1-.233.596l-.083.036L9 6.86v2.559l.974 2.923a.5.5 0 0 1-.233.596l-.083.036a.5.5 0 0 1-.596-.233l-.036-.083l-1-3L8 9.5l-.026.158l-1 3a.5.5 0 0 1-.97-.228l.022-.088L7 9.416V6.86l-2.658-.886a.5.5 0 0 1 .228-.97l.088.022L7.583 6h.833l2.926-.974a.5.5 0 0 1 .632.316M8 3a1 1 0 1 1 0 2a1 1 0 0 1 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Aggregate(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 2a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1m0-1a1.5 1.5 0 0 1 1.415 1h1.908a1.5 1.5 0 0 1 1.393.943L8.839 7H12.5a.5.5 0 0 1 0 1H8.839l-1.623 4.057A1.5 1.5 0 0 1 5.823 13H3.915a1.5 1.5 0 1 1 0-1h1.908a.5.5 0 0 0 .464-.314L7.761 8H3.915a1.5 1.5 0 1 1 0-1H7.76L6.287 3.314A.5.5 0 0 0 5.823 3H3.915A1.5 1.5 0 1 1 2.5 1m0 11a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1M3 7.5a.5.5 0 1 0-1 0a.5.5 0 0 0 1 0m9.354-3.354a.5.5 0 0 0-.708.708L13.793 7a.707.707 0 0 1 0 1l-2.147 2.146a.5.5 0 0 0 .708.708L14.5 8.707a1.707 1.707 0 0 0 0-2.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alert(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.59 10.059L7.35 5.18h1.3l-.25 4.88zm.394 1.901a.61.61 0 0 1-.448-.186a.606.606 0 0 1-.186-.444c0-.174.062-.323.186-.446a.614.614 0 0 1 .448-.184c.169 0 .315.06.44.182c.124.122.186.27.186.448a.6.6 0 0 1-.189.446a.607.607 0 0 1-.437.184M2 14a1 1 0 0 1-.878-1.479l6-11a1 1 0 0 1 1.756 0l6 11A1 1 0 0 1 14 14zm0-1h12L8 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AnalyzeEvent(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.924 4.013a.605.605 0 0 0-.228-.236L7.304.082a.607.607 0 0 0-.608 0L.304 3.777A.62.62 0 0 0 0 4.304v7.392a.61.61 0 0 0 .304.527l6.392 3.695c.188.11.42.11.608 0l6.392-3.695a.609.609 0 0 0 .304-.527V4.304a.607.607 0 0 0-.076-.291M1 5.079v6.391l6 3.47l6-3.47V5.08L7.252 8.432L7 8.579l-.252-.147zm11.476-.852L7 1.06L1.524 4.227L7 7.42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Annotation(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 6.5a4.5 4.5 0 1 1 5 4.473V16H7v-5.027A4.5 4.5 0 0 1 3 6.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AnomalyDetection(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1H2a1 1 0 0 0-1 1v2.5a.5.5 0 1 1-1 0V2a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v2.5a.5.5 0 1 1-1 0V2a1 1 0 0 0-1-1M.5 9a.5.5 0 0 1 .5.5V12a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V9.5a.5.5 0 1 1 1 0V12a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V9.5A.5.5 0 0 1 .5 9"/><path fill="currentColor" d="M9.482 3.968a.5.5 0 0 0-.964 0L7.362 8.606l-.915-1.83A.5.5 0 0 0 6 6.5H.5a.5.5 0 0 0 0 1h5.191l1.362 2.724a.5.5 0 0 0 .93-.092L9 6l1.018 4.132a.5.5 0 0 0 .93.092l1.36-2.724H13.5a.5.5 0 1 0 0-1H12a.5.5 0 0 0-.447.276l-.915 1.83z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ApmTrace(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 0h12a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2m0 1a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1zm.5 2h10a.5.5 0 1 1 0 1h-10a.5.5 0 0 1 0-1m1 3h6a.5.5 0 0 1 0 1h-6a.5.5 0 0 1 0-1m2 3h4a.5.5 0 0 1 0 1h-4a.5.5 0 0 1 0-1m3 3h5a.5.5 0 1 1 0 1h-5a.5.5 0 1 1 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppAddData(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M32 30H0V3h12.57l3 5H32zM2 28h28V10H14.43l-3-5H2z"/><path d="M21 18h-4v-4h-2v4h-4v2h4v4h2v-4h4z" class="ouiIcon__fillSecondary"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppAdvancedSettings(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.909 26.182h1.939v4.848H2.909z" class="ouiIcon__fillSecondary"/><path fill="currentColor" d="M4.848 16.62V0H2.91v16.62a3.879 3.879 0 1 0 1.94 0m-.97 5.683a1.94 1.94 0 1 1 0-3.879a1.94 1.94 0 0 1 0 3.879"/><path fill="currentColor" d="M14.545 16.485h1.939V31.03h-1.939z" class="ouiIcon__fillSecondary"/><path fill="currentColor" d="M16.485 6.924V0h-1.94v6.924a3.879 3.879 0 1 0 1.94 0m-.97 5.682a1.94 1.94 0 1 1 0-3.879a1.94 1.94 0 0 1 0 3.88"/><path fill="currentColor" d="M26.182 26.182h1.939v4.848h-1.939z" class="ouiIcon__fillSecondary"/><path fill="currentColor" d="M28.121 16.62V0h-1.94v16.62a3.879 3.879 0 1 0 1.94 0m-.97 5.683a1.94 1.94 0 1 1 0-3.879a1.94 1.94 0 0 1 0 3.879"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppApm(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 10h4v2H1V1h30v6h-2V3H3zm26 19v-6h2v8H18v-8h2v6z"/><path fill="currentColor" d="M31 10H9v11h12c5.523 0 10-4.477 10-10zm-10 9H11v-7h17.938A8.001 8.001 0 0 1 21 19" class="ouiIcon__fillSecondary"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppAppSearch(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.5 11.916L27 5.278L19.5.938a7.002 7.002 0 0 0-7 0l-8 4.62A7 7 0 0 0 1 11.62v9.237a7 7 0 0 0 3.5 6.061l7.5 4.33V17.98a7 7 0 0 1 3.5-6.062zM10 27.786v-9.809a9 9 0 0 1 4.5-7.792l8.503-4.91L18.5 2.67a5.003 5.003 0 0 0-5 0l-8 4.619A5 5 0 0 0 3 11.62v9.238a5 5 0 0 0 2.5 4.33z"/><path fill="currentColor" fill-rule="evenodd" d="M18.409 13.55a7.089 7.089 0 0 1 1.035 1.711A6.93 6.93 0 0 1 20 17.98v13.27l7.5-4.33a7 7 0 0 0 3.5-6.061V11.62a6.992 6.992 0 0 0-1.587-4.422zm2.777.705A8.933 8.933 0 0 1 22 17.979v9.807l4.5-2.597a5 5 0 0 0 2.5-4.33V11.62c0-.588-.106-1.161-.303-1.7z" class="ouiIcon__fillSecondary"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppAuditbeat(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 0h2v32h-2z" class="ouiIcon__fillSecondary"/><path fill="currentColor" d="M0 32h13v-2H2V2h11V0H0zM19 0v2h11v28H19v2h13V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppCanvas(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 17h2v7H7zm5-3h2v10h-2zm5 2h2v8h-2zm5-2h3v2h-3zm0 4h3v2h-3zm0 4h3v2h-3z" class="ouiIcon__fillSecondary"/><path fill="currentColor" d="M30.73 24a6.47 6.47 0 0 1 .45-2.19c.337-.9.52-1.85.54-2.81a8.55 8.55 0 0 0-.54-2.81a6.47 6.47 0 0 1-.45-2.19a9.2 9.2 0 0 1 .62-2.49c.53-1.57 1.08-3.19.08-4.2c-1-1.01-2.41-.44-3.52.05a5.59 5.59 0 0 1-2.09.64a5.3 5.3 0 0 1-.59 0L16 .28L6.77 8a5.3 5.3 0 0 1-.59 0a5.59 5.59 0 0 1-2.09-.65C3 6.87 1.6 6.25.57 7.31c-1.03 1.06-.45 2.63.08 4.2A9.2 9.2 0 0 1 1.27 14a6.47 6.47 0 0 1-.45 2.19A8.55 8.55 0 0 0 .28 19c.02.96.203 1.91.54 2.81A6.47 6.47 0 0 1 1.27 24a9.2 9.2 0 0 1-.62 2.49c-.53 1.57-1.08 3.19-.08 4.2c.353.38.852.59 1.37.58a5.67 5.67 0 0 0 2.15-.63A5.59 5.59 0 0 1 6.18 30a7.13 7.13 0 0 1 2.29.47a8 8 0 0 0 2.62.53a7.37 7.37 0 0 0 2.47-.51A7.14 7.14 0 0 1 16 30a6.24 6.24 0 0 1 2.14.45a8 8 0 0 0 2.77.55a8.08 8.08 0 0 0 2.77-.55a6.24 6.24 0 0 1 2.14-.45a5.59 5.59 0 0 1 2.09.65c1.11.49 2.49 1.11 3.52.05c1.03-1.06.45-2.63-.08-4.2a9.2 9.2 0 0 1-.62-2.5M21.17 7h-.26a8 8 0 0 0-2.77.55A6.24 6.24 0 0 1 16 8a6.24 6.24 0 0 1-2.14-.45A8 8 0 0 0 11.09 7h-.26L16 2.72zm8.89 22.27a4.42 4.42 0 0 1-1.34-.46a7.08 7.08 0 0 0-2.9-.82a8.14 8.14 0 0 0-2.78.55a6.13 6.13 0 0 1-2.13.45a6.24 6.24 0 0 1-2.14-.45A8 8 0 0 0 16 28a9 9 0 0 0-3.08.6a5.74 5.74 0 0 1-1.83.4a6.36 6.36 0 0 1-2-.43A8.72 8.72 0 0 0 6.18 28a7.08 7.08 0 0 0-2.9.82a9.65 9.65 0 0 1-1.28.52a6.08 6.08 0 0 1 .52-2.21c.403-1 .65-2.055.73-3.13a8.55 8.55 0 0 0-.54-2.81A6.47 6.47 0 0 1 2.27 19a6.47 6.47 0 0 1 .44-2.19c.337-.9.52-1.85.54-2.81a10.48 10.48 0 0 0-.72-3.13a9 9 0 0 1-.59-2.16H2c.447.1.88.255 1.29.46a7.08 7.08 0 0 0 2.9.82A8.14 8.14 0 0 0 9 9.44A6.13 6.13 0 0 1 11.09 9a6.13 6.13 0 0 1 2.13.45A8.14 8.14 0 0 0 16 10a8.14 8.14 0 0 0 2.78-.55A6.13 6.13 0 0 1 20.91 9a6.13 6.13 0 0 1 2.09.44a8.14 8.14 0 0 0 2.78.55a7.08 7.08 0 0 0 2.9-.82A9.65 9.65 0 0 1 30 8.66a6.08 6.08 0 0 1-.52 2.21c-.403 1-.65 2.055-.73 3.13c.02.96.203 1.91.54 2.81a6.47 6.47 0 0 1 .44 2.19a6.47 6.47 0 0 1-.44 2.19a8.55 8.55 0 0 0-.54 2.81c.078 1.074.32 2.13.72 3.13a9 9 0 0 1 .59 2.16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppCode(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.276 29l.594 2H0l7.621-14.29l.811 2.73L3.333 29zm17.644 2l-4.987-16.598A16 16 0 0 0 8.688 3l1.8 6H8.4L6 1h2.607a18 18 0 0 1 17.241 12.828L31 31z"/><path fill="currentColor" d="M12.037 14.02L16.492 29h6.827l-2.333-7.849a10 10 0 0 0-8.949-7.13M9.35 12h2.05a12 12 0 0 1 11.503 8.581L26 31H15z" class="ouiIcon__fillSecondary"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppConsole(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M29 9H3a3 3 0 0 1-3-3V3a3 3 0 0 1 3-3h26a3 3 0 0 1 3 3v3a3 3 0 0 1-3 3M3 2a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h26a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1z"/><path fill="currentColor" d="M29 32H3a3 3 0 0 1-3-3V14a3 3 0 0 1 3-3h26a3 3 0 0 1 3 3v15a3 3 0 0 1-3 3M3 13a1 1 0 0 0-1 1v15a1 1 0 0 0 1 1h26a1 1 0 0 0 1-1V14a1 1 0 0 0-1-1z" class="ouiIcon__fillSecondary"/><path fill="currentColor" d="m7.29 17.71l3.3 3.29l-3.3 3.29l1.42 1.42l4.7-4.71l-4.7-4.71zM15 24h9v2h-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppCrossClusterReplication(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v16h8.7l5.3-6V0zm2 2h10v7H7v5H2zm8.45 9L9 12.64V11zM18 16v16h8.7l5.3-6V16zm2 2h10v7h-5v5h-5zm8.45 9L27 28.64V27z"/><path fill="currentColor" d="M5 18H3c0 6.075 4.925 11 11 11h2v-2h-2a9 9 0 0 1-9-9M18 3h-2v2h2a9 9 0 0 1 9 9h2c0-6.075-4.925-11-11-11" class="ouiIcon__fillSecondary"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppDashboard(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M29 9H3a3 3 0 0 1-3-3V3a3 3 0 0 1 3-3h26a3 3 0 0 1 3 3v3a3 3 0 0 1-3 3M3 2a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h26a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1z"/><path fill="currentColor" d="M12 20H3a3 3 0 0 1-3-3v-3a3 3 0 0 1 3-3h9a3 3 0 0 1 3 3v3a3 3 0 0 1-3 3m-9-7a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1z" class="ouiIcon__fillSecondary"/><path fill="currentColor" d="M12 31H3a3 3 0 0 1-3-3v-3a3 3 0 0 1 3-3h9a3 3 0 0 1 3 3v3a3 3 0 0 1-3 3m-9-7a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1z"/><path fill="currentColor" d="M29 31h-9a3 3 0 0 1-3-3V14a3 3 0 0 1 3-3h9a3 3 0 0 1 3 3v14a3 3 0 0 1-3 3m-9-18a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1V14a1 1 0 0 0-1-1z" class="ouiIcon__fillSecondary"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppDevtools(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M21 28h-2v-8.49l.6-.26A9 9 0 0 0 21 3.52V11H11V3.52a9 9 0 0 0 1.4 15.73l.6.26V28h-2v-7.21A11 11 0 0 1 11.6.92L13 .31V9h6V.31l1.4.61a11 11 0 0 1 .6 19.87z"/><path d="M11 30h10v2H11z" class="ouiIcon__fillSecondary"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppDiscover(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.33 23.67l4.79-10.55l10.55-4.79l-4.79 10.55zm6.3-9l-2.28 5l5-2.28l2.28-5z" class="ouiIcon__fillSecondary"/><path fill="currentColor" d="M16 0C7.163 0 0 7.163 0 16s7.163 16 16 16s16-7.163 16-16A16 16 0 0 0 16 0m1 29.95V28h-2v1.95A14 14 0 0 1 2.05 17H4v-2H2.05A14 14 0 0 1 15 2.05V4h2V2.05A14 14 0 0 1 29.95 15H28v2h1.95A14 14 0 0 1 17 29.95"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppEms(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M3 22h3v2H1V1h23v5h-2V3H3z" class="ouiIcon__fillSecondary"/><path d="M15.228 29c1.492-1.678 1.353-2.859.009-5.654c-.049-.1-.049-.1-.097-.203c-1.369-2.855-1.626-4.491-.325-6.582c2.796-4.498 9.514-2.642 14.185 2.317V10H10v19zm2.47 0H29v-6.983c-3.88-5.406-10.376-7.795-12.487-4.4c-.83 1.336-.669 2.37.43 4.662l.097.2c1.352 2.814 1.731 4.63.659 6.521M31 8v23H8V8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppFilebeat(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8 18h16v2H8zm0-5h9v2H8zm0 10h16v2H8z" class="ouiIcon__fillSecondary"/><path d="M21.41 0H5a3 3 0 0 0-3 3v26a3 3 0 0 0 3 3h22a3 3 0 0 0 3-3V8.59zM22 3.41L26.59 8H22zM27 30H5a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1h15v8h8v19a1 1 0 0 1-1 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppGis(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m29.014 23.89l2.296 1.145L16 32.101L.53 24.961l2.49-1.056l2.453 1.132l-.003.002L16 29.899l10.69-4.934l-.003-.001zm-4.972-7.482l7.268 3.627L16 27.101L.53 19.961l7.668-3.252c.392.486.838 1.02 1.34 1.604L5.47 20.039L16 24.899l10.69-4.934l-3.954-1.973c.493-.58.928-1.107 1.306-1.584" class="ouiIcon__fillSecondary"/><path fill="currentColor" d="M18 9a2 2 0 1 0-4 0a2 2 0 0 0 4 0m2 0a4 4 0 1 1-8 0a4 4 0 0 1 8 0m-3.268 12.681l-.732.787l-.732-.787c-3.557-3.824-5.817-6.462-6.81-7.96A8.746 8.746 0 0 1 7 8.875C7 3.97 11.033 0 16 0s9 3.97 9 8.875a8.746 8.746 0 0 1-1.459 4.846c-.992 1.498-3.252 4.136-6.809 7.96m5.142-9.064A6.747 6.747 0 0 0 23 8.875C23 5.081 19.87 2 16 2S9 5.081 9 8.875c0 1.349.394 2.636 1.126 3.742c.846 1.277 2.812 3.593 5.874 6.912c3.062-3.32 5.028-5.635 5.874-6.912"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppGraph(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24 20a4 4 0 1 1 0-8a4 4 0 0 1 0 8m0-6a2 2 0 1 0 0 4a2 2 0 0 0 0-4m-8.2-5.62A4 4 0 1 1 18 1.06a4 4 0 0 1-2.2 7.32m0-6a2 2 0 1 0 .01 0zm.01 29.24a4 4 0 1 1-.083-8a4 4 0 0 1 .083 8m0-6a2 2 0 1 0 .39 0a2 2 0 0 0-.4 0z" class="ouiIcon__fillSecondary"/><path fill="currentColor" d="M18 17v-2h-6.14a4 4 0 0 0-.86-1.64l2.31-3.44l-1.68-1.12l-2.31 3.44A4 4 0 0 0 8 12a4 4 0 1 0 0 8a4 4 0 0 0 1.32-.24l2.31 3.44l1.66-1.12L11 18.64a4 4 0 0 0 .86-1.64zM6 16a2 2 0 1 1 4 0a2 2 0 0 1-4 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppGrok(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.727 2.91V.97c-1.357 0-4.18 0-5.818 1.94c-.97 1.182-1.328 2.908-.97 4.983c.563 3.307.97 6.332.301 7.157c-.3.31-.97.465-2.24.465v1.94c1.27 0 1.94.155 2.25.494c.698.824.262 3.879-.3 7.156c-.35 2.124-.01 3.802.96 4.956C5.527 32 8.35 32 9.726 32v-1.94c-1.144 0-3.287 0-4.315-1.221a4.112 4.112 0 0 1-.533-3.414c.843-5.042.97-7.37-.136-8.727c-.068-.077-.145-.136-.213-.203c-.068-.068.145-.127.213-.204c1.144-1.367.97-3.695.136-8.727a4.112 4.112 0 0 1 .543-3.414C6.44 2.91 8.583 2.91 9.727 2.91M32.03 17.454v-1.94c-1.27 0-1.94-.155-2.25-.494c-.698-.825-.261-3.88.301-7.157c.35-2.123 0-3.801-.97-4.984C27.493.94 24.672.94 23.294.94v1.97c1.145 0 3.288 0 4.315 1.22a4.112 4.112 0 0 1 .544 3.414c-.844 5.043-.97 7.37.135 8.728c.068.077.146.135.214.203c.067.068-.146.126-.214.204c-1.144 1.367-.97 3.694-.135 8.727a4.111 4.111 0 0 1-.544 3.413c-1.018 1.242-3.16 1.242-4.305 1.242V32c1.358 0 4.18 0 5.818-1.94c.97-1.182 1.329-2.908.97-4.984c-.563-3.306-.97-6.332-.3-7.156c.3-.31.969-.465 2.24-.465"/><path fill="currentColor" d="M26.212 15.515h-2.007a7.758 7.758 0 0 0-6.72-6.72V6.788h-1.94v2.007a7.758 7.758 0 0 0-6.72 6.72H6.818v1.94h2.007a7.758 7.758 0 0 0 6.72 6.72v2.007h1.94v-2.007a7.758 7.758 0 0 0 6.72-6.72h2.007zm-8.727 6.7v-1.851h-1.94v1.852a5.818 5.818 0 0 1-4.76-4.761h1.851v-1.94h-1.852a5.818 5.818 0 0 1 4.761-4.761v1.852h1.94v-1.852a5.818 5.818 0 0 1 4.761 4.761h-1.852v1.94h1.852a5.818 5.818 0 0 1-4.761 4.76" class="ouiIcon__fillSecondary"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppHeartbeat(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.14 15.39a8.058 8.058 0 0 1-2.2-4.043A6.788 6.788 0 0 1 4.198 5.47a6.73 6.73 0 0 1 8.727-.213l1.26-1.464a8.65 8.65 0 0 0-11.277.232A8.727 8.727 0 0 0 .068 11.6a10.172 10.172 0 0 0 2.793 5.275z" class="ouiIcon__fillSecondary"/><path fill="currentColor" d="M15.515 31.274L4.548 18.454L15.855 4.763a8.67 8.67 0 0 1 12.266-.746a8.727 8.727 0 0 1 2.91 7.205c-.243 2.695-2.037 4.732-3.482 6.37zm-8.427-12.82l8.427 9.862l10.55-11.995c1.32-1.503 2.822-3.21 3.007-5.265a6.788 6.788 0 0 0-2.24-5.586a6.73 6.73 0 0 0-9.504.563z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppIndexManagement(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M17 18v-2h-2v2H3v6h2v-4h10v4h2v-4h10v4h2v-6z"/><path d="M4 32a3 3 0 1 1 0-6a3 3 0 0 1 0 6m0-4a1 1 0 1 0 0 2a1 1 0 0 0 0-2m12 4a3 3 0 1 1 0-6a3 3 0 0 1 0 6m0-4a1 1 0 1 0 0 2a1 1 0 0 0 0-2m12 4a3 3 0 1 1 0-6a3 3 0 0 1 0 6m0-4a1 1 0 1 0 0 2a1 1 0 0 0 0-2M23 8V6h-2.1a5 5 0 0 0-.73-1.75l1.49-1.49l-1.42-1.42l-1.49 1.49A5 5 0 0 0 17 2.1V0h-2v2.1a5 5 0 0 0-1.75.73l-1.49-1.49l-1.42 1.42l1.49 1.49A5 5 0 0 0 11.1 6H9v2h2.1a5 5 0 0 0 .73 1.75l-1.49 1.49l1.41 1.41l1.49-1.49a5 5 0 0 0 1.76.74V14h2v-2.1a5 5 0 0 0 1.75-.73l1.49 1.49l1.41-1.41l-1.48-1.5A5 5 0 0 0 20.9 8zm-7 2a3 3 0 1 1 0-6a3 3 0 0 1 0 6" class="ouiIcon__fillSecondary"/><path d="M16 8a1 1 0 0 1-1-1a1.39 1.39 0 0 1 0-.2a.65.65 0 0 1 .06-.18a.74.74 0 0 1 .09-.18a1.61 1.61 0 0 1 .12-.15a.93.93 0 0 1 .33-.21a1 1 0 0 1 1.09.21l.12.15a.78.78 0 0 1 .09.18a.62.62 0 0 1 .1.18a1.27 1.27 0 0 1 0 .2a1 1 0 0 1-1 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppIndexPattern(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M20 14h-8a3 3 0 0 1-3-3V3a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3M12 2a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1zm5 16v-2h-2v2H3v6h2v-4h10v4h2v-4h10v4h2v-6z"/><path d="M4 32a3 3 0 1 1 0-6a3 3 0 0 1 0 6m0-4a1 1 0 1 0 0 2a1 1 0 0 0 0-2m12 4a3 3 0 1 1 0-6a3 3 0 0 1 0 6m0-4a1 1 0 1 0 0 2a1 1 0 0 0 0-2m12 4a3 3 0 1 1 0-6a3 3 0 0 1 0 6m0-4a1 1 0 1 0 0 2a1 1 0 0 0 0-2M13 4h6v2h-6zm0 4h6v2h-6z" class="ouiIcon__fillSecondary"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppIndexRollup(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M32 26v-2h-2.1a5 5 0 0 0-.73-1.75l1.49-1.49l-1.41-1.41l-1.49 1.49A5 5 0 0 0 26 20.1V18h-2v2.1a5 5 0 0 0-1.75.73l-1.49-1.49l-1.41 1.41l1.49 1.49A5 5 0 0 0 20.1 24H18v2h2.1a5 5 0 0 0 .73 1.75l-1.49 1.49l1.41 1.41l1.49-1.49a5 5 0 0 0 1.76.74V32h2v-2.1a5 5 0 0 0 1.75-.73l1.49 1.49l1.41-1.41l-1.49-1.49A5 5 0 0 0 29.9 26zm-7 2a3 3 0 1 1 0-6a3 3 0 0 1 0 6"/><path fill="currentColor" d="M25.71 24.29a1 1 0 0 0-1.09-.21a1.15 1.15 0 0 0-.33.21a.93.93 0 0 0-.21.33a1 1 0 0 0 1.3 1.3a1.15 1.15 0 0 0 .33-.21a1 1 0 0 0 .21-1.09a.94.94 0 0 0-.21-.33"/><path fill="currentColor" d="M5 6h16v2H5zm0 6h16v2H5zm0 6h10v2H5zm0 6h8v2H5z" class="ouiIcon__fillSecondary"/><path fill="currentColor" d="M16 32H3a3 3 0 0 1-3-3V3a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v13h-2V3a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v26a1 1 0 0 0 1 1h13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppLens(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m23.793 14.293l1.414 1.414l-6.408 6.409l-3.75-3.25l-4.342 4.341l-1.414-1.414l5.658-5.659l3.75 3.25zM12 11v5l-2 2v-7zm10-6v8l-2 2V5zm-5 3v7l-2-2V8z" class="ouiIcon__fillSecondary"/><path fill="currentColor" d="M17 0c8.284 0 15 6.716 15 15c0 8.284-6.716 15-15 15c-3.782 0-7.238-1.4-9.876-3.71l-5.417 5.417l-1.414-1.414l5.417-5.417A14.943 14.943 0 0 1 2 15c0-1.05.108-2.074.313-3.062l1.906.672A13.076 13.076 0 0 0 4 15c0 7.18 5.82 13 13 13s13-5.82 13-13S24.18 2 17 2c-1.002 0-1.978.113-2.915.328l-.75-1.877A15.031 15.031 0 0 1 17 0M9.621 1.937l.75 1.877a13.062 13.062 0 0 0-4.82 5.024l-1.907-.673a15.068 15.068 0 0 1 5.977-6.228"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppLogs(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 8c3.983 0 7.732 1.013 11.001 2.797v2.312A20.887 20.887 0 0 0 2 10.023v11.025c4.85.462 9.27 4.183 9.955 8.691l.035.261H12v2h-1.938l-.018-1.236c-.116-4.015-4.336-7.631-8.793-7.76L0 22.986V8zm13-8h1c9.28 0 16.825 7.437 16.997 16.677L32 17v15h-2V17c0-7.84-6.014-14.274-13.68-14.943L16 2.033v7.681l-2-1.143zh1z"/><path fill="currentColor" d="M26.997 30.636L27.009 32H14V11.305l1.483.82c6.994 3.861 11.382 10.735 11.514 18.51zm-2.048-1.04C24.505 23.556 21.205 18.2 16 14.771V30h8.974z" class="ouiIcon__fillSecondary"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppManagement(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 21a5 5 0 1 1 0-10a5 5 0 0 1 0 10m0-8a3 3 0 1 0 0 6a3 3 0 0 0 0-6" class="ouiIcon__fillSecondary"/><path fill="currentColor" d="M20 32h-8v-4.06a1 1 0 0 0-1.61-.67l-2.88 2.87l-5.65-5.65l2.87-2.87a.92.92 0 0 0 .2-1a.93.93 0 0 0-.86-.6H0V12h4.06a.92.92 0 0 0 .85-.58a.94.94 0 0 0-.19-1L1.86 7.51l5.65-5.65l2.87 2.87A1 1 0 0 0 12 4.06V0h8v4.06a1 1 0 0 0 1.61.67l2.87-2.87l5.66 5.66l-2.87 2.87a.92.92 0 0 0-.2 1a.93.93 0 0 0 .86.6H32v8h-4.06a.92.92 0 0 0-.85.58a.94.94 0 0 0 .19 1l2.87 2.87l-5.66 5.66l-2.87-2.87a1 1 0 0 0-1.61.67zm-6-2h4v-2.06a3 3 0 0 1 5-2.08l1.46 1.46l2.83-2.83L25.86 23a3 3 0 0 1 2.08-5H30v-4h-2.06a3 3 0 0 1-2.08-5l1.46-1.46l-2.83-2.85L23 6.14a3 3 0 0 1-5-2.08V2h-4v2.06a3 3 0 0 1-5 2.08L7.51 4.69L4.69 7.51L6.14 9a3 3 0 0 1-2.08 5H2v4h2.06a3 3 0 0 1 2.08 5l-1.45 1.49l2.83 2.83L9 25.86a3 3 0 0 1 5 2.08z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppMetricbeat(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16a16 16 0 0 1-16 16m0-30C8.268 2 2 8.268 2 16s6.268 14 14 14s14-6.268 14-14A14 14 0 0 0 16 2"/><path fill="currentColor" d="M28 16h-2c0-5.523-4.477-10-10-10S6 10.477 6 16H4C4 9.373 9.373 4 16 4s12 5.373 12 12" class="ouiIcon__fillSecondary"/><path fill="currentColor" d="M21.71 11.71L20.3 10.3L18 12.57a4 4 0 0 0-2-.57a4 4 0 1 0 4 4a4 4 0 0 0-.57-2zM16 18a2 2 0 1 1 0-4a2 2 0 0 1 0 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppMetrics(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M30 19.092v12.88H2v-5.386l6.747-6.747l.708.708c.236.236.48.463.733.68L4 27.414v2.558h24v-8.91c.186-.166.369-.339.546-.516zm-20.85-3.19A10.955 10.955 0 0 1 8 11C8 4.925 12.925 0 19 0s11 4.925 11 11c0 1.76-.414 3.425-1.15 4.9l-1.51-1.51A8.973 8.973 0 0 0 28 11a9 9 0 1 0-17.34 3.391z"/><path fill="currentColor" d="M19 20a8.96 8.96 0 0 0 5.618-1.968l-4.202-4.204a2 2 0 0 0-2.828 0l-4.205 4.205A8.96 8.96 0 0 0 19 20m-2.826-7.586a4 4 0 0 1 5.656 0l5.656 5.657l-.707.707A10.967 10.967 0 0 1 19 22a10.967 10.967 0 0 1-7.778-3.221l-.707-.707z" class="ouiIcon__fillSecondary"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppMl(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 18v.56a1 1 0 0 1-.68.95L3 21.61V10a1 1 0 0 1 .4-.8l3.2-2.4l-1.2-1.6l-3.2 2.4A3 3 0 0 0 1 10v12a3 3 0 0 0 1.2 2.4l3.2 2.4l1.2-1.6l-2.47-1.85l5.82-1.95A3 3 0 0 0 12 18.56V18zM29.8 7.6l-3.2-2.4l-1.2 1.6l3.2 2.4a1 1 0 0 1 .4.8v11.61l-6.32-2.11a1 1 0 0 1-.68-.95V18h-2v.56a3 3 0 0 0 2.05 2.85l5.82 1.94l-2.47 1.85l1.2 1.6l3.2-2.4A3 3 0 0 0 31 22V10a3 3 0 0 0-1.2-2.4" class="ouiIcon__fillSecondary"/><path fill="currentColor" d="M11 6A3 3 0 0 1 8.88.88a3.07 3.07 0 0 1 4.24 0A3 3 0 0 1 11 6m0-4a1 1 0 1 0-.012 2A1 1 0 0 0 11 2m0 30a3 3 0 0 1-2.12-5.12a3.07 3.07 0 0 1 4.24 0A3 3 0 0 1 11 32m0-4a1 1 0 1 0-.012 2A1 1 0 0 0 11 28m0-12a3 3 0 0 1-2.12-5.12a3.07 3.07 0 0 1 4.24 0A3 3 0 0 1 11 16m0-4a1 1 0 1 0-.012 2A1 1 0 0 0 11 12m10-6A3 3 0 0 1 18.88.88a3.07 3.07 0 0 1 4.24 0A3 3 0 0 1 21 6m0-4a1 1 0 1 0-.012 2A1 1 0 0 0 21 2m0 30a3 3 0 0 1-2.12-5.12a3.07 3.07 0 0 1 4.24 0A3 3 0 0 1 21 32m0-4a1 1 0 1 0-.012 2A1 1 0 0 0 21 28m0-12a3 3 0 0 1-2.12-5.12a3.07 3.07 0 0 1 4.24 0A3 3 0 0 1 21 16m0-4a1 1 0 1 0-.012 2A1 1 0 0 0 21 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppMonitoring(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.81 15.19A8.94 8.94 0 0 1 15.62 3.86l.38.42l.38-.42a8.94 8.94 0 0 1 14.26 10.68l-1.7-1.07a6.94 6.94 0 0 0-11.07-8.28L16 7.29l-1.87-2.1A6.94 6.94 0 0 0 3.41 14zM16 31.18l-7.74-8.51l1.48-1.34L16 28.21l6.26-6.88l1.48 1.34z"/><path fill="currentColor" d="m16.16 23.29l-4.1-7.17L10.62 19H0v-2h9.38l2.56-5.12l3.9 6.83l4.13-10.32L23.66 17H32v2h-9.66l-2.31-5.39z" class="ouiIcon__fillSecondary"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppNotebook(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M25 2h-5V0h-2v2h-3V0h-2v2h-3V0H8v2H3v26h22zm-2 24H5V4h3v2h2V4h3v2h2V4h3v2h2V4h3z"/><path d="M27 7v23H8v2h21V7z"/><path d="M8 12h12v2H8zm0 5h12v2H8z" class="ouiIcon__fillSecondary"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppPacketbeat(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 20a4 4 0 1 1 0-8a4 4 0 0 1 0 8m0-6a2 2 0 1 0 0 4a2 2 0 0 0 0-4" class="ouiIcon__fillSecondary"/><path fill="currentColor" d="M8 4a4 4 0 1 0-4 4a4 4 0 0 0 2-.57l5.27 5.27l1.41-1.41L7.43 6A4 4 0 0 0 8 4M4 6a2 2 0 1 1 0-4a2 2 0 0 1 0 4m16.71 6.71L26 7.43A4 4 0 0 0 28 8a4 4 0 1 0-4-4a4 4 0 0 0 .57 2l-5.27 5.27zM28 2a2 2 0 1 1 0 4a2 2 0 0 1 0-4M11.29 19.29L6 24.57A4 4 0 0 0 4 24a4 4 0 1 0 4 4a4 4 0 0 0-.57-2l5.27-5.27zM4 30a2 2 0 1 1 0-4a2 2 0 0 1 0 4m24-6a4 4 0 0 0-2 .57l-5.27-5.27l-1.41 1.41L24.57 26a4 4 0 0 0-.57 2a4 4 0 1 0 4-4m0 6a2 2 0 1 1 0-4a2 2 0 0 1 0 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppPipeline(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M29 12a3 3 0 0 0-3 3h-4a3 3 0 0 0-3-3h-6a3 3 0 0 0-3 3H6a3 3 0 0 0-3-3H0v14h3a3 3 0 0 0 3-3h4a3 3 0 0 0 3 3h6a3 3 0 0 0 3-3h4a3 3 0 0 0 3 3h3V12zM3 24H2V14h1a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1m17-3v2a1 1 0 0 1-1 1h-6a1 1 0 0 1-1-1v-2H6v-4h6v-2a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v2h6v4zm10 3h-1a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1h1z"/><path d="M22 6H10v2h5v2h2V8h5z" class="ouiIcon__fillSecondary"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppRecentlyViewed(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16a16 16 0 0 1-16 16m0-30C8.268 2 2 8.268 2 16s6.268 14 14 14s14-6.268 14-14A14 14 0 0 0 16 2"/><path fill="currentColor" d="M15 4h2v9.17A3.009 3.009 0 0 1 18.83 15H26v2h-7.17A3.001 3.001 0 1 1 15 13.17zm1 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2" class="ouiIcon__fillSecondary"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppReporting(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M25 5h-.17v2H25a1 1 0 0 1 1 1v20a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h.17V5H7a3 3 0 0 0-3 3v20a3 3 0 0 0 3 3h18a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3"/><path d="M23 3h-3V0h-8v3H9v6h14zm-2 4H11V5h3V2h4v3h3z"/><path d="M10 13h12v2H10zm0 5h12v2H10zm0 5h12v2H10z" class="ouiIcon__fillSecondary"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppSavedObjects(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.38 7.001L17 9.619V0h-2v9.619l-3.37-2.618l-1.25 1.513L16 12.878l5.63-4.364z" class="ouiIcon__fillSecondary"/><path fill="currentColor" d="m21.06.165l-1.11 1.61l9.25 5.983L16 16.29L2.8 7.758l9.25-5.983l-1.1-1.61L0 7.234v13.653l16 10.337l16-10.337V7.234zM2 9.57l13 8.407v10.279L2 19.84zm15 18.676V17.978l13-8.407V19.85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppSearchProfiler(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.63 8h7.38v2h-7.38z" class="ouiIcon__fillSecondary"/><path fill="currentColor" d="M7 8h3.19v2H7z"/><path fill="currentColor" d="M7 16h7.38v2H7z" class="ouiIcon__fillSecondary"/><path fill="currentColor" d="M15.81 16H19v2h-3.19zM7 12h9v2H7z"/><path fill="currentColor" d="M13 0C5.82 0 0 5.82 0 13s5.82 13 13 13s13-5.82 13-13A13 13 0 0 0 13 0m0 24C6.925 24 2 19.075 2 13S6.925 2 13 2s11 4.925 11 11s-4.925 11-11 11m9.581-.007l1.414-1.414l7.708 7.708l-1.414 1.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppSecurity(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14 32l-.36-.14A21.07 21.07 0 0 1 0 12.07V5.44L14 .06l14 5.38v6.63a21.07 21.07 0 0 1-13.64 19.78zM2 6.82v5.25a19.08 19.08 0 0 0 12 17.77a19.08 19.08 0 0 0 12-17.77V6.82L14 2.2z"/><path fill="currentColor" d="M9 17.83h2V23H9zm2-7.65V7H9v3.18a3 3 0 1 0 2 0M10 14a1 1 0 1 1 0-2a1 1 0 0 1 0 2m7-7h2v5.17h-2zm4 10a3 3 0 1 0-4 2.82V23h2v-3.18A3 3 0 0 0 21 17m-3 1a1 1 0 1 1 0-2a1 1 0 0 1 0 2" class="ouiIcon__fillSecondary"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppSecurityAnalytics(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 10h4v2H5v7.928c0 1.299.808 2.795 2.88 4.48c1.83 1.489 4.524 3.02 8.12 4.584V26h2v5.992l-1.38-.567c-4.372-1.797-7.724-3.613-10-5.465C4.358 24.122 3 22.114 3 19.928z"/><path fill="currentColor" d="M9 10h9v14l-1.272-.458c-1.367-.494-3.23-1.314-4.768-2.39C10.484 20.118 9 18.636 9 16.761zm1.895 1.876v4.887c0 .877.744 1.867 2.158 2.856c.937.656 2.038 1.219 3.052 1.657v-9.4z" class="ouiIcon__fillSecondary"/><path fill="currentColor" d="M29 1H9v7h2V2.966h16V16.73c0 .558-.245 1.128-.756 1.72c-.515.596-1.256 1.158-2.12 1.668c-1.381.818-2.961 1.434-4.124 1.817V24c1.26-.378 3.334-1.12 5.155-2.197c.965-.57 1.905-1.261 2.612-2.08c.712-.822 1.233-1.827 1.233-2.992z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppSpaces(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 4h6v2H4zm18 0h6v2h-6zM4 22h6v2H4z" class="ouiIcon__fillSecondary"/><path fill="currentColor" d="M0 14h14V0H0zM2 2h10v10H2zm16-2v14h14V0zm12 12H20V2h10zM0 32h14V18H0zm2-12h10v10H2zm16 12h14V18H18zm2-12h10v10H20z"/><path fill="currentColor" d="M22 22h6v2h-6z" class="ouiIcon__fillSecondary"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppSql(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 6h9v2h-9zM5 6h9v2H5zm0 6h9v2H5zm13 0h9v2h-9zM5 18h9v2H5zm13 0h9v2h-9zm0 6h9v2h-9zM5 24h9v2H5z" class="ouiIcon__fillSecondary"/><path fill="currentColor" d="M29 32H3a3 3 0 0 1-3-3V3a3 3 0 0 1 3-3h26a3 3 0 0 1 3 3v26a3 3 0 0 1-3 3M3 2a1 1 0 0 0-1 1v26a1 1 0 0 0 1 1h26a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppTimelion(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M4 4v20.34L16 32l12-7.64V4zm22 2v4h-4a3 3 0 0 0-3 3v5h-6v-5a3 3 0 0 0-3-3H6V6zm-7.87 14L16 22.52L13.87 20zM6 12h4a1 1 0 0 1 1 1v6.7l3.69 4.37l-2.58 3.06L6 23.24zm7.81 16.22l2.19-2.6l2.19 2.6L16 29.61zm6.08-1.09l-2.58-3.06L21 19.7V13a1 1 0 0 1 1-1h4v11.24z"/><path d="M4 0h24v2H4z" class="ouiIcon__fillSecondary"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppUpgradeAssistant(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 21a5 5 0 1 1 0-10a5 5 0 0 1 0 10m0-8a3 3 0 1 0 0 6a3 3 0 0 0 0-6" class="ouiIcon__fillSecondary"/><path fill="currentColor" d="M27.42 19.69a12 12 0 0 1-23.11-1l2.27-.45l-4.32-4.47L0 19.55l2.39-.47a14 14 0 0 0 27 1.23zm2.23-6.77a14 14 0 0 0-27-1.23l1.9.62a12 12 0 0 1 23.11 1l-2.27.45l4.32 4.46L32 12.45z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppUptime(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.216 12.377A10.948 10.948 0 0 0 2.181 17H.153a12.941 12.941 0 0 1 2.693-6.118zm2.948-2.703l-1.37-1.495A12.94 12.94 0 0 1 13 6v2c-2.144 0-4.144.613-5.836 1.674" class="ouiIcon__fillSecondary"/><path fill="currentColor" d="M26 4.414V19c0 7.18-5.82 13-13 13C6.5 32 1.115 27.23.153 21H2.18c.94 5.12 5.427 9 10.819 9c6.075 0 11-4.925 11-11V4.414l-4.293 4.293l-1.414-1.414L25 .586l6.707 6.707l-1.414 1.414zm-7.836 9.909l1.472 1.354l-7.577 8.235l-4.835-4.442l1.353-1.473l3.364 3.09z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppUsersRoles(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.307 3.21a2.91 2.91 0 1 0-.223 1.94a11.636 11.636 0 0 1 8.232 7.049l1.775-.698a13.576 13.576 0 0 0-9.784-8.291m-2.822 1.638a.97.97 0 1 1 0-1.939a.97.97 0 0 1 0 1.94m-4.267.805l-.717-1.774a13.576 13.576 0 0 0-8.291 9.784a2.91 2.91 0 1 0 1.94.223a11.636 11.636 0 0 1 7.068-8.233m-8.34 11.802a.97.97 0 1 1 0-1.94a.97.97 0 0 1 0 1.94m12.607 8.727a2.91 2.91 0 0 0-2.599 1.62a11.636 11.636 0 0 1-8.233-7.05l-1.774.717a13.576 13.576 0 0 0 9.813 8.291a2.91 2.91 0 1 0 2.793-3.578m0 3.879a.97.97 0 1 1 0-1.94a.97.97 0 0 1 0 1.94M32 16.485a2.91 2.91 0 1 0-4.199 2.599a11.636 11.636 0 0 1-7.05 8.232l.718 1.775a13.576 13.576 0 0 0 8.291-9.813A2.91 2.91 0 0 0 32 16.485m-2.91.97a.97.97 0 1 1 0-1.94a.97.97 0 0 1 0 1.94"/><path fill="currentColor" d="M19.19 16.35a3.879 3.879 0 1 0-5.42 0a4.848 4.848 0 0 0-2.134 4.014v1.939h9.697v-1.94a4.848 4.848 0 0 0-2.143-4.014m-4.645-2.774a1.94 1.94 0 1 1 3.88 0a1.94 1.94 0 0 1-3.88 0m-.97 6.788a2.91 2.91 0 1 1 5.819 0z" class="ouiIcon__fillSecondary"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppVisualize(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M32 32H4a4 4 0 0 1-4-4V0h2v28a2 2 0 0 0 2 2h28z" class="ouiIcon__fillSecondary"/><path fill="currentColor" d="M6 20h2v7H6zm10-8h2v15h-2zm10 5h2v10h-2z"/><path fill="currentColor" d="M27 6a3 3 0 0 0-2.08.84L20 4.36A2.2 2.2 0 0 0 20 4a3 3 0 0 0-6 0c.001.341.062.68.18 1l-5.6 4.46A3 3 0 0 0 7 9a3 3 0 1 0 3 3a2.93 2.93 0 0 0-.18-1l5.6-4.48A3 3 0 0 0 17 7a3 3 0 0 0 2.08-.84l5 2.48A2.2 2.2 0 0 0 24 9a3 3 0 1 0 3-3M7 13a1 1 0 1 1 0-2a1 1 0 0 1 0 2m10-8a1 1 0 1 1 0-2a1 1 0 0 1 0 2m10 5a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppWatches(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.74 7.73l-1.5-1.32a13 13 0 0 0 0 17.19l1.5-1.32a11 11 0 0 1 0-14.54z"/><path fill="currentColor" d="M6.51 3.66L5 2.34c-6.377 7.24-6.377 18.09 0 25.33l1.5-1.32C.792 19.867.792 10.153 6.5 3.67zm17.25 2.75l-1.5 1.32a11 11 0 0 1 0 14.54l1.5 1.32a13 13 0 0 0 0-17.19z"/><path fill="currentColor" d="m27 2.34l-1.5 1.32c5.708 6.483 5.708 16.197 0 22.68l1.5 1.33c6.377-7.24 6.377-18.09 0-25.33"/><path fill="currentColor" d="M21 15a5 5 0 1 0-6 4.9V31h2V19.9a5 5 0 0 0 4-4.9m-5 3a3 3 0 1 1 0-6a3 3 0 0 1 0 6" class="ouiIcon__fillSecondary"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppWorkplaceSearch(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.81 1H2v30h5.815c1.705 0 3.343-.783 4.38-2.094l3.182-4.023l.006.006l2.607-3.137l.01-.014a9.42 9.42 0 0 0 1.953-4.826a9.112 9.112 0 0 0-1.891-6.522l-.01-.012l-2.603-3.156L12.197 3.1C11.16 1.786 9.52 1 7.81 1m6.118 7.523l-3.31-4.195C9.968 3.512 8.917 3 7.81 3H4v26h3.815c1.1 0 2.15-.508 2.8-1.32l3.33-4.21l-1.796-1.77l-.027-.033c-2.84-3.317-2.83-8.195.028-11.287l.004-.004a98.893 98.893 0 0 1 1.774-1.853m4.043 8.099c-.12 1.279-.616 2.537-1.49 3.64l-1.288 1.63l-1.578-1.556c-2.19-2.59-2.118-6.3.001-8.596c.6-.64 1.119-1.183 1.561-1.635l1.293 1.638v.002l.014.016c.942 1.192 1.446 2.564 1.51 3.949a7.09 7.09 0 0 1-.023.912"/><path fill="currentColor" fill-rule="evenodd" d="M23.801 3c-.382 0-1.052.231-1.925.812c-.636.422-1.228.928-1.677 1.357l1.771 2.245l.005.007l.024.03l.008.01c4.003 5.099 3.99 12.057-.035 17.145l-1.793 2.267a8.206 8.206 0 0 0 1.769 1.498c.698.427 1.337.629 1.852.629H28V3zm-6.148 23.842l.765 1.053l.023.027c.717.851 1.566 1.607 2.464 2.156c.89.544 1.89.922 2.895.922H30V1h-6.199c-.987 0-2.066.505-3.032 1.146a14.934 14.934 0 0 0-2.352 1.974l-.276.288l-.595.628l2.855 3.619l.006.006c3.459 4.375 3.457 10.329-.004 14.704z" class="ouiIcon__fillSecondary"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Apps(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 4V2h2v2zm5 0V2h2v2zm5 0V2h2v2zM2 9V7h2v2zm5 0V7h2v2zm5 0V7h2v2zM2 14v-2h2v2zm5 0v-2h2v2zm5 0v-2h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDown(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="non-zero" d="M13.069 5.157L8.384 9.768a.546.546 0 0 1-.768 0L2.93 5.158a.552.552 0 0 0-.771 0a.53.53 0 0 0 0 .759l4.684 4.61a1.65 1.65 0 0 0 2.312 0l4.684-4.61a.53.53 0 0 0 0-.76a.552.552 0 0 0-.771 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.843 13.069L6.232 8.384a.546.546 0 0 1 0-.768l4.61-4.685a.552.552 0 0 0 0-.771a.53.53 0 0 0-.759 0l-4.61 4.684a1.65 1.65 0 0 0 0 2.312l4.61 4.684a.53.53 0 0 0 .76 0a.552.552 0 0 0 0-.771"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.157 13.069l4.611-4.685a.546.546 0 0 0 0-.768L5.158 2.93a.552.552 0 0 1 0-.771a.53.53 0 0 1 .759 0l4.61 4.684a1.65 1.65 0 0 1 0 2.312l-4.61 4.684a.53.53 0 0 1-.76 0a.552.552 0 0 1 0-.771"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUp(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.931 10.843l4.685-4.611a.546.546 0 0 1 .768 0l4.685 4.61a.55.55 0 0 0 .771 0a.53.53 0 0 0 0-.759l-4.684-4.61a1.65 1.65 0 0 0-2.312 0l-4.684 4.61a.53.53 0 0 0 0 .76a.55.55 0 0 0 .771 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Asterisk(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.928 8.468L4 7.618l.446-1.427L7.375 7.25L7.287 4h1.484l-.097 3.296l2.88-1.039L12 7.693l-2.977.86l1.92 2.56L9.741 12L7.937 9.28l-1.745 2.654l-1.213-.86z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Beaker(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.277 10.088c.02.014.04.03.057.047c.582.55 1.134.812 1.666.812c.586 0 1.84-.293 3.713-.88L9 6.212V2H7v4.212zm-.438.987L3.539 14h8.922l-1.32-2.969C9.096 11.677 7.733 12 7 12c-.74 0-1.463-.315-2.161-.925M6 2H5V1h6v1h-1v4l3.375 7.594A1 1 0 0 1 12.461 15H3.54a1 1 0 0 1-.914-1.406L6 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.316 12h10.368c-.188-.704-.28-1.691-.348-3.037c-.07-1.382-.103-1.888-.19-2.612c-.028-.236-.06-.462-.096-.68c-.31-1.892-1.506-2.923-3.708-3.131a1 1 0 1 0-1.684 0c-2.202.208-3.397 1.24-3.708 3.13a16.01 16.01 0 0 0-.096.68c-.087.725-.12 1.23-.19 2.613c-.068 1.346-.16 2.333-.348 3.037m10.843 1H1.84c-.308.353-.737.5-1.341.5a.5.5 0 1 1 0-1c.786 0 1.024-.783 1.166-3.587c.07-1.407.105-1.926.196-2.681c.03-.25.063-.49.102-.724c.334-2.041 1.546-3.313 3.556-3.792a2 2 0 0 1 3.96 0c2.01.479 3.222 1.75 3.557 3.792a17 17 0 0 1 .102.724c.09.755.125 1.274.196 2.681c.14 2.804.379 3.587 1.165 3.587a.5.5 0 1 1 0 1c-.604 0-1.033-.147-1.341-.5M5.5 14h4a2 2 0 1 1-4 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellSlash(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.5 14l-.005.15a2 2 0 0 1-3.972.153l-.018-.154L5.5 14zm5.38-12.825a.5.5 0 0 1 .008.64l-.063.065l-14 12a.5.5 0 0 1-.713-.695l.063-.065l14-12a.5.5 0 0 1 .705.055M13.036 5.51l.055.353l.048.368c.066.554.103.98.145 1.724l.08 1.491c.146 2.371.402 3.053 1.136 3.053a.5.5 0 1 1 0 1c-.603 0-1.031-.146-1.34-.499H4.185l1.182-1h7.317c-.172-.644-.264-1.528-.33-2.708l-.09-1.638c-.033-.514-.066-.87-.118-1.304l-.011-.077zM7.5 0a2 2 0 0 1 1.98 1.717c.476.113.907.27 1.292.472l-.838.71a5.728 5.728 0 0 0-1.591-.36a1 1 0 1 0-1.684 0C4.455 2.75 3.26 3.78 2.95 5.671l-.05.334l-.046.347c-.08.676-.115 1.161-.176 2.347l-.014.265l-.005.09l-1.058.897c.018-.234.035-.488.05-.763l.077-1.427c.038-.626.073-1.025.134-1.528c.03-.25.063-.49.102-.724c.334-2.04 1.546-3.313 3.555-3.792l.004-.019A2 2 0 0 1 7.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bolt(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.04 13.274a.5.5 0 1 0 .892.453l3.014-5.931a.5.5 0 0 0-.445-.727H5.316L8.03 1.727a.5.5 0 1 0-.892-.453L4.055 7.343a.5.5 0 0 0 .446.726h5.185z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxesHorizontal(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 6h4v4H0zm1 1v2h2V7zm5-1h4v4H6zm1 1v2h2V7zm5-1h4v4h-4zm1 3h2V7h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxesVertical(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 1v2h2V1zM6 0h4v4H6zm0 6h4v4H6zm1 1v2h2V7zm-1 5h4v4H6zm1 1v2h2v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Branch(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 10.038a3.49 3.49 0 0 1 2.5-1.05h2a2.5 2.5 0 0 0 2.462-2.061a2 2 0 1 1 1.008.017A3.5 3.5 0 0 1 9.5 9.987h-2a2.5 2.5 0 0 0-2.466 2.085A2 2 0 1 1 4 12.063V3.937a2 2 0 1 1 1 0zM4.5 3a1 1 0 1 0 0-2a1 1 0 0 0 0 2m0 12a1 1 0 1 0 0-2a1 1 0 0 0 0 2m8-9a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Broom(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.732 13.096l-.197-.197l2.83-2.828l.706.707l-2.829 2.828l.708.708l2.828-2.828a1 1 0 0 0 0-1.414L5.658 7.95a.993.993 0 0 0-.708-.293a.994.994 0 0 0-.708.293l-2.828 2.829l.707.707l2.829-2.83l.707.708l-2.829 2.829zm1.218-6.44c.512 0 1.023.196 1.414.587l2.121 2.12a2 2 0 0 1 0 2.83L4.95 15.728L0 10.778l3.535-3.535a1.993 1.993 0 0 1 1.415-.586M14.02 1l.708.707l-6.95 6.95l-.707-.707z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Brush(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.993 8.17c0 .83-.673 1.507-1.499 1.507H5.498A1.505 1.505 0 0 1 3.999 8.17V6.662h7.994zm-2.998 5.998c0 .455-.448.827-.999.827c-.55 0-1-.372-1-.827v-3.486h2zM4 5.658V1.005h1.262v4.653zm2.261 0V1.005h1.244v4.653zm2.244 0V1.005H9.74v4.653zm2.234 0V1.005h1.254v4.653zM3.008 0L3 8.17a2.509 2.509 0 0 0 2.498 2.512h.5v3.486c0 1.01.896 1.832 1.998 1.832c1.102 0 1.998-.822 1.998-1.832v-3.486h.5a2.509 2.509 0 0 0 2.498-2.512L13 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bug(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.651 5.126l.922-.455l.884-2.343a.5.5 0 0 1 .939.344L12.374 5.39l-1.45.717A5.3 5.3 0 0 1 11 7h1.043l2.3 2.198a.5.5 0 0 1-.692.723L11.642 8h-.737c-.09.466-.24.899-.441 1.283l.892.49l1.278 3.554a.5.5 0 0 1-.94.342l-1.15-3.2l-.655-.36C9.373 10.665 8.716 11 8 11s-1.374-.335-1.89-.893l-.658.361l-1.15 3.201a.5.5 0 1 1-.94-.342l1.279-3.554l.895-.491A4.7 4.7 0 0 1 5.095 8h-.74l-2.01 1.92a.5.5 0 0 1-.69-.722L3.952 7H5a5.3 5.3 0 0 1 .075-.892L3.623 5.39L2.6 2.672a.5.5 0 1 1 .939-.344l.884 2.343l.924.457c.17-.428.397-.81.668-1.128h.57a1.5 1.5 0 1 1 2.83 0h.568c.27.318.497.699.667 1.126M8 4a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1m1.751 1.571A3.326 3.326 0 0 0 9.476 5H6.524c-.107.175-.2.367-.276.573c-.11.295-.186.618-.223.957a4.354 4.354 0 0 0 .09 1.465c.071.294.172.565.295.806c.168.328.38.601.616.803c.295.253.631.396.974.396c.342 0 .678-.142.973-.394c.237-.203.448-.476.616-.803a3.62 3.62 0 0 0 .296-.807a4.263 4.263 0 0 0 .09-1.466a3.988 3.988 0 0 0-.224-.959"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bullseye(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 14A6 6 0 1 0 8 2a6 6 0 0 0 0 12m0 1A7 7 0 1 1 8 1a7 7 0 0 1 0 14m0-3a4 4 0 1 0 0-8a4 4 0 0 0 0 8m0-1a3 3 0 1 1 0-6a3 3 0 0 1 0 6m0-2a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14 4v-.994C14 2.45 13.55 2 12.994 2H11v1h-1V2H6v1H5V2H3.006C2.45 2 2 2.45 2 3.006v9.988C2 13.55 2.45 14 3.006 14h9.988C13.55 14 14 13.55 14 12.994V5H2V4zm-3-3h1.994C14.102 1 15 1.897 15 3.006v9.988A2.005 2.005 0 0 1 12.994 15H3.006A2.005 2.005 0 0 1 1 12.994V3.006C1 1.898 1.897 1 3.006 1H5V0h1v1h4V0h1zM4 7h2v1H4zm3 0h2v1H7zm3 0h2v1h-2zM4 9h2v1H4zm3 0h2v1H7zm3 0h2v1h-2zm-6 2h2v1H4zm3 0h2v1H7zm3 0h2v1h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatLeft(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.46 14.923A1 1 0 0 1 3.846 14v-1.619H3a2 2 0 0 1-2-2v-7.38a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v7.38a2 2 0 0 1-2 2H7.91l-2.36 2.331a1 1 0 0 1-1.089.211m.386-2.542a1 1 0 0 0-1-1H3a1 1 0 0 1-1-1v-7.38a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v7.38a1 1 0 0 1-1 1H7.91a1 1 0 0 0-.702.289L4.846 14z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatRight(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.54 14.923a1 1 0 0 0 .614-.923v-1.619H13a2 2 0 0 0 2-2v-7.38a2 2 0 0 0-2-2H3a2 2 0 0 0-2 2v7.38a2 2 0 0 0 2 2h5.09l2.362 2.331a1 1 0 0 0 1.088.211m-.386-2.542a1 1 0 0 1 1-1H13a1 1 0 0 0 1-1v-7.38a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v7.38a1 1 0 0 0 1 1h5.09a1 1 0 0 1 .702.289L11.154 14z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.5 12a.502.502 0 0 1-.354-.146l-4-4a.502.502 0 0 1 .708-.708L6.5 10.793l6.646-6.647a.502.502 0 0 1 .708.708l-7 7A.502.502 0 0 1 6.5 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckInCircleEmpty(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.91 9.8l4.74-4.657a.5.5 0 1 1 .7.714l-5.09 5a.5.5 0 0 1-.701 0L3.649 8a.5.5 0 1 1 .701-.714z"/><path fill="currentColor" fill-rule="evenodd" d="M8 16A8 8 0 1 1 8 0a8 8 0 0 1 0 16m0-1A7 7 0 1 1 8 1a7 7 0 0 1 0 14" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckInCircleFilled(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 16A8 8 0 1 1 8 0a8 8 0 0 1 0 16m3.65-10.857L6.91 9.8L4.35 7.286a.5.5 0 0 0-.7.714l2.909 2.857a.5.5 0 0 0 .7 0l5.091-5a.5.5 0 1 0-.7-.714"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cheer(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.934 3.063a1.5 1.5 0 0 1 .547.321l.112.115l6.07 6.915a1.5 1.5 0 0 1-.646 2.41l-.142.04l-9.031 2.097A1.5 1.5 0 0 1 .037 13.19l.043-.159L3.04 4.02a1.5 1.5 0 0 1 1.893-.957M4.027 4.25l-.036.083l-2.961 9.011a.5.5 0 0 0 .498.656l.09-.013l2.937-.681l-1.399-1.508a.5.5 0 0 1 .666-.74l.067.06l1.788 1.927l2.634-.611l-3.198-3.601a.5.5 0 0 1 .682-.726l.066.062l3.559 4.007l1.229-.284a.5.5 0 0 0 .15-.063l.067-.049a.5.5 0 0 0 .099-.632l-.053-.073l-6.07-6.916a.5.5 0 0 0-.138-.11l-.082-.035l-.088-.02a.5.5 0 0 0-.507.256m11.66 5.039a2.5 2.5 0 0 1-.975 3.399a.5.5 0 0 1-.485-.875a1.5 1.5 0 0 0-1.454-2.624a.5.5 0 0 1-.485-.875a2.5 2.5 0 0 1 3.399.975m-5.03-6.206a.5.5 0 0 1 .338.544l-.02.088l-.677 2.035l2.068-.721a.5.5 0 0 1 .6.225l.036.082a.5.5 0 0 1-.225.6l-.082.037L9.67 7.028a.5.5 0 0 1-.659-.55l.02-.08l.995-3a.5.5 0 0 1 .632-.316M14.5 4a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1M7.862.403a2.5 2.5 0 0 1 .735 3.459a.5.5 0 0 1-.839-.545a1.5 1.5 0 1 0-2.516-1.634a.5.5 0 0 1-.839-.545A2.5 2.5 0 0 1 7.862.403M13.5 2a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1m-3-1a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1m4-1a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 13A5.506 5.506 0 0 1 2 7.5C2 4.467 4.467 2 7.5 2S13 4.467 13 7.5S10.533 13 7.5 13m0-12a6.5 6.5 0 1 0 0 13a6.5 6.5 0 0 0 0-13m3 6H8V3.5a.5.5 0 0 0-1 0v4a.5.5 0 0 0 .5.5h3a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDrizzle(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.348 3.761A3.995 3.995 0 0 1 8 7a.5.5 0 0 1-1 0a3 3 0 1 0-4.878 2.34a.5.5 0 0 1-.627.779a4 4 0 0 1 3.973-6.84a5.502 5.502 0 0 1 10.096 4.37a.5.5 0 1 1-.92-.39a4.5 4.5 0 1 0-8.296-3.497m-1.61 4.935a.5.5 0 1 1 .775.633l-1.466 1.792a.5.5 0 1 1-.774-.633zm-3.12 3.647a.5.5 0 0 1 .774.634l-1.505 1.84a.5.5 0 0 1-.774-.634zm7.62-3.647a.5.5 0 0 1 .775.633l-1.466 1.792a.5.5 0 1 1-.774-.633zm-3.12 3.647a.5.5 0 0 1 .774.634l-1.505 1.84a.5.5 0 0 1-.774-.634zm7.62-3.647a.5.5 0 1 1 .775.633l-1.466 1.792a.5.5 0 1 1-.774-.633zm-3.12 3.647a.5.5 0 0 1 .774.634l-1.505 1.84a.5.5 0 0 1-.774-.634z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudStormy(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.421 4.93a.5.5 0 1 1-.87.49a3 3 0 1 0-4.43 3.918a.5.5 0 0 1-.626.78a4 4 0 0 1 3.973-6.84l.032.018V3.28a5.5 5.5 0 1 1 7.003 7.357a.5.5 0 1 1-.36-.934a4.5 4.5 0 1 0-5.77-5.923c.42.31.778.701 1.05 1.15zM9.6 11c.669.002.794.67.36 1.003l-4.68 3.882c-.457.378-1.053-.26-.643-.689l3.08-3.193A5411.7 5411.7 0 0 1 5.113 12c-.668-.001-.793-.669-.36-1.003l4.68-3.881c.458-.379 1.053.26.643.688l-3.08 3.193z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudSunny(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.746 5.005A5.5 5.5 0 0 1 10.5 16H4a4 4 0 0 1-1.61-7.663A4.473 4.473 0 0 1 2.029 7H.5a.5.5 0 0 1 0-1h1.527a4.479 4.479 0 0 1 .957-2.309L1.646 2.354a.5.5 0 1 1 .708-.708L3.69 2.984A4.479 4.479 0 0 1 6 2.027V.5a.5.5 0 0 1 1 0v1.528a4.493 4.493 0 0 1 2.309.956l1.337-1.338a.5.5 0 0 1 .708.708L10.016 3.69c.311.388.56.831.73 1.314M4 15h6.5a4.5 4.5 0 1 0-4.152-6.239A3.995 3.995 0 0 1 8 12a.5.5 0 1 1-1 0a3 3 0 1 0-3 3m5.691-9.94a3.5 3.5 0 1 0-6.33 2.991a4.029 4.029 0 0 1 2.106.227a5.505 5.505 0 0 1 4.224-3.219"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Color(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 11c1.334 1.393 2 2.435 2 3.125C15 15.161 14.105 16 13 16c-1.104 0-2-.84-2-1.875c0-.69.667-1.732 2-3.125M5.857.15l6.34 6.45l.016.02l.324.321a1.5 1.5 0 0 1 .11 2.006l-.103.114l-4.474 4.513a1.5 1.5 0 0 1-2.123.008L1.464 9.06a1.5 1.5 0 0 1 .007-2.12l4.472-4.45c.145-.146.313-.254.492-.327L5.144.85a.5.5 0 0 1 .713-.7m1.496 3.049a.5.5 0 0 0-.705 0L2.177 7.65a.498.498 0 0 0-.148.35h9.95a.498.498 0 0 0-.148-.35L7.353 3.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compute(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 3a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1zm10 8v1a2 2 0 0 1-2 2h-1v2h-1v-2H8.5v2h-1v-2H6v2H5v-2H4a2 2 0 0 1-2-2v-1H0v-1h2V8.5H0v-1h2V6H0V5h2V4a2 2 0 0 1 2-2h1V0h1v2h1.5V0h1v2H10V0h1v2h1a2 2 0 0 1 2 2v1h2v1h-2v1.5h2v1h-2V10h2v1z"/><rect width="6" height="6" x="5" y="5" fill="currentColor" rx="1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Console(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.157 12.224L5.768 8.32a.404.404 0 0 0 0-.64l-4.61-3.904a.407.407 0 0 1 0-.643a.608.608 0 0 1 .759 0l4.61 3.904c.631.534.63 1.393 0 1.926l-4.61 3.904a.608.608 0 0 1-.76 0a.407.407 0 0 1 0-.643M9 12h6v1H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContinuityAbove(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 3a.5.5 0 0 0-.5.5v9a.5.5 0 0 0 1 0v-9a.5.5 0 0 0-.5-.5m3 0a.5.5 0 0 0-.5.5v9a.5.5 0 0 0 1 0V9h6v1.5a.5.5 0 0 0 .82.384l3-2.5a.5.5 0 0 0 0-.768l-3-2.5A.5.5 0 0 0 11 5.5V7H5V3.5a.5.5 0 0 0-.5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContinuityAboveBelow(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.5 3a.5.5 0 0 1 .5.5v9a.5.5 0 0 1-1 0V9H4v1.5a.5.5 0 0 1-.82.384l-3-2.5a.5.5 0 0 1 0-.768l3-2.5A.5.5 0 0 1 4 5.5V7h2V3.5a.5.5 0 0 1 .5-.5m3 0a.5.5 0 0 0-.5.5v9a.5.5 0 0 0 1 0V9h2v1.5a.5.5 0 0 0 .82.384l3-2.5a.5.5 0 0 0 0-.768l-3-2.5A.5.5 0 0 0 12 5.5V7h-2V3.5a.5.5 0 0 0-.5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContinuityBelow(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.5 3a.5.5 0 0 1 .5.5v9a.5.5 0 0 1-1 0V9H5v1.5a.5.5 0 0 1-.82.384l-3-2.5a.5.5 0 0 1 0-.768l3-2.5A.5.5 0 0 1 5 5.5V7h6V3.5a.5.5 0 0 1 .5-.5m3 0a.5.5 0 0 1 .5.5v9a.5.5 0 0 1-1 0v-9a.5.5 0 0 1 .5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContinuityWithin(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.5 3a.5.5 0 0 0-.5.5v9a.5.5 0 0 0 1 0v-9A.5.5 0 0 0 .5 3m14.5.5a.5.5 0 0 1 1 0v9a.5.5 0 0 1-1 0zm-4.712 1.547a.5.5 0 0 1 .532.069l3 2.5a.5.5 0 0 1 0 .768l-3 2.5A.5.5 0 0 1 10 10.5V9H6v1.5a.5.5 0 0 1-.82.384l-3-2.5a.5.5 0 0 1 0-.768l3-2.5A.5.5 0 0 1 6 5.5V7h4V5.5a.5.5 0 0 1 .288-.453"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ControlsHorizontal(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.05 10a2.5 2.5 0 0 1 4.9 0h1.55a.5.5 0 1 1 0 1h-1.55a2.5 2.5 0 0 1-4.9 0H1.5a.5.5 0 1 1 0-1zm-.1-4a2.5 2.5 0 0 1-4.9 0H1.5a.5.5 0 0 1 0-1h1.55a2.5 2.5 0 0 1 4.9 0h6.55a.5.5 0 1 1 0 1zM4 5.5a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0-3 0m8 5a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ControlsVertical(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 7.95a2.5 2.5 0 0 1 0-4.9V1.5a.5.5 0 1 1 1 0v1.55a2.5 2.5 0 0 1 0 4.9v6.55a.5.5 0 1 1-1 0zm-4 .1a2.5 2.5 0 0 1 0 4.9v1.55a.5.5 0 1 1-1 0v-1.55a2.5 2.5 0 0 1 0-4.9V1.5a.5.5 0 0 1 1 0zM5.5 12a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m5-8a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copy(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.4 0a.85.85 0 0 1 .622.273l2.743 3a.884.884 0 0 1 .235.602v9.25a.867.867 0 0 1-.857.875H3.857A.867.867 0 0 1 3 13.125V.875C3 .392 3.384 0 3.857 0zM14 4h-2.6a.4.4 0 0 1-.4-.4V1H4v12h10z"/><path fill="currentColor" d="M3 1H2a1 1 0 0 0-1 1v13a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-1h-1v1H2V2h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyClipboard(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 2.729V2a1 1 0 0 1 1-1h2v1H1v12h4v1H1a1 1 0 0 1-1-1V9zM12 5V2a1 1 0 0 0-1-1H9v1h2v3zm-1 1h2v9H6V6zV5H6a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h7a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1h-2z"/><path fill="currentColor" d="M7 10h5V9H7zm0-2h5V7H7zm0 4h5v-1H7zm0 2h5v-1H7zM9 2V1a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v1h1V1h4v1zM3 3h6V2H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cross(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.293 8L3.146 3.854a.5.5 0 1 1 .708-.708L8 7.293l4.146-4.147a.5.5 0 0 1 .708.708L8.707 8l4.147 4.146a.5.5 0 0 1-.708.708L8 8.707l-4.146 4.147a.5.5 0 0 1-.708-.708z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CrossInCircleEmpty(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14m0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16" clip-rule="evenodd"/><path fill="currentColor" d="M11.854 4.854a.5.5 0 0 0-.707-.707L8 7.293L4.854 4.147a.5.5 0 1 0-.707.707L7.293 8l-3.146 3.146a.5.5 0 0 0 .707.708L8 8.707l3.147 3.147a.5.5 0 0 0 .707-.708L8.708 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CrossInCircleFilled(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.746 8l3.1-3.1a.527.527 0 1 0-.746-.746L8 7.254l-3.1-3.1a.527.527 0 1 0-.746.746l3.1 3.1l-3.1 3.1a.527.527 0 1 0 .746.746l3.1-3.1l3.1 3.1a.527.527 0 1 0 .746-.746zM8 16A8 8 0 1 1 8 0a8 8 0 0 1 0 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crosshairs(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.822 1.874a.5.5 0 1 1 .335.942a5.517 5.517 0 0 0-3.34 3.341a.5.5 0 1 1-.943-.335a6.517 6.517 0 0 1 3.948-3.948M1.864 10.15a.5.5 0 1 1 .944-.33a5.517 5.517 0 0 0 3.365 3.37a.5.5 0 0 1-.333.943a6.517 6.517 0 0 1-3.976-3.983m8.302 3.981a.5.5 0 1 1-.333-.943a5.517 5.517 0 0 0 3.347-3.332a.5.5 0 1 1 .941.337a6.517 6.517 0 0 1-3.955 3.938m3.968-8.285a.5.5 0 1 1-.943.331A5.517 5.517 0 0 0 9.85 2.82a.5.5 0 0 1 .337-.942a6.517 6.517 0 0 1 3.946 3.968M8.5 3.5a.5.5 0 0 1-1 0V.997a.5.5 0 0 1 1 0zm-4.997 4a.5.5 0 0 1 0 1H1a.5.5 0 0 1 0-1zM7.5 12.497a.5.5 0 0 1 1 0V15a.5.5 0 1 1-1 0zM12.497 8.5a.5.5 0 0 1 0-1H15a.5.5 0 1 1 0 1zM8 9a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Currency(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.95 1.636l1.414 1.414l-2.192 2.193C12.695 6.033 13 6.98 13 8s-.305 1.967-.828 2.757l2.192 2.193l-1.414 1.414l-2.193-2.192A4.977 4.977 0 0 1 8 13a4.977 4.977 0 0 1-2.757-.828L3.05 14.364L1.636 12.95l2.192-2.193A4.977 4.977 0 0 1 3 8c0-1.02.305-1.967.828-2.757L1.636 3.05L3.05 1.636l2.193 2.192A4.977 4.977 0 0 1 8 3c1.02 0 1.967.305 2.757.828zM8 5a2.99 2.99 0 0 0-1.168.236l-.126.057l-.218.116l-.132.081l-.073.05a3.013 3.013 0 0 0-.241.187l-.113.103l-.147.15c-.05.054-.097.11-.142.168l-.1.135l-.05.073l-.06.097c-.05.082-.096.166-.137.253l-.057.126A2.99 2.99 0 0 0 5 8c0 .414.084.809.236 1.168l.057.126l.116.218l.081.132c.059.089.121.175.189.257l.15.17l.151.147c.056.051.114.1.174.147l.142.105c.054.037.109.072.165.106l-.124-.079l.092.06l.094.055c.436.247.94.388 1.477.388a2.99 2.99 0 0 0 1.168-.236l.125-.056l.213-.113l.151-.094l.05-.034a3.011 3.011 0 0 0 .323-.258l-.15.129l.09-.075l.168-.159l.08-.084c.051-.056.1-.114.147-.174l.105-.142l.106-.165c.047-.08.091-.161.131-.245l.057-.126A2.99 2.99 0 0 0 11 8a2.99 2.99 0 0 0-.236-1.168l-.056-.125l-.112-.211l-.096-.155l-.033-.049a3.011 3.011 0 0 0-.258-.322l.129.15l-.075-.09l-.159-.168l-.084-.08a3.015 3.015 0 0 0-.174-.147l-.183-.132l-.124-.079a2.993 2.993 0 0 0-.245-.131l-.126-.057A2.99 2.99 0 0 0 8 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cut(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m5.142 11.074l-1.912.548a2.532 2.532 0 1 1-1.395-4.867l1.947-.559a2.532 2.532 0 0 1 2.555.713l1.53-5.335c.139-.485.6-.897 1.159-1.238c.27-.164.52-.278.779-.32c.814-.132 1.503.558 1.261 1.422L9.574 6.643l4.988-1.43c.864-.242 1.554.447 1.422 1.26c-.042.26-.156.51-.32.78c-.341.56-.753 1.02-1.238 1.16L9.523 9.817a2.53 2.53 0 0 1 .56 2.4l-.56 1.947a2.532 2.532 0 1 1-4.867-1.395zm.33-1.148l.48-1.673a1.52 1.52 0 0 0-1.89-1.083l-1.948.558a1.52 1.52 0 0 0 .837 2.92zm3.773-2.135l-.33 1.148l5.232-1.5c.324-.093 1.182-1.39.694-1.253zM5.63 13.049a1.52 1.52 0 0 0 2.92.837l.559-1.947a1.52 1.52 0 0 0-1.553-1.935l2.537-8.845c.136-.488-1.16.37-1.253.694zm.973.279l.559-1.947a.506.506 0 1 1 .973.279l-.558 1.947a.506.506 0 1 1-.974-.28m-3.93-3.653a.506.506 0 1 1-.28-.973l1.947-.558a.506.506 0 0 1 .28.973z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Database(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 12h12v-1.97c-1.225.582-3.454.97-6 .97s-4.775-.388-6-.97zm-1 0V3c0-1.105 3.134-2 7-2s7 .895 7 2v9c0 1.105-3.134 2-7 2s-7-.895-7-2m1-3h12V7.03c-1.225.582-3.454.97-6 .97s-4.775-.388-6-.97zm0-4.97V6h12V4.03c-1.225.582-3.454.97-6 .97s-4.775-.388-6-.97m10.675-1.483C11.467 2.202 9.795 2 8 2c-1.794 0-3.467.202-4.675.547c-.492.14-.88.298-1.136.453c.256.155.644.312 1.136.453C4.533 3.798 6.205 4 8 4c1.794 0 3.467-.202 4.675-.547c.492-.14.88-.298 1.136-.453c-.256-.155-.644-.312-1.136-.453M2 6c.257.155.833.312 1.325.453C4.533 6.798 6.205 7 8 7c1.794 0 3.467-.202 4.675-.547c.492-.14 1.07-.298 1.327-.453zm0 3c.257.155.833.312 1.325.453C4.533 9.798 6.205 10 8 10c1.794 0 3.467-.202 4.675-.547c.492-.14 1.07-.298 1.327-.453zm0 3c.257.155.833.312 1.325.453C4.533 12.798 6.205 13 8 13c1.794 0 3.467-.202 4.675-.547c.492-.14 1.07-.298 1.327-.453z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DockedBottom(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 1h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V3a2 2 0 0 1 2-2m0 1a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1z"/><path fill="currentColor" d="M3 9.5a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5h-9a.5.5 0 0 1-.5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DockedDetached(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 1H3a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V6h-1v7a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1h6z"/><rect width="6" height="5" x="10" fill="currentColor" rx=".5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DockedLeft(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 1h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V3a2 2 0 0 1 2-2m0 1a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1z"/><path fill="currentColor" d="M3 3.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 .5.5v9a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1-.5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DockedRight(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 1h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V3a2 2 0 0 1 2-2m0 1a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1z"/><path fill="currentColor" d="M9 3.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 .5.5v9a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1-.5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DockedTakeover(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 1h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V3a2 2 0 0 1 2-2m0 1a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1z"/><path fill="currentColor" d="M3 5.5a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-.5.5h-9a.5.5 0 0 1-.5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DockedTop(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 1h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V3a2 2 0 0 1 2-2m0 1a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1z"/><path fill="currentColor" d="M3 3.5a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5h-9a.5.5 0 0 1-.5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Document(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.8 0c.274 0 .537.113.726.312l3.2 3.428c.176.186.274.433.274.689V15a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1V1a1 1 0 0 1 1-1zM14 5h-3.5a.5.5 0 0 1-.5-.5V1H2v14h12zm-8.5 7a.5.5 0 1 1 0-1h5a.5.5 0 1 1 0 1zm0-3a.5.5 0 0 1 0-1h5a.5.5 0 1 1 0 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentEdit(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.505 8.995l6.453-6.44l-1.5-1.5l-6.453 6.44zM12.968.19c.258-.238.657-.26.91 0l1.928 1.929a.642.642 0 0 1 0 .909l-6.78 6.784A.641.641 0 0 1 8.57 10H6.643A.643.643 0 0 1 6 9.357V7.43c0-.17.067-.335.188-.455zM4.5 13a.5.5 0 1 1 0-1h7a.5.5 0 1 1 0 1zm4-12a.5.5 0 0 1 0 1H2v13h12V7.5a.5.5 0 1 1 1 0V15a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Documentation(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 3.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0M9 5v3h1v1H8V6H7V5z"/><path fill="currentColor" d="M13.855 14.147a1.34 1.34 0 0 1-.158-.246A1.998 1.998 0 0 1 13.5 13c0-.414.103-.713.197-.901a1.34 1.34 0 0 1 .158-.246l.003-.005A.5.5 0 0 0 14 11.5V.5a.5.5 0 0 0-.5-.5H3.461l-.083.005a2.957 2.957 0 0 0-1.102.298a2.257 2.257 0 0 0-.88.763C1.148 1.44 1 1.913 1 2.5V13c0 .463.117.843.318 1.145c.2.298.462.491.708.615a2.344 2.344 0 0 0 .94.24H3v-1c-.005 0-.015 0-.029-.002a1.344 1.344 0 0 1-.498-.133a.817.817 0 0 1-.323-.275C2.07 13.47 2 13.287 2 13s.07-.47.15-.59a.817.817 0 0 1 .324-.275A1.344 1.344 0 0 1 3 12h9.658c-.091.27-.158.605-.158 1s.067.73.158 1H8v1h5.5a.5.5 0 0 0 .359-.848zm-.001 0l.002.002zM2.724 1.197c.092-.046.186-.082.276-.11C3 2.918 3.001 11 2.999 11h-.033a1.977 1.977 0 0 0-.283.03a2.344 2.344 0 0 0-.657.21L2 11.254V2.5c0-.413.102-.689.229-.879c.128-.193.304-.328.495-.424M4 11V1h9v10z"/><path fill="currentColor" d="M7 13H4v2.5a.5.5 0 0 0 .854.354l.646-.647l.646.647A.5.5 0 0 0 7 15.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Documents(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.8 0c.274 0 .537.113.726.312l3.2 3.428c.176.186.274.433.274.689V13a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1V1a1 1 0 0 1 1-1zM12 5H8.5a.5.5 0 0 1-.5-.5V1H2v12h10zm-7.5 6a.5.5 0 1 1 0-1h5a.5.5 0 1 1 0 1zm0-3a.5.5 0 0 1 0-1h5a.5.5 0 1 1 0 1zm1 8a.5.5 0 1 1 0-1H14V6.5a.5.5 0 1 1 1 0V15a1 1 0 0 1-1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dot(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="8" cy="8" r="4" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9 9.114l1.85-1.943a.52.52 0 0 1 .77 0c.214.228.214.6 0 .829l-1.95 2.05a1.552 1.552 0 0 1-2.31 0L5.41 8a.617.617 0 0 1 0-.829a.52.52 0 0 1 .77 0L8 9.082V.556C8 .249 8.224 0 8.5 0s.5.249.5.556z"/><path fill="currentColor" d="M16 13.006V10h-1v3.006a.995.995 0 0 1-.994.994H3.01a.995.995 0 0 1-.994-.994V10h-1v3.006c0 1.1.892 1.994 1.994 1.994h10.996c1.1 0 1.994-.893 1.994-1.994"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorAlignCenter(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 4V3h6v1zM3 7V6h10v1zm2 3V9h6v1zm-2 3v-1h10v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorAlignLeft(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 4V3h6v1zm0 3V6h10v1zm0 3V9h6v1zm0 3v-1h10v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorAlignRight(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 4V3h6v1zM4 7V6h10v1zm4 3V9h6v1zm-4 3v-1h10v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorBold(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.193 13H4V3h4.151c1.816 0 2.987.977 2.987 2.495c0 1.074-.797 2.01-1.823 2.176v.055c1.359.132 2.308 1.11 2.308 2.433c0 1.76-1.296 2.841-3.43 2.841M5.788 4.393v2.82h1.635c1.248 0 1.948-.526 1.948-1.455c0-.873-.603-1.365-1.67-1.365zm0 7.214h1.996c1.316 0 2.016-.547 2.016-1.573c0-1.019-.72-1.552-2.092-1.552h-1.92z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorCodeBlock(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.414 8.036L4.89 10.51a.5.5 0 0 1-.707.708L1.354 8.389a.5.5 0 0 1 0-.707l2.828-2.828a.5.5 0 1 1 .707.707zm8.768 2.474l2.475-2.474l-2.475-2.475a.5.5 0 0 1 .707-.707l2.829 2.828a.5.5 0 0 1 0 .707l-2.829 2.829a.5.5 0 1 1-.707-.708M8.559 2.506a.5.5 0 0 1 .981.19L7.441 13.494a.5.5 0 0 1-.981-.19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorComment(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.111 10H12a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h1v1.297zm.46 1L4 14v-3a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorDistributeHorizontal(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 2h2a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2M1.5 1a.5.5 0 0 1 .5.5v13a.5.5 0 1 1-1 0v-13a.5.5 0 0 1 .5-.5m13 0a.5.5 0 0 1 .5.5v13a.5.5 0 1 1-1 0v-13a.5.5 0 0 1 .5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorDistributeVertical(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 5h8a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2M1.5 1h13a.5.5 0 1 1 0 1h-13a.5.5 0 0 1 0-1m0 13h13a.5.5 0 1 1 0 1h-13a.5.5 0 1 1 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorHeading(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 11a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2V5a1 1 0 1 1 0-2h2a1 1 0 1 1 0 2v2h4V5a1 1 0 1 1 0-2h2a1 1 0 0 1 0 2v6a1 1 0 0 1 0 2h-2a1 1 0 0 1 0-2V9H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorItalic(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.16 12H9.5a.5.5 0 1 1 0 1h-4a.5.5 0 1 1 0-1h1.639l1.7-8H7.5a.5.5 0 0 1 0-1h4a.5.5 0 1 1 0 1H9.861z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorItemAlignBottom(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 1h2a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V3a2 2 0 0 1 2-2M1.5 14h13a.5.5 0 1 1 0 1h-13a.5.5 0 1 1 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorItemAlignCenter(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 4H7V1.5a.5.5 0 0 1 1 0zm0 8v2.5a.5.5 0 1 1-1 0V12zM3 5h9a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorItemAlignLeft(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 5h8a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2M1.5 1a.5.5 0 0 1 .5.5v13a.5.5 0 1 1-1 0v-13a.5.5 0 0 1 .5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorItemAlignMiddle(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 7v1H1.5a.5.5 0 0 1 0-1zm8 0h2.5a.5.5 0 1 1 0 1H12zM7 1h2a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V3a2 2 0 0 1 2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorItemAlignRight(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 5h8a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2m11.5-4a.5.5 0 0 1 .5.5v13a.5.5 0 1 1-1 0v-13a.5.5 0 0 1 .5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorItemAlignTop(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 3h2a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2M1.5 1h13a.5.5 0 1 1 0 1h-13a.5.5 0 0 1 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorLink(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.556 5.051a.45.45 0 0 0 .637.637l1.503-1.504c.432-.431 1.278-.382 1.89.23c.612.612.662 1.458.23 1.89L9.519 8.6c-.432.432-1.278.383-1.89-.23a.45.45 0 1 0-.636.637c.914.914 2.33 1.063 3.162.23l2.297-2.297c.833-.833.684-2.248-.23-3.162c-.914-.915-2.33-1.063-3.162-.23zm.888 5.261a.45.45 0 0 0-.637 0l-1.503 1.504c-.432.431-1.278.382-1.89-.23c-.612-.612-.661-1.458-.23-1.89L6.481 7.4c.432-.432 1.278-.383 1.89.23a.45.45 0 0 0 .636-.637c-.914-.914-2.33-1.063-3.162-.23L3.548 9.06c-.833.833-.685 2.248.23 3.162c.914.915 2.33 1.063 3.162.23l1.504-1.503a.45.45 0 0 0 0-.637M7.877 5.76a.39.39 0 0 0 .274-.114l1.503-1.504l-1.503 1.504a.39.39 0 0 1-.274.114m.912 3.183c-.4.003-.843-.172-1.202-.53c.359.358.802.533 1.202.53M12.18 3.82c-.502-.503-1.155-.766-1.773-.76c.618-.006 1.27.257 1.773.76c.898.898 1.034 2.275.23 3.078l-2.297 2.297l2.297-2.297c.804-.803.668-2.18-.23-3.078m-4.062 6.42a.39.39 0 0 1 .284.667L6.898 12.41l1.504-1.503a.39.39 0 0 0-.284-.667m-.926-3.965c.618-.006 1.27.257 1.773.76c-.502-.503-1.155-.766-1.773-.76M5.414 12.15a1.762 1.762 0 0 1-1.042-.522c-.626-.627-.692-1.511-.23-1.974L6.44 7.358L4.142 9.654c-.462.463-.396 1.348.23 1.974c.311.311.687.484 1.042.522"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorOrderedList(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 8v1h1v1H1V9h1V8H1V7h3v1zm1 4v2H1v-1h1v-1H1v-1h3zM3 5h1v1H1V5h1V4H1V3h2zm2.5-1h8a.5.5 0 1 1 0 1h-8a.5.5 0 0 1 0-1m0 4h8a.5.5 0 1 1 0 1h-8a.5.5 0 0 1 0-1m0 4h8a.5.5 0 1 1 0 1h-8a.5.5 0 1 1 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorPositionBottomLeft(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 8h5a1 1 0 0 1 1 1v5H3a1 1 0 0 1-1-1zm1-7h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V3a2 2 0 0 1 2-2m0 1a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorPositionBottomRight(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 8h5v5a1 1 0 0 1-1 1H8V9a1 1 0 0 1 1-1M3 1h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V3a2 2 0 0 1 2-2m0 1a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorPositionTopLeft(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 2h5v5a1 1 0 0 1-1 1H2V3a1 1 0 0 1 1-1m0-1h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V3a2 2 0 0 1 2-2m0 1a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorPositionTopRight(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 2h5a1 1 0 0 1 1 1v5H9a1 1 0 0 1-1-1zM3 1h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V3a2 2 0 0 1 2-2m0 1a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorRedo(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 6H5C3.057 6 2 7.057 2 9s1.057 3 3 3h1.5v-1H5c-1.39 0-2-.61-2-2c0-1.39.61-2 2-2h5v3l3.5-3.5L10 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorStrike(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.023 10h1.274c.006.08.01.164.01.25a2.557 2.557 0 0 1-.883 1.949c-.284.25-.627.446-1.03.588A4.087 4.087 0 0 1 8.028 13a4.616 4.616 0 0 1-3.382-1.426c-.193-.259-.193-.5 0-.724c.193-.223.438-.266.735-.13c.343.363.748.655 1.213.876c.466.22.949.33 1.449.33c.637 0 1.132-.144 1.485-.433c.353-.29.53-.67.53-1.14a1.72 1.72 0 0 0-.034-.353M5.586 7a2.49 2.49 0 0 1-.294-.507a2.316 2.316 0 0 1-.177-.934c0-.363.076-.701.228-1.015c.152-.314.363-.586.633-.816c.27-.23.588-.41.955-.537A3.683 3.683 0 0 1 8.145 3c.578 0 1.112.11 1.603.33c.49.221.907.508 1.25.861c.16.282.16.512 0 .692c-.16.18-.38.214-.662.102a3.438 3.438 0 0 0-.978-.669a2.914 2.914 0 0 0-1.213-.242c-.54 0-.973.125-1.302.375c-.328.25-.492.595-.492 1.036c0 .236.046.434.14.596c.092.162.217.304.374.426c.157.123.329.23.515.324c.119.06.24.116.362.169zM2.5 8h11a.5.5 0 1 1 0 1h-11a.5.5 0 0 1 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorTable(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 3v2H2V3v10h12zM2 2h12c.552 0 1 .413 1 .923v10.154c0 .51-.448.923-1 .923H2c-.552 0-1-.413-1-.923V2.923C1 2.413 1.448 2 2 2m0 5h12v1H2zm0 3h12v1H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorUnderline(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3.536V7.8c0 1.628 1.224 2.6 3 2.6c1.783 0 3-.972 3-2.6V3.536c0-.357.167-.536.5-.536c.333 0 .5.179.5.536v4.318c0 2.093-1.665 3.546-4 3.546S4 9.893 4 7.8V3.536C4 3.179 4.167 3 4.5 3c.333 0 .5.179.5.536M2.5 13h11a.5.5 0 1 1 0 1h-11a.5.5 0 1 1 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorUndo(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 6h5c1.943 0 3 1.057 3 3s-1.057 3-3 3H9v-1h1.5c1.39 0 2-.61 2-2c0-1.39-.61-2-2-2h-5v3L2 6.5L5.5 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorUnorderedList(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 4h8a.5.5 0 1 1 0 1h-8a.5.5 0 0 1 0-1m0 4h8a.5.5 0 1 1 0 1h-8a.5.5 0 0 1 0-1m0 4h8a.5.5 0 1 1 0 1h-8a.5.5 0 1 1 0-1m-3-7a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0 4a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0 4a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Email(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.95 3.684L8.637 8.912a1 1 0 0 1-1.276 0l-6.31-5.228A.999.999 0 0 0 1 4v8a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V4a.999.999 0 0 0-.05-.316M2 2h12a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2m-.21 1l5.576 4.603a1 1 0 0 0 1.27.003L14.268 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Empty(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(``),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eql(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.862 14.18a.482.482 0 0 1 .1.521a.5.5 0 0 1-.447.299h-3.149c-.07 0-.136 0-.2-.003a.664.664 0 0 1-.414-.14c-.35-.31-.684-.637-1.001-.981l-1.772-1.738a5.941 5.941 0 0 1-3.802.14l-.045-.012a5.89 5.89 0 0 1-2.682-1.712A5.715 5.715 0 0 1 1.08 7.72a5.66 5.66 0 0 1 .348-3.118A5.78 5.78 0 0 1 3.39 2.124a5.979 5.979 0 0 1 6.136-.505a5.836 5.836 0 0 1 2.356 2.123a5.67 5.67 0 0 1 .873 3.017l.001.003c0 .128-.051.25-.143.34a.495.495 0 0 1-.694 0a.476.476 0 0 1-.143-.34c0-.95-.288-1.878-.826-2.668a4.88 4.88 0 0 0-2.198-1.769a4.99 4.99 0 0 0-2.83-.273a4.93 4.93 0 0 0-2.509 1.314a4.774 4.774 0 0 0-1.34 2.46A4.715 4.715 0 0 0 2.352 8.6a4.82 4.82 0 0 0 1.804 2.155c.805.528 1.752.81 2.721.81a4.899 4.899 0 0 0 1.296-.194l.032-.009l.717-.211a.506.506 0 0 1 .483.111l1.11 1.026l1.788 1.752h2.03l-3.657-3.583a.475.475 0 0 1 .014-.666a.495.495 0 0 1 .679-.013z"/><path fill="currentColor" d="M7.097 3.468L9.679 4.93a.433.433 0 0 1 .218.37v2.924a.422.422 0 0 1-.218.37l-2.582 1.461a.438.438 0 0 1-.437 0L4.077 8.594a.433.433 0 0 1-.218-.37V5.3a.422.422 0 0 1 .218-.37L6.66 3.467a.446.446 0 0 1 .437 0zm1.845 4.27V5.784a.328.328 0 0 0-.17-.287L7.047 4.52a.342.342 0 0 0-.338 0l-1.726.977a.336.336 0 0 0-.168.287v1.953a.326.326 0 0 0 .17.287l1.724.977a.342.342 0 0 0 .338 0l1.726-.978a.334.334 0 0 0 .169-.287z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eraser(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.648 9.937l7.29-7.288a2.21 2.21 0 0 1 3.124 0l2.29 2.288a2.21 2.21 0 0 1 0 3.126L10.413 13H12.5a.5.5 0 0 1 0 1H4.501a2.21 2.21 0 0 1-1.563-.647l.707-.707c.227.226.535.354.856.354h4.005a1.21 1.21 0 0 0 .848-.354l1.292-1.293l-4-4l-3.29 3.291a1.21 1.21 0 0 0 0 1.712l.29.29l-.708.707l-.29-.29a2.21 2.21 0 0 1 0-3.126M8 6h6.89a1.208 1.208 0 0 0-.246-.356L14 5H9zm2-2h3l-.645-.644a1.21 1.21 0 0 0-1.71 0zm4.89 3H7.708l1 1H14l.644-.644A1.22 1.22 0 0 0 14.891 7zM9.708 9l1.646 1.646L13 9z"/><path fill="currentColor" d="M14 11.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0m1.5-.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1m-1 2a.5.5 0 0 0 0 1h1a.5.5 0 0 0 0-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Exit(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.535 12.493a.47.47 0 0 1 .468.468v2.564a.473.473 0 0 1-.466.475H3V0h9.595a.45.45 0 0 1 .398.463v2.565a.469.469 0 0 1-.468.467h-.065a.468.468 0 0 1-.467-.467V1H4v14h8.01l-.007-2.04c0-.257.21-.467.467-.467zm-1.096-7.59l2.121 2.122a1.5 1.5 0 0 1 0 2.121l-2.12 2.122a.5.5 0 1 1-.708-.708l2.121-2.12a.5.5 0 0 0 0-.708l-2.121-2.121a.5.5 0 0 1 .707-.707"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Expand(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m4.354 12.354l8-8a.5.5 0 0 0-.708-.708l-8 8a.5.5 0 0 0 .708.708M1 10.5a.5.5 0 1 1 1 0v3a.5.5 0 0 0 .5.5h3a.5.5 0 1 1 0 1h-3A1.5 1.5 0 0 1 1 13.5zm14-5a.5.5 0 1 1-1 0v-3a.5.5 0 0 0-.5-.5h-3a.5.5 0 1 1 0-1h3A1.5 1.5 0 0 1 15 2.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExpandMini(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.707 10L10 6.707A.5.5 0 0 0 9.293 6L6 9.293a.5.5 0 1 0 .707.707M4 9.5a.5.5 0 0 1 1 0v1a.5.5 0 0 0 .5.5h1a.5.5 0 1 1 0 1h-1A1.5 1.5 0 0 1 4 10.5zm8-3a.5.5 0 1 1-1 0v-1a.5.5 0 0 0-.5-.5h-1a.5.5 0 0 1 0-1h1A1.5 1.5 0 0 1 12 5.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Export(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.505 1c.422-.003.844.17 1.166.516l1.95 2.05c.213.228.213.6 0 .828a.52.52 0 0 1-.771 0L9 2.451v7.993c0 .307-.224.556-.5.556s-.5-.249-.5-.556v-7.96l-1.82 1.91a.52.52 0 0 1-.77 0a.617.617 0 0 1 0-.829l1.95-2.05A1.575 1.575 0 0 1 8.5 1zM4.18 7c-.473 0-.88.294-.972.703l-1.189 5.25a.776.776 0 0 0-.019.172c0 .483.444.875.99.875h11.02c.065 0 .13-.006.194-.017c.537-.095.885-.556.778-1.03l-1.19-5.25C13.7 7.294 13.293 7 12.822 7zM6 6v1h5V6h1.825c.946 0 1.76.606 1.946 1.447l1.19 5.4c.215.975-.482 1.923-1.556 2.118a2.18 2.18 0 0 1-.39.035H2.985C1.888 15 1 14.194 1 13.2c0-.119.013-.237.039-.353l1.19-5.4C2.414 6.606 3.229 6 4.174 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.98 7.873c.013.03.02.064.02.098v.06a.24.24 0 0 1-.02.097C15.952 8.188 13.291 14 8 14S.047 8.188.02 8.128A.24.24 0 0 1 0 8.03v-.059c0-.034.007-.068.02-.098C.048 7.813 2.709 2 8 2s7.953 5.813 7.98 5.873m-1.37-.424a12.097 12.097 0 0 0-1.385-1.862C11.739 3.956 9.999 3 8 3c-2 0-3.74.956-5.225 2.587a12.098 12.098 0 0 0-1.701 2.414a12.095 12.095 0 0 0 1.7 2.413C4.26 12.043 6.002 13 8 13s3.74-.956 5.225-2.587A12.097 12.097 0 0 0 14.926 8c-.08-.15-.189-.343-.315-.551M8 4.75A3.253 3.253 0 0 1 11.25 8A3.254 3.254 0 0 1 8 11.25A3.253 3.253 0 0 1 4.75 8A3.252 3.252 0 0 1 8 4.75m0 1C6.76 5.75 5.75 6.76 5.75 8S6.76 10.25 8 10.25S10.25 9.24 10.25 8S9.24 5.75 8 5.75m0 1.5a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeClosed(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.318 13.47l.776-.776A6.04 6.04 0 0 0 8 13c1.999 0 3.74-.956 5.225-2.587A12.097 12.097 0 0 0 14.926 8a12.097 12.097 0 0 0-1.701-2.413l-.011-.012l.707-.708c1.359 1.476 2.045 2.976 2.058 3.006c.014.03.021.064.021.098v.06a.24.24 0 0 1-.02.097C15.952 8.188 13.291 14 8 14a7.03 7.03 0 0 1-2.682-.53M2.04 11.092C.707 9.629.034 8.158.02 8.128A.24.24 0 0 1 0 8.03v-.059c0-.034.007-.068.02-.098C.048 7.813 2.709 2 8 2c.962 0 1.837.192 2.625.507l-.78.78A6.039 6.039 0 0 0 8 3c-2 0-3.74.956-5.225 2.587a12.098 12.098 0 0 0-1.701 2.414a12.11 12.11 0 0 0 1.674 2.383zM8.362 4.77L7.255 5.877a2.262 2.262 0 0 0-1.378 1.378L4.77 8.362A3.252 3.252 0 0 1 8.362 4.77m2.86 2.797a3.254 3.254 0 0 1-3.654 3.654l1.06-1.06a2.262 2.262 0 0 0 1.533-1.533zm-9.368 7.287a.5.5 0 0 1-.708-.708l13-13a.5.5 0 0 1 .708.708z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceHappy(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16A8 8 0 1 1 8 0a8 8 0 0 1 0 16m0-1.067A6.933 6.933 0 1 0 8 1.067a6.933 6.933 0 0 0 0 13.866M5.333 6.4a1.067 1.067 0 1 1 0-2.133a1.067 1.067 0 0 1 0 2.133m5.334 0a1.067 1.067 0 1 1 0-2.133a1.067 1.067 0 0 1 0 2.133M2.739 8.802a.533.533 0 0 1 .922-.537C4.815 10.245 6.249 11.2 8 11.2c1.75 0 3.185-.956 4.34-2.935a.533.533 0 0 1 .92.537c-1.333 2.287-3.1 3.465-5.26 3.465c-2.16 0-3.927-1.178-5.26-3.465"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceNeutral(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><circle cx="5" cy="5" r="1"/><circle cx="10" cy="5" r="1"/><path fill-rule="nonzero" d="M7.5 14a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13m0 1a7.5 7.5 0 1 1 0-15a7.5 7.5 0 0 1 0 15"/><path fill-rule="nonzero" d="M3 10h9a.5.5 0 1 0 0-1H3a.5.5 0 0 0 0 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceSad(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16A8 8 0 1 1 8 0a8 8 0 0 1 0 16m0-1.067A6.933 6.933 0 1 0 8 1.067a6.933 6.933 0 0 0 0 13.866M5.333 6.4a1.067 1.067 0 1 1 0-2.133a1.067 1.067 0 0 1 0 2.133m5.334 0a1.067 1.067 0 1 1 0-2.133a1.067 1.067 0 0 1 0 2.133m2.487 3.868a.533.533 0 1 1-.708.797C10.936 9.725 9.458 9.067 8 9.067c-1.458 0-2.937.657-4.446 1.998a.533.533 0 1 1-.708-.797C4.536 8.765 6.258 8 8 8s3.463.765 5.154 2.268"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filter(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.999 15.999a8 8 0 1 1 0-16a8 8 0 0 1 0 16M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14M3.5 5h9a.5.5 0 1 1 0 1h-9a.5.5 0 0 1 0-1m2 3h5a.5.5 0 1 1 0 1h-5a.5.5 0 0 1 0-1m2 3h1a.5.5 0 1 1 0 1h-1a.5.5 0 1 1 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.686 8.464c1.547-.619 3.08-.619 4.628 0A.5.5 0 0 0 13 8V2a.5.5 0 0 0-.276-.447C11.259.82 9.458.82 7.342 1.526c-1.884.628-3.417.628-4.618.027A.5.5 0 0 0 2 2v12.5a.5.5 0 1 0 1 0V8.553c1.411.627 2.983.592 4.686-.089M3 2.741c1.322.42 2.878.327 4.658-.267C9.4 1.894 10.843 1.85 12 2.322v4.975c-1.56-.464-3.128-.384-4.686.239c-1.54.616-2.892.616-4.09.017A.498.498 0 0 0 3 7.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fold(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.14.192L7.53 2.49a.67.67 0 0 0 .942 0L10.86.192a.677.677 0 0 1 .944 0a.65.65 0 0 1 0 .93l-2.388 2.3a2.02 2.02 0 0 1-2.832 0l-2.388-2.3a.65.65 0 0 1 0-.93a.677.677 0 0 1 .944 0m0 15.616l2.39-2.298a.67.67 0 0 1 .942 0l2.389 2.298c.26.256.685.256.944 0a.65.65 0 0 0 0-.93l-2.388-2.3a2.02 2.02 0 0 0-2.832 0l-2.388 2.3a.65.65 0 0 0 0 .93c.26.256.683.256.944 0zM16 6H0v4h16zM1 9V7h14v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderCheck(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 2H1v11h6.1c.07.348.177.682.316 1H1a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1h5.25a.75.75 0 0 1 .75.75v1.125c0 .069.056.125.125.125H13a1 1 0 0 1 1 1v3.416a4.962 4.962 0 0 0-1-.316V4H7.125A1.125 1.125 0 0 1 6 2.875zm10 10a4 4 0 1 1-8 0a4 4 0 0 1 8 0m-1.646-1.354a.5.5 0 0 1 0 .708l-2.5 2.5a.5.5 0 0 1-.708 0l-1-1a.5.5 0 0 1 .708-.708l.646.647l2.146-2.147a.5.5 0 0 1 .708 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderClosed(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 2H2v11h12V4H8.125A1.125 1.125 0 0 1 7 2.875zm.25-1a.75.75 0 0 1 .75.75v1.125c0 .069.056.125.125.125H14a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderExclamation(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m1 9.5l.826-3.717A1 1 0 0 1 2.802 5H13V4H7.125A1.125 1.125 0 0 1 6 2.875V2H1zm.247 3.5H7.1c.07.348.177.682.316 1H1a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1h5.25a.75.75 0 0 1 .75.75v1.125c0 .069.056.125.125.125H13a1 1 0 0 1 1 1v1h.753a1 1 0 0 1 .977 1.217l-.447 2.011a5.015 5.015 0 0 0-.887-.618L14.753 6H2.803zM16 12a4 4 0 1 1-8 0a4 4 0 0 1 8 0m-4 .5a.577.577 0 0 1-.57-.495l-.29-2.015a.867.867 0 1 1 1.718 0l-.288 2.015a.577.577 0 0 1-.57.495m0 2.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpen(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1 9.5l.826-3.717A1 1 0 0 1 2.802 5H13V4H7.125A1.125 1.125 0 0 1 6 2.875V2H1zm.247 3.5h11.95l1.556-7H2.803zM13 14H1a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1h5.25a.75.75 0 0 1 .75.75v1.125c0 .069.056.125.125.125H13a1 1 0 0 1 1 1v1h.753a1 1 0 0 1 .977 1.217l-1.556 7a1 1 0 0 1-.976.783z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FullScreen(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13 3v4h-1V4H9V3zM3 3h4v1H4v3H3zm10 10H9v-1h3V9h1zM3 13V9h1v3h3v1zM0 1.994C0 .893.895 0 1.994 0h12.012C15.107 0 16 .895 16 1.994v12.012A1.995 1.995 0 0 1 14.006 16H1.994A1.995 1.995 0 0 1 0 14.006zm1 0v12.012c0 .548.446.994.994.994h12.012a.995.995 0 0 0 .994-.994V1.994A.995.995 0 0 0 14.006 1H1.994A.995.995 0 0 0 1 1.994"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FullScreenExit(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 7V3h1v3h3v1zM7 7H3V6h3V3h1zm2 2h4v1h-3v3H9zM7 9v4H6v-3H3V9zM0 1.994C0 .893.895 0 1.994 0h12.012C15.107 0 16 .895 16 1.994v12.012A1.995 1.995 0 0 1 14.006 16H1.994A1.995 1.995 0 0 1 0 14.006zm1 0v12.012c0 .548.446.994.994.994h12.012a.995.995 0 0 0 .994-.994V1.994A.995.995 0 0 0 14.006 1H1.994A.995.995 0 0 0 1 1.994"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Function(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 2.226v2.218c-.359-.143-.845-.218-1.315-.218c-1.059 0-1.631.519-1.802 1.565l-.168.937h2.798v2.159H9.41l-.313 1.674C8.696 12.987 7.261 14 4.785 14c-.718 0-1.35-.092-1.785-.251v-2.243c.418.176.905.268 1.383.268c1.008 0 1.546-.435 1.725-1.523l.24-1.364H3.787V6.728h2.812l.288-1.264C7.286 3.071 8.662 2 11.352 2c.598 0 1.306.1 1.648.226"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gear(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.164 10.329L1.87 8L.163 5.67c.18-.601.43-1.19.758-1.757a8.197 8.197 0 0 1 1.142-1.535l2.872.313L6.099.05a8.166 8.166 0 0 1 3.8-.003l1.166 2.644l2.872-.313a8.166 8.166 0 0 1 1.899 3.293L14.13 8l1.706 2.33c-.18.601-.43 1.19-.758 1.757a8.197 8.197 0 0 1-1.142 1.535l-2.872-.313L9.9 15.95a8.166 8.166 0 0 1-3.8.003l-1.166-2.644l-2.872.313a8.166 8.166 0 0 1-1.899-3.293m4.663 1.986a1 1 0 0 1 1.023.591l.957 2.17c.79.134 1.597.132 2.387-.001l.956-2.169a1 1 0 0 1 1.023-.59l2.358.256a7.23 7.23 0 0 0 1.194-2.068l-1.401-1.913a1 1 0 0 1 0-1.182l1.4-1.912a7.165 7.165 0 0 0-1.192-2.069l-2.359.257a1 1 0 0 1-1.023-.591L9.193.924a7.165 7.165 0 0 0-2.387.001L5.85 3.094a1 1 0 0 1-1.023.59l-2.358-.256a7.23 7.23 0 0 0-1.194 2.068l1.401 1.913a1 1 0 0 1 0 1.182l-1.4 1.912c.28.751.681 1.45 1.192 2.069zM8 11a3 3 0 1 1 0-6a3 3 0 0 1 0 6m0-1a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Glasses(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.035 9A3.5 3.5 0 0 1 15 7.05V4.5c0-1.072-.648-1.72-2.098-2.01a.5.5 0 0 1 .196-.98C14.981 1.886 16 2.905 16 4.5v4.25c0 .072-.015.14-.043.202A3.5 3.5 0 1 1 9.035 10h-2.07A3.5 3.5 0 1 1 .043 8.952A.498.498 0 0 1 0 8.75V4.5c0-1.595 1.019-2.614 2.902-2.99a.5.5 0 0 1 .196.98C1.648 2.78 1 3.428 1 4.5v2.55A3.5 3.5 0 0 1 6.965 9zM3.5 12a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5m9 0a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globe(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.019 8a6.462 6.462 0 0 0 1.003 3h2.382a14.469 14.469 0 0 1-.396-3zm0-1h2.989c.033-1.078.172-2.094.396-3H2.022a6.462 6.462 0 0 0-1.003 3M13.98 8h-2.989a14.469 14.469 0 0 1-.396 3h2.382a6.462 6.462 0 0 0 1.003-3m0-1a6.462 6.462 0 0 0-1.003-3h-2.382c.224.906.363 1.922.396 3zM5.008 8c.037 1.107.195 2.127.429 3h4.126c.234-.873.392-1.893.429-3zm0-1h4.984a13.422 13.422 0 0 0-.429-3H5.437a13.422 13.422 0 0 0-.429 3M.016 8H0V7h.016a7.5 7.5 0 0 1 14.968 0H15v1h-.016A7.5 7.5 0 0 1 .016 8m2.794 4a6.501 6.501 0 0 0 2.717 1.695A7.315 7.315 0 0 1 4.7 12zm9.38 0H10.3c-.23.657-.51 1.23-.827 1.695A6.501 6.501 0 0 0 12.19 12m-6.428 0c.484 1.24 1.132 2 1.738 2c.606 0 1.254-.76 1.738-2zM2.81 3H4.7c.23-.657.51-1.23.827-1.695A6.501 6.501 0 0 0 2.81 3m9.38 0a6.501 6.501 0 0 0-2.717-1.695c.317.465.597 1.038.827 1.695zM5.762 3h3.476C8.754 1.76 8.106 1 7.5 1c-.606 0-1.254.76-1.738 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grab(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 6c.276 0 .5.232.5.5c0 .276-.229.5-.5.5h-11a.505.505 0 0 1-.5-.5c0-.276.229-.5.5-.5zm0 3c.276 0 .5.232.5.5c0 .276-.229.5-.5.5h-11a.505.505 0 0 1-.5-.5c0-.276.229-.5.5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GrabHorizontal(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 2.5c0-.276.232-.5.5-.5c.276 0 .5.229.5.5v11c0 .276-.232.5-.5.5a.503.503 0 0 1-.5-.5zm3 0c0-.276.232-.5.5-.5c.276 0 .5.229.5.5v11c0 .276-.232.5-.5.5a.503.503 0 0 1-.5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grid(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 5V1h4v4zm3-1V2H2v2zm2 1V1h4v4zm3-1V2H7v2zm2 1V1h4v4zm1-1h2V2h-2zM1 10V6h4v4zm3-1V7H2v2zm2 1V6h4v4zm3-1V7H7v2zm2 1V6h4v4zm3-1V7h-2v2zM1 15v-4h4v4zm1-1h2v-2H2zm4 1v-4h4v4zm1-1h2v-2H7zm4 1v-4h4v4zm1-1h2v-2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.402 3.098a3.75 3.75 0 0 0-5.304 5.304l5.558 5.27L8 14l5.892-5.588a3.75 3.75 0 1 0-5.294-5.313L8 3.697zM2.796 7.685a2.747 2.747 0 0 1 .01-3.88a2.75 2.75 0 0 1 3.889 0L8 5.111l1.305-1.306a2.75 2.75 0 1 1 3.89 3.89L8 12.62z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heatmap(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 9a3 3 0 1 1 0 6a3 3 0 0 1 0-6m0 1a2 2 0 1 0 0 4a2 2 0 0 0 0-4M4 2a2 2 0 1 1 0 4a2 2 0 0 1 0-4m0 1a1 1 0 1 0 0 2a1 1 0 0 0 0-2m0 5a4 4 0 1 1 0 8a4 4 0 0 1 0-8m0 1a3 3 0 1 0 0 6a3 3 0 0 0 0-6m8-9a4 4 0 1 1 0 8a4 4 0 0 1 0-8m0 1a3 3 0 1 0 0 6a3 3 0 0 0 0-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Help(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m13.6 12.186l-1.357-1.358c-.025-.025-.058-.034-.084-.056c.53-.794.84-1.746.84-2.773a4.977 4.977 0 0 0-.84-2.772c.026-.02.059-.03.084-.056L13.6 3.813a6.96 6.96 0 0 1 0 8.373M8 15a6.956 6.956 0 0 1-4.186-1.4l1.358-1.358c.025-.025.034-.057.055-.084C6.02 12.688 6.974 13 8 13a4.978 4.978 0 0 0 2.773-.84c.02.026.03.058.056.083l1.357 1.358A6.956 6.956 0 0 1 8 15m-5.601-2.813a6.963 6.963 0 0 1 0-8.373l1.359 1.358c.024.025.057.035.084.056A4.97 4.97 0 0 0 3 8c0 1.027.31 1.98.842 2.773c-.027.022-.06.031-.084.056zm5.6-.187A4 4 0 1 1 8 4a4 4 0 0 1 0 8M8 1c1.573 0 3.019.525 4.187 1.4L10.83 3.758c-.025.025-.035.057-.056.084A4.979 4.979 0 0 0 8 3a4.979 4.979 0 0 0-2.773.842c-.021-.027-.03-.059-.055-.084L3.814 2.4A6.957 6.957 0 0 1 8 1m0-1a8.001 8.001 0 1 0 .003 16.002A8.001 8.001 0 0 0 8 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.13 1.229l5.5 4.47a1 1 0 0 1 .37.777V14a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1V6.476a1 1 0 0 1 .37-.776l5.5-4.471a1 1 0 0 1 1.26 0M13 6.476L7.5 2.005L2 6.475V14h11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IinCircle(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 11.508L7.468 8H6.25V7h2.401l.03 3.508H9.8v1zm-.25-6.202a.83.83 0 0 1 .207-.577c.137-.153.334-.229.59-.229c.256 0 .454.076.594.23c.14.152.209.345.209.576c0 .228-.07.417-.21.568c-.14.15-.337.226-.593.226c-.256 0-.453-.075-.59-.226a.81.81 0 0 1-.207-.568M8 13A5 5 0 1 0 8 3a5 5 0 0 0 0 10m0 1A6 6 0 1 1 8 2a6 6 0 0 1 0 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Image(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 5a2 2 0 1 1-4 0a2 2 0 0 1 4 0m9-4a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H1a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1zm-3.448 6.134l-3.76 2.769a.5.5 0 0 1-.436.077l-.087-.034l-1.713-.87L1 11.8V14h14V9.751zM15 2H1v8.635l4.28-2.558a.5.5 0 0 1 .389-.054l.094.037l1.684.855l3.813-2.807a.5.5 0 0 1 .52-.045l.079.05L15 8.495z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Import(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9 10.114l1.85-1.943a.52.52 0 0 1 .77 0c.214.228.214.6 0 .829l-1.95 2.05a1.552 1.552 0 0 1-2.31 0L5.41 9a.617.617 0 0 1 0-.829a.52.52 0 0 1 .77 0L8 10.082V1.556C8 1.249 8.224 1 8.5 1s.5.249.5.556zM4.18 6a.993.993 0 0 0-.972.804l-1.189 6A.995.995 0 0 0 2.991 14h11.018a1 1 0 0 0 .972-1.196l-1.19-6a.993.993 0 0 0-.97-.804zM6 5v1h5V5h1.825c.946 0 1.76.673 1.946 1.608l1.19 6A2 2 0 0 1 14.016 15H2.984a1.992 1.992 0 0 1-1.945-2.392l1.19-6C2.414 5.673 3.229 5 4.174 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IndexClose(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2H2v11h6v1H1V1h12v6h-1zM5 5h5.999V4H5zM3 5V4h1v1zm2 3V7h3v1zM3 8V7h1v1zm2 3v-1h2v1zm5.5-1.207L9.086 8.379l-.707.707L9.793 10.5l-1.414 1.414l.707.707l1.414-1.414l1.414 1.414l.707-.707l-1.414-1.414l1.414-1.414l-.707-.707zM3 11v-1h1v1zm7.5-5a4.5 4.5 0 1 1 0 9a4.5 4.5 0 0 1 0-9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IndexEdit(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2H2v11h4v1H1V1h12v5h-1zM5 5h5.999V4H5zM3 5V4h1v1zm2 3V7h6v1zM3 8V7h1v1zm2 3v-1h3v1zm-2 0v-1h1v1zm4.502 1.41L12.913 7L15 9.087l-5.41 5.41L7 15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IndexFlush(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.516 9H10.5a.5.5 0 0 1 0-1h4.016L13.11 5.948c-.171-.252-.137-.62.079-.821c.217-.2.531-.159.703.092l2 2.916a.648.648 0 0 1 .108.397a.643.643 0 0 1-.108.332l-2 2.918A.478.478 0 0 1 13.5 12a.457.457 0 0 1-.312-.127a.653.653 0 0 1-.079-.82zM3 15H1a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1h7.8c.274 0 .537.113.726.312l2.2 2.428c.176.186.274.433.274.689V7h-1V5H8.5a.5.5 0 0 1-.5-.5V2H3v12h8v-4h1v4a1 1 0 0 1-1 1zm-1-1V2H1v12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IndexMapping(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 8H4.915a1.5 1.5 0 1 1 0-1H8V2.5A1.5 1.5 0 0 1 9.5 1h2.585a1.5 1.5 0 1 1 0 1H9.5a.5.5 0 0 0-.5.5v10a.5.5 0 0 0 .5.5h2.585a1.5 1.5 0 1 1 0 1H9.5A1.5 1.5 0 0 1 8 12.5zM3.5 3a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3m0 12a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3m10-6a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IndexOpen(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2H2v11h6v1H1V1h12v6h-1zM5 5h5.999V4H5zM3 5V4h1v1zm2 3V7h3v1zM3 8V7h1v1zm2 3v-1h2v1zm5-1H8v1h2v2h1v-2h2v-1h-2V8h-1zm-7 1v-1h1v1zm7.5-5a4.5 4.5 0 1 1 0 9a4.5 4.5 0 0 1 0-9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IndexRuntime(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2H2v11h6v1H1V1h12v6.839l-1-.707z"/><path fill="currentColor" d="M8 11v-1H5v1zm0-3.055c0-.342.082-.664.23-.945H5v1h3zM5 5h5.999V4H5zM3 4v1h1V4zm0 3v1h1V7zm0 3v1h1v-1zm11.607.2l-4.32-3.055C9.727 6.75 9 7.204 9 7.945v6.108c0 .74.726 1.196 1.287.8l4.32-3.055c.524-.37.524-1.228 0-1.598"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IndexSettings(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 5h5.999V4H5zM3 5h1V4H3zm0 3h1V7H3zm6.022-1l-.15.333l-.737-.078l-.467-.05l-.33.342a5.13 5.13 0 0 0-.39.453H5V7zm-3.005 3L6 10.056l.306.411l.399.533H5v-1zM3 11h1v-1H3z"/><path fill="currentColor" d="m13 7.05l-.162-.359l-.2-.447l-.47-.11A5.019 5.019 0 0 0 12 6.098V2H2v11h4.36c.157.354.355.69.59 1H1V1h12z"/><path fill="currentColor" d="M11.004 7c.322 0 .646.036.966.109l.595 1.293l1.465-.152c.457.462.786 1.016.969 1.61l-.87 1.14l.871 1.141a3.94 3.94 0 0 1-.387.859a4.058 4.058 0 0 1-.583.75l-1.465-.152l-.594 1.292a4.37 4.37 0 0 1-1.941.001l-.594-1.293l-1.466.152a3.954 3.954 0 0 1-.969-1.61l.87-1.14L7 9.86a3.947 3.947 0 0 1 .97-1.61l1.466.152l.593-1.292a4.37 4.37 0 0 1 .975-.11M11 12a1 1 0 1 0 .002-1.998A1 1 0 0 0 11 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InputOutput(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10 13.999v-2.99h1V15H0V1h11v3.999h-1V2.001H1v11.998z"/><path d="M4.5 10V9H11v1zm8.5-.5l3-3l-3-3z"/><path d="M5.5 6.5v6l-3-3zM7 7V6h6.5v1z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inspect(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.363 14.658a.5.5 0 1 1-.713.7l-2.97-3.023a.5.5 0 0 1 .001-.7A3.9 3.9 0 1 0 8.9 12.8a.5.5 0 1 1 0 .999a4.9 4.9 0 1 1 3.821-1.833zM3.094 13a.5.5 0 1 1 0 1H2.5A2.5 2.5 0 0 1 0 11.5v-9A2.5 2.5 0 0 1 2.5 0h9A2.5 2.5 0 0 1 14 2.5v.599a.5.5 0 1 1-1 0V2.5A1.5 1.5 0 0 0 11.5 1h-9A1.5 1.5 0 0 0 1 2.5v9A1.5 1.5 0 0 0 2.5 13zM2.5 3a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m2 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m2 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-4 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m2 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-2 1a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1m0 3a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m6-6a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m2 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-8 8a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IntegrationGeneral(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.095 2.903a.525.525 0 0 1 .472-.278l1.05.016a.525.525 0 0 1 .462.292l.42.844l.742-.063a.525.525 0 0 1 .486.239l.53.822a.525.525 0 0 1-.015.59l-.39.546l.415.548c.129.17.142.402.032.586l-.5.838a.525.525 0 0 1-.482.255l-.785-.048l-.483.989a.525.525 0 0 1-.472.294H5.495a.525.525 0 0 1-.471-.294L4.57 8.15l-.84.076a.525.525 0 0 1-.493-.246l-.531-.855a.525.525 0 0 1 .03-.597l.48-.625l-.386-.539a.525.525 0 0 1-.014-.59l.53-.822a.525.525 0 0 1 .485-.239l.797.068zm.598.474l-.472.888a.525.525 0 0 1-.508.277l-.81-.069l-.374.58l.394.55a.525.525 0 0 1-.01.625l-.488.634l.376.606l.856-.078a.525.525 0 0 1 .52.292l.459.941h.8l.488-.997a.525.525 0 0 1 .503-.293l.798.048l.352-.59l-.424-.56a.525.525 0 0 1-.008-.622l.399-.556l-.374-.58l-.76.064a.525.525 0 0 1-.515-.29l-.427-.858z" clip-rule="evenodd"/><path fill="currentColor" d="M10.282 7.4a4.5 4.5 0 0 1-8.666-.376l.851-.168L.847 5.18l-.848 2.167l.897-.176a5.25 5.25 0 0 0 10.125.461zm.836-2.54A5.25 5.25 0 0 0 .993 4.4l.713.232a4.5 4.5 0 0 1 8.666.375l-.851.169l1.62 1.672l.858-2.164z"/><path fill="currentColor" fill-rule="evenodd" d="M6 5.625a.375.375 0 1 0 0 .75a.375.375 0 0 0 0-.75M4.875 6a1.125 1.125 0 1 1 2.25 0a1.125 1.125 0 0 1-2.25 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IntegrationObservability(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.282 7.4a4.5 4.5 0 0 1-8.666-.376l.851-.168L.847 5.18l-.848 2.167l.897-.176a5.25 5.25 0 0 0 10.125.461zm.836-2.54A5.25 5.25 0 0 0 .993 4.4l.713.232a4.5 4.5 0 0 1 8.666.375l-.851.169l1.62 1.672l.858-2.164z"/><path fill="currentColor" fill-rule="evenodd" d="M3.42 5.93c.914-2.328 4.246-2.328 5.16 0c-.914 2.328-4.246 2.328-5.16 0m5.935-.063c-1.06-3.244-5.65-3.244-6.71 0a.202.202 0 0 0 0 .125c1.06 3.245 5.65 3.245 6.71 0a.202.202 0 0 0 0-.125" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M6 5.55a.375.375 0 1 0 0 .75a.375.375 0 0 0 0-.75m-1.125.375a1.125 1.125 0 1 1 2.25 0a1.125 1.125 0 0 1-2.25 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IntegrationSearch(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.282 7.4a4.5 4.5 0 0 1-8.666-.376l.851-.168L.847 5.18l-.848 2.167l.897-.176a5.25 5.25 0 0 0 10.125.461zm.836-2.54A5.25 5.25 0 0 0 .993 4.4l.713.232a4.5 4.5 0 0 1 8.666.375l-.851.169l1.62 1.672l.858-2.164z"/><path fill="currentColor" d="M9.042 8.132a.378.378 0 1 1-.54.53l-1.355-1.38a.378.378 0 0 1 0-.53a1.63 1.63 0 1 0-1.162.487a.378.378 0 1 1 0 .757a2.385 2.385 0 1 1 1.942-.999z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IntegrationSecurity(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.282 7.4a4.5 4.5 0 0 1-8.666-.376l.851-.168L.847 5.18l-.848 2.167l.897-.176a5.25 5.25 0 0 0 10.125.461zm.836-2.54A5.25 5.25 0 0 0 .993 4.4l.713.232a4.5 4.5 0 0 1 8.666.375l-.851.169l1.62 1.672l.858-2.164z"/><path fill="currentColor" fill-rule="evenodd" d="m9.048 3.977l-.244 1.952a3.75 3.75 0 0 1-1.84 2.779L6 9.267l-.964-.56a3.75 3.75 0 0 1-1.84-2.778l-.244-1.952l.344-.262l.454.597l.19 1.524a3 3 0 0 0 1.472 2.223L6 8.4l.588-.341A3 3 0 0 0 8.06 5.836l.19-1.524l.454-.597zm-.798.335S7.313 3.6 6 3.6s-2.25.712-2.25.712l-.454-.597h.001v-.001l.003-.002l.005-.004l.015-.01a2.48 2.48 0 0 1 .204-.135c.132-.082.319-.187.553-.291C4.54 3.064 5.21 2.85 6 2.85s1.459.214 1.923.422a4.779 4.779 0 0 1 .71.393l.047.033l.015.01l.005.004l.002.002l.002.001z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Invert(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 13.25a5.25 5.25 0 1 0 0-10.5a5.25 5.25 0 0 0 0 10.5M8 14A6 6 0 1 1 8 2a6 6 0 0 1 0 12"/><path fill="currentColor" d="M8 2a6 6 0 1 0 0 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ip(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2zm-2 3H8v6h1V9.014h1c.298-.013 2 0 2-2.018c0-1.74-1.314-1.952-1.825-1.987zM6 5H5v6h1zm4 .984c.667 0 1 .336 1 1.008C11 7.664 10.667 8 10 8H9V5.984Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyboardShortcut(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 9a1 1 0 0 1 .993.883L16 10v5a1 1 0 0 1-.883.993L15 16H7a1 1 0 0 1-.993-.883L6 15v-5a1 1 0 0 1 .883-.993L7 9zM2.5 10a.5.5 0 0 1 .492.41L3 10.5V12h1.5a.5.5 0 0 1 .09.992L4.5 13H3v1.5a.5.5 0 0 1-.992.09L2 14.5V13H.5a.5.5 0 0 1-.09-.992L.5 12H2v-1.5a.5.5 0 0 1 .5-.5M15 10H7v5h8zm-1 3v1H8v-1zm1-13a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1H1a1 1 0 0 1-1-1V1a1 1 0 0 1 1-1zm0 1H1v5h14zM8 4v1H2V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KqlField(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 9a5 5 0 1 1 0-8a5 5 0 1 1 0 8m.75-.692a4 4 0 1 0 0-6.615A4.981 4.981 0 0 1 10 5a4.981 4.981 0 0 1-1.25 3.308M4.133 8V5.559h2.496v-.625H4.133V2.996h2.719v-.633H3.43V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KqlFunction(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 7H3v2h4v2l3-3l-3-3zM6 6V5a1 1 0 0 1 1.707-.707l3 3a1 1 0 0 1 0 1.414l-3 3A1 1 0 0 1 6 11v-1H3a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1zm7.5-3a.5.5 0 0 1 .5.5v9a.5.5 0 1 1-1 0v-9a.5.5 0 0 1 .5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KqlOperand(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.192 10.145l2.298-1.792c.259-.196.259-.509 0-.706l-2.298-1.792c-.256-.196-.256-.513 0-.708a.81.81 0 0 1 .93 0l2.3 1.791c.772.59.77 1.537 0 2.124l-2.3 1.791a.81.81 0 0 1-.93 0c-.256-.195-.256-.512 0-.708m-6.384-4.29L2.51 7.647c-.259.196-.259.509 0 .706l2.298 1.792c.256.196.256.513 0 .708a.81.81 0 0 1-.93 0l-2.3-1.791c-.772-.59-.77-1.537 0-2.124l2.3-1.791a.81.81 0 0 1 .93 0c.256.195.256.512 0 .708M6.5 6h3a.5.5 0 0 1 0 1h-3a.5.5 0 0 1 0-1m0 3h3a.5.5 0 0 1 0 1h-3a.5.5 0 0 1 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KqlSelector(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8m0 1A5 5 0 1 1 5 3a5 5 0 0 1 0 10m6-1a4 4 0 1 0 0-8a4 4 0 0 0 0 8m0 1a5 5 0 1 1 0-10a5 5 0 0 1 0 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KqlValue(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 4a5 5 0 1 1 0 8a5 5 0 1 1 0-8m-.75.692a4 4 0 1 0 0 6.615A4.981 4.981 0 0 1 6 8c0-1.268.472-2.426 1.25-3.308M11.348 11l2.078-5.637h-.739l-1.656 4.727h-.062L9.313 5.363h-.739L10.652 11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Layers(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.276 1.053a.5.5 0 0 1 .448 0l6 3a.5.5 0 0 1 0 .894l-6 3a.5.5 0 0 1-.448 0l-6-3a.5.5 0 0 1 0-.894zM2.618 4.5L7.5 6.941L12.382 4.5L7.5 2.059z"/><path fill="currentColor" d="M1.053 7.276a.5.5 0 0 1 .67-.223L7.5 9.94l5.776-2.888a.5.5 0 1 1 .448.894l-6 3a.5.5 0 0 1-.448 0l-6-3a.5.5 0 0 1-.223-.67"/><path fill="currentColor" d="M1.724 10.053a.5.5 0 1 0-.448.894l6 3a.5.5 0 0 0 .448 0l6-3a.5.5 0 1 0-.448-.894L7.5 12.94l-5.776-2.888z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.66 3.803a.5.5 0 1 1-.706-.707L9.268.78c1.187-1.187 3.242-1 4.596.354s1.54 3.409.354 4.596l-3.536 3.536c-1.187 1.187-3.242 1-4.596-.354a.5.5 0 1 1 .707-.707c.99.99 2.417 1.119 3.182.354l3.536-3.536c.765-.765.635-2.193-.354-3.182c-.99-.99-2.417-1.119-3.182-.354zm-.32 7.392a.5.5 0 1 1 .707.707l-2.315 2.314c-1.187 1.188-3.242 1-4.596-.353c-1.354-1.354-1.54-3.41-.353-4.596L4.318 5.73c1.187-1.187 3.242-1 4.596.354a.5.5 0 0 1-.707.707c-.989-.99-2.416-1.12-3.182-.354L1.49 9.974c-.766.765-.636 2.193.353 3.182c.99.989 2.417 1.119 3.182.353z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func List(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 4V3h2v1zm4 0V3h8v1zm0 3V6h8v1zm0 3V9h8v1zM2 7V6h2v1zm0 3V9h2v1zm4 3v-1h8v1zm-4 0v-1h2v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListAdd(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 11H9v1h2v2h1v-2h2v-1h-2V9h-1zM7.758 9a4.5 4.5 0 1 1-.502 4H6v-1h1.027a4.548 4.548 0 0 1 .23-2H6V9zM2 4V3h2v1zm4 0V3h8v1zm0 3V6h8v1zM2 7V6h2v1zm0 3V9h2v1zm0 3v-1h2v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 5v-.8C4 1.88 5.79 0 8 0s4 1.88 4 4.2V5h1.143c.473 0 .857.448.857 1v9c0 .552-.384 1-.857 1H2.857C2.384 16 2 15.552 2 15V6c0-.552.384-1 .857-1zM3 15h10V6H3zm5.998-3.706L9.5 12.5h-3l.502-1.206A1.644 1.644 0 0 1 6.5 10.1c0-.883.672-1.6 1.5-1.6s1.5.717 1.5 1.6c0 .475-.194.901-.502 1.194M11 4.36C11 2.504 9.657 1 8 1S5 2.504 5 4.36V5h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOpen(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.143 5c.473 0 .857.448.857 1v9c0 .552-.384 1-.857 1H.857C.384 16 0 15.552 0 15V6c0-.552.384-1 .857-1H8v-.8C8 1.88 9.79 0 12 0s4 1.88 4 4.2V5h-1v-.64C15 2.504 13.657 1 12 1S9 2.504 9 4.36V5zM1 15h10V6H1zm5.998-3.706L7.5 12.5h-3l.502-1.206A1.644 1.644 0 0 1 4.5 10.1c0-.883.672-1.6 1.5-1.6s1.5.717 1.5 1.6c0 .475-.194.901-.502 1.194"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogstashFilter(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.914 1c-3.6 0-5.897 1.111-5.897 1.876c0 .766 2.297 1.877 5.897 1.877s5.897-1.111 5.897-1.877C13.811 2.111 11.514 1 7.914 1m3.98 7.988c-1.187.331-2.601.502-3.98.502c-1.383 0-2.801-.172-3.991-.505l2.863 3.242v2.515c.136.08.515.214 1.128.214c.61 0 .989-.133 1.126-.213v-2.515zm2.394-1.206l-4.248 4.823v2.246h.004c0 .763-1.069 1.105-2.13 1.105c-1.06 0-2.13-.342-2.13-1.105h.002v-2.245L1.321 7.55l.01-.008A1.53 1.53 0 0 1 1 6.607V2.78h.02C1.14.973 4.627 0 7.913 0c3.286 0 6.774.973 6.894 2.78h.02v3.827c0 .343-.12.657-.335.941l-.003.004c-.06.079-.128.155-.202.23m-.578-.856a.56.56 0 0 0 .118-.319V4.411c-1.288.879-3.649 1.342-5.914 1.342C5.65 5.753 3.288 5.29 2 4.411v2.196C2 7.375 4.304 8.49 7.914 8.49c2.879 0 4.927-.709 5.639-1.385z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogstashIf(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.918 9.799l-4.82 4.819l2.88-.43l.148.987L0 15.943l.768-5.126l.988.148l-.453 3.033l4.936-4.935a2.003 2.003 0 0 1 1.265-2.885V0h1v6.178a2.003 2.003 0 0 1 1.263 2.884l4.865 4.866l-.443-2.963l.988-.147l.768 5.126l-5.127-.769l.15-.988l2.95.441l-4.83-4.83c-.312.203-.685.32-1.085.32c-.4 0-.773-.117-1.085-.32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogstashInput(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.747 10.992h13.1a.123.123 0 0 0 .122-.123V8.51a.122.122 0 0 0-.122-.122H1.122A.122.122 0 0 0 1 8.51v2.36c0 .066.055.122.122.122zm12.011 1H2.21V16h-1v-4.008h-.088A1.124 1.124 0 0 1 0 10.87V8.51c0-.62.503-1.122 1.122-1.122h13.725c.62 0 1.122.502 1.122 1.122v2.36c0 .618-.503 1.122-1.122 1.122h-.089V16h-1zm-6.27-7.487V0h1v4.529l2.407-2.262l.685.73L8 6.356L4.42 2.995l.685-.729z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogstashOutput(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.21 4.008h11.55V0h1v4.008h.088c.619 0 1.122.504 1.122 1.123V7.49c0 .62-.503 1.122-1.122 1.122H1.122A1.122 1.122 0 0 1 0 7.49V5.13c0-.618.503-1.122 1.122-1.122h.089V0h1zm11.549 1H1.12A.123.123 0 0 0 1 5.13v2.36c0 .068.055.122.122.122h13.725a.122.122 0 0 0 .122-.122V5.13a.123.123 0 0 0-.122-.122zm-5.301 9.097l2.405-2.26l.686.728l-3.58 3.363l-3.58-3.363l.686-.728l2.383 2.24V9.578h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogstashQueue(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.339 15.464H4.77a3.248 3.248 0 0 1-3.245-3.244V4.549H0v-1h2.526v8.67a2.247 2.247 0 0 0 2.245 2.245h6.568a2.247 2.247 0 0 0 2.244-2.244V3.549h2.455v1h-1.455v7.67a3.247 3.247 0 0 1-3.244 3.245m.513-5.962v1.095l-3.848 1.72l-3.85-1.72V9.502l3.85 1.72zm0-4.251v1.095l-3.848 1.72l-3.85-1.72V5.25l3.85 1.72zm0-4.251v1.095l-3.848 1.72l-3.85-1.72V1l3.85 1.72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Magnet(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.552 10.71a3.008 3.008 0 0 0 4.19.738l1.642-1.15l1.15 1.642l-1.643 1.15a5.013 5.013 0 1 1-5.75-8.212l1.642-1.15l1.15 1.643l-1.642 1.15a3.007 3.007 0 0 0-.739 4.189m8.296-2.137l1.15 1.643l-1.643 1.149l-1.15-1.642zm-4.6-6.571l1.15 1.643l-1.643 1.15l-1.15-1.642zm1.97 1.068L9.07 1.428a1.003 1.003 0 0 0-1.397-.246L3.566 4.057A5.995 5.995 0 0 0 1.092 7.94a5.993 5.993 0 0 0 .996 4.495a5.99 5.99 0 0 0 3.883 2.473a5.991 5.991 0 0 0 4.495-.996l4.107-2.875c.454-.318.563-.943.246-1.396l-1.15-1.643a1.002 1.002 0 0 0-1.396-.246l-4.107 2.875a2.002 2.002 0 0 1-1.498.332a2 2 0 0 1-1.627-2.323c.09-.505.371-.976.824-1.294l4.107-2.876c.454-.317.564-.942.246-1.396"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifyWithMinus(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.5 6a.5.5 0 0 1 0 1h-6a.5.5 0 0 1 0-1zm.74 4.74a.5.5 0 0 1 .107-.31A5.478 5.478 0 0 0 12 6.5A5.5 5.5 0 1 0 6.5 12a.5.5 0 1 1 0 1a6.5 6.5 0 1 1 4.936-2.27l4.419 4.418a.5.5 0 0 1-.707.707l-4.768-4.768a.499.499 0 0 1-.14-.347"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifyWithPlus(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 6h2.5a.5.5 0 0 1 0 1H7v2.5a.5.5 0 0 1-1 0V7H3.5a.5.5 0 0 1 0-1H6V3.5a.5.5 0 0 1 1 0zm3.24 4.74a.5.5 0 0 1 .107-.31A5.478 5.478 0 0 0 12 6.5A5.5 5.5 0 1 0 6.5 12a.5.5 0 1 1 0 1a6.5 6.5 0 1 1 4.936-2.27l4.419 4.418a.5.5 0 0 1-.707.707l-4.768-4.768a.499.499 0 0 1-.14-.347"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarker(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.55 14.224a.502.502 0 0 0 .897 0l4-8a.5.5 0 0 0 .053-.235a4.672 4.672 0 0 0-.084-.705a5.538 5.538 0 0 0-.505-1.512C11.189 2.362 9.906 1.5 8 1.5c-1.906 0-3.19.862-3.91 2.272c-.248.485-.41.998-.506 1.512c-.058.31-.08.554-.084.705a.5.5 0 0 0 .053.235zM8 5a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-3.493.895c.009-.106.027-.253.06-.429c.079-.424.213-.848.413-1.238C5.537 3.138 6.487 2.5 8 2.5c1.513 0 2.463.638 3.02 1.728c.2.39.334.814.413 1.238c.033.176.051.323.06.43L8 12.881z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Memory(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M7 10h2V6H7zm-4 0h2V6H3zm8.025 0h2V6h-2zM3.5 13.75h1v-2.4h-1zm2.675 0h1.001v-2.4H6.175zm2.675 0h1v-2.4h-1zm2.675 0h1v-2.4h-1z"/><path d="M0 3v7.05h1v3.698h1v-3.699h12v3.699h1v-3.699h1V3zm1 6h14V4H1z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Menu(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" d="M0 2h16v2H0zm0 5h16v2H0zm16 5H0v2h16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuDown(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 7.5c0 .276-.216.5-.495.5h-2.01a.503.503 0 0 1-.487-.412L3 7.5c0-.276.216-.5.495-.5h2.01c.243 0 .445.183.487.412zM3.51 4a.513.513 0 0 1-.502-.412L3 3.5c0-.276.228-.5.51-.5h8.98c.25 0 .459.183.502.412L13 3.5c0 .276-.228.5-.51.5H8.493v7.792l2.06-2.06a.5.5 0 1 1 .707.707L9.14 12.56a1.496 1.496 0 0 1-1.114.44h-.033a.501.501 0 0 1-.118-.014a1.493 1.493 0 0 1-.857-.426l-2.122-2.12a.5.5 0 0 1 .708-.708l1.889 1.89V4zM13 7.5c0 .276-.216.5-.495.5h-2.01a.503.503 0 0 1-.487-.412L10 7.5c0-.276.216-.5.495-.5h2.01c.243 0 .445.183.487.412z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuLeft(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.014 7.382a.501.501 0 0 0-.013.152c-.014.4.133.806.439 1.112l2.12 2.122a.5.5 0 1 0 .708-.708L2.208 8H14.5a.5.5 0 0 0 0-1H2.379l1.889-1.89a.5.5 0 0 0-.707-.706L1.44 6.524a1.5 1.5 0 0 0-.426.858M14.5 3h-7a.5.5 0 0 0 0 1h7a.5.5 0 0 0 0-1m0 8h-7a.5.5 0 1 0 0 1h7a.5.5 0 1 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuRight(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.986 7.382a.501.501 0 0 1 .013.152c.014.4-.133.806-.439 1.112l-2.12 2.122a.5.5 0 1 1-.708-.708L13.792 8H1.5a.5.5 0 0 1 0-1h12.121l-1.889-1.89a.5.5 0 0 1 .707-.706l2.121 2.12a1.5 1.5 0 0 1 .426.858M1.5 3h7a.5.5 0 0 1 0 1h-7a.5.5 0 0 1 0-1m0 8h7a.5.5 0 1 1 0 1h-7a.5.5 0 1 1 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuUp(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.01 8.5c0-.276.216-.5.495-.5h2.01c.243 0 .445.183.487.412l.008.088c0 .276-.216.5-.495.5h-2.01a.503.503 0 0 1-.487-.412zM12.5 12c.25 0 .459.183.502.412l.008.088c0 .276-.228.5-.51.5H3.52a.513.513 0 0 1-.502-.412L3.01 12.5c0-.276.228-.5.51-.5h3.987V4.208l-2.06 2.06a.5.5 0 1 1-.707-.707L6.86 3.44A1.496 1.496 0 0 1 7.974 3h.033c.04 0 .08.005.118.014c.314.043.616.185.857.426l2.122 2.12a.5.5 0 0 1-.708.708l-1.889-1.89V12zM3 8.5c0-.276.216-.5.495-.5h2.01c.243 0 .445.183.487.412L6 8.5c0 .276-.216.5-.495.5h-2.01a.503.503 0 0 1-.487-.412z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Merge(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.352 6H2.5a.5.5 0 0 1 0-1h4.852L5.12 2.721a.4.4 0 0 1 .055-.616a.551.551 0 0 1 .705.048l3 3.062c.16.164.16.405 0 .57l-3 3.062A.532.532 0 0 1 5.5 9a.54.54 0 0 1-.325-.106c-.21-.157-.235-.433-.055-.616zm1.296 4H13.5a.5.5 0 0 1 0 1H8.648l2.232 2.278c.18.183.155.46-.055.617A.54.54 0 0 1 10.5 14a.532.532 0 0 1-.38-.153l-3-3.063a.397.397 0 0 1 0-.568l3-3.063a.551.551 0 0 1 .705-.047a.4.4 0 0 1 .055.616z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minimize(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1.146 14.146l4-4a.5.5 0 0 1 .765.638l-.057.07l-4 4a.5.5 0 0 1-.765-.638zl4-4zM6.5 8A1.5 1.5 0 0 1 8 9.5v3a.5.5 0 1 1-1 0v-3a.5.5 0 0 0-.5-.5h-3a.5.5 0 0 1 0-1zm2-5a.5.5 0 0 1 .5.5v3a.5.5 0 0 0 .5.5h3a.5.5 0 1 1 0 1h-3A1.5 1.5 0 0 1 8 6.5v-3a.5.5 0 0 1 .5-.5m1.651 2.146l4-4a.5.5 0 0 1 .765.638l-.057.07l-4 4a.5.5 0 0 1-.765-.638zl4-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<rect width="10" height="1.5" x="3" y="7.25" fill="currentColor" rx=".5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusInCircle(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 0C11.636 0 15 3.364 15 7.5S11.636 15 7.5 15S0 11.636 0 7.5S3.364 0 7.5 0m0 .882a6.618 6.618 0 1 0 0 13.236A6.618 6.618 0 0 0 7.5.882M3.5 7h8a.5.5 0 1 1 0 1h-8a.5.5 0 0 1 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusInCircleFilled(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 0C11.636 0 15 3.364 15 7.5S11.636 15 7.5 15S0 11.636 0 7.5S3.364 0 7.5 0m-4 7a.5.5 0 0 0 0 1h8a.5.5 0 1 0 0-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MlClassificationJob(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 16v5h2.038a13.179 13.179 0 0 0 0 2H7v5H5v-5H0v-2h5v-5zM7 0v5h5v2H7v5H5V7H0V5h5V0zm16 0v5h5v2h-5v2.038a13.179 13.179 0 0 0-2 0V7h-5V5h5V0z"/><path fill="currentColor" d="M22 10c3.073 0 5.877 1.155 8 3.056v3.252A9.82 9.82 0 1 0 16.307 30h-3.251A11.955 11.955 0 0 1 10 22c0-6.627 5.373-12 12-12m1 8v3h3v2h-3v3h-2v-3h-3v-2h3v-3z" class="ouiIcon__fillSecondary"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MlCreateAdvancedJob(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16h-2c0-7.732-6.268-14-14-14S2 8.268 2 16s6.268 14 14 14z"/><path fill="currentColor" d="M27 20v12h-2V20zm-5 4v8h-2v-8zm10-2v10h-2V22zM17 9v6h6v2h-6v6h-2v-6H9v-2h6V9z" class="ouiIcon__fillSecondary"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MlCreateMultiMetricJob(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 24a4 4 0 1 1 0 8a4 4 0 0 1 0-8m24 0a4 4 0 1 1 0 8a4 4 0 0 1 0-8M4 26a2 2 0 1 0 0 4a2 2 0 0 0 0-4m24 0a2 2 0 1 0 0 4a2 2 0 0 0 0-4M17 9v6h6v2h-6v6h-2v-6H9v-2h6V9zM4 0a4 4 0 1 1 0 8a4 4 0 0 1 0-8m24 0a4 4 0 1 1 0 8a4 4 0 0 1 0-8m0 2a2 2 0 1 0 0 4a2 2 0 0 0 0-4M4 2a2 2 0 1 0 0 4a2 2 0 0 0 0-4" class="ouiIcon__fillSecondary"/><path fill="currentColor" d="M21.078 29.05c.102.677.3 1.322.582 1.92A15.963 15.963 0 0 1 16 32c-1.993 0-3.9-.364-5.66-1.03c.281-.598.48-1.243.582-1.92c1.574.614 3.287.95 5.078.95c1.791 0 3.504-.336 5.078-.95m9.892-18.71A15.963 15.963 0 0 1 32 16c0 1.993-.364 3.9-1.03 5.66a6.948 6.948 0 0 0-1.92-.582c.614-1.574.95-3.287.95-5.078c0-1.791-.336-3.504-.95-5.078a6.948 6.948 0 0 0 1.92-.582m-29.94 0c.598.281 1.243.48 1.92.582A13.965 13.965 0 0 0 2 16c0 1.791.336 3.504.95 5.078a6.91 6.91 0 0 0-1.92.582A15.964 15.964 0 0 1 0 16c0-1.993.364-3.9 1.03-5.66M16 0c1.993 0 3.9.364 5.66 1.03a6.948 6.948 0 0 0-.582 1.92A13.965 13.965 0 0 0 16 2c-1.791 0-3.504.336-5.078.95a6.948 6.948 0 0 0-.582-1.92A15.964 15.964 0 0 1 16 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MlCreatePopulationJob(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S6 15.523 6 10S10.477 0 16 0m0 2a8 8 0 1 0 0 16a8 8 0 0 0 0-16"/><path fill="currentColor" d="M8.4 19.288a12.11 12.11 0 0 0 1.566 1.087L7.26 23.682a4 4 0 1 1-1.437-1.243zm15.2 0l2.577 3.15a4 4 0 1 1-1.437 1.243l-2.706-3.306a12.032 12.032 0 0 0 1.565-1.087M4 24a2 2 0 1 0 0 4a2 2 0 0 0 0-4m24 0a2 2 0 1 0 0 4a2 2 0 0 0 0-4M17 6v3h3v2h-3v3h-2v-3h-3V9h3V6z" class="ouiIcon__fillSecondary"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MlCreateSingleMetricJob(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16a16 16 0 0 1-16 16m0-30C8.268 2 2 8.268 2 16s6.268 14 14 14s14-6.268 14-14A14 14 0 0 0 16 2"/><path fill="currentColor" d="M23 15h-6V9h-2v6H9v2h6v6h2v-6h6z" class="ouiIcon__fillSecondary"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MlDataVisualizer(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 20v10h10v2H0V20zm30 0v12H20v-2h10V20zM12 4a8 8 0 1 1 0 16a8 8 0 0 1 0-16m0 2a6 6 0 1 0 0 12a6 6 0 0 0 0-12m0-6v2H2v10H0V0zm20 0v12h-2V2H20V0z"/><path fill="currentColor" d="M21.997 12.251a10.02 10.02 0 0 1-.253 2.006a6 6 0 1 1-7.487 7.487a9.947 9.947 0 0 1-2.006.253a8 8 0 1 0 9.746-9.746" class="ouiIcon__fillSecondary"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MlOutlierDetectionJob(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 12v6h6v2H0v-8zm18 0v8h-8v-2h6v-6zM8 0v2H2v6H0V0zm12 0v8h-2V2h-6V0z"/><path fill="currentColor" d="M16 24a4 4 0 1 1 0 8a4 4 0 0 1 0-8m12 0a4 4 0 1 1 0 8a4 4 0 0 1 0-8m-12 1.75a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5m12 0a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5M28 12a4 4 0 1 1 0 8a4 4 0 0 1 0-8m0 1.75a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5M10 6a4 4 0 1 1 0 8a4 4 0 0 1 0-8m0 1.75a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5" class="ouiIcon__fillSecondary"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MlRegressionJob(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24 0a8 8 0 1 1-4.906 14.32l-4.774 4.774a8 8 0 1 1-1.414-1.414l4.774-4.774A8 8 0 0 1 24 0M8 18a6 6 0 1 0 0 12a6 6 0 0 0 0-12M24 2a6 6 0 1 0 0 12a6 6 0 0 0 0-12"/><path fill="currentColor" d="M32 20v12H20V20zm-2 2h-8v8h8zM12 0v12H0V0zm-2 2H2v8h8z" class="ouiIcon__fillSecondary"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mobile(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 2.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1h-3a.5.5 0 0 1-.5-.5"/><path fill="currentColor" fill-rule="evenodd" d="M4.5 0A1.5 1.5 0 0 0 3 1.5v13A1.5 1.5 0 0 0 4.5 16h7a1.5 1.5 0 0 0 1.5-1.5v-13A1.5 1.5 0 0 0 11.5 0zM4 1.5a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 .5.5V4H4zM4 13v1.5a.5.5 0 0 0 .5.5h7a.5.5 0 0 0 .5-.5V13zm0-1h8V5H4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Moon(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.796 9.048c-1.552-2.238-1.199-5.323.61-8.1c-3.47-.12-6.6 2.232-7.269 5.672c-.742 3.82 1.83 7.533 5.749 8.294a7.226 7.226 0 0 0 7.526-3.218c-2.794.177-5.27-.711-6.616-2.648"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Nested(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 1a.5.5 0 0 1 0 1H3.006C2.45 2 2 2.45 2 3.006v9.988C2 13.55 2.45 14 3.006 14H5.5a.5.5 0 1 1 0 1H3.006A2.005 2.005 0 0 1 1 12.994V3.006C1 1.898 1.897 1 3.006 1zm7.494 0c1.059 0 1.924.818 2 1.856l.006.15v9.988a2.005 2.005 0 0 1-1.856 2l-.15.006H10.5a.5.5 0 0 1-.09-.992L10.5 14h2.494c.516 0 .941-.388 1-.888l.006-.118V3.006c0-.516-.388-.941-.888-1L12.994 2H10.5a.5.5 0 0 1-.09-.992L10.5 1zM5 7a1 1 0 1 1 0 2a1 1 0 0 1 0-2m3 0a1 1 0 1 1 0 2a1 1 0 0 1 0-2m3 0a1 1 0 1 1 0 2a1 1 0 0 1 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Node(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 1.443a1 1 0 0 0-1 0L2.572 4.29a1 1 0 0 0-.5.866v5.69a1 1 0 0 0 .5.866L7.5 14.557a1 1 0 0 0 1 0l4.928-2.846a1 1 0 0 0 .5-.866v-5.69a1 1 0 0 0-.5-.866zM9 .577l4.928 2.846a2 2 0 0 1 1 1.732v5.69a2 2 0 0 1-1 1.732L9 15.423a2 2 0 0 1-2 0l-4.928-2.846a2 2 0 0 1-1-1.732v-5.69a2 2 0 0 1 1-1.732L7 .577a2 2 0 0 1 2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Number(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.808 10.197H6.796L5.859 13H4.485l.937-2.803H3.966l.219-1.25h1.647l.608-1.805H4.991l.226-1.251h1.64l.95-2.844h1.368l-.95 2.844h1.018l.95-2.844h1.374l-.95 2.844h1.51l-.218 1.25h-1.702l-.608 1.805h1.497l-.219 1.251H9.182L8.252 13H6.878zm-.602-1.25h1.012l.615-1.805H7.814z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Offline(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.517 12.271l1.254-1.254a1.5 1.5 0 1 1-1.254 1.254m2.945-2.944l.74-.74c.361.208.694.467.987.772a.5.5 0 0 1-.721.693a3.424 3.424 0 0 0-1.006-.725m2.162-2.163l.716-.715c.308.232.599.49.87.772a.5.5 0 1 1-.722.692a6.26 6.26 0 0 0-.864-.749M7.061 6.07A6.198 6.198 0 0 0 3.54 7.885a.5.5 0 0 1-.717-.697a7.199 7.199 0 0 1 5.309-2.187zm6.672-1.014l.71-.71c.274.23.536.476.786.736a.5.5 0 0 1-.721.692a9.1 9.1 0 0 0-.775-.718m-3.807-1.85A9.06 9.06 0 0 0 8 3a8.99 8.99 0 0 0-6.469 2.734a.5.5 0 1 1-.717-.697A9.99 9.99 0 0 1 8 2c.944 0 1.868.131 2.75.382zM8 13a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1m-5.424 1a.5.5 0 0 1-.707-.707L14.146 1.146a.5.5 0 0 1 .708.708z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Online(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 14a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3m0-1a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1m3.189-3.64a.5.5 0 0 1-.721.692A3.408 3.408 0 0 0 8 9c-.937 0-1.813.378-2.453 1.037a.5.5 0 0 1-.717-.697A4.408 4.408 0 0 1 8 8c1.22 0 2.361.497 3.189 1.36m2.02-2.14a.5.5 0 1 1-.721.693A6.2 6.2 0 0 0 8 6a6.199 6.199 0 0 0-4.46 1.885a.5.5 0 0 1-.718-.697A7.199 7.199 0 0 1 8 5a7.2 7.2 0 0 1 5.21 2.22m2.02-2.138a.5.5 0 0 1-.721.692A8.99 8.99 0 0 0 8 3a8.99 8.99 0 0 0-6.469 2.734a.5.5 0 1 1-.717-.697A9.99 9.99 0 0 1 8 2a9.99 9.99 0 0 1 7.23 3.082"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Package(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.447 3.724l-6-3a1 1 0 0 0-.894 0l-6 3A1 1 0 0 0 1 4.618v6.764a1 1 0 0 0 .553.894l6 3a1 1 0 0 0 .894 0l6-3a1 1 0 0 0 .553-.894V4.618a1 1 0 0 0-.553-.894M5.871 5.897l5.343-2.672l2.158 1.079L8 6.943ZM8 1.618l2.096 1.048l-5.353 2.677l-2.115-1.039ZM2 5.11l2.25 1.105V9a.5.5 0 0 0 1 0V6.706L7.5 7.811v6.321L2 11.382Zm6.5 9.022v-6.32L14 5.11v6.272Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageSelect(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 1a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h5a4.995 4.995 0 0 1-.584-1H3V2h7v2a1 1 0 0 0 1 1h2v2.1c.348.07.682.177 1 .316V4a1 1 0 0 0-.293-.707l-2-2A1 1 0 0 0 11 1zm13 11a4 4 0 1 1-8 0a4 4 0 0 1 8 0m-1.646-1.354a.5.5 0 0 1 0 .708l-2.5 2.5a.5.5 0 0 1-.708 0l-1-1a.5.5 0 0 1 .708-.708l.646.647l2.146-2.147a.5.5 0 0 1 .708 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PagesSelect(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 1a1 1 0 0 1 1-1h8a1 1 0 0 1 .707.293l2 2A1 1 0 0 1 15 3v5a4.995 4.995 0 0 0-1-.584V4h-2a1 1 0 0 1-1-1V1H4v12h3.1c.07.348.177.682.316 1H4a1 1 0 0 1-1-1zm5 14H2V2a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h7a5.029 5.029 0 0 1-1-1m8-3a4 4 0 1 1-8 0a4 4 0 0 1 8 0m-1.646-1.354a.5.5 0 0 1 0 .708l-2.5 2.5a.5.5 0 0 1-.708 0l-1-1a.5.5 0 0 1 .708-.708l.646.647l2.146-2.147a.5.5 0 0 1 .708 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paint(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.993 8.17c0 .83-.673 1.507-1.499 1.507H5.498A1.505 1.505 0 0 1 3.999 8.17V6.662h7.994zm-2.998 5.998c0 .455-.448.827-.999.827c-.55 0-1-.372-1-.827v-3.486h2zM4 5.658h1.262V1.005H4zm2.261 0h1.244V1.005H6.26zm2.244 0H9.74V1.005H8.504zm2.234 0h1.254V1.005h-1.254zM3.008 0L3 8.17a2.509 2.509 0 0 0 2.498 2.512h.5v3.486c0 1.01.896 1.832 1.998 1.832c1.102 0 1.998-.822 1.998-1.832v-3.486h.5a2.509 2.509 0 0 0 2.498-2.512L13 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperClip(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.84 2.019L3.046 8.57c-.987.952-1.133 2.517-.199 3.516c.951 1.021 2.58 1.106 3.64.19c.034-.03.068-.061.1-.092l5.655-5.452a.484.484 0 0 0 0-.703a.53.53 0 0 0-.729 0L5.92 11.421c-.572.551-1.505.657-2.131.163a1.455 1.455 0 0 1-.118-2.211l6.899-6.651a2.646 2.646 0 0 1 3.644 0a2.422 2.422 0 0 1 0 3.513L7.3 12.901c-1.333 1.285-3.497 1.493-4.95.336c-1.54-1.22-1.764-3.411-.5-4.897a3.33 3.33 0 0 1 .238-.252l5.78-5.572a.484.484 0 0 0 0-.703a.53.53 0 0 0-.73 0l-5.78 5.572a4.36 4.36 0 0 0 0 6.324c2.188 2.109 5.202 1.31 6.66-.095l6.925-6.676a3.39 3.39 0 0 0 0-4.92C13.534.66 11.25.66 9.841 2.019z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Partial(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.664 14.871a7 7 0 0 1-2.42-12.778a.5.5 0 0 1 .612.06c.456.431 8.216 8.212 8.98 9.002a.5.5 0 0 1 .063.618a7.002 7.002 0 0 1-7.235 3.098m6.168-3.312a1961.733 1961.733 0 0 0-8.377-8.4a6 6 0 1 0 8.378 8.4m2.095-2.548a.5.5 0 1 1-.99-.144c.01-.066.01-.066.018-.133a6.007 6.007 0 0 0-.034-1.714a.5.5 0 1 1 .987-.163c.108.655.122 1.326.04 1.999zm-1.273-5.138a.5.5 0 1 1-.808.59a6.026 6.026 0 0 0-1.304-1.308a.5.5 0 0 1 .59-.806a7.026 7.026 0 0 1 1.522 1.524M9.169 1.098a.5.5 0 1 1-.166.986a6.105 6.105 0 0 0-1.849-.026a.5.5 0 0 1-.14-.99a7.02 7.02 0 0 1 2.155.03"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 2.5c0-.276.232-.5.5-.5c.276 0 .5.229.5.5v11c0 .276-.232.5-.5.5a.503.503 0 0 1-.5-.5zm7 0c0-.276.232-.5.5-.5c.276 0 .5.229.5.5v11c0 .276-.232.5-.5.5a.503.503 0 0 1-.5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pencil(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.148 3.148L11 2l-9 9v3h3l9-9l-1.144-1.144l-8.002 7.998a.502.502 0 0 1-.708 0a.502.502 0 0 1 0-.708zM11 1c.256 0 .512.098.707.293l3 3a.999.999 0 0 1 0 1.414l-9 9A.997.997 0 0 1 5 15H2a1 1 0 0 1-1-1v-3c0-.265.105-.52.293-.707l9-9A.997.997 0 0 1 11 1M5 14H2v-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Percent(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 8c1.105 0 2-1.12 2-2.5S6.105 3 5 3S3 4.12 3 5.5S3.895 8 5 8m0-1c.356 0 1-.452 1-1.5S5.356 4 5 4s-1 .452-1 1.5S4.644 7 5 7"/><path fill="currentColor" d="M10.5 3H12L5.5 13H4z"/><path fill="currentColor" fill-rule="evenodd" d="M13 10.5c0 1.38-.895 2.5-2 2.5s-2-1.12-2-2.5S9.895 8 11 8s2 1.12 2 2.5m-1 0c0 1.048-.644 1.5-1 1.5s-1-.452-1-1.5s.644-1.5 1-1.5s1 .452 1 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pin(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 11h4.5a.5.5 0 1 0 0-1h-10a.5.5 0 0 0 0 1H7v4.25c0 .414.224.75.5.75s.5-.336.5-.75zM4 4h1v6H4zm6 0h1v6h-1zM4.286 2C4.08 2 4 2.063 4 2v1c0-.063.08 0 .286 0h6.428C10.92 3 11 2.937 11 3V2c0 .063-.08 0-.286 0zm0-1h6.428C11.424 1 12 1.448 12 2v1c0 .552-.576 1-1.286 1H4.286C3.576 4 3 3.552 3 3V2c0-.552.576-1 1.286-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinFilled(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 11h4.5a.5.5 0 1 0 0-1h-10a.5.5 0 1 0 0 1H8v4.25c0 .414.224.75.5.75s.5-.336.5-.75zM5 4h7v6H5zm.286-3h6.428C12.424 1 13 1.448 13 2v1c0 .552-.576 1-1.286 1H5.286C4.576 4 4 3.552 4 3V2c0-.552.576-1 1.286-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.608 3.063C4.345 2.895 4 3.089 4 3.418v9.167c0 .329.345.523.608.356l7.2-4.584a.426.426 0 0 0 0-.711zm.538-.844l7.2 4.583a1.426 1.426 0 0 1 0 2.399l-7.2 4.583C4.21 14.38 3 13.696 3 12.585V3.418C3 2.307 4.21 1.624 5.146 2.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayFilled(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" d="m12.345 6.801l-7.2-4.581C4.21 1.625 3 2.308 3 3.419v9.162c0 1.111 1.21 1.794 2.146 1.199l7.2-4.581a1.425 1.425 0 0 0 0-2.398z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.25 3a.5.5 0 0 1 .5.5v3.75h3.75a.5.5 0 0 1 .5.5v.5a.5.5 0 0 1-.5.5H8.75v3.75a.5.5 0 0 1-.5.5h-.5a.5.5 0 0 1-.5-.5V8.75H3.5a.5.5 0 0 1-.5-.5v-.5a.5.5 0 0 1 .5-.5h3.75V3.5a.5.5 0 0 1 .5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusInCircle(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 7h3.5a.5.5 0 1 1 0 1H8v3.5a.5.5 0 1 1-1 0V8H3.5a.5.5 0 0 1 0-1H7V3.5a.5.5 0 0 1 1 0zm-.5-7C11.636 0 15 3.364 15 7.5S11.636 15 7.5 15S0 11.636 0 7.5S3.364 0 7.5 0m0 .882a6.618 6.618 0 1 0 0 13.236A6.618 6.618 0 0 0 7.5.882"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusInCircleFilled(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 7V3.5a.5.5 0 0 0-1 0V7H3.5a.5.5 0 0 0 0 1H7v3.5a.5.5 0 1 0 1 0V8h3.5a.5.5 0 1 0 0-1zm-.5 8a7.5 7.5 0 1 1 0-15a7.5 7.5 0 0 1 0 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Polygon(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 11v4h-4v-1.5H5V15H1v-4h1.5V5H1V1h4v1.3h6V1h4v4h-2.9l-.8 1h.7v4h-.8l.8 1zm-1 3v-2h-2v2zm0-10V2h-2v2zm-3 8.5v-1.2L10 10H8V6h2l1-1.3V3.3H5V5H3.5v6H5v1.5zM11 9V7H9v2zm-7 5v-2H2v2zM4 4V2H2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Popout(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 8.5a.5.5 0 1 1 1 0V12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h3.5a.5.5 0 0 1 0 1H4a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1zm-5.12.339a.5.5 0 1 1-.706-.707L13.305 2H10.5a.5.5 0 1 1 0-1H14a1 1 0 0 1 1 1v3.5a.5.5 0 1 1-1 0V2.72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Power(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a.5.5 0 0 1 .5.5v5a.5.5 0 0 1-1 0v-5A.5.5 0 0 1 8 0"/><path fill="currentColor" d="M10.55.932a.5.5 0 0 1 .663-.246C13.849 1.896 16 4.437 16 7.809C16 12.32 12.43 16 8 16s-8-3.68-8-8.191C0 4.436 2.17 1.896 4.79.686a.5.5 0 0 1 .42.908C2.86 2.68 1 4.908 1 7.808C1 11.792 4.146 15 8 15s7-3.208 7-7.192c0-2.901-1.845-5.13-4.204-6.213a.5.5 0 0 1-.246-.663"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Push(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.171 5.15L10.114 7H1.556C1.249 7 1 7.224 1 7.5s.249.5.556.5h8.526l-1.91 1.82a.52.52 0 0 0 0 .77c.227.213.6.213.828 0l2.05-1.95a1.552 1.552 0 0 0 0-2.31L9 4.38a.617.617 0 0 0-.829 0a.52.52 0 0 0 0 .77"/><path fill="currentColor" d="M6.804 12.792A.993.993 0 0 1 6 11.82V10H5v1.826c0 .945.673 1.76 1.608 1.945l6 1.19A1.992 1.992 0 0 0 15 13.016V1.984A2 2 0 0 0 12.608.04l-6 1.19C5.673 1.415 5 2.23 5 3.175V5h1V3.18c0-.472.336-.879.804-.972l6-1.189A1 1 0 0 1 14 1.991v11.018a.995.995 0 0 1-1.196.972z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionInCircle(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 14A6 6 0 1 1 8 2a6 6 0 0 1 0 12m0-1A5 5 0 1 0 8 3a5 5 0 0 0 0 10m-.186-1.065A.785.785 0 0 1 7 11.12c0-.48.34-.82.814-.82c.475 0 .809.34.809.82c0 .475-.334.815-.809.815M5.9 6.317C5.96 5.168 6.755 4.4 8.048 4.4c1.218 0 2.091.759 2.091 1.8c0 .736-.36 1.304-1.03 1.707c-.56.33-.717.56-.717 1.022v.305l-.1.1H7.47l-.1-.1v-.431c-.005-.646.302-1.104.987-1.514c.527-.322.708-.59.708-1.047c0-.536-.416-.91-1.05-.91c-.652 0-1.064.374-1.112.998l-.1.092H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Quote(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.848 2.47a1 1 0 0 1-.318 1.378A7.284 7.284 0 0 0 3.75 7.01A3 3 0 1 1 1 10v-.027a3.521 3.521 0 0 1 .01-.232c.009-.15.027-.36.062-.618c.07-.513.207-1.22.484-2.014c.552-1.59 1.67-3.555 3.914-4.957a1 1 0 0 1 1.378.318m7 0a1 1 0 0 1-.318 1.378a7.283 7.283 0 0 0-2.78 3.162A3 3 0 1 1 8 10v-.027a3.521 3.521 0 0 1 .01-.232c.009-.15.027-.36.062-.618c.07-.513.207-1.22.484-2.014c.552-1.59 1.67-3.555 3.914-4.957a1 1 0 0 1 1.378.318"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Radius(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.05 12.95a7 7 0 1 0 9.9-9.9a7 7 0 0 0-9.9 9.9m.707-.707a6 6 0 0 0 8.825-8.118L10 6.707V10H6V6h3.293l2.582-2.582a6 6 0 0 0-8.118 8.825M9 7v2H7V7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Redeploy(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.18 7c-.473 0-.88.294-.972.703l-1.189 5.25a.776.776 0 0 0-.019.172c0 .483.444.875.99.875h11.02c.065 0 .13-.006.194-.017c.537-.095.885-.556.778-1.03l-1.19-5.25C13.7 7.294 13.293 7 12.822 7zM5 6v1h7V6h.825c.946 0 1.76.606 1.946 1.447l1.19 5.4c.215.975-.482 1.923-1.556 2.118a2.175 2.175 0 0 1-.39.035H2.985C1.888 15 1 14.194 1 13.2c0-.118.013-.237.039-.353l1.19-5.4C2.414 6.606 3.229 6 4.174 6z"/><path fill="currentColor" d="M5.99 1.399a.5.5 0 0 0-.98.202l2.058 9.945c.327 1.582 2.58 1.6 2.933.023l1.845-8.274L12.6 4.3a.5.5 0 0 0 .8-.6l-1.5-2a.5.5 0 0 0-.7-.1l-2 1.5a.5.5 0 1 0 .6.8l1.065-.799l-1.84 8.25c-.118.526-.869.52-.978-.008z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Refresh(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.228 2.942a.5.5 0 1 1-.538.842A5 5 0 1 0 13 8a.5.5 0 1 1 1 0a6 6 0 1 1-2.772-5.058M14 1.5v3A1.5 1.5 0 0 1 12.5 6h-3a.5.5 0 0 1 0-1h3a.5.5 0 0 0 .5-.5v-3a.5.5 0 1 1 1 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reporter(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.684.895L6.074.358a1 1 0 0 0-1.296.753L4.2 4H2.5a.5.5 0 0 0 0 1h1.626a4.007 4.007 0 0 0 .11 2.359l-2.072-.345A1 1 0 0 0 1 8v1c.364 0 .706.097 1 .268V8l1 .167l1.859.31l2.163.36l.478.08v6L2 14v-1.268A1.99 1.99 0 0 1 1 13v1a1 1 0 0 0 .836.986l6 1c.108.018.22.018.328 0l6-1A1 1 0 0 0 15 14v-1a1.99 1.99 0 0 1-1-.268V14l-5.5.917v-6l.478-.08l2.163-.36L13 8.166L14 8v1.268A1.99 1.99 0 0 1 15 9V8a1 1 0 0 0-1.164-.986l-2.073.345A3.991 3.991 0 0 0 11.874 5H13.5a.5.5 0 0 0 0-1h-1.7l-.578-2.89A1 1 0 0 0 9.925.359L8.316.895a1 1 0 0 1-.632 0m2.88 6.664A3.013 3.013 0 0 0 10.83 5H5.17a3.013 3.013 0 0 0 .266 2.559L8 7.986zM10.8 4H9.2L9 3l1.5-.5zM1 12a1 1 0 1 0 0-2a1 1 0 0 0 0 2m14 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReturnKey(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.994 4c1.059 0 1.924.818 2 1.856l.006.15v1.988a2.005 2.005 0 0 1-1.856 2L12 10H3.484l1.91 1.82a.52.52 0 0 1 0 .77a.616.616 0 0 1-.829 0l-2.05-1.95a1.551 1.551 0 0 1 0-2.31l2.05-1.95a.617.617 0 0 1 .83 0a.52.52 0 0 1 0 .77L3.45 9H12c.514-.003.935-.39.993-.888L13 7.994V6.006c0-.516-.388-.941-.888-1L11.994 5H9.5a.5.5 0 0 1-.09-.992L9.5 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Save(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.008 2H2.282c-.181 0-.245.002-.275.007c-.005.03-.007.094-.007.275v11.436c0 .181.002.245.007.275c.03.005.094.007.275.007h11.436c.181 0 .245-.002.275-.007c.005-.03.007-.094.007-.275V4.62c0-.13-.001-.18-.004-.204a2.654 2.654 0 0 0-.141-.147L11.73 2.145a2.654 2.654 0 0 0-.147-.141A2.654 2.654 0 0 0 11.38 2h-.388c.005.08.008.172.008.282v2.436c0 .446-.046.607-.134.77a.909.909 0 0 1-.378.378c-.163.088-.324.134-.77.134H6.282c-.446 0-.607-.046-.77-.134a.909.909 0 0 1-.378-.378C5.046 5.325 5 5.164 5 4.718V2.282c0-.11.003-.202.008-.282M2.282 1h9.098c.259 0 .348.01.447.032a.87.87 0 0 1 .273.113c.086.054.156.11.338.293l2.124 2.124c.182.182.239.252.293.338a.87.87 0 0 1 .113.273c.023.1.032.188.032.447v9.098c0 .446-.046.607-.134.77a.909.909 0 0 1-.378.378c-.163.088-.324.134-.77.134H2.282c-.446 0-.607-.046-.77-.134a.909.909 0 0 1-.378-.378c-.088-.163-.134-.324-.134-.77V2.282c0-.446.046-.607.134-.77a.909.909 0 0 1 .378-.378c.163-.088.324-.134.77-.134M6 2.282v2.436c0 .181.002.245.007.275c.03.005.094.007.275.007h3.436c.181 0 .245-.002.275-.007c.005-.03.007-.094.007-.275V2.282c0-.181-.002-.245-.007-.275A2.248 2.248 0 0 0 9.718 2H6.282c-.181 0-.245.002-.275.007c-.005.03-.007.094-.007.275M8 12a2 2 0 1 1 0-4a2 2 0 0 1 0 4m0-1a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scale(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 12a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1m-2 0a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1m-2 0a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1m4-2a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1m-2 0a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1m2-1a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0-3a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1m-2 2a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1m-2 0a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1m0 2a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1m-2 2a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1m-2 0a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1m2-2a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1m6-6a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1m-2 2a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.271 11.978l3.872 3.873a.502.502 0 0 0 .708 0a.502.502 0 0 0 0-.708l-3.565-3.564c2.38-2.747 2.267-6.923-.342-9.532c-2.73-2.73-7.17-2.73-9.898 0c-2.728 2.729-2.728 7.17 0 9.9a6.955 6.955 0 0 0 4.949 2.05a.5.5 0 0 0 0-1a5.96 5.96 0 0 1-4.242-1.757a6.01 6.01 0 0 1 0-8.486a6.004 6.004 0 0 1 8.484 0a6.01 6.01 0 0 1 0 8.486a.5.5 0 0 0 .034.738"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SecuritySignal(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.875 3.418a6 6 0 1 0 .707.707l-2.46 2.46l-1.156 1.156a1 1 0 1 1-.707-.707l.757-.757a2 2 0 0 0-2.43 3.137a.5.5 0 1 1-.707.707a3 3 0 0 1 3.86-4.567l.714-.714A4 4 0 1 0 8 12a.5.5 0 1 1 0 1a5 5 0 1 1 3.164-8.871l.71-.71zm.709-.709a7 7 0 1 0 .707.707l.366-.366a.5.5 0 1 0-.707-.707z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SecuritySignalDetected(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.657 3.05a.5.5 0 0 0-.707-.707l-.366.366A7 7 0 1 0 8 15a4.994 4.994 0 0 1-.597-1.03a6 6 0 1 1 4.471-10.552l-.71.71a5 5 0 1 0-4.08 8.788a5.027 5.027 0 0 1-.082-1.042A4.002 4.002 0 0 1 8 4a3.98 3.98 0 0 1 2.453.84l-.715.714a3 3 0 0 0-3.86 4.567a.5.5 0 1 0 .708-.707a2 2 0 0 1 2.43-3.137l-.757.757a1 1 0 1 0 .707.707l1.155-1.155l2.46-2.46a5.972 5.972 0 0 1 1.39 3.277c.367.158.713.36 1.029.597c0-1.636-.57-3.271-1.71-4.584zM16 12a4 4 0 1 1-8 0a4 4 0 0 1 8 0m-4 .5a.577.577 0 0 1-.57-.495l-.29-2.015a.867.867 0 1 1 1.718 0l-.288 2.015a.577.577 0 0 1-.57.495m0 2.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SecuritySignalResolved(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.657 3.05a.5.5 0 1 0-.707-.707l-.366.366A7 7 0 1 0 8 15a4.994 4.994 0 0 1-.597-1.03a6 6 0 1 1 4.471-10.552l-.71.71a5 5 0 1 0-4.08 8.788a5.028 5.028 0 0 1-.082-1.042A4.002 4.002 0 0 1 8 4a3.98 3.98 0 0 1 2.453.84l-.715.714a3 3 0 0 0-3.86 4.567a.5.5 0 1 0 .708-.707a2 2 0 0 1 2.43-3.137l-.757.757a1 1 0 1 0 .707.707l1.155-1.155l2.46-2.46a5.972 5.972 0 0 1 1.39 3.277c.367.158.713.36 1.029.597c0-1.636-.57-3.271-1.71-4.584zM16 12a4 4 0 1 1-8 0a4 4 0 0 1 8 0m-1.646-1.354a.5.5 0 0 1 0 .707l-2.5 2.5a.5.5 0 0 1-.708 0l-1-1a.5.5 0 0 1 .708-.707l.646.647l2.146-2.147a.5.5 0 0 1 .708 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shard(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.134 7.031L8 12.764l2.866-5.733L8 2.016zM8 0l4 7l-4 8l-4-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 6.81v6.38c0 .493.448.9.992.9h7.016c.543 0 .992-.406.992-.9V6.81c0-.493-.448-.9-.992-.9H3.992c-.543 0-.992.406-.992.9M6 5v.91h3V5h2.008C12.108 5 13 5.818 13 6.81v6.38c0 1-.9 1.81-1.992 1.81H3.992C2.892 15 2 14.182 2 13.19V6.81C2 5.81 2.9 5 3.992 5zm1.997-3.552A.506.506 0 0 1 8 1.5v8a.5.5 0 0 1-1 0v-8a.51.51 0 0 1 0-.017L5.18 3.394a.52.52 0 0 1-.77 0a.617.617 0 0 1 0-.829L6.36.515a1.552 1.552 0 0 1 2.31 0l1.95 2.05a.617.617 0 0 1 0 .83a.52.52 0 0 1-.77 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Snowflake(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.007.5a.5.5 0 0 1 1 0v2.024a.999.999 0 0 0 .268-.227l1.027-1.233a.5.5 0 0 1 .769.64L8.275 3.86a.999.999 0 0 1-.268.227v2.548l2.207-1.274c0-.114.02-.231.062-.346l.968-2.632a.5.5 0 1 1 .938.345l-.554 1.506a.998.998 0 0 0-.062.346l1.753-1.012a.5.5 0 1 1 .5.866l-1.753 1.012c.1.057.21.098.33.119l1.582.273a.5.5 0 1 1-.17.986l-2.764-.478a1 1 0 0 1-.33-.12L8.506 7.5l2.207 1.274a1 1 0 0 1 .33-.119l2.764-.478a.5.5 0 1 1 .17.986l-1.582.273a.999.999 0 0 0-.33.12l1.753 1.011a.5.5 0 1 1-.5.866l-1.753-1.012c0 .115.02.231.062.346l.554 1.506a.5.5 0 0 1-.938.345l-.968-2.632a.999.999 0 0 1-.062-.346L8.007 8.366v2.548c.098.058.19.133.268.227l1.796 2.155a.5.5 0 0 1-.769.64l-1.027-1.233a.999.999 0 0 0-.268-.226V14.5a.5.5 0 0 1-1 0v-2.024a.999.999 0 0 0-.269.227l-1.027 1.233a.5.5 0 0 1-.768-.64l1.795-2.155a.999.999 0 0 1 .269-.227V8.366L4.8 9.64c0 .114-.02.231-.062.346l-.969 2.632a.5.5 0 1 1-.938-.345l.554-1.506a1 1 0 0 0 .062-.346l-1.753 1.012a.5.5 0 0 1-.5-.866l1.753-1.012a.999.999 0 0 0-.33-.119l-1.582-.273a.5.5 0 0 1 .17-.986l2.764.478c.12.02.232.062.33.12L6.508 7.5L4.3 6.226a.999.999 0 0 1-.33.119l-2.765.478a.5.5 0 1 1-.17-.986l1.582-.273a.999.999 0 0 0 .33-.12l-1.753-1.01a.5.5 0 1 1 .5-.866L3.447 4.58c0-.114-.02-.231-.062-.346L2.83 2.727a.5.5 0 1 1 .938-.345l.969 2.632a.999.999 0 0 1 .062.346l2.207 1.274V4.086a1 1 0 0 1-.269-.227L4.943 1.704a.5.5 0 0 1 .768-.64l1.027 1.233c.079.094.17.17.269.227z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortDown(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 11.692V3.556C7 3.249 7.224 3 7.5 3s.5.249.5.556v8.136l4.096-4.096a.5.5 0 0 1 .707.707l-4.242 4.243a1.494 1.494 0 0 1-.925.433a.454.454 0 0 1-.272 0a1.494 1.494 0 0 1-.925-.433L2.197 8.303a.5.5 0 1 1 .707-.707z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortLeft(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.308 7h8.136c.307 0 .556.224.556.5s-.249.5-.556.5H4.308l4.096 4.096a.5.5 0 0 1-.707.707L3.454 8.561a1.494 1.494 0 0 1-.433-.925a.454.454 0 0 1 0-.272c.03-.338.175-.666.433-.925l4.243-4.242a.5.5 0 1 1 .707.707z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortRight(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.692 7H3.556C3.249 7 3 7.224 3 7.5s.249.5.556.5h8.136l-4.096 4.096a.5.5 0 0 0 .707.707l4.243-4.242c.258-.259.403-.587.433-.925a.454.454 0 0 0 0-.272a1.494 1.494 0 0 0-.433-.925L8.303 2.197a.5.5 0 1 0-.707.707z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortUp(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 4.207v8.237c0 .307-.224.556-.5.556s-.5-.249-.5-.556V4.207L2.904 8.303a.5.5 0 0 1-.707-.707l4.242-4.242a1.5 1.5 0 0 1 2.122 0l4.242 4.242a.5.5 0 1 1-.707.707z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sortable(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 12.786V3.2L3.604 6.596a.5.5 0 0 1-.708-.707l3.536-3.535a1.5 1.5 0 0 1 2.121 0l3.536 3.535a.5.5 0 0 1-.707.707L8 3.214v9.557l3.382-3.382a.5.5 0 0 1 .707.707l-3.536 3.536a1.5 1.5 0 0 1-2.121 0l-3.536-3.536a.5.5 0 0 1 .708-.707z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarEmpty(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 2a.86.86 0 0 0-.792.511l-1.33 2.924l-3.128.446c-.71.102-1.001.976-.496 1.487l2.313 2.336l-.563 3.268A.877.877 0 0 0 4.864 14a.86.86 0 0 0 .429-.116L8 12.343l2.707 1.542c.14.08.287.116.43.116a.877.877 0 0 0 .859-1.027l-.563-3.269l2.313-2.336c.505-.511.214-1.385-.496-1.487l-3.128-.446l-1.33-2.923A.86.86 0 0 0 8 2m0 1c.073 0 .095.049.104.07l1.267 2.783l.162.356l.387.055l2.978.425a.11.11 0 0 1 .095.08a.116.116 0 0 1-.029.126l-2.202 2.226l-.259.261l.063.362l.535 3.112c.007.04 0 .07-.023.098a.127.127 0 0 1-.091.046a.106.106 0 0 1-.055-.016l-2.578-1.469L8 11.314l-.354.201l-2.579 1.469a.103.103 0 0 1-.054.016a.127.127 0 0 1-.091-.046a.115.115 0 0 1-.024-.098l.536-3.112l.063-.362l-.259-.261l-2.202-2.226a.116.116 0 0 1-.029-.126a.108.108 0 0 1 .094-.08l2.98-.425l.386-.055l.162-.356l1.267-2.786C7.905 3.05 7.927 3 8 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarEmptySpace(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 2a.86.86 0 0 0-.792.511l-1.33 2.924l-3.128.446c-.71.102-1.001.976-.496 1.487l2.313 2.336l-.563 3.268A.877.877 0 0 0 2.864 14a.86.86 0 0 0 .429-.116L6 12.343l2.707 1.542c.14.08.287.116.43.116a.877.877 0 0 0 .859-1.027l-.563-3.269l2.313-2.336c.505-.511.214-1.385-.496-1.487l-3.128-.446l-1.33-2.923A.86.86 0 0 0 6 2m0 1c.073 0 .095.049.104.07l1.267 2.783l.162.356l.387.055l2.978.425a.11.11 0 0 1 .095.08a.116.116 0 0 1-.029.126L8.762 9.121l-.259.261l.063.362l.535 3.112c.007.04 0 .07-.023.098a.127.127 0 0 1-.091.046a.106.106 0 0 1-.055-.016l-2.578-1.469L6 11.314l-.354.201l-2.579 1.469a.103.103 0 0 1-.054.016a.127.127 0 0 1-.091-.046a.115.115 0 0 1-.024-.098l.536-3.112l.063-.362l-.259-.261l-2.202-2.226a.116.116 0 0 1-.029-.126a.108.108 0 0 1 .094-.08l2.98-.425l.386-.055l.162-.356l1.267-2.786C5.905 3.05 5.927 3 6 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarFilled(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 2a.86.86 0 0 0-.792.511l-1.33 2.924l-3.128.446c-.71.102-1.001.976-.496 1.487l2.313 2.336l-.563 3.268A.877.877 0 0 0 4.864 14a.86.86 0 0 0 .429-.116L8 12.343l2.707 1.542c.14.08.287.116.43.116a.877.877 0 0 0 .859-1.027l-.563-3.269l2.313-2.336c.505-.511.214-1.385-.496-1.487l-3.128-.446l-1.33-2.923A.86.86 0 0 0 8 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarFilledSpace(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 2a.86.86 0 0 0-.792.511l-1.33 2.924l-3.128.446c-.71.102-1.001.976-.496 1.487l2.313 2.336l-.563 3.268A.877.877 0 0 0 2.864 14a.86.86 0 0 0 .429-.116L6 12.343l2.707 1.542c.14.08.287.116.43.116a.877.877 0 0 0 .859-1.027l-.563-3.269l2.313-2.336c.505-.511.214-1.385-.496-1.487l-3.128-.446l-1.33-2.923A.86.86 0 0 0 6 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarMinusEmpty(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 2a.86.86 0 0 0-.792.511l-1.33 2.924l-3.128.446c-.71.102-1.001.976-.496 1.487l2.313 2.336l-.563 3.268A.877.877 0 0 0 2.864 14a.86.86 0 0 0 .429-.116L6 12.343l2.707 1.542c.14.08.287.116.43.116a.877.877 0 0 0 .859-1.027l-.563-3.269l2.313-2.336c.505-.511.214-1.385-.496-1.487l-3.128-.446l-1.33-2.923A.86.86 0 0 0 6 2m0 1c.073 0 .095.049.104.07l1.267 2.783l.162.356l.387.055l2.978.425a.11.11 0 0 1 .095.08a.116.116 0 0 1-.029.126L8.762 9.121l-.259.261l.063.362l.535 3.112c.007.04 0 .07-.023.098a.127.127 0 0 1-.091.046a.106.106 0 0 1-.055-.016l-2.578-1.469L6 11.314l-.354.201l-2.579 1.469a.103.103 0 0 1-.054.016a.127.127 0 0 1-.091-.046a.115.115 0 0 1-.024-.098l.536-3.112l.063-.362l-.259-.261l-2.202-2.226a.116.116 0 0 1-.029-.126a.108.108 0 0 1 .094-.08l2.98-.425l.386-.055l.162-.356l1.267-2.786C5.905 3.05 5.927 3 6 3m10 7v1h-5v-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarMinusFilled(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 2a.86.86 0 0 1 .792.511l1.33 2.924l3.128.446c.71.102 1.001.976.496 1.487L9.433 9.704l.563 3.268A.877.877 0 0 1 9.136 14a.862.862 0 0 1-.429-.116L6 12.343l-2.707 1.542a.862.862 0 0 1-.43.116a.877.877 0 0 1-.859-1.027l.563-3.269L.254 7.368C-.25 6.857.04 5.983.75 5.88l3.128-.446l1.33-2.923A.86.86 0 0 1 6 2m10 8v1h-5v-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarPlusEmpty(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 2a.86.86 0 0 0-.792.511l-1.33 2.924l-3.128.446c-.71.102-1.001.976-.496 1.487l2.313 2.336l-.563 3.268A.877.877 0 0 0 2.864 14a.86.86 0 0 0 .429-.116L6 12.343l2.707 1.542c.14.08.287.116.43.116a.877.877 0 0 0 .859-1.027l-.563-3.269l2.313-2.336c.505-.511.214-1.385-.496-1.487l-3.128-.446l-1.33-2.923A.86.86 0 0 0 6 2m8 8h2v1h-2v2h-1v-2h-2v-1h2V8h1zM6 3c.073 0 .095.049.104.07l1.267 2.783l.162.356l.387.055l2.978.425a.11.11 0 0 1 .095.08a.116.116 0 0 1-.029.126L8.762 9.121l-.259.261l.063.362l.535 3.112c.007.04 0 .07-.023.098a.127.127 0 0 1-.091.046a.106.106 0 0 1-.055-.016l-2.578-1.469L6 11.314l-.354.201l-2.579 1.469a.103.103 0 0 1-.054.016a.127.127 0 0 1-.091-.046a.115.115 0 0 1-.024-.098l.536-3.112l.063-.362l-.259-.261l-2.202-2.226a.116.116 0 0 1-.029-.126a.108.108 0 0 1 .094-.08l2.98-.425l.386-.055l.162-.356l1.267-2.786C5.905 3.05 5.927 3 6 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarPlusFilled(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14 10h2v1h-2v2h-1v-2h-2v-1h2V8h1zM6 2a.86.86 0 0 0-.792.511l-1.33 2.924l-3.128.446c-.71.102-1.001.976-.496 1.487l2.313 2.336l-.563 3.268A.877.877 0 0 0 2.864 14a.86.86 0 0 0 .429-.116L6 12.343l2.707 1.542c.14.08.287.116.43.116a.877.877 0 0 0 .859-1.027l-.563-3.269l2.313-2.336c.505-.511.214-1.385-.496-1.487l-3.128-.446l-1.33-2.923A.86.86 0 0 0 6 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stats(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 14v-4h1v4h5V5h1v9a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1v-2h1v2zm4.853-10.146l-2.999 3a1.5 1.5 0 0 1-2.538 1.568l-2.714.904L4 9.527a1.5 1.5 0 1 1-.316-.948L7 7.473a1.5 1.5 0 0 1 2.146-1.327l3-3a1.5 1.5 0 1 1 .707.707"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stop(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2m0 1a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopFilled(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<rect width="12" height="12" x="2" y="2" fill="currentColor" fill-rule="evenodd" rx="2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopSlash(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.259 3.034A1.001 1.001 0 0 0 12 3H4a1 1 0 0 0-1 1v8c0 .09.012.176.034.259zm.707.707l-9.225 9.225c.083.022.17.034.259.034h8a1 1 0 0 0 1-1V4c0-.09-.012-.176-.034-.259M4 2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Storage(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(0 2)"><path fill-rule="nonzero" d="M2 6a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1zm13 2v3a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1m1-3V2a2 2 0 0 0-2-2H2a2 2 0 0 0-2 2v3c0 .601.271 1.133.69 1.5C.271 6.867 0 7.399 0 8v3a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8c0-.601-.271-1.133-.689-1.5c.418-.367.689-.899.689-1.5"/><circle cx="4.5" cy="9.5" r="1.5"/><circle cx="4.5" cy="3.5" r="1.5"/><path d="M12 8h1v3h-1zm-2 0h1v3h-1zm2-6h1v3h-1zm-2 0h1v3h-1z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func String(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.297 3L8.93 5.102h1.351l-.32 1.828H8.609l-.656 3.883c-.036.265-.02.466.05.601c.071.135.247.208.528.219c.11.005.334-.008.672-.04L9.016 13.5a4.16 4.16 0 0 1-1.383.195c-.797-.01-1.393-.244-1.79-.703c-.395-.458-.557-1.08-.484-1.867l.688-4.195H5l.313-1.828h1.046L6.727 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Submodule(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" d="M6 2H1v12h3V7a1 1 0 0 1 1-1h5a1 1 0 0 1 1 1v1h4V4H7c-.621 0-1-.379-1-1zm10 6v6a1 1 0 0 1-1 1H1a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1h5.25a.75.75 0 0 1 .75.75l-.004.206C6.99 2.317 6.974 3 7 3h8a1 1 0 0 1 1 1zm-1 1h-4a1 1 0 0 1-1-1V7H5v7h10zM2 4.5a.5.5 0 0 1 .5-.5h2a.5.5 0 0 1 0 1h-2a.5.5 0 0 1-.5-.5M6.5 9a.5.5 0 0 0 0 1h2a.5.5 0 0 0 0-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwatchInput(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<rect width="12" height="12" x="2" y="2" fill="currentColor" rx="3"/><rect width="11" height="11" x="2.5" y="2.5" fill="currentColor" class="ouiSwatchInput__stroke" rx="2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Symlink(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" d="M10.8 0H2a1 1 0 0 0-1 1v8l1-1V1h8v3.5a.5.5 0 0 0 .5.5H14v10H2v-1a3.5 3.5 0 0 1 3.5-3.5H8V13l3-3l-3-3v2.5H5.5A4.5 4.5 0 0 0 1 14v1a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V4.429c0-.256-.098-.503-.274-.689l-3.2-3.428A1.002 1.002 0 0 0 10.8 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableDensityCompact(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3v11a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2zm-1 0V2a1 1 0 0 0-1-1H2a1 1 0 0 0-1 1v1zm0 1H1v10a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1zM4.496 7a.5.5 0 0 1 0 1H2.495a.5.5 0 0 1 0-1zm5.218 0c.158 0 .286.224.286.5s-.128.5-.286.5H6.286C6.128 8 6 7.776 6 7.5s.128-.5.286-.5zM4.496 5a.5.5 0 0 1 0 1H2.495a.5.5 0 0 1 0-1zm5.218 0c.158 0 .286.224.286.5s-.128.5-.286.5H6.286C6.128 6 6 5.776 6 5.5s.128-.5.286-.5zM4.496 9a.5.5 0 0 1 0 1H2.495a.5.5 0 0 1 0-1zm5.218 0c.158 0 .286.224.286.5s-.128.5-.286.5H6.286C6.128 10 6 9.776 6 9.5s.128-.5.286-.5zm-5.218 2a.5.5 0 0 1 0 1H2.495a.5.5 0 0 1 0-1zm5.218 0c.158 0 .286.224.286.5s-.128.5-.286.5H6.286C6.128 12 6 11.776 6 11.5s.128-.5.286-.5zm-5.218 2a.5.5 0 0 1 0 1H2.495a.5.5 0 0 1 0-1zm9-6a.5.5 0 0 1 0 1h-2.001a.5.5 0 0 1 0-1zm0-2a.5.5 0 0 1 0 1h-2.001a.5.5 0 0 1 0-1zm0 4a.5.5 0 0 1 0 1h-2.001a.5.5 0 0 1 0-1zm0 2a.5.5 0 0 1 0 1h-2.001a.5.5 0 0 1 0-1zm0 2a.5.5 0 0 1 0 1h-2.001a.5.5 0 0 1 0-1zm-3.782 0c.158 0 .286.224.286.5s-.128.5-.286.5H6.286C6.128 14 6 13.776 6 13.5s.128-.5.286-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableDensityExpanded(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3v11a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2zm-1 0V2a1 1 0 0 0-1-1H2a1 1 0 0 0-1 1v1zm0 1H1v10a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1zM4.5 7a.5.5 0 0 1 0 1H2.496a.5.5 0 1 1 0-1zm9 0a.5.5 0 1 1 0 1h-6a.5.5 0 0 1 0-1zm-9 4a.5.5 0 1 1 0 1H2.496a.5.5 0 1 1 0-1zm9 0a.5.5 0 1 1 0 1h-6a.5.5 0 1 1 0-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableDensityNormal(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3v11a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2zm-1 0V2a1 1 0 0 0-1-1H2a1 1 0 0 0-1 1v1zm0 1H1v10a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1zM4.5 6a.5.5 0 0 1 0 1H2.496a.5.5 0 1 1 0-1zm5.214 0c.158 0 .286.224.286.5s-.128.5-.286.5H6.286C6.128 7 6 6.776 6 6.5s.128-.5.286-.5zM4.5 9a.5.5 0 0 1 0 1H2.496a.5.5 0 1 1 0-1zm5.214 0c.158 0 .286.224.286.5s-.128.5-.286.5H6.286C6.128 10 6 9.776 6 9.5s.128-.5.286-.5zM4.5 12a.5.5 0 1 1 0 1H2.496a.5.5 0 1 1 0-1zm9-6a.5.5 0 1 1 0 1h-2.004a.5.5 0 0 1 0-1zm0 3a.5.5 0 1 1 0 1h-2.004a.5.5 0 0 1 0-1zm0 3a.5.5 0 1 1 0 1h-2.004a.5.5 0 0 1 0-1zm-3.786 0c.158 0 .286.224.286.5s-.128.5-.286.5H6.286C6.128 13 6 12.776 6 12.5s.128-.5.286-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableOfContents(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 1v14h14V1zM0 0h16v16H0zm9 1v14h1V1zM3 3.5h4v-1H3zm0 3h4v-1H3zm0 3h4v-1H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.254 14.97L.996 9.712c-.315-.316-.397-.463-.45-.64a.909.909 0 0 1 0-.534c.053-.177.135-.324.45-.64L7.43 1.466c.182-.183.252-.24.338-.293a.87.87 0 0 1 .273-.113c.099-.023.188-.032.446-.032h5.173c.445 0 .607.046.77.133c.162.087.29.214.377.377c.088.162.134.324.136.769l.015 5.15c0 .259-.009.348-.032.448a.87.87 0 0 1-.112.273c-.054.087-.111.157-.294.34L8.067 14.97c-.315.315-.462.396-.639.45a.909.909 0 0 1-.535 0c-.176-.054-.324-.135-.639-.45m1.106-.707l6.453-6.453c.092-.092.126-.128.141-.147c.003-.025.004-.074.004-.204l-.015-5.15a1.92 1.92 0 0 0-.009-.275a2.247 2.247 0 0 0-.274-.007H8.487c-.13 0-.179.001-.203.004c-.02.015-.055.05-.147.141L1.703 8.606a2.248 2.248 0 0 0-.189.2c.017.024.061.07.19.198l5.257 5.259c.128.128.175.171.2.188c.024-.017.071-.06.2-.188m4.972-10.607a2 2 0 1 1-2.828 2.828a2 2 0 0 1 2.828-2.828m-.707.707a1 1 0 1 0-1.414 1.414a1 1 0 0 0 1.414-1.414M6.807 11.28L4.686 9.159a.5.5 0 1 1 .707-.707l2.121 2.12a.5.5 0 1 1-.707.708m1.414-1.414l-2.12-2.122a.5.5 0 1 1 .706-.707L8.928 9.16a.5.5 0 1 1-.707.707"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tear(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.553 1.776a.5.5 0 0 1 .894 0l4 8a.5.5 0 0 1 .053.235c-.004.15-.026.396-.084.705a5.538 5.538 0 0 1-.505 1.512C11.189 13.638 9.906 14.5 8 14.5c-1.906 0-3.19-.862-3.91-2.272a5.538 5.538 0 0 1-.506-1.512a4.672 4.672 0 0 1-.084-.705a.5.5 0 0 1 .053-.235zm-3.046 8.329c.009.106.027.253.06.429c.079.424.213.848.413 1.238C5.537 12.862 6.487 13.5 8 13.5c1.513 0 2.463-.638 3.02-1.728c.2-.39.334-.814.413-1.238a4.35 4.35 0 0 0 .06-.43L8 3.119z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Temperature(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 15a3.5 3.5 0 0 1-1.75-6.532L7 8.324V2.5A1.496 1.496 0 0 1 9.908 2H8.5v1H10v1H8.5v1H10v1H8.5v1H10v1.324l.25.144A3.5 3.5 0 0 1 8.5 15M11 7.758V2.5a2.5 2.5 0 1 0-5 0v5.258a4.5 4.5 0 1 0 5 0"/><path fill="currentColor" d="M8.5 9a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsDown(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.509 13.649a.532.532 0 0 1-.845-.053l-.231-.347l-.539-1.078A2.118 2.118 0 0 0 7 11H6a.5.5 0 0 0 0 1h1c.423 0 .81.24 1 .618l.553 1.106a.5.5 0 0 0 .03.053l.25.373c.56.842 1.77.918 2.432.154a2.04 2.04 0 0 0 .492-1.466l-.008-.121a6.358 6.358 0 0 0-.313-1.588L11.393 11h1.773a1.835 1.835 0 0 0 1.6-2.736l-.118-.21a1.934 1.934 0 0 0-.491-1.904l-.095-.095A2.123 2.123 0 0 0 12.9 3.62l-.11-.33A1.887 1.887 0 0 0 11 2h-.921a5.07 5.07 0 0 0-2.268.535a7.63 7.63 0 0 0-1.983 1.43l-.182.181a.5.5 0 1 0 .707.707l.182-.181A6.631 6.631 0 0 1 8.258 3.43a4.07 4.07 0 0 1 1.82-.43H11c.382 0 .721.244.842.607l.179.536a.531.531 0 0 0 .336.336c.573.191.892.802.721 1.382l-.052.178a.535.535 0 0 0 .135.53l.289.288c.273.273.35.687.193 1.04a.5.5 0 0 0 .021.448l.23.41A.836.836 0 0 1 13.166 10h-1.96a.864.864 0 0 0-.82 1.138l.102.307c.144.433.233.883.263 1.338l.008.121c.018.272-.072.54-.25.745M4 4.647a.2.2 0 0 0-.2-.2H2.2a.2.2 0 0 0-.2.2v6.706c0 .11.09.2.2.2h1.6a.2.2 0 0 0 .2-.2zM2 3.5c-.552 0-1 .424-1 .947v7.106c0 .523.448.947 1 .947h2c.552 0 1-.424 1-.947V4.447c0-.523-.448-.947-1-.947z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsUp(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.509 2.351a.532.532 0 0 0-.845.053l-.231.347l-.539 1.078A2.118 2.118 0 0 1 7 5H6a.5.5 0 0 1 0-1h1c.423 0 .81-.238 1-.617l.553-1.106a.5.5 0 0 1 .03-.053l.25-.373a1.532 1.532 0 0 1 2.432-.154c.35.405.527.932.492 1.466l-.008.121a6.358 6.358 0 0 1-.313 1.588L11.393 5h1.773a1.835 1.835 0 0 1 1.6 2.736l-.118.21c.193.67.012 1.401-.491 1.904l-.095.095c.246.99-.25 2.009-1.161 2.434l-.11.33A1.888 1.888 0 0 1 11 14h-.921a5.071 5.071 0 0 1-2.268-.535a7.63 7.63 0 0 1-1.983-1.43l-.182-.181a.5.5 0 0 1 .707-.708l.182.182a6.63 6.63 0 0 0 1.723 1.242a4.07 4.07 0 0 0 1.82.43H11a.887.887 0 0 0 .842-.607l.179-.536a.532.532 0 0 1 .336-.336c.573-.191.892-.802.721-1.382l-.052-.178a.535.535 0 0 1 .135-.53l.289-.288a.934.934 0 0 0 .193-1.04a.5.5 0 0 1 .021-.448l.23-.41A.836.836 0 0 0 13.166 6h-1.96a.864.864 0 0 1-.82-1.138l.102-.307a5.36 5.36 0 0 0 .263-1.338l.008-.121a1.032 1.032 0 0 0-.25-.745M4 4.647a.2.2 0 0 0-.2-.2H2.2a.2.2 0 0 0-.2.2v6.706c0 .11.09.2.2.2h1.6a.2.2 0 0 0 .2-.2zM2 3.5c-.552 0-1 .424-1 .947v7.106c0 .523.448.947 1 .947h2c.552 0 1-.424 1-.947V4.447c0-.523-.448-.947-1-.947z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Timeline(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 4.5a.5.5 0 0 0 1 0V4h1a1 1 0 0 0 1-1V1a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h1zM9 1H6v2h3zM2 7.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0M2.5 9a1.5 1.5 0 0 1-1.415-1H.5a.5.5 0 0 1 0-1h.585a1.5 1.5 0 0 1 2.83 0h2.17a1.5 1.5 0 0 1 2.83 0h2.17a1.5 1.5 0 0 1 2.83 0h.585a.5.5 0 0 1 0 1h-.585a1.5 1.5 0 0 1-2.83 0h-2.17a1.5 1.5 0 0 1-2.83 0h-2.17A1.5 1.5 0 0 1 2.5 9M13 7.5a.5.5 0 1 0-1 0a.5.5 0 0 0 1 0m-5 0a.5.5 0 1 0-1 0a.5.5 0 0 0 1 0M2.5 10a.5.5 0 0 0-.5.5v.5H1a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1H3v-.5a.5.5 0 0 0-.5-.5M4 14v-2H1v2zm8-3.5a.5.5 0 0 1 1 0v.5h1a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-3a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h1zm2 2.5v1h-3v-2h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Timeslider(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" d="M13.923 8A5.93 5.93 0 0 0 8 2.077A5.93 5.93 0 0 0 2.077 8a5.93 5.93 0 0 0 6.296 5.912c.328-.021.625.227.625.556a.504.504 0 0 1-.458.511a7 7 0 1 1 6.43-6.333c-.036.398-.487.58-.817.354a.595.595 0 0 1-.248-.54c.012-.152.018-.305.018-.46m1.684 3.2l-4.32-3.055c-.56-.396-1.287.059-1.287.8v6.108c0 .74.726 1.196 1.287.8l4.32-3.055c.524-.37.524-1.228 0-1.598M7.462 7.462H4.769a.539.539 0 0 0 0 1.076H8A.539.539 0 0 0 8.538 8V3.692a.539.539 0 0 0-1.076 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenAlias(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.075 6.953a.5.5 0 1 1-.707.707a1.5 1.5 0 0 0-2.122 0L4.125 9.782a1.5 1.5 0 1 0 2.121 2.121l1.145-1.144a.5.5 0 0 1 .707.707L6.953 12.61a2.5 2.5 0 1 1-3.535-3.535l2.121-2.122a2.5 2.5 0 0 1 3.536 0m3.535-3.535a2.5 2.5 0 0 1 0 3.535l-2.12 2.122a2.5 2.5 0 0 1-3.536 0a.5.5 0 1 1 .707-.708a1.5 1.5 0 0 0 2.122 0l2.121-2.12a1.5 1.5 0 1 0-2.121-2.122L8.637 5.269a.5.5 0 1 1-.707-.707l1.145-1.144a2.5 2.5 0 0 1 3.535 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenAnnotation(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.15 3.392c2.797 0 4.524 1.644 4.517 4.289c.007 1.816-.708 2.893-2.21 3.004c-.908.076-1.081-.287-1.157-.725h-.041c-.163.42-.964.732-1.744.683c-1.053-.065-2.082-.842-2.09-2.572c.008-1.72 1.071-2.441 1.959-2.586c.804-.135 1.598.158 1.723.462h.051v-.386h1.195v3.452c.007.3.128.425.304.425c.4 0 .677-.583.673-1.861c.004-2.376-1.705-2.914-3.187-2.914c-2.34 0-3.415 1.522-3.422 3.387c.007 2.127 1.22 3.277 3.433 3.277c.808 0 1.598-.176 2.006-.349l.393 1.122c-.435.27-1.419.508-2.493.508c-2.98 0-4.723-1.66-4.727-4.496c.004-2.804 1.748-4.72 4.817-4.72M7.964 6.79c-.76 0-1.185.459-1.188 1.24c.003.683.3 1.332 1.202 1.332c.821 0 1.094-.473 1.077-1.343c-.004-.718-.204-1.23-1.091-1.23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenArray(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.392 12V4h2.713v1.14h-1.21v5.72h1.21V12zm7.692-8v8H9.37v-1.14h1.209V5.14H9.37V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenBinary(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 4H4v8h8zM8.5 5.5h-3v5h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenBoolean(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 13A5 5 0 1 1 8 3a5 5 0 0 1 0 10m-2.828-2.172a4 4 0 0 1 5.656-5.656c.004.013-5.645 5.674-5.656 5.656"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenClass(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.333 7.027H9.375c-.056-.708-.48-1.187-1.222-1.187c-.972 0-1.5.806-1.5 2.16c0 1.43.545 2.16 1.486 2.16c.708 0 1.139-.415 1.236-1.08l1.958.015C11.236 10.418 10.181 12 8.097 12c-1.958 0-3.43-1.41-3.43-4c0-2.6 1.514-4 3.43-4c1.792 0 3.084 1.095 3.236 3.027"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenCompletionSuggester(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v5.996a1 1 0 0 1-1 1h-1.661L7.4 13.2A.25.25 0 0 1 7 13v-2.004H4a1 1 0 0 1-1-1zm1.5 1a.5.5 0 0 1 .5-.5h6a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-.5.5h-.9L8 11V9.5H5a.5.5 0 0 1-.5-.5z" clip-rule="evenodd"/><path fill="currentColor" d="M6.75 6a1 1 0 1 0 0 2a1 1 0 0 0 0-2m2.5 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenConstant(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m9.414 3.757l2.829 2.829a2 2 0 0 1 0 2.828l-2.829 2.829a2 2 0 0 1-2.828 0L3.757 9.414a2 2 0 0 1 0-2.828l2.829-2.829a2 2 0 0 1 2.828 0m-1.747 2.91a1 1 0 0 0-1 1v.666a1 1 0 0 0 1 1h.666a1 1 0 0 0 1-1v-.666a1 1 0 0 0-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenDate(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13 11.567C13 12.36 12.36 13 11.567 13H4.433C3.64 13 3 12.36 3 11.567V4.433C3 3.64 3.64 3 4.433 3H6v-.495a.51.51 0 0 1 .412-.497L6.5 2c.276 0 .5.214.5.505V3h2v-.495a.51.51 0 0 1 .412-.497L9.5 2c.276 0 .5.214.5.505V3h1.567C12.36 3 13 3.64 13 4.433zM4 6v5.33c0 .37.3.67.67.67h6.66c.37 0 .67-.3.67-.67V6zm1.5 4c.245 0 .45.183.492.412L6 10.5c0 .245-.183.45-.412.492L5.5 11a.505.505 0 0 1-.5-.5c0-.245.183-.45.412-.492zM8 10c.245 0 .45.183.492.412l.008.088c0 .245-.183.45-.412.492L8 11a.505.505 0 0 1-.5-.5c0-.245.183-.45.412-.492zm2.5 0c.245 0 .45.183.492.412L11 10.5c0 .245-.183.45-.412.492L10.5 11a.505.505 0 0 1-.5-.5c0-.245.183-.45.412-.492zm-5-1.5c.245 0 .45.183.492.412L6 9c0 .245-.183.45-.412.492L5.5 9.5A.505.505 0 0 1 5 9c0-.245.183-.45.412-.492zm2.5 0c.245 0 .45.183.492.412L8.5 9c0 .245-.183.45-.412.492L8 9.5a.505.505 0 0 1-.5-.5c0-.245.183-.45.412-.492zm2.5 0c.245 0 .45.183.492.412L11 9c0 .245-.183.45-.412.492L10.5 9.5A.505.505 0 0 1 10 9c0-.245.183-.45.412-.492zM5.5 7c.245 0 .45.183.492.412L6 7.5c0 .245-.183.45-.412.492L5.5 8a.505.505 0 0 1-.5-.5c0-.245.183-.45.412-.492zM8 7c.245 0 .45.183.492.412L8.5 7.5c0 .245-.183.45-.412.492L8 8a.505.505 0 0 1-.5-.5c0-.245.183-.45.412-.492zm2.5 0c.245 0 .45.183.492.412L11 7.5c0 .245-.183.45-.412.492L10.5 8a.505.505 0 0 1-.5-.5c0-.245.183-.45.412-.492zM4 5h8v-.33c0-.37-.3-.67-.67-.67H4.67C4.3 4 4 4.3 4 4.67z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenDenseVector(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.154 12V4h2.713v1.14H5.658v5.72h1.21V12zm7.692-8v8H9.133v-1.14h1.209V5.14h-1.21V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenElement(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7.39 9.736l-1.041.94L3.258 8l3.09-2.677l1.041.94l-2.032 1.722v.03zm2.777.94l-1.04-.94l2.032-1.721v-.03L9.126 6.264l1.04-.94L13.259 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenEnum(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.211 12V4h5.578v1.57H7.145v1.641h3.359v1.574H7.145v1.645h3.644V12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenEnumMember(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.682 12V9.247l1.023-.861a.5.5 0 0 0-.003-.768l-1.02-.844V4h5.578v1.57H9.615v1.64h3.36v1.575h-3.36v1.645h3.645V12zm.743-4.103a.138.138 0 0 1 0 .206L6.158 9.97a.133.133 0 0 1-.218-.103v-.934H2.873A.133.133 0 0 1 2.74 8.8V7.2c0-.074.06-.133.133-.133H5.94v-.934a.133.133 0 0 1 .218-.103z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenEvent(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.225 5.656c0 .423-.106.79-.318 1.102c-.211.311-.51.57-.898.775a5.435 5.435 0 0 1-1.392.485c-.54.117-1.14.193-1.798.229a6.047 6.047 0 0 0-.035.67c0 .258.02.51.062.757c.04.247.114.464.22.652s.25.34.432.458s.414.176.696.176c.211 0 .467-.044.766-.132c.3-.088.62-.244.96-.467c.106-.07.192-.129.256-.176a.365.365 0 0 1 .22-.07c.118 0 .197.061.238.185a.99.99 0 0 1 .062.255a7.1 7.1 0 0 1-.573.467a4.93 4.93 0 0 1-.775.467c-.288.141-.6.261-.934.361c-.335.1-.678.15-1.03.15c-.541 0-.982-.088-1.322-.264a2.072 2.072 0 0 1-.793-.688a2.626 2.626 0 0 1-.388-.933a4.949 4.949 0 0 1-.106-1.005c0-.634.103-1.257.309-1.868c.205-.61.499-1.157.88-1.638c.383-.482.838-.87 1.366-1.163A3.567 3.567 0 0 1 9.093 4c.599 0 1.104.126 1.515.379c.411.252.617.678.617 1.277m-2.467-.951c-.223 0-.435.08-.635.238c-.2.158-.381.373-.546.643c-.164.27-.305.578-.423.925a6.42 6.42 0 0 0-.264 1.101c.47-.047.863-.135 1.18-.264c.318-.13.57-.285.758-.467c.188-.182.323-.388.405-.617c.083-.229.124-.467.124-.713c0-.27-.056-.479-.168-.626a.519.519 0 0 0-.431-.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenException(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.461 7.31h3.055a.74.74 0 0 1 .66 1.074l-2.141 4.211a.74.74 0 1 1-1.319-.67L9.31 8.79H6.256a.74.74 0 0 1-.66-1.075l2.19-4.31a.74.74 0 0 1 1.319.67z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenField(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.27 12V4h5.46v1.57H7.203v1.641h3.18v1.574h-3.18V12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenFile(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.867 2.667H4a.667.667 0 0 0-.667.666v9.334c0 .368.299.666.667.666h8a.667.667 0 0 0 .667-.666V5.619a.669.669 0 0 0-.183-.459l-2.133-2.285a.668.668 0 0 0-.484-.208m1.466 4V12H4.667V4h4v2.333c0 .184.149.334.333.334z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenFlattened(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.25 3a.25.25 0 0 0-.25.25v2c0 .138.112.25.25.25h9.5a.25.25 0 0 0 .25-.25v-2a.25.25 0 0 0-.25-.25zm0 3.75A.25.25 0 0 0 3 7v2c0 .138.112.25.25.25H5.5A.25.25 0 0 0 5.75 9V7a.25.25 0 0 0-.25-.25zm-.25 4a.25.25 0 0 1 .25-.25H5.5a.25.25 0 0 1 .25.25v2a.25.25 0 0 1-.25.25H3.25a.25.25 0 0 1-.25-.25zm3.31-.727c-.082-.073-.082-.224 0-.296l3.054-2.683a.17.17 0 0 1 .19-.026c.064.032.104.1.104.174v1.341l3.161-.016c.1 0 .18.086.18.192v2.3c0 .105-.08.191-.18.191l-3.161.017v1.341c0 .074-.04.142-.103.174a.17.17 0 0 1-.19-.025z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenFunction(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.266 4.15v1.48a2.41 2.41 0 0 0-.859-.145c-.692 0-1.065.346-1.177 1.043l-.11.625h1.828v1.44H8.921l-.204 1.115C8.455 11.325 7.517 12 5.9 12c-.469 0-.882-.061-1.166-.167v-1.495c.273.117.591.178.903.178c.659 0 1.01-.29 1.127-1.015l.157-.91H5.247V7.152h1.837l.188-.842C7.534 4.714 8.432 4 10.19 4c.39 0 .853.067 1.076.15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenGeo(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 3a4.99 4.99 0 0 1 4 2c.628.836 1 1.875 1 3a4.978 4.978 0 0 1-.999 3H12a4.992 4.992 0 0 1-4 2a4.992 4.992 0 0 1-4-1.999V11a5 5 0 0 1 4-8m.948 8H7.052c.277.626.623 1 .948 1c.325 0 .67-.374.948-1M6 11l-.645.001c.274.242.581.446.914.606A5.445 5.445 0 0 1 6 11.001m4.645.001H10a5.51 5.51 0 0 1-.269.606c.333-.16.64-.364.914-.606m-5.133-2.5H4.031a3.98 3.98 0 0 0 .505 1.5h1.172a9.186 9.186 0 0 1-.196-1.5m3.975 0H6.513c.03.544.104 1.05.21 1.5h2.553c.107-.45.182-.956.21-1.5m2.482 0h-1.481a9.186 9.186 0 0 1-.196 1.5h1.172a3.98 3.98 0 0 0 .505-1.5M5.708 6H4.535c-.261.452-.437.96-.504 1.5h1.481A9.187 9.187 0 0 1 5.708 6m3.568 0H6.724a8.357 8.357 0 0 0-.21 1.499h2.973a8.479 8.479 0 0 0-.21-1.5M11.465 6h-1.173c.102.467.17.972.196 1.5h1.481a3.974 3.974 0 0 0-.504-1.5M6.269 4.393l-.124.062c-.286.15-.551.333-.79.545H6a5.51 5.51 0 0 1 .269-.607M8 4c-.326 0-.671.375-.948 1h1.896C8.671 4.376 8.326 4 8 4m1.73.393l.038.071c.083.168.161.347.232.536h.646a4.006 4.006 0 0 0-.915-.607"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenHistogram(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h4v5h2V5.5h4V13H3.001v-1H3zm1 9h2V4H4zm3 0h2V9H7zm3 0h2V6.5h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenInterface(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.967 10.653h1.727V12H5.306v-1.347h1.727V5.347H5.306V4h5.388v1.347H8.967z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenIp(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 3a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2zm-1 2H8v6h1V9.014h1c.298-.013 2 0 2-2.018c0-1.74-1.314-1.952-1.825-1.987zM6 5H5v6h1zm4 .984c.667 0 1 .336 1 1.008C11 7.664 10.667 8 10 8H9V5.984Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenJoin(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 4.5v1.025c0 1.269-1.185 1.908-2.112 1.737a.75.75 0 1 0 0 1.475c.927-.17 2.112.47 2.112 1.739v1.023h4v-1.005a2.5 2.5 0 0 1 0-4.988V4.5zM13 4a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v1.525c0 .172-.172.293-.341.262a2.25 2.25 0 1 0 0 4.426c.17-.031.341.09.341.262V12a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1V9.004a.16.16 0 0 0-.04-.105c-.109-.125-.594-.16-.732-.068a1 1 0 1 1 0-1.662c.138.092.623.057.732-.068a.16.16 0 0 0 .04-.105z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenKey(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.667 6.542A3.208 3.208 0 0 1 8.86 9.694l-.438.492a.437.437 0 0 1-.327.147h-.678v.73a.437.437 0 0 1-.438.437H6.25v.73a.437.437 0 0 1-.437.437H3.77a.437.437 0 0 1-.438-.438v-1.423a.44.44 0 0 1 .128-.31l2.95-2.949a3.208 3.208 0 0 1 3.047-4.214a3.202 3.202 0 0 1 3.209 3.209m-3.209-.875a.875.875 0 1 0 1.75 0a.875.875 0 0 0-1.75 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenKeyword(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.75 7.375a.25.25 0 0 0-.25.25v.75c0 .138.112.25.25.25h3.5a.25.25 0 0 0 .25-.25v-.75a.25.25 0 0 0-.25-.25z"/><path fill="currentColor" fill-rule="evenodd" d="M3 5a1 1 0 0 1 1-1h5.989a1 1 0 0 1 .825.436l2.05 3a1 1 0 0 1 0 1.128l-2.05 3A1 1 0 0 1 9.99 12H4a1 1 0 0 1-1-1zm1.25.75a.5.5 0 0 1 .5-.5h4.745a.5.5 0 0 1 .405.206l1.636 2.25a.5.5 0 0 1 0 .588L9.9 10.544a.5.5 0 0 1-.405.206H4.75a.5.5 0 0 1-.5-.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenMethod(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.333 11.027V5.05h2.059v1.136h.063c.25-.747.891-1.214 1.728-1.214c.848 0 1.524.483 1.65 1.214h.063c.204-.731.927-1.214 1.822-1.214c1.155 0 1.949.798 1.949 2.023v4.03h-2.169V7.542c0-.521-.29-.84-.738-.84s-.723.319-.723.84v3.486H6.963V7.54c0-.521-.29-.84-.739-.84c-.447 0-.722.319-.722.84v3.486z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenModule(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8 2l5.196 3v.178l-.866.468V5.5L8 3L3.67 5.5v5L8 13l4.33-2.5V5.77l.866-.474V11L8 14l-5.196-3V5z"/><path fill="currentColor" d="M5.243 4.429L9.597 7.04L8 7.928L3.743 5.563a.5.5 0 1 0-.486.874L7.5 8.794V13.5h1V8.794l4.243-2.357a.508.508 0 0 0 .06-.04l.392-.202V5.047l-.917.505a.573.573 0 0 0-.02.01l-.106.06l-.191.105l-1.355.753l-4.849-2.909z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenNamespace(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.556 8.416l-.804-1.68h-.036v5.64H5V4h1.992l2.292 3.96l.804 1.68h.036V4h1.716v8.376H9.848z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenNested(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M11 3c1.044 0 1.913.757 1.994 1.736l.006.149v6.23c0 1-.82 1.805-1.845 1.88L11 13H9.501a.5.5 0 0 1-.09-.992l.09-.008H11c.52 0 .937-.35.993-.783l.007-.102v-6.23c0-.445-.379-.827-.882-.879L11 4H9.5a.5.5 0 0 1-.09-.992L9.5 3zM6.5 3a.5.5 0 0 1 .09.992L6.5 4H5c-.52 0-.937.35-.993.783L4 4.885v6.23c0 .445.379.827.882.879L5 12h1.5a.5.5 0 0 1 .09.992L6.5 13H5c-1.044 0-1.913-.757-1.994-1.736L3 11.115v-6.23c0-1 .82-1.805 1.845-1.88L5 3z"/><path d="M5.864 7.25a.714.714 0 1 1 0 1.429a.714.714 0 0 1 0-1.429m2.143 0a.714.714 0 1 1 0 1.429a.714.714 0 0 1 0-1.429m2.143 0a.714.714 0 1 1 0 1.429a.714.714 0 0 1 0-1.429"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenNull(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.002 12.412l-.962.962a1 1 0 0 1-1.414-1.414l.962-.962a5.333 5.333 0 0 1 7.41-7.41l.962-.962a1 1 0 1 1 1.414 1.414l-.962.962a5.333 5.333 0 0 1-7.41 7.41m.966-.966a4 4 0 0 0 5.478-5.478zm-1.414-1.414l5.478-5.478a4 4 0 0 0-5.478 5.478"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenNumber(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.8 9.812h-.842l-.78 2.335H5.031l.78-2.335H4.6l.182-1.043h1.373l.507-1.504H5.454l.188-1.042h1.367l.792-2.37H8.94l-.792 2.37h.849l.792-2.37h1.145l-.792 2.37H11.4l-.182 1.042H9.8L9.293 8.77h1.248l-.183 1.043H8.946l-.775 2.335H7.026zm-.5-1.043h.842l.513-1.504h-.849z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenObject(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.63 12c-1.294 0-2.383-.105-2.383-1.802V9.6c0-.638-.247-.914-.98-.914v-1.37c.733 0 .98-.28.98-.915v-.6C4.247 4.105 5.336 4 6.631 4v1.14c-.759 0-.886.272-.886.843v.813c0 .479-.225.936-1.212 1.133v.142c.987.197 1.212.654 1.212 1.133v.813c0 .57.127.844.886.844zm2.266-8c1.295 0 2.384.105 2.384 1.802V6.4c0 .638.247.914.98.914v1.37c-.733 0-.98.28-.98.915v.6c0 1.696-1.09 1.801-2.384 1.801v-1.14c.759 0 .886-.272.886-.843v-.813c0-.479.225-.936 1.212-1.133V7.93c-.987-.197-1.212-.654-1.212-1.133v-.813c0-.57-.127-.844-.886-.844z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenOperator(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.667 8.002c0-1.5.394-2.743 1.248-3.822h1.437c-.652.8-1.14 2.388-1.14 3.822c0 1.43.488 3.018 1.14 3.818H3.915c-.854-1.08-1.248-2.322-1.248-3.818M6.77 9.998l-.818-.803l1.23-1.197l-1.23-1.203l.83-.793l1.221 1.193L9.23 6.002l.818.793l-1.227 1.2l1.227 1.2l-.818.803L8 8.795zm6.563-2c0 1.5-.394 2.743-1.248 3.822h-1.437c.652-.8 1.14-2.388 1.14-3.822c0-1.43-.488-3.018-1.14-3.818h1.437c.854 1.08 1.248 2.322 1.248 3.818"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenPackage(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m8.049 3.785l3.852 1.006l-4.049 1.103L4 4.791l3.951-1.006a.19.19 0 0 1 .098 0m.073 2.654l4.545-1.306v5.45l-.131.184l-4.414 1.455zm-4.789 4.145V5.188L7.498 6.41v5.81l-4.034-1.453z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenParameter(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.889 12V4h3.304c1.797 0 2.922 1.117 2.918 2.77c.004 1.652-1.144 2.746-2.976 2.746H6.822V12zm1.933-4.008h.953c.868 0 1.336-.472 1.332-1.222c.004-.73-.464-1.211-1.332-1.211h-.953z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenPercolator(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.5 7a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0 1a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5M11 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2m0 1a2 2 0 1 0 0-4a2 2 0 0 0 0 4m-2.5 1.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0m1 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenProperty(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.889 12V4h3.304c1.797 0 2.922 1.117 2.918 2.77c.004 1.652-1.144 2.746-2.976 2.746H6.822V12zm1.933-4.008h.953c.868 0 1.336-.472 1.332-1.222c.004-.73-.464-1.211-1.332-1.211h-.953z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenRange(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.82 5.116a.5.5 0 0 0-.704.704l.064.064L11.719 8l-2.54 2.116a.5.5 0 0 0-.114.63l.05.074a.5.5 0 0 0 .63.115l.075-.05l3-2.5a.5.5 0 0 0 .071-.697l-.07-.072zm-3.64 0a.5.5 0 0 1 .704.704l-.064.064L4.281 8l2.54 2.116a.5.5 0 0 1 .114.63l-.05.074a.5.5 0 0 1-.63.115l-.075-.05l-3-2.5a.5.5 0 0 1-.071-.697l.07-.072z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenRankFeature(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 4v8h-2V6H9v4H5v2H3V8h4V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenRankFeatures(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 4v8h-2V6H9v4H5v2H3V8h4V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenRepo(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.533 9.067c-1.792 0-2.378.72-2.57 1.194a1.601 1.601 0 1 1-1.163-.037V5.776a1.595 1.595 0 0 1-1.067-1.51c0-.885.715-1.6 1.6-1.6c.886 0 1.6.715 1.6 1.6c0 .7-.442 1.291-1.066 1.51v2.821C6.336 8.251 7.019 8 8 8c1.424 0 1.899-.715 2.053-1.19a1.603 1.603 0 0 1-.986-1.477c0-.885.714-1.6 1.6-1.6c.885 0 1.6.715 1.6 1.6a1.59 1.59 0 0 1-1.115 1.526c-.139.762-.656 2.208-2.619 2.208m-3.2 2.133a.535.535 0 0 0-.533.533c0 .294.24.534.533.534a.535.535 0 0 0 0-1.067m0-7.467a.535.535 0 0 0-.533.534c0 .293.24.533.533.533c.294 0 .534-.24.534-.533a.535.535 0 0 0-.534-.534M10.667 4.8a.535.535 0 0 0-.534.533a.535.535 0 0 0 1.067 0a.535.535 0 0 0-.533-.533"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenSearchType(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.27 7.5a1 1 0 1 1 2 0a1 1 0 0 1-2 0m2.5 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0"/><path fill="currentColor" fill-rule="evenodd" d="M11.702 10.682a4.501 4.501 0 0 1-5.796.482L4.28 12.789a.75.75 0 0 1-1.06-1.06L4.847 10.1a4.501 4.501 0 1 1 6.855.581m-5.304-1.06a3 3 0 1 0 4.243-4.243A3 3 0 0 0 6.398 9.62" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenShape(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13 10v3h-3v-1H6v1H3v-3h1V6H3V3h3v1h4V3h3v3h-1v4zm-8 1H4v1h1zm7 0h-1v1h1zM5 4H4v1h1zm7 0h-1v1h1zm-1 2h-1V5H6v1H5v4h1v1h4v-1h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenString(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.147 4.297l-.255 1.455h.936l-.222 1.266h-.935l-.455 2.688c-.025.184-.013.323.036.417c.048.093.17.144.365.151c.075.004.23-.005.465-.027l-.13 1.32a2.89 2.89 0 0 1-.957.135c-.552-.007-.965-.17-1.239-.487c-.274-.317-.386-.748-.335-1.293l.476-2.904h-.725l.216-1.266h.725l.254-1.455z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenStruct(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.336 4.667h2.667v2.666H4.336zm0 4h2.667v2.666H4.336zm4-4h2.667v2.666H8.336zm0 4h2.667v2.666H8.336zM3.003 3.333v9.334h9.333V3.333zm0-1.333h9.333c.737 0 1.334.597 1.334 1.333v9.334c0 .736-.597 1.333-1.334 1.333H3.003a1.333 1.333 0 0 1-1.333-1.333V3.333C1.67 2.597 2.267 2 3.003 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenSymbol(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.316 14a6 6 0 1 1 0-12a6 6 0 0 1 0 12m0-1.333a4.667 4.667 0 1 0 0-9.334a4.667 4.667 0 0 0 0 9.334m2.19-5.72h1.143c.019 1.448-.793 2.338-1.922 2.338c-.632 0-1.194-.267-1.706-.811c-.36-.397-.636-.576-1-.576c-.517 0-.849.355-.885 1.083H4.983c.014-1.47.858-2.314 1.95-2.314c.595 0 1.125.249 1.678.802c.392.382.641.595 1.038.595c.484 0 .857-.323.857-1.116"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenText(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.147 4.297l-.255 1.455h.936l-.222 1.266h-.935l-.455 2.688c-.025.184-.013.323.036.417c.048.093.17.144.365.151c.075.004.23-.005.465-.027l-.13 1.32a2.89 2.89 0 0 1-.957.135c-.552-.007-.965-.17-1.239-.487c-.274-.317-.386-.748-.335-1.293l.476-2.904h-.725l.216-1.266h.725l.254-1.455z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenTokenCount(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 4H3v3h5zm5 5H8v3h5zm-3-5h3v3h-3zM6 9H3v3h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TokenVariable(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.649 4.667l1.326 4.7h.05l1.326-4.7h1.982l-2.134 6.666H6.801L4.667 4.667z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Training(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.386 9.836a2.5 2.5 0 1 1 3.611.667C15.212 11.173 16 12.46 16 14v1.5a.5.5 0 1 1-1 0V14c0-1.724-1.276-3-3-3c-.91 0-1.298-.02-1.805-.122c-1.25-.254-2.333-1-3.585-2.566a.5.5 0 1 1 .78-.624c.9 1.124 1.653 1.74 2.434 2.043c.155.052.345.083.562.105m1.785.128c.083.01.167.021.251.034L12.5 10a1.5 1.5 0 1 0-.33-.036M9.78 11.97a.5.5 0 0 1 .5.5a.992.992 0 0 1-.05.231c-.179.38-.23.774-.23 1.302v1.5a.5.5 0 1 1-1 0v-1.5c0-.657.072-1.186.307-1.696a.5.5 0 0 1 .473-.337M5.958 5.772a.5.5 0 0 1-.78.625L3.11 3.812a.5.5 0 1 1 .78-.624zM1 11h5.5a.5.5 0 1 1 0 1h-6a.5.5 0 0 1-.5-.5V.5A.5.5 0 0 1 .5 0h12a.5.5 0 0 1 .5.5v3a.5.5 0 1 1-1 0V1H1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 3h5v1H0V3h5V1a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1zm-7.056 8H7v1H4.1l.392 2.519c.042.269.254.458.493.458h6.03c.239 0 .451-.189.493-.458l1.498-9.576H14l-1.504 9.73c-.116.747-.74 1.304-1.481 1.304h-6.03c-.741 0-1.365-.557-1.481-1.304l-1.511-9.73H9V5.95H3.157L3.476 8H8v1H3.632zM6 3h4V1H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Undeploy(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.18 7c-.473 0-.88.294-.972.703l-1.189 5.25a.776.776 0 0 0-.019.172c0 .483.444.875.99.875h11.02c.065 0 .13-.006.194-.017c.537-.095.885-.556.778-1.03l-1.19-5.25C13.7 7.294 13.293 7 12.822 7zM6 6v1h5V6h1.825c.946 0 1.76.606 1.946 1.447l1.19 5.4c.215.975-.482 1.923-1.556 2.118a2.175 2.175 0 0 1-.39.035H2.985C1.888 15 1 14.194 1 13.2c0-.118.013-.237.039-.353l1.19-5.4C2.414 6.606 3.229 6 4.174 6z"/><path fill="currentColor" fill-rule="evenodd" d="M5.477 2.618a.5.5 0 0 1 .705.059L8.55 5.476l2.368-2.799a.5.5 0 0 1 .764.646L9.205 6.25l2.477 2.927a.5.5 0 1 1-.764.646L8.55 7.024L6.182 9.823a.5.5 0 1 1-.764-.646L7.895 6.25L5.418 3.323a.5.5 0 0 1 .059-.705" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unfold(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.14 3.808L7.53 1.51a.67.67 0 0 1 .942 0l2.389 2.298c.26.256.685.256.944 0a.65.65 0 0 0 0-.93L9.416.578a2.02 2.02 0 0 0-2.832 0l-2.388 2.3a.65.65 0 0 0 0 .93c.26.256.683.256.944 0m0 8.384l2.39 2.298a.67.67 0 0 0 .942 0l2.389-2.298a.677.677 0 0 1 .944 0a.65.65 0 0 1 0 .93l-2.388 2.3a2.02 2.02 0 0 1-2.832 0l-2.388-2.3a.65.65 0 0 1 0-.93a.677.677 0 0 1 .944 0zM16 6H0v4h16zM1 9V7h14v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlink(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><path d="M7.565 10.83a.5.5 0 1 1 .819.573l-1.877 2.68c-.963 1.376-3.02 1.55-4.588.45C.35 13.436-.191 11.445.773 10.07L3.64 5.973c.963-1.375 3.019-1.548 4.587-.45a.5.5 0 1 1-.573.82C6.509 5.54 5.08 5.66 4.46 6.546l-2.868 4.095c-.621.887-.245 2.27.9 3.073c1.146.802 2.575.682 3.196-.205z"/><path d="M7.892 3.43a.5.5 0 1 1-.574-.819L10 .734c1.376-.963 3.367-.422 4.465 1.146c1.098 1.569.926 3.625-.45 4.588L9.918 9.336c-1.375.963-3.366.422-4.464-1.146a.5.5 0 1 1 .819-.574c.802 1.146 2.185 1.522 3.072.9L13.44 5.65c.886-.621 1.006-2.05.204-3.195c-.802-1.146-2.186-1.522-3.072-.9zM6 .5v3a.5.5 0 0 1-1 0v-3a.5.5 0 0 1 1 0M1.058 2.23l2.458 1.72a.5.5 0 0 1-.574.82L.484 3.05a.5.5 0 1 1 .574-.82m12 7.093l2.457 1.72a.5.5 0 1 1-.573.82l-2.457-1.72a.5.5 0 1 1 .573-.82m-2.099 1.181v3a.5.5 0 1 1-1 0v-3a.5.5 0 1 1 1 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func User(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M13.689 11.132c1.155 1.222 1.953 2.879 2.183 4.748a1.007 1.007 0 0 1-1 1.12H3.007a1.005 1.005 0 0 1-1-1.12c.23-1.87 1.028-3.526 2.183-4.748c.247.228.505.442.782.633c-1.038 1.069-1.765 2.55-1.972 4.237L14.872 16c-.204-1.686-.93-3.166-1.966-4.235a7.01 7.01 0 0 0 .783-.633M8.939 1c1.9 0 3 2 4.38 2.633a2.483 2.483 0 0 1-1.88.867c-.298 0-.579-.06-.844-.157A3.726 3.726 0 0 1 7.69 5.75c-1.395 0-3.75.25-3.245-1.903C5.94 3 6.952 1 8.94 1"/><path d="M8.94 2c2.205 0 4 1.794 4 4s-1.795 4-4 4c-2.207 0-4-1.794-4-4s1.793-4 4-4m0 9A5 5 0 1 0 8.937.999A5 5 0 0 0 8.94 11"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Users(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.482 4.344a2 2 0 1 0-2.963 0c-.08.042-.156.087-.23.136c-.457.305-.75.704-.933 1.073A3.457 3.457 0 0 0 1 6.978V9a1 1 0 0 0 1 1h2.5a3.69 3.69 0 0 1 .684-.962L5.171 9H2V7s0-2 2-2c1.007 0 1.507.507 1.755 1.01c.225-.254.493-.47.793-.636a2.717 2.717 0 0 0-1.066-1.03M4 4a1 1 0 1 0 0-2a1 1 0 0 0 0 2m10 6h-2.5a3.684 3.684 0 0 0-.684-.962L10.829 9H14V7s0-2-2-2c-1.007 0-1.507.507-1.755 1.01a3.012 3.012 0 0 0-.793-.636a2.716 2.716 0 0 1 1.066-1.03a2 2 0 1 1 2.963 0c.08.042.156.087.23.136c.457.305.75.704.933 1.073A3.453 3.453 0 0 1 15 6.944V9a1 1 0 0 1-1 1m-2-6a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path fill="currentColor" fill-rule="evenodd" d="M10 8c0 .517-.196.989-.518 1.344a2.755 2.755 0 0 1 1.163 1.21A3.453 3.453 0 0 1 11 11.977V14a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-2.022a2.005 2.005 0 0 1 .006-.135a3.456 3.456 0 0 1 .35-1.29a2.755 2.755 0 0 1 1.162-1.21A2 2 0 1 1 10 8m-4 4v2h4v-2s0-2-2-2s-2 2-2 2m3-4a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vector(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 11V5H11V3.5H5V5H3.5v6H5v1.5h6V11zm1 0H15v4h-4v-1.5H5V15H1v-4h1.5V5H1V1h4v1.5h6V1h4v4h-1.5zM4 4V2H2v2zm8 0h2V2h-2zM2 14h2v-2H2zm10 0h2v-2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoPlayer(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1.994C0 .893.895 0 1.994 0h12.012C15.107 0 16 .895 16 1.994v12.012A1.995 1.995 0 0 1 14.006 16H1.994A1.995 1.995 0 0 1 0 14.006zm1 0v12.012c0 .548.446.994.994.994h12.012a.995.995 0 0 0 .994-.994V1.994A.995.995 0 0 0 14.006 1H1.994A.995.995 0 0 0 1 1.994M1 4h14v1H1zm1.5-1a.5.5 0 0 1 0-1h1a.5.5 0 0 1 0 1zm3 0a.5.5 0 0 1 0-1h1a.5.5 0 0 1 0 1zm4.947 6.106a1 1 0 0 1 0 1.788l-3 2A1 1 0 0 1 6 12V8a1 1 0 0 1 1.447-.894zM10 10L7 8v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisArea(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 13h10V9.913l-2.571-2.826L8.56 8.753a.5.5 0 0 1-.728-.067L4.448 4.317L3 6.191zm5.295-5.35l1.837-1.64a.5.5 0 0 1 .703.037l3.035 3.336a.5.5 0 0 1 .13.337v3.78a.5.5 0 0 1-.5.5h-11a.5.5 0 0 1-.5-.5V6.02a.5.5 0 0 1 .104-.305l1.947-2.52a.5.5 0 0 1 .791-.001zM1 15h13.5a.5.5 0 1 1 0 1H.5a.5.5 0 0 1-.5-.5v-14a.5.5 0 0 1 1 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisAreaStacked(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.5 1a.5.5 0 0 1 .5.5V15h13.5a.5.5 0 1 1 0 1H.5a.5.5 0 0 1-.5-.5v-14A.5.5 0 0 1 .5 1m4.342 2.194L8.295 7.65l1.837-1.64a.5.5 0 0 1 .703.037l3.035 3.336a.5.5 0 0 1 .13.337v3.78a.5.5 0 0 1-.5.5h-11a.5.5 0 0 1-.5-.5V6.02a.5.5 0 0 1 .104-.305l1.947-2.52a.5.5 0 0 1 .791-.001m-.394 1.123L3 6.191v4.101l1.146-1.146a.5.5 0 0 1 .493-.126l.085.033L8.5 10.94l1.776-.888a.5.5 0 0 1 .36-.034l.088.034L13 11.19V9.913l-2.571-2.826L8.56 8.753a.5.5 0 0 1-.728-.067z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisBarHorizontal(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 10h-6a.5.5 0 0 1 0-1H8V6H2.5a.5.5 0 0 1 0-1H13V2H2.5a.5.5 0 0 1 0-1h11a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-.5.5H9v3h2.5a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-.5.5h-9a.5.5 0 1 1 0-1H11v-3zM0 .5a.5.5 0 1 1 1 0v14a.5.5 0 1 1-1 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisBarHorizontalStacked(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.5 0a.5.5 0 0 1 .5.5v14a.5.5 0 1 1-1 0V.5A.5.5 0 0 1 .5 0m13 1a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-.5.5H9v3h2.5a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-.5.5h-9a.5.5 0 1 1 0-1H9v-3H2.5a.5.5 0 0 1 0-1H6V6H2.5a.5.5 0 0 1 0-1H10V2H2.5a.5.5 0 0 1 0-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisBarVertical(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 7.5v6a.5.5 0 1 1-1 0V8H6v5.5a.5.5 0 1 1-1 0V3H2v10.5a.5.5 0 1 1-1 0v-11a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5V7h3V4.5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5v9a.5.5 0 1 1-1 0V5h-3zM.5 16a.5.5 0 1 1 0-1h14a.5.5 0 1 1 0 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisBarVerticalStacked(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 15a.5.5 0 1 1 0 1H.5a.5.5 0 1 1 0-1zm-9-13a.5.5 0 0 1 .5.5L5.999 7H9V4.5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5v9a.5.5 0 1 1-1 0V7h-3v6.5a.5.5 0 0 1-.41.492L9.5 14a.5.5 0 0 1-.5-.5V10H6v3.5a.5.5 0 0 1-.992.09L5 13.5V6H2v7.5a.5.5 0 1 1-1 0v-11a.5.5 0 0 1 .5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisBuilder(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9 15H5v-1h4zm8-7h-4V7h4zm3 7h-9v-1h9zm8-7h-9V7h9z" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M10 9V3a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1h-7v4h-1v-7H3v9h9v1H3a1 1 0 0 1-1-1v-9a1 1 0 0 1 1-1zm1-6h18v9h-7v-2a1 1 0 0 0-1-1H11z" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M19.714 26.286L20.571 28h.858v-.888c0-.34.296-.976.602-1.54c.14-.259.281-.502.398-.696c.154-.257.264-.427.264-.427c.246-.408.45-1.117.45-1.588v-4.046c0-.796-.605-1.445-1.364-1.48a.728.728 0 0 0-.78.74v.295a1.518 1.518 0 0 0-.698-1.282a1.461 1.461 0 0 0-.73-.23a.778.778 0 0 0-.714.772v.296a1.481 1.481 0 0 0-.704-1.277a1.383 1.383 0 0 0-.724-.205a.728.728 0 0 0-.715.741v.296a1.482 1.482 0 0 0-.702-1.276a1.383 1.383 0 0 0-.726-.205a.728.728 0 0 0-.715.74v4.149h-.857v-2.667a1.668 1.668 0 0 0-1 .334A1.8 1.8 0 0 0 12 20v3.089c0 .478.307 1.084.686 1.365c0 0 2.743 1.956 2.743 2.658V28h3.428zm-.857-5.143v3.428h.857v-3.428zm-.857 0h-.857v3.428H18zm2.572 0v3.428h.857v-3.428zm-8.467 4.126l-.008-.006l-.008-.006a2.74 2.74 0 0 1-.781-.962A2.743 2.743 0 0 1 11 23.089V20c0-1.453 1.109-2.698 2.571-2.774v-.485c0-.927.734-1.741 1.715-1.741c.59 0 1.12.21 1.533.556a1.7 1.7 0 0 1 .61-.112c.584 0 1.11.207 1.522.547c.17-.07.351-.115.54-.13l.06-.006l.06.003c.572.023 1.091.24 1.499.585c.186-.07.39-.11.604-.11c1.376 0 2.429 1.146 2.429 2.482v4.046c0 .358-.074.75-.17 1.091a4.276 4.276 0 0 1-.425 1.015l-.007.012m-.008.012l-.002.003l-.01.015l-.036.058a17.848 17.848 0 0 0-.547.93a8.94 8.94 0 0 0-.39.79c-.076.18-.104.282-.114.32c-.005.02-.005.02-.005.005V29h-2.476l-.239-.478l-.239.478H14.43v-1.724a2.434 2.434 0 0 0-.168-.202a9.35 9.35 0 0 0-.847-.789a22.704 22.704 0 0 0-1.283-.997l-.02-.015l-.006-.004m2.372 2.082l-.008-.014z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisBuilderSavedObject(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 2c0 .022 0 .043-.002.064l1.559 1.04a1 1 0 1 1-.555.832l-1.559-1.04a.996.996 0 0 1-.702.07L7.466 4.741A1.001 1.001 0 0 1 6.5 6a1 1 0 0 1-.998-1.064l-1.559-1.04a1 1 0 1 1 .555-.832l1.559 1.04a.996.996 0 0 1 .702-.07l1.775-1.775A1.001 1.001 0 0 1 9.5 1a1 1 0 0 1 1 1"/><path fill="currentColor" d="M8.5 13a.5.5 0 0 1-1 0V8.5h-2V13a.5.5 0 0 1-1 0V7.5h-2V13a.5.5 0 0 1-1 0V7.2a.7.7 0 0 1 .7-.7h2.6a.7.7 0 0 1 .7.7v.3h2V6.2a.7.7 0 0 1 .7-.7h2.6a.7.7 0 0 1 .7.7V7h2.3a.7.7 0 0 1 .7.7V13a.5.5 0 0 1-1 0V8h-2v5a.5.5 0 0 1-1 0V6.5h-2zm-8 1.5A.5.5 0 0 1 1 14h14a.5.5 0 0 1 0 1H1a.5.5 0 0 1-.5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisGauge(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.877 5.847l-1.02 1.02a.5.5 0 0 1-.708-.707l1.1-1.099c-.05-.053-.1-.106-.152-.157A6.471 6.471 0 0 0 8 3.019V4.5a.5.5 0 0 1-1 0V3.019a6.47 6.47 0 0 0-4.261 2.055l1.07 1.071a.5.5 0 0 1-.706.707l-.99-.99A6.46 6.46 0 0 0 1.018 10H2.5a.5.5 0 1 1 0 1H1.174c.083.353.196.697.337 1.03a.5.5 0 1 1-.922.39A7.487 7.487 0 0 1 0 9.5a7.483 7.483 0 0 1 2.197-5.304A7.487 7.487 0 0 1 7.5 2a7.487 7.487 0 0 1 5.304 2.197A7.483 7.483 0 0 1 15 9.5a7.487 7.487 0 0 1-.59 2.92a.5.5 0 0 1-.92-.39a6.49 6.49 0 0 0 .336-1.03H12.5a.5.5 0 1 1 0-1h1.481a6.483 6.483 0 0 0-1.104-4.153m-6.041 5.317a.993.993 0 0 1-.01-1.404c.384-.385 2.882-2.002 3.149-1.735c.267.267-1.35 2.765-1.735 3.15a.993.993 0 0 1-1.404-.01"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisGoal(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.725 3.653a6 6 0 0 1 2.847 7.576a.5.5 0 0 1-.928-.372a5 5 0 1 0-9.293-.014a.5.5 0 0 1-.218.619L1.39 12.47a.5.5 0 0 1-.708-.23A7.99 7.99 0 0 1 0 9a8 8 0 0 1 11.212-7.329a.5.5 0 0 1 .234.704zm-.933-.38l.5-.889a7 7 0 0 0-8.902 8.93l.886-.511a6 6 0 0 1 7.516-7.53M6.73 9.467a1.75 1.75 0 1 1 2.539 0a2 2 0 1 1-2.539 0M8 12.013a1 1 0 1 0 0-2a1 1 0 0 0 0 2m0-3a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisLine(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.654 3.48c.248.225.552.389.888.467L11.24 9.43a1.99 1.99 0 0 0-.915-.404zM9.146 9.19a2.008 2.008 0 0 0-.769.64l-1.572-2a1.99 1.99 0 0 0 .785-.618zM5.581 7.956l-2.134 4.268a.5.5 0 0 1-.894-.448l2.134-4.268c.25.22.557.376.894.448M1 15h13.5a.5.5 0 1 1 0 1H.5a.5.5 0 0 1-.5-.5v-14a.5.5 0 0 1 1 0zm5-8a1 1 0 1 1 0-2a1 1 0 0 1 0 2m4 5a1 1 0 1 1 0-2a1 1 0 0 1 0 2m4-9a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisMapCoordinate(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 15C5.77 15 2 9.97 2 6.595C2 3.507 4.46 1 7.5 1S13 3.507 13 6.595C13 9.97 9.23 15 7.5 15m0-1c.286 0 1.48-1.044 2.459-2.35C11.219 9.969 12 8.153 12 6.596C12 4.055 9.983 2 7.5 2S3 4.055 3 6.595c0 1.557.78 3.373 2.041 5.056C6.02 12.956 7.214 14 7.5 14m0-4.996a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5m0-1a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisMapRegion(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.5 2.309v8.69a.499.499 0 0 1-.032.176L9.5 12.691V3.809zm-1-.04L2 3.825v8.906l3.527-1.568a.5.5 0 0 1-.027-.164zm.274-1.216a.498.498 0 0 1 .471.01l3.768 1.884l4.284-1.904A.5.5 0 0 1 15 1.5v10a.5.5 0 0 1-.297.457l-4.5 2a.5.5 0 0 1-.427-.01l-3.789-1.894l-4.283 1.904a.5.5 0 0 1-.703-.457v-10a.5.5 0 0 1 .297-.457zM10.5 3.825v8.906l3.5-1.556V2.27z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisMetric(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.532 7.34a2.161 2.161 0 1 1 2.936 0a2.746 2.746 0 1 1-2.936 0M2 0h12a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2m0 1a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1zm6 5.915a1.161 1.161 0 1 0 0-2.322a1.161 1.161 0 0 0 0 2.322m0 4.492a1.746 1.746 0 1 0 0-3.492a1.746 1.746 0 0 0 0 3.492"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisPie(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.5 9a.5.5 0 0 1-.5-.5V3.023A5.5 5.5 0 1 0 11.978 9zM7 8h5.5a.5.5 0 0 1 .5.5A6.5 6.5 0 1 1 6.5 2a.5.5 0 0 1 .5.5zm2-6.972V6h4.972C13.696 3.552 11.448 1.304 9 1.028M14.5 7h-6a.5.5 0 0 1-.5-.5v-6a.5.5 0 0 1 .5-.5C11.853 0 15 3.147 15 6.5a.5.5 0 0 1-.5.5M6.146 8.854a.5.5 0 1 1 .708-.708l4 4a.5.5 0 0 1-.708.708z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisQueryDql(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 2c0 .022 0 .043-.002.064l1.559 1.04a1 1 0 1 1-.555.832l-1.559-1.04a.996.996 0 0 1-.702.07L7.466 4.741A1.001 1.001 0 0 1 6.5 6a1 1 0 0 1-.998-1.064l-1.559-1.04a1 1 0 1 1 .555-.832l1.559 1.04a.996.996 0 0 1 .702-.07l1.775-1.775A1.001 1.001 0 0 1 9.5 1a1 1 0 0 1 1 1"/><path fill="currentColor" d="M8.5 6.5h2v8a.5.5 0 0 0 1 0V8h2v6.5a.5.5 0 0 0 1 0V7.7a.7.7 0 0 0-.7-.7h-2.3v-.8a.7.7 0 0 0-.7-.7H8.2a.7.7 0 0 0-.7.7v1.3h-2v-.3a.7.7 0 0 0-.7-.7H2.2a.7.7 0 0 0-.7.7V9a.5.5 0 0 0 1 0V7.5h2V9a.5.5 0 0 0 1 0v-.5h2V9a.5.5 0 0 0 1 0z"/><path fill="currentColor" fill-rule="evenodd" d="M1.6 10a.6.6 0 0 0-.6.6v3.86a.6.6 0 0 0 .6.6h7.476a.6.6 0 0 0 .6-.6V10.6a.6.6 0 0 0-.6-.6zm5.545 3.398v-2.675h.579v2.675c0 .16.13.29.289.29h1.229v.578h-1.23a.868.868 0 0 1-.867-.868m-.867.336v-2.071a.868.868 0 0 0-.868-.868h-.289a.868.868 0 0 0-.868.868v1.735c0 .48.389.868.868.868h.868c.093 0 .215.051.336.132a1.43 1.43 0 0 1 .173.136l.008.008h.001l.205-.204l.204-.205h-.001l-.002-.002l-.005-.005l-.015-.015a2.046 2.046 0 0 0-.247-.195a1.486 1.486 0 0 0-.368-.182m-1.157-2.36a.29.29 0 0 0-.29.289v1.735c0 .16.13.29.29.29h.578v-2.025a.29.29 0 0 0-.289-.29zm-3.398 2.892v-3.47h1.012c.56 0 1.012.453 1.012 1.011v1.446c0 .56-.453 1.013-1.012 1.013zm.578-2.892v2.313h.434c.24 0 .434-.194.434-.434v-1.446a.434.434 0 0 0-.434-.433z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisQueryPpl(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 2a.98.98 0 0 1-.002.064l1.559 1.04a1 1 0 1 1-.554.832l-1.56-1.04a.996.996 0 0 1-.702.07L7.466 4.741A1.001 1.001 0 0 1 6.5 6a1 1 0 0 1-.998-1.064l-1.558-1.04a1 1 0 1 1 .555-.832l1.558 1.04a.996.996 0 0 1 .702-.07l1.775-1.775A1.001 1.001 0 0 1 9.5 1a1 1 0 0 1 1 1"/><path fill="currentColor" d="M8.5 9a.5.5 0 0 1-1 0v-.5h-2V9a.5.5 0 0 1-1 0V7.5h-2V9a.5.5 0 0 1-1 0V7.2a.7.7 0 0 1 .7-.7h2.6a.7.7 0 0 1 .7.7v.3h2V6.2a.7.7 0 0 1 .7-.7h2.6a.7.7 0 0 1 .7.7V7h2.3a.7.7 0 0 1 .7.7v6.8a.5.5 0 0 1-1 0V8h-2v6.5a.5.5 0 0 1-1 0v-8h-2z"/><path fill="currentColor" fill-rule="evenodd" d="M1.6 10a.6.6 0 0 0-.6.6v3.86a.6.6 0 0 0 .6.6h7.476a.6.6 0 0 0 .6-.6V10.6a.6.6 0 0 0-.6-.6zm5.184.723h.578v2.964H8.88v.579H6.784zm-2.53.072H5.41c.48 0 .868.389.868.868v.65c0 .48-.389.868-.868.868h-.578v1.157h-.579zm.578 1.808h.578c.16 0 .29-.13.29-.29v-.65a.29.29 0 0 0-.29-.29h-.578zm-3.109 1.735v-3.543H2.88c.479 0 .867.389.867.868v.65c0 .48-.388.868-.867.868H2.3v1.157zm.578-1.735h.579a.29.29 0 0 0 .289-.29v-.65a.29.29 0 0 0-.29-.29h-.578z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisQueryPromql(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 2c0 .022 0 .043-.002.064l1.559 1.04a1 1 0 1 1-.555.832l-1.559-1.04a.996.996 0 0 1-.702.07L7.466 4.741A1.001 1.001 0 0 1 6.5 6a1 1 0 0 1-.998-1.064l-1.559-1.04a1 1 0 1 1 .555-.832l1.559 1.04a.996.996 0 0 1 .702-.07l1.775-1.775A1.001 1.001 0 0 1 9.5 1a1 1 0 0 1 1 1"/><path fill="currentColor" d="M8.5 9a.5.5 0 0 1-1 0v-.5h-2V9a.5.5 0 0 1-1 0V7.5h-2V9a.5.5 0 0 1-1 0V7.2a.7.7 0 0 1 .7-.7h2.6a.7.7 0 0 1 .7.7v.3h2V6.2a.7.7 0 0 1 .7-.7h2.6a.7.7 0 0 1 .7.7V7h2.3a.7.7 0 0 1 .7.7V9a.5.5 0 0 1-1 0V8h-2v1a.5.5 0 0 1-1 0V6.5h-2z"/><path fill="currentColor" fill-rule="evenodd" d="M1.6 10c-.331 0-.6.315-.6.704v3.592c0 .389.269.704.6.704h12.8c.331 0 .6-.315.6-.704v-3.592c0-.389-.269-.704-.6-.704zm10.965.714v2.643c0 .474.327.857.73.857h1.035v-.571h-1.034c-.135 0-.244-.128-.244-.286v-2.643zm-.73.929v2.047c.124.046.231.118.31.18a1.749 1.749 0 0 1 .207.192l.014.014l.004.005l.001.002h.001l-.172.203l-.172.202v-.001l-.008-.008a1.15 1.15 0 0 0-.145-.134c-.102-.08-.205-.13-.284-.13h-.73c-.404 0-.73-.384-.73-.858v-1.714c0-.474.326-.857.73-.857h.243c.404 0 .73.383.73.857m-1.218 0c0-.158.11-.286.244-.286h.243c.135 0 .244.128.244.286v2h-.487c-.135 0-.244-.128-.244-.286zm-9.008-.857v3.5h.487v-1.143h.487c.403 0 .73-.384.73-.857v-.643c0-.474-.327-.857-.73-.857zm.974 1.785h-.487v-1.214h.487c.134 0 .243.128.243.286v.643c0 .158-.109.285-.243.285m1.643 1.715H3.74v-2.5h.487v.167a.746.746 0 0 1 .365-.096h.305v.572H4.59c-.201 0-.365.191-.365.428zm3.652 0v-1.668l.498.584a.218.218 0 0 0 .344 0l.497-.584v1.668h.487v-2.012c0-.382-.393-.574-.623-.304l-.533.626l-.533-.626c-.23-.27-.624-.078-.624.303v2.013zM5.99 11.857c-.403 0-.73.384-.73.857v.643c0 .474.327.857.73.857h.244c.403 0 .73-.383.73-.857v-.643c0-.473-.327-.857-.73-.857zm-.243.857c0-.158.109-.285.243-.285h.244c.134 0 .243.127.243.285v.643c0 .158-.109.286-.243.286H5.99c-.134 0-.243-.128-.243-.286z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisQuerySql(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 2c0 .022 0 .043-.002.064l1.559 1.04a1 1 0 1 1-.555.832l-1.559-1.04a.996.996 0 0 1-.702.07L7.466 4.741A1.001 1.001 0 0 1 6.5 6a1 1 0 0 1-.998-1.064l-1.559-1.04a1 1 0 1 1 .555-.832l1.559 1.04a.996.996 0 0 1 .702-.07l1.775-1.775A1.001 1.001 0 0 1 9.5 1a1 1 0 0 1 1 1"/><path fill="currentColor" d="M8.5 9a.5.5 0 0 1-1 0v-.5h-2V9a.5.5 0 0 1-1 0V7.5h-2V9a.5.5 0 0 1-1 0V7.2a.7.7 0 0 1 .7-.7h2.6a.7.7 0 0 1 .7.7v.3h2V6.2a.7.7 0 0 1 .7-.7h2.6a.7.7 0 0 1 .7.7V7h2.3a.7.7 0 0 1 .7.7v6.8a.5.5 0 0 1-1 0V8h-2v6.5a.5.5 0 0 1-1 0v-8h-2z"/><path fill="currentColor" fill-rule="evenodd" d="M1.6 10a.6.6 0 0 0-.6.6v3.86a.6.6 0 0 0 .6.6h7.476a.6.6 0 0 0 .6-.6V10.6a.6.6 0 0 0-.6-.6zm5.618.723v2.675c0 .48.388.868.867.868h1.23v-.579h-1.23a.29.29 0 0 1-.289-.289v-2.675zm-5.423 1.084c0-.559.453-1.012 1.012-1.012h1.085v.579H2.807a.434.434 0 1 0 0 .867a1.012 1.012 0 0 1 0 2.025H1.723v-.579h1.084a.434.434 0 1 0 0-.867a1.012 1.012 0 0 1-1.012-1.013m4.555-.144v2.071c.148.047.275.12.368.182a2.01 2.01 0 0 1 .247.195l.016.015l.005.005l.001.001l.001.001l-.204.205l-.205.205v-.001l-.009-.008a1.427 1.427 0 0 0-.173-.136c-.12-.081-.242-.132-.336-.132h-.868a.868.868 0 0 1-.867-.868v-1.735c0-.48.388-.868.867-.868h.29c.479 0 .867.389.867.868m-1.446 0c0-.16.13-.29.29-.29h.288c.16 0 .29.13.29.29v2.024h-.579a.29.29 0 0 1-.289-.289z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisTable(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3v11a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2zm-1 0V2a1 1 0 0 0-1-1H2a1 1 0 0 0-1 1v1zm0 1H1v10a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1zM4.5 6a.5.5 0 0 1 0 1H2.496a.5.5 0 1 1 0-1zm9 0a.5.5 0 1 1 0 1h-6a.5.5 0 0 1 0-1zm-9 3a.5.5 0 0 1 0 1H2.496a.5.5 0 1 1 0-1zm9 0a.5.5 0 1 1 0 1h-6a.5.5 0 0 1 0-1zm-9 3a.5.5 0 1 1 0 1H2.496a.5.5 0 1 1 0-1zm9 0a.5.5 0 1 1 0 1h-6a.5.5 0 1 1 0-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisTagCloud(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 9.047a.5.5 0 1 0 0 1h13a.5.5 0 0 0 0-1zm0-1h13a1.5 1.5 0 0 1 0 3h-13a1.5 1.5 0 0 1 0-3M10 13a.5.5 0 1 1 0 1H4a.5.5 0 1 1 0-1zM8.001 2.015a.5.5 0 1 1-.002 1l-5-.015a.5.5 0 1 1 .003-1zM14 5a.5.5 0 1 1 0 1H6a.5.5 0 0 1 0-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisText(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 5v6h2a.5.5 0 1 1 0 1h-5a.5.5 0 1 1 0-1h2V5H5v.5a.5.5 0 0 1-1 0v-1a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 .5.51l-.021 1a.5.5 0 1 1-1-.02l.01-.49zM1 15h1.5a.5.5 0 1 1 0 1h-2a.5.5 0 0 1-.5-.5v-1.996a.5.5 0 0 1 1 0zM1 1v1.497a.5.5 0 1 1-1 0V.5A.5.5 0 0 1 .5 0h2a.5.5 0 0 1 0 1zm14 0h-1.495a.5.5 0 0 1 0-1H15.5a.5.5 0 0 1 .5.5v2a.5.5 0 1 1-1 0zm0 14v-1.5a.5.5 0 1 1 1 0v2a.5.5 0 0 1-.5.5h-2a.5.5 0 1 1 0-1zM0 6.5a.5.5 0 0 1 1 0v3a.5.5 0 0 1-1 0zM9.5 0a.5.5 0 0 1 0 1h-3a.5.5 0 0 1 0-1zM15 6.5a.5.5 0 1 1 1 0v3a.5.5 0 1 1-1 0zM9.5 15a.5.5 0 1 1 0 1h-3a.5.5 0 1 1 0-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisTimelion(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.884 1.802c1.171-.309 2.22-.24 3.103.07c.295.104.507.212.631.292l.164.106l.048.188c.278 1.066.2 2.175-.091 3.23a6.728 6.728 0 0 1-.194.598c.035.12.074.273.112.454c.227 1.065.247 2.193-.08 3.267a4.855 4.855 0 0 1-.534 1.175a6.1 6.1 0 0 1-1.117 1.323c-.806.712-1.616 1.071-2.235 1.206c-.14.179-.346.381-.629.578c-.643.447-1.493.711-2.561.711c-1.07 0-1.92-.264-2.563-.711a2.874 2.874 0 0 1-.629-.578c-.618-.134-1.428-.493-2.232-1.203a6.05 6.05 0 0 1-1.12-1.326a4.855 4.855 0 0 1-.534-1.175C.096 8.933.116 7.805.343 6.74a6.74 6.74 0 0 1 .112-.454a6.728 6.728 0 0 1-.194-.598C-.03 4.633-.108 3.524.17 2.458l.048-.188l.164-.106c.124-.08.336-.188.63-.291c.885-.31 1.933-.38 3.104-.07c.184-.108.434-.23.748-.348C5.616 1.17 6.497 1 7.501 1c1.002 0 1.883.17 2.634.455c.315.119.565.24.75.347m3.027 1.116a3.13 3.13 0 0 0-.255-.102c-.749-.263-1.66-.309-2.708.007l-.238.072l-.201-.147c-.11-.08-.354-.217-.728-.358C9.14 2.147 8.38 2 7.501 2c-.88 0-1.641.147-2.283.39c-.373.141-.617.278-.727.358l-.2.147l-.24-.072c-1.048-.316-1.958-.27-2.707-.007a3.13 3.13 0 0 0-.255.102c-.167.812-.094 1.674.136 2.503c.084.303.166.524.217.637l.082.182l-.066.19c-.032.09-.084.27-.137.518c-.195.914-.212 1.88.059 2.767c.1.333.241.645.423.935c.273.434.588.801.936 1.108c.434.383.898.655 1.36.834c.27.105.474.155.575.17l.248.034l.12.22c.051.091.199.266.466.452c.476.33 1.127.532 1.993.532c.865 0 1.516-.202 1.99-.532c.268-.186.416-.36.466-.452l.121-.22l.248-.034c.101-.015.305-.065.576-.17c.462-.18.927-.452 1.362-.836c.347-.307.66-.674.933-1.107c.182-.29.322-.601.423-.934c.271-.887.254-1.853.06-2.767a4.597 4.597 0 0 0-.138-.519l-.066-.189l.082-.182c.05-.113.133-.334.217-.637c.23-.829.303-1.691.136-2.503M5 5.008c.004.181-.048.364-.264.616c-.428.498-.985.488-1.431.049c-.326-.321-.325-.732-.29-1.047c.085-.76.748-.722 1.262-.44c.486.27.717.616.723.823m5 0c.006-.206.237-.553.724-.821c.514-.283 1.176-.321 1.26.44c.036.314.038.725-.288 1.046c-.446.44-1.005.449-1.433-.05c-.217-.25-.267-.434-.263-.614M7.5 8c-.503 0-.976-.24-1.375-.672c-.398-.432-.592-1.126 0-1.528c.384-.259 1.082-.3 1.375-.3c.294 0 .991.041 1.375.3c.593.402.397 1.096 0 1.528C8.477 7.76 8.002 8 7.5 8m0 1c.808 0 2.656.994 2.49 2.074c-.153.993-2.114.925-2.49.925c-.374 0-2.336.068-2.49-.925C4.845 9.994 6.693 9 7.5 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisVega(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.414 8.036L4.89 10.51a.5.5 0 0 1-.707.708L1.354 8.389a.5.5 0 0 1 0-.707l2.828-2.828a.5.5 0 1 1 .707.707zm8.768 2.474l2.475-2.474l-2.475-2.475a.5.5 0 0 1 .707-.707l2.829 2.828a.5.5 0 0 1 0 .707l-2.829 2.829a.5.5 0 1 1-.707-.708M8.559 2.506a.5.5 0 0 1 .981.19L7.441 13.494a.5.5 0 0 1-.981-.19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisVisualBuilder(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.837 7c.11.93.165 1.886.165 2.869V13.5a.5.5 0 1 1-1 0V9.869A23.3 23.3 0 0 0 8.83 7H7.29c-.195 1.04-.292 1.985-.292 2.835V13.5a.5.5 0 1 1-1 0V9.835c0-.864.092-1.809.276-2.835H2.5a.5.5 0 0 1-.495-.57c.285-2.023 1.626-3.358 3.931-3.96c1.967-.514 4.22-.606 6.756-.278A1.5 1.5 0 0 1 14 3.679V5.5A1.5 1.5 0 0 1 12.5 7zm-.569-1H12.5a.5.5 0 0 0 .5-.5V3.68a.5.5 0 0 0-.436-.497c-2.416-.311-4.54-.225-6.375.254C4.494 3.88 3.491 4.724 3.117 6h6.15zM2 10v3.5a.5.5 0 1 1-1 0v-4a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 .5.5v4a.5.5 0 1 1-1 0V10zm10 3.5a.5.5 0 1 1-1 0v-2a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 .5.5v2a.5.5 0 1 1-1 0V12h-2zM1.016 16.026a.5.5 0 0 1 0-1H15a.5.5 0 1 1 0 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WordWrap(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 3h12v1H2zm0 8h6v1H2z"/><path fill="currentColor" d="M2 7h9.5v.5V7h.039l.083.005a2.958 2.958 0 0 1 1.102.298c.309.154.633.394.88.763c.248.373.396.847.396 1.434c0 .588-.148 1.061-.396 1.434a2.257 2.257 0 0 1-.88.763a2.957 2.957 0 0 1-1.185.302h-.025l-.009.001h-.003s-.002 0-.002-.5v.5H11v1l-2-1.5l2-1.5v1h.506l.044-.003a1.959 1.959 0 0 0 .726-.195c.191-.095.367-.23.495-.423c.127-.19.229-.466.229-.879s-.102-.689-.229-.879a1.256 1.256 0 0 0-.495-.424a1.958 1.958 0 0 0-.77-.197H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WordWrapDisabled(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 3.5L12 2v1H2v1h10v1zM12 9V8H2V7h10V6l2 1.5zm0 3v1l2-1.5l-2-1.5v1H2v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wrench(children ...ElementRenderer) *OuiIcon {
	return &OuiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.918 9.746l4.537 4.537a2 2 0 1 0 2.828-2.829l-3.157-3.156a.5.5 0 0 1 .708-.708l3.156 3.157a3 3 0 1 1-4.243 4.243l-4.949-4.95a5.001 5.001 0 0 1-5.22-7.106a.5.5 0 0 1 .805-.138L3.676 5.09a1 1 0 1 0 1.415-1.414L2.797 1.382a.5.5 0 0 1 .138-.805a5.001 5.001 0 1 1 3.983 9.169M1.226 4.054a4.002 4.002 0 0 0 6.693 3.865a4 4 0 0 0-3.865-6.693l1.744 1.743a2 2 0 1 1-2.829 2.828zm10.229 8.814a1 1 0 1 1 1.414-1.414a1 1 0 0 1-1.414 1.414"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
