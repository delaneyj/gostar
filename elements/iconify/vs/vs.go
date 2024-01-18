package vs

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "5.11.2"
	hAttr          = "1em"
	viewbox        = "0 0 1792 1792"
	fill           = "currentColor"
)

type VsIcon struct {
	*SVGSVGElement
}

type VsIconFn func(children ...ElementRenderer) *VsIcon

var IconLookup = map[string]VsIconFn{
	"asquare":          Asquare,
	"bsquare":          Bsquare,
	"baby":             Baby,
	"bcCard":           BcCard,
	"butterfly":        Butterfly,
	"csquare":          Csquare,
	"calendar":         Calendar,
	"calendarAlt":      CalendarAlt,
	"calendarAltTwo":   CalendarAltTwo,
	"camel":            Camel,
	"catFace":          CatFace,
	"ccCard":           CcCard,
	"chair":            Chair,
	"chairAlt":         ChairAlt,
	"chicken":          Chicken,
	"clipNote":         ClipNote,
	"clipNoteO":        ClipNoteO,
	"clipboard":        Clipboard,
	"clock":            Clock,
	"clockAlt":         ClockAlt,
	"comment":          Comment,
	"commentBubble":    CommentBubble,
	"comments":         Comments,
	"cow":              Cow,
	"crown":            Crown,
	"cutlery":          Cutlery,
	"dsquare":          Dsquare,
	"doorClosed":       DoorClosed,
	"doorOpen":         DoorOpen,
	"drumstick":        Drumstick,
	"esquare":          Esquare,
	"editPage":         EditPage,
	"eightSquare":      EightSquare,
	"fsquare":          Fsquare,
	"faceAllergy":      FaceAllergy,
	"faceDislike":      FaceDislike,
	"faceLike":         FaceLike,
	"fileDownload":     FileDownload,
	"fileDownloadO":    FileDownloadO,
	"fileMoveO":        FileMoveO,
	"fish":             Fish,
	"fiveSquare":       FiveSquare,
	"floors":           Floors,
	"flower":           Flower,
	"fourSquare":       FourSquare,
	"gsquare":          Gsquare,
	"gantt":            Gantt,
	"ganttO":           GanttO,
	"globe":            Globe,
	"grapes":           Grapes,
	"hsquare":          Hsquare,
	"highchair":        Highchair,
	"hourglass":        Hourglass,
	"isquare":          Isquare,
	"idBadgeAlt":       IdBadgeAlt,
	"idCard":           IdCard,
	"idCardAlt":        IdCardAlt,
	"jsquare":          Jsquare,
	"ksquare":          Ksquare,
	"kakao":            Kakao,
	"kakaoSquare":      KakaoSquare,
	"kakaotalk":        Kakaotalk,
	"kakaotalkSquare":  KakaotalkSquare,
	"kanjiChu":         KanjiChu,
	"kanjiUtage":       KanjiUtage,
	"kanjiYubi":        KanjiYubi,
	"keyboard":         Keyboard,
	"lsquare":          Lsquare,
	"language":         Language,
	"line":             Line,
	"lineSquare":       LineSquare,
	"msquare":          Msquare,
	"magnetNote":       MagnetNote,
	"maleFemale":       MaleFemale,
	"mic":              Mic,
	"mobile":           Mobile,
	"moon":             Moon,
	"multiArrow":       MultiArrow,
	"nsquare":          Nsquare,
	"naver":            Naver,
	"naverSquare":      NaverSquare,
	"neko":             Neko,
	"nekoSleep":        NekoSleep,
	"nineSquare":       NineSquare,
	"ninja":            Ninja,
	"noCommentBubble":  NoCommentBubble,
	"noSmokingAlt":     NoSmokingAlt,
	"osquare":          Osquare,
	"oneSquare":        OneSquare,
	"picon":            Picon,
	"psquare":          Psquare,
	"panther":          Panther,
	"party":            Party,
	"peopleGroup":      PeopleGroup,
	"person":           Person,
	"pig":              Pig,
	"pregnant":         Pregnant,
	"profile":          Profile,
	"qsquare":          Qsquare,
	"questionSquare":   QuestionSquare,
	"rsquare":          Rsquare,
	"rose":             Rose,
	"ssquare":          Ssquare,
	"senior":           Senior,
	"sevenSquare":      SevenSquare,
	"sexFemale":        SexFemale,
	"sexMale":          SexMale,
	"sheep":            Sheep,
	"shieldCheck":      ShieldCheck,
	"shieldTimes":      ShieldTimes,
	"shop":             Shop,
	"sixSquare":        SixSquare,
	"sleep":            Sleep,
	"sleepSquare":      SleepSquare,
	"smoking":          Smoking,
	"smokingAlt":       SmokingAlt,
	"sms":              Sms,
	"sofa":             Sofa,
	"speech":           Speech,
	"spinner":          Spinner,
	"stickyNote":       StickyNote,
	"stroller":         Stroller,
	"sun":              Sun,
	"sunrise":          Sunrise,
	"sunriseO":         SunriseO,
	"tsquare":          Tsquare,
	"table":            Table,
	"tableAlt":         TableAlt,
	"tableO":           TableO,
	"tableQuestion":    TableQuestion,
	"tables":           Tables,
	"tablesolution":    Tablesolution,
	"threeSquare":      ThreeSquare,
	"timeslot":         Timeslot,
	"timeslotQuestion": TimeslotQuestion,
	"timeslots":        Timeslots,
	"twoSquare":        TwoSquare,
	"usquare":          Usquare,
	"userBoss":         UserBoss,
	"userGroup":        UserGroup,
	"userSuit":         UserSuit,
	"userSuitFemale":   UserSuitFemale,
	"userWaiter":       UserWaiter,
	"userWaiterFemale": UserWaiterFemale,
	"vsquare":          Vsquare,
	"volumeOn":         VolumeOn,
	"volumeTimes":      VolumeTimes,
	"wsquare":          Wsquare,
	"walk":             Walk,
	"weddingCake":      WeddingCake,
	"whiteboard":       Whiteboard,
	"window":           Window,
	"wine":             Wine,
	"wineO":            WineO,
	"xsquare":          Xsquare,
	"ysquare":          Ysquare,
	"yahooJapan":       YahooJapan,
	"zsquare":          Zsquare,
	"zeroSquare":       ZeroSquare,
}

func Asquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0m331 1188h458l85 251q11 36 24 45t44 9h157q53 0 35-69l-333-964q-17-64-57.5-112.5T982 299H810q-57 0-97.5 48.5T655 460l-333 964q-18 69 35 69h157q31 0 44-9t24-45zm391-225H734l130-385q8-26 32-26t32 26z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bsquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0m339 507v285h334q49 0 84.5-41.5T1129 650t-35-101t-85-42zM428 374v1044q0 31 21.5 53t52.5 22h507q147 0 251-103t104-248q0-144-102-246q102-102 102-246q0-145-104-248t-251-103H502q-31 0-52.5 22T428 374m581 911H675v-285h334q49 0 84.5 41.5T1129 1142t-35 101t-85 42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Baby(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M802 224q0 93-66 159t-160 66q-92 0-157.5-66T353 224q0-92 65.5-158T576 0q94 0 160 66t66 158m-7 1175l-142 142q-24 24-31 50.5t-.5 47.5t22.5 39q21 24 61.5 22.5T778 1667l268-268q32-32 40-71.5t-3.5-73t-36.5-59.5l-176-174l-227 227zm48-442H310V794L148 956q-34 30-68 28t-58-27Q1 935 .5 906.5T27 848l223-223l115-116v3q24-20 70-20h285q45 0 68 20v-3l116 116l224 223q26 30 25.5 58.5T1133 957q-24 25-58.5 27t-69.5-28L843 794zm-482 442l142 142q24 24 31 50.5t.5 47.5t-22.5 39q-21 24-61.5 22.5T378 1667l-268-268q-32-32-40-71.5t3.5-73T110 1195l176-174l229 227z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BcCard(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2368 192v1280q0 78-57 135t-135 57H192q-77 0-134.5-57T0 1472V192q0-77 57.5-134.5T192 0h1984q78 0 135 57.5t57 134.5M470 517v233h272q41 0 69.5-34t28.5-82t-29-82.5t-69-34.5zM268 409v851q0 25 18 43.5t43 18.5h413q120 0 205-84t85-203q0-116-84-200q84-84 84-201q0-118-85-202t-205-84H329q-25 0-43 18t-18 43m474 743H470V919h272q41 0 69.5 34t28.5 82t-29 82.5t-69 34.5m445-317q0 99 38.5 189t104 155.5t155.5 104t189 38.5q157 0 278-80.5t173-218.5q25-66-19-66h-152q-16 0-26 3.5t-13.5 6.5t-12 16t-14.5 20q-38 49-94.5 77.5T1672 1109q-114 0-194.5-80T1397 835t80.5-194.5T1672 560q64 0 120.5 29t95.5 78q6 7 12.5 17t9.5 13.5t9 8t14 5.5t21 1h152q44 0 19-66q-52-138-173-218t-278-80q-99 0-189 38.5t-155.5 104t-104 155.5t-38.5 189"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Butterfly(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M793 768H256q-51 0-94.5 9.5t-81 30t-59 59.5T0 960q0 69 42.5 148T153 1252.5t161 109t185 42.5q152-2 270-120q-45 58-55 130.5t11 137t57 120t83 88t90 32.5q64 0 111.5-21.5t75-60t40.5-86t13-103.5q0-83-26.5-170t-69.5-148zm48-38l144-532q22-83 73.5-140.5T1188 0q81 0 165 73.5T1492 265t55 235q0 109-51 197t-146 142q83-33 162.5-13.5t141.5 68t100 111.5t38 118q0 30-13 64t-38.5 65.5t-70 52.5t-99.5 21q-101 0-211.5-53.5T1185 1143z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Csquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0m-37 896q0 162 80 299.5T596.5 1413t299.5 80q192 0 340.5-98.5T1449 1127q30-81-23-81h-186q-20 0-32 4t-16 7.5t-15 19.5t-18 25q-47 60-116 95t-149 35q-140 0-238.5-98.5T557 896t98.5-237.5T894 560q80 0 149 35t116 95q6 8 13 18l10.5 15l8.5 10.5l11 7.5l15.5 3.5l22.5 1.5h186q53 0 23-81q-64-169-212.5-267.5T896 299q-162 0-299.5 80T379 596.5T299 896"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M56 333h1553q55 0 55-56V111q0-47-32-79t-79-32H111Q64 0 32 32T0 111v166q0 56 56 56m221 998H56q-56 0-56 56v166q0 47 32 79t79 32h166q56 0 56-56v-221q0-56-56-56m1110 0q-56 0-56 56v221q0 56 56 56h166q47 0 79-32t32-79v-166q0-56-55-56zm-444 0q-56 0-56 56v221q0 56 56 56h222q55 0 55-56v-221q0-56-55-56zm-444 0q-55 0-55 56v221q0 56 55 56h222q56 0 56-56v-221q0-56-56-56zm888-444q-56 0-56 56v222q0 55 56 55h222q55 0 55-55V943q0-56-55-56zm-444 0q-56 0-56 56v222q0 55 56 55h222q55 0 55-55V943q0-56-55-56zm-887 0q-56 0-56 56v222q0 55 56 55h221q56 0 56-55V943q0-56-56-56zm443 0q-55 0-55 56v222q0 55 55 55h222q56 0 56-55V943q0-56-56-56zm888-443q-56 0-56 55v222q0 56 56 56h222q55 0 55-56V499q0-55-55-55zm-444 0q-56 0-56 55v222q0 56 56 56h222q55 0 55-56V499q0-55-55-55zm-887 0q-56 0-56 55v222q0 56 56 56h221q56 0 56-56V499q0-55-56-55zm443 0q-55 0-55 55v222q0 56 55 56h222q56 0 56-56V499q0-55-56-55z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarAlt(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M561 561v224h224V561zm320 0v224h224V561zm320 0v224h224V561zm0 320v224h224V881zm-320 0v224h224V881zm0 320v224h224v-224zm-320 0v224h224v-224zm0-320v224h224V881zm-320 320v224h224v-224zm0-320v224h224V881zm271-634V41q0-16-9.5-28.5T480 0h-64q-13 0-22.5 12.5T384 41v206q0 17 9.5 28.5T416 288h64q14-1 23-12.5t9-28.5M128 1664h1408q52 0 89-36t39-88V252q0-51-38-87.5t-90-36.5h-192v96q0 63-21 95.5t-75 32.5h-64q-96-1-96-127v-97H576v97q0 126-96 127h-64q-54 0-75-32.5T320 224v-96H128q-52 0-90 36.5T0 252v1288q0 51 37.5 87t90.5 37m0-1216h1408v1088H128zm1152-201V41q0-16-9.5-28.5T1248 0h-64q-13 0-22.5 12.5T1152 41v206q0 17 9.5 28.5T1184 288h64q14-1 23-12.5t9-28.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarAltTwo(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1792V644h1792v1148zm1186-780v553h178V840h-178l-111 67v171zm-374 191q49-13 74-56q27-48 27-99q0-67-35-114q-35-48-89-71q-52-23-117-23q-116 0-186 71q-16 17-43.5 69.5T416 1051h165q0-6 1-11.5t1.5-8t3.5-7t4-5.5l7-7l7-7q24-25 62-25q37 0 62 23q24 21 24 61q0 32-25 56q-25 25-60 25h-61v116h61q33 0 60 25q25 24 25 57q0 39-24 60q-23 23-62 23q-38 0-62-25l-5.5-5.5l-5.5-5l-4-4.5l-3-4.5l-2.5-4.5l-2-6l-1-6.5l-.5-8.5H416q-1 18 26.5 70t43.5 69q70 72 186 72q63 0 117-24q54-23 89-71q35-47 35-114q0-51-27-99q-25-42-74-56M0 567V224q0-26 19.5-44.5T66 161h243v199q0 12 9.5 21t22.5 9h280q13 0 22.5-9t9.5-21V161h490v199q0 12 9.5 21t22.5 9h280q13 0 22.5-9t9.5-21V161h239q27 0 46.5 18.5T1792 224v343zM1378 0q13 0 22 8.5t9 20.5v258q0 11-9 19.5t-22 8.5h-126q-12 0-21-8.5t-9-19.5V29q0-12 9-20.5t21-8.5zM418 0q-13 0-22 8.5T387 29v258q0 11 9 19.5t22 8.5h126q12 0 21-8.5t9-19.5V29q0-12-9-20.5T544 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camel(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M123 1539q41 0 70.5 30t29.5 71H0l42-526l124-177l-6-264l-50 160q-3 9-3.5 42t-6 57T81 956H30q-8 0-12-9t0-22l125-388q21-67 50.5-116t57.5-71t53-35t40-15l15-1q19 0 44-27.5t56-67t69.5-79t97.5-67T754 31h15.5l39.5 4l58.5 12.5l68 28l72.5 47l67 73.5l57 103q20 85 77 158.5t128 73.5q7 0 13.5-.5t26-7.5t35-19.5t35-41t33-67t23-101.5t9.5-142V87q0-58 48-58q21 0 32 18.5t12 36.5v18q13 0 29-.5t22.5-.5t15.5 1t19 5q27 9 58 43l159 8q38 0 60 45t25 89l4 45q0 29-19.5 43t-39.5 13q-25 0-53-34l-176-20q2 112-17.5 206.5T1626 722t-126 136t-193 70l74 319l97 291h34q41 0 70 30t29 71h-200l-220-458l-84-161l-116 195l-116 324h25q41 0 70.5 30t29.5 71H770l56-484l101-224H581l-94 310l50 297h32q41 0 70.5 30t29.5 71H452l-156-467l-21-105l-94 125l-73 346z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CatFace(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M774 1404q-1-2-11 0t-27-4t-33.5-14t-30-30t-16.5-53q-3-23-26-22q-24 1-33 25q-2 8-2 17q6 51 43.5 88.5t75 51.5t66.5 14q25 0 55-27t47-54l17-27q6 1 20 17t29.5 34t43 32.5t58.5 13.5q45-2 100.5-45t64.5-108q3-22-10-31.5t-28.5-3t-17.5 23.5q-9 61-46 82t-82 16q-20-1-39.5-24t-33.5-57.5t-23.5-67.5t-14.5-61q7-4 37-46t30-50q-7-6-48-5.5t-89 3.5l-48 3q74 90 83 95q-4 56-39 133.5t-72 80.5m840-669q-8-6-42.5-39.5t-68.5-56t-71-26.5q-41-5-90 9t-64 39q48 25 69.5 66.5t18 83.5t-18 85t-32.5 72.5t-32 44.5q-22-4-45-22t-46.5-56t-29-94t8.5-137q-108 65-152 317q49 13 110 22t131.5 7.5T1369 1029q83-44 149-129t96-165m-1436 0q30 80 96 165t149 129q38 20 108.5 21.5T663 1043t110-22q-44-252-152-317q8 89 2.5 146T595 943t-43.5 51t-42.5 19q-14-15-32-44.5T444.5 896t-18-85t18-83.5T514 661q-15-25-64-39t-90-9q-37 4-71 26.5t-68.5 56T178 735m718-378q107 0 165 7.5t127 34.5q4 2 20.5-19.5t45-58T1317 244t83.5-86t96.5-76q72-47 123-66.5T1728 0q30 49-.5 276.5T1664 576q69 125 98.5 260.5T1792 1074q0 87-101.5 194.5t-243 193t-296 144T896 1664t-255.5-58.5t-296-144t-243-193T0 1074q0-102 29.5-237.5T128 576q-33-72-63.5-299.5T64 0q58-4 108.5 15.5T294 82q48 31 97 76t84 86t63.5 77.5t45 58T604 399q69-27 127-34.5t165-7.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CcCard(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2368 192v1280q0 78-57 135t-135 57H192q-77 0-134.5-57T0 1472V192q0-77 57.5-134.5T192 0h1984q78 0 135 57.5t57 134.5M1219 835q0 128 63 237t172 172t237 63q152 0 269.5-78t167.5-212q26-64-18-64h-147q-16 0-25.5 3.5t-13 6.5t-11.5 15.5t-14 19.5q-37 47-92 75t-118 28q-110 0-188-78t-78-188t78-188t188-78q63 0 118 28t92 75q5 7 11.5 16.5t9.5 13t9 8t13.5 5.5t20.5 1h147q44 0 18-64q-50-134-167.5-211.5T1691 363q-128 0-237 63t-172 172t-63 237m-1019 1q0 128 63 237t172 172t237 63q152 0 269.5-78t167.5-212q26-64-18-64H944q-16 0-25.5 3.5t-13 6.5t-11.5 15.5t-14 19.5q-37 47-92 75t-118 28q-110 0-188-78t-78-188t78-188t188-78q63 0 118 28t92 75q5 7 11.5 16.5t9.5 13t9 8t13.5 5.5t20.5 1h147q44 0 18-64q-50-134-167.5-211.5T672 364q-128 0-237 63T263 599t-63 237"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chair(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 896h768q53 0 90.5 37t37.5 91t-38 91t-90 37h-48l20 608q1 13-8.5 22.5t-22.5 9.5h-69q-12 0-21-9.5t-10-22.5l-18-608H448l-102 609q-5 31-31 31h-74q-13 0-22.5-9.5T210 1761l95-609q-4-2-44-5.5t-66.5-26.5t-42.5-80q-16-59-26.5-150t-20-221.5T92 497q-5-56-48.5-176.5T0 128q0-46 31.5-87T128 0q59 0 88 37.5T269 154q91 298 115 742"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChairAlt(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M429 1056h676l72-748q14-147-41-227.5T960 0H576Q455 0 399 80.5T357 308zm595 288H512v252q0 28-18.5 48t-45.5 20H320q-27 0-45.5-18t-18.5-46l-1-271q-35-15-61.5-38.5t-43-57.5t-27-67.5t-15.5-83t-6.5-89.5t-1.5-100q-44-9-72-42.5T0 768q0-55 38.5-91.5T158 640q35 0 59.5 10t40.5 31.5t24 38t18 46.5l39 331q8 41 16 47q10 8 44 8h736q37 0 44.5-8t15.5-47l39-331q10-30 18-46.5t24-38t40.5-31.5t59.5-10q81 0 120.5 36.5T1536 768q0 48-28.5 82t-73.5 43q0 61-1.5 100.5t-6.5 90t-15 84.5t-26.5 68.5t-43 58t-61.5 38.5v267q0 28-18.5 46t-45.5 18h-128q-27 0-45.5-18t-18.5-46z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chicken(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M159 1189q2-60 16.5-112.5T210 985t39.5-77t34-85.5T300 723q3-86-55-112q-38-23-127.5-56.5T0 507q26-23 66.5-60t70.5-63.5t56-46.5q-62-41-62-117q0-55 43-96.5T267 82q43 0 79 34.5t57 76.5q-7-47-1-87.5t36.5-73T524 0q56 0 96 40.5t40 98.5q0 47-29 86q8 1 15 3t12 3.5t11 5.5t9.5 5.5t9 8t7.5 8t8.5 10.5t8 11.5t9 14T730 309q30 44 56.5 132t45 170.5T857 710q57 49 123.5 82.5T1103 827q63 1 135.5-36.5t136-94.5t123-127t102-133t67.5-114q15 42 24 75.5t14 88t-3.5 98.5t-38.5 93.5t-83 86.5q2 2 27 3.5t52.5 6T1718 804t46 71q-53 0-94 32.5t-52 66.5q35-8 95 25.5t79 89.5q-45 4-71 31t-52 96q-34 92-107.5 151t-162.5 78t-189 13t-188.5-36t-159.5-77.5T758 1241q70 162 203 252.5t280 91.5q108 1 217-68q-31 71-98.5 126.5t-147 85.5t-151.5 45t-131 17q-371 11-575-151t-196-451m224-902q-41 0-69.5 28.5T285 385t28.5 69.5T383 483t69.5-28.5T481 385t-28.5-69.5T383 287"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipNote(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1664 256v640q0 116-.5 170.5T1662 1210t-4.5 132t-9.5 108.5t-15.5 102.5t-22.5 84t-31.5 83t-42.5 72H0q79-242 103.5-447.5T128 768V256l256 2v574q0 86 53 139t139 53q139 0 180-93q12-28 12-99V384H640v448q0 39-12 51.5T576 896q-33-3-49-17.5T512 832V256zM527 159q30-50 82-50q60 0 90 24.5t35 82.5h128q0-56-13-89q-20-49-80.5-88T617 0Q517 0 452 55q-68 59-68 147v54h128v-24l1-27l4.5-21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipNoteO(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1664 256v637q0 110-.5 153t-1.5 130t-3 122.5t-6.5 102t-11 97.5t-16 80.5t-22 80t-29.5 66t-38 67.5H0q79-242 103.5-447.5T128 768V256h256v-54q0-88 68-147Q517 0 617 0q91 0 151.5 39t80.5 88q13 33 13 89H734q-5-58-35-82.5T609 109q-52 0-82 50q-6 11-9.5 25t-4.5 21t-1 27v24zM512 384v448q-1 32 15 46.5t49 17.5q40 0 52-12.5t12-51.5V512h128v320q0 71-12 99q-41 93-180 93q-86 0-139-53t-53-139V384H256v384q0 326-17.5 515T167 1664h1275q28-57 46-122.5t27.5-118.5t14-155t5.5-163.5t1-211.5V384z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clipboard(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M275 662q-8 0-13.5 6t-5.5 15v44q0 8 5.5 14t13.5 6h602q8 0 13.5-6t5.5-14v-44q0-9-5.5-15t-13.5-6zM768 0q-67-2-115.5 45.5T604 160v125q0 22-5.5 36T581 343t-23 11t-31 6q-49 6-82.5 41.5T411 512h714q0-75-33.5-110.5T1009 360q-20-3-31-6t-23-11t-17.5-22t-5.5-36V160q0-67-48.5-114.5T768 0m0 128q24 2 41 20t17 45q0 25-17 43t-41 20q-24-2-41-20t-17-43q0-27 17-45t41-20m674 0q38 0 66 28t28 66v1476q0 38-28 66t-66 28H94q-38 0-66-28t-28-66V222q0-38 28-66t66-28h418v65q1 27-17 45t-46 18H213q-35 0-60 25t-25 60v1238q0 35 25 60t60 25h1110q35 0 60-25t25-60V341q0-35-25-60t-60-25h-234q-29 0-47.5-18t-17.5-45v-65zM287 916q-12 0-21.5 6.5T256 937v44q0 9 9.5 15t21.5 6h962q12 0 21.5-6t9.5-15v-44q0-8-9.5-14.5T1249 916zm-7 258q-9 0-16.5 6t-7.5 15v44q0 8 7.5 14t16.5 6h720q9 0 16.5-6t7.5-14v-44q0-9-7.5-15t-16.5-6zm3 254q-10 0-18.5 6.5T256 1449v44q0 9 8.5 15t18.5 6h842q10 0 18.5-6t8.5-15v-44q0-8-8.5-14.5t-18.5-6.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1088 768H904q-29-32-72-32h-5L475 384q-19-19-45.5-19T384 384t-19 45.5t19 45.5l352 352v5q0 40 28 68t68 28q43 0 72-32h184q26 0 45-19t19-45t-19-45t-45-19M832 256q26 0 45 19t19 45t-19 45t-45 19t-45-19t-19-45t19-45t45-19m0 1024q26 0 45 19t19 45t-19 45t-45 19t-45-19t-19-45t19-45t45-19M320 768q26 0 45 19t19 45t-19 45t-45 19t-45-19t-19-45t19-45t45-19m1024 0q26 0 45 19t19 45t-19 45t-45 19t-45-19t-19-45t19-45t45-19M832 0Q663 0 508.5 66T243 243T66 508.5T0 832t66 323.5T243 1421t265.5 177t323.5 66t323.5-66t265.5-177t177-265.5t66-323.5t-66-323.5T1421 243T1155.5 66T832 0m0 128q143 0 273.5 55.5t225 150t150 225T1536 832t-55.5 273.5t-150 225t-225 150T832 1536t-273.5-55.5t-225-150t-150-225T128 832t55.5-273.5t150-225t225-150T832 128"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockAlt(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 416q26 0 45 19t19 45v192l266 207q22 20 22 49q0 26-19 45t-45 19t-44-18L733 758q-29-19-29-54V480q0-26 19-45t45-19m0-416Q559 0 382.5 103T103 382.5T0 768t103 385.5T382.5 1433T768 1536t385.5-103t279.5-279.5T1536 768t-103-385.5T1153.5 103T768 0m0 224q148 0 273 73t198 198t73 273t-73 273t-198 198t-273 73t-273.5-73t-198-198T224 768t72.5-273t198-198T768 224"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comment(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1408 64q98 0 177 77t79 179v512q0 109-72.5 182.5T1408 1088H896l-512 512v-512H256q-113 0-184.5-72T0 832V320q0-109 73.5-182.5T256 64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentBubble(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 704q0 174-120 321.5t-326 233t-450 85.5q-70 0-145-8q-198 175-460 242q-49 14-114 22q-17 2-30.5-9t-17.5-29v-1q-3-4-.5-12t2-10t4.5-9.5l6-9l7-8.5l8-9q7-8 31-34.5t34.5-38t31-39.5t32.5-51t27-59t26-76q-157-89-247.5-220T0 704q0-130 71-248.5T262 251t286-136.5T896 64q244 0 450 85.5t326 233T1792 704"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comments(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1664 64q49 0 88.5 39.5T1792 192v512q0 52-38 90t-90 38h-128v384l-388-384H768q-50 0-89-39t-39-89V192q0-50 39-89t89-39zM512 704q1 97 80 176q76 76 176 80h384v128q0 49-37.5 88.5T1028 1216H644l-388 384v-384H128q-50 0-89-39t-39-89V576q0-50 39-89t89-39h384z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cow(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 1152q59-25 145.5-44.5T1088 1088q166 0 320 64q-12-81 11-155.5t77-172.5h-157q-34 0-58.5-25.5T1256 738t24.5-60.5T1339 652h209q7-102 1.5-138.5T1522 455q-32-31-109-51t-133-20H896q-56 0-132 20t-108 51q-22 22-27.5 58.5T630 652h208q34 0 59 25.5t25 60.5t-25 60.5t-59 25.5H682q53 96 75.5 171.5T768 1152m0-960h640q61 0 128 64q50 0 90-6.5t71.5-20.5t53.5-28t46-41t38.5-45.5t41-55.5T1920 0q22 60 24.5 113t-11.5 91.5t-40 71t-57 53.5t-67.5 36.5t-68.5 23t-61 10.5q7 27 15.5 88t9.5 89q8-8 131.5-67t132.5-59q11 0 64 41t113 91t71 58q-41 28-69 59.5t-40 53t-27.5 40.5t-43.5 29t-76 10q-38 0-74-12.5t-72.5-39t-56-43.5t-54.5-51q-7 77-28.5 141T1592 927.5t-38 95t-18 129.5q-1 44 31.5 135t32.5 121q0 60-37.5 107.5t-93 74.5t-128 44.5t-134 23.5t-119.5 6t-119.5-6t-134-23.5t-128-44.5t-93-74.5T576 1408q0-30 32.5-121t31.5-135q-1-71-18-129.5t-38-95T541.5 827T513 686q-35 34-54.5 51t-56 43.5t-72.5 39t-74 12.5q-48 0-76-10t-43.5-29t-27.5-40.5t-40-53T0 640q10-7 71-57.5T185 491t64-41q9 0 132 59t131 67q1-28 10-89t16-88q-29-3-61-10.5t-68.5-23t-68-36.5t-57.5-53.5t-39.5-71T232 113T256 0q17 22 43.5 59t41 55.5T379 160t46 41t53.5 28t71.5 20.5t90 6.5q67-64 128-64m552 1063q-38 0-62.5 21.5T1233 1334v79q0 36 24.5 57.5t62.5 21.5q37 0 62-21.5t25-57.5v-79q0-36-25-57.5t-62-21.5m-464 0q-37 0-62 21.5t-25 57.5v79q0 36 25 57.5t62 21.5q38 0 62.5-21.5T943 1413v-79q0-36-24.5-57.5T856 1255"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crown(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1472 1216H192q-24 0-44 38.5t-20 89.5t20 89.5t44 38.5h1280q24 0 44-38.5t20-89.5t-20-89.5t-44-38.5M128 128q-53 0-90.5 37.5T0 256t37.5 90.5T128 384l80 703h1248l80-703q53 0 90.5-37.5T1664 256t-37.5-90.5T1536 128t-90.5 37.5T1408 256q0 56 41 94q-153 183-207 236.5t-90 53.5q-34 0-83-66.5T899 301q28-17 44.5-46t16.5-63q0-53-37.5-90.5T832 64t-90.5 37.5T704 192q0 34 16.5 63t44.5 46Q644 507 595 573.5T512 640q-36 0-90-53.5T215 350q41-38 41-94q0-53-37.5-90.5T128 128"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cutlery(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1280 64v1600q0 36-18 64q-16 30-46 46q-28 18-64 18t-64-18q-29-17-46-46q-18-28-18-64v-512H800q-13 0-22-10q-19-19 4-145q4-25 15.5-101.5t21.5-143T847 591t37-171.5t47.5-158t59.5-136t73-90T1152 0h64q26 0 45 19t19 45m-640 0v640q0 38-8 65t-19 42.5t-32 27.5t-37 17t-46.5 13.5T448 885v779q0 36-18 64q-17 29-46 46q-29 18-64 18q-36 0-64-18q-29-17-46-46q-18-28-18-64V885q-19-7-49.5-15.5T96 856t-37-17t-32-27.5T8 769t-8-65V64q2-28 19-45T64 0q28 2 45 19q19 19 19 45v416q0 26 19 45t45 19t45-19t19-45V64q2-28 19-45t45-19q28 2 45 19t19 45v416q0 26 19 45t45 19t45-19t19-45V64q2-28 19-45t45-19q28 2 45 19t19 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dsquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0m339 552q208 0 247 3q158 11 224 79q76 78 76 258q0 65-9.5 116t-27.5 87t-44 62t-58.5 41.5t-71 25T929 1236t-92.5 4.5t-100.5-.5h-61zM428 374v1044q0 31 22 53t52 22h406q280 0 423-150t143-450q0-594-571-594H502q-30 0-52 22t-22 53"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoorClosed(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 8v1600h1200V8zm194 865q-28 0-50-20.5T122 803t22-50t50-21q26 0 49 21.5t23 49.5t-22.5 49t-49.5 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoorOpen(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 200H0v1600h400v200l800-200V200L400 0zM100 300h300v1400H100zm479 779q-24 0-43.5-25T516 995t19.5-59.5T579 910t43 25t19 60q0 34-19 59t-43 25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drumstick(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1381 1308l-143-170l13-11q32-33 30.5-78.5T1254 962q-14-21-21.5-36t-13-27.5t-7.5-35t-3.5-38.5t-3.5-60.5t-4-81.5q0-119-61-243.5T983 219T761 61.5T512 0Q403 0 307 46.5t-163 125t-105.5 184T0 576q0 88 42.5 186T149 939.5T283.5 1081t130.5 91q42 20 184.5 64.5T791 1301q20 8 51 23t42 18.5t32.5 5t42-8t50.5-29.5l168 200q-40 48-40 110q0 71 50.5 121.5T1310 1792q71 0 121.5-50.5T1482 1620q72 0 127-50.5t55-121.5q0-70-55-121.5t-127-51.5q-57 0-101 33M512 209q4-12 13-22.5t26.5-24T572 146q30-26 84 5q170 91 276.5 236.5T1064 719q5 38-15 65q-18 24-62 46q-16 8-33.5-8.5T935 780q-4-118-63.5-225T722 372.5T522 250q-21-9-10-41m803 1493q-47 0-69-24t-22-59q0-32 16.5-66.5t42.5-57.5l-209-246q-18-21 1-35l46-35q19-13 36 5l207 242q8-6 26-20.5t29.5-22.5t29-15.5t35.5-7.5q33 0 61.5 27.5t28.5 64.5q0 51-50 73t-100 9q-15-4-17 2t3 19q12 53-15 100t-80 47"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Esquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0m339 546h584q31 0 53-21.5t22-52.5v-98q0-31-22-53t-53-22H502q-30 0-52 22t-22 53v1044q0 31 22 53t52 22h757q31 0 53-22t22-53v-98q0-31-22-52.5t-53-21.5H677l-2-226h423q31 0 53-22t22-53v-98q0-31-22-53t-53-22H675z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditPage(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m806 1199l58-58l-149-149l-58 58v68h81v81zm331-588q0-7-4-10q-3-4-10-4q-3 0-5 1q-4 2-5 3L769 945q-1 1-3 5q-1 2-1 5q0 7 4 11q3 3 10 3h5q4-2 6-4l343-343q3-3 3-5q1-2 1-6m-34-122l264 264l-527 527H576v-264zm433 61q0 16-6 31t-17 26l-106 105l-263-263l105-105q8-9 26-18q15-6 31-6t31 6q18 9 27 18l149 148q11 13 17 27q6 15 6 31M128 1408h1024v-320l128-128v480q0 25-13 49q-14 21-35 34q-20 13-48 13H96q-28 0-48-13q-21-13-35-34q-11-22-13-49V96q2-28 13-48q13-22 35-35Q71 0 96 0h640q29 0 58 10q29 8 59 24q24 12 47 34l251 253l-93 93l-249-256q-10-11-37-20q-34-10-52-10H128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EightSquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1456 1H336Q197 1 98.5 99T0 336v1120q0 139 98.5 237.5T336 1792h1120q139 0 237.5-98.5T1792 1456V336q0-139-98.5-237T1456 1M896 299q100 0 187.5 29.5T1229 405t91 104t33 112q0 107-30.5 157T1201 874q99 50 147.5 111.5T1397 1149q0 111-68.5 190T1149 1455.5T896 1493t-253-37.5T463.5 1339T395 1149q0-102 48.5-163.5T591 874q-91-46-121.5-96T439 621q0-55 33-112t91-104t145.5-76.5T896 299m0 979q-95 0-162.5-42T666 1135q0-60 68-107t162-47t162 47t68 107q0 59-67.5 101T896 1278m0-775q-93 0-150.5 41T688 646q0 59 56.5 95T896 777t151.5-36t56.5-95q0-61-57.5-102T896 503"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fsquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 1h1120q139 0 237.5 98t98.5 237v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 99T336 1m339 545h584q31 0 53-21.5t22-52.5v-98q0-31-22-53t-53-22H502q-30 0-52 22t-22 53v1045q0 31 22 52.5t52 21.5h98q31 0 53-21.5t22-52.5v-399h476q31 0 52.5-22t21.5-53v-98q0-31-21.5-53t-52.5-22H675z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceAllergy(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 0Q663 0 508.5 66T243 243T66 508.5T0 832t66 323.5T243 1421t265.5 177t323.5 66t323.5-66t265.5-177t177-265.5t66-323.5t-66-323.5T1421 243T1155.5 66T832 0m128 512l64-64l128 128l128-128l64 64l-128 128l128 128l-64 64l-128-128l-128 128l-64-64l128-128zm-128 640l-160 160l-160-160l-128 128l-64-64l195-192l157 160l160-160l160 160l160-160l192 192l-64 64l-128-128l-160 160zM704 512L576 640l128 128l-64 64l-128-128l-128 128l-64-64l128-128l-128-128l64-64l128 128l128-128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceDislike(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 0q169 0 323 66t265.5 177.5T1598 509t66 323t-66 323.5t-177.5 265.5t-265.5 177t-323 66t-323-66t-265.5-177T66 1155.5T0 832t66-323t177.5-265.5T509 66T832 0M607 1265q122-111 225-113q102 0 225 113q20 19 44.5 14.5t38.5-29.5q30-50-12-89q-148-137-296-137t-296 137q-42 40-12 89q15 25 39 29.5t44-14.5M384 704q0 53 37.5 90.5T512 832t90.5-37.5T640 704t-37.5-90.5T512 576t-90.5 37.5T384 704m768 128q53 0 90.5-37.5T1280 704t-37.5-90.5T1152 576t-90.5 37.5T1024 704t37.5 90.5T1152 832"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceLike(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 0q169 0 323 66t265.5 177.5T1598 509t66 323t-66 323.5t-177.5 265.5t-265.5 177t-323 66t-323-66t-265.5-177T66 1155.5T0 832t66-323t177.5-265.5T509 66T832 0M384 640q0 53 37.5 90.5T512 768t90.5-37.5T640 640t-37.5-90.5T512 512t-90.5 37.5T384 640m448 704q137 0 243.5-89.5T1216 1025H448q34 141 140 230t244 89m320-576q53 0 90.5-37.5T1280 640t-37.5-90.5T1152 512t-90.5 37.5T1024 640t37.5 90.5T1152 768"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileDownload(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1440V96q0-28 13-48q13-21 34-35Q71 0 96 0h640q28 0 58 10q31 9 59 24q26 14 47 34q35 35 98 98.5T1113 282t99 98q20 21 34 47q15 28 24 59q10 30 10 58v896q0 25-13 49q-14 21-35 34q-20 13-48 13H96q-27 0-49-13q-21-13-34-34q-13-22-13-49M768 416q0-14-9-23t-23-9H544q-14 0-23 9t-9 23v416H288q-9 0-17 5q-8 7-12 14q-3 10-2 18q1 9 8 17l351 384q3 3 11 7l13 3q7 0 13-3t10-7l353-384q7-8 8-17q0-12-3-18q-3-9-11-14q-10-5-18-5H768z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileDownloadO(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1440V96q0-28 13-48q13-21 34-35Q71 0 96 0h640q28 0 58 10q31 9 59 24q26 14 47 34q35 35 98 98.5T1113 282t99 98q20 21 34 47q15 28 24 59q10 30 10 58v896q0 25-13 49q-14 21-35 34q-20 13-48 13H96q-27 0-49-13q-21-13-34-34q-13-22-13-49m1152-32V558q0-35-19-70q-8-14-11-17L809 158q-9-10-37-20q-34-10-53-10H128v1280zM768 416v416h224q8 0 18 5q8 5 11 14q3 6 3 18q-1 9-8 17l-353 384q-4 4-10 7t-13 3q-9-1-13-3q-8-4-11-7L265 886q-7-8-8-17q-1-8 2-18q4-7 12-14q8-5 17-5h224V416q0-14 9-23t23-9h192q14 0 23 9t9 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileMoveO(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M672 768h736V544q0-9 5-17q3-7 14-12q10-3 18-2q10 2 17 7l384 353q6 6 7 10q3 6 3 13t-3 13q-1 5-7 11l-384 351q-6 6-17 8q-8 1-18-2q-9-4-14-11q-5-8-5-18v-224H672q-15-2-23-10q-9-8-9-22V800q0-14 9-23t23-9M0 544v896q0 25 13 49q14 21 35 34q20 13 48 13h1088q27 0 49-13q21-13 34-34q13-22 13-49v-320h-128v288H128V640h416q27 0 49-13q21-13 34-34q13-22 13-49V128h512v544h128V96q0-28-13-48q-13-21-34-35q-24-13-49-13H544q-26 0-58 9q-27 8-59 25q-29 16-47 34L68 380q-22 23-34 47q-15 28-24 59q-10 30-10 58m512-32H136q4-11 11-24q6-12 11-17l313-313q4-4 17-12q21-9 24-10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fish(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 1344q-557 0-832-576q69-95 142.5-169t146-122t141-80t141-48.5t133-22.5t128.5-6q117 0 280 57.5T1421 528t208 187q77-81 123-127.5t100.5-93.5t101-72.5T2048 384q-7 24-28 83.5t-32 96t-21.5 96.5t-10.5 119q0 103 25 230t67 207q-41-2-88-17.5t-102-46.5t-114-87t-111-130q-49 74-132 146t-186.5 131t-231 95.5T832 1344M448 512q-81 0-136.5 55.5T256 704t55.5 136.5T448 896t136.5-55.5T640 704t-55.5-136.5T448 512m3 97q-41 0-68.5 27T355 703q0 41 28 71t68 30t69.5-30t29.5-71q0-40-29-67t-70-27"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FiveSquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1456 1H336Q197 1 98.5 99T0 336v1120q0 139 98.5 237.5T336 1792h1120q139 0 237.5-98.5T1792 1456V336q0-139-98.5-237T1456 1M400 1208q-7-19-9.5-30t-1-24.5t14.5-20t38-6.5h143q15 0 21 1.5t15.5 11.5t24.5 34q32 53 98.5 85t147.5 32q111 0 190-57.5t79-139.5q0-81-79-138.5T892 898q-95 0-168 43q-9 5-16.5 10.5T697 959t-10 3t-21 1H507q-38 0-66-35.5T413 847l30-479q0-29 22-49t52-20h766q29 0 50.5 21t21.5 50v77q0 29-21.5 50t-50.5 21H683l-14 217q107-40 223-40q139 0 256.5 53.5t186 145.5t68.5 200q0 109-68.5 201t-186 145t-256.5 53q-97 0-194.5-30.5t-181-97.5T400 1208"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Floors(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M33 431q-18-9-25.5-19.5T0 383t7.5-28.5T33 335L670 17q39-19 98-17q59-2 98 17l637 318q18 9 25.5 19.5t7.5 28.5t-7.5 28.5T1503 431L866 749q-39 19-98 17q-59 2-98-17zm0 770q-18-9-25.5-19.5T0 1153t7.5-28.5T33 1105l160-80l477 238q40 19 98 16q58 3 98-16l477-238l160 80q18 9 25.5 19.5t7.5 28.5t-7.5 28.5t-25.5 19.5l-637 318q-40 19-98 16q-58 3-98-16zm0-384q-18-9-25.5-19.5T0 769t7.5-28.5T33 721l160-80l477 238q40 19 98 16q58 3 98-16l477-238l160 80q18 9 25.5 19.5t7.5 28.5t-7.5 28.5T1503 817l-637 318q-40 19-98 16q-58 3-98-16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flower(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M896 1792q277 0 516-113.5t380-334.5q-80-50-240-45t-298 65q-90 39-159 98t-96 113v-449q201-28 332-153t131-326V81q0-34-23.5-57.5T1381 0q-21 10-67.5 56.5t-103 109T1139 243q-13-13-71-80.5t-105.5-115T896 0q-20 0-67.5 47.5t-105 115T653 243q-15-15-71.5-77.5t-103-109T410 0q-33 0-57 23.5T329 81v566q0 102 36.5 187.5T466 979t147 96.5t179 50.5v449q-26-53-95-112.5T538 1364q-138-60-298-65T0 1344q142 221 380.5 334.5T896 1792"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FourSquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1456 1q139 0 237.5 98t98.5 237v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 99T336 1zM384 926q-15 17-24.5 37.5t-15 34.5t-8 41t-3 37.5t0 45.5t.5 44q0 31 20.5 53t49.5 22h576v178q0 30 22.5 52t52.5 22h127q31 0 53-21.5t22-52.5v-178h132q29 0 49.5-22t20.5-53v-105q0-31-20.5-53t-49.5-22h-132V373q0-30-22-52t-53-22h-127q-149 0-171 25zm256 60q106-123 340-401v401z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gsquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0m819 1097q-36 136-242 136q-140 0-238.5-98.5T576 897t98.5-237.5T913 561q128 0 224 86q10 9 18 18t12 12.5t6.5 6.5t10.5 3.5t14.5 1t27.5.5h192q59 0 38-54q-72-161-216.5-248T915 301q-121 0-231.5 47T493 475T365.5 665T318 897q0 162 80 299.5T615.5 1414t299.5 80q106 0 203.5-51t160.5-124v100q0 31 21.5 53t52.5 22h65q31 0 53-22t22-53V952q0-74-75-74h-379q-30 0-52 21.5T965 952v71q0 31 22 52.5t52 21.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gantt(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M198 992h154V480H70q-28 0-49-21T0 410V166q0-28 21-49t49-21h282V32q0-14 9-23t23-9t23 9t9 23v64h448V32q0-14 9-23t23-9t23 9t9 23v64h154q29 0 49.5 20.5T1152 166v244q0 28-21 49t-49 21H928v64h448V32q0-14 9-23t23-9t23 9t9 23v512h282q29 0 49.5 20.5T1792 614v244q0 28-21 49t-49 21h-282v512q0 14-9 23t-23 9t-23-9t-9-23V928H928v64h282q28 0 49 21t21 49v244q0 28-21 49t-49 21H928v64q0 14-9 23t-23 9t-23-9t-9-23v-64H416v64q0 14-9 23t-23 9t-23-9t-9-23v-64H198q-28 0-49-21t-21-49v-244q0-28 21-49t49-21m666 0v-64H710q-28 0-49-21t-21-49V614q0-29 20.5-49.5T710 544h154v-64H416v512z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GanttO(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1408 106v250q0 18-9 30t-23 12H32q-14 0-23-12t-9-30V106q0-18 9-30t23-12h1344q14 0 23 12t9 30m512 400v250q0 18-9 30t-23 12H544q-14 0-23-12t-9-30V506q0-16 9-28t23-14h1344q14 2 23 14t9 28m-768 402v250q0 17-9.5 29.5T1120 1200H288q-12 0-22-13t-10-29V908q0-18 10-30t22-12h832q14 0 23 12t9 30m512 400v250q0 17-9.5 29.5T1632 1600H800q-12 0-22-13t-10-29v-250q0-16 10-29t22-13h832q13 0 22.5 12.5t9.5 29.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globe(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 0Q663 0 508.5 66T243 243T66 508.5T0 832t66 323.5T243 1421t265.5 177t323.5 66t323.5-66t265.5-177t177-265.5t66-323.5t-66-323.5T1421 243T1155.5 66T832 0m0 1536q-113 0-221-36q16-8 39.5-30t44.5-23q20-1 40-5t46.5-11.5T824 1419q13-3 82.5-29.5T1016 1360q57-5 120 8.5t84 50.5q-176 117-388 117M175 580q68-178 218.5-298.5T735 135q-10 19-45.5 41.5T647 206q-23 21-47.5 36T560 269t-29 34q-13 20-47.5 54T423 419.5T397 464q0 27 29 53t48 19q18-7 62.5-5.5T609 542q6 2 26 6l42.5 8.5L728 571l52.5 22.5l46 33.5l33 47l10 64l-18.5 83q-10 26-38.5 46.5t-57.5 33t-58.5 43T655 1019q-23 90-24 129q-1 11 5 67t4 94.5t-20 38.5q-28 0-88.5-72.5T471 1181q0-6-14.5-76.5T442 989q0-25-20.5-45.5t-44.5-33t-44.5-38.5t-20.5-60q0-39 17.5-76t31-62.5T367 630q-5-12-20-21t-41-14t-44-7.5t-50.5-5T175 580m881-416q212 71 346 254.5T1536 832q0 240-146 429v-1q-23-18-33-51.5t6-79.5q18-39 25-117t2-122q-3-24-9.5-49.5t-18-56t-32-50T1285 715q-57-1-95-25t-69-86q-18-36-19-67.5t11-54.5t29.5-43.5t35-37.5t29.5-33.5t11-35t-19-38.5q-9-10-26-.5t-40 16.5t-44.5 9t-39-22.5T1027 219q-1-19 29-54z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grapes(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 38q27 117 77.5 184.5T215 314t165.5 28.5T589 339q61-3 93-4Q629-106 16 38m271 366q-104 0-178 74.5T35 657q0 103 74.5 177.5T287 909q60 0 112-26t87-72q-37-71-37-154q0-84 37-155q-35-46-87-72t-112-26m554 6q3-104 69.5-175T1074 133t218-33V0q-86 1-164 17T980.5 65.5t-120.5 80t-81 115T749 409q-90 14-149.5 84.5T540 657q0 104 74 178t178 74t178-74t74-178q0-91-58-160.5T841 410m455-6q-60 0-112 26t-87 72q37 71 37 155q0 83-37 154q35 46 87 72t112 26q104 0 178-74t74-178t-74-178.5t-178-74.5m-23 594q-139-9-229-109q-92 100-230 109q-22 48-22 105q0 104 74 178t178 74t178-74t74-178q0-54-23-105m-535 260q-37-72-37-155q0-57 19-112q-109-23-180-102q-92 100-230 109q-23 58-23 105q0 104 74.5 178t178.5 74q59 0 111-26t87-71m283 187q-138-11-229-111q-92 100-231 111q-21 50-21 103q0 104 74 178t178 74t178-74t74-178q0-55-23-103"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hsquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1493 1419V373q0-31-22-52.5t-53-21.5h-149q-31 0-52.5 21.5T1195 373v374H597V373q0-31-21.5-52.5T523 299H374q-31 0-53 21.5T299 373v1046q0 31 22 52.5t53 21.5h149q31 0 52.5-21.5T597 1419v-374h598v374q0 31 21.5 52.5t52.5 21.5h149q31 0 53-21.5t22-52.5m299-1083v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0h1120q139 0 237.5 98.5T1792 336"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Highchair(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 896V512h229q11 0 19-8t8-19v-74q0-11-8-19t-19-8H256V24q0-10-7-17t-17-7h-80q-10 0-17 7t-7 17v788q0 34 17 57.5t47 26.5L0 1792h130l94-384h464l77 384h131L704 896zm-104 384H248l72-384h256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hourglass(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M86 192H32q-14 0-23-9t-9-23V96q0-14 9-23t23-9h960q14 0 23 9t9 23v64q0 14-9 23t-23 9h-54q-8 201-77.5 332.5T668 698v148q121 42 190.5 171.5T938 1344h54q14 0 23 9t9 23v64q0 14-9 23t-23 9H32q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h54q10-197 79.5-326.5T356 846V698q-123-42-192.5-173.5T86 192m784 0H154q3 75 15.5 141.5t39 131.5t75 111T398 639l26 6v254l-26 7q-65 17-113 61.5T210.5 1076T171 1205t-17 139h87l271-271l271 271h87q-4-74-17-139t-39.5-129T739 967.5T626 906l-26-7V645l26-6q66-17 114.5-63t75-111t39-131.5T870 192M363 502h298L512 650z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Isquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0m864 299H591q-30 0-52 22t-22 53v149q0 31 22 52.5t52 21.5h156v598H591q-30 0-52 21.5t-22 52.5v149q0 31 22 53t52 22h609q31 0 53-22t22-53v-149q0-31-22-52.5t-53-21.5h-155V597h155q31 0 53-21.5t22-52.5V374q0-31-22-53t-53-22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IdBadgeAlt(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2368 448v1152q0 78-57.5 135t-134.5 57H192q-77 0-134.5-57T0 1600V448q0-76 58-134t134-58h704l-1 158h-30q-16 0-27 10.5T827 449v69q0 14 11 23.5t27 9.5h643q16 0 26.5-9.5T1545 518v-69q0-14-10.5-24.5T1508 414h-31l-2-158h701q77 0 134.5 57.5T2368 448m-1331 968q0-64-9-118.5t-29-105t-60-79.5t-95-29q-76 76-181 76q-107 0-183-76q-55 0-94.5 29t-59.5 79.5t-28.5 105T289 1416q0 63 36 108.5t88 45.5h497q52 0 89.5-45.5T1037 1416M887 898q0-94-65.5-160.5T663 671t-158.5 66.5T439 898q0 92 65.5 157.5T663 1121t158.5-65.5T887 898m1195 560v-75q0-15-10-26t-25-11h-825q-15 0-26.5 11t-11.5 26v75q0 15 11.5 26t26.5 11h825q15 0 25-10.5t10-26.5m-447-298v-76q0-16-11-26.5t-26-10.5h-376q-15 0-26.5 11t-11.5 26v76q0 15 11.5 26t26.5 11h376q15 0 26-10.5t11-26.5m447 0v-76q0-16-10-26.5t-25-10.5h-225q-16 0-27 10.5t-11 26.5v76q0 16 11 26.5t27 10.5h225q15 0 25-10.5t10-26.5m0-300v-75q0-16-10-26.5t-25-10.5h-825q-15 0-26.5 11t-11.5 26v75q0 16 11.5 27t26.5 11h825q15 0 25-11t10-27M1049 0q-24 0-41 12.5T991 43v397q0 17 17.5 30t40.5 13h274q24 0 40.5-13t16.5-30V43q0-18-16.5-30.5T1323 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IdCard(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M180 11q-75 0-127.5 52.5T0 191v1167q0 74 52.5 127t127.5 53h1741q75 0 127.5-53t52.5-127V191q0-75-52.5-127.5T1921 11zm1798 1403H123V135h1855zM1469 358q-112 0-174 76.5T1233 631q-2 133 117 235q15 12 9 30l-18 50q-8 17-25 26.5t-70 27.5q-46 8-89 24v167h624v-167q-44-16-88-24q-53-18-70.5-27.5T1597 946l-17-50q-6-18 9-30q119-102 117-235q0-120-62.5-196.5T1469 358M302 445v104h649V445zm0 278v104h649V723zm0 277v104h649v-104z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IdCardAlt(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2368 192v1280q0 78-57 135t-135 57H192q-77 0-134.5-57T0 1472V192q0-77 57.5-134.5T192 0h1984q78 0 135 57.5t57 134.5m-1333 940q0-65-8.5-120t-29-106.5t-60-81T842 795q-79 76-185 76q-110 0-186-76q-56 0-96 29.5t-60 81t-28.5 106.5t-8.5 120q0 64 37 109.5t89 45.5h506q52 0 88.5-45.5T1035 1132M884 607q0-95-66.5-162.5T657 377t-161 67.5T429 607q0 93 67 159.5T657 833t160.5-66.5T884 607m1213 567v-76q0-16-11.5-27t-26.5-11h-836q-15 0-26.5 11t-11.5 27v76q0 16 11.5 27t26.5 11h836q15 0 26.5-11t11.5-27m-454-303v-76q0-16-11.5-27t-26.5-11h-382q-15 0-26.5 11t-11.5 27v76q0 17 11.5 28t26.5 11h382q15 0 26.5-11t11.5-28m454 0v-76q0-16-11.5-27t-26.5-11h-229q-16 0-26.5 10.5T1793 795v76q0 17 10.5 28t26.5 11h229q15 0 26.5-11t11.5-28m0-303v-76q0-16-11.5-27t-26.5-11h-836q-15 0-26.5 11t-11.5 27v76q0 16 11.5 27.5T1223 607h836q15 0 26.5-11.5T2097 568"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Jsquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0m562 1221q-103 0-155.5-40T690 1051v-84q0-31-22-53t-53-22H466q-31 0-53 22t-22 53v48q0 58 16.5 120t55.5 127t95 115.5t144.5 83T898 1493t195.5-32.5t144.5-83t94.5-115.5t55-127t16.5-120V374q0-31-21.5-53t-52.5-22h-149q-31 0-53 22t-22 53v677q0 90-52.5 130T898 1221"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ksquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0m28 1419q0 31 22 52.5t53 21.5h149q31 0 53-21.5t22-52.5v-374h123q35 0 55.5 11t45.5 41l204 321q19 30 45 52.5t53 22.5h188q55 0 56-32q0-17-38-83l-314-482l314-482q38-66 38-83q-1-32-56-32h-188q-27 0-53 22.5t-45 52.5L887 695q-25 30-45.5 41T786 747H663V373q0-31-22-52.5T588 299H439q-31 0-53 21.5T364 373z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Kakao(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M896 12q243 0 449.5 94.5t326.5 257T1792 718t-120 355t-326 257.5t-450 94.5q-77 0-159-11q-356 247-379 250q-11 4-21-1q-4-3-6-8t-2-9v-4q6-39 91-325q-193-96-306.5-254.5T0 718q0-192 120-354.5t326.5-257T896 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KakaoSquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 1652V140q0-58-41-99t-99-41H140Q82 0 41 41T0 140v1512q0 58 41 99t99 41h1512q58 0 99-41t41-99M896 252q198 0 365.5 77t265 209t97.5 288t-97.5 288t-265 209t-365.5 77q-66 0-129-9q-68 49-180.5 125T459 1594q-9 4-17-1q-4-2-5.5-6.5t-1.5-7.5v-3q1-5 74-264q-156-77-248.5-206T168 826q0-156 97.5-288t265-209T896 252"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Kakaotalk(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M896 12q243 0 449.5 94.5t326.5 257T1792 718t-120 355t-326 257.5t-450 94.5q-77 0-159-11q-356 247-379 250q-11 4-21-1q-4-3-6-8t-2-9v-4q6-39 91-325q-193-96-306.5-254.5T0 718q0-192 120-354.5t326.5-257T896 12M442 905V598h79q16 0 33.5-11.5T572 555t-16-33.5t-31-13.5H265q-17 0-32 10t-15 33t16 35t36 12h79v307q0 25 12 41t34 16t34.5-18t12.5-39m433 57q37-14 23-71q-4-13-56.5-155.5T781 573q-27-65-74-65q-51 0-79 65q-8 17-61 165.5T512 891q-7 9-3.5 33t21.5 33q20 9 39 2t26-21l24-66h174q11 36 16 47q5 10 11 20t22 19.5t33 3.5m269 0q21 0 38.5-13t17.5-35t-16-34.5t-40-12.5h-111V565q0-21-12.5-39T986 508t-34 16t-12 41v340q0 25 12 41t34 16q1 0 3.5-.5t3.5-.5t3.5.5t3.5.5zm394-7q17-16 17-36.5t-13-36.5q-6-9-126-169q58-59 117-118q15-15 18.5-35.5T1540 521q-17-16-36.5-13t-36.5 20q-3 3-26 26.5t-59.5 60.5t-64.5 66V565q0-21-12.5-39t-34.5-18t-34 16t-12 41v340q0 25 12 41t34 16t34.5-18t12.5-39v-91q4-4 15-15.5t18-18.5q66 90 118 159q14 19 32.5 24.5t37.5-8.5M646 782l61-179l60 179z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KakaotalkSquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 1652V140q0-58-41-99t-99-41H140Q82 0 41 41T0 140v1512q0 58 41 99t99 41h1512q58 0 99-41t41-99M896 252q198 0 365.5 77t265 209t97.5 288t-97.5 288t-265 209t-365.5 77q-66 0-129-9q-68 49-180.5 125T459 1594q-9 4-17-1q-4-2-5.5-6.5t-1.5-7.5v-3q1-5 74-264q-156-77-248.5-206T168 826q0-156 97.5-288t265-209T896 252M527 978V728h64q14 0 28-9t14-26q0-16-13-27t-25-11H383q-38 0-38 35q0 19 13 28.5t29 9.5h64v250q0 20 10 33t28 13q17 0 27.5-14.5T527 978m352 46q32-13 18-58q-2-10-44-125.5T803 708q-22-53-61-53q-41 0-64 53q-6 14-49 134t-45 124q-5 7-2 26.5t17 27.5q16 7 31.5 1.5T652 1004l19-53h142q9 29 13 38q21 46 53 35m219 0q17 0 31-10.5t14-28.5t-12.5-28t-32.5-10h-91V701q0-17-10-31.5T969 655t-27.5 13t-9.5 33v277q0 20 9.5 33t27.5 13q1 0 3-.5t3-.5t3 .5t3 .5zm320-5q14-13 13.5-30t-10.5-30q-1-1-16-21t-41-55t-45-61q94-96 95-97q12-12 14.5-28t-9.5-31q-26-28-59 5q-1 1-21 21.5t-49 50t-52 52.5v-94q0-17-10.5-31.5T1200 655q-18 0-28 13t-10 33v277q0 20 10 33t28 13q17 0 27.5-14.5T1238 978v-74l13-13l14-15q89 120 95 129q12 16 27 20t31-6M693 878l49-146l50 146z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KanjiChu(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m530 1079l-226-75q-32 112-115 313q-87 212-178 344l225 131q87-140 174-351q90-218 120-362m-67-330Q347 633 133 471L0 664q185 148 322 290zm83-435q-96-96-333-282L71 219q186 156 330 300zm1246 1442v-233h-503v-387h422V908h-422V596h471V371h-432l56-45Q1231 163 1049 0L850 151q148 140 222 220H553v225h469v312H585v228h437v387H480v233z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KanjiUtage(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1452 1282h340v-179H761q30-56 45-88h708V411H264v604h264q-14 32-43 88H0v179h384l-55 85q-11 19-11 37q0 26 20 46t38 23q129 13 332 52q-340 76-676 76l49 191q629-23 912-149l36-16q488 122 568 163l138-215q-117-39-429-114q91-88 146-179m-445 134q-246-53-384-79l33-55h499q-66 83-148 134m252-647v99H514v-99zM514 655v-98h745v98zm1259-107V145h-758V0H754v145H6v420h217V325h1341v223z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KanjiYubi(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1482 1741v64h260V892H738v913h251v-64zm0-516H989v-130h493zm-493 315v-132h493v132zm-292-520l-21-207q-38 15-151 47V575h159V370H525V31H274v339H53v205h221v351Q93 969 14 979l49 232l211-56v373q0 59-63 59H57l29 232h219q94 0 157-52q63-53 63-148v-536q114-37 172-63m1097-433V359l-218-26q0 188-25 213q-27 25-263 25q-93 0-171-5q-61-3-87-7q-18-4-27-17q-10-15-10-34v-75q578-100 774-185L1634 43q-168 83-641 178V27H740v552q0 94 54 145q55 52 199 61q139 10 266 10q369 0 453-52q82-52 82-156"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyboard(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1600 192q132 0 226 94t94 226v512q0 132-94 226t-226 94H320q-132 0-226-94T0 1024V512q0-132 94-226t226-94zM320 1216h1280q80 0 136-56t56-136V512q0-80-56-136t-136-56H320q-80 0-136 56t-56 136v512q0 80 56 136t136 56m-64-768v128h128V448zm256 0v128h128V448zm256 0v128h128V448zm256 0v128h128V448zm512 0v128h128V448zm-256 0v128h128V448zM256 704v128h128V704zm256 0v128h128V704zm256 0v128h128V704zm256 0v128h128V704zm512 0v128h128V704zm-256 0v128h128V704zM256 960v128h128V960zm1280 0v128h128V960zm-1024 0h896v128H512z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lsquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0m43 1419q0 31 22 52.5t53 21.5h884q31 0 53-21.5t22-52.5v-150q0-31-22-52.5t-53-21.5H678V373q0-31-22-52.5T603 299H454q-31 0-53 21.5T379 373z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Language(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M891 689q-6-87-66-144q-30-27-50-37.5T715 488q1-5 4.5-26t6.5-35l-91-12l-10 50q-79-14-136 3V354q151-6 306-23l-14-94q-126 21-286 29q2-30 6-49t6-37q-5 0-36-.5t-53-.5q-2 18-5.5 40t-5.5 51l-154 3l14 94l120-6l5 137q-100 50-150.5 114T191 781q0 64 30 97t75 33.5t94-17t89-49.5l18 36l72-46l-21-43q34-34 54-56.5t50.5-71.5T700 562q60 17 82 43t19 86q-48 10-80 48t-32 88v25q-37 10-81 13l46 73q5-1 35-6v68H179q-31 0-53-22.5T104 924V179q0-31 22-53t53-22h745q31 0 53.5 22t22.5 53v510zm212 0V138q0-57-41-97.5T965 0H138Q81 0 40.5 40.5T0 138v827q0 56 40.5 97t97.5 41h551v551q0 57 41 97.5t97 40.5h827q57 0 97.5-40.5t40.5-97.5V827q0-56-40.5-97t-97.5-41zM401 584q10 120 32 181q-17 16-40.5 31T343 821t-44-1.5t-18-47.5q0-60 33-111t87-77m90-33q54-26 119-7q-7 25-19.5 47.5t-22.5 35t-28 34.5l-27 33q-5-13-9-26t-6-21t-3.5-24t-2-20.5t-1-26t-.5-25.5m504 1069H863l310-758h135l310 758h-132l-89-218h-312zm246-599l-116 281h231z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Line(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 736q0-148-71-282.5t-191-232T1244 66T896 8T548 66T262 221.5t-191 232T0 736q0 177 97.5 331.5T366 1324t383 130q66 13 79 45q8 21 3 80q-11 64-13 77l-3 18l1 19.5l9.5 18l25.5 4l45-11.5q98-41 338.5-208.5T1600 1184q192-210 192-448M714 542q8 0 13 5t5 13v391q0 17-18 17h-63q-18 0-18-17V560q0-8 5.5-13t12.5-5zm433 0q17 0 17 18v391q0 17-17 17h-63q-10 0-15-7L890 718v233q0 17-17 17h-63q-18 0-18-17V560q0-18 18-18h63q9 0 14 7l179 243V560q0-8 5-13t13-5zM562 870q18 0 18 17v64q0 17-18 17H311q-18 0-18-17V560q0-18 18-18h63q17 0 17 18v310zm931-230h-170v67h170q18 0 18 17v62q0 8-5.5 13.5T1493 805h-170v65h170q18 0 18 17v64q0 17-18 17h-251q-17 0-17-17V560q0-18 17-18h251q7 0 12.5 5t5.5 13v63q0 17-18 17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineSquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1546 816q0-143-87.5-265T1221 358t-327-71t-327 71t-237.5 193T242 816q0 196 155 343t391 178q47 11 57 33q7 18 2 59q-7 46-9 56q-4 35 9 42q14 7 47-7q73-31 247.5-152.5T1407 1142q139-154 139-326m-217-69h-125v47h125q13 0 13 13v46q0 13-13 13h-125v47h125q13 0 13 13v46q0 13-13 13h-183q-12 0-12-13V688q0-13 12-13h183q13 0 13 13v46q0 13-13 13M651 913q13 0 13 13v46q0 13-13 13H469q-12 0-12-13V688q0-13 12-13h46q12 0 12 13v225zm426-238q13 0 13 13v284q0 13-13 13h-46q-7 0-11-6L891 803v169q0 13-13 13h-46q-13 0-13-13V688q0-13 13-13h46q4 0 10 6l130 176V688q0-13 13-13zm-315 0q13 0 13 13v284q0 13-13 13h-46q-13 0-13-13V688q0-13 13-13zM336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Msquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0h1120q139 0 237.5 98.5T1792 336m-896 779q39 0 77-55l222-312v671q0 31 21.5 52.5t52.5 21.5h149q31 0 53-21.5t22-52.5V373q0-31-22-52.5t-53-21.5h-149q-13 0-24.5 2t-20 4.5t-17.5 8.5t-14 9t-12.5 11t-10 11t-9.5 12t-8 10L896 729L639 367q-1-1-8-10t-9.5-12t-10-11t-12.5-11t-14-9t-17.5-8.5t-20-4.5t-24.5-2H374q-31 0-53 21.5T299 373v1046q0 31 22 52.5t53 21.5h149q31 0 52.5-21.5T597 1419V748l222 312q38 55 77 55"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnetNote(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1664 128v640q0 116-.5 170.5T1662 1082t-4.5 132t-9.5 108.5t-15.5 102.5t-22.5 84t-31.5 83t-42.5 72H0q79-242 103.5-447.5T128 640V128h442q-24 60-24 128q0 145 102.5 247.5T896 606t247.5-102.5T1246 256q0-68-24-128zM896 0q106 0 181 75t75 181t-75 181t-181 75t-181-75t-75-181t75-181T896 0m0 128q8 0 14-7t6-17t-6-17t-14-7q-73 0-124.5 51.5T720 256q0 7 8 13.5t16 6.5q10 0 17-6t7-14q0-52 38-90t90-38"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MaleFemale(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2291 261q0 109-75.5 185T2031 522t-186-76t-77-185t77-185t186-76t184.5 76t75.5 185m-1285 0q0 109-76 185t-185 76q-108 0-184.5-76.5T484 261t76.5-184.5T745 0q109 0 185 76t76 185m168 423l184 277l157-277q56-125 121-125h766q92 0 157.5 65.5T2625 782v484q0 46-33 79.5t-79 33.5t-78.5-33t-32.5-80V858h-73v1062q0 54-39 92t-93 38q-53 0-91-38t-38-92v-541h-75v541q0 54-39 92t-93 38q-53 0-91-38t-38-92V858h-45l-203 371q-16 34-43.5 54.5T1378 1304q-59 0-93-49l-263-397h-53v153l254 424l33 55q10 17 10 37q0 32-21 54t-52 22H969v317q0 54-38.5 92t-91.5 38H652q-53 0-91.5-38t-38.5-92v-317H299q-31 0-53.5-22.5T223 1527q0-17 10-37l289-479V858h-53l-265 397q-31 49-92 49q-46 0-79-32.5T0 1193q0-34 19-61l298-448q37-56 90.5-90.5T522 559h447q61 0 114.5 34.5T1174 684"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mic(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M816 516q-94 54-261.5 150.5T293 818Q177 782 99 688l-11 7q-12 7-26 3.5T40 683L7 626q-7-12-3-27t16-23l11-6q-53-138-2.5-280.5T210 72Q343-4 490 23t241 142l11-7q14-7 28-2.5t22 16.5l32 57q7 12 3.5 26T812 277l-13 7q42 111 17 232m253 1129q37 71 92.5 109t119.5 38q112 0 184.5-73.5T1537 1536q0-35-5.5-70t-11.5-59t-18.5-54.5t-18.5-43t-20-39t-15-29.5q-22-42-45.5-57t-38.5-7q-10 5-14 19t1 36.5t19 47.5q29 50 57.5 131.5T1456 1544q0 65-54.5 113t-125.5 48q-58 0-124-108l8-5q27-16 40-29.5t16-35.5t-12-50L826 612l-163 94l83 143q6 10 2.5 22.5T735 890l-39 22l-39 22q-11 7-22.5 3.5T616 923l-82-142l-163 94l560 760q28 37 57 40.5t65-20.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mobile(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M864 64q40 1 68 28q27 28 28 68v1344q-1 40-28 68q-28 27-68 28H96q-40-1-68-28q-27-28-28-68V160q1-40 28-68q28-27 68-28zm-432 96q-20 0-34 14t-14 34t14 34t34 14h96q20 0 34-14t14-34t-14-34t-34-14zm96 1344q20 0 34-14t14-34t-14-34t-34-14h-96q-20 0-34 14t-14 34t14 34t34 14zm336-192V352H96v960zm458-842q150 158 150 362t-150 362q-20 18-45 18q-26 0-45-19q-20-20-20-44q0-25 20-45q112-118 112-272t-112-272q-20-20-20-45q0-24 20-44q19-19 45-19q25 0 45 18m-180 180q74 80 74 182t-74 182q-20 18-46 18t-45-19t-19-45t18-46q38-38 38-90t-38-90q-18-19-18-46q0-26 19-45t45-19t46 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Moon(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1464 1090q-94 203-283 323.5T768 1536q-157-1-299-62t-244.5-163.5T61 1066T0 768q0-205 97.5-378t267-276.5T738 2q43-2 62 38q17 42-16 72q-176 164-176 400q0 111 43 211.5t115 173t172.5 116T1151 1056q119 0 228-51q41-18 72 13t13 72"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MultiArrow(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1100 1653l382-382h-260V442q39 70 89 143t91 126t110 131t106 117t122 127.5t116 121.5l-200 200h544V808l-200 200q-70-60-157.5-128T1694 766.5t-130.5-103T1447 559t-93-108.5t-72.5-124.5t-43-143.5T1222 8H978q0 94-16.5 177T919 333.5t-73.5 133t-92 117.5t-117 114t-130 110t-149 120T200 1058L0 858v550h550l-200-200q21-22 99-104t104.5-110.5t91-99t97-111.5t81-103.5t85-119.5T978 442v829H718z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Nsquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0h1120q139 0 237.5 98.5T1792 336m-685 1089q14 19 24 30.5t35.5 24.5t57.5 13h149q31 0 52.5-21.5t21.5-52.5V373q0-31-21.5-52.5T1373 299h-149q-31 0-53 21.5t-22 52.5v671L685 367q-13-19-24-31t-36-24.5t-57-12.5H419q-31 0-52.5 21.5T345 373v1046q0 31 21.5 52.5T419 1493h149q31 0 53-21.5t22-52.5V748z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Naver(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v1623h624V811l543 812h625V0h-625v811L624 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NaverSquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 0Q39 0 19.5 19T0 64v1664q0 26 19.5 45t44.5 19h1664q25 0 44.5-19t19.5-45V64q0-26-19.5-45T1728 0zm337 448h345l300 448V448h345v896h-345L746 896v448H401z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Neko(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 720q-49 0-80.5 31.5T352 832t31.5 80.5T464 944t80.5-31.5T576 832t-31.5-80.5T464 720m866-2q-50 0-82 32t-32 82t32 82t82 32t82-32t32-82t-32-82t-82-32M566 366q29-12 344-12t344 12q3 1 9.5-8t18-25.5T1308 294t37-48.5t46.5-53T1449 138t68-51q69-46 119.5-66T1744 1q30 59 15.5 303.5T1714 615q32 58 55 156.5t23 188.5q0 323-235 513.5T896 1664q-167 0-300.5-21.5t-247-72t-189-130T42 1240T0 960q0-136 13-207.5T61 616q-19-40-31.5-175.5t-9-270.5T44 1q26 0 44 3t36.5 11.5t28.5 15T189.5 56T234 87q75 49 156 121t126 116t50 42m733 887q0-29-10.5-47t-21.5-24.5t-32-13.5q-24 48-70 79.5t-84 32.5q-44 1-99-42.5t-86-85.5q-20 39-66.5 82t-86.5 44q-48 1-81-6t-54-26t-31-33.5t-25-43.5q-26-2-46 15.5t-19 41.5q-2 55 75 118.5t142 63.5q51 0 88-12t51-25t29.5-33.5T894 1312q2 1 25 25.5t39 36.5t52 24t78 10q50 0 130.5-59.5t80.5-95.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NekoSleep(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1299 1253q0-29-10.5-47t-21.5-24.5t-32-13.5q-24 48-70 79.5t-84 32.5q-44 1-99-42.5t-86-85.5q-20 39-66.5 82t-86.5 44q-48 1-81-6t-54-26t-31-33.5t-25-43.5q-26-2-46 15.5t-19 41.5q-2 55 75 118.5t142 63.5q51 0 88-12t51-25t29.5-33.5T894 1312q2 1 25 25.5t39 36.5t52 24t78 10q50 0 130.5-59.5t80.5-95.5M566 366q29-12 344-12t344 12q3 1 9.5-8t18-25.5T1308 294t37-48.5t46.5-53T1449 138t68-51q69-46 119.5-66T1744 1q30 59 15.5 303.5T1714 615q32 58 55 156.5t23 188.5q0 323-235 513.5T896 1664q-167 0-300.5-21.5t-247-72t-189-130T42 1240T0 960q0-136 13-207.5T61 616q-19-40-31.5-175.5t-9-270.5T44 1q26 0 44 3t36.5 11.5t28.5 15T189.5 56T234 87q75 49 156 121t126 116t50 42m906 539q23-71 28.5-129t-7.5-62q-7-2-25 11q-9 7-28.5 36.5t-39 57t-52 50T1280 890q-10 0-21-3t-22-9t-20.5-11.5t-21-14.5t-19-14.5t-18-15.5t-14.5-14l-13-13l-9-9q-21-24-29-21q-9 2-12 29t-2 49q-2 44 61.5 109t139.5 66q81 0 129.5-35t62.5-78m-1162-1q13 40 67 77t135 37q76-1 139.5-66T713 843q1-22-2-49t-12-29q-8-3-29 21l-9 9l-13 13l-14.5 14l-18 15.5l-19 14.5l-21 14.5L555 878l-22 9l-21 3q-36 1-68.5-21.5t-52-50t-39-57T324 725q-18-13-25-11q-8 5-11.5 30t2 68.5T310 904"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NineSquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1456 1H336Q197 1 98.5 99T0 336v1120q0 139 98.5 237.5T336 1792h1120q139 0 237.5-98.5T1792 1456V336q0-139-98.5-237T1456 1M867 1493q-170 0-281.5-62.5T446 1299q-21-47-12.5-80.5T483 1185h161q34 1 46 12t17 33q7 36 38.5 58t67 27.5t75.5 3.5q50-2 91.5-35.5t65-82.5t35.5-102t10-100q-11 4-41.5 16.5T1003 1033t-40 13t-47 11t-45 3q-133 0-244.5-49t-177-135.5T384 685q0-105 65-194t177.5-140.5T871 299q129 0 229 49.5T1259.5 483t89 189t29.5 222q0 119-33 225.5t-95 190.5t-161 133.5t-222 49.5m-4-627q85 0 146-56.5t61-136.5t-60.5-137T863 479t-146.5 57T656 673t61 136.5T863 866"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ninja(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M320 512h896q96 0 142 86t40 229q-2 48-25.5 87.5t-57 61.5t-67 33.5t-61.5 11.5q-70 0-127.5-13T966 976.5T899.5 940T836 909t-68-13q-42 0-85 19.5t-78 43t-102.5 43T349 1021q-28 0-61.5-11.5T220 976t-57.5-61.5T137 827q-5-144 41-229.5T320 512m448 1024q156 0 298-61t245-163.5t163.5-245T1535 768q0-118-33-225l205 163l85-194l-313-34q-1-2-2-5t-2-5.5t-2-4.5l302-125l-127-173l-192 263q-95-193-280-310.5T768 0Q559 0 382.5 103T103 382.5T0 768t103 385.5T382.5 1433T768 1536M308 562l-52 142l125-3q-28 33-28 84q0 44 30.5 77.5T448 896q35 0 62.5-33t27.5-78q0-25-7-43l131 28l25-61zm920 0L849 709l25 61l131-28q-7 18-7 43q0 45 27.5 78t62.5 33q34 0 64.5-33.5T1183 785q0-51-28-84l125 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoCommentBubble(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 768q0 150-91.5 282T1450 1271L793 132q68-4 103-4q244 0 450 85.5t326 233T1792 768M539 21q256 444 457.5 793t362.5 629q9 14 4.5 31t-18.5 25l-108 62q-15 8-31.5 4t-24.5-19l-89-153q-95 15-196 15q-70 0-145-8q-198 175-460 242q-49 14-114 22q-17 2-30.5-9t-17.5-29v-1q-3-4-.5-12t2-10t4.5-9.5l6-9l7-8.5l8-9q7-8 31-34.5t34.5-38t31-39.5t32.5-51t27-59t26-76q-157-89-247.5-220T0 768q0-168 113-311.5T419 226l-59-102q-8-14-4-31t19-25L482 6q15-8 31.5-3.5T539 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoSmokingAlt(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1504 640q0-23-5-40t-17-27.5t-22.5-17t-33-9.5t-36.5-4.5t-43-2.5q-47-2-72-4t-72-13t-75.5-29.5t-66-56t-61-89.5T957 214.5T928 32q0-14 9-23t23-9t23 9t9 23q7 86 23 154t35.5 114t47.5 79.5t54.5 52t62 29.5t62.5 14t64 4q10 0 35.5.5t35 1t31.5 2.5t31 5t25.5 9t24 14t18 20t15 27t8 36t3.5 46v96q0 14-9 23t-23 9t-23-9t-9-23zm0-320q-2-160-160-160h-96q-14 0-23-9t-9-23t9-23t23-9h96q101 0 158.5 49.5T1560 292q53 0 85 9.5t61 40.5q56 60 54 170v224q0 14-9 23t-23 9t-23-9t-9-23V512q0-72-31-116t-97-44h-64zm-362 832l221 383q8 13 4 28.5t-18 23.5l-100 58q-13 8-29 3.5t-23-17.5L435 312q-8-14-4-29.5t18-23.5l100-57q13-8 29-4t24 18l355 616h451v320zM0 832h578l185 320H0zm1503 0q-12 0-21.5 9t-9.5 22v258q0 13 9.5 22t21.5 9h66q12 0 21.5-9t9.5-22V863q0-13-9.5-22t-21.5-9zm192 0q-12 0-21.5 9t-9.5 22v258q0 13 9.5 22t21.5 9h66q12 0 21.5-9t9.5-22V863q0-13-9.5-22t-21.5-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Osquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0m221 896q0-142 99.5-238.5T899 560q139 0 237.5 98.5T1235 896q0 142-99.5 239T893 1232q-139 0-237.5-98.5T557 896m339 597q162 0 299.5-80t217.5-217.5t80-299.5t-80-299.5T1195.5 379T896 299t-299.5 80T379 596.5T299 896t80 299.5T596.5 1413t299.5 80"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OneSquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0m709 387q0-48-17-68t-62-20H747q-19 48-76.5 110.5T568 488q-18 6-33.5 30T519 563v98q0 30 22 52t53 22q20 0 38.5-5t30.5-10.5t27-19t20.5-19t20-23.5t16.5-20v557H591q-30 0-52 21.5t-22 52.5v149q0 31 22 53t52 22h609q31 0 53-22t22-53v-149q0-31-22-52.5t-53-21.5h-155z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Picon(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v1408h384V896h320q198 0 322.5-119T1151 448t-124.5-329T704 0zm384 640V256h192q89 0 140.5 48.5T768 448t-53.5 143.5T576 640z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Psquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0m339 1054h334q147 0 251-110t104-267q0-154-102-266q-104-112-253-112H502q-30 0-52 22t-22 53v1044q0 31 22 53t52 22h98q31 0 53-22t22-53zm0-531h334q50 0 85 45t35 109q0 63-35.5 108t-84.5 45H675z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Panther(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1332V232q506-12 736 6q28-43 60-77t84-69.5t130-55T1185 17q-12 229 24 299q44 14 89.5 38.5t78 46.5t78 62t67 60.5T1589 589q49 37 162 222t181 333l68 88l-58 104l-29 93q-52 90-109.5 133.5T1649 1606q38-64 20.5-119.5t-74-92T1471 1357q-65-1-118 41t-47 103t86 104t153 38l-106 53l114 18q-385 153-483 262l51-185q-76-223-235-422t-342-305q-64 16-277.5 119T0 1332M1079 328q-16-136-9-202q-186 62-207 163q108 13 216 39m166 386l167 217l177-25q-59-63-86-89.5t-68-53t-82-35.5t-108-14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Party(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1282 626q-95-100-198-174T891 341t-165-35.5T609 352q-10 10-17 22.5T582 394t-8 25t-7 23L195 1666l23 24q-39 27-69.5 63T95 1812t-55 32q-7 2-17.5 13.5T12 1878q0 18 18 30t40 12q28 0 54-22.5t52.5-57.5t38.5-46q26-23 58-44l26 25l1182-401q39-12 60-36q42-45 43.5-124.5t-33.5-175t-104.5-204T1282 626m-108 109q153 156 232.5 298.5T1445 1221t-179-33t-290-233Q822 799 742.5 657T704 469q41-45 179.5 33T1174 735m-219 414q73 65 145 114q1 2 50 94l-84 29q-60-113-111-237M614 691q22 49 61 119q32 144 100.5 312T925 1437l-83 28q-94-174-164-364.5T592 764zm-116 366q70 224 197 460l-82 29q-99-190-156-357zm-115 365q39 92 83 175l-82 29q-14-27-40-83zM1118 0l50 151l-129 94h159l49 151l49-151h158l-128-94l50-150l-129 93zM838 0l-27 137l137 27l27-137zm901 412l110-115l143 70l-74-140l111-113l-157 27l-74-141l-22 158l-157 27l143 69zm-37 246l133 88l-42 154l124-100l132 88l-55-150l124-98l-159 7l-56-149l-41 153zm-57-371l-298 244l37 45l298-244zm104 472l-233-21l-6 58l234 21zm23 280l-43 134l133 43l44-133z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PeopleGroup(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2099 482h221q126 0 203 58.5t77 186.5v338q0 35-.5 51.5t-5 40.5t-13 34t-25 18t-41.5 8h-88v739q0 25-14.5 42.5T2378 2016h-123q-11 0-17-1.5t-10-7t-5.5-10t-5-16.5t-7.5-22q-3 7-7 18.5t-5.5 15.5t-5 10t-7.5 8t-10 3.5t-15 1.5h-117q-22 0-36.5-17t-14.5-43v-739h-88q-25 0-41.5-8t-25-18t-13-34t-5-40t-.5-52V727q0-128 77-186.5t203-58.5m111-49q-80 0-144-62.5T2002 223q0-84 64-146t144-62t144.5 62t64.5 146q0 85-64.5 147.5T2210 433m-1021 49h222q126 0 202.5 58.5T1690 727v378l-1.5 31l-5 28.5l-9 21.5l-14.5 17.5l-22.5 9.5l-31.5 4h-89v739q0 25-14.5 42.5T1468 2016h-122q-13 0-20.5-4t-10-7.5t-7-20t-8.5-25.5q-3 8-7.5 21.5t-6.5 17t-6.5 9.5t-11 7.5t-17.5 1.5h-118q-22 0-36-17t-14-43v-739h-89q-18 0-31.5-4t-22.5-9.5t-14.5-18t-9-21t-5-28.5t-1.5-31V727q0-128 76.5-186.5T1189 482m111-49q-80 0-144-62.5T1092 223q0-84 64-146t144-62t144.5 62t64.5 146q0 85-64.5 147.5T1300 433M280 482h221q126 0 203 58.5T781 727v338q0 35-.5 51.5t-5 40.5t-13 34t-25 18t-41.5 8h-88v739q0 25-15 42.5t-35 17.5H436q-13 0-20.5-4t-10-7.5t-7-20t-8.5-25.5q-3 8-7.5 21.5t-6.5 17t-6.5 9.5t-11 7.5t-17.5 1.5H224q-22 0-36.5-17t-14.5-43v-739H85q-25 0-41.5-9t-25-20t-13-36t-5-42t-.5-53V719q0-237 280-237m110-49q-80 0-144-62.5T182 223q0-84 64-146t144-62t145 62t65 146q0 85-65 147.5T390 433"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Person(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832 0Q673 0 560.5 112.5T448 384t112.5 271.5T832 768t271.5-112.5T1216 384t-112.5-271.5T832 0m0 896q112 0 227 22t224 69.5t193.5 114t136 162.5t51.5 208q0 75-57 133.5t-135 58.5H192q-78 0-135-58.5T0 1472q0-112 51.5-208t136-162.5t193.5-114T605 918t227-22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pig(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 1536q50 5 128 5t128-5q-2 42 0 139q1 15 0 31t1 31.5t6 27.5t15.5 19.5t30.5 7.5h94q35 0 55.5-12t42.5-43l236-339q103-96 168.5-250t65.5-270q0-41-4.5-79t-17.5-79.5t-18-58t-25-63t-20-47.5q13-41 47-138.5t61-176.5t30-98q-5-5-24.5-25.5t-26-26t-27-22t-34-22T1641 26t-51.5-15t-62.5-7.5t-80-3.5q-64 0-108.5 19T1275 61.5t-43.5 58.5t-44.5 53q-140-46-291-46t-291 46q-20-18-44.5-53T517 61.5T453.5 19T345 0q-43 0-80 3.5T202.5 11T151 26t-39.5 16.5t-34 22t-27 22t-26 26T0 138q3 19 30 98t61 176.5T138 551q0 1-20 47.5t-25 63t-18 58T57.5 799T53 878q0 116 65.5 270T287 1398l236 339q22 31 42.5 43t55.5 12h94q19 0 30.5-7.5T761 1765t6-27.5t1-31.5t0-31q2-97 0-139m-203-467q0-62 28.5-111.5t75-80.5t106-47.5T896 813t121.5 16.5t106 47.5t75 80.5T1227 1069t-28.5 111.5t-75 80.5t-106 47.5T896 1325t-121.5-16.5t-106-47.5t-75-80.5T565 1069m469-104q-29 0-49 26.5t-20 63.5t20 63.5t49 26.5t49-26.5t20-63.5t-20-63.5t-49-26.5m170-456q29 0 50 15t21 36v118q0 23-23 37t-48 14q-29 0-50-15t-21-36V560q0-24 22.5-37.5T1204 509M758 965q-29 0-49 26.5t-20 63.5t20 63.5t49 26.5t49-26.5t20-63.5t-20-63.5t-49-26.5M588 509q26 0 48.5 13.5T659 560v118q0 21-21 36t-50 15q-25 0-48-14t-23-37V560q0-21 21-36t50-15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pregnant(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M862 261q0 108-76.5 184T602 521q-108 0-184.5-76T341 261q0-107 76.5-184T602 0q107 0 183.5 77T862 261M676 558q19 0 33.5 2.5t27 11t20 14t17 21.5t14 24t15 31t15.5 33l28 61q11 23 48.5 51t80.5 58t85.5 68.5t70.5 96.5t28 128q0 83-32 154t-82 109v104q0 40-13.5 57t-59.5 17H824v318q0 55-36 93t-90 38H506q-53 0-91-38.5t-38-92.5v-318h-59q-30 0-52-22t-22-52q0-8 4.5-21t5.5-17l123-478V855h-52l-99 397q-6 25-34 37.5t-59 12.5q-46 0-79.5-33T20 1190q0-10 3-20.5t8.5-22.5t7.5-18l135-448q17-56 76.5-89.5T377 558zm31 472q-48 0-75 35t-27 88q4 43 41 80l165 166l163-166q38-38 42-80q0-53-28-88t-76-35t-73.5 26t-27.5 69q-2-43-28.5-69t-75.5-26"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Profile(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M704 128q-144 0-225 106t-81 271q-1 205 132 325q17 16 12 41l-23 48q-11 24-32.5 37.5T396 995q-3 1-126.5 41T138 1080q-84 35-110 73q-28 63-28 319h1408q0-256-28-319q-26-38-110-73q-8-4-131.5-44T1012 995q-69-25-90.5-38.5T889 919l-23-48q-5-25 12-41q133-120 132-325q0-165-81-271T704 128"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Qsquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0m221 896q0-142 99.5-238.5T899 560q139 0 237.5 98.5T1235 896q0 87-41 161l-13-16q-2-2-10.5-12.5t-13-15t-13-14t-15-14t-15-10t-17-8t-16.5-2.5H922q-27 0-51 15t-24 40q0 18 33 59l115 138q-48 15-102 15q-139 0-237.5-98.5T557 896m808 369q128-163 128-369q0-162-80-299.5T1195.5 379T896 299t-299.5 80T379 596.5T299 896t80 299.5T596.5 1413t299.5 80q142 0 272-65q54 65 91 65h159q27 0 51-15t24-40q0-9-4.5-19t-8.5-15.5t-13.5-16.5t-11.5-14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionSquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1456 1H336Q197 1 98.5 99T0 336v1120q0 139 98.5 237.5T336 1792h1120q139 0 237.5-98.5T1792 1456V336q0-139-98.5-237T1456 1m-411 1455q0 16-10.5 26.5T1008 1493H784q-16 0-26.5-10.5T747 1456v-224q0-16 10.5-27t26.5-11h224q16 0 26.5 11t10.5 27zm299-784q0 43-8.5 79t-27.5 65.5t-35 49.5t-46.5 42t-45 31t-47.5 28q-38 22-63.5 57.5T1045 1082q0 16-10.5 27t-26.5 11H784q-16 0-26.5-11t-10.5-27v-42q0-78 61.5-146T943 793q55-26 78.5-53t23.5-71q0-38-43.5-68.5T902 570q-60 0-100 28q-35 25-101 107q-11 15-29 15q-12 0-22-7L496 596q-26-21-10-49q151-249 433-249q101 0 198.5 48t162 135.5T1344 672"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rsquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0m339 507v285h322q50 0 85-41.5t35-100.5t-35-101t-85-42zm0 911v-418h322q50 0 85 41.5t35 100.5v276q0 31 22 53t53 22h98q30 0 52-22t22-53v-276q0-87-24.5-140T1262 896q102-102 102-246q0-145-104-248t-251-103H502q-30 0-52 22t-22 53v1044q0 31 22 53t52 22h98q31 0 53-22t22-53"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rose(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1404 163q-57 15-97 73.5T1267 364q0 59 4.5 183t4.5 181q0 72-13 124t-33.5 84t-55.5 53t-69.5 30.5t-85.5 17.5q43-75 43-162t-35-179.5T944.5 527T845 398q99-119 226-178.5t333-56.5m-572 925q155-37 156-200q1-69-26-148t-76.5-158T760 430.5t-167-128T384.5 215T142 185q56 15 96.5 84.5T279 410q0 60-4.5 172T270 750q0 67 25 123.5t67.5 95.5t98.5 66t117.5 40t125.5 13v256q-31-38-103.5-77T445 1201q-118-40-250.5-36.5T0 1216q143 186 339 240t365-48v350q0 14 8 24t22 10h68q14 0 22-10t8-24v-158q90 41 188.5 43.5t191-30.5t177.5-102t147-167q-48-29-124.5-46t-170.5-12.5t-169 36.5q-63 27-130 86.5T832 1523zm332-991q0 23-49.5 50T995 200t-77 30q3-24-20.5-44T844 156.5t-56.5-18T760 126q-5-12 51.5-15.5t113.5-6T979 92q-2-6-12.5-10.5t-28-7.5t-37-5t-42-3t-39-1.5l-33-1L768 63q-39-2-125 6t-82 25q3 14 51 30t103.5 29.5T817 192t46 57q0 11-2 18.5t-10 17t-14 16t-24 22t-30 27.5q-45-47-101-79t-147-48q-68-13-112-44t-44-69q0-26 22.5-45.5t62.5-31t80-19t90.5-10t77-3.5T768 0q396 0 396 97"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ssquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0m94 656q0 100 38.5 167t100 98.5t136 48.5t149 24.5t136 19t100 41t38.5 81.5q0 61-68 103.5T896 1282q-71 0-129-24.5t-85-64.5q-25-38-40-47.5t-42-9.5H477q-57 0-39 68q33 124 161.5 206.5T896 1493q126 0 233.5-48t170-130t62.5-179q0-100-38.5-167t-100-98.5t-136-48.5t-149-24.5t-136-19t-100-41T664 656q0-61 68-103.5T896 510q71 0 129 24.5t85 64.5q25 38 40 47.5t42 9.5h123q57 0 39-68q-33-124-161.5-206.5T896 299q-126 0-233.5 48t-170 130T430 656"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Senior(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M642 0q-78 0-132.5 55T455 188t54.5 132.5T642 375q79 0 133.5-54.5T830 188T775.5 55T642 0m-92 428q-24-52-88-56q-24-2-94-2q-159 5-286 205q-61 97-74.5 201T38 980l62 269l-75 406q-9 52 21.5 94t80.5 43q119 0 141-89l89-389q8-38 8-53t-9-56l-42-248l110-292q9 13 32.5 50t39.5 60t40 52.5t48.5 49.5t49.5 31q23 10 328 74q41 8 68.5-24t27.5-66q0-43-57-74l-288-82zm241 767q0-71 50.5-121t122.5-50q70 0 120.5 50t50.5 121l-2 554q0 18-12.5 30.5T1090 1792q-17 0-30-12.5t-13-30.5v-554q0-36-24-61t-59-25q-36 0-60.5 25t-24.5 61q0 18-13 30t-30 12q-18 0-31.5-12t-13.5-30"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SevenSquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1456 1q139 0 237.5 98t98.5 237v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 99T336 1zM971 1419l-1-98q0-59 8-114t20-99t35-92.5t43.5-83.5t54.5-82.5t58.5-79t66-84T1323 600q3-3 24.5-28t37.5-45t30.5-50t14.5-54v-50q0-30-22-52t-53-22H437q-31 0-53 22t-22 52v111q0 31 22 53t53 22h582Q898 671 771.5 887.5T645 1249v170q0 31 22 52.5t53 21.5h177q30 0 52.5-22t21.5-52"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SexFemale(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 1664v-256h384v-256H640v-144q168-43 276-181t108-315q0-212-150-362T512 0T150 150T0 512q0 177 108 315t276 181v144H0v256h384v256zM512 256q106 0 181 75t75 181t-75 181t-181 75t-181-75t-75-181t75-181t181-75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SexMale(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 512q-212 0-362 150T0 1024t150 362t362 150t362-150t150-362q0-140-71-260l327-327v395h256V0H704v257h394L772 583q-120-71-260-71m0 256q106 0 181 75t75 181t-75 181t-181 75t-181-75t-75-181t75-181t181-75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sheep(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M568 787q9 2 30 2q8 0 16.5-2t15-4t15-7.5t13-8.5t14-11t12-10t12.5-11.5t11-10.5q20 44 59.5 77.5T849 835q55 0 98-41.5t55-97.5q72 76 133 76q37 0 74-21t64-52q13 62 63.5 107.5T1448 852q20 80 20 164q0 154-40 280t-107 205.5t-150 122t-173 42.5q-90-1-173-43.5T675 1501t-107.5-205t-40.5-280q0-111 41-229m168 174q-34 0-62 27.5t-28 62.5q0 34 24.5 58.5T729 1134q36 0 63.5-28t27.5-62v-4q-29 18-36 18q-20 0-34-14.5t-14-34.5q0-18 28-43q-15-5-28-5m532 0q-34 0-63 27.5t-29 62.5q0 34 25.5 58.5t59.5 24.5q35 0 62-28t27-62v-4q-29 18-36 18q-20 0-33-14t-13-35q0-20 26-43q-15-5-26-5M671 239q29-71 81.5-121T867 44.5T998 21t131 23.5t114.5 73.5t81.5 121q60-24 121.5-25.5t115 17.5t101 52t80 80.5t51 101t16 115.5t-27.5 122q106 42 162 135.5t51.5 199T1931 1227q15 66 4 122q-8 51-35 80.5t-72 22.5q-3-1-9.5-2.5t-9.5-2.5q8 73-15.5 143t-70 121.5t-111 86.5t-138 36.5T1324 1807q-36 82-103 136t-145 68.5t-156 0t-144.5-68.5T674 1807q-77 29-150.5 27.5T386 1798t-111-86.5t-70.5-121.5t-15.5-143q-3 1-8.5 2.5t-8.5 2.5q-46 7-74.5-22.5T59 1349q-10-53 5-122q-61-85-64.5-191t54-198.5T216 702q-25-60-27.5-122T204 464.5t50.5-101T334 283t100.5-52t115-17.5T671 239m854 578q-1-3 27.5 10.5T1627 870t96 68t100 94.5t79 113.5q29-56 29.5-117.5t-25-117t-83-101.5t-137.5-69q41-73 49.5-146T1725 464.5t-62-101t-100.5-63.5t-131-12.5T1284 335q-22-80-68-137t-101-83.5T998 88t-117 26.5T780 198t-68 137q-74-40-147-47.5T434.5 300T334 363.5t-62 101T262 595t50 146q-81 23-138.5 69T90 911.5t-25.5 117T94 1146q29-58 79-113.5T273 938t95.5-68t74-43t26.5-10q-36 77-37.5 174.5T458 1161q-12 20-38.5 86.5t-60 107T266 1422q-16 74 .5 141t60 112t102.5 69.5t134 16.5t149-49q23 79 69 135t100.5 82t116.5 26t116.5-26t100.5-82t69-135q74 41 149 49t134.5-16.5t103-69.5t60-112t.5-141q-47-21-80-55t-47-63.5t-33-74.5t-32-68q27-72 24.5-169T1525 817"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldCheck(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1000 20l1000 300q0 136-2.5 228t-13 217.5t-30 216t-55 205T1813 1391t-126.5 195t-172.5 196t-226.5 188.5T1000 2160q-158-92-287.5-189.5T486 1782t-172.5-196T187 1391t-86.5-204.5t-55-205t-30-216T2.5 548T0 320zm673 732q0-36-24-60l-118-118q-25-25-59-25q-33 0-58 25l-533 531l-294-295q-24-24-59-24t-59 24L351 928q-25 25-25 61q0 32 25 57l472 472q25 25 58 25q35 0 60-25l708-708q24-24 24-58"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldTimes(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1000 20l1000 300q0 136-2.5 228t-13 217.5t-30 216t-55 205T1813 1391t-126.5 195t-172.5 196t-226.5 188.5T1000 2160q-158-92-287.5-189.5T486 1782t-172.5-196T187 1391t-86.5-204.5t-55-205t-30-216T2.5 548T0 320zm497 1237q0-34-25-59l-236-236l236-235q25-25 25-60q0-36-25-59l-117-118q-23-24-59-24t-60 24l-236 237l-236-237q-24-24-59-24q-37 0-59 24L528 608q-26 24-26 59q0 34 26 60l236 235l-236 236q-26 26-26 59q0 36 26 60l118 118q23 25 59 25q34 0 59-25l236-237l236 237q25 25 60 25q36 0 59-25l117-118q25-23 25-60"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shop(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1664 675V331q0-40-50-89.5T1503.5 160t-95.5-32q0-24-20-44t-44-20H320q-24 0-44 20t-20 44q-35 0-95.5 32T50 241.5T0 331v337q0 39 34.5 72t93.5 42v562q-26 0-45 19t-19 45t19 45t45 19h1408q26 0 45-19t19-45t-19-45t-45-19V785q60-10 94-40.5t34-69.5M896 896v448H256V782q79-24 96-39q57 50 160 50t160-50q57 50 160 50t160-50q57 50 160 50t160-50q17 15 96 39v562h-128V896zm64 64h256v384H960zM384 677l-1-377q0-58 33.5-82t94.5-26h129q0 54 1 223.5t0 261.5q-37 50-129 50q-90 0-128-50m639 0V192h128q24 1 31 1.5t26 3t25 6t18 11.5t16 19t8 28t4 39l1 377q-6 27-41.5 41t-86.5 14t-87-14t-42-41M384 896v320h384V896zm64 64h256v192H448z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SixSquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1456 1H336Q197 1 98.5 99T0 336v1120q0 139 98.5 237.5T336 1792h1120q139 0 237.5-98.5T1792 1456V336q0-139-98.5-237T1456 1M896 299q89 0 166.5 19.5t128 50.5t82.5 62.5t44 61.5q21 47 12.5 81t-49.5 34h-161q-34-1-46.5-12.5T1056 562q-8-37-39.5-58.5t-66.5-27t-76-3.5q-50 2-91 35.5T718.5 591T683 693t-10 100q8-3 51-19.5t59-21.5t48.5-12t61.5-7q203 0 344 108t141 267q0 105-65 193.5t-176.5 140T893 1493q-105 0-190.5-32.5t-144-88.5t-98.5-132.5t-58-162T384 898q0-119 33-225.5T512 482t161.5-133.5T896 299m4 628q-86 0-146.5 56.5T693 1120t60.5 136.5T900 1313t146.5-56.5T1107 1120t-60.5-136.5T900 927"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sleep(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1152 992v64q0 14-9 23t-23 9H672q-14 0-23-9t-9-23v-64q0-32 20-56l300-360H672q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h448q14 0 23 9t9 23v64q0 32-20 56L832 960h288q14 0 23 9t9 23m640 384v64q0 14-9 23t-23 9h-448q-14 0-23-9t-9-23v-64q0-32 20-56l300-360h-288q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h448q14 0 23 9t9 23v64q0 32-20 56l-300 360h288q14 0 23 9t9 23M512 608v64q0 14-9 23t-23 9H32q-14 0-23-9t-9-23v-64q0-32 20-56l300-360H32q-14 0-23-9t-9-23V96q0-14 9-23t23-9h448q14 0 23 9t9 23v64q0 32-20 56L192 576h288q14 0 23 9t9 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SleepSquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M288 0Q170 0 85 85T0 288v960q0 118 85 203t203 85h960q118 0 203-85t85-203V288q0-118-85-203T1248 0zm663 928v46q0 10-6.5 16.5T928 997H608q-10 0-16.5-6.5T585 974v-46q0-23 14-40q36-43 107.5-128.5T814 631H608q-10 0-16.5-6.5T585 608v-46q0-10 6.5-16.5T608 539h320q10 0 16.5 6.5T951 562v46q0 23-14 40q-36 43-107.5 128.5T722 905h206q10 0 16.5 6.5T951 928m457 274v46q0 10-6.5 16.5t-16.5 6.5h-320q-10 0-16.5-6.5t-6.5-16.5v-46q0-22 14-40q43-50 124-147.5t91-109.5h-206q-10 0-16.5-6.5T1042 882v-45q0-10 6.5-16.5t16.5-6.5h320q10 0 16.5 6.5t6.5 16.5v45q0 23-14 40q-10 12-91 109.5T1179 1179h206q10 0 16.5 6.5t6.5 16.5M494 654v45q0 10-6.5 16.5T471 722H151q-10 0-16.5-6.5T128 699v-45q0-23 14-40q10-12 91-109.5T357 357H151q-10 0-16.5-6.5T128 334v-46q0-10 6.5-16.5T151 265h320q10 0 16.5 6.5T494 288v46q0 22-14 40q-43 50-124 147.5T265 631h206q10 0 16.5 6.5T494 654"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Smoking(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 786q1 95 41.5 193.5T154 1152q-16-79-13.5-157T159 868q13-41 38.5-78t50-64.5t48.5-58t39.5-73T353 499q2-55-9.5-114t-39-120t-68-112.5T135 60T0 0q47 49 77.5 104.5t43.5 113t17 97.5t4 93q1 54-21 109.5t-48.5 94t-48.5 86T2 786m557 668l-147-292l1234-619l146 291zm-408 205L4 1368l290-146l146 291z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmokingAlt(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1504 640q0-23-5-40t-17-27.5t-22.5-17t-33-9.5t-36.5-4.5t-43-2.5q-47-2-72-4t-72-13t-75.5-29.5t-66-56t-61-89.5T957 214.5T928 32q0-14 9-23t23-9t23 9t9 23q7 86 23 154t35.5 114t47.5 79.5t54.5 52t62 29.5t62.5 14t64 4q10 0 35.5.5t35 1t31.5 2.5t31 5t25.5 9t24 14t18 20t15 27t8 36t3.5 46v96q0 14-9 23t-23 9t-23-9t-9-23zm0-320q-2-160-160-160h-96q-14 0-23-9t-9-23t9-23t23-9h96q101 0 158.5 49.5T1560 292q53 0 85 9.5t61 40.5q56 60 54 170v224q0 14-9 23t-23 9t-23-9t-9-23V512q0-72-31-116t-97-44h-64zM0 832h1408v320H0zm1503 0q-12 0-21.5 9t-9.5 22v258q0 13 9.5 22t21.5 9h66q12 0 21.5-9t9.5-22V863q0-13-9.5-22t-21.5-9zm192 0q-12 0-21.5 9t-9.5 22v258q0 13 9.5 22t21.5 9h66q12 0 21.5-9t9.5-22V863q0-13-9.5-22t-21.5-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sms(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 704q0 174-120 321.5t-326 233t-450 85.5q-70 0-145-8q-198 175-460 242q-49 14-114 22q-17 2-30.5-9t-17.5-29v-1q-3-4-.5-12t2-10t4.5-9.5l6-9l7-8.5l8-9q7-8 31-34.5t34.5-38t31-39.5t32.5-51t27-59t26-76q-157-89-247.5-220T0 704q0-130 71-248.5T262 251t286-136.5T896 64q244 0 450 85.5t326 233T1792 704M128 802q22 189 218 189q91 0 151.5-42T558 820q0-36-10.5-64T521 709.5t-43-33t-52-23t-62-17.5q-80-18-102-41q-12-13-12-29q0-29 28-42t61-13q44 0 66.5 18.5T435 591l107-4q-3-86-58-127.5T340 418q-80 0-136 39t-56 116q0 71 45 109t123 58q4 1 25 5.5t31 7.5t28 9.5t27 13.5t16.5 19t7.5 26q0 39-31.5 57.5T347 897q-96 0-115-105zm519 180h98l1-436l103 436h103l104-436l1 436h99V429H995l-93 376l-95-376H647zm587-180q22 189 219 189q90 0 150.5-42t60.5-129q0-45-14-77t-44.5-52.5T1546 659t-76-23q-80-18-102-41q-12-13-12-29q0-29 28-42t61-13q44 0 67 18.5t30 61.5l107-4q-3-86-58-127.5T1447 418q-80 0-136.5 39T1254 573q0 71 45.5 109t123.5 58q5 1 25 5.5t30 7.5t27.5 9.5t27 13.5t17 19t7.5 26q0 39-31 57.5t-72 18.5q-98 0-115-105z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sofa(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 1428h-54.5l-30.5-.5l-34.5-1.5l-32-2.5l-33-4.5l-28-7l-26-9l-17.5-12v145q0 29 18.5 46.5T320 1600h128q27 0 45.5-17.5T512 1536zm1024 0h54.5l30.5-.5l34.5-1.5l32-2.5l33-4.5l28-7l26-9l17.5-12v145q0 29-18.5 46.5T1728 1600h-128q-27 0-45.5-17.5T1536 1536zM193 575q-98 0-146 43T0 736q1 86 36.5 131t91.5 45q0 48 1 79t4 72.5t9.5 68.5t17.5 59.5t27.5 52.5t40.5 41t55 32.5t72.5 19t92.5 7.5h1152q51 0 92.5-7.5t72.5-19t55-32.5t40.5-41t27.5-52.5t17.5-59.5t9.5-68.5t4-72.5t1-79q56 0 91.5-45t36.5-131q1-75-46-118t-145-43q-49 0-92 35t-74.5 93.5t-53.5 123t-36 134.5q-266-133-576.5-133T448 961q-14-71-36-135t-53-122.5t-74-93.5t-92-35m319-383q-64 0-111 7.5T320.5 225t-55 40t-36 58t-21.5 72t-13 90q38 5 70 16t55 23.5t44.5 36.5t34 40.5t28.5 51t23.5 54t23 62.5t23.5 64q244-92 527.5-92t527.5 92q8-20 19.5-52t18.5-50.5t18-45.5t20-43.5t23-38.5t27.5-35.5t32-29t39.5-25t47.5-17T1855 485q-7-64-17-106t-31-81t-54.5-60t-87-33.5T1536 192z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Speech(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1664 832q0 203-66.5 386.5T1391 1517q-24 19-47 19q-26 0-45-19t-19-45q0-32 31-55q80-58 133-163.5t72.5-209.5t19.5-212t-19.5-212t-72.5-209.5T1311 247q-31-23-31-55q0-26 19-45t45-19q23 0 47 19q140 115 206.5 298.5T1664 832m-256 0q0-161-51.5-309.5T1198 275q-21-19-46-19q-26 0-45 19t-19 45q0 31 29 54q82 67 122.5 198.5T1280 832t-40.5 259.5T1117 1290q-29 23-29 54q0 26 19 45t45 19q25 0 46-19q107-99 158.5-247.5T1408 832m-448 448q-26 0-45-19t-19-45q0-24 17-43q111-130 111-341q0-210-111-341q-17-19-17-43q0-26 19-45t45-19q27 0 48 21q72 75 108 192t36 235t-36 235t-108 192q-21 21-48 21M64 128q-26 0-45-19T0 64t19-45T64 0q121 0 226 41t182 126.5T575 371q18 66 49 100q23 26 96 62t97 65q7 9 11 21.5t4.5 19t0 25t-.5 21.5q0 26-13 48.5t-37.5 40t-47 29.5t-55.5 26l-96 37q5 20 22.5 57.5t28 71t5.5 59.5q-5 24-43 31q-49 16-154 17.5t-154-6.5q33 39 55 57q34 27 66 39t103 24q69 11 64 70q-9 95-82.5 140.5T320 1472q-19 0-32.5 8.5T268 1503t-9.5 31t-3 34t.5 32q0 26-19 45t-45 19t-45-19t-19-45v-64q0-82 56-137t136-55q34-1 74.5-11t41.5-31q-128-12-218-84t-90-194q0-33 26-50.5t59-10.5q92 21 215 19q44 0 66-7q-1-2-23.5-61T448 832q0-44 48-62q5-2 35-11t58-18t58-21.5t47.5-26.5t14.5-27q-4-10-45-29t-88.5-46t-65.5-53q-49-70-61-143q-15-72-51.5-123t-92-79.5T194 150T64 128"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spinner(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 256q212 0 362 150t150 362q0 104-39.5 197.5T1130 1130l181 181q106-106 165.5-246.5T1536 768q0-209-103-385.5T1153.5 103T768 0T382.5 103T103 382.5T0 768q0 156 59.5 296.5T225 1311l181-181q-71-71-110.5-164.5T256 768q0-212 150-362t362-150"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StickyNote(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1152 1216h412q-13 60-94 148.5t-173.5 156T1152 1600zm-128-64v448H0q79-242 103.5-447.5T128 576V64h1536v512q0 65-3.5 125t-6.5 97t-12 90t-12.5 69t-16 70.5t-13.5 60.5h-512q-40 0-52 12t-12 52"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stroller(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1216 0L960 565h768q-24-207-167-360.5T1216 0M590 651q0-74-19-154t-58-157.5t-92.5-139t-130-99T128 64q-55 0-91.5 34.5T0 192q0 55 39 91.5t89 36.5q43 0 66.5-25t42.5-71q99 69 145 175.5T428 628q7 170 96 312t237 224.5t322 82.5t319.5-76.5t232-213.5t93.5-306zm-14 472q-85 0-145.5 60.5T370 1330t60 146t146 60t146-60t60-146t-60.5-146.5T576 1123m0 91q47 0 81 34t34 82t-34 81.5t-81 33.5t-81-33.5t-34-81.5t34-82t81-34m1010-90q-86 0-146.5 60.5T1379 1330t60.5 145.5T1586 1536t146-60t60-146t-60-146t-146-60m0 90q48 0 81.5 34t33.5 82t-33.5 81.5t-81.5 33.5t-82-33.5t-34-81.5t34-82t82-34"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sun(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M873 416q-130 0-240.5 64.5t-175 175T393 896t64.5 240.5t175 175T873 1376t240.5-64.5t175-175T1353 896t-64.5-240.5t-175-175T873 416m853 757q0 6-7 12t-13 8l-293 97v306q0 16-13 26q-14 9-29 4l-292-94l-180 248q-9 12-26 12t-26-12l-180-248l-292 94q-15 5-29-4q-12-8-14-26v-306l-292-97q-16-5-20-20q-5-16 4-29l180-248L24 648q-9-13-4-29q4-15 20-20l292-97V196q2-18 14-26q14-9 29-4l292 94L847 12q9-12 26-12t26 12l180 248l292-94q15-5 29 4q13 10 13 26v306l293 97q6 2 13 8t7 12q4 15-4 29l-181 248l181 248q8 14 4 29m-277-277q0-157-77-289.5T1162.5 397T873 320t-289.5 77T374 606.5T297 896q0 118 46 225t123 184t183.5 122.5T873 1473q158 0 290.5-77.5t209-210T1449 896"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sunrise(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1565 1216h-93q0-157-77-289.5T1185.5 717T896 640t-289.5 77T397 926.5T320 1216h-93L47 968q-9-13-4-29q4-15 20-20l292-97V516q2-18 14-26q14-9 29-4l292 94l180-248q9-12 26-12t26 12l180 248l292-94q15-5 29 4q13 10 13 26v306l293 97q6 2 13 8t7 12q4 15-4 29l-181 248zM55 1312q-21 0-38 16.5T0 1367v50q0 22 17 38.5t38 16.5h1682q21 0 38-16.5t17-38.5v-50q0-22-17-38.5t-38-16.5zm841-576q-130 0-240.5 64.5t-175 175T416 1216h960q0-130-64.5-240.5t-175-175T896 736"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunriseO(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1565 1216h-93q0-157-77-289.5T1185.5 717T896 640t-289.5 77T397 926.5T320 1216h-93L47 968q-9-13-4-29q4-15 20-20l292-97V516q2-18 14-26q14-9 29-4l292 94l180-248q9-12 26-12t26 12l180 248l292-94q15-5 29 4q13 10 13 26v306l293 97q6 2 13 8t7 12q4 15-4 29l-181 248zM55 1312q-21 0-38 16.5T0 1367v50q0 22 17 38.5t38 16.5h1682q21 0 38-16.5t17-38.5v-50q0-22-17-38.5t-38-16.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tsquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0m1082 299H374q-31 0-53 22t-22 53v149q0 31 22 52.5t53 21.5h373v821q0 31 22 53t53 22h149q31 0 52.5-22t21.5-53V597h373q31 0 53-21.5t22-52.5V374q0-31-22-53t-53-22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Table(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 606L0 576V448L832 64l832 384v128l-64 30v534q0 38-26.5 57t-69.5 19t-69.5-19t-26.5-57V694L928 916v480q0 38-26.5 57t-69.5 19t-69.5-19t-26.5-57V916L256 694v446q0 38-26.5 57t-69.5 19t-69.5-19t-26.5-57z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableAlt(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M289 192h1086q202 185 267 245q22 11 22 34v141q0 22-11.5 50t-27.5 28h-32v678q0 27-8 65.5t-23 38.5h-133q-16 0-35-42t-19-62V690H289v678q0 20-19 62t-35 42H102q-15 0-23-38.5t-8-65.5V690H39q-16 0-27.5-28T0 612V471q0-23 22-34zm1034 866V746h-170v312q0 24 9 51.5t25 27.5h71q17 0 41-30.5t24-48.5m-982 0V746h170v312q0 24-9 51.5t-25 27.5h-71q-17 0-41-30.5t-24-48.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableO(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 606L0 576V448L832 64l832 384v128l-64 30v534q0 38-26.5 57t-69.5 19t-69.5-19t-26.5-57V694L928 916v480q0 38-26.5 57t-69.5 19t-69.5-19t-26.5-57V916L256 694v446q0 38-26.5 57t-69.5 19t-69.5-19t-26.5-57zm768 226l768-352l-768-352L64 480z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableQuestion(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 606L0 576V448L832 64l832 384v128l-64 30v534q0 38-26.5 57t-69.5 19t-69.5-19t-26.5-57V694L928 916v480q0 38-26.5 57t-69.5 19t-69.5-19t-26.5-57V916L256 694v446q0 38-26.5 57t-69.5 19t-69.5-19t-26.5-57zm2176 458v240q0 16-12 28t-28 12h-240q-16 0-28-12t-12-28v-240q0-16 12-28t28-12h240q16 0 28 12t12 28m316-600q0 54-15.5 101t-35 76.5t-55 59.5t-57.5 43.5t-61 35.5q-41 23-68.5 65t-27.5 67q0 17-12 32.5t-28 15.5h-240q-15 0-25.5-18.5T1920 904v-45q0-83 65-156.5T2128 594q59-27 84-56t25-76q0-42-46.5-74T2083 356q-65 0-108 29q-35 25-107 115q-13 16-31 16q-12 0-25-8l-164-125q-13-10-15.5-25t5.5-28q160-266 464-266q80 0 161 31t146 83t106 127.5t41 158.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tables(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 0Q76 0 38 38T0 128v512q0 52 38 90t90 38h512q52 0 90-38t38-90V128q0-52-38-90T640 0zm0 128h512v512H128zM1024 0q-52 0-90 38t-38 90v512q0 52 38 90t90 38h512q52 0 90-38t38-90V128q0-52-38-90t-90-38zM128 896q-52 0-90 38t-38 90v512q0 52 38 90t90 38h512q52 0 90-38t38-90v-512q0-52-38-90t-90-38zm0 128h512v512H128zm896-128q-52 0-90 38t-38 90v512q0 52 38 90t90 38h512q52 0 90-38t38-90v-512q0-52-38-90t-90-38zm0 128h512v512h-512z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tablesolution(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 1248l320-320q5 105 48 202t120 174t174 120t202 48l-320 320q-188-74-328-216Q74 1436 0 1248m1792 0l-320-320q-5 105-48 202t-120 174t-174 120t-202 48l320 320q188-74 328-216q142-140 216-328M1248 0L980 268l-52 52q105 5 202 48t174 120t120 174t48 202l320-320q-74-188-216-328Q1436 74 1248 0M544 0l320 320q-105 5-202 48T488 488T368 662t-48 202l-52-52L0 544q74-188 216-328Q356 74 544 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThreeSquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1456 0H336Q197 0 98.5 99T0 337v1119q0 139 98.5 237.5T336 1792h1120q139 0 237.5-98.5T1792 1456V337q0-139-98.5-238T1456 0M906 514q0-1-12.5-1.5t-32.5 1t-43.5 5t-47.5 13t-44 23t-32.5 36T681 642q0 12-8.5 18t-15.5 6H434q-11 0-17.5-16t-6.5-28q3-71 32.5-127t75-92t107-60t126-34T885 299q77 0 157.5 16.5T1202 366t129.5 97.5T1382 611q0 182-195 285q195 103 195 285q0 84-50.5 147.5T1202 1426t-159.5 50.5T885 1493q-70 0-134.5-10t-126-34t-107-60t-75-92t-32.5-127q0-12 6.5-27.5T434 1127h223q8 0 16 6t8 18q0 29 12.5 51.5T726 1238t44 22.5t47.5 13t43.5 5t32.5 1.5t12.5-1q51-2 96-17t79-48.5t34-81.5q0-130-184-130H823q-34 0-42.5-8.5T771 953V839q1-32 10-40.5t42-8.5h108q184 0 184-130q-1-70-64-106.5T906 514"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Timeslot(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M320 320q-120 0-220 103.5T0 667v324q0 140 100 246.5T320 1344h1024q119 0 219.5-103.5T1664 998V667q0-139-100.5-243T1344 320zm-64 320h1152v384H256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeslotQuestion(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M320 320q-120 0-220 103.5T0 667v324q0 140 100 246.5T320 1344h1024q119 0 219.5-103.5T1664 998V667q0-139-100.5-243T1344 320zm-64 320h1152v384H256zm1984 424v240q0 16-12 28t-28 12h-240q-16 0-28-12t-12-28v-240q0-16 12-28t28-12h240q16 0 28 12t12 28m316-600q0 54-15.5 101t-35 76.5t-55 59.5t-57.5 43.5t-61 35.5q-41 23-68.5 65t-27.5 67q0 17-12 32.5t-28 15.5h-240q-15 0-25.5-18.5T1920 904v-45q0-83 65-156.5T2128 594q59-27 84-56t25-76q0-42-46.5-74T2083 356q-65 0-108 29q-35 25-107 115q-13 16-31 16q-12 0-25-8l-164-125q-13-10-15.5-25t5.5-28q160-266 464-266q80 0 161 31t146 83t106 127.5t41 158.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Timeslots(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M815 61q-44 0-74.5 38T710 191v140q0 54 31 92.5t74 38.5h417q43 0 73.5-38.5T1336 331V191q0-54-30.5-92T1232 61zM105 1071q-44 0-74.5 38.5T0 1202v142q0 54 30.5 91.5T105 1473h417q44 0 74.5-38t30.5-92v-141q0-54-31-92.5t-74-38.5zm0 131h417v141H105zm710-131q-43 0-74 38.5t-31 92.5v141q0 54 30.5 92t74.5 38h417q43 0 73.5-38t30.5-92v-141q0-54-30.5-92.5T1232 1071zM105 61q-44 0-74.5 38T0 191v140q0 54 30.5 92.5T105 462h417q43 0 74-38.5t31-92.5V191q0-54-30.5-92T522 61zm1423 0q-44 0-75 38t-31 92v140q0 54 31 92.5t75 38.5h416q44 0 74.5-38.5T2049 331V191q0-54-30.5-92T1944 61zm0 1010q-44 0-75 38.5t-31 92.5v141q0 54 31 92t75 38h416q44 0 74.5-38t30.5-92v-141q0-54-30.5-92.5T1944 1071zm0 131h416v141h-416zM105 548q-44 0-74.5 38T0 678v178q0 54 30.5 92.5T105 987h417q43 0 74-38.5t31-92.5V678q0-54-30.5-92T522 548zm0 130h417v178H105zm710-130q-44 0-74.5 38T710 678v178q0 54 31 92.5t74 38.5h417q43 0 73.5-38.5T1336 856V678q0-54-30.5-92t-73.5-38zm713 0q-44 0-75 38t-31 92v178q0 54 31 92.5t75 38.5h416q44 0 74.5-38.5T2049 856V678q0-54-30.5-92t-74.5-38z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwoSquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1455 0H336Q197 0 98.5 98.5T0 336v1119q0 139 98 237.5t238 98.5h1119q139 0 237.5-98.5T1791 1455V336q0-139-98.5-237.5T1455 0M346 1437v-11q0-70 59.5-155.5T557 1108t200.5-144T967 857q86-35 126.5-71t41.5-80q1-67-77-119t-157-52q-46 0-104 25t-107 66.5t-68 84.5q-3 6-6.5 14.5T611 736t-3.5 6t-3.5 4t-5 1t-7 0t-10.5-1H404q-30 0-48-26t-10-51q55-168 212-269t354-101q87 0 173.5 26t156.5 73t113.5 119.5T1399 674q0 81-24 145t-65.5 106.5t-92 75.5t-109 60t-112 52.5t-106.5 61t-86 76.5h595q19 0 33.5 16.5t14.5 39.5v130q0 23-14.5 39.5T1399 1493H394q-19 0-33.5-16.5T346 1437"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Usquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0m562 1221q-103 0-155.5-40T690 1051V374q0-31-22-53t-53-22H466q-31 0-53 22t-22 53v641q0 58 16.5 120t55.5 127t95 115.5t144.5 83T898 1493t195.5-32.5t144.5-83t94.5-115.5t55-127t16.5-120V374q0-31-21.5-53t-52.5-22h-149q-31 0-53 22t-22 53v677q0 90-52.5 130T898 1221"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserBoss(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M704 768q-5 3-55.5 25T590 821q38 57 73 98q-9 20-26 93t-36.5 168t-20.5 99q-1-3-21-8t-42.5-18.5T487 1215q-16-54-27-162t-17-192.5t-13-93.5q-7 3-31.5 16T348 806q-30 11-82 30t-82 30.5t-66.5 27.5t-59 32.5T24 960q-24 64-24 512h1408q0-112-5.5-292T1384 960q-12-17-34.5-33.5t-59-32.5t-66.5-27.5t-82-30.5t-82-30q-26-10-50.5-23T978 767q-7 9-11.5 93.5t-14 192.5t-25.5 162q-7 24-29.5 37.5T855 1271t-21 8q-62-303-89-360q35-41 73-98q-8-6-58.5-28T704 768m0-704q140 0 235 103q27-13 61-3q13 2 29 1t30.5-2.5t28.5 14t23 50.5q4 16 6.5 29.5t1.5 25.5t-1 20t-4.5 18t-5.5 14.5t-9 14t-9.5 11.5t-11 12t-10.5 11q-5 25-21.5 44t-28.5 19q-22 111-110 184.5T704 704t-204-73.5T390 446q-12 0-28.5-19T340 383q-1-1-10.5-11t-11-12t-9.5-11.5t-9-14t-5.5-14.5t-4.5-18t-1-20t1.5-25.5T297 227q9-35 23-50.5t28.5-14T379 165t29-1q34-10 61 3Q564 64 704 64m-47 224q6 0 10.5 4.5T672 303v27q20-2 32-2t32 2v-27q0-6 4.5-10.5T751 288h145q-19-34-17-47q1-10 3.5-14t5.5-5t10-4.5t13-9.5q-81-96-207-96t-207 96q6 6 13 9.5t10 4.5t5.5 5t3.5 14q2 13-17 47zm-177 41q-10 5-45 18q-3 18-3 37q0 112 80 192t192 80t192-80t80-192q0-19-3-37q-35-13-45-18v104q0 6-4.5 10.5T913 448H751q-6 0-10.5-4.5T736 433v-63q-15-3-32-3t-32 3v63q0 6-4.5 10.5T657 448H495q-6 0-10.5-4.5T480 433zm396-9h-88q-8 0-14 6t-6 14v56q0 8 6 14t14 6h88q8 0 14-6t6-14v-56q0-8-6-14t-14-6m-344 0q-8 0-14 6t-6 14v56q0 8 6 14t14 6h88q8 0 14-6t6-14v-56q0-8-6-14t-14-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserGroup(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1536 256q0 53-20 99q-21 47-55 82q-38 37-82 55q-46 20-99 20q-52 0-100-20q-43-19-81-55q-37-38-55-82q-20-46-20-99t20-99q18-44 55-82q38-36 81-55q48-20 100-20q53 0 99 20q44 18 82 55q34 35 55 82q20 46 20 99m128 609q0 40-16 72q-16 30-43 49q-30 20-63 29q-34 9-72 9h-134q-52-61-119-93q-65-32-146-35q39-57 60-122q21-64 21-134q0-6-.5-17t-.5-16q-2-22-4-33q28 10 65 17q35 6 68 6q52-2 95-13q49-15 78-29q30-14 56-30q20-11 31-13q34 2 53 21q21 18 35 53q14 29 22 72q8 55 10 79q3 40 3 74q1 22 1 54m-576-225q0 76-30 149q-30 70-83 123q-55 54-122 82q-69 30-149 30t-150-30q-65-28-122-82q-53-54-82-123q-30-69-30-149t30-149q31-72 82-122q53-53 122-83q70-30 150-30t149 30q70 30 122 83q52 49 83 122q30 73 30 149m320 893q-2 83-34 139q-32 57-93 89q-62 31-140 31H267q-79 0-141-31q-60-32-92-89q-34-56-34-139q0-49 5-121q7-70 20-128q15-70 38-122q26-60 63-104q40-47 92-71q52-25 128-27q12 2 31 15q28 18 52 34q25 16 72 41q41 21 92 34q45 13 111 15q66-2 111-15q51-13 92-34q47-25 72-41q10-7 52-34q19-13 31-15q76 2 127 27q52 24 93 71q38 48 63 104q25 59 38 122q13 58 20 128q5 72 5 121"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserSuit(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M704 768q-5 3-55.5 25T590 821q38 57 73 98q-9 20-26 93t-36.5 168t-20.5 99q-1-3-21-8t-42.5-18.5T487 1215q-16-54-27-162t-17-192.5t-13-93.5q-7 3-31.5 16T348 806q-30 11-82 30t-82 30.5t-66.5 27.5t-59 32.5T24 960q-24 64-24 512h1408q0-112-5.5-292T1384 960q-12-17-34.5-33.5t-59-32.5t-66.5-27.5t-82-30.5t-82-30q-26-10-50.5-23T978 767q-7 9-11.5 93.5t-14 192.5t-25.5 162q-7 24-29.5 37.5T855 1271t-21 8q-62-303-89-360q35-41 73-98q-8-6-58.5-28T704 768m0-704q133 0 226.5 93.5T1024 384t-93.5 226.5T704 704t-226.5-93.5T384 384t93.5-226.5T704 64m272 322q-39 8-111.5-7.5T742 342q-34-15-65-34.5t-71-49t-62-43.5q16 119-112 170q1 113 80.5 192T704 656t191.5-79T976 386"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserSuitFemale(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 1472h1152q0-24 32-250t32-298q0-38-92-68.5t-212-54T832 768l-53 178q-67 43-139 43t-139-43l-53-178q-24 10-144 33.5t-212 54T0 924q0 72 32 298t32 250m576-448q58 1 124-29l-124 413l-124-413q66 30 124 29m174-372q104 79 194 63q-48-28-57-85q16 21 46.5 27t59.5-4q-49-43-65-89t-24-126q-4-44-5.5-124.5T953 188t-32-79q-30-43-124-76T640 0q-87 0-183 43T336 160q-7 21-11.5 88T318 381.5t-2 69.5q-18 177-74 220q14 9 44 .5t52-25.5q-1 30-45 80q27 5 89.5-19t96.5-46q75 43 161 43q95 0 174-52m-375-85q-9-60 12-204t62-164q7-3 26 12.5t21 26.5q6 32-12 62q19 17 73.5 21t99 .5T788 313q12 13 22.5 28t17.5 33.5t12 32t9.5 37t6 33t4.5 34t4 28.5q-38 54-97 85.5T640 656q-120 0-201-89"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserWaiter(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M704 64q133 0 226.5 93.5T1024 384t-93.5 226.5T704 704t-226.5-93.5T384 384t93.5-226.5T704 64m272 322q-39 8-111.5-7.5T742 342q-34-15-65-34.5t-71-49t-62-43.5q16 119-112 170q1 113 80.5 192T704 656t191.5-79T976 386m-142 893H580q-10 0-26-4t-38-20t-29-40q-16-54-27-162t-17-192.5t-13-93.5q-7 3-31.5 16T348 806q-30 11-82 30t-82 30.5t-66.5 27.5t-59 32.5T24 960q-24 64-24 512h1408q0-112-5.5-292T1384 960q-12-17-34.5-33.5t-59-32.5t-66.5-27.5t-82-30.5t-82-30q-26-10-50.5-23T978 767q-7 9-11.5 93.5t-14 192.5t-25.5 162q-7 24-29 40t-38 20t-26 4M704 800l192-64v192l-192-64l-192 64V736zm0 224q-12 0-22-9.5T672 992t10-22.5t22-9.5t22 9.5t10 22.5t-10 22.5t-22 9.5m0 128q-12 0-22-9.5t-10-22.5t10-22.5t22-9.5t22 9.5t10 22.5t-10 22.5t-22 9.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserWaiterFemale(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 1472h1152q0-24 32-250t32-298q0-13-7-24.5t-23-20.5t-32-16t-45-15t-51-13t-61.5-14t-64.5-15q-26-10-50.5-23T914 767q-7 9-13 93.5T884 1053t-27 162q-7 24-29 40t-38 20t-26 4H516q-10 0-26-4t-38-20t-29-40q-16-54-27-162t-17-192.5t-13-93.5q-7 3-31.5 16T284 806q-25 6-64.5 15T158 835t-51 13t-45 15t-32 16t-23 20.5T0 924q0 72 32 298t32 250m576-672l192-64v192l-192-64l-192 64V736zm0 224q-12 0-22-9.5T608 992t10-22.5t22-9.5t22 9.5t10 22.5t-10 22.5t-22 9.5m0 128q-12 0-22-9.5t-10-22.5t10-22.5t22-9.5t22 9.5t10 22.5t-10 22.5t-22 9.5m174-500q104 79 194 63q-48-28-57-85q16 21 46.5 27t59.5-4q-49-43-65-89t-24-126q-4-44-5.5-124.5T953 188t-32-79q-30-43-124-76T640 0q-87 0-183 43T336 160q-7 21-11.5 88T318 381.5t-2 69.5q-18 177-74 220q14 9 44 .5t52-25.5q-1 30-45 80q27 5 89.5-19t96.5-46q75 43 161 43q95 0 174-52m-375-85q-9-60 12-204t62-164q7-3 26 12.5t21 26.5q6 32-12 62q19 17 73.5 21t99 .5T788 313q12 13 22.5 28t17.5 33.5t12 32t9.5 37t6 33t4.5 34t4 28.5q-38 54-97 85.5T640 656q-120 0-201-89"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vsquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0m486 1493h147q28 0 44-19.5t31-55.5l447-1044q11-27-15-51t-59-24h-149q-28 0-44 19.5t-31 55.5l-298 695l-298-695q-10-24-17.5-37.5T557 311t-34-12H374q-33 0-59.5 24T299 374l448 1044q15 36 31 55.5t44 19.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeOn(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M920 216v1306q0 32-21.5 54.5T846 1599q-32 0-56-24l-400-399H76q-32 0-54-22t-22-54V638q0-32 22-54t54-22h314l400-400q25-25 56-25t52.5 23.5T920 216m459 654q0 91-49.5 170T1197 1151q-17 6-32 6q-32 0-54.5-22t-22.5-55q0-25 14.5-42.5t35-30t41-27.5t35-42.5T1228 870t-14.5-68.5t-35-43.5t-41-27.5t-35-30T1088 658q0-33 22.5-55.5T1165 580q13 0 32 7q83 32 132.5 111t49.5 172m311 0q0 182-102 336.5T1316 1433q-14 7-29 7q-31 0-54.5-23t-23.5-55q0-43 48-69q53-25 90-54q89-64 139.5-162t50.5-207q0-111-50-208.5T1347 499q-34-27-90-53q-48-26-48-72q0-31 23-54t53-23q17 0 31 7q170 72 272 227t102 339m307 0q0 276-152.5 506.5T1438 1716q-14 7-32 7q-32 0-54.5-22.5T1329 1646q0-46 47-71q8-6 27-13.5t27-12.5q50-25 101-60q146-109 229.5-273.5T1844 870q0-183-83-348t-230-274q-45-31-101-59q-8-5-26.5-13t-27.5-14q-47-25-47-69q0-31 22.5-54t54.5-23q4 0 32 7q254 109 406.5 339.5T1997 870"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeTimes(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1264 421l193 195l194-195q17-17 37.5-23.5t38.5-2t34.5 14.5t27 27t14.5 35.5t-2.5 39T1777 549l-195 193l195 193q17 17 23.5 37.5t2.5 38.5t-14.5 34.5t-27 27t-34.5 14.5t-38.5-3t-37.5-24l-194-192l-193 192q-15 15-33 22t-34 6t-31.5-8.5t-27-19t-19-27t-8.5-31.5t6-34t22-33l192-193l-192-193q-17-17-24-37.5t-3-38.5t14.5-35t27-27.5t34.5-15t38.5 2T1264 421M920 88v1306q0 32-21.5 54.5T846 1471q-32 0-56-24l-400-399H76q-32 0-54-22T0 972V510q0-32 22-54t54-22h314L790 34q25-25 56-25t52.5 23.5T920 88"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wsquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0m918 1493q12 0 21-12.5t13-25t11-37.5l269-1044q7-27-9-51t-36-24h-179q-12 0-21 12.5t-13 25t-11 37.5l-134 521l-134-521q-6-24-11-37.5t-14-25.5t-20-12H806q-16 0-25.5 19.5T761 374L627 895L493 374q-7-25-11-37.5t-13-25t-21-12.5H269q-20 0-36 24t-9 51l269 1044q7 25 11 37.5t13 25t21 12.5h178q27 0 45-75l135-523l135 523q18 75 45 75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Walk(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M640 0q-66 0-113 47t-47 113t47 113t113 47t113-47t47-113t-47-113T640 0M275 1616l169-277q28-46 42-92l58-192l175 183l153 367q11 27 34.5 43t53.5 16q40 0 68-28t28-68q0-17-7-36l-140-345q-13-33-20.5-44.5T855 1103L658 891l60-282l56 109q16 34 44 49q6 3 207 105q48 24 63 24q28 0 46-18t18-46q0-41-38-61L885 649l-90-206q-25-55-66-60l-229-26q-32-3-86 19l-242 96q-20 8-39 27t-25 39Q29 742 4 810q-4 9-4 22q0 26 19 45t45 19q20 0 40.5-9.5T131 862l102-268l210-79l-147 687l-188 320q-12 23-12 46q0 40 28 68t68 28q54 0 83-48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WeddingCake(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M664 162q-31 0-55.5-24T584 81t24.5-57T664 0q32 0 58 24t26 57t-26 57t-58 24m294 0q-32 0-57-24t-25-57t25-57t57-24t57 24t25 57t-25 57t-57 24m627 1620h-69v-326q0 43-40 73.5t-98 30.5q-57 0-97-30.5t-40-73.5q0 43-40.5 73.5T1103 1560t-97.5-30.5T965 1456q0 43-40.5 73.5T827 1560t-97.5-30.5T689 1456q0 43-40.5 73.5T551 1560t-97-30.5t-40-73.5q0 43-40.5 73.5T276 1560t-97.5-30.5T138 1456v326H69q-29 0-49 20t-20 49t20 49t49 20h1516q29 0 49-20t20-49t-20-49t-49-20m-344-455q0 43 40 73.5t97 30.5q58 0 98-30t40-74v-61q0-43-28-73.5t-69-30.5H237q-41 0-70 30.5t-29 73.5v61q0 43 40.5 73.5T276 1431t97.5-30.5T414 1327q0 43 40 73.5t97 30.5t97.5-30.5T689 1327q0 43 40.5 73.5T827 1431t97.5-30.5T965 1327q0 43 40.5 73.5t97.5 30.5t97.5-30.5t40.5-73.5m0-234V757q0 43-40.5 74t-97.5 31t-97.5-31t-40.5-74q0 43-40.5 74T827 862t-97.5-31t-40.5-74q0 43-40.5 74T551 862t-97-30.5t-40-74.5v336zM840 473h234l-78-162v-58q0-17-12-17h-55q-11 0-11 17v58zm290 0h34q34 0 55.5 29.5T1241 576v52q0 43-40.5 74t-97.5 31t-97.5-31t-40.5-74q0 43-40.5 74T827 733t-97.5-31t-40.5-74q0 43-40.5 74T551 733t-97-30.5t-40-74.5v-52q0-43 24-73t58-30h102V345h-45V231q0-11 13.5-25t23.5-14h151q10 0 23 14t13 25v114h-44v128h51q18-46 83-174v-69q0-11 14.5-24.5T908 192h98q12 0 26.5 13.5T1047 230v69q65 128 83 174"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Whiteboard(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 64q0-26 19-45t45-19h1408q26 0 45 19t19 45v1024h98q12 0 21 9t9 21v68q0 12-9 21t-21 9h-579l131 487q7 25-6.5 48.5T1268 1782t-48.5-6.5t-29.5-38.5l-140-521h-90v384q0 26-19 45t-45 19t-45-19t-19-45v-384h-90l-140 521q-7 25-29.5 38.5T524 1782t-39.5-30.5t-6.5-48.5l131-487H30q-12 0-21-9t-9-21v-68q0-12 9-21t21-9h98zm128 1024h1280V128H256zm128-64V915q0-14 9.5-24t23.5-10h318q14 0 23.5 10t9.5 24v109z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Window(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 0Q38 0 19 19T0 64v1536q0 26 19 45t45 19h1536q26 0 45-19t19-45V64q0-26-19-45t-45-19zm1408 896v576H896V896zm-704 576H192V896h576zM896 192h576v576H896zm-704 0h576v576H192zm960 384q33 55 124 91.5t132 36.5V256h-384q27 116 51 175t77 145m-640 0q53-86 77-145t51-175H256v448q41 0 132-36.5T512 576m832 448q-26 66-44.5 171t-19.5 213h128V960q-24 0-37 13.5t-27 50.5m-1024 0q-14-37-27-50.5T256 960v448h128q-1-108-19.5-213T320 1024"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wine(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M500 2000q-253 2-394-26q-48-10-77-64T0 1784v-399q0-151 42-291t99-193q60-56 124.5-159T343 592v-20h318v20q13 47 77.5 150T863 901q56 53 96.5 192.5T1000 1385v399q0 73-27.5 127t-74.5 63q-141 28-398 26M394 0h217q17 0 26 3t16.5 18.5T661 66v386H343V66q0-29 7.5-44.5t17-18.5T394 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WineO(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M547 2190q227 2 394-35q41-9 72.5-40.5t48-74t24.5-85.5t8-84v-589q0-156-70-349q-17-46-64-112.5T868 700t-80.5-107.5T752 514V139q0-71-12.5-105T692 0H402q-36 0-48 33.5T342 139v375q0 25-35.5 78.5T226 700t-92 120.5T70 933Q0 1114 0 1282v589q0 41 8 83.5t24.5 85.5t47.5 74.5t72 40.5q164 37 395 35M465 138h162v393q0 47 29 107.5T721.5 751T801 872.5T864 997q62 164 62 280v581q0 75-20 112t-83 48q-132 23-276 23t-276-23q-64-11-84-48t-20-112v-581q0-116 62-280q20-55 62.5-124.5T371 751t65.5-112.5T465 531z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xsquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0m560 1097l265 321q19 24 29 35t30.5 25.5t38.5 14.5h159q27 0 51-15t24-40q0-20-34-60l-399-482l399-482q34-40 34-60q0-25-24-40t-51-15h-159q-18 0-38.5 14.5T1190 339t-29 35L896 695L631 374q-19-24-29-35t-30.5-25.5T533 299H374q-27 0-51 15t-24 40q0 20 34 60l399 482l-399 482q-34 40-34 60q0 25 24 40t51 15h159q18 0 38.5-14.5T602 1453t29-35z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ysquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0m410 1024l1 394q0 31 22 53t52 22h150q31 0 52.5-22t21.5-53l1-394q146-215 436-645q11-20 11-36q0-44-75-44h-121q-51 0-102 75L896 796L597 374q-51-75-102-75H374q-75 0-75 44q0 16 11 36q290 430 436 645"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YahooJapan(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1733 481l46-73h-715l12 78q19 0 103 11t118 18q-3 30-85.5 103.5t-174 146.5T934 855q-22-32-160-211T568 366q26-7 135.5-15T829 340l12-71H9l-9 75q18 4 128.5 18.5T253 382q66 50 262.5 293.5T726 959v256q0 29-20 45.5t-56 26.5q-86 24-163 17l-21 85q59 2 114 2t145.5-1.5T837 1388q400 0 421 3l10-96l-247-17q-5-298 0-324q13-37 131-142.5t237-196t143-96.5q185-35 201-38m464 6l-390 663l-135-38l217-718zm-674 896l174 49l67-149l-179-50z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Zsquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0h1120q139 0 237.5 98.5T1792 336m-990 859l607-594q30-32 30-78V373q0-31-22-52.5t-53-21.5H428q-31 0-53 21.5T353 373v150q0 31 22 52.5t53 21.5h562l-608 594q-29 32-29 78v150q0 31 22 52.5t52 21.5h937q30 0 52.5-22t22.5-52v-150q0-30-22.5-52t-52.5-22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZeroSquare(children ...ElementRenderer) *VsIcon {
	return &VsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0m560 1493q150 0 276.5-80t200.5-217.5t74-299.5t-74-299.5T1172.5 379T896 299t-276.5 80T419 596.5T345 896t74 299.5T619.5 1413t276.5 80m-77-271q40 10 75 10q127 0 216.5-97t89.5-239q0-107-56-195zm154-652q-40-10-75-10q-128 1-217 97.5T592 896q0 107 56 195z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
