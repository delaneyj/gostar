package heroicons

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "2.1.1"
	hAttr          = "1em"
	viewbox        = "0 0 24 24"
	fill           = "currentColor"
)

type HeroiconsIcon struct {
	*SVGSVGElement
}

type HeroiconsIconFn func(children ...ElementRenderer) *HeroiconsIcon

var IconLookup = map[string]HeroiconsIconFn{
	"academicCap":                            AcademicCap,
	"academicCapSixteenSolid":                AcademicCapSixteenSolid,
	"academicCapSolid":                       AcademicCapSolid,
	"academicCapTwentySolid":                 AcademicCapTwentySolid,
	"adjustmentsHorizontal":                  AdjustmentsHorizontal,
	"adjustmentsHorizontalSixteenSolid":      AdjustmentsHorizontalSixteenSolid,
	"adjustmentsHorizontalSolid":             AdjustmentsHorizontalSolid,
	"adjustmentsHorizontalTwentySolid":       AdjustmentsHorizontalTwentySolid,
	"adjustmentsVertical":                    AdjustmentsVertical,
	"adjustmentsVerticalSixteenSolid":        AdjustmentsVerticalSixteenSolid,
	"adjustmentsVerticalSolid":               AdjustmentsVerticalSolid,
	"adjustmentsVerticalTwentySolid":         AdjustmentsVerticalTwentySolid,
	"archiveBox":                             ArchiveBox,
	"archiveBoxArrowDown":                    ArchiveBoxArrowDown,
	"archiveBoxArrowDownSixteenSolid":        ArchiveBoxArrowDownSixteenSolid,
	"archiveBoxArrowDownSolid":               ArchiveBoxArrowDownSolid,
	"archiveBoxArrowDownTwentySolid":         ArchiveBoxArrowDownTwentySolid,
	"archiveBoxSixteenSolid":                 ArchiveBoxSixteenSolid,
	"archiveBoxSolid":                        ArchiveBoxSolid,
	"archiveBoxTwentySolid":                  ArchiveBoxTwentySolid,
	"archiveBoxXmark":                        ArchiveBoxXmark,
	"archiveBoxXmarkSixteenSolid":            ArchiveBoxXmarkSixteenSolid,
	"archiveBoxXmarkSolid":                   ArchiveBoxXmarkSolid,
	"archiveBoxXmarkTwentySolid":             ArchiveBoxXmarkTwentySolid,
	"arrowDown":                              ArrowDown,
	"arrowDownCircle":                        ArrowDownCircle,
	"arrowDownCircleSixteenSolid":            ArrowDownCircleSixteenSolid,
	"arrowDownCircleSolid":                   ArrowDownCircleSolid,
	"arrowDownCircleTwentySolid":             ArrowDownCircleTwentySolid,
	"arrowDownLeft":                          ArrowDownLeft,
	"arrowDownLeftSixteenSolid":              ArrowDownLeftSixteenSolid,
	"arrowDownLeftSolid":                     ArrowDownLeftSolid,
	"arrowDownLeftTwentySolid":               ArrowDownLeftTwentySolid,
	"arrowDownOnSquare":                      ArrowDownOnSquare,
	"arrowDownOnSquareSixteenSolid":          ArrowDownOnSquareSixteenSolid,
	"arrowDownOnSquareSolid":                 ArrowDownOnSquareSolid,
	"arrowDownOnSquareStack":                 ArrowDownOnSquareStack,
	"arrowDownOnSquareStackSixteenSolid":     ArrowDownOnSquareStackSixteenSolid,
	"arrowDownOnSquareStackSolid":            ArrowDownOnSquareStackSolid,
	"arrowDownOnSquareStackTwentySolid":      ArrowDownOnSquareStackTwentySolid,
	"arrowDownOnSquareTwentySolid":           ArrowDownOnSquareTwentySolid,
	"arrowDownRight":                         ArrowDownRight,
	"arrowDownRightSixteenSolid":             ArrowDownRightSixteenSolid,
	"arrowDownRightSolid":                    ArrowDownRightSolid,
	"arrowDownRightTwentySolid":              ArrowDownRightTwentySolid,
	"arrowDownSixteenSolid":                  ArrowDownSixteenSolid,
	"arrowDownSolid":                         ArrowDownSolid,
	"arrowDownTray":                          ArrowDownTray,
	"arrowDownTraySixteenSolid":              ArrowDownTraySixteenSolid,
	"arrowDownTraySolid":                     ArrowDownTraySolid,
	"arrowDownTrayTwentySolid":               ArrowDownTrayTwentySolid,
	"arrowDownTwentySolid":                   ArrowDownTwentySolid,
	"arrowLeft":                              ArrowLeft,
	"arrowLeftCircle":                        ArrowLeftCircle,
	"arrowLeftCircleSixteenSolid":            ArrowLeftCircleSixteenSolid,
	"arrowLeftCircleSolid":                   ArrowLeftCircleSolid,
	"arrowLeftCircleTwentySolid":             ArrowLeftCircleTwentySolid,
	"arrowLeftEndOnRectangle":                ArrowLeftEndOnRectangle,
	"arrowLeftEndOnRectangleSixteenSolid":    ArrowLeftEndOnRectangleSixteenSolid,
	"arrowLeftEndOnRectangleSolid":           ArrowLeftEndOnRectangleSolid,
	"arrowLeftEndOnRectangleTwentySolid":     ArrowLeftEndOnRectangleTwentySolid,
	"arrowLeftOnRectangle":                   ArrowLeftOnRectangle,
	"arrowLeftOnRectangleSolid":              ArrowLeftOnRectangleSolid,
	"arrowLeftOnRectangleTwentySolid":        ArrowLeftOnRectangleTwentySolid,
	"arrowLeftSixteenSolid":                  ArrowLeftSixteenSolid,
	"arrowLeftSolid":                         ArrowLeftSolid,
	"arrowLeftStartOnRectangle":              ArrowLeftStartOnRectangle,
	"arrowLeftStartOnRectangleSixteenSolid":  ArrowLeftStartOnRectangleSixteenSolid,
	"arrowLeftStartOnRectangleSolid":         ArrowLeftStartOnRectangleSolid,
	"arrowLeftStartOnRectangleTwentySolid":   ArrowLeftStartOnRectangleTwentySolid,
	"arrowLeftTwentySolid":                   ArrowLeftTwentySolid,
	"arrowLongDown":                          ArrowLongDown,
	"arrowLongDownSixteenSolid":              ArrowLongDownSixteenSolid,
	"arrowLongDownSolid":                     ArrowLongDownSolid,
	"arrowLongDownTwentySolid":               ArrowLongDownTwentySolid,
	"arrowLongLeft":                          ArrowLongLeft,
	"arrowLongLeftSixteenSolid":              ArrowLongLeftSixteenSolid,
	"arrowLongLeftSolid":                     ArrowLongLeftSolid,
	"arrowLongLeftTwentySolid":               ArrowLongLeftTwentySolid,
	"arrowLongRight":                         ArrowLongRight,
	"arrowLongRightSixteenSolid":             ArrowLongRightSixteenSolid,
	"arrowLongRightSolid":                    ArrowLongRightSolid,
	"arrowLongRightTwentySolid":              ArrowLongRightTwentySolid,
	"arrowLongUp":                            ArrowLongUp,
	"arrowLongUpSixteenSolid":                ArrowLongUpSixteenSolid,
	"arrowLongUpSolid":                       ArrowLongUpSolid,
	"arrowLongUpTwentySolid":                 ArrowLongUpTwentySolid,
	"arrowPath":                              ArrowPath,
	"arrowPathRoundedSquare":                 ArrowPathRoundedSquare,
	"arrowPathRoundedSquareSixteenSolid":     ArrowPathRoundedSquareSixteenSolid,
	"arrowPathRoundedSquareSolid":            ArrowPathRoundedSquareSolid,
	"arrowPathRoundedSquareTwentySolid":      ArrowPathRoundedSquareTwentySolid,
	"arrowPathSixteenSolid":                  ArrowPathSixteenSolid,
	"arrowPathSolid":                         ArrowPathSolid,
	"arrowPathTwentySolid":                   ArrowPathTwentySolid,
	"arrowRight":                             ArrowRight,
	"arrowRightCircle":                       ArrowRightCircle,
	"arrowRightCircleSixteenSolid":           ArrowRightCircleSixteenSolid,
	"arrowRightCircleSolid":                  ArrowRightCircleSolid,
	"arrowRightCircleTwentySolid":            ArrowRightCircleTwentySolid,
	"arrowRightEndOnRectangle":               ArrowRightEndOnRectangle,
	"arrowRightEndOnRectangleSixteenSolid":   ArrowRightEndOnRectangleSixteenSolid,
	"arrowRightEndOnRectangleSolid":          ArrowRightEndOnRectangleSolid,
	"arrowRightEndOnRectangleTwentySolid":    ArrowRightEndOnRectangleTwentySolid,
	"arrowRightOnRectangle":                  ArrowRightOnRectangle,
	"arrowRightOnRectangleSolid":             ArrowRightOnRectangleSolid,
	"arrowRightOnRectangleTwentySolid":       ArrowRightOnRectangleTwentySolid,
	"arrowRightSixteenSolid":                 ArrowRightSixteenSolid,
	"arrowRightSolid":                        ArrowRightSolid,
	"arrowRightStartOnRectangle":             ArrowRightStartOnRectangle,
	"arrowRightStartOnRectangleSixteenSolid": ArrowRightStartOnRectangleSixteenSolid,
	"arrowRightStartOnRectangleSolid":        ArrowRightStartOnRectangleSolid,
	"arrowRightStartOnRectangleTwentySolid":  ArrowRightStartOnRectangleTwentySolid,
	"arrowRightTwentySolid":                  ArrowRightTwentySolid,
	"arrowSmallDown":                         ArrowSmallDown,
	"arrowSmallDownSolid":                    ArrowSmallDownSolid,
	"arrowSmallDownTwentySolid":              ArrowSmallDownTwentySolid,
	"arrowSmallLeft":                         ArrowSmallLeft,
	"arrowSmallLeftSolid":                    ArrowSmallLeftSolid,
	"arrowSmallLeftTwentySolid":              ArrowSmallLeftTwentySolid,
	"arrowSmallRight":                        ArrowSmallRight,
	"arrowSmallRightSolid":                   ArrowSmallRightSolid,
	"arrowSmallRightTwentySolid":             ArrowSmallRightTwentySolid,
	"arrowSmallUp":                           ArrowSmallUp,
	"arrowSmallUpSolid":                      ArrowSmallUpSolid,
	"arrowSmallUpTwentySolid":                ArrowSmallUpTwentySolid,
	"arrowTopRightOnSquare":                  ArrowTopRightOnSquare,
	"arrowTopRightOnSquareSixteenSolid":      ArrowTopRightOnSquareSixteenSolid,
	"arrowTopRightOnSquareSolid":             ArrowTopRightOnSquareSolid,
	"arrowTopRightOnSquareTwentySolid":       ArrowTopRightOnSquareTwentySolid,
	"arrowTrendingDown":                      ArrowTrendingDown,
	"arrowTrendingDownSixteenSolid":          ArrowTrendingDownSixteenSolid,
	"arrowTrendingDownSolid":                 ArrowTrendingDownSolid,
	"arrowTrendingDownTwentySolid":           ArrowTrendingDownTwentySolid,
	"arrowTrendingUp":                        ArrowTrendingUp,
	"arrowTrendingUpSixteenSolid":            ArrowTrendingUpSixteenSolid,
	"arrowTrendingUpSolid":                   ArrowTrendingUpSolid,
	"arrowTrendingUpTwentySolid":             ArrowTrendingUpTwentySolid,
	"arrowUp":                                ArrowUp,
	"arrowUpCircle":                          ArrowUpCircle,
	"arrowUpCircleSixteenSolid":              ArrowUpCircleSixteenSolid,
	"arrowUpCircleSolid":                     ArrowUpCircleSolid,
	"arrowUpCircleTwentySolid":               ArrowUpCircleTwentySolid,
	"arrowUpLeft":                            ArrowUpLeft,
	"arrowUpLeftSixteenSolid":                ArrowUpLeftSixteenSolid,
	"arrowUpLeftSolid":                       ArrowUpLeftSolid,
	"arrowUpLeftTwentySolid":                 ArrowUpLeftTwentySolid,
	"arrowUpOnSquare":                        ArrowUpOnSquare,
	"arrowUpOnSquareSixteenSolid":            ArrowUpOnSquareSixteenSolid,
	"arrowUpOnSquareSolid":                   ArrowUpOnSquareSolid,
	"arrowUpOnSquareStack":                   ArrowUpOnSquareStack,
	"arrowUpOnSquareStackSixteenSolid":       ArrowUpOnSquareStackSixteenSolid,
	"arrowUpOnSquareStackSolid":              ArrowUpOnSquareStackSolid,
	"arrowUpOnSquareStackTwentySolid":        ArrowUpOnSquareStackTwentySolid,
	"arrowUpOnSquareTwentySolid":             ArrowUpOnSquareTwentySolid,
	"arrowUpRight":                           ArrowUpRight,
	"arrowUpRightSixteenSolid":               ArrowUpRightSixteenSolid,
	"arrowUpRightSolid":                      ArrowUpRightSolid,
	"arrowUpRightTwentySolid":                ArrowUpRightTwentySolid,
	"arrowUpSixteenSolid":                    ArrowUpSixteenSolid,
	"arrowUpSolid":                           ArrowUpSolid,
	"arrowUpTray":                            ArrowUpTray,
	"arrowUpTraySixteenSolid":                ArrowUpTraySixteenSolid,
	"arrowUpTraySolid":                       ArrowUpTraySolid,
	"arrowUpTrayTwentySolid":                 ArrowUpTrayTwentySolid,
	"arrowUpTwentySolid":                     ArrowUpTwentySolid,
	"arrowUturnDown":                         ArrowUturnDown,
	"arrowUturnDownSixteenSolid":             ArrowUturnDownSixteenSolid,
	"arrowUturnDownSolid":                    ArrowUturnDownSolid,
	"arrowUturnDownTwentySolid":              ArrowUturnDownTwentySolid,
	"arrowUturnLeft":                         ArrowUturnLeft,
	"arrowUturnLeftSixteenSolid":             ArrowUturnLeftSixteenSolid,
	"arrowUturnLeftSolid":                    ArrowUturnLeftSolid,
	"arrowUturnLeftTwentySolid":              ArrowUturnLeftTwentySolid,
	"arrowUturnRight":                        ArrowUturnRight,
	"arrowUturnRightSixteenSolid":            ArrowUturnRightSixteenSolid,
	"arrowUturnRightSolid":                   ArrowUturnRightSolid,
	"arrowUturnRightTwentySolid":             ArrowUturnRightTwentySolid,
	"arrowUturnUp":                           ArrowUturnUp,
	"arrowUturnUpSixteenSolid":               ArrowUturnUpSixteenSolid,
	"arrowUturnUpSolid":                      ArrowUturnUpSolid,
	"arrowUturnUpTwentySolid":                ArrowUturnUpTwentySolid,
	"arrowsPointingIn":                       ArrowsPointingIn,
	"arrowsPointingInSixteenSolid":           ArrowsPointingInSixteenSolid,
	"arrowsPointingInSolid":                  ArrowsPointingInSolid,
	"arrowsPointingInTwentySolid":            ArrowsPointingInTwentySolid,
	"arrowsPointingOut":                      ArrowsPointingOut,
	"arrowsPointingOutSixteenSolid":          ArrowsPointingOutSixteenSolid,
	"arrowsPointingOutSolid":                 ArrowsPointingOutSolid,
	"arrowsPointingOutTwentySolid":           ArrowsPointingOutTwentySolid,
	"arrowsRightLeft":                        ArrowsRightLeft,
	"arrowsRightLeftSixteenSolid":            ArrowsRightLeftSixteenSolid,
	"arrowsRightLeftSolid":                   ArrowsRightLeftSolid,
	"arrowsRightLeftTwentySolid":             ArrowsRightLeftTwentySolid,
	"arrowsUpDown":                           ArrowsUpDown,
	"arrowsUpDownSixteenSolid":               ArrowsUpDownSixteenSolid,
	"arrowsUpDownSolid":                      ArrowsUpDownSolid,
	"arrowsUpDownTwentySolid":                ArrowsUpDownTwentySolid,
	"atSymbol":                               AtSymbol,
	"atSymbolSixteenSolid":                   AtSymbolSixteenSolid,
	"atSymbolSolid":                          AtSymbolSolid,
	"atSymbolTwentySolid":                    AtSymbolTwentySolid,
	"backspace":                              Backspace,
	"backspaceSixteenSolid":                  BackspaceSixteenSolid,
	"backspaceSolid":                         BackspaceSolid,
	"backspaceTwentySolid":                   BackspaceTwentySolid,
	"backward":                               Backward,
	"backwardSixteenSolid":                   BackwardSixteenSolid,
	"backwardSolid":                          BackwardSolid,
	"backwardTwentySolid":                    BackwardTwentySolid,
	"banknotes":                              Banknotes,
	"banknotesSixteenSolid":                  BanknotesSixteenSolid,
	"banknotesSolid":                         BanknotesSolid,
	"banknotesTwentySolid":                   BanknotesTwentySolid,
	"barsArrowDown":                          BarsArrowDown,
	"barsArrowDownSixteenSolid":              BarsArrowDownSixteenSolid,
	"barsArrowDownSolid":                     BarsArrowDownSolid,
	"barsArrowDownTwentySolid":               BarsArrowDownTwentySolid,
	"barsArrowUp":                            BarsArrowUp,
	"barsArrowUpSixteenSolid":                BarsArrowUpSixteenSolid,
	"barsArrowUpSolid":                       BarsArrowUpSolid,
	"barsArrowUpTwentySolid":                 BarsArrowUpTwentySolid,
	"barsFour":                               BarsFour,
	"barsFourSixteenSolid":                   BarsFourSixteenSolid,
	"barsFourSolid":                          BarsFourSolid,
	"barsFourTwentySolid":                    BarsFourTwentySolid,
	"barsThree":                              BarsThree,
	"barsThreeBottomLeft":                    BarsThreeBottomLeft,
	"barsThreeBottomLeftSixteenSolid":        BarsThreeBottomLeftSixteenSolid,
	"barsThreeBottomLeftSolid":               BarsThreeBottomLeftSolid,
	"barsThreeBottomLeftTwentySolid":         BarsThreeBottomLeftTwentySolid,
	"barsThreeBottomRight":                   BarsThreeBottomRight,
	"barsThreeBottomRightSixteenSolid":       BarsThreeBottomRightSixteenSolid,
	"barsThreeBottomRightSolid":              BarsThreeBottomRightSolid,
	"barsThreeBottomRightTwentySolid":        BarsThreeBottomRightTwentySolid,
	"barsThreeCenterLeft":                    BarsThreeCenterLeft,
	"barsThreeCenterLeftSixteenSolid":        BarsThreeCenterLeftSixteenSolid,
	"barsThreeCenterLeftSolid":               BarsThreeCenterLeftSolid,
	"barsThreeCenterLeftTwentySolid":         BarsThreeCenterLeftTwentySolid,
	"barsThreeSixteenSolid":                  BarsThreeSixteenSolid,
	"barsThreeSolid":                         BarsThreeSolid,
	"barsThreeTwentySolid":                   BarsThreeTwentySolid,
	"barsTwo":                                BarsTwo,
	"barsTwoSixteenSolid":                    BarsTwoSixteenSolid,
	"barsTwoSolid":                           BarsTwoSolid,
	"barsTwoTwentySolid":                     BarsTwoTwentySolid,
	"batteryFifty":                           BatteryFifty,
	"batteryFiftySixteenSolid":               BatteryFiftySixteenSolid,
	"batteryFiftySolid":                      BatteryFiftySolid,
	"batteryFiftyTwentySolid":                BatteryFiftyTwentySolid,
	"batteryOneHundred":                      BatteryOneHundred,
	"batteryOneHundredSixteenSolid":          BatteryOneHundredSixteenSolid,
	"batteryOneHundredSolid":                 BatteryOneHundredSolid,
	"batteryOneHundredTwentySolid":           BatteryOneHundredTwentySolid,
	"batteryZero":                            BatteryZero,
	"batteryZeroSixteenSolid":                BatteryZeroSixteenSolid,
	"batteryZeroSolid":                       BatteryZeroSolid,
	"batteryZeroTwentySolid":                 BatteryZeroTwentySolid,
	"beaker":                                 Beaker,
	"beakerSixteenSolid":                     BeakerSixteenSolid,
	"beakerSolid":                            BeakerSolid,
	"beakerTwentySolid":                      BeakerTwentySolid,
	"bell":                                   Bell,
	"bellAlert":                              BellAlert,
	"bellAlertSixteenSolid":                  BellAlertSixteenSolid,
	"bellAlertSolid":                         BellAlertSolid,
	"bellAlertTwentySolid":                   BellAlertTwentySolid,
	"bellSixteenSolid":                       BellSixteenSolid,
	"bellSlash":                              BellSlash,
	"bellSlashSixteenSolid":                  BellSlashSixteenSolid,
	"bellSlashSolid":                         BellSlashSolid,
	"bellSlashTwentySolid":                   BellSlashTwentySolid,
	"bellSnooze":                             BellSnooze,
	"bellSnoozeSixteenSolid":                 BellSnoozeSixteenSolid,
	"bellSnoozeSolid":                        BellSnoozeSolid,
	"bellSnoozeTwentySolid":                  BellSnoozeTwentySolid,
	"bellSolid":                              BellSolid,
	"bellTwentySolid":                        BellTwentySolid,
	"bold":                                   Bold,
	"boldSolid":                              BoldSolid,
	"boldTwentySolid":                        BoldTwentySolid,
	"bolt":                                   Bolt,
	"boltSixteenSolid":                       BoltSixteenSolid,
	"boltSlash":                              BoltSlash,
	"boltSlashSixteenSolid":                  BoltSlashSixteenSolid,
	"boltSlashSolid":                         BoltSlashSolid,
	"boltSlashTwentySolid":                   BoltSlashTwentySolid,
	"boltSolid":                              BoltSolid,
	"boltTwentySolid":                        BoltTwentySolid,
	"bookOpen":                               BookOpen,
	"bookOpenSixteenSolid":                   BookOpenSixteenSolid,
	"bookOpenSolid":                          BookOpenSolid,
	"bookOpenTwentySolid":                    BookOpenTwentySolid,
	"bookmark":                               Bookmark,
	"bookmarkSixteenSolid":                   BookmarkSixteenSolid,
	"bookmarkSlash":                          BookmarkSlash,
	"bookmarkSlashSixteenSolid":              BookmarkSlashSixteenSolid,
	"bookmarkSlashSolid":                     BookmarkSlashSolid,
	"bookmarkSlashTwentySolid":               BookmarkSlashTwentySolid,
	"bookmarkSolid":                          BookmarkSolid,
	"bookmarkSquare":                         BookmarkSquare,
	"bookmarkSquareSixteenSolid":             BookmarkSquareSixteenSolid,
	"bookmarkSquareSolid":                    BookmarkSquareSolid,
	"bookmarkSquareTwentySolid":              BookmarkSquareTwentySolid,
	"bookmarkTwentySolid":                    BookmarkTwentySolid,
	"briefcase":                              Briefcase,
	"briefcaseSixteenSolid":                  BriefcaseSixteenSolid,
	"briefcaseSolid":                         BriefcaseSolid,
	"briefcaseTwentySolid":                   BriefcaseTwentySolid,
	"bugAnt":                                 BugAnt,
	"bugAntSixteenSolid":                     BugAntSixteenSolid,
	"bugAntSolid":                            BugAntSolid,
	"bugAntTwentySolid":                      BugAntTwentySolid,
	"buildingLibrary":                        BuildingLibrary,
	"buildingLibrarySixteenSolid":            BuildingLibrarySixteenSolid,
	"buildingLibrarySolid":                   BuildingLibrarySolid,
	"buildingLibraryTwentySolid":             BuildingLibraryTwentySolid,
	"buildingOffice":                         BuildingOffice,
	"buildingOfficeSixteenSolid":             BuildingOfficeSixteenSolid,
	"buildingOfficeSolid":                    BuildingOfficeSolid,
	"buildingOfficeTwentySolid":              BuildingOfficeTwentySolid,
	"buildingOfficeTwo":                      BuildingOfficeTwo,
	"buildingOfficeTwoSixteenSolid":          BuildingOfficeTwoSixteenSolid,
	"buildingOfficeTwoSolid":                 BuildingOfficeTwoSolid,
	"buildingOfficeTwoTwentySolid":           BuildingOfficeTwoTwentySolid,
	"buildingStorefront":                     BuildingStorefront,
	"buildingStorefrontSixteenSolid":         BuildingStorefrontSixteenSolid,
	"buildingStorefrontSolid":                BuildingStorefrontSolid,
	"buildingStorefrontTwentySolid":          BuildingStorefrontTwentySolid,
	"cake":                                   Cake,
	"cakeSixteenSolid":                       CakeSixteenSolid,
	"cakeSolid":                              CakeSolid,
	"cakeTwentySolid":                        CakeTwentySolid,
	"calculator":                             Calculator,
	"calculatorSixteenSolid":                 CalculatorSixteenSolid,
	"calculatorSolid":                        CalculatorSolid,
	"calculatorTwentySolid":                  CalculatorTwentySolid,
	"calendar":                               Calendar,
	"calendarDays":                           CalendarDays,
	"calendarDaysSixteenSolid":               CalendarDaysSixteenSolid,
	"calendarDaysSolid":                      CalendarDaysSolid,
	"calendarDaysTwentySolid":                CalendarDaysTwentySolid,
	"calendarSixteenSolid":                   CalendarSixteenSolid,
	"calendarSolid":                          CalendarSolid,
	"calendarTwentySolid":                    CalendarTwentySolid,
	"camera":                                 Camera,
	"cameraSixteenSolid":                     CameraSixteenSolid,
	"cameraSolid":                            CameraSolid,
	"cameraTwentySolid":                      CameraTwentySolid,
	"chartBar":                               ChartBar,
	"chartBarSixteenSolid":                   ChartBarSixteenSolid,
	"chartBarSolid":                          ChartBarSolid,
	"chartBarSquare":                         ChartBarSquare,
	"chartBarSquareSixteenSolid":             ChartBarSquareSixteenSolid,
	"chartBarSquareSolid":                    ChartBarSquareSolid,
	"chartBarSquareTwentySolid":              ChartBarSquareTwentySolid,
	"chartBarTwentySolid":                    ChartBarTwentySolid,
	"chartPie":                               ChartPie,
	"chartPieSixteenSolid":                   ChartPieSixteenSolid,
	"chartPieSolid":                          ChartPieSolid,
	"chartPieTwentySolid":                    ChartPieTwentySolid,
	"chatBubbleBottomCenter":                 ChatBubbleBottomCenter,
	"chatBubbleBottomCenterSixteenSolid":     ChatBubbleBottomCenterSixteenSolid,
	"chatBubbleBottomCenterSolid":            ChatBubbleBottomCenterSolid,
	"chatBubbleBottomCenterText":             ChatBubbleBottomCenterText,
	"chatBubbleBottomCenterTextSixteenSolid": ChatBubbleBottomCenterTextSixteenSolid,
	"chatBubbleBottomCenterTextSolid":        ChatBubbleBottomCenterTextSolid,
	"chatBubbleBottomCenterTextTwentySolid":  ChatBubbleBottomCenterTextTwentySolid,
	"chatBubbleBottomCenterTwentySolid":      ChatBubbleBottomCenterTwentySolid,
	"chatBubbleLeft":                         ChatBubbleLeft,
	"chatBubbleLeftEllipsis":                 ChatBubbleLeftEllipsis,
	"chatBubbleLeftEllipsisSixteenSolid":     ChatBubbleLeftEllipsisSixteenSolid,
	"chatBubbleLeftEllipsisSolid":            ChatBubbleLeftEllipsisSolid,
	"chatBubbleLeftEllipsisTwentySolid":      ChatBubbleLeftEllipsisTwentySolid,
	"chatBubbleLeftRight":                    ChatBubbleLeftRight,
	"chatBubbleLeftRightSixteenSolid":        ChatBubbleLeftRightSixteenSolid,
	"chatBubbleLeftRightSolid":               ChatBubbleLeftRightSolid,
	"chatBubbleLeftRightTwentySolid":         ChatBubbleLeftRightTwentySolid,
	"chatBubbleLeftSixteenSolid":             ChatBubbleLeftSixteenSolid,
	"chatBubbleLeftSolid":                    ChatBubbleLeftSolid,
	"chatBubbleLeftTwentySolid":              ChatBubbleLeftTwentySolid,
	"chatBubbleOvalLeft":                     ChatBubbleOvalLeft,
	"chatBubbleOvalLeftEllipsis":             ChatBubbleOvalLeftEllipsis,
	"chatBubbleOvalLeftEllipsisSixteenSolid": ChatBubbleOvalLeftEllipsisSixteenSolid,
	"chatBubbleOvalLeftEllipsisSolid":        ChatBubbleOvalLeftEllipsisSolid,
	"chatBubbleOvalLeftEllipsisTwentySolid":  ChatBubbleOvalLeftEllipsisTwentySolid,
	"chatBubbleOvalLeftSixteenSolid":         ChatBubbleOvalLeftSixteenSolid,
	"chatBubbleOvalLeftSolid":                ChatBubbleOvalLeftSolid,
	"chatBubbleOvalLeftTwentySolid":          ChatBubbleOvalLeftTwentySolid,
	"check":                                  Check,
	"checkBadge":                             CheckBadge,
	"checkBadgeSixteenSolid":                 CheckBadgeSixteenSolid,
	"checkBadgeSolid":                        CheckBadgeSolid,
	"checkBadgeTwentySolid":                  CheckBadgeTwentySolid,
	"checkCircle":                            CheckCircle,
	"checkCircleSixteenSolid":                CheckCircleSixteenSolid,
	"checkCircleSolid":                       CheckCircleSolid,
	"checkCircleTwentySolid":                 CheckCircleTwentySolid,
	"checkSixteenSolid":                      CheckSixteenSolid,
	"checkSolid":                             CheckSolid,
	"checkTwentySolid":                       CheckTwentySolid,
	"chevronDoubleDown":                      ChevronDoubleDown,
	"chevronDoubleDownSixteenSolid":          ChevronDoubleDownSixteenSolid,
	"chevronDoubleDownSolid":                 ChevronDoubleDownSolid,
	"chevronDoubleDownTwentySolid":           ChevronDoubleDownTwentySolid,
	"chevronDoubleLeft":                      ChevronDoubleLeft,
	"chevronDoubleLeftSixteenSolid":          ChevronDoubleLeftSixteenSolid,
	"chevronDoubleLeftSolid":                 ChevronDoubleLeftSolid,
	"chevronDoubleLeftTwentySolid":           ChevronDoubleLeftTwentySolid,
	"chevronDoubleRight":                     ChevronDoubleRight,
	"chevronDoubleRightSixteenSolid":         ChevronDoubleRightSixteenSolid,
	"chevronDoubleRightSolid":                ChevronDoubleRightSolid,
	"chevronDoubleRightTwentySolid":          ChevronDoubleRightTwentySolid,
	"chevronDoubleUp":                        ChevronDoubleUp,
	"chevronDoubleUpSixteenSolid":            ChevronDoubleUpSixteenSolid,
	"chevronDoubleUpSolid":                   ChevronDoubleUpSolid,
	"chevronDoubleUpTwentySolid":             ChevronDoubleUpTwentySolid,
	"chevronDown":                            ChevronDown,
	"chevronDownSixteenSolid":                ChevronDownSixteenSolid,
	"chevronDownSolid":                       ChevronDownSolid,
	"chevronDownTwentySolid":                 ChevronDownTwentySolid,
	"chevronLeft":                            ChevronLeft,
	"chevronLeftSixteenSolid":                ChevronLeftSixteenSolid,
	"chevronLeftSolid":                       ChevronLeftSolid,
	"chevronLeftTwentySolid":                 ChevronLeftTwentySolid,
	"chevronRight":                           ChevronRight,
	"chevronRightSixteenSolid":               ChevronRightSixteenSolid,
	"chevronRightSolid":                      ChevronRightSolid,
	"chevronRightTwentySolid":                ChevronRightTwentySolid,
	"chevronUp":                              ChevronUp,
	"chevronUpDown":                          ChevronUpDown,
	"chevronUpDownSixteenSolid":              ChevronUpDownSixteenSolid,
	"chevronUpDownSolid":                     ChevronUpDownSolid,
	"chevronUpDownTwentySolid":               ChevronUpDownTwentySolid,
	"chevronUpSixteenSolid":                  ChevronUpSixteenSolid,
	"chevronUpSolid":                         ChevronUpSolid,
	"chevronUpTwentySolid":                   ChevronUpTwentySolid,
	"circleStack":                            CircleStack,
	"circleStackSixteenSolid":                CircleStackSixteenSolid,
	"circleStackSolid":                       CircleStackSolid,
	"circleStackTwentySolid":                 CircleStackTwentySolid,
	"clipboard":                              Clipboard,
	"clipboardDocument":                      ClipboardDocument,
	"clipboardDocumentCheck":                 ClipboardDocumentCheck,
	"clipboardDocumentCheckSixteenSolid":     ClipboardDocumentCheckSixteenSolid,
	"clipboardDocumentCheckSolid":            ClipboardDocumentCheckSolid,
	"clipboardDocumentCheckTwentySolid":      ClipboardDocumentCheckTwentySolid,
	"clipboardDocumentList":                  ClipboardDocumentList,
	"clipboardDocumentListSixteenSolid":      ClipboardDocumentListSixteenSolid,
	"clipboardDocumentListSolid":             ClipboardDocumentListSolid,
	"clipboardDocumentListTwentySolid":       ClipboardDocumentListTwentySolid,
	"clipboardDocumentSixteenSolid":          ClipboardDocumentSixteenSolid,
	"clipboardDocumentSolid":                 ClipboardDocumentSolid,
	"clipboardDocumentTwentySolid":           ClipboardDocumentTwentySolid,
	"clipboardSixteenSolid":                  ClipboardSixteenSolid,
	"clipboardSolid":                         ClipboardSolid,
	"clipboardTwentySolid":                   ClipboardTwentySolid,
	"clock":                                  Clock,
	"clockSixteenSolid":                      ClockSixteenSolid,
	"clockSolid":                             ClockSolid,
	"clockTwentySolid":                       ClockTwentySolid,
	"cloud":                                  Cloud,
	"cloudArrowDown":                         CloudArrowDown,
	"cloudArrowDownSixteenSolid":             CloudArrowDownSixteenSolid,
	"cloudArrowDownSolid":                    CloudArrowDownSolid,
	"cloudArrowDownTwentySolid":              CloudArrowDownTwentySolid,
	"cloudArrowUp":                           CloudArrowUp,
	"cloudArrowUpSixteenSolid":               CloudArrowUpSixteenSolid,
	"cloudArrowUpSolid":                      CloudArrowUpSolid,
	"cloudArrowUpTwentySolid":                CloudArrowUpTwentySolid,
	"cloudSixteenSolid":                      CloudSixteenSolid,
	"cloudSolid":                             CloudSolid,
	"cloudTwentySolid":                       CloudTwentySolid,
	"codeBracket":                            CodeBracket,
	"codeBracketSixteenSolid":                CodeBracketSixteenSolid,
	"codeBracketSolid":                       CodeBracketSolid,
	"codeBracketSquare":                      CodeBracketSquare,
	"codeBracketSquareSixteenSolid":          CodeBracketSquareSixteenSolid,
	"codeBracketSquareSolid":                 CodeBracketSquareSolid,
	"codeBracketSquareTwentySolid":           CodeBracketSquareTwentySolid,
	"codeBracketTwentySolid":                 CodeBracketTwentySolid,
	"cog":                                    Cog,
	"cogEightTooth":                          CogEightTooth,
	"cogEightToothSixteenSolid":              CogEightToothSixteenSolid,
	"cogEightToothSolid":                     CogEightToothSolid,
	"cogEightToothTwentySolid":               CogEightToothTwentySolid,
	"cogSixTooth":                            CogSixTooth,
	"cogSixToothSixteenSolid":                CogSixToothSixteenSolid,
	"cogSixToothSolid":                       CogSixToothSolid,
	"cogSixToothTwentySolid":                 CogSixToothTwentySolid,
	"cogSixteenSolid":                        CogSixteenSolid,
	"cogSolid":                               CogSolid,
	"cogTwentySolid":                         CogTwentySolid,
	"commandLine":                            CommandLine,
	"commandLineSixteenSolid":                CommandLineSixteenSolid,
	"commandLineSolid":                       CommandLineSolid,
	"commandLineTwentySolid":                 CommandLineTwentySolid,
	"computerDesktop":                        ComputerDesktop,
	"computerDesktopSixteenSolid":            ComputerDesktopSixteenSolid,
	"computerDesktopSolid":                   ComputerDesktopSolid,
	"computerDesktopTwentySolid":             ComputerDesktopTwentySolid,
	"cpuChip":                                CpuChip,
	"cpuChipSixteenSolid":                    CpuChipSixteenSolid,
	"cpuChipSolid":                           CpuChipSolid,
	"cpuChipTwentySolid":                     CpuChipTwentySolid,
	"creditCard":                             CreditCard,
	"creditCardSixteenSolid":                 CreditCardSixteenSolid,
	"creditCardSolid":                        CreditCardSolid,
	"creditCardTwentySolid":                  CreditCardTwentySolid,
	"cube":                                   Cube,
	"cubeSixteenSolid":                       CubeSixteenSolid,
	"cubeSolid":                              CubeSolid,
	"cubeTransparent":                        CubeTransparent,
	"cubeTransparentSixteenSolid":            CubeTransparentSixteenSolid,
	"cubeTransparentSolid":                   CubeTransparentSolid,
	"cubeTransparentTwentySolid":             CubeTransparentTwentySolid,
	"cubeTwentySolid":                        CubeTwentySolid,
	"currencyBangladeshi":                    CurrencyBangladeshi,
	"currencyBangladeshiSixteenSolid":        CurrencyBangladeshiSixteenSolid,
	"currencyBangladeshiSolid":               CurrencyBangladeshiSolid,
	"currencyBangladeshiTwentySolid":         CurrencyBangladeshiTwentySolid,
	"currencyDollar":                         CurrencyDollar,
	"currencyDollarSixteenSolid":             CurrencyDollarSixteenSolid,
	"currencyDollarSolid":                    CurrencyDollarSolid,
	"currencyDollarTwentySolid":              CurrencyDollarTwentySolid,
	"currencyEuro":                           CurrencyEuro,
	"currencyEuroSixteenSolid":               CurrencyEuroSixteenSolid,
	"currencyEuroSolid":                      CurrencyEuroSolid,
	"currencyEuroTwentySolid":                CurrencyEuroTwentySolid,
	"currencyPound":                          CurrencyPound,
	"currencyPoundSixteenSolid":              CurrencyPoundSixteenSolid,
	"currencyPoundSolid":                     CurrencyPoundSolid,
	"currencyPoundTwentySolid":               CurrencyPoundTwentySolid,
	"currencyRupee":                          CurrencyRupee,
	"currencyRupeeSixteenSolid":              CurrencyRupeeSixteenSolid,
	"currencyRupeeSolid":                     CurrencyRupeeSolid,
	"currencyRupeeTwentySolid":               CurrencyRupeeTwentySolid,
	"currencyYen":                            CurrencyYen,
	"currencyYenSixteenSolid":                CurrencyYenSixteenSolid,
	"currencyYenSolid":                       CurrencyYenSolid,
	"currencyYenTwentySolid":                 CurrencyYenTwentySolid,
	"cursorArrowRays":                        CursorArrowRays,
	"cursorArrowRaysSixteenSolid":            CursorArrowRaysSixteenSolid,
	"cursorArrowRaysSolid":                   CursorArrowRaysSolid,
	"cursorArrowRaysTwentySolid":             CursorArrowRaysTwentySolid,
	"cursorArrowRipple":                      CursorArrowRipple,
	"cursorArrowRippleSixteenSolid":          CursorArrowRippleSixteenSolid,
	"cursorArrowRippleSolid":                 CursorArrowRippleSolid,
	"cursorArrowRippleTwentySolid":           CursorArrowRippleTwentySolid,
	"devicePhoneMobile":                      DevicePhoneMobile,
	"devicePhoneMobileSixteenSolid":          DevicePhoneMobileSixteenSolid,
	"devicePhoneMobileSolid":                 DevicePhoneMobileSolid,
	"devicePhoneMobileTwentySolid":           DevicePhoneMobileTwentySolid,
	"deviceTablet":                           DeviceTablet,
	"deviceTabletSixteenSolid":               DeviceTabletSixteenSolid,
	"deviceTabletSolid":                      DeviceTabletSolid,
	"deviceTabletTwentySolid":                DeviceTabletTwentySolid,
	"document":                               Document,
	"documentArrowDown":                      DocumentArrowDown,
	"documentArrowDownSixteenSolid":          DocumentArrowDownSixteenSolid,
	"documentArrowDownSolid":                 DocumentArrowDownSolid,
	"documentArrowDownTwentySolid":           DocumentArrowDownTwentySolid,
	"documentArrowUp":                        DocumentArrowUp,
	"documentArrowUpSixteenSolid":            DocumentArrowUpSixteenSolid,
	"documentArrowUpSolid":                   DocumentArrowUpSolid,
	"documentArrowUpTwentySolid":             DocumentArrowUpTwentySolid,
	"documentChartBar":                       DocumentChartBar,
	"documentChartBarSixteenSolid":           DocumentChartBarSixteenSolid,
	"documentChartBarSolid":                  DocumentChartBarSolid,
	"documentChartBarTwentySolid":            DocumentChartBarTwentySolid,
	"documentCheck":                          DocumentCheck,
	"documentCheckSixteenSolid":              DocumentCheckSixteenSolid,
	"documentCheckSolid":                     DocumentCheckSolid,
	"documentCheckTwentySolid":               DocumentCheckTwentySolid,
	"documentDuplicate":                      DocumentDuplicate,
	"documentDuplicateSixteenSolid":          DocumentDuplicateSixteenSolid,
	"documentDuplicateSolid":                 DocumentDuplicateSolid,
	"documentDuplicateTwentySolid":           DocumentDuplicateTwentySolid,
	"documentMagnifyingGlass":                DocumentMagnifyingGlass,
	"documentMagnifyingGlassSixteenSolid":    DocumentMagnifyingGlassSixteenSolid,
	"documentMagnifyingGlassSolid":           DocumentMagnifyingGlassSolid,
	"documentMagnifyingGlassTwentySolid":     DocumentMagnifyingGlassTwentySolid,
	"documentMinus":                          DocumentMinus,
	"documentMinusSixteenSolid":              DocumentMinusSixteenSolid,
	"documentMinusSolid":                     DocumentMinusSolid,
	"documentMinusTwentySolid":               DocumentMinusTwentySolid,
	"documentPlus":                           DocumentPlus,
	"documentPlusSixteenSolid":               DocumentPlusSixteenSolid,
	"documentPlusSolid":                      DocumentPlusSolid,
	"documentPlusTwentySolid":                DocumentPlusTwentySolid,
	"documentSixteenSolid":                   DocumentSixteenSolid,
	"documentSolid":                          DocumentSolid,
	"documentText":                           DocumentText,
	"documentTextSixteenSolid":               DocumentTextSixteenSolid,
	"documentTextSolid":                      DocumentTextSolid,
	"documentTextTwentySolid":                DocumentTextTwentySolid,
	"documentTwentySolid":                    DocumentTwentySolid,
	"ellipsisHorizontal":                     EllipsisHorizontal,
	"ellipsisHorizontalCircle":               EllipsisHorizontalCircle,
	"ellipsisHorizontalCircleSixteenSolid":   EllipsisHorizontalCircleSixteenSolid,
	"ellipsisHorizontalCircleSolid":          EllipsisHorizontalCircleSolid,
	"ellipsisHorizontalCircleTwentySolid":    EllipsisHorizontalCircleTwentySolid,
	"ellipsisHorizontalSixteenSolid":         EllipsisHorizontalSixteenSolid,
	"ellipsisHorizontalSolid":                EllipsisHorizontalSolid,
	"ellipsisHorizontalTwentySolid":          EllipsisHorizontalTwentySolid,
	"ellipsisVertical":                       EllipsisVertical,
	"ellipsisVerticalSixteenSolid":           EllipsisVerticalSixteenSolid,
	"ellipsisVerticalSolid":                  EllipsisVerticalSolid,
	"ellipsisVerticalTwentySolid":            EllipsisVerticalTwentySolid,
	"envelope":                               Envelope,
	"envelopeOpen":                           EnvelopeOpen,
	"envelopeOpenSixteenSolid":               EnvelopeOpenSixteenSolid,
	"envelopeOpenSolid":                      EnvelopeOpenSolid,
	"envelopeOpenTwentySolid":                EnvelopeOpenTwentySolid,
	"envelopeSixteenSolid":                   EnvelopeSixteenSolid,
	"envelopeSolid":                          EnvelopeSolid,
	"envelopeTwentySolid":                    EnvelopeTwentySolid,
	"exclamationCircle":                      ExclamationCircle,
	"exclamationCircleSixteenSolid":          ExclamationCircleSixteenSolid,
	"exclamationCircleSolid":                 ExclamationCircleSolid,
	"exclamationCircleTwentySolid":           ExclamationCircleTwentySolid,
	"exclamationTriangle":                    ExclamationTriangle,
	"exclamationTriangleSixteenSolid":        ExclamationTriangleSixteenSolid,
	"exclamationTriangleSolid":               ExclamationTriangleSolid,
	"exclamationTriangleTwentySolid":         ExclamationTriangleTwentySolid,
	"eye":                                    Eye,
	"eyeDropper":                             EyeDropper,
	"eyeDropperSixteenSolid":                 EyeDropperSixteenSolid,
	"eyeDropperSolid":                        EyeDropperSolid,
	"eyeDropperTwentySolid":                  EyeDropperTwentySolid,
	"eyeSixteenSolid":                        EyeSixteenSolid,
	"eyeSlash":                               EyeSlash,
	"eyeSlashSixteenSolid":                   EyeSlashSixteenSolid,
	"eyeSlashSolid":                          EyeSlashSolid,
	"eyeSlashTwentySolid":                    EyeSlashTwentySolid,
	"eyeSolid":                               EyeSolid,
	"eyeTwentySolid":                         EyeTwentySolid,
	"faceFrown":                              FaceFrown,
	"faceFrownSixteenSolid":                  FaceFrownSixteenSolid,
	"faceFrownSolid":                         FaceFrownSolid,
	"faceFrownTwentySolid":                   FaceFrownTwentySolid,
	"faceSmile":                              FaceSmile,
	"faceSmileSixteenSolid":                  FaceSmileSixteenSolid,
	"faceSmileSolid":                         FaceSmileSolid,
	"faceSmileTwentySolid":                   FaceSmileTwentySolid,
	"film":                                   Film,
	"filmSixteenSolid":                       FilmSixteenSolid,
	"filmSolid":                              FilmSolid,
	"filmTwentySolid":                        FilmTwentySolid,
	"fingerPrint":                            FingerPrint,
	"fingerPrintSixteenSolid":                FingerPrintSixteenSolid,
	"fingerPrintSolid":                       FingerPrintSolid,
	"fingerPrintTwentySolid":                 FingerPrintTwentySolid,
	"fire":                                   Fire,
	"fireSixteenSolid":                       FireSixteenSolid,
	"fireSolid":                              FireSolid,
	"fireTwentySolid":                        FireTwentySolid,
	"flag":                                   Flag,
	"flagSixteenSolid":                       FlagSixteenSolid,
	"flagSolid":                              FlagSolid,
	"flagTwentySolid":                        FlagTwentySolid,
	"folder":                                 Folder,
	"folderArrowDown":                        FolderArrowDown,
	"folderArrowDownSixteenSolid":            FolderArrowDownSixteenSolid,
	"folderArrowDownSolid":                   FolderArrowDownSolid,
	"folderArrowDownTwentySolid":             FolderArrowDownTwentySolid,
	"folderMinus":                            FolderMinus,
	"folderMinusSixteenSolid":                FolderMinusSixteenSolid,
	"folderMinusSolid":                       FolderMinusSolid,
	"folderMinusTwentySolid":                 FolderMinusTwentySolid,
	"folderOpen":                             FolderOpen,
	"folderOpenSixteenSolid":                 FolderOpenSixteenSolid,
	"folderOpenSolid":                        FolderOpenSolid,
	"folderOpenTwentySolid":                  FolderOpenTwentySolid,
	"folderPlus":                             FolderPlus,
	"folderPlusSixteenSolid":                 FolderPlusSixteenSolid,
	"folderPlusSolid":                        FolderPlusSolid,
	"folderPlusTwentySolid":                  FolderPlusTwentySolid,
	"folderSixteenSolid":                     FolderSixteenSolid,
	"folderSolid":                            FolderSolid,
	"folderTwentySolid":                      FolderTwentySolid,
	"forward":                                Forward,
	"forwardSixteenSolid":                    ForwardSixteenSolid,
	"forwardSolid":                           ForwardSolid,
	"forwardTwentySolid":                     ForwardTwentySolid,
	"funnel":                                 Funnel,
	"funnelSixteenSolid":                     FunnelSixteenSolid,
	"funnelSolid":                            FunnelSolid,
	"funnelTwentySolid":                      FunnelTwentySolid,
	"gif":                                    Gif,
	"gifSixteenSolid":                        GifSixteenSolid,
	"gifSolid":                               GifSolid,
	"gifTwentySolid":                         GifTwentySolid,
	"gift":                                   Gift,
	"giftSixteenSolid":                       GiftSixteenSolid,
	"giftSolid":                              GiftSolid,
	"giftTop":                                GiftTop,
	"giftTopSixteenSolid":                    GiftTopSixteenSolid,
	"giftTopSolid":                           GiftTopSolid,
	"giftTopTwentySolid":                     GiftTopTwentySolid,
	"giftTwentySolid":                        GiftTwentySolid,
	"globeAlt":                               GlobeAlt,
	"globeAltSixteenSolid":                   GlobeAltSixteenSolid,
	"globeAltSolid":                          GlobeAltSolid,
	"globeAltTwentySolid":                    GlobeAltTwentySolid,
	"globeAmericas":                          GlobeAmericas,
	"globeAmericasSixteenSolid":              GlobeAmericasSixteenSolid,
	"globeAmericasSolid":                     GlobeAmericasSolid,
	"globeAmericasTwentySolid":               GlobeAmericasTwentySolid,
	"globeAsiaAustralia":                     GlobeAsiaAustralia,
	"globeAsiaAustraliaSixteenSolid":         GlobeAsiaAustraliaSixteenSolid,
	"globeAsiaAustraliaSolid":                GlobeAsiaAustraliaSolid,
	"globeAsiaAustraliaTwentySolid":          GlobeAsiaAustraliaTwentySolid,
	"globeEuropeAfrica":                      GlobeEuropeAfrica,
	"globeEuropeAfricaSixteenSolid":          GlobeEuropeAfricaSixteenSolid,
	"globeEuropeAfricaSolid":                 GlobeEuropeAfricaSolid,
	"globeEuropeAfricaTwentySolid":           GlobeEuropeAfricaTwentySolid,
	"handRaised":                             HandRaised,
	"handRaisedSixteenSolid":                 HandRaisedSixteenSolid,
	"handRaisedSolid":                        HandRaisedSolid,
	"handRaisedTwentySolid":                  HandRaisedTwentySolid,
	"handThumbDown":                          HandThumbDown,
	"handThumbDownSixteenSolid":              HandThumbDownSixteenSolid,
	"handThumbDownSolid":                     HandThumbDownSolid,
	"handThumbDownTwentySolid":               HandThumbDownTwentySolid,
	"handThumbUp":                            HandThumbUp,
	"handThumbUpSixteenSolid":                HandThumbUpSixteenSolid,
	"handThumbUpSolid":                       HandThumbUpSolid,
	"handThumbUpTwentySolid":                 HandThumbUpTwentySolid,
	"hashtag":                                Hashtag,
	"hashtagSixteenSolid":                    HashtagSixteenSolid,
	"hashtagSolid":                           HashtagSolid,
	"hashtagTwentySolid":                     HashtagTwentySolid,
	"heart":                                  Heart,
	"heartSixteenSolid":                      HeartSixteenSolid,
	"heartSolid":                             HeartSolid,
	"heartTwentySolid":                       HeartTwentySolid,
	"home":                                   Home,
	"homeModern":                             HomeModern,
	"homeModernSixteenSolid":                 HomeModernSixteenSolid,
	"homeModernSolid":                        HomeModernSolid,
	"homeModernTwentySolid":                  HomeModernTwentySolid,
	"homeSixteenSolid":                       HomeSixteenSolid,
	"homeSolid":                              HomeSolid,
	"homeTwentySolid":                        HomeTwentySolid,
	"identification":                         Identification,
	"identificationSixteenSolid":             IdentificationSixteenSolid,
	"identificationSolid":                    IdentificationSolid,
	"identificationTwentySolid":              IdentificationTwentySolid,
	"inbox":                                  Inbox,
	"inboxArrowDown":                         InboxArrowDown,
	"inboxArrowDownSixteenSolid":             InboxArrowDownSixteenSolid,
	"inboxArrowDownSolid":                    InboxArrowDownSolid,
	"inboxArrowDownTwentySolid":              InboxArrowDownTwentySolid,
	"inboxSixteenSolid":                      InboxSixteenSolid,
	"inboxSolid":                             InboxSolid,
	"inboxStack":                             InboxStack,
	"inboxStackSixteenSolid":                 InboxStackSixteenSolid,
	"inboxStackSolid":                        InboxStackSolid,
	"inboxStackTwentySolid":                  InboxStackTwentySolid,
	"inboxTwentySolid":                       InboxTwentySolid,
	"informationCircle":                      InformationCircle,
	"informationCircleSixteenSolid":          InformationCircleSixteenSolid,
	"informationCircleSolid":                 InformationCircleSolid,
	"informationCircleTwentySolid":           InformationCircleTwentySolid,
	"italic":                                 Italic,
	"italicSolid":                            ItalicSolid,
	"italicTwentySolid":                      ItalicTwentySolid,
	"key":                                    Key,
	"keySixteenSolid":                        KeySixteenSolid,
	"keySolid":                               KeySolid,
	"keyTwentySolid":                         KeyTwentySolid,
	"language":                               Language,
	"languageSixteenSolid":                   LanguageSixteenSolid,
	"languageSolid":                          LanguageSolid,
	"languageTwentySolid":                    LanguageTwentySolid,
	"lifebuoy":                               Lifebuoy,
	"lifebuoySixteenSolid":                   LifebuoySixteenSolid,
	"lifebuoySolid":                          LifebuoySolid,
	"lifebuoyTwentySolid":                    LifebuoyTwentySolid,
	"lightBulb":                              LightBulb,
	"lightBulbSixteenSolid":                  LightBulbSixteenSolid,
	"lightBulbSolid":                         LightBulbSolid,
	"lightBulbTwentySolid":                   LightBulbTwentySolid,
	"link":                                   Link,
	"linkSixteenSolid":                       LinkSixteenSolid,
	"linkSolid":                              LinkSolid,
	"linkTwentySolid":                        LinkTwentySolid,
	"listBullet":                             ListBullet,
	"listBulletSixteenSolid":                 ListBulletSixteenSolid,
	"listBulletSolid":                        ListBulletSolid,
	"listBulletTwentySolid":                  ListBulletTwentySolid,
	"lockClosed":                             LockClosed,
	"lockClosedSixteenSolid":                 LockClosedSixteenSolid,
	"lockClosedSolid":                        LockClosedSolid,
	"lockClosedTwentySolid":                  LockClosedTwentySolid,
	"lockOpen":                               LockOpen,
	"lockOpenSixteenSolid":                   LockOpenSixteenSolid,
	"lockOpenSolid":                          LockOpenSolid,
	"lockOpenTwentySolid":                    LockOpenTwentySolid,
	"magnifyingGlass":                        MagnifyingGlass,
	"magnifyingGlassCircle":                  MagnifyingGlassCircle,
	"magnifyingGlassCircleSixteenSolid":      MagnifyingGlassCircleSixteenSolid,
	"magnifyingGlassCircleSolid":             MagnifyingGlassCircleSolid,
	"magnifyingGlassCircleTwentySolid":       MagnifyingGlassCircleTwentySolid,
	"magnifyingGlassMinus":                   MagnifyingGlassMinus,
	"magnifyingGlassMinusSixteenSolid":       MagnifyingGlassMinusSixteenSolid,
	"magnifyingGlassMinusSolid":              MagnifyingGlassMinusSolid,
	"magnifyingGlassMinusTwentySolid":        MagnifyingGlassMinusTwentySolid,
	"magnifyingGlassPlus":                    MagnifyingGlassPlus,
	"magnifyingGlassPlusSixteenSolid":        MagnifyingGlassPlusSixteenSolid,
	"magnifyingGlassPlusSolid":               MagnifyingGlassPlusSolid,
	"magnifyingGlassPlusTwentySolid":         MagnifyingGlassPlusTwentySolid,
	"magnifyingGlassSixteenSolid":            MagnifyingGlassSixteenSolid,
	"magnifyingGlassSolid":                   MagnifyingGlassSolid,
	"magnifyingGlassTwentySolid":             MagnifyingGlassTwentySolid,
	"map":                                    Map,
	"mapPin":                                 MapPin,
	"mapPinSixteenSolid":                     MapPinSixteenSolid,
	"mapPinSolid":                            MapPinSolid,
	"mapPinTwentySolid":                      MapPinTwentySolid,
	"mapSixteenSolid":                        MapSixteenSolid,
	"mapSolid":                               MapSolid,
	"mapTwentySolid":                         MapTwentySolid,
	"megaphone":                              Megaphone,
	"megaphoneSixteenSolid":                  MegaphoneSixteenSolid,
	"megaphoneSolid":                         MegaphoneSolid,
	"megaphoneTwentySolid":                   MegaphoneTwentySolid,
	"microphone":                             Microphone,
	"microphoneSixteenSolid":                 MicrophoneSixteenSolid,
	"microphoneSolid":                        MicrophoneSolid,
	"microphoneTwentySolid":                  MicrophoneTwentySolid,
	"minus":                                  Minus,
	"minusCircle":                            MinusCircle,
	"minusCircleSixteenSolid":                MinusCircleSixteenSolid,
	"minusCircleSolid":                       MinusCircleSolid,
	"minusCircleTwentySolid":                 MinusCircleTwentySolid,
	"minusSixteenSolid":                      MinusSixteenSolid,
	"minusSmall":                             MinusSmall,
	"minusSmallSolid":                        MinusSmallSolid,
	"minusSmallTwentySolid":                  MinusSmallTwentySolid,
	"minusSolid":                             MinusSolid,
	"minusTwentySolid":                       MinusTwentySolid,
	"moon":                                   Moon,
	"moonSixteenSolid":                       MoonSixteenSolid,
	"moonSolid":                              MoonSolid,
	"moonTwentySolid":                        MoonTwentySolid,
	"musicalNote":                            MusicalNote,
	"musicalNoteSixteenSolid":                MusicalNoteSixteenSolid,
	"musicalNoteSolid":                       MusicalNoteSolid,
	"musicalNoteTwentySolid":                 MusicalNoteTwentySolid,
	"newspaper":                              Newspaper,
	"newspaperSixteenSolid":                  NewspaperSixteenSolid,
	"newspaperSolid":                         NewspaperSolid,
	"newspaperTwentySolid":                   NewspaperTwentySolid,
	"noSymbol":                               NoSymbol,
	"noSymbolSixteenSolid":                   NoSymbolSixteenSolid,
	"noSymbolSolid":                          NoSymbolSolid,
	"noSymbolTwentySolid":                    NoSymbolTwentySolid,
	"paintBrush":                             PaintBrush,
	"paintBrushSixteenSolid":                 PaintBrushSixteenSolid,
	"paintBrushSolid":                        PaintBrushSolid,
	"paintBrushTwentySolid":                  PaintBrushTwentySolid,
	"paperAirplane":                          PaperAirplane,
	"paperAirplaneSixteenSolid":              PaperAirplaneSixteenSolid,
	"paperAirplaneSolid":                     PaperAirplaneSolid,
	"paperAirplaneTwentySolid":               PaperAirplaneTwentySolid,
	"paperClip":                              PaperClip,
	"paperClipSixteenSolid":                  PaperClipSixteenSolid,
	"paperClipSolid":                         PaperClipSolid,
	"paperClipTwentySolid":                   PaperClipTwentySolid,
	"pause":                                  Pause,
	"pauseCircle":                            PauseCircle,
	"pauseCircleSixteenSolid":                PauseCircleSixteenSolid,
	"pauseCircleSolid":                       PauseCircleSolid,
	"pauseCircleTwentySolid":                 PauseCircleTwentySolid,
	"pauseSixteenSolid":                      PauseSixteenSolid,
	"pauseSolid":                             PauseSolid,
	"pauseTwentySolid":                       PauseTwentySolid,
	"pencil":                                 Pencil,
	"pencilSixteenSolid":                     PencilSixteenSolid,
	"pencilSolid":                            PencilSolid,
	"pencilSquare":                           PencilSquare,
	"pencilSquareSixteenSolid":               PencilSquareSixteenSolid,
	"pencilSquareSolid":                      PencilSquareSolid,
	"pencilSquareTwentySolid":                PencilSquareTwentySolid,
	"pencilTwentySolid":                      PencilTwentySolid,
	"phone":                                  Phone,
	"phoneArrowDownLeft":                     PhoneArrowDownLeft,
	"phoneArrowDownLeftSixteenSolid":         PhoneArrowDownLeftSixteenSolid,
	"phoneArrowDownLeftSolid":                PhoneArrowDownLeftSolid,
	"phoneArrowDownLeftTwentySolid":          PhoneArrowDownLeftTwentySolid,
	"phoneArrowUpRight":                      PhoneArrowUpRight,
	"phoneArrowUpRightSixteenSolid":          PhoneArrowUpRightSixteenSolid,
	"phoneArrowUpRightSolid":                 PhoneArrowUpRightSolid,
	"phoneArrowUpRightTwentySolid":           PhoneArrowUpRightTwentySolid,
	"phoneSixteenSolid":                      PhoneSixteenSolid,
	"phoneSolid":                             PhoneSolid,
	"phoneTwentySolid":                       PhoneTwentySolid,
	"phoneXmark":                             PhoneXmark,
	"phoneXmarkSixteenSolid":                 PhoneXmarkSixteenSolid,
	"phoneXmarkSolid":                        PhoneXmarkSolid,
	"phoneXmarkTwentySolid":                  PhoneXmarkTwentySolid,
	"photo":                                  Photo,
	"photoSixteenSolid":                      PhotoSixteenSolid,
	"photoSolid":                             PhotoSolid,
	"photoTwentySolid":                       PhotoTwentySolid,
	"play":                                   Play,
	"playCircle":                             PlayCircle,
	"playCircleSixteenSolid":                 PlayCircleSixteenSolid,
	"playCircleSolid":                        PlayCircleSolid,
	"playCircleTwentySolid":                  PlayCircleTwentySolid,
	"playPause":                              PlayPause,
	"playPauseSixteenSolid":                  PlayPauseSixteenSolid,
	"playPauseSolid":                         PlayPauseSolid,
	"playPauseTwentySolid":                   PlayPauseTwentySolid,
	"playSixteenSolid":                       PlaySixteenSolid,
	"playSolid":                              PlaySolid,
	"playTwentySolid":                        PlayTwentySolid,
	"plus":                                   Plus,
	"plusCircle":                             PlusCircle,
	"plusCircleSixteenSolid":                 PlusCircleSixteenSolid,
	"plusCircleSolid":                        PlusCircleSolid,
	"plusCircleTwentySolid":                  PlusCircleTwentySolid,
	"plusSixteenSolid":                       PlusSixteenSolid,
	"plusSmall":                              PlusSmall,
	"plusSmallSolid":                         PlusSmallSolid,
	"plusSmallTwentySolid":                   PlusSmallTwentySolid,
	"plusSolid":                              PlusSolid,
	"plusTwentySolid":                        PlusTwentySolid,
	"power":                                  Power,
	"powerSixteenSolid":                      PowerSixteenSolid,
	"powerSolid":                             PowerSolid,
	"powerTwentySolid":                       PowerTwentySolid,
	"presentationChartBar":                   PresentationChartBar,
	"presentationChartBarSixteenSolid":       PresentationChartBarSixteenSolid,
	"presentationChartBarSolid":              PresentationChartBarSolid,
	"presentationChartBarTwentySolid":        PresentationChartBarTwentySolid,
	"presentationChartLine":                  PresentationChartLine,
	"presentationChartLineSixteenSolid":      PresentationChartLineSixteenSolid,
	"presentationChartLineSolid":             PresentationChartLineSolid,
	"presentationChartLineTwentySolid":       PresentationChartLineTwentySolid,
	"printer":                                Printer,
	"printerSixteenSolid":                    PrinterSixteenSolid,
	"printerSolid":                           PrinterSolid,
	"printerTwentySolid":                     PrinterTwentySolid,
	"puzzlePiece":                            PuzzlePiece,
	"puzzlePieceSixteenSolid":                PuzzlePieceSixteenSolid,
	"puzzlePieceSolid":                       PuzzlePieceSolid,
	"puzzlePieceTwentySolid":                 PuzzlePieceTwentySolid,
	"qrCode":                                 QrCode,
	"qrCodeSixteenSolid":                     QrCodeSixteenSolid,
	"qrCodeSolid":                            QrCodeSolid,
	"qrCodeTwentySolid":                      QrCodeTwentySolid,
	"questionMarkCircle":                     QuestionMarkCircle,
	"questionMarkCircleSixteenSolid":         QuestionMarkCircleSixteenSolid,
	"questionMarkCircleSolid":                QuestionMarkCircleSolid,
	"questionMarkCircleTwentySolid":          QuestionMarkCircleTwentySolid,
	"queueList":                              QueueList,
	"queueListSixteenSolid":                  QueueListSixteenSolid,
	"queueListSolid":                         QueueListSolid,
	"queueListTwentySolid":                   QueueListTwentySolid,
	"radio":                                  Radio,
	"radioSixteenSolid":                      RadioSixteenSolid,
	"radioSolid":                             RadioSolid,
	"radioTwentySolid":                       RadioTwentySolid,
	"receiptPercent":                         ReceiptPercent,
	"receiptPercentSixteenSolid":             ReceiptPercentSixteenSolid,
	"receiptPercentSolid":                    ReceiptPercentSolid,
	"receiptPercentTwentySolid":              ReceiptPercentTwentySolid,
	"receiptRefund":                          ReceiptRefund,
	"receiptRefundSixteenSolid":              ReceiptRefundSixteenSolid,
	"receiptRefundSolid":                     ReceiptRefundSolid,
	"receiptRefundTwentySolid":               ReceiptRefundTwentySolid,
	"rectangleGroup":                         RectangleGroup,
	"rectangleGroupSixteenSolid":             RectangleGroupSixteenSolid,
	"rectangleGroupSolid":                    RectangleGroupSolid,
	"rectangleGroupTwentySolid":              RectangleGroupTwentySolid,
	"rectangleStack":                         RectangleStack,
	"rectangleStackSixteenSolid":             RectangleStackSixteenSolid,
	"rectangleStackSolid":                    RectangleStackSolid,
	"rectangleStackTwentySolid":              RectangleStackTwentySolid,
	"rocketLaunch":                           RocketLaunch,
	"rocketLaunchSixteenSolid":               RocketLaunchSixteenSolid,
	"rocketLaunchSolid":                      RocketLaunchSolid,
	"rocketLaunchTwentySolid":                RocketLaunchTwentySolid,
	"rss":                                    Rss,
	"rssSixteenSolid":                        RssSixteenSolid,
	"rssSolid":                               RssSolid,
	"rssTwentySolid":                         RssTwentySolid,
	"scale":                                  Scale,
	"scaleSixteenSolid":                      ScaleSixteenSolid,
	"scaleSolid":                             ScaleSolid,
	"scaleTwentySolid":                       ScaleTwentySolid,
	"scissors":                               Scissors,
	"scissorsSixteenSolid":                   ScissorsSixteenSolid,
	"scissorsSolid":                          ScissorsSolid,
	"scissorsTwentySolid":                    ScissorsTwentySolid,
	"server":                                 Server,
	"serverSixteenSolid":                     ServerSixteenSolid,
	"serverSolid":                            ServerSolid,
	"serverStack":                            ServerStack,
	"serverStackSixteenSolid":                ServerStackSixteenSolid,
	"serverStackSolid":                       ServerStackSolid,
	"serverStackTwentySolid":                 ServerStackTwentySolid,
	"serverTwentySolid":                      ServerTwentySolid,
	"share":                                  Share,
	"shareSixteenSolid":                      ShareSixteenSolid,
	"shareSolid":                             ShareSolid,
	"shareTwentySolid":                       ShareTwentySolid,
	"shieldCheck":                            ShieldCheck,
	"shieldCheckSixteenSolid":                ShieldCheckSixteenSolid,
	"shieldCheckSolid":                       ShieldCheckSolid,
	"shieldCheckTwentySolid":                 ShieldCheckTwentySolid,
	"shieldExclamation":                      ShieldExclamation,
	"shieldExclamationSixteenSolid":          ShieldExclamationSixteenSolid,
	"shieldExclamationSolid":                 ShieldExclamationSolid,
	"shieldExclamationTwentySolid":           ShieldExclamationTwentySolid,
	"shoppingBag":                            ShoppingBag,
	"shoppingBagSixteenSolid":                ShoppingBagSixteenSolid,
	"shoppingBagSolid":                       ShoppingBagSolid,
	"shoppingBagTwentySolid":                 ShoppingBagTwentySolid,
	"shoppingCart":                           ShoppingCart,
	"shoppingCartSixteenSolid":               ShoppingCartSixteenSolid,
	"shoppingCartSolid":                      ShoppingCartSolid,
	"shoppingCartTwentySolid":                ShoppingCartTwentySolid,
	"signal":                                 Signal,
	"signalSixteenSolid":                     SignalSixteenSolid,
	"signalSlash":                            SignalSlash,
	"signalSlashSixteenSolid":                SignalSlashSixteenSolid,
	"signalSlashSolid":                       SignalSlashSolid,
	"signalSlashTwentySolid":                 SignalSlashTwentySolid,
	"signalSolid":                            SignalSolid,
	"signalTwentySolid":                      SignalTwentySolid,
	"sparkles":                               Sparkles,
	"sparklesSixteenSolid":                   SparklesSixteenSolid,
	"sparklesSolid":                          SparklesSolid,
	"sparklesTwentySolid":                    SparklesTwentySolid,
	"speakerWave":                            SpeakerWave,
	"speakerWaveSixteenSolid":                SpeakerWaveSixteenSolid,
	"speakerWaveSolid":                       SpeakerWaveSolid,
	"speakerWaveTwentySolid":                 SpeakerWaveTwentySolid,
	"speakerXmark":                           SpeakerXmark,
	"speakerXmarkSixteenSolid":               SpeakerXmarkSixteenSolid,
	"speakerXmarkSolid":                      SpeakerXmarkSolid,
	"speakerXmarkTwentySolid":                SpeakerXmarkTwentySolid,
	"squareThreeStackThreeD":                 SquareThreeStackThreeD,
	"squareThreeStackThreeDsixteenSolid":     SquareThreeStackThreeDsixteenSolid,
	"squareThreeStackThreeDsolid":            SquareThreeStackThreeDsolid,
	"squareThreeStackThreeDtwentySolid":      SquareThreeStackThreeDtwentySolid,
	"squareTwoStack":                         SquareTwoStack,
	"squareTwoStackSixteenSolid":             SquareTwoStackSixteenSolid,
	"squareTwoStackSolid":                    SquareTwoStackSolid,
	"squareTwoStackTwentySolid":              SquareTwoStackTwentySolid,
	"squaresPlus":                            SquaresPlus,
	"squaresPlusSixteenSolid":                SquaresPlusSixteenSolid,
	"squaresPlusSolid":                       SquaresPlusSolid,
	"squaresPlusTwentySolid":                 SquaresPlusTwentySolid,
	"squaresTwoXtwo":                         SquaresTwoXtwo,
	"squaresTwoXtwoSixteenSolid":             SquaresTwoXtwoSixteenSolid,
	"squaresTwoXtwoSolid":                    SquaresTwoXtwoSolid,
	"squaresTwoXtwoTwentySolid":              SquaresTwoXtwoTwentySolid,
	"star":                                   Star,
	"starSixteenSolid":                       StarSixteenSolid,
	"starSolid":                              StarSolid,
	"starTwentySolid":                        StarTwentySolid,
	"stop":                                   Stop,
	"stopCircle":                             StopCircle,
	"stopCircleSixteenSolid":                 StopCircleSixteenSolid,
	"stopCircleSolid":                        StopCircleSolid,
	"stopCircleTwentySolid":                  StopCircleTwentySolid,
	"stopSixteenSolid":                       StopSixteenSolid,
	"stopSolid":                              StopSolid,
	"stopTwentySolid":                        StopTwentySolid,
	"sun":                                    Sun,
	"sunSixteenSolid":                        SunSixteenSolid,
	"sunSolid":                               SunSolid,
	"sunTwentySolid":                         SunTwentySolid,
	"swatch":                                 Swatch,
	"swatchSixteenSolid":                     SwatchSixteenSolid,
	"swatchSolid":                            SwatchSolid,
	"swatchTwentySolid":                      SwatchTwentySolid,
	"tableCells":                             TableCells,
	"tableCellsSixteenSolid":                 TableCellsSixteenSolid,
	"tableCellsSolid":                        TableCellsSolid,
	"tableCellsTwentySolid":                  TableCellsTwentySolid,
	"tag":                                    Tag,
	"tagSixteenSolid":                        TagSixteenSolid,
	"tagSolid":                               TagSolid,
	"tagTwentySolid":                         TagTwentySolid,
	"ticket":                                 Ticket,
	"ticketSixteenSolid":                     TicketSixteenSolid,
	"ticketSolid":                            TicketSolid,
	"ticketTwentySolid":                      TicketTwentySolid,
	"trash":                                  Trash,
	"trashSixteenSolid":                      TrashSixteenSolid,
	"trashSolid":                             TrashSolid,
	"trashTwentySolid":                       TrashTwentySolid,
	"trophy":                                 Trophy,
	"trophySixteenSolid":                     TrophySixteenSolid,
	"trophySolid":                            TrophySolid,
	"trophyTwentySolid":                      TrophyTwentySolid,
	"truck":                                  Truck,
	"truckSixteenSolid":                      TruckSixteenSolid,
	"truckSolid":                             TruckSolid,
	"truckTwentySolid":                       TruckTwentySolid,
	"tv":                                     Tv,
	"tvSixteenSolid":                         TvSixteenSolid,
	"tvSolid":                                TvSolid,
	"tvTwentySolid":                          TvTwentySolid,
	"underline":                              Underline,
	"underlineSolid":                         UnderlineSolid,
	"underlineTwentySolid":                   UnderlineTwentySolid,
	"user":                                   User,
	"userCircle":                             UserCircle,
	"userCircleSixteenSolid":                 UserCircleSixteenSolid,
	"userCircleSolid":                        UserCircleSolid,
	"userCircleTwentySolid":                  UserCircleTwentySolid,
	"userGroup":                              UserGroup,
	"userGroupSixteenSolid":                  UserGroupSixteenSolid,
	"userGroupSolid":                         UserGroupSolid,
	"userGroupTwentySolid":                   UserGroupTwentySolid,
	"userMinus":                              UserMinus,
	"userMinusSixteenSolid":                  UserMinusSixteenSolid,
	"userMinusSolid":                         UserMinusSolid,
	"userMinusTwentySolid":                   UserMinusTwentySolid,
	"userPlus":                               UserPlus,
	"userPlusSixteenSolid":                   UserPlusSixteenSolid,
	"userPlusSolid":                          UserPlusSolid,
	"userPlusTwentySolid":                    UserPlusTwentySolid,
	"userSixteenSolid":                       UserSixteenSolid,
	"userSolid":                              UserSolid,
	"userTwentySolid":                        UserTwentySolid,
	"users":                                  Users,
	"usersSixteenSolid":                      UsersSixteenSolid,
	"usersSolid":                             UsersSolid,
	"usersTwentySolid":                       UsersTwentySolid,
	"variable":                               Variable,
	"variableSixteenSolid":                   VariableSixteenSolid,
	"variableSolid":                          VariableSolid,
	"variableTwentySolid":                    VariableTwentySolid,
	"videoCamera":                            VideoCamera,
	"videoCameraSixteenSolid":                VideoCameraSixteenSolid,
	"videoCameraSlash":                       VideoCameraSlash,
	"videoCameraSlashSixteenSolid":           VideoCameraSlashSixteenSolid,
	"videoCameraSlashSolid":                  VideoCameraSlashSolid,
	"videoCameraSlashTwentySolid":            VideoCameraSlashTwentySolid,
	"videoCameraSolid":                       VideoCameraSolid,
	"videoCameraTwentySolid":                 VideoCameraTwentySolid,
	"viewColumns":                            ViewColumns,
	"viewColumnsSixteenSolid":                ViewColumnsSixteenSolid,
	"viewColumnsSolid":                       ViewColumnsSolid,
	"viewColumnsTwentySolid":                 ViewColumnsTwentySolid,
	"viewfinderCircle":                       ViewfinderCircle,
	"viewfinderCircleSixteenSolid":           ViewfinderCircleSixteenSolid,
	"viewfinderCircleSolid":                  ViewfinderCircleSolid,
	"viewfinderCircleTwentySolid":            ViewfinderCircleTwentySolid,
	"wallet":                                 Wallet,
	"walletSixteenSolid":                     WalletSixteenSolid,
	"walletSolid":                            WalletSolid,
	"walletTwentySolid":                      WalletTwentySolid,
	"wifi":                                   Wifi,
	"wifiSixteenSolid":                       WifiSixteenSolid,
	"wifiSolid":                              WifiSolid,
	"wifiTwentySolid":                        WifiTwentySolid,
	"window":                                 Window,
	"windowSixteenSolid":                     WindowSixteenSolid,
	"windowSolid":                            WindowSolid,
	"windowTwentySolid":                      WindowTwentySolid,
	"wrench":                                 Wrench,
	"wrenchScrewdriver":                      WrenchScrewdriver,
	"wrenchScrewdriverSixteenSolid":          WrenchScrewdriverSixteenSolid,
	"wrenchScrewdriverSolid":                 WrenchScrewdriverSolid,
	"wrenchScrewdriverTwentySolid":           WrenchScrewdriverTwentySolid,
	"wrenchSixteenSolid":                     WrenchSixteenSolid,
	"wrenchSolid":                            WrenchSolid,
	"wrenchTwentySolid":                      WrenchTwentySolid,
	"xcircle":                                Xcircle,
	"xcircleSixteenSolid":                    XcircleSixteenSolid,
	"xcircleSolid":                           XcircleSolid,
	"xcircleTwentySolid":                     XcircleTwentySolid,
	"xmark":                                  Xmark,
	"xmarkSixteenSolid":                      XmarkSixteenSolid,
	"xmarkSolid":                             XmarkSolid,
	"xmarkTwentySolid":                       XmarkTwentySolid,
}

func AcademicCap(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.26 10.147a60.438 60.438 0 0 0-.491 6.347A48.62 48.62 0 0 1 12 20.904a48.62 48.62 0 0 1 8.232-4.41a60.46 60.46 0 0 0-.491-6.347m-15.482 0a50.636 50.636 0 0 0-2.658-.813A59.906 59.906 0 0 1 12 3.493a59.903 59.903 0 0 1 10.399 5.84a51.39 51.39 0 0 0-2.658.814m-15.482 0A50.717 50.717 0 0 1 12 13.489a50.702 50.702 0 0 1 7.74-3.342M6.75 15a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5m0 0v-3.675A55.378 55.378 0 0 1 12 8.443m-7.007 11.55A5.981 5.981 0 0 0 6.75 15.75v-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AcademicCapSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M7.702 1.368a.75.75 0 0 1 .597 0c2.098.91 4.105 1.99 6.004 3.223a.75.75 0 0 1-.194 1.348A34.27 34.27 0 0 0 8.341 8.25a.75.75 0 0 1-.682 0c-.625-.32-1.262-.62-1.909-.901v-.542a36.878 36.878 0 0 1 2.568-1.33a.75.75 0 0 0-.636-1.357a38.39 38.39 0 0 0-3.06 1.605a.75.75 0 0 0-.372.648v.365c-.773-.294-1.56-.56-2.359-.8a.75.75 0 0 1-.194-1.347a40.901 40.901 0 0 1 6.005-3.223M4.25 8.348a32.594 32.594 0 0 0-1.611-.596a40.973 40.973 0 0 0-.418 2.97a.75.75 0 0 0 .474.776c.175.068.35.138.524.21a5.544 5.544 0 0 1-.58.681a.75.75 0 1 0 1.06 1.06a6.97 6.97 0 0 0 .915-1.124a29.282 29.282 0 0 0-1.395-.617A5.483 5.483 0 0 0 4.25 8.5z"/><path d="M7.603 13.96c-.96-.6-1.958-1.147-2.989-1.635a6.981 6.981 0 0 0 1.12-3.341c.419.192.834.393 1.244.602a2.25 2.25 0 0 0 2.045 0a32.787 32.787 0 0 1 4.338-1.834c.175.978.315 1.969.419 2.97a.75.75 0 0 1-.474.776a29.385 29.385 0 0 0-4.909 2.461a.75.75 0 0 1-.794 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AcademicCapSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M11.7 2.805a.75.75 0 0 1 .6 0A60.65 60.65 0 0 1 22.83 8.72a.75.75 0 0 1-.231 1.337a49.948 49.948 0 0 0-9.902 3.912l-.003.002c-.114.06-.227.119-.34.18a.75.75 0 0 1-.707 0A50.88 50.88 0 0 0 7.5 12.173v-.224a.36.36 0 0 1 .172-.311a54.615 54.615 0 0 1 4.653-2.52a.75.75 0 0 0-.65-1.352a56.123 56.123 0 0 0-4.78 2.589a1.858 1.858 0 0 0-.859 1.228a49.803 49.803 0 0 0-4.634-1.527a.75.75 0 0 1-.231-1.337A60.653 60.653 0 0 1 11.7 2.805"/><path d="M13.06 15.473a48.45 48.45 0 0 1 7.666-3.282c.134 1.414.22 2.843.255 4.284a.75.75 0 0 1-.46.711a47.87 47.87 0 0 0-8.105 4.342a.75.75 0 0 1-.832 0a47.87 47.87 0 0 0-8.104-4.342a.75.75 0 0 1-.461-.71c.035-1.442.121-2.87.255-4.286c.921.304 1.83.634 2.726.99v1.27a1.5 1.5 0 0 0-.14 2.508c-.09.38-.222.753-.397 1.11c.452.213.901.434 1.346.66a6.727 6.727 0 0 0 .551-1.607a1.5 1.5 0 0 0 .14-2.67v-.645a48.549 48.549 0 0 1 3.44 1.667a2.25 2.25 0 0 0 2.12 0"/><path d="M4.462 19.462c.42-.419.753-.89 1-1.395c.453.214.902.435 1.347.662a6.742 6.742 0 0 1-1.286 1.794a.75.75 0 0 1-1.06-1.06"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AcademicCapTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.664 1.319a.75.75 0 0 1 .672 0a41.059 41.059 0 0 1 8.198 5.424a.75.75 0 0 1-.254 1.285a31.372 31.372 0 0 0-7.86 3.83a.75.75 0 0 1-.84 0a31.508 31.508 0 0 0-2.08-1.287V9.394c0-.244.116-.463.302-.592a35.504 35.504 0 0 1 3.305-2.033a.75.75 0 0 0-.714-1.319a37 37 0 0 0-3.446 2.12A2.216 2.216 0 0 0 6 9.393v.38a31.293 31.293 0 0 0-4.28-1.746a.75.75 0 0 1-.254-1.285a41.059 41.059 0 0 1 8.198-5.424M6 11.459a29.848 29.848 0 0 0-2.455-1.158a41.029 41.029 0 0 0-.39 3.114a.75.75 0 0 0 .419.74c.528.256 1.046.53 1.554.82c-.21.324-.455.63-.739.914a.75.75 0 1 0 1.06 1.06c.37-.369.69-.77.96-1.193a26.61 26.61 0 0 1 3.095 2.348a.75.75 0 0 0 .992 0a26.547 26.547 0 0 1 5.93-3.95a.75.75 0 0 0 .42-.739a41.053 41.053 0 0 0-.39-3.114a29.925 29.925 0 0 0-5.199 2.801a2.25 2.25 0 0 1-2.514 0c-.41-.275-.826-.541-1.25-.797a6.985 6.985 0 0 1-1.084 3.45a26.503 26.503 0 0 0-1.281-.78A5.487 5.487 0 0 0 6 12z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdjustmentsHorizontal(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.5 6h9.75M10.5 6a1.5 1.5 0 1 1-3 0m3 0a1.5 1.5 0 1 0-3 0M3.75 6H7.5m3 12h9.75m-9.75 0a1.5 1.5 0 0 1-3 0m3 0a1.5 1.5 0 0 0-3 0m-3.75 0H7.5m9-6h3.75m-3.75 0a1.5 1.5 0 0 1-3 0m3 0a1.5 1.5 0 0 0-3 0m-9.75 0h9.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdjustmentsHorizontalSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.5 2.25a.75.75 0 0 0-1.5 0v3a.75.75 0 0 0 1.5 0V4.5h6.75a.75.75 0 0 0 0-1.5H6.5zM11 6.5a.75.75 0 0 0-1.5 0v3a.75.75 0 0 0 1.5 0v-.75h2.25a.75.75 0 0 0 0-1.5H11zM5.75 10a.75.75 0 0 1 .75.75v.75h6.75a.75.75 0 0 1 0 1.5H6.5v.75a.75.75 0 0 1-1.5 0v-3a.75.75 0 0 1 .75-.75m-3-2.75H8.5v1.5H2.75a.75.75 0 0 1 0-1.5M4 3H2.75a.75.75 0 0 0 0 1.5H4zm-1.25 8.5H4V13H2.75a.75.75 0 0 1 0-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdjustmentsHorizontalSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.75 12.75h1.5a.75.75 0 0 0 0-1.5h-1.5a.75.75 0 0 0 0 1.5M12 6a.75.75 0 0 1 .75-.75h7.5a.75.75 0 0 1 0 1.5h-7.5A.75.75 0 0 1 12 6m0 12a.75.75 0 0 1 .75-.75h7.5a.75.75 0 0 1 0 1.5h-7.5A.75.75 0 0 1 12 18M3.75 6.75h1.5a.75.75 0 1 0 0-1.5h-1.5a.75.75 0 0 0 0 1.5m1.5 12h-1.5a.75.75 0 0 1 0-1.5h1.5a.75.75 0 0 1 0 1.5M3 12a.75.75 0 0 1 .75-.75h7.5a.75.75 0 0 1 0 1.5h-7.5A.75.75 0 0 1 3 12m6-8.25a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5M12.75 12a2.25 2.25 0 1 1 4.5 0a2.25 2.25 0 0 1-4.5 0M9 15.75a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdjustmentsHorizontalTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 3.75a2 2 0 1 0-4 0a2 2 0 0 0 4 0m7.25.75a.75.75 0 0 0 0-1.5h-5.5a.75.75 0 0 0 0 1.5zM5 3.75a.75.75 0 0 1-.75.75h-1.5a.75.75 0 0 1 0-1.5h1.5a.75.75 0 0 1 .75.75M4.25 17a.75.75 0 0 0 0-1.5h-1.5a.75.75 0 0 0 0 1.5zm13 0a.75.75 0 0 0 0-1.5h-5.5a.75.75 0 0 0 0 1.5zM9 10a.75.75 0 0 1-.75.75h-5.5a.75.75 0 0 1 0-1.5h5.5A.75.75 0 0 1 9 10m8.25.75a.75.75 0 0 0 0-1.5h-1.5a.75.75 0 0 0 0 1.5zM14 10a2 2 0 1 0-4 0a2 2 0 0 0 4 0m-4 6.25a2 2 0 1 0-4 0a2 2 0 0 0 4 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdjustmentsVertical(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 13.5V3.75m0 9.75a1.5 1.5 0 0 1 0 3m0-3a1.5 1.5 0 0 0 0 3m0 3.75V16.5m12-3V3.75m0 9.75a1.5 1.5 0 0 1 0 3m0-3a1.5 1.5 0 0 0 0 3m0 3.75V16.5m-6-9V3.75m0 3.75a1.5 1.5 0 0 1 0 3m0-3a1.5 1.5 0 0 0 0 3m0 9.75V10.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdjustmentsVerticalSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.25 13.25V7.5h1.5v5.75a.75.75 0 0 1-1.5 0m1.5-10.5V5h.75a.75.75 0 0 1 0 1.5h-3a.75.75 0 0 1 0-1.5h.75V2.75a.75.75 0 0 1 1.5 0M2.25 9.5a.75.75 0 0 0 0 1.5h3a.75.75 0 0 0 0-1.5H4.5V2.75a.75.75 0 0 0-1.5 0V9.5zm7.75.75a.75.75 0 0 1 .75-.75h.75V2.75a.75.75 0 0 1 1.5 0V9.5h.75a.75.75 0 0 1 0 1.5h-3a.75.75 0 0 1-.75-.75M3 12v1.25a.75.75 0 0 0 1.5 0V12zm8.5 1.25V12H13v1.25a.75.75 0 0 1-1.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdjustmentsVerticalSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 12a.75.75 0 0 1-.75-.75v-7.5a.75.75 0 1 1 1.5 0v7.5A.75.75 0 0 1 6 12m12 0a.75.75 0 0 1-.75-.75v-7.5a.75.75 0 0 1 1.5 0v7.5A.75.75 0 0 1 18 12M6.75 20.25v-1.5a.75.75 0 0 0-1.5 0v1.5a.75.75 0 0 0 1.5 0m12-1.5v1.5a.75.75 0 0 1-1.5 0v-1.5a.75.75 0 0 1 1.5 0m-6-13.5v-1.5a.75.75 0 0 0-1.5 0v1.5a.75.75 0 0 0 1.5 0M12 21a.75.75 0 0 1-.75-.75v-7.5a.75.75 0 0 1 1.5 0v7.5A.75.75 0 0 1 12 21m-8.25-6a2.25 2.25 0 1 0 4.5 0a2.25 2.25 0 0 0-4.5 0M12 11.25a2.25 2.25 0 1 1 0-4.5a2.25 2.25 0 0 1 0 4.5M15.75 15a2.25 2.25 0 1 0 4.5 0a2.25 2.25 0 0 0-4.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdjustmentsVerticalTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2.75a.75.75 0 0 0-1.5 0v5.5a.75.75 0 0 0 1.5 0zm0 13a.75.75 0 0 0-1.5 0v1.5a.75.75 0 0 0 1.5 0zM3.75 15a.75.75 0 0 1 .75.75v1.5a.75.75 0 0 1-1.5 0v-1.5a.75.75 0 0 1 .75-.75M4.5 2.75a.75.75 0 0 0-1.5 0v5.5a.75.75 0 0 0 1.5 0zM10 11a.75.75 0 0 1 .75.75v5.5a.75.75 0 0 1-1.5 0v-5.5A.75.75 0 0 1 10 11m.75-8.25a.75.75 0 0 0-1.5 0v1.5a.75.75 0 0 0 1.5 0zM10 6a2 2 0 1 0 0 4a2 2 0 0 0 0-4m-6.25 4a2 2 0 1 0 0 4a2 2 0 0 0 0-4m12.5 0a2 2 0 1 0 0 4a2 2 0 0 0 0-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArchiveBox(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20.25 7.5l-.625 10.632a2.25 2.25 0 0 1-2.247 2.118H6.622a2.25 2.25 0 0 1-2.247-2.118L3.75 7.5M10 11.25h4M3.375 7.5h17.25c.621 0 1.125-.504 1.125-1.125v-1.5c0-.621-.504-1.125-1.125-1.125H3.375c-.621 0-1.125.504-1.125 1.125v1.5c0 .621.504 1.125 1.125 1.125"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArchiveBoxArrowDown(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20.25 7.5l-.625 10.632a2.25 2.25 0 0 1-2.247 2.118H6.622a2.25 2.25 0 0 1-2.247-2.118L3.75 7.5m8.25 3v6.75m0 0l-3-3m3 3l3-3M3.375 7.5h17.25c.621 0 1.125-.504 1.125-1.125v-1.5c0-.621-.504-1.125-1.125-1.125H3.375c-.621 0-1.125.504-1.125 1.125v1.5c0 .621.504 1.125 1.125 1.125"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArchiveBoxArrowDownSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M2 3a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1z"/><path fill-rule="evenodd" d="M13 6H3v6a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2zM8.75 7.75a.75.75 0 0 0-1.5 0v2.69L6.03 9.22a.75.75 0 0 0-1.06 1.06l2.5 2.5a.75.75 0 0 0 1.06 0l2.5-2.5a.75.75 0 1 0-1.06-1.06l-1.22 1.22z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArchiveBoxArrowDownSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M3.375 3C2.339 3 1.5 3.84 1.5 4.875v.75c0 1.036.84 1.875 1.875 1.875h17.25c1.035 0 1.875-.84 1.875-1.875v-.75C22.5 3.839 21.66 3 20.625 3z"/><path fill-rule="evenodd" d="m3.087 9l.54 9.176A3 3 0 0 0 6.62 21h10.757a3 3 0 0 0 2.995-2.824L20.913 9zM12 10.5a.75.75 0 0 1 .75.75v4.94l1.72-1.72a.75.75 0 1 1 1.06 1.06l-3 3a.75.75 0 0 1-1.06 0l-3-3a.75.75 0 1 1 1.06-1.06l1.72 1.72v-4.94a.75.75 0 0 1 .75-.75" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArchiveBoxArrowDownTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 3a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1zm0 4.5h16l-.811 7.71a2 2 0 0 1-1.99 1.79H4.802a2 2 0 0 1-1.99-1.79zM10 9a.75.75 0 0 1 .75.75v2.546l.943-1.048a.75.75 0 1 1 1.114 1.004l-2.25 2.5a.75.75 0 0 1-1.114 0l-2.25-2.5a.75.75 0 1 1 1.114-1.004l.943 1.048V9.75A.75.75 0 0 1 10 9" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArchiveBoxSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M3 2a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1z"/><path fill-rule="evenodd" d="M3 6h10v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2zm3 2.75A.75.75 0 0 1 6.75 8h2.5a.75.75 0 0 1 0 1.5h-2.5A.75.75 0 0 1 6 8.75" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArchiveBoxSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M3.375 3C2.339 3 1.5 3.84 1.5 4.875v.75c0 1.036.84 1.875 1.875 1.875h17.25c1.035 0 1.875-.84 1.875-1.875v-.75C22.5 3.839 21.66 3 20.625 3z"/><path fill-rule="evenodd" d="m3.087 9l.54 9.176A3 3 0 0 0 6.62 21h10.757a3 3 0 0 0 2.995-2.824L20.913 9zm6.163 3.75A.75.75 0 0 1 10 12h4a.75.75 0 0 1 0 1.5h-4a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArchiveBoxTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M2 3a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1z"/><path fill-rule="evenodd" d="M2 7.5h16l-.811 7.71a2 2 0 0 1-1.99 1.79H4.802a2 2 0 0 1-1.99-1.79zM7 11a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArchiveBoxXmark(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20.25 7.5l-.625 10.632a2.25 2.25 0 0 1-2.247 2.118H6.622a2.25 2.25 0 0 1-2.247-2.118L3.75 7.5m6 4.125l2.25 2.25m0 0l2.25 2.25M12 13.875l2.25-2.25M12 13.875l-2.25 2.25M3.375 7.5h17.25c.621 0 1.125-.504 1.125-1.125v-1.5c0-.621-.504-1.125-1.125-1.125H3.375c-.621 0-1.125.504-1.125 1.125v1.5c0 .621.504 1.125 1.125 1.125"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArchiveBoxXmarkSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M2 3a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1z"/><path fill-rule="evenodd" d="M13 6H3v6a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2zM5.72 7.47a.75.75 0 0 1 1.06 0L8 8.69l1.22-1.22a.75.75 0 1 1 1.06 1.06L9.06 9.75l1.22 1.22a.75.75 0 1 1-1.06 1.06L8 10.81l-1.22 1.22a.75.75 0 0 1-1.06-1.06l1.22-1.22l-1.22-1.22a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArchiveBoxXmarkSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M3.375 3C2.339 3 1.5 3.84 1.5 4.875v.75c0 1.036.84 1.875 1.875 1.875h17.25c1.035 0 1.875-.84 1.875-1.875v-.75C22.5 3.839 21.66 3 20.625 3z"/><path fill-rule="evenodd" d="m3.087 9l.54 9.176A3 3 0 0 0 6.62 21h10.757a3 3 0 0 0 2.995-2.824L20.913 9zm6.133 2.845a.75.75 0 0 1 1.06 0l1.72 1.72l1.72-1.72a.75.75 0 1 1 1.06 1.06l-1.72 1.72l1.72 1.72a.75.75 0 1 1-1.06 1.06L12 15.685l-1.72 1.72a.75.75 0 1 1-1.06-1.06l1.72-1.72l-1.72-1.72a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArchiveBoxXmarkTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M2 3a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1z"/><path fill-rule="evenodd" d="M2 7.5h16l-.811 7.71a2 2 0 0 1-1.99 1.79H4.802a2 2 0 0 1-1.99-1.79zm5.22 1.72a.75.75 0 0 1 1.06 0L10 10.94l1.72-1.72a.75.75 0 1 1 1.06 1.06L11.06 12l1.72 1.72a.75.75 0 1 1-1.06 1.06L10 13.06l-1.72 1.72a.75.75 0 0 1-1.06-1.06L8.94 12l-1.72-1.72a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDown(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.5 13.5L12 21m0 0l-7.5-7.5M12 21V3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownCircle(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 12.75l3 3m0 0l3-3m-3 3v-7.5M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownCircleSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14m.75-10.25a.75.75 0 0 0-1.5 0v4.69L6.03 8.22a.75.75 0 0 0-1.06 1.06l2.5 2.5a.75.75 0 0 0 1.06 0l2.5-2.5a.75.75 0 1 0-1.06-1.06L8.75 9.44z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownCircleSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75s9.75-4.365 9.75-9.75S17.385 2.25 12 2.25m-.53 14.03a.75.75 0 0 0 1.06 0l3-3a.75.75 0 1 0-1.06-1.06l-1.72 1.72V8.25a.75.75 0 0 0-1.5 0v5.69l-1.72-1.72a.75.75 0 0 0-1.06 1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownCircleTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 18a8 8 0 1 0 0-16a8 8 0 0 0 0 16m.75-11.25a.75.75 0 0 0-1.5 0v4.59L7.3 9.24a.75.75 0 0 0-1.1 1.02l3.25 3.5a.75.75 0 0 0 1.1 0l3.25-3.5a.75.75 0 1 0-1.1-1.02l-1.95 2.1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeft(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m19.5 4.5l-15 15m0 0h11.25m-11.25 0V8.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeftSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.78 4.22a.75.75 0 0 1 0 1.06L6.56 10.5h3.69a.75.75 0 0 1 0 1.5h-5.5a.75.75 0 0 1-.75-.75v-5.5a.75.75 0 0 1 1.5 0v3.69l5.22-5.22a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeftSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M20.03 3.97a.75.75 0 0 1 0 1.06L6.31 18.75h9.44a.75.75 0 0 1 0 1.5H4.5a.75.75 0 0 1-.75-.75V8.25a.75.75 0 0 1 1.5 0v9.44L18.97 3.97a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeftTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.78 5.22a.75.75 0 0 0-1.06 0L6.5 12.44V6.75a.75.75 0 0 0-1.5 0v7.5c0 .414.336.75.75.75h7.5a.75.75 0 0 0 0-1.5H7.56l7.22-7.22a.75.75 0 0 0 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownOnSquare(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 8.25H7.5a2.25 2.25 0 0 0-2.25 2.25v9a2.25 2.25 0 0 0 2.25 2.25h9a2.25 2.25 0 0 0 2.25-2.25v-9a2.25 2.25 0 0 0-2.25-2.25H15M9 12l3 3m0 0l3-3m-3 3V2.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownOnSquareSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1a.75.75 0 0 1 .75.75V5h-1.5V1.75A.75.75 0 0 1 8 1m-.75 4v4.44L6.03 8.22a.75.75 0 0 0-1.06 1.06l2.5 2.5a.75.75 0 0 0 1.06 0l2.5-2.5a.75.75 0 1 0-1.06-1.06L8.75 9.44V5H11a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownOnSquareSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.5a.75.75 0 0 1 .75.75V7.5h-1.5V2.25A.75.75 0 0 1 12 1.5m-.75 6v5.69l-1.72-1.72a.75.75 0 0 0-1.06 1.06l3 3a.75.75 0 0 0 1.06 0l3-3a.75.75 0 1 0-1.06-1.06l-1.72 1.72V7.5h3.75a3 3 0 0 1 3 3v9a3 3 0 0 1-3 3h-9a3 3 0 0 1-3-3v-9a3 3 0 0 1 3-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownOnSquareStack(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.5 7.5h-.75A2.25 2.25 0 0 0 4.5 9.75v7.5a2.25 2.25 0 0 0 2.25 2.25h7.5a2.25 2.25 0 0 0 2.25-2.25v-7.5a2.25 2.25 0 0 0-2.25-2.25h-.75m-6 3.75l3 3m0 0l3-3m-3 3V1.5m6 9h.75a2.25 2.25 0 0 1 2.25 2.25v7.5a2.25 2.25 0 0 1-2.25 2.25h-7.5a2.25 2.25 0 0 1-2.25-2.25v-.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownOnSquareStackSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M7 1a.75.75 0 0 1 .75.75V6h-1.5V1.75A.75.75 0 0 1 7 1m-.75 5v2.94L5.03 7.72a.75.75 0 0 0-1.06 1.06l2.5 2.5a.75.75 0 0 0 1.06 0l2.5-2.5a.75.75 0 1 0-1.06-1.06L7.75 8.94V6H10a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2z"/><path d="M4.268 14A2 2 0 0 0 6 15h6a2 2 0 0 0 2-2v-3a2 2 0 0 0-1-1.732V11a3 3 0 0 1-3 3z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownOnSquareStackSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M9.75 6.75h-3a3 3 0 0 0-3 3v7.5a3 3 0 0 0 3 3h7.5a3 3 0 0 0 3-3v-7.5a3 3 0 0 0-3-3h-3V1.5a.75.75 0 0 0-1.5 0zm0 0h1.5v5.69l1.72-1.72a.75.75 0 1 1 1.06 1.06l-3 3a.75.75 0 0 1-1.06 0l-3-3a.75.75 0 1 1 1.06-1.06l1.72 1.72z" clip-rule="evenodd"/><path d="M7.151 21.75a2.999 2.999 0 0 0 2.599 1.5h7.5a3 3 0 0 0 3-3v-7.5a3 3 0 0 0-1.5-2.599v7.099a4.5 4.5 0 0 1-4.5 4.5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownOnSquareStackTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1a.75.75 0 0 1 .75.75V6h-1.5V1.75A.75.75 0 0 1 8 1m-.75 5v3.296l-.943-1.048a.75.75 0 1 0-1.114 1.004l2.25 2.5a.75.75 0 0 0 1.114 0l2.25-2.5a.75.75 0 0 0-1.114-1.004L8.75 9.296V6h2A2.25 2.25 0 0 1 13 8.25v4.5A2.25 2.25 0 0 1 10.75 15h-5.5A2.25 2.25 0 0 1 3 12.75v-4.5A2.25 2.25 0 0 1 5.25 6zM7 16.75v-.25h3.75a3.75 3.75 0 0 0 3.75-3.75V10h.25A2.25 2.25 0 0 1 17 12.25v4.5A2.25 2.25 0 0 1 14.75 19h-5.5A2.25 2.25 0 0 1 7 16.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownOnSquareTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.75 7h-3v5.296l1.943-2.048a.75.75 0 0 1 1.114 1.004l-3.25 3.5a.75.75 0 0 1-1.114 0l-3.25-3.5a.75.75 0 1 1 1.114-1.004l1.943 2.048V7h1.5V1.75a.75.75 0 0 0-1.5 0V7h-3A2.25 2.25 0 0 0 4 9.25v7.5A2.25 2.25 0 0 0 6.25 19h7.5A2.25 2.25 0 0 0 16 16.75v-7.5A2.25 2.25 0 0 0 13.75 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRight(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4.5 4.5l15 15m0 0V8.25m0 11.25H8.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRightSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.22 4.22a.75.75 0 0 0 0 1.06l5.22 5.22H5.75a.75.75 0 0 0 0 1.5h5.5a.75.75 0 0 0 .75-.75v-5.5a.75.75 0 0 0-1.5 0v3.69L5.28 4.22a.75.75 0 0 0-1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRightSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.97 3.97a.75.75 0 0 1 1.06 0l13.72 13.72V8.25a.75.75 0 0 1 1.5 0V19.5a.75.75 0 0 1-.75.75H8.25a.75.75 0 0 1 0-1.5h9.44L3.97 5.03a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRightTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.28 5.22a.75.75 0 0 0-1.06 1.06l7.22 7.22H6.75a.75.75 0 0 0 0 1.5h7.5a.747.747 0 0 0 .75-.75v-7.5a.75.75 0 0 0-1.5 0v5.69z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 2a.75.75 0 0 1 .75.75v8.69l3.22-3.22a.75.75 0 1 1 1.06 1.06l-4.5 4.5a.75.75 0 0 1-1.06 0l-4.5-4.5a.75.75 0 0 1 1.06-1.06l3.22 3.22V2.75A.75.75 0 0 1 8 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25a.75.75 0 0 1 .75.75v16.19l6.22-6.22a.75.75 0 1 1 1.06 1.06l-7.5 7.5a.75.75 0 0 1-1.06 0l-7.5-7.5a.75.75 0 1 1 1.06-1.06l6.22 6.22V3a.75.75 0 0 1 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownTray(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 16.5v2.25A2.25 2.25 0 0 0 5.25 21h13.5A2.25 2.25 0 0 0 21 18.75V16.5M16.5 12L12 16.5m0 0L7.5 12m4.5 4.5V3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownTraySixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8.75 2.75a.75.75 0 0 0-1.5 0v5.69L5.03 6.22a.75.75 0 0 0-1.06 1.06l3.5 3.5a.75.75 0 0 0 1.06 0l3.5-3.5a.75.75 0 0 0-1.06-1.06L8.75 8.44z"/><path d="M3.5 9.75a.75.75 0 0 0-1.5 0v1.5A2.75 2.75 0 0 0 4.75 14h6.5A2.75 2.75 0 0 0 14 11.25v-1.5a.75.75 0 0 0-1.5 0v1.5c0 .69-.56 1.25-1.25 1.25h-6.5c-.69 0-1.25-.56-1.25-1.25z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownTraySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25a.75.75 0 0 1 .75.75v11.69l3.22-3.22a.75.75 0 1 1 1.06 1.06l-4.5 4.5a.75.75 0 0 1-1.06 0l-4.5-4.5a.75.75 0 1 1 1.06-1.06l3.22 3.22V3a.75.75 0 0 1 .75-.75m-9 13.5a.75.75 0 0 1 .75.75v2.25a1.5 1.5 0 0 0 1.5 1.5h13.5a1.5 1.5 0 0 0 1.5-1.5V16.5a.75.75 0 0 1 1.5 0v2.25a3 3 0 0 1-3 3H5.25a3 3 0 0 1-3-3V16.5a.75.75 0 0 1 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownTrayTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10.75 2.75a.75.75 0 0 0-1.5 0v8.614L6.295 8.235a.75.75 0 1 0-1.09 1.03l4.25 4.5a.75.75 0 0 0 1.09 0l4.25-4.5a.75.75 0 0 0-1.09-1.03l-2.955 3.129z"/><path d="M3.5 12.75a.75.75 0 0 0-1.5 0v2.5A2.75 2.75 0 0 0 4.75 18h10.5A2.75 2.75 0 0 0 18 15.25v-2.5a.75.75 0 0 0-1.5 0v2.5c0 .69-.56 1.25-1.25 1.25H4.75c-.69 0-1.25-.56-1.25-1.25z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 3a.75.75 0 0 1 .75.75v10.638l3.96-4.158a.75.75 0 1 1 1.08 1.04l-5.25 5.5a.75.75 0 0 1-1.08 0l-5.25-5.5a.75.75 0 1 1 1.08-1.04l3.96 4.158V3.75A.75.75 0 0 1 10 3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.5 19.5L3 12m0 0l7.5-7.5M3 12h18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftCircle(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m11.25 9l-3 3m0 0l3 3m-3-3h7.5M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftCircleSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 8a7 7 0 1 0 14 0A7 7 0 0 0 1 8m10.25.75a.75.75 0 0 0 0-1.5H6.56l1.22-1.22a.75.75 0 0 0-1.06-1.06l-2.5 2.5a.75.75 0 0 0 0 1.06l2.5 2.5a.75.75 0 1 0 1.06-1.06L6.56 8.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftCircleSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75s9.75-4.365 9.75-9.75S17.385 2.25 12 2.25m-4.28 9.22a.75.75 0 0 0 0 1.06l3 3a.75.75 0 1 0 1.06-1.06l-1.72-1.72h5.69a.75.75 0 0 0 0-1.5h-5.69l1.72-1.72a.75.75 0 0 0-1.06-1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftCircleTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#heroiconsArrowLeftCircle20Solid0)"><path fill="currentColor" fill-rule="evenodd" d="M10 18a8 8 0 1 0 0-16a8 8 0 0 0 0 16m3.25-7.25a.75.75 0 0 0 0-1.5H8.66l2.1-1.95a.75.75 0 1 0-1.02-1.1l-3.5 3.25a.75.75 0 0 0 0 1.1l3.5 3.25a.75.75 0 0 0 1.02-1.1l-2.1-1.95z" clip-rule="evenodd"/></g><defs><clipPath id="heroiconsArrowLeftCircle20Solid0"><path fill="#fff" d="M0 0h20v20H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftEndOnRectangle(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.75 9V5.25A2.25 2.25 0 0 0 13.5 3h-6a2.25 2.25 0 0 0-2.25 2.25v13.5A2.25 2.25 0 0 0 7.5 21h6a2.25 2.25 0 0 0 2.25-2.25V15M12 9l-3 3m0 0l3 3m-3-3h12.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftEndOnRectangleSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M4.75 2A2.75 2.75 0 0 0 2 4.75v6.5A2.75 2.75 0 0 0 4.75 14h3a2.75 2.75 0 0 0 2.75-2.75v-.5a.75.75 0 0 0-1.5 0v.5c0 .69-.56 1.25-1.25 1.25h-3c-.69 0-1.25-.56-1.25-1.25v-6.5c0-.69.56-1.25 1.25-1.25h3C8.44 3.5 9 4.06 9 4.75v.5a.75.75 0 0 0 1.5 0v-.5A2.75 2.75 0 0 0 7.75 2z"/><path d="M8.03 6.28a.75.75 0 0 0-1.06-1.06L4.72 7.47a.75.75 0 0 0 0 1.06l2.25 2.25a.75.75 0 1 0 1.06-1.06l-.97-.97h7.19a.75.75 0 0 0 0-1.5H7.06z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftEndOnRectangleSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 3.75A1.5 1.5 0 0 0 6 5.25v13.5a1.5 1.5 0 0 0 1.5 1.5h6a1.5 1.5 0 0 0 1.5-1.5V15a.75.75 0 0 1 1.5 0v3.75a3 3 0 0 1-3 3h-6a3 3 0 0 1-3-3V5.25a3 3 0 0 1 3-3h6a3 3 0 0 1 3 3V9A.75.75 0 0 1 15 9V5.25a1.5 1.5 0 0 0-1.5-1.5zm5.03 4.72a.75.75 0 0 1 0 1.06l-1.72 1.72h10.94a.75.75 0 0 1 0 1.5H10.81l1.72 1.72a.75.75 0 1 1-1.06 1.06l-3-3a.75.75 0 0 1 0-1.06l3-3a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftEndOnRectangleTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M3 4.25A2.25 2.25 0 0 1 5.25 2h5.5A2.25 2.25 0 0 1 13 4.25v2a.75.75 0 0 1-1.5 0v-2a.75.75 0 0 0-.75-.75h-5.5a.75.75 0 0 0-.75.75v11.5c0 .414.336.75.75.75h5.5a.75.75 0 0 0 .75-.75v-2a.75.75 0 0 1 1.5 0v2A2.25 2.25 0 0 1 10.75 18h-5.5A2.25 2.25 0 0 1 3 15.75z"/><path d="M19 10a.75.75 0 0 0-.75-.75H8.704l1.048-.943a.75.75 0 1 0-1.004-1.114l-2.5 2.25a.75.75 0 0 0 0 1.114l2.5 2.25a.75.75 0 1 0 1.004-1.114l-1.048-.943h9.546A.75.75 0 0 0 19 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftOnRectangle(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.75 9V5.25A2.25 2.25 0 0 0 13.5 3h-6a2.25 2.25 0 0 0-2.25 2.25v13.5A2.25 2.25 0 0 0 7.5 21h6a2.25 2.25 0 0 0 2.25-2.25V15M12 9l-3 3m0 0l3 3m-3-3h12.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftOnRectangleSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 3.75A1.5 1.5 0 0 0 6 5.25v13.5a1.5 1.5 0 0 0 1.5 1.5h6a1.5 1.5 0 0 0 1.5-1.5V15a.75.75 0 0 1 1.5 0v3.75a3 3 0 0 1-3 3h-6a3 3 0 0 1-3-3V5.25a3 3 0 0 1 3-3h6a3 3 0 0 1 3 3V9A.75.75 0 0 1 15 9V5.25a1.5 1.5 0 0 0-1.5-1.5zm5.03 4.72a.75.75 0 0 1 0 1.06l-1.72 1.72h10.94a.75.75 0 0 1 0 1.5H10.81l1.72 1.72a.75.75 0 1 1-1.06 1.06l-3-3a.75.75 0 0 1 0-1.06l3-3a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftOnRectangleTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M3 4.25A2.25 2.25 0 0 1 5.25 2h5.5A2.25 2.25 0 0 1 13 4.25v2a.75.75 0 0 1-1.5 0v-2a.75.75 0 0 0-.75-.75h-5.5a.75.75 0 0 0-.75.75v11.5c0 .414.336.75.75.75h5.5a.75.75 0 0 0 .75-.75v-2a.75.75 0 0 1 1.5 0v2A2.25 2.25 0 0 1 10.75 18h-5.5A2.25 2.25 0 0 1 3 15.75z"/><path d="M19 10a.75.75 0 0 0-.75-.75H8.704l1.048-.943a.75.75 0 1 0-1.004-1.114l-2.5 2.25a.75.75 0 0 0 0 1.114l2.5 2.25a.75.75 0 1 0 1.004-1.114l-1.048-.943h9.546A.75.75 0 0 0 19 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14 8a.75.75 0 0 1-.75.75H4.56l3.22 3.22a.75.75 0 1 1-1.06 1.06l-4.5-4.5a.75.75 0 0 1 0-1.06l4.5-4.5a.75.75 0 0 1 1.06 1.06L4.56 7.25h8.69A.75.75 0 0 1 14 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.03 3.97a.75.75 0 0 1 0 1.06l-6.22 6.22H21a.75.75 0 0 1 0 1.5H4.81l6.22 6.22a.75.75 0 1 1-1.06 1.06l-7.5-7.5a.75.75 0 0 1 0-1.06l7.5-7.5a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftStartOnRectangle(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.25 9V5.25A2.25 2.25 0 0 1 10.5 3h6a2.25 2.25 0 0 1 2.25 2.25v13.5A2.25 2.25 0 0 1 16.5 21h-6a2.25 2.25 0 0 1-2.25-2.25V15m-3 0l-3-3m0 0l3-3m-3 3H15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftStartOnRectangleSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14 4.75A2.75 2.75 0 0 0 11.25 2h-3A2.75 2.75 0 0 0 5.5 4.75v.5a.75.75 0 0 0 1.5 0v-.5c0-.69.56-1.25 1.25-1.25h3c.69 0 1.25.56 1.25 1.25v6.5c0 .69-.56 1.25-1.25 1.25h-3c-.69 0-1.25-.56-1.25-1.25v-.5a.75.75 0 0 0-1.5 0v.5A2.75 2.75 0 0 0 8.25 14h3A2.75 2.75 0 0 0 14 11.25zm-9.47.47a.75.75 0 0 0-1.06 0L1.22 7.47a.75.75 0 0 0 0 1.06l2.25 2.25a.75.75 0 1 0 1.06-1.06l-.97-.97h7.19a.75.75 0 0 0 0-1.5H3.56l.97-.97a.75.75 0 0 0 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftStartOnRectangleSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.5 3.75a1.5 1.5 0 0 1 1.5 1.5v13.5a1.5 1.5 0 0 1-1.5 1.5h-6a1.5 1.5 0 0 1-1.5-1.5V15a.75.75 0 0 0-1.5 0v3.75a3 3 0 0 0 3 3h6a3 3 0 0 0 3-3V5.25a3 3 0 0 0-3-3h-6a3 3 0 0 0-3 3V9A.75.75 0 1 0 9 9V5.25a1.5 1.5 0 0 1 1.5-1.5zM5.78 8.47a.75.75 0 0 0-1.06 0l-3 3a.75.75 0 0 0 0 1.06l3 3a.75.75 0 0 0 1.06-1.06l-1.72-1.72H15a.75.75 0 0 0 0-1.5H4.06l1.72-1.72a.75.75 0 0 0 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftStartOnRectangleTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M17 4.25A2.25 2.25 0 0 0 14.75 2h-5.5A2.25 2.25 0 0 0 7 4.25v2a.75.75 0 0 0 1.5 0v-2a.75.75 0 0 1 .75-.75h5.5a.75.75 0 0 1 .75.75v11.5a.75.75 0 0 1-.75.75h-5.5a.75.75 0 0 1-.75-.75v-2a.75.75 0 0 0-1.5 0v2A2.25 2.25 0 0 0 9.25 18h5.5A2.25 2.25 0 0 0 17 15.75z"/><path d="M14 10a.75.75 0 0 0-.75-.75H3.704l1.048-.943a.75.75 0 1 0-1.004-1.114l-2.5 2.25a.75.75 0 0 0 0 1.114l2.5 2.25a.75.75 0 1 0 1.004-1.114l-1.048-.943h9.546A.75.75 0 0 0 14 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M17 10a.75.75 0 0 1-.75.75H5.612l4.158 3.96a.75.75 0 1 1-1.04 1.08l-5.5-5.25a.75.75 0 0 1 0-1.08l5.5-5.25a.75.75 0 1 1 1.04 1.08L5.612 9.25H16.25A.75.75 0 0 1 17 10" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongDown(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.75 17.25L12 21m0 0l-3.75-3.75M12 21V3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongDownSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 2a.75.75 0 0 1 .75.75v8.69l1.22-1.22a.75.75 0 1 1 1.06 1.06l-2.5 2.5a.75.75 0 0 1-1.06 0l-2.5-2.5a.75.75 0 1 1 1.06-1.06l1.22 1.22V2.75A.75.75 0 0 1 8 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongDownSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25a.75.75 0 0 1 .75.75v16.19l2.47-2.47a.75.75 0 1 1 1.06 1.06l-3.75 3.75a.75.75 0 0 1-1.06 0l-3.75-3.75a.75.75 0 1 1 1.06-1.06l2.47 2.47V3a.75.75 0 0 1 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongDownTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 2a.75.75 0 0 1 .75.75v12.59l1.95-2.1a.75.75 0 1 1 1.1 1.02l-3.25 3.5a.75.75 0 0 1-1.1 0l-3.25-3.5a.75.75 0 1 1 1.1-1.02l1.95 2.1V2.75A.75.75 0 0 1 10 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongLeft(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.75 15.75L3 12m0 0l3.75-3.75M3 12h18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongLeftSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14 8a.75.75 0 0 1-.75.75H4.56l1.22 1.22a.75.75 0 1 1-1.06 1.06l-2.5-2.5a.75.75 0 0 1 0-1.06l2.5-2.5a.75.75 0 0 1 1.06 1.06L4.56 7.25h8.69A.75.75 0 0 1 14 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongLeftSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.28 7.72a.75.75 0 0 1 0 1.06l-2.47 2.47H21a.75.75 0 0 1 0 1.5H4.81l2.47 2.47a.75.75 0 1 1-1.06 1.06l-3.75-3.75a.75.75 0 0 1 0-1.06l3.75-3.75a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongLeftTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M18 10a.75.75 0 0 1-.75.75H4.66l2.1 1.95a.75.75 0 1 1-1.02 1.1l-3.5-3.25a.75.75 0 0 1 0-1.1l3.5-3.25a.75.75 0 1 1 1.02 1.1l-2.1 1.95h12.59A.75.75 0 0 1 18 10" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongRight(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.25 8.25L21 12m0 0l-3.75 3.75M21 12H3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongRightSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 8c0 .414.336.75.75.75h8.69l-1.22 1.22a.75.75 0 1 0 1.06 1.06l2.5-2.5a.75.75 0 0 0 0-1.06l-2.5-2.5a.75.75 0 1 0-1.06 1.06l1.22 1.22H2.75A.75.75 0 0 0 2 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongRightSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.72 7.72a.75.75 0 0 1 1.06 0l3.75 3.75a.75.75 0 0 1 0 1.06l-3.75 3.75a.75.75 0 1 1-1.06-1.06l2.47-2.47H3a.75.75 0 0 1 0-1.5h16.19l-2.47-2.47a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongRightTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 10a.75.75 0 0 1 .75-.75h12.59l-2.1-1.95a.75.75 0 1 1 1.02-1.1l3.5 3.25a.75.75 0 0 1 0 1.1l-3.5 3.25a.75.75 0 1 1-1.02-1.1l2.1-1.95H2.75A.75.75 0 0 1 2 10" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongUp(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.25 6.75L12 3m0 0l3.75 3.75M12 3v18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongUpSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 14a.75.75 0 0 0 .75-.75V4.56l1.22 1.22a.75.75 0 1 0 1.06-1.06l-2.5-2.5a.75.75 0 0 0-1.06 0l-2.5 2.5a.75.75 0 0 0 1.06 1.06l1.22-1.22v8.69c0 .414.336.75.75.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongUpSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.47 2.47a.75.75 0 0 1 1.06 0l3.75 3.75a.75.75 0 0 1-1.06 1.06l-2.47-2.47V21a.75.75 0 0 1-1.5 0V4.81L8.78 7.28a.75.75 0 0 1-1.06-1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongUpTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 18a.75.75 0 0 1-.75-.75V4.66L7.3 6.76a.75.75 0 0 1-1.1-1.02l3.25-3.5a.75.75 0 0 1 1.1 0l3.25 3.5a.75.75 0 1 1-1.1 1.02l-1.95-2.1v12.59A.75.75 0 0 1 10 18" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowPath(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0l3.181 3.183a8.25 8.25 0 0 0 13.803-3.7M4.031 9.865a8.25 8.25 0 0 1 13.803-3.7l3.181 3.182m0-4.991v4.99"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowPathRoundedSquare(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.5 12c0-1.232-.046-2.453-.138-3.662a4.006 4.006 0 0 0-3.7-3.7a48.678 48.678 0 0 0-7.324 0a4.006 4.006 0 0 0-3.7 3.7c-.017.22-.032.441-.046.662M19.5 12l3-3m-3 3l-3-3m-12 3c0 1.232.046 2.453.138 3.662a4.006 4.006 0 0 0 3.7 3.7a48.656 48.656 0 0 0 7.324 0a4.006 4.006 0 0 0 3.7-3.7c.017-.22.032-.441.046-.662M4.5 12l3 3m-3-3l-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowPathRoundedSquareSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 3.5c-.771 0-1.537.022-2.297.066a1.124 1.124 0 0 0-1.058 1.028l-.018.214a.75.75 0 1 1-1.495-.12l.018-.221a2.624 2.624 0 0 1 2.467-2.399a41.628 41.628 0 0 1 4.766 0a2.624 2.624 0 0 1 2.467 2.399c.056.662.097 1.329.122 2l.748-.748a.75.75 0 1 1 1.06 1.06l-2 2.001a.75.75 0 0 1-1.061 0l-2-1.999a.75.75 0 0 1 1.061-1.06l.689.688a39.89 39.89 0 0 0-.114-1.815a1.124 1.124 0 0 0-1.058-1.028A40.138 40.138 0 0 0 8 3.5M3.22 7.22a.75.75 0 0 1 1.061 0l2 2a.75.75 0 1 1-1.06 1.06l-.69-.69c.025.61.062 1.214.114 1.816c.048.56.496.996 1.058 1.028a40.112 40.112 0 0 0 4.594 0a1.124 1.124 0 0 0 1.058-1.028a39.2 39.2 0 0 0 .018-.219a.75.75 0 1 1 1.495.12l-.018.226a2.624 2.624 0 0 1-2.467 2.399a41.648 41.648 0 0 1-4.766 0a2.624 2.624 0 0 1-2.467-2.399a41.395 41.395 0 0 1-.122-2l-.748.748A.75.75 0 1 1 1.22 9.22z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowPathRoundedSquareSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 5.25c1.213 0 2.415.046 3.605.135a3.256 3.256 0 0 1 3.01 3.01c.044.583.077 1.17.1 1.759L17.03 8.47a.75.75 0 1 0-1.06 1.06l3 3a.75.75 0 0 0 1.06 0l3-3a.75.75 0 0 0-1.06-1.06l-1.752 1.751c-.023-.65-.06-1.296-.108-1.939a4.756 4.756 0 0 0-4.392-4.392a49.422 49.422 0 0 0-7.436 0A4.756 4.756 0 0 0 3.89 8.282a37.26 37.26 0 0 0-.046.672a.75.75 0 1 0 1.497.092c.013-.217.028-.434.044-.651a3.256 3.256 0 0 1 3.01-3.01c1.19-.09 2.392-.135 3.605-.135m-6.97 6.22a.75.75 0 0 0-1.06 0l-3 3a.75.75 0 1 0 1.06 1.06l1.752-1.751c.023.65.06 1.296.108 1.939a4.756 4.756 0 0 0 4.392 4.392a49.413 49.413 0 0 0 7.436 0a4.756 4.756 0 0 0 4.392-4.392c.017-.223.032-.447.046-.672a.75.75 0 0 0-1.497-.092c-.013.217-.028.434-.044.651a3.256 3.256 0 0 1-3.01 3.01a47.953 47.953 0 0 1-7.21 0a3.256 3.256 0 0 1-3.01-3.01a47.759 47.759 0 0 1-.1-1.759L6.97 15.53a.75.75 0 0 0 1.06-1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowPathRoundedSquareTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 4.5c1.215 0 2.417.055 3.604.162a.68.68 0 0 1 .615.597c.124 1.038.208 2.088.25 3.15l-1.689-1.69a.75.75 0 0 0-1.06 1.061l2.999 3a.75.75 0 0 0 1.06 0l3.001-3a.75.75 0 1 0-1.06-1.06l-1.748 1.747a41.31 41.31 0 0 0-.264-3.386a2.18 2.18 0 0 0-1.97-1.913a41.512 41.512 0 0 0-7.477 0a2.18 2.18 0 0 0-1.969 1.913a41.16 41.16 0 0 0-.16 1.61a.75.75 0 1 0 1.495.12c.041-.52.093-1.038.154-1.552a.68.68 0 0 1 .615-.597A40.012 40.012 0 0 1 10 4.5M5.281 9.22a.75.75 0 0 0-1.06 0l-3.001 3a.75.75 0 1 0 1.06 1.06l1.748-1.747c.042 1.141.13 2.27.264 3.386a2.18 2.18 0 0 0 1.97 1.913a41.533 41.533 0 0 0 7.477 0a2.18 2.18 0 0 0 1.969-1.913c.064-.534.117-1.071.16-1.61a.75.75 0 1 0-1.495-.12c-.041.52-.093 1.037-.154 1.552a.68.68 0 0 1-.615.597a40.013 40.013 0 0 1-7.208 0a.68.68 0 0 1-.615-.597a39.785 39.785 0 0 1-.25-3.15l1.689 1.69a.75.75 0 0 0 1.06-1.061z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowPathSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.836 2.477a.75.75 0 0 1 .75.75v3.182a.75.75 0 0 1-.75.75h-3.182a.75.75 0 0 1 0-1.5h1.37l-.84-.841a4.5 4.5 0 0 0-7.08.932a.75.75 0 0 1-1.3-.75a6 6 0 0 1 9.44-1.242l.842.84V3.227a.75.75 0 0 1 .75-.75m-.911 7.5A.75.75 0 0 1 13.199 11a6 6 0 0 1-9.44 1.241l-.84-.84v1.371a.75.75 0 0 1-1.5 0V9.591a.75.75 0 0 1 .75-.75H5.35a.75.75 0 0 1 0 1.5H3.98l.841.841a4.5 4.5 0 0 0 7.08-.932a.75.75 0 0 1 1.025-.273" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowPathSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.755 10.059a7.5 7.5 0 0 1 12.548-3.364l1.903 1.903h-3.183a.75.75 0 1 0 0 1.5h4.992a.75.75 0 0 0 .75-.75V4.356a.75.75 0 0 0-1.5 0v3.18l-1.9-1.9A9 9 0 0 0 3.306 9.67a.75.75 0 1 0 1.45.388m15.408 3.352a.75.75 0 0 0-.919.53a7.5 7.5 0 0 1-12.548 3.364l-1.902-1.903h3.183a.75.75 0 0 0 0-1.5H2.984a.75.75 0 0 0-.75.75v4.992a.75.75 0 0 0 1.5 0v-3.18l1.9 1.9a9 9 0 0 0 15.059-4.035a.75.75 0 0 0-.53-.918" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowPathTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.312 11.424a5.5 5.5 0 0 1-9.201 2.466l-.312-.311h2.433a.75.75 0 0 0 0-1.5H3.989a.75.75 0 0 0-.75.75v4.242a.75.75 0 0 0 1.5 0v-2.43l.31.31a7 7 0 0 0 11.712-3.138a.75.75 0 0 0-1.449-.39m1.23-3.723a.75.75 0 0 0 .219-.53V2.929a.75.75 0 0 0-1.5 0V5.36l-.31-.31A7 7 0 0 0 3.239 8.188a.75.75 0 1 0 1.448.389A5.5 5.5 0 0 1 13.89 6.11l.311.31h-2.432a.75.75 0 0 0 0 1.5h4.243a.75.75 0 0 0 .53-.219" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.5 4.5L21 12m0 0l-7.5 7.5M21 12H3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightCircle(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12.75 15l3-3m0 0l-3-3m3 3h-7.5M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightCircleSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 8A7 7 0 1 0 1 8a7 7 0 0 0 14 0M4.75 7.25a.75.75 0 0 0 0 1.5h4.69L8.22 9.97a.75.75 0 1 0 1.06 1.06l2.5-2.5a.75.75 0 0 0 0-1.06l-2.5-2.5a.75.75 0 0 0-1.06 1.06l1.22 1.22z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightCircleSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75s9.75-4.365 9.75-9.75S17.385 2.25 12 2.25m4.28 10.28a.75.75 0 0 0 0-1.06l-3-3a.75.75 0 1 0-1.06 1.06l1.72 1.72H8.25a.75.75 0 0 0 0 1.5h5.69l-1.72 1.72a.75.75 0 1 0 1.06 1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightCircleTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 18a8 8 0 1 0 0-16a8 8 0 0 0 0 16M6.75 9.25a.75.75 0 0 0 0 1.5h4.59l-2.1 1.95a.75.75 0 0 0 1.02 1.1l3.5-3.25a.75.75 0 0 0 0-1.1l-3.5-3.25a.75.75 0 1 0-1.02 1.1l2.1 1.95z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightEndOnRectangle(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.25 9V5.25A2.25 2.25 0 0 1 10.5 3h6a2.25 2.25 0 0 1 2.25 2.25v13.5A2.25 2.25 0 0 1 16.5 21h-6a2.25 2.25 0 0 1-2.25-2.25V15M12 9l3 3m0 0l-3 3m3-3H2.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightEndOnRectangleSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M11.25 2A2.75 2.75 0 0 1 14 4.75v6.5A2.75 2.75 0 0 1 11.25 14h-3a2.75 2.75 0 0 1-2.75-2.75v-.5a.75.75 0 0 1 1.5 0v.5c0 .69.56 1.25 1.25 1.25h3c.69 0 1.25-.56 1.25-1.25v-6.5c0-.69-.56-1.25-1.25-1.25h-3C7.56 3.5 7 4.06 7 4.75v.5a.75.75 0 0 1-1.5 0v-.5A2.75 2.75 0 0 1 8.25 2z"/><path d="M7.97 6.28a.75.75 0 0 1 1.06-1.06l2.25 2.25a.75.75 0 0 1 0 1.06l-2.25 2.25a.75.75 0 1 1-1.06-1.06l.97-.97H1.75a.75.75 0 0 1 0-1.5h7.19z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightEndOnRectangleSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.5 3.75a1.5 1.5 0 0 1 1.5 1.5v13.5a1.5 1.5 0 0 1-1.5 1.5h-6a1.5 1.5 0 0 1-1.5-1.5V15a.75.75 0 0 0-1.5 0v3.75a3 3 0 0 0 3 3h6a3 3 0 0 0 3-3V5.25a3 3 0 0 0-3-3h-6a3 3 0 0 0-3 3V9A.75.75 0 1 0 9 9V5.25a1.5 1.5 0 0 1 1.5-1.5zm-5.03 4.72a.75.75 0 0 0 0 1.06l1.72 1.72H2.25a.75.75 0 0 0 0 1.5h10.94l-1.72 1.72a.75.75 0 1 0 1.06 1.06l3-3a.75.75 0 0 0 0-1.06l-3-3a.75.75 0 0 0-1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightEndOnRectangleTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M17 4.25A2.25 2.25 0 0 0 14.75 2h-5.5A2.25 2.25 0 0 0 7 4.25v2a.75.75 0 0 0 1.5 0v-2a.75.75 0 0 1 .75-.75h5.5a.75.75 0 0 1 .75.75v11.5a.75.75 0 0 1-.75.75h-5.5a.75.75 0 0 1-.75-.75v-2a.75.75 0 0 0-1.5 0v2A2.25 2.25 0 0 0 9.25 18h5.5A2.25 2.25 0 0 0 17 15.75z"/><path d="M1 10a.75.75 0 0 1 .75-.75h9.546l-1.048-.943a.75.75 0 1 1 1.004-1.114l2.5 2.25a.75.75 0 0 1 0 1.114l-2.5 2.25a.75.75 0 1 1-1.004-1.114l1.048-.943H1.75A.75.75 0 0 1 1 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightOnRectangle(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.75 9V5.25A2.25 2.25 0 0 0 13.5 3h-6a2.25 2.25 0 0 0-2.25 2.25v13.5A2.25 2.25 0 0 0 7.5 21h6a2.25 2.25 0 0 0 2.25-2.25V15m3 0l3-3m0 0l-3-3m3 3H9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightOnRectangleSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 3.75A1.5 1.5 0 0 0 6 5.25v13.5a1.5 1.5 0 0 0 1.5 1.5h6a1.5 1.5 0 0 0 1.5-1.5V15a.75.75 0 0 1 1.5 0v3.75a3 3 0 0 1-3 3h-6a3 3 0 0 1-3-3V5.25a3 3 0 0 1 3-3h6a3 3 0 0 1 3 3V9A.75.75 0 0 1 15 9V5.25a1.5 1.5 0 0 0-1.5-1.5zm10.72 4.72a.75.75 0 0 1 1.06 0l3 3a.75.75 0 0 1 0 1.06l-3 3a.75.75 0 1 1-1.06-1.06l1.72-1.72H9a.75.75 0 0 1 0-1.5h10.94l-1.72-1.72a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightOnRectangleTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M3 4.25A2.25 2.25 0 0 1 5.25 2h5.5A2.25 2.25 0 0 1 13 4.25v2a.75.75 0 0 1-1.5 0v-2a.75.75 0 0 0-.75-.75h-5.5a.75.75 0 0 0-.75.75v11.5c0 .414.336.75.75.75h5.5a.75.75 0 0 0 .75-.75v-2a.75.75 0 0 1 1.5 0v2A2.25 2.25 0 0 1 10.75 18h-5.5A2.25 2.25 0 0 1 3 15.75z"/><path d="M6 10a.75.75 0 0 1 .75-.75h9.546l-1.048-.943a.75.75 0 1 1 1.004-1.114l2.5 2.25a.75.75 0 0 1 0 1.114l-2.5 2.25a.75.75 0 1 1-1.004-1.114l1.048-.943H6.75A.75.75 0 0 1 6 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 8a.75.75 0 0 1 .75-.75h8.69L8.22 4.03a.75.75 0 0 1 1.06-1.06l4.5 4.5a.75.75 0 0 1 0 1.06l-4.5 4.5a.75.75 0 0 1-1.06-1.06l3.22-3.22H2.75A.75.75 0 0 1 2 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.97 3.97a.75.75 0 0 1 1.06 0l7.5 7.5a.75.75 0 0 1 0 1.06l-7.5 7.5a.75.75 0 1 1-1.06-1.06l6.22-6.22H3a.75.75 0 0 1 0-1.5h16.19l-6.22-6.22a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightStartOnRectangle(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.75 9V5.25A2.25 2.25 0 0 0 13.5 3h-6a2.25 2.25 0 0 0-2.25 2.25v13.5A2.25 2.25 0 0 0 7.5 21h6a2.25 2.25 0 0 0 2.25-2.25V15m3 0l3-3m0 0l-3-3m3 3H9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightStartOnRectangleSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 4.75A2.75 2.75 0 0 1 4.75 2h3a2.75 2.75 0 0 1 2.75 2.75v.5a.75.75 0 0 1-1.5 0v-.5c0-.69-.56-1.25-1.25-1.25h-3c-.69 0-1.25.56-1.25 1.25v6.5c0 .69.56 1.25 1.25 1.25h3c.69 0 1.25-.56 1.25-1.25v-.5a.75.75 0 0 1 1.5 0v.5A2.75 2.75 0 0 1 7.75 14h-3A2.75 2.75 0 0 1 2 11.25zm9.47.47a.75.75 0 0 1 1.06 0l2.25 2.25a.75.75 0 0 1 0 1.06l-2.25 2.25a.75.75 0 1 1-1.06-1.06l.97-.97H5.25a.75.75 0 0 1 0-1.5h7.19l-.97-.97a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightStartOnRectangleSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 3.75A1.5 1.5 0 0 0 6 5.25v13.5a1.5 1.5 0 0 0 1.5 1.5h6a1.5 1.5 0 0 0 1.5-1.5V15a.75.75 0 0 1 1.5 0v3.75a3 3 0 0 1-3 3h-6a3 3 0 0 1-3-3V5.25a3 3 0 0 1 3-3h6a3 3 0 0 1 3 3V9A.75.75 0 0 1 15 9V5.25a1.5 1.5 0 0 0-1.5-1.5zm10.72 4.72a.75.75 0 0 1 1.06 0l3 3a.75.75 0 0 1 0 1.06l-3 3a.75.75 0 1 1-1.06-1.06l1.72-1.72H9a.75.75 0 0 1 0-1.5h10.94l-1.72-1.72a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightStartOnRectangleTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M3 4.25A2.25 2.25 0 0 1 5.25 2h5.5A2.25 2.25 0 0 1 13 4.25v2a.75.75 0 0 1-1.5 0v-2a.75.75 0 0 0-.75-.75h-5.5a.75.75 0 0 0-.75.75v11.5c0 .414.336.75.75.75h5.5a.75.75 0 0 0 .75-.75v-2a.75.75 0 0 1 1.5 0v2A2.25 2.25 0 0 1 10.75 18h-5.5A2.25 2.25 0 0 1 3 15.75z"/><path d="M6 10a.75.75 0 0 1 .75-.75h9.546l-1.048-.943a.75.75 0 1 1 1.004-1.114l2.5 2.25a.75.75 0 0 1 0 1.114l-2.5 2.25a.75.75 0 1 1-1.004-1.114l1.048-.943H6.75A.75.75 0 0 1 6 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 10a.75.75 0 0 1 .75-.75h10.638L10.23 5.29a.75.75 0 1 1 1.04-1.08l5.5 5.25a.75.75 0 0 1 0 1.08l-5.5 5.25a.75.75 0 1 1-1.04-1.08l4.158-3.96H3.75A.75.75 0 0 1 3 10" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSmallDown(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 4.5v15m0 0l6.75-6.75M12 19.5l-6.75-6.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSmallDownSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 3.75a.75.75 0 0 1 .75.75v13.19l5.47-5.47a.75.75 0 1 1 1.06 1.06l-6.75 6.75a.75.75 0 0 1-1.06 0l-6.75-6.75a.75.75 0 1 1 1.06-1.06l5.47 5.47V4.5a.75.75 0 0 1 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSmallDownTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 5a.75.75 0 0 1 .75.75v6.638l1.96-2.158a.75.75 0 1 1 1.08 1.04l-3.25 3.5a.75.75 0 0 1-1.08 0l-3.25-3.5a.75.75 0 1 1 1.08-1.04l1.96 2.158V5.75A.75.75 0 0 1 10 5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSmallLeft(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.5 12h-15m0 0l6.75 6.75M4.5 12l6.75-6.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSmallLeftSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M20.25 12a.75.75 0 0 1-.75.75H6.31l5.47 5.47a.75.75 0 1 1-1.06 1.06l-6.75-6.75a.75.75 0 0 1 0-1.06l6.75-6.75a.75.75 0 1 1 1.06 1.06l-5.47 5.47H19.5a.75.75 0 0 1 .75.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSmallLeftTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 10a.75.75 0 0 1-.75.75H7.612l2.158 1.96a.75.75 0 1 1-1.04 1.08l-3.5-3.25a.75.75 0 0 1 0-1.08l3.5-3.25a.75.75 0 1 1 1.04 1.08L7.612 9.25h6.638A.75.75 0 0 1 15 10" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSmallRight(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.5 12h15m0 0l-6.75-6.75M19.5 12l-6.75 6.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSmallRightSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.75 12a.75.75 0 0 1 .75-.75h13.19l-5.47-5.47a.75.75 0 0 1 1.06-1.06l6.75 6.75a.75.75 0 0 1 0 1.06l-6.75 6.75a.75.75 0 1 1-1.06-1.06l5.47-5.47H4.5a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSmallRightTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 10a.75.75 0 0 1 .75-.75h6.638L10.23 7.29a.75.75 0 1 1 1.04-1.08l3.5 3.25a.75.75 0 0 1 0 1.08l-3.5 3.25a.75.75 0 1 1-1.04-1.08l2.158-1.96H5.75A.75.75 0 0 1 5 10" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSmallUp(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 19.5v-15m0 0l-6.75 6.75M12 4.5l6.75 6.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSmallUpSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 20.25a.75.75 0 0 1-.75-.75V6.31l-5.47 5.47a.75.75 0 0 1-1.06-1.06l6.75-6.75a.75.75 0 0 1 1.06 0l6.75 6.75a.75.75 0 1 1-1.06 1.06l-5.47-5.47V19.5a.75.75 0 0 1-.75.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSmallUpTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 15a.75.75 0 0 1-.75-.75V7.612L7.29 9.77a.75.75 0 0 1-1.08-1.04l3.25-3.5a.75.75 0 0 1 1.08 0l3.25 3.5a.75.75 0 1 1-1.08 1.04l-1.96-2.158v6.638A.75.75 0 0 1 10 15" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTopRightOnSquare(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.5 6H5.25A2.25 2.25 0 0 0 3 8.25v10.5A2.25 2.25 0 0 0 5.25 21h10.5A2.25 2.25 0 0 0 18 18.75V10.5m-10.5 6L21 3m0 0h-5.25M21 3v5.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTopRightOnSquareSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M6.22 8.72a.75.75 0 0 0 1.06 1.06l5.22-5.22v1.69a.75.75 0 0 0 1.5 0v-3.5a.75.75 0 0 0-.75-.75h-3.5a.75.75 0 0 0 0 1.5h1.69z"/><path d="M3.5 6.75c0-.69.56-1.25 1.25-1.25H7A.75.75 0 0 0 7 4H4.75A2.75 2.75 0 0 0 2 6.75v4.5A2.75 2.75 0 0 0 4.75 14h4.5A2.75 2.75 0 0 0 12 11.25V9a.75.75 0 0 0-1.5 0v2.25c0 .69-.56 1.25-1.25 1.25h-4.5c-.69 0-1.25-.56-1.25-1.25z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTopRightOnSquareSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.75 2.25H21a.75.75 0 0 1 .75.75v5.25a.75.75 0 0 1-1.5 0V4.81L8.03 17.03a.75.75 0 0 1-1.06-1.06L19.19 3.75h-3.44a.75.75 0 0 1 0-1.5m-10.5 4.5a1.5 1.5 0 0 0-1.5 1.5v10.5a1.5 1.5 0 0 0 1.5 1.5h10.5a1.5 1.5 0 0 0 1.5-1.5V10.5a.75.75 0 0 1 1.5 0v8.25a3 3 0 0 1-3 3H5.25a3 3 0 0 1-3-3V8.25a3 3 0 0 1 3-3h8.25a.75.75 0 0 1 0 1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTopRightOnSquareTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M4.25 5.5a.75.75 0 0 0-.75.75v8.5c0 .414.336.75.75.75h8.5a.75.75 0 0 0 .75-.75v-4a.75.75 0 0 1 1.5 0v4A2.25 2.25 0 0 1 12.75 17h-8.5A2.25 2.25 0 0 1 2 14.75v-8.5A2.25 2.25 0 0 1 4.25 4h5a.75.75 0 0 1 0 1.5z"/><path d="M6.194 12.753a.75.75 0 0 0 1.06.053L16.5 4.44v2.81a.75.75 0 0 0 1.5 0v-4.5a.75.75 0 0 0-.75-.75h-4.5a.75.75 0 0 0 0 1.5h2.553l-9.056 8.194a.75.75 0 0 0-.053 1.06"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTrendingDown(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.25 6L9 12.75l4.286-4.286a11.948 11.948 0 0 1 4.306 6.43l.776 2.898m0 0l3.182-5.511m-3.182 5.51l-5.511-3.181"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTrendingDownSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.22 4.22a.75.75 0 0 1 1.06 0L6 7.94l2.761-2.762a.75.75 0 0 1 1.158.12a24.9 24.9 0 0 1 2.718 5.556l.729-1.261a.75.75 0 0 1 1.299.75l-1.591 2.755a.75.75 0 0 1-1.025.275l-2.756-1.591a.75.75 0 1 1 .75-1.3l1.097.634a23.417 23.417 0 0 0-1.984-4.211L6.53 9.53a.75.75 0 0 1-1.06 0L1.22 5.28a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTrendingDownSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.72 5.47a.75.75 0 0 1 1.06 0L9 11.69l3.756-3.756a.75.75 0 0 1 .985-.066a12.698 12.698 0 0 1 4.575 6.832l.308 1.149l2.277-3.943a.75.75 0 1 1 1.299.75l-3.182 5.51a.75.75 0 0 1-1.025.275l-5.511-3.181a.75.75 0 0 1 .75-1.3l3.943 2.277l-.308-1.149a11.194 11.194 0 0 0-3.528-5.617l-3.809 3.81a.75.75 0 0 1-1.06 0L1.72 6.53a.75.75 0 0 1 0-1.061" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTrendingDownTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.22 5.222a.75.75 0 0 1 1.06 0L7 9.942l3.768-3.769a.75.75 0 0 1 1.113.058a20.908 20.908 0 0 1 3.813 7.254l1.574-2.727a.75.75 0 0 1 1.3.75l-2.475 4.286a.75.75 0 0 1-1.025.275l-4.286-2.475a.75.75 0 0 1 .75-1.3l2.708 1.565a19.422 19.422 0 0 0-3.012-6.024L7.53 11.533a.75.75 0 0 1-1.06 0l-5.25-5.25a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTrendingUp(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.25 18L9 11.25l4.306 4.306a11.95 11.95 0 0 1 5.814-5.518l2.74-1.22m0 0l-5.94-2.281m5.94 2.28l-2.28 5.942"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTrendingUpSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.808 4.057a.75.75 0 0 1 .92-.527l3.116.849a.75.75 0 0 1 .528.915l-.823 3.121a.75.75 0 0 1-1.45-.382l.337-1.281a23.484 23.484 0 0 0-3.609 3.056a.75.75 0 0 1-1.07.01L6 8.06l-3.72 3.72a.75.75 0 1 1-1.06-1.061l4.25-4.25a.75.75 0 0 1 1.06 0l1.756 1.755a25.015 25.015 0 0 1 3.508-2.85l-1.46-.398a.75.75 0 0 1-.526-.92" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTrendingUpSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.22 6.268a.75.75 0 0 1 .968-.431l5.942 2.28a.75.75 0 0 1 .431.97l-2.28 5.94a.75.75 0 1 1-1.4-.537l1.63-4.251l-1.086.484a11.2 11.2 0 0 0-5.45 5.173a.75.75 0 0 1-1.199.19L9 12.312l-6.22 6.22a.75.75 0 0 1-1.06-1.061l6.75-6.75a.75.75 0 0 1 1.06 0l3.606 3.606a12.695 12.695 0 0 1 5.68-4.974l1.086-.483l-4.251-1.632a.75.75 0 0 1-.432-.97" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTrendingUpTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.577 4.878a.75.75 0 0 1 .919-.53l4.78 1.281a.75.75 0 0 1 .531.919l-1.281 4.78a.75.75 0 0 1-1.449-.387l.81-3.022a19.407 19.407 0 0 0-5.594 5.203a.75.75 0 0 1-1.139.093L7 10.06l-4.72 4.72a.75.75 0 0 1-1.06-1.061l5.25-5.25a.75.75 0 0 1 1.06 0l3.074 3.073a20.923 20.923 0 0 1 5.545-4.931l-3.042-.815a.75.75 0 0 1-.53-.919" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUp(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.5 10.5L12 3m0 0l7.5 7.5M12 3v18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpCircle(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15 11.25l-3-3m0 0l-3 3m3-3v7.5M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpCircleSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1a7 7 0 1 0 0 14A7 7 0 0 0 8 1m-.75 10.25a.75.75 0 0 0 1.5 0V6.56l1.22 1.22a.75.75 0 1 0 1.06-1.06l-2.5-2.5a.75.75 0 0 0-1.06 0l-2.5 2.5a.75.75 0 0 0 1.06 1.06l1.22-1.22z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpCircleSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75s9.75-4.365 9.75-9.75S17.385 2.25 12 2.25m.53 5.47a.75.75 0 0 0-1.06 0l-3 3a.75.75 0 1 0 1.06 1.06l1.72-1.72v5.69a.75.75 0 0 0 1.5 0v-5.69l1.72 1.72a.75.75 0 1 0 1.06-1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpCircleTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 18a8 8 0 1 0 0-16a8 8 0 0 0 0 16m-.75-4.75a.75.75 0 0 0 1.5 0V8.66l1.95 2.1a.75.75 0 1 0 1.1-1.02l-3.25-3.5a.75.75 0 0 0-1.1 0L6.2 9.74a.75.75 0 1 0 1.1 1.02l1.95-2.1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeft(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m19.5 19.5l-15-15m0 0v11.25m0-11.25h11.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeftSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.78 11.78a.75.75 0 0 0 0-1.06L6.56 5.5h3.69a.75.75 0 0 0 0-1.5h-5.5a.75.75 0 0 0-.75.75v5.5a.75.75 0 0 0 1.5 0V6.56l5.22 5.22a.75.75 0 0 0 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeftSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.25 6.31v9.44a.75.75 0 0 1-1.5 0V4.5a.75.75 0 0 1 .75-.75h11.25a.75.75 0 0 1 0 1.5H6.31l13.72 13.72a.75.75 0 1 1-1.06 1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeftTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.78 14.78a.75.75 0 0 1-1.06 0L6.5 7.56v5.69a.75.75 0 0 1-1.5 0v-7.5A.75.75 0 0 1 5.75 5h7.5a.75.75 0 0 1 0 1.5H7.56l7.22 7.22a.75.75 0 0 1 0 1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpOnSquare(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 8.25H7.5a2.25 2.25 0 0 0-2.25 2.25v9a2.25 2.25 0 0 0 2.25 2.25h9a2.25 2.25 0 0 0 2.25-2.25v-9a2.25 2.25 0 0 0-2.25-2.25H15m0-3l-3-3m0 0l-3 3m3-3V15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpOnSquareSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.75 6h-1.5V3.56L6.03 4.78a.75.75 0 0 1-1.06-1.06l2.5-2.5a.75.75 0 0 1 1.06 0l2.5 2.5a.75.75 0 1 1-1.06 1.06L8.75 3.56zH11a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h2.25v5.25a.75.75 0 0 0 1.5 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpOnSquareSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.47 1.72a.75.75 0 0 1 1.06 0l3 3a.75.75 0 0 1-1.06 1.06l-1.72-1.72V7.5h-1.5V4.06L9.53 5.78a.75.75 0 0 1-1.06-1.06zm-.22 5.78V15a.75.75 0 0 0 1.5 0V7.5h3.75a3 3 0 0 1 3 3v9a3 3 0 0 1-3 3h-9a3 3 0 0 1-3-3v-9a3 3 0 0 1 3-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpOnSquareStack(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.5 7.5h-.75A2.25 2.25 0 0 0 4.5 9.75v7.5a2.25 2.25 0 0 0 2.25 2.25h7.5a2.25 2.25 0 0 0 2.25-2.25v-7.5a2.25 2.25 0 0 0-2.25-2.25h-.75m0-3l-3-3m0 0l-3 3m3-3v11.25m6-2.25h.75a2.25 2.25 0 0 1 2.25 2.25v7.5a2.25 2.25 0 0 1-2.25 2.25h-7.5a2.25 2.25 0 0 1-2.25-2.25v-.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpOnSquareStackSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.268 14A2 2 0 0 0 7 15h4a2 2 0 0 0 2-2v-3a2 2 0 0 0-1-1.732V11a3 3 0 0 1-3 3zm.982-8h1.5V3.56l1.22 1.22a.75.75 0 1 0 1.06-1.06l-2.5-2.5a.75.75 0 0 0-1.06 0l-2.5 2.5a.75.75 0 0 0 1.06 1.06l1.22-1.22z"/><path fill="currentColor" d="M6.25 8.75a.75.75 0 0 0 1.5 0V6H9a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h1.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpOnSquareStackSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M9.97.97a.75.75 0 0 1 1.06 0l3 3a.75.75 0 0 1-1.06 1.06l-1.72-1.72v3.44h-1.5V3.31L8.03 5.03a.75.75 0 0 1-1.06-1.06zm-.22 5.78v6a.75.75 0 0 0 1.5 0v-6h3a3 3 0 0 1 3 3v7.5a3 3 0 0 1-3 3h-7.5a3 3 0 0 1-3-3v-7.5a3 3 0 0 1 3-3z"/><path d="M7.151 21.75a2.999 2.999 0 0 0 2.599 1.5h7.5a3 3 0 0 0 3-3v-7.5a3 3 0 0 0-1.5-2.599v7.099a4.5 4.5 0 0 1-4.5 4.5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpOnSquareStackTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.75 6h-2v4.25a.75.75 0 0 1-1.5 0V6h1.5V3.704l.943 1.048a.75.75 0 0 0 1.114-1.004l-2.25-2.5a.75.75 0 0 0-1.114 0l-2.25 2.5a.75.75 0 0 0 1.114 1.004l.943-1.048V6h-2A2.25 2.25 0 0 0 3 8.25v4.5A2.25 2.25 0 0 0 5.25 15h5.5A2.25 2.25 0 0 0 13 12.75v-4.5A2.25 2.25 0 0 0 10.75 6M7 16.75v-.25h3.75a3.75 3.75 0 0 0 3.75-3.75V10h.25A2.25 2.25 0 0 1 17 12.25v4.5A2.25 2.25 0 0 1 14.75 19h-5.5A2.25 2.25 0 0 1 7 16.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpOnSquareTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.75 7h-3V3.66l1.95 2.1a.75.75 0 1 0 1.1-1.02l-3.25-3.5a.75.75 0 0 0-1.1 0L6.2 4.74a.75.75 0 0 0 1.1 1.02l1.95-2.1V7h-3A2.25 2.25 0 0 0 4 9.25v7.5A2.25 2.25 0 0 0 6.25 19h7.5A2.25 2.25 0 0 0 16 16.75v-7.5A2.25 2.25 0 0 0 13.75 7m-3 0h-1.5v5.25a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRight(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4.5 19.5l15-15m0 0H8.25m11.25 0v11.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRightSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.22 11.78a.75.75 0 0 1 0-1.06L9.44 5.5H5.75a.75.75 0 0 1 0-1.5h5.5a.75.75 0 0 1 .75.75v5.5a.75.75 0 0 1-1.5 0V6.56l-5.22 5.22a.75.75 0 0 1-1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRightSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.25 3.75H19.5a.75.75 0 0 1 .75.75v11.25a.75.75 0 0 1-1.5 0V6.31L5.03 20.03a.75.75 0 0 1-1.06-1.06L17.69 5.25H8.25a.75.75 0 0 1 0-1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRightTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.22 14.78a.75.75 0 0 0 1.06 0l7.22-7.22v5.69a.75.75 0 0 0 1.5 0v-7.5a.75.75 0 0 0-.75-.75h-7.5a.75.75 0 0 0 0 1.5h5.69l-7.22 7.22a.75.75 0 0 0 0 1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 14a.75.75 0 0 1-.75-.75V4.56L4.03 7.78a.75.75 0 0 1-1.06-1.06l4.5-4.5a.75.75 0 0 1 1.06 0l4.5 4.5a.75.75 0 0 1-1.06 1.06L8.75 4.56v8.69A.75.75 0 0 1 8 14" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.47 2.47a.75.75 0 0 1 1.06 0l7.5 7.5a.75.75 0 1 1-1.06 1.06l-6.22-6.22V21a.75.75 0 0 1-1.5 0V4.81l-6.22 6.22a.75.75 0 1 1-1.06-1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpTray(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 16.5v2.25A2.25 2.25 0 0 0 5.25 21h13.5A2.25 2.25 0 0 0 21 18.75V16.5m-13.5-9L12 3m0 0l4.5 4.5M12 3v13.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpTraySixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M7.25 10.25a.75.75 0 0 0 1.5 0V4.56l2.22 2.22a.75.75 0 1 0 1.06-1.06l-3.5-3.5a.75.75 0 0 0-1.06 0l-3.5 3.5a.75.75 0 0 0 1.06 1.06l2.22-2.22z"/><path d="M3.5 9.75a.75.75 0 0 0-1.5 0v1.5A2.75 2.75 0 0 0 4.75 14h6.5A2.75 2.75 0 0 0 14 11.25v-1.5a.75.75 0 0 0-1.5 0v1.5c0 .69-.56 1.25-1.25 1.25h-6.5c-.69 0-1.25-.56-1.25-1.25z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpTraySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.47 2.47a.75.75 0 0 1 1.06 0l4.5 4.5a.75.75 0 0 1-1.06 1.06l-3.22-3.22V16.5a.75.75 0 0 1-1.5 0V4.81L8.03 8.03a.75.75 0 0 1-1.06-1.06zM3 15.75a.75.75 0 0 1 .75.75v2.25a1.5 1.5 0 0 0 1.5 1.5h13.5a1.5 1.5 0 0 0 1.5-1.5V16.5a.75.75 0 0 1 1.5 0v2.25a3 3 0 0 1-3 3H5.25a3 3 0 0 1-3-3V16.5a.75.75 0 0 1 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpTrayTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M9.25 13.25a.75.75 0 0 0 1.5 0V4.636l2.955 3.129a.75.75 0 0 0 1.09-1.03l-4.25-4.5a.75.75 0 0 0-1.09 0l-4.25 4.5a.75.75 0 1 0 1.09 1.03L9.25 4.636z"/><path d="M3.5 12.75a.75.75 0 0 0-1.5 0v2.5A2.75 2.75 0 0 0 4.75 18h10.5A2.75 2.75 0 0 0 18 15.25v-2.5a.75.75 0 0 0-1.5 0v2.5c0 .69-.56 1.25-1.25 1.25H4.75c-.69 0-1.25-.56-1.25-1.25z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 17a.75.75 0 0 1-.75-.75V5.612L5.29 9.77a.75.75 0 0 1-1.08-1.04l5.25-5.5a.75.75 0 0 1 1.08 0l5.25 5.5a.75.75 0 1 1-1.08 1.04l-3.96-4.158V16.25A.75.75 0 0 1 10 17" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUturnDown(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15 15l-6 6m0 0l-6-6m6 6V9a6 6 0 0 1 12 0v3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUturnDownSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.75 3.5A2.75 2.75 0 0 0 7 6.25v5.19l2.22-2.22a.75.75 0 1 1 1.06 1.06l-3.5 3.5a.75.75 0 0 1-1.06 0l-3.5-3.5a.75.75 0 1 1 1.06-1.06l2.22 2.22V6.25a4.25 4.25 0 0 1 8.5 0v1a.75.75 0 0 1-1.5 0v-1A2.75 2.75 0 0 0 9.75 3.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUturnDownSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 3.75A5.25 5.25 0 0 0 9.75 9v10.19l4.72-4.72a.75.75 0 1 1 1.06 1.06l-6 6a.75.75 0 0 1-1.06 0l-6-6a.75.75 0 1 1 1.06-1.06l4.72 4.72V9a6.75 6.75 0 0 1 13.5 0v3a.75.75 0 0 1-1.5 0V9c0-2.9-2.35-5.25-5.25-5.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUturnDownTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.232 12.208a.75.75 0 0 1 1.06.024l3.958 4.146V6.375a5.375 5.375 0 0 1 10.75 0V9.25a.75.75 0 0 1-1.5 0V6.375a3.875 3.875 0 0 0-7.75 0v10.003l3.957-4.146a.75.75 0 0 1 1.085 1.036l-5.25 5.5a.75.75 0 0 1-1.085 0l-5.25-5.5a.75.75 0 0 1 .025-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUturnLeft(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 15L3 9m0 0l6-6M3 9h12a6 6 0 0 1 0 12h-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUturnLeftSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 9.75A2.75 2.75 0 0 0 9.75 7H4.56l2.22 2.22a.75.75 0 1 1-1.06 1.06l-3.5-3.5a.75.75 0 0 1 0-1.06l3.5-3.5a.75.75 0 0 1 1.06 1.06L4.56 5.5h5.19a4.25 4.25 0 0 1 0 8.5h-1a.75.75 0 0 1 0-1.5h1a2.75 2.75 0 0 0 2.75-2.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUturnLeftSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.53 2.47a.75.75 0 0 1 0 1.06L4.81 8.25H15a6.75 6.75 0 0 1 0 13.5h-3a.75.75 0 0 1 0-1.5h3a5.25 5.25 0 1 0 0-10.5H4.81l4.72 4.72a.75.75 0 1 1-1.06 1.06l-6-6a.75.75 0 0 1 0-1.06l6-6a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUturnLeftTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.793 2.232a.75.75 0 0 1-.025 1.06L3.622 7.25h10.003a5.375 5.375 0 0 1 0 10.75H10.75a.75.75 0 0 1 0-1.5h2.875a3.875 3.875 0 0 0 0-7.75H3.622l4.146 3.957a.75.75 0 0 1-1.036 1.085l-5.5-5.25a.75.75 0 0 1 0-1.085l5.5-5.25a.75.75 0 0 1 1.06.025" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUturnRight(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15 15l6-6m0 0l-6-6m6 6H9a6 6 0 0 0 0 12h3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUturnRightSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 9.75A2.75 2.75 0 0 1 6.25 7h5.19L9.22 9.22a.75.75 0 1 0 1.06 1.06l3.5-3.5a.75.75 0 0 0 0-1.06l-3.5-3.5a.75.75 0 1 0-1.06 1.06l2.22 2.22H6.25a4.25 4.25 0 0 0 0 8.5h1a.75.75 0 0 0 0-1.5h-1A2.75 2.75 0 0 1 3.5 9.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUturnRightSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.47 2.47a.75.75 0 0 1 1.06 0l6 6a.75.75 0 0 1 0 1.06l-6 6a.75.75 0 1 1-1.06-1.06l4.72-4.72H9a5.25 5.25 0 1 0 0 10.5h3a.75.75 0 0 1 0 1.5H9a6.75 6.75 0 0 1 0-13.5h10.19l-4.72-4.72a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUturnRightTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.208 2.232a.75.75 0 0 0 .024 1.06l4.146 3.958H6.375a5.375 5.375 0 0 0 0 10.75H9.25a.75.75 0 0 0 0-1.5H6.375a3.875 3.875 0 0 1 0-7.75h10.003l-4.146 3.957a.75.75 0 0 0 1.036 1.085l5.5-5.25a.75.75 0 0 0 0-1.085l-5.5-5.25a.75.75 0 0 0-1.06.025" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUturnUp(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 9l6-6m0 0l6 6m-6-6v12a6 6 0 0 1-12 0v-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUturnUpSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.25 12.5A2.75 2.75 0 0 0 9 9.75V4.56L6.78 6.78a.75.75 0 0 1-1.06-1.06l3.5-3.5a.75.75 0 0 1 1.06 0l3.5 3.5a.75.75 0 0 1-1.06 1.06L10.5 4.56v5.19a4.25 4.25 0 0 1-8.5 0v-1a.75.75 0 0 1 1.5 0v1a2.75 2.75 0 0 0 2.75 2.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUturnUpSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M21.53 9.53a.75.75 0 0 1-1.06 0l-4.72-4.72V15a6.75 6.75 0 0 1-13.5 0v-3a.75.75 0 0 1 1.5 0v3a5.25 5.25 0 1 0 10.5 0V4.81L9.53 9.53a.75.75 0 0 1-1.06-1.06l6-6a.75.75 0 0 1 1.06 0l6 6a.75.75 0 0 1 0 1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUturnUpTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M17.768 7.793a.75.75 0 0 1-1.06-.025L12.75 3.622v10.003a5.375 5.375 0 0 1-10.75 0V10.75a.75.75 0 0 1 1.5 0v2.875a3.875 3.875 0 0 0 7.75 0V3.622L7.293 7.768a.75.75 0 0 1-1.086-1.036l5.25-5.5a.75.75 0 0 1 1.085 0l5.25 5.5a.75.75 0 0 1-.024 1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsPointingIn(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 9V4.5M9 9H4.5M9 9L3.75 3.75M9 15v4.5M9 15H4.5M9 15l-5.25 5.25M15 9h4.5M15 9V4.5M15 9l5.25-5.25M15 15h4.5M15 15v4.5m0-4.5l5.25 5.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsPointingInSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.22 2.22a.75.75 0 0 1 1.06 0L5.5 4.44V2.75a.75.75 0 0 1 1.5 0v3.5a.75.75 0 0 1-.75.75h-3.5a.75.75 0 0 1 0-1.5h1.69L2.22 3.28a.75.75 0 0 1 0-1.06m10.5 0a.75.75 0 1 1 1.06 1.06L11.56 5.5h1.69a.75.75 0 0 1 0 1.5h-3.5A.75.75 0 0 1 9 6.25v-3.5a.75.75 0 0 1 1.5 0v1.69zM2.75 9h3.5a.75.75 0 0 1 .75.75v3.5a.75.75 0 0 1-1.5 0v-1.69l-2.22 2.22a.75.75 0 0 1-1.06-1.06l2.22-2.22H2.75a.75.75 0 0 1 0-1.5M9 9.75A.75.75 0 0 1 9.75 9h3.5a.75.75 0 0 1 0 1.5h-1.69l2.22 2.22a.75.75 0 1 1-1.06 1.06l-2.22-2.22v1.69a.75.75 0 0 1-1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsPointingInSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.22 3.22a.75.75 0 0 1 1.06 0l3.97 3.97V4.5a.75.75 0 0 1 1.5 0V9a.75.75 0 0 1-.75.75H4.5a.75.75 0 0 1 0-1.5h2.69L3.22 4.28a.75.75 0 0 1 0-1.06m17.56 0a.75.75 0 0 1 0 1.06l-3.97 3.97h2.69a.75.75 0 0 1 0 1.5H15a.75.75 0 0 1-.75-.75V4.5a.75.75 0 0 1 1.5 0v2.69l3.97-3.97a.75.75 0 0 1 1.06 0M3.75 15a.75.75 0 0 1 .75-.75H9a.75.75 0 0 1 .75.75v4.5a.75.75 0 0 1-1.5 0v-2.69l-3.97 3.97a.75.75 0 0 1-1.06-1.06l3.97-3.97H4.5a.75.75 0 0 1-.75-.75m10.5 0a.75.75 0 0 1 .75-.75h4.5a.75.75 0 0 1 0 1.5h-2.69l3.97 3.97a.75.75 0 1 1-1.06 1.06l-3.97-3.97v2.69a.75.75 0 0 1-1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsPointingInTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.28 2.22a.75.75 0 0 0-1.06 1.06L5.44 6.5H2.75a.75.75 0 0 0 0 1.5h4.5A.75.75 0 0 0 8 7.25v-4.5a.75.75 0 0 0-1.5 0v2.69zm10.22.53a.75.75 0 0 0-1.5 0v4.5c0 .414.336.75.75.75h4.5a.75.75 0 0 0 0-1.5h-2.69l3.22-3.22a.75.75 0 0 0-1.06-1.06L13.5 5.44zM3.28 17.78l3.22-3.22v2.69a.75.75 0 0 0 1.5 0v-4.5a.75.75 0 0 0-.75-.75h-4.5a.75.75 0 0 0 0 1.5h2.69l-3.22 3.22a.75.75 0 1 0 1.06 1.06m10.22-3.22l3.22 3.22a.75.75 0 1 0 1.06-1.06l-3.22-3.22h2.69a.75.75 0 0 0 0-1.5h-4.5a.75.75 0 0 0-.75.75v4.5a.75.75 0 0 0 1.5 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsPointingOut(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 3.75v4.5m0-4.5h4.5m-4.5 0L9 9M3.75 20.25v-4.5m0 4.5h4.5m-4.5 0L9 15M20.25 3.75h-4.5m4.5 0v4.5m0-4.5L15 9m5.25 11.25h-4.5m4.5 0v-4.5m0 4.5L15 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsPointingOutSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.75 9a.75.75 0 0 1 .75.75v1.69l2.22-2.22a.75.75 0 0 1 1.06 1.06L4.56 12.5h1.69a.75.75 0 0 1 0 1.5h-3.5a.75.75 0 0 1-.75-.75v-3.5A.75.75 0 0 1 2.75 9m0-2a.75.75 0 0 0 .75-.75V4.56l2.22 2.22a.75.75 0 0 0 1.06-1.06L4.56 3.5h1.69a.75.75 0 0 0 0-1.5h-3.5a.75.75 0 0 0-.75.75v3.5c0 .414.336.75.75.75m10.5 2a.75.75 0 0 0-.75.75v1.69l-2.22-2.22a.75.75 0 1 0-1.06 1.06l2.22 2.22H9.75a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 .75-.75v-3.5a.75.75 0 0 0-.75-.75m0-2a.75.75 0 0 1-.75-.75V4.56l-2.22 2.22a.75.75 0 1 1-1.06-1.06l2.22-2.22H9.75a.75.75 0 0 1 0-1.5h3.5a.75.75 0 0 1 .75.75v3.5a.75.75 0 0 1-.75.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsPointingOutSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 3.75a.75.75 0 0 1 .75-.75h4.5a.75.75 0 0 1 .75.75v4.5a.75.75 0 0 1-1.5 0V5.56l-3.97 3.97a.75.75 0 1 1-1.06-1.06l3.97-3.97h-2.69a.75.75 0 0 1-.75-.75m-12 0A.75.75 0 0 1 3.75 3h4.5a.75.75 0 0 1 0 1.5H5.56l3.97 3.97a.75.75 0 0 1-1.06 1.06L4.5 5.56v2.69a.75.75 0 0 1-1.5 0zm11.47 11.78a.75.75 0 1 1 1.06-1.06l3.97 3.97v-2.69a.75.75 0 0 1 1.5 0v4.5a.75.75 0 0 1-.75.75h-4.5a.75.75 0 0 1 0-1.5h2.69zm-4.94-1.06a.75.75 0 0 1 0 1.06L5.56 19.5h2.69a.75.75 0 0 1 0 1.5h-4.5a.75.75 0 0 1-.75-.75v-4.5a.75.75 0 0 1 1.5 0v2.69l3.97-3.97a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsPointingOutTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.28 7.78l3.22-3.22v2.69a.75.75 0 0 0 1.5 0v-4.5a.75.75 0 0 0-.75-.75h-4.5a.75.75 0 0 0 0 1.5h2.69l-3.22 3.22a.75.75 0 0 0 1.06 1.06M2 17.25v-4.5a.75.75 0 0 1 1.5 0v2.69l3.22-3.22a.75.75 0 0 1 1.06 1.06L4.56 16.5h2.69a.75.75 0 0 1 0 1.5h-4.5a.747.747 0 0 1-.75-.75m10.22-3.97l3.22 3.22h-2.69a.75.75 0 0 0 0 1.5h4.5a.747.747 0 0 0 .75-.75v-4.5a.75.75 0 0 0-1.5 0v2.69l-3.22-3.22a.75.75 0 1 0-1.06 1.06M3.5 4.56l3.22 3.22a.75.75 0 0 0 1.06-1.06L4.56 3.5h2.69a.75.75 0 0 0 0-1.5h-4.5a.75.75 0 0 0-.75.75v4.5a.75.75 0 0 0 1.5 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsRightLeft(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.5 21L3 16.5m0 0L7.5 12M3 16.5h13.5m0-13.5L21 7.5m0 0L16.5 12M21 7.5H7.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsRightLeftSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.47 2.22a.75.75 0 0 1 1.06 0l2.25 2.25a.75.75 0 0 1 0 1.06l-2.25 2.25a.75.75 0 1 1-1.06-1.06l.97-.97H5.75a.75.75 0 0 1 0-1.5h5.69l-.97-.97a.75.75 0 0 1 0-1.06m-4.94 6a.75.75 0 0 1 0 1.06l-.97.97h5.69a.75.75 0 0 1 0 1.5H4.56l.97.97a.75.75 0 1 1-1.06 1.06l-2.25-2.25a.75.75 0 0 1 0-1.06l2.25-2.25a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsRightLeftSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.97 2.47a.75.75 0 0 1 1.06 0l4.5 4.5a.75.75 0 0 1 0 1.06l-4.5 4.5a.75.75 0 1 1-1.06-1.06l3.22-3.22H7.5a.75.75 0 0 1 0-1.5h11.69l-3.22-3.22a.75.75 0 0 1 0-1.06m-7.94 9a.75.75 0 0 1 0 1.06l-3.22 3.22H16.5a.75.75 0 0 1 0 1.5H4.81l3.22 3.22a.75.75 0 1 1-1.06 1.06l-4.5-4.5a.75.75 0 0 1 0-1.06l4.5-4.5a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsRightLeftTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.2 2.24a.75.75 0 0 0 .04 1.06l2.1 1.95H6.75a.75.75 0 0 0 0 1.5h8.59l-2.1 1.95a.75.75 0 1 0 1.02 1.1l3.5-3.25a.75.75 0 0 0 0-1.1l-3.5-3.25a.75.75 0 0 0-1.06.04m-6.4 8a.75.75 0 0 0-1.06-.04l-3.5 3.25a.75.75 0 0 0 0 1.1l3.5 3.25a.75.75 0 1 0 1.02-1.1l-2.1-1.95h8.59a.75.75 0 0 0 0-1.5H4.66l2.1-1.95a.75.75 0 0 0 .04-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsUpDown(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 7.5L7.5 3m0 0L12 7.5M7.5 3v13.5m13.5 0L16.5 21m0 0L12 16.5m4.5 4.5V7.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsUpDownSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.78 10.47a.75.75 0 0 1 0 1.06l-2.25 2.25a.75.75 0 0 1-1.06 0l-2.25-2.25a.75.75 0 1 1 1.06-1.06l.97.97V5.75a.75.75 0 0 1 1.5 0v5.69l.97-.97a.75.75 0 0 1 1.06 0M2.22 5.53a.75.75 0 0 1 0-1.06l2.25-2.25a.75.75 0 0 1 1.06 0l2.25 2.25a.75.75 0 0 1-1.06 1.06l-.97-.97v5.69a.75.75 0 0 1-1.5 0V4.56l-.97.97a.75.75 0 0 1-1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsUpDownSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.97 2.47a.75.75 0 0 1 1.06 0l4.5 4.5a.75.75 0 0 1-1.06 1.06L8.25 4.81V16.5a.75.75 0 0 1-1.5 0V4.81L3.53 8.03a.75.75 0 0 1-1.06-1.06zm9.53 4.28a.75.75 0 0 1 .75.75v11.69l3.22-3.22a.75.75 0 1 1 1.06 1.06l-4.5 4.5a.75.75 0 0 1-1.06 0l-4.5-4.5a.75.75 0 1 1 1.06-1.06l3.22 3.22V7.5a.75.75 0 0 1 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsUpDownTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.24 6.8a.75.75 0 0 0 1.06-.04l1.95-2.1v8.59a.75.75 0 0 0 1.5 0V4.66l1.95 2.1a.75.75 0 1 0 1.1-1.02l-3.25-3.5a.75.75 0 0 0-1.1 0L2.2 5.74a.75.75 0 0 0 .04 1.06m8 6.4a.75.75 0 0 0-.04 1.06l3.25 3.5a.75.75 0 0 0 1.1 0l3.25-3.5a.75.75 0 1 0-1.1-1.02l-1.95 2.1V6.75a.75.75 0 0 0-1.5 0v8.59l-1.95-2.1a.75.75 0 0 0-1.06-.04" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AtSymbol(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.5 12a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0m0 0c0 1.657 1.007 3 2.25 3S21 13.657 21 12a9 9 0 1 0-2.636 6.364M16.5 12V8.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AtSymbolSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.89 4.111a5.5 5.5 0 1 0 0 7.778a.75.75 0 1 1 1.06 1.061A7 7 0 1 1 15 8a2.5 2.5 0 0 1-4.083 1.935A3.5 3.5 0 1 1 11.5 8a1 1 0 0 0 2 0a5.48 5.48 0 0 0-1.61-3.889M10 8a2 2 0 1 0-4 0a2 2 0 0 0 4 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AtSymbolSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M17.834 6.166a8.25 8.25 0 1 0 0 11.668a.75.75 0 0 1 1.06 1.06c-3.807 3.808-9.98 3.808-13.788 0c-3.808-3.807-3.808-9.98 0-13.788c3.807-3.808 9.98-3.808 13.788 0A9.722 9.722 0 0 1 21.75 12c0 .975-.296 1.887-.809 2.571c-.514.685-1.28 1.179-2.191 1.179c-.904 0-1.666-.487-2.18-1.164a5.25 5.25 0 1 1-.82-6.26V8.25a.75.75 0 0 1 1.5 0V12c0 .682.208 1.27.509 1.671c.3.401.659.579.991.579c.332 0 .69-.178.991-.579c.3-.4.509-.99.509-1.671a8.222 8.222 0 0 0-2.416-5.834M15.75 12a3.75 3.75 0 1 0-7.5 0a3.75 3.75 0 0 0 7.5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AtSymbolTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.404 14.596A6.5 6.5 0 1 1 16.5 10a1.25 1.25 0 0 1-2.5 0a4 4 0 1 0-.571 2.06A2.75 2.75 0 0 0 18 10a8 8 0 1 0-2.343 5.657a.75.75 0 0 0-1.06-1.06a6.5 6.5 0 0 1-9.193 0M10 7.5a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Backspace(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 9.75L14.25 12m0 0l2.25 2.25M14.25 12l2.25-2.25M14.25 12L12 14.25m-2.58 4.92l-6.374-6.375a1.125 1.125 0 0 1 0-1.59L9.42 4.83c.21-.211.497-.33.795-.33H19.5a2.25 2.25 0 0 1 2.25 2.25v10.5a2.25 2.25 0 0 1-2.25 2.25h-9.284c-.298 0-.585-.119-.795-.33"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BackspaceSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.414 3c-.464 0-.909.184-1.237.513L1.22 7.47a.75.75 0 0 0 0 1.06l3.957 3.957A1.75 1.75 0 0 0 6.414 13h5.836A2.75 2.75 0 0 0 15 10.25v-4.5A2.75 2.75 0 0 0 12.25 3zM8.28 5.72a.75.75 0 0 0-1.06 1.06L8.44 8L7.22 9.22a.75.75 0 1 0 1.06 1.06L9.5 9.06l1.22 1.22a.75.75 0 1 0 1.06-1.06L10.56 8l1.22-1.22a.75.75 0 0 0-1.06-1.06L9.5 6.94z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BackspaceSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.515 10.674a1.875 1.875 0 0 0 0 2.652L8.89 19.7c.352.351.829.549 1.326.549H19.5a3 3 0 0 0 3-3V6.75a3 3 0 0 0-3-3h-9.284c-.497 0-.974.198-1.326.55zM12.53 9.22a.75.75 0 1 0-1.06 1.06L13.19 12l-1.72 1.72a.75.75 0 1 0 1.06 1.06l1.72-1.72l1.72 1.72a.75.75 0 1 0 1.06-1.06L15.31 12l1.72-1.72a.75.75 0 1 0-1.06-1.06l-1.72 1.72z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BackspaceTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.22 3.22A.75.75 0 0 1 7.75 3h9A2.25 2.25 0 0 1 19 5.25v9.5A2.25 2.25 0 0 1 16.75 17h-9a.75.75 0 0 1-.53-.22L.97 10.53a.75.75 0 0 1 0-1.06zm3.06 4a.75.75 0 1 0-1.06 1.06L10.94 10l-1.72 1.72a.75.75 0 1 0 1.06 1.06L12 11.06l1.72 1.72a.75.75 0 1 0 1.06-1.06L13.06 10l1.72-1.72a.75.75 0 0 0-1.06-1.06L12 8.94z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Backward(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 16.812c0 .863-.933 1.405-1.683.976l-7.108-4.061a1.125 1.125 0 0 1 0-1.954l7.108-4.061A1.125 1.125 0 0 1 21 8.689zm-9.75 0c0 .863-.933 1.405-1.683.976l-7.108-4.061a1.125 1.125 0 0 1 0-1.954l7.108-4.061a1.125 1.125 0 0 1 1.683.977z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BackwardSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 4.75a.75.75 0 0 0-1.107-.66l-6 3.25a.75.75 0 0 0 0 1.32l6 3.25a.75.75 0 0 0 1.107-.66V8.988l5.393 2.922A.75.75 0 0 0 15 11.25v-6.5a.75.75 0 0 0-1.107-.66L8.5 7.013z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BackwardSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.195 18.44c1.25.714 2.805-.189 2.805-1.629v-2.34l6.945 3.968c1.25.715 2.805-.188 2.805-1.628V8.69c0-1.44-1.555-2.343-2.805-1.628L12 11.029v-2.34c0-1.44-1.555-2.343-2.805-1.628l-7.108 4.061c-1.26.72-1.26 2.536 0 3.256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BackwardTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.712 4.818A1.5 1.5 0 0 1 10 6.095v2.972c.104-.13.234-.248.389-.343l6.323-3.906A1.5 1.5 0 0 1 19 6.095v7.81a1.5 1.5 0 0 1-2.288 1.276l-6.323-3.905a1.505 1.505 0 0 1-.389-.344v2.973a1.5 1.5 0 0 1-2.288 1.276l-6.323-3.905a1.5 1.5 0 0 1 0-2.552z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Banknotes(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.25 18.75a60.07 60.07 0 0 1 15.797 2.101c.727.198 1.453-.342 1.453-1.096V18.75M3.75 4.5v.75A.75.75 0 0 1 3 6h-.75m0 0v-.375c0-.621.504-1.125 1.125-1.125H20.25M2.25 6v9m18-10.5v.75c0 .414.336.75.75.75h.75m-1.5-1.5h.375c.621 0 1.125.504 1.125 1.125v9.75c0 .621-.504 1.125-1.125 1.125h-.375m1.5-1.5H21a.75.75 0 0 0-.75.75v.75m0 0H3.75m0 0h-.375a1.125 1.125 0 0 1-1.125-1.125V15m1.5 1.5v-.75A.75.75 0 0 0 3 15h-.75M15 10.5a3 3 0 1 1-6 0a3 3 0 0 1 6 0m3 0h.008v.008H18zm-12 0h.008v.008H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BanknotesSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M1 3a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1zm9 3a2 2 0 1 1-4 0a2 2 0 0 1 4 0m-6.25-.75a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5M11.5 6A.75.75 0 1 1 13 6a.75.75 0 0 1-1.5 0" clip-rule="evenodd"/><path d="M13 11.75a.75.75 0 0 0-1.5 0v.179c0 .15-.138.28-.306.255A65.277 65.277 0 0 0 1.75 11.5a.75.75 0 0 0 0 1.5c3.135 0 6.215.228 9.227.668A1.764 1.764 0 0 0 13 11.928z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BanknotesSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 7.5a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5"/><path fill-rule="evenodd" d="M1.5 4.875C1.5 3.839 2.34 3 3.375 3h17.25c1.035 0 1.875.84 1.875 1.875v9.75c0 1.036-.84 1.875-1.875 1.875H3.375A1.875 1.875 0 0 1 1.5 14.625zM8.25 9.75a3.75 3.75 0 1 1 7.5 0a3.75 3.75 0 0 1-7.5 0M18.75 9a.75.75 0 0 0-.75.75v.008c0 .414.336.75.75.75h.008a.75.75 0 0 0 .75-.75V9.75a.75.75 0 0 0-.75-.75zM4.5 9.75A.75.75 0 0 1 5.25 9h.008a.75.75 0 0 1 .75.75v.008a.75.75 0 0 1-.75.75H5.25a.75.75 0 0 1-.75-.75z" clip-rule="evenodd"/><path d="M2.25 18a.75.75 0 0 0 0 1.5c5.4 0 10.63.722 15.6 2.075c1.19.324 2.4-.558 2.4-1.82V18.75a.75.75 0 0 0-.75-.75z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BanknotesTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1zm12 4a3 3 0 1 1-6 0a3 3 0 0 1 6 0M4 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2m13-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0M1.75 14.5a.75.75 0 0 0 0 1.5c4.417 0 8.693.603 12.749 1.73c1.111.309 2.251-.512 2.251-1.696v-.784a.75.75 0 0 0-1.5 0v.784a.272.272 0 0 1-.35.25A49.043 49.043 0 0 0 1.75 14.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsArrowDown(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 4.5h14.25M3 9h9.75M3 13.5h9.75m4.5-4.5v12m0 0l-3.75-3.75M17.25 21L21 17.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsArrowDownSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 2.75A.75.75 0 0 1 2.75 2h9.5a.75.75 0 0 1 0 1.5h-9.5A.75.75 0 0 1 2 2.75m0 3.5a.75.75 0 0 1 .75-.75h5.5a.75.75 0 0 1 0 1.5h-5.5A.75.75 0 0 1 2 6.25m0 3.5A.75.75 0 0 1 2.75 9h3.5a.75.75 0 0 1 0 1.5h-3.5A.75.75 0 0 1 2 9.75m12.78 1.72a.75.75 0 0 1 0 1.06l-2.25 2.25a.75.75 0 0 1-1.06 0l-2.25-2.25a.75.75 0 1 1 1.06-1.06l.97.97V6.75a.75.75 0 0 1 1.5 0v5.69l.97-.97a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsArrowDownSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 4.5A.75.75 0 0 1 3 3.75h14.25a.75.75 0 0 1 0 1.5H3a.75.75 0 0 1-.75-.75m0 4.5A.75.75 0 0 1 3 8.25h9.75a.75.75 0 0 1 0 1.5H3A.75.75 0 0 1 2.25 9m15-.75A.75.75 0 0 1 18 9v10.19l2.47-2.47a.75.75 0 1 1 1.06 1.06l-3.75 3.75a.75.75 0 0 1-1.06 0l-3.75-3.75a.75.75 0 1 1 1.06-1.06l2.47 2.47V9a.75.75 0 0 1 .75-.75m-15 5.25a.75.75 0 0 1 .75-.75h9.75a.75.75 0 0 1 0 1.5H3a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsArrowDownTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 3.75A.75.75 0 0 1 2.75 3h11.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 3.75M2 7.5a.75.75 0 0 1 .75-.75h7.508a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 7.5M14 7a.75.75 0 0 1 .75.75v6.59l1.95-2.1a.75.75 0 1 1 1.1 1.02l-3.25 3.5a.75.75 0 0 1-1.1 0l-3.25-3.5a.75.75 0 1 1 1.1-1.02l1.95 2.1V7.75A.75.75 0 0 1 14 7M2 11.25a.75.75 0 0 1 .75-.75h4.562a.75.75 0 0 1 0 1.5H2.75a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsArrowUp(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 4.5h14.25M3 9h9.75M3 13.5h5.25m5.25-.75L17.25 9m0 0L21 12.75M17.25 9v12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsArrowUpSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 2.75A.75.75 0 0 1 2.75 2h9.5a.75.75 0 0 1 0 1.5h-9.5A.75.75 0 0 1 2 2.75m0 3.5a.75.75 0 0 1 .75-.75h5.5a.75.75 0 0 1 0 1.5h-5.5A.75.75 0 0 1 2 6.25m0 3.5A.75.75 0 0 1 2.75 9h3.5a.75.75 0 0 1 0 1.5h-3.5A.75.75 0 0 1 2 9.75m7.22-.22a.75.75 0 0 1 0-1.06l2.25-2.25a.75.75 0 0 1 1.06 0l2.25 2.25a.75.75 0 0 1-1.06 1.06l-.97-.97v5.69a.75.75 0 0 1-1.5 0V8.56l-.97.97a.75.75 0 0 1-1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsArrowUpSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 4.5A.75.75 0 0 1 3 3.75h14.25a.75.75 0 0 1 0 1.5H3a.75.75 0 0 1-.75-.75m14.47 3.97a.75.75 0 0 1 1.06 0l3.75 3.75a.75.75 0 1 1-1.06 1.06L18 10.81V21a.75.75 0 0 1-1.5 0V10.81l-2.47 2.47a.75.75 0 1 1-1.06-1.06zM2.25 9A.75.75 0 0 1 3 8.25h9.75a.75.75 0 0 1 0 1.5H3A.75.75 0 0 1 2.25 9m0 4.5a.75.75 0 0 1 .75-.75h5.25a.75.75 0 0 1 0 1.5H3a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsArrowUpTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 3.75A.75.75 0 0 1 2.75 3h11.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 3.75M2 7.5a.75.75 0 0 1 .75-.75h6.365a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 7.5M14 7a.75.75 0 0 1 .55.24l3.25 3.5a.75.75 0 1 1-1.1 1.02l-1.95-2.1v6.59a.75.75 0 0 1-1.5 0V9.66l-1.95 2.1a.75.75 0 1 1-1.1-1.02l3.25-3.5A.75.75 0 0 1 14 7M2 11.25a.75.75 0 0 1 .75-.75H7A.75.75 0 0 1 7 12H2.75a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsFour(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 5.25h16.5m-16.5 4.5h16.5m-16.5 4.5h16.5m-16.5 4.5h16.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsFourSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 2.75A.75.75 0 0 1 2.75 2h10.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 2.75m0 10.5a.75.75 0 0 1 .75-.75h10.5a.75.75 0 0 1 0 1.5H2.75a.75.75 0 0 1-.75-.75m0-7a.75.75 0 0 1 .75-.75h10.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 6.25m0 3.5A.75.75 0 0 1 2.75 9h10.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 9.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsFourSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 5.25a.75.75 0 0 1 .75-.75h16.5a.75.75 0 0 1 0 1.5H3.75A.75.75 0 0 1 3 5.25m0 4.5A.75.75 0 0 1 3.75 9h16.5a.75.75 0 0 1 0 1.5H3.75A.75.75 0 0 1 3 9.75m0 4.5a.75.75 0 0 1 .75-.75h16.5a.75.75 0 0 1 0 1.5H3.75a.75.75 0 0 1-.75-.75m0 4.5a.75.75 0 0 1 .75-.75h16.5a.75.75 0 0 1 0 1.5H3.75a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsFourTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 3.75A.75.75 0 0 1 2.75 3h14.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 3.75m0 4.167a.75.75 0 0 1 .75-.75h14.5a.75.75 0 0 1 0 1.5H2.75a.75.75 0 0 1-.75-.75m0 4.166a.75.75 0 0 1 .75-.75h14.5a.75.75 0 0 1 0 1.5H2.75a.75.75 0 0 1-.75-.75m0 4.167a.75.75 0 0 1 .75-.75h14.5a.75.75 0 0 1 0 1.5H2.75a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsThree(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsThreeBottomLeft(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25H12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsThreeBottomLeftSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 3.75A.75.75 0 0 1 2.75 3h10.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 3.75M2 8a.75.75 0 0 1 .75-.75h10.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 8m0 4.25a.75.75 0 0 1 .75-.75h4.5a.75.75 0 0 1 0 1.5h-4.5a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsThreeBottomLeftSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 6.75A.75.75 0 0 1 3.75 6h16.5a.75.75 0 0 1 0 1.5H3.75A.75.75 0 0 1 3 6.75M3 12a.75.75 0 0 1 .75-.75h16.5a.75.75 0 0 1 0 1.5H3.75A.75.75 0 0 1 3 12m0 5.25a.75.75 0 0 1 .75-.75H12a.75.75 0 0 1 0 1.5H3.75a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsThreeBottomLeftTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 4.75A.75.75 0 0 1 2.75 4h14.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 4.75m0 10.5a.75.75 0 0 1 .75-.75h7.5a.75.75 0 0 1 0 1.5h-7.5a.75.75 0 0 1-.75-.75M2 10a.75.75 0 0 1 .75-.75h14.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 10" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsThreeBottomRight(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 6.75h16.5M3.75 12h16.5M12 17.25h8.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsThreeBottomRightSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 3.75A.75.75 0 0 1 2.75 3h10.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 3.75M2 8a.75.75 0 0 1 .75-.75h10.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 8m6 4.25a.75.75 0 0 1 .75-.75h4.5a.75.75 0 0 1 0 1.5h-4.5a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsThreeBottomRightSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 6.75A.75.75 0 0 1 3.75 6h16.5a.75.75 0 0 1 0 1.5H3.75A.75.75 0 0 1 3 6.75M3 12a.75.75 0 0 1 .75-.75h16.5a.75.75 0 0 1 0 1.5H3.75A.75.75 0 0 1 3 12m8.25 5.25a.75.75 0 0 1 .75-.75h8.25a.75.75 0 0 1 0 1.5H12a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsThreeBottomRightTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 4.75A.75.75 0 0 1 2.75 4h14.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 4.75m7 10.5a.75.75 0 0 1 .75-.75h7.5a.75.75 0 0 1 0 1.5h-7.5a.75.75 0 0 1-.75-.75M2 10a.75.75 0 0 1 .75-.75h14.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 10" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsThreeCenterLeft(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 6.75h16.5M3.75 12H12m-8.25 5.25h16.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsThreeCenterLeftSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 3.75A.75.75 0 0 1 2.75 3h10.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 3.75M2 8a.75.75 0 0 1 .75-.75h4.5a.75.75 0 0 1 0 1.5h-4.5A.75.75 0 0 1 2 8m0 4.25a.75.75 0 0 1 .75-.75h10.5a.75.75 0 0 1 0 1.5H2.75a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsThreeCenterLeftSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 6.75A.75.75 0 0 1 3.75 6h16.5a.75.75 0 0 1 0 1.5H3.75A.75.75 0 0 1 3 6.75M3 12a.75.75 0 0 1 .75-.75H12a.75.75 0 0 1 0 1.5H3.75A.75.75 0 0 1 3 12m0 5.25a.75.75 0 0 1 .75-.75h16.5a.75.75 0 0 1 0 1.5H3.75a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsThreeCenterLeftTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 4.75A.75.75 0 0 1 2.75 4h14.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 4.75m0 10.5a.75.75 0 0 1 .75-.75h14.5a.75.75 0 0 1 0 1.5H2.75a.75.75 0 0 1-.75-.75M2 10a.75.75 0 0 1 .75-.75h7.5a.75.75 0 0 1 0 1.5h-7.5A.75.75 0 0 1 2 10" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsThreeSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 3.75A.75.75 0 0 1 2.75 3h10.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 3.75M2 8a.75.75 0 0 1 .75-.75h10.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 8m0 4.25a.75.75 0 0 1 .75-.75h10.5a.75.75 0 0 1 0 1.5H2.75a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsThreeSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 6.75A.75.75 0 0 1 3.75 6h16.5a.75.75 0 0 1 0 1.5H3.75A.75.75 0 0 1 3 6.75M3 12a.75.75 0 0 1 .75-.75h16.5a.75.75 0 0 1 0 1.5H3.75A.75.75 0 0 1 3 12m0 5.25a.75.75 0 0 1 .75-.75h16.5a.75.75 0 0 1 0 1.5H3.75a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsThreeTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 4.75A.75.75 0 0 1 2.75 4h14.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 4.75M2 10a.75.75 0 0 1 .75-.75h14.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 10m0 5.25a.75.75 0 0 1 .75-.75h14.5a.75.75 0 0 1 0 1.5H2.75a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsTwo(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 9h16.5m-16.5 6.75h16.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsTwoSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 4.75A.75.75 0 0 1 2.75 4h10.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 4.75m0 6.5a.75.75 0 0 1 .75-.75h10.5a.75.75 0 0 1 0 1.5H2.75a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsTwoSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 9a.75.75 0 0 1 .75-.75h16.5a.75.75 0 0 1 0 1.5H3.75A.75.75 0 0 1 3 9m0 6.75a.75.75 0 0 1 .75-.75h16.5a.75.75 0 0 1 0 1.5H3.75a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarsTwoTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 6.75A.75.75 0 0 1 2.75 6h14.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 6.75m0 6.5a.75.75 0 0 1 .75-.75h14.5a.75.75 0 0 1 0 1.5H2.75a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFifty(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 10.5h.375c.621 0 1.125.504 1.125 1.125v2.25c0 .621-.504 1.125-1.125 1.125H21M4.5 10.5h6.75V15H4.5zM3.75 18h15A2.25 2.25 0 0 0 21 15.75v-6a2.25 2.25 0 0 0-2.25-2.25h-15A2.25 2.25 0 0 0 1.5 9.75v6A2.25 2.25 0 0 0 3.75 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFiftySixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M1 6.25A2.25 2.25 0 0 1 3.25 4h8.5A2.25 2.25 0 0 1 14 6.25v.085a1.5 1.5 0 0 1 1 1.415v.5a1.5 1.5 0 0 1-1 1.415v.085A2.25 2.25 0 0 1 11.75 12h-8.5A2.25 2.25 0 0 1 1 9.75zm2.25-.75a.75.75 0 0 0-.75.75v3.5c0 .414.336.75.75.75h8.5a.75.75 0 0 0 .75-.75v-3.5a.75.75 0 0 0-.75-.75z" clip-rule="evenodd"/><path d="M4.75 7a.75.75 0 0 0-.75.75v.5c0 .414.336.75.75.75h2a.75.75 0 0 0 .75-.75v-.5A.75.75 0 0 0 6.75 7z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFiftySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M4.5 9.75a.75.75 0 0 0-.75.75V15c0 .414.336.75.75.75h6.75A.75.75 0 0 0 12 15v-4.5a.75.75 0 0 0-.75-.75z"/><path fill-rule="evenodd" d="M3.75 6.75a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h15a3 3 0 0 0 3-3v-.037c.856-.174 1.5-.93 1.5-1.838v-2.25c0-.907-.644-1.664-1.5-1.837V9.75a3 3 0 0 0-3-3zm15 1.5a1.5 1.5 0 0 1 1.5 1.5v6a1.5 1.5 0 0 1-1.5 1.5h-15a1.5 1.5 0 0 1-1.5-1.5v-6a1.5 1.5 0 0 1 1.5-1.5z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFiftyTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M4.75 8a.75.75 0 0 0-.75.75v2.5c0 .414.336.75.75.75H9.5a.75.75 0 0 0 .75-.75v-2.5A.75.75 0 0 0 9.5 8z"/><path fill-rule="evenodd" d="M3.25 5A2.25 2.25 0 0 0 1 7.25v5.5A2.25 2.25 0 0 0 3.25 15h12.5A2.25 2.25 0 0 0 18 12.75v-1.085a1.5 1.5 0 0 0 1-1.415v-.5a1.5 1.5 0 0 0-1-1.415V7.25A2.25 2.25 0 0 0 15.75 5zM2.5 7.25a.75.75 0 0 1 .75-.75h12.5a.75.75 0 0 1 .75.75v5.5a.75.75 0 0 1-.75.75H3.25a.75.75 0 0 1-.75-.75z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryOneHundred(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 10.5h.375c.621 0 1.125.504 1.125 1.125v2.25c0 .621-.504 1.125-1.125 1.125H21M4.5 10.5H18V15H4.5zM3.75 18h15A2.25 2.25 0 0 0 21 15.75v-6a2.25 2.25 0 0 0-2.25-2.25h-15A2.25 2.25 0 0 0 1.5 9.75v6A2.25 2.25 0 0 0 3.75 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryOneHundredSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M4 7.75A.75.75 0 0 1 4.75 7h5.5a.75.75 0 0 1 .75.75v.5a.75.75 0 0 1-.75.75h-5.5A.75.75 0 0 1 4 8.25z"/><path fill-rule="evenodd" d="M3.25 4A2.25 2.25 0 0 0 1 6.25v3.5A2.25 2.25 0 0 0 3.25 12h8.5A2.25 2.25 0 0 0 14 9.75v-.085a1.5 1.5 0 0 0 1-1.415v-.5a1.5 1.5 0 0 0-1-1.415V6.25A2.25 2.25 0 0 0 11.75 4zM2.5 6.25a.75.75 0 0 1 .75-.75h8.5a.75.75 0 0 1 .75.75v3.5a.75.75 0 0 1-.75.75h-8.5a.75.75 0 0 1-.75-.75z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryOneHundredSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.75 6.75a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h15a3 3 0 0 0 3-3v-.037c.856-.174 1.5-.93 1.5-1.838v-2.25c0-.907-.644-1.664-1.5-1.837V9.75a3 3 0 0 0-3-3zm15 1.5a1.5 1.5 0 0 1 1.5 1.5v6a1.5 1.5 0 0 1-1.5 1.5h-15a1.5 1.5 0 0 1-1.5-1.5v-6a1.5 1.5 0 0 1 1.5-1.5zM4.5 9.75a.75.75 0 0 0-.75.75V15c0 .414.336.75.75.75H18a.75.75 0 0 0 .75-.75v-4.5a.75.75 0 0 0-.75-.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryOneHundredTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M4.75 8a.75.75 0 0 0-.75.75v2.5c0 .414.336.75.75.75h9.5a.75.75 0 0 0 .75-.75v-2.5a.75.75 0 0 0-.75-.75z"/><path fill-rule="evenodd" d="M1 7.25A2.25 2.25 0 0 1 3.25 5h12.5A2.25 2.25 0 0 1 18 7.25v1.085a1.5 1.5 0 0 1 1 1.415v.5a1.5 1.5 0 0 1-1 1.415v1.085A2.25 2.25 0 0 1 15.75 15H3.25A2.25 2.25 0 0 1 1 12.75zm2.25-.75a.75.75 0 0 0-.75.75v5.5c0 .414.336.75.75.75h12.5a.75.75 0 0 0 .75-.75v-5.5a.75.75 0 0 0-.75-.75z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryZero(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 10.5h.375c.621 0 1.125.504 1.125 1.125v2.25c0 .621-.504 1.125-1.125 1.125H21M3.75 18h15A2.25 2.25 0 0 0 21 15.75v-6a2.25 2.25 0 0 0-2.25-2.25h-15A2.25 2.25 0 0 0 1.5 9.75v6A2.25 2.25 0 0 0 3.75 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryZeroSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 6.25A2.25 2.25 0 0 1 3.25 4h8.5A2.25 2.25 0 0 1 14 6.25v.085a1.5 1.5 0 0 1 1 1.415v.5a1.5 1.5 0 0 1-1 1.415v.085A2.25 2.25 0 0 1 11.75 12h-8.5A2.25 2.25 0 0 1 1 9.75zm2.25-.75a.75.75 0 0 0-.75.75v3.5c0 .414.336.75.75.75h8.5a.75.75 0 0 0 .75-.75v-3.5a.75.75 0 0 0-.75-.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryZeroSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M.75 9.75a3 3 0 0 1 3-3h15a3 3 0 0 1 3 3v.038c.856.173 1.5.93 1.5 1.837v2.25c0 .907-.644 1.664-1.5 1.838v.037a3 3 0 0 1-3 3h-15a3 3 0 0 1-3-3zm19.5 0a1.5 1.5 0 0 0-1.5-1.5h-15a1.5 1.5 0 0 0-1.5 1.5v6a1.5 1.5 0 0 0 1.5 1.5h15a1.5 1.5 0 0 0 1.5-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryZeroTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 7.25A2.25 2.25 0 0 1 3.25 5h12.5A2.25 2.25 0 0 1 18 7.25v1.085a1.5 1.5 0 0 1 1 1.415v.5a1.5 1.5 0 0 1-1 1.415v1.085A2.25 2.25 0 0 1 15.75 15H3.25A2.25 2.25 0 0 1 1 12.75zm2.25-.75a.75.75 0 0 0-.75.75v5.5c0 .414.336.75.75.75h12.5a.75.75 0 0 0 .75-.75v-5.5a.75.75 0 0 0-.75-.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Beaker(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.75 3.104v5.714a2.25 2.25 0 0 1-.659 1.591L5 14.5M9.75 3.104c-.251.023-.501.05-.75.082m.75-.082a24.301 24.301 0 0 1 4.5 0m0 0v5.714c0 .597.237 1.17.659 1.591L19.8 15.3M14.25 3.104c.251.023.501.05.75.082M19.8 15.3l-1.57.393A9.065 9.065 0 0 1 12 15a9.065 9.065 0 0 0-6.23-.693L5 14.5m14.8.8l1.402 1.402c1.232 1.232.65 3.318-1.067 3.611A48.309 48.309 0 0 1 12 21a48.25 48.25 0 0 1-8.135-.687c-1.718-.293-2.3-2.379-1.067-3.61L5 14.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BeakerSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 3.5v2.257c0 .597.237 1.17.659 1.591l2.733 2.733c.39.39.608.918.608 1.469a2.04 2.04 0 0 1-1.702 2.024C11.573 13.854 9.803 14 8 14s-3.573-.146-5.298-.426A2.04 2.04 0 0 1 1 11.55c0-.551.219-1.08.608-1.47l2.733-2.732A2.25 2.25 0 0 0 5 5.758V3.5h-.25a.75.75 0 0 1 0-1.5h6.5a.75.75 0 0 1 0 1.5zM6.5 5.757V3.5h3v2.257a3.75 3.75 0 0 0 1.098 2.652l.158.158a3.36 3.36 0 0 0-.075.034c-.424.2-.916.194-1.335-.016l-1.19-.595a4.943 4.943 0 0 0-2.07-.52A3.75 3.75 0 0 0 6.5 5.757" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BeakerSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.5 3.798v5.02a3 3 0 0 1-.879 2.121l-2.377 2.377a9.845 9.845 0 0 1 5.091 1.013a8.315 8.315 0 0 0 5.713.636l.285-.071l-3.954-3.955a3 3 0 0 1-.879-2.121v-5.02a23.614 23.614 0 0 0-3 0m4.5.138a.75.75 0 0 0 .093-1.495A24.837 24.837 0 0 0 12 2.25a25.048 25.048 0 0 0-3.093.191A.75.75 0 0 0 9 3.936v4.882a1.5 1.5 0 0 1-.44 1.06l-6.293 6.294c-1.62 1.621-.903 4.475 1.471 4.88c2.686.46 5.447.698 8.262.698c2.816 0 5.576-.239 8.262-.697c2.374-.406 3.092-3.26 1.47-4.881L15.44 9.879A1.5 1.5 0 0 1 15 8.818z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BeakerTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.5 3.528v4.644c0 .729-.29 1.428-.805 1.944l-1.217 1.217a8.75 8.75 0 0 1 3.55.62l.502.201a7.25 7.25 0 0 0 4.178.365l-2.403-2.403a2.75 2.75 0 0 1-.805-1.944V3.528a40.205 40.205 0 0 0-3 0m4.5.084l.19.015a.75.75 0 1 0 .12-1.495a41.364 41.364 0 0 0-6.62 0a.75.75 0 0 0 .12 1.495L7 3.612v4.56c0 .331-.132.649-.366.883L2.6 13.09c-1.496 1.496-.817 4.15 1.403 4.475C5.961 17.853 7.963 18 10 18s4.039-.148 5.997-.436c2.22-.325 2.9-2.979 1.403-4.475l-4.034-4.034A1.25 1.25 0 0 1 13 8.172z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.857 17.082a23.848 23.848 0 0 0 5.454-1.31A8.967 8.967 0 0 1 18 9.75V9A6 6 0 0 0 6 9v.75a8.967 8.967 0 0 1-2.312 6.022c1.733.64 3.56 1.085 5.455 1.31m5.714 0a24.255 24.255 0 0 1-5.714 0m5.714 0a3 3 0 1 1-5.714 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellAlert(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.857 17.082a23.848 23.848 0 0 0 5.454-1.31A8.967 8.967 0 0 1 18 9.75V9A6 6 0 0 0 6 9v.75a8.967 8.967 0 0 1-2.312 6.022c1.733.64 3.56 1.085 5.455 1.31m5.714 0a24.255 24.255 0 0 1-5.714 0m5.714 0a3 3 0 1 1-5.714 0M3.124 7.5A8.969 8.969 0 0 1 5.292 3m13.416 0a8.969 8.969 0 0 1 2.168 4.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellAlertSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M3.6 1.7A.75.75 0 1 0 2.4.799a6.978 6.978 0 0 0-1.123 2.247a.75.75 0 1 0 1.44.418c.187-.644.489-1.24.883-1.764m10-.901a.75.75 0 1 0-1.2.9a5.48 5.48 0 0 1 .883 1.765a.75.75 0 1 0 1.44-.418A6.978 6.978 0 0 0 13.6.799"/><path fill-rule="evenodd" d="M8 1a4 4 0 0 1 4 4v2.379c0 .398.158.779.44 1.06l1.267 1.268a1 1 0 0 1 .293.707V11a1 1 0 0 1-1 1h-2a3 3 0 1 1-6 0H3a1 1 0 0 1-1-1v-.586a1 1 0 0 1 .293-.707L3.56 8.44A1.5 1.5 0 0 0 4 7.38V5a4 4 0 0 1 4-4m0 12.5A1.5 1.5 0 0 1 6.5 12h3A1.5 1.5 0 0 1 8 13.5" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellAlertSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M5.85 3.5a.75.75 0 0 0-1.117-1a9.719 9.719 0 0 0-2.348 4.876a.75.75 0 0 0 1.479.248A8.219 8.219 0 0 1 5.85 3.5m13.417-1a.75.75 0 1 0-1.118 1a8.22 8.22 0 0 1 1.987 4.124a.75.75 0 0 0 1.48-.248a9.72 9.72 0 0 0-2.35-4.876"/><path fill-rule="evenodd" d="M12 2.25A6.75 6.75 0 0 0 5.25 9v.75a8.217 8.217 0 0 1-2.119 5.52a.75.75 0 0 0 .298 1.206c1.544.57 3.16.99 4.831 1.243a3.75 3.75 0 1 0 7.48 0a24.583 24.583 0 0 0 4.83-1.244a.75.75 0 0 0 .298-1.205a8.217 8.217 0 0 1-2.118-5.52V9A6.75 6.75 0 0 0 12 2.25M9.75 18c0-.034 0-.067.002-.1a25.05 25.05 0 0 0 4.496 0l.002.1a2.25 2.25 0 1 1-4.5 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellAlertTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M4.214 3.227a.75.75 0 0 0-1.156-.955a8.97 8.97 0 0 0-1.856 3.825a.75.75 0 0 0 1.466.316a7.47 7.47 0 0 1 1.546-3.186m12.728-.955a.75.75 0 0 0-1.157.955a7.47 7.47 0 0 1 1.547 3.186a.75.75 0 0 0 1.466-.316a8.971 8.971 0 0 0-1.856-3.825"/><path fill-rule="evenodd" d="M10 2a6 6 0 0 0-6 6c0 1.887-.454 3.665-1.257 5.234a.75.75 0 0 0 .515 1.076a32.91 32.91 0 0 0 3.256.508a3.5 3.5 0 0 0 6.972 0a32.903 32.903 0 0 0 3.256-.508a.75.75 0 0 0 .515-1.076A11.448 11.448 0 0 1 16 8a6 6 0 0 0-6-6m0 14.5a2 2 0 0 1-1.95-1.557a33.54 33.54 0 0 0 3.9 0A2 2 0 0 1 10 16.5" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 5a4 4 0 0 0-8 0v2.379a1.5 1.5 0 0 1-.44 1.06L2.294 9.707a1 1 0 0 0-.293.707V11a1 1 0 0 0 1 1h2a3 3 0 1 0 6 0h2a1 1 0 0 0 1-1v-.586a1 1 0 0 0-.293-.707L12.44 8.44A1.5 1.5 0 0 1 12 7.38zm-5.5 7a1.5 1.5 0 0 0 3 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellSlash(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.143 17.082a24.248 24.248 0 0 0 3.844.148m-3.844-.148a23.856 23.856 0 0 1-5.455-1.31a8.964 8.964 0 0 0 2.3-5.542m3.155 6.852a3 3 0 0 0 5.667 1.97m1.965-2.277L21 21m-4.225-4.225a23.81 23.81 0 0 0 3.536-1.003A8.967 8.967 0 0 1 18 9.75V9A6 6 0 0 0 6.53 6.53m10.245 10.245L6.53 6.53M3 3l3.53 3.53"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellSlashSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M4 7.379v-.904l6.743 6.743A3 3 0 0 1 5 12H3a1 1 0 0 1-1-1v-.586a1 1 0 0 1 .293-.707L3.56 8.44A1.5 1.5 0 0 0 4 7.38M6.5 12a1.5 1.5 0 0 0 3 0z" clip-rule="evenodd"/><path d="M14 11a.997.997 0 0 1-.096.429L4.92 2.446A4 4 0 0 1 12 5v2.379c0 .398.158.779.44 1.06l1.267 1.268a1 1 0 0 1 .293.707zM2.22 2.22a.75.75 0 0 1 1.06 0l10.5 10.5a.75.75 0 1 1-1.06 1.06L2.22 3.28a.75.75 0 0 1 0-1.06"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellSlashSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M3.53 2.47a.75.75 0 0 0-1.06 1.06l18 18a.75.75 0 1 0 1.06-1.06zm17.04 14.006c-.223.082-.448.161-.674.238L7.319 4.137A6.75 6.75 0 0 1 18.75 9v.75c0 2.123.8 4.057 2.118 5.52a.75.75 0 0 1-.297 1.206"/><path fill-rule="evenodd" d="M5.25 9c0-.184.007-.366.022-.546l10.384 10.384a3.751 3.751 0 0 1-7.396-1.119a24.585 24.585 0 0 1-4.831-1.244a.75.75 0 0 1-.298-1.205A8.217 8.217 0 0 0 5.25 9.75zm4.502 8.9a2.25 2.25 0 1 0 4.496 0a25.057 25.057 0 0 1-4.496 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellSlashTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 8c0-.26.017-.517.049-.77l7.722 7.723a33.56 33.56 0 0 1-3.722-.01a2 2 0 0 0 3.862.15l1.134 1.134a3.5 3.5 0 0 1-6.53-1.409a32.91 32.91 0 0 1-3.257-.508a.75.75 0 0 1-.515-1.076A11.448 11.448 0 0 0 4 8m13.266 5.9a.756.756 0 0 1-.068.116L6.389 3.207A6 6 0 0 1 16 8a11.466 11.466 0 0 0 1.258 5.234a.75.75 0 0 1 .01.666M3.28 2.22a.75.75 0 0 0-1.06 1.06l14.5 14.5a.75.75 0 1 0 1.06-1.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellSnooze(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.857 17.082a23.848 23.848 0 0 0 5.454-1.31A8.967 8.967 0 0 1 18 9.75V9A6 6 0 0 0 6 9v.75a8.967 8.967 0 0 1-2.312 6.022c1.733.64 3.56 1.085 5.455 1.31m5.714 0a24.255 24.255 0 0 1-5.714 0m5.714 0a3 3 0 1 1-5.714 0M10.5 8.25h3l-3 4.5h3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellSnoozeSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1a4 4 0 0 1 4 4v2.379c0 .398.158.779.44 1.06l1.267 1.268a1 1 0 0 1 .293.707V11a1 1 0 0 1-1 1h-2a3 3 0 1 1-6 0H3a1 1 0 0 1-1-1v-.586a1 1 0 0 1 .293-.707L3.56 8.44A1.5 1.5 0 0 0 4 7.38V5a4 4 0 0 1 4-4m0 12.5A1.5 1.5 0 0 1 6.5 12h3A1.5 1.5 0 0 1 8 13.5M6.75 4a.75.75 0 0 0 0 1.5h1.043L6.14 7.814A.75.75 0 0 0 6.75 9h2.5a.75.75 0 1 0 0-1.5H8.207L9.86 5.186A.75.75 0 0 0 9.25 4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellSnoozeSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25A6.75 6.75 0 0 0 5.25 9v.75a8.217 8.217 0 0 1-2.119 5.52a.75.75 0 0 0 .298 1.206c1.544.57 3.16.99 4.831 1.243a3.75 3.75 0 1 0 7.48 0a24.583 24.583 0 0 0 4.83-1.244a.75.75 0 0 0 .298-1.205a8.217 8.217 0 0 1-2.118-5.52V9A6.75 6.75 0 0 0 12 2.25M9.75 18c0-.034 0-.067.002-.1a25.05 25.05 0 0 0 4.496 0l.002.1a2.25 2.25 0 1 1-4.5 0m.75-10.5a.75.75 0 0 0 0 1.5h1.599l-2.223 3.334A.75.75 0 0 0 10.5 13.5h3a.75.75 0 0 0 0-1.5h-1.599l2.223-3.334A.75.75 0 0 0 13.5 7.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellSnoozeTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 8a6 6 0 1 1 12 0c0 1.887.454 3.665 1.257 5.234a.75.75 0 0 1-.515 1.076a32.903 32.903 0 0 1-3.256.508a3.5 3.5 0 0 1-6.972 0a32.91 32.91 0 0 1-3.256-.508a.75.75 0 0 1-.515-1.076A11.448 11.448 0 0 0 4 8m6 7c-.655 0-1.305-.02-1.95-.057a2 2 0 0 0 3.9 0A33.15 33.15 0 0 1 10 15M8.75 6a.75.75 0 0 0 0 1.5h1.043L8.14 9.814A.75.75 0 0 0 8.75 11h2.5a.75.75 0 0 0 0-1.5h-1.043l1.653-2.314A.75.75 0 0 0 11.25 6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.25 9a6.75 6.75 0 0 1 13.5 0v.75c0 2.123.8 4.057 2.118 5.52a.75.75 0 0 1-.297 1.206c-1.544.57-3.16.99-4.831 1.243a3.75 3.75 0 1 1-7.48 0a24.585 24.585 0 0 1-4.831-1.244a.75.75 0 0 1-.298-1.205A8.217 8.217 0 0 0 5.25 9.75zm4.502 8.9a2.25 2.25 0 1 0 4.496 0a25.057 25.057 0 0 1-4.496 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 2a6 6 0 0 0-6 6c0 1.887-.454 3.665-1.257 5.234a.75.75 0 0 0 .515 1.076a32.91 32.91 0 0 0 3.256.508a3.5 3.5 0 0 0 6.972 0a32.903 32.903 0 0 0 3.256-.508a.75.75 0 0 0 .515-1.076A11.448 11.448 0 0 1 16 8a6 6 0 0 0-6-6M8.05 14.943a33.54 33.54 0 0 0 3.9 0a2 2 0 0 1-3.9 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bold(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6 3.75h7.125a4.125 4.125 0 1 1 0 8.25H6z"/><path d="M7.5 4.5h4.875a3.375 3.375 0 1 1 0 6.75H7.5zM6 11.25h7.5a4.5 4.5 0 1 1 0 9H6z"/><path d="M7.5 12h5.25a3.75 3.75 0 1 1 0 7.5H7.5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoldSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.25 3.75A.75.75 0 0 1 6 3h7.125a4.875 4.875 0 0 1 3.346 8.42A5.25 5.25 0 0 1 13.5 21H5.999a.75.75 0 0 1-.75-.75zm7.125 6.75a2.625 2.625 0 0 0 0-5.25H8.25v5.25zM8.25 18.75h4.5a3 3 0 1 0 0-6h-4.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoldTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 3.222C4 2.547 4.547 2 5.222 2H10.5a4.5 4.5 0 0 1 3.475 7.359A4.75 4.75 0 0 1 11.25 18H5.333A1.333 1.333 0 0 1 4 16.667zM6.5 8.5v-4h4a2 2 0 1 1 0 4zm0 2.5v4.5h4.75a2.25 2.25 0 0 0 0-4.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bolt(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3.75 13.5l10.5-11.25L12 10.5h8.25L9.75 21.75L12 13.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoltSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.58 1.077a.75.75 0 0 1 .405.82L9.165 6h4.085a.75.75 0 0 1 .567 1.241l-6.5 7.5a.75.75 0 0 1-1.302-.638L6.835 10H2.75a.75.75 0 0 1-.567-1.241l6.5-7.5a.75.75 0 0 1 .897-.182" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoltSlash(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.412 15.655L9.75 21.75l3.745-4.012M9.257 13.5H3.75l2.659-2.849m2.048-2.194L14.25 2.25L12 10.5h8.25l-4.707 5.043M8.457 8.457L3 3m5.457 5.457l7.086 7.086m0 0L21 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoltSlashSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.58 1.077a.75.75 0 0 1 .405.82L9.165 6h4.085a.75.75 0 0 1 .567 1.241l-1.904 2.197L6.385 3.91l2.298-2.65a.75.75 0 0 1 .897-.182M4.087 6.562l5.528 5.528l-2.298 2.651a.75.75 0 0 1-1.302-.638L6.835 10H2.75a.75.75 0 0 1-.567-1.241zM2.22 2.22a.75.75 0 0 1 1.06 0l10.5 10.5a.75.75 0 1 1-1.06 1.06L2.22 3.28a.75.75 0 0 1 0-1.06"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoltSlashSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.798 11.012l-3.188 3.416L9.462 6.28l4.24-4.542a.75.75 0 0 1 1.272.71L12.982 9.75h7.268a.75.75 0 0 1 .548 1.262M3.202 12.988L6.39 9.572l8.148 8.148l-4.24 4.542a.75.75 0 0 1-1.272-.71l1.992-7.302H3.75a.75.75 0 0 1-.548-1.262M3.53 2.47a.75.75 0 0 0-1.06 1.06l18 18a.75.75 0 1 0 1.06-1.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoltSlashTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M2.22 2.22a.75.75 0 0 1 1.06 0l14.5 14.5a.75.75 0 1 1-1.06 1.06L2.22 3.28a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/><path d="M4.73 7.912L2.191 10.75A.75.75 0 0 0 2.75 12h6.068zm4.503 4.503l-1.216 5.678a.75.75 0 0 0 1.292.657l2.956-3.303zm6.037-.327l2.539-2.838A.75.75 0 0 0 17.25 8h-6.068zm-4.503-4.503l1.216-5.678a.75.75 0 0 0-1.292-.657L7.735 4.553z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoltSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.615 1.595a.75.75 0 0 1 .359.852L12.982 9.75h7.268a.75.75 0 0 1 .548 1.262l-10.5 11.25a.75.75 0 0 1-1.272-.71l1.992-7.302H3.75a.75.75 0 0 1-.548-1.262l10.5-11.25a.75.75 0 0 1 .913-.143" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoltTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.983 1.907a.75.75 0 0 0-1.292-.657l-8.5 9.5A.75.75 0 0 0 2.75 12h6.572l-1.305 6.093a.75.75 0 0 0 1.292.657l8.5-9.5A.75.75 0 0 0 17.25 8h-6.572z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookOpen(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 6.042A8.967 8.967 0 0 0 6 3.75c-1.052 0-2.062.18-3 .512v14.25A8.987 8.987 0 0 1 6 18c2.305 0 4.408.867 6 2.292m0-14.25a8.966 8.966 0 0 1 6-2.292c1.052 0 2.062.18 3 .512v14.25A8.987 8.987 0 0 0 18 18a8.967 8.967 0 0 0-6 2.292m0-14.25v14.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookOpenSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.25 3.688a8.035 8.035 0 0 0-4.872-.523A.48.48 0 0 0 2 3.64v7.994c0 .345.342.588.679.512a6.02 6.02 0 0 1 4.571.81zm1.5 9.268a6.02 6.02 0 0 1 4.571-.81c.337.075.679-.167.679-.512V3.64a.48.48 0 0 0-.378-.475a8.034 8.034 0 0 0-4.872.523z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookOpenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.25 4.533A9.707 9.707 0 0 0 6 3a9.735 9.735 0 0 0-3.25.555a.75.75 0 0 0-.5.707v14.25a.75.75 0 0 0 1 .707A8.237 8.237 0 0 1 6 18.75c1.995 0 3.823.707 5.25 1.886zm1.5 16.103A8.214 8.214 0 0 1 18 18.75c.966 0 1.89.166 2.75.47a.75.75 0 0 0 1-.708V4.262a.75.75 0 0 0-.5-.707A9.735 9.735 0 0 0 18 3a9.707 9.707 0 0 0-5.25 1.533z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookOpenTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.75 16.82A7.462 7.462 0 0 1 15 15.5a7.5 7.5 0 0 1 2.046.282a.75.75 0 0 0 .954-.722v-11a.75.75 0 0 0-.546-.721A9.006 9.006 0 0 0 15 3a8.963 8.963 0 0 0-4.25 1.065zM9.25 4.065A8.963 8.963 0 0 0 5 3a9 9 0 0 0-2.454.339A.75.75 0 0 0 2 4.06v11a.75.75 0 0 0 .954.721A7.506 7.506 0 0 1 5 15.5c1.579 0 3.042.487 4.25 1.32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.593 3.322c1.1.128 1.907 1.077 1.907 2.185V21L12 17.25L4.5 21V5.507c0-1.108.806-2.057 1.907-2.185a48.507 48.507 0 0 1 11.186 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.75 2a.75.75 0 0 0-.75.75v10.5a.75.75 0 0 0 1.28.53L8 10.06l3.72 3.72a.75.75 0 0 0 1.28-.53V2.75a.75.75 0 0 0-.75-.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkSlash(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3 3l1.664 1.664M21 21l-1.5-1.5m-5.485-1.242L12 17.25L4.5 21V8.742m.164-4.078a2.15 2.15 0 0 1 1.743-1.342a48.507 48.507 0 0 1 11.186 0c1.1.128 1.907 1.077 1.907 2.185V19.5M4.664 4.664L19.5 19.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkSlashSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 2.75v7.775L4.475 2h7.775a.75.75 0 0 1 .75.75M3 13.25V5.475l4.793 4.793L4.28 13.78A.75.75 0 0 1 3 13.25M2.22 2.22a.75.75 0 0 1 1.06 0l10.5 10.5a.75.75 0 1 1-1.06 1.06L2.22 3.28a.75.75 0 0 1 0-1.06"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkSlashSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.53 2.47a.75.75 0 0 0-1.06 1.06l18 18a.75.75 0 1 0 1.06-1.06zm16.72 3.037v11.561L5.853 2.671c.15-.043.306-.075.467-.094a49.255 49.255 0 0 1 11.36 0c1.497.174 2.57 1.46 2.57 2.93M3.75 21V6.932l14.063 14.063L12 18.088l-7.165 3.583A.75.75 0 0 1 3.75 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkSlashTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 4.517v9.301L5.433 2.252a41.44 41.44 0 0 1 9.637.058A2.213 2.213 0 0 1 17 4.517M3 17.25V6.182l10.654 10.654L10 15.082l-5.925 2.844A.75.75 0 0 1 3 17.25m.28-15.03a.75.75 0 0 0-1.06 1.06l14.5 14.5a.75.75 0 1 0 1.06-1.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.32 2.577a49.255 49.255 0 0 1 11.36 0c1.497.174 2.57 1.46 2.57 2.93V21a.75.75 0 0 1-1.085.67L12 18.089l-7.165 3.583A.75.75 0 0 1 3.75 21V5.507c0-1.47 1.073-2.756 2.57-2.93" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkSquare(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.5 3.75V16.5L12 14.25L7.5 16.5V3.75m9 0H18A2.25 2.25 0 0 1 20.25 6v12A2.25 2.25 0 0 1 18 20.25H6A2.25 2.25 0 0 1 3.75 18V6A2.25 2.25 0 0 1 6 3.75h1.5m9 0h-9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkSquareSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 2a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm1 2.25a.75.75 0 0 1 .75-.75h4.5a.75.75 0 0 1 .75.75v6a.75.75 0 0 1-1.28.53L8 9.06l-1.72 1.72A.75.75 0 0 1 5 10.25z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkSquareSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3zm1.5 1.5a.75.75 0 0 0-.75.75V16.5a.75.75 0 0 0 1.085.67L12 15.089l4.165 2.083a.75.75 0 0 0 1.085-.671V5.25a.75.75 0 0 0-.75-.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkSquareTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.25 2A2.25 2.25 0 0 0 2 4.25v11.5A2.25 2.25 0 0 0 4.25 18h11.5A2.25 2.25 0 0 0 18 15.75V4.25A2.25 2.25 0 0 0 15.75 2zM6 13.25V3.5h8v9.75a.75.75 0 0 1-1.064.681L10 12.576l-2.936 1.355A.75.75 0 0 1 6 13.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 2c-1.716 0-3.408.106-5.07.31A2.213 2.213 0 0 0 3 4.517V17.25a.75.75 0 0 0 1.075.676L10 15.082l5.926 2.844A.75.75 0 0 0 17 17.25V4.517c0-1.103-.806-2.068-1.93-2.207A41.403 41.403 0 0 0 10 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Briefcase(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.25 14.15v4.25c0 1.094-.787 2.036-1.872 2.18c-2.087.277-4.216.42-6.378.42s-4.291-.143-6.378-.42c-1.085-.144-1.872-1.086-1.872-2.18v-4.25m16.5 0a2.18 2.18 0 0 0 .75-1.661V8.706c0-1.081-.768-2.015-1.837-2.175a48.114 48.114 0 0 0-3.413-.387m4.5 8.006c-.194.165-.42.295-.673.38A23.978 23.978 0 0 1 12 15.75a24 24 0 0 1-7.577-1.22a2.016 2.016 0 0 1-.673-.38m0 0A2.18 2.18 0 0 1 3 12.489V8.706c0-1.081.768-2.015 1.837-2.175a48.111 48.111 0 0 1 3.413-.387m7.5 0V5.25A2.25 2.25 0 0 0 13.5 3h-3a2.25 2.25 0 0 0-2.25 2.25v.894m7.5 0a48.667 48.667 0 0 0-7.5 0M12 12.75h.008v.008H12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BriefcaseSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 4V3a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v1H4a2 2 0 0 0-2 2v3a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2zM9 2.5H7a.5.5 0 0 0-.5.5v1h3V3a.5.5 0 0 0-.5-.5M9 9a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/><path fill="currentColor" d="M3 11.83V12a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2v-.17c-.313.11-.65.17-1 .17H4c-.35 0-.687-.06-1-.17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BriefcaseSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M7.5 5.25a3 3 0 0 1 3-3h3a3 3 0 0 1 3 3v.205c.933.085 1.857.197 2.774.334c1.454.218 2.476 1.483 2.476 2.917v3.033c0 1.211-.734 2.352-1.936 2.752A24.726 24.726 0 0 1 12 15.75a24.72 24.72 0 0 1-7.814-1.259c-1.202-.4-1.936-1.541-1.936-2.752V8.706c0-1.434 1.022-2.7 2.476-2.917A48.814 48.814 0 0 1 7.5 5.455zm7.5 0v.09a49.488 49.488 0 0 0-6 0v-.09a1.5 1.5 0 0 1 1.5-1.5h3a1.5 1.5 0 0 1 1.5 1.5m-3 8.25a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5" clip-rule="evenodd"/><path d="M3 18.4v-2.796a4.3 4.3 0 0 0 .713.31A26.226 26.226 0 0 0 12 17.25c2.892 0 5.68-.468 8.287-1.335c.252-.084.49-.189.713-.311V18.4c0 1.452-1.047 2.728-2.523 2.923c-2.12.282-4.282.427-6.477.427a49.19 49.19 0 0 1-6.477-.427C4.047 21.128 3 19.852 3 18.4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BriefcaseTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M6 3.75A2.75 2.75 0 0 1 8.75 1h2.5A2.75 2.75 0 0 1 14 3.75v.443c.572.055 1.14.122 1.706.2C17.053 4.582 18 5.75 18 7.07v3.469c0 1.126-.694 2.191-1.83 2.54c-1.952.599-4.024.921-6.17.921s-4.219-.322-6.17-.921C2.694 12.73 2 11.665 2 10.539V7.07c0-1.321.947-2.489 2.294-2.676A41.047 41.047 0 0 1 6 4.193zm6.5 0v.325a41.622 41.622 0 0 0-5 0V3.75c0-.69.56-1.25 1.25-1.25h2.5c.69 0 1.25.56 1.25 1.25M10 10a1 1 0 0 0-1 1v.01a1 1 0 0 0 1 1h.01a1 1 0 0 0 1-1V11a1 1 0 0 0-1-1z" clip-rule="evenodd"/><path d="M3 15.055v-.684c.126.053.255.1.39.142c2.092.642 4.313.987 6.61.987c2.297 0 4.518-.345 6.61-.987c.135-.041.264-.089.39-.142v.684c0 1.347-.985 2.53-2.363 2.686a41.454 41.454 0 0 1-9.274 0C3.985 17.585 3 16.402 3 15.055"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BugAnt(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 12.75c1.148 0 2.278.08 3.383.237c1.037.146 1.866.966 1.866 2.013c0 3.728-2.35 6.75-5.25 6.75S6.75 18.728 6.75 15c0-1.046.83-1.867 1.866-2.013A24.204 24.204 0 0 1 12 12.75m0 0c2.883 0 5.647.508 8.208 1.44a23.91 23.91 0 0 1-1.153 6.06M12 12.75c-2.883 0-5.647.508-8.208 1.44c.125 2.105.52 4.136 1.153 6.06M12 12.75a2.25 2.25 0 0 0 2.248-2.354M12 12.75a2.25 2.25 0 0 1-2.248-2.354M12 8.25c.995 0 1.971-.08 2.922-.236c.403-.066.74-.358.795-.762a3.778 3.778 0 0 0-.399-2.25M12 8.25c-.995 0-1.97-.08-2.922-.236c-.402-.066-.74-.358-.795-.762a3.734 3.734 0 0 1 .4-2.253M12 8.25a2.25 2.25 0 0 0-2.248 2.146M12 8.25a2.25 2.25 0 0 1 2.248 2.146M8.683 5a6.032 6.032 0 0 1-1.155-1.002c.07-.63.27-1.222.574-1.747M8.683 5a3.75 3.75 0 0 1 6.635 0m0 0c.427-.283.815-.62 1.155-.999a4.471 4.471 0 0 0-.575-1.752M4.921 6a24.048 24.048 0 0 0-.392 3.314a23.88 23.88 0 0 0 5.223 1.082M19.08 6c.205 1.08.337 2.187.392 3.314a23.882 23.882 0 0 1-5.223 1.082"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BugAntSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.983 1.364a.75.75 0 0 0-1.281.78c.096.158.184.321.264.489a5.48 5.48 0 0 1-.713.386A2.993 2.993 0 0 0 8 2a2.99 2.99 0 0 0-2.253 1.02a5.485 5.485 0 0 1-.713-.387c.08-.168.168-.33.264-.489a.75.75 0 1 0-1.28-.78c-.245.401-.45.83-.61 1.278a.75.75 0 0 0 .239.84a7 7 0 0 0 1.422.876A3.01 3.01 0 0 0 5 5c0 .126.072.24.183.3c.386.205.797.37 1.227.487c-.126.165-.227.35-.297.549A10.418 10.418 0 0 1 3.51 5.5a10.686 10.686 0 0 1-.008-.733a.75.75 0 0 0-1.5-.033a12.222 12.222 0 0 0 .041 1.31a.75.75 0 0 0 .4.6A11.922 11.922 0 0 0 6.199 7.87c.04.084.088.166.14.243l-.214.031l-.027.005A11.913 11.913 0 0 0 2.444 9.36a.75.75 0 0 0-.4.6a12.148 12.148 0 0 0 .197 3.443a.75.75 0 0 0 1.47-.299a10.551 10.551 0 0 1-.2-2.6a10.35 10.35 0 0 1 1.085-.441c-.063.3-.096.614-.096.936c0 2.21 1.567 4 3.5 4s3.5-1.79 3.5-4c0-.322-.034-.636-.097-.937c.372.128.734.275 1.085.442a10.703 10.703 0 0 1-.199 2.6a.75.75 0 1 0 1.47.3a12.049 12.049 0 0 0 .197-3.443a.75.75 0 0 0-.4-.6a11.921 11.921 0 0 0-3.671-1.215l-.011-.002a11.95 11.95 0 0 0-.213-.03c.052-.078.1-.16.14-.244c1.336-.202 2.6-.623 3.755-1.227a.75.75 0 0 0 .4-.6a12.178 12.178 0 0 0 .041-1.31a.75.75 0 0 0-1.5.033a11.061 11.061 0 0 1-.008.733a10.43 10.43 0 0 1-2.602.836c-.07-.2-.17-.384-.297-.55a5.9 5.9 0 0 0 1.228-.488A.34.34 0 0 0 11 5c0-.22-.024-.435-.069-.642a7 7 0 0 0 1.422-.876a.75.75 0 0 0 .24-.84a6.97 6.97 0 0 0-.61-1.278"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BugAntSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.478 1.6a.75.75 0 0 1 .273 1.026a3.72 3.72 0 0 0-.425 1.121c.058.058.118.114.18.168A4.491 4.491 0 0 1 12 2.25a4.49 4.49 0 0 1 3.497 1.668c.06-.054.12-.11.178-.167a3.717 3.717 0 0 0-.426-1.125a.75.75 0 1 1 1.298-.752a5.22 5.22 0 0 1 .671 2.046a.75.75 0 0 1-.187.582c-.241.27-.505.52-.787.749a4.494 4.494 0 0 1 .216 2.1c-.106.792-.753 1.295-1.417 1.403c-.182.03-.364.057-.547.081c.152.227.273.476.359.742a23.122 23.122 0 0 0 3.832-.803a23.241 23.241 0 0 0-.345-2.634a.75.75 0 0 1 1.474-.28c.21 1.115.348 2.256.404 3.418a.75.75 0 0 1-.516.75a24.698 24.698 0 0 1-4.76 1.049c-.074.38-.22.735-.423 1.05c2.066.209 4.058.672 5.943 1.358a.75.75 0 0 1 .492.75a24.665 24.665 0 0 1-1.189 6.25a.75.75 0 0 1-1.425-.47a23.14 23.14 0 0 0 1.077-5.306c-.5-.169-1.009-.32-1.524-.455c.068.234.104.484.104.746c0 3.956-2.521 7.5-6 7.5c-3.478 0-6-3.544-6-7.5c0-.262.037-.511.104-.746c-.514.135-1.022.286-1.522.455c.154 1.838.52 3.616 1.077 5.307a.75.75 0 1 1-1.425.468a24.662 24.662 0 0 1-1.19-6.25a.75.75 0 0 1 .493-.749a24.586 24.586 0 0 1 4.964-1.24h.01c.321-.046.644-.085.969-.118a2.983 2.983 0 0 1-.424-1.05a24.614 24.614 0 0 1-4.76-1.05a.75.75 0 0 1-.516-.75c.057-1.16.194-2.302.405-3.417a.75.75 0 0 1 1.474.28c-.164.862-.28 1.74-.345 2.634c1.237.371 2.517.642 3.832.803c.085-.266.207-.515.359-.742a18.698 18.698 0 0 1-.547-.08c-.664-.11-1.311-.612-1.417-1.404a4.535 4.535 0 0 1 .217-2.103a6.788 6.788 0 0 1-.788-.751a.75.75 0 0 1-.187-.583a5.22 5.22 0 0 1 .67-2.04a.75.75 0 0 1 1.026-.273" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BugAntTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.56 1.14a.75.75 0 0 1 .177 1.045a3.989 3.989 0 0 0-.464.86c.185.17.382.329.59.473A3.993 3.993 0 0 1 10 2c1.272 0 2.405.594 3.137 1.518c.208-.144.405-.302.59-.473a3.989 3.989 0 0 0-.464-.86a.75.75 0 0 1 1.222-.869c.369.519.65 1.105.822 1.736a.75.75 0 0 1-.174.707a7.03 7.03 0 0 1-1.299 1.098A4 4 0 0 1 14 6c0 .52-.301.963-.723 1.187a6.961 6.961 0 0 1-1.158.486a2.5 2.5 0 0 1 .296.679a19.34 19.34 0 0 0 4.081-.96a19.655 19.655 0 0 0-.09-2.319a.75.75 0 1 1 1.493-.146a21.239 21.239 0 0 1 .08 3.028a.75.75 0 0 1-.482.667a20.873 20.873 0 0 1-5.153 1.249a2.521 2.521 0 0 1-.107.247a20.945 20.945 0 0 1 5.252 1.257a.75.75 0 0 1 .482.74a20.945 20.945 0 0 1-.908 5.107a.75.75 0 0 1-1.433-.444c.415-1.34.69-2.743.806-4.191c-.495-.173-1-.327-1.512-.46c.05.284.076.575.076.873c0 1.814-.517 3.312-1.426 4.37A4.639 4.639 0 0 1 10 19a4.639 4.639 0 0 1-3.574-1.63C5.516 16.311 5 14.813 5 13c0-.298.026-.59.076-.873c-.513.133-1.017.287-1.512.46c.116 1.448.39 2.85.806 4.191a.75.75 0 1 1-1.433.444a20.94 20.94 0 0 1-.908-5.107a.75.75 0 0 1 .482-.74a20.838 20.838 0 0 1 5.252-1.257a2.493 2.493 0 0 1-.107-.247a20.874 20.874 0 0 1-5.153-1.249a.75.75 0 0 1-.482-.667a21.342 21.342 0 0 1 .08-3.028a.75.75 0 1 1 1.493.146a19.745 19.745 0 0 0-.09 2.319c1.302.46 2.668.786 4.08.96c.066-.243.166-.471.297-.679a6.962 6.962 0 0 1-1.158-.486A1.348 1.348 0 0 1 6 6a4 4 0 0 1 .166-1.143a7.032 7.032 0 0 1-1.3-1.098a.75.75 0 0 1-.173-.707a5.48 5.48 0 0 1 .822-1.736a.75.75 0 0 1 1.046-.177" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingLibrary(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 21v-8.25M15.75 21v-8.25M8.25 21v-8.25M3 9l9-6l9 6m-1.5 12V10.333A48.36 48.36 0 0 0 12 9.75c-2.551 0-5.056.2-7.5.582V21M3 21h18M12 6.75h.008v.008H12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingLibrarySixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.605 2.112a.75.75 0 0 1 .79 0l5.25 3.25A.75.75 0 0 1 13 6.707V12.5h.25a.75.75 0 0 1 0 1.5H2.75a.75.75 0 0 1 0-1.5H3V6.707a.75.75 0 0 1-.645-1.345zM4.5 8.75a.75.75 0 0 1 1.5 0v3a.75.75 0 0 1-1.5 0zM8 8a.75.75 0 0 0-.75.75v3a.75.75 0 0 0 1.5 0v-3A.75.75 0 0 0 8 8m2 .75a.75.75 0 0 1 1.5 0v3a.75.75 0 0 1-1.5 0zM8 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingLibrarySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M11.584 2.376a.75.75 0 0 1 .832 0l9 6a.75.75 0 1 1-.832 1.248L12 3.901L3.416 9.624a.75.75 0 0 1-.832-1.248z"/><path fill-rule="evenodd" d="M20.25 10.333v9.917H21a.75.75 0 0 1 0 1.5H3a.75.75 0 0 1 0-1.5h.75v-9.917a.75.75 0 0 1 .634-.741A49.109 49.109 0 0 1 12 9c2.59 0 5.134.202 7.616.592a.75.75 0 0 1 .634.74m-7.5 2.417a.75.75 0 0 0-1.5 0v6.75a.75.75 0 0 0 1.5 0zm3-.75a.75.75 0 0 1 .75.75v6.75a.75.75 0 0 1-1.5 0v-6.75a.75.75 0 0 1 .75-.75M9 12.75a.75.75 0 0 0-1.5 0v6.75a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/><path d="M12 7.875a1.125 1.125 0 1 0 0-2.25a1.125 1.125 0 0 0 0 2.25"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingLibraryTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.674 2.075a.75.75 0 0 1 .652 0l7.25 3.5A.75.75 0 0 1 17 6.957V16.5h.25a.75.75 0 0 1 0 1.5H2.75a.75.75 0 0 1 0-1.5H3V6.957a.75.75 0 0 1-.576-1.382zM11 6a1 1 0 1 1-2 0a1 1 0 0 1 2 0M7.5 9.75a.75.75 0 0 0-1.5 0v5.5a.75.75 0 0 0 1.5 0zm3.25 0a.75.75 0 0 0-1.5 0v5.5a.75.75 0 0 0 1.5 0zm3.25 0a.75.75 0 0 0-1.5 0v5.5a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingOffice(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 21h16.5M4.5 3h15M5.25 3v18m13.5-18v18M9 6.75h1.5m-1.5 3h1.5m-1.5 3h1.5m3-6H15m-1.5 3H15m-1.5 3H15M9 21v-3.375c0-.621.504-1.125 1.125-1.125h3.75c.621 0 1.125.504 1.125 1.125V21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingOfficeSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.75 2a.75.75 0 0 0 0 1.5H4v9h-.25a.75.75 0 0 0 0 1.5H6a.5.5 0 0 0 .5-.5v-3A.5.5 0 0 1 7 10h2a.5.5 0 0 1 .5.5v3a.5.5 0 0 0 .5.5h2.25a.75.75 0 0 0 0-1.5H12v-9h.25a.75.75 0 0 0 0-1.5zM6.5 4a.5.5 0 0 0-.5.5V5a.5.5 0 0 0 .5.5H7a.5.5 0 0 0 .5-.5v-.5A.5.5 0 0 0 7 4zM6 7a.5.5 0 0 1 .5-.5H7a.5.5 0 0 1 .5.5v.5A.5.5 0 0 1 7 8h-.5a.5.5 0 0 1-.5-.5zm3-3a.5.5 0 0 0-.5.5V5a.5.5 0 0 0 .5.5h.5A.5.5 0 0 0 10 5v-.5a.5.5 0 0 0-.5-.5zm-.5 3a.5.5 0 0 1 .5-.5h.5a.5.5 0 0 1 .5.5v.5a.5.5 0 0 1-.5.5H9a.5.5 0 0 1-.5-.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingOfficeSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 2.25a.75.75 0 0 0 0 1.5v16.5h-.75a.75.75 0 0 0 0 1.5h16.5a.75.75 0 0 0 0-1.5h-.75V3.75a.75.75 0 0 0 0-1.5zM9 6a.75.75 0 0 0 0 1.5h1.5a.75.75 0 0 0 0-1.5zm-.75 3.75A.75.75 0 0 1 9 9h1.5a.75.75 0 0 1 0 1.5H9a.75.75 0 0 1-.75-.75M9 12a.75.75 0 0 0 0 1.5h1.5a.75.75 0 0 0 0-1.5zm3.75-5.25A.75.75 0 0 1 13.5 6H15a.75.75 0 0 1 0 1.5h-1.5a.75.75 0 0 1-.75-.75M13.5 9a.75.75 0 0 0 0 1.5H15A.75.75 0 0 0 15 9zm-.75 3.75a.75.75 0 0 1 .75-.75H15a.75.75 0 0 1 0 1.5h-1.5a.75.75 0 0 1-.75-.75M9 19.5v-2.25a.75.75 0 0 1 .75-.75h4.5a.75.75 0 0 1 .75.75v2.25a.75.75 0 0 1-.75.75h-4.5A.75.75 0 0 1 9 19.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingOfficeTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 16.5v-13h-.25a.75.75 0 0 1 0-1.5h12.5a.75.75 0 0 1 0 1.5H16v13h.25a.75.75 0 0 1 0 1.5h-3.5a.75.75 0 0 1-.75-.75v-2.5a.75.75 0 0 0-.75-.75h-2.5a.75.75 0 0 0-.75.75v2.5a.75.75 0 0 1-.75.75h-3.5a.75.75 0 0 1 0-1.5zm3-11a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1-.5-.5zM7.5 9a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5zM11 5.5a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1-.5-.5zm.5 3.5a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingOfficeTwo(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.25 21h19.5m-18-18v18m10.5-18v18m6-13.5V21M6.75 6.75h.75m-.75 3h.75m-.75 3h.75m3-6h.75m-.75 3h.75m-.75 3h.75M6.75 21v-3.375c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125V21M3 3h12m-.75 4.5H21m-3.75 3.75h.008v.008h-.008zm0 3h.008v.008h-.008zm0 3h.008v.008h-.008z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingOfficeTwoSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.75 2a.75.75 0 0 0 0 1.5H2v9h-.25a.75.75 0 0 0 0 1.5h1.5a.75.75 0 0 0 .75-.75v-1.5a.75.75 0 0 1 .75-.75h1.5a.75.75 0 0 1 .75.75v1.5c0 .414.336.75.75.75h.5a.75.75 0 0 0 .75-.75V3.5h.25a.75.75 0 0 0 0-1.5zM3.5 5.5A.5.5 0 0 1 4 5h.5a.5.5 0 0 1 .5.5V6a.5.5 0 0 1-.5.5H4a.5.5 0 0 1-.5-.5zm.5 2a.5.5 0 0 0-.5.5v.5A.5.5 0 0 0 4 9h.5a.5.5 0 0 0 .5-.5V8a.5.5 0 0 0-.5-.5zm2-2a.5.5 0 0 1 .5-.5H7a.5.5 0 0 1 .5.5V6a.5.5 0 0 1-.5.5h-.5A.5.5 0 0 1 6 6zm.5 2A.5.5 0 0 0 6 8v.5a.5.5 0 0 0 .5.5H7a.5.5 0 0 0 .5-.5V8a.5.5 0 0 0-.5-.5zm5-1.5a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h2.75a.75.75 0 0 0 0-1.5H14v-5h.25a.75.75 0 0 0 0-1.5zm.5 1.5h.5a.5.5 0 0 1 .5.5v.5a.5.5 0 0 1-.5.5H12a.5.5 0 0 1-.5-.5V8a.5.5 0 0 1 .5-.5m0 2.5a.5.5 0 0 0-.5.5v.5a.5.5 0 0 0 .5.5h.5a.5.5 0 0 0 .5-.5v-.5a.5.5 0 0 0-.5-.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingOfficeTwoSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 2.25a.75.75 0 0 0 0 1.5v16.5h-.75a.75.75 0 0 0 0 1.5H15v-18a.75.75 0 0 0 0-1.5zM6.75 19.5v-2.25a.75.75 0 0 1 .75-.75h3a.75.75 0 0 1 .75.75v2.25a.75.75 0 0 1-.75.75h-3a.75.75 0 0 1-.75-.75M6 6.75A.75.75 0 0 1 6.75 6h.75a.75.75 0 0 1 0 1.5h-.75A.75.75 0 0 1 6 6.75M6.75 9a.75.75 0 0 0 0 1.5h.75a.75.75 0 0 0 0-1.5zM6 12.75a.75.75 0 0 1 .75-.75h.75a.75.75 0 0 1 0 1.5h-.75a.75.75 0 0 1-.75-.75M10.5 6a.75.75 0 0 0 0 1.5h.75a.75.75 0 0 0 0-1.5zm-.75 3.75A.75.75 0 0 1 10.5 9h.75a.75.75 0 0 1 0 1.5h-.75a.75.75 0 0 1-.75-.75M10.5 12a.75.75 0 0 0 0 1.5h.75a.75.75 0 0 0 0-1.5zm6-5.25v15h5.25a.75.75 0 0 0 0-1.5H21v-12a.75.75 0 0 0 0-1.5zm1.5 4.5a.75.75 0 0 1 .75-.75h.008a.75.75 0 0 1 .75.75v.008a.75.75 0 0 1-.75.75h-.008a.75.75 0 0 1-.75-.75zm.75 2.25a.75.75 0 0 0-.75.75v.008c0 .414.336.75.75.75h.008a.75.75 0 0 0 .75-.75v-.008a.75.75 0 0 0-.75-.75zM18 17.25a.75.75 0 0 1 .75-.75h.008a.75.75 0 0 1 .75.75v.008a.75.75 0 0 1-.75.75h-.008a.75.75 0 0 1-.75-.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingOfficeTwoTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 2.75A.75.75 0 0 1 1.75 2h10.5a.75.75 0 0 1 0 1.5H12v13.75a.75.75 0 0 1-.75.75h-1.5a.75.75 0 0 1-.75-.75v-2.5a.75.75 0 0 0-.75-.75h-2.5a.75.75 0 0 0-.75.75v2.5a.75.75 0 0 1-.75.75h-2.5a.75.75 0 0 1 0-1.5H2v-13h-.25A.75.75 0 0 1 1 2.75M4 5.5a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1-.5-.5zM4.5 9a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5zM8 5.5a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1-.5-.5zM8.5 9a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5zm5.75-3a.75.75 0 0 0-.75.75V17a1 1 0 0 0 1 1h3.75a.75.75 0 0 0 0-1.5H18v-9h.25a.75.75 0 0 0 0-1.5zm.5 3.5a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1-.5-.5zm.5 3.5a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingStorefront(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.5 21v-7.5a.75.75 0 0 1 .75-.75h3a.75.75 0 0 1 .75.75V21m-4.5 0H2.36m11.14 0H18m0 0h3.64m-1.39 0V9.349M3.75 21V9.349m0 0a3.001 3.001 0 0 0 3.75-.615A2.993 2.993 0 0 0 9.75 9.75c.896 0 1.7-.393 2.25-1.016a2.993 2.993 0 0 0 2.25 1.016c.896 0 1.7-.393 2.25-1.015a3.001 3.001 0 0 0 3.75.614m-16.5 0a3.004 3.004 0 0 1-.621-4.72l1.189-1.19A1.5 1.5 0 0 1 5.378 3h13.243a1.5 1.5 0 0 1 1.06.44l1.19 1.189a3 3 0 0 1-.621 4.72M6.75 18h3.75a.75.75 0 0 0 .75-.75V13.5a.75.75 0 0 0-.75-.75H6.75a.75.75 0 0 0-.75.75v3.75c0 .414.336.75.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingStorefrontSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 7c.681 0 1.3-.273 1.75-.715C6.7 6.727 7.319 7 8 7s1.3-.273 1.75-.715A2.5 2.5 0 1 0 11.5 2h-7a2.5 2.5 0 0 0 0 5m1.75 1.097A3.986 3.986 0 0 1 4.5 8.5A4 4 0 0 1 3 8.21v4.29h-.25a.75.75 0 0 0 0 1.5h.5a.754.754 0 0 0 .138-.013A.5.5 0 0 0 3.5 14H6a.5.5 0 0 0 .5-.5v-3A.5.5 0 0 1 7 10h2a.5.5 0 0 1 .5.5v3a.5.5 0 0 0 .5.5h2.5a.5.5 0 0 0 .112-.013c.045.009.09.013.138.013h.5a.75.75 0 1 0 0-1.5H13V8.21a4 4 0 0 1-1.5.29a3.986 3.986 0 0 1-1.75-.403A3.986 3.986 0 0 1 8 8.5a3.986 3.986 0 0 1-1.75-.403"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingStorefrontSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M5.223 2.25c-.497 0-.974.198-1.325.55l-1.3 1.298A3.75 3.75 0 0 0 7.5 9.75a3.74 3.74 0 0 0 2.25.75c.844 0 1.624-.28 2.25-.75c.626.47 1.406.75 2.25.75a3.74 3.74 0 0 0 2.25-.75a3.75 3.75 0 0 0 4.902-5.652l-1.3-1.299a1.875 1.875 0 0 0-1.325-.549z"/><path fill-rule="evenodd" d="M3 20.25v-8.755c1.42.674 3.08.673 4.5 0A5.234 5.234 0 0 0 9.75 12c.804 0 1.568-.182 2.25-.506a5.234 5.234 0 0 0 2.25.506c.804 0 1.567-.182 2.25-.506a5.26 5.26 0 0 0 4.5.001v8.755h.75a.75.75 0 0 1 0 1.5H2.25a.75.75 0 0 1 0-1.5zm3-6a.75.75 0 0 1 .75-.75h3a.75.75 0 0 1 .75.75v3a.75.75 0 0 1-.75.75h-3a.75.75 0 0 1-.75-.75zm8.25-.75a.75.75 0 0 0-.75.75v5.25c0 .414.336.75.75.75h3a.75.75 0 0 0 .75-.75v-5.25a.75.75 0 0 0-.75-.75z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingStorefrontTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.879 7.121A3 3 0 0 0 7.5 6.66A2.997 2.997 0 0 0 10 8a2.997 2.997 0 0 0 2.5-1.34a3 3 0 1 0 4.622-3.78l-.293-.293A2 2 0 0 0 15.415 2H4.585a2 2 0 0 0-1.414.586l-.292.292a3 3 0 0 0 0 4.243M3 9.032a4.507 4.507 0 0 0 4.5-.29A4.48 4.48 0 0 0 10 9.5a4.48 4.48 0 0 0 2.5-.758a4.507 4.507 0 0 0 4.5.29V16.5h.25a.75.75 0 0 1 0 1.5h-4.5a.75.75 0 0 1-.75-.75v-3.5a.75.75 0 0 0-.75-.75h-2.5a.75.75 0 0 0-.75.75v3.5a.75.75 0 0 1-.75.75h-4.5a.75.75 0 0 1 0-1.5H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cake(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 8.25v-1.5m0 1.5c-1.355 0-2.697.056-4.024.166C6.845 8.51 6 9.473 6 10.608v2.513m6-4.871c1.355 0 2.697.056 4.024.166C17.155 8.51 18 9.473 18 10.608v2.513M15 8.25v-1.5m-6 1.5v-1.5m12 9.75l-1.5.75a3.354 3.354 0 0 1-3 0a3.354 3.354 0 0 0-3 0a3.354 3.354 0 0 1-3 0a3.354 3.354 0 0 0-3 0a3.354 3.354 0 0 1-3 0L3 16.5m15-3.379a48.474 48.474 0 0 0-6-.371c-2.032 0-4.034.126-6 .371m12 0c.39.049.777.102 1.163.16c1.07.16 1.837 1.094 1.837 2.175v5.169c0 .621-.504 1.125-1.125 1.125H4.125A1.125 1.125 0 0 1 3 20.625v-5.17c0-1.08.768-2.014 1.837-2.174A47.78 47.78 0 0 1 6 13.12m6.265-10.01a.375.375 0 1 1-.53 0L12 2.845zm-3 0a.375.375 0 1 1-.53 0L9 2.845zm6 0a.375.375 0 1 1-.53 0L15 2.845z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CakeSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4.75 1l-.884.884a1.25 1.25 0 1 0 1.768 0zm6.5 0l-.884.884a1.25 1.25 0 1 0 1.768 0zm-2.366.884L8 1l-.884.884a1.25 1.25 0 1 0 1.768 0M4 7a2 2 0 0 0-2 2v1.034a3.26 3.26 0 0 0 1.028-.167l.47-.157a4.75 4.75 0 0 1 3.004 0l.47.157a3.25 3.25 0 0 0 2.056 0l.47-.157a4.75 4.75 0 0 1 3.004 0l.47.157a3.26 3.26 0 0 0 1.028.167V9a2 2 0 0 0-2-2V5.75a.75.75 0 0 0-1.5 0V7H8.75V5.75a.75.75 0 0 0-1.5 0V7H5.5V5.75a.75.75 0 0 0-1.5 0zm10 4.534a4.749 4.749 0 0 1-1.502-.244l-.47-.157a3.25 3.25 0 0 0-2.056 0l-.47.157a4.75 4.75 0 0 1-3.004 0l-.47-.157a3.25 3.25 0 0 0-2.056 0l-.47.157A4.748 4.748 0 0 1 2 11.534V13a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CakeSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15 1.784l-.796.795a1.125 1.125 0 1 0 1.592 0zm-3 0l-.796.795a1.125 1.125 0 1 0 1.592 0zm-3 0l-.796.795a1.125 1.125 0 1 0 1.592 0zm.75 5.763c.498-.021.998-.035 1.5-.042V6.75a.75.75 0 0 1 1.5 0v.755c.502.007 1.002.021 1.5.042V6.75a.75.75 0 0 1 1.5 0v.88l.307.022c1.55.117 2.693 1.427 2.693 2.946v1.018a62.182 62.182 0 0 0-13.5 0v-1.018c0-1.519 1.143-2.829 2.693-2.946l.307-.022v-.88a.75.75 0 0 1 1.5 0zM12 12.75c-2.472 0-4.9.184-7.274.54c-1.454.217-2.476 1.482-2.476 2.916v.384a4.104 4.104 0 0 1 2.585.364a2.605 2.605 0 0 0 2.33 0a4.104 4.104 0 0 1 3.67 0a2.605 2.605 0 0 0 2.33 0a4.104 4.104 0 0 1 3.67 0a2.605 2.605 0 0 0 2.33 0a4.104 4.104 0 0 1 2.585-.364v-.384c0-1.434-1.022-2.7-2.476-2.917A49.138 49.138 0 0 0 12 12.75m9.75 5.381a2.604 2.604 0 0 0-1.915.165a4.104 4.104 0 0 1-3.67 0a2.605 2.605 0 0 0-2.33 0a4.104 4.104 0 0 1-3.67 0a2.605 2.605 0 0 0-2.33 0a4.104 4.104 0 0 1-3.67 0a2.604 2.604 0 0 0-1.915-.165v2.494c0 1.035.84 1.875 1.875 1.875h15.75c1.035 0 1.875-.84 1.875-1.875z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CakeTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.75.98l-.884.883a1.25 1.25 0 1 0 1.768 0zm6.5 0l-.884.883a1.25 1.25 0 1 0 1.768 0zM10 .98l.884.883a1.25 1.25 0 1 1-1.768 0zM7.5 5.75a.75.75 0 0 0-1.5 0v.464c-1.179.305-2 1.39-2 2.622v.094c.1-.02.202-.038.306-.052A42.867 42.867 0 0 1 10 8.5c1.93 0 3.83.129 5.694.378c.104.014.206.032.306.052v-.094c0-1.232-.821-2.317-2-2.622V5.75a.75.75 0 0 0-1.5 0v.318a45.645 45.645 0 0 0-1.75-.062V5.75a.75.75 0 0 0-1.5 0v.256c-.586.01-1.17.03-1.75.062zm-2.995 4.615A41.36 41.36 0 0 1 10 10c1.863 0 3.697.124 5.495.365C16.967 10.562 18 11.838 18 13.28v.693a3.72 3.72 0 0 1-1.665-.393a5.222 5.222 0 0 0-4.67 0a3.722 3.722 0 0 1-3.33 0a5.222 5.222 0 0 0-4.67 0A3.72 3.72 0 0 1 2 13.972v-.693c0-1.441 1.033-2.717 2.505-2.914m11.16 4.555a5.22 5.22 0 0 0 2.335.552V16.5a1.5 1.5 0 0 1-1.5 1.5h-13A1.5 1.5 0 0 1 2 16.5v-1.028c.8 0 1.6-.184 2.335-.551a3.722 3.722 0 0 1 3.33 0c1.47.735 3.2.735 4.67 0a3.722 3.722 0 0 1 3.33 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calculator(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.75 15.75V18m-7.5-6.75h.008v.008H8.25zm0 2.25h.008v.008H8.25zm0 2.25h.008v.008H8.25zm0 2.25h.008v.008H8.25zm2.498-6.75h.007v.008h-.007zm0 2.25h.007v.008h-.007zm0 2.25h.007v.008h-.007zm0 2.25h.007v.008h-.007zm2.504-6.75h.008v.008h-.008zm0 2.25h.008v.008h-.008zm0 2.25h.008v.008h-.008zm0 2.25h.008v.008h-.008zm2.498-6.75h.008v.008h-.008zm0 2.25h.008v.008h-.008zM8.25 6h7.5v2.25h-7.5zM12 2.25c-1.892 0-3.758.11-5.593.322C5.307 2.7 4.5 3.65 4.5 4.757V19.5a2.25 2.25 0 0 0 2.25 2.25h10.5a2.25 2.25 0 0 0 2.25-2.25V4.757c0-1.108-.806-2.057-1.907-2.185A48.507 48.507 0 0 0 12 2.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalculatorSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 1a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2zm.75 6a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5M5 3.75A.75.75 0 0 1 5.75 3h4.5a.75.75 0 0 1 0 1.5h-4.5A.75.75 0 0 1 5 3.75m.75 7.75a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5M5 10a.75.75 0 1 1 1.5 0A.75.75 0 0 1 5 10m5.25-3a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5m-.75 3a.75.75 0 0 1 1.5 0v2.25a.75.75 0 0 1-1.5 0zM8 7a.75.75 0 1 0 0 1.5A.75.75 0 0 0 8 7m-.75 5.25a.75.75 0 1 1 1.5 0a.75.75 0 0 1-1.5 0m.75-3a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalculatorSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.32 1.827a49.255 49.255 0 0 1 11.36 0c1.497.174 2.57 1.46 2.57 2.93V19.5a3 3 0 0 1-3 3H6.75a3 3 0 0 1-3-3V4.757c0-1.47 1.073-2.756 2.57-2.93M7.5 11.25a.75.75 0 0 1 .75-.75h.008a.75.75 0 0 1 .75.75v.008a.75.75 0 0 1-.75.75H8.25a.75.75 0 0 1-.75-.75zm.75 1.5a.75.75 0 0 0-.75.75v.008c0 .414.336.75.75.75h.008a.75.75 0 0 0 .75-.75V13.5a.75.75 0 0 0-.75-.75zm-.75 3a.75.75 0 0 1 .75-.75h.008a.75.75 0 0 1 .75.75v.008a.75.75 0 0 1-.75.75H8.25a.75.75 0 0 1-.75-.75zm.75 1.5a.75.75 0 0 0-.75.75v.008c0 .414.336.75.75.75h.008a.75.75 0 0 0 .75-.75V18a.75.75 0 0 0-.75-.75zm1.748-6a.75.75 0 0 1 .75-.75h.007a.75.75 0 0 1 .75.75v.008a.75.75 0 0 1-.75.75h-.007a.75.75 0 0 1-.75-.75zm.75 1.5a.75.75 0 0 0-.75.75v.008c0 .414.335.75.75.75h.007a.75.75 0 0 0 .75-.75V13.5a.75.75 0 0 0-.75-.75zm-.75 3a.75.75 0 0 1 .75-.75h.007a.75.75 0 0 1 .75.75v.008a.75.75 0 0 1-.75.75h-.007a.75.75 0 0 1-.75-.75zm.75 1.5a.75.75 0 0 0-.75.75v.008c0 .414.335.75.75.75h.007a.75.75 0 0 0 .75-.75V18a.75.75 0 0 0-.75-.75zm1.754-6a.75.75 0 0 1 .75-.75h.008a.75.75 0 0 1 .75.75v.008a.75.75 0 0 1-.75.75h-.008a.75.75 0 0 1-.75-.75zm.75 1.5a.75.75 0 0 0-.75.75v.008c0 .414.336.75.75.75h.008a.75.75 0 0 0 .75-.75V13.5a.75.75 0 0 0-.75-.75zm-.75 3a.75.75 0 0 1 .75-.75h.008a.75.75 0 0 1 .75.75v.008a.75.75 0 0 1-.75.75h-.008a.75.75 0 0 1-.75-.75zm.75 1.5a.75.75 0 0 0-.75.75v.008c0 .414.336.75.75.75h.008a.75.75 0 0 0 .75-.75V18a.75.75 0 0 0-.75-.75zm1.748-6a.75.75 0 0 1 .75-.75h.008a.75.75 0 0 1 .75.75v.008a.75.75 0 0 1-.75.75h-.008a.75.75 0 0 1-.75-.75zm.75 1.5a.75.75 0 0 0-.75.75v.008c0 .414.336.75.75.75h.008a.75.75 0 0 0 .75-.75V13.5a.75.75 0 0 0-.75-.75zm-8.25-6A.75.75 0 0 1 8.25 6h7.5a.75.75 0 0 1 .75.75v.75a.75.75 0 0 1-.75.75h-7.5a.75.75 0 0 1-.75-.75zm9 9a.75.75 0 0 0-1.5 0V18a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalculatorTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 1c-1.716 0-3.408.106-5.07.31A2.213 2.213 0 0 0 3 3.517V16.75A2.25 2.25 0 0 0 5.25 19h9.5A2.25 2.25 0 0 0 17 16.75V3.517c0-1.103-.806-2.068-1.93-2.207A41.403 41.403 0 0 0 10 1M5.99 8.75A.75.75 0 0 1 6.74 8h.01a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75h-.01a.75.75 0 0 1-.75-.75zm.75 1.417a.75.75 0 0 0-.75.75v.01c0 .414.336.75.75.75h.01a.75.75 0 0 0 .75-.75v-.01a.75.75 0 0 0-.75-.75zm-.75 2.916a.75.75 0 0 1 .75-.75h.01a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75h-.01a.75.75 0 0 1-.75-.75zm.75 1.417a.75.75 0 0 0-.75.75v.01c0 .414.336.75.75.75h.01a.75.75 0 0 0 .75-.75v-.01a.75.75 0 0 0-.75-.75zm1.417-5.75a.75.75 0 0 1 .75-.75h.01a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75h-.01a.75.75 0 0 1-.75-.75zm.75 1.417a.75.75 0 0 0-.75.75v.01c0 .414.336.75.75.75h.01a.75.75 0 0 0 .75-.75v-.01a.75.75 0 0 0-.75-.75zm-.75 2.916a.75.75 0 0 1 .75-.75h.01a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75h-.01a.75.75 0 0 1-.75-.75zm.75 1.417a.75.75 0 0 0-.75.75v.01c0 .414.336.75.75.75h.01a.75.75 0 0 0 .75-.75v-.01a.75.75 0 0 0-.75-.75zm1.42-5.75a.75.75 0 0 1 .75-.75h.01a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75h-.01a.75.75 0 0 1-.75-.75zm.75 1.417a.75.75 0 0 0-.75.75v.01c0 .414.336.75.75.75h.01a.75.75 0 0 0 .75-.75v-.01a.75.75 0 0 0-.75-.75zm-.75 2.916a.75.75 0 0 1 .75-.75h.01a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75h-.01a.75.75 0 0 1-.75-.75zm.75 1.417a.75.75 0 0 0-.75.75v.01c0 .414.336.75.75.75h.01a.75.75 0 0 0 .75-.75v-.01a.75.75 0 0 0-.75-.75zM12.5 8.75a.75.75 0 0 1 .75-.75h.01a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75h-.01a.75.75 0 0 1-.75-.75zm.75 1.417a.75.75 0 0 0-.75.75v.01c0 .414.336.75.75.75h.01a.75.75 0 0 0 .75-.75v-.01a.75.75 0 0 0-.75-.75zm0 2.166a.75.75 0 0 1 .75.75v2.167a.75.75 0 1 1-1.5 0v-2.167a.75.75 0 0 1 .75-.75M6.75 4a.75.75 0 0 0-.75.75v.5c0 .414.336.75.75.75h6.5a.75.75 0 0 0 .75-.75v-.5a.75.75 0 0 0-.75-.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.75 3v2.25M17.25 3v2.25M3 18.75V7.5a2.25 2.25 0 0 1 2.25-2.25h13.5A2.25 2.25 0 0 1 21 7.5v11.25m-18 0A2.25 2.25 0 0 0 5.25 21h13.5A2.25 2.25 0 0 0 21 18.75m-18 0v-7.5A2.25 2.25 0 0 1 5.25 9h13.5A2.25 2.25 0 0 1 21 11.25v7.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarDays(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.75 3v2.25M17.25 3v2.25M3 18.75V7.5a2.25 2.25 0 0 1 2.25-2.25h13.5A2.25 2.25 0 0 1 21 7.5v11.25m-18 0A2.25 2.25 0 0 0 5.25 21h13.5A2.25 2.25 0 0 0 21 18.75m-18 0v-7.5A2.25 2.25 0 0 1 5.25 9h13.5A2.25 2.25 0 0 1 21 11.25v7.5m-9-6h.008v.008H12zM12 15h.008v.008H12zm0 2.25h.008v.008H12zM9.75 15h.008v.008H9.75zm0 2.25h.008v.008H9.75zM7.5 15h.008v.008H7.5zm0 2.25h.008v.008H7.5zm6.75-4.5h.008v.008h-.008zm0 2.25h.008v.008h-.008zm0 2.25h.008v.008h-.008zm2.25-4.5h.008v.008H16.5zm0 2.25h.008v.008H16.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarDaysSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M5.75 7.5a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5M5 10.25a.75.75 0 1 1 1.5 0a.75.75 0 0 1-1.5 0m5.25-2.75a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5m-3 .75a.75.75 0 1 1 1.5 0a.75.75 0 0 1-1.5 0M8 9.5A.75.75 0 1 0 8 11a.75.75 0 0 0 0-1.5"/><path fill-rule="evenodd" d="M4.75 1a.75.75 0 0 0-.75.75V3a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2V1.75a.75.75 0 0 0-1.5 0V3h-5V1.75A.75.75 0 0 0 4.75 1M3.5 7a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1v4.5a1 1 0 0 1-1 1h-7a1 1 0 0 1-1-1z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarDaysSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12.75 12.75a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0m-5.25 3a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5m.75 1.5a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0m1.5-1.5a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5m.75 1.5a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0m1.5-1.5a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5m.75 1.5a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0m1.5-1.5a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5m.75 1.5a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0m1.5-1.5a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5m-1.5-3a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0m1.5.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5"/><path fill-rule="evenodd" d="M6.75 2.25A.75.75 0 0 1 7.5 3v1.5h9V3A.75.75 0 0 1 18 3v1.5h.75a3 3 0 0 1 3 3v11.25a3 3 0 0 1-3 3H5.25a3 3 0 0 1-3-3V7.5a3 3 0 0 1 3-3H6V3a.75.75 0 0 1 .75-.75m13.5 9a1.5 1.5 0 0 0-1.5-1.5H5.25a1.5 1.5 0 0 0-1.5 1.5v7.5a1.5 1.5 0 0 0 1.5 1.5h13.5a1.5 1.5 0 0 0 1.5-1.5z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarDaysTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M5.25 12a.75.75 0 0 1 .75-.75h.01a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75H6a.75.75 0 0 1-.75-.75zM6 13.25a.75.75 0 0 0-.75.75v.01c0 .414.336.75.75.75h.01a.75.75 0 0 0 .75-.75V14a.75.75 0 0 0-.75-.75zM7.25 12a.75.75 0 0 1 .75-.75h.01a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75H8a.75.75 0 0 1-.75-.75zM8 13.25a.75.75 0 0 0-.75.75v.01c0 .414.336.75.75.75h.01a.75.75 0 0 0 .75-.75V14a.75.75 0 0 0-.75-.75zM9.25 10a.75.75 0 0 1 .75-.75h.01a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75H10a.75.75 0 0 1-.75-.75zm.75 1.25a.75.75 0 0 0-.75.75v.01c0 .414.336.75.75.75h.01a.75.75 0 0 0 .75-.75V12a.75.75 0 0 0-.75-.75zM9.25 14a.75.75 0 0 1 .75-.75h.01a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75H10a.75.75 0 0 1-.75-.75zM12 9.25a.75.75 0 0 0-.75.75v.01c0 .414.336.75.75.75h.01a.75.75 0 0 0 .75-.75V10a.75.75 0 0 0-.75-.75zM11.25 12a.75.75 0 0 1 .75-.75h.01a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75H12a.75.75 0 0 1-.75-.75zm.75 1.25a.75.75 0 0 0-.75.75v.01c0 .414.336.75.75.75h.01a.75.75 0 0 0 .75-.75V14a.75.75 0 0 0-.75-.75zM13.25 10a.75.75 0 0 1 .75-.75h.01a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75H14a.75.75 0 0 1-.75-.75zm.75 1.25a.75.75 0 0 0-.75.75v.01c0 .414.336.75.75.75h.01a.75.75 0 0 0 .75-.75V12a.75.75 0 0 0-.75-.75z"/><path fill-rule="evenodd" d="M5.75 2a.75.75 0 0 1 .75.75V4h7V2.75a.75.75 0 0 1 1.5 0V4h.25A2.75 2.75 0 0 1 18 6.75v8.5A2.75 2.75 0 0 1 15.25 18H4.75A2.75 2.75 0 0 1 2 15.25v-8.5A2.75 2.75 0 0 1 4.75 4H5V2.75A.75.75 0 0 1 5.75 2m-1 5.5c-.69 0-1.25.56-1.25 1.25v6.5c0 .69.56 1.25 1.25 1.25h10.5c.69 0 1.25-.56 1.25-1.25v-6.5c0-.69-.56-1.25-1.25-1.25z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 1.75a.75.75 0 0 1 1.5 0V3h5V1.75a.75.75 0 0 1 1.5 0V3a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2zM4.5 6a1 1 0 0 0-1 1v4.5a1 1 0 0 0 1 1h7a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.75 2.25A.75.75 0 0 1 7.5 3v1.5h9V3A.75.75 0 0 1 18 3v1.5h.75a3 3 0 0 1 3 3v11.25a3 3 0 0 1-3 3H5.25a3 3 0 0 1-3-3V7.5a3 3 0 0 1 3-3H6V3a.75.75 0 0 1 .75-.75m13.5 9a1.5 1.5 0 0 0-1.5-1.5H5.25a1.5 1.5 0 0 0-1.5 1.5v7.5a1.5 1.5 0 0 0 1.5 1.5h13.5a1.5 1.5 0 0 0 1.5-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.75 2a.75.75 0 0 1 .75.75V4h7V2.75a.75.75 0 0 1 1.5 0V4h.25A2.75 2.75 0 0 1 18 6.75v8.5A2.75 2.75 0 0 1 15.25 18H4.75A2.75 2.75 0 0 1 2 15.25v-8.5A2.75 2.75 0 0 1 4.75 4H5V2.75A.75.75 0 0 1 5.75 2m-1 5.5c-.69 0-1.25.56-1.25 1.25v6.5c0 .69.56 1.25 1.25 1.25h10.5c.69 0 1.25-.56 1.25-1.25v-6.5c0-.69-.56-1.25-1.25-1.25z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.827 6.175A2.31 2.31 0 0 1 5.186 7.23c-.38.054-.757.112-1.134.175C2.999 7.58 2.25 8.507 2.25 9.574V18a2.25 2.25 0 0 0 2.25 2.25h15A2.25 2.25 0 0 0 21.75 18V9.574c0-1.067-.75-1.994-1.802-2.169a47.865 47.865 0 0 0-1.134-.175a2.31 2.31 0 0 1-1.64-1.055l-.822-1.316a2.192 2.192 0 0 0-1.736-1.039a48.774 48.774 0 0 0-5.232 0a2.192 2.192 0 0 0-1.736 1.039z"/><path d="M16.5 12.75a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0m2.25-2.25h.008v.008h-.008z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.5 8.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0"/><path fill="currentColor" fill-rule="evenodd" d="M2.5 5A1.5 1.5 0 0 0 1 6.5v5A1.5 1.5 0 0 0 2.5 13h11a1.5 1.5 0 0 0 1.5-1.5v-5A1.5 1.5 0 0 0 13.5 5h-.879a1.5 1.5 0 0 1-1.06-.44l-1.122-1.12A1.5 1.5 0 0 0 9.38 3H6.62a1.5 1.5 0 0 0-1.06.44L4.439 4.56A1.5 1.5 0 0 1 3.38 5zM11 8.5a3 3 0 1 1-6 0a3 3 0 0 1 6 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 9a3.75 3.75 0 1 0 0 7.5A3.75 3.75 0 0 0 12 9"/><path fill-rule="evenodd" d="M9.344 3.071a49.52 49.52 0 0 1 5.312 0c.967.052 1.83.585 2.332 1.39l.821 1.317c.24.383.645.643 1.11.71c.386.054.77.113 1.152.177c1.432.239 2.429 1.493 2.429 2.909V18a3 3 0 0 1-3 3h-15a3 3 0 0 1-3-3V9.574c0-1.416.997-2.67 2.429-2.909c.382-.064.766-.123 1.151-.178a1.56 1.56 0 0 0 1.11-.71l.822-1.315a2.942 2.942 0 0 1 2.332-1.39M6.75 12.75a5.25 5.25 0 1 1 10.5 0a5.25 5.25 0 0 1-10.5 0m12-1.5a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 8a2 2 0 0 1 2-2h.93a2 2 0 0 0 1.664-.89l.812-1.22A2 2 0 0 1 8.07 3h3.86a2 2 0 0 1 1.664.89l.812 1.22A2 2 0 0 0 16.07 6H17a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2zm13.5 3a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0M10 14a3 3 0 1 0 0-6a3 3 0 0 0 0 6" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartBar(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 13.125C3 12.504 3.504 12 4.125 12h2.25c.621 0 1.125.504 1.125 1.125v6.75C7.5 20.496 6.996 21 6.375 21h-2.25A1.125 1.125 0 0 1 3 19.875zm6.75-4.5c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125v11.25c0 .621-.504 1.125-1.125 1.125h-2.25a1.125 1.125 0 0 1-1.125-1.125zm6.75-4.5c0-.621.504-1.125 1.125-1.125h2.25C20.496 3 21 3.504 21 4.125v15.75c0 .621-.504 1.125-1.125 1.125h-2.25a1.125 1.125 0 0 1-1.125-1.125z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartBarSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1zM6.5 6a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1zM2 9a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartBarSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.375 2.25c-1.035 0-1.875.84-1.875 1.875v15.75c0 1.035.84 1.875 1.875 1.875h.75c1.035 0 1.875-.84 1.875-1.875V4.125c0-1.036-.84-1.875-1.875-1.875zM9.75 8.625c0-1.036.84-1.875 1.875-1.875h.75c1.036 0 1.875.84 1.875 1.875v11.25c0 1.035-.84 1.875-1.875 1.875h-.75a1.875 1.875 0 0 1-1.875-1.875zM3 13.125c0-1.036.84-1.875 1.875-1.875h.75c1.036 0 1.875.84 1.875 1.875v6.75c0 1.035-.84 1.875-1.875 1.875h-.75A1.875 1.875 0 0 1 3 19.875z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartBarSquare(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.5 14.25v2.25m3-4.5v4.5m3-6.75v6.75m3-9v9M6 20.25h12A2.25 2.25 0 0 0 20.25 18V6A2.25 2.25 0 0 0 18 3.75H6A2.25 2.25 0 0 0 3.75 6v12A2.25 2.25 0 0 0 6 20.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartBarSquareSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 2a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm.75 7a.75.75 0 0 0-.75.75v1.5a.75.75 0 0 0 1.5 0v-1.5A.75.75 0 0 0 4.75 9m2.5-1.75a.75.75 0 0 1 1.5 0v4a.75.75 0 0 1-1.5 0zm4-3.25a.75.75 0 0 0-.75.75v6.5a.75.75 0 0 0 1.5 0v-6.5a.75.75 0 0 0-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartBarSquareSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm4.5 7.5a.75.75 0 0 1 .75.75v2.25a.75.75 0 0 1-1.5 0v-2.25a.75.75 0 0 1 .75-.75m3.75-1.5a.75.75 0 0 0-1.5 0v4.5a.75.75 0 0 0 1.5 0zm2.25-3a.75.75 0 0 1 .75.75v6.75a.75.75 0 0 1-1.5 0V9.75A.75.75 0 0 1 13.5 9m3.75-1.5a.75.75 0 0 0-1.5 0v9a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartBarSquareTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.25 2A2.25 2.25 0 0 0 2 4.25v11.5A2.25 2.25 0 0 0 4.25 18h11.5A2.25 2.25 0 0 0 18 15.75V4.25A2.25 2.25 0 0 0 15.75 2zM15 5.75a.75.75 0 0 0-1.5 0v8.5a.75.75 0 0 0 1.5 0zm-8.5 6a.75.75 0 0 0-1.5 0v2.5a.75.75 0 0 0 1.5 0zM8.584 9a.75.75 0 0 1 .75.75v4.5a.75.75 0 0 1-1.5 0v-4.5a.75.75 0 0 1 .75-.75m3.58-1.25a.75.75 0 0 0-1.5 0v6.5a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartBarTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 2A1.5 1.5 0 0 0 14 3.5v13a1.5 1.5 0 0 0 1.5 1.5h1a1.5 1.5 0 0 0 1.5-1.5v-13A1.5 1.5 0 0 0 16.5 2zm-6 4A1.5 1.5 0 0 0 8 7.5v9A1.5 1.5 0 0 0 9.5 18h1a1.5 1.5 0 0 0 1.5-1.5v-9A1.5 1.5 0 0 0 10.5 6zm-6 4A1.5 1.5 0 0 0 2 11.5v5A1.5 1.5 0 0 0 3.5 18h1A1.5 1.5 0 0 0 6 16.5v-5A1.5 1.5 0 0 0 4.5 10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartPie(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.5 6a7.5 7.5 0 1 0 7.5 7.5h-7.5z"/><path d="M13.5 10.5H21A7.5 7.5 0 0 0 13.5 3z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartPieSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.975 6.5c.028.276-.199.5-.475.5h-4a.5.5 0 0 1-.5-.5v-4c0-.276.225-.503.5-.475A5.002 5.002 0 0 1 13.974 6.5"/><path fill="currentColor" d="M6.5 4.025c.276-.028.5.199.5.475v4a.5.5 0 0 0 .5.5h4c.276 0 .503.225.475.5a5 5 0 1 1-5.474-5.475"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartPieSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M2.25 13.5a8.25 8.25 0 0 1 8.25-8.25a.75.75 0 0 1 .75.75v6.75H18a.75.75 0 0 1 .75.75a8.25 8.25 0 0 1-16.5 0"/><path d="M12.75 3a.75.75 0 0 1 .75-.75a8.25 8.25 0 0 1 8.25 8.25a.75.75 0 0 1-.75.75h-7.5a.75.75 0 0 1-.75-.75z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartPieTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 9a1 1 0 0 1-1-1V3c0-.552.45-1.007.997-.93a7.004 7.004 0 0 1 5.933 5.933c.078.547-.378.997-.93.997z"/><path d="M8.003 4.07C8.55 3.994 9 4.449 9 5v5a1 1 0 0 0 1 1h5c.552 0 1.008.45.93.997A7.001 7.001 0 0 1 2 11a7.002 7.002 0 0 1 6.003-6.93"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleBottomCenter(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.25 12.76c0 1.6 1.123 2.994 2.707 3.227c1.068.157 2.148.279 3.238.364c.466.037.893.281 1.153.671L12 21l2.652-3.978c.26-.39.687-.634 1.153-.67c1.09-.086 2.17-.208 3.238-.365c1.584-.233 2.707-1.626 2.707-3.228V6.741c0-1.602-1.123-2.995-2.707-3.228A48.394 48.394 0 0 0 12 3c-2.392 0-4.744.175-7.043.513C3.373 3.746 2.25 5.14 2.25 6.741z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleBottomCenterSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 8.74c0 .983.713 1.825 1.69 1.943c.904.108 1.817.19 2.737.243c.363.02.688.231.85.556l1.052 2.103a.75.75 0 0 0 1.342 0l1.052-2.103c.162-.325.487-.535.85-.556c.92-.053 1.833-.134 2.738-.243c.976-.118 1.689-.96 1.689-1.942V4.259c0-.982-.713-1.824-1.69-1.942a44.45 44.45 0 0 0-10.62 0C1.712 2.435 1 3.277 1 4.26z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleBottomCenterSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.848 2.771A49.144 49.144 0 0 1 12 2.25c2.43 0 4.817.178 7.152.52c1.978.292 3.348 2.024 3.348 3.97v6.02c0 1.946-1.37 3.678-3.348 3.97a48.901 48.901 0 0 1-3.476.383a.39.39 0 0 0-.297.17l-2.755 4.133a.75.75 0 0 1-1.248 0l-2.755-4.132a.39.39 0 0 0-.297-.17a48.9 48.9 0 0 1-3.476-.385c-1.978-.29-3.348-2.024-3.348-3.97V6.741c0-1.946 1.37-3.68 3.348-3.97" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleBottomCenterText(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.5 8.25h9m-9 3H12m-9.75 1.51c0 1.6 1.123 2.994 2.707 3.227c1.129.166 2.27.293 3.423.379c.35.026.67.21.865.501L12 21l2.755-4.132a1.14 1.14 0 0 1 .865-.502a48.172 48.172 0 0 0 3.423-.379c1.584-.233 2.707-1.626 2.707-3.228V6.741c0-1.602-1.123-2.995-2.707-3.228A48.394 48.394 0 0 0 12 3c-2.392 0-4.744.175-7.043.513C3.373 3.746 2.25 5.14 2.25 6.741z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleBottomCenterTextSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 8.74c0 .983.713 1.825 1.69 1.943c.904.108 1.817.19 2.737.243c.363.02.688.231.85.556l1.052 2.103a.75.75 0 0 0 1.342 0l1.052-2.103c.162-.325.487-.535.85-.556c.92-.053 1.833-.134 2.738-.243c.976-.118 1.689-.96 1.689-1.942V4.259c0-.982-.713-1.824-1.69-1.942a44.45 44.45 0 0 0-10.62 0C1.712 2.435 1 3.277 1 4.26zm3-3.49a.75.75 0 0 1 .75-.75h6.5a.75.75 0 0 1 0 1.5h-6.5A.75.75 0 0 1 4 5.25M4.75 7a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleBottomCenterTextSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.848 2.771A49.144 49.144 0 0 1 12 2.25c2.43 0 4.817.178 7.152.52c1.978.292 3.348 2.024 3.348 3.97v6.02c0 1.946-1.37 3.678-3.348 3.97a48.901 48.901 0 0 1-3.476.383a.39.39 0 0 0-.297.17l-2.755 4.133a.75.75 0 0 1-1.248 0l-2.755-4.132a.39.39 0 0 0-.297-.17a48.9 48.9 0 0 1-3.476-.385c-1.978-.29-3.348-2.024-3.348-3.97V6.741c0-1.946 1.37-3.68 3.348-3.97M6.75 8.25a.75.75 0 0 1 .75-.75h9a.75.75 0 0 1 0 1.5h-9a.75.75 0 0 1-.75-.75m.75 2.25a.75.75 0 0 0 0 1.5H12a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleBottomCenterTextTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 2c-2.236 0-4.43.18-6.57.524C1.993 2.755 1 4.014 1 5.426v5.148c0 1.413.993 2.67 2.43 2.902c1.168.188 2.352.327 3.55.414c.28.02.521.18.642.413l1.713 3.293a.75.75 0 0 0 1.33 0l1.713-3.293a.783.783 0 0 1 .642-.413a41.102 41.102 0 0 0 3.55-.414c1.437-.231 2.43-1.49 2.43-2.902V5.426c0-1.413-.993-2.67-2.43-2.902A41.289 41.289 0 0 0 10 2M6.75 6a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5zm0 2.5a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleBottomCenterTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.43 2.524A41.29 41.29 0 0 1 10 2c2.236 0 4.43.18 6.57.524c1.437.231 2.43 1.49 2.43 2.902v5.148c0 1.413-.993 2.67-2.43 2.902a41.102 41.102 0 0 1-3.55.414a.785.785 0 0 0-.643.413l-1.712 3.293a.75.75 0 0 1-1.33 0l-1.713-3.293a.783.783 0 0 0-.642-.413a41.108 41.108 0 0 1-3.55-.414C1.993 13.245 1 11.986 1 10.574V5.426c0-1.413.993-2.67 2.43-2.902" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleLeft(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.25 12.76c0 1.6 1.123 2.994 2.707 3.227c1.087.16 2.185.283 3.293.369V21l4.076-4.076a1.526 1.526 0 0 1 1.037-.443a48.282 48.282 0 0 0 5.68-.494c1.584-.233 2.707-1.626 2.707-3.228V6.741c0-1.602-1.123-2.995-2.707-3.228A48.394 48.394 0 0 0 12 3c-2.392 0-4.744.175-7.043.513C3.373 3.746 2.25 5.14 2.25 6.741z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleLeftEllipsis(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.625 9.75a.375.375 0 1 1-.75 0a.375.375 0 0 1 .75 0m0 0H8.25m4.125 0a.375.375 0 1 1-.75 0a.375.375 0 0 1 .75 0m0 0H12m4.125 0a.375.375 0 1 1-.75 0a.375.375 0 0 1 .75 0m0 0h-.375m-13.5 3.01c0 1.6 1.123 2.994 2.707 3.227c1.087.16 2.185.283 3.293.369V21l4.184-4.183a1.14 1.14 0 0 1 .778-.332a48.294 48.294 0 0 0 5.83-.498c1.585-.233 2.708-1.626 2.708-3.228V6.741c0-1.602-1.123-2.995-2.707-3.228A48.394 48.394 0 0 0 12 3c-2.392 0-4.744.175-7.043.513C3.373 3.746 2.25 5.14 2.25 6.741z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleLeftEllipsisSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 8.74c0 .983.713 1.826 1.69 1.943a43.83 43.83 0 0 0 2.31.216v2.351a.75.75 0 0 0 1.28.53l2.51-2.51c.182-.181.427-.286.684-.294a44.298 44.298 0 0 0 3.837-.293C14.287 10.565 15 9.723 15 8.74V4.26c0-.983-.713-1.825-1.69-1.943a44.447 44.447 0 0 0-10.62 0C1.712 2.435 1 3.277 1 4.26zM5.5 6.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0m2.5 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2m3.5 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleLeftEllipsisSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25c-2.429 0-4.817.178-7.152.521C2.87 3.061 1.5 4.795 1.5 6.741v6.018c0 1.946 1.37 3.68 3.348 3.97c.877.129 1.761.234 2.652.316V21a.75.75 0 0 0 1.28.53l4.184-4.183a.39.39 0 0 1 .266-.112c2.006-.05 3.982-.22 5.922-.506c1.978-.29 3.348-2.023 3.348-3.97V6.741c0-1.947-1.37-3.68-3.348-3.97A49.145 49.145 0 0 0 12 2.25M8.25 8.625a1.125 1.125 0 1 0 0 2.25a1.125 1.125 0 0 0 0-2.25m2.625 1.125a1.125 1.125 0 1 1 2.25 0a1.125 1.125 0 0 1-2.25 0m4.875-1.125a1.125 1.125 0 1 0 0 2.25a1.125 1.125 0 0 0 0-2.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleLeftEllipsisTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 2c-2.236 0-4.43.18-6.57.524C1.993 2.755 1 4.014 1 5.426v5.148c0 1.413.993 2.67 2.43 2.902c.848.137 1.705.248 2.57.331v3.443a.75.75 0 0 0 1.28.53l3.58-3.579a.78.78 0 0 1 .527-.224a41.202 41.202 0 0 0 5.183-.5c1.437-.232 2.43-1.49 2.43-2.903V5.426c0-1.413-.993-2.67-2.43-2.902A41.289 41.289 0 0 0 10 2m0 7a1 1 0 1 0 0-2a1 1 0 0 0 0 2M8 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0m5 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleLeftRight(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.25 8.511c.884.284 1.5 1.128 1.5 2.097v4.286c0 1.136-.847 2.1-1.98 2.193c-.34.027-.68.052-1.02.072v3.091l-3-3a49.5 49.5 0 0 1-4.02-.163a2.115 2.115 0 0 1-.825-.242m9.345-8.334a2.126 2.126 0 0 0-.476-.095a48.64 48.64 0 0 0-8.048 0c-1.131.094-1.976 1.057-1.976 2.192v4.286c0 .837.46 1.58 1.155 1.951m9.345-8.334V6.637c0-1.621-1.152-3.026-2.76-3.235A48.455 48.455 0 0 0 11.25 3c-2.115 0-4.198.137-6.24.402c-1.608.209-2.76 1.614-2.76 3.235v6.226c0 1.621 1.152 3.026 2.76 3.235c.577.075 1.157.14 1.74.194V21l4.155-4.155"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleLeftRightSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 8.849c0 1 .738 1.851 1.734 1.947L3 10.82v2.429a.75.75 0 0 0 1.28.53l1.82-1.82A3.484 3.484 0 0 1 5.5 10V9A3.5 3.5 0 0 1 9 5.5h4V4.151c0-1-.739-1.851-1.734-1.947a44.539 44.539 0 0 0-8.532 0C1.738 2.3 1 3.151 1 4.151z"/><path fill="currentColor" d="M7 9a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2h-.25v1.25a.75.75 0 0 1-1.28.53L9.69 12H9a2 2 0 0 1-2-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleLeftRightSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M4.913 2.658a49.272 49.272 0 0 1 12.674 0c1.922.25 3.291 1.861 3.405 3.727a4.403 4.403 0 0 0-1.032-.211a50.89 50.89 0 0 0-8.42 0c-2.358.196-4.04 2.19-4.04 4.434v4.286a4.47 4.47 0 0 0 2.433 3.984L7.28 21.53A.75.75 0 0 1 6 21v-4.03a48.527 48.527 0 0 1-1.087-.128C2.905 16.58 1.5 14.833 1.5 12.862V6.638c0-1.97 1.405-3.718 3.413-3.979"/><path d="M15.75 7.5c-1.376 0-2.739.057-4.086.169C10.124 7.797 9 9.103 9 10.609v4.285c0 1.507 1.128 2.814 2.67 2.94c1.243.102 2.5.157 3.768.165l2.782 2.781a.75.75 0 0 0 1.28-.53v-2.39l.33-.026c1.542-.125 2.67-1.433 2.67-2.94v-4.286c0-1.505-1.125-2.811-2.664-2.94A49.392 49.392 0 0 0 15.75 7.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleLeftRightTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M3.505 2.365A41.369 41.369 0 0 1 9 2c1.863 0 3.697.124 5.495.365c1.247.167 2.18 1.108 2.435 2.268a4.45 4.45 0 0 0-.577-.069a43.141 43.141 0 0 0-4.706 0C9.229 4.696 7.5 6.727 7.5 8.998v2.24c0 1.413.67 2.735 1.76 3.562l-2.98 2.98A.75.75 0 0 1 5 17.25v-3.443c-.501-.048-1-.106-1.495-.172C2.033 13.438 1 12.162 1 10.72V5.28c0-1.441 1.033-2.717 2.505-2.914"/><path d="M14 6c-.762 0-1.52.02-2.271.062C10.157 6.148 9 7.472 9 8.998v2.24c0 1.519 1.147 2.839 2.71 2.935c.214.013.428.024.642.034c.2.009.385.09.518.224l2.35 2.35a.75.75 0 0 0 1.28-.531v-2.07c1.453-.195 2.5-1.463 2.5-2.915V8.998c0-1.526-1.157-2.85-2.729-2.936A41.645 41.645 0 0 0 14 6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleLeftSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 8.74c0 .983.713 1.826 1.69 1.943a43.83 43.83 0 0 0 2.31.216v2.351a.75.75 0 0 0 1.28.53l2.51-2.51c.182-.181.427-.286.684-.294a44.298 44.298 0 0 0 3.837-.293C14.287 10.565 15 9.723 15 8.74V4.26c0-.983-.713-1.825-1.69-1.943a44.447 44.447 0 0 0-10.62 0C1.712 2.435 1 3.277 1 4.26z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleLeftSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.848 2.771A49.144 49.144 0 0 1 12 2.25c2.43 0 4.817.178 7.152.52c1.978.292 3.348 2.024 3.348 3.97v6.02c0 1.946-1.37 3.678-3.348 3.97a49.36 49.36 0 0 1-5.922.505a.39.39 0 0 0-.266.112L8.78 21.53A.75.75 0 0 1 7.5 21v-3.955a48.842 48.842 0 0 1-2.652-.316c-1.978-.29-3.348-2.024-3.348-3.97V6.741c0-1.946 1.37-3.68 3.348-3.97" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleLeftTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.43 2.524A41.29 41.29 0 0 1 10 2c2.236 0 4.43.18 6.57.524c1.437.231 2.43 1.49 2.43 2.902v5.148c0 1.413-.993 2.67-2.43 2.902a41.202 41.202 0 0 1-5.183.501a.78.78 0 0 0-.528.224l-3.579 3.58A.75.75 0 0 1 6 17.25v-3.443a41.033 41.033 0 0 1-2.57-.33C1.993 13.244 1 11.986 1 10.573V5.426c0-1.413.993-2.67 2.43-2.902" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleOvalLeft(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 20.25c4.97 0 9-3.694 9-8.25s-4.03-8.25-9-8.25S3 7.444 3 12c0 2.104.859 4.023 2.273 5.48c.432.447.74 1.04.586 1.641a4.483 4.483 0 0 1-.923 1.785A5.969 5.969 0 0 0 6 21c1.282 0 2.47-.402 3.445-1.087c.81.22 1.668.337 2.555.337"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleOvalLeftEllipsis(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.625 12a.375.375 0 1 1-.75 0a.375.375 0 0 1 .75 0m0 0H8.25m4.125 0a.375.375 0 1 1-.75 0a.375.375 0 0 1 .75 0m0 0H12m4.125 0a.375.375 0 1 1-.75 0a.375.375 0 0 1 .75 0m0 0h-.375M21 12c0 4.556-4.03 8.25-9 8.25a9.764 9.764 0 0 1-2.555-.337A5.972 5.972 0 0 1 5.41 20.97a5.969 5.969 0 0 1-.474-.065a4.48 4.48 0 0 0 .978-2.025c.09-.457-.133-.901-.467-1.226C3.93 16.178 3 14.189 3 12c0-4.556 4.03-8.25 9-8.25s9 3.694 9 8.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleOvalLeftEllipsisSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 2C4.262 2 1 4.57 1 8c0 1.86.98 3.486 2.455 4.566a3.472 3.472 0 0 1-.469 1.26a.75.75 0 0 0 .713 1.14a6.961 6.961 0 0 0 3.06-1.06c.403.062.818.094 1.241.094c3.738 0 7-2.57 7-6s-3.262-6-7-6M5 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2m7-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0M8 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleOvalLeftEllipsisSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.804 21.644A6.707 6.707 0 0 0 6 21.75a6.721 6.721 0 0 0 3.583-1.029c.774.182 1.584.279 2.417.279c5.322 0 9.75-3.97 9.75-9c0-5.03-4.428-9-9.75-9s-9.75 3.97-9.75 9c0 2.409 1.025 4.587 2.674 6.192c.232.226.277.428.254.543a3.73 3.73 0 0 1-.814 1.686a.75.75 0 0 0 .44 1.223M8.25 10.875a1.125 1.125 0 1 0 0 2.25a1.125 1.125 0 0 0 0-2.25M10.875 12a1.125 1.125 0 1 1 2.25 0a1.125 1.125 0 0 1-2.25 0m4.875-1.125a1.125 1.125 0 1 0 0 2.25a1.125 1.125 0 0 0 0-2.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleOvalLeftEllipsisTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 3c-4.31 0-8 3.033-8 7c0 2.024.978 3.825 2.499 5.085a3.478 3.478 0 0 1-.522 1.756a.75.75 0 0 0 .584 1.143a5.976 5.976 0 0 0 3.936-1.108c.487.082.99.124 1.503.124c4.31 0 8-3.033 8-7s-3.69-7-8-7m0 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2m-2-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0m5 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleOvalLeftSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 8c0-3.43 3.262-6 7-6s7 2.57 7 6s-3.262 6-7 6c-.423 0-.838-.032-1.241-.094c-.9.574-1.941.948-3.06 1.06a.75.75 0 0 1-.713-1.14a3.47 3.47 0 0 0 .469-1.26C1.979 11.486 1 9.86 1 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleOvalLeftSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.337 21.718a6.707 6.707 0 0 1-.533-.074a.75.75 0 0 1-.44-1.223a3.73 3.73 0 0 0 .814-1.686c.023-.115-.022-.317-.254-.543C3.274 16.587 2.25 14.41 2.25 12c0-5.03 4.428-9 9.75-9s9.75 3.97 9.75 9c0 5.03-4.428 9-9.75 9c-.833 0-1.643-.097-2.417-.279a6.721 6.721 0 0 1-4.246.997" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleOvalLeftTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 10c0-3.967 3.69-7 8-7c4.31 0 8 3.033 8 7s-3.69 7-8 7a9.165 9.165 0 0 1-1.504-.123a5.976 5.976 0 0 1-3.935 1.107a.75.75 0 0 1-.584-1.143a3.478 3.478 0 0 0 .522-1.756C2.979 13.825 2 12.025 2 10" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4.5 12.75l6 6l9-13.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckBadge(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12.75L11.25 15L15 9.75M21 12c0 1.268-.63 2.39-1.593 3.068a3.745 3.745 0 0 1-1.043 3.296a3.745 3.745 0 0 1-3.296 1.043A3.745 3.745 0 0 1 12 21c-1.268 0-2.39-.63-3.068-1.593a3.746 3.746 0 0 1-3.296-1.043a3.745 3.745 0 0 1-1.043-3.296A3.745 3.745 0 0 1 3 12c0-1.268.63-2.39 1.593-3.068a3.745 3.745 0 0 1 1.043-3.296a3.746 3.746 0 0 1 3.296-1.043A3.746 3.746 0 0 1 12 3c1.268 0 2.39.63 3.068 1.593a3.746 3.746 0 0 1 3.296 1.043a3.746 3.746 0 0 1 1.043 3.296A3.745 3.745 0 0 1 21 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckBadgeSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 8c0 .982-.472 1.854-1.202 2.402a2.995 2.995 0 0 1-.848 2.547a2.995 2.995 0 0 1-2.548.849A2.996 2.996 0 0 1 8 15a2.996 2.996 0 0 1-2.402-1.202a2.995 2.995 0 0 1-2.547-.848a2.995 2.995 0 0 1-.849-2.548A2.996 2.996 0 0 1 1 8c0-.982.472-1.854 1.202-2.402a2.995 2.995 0 0 1 .848-2.547a2.995 2.995 0 0 1 2.548-.849A2.995 2.995 0 0 1 8 1c.982 0 1.854.472 2.402 1.202a2.995 2.995 0 0 1 2.547.848c.695.695.978 1.645.849 2.548A2.996 2.996 0 0 1 15 8m-3.291-2.843a.75.75 0 0 1 .135 1.052l-4.25 5.5a.75.75 0 0 1-1.151.043l-2.25-2.5a.75.75 0 1 1 1.114-1.004l1.65 1.832l3.7-4.789a.75.75 0 0 1 1.052-.134" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckBadgeSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.603 3.799A4.49 4.49 0 0 1 12 2.25c1.357 0 2.573.6 3.397 1.549a4.49 4.49 0 0 1 3.498 1.307a4.491 4.491 0 0 1 1.307 3.497A4.49 4.49 0 0 1 21.75 12a4.49 4.49 0 0 1-1.549 3.397a4.491 4.491 0 0 1-1.307 3.497a4.491 4.491 0 0 1-3.497 1.307A4.49 4.49 0 0 1 12 21.75a4.49 4.49 0 0 1-3.397-1.549a4.49 4.49 0 0 1-3.498-1.306a4.491 4.491 0 0 1-1.307-3.498A4.49 4.49 0 0 1 2.25 12c0-1.357.6-2.573 1.549-3.397a4.49 4.49 0 0 1 1.307-3.497a4.49 4.49 0 0 1 3.497-1.307m7.007 6.387a.75.75 0 1 0-1.22-.872l-3.236 4.53L9.53 12.22a.75.75 0 0 0-1.06 1.06l2.25 2.25a.75.75 0 0 0 1.14-.094z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckBadgeTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.403 12.652a3 3 0 0 0 0-5.304a3 3 0 0 0-3.75-3.751a3 3 0 0 0-5.305 0a3 3 0 0 0-3.751 3.75a3 3 0 0 0 0 5.305a3 3 0 0 0 3.75 3.751a3 3 0 0 0 5.305 0a3 3 0 0 0 3.751-3.75m-2.546-4.46a.75.75 0 0 0-1.214-.883L9.16 12.1l-1.88-1.88a.75.75 0 1 0-1.06 1.061l2.5 2.5a.75.75 0 0 0 1.137-.089z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircle(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12.75L11.25 15L15 9.75M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircleSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14m3.844-8.791a.75.75 0 0 0-1.187-.918l-3.7 4.79l-1.65-1.833a.75.75 0 1 0-1.114 1.004l2.25 2.5a.75.75 0 0 0 1.15-.043z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircleSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 12c0-5.385 4.365-9.75 9.75-9.75s9.75 4.365 9.75 9.75s-4.365 9.75-9.75 9.75S2.25 17.385 2.25 12m13.36-1.814a.75.75 0 1 0-1.22-.872l-3.236 4.53L9.53 12.22a.75.75 0 0 0-1.06 1.06l2.25 2.25a.75.75 0 0 0 1.14-.094z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircleTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 18a8 8 0 1 0 0-16a8 8 0 0 0 0 16m3.857-9.809a.75.75 0 0 0-1.214-.882l-3.483 4.79l-1.88-1.88a.75.75 0 1 0-1.06 1.061l2.5 2.5a.75.75 0 0 0 1.137-.089z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.416 3.376a.75.75 0 0 1 .208 1.04l-5 7.5a.75.75 0 0 1-1.154.114l-3-3a.75.75 0 0 1 1.06-1.06l2.353 2.353l4.493-6.74a.75.75 0 0 1 1.04-.207" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M19.916 4.626a.75.75 0 0 1 .208 1.04l-9 13.5a.75.75 0 0 1-1.154.114l-6-6a.75.75 0 0 1 1.06-1.06l5.353 5.353l8.493-12.74a.75.75 0 0 1 1.04-.207" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.705 4.153a.75.75 0 0 1 .142 1.052l-8 10.5a.75.75 0 0 1-1.127.075l-4.5-4.5a.75.75 0 0 1 1.06-1.06l3.894 3.893l7.48-9.817a.75.75 0 0 1 1.05-.143" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleDown(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4.5 5.25l7.5 7.5l7.5-7.5m-15 6l7.5 7.5l7.5-7.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleDownSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.47 12.78a.75.75 0 0 0 1.06 0l3.25-3.25a.75.75 0 0 0-1.06-1.06L8 11.19L5.28 8.47a.75.75 0 0 0-1.06 1.06zM4.22 4.53l3.25 3.25a.75.75 0 0 0 1.06 0l3.25-3.25a.75.75 0 0 0-1.06-1.06L8 6.19L5.28 3.47a.75.75 0 0 0-1.06 1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleDownSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M11.47 13.28a.75.75 0 0 0 1.06 0l7.5-7.5a.75.75 0 0 0-1.06-1.06L12 11.69L5.03 4.72a.75.75 0 0 0-1.06 1.06z"/><path d="M11.47 19.28a.75.75 0 0 0 1.06 0l7.5-7.5a.75.75 0 1 0-1.06-1.06L12 17.69l-6.97-6.97a.75.75 0 0 0-1.06 1.06z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleDownTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.47 15.28a.75.75 0 0 0 1.06 0l4.25-4.25a.75.75 0 1 0-1.06-1.06L10 13.69L6.28 9.97a.75.75 0 0 0-1.06 1.06zM5.22 6.03l4.25 4.25a.75.75 0 0 0 1.06 0l4.25-4.25a.75.75 0 0 0-1.06-1.06L10 8.69L6.28 4.97a.75.75 0 0 0-1.06 1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleLeft(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m18.75 4.5l-7.5 7.5l7.5 7.5m-6-15L5.25 12l7.5 7.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleLeftSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.22 7.595a.75.75 0 0 0 0 1.06l3.25 3.25a.75.75 0 0 0 1.06-1.06l-2.72-2.72l2.72-2.72a.75.75 0 0 0-1.06-1.06zm8.25-3.25l-3.25 3.25a.75.75 0 0 0 0 1.06l3.25 3.25a.75.75 0 1 0 1.06-1.06l-2.72-2.72l2.72-2.72a.75.75 0 0 0-1.06-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleLeftSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M10.72 11.47a.75.75 0 0 0 0 1.06l7.5 7.5a.75.75 0 1 0 1.06-1.06L12.31 12l6.97-6.97a.75.75 0 0 0-1.06-1.06z"/><path d="M4.72 11.47a.75.75 0 0 0 0 1.06l7.5 7.5a.75.75 0 1 0 1.06-1.06L6.31 12l6.97-6.97a.75.75 0 0 0-1.06-1.06z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleLeftTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.72 9.47a.75.75 0 0 0 0 1.06l4.25 4.25a.75.75 0 1 0 1.06-1.06L6.31 10l3.72-3.72a.75.75 0 1 0-1.06-1.06zm9.25-4.25L9.72 9.47a.75.75 0 0 0 0 1.06l4.25 4.25a.75.75 0 1 0 1.06-1.06L11.31 10l3.72-3.72a.75.75 0 0 0-1.06-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleRight(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5.25 4.5l7.5 7.5l-7.5 7.5m6-15l7.5 7.5l-7.5 7.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleRightSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.78 7.595a.75.75 0 0 1 0 1.06l-3.25 3.25a.75.75 0 0 1-1.06-1.06l2.72-2.72l-2.72-2.72a.75.75 0 0 1 1.06-1.06zm-8.25-3.25l3.25 3.25a.75.75 0 0 1 0 1.06l-3.25 3.25a.75.75 0 0 1-1.06-1.06l2.72-2.72l-2.72-2.72a.75.75 0 0 1 1.06-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleRightSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M13.28 11.47a.75.75 0 0 1 0 1.06l-7.5 7.5a.75.75 0 0 1-1.06-1.06L11.69 12L4.72 5.03a.75.75 0 0 1 1.06-1.06z"/><path d="M19.28 11.47a.75.75 0 0 1 0 1.06l-7.5 7.5a.75.75 0 1 1-1.06-1.06L17.69 12l-6.97-6.97a.75.75 0 0 1 1.06-1.06z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleRightTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.28 9.47a.75.75 0 0 1 0 1.06l-4.25 4.25a.75.75 0 1 1-1.06-1.06L13.69 10L9.97 6.28a.75.75 0 0 1 1.06-1.06zM6.03 5.22l4.25 4.25a.75.75 0 0 1 0 1.06l-4.25 4.25a.75.75 0 0 1-1.06-1.06L8.69 10L4.97 6.28a.75.75 0 0 1 1.06-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleUp(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m4.5 18.75l7.5-7.5l7.5 7.5"/><path d="m4.5 12.75l7.5-7.5l7.5 7.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleUpSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.47 3.22a.75.75 0 0 1 1.06 0l3.25 3.25a.75.75 0 0 1-1.06 1.06L8 4.81L5.28 7.53a.75.75 0 0 1-1.06-1.06zm-3.25 8.25l3.25-3.25a.75.75 0 0 1 1.06 0l3.25 3.25a.75.75 0 1 1-1.06 1.06L8 9.81l-2.72 2.72a.75.75 0 0 1-1.06-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleUpSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M11.47 10.72a.75.75 0 0 1 1.06 0l7.5 7.5a.75.75 0 1 1-1.06 1.06L12 12.31l-6.97 6.97a.75.75 0 0 1-1.06-1.06z"/><path d="M11.47 4.72a.75.75 0 0 1 1.06 0l7.5 7.5a.75.75 0 1 1-1.06 1.06L12 6.31l-6.97 6.97a.75.75 0 0 1-1.06-1.06z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleUpTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.47 4.72a.75.75 0 0 1 1.06 0l4.25 4.25a.75.75 0 1 1-1.06 1.06L10 6.31l-3.72 3.72a.75.75 0 1 1-1.06-1.06zm-4.25 9.25l4.25-4.25a.75.75 0 0 1 1.06 0l4.25 4.25a.75.75 0 1 1-1.06 1.06L10 11.31l-3.72 3.72a.75.75 0 0 1-1.06-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDown(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m19.5 8.25l-7.5 7.5l-7.5-7.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.22 6.22a.75.75 0 0 1 1.06 0L8 8.94l2.72-2.72a.75.75 0 1 1 1.06 1.06l-3.25 3.25a.75.75 0 0 1-1.06 0L4.22 7.28a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.53 16.28a.75.75 0 0 1-1.06 0l-7.5-7.5a.75.75 0 0 1 1.06-1.06L12 14.69l6.97-6.97a.75.75 0 1 1 1.06 1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.22 8.22a.75.75 0 0 1 1.06 0L10 11.94l3.72-3.72a.75.75 0 1 1 1.06 1.06l-4.25 4.25a.75.75 0 0 1-1.06 0L5.22 9.28a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeft(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.75 19.5L8.25 12l7.5-7.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeftSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.78 4.22a.75.75 0 0 1 0 1.06L7.06 8l2.72 2.72a.75.75 0 1 1-1.06 1.06L5.47 8.53a.75.75 0 0 1 0-1.06l3.25-3.25a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeftSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.72 12.53a.75.75 0 0 1 0-1.06l7.5-7.5a.75.75 0 1 1 1.06 1.06L9.31 12l6.97 6.97a.75.75 0 1 1-1.06 1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeftTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.78 5.22a.75.75 0 0 1 0 1.06L8.06 10l3.72 3.72a.75.75 0 1 1-1.06 1.06l-4.25-4.25a.75.75 0 0 1 0-1.06l4.25-4.25a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRight(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8.25 4.5l7.5 7.5l-7.5 7.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRightSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.22 4.22a.75.75 0 0 1 1.06 0l3.25 3.25a.75.75 0 0 1 0 1.06l-3.25 3.25a.75.75 0 0 1-1.06-1.06L8.94 8L6.22 5.28a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRightSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.28 11.47a.75.75 0 0 1 0 1.06l-7.5 7.5a.75.75 0 0 1-1.06-1.06L14.69 12L7.72 5.03a.75.75 0 0 1 1.06-1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRightTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.22 5.22a.75.75 0 0 1 1.06 0l4.25 4.25a.75.75 0 0 1 0 1.06l-4.25 4.25a.75.75 0 0 1-1.06-1.06L11.94 10L8.22 6.28a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUp(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4.5 15.75l7.5-7.5l7.5 7.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpDown(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.25 15L12 18.75L15.75 15m-7.5-6L12 5.25L15.75 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpDownSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.22 10.22a.75.75 0 0 1 1.06 0L8 11.94l1.72-1.72a.75.75 0 1 1 1.06 1.06l-2.25 2.25a.75.75 0 0 1-1.06 0l-2.25-2.25a.75.75 0 0 1 0-1.06m5.56-4.44a.75.75 0 0 1-1.06 0L8 4.06L6.28 5.78a.75.75 0 0 1-1.06-1.06l2.25-2.25a.75.75 0 0 1 1.06 0l2.25 2.25a.75.75 0 0 1 0 1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpDownSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.47 4.72a.75.75 0 0 1 1.06 0l3.75 3.75a.75.75 0 0 1-1.06 1.06L12 6.31L8.78 9.53a.75.75 0 0 1-1.06-1.06zm-3.75 9.75a.75.75 0 0 1 1.06 0L12 17.69l3.22-3.22a.75.75 0 1 1 1.06 1.06l-3.75 3.75a.75.75 0 0 1-1.06 0l-3.75-3.75a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpDownTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.53 3.47a.75.75 0 0 0-1.06 0L6.22 6.72a.75.75 0 0 0 1.06 1.06L10 5.06l2.72 2.72a.75.75 0 1 0 1.06-1.06zm-4.31 9.81l3.25 3.25a.75.75 0 0 0 1.06 0l3.25-3.25a.75.75 0 1 0-1.06-1.06L10 14.94l-2.72-2.72a.75.75 0 0 0-1.06 1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.78 9.78a.75.75 0 0 1-1.06 0L8 7.06L5.28 9.78a.75.75 0 0 1-1.06-1.06l3.25-3.25a.75.75 0 0 1 1.06 0l3.25 3.25a.75.75 0 0 1 0 1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.47 7.72a.75.75 0 0 1 1.06 0l7.5 7.5a.75.75 0 1 1-1.06 1.06L12 9.31l-6.97 6.97a.75.75 0 0 1-1.06-1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.47 6.47a.75.75 0 0 1 1.06 0l4.25 4.25a.75.75 0 1 1-1.06 1.06L10 8.06l-3.72 3.72a.75.75 0 0 1-1.06-1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleStack(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.25 6.375c0 2.278-3.694 4.125-8.25 4.125S3.75 8.653 3.75 6.375m16.5 0c0-2.278-3.694-4.125-8.25-4.125S3.75 4.097 3.75 6.375m16.5 0v11.25c0 2.278-3.694 4.125-8.25 4.125s-8.25-1.847-8.25-4.125V6.375m16.5 0v3.75m-16.5-3.75v3.75m16.5 0v3.75C20.25 16.153 16.556 18 12 18s-8.25-1.847-8.25-4.125v-3.75m16.5 0c0 2.278-3.694 4.125-8.25 4.125s-8.25-1.847-8.25-4.125"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleStackSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 7c3.314 0 6-1.343 6-3s-2.686-3-6-3s-6 1.343-6 3s2.686 3 6 3"/><path fill="currentColor" d="M8 8.5c1.84 0 3.579-.37 4.914-1.037A6.33 6.33 0 0 0 14 6.78V8c0 1.657-2.686 3-6 3S2 9.657 2 8V6.78c.346.273.72.5 1.087.683C4.42 8.131 6.16 8.5 8 8.5"/><path fill="currentColor" d="M8 12.5c1.84 0 3.579-.37 4.914-1.037c.366-.183.74-.41 1.086-.684V12c0 1.657-2.686 3-6 3s-6-1.343-6-3v-1.22c.346.273.72.5 1.087.683C4.42 12.131 6.16 12.5 8 12.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleStackSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M21 6.375c0 2.692-4.03 4.875-9 4.875S3 9.067 3 6.375S7.03 1.5 12 1.5s9 2.183 9 4.875"/><path d="M12 12.75c2.685 0 5.19-.586 7.078-1.609a8.283 8.283 0 0 0 1.897-1.384c.016.121.025.244.025.368C21 12.817 16.97 15 12 15s-9-2.183-9-4.875c0-.124.009-.247.025-.368a8.285 8.285 0 0 0 1.897 1.384C6.809 12.164 9.315 12.75 12 12.75"/><path d="M12 16.5c2.685 0 5.19-.586 7.078-1.609a8.282 8.282 0 0 0 1.897-1.384c.016.121.025.244.025.368c0 2.692-4.03 4.875-9 4.875s-9-2.183-9-4.875c0-.124.009-.247.025-.368a8.284 8.284 0 0 0 1.897 1.384C6.809 15.914 9.315 16.5 12 16.5"/><path d="M12 20.25c2.685 0 5.19-.586 7.078-1.609a8.282 8.282 0 0 0 1.897-1.384c.016.121.025.244.025.368c0 2.692-4.03 4.875-9 4.875s-9-2.183-9-4.875c0-.124.009-.247.025-.368a8.284 8.284 0 0 0 1.897 1.384C6.809 19.664 9.315 20.25 12 20.25"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleStackTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 1c3.866 0 7 1.79 7 4s-3.134 4-7 4s-7-1.79-7-4s3.134-4 7-4m5.694 8.13c.464-.264.91-.583 1.306-.952V10c0 2.21-3.134 4-7 4s-7-1.79-7-4V8.178a7 7 0 0 0 1.306.953C5.838 10.006 7.854 10.5 10 10.5s4.162-.494 5.694-1.37M3 13.179V15c0 2.21 3.134 4 7 4s7-1.79 7-4v-1.822a7 7 0 0 1-1.306.953C14.162 15.006 12.146 15.5 10 15.5s-4.162-.494-5.694-1.37A7.009 7.009 0 0 1 3 13.179" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clipboard(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.666 3.888A2.25 2.25 0 0 0 13.5 2.25h-3c-1.03 0-1.9.693-2.166 1.638m7.332 0c.055.194.084.4.084.612v0a.75.75 0 0 1-.75.75H9a.75.75 0 0 1-.75-.75v0c0-.212.03-.418.084-.612m7.332 0c.646.049 1.288.11 1.927.184c1.1.128 1.907 1.077 1.907 2.185V19.5a2.25 2.25 0 0 1-2.25 2.25H6.75A2.25 2.25 0 0 1 4.5 19.5V6.257c0-1.108.806-2.057 1.907-2.185a48.208 48.208 0 0 1 1.927-.184"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardDocument(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.25 7.5V6.108c0-1.135.845-2.098 1.976-2.192c.373-.03.748-.057 1.124-.08M15.75 18H18a2.25 2.25 0 0 0 2.25-2.25V6.108c0-1.135-.845-2.098-1.976-2.192a48.424 48.424 0 0 0-1.123-.08M15.75 18.75v-1.875a3.375 3.375 0 0 0-3.375-3.375h-1.5a1.125 1.125 0 0 1-1.125-1.125v-1.5A3.375 3.375 0 0 0 6.375 7.5H5.25m11.9-3.664A2.251 2.251 0 0 0 15 2.25h-1.5a2.251 2.251 0 0 0-2.15 1.586m5.8 0c.065.21.1.433.1.664v.75h-6V4.5c0-.231.035-.454.1-.664M6.75 7.5H4.875c-.621 0-1.125.504-1.125 1.125v12c0 .621.504 1.125 1.125 1.125h9.75c.621 0 1.125-.504 1.125-1.125V16.5a9 9 0 0 0-9-9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardDocumentCheck(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.35 3.836c-.065.21-.1.433-.1.664c0 .414.336.75.75.75h4.5a.75.75 0 0 0 .75-.75a2.25 2.25 0 0 0-.1-.664m-5.8 0A2.251 2.251 0 0 1 13.5 2.25H15a2.25 2.25 0 0 1 2.15 1.586m-5.8 0c-.376.023-.75.05-1.124.08C9.095 4.01 8.25 4.973 8.25 6.108V8.25m8.9-4.414c.376.023.75.05 1.124.08c1.131.094 1.976 1.057 1.976 2.192V16.5A2.25 2.25 0 0 1 18 18.75h-2.25m-7.5-10.5H4.875c-.621 0-1.125.504-1.125 1.125v11.25c0 .621.504 1.125 1.125 1.125h9.75c.621 0 1.125-.504 1.125-1.125V18.75m-7.5-10.5h6.375c.621 0 1.125.504 1.125 1.125v9.375m-8.25-3l1.5 1.5l3-3.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardDocumentCheckSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M11.986 3H12a2 2 0 0 1 2 2v6a2 2 0 0 1-1.5 1.937V7A2.5 2.5 0 0 0 10 4.5H4.063A2 2 0 0 1 6 3h.014A2.25 2.25 0 0 1 8.25 1h1.5a2.25 2.25 0 0 1 2.236 2M10.5 4v-.75a.75.75 0 0 0-.75-.75h-1.5a.75.75 0 0 0-.75.75V4z"/><path d="M2 7a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1zm6.585 1.08a.75.75 0 0 1 .336 1.005l-1.75 3.5a.75.75 0 0 1-1.16.234l-1.75-1.5a.75.75 0 0 1 .977-1.139l1.02.875l1.321-2.64a.75.75 0 0 1 1.006-.336"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardDocumentCheckSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M7.502 6h7.128A3.375 3.375 0 0 1 18 9.375v9.375a3 3 0 0 0 3-3V6.108c0-1.505-1.125-2.811-2.664-2.94a48.972 48.972 0 0 0-.673-.05A3 3 0 0 0 15 1.5h-1.5a3 3 0 0 0-2.663 1.618c-.225.015-.45.032-.673.05C8.662 3.295 7.554 4.542 7.502 6M13.5 3A1.5 1.5 0 0 0 12 4.5h4.5A1.5 1.5 0 0 0 15 3z"/><path d="M3 9.375C3 8.339 3.84 7.5 4.875 7.5h9.75c1.036 0 1.875.84 1.875 1.875v11.25c0 1.035-.84 1.875-1.875 1.875h-9.75A1.875 1.875 0 0 1 3 20.625zm9.586 4.594a.75.75 0 0 0-1.172-.937l-2.476 3.095l-.908-.907a.75.75 0 0 0-1.06 1.06l1.5 1.5a.75.75 0 0 0 1.116-.062z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardDocumentCheckTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M18 5.25a2.25 2.25 0 0 0-2.012-2.238A2.25 2.25 0 0 0 13.75 1h-1.5a2.25 2.25 0 0 0-2.238 2.012c-.875.092-1.6.686-1.884 1.488H11A2.5 2.5 0 0 1 13.5 7v7h2.25A2.25 2.25 0 0 0 18 11.75zM12.25 2.5a.75.75 0 0 0-.75.75v.25h3v-.25a.75.75 0 0 0-.75-.75z"/><path d="M3 6a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1zm6.874 4.166a.75.75 0 1 0-1.248-.832l-2.493 3.739l-.853-.853a.75.75 0 0 0-1.06 1.06l1.5 1.5a.75.75 0 0 0 1.154-.114z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardDocumentList(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12h3.75M9 15h3.75M9 18h3.75m3 .75H18a2.25 2.25 0 0 0 2.25-2.25V6.108c0-1.135-.845-2.098-1.976-2.192a48.424 48.424 0 0 0-1.123-.08m-5.801 0c-.065.21-.1.433-.1.664c0 .414.336.75.75.75h4.5a.75.75 0 0 0 .75-.75a2.25 2.25 0 0 0-.1-.664m-5.8 0A2.251 2.251 0 0 1 13.5 2.25H15a2.25 2.25 0 0 1 2.15 1.586m-5.8 0c-.376.023-.75.05-1.124.08C9.095 4.01 8.25 4.973 8.25 6.108V8.25m0 0H4.875c-.621 0-1.125.504-1.125 1.125v11.25c0 .621.504 1.125 1.125 1.125h9.75c.621 0 1.125-.504 1.125-1.125V9.375c0-.621-.504-1.125-1.125-1.125zM6.75 12h.008v.008H6.75zm0 3h.008v.008H6.75zm0 3h.008v.008H6.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardDocumentListSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M11.986 3H12a2 2 0 0 1 2 2v6a2 2 0 0 1-1.5 1.937V7A2.5 2.5 0 0 0 10 4.5H4.063A2 2 0 0 1 6 3h.014A2.25 2.25 0 0 1 8.25 1h1.5a2.25 2.25 0 0 1 2.236 2M10.5 4v-.75a.75.75 0 0 0-.75-.75h-1.5a.75.75 0 0 0-.75.75V4z"/><path d="M3 6a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h7a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1zm1.75 2.5a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5zM4 11.75a.75.75 0 0 1 .75-.75h3.5a.75.75 0 0 1 0 1.5h-3.5a.75.75 0 0 1-.75-.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardDocumentListSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M7.502 6h7.128A3.375 3.375 0 0 1 18 9.375v9.375a3 3 0 0 0 3-3V6.108c0-1.505-1.125-2.811-2.664-2.94a48.972 48.972 0 0 0-.673-.05A3 3 0 0 0 15 1.5h-1.5a3 3 0 0 0-2.663 1.618c-.225.015-.45.032-.673.05C8.662 3.295 7.554 4.542 7.502 6M13.5 3A1.5 1.5 0 0 0 12 4.5h4.5A1.5 1.5 0 0 0 15 3z"/><path d="M3 9.375C3 8.339 3.84 7.5 4.875 7.5h9.75c1.036 0 1.875.84 1.875 1.875v11.25c0 1.035-.84 1.875-1.875 1.875h-9.75A1.875 1.875 0 0 1 3 20.625zM6 12a.75.75 0 0 1 .75-.75h.008a.75.75 0 0 1 .75.75v.008a.75.75 0 0 1-.75.75H6.75a.75.75 0 0 1-.75-.75zm2.25 0a.75.75 0 0 1 .75-.75h3.75a.75.75 0 0 1 0 1.5H9a.75.75 0 0 1-.75-.75M6 15a.75.75 0 0 1 .75-.75h.008a.75.75 0 0 1 .75.75v.008a.75.75 0 0 1-.75.75H6.75a.75.75 0 0 1-.75-.75zm2.25 0a.75.75 0 0 1 .75-.75h3.75a.75.75 0 0 1 0 1.5H9a.75.75 0 0 1-.75-.75M6 18a.75.75 0 0 1 .75-.75h.008a.75.75 0 0 1 .75.75v.008a.75.75 0 0 1-.75.75H6.75a.75.75 0 0 1-.75-.75zm2.25 0a.75.75 0 0 1 .75-.75h3.75a.75.75 0 0 1 0 1.5H9a.75.75 0 0 1-.75-.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardDocumentListTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M15.988 3.012A2.25 2.25 0 0 1 18 5.25v6.5A2.25 2.25 0 0 1 15.75 14H13.5V7A2.5 2.5 0 0 0 11 4.5H8.128a2.252 2.252 0 0 1 1.884-1.488A2.25 2.25 0 0 1 12.25 1h1.5a2.25 2.25 0 0 1 2.238 2.012M11.5 3.25a.75.75 0 0 1 .75-.75h1.5a.75.75 0 0 1 .75.75v.25h-3z"/><path d="M2 7a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1zm2 3.25a.75.75 0 0 1 .75-.75h4.5a.75.75 0 0 1 0 1.5h-4.5a.75.75 0 0 1-.75-.75m0 3.5a.75.75 0 0 1 .75-.75h4.5a.75.75 0 0 1 0 1.5h-4.5a.75.75 0 0 1-.75-.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardDocumentSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.986 3H12a2 2 0 0 1 2 2v6a2 2 0 0 1-1.5 1.937v-2.523a2.5 2.5 0 0 0-.732-1.768L8.354 5.232A2.5 2.5 0 0 0 6.586 4.5H4.063A2 2 0 0 1 6 3h.014A2.25 2.25 0 0 1 8.25 1h1.5a2.25 2.25 0 0 1 2.236 2M10.5 4v-.75a.75.75 0 0 0-.75-.75h-1.5a.75.75 0 0 0-.75.75V4z" clip-rule="evenodd"/><path fill="currentColor" d="M3 6a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h7a1 1 0 0 0 1-1v-3.586a1 1 0 0 0-.293-.707L7.293 6.293A1 1 0 0 0 6.586 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardDocumentSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M17.663 3.118c.225.015.45.032.673.05C19.876 3.298 21 4.604 21 6.109v9.642a3 3 0 0 1-3 3V16.5c0-5.922-4.576-10.775-10.384-11.217c.324-1.132 1.3-2.01 2.548-2.114c.224-.019.448-.036.673-.051A3 3 0 0 1 13.5 1.5H15a3 3 0 0 1 2.663 1.618M12 4.5A1.5 1.5 0 0 1 13.5 3H15a1.5 1.5 0 0 1 1.5 1.5z" clip-rule="evenodd"/><path d="M3 8.625c0-1.036.84-1.875 1.875-1.875h.375A3.75 3.75 0 0 1 9 10.5v1.875c0 1.036.84 1.875 1.875 1.875h1.875A3.75 3.75 0 0 1 16.5 18v2.625c0 1.035-.84 1.875-1.875 1.875h-9.75A1.875 1.875 0 0 1 3 20.625z"/><path d="M10.5 10.5a5.23 5.23 0 0 0-1.279-3.434a9.768 9.768 0 0 1 6.963 6.963a5.23 5.23 0 0 0-3.434-1.279h-1.875a.375.375 0 0 1-.375-.375z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardDocumentTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M15.988 3.012A2.25 2.25 0 0 1 18 5.25v6.5A2.25 2.25 0 0 1 15.75 14H13.5v-3.379a3 3 0 0 0-.879-2.121l-3.12-3.121a3 3 0 0 0-1.402-.791a2.252 2.252 0 0 1 1.913-1.576A2.25 2.25 0 0 1 12.25 1h1.5a2.25 2.25 0 0 1 2.238 2.012M11.5 3.25a.75.75 0 0 1 .75-.75h1.5a.75.75 0 0 1 .75.75v.25h-3z" clip-rule="evenodd"/><path d="M3.5 6A1.5 1.5 0 0 0 2 7.5v9A1.5 1.5 0 0 0 3.5 18h7a1.5 1.5 0 0 0 1.5-1.5v-5.879a1.5 1.5 0 0 0-.44-1.06L8.44 6.439A1.5 1.5 0 0 0 7.378 6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.986 3H12a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h1.014A2.25 2.25 0 0 1 7.25 1h1.5a2.25 2.25 0 0 1 2.236 2M9.5 4v-.75a.75.75 0 0 0-.75-.75h-1.5a.75.75 0 0 0-.75.75V4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.5 3A1.501 1.501 0 0 0 9 4.5h6A1.5 1.5 0 0 0 13.5 3zm-2.693.178A3 3 0 0 1 10.5 1.5h3a3 3 0 0 1 2.694 1.678c.497.042.992.092 1.486.15c1.497.173 2.57 1.46 2.57 2.929V19.5a3 3 0 0 1-3 3H6.75a3 3 0 0 1-3-3V6.257c0-1.47 1.073-2.756 2.57-2.93c.493-.057.989-.107 1.487-.15" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.887 3.182c.396.037.79.08 1.183.128A2.213 2.213 0 0 1 17 5.517V16.75A2.25 2.25 0 0 1 14.75 19h-9.5A2.25 2.25 0 0 1 3 16.75V5.517c0-1.103.806-2.068 1.93-2.207c.393-.048.787-.09 1.183-.128A3.001 3.001 0 0 1 9 1h2c1.373 0 2.531.923 2.887 2.182M7.5 4A1.5 1.5 0 0 1 9 2.5h2A1.5 1.5 0 0 1 12.5 4v.5h-5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 6v6h4.5m4.5 0a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 8a7 7 0 1 1 14 0A7 7 0 0 1 1 8m7.75-4.25a.75.75 0 0 0-1.5 0V8c0 .414.336.75.75.75h3.25a.75.75 0 0 0 0-1.5h-2.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75s9.75-4.365 9.75-9.75S17.385 2.25 12 2.25M12.75 6a.75.75 0 0 0-1.5 0v6c0 .414.336.75.75.75h4.5a.75.75 0 0 0 0-1.5h-3.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 18a8 8 0 1 0 0-16a8 8 0 0 0 0 16m.75-13a.75.75 0 0 0-1.5 0v5c0 .414.336.75.75.75h4a.75.75 0 0 0 0-1.5h-3.25z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.25 15a4.5 4.5 0 0 0 4.5 4.5H18a3.75 3.75 0 0 0 1.332-7.257a3 3 0 0 0-3.758-3.848a5.25 5.25 0 0 0-10.233 2.33A4.502 4.502 0 0 0 2.25 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudArrowDown(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 9.75v6.75m0 0l-3-3m3 3l3-3m-8.25 6a4.5 4.5 0 0 1-1.41-8.775a5.25 5.25 0 0 1 10.233-2.33a3 3 0 0 1 3.758 3.848A3.752 3.752 0 0 1 18 19.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudArrowDownSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 13a3.5 3.5 0 0 1-1.41-6.705a3.5 3.5 0 0 1 6.63-2.171a2.5 2.5 0 0 1 3.197 3.018A3.001 3.001 0 0 1 12 13zm6.28-3.97a.75.75 0 1 0-1.06-1.06l-.97.97V6.25a.75.75 0 0 0-1.5 0v2.69l-.97-.97a.75.75 0 0 0-1.06 1.06l2.25 2.25a.75.75 0 0 0 1.06 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudArrowDownSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.5 3.75a6 6 0 0 0-5.98 6.496A5.25 5.25 0 0 0 6.75 20.25H18a4.5 4.5 0 0 0 2.206-8.423a3.75 3.75 0 0 0-4.132-4.303A6.001 6.001 0 0 0 10.5 3.75m2.25 6a.75.75 0 0 0-1.5 0v4.94l-1.72-1.72a.75.75 0 0 0-1.06 1.06l3 3a.75.75 0 0 0 1.06 0l3-3a.75.75 0 1 0-1.06-1.06l-1.72 1.72z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudArrowDownTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.5 17a4.5 4.5 0 0 1-1.44-8.765a4.5 4.5 0 0 1 8.302-3.046a3.5 3.5 0 0 1 4.504 4.272A4 4 0 0 1 15 17zm5.25-9.25a.75.75 0 0 0-1.5 0v4.59l-1.95-2.1a.75.75 0 1 0-1.1 1.02l3.25 3.5a.75.75 0 0 0 1.1 0l3.25-3.5a.75.75 0 1 0-1.1-1.02l-1.95 2.1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudArrowUp(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 16.5V9.75m0 0l3 3m-3-3l-3 3M6.75 19.5a4.5 4.5 0 0 1-1.41-8.775a5.25 5.25 0 0 1 10.233-2.33a3 3 0 0 1 3.758 3.848A3.752 3.752 0 0 1 18 19.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudArrowUpSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 13a3.5 3.5 0 0 1-1.41-6.705a3.5 3.5 0 0 1 6.63-2.171a2.5 2.5 0 0 1 3.197 3.018A3.001 3.001 0 0 1 12 13zm.72-5.03a.75.75 0 0 0 1.06 1.06l.97-.97v2.69a.75.75 0 0 0 1.5 0V8.06l.97.97a.75.75 0 1 0 1.06-1.06L8.53 5.72a.75.75 0 0 0-1.06 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudArrowUpSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.5 3.75a6 6 0 0 0-5.98 6.496A5.25 5.25 0 0 0 6.75 20.25H18a4.5 4.5 0 0 0 2.206-8.423a3.75 3.75 0 0 0-4.132-4.303A6.001 6.001 0 0 0 10.5 3.75m2.03 5.47a.75.75 0 0 0-1.06 0l-3 3a.75.75 0 1 0 1.06 1.06l1.72-1.72v4.94a.75.75 0 0 0 1.5 0v-4.94l1.72 1.72a.75.75 0 1 0 1.06-1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudArrowUpTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.5 17a4.5 4.5 0 0 1-1.44-8.765a4.5 4.5 0 0 1 8.302-3.046a3.5 3.5 0 0 1 4.504 4.272A4 4 0 0 1 15 17zm3.75-2.75a.75.75 0 0 0 1.5 0V9.66l1.95 2.1a.75.75 0 1 0 1.1-1.02l-3.25-3.5a.75.75 0 0 0-1.1 0l-3.25 3.5a.75.75 0 1 0 1.1 1.02l1.95-2.1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 9.5A3.5 3.5 0 0 0 4.5 13H12a3 3 0 0 0 .917-5.857a2.503 2.503 0 0 0-3.198-3.019a3.5 3.5 0 0 0-6.628 2.171A3.5 3.5 0 0 0 1 9.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 9.75a6 6 0 0 1 11.573-2.226a3.75 3.75 0 0 1 4.133 4.303A4.5 4.5 0 0 1 18 20.25H6.75a5.25 5.25 0 0 1-2.23-10.004a6.072 6.072 0 0 1-.02-.496" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 12.5A4.5 4.5 0 0 0 5.5 17H15a4 4 0 0 0 1.866-7.539a3.504 3.504 0 0 0-4.504-4.272A4.5 4.5 0 0 0 4.06 8.235A4.502 4.502 0 0 0 1 12.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeBracket(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.25 6.75L22.5 12l-5.25 5.25m-10.5 0L1.5 12l5.25-5.25m7.5-3l-4.5 16.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeBracketSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.78 4.97a.75.75 0 0 1 0 1.06L2.81 8l1.97 1.97a.75.75 0 1 1-1.06 1.06l-2.5-2.5a.75.75 0 0 1 0-1.06l2.5-2.5a.75.75 0 0 1 1.06 0m6.44 0a.75.75 0 0 0 0 1.06L13.19 8l-1.97 1.97a.75.75 0 1 0 1.06 1.06l2.5-2.5a.75.75 0 0 0 0-1.06l-2.5-2.5a.75.75 0 0 0-1.06 0M8.856 2.008a.75.75 0 0 1 .636.848l-1.5 10.5a.75.75 0 0 1-1.484-.212l1.5-10.5a.75.75 0 0 1 .848-.636" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeBracketSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.447 3.026a.75.75 0 0 1 .527.921l-4.5 16.5a.75.75 0 0 1-1.448-.394l4.5-16.5a.75.75 0 0 1 .921-.527M16.72 6.22a.75.75 0 0 1 1.06 0l5.25 5.25a.75.75 0 0 1 0 1.06l-5.25 5.25a.75.75 0 1 1-1.06-1.06L21.44 12l-4.72-4.72a.75.75 0 0 1 0-1.06m-9.44 0a.75.75 0 0 1 0 1.06L2.56 12l4.72 4.72a.75.75 0 0 1-1.06 1.06L.97 12.53a.75.75 0 0 1 0-1.06l5.25-5.25a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeBracketSquare(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.25 9.75L16.5 12l-2.25 2.25m-4.5 0L7.5 12l2.25-2.25M6 20.25h12A2.25 2.25 0 0 0 20.25 18V6A2.25 2.25 0 0 0 18 3.75H6A2.25 2.25 0 0 0 3.75 6v12A2.25 2.25 0 0 0 6 20.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeBracketSquareSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2zm4.78 1.97a.75.75 0 0 1 0 1.06L5.81 8l.97.97a.75.75 0 1 1-1.06 1.06l-1.5-1.5a.75.75 0 0 1 0-1.06l1.5-1.5a.75.75 0 0 1 1.06 0m2.44 1.06a.75.75 0 0 1 1.06-1.06l1.5 1.5a.75.75 0 0 1 0 1.06l-1.5 1.5a.75.75 0 1 1-1.06-1.06l.97-.97z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeBracketSquareSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm14.25 6a.75.75 0 0 1-.22.53l-2.25 2.25a.75.75 0 1 1-1.06-1.06L15.44 12l-1.72-1.72a.75.75 0 1 1 1.06-1.06l2.25 2.25c.141.14.22.331.22.53m-10.28-.53a.75.75 0 0 0 0 1.06l2.25 2.25a.75.75 0 1 0 1.06-1.06L8.56 12l1.72-1.72a.75.75 0 1 0-1.06-1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeBracketSquareTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.25 2A2.25 2.25 0 0 0 2 4.25v11.5A2.25 2.25 0 0 0 4.25 18h11.5A2.25 2.25 0 0 0 18 15.75V4.25A2.25 2.25 0 0 0 15.75 2zm4.03 6.28a.75.75 0 0 0-1.06-1.06L4.97 9.47a.75.75 0 0 0 0 1.06l2.25 2.25a.75.75 0 0 0 1.06-1.06L6.56 10zm4.5-1.06a.75.75 0 1 0-1.06 1.06L13.44 10l-1.72 1.72a.75.75 0 1 0 1.06 1.06l2.25-2.25a.75.75 0 0 0 0-1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeBracketTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.28 5.22a.75.75 0 0 1 0 1.06L2.56 10l3.72 3.72a.75.75 0 0 1-1.06 1.06L.97 10.53a.75.75 0 0 1 0-1.06l4.25-4.25a.75.75 0 0 1 1.06 0m7.44 0a.75.75 0 0 1 1.06 0l4.25 4.25a.75.75 0 0 1 0 1.06l-4.25 4.25a.75.75 0 0 1-1.06-1.06L17.44 10l-3.72-3.72a.75.75 0 0 1 0-1.06m-2.343-3.209a.75.75 0 0 1 .612.867l-2.5 14.5a.75.75 0 0 1-1.478-.255l2.5-14.5a.75.75 0 0 1 .866-.612" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cog(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.5 12a7.5 7.5 0 0 0 15 0m-15 0a7.5 7.5 0 1 1 15 0m-15 0H3m16.5 0H21m-1.5 0H12m-8.457 3.077l1.41-.513m14.095-5.13l1.41-.513M5.106 17.785l1.15-.964m11.49-9.642l1.149-.964M7.501 19.795l.75-1.3m7.5-12.99l.75-1.3m-6.063 16.658l.26-1.477m2.605-14.772l.26-1.477m0 17.726l-.26-1.477M10.698 4.614l-.26-1.477M16.5 19.794l-.75-1.299M7.5 4.205L12 12m6.894 5.785l-1.149-.964M6.256 7.178l-1.15-.964m15.352 8.864l-1.41-.513M4.954 9.435l-1.41-.514M12.002 12l-3.75 6.495"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CogEightTooth(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.343 3.94c.09-.542.56-.94 1.11-.94h1.093c.55 0 1.02.398 1.11.94l.149.894c.07.424.384.764.78.93c.398.164.855.142 1.205-.108l.737-.527a1.125 1.125 0 0 1 1.45.12l.773.774c.39.389.44 1.002.12 1.45l-.527.737c-.25.35-.272.806-.107 1.204c.165.397.505.71.93.78l.893.15c.543.09.94.559.94 1.109v1.094c0 .55-.397 1.02-.94 1.11l-.894.149c-.424.07-.764.383-.929.78c-.165.398-.143.854.107 1.204l.527.738c.32.447.269 1.06-.12 1.45l-.774.773a1.125 1.125 0 0 1-1.449.12l-.738-.527c-.35-.25-.806-.272-1.203-.107c-.398.165-.71.505-.781.929l-.149.894c-.09.542-.56.94-1.11.94h-1.094c-.55 0-1.019-.398-1.11-.94l-.148-.894c-.071-.424-.384-.764-.781-.93c-.398-.164-.854-.142-1.204.108l-.738.527c-.447.32-1.06.269-1.45-.12l-.773-.774a1.125 1.125 0 0 1-.12-1.45l.527-.737c.25-.35.272-.806.108-1.204c-.165-.397-.506-.71-.93-.78l-.894-.15c-.542-.09-.94-.56-.94-1.109v-1.094c0-.55.398-1.02.94-1.11l.894-.149c.424-.07.765-.383.93-.78c.165-.398.143-.854-.108-1.204l-.526-.738a1.125 1.125 0 0 1 .12-1.45l.773-.773a1.125 1.125 0 0 1 1.45-.12l.737.527c.35.25.807.272 1.204.107c.397-.165.71-.505.78-.929z"/><path d="M15 12a3 3 0 1 1-6 0a3 3 0 0 1 6 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CogEightToothSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.955 1.45A.5.5 0 0 1 7.452 1h1.096a.5.5 0 0 1 .497.45l.17 1.699c.484.12.94.312 1.356.562l1.321-1.081a.5.5 0 0 1 .67.033l.774.775a.5.5 0 0 1 .034.67l-1.08 1.32c.25.417.44.873.561 1.357l1.699.17a.5.5 0 0 1 .45.497v1.096a.5.5 0 0 1-.45.497l-1.699.17c-.12.484-.312.94-.562 1.356l1.082 1.322a.5.5 0 0 1-.034.67l-.774.774a.5.5 0 0 1-.67.033l-1.322-1.08a4.97 4.97 0 0 1-1.356.561l-.17 1.699a.5.5 0 0 1-.497.45H7.452a.5.5 0 0 1-.497-.45l-.17-1.699a4.973 4.973 0 0 1-1.356-.562L4.108 13.37a.5.5 0 0 1-.67-.033l-.774-.775a.5.5 0 0 1-.034-.67l1.08-1.32a4.971 4.971 0 0 1-.561-1.357l-1.699-.17A.5.5 0 0 1 1 8.548V7.452a.5.5 0 0 1 .45-.497l1.699-.17c.12-.484.312-.94.562-1.356L2.629 4.107a.5.5 0 0 1 .034-.67l.774-.774a.5.5 0 0 1 .67-.033L5.43 3.71a4.97 4.97 0 0 1 1.356-.561zM6 8c0 .538.212 1.026.558 1.385l.057.057a2 2 0 0 0 2.828-2.828l-.058-.056A2 2 0 0 0 6 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CogEightToothSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.828 2.25c-.916 0-1.699.663-1.85 1.567l-.091.549a.798.798 0 0 1-.517.608a7.45 7.45 0 0 0-.478.198a.798.798 0 0 1-.796-.064l-.453-.324a1.875 1.875 0 0 0-2.416.2l-.243.243a1.875 1.875 0 0 0-.2 2.416l.324.453a.798.798 0 0 1 .064.796a7.448 7.448 0 0 0-.198.478a.798.798 0 0 1-.608.517l-.55.092a1.875 1.875 0 0 0-1.566 1.849v.344c0 .916.663 1.699 1.567 1.85l.549.091c.281.047.508.25.608.517c.06.162.127.321.198.478a.798.798 0 0 1-.064.796l-.324.453a1.875 1.875 0 0 0 .2 2.416l.243.243c.648.648 1.67.733 2.416.2l.453-.324a.798.798 0 0 1 .796-.064c.157.071.316.137.478.198c.267.1.47.327.517.608l.092.55c.15.903.932 1.566 1.849 1.566h.344c.916 0 1.699-.663 1.85-1.567l.091-.549a.798.798 0 0 1 .517-.608a7.52 7.52 0 0 0 .478-.198a.798.798 0 0 1 .796.064l.453.324a1.875 1.875 0 0 0 2.416-.2l.243-.243c.648-.648.733-1.67.2-2.416l-.324-.453a.798.798 0 0 1-.064-.796c.071-.157.137-.316.198-.478c.1-.267.327-.47.608-.517l.55-.091a1.875 1.875 0 0 0 1.566-1.85v-.344c0-.916-.663-1.699-1.567-1.85l-.549-.091a.798.798 0 0 1-.608-.517a7.507 7.507 0 0 0-.198-.478a.798.798 0 0 1 .064-.796l.324-.453a1.875 1.875 0 0 0-.2-2.416l-.243-.243a1.875 1.875 0 0 0-2.416-.2l-.453.324a.798.798 0 0 1-.796.064a7.462 7.462 0 0 0-.478-.198a.798.798 0 0 1-.517-.608l-.091-.55a1.875 1.875 0 0 0-1.85-1.566zM12 15.75a3.75 3.75 0 1 0 0-7.5a3.75 3.75 0 0 0 0 7.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CogEightToothTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.34 1.804A1 1 0 0 1 9.32 1h1.36a1 1 0 0 1 .98.804l.295 1.473c.497.144.971.342 1.416.587l1.25-.834a1 1 0 0 1 1.262.125l.962.962a1 1 0 0 1 .125 1.262l-.834 1.25c.245.445.443.919.587 1.416l1.473.294a1 1 0 0 1 .804.98v1.361a1 1 0 0 1-.804.98l-1.473.295a6.95 6.95 0 0 1-.587 1.416l.834 1.25a1 1 0 0 1-.125 1.262l-.962.962a1 1 0 0 1-1.262.125l-1.25-.834a6.953 6.953 0 0 1-1.416.587l-.294 1.473a1 1 0 0 1-.98.804H9.32a1 1 0 0 1-.98-.804l-.295-1.473a6.957 6.957 0 0 1-1.416-.587l-1.25.834a1 1 0 0 1-1.262-.125l-.962-.962a1 1 0 0 1-.125-1.262l.834-1.25a6.957 6.957 0 0 1-.587-1.416l-1.473-.294A1 1 0 0 1 1 10.68V9.32a1 1 0 0 1 .804-.98l1.473-.295c.144-.497.342-.971.587-1.416l-.834-1.25a1 1 0 0 1 .125-1.262l.962-.962A1 1 0 0 1 5.38 3.03l1.25.834a6.957 6.957 0 0 1 1.416-.587zM13 10a3 3 0 1 1-6 0a3 3 0 0 1 6 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CogSixTooth(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.594 3.94c.09-.542.56-.94 1.11-.94h2.593c.55 0 1.02.398 1.11.94l.213 1.281c.063.374.313.686.645.87c.074.04.147.083.22.127c.325.196.72.257 1.075.124l1.217-.456a1.125 1.125 0 0 1 1.37.49l1.296 2.247a1.125 1.125 0 0 1-.26 1.431l-1.003.827c-.293.241-.438.613-.43.992a7.723 7.723 0 0 1 0 .255c-.008.378.137.75.43.991l1.004.827c.424.35.534.955.26 1.43l-1.298 2.247a1.125 1.125 0 0 1-1.369.491l-1.217-.456c-.355-.133-.75-.072-1.076.124a6.47 6.47 0 0 1-.22.128c-.331.183-.581.495-.644.869l-.213 1.281c-.09.543-.56.94-1.11.94h-2.594c-.55 0-1.019-.398-1.11-.94l-.213-1.281c-.062-.374-.312-.686-.644-.87a6.52 6.52 0 0 1-.22-.127c-.325-.196-.72-.257-1.076-.124l-1.217.456a1.125 1.125 0 0 1-1.369-.49l-1.297-2.247a1.125 1.125 0 0 1 .26-1.431l1.004-.827c.292-.24.437-.613.43-.991a6.932 6.932 0 0 1 0-.255c.007-.38-.138-.751-.43-.992l-1.004-.827a1.125 1.125 0 0 1-.26-1.43l1.297-2.247a1.125 1.125 0 0 1 1.37-.491l1.216.456c.356.133.751.072 1.076-.124c.072-.044.146-.086.22-.128c.332-.183.582-.495.644-.869z"/><path d="M15 12a3 3 0 1 1-6 0a3 3 0 0 1 6 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CogSixToothSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.455 1.45A.5.5 0 0 1 6.952 1h2.096a.5.5 0 0 1 .497.45l.186 1.858a4.996 4.996 0 0 1 1.466.848l1.703-.769a.5.5 0 0 1 .639.206l1.047 1.814a.5.5 0 0 1-.14.656l-1.517 1.09a5.026 5.026 0 0 1 0 1.694l1.516 1.09a.5.5 0 0 1 .141.656l-1.047 1.814a.5.5 0 0 1-.639.206l-1.703-.768c-.433.36-.928.649-1.466.847l-.186 1.858a.5.5 0 0 1-.497.45H6.952a.5.5 0 0 1-.497-.45l-.186-1.858a4.993 4.993 0 0 1-1.466-.848l-1.703.769a.5.5 0 0 1-.639-.206l-1.047-1.814a.5.5 0 0 1 .14-.656l1.517-1.09a5.033 5.033 0 0 1 0-1.694l-1.516-1.09a.5.5 0 0 1-.141-.656L2.46 3.593a.5.5 0 0 1 .639-.206l1.703.769c.433-.36.928-.65 1.466-.848zm-.177 7.567l-.022-.037a2 2 0 0 1 3.466-1.997l.022.037a2 2 0 0 1-3.466 1.997" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CogSixToothSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.078 2.25c-.917 0-1.699.663-1.85 1.567L9.05 4.889c-.02.12-.115.26-.297.348a7.493 7.493 0 0 0-.986.57c-.166.115-.334.126-.45.083L6.3 5.508a1.875 1.875 0 0 0-2.282.819l-.922 1.597a1.875 1.875 0 0 0 .432 2.385l.84.692c.095.078.17.229.154.43a7.598 7.598 0 0 0 0 1.139c.015.2-.059.352-.153.43l-.841.692a1.875 1.875 0 0 0-.432 2.385l.922 1.597a1.875 1.875 0 0 0 2.282.818l1.019-.382c.115-.043.283-.031.45.082c.312.214.641.405.985.57c.182.088.277.228.297.35l.178 1.071a1.876 1.876 0 0 0 1.85 1.567h1.844c.916 0 1.699-.663 1.85-1.567l.178-1.072c.02-.12.114-.26.297-.349c.344-.165.673-.356.985-.57c.167-.114.335-.125.45-.082l1.02.382a1.875 1.875 0 0 0 2.28-.819l.923-1.597a1.875 1.875 0 0 0-.432-2.385l-.84-.692c-.095-.078-.17-.229-.154-.43a7.614 7.614 0 0 0 0-1.139c-.016-.2.059-.352.153-.43l.84-.692c.708-.582.891-1.59.433-2.385l-.922-1.597a1.875 1.875 0 0 0-2.282-.818l-1.02.382c-.114.043-.282.031-.449-.083a7.49 7.49 0 0 0-.985-.57c-.183-.087-.277-.227-.297-.348l-.179-1.072a1.875 1.875 0 0 0-1.85-1.567zM12 15.75a3.75 3.75 0 1 0 0-7.5a3.75 3.75 0 0 0 0 7.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CogSixToothTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.84 1.804A1 1 0 0 1 8.82 1h2.36a1 1 0 0 1 .98.804l.331 1.652a6.993 6.993 0 0 1 1.929 1.115l1.598-.54a1 1 0 0 1 1.186.447l1.18 2.044a1 1 0 0 1-.205 1.251l-1.267 1.113a7.047 7.047 0 0 1 0 2.228l1.267 1.113a1 1 0 0 1 .206 1.25l-1.18 2.045a1 1 0 0 1-1.187.447l-1.598-.54a6.993 6.993 0 0 1-1.929 1.115l-.33 1.652a1 1 0 0 1-.98.804H8.82a1 1 0 0 1-.98-.804l-.331-1.652a6.993 6.993 0 0 1-1.929-1.115l-1.598.54a1 1 0 0 1-1.186-.447l-1.18-2.044a1 1 0 0 1 .205-1.251l1.267-1.114a7.05 7.05 0 0 1 0-2.227L1.821 7.773a1 1 0 0 1-.206-1.25l1.18-2.045a1 1 0 0 1 1.187-.447l1.598.54a6.992 6.992 0 0 1 1.93-1.115zM10 13a3 3 0 1 0 0-6a3 3 0 0 0 0 6" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CogSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 1.938a.75.75 0 0 1 1.025.274l.652 1.131c.351-.138.71-.233 1.073-.288V1.75a.75.75 0 0 1 1.5 0v1.306a5.03 5.03 0 0 1 1.072.288l.654-1.132a.75.75 0 1 1 1.298.75l-.652 1.13c.286.23.55.492.785.786l1.13-.653a.75.75 0 1 1 .75 1.3l-1.13.652c.137.351.233.71.288 1.073h1.305a.75.75 0 0 1 0 1.5h-1.306a5.032 5.032 0 0 1-.288 1.072l1.132.654a.75.75 0 0 1-.75 1.298l-1.13-.652c-.23.286-.492.55-.786.785l.652 1.13a.75.75 0 0 1-1.298.75l-.653-1.13c-.351.137-.71.233-1.073.288v1.305a.75.75 0 0 1-1.5 0v-1.306a5.032 5.032 0 0 1-1.072-.288l-.653 1.132a.75.75 0 0 1-1.3-.75l.653-1.13a4.966 4.966 0 0 1-.785-.786l-1.13.652a.75.75 0 0 1-.75-1.298l1.13-.653a4.965 4.965 0 0 1-.288-1.073H1.75a.75.75 0 0 1 0-1.5h1.306a5.03 5.03 0 0 1 .288-1.072l-1.132-.653a.75.75 0 0 1 .75-1.3l1.13.653c.23-.286.492-.55.786-.785l-.653-1.13A.75.75 0 0 1 4.5 1.937m1.14 3.476a3.501 3.501 0 0 0 0 5.172L7.135 8zM8.434 8.75L6.94 11.336a3.491 3.491 0 0 0 2.81-.305a3.49 3.49 0 0 0 1.669-2.281zm2.987-1.5H8.433L6.94 4.664a3.501 3.501 0 0 1 4.48 2.586" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CogSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M17.004 10.407c.138.435-.216.842-.672.842h-3.465a.75.75 0 0 1-.65-.375l-1.732-3c-.229-.396-.053-.907.393-1.004a5.252 5.252 0 0 1 6.126 3.537M8.12 8.464c.307-.338.838-.235 1.066.16l1.732 3a.75.75 0 0 1 0 .75l-1.732 3c-.229.397-.76.5-1.067.161A5.23 5.23 0 0 1 6.75 12a5.23 5.23 0 0 1 1.37-3.536m2.758 8.666c-.447-.098-.623-.608-.394-1.004l1.733-3.002a.75.75 0 0 1 .65-.375h3.465c.457 0 .81.407.672.842a5.252 5.252 0 0 1-6.126 3.539"/><path fill-rule="evenodd" d="M21 12.75a.75.75 0 1 0 0-1.5h-.783a8.22 8.22 0 0 0-.237-1.357l.734-.267a.75.75 0 1 0-.513-1.41l-.735.268a8.24 8.24 0 0 0-.689-1.192l.6-.503a.75.75 0 1 0-.964-1.149l-.6.504a8.3 8.3 0 0 0-1.054-.885l.391-.678a.75.75 0 1 0-1.299-.75l-.39.676a8.188 8.188 0 0 0-1.295-.47l.136-.77a.75.75 0 0 0-1.477-.26l-.136.77a8.36 8.36 0 0 0-1.377 0l-.136-.77a.75.75 0 1 0-1.477.26l.136.77c-.448.121-.88.28-1.294.47l-.39-.676a.75.75 0 0 0-1.3.75l.392.678a8.29 8.29 0 0 0-1.054.885l-.6-.504a.75.75 0 1 0-.965 1.149l.6.503a8.243 8.243 0 0 0-.689 1.192L3.8 8.216a.75.75 0 1 0-.513 1.41l.735.267a8.222 8.222 0 0 0-.238 1.356h-.783a.75.75 0 0 0 0 1.5h.783c.042.464.122.917.238 1.356l-.735.268a.75.75 0 0 0 .513 1.41l.735-.268c.197.417.428.816.69 1.191l-.6.504a.75.75 0 0 0 .963 1.15l.601-.505c.326.323.679.62 1.054.885l-.392.68a.75.75 0 0 0 1.3.75l.39-.679c.414.192.847.35 1.294.471l-.136.77a.75.75 0 0 0 1.477.261l.137-.772a8.332 8.332 0 0 0 1.376 0l.136.772a.75.75 0 1 0 1.477-.26l-.136-.771a8.19 8.19 0 0 0 1.294-.47l.391.677a.75.75 0 0 0 1.3-.75l-.393-.679a8.29 8.29 0 0 0 1.054-.885l.601.504a.75.75 0 0 0 .964-1.15l-.6-.503c.261-.375.492-.774.69-1.191l.735.267a.75.75 0 1 0 .512-1.41l-.734-.267c.115-.439.195-.892.237-1.356zm-2.657-3.06a6.744 6.744 0 0 0-1.19-2.053a6.784 6.784 0 0 0-1.82-1.51A6.705 6.705 0 0 0 12 5.25a6.8 6.8 0 0 0-1.225.11a6.7 6.7 0 0 0-2.15.793a6.784 6.784 0 0 0-2.952 3.489a.76.76 0 0 1-.036.098A6.74 6.74 0 0 0 5.251 12a6.74 6.74 0 0 0 3.366 5.842l.009.005a6.704 6.704 0 0 0 2.18.798l.022.003a6.792 6.792 0 0 0 2.368-.004a6.704 6.704 0 0 0 2.205-.811a6.785 6.785 0 0 0 1.762-1.484l.009-.01l.009-.01a6.743 6.743 0 0 0 1.18-2.065c.253-.708.39-1.47.39-2.264a6.74 6.74 0 0 0-.408-2.309" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CogTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M13.024 9.25c.47 0 .827-.433.637-.863a4 4 0 0 0-4.094-2.364c-.468.05-.665.576-.43.984l1.08 1.868a.75.75 0 0 0 .649.375zM7.84 7.758c-.236-.408-.79-.5-1.068-.12A3.982 3.982 0 0 0 6 10c0 .884.287 1.7.772 2.363c.278.38.832.287 1.068-.12l1.078-1.868a.75.75 0 0 0 0-.75zm1.298 5.235c-.235.408-.039.934.43.984a4 4 0 0 0 4.094-2.364c.19-.43-.168-.863-.638-.863h-2.158a.75.75 0 0 0-.65.375z"/><path fill-rule="evenodd" d="m14.13 4.347l.644-1.117a.75.75 0 0 0-1.299-.75l-.644 1.116a6.954 6.954 0 0 0-2.081-.556V1.75a.75.75 0 0 0-1.5 0v1.29a6.954 6.954 0 0 0-2.081.556L6.525 2.48a.75.75 0 1 0-1.3.75l.645 1.117A7.04 7.04 0 0 0 4.347 5.87L3.23 5.225a.75.75 0 1 0-.75 1.3l1.116.644A6.954 6.954 0 0 0 3.04 9.25H1.75a.75.75 0 0 0 0 1.5h1.29a6.98 6.98 0 0 0 .556 2.081l-1.116.645a.75.75 0 1 0 .75 1.298l1.117-.644a7.04 7.04 0 0 0 1.523 1.523l-.645 1.117a.75.75 0 1 0 1.3.75l.644-1.116a6.954 6.954 0 0 0 2.081.556v1.29a.75.75 0 0 0 1.5 0v-1.29a6.954 6.954 0 0 0 2.081-.556l.645 1.116a.75.75 0 0 0 1.299-.75l-.645-1.117a7.042 7.042 0 0 0 1.523-1.523l1.117.644a.75.75 0 0 0 .75-1.298l-1.116-.645a6.954 6.954 0 0 0 .556-2.081h1.29a.75.75 0 0 0 0-1.5h-1.29a6.954 6.954 0 0 0-.556-2.081l1.116-.644a.75.75 0 0 0-.75-1.3l-1.117.645a7.04 7.04 0 0 0-1.524-1.523M10 4.5a5.475 5.475 0 0 0-2.781.754A5.527 5.527 0 0 0 5.22 7.277A5.475 5.475 0 0 0 4.5 10a5.475 5.475 0 0 0 .752 2.777a5.527 5.527 0 0 0 2.028 2.004c.802.458 1.73.719 2.72.719a5.474 5.474 0 0 0 2.78-.753a5.527 5.527 0 0 0 2.001-2.027c.458-.802.719-1.73.719-2.72a5.475 5.475 0 0 0-.753-2.78a5.528 5.528 0 0 0-2.028-2.002A5.475 5.475 0 0 0 10 4.5" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommandLine(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6.75 7.5l3 2.25l-3 2.25m4.5 0h3m-9 8.25h13.5A2.25 2.25 0 0 0 21 18V6a2.25 2.25 0 0 0-2.25-2.25H5.25A2.25 2.25 0 0 0 3 6v12a2.25 2.25 0 0 0 2.25 2.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommandLineSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2zm2.22 1.97a.75.75 0 0 0 0 1.06l.97.97l-.97.97a.75.75 0 1 0 1.06 1.06l1.5-1.5a.75.75 0 0 0 0-1.06l-1.5-1.5a.75.75 0 0 0-1.06 0M8.75 8.5a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommandLineSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 6a3 3 0 0 1 3-3h13.5a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H5.25a3 3 0 0 1-3-3zm3.97.97a.75.75 0 0 1 1.06 0l2.25 2.25a.75.75 0 0 1 0 1.06l-2.25 2.25a.75.75 0 0 1-1.06-1.06l1.72-1.72l-1.72-1.72a.75.75 0 0 1 0-1.06m4.28 4.28a.75.75 0 0 0 0 1.5h3a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommandLineTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.25 3A2.25 2.25 0 0 0 1 5.25v9.5A2.25 2.25 0 0 0 3.25 17h13.5A2.25 2.25 0 0 0 19 14.75v-9.5A2.25 2.25 0 0 0 16.75 3zm.943 8.752a.75.75 0 0 1 .055-1.06L6.128 9l-1.88-1.693a.75.75 0 1 1 1.004-1.114l2.5 2.25a.75.75 0 0 1 0 1.114l-2.5 2.25a.75.75 0 0 1-1.06-.055M9.75 10.25a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ComputerDesktop(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 17.25v1.007a3 3 0 0 1-.879 2.122L7.5 21h9l-.621-.621A3 3 0 0 1 15 18.257V17.25m6-12V15a2.25 2.25 0 0 1-2.25 2.25H5.25A2.25 2.25 0 0 1 3 15V5.25m18 0A2.25 2.25 0 0 0 18.75 3H5.25A2.25 2.25 0 0 0 3 5.25m18 0V12a2.25 2.25 0 0 1-2.25 2.25H5.25A2.25 2.25 0 0 1 3 12V5.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ComputerDesktopSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 4.25A2.25 2.25 0 0 1 4.25 2h7.5A2.25 2.25 0 0 1 14 4.25v5.5A2.25 2.25 0 0 1 11.75 12h-1.312c.1.128.21.248.328.36a.75.75 0 0 1 .234.545v.345a.75.75 0 0 1-.75.75h-4.5a.75.75 0 0 1-.75-.75v-.345a.75.75 0 0 1 .234-.545c.118-.111.228-.232.328-.36H4.25A2.25 2.25 0 0 1 2 9.75zm2.25-.75a.75.75 0 0 0-.75.75v4.5c0 .414.336.75.75.75h7.5a.75.75 0 0 0 .75-.75v-4.5a.75.75 0 0 0-.75-.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ComputerDesktopSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 5.25a3 3 0 0 1 3-3h13.5a3 3 0 0 1 3 3V15a3 3 0 0 1-3 3h-3v.257c0 .597.237 1.17.659 1.591l.621.622a.75.75 0 0 1-.53 1.28h-9a.75.75 0 0 1-.53-1.28l.621-.622a2.25 2.25 0 0 0 .659-1.59V18h-3a3 3 0 0 1-3-3zm1.5 0v7.5a1.5 1.5 0 0 0 1.5 1.5h13.5a1.5 1.5 0 0 0 1.5-1.5v-7.5a1.5 1.5 0 0 0-1.5-1.5H5.25a1.5 1.5 0 0 0-1.5 1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ComputerDesktopTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 4.25A2.25 2.25 0 0 1 4.25 2h11.5A2.25 2.25 0 0 1 18 4.25v8.5A2.25 2.25 0 0 1 15.75 15h-3.105a3.501 3.501 0 0 0 1.1 1.677A.75.75 0 0 1 13.26 18H6.74a.75.75 0 0 1-.484-1.323A3.501 3.501 0 0 0 7.355 15H4.25A2.25 2.25 0 0 1 2 12.75zm1.5 0a.75.75 0 0 1 .75-.75h11.5a.75.75 0 0 1 .75.75v7.5a.75.75 0 0 1-.75.75H4.25a.75.75 0 0 1-.75-.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CpuChip(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.25 3v1.5M4.5 8.25H3m18 0h-1.5M4.5 12H3m18 0h-1.5m-15 3.75H3m18 0h-1.5M8.25 19.5V21M12 3v1.5m0 15V21m3.75-18v1.5m0 15V21m-9-1.5h10.5a2.25 2.25 0 0 0 2.25-2.25V6.75a2.25 2.25 0 0 0-2.25-2.25H6.75A2.25 2.25 0 0 0 4.5 6.75v10.5a2.25 2.25 0 0 0 2.25 2.25m.75-12h9v9h-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CpuChipSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 6v4h4V6z"/><path fill="currentColor" fill-rule="evenodd" d="M5.75 1a.75.75 0 0 0-.75.75V3a2 2 0 0 0-2 2H1.75a.75.75 0 0 0 0 1.5H3v.75H1.75a.75.75 0 0 0 0 1.5H3v.75H1.75a.75.75 0 0 0 0 1.5H3a2 2 0 0 0 2 2v1.25a.75.75 0 0 0 1.5 0V13h.75v1.25a.75.75 0 0 0 1.5 0V13h.75v1.25a.75.75 0 0 0 1.5 0V13a2 2 0 0 0 2-2h1.25a.75.75 0 0 0 0-1.5H13v-.75h1.25a.75.75 0 0 0 0-1.5H13V6.5h1.25a.75.75 0 0 0 0-1.5H13a2 2 0 0 0-2-2V1.75a.75.75 0 0 0-1.5 0V3h-.75V1.75a.75.75 0 0 0-1.5 0V3H6.5V1.75A.75.75 0 0 0 5.75 1M11 4.5a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-.5.5H5a.5.5 0 0 1-.5-.5V5a.5.5 0 0 1 .5-.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CpuChipSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M16.5 7.5h-9v9h9z"/><path fill-rule="evenodd" d="M8.25 2.25A.75.75 0 0 1 9 3v.75h2.25V3a.75.75 0 0 1 1.5 0v.75H15V3a.75.75 0 0 1 1.5 0v.75h.75a3 3 0 0 1 3 3v.75H21A.75.75 0 0 1 21 9h-.75v2.25H21a.75.75 0 0 1 0 1.5h-.75V15H21a.75.75 0 0 1 0 1.5h-.75v.75a3 3 0 0 1-3 3h-.75V21a.75.75 0 0 1-1.5 0v-.75h-2.25V21a.75.75 0 0 1-1.5 0v-.75H9V21a.75.75 0 0 1-1.5 0v-.75h-.75a3 3 0 0 1-3-3v-.75H3A.75.75 0 0 1 3 15h.75v-2.25H3a.75.75 0 0 1 0-1.5h.75V9H3a.75.75 0 0 1 0-1.5h.75v-.75a3 3 0 0 1 3-3h.75V3a.75.75 0 0 1 .75-.75M6 6.75A.75.75 0 0 1 6.75 6h10.5a.75.75 0 0 1 .75.75v10.5a.75.75 0 0 1-.75.75H6.75a.75.75 0 0 1-.75-.75z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CpuChipTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M14 6H6v8h8z"/><path fill-rule="evenodd" d="M9.25 3V1.75a.75.75 0 0 1 1.5 0V3h1.5V1.75a.75.75 0 0 1 1.5 0V3h.5A2.75 2.75 0 0 1 17 5.75v.5h1.25a.75.75 0 0 1 0 1.5H17v1.5h1.25a.75.75 0 0 1 0 1.5H17v1.5h1.25a.75.75 0 0 1 0 1.5H17v.5A2.75 2.75 0 0 1 14.25 17h-.5v1.25a.75.75 0 0 1-1.5 0V17h-1.5v1.25a.75.75 0 0 1-1.5 0V17h-1.5v1.25a.75.75 0 0 1-1.5 0V17h-.5A2.75 2.75 0 0 1 3 14.25v-.5H1.75a.75.75 0 0 1 0-1.5H3v-1.5H1.75a.75.75 0 0 1 0-1.5H3v-1.5H1.75a.75.75 0 0 1 0-1.5H3v-.5A2.75 2.75 0 0 1 5.75 3h.5V1.75a.75.75 0 0 1 1.5 0V3zM4.5 5.75c0-.69.56-1.25 1.25-1.25h8.5c.69 0 1.25.56 1.25 1.25v8.5c0 .69-.56 1.25-1.25 1.25h-8.5c-.69 0-1.25-.56-1.25-1.25z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCard(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.25 8.25h19.5M2.25 9h19.5m-16.5 5.25h6m-6 2.25h3m-3.75 3h15a2.25 2.25 0 0 0 2.25-2.25V6.75A2.25 2.25 0 0 0 19.5 4.5h-15a2.25 2.25 0 0 0-2.25 2.25v10.5A2.25 2.25 0 0 0 4.5 19.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M2.5 3A1.5 1.5 0 0 0 1 4.5V5h14v-.5A1.5 1.5 0 0 0 13.5 3z"/><path fill-rule="evenodd" d="M15 7H1v4.5A1.5 1.5 0 0 0 2.5 13h11a1.5 1.5 0 0 0 1.5-1.5zM3 10.25a.75.75 0 0 1 .75-.75h.5a.75.75 0 0 1 0 1.5h-.5a.75.75 0 0 1-.75-.75m3.75-.75a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M4.5 3.75a3 3 0 0 0-3 3v.75h21v-.75a3 3 0 0 0-3-3z"/><path fill-rule="evenodd" d="M22.5 9.75h-21v7.5a3 3 0 0 0 3 3h15a3 3 0 0 0 3-3zm-18 3.75a.75.75 0 0 1 .75-.75h6a.75.75 0 0 1 0 1.5h-6a.75.75 0 0 1-.75-.75m.75 2.25a.75.75 0 0 0 0 1.5h3a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 4A1.5 1.5 0 0 0 1 5.5V6h18v-.5A1.5 1.5 0 0 0 17.5 4zM19 8.5H1v6A1.5 1.5 0 0 0 2.5 16h15a1.5 1.5 0 0 0 1.5-1.5zM3 13.25a.75.75 0 0 1 .75-.75h1.5a.75.75 0 0 1 0 1.5h-1.5a.75.75 0 0 1-.75-.75m4.75-.75a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cube(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m21 7.5l-9-5.25L3 7.5m18 0l-9 5.25m9-5.25v9l-9 5.25M3 7.5l9 5.25M3 7.5v9l9 5.25m0-9v9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CubeSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.372 1.349a.75.75 0 0 0-.744 0l-4.81 2.748L8 7.131l5.182-3.034zM14 5.357L8.75 8.43v6.005l4.872-2.784A.75.75 0 0 0 14 11zm-6.75 9.078V8.43L2 5.357V11c0 .27.144.518.378.651z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CubeSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.378 1.602a.75.75 0 0 0-.756 0L3 6.632l9 5.25l9-5.25zM21.75 7.93l-9 5.25v9l8.628-5.032a.75.75 0 0 0 .372-.648zm-10.5 14.25v-9l-9-5.25v8.57a.75.75 0 0 0 .372.648z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CubeTransparent(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m21 7.5l-2.25-1.312M21 7.5v2.25m0-2.25l-2.25 1.313M3 7.5l2.25-1.312M3 7.5l2.25 1.313M3 7.5v2.25m9 3l2.25-1.312M12 12.75l-2.25-1.312M12 12.75V15m0 6.75l2.25-1.312M12 21.75V19.5m0 2.25l-2.25-1.312m0-16.875L12 2.25l2.25 1.313M21 14.25v2.25l-2.25 1.313m-13.5 0L3 16.5v-2.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CubeTransparentSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.628 1.349a.75.75 0 0 1 .744 0l1.247.712a.75.75 0 1 1-.744 1.303L8 2.864l-.875.5a.75.75 0 0 1-.744-1.303zM4.65 3.914a.75.75 0 0 1-.279 1.023L4.262 5l.11.063a.75.75 0 0 1-.744 1.302l-.13-.073A.75.75 0 0 1 2 6.25V5a.75.75 0 0 1 .378-.651l1.25-.714a.75.75 0 0 1 1.023.279m6.698 0a.75.75 0 0 1 1.023-.28l1.25.715A.75.75 0 0 1 14 5v1.25a.75.75 0 0 1-1.499.042l-.129.073a.75.75 0 0 1-.744-1.302l.11-.063l-.11-.063a.75.75 0 0 1-.28-1.023M6.102 6.915a.75.75 0 0 1 1.023-.279l.875.5l.875-.5a.75.75 0 0 1 .744 1.303l-.869.496v.815a.75.75 0 0 1-1.5 0v-.815l-.869-.496a.75.75 0 0 1-.28-1.024M2.75 9a.75.75 0 0 1 .75.75v.815l.872.498a.75.75 0 0 1-.744 1.303l-1.25-.715A.75.75 0 0 1 2 11V9.75A.75.75 0 0 1 2.75 9m10.5 0a.75.75 0 0 1 .75.75V11a.75.75 0 0 1-.378.651l-1.25.715a.75.75 0 0 1-.744-1.303l.872-.498V9.75a.75.75 0 0 1 .75-.75m-4.501 3.708l.126-.072a.75.75 0 0 1 .744 1.303l-1.247.712a.75.75 0 0 1-.744 0L6.38 13.94a.75.75 0 0 1 .744-1.303l.126.072a.75.75 0 0 1 1.498 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CubeTransparentSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.622 1.602a.75.75 0 0 1 .756 0l2.25 1.313a.75.75 0 0 1-.756 1.295L12 3.118L10.128 4.21a.75.75 0 1 1-.756-1.295zM5.898 5.81a.75.75 0 0 1-.27 1.025l-1.14.665l1.14.665a.75.75 0 1 1-.756 1.295L3.75 8.806v.944a.75.75 0 0 1-1.5 0V7.5a.75.75 0 0 1 .372-.648l2.25-1.312a.75.75 0 0 1 1.026.27m12.204 0a.75.75 0 0 1 1.026-.27l2.25 1.312a.75.75 0 0 1 .372.648v2.25a.75.75 0 0 1-1.5 0v-.944l-1.122.654a.75.75 0 1 1-.756-1.295l1.14-.665l-1.14-.665a.75.75 0 0 1-.27-1.025m-9 5.25a.75.75 0 0 1 1.026-.27L12 11.882l1.872-1.092a.75.75 0 1 1 .756 1.295l-1.878 1.096V15a.75.75 0 0 1-1.5 0v-1.82l-1.878-1.095a.75.75 0 0 1-.27-1.025M3 13.5a.75.75 0 0 1 .75.75v1.82l1.878 1.095a.75.75 0 1 1-.756 1.295l-2.25-1.312a.75.75 0 0 1-.372-.648v-2.25A.75.75 0 0 1 3 13.5m18 0a.75.75 0 0 1 .75.75v2.25a.75.75 0 0 1-.372.648l-2.25 1.312a.75.75 0 1 1-.756-1.295l1.878-1.096V14.25a.75.75 0 0 1 .75-.75m-9 5.25a.75.75 0 0 1 .75.75v.944l1.122-.654a.75.75 0 1 1 .756 1.295l-2.25 1.313a.75.75 0 0 1-.756 0l-2.25-1.313a.75.75 0 1 1 .756-1.295l1.122.654V19.5a.75.75 0 0 1 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CubeTransparentTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.638 1.093a.75.75 0 0 1 .724 0l2 1.104a.75.75 0 1 1-.724 1.313L10 2.607l-1.638.903a.75.75 0 1 1-.724-1.313zM5.403 4.287a.75.75 0 0 1-.295 1.019l-.805.444l.805.444a.75.75 0 0 1-.724 1.314L3.5 7.02v.73a.75.75 0 0 1-1.5 0v-2a.75.75 0 0 1 .388-.657l1.996-1.1a.75.75 0 0 1 1.019.294m9.194 0a.75.75 0 0 1 1.02-.295l1.995 1.101A.75.75 0 0 1 18 5.75v2a.75.75 0 0 1-1.5 0v-.73l-.884.488a.75.75 0 1 1-.724-1.314l.806-.444l-.806-.444a.75.75 0 0 1-.295-1.02M7.343 8.284a.75.75 0 0 1 1.02-.294L10 8.893l1.638-.903a.75.75 0 1 1 .724 1.313l-1.612.89v1.557a.75.75 0 0 1-1.5 0v-1.557l-1.612-.89a.75.75 0 0 1-.295-1.019M2.75 11.5a.75.75 0 0 1 .75.75v1.557l1.608.887a.75.75 0 0 1-.724 1.314l-1.996-1.101A.75.75 0 0 1 2 14.25v-2a.75.75 0 0 1 .75-.75m14.5 0a.75.75 0 0 1 .75.75v2a.75.75 0 0 1-.388.657l-1.996 1.1a.75.75 0 1 1-.724-1.313l1.608-.887V12.25a.75.75 0 0 1 .75-.75m-7.25 4a.75.75 0 0 1 .75.75v.73l.888-.49a.75.75 0 0 1 .724 1.313l-2 1.104a.75.75 0 0 1-.724 0l-2-1.104a.75.75 0 1 1 .724-1.313l.888.49v-.73a.75.75 0 0 1 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CubeTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.362 1.093a.75.75 0 0 0-.724 0L2.523 5.018L10 9.143l7.477-4.125zM18 6.443l-7.25 4v8.25l6.862-3.786A.75.75 0 0 0 18 14.25zm-8.75 12.25v-8.25l-7.25-4v7.807a.75.75 0 0 0 .388.657z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrencyBangladeshi(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8.25 7.5l.415-.207a.75.75 0 0 1 1.085.67V10.5m0 0h6m-6 0h-1.5m1.5 0v5.438c0 .354.161.697.473.865a3.751 3.751 0 0 0 5.452-2.553c.083-.409-.263-.75-.68-.75h-.745M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrencyBangladeshiSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M5.25 4.707a.75.75 0 0 1-.78-1.237c.841-.842 2.28-.246 2.28.944V6h5.5a.75.75 0 0 1 0 1.5h-5.5v3.098c0 .549.295.836.545.87a3.241 3.241 0 0 0 2.799-.966H9.75a.75.75 0 0 1 0-1.5h1.708a.75.75 0 0 1 .695 1.032a4.751 4.751 0 0 1-5.066 2.92c-1.266-.177-1.837-1.376-1.837-2.356V7.5h-1.5a.75.75 0 0 1 0-1.5h1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrencyBangladeshiSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 21.75c5.385 0 9.75-4.365 9.75-9.75S17.385 2.25 12 2.25S2.25 6.615 2.25 12s4.365 9.75 9.75 9.75M10.5 7.963a1.5 1.5 0 0 0-2.17-1.341l-.415.207a.75.75 0 0 0 .67 1.342L9 7.963V9.75h-.75a.75.75 0 1 0 0 1.5H9v4.688c0 .563.26 1.198.867 1.525A4.501 4.501 0 0 0 16.41 14.4c.199-.977-.636-1.649-1.415-1.649h-.745a.75.75 0 1 0 0 1.5h.656a3.002 3.002 0 0 1-4.327 1.893a.113.113 0 0 1-.045-.051a.336.336 0 0 1-.034-.154V11.25h5.25a.75.75 0 0 0 0-1.5H10.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrencyBangladeshiTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 2a8 8 0 1 0 0 16a8 8 0 0 0 0-16M5.94 5.5c.944-.945 2.56-.276 2.56 1.06V8h5.75a.75.75 0 0 1 0 1.5H8.5v4.275c0 .296.144.455.26.499a3.5 3.5 0 0 0 4.402-1.77h-.412a.75.75 0 0 1 0-1.5h.537c.462 0 .887.21 1.156.556c.278.355.383.852.184 1.337a5.001 5.001 0 0 1-6.4 2.78C7.376 15.353 7 14.512 7 13.774V9.5H5.75a.75.75 0 0 1 0-1.5H7V6.56l-.22.22a.75.75 0 1 1-1.06-1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrencyDollar(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 6v12m-3-2.818l.879.659c1.171.879 3.07.879 4.242 0c1.172-.879 1.172-2.303 0-3.182C13.536 12.219 12.768 12 12 12c-.725 0-1.45-.22-2.003-.659c-1.106-.879-1.106-2.303 0-3.182s2.9-.879 4.006 0l.415.33M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrencyDollarSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.375 5.5h.875v1.75h-.875a.875.875 0 1 1 0-1.75m2.375 5V8.75h.875a.875.875 0 0 1 0 1.75z"/><path fill="currentColor" fill-rule="evenodd" d="M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M7.25 3.75a.75.75 0 0 1 1.5 0V4h2.5a.75.75 0 0 1 0 1.5h-2.5v1.75h.875a2.375 2.375 0 1 1 0 4.75H8.75v.25a.75.75 0 0 1-1.5 0V12h-2.5a.75.75 0 0 1 0-1.5h2.5V8.75h-.875a2.375 2.375 0 1 1 0-4.75h.875z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrencyDollarSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10.464 8.746c.227-.18.497-.311.786-.394v2.795a2.252 2.252 0 0 1-.786-.393c-.394-.313-.546-.681-.546-1.004c0-.323.152-.691.546-1.004m2.286 6.916v-2.824c.347.085.664.228.921.421c.427.32.579.686.579.991c0 .305-.152.671-.579.991a2.534 2.534 0 0 1-.921.42"/><path fill-rule="evenodd" d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75s9.75-4.365 9.75-9.75S17.385 2.25 12 2.25M12.75 6a.75.75 0 0 0-1.5 0v.816a3.836 3.836 0 0 0-1.72.756c-.712.566-1.112 1.35-1.112 2.178c0 .829.4 1.612 1.113 2.178c.502.4 1.102.647 1.719.756v2.978a2.536 2.536 0 0 1-.921-.421l-.879-.66a.75.75 0 0 0-.9 1.2l.879.66c.533.4 1.169.645 1.821.75V18a.75.75 0 0 0 1.5 0v-.81a4.124 4.124 0 0 0 1.821-.749c.745-.559 1.179-1.344 1.179-2.191c0-.847-.434-1.632-1.179-2.191a4.122 4.122 0 0 0-1.821-.75V8.354c.29.082.559.213.786.393l.415.33a.75.75 0 0 0 .933-1.175l-.415-.33a3.836 3.836 0 0 0-1.719-.755z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrencyDollarTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10.75 10.818v2.614A3.13 3.13 0 0 0 11.888 13c.482-.315.612-.648.612-.875c0-.227-.13-.56-.612-.875a3.13 3.13 0 0 0-1.138-.432M8.33 8.62c.053.055.115.11.184.164c.208.16.46.284.736.363V6.603a2.45 2.45 0 0 0-.35.13a2.01 2.01 0 0 0-.386.233c-.377.292-.514.627-.514.909c0 .184.058.39.202.592c.037.051.08.102.128.152"/><path fill-rule="evenodd" d="M18 10a8 8 0 1 1-16 0a8 8 0 0 1 16 0m-8-6a.75.75 0 0 1 .75.75v.316a3.78 3.78 0 0 1 1.653.713c.426.33.744.74.925 1.2a.75.75 0 0 1-1.395.55a1.35 1.35 0 0 0-.447-.563a2.187 2.187 0 0 0-.736-.363V9.3c.698.093 1.383.32 1.959.696c.787.514 1.29 1.27 1.29 2.13c0 .86-.504 1.616-1.29 2.13c-.576.377-1.261.603-1.96.696v.299a.75.75 0 1 1-1.5 0v-.3c-.697-.092-1.382-.318-1.958-.695c-.482-.315-.857-.717-1.078-1.188a.75.75 0 1 1 1.359-.636c.08.173.245.376.54.569c.313.205.706.353 1.138.432v-2.748a3.782 3.782 0 0 1-1.653-.713C6.9 9.433 6.5 8.681 6.5 7.875c0-.805.4-1.558 1.097-2.096a3.78 3.78 0 0 1 1.653-.713V4.75A.75.75 0 0 1 10 4" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrencyEuro(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.25 7.756a4.5 4.5 0 1 0 0 8.488M7.5 10.5h5.25m-5.25 3h5.25M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrencyEuroSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M6.875 6c.09-.22.195-.42.31-.598c.413-.638.895-.902 1.315-.902c.264 0 .54.1.814.325a.75.75 0 1 0 .953-1.158C9.772 3.259 9.169 3 8.5 3c-1.099 0-1.992.687-2.574 1.587A5.518 5.518 0 0 0 5.285 6H4.75a.75.75 0 0 0 0 1.5h.267a7.372 7.372 0 0 0 0 1H4.75a.75.75 0 0 0 0 1.5h.535c.156.52.372.998.64 1.413C6.509 12.313 7.402 13 8.5 13c.669 0 1.272-.26 1.767-.667a.75.75 0 0 0-.953-1.158c-.275.226-.55.325-.814.325c-.42 0-.902-.264-1.315-.902a3.722 3.722 0 0 1-.31-.598H8.25a.75.75 0 0 0 0-1.5H6.521a5.854 5.854 0 0 1 0-1H8.25a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrencyEuroSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75s9.75-4.365 9.75-9.75S17.385 2.25 12 2.25m-1.902 7.098a3.75 3.75 0 0 1 3.903-.884a.75.75 0 1 0 .498-1.415A5.25 5.25 0 0 0 8.005 9.75H7.5a.75.75 0 0 0 0 1.5h.054a5.281 5.281 0 0 0 0 1.5H7.5a.75.75 0 0 0 0 1.5h.505a5.25 5.25 0 0 0 6.494 2.701a.75.75 0 1 0-.498-1.415a3.75 3.75 0 0 1-4.252-1.286h3.001a.75.75 0 0 0 0-1.5H9.075a3.77 3.77 0 0 1 0-1.5h3.675a.75.75 0 0 0 0-1.5h-3c.105-.14.221-.274.348-.402" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrencyEuroTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 18a8 8 0 1 0 0-16a8 8 0 0 0 0 16M8.798 7.45c.512-.67 1.135-.95 1.702-.95s1.19.28 1.702.95a.75.75 0 0 0 1.192-.91C12.637 5.55 11.596 5 10.5 5s-2.137.55-2.894 1.54A5.205 5.205 0 0 0 6.83 8H5.75a.75.75 0 0 0 0 1.5h.77a6.333 6.333 0 0 0 0 1h-.77a.75.75 0 0 0 0 1.5h1.08a5.21 5.21 0 0 0 .776 1.46C8.363 14.45 9.404 15 10.5 15s2.137-.55 2.894-1.54a.75.75 0 0 0-1.192-.91c-.512.67-1.135.95-1.702.95s-1.19-.28-1.702-.95a3.505 3.505 0 0 1-.343-.55h1.795a.75.75 0 0 0 0-1.5H8.026a4.835 4.835 0 0 1 0-1h2.224a.75.75 0 0 0 0-1.5H8.455c.098-.195.212-.38.343-.55" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrencyPound(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.121 7.629A3 3 0 0 0 9.017 9.43a2.56 2.56 0 0 0 .028.636l.506 3.541a4.5 4.5 0 0 1-.43 2.65L9 16.5l1.539-.513a2.25 2.25 0 0 1 1.422 0l.655.218a2.25 2.25 0 0 0 1.718-.122L15 15.75M8.25 12H12m9 0a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrencyPoundSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M7.94 4.94c-.294.293-.44.675-.44 1.06v1.25h1.25a.75.75 0 1 1 0 1.5H7.5v1c0 .263-.045.516-.128.75h3.878a.75.75 0 0 1 0 1.5h-6.5a.75.75 0 0 1 0-1.5h.5A.75.75 0 0 0 6 9.75v-1H4.75a.75.75 0 0 1 0-1.5H6V6a3 3 0 0 1 5.121-2.121a.75.75 0 1 1-1.06 1.06a1.5 1.5 0 0 0-2.121 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrencyPoundSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75s9.75-4.365 9.75-9.75S17.385 2.25 12 2.25M9.763 9.51a2.25 2.25 0 0 1 3.828-1.351a.75.75 0 0 0 1.06-1.06a3.75 3.75 0 0 0-6.38 2.252c-.033.307 0 .595.032.822l.154 1.077H8.25a.75.75 0 0 0 0 1.5h.421l.138.964a3.75 3.75 0 0 1-.358 2.208l-.122.242a.75.75 0 0 0 .908 1.047l1.539-.512a1.5 1.5 0 0 1 .948 0l.655.218a3 3 0 0 0 2.29-.163l.666-.333a.75.75 0 1 0-.67-1.342l-.667.333a1.5 1.5 0 0 1-1.145.082l-.654-.218a3 3 0 0 0-1.898 0l-.06.02a5.25 5.25 0 0 0 .053-1.794l-.108-.752H12a.75.75 0 0 0 0-1.5H9.972l-.184-1.29a1.863 1.863 0 0 1-.025-.45" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrencyPoundTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 18a8 8 0 1 0 0-16a8 8 0 0 0 0 16M8.732 6.232a2.5 2.5 0 0 1 3.536 0a.75.75 0 1 0 1.06-1.06A4 4 0 0 0 6.5 8v.165c0 .364.034.728.1 1.085h-.35a.75.75 0 0 0 0 1.5h.737a5.25 5.25 0 0 1-.367 3.072l-.055.123a.75.75 0 0 0 .848 1.037l1.272-.283a3.493 3.493 0 0 1 1.604.021a4.992 4.992 0 0 0 2.422 0l.97-.242a.75.75 0 0 0-.363-1.456l-.971.243a3.491 3.491 0 0 1-1.694 0a4.992 4.992 0 0 0-2.258-.038c.19-.811.227-1.651.111-2.477H9.75a.75.75 0 0 0 0-1.5H8.136A4.397 4.397 0 0 1 8 8.165V8c0-.641.244-1.28.732-1.768" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrencyRupee(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 8.25H9m6 3H9m3 6l-3-3h1.5a3 3 0 1 0 0-6M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrencyRupeeSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M4.75 4a.75.75 0 0 0 0 1.5H6.5c.698 0 1.3.409 1.582 1H4.75a.75.75 0 0 0 0 1.5h3.332C7.8 8.591 7.198 9 6.5 9H4.75a.75.75 0 0 0-.53 1.28l2.5 2.5a.75.75 0 0 0 1.06-1.06L6.56 10.5A3.251 3.251 0 0 0 9.663 8h1.587a.75.75 0 0 0 0-1.5H9.663a3.232 3.232 0 0 0-.424-1h2.011a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrencyRupeeSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75s9.75-4.365 9.75-9.75S17.385 2.25 12 2.25M9 7.5A.75.75 0 0 0 9 9h1.5c.98 0 1.813.626 2.122 1.5H9A.75.75 0 0 0 9 12h3.622a2.251 2.251 0 0 1-2.122 1.5H9a.75.75 0 0 0-.53 1.28l3 3a.75.75 0 1 0 1.06-1.06l-1.73-1.732A3.752 3.752 0 0 0 14.175 12H15a.75.75 0 0 0 0-1.5h-.825A3.733 3.733 0 0 0 13.5 9H15a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrencyRupeeTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 18a8 8 0 1 0 0-16a8 8 0 0 0 0 16M6 5.75A.75.75 0 0 1 6.75 5h6.5a.75.75 0 0 1 0 1.5h-2.127c.4.5.683 1.096.807 1.75h1.32a.75.75 0 0 1 0 1.5h-1.32a4.003 4.003 0 0 1-3.404 3.216l1.754 1.754a.75.75 0 0 1-1.06 1.06l-3-3a.75.75 0 0 1 .53-1.28H8c1.12 0 2.067-.736 2.386-1.75H6.75a.75.75 0 0 1 0-1.5h3.636A2.501 2.501 0 0 0 8 6.5H6.75A.75.75 0 0 1 6 5.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrencyYen(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 7.5l3 4.5m0 0l3-4.5M12 12v5.25M15 12H9m6 3H9m12-3a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrencyYenSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M5.6 3.55a.75.75 0 1 0-1.2.9L7.063 8H4.75a.75.75 0 0 0 0 1.5h2.5v1h-2.5a.75.75 0 0 0 0 1.5h2.5v.5a.75.75 0 0 0 1.5 0V12h2.5a.75.75 0 0 0 0-1.5h-2.5v-1h2.5a.75.75 0 0 0 0-1.5H8.938L11.6 4.45a.75.75 0 1 0-1.2-.9L8 6.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrencyYenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75s9.75-4.365 9.75-9.75S17.385 2.25 12 2.25M9.624 7.084a.75.75 0 0 0-1.248.832l2.223 3.334H9a.75.75 0 0 0 0 1.5h2.25v1.5H9a.75.75 0 0 0 0 1.5h2.25v1.5a.75.75 0 0 0 1.5 0v-1.5H15a.75.75 0 0 0 0-1.5h-2.25v-1.5H15a.75.75 0 0 0 0-1.5h-1.599l2.223-3.334a.75.75 0 1 0-1.248-.832L12 10.648z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrencyYenTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 18a8 8 0 1 0 0-16a8 8 0 0 0 0 16M7.346 5.294a.75.75 0 0 0-1.192.912L9.056 10H6.75a.75.75 0 0 0 0 1.5h2.5v1h-2.5a.75.75 0 0 0 0 1.5h2.5v1.25a.75.75 0 0 0 1.5 0V14h2.5a.75.75 0 1 0 0-1.5h-2.5v-1h2.5a.75.75 0 1 0 0-1.5h-2.306l2.902-3.794a.75.75 0 1 0-1.192-.912L10 8.765z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CursorArrowRays(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.042 21.672L13.684 16.6m0 0l-2.51 2.225l.569-9.47l5.227 7.917zM12 2.25V4.5m5.834.166l-1.591 1.591M20.25 10.5H18M7.757 14.743l-1.59 1.59M6 10.5H3.75m4.007-4.243l-1.59-1.59"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CursorArrowRaysSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.25 1.75a.75.75 0 0 1 1.5 0v1.5a.75.75 0 0 1-1.5 0zm4.286 1.154a.75.75 0 1 1 1.06 1.06l-1.06 1.061a.75.75 0 0 1-1.061-1.06zM14.5 7.5a.75.75 0 0 0-.75-.75h-1.5a.75.75 0 0 0 0 1.5h1.5a.75.75 0 0 0 .75-.75M4.464 9.975a.75.75 0 0 1 1.061 1.06l-1.06 1.061a.75.75 0 1 1-1.061-1.06zM4.5 7.5a.75.75 0 0 0-.75-.75h-1.5a.75.75 0 0 0 0 1.5h1.5a.75.75 0 0 0 .75-.75m1.025-3.536a.75.75 0 0 1-1.06 1.061l-1.061-1.06a.75.75 0 0 1 1.06-1.061zm3.254 3.474a.75.75 0 0 0-1.368.366l-.396 5.283a.75.75 0 0 0 1.212.646l.602-.474l.288 1.074a.75.75 0 1 0 1.449-.388l-.288-1.075l.759.11a.75.75 0 0 0 .726-1.165z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CursorArrowRaysSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.5a.75.75 0 0 1 .75.75V4.5a.75.75 0 0 1-1.5 0V2.25A.75.75 0 0 1 12 1.5M5.636 4.136a.75.75 0 0 1 1.06 0l1.592 1.591a.75.75 0 0 1-1.061 1.06l-1.591-1.59a.75.75 0 0 1 0-1.061m12.728 0a.75.75 0 0 1 0 1.06l-1.591 1.592a.75.75 0 0 1-1.06-1.061l1.59-1.591a.75.75 0 0 1 1.061 0m-6.816 4.496a.75.75 0 0 1 .82.311l5.228 7.917a.75.75 0 0 1-.777 1.148l-2.097-.43l1.045 3.9a.75.75 0 0 1-1.45.388l-1.044-3.899l-1.601 1.42a.75.75 0 0 1-1.247-.606l.569-9.47a.75.75 0 0 1 .554-.68M3 10.5a.75.75 0 0 1 .75-.75H6a.75.75 0 0 1 0 1.5H3.75A.75.75 0 0 1 3 10.5m14.25 0a.75.75 0 0 1 .75-.75h2.25a.75.75 0 0 1 0 1.5H18a.75.75 0 0 1-.75-.75m-8.962 3.712a.75.75 0 0 1 0 1.061l-1.591 1.591a.75.75 0 1 1-1.061-1.06l1.591-1.592a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CursorArrowRaysTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 1a.75.75 0 0 1 .75.75v1.5a.75.75 0 0 1-1.5 0v-1.5A.75.75 0 0 1 10 1M5.05 3.05a.75.75 0 0 1 1.06 0l1.062 1.06A.75.75 0 1 1 6.11 5.173L5.05 4.11a.75.75 0 0 1 0-1.06m9.9 0a.75.75 0 0 1 0 1.06l-1.06 1.062a.75.75 0 0 1-1.062-1.061l1.061-1.06a.75.75 0 0 1 1.06 0M3 8a.75.75 0 0 1 .75-.75h1.5a.75.75 0 0 1 0 1.5h-1.5A.75.75 0 0 1 3 8m11 0a.75.75 0 0 1 .75-.75h1.5a.75.75 0 0 1 0 1.5h-1.5A.75.75 0 0 1 14 8m-6.828 2.828a.75.75 0 0 1 0 1.061L6.11 12.95a.75.75 0 0 1-1.06-1.06l1.06-1.06a.75.75 0 0 1 1.06 0m3.596-3.32a.75.75 0 0 0-1.37.365l-.492 6.861a.75.75 0 0 0 1.204.65l1.043-.799l.985 3.678a.75.75 0 0 0 1.45-.388l-.978-3.646l1.292.204a.75.75 0 0 0 .74-1.16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CursorArrowRipple(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.042 21.672L13.684 16.6m0 0l-2.51 2.225l.569-9.47l5.227 7.917zm-7.518-.267A8.25 8.25 0 1 1 20.25 10.5M8.288 14.212A5.25 5.25 0 1 1 17.25 10.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CursorArrowRippleSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.038 4.038a5.25 5.25 0 0 0 0 7.424a.75.75 0 0 1-1.06 1.061A6.75 6.75 0 1 1 14.5 7.75a.75.75 0 1 1-1.5 0a5.25 5.25 0 0 0-8.962-3.712"/><path fill="currentColor" d="M7.712 7.136a.75.75 0 0 1 .814.302l2.984 4.377a.75.75 0 0 1-.726 1.164l-.76-.109l.289 1.075a.75.75 0 0 1-1.45.388l-.287-1.075l-.602.474a.75.75 0 0 1-1.212-.645l.396-5.283a.75.75 0 0 1 .554-.668"/><path fill="currentColor" d="M5.805 9.695A2.75 2.75 0 1 1 10.5 7.75a.75.75 0 0 0 1.5 0a4.25 4.25 0 1 0-7.255 3.005a.75.75 0 1 0 1.06-1.06"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CursorArrowRippleSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M17.303 5.197A7.5 7.5 0 0 0 6.697 15.803a.75.75 0 0 1-1.061 1.061A9 9 0 1 1 21 10.5a.75.75 0 0 1-1.5 0c0-1.92-.732-3.839-2.197-5.303m-2.121 2.121a4.5 4.5 0 0 0-6.364 6.364a.75.75 0 1 1-1.06 1.06A6 6 0 1 1 18 10.5a.75.75 0 0 1-1.5 0c0-1.153-.44-2.303-1.318-3.182m-3.634 1.314a.75.75 0 0 1 .82.311l5.228 7.917a.75.75 0 0 1-.777 1.148l-2.097-.43l1.045 3.9a.75.75 0 0 1-1.45.388l-1.044-3.899l-1.601 1.42a.75.75 0 0 1-1.247-.606l.569-9.47a.75.75 0 0 1 .554-.68" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CursorArrowRippleTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M6.111 11.89A5.5 5.5 0 1 1 15.501 8A.75.75 0 0 0 17 8a7 7 0 1 0-11.95 4.95a.75.75 0 0 0 1.06-1.06"/><path d="M8.232 6.232a2.5 2.5 0 0 0 0 3.536a.75.75 0 1 1-1.06 1.06A4 4 0 1 1 14 8a.75.75 0 0 1-1.5 0a2.5 2.5 0 0 0-4.268-1.768"/><path d="M10.766 7.51a.75.75 0 0 0-1.37.365l-.492 6.861a.75.75 0 0 0 1.204.65l1.043-.799l.985 3.678a.75.75 0 0 0 1.45-.388l-.978-3.646l1.292.204a.75.75 0 0 0 .74-1.16z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DevicePhoneMobile(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.5 1.5H8.25A2.25 2.25 0 0 0 6 3.75v16.5a2.25 2.25 0 0 0 2.25 2.25h7.5A2.25 2.25 0 0 0 18 20.25V3.75a2.25 2.25 0 0 0-2.25-2.25H13.5m-3 0V3h3V1.5m-3 0h3m-3 18.75h3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DevicePhoneMobileSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.25 11.5a.75.75 0 0 0 0 1.5h1.5a.75.75 0 0 0 0-1.5z"/><path fill="currentColor" fill-rule="evenodd" d="M6 1a2.5 2.5 0 0 0-2.5 2.5v9A2.5 2.5 0 0 0 6 15h4a2.5 2.5 0 0 0 2.5-2.5v-9A2.5 2.5 0 0 0 10 1zm4 1.5h-.5V3a.5.5 0 0 1-.5.5H7a.5.5 0 0 1-.5-.5v-.5H6a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DevicePhoneMobileSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10.5 18.75a.75.75 0 0 0 0 1.5h3a.75.75 0 0 0 0-1.5z"/><path fill-rule="evenodd" d="M8.625.75A3.375 3.375 0 0 0 5.25 4.125v15.75a3.375 3.375 0 0 0 3.375 3.375h6.75a3.375 3.375 0 0 0 3.375-3.375V4.125A3.375 3.375 0 0 0 15.375.75zM7.5 4.125C7.5 3.504 8.004 3 8.625 3H9.75v.375c0 .621.504 1.125 1.125 1.125h2.25c.621 0 1.125-.504 1.125-1.125V3h1.125c.621 0 1.125.504 1.125 1.125v15.75c0 .621-.504 1.125-1.125 1.125h-6.75A1.125 1.125 0 0 1 7.5 19.875z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DevicePhoneMobileTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8 16.25a.75.75 0 0 1 .75-.75h2.5a.75.75 0 0 1 0 1.5h-2.5a.75.75 0 0 1-.75-.75"/><path fill-rule="evenodd" d="M4 4a3 3 0 0 1 3-3h6a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3zm4-1.5v.75c0 .414.336.75.75.75h2.5a.75.75 0 0 0 .75-.75V2.5h1A1.5 1.5 0 0 1 14.5 4v12a1.5 1.5 0 0 1-1.5 1.5H7A1.5 1.5 0 0 1 5.5 16V4A1.5 1.5 0 0 1 7 2.5z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeviceTablet(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.5 19.5h3m-6.75 2.25h10.5a2.25 2.25 0 0 0 2.25-2.25v-15a2.25 2.25 0 0 0-2.25-2.25H6.75A2.25 2.25 0 0 0 4.5 4.5v15a2.25 2.25 0 0 0 2.25 2.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeviceTabletSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M7.25 11.5a.75.75 0 0 0 0 1.5h1.5a.75.75 0 0 0 0-1.5z"/><path fill-rule="evenodd" d="M2 3.5A2.5 2.5 0 0 1 4.5 1h7A2.5 2.5 0 0 1 14 3.5v9a2.5 2.5 0 0 1-2.5 2.5h-7A2.5 2.5 0 0 1 2 12.5zm2.5-1h7a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1h-7a1 1 0 0 1-1-1v-9a1 1 0 0 1 1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeviceTabletSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10.5 18a.75.75 0 0 0 0 1.5h3a.75.75 0 0 0 0-1.5z"/><path fill-rule="evenodd" d="M7.125 1.5A3.375 3.375 0 0 0 3.75 4.875v14.25A3.375 3.375 0 0 0 7.125 22.5h9.75a3.375 3.375 0 0 0 3.375-3.375V4.875A3.375 3.375 0 0 0 16.875 1.5zM6 4.875c0-.621.504-1.125 1.125-1.125h9.75c.621 0 1.125.504 1.125 1.125v14.25c0 .621-.504 1.125-1.125 1.125h-9.75A1.125 1.125 0 0 1 6 19.125z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeviceTabletTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 1a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3zM3.5 4A1.5 1.5 0 0 1 5 2.5h10A1.5 1.5 0 0 1 16.5 4v12a1.5 1.5 0 0 1-1.5 1.5H5A1.5 1.5 0 0 1 3.5 16zm5.25 11.5a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Document(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.5 14.25v-2.625a3.375 3.375 0 0 0-3.375-3.375h-1.5A1.125 1.125 0 0 1 13.5 7.125v-1.5a3.375 3.375 0 0 0-3.375-3.375H8.25m2.25 0H5.625c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 0 0-9-9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentArrowDown(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.5 14.25v-2.625a3.375 3.375 0 0 0-3.375-3.375h-1.5A1.125 1.125 0 0 1 13.5 7.125v-1.5a3.375 3.375 0 0 0-3.375-3.375H8.25m.75 12l3 3m0 0l3-3m-3 3v-6m-1.5-9H5.625c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 0 0-9-9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentArrowDownSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 2a1.5 1.5 0 0 0-1.5 1.5v9A1.5 1.5 0 0 0 4 14h8a1.5 1.5 0 0 0 1.5-1.5V6.621a1.5 1.5 0 0 0-.44-1.06L9.94 2.439A1.5 1.5 0 0 0 8.878 2zm4 3.5a.75.75 0 0 1 .75.75v2.69l.72-.72a.75.75 0 1 1 1.06 1.06l-2 2a.75.75 0 0 1-1.06 0l-2-2a.75.75 0 0 1 1.06-1.06l.72.72V6.25A.75.75 0 0 1 8 5.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentArrowDownSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M5.625 1.5H9a3.75 3.75 0 0 1 3.75 3.75v1.875c0 1.036.84 1.875 1.875 1.875H16.5a3.75 3.75 0 0 1 3.75 3.75v7.875c0 1.035-.84 1.875-1.875 1.875H5.625a1.875 1.875 0 0 1-1.875-1.875V3.375c0-1.036.84-1.875 1.875-1.875m5.845 17.03a.75.75 0 0 0 1.06 0l3-3a.75.75 0 1 0-1.06-1.06l-1.72 1.72V12a.75.75 0 0 0-1.5 0v4.19l-1.72-1.72a.75.75 0 0 0-1.06 1.06z" clip-rule="evenodd"/><path d="M14.25 5.25a5.23 5.23 0 0 0-1.279-3.434a9.768 9.768 0 0 1 6.963 6.963A5.23 5.23 0 0 0 16.5 7.5h-1.875a.375.375 0 0 1-.375-.375z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentArrowDownTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 2A1.5 1.5 0 0 0 3 3.5v13A1.5 1.5 0 0 0 4.5 18h11a1.5 1.5 0 0 0 1.5-1.5V7.621a1.5 1.5 0 0 0-.44-1.06l-4.12-4.122A1.5 1.5 0 0 0 11.378 2zm4.75 6.75a.75.75 0 0 1 1.5 0v2.546l.943-1.048a.75.75 0 0 1 1.114 1.004l-2.25 2.5a.75.75 0 0 1-1.114 0l-2.25-2.5a.75.75 0 1 1 1.114-1.004l.943 1.048z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentArrowUp(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.5 14.25v-2.625a3.375 3.375 0 0 0-3.375-3.375h-1.5A1.125 1.125 0 0 1 13.5 7.125v-1.5a3.375 3.375 0 0 0-3.375-3.375H8.25m6.75 12l-3-3m0 0l-3 3m3-3v6m-1.5-15H5.625c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 0 0-9-9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentArrowUpSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 2a1.5 1.5 0 0 0-1.5 1.5v9A1.5 1.5 0 0 0 4 14h8a1.5 1.5 0 0 0 1.5-1.5V6.621a1.5 1.5 0 0 0-.44-1.06L9.94 2.439A1.5 1.5 0 0 0 8.878 2zm4 9.5a.75.75 0 0 1-.75-.75V8.06l-.72.72a.75.75 0 0 1-1.06-1.06l2-2a.75.75 0 0 1 1.06 0l2 2a.75.75 0 1 1-1.06 1.06l-.72-.72v2.69a.75.75 0 0 1-.75.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentArrowUpSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M5.625 1.5H9a3.75 3.75 0 0 1 3.75 3.75v1.875c0 1.036.84 1.875 1.875 1.875H16.5a3.75 3.75 0 0 1 3.75 3.75v7.875c0 1.035-.84 1.875-1.875 1.875H5.625a1.875 1.875 0 0 1-1.875-1.875V3.375c0-1.036.84-1.875 1.875-1.875m6.905 9.97a.75.75 0 0 0-1.06 0l-3 3a.75.75 0 1 0 1.06 1.06l1.72-1.72V18a.75.75 0 0 0 1.5 0v-4.19l1.72 1.72a.75.75 0 1 0 1.06-1.06z" clip-rule="evenodd"/><path d="M14.25 5.25a5.23 5.23 0 0 0-1.279-3.434a9.768 9.768 0 0 1 6.963 6.963A5.23 5.23 0 0 0 16.5 7.5h-1.875a.375.375 0 0 1-.375-.375z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentArrowUpTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 2A1.5 1.5 0 0 0 3 3.5v13A1.5 1.5 0 0 0 4.5 18h11a1.5 1.5 0 0 0 1.5-1.5V7.621a1.5 1.5 0 0 0-.44-1.06l-4.12-4.122A1.5 1.5 0 0 0 11.378 2zm4.75 11.25a.75.75 0 0 0 1.5 0v-2.546l.943 1.048a.75.75 0 1 0 1.114-1.004l-2.25-2.5a.75.75 0 0 0-1.114 0l-2.25 2.5a.75.75 0 1 0 1.114 1.004l.943-1.048z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentChartBar(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.5 14.25v-2.625a3.375 3.375 0 0 0-3.375-3.375h-1.5A1.125 1.125 0 0 1 13.5 7.125v-1.5a3.375 3.375 0 0 0-3.375-3.375H8.25M9 16.5v.75m3-3v3M15 12v5.25m-4.5-15H5.625c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 0 0-9-9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentChartBarSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 2a1.5 1.5 0 0 0-1.5 1.5v9A1.5 1.5 0 0 0 4 14h8a1.5 1.5 0 0 0 1.5-1.5V6.621a1.5 1.5 0 0 0-.44-1.06L9.94 2.439A1.5 1.5 0 0 0 8.878 2zm6 5.75a.75.75 0 0 1 1.5 0v3.5a.75.75 0 0 1-1.5 0zm-2.75 1.5a.75.75 0 0 1 1.5 0v2a.75.75 0 0 1-1.5 0zm-2 .75a.75.75 0 0 0-.75.75v.5a.75.75 0 0 0 1.5 0v-.5a.75.75 0 0 0-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentChartBarSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M5.625 1.5H9a3.75 3.75 0 0 1 3.75 3.75v1.875c0 1.036.84 1.875 1.875 1.875H16.5a3.75 3.75 0 0 1 3.75 3.75v7.875c0 1.035-.84 1.875-1.875 1.875H5.625a1.875 1.875 0 0 1-1.875-1.875V3.375c0-1.036.84-1.875 1.875-1.875M9.75 17.25a.75.75 0 0 0-1.5 0V18a.75.75 0 0 0 1.5 0zm2.25-3a.75.75 0 0 1 .75.75v3a.75.75 0 0 1-1.5 0v-3a.75.75 0 0 1 .75-.75m3.75-1.5a.75.75 0 0 0-1.5 0V18a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/><path d="M14.25 5.25a5.23 5.23 0 0 0-1.279-3.434a9.768 9.768 0 0 1 6.963 6.963A5.23 5.23 0 0 0 16.5 7.5h-1.875a.375.375 0 0 1-.375-.375z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentChartBarTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 3.5A1.5 1.5 0 0 1 4.5 2h6.879a1.5 1.5 0 0 1 1.06.44l4.122 4.12A1.5 1.5 0 0 1 17 7.622V16.5a1.5 1.5 0 0 1-1.5 1.5h-11A1.5 1.5 0 0 1 3 16.5zM13.25 9a.75.75 0 0 1 .75.75v4.5a.75.75 0 0 1-1.5 0v-4.5a.75.75 0 0 1 .75-.75m-6.5 4a.75.75 0 0 1 .75.75v.5a.75.75 0 0 1-1.5 0v-.5a.75.75 0 0 1 .75-.75m4-1.25a.75.75 0 0 0-1.5 0v2.5a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentCheck(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.125 2.25h-4.5c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125v-9M10.125 2.25h.375a9 9 0 0 1 9 9v.375M10.125 2.25A3.375 3.375 0 0 1 13.5 5.625v1.5c0 .621.504 1.125 1.125 1.125h1.5a3.375 3.375 0 0 1 3.375 3.375M9 15l2.25 2.25L15 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentCheckSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 2a1.5 1.5 0 0 0-1.5 1.5v9A1.5 1.5 0 0 0 4 14h8a1.5 1.5 0 0 0 1.5-1.5V6.621a1.5 1.5 0 0 0-.44-1.06L9.94 2.439A1.5 1.5 0 0 0 8.878 2zm6.713 4.16a.75.75 0 0 1 .127 1.053l-2.75 3.5a.75.75 0 0 1-1.078.106l-1.75-1.5a.75.75 0 1 1 .976-1.138l1.156.99L9.66 6.287a.75.75 0 0 1 1.053-.127" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentCheckSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M9 1.5H5.625c-1.036 0-1.875.84-1.875 1.875v17.25c0 1.035.84 1.875 1.875 1.875h12.75c1.035 0 1.875-.84 1.875-1.875V12.75A3.75 3.75 0 0 0 16.5 9h-1.875a1.875 1.875 0 0 1-1.875-1.875V5.25A3.75 3.75 0 0 0 9 1.5m6.61 10.936a.75.75 0 1 0-1.22-.872l-3.236 4.53L9.53 14.47a.75.75 0 0 0-1.06 1.06l2.25 2.25a.75.75 0 0 0 1.14-.094z" clip-rule="evenodd"/><path d="M12.971 1.816A5.23 5.23 0 0 1 14.25 5.25v1.875c0 .207.168.375.375.375H16.5a5.23 5.23 0 0 1 3.434 1.279a9.768 9.768 0 0 0-6.963-6.963"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentCheckTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 3.5A1.5 1.5 0 0 1 4.5 2h6.879a1.5 1.5 0 0 1 1.06.44l4.122 4.12A1.5 1.5 0 0 1 17 7.622V16.5a1.5 1.5 0 0 1-1.5 1.5h-11A1.5 1.5 0 0 1 3 16.5zm10.857 5.691a.75.75 0 0 0-1.214-.882l-3.483 4.79l-1.88-1.88a.75.75 0 0 0-1.06 1.061l2.5 2.5a.75.75 0 0 0 1.137-.089z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentDuplicate(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.75 17.25v3.375c0 .621-.504 1.125-1.125 1.125h-9.75a1.125 1.125 0 0 1-1.125-1.125V7.875c0-.621.504-1.125 1.125-1.125H6.75a9.06 9.06 0 0 1 1.5.124m7.5 10.376h3.375c.621 0 1.125-.504 1.125-1.125V11.25c0-4.46-3.243-8.161-7.5-8.876a9.06 9.06 0 0 0-1.5-.124H9.375c-.621 0-1.125.504-1.125 1.125v3.5m7.5 10.375H9.375a1.125 1.125 0 0 1-1.125-1.125v-9.25m12 6.625v-1.875a3.375 3.375 0 0 0-3.375-3.375h-1.5a1.125 1.125 0 0 1-1.125-1.125v-1.5a3.375 3.375 0 0 0-3.375-3.375H9.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentDuplicateSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 3.5A1.5 1.5 0 0 1 7 2h2.879a1.5 1.5 0 0 1 1.06.44l2.122 2.12a1.5 1.5 0 0 1 .439 1.061V9.5A1.5 1.5 0 0 1 12 11V8.621a3 3 0 0 0-.879-2.121L9 4.379A3 3 0 0 0 6.879 3.5z"/><path fill="currentColor" d="M4 5a1.5 1.5 0 0 0-1.5 1.5v6A1.5 1.5 0 0 0 4 14h5a1.5 1.5 0 0 0 1.5-1.5V8.621a1.5 1.5 0 0 0-.44-1.06L7.94 5.439A1.5 1.5 0 0 0 6.878 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentDuplicateSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M7.5 3.375c0-1.036.84-1.875 1.875-1.875h.375a3.75 3.75 0 0 1 3.75 3.75v1.875C13.5 8.161 14.34 9 15.375 9h1.875A3.75 3.75 0 0 1 21 12.75v3.375C21 17.16 20.16 18 19.125 18h-9.75A1.875 1.875 0 0 1 7.5 16.125z"/><path d="M15 5.25a5.23 5.23 0 0 0-1.279-3.434a9.768 9.768 0 0 1 6.963 6.963A5.23 5.23 0 0 0 17.25 7.5h-1.875A.375.375 0 0 1 15 7.125zM4.875 6H6v10.125A3.375 3.375 0 0 0 9.375 19.5H16.5v1.125c0 1.035-.84 1.875-1.875 1.875h-9.75A1.875 1.875 0 0 1 3 20.625V7.875C3 6.839 3.84 6 4.875 6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentDuplicateTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M7 3.5A1.5 1.5 0 0 1 8.5 2h3.879a1.5 1.5 0 0 1 1.06.44l3.122 3.12A1.5 1.5 0 0 1 17 6.622V12.5a1.5 1.5 0 0 1-1.5 1.5h-1v-3.379a3 3 0 0 0-.879-2.121L10.5 5.379A3 3 0 0 0 8.379 4.5H7z"/><path d="M4.5 6A1.5 1.5 0 0 0 3 7.5v9A1.5 1.5 0 0 0 4.5 18h7a1.5 1.5 0 0 0 1.5-1.5v-5.879a1.5 1.5 0 0 0-.44-1.06L9.44 6.439A1.5 1.5 0 0 0 8.378 6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentMagnifyingGlass(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.5 14.25v-2.625a3.375 3.375 0 0 0-3.375-3.375h-1.5A1.125 1.125 0 0 1 13.5 7.125v-1.5a3.375 3.375 0 0 0-3.375-3.375H8.25m5.231 13.481L15 17.25m-4.5-15H5.625c-.621 0-1.125.504-1.125 1.125v16.5c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 0 0-9-9m3.75 11.625a2.625 2.625 0 1 1-5.25 0a2.625 2.625 0 0 1 5.25 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentMagnifyingGlassSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 7.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0"/><path fill="currentColor" fill-rule="evenodd" d="M4 2a1.5 1.5 0 0 0-1.5 1.5v9A1.5 1.5 0 0 0 4 14h8a1.5 1.5 0 0 0 1.5-1.5V6.621a1.5 1.5 0 0 0-.44-1.06L9.94 2.439A1.5 1.5 0 0 0 8.878 2zm3.5 2.5a3 3 0 1 0 1.524 5.585l1.196 1.195a.75.75 0 1 0 1.06-1.06l-1.195-1.196A3 3 0 0 0 7.5 4.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentMagnifyingGlassSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M11.625 16.5a1.875 1.875 0 1 0 0-3.75a1.875 1.875 0 0 0 0 3.75"/><path fill-rule="evenodd" d="M5.625 1.5H9a3.75 3.75 0 0 1 3.75 3.75v1.875c0 1.036.84 1.875 1.875 1.875H16.5a3.75 3.75 0 0 1 3.75 3.75v7.875c0 1.035-.84 1.875-1.875 1.875H5.625a1.875 1.875 0 0 1-1.875-1.875V3.375c0-1.036.84-1.875 1.875-1.875m6 16.5c.66 0 1.277-.19 1.797-.518l1.048 1.048a.75.75 0 0 0 1.06-1.06l-1.047-1.048A3.375 3.375 0 1 0 11.625 18" clip-rule="evenodd"/><path d="M14.25 5.25a5.23 5.23 0 0 0-1.279-3.434a9.768 9.768 0 0 1 6.963 6.963A5.23 5.23 0 0 0 16.5 7.5h-1.875a.375.375 0 0 1-.375-.375z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentMagnifyingGlassTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8 10a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0"/><path fill-rule="evenodd" d="M4.5 2A1.5 1.5 0 0 0 3 3.5v13A1.5 1.5 0 0 0 4.5 18h11a1.5 1.5 0 0 0 1.5-1.5V7.621a1.5 1.5 0 0 0-.44-1.06l-4.12-4.122A1.5 1.5 0 0 0 11.378 2zm5 5a3 3 0 1 0 1.524 5.585l1.196 1.195a.75.75 0 1 0 1.06-1.06l-1.195-1.196A3 3 0 0 0 9.5 7" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentMinus(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.5 14.25v-2.625a3.375 3.375 0 0 0-3.375-3.375h-1.5A1.125 1.125 0 0 1 13.5 7.125v-1.5a3.375 3.375 0 0 0-3.375-3.375H8.25m6.75 12H9m1.5-12H5.625c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 0 0-9-9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentMinusSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 2a1.5 1.5 0 0 0-1.5 1.5v9A1.5 1.5 0 0 0 4 14h8a1.5 1.5 0 0 0 1.5-1.5V6.621a1.5 1.5 0 0 0-.44-1.06L9.94 2.439A1.5 1.5 0 0 0 8.878 2zm7 7a.75.75 0 0 1-.75.75h-4.5a.75.75 0 0 1 0-1.5h4.5A.75.75 0 0 1 11 9" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentMinusSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M5.625 1.5H9a3.75 3.75 0 0 1 3.75 3.75v1.875c0 1.036.84 1.875 1.875 1.875H16.5a3.75 3.75 0 0 1 3.75 3.75v7.875c0 1.035-.84 1.875-1.875 1.875H5.625a1.875 1.875 0 0 1-1.875-1.875V3.375c0-1.036.84-1.875 1.875-1.875M9.75 14.25a.75.75 0 0 0 0 1.5H15a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/><path d="M14.25 5.25a5.23 5.23 0 0 0-1.279-3.434a9.768 9.768 0 0 1 6.963 6.963A5.23 5.23 0 0 0 16.5 7.5h-1.875a.375.375 0 0 1-.375-.375z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentMinusTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 2A1.5 1.5 0 0 0 3 3.5v13A1.5 1.5 0 0 0 4.5 18h11a1.5 1.5 0 0 0 1.5-1.5V7.621a1.5 1.5 0 0 0-.44-1.06l-4.12-4.122A1.5 1.5 0 0 0 11.378 2zm7.75 9.75a.75.75 0 0 0 0-1.5h-4.5a.75.75 0 0 0 0 1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentPlus(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.5 14.25v-2.625a3.375 3.375 0 0 0-3.375-3.375h-1.5A1.125 1.125 0 0 1 13.5 7.125v-1.5a3.375 3.375 0 0 0-3.375-3.375H8.25m3.75 9v6m3-3H9m1.5-12H5.625c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 0 0-9-9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentPlusSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 2a1.5 1.5 0 0 0-1.5 1.5v9A1.5 1.5 0 0 0 4 14h8a1.5 1.5 0 0 0 1.5-1.5V6.621a1.5 1.5 0 0 0-.44-1.06L9.94 2.439A1.5 1.5 0 0 0 8.878 2zm4.75 4.75a.75.75 0 0 0-1.5 0v1.5h-1.5a.75.75 0 0 0 0 1.5h1.5v1.5a.75.75 0 0 0 1.5 0v-1.5h1.5a.75.75 0 0 0 0-1.5h-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentPlusSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M5.625 1.5H9a3.75 3.75 0 0 1 3.75 3.75v1.875c0 1.036.84 1.875 1.875 1.875H16.5a3.75 3.75 0 0 1 3.75 3.75v7.875c0 1.035-.84 1.875-1.875 1.875H5.625a1.875 1.875 0 0 1-1.875-1.875V3.375c0-1.036.84-1.875 1.875-1.875M12.75 12a.75.75 0 0 0-1.5 0v2.25H9a.75.75 0 0 0 0 1.5h2.25V18a.75.75 0 0 0 1.5 0v-2.25H15a.75.75 0 0 0 0-1.5h-2.25z" clip-rule="evenodd"/><path d="M14.25 5.25a5.23 5.23 0 0 0-1.279-3.434a9.768 9.768 0 0 1 6.963 6.963A5.23 5.23 0 0 0 16.5 7.5h-1.875a.375.375 0 0 1-.375-.375z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentPlusTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 2A1.5 1.5 0 0 0 3 3.5v13A1.5 1.5 0 0 0 4.5 18h11a1.5 1.5 0 0 0 1.5-1.5V7.621a1.5 1.5 0 0 0-.44-1.06l-4.12-4.122A1.5 1.5 0 0 0 11.378 2zM10 8a.75.75 0 0 1 .75.75v1.5h1.5a.75.75 0 0 1 0 1.5h-1.5v1.5a.75.75 0 0 1-1.5 0v-1.5h-1.5a.75.75 0 0 1 0-1.5h1.5v-1.5A.75.75 0 0 1 10 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 3.5A1.5 1.5 0 0 1 4 2h4.879a1.5 1.5 0 0 1 1.06.44l3.122 3.12a1.5 1.5 0 0 1 .439 1.061V12.5A1.5 1.5 0 0 1 12 14H4a1.5 1.5 0 0 1-1.5-1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M5.625 1.5c-1.036 0-1.875.84-1.875 1.875v17.25c0 1.035.84 1.875 1.875 1.875h12.75c1.035 0 1.875-.84 1.875-1.875V12.75A3.75 3.75 0 0 0 16.5 9h-1.875a1.875 1.875 0 0 1-1.875-1.875V5.25A3.75 3.75 0 0 0 9 1.5z"/><path d="M12.971 1.816A5.23 5.23 0 0 1 14.25 5.25v1.875c0 .207.168.375.375.375H16.5a5.23 5.23 0 0 1 3.434 1.279a9.768 9.768 0 0 0-6.963-6.963"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentText(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.5 14.25v-2.625a3.375 3.375 0 0 0-3.375-3.375h-1.5A1.125 1.125 0 0 1 13.5 7.125v-1.5a3.375 3.375 0 0 0-3.375-3.375H8.25m0 12.75h7.5m-7.5 3H12M10.5 2.25H5.625c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 0 0-9-9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentTextSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 2a1.5 1.5 0 0 0-1.5 1.5v9A1.5 1.5 0 0 0 4 14h8a1.5 1.5 0 0 0 1.5-1.5V6.621a1.5 1.5 0 0 0-.44-1.06L9.94 2.439A1.5 1.5 0 0 0 8.878 2zm1 5.75A.75.75 0 0 1 5.75 7h4.5a.75.75 0 0 1 0 1.5h-4.5A.75.75 0 0 1 5 7.75m0 3a.75.75 0 0 1 .75-.75h4.5a.75.75 0 0 1 0 1.5h-4.5a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentTextSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M5.625 1.5c-1.036 0-1.875.84-1.875 1.875v17.25c0 1.035.84 1.875 1.875 1.875h12.75c1.035 0 1.875-.84 1.875-1.875V12.75A3.75 3.75 0 0 0 16.5 9h-1.875a1.875 1.875 0 0 1-1.875-1.875V5.25A3.75 3.75 0 0 0 9 1.5zM7.5 15a.75.75 0 0 1 .75-.75h7.5a.75.75 0 0 1 0 1.5h-7.5A.75.75 0 0 1 7.5 15m.75 2.25a.75.75 0 0 0 0 1.5H12a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/><path d="M12.971 1.816A5.23 5.23 0 0 1 14.25 5.25v1.875c0 .207.168.375.375.375H16.5a5.23 5.23 0 0 1 3.434 1.279a9.768 9.768 0 0 0-6.963-6.963"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentTextTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 2A1.5 1.5 0 0 0 3 3.5v13A1.5 1.5 0 0 0 4.5 18h11a1.5 1.5 0 0 0 1.5-1.5V7.621a1.5 1.5 0 0 0-.44-1.06l-4.12-4.122A1.5 1.5 0 0 0 11.378 2zm2.25 8.5a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5zm0 3a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3.5A1.5 1.5 0 0 1 4.5 2h6.879a1.5 1.5 0 0 1 1.06.44l4.122 4.12A1.5 1.5 0 0 1 17 7.622V16.5a1.5 1.5 0 0 1-1.5 1.5h-11A1.5 1.5 0 0 1 3 16.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisHorizontal(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.75 12a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0m6 0a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0m6 0a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisHorizontalCircle(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.625 12a.375.375 0 1 1-.75 0a.375.375 0 0 1 .75 0m0 0H8.25m4.125 0a.375.375 0 1 1-.75 0a.375.375 0 0 1 .75 0m0 0H12m4.125 0a.375.375 0 1 1-.75 0a.375.375 0 0 1 .75 0m0 0h-.375M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisHorizontalCircleSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M8 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2M5.5 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0m6 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisHorizontalCircleSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75s9.75-4.365 9.75-9.75S17.385 2.25 12 2.25m0 8.625a1.125 1.125 0 1 0 0 2.25a1.125 1.125 0 0 0 0-2.25M15.375 12a1.125 1.125 0 1 1 2.25 0a1.125 1.125 0 0 1-2.25 0M7.5 10.875a1.125 1.125 0 1 0 0 2.25a1.125 1.125 0 0 0 0-2.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisHorizontalCircleTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 10a8 8 0 1 1 16 0a8 8 0 0 1-16 0m8 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2m-3-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0m7 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisHorizontalSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 8a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0m4.5 0a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0m6-1.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisHorizontalSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 12a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0m6 0a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0m6 0a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisHorizontalTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 10a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0m5.5 0a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0m7-1.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisVertical(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 6.75a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5m0 6a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5m0 6a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisVerticalSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 2a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3m0 4.5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3m1.5 6a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisVerticalSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.5 6a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0m0 6a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0m0 6a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisVerticalTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 3a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3m0 5.5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3m1.5 7a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Envelope(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21.75 6.75v10.5a2.25 2.25 0 0 1-2.25 2.25h-15a2.25 2.25 0 0 1-2.25-2.25V6.75m19.5 0A2.25 2.25 0 0 0 19.5 4.5h-15a2.25 2.25 0 0 0-2.25 2.25m19.5 0v.243a2.25 2.25 0 0 1-1.07 1.916l-7.5 4.615a2.25 2.25 0 0 1-2.36 0L3.32 8.91a2.25 2.25 0 0 1-1.07-1.916V6.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeOpen(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21.75 9v.906a2.25 2.25 0 0 1-1.183 1.981l-6.478 3.488M2.25 9v.906a2.25 2.25 0 0 0 1.183 1.981l6.478 3.488m8.839 2.51l-4.66-2.51m0 0l-1.023-.55a2.25 2.25 0 0 0-2.134 0l-1.022.55m0 0l-4.661 2.51m16.5 1.615a2.25 2.25 0 0 1-2.25 2.25h-15a2.25 2.25 0 0 1-2.25-2.25V8.844a2.25 2.25 0 0 1 1.183-1.981l7.5-4.039a2.25 2.25 0 0 1 2.134 0l7.5 4.039a2.25 2.25 0 0 1 1.183 1.98z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeOpenSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.756 4.568A1.5 1.5 0 0 0 1 5.871V12.5A1.5 1.5 0 0 0 2.5 14h11a1.5 1.5 0 0 0 1.5-1.5V5.87a1.5 1.5 0 0 0-.756-1.302l-5.5-3.143a1.5 1.5 0 0 0-1.488 0zm1.82 2.963a.75.75 0 0 0-.653 1.35l4.1 1.98a2.25 2.25 0 0 0 1.955 0l4.1-1.98a.75.75 0 1 0-.653-1.35L8.326 9.51a.75.75 0 0 1-.652 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeOpenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M19.5 22.5a3 3 0 0 0 3-3v-8.174l-6.879 4.022l3.485 1.876a.75.75 0 1 1-.712 1.321l-5.683-3.06a1.5 1.5 0 0 0-1.422 0l-5.683 3.06a.75.75 0 0 1-.712-1.32l3.485-1.877L1.5 11.326V19.5a3 3 0 0 0 3 3z"/><path d="M1.5 9.589v-.745a3 3 0 0 1 1.578-2.642l7.5-4.038a3 3 0 0 1 2.844 0l7.5 4.038A3 3 0 0 1 22.5 8.844v.745l-8.426 4.926l-.652-.351a3 3 0 0 0-2.844 0l-.652.351z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeOpenTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.106 6.447A2 2 0 0 0 1 8.237V16a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V8.236a2 2 0 0 0-1.106-1.789l-7-3.5a2 2 0 0 0-1.788 0zm1.48 4.007a.75.75 0 0 0-.671 1.342l5.855 2.928a2.75 2.75 0 0 0 2.46 0l5.852-2.927a.75.75 0 1 0-.67-1.341l-5.853 2.926a1.25 1.25 0 0 1-1.118 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M2.5 3A1.5 1.5 0 0 0 1 4.5v.793a.72.72 0 0 1 .076.032L7.674 8.51c.206.1.446.1.652 0l6.598-3.185A.755.755 0 0 1 15 5.293V4.5A1.5 1.5 0 0 0 13.5 3z"/><path d="M15 6.954L8.978 9.86a2.25 2.25 0 0 1-1.956 0L1 6.954V11.5A1.5 1.5 0 0 0 2.5 13h11a1.5 1.5 0 0 0 1.5-1.5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M1.5 8.67v8.58a3 3 0 0 0 3 3h15a3 3 0 0 0 3-3V8.67l-8.928 5.493a3 3 0 0 1-3.144 0z"/><path d="M22.5 6.908V6.75a3 3 0 0 0-3-3h-15a3 3 0 0 0-3 3v.158l9.714 5.978a1.5 1.5 0 0 0 1.572 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M3 4a2 2 0 0 0-2 2v1.161l8.441 4.221a1.25 1.25 0 0 0 1.118 0L19 7.162V6a2 2 0 0 0-2-2z"/><path d="m19 8.839l-7.77 3.885a2.75 2.75 0 0 1-2.46 0L1 8.839V14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationCircle(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 9v3.75m9-.75a9 9 0 1 1-18 0a9 9 0 0 1 18 0m-9 3.75h.008v.008H12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationCircleSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14M8 4a.75.75 0 0 1 .75.75v3a.75.75 0 0 1-1.5 0v-3A.75.75 0 0 1 8 4m0 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationCircleSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 12c0-5.385 4.365-9.75 9.75-9.75s9.75 4.365 9.75 9.75s-4.365 9.75-9.75 9.75S2.25 17.385 2.25 12M12 8.25a.75.75 0 0 1 .75.75v3.75a.75.75 0 0 1-1.5 0V9a.75.75 0 0 1 .75-.75m0 8.25a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationCircleTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M18 10a8 8 0 1 1-16 0a8 8 0 0 1 16 0m-8-5a.75.75 0 0 1 .75.75v4.5a.75.75 0 0 1-1.5 0v-4.5A.75.75 0 0 1 10 5m0 10a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationTriangle(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 9v3.75m-9.303 3.376c-.866 1.5.217 3.374 1.948 3.374h14.71c1.73 0 2.813-1.874 1.948-3.374L13.949 3.378c-.866-1.5-3.032-1.5-3.898 0zM12 15.75h.007v.008H12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationTriangleSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.701 2.25c.577-1 2.02-1 2.598 0l5.196 9a1.5 1.5 0 0 1-1.299 2.25H2.804a1.5 1.5 0 0 1-1.3-2.25zM8 4a.75.75 0 0 1 .75.75v3a.75.75 0 1 1-1.5 0v-3A.75.75 0 0 1 8 4m0 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationTriangleSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.401 3.003c1.155-2 4.043-2 5.197 0l7.355 12.748c1.154 2-.29 4.5-2.599 4.5H4.645c-2.309 0-3.752-2.5-2.598-4.5zM12 8.25a.75.75 0 0 1 .75.75v3.75a.75.75 0 0 1-1.5 0V9a.75.75 0 0 1 .75-.75m0 8.25a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationTriangleTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.485 2.495c.673-1.167 2.357-1.167 3.03 0l6.28 10.875c.673 1.167-.17 2.625-1.516 2.625H3.72c-1.347 0-2.189-1.458-1.515-2.625zM10 5a.75.75 0 0 1 .75.75v3.5a.75.75 0 0 1-1.5 0v-3.5A.75.75 0 0 1 10 5m0 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.036 12.322a1.012 1.012 0 0 1 0-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178c.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178"/><path d="M15 12a3 3 0 1 1-6 0a3 3 0 0 1 6 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeDropper(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15 11.25l1.5 1.5l.75-.75V8.758l2.276-.61a3 3 0 1 0-3.675-3.675l-.61 2.277H12l-.75.75l1.5 1.5M15 11.25l-8.47 8.47c-.34.34-.8.53-1.28.53s-.94.19-1.28.53l-.97.97l-.75-.75l.97-.97c.34-.34.53-.8.53-1.28s.19-.94.53-1.28L12.75 9M15 11.25L12.75 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeDropperSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 4a3.001 3.001 0 0 1-2.25 2.905V8.5a.75.75 0 0 1-.22.53l-.5.5a.75.75 0 0 1-1.06 0l-.72-.72l-4.677 4.678A1.75 1.75 0 0 1 4.336 14h-.672a.25.25 0 0 0-.177.073l-.707.707a.75.75 0 0 1-1.06 0l-.5-.5a.75.75 0 0 1 0-1.06l.707-.707A.25.25 0 0 0 2 12.336v-.672c0-.464.184-.909.513-1.237L7.189 5.75l-.72-.72a.75.75 0 0 1 0-1.06l.5-.5a.75.75 0 0 1 .531-.22h1.595A3.001 3.001 0 0 1 15 4M9.19 7.75l-.94-.94l-4.677 4.678a.25.25 0 0 0-.073.176v.672c0 .058-.003.115-.009.173a1.74 1.74 0 0 1 .173-.009h.672a.25.25 0 0 0 .177-.073z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeDropperSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.098 2.598a3.75 3.75 0 1 1 3.622 6.275l-1.72.46V12a.75.75 0 0 1-.22.53l-.75.75a.75.75 0 0 1-1.06 0l-.97-.97l-7.94 7.94a2.56 2.56 0 0 1-1.81.75a1.06 1.06 0 0 0-.75.31l-.97.97a.75.75 0 0 1-1.06 0l-.75-.75a.75.75 0 0 1 0-1.06l.97-.97a1.06 1.06 0 0 0 .31-.75c0-.68.27-1.33.75-1.81L11.69 9l-.97-.97a.75.75 0 0 1 0-1.06l.75-.75A.75.75 0 0 1 12 6h2.666l.461-1.72c.165-.617.49-1.2.971-1.682m-3.348 7.463L4.81 18a1.06 1.06 0 0 0-.31.75c0 .318-.06.63-.172.922a2.56 2.56 0 0 1 .922-.172c.281 0 .551-.112.75-.31l7.94-7.94z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeDropperTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.1 3.667a3.502 3.502 0 1 1 6.782 1.738a3.487 3.487 0 0 1-.907 1.57a3.495 3.495 0 0 1-1.617.919L16 7.99V10a.75.75 0 0 1-.22.53l-.25.25a.75.75 0 0 1-1.06 0l-.845-.844L7.22 16.34a2.25 2.25 0 0 1-1.591.66H5.12a.75.75 0 0 0-.53.22l-1.56 1.56a.75.75 0 0 1-1.061 0l-.75-.75a.75.75 0 0 1 0-1.06l1.56-1.561a.75.75 0 0 0 .22-.53v-.508c0-.596.237-1.169.659-1.59l6.405-6.406l-.844-.845a.75.75 0 0 1 0-1.06l.25-.25A.75.75 0 0 1 10 4h2.01zM4.72 13.84l6.405-6.405l1.44 1.439l-6.406 6.405a.75.75 0 0 1-.53.22H5.12c-.258 0-.511.044-.75.129a2.25 2.25 0 0 0 .129-.75v-.508a.75.75 0 0 1 .22-.53" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8 9.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3"/><path fill-rule="evenodd" d="M1.38 8.28a.87.87 0 0 1 0-.566a7.003 7.003 0 0 1 13.238.006a.87.87 0 0 1 0 .566A7.003 7.003 0 0 1 1.379 8.28M11 8a3 3 0 1 1-6 0a3 3 0 0 1 6 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeSlash(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.98 8.223A10.477 10.477 0 0 0 1.934 12c1.292 4.339 5.31 7.5 10.066 7.5c.993 0 1.953-.138 2.863-.395M6.228 6.228A10.451 10.451 0 0 1 12 4.5c4.756 0 8.773 3.162 10.065 7.499a10.522 10.522 0 0 1-4.293 5.773M6.228 6.228L3 3m3.228 3.228l3.65 3.65m7.894 7.894L21 21m-3.228-3.228l-3.65-3.65m0 0a3 3 0 1 0-4.243-4.243m4.242 4.242L9.88 9.88"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeSlashSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.28 2.22a.75.75 0 0 0-1.06 1.06l10.5 10.5a.75.75 0 1 0 1.06-1.06l-1.322-1.323a7.012 7.012 0 0 0 2.16-3.11a.87.87 0 0 0 0-.567A7.003 7.003 0 0 0 4.82 3.76zm3.196 3.195l1.135 1.136A1.502 1.502 0 0 1 9.45 8.389l1.136 1.135a3 3 0 0 0-4.109-4.109" clip-rule="evenodd"/><path fill="currentColor" d="m7.812 10.994l1.816 1.816A7.003 7.003 0 0 1 1.38 8.28a.87.87 0 0 1 0-.566a6.985 6.985 0 0 1 1.113-2.039l2.513 2.513a3 3 0 0 0 2.806 2.806"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeSlashSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M3.53 2.47a.75.75 0 0 0-1.06 1.06l18 18a.75.75 0 1 0 1.06-1.06zm19.146 10.083a11.249 11.249 0 0 1-2.631 4.31l-3.099-3.099a5.25 5.25 0 0 0-6.71-6.71L7.759 4.577a11.217 11.217 0 0 1 4.242-.827c4.97 0 9.185 3.223 10.675 7.69c.12.362.12.752 0 1.113"/><path d="M15.75 12c0 .18-.013.357-.037.53l-4.244-4.243A3.75 3.75 0 0 1 15.75 12m-3.22 3.713l-4.243-4.244a3.75 3.75 0 0 0 4.244 4.243"/><path d="M6.75 12c0-.619.107-1.213.304-1.764l-3.1-3.1a11.25 11.25 0 0 0-2.63 4.31c-.12.362-.12.752 0 1.114c1.489 4.467 5.704 7.69 10.675 7.69c1.5 0 2.933-.294 4.242-.827l-2.477-2.477A5.25 5.25 0 0 1 6.75 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeSlashTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M3.28 2.22a.75.75 0 0 0-1.06 1.06l14.5 14.5a.75.75 0 1 0 1.06-1.06l-1.745-1.745a10.029 10.029 0 0 0 3.3-4.38a1.651 1.651 0 0 0 0-1.185A10.004 10.004 0 0 0 9.999 3a9.956 9.956 0 0 0-4.744 1.194zm4.472 4.47l1.092 1.092a2.5 2.5 0 0 1 3.374 3.373l1.091 1.092A4 4 0 0 0 7.752 6.69" clip-rule="evenodd"/><path d="m10.748 13.93l2.523 2.523a9.987 9.987 0 0 1-3.27.547c-4.258 0-7.894-2.66-9.337-6.41a1.651 1.651 0 0 1 0-1.186A10.007 10.007 0 0 1 2.839 6.02L6.07 9.252a4 4 0 0 0 4.678 4.678"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/><path fill-rule="evenodd" d="M1.323 11.447C2.811 6.976 7.028 3.75 12.001 3.75c4.97 0 9.185 3.223 10.675 7.69c.12.362.12.752 0 1.113c-1.487 4.471-5.705 7.697-10.677 7.697c-4.97 0-9.186-3.223-10.675-7.69a1.762 1.762 0 0 1 0-1.113M17.25 12a5.25 5.25 0 1 1-10.5 0a5.25 5.25 0 0 1 10.5 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10 12.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5"/><path fill-rule="evenodd" d="M.664 10.59a1.651 1.651 0 0 1 0-1.186A10.004 10.004 0 0 1 10 3c4.257 0 7.893 2.66 9.336 6.41c.147.381.146.804 0 1.186A10.004 10.004 0 0 1 10 17c-4.257 0-7.893-2.66-9.336-6.41M14 10a4 4 0 1 1-8 0a4 4 0 0 1 8 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceFrown(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.182 16.318A4.486 4.486 0 0 0 12.016 15a4.486 4.486 0 0 0-3.198 1.318M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0M9.75 9.75c0 .414-.168.75-.375.75S9 10.164 9 9.75S9.168 9 9.375 9s.375.336.375.75m-.375 0h.008v.015h-.008zm5.625 0c0 .414-.168.75-.375.75s-.375-.336-.375-.75s.168-.75.375-.75s.375.336.375.75m-.375 0h.008v.015h-.008z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceFrownSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M6 8c.552 0 1-.672 1-1.5S6.552 5 6 5s-1 .672-1 1.5S5.448 8 6 8m5-1.5c0 .828-.448 1.5-1 1.5s-1-.672-1-1.5S9.448 5 10 5s1 .672 1 1.5m-6.005 5.805a.75.75 0 0 0 1.06 0a2.75 2.75 0 0 1 3.89 0a.75.75 0 0 0 1.06-1.06a4.25 4.25 0 0 0-6.01 0a.75.75 0 0 0 0 1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceFrownSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75s9.75-4.365 9.75-9.75S17.385 2.25 12 2.25m-2.625 6c-.54 0-.828.419-.936.634a1.96 1.96 0 0 0-.189.866c0 .298.059.605.189.866c.108.215.395.634.936.634c.54 0 .828-.419.936-.634c.13-.26.189-.568.189-.866c0-.298-.059-.605-.189-.866c-.108-.215-.395-.634-.936-.634m4.314.634c.108-.215.395-.634.936-.634c.54 0 .828.419.936.634c.13.26.189.568.189.866c0 .298-.059.605-.189.866c-.108.215-.395.634-.936.634c-.54 0-.828-.419-.936-.634a1.96 1.96 0 0 1-.189-.866c0-.298.059-.605.189-.866m-4.34 7.964a.75.75 0 0 1-1.061-1.06a5.236 5.236 0 0 1 3.73-1.538a5.236 5.236 0 0 1 3.695 1.538a.75.75 0 1 1-1.061 1.06a3.736 3.736 0 0 0-2.639-1.098a3.736 3.736 0 0 0-2.664 1.098" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceFrownTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 18a8 8 0 1 0 0-16a8 8 0 0 0 0 16m-3.536-3.475a.75.75 0 0 0 1.061 0a3.5 3.5 0 0 1 4.95 0a.75.75 0 1 0 1.06-1.06a5 5 0 0 0-7.07 0a.75.75 0 0 0 0 1.06M9 8.5c0 .828-.448 1.5-1 1.5s-1-.672-1-1.5S7.448 7 8 7s1 .672 1 1.5m3 1.5c.552 0 1-.672 1-1.5S12.552 7 12 7s-1 .672-1 1.5s.448 1.5 1 1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceSmile(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.182 15.182a4.5 4.5 0 0 1-6.364 0M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0M9.75 9.75c0 .414-.168.75-.375.75S9 10.164 9 9.75S9.168 9 9.375 9s.375.336.375.75m-.375 0h.008v.015h-.008zm5.625 0c0 .414-.168.75-.375.75s-.375-.336-.375-.75s.168-.75.375-.75s.375.336.375.75m-.375 0h.008v.015h-.008z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceSmileSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M6 8c.552 0 1-.672 1-1.5S6.552 5 6 5s-1 .672-1 1.5S5.448 8 6 8m5-1.5c0 .828-.448 1.5-1 1.5s-1-.672-1-1.5S9.448 5 10 5s1 .672 1 1.5m.005 4.245a.75.75 0 0 0-1.06 0a2.75 2.75 0 0 1-3.89 0a.75.75 0 0 0-1.06 1.06a4.25 4.25 0 0 0 6.01 0a.75.75 0 0 0 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceSmileSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75s9.75-4.365 9.75-9.75S17.385 2.25 12 2.25m-2.625 6c-.54 0-.828.419-.936.634a1.96 1.96 0 0 0-.189.866c0 .298.059.605.189.866c.108.215.395.634.936.634c.54 0 .828-.419.936-.634c.13-.26.189-.568.189-.866c0-.298-.059-.605-.189-.866c-.108-.215-.395-.634-.936-.634m4.314.634c.108-.215.395-.634.936-.634c.54 0 .828.419.936.634c.13.26.189.568.189.866c0 .298-.059.605-.189.866c-.108.215-.395.634-.936.634c-.54 0-.828-.419-.936-.634a1.96 1.96 0 0 1-.189-.866c0-.298.059-.605.189-.866m2.023 6.828a.75.75 0 1 0-1.06-1.06a3.75 3.75 0 0 1-5.304 0a.75.75 0 0 0-1.06 1.06a5.25 5.25 0 0 0 7.424 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceSmileTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 18a8 8 0 1 0 0-16a8 8 0 0 0 0 16m3.536-4.464a.75.75 0 1 0-1.061-1.061a3.5 3.5 0 0 1-4.95 0a.75.75 0 0 0-1.06 1.06a5 5 0 0 0 7.07 0M9 8.5c0 .828-.448 1.5-1 1.5s-1-.672-1-1.5S7.448 7 8 7s1 .672 1 1.5m3 1.5c.552 0 1-.672 1-1.5S12.552 7 12 7s-1 .672-1 1.5s.448 1.5 1 1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Film(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.375 19.5h17.25m-17.25 0a1.125 1.125 0 0 1-1.125-1.125M3.375 19.5h1.5C5.496 19.5 6 18.996 6 18.375m-3.75 0V5.625m0 12.75v-1.5c0-.621.504-1.125 1.125-1.125m18.375 2.625V5.625m0 12.75c0 .621-.504 1.125-1.125 1.125m1.125-1.125v-1.5c0-.621-.504-1.125-1.125-1.125m0 3.75h-1.5A1.125 1.125 0 0 1 18 18.375M20.625 4.5H3.375m17.25 0c.621 0 1.125.504 1.125 1.125M20.625 4.5h-1.5C18.504 4.5 18 5.004 18 5.625m3.75 0v1.5c0 .621-.504 1.125-1.125 1.125M3.375 4.5c-.621 0-1.125.504-1.125 1.125M3.375 4.5h1.5C5.496 4.5 6 5.004 6 5.625m-3.75 0v1.5c0 .621.504 1.125 1.125 1.125m0 0h1.5m-1.5 0c-.621 0-1.125.504-1.125 1.125v1.5c0 .621.504 1.125 1.125 1.125m1.5-3.75C5.496 8.25 6 7.746 6 7.125v-1.5M4.875 8.25C5.496 8.25 6 8.754 6 9.375v1.5m0-5.25v5.25m0-5.25C6 5.004 6.504 4.5 7.125 4.5h9.75c.621 0 1.125.504 1.125 1.125m1.125 2.625h1.5m-1.5 0A1.125 1.125 0 0 1 18 7.125v-1.5m1.125 2.625c-.621 0-1.125.504-1.125 1.125v1.5m2.625-2.625c.621 0 1.125.504 1.125 1.125v1.5c0 .621-.504 1.125-1.125 1.125M18 5.625v5.25M7.125 12h9.75m-9.75 0A1.125 1.125 0 0 1 6 10.875M7.125 12C6.504 12 6 12.504 6 13.125m0-2.25C6 11.496 5.496 12 4.875 12M18 10.875c0 .621-.504 1.125-1.125 1.125M18 10.875c0 .621.504 1.125 1.125 1.125m-2.25 0c.621 0 1.125.504 1.125 1.125m-12 5.25v-5.25m0 5.25c0 .621.504 1.125 1.125 1.125h9.75c.621 0 1.125-.504 1.125-1.125m-12 0v-1.5c0-.621-.504-1.125-1.125-1.125M18 18.375v-5.25m0 5.25v-1.5c0-.621.504-1.125 1.125-1.125M18 13.125v1.5c0 .621.504 1.125 1.125 1.125M18 13.125c0-.621.504-1.125 1.125-1.125M6 13.125v1.5c0 .621-.504 1.125-1.125 1.125M6 13.125C6 12.504 5.496 12 4.875 12m-1.5 0h1.5m-1.5 0c-.621 0-1.125.504-1.125 1.125v1.5c0 .621.504 1.125 1.125 1.125M19.125 12h1.5m0 0c.621 0 1.125.504 1.125 1.125v1.5c0 .621-.504 1.125-1.125 1.125m-17.25 0h1.5m14.25 0h1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilmSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 3.5A1.5 1.5 0 0 1 2.5 2h11A1.5 1.5 0 0 1 15 3.5v9a1.5 1.5 0 0 1-1.5 1.5h-11A1.5 1.5 0 0 1 1 12.5zm1.5.25a.25.25 0 0 1 .25-.25h1.5a.25.25 0 0 1 .25.25v1a.25.25 0 0 1-.25.25h-1.5a.25.25 0 0 1-.25-.25zm3.75-.25a.25.25 0 0 0-.25.25v3.5c0 .138.112.25.25.25h3.5a.25.25 0 0 0 .25-.25v-3.5a.25.25 0 0 0-.25-.25zM6 8.75a.25.25 0 0 1 .25-.25h3.5a.25.25 0 0 1 .25.25v3.5a.25.25 0 0 1-.25.25h-3.5a.25.25 0 0 1-.25-.25zm5.75-5.25a.25.25 0 0 0-.25.25v1c0 .138.112.25.25.25h1.5a.25.25 0 0 0 .25-.25v-1a.25.25 0 0 0-.25-.25zM2.5 11.25a.25.25 0 0 1 .25-.25h1.5a.25.25 0 0 1 .25.25v1a.25.25 0 0 1-.25.25h-1.5a.25.25 0 0 1-.25-.25zm9.25-.25a.25.25 0 0 0-.25.25v1c0 .138.112.25.25.25h1.5a.25.25 0 0 0 .25-.25v-1a.25.25 0 0 0-.25-.25zM2.5 8.75a.25.25 0 0 1 .25-.25h1.5a.25.25 0 0 1 .25.25v1a.25.25 0 0 1-.25.25h-1.5a.25.25 0 0 1-.25-.25zm9.25-.25a.25.25 0 0 0-.25.25v1c0 .138.112.25.25.25h1.5a.25.25 0 0 0 .25-.25v-1a.25.25 0 0 0-.25-.25zM2.5 6.25A.25.25 0 0 1 2.75 6h1.5a.25.25 0 0 1 .25.25v1a.25.25 0 0 1-.25.25h-1.5a.25.25 0 0 1-.25-.25zM11.75 6a.25.25 0 0 0-.25.25v1c0 .138.112.25.25.25h1.5a.25.25 0 0 0 .25-.25v-1a.25.25 0 0 0-.25-.25z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilmSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 5.625c0-1.036.84-1.875 1.875-1.875h17.25c1.035 0 1.875.84 1.875 1.875v12.75c0 1.035-.84 1.875-1.875 1.875H3.375A1.875 1.875 0 0 1 1.5 18.375zm1.5 0v1.5c0 .207.168.375.375.375h1.5a.375.375 0 0 0 .375-.375v-1.5a.375.375 0 0 0-.375-.375h-1.5A.375.375 0 0 0 3 5.625m16.125-.375a.375.375 0 0 0-.375.375v1.5c0 .207.168.375.375.375h1.5A.375.375 0 0 0 21 7.125v-1.5a.375.375 0 0 0-.375-.375zM21 9.375A.375.375 0 0 0 20.625 9h-1.5a.375.375 0 0 0-.375.375v1.5c0 .207.168.375.375.375h1.5a.375.375 0 0 0 .375-.375zm0 3.75a.375.375 0 0 0-.375-.375h-1.5a.375.375 0 0 0-.375.375v1.5c0 .207.168.375.375.375h1.5a.375.375 0 0 0 .375-.375zm0 3.75a.375.375 0 0 0-.375-.375h-1.5a.375.375 0 0 0-.375.375v1.5c0 .207.168.375.375.375h1.5a.375.375 0 0 0 .375-.375zM4.875 18.75a.375.375 0 0 0 .375-.375v-1.5a.375.375 0 0 0-.375-.375h-1.5a.375.375 0 0 0-.375.375v1.5c0 .207.168.375.375.375zM3.375 15h1.5a.375.375 0 0 0 .375-.375v-1.5a.375.375 0 0 0-.375-.375h-1.5a.375.375 0 0 0-.375.375v1.5c0 .207.168.375.375.375m0-3.75h1.5a.375.375 0 0 0 .375-.375v-1.5A.375.375 0 0 0 4.875 9h-1.5A.375.375 0 0 0 3 9.375v1.5c0 .207.168.375.375.375m4.125 0a.75.75 0 0 0 0 1.5h9a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilmTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 4.75C1 3.784 1.784 3 2.75 3h14.5c.966 0 1.75.784 1.75 1.75v10.515a1.75 1.75 0 0 1-1.75 1.75h-1.5c-.078 0-.155-.005-.23-.015H4.48c-.075.01-.152.015-.23.015h-1.5A1.75 1.75 0 0 1 1 15.265zm16.5 7.385V11.01a.25.25 0 0 0-.25-.25h-1.5a.25.25 0 0 0-.25.25v1.125c0 .138.112.25.25.25h1.5a.25.25 0 0 0 .25-.25m0 2.005a.25.25 0 0 0-.25-.25h-1.5a.25.25 0 0 0-.25.25v1.125c0 .108.069.2.165.235h1.585a.25.25 0 0 0 .25-.25zm-15 1.11v-1.11a.25.25 0 0 1 .25-.25h1.5a.25.25 0 0 1 .25.25v1.125a.25.25 0 0 1-.164.235H2.75a.25.25 0 0 1-.25-.25m2-4.24v1.125a.25.25 0 0 1-.25.25h-1.5a.25.25 0 0 1-.25-.25V11.01a.25.25 0 0 1 .25-.25h1.5a.25.25 0 0 1 .25.25m13-2.005V7.88a.25.25 0 0 0-.25-.25h-1.5a.25.25 0 0 0-.25.25v1.125c0 .138.112.25.25.25h1.5a.25.25 0 0 0 .25-.25M4.25 7.63a.25.25 0 0 1 .25.25v1.125a.25.25 0 0 1-.25.25h-1.5a.25.25 0 0 1-.25-.25V7.88a.25.25 0 0 1 .25-.25zm0-3.13a.25.25 0 0 1 .25.25v1.125a.25.25 0 0 1-.25.25h-1.5a.25.25 0 0 1-.25-.25V4.75a.25.25 0 0 1 .25-.25zm11.5 1.625a.25.25 0 0 1-.25-.25V4.75a.25.25 0 0 1 .25-.25h1.5a.25.25 0 0 1 .25.25v1.125a.25.25 0 0 1-.25.25zm-9 3.125a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FingerPrint(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.864 4.243A7.5 7.5 0 0 1 19.5 10.5c0 2.92-.556 5.709-1.568 8.269M5.742 6.364A7.465 7.465 0 0 0 4.5 10.5a7.464 7.464 0 0 1-1.15 3.993m1.989 3.559A11.209 11.209 0 0 0 8.25 10.5a3.75 3.75 0 1 1 7.5 0c0 .527-.021 1.049-.064 1.565M12 10.5a14.94 14.94 0 0 1-3.6 9.75m6.633-4.596a18.666 18.666 0 0 1-2.485 5.33"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FingerPrintSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 3c-.988 0-1.908.286-2.682.78a.75.75 0 0 1-.806-1.266A6.5 6.5 0 0 1 14.5 8c0 1.665-.333 3.254-.936 4.704a.75.75 0 0 1-1.385-.577C12.708 10.857 13 9.464 13 8a5 5 0 0 0-5-5M3.55 4.282a.75.75 0 0 1 .23 1.036A4.973 4.973 0 0 0 3 8a.75.75 0 0 1-1.5 0c0-1.282.372-2.48 1.014-3.488a.75.75 0 0 1 1.036-.23M8 5.875A2.125 2.125 0 0 0 5.875 8a3.625 3.625 0 0 1-3.625 3.625h-.037a.75.75 0 1 1 .008-1.5h.03A2.125 2.125 0 0 0 4.376 8a3.625 3.625 0 1 1 7.25 0c0 .078-.001.156-.003.233a.75.75 0 1 1-1.5-.036A6.47 6.47 0 0 0 10.126 8A2.125 2.125 0 0 0 8 5.875M7.995 7.25a.75.75 0 0 1 .75.75a6.502 6.502 0 0 1-4.343 6.134a.75.75 0 1 1-.498-1.416A5.002 5.002 0 0 0 7.245 8a.75.75 0 0 1 .75-.75m2.651 2.87a.75.75 0 0 1 .463.955a9.39 9.39 0 0 1-3.008 4.25a.75.75 0 0 1-.936-1.171a7.892 7.892 0 0 0 2.527-3.57a.75.75 0 0 1 .954-.463" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FingerPrintSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 3.75a6.715 6.715 0 0 0-3.722 1.118a.75.75 0 1 1-.828-1.25a8.25 8.25 0 0 1 12.8 6.883a23.19 23.19 0 0 1-1.62 8.543a.75.75 0 0 1-1.395-.551A21.69 21.69 0 0 0 18.75 10.5A6.75 6.75 0 0 0 12 3.75M6.157 5.739a.75.75 0 0 1 .21 1.04A6.715 6.715 0 0 0 5.25 10.5c0 1.613-.463 3.12-1.265 4.393a.75.75 0 0 1-1.27-.8A6.715 6.715 0 0 0 3.75 10.5c0-1.68.503-3.246 1.367-4.55a.75.75 0 0 1 1.04-.211M12 7.5a3 3 0 0 0-3 3c0 3.1-1.176 5.927-3.105 8.056a.75.75 0 1 1-1.112-1.008A10.459 10.459 0 0 0 7.5 10.5a4.5 4.5 0 1 1 9 0c0 .547-.022 1.09-.067 1.626a.75.75 0 0 1-1.495-.123A18.2 18.2 0 0 0 15 10.5a3 3 0 0 0-3-3m0 2.25a.75.75 0 0 1 .75.75a15.69 15.69 0 0 1-3.781 10.238a.75.75 0 0 1-1.14-.975A14.19 14.19 0 0 0 11.25 10.5a.75.75 0 0 1 .75-.75m3.239 5.183a.75.75 0 0 1 .515.927a19.417 19.417 0 0 1-2.585 5.544a.75.75 0 0 1-1.243-.84a17.915 17.915 0 0 0 2.386-5.116a.75.75 0 0 1 .927-.515" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FingerPrintTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 2.5c-1.31 0-2.526.386-3.546 1.051a.75.75 0 0 1-.82-1.256A8 8 0 0 1 18 9a22.47 22.47 0 0 1-1.228 7.351a.75.75 0 1 1-1.417-.49A20.97 20.97 0 0 0 16.5 9A6.5 6.5 0 0 0 10 2.5M4.333 4.416a.75.75 0 0 1 .218 1.038A6.466 6.466 0 0 0 3.5 9a7.966 7.966 0 0 1-1.293 4.362a.75.75 0 0 1-1.257-.819A6.466 6.466 0 0 0 2 9c0-1.61.476-3.11 1.295-4.365a.75.75 0 0 1 1.038-.219M10 6.12a3 3 0 0 0-3.001 3.041a11.455 11.455 0 0 1-2.697 7.24a.75.75 0 0 1-1.148-.965A9.957 9.957 0 0 0 5.5 9c0-.028.002-.055.004-.082a4.5 4.5 0 0 1 8.996.084v.148l-.005.297a.75.75 0 1 1-1.5-.034c.003-.11.004-.219.005-.328a3 3 0 0 0-3-2.965m0 2.13a.75.75 0 0 1 .75.75c0 3.51-1.187 6.746-3.181 9.323a.75.75 0 1 1-1.186-.918A13.687 13.687 0 0 0 9.25 9a.75.75 0 0 1 .75-.75m3.529 3.698a.75.75 0 0 1 .584.885a18.883 18.883 0 0 1-2.257 5.84a.75.75 0 1 1-1.29-.764a17.386 17.386 0 0 0 2.078-5.377a.75.75 0 0 1 .885-.584" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fire(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15.362 5.214A8.252 8.252 0 0 1 12 21A8.25 8.25 0 0 1 6.038 7.047A8.287 8.287 0 0 0 9 9.601a8.983 8.983 0 0 1 3.361-6.867a8.21 8.21 0 0 0 3 2.48"/><path d="M12 18a3.75 3.75 0 0 0 .495-7.468a5.99 5.99 0 0 0-1.925 3.547a5.975 5.975 0 0 1-2.133-1.001A3.75 3.75 0 0 0 12 18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FireSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.074.945A4.993 4.993 0 0 0 6 5v.032c.004.6.114 1.176.311 1.709c.16.428-.204.91-.61.7a5.023 5.023 0 0 1-1.868-1.677c-.202-.304-.648-.363-.848-.058a6 6 0 1 0 8.017-1.901l-.004-.007a4.98 4.98 0 0 1-2.18-2.574c-.116-.31-.477-.472-.744-.28m.78 6.178a3.001 3.001 0 1 1-3.473 4.341c-.205-.365.215-.694.62-.59a4.008 4.008 0 0 0 1.873.03c.288-.065.413-.386.321-.666A3.997 3.997 0 0 1 8 8.999c0-.585.126-1.14.351-1.641a.42.42 0 0 1 .503-.235" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FireSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.963 2.286a.75.75 0 0 0-1.071-.136a9.742 9.742 0 0 0-3.539 6.176a7.547 7.547 0 0 1-1.705-1.715a.75.75 0 0 0-1.152-.082A9 9 0 1 0 15.68 4.534a7.46 7.46 0 0 1-2.717-2.248M15.75 14.25a3.75 3.75 0 1 1-7.313-1.172c.628.465 1.35.81 2.133 1a5.99 5.99 0 0 1 1.925-3.546a3.75 3.75 0 0 1 3.255 3.718" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FireTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 4.938a7 7 0 1 1-9.006 1.737c.202-.257.59-.218.793.039c.278.352.594.672.943.954c.332.269.786-.049.773-.476a5.977 5.977 0 0 1 .572-2.759a6.026 6.026 0 0 1 2.486-2.665c.247-.14.55-.016.677.238A6.967 6.967 0 0 0 13.5 4.938M14 12a4 4 0 0 1-4 4c-1.913 0-3.52-1.398-3.91-3.182c-.093-.429.44-.643.814-.413a4.043 4.043 0 0 0 1.601.564c.303.038.531-.24.51-.544a5.975 5.975 0 0 1 1.315-4.192a.447.447 0 0 1 .431-.16A4.001 4.001 0 0 1 14 12" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 3v1.5M3 21v-6m0 0l2.77-.693a9 9 0 0 1 6.208.682l.108.054a9 9 0 0 0 6.086.71l3.114-.732a48.524 48.524 0 0 1-.005-10.499l-3.11.732a9 9 0 0 1-6.085-.711l-.108-.054a9 9 0 0 0-6.208-.682L3 4.5M3 15V4.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.75 2a.75.75 0 0 0-.75.75v10.5a.75.75 0 0 0 1.5 0v-2.624l.33-.083A6.044 6.044 0 0 1 8 11c1.29.645 2.77.807 4.17.457l1.48-.37a.462.462 0 0 0 .35-.448V3.56a.438.438 0 0 0-.544-.425l-1.287.322A6.035 6.035 0 0 1 8 3a6.045 6.045 0 0 0-4.17-.457l-.34.085A.75.75 0 0 0 2.75 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 2.25a.75.75 0 0 1 .75.75v.54l1.838-.46a9.75 9.75 0 0 1 6.725.738l.108.054A8.25 8.25 0 0 0 18 4.524l3.11-.732a.75.75 0 0 1 .917.81a47.784 47.784 0 0 0 .005 10.337a.75.75 0 0 1-.574.812l-3.114.733a9.75 9.75 0 0 1-6.594-.77l-.108-.054a8.25 8.25 0 0 0-5.69-.625l-2.202.55V21a.75.75 0 0 1-1.5 0V3A.75.75 0 0 1 3 2.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 2.75a.75.75 0 0 0-1.5 0v14.5a.75.75 0 0 0 1.5 0v-4.392l1.657-.348a6.449 6.449 0 0 1 4.271.572a7.948 7.948 0 0 0 5.965.524l2.078-.64A.75.75 0 0 0 18 12.25v-8.5a.75.75 0 0 0-.904-.734l-2.38.501a7.25 7.25 0 0 1-4.186-.363l-.502-.2a8.75 8.75 0 0 0-5.053-.439l-1.475.31z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.25 12.75V12A2.25 2.25 0 0 1 4.5 9.75h15A2.25 2.25 0 0 1 21.75 12v.75m-8.69-6.44l-2.12-2.12a1.5 1.5 0 0 0-1.061-.44H4.5A2.25 2.25 0 0 0 2.25 6v12a2.25 2.25 0 0 0 2.25 2.25h15A2.25 2.25 0 0 0 21.75 18V9a2.25 2.25 0 0 0-2.25-2.25h-5.379a1.5 1.5 0 0 1-1.06-.44"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderArrowDown(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 13.5l3 3m0 0l3-3m-3 3v-6m1.06-4.19l-2.12-2.12a1.5 1.5 0 0 0-1.061-.44H4.5A2.25 2.25 0 0 0 2.25 6v12a2.25 2.25 0 0 0 2.25 2.25h15A2.25 2.25 0 0 0 21.75 18V9a2.25 2.25 0 0 0-2.25-2.25h-5.379a1.5 1.5 0 0 1-1.06-.44"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderArrowDownSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 2A1.5 1.5 0 0 0 2 3.5v9A1.5 1.5 0 0 0 3.5 14h9a1.5 1.5 0 0 0 1.5-1.5v-7A1.5 1.5 0 0 0 12.5 4H9.621a1.5 1.5 0 0 1-1.06-.44L7.439 2.44A1.5 1.5 0 0 0 6.38 2zm5.25 4.75a.75.75 0 0 0-1.5 0v2.69l-.72-.72a.75.75 0 0 0-1.06 1.06l2 2a.75.75 0 0 0 1.06 0l2-2a.75.75 0 1 0-1.06-1.06l-.72.72z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderArrowDownSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M19.5 21a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3h-5.379a.75.75 0 0 1-.53-.22L11.47 3.66A2.25 2.25 0 0 0 9.879 3H4.5a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3zm-6.75-10.5a.75.75 0 0 0-1.5 0v4.19l-1.72-1.72a.75.75 0 0 0-1.06 1.06l3 3a.75.75 0 0 0 1.06 0l3-3a.75.75 0 1 0-1.06-1.06l-1.72 1.72z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderArrowDownTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 4.75C2 3.784 2.784 3 3.75 3h4.836c.464 0 .909.184 1.237.513l1.414 1.414a.25.25 0 0 0 .177.073h4.836c.966 0 1.75.784 1.75 1.75v8.5A1.75 1.75 0 0 1 16.25 17H3.75A1.75 1.75 0 0 1 2 15.25zm8.75 4a.75.75 0 0 0-1.5 0v2.546l-.943-1.048a.75.75 0 1 0-1.114 1.004l2.25 2.5a.75.75 0 0 0 1.114 0l2.25-2.5a.75.75 0 1 0-1.114-1.004l-.943 1.048z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderMinus(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 13.5H9m4.06-7.19l-2.12-2.12a1.5 1.5 0 0 0-1.061-.44H4.5A2.25 2.25 0 0 0 2.25 6v12a2.25 2.25 0 0 0 2.25 2.25h15A2.25 2.25 0 0 0 21.75 18V9a2.25 2.25 0 0 0-2.25-2.25h-5.379a1.5 1.5 0 0 1-1.06-.44"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderMinusSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 2A1.5 1.5 0 0 0 2 3.5v9A1.5 1.5 0 0 0 3.5 14h9a1.5 1.5 0 0 0 1.5-1.5v-7A1.5 1.5 0 0 0 12.5 4H9.621a1.5 1.5 0 0 1-1.06-.44L7.439 2.44A1.5 1.5 0 0 0 6.38 2zm6.75 7.75a.75.75 0 0 0 0-1.5h-4.5a.75.75 0 0 0 0 1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderMinusSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M19.5 21a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3h-5.379a.75.75 0 0 1-.53-.22L11.47 3.66A2.25 2.25 0 0 0 9.879 3H4.5a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3zM9 12.75a.75.75 0 0 0 0 1.5h6a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderMinusTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 4.75C2 3.784 2.784 3 3.75 3h4.836c.464 0 .909.184 1.237.513l1.414 1.414a.25.25 0 0 0 .177.073h4.836c.966 0 1.75.784 1.75 1.75v8.5A1.75 1.75 0 0 1 16.25 17H3.75A1.75 1.75 0 0 1 2 15.25zm10.25 7a.75.75 0 0 0 0-1.5h-4.5a.75.75 0 0 0 0 1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpen(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 9.776c.112-.017.227-.026.344-.026h15.812c.117 0 .232.009.344.026m-16.5 0a2.25 2.25 0 0 0-1.883 2.542l.857 6a2.25 2.25 0 0 0 2.227 1.932H19.05a2.25 2.25 0 0 0 2.227-1.932l.857-6a2.25 2.25 0 0 0-1.883-2.542m-16.5 0V6A2.25 2.25 0 0 1 6 3.75h3.879a1.5 1.5 0 0 1 1.06.44l2.122 2.12a1.5 1.5 0 0 0 1.06.44H18A2.25 2.25 0 0 1 20.25 9v.776"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpenSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3.5A1.5 1.5 0 0 1 4.5 2h1.879a1.5 1.5 0 0 1 1.06.44l1.122 1.12A1.5 1.5 0 0 0 9.62 4h1.88A1.5 1.5 0 0 1 13 5.5v1H3zM3.081 8a1.5 1.5 0 0 0-1.423 1.974l1 3A1.5 1.5 0 0 0 4.081 14h7.838a1.5 1.5 0 0 0 1.423-1.026l1-3A1.5 1.5 0 0 0 12.919 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.906 9c.382 0 .749.057 1.094.162V9a3 3 0 0 0-3-3h-3.879a.75.75 0 0 1-.53-.22L11.47 3.66A2.25 2.25 0 0 0 9.879 3H6a3 3 0 0 0-3 3v3.162A3.756 3.756 0 0 1 4.094 9zM4.094 10.5a2.25 2.25 0 0 0-2.227 2.568l.857 6A2.25 2.25 0 0 0 4.951 21H19.05a2.25 2.25 0 0 0 2.227-1.932l.857-6a2.25 2.25 0 0 0-2.227-2.568z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpenTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.75 3A1.75 1.75 0 0 0 3 4.75v2.752l.104-.002h13.792c.035 0 .07 0 .104.002V6.75A1.75 1.75 0 0 0 15.25 5h-3.836a.25.25 0 0 1-.177-.073L9.823 3.513A1.75 1.75 0 0 0 8.586 3zM3.104 9a1.75 1.75 0 0 0-1.673 2.265l1.385 4.5A1.75 1.75 0 0 0 4.488 17h11.023a1.75 1.75 0 0 0 1.673-1.235l1.384-4.5A1.75 1.75 0 0 0 16.896 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderPlus(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 10.5v6m3-3H9m4.06-7.19l-2.12-2.12a1.5 1.5 0 0 0-1.061-.44H4.5A2.25 2.25 0 0 0 2.25 6v12a2.25 2.25 0 0 0 2.25 2.25h15A2.25 2.25 0 0 0 21.75 18V9a2.25 2.25 0 0 0-2.25-2.25h-5.379a1.5 1.5 0 0 1-1.06-.44"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderPlusSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 2A1.5 1.5 0 0 0 2 3.5v9A1.5 1.5 0 0 0 3.5 14h9a1.5 1.5 0 0 0 1.5-1.5v-7A1.5 1.5 0 0 0 12.5 4H9.621a1.5 1.5 0 0 1-1.06-.44L7.439 2.44A1.5 1.5 0 0 0 6.38 2zM8 6a.75.75 0 0 1 .75.75v1.5h1.5a.75.75 0 0 1 0 1.5h-1.5v1.5a.75.75 0 0 1-1.5 0v-1.5h-1.5a.75.75 0 0 1 0-1.5h1.5v-1.5A.75.75 0 0 1 8 6" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderPlusSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M19.5 21a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3h-5.379a.75.75 0 0 1-.53-.22L11.47 3.66A2.25 2.25 0 0 0 9.879 3H4.5a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3zm-6.75-10.5a.75.75 0 0 0-1.5 0v2.25H9a.75.75 0 0 0 0 1.5h2.25v2.25a.75.75 0 0 0 1.5 0v-2.25H15a.75.75 0 0 0 0-1.5h-2.25z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderPlusTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.75 3A1.75 1.75 0 0 0 2 4.75v10.5c0 .966.784 1.75 1.75 1.75h12.5A1.75 1.75 0 0 0 18 15.25v-8.5A1.75 1.75 0 0 0 16.25 5h-4.836a.25.25 0 0 1-.177-.073L9.823 3.513A1.75 1.75 0 0 0 8.586 3zM10 8a.75.75 0 0 1 .75.75v1.5h1.5a.75.75 0 0 1 0 1.5h-1.5v1.5a.75.75 0 0 1-1.5 0v-1.5h-1.5a.75.75 0 0 1 0-1.5h1.5v-1.5A.75.75 0 0 1 10 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 3.5A1.5 1.5 0 0 1 3.5 2h2.879a1.5 1.5 0 0 1 1.06.44l1.122 1.12A1.5 1.5 0 0 0 9.62 4h2.88A1.5 1.5 0 0 1 14 5.5v1.401a2.986 2.986 0 0 0-1.5-.401h-9c-.546 0-1.059.146-1.5.401zm0 6v3A1.5 1.5 0 0 0 3.5 14h9a1.5 1.5 0 0 0 1.5-1.5v-3A1.5 1.5 0 0 0 12.5 8h-9A1.5 1.5 0 0 0 2 9.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.5 21a3 3 0 0 0 3-3v-4.5a3 3 0 0 0-3-3h-15a3 3 0 0 0-3 3V18a3 3 0 0 0 3 3zm-18-10.854V6a3 3 0 0 1 3-3h5.379a2.25 2.25 0 0 1 1.59.659l2.122 2.121c.14.141.331.22.53.22H19.5a3 3 0 0 1 3 3v1.146A4.483 4.483 0 0 0 19.5 9h-15a4.483 4.483 0 0 0-3 1.146"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.75 3A1.75 1.75 0 0 0 2 4.75v3.26a3.235 3.235 0 0 1 1.75-.51h12.5c.644 0 1.245.188 1.75.51V6.75A1.75 1.75 0 0 0 16.25 5h-4.836a.25.25 0 0 1-.177-.073L9.823 3.513A1.75 1.75 0 0 0 8.586 3zm0 6A1.75 1.75 0 0 0 2 10.75v4.5c0 .966.784 1.75 1.75 1.75h12.5A1.75 1.75 0 0 0 18 15.25v-4.5A1.75 1.75 0 0 0 16.25 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Forward(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 8.689c0-.864.933-1.406 1.683-.977l7.108 4.061a1.125 1.125 0 0 1 0 1.954l-7.108 4.061A1.125 1.125 0 0 1 3 16.811zm9.75 0c0-.864.933-1.406 1.683-.977l7.108 4.061a1.125 1.125 0 0 1 0 1.954l-7.108 4.061a1.125 1.125 0 0 1-1.683-.977z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.53 3.956A1 1 0 0 0 1 4.804v6.392a1 1 0 0 0 1.53.848l5.113-3.196c.16-.1.279-.233.357-.383v2.73a1 1 0 0 0 1.53.849l5.113-3.196a1 1 0 0 0 0-1.696L9.53 3.956A1 1 0 0 0 8 4.804v2.731a.992.992 0 0 0-.357-.383z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.055 7.06c-1.25-.713-2.805.19-2.805 1.63v8.122c0 1.44 1.555 2.343 2.805 1.628L12 14.471v2.34c0 1.44 1.555 2.343 2.805 1.628l7.108-4.061c1.26-.72 1.26-2.536 0-3.256l-7.108-4.061C13.555 6.346 12 7.249 12 8.689v2.34z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.288 4.818A1.5 1.5 0 0 0 1 6.095v7.81a1.5 1.5 0 0 0 2.288 1.276l6.323-3.905c.155-.096.285-.213.389-.344v2.973a1.5 1.5 0 0 0 2.288 1.276l6.323-3.905a1.5 1.5 0 0 0 0-2.552l-6.323-3.906A1.5 1.5 0 0 0 10 6.095v2.972a1.506 1.506 0 0 0-.389-.343z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Funnel(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 3c2.755 0 5.455.232 8.083.678c.533.09.917.556.917 1.096v1.044a2.25 2.25 0 0 1-.659 1.591l-5.432 5.432a2.25 2.25 0 0 0-.659 1.591v2.927a2.25 2.25 0 0 1-1.244 2.013L9.75 21v-6.568a2.25 2.25 0 0 0-.659-1.591L3.659 7.409A2.25 2.25 0 0 1 3 5.818V4.774c0-.54.384-1.006.917-1.096A48.32 48.32 0 0 1 12 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FunnelSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 2a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v2.172a2 2 0 0 0 .586 1.414l2.828 2.828A2 2 0 0 1 6 9.828v4.363a.5.5 0 0 0 .724.447l2.17-1.085A2 2 0 0 0 10 11.763V9.829a2 2 0 0 1 .586-1.414l2.828-2.828A2 2 0 0 0 14 4.172z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FunnelSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.792 2.938A49.069 49.069 0 0 1 12 2.25c2.797 0 5.54.236 8.209.688a1.857 1.857 0 0 1 1.541 1.836v1.044a3 3 0 0 1-.879 2.121l-6.182 6.182a1.5 1.5 0 0 0-.439 1.061v2.927a3 3 0 0 1-1.658 2.684l-1.757.878A.75.75 0 0 1 9.75 21v-5.818a1.5 1.5 0 0 0-.44-1.06L3.13 7.938a3 3 0 0 1-.879-2.121V4.774c0-.897.64-1.683 1.542-1.836" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FunnelTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.628 1.601C5.028 1.206 7.49 1 10 1s4.973.206 7.372.601a.75.75 0 0 1 .628.74v2.288a2.25 2.25 0 0 1-.659 1.59l-4.682 4.683a2.25 2.25 0 0 0-.659 1.59v3.037c0 .684-.31 1.33-.844 1.757l-1.937 1.55A.75.75 0 0 1 8 18.25v-5.757a2.25 2.25 0 0 0-.659-1.591L2.659 6.22A2.25 2.25 0 0 1 2 4.629V2.34a.75.75 0 0 1 .628-.74" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gif(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.75 8.25v7.5m6-7.5h-3V12m0 0v3.75m0-3.75H18M9.75 9.348c-1.03-1.464-2.698-1.464-3.728 0c-1.03 1.465-1.03 3.84 0 5.304c1.03 1.464 2.699 1.464 3.728 0V12h-1.5M4.5 19.5h15a2.25 2.25 0 0 0 2.25-2.25V6.75A2.25 2.25 0 0 0 19.5 4.5h-15a2.25 2.25 0 0 0-2.25 2.25v10.5A2.25 2.25 0 0 0 4.5 19.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GifSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 3a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2zm.895 3.458C4.142 6.071 4.38 6 4.5 6s.358.07.605.458a.75.75 0 1 0 1.265-.805C5.933 4.966 5.274 4.5 4.5 4.5s-1.433.466-1.87 1.153C2.195 6.336 2 7.187 2 8s.195 1.664.63 2.347c.437.687 1.096 1.153 1.87 1.153s1.433-.466 1.87-1.153a.75.75 0 0 0 .117-.402V8a.75.75 0 0 0-.75-.75H5a.75.75 0 0 0-.013 1.5v.955C4.785 9.95 4.602 10 4.5 10c-.121 0-.358-.07-.605-.458C3.647 9.15 3.5 8.595 3.5 8c0-.595.147-1.15.395-1.542M9 5.25a.75.75 0 0 0-1.5 0v5.5a.75.75 0 0 0 1.5 0zm1 0a.75.75 0 0 1 .75-.75h2.5a.75.75 0 0 1 0 1.5H11.5v1.25h.75a.75.75 0 0 1 0 1.5h-.75v2a.75.75 0 0 1-1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GifSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 3.75a3 3 0 0 0-3 3v10.5a3 3 0 0 0 3 3h15a3 3 0 0 0 3-3V6.75a3 3 0 0 0-3-3zm9 4.5a.75.75 0 0 0-1.5 0v7.5a.75.75 0 0 0 1.5 0zm1.5 0a.75.75 0 0 1 .75-.75h3a.75.75 0 0 1 0 1.5H16.5v2.25H18a.75.75 0 0 1 0 1.5h-1.5v3a.75.75 0 0 1-1.5 0zM6.636 9.78c.404-.575.867-.78 1.25-.78s.846.205 1.25.78a.75.75 0 0 0 1.228-.863C9.738 8.027 8.853 7.5 7.886 7.5c-.966 0-1.852.527-2.477 1.417c-.62.882-.909 2-.909 3.083c0 1.083.289 2.201.909 3.083c.625.89 1.51 1.417 2.477 1.417c.967 0 1.852-.527 2.478-1.417a.75.75 0 0 0 .136-.431V12a.75.75 0 0 0-.75-.75h-1.5a.75.75 0 0 0 0 1.5H9v1.648c-.37.44-.774.602-1.114.602c-.383 0-.846-.205-1.25-.78C6.226 13.638 6 12.837 6 12c0-.837.226-1.638.636-2.22" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GifTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 5.25A2.25 2.25 0 0 1 3.25 3h13.5A2.25 2.25 0 0 1 19 5.25v9.5A2.25 2.25 0 0 1 16.75 17H3.25A2.25 2.25 0 0 1 1 14.75zm4.026 2.879C5.356 7.65 5.72 7.5 6 7.5s.643.15.974.629a.75.75 0 0 0 1.234-.854C7.66 6.484 6.873 6 6 6c-.873 0-1.66.484-2.208 1.275C3.25 8.059 3 9.048 3 10c0 .952.25 1.941.792 2.725C4.34 13.516 5.127 14 6 14c.873 0 1.66-.484 2.208-1.275a.75.75 0 0 0 .133-.427V10a.75.75 0 0 0-.75-.75H6.25a.75.75 0 0 0 0 1.5h.591v1.295c-.293.342-.6.455-.841.455c-.279 0-.643-.15-.974-.629C4.69 11.386 4.5 10.711 4.5 10c0-.711.19-1.386.526-1.871M10.75 6a.75.75 0 0 1 .75.75v6.5a.75.75 0 0 1-1.5 0v-6.5a.75.75 0 0 1 .75-.75m3 0h2.5a.75.75 0 0 1 0 1.5H14.5v1.75h.75a.75.75 0 0 1 0 1.5h-.75v2.5a.75.75 0 0 1-1.5 0v-6.5a.75.75 0 0 1 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gift(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 11.25v8.25a1.5 1.5 0 0 1-1.5 1.5H5.25a1.5 1.5 0 0 1-1.5-1.5v-8.25M12 4.875A2.625 2.625 0 1 0 9.375 7.5H12m0-2.625V7.5m0-2.625A2.625 2.625 0 1 1 14.625 7.5H12m0 0V21m-8.625-9.75h18c.621 0 1.125-.504 1.125-1.125v-1.5c0-.621-.504-1.125-1.125-1.125h-18c-.621 0-1.125.504-1.125 1.125v1.5c0 .621.504 1.125 1.125 1.125"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GiftSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.75 3.5c0 .563.186 1.082.5 1.5H2a1 1 0 0 0 0 2h5.25V5h1.5v2H14a1 1 0 1 0 0-2h-2.25A2.5 2.5 0 0 0 8 1.714A2.5 2.5 0 0 0 3.75 3.5m3.499 0v-.038A1 1 0 1 0 6.25 4.5h1zm2.5-1a1 1 0 0 0-1 .962l.001.038v1h.999a1 1 0 0 0 0-2" clip-rule="evenodd"/><path fill="currentColor" d="M7.25 8.5H2V12a2 2 0 0 0 2 2h3.25zm1.5 5.5V8.5H14V12a2 2 0 0 1-2 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GiftSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.375 3a1.875 1.875 0 0 0 0 3.75h1.875v4.5H3.375A1.875 1.875 0 0 1 1.5 9.375v-.75c0-1.036.84-1.875 1.875-1.875h3.193A3.375 3.375 0 0 1 12 2.754a3.375 3.375 0 0 1 5.432 3.997h3.943a1.874 1.874 0 0 1 1.875 1.874v.75c0 1.036-.84 1.875-1.875 1.875H12.75v-4.5h1.875a1.875 1.875 0 1 0-1.875-1.875V6.75h-1.5V4.875C11.25 3.839 10.41 3 9.375 3m1.875 9.75H3v6.75a2.25 2.25 0 0 0 2.25 2.25h6zm1.5 0v9h6.75a2.25 2.25 0 0 0 2.25-2.25v-6.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GiftTop(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 3.75v16.5M2.25 12h19.5M6.375 17.25a4.875 4.875 0 0 0 4.875-4.875V12m6.375 5.25a4.875 4.875 0 0 1-4.875-4.875V12m-9 8.25h16.5a1.5 1.5 0 0 0 1.5-1.5V5.25a1.5 1.5 0 0 0-1.5-1.5H3.75a1.5 1.5 0 0 0-1.5 1.5v13.5a1.5 1.5 0 0 0 1.5 1.5m12.621-9.44c-1.409 1.41-4.242 1.061-4.242 1.061s-.349-2.833 1.06-4.242a2.25 2.25 0 0 1 3.182 3.182M10.773 7.63c1.409 1.409 1.06 4.242 1.06 4.242S9 12.22 7.592 10.811a2.25 2.25 0 1 1 3.182-3.182"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GiftTopSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.25 2H3.5A1.5 1.5 0 0 0 2 3.5v3.75h1.718A2.5 2.5 0 0 1 7.25 3.716zM2 8.75v3.75A1.5 1.5 0 0 0 3.5 14h3.75v-3.085a4.743 4.743 0 0 1-3.455 1.826a.75.75 0 1 1-.092-1.497a3.252 3.252 0 0 0 2.96-2.494zM8.75 14h3.75a1.5 1.5 0 0 0 1.5-1.5V8.75H9.337a3.252 3.252 0 0 0 2.96 2.494a.75.75 0 1 1-.093 1.497a4.743 4.743 0 0 1-3.454-1.826zM14 7.25h-1.718A2.5 2.5 0 0 0 8.75 3.717V2h3.75A1.5 1.5 0 0 1 14 3.5z"/><path fill="currentColor" d="M6.352 6.787c.16.012.312.014.448.012c.002-.136 0-.289-.012-.448c-.043-.617-.203-1.181-.525-1.503a1 1 0 0 0-1.414 1.414c.322.322.886.482 1.503.525m3.297 0c-.16.012-.312.014-.448.012c-.003-.136 0-.289.011-.448c.044-.617.203-1.181.526-1.503a1 1 0 1 1 1.414 1.414c-.322.322-.887.482-1.503.525"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GiftTopSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M11.25 3v4.046a3 3 0 0 0-4.277 4.204H1.5v-6A2.25 2.25 0 0 1 3.75 3zm1.5 0v4.011a3 3 0 0 1 4.239 4.239H22.5v-6A2.25 2.25 0 0 0 20.25 3zm9.75 9.75h-8.983a4.125 4.125 0 0 0 4.108 3.75a.75.75 0 0 1 0 1.5a5.623 5.623 0 0 1-4.875-2.817V21h7.5a2.25 2.25 0 0 0 2.25-2.25zM11.25 21v-5.817A5.623 5.623 0 0 1 6.375 18a.75.75 0 0 1 0-1.5a4.126 4.126 0 0 0 4.108-3.75H1.5v6A2.25 2.25 0 0 0 3.75 21z"/><path d="M11.086 10.354c.03.297.037.575.035.805a7.484 7.484 0 0 1-.805-.036c-.833-.084-1.677-.325-2.195-.843a1.5 1.5 0 0 1 2.122-2.12c.517.517.759 1.36.842 2.194m1.792 0c-.03.297-.038.575-.036.805c.23.002.508-.006.805-.036c.833-.084 1.677-.325 2.195-.843a1.5 1.5 0 0 0-2.121-2.12c-.518.518-.76 1.362-.843 2.194"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GiftTopTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.25 3H3.5A1.5 1.5 0 0 0 2 4.5v4.75h3.365A2.75 2.75 0 0 1 9.25 5.362zM2 10.75v4.75A1.5 1.5 0 0 0 3.5 17h5.75v-4.876A4.75 4.75 0 0 1 5 14.75a.75.75 0 0 1 0-1.5a3.251 3.251 0 0 0 3.163-2.5zM10.75 17h5.75a1.5 1.5 0 0 0 1.5-1.5v-4.75h-6.163A3.251 3.251 0 0 0 15 13.25a.75.75 0 0 1 0 1.5a4.75 4.75 0 0 1-4.25-2.626zM18 9.25V4.5A1.5 1.5 0 0 0 16.5 3h-5.75v2.362a2.75 2.75 0 0 1 3.885 3.888zm-4.496-2.755a1.25 1.25 0 0 0-1.768 0c-.36.359-.526.999-.559 1.697c-.01.228-.006.443.004.626c.183.01.398.014.626.003c.698-.033 1.338-.2 1.697-.559a1.25 1.25 0 0 0 0-1.767m-5.24 0a1.25 1.25 0 0 0-1.768 1.767c.36.36 1 .526 1.697.56c.228.01.443.006.626-.004c.01-.183.015-.398.004-.626c-.033-.698-.2-1.338-.56-1.697" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GiftTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M14 6a2.5 2.5 0 0 0-4-3a2.5 2.5 0 0 0-4 3H3.25C2.56 6 2 6.56 2 7.25v.5C2 8.44 2.56 9 3.25 9h6V6h1.5v3h6C17.44 9 18 8.44 18 7.75v-.5C18 6.56 17.44 6 16.75 6zm-1-1.5a1 1 0 0 1-1 1h-1v-1a1 1 0 1 1 2 0m-6 0a1 1 0 0 0 1 1h1v-1a1 1 0 0 0-2 0" clip-rule="evenodd"/><path d="M9.25 10.5H3v4.75A2.75 2.75 0 0 0 5.75 18h3.5zm1.5 7.5v-7.5H17v4.75A2.75 2.75 0 0 1 14.25 18z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeAlt(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 21a9.004 9.004 0 0 0 8.716-6.747M12 21a9.004 9.004 0 0 1-8.716-6.747M12 21c2.485 0 4.5-4.03 4.5-9S14.485 3 12 3m0 18c-2.485 0-4.5-4.03-4.5-9S9.515 3 12 3m0 0a8.997 8.997 0 0 1 7.843 4.582M12 3a8.997 8.997 0 0 0-7.843 4.582m15.686 0A11.953 11.953 0 0 1 12 10.5c-2.998 0-5.74-1.1-7.843-2.918m15.686 0A8.959 8.959 0 0 1 21 12c0 .778-.099 1.533-.284 2.253m0 0A17.919 17.919 0 0 1 12 16.5a17.92 17.92 0 0 1-8.716-2.247m0 0A9.015 9.015 0 0 1 3 12c0-1.605.42-3.113 1.157-4.418"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeAltSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.757 4.5c.18.217.376.42.586.608c.153-.61.354-1.175.596-1.678A5.53 5.53 0 0 0 3.757 4.5M8 1a6.994 6.994 0 0 0-7 7a7 7 0 1 0 7-7m0 1.5c-.476 0-1.091.386-1.633 1.427c-.293.564-.531 1.267-.683 2.063A5.48 5.48 0 0 0 8 6.5a5.48 5.48 0 0 0 2.316-.51c-.152-.796-.39-1.499-.683-2.063C9.09 2.886 8.476 2.5 8 2.5m3.657 2.608a8.823 8.823 0 0 0-.596-1.678c.444.298.842.659 1.182 1.07c-.18.217-.376.42-.586.608m-1.166 2.436A6.983 6.983 0 0 1 8 8a6.983 6.983 0 0 1-2.49-.456a10.703 10.703 0 0 0 .202 2.6c.72.231 1.49.356 2.288.356c.798 0 1.568-.125 2.29-.356a10.705 10.705 0 0 0 .2-2.6m1.433 1.85a12.652 12.652 0 0 0 .018-2.609a6.98 6.98 0 0 0 1.117-.947a5.48 5.48 0 0 1 .44 2.262a7.536 7.536 0 0 1-1.575 1.293m-2.172 2.435a9.046 9.046 0 0 1-3.504 0c.039.084.078.166.12.244C6.907 13.114 7.523 13.5 8 13.5s1.091-.386 1.633-1.427c.04-.078.08-.16.12-.244m1.31.74a8.5 8.5 0 0 0 .492-1.298c.457-.197.893-.43 1.307-.696a5.526 5.526 0 0 1-1.8 1.995m-6.123 0a8.507 8.507 0 0 1-.493-1.298a8.985 8.985 0 0 1-1.307-.696a5.526 5.526 0 0 0 1.8 1.995M2.5 8.1c.463.5.993.935 1.575 1.293a12.652 12.652 0 0 1-.018-2.608a7.037 7.037 0 0 1-1.117-.947A5.48 5.48 0 0 0 2.5 8.1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeAltSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.721 12.752a9.711 9.711 0 0 0-.945-5.003a12.754 12.754 0 0 1-4.339 2.708a18.991 18.991 0 0 1-.214 4.772a17.165 17.165 0 0 0 5.498-2.477m-7.087 2.798a17.324 17.324 0 0 0 .332-4.647c-.952.227-1.945.347-2.966.347c-1.021 0-2.014-.12-2.966-.347a17.515 17.515 0 0 0 .332 4.647a17.385 17.385 0 0 0 5.268 0m-4.862 1.569a18.963 18.963 0 0 0 4.456 0A17.182 17.182 0 0 1 12 21.724a17.18 17.18 0 0 1-2.228-4.605M7.777 15.23a18.87 18.87 0 0 1-.214-4.773a12.753 12.753 0 0 1-4.34-2.709a9.711 9.711 0 0 0-.944 5.004a17.165 17.165 0 0 0 5.498 2.477m13.579-.476a9.765 9.765 0 0 1-7.478 6.816a18.64 18.64 0 0 0 1.988-4.718a18.627 18.627 0 0 0 5.49-2.098m-18.712 0c1.682.97 3.53 1.687 5.49 2.098a18.64 18.64 0 0 0 1.988 4.718a9.765 9.765 0 0 1-7.478-6.816M13.878 2.43a9.755 9.755 0 0 1 6.116 3.986a11.267 11.267 0 0 1-3.746 2.504a18.63 18.63 0 0 0-2.37-6.49M12 2.276a17.152 17.152 0 0 1 2.805 7.121a11.281 11.281 0 0 1-5.61 0A17.151 17.151 0 0 1 12 2.276m-1.878.154a18.629 18.629 0 0 0-2.37 6.49a11.266 11.266 0 0 1-3.746-2.504a9.754 9.754 0 0 1 6.116-3.985"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeAltTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.555 5.412a8.028 8.028 0 0 0-3.503-2.81a14.899 14.899 0 0 1 1.663 4.472a8.547 8.547 0 0 0 1.84-1.662m-3.229 2.413a13.43 13.43 0 0 0-2.413-5.773a8.087 8.087 0 0 0-1.826 0a13.43 13.43 0 0 0-2.413 5.773A8.473 8.473 0 0 0 10 8.5c1.18 0 2.304-.24 3.326-.675M6.514 9.376A9.98 9.98 0 0 0 10 10c1.226 0 2.4-.22 3.486-.624a13.54 13.54 0 0 1-.351 3.759A13.54 13.54 0 0 1 10 13.5c-1.079 0-2.128-.127-3.134-.366a13.538 13.538 0 0 1-.352-3.758M5.285 7.074a14.9 14.9 0 0 1 1.663-4.471a8.028 8.028 0 0 0-3.503 2.81a8.577 8.577 0 0 0 1.84 1.66m12.049-.275a7.973 7.973 0 0 1 .614 4.115a13.47 13.47 0 0 1-3.178 1.72a15.093 15.093 0 0 0 .174-3.939a10.043 10.043 0 0 0 2.39-1.896m-14.668 0a10.042 10.042 0 0 0 2.39 1.896a15.196 15.196 0 0 0 .174 3.94a13.472 13.472 0 0 1-3.178-1.72a7.973 7.973 0 0 1 .615-4.115M10 15c.898 0 1.778-.079 2.633-.23a13.473 13.473 0 0 1-1.72 3.178a8.099 8.099 0 0 1-1.826 0a13.47 13.47 0 0 1-1.72-3.178c.855.151 1.735.23 2.633.23m4.357-.643a14.912 14.912 0 0 1-1.305 3.04a8.027 8.027 0 0 0 4.345-4.345c-.953.542-1.971.981-3.04 1.305m-7.409 3.04a8.027 8.027 0 0 1-4.345-4.345c.953.542 1.971.981 3.04 1.305a14.912 14.912 0 0 0 1.305 3.04"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeAmericas(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6.115 5.19l.319 1.913A6 6 0 0 0 8.11 10.36L9.75 12l-.387.775c-.217.433-.132.956.21 1.298l1.348 1.348c.21.21.329.497.329.795v1.089c0 .426.24.815.622 1.006l.153.076c.433.217.956.132 1.298-.21l.723-.723a8.7 8.7 0 0 0 2.288-4.042a1.087 1.087 0 0 0-.358-1.099l-1.33-1.108a1.122 1.122 0 0 0-.905-.245l-1.17.195a1.125 1.125 0 0 1-.98-.314l-.295-.295a1.125 1.125 0 0 1 0-1.591l.13-.132a1.125 1.125 0 0 1 1.3-.21l.603.302a.809.809 0 0 0 1.086-1.086L14.25 7.5l1.256-.837a4.5 4.5 0 0 0 1.528-1.732l.146-.292M6.115 5.19a9 9 0 1 0 11.065-.55m-11.065.55A8.965 8.965 0 0 1 12 3c1.929 0 3.716.607 5.18 1.64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeAmericasSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1a7 7 0 1 0 0 14A7 7 0 0 0 8 1M4.5 3.757a5.5 5.5 0 1 0 6.857-.114l-.65.65a.707.707 0 0 0-.207.5c0 .39-.317.707-.707.707H8.427a.496.496 0 0 0-.413.771l.25.376a.481.481 0 0 0 .616.163a.962.962 0 0 1 1.11.18l.573.573a1 1 0 0 1 .242 1.023l-1.012 3.035a1 1 0 0 1-1.191.654l-.345-.086a1 1 0 0 1-.757-.97v-.305a1 1 0 0 0-.293-.707L6.1 9.1a.849.849 0 0 1 0-1.2c.22-.22.22-.58 0-.8l-.721-.721A3 3 0 0 1 4.5 4.257z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeAmericasSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75s9.75-4.365 9.75-9.75S17.385 2.25 12 2.25M6.262 6.072a8.25 8.25 0 1 0 10.562-.766a4.5 4.5 0 0 1-1.318 1.357L14.25 7.5l.165.33a.809.809 0 0 1-1.086 1.085l-.604-.302a1.125 1.125 0 0 0-1.298.21l-.132.131a1.127 1.127 0 0 0 0 1.592l.296.295c.256.257.622.374.98.314l1.17-.195c.323-.054.654.036.905.245l1.33 1.108c.32.267.46.694.358 1.1a8.7 8.7 0 0 1-2.288 4.04l-.723.724a1.125 1.125 0 0 1-1.298.21l-.153-.076a1.125 1.125 0 0 1-.622-1.006v-1.089c0-.298-.119-.585-.33-.796l-1.347-1.347a1.125 1.125 0 0 1-.21-1.298L9.75 12l-1.64-1.64a6 6 0 0 1-1.676-3.257z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeAmericasTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M18 10a8 8 0 1 1-16 0a8 8 0 0 1 16 0m-1.5 0a6.5 6.5 0 1 1-11-4.69v.447a3.5 3.5 0 0 0 1.025 2.475L8.293 10L8 10.293a1 1 0 0 0 0 1.414l1.06 1.06a1.5 1.5 0 0 1 .44 1.061v.363a1 1 0 0 0 .553.894l.276.139a1 1 0 0 0 1.342-.448l1.454-2.908a1.5 1.5 0 0 0-.281-1.731l-.772-.772a1 1 0 0 0-1.023-.242l-.384.128a.5.5 0 0 1-.606-.25l-.296-.592a.481.481 0 0 1 .646-.646l.262.131a1 1 0 0 0 .447.106h.188a1 1 0 0 0 .949-1.316l-.068-.204a.5.5 0 0 1 .149-.538l1.44-1.234A6.492 6.492 0 0 1 16.5 10" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeAsiaAustralia(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.75 3.03v.568c0 .334.148.65.405.864l1.068.89c.442.369.535 1.01.216 1.49l-.51.766a2.25 2.25 0 0 1-1.161.886l-.143.048a1.107 1.107 0 0 0-.57 1.664a1.108 1.108 0 0 1-.427 1.605L9 13.125l.423 1.059a.956.956 0 0 1-1.652.928l-.679-.906a1.125 1.125 0 0 0-1.906.172L4.5 15.75l-.612.153M12.75 3.031a9 9 0 0 0-8.862 12.872M12.75 3.031a9 9 0 0 1 6.69 14.036m0 0l-.177-.529A2.25 2.25 0 0 0 17.128 15H16.5l-.324-.324a1.453 1.453 0 0 0-2.328.377l-.036.073a1.586 1.586 0 0 1-.982.816l-.99.282c-.55.157-.894.702-.8 1.267l.073.438c.08.474.49.821.97.821c.846 0 1.598.542 1.865 1.345l.215.643m5.276-3.67a9.012 9.012 0 0 1-5.276 3.67m0 0a9 9 0 0 1-10.275-4.835M15.75 9c0 .896-.393 1.7-1.016 2.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeAsiaAustraliaSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 8a7 7 0 1 1 14 0A7 7 0 0 1 1 8m7 5.5a5.485 5.485 0 0 1-4.007-1.732l.28-.702a.402.402 0 0 1 .658-.135a.804.804 0 0 0 1.138 0l.012-.012a.822.822 0 0 0 .154-.949l-.055-.11a.497.497 0 0 1 .134-.611L8.14 7.788a.57.57 0 0 0 .154-.7a.57.57 0 0 1 .33-.796l.028-.01a1.788 1.788 0 0 0 1.13-1.13l.072-.214a.747.747 0 0 0-.18-.764L8.293 2.793A1 1 0 0 1 8.09 2.5a5.5 5.5 0 0 1 4.81 8h-.486a1 1 0 0 1-.707-.293l-.353-.353a1.207 1.207 0 0 0-1.708 0l-.531.531a1 1 0 0 1-.26.188l-.343.17a.927.927 0 0 0-.512.83v.177c0 .414.336.75.75.75a.75.75 0 0 1 .751.793c-.477.135-.98.207-1.501.207" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeAsiaAustraliaSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M15.75 8.25a.75.75 0 0 1 .75.75a3.74 3.74 0 0 1-1.27 2.812a.75.75 0 1 1-.992-1.124A2.243 2.243 0 0 0 15 9a.75.75 0 0 1 .75-.75"/><path fill-rule="evenodd" d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75s9.75-4.365 9.75-9.75S17.385 2.25 12 2.25M4.575 15.6a8.25 8.25 0 0 0 9.348 4.425a1.966 1.966 0 0 0-1.84-1.275a.983.983 0 0 1-.97-.822l-.073-.437c-.094-.565.25-1.11.8-1.267l.99-.282c.427-.123.783-.418.982-.816l.036-.073a1.453 1.453 0 0 1 2.328-.377L16.5 15h.628a2.25 2.25 0 0 1 1.983 1.186a8.25 8.25 0 0 0-6.345-12.4c.044.262.18.503.389.676l1.068.89c.442.369.535 1.01.216 1.49l-.51.766a2.25 2.25 0 0 1-1.161.886l-.143.048a1.107 1.107 0 0 0-.57 1.664a1.108 1.108 0 0 1-.427 1.605L9 13.125l.423 1.059a.956.956 0 0 1-1.652.928l-.679-.906a1.125 1.125 0 0 0-1.906.172z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeAsiaAustraliaTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M18 10a8 8 0 1 1-16 0a8 8 0 0 1 16 0m-6.5 6.326a6.52 6.52 0 0 1-1.5.174a6.487 6.487 0 0 1-5.011-2.36l.49-.98a.423.423 0 0 1 .614-.164l.294.196a.992.992 0 0 0 1.491-1.139l-.197-.593a.252.252 0 0 1 .126-.304l1.973-.987a.938.938 0 0 0 .361-1.359a.375.375 0 0 1 .239-.576l.125-.025A2.421 2.421 0 0 0 12.327 6.6l.05-.149a1 1 0 0 0-.242-1.023l-1.489-1.489a.5.5 0 0 1-.146-.353v-.067a6.5 6.5 0 0 1 5.392 9.23a1.398 1.398 0 0 0-.68-.244l-.566-.566a1.5 1.5 0 0 0-1.06-.439h-.172a1.5 1.5 0 0 0-1.06.44l-.593.592a.501.501 0 0 1-.13.093l-1.578.79a1 1 0 0 0-.553.894v.191a1 1 0 0 0 1 1h.5a.5.5 0 0 1 .5.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeEuropeAfrica(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20.893 13.393l-1.135-1.135a2.252 2.252 0 0 1-.421-.585l-1.08-2.16a.414.414 0 0 0-.663-.107a.827.827 0 0 1-.812.21l-1.273-.363a.89.89 0 0 0-.738 1.595l.587.39c.59.395.674 1.23.172 1.732l-.2.2c-.212.212-.33.498-.33.796v.41c0 .409-.11.809-.32 1.158l-1.315 2.191a2.11 2.11 0 0 1-1.81 1.025a1.055 1.055 0 0 1-1.055-1.055v-1.172c0-.92-.56-1.747-1.414-2.089l-.655-.261a2.25 2.25 0 0 1-1.383-2.46l.007-.042a2.25 2.25 0 0 1 .29-.787l.09-.15a2.25 2.25 0 0 1 2.37-1.048l1.178.236a1.125 1.125 0 0 0 1.302-.795l.208-.73a1.125 1.125 0 0 0-.578-1.315l-.665-.332l-.091.091a2.25 2.25 0 0 1-1.591.659h-.18a.94.94 0 0 0-.662.274a.931.931 0 0 1-1.458-1.137l1.411-2.353a2.25 2.25 0 0 0 .286-.76m11.928 9.869A9 9 0 0 0 8.965 3.525m11.928 9.868A9 9 0 1 1 8.965 3.525"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeEuropeAfricaSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1a7 7 0 1 0 0 14A7 7 0 0 0 8 1M5.657 3.023a5.5 5.5 0 1 0 7.584 3.304l-.947-.63a.431.431 0 0 0-.544.053a.431.431 0 0 1-.544.054l-.467-.312a.475.475 0 0 0-.689.608l.226.453a2.119 2.119 0 0 1 0 1.894L10.1 8.8a.947.947 0 0 0-.1.424v.11a2 2 0 0 1-.4 1.2L8.8 11.6A1 1 0 0 1 7 11v-.382a1 1 0 0 0-.553-.894l-.422-.212A1.854 1.854 0 0 1 6.855 6h.707a.438.438 0 1 0-.107-.864l-.835.209a1.129 1.129 0 0 1-1.305-1.553z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeEuropeAfricaSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75s9.75-4.365 9.75-9.75S17.385 2.25 12 2.25M8.547 4.505a8.25 8.25 0 1 0 11.672 8.214l-.46-.46a2.252 2.252 0 0 1-.422-.586l-1.08-2.16a.414.414 0 0 0-.663-.107a.827.827 0 0 1-.812.21l-1.273-.363a.89.89 0 0 0-.738 1.595l.587.39c.59.395.674 1.23.172 1.732l-.2.2c-.211.212-.33.498-.33.796v.41c0 .409-.11.809-.32 1.158l-1.315 2.191a2.11 2.11 0 0 1-1.81 1.025a1.055 1.055 0 0 1-1.055-1.055v-1.172c0-.92-.56-1.747-1.414-2.089l-.654-.261a2.25 2.25 0 0 1-1.384-2.46l.007-.042a2.25 2.25 0 0 1 .29-.787l.09-.15a2.25 2.25 0 0 1 2.37-1.048l1.178.236a1.125 1.125 0 0 0 1.302-.795l.208-.73a1.125 1.125 0 0 0-.578-1.315l-.665-.332l-.091.091a2.25 2.25 0 0 1-1.591.659h-.18a.94.94 0 0 0-.662.274a.931.931 0 0 1-1.458-1.137z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeEuropeAfricaTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M18 10a8 8 0 1 1-16 0a8 8 0 0 1 16 0m-1.503.204A6.5 6.5 0 1 1 7.95 3.83L6.927 5.62a1.453 1.453 0 0 0 1.91 2.02l.175-.087a.5.5 0 0 1 .224-.053h.146a.5.5 0 0 1 .447.724l-.028.055a.4.4 0 0 1-.357.221h-.502a2.26 2.26 0 0 0-1.88 1.006l-.044.066a2.099 2.099 0 0 0 1.085 3.156a.58.58 0 0 1 .397.547v1.05a1.175 1.175 0 0 0 2.093.734l1.611-2.014c.192-.24.296-.536.296-.842c0-.316.128-.624.353-.85a1.363 1.363 0 0 0 .173-1.716l-.464-.696a.369.369 0 0 1 .527-.499l.343.257a1.04 1.04 0 0 0 1.091.098a.586.586 0 0 1 .677.11z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandRaised(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.05 4.575a1.575 1.575 0 1 0-3.15 0v3m3.15-3v-1.5a1.575 1.575 0 0 1 3.15 0v1.5m-3.15 0l.075 5.925m3.075.75V4.575m0 0a1.575 1.575 0 0 1 3.15 0V15M6.9 7.575a1.575 1.575 0 1 0-3.15 0v8.175a6.75 6.75 0 0 0 6.75 6.75h2.018a5.25 5.25 0 0 0 3.712-1.538l1.732-1.732a5.25 5.25 0 0 0 1.538-3.712l.003-2.024a.668.668 0 0 1 .198-.471a1.575 1.575 0 1 0-2.228-2.228a3.818 3.818 0 0 0-1.12 2.687M6.9 7.575V12m6.27 4.318A4.49 4.49 0 0 1 16.35 15m.002 0h-.002"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandRaisedSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 1a.75.75 0 0 0-.75.75V6.5a.5.5 0 0 1-1 0V2.75a.75.75 0 0 0-1.5 0V7.5a.5.5 0 0 1-1 0V4.75a.75.75 0 0 0-1.5 0v4.5a5.75 5.75 0 0 0 11.5 0v-2.5a.75.75 0 0 0-1.5 0V9.5a.5.5 0 0 1-1 0V2.75a.75.75 0 0 0-1.5 0V6.5a.5.5 0 0 1-1 0V1.75A.75.75 0 0 0 8.5 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandRaisedSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 1.875a1.125 1.125 0 0 1 2.25 0v8.219c.517.162 1.02.382 1.5.659V3.375a1.125 1.125 0 0 1 2.25 0v10.937a4.505 4.505 0 0 0-3.25 2.373a8.963 8.963 0 0 1 4-.935A.75.75 0 0 0 18 15v-2.266a3.368 3.368 0 0 1 .988-2.37a1.125 1.125 0 0 1 1.592 1.59a1.118 1.118 0 0 0-.33.79v3.006h-.005a6 6 0 0 1-1.752 4.007l-1.736 1.736a6 6 0 0 1-4.242 1.757H10.5a7.5 7.5 0 0 1-7.5-7.5V6.375a1.125 1.125 0 0 1 2.25 0v5.519c.46-.452.965-.832 1.5-1.141V3.375a1.125 1.125 0 0 1 2.25 0v6.526c.495-.1.997-.151 1.5-.151z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandRaisedTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 2a1 1 0 1 0-2 0v6.5a.5.5 0 0 1-1 0V3a1 1 0 1 0-2 0v5.5a.5.5 0 0 1-1 0V5a1 1 0 1 0-2 0v7a7 7 0 1 0 14 0V8a1 1 0 1 0-2 0v3.5a.5.5 0 0 1-1 0V3a1 1 0 1 0-2 0v5.5a.5.5 0 0 1-1 0V8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandThumbDown(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.498 15.25H4.372c-1.026 0-1.945-.694-2.054-1.715a12.137 12.137 0 0 1-.068-1.285c0-2.848.992-5.464 2.649-7.521C5.287 4.247 5.886 4 6.504 4h4.016a4.5 4.5 0 0 1 1.423.23l3.114 1.04a4.5 4.5 0 0 0 1.423.23h1.294M7.499 15.25c.618 0 .991.724.725 1.282A7.471 7.471 0 0 0 7.5 19.75A2.25 2.25 0 0 0 9.75 22a.75.75 0 0 0 .75-.75v-.633c0-.573.11-1.14.322-1.672c.304-.76.93-1.33 1.653-1.715a9.04 9.04 0 0 0 2.86-2.4c.498-.634 1.226-1.08 2.032-1.08h.384m-10.253 1.5H9.7m8.075-9.75c.01.05.027.1.05.148a8.95 8.95 0 0 1 .925 3.977a8.95 8.95 0 0 1-.999 4.125m.023-8.25c-.076-.365.183-.75.575-.75h.908c.889 0 1.713.518 1.972 1.368c.339 1.11.521 2.287.521 3.507c0 1.553-.295 3.036-.831 4.398c-.306.774-1.086 1.227-1.918 1.227h-1.053c-.472 0-.745-.556-.5-.96a8.95 8.95 0 0 0 .303-.54"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandThumbDownSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.325 3H12v5c-.663 0-1.219.466-1.557 1.037a4.02 4.02 0 0 1-1.357 1.377c-.478.292-.907.706-.989 1.26v.005a9.031 9.031 0 0 0 0 2.642c.028.194-.048.394-.224.479A2 2 0 0 1 5 13c0-.812.08-1.605.234-2.371a.521.521 0 0 0-.5-.629H3c-1.104 0-2.01-.898-1.9-1.997a19.827 19.827 0 0 1 1.08-4.788C2.45 2.469 3.178 2 3.973 2h2.703a2 2 0 0 1 .632.103l2.384.794a2 2 0 0 0 .633.103M14 2a1 1 0 0 0-1 1v6a1 1 0 1 0 2 0V3a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandThumbDownSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.73 5.5h1.035A7.465 7.465 0 0 1 18 9.625a7.465 7.465 0 0 1-1.235 4.125h-.148c-.806 0-1.533.446-2.031 1.08a9.04 9.04 0 0 1-2.861 2.4c-.723.384-1.35.956-1.653 1.715a4.499 4.499 0 0 0-.322 1.672v.633A.75.75 0 0 1 9 22a2.25 2.25 0 0 1-2.25-2.25c0-1.152.26-2.243.723-3.218c.266-.558-.107-1.282-.725-1.282H3.622c-1.026 0-1.945-.694-2.054-1.715A12.137 12.137 0 0 1 1.5 12.25c0-2.848.992-5.464 2.649-7.521C4.537 4.247 5.136 4 5.754 4H9.77a4.5 4.5 0 0 1 1.423.23l3.114 1.04a4.5 4.5 0 0 0 1.423.23m5.939 8.523c.536-1.362.831-2.845.831-4.398c0-1.22-.182-2.398-.52-3.507c-.26-.85-1.084-1.368-1.973-1.368H19.1c-.445 0-.72.498-.523.898c.591 1.2.924 2.55.924 3.977a8.958 8.958 0 0 1-1.302 4.666c-.245.403.028.959.5.959h1.053c.832 0 1.612-.453 1.918-1.227"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandThumbDownTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.905 12.75a1.25 1.25 0 1 1-2.5 0v-7.5a1.25 1.25 0 0 1 2.5 0zm-10 4.25v1.3c0 .268-.14.526-.395.607A2 2 0 0 1 5.905 17c0-.995.182-1.948.514-2.826c.204-.54-.166-1.174-.744-1.174h-2.52c-1.243 0-2.261-1.01-2.146-2.247c.193-2.08.651-4.082 1.341-5.974C2.752 3.678 3.833 3 5.005 3h3.192a3 3 0 0 1 1.341.317l2.734 1.366A3 3 0 0 0 13.613 5h1.292v7h-.963c-.685 0-1.258.482-1.612 1.068a4.01 4.01 0 0 1-2.166 1.73c-.432.143-.853.386-1.011.814c-.16.432-.248.9-.248 1.388"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandThumbUp(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.633 10.25c.806 0 1.533-.446 2.031-1.08a9.041 9.041 0 0 1 2.861-2.4c.723-.384 1.35-.956 1.653-1.715a4.498 4.498 0 0 0 .322-1.672V2.75a.75.75 0 0 1 .75-.75a2.25 2.25 0 0 1 2.25 2.25c0 1.152-.26 2.243-.723 3.218c-.266.558.107 1.282.725 1.282m0 0h3.126c1.026 0 1.945.694 2.054 1.715c.045.422.068.85.068 1.285a11.95 11.95 0 0 1-2.649 7.521c-.388.482-.987.729-1.605.729H13.48a4.53 4.53 0 0 1-1.423-.23l-3.114-1.04a4.501 4.501 0 0 0-1.423-.23H5.904m10.598-9.75H14.25M5.904 18.5c.083.205.173.405.27.602c.197.4-.078.898-.523.898h-.908c-.889 0-1.713-.518-1.972-1.368a12 12 0 0 1-.521-3.507c0-1.553.295-3.036.831-4.398C3.387 9.953 4.167 9.5 5 9.5h1.053c.472 0 .745.556.5.96a8.958 8.958 0 0 0-1.302 4.665a8.97 8.97 0 0 0 .654 3.375"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandThumbUpSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.09 15a1 1 0 0 0 1-1V8a1 1 0 1 0-2 0v6a1 1 0 0 0 1 1m3.675-2H4.09V8c.663 0 1.218-.466 1.556-1.037a4.02 4.02 0 0 1 1.358-1.377c.478-.292.907-.706.989-1.26V4.32a9.03 9.03 0 0 0 0-2.642c-.028-.194.048-.394.224-.479A2 2 0 0 1 11.09 3c0 .812-.08 1.605-.235 2.371a.521.521 0 0 0 .502.629h1.733c1.104 0 2.01.898 1.901 1.997a19.831 19.831 0 0 1-1.081 4.788c-.27.747-.998 1.215-1.793 1.215H9.414c-.215 0-.428-.035-.632-.103l-2.384-.794A2.002 2.002 0 0 0 5.765 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandThumbUpSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.493 18.5c-.425 0-.82-.236-.975-.632A7.48 7.48 0 0 1 6 15.125a7.47 7.47 0 0 1 1.602-4.634c.151-.192.373-.309.6-.397c.473-.183.89-.514 1.212-.924a9.042 9.042 0 0 1 2.861-2.4c.723-.384 1.35-.956 1.653-1.715a4.498 4.498 0 0 0 .322-1.672V2.75A.75.75 0 0 1 15 2a2.25 2.25 0 0 1 2.25 2.25c0 1.152-.26 2.243-.723 3.218c-.266.558.107 1.282.725 1.282h3.126c1.026 0 1.945.694 2.054 1.715c.045.422.068.85.068 1.285a11.95 11.95 0 0 1-2.649 7.521c-.388.482-.987.729-1.605.729H14.23a4.53 4.53 0 0 1-1.423-.23l-3.114-1.04a4.501 4.501 0 0 0-1.423-.23zm-5.162-7.773a11.969 11.969 0 0 0-.831 4.398a12 12 0 0 0 .52 3.507C2.28 19.482 3.105 20 3.994 20H4.9c.445 0 .72-.498.523-.898a8.963 8.963 0 0 1-.924-3.977c0-1.708.476-3.305 1.302-4.666c.245-.403-.028-.959-.5-.959H4.25c-.832 0-1.612.453-1.918 1.227"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandThumbUpTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 8.25a1.25 1.25 0 1 1 2.5 0v7.5a1.25 1.25 0 1 1-2.5 0zM11 3V1.7c0-.268.14-.526.395-.607A2 2 0 0 1 14 3c0 .995-.182 1.948-.514 2.826c-.204.54.166 1.174.744 1.174h2.52c1.243 0 2.261 1.01 2.146 2.247a23.864 23.864 0 0 1-1.341 5.974C17.153 16.323 16.072 17 14.9 17h-3.192a3 3 0 0 1-1.341-.317l-2.734-1.366A3 3 0 0 0 6.292 15H5V8h.963c.685 0 1.258-.483 1.612-1.068a4.011 4.011 0 0 1 2.166-1.73c.432-.143.853-.386 1.011-.814c.16-.432.248-.9.248-1.388"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hashtag(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 8.25h15m-16.5 7.5h15m-1.8-13.5l-3.9 19.5m-2.1-19.5l-3.9 19.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HashtagSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.487 2.89a.75.75 0 1 0-1.474-.28l-.455 2.388H3.61a.75.75 0 0 0 0 1.5h1.663l-.571 2.998H2.75a.75.75 0 0 0 0 1.5h1.666l-.403 2.114a.75.75 0 0 0 1.474.28l.456-2.394h2.973l-.403 2.114a.75.75 0 0 0 1.474.28l.456-2.394h1.947a.75.75 0 0 0 0-1.5h-1.661l.57-2.998h1.95a.75.75 0 0 0 0-1.5h-1.664l.402-2.108a.75.75 0 0 0-1.474-.28l-.455 2.388H7.085zM6.8 6.498l-.571 2.998h2.973l.57-2.998z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HashtagSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.097 1.515a.75.75 0 0 1 .589.882L10.666 7.5h4.47l1.079-5.397a.75.75 0 1 1 1.47.294L16.665 7.5h3.585a.75.75 0 0 1 0 1.5h-3.885l-1.2 6h3.585a.75.75 0 0 1 0 1.5h-3.885l-1.08 5.397a.75.75 0 1 1-1.47-.294l1.02-5.103h-4.47l-1.08 5.397a.75.75 0 1 1-1.47-.294l1.02-5.103H3.75a.75.75 0 0 1 0-1.5h3.885l1.2-6H5.25a.75.75 0 0 1 0-1.5h3.885l1.08-5.397a.75.75 0 0 1 .882-.588M10.365 9l-1.2 6h4.47l1.2-6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HashtagTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.493 2.852a.75.75 0 0 0-1.486-.204L7.545 6H4.198a.75.75 0 0 0 0 1.5h3.14l-.69 5H3.302a.75.75 0 0 0 0 1.5h3.14l-.435 3.148a.75.75 0 0 0 1.486.204L7.955 14h2.986l-.434 3.148a.75.75 0 0 0 1.486.204L12.456 14h3.346a.75.75 0 0 0 0-1.5h-3.14l.69-5h3.346a.75.75 0 0 0 0-1.5h-3.14l.435-3.148a.75.75 0 0 0-1.486-.204L12.045 6H9.059zM8.852 7.5l-.69 5h2.986l.69-5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 8.25c0-2.485-2.099-4.5-4.687-4.5c-1.936 0-3.598 1.126-4.313 2.733c-.715-1.607-2.377-2.733-4.312-2.733C5.098 3.75 3 5.765 3 8.25c0 7.22 9 12 9 12s9-4.78 9-12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 6.342a3.375 3.375 0 0 1 6-2.088a3.375 3.375 0 0 1 5.997 2.26c-.063 2.134-1.618 3.76-2.955 4.784a14.437 14.437 0 0 1-2.676 1.61c-.02.01-.038.017-.05.022l-.014.006l-.004.002h-.002a.75.75 0 0 1-.592.001h-.002l-.004-.003l-.015-.006a5.528 5.528 0 0 1-.232-.107a14.395 14.395 0 0 1-2.535-1.557C3.564 10.22 1.999 8.558 1.999 6.38z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.645 20.91l-.007-.003l-.022-.012a15.247 15.247 0 0 1-.383-.218a25.18 25.18 0 0 1-4.244-3.17C4.688 15.36 2.25 12.174 2.25 8.25C2.25 5.322 4.714 3 7.688 3A5.5 5.5 0 0 1 12 5.052A5.5 5.5 0 0 1 16.313 3c2.973 0 5.437 2.322 5.437 5.25c0 3.925-2.438 7.111-4.739 9.256a25.175 25.175 0 0 1-4.244 3.17a15.247 15.247 0 0 1-.383.219l-.022.012l-.007.004l-.003.001a.752.752 0 0 1-.704 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.653 16.915l-.005-.003l-.019-.01a20.759 20.759 0 0 1-1.162-.682a22.045 22.045 0 0 1-2.582-1.9C4.045 12.733 2 10.352 2 7.5a4.5 4.5 0 0 1 8-2.828A4.5 4.5 0 0 1 18 7.5c0 2.852-2.044 5.233-3.885 6.82a22.049 22.049 0 0 1-3.744 2.582l-.019.01l-.005.003h-.002a.739.739 0 0 1-.69.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m2.25 12l8.955-8.955a1.124 1.124 0 0 1 1.59 0L21.75 12M4.5 9.75v10.125c0 .621.504 1.125 1.125 1.125H9.75v-4.875c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125V21h4.125c.621 0 1.125-.504 1.125-1.125V9.75M8.25 21h8.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeModern(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.25 21v-4.875c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125V21m0 0h4.5V3.545M12.75 21h7.5V10.75M2.25 21h1.5m18 0h-18M2.25 9l4.5-1.636M18.75 3l-1.5.545m0 6.205l3 1m1.5.5l-1.5-.5M6.75 7.364V3h-3v18m3-13.636l10.5-3.819"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeModernSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.536 3.444a.75.75 0 0 0-.571-1.387L3.5 4.719V3.75a.75.75 0 0 0-1.5 0v1.586l-.535.22A.75.75 0 0 0 2 6.958V12.5h-.25a.75.75 0 0 0 0 1.5H4a1 1 0 0 0 1-1v-1a1 1 0 1 1 2 0v1a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1V3.664zm1.293 2.358a.75.75 0 0 0-.333.623V8.5c0 .027.001.053.004.08V13a1 1 0 0 0 1 1h.5a1 1 0 0 0 1-1V7.957a.75.75 0 0 0 .535-1.4l-2.004-.826a.75.75 0 0 0-.703.07"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeModernSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M19.006 3.705a.75.75 0 1 0-.512-1.41L6 6.838V3a.75.75 0 0 0-.75-.75h-1.5A.75.75 0 0 0 3 3v4.93l-1.006.365a.75.75 0 0 0 .512 1.41z"/><path fill-rule="evenodd" d="M3.019 11.115L18 5.666v3.421l4.006 1.457a.75.75 0 1 1-.512 1.41l-.494-.18v8.475h.75a.75.75 0 0 1 0 1.5H2.25a.75.75 0 0 1 0-1.5H3V11.12zM18 20.25v-9.565l1.5.545v9.02zm-9-6a.75.75 0 0 0-.75.75v4.5c0 .414.336.75.75.75h3a.75.75 0 0 0 .75-.75V15a.75.75 0 0 0-.75-.75z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeModernTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.916 2.404a.75.75 0 0 1-.32 1.011l-.596.31V17a1 1 0 0 1-1 1h-2.26a.75.75 0 0 1-.75-.75v-3.5a.75.75 0 0 0-.75-.75H6.75a.75.75 0 0 0-.75.75v3.5a.75.75 0 0 1-.75.75h-3.5a.75.75 0 0 1 0-1.5H2V9.957a.75.75 0 0 1-.596-1.372L2 8.275V5.75a.75.75 0 0 1 1.5 0v1.745l10.404-5.41a.75.75 0 0 1 1.012.319m.946 6.166a.75.75 0 0 1 .735-.025l1.999 1.04A.75.75 0 0 1 18 10.957V16.5h.25a.75.75 0 0 1 0 1.5h-2a.75.75 0 0 1-.75-.75V9.21a.75.75 0 0 1 .361-.64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.543 2.232a.75.75 0 0 0-1.085 0l-5.25 5.5A.75.75 0 0 0 2.75 9H4v4a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-1a1 1 0 1 1 2 0v1a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1V9h1.25a.75.75 0 0 0 .543-1.268z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M11.47 3.841a.75.75 0 0 1 1.06 0l8.69 8.69a.75.75 0 1 0 1.06-1.061l-8.689-8.69a2.25 2.25 0 0 0-3.182 0l-8.69 8.69a.75.75 0 1 0 1.061 1.06z"/><path d="m12 5.432l8.159 8.159c.03.03.06.058.091.086v6.198c0 1.035-.84 1.875-1.875 1.875H15a.75.75 0 0 1-.75-.75v-4.5a.75.75 0 0 0-.75-.75h-3a.75.75 0 0 0-.75.75V21a.75.75 0 0 1-.75.75H5.625a1.875 1.875 0 0 1-1.875-1.875v-6.198a2.29 2.29 0 0 0 .091-.086z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.293 2.293a1 1 0 0 1 1.414 0l7 7A1 1 0 0 1 17 11h-1v6a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1v-3a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v3a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-6H3a1 1 0 0 1-.707-1.707z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Identification(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 9h3.75M15 12h3.75M15 15h3.75M4.5 19.5h15a2.25 2.25 0 0 0 2.25-2.25V6.75A2.25 2.25 0 0 0 19.5 4.5h-15a2.25 2.25 0 0 0-2.25 2.25v10.5A2.25 2.25 0 0 0 4.5 19.5m6-10.125a1.875 1.875 0 1 1-3.75 0a1.875 1.875 0 0 1 3.75 0m1.294 6.336a6.721 6.721 0 0 1-3.17.789a6.721 6.721 0 0 1-3.168-.789a3.376 3.376 0 0 1 6.338 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IdentificationSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 3a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2zm2.5 5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3M10 5.75a.75.75 0 0 1 .75-.75h1.5a.75.75 0 0 1 0 1.5h-1.5a.75.75 0 0 1-.75-.75m.75 3.75a.75.75 0 0 0 0 1.5h1.5a.75.75 0 0 0 0-1.5zM10 8a.75.75 0 0 1 .75-.75h1.5a.75.75 0 0 1 0 1.5h-1.5A.75.75 0 0 1 10 8m-2.378 3c.346 0 .583-.343.395-.633A2.998 2.998 0 0 0 5.5 9a2.998 2.998 0 0 0-2.517 1.367c-.188.29.05.633.395.633z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IdentificationSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 3.75a3 3 0 0 0-3 3v10.5a3 3 0 0 0 3 3h15a3 3 0 0 0 3-3V6.75a3 3 0 0 0-3-3zm4.125 3a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5m-3.873 8.703a4.126 4.126 0 0 1 7.746 0a.75.75 0 0 1-.351.92a7.47 7.47 0 0 1-3.522.877a7.47 7.47 0 0 1-3.522-.877a.75.75 0 0 1-.351-.92M15 8.25a.75.75 0 0 0 0 1.5h3.75a.75.75 0 0 0 0-1.5zM14.25 12a.75.75 0 0 1 .75-.75h3.75a.75.75 0 0 1 0 1.5H15a.75.75 0 0 1-.75-.75m.75 2.25a.75.75 0 0 0 0 1.5h3.75a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IdentificationTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3zm4 1.5a2 2 0 1 1 4 0a2 2 0 0 1-4 0m2 3a4 4 0 0 0-3.665 2.395a.75.75 0 0 0 .416 1A8.98 8.98 0 0 0 7 14.5a8.98 8.98 0 0 0 3.249-.604a.75.75 0 0 0 .416-1.001A4.001 4.001 0 0 0 7 10.5m5-3.75a.75.75 0 0 1 .75-.75h2.5a.75.75 0 0 1 0 1.5h-2.5a.75.75 0 0 1-.75-.75m0 6.5a.75.75 0 0 1 .75-.75h2.5a.75.75 0 0 1 0 1.5h-2.5a.75.75 0 0 1-.75-.75m.75-4a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inbox(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.25 13.5h3.86a2.25 2.25 0 0 1 2.012 1.244l.256.512a2.25 2.25 0 0 0 2.013 1.244h3.218a2.25 2.25 0 0 0 2.013-1.244l.256-.512a2.25 2.25 0 0 1 2.013-1.244h3.859m-19.5.338V18a2.25 2.25 0 0 0 2.25 2.25h15A2.25 2.25 0 0 0 21.75 18v-4.162c0-.224-.034-.447-.1-.661l-2.41-7.839a2.25 2.25 0 0 0-2.15-1.588H6.911a2.25 2.25 0 0 0-2.15 1.588L2.35 13.177a2.25 2.25 0 0 0-.1.661"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxArrowDown(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 3.75H6.912a2.25 2.25 0 0 0-2.15 1.588L2.35 13.177a2.25 2.25 0 0 0-.1.661V18a2.25 2.25 0 0 0 2.25 2.25h15A2.25 2.25 0 0 0 21.75 18v-4.162c0-.224-.034-.447-.1-.661l-2.41-7.839a2.25 2.25 0 0 0-2.15-1.588H15M2.25 13.5h3.86a2.25 2.25 0 0 1 2.012 1.244l.256.512a2.25 2.25 0 0 0 2.013 1.244h3.218a2.25 2.25 0 0 0 2.013-1.244l.256-.512a2.25 2.25 0 0 1 2.013-1.244h3.859M12 3v8.25m0 0l-3-3m3 3l3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxArrowDownSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.75 2.75a.75.75 0 0 0-1.5 0v3.69l-.72-.72a.75.75 0 0 0-1.06 1.06l2 2a.75.75 0 0 0 1.06 0l2-2a.75.75 0 1 0-1.06-1.06l-.72.72z"/><path fill="currentColor" d="M4.784 4.5a.75.75 0 0 0-.701.483L2.553 9h2.412a1 1 0 0 1 .832.445l.406.61a1 1 0 0 0 .832.445h1.93a1 1 0 0 0 .832-.445l.406-.61A1 1 0 0 1 11.035 9h2.412l-1.53-4.017a.75.75 0 0 0-.7-.483h-.467a.75.75 0 0 1 0-1.5h.466c.934 0 1.77.577 2.103 1.449l1.534 4.026c.097.256.147.527.147.801v1.474A2.25 2.25 0 0 1 12.75 13h-9.5A2.25 2.25 0 0 1 1 10.75V9.276c0-.274.05-.545.147-.801l1.534-4.026A2.25 2.25 0 0 1 4.784 3h.466a.75.75 0 0 1 0 1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxArrowDownSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M5.478 5.559A1.5 1.5 0 0 1 6.912 4.5H9A.75.75 0 0 0 9 3H6.912a3 3 0 0 0-2.868 2.118l-2.411 7.838a3 3 0 0 0-.133.882V18a3 3 0 0 0 3 3h15a3 3 0 0 0 3-3v-4.162c0-.299-.045-.596-.133-.882l-2.411-7.838A3 3 0 0 0 17.088 3H15a.75.75 0 0 0 0 1.5h2.088a1.5 1.5 0 0 1 1.434 1.059l2.213 7.191H17.89a3 3 0 0 0-2.684 1.658l-.256.513a1.5 1.5 0 0 1-1.342.829H10.39a1.5 1.5 0 0 1-1.342-.83l-.256-.512a3 3 0 0 0-2.684-1.658H3.265z"/><path d="M12 2.25a.75.75 0 0 1 .75.75v6.44l1.72-1.72a.75.75 0 1 1 1.06 1.06l-3 3a.75.75 0 0 1-1.06 0l-3-3a.75.75 0 0 1 1.06-1.06l1.72 1.72V3a.75.75 0 0 1 .75-.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxArrowDownTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10 2a.75.75 0 0 1 .75.75v5.59l1.95-2.1a.75.75 0 1 1 1.1 1.02l-3.25 3.5a.75.75 0 0 1-1.1 0L6.2 7.26a.75.75 0 1 1 1.1-1.02l1.95 2.1V2.75A.75.75 0 0 1 10 2"/><path d="M5.273 4.5a1.25 1.25 0 0 0-1.205.918l-1.523 5.52c-.006.02-.01.041-.015.062H6a1 1 0 0 1 .894.553l.448.894a1 1 0 0 0 .894.553h3.438a1 1 0 0 0 .86-.49l.606-1.02A1 1 0 0 1 14 11h3.47a1.318 1.318 0 0 0-.015-.062l-1.523-5.52a1.25 1.25 0 0 0-1.205-.918h-.977a.75.75 0 0 1 0-1.5h.977a2.75 2.75 0 0 1 2.651 2.019l1.523 5.52c.066.239.099.485.099.732V15a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2v-3.73c0-.246.033-.492.099-.73l1.523-5.521A2.75 2.75 0 0 1 5.273 3h.977a.75.75 0 0 1 0 1.5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.784 3A2.25 2.25 0 0 0 2.68 4.449L1.147 8.475A2.25 2.25 0 0 0 1 9.276v1.474A2.25 2.25 0 0 0 3.25 13h9.5A2.25 2.25 0 0 0 15 10.75V9.276c0-.274-.05-.545-.147-.801l-1.534-4.026A2.25 2.25 0 0 0 11.216 3zm-.701 1.983a.75.75 0 0 1 .7-.483h6.433a.75.75 0 0 1 .701.483L13.447 9h-2.412a1 1 0 0 0-.832.445l-.406.61a1 1 0 0 1-.832.445h-1.93a1 1 0 0 1-.832-.445l-.406-.61A1 1 0 0 0 4.965 9H2.553z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.912 3a3 3 0 0 0-2.868 2.118l-2.411 7.838a3 3 0 0 0-.133.882V18a3 3 0 0 0 3 3h15a3 3 0 0 0 3-3v-4.162c0-.299-.045-.596-.133-.882l-2.411-7.838A3 3 0 0 0 17.088 3zm13.823 9.75l-2.213-7.191A1.5 1.5 0 0 0 17.088 4.5H6.912a1.5 1.5 0 0 0-1.434 1.059L3.265 12.75H6.11a3 3 0 0 1 2.684 1.658l.256.513a1.5 1.5 0 0 0 1.342.829h3.218a1.5 1.5 0 0 0 1.342-.83l.256-.512a3 3 0 0 1 2.684-1.658z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxStack(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m7.875 14.25l1.214 1.942a2.25 2.25 0 0 0 1.908 1.058h2.006c.776 0 1.497-.4 1.908-1.058l1.214-1.942M2.41 9h4.636a2.25 2.25 0 0 1 1.872 1.002l.164.246a2.25 2.25 0 0 0 1.872 1.002h2.092a2.25 2.25 0 0 0 1.872-1.002l.164-.246A2.25 2.25 0 0 1 16.954 9h4.636M2.41 9a2.25 2.25 0 0 0-.16.832V12a2.25 2.25 0 0 0 2.25 2.25h15A2.25 2.25 0 0 0 21.75 12V9.832c0-.287-.055-.57-.16-.832M2.41 9a2.25 2.25 0 0 1 .382-.632l3.285-3.832a2.25 2.25 0 0 1 1.708-.786h8.43c.657 0 1.281.287 1.709.786l3.284 3.832c.163.19.291.404.382.632M4.5 20.25h15A2.25 2.25 0 0 0 21.75 18v-2.625c0-.621-.504-1.125-1.125-1.125H3.375c-.621 0-1.125.504-1.125 1.125V18a2.25 2.25 0 0 0 2.25 2.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxStackSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.742 2.755A2.25 2.25 0 0 1 4.424 2h7.152a2.25 2.25 0 0 1 1.682.755l1.174 1.32c.366.412.568.944.568 1.495v.68a2.25 2.25 0 0 1-2.25 2.25h-9.5A2.25 2.25 0 0 1 1 6.25v-.68c0-.55.202-1.083.568-1.495zm1.682.745a.75.75 0 0 0-.561.252L2.753 5h2.212a1 1 0 0 1 .832.445l.406.61a1 1 0 0 0 .832.445h1.93a1 1 0 0 0 .832-.445l.406-.61A1 1 0 0 1 11.035 5h2.211l-1.109-1.248a.75.75 0 0 0-.56-.252z" clip-rule="evenodd"/><path fill="currentColor" d="M1 10.75a.75.75 0 0 1 .75-.75h3.215a1 1 0 0 1 .832.445l.406.61a1 1 0 0 0 .832.445h1.93a1 1 0 0 0 .832-.445l.406-.61a1 1 0 0 1 .832-.445h3.215a.75.75 0 0 1 .75.75v1A2.25 2.25 0 0 1 12.75 14h-9.5A2.25 2.25 0 0 1 1 11.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxStackSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M1.5 9.832v1.793c0 1.036.84 1.875 1.875 1.875h17.25c1.035 0 1.875-.84 1.875-1.875V9.832a3 3 0 0 0-.722-1.952l-3.285-3.832A3 3 0 0 0 16.215 3h-8.43a3 3 0 0 0-2.278 1.048L2.222 7.88A3 3 0 0 0 1.5 9.832M7.785 4.5a1.5 1.5 0 0 0-1.139.524L3.881 8.25h3.165a3 3 0 0 1 2.496 1.336l.164.246a1.5 1.5 0 0 0 1.248.668h2.092a1.5 1.5 0 0 0 1.248-.668l.164-.246a3 3 0 0 1 2.496-1.336h3.165l-2.765-3.226a1.5 1.5 0 0 0-1.139-.524z" clip-rule="evenodd"/><path d="M2.813 15c-.725 0-1.313.588-1.313 1.313V18a3 3 0 0 0 3 3h15a3 3 0 0 0 3-3v-1.687c0-.725-.588-1.313-1.312-1.313h-4.234a3 3 0 0 0-2.496 1.336l-.164.246a1.5 1.5 0 0 1-1.248.668h-2.092a1.5 1.5 0 0 1-1.248-.668l-.164-.246A3 3 0 0 0 7.046 15z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxStackTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M1.045 6.954a2.75 2.75 0 0 1 .217-.678L2.53 3.58A2.75 2.75 0 0 1 5.019 2h9.962a2.75 2.75 0 0 1 2.488 1.58l1.27 2.696c.101.216.174.444.216.678A1 1 0 0 1 19 7.25v1.5a2.75 2.75 0 0 1-2.75 2.75H3.75A2.75 2.75 0 0 1 1 8.75v-1.5a1 1 0 0 1 .045-.296m2.843-2.736A1.25 1.25 0 0 1 5.02 3.5h9.962c.484 0 .925.28 1.13.718l.957 2.032H14a1 1 0 0 0-.86.49l-.606 1.02a1 1 0 0 1-.86.49H8.236a1 1 0 0 1-.894-.553l-.448-.894A1 1 0 0 0 6 6.25H2.932z" clip-rule="evenodd"/><path d="M1 14a1 1 0 0 1 1-1h4a1 1 0 0 1 .894.553l.448.894a1 1 0 0 0 .894.553h3.438a1 1 0 0 0 .86-.49l.606-1.02A1 1 0 0 1 14 13h4a1 1 0 0 1 1 1v2a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 11.27c0-.246.033-.492.099-.73l1.523-5.521A2.75 2.75 0 0 1 5.273 3h9.454a2.75 2.75 0 0 1 2.651 2.019l1.523 5.52c.066.239.099.485.099.732V15a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2zm3.068-5.852A1.25 1.25 0 0 1 5.273 4.5h9.454a1.25 1.25 0 0 1 1.205.918l1.523 5.52c.006.02.01.041.015.062H14a1 1 0 0 0-.86.49l-.606 1.02a1 1 0 0 1-.86.49H8.236a1 1 0 0 1-.894-.553l-.448-.894A1 1 0 0 0 6 11H2.53l.015-.062z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InformationCircle(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m11.25 11.25l.041-.02a.75.75 0 0 1 1.063.852l-.708 2.836a.75.75 0 0 0 1.063.853l.041-.021M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0m-9-3.75h.008v.008H12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InformationCircleSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M9 5a1 1 0 1 1-2 0a1 1 0 0 1 2 0M6.75 8a.75.75 0 0 0 0 1.5h.75v1.75a.75.75 0 0 0 1.5 0v-2.5A.75.75 0 0 0 8.25 8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InformationCircleSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 12c0-5.385 4.365-9.75 9.75-9.75s9.75 4.365 9.75 9.75s-4.365 9.75-9.75 9.75S2.25 17.385 2.25 12m8.706-1.442c1.146-.573 2.437.463 2.126 1.706l-.709 2.836l.042-.02a.75.75 0 0 1 .67 1.34l-.04.022c-1.147.573-2.438-.463-2.127-1.706l.71-2.836l-.042.02a.75.75 0 1 1-.671-1.34zM12 9a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InformationCircleTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M18 10a8 8 0 1 1-16 0a8 8 0 0 1 16 0m-7-4a1 1 0 1 1-2 0a1 1 0 0 1 2 0M9 9a.75.75 0 0 0 0 1.5h.253a.25.25 0 0 1 .244.304l-.459 2.066A1.75 1.75 0 0 0 10.747 15H11a.75.75 0 0 0 0-1.5h-.253a.25.25 0 0 1-.244-.304l.459-2.066A1.75 1.75 0 0 0 9.253 9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Italic(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.25 3.75h3.696m3.804 0h-3.804M5.25 20.25h3.804m3.696 0H9.054m0 0l5.892-16.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ItalicSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.5 3.75a.75.75 0 0 1 .75-.75h7.5a.75.75 0 0 1 0 1.5h-3.275l-5.357 15h2.632a.75.75 0 0 1 0 1.5h-7.5a.75.75 0 0 1 0-1.5h3.275l5.357-15H11.25a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ItalicTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.05 2.75A.75.75 0 0 1 8.8 2h7.45a.75.75 0 0 1 0 1.5h-3.132l-4.643 13h2.785a.75.75 0 0 1 0 1.5H3.75a.75.75 0 0 1 0-1.5h3.132l4.643-13H8.8a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Key(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.75 5.25a3 3 0 0 1 3 3m3 0a6 6 0 0 1-7.029 5.912c-.563-.097-1.159.026-1.563.43L10.5 17.25H8.25v2.25H6v2.25H2.25v-2.818c0-.597.237-1.17.659-1.591l6.499-6.499c.404-.404.527-1 .43-1.563A6 6 0 1 1 21.75 8.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeySixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14 6a4 4 0 0 1-4.899 3.899l-1.955 1.955a.5.5 0 0 1-.353.146H5v1.5a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1-.5-.5v-2.293a.5.5 0 0 1 .146-.353l3.955-3.955A4 4 0 1 1 14 6m-4-2a.75.75 0 0 0 0 1.5a.5.5 0 0 1 .5.5a.75.75 0 0 0 1.5 0a2 2 0 0 0-2-2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.75 1.5a6.75 6.75 0 0 0-6.651 7.906c.067.39-.032.717-.221.906l-6.5 6.499a3 3 0 0 0-.878 2.121v2.818c0 .414.336.75.75.75H6a.75.75 0 0 0 .75-.75v-1.5h1.5A.75.75 0 0 0 9 19.5V18h1.5a.75.75 0 0 0 .53-.22l2.658-2.658c.19-.189.517-.288.906-.22A6.75 6.75 0 1 0 15.75 1.5m0 3a.75.75 0 0 0 0 1.5A2.25 2.25 0 0 1 18 8.25a.75.75 0 0 0 1.5 0a3.75 3.75 0 0 0-3.75-3.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 7a5 5 0 1 1 3.61 4.804l-1.903 1.903A1 1 0 0 1 9 14H8v1a1 1 0 0 1-1 1H6v1a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1v-2a1 1 0 0 1 .293-.707L8.196 8.39A5.002 5.002 0 0 1 8 7m5-3a.75.75 0 0 0 0 1.5A1.5 1.5 0 0 1 14.5 7A.75.75 0 0 0 16 7a3 3 0 0 0-3-3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Language(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10.5 21l5.25-11.25L21 21m-9-3h7.5M3 5.621a48.474 48.474 0 0 1 6-.371m0 0c1.12 0 2.233.038 3.334.114M9 5.25V3m3.334 2.364C11.176 10.658 7.69 15.08 3 17.502m9.334-12.138A47.63 47.63 0 0 1 15 5.621m-4.589 8.495a18.023 18.023 0 0 1-3.827-5.802"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LanguageSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 5a.75.75 0 0 1 .688.452l3.25 7.5a.75.75 0 1 1-1.376.596L12.89 12H9.109l-.67 1.548a.75.75 0 1 1-1.377-.596l3.25-7.5A.75.75 0 0 1 11 5m-1.24 5.5h2.48L11 7.636zM5 1a.75.75 0 0 1 .75.75v1.261a25.27 25.27 0 0 1 2.598.211a.75.75 0 1 1-.2 1.487c-.22-.03-.44-.056-.662-.08A12.939 12.939 0 0 1 5.92 8.058c.237.304.488.595.752.873a.75.75 0 0 1-1.086 1.035A13.075 13.075 0 0 1 5 9.307a13.068 13.068 0 0 1-2.841 2.546a.75.75 0 0 1-.827-1.251A11.566 11.566 0 0 0 4.08 8.057a12.991 12.991 0 0 1-.554-.938a.75.75 0 1 1 1.323-.707c.049.09.099.181.15.271c.388-.68.708-1.405.952-2.164a23.941 23.941 0 0 0-4.1.19a.75.75 0 0 1-.2-1.487c.853-.114 1.72-.185 2.598-.211V1.75A.75.75 0 0 1 5 1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LanguageSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9 2.25a.75.75 0 0 1 .75.75v1.506a49.384 49.384 0 0 1 5.343.371a.75.75 0 1 1-.186 1.489a46.7 46.7 0 0 0-1.99-.206a18.67 18.67 0 0 1-2.97 6.323c.318.384.65.753 1 1.108a.75.75 0 0 1-1.07 1.05A18.902 18.902 0 0 1 9 13.688a18.823 18.823 0 0 1-5.656 4.482a.75.75 0 0 1-.688-1.333a17.323 17.323 0 0 0 5.396-4.353A18.72 18.72 0 0 1 5.89 8.598a.75.75 0 0 1 1.388-.568A17.21 17.21 0 0 0 9 11.224a17.168 17.168 0 0 0 2.391-5.165a48.04 48.04 0 0 0-8.298.307a.75.75 0 0 1-.186-1.489a49.159 49.159 0 0 1 5.343-.371V3A.75.75 0 0 1 9 2.25M15.75 9a.75.75 0 0 1 .68.433l5.25 11.25a.75.75 0 1 1-1.36.634l-1.198-2.567h-6.744l-1.198 2.567a.75.75 0 0 1-1.36-.634l5.25-11.25A.75.75 0 0 1 15.75 9m-2.672 8.25h5.344l-2.672-5.726z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LanguageTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M7.75 2.75a.75.75 0 0 0-1.5 0v1.258a32.987 32.987 0 0 0-3.599.278a.75.75 0 1 0 .198 1.487A31.545 31.545 0 0 1 8.7 5.545A19.381 19.381 0 0 1 7 9.56a19.418 19.418 0 0 1-1.002-2.05a.75.75 0 0 0-1.384.577a20.935 20.935 0 0 0 1.492 2.91a19.613 19.613 0 0 1-3.828 4.154a.75.75 0 1 0 .945 1.164A21.116 21.116 0 0 0 7 12.331c.095.132.192.262.29.391a.75.75 0 0 0 1.194-.91a18.97 18.97 0 0 1-.59-.815a20.888 20.888 0 0 0 2.333-5.332c.31.031.618.068.924.108a.75.75 0 0 0 .198-1.487a32.832 32.832 0 0 0-3.599-.278z"/><path fill-rule="evenodd" d="M13 8a.75.75 0 0 1 .671.415l4.25 8.5a.75.75 0 1 1-1.342.67L15.787 16h-5.573l-.793 1.585a.75.75 0 1 1-1.342-.67l4.25-8.5A.75.75 0 0 1 13 8m2.037 6.5L13 10.427L10.964 14.5z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lifebuoy(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.712 4.33a9.027 9.027 0 0 1 1.652 1.306c.51.51.944 1.064 1.306 1.652M16.712 4.33l-3.448 4.138m3.448-4.138a9.014 9.014 0 0 0-9.424 0M19.67 7.288l-4.138 3.448m4.138-3.448a9.014 9.014 0 0 1 0 9.424m-4.138-5.976a3.736 3.736 0 0 0-.88-1.388a3.737 3.737 0 0 0-1.388-.88m2.268 2.268a3.765 3.765 0 0 1 0 2.528m-2.268-4.796a3.765 3.765 0 0 0-2.528 0m4.796 4.796a3.754 3.754 0 0 1-.88 1.388a3.736 3.736 0 0 1-1.388.88m2.268-2.268l4.138 3.448m0 0a9.027 9.027 0 0 1-1.306 1.652c-.51.51-1.064.944-1.652 1.306m0 0l-3.448-4.138m3.448 4.138a9.014 9.014 0 0 1-9.424 0m5.976-4.138a3.765 3.765 0 0 1-2.528 0m0 0a3.736 3.736 0 0 1-1.388-.88a3.737 3.737 0 0 1-.88-1.388m2.268 2.268L7.288 19.67m0 0a9.024 9.024 0 0 1-1.652-1.306a9.027 9.027 0 0 1-1.306-1.652m0 0l4.138-3.448M4.33 16.712a9.014 9.014 0 0 1 0-9.424m4.138 5.976a3.765 3.765 0 0 1 0-2.528m0 0c.181-.506.475-.982.88-1.388a3.736 3.736 0 0 1 1.388-.88m-2.268 2.268L4.33 7.288m6.406 1.18L7.288 4.33m0 0a9.024 9.024 0 0 0-1.652 1.306A9.025 9.025 0 0 0 4.33 7.288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LifebuoySixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.95 3.05a7 7 0 1 0-9.9 9.9a7 7 0 0 0 9.9-9.9m-7.262-.042a5.516 5.516 0 0 1 4.624 0L8.698 5.082a3.016 3.016 0 0 0-1.396 0zm-2.68 2.68a5.516 5.516 0 0 0 0 4.624l2.074-1.614a3.015 3.015 0 0 1 0-1.396zm2.68 7.304l1.614-2.074c.458.11.938.11 1.396 0l1.614 2.074a5.516 5.516 0 0 1-4.624 0m7.304-2.68a5.516 5.516 0 0 0 0-4.624l-2.074 1.614c.11.458.11.938 0 1.396zM6.94 6.939a1.5 1.5 0 1 1 2.122 2.122A1.5 1.5 0 0 1 6.94 6.939" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LifebuoySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M19.449 8.448L16.388 11a4.52 4.52 0 0 1 0 2.002l3.061 2.55a8.275 8.275 0 0 0 0-7.103M15.552 19.45L13 16.388a4.52 4.52 0 0 1-2.002 0l-2.55 3.061a8.275 8.275 0 0 0 7.103 0M4.55 15.552L7.612 13a4.52 4.52 0 0 1 0-2.002L4.551 8.45a8.275 8.275 0 0 0 0 7.103M8.448 4.55L11 7.612a4.52 4.52 0 0 1 2.002 0l2.55-3.061a8.275 8.275 0 0 0-7.103 0m8.657-.86a9.776 9.776 0 0 1 1.79 1.415a9.776 9.776 0 0 1 1.414 1.788a9.764 9.764 0 0 1 0 10.212a9.777 9.777 0 0 1-1.415 1.788a9.777 9.777 0 0 1-1.788 1.415a9.764 9.764 0 0 1-10.212 0a9.776 9.776 0 0 1-1.788-1.415a9.776 9.776 0 0 1-1.415-1.788a9.764 9.764 0 0 1 0-10.212a9.774 9.774 0 0 1 1.415-1.788A9.774 9.774 0 0 1 6.894 3.69a9.764 9.764 0 0 1 10.212 0m-2.984 6.19a2.985 2.985 0 0 0-1.11-.704a3.015 3.015 0 0 0-2.022 0a2.985 2.985 0 0 0-1.11.704c-.326.325-.56.705-.704 1.11a3.015 3.015 0 0 0 0 2.022c.144.405.378.785.704 1.11c.325.326.705.56 1.11.704c.652.233 1.37.233 2.022 0a2.985 2.985 0 0 0 1.11-.704c.326-.325.56-.705.704-1.11a3.016 3.016 0 0 0 0-2.022a2.985 2.985 0 0 0-.704-1.11" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LifebuoyTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7.171 4.146l1.947 2.466a3.514 3.514 0 0 1 1.764 0l1.947-2.466a6.52 6.52 0 0 0-5.658 0m8.683 3.025l-2.466 1.947c.15.578.15 1.186 0 1.764l2.466 1.947a6.52 6.52 0 0 0 0-5.658m-3.025 8.683l-1.947-2.466c-.578.15-1.186.15-1.764 0l-1.947 2.466a6.52 6.52 0 0 0 5.658 0M4.146 12.83l2.466-1.947a3.514 3.514 0 0 1 0-1.764L4.146 7.171a6.52 6.52 0 0 0 0 5.658M5.63 3.297a8.01 8.01 0 0 1 8.738 0a8.031 8.031 0 0 1 2.334 2.334a8.01 8.01 0 0 1 0 8.738a8.033 8.033 0 0 1-2.334 2.334a8.01 8.01 0 0 1-8.738 0a8.032 8.032 0 0 1-2.334-2.334a8.01 8.01 0 0 1 0-8.738A8.03 8.03 0 0 1 5.63 3.297m5.198 4.882a2.008 2.008 0 0 0-2.243.407a1.994 1.994 0 0 0-.407 2.243a1.993 1.993 0 0 0 .992.992a2.008 2.008 0 0 0 2.243-.407c.176-.175.31-.374.407-.585a2.008 2.008 0 0 0-.407-2.243a1.993 1.993 0 0 0-.585-.407" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightBulb(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 18v-5.25m0 0a6.01 6.01 0 0 0 1.5-.189m-1.5.189a6.01 6.01 0 0 1-1.5-.189m3.75 7.478a12.06 12.06 0 0 1-4.5 0m3.75 2.383a14.406 14.406 0 0 1-3 0M14.25 18v-.192c0-.983.658-1.823 1.508-2.316a7.5 7.5 0 1 0-7.517 0c.85.493 1.509 1.333 1.509 2.316V18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightBulbSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.618 10.26c-.361.223-.618.598-.618 1.022c0 .226-.142.43-.36.49A6.006 6.006 0 0 1 8 12c-.569 0-1.12-.08-1.64-.227a.504.504 0 0 1-.36-.491c0-.424-.257-.799-.618-1.021a5 5 0 1 1 5.235 0m-3.75 3.154a.75.75 0 1 0-.225 1.483a9.065 9.065 0 0 0 2.716 0a.75.75 0 1 0-.225-1.483a7.563 7.563 0 0 1-2.266 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightBulbSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 .75a8.25 8.25 0 0 0-4.135 15.39c.686.398 1.115 1.008 1.134 1.623a.75.75 0 0 0 .577.706c.352.083.71.148 1.074.195c.323.041.6-.218.6-.544v-4.661a6.714 6.714 0 0 1-.937-.171a.75.75 0 1 1 .374-1.453a5.261 5.261 0 0 0 2.626 0a.75.75 0 1 1 .374 1.452a6.712 6.712 0 0 1-.937.172v4.66a.54.54 0 0 0 .6.545a10.58 10.58 0 0 0 1.074-.195a.75.75 0 0 0 .577-.706c.02-.615.448-1.225 1.134-1.623A8.25 8.25 0 0 0 12 .75"/><path fill-rule="evenodd" d="M9.013 19.9a.75.75 0 0 1 .877-.597a11.319 11.319 0 0 0 4.22 0a.75.75 0 1 1 .28 1.473a12.819 12.819 0 0 1-4.78 0a.75.75 0 0 1-.597-.876m.741 2.444a.75.75 0 0 1 .824-.668a13.682 13.682 0 0 0 2.844 0a.75.75 0 1 1 .156 1.492a15.156 15.156 0 0 1-3.156 0a.75.75 0 0 1-.668-.824" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightBulbTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 1a6 6 0 0 0-3.815 10.631C7.237 12.5 8 13.443 8 14.456v.644a.75.75 0 0 0 .572.729a6.016 6.016 0 0 0 2.856 0A.75.75 0 0 0 12 15.1v-.644c0-1.013.762-1.957 1.815-2.825A6 6 0 0 0 10 1M8.863 17.414a.75.75 0 0 0-.226 1.483a9.066 9.066 0 0 0 2.726 0a.75.75 0 0 0-.226-1.483a7.553 7.553 0 0 1-2.274 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.19 8.688a4.5 4.5 0 0 1 1.242 7.244l-4.5 4.5a4.5 4.5 0 0 1-6.364-6.364l1.757-1.757m13.35-.622l1.757-1.757a4.5 4.5 0 0 0-6.364-6.364l-4.5 4.5a4.5 4.5 0 0 0 1.242 7.244"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8.914 6.025a.75.75 0 0 1 1.06 0a3.5 3.5 0 0 1 0 4.95l-2 2a3.5 3.5 0 0 1-5.396-4.402a.75.75 0 0 1 1.251.827a2 2 0 0 0 3.085 2.514l2-2a2 2 0 0 0 0-2.828a.75.75 0 0 1 0-1.06"/><path d="M7.086 9.975a.75.75 0 0 1-1.06 0a3.5 3.5 0 0 1 0-4.95l2-2a3.5 3.5 0 0 1 5.396 4.402a.75.75 0 0 1-1.251-.827a2 2 0 0 0-3.085-2.514l-2 2a2 2 0 0 0 0 2.828a.75.75 0 0 1 0 1.06"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M19.902 4.098a3.75 3.75 0 0 0-5.304 0l-4.5 4.5a3.75 3.75 0 0 0 1.035 6.037a.75.75 0 0 1-.646 1.353a5.25 5.25 0 0 1-1.449-8.45l4.5-4.5a5.25 5.25 0 1 1 7.424 7.424l-1.757 1.757a.75.75 0 1 1-1.06-1.06l1.757-1.757a3.75 3.75 0 0 0 0-5.304m-7.389 4.267a.75.75 0 0 1 1-.353a5.25 5.25 0 0 1 1.449 8.45l-4.5 4.5a5.25 5.25 0 1 1-7.424-7.424l1.757-1.757a.75.75 0 1 1 1.06 1.06l-1.757 1.757a3.75 3.75 0 1 0 5.304 5.304l4.5-4.5a3.75 3.75 0 0 0-1.035-6.037a.75.75 0 0 1-.354-1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12.232 4.232a2.5 2.5 0 0 1 3.536 3.536l-1.225 1.224a.75.75 0 0 0 1.061 1.06l1.224-1.224a4 4 0 0 0-5.656-5.656l-3 3a4 4 0 0 0 .225 5.865a.75.75 0 0 0 .977-1.138a2.5 2.5 0 0 1-.142-3.667z"/><path d="M11.603 7.963a.75.75 0 0 0-.977 1.138a2.5 2.5 0 0 1 .142 3.667l-3 3a2.5 2.5 0 0 1-3.536-3.536l1.225-1.224a.75.75 0 0 0-1.061-1.06l-1.224 1.224a4 4 0 1 0 5.656 5.656l3-3a4 4 0 0 0-.225-5.865"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListBullet(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.25 6.75h12M8.25 12h12m-12 5.25h12M3.75 6.75h.007v.008H3.75zm.375 0a.375.375 0 1 1-.75 0a.375.375 0 0 1 .75 0M3.75 12h.007v.008H3.75zm.375 0a.375.375 0 1 1-.75 0a.375.375 0 0 1 .75 0m-.375 5.25h.007v.008H3.75zm.375 0a.375.375 0 1 1-.75 0a.375.375 0 0 1 .75 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListBulletSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 4.75a1 1 0 1 0 0-2a1 1 0 0 0 0 2M6.25 3a.75.75 0 0 0 0 1.5h7a.75.75 0 0 0 0-1.5zm0 4.25a.75.75 0 0 0 0 1.5h7a.75.75 0 0 0 0-1.5zm0 4.25a.75.75 0 0 0 0 1.5h7a.75.75 0 0 0 0-1.5zM4 12.25a1 1 0 1 1-2 0a1 1 0 0 1 2 0M3 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListBulletSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.625 6.75a1.125 1.125 0 1 1 2.25 0a1.125 1.125 0 0 1-2.25 0m4.875 0A.75.75 0 0 1 8.25 6h12a.75.75 0 0 1 0 1.5h-12a.75.75 0 0 1-.75-.75M2.625 12a1.125 1.125 0 1 1 2.25 0a1.125 1.125 0 0 1-2.25 0M7.5 12a.75.75 0 0 1 .75-.75h12a.75.75 0 0 1 0 1.5h-12A.75.75 0 0 1 7.5 12m-4.875 5.25a1.125 1.125 0 1 1 2.25 0a1.125 1.125 0 0 1-2.25 0m4.875 0a.75.75 0 0 1 .75-.75h12a.75.75 0 0 1 0 1.5h-12a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListBulletTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 4.75A.75.75 0 0 1 6.75 4h10.5a.75.75 0 0 1 0 1.5H6.75A.75.75 0 0 1 6 4.75M6 10a.75.75 0 0 1 .75-.75h10.5a.75.75 0 0 1 0 1.5H6.75A.75.75 0 0 1 6 10m0 5.25a.75.75 0 0 1 .75-.75h10.5a.75.75 0 0 1 0 1.5H6.75a.75.75 0 0 1-.75-.75M1.99 4.75a1 1 0 0 1 1-1H3a1 1 0 0 1 1 1v.01a1 1 0 0 1-1 1h-.01a1 1 0 0 1-1-1zm0 10.5a1 1 0 0 1 1-1H3a1 1 0 0 1 1 1v.01a1 1 0 0 1-1 1h-.01a1 1 0 0 1-1-1zm0-5.25a1 1 0 0 1 1-1H3a1 1 0 0 1 1 1v.01a1 1 0 0 1-1 1h-.01a1 1 0 0 1-1-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockClosed(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.5 10.5V6.75a4.5 4.5 0 1 0-9 0v3.75m-.75 11.25h10.5a2.25 2.25 0 0 0 2.25-2.25v-6.75a2.25 2.25 0 0 0-2.25-2.25H6.75a2.25 2.25 0 0 0-2.25 2.25v6.75a2.25 2.25 0 0 0 2.25 2.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockClosedSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1a3.5 3.5 0 0 0-3.5 3.5V7A1.5 1.5 0 0 0 3 8.5v5A1.5 1.5 0 0 0 4.5 15h7a1.5 1.5 0 0 0 1.5-1.5v-5A1.5 1.5 0 0 0 11.5 7V4.5A3.5 3.5 0 0 0 8 1m2 6V4.5a2 2 0 1 0-4 0V7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockClosedSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.5a5.25 5.25 0 0 0-5.25 5.25v3a3 3 0 0 0-3 3v6.75a3 3 0 0 0 3 3h10.5a3 3 0 0 0 3-3v-6.75a3 3 0 0 0-3-3v-3c0-2.9-2.35-5.25-5.25-5.25m3.75 8.25v-3a3.75 3.75 0 1 0-7.5 0v3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockClosedTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 1a4.5 4.5 0 0 0-4.5 4.5V9H5a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2h-.5V5.5A4.5 4.5 0 0 0 10 1m3 8V5.5a3 3 0 1 0-6 0V9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOpen(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.5 10.5V6.75a4.5 4.5 0 1 1 9 0v3.75M3.75 21.75h10.5a2.25 2.25 0 0 0 2.25-2.25v-6.75a2.25 2.25 0 0 0-2.25-2.25H3.75a2.25 2.25 0 0 0-2.25 2.25v6.75a2.25 2.25 0 0 0 2.25 2.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOpenSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.5 1A3.5 3.5 0 0 0 8 4.5V7H2.5A1.5 1.5 0 0 0 1 8.5v5A1.5 1.5 0 0 0 2.5 15h7a1.5 1.5 0 0 0 1.5-1.5v-5A1.5 1.5 0 0 0 9.5 7V4.5a2 2 0 1 1 4 0v1.75a.75.75 0 0 0 1.5 0V4.5A3.5 3.5 0 0 0 11.5 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOpenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 1.5c2.9 0 5.25 2.35 5.25 5.25v3.75a.75.75 0 0 1-1.5 0V6.75a3.75 3.75 0 1 0-7.5 0v3a3 3 0 0 1 3 3v6.75a3 3 0 0 1-3 3H3.75a3 3 0 0 1-3-3v-6.75a3 3 0 0 1 3-3h9v-3c0-2.9 2.35-5.25 5.25-5.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOpenTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 1A4.5 4.5 0 0 0 10 5.5V9H3a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2h-1.5V5.5a3 3 0 1 1 6 0v2.75a.75.75 0 0 0 1.5 0V5.5A4.5 4.5 0 0 0 14.5 1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifyingGlass(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m21 21l-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifyingGlassCircle(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15.75 15.75l-2.488-2.488m0 0a3.375 3.375 0 1 0-4.773-4.773a3.375 3.375 0 0 0 4.772 4.772M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifyingGlassCircleSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M5.94 8.06a1.5 1.5 0 1 1 2.12-2.12a1.5 1.5 0 0 1-2.12 2.12"/><path fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14M4.879 4.879a3 3 0 0 0 3.645 4.706L9.72 10.78a.75.75 0 0 0 1.061-1.06L9.585 8.524A3.001 3.001 0 0 0 4.879 4.88" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifyingGlassCircleSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8.25 10.875a2.625 2.625 0 1 1 5.25 0a2.625 2.625 0 0 1-5.25 0"/><path fill-rule="evenodd" d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75s9.75-4.365 9.75-9.75S17.385 2.25 12 2.25m-1.125 4.5a4.125 4.125 0 1 0 2.338 7.524l2.007 2.006a.75.75 0 1 0 1.06-1.06l-2.006-2.007a4.125 4.125 0 0 0-3.399-6.463" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifyingGlassCircleTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M6.5 9a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0"/><path fill-rule="evenodd" d="M10 18a8 8 0 1 0 0-16a8 8 0 0 0 0 16M9 5a4 4 0 1 0 2.248 7.309l1.472 1.471a.75.75 0 1 0 1.06-1.06l-1.471-1.472A4 4 0 0 0 9 5" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifyingGlassMinus(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m21 21l-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607M13.5 10.5h-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifyingGlassMinusSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.75 6.25h-3.5a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5"/><path fill="currentColor" fill-rule="evenodd" d="M7 12c1.11 0 2.136-.362 2.965-.974l2.755 2.754a.75.75 0 1 0 1.06-1.06l-2.754-2.755A5 5 0 1 0 7 12m0-1.5a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifyingGlassMinusSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.5 3.75a6.75 6.75 0 1 0 0 13.5a6.75 6.75 0 0 0 0-13.5M2.25 10.5a8.25 8.25 0 1 1 14.59 5.28l4.69 4.69a.75.75 0 1 1-1.06 1.06l-4.69-4.69A8.25 8.25 0 0 1 2.25 10.5m4.5 0a.75.75 0 0 1 .75-.75h6a.75.75 0 0 1 0 1.5h-6a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifyingGlassMinusTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M6.75 8.25a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 0-1.5z"/><path fill-rule="evenodd" d="M9 2a7 7 0 1 0 4.391 12.452l3.329 3.328a.75.75 0 1 0 1.06-1.06l-3.328-3.329A7 7 0 0 0 9 2M3.5 9a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifyingGlassPlus(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m21 21l-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607M10.5 7.5v6m3-3h-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifyingGlassPlusSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.25 8.75v-1h-1a.75.75 0 0 1 0-1.5h1v-1a.75.75 0 0 1 1.5 0v1h1a.75.75 0 0 1 0 1.5h-1v1a.75.75 0 0 1-1.5 0"/><path fill="currentColor" fill-rule="evenodd" d="M7 12c1.11 0 2.136-.362 2.965-.974l2.755 2.754a.75.75 0 1 0 1.06-1.06l-2.754-2.755A5 5 0 1 0 7 12m0-1.5a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifyingGlassPlusSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.5 3.75a6.75 6.75 0 1 0 0 13.5a6.75 6.75 0 0 0 0-13.5M2.25 10.5a8.25 8.25 0 1 1 14.59 5.28l4.69 4.69a.75.75 0 1 1-1.06 1.06l-4.69-4.69A8.25 8.25 0 0 1 2.25 10.5m8.25-3.75a.75.75 0 0 1 .75.75v2.25h2.25a.75.75 0 0 1 0 1.5h-2.25v2.25a.75.75 0 0 1-1.5 0v-2.25H7.5a.75.75 0 0 1 0-1.5h2.25V7.5a.75.75 0 0 1 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifyingGlassPlusTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M9 6a.75.75 0 0 1 .75.75v1.5h1.5a.75.75 0 0 1 0 1.5h-1.5v1.5a.75.75 0 0 1-1.5 0v-1.5h-1.5a.75.75 0 0 1 0-1.5h1.5v-1.5A.75.75 0 0 1 9 6"/><path fill-rule="evenodd" d="M2 9a7 7 0 1 1 12.452 4.391l3.328 3.329a.75.75 0 1 1-1.06 1.06l-3.329-3.328A7 7 0 0 1 2 9m7-5.5a5.5 5.5 0 1 0 0 11a5.5 5.5 0 0 0 0-11" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifyingGlassSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.965 11.026a5 5 0 1 1 1.06-1.06l2.755 2.754a.75.75 0 1 1-1.06 1.06zM10.5 7a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifyingGlassSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.5 3.75a6.75 6.75 0 1 0 0 13.5a6.75 6.75 0 0 0 0-13.5M2.25 10.5a8.25 8.25 0 1 1 14.59 5.28l4.69 4.69a.75.75 0 1 1-1.06 1.06l-4.69-4.69A8.25 8.25 0 0 1 2.25 10.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifyingGlassTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9 3.5a5.5 5.5 0 1 0 0 11a5.5 5.5 0 0 0 0-11M2 9a7 7 0 1 1 12.452 4.391l3.328 3.329a.75.75 0 1 1-1.06 1.06l-3.329-3.328A7 7 0 0 1 2 9" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Map(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 6.75V15m6-6v8.25m.503 3.499l4.875-2.438c.381-.19.622-.58.622-1.006V4.82c0-.836-.88-1.38-1.628-1.006l-3.869 1.934a1.12 1.12 0 0 1-1.006 0L9.503 3.252a1.125 1.125 0 0 0-1.006 0L3.622 5.689A1.125 1.125 0 0 0 3 6.695V19.18c0 .836.88 1.38 1.628 1.006l3.869-1.934a1.12 1.12 0 0 1 1.006 0l4.994 2.497c.317.158.69.158 1.006 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapPin(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15 10.5a3 3 0 1 1-6 0a3 3 0 0 1 6 0"/><path d="M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1 1 15 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapPinSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7.539 14.841l.003.003l.002.002a.755.755 0 0 0 .912 0l.002-.002l.003-.003l.012-.009a5.57 5.57 0 0 0 .19-.153a15.588 15.588 0 0 0 2.046-2.082C11.81 11.235 13 9.255 13 7A5 5 0 0 0 3 7c0 2.255 1.19 4.235 2.292 5.597a15.591 15.591 0 0 0 2.046 2.082a8.916 8.916 0 0 0 .189.153zM8 8.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapPinSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m11.54 22.351l.07.04l.028.016a.76.76 0 0 0 .723 0l.028-.015l.071-.041a16.975 16.975 0 0 0 1.144-.742a19.58 19.58 0 0 0 2.683-2.282c1.944-1.99 3.963-4.98 3.963-8.827a8.25 8.25 0 0 0-16.5 0c0 3.846 2.02 6.837 3.963 8.827a19.58 19.58 0 0 0 2.682 2.282a16.975 16.975 0 0 0 1.145.742M12 13.5a3 3 0 1 0 0-6a3 3 0 0 0 0 6" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapPinTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m9.69 18.933l.003.001C9.89 19.02 10 19 10 19s.11.02.308-.066l.002-.001l.006-.003l.018-.008a5.741 5.741 0 0 0 .281-.14c.186-.096.446-.24.757-.433c.62-.384 1.445-.966 2.274-1.765C15.302 14.988 17 12.493 17 9A7 7 0 1 0 3 9c0 3.492 1.698 5.988 3.355 7.584a13.731 13.731 0 0 0 2.273 1.765a11.842 11.842 0 0 0 .976.544l.062.029l.018.008zM10 11.25a2.25 2.25 0 1 0 0-4.5a2.25 2.25 0 0 0 0 4.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.37 2.257a1.25 1.25 0 0 1 1.214-.054l3.378 1.69l2.133-1.313A1.25 1.25 0 0 1 14 3.644v7.326c0 .434-.225.837-.595 1.065l-2.775 1.708a1.25 1.25 0 0 1-1.214.053l-3.378-1.689l-2.133 1.313A1.25 1.25 0 0 1 2 12.355V5.029c0-.434.225-.837.595-1.064zM6 4a.75.75 0 0 1 .75.75v4.5a.75.75 0 0 1-1.5 0v-4.5A.75.75 0 0 1 6 4m4.75 2.75a.75.75 0 0 0-1.5 0v4.5a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.161 2.58a1.875 1.875 0 0 1 1.678 0l4.993 2.498c.106.052.23.052.336 0l3.869-1.935A1.875 1.875 0 0 1 21.75 4.82v12.485c0 .71-.401 1.36-1.037 1.677l-4.875 2.437a1.875 1.875 0 0 1-1.676 0l-4.994-2.497a.375.375 0 0 0-.336 0l-3.868 1.935A1.875 1.875 0 0 1 2.25 19.18V6.695c0-.71.401-1.36 1.036-1.677zM9 6a.75.75 0 0 1 .75.75V15a.75.75 0 0 1-1.5 0V6.75A.75.75 0 0 1 9 6m6.75 3a.75.75 0 0 0-1.5 0v8.25a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.157 2.176a1.5 1.5 0 0 0-1.147 0l-4.084 1.69A1.5 1.5 0 0 0 2 5.25v10.877a1.5 1.5 0 0 0 2.074 1.386l3.51-1.452l4.26 1.762a1.5 1.5 0 0 0 1.146 0l4.083-1.69A1.5 1.5 0 0 0 18 14.75V3.872a1.5 1.5 0 0 0-2.073-1.386l-3.51 1.452zM7.58 5a.75.75 0 0 1 .75.75v6.5a.75.75 0 0 1-1.5 0v-6.5A.75.75 0 0 1 7.58 5m5.59 2.75a.75.75 0 0 0-1.5 0v6.5a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Megaphone(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.34 15.84a24.07 24.07 0 0 0-2.09-.09H7.5a4.5 4.5 0 1 1 0-9h.75c.704 0 1.402-.03 2.09-.09m0 9.18c.253.962.584 1.892.985 2.783c.247.55.06 1.21-.463 1.511l-.657.38c-.551.318-1.26.117-1.527-.461a20.845 20.845 0 0 1-1.44-4.282m3.102.069a18.03 18.03 0 0 1-.59-4.59c0-1.586.205-3.124.59-4.59m0 9.18a23.848 23.848 0 0 1 8.835 2.535M10.34 6.66a23.847 23.847 0 0 0 8.835-2.535m0 0A23.74 23.74 0 0 0 18.795 3m.38 1.125a23.91 23.91 0 0 1 1.014 5.395m-1.014 8.855c-.118.38-.245.754-.38 1.125m.38-1.125a23.91 23.91 0 0 0 1.014-5.395m0-3.46a2.249 2.249 0 0 1 0 3.46m0-3.46a24.347 24.347 0 0 1 0 3.46"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MegaphoneSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.407 2.59a.75.75 0 0 0-1.464.326c.365 1.636.557 3.337.557 5.084c0 1.747-.192 3.448-.557 5.084a.75.75 0 0 0 1.464.327c.264-1.185.444-2.402.531-3.644a2 2 0 0 0 0-3.534a24.736 24.736 0 0 0-.531-3.643M4.348 11H4a3 3 0 0 1 0-6h2c1.647 0 3.217-.332 4.646-.933C10.878 5.341 11 6.655 11 8c0 1.345-.122 2.659-.354 3.933a11.946 11.946 0 0 0-4.23-.925c.203.718.478 1.407.816 2.057c.12.23.057.515-.155.663l-.828.58a.484.484 0 0 1-.707-.16A12.91 12.91 0 0 1 4.348 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MegaphoneSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.881 4.345A23.112 23.112 0 0 1 8.25 6H7.5a5.25 5.25 0 0 0-.88 10.427a21.593 21.593 0 0 0 1.378 3.94c.464 1.004 1.674 1.32 2.582.796l.657-.379c.88-.508 1.165-1.593.772-2.468a17.116 17.116 0 0 1-.628-1.607c1.918.258 3.76.75 5.5 1.446A21.727 21.727 0 0 0 18 11.25c0-2.414-.393-4.735-1.119-6.905m1.379-.605a23.22 23.22 0 0 1 1.24 7.51a23.22 23.22 0 0 1-1.41 7.992a.75.75 0 1 0 1.409.516a24.555 24.555 0 0 0 1.415-6.43a2.992 2.992 0 0 0 .836-2.078a2.99 2.99 0 0 0-.836-2.078a24.65 24.65 0 0 0-1.415-6.43a.75.75 0 1 0-1.409.516c.059.16.116.321.17.483"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MegaphoneTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.92 3.845a19.362 19.362 0 0 1-6.3 1.98C6.765 5.942 5.89 6 5 6a4 4 0 0 0-.504 7.969a15.97 15.97 0 0 0 1.271 3.34c.397.771 1.342 1 2.05.59l.867-.5c.726-.419.94-1.32.588-2.02a12.22 12.22 0 0 1-.448-1.004c1.8.357 3.511.963 5.096 1.78A17.964 17.964 0 0 0 15 10c0-2.162-.381-4.235-1.08-6.155m1.323-.748A19.456 19.456 0 0 1 16.5 10c0 2.43-.445 4.758-1.257 6.904l-.03.077a.75.75 0 0 0 1.401.537a20.903 20.903 0 0 0 1.312-5.745a2 2 0 0 0 0-3.546a20.902 20.902 0 0 0-1.312-5.745a.75.75 0 0 0-1.4.537z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microphone(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 18.75a6 6 0 0 0 6-6v-1.5m-6 7.5a6 6 0 0 1-6-6v-1.5m6 7.5v3.75m-3.75 0h7.5M12 15.75a3 3 0 0 1-3-3V4.5a3 3 0 1 1 6 0v8.25a3 3 0 0 1-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1a2 2 0 0 0-2 2v4a2 2 0 1 0 4 0V3a2 2 0 0 0-2-2"/><path fill="currentColor" d="M4.5 7A.75.75 0 0 0 3 7a5.001 5.001 0 0 0 4.25 4.944V13.5h-1.5a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 0-1.5h-1.5v-1.556A5.001 5.001 0 0 0 13 7a.75.75 0 0 0-1.5 0a3.5 3.5 0 1 1-7 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8.25 4.5a3.75 3.75 0 1 1 7.5 0v8.25a3.75 3.75 0 1 1-7.5 0z"/><path d="M6 10.5a.75.75 0 0 1 .75.75v1.5a5.25 5.25 0 1 0 10.5 0v-1.5a.75.75 0 0 1 1.5 0v1.5a6.751 6.751 0 0 1-6 6.709v2.291h3a.75.75 0 0 1 0 1.5h-7.5a.75.75 0 0 1 0-1.5h3v-2.291a6.751 6.751 0 0 1-6-6.709v-1.5A.75.75 0 0 1 6 10.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M7 4a3 3 0 0 1 6 0v6a3 3 0 1 1-6 0z"/><path d="M5.5 9.643a.75.75 0 0 0-1.5 0V10c0 3.06 2.29 5.585 5.25 5.954V17.5h-1.5a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 0-1.5h-1.5v-1.546A6.001 6.001 0 0 0 16 10v-.357a.75.75 0 0 0-1.5 0V10a4.5 4.5 0 0 1-9 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 12h14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircle(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 12H9m12 0a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircleSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14m4-7a.75.75 0 0 0-.75-.75h-6.5a.75.75 0 0 0 0 1.5h6.5A.75.75 0 0 0 12 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircleSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75s9.75-4.365 9.75-9.75S17.385 2.25 12 2.25m3 10.5a.75.75 0 0 0 0-1.5H9a.75.75 0 0 0 0 1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircleTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 18a8 8 0 1 0 0-16a8 8 0 0 0 0 16M6.75 9.25a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.75 7.25a.75.75 0 0 0 0 1.5h8.5a.75.75 0 0 0 0-1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSmall(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 12H6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSmallSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.25 12a.75.75 0 0 1 .75-.75h12a.75.75 0 0 1 0 1.5H6a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSmallTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.75 9.25a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.25 12a.75.75 0 0 1 .75-.75h14a.75.75 0 0 1 0 1.5H5a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 10a.75.75 0 0 1 .75-.75h10.5a.75.75 0 0 1 0 1.5H4.75A.75.75 0 0 1 4 10" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Moon(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21.752 15.002A9.72 9.72 0 0 1 18 15.75A9.75 9.75 0 0 1 8.25 6c0-1.33.266-2.597.748-3.752A9.753 9.753 0 0 0 3 11.25A9.75 9.75 0 0 0 12.75 21a9.753 9.753 0 0 0 9.002-5.998"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.438 10.148c.19-.425-.321-.787-.748-.601A5.5 5.5 0 0 1 6.453 2.31c.186-.427-.176-.938-.6-.748a6.501 6.501 0 1 0 8.585 8.586"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.528 1.718a.75.75 0 0 1 .162.819A8.97 8.97 0 0 0 9 6a9 9 0 0 0 9 9a8.97 8.97 0 0 0 3.463-.69a.75.75 0 0 1 .981.98a10.503 10.503 0 0 1-9.694 6.46c-5.799 0-10.5-4.7-10.5-10.5c0-4.368 2.667-8.112 6.46-9.694a.75.75 0 0 1 .818.162" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.455 2.004a.75.75 0 0 1 .26.77a7 7 0 0 0 9.958 7.967a.75.75 0 0 1 1.067.853A8.5 8.5 0 1 1 6.647 1.921a.75.75 0 0 1 .808.083" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicalNote(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 9l10.5-3m0 6.553v3.75a2.25 2.25 0 0 1-1.632 2.163l-1.32.377a1.803 1.803 0 1 1-.99-3.467l2.31-.66a2.25 2.25 0 0 0 1.632-2.163m0 0V2.25L9 5.25v10.303m0 0v3.75a2.25 2.25 0 0 1-1.632 2.163l-1.32.377a1.803 1.803 0 0 1-.99-3.467l2.31-.66A2.25 2.25 0 0 0 9 15.553"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicalNoteSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 1.75a.75.75 0 0 0-.89-.737l-7.502 1.43a.75.75 0 0 0-.61.736v2.5c0 .018 0 .036.002.054V9.73a1 1 0 0 1-.813.983l-.58.11a1.978 1.978 0 0 0 .741 3.886l.603-.115c.9-.171 1.55-.957 1.55-1.873v-1.543l-.001-.043V6.3l6-1.143v3.146a1 1 0 0 1-.813.982l-.584.111a1.978 1.978 0 0 0 .74 3.887l.326-.063A2.252 2.252 0 0 0 14 11.007z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicalNoteSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M19.952 1.651a.75.75 0 0 1 .298.599v14.053a3 3 0 0 1-2.176 2.884l-1.32.377a2.553 2.553 0 1 1-1.403-4.909l2.311-.66a1.5 1.5 0 0 0 1.088-1.442V6.994l-9 2.572v9.737a3 3 0 0 1-2.176 2.884l-1.32.377a2.553 2.553 0 1 1-1.402-4.909l2.31-.66a1.5 1.5 0 0 0 1.088-1.442V5.25a.75.75 0 0 1 .544-.721l10.5-3a.75.75 0 0 1 .658.122" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicalNoteTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M17.721 1.599a.75.75 0 0 1 .279.583v11.29a2.25 2.25 0 0 1-1.774 2.2l-2.041.44a2.216 2.216 0 0 1-.938-4.332l2.662-.577a.75.75 0 0 0 .591-.733V6.112l-8 1.73v7.684a2.25 2.25 0 0 1-1.774 2.2l-2.042.44a2.216 2.216 0 1 1-.935-4.331l2.659-.573A.75.75 0 0 0 7 12.529V4.236a.75.75 0 0 1 .591-.733l9.5-2.054a.75.75 0 0 1 .63.15" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Newspaper(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 7.5h1.5m-1.5 3h1.5m-7.5 3h7.5m-7.5 3h7.5m3-9h3.375c.621 0 1.125.504 1.125 1.125V18a2.25 2.25 0 0 1-2.25 2.25M16.5 7.5V18a2.25 2.25 0 0 0 2.25 2.25M16.5 7.5V4.875c0-.621-.504-1.125-1.125-1.125H4.125C3.504 3.75 3 4.254 3 4.875V18a2.25 2.25 0 0 0 2.25 2.25h13.5M6 7.5h3v3H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NewspaperSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 3a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v9a2 2 0 0 0 2 2h8a2 2 0 0 1-2-2zM4 4h4v2H4zm4 3.5H4V9h4zm-4 3h4V12H4z" clip-rule="evenodd"/><path fill="currentColor" d="M13 5h-1.5v6.25a1.25 1.25 0 1 0 2.5 0V6a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NewspaperSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M4.125 3C3.089 3 2.25 3.84 2.25 4.875V18a3 3 0 0 0 3 3h15a3 3 0 0 1-3-3V4.875C17.25 3.839 16.41 3 15.375 3zM12 9.75a.75.75 0 0 0 0 1.5h1.5a.75.75 0 0 0 0-1.5zm-.75-2.25a.75.75 0 0 1 .75-.75h1.5a.75.75 0 0 1 0 1.5H12a.75.75 0 0 1-.75-.75M6 12.75a.75.75 0 0 0 0 1.5h7.5a.75.75 0 0 0 0-1.5zm-.75 3.75a.75.75 0 0 1 .75-.75h7.5a.75.75 0 0 1 0 1.5H6a.75.75 0 0 1-.75-.75M6 6.75a.75.75 0 0 0-.75.75v3c0 .414.336.75.75.75h3a.75.75 0 0 0 .75-.75v-3A.75.75 0 0 0 9 6.75z" clip-rule="evenodd"/><path d="M18.75 6.75h1.875c.621 0 1.125.504 1.125 1.125V18a1.5 1.5 0 0 1-3 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NewspaperTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M2 3.5A1.5 1.5 0 0 1 3.5 2h9A1.5 1.5 0 0 1 14 3.5v11.75A2.75 2.75 0 0 0 16.75 18h-12A2.75 2.75 0 0 1 2 15.25zm3.75 7a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 0-1.5zm0 3a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 0-1.5zM5 5.75A.75.75 0 0 1 5.75 5h4.5a.75.75 0 0 1 .75.75v2.5a.75.75 0 0 1-.75.75h-4.5A.75.75 0 0 1 5 8.25z" clip-rule="evenodd"/><path d="M16.5 6.5h-1v8.75a1.25 1.25 0 1 0 2.5 0V8a1.5 1.5 0 0 0-1.5-1.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoSymbol(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18.364 18.364A9 9 0 0 0 5.636 5.636m12.728 12.728A9 9 0 0 1 5.636 5.636m12.728 12.728L5.636 5.636"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoSymbolSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.05 3.05a7 7 0 1 1 9.9 9.9a7 7 0 0 1-9.9-9.9m1.627.566l7.707 7.707a5.501 5.501 0 0 0-7.707-7.707m6.646 8.768L3.616 4.677a5.501 5.501 0 0 0 7.707 7.707" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoSymbolSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.72 5.66l11.62 11.62A8.25 8.25 0 0 0 6.72 5.66m10.56 12.68L5.66 6.72a8.25 8.25 0 0 0 11.62 11.62M5.105 5.106c3.807-3.808 9.98-3.808 13.788 0c3.808 3.807 3.808 9.98 0 13.788c-3.807 3.808-9.98 3.808-13.788 0c-3.808-3.807-3.808-9.98 0-13.788" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoSymbolTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m5.965 4.904l9.131 9.131a6.5 6.5 0 0 0-9.131-9.131m8.07 10.192L4.904 5.965a6.5 6.5 0 0 0 9.131 9.131M4.343 4.343a8 8 0 1 1 11.314 11.314A8 8 0 0 1 4.343 4.343" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaintBrush(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.53 16.122a3 3 0 0 0-5.78 1.128a2.25 2.25 0 0 1-2.4 2.245a4.5 4.5 0 0 0 8.4-2.245c0-.399-.078-.78-.22-1.128m0 0a15.998 15.998 0 0 0 3.388-1.62m-5.043-.025a15.994 15.994 0 0 1 1.622-3.395m3.42 3.42a15.995 15.995 0 0 0 4.764-4.648l3.876-5.814a1.151 1.151 0 0 0-1.597-1.597L14.146 6.32a15.996 15.996 0 0 0-4.649 4.764m3.42 3.42a6.776 6.776 0 0 0-3.42-3.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaintBrushSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.613 1.258a1.535 1.535 0 0 1 2.13 2.129l-1.905 2.856a8 8 0 0 1-3.56 2.939a4.011 4.011 0 0 0-2.46-2.46a8 8 0 0 1 2.94-3.56zM5.5 8A2.5 2.5 0 0 0 3 10.5a.5.5 0 0 1-.7.459a.75.75 0 0 0-.983 1A3.5 3.5 0 0 0 8 10.5A2.5 2.5 0 0 0 5.5 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaintBrushSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M20.599 1.5c-.376 0-.743.111-1.055.32l-5.08 3.385a18.747 18.747 0 0 0-3.471 2.987a10.04 10.04 0 0 1 4.815 4.815a18.748 18.748 0 0 0 2.987-3.472l3.386-5.079A1.902 1.902 0 0 0 20.599 1.5m-8.3 14.025a18.76 18.76 0 0 0 1.896-1.207a8.026 8.026 0 0 0-4.513-4.513A18.75 18.75 0 0 0 8.475 11.7l-.278.5a5.26 5.26 0 0 1 3.601 3.602zM6.75 13.5A3.75 3.75 0 0 0 3 17.25a1.5 1.5 0 0 1-1.601 1.497a.75.75 0 0 0-.7 1.123a5.25 5.25 0 0 0 9.8-2.62a3.75 3.75 0 0 0-3.75-3.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaintBrushTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.993 1.385a1.87 1.87 0 0 1 2.623 2.622l-4.03 5.27a12.749 12.749 0 0 1-4.237 3.562a4.508 4.508 0 0 0-3.188-3.188a12.75 12.75 0 0 1 3.562-4.236zM6 11a3 3 0 0 0-3 3a.5.5 0 0 1-.72.45a.75.75 0 0 0-1.035.931A4.001 4.001 0 0 0 9 14.004V14a3.01 3.01 0 0 0-1.66-2.685A2.99 2.99 0 0 0 6 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperAirplane(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 12L3.269 3.125A59.769 59.769 0 0 1 21.486 12a59.768 59.768 0 0 1-18.217 8.875zm0 0h7.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperAirplaneSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.87 2.298a.75.75 0 0 0-.812 1.021L3.39 6.624a1 1 0 0 0 .928.626H8.25a.75.75 0 0 1 0 1.5H4.318a1 1 0 0 0-.927.626l-1.333 3.305a.75.75 0 0 0 .811 1.022a24.89 24.89 0 0 0 11.668-5.115a.75.75 0 0 0 0-1.175A24.89 24.89 0 0 0 2.869 2.298"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperAirplaneSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.478 2.404a.75.75 0 0 0-.926.941l2.432 7.905H13.5a.75.75 0 0 1 0 1.5H4.984l-2.432 7.905a.75.75 0 0 0 .926.94a60.519 60.519 0 0 0 18.445-8.986a.75.75 0 0 0 0-1.218A60.517 60.517 0 0 0 3.478 2.404"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperAirplaneTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.105 2.288a.75.75 0 0 0-.826.95l1.414 4.926A1.5 1.5 0 0 0 5.135 9.25h6.115a.75.75 0 0 1 0 1.5H5.135a1.5 1.5 0 0 0-1.442 1.086l-1.414 4.926a.75.75 0 0 0 .826.95a28.897 28.897 0 0 0 15.293-7.155a.75.75 0 0 0 0-1.114A28.897 28.897 0 0 0 3.105 2.288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperClip(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m18.375 12.739l-7.693 7.693a4.5 4.5 0 0 1-6.364-6.364l10.94-10.94A3 3 0 1 1 19.5 7.372L8.552 18.32m.009-.01l-.01.01m5.699-9.941l-7.81 7.81a1.5 1.5 0 0 0 2.112 2.13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperClipSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.914 4.086a2 2 0 0 0-2.828 0l-5 5a2 2 0 1 0 2.828 2.828l.556-.555a.75.75 0 0 1 1.06 1.06l-.555.556a3.5 3.5 0 0 1-4.95-4.95l5-5a3.5 3.5 0 0 1 4.95 4.95l-1.972 1.972a2.125 2.125 0 0 1-3.006-3.005L9.97 4.97a.75.75 0 1 1 1.06 1.06L9.058 8.003a.625.625 0 0 0 .884.883l1.972-1.972a2 2 0 0 0 0-2.828" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperClipSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M18.97 3.659a2.25 2.25 0 0 0-3.182 0l-10.94 10.94a3.75 3.75 0 1 0 5.304 5.303l7.693-7.693a.75.75 0 0 1 1.06 1.06l-7.693 7.693a5.25 5.25 0 1 1-7.424-7.424l10.939-10.94a3.75 3.75 0 1 1 5.303 5.304L9.097 18.835l-.008.008l-.007.007l-.002.002l-.003.002A2.25 2.25 0 0 1 5.91 15.66l7.81-7.81a.75.75 0 0 1 1.061 1.06l-7.81 7.81a.75.75 0 0 0 1.054 1.068L18.97 6.84a2.25 2.25 0 0 0 0-3.182" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperClipTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.621 4.379a3 3 0 0 0-4.242 0l-7 7a3 3 0 0 0 4.241 4.243h.001l.497-.5a.75.75 0 0 1 1.064 1.057l-.498.501l-.002.002a4.5 4.5 0 0 1-6.364-6.364l7-7a4.5 4.5 0 0 1 6.368 6.36l-3.455 3.553A2.625 2.625 0 1 1 9.52 9.52l3.45-3.451a.75.75 0 1 1 1.061 1.06l-3.45 3.451a1.125 1.125 0 0 0 1.587 1.595l3.454-3.553a3 3 0 0 0 0-4.242" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.75 5.25v13.5m-7.5-13.5v13.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseCircle(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.25 9v6m-4.5 0V9M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseCircleSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0M5.5 5.5A.5.5 0 0 1 6 5h.5a.5.5 0 0 1 .5.5v5a.5.5 0 0 1-.5.5H6a.5.5 0 0 1-.5-.5zm4-.5a.5.5 0 0 0-.5.5v5a.5.5 0 0 0 .5.5h.5a.5.5 0 0 0 .5-.5v-5A.5.5 0 0 0 10 5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseCircleSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 12c0-5.385 4.365-9.75 9.75-9.75s9.75 4.365 9.75 9.75s-4.365 9.75-9.75 9.75S2.25 17.385 2.25 12M9 8.25a.75.75 0 0 0-.75.75v6c0 .414.336.75.75.75h.75a.75.75 0 0 0 .75-.75V9a.75.75 0 0 0-.75-.75zm5.25 0a.75.75 0 0 0-.75.75v6c0 .414.336.75.75.75H15a.75.75 0 0 0 .75-.75V9a.75.75 0 0 0-.75-.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseCircleTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 10a8 8 0 1 1 16 0a8 8 0 0 1-16 0m5-2.25A.75.75 0 0 1 7.75 7h.5a.75.75 0 0 1 .75.75v4.5a.75.75 0 0 1-.75.75h-.5a.75.75 0 0 1-.75-.75zm4 0a.75.75 0 0 1 .75-.75h.5a.75.75 0 0 1 .75.75v4.5a.75.75 0 0 1-.75.75h-.5a.75.75 0 0 1-.75-.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 2a.5.5 0 0 0-.5.5v11a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 .5-.5v-11a.5.5 0 0 0-.5-.5zm6 0a.5.5 0 0 0-.5.5v11a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 .5-.5v-11a.5.5 0 0 0-.5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.75 5.25a.75.75 0 0 1 .75-.75H9a.75.75 0 0 1 .75.75v13.5a.75.75 0 0 1-.75.75H7.5a.75.75 0 0 1-.75-.75zm7.5 0A.75.75 0 0 1 15 4.5h1.5a.75.75 0 0 1 .75.75v13.5a.75.75 0 0 1-.75.75H15a.75.75 0 0 1-.75-.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.75 3a.75.75 0 0 0-.75.75v12.5c0 .414.336.75.75.75h1.5a.75.75 0 0 0 .75-.75V3.75A.75.75 0 0 0 7.25 3zm7 0a.75.75 0 0 0-.75.75v12.5c0 .414.336.75.75.75h1.5a.75.75 0 0 0 .75-.75V3.75a.75.75 0 0 0-.75-.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pencil(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m16.862 4.487l1.687-1.688a1.875 1.875 0 1 1 2.652 2.652L6.832 19.82a4.5 4.5 0 0 1-1.897 1.13l-2.685.8l.8-2.685a4.5 4.5 0 0 1 1.13-1.897zm0 0L19.5 7.125"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PencilSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.013 2.513a1.75 1.75 0 0 1 2.475 2.474L6.226 12.25a2.751 2.751 0 0 1-.892.596l-2.047.848a.75.75 0 0 1-.98-.98l.848-2.047a2.75 2.75 0 0 1 .596-.892z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PencilSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.731 2.269a2.625 2.625 0 0 0-3.712 0l-1.157 1.157l3.712 3.712l1.157-1.157a2.625 2.625 0 0 0 0-3.712m-2.218 5.93l-3.712-3.712l-12.15 12.15a5.25 5.25 0 0 0-1.32 2.214l-.8 2.685a.75.75 0 0 0 .933.933l2.685-.8a5.25 5.25 0 0 0 2.214-1.32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PencilSquare(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m16.862 4.487l1.687-1.688a1.875 1.875 0 1 1 2.652 2.652L10.582 16.07a4.5 4.5 0 0 1-1.897 1.13L6 18l.8-2.685a4.5 4.5 0 0 1 1.13-1.897zm0 0L19.5 7.125M18 14v4.75A2.25 2.25 0 0 1 15.75 21H5.25A2.25 2.25 0 0 1 3 18.75V8.25A2.25 2.25 0 0 1 5.25 6H10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PencilSquareSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M13.488 2.513a1.75 1.75 0 0 0-2.475 0L6.75 6.774a2.75 2.75 0 0 0-.596.892l-.848 2.047a.75.75 0 0 0 .98.98l2.047-.848a2.75 2.75 0 0 0 .892-.596l4.262-4.262a1.75 1.75 0 0 0 0-2.474"/><path d="M4.75 3.5c-.69 0-1.25.56-1.25 1.25v6.5c0 .69.56 1.25 1.25 1.25h6.5c.69 0 1.25-.56 1.25-1.25V9A.75.75 0 0 1 14 9v2.25A2.75 2.75 0 0 1 11.25 14h-6.5A2.75 2.75 0 0 1 2 11.25v-6.5A2.75 2.75 0 0 1 4.75 2H7a.75.75 0 0 1 0 1.5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PencilSquareSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M21.731 2.269a2.625 2.625 0 0 0-3.712 0l-1.157 1.157l3.712 3.712l1.157-1.157a2.625 2.625 0 0 0 0-3.712m-2.218 5.93l-3.712-3.712l-8.4 8.4a5.25 5.25 0 0 0-1.32 2.214l-.8 2.685a.75.75 0 0 0 .933.933l2.685-.8a5.25 5.25 0 0 0 2.214-1.32z"/><path d="M5.25 5.25a3 3 0 0 0-3 3v10.5a3 3 0 0 0 3 3h10.5a3 3 0 0 0 3-3V13.5a.75.75 0 0 0-1.5 0v5.25a1.5 1.5 0 0 1-1.5 1.5H5.25a1.5 1.5 0 0 1-1.5-1.5V8.25a1.5 1.5 0 0 1 1.5-1.5h5.25a.75.75 0 0 0 0-1.5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PencilSquareTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m5.433 13.917l1.262-3.155A4 4 0 0 1 7.58 9.42l6.92-6.918a2.121 2.121 0 0 1 3 3l-6.92 6.918c-.383.383-.84.685-1.343.886l-3.154 1.262a.5.5 0 0 1-.65-.65"/><path d="M3.5 5.75c0-.69.56-1.25 1.25-1.25H10A.75.75 0 0 0 10 3H4.75A2.75 2.75 0 0 0 2 5.75v9.5A2.75 2.75 0 0 0 4.75 18h9.5A2.75 2.75 0 0 0 17 15.25V10a.75.75 0 0 0-1.5 0v5.25c0 .69-.56 1.25-1.25 1.25h-9.5c-.69 0-1.25-.56-1.25-1.25z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PencilTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.695 14.762l-1.262 3.155a.5.5 0 0 0 .65.65l3.155-1.262a4 4 0 0 0 1.343-.886L17.5 5.501a2.121 2.121 0 0 0-3-3L3.58 13.419a4 4 0 0 0-.885 1.343"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.25 6.75c0 8.284 6.716 15 15 15h2.25a2.25 2.25 0 0 0 2.25-2.25v-1.372c0-.516-.351-.966-.852-1.091l-4.423-1.106c-.44-.11-.902.055-1.173.417l-.97 1.293c-.282.376-.769.542-1.21.38a12.035 12.035 0 0 1-7.143-7.143c-.162-.441.004-.928.38-1.21l1.293-.97c.363-.271.527-.734.417-1.173L6.963 3.102a1.125 1.125 0 0 0-1.091-.852H4.5A2.25 2.25 0 0 0 2.25 4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneArrowDownLeft(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.25 9.75v-4.5m0 4.5h4.5m-4.5 0l6-6m-3 18c-8.284 0-15-6.716-15-15V4.5A2.25 2.25 0 0 1 4.5 2.25h1.372c.516 0 .966.351 1.091.852l1.106 4.423c.11.44-.054.902-.417 1.173l-1.293.97a1.062 1.062 0 0 0-.38 1.21a12.035 12.035 0 0 0 7.143 7.143c.441.162.928-.004 1.21-.38l.97-1.293a1.125 1.125 0 0 1 1.173-.417l4.423 1.106c.5.125.852.575.852 1.091V19.5a2.25 2.25 0 0 1-2.25 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneArrowDownLeftSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m4.922 6.752l-1.067.534a7.52 7.52 0 0 0 4.859 4.86l.534-1.068a1 1 0 0 1 1.046-.542l2.858.44a1 1 0 0 1 .848.988V13a1 1 0 0 1-1 1h-2c-.709 0-1.4-.082-2.062-.238a9.012 9.012 0 0 1-6.7-6.7A9.024 9.024 0 0 1 2 5V3a1 1 0 0 1 1-1h1.036a1 1 0 0 1 .988.848l.44 2.858a1 1 0 0 1-.542 1.046"/><path d="m11.56 5.5l2.22-2.22a.75.75 0 0 0-1.06-1.06L10.5 4.44V2.75a.75.75 0 0 0-1.5 0v3.5c0 .414.336.75.75.75h3.5a.75.75 0 0 0 0-1.5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneArrowDownLeftSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M19.5 9.75a.75.75 0 0 1-.75.75h-4.5a.75.75 0 0 1-.75-.75v-4.5a.75.75 0 0 1 1.5 0v2.69l4.72-4.72a.75.75 0 1 1 1.06 1.06L16.06 9h2.69a.75.75 0 0 1 .75.75"/><path d="M1.5 4.5a3 3 0 0 1 3-3h1.372c.86 0 1.61.586 1.819 1.42l1.105 4.423a1.875 1.875 0 0 1-.694 1.955l-1.293.97c-.135.101-.164.249-.126.352a11.285 11.285 0 0 0 6.697 6.697c.103.038.25.009.352-.126l.97-1.293a1.875 1.875 0 0 1 1.955-.694l4.423 1.105c.834.209 1.42.959 1.42 1.82V19.5a3 3 0 0 1-3 3h-2.25C8.552 22.5 1.5 15.448 1.5 6.75z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneArrowDownLeftTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 2A1.5 1.5 0 0 0 2 3.5V5c0 1.149.15 2.263.43 3.326a13.022 13.022 0 0 0 9.244 9.244c1.063.28 2.177.43 3.326.43h1.5a1.5 1.5 0 0 0 1.5-1.5v-1.148a1.5 1.5 0 0 0-1.175-1.465l-3.223-.716a1.5 1.5 0 0 0-1.767 1.052l-.267.933c-.117.41-.555.643-.95.48a11.542 11.542 0 0 1-6.254-6.254c-.163-.395.07-.833.48-.95l.933-.267a1.5 1.5 0 0 0 1.052-1.767l-.716-3.223A1.5 1.5 0 0 0 4.648 2zm13.22.22a.75.75 0 1 1 1.06 1.06L14.56 6.5h2.69a.75.75 0 0 1 0 1.5h-4.5a.75.75 0 0 1-.75-.75v-4.5a.75.75 0 0 1 1.5 0v2.69z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneArrowUpRight(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.25 3.75v4.5m0-4.5h-4.5m4.5 0l-6 6m3 12c-8.284 0-15-6.716-15-15V4.5A2.25 2.25 0 0 1 4.5 2.25h1.372c.516 0 .966.351 1.091.852l1.106 4.423c.11.44-.054.902-.417 1.173l-1.293.97a1.062 1.062 0 0 0-.38 1.21a12.035 12.035 0 0 0 7.143 7.143c.441.162.928-.004 1.21-.38l.97-1.293a1.125 1.125 0 0 1 1.173-.417l4.423 1.106c.5.125.852.575.852 1.091V19.5a2.25 2.25 0 0 1-2.25 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneArrowUpRightSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m4.922 6.752l-1.067.534a7.52 7.52 0 0 0 4.859 4.86l.534-1.068a1 1 0 0 1 1.046-.542l2.858.44a1 1 0 0 1 .848.988V13a1 1 0 0 1-1 1h-2c-.709 0-1.4-.082-2.062-.238a9.012 9.012 0 0 1-6.7-6.7A9.024 9.024 0 0 1 2 5V3a1 1 0 0 1 1-1h1.036a1 1 0 0 1 .988.848l.44 2.858a1 1 0 0 1-.542 1.046"/><path d="M9.22 5.72a.75.75 0 0 0 1.06 1.06l2.22-2.22v1.69a.75.75 0 0 0 1.5 0v-3.5a.75.75 0 0 0-.75-.75h-3.5a.75.75 0 0 0 0 1.5h1.69z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneArrowUpRightSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M15 3.75a.75.75 0 0 1 .75-.75h4.5a.75.75 0 0 1 .75.75v4.5a.75.75 0 0 1-1.5 0V5.56l-4.72 4.72a.75.75 0 1 1-1.06-1.06l4.72-4.72h-2.69a.75.75 0 0 1-.75-.75"/><path d="M1.5 4.5a3 3 0 0 1 3-3h1.372c.86 0 1.61.586 1.819 1.42l1.105 4.423a1.875 1.875 0 0 1-.694 1.955l-1.293.97c-.135.101-.164.249-.126.352a11.285 11.285 0 0 0 6.697 6.697c.103.038.25.009.352-.126l.97-1.293a1.875 1.875 0 0 1 1.955-.694l4.423 1.105c.834.209 1.42.959 1.42 1.82V19.5a3 3 0 0 1-3 3h-2.25C8.552 22.5 1.5 15.448 1.5 6.75z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneArrowUpRightTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 2A1.5 1.5 0 0 0 2 3.5V5c0 1.149.15 2.263.43 3.326a13.022 13.022 0 0 0 9.244 9.244c1.063.28 2.177.43 3.326.43h1.5a1.5 1.5 0 0 0 1.5-1.5v-1.148a1.5 1.5 0 0 0-1.175-1.465l-3.223-.716a1.5 1.5 0 0 0-1.767 1.052l-.267.933c-.117.41-.555.643-.95.48a11.542 11.542 0 0 1-6.254-6.254c-.163-.395.07-.833.48-.95l.933-.267a1.5 1.5 0 0 0 1.052-1.767l-.716-3.223A1.5 1.5 0 0 0 4.648 2zm13 2.56l-3.22 3.22a.75.75 0 1 1-1.06-1.06l3.22-3.22h-2.69a.75.75 0 0 1 0-1.5h4.5a.75.75 0 0 1 .75.75v4.5a.75.75 0 0 1-1.5 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m3.855 7.286l1.067-.534a1 1 0 0 0 .542-1.046l-.44-2.858A1 1 0 0 0 4.036 2H3a1 1 0 0 0-1 1v2c0 .709.082 1.4.238 2.062a9.012 9.012 0 0 0 6.7 6.7A9.024 9.024 0 0 0 11 14h2a1 1 0 0 0 1-1v-1.036a1 1 0 0 0-.848-.988l-2.858-.44a1 1 0 0 0-1.046.542l-.534 1.067a7.52 7.52 0 0 1-4.86-4.859" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 4.5a3 3 0 0 1 3-3h1.372c.86 0 1.61.586 1.819 1.42l1.105 4.423a1.875 1.875 0 0 1-.694 1.955l-1.293.97c-.135.101-.164.249-.126.352a11.285 11.285 0 0 0 6.697 6.697c.103.038.25.009.352-.126l.97-1.293a1.875 1.875 0 0 1 1.955-.694l4.423 1.105c.834.209 1.42.959 1.42 1.82V19.5a3 3 0 0 1-3 3h-2.25C8.552 22.5 1.5 15.448 1.5 6.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 3.5A1.5 1.5 0 0 1 3.5 2h1.148a1.5 1.5 0 0 1 1.465 1.175l.716 3.223a1.5 1.5 0 0 1-1.052 1.767l-.933.267c-.41.117-.643.555-.48.95a11.542 11.542 0 0 0 6.254 6.254c.395.163.833-.07.95-.48l.267-.933a1.5 1.5 0 0 1 1.767-1.052l3.223.716A1.5 1.5 0 0 1 18 15.352V16.5a1.5 1.5 0 0 1-1.5 1.5H15c-1.149 0-2.263-.15-3.326-.43A13.022 13.022 0 0 1 2.43 8.327A13.019 13.019 0 0 1 2 5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneXmark(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.75 3.75L18 6m0 0l2.25 2.25M18 6l2.25-2.25M18 6l-2.25 2.25m1.5 13.5c-8.284 0-15-6.716-15-15V4.5A2.25 2.25 0 0 1 4.5 2.25h1.372c.516 0 .966.351 1.091.852l1.106 4.423c.11.44-.054.902-.417 1.173l-1.293.97a1.062 1.062 0 0 0-.38 1.21a12.035 12.035 0 0 0 7.143 7.143c.441.162.928-.004 1.21-.38l.97-1.293a1.125 1.125 0 0 1 1.173-.417l4.423 1.106c.5.125.852.575.852 1.091V19.5a2.25 2.25 0 0 1-2.25 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneXmarkSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.855 7.286l1.067-.534a1 1 0 0 0 .542-1.046l-.44-2.858A1 1 0 0 0 4.036 2H3a1 1 0 0 0-1 1v2c0 .709.082 1.4.238 2.062a9.012 9.012 0 0 0 6.7 6.7A9.024 9.024 0 0 0 11 14h2a1 1 0 0 0 1-1v-1.036a1 1 0 0 0-.848-.988l-2.858-.44a1 1 0 0 0-1.046.542l-.534 1.067a7.52 7.52 0 0 1-4.86-4.859"/><path fill="currentColor" d="M13.78 2.22a.75.75 0 0 1 0 1.06L12.56 4.5l1.22 1.22a.75.75 0 0 1-1.06 1.06L11.5 5.56l-1.22 1.22a.75.75 0 1 1-1.06-1.06l1.22-1.22l-1.22-1.22a.75.75 0 0 1 1.06-1.06l1.22 1.22l1.22-1.22a.75.75 0 0 1 1.06 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneXmarkSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.22 3.22a.75.75 0 0 1 1.06 0L18 4.94l1.72-1.72a.75.75 0 1 1 1.06 1.06L19.06 6l1.72 1.72a.75.75 0 0 1-1.06 1.06L18 7.06l-1.72 1.72a.75.75 0 1 1-1.06-1.06L16.94 6l-1.72-1.72a.75.75 0 0 1 0-1.06M1.5 4.5a3 3 0 0 1 3-3h1.372c.86 0 1.61.586 1.819 1.42l1.105 4.423a1.875 1.875 0 0 1-.694 1.955l-1.293.97c-.135.101-.164.249-.126.352a11.285 11.285 0 0 0 6.697 6.697c.103.038.25.009.352-.126l.97-1.293a1.875 1.875 0 0 1 1.955-.694l4.423 1.105c.834.209 1.42.959 1.42 1.82V19.5a3 3 0 0 1-3 3h-2.25C8.552 22.5 1.5 15.448 1.5 6.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneXmarkTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 2A1.5 1.5 0 0 0 2 3.5V5c0 1.149.15 2.263.43 3.326a13.022 13.022 0 0 0 9.244 9.244c1.063.28 2.177.43 3.326.43h1.5a1.5 1.5 0 0 0 1.5-1.5v-1.148a1.5 1.5 0 0 0-1.175-1.465l-3.223-.716a1.5 1.5 0 0 0-1.767 1.052l-.267.933c-.117.41-.555.643-.95.48a11.542 11.542 0 0 1-6.254-6.254c-.163-.395.07-.833.48-.95l.933-.267a1.5 1.5 0 0 0 1.052-1.767l-.716-3.223A1.5 1.5 0 0 0 4.648 2zm9.78.22a.75.75 0 1 0-1.06 1.06L13.94 5l-1.72 1.72a.75.75 0 0 0 1.06 1.06L15 6.06l1.72 1.72a.75.75 0 1 0 1.06-1.06L16.06 5l1.72-1.72a.75.75 0 0 0-1.06-1.06L15 3.94z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Photo(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m2.25 15.75l5.159-5.159a2.25 2.25 0 0 1 3.182 0l5.159 5.159m-1.5-1.5l1.409-1.409a2.25 2.25 0 0 1 3.182 0l2.909 2.909m-18 3.75h16.5a1.5 1.5 0 0 0 1.5-1.5V6a1.5 1.5 0 0 0-1.5-1.5H3.75A1.5 1.5 0 0 0 2.25 6v12a1.5 1.5 0 0 0 1.5 1.5m10.5-11.25h.008v.008h-.008zm.375 0a.375.375 0 1 1-.75 0a.375.375 0 0 1 .75 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhotoSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2zm10.5 5.707a.5.5 0 0 0-.146-.353l-1-1a.5.5 0 0 0-.708 0L9.354 9.646a.5.5 0 0 1-.708 0L6.354 7.354a.5.5 0 0 0-.708 0l-2 2a.5.5 0 0 0-.146.353V12a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5zM12 5a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhotoSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 6a2.25 2.25 0 0 1 2.25-2.25h16.5A2.25 2.25 0 0 1 22.5 6v12a2.25 2.25 0 0 1-2.25 2.25H3.75A2.25 2.25 0 0 1 1.5 18zM3 16.06V18c0 .414.336.75.75.75h16.5A.75.75 0 0 0 21 18v-1.94l-2.69-2.689a1.5 1.5 0 0 0-2.12 0l-.88.879l.97.97a.75.75 0 1 1-1.06 1.06l-5.16-5.159a1.5 1.5 0 0 0-2.12 0zm10.125-7.81a1.125 1.125 0 1 1 2.25 0a1.125 1.125 0 0 1-2.25 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhotoTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 5.25A2.25 2.25 0 0 1 3.25 3h13.5A2.25 2.25 0 0 1 19 5.25v9.5A2.25 2.25 0 0 1 16.75 17H3.25A2.25 2.25 0 0 1 1 14.75zm1.5 5.81v3.69c0 .414.336.75.75.75h13.5a.75.75 0 0 0 .75-.75v-2.69l-2.22-2.219a.75.75 0 0 0-1.06 0l-1.91 1.909l.47.47a.75.75 0 1 1-1.06 1.06L6.53 8.091a.75.75 0 0 0-1.06 0zM12 7a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 5.653c0-.856.917-1.398 1.667-.986l11.54 6.347a1.125 1.125 0 0 1 0 1.972l-11.54 6.347a1.125 1.125 0 0 1-1.667-.986z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayCircle(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/><path d="M15.91 11.672a.375.375 0 0 1 0 .656l-5.603 3.113a.375.375 0 0 1-.557-.328V8.887c0-.286.307-.466.557-.327z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayCircleSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14m-.847-9.766A.75.75 0 0 0 6 5.866v4.268a.75.75 0 0 0 1.153.633l3.353-2.134a.75.75 0 0 0 0-1.266z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayCircleSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 12c0-5.385 4.365-9.75 9.75-9.75s9.75 4.365 9.75 9.75s-4.365 9.75-9.75 9.75S2.25 17.385 2.25 12m14.024-.983a1.125 1.125 0 0 1 0 1.967l-5.603 3.112A1.125 1.125 0 0 1 9 15.113V8.887c0-.857.922-1.4 1.671-.983z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayCircleTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 10a8 8 0 1 1 16 0a8 8 0 0 1-16 0m6.39-2.908a.75.75 0 0 1 .766.027l3.5 2.25a.75.75 0 0 1 0 1.262l-3.5 2.25A.75.75 0 0 1 8 12.25v-4.5a.75.75 0 0 1 .39-.658" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayPause(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 7.5V18M15 7.5V18M3 16.811V8.69c0-.864.933-1.406 1.683-.977l7.108 4.061a1.125 1.125 0 0 1 0 1.954l-7.108 4.061A1.125 1.125 0 0 1 3 16.811"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayPauseSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 4.804a1 1 0 0 1 1.53-.848l5.113 3.196a1 1 0 0 1 0 1.696L2.53 12.044A1 1 0 0 1 1 11.196zM13.5 4.5A.5.5 0 0 1 14 4h.5a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-.5.5H14a.5.5 0 0 1-.5-.5zm-3-.5a.5.5 0 0 0-.5.5v7a.5.5 0 0 0 .5.5h.5a.5.5 0 0 0 .5-.5v-7A.5.5 0 0 0 11 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayPauseSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 6.75a.75.75 0 0 0-.75.75V18a.75.75 0 0 0 .75.75h.75a.75.75 0 0 0 .75-.75V7.5a.75.75 0 0 0-.75-.75zm5.25 0a.75.75 0 0 0-.75.75V18c0 .414.336.75.75.75H21a.75.75 0 0 0 .75-.75V7.5a.75.75 0 0 0-.75-.75zm-15.195.31c-1.25-.713-2.805.19-2.805 1.63v8.122c0 1.44 1.555 2.343 2.805 1.628l7.108-4.061c1.26-.72 1.26-2.536 0-3.256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayPauseTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.75 4a.75.75 0 0 0-.75.75v10.5c0 .414.336.75.75.75h.5a.75.75 0 0 0 .75-.75V4.75a.75.75 0 0 0-.75-.75zm5 0a.75.75 0 0 0-.75.75v10.5c0 .414.336.75.75.75h.5a.75.75 0 0 0 .75-.75V4.75a.75.75 0 0 0-.75-.75zm-14.462.819A1.5 1.5 0 0 0 1 6.095v7.81a1.5 1.5 0 0 0 2.288 1.277l6.323-3.906a1.5 1.5 0 0 0 0-2.552z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaySixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3.732a1.5 1.5 0 0 1 2.305-1.265l6.706 4.267a1.5 1.5 0 0 1 0 2.531l-6.706 4.268A1.5 1.5 0 0 1 3 12.267z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 5.653c0-1.427 1.529-2.33 2.779-1.643l11.54 6.347c1.295.712 1.295 2.573 0 3.286L7.28 19.99c-1.25.687-2.779-.217-2.779-1.643z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.3 2.84A1.5 1.5 0 0 0 4 4.11v11.78a1.5 1.5 0 0 0 2.3 1.27l9.344-5.891a1.5 1.5 0 0 0 0-2.538z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 4.5v15m7.5-7.5h-15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircle(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 9v6m3-3H9m12 0a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircleSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14m.75-10.25v2.5h2.5a.75.75 0 0 1 0 1.5h-2.5v2.5a.75.75 0 0 1-1.5 0v-2.5h-2.5a.75.75 0 0 1 0-1.5h2.5v-2.5a.75.75 0 0 1 1.5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircleSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75s9.75-4.365 9.75-9.75S17.385 2.25 12 2.25M12.75 9a.75.75 0 0 0-1.5 0v2.25H9a.75.75 0 0 0 0 1.5h2.25V15a.75.75 0 0 0 1.5 0v-2.25H15a.75.75 0 0 0 0-1.5h-2.25z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircleTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 18a8 8 0 1 0 0-16a8 8 0 0 0 0 16m.75-11.25a.75.75 0 0 0-1.5 0v2.5h-2.5a.75.75 0 0 0 0 1.5h2.5v2.5a.75.75 0 0 0 1.5 0v-2.5h2.5a.75.75 0 0 0 0-1.5h-2.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.75 3.75a.75.75 0 0 0-1.5 0v3.5h-3.5a.75.75 0 0 0 0 1.5h3.5v3.5a.75.75 0 0 0 1.5 0v-3.5h3.5a.75.75 0 0 0 0-1.5h-3.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSmall(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 6v12m6-6H6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSmallSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 5.25a.75.75 0 0 1 .75.75v5.25H18a.75.75 0 0 1 0 1.5h-5.25V18a.75.75 0 0 1-1.5 0v-5.25H6a.75.75 0 0 1 0-1.5h5.25V6a.75.75 0 0 1 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSmallTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.75 6.75a.75.75 0 0 0-1.5 0v2.5h-2.5a.75.75 0 0 0 0 1.5h2.5v2.5a.75.75 0 0 0 1.5 0v-2.5h2.5a.75.75 0 0 0 0-1.5h-2.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 3.75a.75.75 0 0 1 .75.75v6.75h6.75a.75.75 0 0 1 0 1.5h-6.75v6.75a.75.75 0 0 1-1.5 0v-6.75H4.5a.75.75 0 0 1 0-1.5h6.75V4.5a.75.75 0 0 1 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.75 4.75a.75.75 0 0 0-1.5 0v4.5h-4.5a.75.75 0 0 0 0 1.5h4.5v4.5a.75.75 0 0 0 1.5 0v-4.5h4.5a.75.75 0 0 0 0-1.5h-4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Power(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.636 5.636a9 9 0 1 0 12.728 0M12 3v9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PowerSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1a.75.75 0 0 1 .75.75v6.5a.75.75 0 0 1-1.5 0v-6.5A.75.75 0 0 1 8 1M4.11 3.05a.75.75 0 0 1 0 1.06a5.5 5.5 0 1 0 7.78 0a.75.75 0 0 1 1.06-1.06a7 7 0 1 1-9.9 0a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PowerSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25a.75.75 0 0 1 .75.75v9a.75.75 0 0 1-1.5 0V3a.75.75 0 0 1 .75-.75M6.166 5.106a.75.75 0 0 1 0 1.06a8.25 8.25 0 1 0 11.668 0a.75.75 0 1 1 1.06-1.06c3.808 3.807 3.808 9.98 0 13.788c-3.807 3.808-9.98 3.808-13.788 0c-3.808-3.807-3.808-9.98 0-13.788a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PowerTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 2a.75.75 0 0 1 .75.75v7.5a.75.75 0 0 1-1.5 0v-7.5A.75.75 0 0 1 10 2M5.404 4.343a.75.75 0 0 1 0 1.06a6.5 6.5 0 1 0 9.192 0a.75.75 0 1 1 1.06-1.06a8 8 0 1 1-11.313 0a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PresentationChartBar(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 3v11.25A2.25 2.25 0 0 0 6 16.5h2.25M3.75 3h-1.5m1.5 0h16.5m0 0h1.5m-1.5 0v11.25A2.25 2.25 0 0 1 18 16.5h-2.25m-7.5 0h7.5m-7.5 0l-1 3m8.5-3l1 3m0 0l.5 1.5m-.5-1.5h-9.5m0 0l-.5 1.5M9 11.25v1.5M12 9v3.75m3-6v6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PresentationChartBarSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.75 2a.75.75 0 0 0 0 1.5H2V9a2 2 0 0 0 2 2h.043l-1.004 3.013a.75.75 0 0 0 1.423.474L4.624 14h6.752l.163.487a.75.75 0 1 0 1.422-.474L11.958 11H12a2 2 0 0 0 2-2V3.5h.25a.75.75 0 0 0 0-1.5zm8.626 9l.5 1.5H5.124l.5-1.5zM5.25 7a.75.75 0 0 0-.75.75v.5a.75.75 0 0 0 1.5 0v-.5A.75.75 0 0 0 5.25 7M10 4.75a.75.75 0 0 1 1.5 0v3.5a.75.75 0 0 1-1.5 0zM8 5.5a.75.75 0 0 0-.75.75v2a.75.75 0 0 0 1.5 0v-2A.75.75 0 0 0 8 5.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PresentationChartBarSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 2.25a.75.75 0 0 0 0 1.5H3v10.5a3 3 0 0 0 3 3h1.21l-1.172 3.513a.75.75 0 0 0 1.424.474l.329-.987h8.418l.33.987a.75.75 0 0 0 1.422-.474l-1.17-3.513H18a3 3 0 0 0 3-3V3.75h.75a.75.75 0 0 0 0-1.5zm6.04 16.5l.5-1.5h6.42l.5 1.5zm7.46-12a.75.75 0 0 0-1.5 0v6a.75.75 0 0 0 1.5 0zm-3 2.25a.75.75 0 0 0-1.5 0v3.75a.75.75 0 0 0 1.5 0zm-3 2.25a.75.75 0 0 0-1.5 0v1.5a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PresentationChartBarTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 2.75A.75.75 0 0 1 1.75 2h16.5a.75.75 0 0 1 0 1.5H18v8.75A2.75 2.75 0 0 1 15.25 15h-1.072l.798 3.06a.75.75 0 0 1-1.452.38L13.41 18H6.59l-.114.44a.75.75 0 0 1-1.452-.38L5.823 15H4.75A2.75 2.75 0 0 1 2 12.25V3.5h-.25A.75.75 0 0 1 1 2.75M7.373 15l-.391 1.5h6.037l-.392-1.5zM13.25 5a.75.75 0 0 1 .75.75v5.5a.75.75 0 0 1-1.5 0v-5.5a.75.75 0 0 1 .75-.75m-6.5 4a.75.75 0 0 1 .75.75v1.5a.75.75 0 0 1-1.5 0v-1.5A.75.75 0 0 1 6.75 9m4-1.25a.75.75 0 0 0-1.5 0v3.5a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PresentationChartLine(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 3v11.25A2.25 2.25 0 0 0 6 16.5h2.25M3.75 3h-1.5m1.5 0h16.5m0 0h1.5m-1.5 0v11.25A2.25 2.25 0 0 1 18 16.5h-2.25m-7.5 0h7.5m-7.5 0l-1 3m8.5-3l1 3m0 0l.5 1.5m-.5-1.5h-9.5m0 0l-.5 1.5m.75-9l3-3l2.148 2.148A12.061 12.061 0 0 1 16.5 7.605"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PresentationChartLineSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.75 2a.75.75 0 0 0 0 1.5H2V9a2 2 0 0 0 2 2h.043l-1.005 3.013a.75.75 0 0 0 1.423.474L4.624 14h6.752l.163.487a.75.75 0 0 0 1.423-.474L11.957 11H12a2 2 0 0 0 2-2V3.5h.25a.75.75 0 0 0 0-1.5zm8.626 9l.5 1.5H5.124l.5-1.5zm1.317-5.833a.75.75 0 0 0-.892-1.206a8.789 8.789 0 0 0-2.465 2.814L7.28 5.72a.75.75 0 0 0-1.06 0l-2 2a.75.75 0 0 0 1.06 1.06l1.47-1.47l1.278 1.28a.75.75 0 0 0 1.228-.255a7.275 7.275 0 0 1 2.437-3.167" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PresentationChartLineSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 2.25a.75.75 0 0 0 0 1.5H3v10.5a3 3 0 0 0 3 3h1.21l-1.172 3.513a.75.75 0 0 0 1.424.474l.329-.987h8.418l.33.987a.75.75 0 0 0 1.422-.474l-1.17-3.513H18a3 3 0 0 0 3-3V3.75h.75a.75.75 0 0 0 0-1.5zm6.54 15h6.42l.5 1.5H8.29zm8.086-8.995a.75.75 0 1 0-.751-1.299a12.81 12.81 0 0 0-3.558 3.05L11.03 8.47a.75.75 0 0 0-1.06 0l-3 3a.75.75 0 1 0 1.06 1.06l2.47-2.47l1.617 1.618a.75.75 0 0 0 1.146-.102a11.312 11.312 0 0 1 3.612-3.321" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PresentationChartLineTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 2.75A.75.75 0 0 1 1.75 2h16.5a.75.75 0 0 1 0 1.5H18v8.75A2.75 2.75 0 0 1 15.25 15h-1.072l.798 3.06a.75.75 0 0 1-1.452.38L13.41 18H6.59l-.114.44a.75.75 0 0 1-1.452-.38L5.823 15H4.75A2.75 2.75 0 0 1 2 12.25V3.5h-.25A.75.75 0 0 1 1 2.75M7.373 15l-.391 1.5h6.037l-.392-1.5zm7.49-8.931a.75.75 0 0 1-.175 1.046a19.326 19.326 0 0 0-3.398 3.098a.75.75 0 0 1-1.097.04L8.5 8.561l-2.22 2.22A.75.75 0 1 1 5.22 9.72l2.75-2.75a.75.75 0 0 1 1.06 0l1.664 1.663a20.786 20.786 0 0 1 3.122-2.74a.75.75 0 0 1 1.047.176" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Printer(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.72 13.829c-.24.03-.48.062-.72.096m.72-.096a42.415 42.415 0 0 1 10.56 0m-10.56 0L6.34 18m10.94-4.171c.24.03.48.062.72.096m-.72-.096L17.66 18m0 0l.229 2.523a1.125 1.125 0 0 1-1.12 1.227H7.231c-.662 0-1.18-.568-1.12-1.227L6.34 18m11.318 0h1.091A2.25 2.25 0 0 0 21 15.75V9.456c0-1.081-.768-2.015-1.837-2.175a48.055 48.055 0 0 0-1.913-.247M6.34 18H5.25A2.25 2.25 0 0 1 3 15.75V9.456c0-1.081.768-2.015 1.837-2.175a48.041 48.041 0 0 1 1.913-.247m10.5 0a48.536 48.536 0 0 0-10.5 0m10.5 0V3.375c0-.621-.504-1.125-1.125-1.125h-8.25c-.621 0-1.125.504-1.125 1.125v3.659M18 10.5h.008v.008H18zm-3 0h.008v.008H15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PrinterSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 5a2 2 0 0 0-2 2v3a2 2 0 0 0 1.51 1.94l-.315 1.896A1 1 0 0 0 4.18 15h7.639a1 1 0 0 0 .986-1.164l-.316-1.897A2 2 0 0 0 14 10V7a2 2 0 0 0-2-2V2a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1zm1.5 0V2.5h5V5zm5.23 5.5H5.27l-.5 3h6.459z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PrinterSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.875 1.5C6.839 1.5 6 2.34 6 3.375v2.99c-.426.053-.851.11-1.274.174c-1.454.218-2.476 1.483-2.476 2.917v6.294a3 3 0 0 0 3 3h.27l-.155 1.705A1.875 1.875 0 0 0 7.232 22.5h9.536a1.875 1.875 0 0 0 1.867-2.045l-.155-1.705h.27a3 3 0 0 0 3-3V9.456c0-1.434-1.022-2.7-2.476-2.917A48.716 48.716 0 0 0 18 6.366V3.375c0-1.036-.84-1.875-1.875-1.875zM16.5 6.205v-2.83A.375.375 0 0 0 16.125 3h-8.25a.375.375 0 0 0-.375.375v2.83a49.353 49.353 0 0 1 9 0m-.217 8.265c.178.018.317.16.333.337l.526 5.784a.375.375 0 0 1-.374.409H7.232a.375.375 0 0 1-.374-.409l.526-5.784a.373.373 0 0 1 .333-.337a41.741 41.741 0 0 1 8.566 0m.967-3.97a.75.75 0 0 1 .75-.75h.008a.75.75 0 0 1 .75.75v.008a.75.75 0 0 1-.75.75H18a.75.75 0 0 1-.75-.75zM15 9.75a.75.75 0 0 0-.75.75v.008c0 .414.336.75.75.75h.008a.75.75 0 0 0 .75-.75V10.5a.75.75 0 0 0-.75-.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PrinterTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 2.75C5 1.784 5.784 1 6.75 1h6.5c.966 0 1.75.784 1.75 1.75v3.552c.377.046.752.097 1.126.153A2.212 2.212 0 0 1 18 8.653v4.097A2.25 2.25 0 0 1 15.75 15h-.241l.305 1.984A1.75 1.75 0 0 1 14.084 19H5.915a1.75 1.75 0 0 1-1.73-2.016L4.492 15H4.25A2.25 2.25 0 0 1 2 12.75V8.653c0-1.082.775-2.034 1.874-2.198c.374-.056.75-.107 1.127-.153L5 6.25zm8.5 3.397a41.533 41.533 0 0 0-7 0V2.75a.25.25 0 0 1 .25-.25h6.5a.25.25 0 0 1 .25.25zM6.608 12.5a.25.25 0 0 0-.247.212l-.693 4.5a.25.25 0 0 0 .247.288h8.17a.25.25 0 0 0 .246-.288l-.692-4.5a.25.25 0 0 0-.247-.212z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PuzzlePiece(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.25 6.087c0-.355.186-.676.401-.959c.221-.29.349-.634.349-1.003c0-1.036-1.007-1.875-2.25-1.875s-2.25.84-2.25 1.875c0 .369.128.713.349 1.003c.215.283.401.604.401.959v0a.64.64 0 0 1-.657.643a48.39 48.39 0 0 1-4.163-.3a48.44 48.44 0 0 1 .315 4.907a.656.656 0 0 1-.658.663v0c-.355 0-.676-.186-.959-.401a1.647 1.647 0 0 0-1.003-.349c-1.036 0-1.875 1.007-1.875 2.25s.84 2.25 1.875 2.25c.369 0 .713-.128 1.003-.349c.283-.215.604-.401.959-.401v0c.31 0 .555.26.532.57a48.039 48.039 0 0 1-.642 5.056c1.518.19 3.058.309 4.616.354a.64.64 0 0 0 .657-.643v0c0-.355-.186-.676-.401-.959a1.647 1.647 0 0 1-.349-1.003c0-1.035 1.008-1.875 2.25-1.875c1.243 0 2.25.84 2.25 1.875c0 .369-.128.713-.349 1.003c-.215.283-.4.604-.4.959v0c0 .333.277.599.61.58a48.1 48.1 0 0 0 5.427-.63a48.05 48.05 0 0 0 .582-4.717a.532.532 0 0 0-.533-.57v0c-.355 0-.676.186-.959.401c-.29.221-.634.349-1.003.349c-1.035 0-1.875-1.007-1.875-2.25s.84-2.25 1.875-2.25c.37 0 .713.128 1.003.349c.283.215.604.401.96.401v0a.656.656 0 0 0 .658-.663a48.422 48.422 0 0 0-.37-5.36a48.14 48.14 0 0 1-5.766.689a.578.578 0 0 1-.61-.58"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PuzzlePieceSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 3.889c0-.273.188-.502.417-.65c.355-.229.583-.587.583-.989C10 1.56 9.328 1 8.5 1S7 1.56 7 2.25c0 .41.237.774.603 1.002c.22.137.397.355.397.613a.577.577 0 0 1-.605.579c-.744-.04-1.482-.1-2.214-.18a.75.75 0 0 0-.83.81c.067.764.111 1.535.133 2.312A.6.6 0 0 1 3.882 8c-.268 0-.495-.185-.64-.412C3.015 7.231 2.655 7 2.25 7C1.56 7 1 7.672 1 8.5S1.56 10 2.25 10c.404 0 .764-.23.993-.588c.144-.227.37-.412.64-.412a.6.6 0 0 1 .601.614a39.338 39.338 0 0 1-.231 3.3a.75.75 0 0 0 .661.829c.826.093 1.66.161 2.5.204A.56.56 0 0 0 8 13.386c0-.271-.187-.499-.415-.645C7.23 12.512 7 12.153 7 11.75c0-.69.672-1.25 1.5-1.25s1.5.56 1.5 1.25c0 .403-.23.762-.585.99c-.228.147-.415.375-.415.646v.11c0 .278.223.504.5.504c1.196 0 2.381-.052 3.552-.154a.75.75 0 0 0 .68-.661c.135-1.177.22-2.37.253-3.574a.597.597 0 0 0-.6-.611c-.27 0-.498.187-.644.415c-.229.356-.588.585-.991.585c-.69 0-1.25-.672-1.25-1.5S11.06 7 11.75 7c.403 0 .762.23.99.585c.147.228.375.415.646.415a.597.597 0 0 0 .599-.61a40.914 40.914 0 0 0-.132-2.365a.75.75 0 0 0-.815-.684A39.51 39.51 0 0 1 9.5 4.5a.501.501 0 0 1-.5-.503z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PuzzlePieceSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.25 5.337c0-.355-.186-.676-.401-.959a1.647 1.647 0 0 1-.349-1.003c0-1.036 1.007-1.875 2.25-1.875S15 2.34 15 3.375c0 .369-.128.713-.349 1.003c-.215.283-.401.604-.401.959c0 .332.278.598.61.578c1.91-.114 3.79-.342 5.632-.676a.75.75 0 0 1 .878.645a49.17 49.17 0 0 1 .376 5.452a.657.657 0 0 1-.66.664c-.354 0-.675-.186-.958-.401a1.647 1.647 0 0 0-1.003-.349c-1.035 0-1.875 1.007-1.875 2.25s.84 2.25 1.875 2.25c.369 0 .713-.128 1.003-.349c.283-.215.604-.401.959-.401c.31 0 .557.262.534.571a48.774 48.774 0 0 1-.595 4.845a.75.75 0 0 1-.61.61c-1.82.317-3.673.533-5.555.642a.58.58 0 0 1-.611-.581c0-.355.186-.676.401-.959c.221-.29.349-.634.349-1.003c0-1.035-1.007-1.875-2.25-1.875s-2.25.84-2.25 1.875c0 .369.128.713.349 1.003c.215.283.401.604.401.959a.641.641 0 0 1-.658.643a49.118 49.118 0 0 1-4.708-.36a.75.75 0 0 1-.645-.878c.293-1.614.504-3.257.629-4.924A.53.53 0 0 0 5.337 15c-.355 0-.676.186-.959.401c-.29.221-.634.349-1.003.349c-1.036 0-1.875-1.007-1.875-2.25s.84-2.25 1.875-2.25c.369 0 .713.128 1.003.349c.283.215.604.401.959.401a.656.656 0 0 0 .659-.663a47.703 47.703 0 0 0-.31-4.82a.75.75 0 0 1 .83-.832c1.343.155 2.703.254 4.077.294a.64.64 0 0 0 .657-.642"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PuzzlePieceTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 4.467c0-.405.262-.75.559-1.027c.276-.257.441-.584.441-.94c0-.828-.895-1.5-2-1.5s-2 .672-2 1.5c0 .362.171.694.456.953c.29.265.544.6.544.994a.968.968 0 0 1-1.024.974a39.655 39.655 0 0 1-3.014-.306a.75.75 0 0 0-.847.847c.14.993.242 1.999.306 3.014A.968.968 0 0 1 4.447 10c-.393 0-.729-.253-.994-.544C3.194 9.17 2.862 9 2.5 9C1.672 9 1 9.895 1 11s.672 2 1.5 2c.356 0 .683-.165.94-.441c.276-.297.622-.559 1.027-.559a.997.997 0 0 1 1.004 1.03a39.747 39.747 0 0 1-.319 3.734a.75.75 0 0 0 .64.842c1.05.146 2.111.252 3.184.318A.97.97 0 0 0 10 16.948c0-.394-.254-.73-.545-.995C9.171 15.693 9 15.362 9 15c0-.828.895-1.5 2-1.5s2 .672 2 1.5c0 .356-.165.683-.441.94c-.297.276-.559.622-.559 1.027a.998.998 0 0 0 1.03 1.005a41.333 41.333 0 0 0 3.961-.337a.75.75 0 0 0 .644-.644c.175-1.302.288-2.624.337-3.961A.998.998 0 0 0 16.967 12c-.405 0-.75.262-1.027.559c-.257.276-.584.441-.94.441c-.828 0-1.5-.895-1.5-2s.672-2 1.5-2c.362 0 .694.17.953.455c.265.291.601.545.995.545a.97.97 0 0 0 .976-1.024a41.159 41.159 0 0 0-.318-3.184a.75.75 0 0 0-.842-.64a39.565 39.565 0 0 1-3.734.319A.997.997 0 0 1 12 4.467"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QrCode(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.75 4.875c0-.621.504-1.125 1.125-1.125h4.5c.621 0 1.125.504 1.125 1.125v4.5c0 .621-.504 1.125-1.125 1.125h-4.5A1.125 1.125 0 0 1 3.75 9.375zm0 9.75c0-.621.504-1.125 1.125-1.125h4.5c.621 0 1.125.504 1.125 1.125v4.5c0 .621-.504 1.125-1.125 1.125h-4.5a1.125 1.125 0 0 1-1.125-1.125zm9.75-9.75c0-.621.504-1.125 1.125-1.125h4.5c.621 0 1.125.504 1.125 1.125v4.5c0 .621-.504 1.125-1.125 1.125h-4.5A1.125 1.125 0 0 1 13.5 9.375z"/><path d="M6.75 6.75h.75v.75h-.75zm0 9.75h.75v.75h-.75zm9.75-9.75h.75v.75h-.75zm-3 6.75h.75v.75h-.75zm0 6h.75v.75h-.75zm6-6h.75v.75h-.75zm0 6h.75v.75h-.75zm-3-3h.75v.75h-.75z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QrCodeSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M4.75 4.25a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1"/><path fill-rule="evenodd" d="M2 3.5A1.5 1.5 0 0 1 3.5 2H6a1.5 1.5 0 0 1 1.5 1.5V6A1.5 1.5 0 0 1 6 7.5H3.5A1.5 1.5 0 0 1 2 6zm1.5 0H6V6H3.5z" clip-rule="evenodd"/><path d="M4.25 11.25a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0"/><path fill-rule="evenodd" d="M2 10a1.5 1.5 0 0 1 1.5-1.5H6A1.5 1.5 0 0 1 7.5 10v2.5A1.5 1.5 0 0 1 6 14H3.5A1.5 1.5 0 0 1 2 12.5zm1.5 2.5V10H6v2.5z" clip-rule="evenodd"/><path d="M11.25 4.25a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1"/><path fill-rule="evenodd" d="M10 2a1.5 1.5 0 0 0-1.5 1.5V6A1.5 1.5 0 0 0 10 7.5h2.5A1.5 1.5 0 0 0 14 6V3.5A1.5 1.5 0 0 0 12.5 2zm2.5 1.5H10V6h2.5z" clip-rule="evenodd"/><path d="M8.5 9.417a.917.917 0 1 1 1.833 0a.917.917 0 0 1-1.833 0m0 3.666a.917.917 0 1 1 1.833 0a.917.917 0 0 1-1.833 0M13.083 8.5a.917.917 0 1 0 0 1.833a.917.917 0 0 0 0-1.833m-.917 4.584a.917.917 0 1 1 1.833 0a.917.917 0 0 1-1.833 0m-.916-2.751a.917.917 0 1 0 0 1.833a.917.917 0 0 0 0-1.833"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QrCodeSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.875C3 3.839 3.84 3 4.875 3h4.5c1.036 0 1.875.84 1.875 1.875v4.5c0 1.036-.84 1.875-1.875 1.875h-4.5A1.875 1.875 0 0 1 3 9.375zM4.875 4.5a.375.375 0 0 0-.375.375v4.5c0 .207.168.375.375.375h4.5a.375.375 0 0 0 .375-.375v-4.5a.375.375 0 0 0-.375-.375zm7.875.375c0-1.036.84-1.875 1.875-1.875h4.5C20.16 3 21 3.84 21 4.875v4.5c0 1.036-.84 1.875-1.875 1.875h-4.5a1.875 1.875 0 0 1-1.875-1.875zm1.875-.375a.375.375 0 0 0-.375.375v4.5c0 .207.168.375.375.375h4.5a.375.375 0 0 0 .375-.375v-4.5a.375.375 0 0 0-.375-.375zM6 6.75A.75.75 0 0 1 6.75 6h.75a.75.75 0 0 1 .75.75v.75a.75.75 0 0 1-.75.75h-.75A.75.75 0 0 1 6 7.5zm9.75 0A.75.75 0 0 1 16.5 6h.75a.75.75 0 0 1 .75.75v.75a.75.75 0 0 1-.75.75h-.75a.75.75 0 0 1-.75-.75zM3 14.625c0-1.036.84-1.875 1.875-1.875h4.5c1.036 0 1.875.84 1.875 1.875v4.5c0 1.035-.84 1.875-1.875 1.875h-4.5A1.875 1.875 0 0 1 3 19.125zm1.875-.375a.375.375 0 0 0-.375.375v4.5c0 .207.168.375.375.375h4.5a.375.375 0 0 0 .375-.375v-4.5a.375.375 0 0 0-.375-.375zm7.875-.75a.75.75 0 0 1 .75-.75h.75a.75.75 0 0 1 .75.75v.75a.75.75 0 0 1-.75.75h-.75a.75.75 0 0 1-.75-.75zm6 0a.75.75 0 0 1 .75-.75h.75a.75.75 0 0 1 .75.75v.75a.75.75 0 0 1-.75.75h-.75a.75.75 0 0 1-.75-.75zM6 16.5a.75.75 0 0 1 .75-.75h.75a.75.75 0 0 1 .75.75v.75a.75.75 0 0 1-.75.75h-.75a.75.75 0 0 1-.75-.75zm9.75 0a.75.75 0 0 1 .75-.75h.75a.75.75 0 0 1 .75.75v.75a.75.75 0 0 1-.75.75h-.75a.75.75 0 0 1-.75-.75zm-3 3a.75.75 0 0 1 .75-.75h.75a.75.75 0 0 1 .75.75v.75a.75.75 0 0 1-.75.75h-.75a.75.75 0 0 1-.75-.75zm6 0a.75.75 0 0 1 .75-.75h.75a.75.75 0 0 1 .75.75v.75a.75.75 0 0 1-.75.75h-.75a.75.75 0 0 1-.75-.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QrCodeTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.75 2A1.75 1.75 0 0 0 2 3.75v3.5C2 8.216 2.784 9 3.75 9h3.5A1.75 1.75 0 0 0 9 7.25v-3.5A1.75 1.75 0 0 0 7.25 2zM3.5 3.75a.25.25 0 0 1 .25-.25h3.5a.25.25 0 0 1 .25.25v3.5a.25.25 0 0 1-.25.25h-3.5a.25.25 0 0 1-.25-.25zM3.75 11A1.75 1.75 0 0 0 2 12.75v3.5c0 .966.784 1.75 1.75 1.75h3.5A1.75 1.75 0 0 0 9 16.25v-3.5A1.75 1.75 0 0 0 7.25 11zm-.25 1.75a.25.25 0 0 1 .25-.25h3.5a.25.25 0 0 1 .25.25v3.5a.25.25 0 0 1-.25.25h-3.5a.25.25 0 0 1-.25-.25zm7.5-9c0-.966.784-1.75 1.75-1.75h3.5c.966 0 1.75.784 1.75 1.75v3.5A1.75 1.75 0 0 1 16.25 9h-3.5A1.75 1.75 0 0 1 11 7.25zm1.75-.25a.25.25 0 0 0-.25.25v3.5c0 .138.112.25.25.25h3.5a.25.25 0 0 0 .25-.25v-3.5a.25.25 0 0 0-.25-.25zm-7.26 1a1 1 0 0 0-1 1v.01a1 1 0 0 0 1 1h.01a1 1 0 0 0 1-1V5.5a1 1 0 0 0-1-1zm9 0a1 1 0 0 0-1 1v.01a1 1 0 0 0 1 1h.01a1 1 0 0 0 1-1V5.5a1 1 0 0 0-1-1zm-9 9a1 1 0 0 0-1 1v.01a1 1 0 0 0 1 1h.01a1 1 0 0 0 1-1v-.01a1 1 0 0 0-1-1zm9 0a1 1 0 0 0-1 1v.01a1 1 0 0 0 1 1h.01a1 1 0 0 0 1-1v-.01a1 1 0 0 0-1-1zm-3.5-1.5a1 1 0 0 1 1-1H12a1 1 0 0 1 1 1v.01a1 1 0 0 1-1 1h-.01a1 1 0 0 1-1-1zm6-1a1 1 0 0 0-1 1v.01a1 1 0 0 0 1 1H17a1 1 0 0 0 1-1V12a1 1 0 0 0-1-1zm-1 6a1 1 0 0 1 1-1H17a1 1 0 0 1 1 1v.01a1 1 0 0 1-1 1h-.01a1 1 0 0 1-1-1zm-4-1a1 1 0 0 0-1 1v.01a1 1 0 0 0 1 1H12a1 1 0 0 0 1-1V17a1 1 0 0 0-1-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionMarkCircle(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.879 7.519c1.172-1.025 3.071-1.025 4.243 0c1.171 1.025 1.171 2.687 0 3.712a2.98 2.98 0 0 1-.67.442c-.746.361-1.452.999-1.452 1.827v.75M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0m-9 5.25h.008v.008H12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionMarkCircleSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0m-6 3.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0M7.293 5.293a1 1 0 1 1 .99 1.667c-.459.134-1.033.566-1.033 1.29v.25a.75.75 0 1 0 1.5 0v-.115a2.5 2.5 0 1 0-2.518-4.153a.75.75 0 1 0 1.061 1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionMarkCircleSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 12c0-5.385 4.365-9.75 9.75-9.75s9.75 4.365 9.75 9.75s-4.365 9.75-9.75 9.75S2.25 17.385 2.25 12m11.378-3.917c-.89-.777-2.366-.777-3.255 0a.75.75 0 0 1-.988-1.129c1.454-1.272 3.776-1.272 5.23 0c1.513 1.324 1.513 3.518 0 4.842a3.75 3.75 0 0 1-.837.552c-.676.328-1.028.774-1.028 1.152v.75a.75.75 0 0 1-1.5 0v-.75c0-1.279 1.06-2.107 1.875-2.502a2.28 2.28 0 0 0 .503-.331c.83-.727.83-1.857 0-2.584M12 18a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionMarkCircleTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M18 10a8 8 0 1 1-16 0a8 8 0 0 1 16 0M8.94 6.94a.75.75 0 1 1-1.061-1.061a3 3 0 1 1 2.871 5.026v.345a.75.75 0 0 1-1.5 0v-.5c0-.72.57-1.172 1.081-1.287A1.5 1.5 0 1 0 8.94 6.94M10 15a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QueueList(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 12h16.5m-16.5 3.75h16.5M3.75 19.5h16.5M5.625 4.5h12.75a1.875 1.875 0 0 1 0 3.75H5.625a1.875 1.875 0 0 1 0-3.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QueueListSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 4a2 2 0 0 1 2-2h8a2 2 0 1 1 0 4H4a2 2 0 0 1-2-2m0 5.25a.75.75 0 0 1 .75-.75h10.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 9.25m.75 3.25a.75.75 0 0 0 0 1.5h10.5a.75.75 0 0 0 0-1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QueueListSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.625 3.75a2.625 2.625 0 1 0 0 5.25h12.75a2.625 2.625 0 0 0 0-5.25zm-1.875 7.5a.75.75 0 0 0 0 1.5h16.5a.75.75 0 0 0 0-1.5zM3 15.75a.75.75 0 0 1 .75-.75h16.5a.75.75 0 0 1 0 1.5H3.75a.75.75 0 0 1-.75-.75m.75 3a.75.75 0 0 0 0 1.5h16.5a.75.75 0 0 0 0-1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QueueListTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 4.5A2.5 2.5 0 0 1 4.5 2h11a2.5 2.5 0 0 1 0 5h-11A2.5 2.5 0 0 1 2 4.5m.75 4.583a.75.75 0 0 0 0 1.5h14.5a.75.75 0 0 0 0-1.5zm0 3.58a.75.75 0 0 0 0 1.5h14.5a.75.75 0 0 0 0-1.5zm0 3.587a.75.75 0 0 0 0 1.5h14.5a.75.75 0 1 0 0-1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Radio(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3.75 7.5l16.5-4.125M12 6.75a48.3 48.3 0 0 0-7.948.655C2.999 7.58 2.25 8.507 2.25 9.574v9.176A2.25 2.25 0 0 0 4.5 21h15a2.25 2.25 0 0 0 2.25-2.25V9.574c0-1.067-.75-1.994-1.802-2.169A48.329 48.329 0 0 0 12 6.75m-1.683 6.443l-.005.005l-.006-.005l.006-.005zm-.005 2.127l-.005-.006l.005-.005l.005.005zm-2.116-.006l-.005.006l-.006-.006l.005-.005zm-.005-2.116l-.006-.005l.006-.005l.005.005zM9.255 10.5v.008h-.008V10.5zm3.249 1.88l-.007.004l-.003-.007l.006-.003zm-1.38 5.126l-.003-.006l.006-.004l.004.007zm.007-6.501l-.003.006l-.007-.003l.004-.007zm1.37 5.129l-.007-.004l.004-.006l.006.003zm.504-1.877h-.008v-.007h.008zM9.255 18v.008h-.008V18zm-3.246-1.87l-.007.004L6 16.127l.006-.003zm1.366-5.119l-.004-.006l.006-.004l.004.007zM7.38 17.5l-.003.006l-.007-.003l.004-.007zm-1.376-5.116L6 12.38l.003-.007l.007.004zm-.5 1.873h-.008v-.007h.008zM17.25 12.75a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5m0 4.5a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadioSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.437 1.45a.75.75 0 0 1-.386.987L7.478 4H13a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h.736l6.713-2.937a.75.75 0 0 1 .988.386M12 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2M6.75 6.75a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0m-.75 3a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5m2.323-1.225a.75.75 0 1 1-.75-1.3a.75.75 0 0 1 .75 1.3M7.3 9.75a.75.75 0 1 0 1.299.75a.75.75 0 0 0-1.3-.75m-.549 1.5a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0m-3.348-.75a.75.75 0 1 0 1.3-.75a.75.75 0 0 0-1.3.75m.275-1.975a.75.75 0 1 1 .75-1.3a.75.75 0 0 1-.75 1.3M12 12a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadioSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M20.432 4.103a.75.75 0 0 0-.364-1.456L4.128 6.632l-.2.033C2.498 6.904 1.5 8.158 1.5 9.574v9.176a3 3 0 0 0 3 3h15a3 3 0 0 0 3-3V9.574c0-1.416-.997-2.67-2.429-2.909a49.017 49.017 0 0 0-7.255-.658zm-9.585 8.56a.75.75 0 0 1 0 1.06l-.005.006a.75.75 0 0 1-1.06 0l-.006-.006a.75.75 0 0 1 0-1.06l.005-.005a.75.75 0 0 1 1.06 0zM9.781 15.85a.75.75 0 0 0 1.061 0l.005-.005a.75.75 0 0 0 0-1.061l-.005-.005a.75.75 0 0 0-1.06 0l-.006.005a.75.75 0 0 0 0 1.06zm-1.055-1.066a.75.75 0 0 1 0 1.06l-.005.006a.75.75 0 0 1-1.061 0l-.005-.005a.75.75 0 0 1 0-1.06l.005-.006a.75.75 0 0 1 1.06 0zM7.66 13.73a.75.75 0 0 0 1.061 0l.005-.006a.75.75 0 0 0 0-1.06l-.005-.006a.75.75 0 0 0-1.06 0l-.006.006a.75.75 0 0 0 0 1.06zm1.595-3.98a.75.75 0 0 1 .75.75v.008a.75.75 0 0 1-.75.75h-.008a.75.75 0 0 1-.75-.75V10.5a.75.75 0 0 1 .75-.75zm3.624 3.28a.75.75 0 0 0 .275-1.025L13.15 12a.75.75 0 0 0-1.025-.275l-.006.004a.75.75 0 0 0-.275 1.024l.004.007a.75.75 0 0 0 1.025.274zm-1.38 5.126a.75.75 0 0 1-1.024-.275l-.004-.006a.75.75 0 0 1 .275-1.025l.006-.004a.75.75 0 0 1 1.025.275l.004.007a.75.75 0 0 1-.275 1.024zm.282-6.776a.75.75 0 0 0-.274-1.025l-.007-.003a.75.75 0 0 0-1.024.274l-.004.007a.75.75 0 0 0 .274 1.024l.007.004a.75.75 0 0 0 1.024-.275zm1.369 5.129a.75.75 0 0 1-1.025.274l-.006-.004a.75.75 0 0 1-.275-1.024l.004-.007a.75.75 0 0 1 1.025-.274l.006.004a.75.75 0 0 1 .275 1.024zm-.145-1.502a.75.75 0 0 0 .75-.75v-.007a.75.75 0 0 0-.75-.75h-.008a.75.75 0 0 0-.75.75v.007c0 .415.336.75.75.75zm-3.75 2.243a.75.75 0 0 1 .75.75v.008a.75.75 0 0 1-.75.75h-.008a.75.75 0 0 1-.75-.75V18a.75.75 0 0 1 .75-.75zm-2.871-.47a.75.75 0 0 0 .274-1.025l-.003-.006a.75.75 0 0 0-1.025-.275l-.006.004a.75.75 0 0 0-.275 1.024l.004.007a.75.75 0 0 0 1.024.274zm1.366-5.12a.75.75 0 0 1-1.025-.274l-.004-.006a.75.75 0 0 1 .275-1.025l.006-.004a.75.75 0 0 1 1.025.275l.004.006a.75.75 0 0 1-.275 1.025zm.281 6.215a.75.75 0 0 0-.275-1.024l-.006-.004a.75.75 0 0 0-1.025.274l-.003.007a.75.75 0 0 0 .274 1.024l.007.004a.75.75 0 0 0 1.024-.274zm-1.376-5.116a.75.75 0 0 1-1.025.274l-.006-.003a.75.75 0 0 1-.275-1.025l.004-.007a.75.75 0 0 1 1.025-.274l.006.004a.75.75 0 0 1 .275 1.024zm-1.15 2.248a.75.75 0 0 0 .75-.75v-.007a.75.75 0 0 0-.75-.75h-.008a.75.75 0 0 0-.75.75v.007c0 .415.336.75.75.75zM17.25 10.5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3m1.5 6a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadioTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M17.45 3.473a.75.75 0 1 0-.4-1.446L5.313 5.265c-.84.096-1.671.217-2.495.362A2.212 2.212 0 0 0 1 7.816v7.934A2.25 2.25 0 0 0 3.25 18h13.5A2.25 2.25 0 0 0 19 15.75V7.816c0-1.06-.745-2-1.817-2.189a41.12 41.12 0 0 0-5.406-.59zM16 9.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0M14.5 16a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m-9.26-5a.75.75 0 0 1 .75-.75H6a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75h-.01a.75.75 0 0 1-.75-.75zm2.75-.75a.75.75 0 0 0-.75.75v.01c0 .414.336.75.75.75H8a.75.75 0 0 0 .75-.75V11a.75.75 0 0 0-.75-.75zm-1.75-1.5A.75.75 0 0 1 6.99 8H7a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75h-.01a.75.75 0 0 1-.75-.75zm3.583.42a.75.75 0 0 0-1.06 0l-.007.007a.75.75 0 0 0 0 1.06l.007.007a.75.75 0 0 0 1.06 0l.007-.007a.75.75 0 0 0 0-1.06zm.427 2.08A.75.75 0 0 1 11 12v.01a.75.75 0 0 1-.75.75h-.01a.75.75 0 0 1-.75-.75V12a.75.75 0 0 1 .75-.75zm-.42 3.583a.75.75 0 0 0 0-1.06l-.007-.007a.75.75 0 0 0-1.06 0l-.007.007a.75.75 0 0 0 0 1.06l.007.008a.75.75 0 0 0 1.06 0zm-3.59.417a.75.75 0 0 1 .75-.75H7a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75h-.01a.75.75 0 0 1-.75-.75zm-1.013-1.484a.75.75 0 0 0-1.06 0l-.008.007a.75.75 0 0 0 0 1.06l.007.008a.75.75 0 0 0 1.061 0l.007-.008a.75.75 0 0 0 0-1.06zM3.75 11.25a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75h-.01a.75.75 0 0 1-.75-.75V12a.75.75 0 0 1 .75-.75zm1.484-1.013a.75.75 0 0 0 0-1.06l-.007-.007a.75.75 0 0 0-1.06 0l-.007.007a.75.75 0 0 0 0 1.06l.007.007a.75.75 0 0 0 1.06 0zM7.24 13a.75.75 0 0 1 .75-.75H8a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75h-.01a.75.75 0 0 1-.75-.75zm-1.25-.75a.75.75 0 0 0-.75.75v.01c0 .414.336.75.75.75H6a.75.75 0 0 0 .75-.75V13a.75.75 0 0 0-.75-.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReceiptPercent(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 14.25l6-6m4.5-3.493V21.75l-3.75-1.5l-3.75 1.5l-3.75-1.5l-3.75 1.5V4.757c0-1.108.806-2.057 1.907-2.185a48.507 48.507 0 0 1 11.186 0c1.1.128 1.907 1.077 1.907 2.185M9.75 9h.008v.008H9.75zm.375 0a.375.375 0 1 1-.75 0a.375.375 0 0 1 .75 0m4.125 4.5h.008v.008h-.008zm.375 0a.375.375 0 1 1-.75 0a.375.375 0 0 1 .75 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReceiptPercentSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.25 2A2.25 2.25 0 0 0 3 4.25v9a.75.75 0 0 0 1.183.613l1.692-1.195l1.692 1.195a.75.75 0 0 0 .866 0l1.692-1.195l1.693 1.195A.75.75 0 0 0 13 13.25v-9A2.25 2.25 0 0 0 10.75 2zm5.53 4.28a.75.75 0 1 0-1.06-1.06l-4.5 4.5a.75.75 0 1 0 1.06 1.06zM7 6.25a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0m2.75 4.25a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReceiptPercentSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.5c-1.921 0-3.816.111-5.68.327c-1.497.174-2.57 1.46-2.57 2.93V21.75a.75.75 0 0 0 1.029.696l3.471-1.388l3.472 1.388a.75.75 0 0 0 .556 0l3.472-1.388l3.471 1.388a.75.75 0 0 0 1.029-.696V4.757c0-1.47-1.073-2.756-2.57-2.93A49.255 49.255 0 0 0 12 1.5m3.53 7.28a.75.75 0 0 0-1.06-1.06l-6 6a.75.75 0 1 0 1.06 1.06zM8.625 9a1.125 1.125 0 1 1 2.25 0a1.125 1.125 0 0 1-2.25 0m5.625 3.375a1.125 1.125 0 1 0 0 2.25a1.125 1.125 0 0 0 0-2.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReceiptPercentTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.93 2.31a41.401 41.401 0 0 1 10.14 0A2.213 2.213 0 0 1 17 4.517V17.25a.75.75 0 0 1-1.075.676l-2.8-1.344l-2.8 1.344a.75.75 0 0 1-.65 0l-2.8-1.344l-2.8 1.344A.75.75 0 0 1 3 17.25V4.517c0-1.103.806-2.068 1.93-2.207m8.85 4.97a.75.75 0 0 0-1.06-1.06l-6.5 6.5a.75.75 0 1 0 1.06 1.06zM9 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0m3 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReceiptRefund(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.25 9.75h4.875a2.625 2.625 0 0 1 0 5.25H12M8.25 9.75L10.5 7.5M8.25 9.75L10.5 12m9-7.243V21.75l-3.75-1.5l-3.75 1.5l-3.75-1.5l-3.75 1.5V4.757c0-1.108.806-2.057 1.907-2.185a48.507 48.507 0 0 1 11.186 0c1.1.128 1.907 1.077 1.907 2.185"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReceiptRefundSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.25 2A2.25 2.25 0 0 0 3 4.25v9a.75.75 0 0 0 1.183.613l1.692-1.195l1.692 1.195a.75.75 0 0 0 .866 0l1.692-1.195l1.693 1.195A.75.75 0 0 0 13 13.25v-9A2.25 2.25 0 0 0 10.75 2zm3.03 3.28a.75.75 0 0 0-1.06-1.06L4.97 6.47a.75.75 0 0 0 0 1.06l2.25 2.25a.75.75 0 0 0 1.06-1.06l-.97-.97h1.315c.76 0 1.375.616 1.375 1.375a.75.75 0 0 0 1.5 0A2.875 2.875 0 0 0 8.625 6.25H7.311z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReceiptRefundSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.5c-1.921 0-3.816.111-5.68.327c-1.497.174-2.57 1.46-2.57 2.93V21.75a.75.75 0 0 0 1.029.696l3.471-1.388l3.472 1.388a.75.75 0 0 0 .556 0l3.472-1.388l3.471 1.388a.75.75 0 0 0 1.029-.696V4.757c0-1.47-1.073-2.756-2.57-2.93A49.255 49.255 0 0 0 12 1.5m-.97 6.53a.75.75 0 1 0-1.06-1.06L7.72 9.22a.75.75 0 0 0 0 1.06l2.25 2.25a.75.75 0 1 0 1.06-1.06l-.97-.97h3.065a1.875 1.875 0 0 1 0 3.75H12a.75.75 0 0 0 0 1.5h1.125a3.375 3.375 0 1 0 0-6.75h-3.064z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReceiptRefundTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.93 2.31a41.401 41.401 0 0 1 10.14 0A2.213 2.213 0 0 1 17 4.517V17.25a.75.75 0 0 1-1.075.676l-2.8-1.344l-2.8 1.344a.75.75 0 0 1-.65 0l-2.8-1.344l-2.8 1.344A.75.75 0 0 1 3 17.25V4.517c0-1.103.806-2.068 1.93-2.207m4.822 3.997a.75.75 0 1 0-1.004-1.114l-2.5 2.25a.75.75 0 0 0 0 1.114l2.5 2.25a.75.75 0 0 0 1.004-1.114L8.704 8.75h1.921a1.875 1.875 0 0 1 0 3.75a.75.75 0 0 0 0 1.5a3.375 3.375 0 1 0 0-6.75h-1.92z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RectangleGroup(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.25 7.125C2.25 6.504 2.754 6 3.375 6h6c.621 0 1.125.504 1.125 1.125v3.75c0 .621-.504 1.125-1.125 1.125h-6a1.125 1.125 0 0 1-1.125-1.125zm12 1.5c0-.621.504-1.125 1.125-1.125h5.25c.621 0 1.125.504 1.125 1.125v8.25c0 .621-.504 1.125-1.125 1.125h-5.25a1.125 1.125 0 0 1-1.125-1.125zm-10.5 7.5c0-.621.504-1.125 1.125-1.125h5.25c.621 0 1.125.504 1.125 1.125v2.25c0 .621-.504 1.125-1.125 1.125h-5.25a1.125 1.125 0 0 1-1.125-1.125z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RectangleGroupSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 4a1 1 0 0 1 1-1h5a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1zm9 1a1 1 0 0 1 1-1h3a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-3a1 1 0 0 1-1-1zm-6 5a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-1a1 1 0 0 0-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RectangleGroupSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 7.125c0-1.036.84-1.875 1.875-1.875h6c1.036 0 1.875.84 1.875 1.875v3.75c0 1.036-.84 1.875-1.875 1.875h-6A1.875 1.875 0 0 1 1.5 10.875zm12 1.5c0-1.036.84-1.875 1.875-1.875h5.25c1.035 0 1.875.84 1.875 1.875v8.25c0 1.035-.84 1.875-1.875 1.875h-5.25a1.875 1.875 0 0 1-1.875-1.875zM3 16.125c0-1.036.84-1.875 1.875-1.875h5.25c1.036 0 1.875.84 1.875 1.875v2.25c0 1.035-.84 1.875-1.875 1.875h-5.25A1.875 1.875 0 0 1 3 18.375z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RectangleGroupTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 3A1.5 1.5 0 0 0 1 4.5v4A1.5 1.5 0 0 0 2.5 10h6A1.5 1.5 0 0 0 10 8.5v-4A1.5 1.5 0 0 0 8.5 3zm11 2A1.5 1.5 0 0 0 12 6.5v7a1.5 1.5 0 0 0 1.5 1.5h4a1.5 1.5 0 0 0 1.5-1.5v-7A1.5 1.5 0 0 0 17.5 5zm-10 7A1.5 1.5 0 0 0 2 13.5v2A1.5 1.5 0 0 0 3.5 17h6a1.5 1.5 0 0 0 1.5-1.5v-2A1.5 1.5 0 0 0 9.5 12z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RectangleStack(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 6.878V6a2.25 2.25 0 0 1 2.25-2.25h7.5A2.25 2.25 0 0 1 18 6v.878m-12 0c.235-.083.487-.128.75-.128h10.5c.263 0 .515.045.75.128m-12 0A2.25 2.25 0 0 0 4.5 9v.878m13.5-3A2.25 2.25 0 0 1 19.5 9v.878m0 0a2.246 2.246 0 0 0-.75-.128H5.25c-.263 0-.515.045-.75.128m15 0A2.25 2.25 0 0 1 21 12v6a2.25 2.25 0 0 1-2.25 2.25H5.25A2.25 2.25 0 0 1 3 18v-6c0-.98.626-1.813 1.5-2.122"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RectangleStackSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3.5A1.5 1.5 0 0 1 6.5 2h3A1.5 1.5 0 0 1 11 3.5zM4.5 5A1.5 1.5 0 0 0 3 6.5v.041a3.02 3.02 0 0 1 .5-.041h9c.17 0 .337.014.5.041V6.5A1.5 1.5 0 0 0 11.5 5zm8 3h-9A1.5 1.5 0 0 0 2 9.5v3A1.5 1.5 0 0 0 3.5 14h9a1.5 1.5 0 0 0 1.5-1.5v-3A1.5 1.5 0 0 0 12.5 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RectangleStackSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.566 4.657A4.505 4.505 0 0 1 6.75 4.5h10.5c.41 0 .806.055 1.183.157A3 3 0 0 0 15.75 3h-7.5a3 3 0 0 0-2.684 1.657M2.25 12a3 3 0 0 1 3-3h13.5a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3H5.25a3 3 0 0 1-3-3zm3-4.5c-.41 0-.806.055-1.184.157A3 3 0 0 1 6.75 6h10.5a3 3 0 0 1 2.683 1.657A4.505 4.505 0 0 0 18.75 7.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RectangleStackTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.127 3.502L5.25 3.5h9.5c.041 0 .082 0 .123.002A2.251 2.251 0 0 0 12.75 2h-5.5a2.25 2.25 0 0 0-2.123 1.502M1 10.25A2.25 2.25 0 0 1 3.25 8h13.5A2.25 2.25 0 0 1 19 10.25v5.5A2.25 2.25 0 0 1 16.75 18H3.25A2.25 2.25 0 0 1 1 15.75zM3.25 6.5c-.04 0-.082 0-.123.002A2.25 2.25 0 0 1 5.25 5h9.5c.98 0 1.814.627 2.123 1.502a3.819 3.819 0 0 0-.123-.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RocketLaunch(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.59 14.37a6 6 0 0 1-5.84 7.38v-4.8m5.84-2.58a14.98 14.98 0 0 0 6.16-12.12A14.98 14.98 0 0 0 9.631 8.41m5.96 5.96a14.926 14.926 0 0 1-5.841 2.58m-.119-8.54a6 6 0 0 0-7.381 5.84h4.8m2.581-5.84a14.927 14.927 0 0 0-2.58 5.84m2.699 2.7a15.53 15.53 0 0 1-.311.06a15.09 15.09 0 0 1-2.448-2.448a14.9 14.9 0 0 1 .06-.312m-2.24 2.39a4.493 4.493 0 0 0-1.757 4.306a4.493 4.493 0 0 0 4.306-1.758M16.5 9a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RocketLaunchSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.333 4.478A4 4 0 0 0 1 8.25c0 .414.336.75.75.75h3.322A12.78 12.78 0 0 0 7 10.928v3.322c0 .414.336.75.75.75a4 4 0 0 0 3.772-5.333A10.721 10.721 0 0 0 15 1.75a.75.75 0 0 0-.75-.75a10.72 10.72 0 0 0-7.917 3.478M12 5.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0" clip-rule="evenodd"/><path fill="currentColor" d="M3.902 10.682a.75.75 0 1 0-1.313-.725a4.764 4.764 0 0 0-.469 3.36a.75.75 0 0 0 .564.563a4.76 4.76 0 0 0 3.359-.47a.75.75 0 1 0-.725-1.312a3.231 3.231 0 0 1-1.81.393a3.232 3.232 0 0 1 .394-1.81"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RocketLaunchSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M9.315 7.584A15.723 15.723 0 0 1 21.75 1.5a.75.75 0 0 1 .75.75c0 5.056-2.383 9.555-6.084 12.436A6.75 6.75 0 0 1 9.75 22.5a.75.75 0 0 1-.75-.75v-4.131A15.838 15.838 0 0 1 6.382 15H2.25a.75.75 0 0 1-.75-.75a6.75 6.75 0 0 1 7.815-6.666M15 6.75a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5" clip-rule="evenodd"/><path d="M5.26 17.242a.75.75 0 1 0-.897-1.203a5.243 5.243 0 0 0-2.05 5.022a.75.75 0 0 0 .625.627a5.243 5.243 0 0 0 5.022-2.051a.75.75 0 1 0-1.202-.897a3.744 3.744 0 0 1-3.008 1.51c0-1.23.592-2.323 1.51-3.008"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RocketLaunchTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M4.606 12.97a.75.75 0 0 1-.134 1.051a2.494 2.494 0 0 0-.93 2.437a2.494 2.494 0 0 0 2.437-.93a.75.75 0 1 1 1.186.918a3.995 3.995 0 0 1-4.482 1.332a.75.75 0 0 1-.461-.461a3.994 3.994 0 0 1 1.332-4.482a.75.75 0 0 1 1.052.134"/><path d="M5.752 12A13.07 13.07 0 0 0 8 14.248v4.002c0 .414.336.75.75.75a5 5 0 0 0 4.797-6.414a12.984 12.984 0 0 0 5.45-10.848a.75.75 0 0 0-.735-.735a12.984 12.984 0 0 0-10.849 5.45A5 5 0 0 0 1 11.25c.001.414.337.75.751.75zM13 9a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rss(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.75 19.5v-.75a7.5 7.5 0 0 0-7.5-7.5H4.5m0-6.75h.75c7.87 0 14.25 6.38 14.25 14.25v.75M6 18.75a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RssSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 2.75A.75.75 0 0 1 2.75 2C8.963 2 14 7.037 14 13.25a.75.75 0 0 1-1.5 0A9.75 9.75 0 0 0 2.75 3.5A.75.75 0 0 1 2 2.75m0 4.5a.75.75 0 0 1 .75-.75a6.75 6.75 0 0 1 6.75 6.75a.75.75 0 0 1-1.5 0C8 10.35 5.65 8 2.75 8A.75.75 0 0 1 2 7.25M3.5 11a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RssSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.75 4.5a.75.75 0 0 1 .75-.75h.75c8.284 0 15 6.716 15 15v.75a.75.75 0 0 1-.75.75h-.75a.75.75 0 0 1-.75-.75v-.75C18 11.708 12.292 6 5.25 6H4.5a.75.75 0 0 1-.75-.75zm0 6.75a.75.75 0 0 1 .75-.75h.75a8.25 8.25 0 0 1 8.25 8.25v.75a.75.75 0 0 1-.75.75H12a.75.75 0 0 1-.75-.75v-.75a6 6 0 0 0-6-6H4.5a.75.75 0 0 1-.75-.75zm0 7.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RssTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M3.75 3a.75.75 0 0 0-.75.75v.5c0 .414.336.75.75.75H4c6.075 0 11 4.925 11 11v.25c0 .414.336.75.75.75h.5a.75.75 0 0 0 .75-.75V16C17 8.82 11.18 3 4 3z"/><path d="M3 8.75A.75.75 0 0 1 3.75 8H4a8 8 0 0 1 8 8v.25a.75.75 0 0 1-.75.75h-.5a.75.75 0 0 1-.75-.75V16a6 6 0 0 0-6-6h-.25A.75.75 0 0 1 3 9.25zM7 15a2 2 0 1 1-4 0a2 2 0 0 1 4 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scale(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 3v17.25m0 0c-1.472 0-2.882.265-4.185.75M12 20.25c1.472 0 2.882.265 4.185.75M18.75 4.97A48.416 48.416 0 0 0 12 4.5c-2.291 0-4.545.16-6.75.47m13.5 0c1.01.143 2.01.317 3 .52m-3-.52l2.62 10.726c.122.499-.106 1.028-.589 1.202a5.988 5.988 0 0 1-2.031.352a5.988 5.988 0 0 1-2.031-.352c-.483-.174-.711-.703-.59-1.202zm-16.5.52c.99-.203 1.99-.377 3-.52m0 0l2.62 10.726c.122.499-.106 1.028-.589 1.202a5.989 5.989 0 0 1-2.031.352a5.989 5.989 0 0 1-2.031-.352c-.483-.174-.711-.703-.59-1.202z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScaleSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.75 2.5a.75.75 0 0 0-1.5 0v.508a32.661 32.661 0 0 0-4.624.434a.75.75 0 0 0 .246 1.48l.13-.021l-1.188 4.75a.75.75 0 0 0 .33.817A3.487 3.487 0 0 0 4 11a3.49 3.49 0 0 0 1.856-.532a.75.75 0 0 0 .33-.818l-1.25-5a31.31 31.31 0 0 1 2.314-.141v7.503c-.882.027-1.752.104-2.607.226a.75.75 0 0 0 .213 1.485a22.188 22.188 0 0 1 6.288 0a.75.75 0 1 0 .213-1.485a23.657 23.657 0 0 0-2.607-.226V4.509c.779.018 1.55.066 2.314.14L9.814 9.65a.75.75 0 0 0 .329.818a3.487 3.487 0 0 0 1.856.532a3.49 3.49 0 0 0 1.856-.532a.75.75 0 0 0 .33-.818L12.997 4.9l.13.022a.75.75 0 1 0 .247-1.48a32.66 32.66 0 0 0-4.624-.434zM3.42 9.415a2 2 0 0 0 1.16 0L4 7.092zM12 9.5a2 2 0 0 1-.582-.085L12 7.092l.58 2.323A2 2 0 0 1 12 9.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScaleSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25a.75.75 0 0 1 .75.75v.756a49.106 49.106 0 0 1 9.152 1a.75.75 0 0 1-.152 1.485h-1.918l2.474 10.124a.75.75 0 0 1-.375.84A6.723 6.723 0 0 1 18.75 18a6.723 6.723 0 0 1-3.181-.795a.75.75 0 0 1-.375-.84l2.474-10.124H12.75v13.28c1.293.076 2.534.343 3.697.776a.75.75 0 0 1-.262 1.453h-8.37a.75.75 0 0 1-.262-1.453c1.162-.433 2.404-.7 3.697-.775V6.24H6.332l2.474 10.124a.75.75 0 0 1-.375.84A6.723 6.723 0 0 1 5.25 18a6.723 6.723 0 0 1-3.181-.795a.75.75 0 0 1-.375-.84L4.168 6.241H2.25a.75.75 0 0 1-.152-1.485a49.105 49.105 0 0 1 9.152-1V3a.75.75 0 0 1 .75-.75m4.878 13.543l1.872-7.662l1.872 7.662zm-9.756 0L5.25 8.131l-1.872 7.662z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScaleTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 2a.75.75 0 0 1 .75.75v.258a33.186 33.186 0 0 1 6.668.83a.75.75 0 0 1-.336 1.461a31.28 31.28 0 0 0-1.103-.232l1.702 7.545a.75.75 0 0 1-.387.832A4.981 4.981 0 0 1 15 14a4.98 4.98 0 0 1-2.294-.556a.75.75 0 0 1-.387-.832l1.77-7.849a31.743 31.743 0 0 0-3.339-.254v11.505a20.01 20.01 0 0 1 3.78.501a.75.75 0 1 1-.339 1.462A18.558 18.558 0 0 0 10 17.5c-1.442 0-2.845.165-4.191.477a.75.75 0 0 1-.338-1.462a20.01 20.01 0 0 1 3.779-.501V4.509a31.692 31.692 0 0 0-3.34.254l1.771 7.85a.75.75 0 0 1-.387.83A4.98 4.98 0 0 1 5 14a4.98 4.98 0 0 1-2.294-.556a.75.75 0 0 1-.387-.832L4.02 5.067c-.37.07-.738.148-1.103.232a.75.75 0 0 1-.336-1.462a32.845 32.845 0 0 1 6.668-.829V2.75A.75.75 0 0 1 10 2M5 7.543L3.92 12.33a3.499 3.499 0 0 0 2.16 0zm10 0l-1.08 4.787a3.498 3.498 0 0 0 2.16 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scissors(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m7.848 8.25l1.536.887M7.848 8.25a3 3 0 1 1-5.196-3a3 3 0 0 1 5.196 3m1.536.887a2.165 2.165 0 0 1 1.083 1.839c.005.351.054.695.14 1.024M9.384 9.137l2.077 1.199M7.848 15.75l1.536-.887m-1.536.887a3 3 0 1 1-5.196 3a3 3 0 0 1 5.196-3m1.536-.887a2.165 2.165 0 0 0 1.083-1.838c.005-.352.054-.695.14-1.025m-1.223 2.863l2.077-1.199m0-3.328a4.323 4.323 0 0 1 2.068-1.379l5.325-1.628a4.5 4.5 0 0 1 2.48-.044l.803.215l-7.794 4.5m-2.882-1.664A4.33 4.33 0 0 0 10.607 12m3.736 0l7.794 4.5l-.802.215a4.5 4.5 0 0 1-2.48-.043l-5.326-1.629a4.324 4.324 0 0 1-2.068-1.379M14.343 12l-2.882 1.664"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScissorsSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 6.665a2.5 2.5 0 0 0 2.94-.323l.359.207c.34.196.777.02.97-.322c.19-.337.115-.784-.22-.977l-.359-.207a2.501 2.501 0 1 0-3.69 1.622M4.364 5a1 1 0 1 1-1.732-1a1 1 0 0 1 1.732 1m4.539.465a2.75 2.75 0 0 0-1.775 1.893l-.375 1.398l-1.563.902a2.501 2.501 0 1 0 .75 1.3L14.7 5.9a.75.75 0 0 0-.18-1.374l-.782-.21a2.75 2.75 0 0 0-1.593.052zM4.365 11a1 1 0 1 1-1.732 1a1 1 0 0 1 1.732-1" clip-rule="evenodd"/><path fill="currentColor" d="M8.892 10.408c-.052.03-.047.108.011.128l3.243 1.097a2.75 2.75 0 0 0 1.593.05l.781-.208a.75.75 0 0 0 .18-1.374l-2.137-1.235a1 1 0 0 0-1 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScissorsSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M8.128 9.155a3.751 3.751 0 1 1 .713-1.321l1.136.656a.75.75 0 0 1 .222 1.104l-.006.007a.75.75 0 0 1-1.032.157a1.421 1.421 0 0 0-.113-.072zm-4.827-3.53a2.25 2.25 0 0 1 3.994 2.063a.756.756 0 0 0-.122.23a2.25 2.25 0 0 1-3.872-2.293m10.047 2.647a5.073 5.073 0 0 0-3.428 3.57a5.08 5.08 0 0 0-.165 1.202a1.415 1.415 0 0 1-.707 1.201l-.96.554a3.751 3.751 0 1 0 .734 1.309l13.729-7.926a.75.75 0 0 0-.181-1.374l-.803-.215a5.25 5.25 0 0 0-2.894.05zm-9.223 7.03a2.25 2.25 0 1 0 2.25 3.897a2.25 2.25 0 0 0-2.25-3.897M12 12.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5" clip-rule="evenodd"/><path d="M16.372 12.615a.75.75 0 0 1 .75 0l5.43 3.135a.75.75 0 0 1-.182 1.374l-.802.215a5.25 5.25 0 0 1-2.894-.051l-5.147-1.574a.75.75 0 0 1-.156-1.367z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScissorsTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M1.469 3.75a3.5 3.5 0 0 0 5.617 4.11l.883.51c.025.092.147.116.21.043c.15-.176.318-.338.5-.484c.286-.23.3-.709-.018-.892l-.825-.477A3.501 3.501 0 0 0 1.47 3.75M3.5 7.232a2 2 0 1 1 2-3.464a2 2 0 0 1-2 3.464m6.456 1.09a2.75 2.75 0 0 0-1.588 1.822L7.97 11.63l-.884.51a3.501 3.501 0 0 0-5.616 4.11a3.5 3.5 0 0 0 6.367-2.81l10.68-6.166a.75.75 0 0 0-.182-1.373l-.703-.189a2.75 2.75 0 0 0-1.78.123zM2.768 15.5a2 2 0 1 1 3.464-2a2 2 0 0 1-3.464 2" clip-rule="evenodd"/><path d="M12.52 11.89a.5.5 0 0 0 .056.894l3.274 1.381a2.75 2.75 0 0 0 1.78.123l.704-.189a.75.75 0 0 0 .18-1.373l-3.47-2.004a.5.5 0 0 0-.5 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Server(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21.75 17.25v-.228a4.5 4.5 0 0 0-.12-1.03l-2.268-9.64a3.375 3.375 0 0 0-3.285-2.602H7.923a3.375 3.375 0 0 0-3.285 2.602l-2.268 9.64a4.5 4.5 0 0 0-.12 1.03v.228m19.5 0a3 3 0 0 1-3 3H5.25a3 3 0 0 1-3-3m19.5 0a3 3 0 0 0-3-3H5.25a3 3 0 0 0-3 3m16.5 0h.008v.008h-.008zm-3 0h.008v.008h-.008z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServerSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.665 3.588A2 2 0 0 1 5.622 2h4.754a2 2 0 0 1 1.958 1.588l1.098 5.218a3.487 3.487 0 0 0-1.433-.306H4a3.5 3.5 0 0 0-1.433.306z"/><path fill="currentColor" fill-rule="evenodd" d="M4 10a2 2 0 1 0 0 4h8a2 2 0 1 0 0-4zm8 2.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5M9.75 12a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServerSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M4.08 5.227A3 3 0 0 1 6.979 3H17.02a3 3 0 0 1 2.9 2.227l2.113 7.926A5.228 5.228 0 0 0 18.75 12H5.25a5.228 5.228 0 0 0-3.284 1.153z"/><path fill-rule="evenodd" d="M5.25 13.5a3.75 3.75 0 1 0 0 7.5h13.5a3.75 3.75 0 1 0 0-7.5zm10.5 4.5a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5m3.75-.75a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServerStack(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 14.25h13.5m-13.5 0a3 3 0 0 1-3-3m3 3a3 3 0 1 0 0 6h13.5a3 3 0 1 0 0-6m-16.5-3a3 3 0 0 1 3-3h13.5a3 3 0 0 1 3 3m-19.5 0a4.5 4.5 0 0 1 .9-2.7L5.738 5.1a3.375 3.375 0 0 1 2.7-1.35h7.124c1.063 0 2.063.5 2.7 1.35l2.588 3.45a4.5 4.5 0 0 1 .9 2.7m0 0a3 3 0 0 1-3 3m0 3h.008v.008h-.008zm0-6h.008v.008h-.008zm-3 6h.008v.008h-.008zm0-6h.008v.008h-.008z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServerStackSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.354 2a2 2 0 0 0-1.857 1.257l-.338.845A3.49 3.49 0 0 1 4 4h8c.29 0 .571.035.84.102l-.337-.845A2 2 0 0 0 10.646 2z"/><path fill="currentColor" fill-rule="evenodd" d="M2 13a2 2 0 0 1 2-2h8a2 2 0 1 1 0 4H4a2 2 0 0 1-2-2m10.75 0a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0M9 13.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5M4 5.5a2 2 0 1 0 0 4h8a2 2 0 1 0 0-4zm8 2.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5M9.75 7.5a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServerStackSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M5.507 4.048A3 3 0 0 1 7.785 3h8.43a3 3 0 0 1 2.278 1.048l1.722 2.008A4.533 4.533 0 0 0 19.5 6h-15c-.243 0-.482.02-.715.056z"/><path fill-rule="evenodd" d="M1.5 10.5a3 3 0 0 1 3-3h15a3 3 0 1 1 0 6h-15a3 3 0 0 1-3-3m15 0a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0m2.25.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5M4.5 15a3 3 0 1 0 0 6h15a3 3 0 1 0 0-6zm11.25 3.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5M19.5 18a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServerStackTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M4.464 3.162A2 2 0 0 1 6.28 2h7.44a2 2 0 0 1 1.816 1.162l1.154 2.5c.067.145.115.291.145.438A3.508 3.508 0 0 0 16 6H4c-.288 0-.568.035-.835.1a1.99 1.99 0 0 1 .145-.438z"/><path fill-rule="evenodd" d="M2 9.5a2 2 0 0 1 2-2h12a2 2 0 1 1 0 4H4a2 2 0 0 1-2-2m13.24 0a.75.75 0 0 1 .75-.75H16a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75h-.01a.75.75 0 0 1-.75-.75zm-2.25-.75a.75.75 0 0 0-.75.75v.01c0 .414.336.75.75.75H13a.75.75 0 0 0 .75-.75V9.5a.75.75 0 0 0-.75-.75zM2 15a2 2 0 0 1 2-2h12a2 2 0 1 1 0 4H4a2 2 0 0 1-2-2m13.24 0a.75.75 0 0 1 .75-.75H16a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75h-.01a.75.75 0 0 1-.75-.75zm-2.25-.75a.75.75 0 0 0-.75.75v.01c0 .414.336.75.75.75H13a.75.75 0 0 0 .75-.75V15a.75.75 0 0 0-.75-.75z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServerTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M4.632 3.533A2 2 0 0 1 6.577 2h6.846a2 2 0 0 1 1.945 1.533l1.976 8.234A3.489 3.489 0 0 0 16 11.5H4c-.476 0-.93.095-1.344.267z"/><path fill-rule="evenodd" d="M4 13a2 2 0 1 0 0 4h12a2 2 0 1 0 0-4zm11.24 2a.75.75 0 0 1 .75-.75H16a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75h-.01a.75.75 0 0 1-.75-.75zm-2.25-.75a.75.75 0 0 0-.75.75v.01c0 .414.336.75.75.75H13a.75.75 0 0 0 .75-.75V15a.75.75 0 0 0-.75-.75z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.217 10.907a2.25 2.25 0 1 0 0 2.186m0-2.186c.18.324.283.696.283 1.093s-.103.77-.283 1.093m0-2.186l9.566-5.314m-9.566 7.5l9.566 5.314m0 0a2.25 2.25 0 1 0 3.935 2.186a2.25 2.25 0 0 0-3.935-2.186m0-12.814a2.25 2.25 0 1 0 3.933-2.185a2.25 2.25 0 0 0-3.933 2.185"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 6a2 2 0 1 0-1.994-1.842L5.323 6.5a2 2 0 1 0 0 3l4.683 2.342a2 2 0 1 0 .67-1.342L5.995 8.158a2.03 2.03 0 0 0 0-.316L10.677 5.5c.353.311.816.5 1.323.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.75 4.5a3 3 0 1 1 .825 2.066l-8.421 4.679a3.002 3.002 0 0 1 0 1.51l8.421 4.679a3 3 0 1 1-.729 1.31l-8.421-4.678a3 3 0 1 1 0-4.132l8.421-4.679a3 3 0 0 1-.096-.755" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 4.5a2.5 2.5 0 1 1 .702 1.737L6.97 9.604a2.518 2.518 0 0 1 0 .792l6.733 3.367a2.5 2.5 0 1 1-.671 1.341l-6.733-3.367a2.5 2.5 0 1 1 0-3.475l6.733-3.366A2.52 2.52 0 0 1 13 4.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldCheck(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12.75L11.25 15L15 9.75m-3-7.036A11.959 11.959 0 0 1 3.598 6A11.99 11.99 0 0 0 3 9.749c0 5.592 3.824 10.29 9 11.623c5.176-1.332 9-6.03 9-11.622c0-1.31-.21-2.571-.598-3.751h-.152c-3.196 0-6.1-1.248-8.25-3.285"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldCheckSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.5 1.709a.75.75 0 0 0-1 0a8.963 8.963 0 0 1-4.84 2.217a.75.75 0 0 0-.654.72a10.499 10.499 0 0 0 5.647 9.672a.75.75 0 0 0 .694-.001a10.499 10.499 0 0 0 5.647-9.672a.75.75 0 0 0-.654-.719A8.963 8.963 0 0 1 8.5 1.71m2.34 5.504a.75.75 0 0 0-1.18-.926L7.394 9.17l-1.156-.99a.75.75 0 1 0-.976 1.138l1.75 1.5a.75.75 0 0 0 1.078-.106z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldCheckSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.516 2.17a.75.75 0 0 0-1.032 0a11.209 11.209 0 0 1-7.877 3.08a.75.75 0 0 0-.722.515A12.74 12.74 0 0 0 2.25 9.75c0 5.942 4.064 10.933 9.563 12.348a.749.749 0 0 0 .374 0c5.499-1.415 9.563-6.406 9.563-12.348c0-1.39-.223-2.73-.635-3.985a.75.75 0 0 0-.722-.516l-.143.001c-2.996 0-5.717-1.17-7.734-3.08m3.094 8.016a.75.75 0 1 0-1.22-.872l-3.236 4.53L9.53 12.22a.75.75 0 0 0-1.06 1.06l2.25 2.25a.75.75 0 0 0 1.14-.094z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldCheckTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.661 2.237a.531.531 0 0 1 .678 0a11.947 11.947 0 0 0 7.078 2.749a.5.5 0 0 1 .479.425c.069.52.104 1.05.104 1.59c0 5.162-3.26 9.563-7.834 11.256a.48.48 0 0 1-.332 0C5.26 16.564 2 12.163 2 7c0-.538.035-1.069.104-1.589a.5.5 0 0 1 .48-.425a11.947 11.947 0 0 0 7.077-2.75m4.196 5.954a.75.75 0 0 0-1.214-.882l-3.483 4.79l-1.88-1.88a.75.75 0 1 0-1.06 1.061l2.5 2.5a.75.75 0 0 0 1.137-.089z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldExclamation(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 9v3.75m0-10.036A11.959 11.959 0 0 1 3.598 6A11.99 11.99 0 0 0 3 9.75c0 5.592 3.824 10.29 9 11.622c5.176-1.332 9-6.03 9-11.622c0-1.31-.21-2.57-.598-3.75h-.152c-3.196 0-6.1-1.25-8.25-3.286m0 13.036h.008v.008H12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldExclamationSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 1.709a.75.75 0 0 1 1 0a8.963 8.963 0 0 0 4.84 2.217a.75.75 0 0 1 .654.72a10.499 10.499 0 0 1-5.647 9.672a.75.75 0 0 1-.694-.001a10.499 10.499 0 0 1-5.647-9.672a.75.75 0 0 1 .654-.719A8.963 8.963 0 0 0 7.5 1.71M8 5a.75.75 0 0 1 .75.75v2a.75.75 0 0 1-1.5 0v-2A.75.75 0 0 1 8 5m0 7a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldExclamationSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.484 2.17a.75.75 0 0 1 1.032 0a11.209 11.209 0 0 0 7.877 3.08a.75.75 0 0 1 .722.515a12.74 12.74 0 0 1 .635 3.985c0 5.942-4.064 10.933-9.563 12.348a.749.749 0 0 1-.374 0C6.314 20.683 2.25 15.692 2.25 9.75c0-1.39.223-2.73.635-3.985a.75.75 0 0 1 .722-.516l.143.001c2.996 0 5.718-1.17 7.734-3.08M12 8.25a.75.75 0 0 1 .75.75v3.75a.75.75 0 0 1-1.5 0V9a.75.75 0 0 1 .75-.75M12 15a.75.75 0 0 0-.75.75v.008c0 .414.336.75.75.75h.008a.75.75 0 0 0 .75-.75v-.008a.75.75 0 0 0-.75-.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldExclamationTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.339 2.237a.531.531 0 0 0-.678 0a11.947 11.947 0 0 1-7.078 2.75a.5.5 0 0 0-.479.425A12.11 12.11 0 0 0 2 7c0 5.163 3.26 9.564 7.834 11.257a.48.48 0 0 0 .332 0C14.74 16.564 18 12.163 18 7c0-.538-.035-1.069-.104-1.589a.5.5 0 0 0-.48-.425a11.947 11.947 0 0 1-7.077-2.75M10 6a.75.75 0 0 1 .75.75v3.5a.75.75 0 0 1-1.5 0v-3.5A.75.75 0 0 1 10 6m0 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBag(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.75 10.5V6a3.75 3.75 0 1 0-7.5 0v4.5m11.356-1.993l1.263 12c.07.665-.45 1.243-1.119 1.243H4.25a1.125 1.125 0 0 1-1.12-1.243l1.264-12A1.125 1.125 0 0 1 5.513 7.5h12.974c.576 0 1.059.435 1.119 1.007M8.625 10.5a.375.375 0 1 1-.75 0a.375.375 0 0 1 .75 0m7.5 0a.375.375 0 1 1-.75 0a.375.375 0 0 1 .75 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBagSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 4a3 3 0 0 1 6 0v1h.643a1.5 1.5 0 0 1 1.492 1.35l.7 7A1.5 1.5 0 0 1 12.342 15H3.657a1.5 1.5 0 0 1-1.492-1.65l.7-7A1.5 1.5 0 0 1 4.357 5H5zm4.5 0v1h-3V4a1.5 1.5 0 0 1 3 0m-3 3.75a.75.75 0 0 0-1.5 0v1a3 3 0 1 0 6 0v-1a.75.75 0 0 0-1.5 0v1a1.5 1.5 0 1 1-3 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBagSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 6v.75H5.513c-.96 0-1.764.724-1.865 1.679l-1.263 12A1.875 1.875 0 0 0 4.25 22.5h15.5a1.875 1.875 0 0 0 1.865-2.071l-1.263-12a1.875 1.875 0 0 0-1.865-1.679H16.5V6a4.5 4.5 0 1 0-9 0M12 3a3 3 0 0 0-3 3v.75h6V6a3 3 0 0 0-3-3m-3 8.25a3 3 0 1 0 6 0v-.75a.75.75 0 0 1 1.5 0v.75a4.5 4.5 0 1 1-9 0v-.75a.75.75 0 0 1 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBagTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 5v1H4.667a1.75 1.75 0 0 0-1.743 1.598l-.826 9.5A1.75 1.75 0 0 0 3.84 19h12.32a1.75 1.75 0 0 0 1.743-1.902l-.826-9.5A1.75 1.75 0 0 0 15.333 6H14V5a4 4 0 0 0-8 0m4-2.5A2.5 2.5 0 0 0 7.5 5v1h5V5A2.5 2.5 0 0 0 10 2.5M7.5 10a2.5 2.5 0 0 0 5 0V8.75a.75.75 0 0 1 1.5 0V10a4 4 0 0 1-8 0V8.75a.75.75 0 0 1 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCart(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.25 3h1.386c.51 0 .955.343 1.087.835l.383 1.437M7.5 14.25a3 3 0 0 0-3 3h15.75m-12.75-3h11.218c1.121-2.3 2.1-4.684 2.924-7.137a60.114 60.114 0 0 0-16.536-1.84M7.5 14.25L5.106 5.272M6 20.25a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0m12.75 0a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCartSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.75 1.002a.75.75 0 1 0 0 1.5h1.835l1.24 5.113A3.752 3.752 0 0 0 2 11.25c0 .414.336.75.75.75h10.5a.75.75 0 0 0 0-1.5H3.628A2.25 2.25 0 0 1 5.75 9h6.5a.75.75 0 0 0 .73-.578l.846-3.595a.75.75 0 0 0-.578-.906a44.118 44.118 0 0 0-7.996-.91l-.348-1.436a.75.75 0 0 0-.73-.573zM5 14a1 1 0 1 1-2 0a1 1 0 0 1 2 0m8 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCartSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.25 2.25a.75.75 0 0 0 0 1.5h1.386c.17 0 .318.114.362.278l2.558 9.592a3.752 3.752 0 0 0-2.806 3.63c0 .414.336.75.75.75h15.75a.75.75 0 0 0 0-1.5H5.378A2.25 2.25 0 0 1 7.5 15h11.218a.75.75 0 0 0 .674-.421a60.358 60.358 0 0 0 2.96-7.228a.75.75 0 0 0-.525-.965A60.864 60.864 0 0 0 5.68 4.509l-.232-.867A1.875 1.875 0 0 0 3.636 2.25zm1.5 18a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0m12.75 0a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCartTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 1.75A.75.75 0 0 1 1.75 1h1.628a1.75 1.75 0 0 1 1.734 1.51L5.18 3a65.25 65.25 0 0 1 13.36 1.412a.75.75 0 0 1 .58.875a48.645 48.645 0 0 1-1.618 6.2a.75.75 0 0 1-.712.513H6a2.503 2.503 0 0 0-2.292 1.5H17.25a.75.75 0 0 1 0 1.5H2.76a.75.75 0 0 1-.748-.807a4.002 4.002 0 0 1 2.716-3.486L3.626 2.716a.25.25 0 0 0-.248-.216H1.75A.75.75 0 0 1 1 1.75M6 17.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0m9.5 1.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Signal(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.348 14.652a3.75 3.75 0 0 1 0-5.304m5.304 0a3.75 3.75 0 0 1 0 5.304m-7.425 2.121a6.75 6.75 0 0 1 0-9.546m9.546 0a6.75 6.75 0 0 1 0 9.546M5.106 18.894c-3.808-3.807-3.808-9.98 0-13.788m13.788 0c3.808 3.807 3.808 9.98 0 13.788M12 12h.008v.008H12zm.375 0a.375.375 0 1 1-.75 0a.375.375 0 0 1 .75 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignalSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/><path fill="currentColor" fill-rule="evenodd" d="M9.68 5.26a.75.75 0 0 1 1.06 0a3.875 3.875 0 0 1 0 5.48a.75.75 0 1 1-1.06-1.06a2.375 2.375 0 0 0 0-3.36a.75.75 0 0 1 0-1.06m-3.36 0a.75.75 0 0 1 0 1.06a2.375 2.375 0 0 0 0 3.36a.75.75 0 1 1-1.06 1.06a3.875 3.875 0 0 1 0-5.48a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M11.89 3.05a.75.75 0 0 1 1.06 0a7 7 0 0 1 0 9.9a.75.75 0 1 1-1.06-1.06a5.5 5.5 0 0 0 0-7.78a.75.75 0 0 1 0-1.06m-7.78 0a.75.75 0 0 1 0 1.06a5.5 5.5 0 0 0 0 7.78a.75.75 0 1 1-1.06 1.06a7 7 0 0 1 0-9.9a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignalSlash(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3 3l8.735 8.735m0 0a.374.374 0 1 1 .53.53m-.53-.53l.53.53m0 0L21 21M14.652 9.348a3.75 3.75 0 0 1 0 5.304m2.121-7.425a6.75 6.75 0 0 1 0 9.546m2.121-11.667c3.808 3.807 3.808 9.98 0 13.788m-9.546-4.242a3.733 3.733 0 0 1-1.06-2.122m-1.061 4.243a6.75 6.75 0 0 1-1.625-6.929m-.496 9.05c-3.068-3.067-3.664-7.67-1.79-11.334M12 12h.008v.008H12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignalSlashSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.28 2.22a.75.75 0 0 0-1.06 1.06l4.782 4.783a1 1 0 0 0 .935.935l4.783 4.782a.75.75 0 1 0 1.06-1.06L8.998 7.937a1 1 0 0 0-.935-.935zm-.23 10.73a7.003 7.003 0 0 1-1.33-8.047L2.86 6.04a5.501 5.501 0 0 0 1.25 5.849a.75.75 0 1 1-1.06 1.06m2.21-2.209a3.87 3.87 0 0 1-1.082-3.38L5.87 9.052c.112.226.262.439.45.627a.75.75 0 1 1-1.06 1.061m7.69-7.69a7.003 7.003 0 0 1 1.33 8.048l-1.14-1.139a5.501 5.501 0 0 0-1.25-5.848a.75.75 0 0 1 1.06-1.06M10.74 5.26a3.87 3.87 0 0 1 1.082 3.38L10.13 6.948a2.372 2.372 0 0 0-.45-.627a.75.75 0 0 1 1.06-1.061"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignalSlashSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.47 2.47a.75.75 0 0 1 1.06 0l8.407 8.407a1.125 1.125 0 0 1 1.186 1.186l1.462 1.461a3.001 3.001 0 0 0-.464-3.645a.75.75 0 1 1 1.061-1.061a4.501 4.501 0 0 1 .486 5.79l1.072 1.072a6.001 6.001 0 0 0-.497-7.923a.75.75 0 0 1 1.06-1.06a7.501 7.501 0 0 1 .505 10.05l1.064 1.065a9 9 0 0 0-.508-12.176a.75.75 0 0 1 1.06-1.06c3.923 3.922 4.093 10.175.512 14.3l1.594 1.594a.75.75 0 1 1-1.06 1.06l-2.106-2.105l-2.121-2.122h-.001l-4.705-4.706a.747.747 0 0 1-.127-.126L2.47 3.53a.75.75 0 0 1 0-1.061m1.189 4.422a.75.75 0 0 1 .326 1.01a9.004 9.004 0 0 0 1.651 10.462a.75.75 0 1 1-1.06 1.06C1.27 16.12.63 11.165 2.648 7.219a.75.75 0 0 1 1.01-.326M5.84 9.134a.75.75 0 0 1 .472.95a6 6 0 0 0 1.444 6.159a.75.75 0 0 1-1.06 1.06A7.5 7.5 0 0 1 4.89 9.606a.75.75 0 0 1 .95-.472m2.341 2.653a.75.75 0 0 1 .848.638c.088.62.37 1.218.849 1.696a.75.75 0 0 1-1.061 1.061a4.483 4.483 0 0 1-1.273-2.546a.75.75 0 0 1 .637-.848" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignalSlashTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.22 2.22a.75.75 0 0 1 1.06 0l6.783 6.782a1 1 0 0 1 .935.935l6.782 6.783a.75.75 0 1 1-1.06 1.06l-6.783-6.782a1 1 0 0 1-.935-.935L2.22 3.28a.75.75 0 0 1 0-1.06m1.416 14.144a9.004 9.004 0 0 1-1.39-10.936L3.349 6.53a7.503 7.503 0 0 0 1.348 8.773a.75.75 0 0 1-1.061 1.061m2.828-2.828a5 5 0 0 1-1.213-5.103l1.262 1.262a3.493 3.493 0 0 0 1.012 2.78a.75.75 0 0 1-1.06 1.06m9.899-9.899a9.004 9.004 0 0 1 1.39 10.937l-1.103-1.104a7.503 7.503 0 0 0-1.348-8.772a.75.75 0 1 1 1.061-1.061m-2.828 2.828a5 5 0 0 1 1.213 5.103l-1.262-1.262a3.493 3.493 0 0 0-1.012-2.78a.75.75 0 0 1 1.06-1.06"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignalSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.636 4.575a.75.75 0 0 1 0 1.061a9 9 0 0 0 0 12.728a.75.75 0 1 1-1.06 1.06c-4.101-4.1-4.101-10.748 0-14.849a.75.75 0 0 1 1.06 0m12.728 0a.75.75 0 0 1 1.06 0c4.101 4.1 4.101 10.75 0 14.85a.75.75 0 1 1-1.06-1.061a9 9 0 0 0 0-12.728a.75.75 0 0 1 0-1.06M7.757 6.697a.75.75 0 0 1 0 1.06a6 6 0 0 0 0 8.486a.75.75 0 0 1-1.06 1.06a7.5 7.5 0 0 1 0-10.606a.75.75 0 0 1 1.06 0m8.486 0a.75.75 0 0 1 1.06 0a7.5 7.5 0 0 1 0 10.606a.75.75 0 0 1-1.06-1.06a6 6 0 0 0 0-8.486a.75.75 0 0 1 0-1.06M9.879 8.818a.75.75 0 0 1 0 1.06a3 3 0 0 0 0 4.243a.75.75 0 1 1-1.061 1.061a4.5 4.5 0 0 1 0-6.364a.75.75 0 0 1 1.06 0m4.242 0a.75.75 0 0 1 1.061 0a4.5 4.5 0 0 1 0 6.364a.75.75 0 0 1-1.06-1.06a3 3 0 0 0 0-4.243a.75.75 0 0 1 0-1.061M10.875 12a1.125 1.125 0 1 1 2.25 0a1.125 1.125 0 0 1-2.25 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignalTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M16.364 3.636a.75.75 0 0 0-1.06 1.06a7.5 7.5 0 0 1 0 10.607a.75.75 0 0 0 1.06 1.061a9 9 0 0 0 0-12.728M4.697 4.697a.75.75 0 0 0-1.061-1.061a9 9 0 0 0 0 12.728a.75.75 0 1 0 1.06-1.06a7.5 7.5 0 0 1 0-10.607"/><path d="M12.475 6.464a.75.75 0 0 1 1.06 0a5 5 0 0 1 0 7.072a.75.75 0 0 1-1.06-1.061a3.5 3.5 0 0 0 0-4.95a.75.75 0 0 1 0-1.06m-4.95-.001a.75.75 0 0 1 0 1.061a3.5 3.5 0 0 0 0 4.95a.75.75 0 0 1-1.06 1.06a5 5 0 0 1 0-7.07a.75.75 0 0 1 1.06 0M11 10a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sparkles(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.813 15.904L9 18.75l-.813-2.846a4.5 4.5 0 0 0-3.09-3.09L2.25 12l2.846-.813a4.5 4.5 0 0 0 3.09-3.09L9 5.25l.813 2.846a4.5 4.5 0 0 0 3.09 3.09L15.75 12l-2.846.813a4.5 4.5 0 0 0-3.09 3.09m8.445-7.188L18 9.75l-.259-1.035a3.375 3.375 0 0 0-2.455-2.456L14.25 6l1.036-.259a3.375 3.375 0 0 0 2.455-2.456L18 2.25l.259 1.035a3.375 3.375 0 0 0 2.456 2.456L21.75 6l-1.035.259a3.375 3.375 0 0 0-2.456 2.456m-1.365 11.852L16.5 21.75l-.394-1.183a2.25 2.25 0 0 0-1.423-1.423L13.5 18.75l1.183-.394a2.25 2.25 0 0 0 1.423-1.423l.394-1.183l.394 1.183a2.25 2.25 0 0 0 1.423 1.423l1.183.394l-1.183.394a2.25 2.25 0 0 0-1.423 1.423"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SparklesSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 4a.75.75 0 0 1 .738.616l.252 1.388A1.25 1.25 0 0 0 6.996 7.01l1.388.252a.75.75 0 0 1 0 1.476l-1.388.252A1.25 1.25 0 0 0 5.99 9.996l-.252 1.388a.75.75 0 0 1-1.476 0L4.01 9.996A1.25 1.25 0 0 0 3.004 8.99l-1.388-.252a.75.75 0 0 1 0-1.476l1.388-.252A1.25 1.25 0 0 0 4.01 6.004l.252-1.388A.75.75 0 0 1 5 4m7-3a.75.75 0 0 1 .721.544l.195.682c.118.415.443.74.858.858l.682.195a.75.75 0 0 1 0 1.442l-.682.195a1.25 1.25 0 0 0-.858.858l-.195.682a.75.75 0 0 1-1.442 0l-.195-.682a1.25 1.25 0 0 0-.858-.858l-.682-.195a.75.75 0 0 1 0-1.442l.682-.195a1.25 1.25 0 0 0 .858-.858l.195-.682A.75.75 0 0 1 12 1m-2 10a.75.75 0 0 1 .728.568a.968.968 0 0 0 .704.704a.75.75 0 0 1 0 1.456a.968.968 0 0 0-.704.704a.75.75 0 0 1-1.456 0a.968.968 0 0 0-.704-.704a.75.75 0 0 1 0-1.456a.968.968 0 0 0 .704-.704A.75.75 0 0 1 10 11" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SparklesSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9 4.5a.75.75 0 0 1 .721.544l.813 2.846a3.75 3.75 0 0 0 2.576 2.576l2.846.813a.75.75 0 0 1 0 1.442l-2.846.813a3.75 3.75 0 0 0-2.576 2.576l-.813 2.846a.75.75 0 0 1-1.442 0l-.813-2.846a3.75 3.75 0 0 0-2.576-2.576l-2.846-.813a.75.75 0 0 1 0-1.442l2.846-.813A3.75 3.75 0 0 0 7.466 7.89l.813-2.846A.75.75 0 0 1 9 4.5m9-3a.75.75 0 0 1 .728.568l.258 1.036a2.63 2.63 0 0 0 1.91 1.91l1.036.258a.75.75 0 0 1 0 1.456l-1.036.258a2.63 2.63 0 0 0-1.91 1.91l-.258 1.036a.75.75 0 0 1-1.456 0l-.258-1.036a2.625 2.625 0 0 0-1.91-1.91l-1.036-.258a.75.75 0 0 1 0-1.456l1.036-.258a2.625 2.625 0 0 0 1.91-1.91l.258-1.036A.75.75 0 0 1 18 1.5M16.5 15a.75.75 0 0 1 .712.513l.394 1.183c.15.447.5.799.948.948l1.183.395a.75.75 0 0 1 0 1.422l-1.183.395a1.5 1.5 0 0 0-.948.948l-.395 1.183a.75.75 0 0 1-1.422 0l-.395-1.183a1.5 1.5 0 0 0-.948-.948l-1.183-.395a.75.75 0 0 1 0-1.422l1.183-.395a1.5 1.5 0 0 0 .948-.948l.395-1.183A.75.75 0 0 1 16.5 15" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SparklesTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.98 1.804a1 1 0 0 0-1.96 0l-.24 1.192a1 1 0 0 1-.784.785l-1.192.238a1 1 0 0 0 0 1.962l1.192.238a1 1 0 0 1 .785.785l.238 1.192a1 1 0 0 0 1.962 0l.238-1.192a1 1 0 0 1 .785-.785l1.192-.238a1 1 0 0 0 0-1.962l-1.192-.238a1 1 0 0 1-.785-.785zm-9.031 3.88a1 1 0 0 0-1.898 0l-.683 2.051a1 1 0 0 1-.633.633l-2.051.683a1 1 0 0 0 0 1.898l2.051.684a1 1 0 0 1 .633.632l.683 2.051a1 1 0 0 0 1.898 0l.683-2.051a1 1 0 0 1 .633-.633l2.051-.683a1 1 0 0 0 0-1.898l-2.051-.683a1 1 0 0 1-.633-.633zm7 8a1 1 0 0 0-1.898 0l-.184.551a1 1 0 0 1-.632.633l-.551.183a1 1 0 0 0 0 1.898l.551.183a1 1 0 0 1 .633.633l.183.551a1 1 0 0 0 1.898 0l.184-.551a1 1 0 0 1 .632-.633l.551-.183a1 1 0 0 0 0-1.898l-.551-.184a1 1 0 0 1-.633-.632z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpeakerWave(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.114 5.636a9 9 0 0 1 0 12.728M16.463 8.288a5.25 5.25 0 0 1 0 7.424M6.75 8.25l4.72-4.72a.75.75 0 0 1 1.28.53v15.88a.75.75 0 0 1-1.28.53l-4.72-4.72H4.51c-.88 0-1.704-.507-1.938-1.354A9.009 9.009 0 0 1 2.25 12c0-.83.112-1.633.322-2.396C2.806 8.756 3.63 8.25 4.51 8.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpeakerWaveSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.557 2.066A.75.75 0 0 1 8 2.75v10.5a.75.75 0 0 1-1.248.56L3.59 11H2a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h1.59l3.162-2.81a.75.75 0 0 1 .805-.124m5.393.984a.75.75 0 1 0-1.06 1.06a5.5 5.5 0 0 1 0 7.78a.75.75 0 1 0 1.06 1.06a7 7 0 0 0 0-9.9"/><path fill="currentColor" d="M10.828 5.172a.75.75 0 1 0-1.06 1.06a2.5 2.5 0 0 1 0 3.536a.75.75 0 1 0 1.06 1.06a4 4 0 0 0 0-5.656"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpeakerWaveSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M13.5 4.06c0-1.336-1.616-2.005-2.56-1.06l-4.5 4.5H4.508c-1.141 0-2.318.664-2.66 1.905A9.76 9.76 0 0 0 1.5 12c0 .898.121 1.768.35 2.595c.341 1.24 1.518 1.905 2.659 1.905h1.93l4.5 4.5c.945.945 2.561.276 2.561-1.06zm5.084 1.046a.75.75 0 0 1 1.06 0c3.808 3.807 3.808 9.98 0 13.788a.75.75 0 0 1-1.06-1.06a8.25 8.25 0 0 0 0-11.668a.75.75 0 0 1 0-1.06"/><path d="M15.932 7.757a.75.75 0 0 1 1.061 0a6 6 0 0 1 0 8.486a.75.75 0 0 1-1.06-1.061a4.5 4.5 0 0 0 0-6.364a.75.75 0 0 1 0-1.06"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpeakerWaveTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10.5 3.75a.75.75 0 0 0-1.264-.546L5.203 7H2.667a.75.75 0 0 0-.7.48A6.985 6.985 0 0 0 1.5 10c0 .887.165 1.737.468 2.52c.111.29.39.48.7.48h2.535l4.033 3.796a.75.75 0 0 0 1.264-.546zm5.95 1.3a.75.75 0 0 0-1.06 1.061a5.5 5.5 0 0 1 0 7.778a.75.75 0 0 0 1.06 1.06a7 7 0 0 0 0-9.899"/><path d="M14.329 7.172a.75.75 0 0 0-1.061 1.06a2.5 2.5 0 0 1 0 3.536a.75.75 0 0 0 1.06 1.06a4 4 0 0 0 0-5.656"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpeakerXmark(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.25 9.75L19.5 12m0 0l2.25 2.25M19.5 12l2.25-2.25M19.5 12l-2.25 2.25m-10.5-6l4.72-4.72a.75.75 0 0 1 1.28.53v15.88a.75.75 0 0 1-1.28.53l-4.72-4.72H4.51c-.88 0-1.704-.507-1.938-1.354A9.009 9.009 0 0 1 2.25 12c0-.83.112-1.633.322-2.396C2.806 8.756 3.63 8.25 4.51 8.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpeakerXmarkSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.557 2.066A.75.75 0 0 1 8 2.75v10.5a.75.75 0 0 1-1.248.56L3.59 11H2a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h1.59l3.162-2.81a.75.75 0 0 1 .805-.124M11.28 5.72a.75.75 0 1 0-1.06 1.06L11.44 8l-1.22 1.22a.75.75 0 1 0 1.06 1.06l1.22-1.22l1.22 1.22a.75.75 0 1 0 1.06-1.06L13.56 8l1.22-1.22a.75.75 0 0 0-1.06-1.06L12.5 6.94z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpeakerXmarkSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 4.06c0-1.336-1.616-2.005-2.56-1.06l-4.5 4.5H4.508c-1.141 0-2.318.664-2.66 1.905A9.76 9.76 0 0 0 1.5 12c0 .898.121 1.768.35 2.595c.341 1.24 1.518 1.905 2.659 1.905h1.93l4.5 4.5c.945.945 2.561.276 2.561-1.06zm4.28 5.16a.75.75 0 1 0-1.06 1.06L18.44 12l-1.72 1.72a.75.75 0 1 0 1.06 1.06l1.72-1.72l1.72 1.72a.75.75 0 1 0 1.06-1.06L20.56 12l1.72-1.72a.75.75 0 1 0-1.06-1.06l-1.72 1.72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpeakerXmarkTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.048 3.062a.75.75 0 0 1 .452.688v12.5a.75.75 0 0 1-1.264.546L5.203 13H2.667a.75.75 0 0 1-.7-.48A6.985 6.985 0 0 1 1.5 10c0-.887.165-1.737.468-2.52a.75.75 0 0 1 .7-.48h2.535l4.033-3.796a.75.75 0 0 1 .811-.142M13.78 7.22a.75.75 0 1 0-1.06 1.06L14.44 10l-1.72 1.72a.75.75 0 0 0 1.06 1.06l1.72-1.72l1.72 1.72a.75.75 0 1 0 1.06-1.06L16.56 10l1.72-1.72a.75.75 0 0 0-1.06-1.06L15.5 8.94z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareThreeStackThreeD(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.429 9.75L2.25 12l4.179 2.25m0-4.5l5.571 3l5.571-3m-11.142 0L2.25 7.5L12 2.25l9.75 5.25l-4.179 2.25m0 0L21.75 12l-4.179 2.25m0 0l4.179 2.25L12 21.75L2.25 16.5l4.179-2.25m11.142 0l-5.571 3l-5.571-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareThreeStackThreeDsixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.628 1.099a.75.75 0 0 1 .744 0l5.25 3a.75.75 0 0 1 0 1.302l-5.25 3a.75.75 0 0 1-.744 0l-5.25-3a.75.75 0 0 1 0-1.302z"/><path fill="currentColor" d="m2.57 7.24l-.192.11a.75.75 0 0 0 0 1.302l5.25 3a.75.75 0 0 0 .744 0l5.25-3a.75.75 0 0 0 0-1.303l-.192-.11l-4.314 2.465a2.25 2.25 0 0 1-2.232 0z"/><path fill="currentColor" d="m2.378 10.6l.192-.11l4.314 2.464a2.25 2.25 0 0 0 2.232 0l4.314-2.465l.192.11a.75.75 0 0 1 0 1.303l-5.25 3a.75.75 0 0 1-.744 0l-5.25-3a.75.75 0 0 1 0-1.303"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareThreeStackThreeDsolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M11.644 1.59a.75.75 0 0 1 .712 0l9.75 5.25a.75.75 0 0 1 0 1.32l-9.75 5.25a.75.75 0 0 1-.712 0l-9.75-5.25a.75.75 0 0 1 0-1.32z"/><path d="m3.265 10.602l7.668 4.129a2.25 2.25 0 0 0 2.134 0l7.668-4.13l1.37.739a.75.75 0 0 1 0 1.32l-9.75 5.25a.75.75 0 0 1-.71 0l-9.75-5.25a.75.75 0 0 1 0-1.32z"/><path d="m10.933 19.231l-7.668-4.13l-1.37.739a.75.75 0 0 0 0 1.32l9.75 5.25c.221.12.489.12.71 0l9.75-5.25a.75.75 0 0 0 0-1.32l-1.37-.738l-7.668 4.13a2.25 2.25 0 0 1-2.134-.001"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareThreeStackThreeDtwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m3.196 12.87l-.825.483a.75.75 0 0 0 0 1.294l7.25 4.25a.75.75 0 0 0 .758 0l7.25-4.25a.75.75 0 0 0 0-1.294l-.825-.484l-5.666 3.322a2.25 2.25 0 0 1-2.276 0z"/><path d="m3.196 8.87l-.825.483a.75.75 0 0 0 0 1.294l7.25 4.25a.75.75 0 0 0 .758 0l7.25-4.25a.75.75 0 0 0 0-1.294l-.825-.484l-5.666 3.322a2.25 2.25 0 0 1-2.276 0z"/><path d="M10.38 1.103a.75.75 0 0 0-.76 0l-7.25 4.25a.75.75 0 0 0 0 1.294l7.25 4.25a.75.75 0 0 0 .76 0l7.25-4.25a.75.75 0 0 0 0-1.294z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareTwoStack(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.5 8.25V6a2.25 2.25 0 0 0-2.25-2.25H6A2.25 2.25 0 0 0 3.75 6v8.25A2.25 2.25 0 0 0 6 16.5h2.25m8.25-8.25H18a2.25 2.25 0 0 1 2.25 2.25V18A2.25 2.25 0 0 1 18 20.25h-7.5A2.25 2.25 0 0 1 8.25 18v-1.5m8.25-8.25h-6a2.25 2.25 0 0 0-2.25 2.25v6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareTwoStackSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M5 6.5A1.5 1.5 0 0 1 6.5 5h6A1.5 1.5 0 0 1 14 6.5v6a1.5 1.5 0 0 1-1.5 1.5h-6A1.5 1.5 0 0 1 5 12.5z"/><path d="M3.5 2A1.5 1.5 0 0 0 2 3.5v6A1.5 1.5 0 0 0 3.5 11V6.5a3 3 0 0 1 3-3H11A1.5 1.5 0 0 0 9.5 2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareTwoStackSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M16.5 6a3 3 0 0 0-3-3H6a3 3 0 0 0-3 3v7.5a3 3 0 0 0 3 3v-6A4.5 4.5 0 0 1 10.5 6z"/><path d="M18 7.5a3 3 0 0 1 3 3V18a3 3 0 0 1-3 3h-7.5a3 3 0 0 1-3-3v-7.5a3 3 0 0 1 3-3z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareTwoStackTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M2 4.25A2.25 2.25 0 0 1 4.25 2h6.5A2.25 2.25 0 0 1 13 4.25V5.5H9.25A3.75 3.75 0 0 0 5.5 9.25V13H4.25A2.25 2.25 0 0 1 2 10.75z"/><path d="M9.25 7A2.25 2.25 0 0 0 7 9.25v6.5A2.25 2.25 0 0 0 9.25 18h6.5A2.25 2.25 0 0 0 18 15.75v-6.5A2.25 2.25 0 0 0 15.75 7z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquaresPlus(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.5 16.875h3.375m0 0h3.375m-3.375 0V13.5m0 3.375v3.375M6 10.5h2.25a2.25 2.25 0 0 0 2.25-2.25V6a2.25 2.25 0 0 0-2.25-2.25H6A2.25 2.25 0 0 0 3.75 6v2.25A2.25 2.25 0 0 0 6 10.5m0 9.75h2.25A2.25 2.25 0 0 0 10.5 18v-2.25a2.25 2.25 0 0 0-2.25-2.25H6a2.25 2.25 0 0 0-2.25 2.25V18A2.25 2.25 0 0 0 6 20.25m9.75-9.75H18a2.25 2.25 0 0 0 2.25-2.25V6A2.25 2.25 0 0 0 18 3.75h-2.25A2.25 2.25 0 0 0 13.5 6v2.25a2.25 2.25 0 0 0 2.25 2.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquaresPlusSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 3.5A1.5 1.5 0 0 1 3.5 2h2A1.5 1.5 0 0 1 7 3.5v2A1.5 1.5 0 0 1 5.5 7h-2A1.5 1.5 0 0 1 2 5.5zm0 7A1.5 1.5 0 0 1 3.5 9h2A1.5 1.5 0 0 1 7 10.5v2A1.5 1.5 0 0 1 5.5 14h-2A1.5 1.5 0 0 1 2 12.5zM10.5 2A1.5 1.5 0 0 0 9 3.5v2A1.5 1.5 0 0 0 10.5 7h2A1.5 1.5 0 0 0 14 5.5v-2A1.5 1.5 0 0 0 12.5 2zm1 7a.75.75 0 0 1 .75.75v1h1a.75.75 0 0 1 0 1.5h-1v1a.75.75 0 0 1-1.5 0v-1h-1a.75.75 0 0 1 0-1.5h1v-1A.75.75 0 0 1 11.5 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquaresPlusSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3a3 3 0 0 0-3 3v2.25a3 3 0 0 0 3 3h2.25a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3zm9.75 0a3 3 0 0 0-3 3v2.25a3 3 0 0 0 3 3H18a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3zM6 12.75a3 3 0 0 0-3 3V18a3 3 0 0 0 3 3h2.25a3 3 0 0 0 3-3v-2.25a3 3 0 0 0-3-3zm11.625.75a.75.75 0 0 0-1.5 0v2.625H13.5a.75.75 0 0 0 0 1.5h2.625v2.625a.75.75 0 0 0 1.5 0v-2.625h2.625a.75.75 0 0 0 0-1.5h-2.625z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquaresPlusTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 4.25A2.25 2.25 0 0 1 4.25 2h2.5A2.25 2.25 0 0 1 9 4.25v2.5A2.25 2.25 0 0 1 6.75 9h-2.5A2.25 2.25 0 0 1 2 6.75zm0 9A2.25 2.25 0 0 1 4.25 11h2.5A2.25 2.25 0 0 1 9 13.25v2.5A2.25 2.25 0 0 1 6.75 18h-2.5A2.25 2.25 0 0 1 2 15.75zm9-9A2.25 2.25 0 0 1 13.25 2h2.5A2.25 2.25 0 0 1 18 4.25v2.5A2.25 2.25 0 0 1 15.75 9h-2.5A2.25 2.25 0 0 1 11 6.75zm4.25 7.5a.75.75 0 0 0-1.5 0v2h-2a.75.75 0 0 0 0 1.5h2v2a.75.75 0 0 0 1.5 0v-2h2a.75.75 0 0 0 0-1.5h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquaresTwoXtwo(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 6A2.25 2.25 0 0 1 6 3.75h2.25A2.25 2.25 0 0 1 10.5 6v2.25a2.25 2.25 0 0 1-2.25 2.25H6a2.25 2.25 0 0 1-2.25-2.25zm0 9.75A2.25 2.25 0 0 1 6 13.5h2.25a2.25 2.25 0 0 1 2.25 2.25V18a2.25 2.25 0 0 1-2.25 2.25H6A2.25 2.25 0 0 1 3.75 18zM13.5 6a2.25 2.25 0 0 1 2.25-2.25H18A2.25 2.25 0 0 1 20.25 6v2.25A2.25 2.25 0 0 1 18 10.5h-2.25a2.25 2.25 0 0 1-2.25-2.25zm0 9.75a2.25 2.25 0 0 1 2.25-2.25H18a2.25 2.25 0 0 1 2.25 2.25V18A2.25 2.25 0 0 1 18 20.25h-2.25A2.25 2.25 0 0 1 13.5 18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquaresTwoXtwoSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 2A1.5 1.5 0 0 0 2 3.5v2A1.5 1.5 0 0 0 3.5 7h2A1.5 1.5 0 0 0 7 5.5v-2A1.5 1.5 0 0 0 5.5 2zm0 7A1.5 1.5 0 0 0 2 10.5v2A1.5 1.5 0 0 0 3.5 14h2A1.5 1.5 0 0 0 7 12.5v-2A1.5 1.5 0 0 0 5.5 9zM9 3.5A1.5 1.5 0 0 1 10.5 2h2A1.5 1.5 0 0 1 14 3.5v2A1.5 1.5 0 0 1 12.5 7h-2A1.5 1.5 0 0 1 9 5.5zM10.5 9A1.5 1.5 0 0 0 9 10.5v2a1.5 1.5 0 0 0 1.5 1.5h2a1.5 1.5 0 0 0 1.5-1.5v-2A1.5 1.5 0 0 0 12.5 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquaresTwoXtwoSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 6a3 3 0 0 1 3-3h2.25a3 3 0 0 1 3 3v2.25a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm9.75 0a3 3 0 0 1 3-3H18a3 3 0 0 1 3 3v2.25a3 3 0 0 1-3 3h-2.25a3 3 0 0 1-3-3zM3 15.75a3 3 0 0 1 3-3h2.25a3 3 0 0 1 3 3V18a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm9.75 0a3 3 0 0 1 3-3H18a3 3 0 0 1 3 3V18a3 3 0 0 1-3 3h-2.25a3 3 0 0 1-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquaresTwoXtwoTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.25 2A2.25 2.25 0 0 0 2 4.25v2.5A2.25 2.25 0 0 0 4.25 9h2.5A2.25 2.25 0 0 0 9 6.75v-2.5A2.25 2.25 0 0 0 6.75 2zm0 9A2.25 2.25 0 0 0 2 13.25v2.5A2.25 2.25 0 0 0 4.25 18h2.5A2.25 2.25 0 0 0 9 15.75v-2.5A2.25 2.25 0 0 0 6.75 11zm9-9A2.25 2.25 0 0 0 11 4.25v2.5A2.25 2.25 0 0 0 13.25 9h2.5A2.25 2.25 0 0 0 18 6.75v-2.5A2.25 2.25 0 0 0 15.75 2zm0 9A2.25 2.25 0 0 0 11 13.25v2.5A2.25 2.25 0 0 0 13.25 18h2.5A2.25 2.25 0 0 0 18 15.75v-2.5A2.25 2.25 0 0 0 15.75 11z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.48 3.499a.562.562 0 0 1 1.04 0l2.125 5.111a.563.563 0 0 0 .475.345l5.518.442c.499.04.701.663.321.988l-4.204 3.602a.563.563 0 0 0-.182.557l1.285 5.385a.562.562 0 0 1-.84.61l-4.725-2.885a.562.562 0 0 0-.586 0L6.982 20.54a.562.562 0 0 1-.84-.61l1.285-5.386a.562.562 0 0 0-.182-.557l-4.204-3.602a.562.562 0 0 1 .321-.988l5.518-.442a.563.563 0 0 0 .475-.345z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1.75a.75.75 0 0 1 .692.462l1.41 3.393l3.664.293a.75.75 0 0 1 .428 1.317l-2.791 2.39l.853 3.575a.75.75 0 0 1-1.12.814L7.998 12.08l-3.135 1.915a.75.75 0 0 1-1.12-.814l.852-3.574l-2.79-2.39a.75.75 0 0 1 .427-1.318l3.663-.293l1.41-3.393A.75.75 0 0 1 8 1.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.788 3.21c.448-1.077 1.976-1.077 2.424 0l2.082 5.006l5.404.434c1.164.093 1.636 1.545.749 2.305l-4.117 3.527l1.257 5.273c.271 1.136-.964 2.033-1.96 1.425L12 18.354L7.373 21.18c-.996.608-2.231-.29-1.96-1.425l1.257-5.273l-4.117-3.527c-.887-.76-.415-2.212.749-2.305l5.404-.434z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.868 2.884c-.321-.772-1.415-.772-1.736 0l-1.83 4.401l-4.753.381c-.833.067-1.171 1.107-.536 1.651l3.62 3.102l-1.106 4.637c-.194.813.691 1.456 1.405 1.02L10 15.591l4.069 2.485c.713.436 1.598-.207 1.404-1.02l-1.106-4.637l3.62-3.102c.635-.544.297-1.584-.536-1.65l-4.752-.382z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stop(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 7.5A2.25 2.25 0 0 1 7.5 5.25h9a2.25 2.25 0 0 1 2.25 2.25v9a2.25 2.25 0 0 1-2.25 2.25h-9a2.25 2.25 0 0 1-2.25-2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopCircle(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/><path d="M9 9.563C9 9.252 9.252 9 9.563 9h4.874c.311 0 .563.252.563.563v4.874a.563.563 0 0 1-.562.563H9.561A.562.562 0 0 1 9 14.438z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopCircleSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14M6.5 5.5a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopCircleSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 12c0-5.385 4.365-9.75 9.75-9.75s9.75 4.365 9.75 9.75s-4.365 9.75-9.75 9.75S2.25 17.385 2.25 12m6-2.437c0-.725.588-1.313 1.313-1.313h4.874c.725 0 1.313.588 1.313 1.313v4.874c0 .725-.588 1.313-1.312 1.313H9.561a1.312 1.312 0 0 1-1.312-1.312z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopCircleTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 10a8 8 0 1 1 16 0a8 8 0 0 1-16 0m5-2.25A.75.75 0 0 1 7.75 7h4.5a.75.75 0 0 1 .75.75v4.5a.75.75 0 0 1-.75.75h-4.5a.75.75 0 0 1-.75-.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<rect width="10" height="10" x="3" y="3" fill="currentColor" rx="1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 7.5a3 3 0 0 1 3-3h9a3 3 0 0 1 3 3v9a3 3 0 0 1-3 3h-9a3 3 0 0 1-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.25 3A2.25 2.25 0 0 0 3 5.25v9.5A2.25 2.25 0 0 0 5.25 17h9.5A2.25 2.25 0 0 0 17 14.75v-9.5A2.25 2.25 0 0 0 14.75 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sun(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 3v2.25m6.364.386l-1.591 1.591M21 12h-2.25m-.386 6.364l-1.591-1.591M12 18.75V21m-4.773-4.227l-1.591 1.591M5.25 12H3m4.227-4.773L5.636 5.636M15.75 12a3.75 3.75 0 1 1-7.5 0a3.75 3.75 0 0 1 7.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1a.75.75 0 0 1 .75.75v1.5a.75.75 0 0 1-1.5 0v-1.5A.75.75 0 0 1 8 1m2.5 7a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0m2.45-3.89a.75.75 0 1 0-1.06-1.06l-1.062 1.06a.75.75 0 0 0 1.061 1.062zM15 8a.75.75 0 0 1-.75.75h-1.5a.75.75 0 0 1 0-1.5h1.5A.75.75 0 0 1 15 8m-3.11 4.95a.75.75 0 0 0 1.06-1.06l-1.06-1.062a.75.75 0 0 0-1.062 1.061zM8 12a.75.75 0 0 1 .75.75v1.5a.75.75 0 0 1-1.5 0v-1.5A.75.75 0 0 1 8 12m-2.828-.11a.75.75 0 0 0-1.061-1.062L3.05 11.89a.75.75 0 1 0 1.06 1.06zM4 8a.75.75 0 0 1-.75.75h-1.5a.75.75 0 0 1 0-1.5h1.5A.75.75 0 0 1 4 8m.11-2.828A.75.75 0 0 0 5.173 4.11L4.11 3.05a.75.75 0 1 0-1.06 1.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2.25a.75.75 0 0 1 .75.75v2.25a.75.75 0 0 1-1.5 0V3a.75.75 0 0 1 .75-.75M7.5 12a4.5 4.5 0 1 1 9 0a4.5 4.5 0 0 1-9 0m11.394-5.834a.75.75 0 0 0-1.06-1.06l-1.591 1.59a.75.75 0 1 0 1.06 1.061zM21.75 12a.75.75 0 0 1-.75.75h-2.25a.75.75 0 0 1 0-1.5H21a.75.75 0 0 1 .75.75m-3.916 6.894a.75.75 0 0 0 1.06-1.06l-1.59-1.591a.75.75 0 1 0-1.061 1.06zM12 18a.75.75 0 0 1 .75.75V21a.75.75 0 0 1-1.5 0v-2.25A.75.75 0 0 1 12 18m-4.242-.697a.75.75 0 0 0-1.061-1.06l-1.591 1.59a.75.75 0 0 0 1.06 1.061zM6 12a.75.75 0 0 1-.75.75H3a.75.75 0 0 1 0-1.5h2.25A.75.75 0 0 1 6 12m.697-4.243a.75.75 0 0 0 1.06-1.06l-1.59-1.591a.75.75 0 0 0-1.061 1.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 2a.75.75 0 0 1 .75.75v1.5a.75.75 0 0 1-1.5 0v-1.5A.75.75 0 0 1 10 2m0 13a.75.75 0 0 1 .75.75v1.5a.75.75 0 0 1-1.5 0v-1.5A.75.75 0 0 1 10 15m0-8a3 3 0 1 0 0 6a3 3 0 0 0 0-6m5.657-1.596a.75.75 0 1 0-1.06-1.06l-1.061 1.06a.75.75 0 0 0 1.06 1.06zm-9.193 9.192a.75.75 0 1 0-1.06-1.06l-1.06 1.06a.75.75 0 0 0 1.06 1.06zM18 10a.75.75 0 0 1-.75.75h-1.5a.75.75 0 0 1 0-1.5h1.5A.75.75 0 0 1 18 10M5 10a.75.75 0 0 1-.75.75h-1.5a.75.75 0 0 1 0-1.5h1.5A.75.75 0 0 1 5 10m9.596 5.657a.75.75 0 0 0 1.06-1.06l-1.06-1.061a.75.75 0 1 0-1.06 1.06zM5.404 6.464a.75.75 0 0 0 1.06-1.06l-1.06-1.06a.75.75 0 1 0-1.061 1.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Swatch(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.098 19.902a3.75 3.75 0 0 0 5.304 0l6.401-6.402M6.75 21A3.75 3.75 0 0 1 3 17.25V4.125C3 3.504 3.504 3 4.125 3h5.25c.621 0 1.125.504 1.125 1.125v4.072M6.75 21a3.75 3.75 0 0 0 3.75-3.75V8.197M6.75 21h13.125c.621 0 1.125-.504 1.125-1.125v-5.25c0-.621-.504-1.125-1.125-1.125h-4.072M10.5 8.197l2.88-2.88a1.124 1.124 0 0 1 1.59 0l3.712 3.713c.44.44.44 1.152 0 1.59l-2.879 2.88M6.75 17.25h.008v.008H6.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwatchSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 3a1 1 0 0 1 1-1h3a1 1 0 0 1 1 1v8.5a2.5 2.5 0 0 1-5 0zm3.25 8.5a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0" clip-rule="evenodd"/><path fill="currentColor" d="m8.5 11.035l3.778-3.778a1 1 0 0 0 0-1.414l-2.122-2.121a1 1 0 0 0-1.414 0l-.242.242zM7.656 14H13a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1h-.344z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwatchSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M2.25 4.125c0-1.036.84-1.875 1.875-1.875h5.25c1.036 0 1.875.84 1.875 1.875V17.25a4.5 4.5 0 1 1-9 0zm4.5 14.25a1.125 1.125 0 1 0 0-2.25a1.125 1.125 0 0 0 0 2.25" clip-rule="evenodd"/><path d="M10.719 21.75h9.156c1.036 0 1.875-.84 1.875-1.875v-5.25c0-1.036-.84-1.875-1.875-1.875h-.14l-8.742 8.743a6 6 0 0 1-.274.257m2.019-4.124l6.474-6.475a1.875 1.875 0 0 0 0-2.651L15.5 4.787a1.875 1.875 0 0 0-2.651 0l-.1.099V17.25c0 .126-.003.251-.01.375"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwatchTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 2A1.5 1.5 0 0 0 2 3.5V15a3 3 0 1 0 6 0V3.5A1.5 1.5 0 0 0 6.5 2zm11.753 6.99L9.5 14.743V6.257l1.51-1.51a1.5 1.5 0 0 1 2.122 0l2.121 2.121a1.5 1.5 0 0 1 0 2.122M8.364 18H16.5a1.5 1.5 0 0 0 1.5-1.5v-3a1.5 1.5 0 0 0-1.5-1.5h-2.136zM5 16a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableCells(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.375 19.5h17.25m-17.25 0a1.125 1.125 0 0 1-1.125-1.125M3.375 19.5h7.5c.621 0 1.125-.504 1.125-1.125m-9.75 0V5.625m0 12.75v-1.5c0-.621.504-1.125 1.125-1.125m18.375 2.625V5.625m0 12.75c0 .621-.504 1.125-1.125 1.125m1.125-1.125v-1.5c0-.621-.504-1.125-1.125-1.125m0 3.75h-7.5A1.125 1.125 0 0 1 12 18.375m9.75-12.75c0-.621-.504-1.125-1.125-1.125H3.375c-.621 0-1.125.504-1.125 1.125m19.5 0v1.5c0 .621-.504 1.125-1.125 1.125M2.25 5.625v1.5c0 .621.504 1.125 1.125 1.125m0 0h17.25m-17.25 0h7.5c.621 0 1.125.504 1.125 1.125M3.375 8.25c-.621 0-1.125.504-1.125 1.125v1.5c0 .621.504 1.125 1.125 1.125m17.25-3.75h-7.5c-.621 0-1.125.504-1.125 1.125m8.625-1.125c.621 0 1.125.504 1.125 1.125v1.5c0 .621-.504 1.125-1.125 1.125m-17.25 0h7.5m-7.5 0c-.621 0-1.125.504-1.125 1.125v1.5c0 .621.504 1.125 1.125 1.125M12 10.875v-1.5m0 1.5c0 .621-.504 1.125-1.125 1.125M12 10.875c0 .621.504 1.125 1.125 1.125m-2.25 0c.621 0 1.125.504 1.125 1.125M13.125 12h7.5m-7.5 0c-.621 0-1.125.504-1.125 1.125M20.625 12c.621 0 1.125.504 1.125 1.125v1.5c0 .621-.504 1.125-1.125 1.125m-17.25 0h7.5M12 14.625v-1.5m0 1.5c0 .621-.504 1.125-1.125 1.125M12 14.625c0 .621.504 1.125 1.125 1.125m-2.25 0c.621 0 1.125.504 1.125 1.125m0 1.5v-1.5m0 0c0-.621.504-1.125 1.125-1.125m0 0h7.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableCellsSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 11a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2zM7.25 7.5a.5.5 0 0 0-.5-.5H3a.5.5 0 0 0-.5.5V8a.5.5 0 0 0 .5.5h3.75a.5.5 0 0 0 .5-.5zm1.5 3a.5.5 0 0 1 .5-.5H13a.5.5 0 0 1 .5.5v.5a.5.5 0 0 1-.5.5H9.25a.5.5 0 0 1-.5-.5zM13.5 8v-.5A.5.5 0 0 0 13 7H9.25a.5.5 0 0 0-.5.5V8a.5.5 0 0 0 .5.5H13a.5.5 0 0 0 .5-.5m-6.75 3.5a.5.5 0 0 0 .5-.5v-.5a.5.5 0 0 0-.5-.5H3a.5.5 0 0 0-.5.5v.5a.5.5 0 0 0 .5.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableCellsSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 5.625c0-1.036.84-1.875 1.875-1.875h17.25c1.035 0 1.875.84 1.875 1.875v12.75c0 1.035-.84 1.875-1.875 1.875H3.375A1.875 1.875 0 0 1 1.5 18.375zM21 9.375A.375.375 0 0 0 20.625 9h-7.5a.375.375 0 0 0-.375.375v1.5c0 .207.168.375.375.375h7.5a.375.375 0 0 0 .375-.375zm0 3.75a.375.375 0 0 0-.375-.375h-7.5a.375.375 0 0 0-.375.375v1.5c0 .207.168.375.375.375h7.5a.375.375 0 0 0 .375-.375zm0 3.75a.375.375 0 0 0-.375-.375h-7.5a.375.375 0 0 0-.375.375v1.5c0 .207.168.375.375.375h7.5a.375.375 0 0 0 .375-.375zM10.875 18.75a.375.375 0 0 0 .375-.375v-1.5a.375.375 0 0 0-.375-.375h-7.5a.375.375 0 0 0-.375.375v1.5c0 .207.168.375.375.375zM3.375 15h7.5a.375.375 0 0 0 .375-.375v-1.5a.375.375 0 0 0-.375-.375h-7.5a.375.375 0 0 0-.375.375v1.5c0 .207.168.375.375.375m0-3.75h7.5a.375.375 0 0 0 .375-.375v-1.5A.375.375 0 0 0 10.875 9h-7.5A.375.375 0 0 0 3 9.375v1.5c0 .207.168.375.375.375" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableCellsTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M.99 5.24A2.25 2.25 0 0 1 3.25 3h13.5A2.25 2.25 0 0 1 19 5.25l.01 9.5A2.25 2.25 0 0 1 16.76 17H3.26A2.267 2.267 0 0 1 1 14.74zm8.26 9.52v-.625a.75.75 0 0 0-.75-.75H3.25a.75.75 0 0 0-.75.75v.615c0 .414.336.75.75.75h5.373a.75.75 0 0 0 .627-.74m1.5 0a.75.75 0 0 0 .627.74h5.373a.75.75 0 0 0 .75-.75v-.615a.75.75 0 0 0-.75-.75H11.5a.75.75 0 0 0-.75.75zm6.75-3.63v-.625a.75.75 0 0 0-.75-.75H11.5a.75.75 0 0 0-.75.75v.625c0 .414.336.75.75.75h5.25a.75.75 0 0 0 .75-.75m-8.25 0v-.625a.75.75 0 0 0-.75-.75H3.25a.75.75 0 0 0-.75.75v.625c0 .414.336.75.75.75H8.5a.75.75 0 0 0 .75-.75M17.5 7.5v-.625a.75.75 0 0 0-.75-.75H11.5a.75.75 0 0 0-.75.75V7.5c0 .414.336.75.75.75h5.25a.75.75 0 0 0 .75-.75m-8.25 0v-.625a.75.75 0 0 0-.75-.75H3.25a.75.75 0 0 0-.75.75V7.5c0 .414.336.75.75.75H8.5a.75.75 0 0 0 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.568 3H5.25A2.25 2.25 0 0 0 3 5.25v4.318c0 .597.237 1.17.659 1.591l9.581 9.581c.699.699 1.78.872 2.607.33a18.095 18.095 0 0 0 5.224-5.223c.54-.827.368-1.908-.33-2.607l-9.583-9.58A2.25 2.25 0 0 0 9.568 3"/><path d="M6 6h.008v.008H6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 2A2.5 2.5 0 0 0 2 4.5v2.879a2.5 2.5 0 0 0 .732 1.767l4.5 4.5a2.5 2.5 0 0 0 3.536 0l2.878-2.878a2.5 2.5 0 0 0 0-3.536l-4.5-4.5A2.5 2.5 0 0 0 7.38 2zM5 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.25 2.25a3 3 0 0 0-3 3v4.318a3 3 0 0 0 .879 2.121l9.58 9.581c.92.92 2.39 1.186 3.548.428a18.849 18.849 0 0 0 5.441-5.44c.758-1.16.492-2.629-.428-3.548l-9.58-9.581a3 3 0 0 0-2.122-.879zM6.375 7.5a1.125 1.125 0 1 0 0-2.25a1.125 1.125 0 0 0 0 2.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 2A2.5 2.5 0 0 0 2 4.5v3.879a2.5 2.5 0 0 0 .732 1.767l7.5 7.5a2.5 2.5 0 0 0 3.536 0l3.878-3.878a2.5 2.5 0 0 0 0-3.536l-7.5-7.5A2.5 2.5 0 0 0 8.38 2zM5 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ticket(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.5 6v.75m0 3v.75m0 3v.75m0 3V18m-9-5.25h5.25M7.5 15h3M3.375 5.25c-.621 0-1.125.504-1.125 1.125v3.026a2.999 2.999 0 0 1 0 5.198v3.026c0 .621.504 1.125 1.125 1.125h17.25c.621 0 1.125-.504 1.125-1.125v-3.026a2.999 2.999 0 0 1 0-5.198V6.375c0-.621-.504-1.125-1.125-1.125z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TicketSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 4.5A1.5 1.5 0 0 1 2.5 3h11A1.5 1.5 0 0 1 15 4.5v1c0 .276-.227.494-.495.562a2 2 0 0 0 0 3.876c.268.068.495.286.495.562v1a1.5 1.5 0 0 1-1.5 1.5h-11A1.5 1.5 0 0 1 1 11.5v-1c0-.276.227-.494.495-.562a2 2 0 0 0 0-3.876C1.227 5.994 1 5.776 1 5.5zm9 1.25a.75.75 0 0 1 1.5 0v1a.75.75 0 0 1-1.5 0zm.75 2.75a.75.75 0 0 0-.75.75v1a.75.75 0 0 0 1.5 0v-1a.75.75 0 0 0-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TicketSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 6.375c0-1.036.84-1.875 1.875-1.875h17.25c1.035 0 1.875.84 1.875 1.875v3.026a.75.75 0 0 1-.375.65a2.249 2.249 0 0 0 0 3.898a.75.75 0 0 1 .375.65v3.026c0 1.035-.84 1.875-1.875 1.875H3.375A1.875 1.875 0 0 1 1.5 17.625v-3.026a.75.75 0 0 1 .374-.65a2.249 2.249 0 0 0 0-3.898a.75.75 0 0 1-.374-.65zm15-1.125a.75.75 0 0 1 .75.75v.75a.75.75 0 0 1-1.5 0V6a.75.75 0 0 1 .75-.75m.75 4.5a.75.75 0 0 0-1.5 0v.75a.75.75 0 0 0 1.5 0zm-.75 3a.75.75 0 0 1 .75.75v.75a.75.75 0 0 1-1.5 0v-.75a.75.75 0 0 1 .75-.75m.75 4.5a.75.75 0 0 0-1.5 0V18a.75.75 0 0 0 1.5 0zM6 12a.75.75 0 0 1 .75-.75H12a.75.75 0 0 1 0 1.5H6.75A.75.75 0 0 1 6 12m.75 2.25a.75.75 0 0 0 0 1.5h3a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TicketTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.75 3A2.25 2.25 0 0 1 18 5.25v1.214c0 .423-.277.788-.633 1.019A2.997 2.997 0 0 0 16 10c0 1.055.544 1.982 1.367 2.517c.356.231.633.596.633 1.02v1.213A2.25 2.25 0 0 1 15.75 17H4.25A2.25 2.25 0 0 1 2 14.75v-1.213c0-.424.277-.789.633-1.02A2.998 2.998 0 0 0 4 10a2.997 2.997 0 0 0-1.367-2.517C2.277 7.252 2 6.887 2 6.463V5.25A2.25 2.25 0 0 1 4.25 3zM13.5 7.396a.75.75 0 0 0-1.5 0v1.042a.75.75 0 0 0 1.5 0zm0 4.167a.75.75 0 0 0-1.5 0v1.041a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m14.74 9l-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 0 1-2.244 2.077H8.084a2.25 2.25 0 0 1-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 0 0-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 0 1 3.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 0 0-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 0 0-7.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrashSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 3.25V4H2.75a.75.75 0 0 0 0 1.5h.3l.815 8.15A1.5 1.5 0 0 0 5.357 15h5.286a1.5 1.5 0 0 0 1.492-1.35l.815-8.15h.3a.75.75 0 0 0 0-1.5H11v-.75A2.25 2.25 0 0 0 8.75 1h-1.5A2.25 2.25 0 0 0 5 3.25m2.25-.75a.75.75 0 0 0-.75.75V4h3v-.75a.75.75 0 0 0-.75-.75zM6.05 6a.75.75 0 0 1 .787.713l.275 5.5a.75.75 0 0 1-1.498.075l-.275-5.5A.75.75 0 0 1 6.05 6m3.9 0a.75.75 0 0 1 .712.787l-.275 5.5a.75.75 0 0 1-1.498-.075l.275-5.5a.75.75 0 0 1 .786-.711" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrashSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.5 4.478v.227a48.816 48.816 0 0 1 3.878.512a.75.75 0 1 1-.256 1.478l-.209-.035l-1.005 13.07a3 3 0 0 1-2.991 2.77H8.084a3 3 0 0 1-2.991-2.77L4.087 6.66l-.209.035a.75.75 0 0 1-.256-1.478A48.567 48.567 0 0 1 7.5 4.705v-.227c0-1.564 1.213-2.9 2.816-2.951a52.662 52.662 0 0 1 3.369 0c1.603.051 2.815 1.387 2.815 2.951m-6.136-1.452a51.196 51.196 0 0 1 3.273 0C14.39 3.05 15 3.684 15 4.478v.113a49.488 49.488 0 0 0-6 0v-.113c0-.794.609-1.428 1.364-1.452m-.355 5.945a.75.75 0 1 0-1.5.058l.347 9a.75.75 0 1 0 1.499-.058zm5.48.058a.75.75 0 1 0-1.498-.058l-.347 9a.75.75 0 0 0 1.5.058z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrashTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.75 1A2.75 2.75 0 0 0 6 3.75v.443a40.83 40.83 0 0 0-2.365.298a.75.75 0 1 0 .23 1.482l.149-.022l.841 10.518A2.75 2.75 0 0 0 7.596 19h4.807a2.75 2.75 0 0 0 2.742-2.53l.841-10.52l.149.023a.75.75 0 0 0 .23-1.482A41.03 41.03 0 0 0 14 4.193V3.75A2.75 2.75 0 0 0 11.25 1zM10 4c.84 0 1.673.025 2.5.075V3.75c0-.69-.56-1.25-1.25-1.25h-2.5c-.69 0-1.25.56-1.25 1.25v.325C8.327 4.025 9.16 4 10 4M8.58 7.72a.75.75 0 0 0-1.5.06l.3 7.5a.75.75 0 1 0 1.5-.06zm4.34.06a.75.75 0 1 0-1.5-.06l-.3 7.5a.75.75 0 1 0 1.5.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trophy(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.5 18.75h-9m9 0a3 3 0 0 1 3 3h-15a3 3 0 0 1 3-3m9 0v-3.375c0-.621-.503-1.125-1.125-1.125h-.871M7.5 18.75v-3.375c0-.621.504-1.125 1.125-1.125h.872m5.007 0H9.497m5.007 0a7.454 7.454 0 0 1-.982-3.172M9.497 14.25a7.454 7.454 0 0 0 .981-3.172M5.25 4.236c-.982.143-1.954.317-2.916.52A6.003 6.003 0 0 0 7.73 9.728M5.25 4.236V4.5c0 2.108.966 3.99 2.48 5.228M5.25 4.236V2.721A48.339 48.339 0 0 1 12 2.25c2.291 0 4.545.16 6.75.47v1.516M7.73 9.728a6.726 6.726 0 0 0 2.748 1.35m8.272-6.842V4.5c0 2.108-.966 3.99-2.48 5.228m2.48-5.492a46.32 46.32 0 0 1 2.916.52a6.003 6.003 0 0 1-5.395 4.972m0 0a6.726 6.726 0 0 1-2.749 1.35m0 0a6.772 6.772 0 0 1-3.044 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrophySixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.69a.494.494 0 0 0-.438-.494a32.352 32.352 0 0 0-7.124 0A.494.494 0 0 0 4 1.689v.567c-.811.104-1.612.24-2.403.406a.75.75 0 0 0-.595.714a4.5 4.5 0 0 0 4.35 4.622A3.99 3.99 0 0 0 7 8.874V10H6a1 1 0 0 0-1 1v2h-.667C3.597 13 3 13.597 3 14.333c0 .368.298.667.667.667h8.666a.667.667 0 0 0 .667-.667c0-.736-.597-1.333-1.333-1.333H11v-2a1 1 0 0 0-1-1H9V8.874a3.99 3.99 0 0 0 1.649-.876a4.5 4.5 0 0 0 4.35-4.622a.75.75 0 0 0-.596-.714A30.897 30.897 0 0 0 12 2.256zM4 3.768c-.49.066-.976.145-1.458.235a3.004 3.004 0 0 0 1.64 2.192A3.999 3.999 0 0 1 4 5zm8 0c.49.066.976.145 1.458.235a3.004 3.004 0 0 1-1.64 2.192C11.936 5.818 12 5.416 12 5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrophySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.166 2.621v.858a47.68 47.68 0 0 0-3.071.543a.75.75 0 0 0-.584.859a6.753 6.753 0 0 0 6.138 5.6a6.73 6.73 0 0 0 2.743 1.346A6.707 6.707 0 0 1 9.279 15H8.54c-1.036 0-1.875.84-1.875 1.875V19.5h-.75a2.25 2.25 0 0 0-2.25 2.25c0 .414.336.75.75.75h15a.75.75 0 0 0 .75-.75a2.25 2.25 0 0 0-2.25-2.25h-.75v-2.625c0-1.036-.84-1.875-1.875-1.875h-.739a6.706 6.706 0 0 1-1.112-3.173a6.73 6.73 0 0 0 2.743-1.347a6.753 6.753 0 0 0 6.139-5.6a.75.75 0 0 0-.585-.858a47.077 47.077 0 0 0-3.07-.543V2.62a.75.75 0 0 0-.658-.744a49.22 49.22 0 0 0-6.093-.377c-2.063 0-4.096.128-6.093.377a.75.75 0 0 0-.657.744m0 2.629c0 1.196.312 2.32.857 3.294A5.266 5.266 0 0 1 3.16 5.337a45.6 45.6 0 0 1 2.006-.343zm13.5 0v-.256c.674.1 1.343.214 2.006.343a5.265 5.265 0 0 1-2.863 3.207a6.72 6.72 0 0 0 .857-3.294" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrophyTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 1c-1.828 0-3.623.149-5.371.435a.75.75 0 0 0-.629.74v.387c-.827.157-1.642.345-2.445.564a.75.75 0 0 0-.552.698a5 5 0 0 0 4.503 5.152a6 6 0 0 0 2.946 1.822A6.451 6.451 0 0 1 7.768 13H7.5A1.5 1.5 0 0 0 6 14.5V17h-.75C4.56 17 4 17.56 4 18.25c0 .414.336.75.75.75h10.5a.75.75 0 0 0 .75-.75c0-.69-.56-1.25-1.25-1.25H14v-2.5a1.5 1.5 0 0 0-1.5-1.5h-.268a6.453 6.453 0 0 1-.684-2.202a6 6 0 0 0 2.946-1.822a5 5 0 0 0 4.503-5.152a.75.75 0 0 0-.552-.698A31.804 31.804 0 0 0 16 2.562v-.387a.75.75 0 0 0-.629-.74A33.227 33.227 0 0 0 10 1M2.525 4.422C3.012 4.3 3.504 4.19 4 4.09V5c0 .74.134 1.448.38 2.103a3.503 3.503 0 0 1-1.855-2.68m14.95 0a3.503 3.503 0 0 1-1.854 2.68A5.984 5.984 0 0 0 16 5v-.91c.496.099.988.21 1.475.332" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Truck(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.25 18.75a1.5 1.5 0 0 1-3 0m3 0a1.5 1.5 0 0 0-3 0m3 0h6m-9 0H3.375a1.125 1.125 0 0 1-1.125-1.125V14.25m17.25 4.5a1.5 1.5 0 0 1-3 0m3 0a1.5 1.5 0 0 0-3 0m3 0h1.125c.621 0 1.129-.504 1.09-1.124a17.902 17.902 0 0 0-3.213-9.193a2.056 2.056 0 0 0-1.58-.86H14.25M16.5 18.75h-2.25m0-11.177v-.958c0-.568-.422-1.048-.987-1.106a48.554 48.554 0 0 0-10.026 0a1.106 1.106 0 0 0-.987 1.106v7.635m12-6.677v6.677m0 4.5v-4.5m0 0h-12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TruckSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.908 2.067A.978.978 0 0 0 2 3.05V8h6V3.05a.978.978 0 0 0-.908-.983a32.481 32.481 0 0 0-4.184 0m10.011 2.655A.98.98 0 0 0 11.968 4H10a1 1 0 0 0-1 1v6.268A2 2 0 0 1 12 13h1a.977.977 0 0 0 .985-1a31.99 31.99 0 0 0-1.066-7.278"/><path fill="currentColor" d="M11 13a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-9-1V9h6v3a1 1 0 0 1-1 1a2 2 0 1 0-4 0a1 1 0 0 1-1-1"/><path fill="currentColor" d="M6 13a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TruckSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M3.375 4.5C2.339 4.5 1.5 5.34 1.5 6.375V13.5h12V6.375c0-1.036-.84-1.875-1.875-1.875zM13.5 15h-12v2.625c0 1.035.84 1.875 1.875 1.875h.375a3 3 0 1 1 6 0h3a.75.75 0 0 0 .75-.75z"/><path d="M8.25 19.5a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0m7.5-12.75a.75.75 0 0 0-.75.75v11.25c0 .087.015.17.042.248a3 3 0 0 1 5.958.464c.853-.175 1.522-.935 1.464-1.883a18.659 18.659 0 0 0-3.732-10.104a1.837 1.837 0 0 0-1.47-.725z"/><path d="M19.5 19.5a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TruckTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M6.5 3c-1.051 0-2.093.04-3.125.117A1.49 1.49 0 0 0 2 4.607V10.5h9V4.606c0-.771-.59-1.43-1.375-1.489A41.568 41.568 0 0 0 6.5 3M2 12v2.5A1.5 1.5 0 0 0 3.5 16h.041a3 3 0 0 1 5.918 0h.791a.75.75 0 0 0 .75-.75V12z"/><path d="M6.5 18a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m6.75-13a.75.75 0 0 0-.75.75v8.514a3.001 3.001 0 0 1 4.893 1.44c.37-.275.61-.719.595-1.227a24.905 24.905 0 0 0-1.784-8.549A1.486 1.486 0 0 0 14.823 5zm1.25 13a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tv(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 20.25h12m-7.5-3v3m3-3v3m-10.125-3h17.25c.621 0 1.125-.504 1.125-1.125V4.875c0-.621-.504-1.125-1.125-1.125H3.375c-.621 0-1.125.504-1.125 1.125v11.25c0 .621.504 1.125 1.125 1.125"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TvSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5H4v4h8z"/><path fill="currentColor" fill-rule="evenodd" d="M1 3a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1h-4v1.5h2.25a.75.75 0 0 1 0 1.5h-8.5a.75.75 0 0 1 0-1.5H6V12H2a1 1 0 0 1-1-1zm1.5 7.5v-7h11v7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TvSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M19.5 6h-15v9h15z"/><path fill-rule="evenodd" d="M3.375 3C2.339 3 1.5 3.84 1.5 4.875v11.25C1.5 17.16 2.34 18 3.375 18H9.75v1.5H6A.75.75 0 0 0 6 21h12a.75.75 0 0 0 0-1.5h-3.75V18h6.375c1.035 0 1.875-.84 1.875-1.875V4.875C22.5 3.839 21.66 3 20.625 3zm0 13.5h17.25a.375.375 0 0 0 .375-.375V4.875a.375.375 0 0 0-.375-.375H3.375A.375.375 0 0 0 3 4.875v11.25c0 .207.168.375.375.375" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TvTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M4 5h12v7H4z"/><path fill-rule="evenodd" d="M1 3.5A1.5 1.5 0 0 1 2.5 2h15A1.5 1.5 0 0 1 19 3.5v10a1.5 1.5 0 0 1-1.5 1.5H12v1.5h3.25a.75.75 0 0 1 0 1.5H4.75a.75.75 0 0 1 0-1.5H8V15H2.5A1.5 1.5 0 0 1 1 13.5zm16.5 0h-15v10h15z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Underline(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 3.75v7.5a6 6 0 0 1-12 0v-7.5m-2.25 16.5h16.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnderlineSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 3a.75.75 0 0 1 .75.75v7.5a5.25 5.25 0 1 0 10.5 0v-7.5a.75.75 0 0 1 1.5 0v7.5a6.75 6.75 0 0 1-13.5 0v-7.5A.75.75 0 0 1 6 3M3 20.25a.75.75 0 0 1 .75-.75h16.5a.75.75 0 0 1 0 1.5H3.75a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnderlineTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.75 2a.75.75 0 0 1 .75.75V10a4.5 4.5 0 1 0 9 0V2.75a.75.75 0 0 1 1.5 0V10a6 6 0 0 1-12 0V2.75A.75.75 0 0 1 4.75 2m-2.5 16.25A.75.75 0 0 1 3 17.5h14a.75.75 0 0 1 0 1.5H3a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func User(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.75 6a3.75 3.75 0 1 1-7.5 0a3.75 3.75 0 0 1 7.5 0M4.501 20.118a7.5 7.5 0 0 1 14.998 0A17.933 17.933 0 0 1 12 21.75c-2.676 0-5.216-.584-7.499-1.632"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCircle(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.982 18.725A7.488 7.488 0 0 0 12 15.75a7.488 7.488 0 0 0-5.982 2.975m11.964 0a9 9 0 1 0-11.963 0m11.962 0A8.966 8.966 0 0 1 12 21a8.966 8.966 0 0 1-5.982-2.275M15 9.75a3 3 0 1 1-6 0a3 3 0 0 1 6 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCircleSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0m-5-2a2 2 0 1 1-4 0a2 2 0 0 1 4 0M8 9a4.998 4.998 0 0 0-4.295 2.437A5.49 5.49 0 0 0 8 13.5a5.49 5.49 0 0 0 4.294-2.063A4.997 4.997 0 0 0 8 9" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCircleSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M18.685 19.097A9.723 9.723 0 0 0 21.75 12c0-5.385-4.365-9.75-9.75-9.75S2.25 6.615 2.25 12a9.723 9.723 0 0 0 3.065 7.097A9.716 9.716 0 0 0 12 21.75a9.716 9.716 0 0 0 6.685-2.653m-12.54-1.285A7.486 7.486 0 0 1 12 15a7.486 7.486 0 0 1 5.855 2.812A8.224 8.224 0 0 1 12 20.25a8.224 8.224 0 0 1-5.855-2.438M15.75 9a3.75 3.75 0 1 1-7.5 0a3.75 3.75 0 0 1 7.5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCircleTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M18 10a8 8 0 1 1-16 0a8 8 0 0 1 16 0m-5.5-2.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0M10 12a5.99 5.99 0 0 0-4.793 2.39A6.483 6.483 0 0 0 10 16.5a6.483 6.483 0 0 0 4.793-2.11A5.99 5.99 0 0 0 10 12" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserGroup(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 18.72a9.094 9.094 0 0 0 3.741-.479a3 3 0 0 0-4.682-2.72m.94 3.198l.001.031c0 .225-.012.447-.037.666A11.944 11.944 0 0 1 12 21c-2.17 0-4.207-.576-5.963-1.584A6.062 6.062 0 0 1 6 18.719m12 0a5.971 5.971 0 0 0-.941-3.197m0 0A5.995 5.995 0 0 0 12 12.75a5.995 5.995 0 0 0-5.058 2.772m0 0a3 3 0 0 0-4.681 2.72a8.986 8.986 0 0 0 3.74.477m.94-3.197a5.971 5.971 0 0 0-.94 3.197M15 6.75a3 3 0 1 1-6 0a3 3 0 0 1 6 0m6 3a2.25 2.25 0 1 1-4.5 0a2.25 2.25 0 0 1 4.5 0m-13.5 0a2.25 2.25 0 1 1-4.5 0a2.25 2.25 0 0 1 4.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserGroupSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 8a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5m-4.844 3.763c.16-.629.44-1.21.813-1.72a2.5 2.5 0 0 0-2.725 1.377c-.136.287.102.58.418.58h1.449c.01-.077.025-.156.045-.237m9.691 0c.02.08.036.16.046.237h1.446c.316 0 .554-.293.417-.579a2.5 2.5 0 0 0-2.722-1.378c.374.51.653 1.09.813 1.72M14 7.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0M3.5 9a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3M5 13c-.552 0-1.013-.455-.876-.99a4.002 4.002 0 0 1 7.753 0c.136.535-.324.99-.877.99z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserGroupSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M8.25 6.75a3.75 3.75 0 1 1 7.5 0a3.75 3.75 0 0 1-7.5 0m7.5 3a3 3 0 1 1 6 0a3 3 0 0 1-6 0m-13.5 0a3 3 0 1 1 6 0a3 3 0 0 1-6 0m4.06 5.368A6.745 6.745 0 0 1 12 12a6.745 6.745 0 0 1 6.709 7.498a.75.75 0 0 1-.372.568A12.696 12.696 0 0 1 12 21.75a12.69 12.69 0 0 1-6.337-1.684a.75.75 0 0 1-.372-.568a6.787 6.787 0 0 1 1.019-4.38" clip-rule="evenodd"/><path d="M5.082 14.254a8.287 8.287 0 0 0-1.308 5.135a9.687 9.687 0 0 1-1.764-.44l-.115-.04a.563.563 0 0 1-.373-.487l-.01-.121a3.75 3.75 0 0 1 3.57-4.047m15.144 5.135a8.287 8.287 0 0 0-1.308-5.135a3.75 3.75 0 0 1 3.57 4.047l-.01.121a.563.563 0 0 1-.373.486l-.115.04c-.567.2-1.156.349-1.764.441"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserGroupTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 9a3 3 0 1 0 0-6a3 3 0 0 0 0 6M6 8a2 2 0 1 1-4 0a2 2 0 0 1 4 0m-4.51 7.326a.78.78 0 0 1-.358-.442a3 3 0 0 1 4.308-3.516a6.484 6.484 0 0 0-1.905 3.959c-.023.222-.014.442.025.654a4.97 4.97 0 0 1-2.07-.655m14.95.654a4.97 4.97 0 0 0 2.07-.654a.78.78 0 0 0 .357-.442a3 3 0 0 0-4.308-3.517a6.484 6.484 0 0 1 1.907 3.96a2.32 2.32 0 0 1-.026.654M18 8a2 2 0 1 1-4 0a2 2 0 0 1 4 0M5.304 16.19a.844.844 0 0 1-.277-.71a5 5 0 0 1 9.947 0a.843.843 0 0 1-.277.71A6.975 6.975 0 0 1 10 18a6.974 6.974 0 0 1-4.696-1.81"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserMinus(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 10.5h-6m-2.25-4.125a3.375 3.375 0 1 1-6.75 0a3.375 3.375 0 0 1 6.75 0M4 19.235v-.11a6.375 6.375 0 0 1 12.75 0v.109A12.318 12.318 0 0 1 10.374 21C8.043 21 5.862 20.355 4 19.234"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserMinusSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 4.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0M10 13c.552 0 1.01-.452.9-.994a5.002 5.002 0 0 0-9.802 0c-.109.542.35.994.902.994zm.75-7.75a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserMinusSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.375 2.25a4.125 4.125 0 1 0 0 8.25a4.125 4.125 0 0 0 0-8.25m0 9.75a7.125 7.125 0 0 0-7.124 7.247a.75.75 0 0 0 .363.63a13.067 13.067 0 0 0 6.761 1.873c2.472 0 4.786-.684 6.76-1.873a.75.75 0 0 0 .364-.63l.001-.12v-.002A7.125 7.125 0 0 0 10.375 12M16 9.75a.75.75 0 0 0 0 1.5h6a.75.75 0 0 0 0-1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserMinusTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 5a3 3 0 1 1-6 0a3 3 0 0 1 6 0M2.046 15.253c-.058.468.172.92.57 1.175A9.953 9.953 0 0 0 8 18c1.982 0 3.83-.578 5.384-1.573c.398-.254.628-.707.57-1.175a6.001 6.001 0 0 0-11.908 0M12.75 7.75a.75.75 0 0 0 0 1.5h5.5a.75.75 0 0 0 0-1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserPlus(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 7.5v3m0 0v3m0-3h3m-3 0h-3m-2.25-4.125a3.375 3.375 0 1 1-6.75 0a3.375 3.375 0 0 1 6.75 0M3 19.235v-.11a6.375 6.375 0 0 1 12.75 0v.109A12.318 12.318 0 0 1 9.374 21C7.043 21 4.862 20.355 3 19.234"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserPlusSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 4.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0M10 13c.552 0 1.01-.452.9-.994a5.002 5.002 0 0 0-9.802 0c-.109.542.35.994.902.994zm2.5-9.5a.75.75 0 0 1 .75.75v1h1a.75.75 0 0 1 0 1.5h-1v1a.75.75 0 0 1-1.5 0v-1h-1a.75.75 0 0 1 0-1.5h1v-1a.75.75 0 0 1 .75-.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserPlusSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.25 6.375a4.125 4.125 0 1 1 8.25 0a4.125 4.125 0 0 1-8.25 0m-3 12.75a7.125 7.125 0 0 1 14.25 0v.003l-.001.119a.75.75 0 0 1-.363.63a13.067 13.067 0 0 1-6.761 1.873c-2.472 0-4.786-.684-6.76-1.873a.75.75 0 0 1-.364-.63zM18.75 7.5a.75.75 0 0 0-1.5 0v2.25H15a.75.75 0 0 0 0 1.5h2.25v2.25a.75.75 0 0 0 1.5 0v-2.25H21a.75.75 0 0 0 0-1.5h-2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserPlusTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 5a3 3 0 1 1-6 0a3 3 0 0 1 6 0M1.615 16.428a1.224 1.224 0 0 1-.569-1.175a6.002 6.002 0 0 1 11.908 0c.058.467-.172.92-.57 1.174A9.953 9.953 0 0 1 7 18a9.953 9.953 0 0 1-5.385-1.572M16.25 5.75a.75.75 0 0 0-1.5 0v2h-2a.75.75 0 0 0 0 1.5h2v2a.75.75 0 0 0 1.5 0v-2h2a.75.75 0 0 0 0-1.5h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 8a3 3 0 1 0 0-6a3 3 0 0 0 0 6m4.735 6c.618 0 1.093-.561.872-1.139a6.002 6.002 0 0 0-11.215 0c-.22.578.254 1.139.872 1.139z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 6a4.5 4.5 0 1 1 9 0a4.5 4.5 0 0 1-9 0M3.751 20.105a8.25 8.25 0 0 1 16.498 0a.75.75 0 0 1-.437.695A18.683 18.683 0 0 1 12 22.5c-2.786 0-5.433-.608-7.812-1.7a.75.75 0 0 1-.437-.695" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 8a3 3 0 1 0 0-6a3 3 0 0 0 0 6m-6.535 6.494a1.23 1.23 0 0 0 .41 1.41A9.957 9.957 0 0 0 10 18c2.31 0 4.438-.784 6.131-2.1c.43-.333.604-.903.408-1.41a7.002 7.002 0 0 0-13.074.003"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Users(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 19.128a9.38 9.38 0 0 0 2.625.372a9.337 9.337 0 0 0 4.121-.952a4.125 4.125 0 0 0-7.533-2.493M15 19.128v-.003c0-1.113-.285-2.16-.786-3.07M15 19.128v.106A12.318 12.318 0 0 1 8.624 21c-2.331 0-4.512-.645-6.374-1.766l-.001-.109a6.375 6.375 0 0 1 11.964-3.07M12 6.375a3.375 3.375 0 1 1-6.75 0a3.375 3.375 0 0 1 6.75 0m8.25 2.25a2.625 2.625 0 1 1-5.25 0a2.625 2.625 0 0 1 5.25 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UsersSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 4.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0m2.4 7.506c.11.542-.348.994-.9.994H2c-.553 0-1.01-.452-.902-.994a5.002 5.002 0 0 1 9.803 0M14.002 12h-1.59a2.556 2.556 0 0 0-.04-.29a6.476 6.476 0 0 0-1.167-2.603a3.002 3.002 0 0 1 3.633 1.911c.18.522-.283.982-.836.982M12 8a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UsersSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 6.375a4.125 4.125 0 1 1 8.25 0a4.125 4.125 0 0 1-8.25 0m9.75 2.25a3.375 3.375 0 1 1 6.75 0a3.375 3.375 0 0 1-6.75 0M1.5 19.125a7.125 7.125 0 0 1 14.25 0v.003l-.001.119a.75.75 0 0 1-.363.63a13.067 13.067 0 0 1-6.761 1.873c-2.472 0-4.786-.684-6.76-1.873a.75.75 0 0 1-.364-.63zm15.75.003l-.001.144a2.25 2.25 0 0 1-.233.96a10.088 10.088 0 0 0 5.06-1.01a.75.75 0 0 0 .42-.643a4.875 4.875 0 0 0-6.957-4.611a8.586 8.586 0 0 1 1.71 5.157z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UsersTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 8a3 3 0 1 0 0-6a3 3 0 0 0 0 6m7.5 1a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5M1.615 16.428a1.224 1.224 0 0 1-.569-1.175a6.002 6.002 0 0 1 11.908 0c.058.467-.172.92-.57 1.174A9.953 9.953 0 0 1 7 18a9.953 9.953 0 0 1-5.385-1.572M14.5 16h-.106c.07-.297.088-.611.048-.933a7.47 7.47 0 0 0-1.588-3.755a4.502 4.502 0 0 1 5.874 2.636a.818.818 0 0 1-.36.98A7.465 7.465 0 0 1 14.5 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Variable(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.745 3A23.933 23.933 0 0 0 3 12c0 3.183.62 6.22 1.745 9M19.5 3c.967 2.78 1.5 5.817 1.5 9s-.533 6.22-1.5 9M8.25 8.885l1.444-.89a.75.75 0 0 1 1.105.402l2.402 7.206a.75.75 0 0 0 1.104.401l1.445-.889m-8.25.75l.213.09a1.687 1.687 0 0 0 2.062-.617l4.45-6.676a1.688 1.688 0 0 1 2.062-.618l.213.09"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VariableSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.38 3.012a.75.75 0 1 0-1.408-.516A15.97 15.97 0 0 0 1 8c0 1.932.343 3.786.972 5.503a.75.75 0 0 0 1.408-.516A14.47 14.47 0 0 1 2.5 8c0-1.754.311-3.434.88-4.988m9.24 0a.75.75 0 1 1 1.408-.516A15.97 15.97 0 0 1 15 8a15.97 15.97 0 0 1-.972 5.503a.75.75 0 0 1-1.408-.516c.569-1.554.88-3.233.88-4.987s-.311-3.434-.88-4.988M6.523 4.785a.75.75 0 0 1 .898.38l.758 1.515l.812-.902a2.376 2.376 0 0 1 2.486-.674a.75.75 0 1 1-.454 1.429a.876.876 0 0 0-.918.249L8.9 8.122l.734 1.468l.388-.124a.75.75 0 0 1 .457 1.428l-1 .32a.75.75 0 0 1-.899-.379L7.821 9.32l-.811.901a2.374 2.374 0 0 1-2.489.673a.75.75 0 0 1 .458-1.428a.874.874 0 0 0 .916-.248L7.1 7.878L6.366 6.41l-.389.124a.75.75 0 1 1-.454-1.43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VariableSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M19.254 2.292a.75.75 0 0 1 .954.461A28.123 28.123 0 0 1 21.75 12a28.13 28.13 0 0 1-1.542 9.247a.75.75 0 1 1-1.416-.494c.94-2.7 1.458-5.654 1.458-8.753s-.519-6.054-1.458-8.754a.75.75 0 0 1 .461-.954m-14.228.013a.75.75 0 0 1 .414.976A23.183 23.183 0 0 0 3.75 12c0 3.085.6 6.027 1.69 8.718a.75.75 0 0 1-1.39.563c-1.161-2.867-1.8-6-1.8-9.281c0-3.28.639-6.414 1.8-9.281a.75.75 0 0 1 .976-.414m4.275 5.052a1.5 1.5 0 0 1 2.21.803l.716 2.148L13.6 8.246a2.438 2.438 0 0 1 2.978-.892l.213.09a.75.75 0 1 1-.584 1.381l-.214-.09a.937.937 0 0 0-1.145.343l-2.021 3.033l1.084 3.255l1.445-.89a.75.75 0 1 1 .786 1.278l-1.444.889a1.5 1.5 0 0 1-2.21-.803l-.716-2.148l-1.374 2.062a2.437 2.437 0 0 1-2.978.892l-.213-.09a.75.75 0 0 1 .584-1.381l.214.09a.938.938 0 0 0 1.145-.344l2.021-3.032l-1.084-3.255l-1.445.89a.75.75 0 1 1-.786-1.278z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VariableTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.212 2.079a.75.75 0 0 1 1.006.336A16.932 16.932 0 0 1 18 10c0 2.724-.641 5.3-1.782 7.585a.75.75 0 1 1-1.342-.67A15.432 15.432 0 0 0 16.5 10c0-2.486-.585-4.834-1.624-6.915a.75.75 0 0 1 .336-1.006m-10.424 0a.75.75 0 0 1 .336 1.006A15.433 15.433 0 0 0 3.5 10c0 2.486.585 4.834 1.624 6.915a.75.75 0 1 1-1.342.67A16.933 16.933 0 0 1 2 10c0-2.724.641-5.3 1.782-7.585a.75.75 0 0 1 1.006-.336m2.285 3.554a1.5 1.5 0 0 1 2.219.677l.856 2.08l1.146-1.77a2.25 2.25 0 0 1 3.137-.65l.235.156a.75.75 0 1 1-.832 1.248l-.235-.156a.75.75 0 0 0-1.045.216l-1.71 2.644l1.251 3.04l.739-.492a.75.75 0 1 1 .832 1.248l-.739.493a1.5 1.5 0 0 1-2.219-.677l-.856-2.08l-1.146 1.77a2.25 2.25 0 0 1-3.137.65l-.235-.156a.75.75 0 0 1 .832-1.248l.235.157a.75.75 0 0 0 1.045-.217l1.71-2.644l-1.251-3.04l-.739.492a.75.75 0 0 1-.832-1.248z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoCamera(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15.75 10.5l4.72-4.72a.75.75 0 0 1 1.28.53v11.38a.75.75 0 0 1-1.28.53l-4.72-4.72M4.5 18.75h9a2.25 2.25 0 0 0 2.25-2.25v-9a2.25 2.25 0 0 0-2.25-2.25h-9A2.25 2.25 0 0 0 2.25 7.5v9a2.25 2.25 0 0 0 2.25 2.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoCameraSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 4a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h5a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2zm12 .75a.75.75 0 0 0-1.28-.53l-2 2a.75.75 0 0 0-.22.53v2.5c0 .199.079.39.22.53l2 2a.75.75 0 0 0 1.28-.53z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoCameraSlash(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15.75 10.5l4.72-4.72a.75.75 0 0 1 1.28.53v11.38a.75.75 0 0 1-1.28.53l-4.72-4.72M12 18.75H4.5a2.25 2.25 0 0 1-2.25-2.25V9m12.841 9.091L16.5 19.5m-1.409-1.409c.407-.407.659-.97.659-1.591v-9a2.25 2.25 0 0 0-2.25-2.25h-9c-.621 0-1.184.252-1.591.659m12.182 12.182L2.909 5.909M1.5 4.5l1.409 1.409"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoCameraSlashSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 10V6.682L6.318 12H3a2 2 0 0 1-2-2m9-4v3.318L4.682 4H8a2 2 0 0 1 2 2m4.537-1.943A.75.75 0 0 1 15 4.75v6.5a.75.75 0 0 1-1.28.53l-2-2a.75.75 0 0 1-.22-.53v-2.5a.75.75 0 0 1 .22-.53l2-2a.75.75 0 0 1 .817-.163M2.78 4.22a.75.75 0 0 0-1.06 1.06l6.5 6.5a.75.75 0 0 0 1.06-1.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoCameraSlashSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.97 3.97a.75.75 0 0 1 1.06 0l15 15a.75.75 0 1 1-1.06 1.06l-15-15a.75.75 0 0 1 0-1.06m16.28 12.09l2.69 2.69c.944.945 2.56.276 2.56-1.06V6.31c0-1.336-1.616-2.005-2.56-1.06l-2.69 2.69zm-1.5-8.56v8.068L4.682 4.5h8.068a3 3 0 0 1 3 3m-14.25 9V7.682l11.773 11.773a3.01 3.01 0 0 1-.523.045H4.5a3 3 0 0 1-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoCameraSlashTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 13.75V7.182L9.818 16H3.25A2.25 2.25 0 0 1 1 13.75m12-7.5v6.568L4.182 4h6.568A2.25 2.25 0 0 1 13 6.25m6-1.5a.75.75 0 0 0-1.28-.53l-3 3a.75.75 0 0 0-.22.53v4.5c0 .199.079.39.22.53l3 3a.75.75 0 0 0 1.28-.53zM2.28 4.22a.75.75 0 0 0-1.06 1.06l10.5 10.5a.75.75 0 1 0 1.06-1.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoCameraSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 4.5a3 3 0 0 0-3 3v9a3 3 0 0 0 3 3h8.25a3 3 0 0 0 3-3v-9a3 3 0 0 0-3-3zm15.44 14.25l-2.69-2.69V7.94l2.69-2.69c.944-.945 2.56-.276 2.56 1.06v11.38c0 1.336-1.616 2.005-2.56 1.06"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoCameraTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.25 4A2.25 2.25 0 0 0 1 6.25v7.5A2.25 2.25 0 0 0 3.25 16h7.5A2.25 2.25 0 0 0 13 13.75v-7.5A2.25 2.25 0 0 0 10.75 4zM19 4.75a.75.75 0 0 0-1.28-.53l-3 3a.75.75 0 0 0-.22.53v4.5c0 .199.079.39.22.53l3 3a.75.75 0 0 0 1.28-.53z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewColumns(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 4.5v15m6-15v15m-10.875 0h15.75c.621 0 1.125-.504 1.125-1.125V5.625c0-.621-.504-1.125-1.125-1.125H4.125C3.504 4.5 3 5.004 3 5.625v12.75c0 .621.504 1.125 1.125 1.125"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewColumnsSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.836 3h-3.67v10h3.67zm1.5 10H13.5a1.5 1.5 0 0 0 1.5-1.5v-7A1.5 1.5 0 0 0 13.5 3h-2.164zM2.5 3h2.166v10H2.5A1.5 1.5 0 0 1 1 11.5v-7A1.5 1.5 0 0 1 2.5 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewColumnsSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 3.75H9v16.5h6zm1.5 16.5h3.375c1.035 0 1.875-.84 1.875-1.875V5.625c0-1.036-.84-1.875-1.875-1.875H16.5zM4.125 3.75H7.5v16.5H4.125a1.875 1.875 0 0 1-1.875-1.875V5.625c0-1.036.84-1.875 1.875-1.875"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewColumnsTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 17h2.75A2.25 2.25 0 0 0 19 14.75v-9.5A2.25 2.25 0 0 0 16.75 3H14zM12.5 3h-5v14h5zM3.25 3H6v14H3.25A2.25 2.25 0 0 1 1 14.75v-9.5A2.25 2.25 0 0 1 3.25 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewfinderCircle(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.5 3.75H6A2.25 2.25 0 0 0 3.75 6v1.5M16.5 3.75H18A2.25 2.25 0 0 1 20.25 6v1.5m0 9V18A2.25 2.25 0 0 1 18 20.25h-1.5m-9 0H6A2.25 2.25 0 0 1 3.75 18v-1.5M15 12a3 3 0 1 1-6 0a3 3 0 0 1 6 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewfinderCircleSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.75 2A1.75 1.75 0 0 0 2 3.75v1.5a.75.75 0 0 0 1.5 0v-1.5a.25.25 0 0 1 .25-.25h1.5a.75.75 0 0 0 0-1.5zm7 0a.75.75 0 0 0 0 1.5h1.5a.25.25 0 0 1 .25.25v1.5a.75.75 0 0 0 1.5 0v-1.5A1.75 1.75 0 0 0 12.25 2zM3.5 10.75a.75.75 0 0 0-1.5 0v1.5c0 .966.784 1.75 1.75 1.75h1.5a.75.75 0 0 0 0-1.5h-1.5a.25.25 0 0 1-.25-.25zm10.5 0a.75.75 0 0 0-1.5 0v1.5a.25.25 0 0 1-.25.25h-1.5a.75.75 0 0 0 0 1.5h1.5A1.75 1.75 0 0 0 14 12.25zM8 10a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewfinderCircleSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3a3 3 0 0 0-3 3v1.5a.75.75 0 0 0 1.5 0V6A1.5 1.5 0 0 1 6 4.5h1.5a.75.75 0 0 0 0-1.5zm10.5 0a.75.75 0 0 0 0 1.5H18A1.5 1.5 0 0 1 19.5 6v1.5a.75.75 0 0 0 1.5 0V6a3 3 0 0 0-3-3zM12 8.25a3.75 3.75 0 1 0 0 7.5a3.75 3.75 0 0 0 0-7.5M4.5 16.5a.75.75 0 0 0-1.5 0V18a3 3 0 0 0 3 3h1.5a.75.75 0 0 0 0-1.5H6A1.5 1.5 0 0 1 4.5 18zm16.5 0a.75.75 0 0 0-1.5 0V18a1.5 1.5 0 0 1-1.5 1.5h-1.5a.75.75 0 0 0 0 1.5H18a3 3 0 0 0 3-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewfinderCircleTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.25 2A2.25 2.25 0 0 0 2 4.25v2a.75.75 0 0 0 1.5 0v-2a.75.75 0 0 1 .75-.75h2a.75.75 0 0 0 0-1.5zm9.5 0a.75.75 0 0 0 0 1.5h2a.75.75 0 0 1 .75.75v2a.75.75 0 0 0 1.5 0v-2A2.25 2.25 0 0 0 15.75 2zM3.5 13.75a.75.75 0 0 0-1.5 0v2A2.25 2.25 0 0 0 4.25 18h2a.75.75 0 0 0 0-1.5h-2a.75.75 0 0 1-.75-.75zm14.5 0a.75.75 0 0 0-1.5 0v2a.75.75 0 0 1-.75.75h-2a.75.75 0 0 0 0 1.5h2A2.25 2.25 0 0 0 18 15.75zM7 10a3 3 0 1 1 6 0a3 3 0 0 1-6 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wallet(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 12a2.25 2.25 0 0 0-2.25-2.25H15a3 3 0 1 1-6 0H5.25A2.25 2.25 0 0 0 3 12m18 0v6a2.25 2.25 0 0 1-2.25 2.25H5.25A2.25 2.25 0 0 1 3 18v-6m18 0V9M3 12V9m18 0a2.25 2.25 0 0 0-2.25-2.25H5.25A2.25 2.25 0 0 0 3 9m18 0V6a2.25 2.25 0 0 0-2.25-2.25H5.25A2.25 2.25 0 0 0 3 6v3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WalletSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 3.5A1.5 1.5 0 0 1 3.5 2h9A1.5 1.5 0 0 1 14 3.5v.401a2.986 2.986 0 0 0-1.5-.401h-9c-.546 0-1.059.146-1.5.401zM3.5 5A1.5 1.5 0 0 0 2 6.5v.401A2.986 2.986 0 0 1 3.5 6.5h9c.546 0 1.059.146 1.5.401V6.5A1.5 1.5 0 0 0 12.5 5zM8 10a2 2 0 0 0 1.938-1.505c.068-.268.286-.495.562-.495h2A1.5 1.5 0 0 1 14 9.5v3a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 2 12.5v-3A1.5 1.5 0 0 1 3.5 8h2c.276 0 .494.227.562.495A2 2 0 0 0 8 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WalletSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.273 5.625A4.483 4.483 0 0 1 5.25 4.5h13.5c1.141 0 2.183.425 2.977 1.125A3 3 0 0 0 18.75 3H5.25a3 3 0 0 0-2.977 2.625m0 3A4.483 4.483 0 0 1 5.25 7.5h13.5c1.141 0 2.183.425 2.977 1.125A3 3 0 0 0 18.75 6H5.25a3 3 0 0 0-2.977 2.625M5.25 9a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h13.5a3 3 0 0 0 3-3v-6a3 3 0 0 0-3-3H15a.75.75 0 0 0-.75.75a2.25 2.25 0 0 1-4.5 0A.75.75 0 0 0 9 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WalletTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 4.25a3.733 3.733 0 0 1 2.25-.75h13.5c.844 0 1.623.279 2.25.75A2.25 2.25 0 0 0 16.75 2H3.25A2.25 2.25 0 0 0 1 4.25m0 3a3.733 3.733 0 0 1 2.25-.75h13.5c.844 0 1.623.279 2.25.75A2.25 2.25 0 0 0 16.75 5H3.25A2.25 2.25 0 0 0 1 7.25M7 8a1 1 0 0 1 1 1a2 2 0 1 0 4 0a1 1 0 0 1 1-1h3.75A2.25 2.25 0 0 1 19 10.25v5.5A2.25 2.25 0 0 1 16.75 18H3.25A2.25 2.25 0 0 1 1 15.75v-5.5A2.25 2.25 0 0 1 3.25 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wifi(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.288 15.038a5.25 5.25 0 0 1 7.424 0M5.106 11.856c3.807-3.808 9.98-3.808 13.788 0M1.924 8.674c5.565-5.565 14.587-5.565 20.152 0M12.53 18.22l-.53.53l-.53-.53a.75.75 0 0 1 1.06 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.188 7.063a8.75 8.75 0 0 0-12.374 0a.75.75 0 0 1-1.061-1.06c4.003-4.004 10.493-4.004 14.496 0a.75.75 0 1 1-1.061 1.06m-2.121 2.121a5.75 5.75 0 0 0-8.132 0a.75.75 0 0 1-1.06-1.06a7.25 7.25 0 0 1 10.252 0a.75.75 0 0 1-1.06 1.06m-2.122 2.122a2.75 2.75 0 0 0-3.889 0a.75.75 0 1 1-1.06-1.061a4.25 4.25 0 0 1 6.01 0a.75.75 0 0 1-1.06 1.06m-2.828 1.06a1.25 1.25 0 0 1 1.768 0a.75.75 0 0 1 0 1.06l-.355.355a.75.75 0 0 1-1.06 0l-.354-.354a.75.75 0 0 1 0-1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.371 8.143c5.858-5.857 15.356-5.857 21.214 0a.75.75 0 0 1 0 1.061l-.53.53a.75.75 0 0 1-1.061 0c-4.98-4.979-13.053-4.979-18.032 0a.75.75 0 0 1-1.06 0l-.53-.53a.75.75 0 0 1 0-1.06m3.182 3.182c4.1-4.1 10.749-4.1 14.85 0a.75.75 0 0 1 0 1.061l-.53.53a.75.75 0 0 1-1.062 0a8.25 8.25 0 0 0-11.667 0a.75.75 0 0 1-1.06 0l-.53-.53a.75.75 0 0 1 0-1.06m3.204 3.182a6 6 0 0 1 8.486 0a.75.75 0 0 1 0 1.061l-.53.53a.75.75 0 0 1-1.061 0a3.75 3.75 0 0 0-5.304 0a.75.75 0 0 1-1.06 0l-.53-.53a.75.75 0 0 1 0-1.06m3.182 3.182a1.5 1.5 0 0 1 2.122 0a.75.75 0 0 1 0 1.061l-.53.53a.75.75 0 0 1-1.061 0l-.53-.53a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M.676 6.941A12.964 12.964 0 0 1 10 3c3.657 0 6.963 1.511 9.324 3.941a.75.75 0 0 1-.008 1.053l-.353.354a.75.75 0 0 1-1.069-.008C15.894 6.28 13.097 5 10 5C6.903 5 4.106 6.28 2.106 8.34a.75.75 0 0 1-1.069.008l-.353-.354a.75.75 0 0 1-.008-1.053m2.825 2.833A8.976 8.976 0 0 1 10 7a8.976 8.976 0 0 1 6.499 2.774a.75.75 0 0 1-.011 1.049l-.354.354a.75.75 0 0 1-1.072-.012A6.978 6.978 0 0 0 10 9c-1.99 0-3.786.83-5.061 2.165a.75.75 0 0 1-1.073.012l-.354-.354a.75.75 0 0 1-.01-1.05m2.82 2.84A4.989 4.989 0 0 1 10 11c1.456 0 2.767.623 3.68 1.614a.75.75 0 0 1-.022 1.039l-.354.354a.75.75 0 0 1-1.085-.026A2.99 2.99 0 0 0 10 13c-.88 0-1.67.377-2.22.981a.75.75 0 0 1-1.084.026l-.354-.354a.75.75 0 0 1-.021-1.039m2.795 2.752a1.248 1.248 0 0 1 1.768 0a.75.75 0 0 1 0 1.06l-.354.354a.75.75 0 0 1-1.06 0l-.354-.353a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Window(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 8.25V18a2.25 2.25 0 0 0 2.25 2.25h13.5A2.25 2.25 0 0 0 21 18V8.25m-18 0V6a2.25 2.25 0 0 1 2.25-2.25h13.5A2.25 2.25 0 0 1 21 6v2.25m-18 0h18M5.25 6h.008v.008H5.25zM7.5 6h.008v.008H7.5zm2.25 0h.008v.008H9.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 12V4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2m1.5-5.5V12a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5V6.5A.5.5 0 0 0 12 6H4a.5.5 0 0 0-.5.5m.75-1.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5M7 4a.75.75 0 1 1-1.5 0A.75.75 0 0 1 7 4m1.25.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 6a3 3 0 0 1 3-3h13.5a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H5.25a3 3 0 0 1-3-3zm18 3H3.75v9a1.5 1.5 0 0 0 1.5 1.5h13.5a1.5 1.5 0 0 0 1.5-1.5zm-15-3.75A.75.75 0 0 0 4.5 6v.008c0 .414.336.75.75.75h.008a.75.75 0 0 0 .75-.75V6a.75.75 0 0 0-.75-.75zm1.5.75a.75.75 0 0 1 .75-.75h.008a.75.75 0 0 1 .75.75v.008a.75.75 0 0 1-.75.75H7.5a.75.75 0 0 1-.75-.75zm3-.75A.75.75 0 0 0 9 6v.008c0 .414.336.75.75.75h.008a.75.75 0 0 0 .75-.75V6a.75.75 0 0 0-.75-.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.25 2A2.25 2.25 0 0 0 2 4.25v11.5A2.25 2.25 0 0 0 4.25 18h11.5A2.25 2.25 0 0 0 18 15.75V4.25A2.25 2.25 0 0 0 15.75 2zM3.5 8v7.75c0 .414.336.75.75.75h11.5a.75.75 0 0 0 .75-.75V8zM5 4.25a.75.75 0 0 0-.75.75v.01c0 .414.336.75.75.75h.01a.75.75 0 0 0 .75-.75V5a.75.75 0 0 0-.75-.75zM7.25 5A.75.75 0 0 1 8 4.25h.01a.75.75 0 0 1 .75.75v.01a.75.75 0 0 1-.75.75H8a.75.75 0 0 1-.75-.75zM11 4.25a.75.75 0 0 0-.75.75v.01c0 .414.336.75.75.75h.01a.75.75 0 0 0 .75-.75V5a.75.75 0 0 0-.75-.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wrench(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.75 6.75a4.5 4.5 0 0 1-4.884 4.484c-1.076-.091-2.264.071-2.95.904l-7.152 8.684a2.548 2.548 0 1 1-3.586-3.586l8.684-7.151c.833-.687.995-1.875.904-2.95a4.5 4.5 0 0 1 6.336-4.486l-3.276 3.275a3.004 3.004 0 0 0 2.25 2.25l3.276-3.276c.256.565.398 1.192.398 1.852"/><path d="M4.867 19.125h.008v.008h-.008z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WrenchScrewdriver(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.42 15.17L17.25 21A2.652 2.652 0 0 0 21 17.25l-5.877-5.877M11.42 15.17l2.496-3.03c.317-.384.74-.626 1.208-.766M11.42 15.17l-4.655 5.653a2.548 2.548 0 1 1-3.586-3.586l6.837-5.63m5.108-.233c.55-.164 1.163-.188 1.743-.14a4.5 4.5 0 0 0 4.486-6.336l-3.276 3.277a3.004 3.004 0 0 1-2.25-2.25l3.276-3.276a4.5 4.5 0 0 0-6.336 4.486c.091 1.076-.071 2.264-.904 2.95l-.102.085m-1.745 1.437L5.909 7.5H4.5L2.25 3.75l1.5-1.5L7.5 4.5v1.409l4.26 4.26m-1.745 1.437l1.745-1.437m6.615 8.206L15.75 15.75M4.867 19.125h.008v.008h-.008z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WrenchScrewdriverSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 4.5A3.5 3.5 0 0 1 11.435 8c-.99-.019-2.093.132-2.7.913l-4.13 5.31a2.015 2.015 0 1 1-2.827-2.828l5.309-4.13c.78-.607.932-1.71.914-2.7L8 4.5a3.5 3.5 0 0 1 4.477-3.362c.325.094.39.497.15.736L10.6 3.902a.48.48 0 0 0-.033.653c.271.314.565.608.879.879a.48.48 0 0 0 .653-.033l2.027-2.027c.239-.24.642-.175.736.15c.09.31.138.637.138.976M3.75 13a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0" clip-rule="evenodd"/><path fill="currentColor" d="M11.5 9.5c.313 0 .62-.029.917-.084l1.962 1.962a2.121 2.121 0 0 1-3 3l-2.81-2.81l1.35-1.734c.05-.064.158-.158.426-.233c.278-.078.639-.11 1.062-.102zM5 4l1.446 1.445a2.256 2.256 0 0 1-.047.21c-.075.268-.169.377-.233.427l-.61.474L4 5H2.656a.25.25 0 0 1-.224-.139l-1.35-2.7a.25.25 0 0 1 .047-.289l.745-.745a.25.25 0 0 1 .289-.047l2.7 1.35A.25.25 0 0 1 5 2.654z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WrenchScrewdriverSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M12 6.75a5.25 5.25 0 0 1 6.775-5.025a.75.75 0 0 1 .313 1.248l-3.32 3.319a2.248 2.248 0 0 0 1.941 1.939l3.318-3.319a.75.75 0 0 1 1.248.313a5.25 5.25 0 0 1-5.472 6.756c-1.018-.086-1.87.1-2.309.634L7.344 21.3A3.298 3.298 0 1 1 2.7 16.657l8.684-7.151c.533-.44.72-1.291.634-2.309A5.342 5.342 0 0 1 12 6.75M4.117 19.125a.75.75 0 0 1 .75-.75h.008a.75.75 0 0 1 .75.75v.008a.75.75 0 0 1-.75.75h-.008a.75.75 0 0 1-.75-.75z" clip-rule="evenodd"/><path d="m10.076 8.64l-2.201-2.2V4.874a.75.75 0 0 0-.364-.643l-3.75-2.25a.75.75 0 0 0-.916.113l-.75.75a.75.75 0 0 0-.113.916l2.25 3.75a.75.75 0 0 0 .643.364h1.564l2.062 2.062z"/><path fill-rule="evenodd" d="m12.556 17.329l4.183 4.182a3.375 3.375 0 0 0 4.773-4.773l-3.306-3.305a6.803 6.803 0 0 1-1.53.043c-.394-.034-.682-.006-.867.042a.589.589 0 0 0-.167.063zm3.414-1.36a.75.75 0 0 1 1.06 0l1.875 1.876a.75.75 0 1 1-1.06 1.06L15.97 17.03a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WrenchScrewdriverTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M14.5 10a4.5 4.5 0 0 0 4.284-5.882c-.105-.324-.51-.391-.751-.15L15.34 6.66a.454.454 0 0 1-.493.11a3.01 3.01 0 0 1-1.618-1.616a.455.455 0 0 1 .11-.494l2.694-2.692c.24-.241.174-.647-.15-.752a4.5 4.5 0 0 0-5.873 4.575c.055.873-.128 1.808-.8 2.368l-7.23 6.024a2.724 2.724 0 1 0 3.837 3.837l6.024-7.23c.56-.672 1.495-.855 2.368-.8c.096.007.193.01.291.01M5 16a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/><path d="M14.5 11.5c.173 0 .345-.007.514-.022l3.754 3.754a2.5 2.5 0 0 1-3.536 3.536l-4.41-4.41l2.172-2.607c.052-.062.147-.138.342-.196a2.28 2.28 0 0 1 .777-.067c.128.008.257.012.387.012M6 4.586l2.33 2.33a.452.452 0 0 1-.08.09L6.8 8.214L4.586 6H3.309a.5.5 0 0 1-.447-.276l-1.7-3.402a.5.5 0 0 1 .093-.577l.49-.49a.5.5 0 0 1 .577-.094l3.402 1.7A.5.5 0 0 1 6 3.31z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WrenchSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 8a3.5 3.5 0 0 0 3.362-4.476c-.094-.325-.497-.39-.736-.15L12.099 5.4a.48.48 0 0 1-.653.033a8.554 8.554 0 0 1-.879-.879a.48.48 0 0 1 .033-.653l2.027-2.028c.24-.239.175-.642-.15-.736a3.502 3.502 0 0 0-4.476 3.427c.018.99-.133 2.093-.914 2.7l-5.31 4.13a2.015 2.015 0 1 0 2.828 2.827l4.13-5.309c.607-.78 1.71-.932 2.7-.914zM3 13.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WrenchSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 6.75a5.25 5.25 0 0 1 6.775-5.025a.75.75 0 0 1 .313 1.248l-3.32 3.319a2.248 2.248 0 0 0 1.941 1.939l3.318-3.319a.75.75 0 0 1 1.248.313a5.25 5.25 0 0 1-5.472 6.756c-1.018-.086-1.87.1-2.309.634L7.344 21.3A3.298 3.298 0 1 1 2.7 16.657l8.684-7.151c.533-.44.72-1.291.634-2.309A5.342 5.342 0 0 1 12 6.75M4.117 19.125a.75.75 0 0 1 .75-.75h.008a.75.75 0 0 1 .75.75v.008a.75.75 0 0 1-.75.75h-.008a.75.75 0 0 1-.75-.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WrenchTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M19 5.5a4.5 4.5 0 0 1-4.791 4.49c-.873-.055-1.808.128-2.368.8l-6.024 7.23a2.724 2.724 0 1 1-3.837-3.837L9.21 8.16c.672-.56.855-1.495.8-2.368a4.5 4.5 0 0 1 5.873-4.575c.324.105.39.51.15.752L13.34 4.66a.455.455 0 0 0-.11.494a3.01 3.01 0 0 0 1.617 1.617c.17.07.363.02.493-.111l2.692-2.692c.241-.241.647-.174.752.15c.14.435.216.9.216 1.382M4 17a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xcircle(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9.75 9.75l4.5 4.5m0-4.5l-4.5 4.5M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func XcircleSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14m2.78-4.22a.75.75 0 0 1-1.06 0L8 9.06l-1.72 1.72a.75.75 0 1 1-1.06-1.06L6.94 8L5.22 6.28a.75.75 0 0 1 1.06-1.06L8 6.94l1.72-1.72a.75.75 0 1 1 1.06 1.06L9.06 8l1.72 1.72a.75.75 0 0 1 0 1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func XcircleSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75s9.75-4.365 9.75-9.75S17.385 2.25 12 2.25m-1.72 6.97a.75.75 0 1 0-1.06 1.06L10.94 12l-1.72 1.72a.75.75 0 1 0 1.06 1.06L12 13.06l1.72 1.72a.75.75 0 1 0 1.06-1.06L13.06 12l1.72-1.72a.75.75 0 1 0-1.06-1.06L12 10.94z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func XcircleTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 18a8 8 0 1 0 0-16a8 8 0 0 0 0 16M8.28 7.22a.75.75 0 0 0-1.06 1.06L8.94 10l-1.72 1.72a.75.75 0 1 0 1.06 1.06L10 11.06l1.72 1.72a.75.75 0 1 0 1.06-1.06L11.06 10l1.72-1.72a.75.75 0 0 0-1.06-1.06L10 8.94z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xmark(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 18L18 6M6 6l12 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func XmarkSixteenSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.28 4.22a.75.75 0 0 0-1.06 1.06L6.94 8l-2.72 2.72a.75.75 0 1 0 1.06 1.06L8 9.06l2.72 2.72a.75.75 0 1 0 1.06-1.06L9.06 8l2.72-2.72a.75.75 0 0 0-1.06-1.06L8 6.94z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func XmarkSolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.47 5.47a.75.75 0 0 1 1.06 0L12 10.94l5.47-5.47a.75.75 0 1 1 1.06 1.06L13.06 12l5.47 5.47a.75.75 0 1 1-1.06 1.06L12 13.06l-5.47 5.47a.75.75 0 0 1-1.06-1.06L10.94 12L5.47 6.53a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func XmarkTwentySolid(children ...ElementRenderer) *HeroiconsIcon {
	return &HeroiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.28 5.22a.75.75 0 0 0-1.06 1.06L8.94 10l-3.72 3.72a.75.75 0 1 0 1.06 1.06L10 11.06l3.72 3.72a.75.75 0 1 0 1.06-1.06L11.06 10l3.72-3.72a.75.75 0 0 0-1.06-1.06L10 8.94z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
