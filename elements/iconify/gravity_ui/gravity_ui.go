package gravity_ui

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "2.9.1"
	hAttr          = "1em"
	viewbox        = "0 0 0 0"
	fill           = "currentColor"
)

type GravityUiIcon struct {
	*SVGSVGElement
}

type GravityUiIconFn func(children ...ElementRenderer) *GravityUiIcon

var IconLookup = map[string]GravityUiIconFn{
	"abbrApi":                           AbbrApi,
	"abbrMl":                            AbbrMl,
	"abbrQl":                            AbbrQl,
	"abbrSql":                           AbbrSql,
	"abbrZip":                           AbbrZip,
	"antennaSignal":                     AntennaSignal,
	"aperture":                          Aperture,
	"archive":                           Archive,
	"arrowDown":                         ArrowDown,
	"arrowDownFromLine":                 ArrowDownFromLine,
	"arrowDownToLine":                   ArrowDownToLine,
	"arrowDownToSquare":                 ArrowDownToSquare,
	"arrowLeft":                         ArrowLeft,
	"arrowLeftFromLine":                 ArrowLeftFromLine,
	"arrowLeftToLine":                   ArrowLeftToLine,
	"arrowRight":                        ArrowRight,
	"arrowRightArrowLeft":               ArrowRightArrowLeft,
	"arrowRightFromLine":                ArrowRightFromLine,
	"arrowRightFromSquare":              ArrowRightFromSquare,
	"arrowRightToLine":                  ArrowRightToLine,
	"arrowRightToSquare":                ArrowRightToSquare,
	"arrowRotateLeft":                   ArrowRotateLeft,
	"arrowRotateRight":                  ArrowRotateRight,
	"arrowShapeDown":                    ArrowShapeDown,
	"arrowShapeDownFromLine":            ArrowShapeDownFromLine,
	"arrowShapeDownToLine":              ArrowShapeDownToLine,
	"arrowShapeLeft":                    ArrowShapeLeft,
	"arrowShapeLeftFromLine":            ArrowShapeLeftFromLine,
	"arrowShapeLeftToLine":              ArrowShapeLeftToLine,
	"arrowShapeRight":                   ArrowShapeRight,
	"arrowShapeRightFromLine":           ArrowShapeRightFromLine,
	"arrowShapeRightToLine":             ArrowShapeRightToLine,
	"arrowShapeTurnUpLeft":              ArrowShapeTurnUpLeft,
	"arrowShapeTurnUpRight":             ArrowShapeTurnUpRight,
	"arrowShapeUp":                      ArrowShapeUp,
	"arrowShapeUpFromLine":              ArrowShapeUpFromLine,
	"arrowShapeUpToLine":                ArrowShapeUpToLine,
	"arrowUp":                           ArrowUp,
	"arrowUpArrowDown":                  ArrowUpArrowDown,
	"arrowUpFromLine":                   ArrowUpFromLine,
	"arrowUpFromSquare":                 ArrowUpFromSquare,
	"arrowUpFromSquareSlash":            ArrowUpFromSquareSlash,
	"arrowUpRightFromSquare":            ArrowUpRightFromSquare,
	"arrowUpToLine":                     ArrowUpToLine,
	"arrowUturnCcwDown":                 ArrowUturnCcwDown,
	"arrowUturnCcwLeft":                 ArrowUturnCcwLeft,
	"arrowUturnCcwRight":                ArrowUturnCcwRight,
	"arrowUturnCwDown":                  ArrowUturnCwDown,
	"arrowUturnCwLeft":                  ArrowUturnCwLeft,
	"arrowUturnCwRight":                 ArrowUturnCwRight,
	"arrowsOppositeToDots":              ArrowsOppositeToDots,
	"arrowsRotateLeft":                  ArrowsRotateLeft,
	"arrowsRotateLeftSlash":             ArrowsRotateLeftSlash,
	"arrowsRotateRight":                 ArrowsRotateRight,
	"arrowsRotateRightSlash":            ArrowsRotateRightSlash,
	"arrowsThreeRotateLeft":             ArrowsThreeRotateLeft,
	"arrowsThreeRotateLeftLetterA":      ArrowsThreeRotateLeftLetterA,
	"arrowsThreeRotateRight":            ArrowsThreeRotateRight,
	"at":                                At,
	"ban":                               Ban,
	"bars":                              Bars,
	"barsAscendingAlignCenter":          BarsAscendingAlignCenter,
	"barsAscendingAlignLeft":            BarsAscendingAlignLeft,
	"barsAscendingAlignLeftArrowDown":   BarsAscendingAlignLeftArrowDown,
	"barsAscendingAlignLeftArrowUp":     BarsAscendingAlignLeftArrowUp,
	"barsAscendingAlignRight":           BarsAscendingAlignRight,
	"barsDescendingAlignCenter":         BarsDescendingAlignCenter,
	"barsDescendingAlignLeft":           BarsDescendingAlignLeft,
	"barsDescendingAlignLeftArrowDown":  BarsDescendingAlignLeftArrowDown,
	"barsDescendingAlignLeftArrowUp":    BarsDescendingAlignLeftArrowUp,
	"barsDescendingAlignRight":          BarsDescendingAlignRight,
	"barsPlay":                          BarsPlay,
	"barsUnaligned":                     BarsUnaligned,
	"bell":                              Bell,
	"bellDot":                           BellDot,
	"bellSlash":                         BellSlash,
	"binoculars":                        Binoculars,
	"bold":                              Bold,
	"book":                              Book,
	"bookOpen":                          BookOpen,
	"bookmark":                          Bookmark,
	"bookmarkFill":                      BookmarkFill,
	"box":                               Box,
	"boxesThree":                        BoxesThree,
	"branchesDown":                      BranchesDown,
	"branchesRight":                     BranchesRight,
	"branchesRightArrowRight":           BranchesRightArrowRight,
	"briefcase":                         Briefcase,
	"broadcastSignal":                   BroadcastSignal,
	"broomMotion":                       BroomMotion,
	"bucket":                            Bucket,
	"bucketPaint":                       BucketPaint,
	"bug":                               Bug,
	"bulb":                              Bulb,
	"calculator":                        Calculator,
	"calendar":                          Calendar,
	"camera":                            Camera,
	"car":                               Car,
	"cardClub":                          CardClub,
	"cardDiamond":                       CardDiamond,
	"cardHeart":                         CardHeart,
	"cardSpade":                         CardSpade,
	"caretDown":                         CaretDown,
	"caretLeft":                         CaretLeft,
	"caretRight":                        CaretRight,
	"caretUp":                           CaretUp,
	"caretsExpandVertical":              CaretsExpandVertical,
	"chartAreaStacked":                  ChartAreaStacked,
	"chartAreaStackedNormalized":        ChartAreaStackedNormalized,
	"chartBar":                          ChartBar,
	"chartBarStacked":                   ChartBarStacked,
	"chartColumn":                       ChartColumn,
	"chartColumnStacked":                ChartColumnStacked,
	"chartDonut":                        ChartDonut,
	"chartLine":                         ChartLine,
	"chartLineLabel":                    ChartLineLabel,
	"chartMixed":                        ChartMixed,
	"chartPie":                          ChartPie,
	"chartTreemap":                      ChartTreemap,
	"check":                             Check,
	"cherry":                            Cherry,
	"chevronDown":                       ChevronDown,
	"chevronDownWide":                   ChevronDownWide,
	"chevronLeft":                       ChevronLeft,
	"chevronRight":                      ChevronRight,
	"chevronUp":                         ChevronUp,
	"chevronUpWide":                     ChevronUpWide,
	"chevronsCollapseFromLines":         ChevronsCollapseFromLines,
	"chevronsCollapseHorizontal":        ChevronsCollapseHorizontal,
	"chevronsCollapseToLine":            ChevronsCollapseToLine,
	"chevronsCollapseUpRight":           ChevronsCollapseUpRight,
	"chevronsCollapseVertical":          ChevronsCollapseVertical,
	"chevronsDown":                      ChevronsDown,
	"chevronsDownWide":                  ChevronsDownWide,
	"chevronsExpandFromLine":            ChevronsExpandFromLine,
	"chevronsExpandHorizontal":          ChevronsExpandHorizontal,
	"chevronsExpandToLines":             ChevronsExpandToLines,
	"chevronsExpandUpRight":             ChevronsExpandUpRight,
	"chevronsExpandVertical":            ChevronsExpandVertical,
	"chevronsLeft":                      ChevronsLeft,
	"chevronsRight":                     ChevronsRight,
	"chevronsUp":                        ChevronsUp,
	"chevronsUpWide":                    ChevronsUpWide,
	"circle":                            Circle,
	"circleArrowDown":                   CircleArrowDown,
	"circleArrowDownFill":               CircleArrowDownFill,
	"circleArrowLeft":                   CircleArrowLeft,
	"circleArrowLeftFill":               CircleArrowLeftFill,
	"circleArrowRight":                  CircleArrowRight,
	"circleArrowRightFill":              CircleArrowRightFill,
	"circleArrowUp":                     CircleArrowUp,
	"circleArrowUpFill":                 CircleArrowUpFill,
	"circleCheck":                       CircleCheck,
	"circleCheckFill":                   CircleCheckFill,
	"circleChevronDown":                 CircleChevronDown,
	"circleChevronDownFill":             CircleChevronDownFill,
	"circleChevronLeft":                 CircleChevronLeft,
	"circleChevronLeftFill":             CircleChevronLeftFill,
	"circleChevronRight":                CircleChevronRight,
	"circleChevronRightFill":            CircleChevronRightFill,
	"circleChevronUp":                   CircleChevronUp,
	"circleChevronUpFill":               CircleChevronUpFill,
	"circleChevronsDown":                CircleChevronsDown,
	"circleChevronsLeft":                CircleChevronsLeft,
	"circleChevronsRight":               CircleChevronsRight,
	"circleChevronsUp":                  CircleChevronsUp,
	"circleDashed":                      CircleDashed,
	"circleDollar":                      CircleDollar,
	"circleExclamation":                 CircleExclamation,
	"circleExclamationFill":             CircleExclamationFill,
	"circleFill":                        CircleFill,
	"circleInfo":                        CircleInfo,
	"circleInfoFill":                    CircleInfoFill,
	"circleLetterC":                     CircleLetterC,
	"circleLetterF":                     CircleLetterF,
	"circleLetterH":                     CircleLetterH,
	"circleLetterP":                     CircleLetterP,
	"circleLetterR":                     CircleLetterR,
	"circleLetterS":                     CircleLetterS,
	"circleLetterT":                     CircleLetterT,
	"circleLetterW":                     CircleLetterW,
	"circleLink":                        CircleLink,
	"circleMinus":                       CircleMinus,
	"circleMinusFill":                   CircleMinusFill,
	"circleNumberEight":                 CircleNumberEight,
	"circleNumberNine":                  CircleNumberNine,
	"circleNumberOne":                   CircleNumberOne,
	"circleNumberSeven":                 CircleNumberSeven,
	"circleNumberSix":                   CircleNumberSix,
	"circleNumberZero":                  CircleNumberZero,
	"circlePause":                       CirclePause,
	"circlePlay":                        CirclePlay,
	"circlePlus":                        CirclePlus,
	"circlePlusFill":                    CirclePlusFill,
	"circleQuestion":                    CircleQuestion,
	"circleQuestionFill":                CircleQuestionFill,
	"circleRuble":                       CircleRuble,
	"circleStop":                        CircleStop,
	"circleXmark":                       CircleXmark,
	"circleXmarkFill":                   CircleXmarkFill,
	"circlesConcentric":                 CirclesConcentric,
	"circlesFiveRandom":                 CirclesFiveRandom,
	"circlesFourDiamond":                CirclesFourDiamond,
	"circlesFourSquare":                 CirclesFourSquare,
	"circlesIntersection":               CirclesIntersection,
	"circlesThreePlus":                  CirclesThreePlus,
	"clock":                             Clock,
	"clockArrowRotateLeft":              ClockArrowRotateLeft,
	"clockFill":                         ClockFill,
	"cloud":                             Cloud,
	"cloudArrowUpIn":                    CloudArrowUpIn,
	"cloudCheck":                        CloudCheck,
	"cloudGear":                         CloudGear,
	"cloudNutHex":                       CloudNutHex,
	"clouds":                            Clouds,
	"code":                              Code,
	"codeCommit":                        CodeCommit,
	"codeCommitHorizontal":              CodeCommitHorizontal,
	"codeCommits":                       CodeCommits,
	"codeCompare":                       CodeCompare,
	"codeFork":                          CodeFork,
	"codeMerge":                         CodeMerge,
	"codePullRequest":                   CodePullRequest,
	"codePullRequestArrowLeft":          CodePullRequestArrowLeft,
	"codePullRequestArrowRight":         CodePullRequestArrowRight,
	"codePullRequestCheck":              CodePullRequestCheck,
	"codePullRequestXmark":              CodePullRequestXmark,
	"codeTrunk":                         CodeTrunk,
	"comment":                           Comment,
	"commentDot":                        CommentDot,
	"commentFill":                       CommentFill,
	"commentPlus":                       CommentPlus,
	"commentSlash":                      CommentSlash,
	"comments":                          Comments,
	"compass":                           Compass,
	"copy":                              Copy,
	"copyArrowRight":                    CopyArrowRight,
	"copyCheck":                         CopyCheck,
	"copyCheckXmark":                    CopyCheckXmark,
	"copyChevronRight":                  CopyChevronRight,
	"copyMinus":                         CopyMinus,
	"copyPicture":                       CopyPicture,
	"copyPlus":                          CopyPlus,
	"copyTransparent":                   CopyTransparent,
	"copyXmark":                         CopyXmark,
	"cpu":                               Cpu,
	"cpus":                              Cpus,
	"creditCard":                        CreditCard,
	"crop":                              Crop,
	"crownDiamond":                      CrownDiamond,
	"cube":                              Cube,
	"cubesThree":                        CubesThree,
	"cubesThreeOverlap":                 CubesThreeOverlap,
	"cup":                               Cup,
	"curlyBrackets":                     CurlyBrackets,
	"curlyBracketsFunction":             CurlyBracketsFunction,
	"curlyBracketsLock":                 CurlyBracketsLock,
	"database":                          Database,
	"databaseArrowRight":                DatabaseArrowRight,
	"databaseFill":                      DatabaseFill,
	"databaseMagnifier":                 DatabaseMagnifier,
	"databases":                         Databases,
	"delete":                            Delete,
	"diamond":                           Diamond,
	"diamondExclamation":                DiamondExclamation,
	"diamondFill":                       DiamondFill,
	"diceFive":                          DiceFive,
	"diceFour":                          DiceFour,
	"diceOne":                           DiceOne,
	"diceSix":                           DiceSix,
	"diceThree":                         DiceThree,
	"diceTwo":                           DiceTwo,
	"display":                           Display,
	"displayPulse":                      DisplayPulse,
	"dotsNine":                          DotsNine,
	"droplet":                           Droplet,
	"ear":                               Ear,
	"ellipsis":                          Ellipsis,
	"ellipsisVertical":                  EllipsisVertical,
	"envelope":                          Envelope,
	"envelopeOpen":                      EnvelopeOpen,
	"envelopeOpenXmark":                 EnvelopeOpenXmark,
	"equal":                             Equal,
	"eraser":                            Eraser,
	"exclamationShape":                  ExclamationShape,
	"eye":                               Eye,
	"eyeSlash":                          EyeSlash,
	"eyesLookLeft":                      EyesLookLeft,
	"eyesLookRight":                     EyesLookRight,
	"faceAlien":                         FaceAlien,
	"faceFun":                           FaceFun,
	"faceNeutral":                       FaceNeutral,
	"faceNeutralDashed":                 FaceNeutralDashed,
	"faceRobot":                         FaceRobot,
	"faceSad":                           FaceSad,
	"faceSmile":                         FaceSmile,
	"faceSurprise":                      FaceSurprise,
	"factory":                           Factory,
	"file":                              File,
	"fileArrowDown":                     FileArrowDown,
	"fileArrowLeft":                     FileArrowLeft,
	"fileArrowRight":                    FileArrowRight,
	"fileArrowRightOut":                 FileArrowRightOut,
	"fileArrowUp":                       FileArrowUp,
	"fileCheck":                         FileCheck,
	"fileCode":                          FileCode,
	"fileDollar":                        FileDollar,
	"fileExclamation":                   FileExclamation,
	"fileMagnifier":                     FileMagnifier,
	"fileMinus":                         FileMinus,
	"filePlus":                          FilePlus,
	"fileQuestion":                      FileQuestion,
	"fileRuble":                         FileRuble,
	"fileText":                          FileText,
	"fileXmark":                         FileXmark,
	"fileZipper":                        FileZipper,
	"files":                             Files,
	"filmstrip":                         Filmstrip,
	"fingerprint":                       Fingerprint,
	"flag":                              Flag,
	"flame":                             Flame,
	"flask":                             Flask,
	"floppyDisk":                        FloppyDisk,
	"folder":                            Folder,
	"folderArrowDown":                   FolderArrowDown,
	"folderArrowLeft":                   FolderArrowLeft,
	"folderArrowRight":                  FolderArrowRight,
	"folderArrowUp":                     FolderArrowUp,
	"folderArrowUpIn":                   FolderArrowUpIn,
	"folderCheck":                       FolderCheck,
	"folderCode":                        FolderCode,
	"folderExclamation":                 FolderExclamation,
	"folderFill":                        FolderFill,
	"folderFlows":                       FolderFlows,
	"folderHouse":                       FolderHouse,
	"folderKeyhole":                     FolderKeyhole,
	"folderLock":                        FolderLock,
	"folderMagnifier":                   FolderMagnifier,
	"folderOpen":                        FolderOpen,
	"folderOpenFill":                    FolderOpenFill,
	"folderPlus":                        FolderPlus,
	"folderTree":                        FolderTree,
	"folders":                           Folders,
	"font":                              Font,
	"fontCase":                          FontCase,
	"fontCursor":                        FontCursor,
	"frame":                             Frame,
	"frames":                            Frames,
	"function":                          Function,
	"funnel":                            Funnel,
	"funnelXmark":                       FunnelXmark,
	"gear":                              Gear,
	"gearBranches":                      GearBranches,
	"gearDot":                           GearDot,
	"gearPlay":                          GearPlay,
	"geo":                               Geo,
	"geoDots":                           GeoDots,
	"geoFill":                           GeoFill,
	"geoPin":                            GeoPin,
	"geoPolygons":                       GeoPolygons,
	"ghost":                             Ghost,
	"gift":                              Gift,
	"globe":                             Globe,
	"gpu":                               Gpu,
	"graduationCap":                     GraduationCap,
	"graphNode":                         GraphNode,
	"grip":                              Grip,
	"gripHorizontal":                    GripHorizontal,
	"hand":                              Hand,
	"handOk":                            HandOk,
	"handPointDown":                     HandPointDown,
	"handPointLeft":                     HandPointLeft,
	"handPointRight":                    HandPointRight,
	"handPointUp":                       HandPointUp,
	"handStop":                          HandStop,
	"handset":                           Handset,
	"hardDrive":                         HardDrive,
	"hashtag":                           Hashtag,
	"heading":                           Heading,
	"headingFive":                       HeadingFive,
	"headingFour":                       HeadingFour,
	"headingOne":                        HeadingOne,
	"headingSix":                        HeadingSix,
	"headingThree":                      HeadingThree,
	"headingTwo":                        HeadingTwo,
	"headphones":                        Headphones,
	"heart":                             Heart,
	"heartCrack":                        HeartCrack,
	"heartFill":                         HeartFill,
	"heartPulse":                        HeartPulse,
	"hierarchy":                         Hierarchy,
	"house":                             House,
	"italic":                            Italic,
	"key":                               Key,
	"keyboard":                          Keyboard,
	"layers":                            Layers,
	"layersThreeDiagonal":               LayersThreeDiagonal,
	"layersVertical":                    LayersVertical,
	"layoutCells":                       LayoutCells,
	"layoutCellsLarge":                  LayoutCellsLarge,
	"layoutColumns":                     LayoutColumns,
	"layoutColumnsThree":                LayoutColumnsThree,
	"layoutFooter":                      LayoutFooter,
	"layoutHeader":                      LayoutHeader,
	"layoutHeaderCells":                 LayoutHeaderCells,
	"layoutHeaderCellsLarge":            LayoutHeaderCellsLarge,
	"layoutHeaderCellsLargeFill":        LayoutHeaderCellsLargeFill,
	"layoutHeaderCellsLargeLetterD":     LayoutHeaderCellsLargeLetterD,
	"layoutHeaderCellsLargeThunderbolt": LayoutHeaderCellsLargeThunderbolt,
	"layoutHeaderColumns":               LayoutHeaderColumns,
	"layoutHeaderCursor":                LayoutHeaderCursor,
	"layoutHeaderSideContent":           LayoutHeaderSideContent,
	"layoutList":                        LayoutList,
	"layoutRows":                        LayoutRows,
	"layoutRowsThree":                   LayoutRowsThree,
	"layoutSideContent":                 LayoutSideContent,
	"layoutSideContentLeft":             LayoutSideContentLeft,
	"layoutSideContentRight":            LayoutSideContentRight,
	"layoutSplitColumns":                LayoutSplitColumns,
	"layoutSplitColumnsThree":           LayoutSplitColumnsThree,
	"layoutSplitRows":                   LayoutSplitRows,
	"layoutSplitSideContentLeft":        LayoutSplitSideContentLeft,
	"layoutSplitSideContentRight":       LayoutSplitSideContentRight,
	"layoutTabs":                        LayoutTabs,
	"letterGroup":                       LetterGroup,
	"letterM":                           LetterM,
	"lifeRing":                          LifeRing,
	"link":                              Link,
	"linkSlash":                         LinkSlash,
	"listCheck":                         ListCheck,
	"listCheckLock":                     ListCheckLock,
	"listOl":                            ListOl,
	"listTimeline":                      ListTimeline,
	"listUl":                            ListUl,
	"lock":                              Lock,
	"lockOpen":                          LockOpen,
	"logoAcrobat":                       LogoAcrobat,
	"logoDocker":                        LogoDocker,
	"logoDrawIo":                        LogoDrawIo,
	"logoFacebook":                      LogoFacebook,
	"logoGitlab":                        LogoGitlab,
	"logoLinux":                         LogoLinux,
	"logoMacos":                         LogoMacos,
	"logoMarkdown":                      LogoMarkdown,
	"logoMermaid":                       LogoMermaid,
	"logoNotion":                        LogoNotion,
	"logoOsi":                           LogoOsi,
	"logoPython":                        LogoPython,
	"logoStackOverflow":                 LogoStackOverflow,
	"logoTelegram":                      LogoTelegram,
	"logoUbuntu":                        LogoUbuntu,
	"logoWindows":                       LogoWindows,
	"logoYandex":                        LogoYandex,
	"logoYandexCloud":                   LogoYandexCloud,
	"logoYandexMessenger":               LogoYandexMessenger,
	"logoYandexTracker":                 LogoYandexTracker,
	"magicWand":                         MagicWand,
	"magnet":                            Magnet,
	"magnifier":                         Magnifier,
	"magnifierMinus":                    MagnifierMinus,
	"magnifierPlus":                     MagnifierPlus,
	"mapPin":                            MapPin,
	"mapPinMinus":                       MapPinMinus,
	"mapPinPlus":                        MapPinPlus,
	"mathIntersectionShape":             MathIntersectionShape,
	"mathOperations":                    MathOperations,
	"mathUnionShape":                    MathUnionShape,
	"medal":                             Medal,
	"megaphone":                         Megaphone,
	"microphone":                        Microphone,
	"microphoneSlash":                   MicrophoneSlash,
	"minus":                             Minus,
	"molecule":                          Molecule,
	"moon":                              Moon,
	"mug":                               Mug,
	"musicNote":                         MusicNote,
	"nodesDown":                         NodesDown,
	"nodesLeft":                         NodesLeft,
	"nodesRight":                        NodesRight,
	"nodesUp":                           NodesUp,
	"nutHex":                            NutHex,
	"objectAlignBottom":                 ObjectAlignBottom,
	"objectAlignCenterHorizontal":       ObjectAlignCenterHorizontal,
	"objectAlignCenterVertical":         ObjectAlignCenterVertical,
	"objectAlignJustifyHorizontal":      ObjectAlignJustifyHorizontal,
	"objectAlignJustifyVertical":        ObjectAlignJustifyVertical,
	"objectAlignLeft":                   ObjectAlignLeft,
	"objectAlignRight":                  ObjectAlignRight,
	"objectAlignTop":                    ObjectAlignTop,
	"objectsAlignBottom":                ObjectsAlignBottom,
	"objectsAlignCenterHorizontal":      ObjectsAlignCenterHorizontal,
	"objectsAlignCenterVertical":        ObjectsAlignCenterVertical,
	"objectsAlignJustifyHorizontal":     ObjectsAlignJustifyHorizontal,
	"objectsAlignJustifyVertical":       ObjectsAlignJustifyVertical,
	"objectsAlignLeft":                  ObjectsAlignLeft,
	"objectsAlignRight":                 ObjectsAlignRight,
	"objectsAlignTop":                   ObjectsAlignTop,
	"octagonXmark":                      OctagonXmark,
	"officeBadge":                       OfficeBadge,
	"palette":                           Palette,
	"paperPlane":                        PaperPlane,
	"paperclip":                         Paperclip,
	"passport":                          Passport,
	"pause":                             Pause,
	"pauseFill":                         PauseFill,
	"pencil":                            Pencil,
	"pencilToLine":                      PencilToLine,
	"pencilToSquare":                    PencilToSquare,
	"percent":                           Percent,
	"person":                            Person,
	"personGear":                        PersonGear,
	"personMagnifier":                   PersonMagnifier,
	"personNutHex":                      PersonNutHex,
	"personPencil":                      PersonPencil,
	"personPlanetEarth":                 PersonPlanetEarth,
	"personPlus":                        PersonPlus,
	"personSpeaker":                     PersonSpeaker,
	"personWorker":                      PersonWorker,
	"personXmark":                       PersonXmark,
	"persons":                           Persons,
	"personsLock":                       PersonsLock,
	"picture":                           Picture,
	"pill":                              Pill,
	"pin":                               Pin,
	"pinFill":                           PinFill,
	"pinSlash":                          PinSlash,
	"pinSlashFill":                      PinSlashFill,
	"pipeline":                          Pipeline,
	"planetEarth":                       PlanetEarth,
	"play":                              Play,
	"playFill":                          PlayFill,
	"plugConnection":                    PlugConnection,
	"plugWire":                          PlugWire,
	"plus":                              Plus,
	"power":                             Power,
	"printer":                           Printer,
	"pulse":                             Pulse,
	"puzzle":                            Puzzle,
	"qrCode":                            QrCode,
	"quoteClose":                        QuoteClose,
	"quoteOpen":                         QuoteOpen,
	"receipt":                           Receipt,
	"rectanglePulse":                    RectanglePulse,
	"rectanglesFour":                    RectanglesFour,
	"rocket":                            Rocket,
	"roundBrackets":                     RoundBrackets,
	"route":                             Route,
	"sack":                              Sack,
	"scalesBalanced":                    ScalesBalanced,
	"scalesUnbalanced":                  ScalesUnbalanced,
	"scissors":                          Scissors,
	"sealCheck":                         SealCheck,
	"sealPercent":                       SealPercent,
	"server":                            Server,
	"shapesFour":                        ShapesFour,
	"shapesThree":                       ShapesThree,
	"shield":                            Shield,
	"shieldCheck":                       ShieldCheck,
	"shieldExclamation":                 ShieldExclamation,
	"shieldKeyhole":                     ShieldKeyhole,
	"shoppingBag":                       ShoppingBag,
	"shoppingBasket":                    ShoppingBasket,
	"shoppingCart":                      ShoppingCart,
	"shuffle":                           Shuffle,
	"signal":                            Signal,
	"sliders":                           Sliders,
	"slidersVertical":                   SlidersVertical,
	"smartphone":                        Smartphone,
	"snail":                             Snail,
	"snowflake":                         Snowflake,
	"sphere":                            Sphere,
	"square":                            Square,
	"squareArticle":                     SquareArticle,
	"squareBars":                        SquareBars,
	"squareBarsVertical":                SquareBarsVertical,
	"squareBrackets":                    SquareBrackets,
	"squareBracketsBarsVertical":        SquareBracketsBarsVertical,
	"squareBracketsLetterA":             SquareBracketsLetterA,
	"squareChartBar":                    SquareChartBar,
	"squareChartColumn":                 SquareChartColumn,
	"squareCheck":                       SquareCheck,
	"squareDashed":                      SquareDashed,
	"squareDashedCircle":                SquareDashedCircle,
	"squareDashedLetterA":               SquareDashedLetterA,
	"squareDashedLetterT":               SquareDashedLetterT,
	"squareDashedText":                  SquareDashedText,
	"squareDot":                         SquareDot,
	"squareExclamation":                 SquareExclamation,
	"squareFill":                        SquareFill,
	"squareHashtag":                     SquareHashtag,
	"squareLetterP":                     SquareLetterP,
	"squareLetterT":                     SquareLetterT,
	"squareListUl":                      SquareListUl,
	"squareMinus":                       SquareMinus,
	"squarePlus":                        SquarePlus,
	"squareXmark":                       SquareXmark,
	"star":                              Star,
	"starFill":                          StarFill,
	"stethoscope":                       Stethoscope,
	"sticker":                           Sticker,
	"stop":                              Stop,
	"stopFill":                          StopFill,
	"stopwatch":                         Stopwatch,
	"strikethrough":                     Strikethrough,
	"suitcase":                          Suitcase,
	"sun":                               Sun,
	"tshirt":                            Tshirt,
	"tag":                               Tag,
	"tagDollar":                         TagDollar,
	"tagRuble":                          TagRuble,
	"tags":                              Tags,
	"target":                            Target,
	"targetDart":                        TargetDart,
	"terminal":                          Terminal,
	"terminalLine":                      TerminalLine,
	"textAlignCenter":                   TextAlignCenter,
	"textAlignJustify":                  TextAlignJustify,
	"textAlignLeft":                     TextAlignLeft,
	"textAlignRight":                    TextAlignRight,
	"textIcon":                          TextIcon,
	"textIndent":                        TextIndent,
	"textOutdent":                       TextOutdent,
	"thumbsDown":                        ThumbsDown,
	"thumbsDownFill":                    ThumbsDownFill,
	"thumbsUp":                          ThumbsUp,
	"thumbsUpFill":                      ThumbsUpFill,
	"thumbsUpTwo":                       ThumbsUpTwo,
	"thunderbolt":                       Thunderbolt,
	"thunderboltFill":                   ThunderboltFill,
	"ticket":                            Ticket,
	"timeline":                          Timeline,
	"timestamps":                        Timestamps,
	"trafficLight":                      TrafficLight,
	"trashBin":                          TrashBin,
	"tray":                              Tray,
	"triangleDown":                      TriangleDown,
	"triangleDownFill":                  TriangleDownFill,
	"triangleExclamation":               TriangleExclamation,
	"triangleExclamationFill":           TriangleExclamationFill,
	"triangleLeft":                      TriangleLeft,
	"triangleLeftFill":                  TriangleLeftFill,
	"triangleRight":                     TriangleRight,
	"triangleRightFill":                 TriangleRightFill,
	"triangleThunderbolt":               TriangleThunderbolt,
	"triangleUp":                        TriangleUp,
	"triangleUpFill":                    TriangleUpFill,
	"trolley":                           Trolley,
	"tv":                                Tv,
	"tvRetro":                           TvRetro,
	"umbrella":                          Umbrella,
	"underline":                         Underline,
	"vault":                             Vault,
	"vectorCircle":                      VectorCircle,
	"vectorSquare":                      VectorSquare,
	"video":                             Video,
	"volume":                            Volume,
	"volumeFill":                        VolumeFill,
	"volumeLow":                         VolumeLow,
	"volumeOff":                         VolumeOff,
	"volumeXmark":                       VolumeXmark,
	"weightHanging":                     WeightHanging,
	"wrench":                            Wrench,
	"xmark":                             Xmark,
}

func AbbrApi(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.02 11.419L4.691 10H2.808l-.327 1.419a.75.75 0 1 1-1.462-.338L2.407 5.07a1.379 1.379 0 0 1 2.686 0l1.388 6.012a.75.75 0 1 1-1.462.338Zm-1.27-5.5L4.346 8.5H3.154zM14.5 4.75a.75.75 0 1 0-1.5 0v6.5a.75.75 0 0 0 1.5 0zM8.25 4a.75.75 0 0 0-.75.75v6.5a.75.75 0 0 0 1.5 0V9.5h.5a2.75 2.75 0 0 0 0-5.5zM9.5 8H9V5.5h.5a1.25 1.25 0 0 1 0 2.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AbbrMl(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 11.25V6.672l1.436 3.39a1.129 1.129 0 0 0 2.076.008l1.513-3.492v4.672a.75.75 0 0 0 1.5 0V5.218a1.218 1.218 0 0 0-2.335-.484L4.98 8.679L3.307 4.732A1.201 1.201 0 0 0 1 5.202v6.048a.75.75 0 1 0 1.5 0m9.5-6.5a.75.75 0 0 0-1.5 0v6.5c0 .414.336.75.75.75h3.5a.75.75 0 0 0 0-1.5H12z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AbbrQl(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.75 4a.75.75 0 0 1 .75.75v5.75h2.75a.75.75 0 0 1 0 1.5h-3.5a.75.75 0 0 1-.75-.75v-6.5a.75.75 0 0 1 .75-.75m-3.21 6.997c.611-.666.96-1.665.96-2.997c0-2.667-1.4-4-3.5-4S1.5 5.333 1.5 8s1.4 4 3.5 4c.448 0 .864-.06 1.24-.182l1.456 1.461a.75.75 0 1 0 1.063-1.058zM7 8c0 1.083-.282 1.675-.567 1.985c-.27.294-.708.515-1.433.515c-.725 0-1.163-.22-1.433-.515C3.282 9.675 3 9.083 3 8c0-1.082.282-1.675.567-1.985c.27-.294.708-.515 1.433-.515c.725 0 1.163.22 1.433.515C6.718 6.325 7 6.918 7 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AbbrSql(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiAbbrSql0)"><path fill="currentColor" fill-rule="evenodd" d="M13.5 4.75a.75.75 0 0 0-1.5 0v6.5c0 .414.336.75.75.75h2.5a.75.75 0 0 0 0-1.5H13.5zM11 8c0 1.32-.257 2.314-.709 2.98l.99.99a.75.75 0 1 1-1.061 1.06l-1.132-1.132a2.43 2.43 0 0 1-.713.102C6.8 12 5.75 10.667 5.75 8S6.8 4 8.375 4S11 5.333 11 8m-1.957 2.15c.222-.322.457-.98.457-2.149c0-1.168-.235-1.826-.457-2.149c-.175-.254-.366-.35-.668-.35c-.302 0-.493.096-.668.35c-.222.323-.457.981-.457 2.15c0 1.167.235 1.826.457 2.148c.175.254.366.351.668.351c.302 0 .493-.097.668-.351M2.562 4.002c-.689 0-1.285.249-1.702.713C.453 5.17.277 5.76.277 6.352c0 .732.326 1.28.759 1.658c.41.358.908.561 1.305.684c.398.122.684.217.89.372c.147.113.27.266.27.625c0 .29-.109.469-.253.589a1.08 1.08 0 0 1-.686.222c-.202 0-.438-.093-.658-.293a1.391 1.391 0 0 1-.423-.684a.75.75 0 1 0-1.462.333c.127.557.453 1.079.878 1.463c.426.386 1.006.681 1.665.681c.578 0 1.177-.18 1.647-.57c.485-.405.791-1.006.791-1.74c0-.843-.352-1.433-.864-1.82c-.434-.33-.962-.492-1.306-.597l-.047-.015c-.31-.095-.579-.221-.76-.38a.65.65 0 0 1-.246-.528c0-.304.09-.512.2-.635c.099-.11.27-.215.585-.215c.322 0 .498.099.612.206c.13.123.233.313.3.56a.75.75 0 1 0 1.449-.386c-.113-.42-.325-.892-.719-1.263c-.409-.387-.961-.617-1.642-.617" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiAbbrSql0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AbbrZip(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 4.75A.75.75 0 0 1 1.75 4h3.113a1.137 1.137 0 0 1 .96 1.748L2.797 10.5H5.25a.75.75 0 0 1 0 1.5H2.137a1.137 1.137 0 0 1-.96-1.748L4.203 5.5H1.75A.75.75 0 0 1 1 4.75M7.75 4a.75.75 0 0 1 .75.75v6.5a.75.75 0 0 1-1.5 0v-6.5A.75.75 0 0 1 7.75 4m3 0a.75.75 0 0 0-.75.75v6.5a.75.75 0 0 0 1.5 0V9.5h.5A2.75 2.75 0 1 0 12 4zM12 8h-.5V5.5h.5A1.25 1.25 0 1 1 12 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AntennaSignal(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.56 4.403c.27.314.223.784-.015 1.123A2.986 2.986 0 0 0 5 7.25c0 .642.202 1.237.545 1.724c.238.339.284.809.015 1.123c-.27.315-.75.354-1.015.036A4.482 4.482 0 0 1 3.5 7.25c0-1.097.393-2.102 1.045-2.883c.266-.318.745-.279 1.015.036m4.88 0c-.27.314-.223.784.015 1.123c.344.487.545 1.082.545 1.724c0 .642-.201 1.237-.545 1.724c-.238.339-.284.809-.015 1.123c.27.315.75.354 1.015.036A4.482 4.482 0 0 0 12.5 7.25a4.482 4.482 0 0 0-1.045-2.883c-.265-.318-.745-.279-1.015.036m1.953-2.278c-.27.315-.23.785.05 1.092A5.978 5.978 0 0 1 14 7.25c0 1.553-.59 2.968-1.558 4.033c-.278.307-.319.777-.05 1.092c.27.314.747.353 1.033.053A7.474 7.474 0 0 0 15.5 7.25c0-2.008-.79-3.832-2.075-5.178c-.286-.3-.763-.261-1.032.053m-8.786 0c-.27-.314-.746-.353-1.032-.053A7.475 7.475 0 0 0 .5 7.25c0 2.008.79 3.832 2.075 5.178c.286.3.763.261 1.032-.053c.27-.315.23-.785-.05-1.092A5.978 5.978 0 0 1 2 7.25c0-1.553.59-2.968 1.558-4.033c.278-.307.319-.777.05-1.092ZM8.75 8.55a1.5 1.5 0 1 0-1.5 0v5.701a.75.75 0 0 0 1.5 0v-5.7Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Aperture(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.48 5.48 0 0 1-.48 2.251a6.965 6.965 0 0 0-.897-2.61c-.64-1.109-1.517-1.984-2.467-2.508l-.013-.008a5.037 5.037 0 0 0-.19-.098c-.634-.314-1.33-.484-2.006-.428c.464-.575.995-1.002 1.514-1.259c.733-.362 1.356-.353 1.789-.103v-.001A5.498 5.498 0 0 1 13.5 8M7.438 2.528A5.5 5.5 0 0 0 2.5 8c0 .5.304 1.044.984 1.497c.482.322 1.117.568 1.847.682c-.386-.558-.587-1.245-.633-1.951a5.047 5.047 0 0 1-.009-.213v-.012c-.022-1.086.298-2.284.938-3.394a6.966 6.966 0 0 1 1.811-2.08ZM6.189 8.002c.002.043.003.087.006.13c.053.815.372 1.35.805 1.6c.433.25 1.056.259 1.789-.104l.117-.06l.11-.07C9.696 9.043 10 8.5 10 8s-.304-1.044-.984-1.497a3.492 3.492 0 0 0-.11-.071a3.511 3.511 0 0 0-.117-.06C8.056 6.009 7.433 6.018 7 6.268c-.433.25-.752.785-.805 1.6a3.51 3.51 0 0 0-.006.134m3.468 2.865l.011-.007a4.22 4.22 0 0 0 .18-.115c.59-.392 1.084-.91 1.374-1.523c.266.69.37 1.362.333 1.94c-.053.816-.372 1.351-.805 1.601A5.474 5.474 0 0 1 8 13.5a5.493 5.493 0 0 1-4.458-2.278c.8.34 1.73.528 2.708.528c1.28 0 2.477-.322 3.407-.883M1 8a7 7 0 1 1 14 0A7 7 0 0 1 1 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Archive(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 8h11v3a1.5 1.5 0 0 1-1.5 1.5H4A1.5 1.5 0 0 1 2.5 11zm10.697-1.5l-1.851-2.777a.5.5 0 0 0-.416-.223H5.07a.5.5 0 0 0-.416.223L2.803 6.5zM15 7.408V11a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3V7.408a3 3 0 0 1 .504-1.664l1.902-2.853A2 2 0 0 1 5.07 2h5.86a2 2 0 0 1 1.664.89l1.902 2.854A3 3 0 0 1 15 7.408M9.25 11a.75.75 0 0 0 0-1.5h-2.5a.75.75 0 0 0 0 1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDown(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1.25a.75.75 0 0 1 .75.75v10.19l2.72-2.72a.75.75 0 1 1 1.06 1.06l-4 4a.75.75 0 0 1-1.06 0l-4-4a.75.75 0 1 1 1.06-1.06l2.72 2.72V2A.75.75 0 0 1 8 1.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownFromLine(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.53 14.78a.75.75 0 0 1-1.06 0l-2.5-2.5a.75.75 0 1 1 1.06-1.06l1.22 1.22V4.75a.75.75 0 0 1 1.5 0v7.69l1.22-1.22a.75.75 0 1 1 1.06 1.06zM14.25 2.5a.75.75 0 0 0 0-1.5H1.75a.75.75 0 0 0 0 1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownToLine(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.53 11.78a.75.75 0 0 1-1.06 0l-2.5-2.5a.75.75 0 0 1 1.06-1.06l1.22 1.22V1.75a.75.75 0 1 1 1.5 0v7.69l1.22-1.22a.75.75 0 1 1 1.06 1.06zM1.75 13.5a.75.75 0 0 0 0 1.5h12.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownToSquare(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.53 11.78a.75.75 0 0 1-1.06 0l-2.5-2.5a.75.75 0 0 1 1.06-1.06l1.22 1.22V1.75a.75.75 0 0 1 1.5 0v7.69l1.22-1.22a.75.75 0 1 1 1.06 1.06zM4.25 4a.75.75 0 1 1 0 1.5H4A1.5 1.5 0 0 0 2.5 7v5A1.5 1.5 0 0 0 4 13.5h8a1.5 1.5 0 0 0 1.5-1.5V7A1.5 1.5 0 0 0 12 5.5h-.25a.75.75 0 0 1 0-1.5H12a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3V7a3 3 0 0 1 3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.75 8a.75.75 0 0 1-.75.75H3.81l2.72 2.72a.75.75 0 1 1-1.06 1.06l-4-4a.75.75 0 0 1 0-1.06l4-4a.75.75 0 0 1 1.06 1.06L3.81 7.25H14a.75.75 0 0 1 .75.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftFromLine(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.22 8.53a.75.75 0 0 1 0-1.06l2.5-2.5a.75.75 0 0 1 1.06 1.06L3.56 7.25h7.69a.75.75 0 0 1 0 1.5H3.56l1.22 1.22a.75.75 0 1 1-1.06 1.06zm12.28 5.72a.75.75 0 0 0 1.5 0V1.75a.75.75 0 0 0-1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftToLine(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.22 8.53a.75.75 0 0 1 0-1.06l2.5-2.5a.75.75 0 0 1 1.06 1.06L6.56 7.25h7.69a.75.75 0 0 1 0 1.5H6.56l1.22 1.22a.75.75 0 1 1-1.06 1.06zM2.5 1.75a.75.75 0 0 0-1.5 0v12.5a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.25 8A.75.75 0 0 1 2 7.25h10.19L9.47 4.53a.75.75 0 0 1 1.06-1.06l4 4a.75.75 0 0 1 0 1.06l-4 4a.75.75 0 1 1-1.06-1.06l2.72-2.72H2A.75.75 0 0 1 1.25 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightArrowLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.78 3.72a.75.75 0 0 1 0 1.06l-3 3a.75.75 0 1 1-1.06-1.06L11.44 5H2.75a.75.75 0 1 1 0-1.5h8.69L9.72 1.78A.75.75 0 0 1 10.78.72zM2 11.75a.75.75 0 0 1 .22-.53l3-3a.75.75 0 1 1 1.06 1.06L4.56 11h8.69a.75.75 0 0 1 0 1.5H4.56l1.72 1.72a.75.75 0 1 1-1.06 1.06l-3-3a.75.75 0 0 1-.22-.53" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightFromLine(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.78 7.47a.75.75 0 0 1 0 1.06l-2.5 2.5a.75.75 0 1 1-1.06-1.06l1.22-1.22H4.75a.75.75 0 0 1 0-1.5h7.69l-1.22-1.22a.75.75 0 0 1 1.06-1.06zM2.5 1.75a.75.75 0 0 0-1.5 0v12.5a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightFromSquare(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.78 7.47a.75.75 0 0 1 0 1.06l-2.5 2.5a.75.75 0 1 1-1.06-1.06l1.22-1.22H4.75a.75.75 0 0 1 0-1.5h7.69l-1.22-1.22a.75.75 0 0 1 1.06-1.06zM9.5 4.25a.75.75 0 0 1-1.5 0V4a1.5 1.5 0 0 0-1.5-1.5H4A1.5 1.5 0 0 0 2.5 4v8A1.5 1.5 0 0 0 4 13.5h2.5A1.5 1.5 0 0 0 8 12v-.25a.75.75 0 0 1 1.5 0V12a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3V4a3 3 0 0 1 3-3h2.5a3 3 0 0 1 3 3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightToLine(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.78 7.47a.75.75 0 0 1 0 1.06l-2.5 2.5a.75.75 0 1 1-1.06-1.06l1.22-1.22H1.75a.75.75 0 0 1 0-1.5h7.69L8.22 6.03a.75.75 0 0 1 1.06-1.06zm1.72 6.78a.75.75 0 0 0 1.5 0V1.75a.75.75 0 0 0-1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightToSquare(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.78 7.47a.75.75 0 0 1 0 1.06l-2.5 2.5a.75.75 0 1 1-1.06-1.06l1.22-1.22H1.75a.75.75 0 0 1 0-1.5h7.69L8.22 6.03a.75.75 0 0 1 1.06-1.06zM4 11.75a.75.75 0 0 1 1.5 0V12A1.5 1.5 0 0 0 7 13.5h5a1.5 1.5 0 0 0 1.5-1.5V4A1.5 1.5 0 0 0 12 2.5H7A1.5 1.5 0 0 0 5.5 4v.25a.75.75 0 0 1-1.5 0V4a3 3 0 0 1 3-3h5a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRotateLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1.5a6.5 6.5 0 1 1-6.445 7.348a.75.75 0 1 1 1.487-.194A5.001 5.001 0 1 0 4.43 4.5h1.32a.75.75 0 0 1 0 1.5h-3A.75.75 0 0 1 2 5.25v-3a.75.75 0 0 1 1.5 0v1.06A6.48 6.48 0 0 1 8 1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRotateRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1.5a6.5 6.5 0 1 0 6.445 7.348a.75.75 0 1 0-1.487-.194A5.001 5.001 0 1 1 11.57 4.5h-1.32a.75.75 0 0 0 0 1.5h3a.75.75 0 0 0 .75-.75v-3a.75.75 0 0 0-1.5 0v1.06A6.48 6.48 0 0 0 8 1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowShapeDown(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 7.5H3.33a.33.33 0 0 0-.252.542l4.42 5.225a.656.656 0 0 0 1.003 0l4.42-5.225a.33.33 0 0 0-.251-.542H10V3a.5.5 0 0 0-.5-.5h-3A.5.5 0 0 0 6 3zM4.5 3a2 2 0 0 1 2-2h3a2 2 0 0 1 2 2v3h1.17a1.83 1.83 0 0 1 1.397 3.011l-4.421 5.225a2.156 2.156 0 0 1-3.292 0L1.933 9.011A1.83 1.83 0 0 1 3.329 6H4.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowShapeDownFromLine(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.5 9H4.029l3.499 4.276a.61.61 0 0 0 .944 0L11.971 9H9.5V5.75a.25.25 0 0 0-.25-.25h-2.5a.25.25 0 0 0-.25.25zm2.75-5H6.749A1.75 1.75 0 0 0 5 5.75V7.5H3.974a1.474 1.474 0 0 0-1.14 2.407l3.533 4.319a2.11 2.11 0 0 0 3.266 0l3.534-4.319a1.474 1.474 0 0 0-1.14-2.407H11V5.75A1.75 1.75 0 0 0 9.25 4m-7.5-1.5a.75.75 0 0 1 0-1.5h12.5a.75.75 0 0 1 0 1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowShapeDownToLine(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.5 6H4.029l3.499 4.276a.61.61 0 0 0 .944 0L11.971 6H9.5V2.75a.25.25 0 0 0-.25-.25h-2.5a.25.25 0 0 0-.25.25zm1.498 7.5H1.75a.75.75 0 0 0 0 1.5h12.5a.75.75 0 0 0 0-1.5zm0-1.5a2.11 2.11 0 0 1-1.631-.774L2.833 6.907A1.474 1.474 0 0 1 3.973 4.5H5V2.75C5 1.784 5.784 1 6.75 1h2.5c.966 0 1.75.784 1.75 1.75V4.5h1.026a1.474 1.474 0 0 1 1.14 2.407l-3.533 4.319c-.4.49-1 .774-1.632.774H8Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowShapeLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.5 6V3.33a.33.33 0 0 0-.543-.252L2.732 7.499a.656.656 0 0 0 0 1.002l5.225 4.421a.33.33 0 0 0 .543-.252V10H13a.5.5 0 0 0 .5-.5v-3A.5.5 0 0 0 13 6zM13 4.5a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2h-3v1.17a1.83 1.83 0 0 1-3.012 1.397L1.763 9.646a2.156 2.156 0 0 1 0-3.292l5.225-4.421A1.83 1.83 0 0 1 10 3.33V4.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowShapeLeftFromLine(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 6.5V4.029L2.724 7.528a.61.61 0 0 0 0 .944L7 11.971V9.5h3.25a.25.25 0 0 0 .25-.25v-2.5a.25.25 0 0 0-.25-.25zm5 2.75V6.749A1.75 1.75 0 0 0 10.25 5H8.5V3.974a1.474 1.474 0 0 0-2.407-1.14L1.774 6.366a2.11 2.11 0 0 0 0 3.266l4.319 3.534a1.474 1.474 0 0 0 2.407-1.14V11h1.75A1.75 1.75 0 0 0 12 9.25m1.5-7.5a.75.75 0 0 1 1.5 0v12.5a.75.75 0 0 1-1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowShapeLeftToLine(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 6.5V4.029L5.724 7.528a.61.61 0 0 0 0 .944L10 11.971V9.5h3.25a.25.25 0 0 0 .25-.25v-2.5a.25.25 0 0 0-.25-.25zM2.5 7.998V1.75a.75.75 0 0 0-1.5 0v12.5a.75.75 0 0 0 1.5 0zm1.5 0c0-.632.284-1.23.774-1.631l4.319-3.534a1.474 1.474 0 0 1 2.407 1.14V5h1.75c.966 0 1.75.784 1.75 1.75v2.5A1.75 1.75 0 0 1 13.25 11H11.5v1.026a1.474 1.474 0 0 1-2.407 1.14L4.774 9.634c-.49-.4-.774-1-.774-1.632V8Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowShapeRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 10v2.67a.33.33 0 0 0 .542.252l5.225-4.421a.656.656 0 0 0 0-1.002L8.042 3.078a.33.33 0 0 0-.543.252V6H3a.5.5 0 0 0-.5.5v3a.5.5 0 0 0 .5.5zM3 11.5a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2h3V3.33a1.83 1.83 0 0 1 3.01-1.397l5.226 4.421a2.156 2.156 0 0 1 0 3.292L9.01 14.067a1.83 1.83 0 0 1-3.012-1.397V11.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowShapeRightFromLine(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9 9.5v2.471l4.276-3.499a.61.61 0 0 0 0-.944L9 4.029V6.5H5.75a.25.25 0 0 0-.25.25v2.5c0 .138.112.25.25.25zM4 6.75v2.501A1.75 1.75 0 0 0 5.75 11H7.5v1.026a1.474 1.474 0 0 0 2.407 1.14l4.319-3.533a2.11 2.11 0 0 0 0-3.266L9.907 2.833A1.474 1.474 0 0 0 7.5 3.973V5H5.75A1.75 1.75 0 0 0 4 6.75m-1.5 7.5a.75.75 0 0 1-1.5 0V1.75a.75.75 0 0 1 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowShapeRightToLine(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 9.5v2.471l4.276-3.499a.61.61 0 0 0 0-.944L6 4.029V6.5H2.75a.25.25 0 0 0-.25.25v2.5c0 .138.112.25.25.25zm7.5-1.498v6.248a.75.75 0 0 0 1.5 0V1.75a.75.75 0 0 0-1.5 0zm-1.5 0a2.11 2.11 0 0 1-.774 1.631l-4.319 3.534a1.474 1.474 0 0 1-2.407-1.14V11H2.75A1.75 1.75 0 0 1 1 9.25v-2.5C1 5.784 1.784 5 2.75 5H4.5V3.974a1.474 1.474 0 0 1 2.407-1.14l4.319 3.533c.49.4.774 1 .774 1.632V8Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowShapeTurnUpLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 9v2.665a.335.335 0 0 1-.55.257L1.73 7.988a.635.635 0 0 1 0-.976l4.72-3.934a.335.335 0 0 1 .55.257V6h1.5c1.584 0 3.182.571 4.241 1.692c.9.951 1.549 2.446 1.31 4.723c-.65-1.026-1.365-1.837-2.201-2.413C10.802 9.278 9.677 9 8.5 9zm3 1.731c1.162.396 2.165 1.337 3.151 3.106c.223.4.64.663 1.098.663c.552 0 1.04-.376 1.143-.917C16.598 7.237 12.322 4.5 8.501 4.5V3.335a1.835 1.835 0 0 0-3.01-1.41L.768 5.86a2.135 2.135 0 0 0 0 3.28l4.721 3.935a1.835 1.835 0 0 0 3.01-1.41V10.5c.533 0 1.03.07 1.5.231Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowShapeTurnUpRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9 9v2.665a.335.335 0 0 0 .55.257l4.72-3.934a.635.635 0 0 0 0-.976L9.55 3.078a.335.335 0 0 0-.55.257V6H7.5c-1.584 0-3.182.571-4.241 1.692c-.9.951-1.549 2.446-1.31 4.723c.65-1.026 1.365-1.837 2.201-2.413C5.198 9.278 6.323 9 7.5 9zm-3 1.731c-1.162.396-2.165 1.337-3.151 3.106c-.223.4-.64.663-1.098.663c-.552 0-1.04-.376-1.143-.917C-.598 7.237 3.678 4.5 7.499 4.5V3.335a1.835 1.835 0 0 1 3.01-1.41l4.722 3.935a2.135 2.135 0 0 1 0 3.28l-4.721 3.935a1.835 1.835 0 0 1-3.01-1.41V10.5c-.533 0-1.03.07-1.5.231" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowShapeUp(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 8.5h2.67a.33.33 0 0 0 .252-.542L8.5 2.733a.656.656 0 0 0-1.002 0l-4.42 5.225a.33.33 0 0 0 .252.542H6V13a.5.5 0 0 0 .5.5h3a.5.5 0 0 0 .5-.5zm1.5 4.5a2 2 0 0 1-2 2h-3a2 2 0 0 1-2-2v-3H3.33a1.83 1.83 0 0 1-1.398-3.01l4.42-5.225a2.156 2.156 0 0 1 3.293 0l4.42 5.225A1.83 1.83 0 0 1 12.67 10H11.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowShapeUpFromLine(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.5 7h2.471L8.472 2.724a.61.61 0 0 0-.944 0L4.029 7H6.5v3.25c0 .138.112.25.25.25h2.5a.25.25 0 0 0 .25-.25zm-2.75 5h2.501A1.75 1.75 0 0 0 11 10.25V8.5h1.026a1.474 1.474 0 0 0 1.14-2.407L9.634 1.774a2.11 2.11 0 0 0-3.266 0L2.833 6.093A1.474 1.474 0 0 0 3.973 8.5H5v1.75c0 .966.784 1.75 1.75 1.75m7.5 1.5a.75.75 0 0 1 0 1.5H1.75a.75.75 0 0 1 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowShapeUpToLine(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.5 10h2.471L8.472 5.724a.61.61 0 0 0-.944 0L4.029 10H6.5v3.25c0 .138.112.25.25.25h2.5a.25.25 0 0 0 .25-.25zM8.002 2.5h6.248a.75.75 0 0 0 0-1.5H1.75a.75.75 0 0 0 0 1.5zm0 1.5a2.11 2.11 0 0 1 1.631.774l3.534 4.319a1.474 1.474 0 0 1-1.14 2.407H11v1.75A1.75 1.75 0 0 1 9.25 15h-2.5A1.75 1.75 0 0 1 5 13.25V11.5H3.974a1.474 1.474 0 0 1-1.14-2.407l3.533-4.319c.4-.49 1-.774 1.632-.774H8Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUp(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 14.75a.75.75 0 0 1-.75-.75V3.81L4.53 6.53a.75.75 0 0 1-1.06-1.06l4-4a.75.75 0 0 1 1.06 0l4 4a.75.75 0 0 1-1.06 1.06L8.75 3.81V14a.75.75 0 0 1-.75.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpArrowDown(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.72 2.22a.75.75 0 0 1 1.06 0l3 3a.75.75 0 0 1-1.06 1.06L5 4.56v8.69a.75.75 0 0 1-1.5 0V4.56L1.78 6.28A.75.75 0 0 1 .72 5.22zM11.75 14a.75.75 0 0 1-.53-.22l-3-3a.75.75 0 1 1 1.06-1.06L11 11.44V2.75a.75.75 0 0 1 1.5 0v8.69l1.72-1.72a.75.75 0 1 1 1.06 1.06l-3 3a.75.75 0 0 1-.53.22" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpFromLine(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.47 1.22a.75.75 0 0 1 1.06 0l2.5 2.5a.75.75 0 1 1-1.06 1.06L8.75 3.56v7.69a.75.75 0 0 1-1.5 0V3.56L6.03 4.78a.75.75 0 0 1-1.06-1.06zM1.75 13.5a.75.75 0 0 0 0 1.5h12.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpFromSquare(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.47 1.22a.75.75 0 0 1 1.06 0l2.5 2.5a.75.75 0 1 1-1.06 1.06L8.75 3.56v7.69a.75.75 0 0 1-1.5 0V3.56L6.03 4.78a.75.75 0 0 1-1.06-1.06zM4.25 6.5a.75.75 0 0 1 0 1.5H4a1.5 1.5 0 0 0-1.5 1.5V12A1.5 1.5 0 0 0 4 13.5h8a1.5 1.5 0 0 0 1.5-1.5V9.5A1.5 1.5 0 0 0 12 8h-.25a.75.75 0 0 1 0-1.5H12a3 3 0 0 1 3 3V12a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3V9.5a3 3 0 0 1 3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpFromSquareSlash(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.47 1.22a.75.75 0 0 1 1.06 0l2.5 2.5a.75.75 0 1 1-1.06 1.06L8.75 3.56v3.738l-1.5-1.5V3.56L6.131 4.679l-1.06-1.06zm-.22 8.84v1.19a.75.75 0 0 0 1.453.263l1.986 1.987H4A1.5 1.5 0 0 1 2.5 12V9.5A1.5 1.5 0 0 1 4 8h1.19zM3.704 6.515L1.97 4.78a.75.75 0 0 1 1.06-1.06l11 11a.75.75 0 1 1-1.06 1.06l-.786-.785A3.11 3.11 0 0 1 12 15H4a3 3 0 0 1-3-3V9.5a3 3 0 0 1 2.704-2.986Zm9.795 5.533l1.22 1.22A2.98 2.98 0 0 0 15 12.001V9.5a3 3 0 0 0-3-3h-.25a.75.75 0 0 0 0 1.5H12a1.5 1.5 0 0 1 1.5 1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRightFromSquare(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 1.5A.75.75 0 0 0 10 3h1.94L6.97 7.97a.75.75 0 0 0 1.06 1.06L13 4.06V6a.75.75 0 0 0 1.5 0V2.25a.75.75 0 0 0-.75-.75zM7.5 3.25a.75.75 0 0 0-.75-.75H4.5a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h6a3 3 0 0 0 3-3V9.25a.75.75 0 0 0-1.5 0v2.25a1.5 1.5 0 0 1-1.5 1.5h-6A1.5 1.5 0 0 1 3 11.5v-6A1.5 1.5 0 0 1 4.5 4h2.25a.75.75 0 0 0 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpToLine(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.47 4.22a.75.75 0 0 1 1.06 0l2.5 2.5a.75.75 0 1 1-1.06 1.06L8.75 6.56v7.69a.75.75 0 0 1-1.5 0V6.56L6.03 7.78a.75.75 0 0 1-1.06-1.06zm6.78-1.72a.75.75 0 0 0 0-1.5H1.75a.75.75 0 0 0 0 1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUturnCcwDown(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.72 13.53a.75.75 0 0 0 1.06 0l3-3a.75.75 0 1 0-1.06-1.06L6 11.19V7a3.25 3.25 0 0 1 6.5 0v1A.75.75 0 0 0 14 8V7a4.75 4.75 0 1 0-9.5 0v4.19L2.78 9.47a.75.75 0 0 0-1.06 1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUturnCcwLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.47 4.72a.75.75 0 0 0 0 1.06l3 3a.75.75 0 0 0 1.06-1.06L4.81 6H9a3.25 3.25 0 0 1 0 6.5H8A.75.75 0 0 0 8 14h1a4.75 4.75 0 1 0 0-9.5H4.81l1.72-1.72a.75.75 0 0 0-1.06-1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUturnCcwRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.53 11.28a.75.75 0 0 0 0-1.06l-3-3a.75.75 0 1 0-1.06 1.06L11.19 10H7a3.25 3.25 0 0 1 0-6.5h1A.75.75 0 0 0 8 2H7a4.75 4.75 0 0 0 0 9.5h4.19l-1.72 1.72a.75.75 0 1 0 1.06 1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUturnCwDown(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.28 13.53a.75.75 0 0 1-1.06 0l-3-3a.75.75 0 1 1 1.06-1.06L10 11.19V7a3.25 3.25 0 0 0-6.5 0v1A.75.75 0 0 1 2 8V7a4.75 4.75 0 0 1 9.5 0v4.19l1.72-1.72a.75.75 0 1 1 1.06 1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUturnCwLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.47 11.28a.75.75 0 0 1 0-1.06l3-3a.75.75 0 0 1 1.06 1.06L4.81 10H9a3.25 3.25 0 0 0 0-6.5H8A.75.75 0 0 1 8 2h1a4.75 4.75 0 1 1 0 9.5H4.81l1.72 1.72a.75.75 0 1 1-1.06 1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUturnCwRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.53 4.72a.75.75 0 0 1 0 1.06l-3 3a.75.75 0 1 1-1.06-1.06L11.19 6H7a3.25 3.25 0 0 0 0 6.5h1A.75.75 0 0 1 8 14H7a4.75 4.75 0 1 1 0-9.5h4.19L9.47 2.78a.75.75 0 0 1 1.06-1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsOppositeToDots(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.28 4.78a.75.75 0 0 0 0-1.06l-2.5-2.5a.75.75 0 1 0-1.06 1.06L6.94 3.5H1.75a.75.75 0 1 0 0 1.5h5.19L5.72 6.22a.75.75 0 1 0 1.06 1.06zm-.06 3.94l-2.5 2.5a.75.75 0 0 0 0 1.06l2.5 2.5a.75.75 0 1 0 1.06-1.06L9.06 12.5h5.19a.75.75 0 0 0 0-1.5H9.06l1.22-1.22a.75.75 0 1 0-1.06-1.06M14 4.25a1.75 1.75 0 1 1-3.5 0a1.75 1.75 0 0 1 3.5 0M3.75 13.5a1.75 1.75 0 1 0 0-3.5a1.75 1.75 0 0 0 0 3.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsRotateLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1.5a6.501 6.501 0 0 1 6.445 5.649a.75.75 0 1 1-1.488.194A5.001 5.001 0 0 0 4.43 4.5h1.32a.75.75 0 0 1 0 1.5h-3A.75.75 0 0 1 2 5.25v-3a.75.75 0 1 1 1.5 0v1.06A6.48 6.48 0 0 1 8 1.5m5.25 13a.75.75 0 0 0 .75-.75v-3a.75.75 0 0 0-.75-.75h-3a.75.75 0 1 0 0 1.5h1.32a5.001 5.001 0 0 1-8.528-2.843a.75.75 0 1 0-1.487.194a6.501 6.501 0 0 0 10.945 3.84v1.059c0 .414.336.75.75.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsRotateLeftSlash(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1.5a6.47 6.47 0 0 1 2.709.59L9.552 3.245A4.996 4.996 0 0 0 8 3a4.983 4.983 0 0 0-3.57 1.5h1.32a.75.75 0 0 1 0 1.5h-3A.75.75 0 0 1 2 5.25v-3a.75.75 0 1 1 1.5 0v1.06A6.48 6.48 0 0 1 8 1.5m4.026 3.534l-6.991 6.992C5.865 12.638 6.89 13 8 13a4.983 4.983 0 0 0 3.57-1.5h-1.32a.75.75 0 0 1 0-1.5h3a.75.75 0 0 1 .75.75v3a.75.75 0 1 1-1.5 0v-1.06A6.48 6.48 0 0 1 8 14.5a6.472 6.472 0 0 1-4.035-1.404l-.935.934a.75.75 0 0 1-1.06-1.06l11-11a.75.75 0 0 1 1.06 1.06l-.934.935a6.472 6.472 0 0 1 1.349 3.184a.75.75 0 1 1-1.488.194a4.972 4.972 0 0 0-.93-2.309ZM3.043 8.657c.04.308.109.607.203.895L2.09 10.708a6.46 6.46 0 0 1-.534-1.857a.75.75 0 0 1 1.487-.194" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsRotateRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1.5a6.501 6.501 0 0 0-6.445 5.649a.75.75 0 1 0 1.488.194A5.001 5.001 0 0 1 11.57 4.5h-1.32a.75.75 0 0 0 0 1.5h3a.75.75 0 0 0 .75-.75v-3a.75.75 0 0 0-1.5 0v1.06A6.48 6.48 0 0 0 8 1.5m-5.25 13a.75.75 0 0 1-.75-.75v-3a.75.75 0 0 1 .75-.75h3a.75.75 0 0 1 0 1.5H4.43a5.001 5.001 0 0 0 8.528-2.843a.75.75 0 1 1 1.487.194A6.501 6.501 0 0 1 3.5 12.691v1.059a.75.75 0 0 1-.75.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsRotateRightSlash(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.904 3.965a6.472 6.472 0 0 0-1.349 3.184a.75.75 0 1 0 1.488.194a4.97 4.97 0 0 1 .93-2.309l6.992 6.992A4.978 4.978 0 0 1 8 13a4.983 4.983 0 0 1-3.57-1.5h1.32a.75.75 0 0 0 0-1.5h-3a.75.75 0 0 0-.75.75v3a.75.75 0 0 0 1.5 0v-1.06A6.48 6.48 0 0 0 8 14.5a6.472 6.472 0 0 0 4.035-1.404l.935.934a.75.75 0 1 0 1.06-1.06l-11-11a.75.75 0 0 0-1.06 1.06zm9.85 5.587l1.156 1.156a6.457 6.457 0 0 0 .534-1.857a.75.75 0 1 0-1.487-.194c-.04.308-.109.607-.203.895M5.291 2.09l1.157 1.157A4.996 4.996 0 0 1 8 3a4.98 4.98 0 0 1 3.57 1.5h-1.32a.75.75 0 0 0 0 1.5h3a.75.75 0 0 0 .75-.75v-3a.75.75 0 0 0-1.5 0v1.06A6.48 6.48 0 0 0 8 1.5a6.47 6.47 0 0 0-2.709.59" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsThreeRotateLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiArrows3RotateLeft0)"><path fill="currentColor" fill-rule="evenodd" d="M11.39 2.503a6.473 6.473 0 0 0-3.383-.984a6.48 6.48 0 0 0-4.515 1.77l.005-.559a.75.75 0 1 0-1.5-.013l-.022 2.5a.75.75 0 0 0 .743.756l2.497.022a.75.75 0 1 0 .013-1.5l-.817-.007a4.983 4.983 0 0 1 3.583-1.469a4.973 4.973 0 0 1 2.602.756a.75.75 0 0 0 .795-1.272Zm-9.097 8.716a6.473 6.473 0 0 1-.84-3.422a.75.75 0 1 1 1.5.053a4.955 4.955 0 0 0 .646 2.63a4.983 4.983 0 0 0 3.064 2.37l-.403-.712a.75.75 0 0 1 1.306-.738l1.229 2.173a.75.75 0 0 1-.283 1.022l-2.176 1.23a.75.75 0 1 1-.739-1.305l.487-.275a6.48 6.48 0 0 1-3.79-3.026Zm11.258.099a6.473 6.473 0 0 1-2.544 2.438a.75.75 0 0 1-.704-1.325a4.974 4.974 0 0 0 1.955-1.875a4.983 4.983 0 0 0 .52-3.838l-.415.705a.75.75 0 1 1-1.292-.762l1.267-2.15a.75.75 0 0 1 1.027-.266l2.154 1.268a.75.75 0 1 1-.761 1.293l-.483-.284a6.48 6.48 0 0 1-.724 4.796" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiArrows3RotateLeft0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsThreeRotateLeftLetterA(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiArrows3RotateLeftLetterA0)"><path fill="currentColor" fill-rule="evenodd" d="M8.007 1.52a6.466 6.466 0 0 1 3.384.983a.75.75 0 1 1-.795 1.272a4.973 4.973 0 0 0-2.602-.756a4.983 4.983 0 0 0-3.583 1.469l.817.007a.75.75 0 1 1-.013 1.5l-2.497-.022a.75.75 0 0 1-.743-.756l.022-2.5a.75.75 0 1 1 1.5.013l-.005.56a6.48 6.48 0 0 1 4.515-1.77M1.453 7.796a6.473 6.473 0 0 0 .84 3.422a6.48 6.48 0 0 0 3.791 3.026l-.487.275a.75.75 0 0 0 .739 1.306l2.176-1.231a.75.75 0 0 0 .283-1.022L7.565 11.4a.75.75 0 0 0-1.305.738l.403.712a4.983 4.983 0 0 1-3.064-2.37a4.973 4.973 0 0 1-.647-2.63a.75.75 0 1 0-1.499-.053Zm9.554 5.959a6.473 6.473 0 0 0 2.544-2.438a6.48 6.48 0 0 0 .724-4.796l.483.284a.75.75 0 1 0 .761-1.293l-2.154-1.268a.75.75 0 0 0-1.027.265L11.071 6.66a.75.75 0 0 0 1.292.762l.415-.705a4.983 4.983 0 0 1-.52 3.838a4.974 4.974 0 0 1-1.955 1.876a.75.75 0 1 0 .704 1.324M7.003 5.284a1.056 1.056 0 0 1 1.994 0l1.45 4.142a.75.75 0 0 1-1.416.496L8.9 9.543H7.1l-.132.379a.75.75 0 0 1-1.416-.496l1.45-4.142ZM8 6.975l.374 1.068h-.748z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiArrows3RotateLeftLetterA0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsThreeRotateRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiArrows3RotateRight0)"><path fill="currentColor" fill-rule="evenodd" d="M4.61 2.503a6.473 6.473 0 0 1 3.383-.984a6.48 6.48 0 0 1 4.515 1.77l-.004-.559a.75.75 0 1 1 1.5-.013l.021 2.5a.75.75 0 0 1-.743.756l-2.497.022a.75.75 0 1 1-.013-1.5l.817-.007a4.983 4.983 0 0 0-3.583-1.469a4.973 4.973 0 0 0-2.602.756a.75.75 0 0 1-.795-1.272Zm9.097 8.716a6.48 6.48 0 0 0 .84-3.422a.75.75 0 1 0-1.5.053a4.973 4.973 0 0 1-.646 2.63a4.983 4.983 0 0 1-3.064 2.37l.403-.712a.75.75 0 0 0-1.306-.738l-1.229 2.173a.75.75 0 0 0 .283 1.022l2.176 1.23a.75.75 0 1 0 .739-1.305l-.487-.275a6.48 6.48 0 0 0 3.79-3.026Zm-11.258.099a6.473 6.473 0 0 0 2.544 2.438a.75.75 0 0 0 .704-1.325a4.973 4.973 0 0 1-1.955-1.875a4.983 4.983 0 0 1-.52-3.838l.415.705a.75.75 0 1 0 1.292-.762l-1.267-2.15a.75.75 0 0 0-1.027-.266L.481 5.513a.75.75 0 1 0 .761 1.293l.483-.284a6.48 6.48 0 0 0 .724 4.796" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiArrows3RotateRight0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func At(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 2.5a5.5 5.5 0 1 0 3.52 9.726a.75.75 0 0 1 .96 1.153A7 7 0 1 1 15 8a2.5 2.5 0 0 1-4.083 1.935A3.5 3.5 0 1 1 11.5 8a1 1 0 1 0 2 0A5.5 5.5 0 0 0 8 2.5M10 8a2 2 0 1 0-4 0a2 2 0 0 0 4 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ban(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.323 12.383a5.5 5.5 0 0 1-7.706-7.706zm1.06-1.06L4.677 3.617a5.5 5.5 0 0 1 7.706 7.706M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bars(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.25 3.25A.75.75 0 0 1 2 2.5h12A.75.75 0 0 1 14 4H2a.75.75 0 0 1-.75-.75m0 4.75A.75.75 0 0 1 2 7.25h12a.75.75 0 0 1 0 1.5H2A.75.75 0 0 1 1.25 8M2 12a.75.75 0 0 0 0 1.5h12a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsAscendingAlignCenter(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 12.75c0 .414.336.75.75.75h12.5a.75.75 0 0 0 0-1.5H1.75a.75.75 0 0 0-.75.75M3 8c0 .414.336.75.75.75h8.5a.75.75 0 0 0 0-1.5h-8.5A.75.75 0 0 0 3 8m3.75-4a.75.75 0 0 1 0-1.5h2.5a.75.75 0 0 1 0 1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsAscendingAlignLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 12.75c0 .414.336.75.75.75h12.5a.75.75 0 0 0 0-1.5H1.75a.75.75 0 0 0-.75.75M1 8c0 .414.336.75.75.75h8.5a.75.75 0 0 0 0-1.5h-8.5A.75.75 0 0 0 1 8m.75-4a.75.75 0 0 1 0-1.5h2.5a.75.75 0 0 1 0 1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsAscendingAlignLeftArrowDown(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.22 13.28a.75.75 0 0 0 1.06 0l2-2a.75.75 0 1 0-1.06-1.06l-.72.72V3.25a.75.75 0 0 0-1.5 0v7.69l-.72-.72a.75.75 0 1 0-1.06 1.06zM7.75 12a.75.75 0 0 0 0 1.5h7.5a.75.75 0 0 0 0-1.5zm0-3.25a.75.75 0 0 1 0-1.5h5.5a.75.75 0 0 1 0 1.5zm0-4.75a.75.75 0 0 1 0-1.5h2.5a.75.75 0 0 1 0 1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsAscendingAlignLeftArrowUp(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.22 2.72a.75.75 0 0 1 1.06 0l2 2a.75.75 0 0 1-1.06 1.06l-.72-.72v7.69a.75.75 0 0 1-1.5 0V5.06l-.72.72A.75.75 0 0 1 .22 4.72zM7 12.75c0 .414.336.75.75.75h7.5a.75.75 0 0 0 0-1.5h-7.5a.75.75 0 0 0-.75.75m.75-4a.75.75 0 0 1 0-1.5h5.5a.75.75 0 0 1 0 1.5zm0-4.75a.75.75 0 0 1 0-1.5h2.5a.75.75 0 0 1 0 1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsAscendingAlignRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 12.75a.75.75 0 0 1-.75.75H1.75a.75.75 0 0 1 0-1.5h12.5a.75.75 0 0 1 .75.75M15 8a.75.75 0 0 1-.75.75h-8.5a.75.75 0 0 1 0-1.5h8.5A.75.75 0 0 1 15 8m-.75-4a.75.75 0 0 0 0-1.5h-2.5a.75.75 0 0 0 0 1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsDescendingAlignCenter(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 3.25a.75.75 0 0 1 .75-.75h12.5a.75.75 0 0 1 0 1.5H1.75A.75.75 0 0 1 1 3.25M3 8a.75.75 0 0 1 .75-.75h8.5a.75.75 0 0 1 0 1.5h-8.5A.75.75 0 0 1 3 8m3.75 4a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsDescendingAlignLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 3.25a.75.75 0 0 1 .75-.75h12.5a.75.75 0 0 1 0 1.5H1.75A.75.75 0 0 1 1 3.25M1 8a.75.75 0 0 1 .75-.75h8.5a.75.75 0 0 1 0 1.5h-8.5A.75.75 0 0 1 1 8m.75 4a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsDescendingAlignLeftArrowDown(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.22 13.28a.75.75 0 0 0 1.06 0l2-2a.75.75 0 1 0-1.06-1.06l-.72.72V3.25a.75.75 0 0 0-1.5 0v7.69l-.72-.72a.75.75 0 1 0-1.06 1.06zM7 3.25a.75.75 0 0 1 .75-.75h7.5a.75.75 0 0 1 0 1.5h-7.5A.75.75 0 0 1 7 3.25m.75 4a.75.75 0 0 0 0 1.5h5.5a.75.75 0 0 0 0-1.5zm0 4.75a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsDescendingAlignLeftArrowUp(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.22 2.72a.75.75 0 0 1 1.06 0l2 2a.75.75 0 0 1-1.06 1.06l-.72-.72v7.69a.75.75 0 0 1-1.5 0V5.06l-.72.72A.75.75 0 0 1 .22 4.72zM7.75 4a.75.75 0 0 1 0-1.5h7.5a.75.75 0 0 1 0 1.5zm0 3.25a.75.75 0 0 0 0 1.5h5.5a.75.75 0 0 0 0-1.5zm0 4.75a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsDescendingAlignRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 3.25a.75.75 0 0 0-.75-.75H1.75a.75.75 0 0 0 0 1.5h12.5a.75.75 0 0 0 .75-.75M15 8a.75.75 0 0 0-.75-.75h-8.5a.75.75 0 0 0 0 1.5h8.5A.75.75 0 0 0 15 8m-.75 4a.75.75 0 0 1 0 1.5h-2.5a.75.75 0 0 1 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsPlay(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.75 3.25a.75.75 0 0 1 .75-.75h11a.75.75 0 0 1 0 1.5h-11a.75.75 0 0 1-.75-.75m.75 3a.75.75 0 0 0 0 1.5h11a.75.75 0 0 0 0-1.5zm0 3.75a.75.75 0 0 0 0 1.5h4.25a.75.75 0 0 0 0-1.5zm11 3.116a1 1 0 0 0 0-1.732l-3-1.732a1 1 0 0 0-1.5.866v3.464a1 1 0 0 0 1.5.866z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsUnaligned(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.75 2.5a.75.75 0 0 0 0 1.5h10.5a.75.75 0 0 0 0-1.5zM4 8a.75.75 0 0 1 .75-.75h10.5a.75.75 0 0 1 0 1.5H4.75A.75.75 0 0 1 4 8m-4 4.75A.75.75 0 0 1 .75 12h10.5a.75.75 0 0 1 0 1.5H.75a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m4.665 7.822l.62-3.096a2.77 2.77 0 0 1 5.43 0l.62 3.096a4.782 4.782 0 0 0 1.305 2.44l.194.193a.567.567 0 0 1-.273.953l-.821.19a16.63 16.63 0 0 1-7.48 0l-.82-.19a.567.567 0 0 1-.274-.953l.194-.193a4.774 4.774 0 0 0 1.305-2.44m-1.47-.294l.619-3.096a4.27 4.27 0 0 1 8.372 0l.62 3.096c.126.634.438 1.216.895 1.673l.194.194a2.066 2.066 0 0 1-.997 3.475l-.821.19c-.701.16-1.41.28-2.12.358a2 2 0 0 1-3.913 0a18.134 18.134 0 0 1-2.12-.359l-.822-.19a2.067 2.067 0 0 1-.997-3.474L2.3 9.2c.457-.457.769-1.04.895-1.673Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellDot(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8 2.5a2.77 2.77 0 0 0-2.716 2.226l-.619 3.096a4.774 4.774 0 0 1-1.305 2.44l-.194.193a.566.566 0 0 0 .273.953l.821.19a16.63 16.63 0 0 0 7.48 0l.82-.19a.567.567 0 0 0 .274-.953l-.194-.193a4.774 4.774 0 0 1-1.305-2.44l-.068-.337a.75.75 0 0 1 1.471-.295l.068.338c.126.634.438 1.216.895 1.673l.194.194a2.066 2.066 0 0 1-.997 3.475l-.821.19c-.701.16-1.41.28-2.12.358a2 2 0 0 1-3.913 0a18.134 18.134 0 0 1-2.12-.359l-.822-.19a2.067 2.067 0 0 1-.997-3.474L2.3 9.2c.457-.457.769-1.04.895-1.673l.62-3.096a4.27 4.27 0 0 1 5.4-3.256a.75.75 0 1 1-.427 1.438A2.766 2.766 0 0 0 8 2.5"/><path d="M12.5 5.5a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellSlash(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.74 4.8L1.97 3.03a.75.75 0 0 1 1.06-1.06l11 11a.75.75 0 1 1-1.06 1.06l-.957-.956c-.68.154-1.367.269-2.057.344a2 2 0 0 1-3.912 0a18.134 18.134 0 0 1-2.12-.359l-.822-.19a2.067 2.067 0 0 1-.997-3.474L2.3 9.2c.457-.457.769-1.04.895-1.673zm6.996 6.997a16.632 16.632 0 0 1-6.476-.2l-.82-.189a.567.567 0 0 1-.274-.953l.194-.193a4.774 4.774 0 0 0 1.305-2.44l.35-1.747zm.599-3.975c.028.14.061.277.101.412l3.025 3.024a2.068 2.068 0 0 0-.566-1.863L13.7 9.2a3.274 3.274 0 0 1-.895-1.673l-.62-3.096a4.27 4.27 0 0 0-6.96-2.408L6.292 3.09a2.77 2.77 0 0 1 4.424 1.637z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Binoculars(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiBinoculars0)"><path fill="currentColor" fill-rule="evenodd" d="M1.522 10.68a2.14 2.14 0 0 0 2.11 1.82a2.125 2.125 0 1 0-2.11-1.82M3.64 14h-.015a3.626 3.626 0 0 1-3.558-2.927a3.6 3.6 0 0 1 .256-2.212L2.98 2.98a2.516 2.516 0 0 1 4.802 1.237L7.673 5.6a1.474 1.474 0 0 1 .655 0l-.11-1.382a2.516 2.516 0 0 1 4.801-1.237l2.658 5.88a3.6 3.6 0 0 1 .256 2.213A3.626 3.626 0 0 1 12.375 14h-.015a3.64 3.64 0 0 1-3.628-3.35l-.174-2.176A.963.963 0 0 0 8 8.312a.963.963 0 0 0-.558.162l-.174 2.176A3.64 3.64 0 0 1 3.64 14m-.015-7.25c.91 0 1.742.336 2.379.89l.283-3.542a1.016 1.016 0 0 0-1.94-.5L2.89 6.825a3.64 3.64 0 0 1 .736-.075Zm10.853 3.93a2.125 2.125 0 1 0-2.11 1.82a2.14 2.14 0 0 0 2.11-1.82m-2.826-7.082l1.459 3.227a3.61 3.61 0 0 0-3.115.815l-.283-3.542a1.016 1.016 0 0 1 1.94-.5Z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiBinoculars0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bold(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.25 2.25A.75.75 0 0 0 3.5 3v10c0 .414.336.75.75.75H9.5a3.25 3.25 0 0 0 1.477-6.146A3.25 3.25 0 0 0 8.5 2.25zm3.5 5a1.75 1.75 0 1 0 0-3.5h-2v3.5zm-2 1.5v3.5h3a1.75 1.75 0 1 0 0-3.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Book(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 11.937V13a.5.5 0 0 0 .5.5h7a1.5 1.5 0 0 0 1.5-1.5v-.401A2.987 2.987 0 0 1 11 12H4c-.173 0-.34-.022-.5-.063M2 10V3a1.987 1.987 0 0 1 .686-1.508A1.994 1.994 0 0 1 4 1h7a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H4a2 2 0 0 1-2-2zm1.5 0a.5.5 0 0 0 .5.5h7A1.5 1.5 0 0 0 12.5 9V4A1.5 1.5 0 0 0 11 2.5H4a.498.498 0 0 0-.5.5zm2-4.75a.75.75 0 0 1 .75-.75h3.5a.75.75 0 0 1 0 1.5h-3.5a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookOpen(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.345 2.634c.09.046.18.094.268.145L8 3l.387-.221a6.699 6.699 0 0 1 6.32-.175c.486.242.793.74.793 1.283v8.938c0 .65-.526 1.175-1.175 1.175h-.04c-.187 0-.37-.05-.529-.146a4.793 4.793 0 0 0-4.61-.177l-.199.1A2.123 2.123 0 0 1 8 14h-.117a1.624 1.624 0 0 1-.726-.171l-.233-.117a4.937 4.937 0 0 0-4.748.183a.741.741 0 0 1-.381.105h-.12A1.175 1.175 0 0 1 .5 12.825V3.887c0-.544.307-1.04.793-1.284a6.699 6.699 0 0 1 6.052.03Zm1.405 9.572V4.3l.382-.218A5.199 5.199 0 0 1 14 3.927v8.357a6.293 6.293 0 0 0-5.25-.078m-1.5.005V4.298l-.382-.218A5.199 5.199 0 0 0 2 3.927v8.365a6.437 6.437 0 0 1 5.25-.082Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m8 9.524l.976.837l2.988 2.56a.325.325 0 0 0 .536-.246V4.5A1.5 1.5 0 0 0 11 3H5a1.5 1.5 0 0 0-1.5 1.5v8.175a.325.325 0 0 0 .536.247l2.988-2.56zM14 4.5a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v8.175a1.825 1.825 0 0 0 3.013 1.386L8 11.5l2.987 2.56A1.825 1.825 0 0 0 14 12.676z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 1.5a3 3 0 0 1 3 3v8.546a1.454 1.454 0 0 1-2.411 1.094L8 11l-3.589 3.14A1.454 1.454 0 0 1 2 13.046V4.5a3 3 0 0 1 3-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Box(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 10.421V5.475l-2 .714V8.25a.75.75 0 0 1-1.5 0V6.725l-2.25.804v6.088l4.777-1.792a1.5 1.5 0 0 0 .973-1.404m-2.254-5.734l1.6-.571a1.5 1.5 0 0 0-.175-.104L9.499 2.427a1.5 1.5 0 0 0-1.197-.063l-.941.353l3.724 1.862c.06.03.113.066.16.108ZM5.444 3.435l3.878 1.94l-2.273.811l-3.805-1.903c.072-.042.149-.078.23-.109zm.806 4.029L2.5 5.589v5.057a1.5 1.5 0 0 0 .83 1.341l2.92 1.46zM1 5.579c0-.436.094-.856.266-1.236a.752.752 0 0 1 .2-.37c.342-.54.855-.968 1.48-1.203L7.777.96a3 3 0 0 1 2.394.125l3.172 1.586A3 3 0 0 1 15 5.354v5.067a3 3 0 0 1-1.947 2.809l-4.828 1.81a3 3 0 0 1-2.395-.125l-3.172-1.586A3 3 0 0 1 1 10.646z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxesThree(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.25 2H6.5a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1h-.75v1.75a.75.75 0 0 1-1.5 0zM6.5.5A2.5 2.5 0 0 0 4 3v3c0 .356.074.694.208 1H3A2.5 2.5 0 0 0 .5 9.5v3A2.5 2.5 0 0 0 3 15h10a2.5 2.5 0 0 0 2.5-2.5v-3A2.5 2.5 0 0 0 13 7h-1.208c.134-.306.208-.644.208-1V3A2.5 2.5 0 0 0 9.5.5zm.75 9a1 1 0 0 0-1-1h-1v1.75a.75.75 0 0 1-1.5 0V8.5H3a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h3.25a1 1 0 0 0 1-1zm1.5 3a1 1 0 0 0 1 1H13a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1h-.75v1.75a.75.75 0 0 1-1.5 0V8.5h-1a1 1 0 0 0-1 1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BranchesDown(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.25 12a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5m-.75-1.372a2.251 2.251 0 1 0 1.5 0v-.378a3 3 0 0 0-3-3H8.75V5.372a2.25 2.25 0 1 0-1.5 0V7.25H5a3 3 0 0 0-3 3v.378a2.251 2.251 0 1 0 1.5 0v-.378A1.5 1.5 0 0 1 5 8.75h2.25v1.878a2.251 2.251 0 1 0 1.5 0V8.75H11a1.5 1.5 0 0 1 1.5 1.5zM2.75 12a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5m4.5.75a.75.75 0 1 1 1.5 0a.75.75 0 0 1-1.5 0M8 2.5A.75.75 0 1 0 8 4a.75.75 0 0 0 0-1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BranchesRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.75a.75.75 0 1 0 1.5 0a.75.75 0 0 0-1.5 0m-1.372.75a2.251 2.251 0 1 0 0-1.5h-.378a3 3 0 0 0-3 3v2.25H5.372a2.25 2.25 0 1 0 0 1.5H7.25V11a3 3 0 0 0 3 3h.378a2.251 2.251 0 1 0 0-1.5h-.378a1.5 1.5 0 0 1-1.5-1.5V8.75h1.878a2.251 2.251 0 1 0 0-1.5H8.75V5a1.5 1.5 0 0 1 1.5-1.5zM12 13.25a.75.75 0 1 0 1.5 0a.75.75 0 0 0-1.5 0m.75-4.5a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5M2.5 8A.75.75 0 1 0 4 8a.75.75 0 0 0-1.5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BranchesRightArrowRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.75a.75.75 0 1 0 1.5 0a.75.75 0 0 0-1.5 0M12.75.5A2.25 2.25 0 0 0 10.628 2h-.378a3 3 0 0 0-3 3v2.19L5.03 4.97a.75.75 0 0 0-1.06 1.06l1.22 1.22H1.75a.75.75 0 0 0 0 1.5h3.44L3.97 9.97a.75.75 0 1 0 1.06 1.06l2.22-2.22V11a3 3 0 0 0 3 3h.378a2.251 2.251 0 1 0 0-1.5h-.378a1.5 1.5 0 0 1-1.5-1.5V8.75h1.878a2.251 2.251 0 1 0 0-1.5H8.75V5a1.5 1.5 0 0 1 1.5-1.5h.378a2.251 2.251 0 1 0 2.122-3M12 13.25a.75.75 0 1 0 1.5 0a.75.75 0 0 0-1.5 0m.75-4.5a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Briefcase(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9 2.5H7a.5.5 0 0 0-.5.5v1h3V3a.5.5 0 0 0-.5-.5M5 3v1H4a3 3 0 0 0-3 3v4a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3h-1V3a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2m4.5 2.5H12A1.5 1.5 0 0 1 13.5 7v4a1.5 1.5 0 0 1-1.5 1.5H4A1.5 1.5 0 0 1 2.5 11V7A1.5 1.5 0 0 1 4 5.5zM4.75 7a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BroadcastSignal(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.442 13.033c-.278.307-.319.777-.05 1.092c.27.314.747.353 1.033.053a7.5 7.5 0 1 0-10.85 0c.286.3.763.261 1.032-.053c.27-.315.23-.785-.05-1.092a6 6 0 1 1 8.884 0Zm-.987-1.15c-.265.318-.745.279-1.015-.036c-.27-.314-.223-.784.015-1.123a3 3 0 1 0-4.91 0c.238.339.284.809.015 1.123c-.27.315-.75.354-1.015.036a4.5 4.5 0 1 1 6.91 0M8 10.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BroomMotion(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiBroomMotion0)"><path fill="currentColor" fill-rule="evenodd" d="M11.995.667a.75.75 0 1 0-1.49.166L11.19 7h-.867c-1.64 0-2.896 1.449-3.197 3.06a12.556 12.556 0 0 1-1.2 3.44C5.434 14.448 5 15 5 15h8.5s2.08-1.734 2.488-5.49C16.14 8.094 14.91 7 13.488 7H12.7zM3.75 2.5a.75.75 0 1 0 0 1.5h4.5a.75.75 0 0 0 0-1.5zM.75 6a.75.75 0 1 0 0 1.5h5.5a.75.75 0 0 0 0-1.5zM1 10.25a.75.75 0 0 1 .75-.75h2.5a.75.75 0 0 1 0 1.5h-2.5a.75.75 0 0 1-.75-.75m6.593 3.25c.393-.866.78-1.94 1.008-3.165C8.819 9.167 9.646 8.5 10.322 8.5h3.167c.332 0 .618.13.797.303a.63.63 0 0 1 .21.545c-.175 1.622-.708 2.779-1.173 3.514a6.003 6.003 0 0 1-.461.638h-.999c.539-.706.887-1.728.887-2.75H12c-.351 1.229-1.072 2.088-2.162 2.75z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiBroomMotion0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bucket(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.615 4.888c.814-.375.885-.714.885-.888c0-.174-.071-.513-.885-.888C10.8 2.737 9.538 2.5 8 2.5c-1.538 0-2.799.237-3.615.612c-.814.375-.885.714-.885.888c0 .174.071.513.885.888C5.2 5.263 6.462 5.5 8 5.5c1.538 0 2.799-.237 3.615-.612M8 7c1.688 0 3.213-.26 4.304-.779a12.67 12.67 0 0 1-.085.431l-1.18 5.51a.86.86 0 0 1-.039.094a1.544 1.544 0 0 1-.31.433c-.306.306-1.03.811-2.69.811s-2.384-.505-2.69-.81a1.535 1.535 0 0 1-.31-.434a.867.867 0 0 1-.038-.095l-.022-.1c.109-.017.219-.038.33-.063c1.518-.331 3.284-1.253 5.221-2.931a.75.75 0 1 0-.982-1.134c-1.813 1.572-3.36 2.338-4.56 2.6c-.11.024-.219.044-.324.06l-.844-3.94c-.03-.144-.059-.287-.085-.432C4.786 6.741 6.312 7 8 7m-4.93 3.496l-.697-3.254c-.28.426-.434.846-.49 1.227c-.107.72.13 1.323.592 1.7c.158.13.356.244.596.327ZM2.062 5.314A14.027 14.027 0 0 1 2 4c0-2 2.686-3 6-3s6 1 6 3c0 .997-.105 1.992-.314 2.967L12.5 12.5S12 15 8 15s-4.5-2.5-4.5-2.5l-.084-.394c-.736-.091-1.378-.356-1.89-.775C.611 10.583.223 9.434.398 8.25c.15-1.02.708-2.049 1.662-2.936Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BucketPaint(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.773 7.412c-.064.064-.27.249-1.017-.027c-.75-.277-1.706-.928-2.695-1.918c-.99-.99-1.64-1.945-1.918-2.695c-.276-.747-.091-.953-.027-1.017c.064-.064.27-.249 1.017.027c.094.035.19.075.29.121c.7.324 1.54.93 2.405 1.797c.99.99 1.641 1.945 1.918 2.695c.276.747.091.953.027 1.017M7 6.528c.85.85 1.738 1.535 2.581 1.972H1.694v-.027a1.292 1.292 0 0 1 .1-.507l2.802-4.33c.056-.087.114-.174.172-.26C5.16 4.383 5.956 5.485 7 6.529Zm3.89-3.889c2.147 2.148 3.24 4.537 1.944 5.834a12.976 12.976 0 0 1-2.127 1.719L6.352 13.01s-1.945 1.296-4.537-1.296C-.778 9.12.518 7.175.518 7.175L3.336 2.82A12.98 12.98 0 0 1 5.056.694C6.351-.602 8.74.491 10.888 2.64Zm1.86 12.36c.966 0 1.75-.765 1.75-1.71c0-1.234-1.17-2.76-1.512-3.178a.304.304 0 0 0-.237-.111a.31.31 0 0 0-.24.112c-.34.422-1.511 1.96-1.511 3.178c0 .944.784 1.71 1.75 1.71Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bug(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiBug0)"><path fill="currentColor" fill-rule="evenodd" d="M5.865.031a.75.75 0 0 1 .918.53l.531 1.981a5.554 5.554 0 0 1 2.384.225a2.5 2.5 0 1 1 3.535 3.535a5.496 5.496 0 0 1 .225 2.384l1.98.53a.75.75 0 0 1-.388 1.45l-1.98-.53a5.49 5.49 0 0 1-.687 1.187l1.45 1.45a.75.75 0 0 1-1.06 1.06l-1.45-1.45a5.493 5.493 0 0 1-1.188.687l.53 1.98a.75.75 0 1 1-1.448.388l-.531-1.98a5.5 5.5 0 0 1-6.144-6.143l-1.98-.532A.75.75 0 0 1 .95 5.334l1.98.531c.18-.426.411-.824.687-1.188l-1.45-1.45a.75.75 0 1 1 1.06-1.06l1.45 1.45a5.494 5.494 0 0 1 1.188-.687L5.335.95a.75.75 0 0 1 .53-.919M8 12a4 4 0 1 0-3.309-1.752L8.42 6.52a.75.75 0 0 1 1.06 1.06l-3.728 3.73c.64.435 1.414.69 2.248.69" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiBug0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bulb(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M6.26 15.109a4 4 0 0 0 3.48 0l.13-.063a2 2 0 0 0 1.13-1.8v-.468c0-1.352.776-2.557 1.54-3.673a5.5 5.5 0 1 0-9.08 0C4.224 10.221 5 11.425 5 12.778v.467a2 2 0 0 0 1.13 1.801zm2.828-1.35l.13-.064a.5.5 0 0 0 .282-.45v-.467c0-.17.009-.337.025-.5a5.328 5.328 0 0 1-3.05 0c.016.163.025.33.025.5v.467a.5.5 0 0 0 .282.45l.13.063a2.5 2.5 0 0 0 2.176 0Zm-4.39-5.501c.394.576.891 1.302 1.263 2.148a3.793 3.793 0 0 0 4.078 0c.372-.846.869-1.572 1.264-2.148a4 4 0 1 0-6.605 0" clip-rule="evenodd"/><path d="M8 3.5A.75.75 0 0 0 8 5a1 1 0 0 1 1 1a.75.75 0 0 0 1.5 0A2.5 2.5 0 0 0 8 3.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calculator(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 12V4A1.5 1.5 0 0 1 5 2.5h6A1.5 1.5 0 0 1 12.5 4v8a1.5 1.5 0 0 1-1.5 1.5H5A1.5 1.5 0 0 1 3.5 12M5 15a3 3 0 0 1-3-3V4a3 3 0 0 1 3-3h6a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3zm.5-11a1 1 0 1 0 0 2h5a1 1 0 1 0 0-2zm1 7a1 1 0 1 1-2 0a1 1 0 0 1 2 0M8 12a1 1 0 1 0 0-2a1 1 0 0 0 0 2m3.5-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-6-2a1 1 0 1 0 0-2a1 1 0 0 0 0 2M9 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0m1.5 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.25 5.497a.75.75 0 0 1-.75-.75V4A1.5 1.5 0 0 0 3 5.5v1h10v-1A1.5 1.5 0 0 0 11.5 4v.75a.75.75 0 0 1-1.5 0V4H6v.747a.75.75 0 0 1-.75.75M10 2.5H6v-.752a.75.75 0 1 0-1.5 0V2.5a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h7a3 3 0 0 0 3-3v-6a3 3 0 0 0-3-3v-.75a.75.75 0 0 0-1.5 0zM3 8v3.5A1.5 1.5 0 0 0 4.5 13h7a1.5 1.5 0 0 0 1.5-1.5V8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.273 5h1.05l.36-.987l.248-.684A.5.5 0 0 1 6.401 3h3.198a.5.5 0 0 1 .47.33l.249.683l.359.987H12a1.5 1.5 0 0 1 1.5 1.5V11a1.5 1.5 0 0 1-1.5 1.5H4A1.5 1.5 0 0 1 2.5 11V6.5A1.5 1.5 0 0 1 4 5zM6.4 1.5a2 2 0 0 0-1.88 1.317l-.248.683H4a3 3 0 0 0-3 3V11a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V6.5a3 3 0 0 0-3-3h-.273l-.248-.683A2 2 0 0 0 9.599 1.5zm3.099 7a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0m1.5 0a3 3 0 1 1-6 0a3 3 0 0 1 6 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Car(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.5 11H4.419l-.342 1.026l-.158.474H2V9.52c.496.129 1.213.23 2.25.23a.75.75 0 1 0 0-1.5c-1.073 0-1.682-.12-1.998-.217a1.765 1.765 0 0 1-.204-.075a1.82 1.82 0 0 1 .485-.87c.074-.073.145-.15.214-.228C4.272 7.293 6.15 7.5 8 7.5c1.849 0 3.728-.207 5.253-.64c.069.079.14.155.214.228c.241.242.408.544.485.87c-.045.02-.111.047-.204.075c-.316.097-.925.217-1.998.217a.75.75 0 0 0 0 1.5c1.037 0 1.754-.101 2.25-.23v2.98h-1.919l-.158-.474L11.581 11zm6.924-5.472C11.144 5.838 9.584 6 8 6c-1.584 0-3.144-.162-4.424-.472c.031-.075.06-.15.088-.226l.448-1.257c.18-.505.57-.806.96-.863a20.756 20.756 0 0 1 5.855 0c.392.057.78.358.96.863l.45 1.257c.027.076.056.151.087.226m-1.652 7.788L10.5 12.5h-5l-.272.816a1 1 0 0 1-.949.684H1.5a1 1 0 0 1-1-1V8.375c0-.88.35-1.725.972-2.347a3.32 3.32 0 0 0 .43-.528H1.25a.75.75 0 1 1 0-1.5h1.286l.164-.46c.343-.96 1.148-1.696 2.157-1.842a22.256 22.256 0 0 1 6.286 0c1.009.146 1.814.882 2.157 1.843l.164.459h1.286a.75.75 0 0 1 0 1.5h-.651c.124.19.268.367.429.528a3.32 3.32 0 0 1 .972 2.347V13a1 1 0 0 1-1 1h-2.78a1 1 0 0 1-.948-.684" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardClub(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 4v8a1.5 1.5 0 0 1-1.5 1.5H5A1.5 1.5 0 0 1 3.5 12V4A1.5 1.5 0 0 1 5 2.5h6A1.5 1.5 0 0 1 12.5 4M11 1a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V4a3 3 0 0 1 3-3zM9.5 6c0 .176-.03.345-.086.502L9.5 6.5a1.5 1.5 0 1 1-1.228 2.361c.08.667.382 1.293.862 1.773L9.5 11v.5h-3V11l.366-.366c.48-.48.782-1.106.862-1.773a1.5 1.5 0 1 1-1.142-2.359A1.5 1.5 0 1 1 9.5 6" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardDiamond(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 4v8a1.5 1.5 0 0 1-1.5 1.5H5A1.5 1.5 0 0 1 3.5 12V4A1.5 1.5 0 0 1 5 2.5h6A1.5 1.5 0 0 1 12.5 4M11 1a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V4a3 3 0 0 1 3-3zM5 8.25v-.5A5.564 5.564 0 0 0 7.75 4h.5A5.564 5.564 0 0 0 11 7.75v.5A5.564 5.564 0 0 0 8.25 12h-.5A5.564 5.564 0 0 0 5 8.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardHeart(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 4v8a1.5 1.5 0 0 1-1.5 1.5H5A1.5 1.5 0 0 1 3.5 12V4A1.5 1.5 0 0 1 5 2.5h6A1.5 1.5 0 0 1 12.5 4M11 1a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V4a3 3 0 0 1 3-3zM4.75 7.333c0-1.148.734-2.084 1.846-2.084c.613 0 1.253.595 1.404 1.501c.15-.915.791-1.5 1.404-1.5c1.112 0 1.846.935 1.846 2.083c0 1.895-1.69 3.1-3.1 3.878a.307.307 0 0 1-.3-.001c-1.412-.786-3.1-1.973-3.1-3.877" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardSpade(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 4v8a1.5 1.5 0 0 1-1.5 1.5H5A1.5 1.5 0 0 1 3.5 12V4A1.5 1.5 0 0 1 5 2.5h6A1.5 1.5 0 0 1 12.5 4M11 1a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V4a3 3 0 0 1 3-3zM9.296 9.5C10.279 9.5 11 8.803 11 7.855c0-1.308-1.38-2.246-2.464-2.983c-.19-.13-.372-.254-.536-.372c-.162.117-.342.24-.53.368C6.384 5.608 5 6.55 5 7.855C5 8.803 5.721 9.5 6.704 9.5c.388 0 .716-.124.974-.345a3.018 3.018 0 0 1-.812 1.479L6.5 11v.5h3V11l-.366-.366a3.018 3.018 0 0 1-.812-1.479c.258.221.587.345.974.345" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretDown(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.177 6.705A.73.73 0 0 1 4.729 5.5h6.542a.73.73 0 0 1 .552 1.205L8.8 10.214a1 1 0 0 1-.757.347h-.084a1 1 0 0 1-.757-.347z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.295 4.177a.73.73 0 0 1 1.205.552v6.542a.73.73 0 0 1-1.205.552L5.786 8.8a1 1 0 0 1-.347-.757v-.084a1 1 0 0 1 .347-.757l3.509-3.024Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.705 11.823a.73.73 0 0 1-1.205-.552V4.729a.73.73 0 0 1 1.205-.552L10.214 7.2a1 1 0 0 1 .347.757v.084a1 1 0 0 1-.347.757l-3.509 3.024Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretUp(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.823 9.295a.73.73 0 0 1-.552 1.205H4.729a.73.73 0 0 1-.552-1.205L7.2 5.786a1 1 0 0 1 .757-.347h.084a1 1 0 0 1 .757.347l3.024 3.509Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretsExpandVertical(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 6.273a.727.727 0 0 0-.18-.479L8.8 2.342A1 1 0 0 0 8.046 2h-.092a1 1 0 0 0-.753.341L4.18 5.794A.727.727 0 0 0 4.727 7h6.546A.727.727 0 0 0 12 6.273M4 9.727c0 .176.064.346.18.479l3.02 3.453a1 1 0 0 0 .753.341h.092a1 1 0 0 0 .753-.341l3.021-3.454A.727.727 0 0 0 11.273 9H4.727A.727.727 0 0 0 4 9.727" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartAreaStacked(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiChartAreaStacked0)"><path fill="currentColor" fill-rule="evenodd" d="M10.257 4.2a.858.858 0 0 1-.86.274l-3.103-.776a2.81 2.81 0 0 0-2.796.876L1.242 7.152A3 3 0 0 0 .5 9.127V12a2 2 0 0 0 2 2h11a2 2 0 0 0 2-2V3.309a1.933 1.933 0 0 0-3.4-1.258l-1.31 1.528l-.533.622ZM14 6.48V3.31a.433.433 0 0 0-.762-.282l-1.842 2.15a2.358 2.358 0 0 1-2.362.753L5.93 5.154a1.31 1.31 0 0 0-1.303.408L2.37 8.139a1.5 1.5 0 0 0-.37.988v.685l2.304-1.44a2.59 2.59 0 0 1 2.458-.155l1.923.888a1.578 1.578 0 0 0 1.777-.317l.22-.22l1.575-1.574A1.862 1.862 0 0 1 14 6.479ZM2 12c0 .277.226.501.5.501h11a.5.5 0 0 0 .5-.5V8.337a.4.4 0 0 0-.683-.283L11.523 9.85a3.078 3.078 0 0 1-3.466.618L6.134 9.58a1.09 1.09 0 0 0-1.035.066L2.352 11.36a.752.752 0 0 0-.352.637Z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiChartAreaStacked0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartAreaStackedNormalized(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiChartAreaStackedNormalized0)"><path fill="currentColor" fill-rule="evenodd" d="M13.5 12.5h-11A.5.5 0 0 1 2 12v-1.611l2.812-1.985a1.25 1.25 0 0 1 1.414-.019l1.753 1.169a2.75 2.75 0 0 0 3.757-.681L14 5.728V12a.5.5 0 0 1-.5.5m.22-8.95l-3.201 4.446a1.25 1.25 0 0 1-1.708.31L7.058 7.137a2.75 2.75 0 0 0-3.111.042L2 8.553V4a.5.5 0 0 1 .5-.5h11c.079 0 .153.018.22.05M.5 9.992V4a2 2 0 0 1 2-2h11a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-11a2 2 0 0 1-2-2V9.991Z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiChartAreaStackedNormalized0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartBar(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiChartBar0)"><path fill="currentColor" fill-rule="evenodd" d="M6 4v1a.5.5 0 0 1-.5.5h-3A.5.5 0 0 1 2 5V4a.5.5 0 0 1 .5-.5h3A.5.5 0 0 1 6 4M2 7.5v1a.5.5 0 0 0 .5.5h7a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5h-7a.5.5 0 0 0-.5.5M2 11v1a.5.5 0 0 0 .5.5h11a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5h-11a.5.5 0 0 0-.5.5m-1.5.503V4a2 2 0 0 1 2-2h3a2 2 0 0 1 2 2v1c0 .173-.022.34-.063.5H9.5a2 2 0 0 1 2 2v1c0 .173-.022.34-.063.5H13.5a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2h-11a2 2 0 0 1-2-2z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiChartBar0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartBarStacked(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiChartBarStacked0)"><path fill="currentColor" fill-rule="evenodd" d="M14 12a.5.5 0 0 1-.5.5h-11A.5.5 0 0 1 2 12v-1a.5.5 0 0 1 .5-.5h11a.5.5 0 0 1 .5.5zm0-3.5a.5.5 0 0 1-.5.5H12a.5.5 0 0 1-.5-.5v-1A.5.5 0 0 1 12 7h1.5a.5.5 0 0 1 .5.5zM14 5V4a.5.5 0 0 0-.5-.5H8a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5h5.5A.5.5 0 0 0 14 5M2.5 3.5A.5.5 0 0 0 2 4v1a.5.5 0 0 0 .5.5h3A.5.5 0 0 0 6 5V4a.5.5 0 0 0-.5-.5zm-.5 4v1a.5.5 0 0 0 .5.5h7a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5h-7a.5.5 0 0 0-.5.5M.5 12a2 2 0 0 0 2 2h11a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2h-11a2 2 0 0 0-2 2z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiChartBarStacked0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartColumn(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiChartColumn0)"><path fill="currentColor" fill-rule="evenodd" d="M11.5 3.5h2a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1-.5-.5V4a.5.5 0 0 1 .5-.5m-2.5 9a.5.5 0 0 0 .5-.5V7a.5.5 0 0 0-.5-.5H7a.5.5 0 0 0-.5.5v5a.5.5 0 0 0 .5.5zm-4.5 0A.5.5 0 0 0 5 12v-2a.5.5 0 0 0-.5-.5h-2a.5.5 0 0 0-.5.5v2a.5.5 0 0 0 .5.5zm-1 1.5h-1a2 2 0 0 1-2-2v-2a2 2 0 0 1 2-2h2c.173 0 .34.022.5.063V7a2 2 0 0 1 2-2h2c.173 0 .34.022.5.063V4a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiChartColumn0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartColumnStacked(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiChartColumnStacked0)"><path fill="currentColor" fill-rule="evenodd" d="M11 4a.5.5 0 0 1 .5-.5h2a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1-.5-.5zm-2 8.5a.5.5 0 0 0 .5-.5V7a.5.5 0 0 0-.5-.5H7a.5.5 0 0 0-.5.5v5a.5.5 0 0 0 .5.5zm-4.5 0A.5.5 0 0 0 5 12v-2a.5.5 0 0 0-.5-.5h-2a.5.5 0 0 0-.5.5v2a.5.5 0 0 0 .5.5zM2 7.5V4a.5.5 0 0 1 .5-.5h2A.5.5 0 0 1 5 4v3.5a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1-.5-.5m5-4h2a.5.5 0 0 1 .5.5v.5A.5.5 0 0 1 9 5H7a.5.5 0 0 1-.5-.5V4a.5.5 0 0 1 .5-.5M7 2H2.5a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h11a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiChartColumnStacked0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartDonut(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.111 2.18a7 7 0 1 1 7.778 11.64A7 7 0 0 1 4.111 2.18M8 6.02a1.98 1.98 0 1 0 0 3.96a1.98 1.98 0 0 0 0-3.96m-.75-3.469a5.5 5.5 0 1 0 6.199 6.199h-2.05a3.481 3.481 0 0 1-5.86 1.71A3.48 3.48 0 0 1 7.25 4.603v-2.05Zm1.5 0v2.05a3.48 3.48 0 0 1 2.648 2.649h2.05A5.498 5.498 0 0 0 8.75 2.551" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartLine(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiChartLine0)"><path fill="currentColor" fill-rule="evenodd" d="M15.326 1.27a.75.75 0 0 1-.096 1.056l-3.674 3.062a2.75 2.75 0 0 1-2.55.522l-2.869-.86a1.25 1.25 0 0 0-1.214.285l-3.16 2.962A.75.75 0 1 1 .737 7.203l3.16-2.963a2.75 2.75 0 0 1 2.671-.628l2.868.86c.402.121.837.032 1.16-.236l3.674-3.062a.75.75 0 0 1 1.056.096m.113 6.185a.75.75 0 0 1-.393.984l-4.398 1.885a2.75 2.75 0 0 1-2.313-.068L6.186 9.182a1.25 1.25 0 0 0-1.238.068l-3.29 2.13a.75.75 0 0 1-.815-1.26l3.29-2.129a2.75 2.75 0 0 1 2.724-.15l2.149 1.073a1.25 1.25 0 0 0 1.051.031l4.398-1.884a.75.75 0 0 1 .984.394M1.25 12.5a.75.75 0 0 0 0 1.5h13.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiChartLine0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartLineLabel(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiChartLineLabel0)"><path fill="currentColor" fill-rule="evenodd" d="M16 4a.75.75 0 0 0-1.4-.375l-1.176 2.037l-1.288-2.06A.75.75 0 0 0 10.75 4v4.25a.75.75 0 0 0 1.5 0V6.613l.564.901a.75.75 0 0 0 1.286-.022l.4-.694v1.45a.75.75 0 0 0 1.5 0zM2.828 4a.75.75 0 0 0-1.218-.586L.282 4.477a.75.75 0 1 0 .937 1.171l.11-.088v2.69a.75.75 0 1 0 1.5 0V4Zm4.583.75a.695.695 0 0 0-.656.464l-.085.242a.75.75 0 1 1-1.415-.498l.085-.242A2.195 2.195 0 0 1 7.411 3.25h.257A1.858 1.858 0 0 1 8.776 6.6l-1.213.9h1.375a.75.75 0 1 1 0 1.5H5.975a.975.975 0 0 1-.581-1.758L7.88 5.395a.358.358 0 0 0-.213-.645zm7.44 5.226a.778.778 0 1 1 .598 1.436l-4.717 1.968a2.85 2.85 0 0 1-2.374-.081l-2.263-1.132a1.296 1.296 0 0 0-1.262.057l-3.624 2.243a.778.778 0 1 1-.818-1.323l3.623-2.243a2.852 2.852 0 0 1 2.776-.126l2.264 1.132c.337.169.73.182 1.079.037zM4.75 8.25a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiChartLineLabel0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartMixed(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiChartMixed0)"><path fill="currentColor" fill-rule="evenodd" d="M14.98 1.826a.75.75 0 0 0-.96-1.152L10.376 3.71a1.25 1.25 0 0 1-1.196.225l-2.155-.718a2.75 2.75 0 0 0-2.973.837L.926 7.767a.75.75 0 1 0 1.148.966l3.125-3.712a1.25 1.25 0 0 1 1.352-.38l2.155.718a2.75 2.75 0 0 0 2.63-.496zM13.5 8h-2a.5.5 0 0 0-.5.5V13a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5V8.5a.5.5 0 0 0-.5-.5m-4 5a.5.5 0 0 1-.5.5H7a.5.5 0 0 1-.5-.5v-2.5A.5.5 0 0 1 7 10h2a.5.5 0 0 1 .5.5zM5 13a.5.5 0 0 1-.5.5h-2A.5.5 0 0 1 2 13v-.5a.5.5 0 0 1 .5-.5h2a.5.5 0 0 1 .5.5zm-2.5 2h11a2 2 0 0 0 2-2V8.5a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v.063A2.004 2.004 0 0 0 9 8.5H7a2 2 0 0 0-2 2v.063a2.005 2.005 0 0 0-.5-.063h-2a2 2 0 0 0-2 2v.5a2 2 0 0 0 2 2" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiChartMixed0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartPie(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.45 8.75a5.501 5.501 0 1 1-6.2-6.2V8c0 .414.336.75.75.75zm0-1.5h-4.7v-4.7a5.503 5.503 0 0 1 4.7 4.7M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartTreemap(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 3H4.5A1.5 1.5 0 0 0 3 4.5v4h5zm1.5 0v5.5H13v-4A1.5 1.5 0 0 0 11.5 3zM3 11.5V10h3v3H4.5A1.5 1.5 0 0 1 3 11.5m8.25 1.5H7.5v-3h3.75zM4.5 1.5a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h7a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.488 3.43a.75.75 0 0 1 .081 1.058l-6 7a.75.75 0 0 1-1.1.042l-3.5-3.5A.75.75 0 0 1 4.03 6.97l2.928 2.927l5.473-6.385a.75.75 0 0 1 1.057-.081Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cherry(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiCherry0)"><path fill="currentColor" fill-rule="evenodd" d="M7.71.058c-1.62.676-2.758 1.724-3.414 3.04l-.033.066a6.542 6.542 0 0 0-.624 2.516c-.026.493-.005.998.058 1.51l.002.01l.002.016A4.002 4.002 0 0 0 5 15a3.985 3.985 0 0 0 3.317-1.764c.038-.056.075-.114.11-.173A4.002 4.002 0 0 0 11 14a4 4 0 1 0-1.391-7.752a4.105 4.105 0 0 1-.747-.327a3.322 3.322 0 0 1-.423-.285a2.792 2.792 0 0 1-.818-1.039c-.303-.67-.355-1.519-.048-2.477c.13.362.292.693.48.995a4.52 4.52 0 0 0 2.006 1.753c.3.137.615.243.941.322a6.44 6.44 0 0 0 1.062.163c.36.025.728.024 1.1 0l.024-.002a9.551 9.551 0 0 0 .937-.111a.75.75 0 0 0 .617-.863a5.859 5.859 0 0 0-.267-1.032a4.559 4.559 0 0 0-.965-1.59a4.662 4.662 0 0 0-1.22-.918C11.068.191 9.553.001 8.018 0a.752.752 0 0 0-.307.058Zm1.605 2.247a3.028 3.028 0 0 0 1.113 1.072l.02.01c.227.125.476.226.745.303c.546.156 1.174.212 1.87.166a2.951 2.951 0 0 0-1.238-1.555a3.594 3.594 0 0 0-.238-.138a4.458 4.458 0 0 0-.707-.296C10.33 1.69 9.7 1.582 9 1.533a10.43 10.43 0 0 0-.045-.003a4.328 4.328 0 0 0 .028.08c.094.258.205.489.332.695M5.867 3.363c0 .2.011.399.036.594a4.367 4.367 0 0 0 .365 1.29a4.204 4.204 0 0 0 .679 1.035c.215.243.46.463.731.659c.14.1.288.195.442.283a4.016 4.016 0 0 0-.548.713a3.999 3.999 0 0 0-2.387-.933a7.487 7.487 0 0 1-.055-.694a6.515 6.515 0 0 1 .009-.591a5.04 5.04 0 0 1 .466-1.886l.008-.016a4.337 4.337 0 0 1 .254-.454M8.608 9.27a3.98 3.98 0 0 1 .365 2.193a2.5 2.5 0 1 0 .229-3.201l-.002.002a2.48 2.48 0 0 0-.592 1.006M7.5 11.037V11a2.5 2.5 0 1 0 0 .037" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiCherry0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDown(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.97 5.47a.75.75 0 0 1 1.06 0L8 9.44l3.97-3.97a.75.75 0 1 1 1.06 1.06l-4.5 4.5a.75.75 0 0 1-1.06 0l-4.5-4.5a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownWide(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.867 6.097a.75.75 0 0 1 1.036-.23L8 9.111l5.097-3.244a.75.75 0 0 1 .806 1.266l-5.5 3.5a.75.75 0 0 1-.806 0l-5.5-3.5a.75.75 0 0 1-.23-1.036" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.53 2.97a.75.75 0 0 1 0 1.06L6.56 8l3.97 3.97a.75.75 0 1 1-1.06 1.06l-4.5-4.5a.75.75 0 0 1 0-1.06l4.5-4.5a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.47 13.03a.75.75 0 0 1 0-1.06L9.44 8L5.47 4.03a.75.75 0 0 1 1.06-1.06l4.5 4.5a.75.75 0 0 1 0 1.06l-4.5 4.5a.75.75 0 0 1-1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUp(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.03 10.53a.75.75 0 0 1-1.06 0L8 6.56l-3.97 3.97a.75.75 0 1 1-1.06-1.06l4.5-4.5a.75.75 0 0 1 1.06 0l4.5 4.5a.75.75 0 0 1 0 1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpWide(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.867 9.903a.75.75 0 0 0 1.036.23L8 6.889l5.097 3.244a.75.75 0 0 0 .806-1.266l-5.5-3.5a.75.75 0 0 0-.806 0l-5.5 3.5a.75.75 0 0 0-.23 1.036" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronsCollapseFromLines(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 2.75a.75.75 0 0 0-1.5 0v10.5a.75.75 0 0 0 1.5 0zm14.5 0a.75.75 0 0 0-1.5 0v10.5a.75.75 0 0 0 1.5 0zM3.47 4.97a.75.75 0 0 0 0 1.06L5.44 8L3.47 9.97a.75.75 0 1 0 1.06 1.06l2.5-2.5a.75.75 0 0 0 0-1.06l-2.5-2.5a.75.75 0 0 0-1.06 0m9.06 1.06a.75.75 0 0 0-1.06-1.06l-2.5 2.5a.75.75 0 0 0 0 1.06l2.5 2.5a.75.75 0 1 0 1.06-1.06L10.56 8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronsCollapseHorizontal(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.141 3.58a.75.75 0 0 0-1.06 1.061L4.439 8l-3.353 3.354a.75.75 0 1 0 1.06 1.06L6.03 8.53a.75.75 0 0 0 0-1.06zm11.718 8.84a.75.75 0 0 0 1.06-1.061L11.561 8l3.353-3.354a.75.75 0 0 0-1.06-1.06L9.97 7.47a.75.75 0 0 0 0 1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronsCollapseToLine(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 2a.75.75 0 0 1 .75.75v10.5a.75.75 0 0 1-1.5 0V2.75A.75.75 0 0 1 8 2M2.22 4.97a.75.75 0 0 0 0 1.06L4.19 8L2.22 9.97a.75.75 0 1 0 1.06 1.06l2.5-2.5a.75.75 0 0 0 0-1.06l-2.5-2.5a.75.75 0 0 0-1.06 0m11.56 1.06a.75.75 0 0 0-1.06-1.06l-2.5 2.5a.75.75 0 0 0 0 1.06l2.5 2.5a.75.75 0 1 0 1.06-1.06L11.81 8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronsCollapseUpRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiChevronsCollapseUpRight0)"><path fill="currentColor" fill-rule="evenodd" d="M15.25 6.993a.75.75 0 0 0 0-1.5H10.5V.75a.75.75 0 1 0-1.5 0v5.493c0 .414.336.75.75.75zM.75 9.007a.75.75 0 1 0 0 1.5H5.5v4.743a.75.75 0 0 0 1.5 0V9.757a.75.75 0 0 0-.75-.75z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiChevronsCollapseUpRight0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronsCollapseVertical(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.42 2.141a.75.75 0 1 0-1.061-1.06L8 4.439L4.646 1.086a.75.75 0 1 0-1.06 1.06L7.47 6.03a.75.75 0 0 0 1.06 0zM3.58 13.86a.75.75 0 0 0 1.061 1.06L8 11.561l3.354 3.353a.75.75 0 0 0 1.06-1.06L8.53 9.97a.75.75 0 0 0-1.06 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronsDown(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.03 3.47a.75.75 0 0 0-1.06 1.06l3.5 3.5a.75.75 0 0 0 1.06 0l3.5-3.5a.75.75 0 0 0-1.06-1.06L8 6.44zm0 5a.75.75 0 0 0-1.06 1.06l3.5 3.5a.75.75 0 0 0 1.06 0l3.5-3.5a.75.75 0 0 0-1.06-1.06L8 11.44z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronsDownWide(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.903 3.617a.75.75 0 1 0-.806 1.266l5.5 3.5a.75.75 0 0 0 .806 0l5.5-3.5a.75.75 0 1 0-.806-1.266L8 6.861zm0 4.5a.75.75 0 1 0-.806 1.266l5.5 3.5a.75.75 0 0 0 .806 0l5.5-3.5a.75.75 0 1 0-.806-1.266L8 11.361z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronsExpandFromLine(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 2a.75.75 0 0 1 .75.75v10.5a.75.75 0 0 1-1.5 0V2.75A.75.75 0 0 1 8 2M4.78 4.97a.75.75 0 0 1 0 1.06L2.81 8l1.97 1.97a.75.75 0 1 1-1.06 1.06l-2.5-2.5a.75.75 0 0 1 0-1.06l2.5-2.5a.75.75 0 0 1 1.06 0m6.44 1.06a.75.75 0 0 1 1.06-1.06l2.5 2.5a.75.75 0 0 1 0 1.06l-2.5 2.5a.75.75 0 1 1-1.06-1.06L13.19 8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronsExpandHorizontal(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiChevronsExpandHorizontal0)"><path fill="currentColor" fill-rule="evenodd" d="M11.891 3.58a.75.75 0 0 0-1.06 1.061L14.188 8l-3.353 3.354a.75.75 0 1 0 1.06 1.06L15.78 8.53a.75.75 0 0 0 0-1.06zM4.11 12.42a.75.75 0 0 0 1.06-1.061L1.811 8l3.353-3.354a.75.75 0 1 0-1.06-1.06L.22 7.47a.75.75 0 0 0 0 1.06l3.889 3.89Z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiChevronsExpandHorizontal0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronsExpandToLines(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 2.75a.75.75 0 0 0-1.5 0v10.5a.75.75 0 0 0 1.5 0zm14.5 0a.75.75 0 0 0-1.5 0v10.5a.75.75 0 0 0 1.5 0zM6.53 4.97a.75.75 0 0 1 0 1.06L4.56 8l1.97 1.97a.75.75 0 1 1-1.06 1.06l-2.5-2.5a.75.75 0 0 1 0-1.06l2.5-2.5a.75.75 0 0 1 1.06 0m2.94 1.06a.75.75 0 0 1 1.06-1.06l2.5 2.5a.75.75 0 0 1 0 1.06l-2.5 2.5a.75.75 0 1 1-1.06-1.06L11.44 8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronsExpandUpRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.754 2.004a.75.75 0 0 0 0 1.5h4.75v4.742a.75.75 0 0 0 1.5 0V2.754a.75.75 0 0 0-.75-.75zm.492 11.992a.75.75 0 0 0 0-1.5h-4.75V7.754a.75.75 0 0 0-1.5 0v5.492a.75.75 0 0 0 .75.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronsExpandVertical(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiChevronsExpandVertical0)"><path fill="currentColor" fill-rule="evenodd" d="M3.58 4.109a.75.75 0 0 0 1.061 1.06L8 1.811l3.354 3.353a.75.75 0 0 0 1.06-1.06L8.53.22a.75.75 0 0 0-1.06 0zm8.84 7.782a.75.75 0 1 0-1.061-1.06l-3.36 3.358l-3.353-3.353a.75.75 0 1 0-1.06 1.06L7.47 15.78a.75.75 0 0 0 1.06 0z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiChevronsExpandVertical0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronsLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.53 5.03a.75.75 0 0 0-1.06-1.06l-3.5 3.5a.75.75 0 0 0 0 1.06l3.5 3.5a.75.75 0 1 0 1.06-1.06L9.56 8zm-5 0a.75.75 0 0 0-1.06-1.06l-3.5 3.5a.75.75 0 0 0 0 1.06l3.5 3.5a.75.75 0 0 0 1.06-1.06L4.56 8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronsRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.47 10.97a.75.75 0 1 0 1.06 1.06l3.5-3.5a.75.75 0 0 0 0-1.06l-3.5-3.5a.75.75 0 0 0-1.06 1.06L6.44 8zm5 0a.75.75 0 1 0 1.06 1.06l3.5-3.5a.75.75 0 0 0 0-1.06l-3.5-3.5a.75.75 0 0 0-1.06 1.06L11.44 8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronsUp(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.97 12.53a.75.75 0 1 0 1.06-1.06l-3.5-3.5a.75.75 0 0 0-1.06 0l-3.5 3.5a.75.75 0 1 0 1.06 1.06L8 9.56zm0-5a.75.75 0 1 0 1.06-1.06l-3.5-3.5a.75.75 0 0 0-1.06 0l-3.5 3.5a.75.75 0 0 0 1.06 1.06L8 4.56z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronsUpWide(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.097 12.383a.75.75 0 0 0 .806-1.266l-5.5-3.5a.75.75 0 0 0-.806 0l-5.5 3.5a.75.75 0 1 0 .806 1.266L8 9.139zm0-4.5a.75.75 0 0 0 .806-1.266l-5.5-3.5a.75.75 0 0 0-.806 0l-5.5 3.5a.75.75 0 0 0 .806 1.266L8 4.639z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circle(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 13.5a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleArrowDown(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M8.75 5a.75.75 0 0 0-1.5 0v4.19l-.72-.72a.75.75 0 0 0-1.06 1.06l2 2a.75.75 0 0 0 1.06 0l2-2a.75.75 0 1 0-1.06-1.06l-.72.72z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleArrowDownFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14m.75-10a.75.75 0 0 0-1.5 0v4.19l-.72-.72a.75.75 0 0 0-1.06 1.06l2 2a.75.75 0 0 0 1.06 0l2-2a.75.75 0 1 0-1.06-1.06l-.72.72z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleArrowLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0m-4 .75a.75.75 0 0 0 0-1.5H6.81l.72-.72a.75.75 0 0 0-1.06-1.06l-2 2a.75.75 0 0 0 0 1.06l2 2a.75.75 0 1 0 1.06-1.06l-.72-.72z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleArrowLeftFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14m3-6.25a.75.75 0 0 0 0-1.5H6.81l.72-.72a.75.75 0 0 0-1.06-1.06l-2 2a.75.75 0 0 0 0 1.06l2 2a.75.75 0 1 0 1.06-1.06l-.72-.72z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleArrowRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M5 7.25a.75.75 0 0 0 0 1.5h4.19l-.72.72a.75.75 0 0 0 1.06 1.06l2-2a.75.75 0 0 0 0-1.06l-2-2a.75.75 0 0 0-1.06 1.06l.72.72z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleArrowRightFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14M5 7.25a.75.75 0 0 0 0 1.5h4.19l-.72.72a.75.75 0 0 0 1.06 1.06l2-2a.75.75 0 0 0 0-1.06l-2-2a.75.75 0 0 0-1.06 1.06l.72.72z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleArrowUp(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0m-7.75 3a.75.75 0 0 0 1.5 0V6.81l.72.72a.75.75 0 1 0 1.06-1.06l-2-2a.75.75 0 0 0-1.06 0l-2 2a.75.75 0 1 0 1.06 1.06l.72-.72z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleArrowUpFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14m-.75-4a.75.75 0 0 0 1.5 0V6.81l.72.72a.75.75 0 1 0 1.06-1.06l-2-2a.75.75 0 0 0-1.06 0l-2 2a.75.75 0 1 0 1.06 1.06l.72-.72z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleCheck(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0m-3.9-1.55a.75.75 0 1 0-1.2-.9L7.419 8.858L6.03 7.47a.75.75 0 0 0-1.06 1.06l2 2a.75.75 0 0 0 1.13-.08z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleCheckFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14m3.1-8.55a.75.75 0 1 0-1.2-.9L7.419 8.858L6.03 7.47a.75.75 0 0 0-1.06 1.06l2 2a.75.75 0 0 0 1.13-.08z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevronDown(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 8a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0M1 8a7 7 0 1 1 14 0A7 7 0 0 1 1 8m5.03-1.28a.75.75 0 0 0-1.06 1.06l2.5 2.5a.75.75 0 0 0 1.06 0l2.5-2.5a.75.75 0 1 0-1.06-1.06L8 8.69z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevronDownFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1a7 7 0 1 0 0 14A7 7 0 0 0 8 1M6.03 6.72a.75.75 0 0 0-1.06 1.06l2.5 2.5a.75.75 0 0 0 1.06 0l2.5-2.5a.75.75 0 1 0-1.06-1.06L8 8.69z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevronLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 2.5a5.5 5.5 0 1 1 0 11a5.5 5.5 0 0 1 0-11M8 1a7 7 0 1 1 0 14A7 7 0 0 1 8 1m1.28 5.03a.75.75 0 0 0-1.06-1.06l-2.5 2.5a.75.75 0 0 0 0 1.06l2.5 2.5a.75.75 0 1 0 1.06-1.06L7.31 8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevronLeftFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 8A7 7 0 1 0 1 8a7 7 0 0 0 14 0M9.28 6.03a.75.75 0 0 0-1.06-1.06l-2.5 2.5a.75.75 0 0 0 0 1.06l2.5 2.5a.75.75 0 1 0 1.06-1.06L7.31 8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevronRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 13.5a5.5 5.5 0 1 1 0-11a5.5 5.5 0 0 1 0 11M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14M6.72 9.97a.75.75 0 0 0 1.06 1.06l2.5-2.5a.75.75 0 0 0 0-1.06l-2.5-2.5a.75.75 0 0 0-1.06 1.06L8.69 8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevronRightFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 8a7 7 0 1 0 14 0A7 7 0 0 0 1 8m5.72 1.97a.75.75 0 1 0 1.06 1.06l2.5-2.5a.75.75 0 0 0 0-1.06l-2.5-2.5a.75.75 0 0 0-1.06 1.06L8.69 8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevronUp(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M9.97 9.28a.75.75 0 1 0 1.06-1.06l-2.5-2.5a.75.75 0 0 0-1.06 0l-2.5 2.5a.75.75 0 0 0 1.06 1.06L8 7.31z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevronUpFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14m1.97-5.72a.75.75 0 1 0 1.06-1.06l-2.5-2.5a.75.75 0 0 0-1.06 0l-2.5 2.5a.75.75 0 0 0 1.06 1.06L8 7.31z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevronsDown(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 2.5a5.5 5.5 0 1 0 0 11a5.5 5.5 0 0 0 0-11M8 1a7 7 0 1 0 0 14A7 7 0 0 0 8 1M5.72 4.97a.75.75 0 0 1 1.06 0L8 6.19l1.22-1.22a.75.75 0 1 1 1.06 1.06L8.53 7.78a.75.75 0 0 1-1.06 0L5.72 6.03a.75.75 0 0 1 0-1.06m1.06 3.75a.75.75 0 0 0-1.06 1.06l1.75 1.75a.75.75 0 0 0 1.06 0l1.75-1.75a.75.75 0 1 0-1.06-1.06L8 9.94z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevronsLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 0-11 0a5.5 5.5 0 0 0 11 0M15 8A7 7 0 1 0 1 8a7 7 0 0 0 14 0m-3.97-2.28a.75.75 0 0 1 0 1.06L9.81 8l1.22 1.22a.75.75 0 1 1-1.06 1.06L8.22 8.53a.75.75 0 0 1 0-1.06l1.75-1.75a.75.75 0 0 1 1.06 0M7.28 6.78a.75.75 0 0 0-1.06-1.06L4.47 7.47a.75.75 0 0 0 0 1.06l1.75 1.75a.75.75 0 1 0 1.06-1.06L6.06 8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevronsRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 8a5.5 5.5 0 1 0 11 0a5.5 5.5 0 0 0-11 0M1 8a7 7 0 1 0 14 0A7 7 0 0 0 1 8m3.97 2.28a.75.75 0 0 1 0-1.06L6.19 8L4.97 6.78a.75.75 0 0 1 1.06-1.06l1.75 1.75a.75.75 0 0 1 0 1.06l-1.75 1.75a.75.75 0 0 1-1.06 0m3.75-1.06a.75.75 0 1 0 1.06 1.06l1.75-1.75a.75.75 0 0 0 0-1.06L9.78 5.72a.75.75 0 0 0-1.06 1.06L9.94 8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevronsUp(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 13.5a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14m2.28-3.97a.75.75 0 0 1-1.06 0L8 9.81l-1.22 1.22a.75.75 0 1 1-1.06-1.06l1.75-1.75a.75.75 0 0 1 1.06 0l1.75 1.75a.75.75 0 0 1 0 1.06M9.22 7.28a.75.75 0 1 0 1.06-1.06L8.53 4.47a.75.75 0 0 0-1.06 0L5.72 6.22a.75.75 0 0 0 1.06 1.06L8 6.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleDashed(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.906 1.085a7.047 7.047 0 0 1 2.188 0a.75.75 0 0 1-.232 1.482a5.546 5.546 0 0 0-1.724 0a.75.75 0 0 1-.232-1.482M4.933 2.502a.75.75 0 0 1-.166 1.048c-.466.34-.878.75-1.217 1.217a.75.75 0 0 1-1.213-.882a7.036 7.036 0 0 1 1.548-1.548a.75.75 0 0 1 1.048.165m6.135 0a.75.75 0 0 1 1.047-.165a7.037 7.037 0 0 1 1.548 1.548a.75.75 0 0 1-1.213.882a5.533 5.533 0 0 0-1.217-1.217a.75.75 0 0 1-.165-1.048M1.943 6.28a.75.75 0 0 1 .624.857a5.546 5.546 0 0 0 0 1.724a.75.75 0 0 1-1.482.232a7.047 7.047 0 0 1 0-2.188a.75.75 0 0 1 .858-.625m12.114 0a.75.75 0 0 1 .858.625a7.048 7.048 0 0 1 0 2.188a.75.75 0 1 1-1.482-.232a5.54 5.54 0 0 0 0-1.724a.75.75 0 0 1 .624-.857M2.502 11.068a.75.75 0 0 1 1.048.165c.34.466.75.878 1.217 1.217a.75.75 0 0 1-.882 1.213a7.037 7.037 0 0 1-1.548-1.548a.75.75 0 0 1 .165-1.047m10.996 0a.75.75 0 0 1 .165 1.047a7.037 7.037 0 0 1-1.548 1.548a.75.75 0 0 1-.883-1.213a5.53 5.53 0 0 0 1.218-1.217a.75.75 0 0 1 1.048-.165m-7.217 2.99a.75.75 0 0 1 .857-.625a5.54 5.54 0 0 0 1.724 0a.75.75 0 0 1 .232 1.482a7.048 7.048 0 0 1-2.188 0a.75.75 0 0 1-.625-.857" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleDollar(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M8.75 4.25a.75.75 0 0 0-1.5 0v.339a2.477 2.477 0 0 0-1.007.47a1.95 1.95 0 0 0-.74 1.546c0 .764.474 1.265.94 1.559c.456.287 1.007.448 1.448.547c.462.102.843.191 1.118.341c.228.125.275.224.275.376c0 .102-.04.217-.248.341c-.224.135-.577.229-.982.229c-.344 0-.683-.114-.953-.29c-.281-.184-.42-.388-.457-.506a.75.75 0 1 0-1.43.452c.171.543.591 1 1.068 1.31c.284.185.612.335.968.429v.357a.75.75 0 0 0 1.5 0v-.313c.375-.067.74-.19 1.058-.382c.53-.319.976-.864.976-1.627c0-.864-.51-1.394-1.055-1.692c-.478-.26-1.056-.389-1.46-.478l-.053-.012c-.386-.086-.736-.202-.973-.352c-.227-.142-.24-.236-.24-.29a.45.45 0 0 1 .18-.375c.134-.108.403-.227.87-.227c.47 0 .742.11.9.218a.832.832 0 0 1 .316.41a.75.75 0 0 0 1.407-.52a2.331 2.331 0 0 0-.878-1.13a2.713 2.713 0 0 0-1.048-.417z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleExclamation(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 13.5a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14m1-4.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0M8.75 5a.75.75 0 0 0-1.5 0v2.5a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleExclamationFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0m-6 2.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0M8.75 5a.75.75 0 0 0-1.5 0v2.5a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleInfo(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 13.5a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14m1-9.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-.25 3a.75.75 0 0 0-1.5 0V11a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleInfoFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14m1-9.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0M8 7.75a.75.75 0 0 1 .75.75V11a.75.75 0 0 1-1.5 0V8.5A.75.75 0 0 1 8 7.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleLetterC(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M6.458 8c0-.8.2-1.289.445-1.569C7.136 6.165 7.487 6 8 6c.368 0 .648.086.857.222c.205.134.384.345.51.672a.75.75 0 0 0 1.4-.538c-.226-.59-.593-1.066-1.09-1.39C9.181 4.643 8.604 4.5 8 4.5c-.862 0-1.657.293-2.226.944c-.557.637-.816 1.523-.816 2.556c0 1.033.259 1.92.816 2.556c.569.65 1.364.944 2.226.944c.605 0 1.182-.143 1.676-.466c.498-.324.865-.8 1.091-1.39a.75.75 0 0 0-1.4-.538c-.126.327-.305.538-.51.672c-.209.136-.49.222-.857.222c-.513 0-.864-.165-1.097-.431c-.245-.28-.445-.77-.445-1.569" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleLetterF(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M6.75 5a.75.75 0 0 0-.75.75v5a.75.75 0 0 0 1.5 0v-1.5h1.75a.75.75 0 0 0 0-1.5H7.5V6.5h2.25a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleLetterH(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M7 5.5a.75.75 0 0 0-1.5 0v5a.75.75 0 0 0 1.5 0V8.75h2v1.75a.75.75 0 0 0 1.5 0v-5a.75.75 0 0 0-1.5 0v1.75H7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleLetterP(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M6.75 5a.75.75 0 0 0-.75.75v5a.75.75 0 0 0 1.5 0V9.5h1a2.25 2.25 0 0 0 0-4.5zm2.5 2.25A.75.75 0 0 1 8.5 8h-1V6.5h1a.75.75 0 0 1 .75.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleLetterR(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M5.75 5.5a.75.75 0 0 1 .75-.75h1.75a2.25 2.25 0 0 1 1.38 4.028l.771 1.35a.75.75 0 0 1-1.302.744L8.172 9.25H7.25v1.25a.75.75 0 0 1-1.5 0zm1.5 2.25h1a.75.75 0 0 0 0-1.5h-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleLetterS(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M8.054 4.502c-.708 0-1.34.18-1.811.558a1.95 1.95 0 0 0-.74 1.545c0 .764.474 1.265.94 1.559c.457.287 1.007.448 1.448.547c.462.102.844.191 1.119.341c.227.125.274.224.274.376c0 .102-.04.217-.248.341c-.224.135-.577.229-.982.229c-.344 0-.682-.114-.952-.29c-.282-.184-.42-.388-.457-.506a.75.75 0 1 0-1.43.452c.17.543.59 1 1.067 1.31a3.273 3.273 0 0 0 1.772.534c.602 0 1.24-.134 1.754-.443c.53-.319.976-.864.976-1.627c0-.864-.51-1.394-1.055-1.692c-.478-.26-1.056-.389-1.46-.478l-.053-.012c-.386-.086-.736-.202-.973-.352c-.227-.142-.24-.236-.24-.29a.45.45 0 0 1 .18-.375c.134-.108.403-.227.87-.227c.47 0 .742.11.9.218a.832.832 0 0 1 .316.41a.75.75 0 0 0 1.407-.52A2.33 2.33 0 0 0 9.8 4.98c-.455-.31-1.037-.478-1.745-.478Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleLetterT(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M6 5a.75.75 0 0 0 0 1.5h1.25v4.25a.75.75 0 0 0 1.5 0V6.5H10A.75.75 0 0 0 10 5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleLetterW(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M5.712 5.763a.75.75 0 1 0-1.424.474l1.5 4.5a.75.75 0 0 0 1.424 0L8 8.372l.788 2.365a.75.75 0 0 0 1.423 0l1.5-4.5a.75.75 0 1 0-1.422-.474L9.5 8.128l-.788-2.365a.75.75 0 0 0-1.424 0L6.5 8.128z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleLink(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 13.5a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14M6.221 7.298a.707.707 0 0 0-1 0A2.461 2.461 0 1 0 8.7 10.78a.707.707 0 0 0 0-1l-.06-.06a.707.707 0 0 0-1 0a.961.961 0 0 1-1.36-1.36a.707.707 0 0 0 0-1l-.06-.06Zm3.497 1.344a.707.707 0 0 1 0-1.001a.961.961 0 0 0-1.359-1.36a.707.707 0 0 1-1 0l-.06-.06a.707.707 0 0 1 0-1a2.461 2.461 0 1 1 3.48 3.48a.707.707 0 0 1-1 0zm-.781-.518a.75.75 0 0 0-1.061-1.06l-.813.812a.75.75 0 0 0 1.061 1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleMinus(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0m-9.5-.75a.75.75 0 0 0 0 1.5h5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleMinusFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14M5.5 7.25a.75.75 0 0 0 0 1.5h5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleNumberEight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 13.5a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14m1.443-5.436A.692.692 0 0 0 9.5 9.25a.692.692 0 0 0-.057-.314a.413.413 0 0 0-.14-.153C9.124 8.655 8.723 8.5 8 8.5s-1.124.155-1.303.283a.413.413 0 0 0-.14.153a.692.692 0 0 0-.057.314c0 .172.033.265.057.314a.413.413 0 0 0 .14.153c.179.128.58.283 1.303.283s1.124-.155 1.303-.283a.413.413 0 0 0 .14-.153m-3.647-1.98C5.29 7.957 5 8.513 5 9.25c0 1.5 1.2 2.25 3 2.25s3-.75 3-2.25c0-.736-.289-1.292-.796-1.667c.192-.291.296-.652.296-1.083c0-1.333-1-2-2.5-2s-2.5.667-2.5 2c0 .43.104.792.296 1.083ZM8 7c.556 0 .817-.127.903-.193a.173.173 0 0 0 .06-.069A.567.567 0 0 0 9 6.5a.567.567 0 0 0-.038-.238a.173.173 0 0 0-.059-.069C8.817 6.127 8.556 6 8 6s-.817.127-.903.193a.173.173 0 0 0-.06.069A.567.567 0 0 0 7 6.5c0 .144.026.214.038.238c.01.022.024.042.059.069c.086.066.347.193.903.193" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleNumberNine(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8 13.5a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14"/><path d="M9.491 8.81a5.344 5.344 0 0 1-1.527.289c-.623.025-1.319-.059-1.9-.37c-.62-.334-1.085-.921-1.144-1.767C4.822 5.56 6.151 4.621 7.756 4.51c.883-.062 1.704.17 2.315.79c.594.601.892 1.47.963 2.49c.072 1.03-.124 1.932-.635 2.606c-.523.688-1.295 1.037-2.155 1.097c-1.167.08-1.81-.183-2.37-.416a.75.75 0 1 1 .575-1.385l.012.005c.467.194.857.356 1.678.3c.512-.037.85-.225 1.065-.508c.122-.16.226-.382.287-.678Zm-.334-1.48a1.65 1.65 0 0 0 .273-.144c-.102-.397-.258-.663-.427-.834c-.237-.24-.592-.386-1.142-.347c-1.262.088-1.456.689-1.444.853a.611.611 0 0 0 .357.549c.258.138.655.212 1.128.193a3.71 3.71 0 0 0 1.255-.27"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleNumberOne(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M9 5.5a.75.75 0 0 0-1.23-.576l-1.5 1.25a.75.75 0 1 0 .96 1.152l.27-.225V10.5a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleNumberSeven(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M6.5 5a.75.75 0 1 0 0 1.5h2.174l-1.378 3.74a.75.75 0 0 0 1.408.52l1.75-4.75A.75.75 0 0 0 9.75 5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleNumberSix(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M6.569 7.088a5.344 5.344 0 0 1 1.544-.181c.623.017 1.311.15 1.869.502c.596.375 1.018.993 1.018 1.841c0 1.405-1.39 2.25-3 2.25c-.885 0-1.688-.289-2.254-.95c-.55-.641-.788-1.529-.788-2.55c0-1.033.259-1.92.816-2.556C6.343 4.794 7.138 4.5 8 4.5c1.17 0 1.793.308 2.335.58a.75.75 0 1 1-.67 1.34l-.012-.005C9.201 6.19 8.823 6 8 6c-.513 0-.864.165-1.097.431a1.665 1.665 0 0 0-.334.657M8.07 8.406a3.71 3.71 0 0 0-1.271.18a1.648 1.648 0 0 0-.282.126c.074.403.21.68.367.862c.22.257.564.426 1.116.426c1.265 0 1.5-.586 1.5-.75c0-.277-.109-.44-.318-.573c-.248-.156-.639-.258-1.112-.271"/><path d="M8 13.5a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleNumberZero(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M9.5 8c0 .923-.236 1.396-.44 1.622C8.875 9.828 8.565 10 8 10c-.565 0-.875-.172-1.06-.378C6.736 9.396 6.5 8.923 6.5 8c0-.923.236-1.395.44-1.621C7.125 6.173 7.435 6 8 6c.565 0 .875.173 1.06.379c.204.226.44.698.44 1.621M11 8c0 2.333-1.2 3.5-3 3.5S5 10.333 5 8s1.2-3.5 3-3.5s3 1.167 3 3.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CirclePause(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 13.5a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14m1.75-9.75a1 1 0 0 1 1 1v3.5a1 1 0 1 1-2 0v-3.5a1 1 0 0 1 1-1m-2.5 1a1 1 0 0 0-2 0v3.5a1 1 0 1 0 2 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CirclePlay(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0m-7.75 3.031L11 8.866a1 1 0 0 0 0-1.732L7.25 4.969a1 1 0 0 0-1.5.866v4.33a1 1 0 0 0 1.5.866" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CirclePlus(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M8.75 5.5a.75.75 0 0 0-1.5 0v1.75H5.5a.75.75 0 0 0 0 1.5h1.75v1.75a.75.75 0 0 0 1.5 0V8.75h1.75a.75.75 0 0 0 0-1.5H8.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CirclePlusFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14m.75-9.5a.75.75 0 0 0-1.5 0v1.75H5.5a.75.75 0 0 0 0 1.5h1.75v1.75a.75.75 0 0 0 1.5 0V8.75h1.75a.75.75 0 0 0 0-1.5H8.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleQuestion(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 13.5a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14M6.44 4.54c.43-.354.994-.565 1.56-.565c1.217 0 2.34.82 2.34 2.14c0 .377-.079.745-.298 1.1c-.208.339-.513.614-.875.867c-.217.153-.326.257-.379.328c-.038.052-.038.07-.038.089a.75.75 0 0 1-1.5 0c0-.794.544-1.286 1.057-1.645c.28-.196.4-.332.458-.426a.543.543 0 0 0 .074-.312c0-.3-.243-.641-.839-.641a.997.997 0 0 0-.608.223c-.167.138-.231.287-.231.418a.75.75 0 1 1-1.5 0c0-.674.345-1.22.78-1.577ZM8 12a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleQuestionFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14M6.44 4.54c.43-.354.994-.565 1.56-.565c1.217 0 2.34.82 2.34 2.14c0 .377-.079.745-.298 1.1c-.208.339-.513.614-.875.867c-.217.153-.326.257-.379.328c-.038.052-.038.07-.038.089a.75.75 0 0 1-1.5 0c0-.794.544-1.286 1.056-1.645c.28-.196.402-.332.46-.425a.543.543 0 0 0 .073-.313c0-.3-.243-.641-.839-.641a.997.997 0 0 0-.608.224c-.167.137-.231.286-.231.417a.75.75 0 0 1-1.5 0c0-.673.345-1.22.78-1.577ZM9 11a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleRuble(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M6.75 4a.75.75 0 0 0-.75.75v2.5h-.125a.75.75 0 0 0 0 1.5H6v.5h-.125a.75.75 0 0 0 0 1.5H6v.5a.75.75 0 0 0 1.5 0v-.5H9a.75.75 0 0 0 0-1.5H7.5v-.5h1.125a2.375 2.375 0 1 0 0-4.75zm1.875 3.25H7.5V5.5h1.125a.875.875 0 1 1 0 1.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleStop(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M5.25 6.25a1 1 0 0 1 1-1h3.5a1 1 0 0 1 1 1v3.5a1 1 0 0 1-1 1h-3.5a1 1 0 0 1-1-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleXmark(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M6.53 5.47a.75.75 0 0 0-1.06 1.06L6.94 8L5.47 9.47a.75.75 0 1 0 1.06 1.06L8 9.06l1.47 1.47a.75.75 0 1 0 1.06-1.06L9.06 8l1.47-1.47a.75.75 0 1 0-1.06-1.06L8 6.94z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleXmarkFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14M6.53 5.47a.75.75 0 0 0-1.06 1.06L6.94 8L5.47 9.47a.75.75 0 1 0 1.06 1.06L8 9.06l1.47 1.47a.75.75 0 1 0 1.06-1.06L9.06 8l1.47-1.47a.75.75 0 1 0-1.06-1.06L8 6.94z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CirclesConcentric(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 8a5.5 5.5 0 0 1 8.14-4.827a.75.75 0 1 0 .72-1.315a7 7 0 0 0-9.502 9.502a.75.75 0 0 0 1.315-.72A5.472 5.472 0 0 1 2.5 8m11.642-3.36a.75.75 0 1 0-1.315.72a5.5 5.5 0 0 1-7.466 7.466a.75.75 0 1 0-.722 1.316a7 7 0 0 0 9.502-9.502ZM9.5 8a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0M11 8a3 3 0 1 1-6 0a3 3 0 0 1 6 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CirclesFiveRandom(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.069 5.55A2.497 2.497 0 0 0 5 3.5a2.5 2.5 0 1 0-3.931 2.05M2.5 4.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2m11 1.5A2.497 2.497 0 0 1 11 3.5A2.5 2.5 0 1 1 13.5 6m1-2.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0M6.069 8.55A2.497 2.497 0 0 0 10 6.5a2.5 2.5 0 1 0-3.931 2.05M7.5 7.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2M12 14a2.497 2.497 0 0 1-2.5-2.5A2.5 2.5 0 1 1 12 14m1-2.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0M4.5 16A2.497 2.497 0 0 1 2 13.5A2.5 2.5 0 1 1 4.5 16m1-2.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CirclesFourDiamond(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiCircles4Diamond0)"><path fill="currentColor" fill-rule="evenodd" d="M10.502 4.392a2.75 2.75 0 1 0-5.004-2.285a2.75 2.75 0 0 0 5.004 2.285m-4.75 2.466a2.76 2.76 0 0 0-1.36-1.36a2.75 2.75 0 1 0 1.36 1.36m1.106 3.39a2.76 2.76 0 0 0-1.36 1.36a2.75 2.75 0 1 0 1.36-1.36m3.39-1.106a2.753 2.753 0 0 0 1.36 1.36a2.75 2.75 0 1 0-1.36-1.36M8 2a1.25 1.25 0 1 0 0 2.5A1.25 1.25 0 0 0 8 2m6 6a1.25 1.25 0 1 0-2.5 0A1.25 1.25 0 0 0 14 8m-6 3.5A1.25 1.25 0 1 0 8 14a1.25 1.25 0 0 0 0-2.5M4.5 8A1.25 1.25 0 1 0 2 8a1.25 1.25 0 0 0 2.5 0" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiCircles4Diamond0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CirclesFourSquare(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.047 7.125A3 3 0 1 0 4.5 1.5a3 3 0 0 0-1.453 5.625M4.5 6a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m5.547 1.125A3 3 0 1 0 11.5 1.5a3 3 0 0 0-1.453 5.625M11.5 6a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m-7 8.5a3 3 0 1 1 0-6a3 3 0 0 1 0 6m1.5-3a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0m4.047 2.625A3 3 0 1 0 11.5 8.5a3 3 0 0 0-1.453 5.625M11.5 13a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CirclesIntersection(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.576 11.854a4 4 0 1 1 0-7.707A5.482 5.482 0 0 0 5 8c0 1.5.601 2.861 1.576 3.854M8 12.9a5.5 5.5 0 1 1 0-9.8a5.5 5.5 0 1 1 0 9.8m1.424-8.754A4.003 4.003 0 0 1 14.5 8a4 4 0 0 1-5.076 3.854A5.483 5.483 0 0 0 11 8c0-1.5-.601-2.861-1.576-3.854M8 4.877C8.914 5.61 9.5 6.737 9.5 8c0 1.263-.586 2.39-1.5 3.123A3.993 3.993 0 0 1 6.5 8c0-1.263.586-2.39 1.5-3.123" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CirclesThreePlus(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.047 7.125A3 3 0 1 0 4.5 1.5a3 3 0 0 0-1.453 5.625M4.5 6a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m-1.453 8.125A3 3 0 1 0 4.5 8.5a3 3 0 0 0-1.453 5.625M4.5 13a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m7 1.5a3 3 0 1 1 0-6a3 3 0 1 1 0 6m1.5-3a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0m-.75-9.25a.75.75 0 0 0-1.5 0v1.501h-1.5a.75.75 0 1 0 0 1.5h1.501v1.502a.75.75 0 0 0 1.5 0V5.25h1.501a.75.75 0 0 0 0-1.5h-1.5v-1.5Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M8.75 4.5a.75.75 0 0 0-1.5 0V8a.75.75 0 0 0 .3.6l2 1.5a.75.75 0 1 0 .9-1.2l-1.7-1.275z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockArrowRotateLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 8a6.5 6.5 0 1 1 7.348 6.445a.75.75 0 1 1-.194-1.487A5.001 5.001 0 1 0 4.5 11.57v-1.32a.75.75 0 0 1 1.5 0v3a.75.75 0 0 1-.75.75h-3a.75.75 0 0 1 0-1.5h1.06A6.48 6.48 0 0 1 1.5 8M8 4.25a.75.75 0 0 1 .75.75v2.625l1.033.775a.75.75 0 1 1-.9 1.2l-1.333-1a.75.75 0 0 1-.3-.6V5A.75.75 0 0 1 8 4.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14m.75-10.5a.75.75 0 0 0-1.5 0V8a.75.75 0 0 0 .3.6l2 1.5a.75.75 0 1 0 .9-1.2l-1.7-1.275z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 6.25a3.25 3.25 0 0 1 6.051-1.65a4.497 4.497 0 0 0-2.35 1.34A.75.75 0 0 0 9.3 6.96a2.99 2.99 0 0 1 2.3-.958A3 3 0 0 1 11.5 12H3.75a2.25 2.25 0 0 1-.002-4.5h.03a.75.75 0 0 0 .747-.843A3.289 3.289 0 0 1 4.5 6.25M7.75 1.5a4.75 4.75 0 0 0-4.747 4.574A3.751 3.751 0 0 0 3.75 13.5h7.75a4.5 4.5 0 0 0 .687-8.948A4.751 4.751 0 0 0 7.75 1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudArrowUpIn(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiCloudArrowUpIn0)"><path fill="currentColor" fill-rule="evenodd" d="M4.5 5.25a3.25 3.25 0 0 1 6.398-.811a.75.75 0 0 0 .702.563A3 3 0 0 1 11.5 11h-.75a.75.75 0 0 0 0 1.5h.75a4.5 4.5 0 0 0 .687-8.948a4.751 4.751 0 0 0-9.184 1.522A3.751 3.751 0 0 0 3.75 12.5h1.5a.75.75 0 0 0 0-1.5H3.751a2.25 2.25 0 0 1-.003-4.5h.03a.75.75 0 0 0 .747-.843A3.289 3.289 0 0 1 4.5 5.25m4.25 3.31l.72.72a.75.75 0 1 0 1.06-1.06l-2-2a.75.75 0 0 0-1.06 0l-2 2a.75.75 0 0 0 1.06 1.06l.72-.72v6.69a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiCloudArrowUpIn0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudCheck(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 6.25a3.25 3.25 0 0 1 6.398-.811a.75.75 0 0 0 .702.563A3 3 0 0 1 11.5 12H3.75a2.25 2.25 0 0 1-.002-4.5h.03a.75.75 0 0 0 .747-.843A3.289 3.289 0 0 1 4.5 6.25M7.75 1.5a4.75 4.75 0 0 0-4.747 4.574A3.751 3.751 0 0 0 3.75 13.5h7.75a4.5 4.5 0 0 0 .687-8.948A4.751 4.751 0 0 0 7.75 1.5m2.85 5.75a.75.75 0 1 0-1.2-.9L7.369 9.058L6.28 7.97a.75.75 0 1 0-1.06 1.06l1.7 1.7a.75.75 0 0 0 1.13-.08z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudGear(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 5.25a3.25 3.25 0 0 1 6.398-.811a.75.75 0 0 0 .702.563a2.986 2.986 0 0 1 1.94.798c.591.549.96 1.33.96 2.2A.75.75 0 0 0 16 8a4.49 4.49 0 0 0-1.44-3.3a4.489 4.489 0 0 0-2.373-1.148a4.751 4.751 0 0 0-9.184 1.522A3.751 3.751 0 0 0 3.75 12.5h.5a.75.75 0 0 0 0-1.5h-.5a2.25 2.25 0 0 1-.002-4.5h.03a.75.75 0 0 0 .747-.843A3.289 3.289 0 0 1 4.5 5.25m8.509 5.939a2.223 2.223 0 0 1-1.66-2.138l-.68-.256a2.38 2.38 0 0 1-2.748.816l-.406.569c.57.788.57 1.87 0 2.657l.406.569a2.38 2.38 0 0 1 2.747.816l.68-.255a2.223 2.223 0 0 1 1.66-2.139v-.639Zm-.113-2.396a.715.715 0 0 0-.418-.921l-1.798-.674a.865.865 0 0 0-1.114.506a.87.87 0 0 1-1.32.395a.884.884 0 0 0-1.23.2L5.94 9.805a.771.771 0 0 0 .18 1.076c.43.307.43.948 0 1.255a.771.771 0 0 0-.18 1.076l1.076 1.506a.884.884 0 0 0 1.23.2a.87.87 0 0 1 1.32.395a.865.865 0 0 0 1.114.506l1.798-.674a.715.715 0 0 0 .418-.92a.715.715 0 0 1 .67-.966h.134a.808.808 0 0 0 .809-.809v-1.883a.808.808 0 0 0-.809-.808h-.134a.715.715 0 0 1-.67-.966m-1.833 2.709a1.052 1.052 0 1 1-2.103 0a1.052 1.052 0 0 1 2.103 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudNutHex(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 5.25a3.25 3.25 0 0 1 6.398-.811a.75.75 0 0 0 .702.563a2.986 2.986 0 0 1 1.94.798c.591.549.96 1.33.96 2.2A.75.75 0 0 0 16 8a4.49 4.49 0 0 0-1.44-3.3a4.489 4.489 0 0 0-2.373-1.148a4.751 4.751 0 0 0-9.184 1.522A3.751 3.751 0 0 0 3.75 12.5h.5a.75.75 0 0 0 0-1.5h-.5a2.25 2.25 0 0 1-.002-4.5h.03a.75.75 0 0 0 .747-.843A3.289 3.289 0 0 1 4.5 5.25m7.123 8.25l1.406-2.5l-1.406-2.5H8.877L7.471 11l1.406 2.5zm2.713-1.765a1.5 1.5 0 0 0 0-1.47l-1.406-2.5A1.5 1.5 0 0 0 11.623 7H8.877a1.5 1.5 0 0 0-1.307.765l-1.406 2.5a1.5 1.5 0 0 0 0 1.47l1.406 2.5A1.5 1.5 0 0 0 8.877 15h2.746a1.5 1.5 0 0 0 1.307-.765zM11.25 11a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clouds(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiClouds0)"><path fill="currentColor" fill-rule="evenodd" d="M4.977 4.563a3.208 3.208 0 0 0-.196.053l.062.006A4.062 4.062 0 0 0 2.49 8.06a3.234 3.234 0 0 0 .743 6.38h6.414a3.854 3.854 0 0 0 3.851-3.744a3.855 3.855 0 0 0-.69-7.406a4.061 4.061 0 0 0-7.817 1.267a5.299 5.299 0 0 0-.014.006m8.256 4.611a2.353 2.353 0 0 0-1.008-4.44a.75.75 0 0 1-.702-.563a2.562 2.562 0 0 0-4.979.079a4.061 4.061 0 0 1 3.765 2.54a3.86 3.86 0 0 1 2.924 2.384m-9.25-.864a2.56 2.56 0 0 1 5.04-.639a.75.75 0 0 0 .702.563a2.353 2.353 0 0 1-.078 4.706H3.233a1.733 1.733 0 0 1-.002-3.466h.024a.75.75 0 0 0 .748-.843a2.593 2.593 0 0 1-.02-.32Zm5.72 4.632H9.7zm-.11 1.496H9.59z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiClouds0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Code(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.218 3.216a.75.75 0 1 0-1.436-.431l-3 10a.75.75 0 0 0 1.436.43l3-10ZM4.53 4.97a.75.75 0 0 1 0 1.06L2.56 8l1.97 1.97a.75.75 0 0 1-1.06 1.06l-2.5-2.5a.75.75 0 0 1 0-1.06l2.5-2.5a.75.75 0 0 1 1.06 0m6.94 6.06a.75.75 0 0 1 0-1.06L13.44 8l-1.97-1.97a.75.75 0 0 1 1.06-1.06l2.5 2.5a.75.75 0 0 1 0 1.06l-2.5 2.5a.75.75 0 0 1-1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeCommit(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.25 1.75v2.32a4.001 4.001 0 0 0 0 7.86v2.32a.75.75 0 0 0 1.5 0v-2.32a4.001 4.001 0 0 0 0-7.86V1.75a.75.75 0 0 0-1.5 0M8 10.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeCommitHorizontal(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.25 7.25h-2.32a4.001 4.001 0 0 0-7.86 0H1.75a.75.75 0 0 0 0 1.5h2.32a4.001 4.001 0 0 0 7.86 0h2.32a.75.75 0 0 0 0-1.5M5.5 8a2.5 2.5 0 1 0 5 0a2.5 2.5 0 0 0-5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeCommits(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiCodeCommits0)"><path fill="currentColor" fill-rule="evenodd" d="M8 3.25h-.008a1.25 1.25 0 1 0 .016 0zm.75-2.5v1.104a2.751 2.751 0 0 1 0 5.292v1.708a2.751 2.751 0 0 1 0 5.293v1.103a.75.75 0 0 1-1.5 0v-1.104a2.751 2.751 0 0 1 0-5.292V7.147a2.751 2.751 0 0 1 0-5.293V.75a.75.75 0 0 1 1.5 0m-.745 12h-.01a1.25 1.25 0 1 1 .01 0" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiCodeCommits0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeCompare(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.22 14.03s-.001 0 0 0a.75.75 0 0 0 1.06-1.06l-.47-.47H10a3.016 3.016 0 0 0 1.507-.405A2.999 2.999 0 0 0 13 9.5V6.896h.003a2.735 2.735 0 0 0 .785-.366a2.75 2.75 0 1 0-2.288.366V9.5A1.5 1.5 0 0 1 10 11h-.19l.47-.47s0 .001 0 0a.75.75 0 0 0-1.06-1.06l-.47.47l-1.28 1.28a.75.75 0 0 0 0 1.06zM5.72 1.97a.75.75 0 0 1 1.06 0l.47.47l1.28 1.28a.748.748 0 0 1 0 1.06L6.78 6.53c.001 0 0 0 0 0a.751.751 0 0 1-1.06-1.06L6.19 5H6a1.5 1.5 0 0 0-1.5 1.5v2.604a2.757 2.757 0 0 1 2 2.646a2.738 2.738 0 0 1-1.212 2.28a2.737 2.737 0 0 1-1.538.47A2.747 2.747 0 0 1 1 11.75a2.751 2.751 0 0 1 2-2.646V6.5a2.999 2.999 0 0 1 3-3h.19l-.47-.47a.75.75 0 0 1 0-1.06m-.908 9.121A1.246 1.246 0 0 1 5 11.75a1.25 1.25 0 1 1-.188-.659M11 4.25a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeFork(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.504 5.897a2.751 2.751 0 1 1 1.503-.002A1.5 1.5 0 0 0 6.5 7.25h3a1.5 1.5 0 0 0 1.493-1.355a2.751 2.751 0 1 1 1.503.002A3 3 0 0 1 9.5 8.75h-.75v1.354a2.751 2.751 0 1 1-1.5 0V8.75H6.5a3 3 0 0 1-2.996-2.853M3 3.25a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0m3.75 9.5a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0m3.75-9.5a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeMerge(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.312 11.091a1.25 1.25 0 1 0-2.123 1.316a1.25 1.25 0 0 0 2.123-1.316M3.188 4.909a1.25 1.25 0 1 1 2.125-1.318a1.25 1.25 0 0 1-2.125 1.318M5 6.896v2.208a2.751 2.751 0 1 1-1.5 0V6.896A2.751 2.751 0 1 1 6.896 3.5H9.5a3 3 0 0 1 3 3v2.604a2.751 2.751 0 1 1-1.5 0V6.5A1.5 1.5 0 0 0 9.5 5H6.896A2.756 2.756 0 0 1 5 6.896m7.812 4.195a1.25 1.25 0 1 1-2.125 1.318a1.25 1.25 0 0 1 2.125-1.318" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodePullRequest(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.788 3.121a1.25 1.25 0 1 1-1.076 2.257a1.25 1.25 0 0 1 1.076-2.257M3.5 6.896a2.751 2.751 0 1 1 3.347-3.554L8.219 1.97A.75.75 0 0 1 9.28 3.03l-.47.47h.69a3.04 3.04 0 0 1 1.139.224A3 3 0 0 1 12.5 6.5v2.604l.068.02a2.751 2.751 0 0 1-.818 5.376A2.75 2.75 0 0 1 11 9.104V6.5A1.501 1.501 0 0 0 9.5 5h-.69l.47.47a.75.75 0 0 1-1.06 1.06L6.847 5.157a2.737 2.737 0 0 1-.88 1.242c-.282.225-.61.397-.967.497v2.208l.068.02a2.751 2.751 0 1 1-1.568-.02zm-.312 4.195a1.246 1.246 0 0 0-.182.787a1.25 1.25 0 1 0 .182-.787m7.5 0a1.247 1.247 0 0 0-.181.787a1.25 1.25 0 1 0 .18-.787Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodePullRequestArrowLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.967 6.399a2.765 2.765 0 0 0 .88-1.242L8.22 6.53a.75.75 0 1 0 1.06-1.06L8.81 5h.69A1.501 1.501 0 0 1 11 6.5v2.75a.747.747 0 0 0 .75.75a.748.748 0 0 0 .75-.75V6.5a3 3 0 0 0-3-3h-.69l.47-.47a.75.75 0 0 0-1.06-1.06L6.847 3.343A2.755 2.755 0 0 0 4.25 1.5a2.75 2.75 0 0 0-.753 5.396H3.5V7.5a.75.75 0 0 0 1.5 0v-.603c.357-.101.685-.273.967-.498m-.75 3.133a.75.75 0 0 0-.747.188h-.001l-1.75 1.75a.75.75 0 0 0 0 1.06l1.75 1.75a.75.75 0 0 0 1.061-1.06l-.47-.47h7.69a.75.75 0 0 0 0-1.5H5.06l.47-.469a.75.75 0 0 0-.313-1.249m-.43-6.41a1.25 1.25 0 1 1-1.075 2.256a1.25 1.25 0 0 1 1.076-2.257Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodePullRequestArrowRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.53 5.788A2.765 2.765 0 0 1 5 6.896V9.25a.75.75 0 0 1-1.5 0V6.896h-.003a2.736 2.736 0 0 1-.785-.366a2.75 2.75 0 1 1 4.134-3.188L8.22 1.97a.75.75 0 0 1 1.06 1.06l-.47.47h.69a3.04 3.04 0 0 1 1.139.224A3 3 0 0 1 12.5 6.5v.75a.75.75 0 0 1-1.5 0V6.5A1.501 1.501 0 0 0 9.5 5h-.69l.47.47a.75.75 0 0 1-1.06 1.06L6.847 5.157v.001a2.738 2.738 0 0 1-.317.63M5.5 4.25a1.25 1.25 0 1 0-2.5 0a1.25 1.25 0 0 0 2.5 0m4.97 6.53a.75.75 0 1 1 1.06-1.06l1.75 1.75a.75.75 0 0 1 0 1.06l-1.75 1.75a.75.75 0 1 1-1.06-1.06l.47-.47H3.25a.75.75 0 0 1 0-1.5h7.69z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodePullRequestCheck(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.967 6.399a2.765 2.765 0 0 0 .88-1.24v-.002L8.22 6.53a.75.75 0 0 0 1.06-1.06L8.81 5h.69A1.501 1.501 0 0 1 11 6.5v1.75a.75.75 0 0 0 1.5 0V6.5a3 3 0 0 0-3-3h-.69l.47-.47a.75.75 0 0 0-1.06-1.06L6.847 3.343A2.755 2.755 0 0 0 4.25 1.5a2.75 2.75 0 0 0-.75 5.396v2.208a2.751 2.751 0 1 0 1.5 0V6.896c.357-.1.685-.272.967-.497m4.063 4.07a.75.75 0 1 0-1.06 1.061l2 2a.75.75 0 0 0 1.13-.08l3-4a.75.75 0 0 0-1.2-.9l-2.481 3.308l-1.389-1.388Zm-6.842.622A1.246 1.246 0 0 0 3 11.75a1.25 1.25 0 1 0 .188-.659m1.6-7.97a1.25 1.25 0 1 1-1.076 2.257a1.25 1.25 0 0 1 1.076-2.257" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodePullRequestXmark(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.967 6.399a2.765 2.765 0 0 0 .88-1.24v-.002L8.22 6.53a.75.75 0 0 0 1.06-1.06L8.81 5h.69A1.501 1.501 0 0 1 11 6.5v1.75a.75.75 0 1 0 1.5 0V6.5a3 3 0 0 0-3-3h-.69l.47-.47a.75.75 0 0 0-1.06-1.06L6.847 3.343A2.755 2.755 0 0 0 4.25 1.5a2.75 2.75 0 0 0-.753 5.396H3.5v2.208a2.751 2.751 0 1 0 1.5 0V6.896c.357-.1.685-.272.967-.497m3.253 6.82a.75.75 0 1 0 1.06 1.061l1.47-1.47l1.47 1.47a.75.75 0 1 0 1.06-1.06l-1.47-1.47l1.47-1.47a.75.75 0 1 0-1.06-1.06l-1.47 1.47l-1.47-1.47a.749.749 0 1 0-1.06 1.06l1.47 1.47l-1.47 1.47Zm-6.032-2.128A1.246 1.246 0 0 0 3 11.75a1.25 1.25 0 1 0 .188-.659m1.6-7.97a1.25 1.25 0 1 1-1.076 2.257a1.25 1.25 0 0 1 1.076-2.257" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeTrunk(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.312 4.909A1.25 1.25 0 1 1 3.19 3.593a1.25 1.25 0 0 1 2.123 1.316Zm-2.124 6.182a1.25 1.25 0 1 0 2.125 1.318a1.25 1.25 0 0 0-2.125-1.318M5 9.104V6.896a2.751 2.751 0 1 0-1.5 0v2.208A2.751 2.751 0 1 0 6.896 12.5H9.5a3 3 0 0 0 3-2.98V6.895a2.751 2.751 0 1 0-1.5 0V9.5A1.5 1.5 0 0 1 9.5 11H6.896A2.756 2.756 0 0 0 5 9.104m7.812-4.195a1.25 1.25 0 1 0-2.125-1.318a1.25 1.25 0 0 0 2.125 1.318" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comment(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m4.843 10.944l-.194 2.335a.204.204 0 0 0 .339.17l2.21-1.964l.589.013L8 11.5c1.695 0 3.087-.44 4.02-1.177c.89-.702 1.48-1.76 1.48-3.323s-.59-2.62-1.48-3.323C11.087 2.94 9.695 2.5 8 2.5c-1.695 0-3.087.44-4.02 1.177C3.09 4.38 2.5 5.437 2.5 7c0 1.648.656 2.742 1.648 3.448zm1.141 3.625l1.77-1.572C7.834 13 7.916 13 8 13c3.866 0 7-2 7-6s-3.134-6-7-6s-7 2-7 6c0 2.117.878 3.674 2.277 4.67l-.123 1.484a1.704 1.704 0 0 0 2.83 1.415" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentDot(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.98 3.677C4.913 2.94 6.305 2.5 8 2.5a9.1 9.1 0 0 1 .924.046a.75.75 0 0 0 .152-1.492A10.6 10.6 0 0 0 8 1c-1.933 0-3.683.5-4.95 1.5C1.784 3.5 1 5 1 7c0 2.117.878 3.674 2.277 4.67l-.123 1.484a1.704 1.704 0 0 0 2.83 1.415l1.77-1.572C7.834 13 7.916 13 8 13c1.933 0 3.683-.5 4.95-1.5C14.216 10.5 15 9 15 7a.75.75 0 0 0-1.5 0c0 1.563-.59 2.62-1.48 3.323C11.087 11.06 9.695 11.5 8 11.5c-.072 0-.143 0-.213-.002l-.295-.007l-.295-.006l-.22.195l-1.99 1.768a.204.204 0 0 1-.338-.17l.159-1.909l.035-.425l-.348-.248l-.347-.248C3.156 9.742 2.5 8.648 2.5 7c0-1.563.59-2.62 1.48-3.323M12.5 5.5a2 2 0 1 0 0-4a2 2 0 0 0 0 4" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 13c3.866 0 7-2 7-6s-3.134-6-7-6s-7 2-7 6c0 2.16.914 3.737 2.364 4.73l.09 2.161a1.157 1.157 0 0 0 1.857.872l2.322-1.77c.122.005.244.007.367.007" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentPlus(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m4.772 11.795l.071-.851l-.695-.496C3.156 9.742 2.5 8.648 2.5 7c0-1.563.59-2.62 1.48-3.323C4.913 2.94 6.305 2.5 8 2.5c1.695 0 3.087.44 4.02 1.177c.89.702 1.48 1.76 1.48 3.323s-.59 2.62-1.48 3.323C11.087 11.06 9.695 11.5 8 11.5c-.072 0-.143 0-.213-.002l-.59-.013l-.44.391l-1.77 1.572a.204.204 0 0 1-.338-.17zm2.981 1.202L5.984 14.57a1.704 1.704 0 0 1-2.83-1.415l.123-1.484C1.877 10.674 1 9.117 1 7c0-4 3.134-6 7-6s7 2 7 6s-3.134 6-7 6c-.083 0-.165 0-.247-.003M8.75 5a.75.75 0 0 0-1.5 0v1.25H6a.75.75 0 0 0 0 1.5h1.25V9a.75.75 0 0 0 1.5 0V7.75H10a.75.75 0 0 0 0-1.5H8.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentSlash(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.235 3.296L1.97 3.03a.75.75 0 0 1 1.06-1.06l11 11a.75.75 0 0 1-1.06 1.06l-1.621-1.62c-.995.393-2.136.59-3.349.59c-.083 0-.165 0-.247-.003L5.984 14.57a1.704 1.704 0 0 1-2.83-1.416l.123-1.484C1.877 10.674 1 9.117 1 7c0-1.526.456-2.76 1.235-3.704m7.93 7.93a8.14 8.14 0 0 1-2.378.272l-.59-.013l-.44.391l-1.77 1.572a.204.204 0 0 1-.338-.169l.123-1.484l.071-.851l-.695-.496C3.156 9.742 2.5 8.648 2.5 7c0-1.12.303-1.98.802-2.637l6.863 6.862ZM13.5 7c0 1.075-.28 1.91-.742 2.556l1.07 1.07C14.568 9.693 15 8.484 15 7c0-4-3.134-6-7-6c-1.172 0-2.276.184-3.247.551l1.194 1.194A8.231 8.231 0 0 1 8 2.5c1.695 0 3.087.44 4.02 1.177c.89.702 1.48 1.76 1.48 3.323" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comments(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 9.5h.621l.44.44l1.51 1.51a.174.174 0 0 0 .295-.136l-.112-1.454l-.062-.809l.642-.495C14.037 8.016 14.5 7.211 14.5 6c0-1.214-.465-2.019-1.17-2.56c-.754-.578-1.902-.94-3.33-.94c-1.428 0-2.576.362-3.33.94C5.966 3.98 5.5 4.786 5.5 6s.465 2.019 1.17 2.56c.754.578 1.902.94 3.33.94m.52 2.02l.99.99a1.673 1.673 0 0 0 2.851-1.312l-.111-1.453C15.33 8.91 16 7.663 16 6c0-3.333-2.686-5-6-5c-2.127 0-3.995.687-5.06 2.06C2.131 3.384 0 5.03 0 8c0 1.663.669 2.911 1.75 3.745l-.111 1.453A1.673 1.673 0 0 0 4.49 14.51L6 13c1.803 0 3.42-.493 4.52-1.48M4.143 4.736C4.05 5.126 4 5.548 4 6c0 2.905 2.04 4.544 4.759 4.918c-.717.366-1.654.582-2.759.582h-.621l-.44.44l-1.51 1.51a.174.174 0 0 1-.295-.136l.112-1.454l.062-.809l-.642-.495C1.963 10.016 1.5 9.211 1.5 8c0-1.214.465-2.019 1.17-2.56c.391-.3.887-.541 1.473-.704" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compass(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0m-6.09 2.303l-2.899.805a.909.909 0 0 1-1.12-1.119l.806-2.9A2 2 0 0 1 7.09 5.697l2.9-.805a.909.909 0 0 1 1.12 1.119l-.806 2.9a2 2 0 0 1-1.392 1.392ZM9 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copy(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.5H8A1.5 1.5 0 0 0 6.5 4v1H8a3 3 0 0 1 3 3v1.5h1A1.5 1.5 0 0 0 13.5 8V4A1.5 1.5 0 0 0 12 2.5M11 11h1a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3H8a3 3 0 0 0-3 3v1H4a3 3 0 0 0-3 3v4a3 3 0 0 0 3 3h4a3 3 0 0 0 3-3zM4 6.5h4A1.5 1.5 0 0 1 9.5 8v4A1.5 1.5 0 0 1 8 13.5H4A1.5 1.5 0 0 1 2.5 12V8A1.5 1.5 0 0 1 4 6.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyArrowRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 2.5h4A1.5 1.5 0 0 1 13.5 4v4A1.5 1.5 0 0 1 12 9.5h-1V8a3 3 0 0 0-3-3H6.5V4A1.5 1.5 0 0 1 8 2.5m4 8.5h-1v1a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3V8a3 3 0 0 1 3-3h1V4a3 3 0 0 1 3-3h4a3 3 0 0 1 3 3v4a3 3 0 0 1-3 3M8 6.5H4A1.5 1.5 0 0 0 2.5 8v4A1.5 1.5 0 0 0 4 13.5h4A1.5 1.5 0 0 0 9.5 12V8A1.5 1.5 0 0 0 8 6.5m-1.47 6.03a.75.75 0 0 1-1.06-1.06l.72-.72H4a.75.75 0 0 1 0-1.5h2.19l-.72-.72a.75.75 0 0 1 1.06-1.06l2 2a.75.75 0 0 1 0 1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyCheck(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m8 9.076l.085-.107a.751.751 0 1 0-1.171-.937L5.438 9.877L5.03 9.47a.747.747 0 0 0-1.06 0a.75.75 0 0 0 0 1.06l.407.408l.593.592a.75.75 0 0 0 1.116-.061l.522-.654h.001L8 9.074Z"/><path fill-rule="evenodd" d="M12 11a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3H8a3 3 0 0 0-3 3v1H4a3 3 0 0 0-3 3v4a3 3 0 0 0 3 3h4a3 3 0 0 0 3-3v-1zM4 6.5A1.5 1.5 0 0 0 2.5 8v4A1.5 1.5 0 0 0 4 13.5h4A1.5 1.5 0 0 0 9.5 12V8A1.498 1.498 0 0 0 8 6.5zM13.5 4A1.5 1.5 0 0 0 12 2.5H8A1.5 1.5 0 0 0 6.5 4v1H8a3 3 0 0 1 3 3v1.5h1A1.498 1.498 0 0 0 13.5 8z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyCheckXmark(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 6.5A1.5 1.5 0 0 0 2.5 8v4A1.5 1.5 0 0 0 4 13.5h4A1.5 1.5 0 0 0 9.5 12V8A1.498 1.498 0 0 0 8 6.5zm6-1.56l.97-.97a.75.75 0 0 1 1.06 1.06s.001 0 0 0l-.97.97l.97.97A.75.75 0 0 1 11 8.059V9.5h1A1.498 1.498 0 0 0 13.5 8V4c0-.414-.168-.79-.44-1.06A1.496 1.496 0 0 0 12 2.5H8a1.495 1.495 0 0 0-1.415 1A1.498 1.498 0 0 0 6.5 4v1h1.441A.748.748 0 0 1 8 3.941a.75.75 0 0 1 1.03.028l.97.97ZM8 9.075l.085-.107L8 9.076Zm0 0l-1.39 1.738l-.002.001l-.522.654a.75.75 0 0 1-1.116.061l-1-1s0 .001 0 0a.75.75 0 0 1 1.06-1.06l.408.407l1.476-1.845a.75.75 0 1 1 1.171.937M15 8a3 3 0 0 1-3 3h-1v1a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3V8a3 3 0 0 1 3-3h1V4a3 3 0 0 1 3-3h4a3 3 0 0 1 3 3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyChevronRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 2.5h4A1.5 1.5 0 0 1 13.5 4v4A1.5 1.5 0 0 1 12 9.5h-1V8a3 3 0 0 0-3-3H6.5V4A1.5 1.5 0 0 1 8 2.5m4 8.5h-1v1a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3V8a3 3 0 0 1 3-3h1V4a3 3 0 0 1 3-3h4a3 3 0 0 1 3 3v4a3 3 0 0 1-3 3M8 6.5H4A1.5 1.5 0 0 0 2.5 8v4A1.5 1.5 0 0 0 4 13.5h4A1.5 1.5 0 0 0 9.5 12V8A1.5 1.5 0 0 0 8 6.5m-1.97 5.78a.75.75 0 0 1-1.06-1.06L6.19 10L4.97 8.78a.75.75 0 0 1 1.06-1.06l1.75 1.75a.75.75 0 0 1 0 1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyMinus(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.5H8A1.5 1.5 0 0 0 6.5 4v1H8a3 3 0 0 1 3 3v1.5h1A1.5 1.5 0 0 0 13.5 8V4A1.5 1.5 0 0 0 12 2.5M11 11h1a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3H8a3 3 0 0 0-3 3v1H4a3 3 0 0 0-3 3v4a3 3 0 0 0 3 3h4a3 3 0 0 0 3-3zM8 6.5H4A1.5 1.5 0 0 0 2.5 8v4A1.5 1.5 0 0 0 4 13.5h4A1.5 1.5 0 0 0 9.5 12V8A1.5 1.5 0 0 0 8 6.5M3.75 10a.75.75 0 0 1 .75-.75h3a.75.75 0 0 1 0 1.5h-3a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyPicture(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.5H8A1.5 1.5 0 0 0 6.5 4v1H8a3 3 0 0 1 3 3v1.5h1A1.5 1.5 0 0 0 13.5 8V4A1.5 1.5 0 0 0 12 2.5M11 11h1a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3H8a3 3 0 0 0-3 3v1H4a3 3 0 0 0-3 3v4a3 3 0 0 0 3 3h4a3 3 0 0 0 3-3zM4 6.5h4A1.5 1.5 0 0 1 9.5 8v2.041l-.568-.52a1.52 1.52 0 0 0-2.014-.032l-1.719 1.473l-.602-.452a1.52 1.52 0 0 0-1.805-.013l-.292.212V8A1.5 1.5 0 0 1 4 6.5m-1.414 6.001A1.5 1.5 0 0 0 4 13.5h4a1.5 1.5 0 0 0 1.498-1.428L7.92 10.63a.02.02 0 0 0-.026 0l-2.175 1.864l-.457.391l-.481-.36l-1.084-.814a.02.02 0 0 0-.023 0l-1.088.791Zm3.03-4.04a1.154 1.154 0 1 1-2.308 0a1.154 1.154 0 0 1 2.307 0Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyPlus(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.5H8A1.5 1.5 0 0 0 6.5 4v1H8a3 3 0 0 1 3 3v1.5h1A1.5 1.5 0 0 0 13.5 8V4A1.5 1.5 0 0 0 12 2.5M11 11h1a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3H8a3 3 0 0 0-3 3v1H4a3 3 0 0 0-3 3v4a3 3 0 0 0 3 3h4a3 3 0 0 0 3-3zM8 6.5H4A1.5 1.5 0 0 0 2.5 8v4A1.5 1.5 0 0 0 4 13.5h4A1.5 1.5 0 0 0 9.5 12V8A1.5 1.5 0 0 0 8 6.5M6 7.75a.75.75 0 0 1 .75.75v.75h.75a.75.75 0 0 1 0 1.5h-.75v.75a.75.75 0 0 1-1.5 0v-.75H4.5a.75.75 0 0 1 0-1.5h.75V8.5A.75.75 0 0 1 6 7.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyTransparent(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 2.5h4A1.5 1.5 0 0 1 13.5 4v4A1.5 1.5 0 0 1 12 9.5h-1V8a3 3 0 0 0-3-3H6.5V4A1.5 1.5 0 0 1 8 2.5M5 5V4a3 3 0 0 1 3-3h4a3 3 0 0 1 3 3v4a3 3 0 0 1-3 3h-1v1a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3V8a3 3 0 0 1 3-3zm4.5 6v1A1.5 1.5 0 0 1 8 13.5H4A1.5 1.5 0 0 1 2.5 12V8A1.5 1.5 0 0 1 4 6.5h1V8a3 3 0 0 0 3 3zm0-1.5H8A1.5 1.5 0 0 1 6.5 8V6.5H8A1.5 1.5 0 0 1 9.5 8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyXmark(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.5 12c0 .414-.168.79-.44 1.06A1.49 1.49 0 0 1 8 13.5H4a1.49 1.49 0 0 1-1.06-.44A1.496 1.496 0 0 1 2.5 12V8c0-.414.168-.79.44-1.06A1.49 1.49 0 0 1 4 6.5h4c.414 0 .79.168 1.06.44c.272.27.44.646.44 1.06zm-1.47-1.03s.001 0 0 0L7.06 10l.97-.97a.75.75 0 0 0-1.06-1.06L6 8.94l-.97-.97a.75.75 0 0 0-1.06 1.06l.97.97l-.97.97a.75.75 0 0 0 1.06 1.06c0 .001 0 0 0 0l.97-.97l.97.97a.75.75 0 0 0 1.06-1.06M6.5 5H8a3 3 0 0 1 3 3v1.5h1A1.498 1.498 0 0 0 13.5 8V4A1.5 1.5 0 0 0 12 2.5H8A1.5 1.5 0 0 0 6.5 4zM5 4a3 3 0 0 1 3-3h4a3 3 0 0 1 3 3v4a3 3 0 0 1-3 3h-1v1a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3V8a3 3 0 0 1 3-3h1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cpu(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 1.25a.75.75 0 0 0-1.5 0V2.5a3 3 0 0 0-3 3H1.25a.75.75 0 0 0 0 1.5H2.5v2H1.25a.75.75 0 0 0 0 1.5H2.5a3 3 0 0 0 3 3v1.25a.75.75 0 0 0 1.5 0V13.5h2v1.25a.75.75 0 0 0 1.5 0V13.5a3 3 0 0 0 3-3h1.25a.75.75 0 1 0 0-1.5H13.5V7h1.25a.75.75 0 1 0 0-1.5H13.5a3 3 0 0 0-3-3V1.25a.75.75 0 0 0-1.5 0V2.5H7zM10.5 4h-5A1.5 1.5 0 0 0 4 5.5v5A1.5 1.5 0 0 0 5.5 12h5a1.5 1.5 0 0 0 1.5-1.5v-5A1.5 1.5 0 0 0 10.5 4m0 2.25a.75.75 0 0 0-.75-.75h-3.5a.75.75 0 0 0-.75.75v3.5a.75.75 0 0 0 .75.75h3.5a.75.75 0 0 0 .75-.75zM7 7h2v2H7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cpus(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.75 0a.75.75 0 0 1 .75.75V2h2V.75a.75.75 0 0 1 1.5 0V2a2 2 0 0 1 2 2h1.25a.75.75 0 0 1 0 1.5H14v2h1.25a.75.75 0 0 1 0 1.5H14a2 2 0 0 1-2 2h-1v1a2 2 0 0 1-2 2v1.25a.75.75 0 0 1-1.5 0V14h-2v1.25a.75.75 0 0 1-1.5 0V14a2 2 0 0 1-2-2H.75a.75.75 0 0 1 0-1.5H2v-2H.75a.75.75 0 0 1 0-1.5H2a2 2 0 0 1 2-2h1V4a2 2 0 0 1 2-2V.75A.75.75 0 0 1 7.75 0M6.5 5H9a2 2 0 0 1 2 2v2.5h1a.5.5 0 0 0 .5-.5V4a.5.5 0 0 0-.5-.5H7a.5.5 0 0 0-.5.5zm3 2a.5.5 0 0 0-.5-.5H4a.5.5 0 0 0-.5.5v5a.5.5 0 0 0 .5.5h5a.5.5 0 0 0 .5-.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCard(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 4h-9A1.5 1.5 0 0 0 2 5.5h12A1.5 1.5 0 0 0 12.5 4M2 10.5V7h12v3.5a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 2 10.5m1.5-8a3 3 0 0 0-3 3v5a3 3 0 0 0 3 3h9a3 3 0 0 0 3-3v-5a3 3 0 0 0-3-3zM4.25 9a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crop(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 0a.75.75 0 0 1 .75.75V2.5h8.25a.75.75 0 0 1 .75.75v8.5h2a.75.75 0 0 1 0 1.5h-2v2a.75.75 0 0 1-1.5 0v-2H3.5a.75.75 0 0 1-.75-.75V4h-2a.75.75 0 0 1 0-1.5h2V.75A.75.75 0 0 1 3.5 0m.75 4v7.75h7.5V4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CrownDiamond(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="m8.427 11.073l1.205-1.205a.402.402 0 0 0 .118-.285a.805.805 0 0 0-.236-.569L8.427 7.927a.603.603 0 0 0-.854 0L6.486 9.014a.805.805 0 0 0-.236.57c0 .106.042.208.118.284l1.205 1.205a.604.604 0 0 0 .854 0"/><path d="M16 5.796v-.028a1.768 1.768 0 0 0-3.018-1.25l-.76.76l-.024.024l-.374.374l-.415.415a.335.335 0 0 1-.561-.149l-.155-.566l-.139-.51l-.009-.033l-.65-2.386a1.964 1.964 0 0 0-3.79 0l-.65 2.386l-.01.032l-.139.511l-.154.566a.335.335 0 0 1-.56.15l-.416-.416l-.374-.374l-.024-.024l-.76-.76A1.768 1.768 0 0 0 0 5.768v.028c0 .135.015.27.046.403l1.3 5.631a1.4 1.4 0 0 0 .778.958a14.021 14.021 0 0 0 11.752 0c.394-.182.681-.535.779-.958l1.299-5.63A1.84 1.84 0 0 0 16 5.796M3.53 7.152c.997.997 2.698.545 3.07-.815l.952-3.495a.464.464 0 0 1 .896 0L9.4 6.337c.37 1.36 2.072 1.812 3.068.815l1.574-1.574a.268.268 0 0 1 .457.19v.028a.291.291 0 0 1-.008.066l-1.288 5.584a12.522 12.522 0 0 1-10.408 0L1.508 5.862a.295.295 0 0 1-.008-.066v-.028a.268.268 0 0 1 .457-.19l1.574 1.574Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cube(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 5.475v4.946a1.5 1.5 0 0 1-.973 1.405l-4.777 1.79V7.53l5.75-2.054Zm-.654-1.36a1.5 1.5 0 0 0-.175-.103L9.499 2.427a1.5 1.5 0 0 0-1.197-.063l-4.829 1.81c-.08.03-.157.067-.23.11L7.05 6.185zM2.5 5.59l3.75 1.875v5.984l-2.92-1.46a1.5 1.5 0 0 1-.83-1.342zM1.266 4.343c-.172.38-.266.8-.266 1.236v5.067a3 3 0 0 0 1.658 2.683l3.172 1.586a3 3 0 0 0 2.395.126l4.828-1.811A3 3 0 0 0 15 10.421V5.354a3 3 0 0 0-1.658-2.683L10.17 1.085A3 3 0 0 0 7.775.959L2.947 2.77a2.997 2.997 0 0 0-1.48 1.203a.75.75 0 0 0-.2.37Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CubesThree(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.5 2h3a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1h-3a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1M4 3A2.5 2.5 0 0 1 6.5.5h3A2.5 2.5 0 0 1 12 3v3c0 .356-.074.694-.208 1H13a2.5 2.5 0 0 1 2.5 2.5v3A2.5 2.5 0 0 1 13 15H3a2.5 2.5 0 0 1-2.5-2.5v-3A2.5 2.5 0 0 1 3 7h1.208A2.492 2.492 0 0 1 4 6zm2.25 5.5a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1v-3a1 1 0 0 1 1-1zm3.5 5a1 1 0 0 1-1-1v-3a1 1 0 0 1 1-1H13a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CubesThreeOverlap(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.5 2.5h-3A1.5 1.5 0 0 0 5 4v1h1a3 3 0 0 1 3 3v.5h.5A1.5 1.5 0 0 0 11 7V4a1.5 1.5 0 0 0-1.5-1.5M9 10v1c0 .546-.146 1.059-.401 1.5H13a1.5 1.5 0 0 0 1.5-1.5V8A1.5 1.5 0 0 0 13 6.5h-.5V7a3 3 0 0 1-3 3zm3.5-5V4a3 3 0 0 0-3-3h-3a3 3 0 0 0-3 3v1H3a3 3 0 0 0-3 3v3a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3zM6 6.5H3A1.5 1.5 0 0 0 1.5 8v3A1.5 1.5 0 0 0 3 12.5h3A1.5 1.5 0 0 0 7.5 11V8A1.5 1.5 0 0 0 6 6.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cup(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiCup0)"><path fill="currentColor" fill-rule="evenodd" d="M10.521 5.835c.84-.402.979-.795.979-1.085c0-.272-.123-.634-.83-1.01a4.02 4.02 0 0 0-.149-.075C9.666 3.255 8.351 3 6.75 3s-2.916.256-3.772.665C2.14 4.067 2 4.46 2 4.75c0 .29.14.683.978 1.085c.856.41 2.171.665 3.772.665s2.916-.256 3.771-.665M6.75 8c1.882 0 3.57-.322 4.715-.966c-.059 1.35-.216 2.595-.634 3.616c-.303.74-.72 1.293-1.296 1.674c-.579.383-1.444.676-2.785.676c-1.34 0-2.206-.293-2.785-.676c-.575-.381-.992-.934-1.296-1.674c-.418-1.021-.575-2.267-.634-3.616C3.18 7.678 4.868 8 6.75 8m6.246-3.388C12.88 2.537 10.128 1.5 6.75 1.5C3.298 1.5.5 2.583.5 4.75C.5 9 .5 14.5 6.75 14.5c3.36 0 4.913-1.589 5.632-3.719a9.67 9.67 0 0 0 .818-.293c.615-.256 1.268-.613 1.79-1.112C15.524 8.865 16 8.12 16 7.158c0-.437-.097-.89-.345-1.304a2.38 2.38 0 0 0-.956-.89c-.563-.293-1.187-.358-1.703-.352m-.003 1.5c-.015.97-.06 1.967-.204 2.92c.964-.432 1.711-1.056 1.711-1.874c0-.808-.656-1.059-1.507-1.046" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiCup0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurlyBrackets(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.413 3.421A.841.841 0 0 1 4.25 2.5a.75.75 0 0 0 0-1.5a2.341 2.341 0 0 0-2.33 2.563l.199 2.096a.898.898 0 0 1-.677.957l-.139.035a1.39 1.39 0 0 0 0 2.698l.14.035a.898.898 0 0 1 .676.957l-.2 2.096A2.341 2.341 0 0 0 4.25 15a.75.75 0 0 0 0-1.5a.841.841 0 0 1-.837-.921l.2-2.096A2.399 2.399 0 0 0 2.04 8a2.399 2.399 0 0 0 1.571-2.483l-.2-2.096Zm9.175 9.158a.841.841 0 0 1-.838.921a.75.75 0 0 0 0 1.5a2.341 2.341 0 0 0 2.33-2.563l-.199-2.096a.898.898 0 0 1 .677-.957l.139-.035a1.39 1.39 0 0 0 0-2.698l-.14-.035a.898.898 0 0 1-.676-.957l.2-2.096A2.341 2.341 0 0 0 11.75 1a.75.75 0 0 0 0 1.5c.496 0 .884.427.838.921l-.2 2.096A2.399 2.399 0 0 0 13.959 8a2.399 2.399 0 0 0-1.571 2.483z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurlyBracketsFunction(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.25 2.5a.841.841 0 0 0-.837.921l.2 2.096A2.399 2.399 0 0 1 2.04 8a2.399 2.399 0 0 1 1.571 2.483l-.2 2.096a.841.841 0 0 0 .838.921a.75.75 0 0 1 0 1.5a2.341 2.341 0 0 1-2.33-2.563l.199-2.096a.898.898 0 0 0-.677-.957l-.139-.035a1.39 1.39 0 0 1 0-2.698l.14-.035a.898.898 0 0 0 .676-.957l-.2-2.096A2.341 2.341 0 0 1 4.25 1a.75.75 0 0 1 0 1.5m4.805 2.639a.78.78 0 0 1 .989-.668a.75.75 0 1 0 .412-1.442a2.279 2.279 0 0 0-2.892 1.953L7.456 6H6.5a.75.75 0 0 0 0 1.5h.798l-.353 3.361a.779.779 0 0 1-.989.668a.75.75 0 1 0-.412 1.442a2.279 2.279 0 0 0 2.892-1.953l.37-3.518h.944a.75.75 0 0 0 0-1.5h-.785zm3.533 7.44a.841.841 0 0 1-.838.921a.75.75 0 0 0 0 1.5a2.341 2.341 0 0 0 2.33-2.563l-.199-2.096a.898.898 0 0 1 .677-.957l.139-.035a1.39 1.39 0 0 0 0-2.698l-.14-.035a.898.898 0 0 1-.676-.957l.2-2.096A2.341 2.341 0 0 0 11.75 1a.75.75 0 0 0 0 1.5c.496 0 .884.427.838.921l-.2 2.096A2.399 2.399 0 0 0 13.959 8a2.399 2.399 0 0 0-1.571 2.483z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurlyBracketsLock(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.413 3.421A.841.841 0 0 1 4.25 2.5a.75.75 0 0 0 0-1.5a2.341 2.341 0 0 0-2.33 2.563l.199 2.096a.898.898 0 0 1-.677.957l-.139.035a1.39 1.39 0 0 0 0 2.698l.14.035a.898.898 0 0 1 .676.957l-.2 2.096A2.341 2.341 0 0 0 4.25 15a.75.75 0 0 0 0-1.5a.841.841 0 0 1-.837-.921l.2-2.096A2.399 2.399 0 0 0 2.04 8a2.399 2.399 0 0 0 1.571-2.483l-.2-2.096Zm9.175 9.158a.841.841 0 0 1-.838.921a.75.75 0 0 0 0 1.5a2.341 2.341 0 0 0 2.33-2.563l-.199-2.096a.898.898 0 0 1 .677-.957l.139-.035a1.39 1.39 0 0 0 0-2.698l-.14-.035a.898.898 0 0 1-.676-.957l.2-2.096A2.341 2.341 0 0 0 11.75 1a.75.75 0 0 0 0 1.5c.496 0 .884.427.838.921l-.2 2.096A2.399 2.399 0 0 0 13.959 8a2.399 2.399 0 0 0-1.571 2.483zM6 8v1.5h4V8zm-.1-1.5a1.4 1.4 0 0 0-1.4 1.4v1.7A1.4 1.4 0 0 0 5.9 11h4.2a1.4 1.4 0 0 0 1.4-1.4V7.9a1.4 1.4 0 0 0-1.4-1.4H10V6a2 2 0 1 0-4 0v.5zm1.6 0h1V6a.5.5 0 0 0-1 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Database(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.615 4.888c.814-.375.885-.714.885-.888c0-.174-.071-.513-.885-.888C10.8 2.737 9.538 2.5 8 2.5c-1.538 0-2.799.237-3.615.612c-.814.375-.885.714-.885.888c0 .174.071.513.885.888C5.2 5.263 6.462 5.5 8 5.5c1.538 0 2.799-.237 3.615-.612m.885 1.235C11.4 6.708 9.792 7 8 7c-1.792 0-3.4-.292-4.5-.877V8c0 .174.071.513.885.888C5.2 9.263 6.462 9.5 8 9.5c1.538 0 2.799-.237 3.615-.612c.814-.375.885-.714.885-.888zm0 4C11.4 10.708 9.792 11 8 11c-1.792 0-3.4-.293-4.5-.877V12c0 .174.071.513.885.887c.816.377 2.077.613 3.615.613c1.538 0 2.799-.236 3.615-.613c.814-.374.885-.713.885-.887zM14 4c0-2-2.686-3-6-3S2 2 2 4v8c0 2 2.686 3 6 3s6-1 6-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseArrowRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 4c0 .174-.071.513-.885.888C9.8 5.263 8.538 5.5 7 5.5c-1.538 0-2.799-.237-3.615-.612C2.57 4.513 2.5 4.174 2.5 4c0-.174.071-.513.885-.888C4.2 2.737 5.462 2.5 7 2.5c1.538 0 2.799.237 3.615.612c.814.375.885.714.885.888M3.385 8.888C2.578 8.516 2.5 8.179 2.5 8V6.123C3.6 6.708 5.208 7 7 7c1.792 0 3.4-.292 4.5-.877V7.25a.75.75 0 0 0 1.5 0V4c0-2-2.686-3-6-3S1 2 1 4v8c0 .995.665 1.747 1.757 2.25l.314-.681l-.314.681c1.086.5 2.586.75 4.243.75c.088 0 .176 0 .263-.002a.75.75 0 0 0-.025-1.5L7 13.5c-1.538 0-2.799-.236-3.615-.613c-.807-.371-.885-.708-.885-.887v-1.877c.083.044.169.086.257.127c1.086.5 2.586.75 4.243.75c.09 0 .178 0 .266-.002a.75.75 0 1 0-.024-1.5L7 9.5c-1.538 0-2.799-.237-3.615-.612M13.78 14.53a.75.75 0 1 1-1.06-1.06l.72-.72H10a.75.75 0 0 1 0-1.5h3.44l-.72-.72a.75.75 0 1 1 1.06-1.06l2 2a.75.75 0 0 1 0 1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 8.5c1.776 0 3.515-.263 4.87-.888A5.64 5.64 0 0 0 14 6.931V8c0 2-2.686 3-6 3s-6-1-6-3V6.93c.35.275.736.501 1.13.682C4.485 8.237 6.224 8.5 8 8.5M14 4c0 2-2.686 3-6 3S2 6 2 4c0-.336.076-.643.217-.923C2.92 1.692 5.242 1 8 1c.828 0 1.618.063 2.335.188C12.49 1.563 14 2.5 14 4M8 15c3.314 0 6-1 6-3v-1.07c-.35.275-.736.501-1.13.683c-1.355.623-3.094.887-4.87.887c-1.776 0-3.515-.264-4.87-.887A5.701 5.701 0 0 1 2 10.93V12c0 2 2.686 3 6 3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseMagnifier(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.615 4.888c.814-.375.885-.714.885-.888c0-.174-.071-.513-.885-.888C10.8 2.737 9.538 2.5 8 2.5c-1.538 0-2.799.237-3.615.612c-.814.375-.885.714-.885.888c0 .174.071.513.885.888C5.2 5.263 6.462 5.5 8 5.5c1.538 0 2.799-.237 3.615-.612m.885 1.235C11.4 6.708 9.792 7 8 7c-1.792 0-3.4-.292-4.5-.877V8c0 .174.071.513.885.888C5.2 9.263 6.462 9.5 8 9.5h.04c-.187.463-.29.968-.29 1.498c-1.691-.026-3.202-.318-4.25-.875V12c0 .174.071.513.885.887c.816.377 2.077.613 3.615.613c.211 0 .417-.005.617-.013a4.008 4.008 0 0 0 1.848 1.302C9.713 14.93 8.879 15 8 15c-3.314 0-6-1-6-3V4c0-2 2.686-3 6-3s6 1 6 3v3.692a3.978 3.978 0 0 0-1.5-.622zm-.75 6.377a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0 1.5c.556 0 1.077-.152 1.524-.415l1.446 1.445a.75.75 0 1 0 1.06-1.06l-1.445-1.446A3 3 0 1 0 11.75 14" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Databases(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.499 10.04c.492-.252.501-.446.501-.54V5.462c-.796.359-1.848.538-3 .538a9.467 9.467 0 0 1-1.606-.13c.07.195.106.405.106.63v3.864c.428.086.932.136 1.5.136c1.107 0 1.971-.19 2.499-.46M8 8.462V12.5c0 .094-.009.288-.501.54c-.528.27-1.392.46-2.499.46c-1.107 0-1.971-.19-2.499-.46C2.01 12.788 2 12.594 2 12.5V8.462C2.796 8.821 3.847 9 5 9s2.204-.18 3-.538M8 6.5c0-.093-.009-.288-.501-.54C6.97 5.69 6.107 5.5 5 5.5c-1.107 0-1.971.19-2.499.46C2.01 6.212 2 6.407 2 6.5c0 .093.009.288.501.54c.528.27 1.392.46 2.499.46c1.107 0 1.971-.19 2.499-.46C7.99 6.788 8 6.593 8 6.5m-7.5 0C.5 4.833 2.515 4 5 4c.526 0 1.03.037 1.5.112V3.5C6.5 1.833 8.515 1 11 1s4.5.833 4.5 2.5v6c0 1.667-2.015 2.5-4.5 2.5a9.56 9.56 0 0 1-1.5-.112v.612C9.5 14.167 7.485 15 5 15S.5 14.167.5 12.5zm7.5-3c0 .094.009.288.501.54c.528.27 1.392.46 2.499.46c1.107 0 1.971-.19 2.499-.46c.492-.252.501-.446.501-.54c0-.094-.009-.288-.501-.54c-.528-.27-1.392-.46-2.499-.46c-1.107 0-1.971.19-2.499.46C8.01 3.212 8 3.406 8 3.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Delete(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.442 3.5H12.5A1.5 1.5 0 0 1 14 5v6a1.5 1.5 0 0 1-1.5 1.5H5.442a1.5 1.5 0 0 1-1.171-.563L1.796 8.844a1.35 1.35 0 0 1 0-1.688l2.475-3.093A1.5 1.5 0 0 1 5.44 3.5Zm-2.343-.374A3 3 0 0 1 5.442 2H12.5a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3H5.442a3 3 0 0 1-2.343-1.126L.625 9.781a2.85 2.85 0 0 1 0-3.562zM7.28 5.47a.75.75 0 0 0-1.06 1.06L7.69 8L6.22 9.47a.75.75 0 1 0 1.06 1.06l1.47-1.47l1.47 1.47a.75.75 0 1 0 1.06-1.06L9.81 8l1.47-1.47a.75.75 0 0 0-1.06-1.06L8.75 6.94z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Diamond(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiDiamond0)"><path fill="currentColor" fill-rule="evenodd" d="M14.01 6.94L9.06 1.99a1.5 1.5 0 0 0-2.12 0L1.99 6.94a1.5 1.5 0 0 0 0 2.12l4.95 4.95a1.5 1.5 0 0 0 2.12 0l4.95-4.95a1.5 1.5 0 0 0 0-2.12M10.121.928a3 3 0 0 0-4.242 0l-4.95 4.95a3 3 0 0 0 0 4.242l4.95 4.95a3 3 0 0 0 4.242 0l4.95-4.95a3 3 0 0 0 0-4.242z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiDiamond0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiamondExclamation(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.01 6.94L9.06 1.99a1.5 1.5 0 0 0-2.12 0L1.99 6.94a1.5 1.5 0 0 0 0 2.12l4.95 4.95a1.5 1.5 0 0 0 2.12 0l4.95-4.95a1.5 1.5 0 0 0 0-2.12M10.121.928a3 3 0 0 0-4.242 0l-4.95 4.95a3 3 0 0 0 0 4.242l4.95 4.95a3 3 0 0 0 4.242 0l4.95-4.95a3 3 0 0 0 0-4.242zM9 10.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0M8.75 5a.75.75 0 1 0-1.5 0v2.5a.75.75 0 1 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiamondFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiDiamondFill0)"><path fill="currentColor" d="M5.879.929a3 3 0 0 1 4.242 0l4.95 4.95a3 3 0 0 1 0 4.242l-4.95 4.95a3 3 0 0 1-4.242 0l-4.95-4.95a3 3 0 0 1 0-4.242z"/></g><defs><clipPath id="gravityUiDiamondFill0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceFive(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 3h7A1.5 1.5 0 0 1 13 4.5v7a1.5 1.5 0 0 1-1.5 1.5h-7A1.5 1.5 0 0 1 3 11.5v-7A1.5 1.5 0 0 1 4.5 3m-3 1.5a3 3 0 0 1 3-3h7a3 3 0 0 1 3 3v7a3 3 0 0 1-3 3h-7a3 3 0 0 1-3-3zm9 2.15a1.15 1.15 0 1 0 0-2.3a1.15 1.15 0 0 0 0 2.3M9.15 8a1.15 1.15 0 1 1-2.3 0a1.15 1.15 0 0 1 2.3 0M5.5 11.65a1.15 1.15 0 1 0 0-2.3a1.15 1.15 0 0 0 0 2.3m6.15-1.15a1.15 1.15 0 1 1-2.3 0a1.15 1.15 0 0 1 2.3 0M5.5 6.65a1.15 1.15 0 1 0 0-2.3a1.15 1.15 0 0 0 0 2.3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceFour(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 3h-7A1.5 1.5 0 0 0 3 4.5v7A1.5 1.5 0 0 0 4.5 13h7a1.5 1.5 0 0 0 1.5-1.5v-7A1.5 1.5 0 0 0 11.5 3m-7-1.5a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h7a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3zM11.15 6a1.15 1.15 0 1 1-2.3 0a1.15 1.15 0 0 1 2.3 0M6 11.15a1.15 1.15 0 1 0 0-2.3a1.15 1.15 0 0 0 0 2.3M11.15 10a1.15 1.15 0 1 1-2.3 0a1.15 1.15 0 0 1 2.3 0M6 7.15a1.15 1.15 0 1 0 0-2.3a1.15 1.15 0 0 0 0 2.3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceOne(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 3h7A1.5 1.5 0 0 1 13 4.5v7a1.5 1.5 0 0 1-1.5 1.5h-7A1.5 1.5 0 0 1 3 11.5v-7A1.5 1.5 0 0 1 4.5 3m-3 1.5a3 3 0 0 1 3-3h7a3 3 0 0 1 3 3v7a3 3 0 0 1-3 3h-7a3 3 0 0 1-3-3zM8 9.15a1.15 1.15 0 1 0 0-2.3a1.15 1.15 0 0 0 0 2.3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceSix(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 3h-7A1.5 1.5 0 0 0 3 4.5v7A1.5 1.5 0 0 0 4.5 13h7a1.5 1.5 0 0 0 1.5-1.5v-7A1.5 1.5 0 0 0 11.5 3m-7-1.5a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h7a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3zM11.65 5a1.15 1.15 0 1 1-2.3 0a1.15 1.15 0 0 1 2.3 0M5.5 12.15a1.15 1.15 0 1 0 0-2.3a1.15 1.15 0 0 0 0 2.3M11.65 11a1.15 1.15 0 1 1-2.3 0a1.15 1.15 0 0 1 2.3 0M5.5 6.15a1.15 1.15 0 1 0 0-2.3a1.15 1.15 0 0 0 0 2.3M11.65 8a1.15 1.15 0 1 1-2.3 0a1.15 1.15 0 0 1 2.3 0M5.5 9.15a1.15 1.15 0 1 0 0-2.3a1.15 1.15 0 0 0 0 2.3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceThree(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 3h7A1.5 1.5 0 0 1 13 4.5v7a1.5 1.5 0 0 1-1.5 1.5h-7A1.5 1.5 0 0 1 3 11.5v-7A1.5 1.5 0 0 1 4.5 3m-3 1.5a3 3 0 0 1 3-3h7a3 3 0 0 1 3 3v7a3 3 0 0 1-3 3h-7a3 3 0 0 1-3-3zm9 2.15a1.15 1.15 0 1 0 0-2.3a1.15 1.15 0 0 0 0 2.3M9.15 8a1.15 1.15 0 1 1-2.3 0a1.15 1.15 0 0 1 2.3 0M5.5 11.65a1.15 1.15 0 1 0 0-2.3a1.15 1.15 0 0 0 0 2.3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceTwo(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 3h-7A1.5 1.5 0 0 0 3 4.5v7A1.5 1.5 0 0 0 4.5 13h7a1.5 1.5 0 0 0 1.5-1.5v-7A1.5 1.5 0 0 0 11.5 3m-7-1.5a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h7a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3zM11.15 6a1.15 1.15 0 1 1-2.3 0a1.15 1.15 0 0 1 2.3 0M6 11.15a1.15 1.15 0 1 0 0-2.3a1.15 1.15 0 0 0 0 2.3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Display(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 3H4a1.5 1.5 0 0 0-1.5 1.5v4A1.5 1.5 0 0 0 4 10h8a1.5 1.5 0 0 0 1.5-1.5v-4A1.5 1.5 0 0 0 12 3M4 1.5a3 3 0 0 0-3 3v4a3 3 0 0 0 3 3h3.25V13h-2.5a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5h-2.5v-1.5H12a3 3 0 0 0 3-3v-4a3 3 0 0 0-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DisplayPulse(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 3h8a1.5 1.5 0 0 1 1.5 1.5v4A1.5 1.5 0 0 1 12 10H4a1.5 1.5 0 0 1-1.5-1.5v-1h2.25a.75.75 0 0 0 .564-.256l1.057-1.208L7.85 8.622A.75.75 0 0 0 9.1 8.7L10.375 7h1.375a.75.75 0 0 0 0-1.5H10a.75.75 0 0 0-.6.3l-.815 1.087l-1.434-2.51a.75.75 0 0 0-1.215-.12L4.41 6H2.5V4.5A1.5 1.5 0 0 1 4 3M1 6.75V8.5a3 3 0 0 0 3 3h3.25V13h-2.5a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5h-2.5v-1.5H12a3 3 0 0 0 3-3v-4a3 3 0 0 0-3-3H4a3 3 0 0 0-3 3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotsNine(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.5 3a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0M3 9.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3M9.5 8a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0m5 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0M13 4.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3M4.5 3a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0M8 14.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m6.5-1.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0M3 14.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Droplet(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 9a4.5 4.5 0 1 1-9 0c0-1.638.761-3.242 1.789-4.64a16.57 16.57 0 0 1 2.714-2.847c.591.48 1.722 1.483 2.707 2.817c1.024 1.389 1.79 3 1.79 4.67M14 9A6 6 0 0 1 2 9C2 4.819 5.846 1.337 7.106.309C7.359.102 7.676 0 8.003 0c.323 0 .637.1.888.302C10.148 1.312 14 4.759 14 9m-6 3.25a.75.75 0 0 1 0-1.5A1.75 1.75 0 0 0 9.75 9a.75.75 0 0 1 1.5 0A3.25 3.25 0 0 1 8 12.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ear(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 12.372V5.75C4 3.419 5.934 1.5 8.298 1.5c2.31 0 4.202 1.877 4.202 4.156v.459a4.773 4.773 0 0 1-1.667 3.624a11.26 11.26 0 0 0-2.486 3.028l-.364.648A2.128 2.128 0 0 1 4 12.372M2.5 5.75v6.622a3.628 3.628 0 0 0 6.79 1.779l.365-.648a9.76 9.76 0 0 1 2.154-2.625A6.274 6.274 0 0 0 14 6.115v-.46C14 2.533 11.421 0 8.298 0C5.122 0 2.5 2.574 2.5 5.75m4.526-.162a1.242 1.242 0 0 1 2.474.154v.508a.75.75 0 0 0 1.5 0v-.508a2.742 2.742 0 0 0-5.462-.34l-.282 2.255a.75.75 0 0 0 .827.838l.85-.094a.511.511 0 0 1 .383.901l-1.046.872a.75.75 0 0 0 .96 1.152l1.046-.872a2.011 2.011 0 0 0-1.414-3.552z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ellipsis(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 9.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3M9.5 8a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0m5 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisVertical(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 4.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3M9.5 8a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0m0 5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Envelope(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 4h9c.25 0 .485.06.692.169L8.75 7.5a1.25 1.25 0 0 1-1.5 0L2.808 4.169C3.015 4.06 3.251 4 3.5 4M2.001 5.438L2 5.5v5A1.5 1.5 0 0 0 3.5 12h9a1.5 1.5 0 0 0 1.5-1.5v-5l-.001-.062L9.65 8.7a2.75 2.75 0 0 1-3.3 0zM.5 5.5a3 3 0 0 1 3-3h9a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3h-9a3 3 0 0 1-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeOpen(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14 6.498V11.5a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 2 11.5V6.498l.001-.06L6.35 9.7a2.75 2.75 0 0 0 3.3 0l4.349-3.262zm-.806-1.33L8.74 2.642a1.5 1.5 0 0 0-1.48 0L2.806 5.167L7.25 8.5a1.25 1.25 0 0 0 1.5 0l4.444-3.333ZM.5 6.497a3 3 0 0 1 1.521-2.61l4.5-2.55a3 3 0 0 1 2.958 0l4.5 2.55a3 3 0 0 1 1.521 2.61V11.5a3 3 0 0 1-3 3h-9a3 3 0 0 1-3-3V6.498Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeOpenXmark(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14 6.498V11.5a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 2 11.5V6.498c0-.54.29-1.039.76-1.305l4.5-2.55a1.5 1.5 0 0 1 1.48 0l4.5 2.55c.47.266.76.765.76 1.305m-13.5 0a3 3 0 0 1 1.521-2.61l4.5-2.55a3 3 0 0 1 2.958 0l4.5 2.55a3 3 0 0 1 1.521 2.61V11.5a3 3 0 0 1-3 3h-9a3 3 0 0 1-3-3zm6.03-.528a.75.75 0 0 0-1.06 1.06L6.94 8.5L5.47 9.97a.75.75 0 1 0 1.06 1.06L8 9.56l1.47 1.47a.75.75 0 1 0 1.06-1.06L9.06 8.5l1.47-1.47a.75.75 0 1 0-1.06-1.06L8 7.44z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Equal(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 11a.75.75 0 0 0 0-1.5h-11a.75.75 0 0 0 0 1.5zm0-4.5a.75.75 0 0 0 0-1.5h-11a.75.75 0 0 0 0 1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eraser(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m9.646 3.268l2.586 2.586a.914.914 0 0 1 0 1.292L8.72 10.66L4.84 6.78l3.513-3.512a.914.914 0 0 1 1.292 0ZM3.78 7.84L1.768 9.854a.914.914 0 0 0 0 1.292L3.12 12.5h3.76l.78-.78zm9.513.366L9 12.5h6.25a.75.75 0 0 1 0 1.5H2.5L.707 12.207a2.414 2.414 0 0 1 0-3.414l6.586-6.586a2.414 2.414 0 0 1 3.414 0l2.586 2.586a2.414 2.414 0 0 1 0 3.414Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationShape(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiExclamationShape0)"><path fill="currentColor" fill-rule="evenodd" d="m8.91 1.762l-.325 6.5a.25.25 0 0 1-.25.238h-.67a.25.25 0 0 1-.25-.238l-.325-6.5a.25.25 0 0 1 .25-.262h1.32a.25.25 0 0 1 .25.262m1.173 6.575l.325-6.5A1.75 1.75 0 0 0 8.66 0H7.34a1.75 1.75 0 0 0-1.748 1.837l.325 6.5A1.75 1.75 0 0 0 7.665 10h.67a1.75 1.75 0 0 0 1.748-1.663M8 12.5a1 1 0 1 1 0 2a1 1 0 0 1 0-2m2.5 1a2.5 2.5 0 1 0-5 0a2.5 2.5 0 0 0 5 0" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiExclamationShape0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.87 8.515L1.641 8l.229-.515a6.708 6.708 0 0 1 12.26 0l.228.515l-.229.515a6.708 6.708 0 0 1-12.259 0M.5 6.876l-.26.585a1.328 1.328 0 0 0 0 1.079l.26.584a8.208 8.208 0 0 0 15 0l.26-.584a1.328 1.328 0 0 0 0-1.08l-.26-.584a8.208 8.208 0 0 0-15 0M9.5 8a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0M11 8a3 3 0 1 1-6 0a3 3 0 0 1 6 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeSlash(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.03 1.97a.75.75 0 0 0-1.06 1.06l.83.83A8.206 8.206 0 0 0 .5 6.876l-.26.585a1.328 1.328 0 0 0 0 1.079l.26.585a8.208 8.208 0 0 0 11.434 3.87l1.036 1.035a.75.75 0 1 0 1.06-1.06zm7.788 9.908l-1.294-1.293a3 3 0 0 1-4.109-4.109L3.866 4.927A6.707 6.707 0 0 0 1.87 7.486L1.641 8l.23.515a6.708 6.708 0 0 0 8.947 3.363M6.55 7.611A1.502 1.502 0 0 0 8.389 9.45zm1.658-2.604l2.784 2.784a3 3 0 0 0-2.784-2.784m5.92 3.508a6.704 6.704 0 0 1-.915 1.496l1.065 1.066A8.203 8.203 0 0 0 15.5 9.125l.26-.585a1.328 1.328 0 0 0 0-1.08l-.26-.584A8.208 8.208 0 0 0 5.572 2.37L6.81 3.61a6.708 6.708 0 0 1 7.32 3.877l.228.514l-.228.515Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyesLookLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 8c0 1.75-.335 3.094-.816 3.944C4.721 12.764 4.217 13 3.75 13c-.467 0-.97-.237-1.434-1.056c-.272-.48-.496-1.116-.64-1.895c.283.289.647.451 1.074.451C3.826 10.5 4.5 9.47 4.5 8s-.674-2.5-1.75-2.5a1.47 1.47 0 0 0-1.075.45c.145-.778.37-1.415.64-1.894C2.78 3.236 3.284 3 3.75 3c.467 0 .97.237 1.434 1.056C5.665 4.906 6 6.25 6 8m1.5 0c0 3.822-1.445 6.5-3.75 6.5C1.445 14.5 0 11.822 0 8s1.445-6.5 3.75-6.5C6.055 1.5 7.5 4.178 7.5 8m7 0c0 1.75-.335 3.094-.816 3.944c-.463.82-.967 1.056-1.434 1.056c-.467 0-.97-.237-1.434-1.056c-.272-.48-.496-1.116-.64-1.895c.283.289.647.451 1.074.451C12.326 10.5 13 9.47 13 8s-.674-2.5-1.75-2.5a1.47 1.47 0 0 0-1.075.45c.145-.778.37-1.415.64-1.894C11.28 3.236 11.784 3 12.25 3c.467 0 .97.237 1.434 1.056c.481.85.816 2.195.816 3.944M16 8c0 3.822-1.445 6.5-3.75 6.5c-2.305 0-3.75-2.678-3.75-6.5s1.445-6.5 3.75-6.5C14.555 1.5 16 4.178 16 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyesLookRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.75 5.5c.427 0 .791.163 1.075.45c-.145-.778-.37-1.415-.64-1.894C4.72 3.236 4.216 3 3.75 3c-.467 0-.97.237-1.434 1.056C1.835 4.906 1.5 6.25 1.5 8s.335 3.094.816 3.944c.463.82.967 1.056 1.434 1.056c.467 0 .97-.237 1.434-1.056c.272-.48.496-1.116.64-1.895a1.47 1.47 0 0 1-1.074.451C3.674 10.5 3 9.47 3 8s.674-2.5 1.75-2.5M7.5 8c0 3.822-1.445 6.5-3.75 6.5C1.445 14.5 0 11.822 0 8s1.445-6.5 3.75-6.5C6.055 1.5 7.5 4.178 7.5 8m6.825 2.05c-.145.778-.37 1.415-.64 1.894c-.464.82-.968 1.056-1.435 1.056c-.467 0-.97-.237-1.434-1.056C10.335 11.094 10 9.75 10 8s.335-3.094.816-3.944C11.279 3.236 11.783 3 12.25 3c.467 0 .97.237 1.434 1.056c.272.48.496 1.116.64 1.895A1.47 1.47 0 0 0 13.25 5.5c-1.076 0-1.75 1.03-1.75 2.5s.674 2.5 1.75 2.5a1.47 1.47 0 0 0 1.075-.45M16 8c0 3.822-1.445 6.5-3.75 6.5c-2.305 0-3.75-2.678-3.75-6.5s1.445-6.5 3.75-6.5C14.555 1.5 16 4.178 16 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceAlien(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.072 11.749C12.447 10.539 13.5 8.915 13.5 7c0-1.376-.48-2.45-1.313-3.195C11.34 3.047 9.98 2.5 8 2.5s-3.34.547-4.187 1.305C2.98 4.55 2.5 5.625 2.5 7c0 1.915 1.053 3.539 2.428 4.749a9.484 9.484 0 0 0 2.01 1.36a5.5 5.5 0 0 0 .778.32c.17.054.262.067.284.07c.022-.003.113-.016.284-.07a5.5 5.5 0 0 0 .778-.32a9.484 9.484 0 0 0 2.01-1.36M8 15c1.5 0 7-3 7-8c0-3.5-2.5-6-7-6S1 3.5 1 7c0 5 5.5 8 7 8m1.25-5a.75.75 0 0 1-.75-.75c0-.842.345-1.553.922-2.041C9.986 6.73 10.73 6.5 11.5 6.5a.75.75 0 0 1 .75.75c0 .842-.345 1.553-.922 2.041c-.564.478-1.308.709-2.078.709M7.5 9.25a.75.75 0 0 1-.75.75c-.77 0-1.514-.231-2.078-.709c-.577-.488-.922-1.199-.922-2.041a.75.75 0 0 1 .75-.75c.77 0 1.514.231 2.078.709c.577.488.922 1.199.922 2.041" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceFun(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 13.5a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14M4.75 9.25a.75.75 0 0 1 .75-.75h5a.75.75 0 0 1 .75.75c0 .686-.43 1.319-.974 1.746c-.57.447-1.358.754-2.276.754c-.918 0-1.706-.307-2.276-.754c-.543-.427-.974-1.06-.974-1.746M10 7.5a.75.75 0 0 1-.75-.75v-1a.75.75 0 0 1 1.5 0v1a.75.75 0 0 1-.75.75m-4.75-.75a.75.75 0 0 0 1.5 0v-1a.75.75 0 0 0-1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceNeutral(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0m-5 3a.75.75 0 0 0 0-1.5H6A.75.75 0 0 0 6 11zm0-3a.75.75 0 0 1-.75-.75v-1a.75.75 0 0 1 1.5 0v1A.75.75 0 0 1 10 8m-4.75-.75a.75.75 0 0 0 1.5 0v-1a.75.75 0 0 0-1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceNeutralDashed(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1c-.372 0-.737.03-1.094.085a.75.75 0 0 0 .232 1.482a5.546 5.546 0 0 1 1.724 0a.75.75 0 0 0 .232-1.482A7.047 7.047 0 0 0 8 1M4.767 3.55a.75.75 0 1 0-.882-1.213a7.036 7.036 0 0 0-1.548 1.548a.75.75 0 0 0 1.213.882c.34-.466.75-.878 1.217-1.217m7.348-1.213a.75.75 0 1 0-.883 1.213c.467.34.879.75 1.218 1.217a.75.75 0 0 0 1.213-.882a7.037 7.037 0 0 0-1.548-1.548M2.567 7.138a.75.75 0 0 0-1.482-.232a7.047 7.047 0 0 0 0 2.188a.75.75 0 0 0 1.482-.232a5.546 5.546 0 0 1 0-1.724m12.348-.232a.75.75 0 1 0-1.482.232a5.54 5.54 0 0 1 0 1.724a.75.75 0 0 0 1.482.232a7.048 7.048 0 0 0 0-2.188M3.55 11.233a.75.75 0 1 0-1.213.882a7.037 7.037 0 0 0 1.548 1.548a.75.75 0 0 0 .882-1.213a5.533 5.533 0 0 1-1.217-1.217m10.113.882a.75.75 0 0 0-1.213-.883a5.53 5.53 0 0 1-1.217 1.218a.75.75 0 0 0 .882 1.213a7.037 7.037 0 0 0 1.548-1.548m-6.525 1.318a.75.75 0 0 0-.232 1.482a7.048 7.048 0 0 0 2.188 0a.75.75 0 1 0-.232-1.482a5.54 5.54 0 0 1-1.724 0M10 11a.75.75 0 0 0 0-1.5H6A.75.75 0 0 0 6 11zm0-3a.75.75 0 0 1-.75-.75v-1a.75.75 0 0 1 1.5 0v1A.75.75 0 0 1 10 8m-4.75-.75a.75.75 0 0 0 1.5 0v-1a.75.75 0 0 0-1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceRobot(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.5 10.071c0-2.582-1.426-4.853-3.633-6.087l1.039-1.87a.75.75 0 1 0-1.312-.728l-1.11 1.997A8.1 8.1 0 0 0 8 3a8.1 8.1 0 0 0-2.485.383l-1.11-1.997a.75.75 0 1 0-1.31.728l1.038 1.87C1.926 5.218.5 7.489.5 10.07c0 .813.169 1.603.614 2.294c.448.697 1.09 1.158 1.795 1.46C4.227 14.39 6.02 14.5 8 14.5s3.773-.11 5.09-.675c.707-.302 1.348-.763 1.796-1.46c.446-.691.614-1.481.614-2.294m-13.5 0C2 12.5 4 13 8 13s6-.5 6-2.929c0-3-2.5-5.571-6-5.571s-6 2.57-6 5.57Zm8.5 1.179a.75.75 0 0 1-.75-.75V9a.75.75 0 0 1 1.5 0v1.5a.75.75 0 0 1-.75.75m-5.75-.75a.75.75 0 0 0 1.5 0V9a.75.75 0 0 0-1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceSad(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0m-5.67 2.835a.75.75 0 1 0 1.34-.67C10.268 9.356 9.219 8.75 8 8.75s-2.267.606-2.67 1.415a.75.75 0 1 0 1.34.67c.097-.191.548-.585 1.33-.585s1.233.394 1.33.585M10 8a.75.75 0 0 1-.75-.75v-1a.75.75 0 0 1 1.5 0v1A.75.75 0 0 1 10 8m-4.75-.75a.75.75 0 0 0 1.5 0v-1a.75.75 0 0 0-1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceSmile(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M6.67 9.665a.75.75 0 0 0-1.34.67c.403.809 1.452 1.415 2.67 1.415s2.267-.606 2.67-1.415a.75.75 0 1 0-1.34-.67c-.097.191-.548.585-1.33.585s-1.233-.394-1.33-.585M10 8a.75.75 0 0 1-.75-.75v-1a.75.75 0 0 1 1.5 0v1A.75.75 0 0 1 10 8m-4.75-.75a.75.75 0 0 0 1.5 0v-1a.75.75 0 0 0-1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceSurprise(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0m-5.75-.75a.75.75 0 0 0 1.5 0v-1a.75.75 0 0 0-1.5 0zM6 8a.75.75 0 0 1-.75-.75v-1a.75.75 0 0 1 1.5 0v1A.75.75 0 0 1 6 8m2 4c1.37 0 2.5-.5 2.5-1.5S9.5 9 8 9s-2.5.5-2.5 1.5S6.5 12 8 12" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Factory(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.5 5.5v2.803l2.332-1.555L13.5 5.636V12.5a.5.5 0 0 1-.5.5H3a.5.5 0 0 1-.5-.5V3a.5.5 0 0 1 .5-.5h1.5A.5.5 0 0 1 5 3v5.446l2.382-1.733L9.5 5.173zM6.5 3v2.5L8 4.41l.903-.657A1.32 1.32 0 0 1 11 4.82v.68l1.02-.68l.463-.309l.429-.285A1.343 1.343 0 0 1 15 5.343V12.5a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V3a2 2 0 0 1 2-2h1.5a2 2 0 0 1 2 2m-1.75 7a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func File(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 13.5H5A1.5 1.5 0 0 1 3.5 12V4A1.5 1.5 0 0 1 5 2.5h2V5a3 3 0 0 0 3 3h2.5v4a1.5 1.5 0 0 1-1.5 1.5m1.303-7a1.503 1.503 0 0 0-.242-.318L8.818 2.939a1.5 1.5 0 0 0-.318-.242V5A1.5 1.5 0 0 0 10 6.5zm.818-1.379A3 3 0 0 1 14 7.243V12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V4a3 3 0 0 1 3-3h2.757a3 3 0 0 1 2.122.879L13.12 5.12Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileArrowDown(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 12a1.5 1.5 0 0 1-1.5 1.5H5A1.5 1.5 0 0 1 3.5 12V4A1.5 1.5 0 0 1 5 2.5h2.757a1.5 1.5 0 0 1 1.061.44l3.243 3.242a1.5 1.5 0 0 1 .439 1.06zm.621-6.879A3 3 0 0 1 14 7.243V12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V4a3 3 0 0 1 3-3h2.757a3 3 0 0 1 2.122.879L13.12 5.12ZM5.47 9.781A.75.75 0 0 1 6.53 8.72l.72.72V7.25a.75.75 0 0 1 1.5 0v2.19l.72-.72a.75.75 0 1 1 1.06 1.06l-2 2a.75.75 0 0 1-1.06 0l-2-2Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileArrowLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 13.5H5A1.5 1.5 0 0 1 3.5 12V4A1.5 1.5 0 0 1 5 2.5h2.757a1.5 1.5 0 0 1 1.061.44l3.243 3.242a1.5 1.5 0 0 1 .439 1.06V12a1.5 1.5 0 0 1-1.5 1.5m2.121-8.379A3 3 0 0 1 14 7.243V12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V4a3 3 0 0 1 3-3h2.757a3 3 0 0 1 2.122.879L13.12 5.12Zm-5.651 6.66a.75.75 0 0 0 1.06-1.061L7.81 10H10a.75.75 0 0 0 0-1.5H7.81l.72-.72a.75.75 0 0 0-1.06-1.06l-2 2a.75.75 0 0 0 0 1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileArrowRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 13.5H5A1.5 1.5 0 0 1 3.5 12V4A1.5 1.5 0 0 1 5 2.5h2.757a1.5 1.5 0 0 1 1.061.44l3.243 3.242a1.5 1.5 0 0 1 .439 1.06V12a1.5 1.5 0 0 1-1.5 1.5m2.121-8.379A3 3 0 0 1 14 7.243V12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V4a3 3 0 0 1 3-3h2.757a3 3 0 0 1 2.122.879L13.12 5.12Zm-4.59 6.66A.75.75 0 0 1 7.47 10.72l.72-.72H6a.75.75 0 0 1 0-1.5h2.19l-.72-.72a.75.75 0 0 1 1.06-1.06l2 2a.75.75 0 0 1 0 1.06l-2 2Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileArrowRightOut(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 7.14a3 3 0 0 0-.758-1.992L8.395 2.507A3 3 0 0 0 6.153 1.5H4a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h6.25a.75.75 0 0 0 0-1.5H4a1.5 1.5 0 0 1-1.5-1.5v-7A1.5 1.5 0 0 1 4 3h1.5v2a3 3 0 0 0 3 3h2.25a.75.75 0 0 0 .75-.75zM9.856 6.5H8.5A1.5 1.5 0 0 1 7 5V3.262c.1.068.192.15.274.241L9.62 6.144a1.5 1.5 0 0 1 .235.356ZM13.44 11l-.72.72a.75.75 0 1 0 1.061 1.06l2-2a.75.75 0 0 0 0-1.06l-2-2a.75.75 0 1 0-1.06 1.06l.72.72H6.75a.75.75 0 0 0 0 1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileArrowUp(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 12a1.5 1.5 0 0 1-1.5 1.5H5A1.5 1.5 0 0 1 3.5 12V4A1.5 1.5 0 0 1 5 2.5h2.757a1.5 1.5 0 0 1 1.061.44l3.243 3.242a1.5 1.5 0 0 1 .439 1.06zm.621-6.879A3 3 0 0 1 14 7.243V12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V4a3 3 0 0 1 3-3h2.757a3 3 0 0 1 2.122.879L13.12 5.12ZM5.47 8.72a.75.75 0 0 0 1.06 1.06l.72-.72v2.19a.75.75 0 0 0 1.5 0V9.06l.72.72a.75.75 0 0 0 1.06-1.06l-2-2a.75.75 0 0 0-1.06 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileCheck(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 13.5H5A1.5 1.5 0 0 1 3.5 12V4A1.5 1.5 0 0 1 5 2.5h2.757a1.5 1.5 0 0 1 1.061.44l3.243 3.242a1.5 1.5 0 0 1 .439 1.06V12a1.5 1.5 0 0 1-1.5 1.5m2.121-8.379A3 3 0 0 1 14 7.243V12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V4a3 3 0 0 1 3-3h2.757a3 3 0 0 1 2.122.879L13.12 5.12ZM10.85 7.65a.75.75 0 1 0-1.2-.9L7.469 9.658L6.28 8.47a.75.75 0 0 0-1.06 1.06l1.8 1.8a.75.75 0 0 0 1.13-.08z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileCode(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 13.5h6a1.5 1.5 0 0 0 1.5-1.5V7.243a1.5 1.5 0 0 0-.44-1.061L8.819 2.939a1.5 1.5 0 0 0-1.06-.439H5A1.5 1.5 0 0 0 3.5 4v8A1.5 1.5 0 0 0 5 13.5m9-6.257a3 3 0 0 0-.879-2.122L9.88 1.88A3 3 0 0 0 7.757 1H5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h6a3 3 0 0 0 3-3zM8.72 10.53a.75.75 0 0 1 0-1.06l.97-.97l-.97-.97a.75.75 0 0 1 1.06-1.06l1.5 1.5a.75.75 0 0 1 0 1.06l-1.5 1.5a.75.75 0 0 1-1.06 0m-1.44-3a.75.75 0 0 0-1.06-1.06l-1.5 1.5a.75.75 0 0 0 0 1.06l1.5 1.5a.75.75 0 1 0 1.06-1.06l-.97-.97z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileDollar(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 13.5H5A1.5 1.5 0 0 1 3.5 12V4A1.5 1.5 0 0 1 5 2.5h2.757a1.5 1.5 0 0 1 1.061.44l3.243 3.242a1.5 1.5 0 0 1 .439 1.06V12a1.5 1.5 0 0 1-1.5 1.5m2.121-8.379A3 3 0 0 1 14 7.243V12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V4a3 3 0 0 1 3-3h2.757a3 3 0 0 1 2.122.879L13.12 5.12Zm-6.878.388c.272-.203.596-.343.955-.423V4.75a.75.75 0 0 1 1.5 0v.311c.537.097.973.309 1.309.616a.75.75 0 1 1-1.014 1.106c-.134-.123-.412-.28-.993-.28c-.464 0-.728.11-.857.206c-.112.084-.147.172-.147.274c0 .014 0 .022.016.043a.628.628 0 0 0 .189.152c.23.135.573.242.952.32l.05.011c.394.081.956.197 1.42.434c.523.267 1.057.763 1.057 1.609c0 .75-.468 1.263-.98 1.55c-.302.17-.648.28-1.002.342v.306a.75.75 0 0 1-1.5 0v-.349a3.347 3.347 0 0 1-.915-.383c-.463-.281-.887-.704-1.064-1.226a.75.75 0 1 1 1.421-.48c.027.08.147.257.423.425c.262.16.595.265.937.265c.4 0 .748-.086.966-.209c.208-.116.214-.204.214-.241c0-.08-.008-.155-.239-.273c-.269-.137-.643-.218-1.093-.312c-.428-.089-.963-.235-1.408-.496c-.446-.263-.944-.735-.944-1.488c0-.587.268-1.115.747-1.474" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileExclamation(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 12a1.5 1.5 0 0 1-1.5 1.5H5A1.5 1.5 0 0 1 3.5 12V4A1.5 1.5 0 0 1 5 2.5h2.757a1.5 1.5 0 0 1 1.061.44l3.243 3.242a1.5 1.5 0 0 1 .439 1.06zm.621-6.879A3 3 0 0 1 14 7.243V12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V4a3 3 0 0 1 3-3h2.757a3 3 0 0 1 2.122.879L13.12 5.12ZM8 5a.75.75 0 0 1 .75.75V8.5a.75.75 0 0 1-1.5 0V5.75A.75.75 0 0 1 8 5m1 6.25a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileMagnifier(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 13.5H5A1.5 1.5 0 0 1 3.5 12V4A1.5 1.5 0 0 1 5 2.5h2.757a1.5 1.5 0 0 1 1.061.44l3.243 3.242a1.5 1.5 0 0 1 .439 1.06V12a1.5 1.5 0 0 1-1.5 1.5m2.121-8.379A3 3 0 0 1 14 7.243V12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V4a3 3 0 0 1 3-3h2.757a3 3 0 0 1 2.122.879L13.12 5.12ZM7.5 11.25A2.497 2.497 0 0 1 5 8.75a2.5 2.5 0 1 1 4.717 1.156l.813.814a.75.75 0 1 1-1.06 1.06l-.814-.813a2.49 2.49 0 0 1-1.156.283m1-2.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileMinus(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 13.5H5A1.5 1.5 0 0 1 3.5 12V4A1.5 1.5 0 0 1 5 2.5h2.757a1.5 1.5 0 0 1 1.061.44l3.243 3.242a1.5 1.5 0 0 1 .439 1.06V12a1.5 1.5 0 0 1-1.5 1.5m2.121-8.379A3 3 0 0 1 14 7.243V12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V4a3 3 0 0 1 3-3h2.757a3 3 0 0 1 2.122.879L13.12 5.12ZM5.75 8.25a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePlus(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 12a1.5 1.5 0 0 1-1.5 1.5H5A1.5 1.5 0 0 1 3.5 12V4A1.5 1.5 0 0 1 5 2.5h2.757a1.5 1.5 0 0 1 1.061.44l3.243 3.242a1.5 1.5 0 0 1 .439 1.06zm.621-6.879A3 3 0 0 1 14 7.243V12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V4a3 3 0 0 1 3-3h2.757a3 3 0 0 1 2.122.879L13.12 5.12ZM8.75 6.75a.75.75 0 0 0-1.5 0v1.5h-1.5a.75.75 0 0 0 0 1.5h1.5v1.5a.75.75 0 0 0 1.5 0v-1.5h1.5a.75.75 0 0 0 0-1.5h-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileQuestion(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 13.5H5A1.5 1.5 0 0 1 3.5 12V4A1.5 1.5 0 0 1 5 2.5h2.757a1.5 1.5 0 0 1 1.061.44l3.243 3.242a1.5 1.5 0 0 1 .439 1.06V12a1.5 1.5 0 0 1-1.5 1.5m2.121-8.379A3 3 0 0 1 14 7.243V12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V4a3 3 0 0 1 3-3h2.757a3 3 0 0 1 2.122.879L13.12 5.12Zm-6.533.388C6.978 5.19 7.488 5 8 5c1.092 0 2.125.736 2.125 1.947c0 .34-.07.678-.273 1.004c-.19.308-.467.554-.785.776c-.185.129-.269.211-.305.26a.197.197 0 0 0-.012.017A.75.75 0 0 1 7.25 9c0-.741.512-1.192.959-1.503c.236-.165.329-.273.368-.336c.028-.046.048-.097.048-.214a.367.367 0 0 0-.133-.292C8.398 6.575 8.235 6.5 8 6.5a.766.766 0 0 0-.463.17c-.126.104-.162.204-.162.277a.75.75 0 1 1-1.5 0c0-.618.32-1.116.713-1.438M9 11.25a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileRuble(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 13.5H5A1.5 1.5 0 0 1 3.5 12V4A1.5 1.5 0 0 1 5 2.5h2.757a1.5 1.5 0 0 1 1.061.44l3.243 3.242a1.5 1.5 0 0 1 .439 1.06V12a1.5 1.5 0 0 1-1.5 1.5m2.121-8.379A3 3 0 0 1 14 7.243V12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V4a3 3 0 0 1 3-3h2.757a3 3 0 0 1 2.122.879L13.12 5.12ZM6.75 4.5a.75.75 0 0 0-.75.75v2h-.25a.75.75 0 0 0 0 1.5H6v.75h-.25a.75.75 0 0 0 0 1.5H6v.25a.75.75 0 0 0 1.5 0V11h.75a.75.75 0 0 0 0-1.5H7.5v-.75h.875a2.125 2.125 0 0 0 0-4.25zm1.625 2.75H7.5V6h.875a.625.625 0 1 1 0 1.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileText(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 13.5h6a1.5 1.5 0 0 0 1.5-1.5V7.243a1.5 1.5 0 0 0-.44-1.061L8.819 2.939a1.5 1.5 0 0 0-1.06-.439H5A1.5 1.5 0 0 0 3.5 4v8A1.5 1.5 0 0 0 5 13.5m9-6.257a3 3 0 0 0-.879-2.122L9.88 1.88A3 3 0 0 0 7.757 1H5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h6a3 3 0 0 0 3-3zM5 8.25a.75.75 0 0 1 .75-.75h4.5a.75.75 0 0 1 0 1.5h-4.5A.75.75 0 0 1 5 8.25m.75 2.25a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileXmark(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 12a1.5 1.5 0 0 1-1.5 1.5H5A1.5 1.5 0 0 1 3.5 12V4A1.5 1.5 0 0 1 5 2.5h2.757a1.5 1.5 0 0 1 1.061.44l3.243 3.242a1.5 1.5 0 0 1 .439 1.06zm.621-6.879A3 3 0 0 1 14 7.243V12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V4a3 3 0 0 1 3-3h2.757a3 3 0 0 1 2.122.879L13.12 5.12ZM6.781 6.72A.75.75 0 0 0 5.72 7.78L6.94 9l-1.22 1.22a.75.75 0 1 0 1.06 1.06L8 10.06l1.22 1.22a.75.75 0 1 0 1.06-1.06L9.06 9l1.22-1.22a.75.75 0 1 0-1.06-1.06L8 7.94z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileZipper(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 13.5H5A1.5 1.5 0 0 1 3.5 12V4A1.5 1.5 0 0 1 5 2.5h.5v.75c0 .414.336.75.75.75H7v2h-.75a.75.75 0 0 0-.75.75v.5c0 .414.336.75.75.75H7v2h-.75a.75.75 0 0 0-.75.75v.5c0 .414.336.75.75.75H7v-2h.75a.75.75 0 0 0 .75-.75v-.5A.75.75 0 0 0 7.75 8H7V6h.75a.75.75 0 0 0 .75-.75v-.5A.75.75 0 0 0 7.75 4H7V2.5h.757a1.5 1.5 0 0 1 1.061.44l3.243 3.242a1.5 1.5 0 0 1 .439 1.06V12a1.5 1.5 0 0 1-1.5 1.5m2.121-8.379A3 3 0 0 1 14 7.243V12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V4a3 3 0 0 1 3-3h2.757a3 3 0 0 1 2.122.879L13.12 5.12Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Files(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 10.5h4A1.5 1.5 0 0 0 13.5 9V7H12a3 3 0 0 1-3-3V2.5H8A1.5 1.5 0 0 0 6.5 4v5A1.5 1.5 0 0 0 8 10.5m5.06-5.318c.096.096.178.203.243.318H12A1.5 1.5 0 0 1 10.5 4V2.697c.115.065.223.147.318.242zM15 6.242a3 3 0 0 0-.879-2.12L11.88 1.878A3 3 0 0 0 9.757 1H8a3 3 0 0 0-3 3H4a3 3 0 0 0-3 3v5a3 3 0 0 0 3 3h4a3 3 0 0 0 3-3h1a3 3 0 0 0 3-3V6.243ZM9.5 12H8a3 3 0 0 1-3-3V5.5H4A1.5 1.5 0 0 0 2.5 7v5A1.5 1.5 0 0 0 4 13.5h4A1.5 1.5 0 0 0 9.5 12" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filmstrip(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 3.5h.5A1.5 1.5 0 0 1 13.5 5v.5h-2zm0 3.5v2h2V7zM15 7v4a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3V5a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3zm-1.5 3.5h-2v2h.5a1.5 1.5 0 0 0 1.5-1.5zm-3.5-7H6v9h4zm-5.5 9v-2h-2v.5A1.5 1.5 0 0 0 4 12.5zm0-5.5v2h-2V7zm0-1.5h-2V5A1.5 1.5 0 0 1 4 3.5h.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fingerprint(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.964 1.79a.75.75 0 0 1-.51.93A5.513 5.513 0 0 0 2.72 6.455a.75.75 0 0 1-1.44-.421A7.013 7.013 0 0 1 6.034 1.28a.75.75 0 0 1 .93.51m2.072 0a.75.75 0 0 1 .93-.51a7.003 7.003 0 0 1 4.176 10.08a.75.75 0 1 1-1.315-.72a5.503 5.503 0 0 0-3.281-7.92a.75.75 0 0 1-.51-.93M1.79 9.036a.75.75 0 0 1 .93.51a5.503 5.503 0 0 0 7.92 3.28a.75.75 0 1 1 .72 1.316A7.003 7.003 0 0 1 1.28 9.966a.75.75 0 0 1 .51-.93M8 5.5a2.5 2.5 0 0 0-1.153 4.719a.75.75 0 0 1-.694 1.33A4 4 0 1 1 12 8a.75.75 0 0 1-1.5 0A2.5 2.5 0 0 0 8 5.5m.742 2.392a.75.75 0 1 0-1.484.216c.252 1.726 1.493 3.001 3.18 3.14a.75.75 0 1 0 .123-1.495c-.918-.076-1.657-.751-1.819-1.861" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.47 3.588a4.447 4.447 0 0 0-4.15-.224a.55.55 0 0 0-.32.499v5.533a6.25 6.25 0 0 1 5.547.439l.344.207a4.018 4.018 0 0 0 3.865.148a.442.442 0 0 0 .244-.395V4.182a6.264 6.264 0 0 1-5.386-.508zm5.957 7.944a5.518 5.518 0 0 1-5.307-.204l-.345-.207a4.75 4.75 0 0 0-4.314-.293L3 11.026v3.255a.75.75 0 0 1-1.5 0V3.863c0-.8.465-1.526 1.19-1.861a5.947 5.947 0 0 1 5.552.3l.144.086a4.761 4.761 0 0 0 4.447.24l.603-.278a.75.75 0 0 1 1.064.681v6.764c0 .735-.416 1.408-1.073 1.737" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flame(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.452 6.864l1.13-2.173a31.715 31.715 0 0 1 1.872-3.095c.964 1.045 1.906 2.3 2.612 3.622c.748 1.402 1.184 2.789 1.184 4.032c0 1.427-.904 2.83-2.153 3.613c.058-.265.09-.553.09-.863c0-1.255-.674-2.336-1.143-2.963a8.82 8.82 0 0 0-1.01-1.125l-.024-.02l-.008-.008L9 7.88l-.001-.001C8.996 7.88 8.996 7.878 8 9a7.03 7.03 0 0 0 .984 1.133c.37.534.704 1.2.704 1.867c0 .77-.313 1.276-.618 1.587c-.159.162-.279.38-.314.6a.786.786 0 0 0 0 .264a.694.694 0 0 0 .06.182c.113.225.343.37.594.35c2.836-.235 5.34-2.87 5.34-5.733c0-3.149-2.177-6.538-4.357-8.845A1.313 1.313 0 0 0 9.435 0A1.32 1.32 0 0 0 8.35.556A33.486 33.486 0 0 0 6.25 4l-.955-1.337a.986.986 0 0 0-1.589-.018C2.62 4.123 1.25 6.249 1.25 9.25c0 2.863 2.504 5.498 5.34 5.733c.25.02.481-.125.593-.35a.672.672 0 0 0 .06-.182a.786.786 0 0 0 .001-.263a1.145 1.145 0 0 0-.314-.601c-.305-.31-.617-.817-.617-1.587c0-.666.333-1.333.703-1.867l.09-.128C7.544 9.405 8 9 8 9l-.997-1.12H7l-.003.003l-.008.007l-.024.021l-.073.07a8.827 8.827 0 0 0-.937 1.056c-.47.626-1.143 1.707-1.143 2.962c0 .31.033.598.09.863C3.654 12.08 2.75 10.677 2.75 9.25c0-2.171.847-3.812 1.745-5.126l.534.748zM8 9l.997-1.121L8 6.993l-.997.886z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flask(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.494 13.2c.837-.482 1.006-.946 1.006-1.2c0-.35-.105-.692-.303-.981L9.072 6.435a1.854 1.854 0 0 1-.322-1.044V2.909C8.518 2.968 8.265 3 8 3s-.518-.032-.75-.09v2.48c0 .373-.112.737-.322 1.045L3.803 11.02c-.198.289-.303.63-.303.981c0 .254.169.718 1.006 1.2c.813.468 2.043.8 3.494.8s2.68-.332 3.494-.8ZM8 .5c2 0 2.25 1 2.25 1.5v3.39c0 .072.021.141.062.2l3.125 4.584c.367.538.563 1.175.563 1.826c0 2-2.686 3.5-6 3.5S2 14 2 12c0-.652.196-1.288.563-1.826L5.69 5.59a.354.354 0 0 0 .061-.2V2C5.75 1.5 6 .5 8 .5m.084 7.626a.75.75 0 0 1 1.04.208l1.5 2.25a.75.75 0 1 1-1.248.832l-1.5-2.25a.75.75 0 0 1 .208-1.04" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FloppyDisk(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 11.5A1.5 1.5 0 0 0 4.5 13v-2.5a2 2 0 0 1 2-2h3a2 2 0 0 1 2 2V13a1.5 1.5 0 0 0 1.5-1.5V6.036a1 1 0 0 0-.293-.708l-2.035-2.035A1 1 0 0 0 9.964 3H6v1a.5.5 0 0 0 .5.5h3a.75.75 0 0 1 0 1.5h-3a2 2 0 0 1-2-2V3A1.5 1.5 0 0 0 3 4.5zm-1.5 0a3 3 0 0 0 3 3h7a3 3 0 0 0 3-3V6.036a2.5 2.5 0 0 0-.732-1.768l-2.036-2.036A2.5 2.5 0 0 0 9.964 1.5H4.5a3 3 0 0 0-3 3zm8.5-1V13H6v-2.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 .5.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.44 4.06l.439.44H12.5A1.5 1.5 0 0 1 14 6v5a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 2 11V4.5A1.5 1.5 0 0 1 3.5 3h1.257a1.5 1.5 0 0 1 1.061.44l.621.62ZM.5 4.5a3 3 0 0 1 3-3h1.257a3 3 0 0 1 2.122.879L7.5 3h5a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3h-9a3 3 0 0 1-3-3zm4.25 2a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderArrowDown(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.44 4.06l.439.44H12.5A1.5 1.5 0 0 1 14 6v5a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 2 11V4.5A1.5 1.5 0 0 1 3.5 3h1.257a1.5 1.5 0 0 1 1.061.44l.621.62ZM.5 4.5a3 3 0 0 1 3-3h1.257a3 3 0 0 1 2.122.879L7.5 3h5a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3h-9a3 3 0 0 1-3-3zm10.03 4.53a.75.75 0 1 0-1.06-1.06l-.72.72V6.5a.75.75 0 0 0-1.5 0v2.19l-.72-.72a.75.75 0 0 0-1.06 1.06l2 2a.75.75 0 0 0 1.06 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderArrowLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.44 4.06l.439.44H12.5A1.5 1.5 0 0 1 14 6v5a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 2 11V4.5A1.5 1.5 0 0 1 3.5 3h1.257a1.5 1.5 0 0 1 1.061.44l.621.62ZM.5 4.5a3 3 0 0 1 3-3h1.257a3 3 0 0 1 2.122.879L7.5 3h5a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3h-9a3 3 0 0 1-3-3zm6.97 6.53a.75.75 0 1 0 1.06-1.06l-.72-.72H10a.75.75 0 0 0 0-1.5H7.81l.72-.72a.75.75 0 0 0-1.06-1.06l-2 2a.75.75 0 0 0 0 1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderArrowRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.44 4.06l.439.44H12.5A1.5 1.5 0 0 1 14 6v5a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 2 11V4.5A1.5 1.5 0 0 1 3.5 3h1.257a1.5 1.5 0 0 1 1.061.44l.621.62ZM.5 4.5a3 3 0 0 1 3-3h1.257a3 3 0 0 1 2.122.879L7.5 3h5a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3h-9a3 3 0 0 1-3-3zm8.03 1.47a.75.75 0 0 0-1.06 1.06l.72.72H6a.75.75 0 1 0 0 1.5h2.19l-.72.72a.75.75 0 1 0 1.06 1.06l2-2a.75.75 0 0 0 0-1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderArrowUp(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.44 4.06l.439.44H12.5A1.5 1.5 0 0 1 14 6v5a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 2 11V4.5A1.5 1.5 0 0 1 3.5 3h1.257a1.5 1.5 0 0 1 1.061.44l.621.62ZM.5 4.5a3 3 0 0 1 3-3h1.257a3 3 0 0 1 2.122.879L7.5 3h5a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3h-9a3 3 0 0 1-3-3zm4.97 3.47a.75.75 0 0 0 1.06 1.06l.72-.72v2.19a.75.75 0 0 0 1.5 0V8.31l.72.72a.75.75 0 1 0 1.06-1.06l-2-2a.75.75 0 0 0-1.06 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderArrowUpIn(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.44 3.56l.439.44H12.5A1.5 1.5 0 0 1 14 5.5v5a1.5 1.5 0 0 1-1.5 1.5h-1.75a.75.75 0 0 0 0 1.5h1.75a3 3 0 0 0 3-3v-5a3 3 0 0 0-3-3h-5l-.621-.621A3 3 0 0 0 4.757 1H3.5a3 3 0 0 0-3 3v6.5a3 3 0 0 0 3 3h1.75a.75.75 0 0 0 0-1.5H3.5A1.5 1.5 0 0 1 2 10.5V4a1.5 1.5 0 0 1 1.5-1.5h1.257a1.5 1.5 0 0 1 1.061.44l.621.62Zm2.31 5l.72.72a.75.75 0 1 0 1.06-1.06l-2-2a.75.75 0 0 0-1.06 0l-2 2a.75.75 0 0 0 1.06 1.06l.72-.72v6.69a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderCheck(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.44 4.06l.439.44H12.5A1.5 1.5 0 0 1 14 6v5a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 2 11V4.5A1.5 1.5 0 0 1 3.5 3h1.257a1.5 1.5 0 0 1 1.061.44l.621.62ZM.5 4.5a3 3 0 0 1 3-3h1.257a3 3 0 0 1 2.122.879L7.5 3h5a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3h-9a3 3 0 0 1-3-3zm10.35 2.45a.75.75 0 1 0-1.2-.9L7.469 8.958L6.28 7.77a.75.75 0 0 0-1.06 1.06l1.8 1.8a.75.75 0 0 0 1.13-.08z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderCode(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.44 4.06l.439.44H12.5A1.5 1.5 0 0 1 14 6v5a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 2 11V4.5A1.5 1.5 0 0 1 3.5 3h1.257a1.5 1.5 0 0 1 1.061.44l.621.62ZM.5 4.5a3 3 0 0 1 3-3h1.257a3 3 0 0 1 2.122.879L7.5 3h5a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3h-9a3 3 0 0 1-3-3zm8.182 1.273a.75.75 0 0 1 .546.91l-1 4a.75.75 0 0 1-1.455-.365l1-4a.75.75 0 0 1 .909-.545M4.97 10.03a.75.75 0 0 0 1.06-1.06l-.47-.47l.47-.47a.75.75 0 0 0-1.06-1.06l-1 1a.75.75 0 0 0 0 1.06zm5-1.06a.75.75 0 1 0 1.06 1.06l1-1a.75.75 0 0 0 0-1.06l-1-1a.75.75 0 1 0-1.06 1.06l.47.47z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderExclamation(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.879 4.5l-.44-.44l-.621-.62A1.5 1.5 0 0 0 4.758 3H3.5A1.5 1.5 0 0 0 2 4.5V11a1.5 1.5 0 0 0 1.5 1.5h9A1.5 1.5 0 0 0 14 11V6a1.5 1.5 0 0 0-1.5-1.5zM3.5 1.5a3 3 0 0 0-3 3V11a3 3 0 0 0 3 3h9a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3h-5l-.621-.621A3 3 0 0 0 4.757 1.5zm5.5 9a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-.25-4.25a.75.75 0 0 0-1.5 0v1.5a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 1.5a3 3 0 0 0-3 3V11a3 3 0 0 0 3 3h9a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3h-5l-.621-.621A3 3 0 0 0 4.757 1.5zm1.25 5a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderFlows(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.44 4.06l.439.44H12.5A1.5 1.5 0 0 1 14 6v5a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 2 11V4.5A1.5 1.5 0 0 1 3.5 3h1.257a1.5 1.5 0 0 1 1.061.44l.621.62ZM.5 4.5a3 3 0 0 1 3-3h1.257a3 3 0 0 1 2.122.879L7.5 3h5a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3h-9a3 3 0 0 1-3-3zm8.25 2.75A.25.25 0 0 1 9 7h1a.75.75 0 0 0 0-1.5H9a1.75 1.75 0 0 0-1.75 1.75v.5H6a.75.75 0 1 0 0 1.5h1.25v.5c0 .966.784 1.75 1.75 1.75h1a.75.75 0 0 0 0-1.5H9a.25.25 0 0 1-.25-.25v-.5H10a.75.75 0 0 0 0-1.5H8.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderHouse(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.44 4.06l.439.44H12.5A1.5 1.5 0 0 1 14 6v5a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 2 11V4.5A1.5 1.5 0 0 1 3.5 3h1.257a1.5 1.5 0 0 1 1.061.44l.621.62ZM.5 4.5a3 3 0 0 1 3-3h1.257a3 3 0 0 1 2.122.879L7.5 3h5a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3h-9a3 3 0 0 1-3-3zM8 7.453l1.5 1.25V9.5h-3v-.797zM5 8.468a1 1 0 0 1 .36-.768l2-1.667a1 1 0 0 1 1.28 0l2 1.667a1 1 0 0 1 .36.768V10a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderKeyhole(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.879 4.5l-.44-.44l-.621-.62A1.5 1.5 0 0 0 4.758 3H3.5A1.5 1.5 0 0 0 2 4.5V11a1.5 1.5 0 0 0 1.5 1.5h9A1.5 1.5 0 0 0 14 11V6a1.5 1.5 0 0 0-1.5-1.5zM3.5 1.5a3 3 0 0 0-3 3V11a3 3 0 0 0 3 3h9a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3h-5l-.621-.621A3 3 0 0 0 4.757 1.5zm5.25 7.3a1.5 1.5 0 1 0-1.5 0v1.45a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderLock(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.44 4.06l.439.44H12.5A1.5 1.5 0 0 1 14 6v5a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 2 11V4.5A1.5 1.5 0 0 1 3.5 3h1.257a1.5 1.5 0 0 1 1.061.44l.621.62ZM.5 4.5a3 3 0 0 1 3-3h1.257a3 3 0 0 1 2.122.879L7.5 3h5a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3h-9a3 3 0 0 1-3-3zm8 2.75v.25h-1v-.25a.5.5 0 0 1 1 0m1.6.25H10v-.25a2 2 0 1 0-4 0v.25h-.1a1.4 1.4 0 0 0-1.4 1.4v1.2a1.4 1.4 0 0 0 1.4 1.4h4.2a1.4 1.4 0 0 0 1.4-1.4V8.9a1.4 1.4 0 0 0-1.4-1.4M6 9v1h4V9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderMagnifier(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.879 4.5l-.44-.44l-.621-.62A1.5 1.5 0 0 0 4.758 3H3.5A1.5 1.5 0 0 0 2 4.5V11a1.5 1.5 0 0 0 1.5 1.5h9A1.5 1.5 0 0 0 14 11V6a1.5 1.5 0 0 0-1.5-1.5zM3.5 1.5a3 3 0 0 0-3 3V11a3 3 0 0 0 3 3h9a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3h-5l-.621-.621A3 3 0 0 0 4.757 1.5zm6.217 7.656a2.5 2.5 0 1 0-1.06 1.06l.813.814a.75.75 0 1 0 1.06-1.06zM7.5 9a1 1 0 1 1 0-2a1 1 0 0 1 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpen(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.379 4.5l-.44-.44l-.621-.62A1.5 1.5 0 0 0 4.258 3H3a1.5 1.5 0 0 0-1.5 1.5v5.25l1.376-2.293A3 3 0 0 1 5.45 6h7.05A1.5 1.5 0 0 0 11 4.5zM14 6.026V6a3 3 0 0 0-3-3H7l-.621-.621A3 3 0 0 0 4.257 1.5H3a3 3 0 0 0-3 3V11a3 3 0 0 0 3 3h8.301a3 3 0 0 0 2.573-1.457l1.791-2.985A2.349 2.349 0 0 0 14 6.026M10 12.5h1.301a1.5 1.5 0 0 0 1.287-.728l1.791-2.986l1.286.772l-1.286-.772a.85.85 0 0 0-.728-1.286H5.449a1.5 1.5 0 0 0-1.287.728l-1.791 2.986a.85.85 0 0 0 .728 1.286z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpenFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 4.5a3 3 0 0 1 3-3h1.257a3 3 0 0 1 2.122.879L7 3h4a3 3 0 0 1 3 3H5.449a3 3 0 0 0-2.573 1.457L1.43 9.867a.75.75 0 0 0 1.287.772l1.446-2.41A1.5 1.5 0 0 1 5.45 7.5h9.376a1.175 1.175 0 0 1 1.008 1.779l-1.96 3.264A3 3 0 0 1 11.302 14H2.94A3 3 0 0 1 0 11z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderPlus(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.44 4.06l.439.44H12.5A1.5 1.5 0 0 1 14 6v5a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 2 11V4.5A1.5 1.5 0 0 1 3.5 3h1.257a1.5 1.5 0 0 1 1.061.44l.621.62ZM.5 4.5a3 3 0 0 1 3-3h1.257a3 3 0 0 1 2.122.879L7.5 3h5a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3h-9a3 3 0 0 1-3-3zm8.25 2a.75.75 0 0 0-1.5 0v1.25H6a.75.75 0 0 0 0 1.5h1.25v1.25a.75.75 0 0 0 1.5 0V9.25H10a.75.75 0 0 0 0-1.5H8.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderTree(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 13.5v-3A.5.5 0 0 1 8 10h.672a.5.5 0 0 1 .353.146l.414.415l.44.439H13a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5H8a.5.5 0 0 1-.5-.5M6 13H4.5a3 3 0 0 1-3-3V1.25a.75.75 0 0 1 1.5 0V2a1.5 1.5 0 0 0 1.5 1.5H6v-1a2 2 0 0 1 2-2h.672a2 2 0 0 1 1.414.586l.414.414H13a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V5H4.5A2.986 2.986 0 0 1 3 4.599V10a1.5 1.5 0 0 0 1.5 1.5H6v-1a2 2 0 0 1 2-2h.672a2 2 0 0 1 1.414.586l.414.414H13a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2zm1.5-8.75V5.5A.5.5 0 0 0 8 6h5a.5.5 0 0 0 .5-.5v-2A.5.5 0 0 0 13 3H9.879l-.44-.44l-.414-.414A.5.5 0 0 0 8.672 2H8a.5.5 0 0 0-.5.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folders(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m8.879 4l-.44-.44l-.767-.767a1 1 0 0 0-.708-.293H6A1.5 1.5 0 0 0 4.5 4v5A1.5 1.5 0 0 0 6 10.5h7A1.5 1.5 0 0 0 14.5 9V5.5A1.5 1.5 0 0 0 13 4zM6 1a3 3 0 0 0-3 3a3 3 0 0 0-3 3v5a3 3 0 0 0 3 3h7a3 3 0 0 0 3-3a3 3 0 0 0 3-3V5.5a3 3 0 0 0-3-3H9.5l-.768-.768A2.5 2.5 0 0 0 6.964 1zM1.5 7A1.5 1.5 0 0 1 3 5.5V9a3 3 0 0 0 3 3h5.5a1.5 1.5 0 0 1-1.5 1.5H3A1.5 1.5 0 0 1 1.5 12zm10.75 0a.75.75 0 0 0 0-1.5h-5.5a.75.75 0 0 0 0 1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Font(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 2.25c-.618 0-1.169.39-1.373.974l-3.335 9.528a.75.75 0 0 0 1.416.496L5.845 10h4.31l1.137 3.248a.75.75 0 0 0 1.416-.496L9.373 3.224A1.455 1.455 0 0 0 8 2.25M9.63 8.5L8 3.842L6.37 8.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FontCase(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiFontCase0)"><path fill="currentColor" fill-rule="evenodd" d="M4.75 2.25c-.618 0-1.169.39-1.373.974L.042 12.752a.75.75 0 0 0 1.416.496L2.595 10h4.31l1.137 3.248a.75.75 0 0 0 1.416-.496L6.123 3.224A1.455 1.455 0 0 0 4.75 2.25M6.38 8.5L4.75 3.842L3.12 8.5zm5.135 2.996c0-.223.28-.746 1.152-.746H14.4c-.294 1.024-1.178 1.5-1.9 1.5c-.45 0-.677-.134-.792-.249a.706.706 0 0 1-.193-.505m2.985.754V13a.75.75 0 1 0 1.5 0v-3c0-1.117-.28-2.065-.873-2.744c-.606-.692-1.453-1.006-2.377-1.006c-.53 0-.946.07-1.306.195c-.338.117-.6.274-.804.396l-.025.015a.75.75 0 1 0 .77 1.288c.22-.132.365-.217.55-.281c.178-.062.423-.113.815-.113c.576 0 .978.186 1.248.494c.191.218.354.543.44 1.006h-1.771c-1.462 0-2.658.977-2.652 2.254c.003.542.191 1.116.632 1.557c.447.448 1.085.689 1.853.689c1 0 1.75-.75 1.75-1.5z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiFontCase0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FontCursor(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 2.25a.75.75 0 0 1 .75-.75c.576 0 1.102.217 1.5.573a2.242 2.242 0 0 1 1.5-.573a.75.75 0 0 1 0 1.5a.75.75 0 0 0-.75.75v8.75c0 .414.336.75.75.75a.75.75 0 0 1 0 1.5a2.242 2.242 0 0 1-1.5-.573a2.242 2.242 0 0 1-1.5.573a.75.75 0 0 1 0-1.5a.75.75 0 0 0 .75-.75V3.75a.75.75 0 0 0-.75-.75a.75.75 0 0 1-.75-.75M5.75 3c-.618 0-1.168.39-1.374.972L1.543 12a.75.75 0 1 0 1.414.5L3.84 10h3.82l.883 2.5a.75.75 0 1 0 1.414-.5L7.124 3.972A1.457 1.457 0 0 0 5.75 3m1.381 5.5L5.75 4.587L4.369 8.5H7.13Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Frame(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.25 1a.75.75 0 0 1 .75.75V3h1.25a.75.75 0 0 1 0 1.5H13v7h1.25a.75.75 0 0 1 0 1.5H13v1.25a.75.75 0 0 1-1.5 0V13h-7v1.25a.75.75 0 0 1-1.5 0V13H1.75a.75.75 0 0 1 0-1.5H3v-7H1.75a.75.75 0 0 1 0-1.5H3V1.75a.75.75 0 0 1 1.5 0V3h7V1.75a.75.75 0 0 1 .75-.75M4.5 4.5v7h7v-7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Frames(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiFrames0)"><path fill="currentColor" fill-rule="evenodd" d="M13.25.75a.75.75 0 0 1 .75.75V2h.5a.75.75 0 0 1 0 1.5H14v5h.5a.75.75 0 0 1 0 1.5H14v.5a.75.75 0 0 1-1.5 0V10H10v2.5h.5a.75.75 0 0 1 0 1.5H10v.5a.75.75 0 0 1-1.5 0V14h-5v.5a.75.75 0 0 1-1.5 0V14h-.5a.75.75 0 0 1 0-1.5H2v-5h-.5a.75.75 0 0 1 0-1.5H2v-.5a.75.75 0 0 1 1.5 0V6H6V3.5h-.5a.75.75 0 0 1 0-1.5H6v-.5a.75.75 0 0 1 1.5 0V2h5v-.5a.75.75 0 0 1 .75-.75M7.5 6h1v-.5a.75.75 0 0 1 1.5 0V6h.5a.75.75 0 0 1 0 1.5H10v1h2.5v-5h-5zm-4 6.5v-5h5v5z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiFrames0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Function(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiFunction0)"><path fill="currentColor" fill-rule="evenodd" d="M4.312 4.29a.764.764 0 0 1 1.103-.62a.75.75 0 1 0 .67-1.34a2.264 2.264 0 0 0-3.268 1.836L2.706 5.5H1.75a.75.75 0 0 0 0 1.5h.83l-.392 4.71a.764.764 0 0 1-1.103.62a.75.75 0 0 0-.67 1.34a2.264 2.264 0 0 0 3.268-1.836L4.086 7H5.25a.75.75 0 1 0 0-1.5H4.21zm6.014 2.23a.75.75 0 0 0-1.152.96l.85 1.02l-.85 1.02a.75.75 0 0 0 1.152.96L11 9.672l.674.808a.75.75 0 0 0 1.152-.96l-.85-1.02l.85-1.02a.75.75 0 0 0-1.152-.96L11 7.328zM8.02 4.55a.75.75 0 0 1 .43.969l-.145.378a7.25 7.25 0 0 0 0 5.205l.145.378a.75.75 0 0 1-1.4.539l-.145-.378a8.75 8.75 0 0 1 0-6.282l.145-.378a.75.75 0 0 1 .97-.431m5.961 0a.75.75 0 0 1 .97.43l.145.379a8.75 8.75 0 0 1 0 6.282l-.146.378a.75.75 0 1 1-1.4-.538l.146-.379a7.25 7.25 0 0 0 0-5.205l-.146-.378a.75.75 0 0 1 .431-.97Z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiFunction0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Funnel(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 4c0 .174-.071.513-.885.888C10.8 5.263 9.538 5.5 8 5.5c-1.538 0-2.799-.237-3.615-.612C3.57 4.513 3.5 4.174 3.5 4c0-.174.071-.513.885-.888C5.2 2.737 6.462 2.5 8 2.5c1.538 0 2.799.237 3.615.612c.814.375.885.714.885.888m-1.448 2.66C10.158 6.888 9.114 7 8 7s-2.158-.113-3.052-.34l1.98 2.905c.21.308.322.672.322 1.044v3.37c.06.012.141.021.25.021c.422 0 .749-.14.95-.316c.185-.162.3-.38.3-.684v-2.39c0-.373.112-.737.322-1.045zM8 1c3.314 0 6 1 6 3a3.24 3.24 0 0 1-.563 1.826l-3.125 4.584a.354.354 0 0 0-.062.2V13c0 1.5-1.25 2.5-2.75 2.5s-1.75-1-1.75-1v-3.89a.354.354 0 0 0-.061-.2L2.563 5.826A3.242 3.242 0 0 1 2 4c0-2 2.686-3 6-3m-.88 12.936c-.01-.005-.014-.009-.013-.01z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FunnelXmark(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.615 4.888c.814-.375.885-.714.885-.888c0-.174-.071-.513-.885-.888C9.8 2.737 8.538 2.5 7 2.5c-1.538 0-2.799.237-3.615.612c-.814.375-.885.714-.885.888c0 .174.071.513.885.888C4.2 5.263 5.462 5.5 7 5.5c1.538 0 2.799-.237 3.615-.612M7 7c1.114 0 2.158-.113 3.052-.34l-1.98 2.905c-.21.308-.322.672-.322 1.044V13a.861.861 0 0 1-.3.684c-.201.175-.528.316-.95.316a1.3 1.3 0 0 1-.25-.02v-3.37c0-.373-.112-.737-.322-1.045L3.948 6.66C4.842 6.887 5.886 7 7 7m6-3c0-2-2.686-3-6-3S1 2 1 4c0 .652.196 1.288.563 1.826L4.69 10.41c.04.059.061.128.061.2v3.89s.25 1 1.75 1s2.75-1 2.75-2.5v-2.39c0-.072.021-.141.061-.2l3.126-4.584A3.242 3.242 0 0 0 13 4m-6.88 9.936c-.01-.005-.014-.009-.013-.01zm7.13-2.247l-1.22-1.22a.75.75 0 1 0-1.06 1.061l1.22 1.22l-1.22 1.22a.75.75 0 1 0 1.06 1.06l1.22-1.22l1.22 1.22a.75.75 0 1 0 1.06-1.06l-1.22-1.22l1.22-1.22a.75.75 0 1 0-1.06-1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gear(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiGear0)"><path fill="currentColor" fill-rule="evenodd" d="M7.199 2H8.8a.2.2 0 0 1 .2.2c0 1.808 1.958 2.939 3.524 2.034a.199.199 0 0 1 .271.073l.802 1.388a.199.199 0 0 1-.073.272c-1.566.904-1.566 3.164 0 4.069a.199.199 0 0 1 .073.271l-.802 1.388a.199.199 0 0 1-.271.073C10.958 10.863 9 11.993 9 13.8a.2.2 0 0 1-.199.2H7.2a.199.199 0 0 1-.2-.2c0-1.808-1.958-2.938-3.524-2.034a.199.199 0 0 1-.272-.073l-.8-1.388a.199.199 0 0 1 .072-.271c1.566-.905 1.566-3.165 0-4.07a.199.199 0 0 1-.073-.271l.801-1.388a.199.199 0 0 1 .272-.073C5.042 5.138 7 4.007 7 2.2c0-.11.089-.199.199-.199ZM5.5 2.2c0-.94.76-1.7 1.699-1.7H8.8c.94 0 1.7.76 1.7 1.7a.85.85 0 0 0 1.274.735a1.699 1.699 0 0 1 2.32.622l.802 1.388c.469.813.19 1.851-.622 2.32a.85.85 0 0 0 0 1.472a1.7 1.7 0 0 1 .622 2.32l-.802 1.388a1.699 1.699 0 0 1-2.32.622a.85.85 0 0 0-1.274.735c0 .939-.76 1.7-1.699 1.7H7.2a1.7 1.7 0 0 1-1.699-1.7a.85.85 0 0 0-1.274-.735a1.698 1.698 0 0 1-2.32-.622l-.802-1.388a1.699 1.699 0 0 1 .622-2.32a.85.85 0 0 0 0-1.471a1.699 1.699 0 0 1-.622-2.321l.801-1.388a1.699 1.699 0 0 1 2.32-.622A.85.85 0 0 0 5.5 2.2m4 5.8a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0M11 8a3 3 0 1 1-6 0a3 3 0 0 1 6 0" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiGear0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GearBranches(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiGearBranches0)"><path fill="currentColor" fill-rule="evenodd" d="M7.295 2a.25.25 0 0 0-.243.194l-.131.568a2.35 2.35 0 0 1-3.263 1.611l-.122-.055a.25.25 0 0 0-.323.107l-.696 1.277A.25.25 0 0 0 2.559 6l.971.97a.75.75 0 1 1-1.06 1.06l-.971-.97a1.75 1.75 0 0 1-.3-2.076l.697-1.277a1.75 1.75 0 0 1 2.26-.755l.123.056a.85.85 0 0 0 1.18-.583l.131-.568A1.75 1.75 0 0 1 7.295.5h1.41a1.75 1.75 0 0 1 1.705 1.357l.13.568a.85.85 0 0 0 1.18.583l.123-.056a1.75 1.75 0 0 1 2.26.755l.697 1.277a1.75 1.75 0 0 1-.299 2.075l-.97.971a.75.75 0 0 1-1.061-1.06l.97-.971a.25.25 0 0 0 .043-.297l-.696-1.277a.25.25 0 0 0-.323-.107l-.123.055a2.35 2.35 0 0 1-3.262-1.61l-.131-.57A.25.25 0 0 0 8.704 2zM3.25 12a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5m0 1.5a2.25 2.25 0 0 0 1.978-3.322l.553-.442a1.25 1.25 0 0 0 .469-.976V7.75a.75.75 0 1 0-1.5 0v.89l-.655.524A2.25 2.25 0 1 0 3.25 13.5m5.5-.25a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0m1.5 0a2.25 2.25 0 1 1-3-2.122V7.754a.75.75 0 0 1 1.5 0v3.374a2.25 2.25 0 0 1 1.5 2.122m2.5-1.25a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5m0 1.5a2.25 2.25 0 1 0-.795-4.356L11.5 8.69v-.94a.75.75 0 1 0-1.5 0v1.043c0 .331.132.65.366.884l.44.44A2.25 2.25 0 0 0 12.75 13.5" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiGearBranches0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GearDot(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g fill="currentColor" fill-rule="evenodd" clip-path="url(#gravityUiGearDot0)" clip-rule="evenodd"><path d="M7.199 2A.199.199 0 0 0 7 2.199c0 1.808-1.958 2.939-3.524 2.034a.199.199 0 0 0-.272.073l-.8 1.388a.199.199 0 0 0 .072.271c1.566.905 1.566 3.165 0 4.07a.199.199 0 0 0-.073.271l.801 1.388a.199.199 0 0 0 .272.073C5.042 10.862 7 11.993 7 13.8c0 .11.089.199.199.199H8.8c.11 0 .199-.089.199-.199c0-1.808 1.958-2.939 3.524-2.034a.199.199 0 0 0 .271-.073l.802-1.388a.199.199 0 0 0-.073-.271c-1.303-.753-1.516-2.434-.665-3.5a.75.75 0 0 1 1.172.936a.852.852 0 0 0 .243 1.265a1.7 1.7 0 0 1 .622 2.32l-.802 1.388a1.699 1.699 0 0 1-2.32.622a.85.85 0 0 0-1.274.735c0 .938-.76 1.699-1.699 1.699H7.2c-.938 0-1.699-.76-1.699-1.699a.85.85 0 0 0-1.274-.735a1.698 1.698 0 0 1-2.32-.622l-.802-1.388a1.699 1.699 0 0 1 .622-2.32a.85.85 0 0 0 0-1.472a1.699 1.699 0 0 1-.622-2.32l.801-1.388a1.699 1.699 0 0 1 2.32-.622A.85.85 0 0 0 5.5 2.2c0-.94.76-1.7 1.699-1.7H9.3a.75.75 0 1 1 0 1.5H7.2Zm.8 7.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3M8 11a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/><path d="M12.5 5.5a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/></g><defs><clipPath id="gravityUiGearDot0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GearPlay(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.199 2H8.8c.11 0 .199.089.199.199c0 1.808 1.958 2.939 3.524 2.034a.199.199 0 0 1 .271.073l.802 1.388a.199.199 0 0 1-.073.271c-1.566.905-1.566 3.165 0 4.07a.199.199 0 0 1 .073.271l-.802 1.388a.199.199 0 0 1-.271.073C10.958 10.862 9 11.993 9 13.8c0 .11-.089.199-.199.199H7.2a.199.199 0 0 1-.2-.198c0-1.808-1.958-2.939-3.524-2.034a.199.199 0 0 1-.272-.073l-.8-1.388a.199.199 0 0 1 .072-.271c1.566-.905 1.566-3.165 0-4.07a.199.199 0 0 1-.073-.271l.801-1.388a.199.199 0 0 1 .272-.073C5.042 5.138 7 4.007 7 2.2c0-.11.089-.199.199-.199ZM5.5 2.199C5.5 1.26 6.26.5 7.199.5H8.8c.938 0 1.699.76 1.699 1.699a.85.85 0 0 0 1.274.735a1.699 1.699 0 0 1 2.32.622l.802 1.388c.469.812.19 1.851-.622 2.32a.85.85 0 0 0 0 1.472a1.7 1.7 0 0 1 .622 2.32l-.802 1.388a1.699 1.699 0 0 1-2.32.622a.85.85 0 0 0-1.274.735c0 .938-.76 1.699-1.699 1.699H7.2c-.938 0-1.699-.76-1.699-1.699a.85.85 0 0 0-1.274-.735a1.698 1.698 0 0 1-2.32-.622l-.802-1.388a1.699 1.699 0 0 1 .622-2.32a.85.85 0 0 0 0-1.472a1.699 1.699 0 0 1-.622-2.32l.801-1.388a1.699 1.699 0 0 1 2.32-.622A.85.85 0 0 0 5.5 2.2Zm5 6.667a1 1 0 0 0 0-1.732l-3-1.732a1 1 0 0 0-1.5.866v3.464a1 1 0 0 0 1.5.866z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Geo(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.293 3H4.5a1.5 1.5 0 0 0-1.41.987l4.68 2.978zM3 5.707V11.5c0 .648.411 1.2.987 1.41l2.978-4.68zM5.707 13H11.5a1.5 1.5 0 0 0 1.5-1.5v-7a1.5 1.5 0 0 0-.987-1.41l-3.38 5.313zM1.5 4.5a3 3 0 0 1 3-3h7a3 3 0 0 1 3 3v7a3 3 0 0 1-3 3h-7a3 3 0 0 1-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GeoDots(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.293 3H4.5a1.5 1.5 0 0 0-1.41.987L7 6.475a2.5 2.5 0 0 1 2.654-2.47zM8.41 8.751a2.5 2.5 0 0 0 2.667-4.19l.935-1.47A1.5 1.5 0 0 1 13 4.5v7a1.5 1.5 0 0 1-1.5 1.5H5.707zM3 5.707V11.5c0 .648.411 1.2.987 1.41l.963-1.514A1.5 1.5 0 1 1 6.531 8.91l.434-.68zM1.5 4.5a3 3 0 0 1 3-3h7a3 3 0 0 1 3 3v7a3 3 0 0 1-3 3h-7a3 3 0 0 1-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GeoFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.247 1.5H4.5a3 3 0 0 0-2.69 1.672l5.96 3.793zM1.5 4.753V11.5a3 3 0 0 0 1.672 2.69l3.793-5.96zM4.753 14.5H11.5a3 3 0 0 0 3-3v-7a3 3 0 0 0-1.671-2.69L8.633 8.402z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GeoPin(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.293 3H4.5a1.5 1.5 0 0 0-1.41.987L7 6.475a2.5 2.5 0 0 1 2.654-2.47zm.785 1.56a2.501 2.501 0 0 1-.836 4.328a.75.75 0 0 1 .008.112v2a.75.75 0 0 1-1.5 0V9c0-.038.003-.075.008-.112a2.483 2.483 0 0 1-.347-.137L5.707 13H11.5a1.5 1.5 0 0 0 1.5-1.5v-7a1.5 1.5 0 0 0-.987-1.41zm-7.091 8.35l2.978-4.68L3 5.707V11.5c0 .648.411 1.2.987 1.41M1.5 4.5a3 3 0 0 1 3-3h7a3 3 0 0 1 3 3v7a3 3 0 0 1-3 3h-7a3 3 0 0 1-3-3zm9 2a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GeoPolygons(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 3h5.793L7.77 6.965L3.09 3.987A1.5 1.5 0 0 1 4.5 3m7 10H5.707l2.926-4.597l3.38-5.313A1.5 1.5 0 0 1 13 4.5v7a1.5 1.5 0 0 1-1.5 1.5m-7-11.5a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h7a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ghost(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.277 11.702L13.5 12V8a5.5 5.5 0 1 0-11 0v4.547l1.956-1.63a1.801 1.801 0 0 1 2.537.231l1.935 2.323a.08.08 0 0 0 .125-.001l1.45-1.811a1.755 1.755 0 0 1 2.774.043m-3.052 2.705l.686-.859h.001l.144-.18l.618-.772a.255.255 0 0 1 .402.006l.593.79l.139.185v.001l.392.522a1 1 0 0 0 1.8-.6V8A7 7 0 1 0 1 8v5.399a1.101 1.101 0 0 0 1.806.846l2.61-2.175a.301.301 0 0 1 .424.038l1.936 2.323a1.58 1.58 0 0 0 2.449-.024M7 8a.75.75 0 0 0 .75-.75v-1a.75.75 0 1 0-1.5 0v1c0 .414.336.75.75.75m4 0a.75.75 0 0 0 .75-.75v-1a.75.75 0 0 0-1.5 0v1c0 .414.336.75.75.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gift(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m9.035 3.863l.182-1.276a.918.918 0 0 1 .818-.783a.914.914 0 0 1 .272 1.805zM7.25 5.5v2H3a.5.5 0 0 1-.5-.5V6a.5.5 0 0 1 .5-.5zM8 1.565A2.415 2.415 0 1 0 3.83 4H3a2 2 0 0 0-2 2v1a2 2 0 0 0 1 1.732V12a3 3 0 0 0 3 3h6a3 3 0 0 0 3-3V8.732A2 2 0 0 0 15 7V6a2 2 0 0 0-2-2h-.83A2.415 2.415 0 1 0 8 1.565m.75 4.185V7.5H13a.5.5 0 0 0 .5-.5V6a.5.5 0 0 0-.5-.5H8.75zM7.25 9H3.5v3A1.5 1.5 0 0 0 5 13.5h2.25zm1.5 4.5V9h3.75v3a1.5 1.5 0 0 1-1.5 1.5zM6.783 2.587l.182 1.276l-1.272-.254a.914.914 0 0 1 .272-1.805a.919.919 0 0 1 .818.783" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globe(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.208 12.346c-.485 1-.953 1.154-1.208 1.154s-.723-.154-1.208-1.154c-.372-.768-.647-1.858-.749-3.187a20.77 20.77 0 0 0 3.914 0c-.102 1.329-.377 2.419-.75 3.187Zm.788-4.699C9.358 7.714 8.69 7.75 8 7.75s-1.358-.036-1.996-.103c.037-1.696.343-3.075.788-3.993C7.277 2.654 7.745 2.5 8 2.5s.723.154 1.208 1.154c.445.918.75 2.297.788 3.993m1.478 1.306c-.085 1.516-.375 2.848-.836 3.874a5.501 5.501 0 0 0 2.843-4.364c-.621.199-1.295.364-2.007.49m1.918-2.043c-.572.204-1.21.379-1.901.514c-.056-1.671-.354-3.14-.853-4.251a5.508 5.508 0 0 1 2.754 3.737m-8.883.514c.056-1.671.354-3.14.853-4.251A5.508 5.508 0 0 0 2.608 6.91c.572.204 1.21.379 1.901.514M2.52 8.463a5.501 5.501 0 0 0 2.843 4.364c-.46-1.026-.75-2.358-.836-3.874a15.535 15.535 0 0 1-2.007-.49M15 8A7 7 0 1 0 1 8a7 7 0 0 0 14 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gpu(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiGpu0)"><path fill="currentColor" fill-rule="evenodd" d="M0 1.25A.75.75 0 0 1 .75.5h.5A2.25 2.25 0 0 1 3.5 2.75V3H13a3 3 0 0 1 3 3v4a3 3 0 0 1-3 3H3.5v1.25a.75.75 0 0 1-1.5 0V2.75A.75.75 0 0 0 1.25 2h-.5A.75.75 0 0 1 0 1.25M13 4.5H3.5v7H13a1.5 1.5 0 0 0 1.5-1.5V6A1.5 1.5 0 0 0 13 4.5M8.5 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0M10 8a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiGpu0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraduationCap(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.836 3.202L1.74 5.386a.396.396 0 0 0 0 .728l5.096 2.184a2.5 2.5 0 0 0 .985.202h.358a2.5 2.5 0 0 0 .985-.202l5.096-2.184a.396.396 0 0 0 0-.728L9.164 3.202A2.5 2.5 0 0 0 8.179 3h-.358a2.5 2.5 0 0 0-.985.202M1.5 7.642l1.5.644v3.228a2 2 0 0 0 1.106 1.789l.806.403a7 7 0 0 0 6.193.033l.909-.442a2 2 0 0 0 1.125-1.798V8.226l1.712-.734a1.896 1.896 0 0 0 0-3.484L9.755 1.823A4 4 0 0 0 8.179 1.5h-.358a4 4 0 0 0-1.576.323L1.15 4.008A1.896 1.896 0 0 0 0 5.75v4.5a.75.75 0 0 0 1.5 0zm3 3.872V8.929l1.745.748A4 4 0 0 0 7.821 10h.358a4 4 0 0 0 1.576-.323l1.884-.808v2.63a.5.5 0 0 1-.282.45l-.909.442a5.5 5.5 0 0 1-4.865-.027l-.807-.403a.5.5 0 0 1-.276-.447" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraphNode(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.75 2.5h4.5a1.5 1.5 0 0 1 1.269.7a2.5 2.5 0 0 0 .231 4.686V12a1.5 1.5 0 0 1-1.5 1.5h-4.5c-.534 0-1.003-.28-1.269-.7a2.5 2.5 0 0 0-.231-4.686v-.228A2.501 2.501 0 0 0 4.481 3.2c.266-.42.735-.7 1.269-.7m-3 5.614v-.228a2.501 2.501 0 0 1 .146-4.813A3.001 3.001 0 0 1 5.75 1h4.5a3 3 0 0 1 2.854 2.074a2.501 2.501 0 0 1 .146 4.812V12a3 3 0 0 1-3 3h-4.5a3.001 3.001 0 0 1-2.854-2.073a2.501 2.501 0 0 1-.146-4.813M3.5 11.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2m-1-6a1 1 0 1 1 2 0a1 1 0 0 1-2 0m10-1a1 1 0 1 0 0 2a1 1 0 0 0 0-2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grip(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 3a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0M5.5 9.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m5 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3M7 13a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0m3.5 1.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GripHorizontal(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 9a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3m6.5 1.5a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0m0-5a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0m-5 0a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0M13 9a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3m1.5-3.5a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hand(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiHand0)"><path fill="currentColor" fill-rule="evenodd" d="M6.5 3.325v5.971L5 8.921l-.53-.132l-.672-.168l-.205-.051a.478.478 0 0 0-.409.84l.167.13l.547.425l.43.334l2.142 1.666l2.728-.181L12.5 11.5l.667-3.558l.108-.578l.455-2.426l.148-.787l.009-.047a.51.51 0 0 0-.993-.234l-.013.046l-.22.77l-.161.564L12 7h-1V2.023a.523.523 0 0 0-1.043-.048l-.005.048l-.07.775l-.216 2.381l-.147 1.617L9.5 7h-1l-.019-.093l-.329-1.647L8 4.5l-.26-1.298l-.13-.652l-.022-.108A.55.55 0 0 0 6.5 2.55zm-2.543 3.79L5 7.374V2.55a2.05 2.05 0 0 1 3.617-1.321a2.023 2.023 0 0 1 3.883.794v.182a2.01 2.01 0 0 1 2.861 2.176l-1.387 7.395a1.5 1.5 0 0 1-1.346 1.219l-3.302.283l-.028.002l-2.728.182a1.5 1.5 0 0 1-1.02-.313l-3.287-2.555a1.978 1.978 0 0 1 1.694-3.48Zm9.346 8.383a.75.75 0 0 0-.106-1.496l-7 .5a.75.75 0 0 0 .106 1.496z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiHand0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandOk(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiHandOk0)"><path fill="currentColor" fill-rule="evenodd" d="M6.5 6.963a1.186 1.186 0 0 0-1.652-1.09a1.213 1.213 0 0 0-.271.165l-.207.166l-.081.065a.578.578 0 0 1-.796-.832l.068-.078l.247-.282l.158-.18l.028-.033c.178-.203.38-.366.596-.49c1.175-.677 2.777-.206 3.323 1.16l.174.435l.375.937L8.5 7h1l.038-.415l.076-.839l.268-2.948l.07-.775l.005-.048A.523.523 0 0 1 11 2.023V7h1l.5-1.75l.161-.565l.22-.769l.013-.046a.51.51 0 0 1 .993.234l-.01.047l-.147.787l-.455 2.426l-.108.579L12.5 11.5l-3.302.284l-2.728.181l-.68-.189a4.431 4.431 0 0 1-3.2-3.643a.554.554 0 0 1 .549-.633h.133c.14 0 .27.076.34.2l.102.185l.045.08l.253.454l.128.23a1.258 1.258 0 0 0 2.36-.612zM1.105 8.345a2.052 2.052 0 0 1 .826-1.953a2.07 2.07 0 0 1 .433-1.943l.5-.572C4.34 2.192 6.799 2.19 8.309 3.542l.155-1.702a2.023 2.023 0 0 1 4.037.183v.182a2.01 2.01 0 0 1 2.861 2.176l-1.387 7.396a1.5 1.5 0 0 1-1.346 1.218l-3.302.283l-.028.002l-2.728.182a1.499 1.499 0 0 1-.502-.051l-.68-.19a5.933 5.933 0 0 1-4.283-4.876Zm12.198 7.153a.75.75 0 1 0-.107-1.496l-7 .5a.75.75 0 0 0 .107 1.496z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiHandOk0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandPointDown(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiHandPointDown0)"><path fill="currentColor" fill-rule="evenodd" d="M6 2.5L2.263 5.406a1.978 1.978 0 0 0 1.694 3.48L5 8.625v5.433a1.942 1.942 0 0 0 3.838.421L9.5 11.5l2.744-.457a3 3 0 0 0 2.405-3.732L13.5 3zm7.2 5.197l-.872-3.272l-5.858-.39L3.184 6.59a.478.478 0 0 0 .41.84l1.042-.26L6.5 6.704v7.354a.442.442 0 0 0 .874.096l.662-2.98l.22-.987l.997-.167l2.744-.457A1.5 1.5 0 0 0 13.2 7.697m.103-7.195a.75.75 0 0 1-.106 1.496l-7-.5A.75.75 0 1 1 6.303.002z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiHandPointDown0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandPointLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiHandPointLeft0)"><path fill="currentColor" fill-rule="evenodd" d="m13.5 6l-2.906-3.737a1.978 1.978 0 0 0-3.48 1.694L7.375 5H1.942a1.942 1.942 0 0 0-.421 3.838L4.5 9.5l.457 2.744a3 3 0 0 0 3.733 2.405L13 13.5zm-5.197 7.2l3.272-.872l.39-5.858L9.41 3.184a.478.478 0 0 0-.84.41l.26 1.042l.466 1.864H1.942a.442.442 0 0 0-.096.874l2.98.662l.987.22l.167.997l.457 2.744A1.5 1.5 0 0 0 8.303 13.2m7.195.103a.75.75 0 0 1-1.496-.106l.5-7a.75.75 0 1 1 1.496.106z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiHandPointLeft0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandPointRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiHandPointRight0)"><path fill="currentColor" fill-rule="evenodd" d="m2.5 6l2.906-3.737a1.978 1.978 0 0 1 3.48 1.694L8.626 5h5.432a1.942 1.942 0 0 1 .421 3.838L11.5 9.5l-.457 2.744a3 3 0 0 1-3.732 2.405L3 13.5zm5.197 7.2l-3.272-.872l-.39-5.858L6.59 3.184a.478.478 0 0 1 .84.41l-.26 1.042L6.704 6.5h7.354a.442.442 0 0 1 .096.874l-2.98.662l-.987.22l-.166.997l-.458 2.744A1.5 1.5 0 0 1 7.697 13.2m-7.195.103a.75.75 0 0 0 1.496-.106l-.5-7a.75.75 0 1 0-1.496.106z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiHandPointRight0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandPointUp(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiHandPointUp0)"><path fill="currentColor" fill-rule="evenodd" d="m6 13.5l-3.737-2.906a1.978 1.978 0 0 1 1.694-3.48L5 7.375V1.942a1.942 1.942 0 0 1 3.838-.421L9.5 4.5l2.744.457a3 3 0 0 1 2.405 3.733L13.5 13zm7.2-5.197l-.872 3.272l-5.858.39L3.184 9.41a.478.478 0 0 1 .41-.84l1.042.26l1.864.466V1.942a.442.442 0 0 1 .874-.096l.662 2.98l.22.987l.997.167l2.744.457A1.5 1.5 0 0 1 13.2 8.303m.103 7.195a.75.75 0 0 0-.106-1.496l-7 .5a.75.75 0 1 0 .106 1.496z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiHandPointUp0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandStop(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.618 8.405L3.47 8.118a1.988 1.988 0 0 0-1.8 3.417l2.198 1.946a7.897 7.897 0 0 0 6.14 1.932A5.339 5.339 0 0 0 14.72 9.71l-.437-5.83a2.033 2.033 0 0 0-2.511-1.822l-.005-.113a2.02 2.02 0 0 0-3.866-.739A2.04 2.04 0 0 0 4.97 2.91l-.353 5.495Zm1.673-2.64l.126-1.968l.045-.707l.006-.084a.54.54 0 0 1 1.077-.015l.007.084l.064.706l.032.346l.147 1.618l.201 2.216L8 8h1l.004-.091l.146-3.822l.049-1.256l.03-.792l.002-.04A.52.52 0 0 1 10.268 2l.002.04l.03.792l.02.576l.02.524l.153 4.125l.002.063l.003.077l.002.053h1l.002-.037l.04-.762l.139-2.64l.039-.738l.003-.068a.53.53 0 0 1 .533-.505c.278 0 .51.215.531.493l.005.068l.055.737l.198 2.636l.18 2.388a3.839 3.839 0 0 1-3.389 4.1a6.398 6.398 0 0 1-4.974-1.565l-1.259-1.114L3.328 11l-.473-.42l-.19-.168a.488.488 0 0 1 .442-.84l.246.062l.613.154l.357.089l.198.05l1.479.37l.098-1.523z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Handset(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.124 9.833a3.712 3.712 0 0 0-1.59-1.22l-.033-.012a2.877 2.877 0 0 0-2.795.37l-.36.269a.5.5 0 0 1-.653-.047L6.807 7.307a.5.5 0 0 1-.047-.654l.27-.36A2.877 2.877 0 0 0 7.4 3.5l-.013-.032a3.713 3.713 0 0 0-1.22-1.592l-.19-.143a2.33 2.33 0 0 0-2.135-.346l-.19.063A3.484 3.484 0 0 0 1.45 3.653a4.878 4.878 0 0 0-.105 2.725a9.758 9.758 0 0 0 2.567 4.533l1.178 1.178a9.758 9.758 0 0 0 4.533 2.566c.9.226 1.845.19 2.725-.104a3.484 3.484 0 0 0 2.204-2.204l.063-.189a2.33 2.33 0 0 0-.347-2.135l-.143-.19ZM5.16 3.219l-.19-.143a.65.65 0 0 0-.596-.096l-.19.063c-.538.18-.96.602-1.14 1.14a3.198 3.198 0 0 0-.069 1.788A8.077 8.077 0 0 0 5.1 9.723l1.178 1.178a8.078 8.078 0 0 0 3.752 2.125c.59.147 1.21.123 1.787-.069c.539-.18.961-.602 1.14-1.14l.064-.19a.65.65 0 0 0-.097-.596l-.143-.19a2.034 2.034 0 0 0-.87-.668l-.033-.013a1.197 1.197 0 0 0-1.163.154l-.36.27a2.18 2.18 0 0 1-2.849-.203L5.62 8.495a2.18 2.18 0 0 1-.203-2.85l.27-.36c.25-.334.309-.774.154-1.162l-.013-.032a2.033 2.033 0 0 0-.668-.872" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HardDrive(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 8.5a1.5 1.5 0 0 1 1.5 1.5v1a1.5 1.5 0 0 1-1.5 1.5H4A1.5 1.5 0 0 1 2.5 11v-1A1.5 1.5 0 0 1 4 8.5zm.89-1.366L11.488 4.33a1.5 1.5 0 0 0-1.342-.829H5.854a1.5 1.5 0 0 0-1.342.83L3.11 7.133A3 3 0 0 1 4 7h8a3 3 0 0 1 .89.134M15 9.18V11a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3V9.18a5 5 0 0 1 .528-2.236L3.17 3.658A3 3 0 0 1 5.854 2h4.292a3 3 0 0 1 2.683 1.658l1.643 3.286A5 5 0 0 1 15 9.18m-6 .57a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hashtag(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.238 2.634a.75.75 0 1 0-1.476-.268L5.283 5H3a.75.75 0 1 0 0 1.5h2.01l-.545 3H2A.75.75 0 1 0 2 11h2.192l-.43 2.366a.75.75 0 1 0 1.476.268L5.717 11h3.475l-.43 2.366a.75.75 0 1 0 1.476.268L10.717 11H13a.75.75 0 0 0 0-1.5h-2.01l.545-3H14A.75.75 0 0 0 14 5h-2.192l.43-2.366a.75.75 0 1 0-1.476-.268L10.283 5H6.808zM9.465 9.5l.545-3H6.535l-.545 3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heading(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.25 2.25A.75.75 0 0 1 5 3v4.25h6V3a.75.75 0 0 1 1.5 0v10a.75.75 0 0 1-1.5 0V8.75H5V13a.75.75 0 0 1-1.5 0V3a.75.75 0 0 1 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadingFive(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 4.25a.75.75 0 0 0-1.5 0v7.496a.75.75 0 0 0 1.5 0V8.75h4v2.996a.75.75 0 0 0 1.5 0V4.25a.75.75 0 0 0-1.5 0v3h-4zm8-.75a1.25 1.25 0 0 0-1.25 1.25V7.5a.75.75 0 0 0 1.221.584a1.497 1.497 0 0 1 .269-.134c.235-.091.643-.2 1.26-.2c.892 0 1.534.751 1.534 1.75c0 .366-.139.735-.413 1.01c-.265.268-.7.49-1.371.49c-.568 0-.905-.184-1.094-.336a1.168 1.168 0 0 1-.24-.26l-.002-.003a.75.75 0 0 0-1.335.684l.671-.335l-.67.336v.003l.003.003l.004.01l.012.02l.034.06a2.668 2.668 0 0 0 .586.653c.436.35 1.099.665 2.031.665c1.027 0 1.86-.352 2.437-.933a2.94 2.94 0 0 0 .847-2.067c0-1.591-1.096-3.25-3.034-3.25a5.55 5.55 0 0 0-1.25.133V5H14a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadingFour(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.25 4.25a.75.75 0 0 0-1.347-.454l-4 5.25A.75.75 0 0 0 9.5 10.25h3.25v1.496a.75.75 0 0 0 1.5 0V10.25h.25a.75.75 0 1 0 0-1.5h-.25zm-1.5 2.222V8.75h-1.736zM2.5 4.25a.75.75 0 1 0-1.5 0v7.496a.75.75 0 0 0 1.5 0V8.75h4v2.996a.75.75 0 0 0 1.5 0V4.25a.75.75 0 1 0-1.5 0v3h-4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadingOne(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14 4.25a.75.75 0 0 0-1.248-.56l-2.25 2a.75.75 0 0 0 .996 1.12l1.002-.89v5.83a.75.75 0 0 0 1.5 0zm-11.5 0a.75.75 0 0 0-1.5 0v7.496a.75.75 0 0 0 1.5 0V8.75h4v2.996a.75.75 0 0 0 1.5 0V4.25a.75.75 0 0 0-1.5 0v3h-4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadingSix(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 4.25a.75.75 0 0 0-1.5 0v7.496a.75.75 0 0 0 1.5 0V8.75h4v2.996a.75.75 0 0 0 1.5 0V4.25a.75.75 0 0 0-1.5 0v3h-4zm9.9 2.532a4.857 4.857 0 0 0-1.843.293L12.273 5c.888 0 1.293.251 1.801.566l.006.003a.75.75 0 0 0 .79-1.275c-.59-.365-1.286-.794-2.597-.794c-.992 0-1.855.424-2.442 1.255C9.267 5.553 9 6.665 9 8c0 1.326.246 2.438.802 3.24c.583.842 1.453 1.26 2.47 1.26c1.918 0 3.228-1.236 3.228-2.795c0-.984-.388-1.74-1.024-2.237c-.607-.474-1.372-.661-2.076-.686M11.056 5.62c-.225.32-.41.794-.498 1.455L12.272 5c-.522 0-.92.201-1.217.62Zm-.021 4.765c-.233-.336-.418-.849-.495-1.58c.048-.056.165-.165.436-.282c.372-.16.873-.26 1.371-.242c.503.018.926.15 1.206.369c.25.195.447.504.447 1.055c0 .58-.48 1.295-1.727 1.295c-.564 0-.956-.207-1.238-.615" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadingThree(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 4.25a.75.75 0 0 0-1.5 0v7.496a.75.75 0 0 0 1.5 0V8.75h4v2.996a.75.75 0 0 0 1.5 0V4.25a.75.75 0 0 0-1.5 0v3h-4zm8.114 1.496c.202-.504.69-.834 1.232-.834h.28a.94.94 0 0 1 .929.796l.027.18a1.15 1.15 0 0 1-.911 1.303l-.8.16a.662.662 0 0 0 .129 1.31h1.21a.89.89 0 0 1 .882 1.017a1.666 1.666 0 0 1-1.414 1.414l-.103.015a1.806 1.806 0 0 1-1.828-.9l-.018-.033a.662.662 0 0 0-1.152.652l.018.032a3.129 3.129 0 0 0 3.167 1.559l.103-.015a2.99 2.99 0 0 0 2.537-2.537a2.213 2.213 0 0 0-1.058-2.216a2.47 2.47 0 0 0 .546-1.963l-.027-.179a2.263 2.263 0 0 0-2.237-1.919h-.28a2.65 2.65 0 0 0-2.46 1.666a.662.662 0 1 0 1.228.492" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadingTwo(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 4.25a.75.75 0 0 0-1.5 0v7.496a.75.75 0 0 0 1.5 0V8.75h4v2.996a.75.75 0 0 0 1.5 0V4.25a.75.75 0 0 0-1.5 0v3h-4zm8.403 1.783A1.364 1.364 0 0 1 12.226 5h.226a1.071 1.071 0 0 1 .672 1.906l-3.61 2.906a1.51 1.51 0 0 0 .947 2.688h3.789a.75.75 0 0 0 0-1.5h-3.793l-.003-.003a.012.012 0 0 1-.003-.004v-.004a.01.01 0 0 1 .004-.008l3.61-2.907A2.571 2.571 0 0 0 12.452 3.5h-.226c-1.314 0-2.46.894-2.778 2.17l-.038.148a.75.75 0 1 0 1.456.364z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Headphones(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 8A7 7 0 1 0 1 8v4h.01a3.25 3.25 0 0 0 3.24 3h.083C5.253 15 6 14.254 6 13.333v-3.166C6 9.247 5.254 8.5 4.333 8.5H4.25c-.644 0-1.245.188-1.75.51V8a5.5 5.5 0 1 1 11 0v1.01a3.235 3.235 0 0 0-1.75-.51h-.083c-.92 0-1.667.746-1.667 1.667v3.166c0 .92.746 1.667 1.667 1.667h.083a3.25 3.25 0 0 0 3.24-3H15zm-1.5 3.75A1.75 1.75 0 0 0 11.75 10h-.083a.167.167 0 0 0-.167.167v3.166c0 .092.075.167.167.167h.083a1.75 1.75 0 0 0 1.75-1.75M4.25 13.5a1.75 1.75 0 1 1 0-3.5h.083c.092 0 .167.075.167.167v3.166a.167.167 0 0 1-.167.167z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.633 2.796c.762-.837 1.85-1.297 3.127-1.297c1.164 0 2.407.55 3.24 1.626c.828-1.075 2.066-1.626 3.24-1.626c1.274 0 2.36.458 3.124 1.293c.756.828 1.136 1.962 1.136 3.22c0 2.166-1.113 3.909-2.522 5.264c-1.405 1.352-3.17 2.383-4.633 3.14a.75.75 0 0 1-.693-.002c-1.463-.765-3.228-1.788-4.633-3.133C1.61 9.93.5 8.193.5 6.013c0-1.255.378-2.389 1.133-3.217m1.109 1.01C2.287 4.306 2 5.053 2 6.013c0 1.624.816 2.996 2.057 4.184c1.146 1.098 2.6 1.985 3.945 2.705c1.335-.71 2.79-1.604 3.937-2.707C13.182 8.998 14 7.62 14 6.013c0-.963-.288-1.71-.744-2.21C12.808 3.314 12.14 3 11.24 3c-.976 0-2.093.628-2.527 1.95a.75.75 0 0 1-1.426 0C6.854 3.63 5.725 3 4.76 3c-.903 0-1.57.315-2.018.807Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartCrack(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.76 1.5c-1.278 0-2.365.459-3.127 1.296C.878 3.624.5 4.758.5 6.013c0 2.18 1.11 3.917 2.52 5.268c1.404 1.345 3.17 2.368 4.632 3.133a.75.75 0 0 0 .693.002c1.463-.757 3.228-1.788 4.633-3.14c1.41-1.355 2.522-3.098 2.522-5.263c0-1.26-.38-2.393-1.136-3.221c-.764-.835-1.85-1.293-3.124-1.293c-1.174 0-2.412.551-3.24 1.626c-.833-1.075-2.076-1.626-3.24-1.626ZM2 6.012c0-.96.287-1.708.742-2.207c.447-.492 1.115-.807 2.018-.807c.9 0 1.94.547 2.43 1.69L6.066 7.192a.75.75 0 0 0 .298.95l2.01 1.206l-.905 3.261c-1.187-.664-2.416-1.458-3.412-2.413C2.817 9.01 2 7.637 2 6.013Zm7.119 6.262c.993-.59 1.987-1.28 2.82-2.08C13.182 8.998 14 7.62 14 6.013c0-.963-.288-1.71-.744-2.21C12.808 3.314 12.14 3 11.24 3c-.974 0-2.088.625-2.525 1.94a.746.746 0 0 1-.034.088L7.707 7.2l1.929 1.158a.75.75 0 0 1 .337.844l-.854 3.074Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.76 1.5c-1.278 0-2.365.459-3.127 1.296C.878 3.624.5 4.758.5 6.013c0 2.18 1.11 3.917 2.52 5.268c1.404 1.345 3.17 2.368 4.632 3.133a.75.75 0 0 0 .693.002c1.463-.757 3.228-1.788 4.633-3.14c1.41-1.355 2.522-3.098 2.522-5.263c0-1.26-.38-2.393-1.136-3.221c-.763-.835-1.85-1.293-3.124-1.293c-1.076 0-1.966.399-2.643 1.151A4.515 4.515 0 0 0 8 3.504a4.515 4.515 0 0 0-.597-.854C6.726 1.898 5.836 1.5 4.76 1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartPulse(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.76 1.5c-1.123 0-2.102.354-2.844 1.015C1.177 3.174.723 4.092.564 5.138a.75.75 0 1 0 1.483.225c.115-.754.428-1.336.867-1.728c.437-.39 1.048-.636 1.846-.636c.965 0 2.094.63 2.527 1.95a.75.75 0 0 0 1.426 0c.434-1.322 1.551-1.95 2.527-1.95c.9 0 1.568.314 2.016.805c.456.498.744 1.246.744 2.209c0 1.607-.818 2.985-2.061 4.182c-1.147 1.103-2.602 1.996-3.937 2.707c-1.339-.717-2.788-1.6-3.932-2.692a.75.75 0 1 0-1.036 1.085c1.402 1.338 3.16 2.357 4.618 3.12a.75.75 0 0 0 .693 0c1.463-.756 3.228-1.787 4.633-3.139c1.41-1.355 2.522-3.098 2.522-5.263c0-1.26-.38-2.393-1.136-3.221c-.764-.835-1.85-1.293-3.124-1.293c-1.174 0-2.412.551-3.24 1.626c-.833-1.075-2.076-1.626-3.24-1.626Zm.85 4.314a.75.75 0 0 0-1.108-.125L2.465 7.5H.75a.75.75 0 0 0 0 1.5h2a.75.75 0 0 0 .498-.19l1.627-1.445l2.015 2.821a.75.75 0 0 0 1.21.014l1.275-1.7H11A.75.75 0 0 0 11 7H9a.75.75 0 0 0-.6.3l-.885 1.18z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hierarchy(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.25 2.5h-8.5a.75.75 0 0 0-.75.75v1.5c0 .414.336.75.75.75h8.5a.75.75 0 0 0 .75-.75v-1.5a.75.75 0 0 0-.75-.75m0 4.5H5v3.13c0 .69.56 1.25 1.25 1.25H7v-.13A2.25 2.25 0 0 1 9.25 9h3a2.25 2.25 0 0 1 2.25 2.25v1.5A2.25 2.25 0 0 1 12.25 15h-3a2.25 2.25 0 0 1-2.246-2.12H6.25a2.75 2.75 0 0 1-2.75-2.75V6.986a2.25 2.25 0 0 1-2-2.236v-1.5A2.25 2.25 0 0 1 3.75 1h8.5a2.25 2.25 0 0 1 2.25 2.25v1.5A2.25 2.25 0 0 1 12.25 7m-3 3.5h3a.75.75 0 0 1 .75.75v1.5a.75.75 0 0 1-.75.75h-3a.75.75 0 0 1-.75-.75v-1.5a.75.75 0 0 1 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func House(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 12.618c.307-.275.5-.674.5-1.118V6.977a1.5 1.5 0 0 0-.585-1.189l-3.5-2.692a1.5 1.5 0 0 0-1.83 0l-3.5 2.692A1.5 1.5 0 0 0 3 6.978V11.5A1.496 1.496 0 0 0 4.493 13H5V9.5a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2V13h.507c.381-.002.73-.146.993-.382m2-1.118a3 3 0 0 1-3 3h-7a3 3 0 0 1-3-3V6.977A3 3 0 0 1 2.67 4.6l3.5-2.692a3 3 0 0 1 3.66 0l3.5 2.692a3.003 3.003 0 0 1 1.17 2.378zm-5-2A.5.5 0 0 0 9 9H7a.5.5 0 0 0-.5.5V13h3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Italic(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.25 2a.75.75 0 0 0 0 1.5h1.317l-2.7 9H4.25a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 0-1.5H7.433l2.7-9h1.617a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Key(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.313 7.488L9 7.653v5.37a.5.5 0 0 1-.353.478l-1.62.498l-.006.001h-.008a.024.024 0 0 1-.007-.006a.023.023 0 0 1-.005-.007v-.003L7 13.979V7.653l-1.313-.165a1.5 1.5 0 0 1-1.271-1.144l-.588-2.5A1.5 1.5 0 0 1 5.288 2h5.424a1.5 1.5 0 0 1 1.46 1.844l-.588 2.5a1.5 1.5 0 0 1-1.271 1.144m2.731-.8A3 3 0 0 1 10.5 8.975v4.047a2 2 0 0 1-1.412 1.911l-1.62.499A1.521 1.521 0 0 1 5.5 13.979V8.977a3 3 0 0 1-2.544-2.29l-.588-2.5A3 3 0 0 1 5.288.5h5.424a3 3 0 0 1 2.92 3.687l-.588 2.5ZM6.75 3.5a.75.75 0 0 0 0 1.5h2.5a.75.75 0 1 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyboard(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 4h10a1.5 1.5 0 0 1 1.5 1.5v5A1.5 1.5 0 0 1 13 12H3a1.5 1.5 0 0 1-1.5-1.5v-5A1.5 1.5 0 0 1 3 4M0 5.5a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3H3a3 3 0 0 1-3-3zm6.25 3.25a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5zM4.5 6.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0m2 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2m4-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0m2 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2m-8 2a1 1 0 1 1-2 0a1 1 0 0 1 2 0m8 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Layers(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m2.789 5.283l4.037-2.02A2.5 2.5 0 0 1 7.944 3h.112c.388 0 .77.09 1.118.264l4.037 2.019a.522.522 0 0 1 0 .934l-4.037 2.02a2.5 2.5 0 0 1-1.118.263h-.112a2.5 2.5 0 0 1-1.118-.264L2.79 6.217a.523.523 0 0 1 0-.934ZM1 5.75c0-.766.433-1.466 1.118-1.809l4.037-2.019a4 4 0 0 1 1.79-.422h.11a4 4 0 0 1 1.79.422l4.037 2.019a2.023 2.023 0 0 1 0 3.618l-.882.44l.882.442a2.023 2.023 0 0 1 0 3.618l-4.037 2.019a4 4 0 0 1-1.79.422h-.11a4 4 0 0 1-1.79-.422l-4.037-2.02a2.023 2.023 0 0 1 0-3.617L3 8l-.882-.441A2.023 2.023 0 0 1 1 5.75m3.677 3.088l-1.888.945a.523.523 0 0 0 0 .934l4.037 2.019A2.5 2.5 0 0 0 7.944 13h.112a2.5 2.5 0 0 0 1.118-.264l4.037-2.019a.523.523 0 0 0 0-.934l-1.888-.945l-1.478.74a4 4 0 0 1-1.79.422h-.11a4 4 0 0 1-1.79-.422z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayersThreeDiagonal(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiLayers3Diagonal0)"><path fill="currentColor" fill-rule="evenodd" d="M6.702 3.013L8 3.5V2.386A2 2 0 0 1 10.702.513l3.351 1.257A3 3 0 0 1 16 4.579v4.035a2 2 0 0 1-2.702 1.873L12 10v1.114a2 2 0 0 1-2.702 1.873L8 12.5v1.114a2 2 0 0 1-2.702 1.873L1.947 14.23A3 3 0 0 1 0 11.421V7.386a2 2 0 0 1 2.702-1.873L4 6V4.886a2 2 0 0 1 2.702-1.873M5.5 6.563l.553.207A3 3 0 0 1 8 9.579v1.319l1.824.684a.5.5 0 0 0 .676-.468V7.079a1.5 1.5 0 0 0-.973-1.405L6.176 4.418a.5.5 0 0 0-.676.468zm4.553-2.293L9.5 4.062V2.386a.5.5 0 0 1 .676-.468l3.35 1.256c.586.22.974.78.974 1.405v4.035a.5.5 0 0 1-.676.468L12 8.398V7.079a3 3 0 0 0-1.947-2.809M1.5 11.421V7.386a.5.5 0 0 1 .676-.468l3.35 1.257c.586.219.974.779.974 1.404v4.035a.5.5 0 0 1-.676.468l-3.35-1.257a1.5 1.5 0 0 1-.974-1.404" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiLayers3Diagonal0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayersVertical(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.283 13.211L3.264 9.174A2.5 2.5 0 0 1 3 8.056v-.112c0-.388.09-.77.264-1.118L5.283 2.79a.523.523 0 0 1 .934 0l2.02 4.037A2.5 2.5 0 0 1 8.5 7.944v.112c0 .388-.09.77-.264 1.118L6.217 13.21a.522.522 0 0 1-.934 0ZM5.75 15a2.023 2.023 0 0 1-1.809-1.118L1.922 9.845a4 4 0 0 1-.422-1.79v-.11a4 4 0 0 1 .422-1.79l2.019-4.037a2.023 2.023 0 0 1 3.618 0L8 3l.441-.882a2.023 2.023 0 0 1 3.618 0l2.019 4.037a4 4 0 0 1 .422 1.79v.11a4 4 0 0 1-.422 1.79l-2.019 4.037a2.022 2.022 0 0 1-3.618 0L8 13l-.44.882A2.023 2.023 0 0 1 5.75 15m3.089-3.677l.944 1.888a.522.522 0 0 0 .934 0l2.019-4.037A2.5 2.5 0 0 0 13 8.056v-.112a2.5 2.5 0 0 0-.264-1.118L10.717 2.79a.522.522 0 0 0-.934 0l-.944 1.888l.739 1.478A4 4 0 0 1 10 7.945v.11a4 4 0 0 1-.422 1.79l-.74 1.478Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutCells(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.25 3.5h-2.5v2h2.5zM4 3.5h1.25v2H2.5V5A1.5 1.5 0 0 1 4 3.5M2.5 7h2.75v2H2.5zm0 3.5v.5A1.5 1.5 0 0 0 4 12.5h1.25v-2zm4.25 0v2h2.5v-2zm4 0v2H12a1.5 1.5 0 0 0 1.5-1.5v-.5zM13.5 9h-2.75V7h2.75zM9.25 9h-2.5V7h2.5zm1.5-5.5v2h2.75V5A1.5 1.5 0 0 0 12 3.5zM4 2a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutCellsLarge(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 3.5H8.75v3.75h4.75V5A1.5 1.5 0 0 0 12 3.5m1.5 5.25H8.75v3.75H12a1.5 1.5 0 0 0 1.5-1.5zm-6.25-1.5V3.5H4A1.5 1.5 0 0 0 2.5 5v2.25zM2.5 8.75h4.75v3.75H4A1.5 1.5 0 0 1 2.5 11zM4 2a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutColumns(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.75 3.5H12A1.5 1.5 0 0 1 13.5 5v6a1.5 1.5 0 0 1-1.5 1.5H8.75zm-1.5 0H4A1.5 1.5 0 0 0 2.5 5v6A1.5 1.5 0 0 0 4 12.5h3.25zM1 5a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutColumnsThree(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 3.5h-1.25v9H12a1.5 1.5 0 0 0 1.5-1.5V5A1.5 1.5 0 0 0 12 3.5m-5.25 0h2.5v9h-2.5zm-1.5 0H4A1.5 1.5 0 0 0 2.5 5v6A1.5 1.5 0 0 0 4 12.5h1.25zM4 2a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutFooter(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 12.5h8a1.5 1.5 0 0 0 1.5-1.5v-.5h-11v.5A1.5 1.5 0 0 0 4 12.5M2.5 9V5A1.5 1.5 0 0 1 4 3.5h8A1.5 1.5 0 0 1 13.5 5v4zM1 11a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3H4a3 3 0 0 0-3 3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutHeader(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 3.5h8A1.5 1.5 0 0 1 13.5 5v.5h-11V5A1.5 1.5 0 0 1 4 3.5M2.5 7v4A1.5 1.5 0 0 0 4 12.5h8a1.5 1.5 0 0 0 1.5-1.5V7zM1 5a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutHeaderCells(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 3.5H4A1.5 1.5 0 0 0 2.5 5v.5h11V5A1.5 1.5 0 0 0 12 3.5M2.5 9V7h2.75v2zm0 1.5v.5A1.5 1.5 0 0 0 4 12.5h1.25v-2zm4.25 0v2h2.5v-2zm4 0v2H12a1.5 1.5 0 0 0 1.5-1.5v-.5zM13.5 9h-2.75V7h2.75zM9.25 9h-2.5V7h2.5zM4 2a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutHeaderCellsLarge(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 3.5h8A1.5 1.5 0 0 1 13.5 5v.5h-11V5A1.5 1.5 0 0 1 4 3.5m-1.5 7v.5A1.5 1.5 0 0 0 4 12.5h3.25v-2zm0-1.5h4.75V7H2.5zm6.25 1.5v2H12a1.5 1.5 0 0 0 1.5-1.5v-.5zM13.5 9H8.75V7h4.75zM1 5a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutHeaderCellsLargeFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 6H8.75v2.5h4.75zM7.25 6H2.5v2.5h4.75zM1 6V5a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3zm7.75 4h4.75v1a1.5 1.5 0 0 1-1.5 1.5H8.75zM2.5 10h4.75v2.5H4A1.5 1.5 0 0 1 2.5 11z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutHeaderCellsLargeLetterD(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 3.5h8A1.5 1.5 0 0 1 13.5 5v.5h-11V5A1.5 1.5 0 0 1 4 3.5m-1.5 7v.5A1.5 1.5 0 0 0 4 12.5h3.25v-2zm0-1.5h4.75V7H2.5zm6.25 1.5V14H4a3 3 0 0 1-3-3V5a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v2H8.75zm3.67-1.995a3.25 3.25 0 0 1 0 6.5h-1.17a.75.75 0 0 1-.75-.75v-5a.75.75 0 0 1 .75-.75zm0 5a1.75 1.75 0 1 0 0-3.5H12v3.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutHeaderCellsLargeThunderbolt(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 3.5h8A1.5 1.5 0 0 1 13.5 5v.5h-11V5A1.5 1.5 0 0 1 4 3.5m-1.5 7v.5A1.5 1.5 0 0 0 4 12.5h3.25v-2zm0-1.5h4.75V7H2.5zm6.25 1.5V14H4a3 3 0 0 1-3-3V5a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v2H8.75zm4.591-.91A.75.75 0 0 0 12 8.92l-1.25 2.5a.75.75 0 0 0 .67 1.085h1.287L12 13.92a.75.75 0 0 0 1.341.67l1.25-2.5a.75.75 0 0 0-.67-1.085h-1.287z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutHeaderColumns(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 3.5H4A1.5 1.5 0 0 0 2.5 5v.5h11V5A1.5 1.5 0 0 0 12 3.5M2.5 11V7h4.75v5.5H4A1.5 1.5 0 0 1 2.5 11m6.25 1.5H12a1.5 1.5 0 0 0 1.5-1.5V7H8.75zM4 2a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutHeaderCursor(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 3h8a1.5 1.5 0 0 1 1.5 1.5V5h-11v-.5A1.5 1.5 0 0 1 4 3m9.5 5.25V6.5h-11V10A1.5 1.5 0 0 0 4 11.5h2.75a.75.75 0 0 1 0 1.5H4a3 3 0 0 1-3-3V4.5a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v3.75a.75.75 0 0 1-1.5 0m-3.272-.078A.75.75 0 0 0 9 8.75v5.543a.75.75 0 0 0 1.335.47l.898-1.12l.846 1.692a.75.75 0 1 0 1.342-.67l-.827-1.654H14a.75.75 0 0 0 .478-1.328z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutHeaderSideContent(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 3.5H4A1.5 1.5 0 0 0 2.5 5v.5h11V5A1.5 1.5 0 0 0 12 3.5M2.5 11V7h2v5.5H4A1.5 1.5 0 0 1 2.5 11M6 12.5h6a1.5 1.5 0 0 0 1.5-1.5V7H6zM4 2a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutList(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 3.5h1.25v2H2.5V5A1.5 1.5 0 0 1 4 3.5m2.75 2v-2H12A1.5 1.5 0 0 1 13.5 5v.5zM2.5 7h2.75v2H2.5zm0 3.5v.5A1.5 1.5 0 0 0 4 12.5h1.25v-2zm4.25 0v2H12a1.5 1.5 0 0 0 1.5-1.5v-.5zM13.5 9V7H6.75v2zM1 5a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutRows(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 3.5h8A1.5 1.5 0 0 1 13.5 5v2.25h-11V5A1.5 1.5 0 0 1 4 3.5M2.5 8.75V11A1.5 1.5 0 0 0 4 12.5h8a1.5 1.5 0 0 0 1.5-1.5V8.75zM1 5a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutRowsThree(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 3.5H4A1.5 1.5 0 0 0 2.5 5v.5h11V5A1.5 1.5 0 0 0 12 3.5M2.5 9V7h11v2zm0 1.5v.5A1.5 1.5 0 0 0 4 12.5h8a1.5 1.5 0 0 0 1.5-1.5v-.5zM4 2a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutSideContent(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 3.5h6A1.5 1.5 0 0 1 13.5 5v6a1.5 1.5 0 0 1-1.5 1.5H6zm-1.5 0H4A1.5 1.5 0 0 0 2.5 5v6A1.5 1.5 0 0 0 4 12.5h.5zM1 5a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutSideContentLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 3.5h6A1.5 1.5 0 0 1 13.5 5v6a1.5 1.5 0 0 1-1.5 1.5H6zm-1.5 0H4A1.5 1.5 0 0 0 2.5 5v6A1.5 1.5 0 0 0 4 12.5h.5zM1 5a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutSideContentRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 3.5H4A1.5 1.5 0 0 0 2.5 5v6A1.5 1.5 0 0 0 4 12.5h6zm1.5 0h.5A1.5 1.5 0 0 1 13.5 5v6a1.5 1.5 0 0 1-1.5 1.5h-.5zM15 5a3 3 0 0 0-3-3H4a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutSplitColumns(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiLayoutSplitColumns0)"><path fill="currentColor" fill-rule="evenodd" d="M5 12.5H3a1 1 0 0 1-1-1v-7a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1M3 14h2a2.5 2.5 0 0 0 2.5-2.5v-7A2.5 2.5 0 0 0 5 2H3A2.5 2.5 0 0 0 .5 4.5v7A2.5 2.5 0 0 0 3 14m10-1.5h-2a1 1 0 0 1-1-1v-7a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1M11 14h2a2.5 2.5 0 0 0 2.5-2.5v-7A2.5 2.5 0 0 0 13 2h-2a2.5 2.5 0 0 0-2.5 2.5v7A2.5 2.5 0 0 0 11 14" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiLayoutSplitColumns0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutSplitColumnsThree(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiLayoutSplitColumns30)"><path fill="currentColor" fill-rule="evenodd" d="M2.5 12.5H3a.5.5 0 0 0 .5-.5V4a.5.5 0 0 0-.5-.5h-.5A.5.5 0 0 0 2 4v8a.5.5 0 0 0 .5.5M3 14h-.5a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2H3a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2m4.75-1.5h.5a.5.5 0 0 0 .5-.5V4a.5.5 0 0 0-.5-.5h-.5a.5.5 0 0 0-.5.5v8a.5.5 0 0 0 .5.5m.5 1.5h-.5a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h.5a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2m5.25-1.5H13a.5.5 0 0 1-.5-.5V4a.5.5 0 0 1 .5-.5h.5a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-.5.5M13 14h.5a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H13a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiLayoutSplitColumns30"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutSplitRows(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 4v1a1 1 0 0 1-1 1h-9a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1M15 5V4a2.5 2.5 0 0 0-2.5-2.5h-9A2.5 2.5 0 0 0 1 4v1a2.5 2.5 0 0 0 2.5 2.5h9A2.5 2.5 0 0 0 15 5m-1.5 6v1a1 1 0 0 1-1 1h-9a1 1 0 0 1-1-1v-1a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1m1.5 1v-1a2.5 2.5 0 0 0-2.5-2.5h-9A2.5 2.5 0 0 0 1 11v1a2.5 2.5 0 0 0 2.5 2.5h9A2.5 2.5 0 0 0 15 12" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutSplitSideContentLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiLayoutSplitSideContentLeft0)"><path fill="currentColor" fill-rule="evenodd" d="M3 12.5h-.5A.5.5 0 0 1 2 12V4a.5.5 0 0 1 .5-.5H3a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-.5.5M2.5 14H3a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2h-.5a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2M13 12.5H8.5a1 1 0 0 1-1-1v-7a1 1 0 0 1 1-1H13a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1M8.5 14H13a2.5 2.5 0 0 0 2.5-2.5v-7A2.5 2.5 0 0 0 13 2H8.5A2.5 2.5 0 0 0 6 4.5v7A2.5 2.5 0 0 0 8.5 14" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiLayoutSplitSideContentLeft0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutSplitSideContentRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiLayoutSplitSideContentRight0)"><path fill="currentColor" fill-rule="evenodd" d="M3 12.5h4.5a1 1 0 0 0 1-1v-7a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1M7.5 14H3a2.5 2.5 0 0 1-2.5-2.5v-7A2.5 2.5 0 0 1 3 2h4.5A2.5 2.5 0 0 1 10 4.5v7A2.5 2.5 0 0 1 7.5 14m5.5-1.5h.5a.5.5 0 0 0 .5-.5V4a.5.5 0 0 0-.5-.5H13a.5.5 0 0 0-.5.5v8a.5.5 0 0 0 .5.5m.5 1.5H13a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h.5a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiLayoutSplitSideContentRight0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutTabs(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.999 3.5h.242c.559 0 1.049.37 1.202.907l.171.598A2.75 2.75 0 0 0 8.26 7h5.24v4a1.5 1.5 0 0 1-1.5 1.5H4A1.5 1.5 0 0 1 2.5 11V5a1.5 1.5 0 0 1 1.499-1.5m9.501 2h-1.242a1.25 1.25 0 0 1-1.201-.907l-.171-.598a2.758 2.758 0 0 0-.195-.495H12A1.5 1.5 0 0 1 13.5 5zM8.241 2H4a3.057 3.057 0 0 0-.457.035A2.99 2.99 0 0 0 2 2.764A2.99 2.99 0 0 0 1 5v6a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3zM7 3.5h1.242c.559 0 1.049.37 1.202.907l.171.598c.05.174.115.339.195.495H8.26a1.25 1.25 0 0 1-1.202-.907l-.171-.598A2.753 2.753 0 0 0 6.69 3.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterGroup(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 2.75a.75.75 0 0 0-1.5 0v7.506a.75.75 0 0 0 1.388.395c.38.223.834.349 1.362.349C13.045 11 14 9.545 14 7.75s-.955-3.25-2.75-3.25a2.69 2.69 0 0 0-1.25.287zm0 5c0 .65.175 1.109.38 1.37c.174.22.424.38.87.38s.696-.16.87-.38c.205-.261.38-.72.38-1.37c0-.65-.175-1.109-.38-1.37c-.174-.22-.424-.38-.87-.38s-.696.16-.87.38c-.205.261-.38.72-.38 1.37M3.58 6.339a2.26 2.26 0 0 1 .432-.222c.132-.046.321-.086.634-.086c.447 0 .743.143.938.365c.12.138.23.338.302.622h-1.31c-1.262 0-2.332.849-2.326 1.998c.003.478.168.988.561 1.381c.4.4.963.61 1.627.61c.53 0 1.091-.16 1.584-.472a.75.75 0 0 0 1.447-.279V7.768c0-.946-.237-1.766-.756-2.36c-.53-.606-1.271-.877-2.067-.877c-.452 0-.811.06-1.125.169c-.293.101-.52.237-.689.338l-.022.013a.75.75 0 1 0 .77 1.288m.17 2.67c0-.095.154-.49.827-.49h1.258c-.268.677-.89.987-1.397.987c-.345 0-.499-.102-.566-.169a.458.458 0 0 1-.122-.329ZM0 11.75A2.25 2.25 0 0 0 2.25 14h11.5A2.25 2.25 0 0 0 16 11.75v-1a.75.75 0 0 0-1.5 0v1a.75.75 0 0 1-.75.75H2.25a.75.75 0 0 1-.75-.75v-1a.75.75 0 0 0-1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterM(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.062 2c.454 0 .876.234 1.115.62L8 7.162l2.823-4.542a1.312 1.312 0 0 1 2.427.692v9.938a.75.75 0 0 1-1.5 0V3.97L8.637 8.98a.75.75 0 0 1-1.274 0L4.25 3.97v9.28a.75.75 0 0 1-1.5 0V3.312C2.75 2.588 3.338 2 4.062 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LifeRing(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 13.5c1.248 0 2.4-.416 3.323-1.117l-1.436-1.435A3.483 3.483 0 0 1 8 11.5a3.483 3.483 0 0 1-1.887-.552l-1.436 1.435A5.476 5.476 0 0 0 8 13.5M5.052 9.887A3.484 3.484 0 0 1 4.5 8c0-.695.203-1.343.552-1.887L3.617 4.677A5.476 5.476 0 0 0 2.5 8c0 1.248.416 2.4 1.117 3.323zm1.06-4.835L4.678 3.617A5.476 5.476 0 0 1 8 2.5c1.248 0 2.4.416 3.323 1.117L9.887 5.052A3.484 3.484 0 0 0 8 4.5c-.695 0-1.343.203-1.887.552Zm4.836 1.06c.35.545.552 1.193.552 1.888c0 .695-.203 1.343-.552 1.887l1.435 1.436A5.476 5.476 0 0 0 13.5 8c0-1.248-.416-2.4-1.117-3.323zM8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14m2-7a2 2 0 1 1-4 0a2 2 0 0 1 4 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.47 6.53a.75.75 0 0 1 1.06 1.061l-.727.727a2.743 2.743 0 0 0 3.879 3.879l.727-.727a.75.75 0 0 1 1.06 1.06l-.726.727a4.243 4.243 0 0 1-6-6zm8 1.879a.75.75 0 0 0 1.06 1.06l.727-.726a4.243 4.243 0 0 0-6-6l-.727.727a.75.75 0 0 0 1.061 1.06l.727-.727a2.743 2.743 0 0 1 3.879 3.879zm-.94-1.879a.75.75 0 1 0-1.06-1.06l-4 4a.75.75 0 1 0 1.06 1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkSlash(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m8 9.06l4.97 4.97a.75.75 0 1 0 1.06-1.06l-11-11a.75.75 0 0 0-1.06 1.06L6.94 8L5.47 9.47a.75.75 0 1 0 1.06 1.06zm3.54-.722L12.601 9.4l.656-.656a4.243 4.243 0 0 0-6-6l-.656.656l1.06 1.06l.657-.656a2.743 2.743 0 0 1 3.879 3.879l-.657.656ZM9.47 5.47l-.4.399l1.061 1.06l.4-.399A.75.75 0 1 0 9.47 5.47M3.22 6.78a.75.75 0 0 1 1.06 1.061l-.477.477a2.743 2.743 0 0 0 3.879 3.879l.477-.477a.75.75 0 0 1 1.06 1.06l-.476.477a4.243 4.243 0 0 1-6-6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListCheck(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.28 2.22a.75.75 0 0 1 0 1.06l-2 2a.75.75 0 0 1-1.06 0l-1-1a.75.75 0 0 1 1.06-1.06l.47.47l1.47-1.47a.75.75 0 0 1 1.06 0M6.5 3.75c0 .414.336.75.75.75h7a.75.75 0 0 0 0-1.5h-7a.75.75 0 0 0-.75.75m.75 3.5a.75.75 0 0 0 0 1.5h7a.75.75 0 0 0 0-1.5zm-1.97.03a.75.75 0 0 0-1.06-1.06L2.75 7.69l-.47-.47a.75.75 0 0 0-1.06 1.06l1 1a.75.75 0 0 0 1.06 0zm0 3.19a.75.75 0 0 1 0 1.06l-2 2a.75.75 0 0 1-1.06 0l-1-1a.75.75 0 1 1 1.06-1.06l.47.47l1.47-1.47a.75.75 0 0 1 1.06 0m1.97 1.03a.75.75 0 0 0 0 1.5h7a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListCheckLock(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.28 2.22a.75.75 0 0 1 0 1.06l-2 2a.75.75 0 0 1-1.06 0l-1-1a.75.75 0 0 1 1.06-1.06l.47.47l1.47-1.47a.75.75 0 0 1 1.06 0M7.25 3h7a.75.75 0 0 1 0 1.5h-7a.75.75 0 0 1 0-1.5M5.28 7.28a.75.75 0 0 0-1.06-1.06L2.75 7.69l-.47-.47a.75.75 0 0 0-1.06 1.06l1 1a.75.75 0 0 0 1.06 0zm0 3.19a.75.75 0 0 1 0 1.06l-2 2a.75.75 0 0 1-1.06 0l-1-1a.75.75 0 1 1 1.06-1.06l.47.47l1.47-1.47a.75.75 0 0 1 1.06 0M8.4 9h.1a2.5 2.5 0 0 1 5 0h.1a1.4 1.4 0 0 1 1.4 1.4v2.2a1.4 1.4 0 0 1-1.4 1.4H8.4A1.4 1.4 0 0 1 7 12.6v-2.2A1.4 1.4 0 0 1 8.4 9M11 8a1 1 0 0 1 1 1h-2a1 1 0 0 1 1-1m-2.5 2.5v2h5v-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListOl(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.21 2.5a.5.5 0 0 0-.796-.403l-.71.524a.5.5 0 0 0 .507.856V5a.5.5 0 0 0 1 0zm2.54.5a.75.75 0 1 0 0 1.5h8.5a.75.75 0 0 0 0-1.5zm0 8.5a.75.75 0 0 0 0 1.5h8.5a.75.75 0 0 0 0-1.5zM5 8a.75.75 0 0 1 .75-.75h8.5a.75.75 0 0 1 0 1.5h-8.5A.75.75 0 0 1 5 8m-2.596-.67c-.07.051-.087.097-.09.116a.5.5 0 0 1-.988-.149c.051-.345.26-.61.495-.779c.236-.169.531-.268.83-.268c.345 0 .666.109.902.343c.236.234.337.543.337.85c0 .51-.321.855-.76 1.278a.495.495 0 0 1-.02.018l-.013.011h.493a.5.5 0 0 1 0 1H1.75a.5.5 0 0 1-.326-.879l1.022-.88c.223-.216.338-.344.398-.434c.046-.069.046-.092.046-.111v-.002c0-.09-.027-.127-.041-.141c-.013-.013-.062-.053-.199-.053a.439.439 0 0 0-.246.08M1.9 10.5a.5.5 0 0 0 0 1h.38l-.15.18a.5.5 0 0 0 .383.82c.259 0 .365.077.407.122c.051.054.08.133.08.218a.02.02 0 0 1-.004.014a.188.188 0 0 1-.057.052a.702.702 0 0 1-.371.094a.545.545 0 0 1-.258-.057c-.059-.032-.075-.061-.076-.064a.5.5 0 0 0-.91.416c.213.464.728.705 1.244.705c.319 0 .651-.082.92-.258c.275-.181.512-.487.512-.902c0-.293-.096-.634-.355-.906a1.28 1.28 0 0 0-.252-.206l.34-.407a.5.5 0 0 0-.383-.821z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListTimeline(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.75.5a.75.75 0 0 0-.75.75v13.5a.75.75 0 0 0 1.5 0V1.25A.75.75 0 0 0 5.75.5m-2 3.5a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0m4 0a.75.75 0 0 1 .75-.75h5.75a.75.75 0 0 1 0 1.5H8.5A.75.75 0 0 1 7.75 4m0 8a.75.75 0 0 1 .75-.75h5.75a.75.75 0 0 1 0 1.5H8.5a.75.75 0 0 1-.75-.75m.75-4.75a.75.75 0 1 0 0 1.5h5.75a.75.75 0 0 0 0-1.5zm-6 2a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5M3.75 12a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListUl(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 5a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5m3.25-2a.75.75 0 0 0 0 1.5h8.5a.75.75 0 0 0 0-1.5zm0 8.5a.75.75 0 0 0 0 1.5h8.5a.75.75 0 0 0 0-1.5zM5 8a.75.75 0 0 1 .75-.75h8.5a.75.75 0 0 1 0 1.5h-8.5A.75.75 0 0 1 5 8M3.75 8a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0M2.5 13.5a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.5 6V5a2.5 2.5 0 0 0-5 0v1zM4 5v1a3 3 0 0 0-3 3v3a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3V5a4 4 0 0 0-8 0m6.5 2.5H12A1.5 1.5 0 0 1 13.5 9v3a1.5 1.5 0 0 1-1.5 1.5H4A1.5 1.5 0 0 1 2.5 12V9A1.5 1.5 0 0 1 4 7.5zm-1.75 2a.75.75 0 0 0-1.5 0v2a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOpen(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.5 5a2.5 2.5 0 0 0-4.532-1.456c-.242.336-.66.559-1.052.428c-.393-.131-.611-.56-.41-.922A4 4 0 0 1 12 5v1h.001a3 3 0 0 1 3 3v3a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3V9a3 3 0 0 1 3-3h6.5zm.75 2.5H4A1.5 1.5 0 0 0 2.5 9v3A1.5 1.5 0 0 0 4 13.5h8a1.5 1.5 0 0 0 1.5-1.5V9A1.5 1.5 0 0 0 12 7.5zM8 8.75a.75.75 0 0 1 .75.75v2a.75.75 0 0 1-1.5 0v-2A.75.75 0 0 1 8 8.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoAcrobat(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiLogoAcrobat0)"><path fill="currentColor" fill-rule="evenodd" d="M7.363 1.5c-.094.007-.377.122-.377.622c0 .593.12 1.248.337 1.927c.209-.799.319-1.463.319-1.927c0-.359-.097-.513-.144-.565a.155.155 0 0 0-.135-.057m.885 4.637c.061-.175.12-.347.174-.518c.431-1.326.72-2.587.72-3.497c0-.625-.173-1.173-.53-1.57A1.653 1.653 0 0 0 7.265.004c-.984.064-1.779.933-1.779 2.118c0 1.314.44 2.758 1.126 4.13a31.082 31.082 0 0 1-1.465 3.305c-.9.302-1.745.648-2.475 1.025c-.714.368-1.352.783-1.822 1.24c-.46.448-.85 1.026-.85 1.714c0 .465.157.91.488 1.243c.33.332.776.491 1.242.491c.541 0 1.038-.242 1.458-.547c.428-.311.844-.738 1.24-1.225c.604-.744 1.203-1.683 1.758-2.699a24.497 24.497 0 0 1 3.143-.739a11.26 11.26 0 0 0 1.606 1.369c.925.639 1.945 1.086 2.943 1.086c.507 0 1.034-.126 1.45-.46c.438-.353.672-.871.672-1.47c0-.758-.412-1.313-.983-1.66c-.535-.327-1.231-.49-1.958-.563c-.925-.092-2.02-.047-3.167.1a14.055 14.055 0 0 1-1.644-2.325m-.675 1.76a32.96 32.96 0 0 1-.481 1.095a26.209 26.209 0 0 1 1.093-.246a15.666 15.666 0 0 1-.612-.849m3.734 1.938c.16.13.321.25.481.36c.802.554 1.521.82 2.09.82c.277 0 .436-.07.51-.13c.052-.041.112-.113.112-.3c0-.129-.04-.244-.263-.38c-.258-.156-.695-.288-1.327-.35a10.518 10.518 0 0 0-1.603-.02m-7.374 1.803c-.2.09-.391.183-.574.277c-.645.333-1.14.668-1.463.982c-.334.326-.396.54-.396.639c0 .126.039.172.05.184c.013.012.057.05.18.05c.095 0 .284-.047.576-.26c.285-.207.608-.528.957-.957c.221-.273.446-.58.67-.915" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiLogoAcrobat0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoDocker(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiLogoDocker0)"><path fill="currentColor" fill-rule="evenodd" d="M5.5 0c-.69 0-1.25.56-1.25 1.25v2h-2C1.56 3.25 1 3.81 1 4.5v2.382l-.06.057L1 7v.002a.847.847 0 0 0-.062-.06l-.002.001l-.003.004l-.006.006l-.014.014a1.76 1.76 0 0 0-.092.105c-.045.055-.095.12-.147.198a3.326 3.326 0 0 0-.317.611C.154 8.39 0 9.08 0 10v.009a4.897 4.897 0 0 0 1.455 3.46C2.434 14.437 3.825 15 5.5 15c2.071 0 3.785-.59 5.107-1.637c.997-.79 1.713-1.796 2.195-2.886a5.28 5.28 0 0 0 .99-.153c.819-.218 1.579-.711 2.05-1.653a1.5 1.5 0 0 0 .081-1.146c-.321-.962-1.12-1.636-2.071-1.877c-.24-.95-.915-1.75-1.877-2.071a1.5 1.5 0 0 0-1.146.081c-1.092.546-1.502 1.583-1.668 2.329c-.038.17-.067.343-.089.513H9V1.25C9 .56 8.44 0 7.75 0zm2 6.5V4.75H5.75V6.5zm-3.25 0V4.75H2.5V6.5zM2 8s-.5.5-.5 2c.01 1.738 1.344 3.5 4 3.5c3.55 0 5.5-2 6.25-4.5C13 9 14 9 14.5 8a1.374 1.374 0 0 0-.655-.756c-.434-.237-1.014-.284-1.595.006c.29-.58.243-1.161.006-1.595A1.373 1.373 0 0 0 11.5 5c-1 .5-1 2-1 3zm3.75-4.75H7.5V1.5H5.75zM4.25 9a.75.75 0 0 1 .75.75v1a.75.75 0 0 1-1.5 0v-1A.75.75 0 0 1 4.25 9" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiLogoDocker0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoDrawIo(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 12.5h2a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 0-.5-.5h-2a.5.5 0 0 0-.5.5v2a.5.5 0 0 0 .5.5m2 1.5a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2h-1.85l-.847-1.482A2 2 0 0 0 11 5V3a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v2c0 .607.27 1.151.697 1.518L4.85 8H3a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-2c0-.607-.27-1.151-.697-1.518L7.15 7h1.7l.848 1.482C9.27 8.85 9 9.392 9 10v2a2 2 0 0 0 2 2zM3 12.5h2a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 0-.5-.5H3a.5.5 0 0 0-.5.5v2a.5.5 0 0 0 .5.5m6-7H7a.5.5 0 0 1-.5-.5V3a.5.5 0 0 1 .5-.5h2a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoFacebook(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.5 13.478a5.5 5.5 0 1 0-1.5-.069V9.75H5.75a.75.75 0 1 1 0-1.5H7V7.24c0-.884.262-1.568.722-2.032c.46-.464 1.121-.708 1.922-.708c.273 0 .612.04.948.213a.75.75 0 0 1-.685 1.334A.57.57 0 0 0 9.644 6c-.493 0-.737.144-.857.265c-.12.12-.287.39-.287.975v1.01h1.25a.75.75 0 0 1 0 1.5H8.5zM8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoGitlab(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.357 7H6.643l1.08 6.267a.28.28 0 0 0 .554 0zM4.511 3.465l-.153-.691a.35.35 0 0 0-.684.007L2.522 8.537a1.144 1.144 0 0 0 .39 1.103l3.11 2.593l-1.51-8.768Zm5.466 8.768l3.111-2.593a1.144 1.144 0 0 0 .39-1.103l-1.152-5.756a.35.35 0 0 0-.684-.007l-.153.692zm-8.026-1.44l4.493 3.743a2 2 0 0 0 1.28.464h.552a2 2 0 0 0 1.28-.464l4.493-3.743a2.644 2.644 0 0 0 .9-2.55l-1.152-5.756a1.85 1.85 0 0 0-3.619-.039L9.5 5.5h-3l-.678-3.052a1.85 1.85 0 0 0-3.62.039l-1.15 5.756a2.644 2.644 0 0 0 .9 2.55Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoLinux(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.854 3.47a2.711 2.711 0 0 1 1.934.102l.618.275c.072.032.14.069.205.11c.003-.126.01-.251.023-.377l.05-.495a1.441 1.441 0 1 0-2.868 0l.038.384Zm3.243 2.725a1.61 1.61 0 0 1-.423.433L7.95 7.835a1 1 0 0 1-1.468-.372l-.23-.46l-.99 2.226a3.037 3.037 0 0 0-.217 1.758c.46.333.81.817.968 1.393l.123.451c.52.419 1.181.669 1.901.669h.54c.036 0 .07 0 .105-.002l.305-1.118a2.564 2.564 0 0 1 2.98-1.829a3.423 3.423 0 0 0-.605-2.463l-.84-1.177a4.897 4.897 0 0 1-.425-.716m3.242 5.135a4.93 4.93 0 0 0-.756-4.114l-.841-1.177a3.392 3.392 0 0 1-.615-2.31l.05-.495a2.941 2.941 0 1 0-5.854 0l.08.8a3.668 3.668 0 0 1-.298 1.854L3.891 8.62a4.537 4.537 0 0 0-.39 1.88A2.539 2.539 0 0 0 1 13.027c0 1.16.79 2.17 1.914 2.452l1.92.48a1.396 1.396 0 0 0 1.727-1.204c.463.159.96.245 1.476.245h.451a1.396 1.396 0 0 0 1.679.958l1.919-.48A2.527 2.527 0 0 0 14 13.027c0-.653-.25-1.248-.66-1.697ZM8.18 4.943l.617.275c.03.013.04.025.044.032c.008.01.016.029.018.053a.115.115 0 0 1-.007.056c-.004.008-.01.022-.037.04l-1.25.875l-.549-1.097a.127.127 0 0 1 .048-.166a1.211 1.211 0 0 1 1.115-.068Zm-4.901 9.08l1.747.437l-.46-1.685A1.064 1.064 0 0 0 3.542 12c-.578 0-1.041.47-1.041 1.027c0 .471.32.882.778.996m6.697.437l1.747-.437c.457-.114.778-.525.778-.996c0-.557-.463-1.027-1.04-1.027c-.482 0-.903.324-1.026.775l-.46 1.685Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoMacos(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.063 3.5H12A1.5 1.5 0 0 1 13.5 5v6a1.5 1.5 0 0 1-1.5 1.5h-1.441a.741.741 0 0 0-.029-.03c-.75-.75-.78-1.425-.78-3.22A.75.75 0 0 0 9 8.5H7.753c.018-1.895.162-3.441 1.31-5m-1.777 0H4A1.5 1.5 0 0 0 2.5 5v6A1.5 1.5 0 0 0 4 12.5h4.714c-.38-.76-.45-1.574-.462-2.5H7a.75.75 0 0 1-.75-.75v-.07c0-1.89 0-3.791 1.036-5.68M1 5a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3zm9.25 2.25a.75.75 0 0 0 1.5 0v-1a.75.75 0 0 0-1.5 0zM4.75 8A.75.75 0 0 1 4 7.25v-1a.75.75 0 0 1 1.5 0v1a.75.75 0 0 1-.75.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoMarkdown(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.443 4.699A1.294 1.294 0 0 0 1 5.294v5.956a.75.75 0 0 0 1.5 0V6.139l1.834 3.54a.75.75 0 0 0 1.332 0L7.5 6.138v5.111a.75.75 0 0 0 1.5 0V5.294A1.294 1.294 0 0 0 6.557 4.7L5 7.703L3.443 4.7Zm9.807.051a.75.75 0 0 0-1.5 0v4.69l-.47-.47a.75.75 0 1 0-1.06 1.06l1.75 1.75a.75.75 0 0 0 1.06 0l1.75-1.75a.75.75 0 1 0-1.06-1.06l-.47.47z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoMermaid(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 12.5A1.5 1.5 0 0 1 10 14H6a1.5 1.5 0 0 1-1.5-1.5v-.823a3.114 3.114 0 0 0-1.35-2.566A6.114 6.114 0 0 1 .5 4.073V3A1.5 1.5 0 0 1 2 1.5h.666A6.425 6.425 0 0 1 8 4.343A6.425 6.425 0 0 1 13.334 1.5H14A1.5 1.5 0 0 1 15.5 3v1.073a6.114 6.114 0 0 1-2.65 5.038a3.114 3.114 0 0 0-1.35 2.566zm-8-9.43a4.924 4.924 0 0 1 3.738 3.025c.275.688 1.249.688 1.524 0A4.924 4.924 0 0 1 13.334 3H14v1.073a4.614 4.614 0 0 1-2 3.802c-1.252.86-2 2.283-2 3.802v.823H6v-.823c0-1.52-.748-2.941-2-3.802a4.614 4.614 0 0 1-2-3.802V3h.666c.283 0 .562.024.834.07" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoNotion(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 2.5h4.98a2.5 2.5 0 0 1 1.66.631l.977.869h-5.93a1.511 1.511 0 0 1-1.069-.443l-.092-.091l-.915-.915c.124-.033.255-.051.389-.051M2.55 3.611A1.56 1.56 0 0 0 2.5 4v4.98a2.5 2.5 0 0 0 .631 1.66l.869.977v-5.8a1.826 1.826 0 0 0-.535-1.291zM7 5.5A1.5 1.5 0 0 0 5.5 7v5A1.5 1.5 0 0 0 7 13.5h5a1.5 1.5 0 0 0 1.5-1.5V7A1.5 1.5 0 0 0 12 5.5zM7 15h-.653a3 3 0 0 1-2.242-1.007L2.01 11.637A4 4 0 0 1 1 8.979V4a3 3 0 0 1 3-3h4.98a4 4 0 0 1 2.657 1.01l2.356 2.095A3 3 0 0 1 15 6.347V12a3 3 0 0 1-3 3zm.75-3a.75.75 0 0 1-.75-.75V8.08a1.08 1.08 0 0 1 1.96-.627l1.54 2.16V7.75a.75.75 0 0 1 1.5 0v3.168a1.082 1.082 0 0 1-1.963.628L8.5 9.392v1.858a.75.75 0 0 1-.75.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoOsi(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 2.5a5.5 5.5 0 0 0-2.977 10.126l.79-2.571a3 3 0 1 1 4.372 0l.792 2.57A5.5 5.5 0 0 0 8 2.5M1 8a7 7 0 1 1 9.832 6.403a.75.75 0 0 1-1.02-.465l-1.2-3.901a.75.75 0 0 1 .273-.826a1.5 1.5 0 1 0-1.77 0a.75.75 0 0 1 .274.826l-1.2 3.901a.75.75 0 0 1-1.021.465A7.001 7.001 0 0 1 1 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoPython(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 5.5h3.752a.75.75 0 0 0 0-1.5H5.5A1.5 1.5 0 0 1 7 2.5h2A1.5 1.5 0 0 1 10.5 4v2c0 .69-.56 1.25-1.25 1.25h-2.5A2.75 2.75 0 0 0 4 10v.5A1.5 1.5 0 0 1 2.5 9V7A1.5 1.5 0 0 1 4 5.5m1.5 5V12A1.5 1.5 0 0 0 7 13.5h2a1.5 1.5 0 0 0 1.5-1.5H8.25a.75.75 0 0 1 0-1.5H12A1.5 1.5 0 0 0 13.5 9V7A1.5 1.5 0 0 0 12 5.5V6a2.75 2.75 0 0 1-2.75 2.75h-2.5c-.69 0-1.25.56-1.25 1.25zM7 1a3 3 0 0 0-3 3a3 3 0 0 0-3 3v2a3 3 0 0 0 3 3a3 3 0 0 0 3 3h2a3 3 0 0 0 3-3a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3a3 3 0 0 0-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoStackOverflow(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.849.33a.75.75 0 0 0-1.115 1.003l3.68 4.088a.75.75 0 1 0 1.115-1.004zM7.6 2.443a.75.75 0 0 1 1.048-.166L13.1 5.51a.75.75 0 1 1-.882 1.213L7.767 3.49a.75.75 0 0 1-.166-1.047Zm-.589 2.424a.75.75 0 1 0-.61 1.37l5.025 2.238a.75.75 0 1 0 .61-1.37L7.012 4.866ZM4.82 8.11a.75.75 0 0 1 .89-.577l5.38 1.143a.75.75 0 0 1-.313 1.467L5.396 9a.75.75 0 0 1-.577-.89ZM2 10.25a.75.75 0 0 1 1.5 0V12A1.5 1.5 0 0 0 5 13.5h6a1.5 1.5 0 0 0 1.5-1.5v-1.75a.75.75 0 0 1 1.5 0V12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3zm3.25.25a.75.75 0 0 0 0 1.5h5.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoTelegram(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.254 8.521L9.61 5.86a.75.75 0 0 1 .782 1.28L6.586 9.465L9.77 12.65a1.199 1.199 0 0 0 1.973-.433l2.692-7.308a1.045 1.045 0 0 0-.98-1.408h-.105c-.1 0-.201.013-.298.04L2.022 6.509a.707.707 0 0 0 .046 1.375zm-3.48.834L5 10l3.71 3.71a2.7 2.7 0 0 0 4.44-.976l2.693-7.308A2.544 2.544 0 0 0 13.454 2h-.104c-.232 0-.464.03-.688.091l-11.03 2.97a2.207 2.207 0 0 0 .142 4.294" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoUbuntu(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.034 3.902a1.5 1.5 0 1 0-2.03-1.293A5.738 5.738 0 0 0 9 2.25a5.75 5.75 0 0 0-1.681.25c-.427.13-.58.63-.356 1.017c.195.333.607.46.982.365a4.258 4.258 0 0 1 5.144 2.954c.105.372.423.664.81.664c.446 0 .801-.384.699-.82a5.742 5.742 0 0 0-1.564-2.778m-5.09 8.216c-.373-.096-.786.032-.98.365c-.226.386-.072.887.355 1.017A5.75 5.75 0 0 0 9 13.75c.705 0 1.38-.127 2.004-.359a1.5 1.5 0 1 0 2.03-1.293a5.742 5.742 0 0 0 1.564-2.779c.102-.435-.253-.819-.7-.819c-.386 0-.704.292-.81.664a4.252 4.252 0 0 1-5.143 2.954ZM5.937 5.055c.268-.278.359-.7.165-1.033c-.226-.386-.737-.499-1.061-.19A5.742 5.742 0 0 0 3.38 6.784a1.5 1.5 0 1 0 0 2.432a5.742 5.742 0 0 0 1.66 2.953c.325.308.836.195 1.062-.191c.194-.333.103-.755-.165-1.033A4.236 4.236 0 0 1 4.75 8c0-1.143.452-2.181 1.186-2.945Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoWindows(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m11.788 2.974l-3.038.434V7.25h4.75V4.46a1.5 1.5 0 0 0-1.712-1.486M13.5 8.75H8.75v3.842l3.038.434A1.5 1.5 0 0 0 13.5 11.54zm-6.25-1.5V3.622l-3.462.495A1.5 1.5 0 0 0 2.5 5.602V7.25zM2.5 8.75h4.75v3.628l-3.462-.495A1.5 1.5 0 0 1 2.5 10.398zm1.076-6.118A3 3 0 0 0 1 5.602v4.796a3 3 0 0 0 2.576 2.97l8 1.143A3 3 0 0 0 15 11.54V4.459a3 3 0 0 0-3.424-2.97z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoYandex(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M6.136 5.103a.75.75 0 0 0-1.272.795l2.044 3.27c.223.357.342.77.342 1.192v1.14a.75.75 0 0 0 1.5 0v-1.14a3.75 3.75 0 0 0-.57-1.987zm5 .795a.75.75 0 1 0-1.272-.795L8.77 6.853a.75.75 0 0 0 1.272.795z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoYandexCloud(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 8a5.5 5.5 0 0 1 5.492-5.5c1.318.032 1.88.796 2.02 1.48c.027.137.02.389-.027.713a7.242 7.242 0 0 1-.042.26l-1.759.308h-.001c-.838.144-1.537.432-2.068.933c-.535.504-.829 1.155-.98 1.887l-.444 2.146a.734.734 0 0 0-.009.043l-.02.098c-.034.168-.091.444-.132.725c-.047.325-.09.763-.025 1.154A5.489 5.489 0 0 1 2.5 8m5.501 7a7 7 0 0 0 .002-14h-.006a7 7 0 0 0 .001 14zm.007-1.5c-1.318-.032-1.88-.796-2.02-1.48c-.027-.137-.02-.389.027-.713a7.25 7.25 0 0 1 .042-.26l1.759-.308h.001c.838-.144 1.537-.432 2.068-.933c.535-.504.829-1.155.98-1.887l.444-2.146a.857.857 0 0 0 .009-.043c.005-.026.011-.059.02-.097c.034-.17.092-.446.132-.726c.047-.325.09-.763.024-1.154A5.5 5.5 0 0 1 8.008 13.5M9.62 6.532l-1.18.207h-.002c-.652.112-1.044.31-1.295.547c-.247.233-.43.57-.539 1.098L6.38 9.468l1.18-.207h.002c.652-.112 1.044-.31 1.295-.547c.247-.233.43-.57.539-1.098z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoYandexMessenger(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m3.585 11.965l-.716.148a1.863 1.863 0 0 1-1.792-3.036L2 8l-.923-1.077A1.863 1.863 0 0 1 2.87 3.887l.716.148l-.07-.303a1.822 1.822 0 0 1 2.64-2.014l8.293 4.466a2 2 0 0 1 1.052 1.76v.111a2 2 0 0 1-1.052 1.761l-8.293 4.466a1.822 1.822 0 0 1-2.64-2.014l.07-.303Zm1.617-.334l-.225.974a.322.322 0 0 0 .466.356l1.625-.875l3.895-1.647zm5.761-6.07L7.068 3.914l-1.625-.875a.322.322 0 0 0-.466.356l.225.974zM3.14 8.976L3.976 8l-.837-.976l-.923-1.077a.363.363 0 0 1 .349-.591l11.162 2.309a.342.342 0 0 1 0 .67l-11.162 2.31a.363.363 0 0 1-.349-.592l.923-1.077Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoYandexTracker(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.75 2.5a.25.25 0 0 0-.25.25v2.167c0 .138.112.25.25.25h2.417V2.5zm3.917 0v2.667h2.666V2.5zm4.166 0v2.667h2.417a.25.25 0 0 0 .25-.25V2.75a.25.25 0 0 0-.25-.25zm0 4.167h2.417A1.75 1.75 0 0 0 15 4.917V2.75A1.75 1.75 0 0 0 13.25 1H2.75A1.75 1.75 0 0 0 1 2.75v2.167c0 .966.784 1.75 1.75 1.75h2.417v6.583c0 .966.783 1.75 1.75 1.75h2.166a1.75 1.75 0 0 0 1.75-1.75zm-1.5 0H6.667v2.666h2.666zm0 4.166H6.667v2.417c0 .138.112.25.25.25h2.166a.25.25 0 0 0 .25-.25z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagicWand(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m11.1 5.894l-.412.903l.671.732l1.356 1.48l-1.994-.227l-.986-.113l-.49.864l-.987 1.747l-.4-1.967l-.198-.972l-.973-.198l-1.966-.4l1.746-.989l.864-.489l-.112-.986l-.227-1.994L8.47 4.641l.732.67l.903-.411l1.826-.832zM7.02 1.276l2.465 2.26l3.043-1.387c.842-.384 1.708.483 1.325 1.324l-1.387 3.043l2.259 2.465c.625.682.069 1.774-.85 1.67l-3.323-.38l-1.646 2.911c-.456.805-1.666.613-1.85-.293l-.667-3.277l-3.277-.666c-.906-.185-1.098-1.395-.293-1.85l2.91-1.647l-.378-3.322c-.105-.92.987-1.476 1.669-.85ZM5.53 11.53a.75.75 0 1 0-1.06-1.06l-3.5 3.5a.75.75 0 1 0 1.06 1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Magnet(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.28 3.841L3.925 5.197a4.864 4.864 0 0 0 6.878 6.878l1.356-1.355l-1.379-1.38l-1.305 1.306a2.914 2.914 0 0 1-4.121-4.122L6.659 5.22L5.28 3.84Zm1.061-1.06L7.72 4.158L8.879 3L7.5 1.621zM14.379 8.5l-1.16 1.159l-1.378-1.379L13 7.121zM2.864 4.136l3.64-3.64a1.41 1.41 0 0 1 1.992 0l1.737 1.737c.424.424.424 1.11 0 1.534L6.414 7.586a1.414 1.414 0 0 0 2 2l3.82-3.82a1.084 1.084 0 0 1 1.533 0l1.737 1.738c.55.55.55 1.442 0 1.992l-3.64 3.64a6.364 6.364 0 1 1-9-9" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Magnifier(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 7a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0m-.82 4.74a6 6 0 1 1 1.06-1.06l2.79 2.79a.75.75 0 1 1-1.06 1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifierMinus(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.75 11a4.25 4.25 0 1 0 0-8.5a4.25 4.25 0 0 0 0 8.5m0 1.5a5.725 5.725 0 0 0 3.501-1.188l2.719 2.718a.75.75 0 1 0 1.06-1.06l-2.718-2.719A5.75 5.75 0 1 0 6.75 12.5m-2-6.5a.75.75 0 0 0 0 1.5h4a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifierPlus(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.75 11a4.25 4.25 0 1 0 0-8.5a4.25 4.25 0 0 0 0 8.5m0 1.5a5.725 5.725 0 0 0 3.501-1.188l2.719 2.718a.75.75 0 1 0 1.06-1.06l-2.718-2.719A5.75 5.75 0 1 0 6.75 12.5m.75-7.75a.75.75 0 0 0-1.5 0V6H4.75a.75.75 0 0 0 0 1.5H6v1.25a.75.75 0 0 0 1.5 0V7.5h1.25a.75.75 0 0 0 0-1.5H7.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapPin(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.125 7a4.875 4.875 0 1 1 9.75 0c0 1.864-.774 2.962-1.687 3.815c-.385.36-.765.65-1.17.958l-.365.28a8.807 8.807 0 0 0-.781.668c-.243.24-.535.575-.73 1.01a.337.337 0 0 1-.095.132a.05.05 0 0 1-.015.008s-.01.004-.032.004s-.031-.003-.032-.003a.05.05 0 0 1-.015-.009a.337.337 0 0 1-.095-.131a3.385 3.385 0 0 0-.73-1.01a8.807 8.807 0 0 0-.781-.668c-.125-.097-.246-.19-.366-.28a14.78 14.78 0 0 1-1.169-.96C3.9 9.963 3.125 8.865 3.125 7M14.5 7c0 3.4-2.066 4.975-3.53 6.091c-.634.485-1.156.882-1.345 1.305C9.355 15 8.788 15.5 8 15.5s-1.354-.5-1.625-1.104c-.19-.423-.71-.82-1.346-1.305C3.566 11.975 1.5 10.399 1.5 7a6.5 6.5 0 0 1 13 0m-5 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0M11 7a3 3 0 1 1-6 0a3 3 0 0 1 6 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapPinMinus(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.125 7a4.875 4.875 0 1 1 9.75 0c0 1.864-.774 2.962-1.687 3.815c-.385.36-.765.65-1.17.958l-.365.28a8.807 8.807 0 0 0-.781.668c-.243.24-.535.575-.73 1.01a.337.337 0 0 1-.095.132a.05.05 0 0 1-.015.008s-.01.004-.032.004s-.031-.003-.032-.003a.05.05 0 0 1-.015-.009a.337.337 0 0 1-.095-.131a3.385 3.385 0 0 0-.73-1.01a8.807 8.807 0 0 0-.781-.668c-.125-.097-.246-.19-.366-.28a14.78 14.78 0 0 1-1.169-.96C3.9 9.963 3.125 8.865 3.125 7M14.5 7c0 3.4-2.066 4.975-3.53 6.091c-.634.485-1.156.882-1.345 1.305C9.355 15 8.788 15.5 8 15.5s-1.354-.5-1.625-1.104c-.19-.423-.71-.82-1.346-1.305C3.566 11.975 1.5 10.399 1.5 7a6.5 6.5 0 0 1 13 0m-8.75-.75a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapPinPlus(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.125 7a4.875 4.875 0 1 1 9.75 0c0 1.864-.774 2.962-1.687 3.815c-.385.36-.765.65-1.17.958l-.365.28a8.807 8.807 0 0 0-.781.668c-.243.24-.535.575-.73 1.01a.337.337 0 0 1-.095.132a.05.05 0 0 1-.015.008s-.01.004-.032.004s-.031-.003-.032-.003a.05.05 0 0 1-.015-.009a.337.337 0 0 1-.095-.131a3.385 3.385 0 0 0-.73-1.01a8.807 8.807 0 0 0-.781-.668c-.125-.097-.246-.19-.366-.28a14.78 14.78 0 0 1-1.169-.96C3.9 9.963 3.125 8.865 3.125 7M14.5 7c0 3.4-2.066 4.975-3.53 6.091c-.634.485-1.156.882-1.345 1.305C9.355 15 8.788 15.5 8 15.5s-1.354-.5-1.625-1.104c-.19-.423-.71-.82-1.346-1.305C3.566 11.975 1.5 10.399 1.5 7a6.5 6.5 0 0 1 13 0M8.75 4.75a.75.75 0 0 0-1.5 0v1.5h-1.5a.75.75 0 0 0 0 1.5h1.5v1.5a.75.75 0 0 0 1.5 0v-1.5h1.5a.75.75 0 0 0 0-1.5h-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MathIntersectionShape(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13 13.25V7.5a5 5 0 0 0-10 0v5.75c0 .138.112.25.25.25h1.5a.25.25 0 0 0 .25-.25V7.5a3 3 0 0 1 6 0v5.75c0 .138.112.25.25.25h1.5a.25.25 0 0 0 .25-.25m1.5-5.75v5.75A1.75 1.75 0 0 1 12.75 15h-1.5a1.75 1.75 0 0 1-1.75-1.75V7.5a1.5 1.5 0 1 0-3 0v5.75A1.75 1.75 0 0 1 4.75 15h-1.5a1.75 1.75 0 0 1-1.75-1.75V7.5a6.5 6.5 0 0 1 13 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MathOperations(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.873 3.44a.75.75 0 1 0-1.06-1.06L11.75 3.44l-1.06-1.06a.75.75 0 1 0-1.061 1.06l1.062 1.06l-1.062 1.063a.75.75 0 0 0 1.06 1.06l1.062-1.061l1.062 1.062a.75.75 0 0 0 1.06-1.061L12.813 4.5l1.062-1.061ZM2.25 3.75a.75.75 0 0 0 0 1.5h4.5a.75.75 0 1 0 0-1.5zm3 5.5a.75.75 0 0 0-1.5 0v1.501h-1.5a.75.75 0 0 0 0 1.5h1.5v1.502a.75.75 0 0 0 1.5 0V12.25h1.501a.75.75 0 0 0 0-1.5h-1.5zm4.5 0a.75.75 0 0 0 0 1.5h4a.75.75 0 0 0 0-1.5zm0 3a.75.75 0 0 0 0 1.5h4a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MathUnionShape(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 2.75V8.5a5 5 0 0 0 10 0V2.75a.25.25 0 0 0-.25-.25h-1.5a.25.25 0 0 0-.25.25V8.5a3 3 0 1 1-6 0V2.75a.25.25 0 0 0-.25-.25h-1.5a.25.25 0 0 0-.25.25M1.5 8.5V2.75C1.5 1.784 2.284 1 3.25 1h1.5c.966 0 1.75.784 1.75 1.75V8.5a1.5 1.5 0 1 0 3 0V2.75c0-.966.784-1.75 1.75-1.75h1.5c.966 0 1.75.784 1.75 1.75V8.5a6.5 6.5 0 0 1-13 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Medal(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.836 1.5a.664.664 0 0 1 .491 1.11l-.354.39H4.027l-.354-.39a.664.664 0 0 1 .49-1.11zm-6.445 3l2.124 2.336a.5.5 0 0 0 .37.164h.23a.5.5 0 0 0 .37-.164L10.61 4.5zm.652 2.947L3.5 4.65v1.852a.5.5 0 0 0 .123.328L4.87 8.266a4.51 4.51 0 0 1 1.172-.82Zm-2.08 2.061L2.492 7.814A2 2 0 0 1 2 6.502V2a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v4.502a2 2 0 0 1-.49 1.312l-1.474 1.694a4.5 4.5 0 1 1-8.073 0m7.166-1.242a4.51 4.51 0 0 0-1.172-.82L12.5 4.65v1.852a.5.5 0 0 1-.123.328L11.13 8.266ZM11 11.5a3 3 0 1 1-6 0a3 3 0 0 1 6 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Megaphone(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.113 11.615c.374.814.713.885.887.885c.174 0 .513-.071.887-.885c.377-.816.613-2.077.613-3.615c0-1.538-.236-2.799-.613-3.615c-.374-.814-.713-.885-.887-.885c-.174 0-.513.071-.887.885C10.736 5.2 10.5 6.462 10.5 8c0 1.538.236 2.799.613 3.615M9 8c0 1.469.197 2.815.59 3.857L2.902 9.31a1.402 1.402 0 0 1 0-2.62l6.686-2.548C9.196 5.185 9 6.532 9 8m3 6c2 0 3-2.686 3-6s-1-6-3-6c-.661 0-1.317.12-1.934.356L2.369 5.288a2.902 2.902 0 0 0 0 5.424l.827.315a2.5 2.5 0 1 0 4.67 1.78l2.2.837A5.433 5.433 0 0 0 12 14m-5.537-1.729L4.6 11.563a1 1 0 1 0 1.862.71Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microphone(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.5 7V3.5a1.5 1.5 0 1 0-3 0V7a1.5 1.5 0 1 0 3 0M8 .5a3 3 0 0 0-3 3V7a3 3 0 0 0 6 0V3.5a3 3 0 0 0-3-3m.75 12.454A6.001 6.001 0 0 0 14 7v-.25a.75.75 0 0 0-1.5 0V7a4.5 4.5 0 1 1-9 0v-.25a.75.75 0 0 0-1.5 0V7c0 3.06 2.29 5.585 5.25 5.954v1.796a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneSlash(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.773 9.9A3.004 3.004 0 0 1 5 7v-.94L1.97 3.03a.75.75 0 0 1 1.06-1.06l11 11a.75.75 0 1 1-1.06 1.06l-1.884-1.883a5.97 5.97 0 0 1-2.336.807v1.796a.75.75 0 0 1-1.5 0v-1.796A6.001 6.001 0 0 1 2 7v-.25a.75.75 0 0 1 1.5 0V7a4.5 4.5 0 0 0 6.481 4.042L8.825 9.885zM9.5 3.5v2.798l1.415 1.415C10.97 7.484 11 7.246 11 7V3.5a3 3 0 0 0-5.669-1.371l1.18 1.18A1.5 1.5 0 0 1 9.5 3.5m2.587 5.385L13.2 9.997c.51-.882.8-1.905.8-2.997v-.25a.75.75 0 0 0-1.5 0V7a4.48 4.48 0 0 1-.413 1.885" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.75 8a.75.75 0 0 1 .75-.75h11a.75.75 0 0 1 0 1.5h-11A.75.75 0 0 1 1.75 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Molecule(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiMolecule0)"><path fill="currentColor" fill-rule="evenodd" d="M4.706 3.39a1 1 0 1 1 .108-.064a.76.76 0 0 0-.108.064m.108 1.546a2.5 2.5 0 1 1 1.292-.762l.849 1.485a3.497 3.497 0 0 1 3.191.576l2.222-1.851a1.75 1.75 0 1 1 1.101 1.035l-2.362 1.968c.251.483.393 1.031.393 1.613c0 .557-.13 1.083-.361 1.55l.958.767a2.25 2.25 0 1 1-.959 1.154l-.937-.75a3.501 3.501 0 0 1-5.62-1.971H3.33a1.75 1.75 0 1 1 0-1.5h1.25a3.497 3.497 0 0 1 1.072-1.846l-.84-1.468ZM10 9a2 2 0 1 1-4 0a2 2 0 0 1 4 0m3.25 5a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiMolecule0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Moon(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 13.5a5.5 5.5 0 0 0 2.263-10.514a5.5 5.5 0 0 1-7.278 7.278A5.501 5.501 0 0 0 8 13.5M1.045 8.795a7.001 7.001 0 1 0 7.75-7.75l-.028-.003A7.078 7.078 0 0 0 8 1c-.527 0-.59.842-.185 1.18a4.02 4.02 0 0 1 .342.322A4 4 0 1 1 2.18 7.814C1.842 7.41 1 7.474 1 8a7.078 7.078 0 0 0 .045.794Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mug(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.123 5.357c.712-.37.877-.75.877-1.107s-.165-.736-.877-1.107C9.373 2.752 8.2 2.5 6.75 2.5c-1.45 0-2.622.252-3.373.643c-.712.37-.877.75-.877 1.107s.165.736.877 1.107C4.127 5.748 5.3 6 6.75 6c1.45 0 2.622-.252 3.373-.643M6.75 7.5c1.684 0 3.198-.305 4.25-.914v5.164c0 .357-.165.736-.877 1.107c-.75.391-1.924.643-3.373.643c-1.45 0-2.622-.252-3.373-.643c-.712-.37-.877-.75-.877-1.107V6.586c1.052.61 2.566.914 4.25.914m5.75-3.25C12.5 2.083 9.926 1 6.75 1S1 2.083 1 4.25v7.5C1 13.917 3.574 15 6.75 15s5.75-1.083 5.75-3.25v-1.182c.068-.026.134-.053.2-.08c.616-.256 1.27-.613 1.791-1.112c.534-.512 1.01-1.256 1.01-2.218c0-.437-.097-.89-.345-1.304a2.38 2.38 0 0 0-.956-.89c-.562-.292-1.185-.358-1.7-.352zm0 1.862v2.82c.86-.425 1.501-1.016 1.501-1.774c0-.806-.652-1.058-1.501-1.046" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicNote(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13 5.46v3.08a2.94 2.94 0 0 0-.976 0a4.15 4.15 0 0 0-.274.055C10.231 8.96 9 10.263 9 11.505c0 1.243 1.231 1.955 2.75 1.59c1.519-.364 2.75-1.667 2.75-2.91c0-.038-.001-.076-.004-.113A.745.745 0 0 0 14.5 10V1.25a.75.75 0 0 0-.932-.728L6.136 2.38l-.568.142A.75.75 0 0 0 5 3.25v7.291a2.941 2.941 0 0 0-.976 0c-.09.014-.181.032-.274.054C2.231 10.96 1 12.263 1 13.505c0 1.243 1.231 1.955 2.75 1.59c1.519-.364 2.75-1.667 2.75-2.91c0-.038-.001-.076-.003-.113A.758.758 0 0 0 6.5 12V7.086zm0-1.546V2.211L6.5 3.836v1.703l6.136-1.534zm-8.003 8.127a1.5 1.5 0 0 0 .003.144c0 .133-.079.419-.396.754a2.497 2.497 0 0 1-1.204.698c-.47.113-.748.023-.844-.032a.215.215 0 0 1-.047-.036l-.001-.002l-.003-.007a.214.214 0 0 1-.005-.054c0-.133.079-.419.396-.754a2.5 2.5 0 0 1 1.204-.698a1.58 1.58 0 0 1 .637-.037c.086.016.173.024.26.024m8-2a1.465 1.465 0 0 0 .003.144c0 .133-.079.419-.396.754a2.497 2.497 0 0 1-1.204.698c-.47.113-.748.023-.844-.032a.214.214 0 0 1-.047-.036l-.001-.002l-.003-.007a.216.216 0 0 1-.005-.054c0-.133.079-.419.396-.754a2.497 2.497 0 0 1 1.204-.698a1.58 1.58 0 0 1 .637-.037c.086.016.173.024.26.024" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NodesDown(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 11a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0M12 14a3 3 0 1 0-.79-5.895L10.092 6.15a3 3 0 1 0-4.185 0L4.79 8.105A3.003 3.003 0 0 0 1 11a3 3 0 1 0 5.092-2.15L7.21 6.895a3.003 3.003 0 0 0 1.58 0L9.908 8.85A3 3 0 0 0 12 14m-6.5-3a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0M8 2.5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NodesLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 2.5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3M2 4a3 3 0 1 1 5.895.79L9.85 5.908a3 3 0 1 1 0 4.185L7.895 11.21A3.003 3.003 0 0 1 5 15a3 3 0 1 1 2.15-5.092L9.105 8.79a3.003 3.003 0 0 1 0-1.58L7.15 6.092A3 3 0 0 1 2 4m3 6.5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3M13.5 8a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NodesRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 2.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3M14 4a3 3 0 1 0-5.895.79L6.15 5.908a3 3 0 1 0 0 4.185l1.955 1.117A3.003 3.003 0 0 0 11 15a3 3 0 1 0-2.15-5.092L6.895 8.79a3.003 3.003 0 0 0 0-1.58L8.85 6.092A3 3 0 0 0 14 4m-3 6.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3M2.5 8a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NodesUp(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 5a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0-3 0M4 2a3 3 0 1 0 .79 5.895L5.908 9.85a3 3 0 1 0 4.185 0l1.117-1.955A3.003 3.003 0 0 0 15 5a3 3 0 1 0-5.092 2.15L8.79 9.105a3.003 3.003 0 0 0-1.58 0L6.092 7.15A3 3 0 0 0 4 2m6.5 3a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0-3 0M8 13.5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NutHex(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m13.873 8.25l-2.72 4.712a.5.5 0 0 1-.433.25H5.28a.5.5 0 0 1-.433-.25L2.127 8.25a.5.5 0 0 1 0-.5l2.72-4.712a.5.5 0 0 1 .433-.25h5.44a.5.5 0 0 1 .433.25l2.72 4.712a.5.5 0 0 1 0 .5m1.3-1.25a2 2 0 0 1 0 2l-2.72 4.712a2 2 0 0 1-1.733 1H5.28a2 2 0 0 1-1.732-1L.828 9a2 2 0 0 1 0-2l2.72-4.712a2 2 0 0 1 1.732-1h5.44a2 2 0 0 1 1.732 1zM9.5 8a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0M11 8a3 3 0 1 1-6 0a3 3 0 0 1 6 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectAlignBottom(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.75 12.5a.75.75 0 0 0 0 1.5h12.5a.75.75 0 0 0 0-1.5zM7 9.5h2a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v4.5a1 1 0 0 0 1 1m-2.5-1A2.5 2.5 0 0 0 7 11h2a2.5 2.5 0 0 0 2.5-2.5V4A2.5 2.5 0 0 0 9 1.5H7A2.5 2.5 0 0 0 4.5 4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectAlignCenterHorizontal(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13 7v2a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1m-1-2.5A2.5 2.5 0 0 1 14.5 7v2a2.5 2.5 0 0 1-2.5 2.5H8.75v2.75a.75.75 0 0 1-1.5 0V11.5H4A2.5 2.5 0 0 1 1.5 9V7A2.5 2.5 0 0 1 4 4.5h3.25V1.75a.75.75 0 1 1 1.5 0V4.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectAlignCenterVertical(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 3h2a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1M4.5 4A2.5 2.5 0 0 1 7 1.5h2A2.5 2.5 0 0 1 11.5 4v3.25h2.75a.75.75 0 0 1 0 1.5H11.5V12A2.5 2.5 0 0 1 9 14.5H7A2.5 2.5 0 0 1 4.5 12V8.75H1.75a.75.75 0 0 1 0-1.5H4.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectAlignJustifyHorizontal(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.25 15a.75.75 0 0 1-.75-.75V1.75a.75.75 0 1 1 1.5 0v12.5a.75.75 0 0 1-.75.75M11 9V7a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1m1.5-2A2.5 2.5 0 0 0 10 4.5H6A2.5 2.5 0 0 0 3.5 7v2A2.5 2.5 0 0 0 6 11.5h4A2.5 2.5 0 0 0 12.5 9zM1 14.25a.75.75 0 0 0 1.5 0V1.75a.75.75 0 1 0-1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectAlignJustifyVertical(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 14.25a.75.75 0 0 1 .75-.75h12.5a.75.75 0 0 1 0 1.5H1.75a.75.75 0 0 1-.75-.75M7 11h2a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1m2 1.5a2.5 2.5 0 0 0 2.5-2.5V6A2.5 2.5 0 0 0 9 3.5H7A2.5 2.5 0 0 0 4.5 6v4A2.5 2.5 0 0 0 7 12.5zM1.75 1a.75.75 0 0 0 0 1.5h12.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectAlignLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 14.25a.75.75 0 0 0 1.5 0V1.75a.75.75 0 0 0-1.5 0zM6.5 7v2a1 1 0 0 0 1 1H12a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1H7.5a1 1 0 0 0-1 1m1-2.5A2.5 2.5 0 0 0 5 7v2a2.5 2.5 0 0 0 2.5 2.5H12A2.5 2.5 0 0 0 14.5 9V7A2.5 2.5 0 0 0 12 4.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectAlignRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 14.25a.75.75 0 0 0 1.5 0V1.75a.75.75 0 0 0-1.5 0zM9.5 7v2a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h4.5a1 1 0 0 1 1 1m-1-2.5A2.5 2.5 0 0 1 11 7v2a2.5 2.5 0 0 1-2.5 2.5H4A2.5 2.5 0 0 1 1.5 9V7A2.5 2.5 0 0 1 4 4.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectAlignTop(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.75 2a.75.75 0 0 0 0 1.5h12.5a.75.75 0 0 0 0-1.5zM7 6.5h2a1 1 0 0 1 1 1V12a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V7.5a1 1 0 0 1 1-1m-2.5 1A2.5 2.5 0 0 1 7 5h2a2.5 2.5 0 0 1 2.5 2.5V12A2.5 2.5 0 0 1 9 14.5H7A2.5 2.5 0 0 1 4.5 12z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectsAlignBottom(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.75 14a.75.75 0 0 1 0-1.5h12.5a.75.75 0 0 1 0 1.5zM5.5 9.5H4a.5.5 0 0 1-.5-.5V3.5A.5.5 0 0 1 4 3h1.5a.5.5 0 0 1 .5.5V9a.5.5 0 0 1-.5.5M4 11a2 2 0 0 1-2-2V3.5a2 2 0 0 1 2-2h1.5a2 2 0 0 1 2 2V9a2 2 0 0 1-2 2zm8-1.5h-1.5A.5.5 0 0 1 10 9V7a.5.5 0 0 1 .5-.5H12a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5M10.5 11a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2H12a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectsAlignCenterHorizontal(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13 5.5V4a.5.5 0 0 0-.5-.5h-9A.5.5 0 0 0 3 4v1.5a.5.5 0 0 0 .5.5h9a.5.5 0 0 0 .5-.5m-4.25 2h3.75a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H8.75v-.25a.75.75 0 0 0-1.5 0V2H3.5a2 2 0 0 0-2 2v1.5a2 2 0 0 0 2 2h3.75v1H6a2 2 0 0 0-2 2V12a2 2 0 0 0 2 2h1.25v.25a.75.75 0 0 0 1.5 0V14H10a2 2 0 0 0 2-2v-1.5a2 2 0 0 0-2-2H8.75zM8 10H6a.5.5 0 0 0-.5.5V12a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-1.5a.5.5 0 0 0-.5-.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectsAlignCenterVertical(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.5 3H4a.5.5 0 0 0-.5.5v9a.5.5 0 0 0 .5.5h1.5a.5.5 0 0 0 .5-.5v-9a.5.5 0 0 0-.5-.5m2 4.25V3.5a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v3.75h-.25a.75.75 0 0 0 0 1.5H2v3.75a2 2 0 0 0 2 2h1.5a2 2 0 0 0 2-2V8.75h1V10a2 2 0 0 0 2 2H12a2 2 0 0 0 2-2V8.75h.25a.75.75 0 0 0 0-1.5H14V6a2 2 0 0 0-2-2h-1.5a2 2 0 0 0-2 2v1.25zM10 8v2a.5.5 0 0 0 .5.5H12a.5.5 0 0 0 .5-.5V6a.5.5 0 0 0-.5-.5h-1.5a.5.5 0 0 0-.5.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectsAlignJustifyHorizontal(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.25 1a.75.75 0 0 1 .75.75v12.5a.75.75 0 0 1-1.5 0V1.75a.75.75 0 0 1 .75-.75M11 4v1.5a.5.5 0 0 1-.5.5h-5a.5.5 0 0 1-.5-.5V4a.5.5 0 0 1 .5-.5h5a.5.5 0 0 1 .5.5m-.5-2a2 2 0 0 1 2 2v1.5a2 2 0 0 1-2 2h-5a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2zm.5 8.5V12a.5.5 0 0 1-.5.5h-5A.5.5 0 0 1 5 12v-1.5a.5.5 0 0 1 .5-.5h5a.5.5 0 0 1 .5.5m-.5-2a2 2 0 0 1 2 2V12a2 2 0 0 1-2 2h-5a2 2 0 0 1-2-2v-1.5a2 2 0 0 1 2-2zM1 1.75a.75.75 0 0 1 1.5 0v12.5a.75.75 0 0 1-1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectsAlignJustifyVertical(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 14.25a.75.75 0 0 1-.75.75H1.75a.75.75 0 0 1 0-1.5h12.5a.75.75 0 0 1 .75.75M12 11h-1.5a.5.5 0 0 1-.5-.5v-5a.5.5 0 0 1 .5-.5H12a.5.5 0 0 1 .5.5v5a.5.5 0 0 1-.5.5m2-.5a2 2 0 0 1-2 2h-1.5a2 2 0 0 1-2-2v-5a2 2 0 0 1 2-2H12a2 2 0 0 1 2 2zm-8.5.5H4a.5.5 0 0 1-.5-.5v-5A.5.5 0 0 1 4 5h1.5a.5.5 0 0 1 .5.5v5a.5.5 0 0 1-.5.5m2-.5a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-5a2 2 0 0 1 2-2h1.5a2 2 0 0 1 2 2zM14.25 1a.75.75 0 0 1 0 1.5H1.75a.75.75 0 0 1 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectsAlignLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 1.75a.75.75 0 0 1 1.5 0v12.5a.75.75 0 0 1-1.5 0zM6.5 5.5V4a.5.5 0 0 1 .5-.5h5.5a.5.5 0 0 1 .5.5v1.5a.5.5 0 0 1-.5.5H7a.5.5 0 0 1-.5-.5M5 4a2 2 0 0 1 2-2h5.5a2 2 0 0 1 2 2v1.5a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2zm1.5 8v-1.5A.5.5 0 0 1 7 10h2a.5.5 0 0 1 .5.5V12a.5.5 0 0 1-.5.5H7a.5.5 0 0 1-.5-.5M5 10.5a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2V12a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectsAlignRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14 1.75a.75.75 0 0 0-1.5 0v12.5a.75.75 0 0 0 1.5 0zM9.5 5.5V4a.5.5 0 0 0-.5-.5H3.5A.5.5 0 0 0 3 4v1.5a.5.5 0 0 0 .5.5H9a.5.5 0 0 0 .5-.5M11 4a2 2 0 0 0-2-2H3.5a2 2 0 0 0-2 2v1.5a2 2 0 0 0 2 2H9a2 2 0 0 0 2-2zm-1.5 8v-1.5A.5.5 0 0 0 9 10H7a.5.5 0 0 0-.5.5V12a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5m1.5-1.5a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2V12a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectsAlignTop(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.75 2a.75.75 0 0 0 0 1.5h12.5a.75.75 0 0 0 0-1.5zM5.5 6.5H4a.5.5 0 0 0-.5.5v5.5a.5.5 0 0 0 .5.5h1.5a.5.5 0 0 0 .5-.5V7a.5.5 0 0 0-.5-.5M4 5a2 2 0 0 0-2 2v5.5a2 2 0 0 0 2 2h1.5a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2zm8 1.5h-1.5a.5.5 0 0 0-.5.5v2a.5.5 0 0 0 .5.5H12a.5.5 0 0 0 .5-.5V7a.5.5 0 0 0-.5-.5M10.5 5a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2H12a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OctagonXmark(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.829 2.5h4.342a.5.5 0 0 1 .354.146l2.829 2.829a.5.5 0 0 1 .146.353v4.343a.5.5 0 0 1-.146.354l-2.829 2.829a.5.5 0 0 1-.354.146H5.83a.5.5 0 0 1-.354-.146l-2.829-2.829a.5.5 0 0 1-.146-.354V5.829a.5.5 0 0 1 .147-.353l2.828-2.829a.5.5 0 0 1 .354-.146Zm-1.415-.914A2 2 0 0 1 5.83 1h4.342a2 2 0 0 1 1.415.586l2.828 2.828A2 2 0 0 1 15 5.828v4.343a2 2 0 0 1-.586 1.415l-2.828 2.828A2 2 0 0 1 10.17 15H5.83a2 2 0 0 1-1.415-.586l-2.828-2.828A2 2 0 0 1 1 10.17V5.828a2 2 0 0 1 .586-1.414zM6.53 5.47a.75.75 0 1 0-1.06 1.06L6.94 8L5.47 9.47a.75.75 0 1 0 1.06 1.06L8 9.06l1.47 1.47a.75.75 0 1 0 1.06-1.06L9.06 8l1.47-1.47a.75.75 0 1 0-1.06-1.06L8 6.94z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OfficeBadge(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.75 1.75a.75.75 0 0 0-1.5 0v3a.75.75 0 0 0 1.5 0zM5.25 4a.75.75 0 0 1 0 1.5H4A1.5 1.5 0 0 0 2.5 7v4A1.5 1.5 0 0 0 4 12.5h8a1.5 1.5 0 0 0 1.5-1.5V7A1.5 1.5 0 0 0 12 5.5h-1.25a.75.75 0 0 1 0-1.5H12a3 3 0 0 1 3 3v4a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3V7a3 3 0 0 1 3-3zm6 4.5a.75.75 0 0 0 0-1.5h-6.5a.75.75 0 0 0 0 1.5zM9 10.25a.75.75 0 0 1-.75.75h-3.5a.75.75 0 0 1 0-1.5h3.5a.75.75 0 0 1 .75.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Palette(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.012 10c.431.004.764-.15 1.002-.411c.244-.268.486-.762.486-1.589a5.5 5.5 0 1 0-5.17 5.491a4.279 4.279 0 0 1-.106-.89a2.37 2.37 0 0 1 .495-1.48c.386-.493.92-.763 1.448-.914C10.69 10.06 11.303 10 12 10zM8.43 14.01v-.005zM12 11.5c1.66.013 3-1.25 3-3.5a7 7 0 1 0-7 7c2.19 0 2.011-.83 1.827-1.68c-.194-.898-.393-1.82 2.173-1.82M9 5a1 1 0 1 1-2 0a1 1 0 0 1 2 0m2 2.75a1 1 0 1 0 0-2a1 1 0 0 0 0 2m-4.75-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0M5.75 11a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperPlane(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.29 13.904L5.25 10.75L2.096 8.71a2.4 2.4 0 0 1 .5-4.278l9.273-3.296a2.346 2.346 0 0 1 2.995 2.995l-1.413-.502a.844.844 0 0 0-1.08-1.08L3.1 5.846a.9.9 0 0 0-.19 1.604l2.78 1.799l3.279-3.28a.75.75 0 1 1 1.06 1.061L6.75 10.31l1.799 2.779a.9.9 0 0 0 1.604-.188l3.297-9.272l1.413.502l-3.296 9.273a2.4 2.4 0 0 1-4.277.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paperclip(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.77 10.73a.717.717 0 0 0 .01-.01l3.08-3.08a3.889 3.889 0 1 0-5.5-5.5L4.73 4.77l-.01.01l-1.667 1.666a5.303 5.303 0 0 0 7.5 7.5l3.167-3.166a.75.75 0 1 0-1.061-1.06l-3.166 3.166a3.803 3.803 0 1 1-5.379-5.379L5.33 6.291l.011-.01L8.421 3.2a2.39 2.39 0 0 1 3.38 3.378l-1.13 1.13l-.012.012l-2.995 2.994a.975.975 0 1 1-1.378-1.378L9.28 6.34a.75.75 0 0 0-1.06-1.06L5.225 8.274a2.475 2.475 0 0 0 3.5 3.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Passport(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 3v10a.5.5 0 0 0 .5.5h7a1.5 1.5 0 0 0 1.5-1.5V4A1.5 1.5 0 0 0 11 2.5H4a.5.5 0 0 0-.5.5M2 13a2 2 0 0 0 2 2h7a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3H4a2 2 0 0 0-2 2zm3.5-2.25a.75.75 0 0 1 .75-.75h3.5a.75.75 0 0 1 0 1.5h-3.5a.75.75 0 0 1-.75-.75M9 6.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0m1.5 0a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 3h1.75a.5.5 0 0 1 .5.5v9a.5.5 0 0 1-.5.5H3.5a.5.5 0 0 1-.5-.5v-9a.5.5 0 0 1 .5-.5m-2 .5a2 2 0 0 1 2-2h1.75a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2H3.5a2 2 0 0 1-2-2zm9.25-.5h1.75a.5.5 0 0 1 .5.5v9a.5.5 0 0 1-.5.5h-1.75a.5.5 0 0 1-.5-.5v-9a.5.5 0 0 1 .5-.5m-2 .5a2 2 0 0 1 2-2h1.75a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2h-1.75a2 2 0 0 1-2-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 3.5a2 2 0 0 1 2-2H5a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2H3.5a2 2 0 0 1-2-2zm7.5 0a2 2 0 0 1 2-2h1.5a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2H11a2 2 0 0 1-2-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pencil(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.423 1A3.577 3.577 0 0 1 15 4.577c0 .27-.108.53-.3.722l-.528.529l-1.971 1.971l-5.059 5.059a3 3 0 0 1-1.533.82l-2.638.528a1 1 0 0 1-1.177-1.177l.528-2.638a3 3 0 0 1 .82-1.533l5.059-5.059l2.5-2.5c.191-.191.451-.299.722-.299m-2.31 4.009l-4.91 4.91a1.5 1.5 0 0 0-.41.766l-.38 1.903l1.902-.38a1.5 1.5 0 0 0 .767-.41l4.91-4.91a2.077 2.077 0 0 0-1.88-1.88Zm3.098.658a3.59 3.59 0 0 0-1.878-1.879l1.28-1.28c.995.09 1.788.884 1.878 1.88l-1.28 1.28Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PencilToLine(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.238 3.64a1.854 1.854 0 0 0-1.629-1.628l-.8.8a3.367 3.367 0 0 1 1.63 1.628zM4.74 7.88l3.87-3.868a1.854 1.854 0 0 1 1.628 1.629L6.369 9.51a1.5 1.5 0 0 1-.814.418l-1.48.247l.247-1.48a1.5 1.5 0 0 1 .418-.814ZM9.72.78l-2 2l-4.04 4.04a3 3 0 0 0-.838 1.628L2.48 10.62a1 1 0 0 0 1.151 1.15l2.17-.36a3 3 0 0 0 1.629-.839l4.04-4.04l2-2c.18-.18.28-.423.28-.677A3.353 3.353 0 0 0 10.397.5c-.254 0-.498.1-.678.28ZM2.75 13a.75.75 0 0 0 0 1.5h10.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PencilToSquare(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.169 6.331a3 3 0 0 0-.833 1.6l-.338 1.912a1 1 0 0 0 1.159 1.159l1.912-.338a3 3 0 0 0 1.6-.833l3.07-3.07l2-2A.894.894 0 0 0 15 4.13A3.13 3.13 0 0 0 11.87 1a.894.894 0 0 0-.632.262l-2 2l-3.07 3.07Zm3.936-1.814L7.229 7.392a1.5 1.5 0 0 0-.416.8L6.6 9.4l1.208-.213l.057-.01a1.5 1.5 0 0 0 .743-.406l2.875-2.876a1.63 1.63 0 0 0-1.378-1.378m2.558.199a3.143 3.143 0 0 0-1.379-1.38l.82-.82a1.63 1.63 0 0 1 1.38 1.38l-.82.82ZM8 2.25a.75.75 0 0 0-.75-.75H4.5a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h7a3 3 0 0 0 3-3V8.75a.75.75 0 0 0-1.5 0v2.75a1.5 1.5 0 0 1-1.5 1.5h-7A1.5 1.5 0 0 1 3 11.5v-7A1.5 1.5 0 0 1 4.5 3h2.75A.75.75 0 0 0 8 2.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Percent(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 7.5a3 3 0 1 1 0-6a3 3 0 0 1 0 6m1.5-3a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0m5.5 10a3 3 0 1 1 0-6a3 3 0 1 1 0 6m1.5-3a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0m.78-8.22a.75.75 0 0 0-1.06-1.06l-10.5 10.5a.75.75 0 1 0 1.06 1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Person(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 4.5a2 2 0 1 1-4 0a2 2 0 0 1 4 0m1.5 0a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0m-9 8c0-.204.22-.809 1.32-1.459C4.838 10.44 6.32 10 8 10c1.68 0 3.162.44 4.18 1.041c1.1.65 1.32 1.255 1.32 1.459a1 1 0 0 1-1 1h-9a1 1 0 0 1-1-1m5.5-4c-3.85 0-7 2-7 4A2.5 2.5 0 0 0 3.5 15h9a2.5 2.5 0 0 0 2.5-2.5c0-2-3.15-4-7-4" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonGear(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiPersonGear0)"><path fill="currentColor" fill-rule="evenodd" d="M7 6.5a2 2 0 1 0 0-4a2 2 0 0 0 0 4M7 8a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7m-.5 1.299c0-.43-.36-.774-.785-.724C2.473 8.955 0 10.728 0 12.5A2.5 2.5 0 0 0 2.5 15h3.25a.75.75 0 0 0 0-1.5H2.5a1 1 0 0 1-1-1c0-.205.22-.809 1.32-1.459c.765-.452 1.792-.813 2.969-.964c.397-.051.711-.378.711-.778m8.002 1.89a2.223 2.223 0 0 1-1.66-2.138l-.68-.256a2.38 2.38 0 0 1-2.747.816l-.406.569c.57.788.57 1.87 0 2.657l.406.569a2.38 2.38 0 0 1 2.747.816l.68-.255a2.223 2.223 0 0 1 1.66-2.139zm-.112-2.396a.715.715 0 0 0-.418-.921l-1.798-.674a.865.865 0 0 0-1.114.506a.87.87 0 0 1-1.32.395a.884.884 0 0 0-1.23.2L7.434 9.805a.771.771 0 0 0 .179 1.076c.43.308.43.948 0 1.255a.771.771 0 0 0-.18 1.076l1.077 1.506a.884.884 0 0 0 1.23.2a.87.87 0 0 1 1.32.395a.865.865 0 0 0 1.114.506l1.798-.674a.715.715 0 0 0 .418-.92a.715.715 0 0 1 .67-.966h.134a.808.808 0 0 0 .809-.809v-1.883a.808.808 0 0 0-.809-.808h-.135a.715.715 0 0 1-.669-.966m-1.833 2.659a1.052 1.052 0 1 1-2.103 0a1.052 1.052 0 0 1 2.103 0" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiPersonGear0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonMagnifier(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 6.5a2 2 0 1 0 0-4a2 2 0 0 0 0 4M8 8a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7m1 1.225a.71.71 0 0 0-.679-.72A11.087 11.087 0 0 0 8 8.5c-3.85 0-7 2-7 4A2.5 2.5 0 0 0 3.5 15h3.25a.75.75 0 0 0 0-1.5H3.5a1 1 0 0 1-1-1c0-.204.22-.809 1.32-1.459C4.838 10.44 6.32 10 8 10c.058 0 .117 0 .175.002c.442.008.825-.335.825-.777m4 3.275a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3m0 1.5a2.986 2.986 0 0 1-1.524-.415L10.03 15.03a.75.75 0 1 1-1.06-1.06l1.445-1.446A3 3 0 1 1 13 14" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonNutHex(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 6.5a2 2 0 1 0 0-4a2 2 0 0 0 0 4M7 8a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7m-.5 1.299c0-.43-.36-.774-.785-.724C2.473 8.955 0 10.728 0 12.5A2.5 2.5 0 0 0 2.5 15h3.25a.75.75 0 0 0 0-1.5H2.5a1 1 0 0 1-1-1c0-.205.22-.809 1.32-1.459c.765-.452 1.792-.813 2.969-.964c.397-.051.711-.378.711-.778m6.373 5.201l1.406-2.5l-1.406-2.5h-2.746L8.721 12l1.406 2.5zm2.713-1.765a1.5 1.5 0 0 0 0-1.47l-1.406-2.5A1.5 1.5 0 0 0 12.873 8h-2.746a1.5 1.5 0 0 0-1.307.765l-1.406 2.5a1.5 1.5 0 0 0 0 1.47l1.406 2.5a1.5 1.5 0 0 0 1.307.765h2.746a1.5 1.5 0 0 0 1.307-.765zM12.5 12a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonPencil(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 6.5a2 2 0 1 0 0-4a2 2 0 0 0 0 4M8 8a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7m1 1.225a.71.71 0 0 0-.679-.72A11.087 11.087 0 0 0 8 8.5c-3.85 0-7 2-7 4A2.5 2.5 0 0 0 3.5 15h2.25a.75.75 0 0 0 0-1.5H3.5a1 1 0 0 1-1-1c0-.204.22-.809 1.32-1.459C4.838 10.44 6.32 10 8 10c.058 0 .117 0 .175.002c.442.008.825-.335.825-.777m3.59.307c.434.102.776.444.879.878l-2.823 2.822a1.5 1.5 0 0 1-.848.425l-.53.075l.075-.53a1.5 1.5 0 0 1 .425-.848zm-.883 4.76l3.068-3.067a.767.767 0 0 0 .225-.543A2.683 2.683 0 0 0 12.318 8a.766.766 0 0 0-.543.224l-3.068 3.069a3 3 0 0 0-.848 1.697l-.17 1.19a1 1 0 0 0 1.13 1.131l1.191-.17a3 3 0 0 0 1.697-.848Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonPlanetEarth(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 6.5a2 2 0 1 0 0-4a2 2 0 0 0 0 4M7 8a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7m3.229 3.909l-.702-.28a2.5 2.5 0 0 1 3.339-1.975c-.121.185-.303.33-.523.403l-.035.012a.951.951 0 0 0 0 1.805l.46.154a.66.66 0 0 1 .383.331a.66.66 0 0 0 .957.254l.365-.242a2.5 2.5 0 0 1-2.4 2.128a.768.768 0 0 0 .041-.247v-.08a.686.686 0 0 0-.685-.686a.686.686 0 0 1-.686-.686v-.132a.818.818 0 0 0-.514-.76ZM12 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8M6.25 8.525c.414-.028.75.31.75.725c0 .414-.336.747-.75.78c-1.369.107-2.567.5-3.43 1.011c-1.1.65-1.32 1.255-1.32 1.459a1 1 0 0 0 1 1h3.75a.75.75 0 0 1 0 1.5H2.5A2.5 2.5 0 0 1 0 12.5c0-1.868 2.75-3.737 6.25-3.975" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonPlus(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 6.5a2 2 0 1 0 0-4a2 2 0 0 0 0 4M8 8a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7m1 1.225a.71.71 0 0 0-.679-.72A11.087 11.087 0 0 0 8 8.5c-3.85 0-7 2-7 4A2.5 2.5 0 0 0 3.5 15h8.75a.75.75 0 0 0 0-1.5H3.5a1 1 0 0 1-1-1c0-.204.22-.809 1.32-1.459C4.838 10.44 6.32 10 8 10c.058 0 .117 0 .175.002c.442.008.825-.335.825-.777M13.75 8a.75.75 0 0 0-1.5 0v1.25H11a.75.75 0 0 0 0 1.5h1.25V12a.75.75 0 0 0 1.5 0v-1.25H15a.75.75 0 0 0 0-1.5h-1.25z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonSpeaker(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 4.5a2 2 0 1 1-4 0a2 2 0 0 1 4 0m1.5 0a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0m-9 8c0-.204.22-.809 1.32-1.459C4.838 10.44 6.32 10 8 10c1.68 0 3.162.44 4.18 1.041c1.1.65 1.32 1.255 1.32 1.459a1 1 0 0 1-1 1h-.25a.75.75 0 0 0 0 1.5h.25a2.5 2.5 0 0 0 2.5-2.5c0-2-3.15-4-7-4s-7 2-7 4A2.5 2.5 0 0 0 3.5 15h.25a.75.75 0 0 0 0-1.5H3.5a1 1 0 0 1-1-1m3.335-.92a.75.75 0 0 0-.68 1.336l.02.012c.022.013.06.038.11.074c.098.074.237.195.38.373c.193.242.404.604.512 1.125H5.75a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 0-1.5h-.427c.108-.52.319-.883.513-1.125c.142-.178.281-.299.38-.373a1.416 1.416 0 0 1 .128-.086a.75.75 0 0 0-.68-1.337l.336.671a61.14 61.14 0 0 0-.336-.67h-.003l-.004.003l-.01.005a1.686 1.686 0 0 0-.1.058a2.9 2.9 0 0 0-.231.156a3.629 3.629 0 0 0-.652.636c-.377.471-.733 1.146-.863 2.062H7.7c-.13-.916-.486-1.59-.863-2.062a3.629 3.629 0 0 0-.652-.636a2.9 2.9 0 0 0-.331-.214l-.01-.005l-.004-.002l-.002-.001c-.001 0-.002 0-.337.67l.335-.67Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonWorker(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.535 3A3.503 3.503 0 0 1 7.25.08v1.67a.75.75 0 1 0 1.5 0V.08A3.503 3.503 0 0 1 11.464 3h.286a.75.75 0 0 1 0 1.5h-.25a3.5 3.5 0 0 1-7 0h-.25a.75.75 0 0 1 0-1.5zM8 6.5a2 2 0 0 1-2-2h4a2 2 0 0 1-2 2m-5.5 6c0-.204.22-.809 1.32-1.459a5.74 5.74 0 0 1 .223-.125L5.01 13.5H3.5a1 1 0 0 1-1-1m4.114 1l-1.179-3.143A9.213 9.213 0 0 1 8 10c.93 0 1.8.135 2.565.357L9.387 13.5H6.612Zm4.375 0H12.5a1 1 0 0 0 1-1c0-.204-.22-.809-1.32-1.459a5.7 5.7 0 0 0-.223-.125l-.969 2.584ZM8 8.5c-3.85 0-7 2-7 4A2.5 2.5 0 0 0 3.5 15h9a2.5 2.5 0 0 0 2.5-2.5c0-2-3.15-4-7-4" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonXmark(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 6.5a2 2 0 1 0 0-4a2 2 0 0 0 0 4M8 8a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7m1 1.225a.71.71 0 0 0-.679-.72A11.087 11.087 0 0 0 8 8.5c-3.85 0-7 2-7 4A2.5 2.5 0 0 0 3.5 15h8.75a.75.75 0 0 0 0-1.5H3.5a1 1 0 0 1-1-1c0-.204.22-.809 1.32-1.459C4.838 10.44 6.32 10 8 10c.058 0 .117 0 .175.002c.442.008.825-.335.825-.777m4-.286l-.97-.97a.75.75 0 1 0-1.06 1.061l.97.97l-.97.97a.75.75 0 1 0 1.06 1.06l.97-.97l.97.97a.75.75 0 1 0 1.06-1.06l-.97-.97l.97-.97a.75.75 0 0 0-1.06-1.06l-.97.97Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Persons(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.5 6a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0 1.5a3 3 0 1 0 0-6a3 3 0 0 0 0 6m-3.029 2.886c-.777.54-.971 1.063-.971 1.306c0 .446.362.808.808.808h6.384a.808.808 0 0 0 .808-.808c0-.244-.194-.767-.971-1.306C7.792 9.875 6.719 9.5 5.5 9.5c-1.218 0-2.292.375-3.029.886M0 11.692C0 9.846 2.475 8 5.5 8c1.18 0 2.278.281 3.177.734A5.671 5.671 0 0 1 11.5 8c2.475 0 4.5 1.538 4.5 3.077A1.923 1.923 0 0 1 14.077 13h-3.483c-.416.604-1.113 1-1.902 1H2.308A2.308 2.308 0 0 1 0 11.692m10.991-.192h3.086c.234 0 .423-.19.423-.423c0-.103-.096-.472-.688-.89c-.554-.393-1.375-.687-2.312-.687c-.517 0-.999.09-1.42.236c.526.534.854 1.146.911 1.764M12.5 5a1 1 0 1 1-2 0a1 1 0 0 1 2 0M14 5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonsLock(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 4.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2m0 1.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5m-.75 1.25a.75.75 0 0 1 .75-.75c1.238 0 2.363.385 3.178.962C15.496 8.04 16 8.81 16 9.577a1.923 1.923 0 0 1-1.923 1.923h-.327a.75.75 0 0 1 0-1.5h.327c.234 0 .423-.19.423-.423c0-.105-.099-.474-.688-.89C13.257 8.293 12.437 8 11.5 8a.75.75 0 0 1-.75-.75M2.188 8.686C2.743 8.294 3.563 8 4.5 8a.75.75 0 0 0 0-1.5c-1.238 0-2.363.385-3.178.962C.504 8.04 0 8.81 0 9.577C0 10.639.861 11.5 1.923 11.5h.327a.75.75 0 0 0 0-1.5h-.327a.423.423 0 0 1-.423-.423c0-.105.099-.474.688-.89ZM4.5 4.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2m0 1.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5m6.25 3.5h-.25V9a2.5 2.5 0 1 0-5 0v.5h-.25c-.69 0-1.25.56-1.25 1.25v3c0 .69.56 1.25 1.25 1.25h5.5c.69 0 1.25-.56 1.25-1.25v-3c0-.69-.56-1.25-1.25-1.25M9 9v.5H7V9a1 1 0 1 1 2 0m-3.5 2v2.5h5V11z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Picture(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 3h-7A1.5 1.5 0 0 0 3 4.5v5.027l.962-.7a1.75 1.75 0 0 1 2.079.016l.928.696l2.368-2.03a1.75 1.75 0 0 1 2.325.043L13 8.787V4.5A1.5 1.5 0 0 0 11.5 3m3 7.498V4.5a3 3 0 0 0-3-3h-7a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h7a3 3 0 0 0 3-3zm-1.5.33l-2.355-2.174a.25.25 0 0 0-.332-.006L7.488 11.07l-.457.392l-.481-.361l-1.41-1.057a.25.25 0 0 0-.296-.002L3 11.381v.119A1.5 1.5 0 0 0 4.5 13h7a1.5 1.5 0 0 0 1.5-1.5zM7.5 6a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.922 3.7L7.81 4.81L6.444 6.179a3.326 3.326 0 0 0 3.378 3.378L11.19 8.19l1.11-1.11A2.39 2.39 0 0 0 8.922 3.7M7 3.5l-2 2l-2.36 2.36a3.89 3.89 0 1 0 5.5 5.5L10.5 11l2-2l.86-.86a3.889 3.889 0 1 0-5.5-5.5zm.078 8.8l1.413-1.412A4.852 4.852 0 0 1 5.11 7.51L3.7 8.922A2.39 2.39 0 0 0 7.078 12.3m.892-6.33a.75.75 0 1 0 1.06 1.06l1.5-1.5a.75.75 0 1 0-1.06-1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pin(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.5 2.255v-.01c.003-.03.013-.157-.361-.35C9.703 1.668 8.966 1.5 8 1.5c-.967 0-1.703.169-2.138.394c-.375.194-.365.32-.362.351v.01c-.003.03-.013.157.362.35C6.297 2.832 7.033 3 8 3c.967 0 1.703-.169 2.139-.394c.374-.194.364-.32.361-.351M8 4.5c.506 0 .99-.04 1.436-.118l.84 2.352l.253.707l.717.221c.648.2 1.055.44 1.277.65c.192.18.227.31.227.438c0 .14-.055.488-.937.878c-.869.384-2.2.622-3.813.622s-2.944-.238-3.813-.622c-.882-.39-.937-.738-.937-.878c0-.128.035-.259.227-.439c.222-.209.629-.448 1.277-.649l.717-.221l.253-.707l.84-2.352c.445.079.93.118 1.436.118m4-2.25c0 .738-.433 1.294-1.136 1.669l.825 2.31c1.553.48 2.561 1.32 2.561 2.52c0 1.854-2.402 2.848-5.5 2.985V15a.75.75 0 0 1-1.5 0v-3.266c-3.098-.136-5.5-1.131-5.5-2.984c0-1.2 1.008-2.04 2.561-2.52l.825-2.311C4.433 3.544 4 2.988 4 2.25C4 .75 5.79 0 8 0s4 .75 4 2.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.5 2.255v-.01c.003-.03.013-.157-.361-.35C9.703 1.668 8.966 1.5 8 1.5c-.967 0-1.703.169-2.138.394c-.375.194-.365.32-.362.351v.01c-.003.03-.013.157.362.35C6.297 2.832 7.033 3 8 3c.967 0 1.703-.169 2.139-.394c.374-.194.364-.32.361-.351M12 2.25c0 .738-.433 1.294-1.136 1.669l.825 2.31c1.553.48 2.561 1.32 2.561 2.52c0 1.854-2.402 2.848-5.5 2.985V15a.75.75 0 0 1-1.5 0v-3.266c-3.098-.136-5.5-1.131-5.5-2.984c0-1.2 1.008-2.04 2.561-2.52l.825-2.311C4.433 3.544 4 2.988 4 2.25C4 .75 5.79 0 8 0s4 .75 4 2.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinSlash(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.5 2.255v-.01c.003-.03.013-.157-.361-.35C9.703 1.668 8.966 1.5 8 1.5c-.967 0-1.703.169-2.138.394c-.375.194-.365.32-.362.351v.01c0 .008-.002.024.005.048L4.38 1.179C5.02.393 6.4 0 8 0c2.21 0 4 .75 4 2.25c0 .738-.433 1.294-1.136 1.669l.825 2.31c1.553.48 2.561 1.32 2.561 2.52c0 .617-.266 1.138-.734 1.565l-1.061-1.061c.27-.227.295-.41.295-.503c0-.128-.035-.259-.226-.439c-.223-.209-.63-.448-1.278-.649l-.552-.17l-.228-.228l-.19-.53l-.84-2.352a8.304 8.304 0 0 1-1.739.113l-1.96-1.96c.037.023.077.046.125.07C6.297 2.832 7.033 3 8 3c.967 0 1.703-.169 2.139-.394c.374-.194.364-.32.361-.351m-.008 9.297l2.478 2.478a.75.75 0 1 0 1.06-1.06l-11-11a.75.75 0 0 0-1.06 1.06l2.567 2.567l-.226.632C2.758 6.71 1.75 7.55 1.75 8.75c0 1.854 2.402 2.848 5.5 2.985V15a.75.75 0 0 0 1.5 0v-3.266a13.889 13.889 0 0 0 1.742-.182m-1.345-1.345L5.71 6.771l-.239.67l-.717.221c-.648.2-1.055.44-1.277.65c-.192.18-.227.31-.227.438c0 .14.055.488.937.878c.869.384 2.2.622 3.813.622c.4 0 .784-.015 1.147-.043" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinSlashFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.5 2.255v-.01c.003-.03.013-.157-.361-.35C9.703 1.668 8.966 1.5 8 1.5c-.967 0-1.703.169-2.138.394c-.375.194-.365.32-.362.351v.01c0 .008-.002.024.005.048L4.38 1.179C5.02.393 6.4 0 8 0c2.21 0 4 .75 4 2.25c0 .738-.433 1.294-1.136 1.669l.825 2.31c1.553.48 2.561 1.32 2.561 2.52c0 .617-.266 1.138-.734 1.565L5.738 2.536c.035.022.076.045.124.07C6.297 2.83 7.033 3 8 3c.967 0 1.703-.169 2.139-.394c.374-.194.364-.32.361-.351M4.537 5.597l-.226.632C2.758 6.71 1.75 7.55 1.75 8.75c0 1.854 2.402 2.848 5.5 2.985V15a.75.75 0 0 0 1.5 0v-3.266a13.889 13.889 0 0 0 1.742-.182l2.478 2.478a.75.75 0 1 0 1.06-1.06l-11-11a.75.75 0 0 0-1.06 1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pipeline(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 1.75A.75.75 0 0 1 .75 1h6a.75.75 0 0 1 0 1.5H6.5v2.25a.5.5 0 0 0 .5.5h2.25A5.75 5.75 0 0 1 15 11v2.5h.25a.75.75 0 0 1 0 1.5h-6a.75.75 0 0 1 0-1.5h.25v-2.25a.5.5 0 0 0-.5-.5H6.75A5.75 5.75 0 0 1 1 5V2.5H.75A.75.75 0 0 1 0 1.75M11 13.5h2.5V11a4.25 4.25 0 0 0-4.25-4.25h-.5v2.5H9a2 2 0 0 1 2 2zm-8.5-11H5v2.25a2 2 0 0 0 2 2h.25v2.5h-.5A4.25 4.25 0 0 1 2.5 5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlanetEarth(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 13.5c1.63 0 3.094-.709 4.101-1.835A2.501 2.501 0 0 1 10.25 9.25V9a.75.75 0 0 0-.75-.75a2.25 2.25 0 0 1 0-4.5a.75.75 0 0 0 .75-.75v-.02a5.502 5.502 0 0 0-7.471 3.287A2.25 2.25 0 0 1 4.75 8.5c0 .414.336.75.75.75a2.25 2.25 0 0 1 1.265 4.11c.397.092.81.14 1.235.14m-3.491-1.25H5.5a.75.75 0 0 0 0-1.5A2.25 2.25 0 0 1 3.25 8.5a.75.75 0 0 0-.744-.75a5.489 5.489 0 0 0 2.003 4.5m8.241-2h.27A5.48 5.48 0 0 0 13.5 8c0-1.665-.74-3.158-1.91-4.166A2.25 2.25 0 0 1 9.5 5.25a.75.75 0 0 0 0 1.5A2.25 2.25 0 0 1 11.75 9v.25a1 1 0 0 0 1 1M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiPlay0)"><path fill="currentColor" fill-rule="evenodd" d="M14.005 7.134L5.5 2.217a1 1 0 0 0-1.5.866v9.834a1 1 0 0 0 1.5.866l8.505-4.917a1 1 0 0 0 0-1.732m.751 3.03c1.665-.962 1.665-3.366 0-4.329L6.251.918C4.585-.045 2.5 1.158 2.5 3.083v9.834c0 1.925 2.085 3.128 3.751 2.164z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiPlay0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiPlayFill0)"><path fill="currentColor" fill-rule="evenodd" d="M14.756 10.164c1.665-.962 1.665-3.366 0-4.329L6.251.918C4.585-.045 2.5 1.158 2.5 3.083v9.834c0 1.925 2.085 3.128 3.751 2.164z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiPlayFill0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlugConnection(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.53 1.53A.75.75 0 0 0 14.47.47l-1.29 1.29a4.243 4.243 0 0 0-5.423.483l-.58.58a.958.958 0 0 0 0 1.354l4.646 4.646a.958.958 0 0 0 1.354 0l.58-.58a4.243 4.243 0 0 0 .484-5.423zm-8.5 4.94a.75.75 0 0 1 0 1.06L5.78 8.78l1.44 1.44l1.25-1.25a.75.75 0 0 1 1.06 1.06l-1.25 1.25l.543.543a.958.958 0 0 1 0 1.354l-.58.58a4.243 4.243 0 0 1-5.423.484l-1.29 1.29A.75.75 0 0 1 .47 14.47l1.29-1.29a4.243 4.243 0 0 1 .483-5.423l.58-.58a.958.958 0 0 1 1.354 0l.543.543l1.25-1.25a.75.75 0 0 1 1.06 0M3.5 8.62l-.197.197a2.743 2.743 0 0 0 3.879 3.879l.197-.197L3.5 8.621Zm9.197-1.439l-.197.197L8.621 3.5l.197-.197a2.743 2.743 0 0 1 3.879 3.879Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlugWire(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiPlugWire0)"><path fill="currentColor" fill-rule="evenodd" d="M11.89 4.111a5.5 5.5 0 0 0-7.78 0c-2.599 2.6-2.12 5.657-1.06 6.718a.75.75 0 0 0 1.06 0l.925-.925a3.09 3.09 0 0 1-.173-2.843l.547-1.23a1 1 0 0 1 1.62-.302l.19.19l.75-.75A.75.75 0 1 1 9.03 6.03l-.75.75l.94.94l.75-.75a.75.75 0 1 1 1.06 1.06l-.75.75l.19.19a1 1 0 0 1-.3 1.621l-1.231.547a3.089 3.089 0 0 1-2.843-.173l-.924.924a2.25 2.25 0 0 1-3.182 0C.222 10.122-.007 6.107 3.05 3.05a7 7 0 1 1 2.983 11.67a.75.75 0 1 1 .42-1.44a5.5 5.5 0 0 0 5.436-9.168Zm-5.658 3.56l.252-.566l2.411 2.411l-.566.252A1.589 1.589 0 0 1 6.232 7.67Z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiPlugWire0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1.75a.75.75 0 0 1 .75.75v4.75h4.75a.75.75 0 0 1 0 1.5H8.75v4.75a.75.75 0 0 1-1.5 0V8.75H2.5a.75.75 0 0 1 0-1.5h4.75V2.5A.75.75 0 0 1 8 1.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Power(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.75 1.75a.75.75 0 0 0-1.5 0v5.5a.75.75 0 0 0 1.5 0zM4.92 3.442A.75.75 0 1 0 4.08 2.2a7 7 0 1 0 7.841 0a.75.75 0 1 0-.841 1.242a5.5 5.5 0 1 1-6.159 0Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Printer(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.75 2.5h-5.5a.75.75 0 0 0-.75.75V4h7v-.75a.75.75 0 0 0-.75-.75M13 4v-.75A2.25 2.25 0 0 0 10.75 1h-5.5A2.25 2.25 0 0 0 3 3.25V4a3 3 0 0 0-3 3v2a3 3 0 0 0 3 3h1v1a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-1h1a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3m-9 6v.5H3A1.5 1.5 0 0 1 1.5 9V7A1.5 1.5 0 0 1 3 5.5h10A1.5 1.5 0 0 1 14.5 7v2a1.5 1.5 0 0 1-1.5 1.5h-1V10a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2m6-.5H6a.5.5 0 0 0-.5.5v3a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-3a.5.5 0 0 0-.5-.5m2.5-1a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pulse(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiPulse0)"><path fill="currentColor" fill-rule="evenodd" d="M9.235 1a.75.75 0 0 1 .74.56l2.034 7.726l1.09-1.908A.75.75 0 0 1 13.75 7h1.5a.75.75 0 0 1 0 1.5h-1.065l-1.784 3.122a.75.75 0 0 1-1.376-.181l-1.71-6.496l-2.083 9.466a.75.75 0 0 1-1.446.07L3.544 7.55l-.65 1.085A.75.75 0 0 1 2.25 9H.75a.75.75 0 1 1 0-1.5h1.075l1.282-2.136a.75.75 0 0 1 1.357.155l1.898 5.868l2.156-9.798A.75.75 0 0 1 9.235 1" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiPulse0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Puzzle(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.731 4H4.5A1.5 1.5 0 0 0 3 5.5v.377a2.72 2.72 0 0 1 0 5.246v.377A1.5 1.5 0 0 0 4.5 13h.377a2.72 2.72 0 0 1 5.246 0h.377a1.5 1.5 0 0 0 1.5-1.5v-1.232l1-.353a1.501 1.501 0 0 0 0-2.83l-1-.354V5.5A1.5 1.5 0 0 0 10.5 4H9.269l-.354-1a1.501 1.501 0 0 0-2.83 0zM8.9 14.5l-.204-1.02a1.22 1.22 0 0 0-2.392 0L6.1 14.5H4.5a3 3 0 0 1-3-3V9.9l1.02-.204a1.22 1.22 0 0 0 0-2.392L1.5 7.1V5.5a3 3 0 0 1 3-3h.17a3.001 3.001 0 0 1 5.66 0h.17a3 3 0 0 1 3 3v.17a3.001 3.001 0 0 1 0 5.66v.17a3 3 0 0 1-3 3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QrCode(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.75 2.25a.75.75 0 0 0-1.5 0v5h-5a.75.75 0 0 0 0 1.5h5.5a1 1 0 0 0 1-1zM7.25 11a1 1 0 0 1 1-1H11a.75.75 0 0 1 0 1.5H8.75v2.25a.75.75 0 0 1-1.5 0zM13 13h-2.25a.75.75 0 0 0 0 1.5h2.75a1 1 0 0 0 1-1V8.25a1 1 0 0 0-1-1h-2.75a.75.75 0 0 0 0 1.5H13zM3 4.5V3h1.5v1.5zm-1.5-2a1 1 0 0 1 1-1H5a1 1 0 0 1 1 1V5a1 1 0 0 1-1 1H2.5a1 1 0 0 1-1-1zm1.5 9V13h1.5v-1.5zM2.5 10a1 1 0 0 0-1 1v2.5a1 1 0 0 0 1 1H5a1 1 0 0 0 1-1V11a1 1 0 0 0-1-1zm9-5.5V3H13v1.5zm-1.5-2a1 1 0 0 1 1-1h2.5a1 1 0 0 1 1 1V5a1 1 0 0 1-1 1H11a1 1 0 0 1-1-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuoteClose(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiQuoteClose0)"><path fill="currentColor" fill-rule="evenodd" d="M3.589 9.67A2.751 2.751 0 0 1 1.5 7v-.25A2.75 2.75 0 0 1 4.25 4h.25a2.747 2.747 0 0 1 2.748 2.657v.003h.002V7c0 .131-.01.26-.027.386c-.02.261-.05.518-.09.77a8.028 8.028 0 0 1-.559 1.918a7.208 7.208 0 0 1-2.162 2.801a7.043 7.043 0 0 1-.098.076a.237.237 0 0 1-.144.049a.251.251 0 0 1-.22-.367a17.026 17.026 0 0 0 .65-1.384c.197-.474.396-1.013.557-1.578a2.712 2.712 0 0 1-.657.079h-.25c-.228 0-.45-.028-.661-.08m-.549 1.405A4.252 4.252 0 0 1 0 7v-.25A4.25 4.25 0 0 1 4.25 2.5h.25c1.452 0 2.733.728 3.5 1.838A4.245 4.245 0 0 1 11.5 2.5h.25A4.25 4.25 0 0 1 16 6.664V7c0 .183-.012.365-.035.543c-.207 2.62-1.358 4.966-3.488 6.599a1.738 1.738 0 0 1-1.057.358c-1.341 0-2.146-1.425-1.548-2.564c.111-.211.26-.508.418-.86a4.252 4.252 0 0 1-2.005-1.297a8.763 8.763 0 0 1-3.058 4.363a1.738 1.738 0 0 1-1.057.358c-1.341 0-2.146-1.425-1.548-2.564c.111-.211.26-.508.418-.86Zm7.8-1.405c.211.052.433.08.661.08h.25a2.76 2.76 0 0 0 .657-.079a14.405 14.405 0 0 1-.68 1.865a17.927 17.927 0 0 1-.527 1.097a.251.251 0 0 0 .22.367a.238.238 0 0 0 .144-.049l.098-.076a7.207 7.207 0 0 0 2.221-2.94a8.027 8.027 0 0 0 .5-1.779a9.01 9.01 0 0 0 .09-.77A2.73 2.73 0 0 0 14.501 7v-.339H14.5v-.004A2.747 2.747 0 0 0 11.75 4h-.251a2.75 2.75 0 0 0-2.75 2.75V7c0 1.29.89 2.374 2.089 2.67Z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiQuoteClose0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuoteOpen(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiQuoteOpen0)"><path fill="currentColor" fill-rule="evenodd" d="M12.411 6.33A2.751 2.751 0 0 1 14.5 9v.25A2.75 2.75 0 0 1 11.75 12h-.25a2.747 2.747 0 0 1-2.748-2.657V9.34H8.75V9c0-.131.01-.26.027-.386c.02-.261.05-.518.09-.77a8.028 8.028 0 0 1 .559-1.918a7.207 7.207 0 0 1 2.162-2.801l.098-.076A.238.238 0 0 1 11.83 3c.186 0 .306.202.22.367a18.552 18.552 0 0 0-.22.433a17.949 17.949 0 0 0-.43.951a14.39 14.39 0 0 0-.557 1.578l.054-.013a2.76 2.76 0 0 1 .603-.066h.25c.228 0 .45.028.661.08m.549-1.405A4.252 4.252 0 0 1 16 9v.25a4.25 4.25 0 0 1-4.25 4.25h-.25A4.245 4.245 0 0 1 8 11.662A4.245 4.245 0 0 1 4.5 13.5h-.25A4.25 4.25 0 0 1 0 9.336V9c0-.183.012-.365.035-.543c.207-2.62 1.358-4.966 3.488-6.599A1.738 1.738 0 0 1 4.58 1.5c1.341 0 2.146 1.425 1.548 2.564c-.111.211-.26.508-.418.86c.788.234 1.481.69 2.005 1.297a8.763 8.763 0 0 1 3.058-4.363A1.738 1.738 0 0 1 11.83 1.5c1.341 0 2.146 1.425 1.548 2.564c-.111.211-.26.508-.418.86ZM5.16 6.33a2.756 2.756 0 0 0-.661-.08h-.25a2.76 2.76 0 0 0-.657.079a14.398 14.398 0 0 1 .68-1.865A17.736 17.736 0 0 1 4.8 3.367A.251.251 0 0 0 4.58 3a.238.238 0 0 0-.144.049a7.737 7.737 0 0 0-.93.844a7.208 7.208 0 0 0-1.39 2.172a8.029 8.029 0 0 0-.498 1.779a8.753 8.753 0 0 0-.091.77A2.773 2.773 0 0 0 1.5 9v.339h.001v.004A2.747 2.747 0 0 0 4.25 12h.251a2.75 2.75 0 0 0 2.75-2.75V9c0-1.29-.89-2.374-2.089-2.67Z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiQuoteOpen0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Receipt(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.3 2L3.9.5L2 2v13.5l1.5-.776L4.9 14l2.4 1.5L9.7 14l2.4 1.5L14 14V.5l-1.5.776L11.1 2L8.7.5zm2.4.269L7.095 3.272l-.795.497l-.795-.497l-1.504-.94l-.501.395v10.308l.71-.367l.76-.393l.725.453L7.3 13.731l1.605-1.003l.795-.497l.795.497l1.504.94l.501-.395V2.965l-.71.367l-.76.393l-.725-.453zM5 6.5a.75.75 0 0 1 .75-.75h4.5a.75.75 0 0 1 0 1.5h-4.5A.75.75 0 0 1 5 6.5m.75 2.25a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RectanglePulse(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiRectanglePulse0)"><path fill="currentColor" fill-rule="evenodd" d="M12.5 12.5h-9A1.5 1.5 0 0 1 2 11V9h3a.75.75 0 0 0 .648-.372l1.02-1.748l1.9 4.18a.75.75 0 0 0 1.33.068L11.43 8.5H14V11a1.5 1.5 0 0 1-1.5 1.5M14 7h-3a.75.75 0 0 0-.648.372L9.332 9.12l-1.9-4.18a.75.75 0 0 0-1.33-.068L4.57 7.5H2V5a1.5 1.5 0 0 1 1.5-1.5h9A1.5 1.5 0 0 1 14 5zM3.5 14h9a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3h-9a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiRectanglePulse0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RectanglesFour(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 3h2a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1-.5-.5v-2a.5.5 0 0 1 .5-.5m-2 .5a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2zm9-.5h2a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1-.5-.5v-2a.5.5 0 0 1 .5-.5m-2 .5a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2zm-3 6.5h-2a.5.5 0 0 0-.5.5v2a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 0-.5-.5m-2-1.5a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2zm7 1.5h2a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1-.5-.5v-2a.5.5 0 0 1 .5-.5m-2 .5a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rocket(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiRocket0)"><path fill="currentColor" fill-rule="evenodd" d="M15.37 1.268a.75.75 0 0 0-.62-.634a7.745 7.745 0 0 0-7.516 3.055l-.156.212l-2.59.112a.75.75 0 0 0-.433.16L.696 6.827a.75.75 0 0 0 .206 1.292L4.25 9.352c.673.273 1.13.56 1.484.913c.352.353.64.811.913 1.484l1.234 3.35a.75.75 0 0 0 1.292.205l2.652-3.36a.75.75 0 0 0 .16-.431l.113-2.591l.212-.156a7.745 7.745 0 0 0 3.058-7.498ZM4.794 5.501l1.144-.05l-1.69 2.302l-1.572-.58zm4.032 7.822l-.58-1.572l2.302-1.69l-.05 1.145zm5.127-11.277a6.246 6.246 0 0 0-5.511 2.531l-2.78 3.786c.425.237.8.51 1.132.842c.332.332.606.707.842 1.133l3.786-2.78a6.246 6.246 0 0 0 2.53-5.512ZM2.378 13.952a5.36 5.36 0 0 1-.377.053a5.52 5.52 0 0 1 .05-.366c.104-.59.294-1.014.527-1.247c.244-.244.694-.274 1.004.036c.31.31.281.76.036 1.005c-.223.223-.644.413-1.24.519M.48 15.069a7.796 7.796 0 0 1 .025-1.18c.082-.838.33-1.876 1.012-2.557c.853-.854 2.253-.838 3.126.035c.873.874.89 2.273.036 3.126c-1.082 1.082-3.112 1.068-3.735 1.036a.487.487 0 0 1-.319-.145a.486.486 0 0 1-.145-.316Z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiRocket0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoundBrackets(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.22 2.28a.75.75 0 0 1 1.06-1.06a9.589 9.589 0 0 1 0 13.56a.75.75 0 1 1-1.06-1.06a8.089 8.089 0 0 0 0-11.44m-6.44 0a.75.75 0 0 0-1.06-1.06a9.589 9.589 0 0 0 0 13.56a.75.75 0 0 0 1.06-1.06a8.089 8.089 0 0 1 0-11.44" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Route(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.25 5a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5m0 1.5a2.751 2.751 0 0 1-2.646-2H4.577a.827.827 0 0 0-.238 1.619l7.753 2.326A2.327 2.327 0 0 1 11.423 13H6.396a2.751 2.751 0 1 1 0-1.5h5.027a.827.827 0 0 0 .238-1.619L3.908 7.556A2.327 2.327 0 0 1 4.577 3h5.027a2.751 2.751 0 1 1 2.646 3.5M5 12.25a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sack(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiSack0)"><path fill="currentColor" fill-rule="evenodd" d="M5.042.156a.75.75 0 0 1 .83-.057c.832.475 1.507.75 2.156.886c.644.134 1.304.139 2.116.023a.75.75 0 0 1 .83.94l-.562 2.06C12.957 5.758 14 10.142 14 12.75c0 2-3 3.25-6 3.25s-6-1.25-6-3.25c0-2.608 1.044-6.995 3.592-8.745L4.775.943a.75.75 0 0 1 .267-.787M7.992 5.6c.536.005 1.044-.116 1.605-.332c.55.385 1.046.983 1.486 1.785a3.895 3.895 0 0 1-2.96.706a.75.75 0 1 0-.246 1.48a5.408 5.408 0 0 0 3.83-.793c.543 1.494.793 3.122.793 4.303c0 .15-.114.568-1.046 1.033c-.875.438-2.136.717-3.454.717s-2.579-.279-3.454-.717C3.614 13.318 3.5 12.9 3.5 12.75c0-.41.03-.874.092-1.368a8.153 8.153 0 0 0 5.251 1.112a.75.75 0 1 0-.186-1.488c-1.719.215-3.356-.225-4.768-1.219c.199-.808.477-1.627.84-2.37c.485-.989 1.047-1.71 1.678-2.15c.549.2 1.053.329 1.586.334ZM6.646 2.135l.479 1.793c.364.12.633.17.882.173c.247.002.513-.042.87-.166l.368-1.347a6.67 6.67 0 0 1-2.599-.453" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiSack0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScalesBalanced(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.25 3h-4.5v10h2.5a.75.75 0 0 1 0 1.5h-6.5a.75.75 0 0 1 0-1.5h2.5V3h-4.5a.75.75 0 0 1 0-1.5h10.5a.75.75 0 0 1 0 1.5M1.643 8.568l.007.018a1.454 1.454 0 0 0 2.7 0l.007-.018L3 5.854zM3 4c-.46 0-.88.26-1.085.67L.2 8.099a1 1 0 0 0-.034.819l.09.226a2.954 2.954 0 0 0 5.486 0l.09-.226a1 1 0 0 0-.034-.819L4.085 4.671A1.214 1.214 0 0 0 3 4m8.643 4.568l.007.018a1.454 1.454 0 0 0 2.7 0l.007-.018L13 5.854zM13 4c-.46 0-.88.26-1.085.67L10.2 8.099a1 1 0 0 0-.034.819l.09.226a2.954 2.954 0 0 0 5.486 0l.09-.226a1 1 0 0 0-.034-.819l-1.714-3.427A1.214 1.214 0 0 0 13 4" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScalesUnbalanced(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.983.843a.75.75 0 0 1-.576.89l-4.657.998V13h2.5a.75.75 0 0 1 0 1.5h-6.5a.75.75 0 0 1 0-1.5h2.5V3.053l-4.343.93a.75.75 0 1 1-.314-1.466l10.5-2.25a.75.75 0 0 1 .89.576m-2.34 6.725l.007.018a1.454 1.454 0 0 0 2.7 0l.007-.018L13 4.854zM13 3c-.46 0-.88.26-1.085.67L10.2 7.099a1 1 0 0 0-.034.819l.09.226a2.954 2.954 0 0 0 5.486 0l.09-.226a1 1 0 0 0-.034-.819l-1.714-3.427A1.214 1.214 0 0 0 13 3M1.643 9.568l.007.018a1.454 1.454 0 0 0 2.7 0l.007-.018L3 6.854zM3 5c-.46 0-.88.26-1.085.67L.2 9.099a1 1 0 0 0-.034.819l.09.226a2.954 2.954 0 0 0 5.486 0l.09-.226a1 1 0 0 0-.034-.819L4.085 5.671A1.214 1.214 0 0 0 3 5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scissors(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 6a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m-3-1.5a3 3 0 0 0 4.524 2.585L6.939 8l-.915.915a3 3 0 1 0 1.06 1.06L8.122 8.94l5.501 3.209a.75.75 0 1 0 .756-1.296L9.488 8l4.89-2.852a.75.75 0 0 0-.756-1.296l-5.5 3.209l-1.037-1.037A3 3 0 1 0 1.5 4.5m3 5.5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SealCheck(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiSealCheck0)"><path fill="currentColor" fill-rule="evenodd" d="M5.515 3.5h.621l.44-.44l1.07-1.07a.5.5 0 0 1 .708 0l1.07 1.07l.44.44H12a.5.5 0 0 1 .5.5v2.136l.44.44l1.07 1.07a.5.5 0 0 1 0 .708l-1.07 1.07l-.44.44V12a.5.5 0 0 1-.5.5H9.865l-.44.44l-1.071 1.07a.5.5 0 0 1-.707 0l-1.072-1.07l-.439-.44H4a.5.5 0 0 1-.5-.5V9.864l-.44-.44l-1.07-1.07a.5.5 0 0 1 0-.708l1.07-1.07l.44-.44V4a.5.5 0 0 1 .5-.5zm3.9-2.571a2 2 0 0 0-2.83 0L5.516 2H4a2 2 0 0 0-2 2v1.515L.93 6.585a2 2 0 0 0 0 2.83L2 10.484V12a2 2 0 0 0 2 2h1.515l1.07 1.071a2 2 0 0 0 2.83 0L10.484 14H12a2 2 0 0 0 2-2v-1.515l1.071-1.07a2 2 0 0 0 0-2.83l-1.07-1.07V4a2 2 0 0 0-2-2h-1.516zM11.1 6.45a.75.75 0 0 0-1.2-.9L7.42 8.858L6.03 7.47a.75.75 0 0 0-1.06 1.06l2 2a.75.75 0 0 0 1.13-.08z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiSealCheck0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SealPercent(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiSealPercent0)"><path fill="currentColor" fill-rule="evenodd" d="M5.515 3.5h.621l.44-.44l1.07-1.07a.5.5 0 0 1 .708 0l1.07 1.07l.44.44H12a.5.5 0 0 1 .5.5v2.136l.44.44l1.07 1.07a.5.5 0 0 1 0 .708l-1.07 1.07l-.44.44V12a.5.5 0 0 1-.5.5H9.864l-.44.44l-1.07 1.07a.5.5 0 0 1-.708 0l-1.07-1.07l-.44-.44H4a.5.5 0 0 1-.5-.5V9.864l-.44-.44l-1.07-1.07a.5.5 0 0 1 0-.708l1.07-1.07l.44-.44V4a.5.5 0 0 1 .5-.5zm3.9-2.571a2 2 0 0 0-2.83 0L5.516 2H4a2 2 0 0 0-2 2v1.515L.929 6.585a2 2 0 0 0 0 2.83L2 10.484V12a2 2 0 0 0 2 2h1.515l1.07 1.071a2 2 0 0 0 2.83 0L10.484 14H12a2 2 0 0 0 2-2v-1.515l1.071-1.07a2 2 0 0 0 0-2.83L14 5.516V4a2 2 0 0 0-2-2h-1.515zM6.53 10.53l4-4a.75.75 0 1 0-1.06-1.06l-4 4a.75.75 0 1 0 1.06 1.06M11 10a1 1 0 1 1-2 0a1 1 0 0 1 2 0M6 7a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiSealPercent0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Server(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 3.5h8A1.5 1.5 0 0 1 13.5 5v2.25h-11V5A1.5 1.5 0 0 1 4 3.5M2.5 8.75V11A1.5 1.5 0 0 0 4 12.5h8a1.5 1.5 0 0 0 1.5-1.5V8.75zM1 5a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3zm2.75.5a.75.75 0 0 1 .75-.75H7a.75.75 0 0 1 0 1.5H4.5a.75.75 0 0 1-.75-.75m.75 4.25a.75.75 0 0 0 0 1.5H7a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShapesFour(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiShapes40)"><path fill="currentColor" fill-rule="evenodd" d="M6.193 6L4 2.075L1.807 6zM5.091.953a1.25 1.25 0 0 0-2.182 0L.289 5.64A1.25 1.25 0 0 0 1.38 7.5h5.24a1.25 1.25 0 0 0 1.09-1.86zM2.534 14.929a3.5 3.5 0 1 0 2.932-6.36a3.5 3.5 0 0 0-2.932 6.36M4 13.75a2 2 0 1 0 0-4a2 2 0 0 0 0 4M15.28 1.22a.75.75 0 0 1 0 1.06l-5 5a.75.75 0 1 1-1.06-1.06l5-5a.75.75 0 0 1 1.06 0m-5.03 12.03v-3.5h3.5v3.5zM8.75 9.5c0-.69.56-1.25 1.25-1.25h4c.69 0 1.25.56 1.25 1.25v4c0 .69-.56 1.25-1.25 1.25h-4c-.69 0-1.25-.56-1.25-1.25z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiShapes40"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShapesThree(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiShapes30)"><path fill="currentColor" fill-rule="evenodd" d="M8 2.075L10.193 6H5.807zM6.909.953a1.25 1.25 0 0 1 2.182 0l2.62 4.687A1.25 1.25 0 0 1 10.62 7.5H5.38a1.25 1.25 0 0 1-1.09-1.86zM4 15.25a3.5 3.5 0 1 1 .001-6.999A3.5 3.5 0 0 1 4 15.25m2-3.5a2 2 0 1 1-4 0a2 2 0 0 1 4 0m8 1.75h-3.5V10H14zM9 9.75c0-.69.56-1.25 1.25-1.25h4c.69 0 1.25.56 1.25 1.25v4c0 .69-.56 1.25-1.25 1.25h-4C9.56 15 9 14.44 9 13.75z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiShapes30"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shield(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m3.003 4.702l4.22-2.025a1.796 1.796 0 0 1 1.554 0l4.22 2.025a.886.886 0 0 1 .503.8V6a8.55 8.55 0 0 1-3.941 7.201l-.986.631a1.063 1.063 0 0 1-1.146 0l-.986-.63A8.55 8.55 0 0 1 2.5 6v-.498c0-.341.196-.652.503-.8m3.57-3.377L2.354 3.35A2.387 2.387 0 0 0 1 5.502V6a10.05 10.05 0 0 0 4.632 8.465l.986.63a2.563 2.563 0 0 0 2.764 0l.986-.63A10.05 10.05 0 0 0 15 6v-.498c0-.918-.526-1.755-1.354-2.152l-4.22-2.025a3.296 3.296 0 0 0-2.852 0ZM8.47 9.97a.75.75 0 1 0 1.06 1.06c.575-.574 1.118-1.398 1.516-2.195c.386-.772.704-1.653.704-2.335a.75.75 0 0 0-1.5 0c0 .318-.182.937-.546 1.665c-.352.703-.809 1.379-1.234 1.805" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldCheck(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m3.003 4.702l4.22-2.025a1.796 1.796 0 0 1 1.554 0l4.22 2.025a.886.886 0 0 1 .503.8V6a8.55 8.55 0 0 1-3.941 7.201l-.986.631a1.063 1.063 0 0 1-1.146 0l-.986-.63A8.55 8.55 0 0 1 2.5 6v-.498c0-.341.196-.652.503-.8m3.57-3.377L2.354 3.35A2.387 2.387 0 0 0 1 5.502V6a10.05 10.05 0 0 0 4.632 8.465l.986.63a2.563 2.563 0 0 0 2.764 0l.986-.63A10.05 10.05 0 0 0 15 6v-.498c0-.918-.526-1.755-1.354-2.152l-4.22-2.025a3.296 3.296 0 0 0-2.852 0ZM11.1 6.45a.75.75 0 1 0-1.2-.9L7.419 8.858L6.03 7.47a.75.75 0 0 0-1.06 1.06l2 2a.75.75 0 0 0 1.13-.08z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldExclamation(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m3.003 4.702l4.22-2.025a1.796 1.796 0 0 1 1.554 0l4.22 2.025a.886.886 0 0 1 .503.8V6a8.55 8.55 0 0 1-3.941 7.201l-.986.631a1.063 1.063 0 0 1-1.146 0l-.986-.63A8.55 8.55 0 0 1 2.5 6v-.498c0-.341.196-.652.503-.8m3.57-3.377L2.354 3.35A2.387 2.387 0 0 0 1 5.502V6a10.05 10.05 0 0 0 4.632 8.465l.986.63a2.563 2.563 0 0 0 2.764 0l.986-.63A10.05 10.05 0 0 0 15 6v-.498c0-.918-.526-1.755-1.354-2.152l-4.22-2.025a3.296 3.296 0 0 0-2.852 0ZM8 4.75a.75.75 0 0 1 .75.75v2a.75.75 0 0 1-1.5 0v-2A.75.75 0 0 1 8 4.75m1 5.75a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldKeyhole(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m3.003 4.702l4.22-2.025a1.796 1.796 0 0 1 1.554 0l4.22 2.025a.886.886 0 0 1 .503.8V6a8.55 8.55 0 0 1-3.941 7.201l-.986.631a1.063 1.063 0 0 1-1.146 0l-.986-.63A8.55 8.55 0 0 1 2.5 6v-.498c0-.341.196-.652.503-.8m3.57-3.377L2.354 3.35A2.387 2.387 0 0 0 1 5.502V6a10.05 10.05 0 0 0 4.632 8.465l.986.63a2.563 2.563 0 0 0 2.764 0l.986-.63A10.05 10.05 0 0 0 15 6v-.498c0-.918-.526-1.755-1.354-2.152l-4.22-2.025a3.296 3.296 0 0 0-2.852 0ZM9.5 7a1.5 1.5 0 0 1-.75 1.3v1.95a.75.75 0 0 1-1.5 0V8.3A1.5 1.5 0 1 1 9.5 7" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBag(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.174 3h5.652a1.5 1.5 0 0 1 1.49 1.328l.808 7A1.5 1.5 0 0 1 11.634 13H4.366a1.5 1.5 0 0 1-1.49-1.672l.808-7A1.5 1.5 0 0 1 5.174 3m-2.98 1.156A3 3 0 0 1 5.174 1.5h5.652a3 3 0 0 1 2.98 2.656l.808 7a3 3 0 0 1-2.98 3.344H4.366a3 3 0 0 1-2.98-3.344zM5 5.25a.75.75 0 0 1 1.5 0v.25a1.5 1.5 0 1 0 3 0v-.25a.75.75 0 0 1 1.5 0v.25a3 3 0 0 1-6 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBasket(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 4.5h1.307a2.193 2.193 0 0 1 2.109 2.795l-1.294 4.53A3 3 0 0 1 11.237 14H4.763a3 3 0 0 1-2.885-2.176L.584 7.295A2.193 2.193 0 0 1 2.693 4.5H4V3a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2zm-6.5 0V3a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5v1.5zm5.737 8H4.763a1.5 1.5 0 0 1-1.442-1.088L2.027 6.883A.693.693 0 0 1 2.693 6h10.614a.693.693 0 0 1 .666.883l-1.294 4.53a1.5 1.5 0 0 1-1.442 1.087m-4.59-1.265a.75.75 0 0 1-.882-.588l-.5-2.5a.75.75 0 0 1 1.47-.294l.5 2.5a.75.75 0 0 1-.588.882m4.088-3.088a.75.75 0 1 0-1.47-.294l-.5 2.5a.75.75 0 1 0 1.47.294z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCart(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.018 3.068L3.395 4.5L4.58 9.005a3 3 0 0 0 4.186 1.948l4.518-2.14A3 3 0 0 0 15 6.102V5a2 2 0 0 0-2-2H4.556l-.15-.535A2 2 0 0 0 2.48 1H.75a.75.75 0 0 0 0 1.5h1.73a.5.5 0 0 1 .482.366zm5.106 6.53l4.518-2.14a1.5 1.5 0 0 0 .858-1.356V5a.5.5 0 0 0-.5-.5H4.946L6.03 8.624a1.5 1.5 0 0 0 2.093.973ZM12 14.75a1.75 1.75 0 1 0 0-3.5a1.75 1.75 0 0 0 0 3.5M4.75 13a1.75 1.75 0 1 1-3.5 0a1.75 1.75 0 0 1 3.5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shuffle(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.75 12.5a.75.75 0 0 1 0-1.5h2.044c.86 0 1.644-.49 2.021-1.262L6.665 8l-.85-1.738A2.25 2.25 0 0 0 3.794 5H1.75a.75.75 0 1 1 0-1.5h2.044a3.75 3.75 0 0 1 3.369 2.103l.337.69l.337-.69A3.75 3.75 0 0 1 11.206 3.5h1.233l-.97-.97a.75.75 0 0 1 1.061-1.06l2.25 2.25a.75.75 0 0 1 0 1.06l-2.25 2.25a.75.75 0 1 1-1.06-1.06l.97-.97h-1.234c-.86 0-1.644.49-2.021 1.262l-2.022 4.135A3.75 3.75 0 0 1 3.794 12.5zm6.639-1.542l.696-1.424l.1.204A2.25 2.25 0 0 0 11.206 11h1.233l-.97-.97a.75.75 0 1 1 1.061-1.06l2.25 2.25a.75.75 0 0 1 0 1.06l-2.25 2.25a.75.75 0 1 1-1.06-1.06l.97-.97h-1.234a3.75 3.75 0 0 1-2.905-1.378c.031-.053.06-.108.088-.164" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Signal(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.261 2.186c.337-.274.829-.154 1.044.223c.197.344.09.777-.21 1.035A5.987 5.987 0 0 0 2 8a5.99 5.99 0 0 0 2.095 4.556c.3.258.407.69.21 1.035c-.215.377-.707.497-1.044.223A7.485 7.485 0 0 1 .5 8a7.485 7.485 0 0 1 2.761-5.814m8.434.223c-.197.344-.09.777.21 1.035A5.986 5.986 0 0 1 14 8a5.986 5.986 0 0 1-2.095 4.556c-.3.258-.407.69-.21 1.035c.215.377.707.497 1.044.223A7.485 7.485 0 0 0 15.5 8a7.485 7.485 0 0 0-2.761-5.814c-.337-.274-.829-.154-1.044.223M4.759 4.878c.315-.327.837-.21 1.062.184c.19.33.097.744-.144 1.04A2.988 2.988 0 0 0 5 8c0 .72.254 1.381.677 1.899c.241.295.333.708.144 1.04c-.225.393-.747.51-1.062.183A4.485 4.485 0 0 1 3.5 8c0-1.213.48-2.313 1.26-3.122Zm5.42.184c-.19.33-.098.744.144 1.04C10.746 6.618 11 7.28 11 8s-.254 1.381-.677 1.899c-.242.295-.333.708-.144 1.04c.225.393.747.51 1.062.183A4.484 4.484 0 0 0 12.5 8c0-1.213-.48-2.313-1.26-3.122c-.314-.327-.836-.21-1.061.184M8 9.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sliders(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0m1.405.75a3.001 3.001 0 0 1-5.81 0H1.747a.75.75 0 0 1 0-1.5h1.348a3.001 3.001 0 0 1 5.81 0h5.345a.75.75 0 0 1 0 1.5zm-7.158 4.5h5.345a3.001 3.001 0 0 1 5.811 0h1.347a.75.75 0 1 1 0 1.5h-1.347a3.001 3.001 0 0 1-5.81 0H1.746a.75.75 0 0 1 0-1.5Zm8.25-.75a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlidersVertical(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.998 8.499a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m-.75-1.406a3.001 3.001 0 0 0 0 5.811v1.347a.75.75 0 0 0 1.5 0v-1.347a3.001 3.001 0 0 0 0-5.811V1.748a.75.75 0 0 0-1.5 0zm-4.5 7.158V8.906a3.001 3.001 0 0 0 0-5.81V1.747a.75.75 0 1 0-1.5 0v1.347a3.001 3.001 0 0 0 0 5.811v5.345a.75.75 0 0 0 1.5 0Zm.75-8.25a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Smartphone(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 3.5v9a1.5 1.5 0 0 1-1.5 1.5h-5A1.5 1.5 0 0 1 4 12.5v-9A1.5 1.5 0 0 1 5.5 2h5A1.5 1.5 0 0 1 12 3.5m-1.5-3a3 3 0 0 1 3 3v9a3 3 0 0 1-3 3h-5a3 3 0 0 1-3-3v-9a3 3 0 0 1 3-3zM6.25 11a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Snail(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.864 1.107a.75.75 0 1 1 .768 1.29a1.982 1.982 0 0 0-.397.414c-.205.287-.432.748-.477 1.44a2.112 2.112 0 0 1 .388 2.78l-.028.043a2.277 2.277 0 0 0-.266 1.982l.471 1.414A3.44 3.44 0 0 1 11.06 15H2.332A1.333 1.333 0 0 1 1 13.667c0-1.223.822-2.253 1.944-2.568C2.344 10.27 2 9.15 2 8c0-1.412.458-2.674 1.349-3.587C4.243 3.498 5.509 3 7 3c1.522 0 2.643.802 3.246 1.93c.105-.169.23-.323.37-.46a3.463 3.463 0 0 1-.36-2a3.254 3.254 0 0 1 .293-.999l.036-.069l.013-.023l.005-.008l.002-.004l.001-.001v-.002a.75.75 0 1 1 1.285.778a1.762 1.762 0 0 0-.146.513a1.95 1.95 0 0 0 .207 1.132c.115-.02.232-.03.351-.034a4.122 4.122 0 0 1 .712-1.814a3.48 3.48 0 0 1 .548-.607a2.664 2.664 0 0 1 .263-.201l.022-.015l.01-.005l.003-.002l.002-.001zM7.75 10.75c.877 0 1.57-.38 2.06-.94a5.188 5.188 0 0 0 1.165-2.169L11.401 6a.998.998 0 0 1 1.09-.74a.6.6 0 0 1 .509.595a.63.63 0 0 1-.102.344l-.028.043a3.776 3.776 0 0 0-.44 3.289l.47 1.413a1.941 1.941 0 0 1-1.84 2.556H2.511c.08-.565.567-1 1.155-1h.678a.905.905 0 0 0 .436-1.7C4.131 10.438 3.5 9.374 3.5 8c0-1.088.348-1.95.922-2.538C4.993 4.877 5.852 4.5 7 4.5c1.346 0 2.25 1.056 2.25 2.5c0 .639-.162 1.235-.437 1.646c-.26.39-.604.604-1.063.604c-.87 0-1.25-.568-1.25-1c0-.551.38-.75.5-.75A.75.75 0 0 0 7 6C5.89 6 5 7.04 5 8.25c0 1.329 1.12 2.5 2.75 2.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Snowflake(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.719 1.22a.75.75 0 0 1 1.06 0L8 2.44l1.22-1.22a.75.75 0 1 1 1.06 1.06L8.75 3.81V6.7l2.502-1.445l.56-2.091a.75.75 0 1 1 1.45.388l-.447 1.667l1.667.447a.75.75 0 1 1-.388 1.449l-2.092-.56L9.5 7.998l2.502 1.445l2.091-.56a.75.75 0 0 1 .389 1.448l-1.667.447l.447 1.667a.75.75 0 1 1-1.45.388l-.56-2.091L8.75 9.298v2.891l1.53 1.53a.75.75 0 0 1-1.062 1.06L8 13.561l-1.218 1.218a.75.75 0 1 1-1.061-1.06l1.529-1.53V9.3l-2.504 1.445l-.56 2.088a.75.75 0 1 1-1.449-.388l.446-1.664l-1.664-.446a.75.75 0 0 1 .388-1.449l2.09.56L6.5 7.999L3.996 6.553l-2.089.56a.75.75 0 1 1-.388-1.449l1.664-.446l-.445-1.664a.75.75 0 0 1 1.448-.388l.56 2.088L7.25 6.7V3.81L5.719 2.28a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sphere(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M8 4a.75.75 0 0 0 0 1.5A2.5 2.5 0 0 1 10.5 8A.75.75 0 1 0 12 8a4 4 0 0 0-4-4" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Square(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 3h-7A1.5 1.5 0 0 0 3 4.5v7A1.5 1.5 0 0 0 4.5 13h7a1.5 1.5 0 0 0 1.5-1.5v-7A1.5 1.5 0 0 0 11.5 3m-7-1.5a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h7a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareArticle(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 3h-7A1.5 1.5 0 0 0 3 4.5v7A1.5 1.5 0 0 0 4.5 13h7a1.5 1.5 0 0 0 1.5-1.5v-7A1.5 1.5 0 0 0 11.5 3m-7-1.5a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h7a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3zm6 6H5.43a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h5.07a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1m-5.32-3h3.57a.75.75 0 0 1 0 1.5H5.18a.75.75 0 0 1 0-1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareBars(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 3h7A1.5 1.5 0 0 1 13 4.5v7a1.5 1.5 0 0 1-1.5 1.5h-7A1.5 1.5 0 0 1 3 11.5v-7A1.5 1.5 0 0 1 4.5 3m-3 1.5a3 3 0 0 1 3-3h7a3 3 0 0 1 3 3v7a3 3 0 0 1-3 3h-7a3 3 0 0 1-3-3zm3.75 5.498a.75.75 0 0 0 0 1.5h5.502a.75.75 0 0 0 0-1.5zM4.5 8a.75.75 0 0 1 .75-.75h5.502a.75.75 0 0 1 0 1.5H5.25A.75.75 0 0 1 4.5 8m.75-3.498a.75.75 0 0 0 0 1.5h5.502a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareBarsVertical(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 3h7A1.5 1.5 0 0 1 13 4.5v7a1.5 1.5 0 0 1-1.5 1.5h-7A1.5 1.5 0 0 1 3 11.5v-7A1.5 1.5 0 0 1 4.5 3m-3 1.5a3 3 0 0 1 3-3h7a3 3 0 0 1 3 3v7a3 3 0 0 1-3 3h-7a3 3 0 0 1-3-3zm8.5 6.251a.75.75 0 1 0 1.5 0V5.249a.75.75 0 0 0-1.5 0zm-2 .75a.75.75 0 0 1-.75-.75V5.249a.75.75 0 0 1 1.5 0v5.502a.75.75 0 0 1-.75.75m-3.497-.75a.75.75 0 0 0 1.5 0V5.249a.75.75 0 1 0-1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareBrackets(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.75 1A2.25 2.25 0 0 0 1.5 3.25v9.5A2.25 2.25 0 0 0 3.75 15h.5a.75.75 0 0 0 0-1.5h-.5a.75.75 0 0 1-.75-.75v-9.5a.75.75 0 0 1 .75-.75h.5a.75.75 0 0 0 0-1.5zm8.5 14a2.25 2.25 0 0 0 2.25-2.25v-9.5A2.25 2.25 0 0 0 12.25 1h-.5a.75.75 0 0 0 0 1.5h.5a.75.75 0 0 1 .75.75v9.5a.75.75 0 0 1-.75.75h-.5a.75.75 0 0 0 0 1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareBracketsBarsVertical(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 4.25a.75.75 0 0 1 .75-.75h1.01a.75.75 0 0 0 0-1.5H3.25A2.25 2.25 0 0 0 1 4.25v7.5A2.25 2.25 0 0 0 3.25 14h1.01a.75.75 0 0 0 0-1.5H3.25a.75.75 0 0 1-.75-.75zM11.74 2a.75.75 0 0 0 0 1.5h1.01a.75.75 0 0 1 .75.75v7.5a.75.75 0 0 1-.75.75h-1.01a.75.75 0 0 0 0 1.5h1.01A2.25 2.25 0 0 0 15 11.75v-7.5A2.25 2.25 0 0 0 12.75 2zM8.75 5.75a.75.75 0 0 0-1.5 0v4.5a.75.75 0 0 0 1.5 0zm2-.75a.75.75 0 0 1 .75.75v4.5a.75.75 0 0 1-1.5 0v-4.5a.75.75 0 0 1 .75-.75M6 5.75a.75.75 0 0 0-1.5 0v4.5a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareBracketsLetterA(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.25 3.5a.75.75 0 0 0-.75.75v7.5c0 .414.336.75.75.75h1.01a.75.75 0 0 1 0 1.5H3.25A2.25 2.25 0 0 1 1 11.75v-7.5A2.25 2.25 0 0 1 3.25 2h1.01a.75.75 0 0 1 0 1.5zm7.74-.75a.75.75 0 0 1 .75-.75h1.01A2.25 2.25 0 0 1 15 4.25v7.5A2.25 2.25 0 0 1 12.75 14h-1.01a.75.75 0 0 1 0-1.5h1.01a.75.75 0 0 0 .75-.75v-7.5a.75.75 0 0 0-.75-.75h-1.01a.75.75 0 0 1-.75-.75m-4.2 2.588a1.291 1.291 0 0 1 2.418 0l1.743 4.649a.75.75 0 1 1-1.404.526L9.168 9.5H6.832l-.38 1.013a.75.75 0 1 1-1.404-.526l1.743-4.65ZM7.395 8h1.21L8 6.386z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareChartBar(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 3h7A1.5 1.5 0 0 1 13 4.5v7a1.5 1.5 0 0 1-1.5 1.5h-7A1.5 1.5 0 0 1 3 11.5v-7A1.5 1.5 0 0 1 4.5 3m-3 1.5a3 3 0 0 1 3-3h7a3 3 0 0 1 3 3v7a3 3 0 0 1-3 3h-7a3 3 0 0 1-3-3zm3.75 5.498a.75.75 0 0 0 0 1.5h2.002a.75.75 0 0 0 0-1.5zM4.5 8a.75.75 0 0 1 .75-.75h5.502a.75.75 0 0 1 0 1.5H5.25A.75.75 0 0 1 4.5 8m.75-3.498a.75.75 0 0 0 0 1.5h3.502a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareChartColumn(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 3h7A1.5 1.5 0 0 1 13 4.5v7a1.5 1.5 0 0 1-1.5 1.5h-7A1.5 1.5 0 0 1 3 11.5v-7A1.5 1.5 0 0 1 4.5 3m-3 1.5a3 3 0 0 1 3-3h7a3 3 0 0 1 3 3v7a3 3 0 0 1-3 3h-7a3 3 0 0 1-3-3zm8.499 6.251a.75.75 0 0 0 1.5 0V8.749a.75.75 0 0 0-1.5 0zm-1.998.75a.75.75 0 0 1-.75-.75V5.249a.75.75 0 1 1 1.5 0v5.502a.75.75 0 0 1-.75.75m-3.498-.75a.75.75 0 0 0 1.5 0V7.249a.75.75 0 1 0-1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareCheck(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 3h7A1.5 1.5 0 0 1 13 4.5v7a1.5 1.5 0 0 1-1.5 1.5h-7A1.5 1.5 0 0 1 3 11.5v-7A1.5 1.5 0 0 1 4.5 3m-3 1.5a3 3 0 0 1 3-3h7a3 3 0 0 1 3 3v7a3 3 0 0 1-3 3h-7a3 3 0 0 1-3-3zm10.092 1.46a.75.75 0 0 0-1.184-.92L7.43 8.869l-1.4-1.4A.75.75 0 0 0 4.97 8.53l2 2a.75.75 0 0 0 1.122-.07z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareDashed(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 3A1.5 1.5 0 0 0 3 4.5v1.75a.75.75 0 0 1-1.5 0V4.5a3 3 0 0 1 3-3h1.75a.75.75 0 0 1 0 1.5zM9 2.25a.75.75 0 0 1 .75-.75h1.75a3 3 0 0 1 3 3v1.75a.75.75 0 0 1-1.5 0V4.5A1.5 1.5 0 0 0 11.5 3H9.75A.75.75 0 0 1 9 2.25M2.25 9a.75.75 0 0 1 .75.75v1.75A1.5 1.5 0 0 0 4.5 13h1.75a.75.75 0 0 1 0 1.5H4.5a3 3 0 0 1-3-3V9.75A.75.75 0 0 1 2.25 9m11.5 0a.75.75 0 0 1 .75.75v1.75a3 3 0 0 1-3 3H9.75a.75.75 0 0 1 0-1.5h1.75a1.5 1.5 0 0 0 1.5-1.5V9.75a.75.75 0 0 1 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareDashedCircle(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.5A1.5 1.5 0 0 1 4.5 3h1.75a.75.75 0 0 0 0-1.5H4.5a3 3 0 0 0-3 3v1.75a.75.75 0 0 0 1.5 0zm6.75-3a.75.75 0 0 0 0 1.5h1.75A1.5 1.5 0 0 1 13 4.5v1.75a.75.75 0 0 0 1.5 0V4.5a3 3 0 0 0-3-3zM3 9.75a.75.75 0 0 0-1.5 0v1.75a3 3 0 0 0 3 3h1.75a.75.75 0 0 0 0-1.5H4.5A1.5 1.5 0 0 1 3 11.5zm11.5 0a.75.75 0 0 0-1.5 0v1.75a1.5 1.5 0 0 1-1.5 1.5H9.75a.75.75 0 0 0 0 1.5h1.75a3 3 0 0 0 3-3zM10 8a2 2 0 1 1-4 0a2 2 0 0 1 4 0m1.5 0a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareDashedLetterA(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 3A1.5 1.5 0 0 0 3 4.5v1.75a.75.75 0 0 1-1.5 0V4.5a3 3 0 0 1 3-3h1.75a.75.75 0 0 1 0 1.5zM9 2.25a.75.75 0 0 1 .75-.75h1.75a3 3 0 0 1 3 3v1.75a.75.75 0 0 1-1.5 0V4.5A1.5 1.5 0 0 0 11.5 3H9.75A.75.75 0 0 1 9 2.25M2.25 9a.75.75 0 0 1 .75.75v1.75A1.5 1.5 0 0 0 4.5 13h1.75a.75.75 0 0 1 0 1.5H4.5a3 3 0 0 1-3-3V9.75A.75.75 0 0 1 2.25 9m11.5 0a.75.75 0 0 1 .75.75v1.75a3 3 0 0 1-3 3H9.75a.75.75 0 0 1 0-1.5h1.75a1.5 1.5 0 0 0 1.5-1.5V9.75a.75.75 0 0 1 .75-.75M6.791 5.338a1.291 1.291 0 0 1 2.418 0l1.743 4.649a.75.75 0 1 1-1.404.526L9.168 9.5H6.832l-.38 1.013a.75.75 0 1 1-1.404-.526l1.743-4.65ZM7.395 8h1.21L8 6.386z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareDashedLetterT(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.5A1.5 1.5 0 0 1 4.5 3h1.75a.75.75 0 0 0 0-1.5H4.5a3 3 0 0 0-3 3v1.75a.75.75 0 0 0 1.5 0zm6.75-3a.75.75 0 0 0 0 1.5h1.75A1.5 1.5 0 0 1 13 4.5v1.75a.75.75 0 0 0 1.5 0V4.5a3 3 0 0 0-3-3zM3 9.75a.75.75 0 0 0-1.5 0v1.75a3 3 0 0 0 3 3h1.75a.75.75 0 0 0 0-1.5H4.5A1.5 1.5 0 0 1 3 11.5zm11.5 0a.75.75 0 0 0-1.5 0v1.75a1.5 1.5 0 0 1-1.5 1.5H9.75a.75.75 0 0 0 0 1.5h1.75a3 3 0 0 0 3-3zM5.752 5.002a.75.75 0 1 0 0 1.5h1.5l-.002 3.75a.75.75 0 1 0 1.5 0l.002-3.75h1.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareDashedText(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.5A1.5 1.5 0 0 1 4.5 3h1.75a.75.75 0 0 0 0-1.5H4.5a3 3 0 0 0-3 3v1.75a.75.75 0 0 0 1.5 0zm6.75-3a.75.75 0 0 0 0 1.5h1.75A1.5 1.5 0 0 1 13 4.5v1.75a.75.75 0 0 0 1.5 0V4.5a3 3 0 0 0-3-3zM3 9.75a.75.75 0 0 0-1.5 0v1.75a3 3 0 0 0 3 3h1.75a.75.75 0 0 0 0-1.5H4.5A1.5 1.5 0 0 1 3 11.5zm11.5 0a.75.75 0 0 0-1.5 0v1.75a1.5 1.5 0 0 1-1.5 1.5H9.75a.75.75 0 0 0 0 1.5h1.75a3 3 0 0 0 3-3zM5.752 4.502a.75.75 0 1 0 0 1.5h4.5a.75.75 0 0 0 0-1.5zm0 2.75a.75.75 0 1 0 0 1.5h4.5a.75.75 0 0 0 0-1.5zm0 2.75a.75.75 0 1 0 0 1.5h2.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareDot(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.25 1.5a.75.75 0 0 1 0 1.5H4.5A1.5 1.5 0 0 0 3 4.5v7A1.5 1.5 0 0 0 4.5 13h7a1.5 1.5 0 0 0 1.5-1.5V7.75a.75.75 0 0 1 1.5 0v3.75a3 3 0 0 1-3 3h-7a3 3 0 0 1-3-3v-7a3 3 0 0 1 3-3zm4.25 4a2 2 0 1 0 0-4a2 2 0 0 0 0 4" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareExclamation(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 3h-7A1.5 1.5 0 0 0 3 4.5v7A1.5 1.5 0 0 0 4.5 13h7a1.5 1.5 0 0 0 1.5-1.5v-7A1.5 1.5 0 0 0 11.5 3m-7-1.5a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h7a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3zm4.5 9a1 1 0 1 1-2 0a1 1 0 0 1 2 0M8.75 5a.75.75 0 0 0-1.5 0v2.5a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 1.5a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h7a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareHashtag(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 3h-7A1.5 1.5 0 0 0 3 4.5v7A1.5 1.5 0 0 0 4.5 13h7a1.5 1.5 0 0 0 1.5-1.5v-7A1.5 1.5 0 0 0 11.5 3m-7-1.5a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h7a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3zm5.623 2.76a.75.75 0 0 1 .617.863l-.105.627h.865a.75.75 0 0 1 0 1.5h-1.115l-.25 1.5H11a.75.75 0 0 1 0 1.5H9.885l-.145.873a.75.75 0 1 1-1.48-.246l.105-.627h-1.48l-.145.873a.75.75 0 1 1-1.48-.246l.105-.627H4.5a.75.75 0 0 1 0-1.5h1.115l.25-1.5H5a.75.75 0 0 1 0-1.5h1.115l.145-.873a.75.75 0 1 1 1.48.246l-.105.627h1.48l.145-.873a.75.75 0 0 1 .863-.617M7.135 8.75h1.48l.25-1.5h-1.48z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareLetterP(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 3h7A1.5 1.5 0 0 1 13 4.5v7a1.5 1.5 0 0 1-1.5 1.5h-7A1.5 1.5 0 0 1 3 11.5v-7A1.5 1.5 0 0 1 4.5 3m-3 1.5a3 3 0 0 1 3-3h7a3 3 0 0 1 3 3v7a3 3 0 0 1-3 3h-7a3 3 0 0 1-3-3zm5.25 0a.75.75 0 0 0-.75.75v5.5a.75.75 0 0 0 1.5 0V9.5h.75a2.5 2.5 0 0 0 0-5zM8.25 8H7.5V6h.75a1 1 0 0 1 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareLetterT(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 3h7A1.5 1.5 0 0 1 13 4.5v7a1.5 1.5 0 0 1-1.5 1.5h-7A1.5 1.5 0 0 1 3 11.5v-7A1.5 1.5 0 0 1 4.5 3m-3 1.5a3 3 0 0 1 3-3h7a3 3 0 0 1 3 3v7a3 3 0 0 1-3 3h-7a3 3 0 0 1-3-3zm4.5.25a.75.75 0 0 0 0 1.5h1.25v4.25a.75.75 0 0 0 1.5 0V6.25H10a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareListUl(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 3h7A1.5 1.5 0 0 1 13 4.5v7a1.5 1.5 0 0 1-1.5 1.5h-7A1.5 1.5 0 0 1 3 11.5v-7A1.5 1.5 0 0 1 4.5 3m-3 1.5a3 3 0 0 1 3-3h7a3 3 0 0 1 3 3v7a3 3 0 0 1-3 3h-7a3 3 0 0 1-3-3zm4.75.75a1 1 0 1 1-2 0a1 1 0 0 1 2 0m1 0A.75.75 0 0 1 8 4.5h2.75a.75.75 0 0 1 0 1.5H8a.75.75 0 0 1-.75-.75M5.25 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2m1 1.75a1 1 0 1 1-2 0a1 1 0 0 1 2 0M8 7.25a.75.75 0 0 0 0 1.5h2.75a.75.75 0 0 0 0-1.5zm-.75 3.5A.75.75 0 0 1 8 10h2.75a.75.75 0 0 1 0 1.5H8a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareMinus(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 3h7A1.5 1.5 0 0 1 13 4.5v7a1.5 1.5 0 0 1-1.5 1.5h-7A1.5 1.5 0 0 1 3 11.5v-7A1.5 1.5 0 0 1 4.5 3m-3 1.5a3 3 0 0 1 3-3h7a3 3 0 0 1 3 3v7a3 3 0 0 1-3 3h-7a3 3 0 0 1-3-3zm4 2.75a.75.75 0 0 0 0 1.5h5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquarePlus(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 3h7A1.5 1.5 0 0 1 13 4.5v7a1.5 1.5 0 0 1-1.5 1.5h-7A1.5 1.5 0 0 1 3 11.5v-7A1.5 1.5 0 0 1 4.5 3m-3 1.5a3 3 0 0 1 3-3h7a3 3 0 0 1 3 3v7a3 3 0 0 1-3 3h-7a3 3 0 0 1-3-3zm7.25 1a.75.75 0 0 0-1.5 0v1.75H5.5a.75.75 0 0 0 0 1.5h1.75v1.75a.75.75 0 0 0 1.5 0V8.75h1.75a.75.75 0 0 0 0-1.5H8.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareXmark(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 3h7A1.5 1.5 0 0 1 13 4.5v7a1.5 1.5 0 0 1-1.5 1.5h-7A1.5 1.5 0 0 1 3 11.5v-7A1.5 1.5 0 0 1 4.5 3m-3 1.5a3 3 0 0 1 3-3h7a3 3 0 0 1 3 3v7a3 3 0 0 1-3 3h-7a3 3 0 0 1-3-3zm5.03.97a.75.75 0 0 0-1.06 1.06L6.94 8L5.47 9.47a.75.75 0 1 0 1.06 1.06L8 9.06l1.47 1.47a.75.75 0 1 0 1.06-1.06L9.06 8l1.47-1.47a.75.75 0 1 0-1.06-1.06L8 6.94z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m9.194 5l.351.873l.94.064l3.197.217l-2.46 2.055l-.722.603l.23.914l.782 3.108l-2.714-1.704L8 10.629l-.798.5l-2.714 1.705l.782-3.108l.23-.914l-.723-.603l-2.46-2.055l3.198-.217l.94-.064l.35-.874L8 2.025zm-7.723-.292l3.943-.268L6.886.773C7.29-.231 8.71-.231 9.114.773l1.472 3.667l3.943.268c1.08.073 1.518 1.424.688 2.118L12.185 9.36l.964 3.832c.264 1.05-.886 1.884-1.802 1.31L8 12.4l-3.347 2.101c-.916.575-2.066-.26-1.802-1.309l.964-3.832L.783 6.826c-.83-.694-.391-2.045.688-2.118" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.886.773C7.29-.231 8.71-.231 9.114.773l1.472 3.667l3.943.268c1.08.073 1.518 1.424.688 2.118L12.185 9.36l.964 3.832c.264 1.05-.886 1.884-1.802 1.31L8 12.4l-3.347 2.101c-.916.575-2.066-.26-1.802-1.309l.964-3.832L.783 6.826c-.83-.694-.391-2.045.688-2.118l3.943-.268z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stethoscope(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.513 2.452a.75.75 0 0 0-.526-1.404l-1.027.385A2.25 2.25 0 0 0 .5 3.539V6a4.75 4.75 0 0 0 4.274 4.727a4.75 4.75 0 0 0 9.476-.477v-.364a2.501 2.501 0 1 0-1.5 0v.364a3.25 3.25 0 0 1-6.477.39A4.752 4.752 0 0 0 10 6V3.54a2.25 2.25 0 0 0-1.46-2.107l-1.027-.385a.75.75 0 0 0-.526 1.404l1.026.385a.75.75 0 0 1 .487.702V6A3.25 3.25 0 1 1 2 6V3.54a.75.75 0 0 1 .487-.703zM13.5 8.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sticker(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 3h-7A1.5 1.5 0 0 0 3 4.5v7A1.5 1.5 0 0 0 4.5 13h3v-2.5a3 3 0 0 1 3-3H13v-3A1.5 1.5 0 0 0 11.5 3m1.303 6a1.503 1.503 0 0 1-.242.318l-3.243 3.243a1.503 1.503 0 0 1-.318.242V10.5A1.5 1.5 0 0 1 10.5 9zm.818 1.379a3 3 0 0 0 .879-2.122V4.5a3 3 0 0 0-3-3h-7a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h3.757a3 3 0 0 0 2.122-.879z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stop(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 3h-7A1.5 1.5 0 0 0 3 4.5v7A1.5 1.5 0 0 0 4.5 13h7a1.5 1.5 0 0 0 1.5-1.5v-7A1.5 1.5 0 0 0 11.5 3m-7-1.5a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h7a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 1.5a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h7a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stopwatch(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.25 0a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5zM13 9A5 5 0 1 1 3 9a5 5 0 0 1 10 0m1.5 0a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0m-2.28-7.28a.75.75 0 0 1 1.06 0l1.5 1.5a.75.75 0 0 1-1.06 1.06l-1.5-1.5a.75.75 0 0 1 0-1.06M8.75 6a.75.75 0 0 0-1.5 0v3a.75.75 0 0 0 .3.6l1.333 1a.75.75 0 1 0 .9-1.2L8.75 8.625z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Strikethrough(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M5.502 2.757C6.214 2.236 7.122 2 7.984 2c1.685 0 3.015.572 3.687 1.915a.75.75 0 1 1-1.342.67C10.001 3.928 9.331 3.5 7.984 3.5c-.611 0-1.19.17-1.597.467a1.46 1.46 0 0 0-.627 1.242c0 .404.165.76.463 1.041H4.447a2.89 2.89 0 0 1-.187-1.04c0-1.084.507-1.916 1.242-2.453m6.047 5.993h1.201a.75.75 0 0 0 0-1.5h-9.5a.75.75 0 0 0 0 1.5h5.475l.043.012h.002c1.196.323 1.98 1.005 1.98 1.988c0 .669-.289 1.063-.742 1.33c-.5.296-1.222.437-2 .437c-1.398 0-2.453-.473-2.796-1.504a.75.75 0 0 0-1.424.474c.657 1.969 2.602 2.53 4.22 2.53c.915 0 1.94-.16 2.762-.644c.869-.513 1.48-1.377 1.48-2.623c0-.822-.276-1.482-.7-2Z"/><path d="M5.502 2.757C6.214 2.236 7.122 2 7.984 2c1.685 0 3.015.572 3.687 1.915a.75.75 0 1 1-1.342.67C10.001 3.93 9.331 3.5 7.984 3.5c-.611 0-1.19.17-1.597.468c-.384.28-.627.678-.627 1.242c0 .403.165.758.463 1.04H4.447a2.89 2.89 0 0 1-.187-1.04c0-1.084.507-1.916 1.242-2.453m6.047 5.993h1.201a.75.75 0 0 0 0-1.5h-9.5a.75.75 0 0 0 0 1.5h5.475l.043.012h.002c1.196.323 1.98 1.005 1.98 1.988c0 .669-.289 1.063-.742 1.33c-.5.296-1.222.437-2 .437c-1.398 0-2.453-.472-2.796-1.504a.75.75 0 1 0-1.424.474c.657 1.969 2.602 2.53 4.22 2.53c.915 0 1.94-.16 2.762-.644c.869-.512 1.48-1.377 1.48-2.623c0-.822-.276-1.481-.7-2Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Suitcase(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9 1.5H7a.5.5 0 0 0-.5.5v1h3V2a.5.5 0 0 0-.5-.5M5 2v1a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h6a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3V2a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2m4.5 2.5H11A1.5 1.5 0 0 1 12.5 6v6a1.5 1.5 0 0 1-1.5 1.5H5A1.5 1.5 0 0 1 3.5 12V6A1.5 1.5 0 0 1 5 4.5zM5.75 6a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sun(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 3a.75.75 0 0 1-.75-.75V.75a.75.75 0 1 1 1.5 0v1.5A.75.75 0 0 1 8 3m0 7.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5M8 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8m-.75 3.25a.75.75 0 0 0 1.5 0v-1.5a.75.75 0 0 0-1.5 0zM13 8a.75.75 0 0 1 .75-.75h1.5a.75.75 0 0 1 0 1.5h-1.5A.75.75 0 0 1 13 8M.75 7.25a.75.75 0 1 0 0 1.5h1.5a.75.75 0 0 0 0-1.5zm10.786-2.786a.75.75 0 0 1 0-1.06l1.06-1.06a.75.75 0 0 1 1.06 1.06l-1.06 1.06a.75.75 0 0 1-1.06 0m-9.193 8.132a.75.75 0 0 0 1.06 1.06l1.062-1.06a.75.75 0 0 0-1.061-1.06zm9.193-1.06a.75.75 0 0 1 1.06 0l1.06 1.06a.75.75 0 0 1-1.06 1.06l-1.06-1.06a.75.75 0 0 1 0-1.06M3.404 2.343a.75.75 0 0 0-1.06 1.06l1.06 1.061a.75.75 0 1 0 1.06-1.06l-1.06-1.06Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tshirt(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m15.177 3.088l-3.544-1.771A3 3 0 0 0 10.292 1h-.877a1.5 1.5 0 0 1-2.83 0h-.877a3 3 0 0 0-1.341.317L.823 3.088a1 1 0 0 0-.481 1.266l1.295 3.237a1 1 0 0 0 1.28.565l.351-.132l-.188 4.9A2 2 0 0 0 5.078 15h5.844a2 2 0 0 0 1.998-2.077l-.188-4.899l.352.132a1 1 0 0 0 1.28-.565l1.294-3.237a1 1 0 0 0-.481-1.266M5.764 2.5h-.056a1.5 1.5 0 0 0-.67.158L1.904 4.224l.943 2.356l2.006-.752l-.275 7.153a.5.5 0 0 0 .5.519h5.843a.5.5 0 0 0 .5-.52l-.276-7.152l2.006.752l.943-2.356l-3.132-1.566a1.5 1.5 0 0 0-.671-.158h-.056A2.99 2.99 0 0 1 8 3.5a2.99 2.99 0 0 1-2.236-1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m13.06 8.818l-4.869 4.87a1 1 0 0 1-1.408.006l-4.45-4.37a1 1 0 0 1-.012-1.414l4.868-4.96a1.5 1.5 0 0 1 1.07-.45H12.5a1 1 0 0 1 1 1v4.257a1.5 1.5 0 0 1-.44 1.061m-6.942-6.92A3 3 0 0 1 8.259 1H12.5A2.5 2.5 0 0 1 15 3.5v4.257a3 3 0 0 1-.879 2.122l-4.87 4.87a2.5 2.5 0 0 1-3.519.015l-4.45-4.37a2.5 2.5 0 0 1-.032-3.535L6.118 1.9ZM10.5 6.5a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagDollar(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7.182 2.94l-4.87 4.869a1 1 0 0 0-.006 1.408l4.37 4.45a1 1 0 0 0 1.414.012l4.96-4.868a1.5 1.5 0 0 0 .45-1.07V3.5a1 1 0 0 0-1-1H8.243a1.5 1.5 0 0 0-1.061.44m6.92 6.942A3 3 0 0 0 15 7.741V3.5A2.5 2.5 0 0 0 12.5 1H8.243a3 3 0 0 0-2.122.879l-4.87 4.87a2.5 2.5 0 0 0-.015 3.519l4.37 4.45a2.5 2.5 0 0 0 3.535.032zm-3.72-5.324a2.477 2.477 0 0 0-1.045-.38a1.95 1.95 0 0 0-1.616.57c-.54.54-.56 1.23-.437 1.767c.119.526.394 1.029.637 1.41c.254.4.46.732.549 1.033c.073.249.036.352-.072.46c-.072.072-.182.124-.417.066c-.253-.064-.569-.247-.855-.533a1.773 1.773 0 0 1-.469-.879c-.069-.329-.023-.57.034-.68a.75.75 0 1 0-1.33-.693c-.263.506-.289 1.125-.172 1.682c.07.332.196.67.382.988l-.253.252a.75.75 0 0 0 1.06 1.061l.222-.221a3.12 3.12 0 0 0 1.018.478c.6.15 1.301.079 1.84-.46c.612-.612.626-1.346.451-1.943c-.153-.522-.472-1.022-.694-1.37l-.029-.046c-.213-.334-.378-.664-.44-.937c-.06-.261-.002-.337.036-.374a.451.451 0 0 1 .392-.14c.172.02.445.126.777.457c.33.33.447.6.482.79a.832.832 0 0 1-.067.513a.75.75 0 1 0 1.362.627c.18-.392.281-.879.179-1.42a2.712 2.712 0 0 0-.446-1.036l.221-.221a.75.75 0 1 0-1.06-1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagRuble(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7.182 2.94l-4.87 4.869a1 1 0 0 0-.006 1.408l4.37 4.45a1 1 0 0 0 1.414.012l4.96-4.868a1.5 1.5 0 0 0 .45-1.07V3.5a1 1 0 0 0-1-1H8.243a1.5 1.5 0 0 0-1.061.44m6.92 6.942A3 3 0 0 0 15 7.741V3.5A2.5 2.5 0 0 0 12.5 1H8.243a3 3 0 0 0-2.122.879l-4.87 4.87a2.5 2.5 0 0 0-.015 3.519l4.37 4.45a2.5 2.5 0 0 0 3.535.032zm-3.908-5.844a.75.75 0 0 0-1.06 0L7.366 5.805l-.088-.088a.75.75 0 0 0-1.06 1.06l.087.09l-.353.353l-.088-.089a.75.75 0 0 0-1.061 1.061l.088.088l-.353.354a.75.75 0 0 0 1.06 1.06l.354-.353l1.06 1.06a.75.75 0 1 0 1.061-1.06l-1.06-1.06l.353-.354l.796.795a2.375 2.375 0 0 0 3.358-3.358zm-.972 3.624l-.795-.796L9.664 5.63l.796.795a.875.875 0 0 1-1.238 1.238Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tags(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m10.884 6.995l-4.12 4.12a.75.75 0 0 1-1.055.005L1.906 7.395a.75.75 0 0 1-.011-1.06l4.118-4.21a1.25 1.25 0 0 1 .894-.375H10.5a.75.75 0 0 1 .75.75v3.61c0 .332-.132.65-.366.885M4.94 1.077A2.75 2.75 0 0 1 6.907.25H10.5a2.25 2.25 0 0 1 2.25 2.25v.75h.75a2.25 2.25 0 0 1 2.25 2.25v3.61c0 .73-.29 1.43-.806 1.946l-4.12 4.12a2.25 2.25 0 0 1-3.165.016l-3.803-3.726a2.264 2.264 0 0 1-.286-.341L.856 8.466a2.25 2.25 0 0 1-.033-3.18l4.118-4.21Zm2.242 11.548a2.25 2.25 0 0 0 .642-.45L11.52 8.48a1.25 1.25 0 0 0 1.206-2.01c.015-.118.023-.237.023-.358V4.75h.75a.75.75 0 0 1 .75.75v3.61c0 .332-.132.65-.366.885l-4.12 4.12a.75.75 0 0 1-1.055.005zM8.75 5.5a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Target(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 13.5a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14m0-4.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5M8 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8m0-3a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TargetDart(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.293 0c.39 0 .707.317.707.707V2h1.293a.707.707 0 0 1 .5 1.207l-1.46 1.46A1.138 1.138 0 0 1 13.53 5h-1.47L8.53 8.53a.75.75 0 0 1-1.06-1.06L11 3.94V2.47c0-.301.12-.59.333-.804l1.46-1.46a.707.707 0 0 1 .5-.207ZM2.5 8a5.5 5.5 0 0 1 6.598-5.39a.75.75 0 0 0 .298-1.47A7 7 0 1 0 14.86 6.6a.75.75 0 0 0-1.47.299A5.5 5.5 0 1 1 2.5 8m5.364-2.496a.75.75 0 0 0-.08-1.498A4 4 0 1 0 11.988 8.3a.75.75 0 0 0-1.496-.111a2.5 2.5 0 1 1-2.63-2.686Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Terminal(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 3.5H4A1.5 1.5 0 0 0 2.5 5v6A1.5 1.5 0 0 0 4 12.5h8a1.5 1.5 0 0 0 1.5-1.5V5A1.5 1.5 0 0 0 12 3.5M4 2a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3zm.47 8.53a.75.75 0 0 1 0-1.06L5.94 8L4.47 6.53a.75.75 0 0 1 1.06-1.06l2 2a.75.75 0 0 1 0 1.06l-2 2a.75.75 0 0 1-1.06 0M8.75 9.5a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TerminalLine(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.22 11.22a.75.75 0 1 0 1.06 1.06l3.75-3.75a.75.75 0 0 0 0-1.06L2.28 3.72a.75.75 0 0 0-1.06 1.06L4.44 8zm13.03 1.28a.75.75 0 0 0 0-1.5h-6.5a.75.75 0 0 0 0 1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextAlignCenter(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.75 2a.75.75 0 0 0 0 1.5h10.5a.75.75 0 0 0 0-1.5zm0 7a.75.75 0 0 0 0 1.5h10.5a.75.75 0 0 0 0-1.5zm2 3.5a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5zM4 6.25a.75.75 0 0 1 .75-.75h6.5a.75.75 0 0 1 0 1.5h-6.5A.75.75 0 0 1 4 6.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextAlignJustify(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.75 2a.75.75 0 0 0 0 1.5h10.5a.75.75 0 0 0 0-1.5zM2 6.25a.75.75 0 0 1 .75-.75h10.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 6.25M2.75 9a.75.75 0 0 0 0 1.5h10.5a.75.75 0 0 0 0-1.5zm0 3.5a.75.75 0 0 0 0 1.5h10.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextAlignLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.75 2a.75.75 0 0 0 0 1.5h10.5a.75.75 0 0 0 0-1.5zm0 7a.75.75 0 0 0 0 1.5h10.5a.75.75 0 0 0 0-1.5zm0 3.5a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5zM2 6.25a.75.75 0 0 1 .75-.75h6.5a.75.75 0 0 1 0 1.5h-6.5A.75.75 0 0 1 2 6.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextAlignRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.75 2a.75.75 0 0 0 0 1.5h10.5a.75.75 0 0 0 0-1.5zM6 6.25a.75.75 0 0 1 .75-.75h6.5a.75.75 0 0 1 0 1.5h-6.5A.75.75 0 0 1 6 6.25M2.75 9a.75.75 0 0 0 0 1.5h10.5a.75.75 0 0 0 0-1.5zm4 3.5a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextIcon(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.279 2.544A.75.75 0 0 1 4 2h8a.75.75 0 0 1 .721.544l.5 1.75a.75.75 0 1 1-1.442.412L11.434 3.5H8.75l-.004 9H9.5a.75.75 0 0 1 0 1.5h-3a.75.75 0 0 1 0-1.5h.746l.004-9H4.566L4.22 4.706a.75.75 0 1 1-1.442-.412l.5-1.75Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextIndent(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.25 2H2.75a.75.75 0 0 0 0 1.5h10.5a.75.75 0 0 0 0-1.5m0 3.5h-5.5a.75.75 0 0 0 0 1.5h5.5a.75.75 0 0 0 0-1.5m0 3.5h-5.5a.75.75 0 0 0 0 1.5h5.5a.75.75 0 0 0 0-1.5m-10.5 3.5h10.5a.75.75 0 0 1 0 1.5H2.75a.75.75 0 0 1 0-1.5m.49-7a.74.74 0 0 1 .463.162l1.906 1.526a1.04 1.04 0 0 1 0 1.624l-1.906 1.526A.74.74 0 0 1 2.5 9.76V6.24a.74.74 0 0 1 .74-.74" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextOutdent(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.25 2H2.75a.75.75 0 0 0 0 1.5h10.5a.75.75 0 0 0 0-1.5m0 3.5h-5.5a.75.75 0 0 0 0 1.5h5.5a.75.75 0 0 0 0-1.5m0 3.5h-5.5a.75.75 0 0 0 0 1.5h5.5a.75.75 0 0 0 0-1.5m-10.5 3.5h10.5a.75.75 0 0 1 0 1.5H2.75a.75.75 0 0 1 0-1.5m2.01-7a.74.74 0 0 0-.463.162L2.39 7.188a1.04 1.04 0 0 0 0 1.624l1.907 1.526A.74.74 0 0 0 5.5 9.76V6.24a.74.74 0 0 0-.74-.74" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsDown(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m12 9l-2.94 5.041a1.932 1.932 0 0 1-3.56-1.378l.25-1.163l.321-1.5h-2.94a2 2 0 0 1-1.927-2.535l.879-3.162A4 4 0 0 1 6.404 1.4L11.5 2zM6.229 2.89l3.863.455l.379 5.3l-2.708 4.64a.432.432 0 0 1-.796-.308l.571-2.663l.389-1.814H3.13a.5.5 0 0 1-.482-.634l.879-3.162a2.5 2.5 0 0 1 2.7-1.814Zm7.023 5.663a.75.75 0 1 0 1.496-.106l-.5-7a.75.75 0 0 0-1.496.106z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsDownFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiThumbsDownFill0)"><path fill="currentColor" fill-rule="evenodd" d="M8.9 14.315a1.38 1.38 0 0 1-2.542-.984L7.07 10H3.131a2 2 0 0 1-1.927-2.535l.879-3.162A4 4 0 0 1 6.404 1.4L11.5 2l.5 7zm4.352-5.762a.75.75 0 1 0 1.496-.106l-.5-7a.75.75 0 0 0-1.496.106z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiThumbsDownFill0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsUp(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m4 7l2.94-5.041a1.932 1.932 0 0 1 3.56 1.378L10.25 4.5L9.93 6h2.94a2 2 0 0 1 1.927 2.535l-.879 3.162A4 4 0 0 1 9.596 14.6L4.5 14zm5.771 6.11l-3.863-.455l-.379-5.3l2.708-4.64a.432.432 0 0 1 .796.308l-.571 2.663L8.073 7.5h4.796a.5.5 0 0 1 .482.634l-.879 3.162a2.5 2.5 0 0 1-2.7 1.814ZM2.748 7.447a.75.75 0 1 0-1.496.106l.5 7a.75.75 0 0 0 1.496-.106z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsUpFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.1 1.685a1.38 1.38 0 0 1 2.542.984L8.93 6h3.939a2 2 0 0 1 1.927 2.535l-.879 3.162A4 4 0 0 1 9.596 14.6L4.5 14L4 7zM2.749 7.447a.75.75 0 1 0-1.496.106l.5 7a.75.75 0 0 0 1.496-.106z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsUpTwo(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiThumbsUp20)"><path fill="currentColor" fill-rule="evenodd" d="m13.345 2.978l-.135.609l-.33 1.5h1.158A2 2 0 0 1 16 7.472l-.259 1.321a4 4 0 0 1-4.322 3.21a4 4 0 0 1-4.286 2.767l-3.984-.473l-.48-6.352l2.813-4.564a1.854 1.854 0 0 1 3.427.826l1.452-2.418a1.627 1.627 0 0 1 2.984 1.189m-4.38 4.06l.215-.37l2.467-4.106a.127.127 0 0 1 .233.092l-.465 2.11l-.403 1.823h3.027a.5.5 0 0 1 .49.596l-.26 1.321a2.5 2.5 0 0 1-2.528 2.018l.219-1.09A2 2 0 0 0 10 7.037H8.964ZM4.55 12.952l2.759.328a2.5 2.5 0 0 0 2.745-1.99l.434-2.155a.5.5 0 0 0-.49-.599H6.476l.414-1.83l.515-2.275a.354.354 0 0 0-.646-.264L4.2 8.318zM.695 7.651a.75.75 0 0 1 .804.691l.478 6.353a.75.75 0 0 1-1.495.112L.004 8.455a.75.75 0 0 1 .691-.804" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiThumbsUp20"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Thunderbolt(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7.835 6l.76-2.027L9.336 2H5.5a.716.716 0 0 0-.664.45L2.513 8.257a.177.177 0 0 0 .164.243h4.965l-.732 2.013l-1.082 2.975a.382.382 0 0 0 .637.392l6.956-7.391A.29.29 0 0 0 13.21 6zm2.728-3l.235-.627A1.386 1.386 0 0 0 9.5.5h-4c-.906 0-1.72.552-2.057 1.393L1.12 7.7A1.677 1.677 0 0 0 2.677 10H5.5l-.545 1.5l-.537 1.475a1.882 1.882 0 0 0 3.14 1.933l6.956-7.391A1.79 1.79 0 0 0 13.21 4.5H10z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThunderboltFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.61 6.914l-7.632 8.08a1.614 1.614 0 0 1-2.69-1.66L5.5 10H2.677A1.677 1.677 0 0 1 1.12 7.7l2.323-5.807A2.216 2.216 0 0 1 5.5.5h4c.968 0 1.637.967 1.298 1.873L10 4.5h3.569a1.431 1.431 0 0 1 1.04 2.414Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ticket(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiTicket0)"><path fill="currentColor" fill-rule="evenodd" d="M11.091 2.318a.986.986 0 0 1 .106 1.485L11 4a.707.707 0 0 0 1 1l.197-.197a1.02 1.02 0 0 1 .288-.2a.986.986 0 0 1 1.197.306l1.243 1.658c.575.766.36 1.862-.461 2.355a16.158 16.158 0 0 0-5.542 5.542a1.616 1.616 0 0 1-2.355.461l-1.658-1.243a.986.986 0 0 1-.106-1.485L5 12a.707.707 0 1 0-1-1l-.197.197a.968.968 0 0 1-.288.2a.986.986 0 0 1-1.197-.306L1.075 9.433a1.616 1.616 0 0 1 .461-2.355a16.158 16.158 0 0 0 5.542-5.542a1.616 1.616 0 0 1 2.355-.461zm-2.558-.043l1.201.9a2.208 2.208 0 0 0 3.09 3.09l.901 1.202a.116.116 0 0 1-.033.169a17.658 17.658 0 0 0-6.057 6.056a.116.116 0 0 1-.168.033l-1.201-.9a2.208 2.208 0 0 0-3.09-3.09l-.901-1.202a.116.116 0 0 1 .033-.168a17.658 17.658 0 0 0 6.056-6.057a.116.116 0 0 1 .17-.033ZM7.53 6.47a.75.75 0 1 0-1.06 1.06l2 2a.75.75 0 0 0 1.06-1.06z" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiTicket0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Timeline(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.25 1.25v13.5a.75.75 0 0 0 1.5 0V1.25a.75.75 0 0 0-1.5 0M2.5 2H6v1.5H2.5a1 1 0 0 0-1 1v1.75a1 1 0 0 0 1 1H6v1.5h-.5a1 1 0 0 0-1 1v1.75a1 1 0 0 0 1 1H6V14h-.5A2.5 2.5 0 0 1 3 11.5V9.75c0-.356.074-.694.208-1H2.5A2.5 2.5 0 0 1 0 6.25V4.5A2.5 2.5 0 0 1 2.5 2m8 5.25H10v1.5h3.5a1 1 0 0 1 1 1v1.75a1 1 0 0 1-1 1H10V14h3.5a2.5 2.5 0 0 0 2.5-2.5V9.75a2.5 2.5 0 0 0-2.5-2.5h-.708c.134-.306.208-.644.208-1V4.5A2.5 2.5 0 0 0 10.5 2H10v1.5h.5a1 1 0 0 1 1 1v1.75a1 1 0 0 1-1 1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Timestamps(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 12.75a2.75 2.75 0 0 0-2-2.646V8.75h2.25a.75.75 0 0 0 0-1.5h-6V5.896a2.752 2.752 0 0 0 2-2.646a2.75 2.75 0 1 0-3.5 2.646V7.25h-6a.75.75 0 0 0 0 1.5H3.5v1.354a2.752 2.752 0 0 0-2 2.646A2.75 2.75 0 1 0 5 10.104V8.75h6v1.354a2.751 2.751 0 1 0 3.5 2.646M11.75 14a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5M8 4.5A1.25 1.25 0 1 0 8 2a1.25 1.25 0 0 0 0 2.5M4.25 14a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrafficLight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 12.5v-9A1.5 1.5 0 0 1 6 2h4a1.5 1.5 0 0 1 1.5 1.5v9A1.5 1.5 0 0 1 10 14H6a1.5 1.5 0 0 1-1.5-1.5m10.374-7.834L13 7.477v.773h1.25a.75.75 0 0 1 .624 1.166L13 12.227v.273a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3v-.273L1.126 9.416A.75.75 0 0 1 1.75 8.25H3v-.773L1.126 4.666A.75.75 0 0 1 1.75 3.5H3a3 3 0 0 1 3-3h4a3 3 0 0 1 3 3h1.25a.75.75 0 0 1 .624 1.166M8 7.25a1.75 1.75 0 1 0 0-3.5a1.75 1.75 0 0 0 0 3.5m0 5a1.75 1.75 0 1 0 0-3.5a1.75 1.75 0 0 0 0 3.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrashBin(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9 2H7a.5.5 0 0 0-.5.5V3h3v-.5A.5.5 0 0 0 9 2m2 1v-.5a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2V3H2.251a.75.75 0 0 0 0 1.5h.312l.317 7.625A3 3 0 0 0 5.878 15h4.245a3 3 0 0 0 2.997-2.875l.318-7.625h.312a.75.75 0 0 0 0-1.5zm.936 1.5H4.064l.315 7.562A1.5 1.5 0 0 0 5.878 13.5h4.245a1.5 1.5 0 0 0 1.498-1.438zm-6.186 2v5a.75.75 0 0 0 1.5 0v-5a.75.75 0 0 0-1.5 0m3.75-.75a.75.75 0 0 1 .75.75v5a.75.75 0 0 1-1.5 0v-5a.75.75 0 0 1 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tray(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 8.5h-2.388a3.422 3.422 0 0 1-6.224 0H2.5V11A1.5 1.5 0 0 0 4 12.5h8a1.5 1.5 0 0 0 1.5-1.5zm-2.204-4.757L13.251 7H10l-.136.545a1.921 1.921 0 0 1-3.728 0L6 7H2.75l1.954-3.257a.5.5 0 0 1 .428-.243h5.736a.5.5 0 0 1 .428.243M15 8.5v-.67a3 3 0 0 0-.428-1.543l-1.99-3.316A2 2 0 0 0 10.869 2H5.132a2 2 0 0 0-1.715.971l-1.99 3.316A3 3 0 0 0 1 7.831V11a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleDown(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.134 13.005L2.217 4.5A1 1 0 0 1 3.083 3h9.834a1 1 0 0 1 .866 1.5l-4.917 8.505a1 1 0 0 1-1.732 0m3.03.751c-.962 1.665-3.366 1.665-4.329 0L.918 5.251C-.045 3.584 1.158 1.5 3.083 1.5h9.834c1.925 0 3.128 2.084 2.164 3.751z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleDownFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.164 13.756c-.962 1.665-3.366 1.665-4.329 0L.918 5.251C-.045 3.584 1.158 1.5 3.083 1.5h9.834c1.925 0 3.128 2.084 2.164 3.751z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleExclamation(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.134 2.994L2.217 11.5a1 1 0 0 0 .866 1.5h9.834a1 1 0 0 0 .866-1.5L8.866 2.993a1 1 0 0 0-1.732 0Zm3.03-.75c-.962-1.665-3.366-1.665-4.328 0L.919 10.749c-.964 1.666.239 3.751 2.164 3.751h9.834c1.925 0 3.128-2.085 2.164-3.751zM8 5a.75.75 0 0 1 .75.75v2a.75.75 0 0 1-1.5 0v-2A.75.75 0 0 1 8 5m1 5.75a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleExclamationFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.836 2.244c.962-1.665 3.366-1.665 4.328 0l4.917 8.505c.964 1.666-.239 3.751-2.164 3.751H3.083c-1.925 0-3.128-2.085-2.164-3.751zM8 5a.75.75 0 0 1 .75.75v2a.75.75 0 1 1-1.5 0v-2A.75.75 0 0 1 8 5m1 5.75a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleLeft(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.995 7.134L11.5 2.217a1 1 0 0 1 1.5.866v9.834a1 1 0 0 1-1.5.866L2.995 8.866a1 1 0 0 1 0-1.732m-.751 3.03c-1.665-.962-1.665-3.366 0-4.329L10.749.918c1.666-.963 3.751.24 3.751 2.165v9.834c0 1.925-2.085 3.128-3.751 2.164z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleLeftFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.244 10.164c-1.665-.962-1.665-3.366 0-4.328L10.749.919c1.666-.964 3.751.239 3.751 2.164v9.834c0 1.925-2.085 3.128-3.751 2.164z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleRight(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.005 7.134L4.5 2.217a1 1 0 0 0-1.5.866v9.834a1 1 0 0 0 1.5.866l8.505-4.917a1 1 0 0 0 0-1.732m.751 3.03c1.665-.962 1.665-3.366 0-4.329L5.251.918C3.584-.045 1.5 1.158 1.5 3.083v9.834c0 1.925 2.084 3.128 3.751 2.164z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleRightFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.756 10.164c1.665-.962 1.665-3.366 0-4.328L5.251.919C3.584-.045 1.5 1.158 1.5 3.083v9.834c0 1.925 2.084 3.128 3.751 2.164z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleThunderbolt(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.134 2.994L2.217 11.5a1 1 0 0 0 .866 1.5h9.834a1 1 0 0 0 .866-1.5L8.866 2.993a1 1 0 0 0-1.732 0Zm3.03-.75c-.962-1.665-3.366-1.665-4.329 0L.918 10.749c-.963 1.666.24 3.751 2.165 3.751h9.834c1.925 0 3.128-2.085 2.164-3.751zM8.671 6.585a.75.75 0 0 0-1.342-.67l-1.25 2.5a.75.75 0 0 0 .67 1.085h1.287l-.707 1.415a.75.75 0 1 0 1.342.67l1.25-2.5A.75.75 0 0 0 9.25 8H7.963z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleUp(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.134 2.995L2.217 11.5a1 1 0 0 0 .866 1.5h9.834a1 1 0 0 0 .866-1.5L8.866 2.995a1 1 0 0 0-1.732 0m3.03-.751c-.962-1.665-3.366-1.665-4.329 0L.918 10.749c-.963 1.666.24 3.751 2.165 3.751h9.834c1.925 0 3.128-2.085 2.164-3.751z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleUpFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.164 2.244c-.962-1.665-3.366-1.665-4.329 0L.918 10.749c-.963 1.666.24 3.751 2.165 3.751h9.834c1.925 0 3.128-2.085 2.164-3.751z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trolley(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.988 11A1.5 1.5 0 1 0 7 11zM10 12.476v.024a3 3 0 1 1-4.473-2.614l-1.955-7.33A.75.75 0 0 0 2.848 2H1.25a.75.75 0 0 1 0-1.5h1.598a2.25 2.25 0 0 1 2.174 1.67l.542 2.035a2.251 2.251 0 0 1 .69-.317l4.602-1.231A2.25 2.25 0 0 1 13.61 4.25l1.228 4.607a2.25 2.25 0 0 1-1.595 2.753zm1.243-8.37a.75.75 0 0 1 .919.531l1.227 4.607a.75.75 0 0 1-.531.918l-3.244.865A2.999 2.999 0 0 0 7 9.5h-.024l-.865-3.245a.75.75 0 0 1 .53-.918l1.595-.427l.642 2.291a.75.75 0 0 0 1.444-.403l-.637-2.275z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tv(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 3.5h10A1.5 1.5 0 0 1 14.5 5v5a1.5 1.5 0 0 1-1.5 1.5H3A1.5 1.5 0 0 1 1.5 10V5A1.5 1.5 0 0 1 3 3.5m-.21 9.493A3 3 0 0 1 0 10V5a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v5a3 3 0 0 1-2.79 2.993l.46.922a.75.75 0 1 1-1.34.67L11.536 13H4.463l-.793 1.585a.75.75 0 1 1-1.342-.67l.461-.922Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TvRetro(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.086.782a.75.75 0 0 0-1.172.937L5.94 3H4a3 3 0 0 0-3 3v5a3 3 0 0 0 2.79 2.993l-.21.421a.75.75 0 0 0 1.34.671l.5-1A.762.762 0 0 0 5.458 14h5.086c.01.029.022.057.036.085l.5 1a.75.75 0 0 0 1.342-.67l-.211-.422A3 3 0 0 0 15 11V6a3 3 0 0 0-3-3h-1.94l1.026-1.281A.75.75 0 0 0 9.914.782L8.14 3h-.28zM7.494 4.5H12A1.5 1.5 0 0 1 13.5 6v5a1.5 1.5 0 0 1-1.5 1.5H4A1.5 1.5 0 0 1 2.5 11V6A1.5 1.5 0 0 1 4 4.5zM5.5 7.5v2H8v-2zM5 6a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h3.5a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1zm7 .75a.75.75 0 0 0-1.5 0v3.5a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Umbrella(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.183 11.595c-.05.02-.34.112-1.113-.187c-.953-.368-2.258-1.218-3.76-2.719c-1.5-1.5-2.35-2.806-2.718-3.759c-.299-.774-.206-1.063-.187-1.113c1.57-1.535 5.177-2.013 7.484.294c2.32 2.32 1.914 5.83.294 7.484M5.79 10.27c2.947 2.747 5.427 3.411 6.427 2.41c2.063-2.061 2.617-6.141.236-9.073l.577-.578a.75.75 0 1 0-1.06-1.06l-.579.578C8.448.178 4.286.814 2.318 2.782c-1 1-.336 3.48 2.41 6.428L2.47 11.47a2.164 2.164 0 1 0 3.06 3.06l.25-.25a.75.75 0 1 0-1.06-1.06l-.25.25a.664.664 0 1 1-.94-.94l2.26-2.259Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Underline(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 2.75a.75.75 0 0 0-1.5 0V7a4.5 4.5 0 0 0 9 0V2.75a.75.75 0 0 0-1.5 0V7a3 3 0 0 1-6 0zm-.75 9.75a.75.75 0 0 0 0 1.5h7.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vault(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 14.25v-.75h7v.75a.75.75 0 0 0 1.5 0v-.791a3 3 0 0 0 2.5-2.959v-6a3 3 0 0 0-3-3h-9a3 3 0 0 0-3 3v6A3 3 0 0 0 3 13.459v.791a.75.75 0 0 0 1.5 0M3.5 12h9a1.5 1.5 0 0 0 1.5-1.5v-6A1.5 1.5 0 0 0 12.5 3h-9A1.5 1.5 0 0 0 2 4.5v6A1.5 1.5 0 0 0 3.5 12M8 8.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2m2.384-.246a2.498 2.498 0 0 0-.402-2.278l.66-.942a.75.75 0 0 0-1.229-.86l-.66.942a2.499 2.499 0 0 0-2.277.402l-.942-.66a.75.75 0 0 0-.86 1.228l.942.66a2.499 2.499 0 0 0 .402 2.277l-.66.943a.75.75 0 0 0 1.228.86l.66-.942a2.499 2.499 0 0 0 2.277-.402l.943.66a.75.75 0 1 0 .86-1.228z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VectorCircle(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 5.5a2.74 2.74 0 0 1-1.444-.409a2.764 2.764 0 0 1-.914-.925c-.6.37-1.106.877-1.476 1.476a2.749 2.749 0 0 1 0 4.716c.37.6.877 1.106 1.476 1.476a2.749 2.749 0 0 1 4.716 0c.6-.37 1.106-.877 1.476-1.476l-.028-.017a2.748 2.748 0 0 1 .028-4.699a4.524 4.524 0 0 0-1.476-1.476l-.017.028A2.748 2.748 0 0 1 8 5.5m2.749-2.835a2.75 2.75 0 0 0-5.498 0a6.026 6.026 0 0 0-2.586 2.586a2.75 2.75 0 0 0 0 5.498a6.026 6.026 0 0 0 2.586 2.586a2.749 2.749 0 0 0 5.09 1.359c.245-.398.393-.862.408-1.36a6.027 6.027 0 0 0 2.586-2.585A2.749 2.749 0 0 0 16 8a2.75 2.75 0 0 0-2.665-2.749a6.027 6.027 0 0 0-2.586-2.586M9.25 2.75a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0M4 8a1.25 1.25 0 1 1-2.5 0A1.25 1.25 0 0 1 4 8m5.25 5.25a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0m4-4a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VectorSquare(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.091 5.194c.132-.214.236-.447.305-.694h3.208c.259.916.98 1.637 1.896 1.896v3.208A2.756 2.756 0 0 0 9.604 11.5H6.396A2.756 2.756 0 0 0 4.5 9.604V6.396a2.756 2.756 0 0 0 1.591-1.202M13 6.396v3.208a2.748 2.748 0 0 1 1.591 4.09A2.751 2.751 0 0 1 9.603 13H6.397a2.748 2.748 0 0 1-2.646 2A2.75 2.75 0 0 1 3 9.604V6.396A2.751 2.751 0 1 1 6.396 3h3.208A2.751 2.751 0 1 1 13 6.396M3.75 2.5a1.25 1.25 0 1 1 0 2.5a1.25 1.25 0 0 1 0-2.5M5 12.25a1.25 1.25 0 1 0-2.5 0a1.25 1.25 0 0 0 2.5 0M12.25 11a1.25 1.25 0 1 1 0 2.5a1.25 1.25 0 0 1 0-2.5m1.25-7.25a1.25 1.25 0 1 0-2.5 0a1.25 1.25 0 0 0 2.5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Video(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.5h5.5A1.5 1.5 0 0 1 10 6v4a1.5 1.5 0 0 1-1.5 1.5H3A1.5 1.5 0 0 1 1.5 10V6A1.5 1.5 0 0 1 3 4.5m8.452 6.037A3 3 0 0 1 8.5 13H3a3 3 0 0 1-3-3V6a3 3 0 0 1 3-3h5.5a3 3 0 0 1 2.952 2.463l1.554-1.11A1.893 1.893 0 0 1 16 5.893v4.214a1.893 1.893 0 0 1-2.994 1.54zm.048-1.809l2.378 1.699a.393.393 0 0 0 .622-.32V5.893a.393.393 0 0 0-.622-.32L11.5 7.272z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Volume(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiVolume0)"><path fill="currentColor" fill-rule="evenodd" d="M5.06 9.94A1.5 1.5 0 0 0 4 9.5H2a.5.5 0 0 1-.5-.5V7a.5.5 0 0 1 .5-.5h2a1.5 1.5 0 0 0 1.06-.44l2.483-2.482a.268.268 0 0 1 .457.19v8.464a.268.268 0 0 1-.457.19zM2 5h2l2.482-2.482A1.768 1.768 0 0 1 9.5 3.768v8.464a1.768 1.768 0 0 1-3.018 1.25L4 11H2a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2m12.162 8.103c-.265.319-.743.316-1.036.024c-.292-.293-.288-.766-.031-1.09A6.473 6.473 0 0 0 14.5 8a6.473 6.473 0 0 0-1.405-4.037c-.257-.324-.261-.797.031-1.09c.293-.292.771-.294 1.036.024A7.967 7.967 0 0 1 16 8c0 1.94-.69 3.718-1.838 5.103m-2.138-2.135c-.246.333-.726.33-1.019.037c-.293-.293-.284-.764-.06-1.113A3.484 3.484 0 0 0 11.5 8c0-.697-.204-1.347-.555-1.892c-.224-.348-.233-.82.06-1.113c.293-.293.773-.296 1.02.037C12.637 5.862 13 6.89 13 8a4.977 4.977 0 0 1-.976 2.968" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiVolume0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeFill(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#gravityUiVolumeFill0)"><path fill="currentColor" fill-rule="evenodd" d="M2 5h2l2.482-2.482A1.768 1.768 0 0 1 9.5 3.768v8.464a1.768 1.768 0 0 1-3.018 1.25L4 11H2a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2m12.162 8.103c-.265.319-.743.316-1.036.024c-.292-.293-.288-.766-.031-1.09A6.473 6.473 0 0 0 14.5 8a6.473 6.473 0 0 0-1.405-4.037c-.257-.324-.261-.797.031-1.09c.293-.292.771-.294 1.036.024A7.967 7.967 0 0 1 16 8c0 1.94-.69 3.718-1.838 5.103m-2.138-2.135c-.246.333-.726.33-1.019.037c-.293-.293-.284-.764-.06-1.113A3.484 3.484 0 0 0 11.5 8c0-.697-.204-1.347-.555-1.892c-.224-.348-.233-.82.06-1.113c.293-.293.773-.296 1.02.037C12.637 5.862 13 6.89 13 8a4.977 4.977 0 0 1-.976 2.968" clip-rule="evenodd"/></g><defs><clipPath id="gravityUiVolumeFill0"><path fill="currentColor" d="M0 0h16v16H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeLow(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.06 9.94A1.5 1.5 0 0 0 4 9.5H2a.5.5 0 0 1-.5-.5V7a.5.5 0 0 1 .5-.5h2a1.5 1.5 0 0 0 1.06-.44l2.483-2.482a.268.268 0 0 1 .457.19v8.464a.268.268 0 0 1-.457.19zM2 5h2l2.482-2.482A1.768 1.768 0 0 1 9.5 3.768v8.464a1.768 1.768 0 0 1-3.018 1.25L4 11H2a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2m10.024 5.968c-.246.333-.726.33-1.019.037c-.293-.293-.284-.764-.06-1.113A3.484 3.484 0 0 0 11.5 8c0-.697-.204-1.347-.555-1.892c-.224-.348-.233-.82.06-1.113c.293-.293.773-.296 1.02.037C12.637 5.862 13 6.89 13 8a4.977 4.977 0 0 1-.976 2.968" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeOff(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.06 9.94A1.5 1.5 0 0 0 4 9.5H2a.5.5 0 0 1-.5-.5V7a.5.5 0 0 1 .5-.5h2a1.5 1.5 0 0 0 1.06-.44l2.483-2.482a.268.268 0 0 1 .457.19v8.464a.268.268 0 0 1-.457.19zM2 5h2l2.482-2.482A1.768 1.768 0 0 1 9.5 3.768v8.464a1.768 1.768 0 0 1-3.018 1.25L4 11H2a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeXmark(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.06 9.94A1.5 1.5 0 0 0 4 9.5H2a.5.5 0 0 1-.5-.5V7a.5.5 0 0 1 .5-.5h2a1.5 1.5 0 0 0 1.06-.44l2.483-2.482a.268.268 0 0 1 .457.19v8.464a.268.268 0 0 1-.457.19zM2 5h2l2.482-2.482A1.768 1.768 0 0 1 9.5 3.768v8.464a1.768 1.768 0 0 1-3.018 1.25L4 11H2a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2m10.28.72a.75.75 0 1 0-1.06 1.06L12.44 8l-1.22 1.22a.75.75 0 1 0 1.06 1.06l1.22-1.22l1.22 1.22a.75.75 0 1 0 1.06-1.06L14.56 8l1.22-1.22a.75.75 0 0 0-1.06-1.06L13.5 6.94z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WeightHanging(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.25 4.886A2.497 2.497 0 0 1 5.5 2.5a2.5 2.5 0 1 1 3.25 2.386V6h2.904a2.5 2.5 0 0 1 2.389 1.765l1.36 4.423A2.173 2.173 0 0 1 13.328 15H2.673a2.173 2.173 0 0 1-2.077-2.812l.252-.82l1.109-3.603A2.5 2.5 0 0 1 4.347 6H7.25zM9 2.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-4.654 5a1 1 0 0 0-.955.706L2.03 12.629a.673.673 0 0 0 .643.871h10.654a.673.673 0 0 0 .643-.871l-1.36-4.423a1 1 0 0 0-.956-.706z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wrench(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m8.581 9.298l.776.143A3.5 3.5 0 0 0 13.5 6c0-.118-.01-.236-.027-.352l-1.39 1.39a1.575 1.575 0 0 1-1.114.462A2.47 2.47 0 0 1 8.5 5.03c0-.417.166-.817.461-1.112l1.39-1.39A2.402 2.402 0 0 0 10 2.5a3.5 3.5 0 0 0-3.441 4.143l.143.776l-3.813 3.813a1.329 1.329 0 0 0 1.879 1.879zm3.817-6.787a.795.795 0 0 0-.411-1.018C11.87 1.432 11.014 1 10 1a5 5 0 0 0-4.916 5.916l-3.256 3.256a2.828 2.828 0 1 0 4 4l3.256-3.256A5 5 0 0 0 15 6c0-1.014-.432-1.87-.493-1.987l-.014-.027a.795.795 0 0 0-1.273-.207l-2.198 2.2a.074.074 0 0 1-.053.021a.97.97 0 0 1-.969-.97c0-.019.008-.038.022-.052L12.22 2.78a.792.792 0 0 0 .178-.27Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xmark(children ...ElementRenderer) *GravityUiIcon {
	return &GravityUiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.47 3.47a.75.75 0 0 1 1.06 0L8 6.94l3.47-3.47a.75.75 0 1 1 1.06 1.06L9.06 8l3.47 3.47a.75.75 0 1 1-1.06 1.06L8 9.06l-3.47 3.47a.75.75 0 0 1-1.06-1.06L6.94 8L3.47 4.53a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
