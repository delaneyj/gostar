package fa_6_regular

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "6.5.1"
	hAttr          = "1em"
	viewbox        = "0 0 512 512"
	fill           = "currentColor"
)

type Fa6RegularIcon struct {
	*SVGSVGElement
}

type Fa6RegularIconFn func(children ...ElementRenderer) *Fa6RegularIcon

var IconLookup = map[string]Fa6RegularIconFn{
	"addressBook":          AddressBook,
	"addressCard":          AddressCard,
	"bell":                 Bell,
	"bellSlash":            BellSlash,
	"bookmark":             Bookmark,
	"building":             Building,
	"calendar":             Calendar,
	"calendarCheck":        CalendarCheck,
	"calendarDays":         CalendarDays,
	"calendarMinus":        CalendarMinus,
	"calendarPlus":         CalendarPlus,
	"calendarXmark":        CalendarXmark,
	"chartBar":             ChartBar,
	"chessBishop":          ChessBishop,
	"chessKing":            ChessKing,
	"chessKnight":          ChessKnight,
	"chessPawn":            ChessPawn,
	"chessQueen":           ChessQueen,
	"chessRook":            ChessRook,
	"circle":               Circle,
	"circleCheck":          CircleCheck,
	"circleDot":            CircleDot,
	"circleDown":           CircleDown,
	"circleLeft":           CircleLeft,
	"circlePause":          CirclePause,
	"circlePlay":           CirclePlay,
	"circleQuestion":       CircleQuestion,
	"circleRight":          CircleRight,
	"circleStop":           CircleStop,
	"circleUp":             CircleUp,
	"circleUser":           CircleUser,
	"circleXmark":          CircleXmark,
	"clipboard":            Clipboard,
	"clock":                Clock,
	"clone":                Clone,
	"closedCaptioning":     ClosedCaptioning,
	"comment":              Comment,
	"commentDots":          CommentDots,
	"comments":             Comments,
	"compass":              Compass,
	"copy":                 Copy,
	"copyright":            Copyright,
	"creditCard":           CreditCard,
	"envelope":             Envelope,
	"envelopeOpen":         EnvelopeOpen,
	"eye":                  Eye,
	"eyeSlash":             EyeSlash,
	"faceAngry":            FaceAngry,
	"faceDizzy":            FaceDizzy,
	"faceFlushed":          FaceFlushed,
	"faceFrown":            FaceFrown,
	"faceFrownOpen":        FaceFrownOpen,
	"faceGrimace":          FaceGrimace,
	"faceGrin":             FaceGrin,
	"faceGrinBeam":         FaceGrinBeam,
	"faceGrinBeamSweat":    FaceGrinBeamSweat,
	"faceGrinHearts":       FaceGrinHearts,
	"faceGrinSquint":       FaceGrinSquint,
	"faceGrinSquintTears":  FaceGrinSquintTears,
	"faceGrinStars":        FaceGrinStars,
	"faceGrinTears":        FaceGrinTears,
	"faceGrinTongue":       FaceGrinTongue,
	"faceGrinTongueSquint": FaceGrinTongueSquint,
	"faceGrinTongueWink":   FaceGrinTongueWink,
	"faceGrinWide":         FaceGrinWide,
	"faceGrinWink":         FaceGrinWink,
	"faceKiss":             FaceKiss,
	"faceKissBeam":         FaceKissBeam,
	"faceKissWinkHeart":    FaceKissWinkHeart,
	"faceLaugh":            FaceLaugh,
	"faceLaughBeam":        FaceLaughBeam,
	"faceLaughSquint":      FaceLaughSquint,
	"faceLaughWink":        FaceLaughWink,
	"faceMeh":              FaceMeh,
	"faceMehBlank":         FaceMehBlank,
	"faceRollingEyes":      FaceRollingEyes,
	"faceSadCry":           FaceSadCry,
	"faceSadTear":          FaceSadTear,
	"faceSmile":            FaceSmile,
	"faceSmileBeam":        FaceSmileBeam,
	"faceSmileWink":        FaceSmileWink,
	"faceSurprise":         FaceSurprise,
	"faceTired":            FaceTired,
	"file":                 File,
	"fileAudio":            FileAudio,
	"fileCode":             FileCode,
	"fileExcel":            FileExcel,
	"fileImage":            FileImage,
	"fileLines":            FileLines,
	"filePdf":              FilePdf,
	"filePowerpoint":       FilePowerpoint,
	"fileVideo":            FileVideo,
	"fileWord":             FileWord,
	"fileZipper":           FileZipper,
	"flag":                 Flag,
	"floppyDisk":           FloppyDisk,
	"folder":               Folder,
	"folderClosed":         FolderClosed,
	"folderOpen":           FolderOpen,
	"fontAwesome":          FontAwesome,
	"futbol":               Futbol,
	"gem":                  Gem,
	"hand":                 Hand,
	"handBackFist":         HandBackFist,
	"handLizard":           HandLizard,
	"handPeace":            HandPeace,
	"handPointDown":        HandPointDown,
	"handPointLeft":        HandPointLeft,
	"handPointRight":       HandPointRight,
	"handPointUp":          HandPointUp,
	"handPointer":          HandPointer,
	"handScissors":         HandScissors,
	"handSpock":            HandSpock,
	"handshake":            Handshake,
	"hardDrive":            HardDrive,
	"heart":                Heart,
	"hospital":             Hospital,
	"hourglass":            Hourglass,
	"hourglassHalf":        HourglassHalf,
	"idBadge":              IdBadge,
	"idCard":               IdCard,
	"image":                Image,
	"images":               Images,
	"keyboard":             Keyboard,
	"lemon":                Lemon,
	"lifeRing":             LifeRing,
	"lightbulb":            Lightbulb,
	"map":                  Map,
	"message":              Message,
	"moneyBillOne":         MoneyBillOne,
	"moon":                 Moon,
	"newspaper":            Newspaper,
	"notdef":               Notdef,
	"noteSticky":           NoteSticky,
	"objectGroup":          ObjectGroup,
	"objectUngroup":        ObjectUngroup,
	"paperPlane":           PaperPlane,
	"paste":                Paste,
	"penToSquare":          PenToSquare,
	"rectangleList":        RectangleList,
	"rectangleXmark":       RectangleXmark,
	"registered":           Registered,
	"shareFromSquare":      ShareFromSquare,
	"snowflake":            Snowflake,
	"square":               Square,
	"squareCaretDown":      SquareCaretDown,
	"squareCaretLeft":      SquareCaretLeft,
	"squareCaretRight":     SquareCaretRight,
	"squareCaretUp":        SquareCaretUp,
	"squareCheck":          SquareCheck,
	"squareFull":           SquareFull,
	"squareMinus":          SquareMinus,
	"squarePlus":           SquarePlus,
	"star":                 Star,
	"starHalf":             StarHalf,
	"starHalfStroke":       StarHalfStroke,
	"sun":                  Sun,
	"thumbsDown":           ThumbsDown,
	"thumbsUp":             ThumbsUp,
	"trashCan":             TrashCan,
	"user":                 User,
	"windowMaximize":       WindowMaximize,
	"windowMinimize":       WindowMinimize,
	"windowRestore":        WindowRestore,
}

