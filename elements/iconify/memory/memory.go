package memory

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = ""
	hAttr          = "1em"
	viewbox        = "0 0 22 22"
	fill           = "currentColor"
)

type MemoryIcon struct {
	*SVGSVGElement
}

type MemoryIconFn func(children ...ElementRenderer) *MemoryIcon

var IconLookup = map[string]MemoryIconFn{
	"account":                                Account,
	"accountBox":                             AccountBox,
	"alert":                                  Alert,
	"alertBox":                               AlertBox,
	"alertBoxFill":                           AlertBoxFill,
	"alertCircle":                            AlertCircle,
	"alertCircleFill":                        AlertCircleFill,
	"alertHexagon":                           AlertHexagon,
	"alertRhombus":                           AlertRhombus,
	"alertRhombusFill":                       AlertRhombusFill,
	"alignHorizontalCenter":                  AlignHorizontalCenter,
	"alignHorizontalDistribute":              AlignHorizontalDistribute,
	"alignHorizontalLeft":                    AlignHorizontalLeft,
	"alignHorizontalRight":                   AlignHorizontalRight,
	"alignVerticalBottom":                    AlignVerticalBottom,
	"alignVerticalCenter":                    AlignVerticalCenter,
	"alignVerticalDistribute":                AlignVerticalDistribute,
	"alignVerticalTop":                       AlignVerticalTop,
	"alphaA":                                 AlphaA,
	"alphaAfill":                             AlphaAfill,
	"alphaB":                                 AlphaB,
	"alphaBfill":                             AlphaBfill,
	"alphaC":                                 AlphaC,
	"alphaCfill":                             AlphaCfill,
	"alphaD":                                 AlphaD,
	"alphaDfill":                             AlphaDfill,
	"alphaE":                                 AlphaE,
	"alphaEfill":                             AlphaEfill,
	"alphaF":                                 AlphaF,
	"alphaFfill":                             AlphaFfill,
	"alphaG":                                 AlphaG,
	"alphaGfill":                             AlphaGfill,
	"alphaH":                                 AlphaH,
	"alphaHfill":                             AlphaHfill,
	"alphaI":                                 AlphaI,
	"alphaIfill":                             AlphaIfill,
	"alphaJ":                                 AlphaJ,
	"alphaJfill":                             AlphaJfill,
	"alphaK":                                 AlphaK,
	"alphaKfill":                             AlphaKfill,
	"alphaL":                                 AlphaL,
	"alphaLfill":                             AlphaLfill,
	"alphaM":                                 AlphaM,
	"alphaMfill":                             AlphaMfill,
	"alphaN":                                 AlphaN,
	"alphaNfill":                             AlphaNfill,
	"alphaO":                                 AlphaO,
	"alphaOfill":                             AlphaOfill,
	"alphaP":                                 AlphaP,
	"alphaPfill":                             AlphaPfill,
	"alphaQ":                                 AlphaQ,
	"alphaQfill":                             AlphaQfill,
	"alphaR":                                 AlphaR,
	"alphaRfill":                             AlphaRfill,
	"alphaS":                                 AlphaS,
	"alphaSfill":                             AlphaSfill,
	"alphaT":                                 AlphaT,
	"alphaTfill":                             AlphaTfill,
	"alphaU":                                 AlphaU,
	"alphaUfill":                             AlphaUfill,
	"alphaV":                                 AlphaV,
	"alphaVfill":                             AlphaVfill,
	"alphaW":                                 AlphaW,
	"alphaWfill":                             AlphaWfill,
	"alphaX":                                 AlphaX,
	"alphaXfill":                             AlphaXfill,
	"alphaY":                                 AlphaY,
	"alphaYfill":                             AlphaYfill,
	"alphaZ":                                 AlphaZ,
	"alphaZfill":                             AlphaZfill,
	"application":                            Application,
	"applicationCode":                        ApplicationCode,
	"apps":                                   Apps,
	"appsBox":                                AppsBox,
	"appsBoxFill":                            AppsBoxFill,
	"archive":                                Archive,
	"arrow":                                  Arrow,
	"arrowBottomLeft":                        ArrowBottomLeft,
	"arrowBottomLeftCircle":                  ArrowBottomLeftCircle,
	"arrowBottomRight":                       ArrowBottomRight,
	"arrowBottomRightCircle":                 ArrowBottomRightCircle,
	"arrowDown":                              ArrowDown,
	"arrowDownBold":                          ArrowDownBold,
	"arrowDownBox":                           ArrowDownBox,
	"arrowDownCircle":                        ArrowDownCircle,
	"arrowDownLeft":                          ArrowDownLeft,
	"arrowDownLeftBox":                       ArrowDownLeftBox,
	"arrowDownRight":                         ArrowDownRight,
	"arrowDownRightBox":                      ArrowDownRightBox,
	"arrowLeft":                              ArrowLeft,
	"arrowLeftBold":                          ArrowLeftBold,
	"arrowLeftBox":                           ArrowLeftBox,
	"arrowLeftCircle":                        ArrowLeftCircle,
	"arrowLeftDown":                          ArrowLeftDown,
	"arrowLeftRight":                         ArrowLeftRight,
	"arrowLeftUp":                            ArrowLeftUp,
	"arrowRight":                             ArrowRight,
	"arrowRightBold":                         ArrowRightBold,
	"arrowRightBox":                          ArrowRightBox,
	"arrowRightCircle":                       ArrowRightCircle,
	"arrowRightDown":                         ArrowRightDown,
	"arrowRightUp":                           ArrowRightUp,
	"arrowTopLeft":                           ArrowTopLeft,
	"arrowTopLeftCircle":                     ArrowTopLeftCircle,
	"arrowTopRight":                          ArrowTopRight,
	"arrowTopRightCircle":                    ArrowTopRightCircle,
	"arrowUp":                                ArrowUp,
	"arrowUpBold":                            ArrowUpBold,
	"arrowUpBox":                             ArrowUpBox,
	"arrowUpCircle":                          ArrowUpCircle,
	"arrowUpDown":                            ArrowUpDown,
	"arrowUpLeft":                            ArrowUpLeft,
	"arrowUpLeftBox":                         ArrowUpLeftBox,
	"arrowUpRight":                           ArrowUpRight,
	"arrowUpRightBox":                        ArrowUpRightBox,
	"aspectRatio":                            AspectRatio,
	"axe":                                    Axe,
	"bagPersonal":                            BagPersonal,
	"bagPersonalFill":                        BagPersonalFill,
	"bank":                                   Bank,
	"batteryFifty":                           BatteryFifty,
	"batteryOneHundred":                      BatteryOneHundred,
	"batterySeventyFive":                     BatterySeventyFive,
	"batteryTwentyFive":                      BatteryTwentyFive,
	"batteryZero":                            BatteryZero,
	"battleAxe":                              BattleAxe,
	"beer":                                   Beer,
	"bell":                                   Bell,
	"blood":                                  Blood,
	"book":                                   Book,
	"bookmark":                               Bookmark,
	"borderBottom":                           BorderBottom,
	"borderBottomLeft":                       BorderBottomLeft,
	"borderBottomLeftRight":                  BorderBottomLeftRight,
	"borderBottomRight":                      BorderBottomRight,
	"borderInside":                           BorderInside,
	"borderLeft":                             BorderLeft,
	"borderLeftRight":                        BorderLeftRight,
	"borderNone":                             BorderNone,
	"borderOutside":                          BorderOutside,
	"borderRight":                            BorderRight,
	"borderTop":                              BorderTop,
	"borderTopBottom":                        BorderTopBottom,
	"borderTopLeft":                          BorderTopLeft,
	"borderTopLeftBottom":                    BorderTopLeftBottom,
	"borderTopLeftRight":                     BorderTopLeftRight,
	"borderTopRight":                         BorderTopRight,
	"borderTopRightBottom":                   BorderTopRightBottom,
	"bow":                                    Bow,
	"bowArrow":                               BowArrow,
	"box":                                    Box,
	"boxLightDashedDownLeft":                 BoxLightDashedDownLeft,
	"boxLightDashedDownRight":                BoxLightDashedDownRight,
	"boxLightDashedHorizontal":               BoxLightDashedHorizontal,
	"boxLightDashedUpLeft":                   BoxLightDashedUpLeft,
	"boxLightDashedUpRight":                  BoxLightDashedUpRight,
	"boxLightDashedVertical":                 BoxLightDashedVertical,
	"boxLightDashedVerticalHorizontal":       BoxLightDashedVerticalHorizontal,
	"boxLightDoubleDownLeft":                 BoxLightDoubleDownLeft,
	"boxLightDoubleDownRight":                BoxLightDoubleDownRight,
	"boxLightDoubleHorizontal":               BoxLightDoubleHorizontal,
	"boxLightDoubleHorizontalDown":           BoxLightDoubleHorizontalDown,
	"boxLightDoubleHorizontalLightDown":      BoxLightDoubleHorizontalLightDown,
	"boxLightDoubleHorizontalLightUp":        BoxLightDoubleHorizontalLightUp,
	"boxLightDoubleHorizontalUp":             BoxLightDoubleHorizontalUp,
	"boxLightDoubleRoundDownLeft":            BoxLightDoubleRoundDownLeft,
	"boxLightDoubleRoundDownRight":           BoxLightDoubleRoundDownRight,
	"boxLightDoubleRoundUpLeft":              BoxLightDoubleRoundUpLeft,
	"boxLightDoubleRoundUpRight":             BoxLightDoubleRoundUpRight,
	"boxLightDoubleUpLeft":                   BoxLightDoubleUpLeft,
	"boxLightDoubleUpRight":                  BoxLightDoubleUpRight,
	"boxLightDoubleVertical":                 BoxLightDoubleVertical,
	"boxLightDoubleVerticalHorizontal":       BoxLightDoubleVerticalHorizontal,
	"boxLightDoubleVerticalLeft":             BoxLightDoubleVerticalLeft,
	"boxLightDoubleVerticalLightLeft":        BoxLightDoubleVerticalLightLeft,
	"boxLightDoubleVerticalLightRight":       BoxLightDoubleVerticalLightRight,
	"boxLightDoubleVerticalRight":            BoxLightDoubleVerticalRight,
	"boxLightDownLeft":                       BoxLightDownLeft,
	"boxLightDownLeftCircle":                 BoxLightDownLeftCircle,
	"boxLightDownLeftStipple":                BoxLightDownLeftStipple,
	"boxLightDownLeftStippleInner":           BoxLightDownLeftStippleInner,
	"boxLightDownLeftStippleOuter":           BoxLightDownLeftStippleOuter,
	"boxLightDownRight":                      BoxLightDownRight,
	"boxLightDownRightCircle":                BoxLightDownRightCircle,
	"boxLightDownRightStipple":               BoxLightDownRightStipple,
	"boxLightDownRightStippleInner":          BoxLightDownRightStippleInner,
	"boxLightDownRightStippleOuter":          BoxLightDownRightStippleOuter,
	"boxLightFoldDownLeft":                   BoxLightFoldDownLeft,
	"boxLightFoldDownRight":                  BoxLightFoldDownRight,
	"boxLightFoldUpLeft":                     BoxLightFoldUpLeft,
	"boxLightFoldUpRight":                    BoxLightFoldUpRight,
	"boxLightHorizontal":                     BoxLightHorizontal,
	"boxLightHorizontalCircle":               BoxLightHorizontalCircle,
	"boxLightHorizontalDown":                 BoxLightHorizontalDown,
	"boxLightHorizontalDownStipple":          BoxLightHorizontalDownStipple,
	"boxLightHorizontalDownStippleDown":      BoxLightHorizontalDownStippleDown,
	"boxLightHorizontalDownStippleDownLeft":  BoxLightHorizontalDownStippleDownLeft,
	"boxLightHorizontalDownStippleDownRight": BoxLightHorizontalDownStippleDownRight,
	"boxLightHorizontalMenuDown":             BoxLightHorizontalMenuDown,
	"boxLightHorizontalMenuLeft":             BoxLightHorizontalMenuLeft,
	"boxLightHorizontalMenuRight":            BoxLightHorizontalMenuRight,
	"boxLightHorizontalMenuUp":               BoxLightHorizontalMenuUp,
	"boxLightHorizontalStipple":              BoxLightHorizontalStipple,
	"boxLightHorizontalStippleDown":          BoxLightHorizontalStippleDown,
	"boxLightHorizontalStippleUp":            BoxLightHorizontalStippleUp,
	"boxLightHorizontalUp":                   BoxLightHorizontalUp,
	"boxLightHorizontalUpStipple":            BoxLightHorizontalUpStipple,
	"boxLightHorizontalUpStippleDown":        BoxLightHorizontalUpStippleDown,
	"boxLightHorizontalUpStippleUp":          BoxLightHorizontalUpStippleUp,
	"boxLightHorizontalUpStippleUpLeft":      BoxLightHorizontalUpStippleUpLeft,
	"boxLightHorizontalUpStippleUpRight":     BoxLightHorizontalUpStippleUpRight,
	"boxLightRoundDownLeft":                  BoxLightRoundDownLeft,
	"boxLightRoundDownLeftStipple":           BoxLightRoundDownLeftStipple,
	"boxLightRoundDownLeftStippleInner":      BoxLightRoundDownLeftStippleInner,
	"boxLightRoundDownLeftStippleOuter":      BoxLightRoundDownLeftStippleOuter,
	"boxLightRoundDownRight":                 BoxLightRoundDownRight,
	"boxLightRoundDownRightStipple":          BoxLightRoundDownRightStipple,
	"boxLightRoundDownRightStippleInner":     BoxLightRoundDownRightStippleInner,
	"boxLightRoundDownRightStippleOuter":     BoxLightRoundDownRightStippleOuter,
	"boxLightRoundUpLeft":                    BoxLightRoundUpLeft,
	"boxLightRoundUpLeftStipple":             BoxLightRoundUpLeftStipple,
	"boxLightRoundUpLeftStippleInner":        BoxLightRoundUpLeftStippleInner,
	"boxLightRoundUpLeftStippleOuter":        BoxLightRoundUpLeftStippleOuter,
	"boxLightRoundUpRight":                   BoxLightRoundUpRight,
	"boxLightRoundUpRightStipple":            BoxLightRoundUpRightStipple,
	"boxLightRoundUpRightStippleInner":       BoxLightRoundUpRightStippleInner,
	"boxLightRoundUpRightStippleOuter":       BoxLightRoundUpRightStippleOuter,
	"boxLightUpLeft":                         BoxLightUpLeft,
	"boxLightUpLeftCircle":                   BoxLightUpLeftCircle,
	"boxLightUpLeftStipple":                  BoxLightUpLeftStipple,
	"boxLightUpLeftStippleInner":             BoxLightUpLeftStippleInner,
	"boxLightUpLeftStippleOuter":             BoxLightUpLeftStippleOuter,
	"boxLightUpRight":                        BoxLightUpRight,
	"boxLightUpRightCircle":                  BoxLightUpRightCircle,
	"boxLightUpRightStipple":                 BoxLightUpRightStipple,
	"boxLightUpRightStippleInner":            BoxLightUpRightStippleInner,
	"boxLightUpRightStippleOuter":            BoxLightUpRightStippleOuter,
	"boxLightVertical":                       BoxLightVertical,
	"boxLightVerticalCircle":                 BoxLightVerticalCircle,
	"boxLightVerticalHorizontal":             BoxLightVerticalHorizontal,
	"boxLightVerticalHorizontalStipple":      BoxLightVerticalHorizontalStipple,
	"boxLightVerticalHorizontalStippleDown":  BoxLightVerticalHorizontalStippleDown,
	"boxLightVerticalHorizontalStippleDownLeft":      BoxLightVerticalHorizontalStippleDownLeft,
	"boxLightVerticalHorizontalStippleDownRight":     BoxLightVerticalHorizontalStippleDownRight,
	"boxLightVerticalHorizontalStippleLeft":          BoxLightVerticalHorizontalStippleLeft,
	"boxLightVerticalHorizontalStippleLeftDownRight": BoxLightVerticalHorizontalStippleLeftDownRight,
	"boxLightVerticalHorizontalStippleLeftUpRight":   BoxLightVerticalHorizontalStippleLeftUpRight,
	"boxLightVerticalHorizontalStippleRight":         BoxLightVerticalHorizontalStippleRight,
	"boxLightVerticalHorizontalStippleRightDownLeft": BoxLightVerticalHorizontalStippleRightDownLeft,
	"boxLightVerticalHorizontalStippleRightUpLeft":   BoxLightVerticalHorizontalStippleRightUpLeft,
	"boxLightVerticalHorizontalStippleUp":            BoxLightVerticalHorizontalStippleUp,
	"boxLightVerticalHorizontalStippleUpLeft":        BoxLightVerticalHorizontalStippleUpLeft,
	"boxLightVerticalHorizontalStippleUpRight":       BoxLightVerticalHorizontalStippleUpRight,
	"boxLightVerticalLeft":                           BoxLightVerticalLeft,
	"boxLightVerticalLeftStipple":                    BoxLightVerticalLeftStipple,
	"boxLightVerticalLeftStippleDownLeft":            BoxLightVerticalLeftStippleDownLeft,
	"boxLightVerticalLeftStippleLeft":                BoxLightVerticalLeftStippleLeft,
	"boxLightVerticalLeftStippleUpLeft":              BoxLightVerticalLeftStippleUpLeft,
	"boxLightVerticalMenuDown":                       BoxLightVerticalMenuDown,
	"boxLightVerticalMenuLeft":                       BoxLightVerticalMenuLeft,
	"boxLightVerticalMenuRight":                      BoxLightVerticalMenuRight,
	"boxLightVerticalMenuUp":                         BoxLightVerticalMenuUp,
	"boxLightVerticalRight":                          BoxLightVerticalRight,
	"boxLightVerticalRightStipple":                   BoxLightVerticalRightStipple,
	"boxLightVerticalRightStippleDownRight":          BoxLightVerticalRightStippleDownRight,
	"boxLightVerticalRightStippleLeft":               BoxLightVerticalRightStippleLeft,
	"boxLightVerticalRightStippleRight":              BoxLightVerticalRightStippleRight,
	"boxLightVerticalRightStippleUpRight":            BoxLightVerticalRightStippleUpRight,
	"boxLightVerticalStipple":                        BoxLightVerticalStipple,
	"boxLightVerticalStippleLeft":                    BoxLightVerticalStippleLeft,
	"boxLightVerticalStippleRight":                   BoxLightVerticalStippleRight,
	"boxOuterLightAll":                               BoxOuterLightAll,
	"boxOuterLightDashedAll":                         BoxOuterLightDashedAll,
	"boxOuterLightDashedDown":                        BoxOuterLightDashedDown,
	"boxOuterLightDashedDownLeft":                    BoxOuterLightDashedDownLeft,
	"boxOuterLightDashedDownLeftRight":               BoxOuterLightDashedDownLeftRight,
	"boxOuterLightDashedDownRight":                   BoxOuterLightDashedDownRight,
	"boxOuterLightDashedFoldDownLeft":                BoxOuterLightDashedFoldDownLeft,
	"boxOuterLightDashedFoldDownRight":               BoxOuterLightDashedFoldDownRight,
	"boxOuterLightDashedFoldUpLeft":                  BoxOuterLightDashedFoldUpLeft,
	"boxOuterLightDashedFoldUpRight":                 BoxOuterLightDashedFoldUpRight,
	"boxOuterLightDashedLeft":                        BoxOuterLightDashedLeft,
	"boxOuterLightDashedLeftRight":                   BoxOuterLightDashedLeftRight,
	"boxOuterLightDashedRight":                       BoxOuterLightDashedRight,
	"boxOuterLightDashedUp":                          BoxOuterLightDashedUp,
	"boxOuterLightDashedUpDown":                      BoxOuterLightDashedUpDown,
	"boxOuterLightDashedUpDownLeft":                  BoxOuterLightDashedUpDownLeft,
	"boxOuterLightDashedUpDownRight":                 BoxOuterLightDashedUpDownRight,
	"boxOuterLightDashedUpLeft":                      BoxOuterLightDashedUpLeft,
	"boxOuterLightDashedUpLeftRight":                 BoxOuterLightDashedUpLeftRight,
	"boxOuterLightDashedUpRight":                     BoxOuterLightDashedUpRight,
	"boxOuterLightDown":                              BoxOuterLightDown,
	"boxOuterLightDownLeft":                          BoxOuterLightDownLeft,
	"boxOuterLightDownLeftRight":                     BoxOuterLightDownLeftRight,
	"boxOuterLightDownLeftStipple":                   BoxOuterLightDownLeftStipple,
	"boxOuterLightDownRight":                         BoxOuterLightDownRight,
	"boxOuterLightDownRightStipple":                  BoxOuterLightDownRightStipple,
	"boxOuterLightDownStipple":                       BoxOuterLightDownStipple,
	"boxOuterLightDownVerticalStipple":               BoxOuterLightDownVerticalStipple,
	"boxOuterLightDownVerticalStippleLeft":           BoxOuterLightDownVerticalStippleLeft,
	"boxOuterLightDownVerticalStippleRight":          BoxOuterLightDownVerticalStippleRight,
	"boxOuterLightLeft":                              BoxOuterLightLeft,
	"boxOuterLightLeftHorizontalStipple":             BoxOuterLightLeftHorizontalStipple,
	"boxOuterLightLeftHorizontalStippleDown":         BoxOuterLightLeftHorizontalStippleDown,
	"boxOuterLightLeftHorizontalStippleUp":           BoxOuterLightLeftHorizontalStippleUp,
	"boxOuterLightLeftRight":                         BoxOuterLightLeftRight,
	"boxOuterLightLeftRightStipple":                  BoxOuterLightLeftRightStipple,
	"boxOuterLightLeftStipple":                       BoxOuterLightLeftStipple,
	"boxOuterLightRight":                             BoxOuterLightRight,
	"boxOuterLightRightHorizontalStipple":            BoxOuterLightRightHorizontalStipple,
	"boxOuterLightRightHorizontalStippleDown":        BoxOuterLightRightHorizontalStippleDown,
	"boxOuterLightRightHorizontalStippleUp":          BoxOuterLightRightHorizontalStippleUp,
	"boxOuterLightRightStipple":                      BoxOuterLightRightStipple,
	"boxOuterLightRoundDownLeft":                     BoxOuterLightRoundDownLeft,
	"boxOuterLightRoundDownRight":                    BoxOuterLightRoundDownRight,
	"boxOuterLightRoundUpLeft":                       BoxOuterLightRoundUpLeft,
	"boxOuterLightRoundUpRight":                      BoxOuterLightRoundUpRight,
	"boxOuterLightUp":                                BoxOuterLightUp,
	"boxOuterLightUpDown":                            BoxOuterLightUpDown,
	"boxOuterLightUpDownLeft":                        BoxOuterLightUpDownLeft,
	"boxOuterLightUpDownRight":                       BoxOuterLightUpDownRight,
	"boxOuterLightUpDownStipple":                     BoxOuterLightUpDownStipple,
	"boxOuterLightUpLeft":                            BoxOuterLightUpLeft,
	"boxOuterLightUpLeftRight":                       BoxOuterLightUpLeftRight,
	"boxOuterLightUpLeftStipple":                     BoxOuterLightUpLeftStipple,
	"boxOuterLightUpRight":                           BoxOuterLightUpRight,
	"boxOuterLightUpRightStipple":                    BoxOuterLightUpRightStipple,
	"boxOuterLightUpStipple":                         BoxOuterLightUpStipple,
	"boxOuterLightUpVerticalStipple":                 BoxOuterLightUpVerticalStipple,
	"boxOuterLightUpVerticalStippleLeft":             BoxOuterLightUpVerticalStippleLeft,
	"boxOuterLightUpVerticalStippleRight":            BoxOuterLightUpVerticalStippleRight,
	"briefcase":                                      Briefcase,
	"broadcast":                                      Broadcast,
	"bug":                                            Bug,
	"bugFill":                                        BugFill,
	"calculator":                                     Calculator,
	"calendar":                                       Calendar,
	"calendarImport":                                 CalendarImport,
	"calendarMonth":                                  CalendarMonth,
	"cancel":                                         Cancel,
	"card":                                           Card,
	"cardText":                                       CardText,
	"cart":                                           Cart,
	"cast":                                           Cast,
	"castle":                                         Castle,
	"chartBar":                                       ChartBar,
	"chat":                                           Chat,
	"chatProcessing":                                 ChatProcessing,
	"check":                                          Check,
	"checkboxBlank":                                  CheckboxBlank,
	"checkboxCross":                                  CheckboxCross,
	"checkboxIntermediate":                           CheckboxIntermediate,
	"checkboxIntermediateVariant":                    CheckboxIntermediateVariant,
	"checkboxMarked":                                 CheckboxMarked,
	"checkerLarge":                                   CheckerLarge,
	"checkerMedium":                                  CheckerMedium,
	"checkerSmall":                                   CheckerSmall,
	"checkerboard":                                   Checkerboard,
	"chevronDown":                                    ChevronDown,
	"chevronDownCircle":                              ChevronDownCircle,
	"chevronLeft":                                    ChevronLeft,
	"chevronLeftCircle":                              ChevronLeftCircle,
	"chevronRight":                                   ChevronRight,
	"chevronRightCircle":                             ChevronRightCircle,
	"chevronUp":                                      ChevronUp,
	"chevronUpCircle":                                ChevronUpCircle,
	"circle":                                         Circle,
	"clipboard":                                      Clipboard,
	"clock":                                          Clock,
	"coffee":                                         Coffee,
	"coinCopper":                                     CoinCopper,
	"coinElectrum":                                   CoinElectrum,
	"coinGold":                                       CoinGold,
	"coinPlatinum":                                   CoinPlatinum,
	"coinSilver":                                     CoinSilver,
	"comment":                                        Comment,
	"commentText":                                    CommentText,
	"commentUser":                                    CommentUser,
	"compass":                                        Compass,
	"compassEastArrow":                               CompassEastArrow,
	"compassNorthArrow":                              CompassNorthArrow,
	"compassNorthEast":                               CompassNorthEast,
	"compassNorthWest":                               CompassNorthWest,
	"compassSouthArrow":                              CompassSouthArrow,
	"compassSouthEast":                               CompassSouthEast,
	"compassSouthWest":                               CompassSouthWest,
	"compassWestArrow":                               CompassWestArrow,
	"creditCard":                                     CreditCard,
	"crown":                                          Crown,
	"cube":                                           Cube,
	"cubeUnfolded":                                   CubeUnfolded,
	"dagger":                                         Dagger,
	"database":                                       Database,
	"device":                                         Device,
	"diamond":                                        Diamond,
	"division":                                       Division,
	"door":                                           Door,
	"doorBox":                                        DoorBox,
	"doorOpen":                                       DoorOpen,
	"dotHexagon":                                     DotHexagon,
	"dotOctagon":                                     DotOctagon,
	"download":                                       Download,
	"email":                                          Email,
	"eye":                                            Eye,
	"eyeFill":                                        EyeFill,
	"file":                                           File,
	"fill":                                           Fill,
	"filter":                                         Filter,
	"fire":                                           Fire,
	"flask":                                          Flask,
	"flaskEmpty":                                     FlaskEmpty,
	"flaskRoundBottom":                               FlaskRoundBottom,
	"flaskRoundBottomEmpty":                          FlaskRoundBottomEmpty,
	"floppyDisk":                                     FloppyDisk,
	"folder":                                         Folder,
	"folderOpen":                                     FolderOpen,
	"formatAlignBottom":                              FormatAlignBottom,
	"formatAlignCenter":                              FormatAlignCenter,
	"formatAlignJustify":                             FormatAlignJustify,
	"formatAlignLeft":                                FormatAlignLeft,
	"formatAlignRight":                               FormatAlignRight,
	"formatAlignTop":                                 FormatAlignTop,
	"formatBold":                                     FormatBold,
	"formatFloatLeft":                                FormatFloatLeft,
	"formatFloatRight":                               FormatFloatRight,
	"formatHorizontalAlignCenter":                    FormatHorizontalAlignCenter,
	"formatItalic":                                   FormatItalic,
	"formatLineSpacing":                              FormatLineSpacing,
	"formatVerticalAlignCenter":                      FormatVerticalAlignCenter,
	"gamepad":                                        Gamepad,
	"gamepadCenter":                                  GamepadCenter,
	"gamepadCenterFill":                              GamepadCenterFill,
	"gamepadDown":                                    GamepadDown,
	"gamepadDownFill":                                GamepadDownFill,
	"gamepadDownLeft":                                GamepadDownLeft,
	"gamepadDownLeftFill":                            GamepadDownLeftFill,
	"gamepadDownRight":                               GamepadDownRight,
	"gamepadDownRightFill":                           GamepadDownRightFill,
	"gamepadFill":                                    GamepadFill,
	"gamepadLeft":                                    GamepadLeft,
	"gamepadLeftFill":                                GamepadLeftFill,
	"gamepadRight":                                   GamepadRight,
	"gamepadRightFill":                               GamepadRightFill,
	"gamepadUp":                                      GamepadUp,
	"gamepadUpFill":                                  GamepadUpFill,
	"gamepadUpLeft":                                  GamepadUpLeft,
	"gamepadUpLeftFill":                              GamepadUpLeftFill,
	"gamepadUpRight":                                 GamepadUpRight,
	"gamepadUpRightFill":                             GamepadUpRightFill,
	"glasses":                                        Glasses,
	"heart":                                          Heart,
	"heartBroken":                                    HeartBroken,
	"help":                                           Help,
	"hexagon":                                        Hexagon,
	"image":                                          Image,
	"key":                                            Key,
	"label":                                          Label,
	"labelVariant":                                   LabelVariant,
	"led":                                            Led,
	"lightbulb":                                      Lightbulb,
	"lock":                                           Lock,
	"lockOpen":                                       LockOpen,
	"login":                                          Login,
	"logout":                                         Logout,
	"magnifyMinus":                                   MagnifyMinus,
	"magnifyPlus":                                    MagnifyPlus,
	"map":                                            Map,
	"menuBottomLeft":                                 MenuBottomLeft,
	"menuBottomRight":                                MenuBottomRight,
	"menuDown":                                       MenuDown,
	"menuDownFill":                                   MenuDownFill,
	"menuLeft":                                       MenuLeft,
	"menuLeftFill":                                   MenuLeftFill,
	"menuLeftRight":                                  MenuLeftRight,
	"menuRight":                                      MenuRight,
	"menuRightFill":                                  MenuRightFill,
	"menuTopLeft":                                    MenuTopLeft,
	"menuTopRight":                                   MenuTopRight,
	"menuUp":                                         MenuUp,
	"menuUpDown":                                     MenuUpDown,
	"menuUpFill":                                     MenuUpFill,
	"message":                                        Message,
	"messageProcessing":                              MessageProcessing,
	"messageText":                                    MessageText,
	"messageUser":                                    MessageUser,
	"microphone":                                     Microphone,
	"minus":                                          Minus,
	"minusBox":                                       MinusBox,
	"minusBoxFill":                                   MinusBoxFill,
	"minusCircle":                                    MinusCircle,
	"minusCircleFill":                                MinusCircleFill,
	"monitor":                                        Monitor,
	"monitorImage":                                   MonitorImage,
	"multiply":                                       Multiply,
	"musicNote":                                      MusicNote,
	"necklace":                                       Necklace,
	"note":                                           Note,
	"notebook":                                       Notebook,
	"notification":                                   Notification,
	"octagon":                                        Octagon,
	"octagonAlert":                                   OctagonAlert,
	"paperclip":                                      Paperclip,
	"pause":                                          Pause,
	"peace":                                          Peace,
	"pencil":                                         Pencil,
	"pickaxe":                                        Pickaxe,
	"pictogrammers":                                  Pictogrammers,
	"play":                                           Play,
	"plus":                                           Plus,
	"plusBox":                                        PlusBox,
	"plusBoxFill":                                    PlusBoxFill,
	"plusCircle":                                     PlusCircle,
	"plusCircleFill":                                 PlusCircleFill,
	"poll":                                           Poll,
	"pound":                                          Pound,
	"radiobox":                                       Radiobox,
	"radioboxMarked":                                 RadioboxMarked,
	"relativeScale":                                  RelativeScale,
	"removeCircle":                                   RemoveCircle,
	"rotateClockwise":                                RotateClockwise,
	"rotateCounterclockwise":                         RotateCounterclockwise,
	"script":                                         Script,
	"search":                                         Search,
	"shield":                                         Shield,
	"skull":                                          Skull,
	"speaker":                                        Speaker,
	"stop":                                           Stop,
	"sword":                                          Sword,
	"tableTopDoorHorizontal":                         TableTopDoorHorizontal,
	"tableTopDoorOneWayDown":                         TableTopDoorOneWayDown,
	"tableTopDoorOneWayLeft":                         TableTopDoorOneWayLeft,
	"tableTopDoorOneWayRight":                        TableTopDoorOneWayRight,
	"tableTopDoorOneWayUp":                           TableTopDoorOneWayUp,
	"tableTopDoorSecretHorizontal":                   TableTopDoorSecretHorizontal,
	"tableTopDoorSecretVertical":                     TableTopDoorSecretVertical,
	"tableTopDoorVertical":                           TableTopDoorVertical,
	"tableTopHorizontalRotateClockwise":              TableTopHorizontalRotateClockwise,
	"tableTopHorizontalRotateCounterclockwise": TableTopHorizontalRotateCounterclockwise,
	"tableTopHorizontalStairsAscendLeft":       TableTopHorizontalStairsAscendLeft,
	"tableTopHorizontalStairsAscendRight":      TableTopHorizontalStairsAscendRight,
	"tableTopHorizontalStairsDescendDown":      TableTopHorizontalStairsDescendDown,
	"tableTopHorizontalStairsDescendLeft":      TableTopHorizontalStairsDescendLeft,
	"tableTopHorizontalStairsDescendRight":     TableTopHorizontalStairsDescendRight,
	"tableTopHorizontalStairsDescendUp":        TableTopHorizontalStairsDescendUp,
	"tableTopSpiralStairsDown":                 TableTopSpiralStairsDown,
	"tableTopSpiralStairsLeft":                 TableTopSpiralStairsLeft,
	"tableTopSpiralStairsRight":                TableTopSpiralStairsRight,
	"tableTopSpiralStairsRoundDown":            TableTopSpiralStairsRoundDown,
	"tableTopSpiralStairsRoundLeft":            TableTopSpiralStairsRoundLeft,
	"tableTopSpiralStairsRoundRight":           TableTopSpiralStairsRoundRight,
	"tableTopSpiralStairsRoundUp":              TableTopSpiralStairsRoundUp,
	"tableTopSpiralStairsUp":                   TableTopSpiralStairsUp,
	"tableTopStairsDown":                       TableTopStairsDown,
	"tableTopStairsLeft":                       TableTopStairsLeft,
	"tableTopStairsRight":                      TableTopStairsRight,
	"tableTopStairsUp":                         TableTopStairsUp,
	"tableTopVerticalRotateClockwise":          TableTopVerticalRotateClockwise,
	"tableTopVerticalRotateCounterclockwise":   TableTopVerticalRotateCounterclockwise,
	"tableTopVerticalStairsAscendDown":         TableTopVerticalStairsAscendDown,
	"tableTopVerticalStairsAscendUp":           TableTopVerticalStairsAscendUp,
	"tag":                                      Tag,
	"tagText":                                  TagText,
	"target":                                   Target,
	"tent":                                     Tent,
	"terminal":                                 Terminal,
	"textBox":                                  TextBox,
	"textImage":                                TextImage,
	"tileCautionHeavy":                         TileCautionHeavy,
	"tileCautionThin":                          TileCautionThin,
	"tileDiamondHex":                           TileDiamondHex,
	"toggleSwitchOff":                          ToggleSwitchOff,
	"toggleSwitchOn":                           ToggleSwitchOn,
	"toolbox":                                  Toolbox,
	"tooltipAbove":                             TooltipAbove,
	"tooltipAboveAlert":                        TooltipAboveAlert,
	"tooltipAboveText":                         TooltipAboveText,
	"tooltipBelow":                             TooltipBelow,
	"tooltipBelowAlert":                        TooltipBelowAlert,
	"tooltipBelowText":                         TooltipBelowText,
	"tooltipEnd":                               TooltipEnd,
	"tooltipEndAlert":                          TooltipEndAlert,
	"tooltipEndText":                           TooltipEndText,
	"tooltipStart":                             TooltipStart,
	"tooltipStartAlert":                        TooltipStartAlert,
	"tooltipStartText":                         TooltipStartText,
	"toyBrick":                                 ToyBrick,
	"trash":                                    Trash,
	"umbrella":                                 Umbrella,
	"upload":                                   Upload,
	"volumeHigh":                               VolumeHigh,
	"volumeLow":                                VolumeLow,
	"volumeMedium":                             VolumeMedium,
	"volumeMute":                               VolumeMute,
	"wall":                                     Wall,
	"wallFill":                                 WallFill,
	"wallFront":                                WallFront,
	"wallFrontDamaged":                         WallFrontDamaged,
	"wallFrontGate":                            WallFrontGate,
	"water":                                    Water,
	"waterFill":                                WaterFill,
	"well":                                     Well,
}

func Account(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 3h4v1h1v1h1v4h-1v1h-1v1H9v-1H8V9H7V5h1V4h1zm1 5v1h2V8h1V6h-1V5h-2v1H9v2zm-3 4h8v1h2v1h1v1h1v4H3v-4h1v-1h1v-1h2zm-1 4H5v1h12v-1h-1v-1h-2v-1H8v1H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccountBox(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2h14v1h1v1h1v14h-1v1h-1v1H4v-1H3v-1H2V4h1V3h1zm0 14h1v-1h2v-1h8v1h2v1h1V5h-1V4H5v1H4zm12 2v-1h-2v-1H8v1H6v1zM9 5h4v1h1v1h1v4h-1v1h-1v1H9v-1H8v-1H7V7h1V6h1zm3 3V7h-2v1H9v2h1v1h2v-1h1V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alert(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 20H2v-1H1v-4h1v-2h1v-2h1V9h1V7h1V5h1V3h1V2h6v1h1v2h1v2h1v2h1v2h1v2h1v2h1v4h-1zM9 6H8v2H7v2H6v2H5v2H4v2H3v2h16v-2h-1v-2h-1v-2h-1v-2h-1V8h-1V6h-1V4H9zm1 1h2v6h-2zm0 7h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertBox(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 12h-2V6h2m0 10h-2v-2h2m6 6H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1m-1-1v-1h1V5h-1V4H5v1H4v12h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertBoxFill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 20H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1Zm-6-8V6h-2v6Zm0 4v-2h-2v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertCircle(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 12h-2V6h2Zm0 4h-2v-2h2Zm3 5H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2Zm-1-2v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertCircleFill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 21H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2Zm-3-9V6h-2v6Zm0 4v-2h-2v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertHexagon(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 12h-2V6h2m0 10h-2v-2h2m0 7h-2v-1H8v-1H6v-1H4v-1H2V5h2V4h2V3h2V2h2V1h2v1h2v1h2v1h2v1h2v12h-2v1h-2v1h-2v1h-2m0-1v-1h2v-1h2v-1h2V6h-2V5h-2V4h-2V3h-2v1H8v1H6v1H4v10h2v1h2v1h2v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertRhombus(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 12h-2V6h2Zm0 4h-2v-2h2Zm0 5h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-1H1v-2h1V9h1V8h1V7h1V6h1V5h1V4h1V3h1V2h1V1h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1Zm0-3v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6h-1V5h-1V4h-2v1H9v1H8v1H7v1H6v1H5v1H4v2h1v1h1v1h1v1h1v1h1v1h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertRhombusFill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-1H1v-2h1V9h1V8h1V7h1V6h1V5h1V4h1V3h1V2h1V1h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1Zm0-9V6h-2v6Zm0 4v-2h-2v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignHorizontalCenter(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 20h-2v-4H6v-4h4v-2H4V6h6V2h2v4h6v4h-6v2h4v4h-4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignHorizontalDistribute(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 16H8V6h6M4 20H2V2h2m16 18h-2V2h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignHorizontalLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 10H6V6h12m-4 10H6v-4h8M4 20H2V2h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignHorizontalRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 10H4V6h12m0 10H8v-4h8m4 8h-2V2h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignVerticalBottom(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 16H6V4h4m6 12h-4V8h4m4 12H2v-2h18Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignVerticalCenter(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 18H6v-6H2v-2h4V4h4v6h2V6h4v4h4v2h-4v4h-4v-4h-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignVerticalDistribute(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 4H2V2h18m-4 12H6V8h10m4 12H2v-2h18Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignVerticalTop(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 4H2V2h18m-4 12h-4V6h4m-6 12H6V6h4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaA(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2zm-1 3v1h1v9h-2v-4h-2v4H8V7h1V6zm-1 2h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaAfill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 8h2v2h-2zm5-7v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-2 5H9v1H8v9h2v-4h2v4h2V7h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaB(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2zM8 6h5v1h1v3h-1v2h1v3h-1v1H8zm2 2v2h2V8zm2 4h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaBfill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zM8 6v10h5v-1h1v-3h-1v-2h1V7h-1V6zm2 2h2v2h-2zm2 4v2h-2v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaC(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2zM9 6h4v1h1v2h-2V8h-2v6h2v-1h2v2h-1v1H9v-1H8V7h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaCfill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zM9 6v1H8v8h1v1h4v-1h1v-2h-2v1h-2V8h2v1h2V7h-1V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaD(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2zM8 6h5v1h1v1h1v6h-1v1h-1v1H8zm2 2v6h2v-1h1V9h-1V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaDfill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zM8 6v10h5v-1h1v-1h1V8h-1V7h-1V6zm2 2h2v1h1v4h-1v1h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaE(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2zM8 6h6v2h-4v2h4v2h-4v2h4v2H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaEfill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zM8 6v10h6v-2h-4v-2h4v-2h-4V8h4V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaF(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2zM8 6h6v2h-4v2h3v2h-3v4H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaFfill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zM8 6v10h2v-4h3v-2h-3V8h4V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaG(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2zM9 6h5v2h-4v6h2v-2h-1v-2h3v5h-1v1H9v-1H8V7h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaGfill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zM9 6v1H8v8h1v1h4v-1h1v-5h-3v2h1v2h-2V8h4V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaH(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2zM8 6h2v4h2V6h2v10h-2v-4h-2v4H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaHfill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zM8 6v10h2v-4h2v4h2V6h-2v4h-2V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaI(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2zm-1 3v2h-1v6h1v2H9v-2h1V8H9V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaIfill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-2 5H9v2h1v6H9v2h4v-2h-1V8h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaJ(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2zm-2 3h2v9h-1v1H9v-1H8v-2h2v1h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaJfill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-3 5v8h-2v-1H8v2h1v1h4v-1h1V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaK(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2zM8 6h2v2h1V7h1V6h2v2h-1v1h-1v1h-1v1h1v1h1v1h1v3h-2v-2h-1v-1h-1v3H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaKfill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zM8 6v10h2v-3h1v1h1v2h2v-3h-1v-1h-1v-1h-1v-1h1V9h1V8h1V6h-2v1h-1v1h-1V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaL(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2zM8 6h2v8h4v2H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaLfill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zM8 6v10h6v-2h-4V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaM(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2zM7 6h8v1h1v9h-2V8h-2v7h-2V8H8v8H6V7h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaMfill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zM7 6v1H6v9h2V8h2v7h2V8h2v8h2V7h-1V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaN(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2zM8 6h2v2h1v2h1V6h2v10h-2v-2h-1v-2h-1v4H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaNfill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zM8 6v10h2v-4h1v2h1v2h2V6h-2v4h-1V8h-1V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaO(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2zM9 6h4v1h1v8h-1v1H9v-1H8V7h1zm1 2v6h2V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaOfill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zM9 6v1H8v8h1v1h4v-1h1V7h-1V6zm1 2h2v6h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaP(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2zM8 6h5v1h1v4h-1v1h-3v4H8zm2 2v2h2V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaPfill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zM8 6v10h2v-4h3v-1h1V7h-1V6zm2 2h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaQ(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2zM9 6h4v1h1v8h1v1h-1v1h-1v-1h-1v-1h-1v1H9v-1H8V7h1zm1 2v6h1v-1h1V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaQfill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zM9 6v1H8v8h1v1h2v-1h1v1h1v1h1v-1h1v-1h-1V7h-1V6zm1 2h2v5h-1v1h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaR(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2zM8 6h5v1h1v4h-1v2h1v3h-2v-2h-1v-2h-1v4H8zm2 2v2h2V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaRfill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zM8 6v10h2v-4h1v2h1v2h2v-3h-1v-2h1V7h-1V6zm2 2h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaS(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2zM9 6h5v2h-4v2h3v1h1v4h-1v1H8v-2h4v-2H9v-1H8V7h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaSfill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zM9 6v1H8v4h1v1h3v2H8v2h5v-1h1v-4h-1v-1h-3V8h4V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaT(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2zM8 6h6v2h-2v8h-2V8H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaTfill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zM8 6v2h2v8h2V8h2V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaU(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2zM8 6h2v8h2V6h2v9h-1v1H9v-1H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaUfill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zM8 6v9h1v1h4v-1h1V6h-2v8h-2V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaV(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2zM7 6h2v3h1v2h2V9h1V6h2v4h-1v2h-1v2h-1v2h-2v-2H9v-2H8v-2H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaVfill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zM7 6v4h1v2h1v2h1v2h2v-2h1v-2h1v-2h1V6h-2v3h-1v2h-2V9H9V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaW(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2zM6 6h2v6h1v1h1V8h2v5h1v-1h1V6h2v7h-1v2h-1v1h-2v-1h-2v1H8v-1H7v-2H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaWfill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zM6 6v7h1v2h1v1h2v-1h2v1h2v-1h1v-2h1V6h-2v6h-1v1h-1V8h-2v5H9v-1H8V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaX(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2zm-1 13v-2h-1v-1h-2v1H9v2H7v-3h1v-1h1v-2H8V9H7V6h2v2h1v1h2V8h1V6h2v3h-1v1h-1v2h1v1h1v3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaXfill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-2 15h2v-3h-1v-1h-1v-2h1V9h1V6h-2v2h-1v1h-2V8H9V6H7v3h1v1h1v2H8v1H7v3h2v-2h1v-1h2v1h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaY(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2zM7 6h2v2h1v2h2V8h1V6h2v3h-1v2h-1v2h-1v3h-2v-3H9v-2H8V9H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaYfill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zM7 6v3h1v2h1v2h1v3h2v-3h1v-2h1V9h1V6h-2v2h-1v2h-2V8H9V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaZ(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2zM8 6h6v4h-1v1h-1v1h-1v1h-1v1h4v2H8v-4h1v-1h1v-1h1V9h1V8H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphaZfill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zM8 6v2h4v1h-1v1h-1v1H9v1H8v4h6v-2h-4v-1h1v-1h1v-1h1v-1h1V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Application(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 20H3v-1H2V3h1V2h16v1h1v16h-1ZM18 6V4H4v2Zm0 12V8H4v10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ApplicationCode(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 16H9v-1H8v-4h1v-1h2v2h-1v2h1m4 2h-2v-2h1v-2h-1v-2h2v1h1v4h-1m4 5H3v-1H2V3h1V2h16v1h1v16h-1M18 6V4H4v2m14 12V8H4v10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Apps(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 7H3V3h4m6 4H9V3h4m6 4h-4V3h4M7 13H3V9h4m6 4H9V9h4m6 4h-4V9h4M7 19H3v-4h4m6 4H9v-4h4m6 4h-4v-4h4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppsBox(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 8H6V6h2m4 2h-2V6h2m4 2h-2V6h2m-8 6H6v-2h2m4 2h-2v-2h2m4 2h-2v-2h2m-8 6H6v-2h2m4 2h-2v-2h2m4 2h-2v-2h2m2 6H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1m-1-1v-1h1V5h-1V4H5v1H4v12h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppsBoxFill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 20H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1M8 8V6H6v2m6 0V6h-2v2m6 0V6h-2v2m-6 4v-2H6v2m6 0v-2h-2v2m6 0v-2h-2v2m-6 4v-2H6v2m6 0v-2h-2v2m6 0v-2h-2v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Archive(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2h18v6h-1v12H3V8H2zm15 16V8H5v10zM4 4v2h14V4zm3 6h8v2H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arrow(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2h-3v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H9v1H8v1H7v1H6v1H3v1H2v1H1v1h1v1h1v1h1v1h1v-1h1v-1h1v-3h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1V9h1V8h1V7h1V6h1V5h2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowBottomLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 17H5V8h2v5h1v-1h1v-1h1v-1h1V9h1V8h1V7h1V6h1v1h1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H9v1h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowBottomLeftCircle(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 15V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2zm4 1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2zm9-1H7V8h2v3h1v-1h1V9h1V8h1V7h1v1h1v1h-1v1h-1v1h-1v1h-1v1h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowBottomRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 8v9H8v-2h5v-1h-1v-1h-1v-1h-1v-1H9v-1H8V9H7V8H6V7h1V6h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowBottomRightCircle(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 21H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2zm1-4h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2zm-1-9v7H8v-2h3v-1h-1v-1H9v-1H8V9H7V8h1V7h1v1h1v1h1v1h1v1h1V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 17h-2v-1H9v-1H8v-1H7v-1H6v-2h2v1h1v1h1V4h2v9h1v-1h1v-1h2v2h-1v1h-1v1h-1v1h-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownBold(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 10v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-2h5V2h8v8zm-4 2h-3V4H9v8H6v1h1v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownBox(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 16h-2v-1H9v-1H8v-1H7v-1H6v-2h2v1h1v1h1V6h2v6h1v-1h1v-1h2v2h-1v1h-1v1h-1v1h-1m6 5H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1m-1-1v-1h1V5h-1V4H5v1H4v12h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownCircle(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 10v2h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-2h2v1h1v1h1V6h2v6h1v-1h1v-1zm5-3v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2zm-4-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 12v2h1v1h1v1h1v1h1v1h2v-2H8v-1H7v-1h4v-1h2v-1h1v-2h1V3h-2v6h-1v2h-2v1H7v-1h1v-1h1V8H7v1H6v1H5v1H4v1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeftBox(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 15H7V8h2v3h1v-1h1V9h1V8h1V7h1v1h1v1h-1v1h-1v1h-1v1h-1v1h3m4 7H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1m-1-1v-1h1V5h-1V4H5v1H4v12h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 12v2h-1v1h-1v1h-1v1h-1v1h-2v-2h1v-1h1v-1h-4v-1H9v-1H8v-2H7V3h2v6h1v2h2v1h3v-1h-1v-1h-1V8h2v1h1v1h1v1h1v1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRightBox(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 15H8v-2h3v-1h-1v-1H9v-1H8V9H7V8h1V7h1v1h1v1h1v1h1v1h1V8h2m3 12H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1m-1-1v-1h1V5h-1V4H5v1H4v12h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 12v-2h1V9h1V8h1V7h1V6h2v2h-1v1H9v1h9v2H9v1h1v1h1v2H9v-1H8v-1H7v-1H6v-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftBold(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 20h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-2h1V9h1V8h1V7h1V6h1V5h1V4h1V3h1V2h2v5h8v8h-8zm-2-4v-3h8V9h-8V6H9v1H8v1H7v1H6v1H5v2h1v1h1v1h1v1h1v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftBox(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 16h-2v-1H9v-1H8v-1H7v-1H6v-2h1V9h1V8h1V7h1V6h2v2h-1v1h-1v1h6v2h-6v1h1v1h1m6 6H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1m-1-1v-1h1V5h-1V4H5v1H4v12h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftCircle(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 16h-2v-1H9v-1H8v-1H7v-1H6v-2h1V9h1V8h1V7h1V6h2v2h-1v1h-1v1h6v2h-6v1h1v1h1zm3 5H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2zm1-4h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 19H8v-1H7v-1H6v-1H5v-1H4v-2h2v1h1v1h1v-4h1V9h1V8h2V7h7v2h-6v1h-2v2h-1v3h1v-1h1v-1h2v2h-1v1h-1v1h-1v1h-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 10v2h-1v1h-1v1h-1v1h-1v1h-2v-2h1v-1h1v-1H7v1h1v1h1v2H7v-1H6v-1H5v-1H4v-1H3v-2h1V9h1V8h1V7h1V6h2v2H8v1H7v1h8V9h-1V8h-1V6h2v1h1v1h1v1h1v1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftUp(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 3H8v1H7v1H6v1H5v1H4v2h2V8h1V7h1v4h1v2h1v1h2v1h7v-2h-6v-1h-2v-2h-1V7h1v1h1v1h2V7h-1V6h-1V5h-1V4h-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 10v2h-1v1h-1v1h-1v1h-1v1h-2v-2h1v-1h1v-1H4v-2h9V9h-1V8h-1V6h2v1h1v1h1v1h1v1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightBold(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 2h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-5H2V7h8zm2 4v3H4v4h8v3h1v-1h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightBox(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 16h-2v-2h1v-1h1v-1H6v-2h6V9h-1V8h-1V6h2v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1m6 5H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1m-1-1v-1h1V5h-1V4H5v1H4v12h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightCircle(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2zM6 5H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6zm4 1h2v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-2v-2h1v-1h1v-1H6v-2h6V9h-1V8h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 19h2v-1h1v-1h1v-1h1v-1h1v-2h-2v1h-1v1h-1v-4h-1V9h-1V8h-2V7H3v2h6v1h2v2h1v3h-1v-1h-1v-1H8v2h1v1h1v1h1v1h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightUp(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3h2v1h1v1h1v1h1v1h1v2h-2V8h-1V7h-1v4h-1v2h-1v1h-2v1H3v-2h6v-1h2v-2h1V7h-1v1h-1v1H8V7h1V6h1V5h1V4h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTopLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 14V5h9v2H9v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h-1v1h-1v-1h-1v-1h-1v-1h-1v-1h-1v-1H9v-1H8V9H7v5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTopLeftCircle(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2zM6 5H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6zm1 9V7h7v2h-3v1h1v1h1v1h1v1h1v1h-1v1h-1v-1h-1v-1h-1v-1h-1v-1H9v3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTopRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 5h9v9h-2V9h-1v1h-1v1h-1v1h-1v1h-1v1H9v1H8v1H7v-1H6v-1h1v-1h1v-1h1v-1h1v-1h1V9h1V8h1V7H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTopRightCircle(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 7v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2zm-4-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6zM8 7h7v7h-2v-3h-1v1h-1v1h-1v1H9v1H8v-1H7v-1h1v-1h1v-1h1v-1h1V9H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUp(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 5h2v1h1v1h1v1h1v1h1v2h-2v-1h-1V9h-1v9h-2V9H9v1H8v1H6V9h1V8h1V7h1V6h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpBold(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 12v-2h1V9h1V8h1V7h1V6h1V5h1V4h1V3h1V2h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2h-5v8H7v-8zm4-2h3v8h4v-8h3V9h-1V8h-1V7h-1V6h-1V5h-2v1H9v1H8v1H7v1H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpBox(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 16h-2v-6H9v1H8v1H6v-2h1V9h1V8h1V7h1V6h2v1h1v1h1v1h1v1h1v2h-2v-1h-1v-1h-1m6 10H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1m-1-1v-1h1V5h-1V4H5v1H4v12h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpCircle(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 12v-2h1V9h1V8h1V7h1V6h2v1h1v1h1v1h1v1h1v2h-2v-1h-1v-1h-1v6h-2v-6H9v1H8v1zm-5 3V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2zm4 1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 3h2v1h1v1h1v1h1v1h1v2h-2V8h-1V7h-1v8h1v-1h1v-1h2v2h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-2h2v1h1v1h1V7H9v1H8v1H6V7h1V6h1V5h1V4h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 10V8h1V7h1V6h1V5h1V4h2v2H8v1H7v1h4v1h2v1h1v2h1v7h-2v-6h-1v-2h-2v-1H7v1h1v1h1v2H7v-1H6v-1H5v-1H4v-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeftBox(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 15h-1v-1h-1v-1h-1v-1h-1v-1H9v3H7V7h7v2h-3v1h1v1h1v1h1v1h1v1h-1m4 6H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1m-1-1v-1h1V5h-1V4H5v1H4v12h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 10V8h-1V7h-1V6h-1V5h-1V4h-2v2h1v1h1v1h-4v1H9v1H8v2H7v7h2v-6h1v-2h2v-1h3v1h-1v1h-1v2h2v-1h1v-1h1v-1h1v-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRightBox(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 15H8v-1H7v-1h1v-1h1v-1h1v-1h1V9H8V7h7v7h-2v-3h-1v1h-1v1h-1v1H9m9 6H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1m-1-1v-1h1V5h-1V4H5v1H4v12h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AspectRatio(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 10h-2V8h-2V6h4M9 16H5v-4h2v2h2m10 5H3v-1H2V4h1V3h16v1h1v14h-1m-1-1V5H4v12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Axe(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 3h2v1h1v1h1v1h2v1h2v1h1v1h1v2h-1v2h-1v1h-1v1h-2v1h-1v-1h-1v-1h-1v-2h-1v-1h-1v-1h-1V9H9V8H8V6h1V5h1V4h1m-1 6v1h1v2h-1v1H9v1H8v1H7v1H6v1H5v1H4v1H3v-1H2v-2h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BagPersonal(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 15H9v2H7v-2H5v4h12zm0-6h-1V8h-1V7H7v1H6v1H5v4h12zm-4 2H9v-1h1V9h2v1h1zM3 8h1V6h2V5h1V2h1V1h6v1h1v3h1v1h2v2h1v12h-1v1H4v-1H3zm6-5v2h4V3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BagPersonalFill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 8h1V6h2V5h1V2h1V1h6v1h1v3h1v1h2v2h1v12h-1v1H4v-1H3zm6-5v2h4V3zm8 11H5v1h2v2h1v-2h9zm-5-3h1V9h-1V8h-2v1H9v2h1v1h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bank(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 19H2v-2h2V9H2V7h1V6h2V5h2V4h2V3h4v1h2v1h2v1h2v1h1v2h-2v8h2m-3-8V7h-2V6h-2V5H9v1H7v1H5v2m3 8v-6H6v6m6 0v-6h-2v6m6 0v-6h-2v6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFifty(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 8h2v6H5zm3 0h2v6H8zm10-3v1h1v3h1v4h-1v3h-1v1H3v-1H2V6h1V5zm-1 2H4v8h13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryOneHundred(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 8h2v6H5zm3 0h2v6H8zm10-3v1h1v3h1v4h-1v3h-1v1H3v-1H2V6h1V5zm-1 2H4v8h13zm-6 1h2v6h-2zm3 0h2v6h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatterySeventyFive(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 8h2v6H5zm3 0h2v6H8zm10-3v1h1v3h1v4h-1v3h-1v1H3v-1H2V6h1V5zm-1 2H4v8h13zm-6 1h2v6h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryTwentyFive(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 8v6H5V8zm11-3v1h1v3h1v4h-1v3h-1v1H3v-1H2V6h1V5zm-1 2H4v8h13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryZero(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 5h15v1h1v3h1v4h-1v3h-1v1H3v-1H2V6h1zm1 2v8h13V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BattleAxe(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1h-4v1h-1v1H9v1H8v4h4v1h-1v1h-1v1H9v1H8v1H7v1H6v1H5v1H4v1H3v1H2v1H1v1h1v1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v4h4v-1h1v-1h1v-1h1V7h-4V5h-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Beer(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 15H7V8h2m3 7h-2V8h2m3 11H4V3h11v4h3v1h1v8h-1v1h-3m2-2V9h-2v6m-2 2V6H6v11Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 17H3v-1h1v-1h1V7h1V6h1V5h2V4h1V2h2v2h1v1h2v1h1v1h1v8h1v1h1m-7 4h-2v-1H9v-1h4v1h-1m3-4V8h-1V7h-1V6H9v1H8v1H7v7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blood(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 16h1v-2H5Zm1 1h1v-1H6Zm8 4H8v-1H6v-1H5v-1H4v-2H3v-3h1v-1h1v-2h1V9h1V7h1V5h1V3h1V1h2v2h1v2h1v2h1v2h1v1h1v2h1v1h1v3h-1v2h-1v1h-1v1h-2Zm-5-3v-1H7v1Zm3 1v-1H9v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Book(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 2h1V1h14v1h1v18h-1v1H4v-1H3zm8 7h-1V8H9v1H8v1H7V3H5v16h12V3h-5v7h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 2h12v1h1v17h-2v-1h-2v-1h-2v-1h-2v1H8v1H6v1H4V3h1zm1 2v13h1v-1h2v-1h1v-1h2v1h1v1h2v1h1V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderBottom(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 12h-2v-2h2zM4 12H2v-2h2zm8-8h-2V2h2zm4 0h-2V2h2zm4 4h-2V6h2zm0-4h-2V2h2zm0 12h-2v-2h2zM4 16H2v-2h2zM8 4H6V2h2zM4 4H2V2h2zm0 4H2V6h2zm16 12H2v-2h18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderBottomLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 4V2h2v2zm8 8v-2h2v2zm0 4v-2h2v2zM6 4V2h2v2zm12 4V6h2v2zm0-4V2h2v2zm-4 0V2h2v2zM2 20V2h2v16h16v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderBottomLeftRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 4V2h2v2zM6 4V2h2v2zm8 0V2h2v2zm4-2h2v18H2V2h2v16h14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderBottomRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 12H2v-2h2zm8-8h-2V2h2zm4 0h-2V2h2zM4 16H2v-2h2zM8 4H6V2h2zM4 4H2V2h2zm0 4H2V6h2zm16 12H2v-2h16V2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderInside(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 10h8V2h2v8h8v2h-8v8h-2v-8H2zm4 8h2v2H6zm-4-4h2v2H2zm0 4h2v2H2zM2 2h2v2H2zm0 4h2v2H2zm4-4h2v2H6zm8 0h2v2h-2zm4 0h2v2h-2zm0 4h2v2h-2zm-4 12h2v2h-2zm4 0h2v2h-2zm0-4h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 20v-2h2v2zm0-16V2h2v2zm8 8v-2h2v2zm0 4v-2h2v2zm-4 4v-2h2v2zm4 0v-2h2v2zM6 20v-2h2v2zM6 4V2h2v2zm12 4V6h2v2zm0-4V2h2v2zm-4 0V2h2v2zM2 20V2h2v18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderLeftRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2v2h-2V2zm0 16v2h-2v-2zM8 2v2H6V2zm8 0v2h-2V2zm0 16v2h-2v-2zm-8 0v2H6v-2zM4 2v18H2V2zm14 0h2v18h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderNone(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 10h2v2H2zm16 0h2v2h-2zm-8-8h2v2h-2zm0 16h2v2h-2zm-4 0h2v2H6zm-4-4h2v2H2zm0 4h2v2H2zM2 2h2v2H2zm0 4h2v2H2zm4-4h2v2H6zm8 0h2v2h-2zm4 0h2v2h-2zm0 4h2v2h-2zm-4 12h2v2h-2zm4 0h2v2h-2zm0-4h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderOutside(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2h18v18H2zm2 2v14h14V4zm6 2h2v2h-2zm0 4h2v2h-2zm-4 0h2v2H6zm8 0h2v2h-2zm-4 4h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2v2h-2V2zm0 16v2h-2v-2zm-8-8v2H2v-2zm0-4v2H2V6zm4-4v2H6V2zM4 2v2H2V2zm12 0v2h-2V2zm0 16v2h-2v-2zM4 14v2H2v-2zm0 4v2H2v-2zm4 0v2H6v-2zM20 2v18h-2V2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderTop(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 10h2v2H2zm16 0h2v2h-2zm-8 8h2v2h-2zm-4 0h2v2H6zm-4-4h2v2H2zm0 4h2v2H2zM2 6h2v2H2zm16 0h2v2h-2zm-4 12h2v2h-2zm4 0h2v2h-2zm0-4h2v2h-2zM2 2h18v2H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderTopBottom(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 10h2v2H2zm16 0h2v2h-2zM2 14h2v2H2zm0-8h2v2H2zm16 0h2v2h-2zm0 8h2v2h-2zM2 18h18v2H2zM2 4V2h18v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderTopLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 10h2v2h-2zm-8 8h2v2h-2zm-4 0h2v2H6zM18 6h2v2h-2zm-4 12h2v2h-2zm4 0h2v2h-2zm0-4h2v2h-2zM2 2h18v2H4v16H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderTopLeftBottom(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 10h2v2h-2zm0-4h2v2h-2zm0 8h2v2h-2zm2 4v2H2V2h18v2H4v14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderTopLeftRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 18v2h-2v-2zm4 0v2h-2v-2zm-8 0v2H6v-2zm-4 2H2V2h18v18h-2V4H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderTopRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 18v2h-2v-2zm-8-8v2H2v-2zm0-4v2H2V6zm12 12v2h-2v-2zM4 14v2H2v-2zm0 4v2H2v-2zm4 0v2H6v-2zM20 2v18h-2V4H2V2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderTopRightBottom(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 12H2v-2h2zm0 4H2v-2h2zm0-8H2V6h2zM2 4V2h18v18H2v-2h16V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bow(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 3h10v1h2v1h2v1h1v1h1v2h1v2h1v10h-2v-2h-2v-1h-1v-1h-1v-1h-1v-1h-1v-1h-1v-1H9v-1H8v-1H7v-1H6V9H5V8H4V7H3V5H1m15 13h1v-6h-1v-2h-1V8h-1V7h-2V6h-2V5H4v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BowArrow(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 3h10v1h2v1h3V4h1V2h3v3h-2v1h-1v3h1v2h1v10h-2v-2h-2v-1h-1v-1h-1v-1h-1v-1h-1v-1H9v1H8v1H7v3H6v1H5v1H4v-1H3v-1H2v-1H1v-1h1v-1h1v-1h3v-1h1v-1h1v-2H7v-1H6V9H5V8H4V7H3V5H1zm15 15h1v-6h-1v-2h-1V9h-1v1h-1v1h-1v1h-1v1h1v1h1v1h1v1h1v1h1zM12 7V6h-2V5H4v1h1v1h1v1h1v1h1v1h1v1h1v-1h1V9h1V8h1V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Box(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2h14v1h1v1h1v14h-1v1h-1v1H4v-1H3v-1H2V4h1V3h1zm13 3V4H5v1H4v12h1v1h12v-1h1V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDashedDownLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 12h-2v-2h2zm-4 0h-3v-2h3zm-5 0h-3V9h2v1h1zM12 0v2h-2V0zm0 4v3h-2V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDashedDownRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 0v2h-2V0zm0 4v3h-2V4zm0 5v3H9v-2h1V9zM.002 10H2v2H.002zM4 10h3v2H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDashedHorizontal(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 12H0v-2h2m5 2H4v-2h3m6 2H9v-2h4m5 2h-3v-2h3m4 2h-2v-2h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDashedUpLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 22v-2h2v2zm0-4v-3h2v3zm0-5v-3h3v2h-1v1zm12-1h-2v-2h2zm-4 0h-3v-2h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDashedUpRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 10h2v2H0zm4 0h3v2H4zm5 0h3v3h-2v-1H9zm1 12v-2h2v2zm0-4v-3h2v3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDashedVertical(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2h-2V0h2m0 7h-2V4h2m0 9h-2V9h2m0 9h-2v-3h2m0 7h-2v-2h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDashedVerticalHorizontal(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2h-2V0h2m0 7h-2V4h2M2 12H0v-2h2m5 2H4v-2h3m11 2h-3v-2h3m4 2h-2v-2h2m-10 3h-2v-1H9v-2h1V9h2v1h1v2h-1m0 6h-2v-3h2m0 7h-2v-2h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDoubleDownLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 10H12V0h2v8h8m0 6H8V0h2v12h12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDoubleDownRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 10H0V8h8V0h2m4 14H0v-2h12V0h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDoubleHorizontal(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 12h22v2H0zm0-4h22v2H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDoubleHorizontalDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 10H0V8h22M10 22H8v-8H0v-2h10m4 10h-2V12h10v2h-8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDoubleHorizontalLightDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 10H0V8h22M12 22h-2v-8H0v-2h22v2H12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDoubleHorizontalLightUp(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 10H0V8h10V0h2v8h10m0 6H0v-2h22Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDoubleHorizontalUp(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 10H0V8h8V0h2m12 10H12V0h2v8h8m0 6H0v-2h22Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDoubleRoundDownLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0v3h1v2h1v2h1v1h1v1h1v1h2v1h2v1h3v2h-4v-1h-2v-1h-2v-1h-1v-1h-1V9h-1V8h-1V6H9V4H8V0zm12 10h-2V9h-3V8h-1V7h-1V6h-1V5h-1V2h-1V0h2v1h1v3h1v1h1v1h1v1h3v1h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDoubleRoundDownRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 12h3v-1h2v-1h2V9h1V8h1V7h1V5h1V3h1V0h2v4h-1v2h-1v2h-1v1h-1v1H9v1H8v1H6v1H4v1H0zM10 0v2H9v3H8v1H7v1H6v1H5v1H2v1H0V8h1V7h3V6h1V5h1V4h1V1h1V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDoubleRoundUpLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 10h-3v1h-2v1h-2v1h-1v1h-1v1h-1v2h-1v2h-1v3H8v-4h1v-2h1v-2h1v-1h1v-1h1v-1h1v-1h2V9h2V8h4zM12 22v-2h1v-3h1v-1h1v-1h1v-1h1v-1h3v-1h2v2h-1v1h-3v1h-1v1h-1v1h-1v3h-1v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDoubleRoundUpRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22v-3h-1v-2h-1v-2H9v-1H8v-1H7v-1H5v-1H3v-1H0V8h4v1h2v1h2v1h1v1h1v1h1v1h1v2h1v2h1v4zM0 12h2v1h3v1h1v1h1v1h1v1h1v3h1v2H8v-1H7v-3H6v-1H5v-1H4v-1H1v-1H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDoubleUpLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 22H8V8h14v2H10m4 12h-2V12h10v2h-8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDoubleUpRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 22H8v-8H0v-2h10m4 10h-2V10H0V8h14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDoubleVertical(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22V0h2v22zm-4 0V0h2v22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDoubleVerticalHorizontal(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 10H0V8h8V0h2m12 10H12V0h2v8h8M10 22H8v-8H0v-2h10m4 10h-2V12h10v2h-8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDoubleVerticalLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 10H0V8h8V0h2Zm0 12H8v-8H0v-2h10Zm4 0h-2V0h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDoubleVerticalLightLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 22H8V12H0v-2h8V0h2m4 22h-2V0h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDoubleVerticalLightRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 22H8V0h2m4 22h-2V0h2v10h8v2h-8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDoubleVerticalRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 10H12V0h2v8h8M10 22H8V0h2m4 22h-2V12h10v2h-8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDownLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 12H10V0h2v10h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDownLeftCircle(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 15H9v-1H8v-1H7V9h1V8h1V7h1V0h2v7h1v1h1v1h1v1h7v2h-7v1h-1v1h-1zm-4-3h1v1h2v-1h1v-2h-1V9h-2v1H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDownLeftStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-1V1h1M6 3H5V2h1m8 2h1V3h-1M7 5h1V4H7m10 2h-1V5h1m4 1h-1V5h1M6 7H5V6h1m8 2h1V7h-1m4 1h1V7h-1m3 2h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1V7h1V6h-1V5h1V4h-1V3h1V2h-1V1h1V0h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1v1h-1M7 9h1V8H7m-1 3H5v-1h1m16 2H10V0h2v10h10M8 15H7v-1h1m5 1h1v-1h-1m4 1h1v-1h-1m1 2h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v1H9v-1h1v-1H9v-1H8v-1H7v1H6v-1h1v-1h1v-1H7V9H6V8h1V7h1V6H7V5H6V4h1V3h1V2H7V1H6V0h1v1h1V0h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h-1v1h-1v-1h-1v1h-1m-6 2h-1v-1h1m4 1h-1v-1h1m4 1h-1v-1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDownLeftStippleInner(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-1V1h1m-3 3h1V3h-1m3 3h-1V5h1m4 1h-1V5h1m-7 3h1V7h-1m4 1h1V7h-1m3 2h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1V7h1V6h-1V5h1V4h-1V3h1V2h-1V1h1V0h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1v1h-1m1 4H10V0h2v10h10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDownLeftStippleOuter(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3H5V2h1m1 3h1V4H7M6 7H5V6h1m1 3h1V8H7m-1 3H5v-1h1m16 2H10V0h2v10h10M8 15H7v-1h1m5 1h1v-1h-1m4 1h1v-1h-1m1 2h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v1H9v-1h1v-1H9v-1H8v-1H7v1H6v-1h1v-1h1v-1H7V9H6V8h1V7h1V6H7V5H6V4h1V3h1V2H7V1H6V0h1v1h1V0h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h-1v1h-1v-1h-1v1h-1m-6 2h-1v-1h1m4 1h-1v-1h1m4 1h-1v-1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDownRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 0v12H0v-2h10V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDownRightCircle(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 9v4h-1v1h-1v1H9v-1H8v-1H7v-1H0v-2h7V9h1V8h1V7h1V0h2v7h1v1h1v1zm-3 4v-1h1v-2h-1V9h-2v1H9v2h1v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDownRightStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-1V1h1M6 3H5V2h1m8 2h1V3h-1M3 6H2V5h1m4 0h1V4H7m10 2h-1V5h1M4 8h1V7H4m5 2H8V8H7v1H6V8H5v1H4V8H3v1H2V8H1v1H0V8h1V7H0V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1V6H7V5H6V4h1V3h1V2H7V1H6V0h1v1h1V0h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1m5 0h1V7h-1m3 3h-1V9h1m-3 3h1v-1h-1m3 3h-1v-1h1m-2 2h-1v-1h1M3 15h1v-1H3m4 1h1v-1H7m4 1h1v-1h-1m1 2h-1v-1h-1v-1H9v1H8v1H7v-1H6v-1H5v1H4v1H3v-1H2v-1H1v1H0v-1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v-1H0v-2h10V0h2v12h1v-1h1v-1h-1V9h1V8h-1V7h1V6h-1V5h1V4h-1V3h1V2h-1V1h1V0h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h-1v1h-1v-1h-1v1h-1v1h1v1h-1M2 17H1v-1h1m4 1H5v-1h1m4 1H9v-1h1m4 1h-1v-1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDownRightStippleInner(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3H5V2h1M3 6H2V5h1m4 0h1V4H7M4 8h1V7H4m5 2H8V8H7v1H6V8H5v1H4V8H3v1H2V8H1v1H0V8h1V7H0V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1V6H7V5H6V4h1V3h1V2H7V1H6V0h1v1h1V0h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1m3 4H0v-2h10V0h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightDownRightStippleOuter(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-1V1h1m-3 3h1V3h-1m3 3h-1V5h1m-3 3h1V7h-1m3 3h-1V9h1m-3 3h1v-1h-1m3 3h-1v-1h1m-2 2h-1v-1h1M3 15h1v-1H3m4 1h1v-1H7m4 1h1v-1h-1m1 2h-1v-1h-1v-1H9v1H8v1H7v-1H6v-1H5v1H4v1H3v-1H2v-1H1v1H0v-1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v-1H0v-2h10V0h2v12h1v-1h1v-1h-1V9h1V8h-1V7h1V6h-1V5h1V4h-1V3h1V2h-1V1h1V0h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h-1v1h-1v-1h-1v1h-1v1h1v1h-1M2 17H1v-1h1m4 1H5v-1h1m4 1H9v-1h1m4 1h-1v-1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightFoldDownLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 9h1V3h-6v1h1v1h1v1h1v1h1v1h1m4 4h-3v-1h-1v-1h-1V9h-1V8h-1V7h-1V6h-1V5h-1V4h-1V3h-1V0h2v1h9v9h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightFoldDownRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 9h1V8h1V7h1V6h1V5h1V4h1V3H3m0 9H0v-2h1V1h9V0h2v3h-1v1h-1v1H9v1H8v1H7v1H6v1H5v1H4v1H3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightFoldUpLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22h-2v-3h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h3v2h-1v9h-9m7-2v-6h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightFoldUpRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22h-2v-1H1v-9H0v-2h3v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1m-3 0v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightHorizontal(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 10h22v2H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightHorizontalCircle(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 15H9v-1H8v-1H7v-1H0v-2h7V9h1V8h1V7h4v1h1v1h1v1h7v2h-7v1h-1v1h-1m-1-1v-1h1v-2h-1V9h-2v1H9v2h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightHorizontalDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22h-2V12H0v-2h22v2H12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightHorizontalDownStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 6H2V5h1m4 1H6V5h1m5 1h-1V5h1m5 1h-1V5h1m4 1h-1V5h1M4 8h1V7H4m4 1h1V7H8m6 1h1V7h-1m4 1h1V7h-1m3 2h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8H9v1H8V8H7v1H6V8H5v1H4V8H3v1H2V8H1v1H0V8h1V7H0V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1v1h1V7h1V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1v1h-1M3 15h1v-1H3m4 1h1v-1H7m10 1h1v-1h-1M2 17H1v-1h1m4 1H5v-1h1m14 1h-1v-1h1m-6 2h1v-1h-1m-7 2h1v-1H7m10 2h-1v-1h1M6 21H5v-1h1m2 2H7v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1H6v-1H5v1H4v1H3v-1H2v-1H1v1H0v-1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8m4 1h-2V12H0v-2h22v2H12m4 10h-1v-1h-1v1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h-1v1h1v1h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightHorizontalDownStippleDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 15h1v-1H3m4 1h1v-1H7m10 1h1v-1h-1M2 17H1v-1h1m4 1H5v-1h1m14 1h-1v-1h1m-6 2h1v-1h-1m-7 2h1v-1H7m10 2h-1v-1h1M6 21H5v-1h1m2 2H7v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1H6v-1H5v1H4v1H3v-1H2v-1H1v1H0v-1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8m4 1h-2V12H0v-2h22v2H12m4 10h-1v-1h-1v1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h-1v1h1v1h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightHorizontalDownStippleDownLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 15h1v-1H3m4 1h1v-1H7m-5 3H1v-1h1m4 1H5v-1h1m1 3h1v-1H7m-1 3H5v-1h1m2 2H7v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1H6v-1H5v1H4v1H3v-1H2v-1H1v1H0v-1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8m4 1h-2V12H0v-2h22v2H12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightHorizontalDownStippleDownRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 15h1v-1h-1m3 3h-1v-1h1m-6 2h1v-1h-1m3 3h-1v-1h1m-5 3h-2V12H0v-2h22v2H12m4 10h-1v-1h-1v1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h-1v1h1v1h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightHorizontalMenuDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 12h-2v-1H9v-1H8V9H7V8H6V7H5V5h12v2h-1v1h-1v1h-1v1h-1v1h-1m0 6h-2v-1H9v-1H8v-1H7v-1H6v-1H0v-2h7v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h7v2h-6v1h-1v1h-1v1h-1v1h-1m0-7V8h1V7H9v1h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightHorizontalMenuLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 12H0v-2h5m17 2h-6v-2h6m-11 3h1V9h-1v1h-1v2h1m3 5h-2v-1h-1v-1h-1v-1H9v-1H8v-1H7v-2h1V9h1V8h1V7h1V6h1V5h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightHorizontalMenuRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 12H0v-2h6m16 2h-5v-2h5m-12 3h1v-1h1v-2h-1V9h-1m0 8H8V5h2v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightHorizontalMenuUp(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 12h-7v-1h-1v-1h-1V9h-1V8h-2v1H9v1H8v1H7v1H0v-2h6V9h1V8h1V7h1V6h1V5h2v1h1v1h1v1h1v1h1v1h6m-5 7H5v-2h1v-1h1v-1h1v-1h1v-1h1v-1h2v1h1v1h1v1h1v1h1v1h1m-4 0v-1h-1v-1h-2v1H9v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightHorizontalStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 6H2V5h1m4 1H6V5h1m5 1h-1V5h1m5 1h-1V5h1m4 1h-1V5h1M4 8h1V7H4m4 1h1V7H8m6 1h1V7h-1m4 1h1V7h-1m3 2h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8H9v1H8V8H7v1H6V8H5v1H4V8H3v1H2V8H1v1H0V8h1V7H0V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1v1h1V7h1V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1v1h-1m1 4H0v-2h22M3 15h1v-1H3m4 1h1v-1H7m6 1h1v-1h-1m4 1h1v-1h-1m5 2h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v-1H9v1H8v1H7v-1H6v-1H5v1H4v1H3v-1H2v-1H1v1H0v-1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h-1v1h1M2 17H1v-1h1m4 1H5v-1h1m5 1h-1v-1h1m5 1h-1v-1h1m4 1h-1v-1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightHorizontalStippleDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 12H0v-2h22M3 15h1v-1H3m4 1h1v-1H7m6 1h1v-1h-1m4 1h1v-1h-1m5 2h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v-1H9v1H8v1H7v-1H6v-1H5v1H4v1H3v-1H2v-1H1v1H0v-1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h-1v1h1M2 17H1v-1h1m4 1H5v-1h1m5 1h-1v-1h1m5 1h-1v-1h1m4 1h-1v-1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightHorizontalStippleUp(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 6H2V5h1m4 1H6V5h1m5 1h-1V5h1m5 1h-1V5h1m4 1h-1V5h1M4 8h1V7H4m4 1h1V7H8m6 1h1V7h-1m4 1h1V7h-1m3 2h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8H9v1H8V8H7v1H6V8H5v1H4V8H3v1H2V8H1v1H0V8h1V7H0V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1v1h1V7h1V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1v1h-1m1 4H0v-2h22Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightHorizontalUp(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 12H0v-2h10V0h2v10h10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightHorizontalUpStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-1V1h1M6 3H5V2h1m8 2h1V3h-1M3 6H2V5h1m4 0h1V4H7m10 2h-1V5h1m4 1h-1V5h1M4 8h1V7H4m5 2H8V8H7v1H6V8H5v1H4V8H3v1H2V8H1v1H0V8h1V7H0V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1V6H7V5H6V4h1V3h1V2H7V1H6V0h1v1h1V0h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1m5 0h1V7h-1m4 1h1V7h-1m3 2h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1V7h1V6h-1V5h1V4h-1V3h1V2h-1V1h1V0h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1v1h-1m1 4H0v-2h10V0h2v10h10M3 15h1v-1H3m4 1h1v-1H7m6 1h1v-1h-1m4 1h1v-1h-1m5 2h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v-1H9v1H8v1H7v-1H6v-1H5v1H4v1H3v-1H2v-1H1v1H0v-1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h-1v1h1M2 17H1v-1h1m4 1H5v-1h1m5 1h-1v-1h1m5 1h-1v-1h1m4 1h-1v-1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightHorizontalUpStippleDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 12H0v-2h10V0h2v10h10M3 15h1v-1H3m4 1h1v-1H7m6 1h1v-1h-1m4 1h1v-1h-1m5 2h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v-1H9v1H8v1H7v-1H6v-1H5v1H4v1H3v-1H2v-1H1v1H0v-1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h-1v1h1M2 17H1v-1h1m4 1H5v-1h1m5 1h-1v-1h1m5 1h-1v-1h1m4 1h-1v-1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightHorizontalUpStippleUp(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-1V1h1M6 3H5V2h1m8 2h1V3h-1M3 6H2V5h1m4 0h1V4H7m10 2h-1V5h1m4 1h-1V5h1M4 8h1V7H4m5 2H8V8H7v1H6V8H5v1H4V8H3v1H2V8H1v1H0V8h1V7H0V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1V6H7V5H6V4h1V3h1V2H7V1H6V0h1v1h1V0h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1m5 0h1V7h-1m4 1h1V7h-1m3 2h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1V7h1V6h-1V5h1V4h-1V3h1V2h-1V1h1V0h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1v1h-1m1 4H0v-2h10V0h2v10h10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightHorizontalUpStippleUpLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3H5V2h1M3 6H2V5h1m4 0h1V4H7M4 8h1V7H4m5 2H8V8H7v1H6V8H5v1H4V8H3v1H2V8H1v1H0V8h1V7H0V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1V6H7V5H6V4h1V3h1V2H7V1H6V0h1v1h1V0h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1m13 4H0v-2h10V0h2v10h10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightHorizontalUpStippleUpRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-1V1h1m-3 3h1V3h-1m3 3h-1V5h1m4 1h-1V5h1m-7 3h1V7h-1m4 1h1V7h-1m3 2h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1V7h1V6h-1V5h1V4h-1V3h1V2h-1V1h1V0h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1v1h-1m1 4H0v-2h10V0h2v10h10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightRoundDownLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 12h-4v-1h-3v-1h-1V9h-1V8h-1V7h-1V4h-1V0h2v3h1v3h1v1h1v1h1v1h3v1h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightRoundDownLeftStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3H5V2h1m12 1h-1V2h1m2 3h-1V4h1M7 5h1V4H7m8 1h1V4h-1M8 6h1V5H8m13 4h-1V8h-1V7h-1v1h-1V7h-1V6h-1V5h-1V4h1V3h-1V2h-1V1h1V0h1v1h-1v1h1v1h1v1h1v1h-1v1h1v1h1V6h1v1h1v1h1V7h1v1h-1M9 10H8V9h1m13 3h-4v-1h-3v-1h-1V9h-1V8h-1V7h-1V4h-1V0h2v3h1v3h1v1h1v1h1v1h3v1h3m-10 3h-1v-1h1m4 2h1v-1h-1m-1 3h-1v-1h1m2 0h1v-1h-1m5 2h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v-1h-1v1h-1v-1h1v-1h-1v-1h-1v-1h-1v1h-1v-1h1V9h-1V8H9V7H8v1H7V7h1V6H7V5H6V4h1V3h1V2H7V1H6V0h1v1h1V0h1v1H8v1h1v1H8v1h1v1h1v1H9v1h1v1h1v1h1v1h1v1h1v1h1v1h1v-1h1v1h1v1h1v-1h1v1h1v-1h1v1h-1v1h1m-2 2h-1v-1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightRoundDownLeftStippleInner(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 3h-1V2h1m2 3h-1V4h1m-5 1h1V4h-1m6 5h-1V8h-1V7h-1v1h-1V7h-1V6h-1V5h-1V4h1V3h-1V2h-1V1h1V0h1v1h-1v1h1v1h1v1h1v1h-1v1h1v1h1V6h1v1h1v1h1V7h1v1h-1m1 4h-4v-1h-3v-1h-1V9h-1V8h-1V7h-1V4h-1V0h2v3h1v3h1v1h1v1h1v1h3v1h3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightRoundDownLeftStippleOuter(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3H5V2h1m1 3h1V4H7m1 2h1V5H8m1 5H8V9h1m13 3h-4v-1h-3v-1h-1V9h-1V8h-1V7h-1V4h-1V0h2v3h1v3h1v1h1v1h1v1h3v1h3m-10 3h-1v-1h1m4 2h1v-1h-1m-1 3h-1v-1h1m2 0h1v-1h-1m5 2h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v-1h-1v1h-1v-1h1v-1h-1v-1h-1v-1h-1v1h-1v-1h1V9h-1V8H9V7H8v1H7V7h1V6H7V5H6V4h1V3h1V2H7V1H6V0h1v1h1V0h1v1H8v1h1v1H8v1h1v1h1v1H9v1h1v1h1v1h1v1h1v1h1v1h1v1h1v-1h1v1h1v1h1v-1h1v1h1v-1h1v1h-1v1h1m-2 2h-1v-1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightRoundDownRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 0v4h-1v3h-1v1H9v1H8v1H7v1H4v1H0v-2h3V9h3V8h1V7h1V6h1V3h1V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightRoundDownRightStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 2H4V1h1m12 1h-1V1h1M2 5H1V4h1m4 0h1V3H6m8 1h1V3h-1m2 4h-1V6h1M3 7h1V6H3M1 9H0V8h1V7h1V6h1V5h1v1h1V5h1V4H5V3h1V2h1V1h1V0h1v1H8v1H7v1h1v1H7v1H6v1H5v1H4v1H3V7H2v1H1m13 3h-1v-1h1M4 12H0v-2h3V9h3V8h1V7h1V6h1V3h1V0h2v4h-1v3h-1v1H9v1H8v1H7v1H4m7 3h-1v-1h1m-5 1h1v-1H6m-3 2h1v-1H3m1 2H3v-1H2v-1H1v1H0v-1h1v-1h1v1h1v-1h1v1h1v-1h1v-1h1v1h1v-1h1v-1h1v-1h1V9h1V8h1V7h-1V6h1V5h1V4h-1V3h1V2h-1V1h1V0h1v1h-1v1h1v1h1v1h-1v1h-1v1h-1v1h1v1h-1v1h-1v1h-1v1h1v1h-1v-1h-1v1H9v1H8v1H7v1H6v-1H5v1H4m5 1H8v-1h1m-7 2H1v-1h1m4 1H5v-1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightRoundDownRightStippleInner(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 2H4V1h1M2 5H1V4h1m4 0h1V3H6M3 7h1V6H3M1 9H0V8h1V7h1V6h1V5h1v1h1V5h1V4H5V3h1V2h1V1h1V0h1v1H8v1H7v1h1v1H7v1H6v1H5v1H4v1H3V7H2v1H1m3 4H0v-2h3V9h3V8h1V7h1V6h1V3h1V0h2v4h-1v3h-1v1H9v1H8v1H7v1H4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightRoundDownRightStippleOuter(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-1V1h1m-3 3h1V3h-1m2 4h-1V6h1m-2 5h-1v-1h1M4 12H0v-2h3V9h3V8h1V7h1V6h1V3h1V0h2v4h-1v3h-1v1H9v1H8v1H7v1H4m7 3h-1v-1h1m-5 1h1v-1H6m-3 2h1v-1H3m1 2H3v-1H2v-1H1v1H0v-1h1v-1h1v1h1v-1h1v1h1v-1h1v-1h1v1h1v-1h1v-1h1v-1h1V9h1V8h1V7h-1V6h1V5h1V4h-1V3h1V2h-1V1h1V0h1v1h-1v1h1v1h1v1h-1v1h-1v1h-1v1h1v1h-1v1h-1v1h-1v1h1v1h-1v-1h-1v1H9v1H8v1H7v1H6v-1H5v1H4m5 1H8v-1h1m-7 2H1v-1h1m4 1H5v-1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightRoundUpLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 22v-4h1v-3h1v-1h1v-1h1v-1h1v-1h3v-1h4v2h-3v1h-3v1h-1v1h-1v1h-1v3h-1v3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightRoundUpLeftStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 6h-1V5h1m4 1h-1V5h1m-7 2h-1V6h1m-2 3h-1V8h1m6 0h1V7h-1m-3 2h1V8h-1m-6 4H8v-1h1m-2 5H6v-1h1m11 1h1v-1h-1m3 3h-1v-1h1M7 19h1v-1H7m8 1h1v-1h-1m-9 3H5v-1h1m12 1h-1v-1h1M8 22H7v-1h1v-1H7v-1H6v-1h1v-1h1v-1h1v-1H8v-1h1v-1h1v-1h1v-1h-1v-1h1v1h1v-1h1V9h1V8h1V7h1v1h1V7h1V6h1v1h1v1h1V7h1v1h-1v1h-1V8h-1v1h-1V8h-1v1h-1v1h-1V9h-1v1h-1v1h-1v1h-1v1h-1v1H9v1h1v1H9v1H8v1h1v1H8v1h1v1H8m4 1h-2v-4h1v-3h1v-1h1v-1h1v-1h1v-1h3v-1h4v2h-3v1h-3v1h-1v1h-1v1h-1v3h-1m2 3h-1v-1h1v-1h1v-1h-1v-1h1v-1h1v-1h1v-1h1v-1h1v1h1v-1h1v-1h1v1h-1v1h-1v1h-1v1h-1v-1h-1v1h-1v1h1v1h-1v1h-1v1h-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightRoundUpLeftStippleInner(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 16h1v-1h-1m3 3h-1v-1h1m-6 2h1v-1h-1m3 3h-1v-1h1m-6 2h-2v-4h1v-3h1v-1h1v-1h1v-1h1v-1h3v-1h4v2h-3v1h-3v1h-1v1h-1v1h-1v3h-1m2 3h-1v-1h1v-1h1v-1h-1v-1h1v-1h1v-1h1v-1h1v-1h1v1h1v-1h1v-1h1v1h-1v1h-1v1h-1v1h-1v-1h-1v1h-1v1h1v1h-1v1h-1v1h-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightRoundUpLeftStippleOuter(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 6h-1V5h1m4 1h-1V5h1m-7 2h-1V6h1m-2 3h-1V8h1m6 0h1V7h-1m-3 2h1V8h-1m-6 4H8v-1h1m-2 5H6v-1h1m0 4h1v-1H7m-1 3H5v-1h1m2 2H7v-1h1v-1H7v-1H6v-1h1v-1h1v-1h1v-1H8v-1h1v-1h1v-1h1v-1h-1v-1h1v1h1v-1h1V9h1V8h1V7h1v1h1V7h1V6h1v1h1v1h1V7h1v1h-1v1h-1V8h-1v1h-1V8h-1v1h-1v1h-1V9h-1v1h-1v1h-1v1h-1v1h-1v1H9v1h1v1H9v1H8v1h1v1H8v1h1v1H8m4 1h-2v-4h1v-3h1v-1h1v-1h1v-1h1v-1h3v-1h4v2h-3v1h-3v1h-1v1h-1v1h-1v3h-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightRoundUpRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 10h4v1h3v1h1v1h1v1h1v1h1v3h1v4h-2v-3H9v-3H8v-1H7v-1H6v-1H3v-1H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightRoundUpRightStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 6H2V5h1m5 2H7V6h1M4 8h1V7H4m1 2h1V8H5m6 2h-1V9h1m3 4h-1v-1h1M3 18H2v-1h1m10 0h1v-1h-1m-7 2h1v-1H6m8 1h1v-1h-1m-9 3H4v-1h1m12 1h-1v-1h1m-9 3H7v-1h1v-1H7v-1H6v-1H5v-1h1v-1H5v-1H4v1H3v-1H2v-1H1v1H0v-1h1v-1h1v1h1v1h1v-1h1v1h1v1h1v1h1v1H7v1h1v1h1v1H8m4 1h-2v-3H9v-3H8v-1H7v-1H6v-1H3v-1H0v-2h4v1h3v1h1v1h1v1h1v1h1v3h1m4 4h-1v-1h-1v1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h-1v-1h1v-1h-1v-1h-1v-1h-1v-1H9v-1H8v-1H7V9H6v1H5V9H4V8H3v1H2V8H1v1H0V8h1V7H0V6h1v1h1v1h1V7h1V6h1v1h1v1h1v1h1V8h1v1H8v1h1v1h1v1h1v-1h1v1h-1v1h1v1h1v1h1v-1h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightRoundUpRightStippleInner(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 18H2v-1h1m3 1h1v-1H6m-1 3H4v-1h1m3 3H7v-1h1v-1H7v-1H6v-1H5v-1h1v-1H5v-1H4v1H3v-1H2v-1H1v1H0v-1h1v-1h1v1h1v1h1v-1h1v1h1v1h1v1h1v1H7v1h1v1h1v1H8m4 1h-2v-3H9v-3H8v-1H7v-1H6v-1H3v-1H0v-2h4v1h3v1h1v1h1v1h1v1h1v3h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightRoundUpRightStippleOuter(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 6H2V5h1m5 2H7V6h1M4 8h1V7H4m1 2h1V8H5m6 2h-1V9h1m3 4h-1v-1h1m-1 5h1v-1h-1m1 2h1v-1h-1m3 3h-1v-1h1m-5 3h-2v-3H9v-3H8v-1H7v-1H6v-1H3v-1H0v-2h4v1h3v1h1v1h1v1h1v1h1v3h1m4 4h-1v-1h-1v1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h-1v-1h1v-1h-1v-1h-1v-1h-1v-1H9v-1H8v-1H7V9H6v1H5V9H4V8H3v1H2V8H1v1H0V8h1V7H0V6h1v1h1v1h1V7h1V6h1v1h1v1h1v1h1V8h1v1H8v1h1v1h1v1h1v-1h1v1h-1v1h1v1h1v1h1v-1h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightUpLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 22V10h12v2H12v10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightUpLeftCircle(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 13V9h1V8h1V7h4v1h1v1h1v1h7v2h-7v1h-1v1h-1v1h-1v7h-2v-7H9v-1H8v-1zm3-4v1H9v2h1v1h2v-1h1v-2h-1V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightUpLeftStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 6H8V5h1m4 1h-1V5h1m4 1h-1V5h1m4 1h-1V5h1M8 8H7V7h1M6 9H5V8h1m4 0h1V7h-1m4 1h1V7h-1m4 1h1V7h-1M7 11h1v-1H7m-1 3H5v-1h1m1 3h1v-1H7m10 1h1v-1h-1M6 17H5v-1h1m14 1h-1v-1h1m-6 2h1v-1h-1m-7 2h1v-1H7m10 2h-1v-1h1M6 21H5v-1h1m6 2h-2V10H9v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1H7v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1H6v-1h1V9h1v1h1V9h1V8H9V7h1V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1v1h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1v1h12v2H12m4 10h-1v-1h-1v1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h-1v1h1v1h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightUpLeftStippleInner(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 15h1v-1h-1m3 3h-1v-1h1m-6 2h1v-1h-1m3 3h-1v-1h1m-5 3h-2V10h12v2H12m4 10h-1v-1h-1v1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h-1v1h1v1h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightUpLeftStippleOuter(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 6H8V5h1m4 1h-1V5h1m4 1h-1V5h1m4 1h-1V5h1M8 8H7V7h1M6 9H5V8h1m4 0h1V7h-1m4 1h1V7h-1m4 1h1V7h-1M7 11h1v-1H7m-1 3H5v-1h1m1 3h1v-1H7m-1 3H5v-1h1m1 3h1v-1H7m-1 3H5v-1h1m6 2h-2V10H9v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1H7v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1H6v-1h1V9h1v1h1V9h1V8H9V7h1V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1v1h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1v1h12v2H12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightUpRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 10h12v12h-2V12H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightUpRightCircle(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 7h4v1h1v1h1v4h-1v1h-1v1h-1v7h-2v-7H9v-1H8v-1H7v-1H0v-2h7V9h1V8h1zm4 3h-1V9h-2v1H9v2h1v1h2v-1h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightUpRightStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 6H2V5h1m4 1H6V5h1m4 1h-1V5h1m4 3h-1V7h1M4 8h1V7H4m4 1h1V7H8m9 5h-1v-1h1m-3 3h1v-1h-1M3 15h1v-1H3m4 1h1v-1H7m10 2h-1v-1h1M2 17H1v-1h1m4 1H5v-1h1m8 2h1v-1h-1m-7 2h1v-1H7m10 2h-1v-1h1M6 21H5v-1h1m2 2H7v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1H6v-1H5v1H4v1H3v-1H2v-1H1v1H0v-1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8m4 1h-2V12H0v-2h12m4 12h-1v-1h-1v1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1V9h-1V8h-1v1h-1V8H9v1H8V8H7v1H6V8H5v1H4V8H3v1H2V8H1v1H0V8h1V7h1v1h1V7h1V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1V6h1v1h-1v1h1v1h1v1h1V9h1v1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightUpRightStippleInner(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 15h1v-1H3m4 1h1v-1H7m-5 3H1v-1h1m4 1H5v-1h1m1 3h1v-1H7m-1 3H5v-1h1m2 2H7v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1H6v-1H5v1H4v1H3v-1H2v-1H1v1H0v-1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8m4 1h-2V12H0v-2h12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightUpRightStippleOuter(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 6H2V5h1m4 1H6V5h1m4 1h-1V5h1m4 3h-1V7h1M4 8h1V7H4m4 1h1V7H8m9 5h-1v-1h1m-3 3h1v-1h-1m3 3h-1v-1h1m-3 3h1v-1h-1m3 3h-1v-1h1m-5 3h-2V12H0v-2h12m4 12h-1v-1h-1v1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1V9h-1V8h-1v1h-1V8H9v1H8V8H7v1H6V8H5v1H4V8H3v1H2V8H1v1H0V8h1V7h1v1h1V7h1V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1V6h1v1h-1v1h1v1h1v1h1V9h1v1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVertical(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 0v22h-2V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalCircle(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22h-2v-7H9v-1H8v-1H7V9h1V8h1V7h1V0h2v7h1v1h1v1h1v4h-1v1h-1v1h-1m0-2v-1h1v-2h-1V9h-2v1H9v2h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalHorizontal(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22h-2V12H0v-2h10V0h2v10h10v2H12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalHorizontalStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-1V1h1M6 3H5V2h1m8 2h1V3h-1M3 6H2V5h1m4 0h1V4H7m10 2h-1V5h1m4 1h-1V5h1M4 8h1V7H4m5 2H8V8H7v1H6V8H5v1H4V8H3v1H2V8H1v1H0V8h1V7H0V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1V6H7V5H6V4h1V3h1V2H7V1H6V0h1v1h1V0h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1m5 0h1V7h-1m4 1h1V7h-1m3 2h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1V7h1V6h-1V5h1V4h-1V3h1V2h-1V1h1V0h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1v1h-1M3 15h1v-1H3m4 1h1v-1H7m10 1h1v-1h-1M2 17H1v-1h1m4 1H5v-1h1m14 1h-1v-1h1m-6 2h1v-1h-1m-7 2h1v-1H7m10 2h-1v-1h1M6 21H5v-1h1m2 2H7v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1H6v-1H5v1H4v1H3v-1H2v-1H1v1H0v-1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8m4 1h-2V12H0v-2h10V0h2v10h10v2H12m4 10h-1v-1h-1v1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h-1v1h1v1h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalHorizontalStippleDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 15h1v-1H3m4 1h1v-1H7m10 1h1v-1h-1M2 17H1v-1h1m4 1H5v-1h1m14 1h-1v-1h1m-6 2h1v-1h-1m-7 2h1v-1H7m10 2h-1v-1h1M6 21H5v-1h1m2 2H7v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1H6v-1H5v1H4v1H3v-1H2v-1H1v1H0v-1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8m4 1h-2V12H0v-2h10V0h2v10h10v2H12m4 10h-1v-1h-1v1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h-1v1h1v1h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalHorizontalStippleDownLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 15h1v-1h-1m3 3h-1v-1h1m-6 2h1v-1h-1m3 3h-1v-1h1m-5 3h-2V12H0v-2h10V0h2v10h10v2H12m4 10h-1v-1h-1v1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h-1v1h1v1h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalHorizontalStippleDownRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 15h1v-1H3m4 1h1v-1H7m-5 3H1v-1h1m4 1H5v-1h1m1 3h1v-1H7m-1 3H5v-1h1m2 2H7v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1H6v-1H5v1H4v1H3v-1H2v-1H1v1H0v-1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8m4 1h-2V12H0v-2h10V0h2v10h10v2H12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalHorizontalStippleLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3H5V2h1M3 6H2V5h1m4 0h1V4H7M4 8h1V7H4m5 2H8V8H7v1H6V8H5v1H4V8H3v1H2V8H1v1H0V8h1V7H0V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1V6H7V5H6V4h1V3h1V2H7V1H6V0h1v1h1V0h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1m-6 7h1v-1H3m4 1h1v-1H7m-5 3H1v-1h1m4 1H5v-1h1m1 3h1v-1H7m-1 3H5v-1h1m2 2H7v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1H6v-1H5v1H4v1H3v-1H2v-1H1v1H0v-1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8m4 1h-2V12H0v-2h10V0h2v10h10v2H12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalHorizontalStippleLeftDownRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3H5V2h1M3 6H2V5h1m4 0h1V4H7M4 8h1V7H4m5 2H8V8H7v1H6V8H5v1H4V8H3v1H2V8H1v1H0V8h1V7H0V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1V6H7V5H6V4h1V3h1V2H7V1H6V0h1v1h1V0h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1m-6 7h1v-1H3m4 1h1v-1H7m10 1h1v-1h-1M2 17H1v-1h1m4 1H5v-1h1m14 1h-1v-1h1m-6 2h1v-1h-1m-7 2h1v-1H7m10 2h-1v-1h1M6 21H5v-1h1m2 2H7v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1H6v-1H5v1H4v1H3v-1H2v-1H1v1H0v-1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8m4 1h-2V12H0v-2h10V0h2v10h10v2H12m4 10h-1v-1h-1v1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h-1v1h1v1h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalHorizontalStippleLeftUpRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-1V1h1M6 3H5V2h1m8 2h1V3h-1M3 6H2V5h1m4 0h1V4H7m10 2h-1V5h1m4 1h-1V5h1M4 8h1V7H4m5 2H8V8H7v1H6V8H5v1H4V8H3v1H2V8H1v1H0V8h1V7H0V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1V6H7V5H6V4h1V3h1V2H7V1H6V0h1v1h1V0h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1m5 0h1V7h-1m4 1h1V7h-1m3 2h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1V7h1V6h-1V5h1V4h-1V3h1V2h-1V1h1V0h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1v1h-1M3 15h1v-1H3m4 1h1v-1H7m-5 3H1v-1h1m4 1H5v-1h1m1 3h1v-1H7m-1 3H5v-1h1m2 2H7v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1H6v-1H5v1H4v1H3v-1H2v-1H1v1H0v-1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8m4 1h-2V12H0v-2h10V0h2v10h10v2H12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalHorizontalStippleRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-1V1h1m-3 3h1V3h-1m3 3h-1V5h1m4 1h-1V5h1m-7 3h1V7h-1m4 1h1V7h-1m3 2h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1V7h1V6h-1V5h1V4h-1V3h1V2h-1V1h1V0h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1v1h-1m-4 7h1v-1h-1m3 3h-1v-1h1m-6 2h1v-1h-1m3 3h-1v-1h1m-5 3h-2V12H0v-2h10V0h2v10h10v2H12m4 10h-1v-1h-1v1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h-1v1h1v1h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalHorizontalStippleRightDownLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-1V1h1m-3 3h1V3h-1m3 3h-1V5h1m4 1h-1V5h1m-7 3h1V7h-1m4 1h1V7h-1m3 2h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1V7h1V6h-1V5h1V4h-1V3h1V2h-1V1h1V0h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1v1h-1M3 15h1v-1H3m4 1h1v-1H7m10 1h1v-1h-1M2 17H1v-1h1m4 1H5v-1h1m14 1h-1v-1h1m-6 2h1v-1h-1m-7 2h1v-1H7m10 2h-1v-1h1M6 21H5v-1h1m2 2H7v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1H6v-1H5v1H4v1H3v-1H2v-1H1v1H0v-1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8m4 1h-2V12H0v-2h10V0h2v10h10v2H12m4 10h-1v-1h-1v1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h-1v1h1v1h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalHorizontalStippleRightUpLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-1V1h1M6 3H5V2h1m8 2h1V3h-1M3 6H2V5h1m4 0h1V4H7m10 2h-1V5h1m4 1h-1V5h1M4 8h1V7H4m5 2H8V8H7v1H6V8H5v1H4V8H3v1H2V8H1v1H0V8h1V7H0V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1V6H7V5H6V4h1V3h1V2H7V1H6V0h1v1h1V0h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1m5 0h1V7h-1m4 1h1V7h-1m3 2h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1V7h1V6h-1V5h1V4h-1V3h1V2h-1V1h1V0h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1v1h-1m-4 7h1v-1h-1m3 3h-1v-1h1m-6 2h1v-1h-1m3 3h-1v-1h1m-5 3h-2V12H0v-2h10V0h2v10h10v2H12m4 10h-1v-1h-1v1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h-1v1h1v1h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalHorizontalStippleUp(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-1V1h1M6 3H5V2h1m8 2h1V3h-1M3 6H2V5h1m4 0h1V4H7m10 2h-1V5h1m4 1h-1V5h1M4 8h1V7H4m5 2H8V8H7v1H6V8H5v1H4V8H3v1H2V8H1v1H0V8h1V7H0V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1V6H7V5H6V4h1V3h1V2H7V1H6V0h1v1h1V0h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1m5 0h1V7h-1m4 1h1V7h-1m3 2h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1V7h1V6h-1V5h1V4h-1V3h1V2h-1V1h1V0h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1v1h-1m-9 14h-2V12H0v-2h10V0h2v10h10v2H12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalHorizontalStippleUpLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3H5V2h1M3 6H2V5h1m4 0h1V4H7M4 8h1V7H4m5 2H8V8H7v1H6V8H5v1H4V8H3v1H2V8H1v1H0V8h1V7H0V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1V6H7V5H6V4h1V3h1V2H7V1H6V0h1v1h1V0h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1m3 14h-2V12H0v-2h10V0h2v10h10v2H12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalHorizontalStippleUpRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-1V1h1m-3 3h1V3h-1m3 3h-1V5h1m4 1h-1V5h1m-7 3h1V7h-1m4 1h1V7h-1m3 2h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1V7h1V6h-1V5h1V4h-1V3h1V2h-1V1h1V0h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1v1h-1m-9 14h-2V12H0v-2h10V0h2v10h10v2H12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22h-2V12H0v-2h10V0h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalLeftStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-1V1h1M6 3H5V2h1m8 2h1V3h-1M3 6H2V5h1m4 0h1V4H7m10 2h-1V5h1M4 8h1V7H4m5 2H8V8H7v1H6V8H5v1H4V8H3v1H2V8H1v1H0V8h1V7H0V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1V6H7V5H6V4h1V3h1V2H7V1H6V0h1v1h1V0h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1m5 0h1V7h-1m3 4h-1v-1h1m-3 4h1v-1h-1M3 15h1v-1H3m4 1h1v-1H7m10 2h-1v-1h1M2 17H1v-1h1m4 1H5v-1h1m8 2h1v-1h-1m-7 2h1v-1H7m10 2h-1v-1h1M6 21H5v-1h1m2 2H7v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1H6v-1H5v1H4v1H3v-1H2v-1H1v1H0v-1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8m4 1h-2V12H0v-2h10V0h2m4 22h-1v-1h-1v1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1V9h1V8h-1V7h1V6h-1V5h1V4h-1V3h1V2h-1V1h1V0h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalLeftStippleDownLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 15h1v-1H3m4 1h1v-1H7m-5 3H1v-1h1m4 1H5v-1h1m1 3h1v-1H7m-1 3H5v-1h1m2 2H7v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1H6v-1H5v1H4v1H3v-1H2v-1H1v1H0v-1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8m4 1h-2V12H0v-2h10V0h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalLeftStippleLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3H5V2h1M3 6H2V5h1m4 0h1V4H7M4 8h1V7H4m5 2H8V8H7v1H6V8H5v1H4V8H3v1H2V8H1v1H0V8h1V7H0V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1V6H7V5H6V4h1V3h1V2H7V1H6V0h1v1h1V0h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1m-6 7h1v-1H3m4 1h1v-1H7m-5 3H1v-1h1m4 1H5v-1h1m1 3h1v-1H7m-1 3H5v-1h1m2 2H7v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1H6v-1H5v1H4v1H3v-1H2v-1H1v1H0v-1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8m4 1h-2V12H0v-2h10V0h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalLeftStippleUpLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3H5V2h1M3 6H2V5h1m4 0h1V4H7M4 8h1V7H4m5 2H8V8H7v1H6V8H5v1H4V8H3v1H2V8H1v1H0V8h1V7H0V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1V6H7V5H6V4h1V3h1V2H7V1H6V0h1v1h1V0h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1m3 14h-2V12H0v-2h10V0h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalMenuDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 6h-2V0h2m0 15h-2v-1H9v-1H8v-1H7v-1H6v-1H5V8h12v2h-1v1h-1v1h-1v1h-1v1h-1m0 8h-2v-5h2m0-5v-1h1v-1H9v1h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalMenuLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 13h1V9h-1v1h-1v2h1m3 5h-2v-1h-1v-1h-1v-1h-1v-1h-1v-1h-1v-2h1V9h1V8h1V7h1V6h1V5h2m-5 17h-2v-6H9v-1H8v-1H7v-1H6v-1H5v-2h1V9h1V8h1V7h1V6h1V0h2v7h-1v1h-1v1H9v1H8v2h1v1h1v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalMenuRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 13h1v-1h1v-2H8V9H7Zm0 4H5V5h2v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1H9v1H8v1H7Zm5 5h-2v-7h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V0h2v6h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalMenuUp(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5h-2V0h2m5 14H5v-2h1v-1h1v-1h1V9h1V8h1V7h2v1h1v1h1v1h1v1h1v1h1m-5 10h-2v-6h2m1-4v-1h-1v-1h-2v1H9v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22h-2V0h2v10h10v2H12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalRightStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-1V1h1M6 3H5V2h1m8 2h1V3h-1M7 5h1V4H7m10 2h-1V5h1m4 1h-1V5h1M6 7H5V6h1m8 2h1V7h-1m4 1h1V7h-1m3 2h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1V7h1V6h-1V5h1V4h-1V3h1V2h-1V1h1V0h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1v1h-1M7 9h1V8H7m-1 4H5v-1h1m1 4h1v-1H7m10 1h1v-1h-1M6 17H5v-1h1m14 1h-1v-1h1m-6 2h1v-1h-1m-7 2h1v-1H7m10 2h-1v-1h1M6 21H5v-1h1m2 2H7v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1h1v-1H7V9H6V8h1V7h1V6H7V5H6V4h1V3h1V2H7V1H6V0h1v1h1V0h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8m4 1h-2V0h2v10h10v2H12m4 10h-1v-1h-1v1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h-1v1h1v1h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalRightStippleDownRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 15h1v-1h-1m3 3h-1v-1h1m-6 2h1v-1h-1m3 3h-1v-1h1m-5 3h-2V0h2v10h10v2H12m4 10h-1v-1h-1v1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h-1v1h1v1h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalRightStippleLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3H5V2h1m1 3h1V4H7M6 7H5V6h1m1 3h1V8H7m-1 4H5v-1h1m1 4h1v-1H7m-1 3H5v-1h1m1 3h1v-1H7m-1 3H5v-1h1m2 2H7v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1h1v-1H7V9H6V8h1V7h1V6H7V5H6V4h1V3h1V2H7V1H6V0h1v1h1V0h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8m4 1h-2V0h2v10h10v2H12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalRightStippleRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-1V1h1m-3 3h1V3h-1m3 3h-1V5h1m4 1h-1V5h1m-7 3h1V7h-1m4 1h1V7h-1m3 2h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1V7h1V6h-1V5h1V4h-1V3h1V2h-1V1h1V0h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1v1h-1m-4 7h1v-1h-1m3 3h-1v-1h1m-6 2h1v-1h-1m3 3h-1v-1h1m-5 3h-2V0h2v10h10v2H12m4 10h-1v-1h-1v1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h-1v1h1v1h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalRightStippleUpRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-1V1h1m-3 3h1V3h-1m3 3h-1V5h1m4 1h-1V5h1m-7 3h1V7h-1m4 1h1V7h-1m3 2h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1V7h1V6h-1V5h1V4h-1V3h1V2h-1V1h1V0h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1v1h-1m-9 14h-2V0h2v10h10v2H12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-1V1h1M6 3H5V2h1m8 2h1V3h-1M7 5h1V4H7m10 2h-1V5h1M6 7H5V6h1m8 2h1V7h-1M7 9h1V8H7m10 3h-1v-1h1M6 12H5v-1h1m8 3h1v-1h-1m-7 2h1v-1H7m10 2h-1v-1h1M6 17H5v-1h1m8 2h1v-1h-1m-7 2h1v-1H7m10 2h-1v-1h1M6 21H5v-1h1m2 2H7v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1h1v-1H7V9H6V8h1V7h1V6H7V5H6V4h1V3h1V2H7V1H6V0h1v1h1V0h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8m4 1h-2V0h2m4 22h-1v-1h-1v1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1V9h1V8h-1V7h1V6h-1V5h1V4h-1V3h1V2h-1V1h1V0h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalStippleLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3H5V2h1m1 3h1V4H7M6 7H5V6h1m1 3h1V8H7m-1 4H5v-1h1m1 4h1v-1H7m-1 3H5v-1h1m1 3h1v-1H7m-1 3H5v-1h1m2 2H7v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1h1v-1H7V9H6V8h1V7h1V6H7V5H6V4h1V3h1V2H7V1H6V0h1v1h1V0h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8m4 1h-2V0h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxLightVerticalStippleRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-1V1h1m-3 3h1V3h-1m3 3h-1V5h1m-3 3h1V7h-1m3 4h-1v-1h1m-3 4h1v-1h-1m3 3h-1v-1h1m-3 3h1v-1h-1m3 3h-1v-1h1m-5 3h-2V0h2m4 22h-1v-1h-1v1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1V9h1V8h-1V7h1V6h-1V5h1V4h-1V3h1V2h-1V1h1V0h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightAll(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h22v22H0zm2 2v18h18V2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDashedAll(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 0v2H2v2H0V0zM2 6v4H0V6zm0 6v4H0v-4zm0 6v2h2v2H0v-4zM6 0h4v2H6zm6 0h4v2h-4zm6 0h4v4h-2V2h-2zm0 22v-2h2v-2h2v4zm-2 0h-4v-2h4zm-6 0H6v-2h4zM20 6h2v4h-2zm0 6h2v4h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDashedDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 20h4v2H1zm6 0h3v2H7zm5 0h4v2h-4zm6 0h3v2h-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDashedDownLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 18h2v2h2v2H0zm6 2h4v2H6zm6 0h4v2h-4zm6 0h3v2h-3zM0 16v-4h2v4zm0-6V7h2v3zm0-5V1h2v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDashedDownLeftRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 18h2v2h2v2H0zm6 2h4v2H6zm6 0h4v2h-4zm6 0h2v-2h2v4h-4zM0 16v-4h2v4zm0-6V7h2v3zm0-5V1h2v4zm22-4v4h-2V1zm0 6v3h-2V7zm0 5v4h-2v-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDashedDownRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 22v-2h2v-2h2v4zm2-6v-4h2v4zm0-6V6h2v4zm0-6V1h2v3zm-4 18h-4v-2h4zm-6 0H7v-2h3zm-5 0H1v-2h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDashedFoldDownLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 4H4V2h4m6 2h-4V2h4M2 6H1V5H0V1h2v3h1v1H2m18 1h-2V4h-2V2h4M7 10H5V9H4V8H3V6h2v1h1v1h1m13 4h-2V8h2m-9 7h-1v-1H9v-1H8v-1H7v-1h1v-1h1v1h1v1h1v1h1v1h-1m9 4h-2v-4h2m-4 5h-2v-1h-1v-1h-1v-2h2v1h1v1h1m5 5h-4v-1h-1v-1h1v-1h1v1h3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDashedFoldDownRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 4H8V2h4m6 2h-4V2h4M4 6H2V2h4v2H4m17 2h-1V5h-1V4h1V1h2v4h-1m-4 5h-2V8h1V7h1V6h2v2h-1v1h-1M4 12H2V8h2m8 7h-1v-1h-1v-1h1v-1h1v-1h1v-1h1v1h1v1h-1v1h-1v1h-1m-8 4H2v-4h2m4 5H6v-2h1v-1h1v-1h2v2H9v1H8m-3 4H1v-2h3v-1h1v1h1v1H5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDashedFoldUpLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 3h-1V2h-1V1h1V0h4v2h-3m-4 5h-2V5h1V4h1V3h2v2h-1v1h-1m6 2h-2V4h2M9 12H8v-1H7v-1h1V9h1V8h1V7h1v1h1v1h-1v1h-1v1H9m11 3h-2v-4h2M5 16H3v-2h1v-1h1v-1h2v2H6v1H5m3 5H4v-2h4m6 2h-4v-2h4m6 2h-4v-2h2v-2h2M2 21H0v-4h1v-1h1v1h1v1H2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDashedFoldUpRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3H4V2H1V0h4v1h1v1H5m5 5H8V6H7V5H6V3h2v1h1v1h1M4 8H2V4h2m10 8h-1v-1h-1v-1h-1V9h-1V8h1V7h1v1h1v1h1v1h1v1h-1M4 14H2v-4h2m15 6h-2v-1h-1v-1h-1v-2h2v1h1v1h1M6 20H2v-4h2v2h2m6 2H8v-2h4m6 2h-4v-2h4m4 3h-2v-3h-1v-1h1v-1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDashedLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 21v-4h2v4zm0-6v-3h2v3zm0-5V6h2v4zm0-6V1h2v3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDashedLeftRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 21v-4h2v4zm0-6v-3h2v3zm0-5V6h2v4zm0-6V1h2v3zM0 21v-4h2v4zm0-6v-3h2v3zm0-5V6h2v4zm0-6V1h2v3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDashedRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 21v-4h2v4zm0-6v-3h2v3zm0-5V6h2v4zm0-6V1h2v3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDashedUp(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 0h4v2H1zm6 0h3v2H7zm5 0h4v2h-4zm6 0h3v2h-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDashedUpDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2h-4V0h4zm-6 0h-3V0h3zm-5 0H6V0h4zM4 2H1V0h3zm17 20h-4v-2h4zm-6 0h-3v-2h3zm-5 0H6v-2h4zm-6 0H1v-2h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDashedUpDownLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 0v2H2v2H0V0zM2 6v4H0V6zm0 6v4H0v-4zm0 6v2h2v2H0v-4zM6 0h4v2H6zm6 0h3v2h-3zm5 0h4v2h-4zm4 22h-4v-2h4zm-6 0h-3v-2h3zm-5 0H6v-2h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDashedUpDownRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 22v-2h2v-2h2v4zm2-6v-4h2v4zm0-6V6h2v4zm0-6V2h-2V0h4v4zm-4 18h-4v-2h4zm-6 0H7v-2h3zm-5 0H1v-2h4zM1 0h4v2H1zm6 0h3v2H7zm5 0h4v2h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDashedUpLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 0v2H2v2H0V0zM2 6v4H0V6zm0 6v4H0v-4zm0 6v3H0v-3zM6 0h4v2H6zm6 0h3v2h-3zm5 0h4v2h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDashedUpLeftRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 4h-2V2h-2V0h4zm-6-2h-4V0h4zm-6 0H6V0h4zM4 2H2v2H0V0h4zm18 4v4h-2V6zm0 6v3h-2v-3zm0 5v4h-2v-4zM0 21v-4h2v4zm0-6v-3h2v3zm0-5V6h2v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDashedUpRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 4h-2V2h-2V0h4zm-6-2h-4V0h4zm-6 0H6V0h4zM4 2H1V0h3zm18 4v4h-2V6zm0 6v3h-2v-3zm0 5v4h-2v-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 20h22v2H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDownLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h2v20h20v2H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDownLeftRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 22H0V0h2v20h18V0h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDownLeftStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 2H6V1h1M4 4h1V3H4m3 3H6V5h1M4 8h1V7H4m3 4H6v-1h1m-3 4h1v-1H4m3 3H6v-1h1m5 1h-1v-1h1m5 1h-1v-1h1m4 1h-1v-1h1M4 18h1v-1H4m4 1h1v-1H8m6 1h1v-1h-1m4 1h1v-1h-1m3 2h-1v-1h-1v1h-1v-1h-1v1h-1v-1h-1v1h-1v-1h-1v1h-1v-1h-1v1h-1v-1H9v1H8v-1H7v1H6v-1H5v1H4v-1H3v-1h1v-1H3v-1h1v-1H3v-1h1v-1H3v-1h1v-1H3V9h1V8H3V7h1V6H3V5h1V4H3V3h1V2H3V1h1V0h1v1H4v1h1v1h1v1H5v1H4v1h1v1h1v1H5v1H4v1h1v1H4v1h1v1h1v1H5v1H4v1h1v1h1v1h1v-1h1v-1h1v1h1v1h1v-1h1v1h1v-1h1v-1h1v1h1v1h1v-1h1v-1h1v1h1v1h1v-1h1v1h-1m1 4H0V0h2v20h20Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDownRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 22v-2h20V0h2v22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDownRightStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 2H6V1h1M4 4h1V3H4m3 3H6V5h1M4 8h1V7H4m3 4H6v-1h1m-3 4h1v-1H4m3 3H6v-1h1m5 1h-1v-1h1m5 1h-1v-1h1m4 1h-1v-1h1M4 18h1v-1H4m4 1h1v-1H8m6 1h1v-1h-1m4 1h1v-1h-1m3 2h-1v-1h-1v1h-1v-1h-1v1h-1v-1h-1v1h-1v-1h-1v1h-1v-1h-1v1h-1v-1H9v1H8v-1H7v1H6v-1H5v1H4v-1H3v-1h1v-1H3v-1h1v-1H3v-1h1v-1H3v-1h1v-1H3V9h1V8H3V7h1V6H3V5h1V4H3V3h1V2H3V1h1V0h1v1H4v1h1v1h1v1H5v1H4v1h1v1h1v1H5v1H4v1h1v1H4v1h1v1h1v1H5v1H4v1h1v1h1v1h1v-1h1v-1h1v1h1v1h1v-1h1v1h1v-1h1v-1h1v1h1v1h1v-1h1v-1h1v1h1v1h1v-1h1v1h-1m1 4H0V0h2v20h20Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDownStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 16H2v-1h1m4 1H6v-1h1m5 1h-1v-1h1m5 1h-1v-1h1m4 1h-1v-1h1M4 18h1v-1H4m4 1h1v-1H8m6 1h1v-1h-1m4 1h1v-1h-1m3 2h-1v-1h-1v1h-1v-1h-1v1h-1v-1h-1v1h-1v-1h-1v1h-1v-1h-1v1h-1v-1H9v1H8v-1H7v1H6v-1H5v1H4v-1H3v1H2v-1H1v1H0v-1h1v-1H0v-1h1v1h1v1h1v-1h1v-1h1v1h1v1h1v-1h1v-1h1v1h1v1h1v-1h1v1h1v-1h1v-1h1v1h1v1h1v-1h1v-1h1v1h1v1h1v-1h1v1h-1m1 4H0v-2h22Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDownVerticalStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-1V1h1M6 3H5V2h1m8 2h1V3h-1M7 5h1V4H7m10 3h-1V6h1M6 8H5V7h1m8 3h1V9h-1m-7 2h1v-1H7m10 2h-1v-1h1M6 13H5v-1h1m8 2h1v-1h-1M3 16H2v-1h1m4 0h1v-1H7m10 2h-1v-1h1m4 1h-1v-1h1M4 18h1v-1H4m5 2H8v-1H7v1H6v-1H5v1H4v-1H3v1H2v-1H1v1H0v-1h1v-1H0v-1h1v1h1v1h1v-1h1v-1h1v1h1v1h1v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1H6v-1h1V9h1V8H7V7h1V6H7V5H6V4h1V3h1V2H7V1H6V0h1v1h1V0h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1m5 0h1v-1h-1m4 1h1v-1h-1m3 2h-1v-1h-1v1h-1v-1h-1v1h-1v-1h-1v1h-1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1V9h1V8h-1V7h1V6h-1V5h1V4h-1V3h1V2h-1V1h1V0h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h1v-1h1v-1h1v1h1v1h1v-1h1v1h-1m1 4H0v-2h10V0h2v20h10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDownVerticalStippleLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3H5V2h1m1 3h1V4H7M6 8H5V7h1m1 4h1v-1H7m-1 3H5v-1h1m-3 4H2v-1h1m4 0h1v-1H7m-3 4h1v-1H4m5 2H8v-1H7v1H6v-1H5v1H4v-1H3v1H2v-1H1v1H0v-1h1v-1H0v-1h1v1h1v1h1v-1h1v-1h1v1h1v1h1v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1H6v-1h1V9h1V8H7V7h1V6H7V5H6V4h1V3h1V2H7V1H6V0h1v1h1V0h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1m13 4H0v-2h10V0h2v20h10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightDownVerticalStippleRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-1V1h1m-3 3h1V3h-1m3 4h-1V6h1m-3 4h1V9h-1m3 3h-1v-1h1m-3 3h1v-1h-1m3 3h-1v-1h1m4 1h-1v-1h1m-7 3h1v-1h-1m4 1h1v-1h-1m3 2h-1v-1h-1v1h-1v-1h-1v1h-1v-1h-1v1h-1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1V9h1V8h-1V7h1V6h-1V5h1V4h-1V3h1V2h-1V1h1V0h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h1v-1h1v-1h1v1h1v1h1v-1h1v1h-1m1 4H0v-2h10V0h2v20h10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 22V0h2v22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightLeftHorizontalStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 2H6V1h1M4 4h1V3H4m3 3H6V5h1m4 1h-1V5h1m5 1h-1V5h1m5 1h-1V5h1M4 8h1V7H4m4 1h1V7H8m4 1h1V7h-1m6 1h1V7h-1m3 2h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8H9v1H8V8H7v1H6V8H5v1H4V8H3V7h1V6H3V5h1V4H3V3h1V2H3V1h1V0h1v1H4v1h1v1h1v1H5v1H4v1h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1v1h1V7h1V6h1v1h1v1h1V7h1v1h-1M7 15h1v-1H7m4 1h1v-1h-1m6 1h1v-1h-1m-7 3H9v-1h1m5 1h-1v-1h1m5 1h-1v-1h1M4 18h1v-1H4m3 3H6v-1h1m-5 3H0V0h2v10h20v2H2m4 10H5v-1H4v1H3v-1h1v-1H3v-1h1v-1H3v-1h1v-1H3v-1h1v-1H3v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h-1v1h1v1h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v-1h-1v1h-1v1h-1v-1h-1v-1H9v1H8v1H7v-1H6v-1H5v1H4v1h1v1h1v1H5v1H4v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightLeftHorizontalStippleDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 15h1v-1H7m4 1h1v-1h-1m6 1h1v-1h-1m-7 3H9v-1h1m5 1h-1v-1h1m5 1h-1v-1h1M4 18h1v-1H4m3 3H6v-1h1m-5 3H0V0h2v10h20v2H2m4 10H5v-1H4v1H3v-1h1v-1H3v-1h1v-1H3v-1h1v-1H3v-1h1v-1H3v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h-1v1h1v1h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v-1h-1v1h-1v1h-1v-1h-1v-1H9v1H8v1H7v-1H6v-1H5v1H4v1h1v1h1v1H5v1H4v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightLeftHorizontalStippleUp(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 2H6V1h1M4 4h1V3H4m3 3H6V5h1m4 1h-1V5h1m5 1h-1V5h1m5 1h-1V5h1M4 8h1V7H4m4 1h1V7H8m4 1h1V7h-1m6 1h1V7h-1m3 2h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8H9v1H8V8H7v1H6V8H5v1H4V8H3V7h1V6H3V5h1V4H3V3h1V2H3V1h1V0h1v1H4v1h1v1h1v1H5v1H4v1h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1v1h1V7h1V6h1v1h1v1h1V7h1v1h-1M2 22H0V0h2v10h20v2H2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightLeftRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 22V0h2v22zm20 0V0h2v22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightLeftRightStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 2H6V1h1m9 2h-1V2h1M4 4h1V3H4m3 3H6V5h1m10 0h1V4h-1m-1 3h-1V6h1M4 8h1V7H4m13 2h1V8h-1M7 11H6v-1h1m9 2h-1v-1h1M4 14h1v-1H4m3 3H6v-1h1m10 0h1v-1h-1m-1 3h-1v-1h1M4 18h1v-1H4m3 3H6v-1h1m10 0h1v-1h-1m-1 3h-1v-1h1M2 22H0V0h2m4 22H5v-1H4v1H3v-1h1v-1H3v-1h1v-1H3v-1h1v-1H3v-1h1v-1H3v-1h1v-1H3v-1h1v-1H3V9h1V8H3V7h1V6H3V5h1V4H3V3h1V2H3V1h1V0h1v1H4v1h1v1h1v1H5v1H4v1h1v1h1v1H5v1H4v1h1v1H4v1h1v1h1v1H5v1H4v1h1v1h1v1H5v1H4v1h1v1h1m12 1h-1v-1h1v-1h-1v-1h-1v-1h1v-1h1v-1h-1v-1h-1v-1h1v-1h1v-1h-1v-1h1v-1h-1V9h-1V8h1V7h1V6h-1V5h-1V4h1V3h1V2h-1V1h-1V0h1v1h1V0h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1m4 1h-2V0h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightLeftStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 2H6V1h1M4 4h1V3H4m3 3H6V5h1M4 8h1V7H4m3 4H6v-1h1m-3 4h1v-1H4m3 3H6v-1h1m-3 3h1v-1H4m3 3H6v-1h1m-5 3H0V0h2m4 22H5v-1H4v1H3v-1h1v-1H3v-1h1v-1H3v-1h1v-1H3v-1h1v-1H3v-1h1v-1H3v-1h1v-1H3V9h1V8H3V7h1V6H3V5h1V4H3V3h1V2H3V1h1V0h1v1H4v1h1v1h1v1H5v1H4v1h1v1h1v1H5v1H4v1h1v1H4v1h1v1h1v1H5v1H4v1h1v1h1v1H5v1H4v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 22V0h2v22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightRightHorizontalStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3h-1V2h1M3 6H2V5h1m5 1H7V5h1m5 1h-1V5h1m4 0h1V4h-1M4 8h1V7H4m6 1h1V7h-1m4 1h1V7h-1m5 2h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8H9v1H8V8H7v1H6V8H5v1H4V8H3v1H2V8H1v1H0V8h1V7H0V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1v1h1V7h1V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1V6h-1V5h-1V4h1V3h1V2h-1V1h-1V0h1v1h1V0h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1M3 15h1v-1H3m6 1h1v-1H9m4 1h1v-1h-1m4 1h1v-1h-1M2 17H1v-1h1m5 1H6v-1h1m5 1h-1v-1h1m4 1h-1v-1h1m1 3h1v-1h-1m-1 3h-1v-1h1m2 2h-1v-1h1v-1h-1v-1h-1v-1h1v-1h1v-1h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v1H9v-1H8v-1H7v1H6v-1H5v1H4v1H3v-1H2v-1H1v1H0v-1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1m4 1h-2V12H0v-2h20V0h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightRightHorizontalStippleDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 15h1v-1H3m6 1h1v-1H9m4 1h1v-1h-1m4 1h1v-1h-1M2 17H1v-1h1m5 1H6v-1h1m5 1h-1v-1h1m4 1h-1v-1h1m1 3h1v-1h-1m-1 3h-1v-1h1m2 2h-1v-1h1v-1h-1v-1h-1v-1h1v-1h1v-1h-1v-1h-1v-1h-1v1h-1v1h-1v-1h-1v-1h-1v1h-1v1H9v-1H8v-1H7v1H6v-1H5v1H4v1H3v-1H2v-1H1v1H0v-1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1m4 1h-2V12H0v-2h20V0h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightRightHorizontalStippleUp(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3h-1V2h1M3 6H2V5h1m5 1H7V5h1m5 1h-1V5h1m4 0h1V4h-1M4 8h1V7H4m6 1h1V7h-1m4 1h1V7h-1m5 2h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8h-1v1h-1V8H9v1H8V8H7v1H6V8H5v1H4V8H3v1H2V8H1v1H0V8h1V7H0V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1v1h1V7h1V6h1v1h1v1h1V7h1V6h1v1h1v1h1V7h1V6h-1V5h-1V4h1V3h1V2h-1V1h-1V0h1v1h1V0h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1m3 14h-2V12H0v-2h20V0h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightRightStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3h-1V2h1m1 3h1V4h-1m-1 3h-1V6h1m1 3h1V8h-1m-1 4h-1v-1h1m1 4h1v-1h-1m-1 3h-1v-1h1m1 3h1v-1h-1m-1 3h-1v-1h1m2 2h-1v-1h1v-1h-1v-1h-1v-1h1v-1h1v-1h-1v-1h-1v-1h1v-1h1v-1h-1v-1h1v-1h-1V9h-1V8h1V7h1V6h-1V5h-1V4h1V3h1V2h-1V1h-1V0h1v1h1V0h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1m4 1h-2V0h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightRoundDownLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h2v5h1v4h1v2h1v2h1v1h1v1h1v1h1v1h2v1h2v1h4v1h5v2h-6v-1h-4v-1h-2v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-2H2v-2H1V6H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightRoundDownRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 22v-2h5v-1h4v-1h2v-1h2v-1h1v-1h1v-1h1v-1h1v-2h1V9h1V5h1V0h2v6h-1v4h-1v2h-1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-2v1h-2v1H6v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightRoundUpLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 0v2h-5v1h-4v1h-2v1H9v1H8v1H7v1H6v1H5v2H4v2H3v4H2v5H.004v-6H1v-4h1v-2h1V8h1V7h1V6h1V5h1V4h1V3h2V2h2V1h4V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightRoundUpRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 22h-2v-5h-1v-4h-1v-2h-1V9h-1V8h-1V7h-1V6h-1V5h-2V4H9V3H5V2H.004V0H6v1h4v1h2v1h2v1h1v1h1v1h1v1h1v1h1v2h1v2h1v4h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightUp(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h22v2H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightUpDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h22v2H0zm0 20h22v2H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightUpDownLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 22V0h22v2H2v18h20v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightUpDownRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 0v22H0v-2h20V2H0V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightUpDownStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 2H0V0h22M3 5h1V4H3m4 1h1V4H7m6 1h1V4h-1m4 1h1V4h-1m5 2h-1V5h-1V4h-1v1h-1v1h-1V5h-1V4h-1v1h-1v1h-1V5h-1V4h-1v1h-1V4H9v1H8v1H7V5H6V4H5v1H4v1H3V5H2V4H1v1H0V4h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h-1v1h1M2 7H1V6h1m4 1H5V6h1m5 1h-1V6h1m5 1h-1V6h1m4 1h-1V6h1M3 16H2v-1h1m4 1H6v-1h1m5 1h-1v-1h1m5 1h-1v-1h1m4 1h-1v-1h1M4 18h1v-1H4m4 1h1v-1H8m6 1h1v-1h-1m4 1h1v-1h-1m3 2h-1v-1h-1v1h-1v-1h-1v1h-1v-1h-1v1h-1v-1h-1v1h-1v-1h-1v1h-1v-1H9v1H8v-1H7v1H6v-1H5v1H4v-1H3v1H2v-1H1v1H0v-1h1v-1H0v-1h1v1h1v1h1v-1h1v-1h1v1h1v1h1v-1h1v-1h1v1h1v1h1v-1h1v1h1v-1h1v-1h1v1h1v1h1v-1h1v-1h1v1h1v1h1v-1h1v1h-1m1 4H0v-2h22Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightUpLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 0v2H2v20H0V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightUpLeftRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h22v22h-2V2H2v20H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightUpLeftStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 5h1V4H5m2 1h1V4H7m6 1h1V4h-1m4 1h1V4h-1M4 6h1V5H4m7 2h-1V6h1m5 1h-1V6h1m4 1h-1V6h1M8 8H7V7h1M4 8h1V7H4m3 4H6v-1h1m-3 4h1v-1H4m3 3H6v-1h1m-3 3h1v-1H4m3 3H6v-1h1m-5 3H0V0h22v2H2m4 20H5v-1H4v1H3v-1h1v-1H3v-1h1v-1H3v-1h1v-1H3v-1h1v-1H3v-1h1v-1H3v-1h1v-1H3V9h1V8H3V7h1V6H3V5h1V4H3V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h-1v1h1v1h-1V5h-1V4h-1v1h-1v1h-1V5h-1V4h-1v1h-1v1h-1V5h-1V4h-1v1h-1V4H9v1H8v1H7V5H6v1H5v1h1v1H5v1H4v1h1v1H4v1h1v1h1v1H5v1H4v1h1v1h1v1H5v1H4v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightUpRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 22h-2V2H.002V0H22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightUpRightStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 5h1V4H3m4 1h1V4H7m6 1h1V4h-1m4 1h1V4h-1M2 7H1V6h1m4 1H5V6h1m5 1h-1V6h1m5 1h-1V6h1m1 3h1V8h-1m-1 4h-1v-1h1m1 4h1v-1h-1m-1 3h-1v-1h1m1 3h1v-1h-1m-1 3h-1v-1h1m2 2h-1v-1h1v-1h-1v-1h-1v-1h1v-1h1v-1h-1v-1h-1v-1h1v-1h1v-1h-1v-1h1v-1h-1V9h-1V8h1V7h1V6h-1V5h-1V4h-1v1h-1v1h-1V5h-1V4h-1v1h-1V4H9v1H8v1H7V5H6V4H5v1H4v1H3V5H2V4H1v1H0V4h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1m4 1h-2V2H0V0h22Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightUpStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 2H0V0h22M3 5h1V4H3m4 1h1V4H7m6 1h1V4h-1m4 1h1V4h-1m5 2h-1V5h-1V4h-1v1h-1v1h-1V5h-1V4h-1v1h-1v1h-1V5h-1V4h-1v1h-1V4H9v1H8v1H7V5H6V4H5v1H4v1H3V5H2V4H1v1H0V4h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h-1v1h1M2 7H1V6h1m4 1H5V6h1m5 1h-1V6h1m5 1h-1V6h1m4 1h-1V6h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightUpVerticalStipple(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 5h1V4H3m4 1h1V4H7m10 1h1V4h-1M2 7H1V6h1m4 1H5V6h1m14 1h-1V6h1m-6 2h1V7h-1M7 9h1V8H7m10 2h-1V9h1M6 11H5v-1h1m8 2h1v-1h-1m-7 2h1v-1H7m10 3h-1v-1h1M6 16H5v-1h1m8 3h1v-1h-1m-7 2h1v-1H7m10 2h-1v-1h1M6 21H5v-1h1m2 2H7v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7V9H6V8h1V7h1V6H7V5H6V4H5v1H4v1H3V5H2V4H1v1H0V4h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8m4 1h-2V2H0V0h22v2H12m4 20h-1v-1h-1v1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1V9h1V8h-1V7h1V6h-1V5h1V4h-1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h-1v1h1v1h-1V5h-1V4h-1v1h-1v1h-1V5h-1V4h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightUpVerticalStippleLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 5h1V4H3m4 1h1V4H7M2 7H1V6h1m4 1H5V6h1m1 3h1V8H7m-1 3H5v-1h1m1 3h1v-1H7m-1 4H5v-1h1m1 4h1v-1H7m-1 3H5v-1h1m2 2H7v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7v-1h1v-1H7v-1H6v-1h1v-1h1v-1H7V9H6V8h1V7h1V6H7V5H6V4H5v1H4v1H3V5H2V4H1v1H0V4h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8v1h1v1H8m4 1h-2V2H0V0h22v2H12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOuterLightUpVerticalStippleRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 5h1V4h-1m3 3h-1V6h1m-6 2h1V7h-1m3 3h-1V9h1m-3 3h1v-1h-1m3 4h-1v-1h1m-3 4h1v-1h-1m3 3h-1v-1h1m-5 3h-2V2H0V0h22v2H12m4 20h-1v-1h-1v1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1v-1h-1V9h1V8h-1V7h1V6h-1V5h1V4h-1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h1V3h1v1h-1v1h1v1h-1V5h-1V4h-1v1h-1v1h-1V5h-1V4h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h-1v1h1v1h1v1h-1v1h-1v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Briefcase(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 6h5V3h1V2h6v1h1v3h5v1h1v12h-1v1H2v-1H1V7h1zm7 0h4V4H9zm10 2H3v10h16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Broadcast(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 13h-2v-1H9v-2h1V9h2v1h1v2h-1m3 4h-1v-1h-1v-1h1v-1h1V9h-1V8h-1V7H9v1H8v1H7v4h1v1h1v1H8v1H7v-1H6v-1H5V8h1V7h1V6h1V5h6v1h1v1h1v1h1v6h-1v1h-1m2 5h-1v-1h-1v-1h1v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h1v1H6v1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bug(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7h3V6h1V5h1V4H7V3H6V2h1V1h1v1h1v1h1v1h2V3h1V2h1V1h1v1h1v1h-1v1h-1v1h1v1h1v1h3v2h-2v2h2v2h-2v2h2v2h-3v1h-1v1h-1v1H8v-1H7v-1H6v-1H3v-2h2v-2H3v-2h2V9H3zm10 11v-1h1v-1h1V8h-1V7h-1V6H9v1H8v1H7v8h1v1h1v1zm-4-5h4v2H9zm0-4h4v2H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BugFill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7h3V6h1V5h1V4H7V3H6V2h1V1h1v1h1v1h1v1h2V3h1V2h1V1h1v1h1v1h-1v1h-1v1h1v1h1v1h3v2h-2v2h2v2h-2v2h2v2h-3v1h-1v1h-1v1H8v-1H7v-1H6v-1H3v-2h2v-2H3v-2h2V9H3zm6 6v2h4v-2zm0-4v2h4V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calculator(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 21H4v-1H3V3h1V2h14v1h1v17h-1ZM16 7V4H6v3Zm-8 4V9H6v2Zm4 0V9h-2v2Zm4 0V9h-2v2Zm-8 4v-2H6v2Zm4 0v-2h-2v2Zm4 0v-2h-2v2Zm-8 4v-2H6v2Zm4 0v-2h-2v2Zm4 0v-2h-2v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 20H3v-1H2V3h1V2h2V0h2v2h8V0h2v2h2v1h1v16h-1zM4 4v2h14V4zm0 4v10h14V8zm8 4h4v4h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarImport(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 19h-6v-2h5V9H4v8h5v2H3v-1H2V4h1V3h2V1h2v2h8V1h2v2h2v1h1v14h-1m-7 3h-2v-6H7v-1h1v-1h1v-1h1v-1h2v1h1v1h1v1h1v1h-3m6-8V5H4v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarMonth(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 20H3v-1H2V3h1V2h2V0h2v2h8V0h2v2h2v1h1v16h-1zM4 4v2h14V4zm0 4v10h14V8zm10 6h2v2h-2zm-4 0h2v2h-2zm-4 0h2v2H6zm0-4h2v2H6zm4 0h2v2h-2zm4 0h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cancel(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 7v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2zm-4-1V5h-1V4h-2V3H8v1H6v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v-2h1V8h-1V6zm-3 10v-1h-1v-1h-1v-1h-1v-1h-1v-1H9v-1H8V9H7V8H6V7H5V6H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h-1v-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Card(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 3h18v1h1v14h-1v1H2v-1H1V4h1zm1 2v12h16V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardText(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 8v2H5V8zM5 12h10v2H5zM2 3h18v1h1v14h-1v1H2v-1H1V4h1zm1 2v12h16V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cart(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 14v2H6v-1H5v-4H4V8H3V3H1V1h4v3h16v4h-1v3h-1v1H7v2zM5 7h1v3h12V7h1V6H5zm2 10h2v1h1v2H9v1H7v-1H6v-2h1zm8 0h2v1h1v2h-1v1h-2v-1h-1v-2h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cast(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 19H2v-3h2v1h1m3 2H6v-3H5v-1H2v-2h3v1h2v2h1m3 3H9v-4H8v-2H6v-1H2v-2h4v1h2v1h1v1h1v2h1m8 4h-6v-2h5V5H4v3H2V4h1V3h16v1h1v14h-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Castle(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 19h-7v-6h-2v6H3v-6H2V4h5v2h1V4h6v2h1V4h5v9h-1m-2 4v-5h1V6h-1v2h-5V6h-2v2H5V6H4v6h1v5h3v-5h1v-1h4v1h1v5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartBar(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2h2v16h16v2H2zm4 14V8h4v8zm5 0V4h4v12zm5 0v-5h4v5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chat(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 20H1v-2h1v-1h1v-1h1v-2H3V8h1V7h1V6h1V5h2V4h8v1h2v1h1v1h1v1h1v6h-1v1h-1v1h-1v1h-2v1H8v1H6m0-1v-1h1v-1h8v-1h2v-1h1v-1h1V9h-1V8h-1V7h-2V6H9v1H7v1H6v1H5v4h1v3H5v1H4v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatProcessing(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 12H7v-2h2m4 2h-2v-2h2m4 2h-2v-2h2M6 20H1v-2h1v-1h1v-1h1v-2H3V8h1V7h1V6h1V5h2V4h8v1h2v1h1v1h1v1h1v6h-1v1h-1v1h-1v1h-2v1H8v1H6m0-1v-1h1v-1h8v-1h2v-1h1v-1h1V9h-1V8h-1V7h-2V6H9v1H7v1H6v1H5v4h1v3H5v1H4v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 11h2v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1V9h1V8h1V7h1V6h2v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H8v-1H7v-1H6v-1H5v-1H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckboxBlank(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 4h1V3h14v1h1v14h-1v1H4v-1H3zm2 13h12V5H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckboxCross(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 12h1v1h1v2h-2v-1h-1v-1h-2v1H9v1H7v-2h1v-1h1v-2H8V9H7V7h2v1h1v1h2V8h1V7h2v2h-1v1h-1zm5 7H4v-1H3V4h1V3h14v1h1v14h-1zM5 5v12h12V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckboxIntermediate(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 15H7V7h8m3 12H4v-1H3V4h1V3h14v1h1v14h-1m-1-1V5H5v12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckboxIntermediateVariant(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 15H7V7h8v1h-1v1h-1v1h-1v1h-1v1h-1v1H9v1H8m10 5H4v-1H3V4h1V3h14v1h1v14h-1m-1-1V5H5v12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckboxMarked(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 4h1V3h14v1h-1v1H5v12h12v-6h1v-1h1v8h-1v1H4v-1H3zm3 5h2v1h1v1h1v1h2v-1h1v-1h1V9h1V8h1V7h1V6h1V5h1V4h2v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckerLarge(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 22h-6v-6h-6v6H5v-6H0v-6h5V5H0V0h5v5h5V0h6v5h6v5h-6v6h6m-6-6V5h-6v5m0 6v-6H5v6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckerMedium(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 22h-4v-3h-3v3h-4v-3H7v3H4v-3H0v-4h4v-4H0V8h4V4H0V0h4v4h3V0h4v4h4V0h3v4h4v4h-4v3h4v4h-4v4h4M11 8V4H7v4m11 0V4h-3v4m-8 3V8H4v3m11 0V8h-4v3m0 4v-4H7v4m11 0v-4h-3v4m-8 4v-4H4v4m11 0v-4h-4v4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckerSmall(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2h1V1H2Zm2 0h1V1H4Zm2 0h1V1H6Zm2 0h1V1H8Zm2 0h1V1h-1Zm2 0h1V1h-1Zm2 0h1V1h-1Zm2 0h1V1h-1Zm2 0h1V1h-1Zm2 0h1V1h-1ZM1 3h1V2H1Zm2 0h1V2H3Zm2 0h1V2H5Zm2 0h1V2H7Zm2 0h1V2H9Zm2 0h1V2h-1Zm2 0h1V2h-1Zm2 0h1V2h-1Zm2 0h1V2h-1Zm2 0h1V2h-1ZM2 4h1V3H2Zm2 0h1V3H4Zm2 0h1V3H6Zm2 0h1V3H8Zm2 0h1V3h-1Zm2 0h1V3h-1Zm2 0h1V3h-1Zm2 0h1V3h-1Zm2 0h1V3h-1Zm2 0h1V3h-1ZM1 5h1V4H1Zm2 0h1V4H3Zm2 0h1V4H5Zm2 0h1V4H7Zm2 0h1V4H9Zm2 0h1V4h-1Zm2 0h1V4h-1Zm2 0h1V4h-1Zm2 0h1V4h-1Zm2 0h1V4h-1ZM2 6h1V5H2Zm2 0h1V5H4Zm2 0h1V5H6Zm2 0h1V5H8Zm2 0h1V5h-1Zm2 0h1V5h-1Zm2 0h1V5h-1Zm2 0h1V5h-1Zm2 0h1V5h-1Zm2 0h1V5h-1ZM1 7h1V6H1Zm2 0h1V6H3Zm2 0h1V6H5Zm2 0h1V6H7Zm2 0h1V6H9Zm2 0h1V6h-1Zm2 0h1V6h-1Zm2 0h1V6h-1Zm2 0h1V6h-1Zm2 0h1V6h-1ZM2 8h1V7H2Zm2 0h1V7H4Zm2 0h1V7H6Zm2 0h1V7H8Zm2 0h1V7h-1Zm2 0h1V7h-1Zm2 0h1V7h-1Zm2 0h1V7h-1Zm2 0h1V7h-1Zm2 0h1V7h-1ZM1 9h1V8H1Zm2 0h1V8H3Zm2 0h1V8H5Zm2 0h1V8H7Zm2 0h1V8H9Zm2 0h1V8h-1Zm2 0h1V8h-1Zm2 0h1V8h-1Zm2 0h1V8h-1Zm2 0h1V8h-1ZM2 10h1V9H2Zm2 0h1V9H4Zm2 0h1V9H6Zm2 0h1V9H8Zm2 0h1V9h-1Zm2 0h1V9h-1Zm2 0h1V9h-1Zm2 0h1V9h-1Zm2 0h1V9h-1Zm2 0h1V9h-1ZM1 11h1v-1H1Zm2 0h1v-1H3Zm2 0h1v-1H5Zm2 0h1v-1H7Zm2 0h1v-1H9Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1ZM2 12h1v-1H2Zm2 0h1v-1H4Zm2 0h1v-1H6Zm2 0h1v-1H8Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1ZM1 13h1v-1H1Zm2 0h1v-1H3Zm2 0h1v-1H5Zm2 0h1v-1H7Zm2 0h1v-1H9Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1ZM2 14h1v-1H2Zm2 0h1v-1H4Zm2 0h1v-1H6Zm2 0h1v-1H8Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1ZM1 15h1v-1H1Zm2 0h1v-1H3Zm2 0h1v-1H5Zm2 0h1v-1H7Zm2 0h1v-1H9Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1ZM2 16h1v-1H2Zm2 0h1v-1H4Zm2 0h1v-1H6Zm2 0h1v-1H8Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1ZM1 17h1v-1H1Zm2 0h1v-1H3Zm2 0h1v-1H5Zm2 0h1v-1H7Zm2 0h1v-1H9Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1ZM2 18h1v-1H2Zm2 0h1v-1H4Zm2 0h1v-1H6Zm2 0h1v-1H8Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1ZM1 19h1v-1H1Zm2 0h1v-1H3Zm2 0h1v-1H5Zm2 0h1v-1H7Zm2 0h1v-1H9Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1ZM2 20h1v-1H2Zm2 0h1v-1H4Zm2 0h1v-1H6Zm2 0h1v-1H8Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1ZM1 21h1v-1H1Zm2 0h1v-1H3Zm2 0h1v-1H5Zm2 0h1v-1H7Zm2 0h1v-1H9Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm2 0h1v-1h-1Zm3 1h-1v-1h-1v1h-1v-1h-1v1h-1v-1h-1v1h-1v-1h-1v1h-1v-1h-1v1h-1v-1h-1v1H9v-1H8v1H7v-1H6v1H5v-1H4v1H3v-1H2v1H1v-1H0v-1h1v-1H0v-1h1v-1H0v-1h1v-1H0v-1h1v-1H0v-1h1v-1H0v-1h1V9H0V8h1V7H0V6h1V5H0V4h1V3H0V2h1V1H0V0h1v1h1V0h1v1h1V0h1v1h1V0h1v1h1V0h1v1h1V0h1v1h1V0h1v1h1V0h1v1h1V0h1v1h1V0h1v1h1V0h1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Checkerboard(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 20H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1M11 7V4H7v3m11 0V5h-1V4h-2v3m-8 4V7H4v4m11 0V7h-4v4m0 4v-4H7v4m11 0v-4h-3v4m-8 3v-3H4v2h1v1m10 0v-3h-4v3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 10h1V9h1V7h-2v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8V9H7V8H6V7H4v2h1v1h1v1h1v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownCircle(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 9v2h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6V9h2v1h1v1h1v1h2v-1h1v-1h1V9zm-1-8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 16v1h1v1h2v-2h-1v-1h-1v-1h-1v-1h-1v-1h-1v-2h1V9h1V8h1V7h1V6h1V4h-2v1h-1v1h-1v1h-1v1H9v1H8v1H7v2h1v1h1v1h1v1h1v1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeftCircle(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 16h-2v-1h-1v-1H9v-1H8v-1H7v-2h1V9h1V8h1V7h1V6h2v2h-1v1h-1v1h-1v2h1v1h1v1h1zm2-15v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 6V5H9V4H7v2h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1H9v1H8v1H7v2h2v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRightCircle(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 16h2v-1h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6H9v2h1v1h1v1h1v2h-1v1h-1v1H9zm6-15v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUp(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 12H5v1H4v2h2v-1h1v-1h1v-1h1v-1h1v-1h2v1h1v1h1v1h1v1h1v1h2v-2h-1v-1h-1v-1h-1v-1h-1V9h-1V8h-1V7h-2v1H9v1H8v1H7v1H6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpCircle(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 13v-2h1v-1h1V9h1V8h1V7h2v1h1v1h1v1h1v1h1v2h-2v-1h-1v-1h-1v-1h-2v1H9v1H8v1zm9-12v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circle(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clipboard(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 5h1V4h4V2h2V1h4v1h2v2h4v1h1v15h-1v1H3v-1H2zm8-2v2h2V3zm8 3h-2v2H6V6H4v13h14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 5h2v6h1v1h1v1h1v2h-2v-1h-1v-1h-1v-1h-1zm5-4v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coffee(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 20v-2h16v2zM2 3h17v1h1v6h-1v1h-3v3h-1v1h-1v1H4v-1H3v-1H2zm14 2v4h2V5zM4 5v8h1v1h8v-1h1V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoinCopper(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 14h-2v-1H9v-1H8v-2h1V9h1V8h2v1h1v1h1v2h-1v1h-1Zm3 5H7v-1H6v-1H5v-1H4v-1H3V7h1V6h1V5h1V4h1V3h8v1h1v1h1v1h1v1h1v8h-1v1h-1v1h-1v1h-1Zm-3-7v-2h-2v2Zm2 5v-1h1v-1h1v-1h1V8h-1V7h-1V6h-1V5H8v1H7v1H6v1H5v6h1v1h1v1h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoinElectrum(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 19H5v-1H4v-2H3v-2H2V8h1V6h1V4h1V3h12v1h1v2h1v2h1v6h-1v2h-1v2h-1Zm-5-3v-2h-2v2Zm4 1v-2h1v-3h1v-2h-1V7h-1V5H6v2H5v3H4v2h1v3h1v2h3v-1H8v-2h1v-1h1v-1h2v1h1v1h1v2h-1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoinGold(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 19H3v-2h1v-2h1v-3h1v-2H5V7H4V5H3V3h16v2h-1v2h-1v3h-1v2h1v3h1v2h1Zm-3-2v-1h-1v-2h-1V8h1V6h1V5H6v1h1v2h1v6H7v2H6v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoinPlatinum(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 21H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-1H1V8h1V7h1V6h1V5h1V4h1V3h1V2h1V1h6v1h1v1h1v1h1v1h1v1h1v1h1v1h1v6h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1Zm-1-2v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1V9h-1V8h-1V7h-1V6h-1V5h-1V4h-1V3H9v1H8v1H7v1H6v1H5v1H4v1H3v4h1v1h1v1h1v1h1v1h1v1h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoinSilver(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 11h1V9h-1V7h-1V5h-2v2H9v2H8v2h1v-1h1V9h2v1h1Zm6 8H3v-1H2v-2h1v-2h1v-2h1v-2h1V8h1V6h1V4h1V3h1V2h2v1h1v1h1v2h1v2h1v2h1v2h1v2h1v2h1v2h-1Zm-7-6v-2h-2v2Zm5 4v-2h-1v-2h-1v-2h-1v2h-1v1h-1v1h-2v-1H9v-1H8v-2H7v2H6v2H5v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comment(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2h18v1h1v14h-1v1h-8v1h-1v1h-1v1H6v-3H2v-1H1V3h1zm1 2v12h5v3h1v-1h1v-1h1v-1h8V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentText(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2h18v1h1v14h-1v1h-8v1h-1v1h-1v1H6v-3H2v-1H1V3h1zm1 2v12h5v3h1v-1h1v-1h1v-1h8V4zm2 3h12v2H5zm0 4h10v2H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentUser(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 10h-2V9H9V7h1V6h2v1h1v2h-1m3 5H7v-2h1v-1h6v1h1m-7 7h1v-1h1v-1h1v-1h8V4H3v12h5m2 5H6v-3H2v-1H1V3h1V2h18v1h1v14h-1v1h-8v1h-1v1h-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compass(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1zm-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2zm-4 6V8h2V7h2V6h2v2h-1v2h-1v2h-1v1h-1v1h-2v1H8v1H6v-2h1v-2h1v-2h1V9zm2 1h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompassEastArrow(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 10v2H5v1H4v1H3v1H2v1H0v-2h1v-1h1v-1h1v-2H2V9H1V8H0V6h2v1h1v1h1v1h1v1m5-4h6v2h-4v2h4v2h-4v2h4v2h-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompassNorthArrow(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 4h2v1.5h1V8h1V4h2v10h-2v-2h-1v-2h-1v4H8m2 3v-1h2v1h1v1h1v1h1v1h1v2h-2v-1h-1v-1h-1v-1h-2v1H9v1H8v1H6v-2h1v-1h1v-1h1v-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompassNorthEast(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 6h2v1.5h1V10h1V6h2v10H8v-2H7v-2H6v4H4m8-10h6v2h-4v2h4v2h-4v2h4v2h-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompassNorthWest(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 6h2v1.5h1V10h1V6h2v10H6v-2H5v-2H4v4H2m8-10h2v6h1v1h1V8h2v5h1v-1h1V6h2v7h-1v2h-1v1h-2v-1h-2v1h-2v-1h-1v-2h-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompassSouthArrow(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 6h2V5h1V4h1V3h1V2h1V0h-2v1h-1v1h-1v1h-2V2H9V1H8V0H6v2h1v1h1v1h1v1h1M9 8h5v2h-4v2h3v1h1v4h-1v1H8v-2h4v-2H9v-1H8V9h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompassSouthEast(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 6h5v2H6v2h3v1h1v4H9v1H4v-2h4v-2H5v-1H4V7h1m7-1h6v2h-4v2h4v2h-4v2h4v2h-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompassSouthWest(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 6h5v2H4v2h3v1h1v4H7v1H2v-2h4v-2H3v-1H2V7h1m7-1h2v6h1v1h1V8h2v5h1v-1h1V6h2v7h-1v2h-1v1h-2v-1h-2v1h-2v-1h-1v-2h-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompassWestArrow(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 6h2v6h1v1h1V8h2v5h1v-1h1V6h2v7h-1v2h-1v1h-2v-1H8v1H6v-1H5v-2H4m12-3v2h1v1h1v1h1v1h1v1h2v-2h-1v-1h-1v-1h-1v-2h1V9h1V8h1V6h-2v1h-1v1h-1v1h-1v1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCard(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 4h18v1h1v12h-1v1H2v-1H1V5h1zm1 2v2h16V6zm0 10h16v-5H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 17h18v2H2zM4 6v1h1v1h1V7h1V6h1V5h1V4h1V3h2v1h1v1h1v1h1v1h1v1h1V7h1V6h1V5h1v11H2V5h1v1zm3 8h11v-4h-3V9h-1V8h-1V7h-1V6h-2v1H9v1H8v1H7v1H4v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cube(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21h-2v-1H8v-1H6v-1H4v-1H2V5h2V4h2V3h2V2h2V1h2v1h2v1h2v1h2v1h2v12h-2v1h-2v1h-2v1h-2m0-10V9h2V8h2V7h2V6h-2V5h-2V4h-2V3h-2v1H8v1H6v1H4v1h2v1h2v1h2v1m0 8v-7H8v-1H6V9H4v7h2v1h2v1m6 0v-1h2v-1h2V9h-2v1h-2v1h-2v7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CubeUnfolded(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3v5h10v7h-5v5h-7v-5H0V8h5V3zm-2 2H7v3h3zm-3 5v3h3v-3zm-2 0H2v3h3zm12 0v3h3v-3zm-2 5h-3v3h3zm-3-5v3h3v-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dagger(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 13h1v-1h1v-1h1v-1h1V9h1V8h1V7h1V5h-2v1h-1v1h-1v1h-1v1h-1v1h-1v1H9v1h1Zm-4 6H4v-1H3v-2h1v-1h1v-1h1v-1H5v-1H4v-2h1V9h2v1h1V9h1V8h1V7h1V6h1V5h1V4h1V3h5v5h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h1v2h-1v1h-2v-1H9v-1H8v1H7v1H6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Database(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 2h8v1h2v1h1v1h1v12h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3V5h1V4h1V3h2zm1 14v-1H6v-1H5v2h1v1h2v1h6v-1h2v-1h1v-2h-1v1h-2v1zm0-5v-1H6V9H5v3h2v1h2v1h4v-1h2v-1h2V9h-1v1h-2v1zm1-3v1h4V8h2V7h2V6h-1V5h-2V4H8v1H6v1H5v1h2v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Device(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 1h18v1h.94v18H20v1H2v-1h-.94V2H2zm1 2v16h16V3zm1 1h14v8H4zm1 10h3v3H5zm7 1h2v2h-2zm3-1h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Diamond(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 2h10v1h1v1h1v1h1v1h1v6h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2V6h1V5h1V4h1V3h1zm9 3V4h-1v2h1v1h2V6h-1V5zm-3 1V4h-2v2H9v1h4V6zM8 6V4H7v1H6v1H5v1h2V6zm-4 5h1v1h1v1h1v1h1v-2H7V9H4zm6 1v4h2v-4h1V9H9v3zm4 0v2h1v-1h1v-1h1v-1h1V9h-3v3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Division(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 8h-2V7H9V5h1V4h2v1h1v2h-1Zm5 4H5v-2h12Zm-5 6h-2v-1H9v-2h1v-1h2v1h1v2h-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Door(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 10h2v2h-2zm4-8v1h1v15h2v2H3v-2h2V3h1V2zm-1 2H7v14h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoorBox(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 14h-2v-2h2Zm3 4h1v-1h1V5h-1V4H5v1H4v12h1v1h1V6h10Zm2 2H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1Zm-4-2V8H8v10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoorOpen(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 10v2H9v-2zM6 2h10v1h1v15h2v2H3v-2h2V3h1zm1 2v14h4V4zm6 0v1h1V4zm1 1v1h1V5zm0 1h-1v1h1zm0 1v1h1V7zm0 1h-1v1h1zm0 1v1h1V9zm0 1h-1v1h1zm0 1v1h1v-1zm0 1h-1v1h1zm0 1v1h1v-1zm0 1h-1v1h1zm0 1v1h1v-1zm0 1h-1v1h1zm0 1v1h1v-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotHexagon(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 13h-2v-1H9v-2h1V9h2v1h1v2h-1m0 9h-2v-1H8v-1H6v-1H4v-1H2V5h2V4h2V3h2V2h2V1h2v1h2v1h2v1h2v1h2v12h-2v1h-2v1h-2v1h-2m0-1v-1h2v-1h2v-1h2V6h-2V5h-2V4h-2V3h-2v1H8v1H6v1H4v10h2v1h2v1h2v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotOctagon(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 13h-2v-1H9v-2h1V9h2v1h1v2h-1m3 9H7v-1H6v-1H5v-1H4v-1H3v-1H2v-1H1V7h1V6h1V5h1V4h1V3h1V2h1V1h8v1h1v1h1v1h1v1h1v1h1v1h1v8h-1v1h-1v1h-1v1h-1v1h-1v1h-1m-1-1v-1h1v-1h1v-1h1v-1h1v-1h1V8h-1V7h-1V6h-1V5h-1V4h-1V3H8v1H7v1H6v1H5v1H4v1H3v6h1v1h1v1h1v1h1v1h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 17v2H4v-2zM14 2v6h4v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-1H5V9H4V8h4V2zm-2 2h-2v6H9v1h1v1h2v-1h1v-1h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Email(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 5h1V4h18v1h1v13h-1v1H2v-1H1zm2 12h16V9h-1v1h-2v1h-2v1h-2v1h-2v-1H8v-1H6v-1H4V9H3zM19 6H3v1h2v1h2v1h2v1h4V9h2V8h2V7h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 14H9v-1H8V9h1V8h4v1h1v4h-1m2 4H7v-1H5v-1H3v-1H2v-1H1V9h1V8h1V7h2V6h2V5h8v1h2v1h2v1h1v1h1v4h-1v1h-1v1h-2v1h-2m-3-4v-2h-2v2m5 3v-1h2v-1h1v-1h1v-2h-1V9h-1V8h-2V7H7v1H5v1H4v1H3v2h1v1h1v1h2v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeFill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 14h-2v-1H9v-1H8v-2h1V9h1V8h2v1h1v1h1v2h-1v1h-1m3 4H7v-1H5v-1H3v-1H2v-1H1V9h1V8h1V7h2V6h2V5h8v1h2v1h2v1h1v1h1v4h-1v1h-1v1h-2v1h-2m-2-1v-1h1v-1h1V9h-1V8h-1V7H9v1H8v1H7v4h1v1h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func File(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 1v1h1v1h1v1h1v1h1v1h1v1h1v13h-1v1H4v-1H3V2h1V1zm0 3h-1v4h4V7h-1V6h-1V5h-1zM5 3v16h12v-9h-6V9h-1V3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 22H0V0h22Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filter(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 17h1v-7h1V9h1V8h1V7h1V5H6v2h1v1h1v1h1v1h1v6h1m3 4h-3v-1h-1v-1H9v-1H8v-6H7v-1H6V9H5V8H4V4h1V3h12v1h1v4h-1v1h-1v1h-1v1h-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fire(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 20H7v-1H6v-1H5v-1H4v-5h1v-2h1V9h1V8h1v1h1v2h1V9h1V5h-1V4H9V3H8V2h3v1h2v1h1v1h1v1h1v1h1v2h1v7h-1v2h-1v1h-2m-2-1v-1h2v-1h1v-2h1v-4h-1V8h-1V7h-1v4h-1v2h-1v1h-1v1H9v-1H8v-3H7v1H6v4h1v1h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flask(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 12h2v2h-2zm3-11v1h1v3h-1v2h1v2h1v2h1v2h1v1h1v2h1v4h-1v1H3v-1H2v-4h1v-2h1v-1h1v-2h1V9h1V7h1V5H7V2h1V1zm-2 2h-2v5H9v2H8v2H7v1H6v2H5v2h1v-1h1v-1h1v-1h1v1h1v1h1v1h1v1h2v-1h1v-2h1v-1h-1v-2h-1v-2h-1V8h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlaskEmpty(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1h6v1h1v3h-1v2h1v2h1v2h1v2h1v1h1v2h1v4h-1v1H3v-1H2v-4h1v-2h1v-1h1v-2h1V9h1V7h1V5H7V2h1zm2 2v5H9v2H8v2H7v1H6v2H5v2H4v2h14v-2h-1v-2h-1v-1h-1v-2h-1v-2h-1V8h-1V3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlaskRoundBottom(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 11h2v2h-2zm2-10v2h1v5h1v1h1v1h1v1h1v6h-1v1h-1v1h-1v1h-1v1H8v-1H7v-1H6v-1H5v-1H4v-6h1v-1h1V9h1V8h1V3h1V1zm-1 4h-2v4H9v1H8v1H7v1H6v2h1v-1h2v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h-1v-1h-1v-1h-1V9h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlaskRoundBottomEmpty(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 1h4v2h1v5h1v1h1v1h1v1h1v6h-1v1h-1v1h-1v1h-1v1H8v-1H7v-1H6v-1H5v-1H4v-6h1v-1h1V9h1V8h1V3h1zm1 4v4H9v1H8v1H7v1H6v4h1v1h1v1h1v1h4v-1h1v-1h1v-1h1v-4h-1v-1h-1v-1h-1V9h-1V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FloppyDisk(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 3h1V2h13v1h1v1h1v1h1v1h1v13h-1v1H3v-1H2zm16 4h-1V6h-1V5h-1v4H6V4H4v14h2v-5h10v5h2zm-7-3v3h2V4zm3 14v-3H8v3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 3h7v1h1v1h10v1h1v12h-1v1H2v-1H1V4h1zm1 4v10h16V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpen(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 4h1V3h7v1h1v1h10v1h1v12h-1v1H2v-1H1zm2 5h16V7H9V6H8V5H3zm0 8h16v-6H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatAlignBottom(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 11h-2v-1H9V9H8V8H7V7H6V6h4V2h2v4h4v1h-1v1h-1v1h-1v1h-1m6 5H4v-2h14m-4 5H4v-2h10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatAlignCenter(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 6H4V4h14m-3 5H7V7h8m3 5H4v-2h14m-3 5H7v-2h8m3 5H4v-2h14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatAlignJustify(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 6H4V4h14m0 5H4V7h14m0 5H4v-2h14m0 5H4v-2h14m0 5H4v-2h14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatAlignLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 6H4V4h14m-4 5H4V7h10m4 5H4v-2h14m-4 5H4v-2h10m4 5H4v-2h14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatAlignRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 6H4V4h14m0 5H8V7h10m0 5H4v-2h14m0 5H8v-2h10m0 5H4v-2h14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatAlignTop(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 6H4V4h14m-4 5H4V7h10m-2 13h-2v-4H6v-1h1v-1h1v-1h1v-1h1v-1h2v1h1v1h1v1h1v1h1v1h-4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatBold(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 19H6v-2h1V5H6V3h7v1h1v1h1v1h1v4h-1v2h1v1h1v4h-1v1h-1m-3-8V9h1V7h-1V6h-2v4m3 6v-1h1v-1h-1v-1h-3v3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatFloatLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 6H4V4h14m-2 5h-6V7h6m-7 5H4V7h5m9 5h-8v-2h8m-2 5H4v-2h12m2 5H4v-2h14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatFloatRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 6H4V4h14m-2 5h-6V7h6m-7 5H4V7h5m9 5h-8v-2h8m-2 5H4v-2h12m2 5H4v-2h14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatHorizontalAlignCenter(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 15H5v-3H2v-2h3V7h1v1h1v1h1v1h1v2H8v1H7v1H6m11 1h-1v-1h-1v-1h-1v-1h-1v-2h1V9h1V8h1V7h1v3h3v2h-3m-5 6h-2V4h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatItalic(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 19H4v-2h3v-1h1v-2h1v-2h1v-2h1V8h1V5h-2V3h8v2h-3v1h-1v2h-1v2h-1v2h-1v2h-1v3h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatLineSpacing(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 7h-9V5h9m0 7h-9v-2h9m0 7h-9v-2h9M7 19H5v-1H4v-1H3v-1H2v-1h3V7H2V6h1V5h1V4h1V3h2v1h1v1h1v1h1v1H7v8h3v1H9v1H8v1H7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatVerticalAlignCenter(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 9h-2V8H9V7H8V6H7V5h3V2h2v3h3v1h-1v1h-1v1h-1m6 4H4v-2h14m-6 10h-2v-3H7v-1h1v-1h1v-1h1v-1h2v1h1v1h1v1h1v1h-3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gamepad(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1zm-1 2H9v6H3v4h6v6h4v-6h6V9h-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GamepadCenter(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 13h-2v-1H9v-2h1V9h2v1h1v2h-1m2 9H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6v1h1v5h5v1h1v6h-1v1h-5v5h-1m-1-1v-6h6V9h-6V3H9v6H3v4h6v6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GamepadCenterFill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 21H9v-1H8v-6H2v-1H1V9h1V8h6V2h1V1h4v1h1v6h6v1h1v4h-1v1h-6v6h-1m-1-7v-1h1v-2h-1V9h-2v1H9v2h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GamepadDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 18h-2v-2h2m2 5H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6v1h1v5h5v1h1v6h-1v1h-5v5h-1m-1-1v-6h6V9h-6V3H9v6H3v4h6v6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GamepadDownFill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 21H9v-1H8v-6H2v-1H1V9h1V8h6V2h1V1h4v1h1v6h6v1h1v4h-1v1h-6v6h-1m-1-1v-2h-2v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GamepadDownLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 12H4v-2h2m6 8h-2v-2h2m2 5H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6v1h1v5h5v1h1v6h-1v1h-5v5h-1m-1-1v-6h6V9h-6V3H9v6H3v4h6v6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GamepadDownLeftFill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 21H9v-1H8v-6H2v-1H1V9h1V8h6V2h1V1h4v1h1v6h6v1h1v4h-1v1h-6v6h-1m-8-8v-2H3v2m9 7v-2h-2v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GamepadDownRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 12h-2v-2h2m-6 8h-2v-2h2m2 5H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6v1h1v5h5v1h1v6h-1v1h-5v5h-1m-1-1v-6h6V9h-6V3H9v6H3v4h6v6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GamepadDownRightFill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 21H9v-1H8v-6H2v-1H1V9h1V8h6V2h1V1h4v1h1v6h6v1h1v4h-1v1h-6v6h-1m6-8v-2h-2v2m-5 7v-2h-2v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GamepadFill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 21H9v-1H8v-6H2v-1H1V9h1V8h6V2h1V1h4v1h1v6h6v1h1v4h-1v1h-6v6h-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GamepadLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 12H4v-2h2m8 11H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6v1h1v5h5v1h1v6h-1v1h-5v5h-1m-1-1v-6h6V9h-6V3H9v6H3v4h6v6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GamepadLeftFill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 21H9v-1H8v-6H2v-1H1V9h1V8h6V2h1V1h4v1h1v6h6v1h1v4h-1v1h-6v6h-1m-8-8v-2H3v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GamepadRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 12h-2v-2h2m-4 11H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6v1h1v5h5v1h1v6h-1v1h-5v5h-1m-1-1v-6h6V9h-6V3H9v6H3v4h6v6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GamepadRightFill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 21H9v-1H8v-6H2v-1H1V9h1V8h6V2h1V1h4v1h1v6h6v1h1v4h-1v1h-6v6h-1m6-8v-2h-2v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GamepadUp(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 6h-2V4h2m2 17H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6v1h1v5h5v1h1v6h-1v1h-5v5h-1m-1-1v-6h6V9h-6V3H9v6H3v4h6v6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GamepadUpFill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 21H9v-1H8v-6H2v-1H1V9h1V8h6V2h1V1h4v1h1v6h6v1h1v4h-1v1h-6v6h-1M12 5V3h-2v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GamepadUpLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 6h-2V4h2m-6 8H4v-2h2m8 11H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6v1h1v5h5v1h1v6h-1v1h-5v5h-1m-1-1v-6h6V9h-6V3H9v6H3v4h6v6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GamepadUpLeftFill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 21H9v-1H8v-6H2v-1H1V9h1V8h6V2h1V1h4v1h1v6h6v1h1v4h-1v1h-6v6h-1M12 5V3h-2v2m-5 7v-2H3v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GamepadUpRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 6h-2V4h2m6 8h-2v-2h2m-4 11H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6v1h1v5h5v1h1v6h-1v1h-5v5h-1m-1-1v-6h6V9h-6V3H9v6H3v4h6v6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GamepadUpRightFill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 21H9v-1H8v-6H2v-1H1V9h1V8h6V2h1V1h4v1h1v6h6v1h1v4h-1v1h-6v6h-1M12 5V3h-2v2m9 7v-2h-2v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Glasses(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 16h-6v-1h-1v-1h-2v1H9v1H3v-1H2v-1H1V8h1V7h18v1h1v6h-1v1h-1M8 14v-1h1v-1h1V9H3v4h1v1m14 0v-1h1V9h-7v3h1v1h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 20h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-2H1V5h1V4h1V3h1V2h5v1h1v1h2V3h1V2h5v1h1v1h1v1h1v5h-1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1zm-7-9v1h1v1h1v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1V9h1V6h-1V5h-1V4h-3v1h-1v1h-1v1h-2V6H9V5H8V4H5v1H4v1H3v3h1v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartBroken(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 14h1v-1h1v-1h1v-1h1V9h1V6h-1V5h-1V4h-3v5h-1v2h1Zm-6 1h1v-2H8v-1H7V8h1V6h1V5H8V4H5v1H4v1H3v3h1v2h1v1h1v1h1v1h1Zm6 3h-1v-1h-1v-5h-1V8h1V3h1V2h5v1h1v1h1v1h1v5h-1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1Zm-2 2h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-2H1V5h1V4h1V3h1V2h5v1h1v1h1v3h-1v2H9v2h1v2h1v5h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Help(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 15h-2v-3h1v-1h1v-1h1V9h1V6h-1V5H9v1H8v3H6V5h1V4h1V3h6v1h1v1h1v5h-1v1h-1v1h-1v1h-1m0 6h-2v-2h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hexagon(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21h-2v-1H8v-1H6v-1H4v-1H2V5h2V4h2V3h2V2h2V1h2v1h2v1h2v1h2v1h2v12h-2v1h-2v1h-2v1h-2m0-1v-1h2v-1h2v-1h2V6h-2V5h-2V4h-2V3h-2v1H8v1H6v1H4v10h2v1h2v1h2v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Image(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 4h1V3h18v1h1v14h-1v1H2v-1H1zm2 10h1v-1h1v-1h1v-1h1v-1h1V9h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h2V5H3zm11 3v-1h-1v-1h-1v-1h-1v-1h-1v-1H8v1H7v1H6v1H5v1H4v1zm-1-9h1V7h2v1h1v2h-1v1h-2v-1h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Key(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 15H5v-1H4v-1H3V9h1V8h1V7h4v1h1v1h1v1h8v2h-2v2h-2v-2h-4v1h-1v1H9m-1-1v-1h1v-2H8V9H6v1H5v2h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Label(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 4h13v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H2v-1H1V5h1zm14 9h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6H3v10h11v-1h1v-1h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LabelVariant(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 4v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H2v-2h1v-1h1v-1h1v-1h1v-1h1v-2H6V9H5V8H4V7H3V6H2V4zm-1 12v-1h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6H6v1h1v1h1v1h1v1h1v2H9v1H8v1H7v1H6v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Led(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 21v-6H4v-2h2V6h1V4h1V3h1V2h4v1h1v1h1v2h1v7h2v2h-4v6h-2v-6h-2v6zm4-16V4h-2v1H9v2H8v6h1v-1h4v1h1V7h-1V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lightbulb(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 19h6v2H8zm0-1v-4H7v-1H6v-1H5v-1H4V5h1V4h1V3h1V2h1V1h6v1h1v1h1v1h1v1h1v6h-1v1h-1v1h-1v1h-1v4zm5-6h1v-1h1v-1h1V6h-1V5h-1V4h-1V3H9v1H8v1H7v1H6v4h1v1h1v1h1v1h1v3h2v-3h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 12h2v1h1v2h-1v1h-2v-1H9v-2h1zM8 2h6v1h1v1h1v4h1v1h1v10h-1v1H5v-1H4V9h1V8h1V4h1V3h1zm1 2v1H8v3h6V5h-1V4zm7 6H6v8h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOpen(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 13h2v1h1v2h-1v1h-2v-1H9v-2h1zm4-11v1h1v1h1v5h1v1h1v10h-1v1H5v-1H4V10h1V9h9V5h-1V4H9v1H8v2H6V4h1V3h1V2zm2 9H6v8h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Login(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 1h12v1h1v18h-1v1H5v-1H4v-6h2v5h10V3H6v5H4V2h1zm3 5h2v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1H8v-2h1v-1h1v-1H2v-2h8V9H9V8H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Logout(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 1v1h1v4h-1V5h-1V3H6v16h10v-2h1v-1h1v4h-1v1H5v-1H4V2h1V1zm-4 5h2v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-2v-2h1v-1h1v-1H8v-2h7V9h-1V8h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifyMinus(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 10H6V8h6m7 12h-2v-1h-1v-1h-1v-1h-1v-2h-2v1H6v-1H5v-1H4v-1H3v-1H2V6h1V5h1V4h1V3h1V2h6v1h1v1h1v1h1v1h1v6h-1v2h2v1h1v1h1v1h1v2h-1m-8-5v-1h2v-2h1V7h-1V5h-2V4H7v1H5v2H4v4h1v2h2v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifyPlus(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 12H8v-2H6V8h2V6h2v2h2v2h-2m9 10h-2v-1h-1v-1h-1v-1h-1v-2h-2v1H6v-1H5v-1H4v-1H3v-1H2V6h1V5h1V4h1V3h1V2h6v1h1v1h1v1h1v1h1v6h-1v2h2v1h1v1h1v1h1v2h-1m-8-5v-1h2v-2h1V7h-1V5h-2V4H7v1H5v2H4v4h1v2h2v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Map(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 4h2V3h2V2h4v1h3v1h2V3h2V2h3v16h-2v1h-2v1h-4v-1H9v-1H7v1H5v1H2zm2 2v11h2v-1h1V5H5v1zm8-1H9v11h1v1h3V6h-1zm4 1h-1v11h2v-1h1V5h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuBottomLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 4h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2H7zm2 4v5h5v-1h-1v-1h-1v-1h-1V9h-1V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuBottomRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 15v-2h1v-1h1v-1h1v-1h1V9h1V8h1V7h1V6h1V5h1V4h2v11zm4-2h5V8h-1v1h-1v1h-1v1H9v1H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 8h14v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4zm4 2v1h1v1h1v1h2v-1h1v-1h1v-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuDownFill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 9V8H5v1h1v1h1v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1V9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 4v14h-2v-1h-1v-1h-1v-1h-1v-1H9v-1H8v-1H7v-2h1V9h1V8h1V7h1V6h1V5h1V4zm-2 4h-1v1h-1v1h-1v2h1v1h1v1h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuLeftFill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 17h1V5h-1v1h-1v1h-1v1h-1v1H9v1H8v2h1v1h1v1h1v1h1v1h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuLeftRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 4h2v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2zm-2 0v14H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-2h1V9h1V8h1V7h1V6h1V5h1V4zm4 4v6h1v-1h1v-1h1v-2h-1V9h-1V8zM8 8H7v1H6v1H5v2h1v1h1v1h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 18V4h2v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1H9v1zm2-4h1v-1h1v-1h1v-2h-1V9h-1V8H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuRightFill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 5H8v12h1v-1h1v-1h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6H9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuTopLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 7v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H9v1H7V7zm-4 2H9v5h1v-1h1v-1h1v-1h1v-1h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuTopRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 18h-2v-1h-1v-1h-1v-1h-1v-1H9v-1H8v-1H7v-1H6v-1H5V9H4V7h11zm-2-4V9H8v1h1v1h1v1h1v1h1v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuUp(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 14H4v-2h1v-1h1v-1h1V9h1V8h1V7h1V6h2v1h1v1h1v1h1v1h1v1h1v1h1zm-4-2v-1h-1v-1h-1V9h-2v1H9v1H8v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuUpDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 10V8h1V7h1V6h1V5h1V4h1V3h1V2h2v1h1v1h1v1h1v1h1v1h1v1h1v2zm0 2h14v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4zm4-4h6V7h-1V6h-1V5h-2v1H9v1H8zm0 6v1h1v1h1v1h2v-1h1v-1h1v-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuUpFill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 13v1h12v-1h-1v-1h-1v-1h-1v-1h-1V9h-1V8h-2v1H9v1H8v1H7v1H6v1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Message(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 1h18v1h1v14h-1v1H5v1H4v1H3v1H2v1H1V2h1zm1 2v13h1v-1h15V3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageProcessing(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 1h18v1h1v14h-1v1H5v1H4v1H3v1H2v1H1V2h1zm2 14h15V3H3v13h1zm2-7h2v2H6zm4 0h2v2h-2zm4 0h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageText(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 1h18v1h1v14h-1v1H5v1H4v1H3v1H2v1H1V2h1zm1 2v13h1v-1h15V3zm2 3h12v2H5zm0 4h9v2H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageUser(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 9h-2V8H9V6h1V5h2v1h1v2h-1m3 5H7v-2h1v-1h6v1h1M3 16h1v-1h15V3H3M2 21H1V2h1V1h18v1h1v14h-1v1H5v1H4v1H3v1H2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microphone(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 21v-3H8v-1H6v-1H5v-1H4v-2H3V9h2v3h1v2h1v1h2v1h4v-1h2v-1h1v-2h1V9h2v4h-1v2h-1v1h-1v1h-2v1h-2v3zm-2-8v-1H7V3h1V2h1V1h4v1h1v1h1v9h-1v1h-1v1H9v-1zm1-2h1v1h2v-1h1V4h-1V3h-2v1H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 12H5v-2h12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusBox(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 12H6v-2h10Zm2 8H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1Zm-1-2v-1h1V5h-1V4H5v1H4v12h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusBoxFill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 20H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1Zm-2-8v-2H6v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircle(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 7v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2zm-4-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6zm-1 4v2H6v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircleFill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 21H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2m1-8v-2H6v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Monitor(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2h18v1h1v12h-1v1h-7v2h2v2H7v-2h2v-2H2v-1H1V3h1zm1 2v10h16V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MonitorImage(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 6h2v1h1v2h-1v1h-2V9h-1V7h1zM2 2h18v1h1v12h-1v1h-7v2h2v2H7v-2h2v-2H2v-1H1V3h1zm1 2v5h1V8h1V7h1V6h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h4V4zm7 7H9v-1H8V9H6v1H5v1H4v1H3v2h9v-1h-1v-1h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Multiply(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 16h-2v-1h-1v-1h-1v-1h-2v1H9v1H8v1H6v-2h1v-1h1v-1h1v-2H8V9H7V8H6V6h2v1h1v1h1v1h2V8h1V7h1V6h2v2h-1v1h-1v1h-1v2h1v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicNote(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 2h7v5h-5v11h-1v1h-1v1H7v-1H6v-1H5v-4h1v-1h1v-1h4zm0 13h-1v-1H8v1H7v2h1v1h2v-1h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Necklace(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 17h1v-1h2v1h1v2h-1v1h-2v-1H9zm1-2v-1H9v-1h-.91v-1H7v-2H6V8H5V6H4V3h1V2h12v1h1v3h-1v2h-1v2h-1v2h-1v1h-1v1h-1v1zM7 5v2h1v2h1v2h1.09v1H12v-1h1V9h1V7h1V5h1V4H6v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Note(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 3v1h1v1h1v1h1v1h1v1h1v1h1v9h-1v1H2v-1H1V4h1V3zm0 3h-1v4h4V9h-1V8h-1V7h-1zM3 5v12h16v-5h-6v-1h-1V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notebook(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2v18h-1v1H4v-1H3v-2H1v-2h2v-4H1v-2h2V6H1V4h2V2h1V1h14v1zm-5 7h-1V8h-1v1h-1v1h-1V3H7v16h10V3h-2v7h-1zM3 4v2h2V4zm2 6H3v2h2zm0 6H3v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notification(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 20v1h-2v-1H8v-2h6v2zM2 15h1v-1h1V6h1V4h1V3h2V2h2V1h2v1h2v1h2v1h1v2h1v8h1v1h1v2H2zm4 0h10V7h-1V5h-2V4H9v1H7v2H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Octagon(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 6h1V5h1V4h1V3h1V2h1V1h8v1h1v1h1v1h1v1h1v1h1v1h1v8h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-1H1V7h1zm13 11h1v-1h1v-1h1v-1h1V8h-1V7h-1V6h-1V5h-1V4h-1V3H8v1H7v1H6v1H5v1H4v1H3v6h1v1h1v1h1v1h1v1h1v1h6v-1h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OctagonAlert(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 6h1V5h1V4h1V3h1V2h1V1h8v1h1v1h1v1h1v1h1v1h1v1h1v8h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-1H1V7h1zm13 11h1v-1h1v-1h1v-1h1V8h-1V7h-1V6h-1V5h-1V4h-1V3H8v1H7v1H6v1H5v1H4v1H3v6h1v1h1v1h1v1h1v1h1v1h6v-1h1zM10 6h2v7h-2zm0 8h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paperclip(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 21H8v-1H7v-1H6V3h1V2h1V1h4v1h1v1h1v14h-1v1h-3v-1H9V5h2v11h1V4h-1V3H9v1H8v14h1v1h5v-1h1V5h2v14h-1v1h-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 4h3v14H6zm7 14V4h3v14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Peace(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 15h1v-1h1v-1h1v-1h1v-1h1v-1h1V3H8v1H6v1H5v1H4v2H3v6h1m13 1h1v-1h1V8h-1V6h-1V5h-1V4h-2V3h-2v7h1v1h1v1h1v1h1v1h1m-2 7H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2m-5-1v-6H9v1H8v1H7v1H6v2h2v1m6 0v-1h2v-2h-1v-1h-1v-1h-1v-1h-1v6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pencil(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 2h1v1h1v1h1v1h1v1h-1v1h-1v1h-1V7h-1V6h-1V5h-1V4h1V3h1m-4 3h2v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H9v1H8v1H7v1H6v1H2v-4h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1V9h1V8h1V7h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pickaxe(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 2h3v1h2v1h2v1h2v2h1v2h1v2h1v3h-2v-2h-1v-1h-1v-1h-2V9h-1V8h-1V6h-1V5h-1V4H8m3 5h1v1h1v1h-1v1h-1v1h-1v1H9v1H8v1H7v1H6v1H5v1H4v1H3v1H2v-1H1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pictogrammers(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 0h14v1h1v1h1v18h-1v1h-1v1H4v-1H3v-1H2V2h1V1h1zm0 18v1h1v1h12v-1h1v-1zM17 2H5v1H4v12h1v1h12v-1h1V3h-1zm-4 2v1h1v1h1v2h-1v1h-1v1h-3v4H8V4zm-1 2h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 5v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1H9v1H7V4h2v1zm2 5h-1V9h-1V8H9v6h1v-1h1v-1h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 17h-2v-5H5v-2h5V5h2v5h5v2h-5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusBox(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 16h-2v-4H6v-2h4V6h2v4h4v2h-4Zm6 4H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1Zm-1-2v-1h1V5h-1V4H5v1H4v12h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusBoxFill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 20H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1Zm-6-4v-4h4v-2h-4V6h-2v4H6v2h4v4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircle(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 7v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2zm-4-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6zm-7 0h2v4h4v2h-4v4h-2v-4H6v-2h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircleFill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 21H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2Zm-3-5v-4h4v-2h-4V6h-2v4H6v2h4v4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Poll(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 18H4V8h3m5 10H9V4h3m5 14h-3v-7h3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pound(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 19h-2v-4H8v4H6v-4H3v-2h4V9H4V7h4V3h2v4h4V3h2v4h3v2h-4v4h3v2h-4m-1-2V9H9v4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Radiobox(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 2h6v1h2v1h1v1h1v1h1v2h1v6h-1v2h-1v1h-1v1h-1v1h-2v1H8v-1H6v-1H5v-1H4v-1H3v-2H2V8h1V6h1V5h1V4h1V3h2zm1 2v1H7v1H6v1H5v2H4v4h1v2h1v1h1v1h2v1h4v-1h2v-1h1v-1h1v-2h1V9h-1V7h-1V6h-1V5h-2V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadioboxMarked(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 2h6v1h2v1h1v1h1v1h1v2h1v6h-1v2h-1v1h-1v1h-1v1h-2v1H8v-1H6v-1H5v-1H4v-1H3v-2H2V8h1V6h1V5h1V4h1V3h2zm1 2v1H7v1H6v1H5v2H4v4h1v2h1v1h1v1h2v1h4v-1h2v-1h1v-1h1v-2h1V9h-1V7h-1V6h-1V5h-2V4zm4 3v1h1v1h1v4h-1v1h-1v1H9v-1H8v-1H7V9h1V8h1V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RelativeScale(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 10H5V9h1V8h2v1H7m4 1H9V9h1V8h2v1h-1m4 1h-2V9h1V8h2v1h-1m0 5h-1v-2h1v-1h1v2h-1m4 6H3v-1H2V4h1V3h16v1h1v14h-1m-1-1V5H4v12h10v-1h1v-1h1v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RemoveCircle(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 7v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2zm-4-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6zm-3 1v1h1v1h-1v1h-1v2h1v1h1v1h-1v1h-1v-1h-1v-1h-2v1H9v1H8v-1H7v-1h1v-1h1v-2H8V9H7V8h1V7h1v1h1v1h2V8h1V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RotateClockwise(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 11v1h-1v1h-1v1h-1v1h-2v-1h-1v-1h-1v-1h-1v-1h3V9h-1V7h-1V6h-2V5H9v1H7v1H6v2H5v4h1v2h1v1h2v1h4v-1h3v2h-2v1H8v-1H6v-1H5v-1H4v-2H3V8h1V6h1V5h1V4h2V3h6v1h2v1h1v1h1v2h1v3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RotateCounterclockwise(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 11v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1H5V9h1V7h1V6h2V5h4v1h2v1h1v2h1v4h-1v2h-1v1h-2v1H9v-1H6v2h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Script(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 1H5v1H4v13h2V3h9v16h-2v-1h-1v-1H1v3h1v1h14v-1h1V3h2v2h2V2h-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 20h-1v-1h-1v-1h-1v-1h-1v-1h-1v-1h-1v-1h-2v1H6v-1H4v-1H3v-2H2V6h1V4h1V3h2V2h5v1h2v1h1v2h1v5h-1v2h1v1h1v1h1v1h1v1h1v1h1v1h-1m-8-6v-1h1v-1h1V6h-1V5h-1V4H6v1H5v1H4v5h1v1h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shield(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 4h2V3h2V2h2V1h4v1h2v1h2v1h2v10h-1v2h-1v2h-1v1h-1v1h-2v1H9v-1H7v-1H6v-1H5v-2H4v-2H3zm7-1v1H8v1H6v1H5v7h1v2h1v2h1v1h2v1h2v-1h2v-1h1v-2h1v-2h1V6h-1V5h-2V4h-2V3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Skull(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 2h2V1h6v1h2v1h1v1h1v1h1v2h1v7h-1v2h-1v4h-1v1H5v-1H4v-4H3v-2H2V8h1V5h1V4h1V3h1zm9 3V4h-2V3H9v1H7v1H6v1H5v3H4v4h1v2h1v4h2v-2h2v2h2v-2h2v2h2v-4h1v-2h1V8h-1V6h-1V5zM7 8h3v3H7zm5 3V8h3v3zm-2 2h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Speaker(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 1h14v1h1v18h-1v1H4v-1H3V2h1zm1 2v16h12V3zm4 2h1V4h2v1h1v2h-1v1h-2V7H9zm0 13v-1H8v-1H7v-4h1v-1h1v-1h4v1h1v1h1v4h-1v1h-1v1zm1-5H9v2h1v1h2v-1h1v-2h-1v-1h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stop(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 5v1h1v10h-1v1H6v-1H5V6h1V5zM7 7v8h8V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sword(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 20H3v-1H2v-2h1v-1h1v-1h1v-1H4v-1H3v-2h2v1h1v-1h1v-1h1V9h1V8h1V7h1V6h1V5h1V4h1V3h1V2h5v5h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h1v2H9v-1H8v-1H7v1H6v1H5m5-6v-1h1v-1h1v-1h1V9h1V8h1V7h1V6h1V5h1V4h-2v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H9v1H8v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopDoorHorizontal(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 15H4v-3H0v-2h4V7h14v3h4v2h-4m-2 1V9H6v4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopDoorOneWayDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 12H0v-2h4V6h14v4h4M12 20h-2v-1H9v-1H8v-1H7v-1h3v-3h2v3h3v1h-1v1h-1v1h-1m4-8V8H6v3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopDoorOneWayLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 15H5v-1H4v-1H3v-1H2v-2h1V9h1V8h1V7h1v3h3v2H6m6 10h-2V0h2v4h4v14h-4m2-2V6h-3v10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopDoorOneWayRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 15h-1v-3h-3v-2h3V7h1v1h1v1h1v1h1v2h-1v1h-1v1h-1m-5 8h-2v-4H6V4h4V0h2m-1 16V6H8v10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopDoorOneWayUp(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 9h-2V6H7V5h1V4h1V3h1V2h2v1h1v1h1v1h1v1h-3m6 10H4v-4H0v-2h22v2h-4m-2 2v-3H6v3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopDoorSecretHorizontal(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 14h-2v-4h-2v3h-1v1H7v-1H6V8h2v4h2V9h1V8h4v1h1m2 8H6v-1H4v-1H3v-3H0v-2h3V6h1V5h12v1h2v2h1v2h3v2h-3v4h-1m-1-1V8h-1V7H5v7h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopDoorSecretVertical(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 16H8v-2h4v-2H9v-1H8V7h1V6h5v2h-4v2h3v1h1v4h-1m-1 7h-2v-3H6v-1H5V6h1V4h1V3h3V0h2v3h4v1h1v12h-1v2h-2v1h-2m2-2v-1h1V5H8v1H7v11Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopDoorVertical(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22h-2v-4H7V4h3V0h2v4h3v14h-3m1-2V6H9v10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopHorizontalRotateClockwise(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 8h-1V5H9v1H7v1H6v1H4V6h1V5h1V4h2V3h3V0h1v1h1v1h1v1h1v2h-1v1h-1v1h-1M2 12H0v-2h2m7 2H3v-2h6m3 2h-2v-2h2m7 2h-6v-2h6m3 2h-2v-2h2m-8 9H8v-1H6v-1H5v-1H4v-2h2v1h1v1h2v1h4v-1h2v-1h1v-1h2v2h-1v1h-1v1h-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopHorizontalRotateCounterclockwise(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 8h-2V7h-1V6h-2V5h-2v3h-1V7H9V6H8V5H7V3h1V2h1V1h1V0h1v3h3v1h2v1h1v1h1M2 12H0v-2h2m7 2H3v-2h6m3 2h-2v-2h2m7 2h-6v-2h6m3 2h-2v-2h2m-8 9H8v-1H6v-1H5v-1H4v-2h2v1h1v1h2v1h4v-1h2v-1h1v-1h2v2h-1v1h-1v1h-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopHorizontalStairsAscendLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2h-2V0h2m0 17h-2V5h2M9 18H7V4h2M6 19H4V3h2M3 20H1V2h2m9 20h-2v-2h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopHorizontalStairsAscendRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2h-2V0h2m0 17h-2V5h2m3 13h-2V4h2m3 15h-2V3h2m3 17h-2V2h2m-9 20h-2v-2h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopHorizontalStairsDescendDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 12H0v-2h2m16 2H4v-2h14m4 2h-2v-2h2m-5 5H5v-2h12m-1 5H6v-2h10m-1 5H7v-2h8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopHorizontalStairsDescendLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2h-2V0h2M3 15H1V7h2m3 9H4V6h2m3 11H7V5h2m3 13h-2V4h2m0 18h-2v-2h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopHorizontalStairsDescendRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2h-2V0h2m9 15h-2V7h2m-3 9h-2V6h2m-3 11h-2V5h2m-3 13h-2V4h2m0 18h-2v-2h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopHorizontalStairsDescendUp(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 3H7V1h8m1 5H6V4h10m1 5H5V7h12M2 12H0v-2h2m16 2H4v-2h14m4 2h-2v-2h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopSpiralStairsDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3h1V2h-1m2 1h1V2h-1m2 1h1V2h-1m2 1h1V2h-1m-5 2h1V3h-1m2 1h1V3h-1m2 1h1V3h-1m-5 2h1V4h-1m2 1h1V4h-1m2 1h1V4h-1m-3 2h1V5h-1m2 1h1V5h-1m4 2h-1V6h1m-7 1h1V6h-1m2 1h1V6h-1m-1 2h1V7h-1m4 2h-1V8h1m2 1h-1V8h1m-7 1h1V8h-1M2 19h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1H2m17 7h1v-8h-8v1h1v1h1v1h1v1h1v1h1v1h1v1h1m3 4H0V0h2v2h8V0h12M10 5V3H2v2m8 3V6H2v2m18 2V3h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1m-3 1V9H2v2m9 9v-8h-1v1H9v1H8v1H7v1H6v1H5v1H4v1H3v1m16 0v-1h-1v-1h-1v-1h-1v-1h-1v-1h-1v-1h-1v-1h-1v7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopSpiralStairsLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 9h1V2H3v1h1v1h1v1h1v1h1v1h1v1h1m4 5h1v-1h-1m2 1h1v-1h-1m2 1h1v-1h-1m2 1h1v-1h-1m-5 2h1v-1h-1m2 1h1v-1h-1m2 1h1v-1h-1m-3 2h1v-1h-1m2 1h1v-1h-1m2 1h1v-1h-1m-5 3h-1v-1h1m2 0h1v-1h-1m2 1h1v-1h-1m-1 2h1v-1h-1m2 1h1v-1h-1m-5 3h-1v-1h1m2 1h-1v-1h1m2 0h1v-1h-1M2 19h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1H2m17 7h1v-1h-1m3 4H0V0h22v2h-2v8h2m-9 0V2h-2v8m5 0V2h-2v8m5 0V2h-2v8m-7 1v-1H9V9H8V8H7V7H6V6H5V5H4V4H3V3H2v8m9 9v-8h-1v1H9v1H8v1H7v1H6v1H5v1H4v1H3v1m16 0v-1h-1v-1h-1v-1h-1v-1h-1v-1h-1v-1h-1v-1h-1v7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopSpiralStairsRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 4H6V3h1m2 1H8V3h1M2 4h1V3H2m1 2h1V4H3m6 2H8V5h1M2 6h1V5H2m2 1h1V5H4M3 7h1V6H3m2 1h1V6H5M2 8h1V7H2m2 1h1V7H4m2 1h1V7H6M3 9h1V8H3m2 1h1V8H5m2 1h1V8H7m2 1h1V2H3v1h1v1h1v1h1v1h1v1h1v1h1m-7 2h1V9H2m2 1h1V9H4m2 1h1V9H6m2 1h1V9H8m3 1h1V9h1V8h1V7h1V6h1V5h1V4h1V3h1V2h-8m8 17h1v-8h-8v1h1v1h1v1h1v1h1v1h1v1h1v1h1m3 4H0v-2h2v-8H0V0h22m-2 10V3h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1M5 20v-8H3v8m5 0v-8H6v8m5 0v-8H9v8m10 0v-1h-1v-1h-1v-1h-1v-1h-1v-1h-1v-1h-1v-1h-1v7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopSpiralStairsRoundDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3h1V2h-1m1 2h1V3h-1m2 1h1V3h-1m-3 2h1V4h-1m2 1h1V4h-1m2 1h1V4h-1m-3 2h1V5h-1m2 1h1V5h-1m-3 2h1V6h-1m2 1h1V6h-1m-1 2h1V7h-1m4 2h-1V8h1V7h1v1h-1m-5 1h1V8h-1m-8 9h1v-1h1v-1h1v-1h1v-1h1v-1H2v2h1v2h1m13 1h1v-1h1v-2h1v-3h-8v1h1v1h1v1h1v1h1v1h1m-3 6H8v-1H5v-1H4v-1H3v-1H2v-1H1v-3H0V0h2v2h8V0h4v1h3v1h1v1h1v1h1v1h1v3h1v6h-1v3h-1v1h-1v1h-1v1h-1v1h-3M10 5V3H2v2m8 3V6H2v2m18 2V8h-1V6h-1V5h-1v1h-1v1h-1v1h-1v1h-1v1m-3 1V9H2v2m9 9v-8h-1v1H9v1H8v1H7v1H6v1H5v1h1v1h2v1m6 0v-1h2v-1h1v-1h-1v-1h-1v-1h-1v-1h-1v-1h-1v7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopSpiralStairsRoundLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 9h1V2H8v1H6v1H5v1h1v1h1v1h1v1h1m4 5h1v-1h-1m2 1h1v-1h-1m2 1h1v-1h-1m2 1h1v-1h-1m-5 2h1v-1h-1m2 1h1v-1h-1m2 1h1v-1h-1m-3 2h1v-1h-1m2 1h1v-1h-1m-1 2h1v-1h-1m2 1h1v-1h-1M4 17h1v-1h1v-1h1v-1h1v-1h1v-1H2v2h1v2h1m11 2h-1v-1h-1v-1h1v1h1m2 0h1v-1h-1m-3 6H8v-1H5v-1H4v-1H3v-1H2v-1H1v-3H0V8h1V5h1V4h1V3h1V2h1V1h3V0h14v2h-2v8h2v4h-1v3h-1v1h-1v1h-1v1h-1v1h-3m-1-11V2h-2v8m5 0V2h-2v8m5 0V2h-2v8m-7 1v-1H9V9H8V8H7V7H6V6H5V5H4v1H3v2H2v3m9 9v-8h-1v1H9v1H8v1H7v1H6v1H5v1h1v1h2v1m6 0v-1h2v-1h1v-1h-1v-1h-1v-1h-1v-1h-1v-1h-1v7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopSpiralStairsRoundRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 6H8V5H7V4h1v1h1M4 6h1V5H4M3 7h1V6H3m2 1h1V6H5M4 8h1V7H4m2 1h1V7H6M3 9h1V8H3m2 1h1V8H5m2 1h1V8H7m2 1h1V2H8v1H6v1H5v1h1v1h1v1h1v1h1m-7 2h1V9H2m2 1h1V9H4m2 1h1V9H6m2 1h1V9H8m3 1h1V9h1V8h1V7h1V6h1V5h1V4h-1V3h-2V2h-3m6 15h1v-1h1v-2h1v-3h-8v1h1v1h1v1h1v1h1v1h1m-3 6H0v-2h2v-8H0V8h1V5h1V4h1V3h1V2h1V1h3V0h6v1h3v1h1v1h1v1h1v1h1v3h1v6h-1v3h-1v1h-1v1h-1v1h-1v1h-3m6-11V8h-1V6h-1V5h-1v1h-1v1h-1v1h-1v1h-1v1M5 20v-8H3v8m5 0v-8H6v8m5 0v-8H9v8m5 0v-1h2v-1h1v-1h-1v-1h-1v-1h-1v-1h-1v-1h-1v7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopSpiralStairsRoundUp(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 9h1V2H8v1H6v1H5v1h1v1h1v1h1v1h1m2 2h1V9h1V8h1V7h1V6h1V5h1V4h-1V3h-2V2h-3M5 15H4v-1h1v-1h1v1H5m4 0h1v-1H9m-1 2h1v-1H8m-1 2h1v-1H7m2 1h1v-1H9m-5 2h1v-1h1v-1h1v-1h1v-1h1v-1H2v2h1v2h1m2 1h1v-1H6m2 1h1v-1H8m-3 2h1v-1H5m2 1h1v-1H7m2 1h1v-1H9m-3 2h1v-1H6m2 1h1v-1H8m1 2h1v-1H9m13 3h-2v-2h-8v2H8v-1H5v-1H4v-1H3v-1H2v-1H1v-3H0V8h1V5h1V4h1V3h1V2h1V1h3V0h6v1h3v1h1v1h1v1h1v1h1v3h1m-2 2V8h-1V6h-1V5h-1v1h-1v1h-1v1h-1v1h-1v1m-3 1v-1H9V9H8V8H7V7H6V6H5V5H4v1H3v2H2v3m18 2v-2h-8v2m8 3v-2h-8v2m8 3v-2h-8v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopSpiralStairsUp(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 9h1V2H3v1h1v1h1v1h1v1h1v1h1v1h1m2 2h1V9h1V8h1V7h1V6h1V5h1V4h1V3h1V2h-8M4 14H3v-1h1m2 1H5v-1h1m3 1h1v-1H9m-5 3H3v-1h1m4 0h1v-1H8m-1 2h1v-1H7m2 1h1v-1H9m-3 2h1v-1H6m2 1h1v-1H8m-3 2h1v-1H5m2 1h1v-1H7m2 1h1v-1H9m-7 2h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1H2m2 7h1v-1H4m2 1h1v-1H6m2 1h1v-1H8m-5 2h1v-1H3m2 1h1v-1H5m2 1h1v-1H7m2 1h1v-1H9m13 3h-2v-2h-8v2H0V0h22m-2 10V3h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1m-3 1v-1H9V9H8V8H7V7H6V6H5V5H4V4H3V3H2v8m18 2v-2h-8v2m8 3v-2h-8v2m8 3v-2h-8v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopStairsDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 4H7V2h8m1 6H6V6h10m1 6H5v-2h12m1 6H4v-2h14m1 6H3v-2h16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopStairsLeft(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 15h-2V7h2m-4 9h-2V6h2m-4 11h-2V5h2M8 18H6V4h2M4 19H2V3h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopStairsRight(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 15H2V7h2m4 9H6V6h2m4 11h-2V5h2m4 13h-2V4h2m4 15h-2V3h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopStairsUp(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 4H3V2h16m-1 6H4V6h14m-1 6H5v-2h12m-1 6H6v-2h10m-1 6H7v-2h8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopVerticalRotateClockwise(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2h-2V0h2m0 9h-2V3h2m0 9h-2v-2h2m7 5h-2v-1h-1v-1h-1v-1h-1v-1h3V9h-1V7h-1V6h-1V4h2v1h1v1h1v2h1v3h3v1h-1v1h-1v1h-1M8 18H6v-1H5v-1H4v-2H3V8h1V6h1V5h1V4h2v2H7v1H6v2H5v4h1v2h1v1h1m4 3h-2v-6h2m0 9h-2v-2h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopVerticalRotateCounterclockwise(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2h-2V0h2m0 9h-2V3h2m0 9h-2v-2h2m-7 5H3v-1H2v-1H1v-1H0v-1h3V8h1V6h1V5h1V4h2v2H7v1H6v2H5v2h3v1H7v1H6v1H5m11 4h-2v-2h1v-1h1v-2h1V9h-1V7h-1V6h-1V4h2v1h1v1h1v2h1v6h-1v2h-1v1h-1m-4 2h-2v-6h2m0 9h-2v-2h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopVerticalStairsAscendDown(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 12H0v-2h2m15 2H5v-2h12m5 2h-2v-2h2m-4 5H4v-2h14m1 5H3v-2h16m1 5H2v-2h18Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTopVerticalStairsAscendUp(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 3H2V1h18m-1 5H3V4h16m-1 5H4V7h14M2 12H0v-2h2m15 2H5v-2h12m5 2h-2v-2h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 2h1V1h9v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1h-1v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-1H1zm2 8h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1v-1h1v-2h-1v-1h-1V9h-1V8h-1V7h-1V6h-1V5h-1V4h-1V3H3zm2-6h2v1h1v2H7v1H5V7H4V5h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagText(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 2h1V1h9v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1h-1v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-1H1zm2 8h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1v-1h1v-2h-1v-1h-1V9h-1V8h-1V7h-1V6h-1V5h-1V4h-1V3H3zm11 1h1v2h-1v-1h-1v-1h-1v-1h-1V9h-1V7h1v1h1v1h1v1h1zm-4 1h1v1h1v2h-1v-1h-1v-1H9v-1H8v-2h1v1h1zM5 4h2v1h1v2H7v1H5V7H4V5h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Target(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 13h-2v-1H9v-2h1V9h2v1h1v2h-1Zm2 4H8v-1H7v-1H6v-1H5V8h1V7h1V6h1V5h6v1h1v1h1v1h1v6h-1v1h-1v1h-1Zm1 4H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2Zm-2-6v-1h1v-1h1V9h-1V8h-1V7H9v1H8v1H7v4h1v1h1v1Zm1 4v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tent(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 19H3v-2h7V8H9v2H8v2H7v2H6v2H3v-1h1v-2h1v-2h1v-1h1V8h1V6h1V4h1V3h2v1h1v2h1v2h1v2h1v1h1v2h1v2h1v1h-3v-2h-1v-2h-1v-2h-1V8h-1v9h7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Terminal(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 3v16h-1v1H3v-1H2V3h1V2h16v1zm-2 3H4v12h14zM9 9v1h1v1h1v2h-1v1H9v1H8v1H6v-2h1v-1h1v-2H7v-1H6V8h2v1zm2 7v-2h5v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextBox(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2h14v1h1v1h1v14h-1v1h-1v1H4v-1H3v-1H2V4h1V3h1zm13 3V4H5v1H4v12h1v1h12v-1h1V5zM6 8h10v2H6zm0 4h7v2H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextImage(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 4h10v10H2zm2 2v6h6V6zm10-2h6v2h-6zm0 4h6v2h-6zm0 4h6v2h-6zM2 16h16v2H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TileCautionHeavy(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 10H0V6h1V5h2V4h2V3h2V2h2V1h2V0h9v1h-2v1h-2v1h-2v1h-2v1h-2v1H8v1H6v1H4v1H2m0 12H0v-4h1v-1h2v-1h2v-1h2v-1h2v-1h2v-1h2v-1h2V9h2V8h2V7h2V6h1v5h-2v1h-2v1h-2v1h-2v1h-2v1h-2v1H8v1H6v1H4v1H2m20 2h-9v-1h2v-1h2v-1h2v-1h2v-1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TileCautionThin(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 10H0V8h1V7h2V6h2V5h2V4h2V3h2V2h2V1h2V0h5v1h-2v1h-2v1h-2v1h-2v1h-2v1H8v1H6v1H4v1H2m0 12H0v-2h1v-1h2v-1h2v-1h2v-1h2v-1h2v-1h2v-1h2v-1h2v-1h2V9h2V8h1v3h-2v1h-2v1h-2v1h-2v1h-2v1h-2v1H8v1H6v1H4v1H2m20 2h-5v-1h2v-1h2v-1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TileDiamondHex(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22h-2v-4H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H0v-2h4V9h1V8h1V7h1V6h1V5h1V4h1V0h2v4h1v1h1v1h1v1h1v1h1v1h1v1h4v2h-4v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1m0-2v-1h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6h-2v1H9v1H8v1H7v1H6v2h1v1h1v1h1v1h1v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleSwitchOff(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 8h4v1h1v4H9v1H5v-1H4V9h1zm14-3v1h1v1h1v8h-1v1h-1v1H3v-1H2v-1H1V7h1V6h1V5zm-1 2H4v1H3v6h1v1h14v-1h1V8h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleSwitchOn(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 5h16v1h1v1h1v8h-1v1h-1v1H3v-1H2v-1H1V7h1V6h1zm10 3v1h-1v4h1v1h4v-1h1V9h-1V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Toolbox(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 6h5V3h1V2h6v1h1v3h5v1h1v12h-1v1H2v-1H1V7h1zm7 0h4V4H9zm10 2H3v4h3v-2h3v2h4v-2h3v2h3zM3 18h16v-4h-3v2h-3v-2H9v2H6v-2H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TooltipAbove(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 1h18v1h1v14h-1v1h-5v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H2v-1H1V2h1zm1 2v12h5v1h1v1h1v1h2v-1h1v-1h1v-1h5V3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TooltipAboveAlert(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 5h2v5h-2zm0 6h2v2h-2zM2 1h18v1h1v14h-1v1h-5v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H2v-1H1V2h1zm1 2v12h5v1h1v1h1v1h2v-1h1v-1h1v-1h5V3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TooltipAboveText(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 1h18v1h1v14h-1v1h-5v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H2v-1H1V2h1zm1 2v12h5v1h1v1h1v1h2v-1h1v-1h1v-1h5V3zm2 3h12v2H5zm0 4h10v2H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TooltipBelow(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 21h18v-1h1V6h-1V5h-5V4h-1V3h-1V2h-1V1h-2v1H9v1H8v1H7v1H2v1H1v14h1zm1-2V7h5V6h1V5h1V4h2v1h1v1h1v1h5v12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TooltipBelowAlert(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 17h2v-2h-2zm0-3h2V9h-2zm-8 7h18v-1h1V6h-1V5h-5V4h-1V3h-1V2h-1V1h-2v1H9v1H8v1H7v1H2v1H1v14h1zm1-2V7h5V6h1V5h1V4h2v1h1v1h1v1h5v12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TooltipBelowText(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 21h18v-1h1V6h-1V5h-5V4h-1V3h-1V2h-1V1h-2v1H9v1H8v1H7v1H2v1H1v14h1zm1-2V7h5V6h1V5h1V4h2v1h1v1h1v1h5v12zm2-3h10v-2H5zm0-4h12v-2H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TooltipEnd(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 20V2h-1V1H6v1H5v5H4v1H3v1H2v1H1v2h1v1h1v1h1v1h1v5h1v1h14v-1zm-2-1H7v-5H6v-1H5v-1H4v-2h1V9h1V8h1V3h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TooltipEndAlert(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 15h-2v-2h2zm0-3h-2V7h2zm7-10v18h-1v1H6v-1H5v-5H4v-1H3v-1H2v-1H1v-2h1V9h1V8h1V7h1V2h1V1h14v1zm-2 1H7v5H6v1H5v1H4v2h1v1h1v1h1v5h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TooltipEndText(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 8V6H9v2zm-2 8v-2H9v2zm2-4v-2H9v2zm4 8V2h-1V1H6v1H5v5H4v1H3v1H2v1H1v2h1v1h1v1h1v1h1v5h1v1h14v-1zm-2-1H7v-5H6v-1H5v-1H4v-2h1V9h1V8h1V3h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TooltipStart(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 20V2h1V1h14v1h1v5h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v5h-1v1H2v-1zm2-1h12v-5h1v-1h1v-1h1v-2h-1V9h-1V8h-1V3H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TooltipStartAlert(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 15h2v-2H8zm0-3h2V7H8zM1 2v18h1v1h14v-1h1v-5h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V2h-1V1H2v1zm2 1h12v5h1v1h1v1h1v2h-1v1h-1v1h-1v5H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TooltipStartText(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 8V6h8v2zm0 8v-2h6v2zm0-4v-2h8v2zm-4 8V2h1V1h14v1h1v5h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v5h-1v1H2v-1zm2-1h12v-5h1v-1h1v-1h1v-2h-1V9h-1V8h-1V3H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToyBrick(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 7h1V6H7m7 1h1V6h-1m5 12H3V7h2V5h1V4h3v1h1v2h2V5h1V4h3v1h1v2h2m-2 9V9H5v7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 7v9H8V7zm2 0h2v9h-2zM8 2h6v1h5v2h-1v14h-1v1H5v-1H4V5H3V3h5zM6 5v13h10V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Umbrella(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 20H7v-1H6v-3h2v2h2v-6H3V7h1V6h1V5h1V4h2V3h6v1h2v1h1v1h1v1h1v5h-7v7h-1m6-9V8h-1V7h-1V6h-2V5H9v1H7v1H6v1H5v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Upload(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 17v2H4v-2zM8 15V9H4V8h1V7h1V6h1V5h1V4h1V3h1V2h2v1h1v1h1v1h1v1h1v1h1v1h1v1h-4v6zm2-2h2V7h1V6h-1V5h-2v1H9v1h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeHigh(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 2h2v1h1v1h1v1h1v1h1v2h1v6h-1v2h-1v1h-1v1h-1v1h-1v1h-2v-2h1v-1h1v-1h1v-1h1v-2h1V9h-1V7h-1V6h-1V5h-1V4h-1zm1 5v1h1v2h1v2h-1v2h-1v1h-1V7zM2 8h4V7h1V6h1V5h1V4h1V3h1v16h-1v-1H9v-1H8v-1H7v-1H6v-1H2zm2 2v2h3v1h1v1h1V8H8v1H7v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeLow(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 8h4V7h1V6h1V5h1V4h1V3h1v16h-1v-1h-1v-1h-1v-1h-1v-1h-1v-1H6zm2 2v2h3v1h1v1h1V8h-1v1h-1v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeMedium(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 7v1h1v2h1v2h-1v2h-1v1h-1V7zM8 8V7h1V6h1V5h1V4h1V3h1v16h-1v-1h-1v-1h-1v-1H9v-1H8v-1H4V8zm-2 2v2h3v1h1v1h1V8h-1v1H9v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeMute(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 7h2v1h1v1h2V8h1V7h2v2h-1v1h-1v2h1v1h1v2h-2v-1h-1v-1h-2v1h-1v1h-2v-2h1v-1h1v-2h-1V9h-1zM6 8V7h1V6h1V5h1V4h1V3h1v16h-1v-1H9v-1H8v-1H7v-1H6v-1H2V8zm1 2H4v2h3v1h1v1h1V8H8v1H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wall(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 20H4v-5H1V7h3V2h15v5h2v8h-2ZM12 7V4H6v3Zm5 0V4h-3v3Zm-9 6V9H3v4Zm11 0V9h-9v4Zm-8 5v-3H6v3Zm6 0v-3h-4v3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WallFill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3h6v4H5m8-4h5v4h-5m-8 8v4h6v-4m2 0h5v4h-5M8 9H3v4h5m2-4h10v4H10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WallFront(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 18H2V4h2v2h2V4h2v2h2V4h2v2h2V4h2v2h2V4h2m-2 12V8H4v8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WallFrontDamaged(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 16h1v-2h1v-1h1v-1h1v-1h2v-1h5v1h1v2h1v2h1V8H4m16 10h-4v-2h-1v-2h-1v-2h-3v1H9v1H8v1H7v2H6v1H2V4h2v2h2V4h2v2h2V4h2v2h2V4h2v2h2V4h2m-9 16H6v-1h1v-1h3v1h1m5 1h-4v-2h1v-1h1v1h1v1h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WallFrontGate(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 18h-7v-5h-1v-1h-2v1H9v5H2V4h2v2h2V4h2v2h2V4h2v2h2V4h2v2h2V4h2m-2 12V8H4v8h3v-4h1v-1h1v-1h4v1h1v1h1v4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Water(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 21H8v-1H6v-1H5v-1H4v-2H3v-3h1v-2h1V9h1V7h1V6h1V4h1V3h1V1h2v2h1v1h1v2h1v1h1v2h1v2h1v2h1v3h-1v2h-1v1h-1v1h-2Zm-1-2v-1h2v-1h1v-2h1v-1h-1v-2h-1v-2h-1V8h-1V7h-1V5h-2v2H9v1H8v2H7v2H6v2H5v1h1v2h1v1h2v1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WaterFill(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 21H8v-1H6v-1H5v-1H4v-2H3v-3h1v-2h1V9h1V7h1V6h1V4h1V3h1V1h2v2h1v1h1v2h1v1h1v2h1v2h1v2h1v3h-1v2h-1v1h-1v1h-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Well(children ...ElementRenderer) *MemoryIcon {
	return &MemoryIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 21H3v-5H1v-2h20v2h-2zM5 16v3h12v-3zM2 7V5h1V4h1V3h1V2h1V1h10v1h1v1h1v1h1v1h1v2h-2v6h-2V7h-4v2h2v4H8V9h2V7H6v6H4V7zm5-4v1H6v1h10V4h-1V3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