func AddressBook(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 48c8.8 0 16 7.2 16 16v384c0 8.8-7.2 16-16 16H96c-8.8 0-16-7.2-16-16V64c0-8.8 7.2-16 16-16zM96 0C60.7 0 32 28.7 32 64v384c0 35.3 28.7 64 64 64h288c35.3 0 64-28.7 64-64V64c0-35.3-28.7-64-64-64zm144 256a64 64 0 1 0 0-128a64 64 0 1 0 0 128m-32 32c-44.2 0-80 35.8-80 80c0 8.8 7.2 16 16 16h192c8.8 0 16-7.2 16-16c0-44.2-35.8-80-80-80zM512 80c0-8.8-7.2-16-16-16s-16 7.2-16 16v64c0 8.8 7.2 16 16 16s16-7.2 16-16zm-16 112c-8.8 0-16 7.2-16 16v64c0 8.8 7.2 16 16 16s16-7.2 16-16v-64c0-8.8-7.2-16-16-16m16 144c0-8.8-7.2-16-16-16s-16 7.2-16 16v64c0 8.8 7.2 16 16 16s16-7.2 16-16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddressCard(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 80c8.8 0 16 7.2 16 16v320c0 8.8-7.2 16-16 16H64c-8.8 0-16-7.2-16-16V96c0-8.8 7.2-16 16-16zM64 32C28.7 32 0 60.7 0 96v320c0 35.3 28.7 64 64 64h448c35.3 0 64-28.7 64-64V96c0-35.3-28.7-64-64-64zm144 224a64 64 0 1 0 0-128a64 64 0 1 0 0 128m-32 32c-44.2 0-80 35.8-80 80c0 8.8 7.2 16 16 16h192c8.8 0 16-7.2 16-16c0-44.2-35.8-80-80-80zm200-144c-13.3 0-24 10.7-24 24s10.7 24 24 24h80c13.3 0 24-10.7 24-24s-10.7-24-24-24zm0 96c-13.3 0-24 10.7-24 24s10.7 24 24 24h80c13.3 0 24-10.7 24-24s-10.7-24-24-24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M224 0c-17.7 0-32 14.3-32 32v19.2C119 66 64 130.6 64 208v25.4c0 45.4-15.5 89.5-43.8 124.9L5.3 377c-5.8 7.2-6.9 17.1-2.9 25.4S14.8 416 24 416h400c9.2 0 17.6-5.3 21.6-13.6s2.9-18.2-2.9-25.4l-14.9-18.6c-28.3-35.5-43.8-79.6-43.8-125V208c0-77.4-55-142-128-156.8V32c0-17.7-14.3-32-32-32m0 96c61.9 0 112 50.1 112 112v25.4c0 47.9 13.9 94.6 39.7 134.6H72.3c25.8-40 39.7-86.7 39.7-134.6V208c0-61.9 50.1-112 112-112m64 352H160c0 17 6.7 33.3 18.7 45.3S207 512 224 512s33.3-6.7 45.3-18.7S288 465 288 448"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellSlash(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M38.8 5.1C28.4-3.1 13.3-1.2 5.1 9.2s-6.3 25.5 4.1 33.7l592 464c10.4 8.2 25.5 6.3 33.7-4.1s6.3-25.5-4.1-33.7L542.6 400c2.7-7.8 1.3-16.5-3.9-23l-14.9-18.6c-28.3-35.5-43.8-79.6-43.8-125V208c0-77.4-55-142-128-156.8V32c0-17.7-14.3-32-32-32s-32 14.3-32 32v19.2c-42.6 8.6-79 34.2-102 69.3zM224 150.3c19.6-32.6 55.3-54.3 96-54.3c61.9 0 112 50.1 112 112v25.4c0 32.7 6.4 64.8 18.7 94.5zM406.2 416l-60.9-48h-177c21.2-32.8 34.4-70.3 38.4-109.1L160 222.1v11.4c0 45.4-15.5 89.5-43.8 124.9L101.3 377c-5.8 7.2-6.9 17.1-2.9 25.4S110.8 416 120 416zM384 448H256c0 17 6.7 33.3 18.7 45.3S303 512 320 512s33.3-6.7 45.3-18.7S384 465 384 448"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 48C0 21.5 21.5 0 48 0v441.4l130.1-92.9c8.3-6 19.6-6 27.9 0l130 92.9V48H48V0h288c26.5 0 48 21.5 48 48v440c0 9-5 17.2-13 21.3s-17.6 3.4-24.9-1.8L192 397.5l-154.1 110c-7.3 5.2-16.9 5.9-24.9 1.8S0 497 0 488z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Building(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 48c-8.8 0-16 7.2-16 16v384c0 8.8 7.2 16 16 16h80v-64c0-26.5 21.5-48 48-48s48 21.5 48 48v64h80c8.8 0 16-7.2 16-16V64c0-8.8-7.2-16-16-16zM0 64C0 28.7 28.7 0 64 0h256c35.3 0 64 28.7 64 64v384c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64zm88 40c0-8.8 7.2-16 16-16h48c8.8 0 16 7.2 16 16v48c0 8.8-7.2 16-16 16h-48c-8.8 0-16-7.2-16-16zm144-16h48c8.8 0 16 7.2 16 16v48c0 8.8-7.2 16-16 16h-48c-8.8 0-16-7.2-16-16v-48c0-8.8 7.2-16 16-16M88 232c0-8.8 7.2-16 16-16h48c8.8 0 16 7.2 16 16v48c0 8.8-7.2 16-16 16h-48c-8.8 0-16-7.2-16-16zm144-16h48c8.8 0 16 7.2 16 16v48c0 8.8-7.2 16-16 16h-48c-8.8 0-16-7.2-16-16v-48c0-8.8 7.2-16 16-16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M152 24c0-13.3-10.7-24-24-24s-24 10.7-24 24v40H64C28.7 64 0 92.7 0 128v320c0 35.3 28.7 64 64 64h320c35.3 0 64-28.7 64-64V128c0-35.3-28.7-64-64-64h-40V24c0-13.3-10.7-24-24-24s-24 10.7-24 24v40H152zM48 192h352v256c0 8.8-7.2 16-16 16H64c-8.8 0-16-7.2-16-16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarCheck(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 0c13.3 0 24 10.7 24 24v40h144V24c0-13.3 10.7-24 24-24s24 10.7 24 24v40h40c35.3 0 64 28.7 64 64v320c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64V128c0-35.3 28.7-64 64-64h40V24c0-13.3 10.7-24 24-24m272 192H48v256c0 8.8 7.2 16 16 16h320c8.8 0 16-7.2 16-16zm-71 105L217 409c-9.4 9.4-24.6 9.4-33.9 0l-64-64c-9.4-9.4-9.4-24.6 0-33.9s24.6-9.4 33.9 0l47 47l95-95c9.4-9.4 24.6-9.4 33.9 0s9.4 24.6 0 33.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarDays(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M152 24c0-13.3-10.7-24-24-24s-24 10.7-24 24v40H64C28.7 64 0 92.7 0 128v320c0 35.3 28.7 64 64 64h320c35.3 0 64-28.7 64-64V128c0-35.3-28.7-64-64-64h-40V24c0-13.3-10.7-24-24-24s-24 10.7-24 24v40H152zM48 192h80v56H48zm0 104h80v64H48zm128 0h96v64h-96zm144 0h80v64h-80zm80-48h-80v-56h80zm0 160v40c0 8.8-7.2 16-16 16h-64v-56zm-128 0v56h-96v-56zm-144 0v56H64c-8.8 0-16-7.2-16-16v-40zm144-160h-96v-56h96z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarMinus(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 0c13.3 0 24 10.7 24 24v40h144V24c0-13.3 10.7-24 24-24s24 10.7 24 24v40h40c35.3 0 64 28.7 64 64v320c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64V128c0-35.3 28.7-64 64-64h40V24c0-13.3 10.7-24 24-24m272 192H48v256c0 8.8 7.2 16 16 16h320c8.8 0 16-7.2 16-16zM296 352H152c-13.3 0-24-10.7-24-24s10.7-24 24-24h144c13.3 0 24 10.7 24 24s-10.7 24-24 24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarPlus(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M152 24c0-13.3-10.7-24-24-24s-24 10.7-24 24v40H64C28.7 64 0 92.7 0 128v320c0 35.3 28.7 64 64 64h320c35.3 0 64-28.7 64-64V128c0-35.3-28.7-64-64-64h-40V24c0-13.3-10.7-24-24-24s-24 10.7-24 24v40H152zM48 192h352v256c0 8.8-7.2 16-16 16H64c-8.8 0-16-7.2-16-16zm176 40c-13.3 0-24 10.7-24 24v48h-48c-13.3 0-24 10.7-24 24s10.7 24 24 24h48v48c0 13.3 10.7 24 24 24s24-10.7 24-24v-48h48c13.3 0 24-10.7 24-24s-10.7-24-24-24h-48v-48c0-13.3-10.7-24-24-24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarXmark(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 0c13.3 0 24 10.7 24 24v40h144V24c0-13.3 10.7-24 24-24s24 10.7 24 24v40h40c35.3 0 64 28.7 64 64v320c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64V128c0-35.3 28.7-64 64-64h40V24c0-13.3 10.7-24 24-24m272 192H48v256c0 8.8 7.2 16 16 16h320c8.8 0 16-7.2 16-16zm-95 89l-47 47l47 47c9.4 9.4 9.4 24.6 0 33.9s-24.6 9.4-33.9 0l-47-47l-47 47c-9.4 9.4-24.6 9.4-33.9 0s-9.4-24.6 0-33.9l47-47l-47-47c-9.4-9.4-9.4-24.6 0-33.9s24.6-9.4 33.9 0l47 47l47-47c9.4-9.4 24.6-9.4 33.9 0s9.4 24.6 0 33.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartBar(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24 32c13.3 0 24 10.7 24 24v352c0 13.3 10.7 24 24 24h416c13.3 0 24 10.7 24 24s-10.7 24-24 24H72c-39.8 0-72-32.2-72-72V56c0-13.3 10.7-24 24-24m104 104c0-13.3 10.7-24 24-24h208c13.3 0 24 10.7 24 24s-10.7 24-24 24H152c-13.3 0-24-10.7-24-24m24 72h144c13.3 0 24 10.7 24 24s-10.7 24-24 24H152c-13.3 0-24-10.7-24-24s10.7-24 24-24m0 96h272c13.3 0 24 10.7 24 24s-10.7 24-24 24H152c-13.3 0-24-10.7-24-24s10.7-24 24-24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChessBishop(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M104 0C90.7 0 80 10.7 80 24c0 11.2 7.6 20.6 18 23.2c-7.8 8-16.1 17-24.4 27C38.2 116.7 0 178.8 0 250.9c0 44.8 24.6 72.2 48 87.8V352h48v-27c0-9-5-17.2-13-21.3c-18-9.3-35-24.7-35-52.7c0-55.5 29.8-106.8 62.4-145.9c16-19.2 32.1-34.8 44.2-45.5c1.9-1.7 3.7-3.2 5.3-4.6c1.7 1.4 3.4 3 5.3 4.6c12.1 10.7 28.2 26.3 44.2 45.5c5.3 6.3 10.5 13 15.5 20L159 191c-9.4 9.4-9.4 24.6 0 33.9s24.6 9.4 33.9 0l57.8-57.8c12.8 25.9 21.2 54.3 21.2 83.8c0 28-17 43.4-35 52.7c-8 4.1-13 12.3-13 21.3v27h48v-13.2c23.4-15.6 48-42.9 48-87.8c0-72.1-38.2-134.2-73.6-176.7c-8.3-9.9-16.6-19-24.4-27c10.3-2.7 18-12.1 18-23.2c0-13.3-10.7-24-24-24H160zM52.7 464l16.6-32h181.5l16.6 32zm207.9-80H59.5c-12 0-22.9 6.7-28.4 17.3L4.6 452.5c-3 5.8-4.6 12.2-4.6 18.7C0 493.8 18.2 512 40.8 512h238.4c22.5 0 40.8-18.2 40.8-40.8c0-6.5-1.6-12.9-4.6-18.7l-26.5-51.2c-5.5-10.6-16.5-17.3-28.4-17.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChessKing(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M248 24c0-13.3-10.7-24-24-24s-24 10.7-24 24v32h-32c-13.3 0-24 10.7-24 24s10.7 24 24 24h32v40H59.6C26.7 144 0 170.7 0 203.6c0 8.2 1.7 16.3 4.9 23.8L59.1 352h52.3L49 208.2c-.6-1.5-1-3-1-4.6c0-6.4 5.2-11.6 11.6-11.6h328.8c6.4 0 11.6 5.2 11.6 11.6c0 1.6-.3 3.2-1 4.6L336.5 352h52.3L443 227.4c3.3-7.5 4.9-15.6 4.9-23.8c0-32.9-26.7-59.6-59.6-59.6H248v-40h32c13.3 0 24-10.7 24-24s-10.7-24-24-24h-32zM101.2 432h245.6l16.6 32H84.7l16.6-32zm283.7-30.7c-5.5-10.6-16.5-17.3-28.4-17.3h-265c-12 0-22.9 6.7-28.4 17.3l-26.5 51.2c-3 5.8-4.6 12.2-4.6 18.7c0 22.6 18.2 40.8 40.8 40.8h302.4c22.5 0 40.8-18.2 40.8-40.8c0-6.5-1.6-12.9-4.6-18.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChessKnight(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M226.6 48H117.3l17.1 12.8c6 4.5 9.6 11.6 9.6 19.2s-3.6 14.7-9.6 19.2l-6.5 4.9c-10 7.5-16 19.3-16 31.9l-.3 91c0 10.2 4.9 19.9 13.2 25.8l1.9 1.3c9.9 7.1 23.3 7 33.2-.1l49.9-36.3c10.7-7.8 25.7-5.4 33.5 5.3s5.4 25.7-5.3 33.5l-49.9 36.3l-53.8 39.1c-7.3 5.3-13 12.2-16.9 20.1H66.8c5.3-22.1 17.8-41.9 35.9-56.3c-1.3-.8-2.6-1.7-3.8-2.6l-1.9-1.3c-21-15-33.4-39.2-33.3-65l.3-91c.1-19.8 6.7-38.7 18.6-53.9l-.4-.3C70.7 73 64 59.6 64 45.3C64 20.3 84.3 0 109.3 0h117.3C331.2 0 416 84.8 416 189.4c0 11.1-1 22.2-2.9 33.2l-23 129.4h-48.8l24.5-137.8c1.5-8.2 2.2-16.5 2.2-24.8C368 111.3 304.7 48 226.6 48M85.2 432l-16.5 32h310.6l-16.6-32zm315.7-30.7l26.5 51.2c3 5.8 4.6 12.2 4.6 18.7c0 22.5-18.2 40.8-40.8 40.8H56.8C34.2 512 16 493.8 16 471.2c0-6.5 1.6-12.9 4.6-18.7l26.5-51.2c5.4-10.6 16.4-17.3 28.4-17.3h297c12 0 22.9 6.7 28.4 17.3M172 128a20 20 0 1 1 0 40a20 20 0 1 1 0-40"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChessPawn(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M232 152a72 72 0 1 0-144 0a72 72 0 1 0 144 0m24 120h-12.6l10.7 80h-48.4L195 272h-70l-10.7 80H65.9l10.7-80H64c-13.3 0-24-10.7-24-24s10.7-24 24-24c-15.1-20.1-24-45-24-72C40 85.7 93.7 32 160 32s120 53.7 120 120c0 27-8.9 51.9-24 72c13.3 0 24 10.7 24 24s-10.7 24-24 24M52.7 464h214.6l-16.6-32H69.2zm207.9-80c12 0 22.9 6.7 28.4 17.3l26.5 51.2c3 5.8 4.6 12.2 4.6 18.7c0 22.5-18.2 40.8-40.8 40.8H40.8C18.2 512 0 493.8 0 471.2c0-6.5 1.6-12.9 4.6-18.7l26.5-51.2c5.4-10.6 16.4-17.3 28.4-17.3h201z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChessQueen(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 96a48 48 0 1 0 0-96a48 48 0 1 0 0 96m-95.2-8c-18.1 0-31.3 12.8-35.6 26.9c-8 26.2-32.4 45.2-61.2 45.2c-10 0-19.4-2.3-27.7-6.3c-7.6-3.7-16.7-3.3-24 1.2c-11.6 7.1-15.4 22.1-8.6 33.9L97.6 352H153L70 207.9c40.5-2.2 75.3-25.9 93.1-59.8c22 26.8 55.4 43.9 92.8 43.9s70.8-17.1 92.8-43.9c17.8 34 52.6 57.7 93.1 59.8L359 352h55.4l93.9-163.1c6.8-11.7 3-26.7-8.6-33.8c-7.3-4.5-16.4-4.9-24-1.2c-8.4 4-17.7 6.3-27.7 6.3c-28.8 0-53.2-19-61.2-45.2c-4.3-14.2-17.5-27-35.6-27c-14.5 0-26.3 8.5-32.4 19.3c-12.4 22-35.9 36.7-62.8 36.7s-50.4-14.8-62.8-36.7C187.1 96.5 175.4 88 160.8 88m-27.6 344h245.6l16.6 32H116.7l16.6-32zm283.7-30.7c-5.5-10.6-16.5-17.3-28.4-17.3h-265c-12 0-22.9 6.7-28.4 17.3l-26.5 51.2c-3 5.8-4.6 12.2-4.6 18.7c0 22.5 18.2 40.8 40.8 40.8h302.4c22.5 0 40.8-18.2 40.8-40.8c0-6.5-1.6-12.9-4.6-18.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChessRook(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M80 80v112c0 2.5 1.2 4.9 3.2 6.4l51.2 38.4c6.8 5.1 10.4 13.4 9.5 21.9L133.5 352H85.2l9.4-85l-40.2-30.2C40.3 226.2 32 209.6 32 192V72c0-22.1 17.9-40 40-40h304c22.1 0 40 17.9 40 40v120c0 17.6-8.3 34.2-22.4 44.8L353.4 267l9.4 85h-48.3l-10.4-93.3c-.9-8.4 2.7-16.8 9.5-21.9l51.2-38.4c2-1.5 3.2-3.9 3.2-6.4V80h-64v24c0 13.3-10.7 24-24 24s-24-10.7-24-24V80h-64v24c0 13.3-10.7 24-24 24s-24-10.7-24-24V80zm4.7 384h278.6l-16.6-32H101.2zm271.9-80c12 0 22.9 6.7 28.4 17.3l26.5 51.2c3 5.8 4.6 12.2 4.6 18.7c0 22.5-18.2 40.8-40.8 40.8H72.8C50.2 512 32 493.8 32 471.2c0-6.5 1.6-12.9 4.6-18.7l26.5-51.2c5.4-10.6 16.4-17.3 28.4-17.3h265zM208 288c-8.8 0-16-7.2-16-16v-48c0-17.7 14.3-32 32-32s32 14.3 32 32v48c0 8.8-7.2 16-16 16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circle(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleCheck(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 48a208 208 0 1 1 0 416a208 208 0 1 1 0-416m0 464a256 256 0 1 0 0-512a256 256 0 1 0 0 512m113-303c9.4-9.4 9.4-24.6 0-33.9s-24.6-9.4-33.9 0l-111 111l-47-47c-9.4-9.4-24.6-9.4-33.9 0s-9.4 24.6 0 33.9l64 64c9.4 9.4 24.6 9.4 33.9 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleDot(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m256-96a96 96 0 1 1 0 192a96 96 0 1 1 0-192"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleDown(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 464a208 208 0 1 1 0-416a208 208 0 1 1 0 416m0-464a256 256 0 1 0 0 512a256 256 0 1 0 0-512m120.9 294.6c4.5-4.2 7.1-10.1 7.1-16.3c0-12.3-10-22.3-22.3-22.3H304v-96c0-17.7-14.3-32-32-32h-32c-17.7 0-32 14.3-32 32v96h-57.7c-12.3 0-22.3 10-22.3 22.3c0 6.2 2.6 12.1 7.1 16.3l107.1 99.9c3.8 3.5 8.7 5.5 13.8 5.5s10.1-2 13.8-5.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleLeft(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M48 256a208 208 0 1 1 416 0a208 208 0 1 1-416 0m464 0a256 256 0 1 0-512 0a256 256 0 1 0 512 0M217.4 376.9c4.2 4.5 10.1 7.1 16.3 7.1c12.3 0 22.3-10 22.3-22.3V304h96c17.7 0 32-14.3 32-32v-32c0-17.7-14.3-32-32-32h-96v-57.7c0-12.3-10-22.3-22.3-22.3c-6.2 0-12.1 2.6-16.3 7.1l-99.9 107.1c-3.5 3.8-5.5 8.7-5.5 13.8s2 10.1 5.5 13.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CirclePause(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m224-72v144c0 13.3-10.7 24-24 24s-24-10.7-24-24V184c0-13.3 10.7-24 24-24s24 10.7 24 24m112 0v144c0 13.3-10.7 24-24 24s-24-10.7-24-24V184c0-13.3 10.7-24 24-24s24 10.7 24 24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CirclePlay(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m188.3-108.9c7.6-4.2 16.8-4.1 24.3.5l144 88c7.1 4.4 11.5 12.1 11.5 20.5s-4.4 16.1-11.5 20.5l-144 88c-7.4 4.5-16.7 4.7-24.3.5S176 352.9 176 344.2V168c0-8.7 4.7-16.7 12.3-20.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleQuestion(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m169.8-90.7c7.9-22.3 29.1-37.3 52.8-37.3h58.3c34.9 0 63.1 28.3 63.1 63.1c0 22.6-12.1 43.5-31.7 54.8L280 264.4c-.2 13-10.9 23.6-24 23.6c-13.3 0-24-10.7-24-24v-13.5c0-8.6 4.6-16.5 12.1-20.8l44.3-25.4c4.7-2.7 7.6-7.7 7.6-13.1c0-8.4-6.8-15.1-15.1-15.1h-58.3c-3.4 0-6.4 2.1-7.5 5.3l-.4 1.2c-4.4 12.5-18.2 19-30.6 14.6s-19-18.2-14.6-30.6l.4-1.2zM224 352a32 32 0 1 1 64 0a32 32 0 1 1-64 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleRight(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 1-416 0a208 208 0 1 1 416 0M0 256a256 256 0 1 0 512 0a256 256 0 1 0-512 0m294.6-120.9c-4.2-4.5-10.1-7.1-16.3-7.1c-12.3 0-22.3 10-22.3 22.3V208h-96c-17.7 0-32 14.3-32 32v32c0 17.7 14.3 32 32 32h96v57.7c0 12.3 10 22.3 22.3 22.3c6.2 0 12.1-2.6 16.3-7.1l99.9-107.1c3.5-3.8 5.5-8.7 5.5-13.8s-2-10.1-5.5-13.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleStop(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m192-96h128c17.7 0 32 14.3 32 32v128c0 17.7-14.3 32-32 32H192c-17.7 0-32-14.3-32-32V192c0-17.7 14.3-32 32-32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleUp(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 48a208 208 0 1 1 0 416a208 208 0 1 1 0-416m0 464a256 256 0 1 0 0-512a256 256 0 1 0 0 512M135.1 217.4c-4.5 4.2-7.1 10.1-7.1 16.3c0 12.3 10 22.3 22.3 22.3H208v96c0 17.7 14.3 32 32 32h32c17.7 0 32-14.3 32-32v-96h57.7c12.3 0 22.3-10 22.3-22.3c0-6.2-2.6-12.1-7.1-16.3l-107.1-99.9c-3.8-3.5-8.7-5.5-13.8-5.5s-10.1 2-13.8 5.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleUser(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M406.5 399.6c-19.1-46.7-65-79.6-118.5-79.6h-64c-53.5 0-99.4 32.9-118.5 79.6C69.9 362.2 48 311.7 48 256c0-114.9 93.1-208 208-208s208 93.1 208 208c0 55.7-21.9 106.2-57.5 143.6m-40.1 32.7c-32 20.1-69.8 31.7-110.4 31.7s-78.4-11.6-110.5-31.7c7.3-36.7 39.7-64.3 78.5-64.3h64c38.8 0 71.2 27.6 78.5 64.3zM256 512a256 256 0 1 0 0-512a256 256 0 1 0 0 512m0-272a40 40 0 1 1 0-80a40 40 0 1 1 0 80m-88-40a88 88 0 1 0 176 0a88 88 0 1 0-176 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleXmark(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 48a208 208 0 1 1 0 416a208 208 0 1 1 0-416m0 464a256 256 0 1 0 0-512a256 256 0 1 0 0 512m-81-337c-9.4 9.4-9.4 24.6 0 33.9l47 47l-47 47c-9.4 9.4-9.4 24.6 0 33.9s24.6 9.4 33.9 0l47-47l47 47c9.4 9.4 24.6 9.4 33.9 0s9.4-24.6 0-33.9l-47-47l47-47c9.4-9.4 9.4-24.6 0-33.9s-24.6-9.4-33.9 0l-47 47l-47-47c-9.4-9.4-24.6-9.4-33.9 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clipboard(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M280 64h40c35.3 0 64 28.7 64 64v320c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64V128c0-35.3 28.7-64 64-64h49.6C121 27.5 153.3 0 192 0s71 27.5 78.4 64zM64 112c-8.8 0-16 7.2-16 16v320c0 8.8 7.2 16 16 16h256c8.8 0 16-7.2 16-16V128c0-8.8-7.2-16-16-16h-16v24c0 13.3-10.7 24-24 24H104c-13.3 0-24-10.7-24-24v-24zm128-8a24 24 0 1 0 0-48a24 24 0 1 0 0 48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 1-416 0a208 208 0 1 1 416 0M0 256a256 256 0 1 0 512 0a256 256 0 1 0-512 0m232-136v136c0 8 4 15.5 10.7 20l96 64c11 7.4 25.9 4.4 33.3-6.7s4.4-25.9-6.7-33.3L280 243.2V120c0-13.3-10.7-24-24-24s-24 10.7-24 24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clone(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 464h224c8.8 0 16-7.2 16-16v-64h48v64c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64V224c0-35.3 28.7-64 64-64h64v48H64c-8.8 0-16 7.2-16 16v224c0 8.8 7.2 16 16 16m160-160h224c8.8 0 16-7.2 16-16V64c0-8.8-7.2-16-16-16H224c-8.8 0-16 7.2-16 16v224c0 8.8 7.2 16 16 16m-64-16V64c0-35.3 28.7-64 64-64h224c35.3 0 64 28.7 64 64v224c0 35.3-28.7 64-64 64H224c-35.3 0-64-28.7-64-64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClosedCaptioning(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 80c8.8 0 16 7.2 16 16v320c0 8.8-7.2 16-16 16H64c-8.8 0-16-7.2-16-16V96c0-8.8 7.2-16 16-16zM64 32C28.7 32 0 60.7 0 96v320c0 35.3 28.7 64 64 64h448c35.3 0 64-28.7 64-64V96c0-35.3-28.7-64-64-64zm136 176c14.2 0 27 6.1 35.8 16c8.8 9.9 24 10.7 33.9 1.9s10.7-24 1.9-33.9c-17.5-19.6-43.1-32-71.5-32c-53 0-96 43-96 96s43 96 96 96c28.4 0 54-12.4 71.5-32c8.8-9.9 8-25-1.9-33.9s-25-8-33.9 1.9c-8.8 9.9-21.6 16-35.8 16c-26.5 0-48-21.5-48-48s21.5-48 48-48m144 48c0-26.5 21.5-48 48-48c14.2 0 27 6.1 35.8 16c8.8 9.9 24 10.7 33.9 1.9s10.7-24 1.9-33.9c-17.5-19.6-43.1-32-71.5-32c-53 0-96 43-96 96s43 96 96 96c28.4 0 54-12.4 71.5-32c8.8-9.9 8-25-1.9-33.9s-25-8-33.9 1.9c-8.8 9.9-21.6 16-35.8 16c-26.5 0-48-21.5-48-48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comment(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M123.6 391.3c12.9-9.4 29.6-11.8 44.6-6.4c26.5 9.6 56.2 15.1 87.8 15.1c124.7 0 208-80.5 208-160S380.7 80 256 80S48 160.5 48 240c0 32 12.4 62.8 35.7 89.2c8.6 9.7 12.8 22.5 11.8 35.5c-1.4 18.1-5.7 34.7-11.3 49.4c17-7.9 31.1-16.7 39.4-22.7zM21.2 431.9c1.8-2.7 3.5-5.4 5.1-8.1c10-16.6 19.5-38.4 21.4-62.9C17.7 326.8 0 285.1 0 240C0 125.1 114.6 32 256 32s256 93.1 256 208s-114.6 208-256 208c-37.1 0-72.3-6.4-104.1-17.9c-11.9 8.7-31.3 20.6-54.3 30.6c-15.1 6.6-32.3 12.6-50.1 16.1c-.8.2-1.6.3-2.4.5c-4.4.8-8.7 1.5-13.2 1.9c-.2 0-.5.1-.7.1c-5.1.5-10.2.8-15.3.8c-6.5 0-12.3-3.9-14.8-9.9S0 457.4 4.5 452.8c4.1-4.2 7.8-8.7 11.3-13.5c1.7-2.3 3.3-4.6 4.8-6.9c.1-.2.2-.3.3-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentDots(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M168.2 384.9c-15-5.4-31.7-3.1-44.6 6.4c-8.2 6-22.3 14.8-39.4 22.7c5.6-14.7 9.9-31.3 11.3-49.4c1-12.9-3.3-25.7-11.8-35.5C60.4 302.8 48 272 48 240c0-79.5 83.3-160 208-160s208 80.5 208 160s-83.3 160-208 160c-31.6 0-61.3-5.5-87.8-15.1M26.3 423.8c-1.6 2.7-3.3 5.4-5.1 8.1l-.3.5l-4.8 6.9c-3.5 4.7-7.3 9.3-11.3 13.5c-4.6 4.6-5.9 11.4-3.4 17.4c2.5 6 8.3 9.9 14.8 9.9c5.1 0 10.2-.3 15.3-.8l.7-.1c4.4-.5 8.8-1.1 13.2-1.9c.8-.1 1.6-.3 2.4-.5c17.8-3.5 34.9-9.5 50.1-16.1c22.9-10 42.4-21.9 54.3-30.6c31.8 11.5 67 17.9 104.1 17.9c141.4 0 256-93.1 256-208S397.4 32 256 32S0 125.1 0 240c0 45.1 17.7 86.8 47.7 120.9c-1.9 24.5-11.4 46.3-21.4 62.9M144 272a32 32 0 1 0 0-64a32 32 0 1 0 0 64m144-32a32 32 0 1 0-64 0a32 32 0 1 0 64 0m80 32a32 32 0 1 0 0-64a32 32 0 1 0 0 64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comments(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M88.2 309.1c9.8-18.3 6.8-40.8-7.5-55.8C59.4 230.9 48 204 48 176c0-63.5 63.8-128 160-128s160 64.5 160 128s-63.8 128-160 128c-13.1 0-25.8-1.3-37.8-3.6c-10.4-2-21.2-.6-30.7 4.2c-4.1 2.1-8.3 4.1-12.6 6c-16 7.2-32.9 13.5-49.9 18c2.8-4.6 5.4-9.1 7.9-13.6c1.1-1.9 2.2-3.9 3.2-5.9zM0 176c0 41.8 17.2 80.1 45.9 110.3c-.9 1.7-1.9 3.5-2.8 5.1c-10.3 18.4-22.3 36.5-36.6 52.1c-6.6 7-8.3 17.2-4.6 25.9C5.8 378.3 14.4 384 24 384c43 0 86.5-13.3 122.7-29.7c4.8-2.2 9.6-4.5 14.2-6.8c15.1 3 30.9 4.5 47.1 4.5c114.9 0 208-78.8 208-176S322.9 0 208 0S0 78.8 0 176m432 304c16.2 0 31.9-1.6 47.1-4.5c4.6 2.3 9.4 4.6 14.2 6.8C529.5 498.7 573 512 616 512c9.6 0 18.2-5.7 22-14.5c3.8-8.8 2-19-4.6-25.9c-14.2-15.6-26.2-33.7-36.6-52.1c-.9-1.7-1.9-3.4-2.8-5.1c28.8-30.3 46-68.6 46-110.4c0-94.4-87.9-171.5-198.2-175.8c4.1 15.2 6.2 31.2 6.2 47.8v.6c87.2 6.7 144 67.5 144 127.4c0 28-11.4 54.9-32.7 77.2c-14.3 15-17.3 37.6-7.5 55.8c1.1 2 2.2 4 3.2 5.9c2.5 4.5 5.2 9 7.9 13.6c-17-4.5-33.9-10.7-49.9-18c-4.3-1.9-8.5-3.9-12.6-6c-9.5-4.8-20.3-6.2-30.7-4.2c-12.1 2.4-24.7 3.6-37.8 3.6c-61.7 0-110-26.5-136.8-62.3c-16 5.4-32.8 9.4-50 11.8C279 439.8 350 480 432 480"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compass(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m306.7 69.1l-144.3 55.5c-19.4 7.5-38.5-11.6-31-31l55.5-144.3c3.3-8.5 9.9-15.1 18.4-18.4l144.3-55.5c19.4-7.5 38.5 11.6 31 31l-55.5 144.3c-3.2 8.5-9.9 15.1-18.4 18.4M288 256a32 32 0 1 0-64 0a32 32 0 1 0 64 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copy(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 336H192c-8.8 0-16-7.2-16-16V64c0-8.8 7.2-16 16-16h140.1l67.9 67.9V320c0 8.8-7.2 16-16 16m-192 48h192c35.3 0 64-28.7 64-64V115.9c0-12.7-5.1-24.9-14.1-33.9l-67.8-67.9c-9-9-21.2-14.1-33.9-14.1H192c-35.3 0-64 28.7-64 64v256c0 35.3 28.7 64 64 64M64 128c-35.3 0-64 28.7-64 64v256c0 35.3 28.7 64 64 64h192c35.3 0 64-28.7 64-64v-32h-48v32c0 8.8-7.2 16-16 16H64c-8.8 0-16-7.2-16-16V192c0-8.8 7.2-16 16-16h32v-48z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copyright(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 48a208 208 0 1 1 0 416a208 208 0 1 1 0-416m0 464a256 256 0 1 0 0-512a256 256 0 1 0 0 512m-56.6-199.4c-31.2-31.2-31.2-81.9 0-113.1s81.9-31.2 113.1 0c9.4 9.4 24.6 9.4 33.9 0s9.4-24.6 0-33.9c-50-50-131-50-181 0s-50 131 0 181s131 50 181 0c9.4-9.4 9.4-24.6 0-33.9s-24.6-9.4-33.9 0c-31.2 31.2-81.9 31.2-113.1 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCard(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 80c8.8 0 16 7.2 16 16v32H48V96c0-8.8 7.2-16 16-16zm16 144v192c0 8.8-7.2 16-16 16H64c-8.8 0-16-7.2-16-16V224zM64 32C28.7 32 0 60.7 0 96v320c0 35.3 28.7 64 64 64h448c35.3 0 64-28.7 64-64V96c0-35.3-28.7-64-64-64zm56 304c-13.3 0-24 10.7-24 24s10.7 24 24 24h48c13.3 0 24-10.7 24-24s-10.7-24-24-24zm128 0c-13.3 0-24 10.7-24 24s10.7 24 24 24h112c13.3 0 24-10.7 24-24s-10.7-24-24-24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Envelope(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 112c-8.8 0-16 7.2-16 16v22.1l172.5 141.6c20.7 17 50.4 17 71.1 0L464 150.1V128c0-8.8-7.2-16-16-16zM48 212.2V384c0 8.8 7.2 16 16 16h384c8.8 0 16-7.2 16-16V212.2L322 328.8c-38.4 31.5-93.7 31.5-132 0zM0 128c0-35.3 28.7-64 64-64h384c35.3 0 64 28.7 64 64v256c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeOpen(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M255.4 48.2c.2-.1.4-.2.6-.2s.4.1.6.2l204 145.8c2.1 1.5 3.4 3.9 3.4 6.5v13.6L291.5 355.7c-20.7 17-50.4 17-71.1 0L48 214.1v-13.6c0-2.6 1.2-5 3.4-6.5zM48 276.2l142 116.6c38.4 31.5 93.7 31.5 132 0l142-116.6V456c0 4.4-3.6 8-8 8H56c-4.4 0-8-3.6-8-8zM256 0c-10.2 0-20.2 3.2-28.5 9.1l-204 145.8C8.7 165.4 0 182.4 0 200.5V456c0 30.9 25.1 56 56 56h400c30.9 0 56-25.1 56-56V200.5c0-18.1-8.7-35.1-23.4-45.6L284.5 9.1C276.2 3.2 266.2 0 256 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M288 80c-65.2 0-118.8 29.6-159.9 67.7C89.6 183.5 63 226 49.4 256C63 286 89.6 328.5 128 364.3c41.2 38.1 94.8 67.7 160 67.7s118.8-29.6 159.9-67.7C486.4 328.5 513 286 526.6 256c-13.6-30-40.2-72.5-78.6-108.3C406.8 109.6 353.2 80 288 80M95.4 112.6C142.5 68.8 207.2 32 288 32s145.5 36.8 192.6 80.6c46.8 43.5 78.1 95.4 93 131.1c3.3 7.9 3.3 16.7 0 24.6c-14.9 35.7-46.2 87.7-93 131.1C433.5 443.2 368.8 480 288 480s-145.5-36.8-192.6-80.6C48.6 356 17.3 304 2.5 268.3c-3.3-7.9-3.3-16.7 0-24.6C17.3 208 48.6 156 95.4 112.6M288 336c44.2 0 80-35.8 80-80s-35.8-80-80-80h-2c1.3 5.1 2 10.5 2 16c0 35.3-28.7 64-64 64c-5.5 0-10.9-.7-16-2v2c0 44.2 35.8 80 80 80m0-208a128 128 0 1 1 0 256a128 128 0 1 1 0-256"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeSlash(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M38.8 5.1C28.4-3.1 13.3-1.2 5.1 9.2s-6.3 25.5 4.1 33.7l592 464c10.4 8.2 25.5 6.3 33.7-4.1s6.3-25.5-4.1-33.7l-105.2-82.4c39.6-40.6 66.4-86.1 79.9-118.4c3.3-7.9 3.3-16.7 0-24.6c-14.9-35.7-46.2-87.7-93-131.1C465.5 68.8 400.8 32 320 32c-68.2 0-125 26.3-169.3 60.8zm151 118.3C226 97.7 269.5 80 320 80c65.2 0 118.8 29.6 159.9 67.7C518.4 183.5 545 226 558.6 256c-12.6 28-36.6 66.8-70.9 100.9l-53.8-42.2c9.1-17.6 14.2-37.5 14.2-58.7c0-70.7-57.3-128-128-128c-32.2 0-61.7 11.9-84.2 31.5zm205.1 160.8l-81.5-63.9c4.2-8.5 6.6-18.2 6.6-28.3c0-5.5-.7-10.9-2-16h2c44.2 0 80 35.8 80 80c0 9.9-1.8 19.4-5.1 28.2m51.3 163.3l-41.9-33C378.8 425.4 350.7 432 320 432c-65.2 0-118.8-29.6-159.9-67.7C121.6 328.5 95 286 81.4 256c8.3-18.4 21.5-41.5 39.4-64.8l-37.7-29.7c-22.8 29.7-39.1 59.3-48.6 82.2c-3.3 7.9-3.3 16.7 0 24.6c14.9 35.7 46.2 87.7 93 131.1c47 43.8 111.7 80.6 192.5 80.6c47.8 0 89.9-12.9 126.2-32.5m-88-69.3L302 334c-23.5-5.4-43.1-21.2-53.7-42.3l-56.1-44.2c-.2 2.8-.3 5.6-.3 8.5c0 70.7 57.3 128 128 128c13.3 0 26.1-2 38.2-5.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceAngry(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 48a208 208 0 1 1 0 416a208 208 0 1 1 0-416m0 464a256 256 0 1 0 0-512a256 256 0 1 0 0 512m72.4-118.5c9.7-9 10.2-24.2 1.2-33.9c-14.3-15.3-39-31.6-73.6-31.6s-59.3 16.3-73.5 31.6c-9 9.7-8.5 24.9 1.2 33.9s24.9 8.5 33.9-1.2c7.4-7.9 20-16.4 38.5-16.4s31.1 8.5 38.5 16.4c9 9.7 24.2 10.2 33.9 1.2zM176.4 272c17.7 0 32-14.3 32-32c0-1.5-.1-3-.3-4.4l10.9 3.6c8.4 2.8 17.4-1.7 20.2-10.1s-1.7-17.4-10.1-20.2l-96-32c-8.4-2.8-17.4 1.7-20.2 10.1s1.7 17.4 10.1 20.2l30.7 10.2c-5.8 5.8-9.3 13.8-9.3 22.6c0 17.7 14.3 32 32 32m192-32c0-8.9-3.6-17-9.5-22.8l30.2-10.1c8.4-2.8 12.9-11.9 10.1-20.2S387.3 174 379 176.8l-96 32c-8.4 2.8-12.9 11.9-10.1 20.2s11.9 12.9 20.2 10.1l11.7-3.9c-.2 1.5-.3 3.1-.3 4.7c0 17.7 14.3 32 32 32s32-14.3 32-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceDizzy(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m256 32a64 64 0 1 1 0 128a64 64 0 1 1 0-128M103 135c9.4-9.4 24.6-9.4 33.9 0l23 23l23-23c9.4-9.4 24.6-9.4 33.9 0s9.4 24.6 0 33.9l-23 23l23 23c9.4 9.4 9.4 24.6 0 33.9s-24.6 9.4-33.9 0l-23-23l-23 23c-9.4 9.4-24.6 9.4-33.9 0s-9.4-24.6 0-33.9l23-23l-23-23c-9.4-9.4-9.4-24.6 0-33.9m192 0c9.4-9.4 24.6-9.4 33.9 0l23 23l23-23c9.4-9.4 24.6-9.4 33.9 0s9.4 24.6 0 33.9l-23 23l23 23c9.4 9.4 9.4 24.6 0 33.9s-24.6 9.4-33.9 0l-23-23l-23 23c-9.4 9.4-24.6 9.4-33.9 0s-9.4-24.6 0-33.9l23-23l-23-23c-9.4-9.4-9.4-24.6 0-33.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceFlushed(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 1-416 0a208 208 0 1 1 416 0M256 0a256 256 0 1 0 0 512a256 256 0 1 0 0-512m-95.6 248a24 24 0 1 0 0-48a24 24 0 1 0 0 48m216-24a24 24 0 1 0-48 0a24 24 0 1 0 48 0M192 336c-13.3 0-24 10.7-24 24s10.7 24 24 24h128c13.3 0 24-10.7 24-24s-10.7-24-24-24zm-32-160a48 48 0 1 1 0 96a48 48 0 1 1 0-96m0 128a80 80 0 1 0 0-160a80 80 0 1 0 0 160m144-80a48 48 0 1 1 96 0a48 48 0 1 1-96 0m128 0a80 80 0 1 0-160 0a80 80 0 1 0 160 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceFrown(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m174.6 128.1c-4.5 12.5-18.2 18.9-30.7 14.4s-18.9-18.2-14.4-30.7C146.9 319.4 198.9 288 256 288s109.1 31.4 126.6 79.9c4.5 12.5-2 26.2-14.4 30.7s-26.2-2-30.7-14.4c-9.3-25.7-40.3-48.2-81.5-48.2s-72.2 22.5-81.4 48.1M144.4 208a32 32 0 1 1 64 0a32 32 0 1 1-64 0m192-32a32 32 0 1 1 0 64a32 32 0 1 1 0-64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceFrownOpen(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m182.4 126.5c-12.4 5.2-26.5-4.1-21.1-16.4c16-36.6 52.4-62.1 94.8-62.1s78.8 25.6 94.8 62.1c5.4 12.3-8.7 21.6-21.1 16.4c-22.4-9.5-47.4-14.8-73.7-14.8s-51.3 5.3-73.7 14.8m-38-174.5a32 32 0 1 1 64 0a32 32 0 1 1-64 0m192-32a32 32 0 1 1 0 64a32 32 0 1 1 0-64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceGrimace(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 48a208 208 0 1 0 0 416a208 208 0 1 0 0-416m256 208a256 256 0 1 1-512 0a256 256 0 1 1 512 0m-344 64c-13.3 0-24 10.7-24 24s10.7 24 24 24h8v-48zm40 48h32v-48h-32zm96 0v-48h-32v48zm32 0h8c13.3 0 24-10.7 24-24s-10.7-24-24-24h-8zm-168-80h176c30.9 0 56 25.1 56 56s-25.1 56-56 56H168c-30.9 0-56-25.1-56-56s25.1-56 56-56m-23.6-80a32 32 0 1 1 64 0a32 32 0 1 1-64 0m192-32a32 32 0 1 1 0 64a32 32 0 1 1 0-64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceGrin(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m349.5 52.4c18.7-4.4 35.9 12 25.5 28.1c-24.6 38.1-68.7 63.5-119.1 63.5s-94.5-25.4-119.1-63.5c-10.4-16.1 6.8-32.5 25.5-28.1c28.9 6.8 60.5 10.5 93.6 10.5s64.7-3.7 93.6-10.5M144.4 208a32 32 0 1 1 64 0a32 32 0 1 1-64 0m192-32a32 32 0 1 1 0 64a32 32 0 1 1 0-64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceGrinBeam(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m349.5 52.4c18.7-4.4 35.9 12 25.5 28.1c-24.6 38.1-68.7 63.5-119.1 63.5s-94.5-25.4-119.1-63.5c-10.4-16.1 6.8-32.5 25.5-28.1c28.9 6.8 60.5 10.5 93.6 10.5s64.7-3.7 93.6-10.5m-131.9-79.6l-.2-.2c-.2-.2-.4-.5-.7-.9c-.6-.8-1.6-2-2.8-3.4c-2.5-2.8-6-6.6-10.2-10.3c-8.8-7.8-18.8-14-27.7-14s-18.9 6.2-27.7 14c-4.2 3.7-7.7 7.5-10.2 10.3c-1.2 1.4-2.2 2.6-2.8 3.4c-.3.4-.6.7-.7.9l-.2.2c-2.1 2.8-5.7 3.9-8.9 2.8s-5.5-4.1-5.5-7.6c0-17.9 6.7-35.6 16.6-48.8c9.8-13 23.9-23.2 39.4-23.2s29.6 10.2 39.4 23.2c9.9 13.2 16.6 30.9 16.6 48.8c0 3.4-2.2 6.5-5.5 7.6s-6.9 0-8.9-2.8m160 0l-.2-.2c-.2-.2-.4-.5-.7-.9c-.6-.8-1.6-2-2.8-3.4c-2.5-2.8-6-6.6-10.2-10.3c-8.8-7.8-18.8-14-27.7-14s-18.9 6.2-27.7 14c-4.2 3.7-7.7 7.5-10.2 10.3c-1.2 1.4-2.2 2.6-2.8 3.4c-.3.4-.6.7-.7.9l-.2.2c-2.1 2.8-5.7 3.9-8.9 2.8s-5.5-4.1-5.5-7.6c0-17.9 6.7-35.6 16.6-48.8c9.8-13 23.9-23.2 39.4-23.2s29.6 10.2 39.4 23.2c9.9 13.2 16.6 30.9 16.6 48.8c0 3.4-2.2 6.5-5.5 7.6s-6.9 0-8.9-2.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceGrinBeamSweat(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M476.8 126.3c20.3-5.5 35.2-23.6 35.2-45.3c0-20-28.6-60.4-41.6-77.7c-3.2-4.4-9.6-4.4-12.8 0c-9.5 12.6-27.1 37.2-36 57.5l-.9 2.1C417.8 69.7 416 76 416 81c0 26 21.5 47 48 47c4.4 0 8.7-.6 12.8-1.7m-81.4-85.1C355.3 15.2 307.4 0 256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256c0-35.8-7.3-69.9-20.6-100.8c-8.6 3.1-17.8 4.8-27.4 4.8c-8.9 0-17.6-1.5-25.7-4.2C454.7 185.5 464 219.7 464 256c0 114.9-93.1 208-208 208S48 370.9 48 256S141.1 48 256 48c48.7 0 93.4 16.7 128.9 44.7c-.6-3.8-.9-7.7-.9-11.7c0-11.4 3.8-22.4 7.1-30.5c1.3-3.1 2.7-6.2 4.3-9.3M375 336.5c10.4-16.1-6.8-32.5-25.5-28.1c-28.9 6.8-60.5 10.5-93.6 10.5s-64.7-3.7-93.6-10.5c-18.7-4.4-35.9 12-25.5 28.1c24.6 38.1 68.7 63.5 119.1 63.5s94.5-25.4 119.1-63.5M217.6 228.8c2.1 2.8 5.7 3.9 8.9 2.8s5.5-4.1 5.5-7.6c0-17.9-6.7-35.6-16.6-48.8c-9.8-13-23.9-23.2-39.4-23.2s-29.6 10.2-39.4 23.2c-9.9 13.2-16.6 30.9-16.6 48.8c0 3.4 2.2 6.5 5.5 7.6s6.9 0 8.9-2.8l.2-.2c.2-.2.4-.5.7-.9c.6-.8 1.6-2 2.8-3.4c2.5-2.8 6-6.6 10.2-10.3c8.8-7.8 18.8-14 27.7-14s18.9 6.2 27.7 14c4.2 3.7 7.7 7.5 10.2 10.3c1.2 1.4 2.2 2.6 2.8 3.4c.3.4.6.7.7.9zm160 0c2.1 2.8 5.7 3.9 8.9 2.8s5.5-4.1 5.5-7.6c0-17.9-6.7-35.6-16.6-48.8c-9.8-13-23.9-23.2-39.4-23.2s-29.6 10.2-39.4 23.2c-9.9 13.2-16.6 30.9-16.6 48.8c0 3.4 2.2 6.5 5.5 7.6s6.9 0 8.9-2.8l.2-.2c.2-.2.4-.5.7-.9c.6-.8 1.6-2 2.8-3.4c2.5-2.8 6-6.6 10.2-10.3c8.8-7.8 18.8-14 27.7-14s18.9 6.2 27.7 14c4.2 3.7 7.7 7.5 10.2 10.3c1.2 1.4 2.2 2.6 2.8 3.4c.3.4.6.7.7.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceGrinHearts(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m349.5 52.4c18.7-4.4 35.9 12 25.5 28.1c-24.6 38.1-68.7 63.5-119.1 63.5s-94.5-25.4-119.1-63.5c-10.4-16.1 6.8-32.5 25.5-28.1c28.9 6.8 60.5 10.5 93.6 10.5s64.7-3.7 93.6-10.5M215.3 137.1c17.8 4.8 28.4 23.1 23.6 40.8l-17.4 65c-2.3 8.5-11.1 13.6-19.6 11.3l-65.1-17.4c-17.8-4.8-28.4-23.1-23.6-40.8s23.1-28.4 40.8-23.6l16.1 4.3l4.3-16.1c4.8-17.8 23.1-28.4 40.8-23.6zm122.3 23.6l4.3 16.1l16.1-4.3c17.8-4.8 36.1 5.8 40.8 23.6s-5.8 36.1-23.6 40.8l-65.1 17.4c-8.5 2.3-17.3-2.8-19.6-11.3l-17.4-65c-4.8-17.8 5.8-36.1 23.6-40.8s36.1 5.8 40.9 23.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceGrinSquint(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m349.5 52.4c18.7-4.4 35.9 12 25.5 28.1c-24.6 38.1-68.7 63.5-119.1 63.5s-94.5-25.4-119.1-63.5c-10.4-16.1 6.8-32.5 25.5-28.1c28.9 6.8 60.5 10.5 93.6 10.5s64.7-3.7 93.6-10.5m-216-161.7l89.9 47.9c10.7 5.7 10.7 21.1 0 26.8l-89.9 47.9c-7.9 4.2-17.5-1.5-17.5-10.5c0-2.8 1-5.5 2.8-7.6l36-43.2l-36-43.2c-1.8-2.1-2.8-4.8-2.8-7.6c0-9 9.6-14.7 17.5-10.5M396 157.1c0 2.8-1 5.5-2.8 7.6l-36 43.2l36 43.2c1.8 2.1 2.8 4.8 2.8 7.6c0 9-9.6 14.7-17.5 10.5l-89.9-47.9c-10.7-5.7-10.7-21.1 0-26.8l89.9-47.9c7.9-4.2 17.5 1.5 17.5 10.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceGrinSquintTears(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M426.8 14.2c19.2-19.2 50.7-18.8 70.3.7s20 51 .7 70.3c-14.8 14.8-65.7 23.6-88.3 26.7c-5.6.9-10.3-3.9-9.5-9.5c3.3-22.5 12-73.4 26.8-88.2M75 75c83.2-83.3 209.5-97.2 307.2-41.8c-1.5 4.8-2.9 9.6-4.1 14.3c-3.1 12.2-5.5 24.6-7.3 35c-80.8-53.6-190.7-44.8-261.9 26.4c-71.2 71.2-80 181.1-26.4 261.9c-10.5 1.8-22.9 4.2-35 7.3c-4.7 1.2-9.5 2.5-14.3 4.1C-22.2 284.5-8.2 158.2 75 75m389.6 58.9c4.7-1.2 9.5-2.5 14.3-4.1c55.3 97.7 41.3 224-41.9 307.2c-83.2 83.2-209.5 97.2-307.2 41.8c1.5-4.8 2.8-9.6 4-14.3c3.1-12.2 5.5-24.6 7.3-35c80.8 53.6 190.7 44.8 261.9-26.4c71.2-71.2 80-181.1 26.4-261.9c10.5-1.8 22.9-4.2 35-7.3zm-105.4 93c10.1-16.3 33.9-16.9 37.9 1.9c9.5 44.4-3.7 93.5-39.3 129.1s-84.8 48.8-129.1 39.3c-18.7-4-18.2-27.8-1.9-37.9c25.2-15.7 50.2-35.4 73.6-58.8s43.1-48.4 58.8-73.6M92 265.3l97.4-29.7c11.6-3.5 22.5 7.3 19 19L178.7 352c-2.6 8.6-13.4 11.3-19.8 4.9c-2-2-3.2-4.6-3.4-7.3l-5.1-56.1l-56.1-5.1c-2.8-.3-5.4-1.5-7.3-3.4c-6.3-6.3-3.6-17.2 4.9-19.8zM285 87.1c2 2 3.2 4.6 3.4 7.3l5.1 56.1l56.1 5.1c2.8.3 5.4 1.5 7.3 3.4c6.3 6.3 3.6 17.2-4.9 19.8l-97.4 29.7c-11.6 3.5-22.5-7.3-19-19L265.3 92c2.6-8.6 13.4-11.3 19.8-4.9zm-270.1 410c-19.6-19.6-20-51-.7-70.3c14.8-14.8 65.6-23.6 88.2-26.7c5.6-.9 10.3 3.9 9.5 9.5c-3.2 22.5-11.9 73.5-26.7 88.3c-19.2 19.1-50.7 18.7-70.3-.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceGrinStars(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 48a208 208 0 1 1 0 416a208 208 0 1 1 0-416m0 464a256 256 0 1 0 0-512a256 256 0 1 0 0 512m-72.8-379.4a7.934 7.934 0 0 0-14.4 0l-16.6 34.7l-38.1 5c-3.1.4-5.6 2.5-6.6 5.5s-.1 6.2 2.1 8.3l27.9 26.5l-7 37.8c-.6 3 .7 6.1 3.2 7.9s5.8 2 8.5.6l33.8-18.4l33.8 18.3c2.7 1.5 6 1.3 8.5-.6s3.7-4.9 3.2-7.9l-7-37.8l27.9-26.5c2.2-2.1 3.1-5.3 2.1-8.3s-3.5-5.1-6.6-5.5l-38.1-5l-16.6-34.7zm160 0a7.934 7.934 0 0 0-14.4 0l-16.6 34.7l-38.1 5c-3.1.4-5.6 2.5-6.6 5.5s-.1 6.2 2.1 8.3l27.9 26.5l-7 37.8c-.6 3 .7 6.1 3.2 7.9s5.8 2 8.5.6l33.8-18.4l33.8 18.3c2.7 1.5 6 1.3 8.5-.6s3.7-4.9 3.2-7.9l-7-37.8l27.9-26.5c2.2-2.1 3.1-5.3 2.1-8.3s-3.5-5.1-6.6-5.5l-38.1-5l-16.6-34.7zm6.3 175.8c-28.9 6.8-60.5 10.5-93.6 10.5s-64.7-3.7-93.6-10.5c-18.7-4.4-35.9 12-25.5 28.1c24.6 38.1 68.7 63.5 119.1 63.5s94.5-25.4 119.1-63.5c10.4-16.1-6.8-32.5-25.5-28.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceGrinTears(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M516.1 325.5c1 3 2.1 6 3.3 8.9c3.3 8.1 8.4 18.5 16.5 26.6c3.9 3.9 8.2 7.4 12.7 10.3C506.4 454.8 419.9 512 320 512S133.6 454.8 91.4 371.4c4.5-2.9 8.7-6.3 12.7-10.3c8.1-8.1 13.2-18.6 16.5-26.6c1.2-2.9 2.3-5.9 3.3-8.9C152.5 406.2 229.5 464 320 464s167.5-57.8 196.1-138.5M320 48c-101.4 0-185.8 72.5-204.3 168.5c-6.7-3.1-14.3-4.3-22.3-3.1c-6.8.9-16.2 2.4-26.6 4.4C85.3 94.5 191.6 0 320 0s234.7 94.5 253.2 217.7c-10.3-2-19.8-3.5-26.6-4.4c-8-1.2-15.7.1-22.3 3.1C505.8 120.5 421.4 48 320 48M78.5 341.1C60 356.7 32 355.5 14.3 337.7c-18.7-18.7-19.1-48.8-.7-67.2c8.6-8.6 30.1-15.1 50.5-19.6c13-2.8 25.5-4.8 33.9-6c5.4-.8 9.9 3.7 9 9c-3.1 21.5-11.4 70.2-25.5 84.4c-.9 1-1.9 1.8-2.9 2.7zm483 0c-.8-.6-1.5-1.3-2.3-2c-.2-.2-.5-.4-.7-.7c-14.1-14.1-22.5-62.9-25.5-84.4c-.8-5.4 3.7-9.9 9-9c1 .1 2.2.3 3.3.5c8.2 1.2 19.2 3 30.6 5.5c20.4 4.4 41.9 10.9 50.5 19.6c18.4 18.4 18 48.5-.7 67.2c-17.7 17.7-45.7 19-64.2 3.4zM439 336.5c-24.6 38.1-68.7 63.5-119.1 63.5s-94.5-25.4-119.1-63.5c-10.4-16.1 6.8-32.5 25.5-28.1c28.9 6.8 60.5 10.5 93.6 10.5s64.7-3.7 93.6-10.5c18.7-4.4 35.9 12 25.5 28.1M281.6 228.8l-.2-.2c-.2-.2-.4-.5-.7-.9c-.6-.8-1.6-2-2.8-3.4c-2.5-2.8-6-6.6-10.2-10.3c-8.8-7.8-18.8-14-27.7-14s-18.9 6.2-27.7 14c-4.2 3.7-7.7 7.5-10.2 10.3c-1.2 1.4-2.2 2.6-2.8 3.4c-.3.4-.6.7-.7.9l-.2.2c-2.1 2.8-5.7 3.9-8.9 2.8s-5.5-4.1-5.5-7.6c0-17.9 6.7-35.6 16.6-48.8c9.8-13 23.9-23.2 39.4-23.2s29.6 10.2 39.4 23.2c9.9 13.2 16.6 30.9 16.6 48.8c0 3.4-2.2 6.5-5.5 7.6s-6.9 0-8.9-2.8m160 0l-.2-.2c-.2-.2-.4-.5-.7-.9c-.6-.8-1.6-2-2.8-3.4c-2.5-2.8-6-6.6-10.2-10.3c-8.8-7.8-18.8-14-27.7-14s-18.9 6.2-27.7 14c-4.2 3.7-7.7 7.5-10.2 10.3c-1.2 1.4-2.2 2.6-2.8 3.4c-.3.4-.6.7-.7.9l-.2.2c-2.1 2.8-5.7 3.9-8.9 2.8s-5.5-4.1-5.5-7.6c0-17.9 6.7-35.6 16.6-48.8c9.8-13 23.9-23.2 39.4-23.2s29.6 10.2 39.4 23.2c9.9 13.2 16.6 30.9 16.6 48.8c0 3.4-2.2 6.5-5.5 7.6s-6.9 0-8.9-2.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceGrinTongue(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256c0-114.9-93.1-208-208-208S48 141.1 48 256c0 81.7 47.1 152.4 115.7 186.4c-2.4-8.4-3.7-17.3-3.7-26.4v-52.4c-8.9-8-16.7-17.1-23.1-27.1c-10.4-16.1 6.8-32.5 25.5-28.1c28.9 6.8 60.5 10.5 93.6 10.5s64.7-3.7 93.6-10.5c18.7-4.4 35.9 12 25.5 28.1c-6.4 9.9-14.2 19-23 27V416c0 9.2-1.3 18-3.7 26.4C416.9 408.4 464 337.7 464 256M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m176.4-80a32 32 0 1 1 0 64a32 32 0 1 1 0-64m128 32a32 32 0 1 1 64 0a32 32 0 1 1-64 0M320 416v-37.4c0-14.7-11.9-26.6-26.6-26.6h-2c-11.3 0-21.1 7.9-23.6 18.9c-2.8 12.6-20.8 12.6-23.6 0c-2.5-11.1-12.3-18.9-23.6-18.9h-2c-14.7 0-26.6 11.9-26.6 26.6V416c0 35.3 28.7 64 64 64s64-28.7 64-64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceGrinTongueSquint(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256c0-114.9-93.1-208-208-208S48 141.1 48 256c0 81.7 47.1 152.4 115.7 186.4c-2.4-8.4-3.7-17.3-3.7-26.4v-23.3c-24-17.5-43.1-41.4-54.8-69.2c-5-11.8 7-22.5 19.3-18.7c39.7 12.2 84.5 19 131.8 19s92.1-6.8 131.8-19c12.3-3.8 24.3 6.9 19.3 18.7c-11.8 28-31.1 52-55.4 69.6V416c0 9.2-1.3 18-3.7 26.4C416.9 408.4 464 337.7 464 256M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m116-98.9c0-9 9.6-14.7 17.5-10.5l89.9 47.9c10.7 5.7 10.7 21.1 0 26.8l-89.9 47.9c-7.9 4.2-17.5-1.5-17.5-10.5c0-2.8 1-5.5 2.8-7.6l36-43.2l-36-43.2c-1.8-2.1-2.8-4.8-2.8-7.6m262.5-10.5c7.9-4.2 17.5 1.5 17.5 10.5c0 2.8-1 5.5-2.8 7.6l-36 43.2l36 43.2c1.8 2.1 2.8 4.8 2.8 7.6c0 9-9.6 14.7-17.5 10.5l-89.9-47.9c-10.7-5.7-10.7-21.1 0-26.8zM320 416v-37.4c0-14.7-11.9-26.6-26.6-26.6h-2c-11.3 0-21.1 7.9-23.6 18.9c-2.8 12.6-20.8 12.6-23.6 0c-2.5-11.1-12.3-18.9-23.6-18.9h-2c-14.7 0-26.6 11.9-26.6 26.6V416c0 35.3 28.7 64 64 64s64-28.7 64-64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceGrinTongueWink(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M348.3 442.4c2.4-8.4 3.7-17.3 3.7-26.4v-52.5c8.8-8 16.6-17.1 23-27c10.4-16.1-6.8-32.5-25.5-28.1c-28.9 6.8-60.5 10.5-93.6 10.5s-64.7-3.7-93.6-10.5c-18.7-4.4-35.9 12-25.5 28.1c6.5 10 14.3 19.1 23.1 27.1V416c0 9.2 1.3 18 3.7 26.4C95.1 408.4 48 337.7 48 256c0-114.9 93.1-208 208-208s208 93.1 208 208c0 81.7-47.1 152.4-115.7 186.4M256 512a256 256 0 1 0 0-512a256 256 0 1 0 0 512m-96.4-292c10.6 0 19.9 3.8 25.4 9.7c7.6 8.1 20.2 8.5 28.3.9s8.5-20.2.9-28.3C199.7 186.8 179 180 159.6 180s-40.1 6.8-54.6 22.3c-7.6 8.1-7.1 20.7.9 28.3s20.7 7.1 28.3-.9c5.5-5.8 14.8-9.7 25.4-9.7m176.7 12a24 24 0 1 0 0-48a24 24 0 1 0 0 48m-.4-72a48 48 0 1 1 0 96a48 48 0 1 1 0-96m0 128a80 80 0 1 0 0-160a80 80 0 1 0 0 160M320 416c0 35.3-28.7 64-64 64s-64-28.7-64-64v-37.4c0-14.7 11.9-26.6 26.6-26.6h2c11.3 0 21.1 7.9 23.6 18.9c2.8 12.6 20.8 12.6 23.6 0c2.5-11.1 12.3-18.9 23.6-18.9h2c14.7 0 26.6 11.9 26.6 26.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceGrinWide(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m349.5 52.4c18.7-4.4 35.9 12 25.5 28.1c-24.6 38.1-68.7 63.5-119.1 63.5s-94.5-25.4-119.1-63.5c-10.4-16.1 6.8-32.5 25.5-28.1c28.9 6.8 60.5 10.5 93.6 10.5s64.7-3.7 93.6-10.5M224 192c0 35.3-14.3 64-32 64s-32-28.7-32-64s14.3-64 32-64s32 28.7 32 64m96 64c-17.7 0-32-28.7-32-64s14.3-64 32-64s32 28.7 32 64s-14.3 64-32 64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceGrinWink(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m349.5 52.4c18.7-4.4 35.9 12 25.5 28.1c-24.6 38.1-68.7 63.5-119.1 63.5s-94.5-25.4-119.1-63.5c-10.4-16.1 6.8-32.5 25.5-28.1c28.9 6.8 60.5 10.5 93.6 10.5s64.7-3.7 93.6-10.5M144.4 208a32 32 0 1 1 64 0a32 32 0 1 1-64 0m165.8 21.7c-7.6 8.1-20.2 8.5-28.3.9s-8.5-20.2-.9-28.3c14.5-15.5 35.2-22.3 54.6-22.3s40.1 6.8 54.6 22.3c7.6 8.1 7.1 20.7-.9 28.3s-20.7 7.1-28.3-.9c-5.5-5.8-14.8-9.7-25.4-9.7s-19.9 3.8-25.4 9.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceKiss(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m304.7 25.7c4.3 5.1 7.3 11.4 7.3 18.3s-3.1 13.2-7.3 18.3c-4.3 5.2-10.1 9.7-16.7 13.4c-2.7 1.5-5.7 3-8.7 4.3c3.1 1.3 6 2.7 8.7 4.3c6.6 3.7 12.5 8.2 16.7 13.4c4.3 5.1 7.3 11.4 7.3 18.3s-3.1 13.2-7.3 18.3c-4.3 5.2-10.1 9.7-16.7 13.4c-13.3 7.4-30.6 12.3-48 12.3c-3.6 0-6.8-2.5-7.7-6s.6-7.2 3.8-9l.2-.1c.2-.1.5-.3.9-.5c.8-.5 2-1.2 3.4-2.1c2.8-1.9 6.5-4.5 10.2-7.6c3.7-3.1 7.2-6.6 9.6-10.1c2.5-3.5 3.5-6.4 3.5-8.6s-1-5-3.5-8.6c-2.5-3.5-5.9-6.9-9.6-10.1c-3.7-3.1-7.4-5.7-10.2-7.6c-1.4-.9-2.6-1.6-3.4-2.1l-.8-.5l-.1-.1l-.2-.1c-2.5-1.4-4.1-4.1-4.1-7s1.6-5.6 4.1-7l.2-.1c.2-.1.5-.3.9-.5c.8-.5 2-1.2 3.4-2.1c2.8-1.9 6.5-4.5 10.2-7.6c3.7-3.1 7.2-6.6 9.6-10.1c2.5-3.5 3.5-6.4 3.5-8.6s-1-5-3.5-8.6c-2.5-3.5-5.9-6.9-9.6-10.1c-3.7-3.1-7.4-5.7-10.2-7.6c-1.4-.9-2.6-1.6-3.4-2.1c-.4-.2-.7-.4-.9-.5l-.2-.1c-3.2-1.8-4.7-5.5-3.8-9s4.1-6 7.7-6c17.4 0 34.7 4.9 47.9 12.3c6.6 3.7 12.5 8.2 16.7 13.4zM144.4 208a32 32 0 1 1 64 0a32 32 0 1 1-64 0m192-32a32 32 0 1 1 0 64a32 32 0 1 1 0-64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceKissBeam(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m304.7 41.7c4.3 5.1 7.3 11.4 7.3 18.3s-3.1 13.2-7.3 18.3c-4.3 5.2-10.1 9.7-16.7 13.4c-2.7 1.5-5.7 3-8.7 4.3c3.1 1.3 6 2.7 8.7 4.3c6.6 3.7 12.5 8.2 16.7 13.4c4.3 5.1 7.3 11.4 7.3 18.3s-3.1 13.2-7.3 18.3c-4.3 5.2-10.1 9.7-16.7 13.4c-13.3 7.4-30.6 12.3-48 12.3c-3.6 0-6.8-2.5-7.7-6s.6-7.2 3.8-9l.2-.1c.2-.1.5-.3.9-.5c.8-.5 2-1.2 3.4-2.1c2.8-1.9 6.5-4.5 10.2-7.6c3.7-3.1 7.2-6.6 9.6-10.1c2.5-3.5 3.5-6.4 3.5-8.6s-1-5-3.5-8.6c-2.5-3.5-5.9-6.9-9.6-10.1c-3.7-3.1-7.4-5.7-10.2-7.6c-1.4-.9-2.6-1.6-3.4-2.1c-.4-.2-.7-.4-.9-.5l-.2-.1c-2.5-1.4-4.1-4.1-4.1-7s1.6-5.6 4.1-7l.2-.1c.2-.1.5-.3.9-.5c.8-.5 2-1.2 3.4-2.1c2.8-1.9 6.5-4.5 10.2-7.6c3.7-3.1 7.2-6.6 9.6-10.1c2.5-3.5 3.5-6.4 3.5-8.6s-1-5-3.5-8.6c-2.5-3.5-5.9-6.9-9.6-10.1c-3.7-3.1-7.4-5.7-10.2-7.6c-1.4-.9-2.6-1.6-3.4-2.1c-.4-.2-.7-.4-.9-.5l-.2-.1c-3.2-1.8-4.7-5.5-3.8-9s4.1-6 7.7-6c17.4 0 34.7 4.9 47.9 12.3c6.6 3.7 12.5 8.2 16.7 13.4zm-87.1-68.9l-.2-.2c-.2-.2-.4-.5-.7-.9c-.6-.8-1.6-2-2.8-3.4c-2.5-2.8-6-6.6-10.2-10.3c-8.8-7.8-18.8-14-27.7-14s-18.9 6.2-27.7 14c-4.2 3.7-7.7 7.5-10.2 10.3c-1.2 1.4-2.2 2.6-2.8 3.4c-.3.4-.6.7-.7.9l-.2.2c-2.1 2.8-5.7 3.9-8.9 2.8s-5.5-4.1-5.5-7.6c0-17.9 6.7-35.6 16.6-48.8c9.8-13 23.9-23.2 39.4-23.2s29.6 10.2 39.4 23.2c9.9 13.2 16.6 30.9 16.6 48.8c0 3.4-2.2 6.5-5.5 7.6s-6.9 0-8.9-2.8m160 0l-.2-.2c-.2-.2-.4-.5-.7-.9c-.6-.8-1.6-2-2.8-3.4c-2.5-2.8-6-6.6-10.2-10.3c-8.8-7.8-18.8-14-27.7-14s-18.9 6.2-27.7 14c-4.2 3.7-7.7 7.5-10.2 10.3c-1.2 1.4-2.2 2.6-2.8 3.4c-.3.4-.6.7-.7.9l-.2.2c-2.1 2.8-5.7 3.9-8.9 2.8s-5.5-4.1-5.5-7.6c0-17.9 6.7-35.6 16.6-48.8c9.8-13 23.9-23.2 39.4-23.2s29.6 10.2 39.4 23.2c9.9 13.2 16.6 30.9 16.6 48.8c0 3.4-2.2 6.5-5.5 7.6s-6.9 0-8.9-2.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceKissWinkHeart(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M338.9 446.8c-25.4 11-53.4 17.2-82.9 17.2c-114.9 0-208-93.1-208-208S141.1 48 256 48s208 93.1 208 208c0 22.4-3.5 43.9-10.1 64.1c3.1 4.5 5.7 9.4 7.8 14.6c12.7-1.6 25.1.4 36.2 5c9.1-26.2 14-54.4 14-83.7C512 114.6 397.4 0 256 0S0 114.6 0 256s114.6 256 256 256c35.4 0 69.1-7.2 99.7-20.2c-4.8-5.5-8.5-12.2-10.4-19.7l-6.5-25.3zM296 316c0-6.9-3.1-13.2-7.3-18.3c-4.3-5.2-10.1-9.7-16.7-13.4c-13.3-7.4-30.6-12.3-48-12.3c-3.6 0-6.8 2.5-7.7 6s.6 7.2 3.8 9l.2.1c.2.1.5.3.9.5c.8.5 2 1.2 3.4 2.1c2.8 1.9 6.5 4.5 10.2 7.6c3.7 3.1 7.2 6.6 9.6 10.1c2.5 3.5 3.5 6.4 3.5 8.6s-1 5-3.5 8.6c-2.5 3.5-5.9 6.9-9.6 10.1c-3.7 3.1-7.4 5.7-10.2 7.6c-1.4.9-2.6 1.6-3.4 2.1c-.4.2-.7.4-.9.5l-.2.1c-2.5 1.4-4.1 4.1-4.1 7s1.6 5.6 4.1 7l.2.1c.2.1.5.3.9.5c.8.5 2 1.2 3.4 2.1c2.8 1.9 6.5 4.5 10.2 7.6c3.7 3.1 7.2 6.6 9.6 10.1c2.5 3.5 3.5 6.4 3.5 8.6s-1 5-3.5 8.6c-2.5 3.5-5.9 6.9-9.6 10.1c-3.7 3.1-7.4 5.7-10.2 7.6c-1.4.9-2.6 1.6-3.4 2.1c-.4.2-.7.4-.9.5l-.2.1c-3.2 1.8-4.7 5.5-3.8 9s4.1 6 7.7 6c17.4 0 34.7-4.9 47.9-12.3c6.6-3.7 12.5-8.2 16.7-13.4c4.3-5.1 7.3-11.4 7.3-18.3s-3.1-13.2-7.3-18.3c-4.3-5.2-10.1-9.7-16.7-13.4c-2.7-1.5-5.7-3-8.7-4.3c3.1-1.3 6-2.7 8.7-4.3c6.6-3.7 12.5-8.2 16.7-13.4c4.3-5.1 7.3-11.4 7.3-18.3zm-119.6-76a32 32 0 1 0 0-64a32 32 0 1 0 0 64m159.3-20c10.6 0 19.9 3.8 25.4 9.7c7.6 8.1 20.2 8.5 28.3.9s8.5-20.2.9-28.3C375.7 186.8 355 180 335.6 180s-40.1 6.8-54.6 22.3c-7.6 8.1-7.1 20.7.9 28.3s20.7 7.1 28.3-.9c5.5-5.8 14.8-9.7 25.4-9.7zM434 352.3c-6-23.2-28.8-37-51.1-30.8s-35.4 30.1-29.5 53.4l22.9 89.3c2.2 8.7 11.2 13.9 19.8 11.4l84.9-23.8c22.2-6.2 35.4-30.1 29.5-53.4s-28.8-37-51.1-30.8l-20.2 5.6l-5.4-21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceLaugh(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m130.7 57.9c-4.2-13.6 7.1-25.9 21.3-25.9h212.5c14.2 0 25.5 12.4 21.3 25.9C369 368.4 318.2 408 258.2 408s-110.8-39.6-127.5-94.1M144.4 192a32 32 0 1 1 64 0a32 32 0 1 1-64 0m192-32a32 32 0 1 1 0 64a32 32 0 1 1 0-64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceLaughBeam(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m130.7 57.9c-4.2-13.6 7.1-25.9 21.3-25.9h212.5c14.2 0 25.5 12.4 21.3 25.9C369 368.4 318.2 408 258.2 408s-110.8-39.6-127.5-94.1m86.9-85.1l-.2-.2c-.2-.2-.4-.5-.7-.9c-.6-.8-1.6-2-2.8-3.4c-2.5-2.8-6-6.6-10.2-10.3c-8.8-7.8-18.8-14-27.7-14s-18.9 6.2-27.7 14c-4.2 3.7-7.7 7.5-10.2 10.3c-1.2 1.4-2.2 2.6-2.8 3.4c-.3.4-.6.7-.7.9l-.2.2c-2.1 2.8-5.7 3.9-8.9 2.8s-5.5-4.1-5.5-7.6c0-17.9 6.7-35.6 16.6-48.8c9.8-13 23.9-23.2 39.4-23.2s29.6 10.2 39.4 23.2c9.9 13.2 16.6 30.9 16.6 48.8c0 3.4-2.2 6.5-5.5 7.6s-6.9 0-8.9-2.8m160 0l-.2-.2c-.2-.2-.4-.5-.7-.9c-.6-.8-1.6-2-2.8-3.4c-2.5-2.8-6-6.6-10.2-10.3c-8.8-7.8-18.8-14-27.7-14s-18.9 6.2-27.7 14c-4.2 3.7-7.7 7.5-10.2 10.3c-1.2 1.4-2.2 2.6-2.8 3.4c-.3.4-.6.7-.7.9l-.2.2c-2.1 2.8-5.7 3.9-8.9 2.8s-5.5-4.1-5.5-7.6c0-17.9 6.7-35.6 16.6-48.8c9.8-13 23.9-23.2 39.4-23.2s29.6 10.2 39.4 23.2c9.9 13.2 16.6 30.9 16.6 48.8c0 3.4-2.2 6.5-5.5 7.6s-6.9 0-8.9-2.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceLaughSquint(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m130.7 57.9c-4.2-13.6 7.1-25.9 21.3-25.9h212.5c14.2 0 25.5 12.4 21.3 25.9C369 368.4 318.2 408 258.2 408s-110.8-39.6-127.5-94.1m2.8-183.3l89.9 47.9c10.7 5.7 10.7 21.1 0 26.8l-89.9 47.9c-7.9 4.2-17.5-1.5-17.5-10.5c0-2.8 1-5.5 2.8-7.6l36-43.2l-36-43.2c-1.8-2.1-2.8-4.8-2.8-7.6c0-9 9.6-14.7 17.5-10.5M396 141.1c0 2.8-1 5.5-2.8 7.6l-36 43.2l36 43.2c1.8 2.1 2.8 4.8 2.8 7.6c0 9-9.6 14.7-17.5 10.5l-89.9-47.9c-10.7-5.7-10.7-21.1 0-26.8l89.9-47.9c7.9-4.2 17.5 1.5 17.5 10.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceLaughWink(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m130.7 57.9c-4.2-13.6 7.1-25.9 21.3-25.9h212.5c14.2 0 25.5 12.4 21.3 25.9C369 368.4 318.2 408 258.2 408s-110.8-39.6-127.5-94.1M144.4 192a32 32 0 1 1 64 0a32 32 0 1 1-64 0m165.8 21.7c-7.6 8.1-20.2 8.5-28.3.9s-8.5-20.2-.9-28.3c14.5-15.5 35.2-22.3 54.6-22.3s40.1 6.8 54.6 22.3c7.6 8.1 7.1 20.7-.9 28.3s-20.7 7.1-28.3-.9c-5.5-5.8-14.8-9.7-25.4-9.7s-19.9 3.8-25.4 9.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceMeh(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 1-416 0a208 208 0 1 1 416 0M256 0a256 256 0 1 0 0 512a256 256 0 1 0 0-512m-79.6 240a32 32 0 1 0 0-64a32 32 0 1 0 0 64m192-32a32 32 0 1 0-64 0a32 32 0 1 0 64 0M184 328c-13.3 0-24 10.7-24 24s10.7 24 24 24h144c13.3 0 24-10.7 24-24s-10.7-24-24-24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceMehBlank(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 48a208 208 0 1 0 0 416a208 208 0 1 0 0-416m256 208a256 256 0 1 1-512 0a256 256 0 1 1 512 0m-367.6-48a32 32 0 1 1 64 0a32 32 0 1 1-64 0m192-32a32 32 0 1 1 0 64a32 32 0 1 1 0-64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceRollingEyes(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 48a208 208 0 1 1 0 416a208 208 0 1 1 0-416m0 464a256 256 0 1 0 0-512a256 256 0 1 0 0 512m-88-136c0 13.3 10.7 24 24 24h128c13.3 0 24-10.7 24-24s-10.7-24-24-24H192c-13.3 0-24 10.7-24 24m-8-104c-26.5 0-48-21.5-48-48c0-14.3 6.3-27.2 16.2-36c-.2 1.3-.2 2.6-.2 4c0 17.7 14.3 32 32 32s32-14.3 32-32c0-1.4-.1-2.7-.2-4c10 8.8 16.2 21.7 16.2 36c0 26.5-21.5 48-48 48m0 32a80 80 0 1 0 0-160a80 80 0 1 0 0 160m192-32c-26.5 0-48-21.5-48-48c0-14.3 6.3-27.2 16.2-36c-.2 1.3-.2 2.6-.2 4c0 17.7 14.3 32 32 32s32-14.3 32-32c0-1.4-.1-2.7-.2-4c10 8.8 16.2 21.7 16.2 36c0 26.5-21.5 48-48 48m0 32a80 80 0 1 0 0-160a80 80 0 1 0 0 160"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceSadCry(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 406.1V288c0-13.3-10.7-24-24-24s-24 10.7-24 24v152.6c-28.7 15-61.4 23.4-96 23.4s-67.3-8.5-96-23.4V288c0-13.3-10.7-24-24-24s-24 10.7-24 24v118.1C72.6 368.2 48 315 48 256c0-114.9 93.1-208 208-208s208 93.1 208 208c0 59-24.6 112.2-64 150.1M256 512a256 256 0 1 0 0-512a256 256 0 1 0 0 512m-96.4-292c10.6 0 19.9 3.8 25.4 9.7c7.6 8.1 20.2 8.5 28.3.9s8.5-20.2.9-28.3C199.7 186.8 179 180 159.6 180s-40.1 6.8-54.6 22.3c-7.6 8.1-7.1 20.7.9 28.3s20.7 7.1 28.3-.9c5.5-5.8 14.8-9.7 25.4-9.7m166.6 9.7c5.5-5.8 14.8-9.7 25.4-9.7s19.9 3.8 25.4 9.7c7.6 8.1 20.2 8.5 28.3.9s8.5-20.2.9-28.3C391.7 186.8 371 180 351.6 180s-40.1 6.8-54.6 22.3c-7.6 8.1-7.1 20.7.9 28.3s20.7 7.1 28.3-.9M208 320v32c0 26.5 21.5 48 48 48s48-21.5 48-48v-32c0-26.5-21.5-48-48-48s-48 21.5-48 48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceSadTear(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M175.9 448c-35-.1-65.5-22.6-76-54.6C67.6 356.8 48 308.7 48 256c0-114.9 93.1-208 208-208s208 93.1 208 208s-93.1 208-208 208c-28.4 0-55.5-5.7-80.1-16M0 256a256 256 0 1 0 512 0a256 256 0 1 0-512 0m128 113c0 26 21.5 47 48 47s48-21 48-47c0-20-28.4-60.4-41.6-77.7c-3.2-4.4-9.6-4.4-12.8 0c-13 17.3-41.6 57.7-41.6 77.7m128-65c-13.3 0-24 10.7-24 24s10.7 24 24 24c30.7 0 58.7 11.5 80 30.6c9.9 8.8 25 8 33.9-1.9s8-25-1.9-33.9c-29.7-26.6-69-42.8-112-42.8m47.6-96a32 32 0 1 0 64 0a32 32 0 1 0-64 0m-128 32a32 32 0 1 0 0-64a32 32 0 1 0 0 64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceSmile(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m177.6 62.1c15.2 16.4 41.2 33.9 78.4 33.9s63.2-17.5 78.4-33.9c9-9.7 24.2-10.4 33.9-1.4s10.4 24.2 1.4 33.9c-22 23.8-60 49.4-113.6 49.4s-91.7-25.5-113.6-49.4c-9-9.7-8.4-24.9 1.4-33.9s24.9-8.4 33.9 1.4zM144.4 208a32 32 0 1 1 64 0a32 32 0 1 1-64 0m192-32a32 32 0 1 1 0 64a32 32 0 1 1 0-64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceSmileBeam(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m177.6 62.1c15.2 16.4 41.2 33.9 78.4 33.9s63.2-17.5 78.4-33.9c9-9.7 24.2-10.4 33.9-1.4s10.4 24.2 1.4 33.9c-22 23.8-60 49.4-113.6 49.4s-91.7-25.5-113.6-49.4c-9-9.7-8.4-24.9 1.4-33.9s24.9-8.4 33.9 1.4zm40-89.3l-.2-.2c-.2-.2-.4-.5-.7-.9c-.6-.8-1.6-2-2.8-3.4c-2.5-2.8-6-6.6-10.2-10.3c-8.8-7.8-18.8-14-27.7-14s-18.9 6.2-27.7 14c-4.2 3.7-7.7 7.5-10.2 10.3c-1.2 1.4-2.2 2.6-2.8 3.4c-.3.4-.6.7-.7.9l-.2.2c-2.1 2.8-5.7 3.9-8.9 2.8s-5.5-4.1-5.5-7.6c0-17.9 6.7-35.6 16.6-48.8c9.8-13 23.9-23.2 39.4-23.2s29.6 10.2 39.4 23.2c9.9 13.2 16.6 30.9 16.6 48.8c0 3.4-2.2 6.5-5.5 7.6s-6.9 0-8.9-2.8m160 0l-.2-.2c-.2-.2-.4-.5-.7-.9c-.6-.8-1.6-2-2.8-3.4c-2.5-2.8-6-6.6-10.2-10.3c-8.8-7.8-18.8-14-27.7-14s-18.9 6.2-27.7 14c-4.2 3.7-7.7 7.5-10.2 10.3c-1.2 1.4-2.2 2.6-2.8 3.4c-.3.4-.6.7-.7.9l-.2.2c-2.1 2.8-5.7 3.9-8.9 2.8s-5.5-4.1-5.5-7.6c0-17.9 6.7-35.6 16.6-48.8c9.8-13 23.9-23.2 39.4-23.2s29.6 10.2 39.4 23.2c9.9 13.2 16.6 30.9 16.6 48.8c0 3.4-2.2 6.5-5.5 7.6s-6.9 0-8.9-2.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceSmileWink(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m177.6 62.1c15.2 16.4 41.2 33.9 78.4 33.9s63.2-17.5 78.4-33.9c9-9.7 24.2-10.4 33.9-1.4s10.4 24.2 1.4 33.9c-22 23.8-60 49.4-113.6 49.4s-91.7-25.5-113.6-49.4c-9-9.7-8.4-24.9 1.4-33.9s24.9-8.4 33.9 1.4zM144.4 208a32 32 0 1 1 64 0a32 32 0 1 1-64 0m165.8 21.7c-7.6 8.1-20.2 8.5-28.3.9s-8.5-20.2-.9-28.3c14.5-15.5 35.2-22.3 54.6-22.3s40.1 6.8 54.6 22.3c7.6 8.1 7.1 20.7-.9 28.3s-20.7 7.1-28.3-.9c-5.5-5.8-14.8-9.7-25.4-9.7s-19.9 3.8-25.4 9.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceSurprise(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m176.4-80a32 32 0 1 1 0 64a32 32 0 1 1 0-64m128 32a32 32 0 1 1 64 0a32 32 0 1 1-64 0M256 288a64 64 0 1 1 0 128a64 64 0 1 1 0-128"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceTired(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0m176.5 64.3c19.6-18.2 47.3-32.3 79.5-32.3s59.9 14.1 79.5 32.3c19 17.8 32.5 41.7 32.5 63.7c0 5.4-2.7 10.4-7.2 13.4s-10.2 3.4-15.2 1.3l-17.2-7.5c-22.8-10-47.5-15.1-72.4-15.1s-49.6 5.2-72.4 15.1l-17.2 7.5c-4.9 2.2-10.7 1.7-15.2-1.3s-7.2-8-7.2-13.4c0-22 13.5-45.9 32.5-63.7m-43-173.6l89.9 47.9c10.7 5.7 10.7 21.1 0 26.8l-89.9 47.9c-7.9 4.2-17.5-1.5-17.5-10.5c0-2.8 1-5.5 2.8-7.6l36-43.2l-36-43.2c-1.8-2.1-2.8-4.8-2.8-7.6c0-9 9.6-14.7 17.5-10.5M396 157.1c0 2.8-1 5.5-2.8 7.6l-36 43.2l36 43.2c1.8 2.1 2.8 4.8 2.8 7.6c0 9-9.6 14.7-17.5 10.5l-89.9-47.9c-10.7-5.7-10.7-21.1 0-26.8l89.9-47.9c7.9-4.2 17.5 1.5 17.5 10.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func File(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M320 464c8.8 0 16-7.2 16-16V160h-80c-17.7 0-32-14.3-32-32V48H64c-8.8 0-16 7.2-16 16v384c0 8.8 7.2 16 16 16zM0 64C0 28.7 28.7 0 64 0h165.5c17 0 33.3 6.7 45.3 18.7l90.5 90.5c12 12 18.7 28.3 18.7 45.3V448c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileAudio(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 464h256c8.8 0 16-7.2 16-16V160h-80c-17.7 0-32-14.3-32-32V48H64c-8.8 0-16 7.2-16 16v384c0 8.8 7.2 16 16 16M0 64C0 28.7 28.7 0 64 0h165.5c17 0 33.3 6.7 45.3 18.7l90.5 90.5c12 12 18.7 28.3 18.7 45.3V448c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64zm192 208v128c0 6.5-3.9 12.3-9.9 14.8s-12.9 1.1-17.4-3.5L129.4 376H112c-8.8 0-16-7.2-16-16v-48c0-8.8 7.2-16 16-16h17.4l35.3-35.3c4.6-4.6 11.5-5.9 17.4-3.5s9.9 8.3 9.9 14.8m85.8-4c11.6 20 18.2 43.3 18.2 68s-6.6 48-18.2 68c-6.6 11.5-21.3 15.4-32.8 8.8s-15.4-21.3-8.8-32.8c7.5-12.9 11.8-27.9 11.8-44s-4.3-31.1-11.8-44c-6.6-11.5-2.7-26.2 8.8-32.8s26.2-2.7 32.8 8.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileCode(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 464c-8.8 0-16-7.2-16-16V64c0-8.8 7.2-16 16-16h160v80c0 17.7 14.3 32 32 32h80v288c0 8.8-7.2 16-16 16zM64 0C28.7 0 0 28.7 0 64v384c0 35.3 28.7 64 64 64h256c35.3 0 64-28.7 64-64V154.5c0-17-6.7-33.3-18.7-45.3l-90.6-90.5C262.7 6.7 246.5 0 229.5 0zm97 289c9.4-9.4 9.4-24.6 0-33.9s-24.6-9.4-33.9 0L79 303c-9.4 9.4-9.4 24.6 0 33.9l48 48c9.4 9.4 24.6 9.4 33.9 0s9.4-24.6 0-33.9l-31-31l31-31zm96-34c-9.4-9.4-24.6-9.4-33.9 0s-9.4 24.6 0 33.9l31 31l-31 31c-9.4 9.4-9.4 24.6 0 33.9s24.6 9.4 33.9 0l48-48c9.4-9.4 9.4-24.6 0-33.9l-48-48z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileExcel(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M48 448V64c0-8.8 7.2-16 16-16h160v80c0 17.7 14.3 32 32 32h80v288c0 8.8-7.2 16-16 16H64c-8.8 0-16-7.2-16-16M64 0C28.7 0 0 28.7 0 64v384c0 35.3 28.7 64 64 64h256c35.3 0 64-28.7 64-64V154.5c0-17-6.7-33.3-18.7-45.3l-90.6-90.5C262.7 6.7 246.5 0 229.5 0zm90.9 233.3c-8.1-10.5-23.2-12.3-33.7-4.2s-12.3 23.2-4.2 33.7l44.6 57.2l-44.5 57.3c-8.1 10.5-6.3 25.5 4.2 33.7s25.5 6.3 33.7-4.2l37-47.7l37.1 47.6c8.1 10.5 23.2 12.3 33.7 4.2s12.3-23.2 4.2-33.7L222.4 320l44.5-57.3c8.1-10.5 6.3-25.5-4.2-33.7s-25.5-6.3-33.7 4.2l-37 47.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileImage(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 464c-8.8 0-16-7.2-16-16V64c0-8.8 7.2-16 16-16h160v80c0 17.7 14.3 32 32 32h80v288c0 8.8-7.2 16-16 16zM64 0C28.7 0 0 28.7 0 64v384c0 35.3 28.7 64 64 64h256c35.3 0 64-28.7 64-64V154.5c0-17-6.7-33.3-18.7-45.3l-90.6-90.5C262.7 6.7 246.5 0 229.5 0zm96 256a32 32 0 1 0-64 0a32 32 0 1 0 64 0m69.2 46.9c-3-4.3-7.9-6.9-13.2-6.9s-10.2 2.6-13.2 6.9l-41.3 59.7l-11.9-19.1c-2.9-4.7-8.1-7.5-13.6-7.5s-10.6 2.8-13.6 7.5l-40 64c-3.1 4.9-3.2 11.1-.4 16.2s8.2 8.2 14 8.2h192c6 0 11.4-3.3 14.2-8.6s2.4-11.6-1-16.5l-72-104z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileLines(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 464c-8.8 0-16-7.2-16-16V64c0-8.8 7.2-16 16-16h160v80c0 17.7 14.3 32 32 32h80v288c0 8.8-7.2 16-16 16zM64 0C28.7 0 0 28.7 0 64v384c0 35.3 28.7 64 64 64h256c35.3 0 64-28.7 64-64V154.5c0-17-6.7-33.3-18.7-45.3l-90.6-90.5C262.7 6.7 246.5 0 229.5 0zm56 256c-13.3 0-24 10.7-24 24s10.7 24 24 24h144c13.3 0 24-10.7 24-24s-10.7-24-24-24zm0 96c-13.3 0-24 10.7-24 24s10.7 24 24 24h144c13.3 0 24-10.7 24-24s-10.7-24-24-24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePdf(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 464h48v48H64c-35.3 0-64-28.7-64-64V64C0 28.7 28.7 0 64 0h165.5c17 0 33.3 6.7 45.3 18.7l90.5 90.5c12 12 18.7 28.3 18.7 45.3V304h-48V160h-80c-17.7 0-32-14.3-32-32V48H64c-8.8 0-16 7.2-16 16v384c0 8.8 7.2 16 16 16m112-112h32c30.9 0 56 25.1 56 56s-25.1 56-56 56h-16v32c0 8.8-7.2 16-16 16s-16-7.2-16-16V368c0-8.8 7.2-16 16-16m32 80c13.3 0 24-10.7 24-24s-10.7-24-24-24h-16v48zm96-80h32c26.5 0 48 21.5 48 48v64c0 26.5-21.5 48-48 48h-32c-8.8 0-16-7.2-16-16V368c0-8.8 7.2-16 16-16m32 128c8.8 0 16-7.2 16-16v-64c0-8.8-7.2-16-16-16h-16v96zm80-112c0-8.8 7.2-16 16-16h48c8.8 0 16 7.2 16 16s-7.2 16-16 16h-32v32h32c8.8 0 16 7.2 16 16s-7.2 16-16 16h-32v48c0 8.8-7.2 16-16 16s-16-7.2-16-16v-64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePowerpoint(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 464c-8.8 0-16-7.2-16-16V64c0-8.8 7.2-16 16-16h160v80c0 17.7 14.3 32 32 32h80v288c0 8.8-7.2 16-16 16zM64 0C28.7 0 0 28.7 0 64v384c0 35.3 28.7 64 64 64h256c35.3 0 64-28.7 64-64V154.5c0-17-6.7-33.3-18.7-45.3l-90.6-90.5C262.7 6.7 246.5 0 229.5 0zm72 208c-13.3 0-24 10.7-24 24v160c0 13.3 10.7 24 24 24s24-10.7 24-24v-32h44c42 0 76-34 76-76s-34-76-76-76zm68 104h-44v-56h44c15.5 0 28 12.5 28 28s-12.5 28-28 28"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileVideo(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M320 464c8.8 0 16-7.2 16-16V160h-80c-17.7 0-32-14.3-32-32V48H64c-8.8 0-16 7.2-16 16v384c0 8.8 7.2 16 16 16zM0 64C0 28.7 28.7 0 64 0h165.5c17 0 33.3 6.7 45.3 18.7l90.5 90.5c12 12 18.7 28.3 18.7 45.3V448c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64zm80 224c0-17.7 14.3-32 32-32h96c17.7 0 32 14.3 32 32v16l44.9-29.9c2-1.3 4.4-2.1 6.8-2.1c6.8 0 12.3 5.5 12.3 12.3v103.4c0 6.8-5.5 12.3-12.3 12.3c-2.4 0-4.8-.7-6.8-2.1L240 368v16c0 17.7-14.3 32-32 32h-96c-17.7 0-32-14.3-32-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileWord(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M48 448V64c0-8.8 7.2-16 16-16h160v80c0 17.7 14.3 32 32 32h80v288c0 8.8-7.2 16-16 16H64c-8.8 0-16-7.2-16-16M64 0C28.7 0 0 28.7 0 64v384c0 35.3 28.7 64 64 64h256c35.3 0 64-28.7 64-64V154.5c0-17-6.7-33.3-18.7-45.3l-90.6-90.5C262.7 6.7 246.5 0 229.5 0zm55 241.1c-3.8-12.7-17.2-19.9-29.9-16.1S69.2 242.2 73 254.9l48 160c3 10.2 12.4 17.1 23 17.1s19.9-7 23-17.1l25-83.4l25 83.4c3 10.2 12.4 17.1 23 17.1s19.9-7 23-17.1l48-160c3.8-12.7-3.4-26.1-16.1-29.9s-26.1 3.4-29.9 16.1l-25 83.4l-25-83.4c-3-10.2-12.4-17.1-23-17.1s-19.9 7-23 17.1l-25 83.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileZipper(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 464c-8.8 0-16-7.2-16-16V64c0-8.8 7.2-16 16-16h48c0 8.8 7.2 16 16 16h32c8.8 0 16-7.2 16-16h48v80c0 17.7 14.3 32 32 32h80v288c0 8.8-7.2 16-16 16zM64 0C28.7 0 0 28.7 0 64v384c0 35.3 28.7 64 64 64h256c35.3 0 64-28.7 64-64V154.5c0-17-6.7-33.3-18.7-45.3l-90.6-90.5C262.7 6.7 246.5 0 229.5 0zm48 112c0 8.8 7.2 16 16 16h32c8.8 0 16-7.2 16-16s-7.2-16-16-16h-32c-8.8 0-16 7.2-16 16m0 64c0 8.8 7.2 16 16 16h32c8.8 0 16-7.2 16-16s-7.2-16-16-16h-32c-8.8 0-16 7.2-16 16m-6.3 71.8l-23.6 88.1c-1.4 5.4-2.1 10.9-2.1 16.4c0 35.2 28.8 63.7 64 63.7s64-28.5 64-63.7c0-5.5-.7-11.1-2.1-16.4l-23.5-88.2c-3.7-14-16.4-23.8-30.9-23.8h-14.9c-14.5 0-27.2 9.7-30.9 23.8zM128 336h32c8.8 0 16 7.2 16 16s-7.2 16-16 16h-32c-8.8 0-16-7.2-16-16s7.2-16 16-16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M48 24C48 10.7 37.3 0 24 0S0 10.7 0 24v464c0 13.3 10.7 24 24 24s24-10.7 24-24V388l80.3-20.1c41.1-10.3 84.6-5.5 122.5 13.4c44.2 22.1 95.5 24.8 141.7 7.4l34.7-13c12.5-4.7 20.8-16.6 20.8-30V66.1c0-23-24.2-38-44.8-27.7l-9.6 4.8c-46.3 23.2-100.8 23.2-147.1 0c-35.1-17.6-75.4-22-113.5-12.5L48 52zm0 77.5l96.6-24.2c27-6.7 55.5-3.6 80.4 8.8c54.9 27.4 118.7 29.7 175 6.8v241.8l-24.4 9.1c-33.7 12.6-71.2 10.7-103.4-5.4c-48.2-24.1-103.3-30.1-155.6-17.1L48 338.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FloppyDisk(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M48 96v320c0 8.8 7.2 16 16 16h320c8.8 0 16-7.2 16-16V170.5c0-4.2-1.7-8.3-4.7-11.3l33.9-33.9c12 12 18.7 28.3 18.7 45.3V416c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64V96c0-35.3 28.7-64 64-64h245.5c17 0 33.3 6.7 45.3 18.7l74.5 74.5l-33.9 33.9l-74.6-74.4l-.8-.8V184c0 13.3-10.7 24-24 24H104c-13.3 0-24-10.7-24-24V80H64c-8.8 0-16 7.2-16 16m80-16v80h144V80zm32 240a64 64 0 1 1 128 0a64 64 0 1 1-128 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 96c0-35.3 28.7-64 64-64h132.1c19.1 0 37.4 7.6 50.9 21.1L289.9 96H448c35.3 0 64 28.7 64 64v256c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64zm64-16c-8.8 0-16 7.2-16 16v320c0 8.8 7.2 16 16 16h384c8.8 0 16-7.2 16-16V160c0-8.8-7.2-16-16-16H286.6c-10.6 0-20.8-4.2-28.3-11.7L213.1 87c-4.5-4.5-10.6-7-17-7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderClosed(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 32C28.7 32 0 60.7 0 96v320c0 35.3 28.7 64 64 64h384c35.3 0 64-28.7 64-64V160c0-35.3-28.7-64-64-64H289.9L247 53.1A71.983 71.983 0 0 0 196.1 32zM48 96c0-8.8 7.2-16 16-16h132.1c6.4 0 12.5 2.5 17 7l45.3 45.3c7.5 7.5 17.7 11.7 28.3 11.7H448c8.8 0 16 7.2 16 16v32H48zm0 144h416v176c0 8.8-7.2 16-16 16H64c-8.8 0-16-7.2-16-16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpen(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 480h48c11.4 0 21.9-6 27.6-15.9l112-192c5.8-9.9 5.8-22.1.1-32.1s-16.2-16-27.7-16H144c-11.4 0-21.9 6-27.6 15.9L48 357.1V96c0-8.8 7.2-16 16-16h117.5c4.2 0 8.3 1.7 11.3 4.7l26.5 26.5c21 21 49.5 32.8 79.2 32.8H416c8.8 0 16 7.2 16 16v32h48v-32c0-35.3-28.7-64-64-64H298.5c-17 0-33.3-6.7-45.3-18.7l-26.5-26.6c-12-12-28.3-18.7-45.3-18.7H64C28.7 32 0 60.7 0 96v320c0 35.3 28.7 64 64 64h23.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FontAwesome(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M48 56c0-13.3-10.7-24-24-24S0 42.7 0 56v400c0 13.3 10.7 24 24 24s24-10.7 24-24V124.2l12.5-2.4c16.7-3.2 31.5-8.5 44.2-13.1c3.7-1.3 7.1-2.6 10.4-3.7c15.2-5.2 30.4-9.1 51.2-9.1c25.6 0 43 6 63.5 13.3l.5.2c20.9 7.4 44.8 15.9 79.1 15.9c32.4 0 53.7-6.8 90.5-19.6v237.2l-9.5 3.3c-41.5 14.4-55.2 19.2-81 19.2c-25.7 0-43.1-6-63.6-13.3l-.6-.2c-20.8-7.4-44.8-15.8-79-15.8c-16.8 0-31 2-43.9 5c-12.9 3-20.9 16-17.9 28.9s16 20.9 28.9 17.9c9.6-2.2 20.1-3.7 32.9-3.7c25.6 0 43 6 63.5 13.3l.5.2c20.9 7.4 44.8 15.9 79.1 15.9c34.4 0 56.4-7.7 97.8-22.2c7.5-2.6 15.5-5.4 24.4-8.5l16.2-5.5v-339l-31.5 10.9c-9.7 3.3-18.2 6.3-25.7 8.9c-41.5 14.4-55.2 19.2-81 19.2c-25.7 0-43.1-6-63.6-13.3l-.6-.2c-20.8-7.4-44.8-15.8-79-15.8c-27.8 0-48.5 5.5-66.6 11.6c-4.9 1.7-9.3 3.3-13.6 4.8c-11.9 4.3-22 7.9-34.7 10.3l-3.4.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Futbol(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m435.4 361.3l-89.7-6c-5.2-.3-10.3 1.1-14.5 4.2s-7.2 7.4-8.4 12.5l-22 87.2c-14.4 3.2-29.4 4.8-44.8 4.8s-30.3-1.7-44.8-4.8l-22-87.2c-1.3-5-4.3-9.4-8.4-12.5s-9.3-4.5-14.5-4.2l-89.7 6C61.7 335.9 51.9 307 49 276.2l76-47.9c4.4-2.8 7.6-7 9.2-11.9s1.4-10.2-.5-15L100.4 118c19.9-22.4 44.6-40.5 72.4-52.7l69.1 57.6c4 3.3 9 5.1 14.1 5.1s10.2-1.8 14.1-5.1l69.1-57.6c27.8 12.2 52.5 30.3 72.4 52.7l-33.4 83.4c-1.9 4.8-2.1 10.1-.5 15s4.9 9.1 9.2 11.9l76.1 47.9c-3 30.8-12.7 59.7-27.6 85.1M256 48h.9h-1.8zM56.7 196.2c.9-3 1.9-6.1 2.9-9.1zM132 423l3.8 2.7c-1.3-.9-2.5-1.8-3.8-2.7m248.1-.1c-1.3 1-2.7 2-4 2.9zm75.2-226.6l-3-9.2c1.1 3 2.1 6.1 3 9.2M256 512a256 256 0 1 0 0-512a256 256 0 1 0 0 512m14.1-325.7c-8.4-6.1-19.8-6.1-28.2 0L194 221c-8.4 6.1-11.9 16.9-8.7 26.8l18.3 56.3c3.2 9.9 12.4 16.6 22.8 16.6h59.2c10.4 0 19.6-6.7 22.8-16.6l18.3-56.3c3.2-9.9-.3-20.7-8.7-26.8l-47.9-34.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gem(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m168.5 72l87.5 93l87.5-93zm215.4 27.1L311.5 176h129zm50 124.9H78.1L256 420.3zM71.5 176h129l-72.4-76.9zm434.3 40.1l-232 256c-4.5 5-11 7.9-17.8 7.9s-13.2-2.9-17.8-7.9l-232-256c-7.7-8.5-8.3-21.2-1.5-30.4l112-152c4.5-6.1 11.7-9.8 19.3-9.8h240c7.6 0 14.8 3.6 19.3 9.8l112 152c6.8 9.2 6.1 21.9-1.5 30.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hand(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0c-25.3 0-47.2 14.7-57.6 36c-7-2.6-14.5-4-22.4-4c-35.3 0-64 28.7-64 64v165.5l-2.7-2.7c-25-25-65.5-25-90.5 0s-25 65.5 0 90.5l87.7 87.7c48 48 113.1 75 181 75H304c1.5 0 3-.1 4.5-.4c91.7-6.2 165-79.4 171.1-171.1c.3-1.5.4-3 .4-4.5V160c0-35.3-28.7-64-64-64c-5.5 0-10.9.7-16 2v-2c0-35.3-28.7-64-64-64c-7.9 0-15.4 1.4-22.4 4C303.2 14.7 281.3 0 256 0m-16 96.1V64c0-8.8 7.2-16 16-16s16 7.2 16 16v168c0 13.3 10.7 24 24 24s24-10.7 24-24V95.9c0-8.8 7.2-16 16-16s16 7.2 16 16v136c0 13.3 10.7 24 24 24s24-10.7 24-24V160c0-8.8 7.2-16 16-16s16 7.2 16 16v172.9c-.1.6-.1 1.3-.2 1.9c-3.4 69.7-59.3 125.6-129 129c-.6 0-1.3.1-1.9.2h-13.4c-55.2 0-108.1-21.9-147.1-60.9l-87.7-87.8c-6.2-6.2-6.2-16.4 0-22.6s16.4-6.2 22.6 0l43.7 43.7c6.9 6.9 17.2 8.9 26.2 5.2s14.8-12.5 14.8-22.2V96c0-8.8 7.2-16 16-16s16 7.1 16 15.9V232c0 13.3 10.7 24 24 24s24-10.7 24-24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandBackFist(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M144 64c0-8.8 7.2-16 16-16s16 7.2 16 16c0 9.1 5.1 17.4 13.3 21.5s17.9 3.2 25.1-2.3c2.7-2 6-3.2 9.6-3.2c8.8 0 16 7.2 16 16c0 9.1 5.1 17.4 13.3 21.5s17.9 3.2 25.1-2.3c2.7-2 6-3.2 9.6-3.2c8.8 0 16 7.2 16 16c0 9.1 5.1 17.4 13.3 21.5s17.9 3.2 25.1-2.3c2.7-2 6-3.2 9.6-3.2c8.8 0 16 7.2 16 16v104c0 31.3-20 58-48 67.9c-9.6 3.4-16 12.5-16 22.6V488c0 13.3 10.7 24 24 24s24-10.7 24-24V370.2c38-20.1 64-60.1 64-106.2V160c0-35.3-28.7-64-64-64c-2.8 0-5.6.2-8.3.5A63.831 63.831 0 0 0 288 64c-2.8 0-5.6.2-8.3.5A63.831 63.831 0 0 0 224 32c-2.8 0-5.6.2-8.3.5A63.831 63.831 0 0 0 160 0c-35.3 0-64 28.7-64 64v64.3c-11.7 7.4-22.5 16.4-32 26.9l17.8 16.1L64 155.2l-9.4 10.5C40 181.8 32 202.8 32 224.6v12.8c0 49.6 24.2 96.1 64.8 124.5l13.8-19.7l-13.8 19.7l8.9 6.2c6.9 4.8 14.4 8.6 22.3 11.3V488c0 13.3 10.7 24 24 24s24-10.7 24-24V359.9c0-12.6-9.8-23.1-22.4-23.9c-7.3-.5-14.3-2.9-20.3-7.1l-13.1 18.7l13.1-18.7l-8.9-6.2C96.6 303.1 80 271.3 80 237.4v-12.8c0-9.9 3.7-19.4 10.3-26.8l9.4-10.5c3.8-4.2 7.9-8.1 12.3-11.6V208c0 8.8 7.2 16 16 16s16-7.2 16-16v-80z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandLizard(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M72 112c-13.3 0-24 10.7-24 24s10.7 24 24 24h168c35.3 0 64 28.7 64 64s-28.7 64-64 64H136c-13.3 0-24 10.7-24 24s10.7 24 24 24h152c4.5 0 8.9 1.3 12.7 3.6l64 40c7 4.4 11.3 12.1 11.3 20.4v24c0 13.3-10.7 24-24 24s-24-10.7-24-24v-10.7L281.1 384H136c-39.8 0-72-32.2-72-72s32.2-72 72-72h104c8.8 0 16-7.2 16-16s-7.2-16-16-16H72c-39.8 0-72-32.2-72-72s32.2-72 72-72h209.6c46.7 0 90.9 21.5 119.7 58.3l78.4 100.1c20.9 26.7 32.3 59.7 32.3 93.7V424c0 13.3-10.7 24-24 24s-24-10.7-24-24V316.1c0-23.2-7.8-45.8-22.1-64.1l-78.4-100.1c-19.7-25.2-49.9-39.9-81.9-39.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandPeace(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M250.8 1.4c-35.2-3.7-66.6 21.8-70.3 57L174 119l-17.3-49.4C145 36.3 108.4 18.8 75.1 30.5s-50.9 48.3-39.2 81.6l52.8 150.1A79.77 79.77 0 0 0 64 320v24c0 92.8 75.2 168 168 168h48c92.8 0 168-75.2 168-168V224c0-35.3-28.7-64-64-64c-7.9 0-15.4 1.4-22.4 4c-10.4-21.3-32.3-36-57.6-36h-2.2l5.9-56.3c3.7-35.2-21.8-66.6-57-70.3zm-.2 155.4C243.9 166.9 240 179 240 192v50c-5.1-1.3-10.5-2-16-2h-7.4l-5.4-15.3l17-161.3c.9-8.8 8.8-15.2 17.6-14.2S261 58 260 66.8l-9.5 90.1zM111.4 85.6L165.7 240H144c-4 0-8 .3-11.9.9L81.2 96.2c-2.9-8.3 1.5-17.5 9.8-20.4s17.5 1.5 20.4 9.8M288 192c0-8.8 7.2-16 16-16s16 7.2 16 16v48c0 8.8-7.2 16-16 16s-16-7.2-16-16zm38.4 108c10.4 21.3 32.3 36 57.6 36c5.5 0 10.9-.7 16-2v10c0 66.3-53.7 120-120 120h-48c-66.3 0-120-53.7-120-120v-24c0-17.7 14.3-32 32-32h80c8.8 0 16 7.2 16 16s-7.2 16-16 16h-40c-13.3 0-24 10.7-24 24s10.7 24 24 24h40c35.3 0 64-28.7 64-64v-2c5.1 1.3 10.5 2 16 2c7.9 0 15.4-1.4 22.4-4m73.6-28c0 8.8-7.2 16-16 16s-16-7.2-16-16v-48c0-8.8 7.2-16 16-16s16 7.2 16 16v32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandPointDown(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 448V270.4c5.2 1 10.5 1.6 16 1.6h16v176c0 8.8-7.2 16-16 16s-16-7.2-16-16m16-224c-17.7 0-32-14.3-32-32v-24c0-66.3 53.7-120 120-120h48c52.5 0 97.1 33.7 113.4 80.7c-3.1-.5-6.2-.7-9.4-.7c-20 0-37.9 9.2-49.7 23.6c-9-4.9-19.4-7.6-30.3-7.6c-15.1 0-29 5.3-40 14c-11-8.8-24.9-14-40-14h-40c-13.3 0-24 10.7-24 24s10.7 24 24 24h40c8.8 0 16 7.2 16 16s-7.2 16-16 16h-40zM0 192c0 18 6 34.6 16 48v208c0 35.3 28.7 64 64 64s64-28.7 64-64v-82c5.1 1.3 10.5 2 16 2c25.3 0 47.2-14.7 57.6-36c7 2.6 14.5 4 22.4 4c20 0 37.9-9.2 49.7-23.6c9 4.9 19.4 7.6 30.3 7.6c35.3 0 64-28.7 64-64v-88C384 75.2 308.8 0 216 0h-48C75.2 0 0 75.2 0 168zm336 64c0 8.8-7.2 16-16 16s-16-7.2-16-16v-64c0-8.8 7.2-16 16-16s16 7.2 16 16zm-176 16c5.5 0 10.9-.7 16-2v34c0 8.8-7.2 16-16 16s-16-7.2-16-16v-32zm64-24v-40c0-8.8 7.2-16 16-16s16 7.2 16 16v64c0 8.8-7.2 16-16 16s-16-7.2-16-16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandPointLeft(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 128h177.6c-1 5.2-1.6 10.5-1.6 16v16H64c-8.8 0-16-7.2-16-16s7.2-16 16-16m224 16c0-17.7 14.3-32 32-32h24c66.3 0 120 53.7 120 120v48c0 52.5-33.7 97.1-80.7 113.4c.5-3.1.7-6.2.7-9.4c0-20-9.2-37.9-23.6-49.7c4.9-9 7.6-19.4 7.6-30.3c0-15.1-5.3-29-14-40c8.8-11 14-24.9 14-40v-40c0-13.3-10.7-24-24-24s-24 10.7-24 24v40c0 8.8-7.2 16-16 16s-16-7.2-16-16v-40zm32-80c-18 0-34.6 6-48 16H64c-35.3 0-64 28.7-64 64s28.7 64 64 64h82c-1.3 5.1-2 10.5-2 16c0 25.3 14.7 47.2 36 57.6c-2.6 7-4 14.5-4 22.4c0 20 9.2 37.9 23.6 49.7c-4.9 9-7.6 19.4-7.6 30.3c0 35.3 28.7 64 64 64h88c92.8 0 168-75.2 168-168v-48c0-92.8-75.2-168-168-168zm-64 336c-8.8 0-16-7.2-16-16s7.2-16 16-16h64c8.8 0 16 7.2 16 16s-7.2 16-16 16zm-16-176c0 5.5.7 10.9 2 16h-34c-8.8 0-16-7.2-16-16s7.2-16 16-16h32zm24 64h40c8.8 0 16 7.2 16 16s-7.2 16-16 16h-64c-8.8 0-16-7.2-16-16s7.2-16 16-16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandPointRight(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 128H270.4c1 5.2 1.6 10.5 1.6 16v16h176c8.8 0 16-7.2 16-16s-7.2-16-16-16m-224 16c0-17.7-14.3-32-32-32h-24c-66.3 0-120 53.7-120 120v48c0 52.5 33.7 97.1 80.7 113.4c-.5-3.1-.7-6.2-.7-9.4c0-20 9.2-37.9 23.6-49.7c-4.9-9-7.6-19.4-7.6-30.3c0-15.1 5.3-29 14-40c-8.8-11-14-24.9-14-40v-40c0-13.3 10.7-24 24-24s24 10.7 24 24v40c0 8.8 7.2 16 16 16s16-7.2 16-16v-40zm-32-80c18 0 34.6 6 48 16h208c35.3 0 64 28.7 64 64s-28.7 64-64 64h-82c1.3 5.1 2 10.5 2 16c0 25.3-14.7 47.2-36 57.6c2.6 7 4 14.5 4 22.4c0 20-9.2 37.9-23.6 49.7c4.9 9 7.6 19.4 7.6 30.3c0 35.3-28.7 64-64 64h-88C75.2 448 0 372.8 0 280v-48C0 139.2 75.2 64 168 64zm64 336c8.8 0 16-7.2 16-16s-7.2-16-16-16h-64c-8.8 0-16 7.2-16 16s7.2 16 16 16zm16-176c0 5.5-.7 10.9-2 16h34c8.8 0 16-7.2 16-16s-7.2-16-16-16h-32zm-24 64h-40c-8.8 0-16 7.2-16 16s7.2 16 16 16h64c8.8 0 16-7.2 16-16s-7.2-16-16-16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandPointUp(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 64v177.6c5.2-1 10.5-1.6 16-1.6h16V64c0-8.8-7.2-16-16-16s-16 7.2-16 16m16 224c-17.7 0-32 14.3-32 32v24c0 66.3 53.7 120 120 120h48c52.5 0 97.1-33.7 113.4-80.7c-3.1.5-6.2.7-9.4.7c-20 0-37.9-9.2-49.7-23.6c-9 4.9-19.4 7.6-30.3 7.6c-15.1 0-29-5.3-40-14c-11 8.8-24.9 14-40 14h-40c-13.3 0-24-10.7-24-24s10.7-24 24-24h40c8.8 0 16-7.2 16-16s-7.2-16-16-16h-40zM0 320c0-18 6-34.6 16-48V64C16 28.7 44.7 0 80 0s64 28.7 64 64v82c5.1-1.3 10.5-2 16-2c25.3 0 47.2 14.7 57.6 36c7-2.6 14.5-4 22.4-4c20 0 37.9 9.2 49.7 23.6c9-4.9 19.4-7.6 30.3-7.6c35.3 0 64 28.7 64 64v88c0 92.8-75.2 168-168 168h-48C75.2 512 0 436.8 0 344zm336-64c0-8.8-7.2-16-16-16s-16 7.2-16 16v64c0 8.8 7.2 16 16 16s16-7.2 16-16zm-176-16c5.5 0 10.9.7 16 2v-34c0-8.8-7.2-16-16-16s-16 7.2-16 16v32zm64 24v40c0 8.8 7.2 16 16 16s16-7.2 16-16v-64c0-8.8-7.2-16-16-16s-16 7.2-16 16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandPointer(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M160 64c0-8.8 7.2-16 16-16s16 7.2 16 16v136c0 10.3 6.6 19.5 16.4 22.8s20.6-.1 26.8-8.3c3-3.9 7.6-6.4 12.8-6.4c8.8 0 16 7.2 16 16c0 10.3 6.6 19.5 16.4 22.8s20.6-.1 26.8-8.3c3-3.9 7.6-6.4 12.8-6.4c7.8 0 14.3 5.6 15.7 13c1.6 8.2 7.3 15.1 15.1 18s16.7 1.6 23.3-3.6c2.7-2.1 6.1-3.4 9.9-3.4c8.8 0 16 7.2 16 16V392c0 39.8-32.2 72-72 72H211.4c-37.4 0-72.4-18.7-93.2-49.9L50.7 312.9c-4.9-7.4-2.9-17.3 4.4-22.2s17.3-2.9 22.2 4.4l38.7 58.1c5.9 8.8 16.8 12.7 26.9 9.7s17-12.4 17-23V64zm16-64c-35.3 0-64 28.7-64 64v197.7c-20.8-23.7-56.5-28.9-83.5-11c-29.4 19.7-37.4 59.4-17.7 88.8l67.5 101.3C108 485.3 157.9 512 211.4 512H328c66.3 0 120-53.7 120-120V272c0-35.3-28.7-64-64-64c-4.5 0-8.8.5-13 1.3c-11.7-15.4-30.2-25.3-51-25.3c-6.9 0-13.5 1.1-19.7 3.1c-11.6-16.4-30.7-27.1-52.3-27.1c-2.7 0-5.4.2-8 .5V64c0-35.3-28.7-64-64-64m48 304c0-8.8-7.2-16-16-16s-16 7.2-16 16v96c0 8.8 7.2 16 16 16s16-7.2 16-16zm48-16c-8.8 0-16 7.2-16 16v96c0 8.8 7.2 16 16 16s16-7.2 16-16v-96c0-8.8-7.2-16-16-16m80 16c0-8.8-7.2-16-16-16s-16 7.2-16 16v96c0 8.8 7.2 16 16 16s16-7.2 16-16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandScissors(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.2 276.3c-1.2-35.3 26.4-65 61.7-66.2l3.3-.1l-8.2-1.9C22.5 200.5.7 166.3 8.3 131.8s41.9-56.3 76.4-48.6l173 38.3c2.3-2.9 4.7-5.7 7.1-8.5l18.4-20.3C299.9 74.5 323.5 64 348.3 64h10.2c54.1 0 104.1 28.7 131.3 75.4l1.5 2.6c13.6 23.2 20.7 49.7 20.7 76.6V344c0 66.3-53.7 120-120 120H288c-35.3 0-64-28.7-64-64c0-2.8.2-5.6.5-8.3c-19.4-11-32.5-31.8-32.5-55.7v-2.4L66.4 338c-35.3 1.2-65-26.4-66.2-61.7m63.4-18.2c-8.8.3-15.7 7.7-15.4 16.5s7.7 15.7 16.5 15.4l161.5-5.6c9.8-.3 18.7 5.3 22.7 14.2s2.2 19.3-4.5 26.4c-2.8 2.9-4.4 6.7-4.4 11c0 8.8 7.2 16 16 16c9.1 0 17.4 5.1 21.5 13.3s3.2 17.9-2.3 25.1c-2 2.7-3.2 6-3.2 9.6c0 8.8 7.2 16 16 16h104c39.8 0 72-32.2 72-72V218.6c0-18.4-4.9-36.5-14.2-52.4l-1.5-2.6c-18.6-32-52.8-51.6-89.8-51.6h-10.2c-11.3 0-22 4.8-29.6 13.1l-17.5-15.9l17.5 15.9l-18.4 20.3c-.6.6-1.1 1.3-1.7 1.9l57 13.2c8.6 2 14 10.6 12 19.2s-10.6 14-19.2 12L262.8 172L74.3 130c-8.6-1.9-17.2 3.5-19.1 12.2s3.5 17.2 12.2 19.1l187.5 41.6c10.2 2.3 17.8 10.9 18.7 21.4l.1 1c.6 6.6-1.5 13.1-5.8 18.1s-10.6 7.9-17.2 8.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandSpock(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M170.2 80.8C161 47 180.8 12 214.6 2.4c34-9.6 69.4 10.2 79 44.2l30.3 107.1L337.1 84c6.6-34.7 40.1-57.5 74.8-50.9c31.4 6 53 33.9 52 64.9c10-2.6 20.8-2.8 31.5-.1c34.3 8.6 55.1 43.3 46.6 77.6l-55.3 221.7C469.8 464.7 409.2 512 339.6 512h-33.7c-56.9 0-112.2-19-157.2-53.9l-92-71.6c-27.9-21.7-32.9-61.9-11.2-89.8s61.9-32.9 89.8-11.2l17 13.2l-51.8-131.2c-13-32.9 3.2-70.1 36-83c11.1-4.4 22.7-5.4 33.7-3.7m77.1-21.2c-2.4-8.5-11.2-13.4-19.7-11s-13.4 11.2-11 19.7l54.8 182.4c3.5 12.3-3.3 25.2-15.4 29.3s-25.3-2-30-13.9l-51.1-128c-3.2-8.2-12.5-12.3-20.8-9s-12.3 12.5-9 20.8l73.3 185.6c12 30.3-23.7 57-49.4 37l-63.1-49.1c-7-5.4-17-4.2-22.5 2.8s-4.2 17 2.8 22.5l92 71.6c36.5 28.4 81.4 43.8 127.7 43.8h33.7c47.5 0 89-32.4 100.5-78.5L495.5 164c2.1-8.6-3.1-17.3-11.6-19.4s-17.3 3.1-19.4 11.6l-26 104C435.6 271.8 425 280 413 280c-16.5 0-28.9-15-25.8-31.2L415.7 99c1.7-8.7-4-17.1-12.7-18.7s-17.1 4-18.7 12.7l-31.8 167c-2.2 11.6-12.4 20-24.2 20c-11 0-20.7-7.3-23.7-17.9L247.4 59.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Handshake(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m272.2 64.6l-51.1 51.1c-15.3 4.2-29.5 11.9-41.5 22.5L153 161.9c-10.2 9.1-23.5 14.1-37.2 14.1H96v128c20.4.6 39.8 8.9 54.3 23.4l35.6 35.6l7 7l27 27c6.2 6.2 16.4 6.2 22.6 0c1.7-1.7 3-3.7 3.7-5.8c2.8-7.7 9.3-13.5 17.3-15.3s16.4.6 22.2 6.5l10.8 10.6c11.6 11.6 30.4 11.6 41.9 0c5.4-5.4 8.3-12.3 8.6-19.4c.4-8.8 5.6-16.6 13.6-20.4s17.3-3 24.4 2.1c9.4 6.7 22.5 5.8 30.9-2.6c9.4-9.4 9.4-24.6 0-33.9L340.1 243l-35.8 33c-27.3 25.2-69.2 25.6-97 .9c-31.7-28.2-32.4-77.4-1.6-106.5l70.1-66.2C303.2 78.4 339.4 64 377.1 64c36.1 0 71 13.3 97.9 37.2l30.1 26.8H624c8.8 0 16 7.2 16 16v208c0 17.7-14.3 32-32 32h-32c-11.8 0-22.2-6.4-27.7-16h-84.9c-3.4 6.7-7.9 13.1-13.5 18.7c-17.1 17.1-40.8 23.8-63 20.1c-3.6 7.3-8.5 14.1-14.6 20.2c-27.3 27.3-70 30-100.4 8.1c-25.1 20.8-62.5 19.5-86-4.1L159 404l-7-7l-35.6-35.6c-5.5-5.5-12.7-8.7-20.4-9.3c0 17.6-14.4 31.9-32 31.9H32c-17.7 0-32-14.3-32-32V144c0-8.8 7.2-16 16-16h99.8c2 0 3.9-.7 5.3-2l26.5-23.6C175.5 77.7 211.4 64 248.7 64H259c4.4 0 8.9.2 13.2.6M544 320V176h-48c-5.9 0-11.6-2.2-15.9-6.1l-36.9-32.8C425 120.9 401.5 112 377.1 112c-25.4 0-49.8 9.7-68.3 27.1l-70.1 66.2c-10.3 9.8-10.1 26.3.5 35.7c9.3 8.3 23.4 8.1 32.5-.3l71.9-66.4c9.7-9 24.9-8.4 33.9 1.4s8.4 24.9-1.4 33.9l-.8.8l74.4 74.4c10 10 16.5 22.3 19.4 35.1H544zM64 336a16 16 0 1 0-32 0a16 16 0 1 0 32 0m528 16a16 16 0 1 0 0-32a16 16 0 1 0 0 32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HardDrive(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 80c-8.8 0-16 7.2-16 16v162c5.1-1.3 10.5-2 16-2h384c5.5 0 10.9.7 16 2V96c0-8.8-7.2-16-16-16zM48 320v96c0 8.8 7.2 16 16 16h384c8.8 0 16-7.2 16-16v-96c0-8.8-7.2-16-16-16H64c-8.8 0-16 7.2-16 16m-48 0V96c0-35.3 28.7-64 64-64h384c35.3 0 64 28.7 64 64v320c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64zm280 48a24 24 0 1 1 48 0a24 24 0 1 1-48 0m120-24a24 24 0 1 1 0 48a24 24 0 1 1 0-48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m225.8 468.2l-2.5-2.3L48.1 303.2C17.4 274.7 0 234.7 0 192.8v-3.3c0-70.4 50-130.8 119.2-144c39.4-7.6 79.7 1.5 111.8 24.1c9 6.4 17.4 13.8 25 22.3c4.2-4.8 8.7-9.2 13.5-13.3c3.7-3.2 7.5-6.2 11.5-9c32.1-22.6 72.4-31.7 111.8-24.2C462 58.6 512 119.1 512 189.5v3.3c0 41.9-17.4 81.9-48.1 110.4L288.7 465.9l-2.5 2.3c-8.2 7.6-19 11.9-30.2 11.9s-22-4.2-30.2-11.9M239.1 145c-.4-.3-.7-.7-1-1.1l-17.8-20l-.1-.1c-23.1-25.9-58-37.7-92-31.2c-46.6 8.9-80.2 49.5-80.2 96.9v3.3c0 28.5 11.9 55.8 32.8 75.2L256 430.7L431.2 268a102.7 102.7 0 0 0 32.8-75.2v-3.3c0-47.3-33.6-88-80.1-96.9c-34-6.5-69 5.4-92 31.2l-.1.1l-.1.1l-17.8 20c-.3.4-.7.7-1 1.1c-4.5 4.5-10.6 7-16.9 7s-12.4-2.5-16.9-7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hospital(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M232 0c-39.8 0-72 32.2-72 72v8H72c-39.8 0-72 32.2-72 72v288c0 39.8 32.2 72 72 72h496c39.8 0 72-32.2 72-72V152c0-39.8-32.2-72-72-72h-88v-8c0-39.8-32.2-72-72-72zm248 128h88c13.3 0 24 10.7 24 24v40h-56c-13.3 0-24 10.7-24 24s10.7 24 24 24h56v48h-56c-13.3 0-24 10.7-24 24s10.7 24 24 24h56v104c0 13.3-10.7 24-24 24h-88V336zm-408 0h88v336H72c-13.2 0-24-10.7-24-24V336h56c13.3 0 24-10.7 24-24s-10.7-24-24-24H48v-48h56c13.3 0 24-10.7 24-24s-10.7-24-24-24H48v-40c0-13.3 10.7-24 24-24m136-56c0-13.3 10.7-24 24-24h176c13.3 0 24 10.7 24 24v392h-64v-64c0-26.5-21.5-48-48-48s-48 21.5-48 48v64h-64zm88 24v24h-24c-8.8 0-16 7.2-16 16v16c0 8.8 7.2 16 16 16h24v24c0 8.8 7.2 16 16 16h16c8.8 0 16-7.2 16-16v-24h24c8.8 0 16-7.2 16-16v-16c0-8.8-7.2-16-16-16h-24V96c0-8.8-7.2-16-16-16h-16c-8.8 0-16 7.2-16 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hourglass(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24 0C10.7 0 0 10.7 0 24s10.7 24 24 24h8v19c0 40.3 16 79 44.5 107.5l81.6 81.5l-81.6 81.5C48 366 32 404.7 32 445v19h-8c-13.3 0-24 10.7-24 24s10.7 24 24 24h336c13.3 0 24-10.7 24-24s-10.7-24-24-24h-8v-19c0-40.3-16-79-44.5-107.5L225.9 256l81.5-81.5C336 146 352 107.3 352 67V48h8c13.3 0 24-10.7 24-24S373.3 0 360 0zm168 289.9l81.5 81.5C293 391 304 417.4 304 445v19H80v-19c0-27.6 11-54 30.5-73.5zm0-67.9l-81.5-81.5C91 121 80 94.6 80 67V48h224v19c0 27.6-11 54-30.5 73.5L192 222.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HourglassHalf(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 24C0 10.7 10.7 0 24 0h336c13.3 0 24 10.7 24 24s-10.7 24-24 24h-8v19c0 40.3-16 79-44.5 107.5L225.9 256l81.5 81.5C336 366 352 404.7 352 445v19h8c13.3 0 24 10.7 24 24s-10.7 24-24 24H24c-13.3 0-24-10.7-24-24s10.7-24 24-24h8v-19c0-40.3 16-79 44.5-107.5l81.6-81.5l-81.6-81.5C48 146 32 107.3 32 67V48h-8C10.7 48 0 37.3 0 24m110.5 347.5c-3.9 3.9-7.5 8.1-10.7 12.5h184.4c-3.2-4.4-6.8-8.6-10.7-12.5L192 289.9l-81.5 81.5zM284.2 128C297 110.4 304 89 304 67V48H80v19c0 22.1 7 43.4 19.8 61z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IdBadge(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 48v16c0 17.7-14.3 32-32 32h-64c-17.7 0-32-14.3-32-32V48H64c-8.8 0-16 7.2-16 16v384c0 8.8 7.2 16 16 16h256c8.8 0 16-7.2 16-16V64c0-8.8-7.2-16-16-16zM0 64C0 28.7 28.7 0 64 0h256c35.3 0 64 28.7 64 64v384c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64zm160 256h64c44.2 0 80 35.8 80 80c0 8.8-7.2 16-16 16H96c-8.8 0-16-7.2-16-16c0-44.2 35.8-80 80-80m-32-96a64 64 0 1 1 128 0a64 64 0 1 1-128 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IdCard(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M528 160v256c0 8.8-7.2 16-16 16H320c0-44.2-35.8-80-80-80h-64c-44.2 0-80 35.8-80 80H64c-8.8 0-16-7.2-16-16V160zM64 32C28.7 32 0 60.7 0 96v320c0 35.3 28.7 64 64 64h448c35.3 0 64-28.7 64-64V96c0-35.3-28.7-64-64-64zm208 224a64 64 0 1 0-128 0a64 64 0 1 0 128 0m104-48c-13.3 0-24 10.7-24 24s10.7 24 24 24h80c13.3 0 24-10.7 24-24s-10.7-24-24-24zm0 96c-13.3 0-24 10.7-24 24s10.7 24 24 24h80c13.3 0 24-10.7 24-24s-10.7-24-24-24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Image(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 80c8.8 0 16 7.2 16 16v319.8l-5-6.5l-136-176c-4.5-5.9-11.6-9.3-19-9.3s-14.4 3.4-19 9.3l-83 107.4l-30.5-42.7c-4.5-6.3-11.7-10-19.5-10s-15 3.7-19.5 10.1l-80 112l-4.5 6.2V96c0-8.8 7.2-16 16-16zM64 32C28.7 32 0 60.7 0 96v320c0 35.3 28.7 64 64 64h384c35.3 0 64-28.7 64-64V96c0-35.3-28.7-64-64-64zm80 192a48 48 0 1 0 0-96a48 48 0 1 0 0 96"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Images(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M160 80h352c8.8 0 16 7.2 16 16v224c0 8.8-7.2 16-16 16h-21.2L388.1 178.9c-4.4-6.8-12-10.9-20.1-10.9s-15.7 4.1-20.1 10.9l-52.2 79.8l-12.4-16.9c-4.5-6.2-11.7-9.8-19.4-9.8s-14.8 3.6-19.4 9.8L175.6 336H160c-8.8 0-16-7.2-16-16V96c0-8.8 7.2-16 16-16M96 96v224c0 35.3 28.7 64 64 64h352c35.3 0 64-28.7 64-64V96c0-35.3-28.7-64-64-64H160c-35.3 0-64 28.7-64 64m-48 24c0-13.3-10.7-24-24-24S0 106.7 0 120v224c0 75.1 60.9 136 136 136h320c13.3 0 24-10.7 24-24s-10.7-24-24-24H136c-48.6 0-88-39.4-88-88zm208 24a32 32 0 1 0-64 0a32 32 0 1 0 64 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyboard(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 112c-8.8 0-16 7.2-16 16v256c0 8.8 7.2 16 16 16h448c8.8 0 16-7.2 16-16V128c0-8.8-7.2-16-16-16zM0 128c0-35.3 28.7-64 64-64h448c35.3 0 64 28.7 64 64v256c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64zm176 192h224c8.8 0 16 7.2 16 16v16c0 8.8-7.2 16-16 16H176c-8.8 0-16-7.2-16-16v-16c0-8.8 7.2-16 16-16m-72-72c0-8.8 7.2-16 16-16h16c8.8 0 16 7.2 16 16v16c0 8.8-7.2 16-16 16h-16c-8.8 0-16-7.2-16-16zm16-96h16c8.8 0 16 7.2 16 16v16c0 8.8-7.2 16-16 16h-16c-8.8 0-16-7.2-16-16v-16c0-8.8 7.2-16 16-16m64 96c0-8.8 7.2-16 16-16h16c8.8 0 16 7.2 16 16v16c0 8.8-7.2 16-16 16h-16c-8.8 0-16-7.2-16-16zm16-96h16c8.8 0 16 7.2 16 16v16c0 8.8-7.2 16-16 16h-16c-8.8 0-16-7.2-16-16v-16c0-8.8 7.2-16 16-16m64 96c0-8.8 7.2-16 16-16h16c8.8 0 16 7.2 16 16v16c0 8.8-7.2 16-16 16h-16c-8.8 0-16-7.2-16-16zm16-96h16c8.8 0 16 7.2 16 16v16c0 8.8-7.2 16-16 16h-16c-8.8 0-16-7.2-16-16v-16c0-8.8 7.2-16 16-16m64 96c0-8.8 7.2-16 16-16h16c8.8 0 16 7.2 16 16v16c0 8.8-7.2 16-16 16h-16c-8.8 0-16-7.2-16-16zm16-96h16c8.8 0 16 7.2 16 16v16c0 8.8-7.2 16-16 16h-16c-8.8 0-16-7.2-16-16v-16c0-8.8 7.2-16 16-16m64 96c0-8.8 7.2-16 16-16h16c8.8 0 16 7.2 16 16v16c0 8.8-7.2 16-16 16h-16c-8.8 0-16-7.2-16-16zm16-96h16c8.8 0 16 7.2 16 16v16c0 8.8-7.2 16-16 16h-16c-8.8 0-16-7.2-16-16v-16c0-8.8 7.2-16 16-16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lemon(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M368 80c-3.2 0-6.2.4-8.9 1.3c-19.1 5.5-46.1 10.6-74.3 3.3c-57.4-14.9-124.6 7.4-174.7 57.5S37.7 259.4 52.6 316.8c7.3 28.2 2.2 55.2-3.3 74.3c-.8 2.8-1.3 5.8-1.3 8.9c0 17.7 14.3 32 32 32c3.2 0 6.2-.4 8.9-1.3c19.1-5.5 46.1-10.7 74.3-3.3c57.4 14.9 124.6-7.4 174.7-57.5s72.4-117.3 57.5-174.7c-7.3-28.2-2.2-55.2 3.3-74.3c.8-2.8 1.3-5.8 1.3-8.9c0-17.7-14.3-32-32-32m0-48c44.2 0 80 35.8 80 80c0 7.7-1.1 15.2-3.1 22.3c-4.6 15.8-7.1 32.9-3 48.9c20.1 77.6-10.9 161.5-70 220.7s-143.1 90.2-220.7 70c-16-4.1-33-1.6-48.9 3c-7.1 2-14.6 3.1-22.3 3.1c-44.2 0-80-35.8-80-80c0-7.7 1.1-15.2 3.1-22.3c4.6-15.8 7.1-32.9 3-48.9C-14 251.3 17 167.3 76.2 108.2S219.3 18 296.8 38.1c16 4.1 33 1.6 48.9-3c7.1-2 14.6-3.1 22.3-3.1M246.7 167c-52 15.2-96.5 59.7-111.7 111.7c-3.7 12.7-17.1 20-29.8 16.3s-20-17.1-16.3-29.8c19.8-67.7 76.6-124.5 144.3-144.3c12.7-3.7 26.1 3.6 29.8 16.3s-3.6 26.1-16.3 29.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LifeRing(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M385.1 419.1C349.7 447.2 304.8 464 256 464s-93.7-16.8-129.1-44.9l80.4-80.4c14.3 8.4 31 13.3 48.8 13.3s34.5-4.8 48.8-13.3l80.4 80.4zm68.1.2C489.9 374.9 512 318.1 512 256s-22.1-118.9-58.8-163.3L465 81c9.4-9.4 9.4-24.6 0-33.9s-24.6-9.4-33.9 0l-11.8 11.7C374.9 22.1 318.1 0 256 0S137.1 22.1 92.7 58.8L81 47c-9.4-9.4-24.6-9.4-33.9 0s-9.4 24.6 0 33.9l11.7 11.8C22.1 137.1 0 193.9 0 256s22.1 118.9 58.8 163.3L47 431c-9.4 9.4-9.4 24.6 0 33.9s24.6 9.4 33.9 0l11.8-11.8C137.1 489.9 193.9 512 256 512s118.9-22.1 163.3-58.8L431 465c9.4 9.4 24.6 9.4 33.9 0s9.4-24.6 0-33.9l-11.8-11.8zm-34.1-34.1l-80.4-80.4c8.4-14.3 13.3-31 13.3-48.8s-4.8-34.5-13.3-48.8l80.4-80.4C447.2 162.3 464 207.2 464 256s-16.8 93.7-44.9 129.1zm-34-292.3l-80.4 80.4c-14.3-8.4-31-13.3-48.8-13.3s-34.5 4.8-48.8 13.3l-80.2-80.4C162.3 64.8 207.2 48 256 48s93.7 16.8 129.1 44.9M173.3 304.8l-80.4 80.3C64.8 349.7 48 304.8 48 256s16.8-93.7 44.9-129.1l80.4 80.4c-8.4 14.3-13.3 31-13.3 48.8s4.8 34.5 13.3 48.8zM208 256a48 48 0 1 1 96 0a48 48 0 1 1-96 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lightbulb(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M297.2 248.9c14.4-20.6 22.8-45.7 22.8-72.9c0-70.7-57.3-128-128-128S64 105.3 64 176c0 27.2 8.4 52.3 22.8 72.9c3.7 5.3 8.1 11.3 12.8 17.7c12.9 17.7 28.3 38.9 39.8 59.8c10.4 19 15.7 38.8 18.3 57.5H109c-2.2-12-5.9-23.7-11.8-34.5c-9.9-18-22.2-34.9-34.5-51.8c-5.2-7.1-10.4-14.2-15.4-21.4C27.6 247.9 16 213.3 16 176C16 78.8 94.8 0 192 0s176 78.8 176 176c0 37.3-11.6 71.9-31.4 100.3c-5 7.2-10.2 14.3-15.4 21.4c-12.3 16.8-24.6 33.7-34.5 51.8c-5.9 10.8-9.6 22.5-11.8 34.5h-48.5c2.6-18.7 7.9-38.6 18.3-57.5c11.5-20.9 26.9-42.1 39.8-59.8c4.7-6.4 9-12.4 12.7-17.7zM192 128c-26.5 0-48 21.5-48 48c0 8.8-7.2 16-16 16s-16-7.2-16-16c0-44.2 35.8-80 80-80c8.8 0 16 7.2 16 16s-7.2 16-16 16m0 384c-44.2 0-80-35.8-80-80v-16h160v16c0 44.2-35.8 80-80 80"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Map(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M565.6 36.2C572.1 40.7 576 48.1 576 56v336c0 10-6.2 18.9-15.5 22.4l-168 64c-5.2 2-10.9 2.1-16.1.3l-183.9-61.2l-160 61c-7.4 2.8-15.7 1.8-22.2-2.7S0 463.9 0 456V120c0-10 6.1-18.9 15.5-22.4l168-64c5.2-2 10.9-2.1 16.1-.3l183.9 61.2l160-61c7.4-2.8 15.7-1.8 22.2 2.7zM48 136.5v284.7l120-45.7V90.8zm312 286.2V137.3l-144-48v285.4zm48-1.5l120-45.7V90.8l-120 45.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Message(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M160 368c26.5 0 48 21.5 48 48v16l72.5-54.4c8.3-6.2 18.4-9.6 28.8-9.6H448c8.8 0 16-7.2 16-16V64c0-8.8-7.2-16-16-16H64c-8.8 0-16 7.2-16 16v288c0 8.8 7.2 16 16 16zm48 124l-.2.2l-5.1 3.8l-17.1 12.8c-4.8 3.6-11.3 4.2-16.8 1.5s-8.8-8.2-8.8-14.3v-80H64c-35.3 0-64-28.7-64-64V64C0 28.7 28.7 0 64 0h384c35.3 0 64 28.7 64 64v288c0 35.3-28.7 64-64 64H309.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneyBillOne(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M112 112c0 35.3-28.7 64-64 64v160c35.3 0 64 28.7 64 64h352c0-35.3 28.7-64 64-64V176c-35.3 0-64-28.7-64-64zM0 128c0-35.3 28.7-64 64-64h448c35.3 0 64 28.7 64 64v256c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64zm176 128a112 112 0 1 1 224 0a112 112 0 1 1-224 0m80-48c0 8.8 7.2 16 16 16v64h-8c-8.8 0-16 7.2-16 16s7.2 16 16 16h48c8.8 0 16-7.2 16-16s-7.2-16-16-16h-8v-80c0-8.8-7.2-16-16-16h-16c-8.8 0-16 7.2-16 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Moon(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M144.7 98.7c-21 34.1-33.1 74.3-33.1 117.3c0 98 62.8 181.4 150.4 211.7c-12.4 2.8-25.3 4.3-38.6 4.3C126.6 432 48 353.3 48 256c0-68.9 39.4-128.4 96.8-157.3zm62.1-66C91.1 41.2 0 137.9 0 256c0 123.7 100 224 223.5 224c47.8 0 92-15 128.4-40.6c1.9-1.3 3.7-2.7 5.5-4c4.8-3.6 9.4-7.4 13.9-11.4c2.7-2.4 5.3-4.8 7.9-7.3c5-4.9 6.3-12.5 3.1-18.7s-10.1-9.7-17-8.5c-3.7.6-7.4 1.2-11.1 1.6c-5 .5-10.1.9-15.3 1h-4c-96.8-.2-175.2-78.9-175.2-176c0-54.8 24.9-103.7 64.1-136c1-.9 2.1-1.7 3.2-2.6c4-3.2 8.2-6.2 12.5-9c3.1-2 6.3-4 9.6-5.8c6.1-3.5 9.2-10.5 7.7-17.3s-7.3-11.9-14.3-12.5c-3.6-.3-7.1-.5-10.7-.6c-2.7-.1-5.5-.1-8.2-.1c-3.3 0-6.5.1-9.8.2c-2.3.1-4.6.2-6.9.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Newspaper(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M168 80c-13.3 0-24 10.7-24 24v304c0 8.4-1.4 16.5-4.1 24H440c13.3 0 24-10.7 24-24V104c0-13.3-10.7-24-24-24zM72 480c-39.8 0-72-32.2-72-72V112c0-13.3 10.7-24 24-24s24 10.7 24 24v296c0 13.3 10.7 24 24 24s24-10.7 24-24V104c0-39.8 32.2-72 72-72h272c39.8 0 72 32.2 72 72v304c0 39.8-32.2 72-72 72zm104-344c0-13.3 10.7-24 24-24h96c13.3 0 24 10.7 24 24v80c0 13.3-10.7 24-24 24h-96c-13.3 0-24-10.7-24-24zm200-24h32c13.3 0 24 10.7 24 24s-10.7 24-24 24h-32c-13.3 0-24-10.7-24-24s10.7-24 24-24m0 80h32c13.3 0 24 10.7 24 24s-10.7 24-24 24h-32c-13.3 0-24-10.7-24-24s10.7-24 24-24m-176 80h208c13.3 0 24 10.7 24 24s-10.7 24-24 24H200c-13.3 0-24-10.7-24-24s10.7-24 24-24m0 80h208c13.3 0 24 10.7 24 24s-10.7 24-24 24H200c-13.3 0-24-10.7-24-24s10.7-24 24-24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notdef(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 0H48C21.49 0 0 21.49 0 48v416c0 26.5 21.49 48 48 48h288c26.51 0 48-21.49 48-48V48c0-26.51-21.5-48-48-48zm0 90.16V421.8L221.2 256L336 90.16zM192 213.8L77.19 48h229.6L192 213.8zM162.8 256L48 421.8V90.16L162.8 256zm29.2 42.2L306.8 464H77.19L192 298.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoteSticky(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 80c-8.8 0-16 7.2-16 16v320c0 8.8 7.2 16 16 16h224v-80c0-17.7 14.3-32 32-32h80V96c0-8.8-7.2-16-16-16zm224 400H64c-35.3 0-64-28.7-64-64V96c0-35.3 28.7-64 64-64h320c35.3 0 64 28.7 64 64v229.5c0 17-6.7 33.3-18.7 45.3l-90.5 90.5c-12 12-28.3 18.7-45.3 18.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectGroup(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M48 115.8C38.2 107 32 94.2 32 80c0-26.5 21.5-48 48-48c14.2 0 27 6.2 35.8 16h344.4c8.8-9.8 21.6-16 35.8-16c26.5 0 48 21.5 48 48c0 14.2-6.2 27-16 35.8v280.4c9.8 8.8 16 21.6 16 35.8c0 26.5-21.5 48-48 48c-14.2 0-27-6.2-35.8-16H115.8c-8.8 9.8-21.6 16-35.8 16c-26.5 0-48-21.5-48-48c0-14.2 6.2-27 16-35.8zM125.3 96c-4.8 13.6-15.6 24.4-29.3 29.3v261.4c13.6 4.8 24.4 15.6 29.3 29.3h325.4c4.8-13.6 15.6-24.4 29.3-29.3V125.3c-13.6-4.8-24.4-15.6-29.3-29.3zm2.7 64c0-17.7 14.3-32 32-32h128c17.7 0 32 14.3 32 32v96c0 17.7-14.3 32-32 32H160c-17.7 0-32-14.3-32-32zm128 160h32c35.3 0 64-28.7 64-64v-32h64c17.7 0 32 14.3 32 32v96c0 17.7-14.3 32-32 32H288c-17.7 0-32-14.3-32-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectUngroup(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M48.2 66.8c-.1-.8-.2-1.7-.2-2.5v-.2c0-8.8 7.2-16 16-16c.9 0 1.9.1 2.8.2C74.3 49.5 80 56.1 80 64c0 8.8-7.2 16-16 16c-7.9 0-14.5-5.7-15.8-13.2M0 64c0 26.9 16.5 49.9 40 59.3v105.4C16.5 238.1 0 261.1 0 288c0 35.3 28.7 64 64 64c26.9 0 49.9-16.5 59.3-40h201.4c9.5 23.5 32.5 40 59.3 40c35.3 0 64-28.7 64-64c0-26.9-16.5-49.9-40-59.3V123.3c23.5-9.5 40-32.5 40-59.3c0-35.3-28.7-64-64-64c-26.9 0-49.9 16.5-59.3 40H123.3C113.9 16.5 90.9 0 64 0C28.7 0 0 28.7 0 64m368 0a16 16 0 1 1 32 0a16 16 0 1 1-32 0m-43.3 24c6.5 16 19.3 28.9 35.3 35.3v105.4c-16 6.5-28.9 19.3-35.3 35.3H123.3c-6.5-16-19.3-28.9-35.3-35.3V123.3c16-6.5 28.9-19.3 35.3-35.3zM384 272a16 16 0 1 1 0 32a16 16 0 1 1 0-32M80 288c0 7.9-5.7 14.5-13.2 15.8c-.8.1-1.7.2-2.5.2h-.2c-8.8 0-16-7.2-16-16c0-.9.1-1.9.2-2.8c1.2-7.5 7.8-13.2 15.7-13.2c8.8 0 16 7.2 16 16m391.3-40h45.4c6.5 16 19.3 28.9 35.3 35.3v105.4c-16 6.5-28.9 19.3-35.3 35.3H315.3c-6.5-16-19.3-28.9-35.3-35.3V352h-48v36.7c-23.5 9.5-40 32.5-40 59.3c0 35.3 28.7 64 64 64c26.9 0 49.9-16.5 59.3-40h201.4c9.5 23.5 32.5 40 59.3 40c35.3 0 64-28.7 64-64c0-26.9-16.5-49.9-40-59.3V283.3c23.5-9.5 40-32.5 40-59.3c0-35.3-28.7-64-64-64c-26.9 0-49.9 16.5-59.3 40H448v16.4c9.8 8.8 17.8 19.5 23.3 31.6m88.9-26.7a16 16 0 1 1 31.5 5.5a16 16 0 1 1-31.5-5.5M271.8 450.7a16 16 0 1 1-31.5-5.5a16 16 0 1 1 31.5 5.5m301.5 13c-7.5-1.3-13.2-7.9-13.2-15.8c0-8.8 7.2-16 16-16c7.9 0 14.5 5.7 15.8 13.2v.1c.1.9.2 1.8.2 2.7c0 8.8-7.2 16-16 16c-.9 0-1.9-.1-2.8-.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperPlane(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.1 260.2c-22.6 12.9-20.5 47.3 3.6 57.3L160 376v103.3c0 18.1 14.6 32.7 32.7 32.7c9.7 0 18.9-4.3 25.1-11.8l62-74.3l123.9 51.6c18.9 7.9 40.8-4.5 43.9-24.7l64-416c1.9-12.1-3.4-24.3-13.5-31.2s-23.3-7.5-34-1.4zm52.1 25.5L409.7 90.6L190.1 336l1.2 1zm335.1 139.7l-166.6-69.5l214.1-239.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paste(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M104.6 48H64C28.7 48 0 76.7 0 112v272c0 35.3 28.7 64 64 64h96v-48H64c-8.8 0-16-7.2-16-16V112c0-8.8 7.2-16 16-16h16c0 17.7 14.3 32 32 32h72.4c17.6-19.6 43.2-32 71.6-32h62c-7.1-27.6-32.2-48-62-48h-40.6C211.6 20.9 188.2 0 160 0s-51.6 20.9-55.4 48m39.4 8a16 16 0 1 1 32 0a16 16 0 1 1-32 0m304 408H256c-8.8 0-16-7.2-16-16V192c0-8.8 7.2-16 16-16h140.1l67.9 67.9V448c0 8.8-7.2 16-16 16m-192 48h192c35.3 0 64-28.7 64-64V243.9c0-12.7-5.1-24.9-14.1-33.9L430 142.1c-9-9-21.2-14.1-33.9-14.1H256c-35.3 0-64 28.7-64 64v256c0 35.3 28.7 64 64 64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PenToSquare(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M441 58.9L453.1 71c9.4 9.4 9.4 24.6 0 33.9L424 134.1L377.9 88L407 58.9c9.4-9.4 24.6-9.4 33.9 0zM209.8 256.2L344 121.9l46.1 46.1l-134.3 134.2c-2.9 2.9-6.5 5-10.4 6.1L186.9 325l16.7-58.5c1.1-3.9 3.2-7.5 6.1-10.4zM373.1 25L175.8 222.2c-8.7 8.7-15 19.4-18.3 31.1l-28.6 100c-2.4 8.4-.1 17.4 6.1 23.6s15.2 8.5 23.6 6.1l100-28.6c11.8-3.4 22.5-9.7 31.1-18.3L487 138.9c28.1-28.1 28.1-73.7 0-101.8L474.9 25c-28.1-28.1-73.7-28.1-101.8 0M88 64c-48.6 0-88 39.4-88 88v272c0 48.6 39.4 88 88 88h272c48.6 0 88-39.4 88-88V312c0-13.3-10.7-24-24-24s-24 10.7-24 24v112c0 22.1-17.9 40-40 40H88c-22.1 0-40-17.9-40-40V152c0-22.1 17.9-40 40-40h112c13.3 0 24-10.7 24-24s-10.7-24-24-24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RectangleList(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 80c-8.8 0-16 7.2-16 16v320c0 8.8 7.2 16 16 16h448c8.8 0 16-7.2 16-16V96c0-8.8-7.2-16-16-16zM0 96c0-35.3 28.7-64 64-64h448c35.3 0 64 28.7 64 64v320c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64zm96 64a32 32 0 1 1 64 0a32 32 0 1 1-64 0m104 0c0-13.3 10.7-24 24-24h224c13.3 0 24 10.7 24 24s-10.7 24-24 24H224c-13.3 0-24-10.7-24-24m0 96c0-13.3 10.7-24 24-24h224c13.3 0 24 10.7 24 24s-10.7 24-24 24H224c-13.3 0-24-10.7-24-24m0 96c0-13.3 10.7-24 24-24h224c13.3 0 24 10.7 24 24s-10.7 24-24 24H224c-13.3 0-24-10.7-24-24m-72-64a32 32 0 1 1 0-64a32 32 0 1 1 0 64m-32 64a32 32 0 1 1 64 0a32 32 0 1 1-64 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RectangleXmark(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 80c-8.8 0-16 7.2-16 16v320c0 8.8 7.2 16 16 16h384c8.8 0 16-7.2 16-16V96c0-8.8-7.2-16-16-16zM0 96c0-35.3 28.7-64 64-64h384c35.3 0 64 28.7 64 64v320c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64zm175 79c9.4-9.4 24.6-9.4 33.9 0l47 47l47-47c9.4-9.4 24.6-9.4 33.9 0s9.4 24.6 0 33.9l-47 47l47 47c9.4 9.4 9.4 24.6 0 33.9s-24.6 9.4-33.9 0l-47-47l-47 47c-9.4 9.4-24.6 9.4-33.9 0s-9.4-24.6 0-33.9l47-47l-47-47c-9.4-9.4-9.4-24.6 0-33.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Registered(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 48a208 208 0 1 1 0 416a208 208 0 1 1 0-416m0 464a256 256 0 1 0 0-512a256 256 0 1 0 0 512m-96-360v208c0 13.3 10.7 24 24 24s24-10.7 24-24v-72h60.9l37.2 81.9c5.5 12.1 19.7 17.4 31.8 11.9s17.4-19.7 11.9-31.8l-34.1-75c21.8-14.3 36.3-39 36.3-67c0-44.2-35.8-80-80-80h-88c-13.3 0-24 10.7-24 24m48 88v-64h64c17.7 0 32 14.3 32 32s-14.3 32-32 32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareFromSquare(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 255.4V208c0-8.8-7.2-16-16-16h-94.5c-50.9 0-93.9 33.5-108.3 79.6c-3.3-9.4-5.2-19.8-5.2-31.6c0-61.9 50.1-112 112-112h96c8.8 0 16-7.2 16-16V64.6L506 160zM336 240h16v48c0 17.7 14.3 32 32 32h3.7c7.9 0 15.5-2.9 21.4-8.2l139-125.1c7.6-6.8 11.9-16.5 11.9-26.7s-4.3-19.9-11.9-26.7L409.9 8.9C403.5 3.2 395.3 0 386.7 0C367.5 0 352 15.5 352 34.7V80h-64c-88.4 0-160 71.6-160 160c0 60.4 34.6 99.1 63.9 120.9c5.9 4.4 11.5 8.1 16.7 11.2c4.4 2.7 8.5 4.9 11.9 6.6c3.4 1.7 6.2 3 8.2 3.9c2.2 1 4.6 1.4 7.1 1.4h2.5c9.8 0 17.8-8 17.8-17.8c0-7.8-5.3-14.7-11.6-19.5c-.4-.3-.7-.5-1.1-.8c-1.7-1.1-3.4-2.5-5-4.1c-.8-.8-1.7-1.6-2.5-2.6s-1.6-1.9-2.4-2.9c-1.8-2.5-3.5-5.3-5-8.5c-2.6-6-4.3-13.3-4.3-22.4c0-36.1 29.3-65.5 65.5-65.5H336zM72 32C32.2 32 0 64.2 0 104v336c0 39.8 32.2 72 72 72h336c39.8 0 72-32.2 72-72v-64c0-13.3-10.7-24-24-24s-24 10.7-24 24v64c0 13.3-10.7 24-24 24H72c-13.3 0-24-10.7-24-24V104c0-13.3 10.7-24 24-24h64c13.3 0 24-10.7 24-24s-10.7-24-24-24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Snowflake(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M224 0c13.3 0 24 10.7 24 24v46.1l23-23c9.4-9.4 24.6-9.4 33.9 0s9.4 24.6 0 33.9l-57 57v76.5l66.2-38.2L335 98.5c3.4-12.8 16.6-20.4 29.4-17s20.4 16.6 17 29.4l-8.4 31.3l37.1-21.4c11.5-6.6 26.2-2.7 32.8 8.8s2.7 26.2-8.8 32.8L397 183.8l31.5 8.4c12.8 3.4 20.4 16.6 17 29.4s-16.6 20.4-29.4 17l-77.8-20.9L272 256l66.2 38.2l77.8-20.9c12.8-3.4 26 4.2 29.4 17s-4.2 26-17 29.4l-31.4 8.5l37.1 21.4c11.5 6.6 15.4 21.3 8.8 32.8s-21.3 15.4-32.8 8.8L373 369.8l8.4 31.5c3.4 12.8-4.2 26-17 29.4s-26-4.2-29.4-17l-20.9-77.8l-66.1-38.3v76.5l57 57c9.4 9.4 9.4 24.6 0 33.9s-24.6 9.4-33.9 0l-23-23v46c0 13.3-10.7 24-24 24s-24-10.7-24-24v-46.1l-23 23c-9.4 9.4-24.6 9.4-33.9 0s-9.4-24.6 0-33.9l57-57v-76.4L134 335.8l-20.9 77.8c-3.4 12.8-16.6 20.4-29.4 17s-20.4-16.6-17-29.4l8.3-31.4l-37.1 21.4c-11.5 6.6-26.2 2.7-32.8-8.8s-2.7-26.2 8.8-32.8L51 328.2l-31.5-8.4c-12.8-3.4-20.4-16.6-17-29.4s16.6-20.4 29.4-17l77.8 20.9L176 256l-66.2-38.2l-77.9 20.8c-12.8 3.4-26-4.2-29.4-17s4.2-26 17-29.4l31.5-8.4l-37.1-21.4c-11.5-6.6-15.4-21.3-8.8-32.8s21.3-15.4 32.8-8.8L75 142.2l-8.4-31.5c-3.4-12.8 4.2-26 17-29.4s26 4.2 29.4 17l20.9 77.8l66.1 38.3v-76.5L143 81c-9.4-9.4-9.4-24.6 0-33.9s24.6-9.4 33.9 0l23 23V24c0-13.3 10.7-24 24-24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Square(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 80c8.8 0 16 7.2 16 16v320c0 8.8-7.2 16-16 16H64c-8.8 0-16-7.2-16-16V96c0-8.8 7.2-16 16-16zM64 32C28.7 32 0 60.7 0 96v320c0 35.3 28.7 64 64 64h320c35.3 0 64-28.7 64-64V96c0-35.3-28.7-64-64-64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareCaretDown(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 432c8.8 0 16-7.2 16-16V96c0-8.8-7.2-16-16-16H64c-8.8 0-16 7.2-16 16v320c0 8.8 7.2 16 16 16zm64-16c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64V96c0-35.3 28.7-64 64-64h320c35.3 0 64 28.7 64 64zm-224-64c-6.7 0-13-2.8-17.6-7.7l-104-112c-6.5-7-8.2-17.2-4.4-25.9s12.5-14.4 22-14.4h208c9.5 0 18.2 5.7 22 14.4s2.1 18.9-4.4 25.9l-104 112c-4.5 4.9-10.9 7.7-17.6 7.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareCaretLeft(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M48 416c0 8.8 7.2 16 16 16h320c8.8 0 16-7.2 16-16V96c0-8.8-7.2-16-16-16H64c-8.8 0-16 7.2-16 16zm16 64c-35.3 0-64-28.7-64-64V96c0-35.3 28.7-64 64-64h320c35.3 0 64 28.7 64 64v320c0 35.3-28.7 64-64 64zm64-224c0-6.7 2.8-13 7.7-17.6l112-104c7-6.5 17.2-8.2 25.9-4.4s14.4 12.5 14.4 22v208c0 9.5-5.7 18.2-14.4 22s-18.9 2.1-25.9-4.4l-112-104c-4.9-4.5-7.7-10.9-7.7-17.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareCaretRight(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 96c0-8.8-7.2-16-16-16H64c-8.8 0-16 7.2-16 16v320c0 8.8 7.2 16 16 16h320c8.8 0 16-7.2 16-16zm-16-64c35.3 0 64 28.7 64 64v320c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64V96c0-35.3 28.7-64 64-64zm-64 224c0 6.7-2.8 13-7.7 17.6l-112 104c-7 6.5-17.2 8.2-25.9 4.4S160 369.5 160 360V152c0-9.5 5.7-18.2 14.4-22s18.9-2.1 25.9 4.4l112 104c4.9 4.5 7.7 10.9 7.7 17.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareCaretUp(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 80c-8.8 0-16 7.2-16 16v320c0 8.8 7.2 16 16 16h320c8.8 0 16-7.2 16-16V96c0-8.8-7.2-16-16-16zM0 96c0-35.3 28.7-64 64-64h320c35.3 0 64 28.7 64 64v320c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64zm224 64c6.7 0 13 2.8 17.6 7.7l104 112c6.5 7 8.2 17.2 4.4 25.9S337.5 320 328 320H120c-9.5 0-18.2-5.7-22-14.4s-2.1-18.9 4.4-25.9l104-112c4.5-4.9 10.9-7.7 17.6-7.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareCheck(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 80c-8.8 0-16 7.2-16 16v320c0 8.8 7.2 16 16 16h320c8.8 0 16-7.2 16-16V96c0-8.8-7.2-16-16-16zM0 96c0-35.3 28.7-64 64-64h320c35.3 0 64 28.7 64 64v320c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64zm337 113L209 337c-9.4 9.4-24.6 9.4-33.9 0l-64-64c-9.4-9.4-9.4-24.6 0-33.9s24.6-9.4 33.9 0l47 47L303 175c9.4-9.4 24.6-9.4 33.9 0s9.4 24.6 0 33.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareFull(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 48v416H48V48zM48 0H0v512h512V0h-48z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareMinus(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 80c-8.8 0-16 7.2-16 16v320c0 8.8 7.2 16 16 16h320c8.8 0 16-7.2 16-16V96c0-8.8-7.2-16-16-16zM0 96c0-35.3 28.7-64 64-64h320c35.3 0 64 28.7 64 64v320c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64zm152 136h144c13.3 0 24 10.7 24 24s-10.7 24-24 24H152c-13.3 0-24-10.7-24-24s10.7-24 24-24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquarePlus(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 80c-8.8 0-16 7.2-16 16v320c0 8.8 7.2 16 16 16h320c8.8 0 16-7.2 16-16V96c0-8.8-7.2-16-16-16zM0 96c0-35.3 28.7-64 64-64h320c35.3 0 64 28.7 64 64v320c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64zm200 248v-64h-64c-13.3 0-24-10.7-24-24s10.7-24 24-24h64v-64c0-13.3 10.7-24 24-24s24 10.7 24 24v64h64c13.3 0 24 10.7 24 24s-10.7 24-24 24h-64v64c0 13.3-10.7 24-24 24s-24-10.7-24-24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M287.9 0c9.2 0 17.6 5.2 21.6 13.5l68.6 141.3l153.2 22.6c9 1.3 16.5 7.6 19.3 16.3s.5 18.1-5.9 24.5L433.6 328.4L459.8 484c1.5 9-2.2 18.1-9.7 23.5s-17.3 6-25.3 1.7l-137-73.2L151 509.1c-8.1 4.3-17.9 3.7-25.3-1.7s-11.2-14.5-9.7-23.5l26.2-155.6L31.1 218.2c-6.5-6.4-8.7-15.9-5.9-24.5s10.3-14.9 19.3-16.3l153.2-22.6l68.6-141.3C270.4 5.2 278.7 0 287.9 0m0 79l-52.5 108.2c-3.5 7.1-10.2 12.1-18.1 13.3L99 217.9l85.9 85.1c5.5 5.5 8.1 13.3 6.8 21l-20.3 119.7l105.2-56.2c7.1-3.8 15.6-3.8 22.6 0l105.2 56.2l-20.2-119.6c-1.3-7.7 1.2-15.5 6.8-21l85.9-85.1l-118.3-17.5c-7.8-1.2-14.6-6.1-18.1-13.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarHalf(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M293.3.6c10.9 2.5 18.6 12.2 18.6 23.4v384.7c0 8.9-4.9 17-12.7 21.2L151 509.1c-8.1 4.3-17.9 3.7-25.3-1.7s-11.2-14.5-9.7-23.5l26.2-155.6l-111.1-110c-6.5-6.4-8.7-15.9-5.9-24.5s10.3-14.9 19.3-16.3l153.2-22.6l68.6-141.4c4.9-10.1 16.1-15.4 27-12.9m-29.4 127.8l-28.6 58.8c-3.5 7.1-10.2 12.1-18.1 13.3L99 217.9l85.9 85.1c5.5 5.5 8.1 13.3 6.8 21l-20.3 119.7l92.5-49.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarHalfStroke(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M309.5 13.5C305.5 5.2 297.1 0 287.9 0s-17.6 5.2-21.6 13.5l-68.6 141.3l-153.2 22.7c-9 1.3-16.5 7.6-19.3 16.3s-.5 18.1 5.9 24.5l111.1 110.1L116 483.9c-1.5 9 2.2 18.1 9.7 23.5s17.3 6 25.3 1.7l137-73.2l137 73.2c8.1 4.3 17.9 3.7 25.3-1.7s11.2-14.5 9.7-23.5l-26.4-155.5l111.2-110.2c6.5-6.4 8.7-15.9 5.9-24.5s-10.3-14.9-19.3-16.3l-153.3-22.6zM288 384.7V79.1l52.5 108.1c3.5 7.1 10.2 12.1 18.1 13.3L476.9 218L391 303c-5.5 5.5-8.1 13.3-6.8 21l20.2 119.6l-105.2-56.1c-3.5-1.9-7.4-2.8-11.2-2.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sun(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M375.7 19.7c-1.5-8-6.9-14.7-14.4-17.8s-16.1-2.2-22.8 2.4L256 61.1L173.5 4.2c-6.7-4.6-15.3-5.5-22.8-2.4s-12.9 9.8-14.4 17.8l-18.1 98.5l-98.5 18.2c-8 1.5-14.7 6.9-17.8 14.4s-2.2 16.1 2.4 22.8L61.1 256L4.2 338.5c-4.6 6.7-5.5 15.3-2.4 22.8s9.8 13 17.8 14.4l98.5 18.1l18.1 98.5c1.5 8 6.9 14.7 14.4 17.8s16.1 2.2 22.8-2.4l82.6-56.8l82.5 56.9c6.7 4.6 15.3 5.5 22.8 2.4s12.9-9.8 14.4-17.8l18.1-98.5l98.5-18.1c8-1.5 14.7-6.9 17.8-14.4s2.2-16.1-2.4-22.8L450.9 256l56.9-82.5c4.6-6.7 5.5-15.3 2.4-22.8s-9.8-12.9-17.8-14.4l-98.5-18.1zM269.6 110l65.6-45.2l14.4 78.3c1.8 9.8 9.5 17.5 19.3 19.3l78.3 14.4l-45.2 65.6c-5.7 8.2-5.7 19 0 27.2l45.2 65.6l-78.3 14.4c-9.8 1.8-17.5 9.5-19.3 19.3l-14.4 78.3l-65.6-45.2c-8.2-5.7-19-5.7-27.2 0l-65.6 45.2l-14.4-78.3c-1.8-9.8-9.5-17.5-19.3-19.3l-78.3-14.4l45.2-65.6c5.7-8.2 5.7-19 0-27.2l-45.2-65.6l78.3-14.4c9.8-1.8 17.5-9.5 19.3-19.3l14.4-78.3l65.6 45.2c8.2 5.7 19 5.7 27.2 0M256 368a112 112 0 1 0 0-224a112 112 0 1 0 0 224m-64-112a64 64 0 1 1 128 0a64 64 0 1 1-128 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsDown(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M323.8 477.2c-38.2 10.9-78.1-11.2-89-49.4l-5.7-20c-3.7-13-10.4-25-19.5-35l-51.3-56.4c-8.9-9.8-8.2-25 1.6-33.9s25-8.2 33.9 1.6l51.3 56.4c14.1 15.5 24.4 34 30.1 54.1l5.7 20c3.6 12.7 16.9 20.1 29.7 16.5s20.1-16.9 16.5-29.7l-5.7-20c-5.7-19.9-14.7-38.7-26.6-55.5c-5.2-7.3-5.8-16.9-1.7-24.9s12.3-13 21.3-13H448c8.8 0 16-7.2 16-16c0-6.8-4.3-12.7-10.4-15c-7.4-2.8-13-9-14.9-16.7s.1-15.8 5.3-21.7c2.5-2.8 4-6.5 4-10.6c0-7.8-5.6-14.3-13-15.7c-8.2-1.6-15.1-7.3-18-15.2s-1.6-16.7 3.6-23.3c2.1-2.7 3.4-6.1 3.4-9.9c0-6.7-4.2-12.6-10.2-14.9c-11.5-4.5-17.7-16.9-14.4-28.8c.4-1.3.6-2.8.6-4.3c0-8.8-7.2-16-16-16h-97.5c-12.6 0-25 3.7-35.5 10.7l-61.7 41.1c-11 7.4-25.9 4.4-33.3-6.7s-4.4-25.9 6.7-33.3l61.7-41.1c18.4-12.3 40-18.8 62.1-18.8H384c34.7 0 62.9 27.6 64 62c14.6 11.7 24 29.7 24 50c0 4.5-.5 8.8-1.3 13c15.4 11.7 25.3 30.2 25.3 51c0 6.5-1 12.8-2.8 18.7c11.6 11.8 18.8 27.8 18.8 45.5c0 35.3-28.6 64-64 64h-92.3c4.7 10.4 8.7 21.2 11.8 32.2l5.7 20c10.9 38.2-11.2 78.1-49.4 89M32 384c-17.7 0-32-14.3-32-32V128c0-17.7 14.3-32 32-32h64c17.7 0 32 14.3 32 32v224c0 17.7-14.3 32-32 32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsUp(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M323.8 34.8c-38.2-10.9-78.1 11.2-89 49.4l-5.7 20c-3.7 13-10.4 25-19.5 35l-51.3 56.4c-8.9 9.8-8.2 25 1.6 33.9s25 8.2 33.9-1.6l51.3-56.4c14.1-15.5 24.4-34 30.1-54.1l5.7-20c3.6-12.7 16.9-20.1 29.7-16.5s20.1 16.9 16.5 29.7l-5.7 20c-5.7 19.9-14.7 38.7-26.6 55.5c-5.2 7.3-5.8 16.9-1.7 24.9s12.3 13 21.3 13H448c8.8 0 16 7.2 16 16c0 6.8-4.3 12.7-10.4 15c-7.4 2.8-13 9-14.9 16.7s.1 15.8 5.3 21.7c2.5 2.8 4 6.5 4 10.6c0 7.8-5.6 14.3-13 15.7c-8.2 1.6-15.1 7.3-18 15.2s-1.6 16.7 3.6 23.3c2.1 2.7 3.4 6.1 3.4 9.9c0 6.7-4.2 12.6-10.2 14.9c-11.5 4.5-17.7 16.9-14.4 28.8c.4 1.3.6 2.8.6 4.3c0 8.8-7.2 16-16 16h-97.5c-12.6 0-25-3.7-35.5-10.7l-61.7-41.1c-11-7.4-25.9-4.4-33.3 6.7s-4.4 25.9 6.7 33.3l61.7 41.1c18.4 12.3 40 18.8 62.1 18.8H384c34.7 0 62.9-27.6 64-62c14.6-11.7 24-29.7 24-50c0-4.5-.5-8.8-1.3-13c15.4-11.7 25.3-30.2 25.3-51c0-6.5-1-12.8-2.8-18.7c11.6-11.8 18.8-27.8 18.8-45.5c0-35.3-28.6-64-64-64h-92.3c4.7-10.4 8.7-21.2 11.8-32.2l5.7-20c10.9-38.2-11.2-78.1-49.4-89M32 192c-17.7 0-32 14.3-32 32v224c0 17.7 14.3 32 32 32h64c17.7 0 32-14.3 32-32V224c0-17.7-14.3-32-32-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrashCan(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m170.5 51.6l-19 28.4h145l-19-28.4c-1.5-2.2-4-3.6-6.7-3.6h-93.7c-2.7 0-5.2 1.3-6.7 3.6zm147-26.6l36.7 55H424c13.3 0 24 10.7 24 24s-10.7 24-24 24h-8v304c0 44.2-35.8 80-80 80H112c-44.2 0-80-35.8-80-80V128h-8c-13.3 0-24-10.7-24-24s10.7-24 24-24h69.8l36.7-55.1C140.9 9.4 158.4 0 177.1 0h93.7c18.7 0 36.2 9.4 46.6 24.9zM80 128v304c0 17.7 14.3 32 32 32h224c17.7 0 32-14.3 32-32V128zm80 64v208c0 8.8-7.2 16-16 16s-16-7.2-16-16V192c0-8.8 7.2-16 16-16s16 7.2 16 16m80 0v208c0 8.8-7.2 16-16 16s-16-7.2-16-16V192c0-8.8 7.2-16 16-16s16 7.2 16 16m80 0v208c0 8.8-7.2 16-16 16s-16-7.2-16-16V192c0-8.8 7.2-16 16-16s16 7.2 16 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func User(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M304 128a80 80 0 1 0-160 0a80 80 0 1 0 160 0m-208 0a128 128 0 1 1 256 0a128 128 0 1 1-256 0M49.3 464h349.4c-8.9-63.3-63.3-112-129-112h-91.4c-65.7 0-120.1 48.7-129 112M0 482.3C0 383.8 79.8 304 178.3 304h91.4c98.5 0 178.3 79.8 178.3 178.3c0 16.4-13.3 29.7-29.7 29.7H29.7C13.3 512 0 498.7 0 482.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowMaximize(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.3 89.5C.1 91.6 0 93.8 0 96v320c0 35.3 28.7 64 64 64h384c35.3 0 64-28.7 64-64V96c0-35.3-28.7-64-64-64H64c-2.2 0-4.4.1-6.5.3c-9.2.9-17.8 3.8-25.5 8.2c-10.2 6-18.6 14.6-24.3 25c-3.9 7.3-6.5 15.4-7.4 24M48 224h416v192c0 8.8-7.2 16-16 16H64c-8.8 0-16-7.2-16-16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowMinimize(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24 432c-13.3 0-24 10.7-24 24s10.7 24 24 24h464c13.3 0 24-10.7 24-24s-10.7-24-24-24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowRestore(children ...ElementRenderer) *Fa6RegularIcon {
	return &Fa6RegularIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M432 48H208c-17.7 0-32 14.3-32 32v16h-48V80c0-44.2 35.8-80 80-80h224c44.2 0 80 35.8 80 80v224c0 44.2-35.8 80-80 80h-16v-48h16c17.7 0 32-14.3 32-32V80c0-17.7-14.3-32-32-32M48 448c0 8.8 7.2 16 16 16h256c8.8 0 16-7.2 16-16V256H48zm16-320h256c35.3 0 64 28.7 64 64v256c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64V192c0-35.3 28.7-64 64-64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
