package icons_8

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "1.0.0"
	hAttr          = "1em"
	viewbox        = "0 0 32 32"
	fill           = "currentColor"
)

type Icons8Icon struct {
	*SVGSVGElement
}

type Icons8IconFn func(children ...ElementRenderer) *Icons8Icon

var IconLookup = map[string]Icons8IconFn{
	"addShoppingCart":           AddShoppingCart,
	"addUser":                   AddUser,
	"adventures":                Adventures,
	"advertising":               Advertising,
	"airport":                   Airport,
	"alignCenter":               AlignCenter,
	"alignJustify":              AlignJustify,
	"alignLeft":                 AlignLeft,
	"alignRight":                AlignRight,
	"alphabeticalSorting":       AlphabeticalSorting,
	"alphabeticalSortingTwo":    AlphabeticalSortingTwo,
	"amex":                      Amex,
	"angleDown":                 AngleDown,
	"angleLeft":                 AngleLeft,
	"angleRight":                AngleRight,
	"angleUp":                   AngleUp,
	"archive":                   Archive,
	"areaChart":                 AreaChart,
	"arrowsLongDown":            ArrowsLongDown,
	"arrowsLongLeft":            ArrowsLongLeft,
	"arrowsLongRight":           ArrowsLongRight,
	"arrowsLongUp":              ArrowsLongUp,
	"asterisk":                  Asterisk,
	"audioFile":                 AudioFile,
	"babysRoom":                 BabysRoom,
	"bankCard":                  BankCard,
	"banknotes":                 Banknotes,
	"barChart":                  BarChart,
	"barcode":                   Barcode,
	"bed":                       Bed,
	"binoculars":                Binoculars,
	"bitcoin":                   Bitcoin,
	"bold":                      Bold,
	"bookmark":                  Bookmark,
	"box":                       Box,
	"briefcase":                 Briefcase,
	"britishPound":              BritishPound,
	"brush":                     Brush,
	"buy":                       Buy,
	"calculator":                Calculator,
	"camera":                    Camera,
	"cancel":                    Cancel,
	"cancelTwo":                 CancelTwo,
	"catFootprint":              CatFootprint,
	"checked":                   Checked,
	"checkedTwo":                CheckedTwo,
	"checkmark":                 Checkmark,
	"chevronDown":               ChevronDown,
	"chevronDownRound":          ChevronDownRound,
	"chevronLeft":               ChevronLeft,
	"chevronLeftRound":          ChevronLeftRound,
	"chevronRight":              ChevronRight,
	"chevronRightRound":         ChevronRightRound,
	"chevronUp":                 ChevronUp,
	"chevronUpRound":            ChevronUpRound,
	"circle":                    Circle,
	"circleNotch":               CircleNotch,
	"circleThin":                CircleThin,
	"clipboard":                 Clipboard,
	"codeFile":                  CodeFile,
	"colorDropper":              ColorDropper,
	"columns":                   Columns,
	"comments":                  Comments,
	"compress":                  Compress,
	"controller":                Controller,
	"copy":                      Copy,
	"copyright":                 Copyright,
	"createNew":                 CreateNew,
	"crop":                      Crop,
	"cut":                       Cut,
	"database":                  Database,
	"diningRoom":                DiningRoom,
	"diplomaOne":                DiplomaOne,
	"doctor":                    Doctor,
	"document":                  Document,
	"doubleLeft":                DoubleLeft,
	"doubleRight":               DoubleRight,
	"doubleUp":                  DoubleUp,
	"downArrow":                 DownArrow,
	"downRound":                 DownRound,
	"downSquared":               DownSquared,
	"download":                  Download,
	"downloadTwo":               DownloadTwo,
	"electrical":                Electrical,
	"electricity":               Electricity,
	"eraser":                    Eraser,
	"euro":                      Euro,
	"exclamationMark":           ExclamationMark,
	"export":                    Export,
	"fantasy":                   Fantasy,
	"fax":                       Fax,
	"female":                    Female,
	"file":                      File,
	"film":                      Film,
	"filter":                    Filter,
	"finishFlag":                FinishFlag,
	"fireExtinguisher":          FireExtinguisher,
	"folder":                    Folder,
	"footballTwo":               FootballTwo,
	"gender":                    Gender,
	"genderNeutralUser":         GenderNeutralUser,
	"genderqueer":               Genderqueer,
	"genericSorting":            GenericSorting,
	"genericSortingTwo":         GenericSortingTwo,
	"genericText":               GenericText,
	"gift":                      Gift,
	"googleWallet":              GoogleWallet,
	"gpsDevice":                 GpsDevice,
	"grid":                      Grid,
	"gridThree":                 GridThree,
	"gridTwo":                   GridTwo,
	"groupIcon":                 GroupIcon,
	"hdd":                       Hdd,
	"header":                    Header,
	"home":                      Home,
	"hospitalTwo":               HospitalTwo,
	"idea":                      Idea,
	"imageFile":                 ImageFile,
	"import":                    Import,
	"indent":                    Indent,
	"info":                      Info,
	"insertTable":               InsertTable,
	"ipad":                      Ipad,
	"iphone":                    Iphone,
	"italic":                    Italic,
	"japaneseYen":               JapaneseYen,
	"key":                       Key,
	"keyboard":                  Keyboard,
	"lastQuarter":               LastQuarter,
	"leftArrow":                 LeftArrow,
	"leftRound":                 LeftRound,
	"leftSquared":               LeftSquared,
	"levelDown":                 LevelDown,
	"levelUp":                   LevelUp,
	"library":                   Library,
	"list":                      List,
	"lock":                      Lock,
	"lockTwo":                   LockTwo,
	"male":                      Male,
	"mastercard":                Mastercard,
	"minus":                     Minus,
	"monitor":                   Monitor,
	"moon":                      Moon,
	"music":                     Music,
	"news":                      News,
	"notebook":                  Notebook,
	"numberedList":              NumberedList,
	"numericalSortingTwelve":    NumericalSortingTwelve,
	"numericalSortingTwentyOne": NumericalSortingTwentyOne,
	"oldTimeCamera":             OldTimeCamera,
	"openedFolder":              OpenedFolder,
	"organization":              Organization,
	"outdent":                   Outdent,
	"paragraph":                 Paragraph,
	"parallelTasks":             ParallelTasks,
	"paste":                     Paste,
	"paypal":                    Paypal,
	"pdf":                       Pdf,
	"pencil":                    Pencil,
	"phone":                     Phone,
	"picture":                   Picture,
	"pieChart":                  PieChart,
	"pinThree":                  PinThree,
	"plus":                      Plus,
	"powerpoint":                Powerpoint,
	"priceTag":                  PriceTag,
	"puzzle":                    Puzzle,
	"qrCode":                    QrCode,
	"questionMark":              QuestionMark,
	"recycling":                 Recycling,
	"refresh":                   Refresh,
	"removeUser":                RemoveUser,
	"resizeDiagonal":            ResizeDiagonal,
	"resizeFourDirections":      ResizeFourDirections,
	"resizeHorizontal":          ResizeHorizontal,
	"resizeVertical":            ResizeVertical,
	"rightArrow":                RightArrow,
	"rightRound":                RightRound,
	"rightSquared":              RightSquared,
	"rotateLeft":                RotateLeft,
	"rotateRight":               RotateRight,
	"rouble":                    Rouble,
	"roundedRectangle":          RoundedRectangle,
	"roundedRectangleFilled":    RoundedRectangleFilled,
	"rupee":                     Rupee,
	"search":                    Search,
	"sensor":                    Sensor,
	"services":                  Services,
	"settings":                  Settings,
	"share":                     Share,
	"shekel":                    Shekel,
	"shoppingCart":              ShoppingCart,
	"shutdown":                  Shutdown,
	"sort":                      Sort,
	"sortDown":                  SortDown,
	"sortLeft":                  SortLeft,
	"sortRight":                 SortRight,
	"sortUp":                    SortUp,
	"spy":                       Spy,
	"strikethrough":             Strikethrough,
	"stripe":                    Stripe,
	"student":                   Student,
	"subscript":                 Subscript,
	"superscript":               Superscript,
	"support":                   Support,
	"tags":                      Tags,
	"tasks":                     Tasks,
	"textHeight":                TextHeight,
	"textWidth":                 TextWidth,
	"ticket":                    Ticket,
	"timeline":                  Timeline,
	"todoList":                  TodoList,
	"translation":               Translation,
	"trash":                     Trash,
	"trophy":                    Trophy,
	"turkishLira":               TurkishLira,
	"umbrella":                  Umbrella,
	"underline":                 Underline,
	"undo":                      Undo,
	"unlockTwo":                 UnlockTwo,
	"upArrow":                   UpArrow,
	"upRound":                   UpRound,
	"upSquared":                 UpSquared,
	"upload":                    Upload,
	"uploadTwo":                 UploadTwo,
	"usDollar":                  UsDollar,
	"userFemale":                UserFemale,
	"userMale":                  UserMale,
	"videoCall":                 VideoCall,
	"videoFile":                 VideoFile,
	"visa":                      Visa,
	"won":                       Won,
	"word":                      Word,
	"xls":                       Xls,
}

func AddShoppingCart(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 7a1 1 0 0 0 0 2h2.22l2.624 10.5c.223.89 1.02 1.5 1.937 1.5h12.47c.903 0 1.67-.6 1.907-1.47L27.75 10h-2.094l-2.406 9H10.78L8.157 8.5A1.984 1.984 0 0 0 6.22 7zm18 14c-1.645 0-3 1.355-3 3s1.355 3 3 3s3-1.355 3-3s-1.355-3-3-3m-9 0c-1.645 0-3 1.355-3 3s1.355 3 3 3s3-1.355 3-3s-1.355-3-3-3m3-14v3h-3v2h3v3h2v-3h3v-2h-3V7zm-3 16c.564 0 1 .436 1 1c0 .564-.436 1-1 1c-.564 0-1-.436-1-1c0-.564.436-1 1-1m9 0c.564 0 1 .436 1 1c0 .564-.436 1-1 1c-.564 0-1-.436-1-1c0-.564.436-1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddUser(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 4c-3.854 0-7 3.146-7 7c0 2.41 1.23 4.552 3.094 5.813C6.527 18.343 4 21.88 4 26h2c0-4.43 3.57-8 8-8c1.376 0 2.654.358 3.78.97A7.993 7.993 0 0 0 16 24c0 4.406 3.594 8 8 8c4.406 0 8-3.594 8-8c0-4.406-3.594-8-8-8a7.98 7.98 0 0 0-4.688 1.53c-.442-.277-.92-.51-1.406-.718A7.018 7.018 0 0 0 21 11c0-3.854-3.146-7-7-7m0 2c2.773 0 5 2.227 5 5s-2.227 5-5 5s-5-2.227-5-5s2.227-5 5-5m10 12c3.326 0 6 2.674 6 6s-2.674 6-6 6s-6-2.674-6-6s2.674-6 6-6m-1 2v3h-3v2h3v3h2v-3h3v-2h-3v-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Adventures(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3m-1.125 2.063c.04-.005.084.003.125 0V6h2v-.938A10.958 10.958 0 0 1 26.938 15H26v2h.938A10.958 10.958 0 0 1 17 26.938V26h-2v.938A10.958 10.958 0 0 1 5.062 17H6v-2h-.938c.472-5.243 4.587-9.41 9.813-9.938zm7.22 4.843l-8.5 3.688l-3.69 8.5l8.5-3.688zM16 14.5c.8 0 1.5.7 1.5 1.5s-.7 1.5-1.5 1.5s-1.5-.7-1.5-1.5s.7-1.5 1.5-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Advertising(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m30 4.656l-1.28.375L5.812 12H2v8h3.813l1.968.594l-.03.093v.032c-.642 2.112.547 4.46 2.75 5.124c.01.003.022-.003.03 0c2.123.645 4.473-.53 5.126-2.75l.03-.094l13.033 3.97l1.28.374zm-2 2.688v17.312L6.28 18.03L6.157 18H4v-4h2.156l.125-.03L28 7.343zM9.687 21.187l4.094 1.22l-.03.093v.03c-.344 1.17-1.586 1.742-2.656 1.407c-1.17-.343-1.772-1.554-1.438-2.625v-.03z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Airport(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3c-1.645 0-3 1.355-3 3v7.344l-8.406 3.75l-.594.25v4.78L5.125 22L13 21.125v1.844l-2.563 1.717l-.437.28v4.253l1.188-.25L16 28l4.813.97l1.187.25v-4.25l-.438-.282L19 22.968v-1.843l7.875.875l1.125.125v-4.78l-.594-.25L19 13.344V6c0-1.645-1.355-3-3-3m0 2c.565 0 1 .435 1 1v8.656l.594.25L26 18.656v1.22L18.125 19L17 18.875v5.187l.438.282L20 26.062v.72l-3.813-.75L16 25.97l-.188.06l-3.812.75v-.72l2.563-1.717l.437-.282v-5.186L13.875 19L6 19.875v-1.22l8.406-3.75l.594-.25V6c0-.565.435-1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenter(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7v2h26V7zm4 4v2h18v-2zm-4 4v2h26v-2zm4 4v2h18v-2zm-4 4v2h26v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignJustify(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7v2h26V7zm0 4v2h26v-2zm0 4v2h26v-2zm0 4v2h26v-2zm0 4v2h26v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeft(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7v2h26V7zm0 4v2h18v-2zm0 4v2h26v-2zm0 4v2h18v-2zm0 4v2h26v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRight(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7v2h26V7zm8 4v2h18v-2zm-8 4v2h26v-2zm8 4v2h18v-2zm-8 4v2h26v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphabeticalSorting(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.188 5l-.22.656L6.032 11H6v.063l-.938 2.593l-.062.156V15h2v-.844L7.406 13h3.188L11 14.156V15h2v-1.188l-.063-.156L12 11.062V11h-.03l-1.94-5.344L9.814 5H8.185zM22 5v18.688l-2.594-2.594L18 22.5l4.28 4.313l.72.687l.72-.688L28 22.5l-1.406-1.406L24 23.687V5zM9 8.656L9.844 11H8.156zM5 17v2h5.563L5.28 24.28l-.28.314V27h8v-2H7.437l5.282-5.28l.28-.314V17z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphabeticalSortingTwo(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 5v2h5.563L5.28 12.28l-.28.314V15h8v-2H7.437l5.282-5.28l.28-.314V5zm17 0v18.688l-2.594-2.594L18 22.5l4.28 4.313l.72.687l.72-.688L28 22.5l-1.406-1.406L24 23.687V5zM8.187 17l-.218.656L6.03 23H6v.063l-.938 2.593l-.062.157V27h2v-.844L7.406 25h3.188L11 26.156V27h2v-1.188l-.063-.156L12 23.063V23h-.03l-1.94-5.344L9.814 17H8.185zM9 20.656L9.844 23H8.156z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Amex(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 6C3.355 6 2 7.355 2 9v14c0 1.645 1.355 3 3 3h22c1.645 0 3-1.355 3-3V9c0-1.645-1.355-3-3-3zm0 2h22c.565 0 1 .435 1 1v14c0 .565-.435 1-1 1H5c-.565 0-1-.435-1-1V9c0-.565.435-1 1-1m2.063 5.25L5 18.75h1.25l.406-1.25H9l.406 1.25h2.344v-4.125l1.5 4.125h1.094l1.53-4v4h1.095v-5.5h-1.783l-1.376 3.72l-1.373-3.72H10.5v5.22l-1.938-5.22h-1.5zm11.406 0v5.5h4.404l1.375-1.78l1.374 1.78H27L24.937 16L27 13.25h-1.5l-1.375 1.656L23 13.25zM7.75 14.344l.688 1.937H7.062l.688-1.935zm11.813.156h2.75l1.125 1.5l-1.25 1.656h-2.625v-1.093h2.5v-1.125h-2.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDown(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4.22 10.78l-1.44 1.44l12.5 12.5l.72.686l.72-.687l12.5-12.5l-1.44-1.44L16 22.564z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleLeft(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.03 4.28l-11 11l-.686.72l.687.72l11 11l1.44-1.44L10.187 16l10.28-10.28z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleRight(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.97 4.28l-1.44 1.44L21.814 16L11.53 26.28l1.44 1.44l11-11l.686-.72l-.687-.72l-11-11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleUp(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 6.594l-.72.687l-12.5 12.5l1.44 1.44L16 9.437l11.78 11.78l1.44-1.437l-12.5-12.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Archive(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3v26h20V3zm2 2h7v1h2V5h7v22H8zm7 2v2h2V7zm0 3v2h2v-2zm0 3v2.188c-1.156.418-2 1.52-2 2.812c0 1.645 1.355 3 3 3s3-1.355 3-3c0-1.292-.844-2.394-2-2.813V13zm1 4c.564 0 1 .436 1 1c0 .564-.436 1-1 1c-.564 0-1-.436-1-1c0-.564.436-1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AreaChart(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m28 4.063l-1.625 1.25l-4.625 3.625L16.156 8l-.375-.063l-.342.22l-5.688 3.78l-4.563-.906L4 10.78V28h24zm-2 4.093v5.375l-4.22 3.345l-5.468-1.813l-.468-.156l-.406.25l-5.563 3.72L6 17.31v-4.09l3.813.75l.406.092l.342-.218l5.657-3.78l5.624.936l.437.063l.345-.282zm0 7.938V26H6v-6.5l3.625 1.438l.5.187l.438-.28l5.624-3.75l5.5 1.843l.5.187l.438-.344z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsLongDown(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 4v20.063l-4.28-4.282l-1.44 1.407l6 6l.72.72l.72-.72l6-6l-1.44-1.406L17 24.063V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsLongLeft(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.813 9.28l-6 6l-.72.72l.72.72l6 6l1.406-1.44L7.936 17H28v-2H7.937l4.282-4.28l-1.408-1.44z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsLongRight(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.188 9.28l-1.407 1.44L24.063 15H4v2h20.063l-4.282 4.28l1.407 1.44l6-6l.72-.72l-.72-.72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsLongUp(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 4.094l-.72.718l-6 6l1.44 1.407L15 7.936V28h2V7.937l4.28 4.282l1.44-1.408l-6-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Asterisk(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m17.7 17l6.2 8.4l-2.4 1.6l-5.5-8.7l-5.5 8.7l-2.3-1.6l6.2-8.4l-9.3-2.4L6 12l9.1 3.2L14.5 5h3L17 15.2l9-3.2l.8 2.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AudioFile(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3v26h20V3zm2 2h16v22H8zm8 4.72v6.468A2.902 2.902 0 0 0 15 16c-1.645 0-3 1.355-3 3s1.355 3 3 3s3-1.355 3-3v-6.72l2.75.69l.5-1.94l-4-1zM15 18c.564 0 1 .436 1 1c0 .564-.436 1-1 1c-.564 0-1-.436-1-1c0-.564.436-1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BabysRoom(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3v2c1.496 0 2.298.52 2.688 1.094c-4.284.58-7.665 3.948-8.407 8.218C5.025 14.67 4 15.636 4 17c0 1.365 1.024 2.33 2.28 2.688C7.098 24.387 11.1 28 16 28s8.903-3.61 9.72-8.313C26.975 19.33 28 18.366 28 17c0-1.32-.963-2.26-2.156-2.656c-.75-3.958-3.718-7.142-7.563-8.063c.132-.17.267-.28.72-.28V4c-1.092 0-1.954.546-2.406 1.188c-.03.04-.036.084-.063.125C15.79 4.006 14.203 3 12 3m4 5c4.093 0 7.464 3.106 8 7.125l.125.875H25c.555 0 1 .445 1 1c0 .555-.445 1-1 1h-1l-.094.875C23.46 22.88 20.093 26 16 26s-7.46-3.12-7.906-7.125L8 18H7c-.555 0-1-.445-1-1c0-.555.445-1 1-1h1l.094-.875C8.54 11.12 11.907 8 16 8m-3.5 8a1.5 1.5 0 1 0 .001 3.001A1.5 1.5 0 0 0 12.5 16m7 0a1.5 1.5 0 1 0 .001 3.001A1.5 1.5 0 0 0 19.5 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BankCard(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 6C3.355 6 2 7.355 2 9v14c0 1.645 1.355 3 3 3h22c1.645 0 3-1.355 3-3V9c0-1.645-1.355-3-3-3zm0 2h22c.565 0 1 .435 1 1v2H5v2h23v10c0 .565-.435 1-1 1H5c-.565 0-1-.435-1-1V9c0-.565.435-1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Banknotes(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27.938 4.75L3.155 9H15l11.313-1.938L26.655 9h2.032zM2 10v16h28V10zm4.938 2h18.125a2.426 2.426 0 0 0-.063.5a2.5 2.5 0 0 0 2.5 2.5c.173 0 .337-.03.5-.063v6.126a2.426 2.426 0 0 0-.5-.063a2.5 2.5 0 0 0-2.5 2.5c0 .173.03.337.063.5H6.938c.033-.163.062-.327.062-.5A2.5 2.5 0 0 0 4.5 21c-.173 0-.337.03-.5.063v-6.125c.163.033.327.062.5.062A2.5 2.5 0 0 0 7 12.5c0-.173-.03-.337-.063-.5zM16 13c-2.75 0-5 2.25-5 5s2.25 5 5 5s5-2.25 5-5s-2.25-5-5-5m0 2c1.67 0 3 1.33 3 3s-1.33 3-3 3s-3-1.33-3-3s1.33-3 3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarChart(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 4v24h8V4zm2 2h4v20h-4zM3 10v18h8V10zm2 2h4v14H5zm7 4v12h8V16zm2 2h4v8h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Barcode(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 7v18h2V7zm4 0v18h6V7zm8 0v18h2V7zm4 0v18h4V7zm6 0v18h2V7zm4 0v18h2V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bed(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 6C4.355 6 3 7.355 3 9v6.78c-.61.552-1 1.342-1 2.22v9h5v-2h18v2h5v-9c0-.878-.39-1.668-1-2.22V9c0-1.645-1.355-3-3-3zm0 2h20c.555 0 1 .445 1 1v6h-2v-1c0-1.645-1.355-3-3-3h-4c-.767 0-1.467.3-2 .78a2.985 2.985 0 0 0-2-.78h-4c-1.645 0-3 1.355-3 3v1H5V9c0-.555.445-1 1-1m4 5h4c.555 0 1 .445 1 1v1H9v-1c0-.555.445-1 1-1m8 0h4c.555 0 1 .445 1 1v1h-6v-1c0-.555.445-1 1-1M5 17h22c.555 0 1 .445 1 1v7h-1v-2H5v2H4v-7c0-.555.445-1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Binoculars(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 5C9.032 5 7.46 6.44 7.125 8.313c-1.89.476-3.53 1.705-4.25 3.656h-.03L.718 17.936A6.972 6.972 0 0 0 0 21c0 3.854 3.146 7 7 7c3.472 0 6.365-2.552 6.906-5.875c.543.535 1.28.875 2.094.875c.814 0 1.55-.34 2.094-.875C18.634 25.448 21.528 28 25 28c3.854 0 7-3.146 7-7c0-.974-.2-1.906-.563-2.75l-2.28-6.375l-.032-.03v-.033c-.73-1.77-2.348-3.012-4.25-3.5C24.54 6.442 22.968 5 21 5c-1.857 0-3.362 1.285-3.813 3h-2.375c-.45-1.715-1.955-3-3.812-3m0 2c1.19 0 2 .81 2 2v1h6V9c0-1.19.81-2 2-2s2 .81 2 2v.906l.906.094c1.486.156 2.766 1.192 3.344 2.53c.01.022.022.042.03.064l.75 2.125a6.926 6.926 0 0 0-3.03-.72c-2.924 0-5.425 1.817-6.47 4.375A3.011 3.011 0 0 0 16 17a3.01 3.01 0 0 0-2.53 1.375C12.424 15.817 9.923 14 7 14c-1.072 0-2.09.253-3 .688l.75-2.032v-.03c.58-1.55 1.808-2.464 3.344-2.626L9 9.906V9c0-1.19.81-2 2-2m-4 9c2.773 0 5 2.227 5 5s-2.227 5-5 5s-5-2.227-5-5c0-.708.175-1.36.438-1.97c.016-.037.013-.086.03-.124A4.982 4.982 0 0 1 7 16m18 0c2.773 0 5 2.227 5 5s-2.227 5-5 5s-5-2.227-5-5s2.227-5 5-5m-9 3c.564 0 1 .436 1 1c0 .564-.436 1-1 1c-.564 0-1-.436-1-1c0-.564.436-1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bitcoin(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3v3H8v20h4v3h2v-3h2v3h2v-3h1.5c3.026 0 5.5-2.474 5.5-5.5a5.52 5.52 0 0 0-2.875-4.844A5.517 5.517 0 0 0 24 11.5C24 8.474 21.526 6 18.5 6H18V3h-2v3h-2V3zm-2 5h8.5c1.944 0 3.5 1.556 3.5 3.5S20.444 15 18.5 15H10zm0 9h9.5c1.944 0 3.5 1.556 3.5 3.5S21.444 24 19.5 24H10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bold(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 7c2.8 0 5 2.2 5 5c0 1.3-.5 2.4-1.3 3.3c1.9.7 3.3 2.5 3.3 4.7c0 2.8-2.2 5-5 5H7V7zm-7 8h7c1.7 0 3-1.3 3-3s-1.3-3-3-3H9zm0 8h9c1.7 0 3-1.3 3-3s-1.3-3-3-3H9zm7-18H5v22h13c3.9 0 7-3.1 7-7c0-2.1-1-4.1-2.5-5.4c.3-.8.5-1.7.5-2.6c0-3.9-3.1-7-7-7m-5 6h5c.6 0 1 .4 1 1s-.4 1-1 1h-5zm0 8h7c.6 0 1 .4 1 1s-.4 1-1 1h-7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 5v23l1.594-1.188L16 21.25l7.406 5.563L25 28V5zm2 2h14v17l-6.406-4.813L16 18.75l-.594.438L9 24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Box(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 5v6h1v16h22V11h1V5zm2 2h20v2H6zm1 4h18v14H7zm5.813 2A1 1 0 0 0 13 15h6a1 1 0 1 0 0-2h-6a1 1 0 0 0-.094 0a1 1 0 0 0-.094 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Briefcase(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3c-1.864 0-3.4 1.275-3.844 3H3v20h26V6h-9.156C19.4 4.275 17.864 3 16 3m0 2c.81 0 1.428.385 1.75 1h-3.5c.322-.615.94-1 1.75-1M5 8h22v9H5zm11 6a1 1 0 1 0 0 2a1 1 0 0 0 0-2M5 19h22v5H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BritishPound(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.688 5c-.27 0-.548.028-.813.063c-3.177.41-5.875 3.14-5.875 6.592c0 1.13.283 2.242.656 3.344H9v2h3.406c.198.554.36 1.102.5 1.655c.563 2.206.62 4.182-1.375 6.344H8v2h16v-5h-2v3h-7.938c1.393-2.307 1.318-4.746.782-6.845c-.102-.4-.225-.772-.344-1.156H20v-2h-6.156C13.39 13.692 13 12.53 13 11.655c0-3.232 3.3-5.543 6.375-4.312l.75-1.844c-.815-.326-1.63-.498-2.438-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Brush(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.813 4.03c-.837 0-1.68.337-2.313.97L9.812 17.5l-.624.594c-.75.14-1.467.51-2.032 1.125c-.833.908-.995 2.02-1.312 3.218c-.318 1.196-.72 2.526-1.688 4l-1 1.562H5c3.915 0 6.495-1.87 7.813-3.188l.03-.03c.603-.617.94-1.395 1.063-2.188l.5-.47l12.688-12.5v-.03a3.253 3.253 0 0 0 0-4.594c-.634-.633-1.445-.97-2.282-.97zm0 1.97c.313 0 .608.14.875.406c.533.534.533 1.248 0 1.782l-9.032 8.937l-1.812-1.813l9.062-8.906c.267-.266.593-.406.907-.406M13.405 16.72l1.813 1.81l-1.5 1.47l-1.814-1.813l1.5-1.468zM9.97 20.03a2.31 2.31 0 0 1 1.467.595c.778.712.796 1.954-.03 2.78c-.87.87-2.408 1.986-4.657 2.407c.48-1.05.815-2.03 1.03-2.843c.334-1.255.578-2.116.845-2.407c.347-.38.838-.546 1.344-.532z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Buy(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 7a1 1 0 0 0 0 2h2.22l2.624 10.5c.223.89 1.02 1.5 1.937 1.5h12.47c.903 0 1.67-.6 1.907-1.47L27.75 10h-2.094l-2.406 9H10.78L8.157 8.5A1.984 1.984 0 0 0 6.22 7zm18 14c-1.645 0-3 1.355-3 3s1.355 3 3 3s3-1.355 3-3s-1.355-3-3-3m-9 0c-1.645 0-3 1.355-3 3s1.355 3 3 3s3-1.355 3-3s-1.355-3-3-3m3-14v5h-3l4 4l4-4h-3V7zm-3 16c.564 0 1 .436 1 1c0 .564-.436 1-1 1c-.564 0-1-.436-1-1c0-.564.436-1 1-1m9 0c.564 0 1 .436 1 1c0 .564-.436 1-1 1c-.564 0-1-.436-1-1c0-.564.436-1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calculator(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3v26h20V3zm2 2h16v22H8zm2 2v6h12V7zm2 2h8v2h-8zm-1 6v2h2v-2zm4 0v2h2v-2zm4 0v2h2v-2zm-8 4v2h2v-2zm4 0v2h2v-2zm4 0v2h2v-2zm-8 4v2h2v-2zm4 0v2h2v-2zm4 0v2h2v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.5 6l-.313.406L10 8H3v18h26V8h-7l-1.188-1.594L20.5 6zm1 2h7l1.188 1.594L21 10h6v14H5V10h6l.313-.406zM8 11a1 1 0 1 0 0 2a1 1 0 0 0 0-2m8 0c-3.302 0-6 2.698-6 6s2.698 6 6 6s6-2.698 6-6s-2.698-6-6-6m0 2c2.22 0 4 1.78 4 4c0 2.22-1.78 4-4 4c-2.22 0-4-1.78-4-4c0-2.22 1.78-4 4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cancel(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3m0 2c6.087 0 11 4.913 11 11s-4.913 11-11 11S5 22.087 5 16S9.913 5 16 5m-3.78 5.78l-1.44 1.44L14.564 16l-3.782 3.78l1.44 1.44L16 17.437l3.78 3.78l1.44-1.437L17.437 16l3.78-3.78l-1.437-1.44L16 14.564l-3.78-3.782z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CancelTwo(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3m0 2c6.087 0 11 4.913 11 11c0 2.726-.99 5.206-2.625 7.125L9.03 7.47A10.954 10.954 0 0 1 16 5M7.625 8.875L22.97 24.53A10.954 10.954 0 0 1 16 27C9.913 27 5 22.087 5 16c0-2.726.99-5.206 2.625-7.125"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CatFootprint(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 5c-1.07 0-2.002.608-2.594 1.438C9.314 7.268 9 8.34 9 9.5s.314 2.233.906 3.063C10.498 13.393 11.43 14 12.5 14c1.07 0 2.002-.608 2.594-1.438c.592-.83.906-1.902.906-3.062s-.314-2.233-.906-3.063C14.502 5.607 13.57 5 12.5 5M16 9.5c0 1.16.314 2.233.906 3.063c.592.83 1.525 1.437 2.594 1.437c1.07 0 2.002-.608 2.594-1.438c.592-.83.906-1.902.906-3.062s-.314-2.233-.906-3.063C21.502 5.607 20.57 5 19.5 5c-1.07 0-2.002.608-2.594 1.438C16.314 7.268 16 8.34 16 9.5M12.5 7c.312 0 .656.156.97.594c.31.437.53 1.133.53 1.906c0 .773-.22 1.47-.53 1.906c-.314.438-.658.594-.97.594c-.312 0-.656-.156-.97-.594c-.31-.437-.53-1.133-.53-1.906c0-.773.22-1.47.53-1.906c.314-.438.658-.594.97-.594m7 0c.312 0 .656.156.97.594c.31.437.53 1.133.53 1.906c0 .773-.22 1.47-.53 1.906c-.314.438-.658.594-.97.594c-.312 0-.656-.156-.97-.594c-.31-.437-.53-1.133-.53-1.906c0-.773.22-1.47.53-1.906c.314-.438.658-.594.97-.594m-12 5c-1.07 0-2.002.608-2.594 1.438C4.314 14.268 4 15.34 4 16.5s.314 2.233.906 3.063C5.498 20.393 6.43 21 7.5 21c1.07 0 2.002-.608 2.594-1.438c.592-.83.906-1.902.906-3.062s-.314-2.233-.906-3.063C9.502 12.607 8.57 12 7.5 12m17 0c-1.07 0-2.002.608-2.594 1.438c-.592.83-.906 1.902-.906 3.062s.314 2.233.906 3.063C22.498 20.393 23.43 21 24.5 21c1.07 0 2.002-.608 2.594-1.438c.592-.83.906-1.902.906-3.062s-.314-2.233-.906-3.063C26.502 12.607 25.57 12 24.5 12m-17 2c.312 0 .656.156.97.594c.31.437.53 1.133.53 1.906c0 .773-.22 1.47-.53 1.906c-.314.438-.658.594-.97.594c-.312 0-.656-.156-.97-.594C6.22 17.97 6 17.273 6 16.5c0-.773.22-1.47.53-1.906c.314-.438.658-.594.97-.594m17 0c.312 0 .656.156.97.594c.31.437.53 1.133.53 1.906c0 .773-.22 1.47-.53 1.906c-.314.438-.658.594-.97.594c-.312 0-.656-.156-.97-.594c-.31-.437-.53-1.133-.53-1.906c0-.773.22-1.47.53-1.906c.314-.438.658-.594.97-.594M16 16c-1.333 0-2.263.865-2.72 1.625c-.454.76-.734 1.392-1 1.656c-.153.157-1.118.52-2.124 1.033c-.503.257-1.01.6-1.437 1.125c-.43.525-.72 1.25-.72 2.063c0 1.92 1.58 3.5 3.5 3.5c.866 0 1.77-.28 2.655-.53C15.043 26.216 16 26 16 26s.957.216 1.844.47c.886.25 1.79.53 2.656.53c1.92 0 3.5-1.58 3.5-3.5c0-.794-.293-1.52-.72-2.03c-.425-.513-.937-.83-1.436-1.095c-1-.53-1.983-.95-2.125-1.094c-.24-.237-.51-.89-.97-1.654C18.29 16.86 17.34 16 16 16m0 2c.66 0 .733.16 1.03.656c.3.495.524 1.336 1.25 2.063c.825.82 1.867 1 2.626 1.404c.38.202.67.417.844.625c.173.207.25.396.25.75c0 .84-.66 1.5-1.5 1.5c-.29 0-1.224-.22-2.094-.47c-.87-.247-1.59-.53-2.406-.53c-.816 0-1.536.283-2.406.53c-.87.25-1.804.47-2.094.47c-.84 0-1.5-.66-1.5-1.5c0-.403.084-.61.25-.814c.166-.203.472-.404.844-.593c.744-.38 1.78-.53 2.625-1.375c.733-.737.953-1.605 1.25-2.095c.292-.49.362-.625 1.03-.625z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Checked(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3C8.8 3 3 8.8 3 16s5.8 13 13 13s13-5.8 13-13c0-1.4-.188-2.794-.688-4.094L26.688 13.5c.2.8.313 1.6.313 2.5c0 6.1-4.9 11-11 11S5 22.1 5 16S9.9 5 16 5c3 0 5.694 1.194 7.594 3.094L25 6.688C22.7 4.388 19.5 3 16 3m11.28 4.28L16 18.563l-4.28-4.28l-1.44 1.437l5 5l.72.686l.72-.687l12-12l-1.44-1.44z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckedTwo(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 4v24h24V12.187l-2 2V26H6V6h19.813l2-2zm23.28 3.28L16 18.563l-4.28-4.28l-1.44 1.437l5 5l.72.686l.72-.687l12-12l-1.44-1.44z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Checkmark(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28.28 6.28L11 23.563l-7.28-7.28l-1.44 1.437l8 8l.72.686l.72-.687l18-18l-1.44-1.44z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDown(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.906 6.594l-.718.687l-3.907 3.907l-.686.72l.687.687l13 13l.72.718l.72-.718l13-13l.686-.688l-.687-.72l-3.907-3.905l-.72-.686l-.687.687L16 15.688L7.594 7.28zm-.03 2.843l8.405 8.376l.72.687l.72-.688l8.405-8.375l2.438 2.438L16 23.47L4.437 11.874z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownRound(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3m0 2c6.087 0 11 4.913 11 11s-4.913 11-11 11S5 22.087 5 16S9.913 5 16 5m-5.28 7.78l-1.44 1.44l6 6l.72.686l.72-.687l6-6l-1.44-1.44L16 18.064l-5.28-5.282z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeft(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.75 2.594l-.72.687l-12 12l-.686.72l.687.72l12 12l.72.686l.72-.687l3.593-3.626l.687-.688l-.688-.718L16.375 16l7.688-7.688l.687-.718l-.688-.688L20.47 3.28zm0 2.844l2.188 2.187l-7.688 7.656l-.72.72l.72.72l7.688 7.655l-2.188 2.188L9.187 16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeftRound(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3m0 2c6.087 0 11 4.913 11 11s-4.913 11-11 11S5 22.087 5 16S9.913 5 16 5m1.78 4.28l-6 6l-.686.72l.687.72l6 6l1.44-1.44L13.937 16l5.28-5.28z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRight(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.25 2.594l-.72.687l-3.59 3.626l-.688.688l.688.718L15.625 16l-7.688 7.688l-.687.718l.688.688l3.593 3.625l.72.686l.72-.687l12-12l.686-.72l-.687-.72l-12-12l-.72-.686zm0 2.844L22.813 16L12.25 26.563l-2.188-2.188l7.688-7.656l.72-.72l-.72-.72l-7.688-7.655l2.188-2.188z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRightRound(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3m0 2c6.087 0 11 4.913 11 11s-4.913 11-11 11S5 22.087 5 16S9.913 5 16 5m-1.78 4.28l-1.44 1.44L18.064 16l-5.282 5.28l1.44 1.44l6-6l.686-.72l-.687-.72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUp(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 5.688l-.72.718l-13 13l-.686.688l.687.718L6.19 24.72l.718.686l.688-.687L16 16.312l8.406 8.406l.688.686l.718-.687l3.907-3.907l.686-.72l-.687-.687l-13-13l-.72-.718zm0 2.843l11.563 11.595l-2.438 2.438l-8.406-8.375L16 13.5l-.72.688l-8.405 8.374l-2.438-2.437L16 8.53z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpRound(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3m0 2c6.087 0 11 4.913 11 11s-4.913 11-11 11S5 22.087 5 16S9.913 5 16 5m0 6.094l-.72.687l-6 6l1.44 1.44L16 13.937l5.28 5.28l1.44-1.437l-6-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circle(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4C9.384 4 4 9.384 4 16s5.384 12 12 12s12-5.384 12-12S22.616 4 16 4m0 2c5.535 0 10 4.465 10 10s-4.465 10-10 10S6 21.535 6 16S10.465 6 16 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleNotch(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 4.18v2.022c4.56.93 8 4.97 8 9.798c0 5.514-4.486 10-10 10S6 21.514 6 16c0-4.83 3.44-8.87 8-9.798v-2.02C8.334 5.136 4 10.065 4 16c0 6.617 5.383 12 12 12s12-5.383 12-12c0-5.934-4.334-10.863-10-11.82"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleThin(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4C9.38 4 4 9.38 4 16s5.38 12 12 12s12-5.38 12-12S22.62 4 16 4m0 1c6.08 0 11 4.92 11 11s-4.92 11-11 11S5 22.08 5 16S9.92 5 16 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clipboard(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3c-1.26 0-2.152.89-2.594 2H6v23h20V5h-7.406C18.152 3.89 17.26 3 16 3m0 2c.555 0 1 .445 1 1v1h3v2h-8V7h3V6c0-.555.445-1 1-1M8 7h2v4h12V7h2v19H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeFile(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3v26h20V9.594l-.28-.313l-6-6l-.314-.28H6zm2 2h10v6h6v16H8zm12 1.438L22.563 9H20zM16 13l-2 12h2l2-12zm-3.78 2.375l-2.5 3l-.533.625l.532.625l2.5 3l1.56-1.25L11.813 19l1.968-2.375zm7.56 0l-1.56 1.25L20.187 19l-1.97 2.375l1.563 1.25l2.5-3l.532-.625l-.53-.625l-2.5-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ColorDropper(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.656 3.03c-1.108 0-2.222.41-3.062 1.25l-2.782 2.814l-1-1L16.407 7.5l1 1l-9.5 9.5c-1.04 1.04-1.633 1.792-2 2.47a5.288 5.288 0 0 0-.562 1.686c-.068.402-.1.62-.344 1.063c-.244.44-.733 1.075-1.72 2.06l-.686.72l.687.72l2 2l.72.686l.72-.687c.96-.96 1.573-1.427 2-1.657c.425-.23.65-.275 1.06-.344a5.081 5.081 0 0 0 1.69-.564c.685-.378 1.466-.998 2.53-2.062l9.5-9.5l1 1l1.406-1.406l-1-1l2.813-2.782c1.68-1.68 1.68-4.444 0-6.125a4.318 4.318 0 0 0-3.064-1.25zm0 2c.592 0 1.166.23 1.625.69c.92.918.92 2.36 0 3.28l-2.78 2.78l-3.28-3.28L23 5.72c.46-.46 1.065-.69 1.656-.69m-5.843 4.876l3.28 3.28l-9.5 9.502c-.985.986-1.65 1.475-2.093 1.718c-.442.244-.66.276-1.063.344c-.402.068-.98.165-1.656.53c-.46.25-1.086.848-1.686 1.376l-.75-.75c.547-.62 1.147-1.25 1.406-1.718c.378-.687.493-1.31.563-1.72c.07-.41.113-.604.343-1.03c.23-.427.696-1.072 1.657-2.032l9.5-9.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Columns(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 5v22h22V5zm2 2h8v18H7zm10 0h8v18h-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comments(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 6v20h9.563l2.718 2.72l.72.686l.72-.687L19.437 26H29V6zm2 2h22v16h-8.406l-.313.28L16 26.563l-2.28-2.28l-.314-.282H5zm4 3v2h14v-2zm0 4v2h14v-2zm0 4v2h10v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compress(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28.28 2.28L19 11.595V3h-2v12h12v-2h-8.594l9.313-9.28l-1.44-1.44zM3 17v2h8.594L2.28 28.28l1.44 1.44L13 20.405V29h2V17z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Controller(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 7C9.617 7 4.03 9.063 4.03 9.063l-.593.218l-.062.596l-1.344 10.25c-.363 2.835 1.667 5.48 4.5 5.844c2.64.34 5.01-1.44 5.627-3.97h7.688c.62 2.528 2.986 4.31 5.625 3.97c2.834-.366 4.864-3.01 4.5-5.845l-1.345-10.25l-.063-.594l-.593-.216S22.38 7 16 7m0 2c5.486 0 10.007 1.524 10.75 1.78l1.22 9.626a3.159 3.159 0 0 1-2.75 3.563c-1.763.225-3.368-.99-3.595-2.75l-.03-.346L21.47 20H10.53l-.124.874l-.03.344c-.227 1.762-1.832 2.976-3.595 2.75a3.162 3.162 0 0 1-2.75-3.564l1.22-9.625C5.994 10.524 10.515 9 16 9m-7 3v2H7v2h2v2h2v-2h2v-2h-2v-2zm13 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-2 2a1 1 0 1 0 0 2a1 1 0 0 0 0-2m4 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-2 2a1 1 0 1 0 0 2a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copy(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 4v20h7v-2H6V6h12v1h2V4zm8 4v20h16V8zm2 2h12v16H14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copyright(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3m0 2c6.087 0 11 4.913 11 11s-4.913 11-11 11S5 22.087 5 16S9.913 5 16 5m-.094 5c-3.324 0-6 2.676-6 6s2.676 6 6 6c2.4 0 4.45-1.44 5.406-3.47l-1.812-.843C18.855 19.058 17.506 20 15.906 20c-2.276 0-4-1.724-4-4s1.724-4 4-4c1.6 0 2.95.942 3.594 2.313l1.813-.844c-.956-2.03-3.007-3.47-5.407-3.47z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreateNew(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25 4.03c-.765 0-1.517.3-2.094.876L13 14.78l-.22.22l-.06.313l-.69 3.5l-.31 1.468l1.467-.31l3.5-.69l.313-.06l.22-.22l9.874-9.906A2.968 2.968 0 0 0 25 4.032zm0 1.94c.235 0 .464.12.688.343c.446.446.446.928 0 1.375L16 17.374l-1.72.344l.345-1.72l9.688-9.688c.223-.223.452-.343.687-.343zM4 8v20h20V14.812l-2 2V26H6V10h9.188l2-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crop(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 4v4H4v2h4v14h14v4h2v-4h4v-2H11.437L22 11.437V21h2V9.437l3.72-3.718l-1.44-1.44L22.563 8H11v2h9.563L10 20.563V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cut(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 6c-2.197 0-4 1.803-4 4s1.803 4 4 4c.957 0 1.81-.37 2.5-.938l5.22 2.97l-5.158 2.937C7.864 18.374 6.982 18 6 18c-2.197 0-4 1.803-4 4s1.803 4 4 4s4-1.803 4-4c0-.494-.115-.97-.28-1.406L30 9h-4l-10.22 5.844l-6.06-3.438c.072-.194.144-.387.186-.594c.054-.26.094-.537.094-.812c0-2.197-1.803-4-4-4m0 2c1.116 0 2 .884 2 2s-.884 2-2 2s-2-.884-2-2c0-.14.005-.275.03-.406A1.997 1.997 0 0 1 6 8m13.094 8.813l-2 1.093L26 23h4zM6 20c1.116 0 2 .884 2 2s-.884 2-2 2s-2-.884-2-2c0-.14.005-.275.03-.406A1.997 1.997 0 0 1 6 20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Database(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 4v24h20V4zm2 2h16v5H8zm0 7h16v6H8zm0 8h16v5H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiningRoom(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.406 4L17.78 8.594c-1.152 1.153-1.152 3.065 0 4.22L16 14.56l-1.28-1.28l-7.5-7.5l-.72-.688l-.72.687a4.569 4.569 0 0 0 0 6.44l6.282 6.28l-7.78 7.78l1.437 1.44l7.78-7.782l1.188 1.156l.718.718l.688-.718l1.78-1.782l8.407 8.407l1.44-1.44l-9.126-9.093L17.437 16l1.75-1.78h.032l.03.03c1.158 1.1 3.02 1.105 4.156-.03l.094-.126l4.5-4.5l-1.406-1.406L22 12.78c-.435.436-.937.445-1.375.033l5.28-5.313L24.5 6.094l-5.313 5.28c-.412-.437-.403-.94.032-1.374l4.593-4.594zM6.78 8.22l9.69 9.686l-1.064 1.063l-8.187-8.19c-.742-.74-.764-1.676-.44-2.56z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiplomaOne(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3c-.624 0-1.248.213-1.78.594L12.56 4.75L10.656 5h-.03l-.032.03a3.39 3.39 0 0 0-2.563 2.564l-.03.03v.032l-.25 1.938l-1.156 1.5l-.032.03v.032c-.698 1.118-.73 2.56.032 3.625L7.78 16.44l.314 1.718l-3.22 4.907l-1.03 1.53h4.781l1.156 2.688l.72 1.72l1.03-1.563l3.158-4.75c.846.35 1.806.38 2.624 0l3.157 4.75L21.5 29l.72-1.72l1.155-2.686h4.781l-1.03-1.532L24 18.313l.25-1.875l1.156-1.656l.032-.03v-.03c.698-1.12.73-2.528-.032-3.595L24.25 9.47l-.375-1.876h.03c0-.025-.027-.04-.03-.063c-.18-1.307-1.217-2.37-2.53-2.53h-.032l-1.875-.25l-1.657-1.156A3.078 3.078 0 0 0 16 3m0 2.03c.23 0 .458.07.625.19l1.78 1.28l.22.156l.25.032L21.063 7h.03c.448.05.764.365.813.813v.062l.407 2.22l.03.217l.157.188l1.28 1.78c.24.335.27.894-.03 1.376l-1.25 1.75l-.156.22l-.032.25L22 18.062v.03a.966.966 0 0 1-.156.438l-.063.032v.032a.883.883 0 0 1-.593.312h-.062l-2.28.407l-.25.03l-.19.157l-1.78 1.28c-.333.24-.924.27-1.406-.03l-1.626-1.25l-.188-.156l-.28-.032L10.937 19h-.032a.872.872 0 0 1-.687-.438a.98.98 0 0 1-.126-.375v-.062l-.406-2.28l-.032-.25l-.156-.19l-1.28-1.78c-.24-.333-.27-.924.03-1.406l1.25-1.626l.156-.188l.03-.28l.283-2.095c.002-.014.026-.015.03-.03a1.398 1.398 0 0 1 1-1c.015-.004.016-.028.03-.03l2.095-.282l.25-.032l.22-.156l1.78-1.28c.167-.12.396-.19.625-.19zm6.906 15.22l1.532 2.344H22.03l-.25.625l-.686 1.593l-2.125-3.25l.468-.344l1.968-.345v.03c.025 0 .04-.027.063-.03a2.973 2.973 0 0 0 1.436-.625zm-13.812.03c.427.382.967.648 1.562.72h.03l1.908.25l.437.344l-2.124 3.218l-.687-1.593l-.25-.626H7.56l1.532-2.313z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Doctor(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.438 6l-.282.47S5 11.652 5 20v1h6.094A4.923 4.923 0 0 0 11 22c0 2.75 2.25 5 5 5s5-2.25 5-5c0-.344-.027-.675-.094-1H27v-1c0-4.61-.776-7.99-1.563-10.22c-.786-2.228-1.625-3.374-1.625-3.374L23.5 6H8.437zm1.218 2h12.72c.145.208.573.732 1.186 2.47c.65 1.84 1.232 4.73 1.344 8.53H20c-.915-1.21-2.376-2-4-2s-3.085.79-4 2H7.094c.228-6.65 2.232-10.43 2.562-11M15 10v2h-2v2h2v2h2v-2h2v-2h-2v-2zm1 9c1.67 0 3 1.33 3 3s-1.33 3-3 3s-3-1.33-3-3s1.33-3 3-3m0 2a1 1 0 1 0 0 2a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Document(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3v26h20V9.594l-.28-.313l-6-6l-.314-.28H6zm2 2h10v6h6v16H8zm12 1.438L22.563 9H20zM11 13v2h10v-2zm0 4v2h10v-2zm0 4v2h10v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleLeft(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.906 4.78l-10.5 10.5l-.718.72l.718.72l10.5 10.5l1.406-1.44L7.533 16l9.78-9.78l-1.406-1.44zm7 0l-10.5 10.5l-.72.72l.72.72l10.5 10.5l1.407-1.44L14.53 16l9.783-9.78z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleRight(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.094 4.78L7.688 6.22l9.78 9.78l-9.78 9.78l1.406 1.44l10.5-10.5l.718-.72l-.718-.72zm7 0l-1.407 1.44L24.47 16l-9.782 9.78l1.406 1.44l10.5-10.5l.718-.72l-.718-.72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleUp(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 4.688l-.72.718l-11.5 11.5l1.44 1.407L16 7.53l10.78 10.782l1.44-1.406l-11.5-11.5zm0 7l-.72.718l-11.5 11.5l1.44 1.407L16 14.53l10.78 10.783l1.44-1.407l-11.5-11.5l-.72-.72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownArrow(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 4v20.063L8.22 17.28l-1.44 1.44l8.5 8.5l.72.686l.72-.687l8.5-8.5l-1.44-1.44L17 24.063V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownRound(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3m0 2c6.087 0 11 4.913 11 11s-4.913 11-11 11S5 22.087 5 16S9.913 5 16 5m-1 4v10.28l-4-4l-1.406 1.44l5.687 5.686l.72.72l.72-.72l5.686-5.687L21 15.28l-4 4V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownSquared(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 5v22h22V5zm2 2h18v18H7zm3.72 5.78l-1.44 1.44l6 6l.72.686l.72-.687l6-6l-1.44-1.44L16 18.064l-5.28-5.282z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 4v16.563L9.72 15.28l-1.44 1.44l7 7l.72.686l.72-.687l7-7l-1.44-1.44L17 20.564V4zM7 27v2h18v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadTwo(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 5c-4.12 0-7.36 3.13-7.813 7.125a4.895 4.895 0 0 0-3.843 3.22C1.884 16.054 0 18.248 0 21c0 3.324 2.676 6 6 6h20c3.324 0 6-2.676 6-6c0-1.76-.855-3.336-2.094-4.438c-.232-3.514-3.035-6.318-6.562-6.5C22.14 7.133 19.378 5 16 5m0 2c2.762 0 4.97 1.77 5.75 4.28l.22.72H23c2.755 0 5 2.245 5 5v.5l.406.313A4.07 4.07 0 0 1 30 21c0 2.276-1.724 4-4 4H6c-2.276 0-4-1.724-4-4c0-2.02 1.45-3.588 3.28-3.906l.657-.125l.125-.657C6.362 14.963 7.556 14 9 14h1v-1c0-3.37 2.63-6 6-6m-1 5v6.563l-2.28-2.282l-1.44 1.44l4 4l.72.686l.72-.687l4-4l-1.44-1.44L17 18.563V12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Electrical(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m22 3.594l-4 3.969l-2.28-2.282l-1.44 1.44l.75.75l-5.124 5.124a3.124 3.124 0 0 0 0 4.406l1.844 1.844l-7.47 7.437l1.44 1.44l7.436-7.47L15 22.094a3.124 3.124 0 0 0 4.406 0l5.125-5.125l.75.75l1.44-1.44L24.437 14l3.968-4L27 8.594l-4 3.97L19.437 9l3.97-4zm-5.563 5.28l6.688 6.688L18 20.689c-.388.387-1.206.387-1.594 0l-5.093-5.094c-.388-.388-.388-1.206 0-1.594l5.124-5.125z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Electricity(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.875 4l-6.906 7.313l-.845.906l1.063.624l7.187 4.344L12 24.563V20h-2v8h8v-2h-4.563l8.282-8.28l.905-.907l-1.094-.657l-7.25-4.375l6.033-6.405z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eraser(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.906 4.094c-.805-.002-1.64.272-2.28.843v.032l-.032.03L4.906 16.594c-1.212 1.212-1.204 3.184-.062 4.468l.03.032h.032l6 6c1.213 1.212 3.184 1.204 4.47.062v-.03l.03-.032L27 15.5c1.267-1.267 1.306-3.288.094-4.5l-6-6a3.06 3.06 0 0 0-2.188-.906m-.03 2.03c.318 0 .618.088.81.282l6.002 6c.387.388.44 1.154-.094 1.688l-5.032 5.03l-7.656-7.655l5.063-5.033l.03-.03c.253-.21.57-.28.875-.282zm-7.407 6.782l7.655 7.656l-5.094 5.094l-.03.032c-.517.428-1.308.38-1.688 0L6.345 19.75c-.016-.018-.017-.044-.032-.063c-.41-.517-.374-1.313 0-1.687z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Euro(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 4c-4.738 0-8.587 3.887-9.688 9H6v2h2.063c-.023.328-.063.666-.063 1c0 .334.04.672.063 1H6v2h2.313c1.1 5.113 4.95 9 9.687 9c2.71 0 5.17-1.303 6.938-3.344l-1.532-1.312C21.954 25.02 20.07 26 18 26c-3.502 0-6.59-2.898-7.625-7H19v-2h-8.97c-.025-.33-.03-.66-.03-1c0-.34.005-.67.03-1H19v-2h-8.625C11.41 8.898 14.498 6 18 6c2.07 0 3.954.98 5.406 2.656l1.532-1.312C23.168 5.304 20.708 4 18 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationMark(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 4v16h6V4zm2 2h2v12h-2zm-2 16v6h6v-6zm2 2h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Export(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 4v24h20v-8l-2 2v4H8V6h16v4l2 2V4zm16.406 7L21 12.406L23.563 15h-9.657v2h9.656L21 19.594L22.406 21l4.313-4.28l.686-.72l-.687-.72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fantasy(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.875 2.563l-.688.75l-1.687 1.78h-3.594v3.5l-1.72 1.813l-.686.72l.688.687l1.5 1.5L3.03 25l-.717.72l.718.686l2.564 2.563l.718.718l.688-.72l11.688-11.655l1.5 1.5l.687.687l.72-.688l1.81-1.718h3.5V13.5l1.782-1.688l.75-.687l-.718-.72l-1.814-1.81v-3.5h-3.5L21.594 3.28zm.03 2.874l1.376 1.375l.314.282h2.312v2.312l.282.313l1.375 1.374l-1.344 1.28l-.314.282v2.438h-2.312l-.282.28l-1.406 1.345l-.843-.845l3.843-3.813l.72-.687l-.72-.72l-2.562-2.56l-.688-.69l-.718.69l-3.813 3.842l-.844-.843l1.345-1.406l.28-.282v-2.31h2.44l.28-.313l1.28-1.343zm-.25 4.782l1.126 1.155l-15.468 15.47l-1.156-1.126l15.5-15.5zM19 21v1h-1v2h1v1h2v-1h1v-2h-1v-1zm6 2v2h-2v2h2v2h2v-2h2v-2h-2v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fax(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5v5h-2V8H4v18h2v1c0 1.645 1.355 3 3 3s3-1.355 3-3v-1h16V10h-4V5zm2 2h8v5h-8zm-8 3h2v14H6zm4 2h2v2h12v-2h2v12H10zm3 4v2h2v-2zm4 0v2h2v-2zm4 0v2h2v-2zm-8 4v2h2v-2zm4 0v2h2v-2zm4 0v2h2v-2zM8 26h2v1c0 .555-.445 1-1 1c-.555 0-1-.445-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Female(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.72 4.28l-1.44 1.44L21.563 9l-2.968 2.97A8.935 8.935 0 0 0 13 10c-4.96 0-9 4.04-9 9s4.04 9 9 9s9-4.04 9-9a8.94 8.94 0 0 0-1.97-5.594l2.97-2.97l3.28 3.283l1.44-1.44L24.437 9l3.28-3.28l-1.437-1.44L23 7.563l-3.28-3.28zM13 12c3.878 0 7 3.122 7 7s-3.122 7-7 7s-7-3.122-7-7s3.122-7 7-7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func File(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3v26h20V9.594l-.28-.313l-6-6l-.314-.28H6zm2 2h10v6h6v16H8zm12 1.438L22.563 9H20z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Film(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 4v24h24V4zm2 2h2v1h2V6h12v1h2V6h2v20h-2v-1h-2v1H10v-1H8v1H6zm2 3v2h2V9zm14 0v2h2V9zM8 13v2h2v-2zm14 0v2h2v-2zM8 17v2h2v-2zm14 0v2h2v-2zM8 21v2h2v-2zm14 0v2h2v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filter(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 4v2.344l.22.28l7.78 9.72V28l1.594-1.188l4-3L19 23.5v-7.156l7.78-9.72l.22-.28V4zm2.28 2h17.44l-7.19 9h-3.06zM15 17h2v5.5L15 24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FinishFlag(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 4v24h2v-8h20V4zm2 2h3v3h3V6h3v3h3V6h3v3h3v3h-3v3h3v3h-3v-3h-3v3h-3v-3h-3v3h-3v-3H7v-3h3V9H7zm3 6v3h3v-3zm3 0h3V9h-3zm3 0v3h3v-3zm3 0h3V9h-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FireExtinguisher(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4c-1.292 0-2.394.844-2.813 2H12c-2.745 0-5 2.255-5 5h2c0-1.655 1.345-3 3-3h1v2.47c-.32.237-.733.575-1.22 1.06C10.954 12.36 10 13.5 10 15v13h12V15c0-1.5-.953-2.64-1.78-3.47c-.487-.485-.9-.823-1.22-1.06v-.283l4.844.813l1.156.188V4.812L23.844 5l-5.094.844C18.292 4.77 17.234 4 16 4m0 2c.555 0 1 .445 1 1v3h-2V7c0-.555.445-1 1-1m7 1.188v1.625l-4-.688v-.25l4-.688zM14.375 12h3.25c.15.105.578.39 1.156.97c.673.67 1.22 1.53 1.22 2.03v11h-8V15c0-.5.547-1.36 1.22-2.03c.577-.58 1.004-.865 1.155-.97zM14 17v2h4v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3v26h20V15.437l1.72-1.718l.28-.314V3zm2 2h14v8.406l.28.313L24 15.436V27H8zm16 0h2v7.563l-1 1l-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FootballTwo(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3m0 2c.602 0 1.175.063 1.75.156L16 6.406l-1.75-1.25C14.822 5.066 15.402 5 16 5m-4.188.844l3.594 2.625l.594.436l.594-.437l3.593-2.626c1.6.663 3.01 1.68 4.125 2.968l-1.375 4.282l-.218.687l.593.408l3.625 2.656a11.005 11.005 0 0 1-1.532 4.843h-5.282l-.218.688L18.5 26.72a11.119 11.119 0 0 1-5.094-.033l-1.375-4.28l-.217-.688h-5.22c-.875-1.44-1.396-3.1-1.53-4.876l3.593-2.625l.594-.408l-.22-.687l-1.405-4.25a10.989 10.989 0 0 1 4.188-3.03zM16 10.094l-.594.437l-4.562 3.314l-.563.437l.22.69l1.75 5.342l.22.688h7.061l.22-.688l1.75-5.343l.22-.69l-.564-.436l-4.562-3.313l-.594-.436zm9.75.812a10.922 10.922 0 0 1 1.094 3.406l-1.782-1.28zm-19.53.063l.686 2.06l-1.75 1.283c.183-1.19.533-2.315 1.063-3.344zM16 12.593l3.375 2.437l-1.28 3.97h-4.19l-1.28-3.97L16 12.595zm5.594 11.094h2.25a10.82 10.82 0 0 1-2.938 2.156l.688-2.157zm-13.438.03h2.188l.687 2.095a11.075 11.075 0 0 1-2.874-2.094z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gender(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 3v2h3.563l-3.375 3.406A6.962 6.962 0 0 0 18 7c-1.87 0-3.616.74-4.938 2.063a6.94 6.94 0 0 0 .001 9.875c.87.87 1.906 1.495 3.062 1.812c.114-.087.242-.178.344-.28a3.45 3.45 0 0 0 .874-1.532a4.906 4.906 0 0 1-2.875-1.407C13.524 16.588 13 15.336 13 14s.525-2.586 1.47-3.53C15.412 9.523 16.664 9 18 9s2.587.525 3.53 1.47A4.956 4.956 0 0 1 23 14c0 .865-.245 1.67-.656 2.406c.096.516.156 1.058.156 1.594c0 .498-.042.99-.125 1.47c.2-.163.378-.348.563-.532C24.26 17.614 25 15.87 25 14c0-1.53-.504-2.984-1.406-4.188L27 6.438V10h2V3zm-6.125 8.25c-.114.087-.242.178-.344.28c-.432.434-.714.96-.874 1.533c1.09.14 2.085.616 2.875 1.406c.945.943 1.47 2.195 1.47 3.53s-.525 2.586-1.47 3.53C16.588 22.477 15.336 23 14 23s-2.587-.525-3.53-1.47A4.948 4.948 0 0 1 9 18c0-.865.245-1.67.656-2.406A8.789 8.789 0 0 1 9.5 14c0-.498.042-.99.125-1.47c-.2.163-.377.348-.563.533C7.742 14.384 7 16.13 7 18c0 1.53.504 2.984 1.406 4.188L6.72 23.875l-2-2l-1.44 1.406l2 2l-2 2l1.44 1.44l2-2l2 2l1.405-1.44l-2-2l1.688-1.686A6.932 6.932 0 0 0 14 25c1.87 0 3.616-.74 4.938-2.063C20.26 21.616 21 19.87 21 18s-.74-3.614-2.063-4.938c-.87-.87-1.906-1.495-3.062-1.812"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GenderNeutralUser(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 5c-3.854 0-7 3.146-7 7c0 2.41 1.23 4.552 3.094 5.813C8.527 19.343 6 22.88 6 27h2c0-4.43 3.57-8 8-8s8 3.57 8 8h2c0-4.12-2.527-7.658-6.094-9.188A7.02 7.02 0 0 0 23 12c0-3.854-3.146-7-7-7m0 2c2.773 0 5 2.227 5 5s-2.227 5-5 5s-5-2.227-5-5s2.227-5 5-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Genderqueer(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M26.594 3.97L23.5 7.06l-2.28-2.28l-1.44 1.437l2.282 2.28l-3.312 3.313c-.79-.524-1.734-.81-2.75-.81s-1.96.286-2.75.813L11.937 10.5l1.782-1.78l-1.44-1.44l-1.78 1.782L7.437 6H11V4H4v7h2V7.437L9.063 10.5L7.28 12.28l1.44 1.44l1.78-1.783l1.313 1.313A4.919 4.919 0 0 0 11 16c0 2.406 1.727 4.438 4 4.906v4.657l-2.28-2.282l-1.44 1.44l4 4l.72.686l.72-.687l4-4l-1.44-1.44L17 25.563v-4.656c2.273-.468 4-2.5 4-4.906c0-1.016-.286-1.96-.813-2.75L23.5 9.937l2.28 2.282l1.44-1.44l-2.282-2.28l3.093-3.094l-1.436-1.437zM16 13c1.67 0 3 1.33 3 3s-1.33 3-3 3s-3-1.33-3-3s1.33-3 3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GenericSorting(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 5v2h12V5zm17 0v18.688l-2.594-2.594L17 22.5l4.28 4.313l.72.687l.72-.688L27 22.5l-1.406-1.406L23 23.687V5zM4 9v2h10V9zm0 4v2h8v-2zm0 4v2h6v-2zm0 4v2h4v-2zm0 4v2h2v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GenericSortingTwo(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 5v2h2V5zm17 0v18.688l-2.594-2.594L17 22.5l4.28 4.313l.72.687l.72-.688L27 22.5l-1.406-1.406L23 23.687V5zM4 9v2h4V9zm0 4v2h6v-2zm0 4v2h8v-2zm0 4v2h10v-2zm0 4v2h12v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GenericText(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 6L8 26h2l2.094-6h7.812L22 26h2L17 6zm1 2.844L19.188 18h-6.375z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gift(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5c-1.645 0-3 1.355-3 3c0 .353.073.684.188 1H4v6h1v13h22V15h1V9h-5.188c.115-.316.188-.647.188-1c0-1.645-1.355-3-3-3c-1.75 0-2.94 1.33-3.72 2.438c-.103.148-.188.292-.28.437c-.092-.145-.177-.29-.28-.438C14.94 6.33 13.75 5 12 5m0 2c.626 0 1.436.67 2.063 1.563c.152.217.13.23.25.437H12c-.565 0-1-.435-1-1s.435-1 1-1m8 0c.565 0 1 .435 1 1s-.435 1-1 1h-2.313c.12-.206.098-.22.25-.438C18.564 7.672 19.375 7 20 7M6 11h20v2h-9v-1h-2v1H6zm1 4h18v11h-8V16h-2v10H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleWallet(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 3.5c-.58 0-1.143.16-1.656.438a3.48 3.48 0 0 0-1.407 4.718c.006.01.45.945.875 2.47c-2.123-3.064-4.39-4.675-4.5-4.75c-.017-.016-.045-.018-.062-.032a1 1 0 0 0-.094-.063c-.613-.484-1.358-.78-2.156-.78c-.71 0-1.383.22-1.97.626a3.493 3.493 0 0 0-.905 4.844s.254.448.53 1.092c-.354-.28-.708-.532-1.06-.75C8.18 10.442 7.005 10 6 10c-1.645 0-3 1.355-3 3c0 1.008.614 1.788 1.28 2.344c.67.555 1.454 1.024 2.25 1.562c1.484 1 2.976 2.2 3.814 4.47c-.223.522-.402.927-.406.936c-.937 1.662-.333 3.77 1.343 4.72a3.516 3.516 0 0 0 1.72.468c1.26 0 2.41-.686 3.03-1.78c.165-.29 1.114-2.056 1.783-4.595c.42 1.26.784 2.9.688 4.813a1 1 0 0 0 .064.593c.003.023-.004.044 0 .065c.196 1.104.9 2.102 2 2.593a3.485 3.485 0 0 0 4.627-1.749c.235-.53 2.31-5.34 2.31-11.343c0-6.145-2.19-10.327-2.437-10.782A3.492 3.492 0 0 0 22 3.5m0 2a1.5 1.5 0 0 1 1.313.78c.088.164 2.187 4.057 2.187 9.814c0 5.694-2.038 10.337-2.125 10.53c-.25.558-.802.876-1.375.876a1.55 1.55 0 0 1-.625-.125c-.756-.338-1.087-1.244-.75-2c.028-.064 1.64-3.755 1.844-8.344a1 1 0 0 0 .03-.53c.002-.136 0-.267 0-.405c0-4.95-1.795-8.34-1.814-8.375a1.499 1.499 0 0 1 .594-2.032c.23-.125.477-.188.72-.188zm-9 2c.355 0 .693.136.97.375a1 1 0 0 0 .03.03c.08.074.155.16.22.25c.09.135 2.28 3.322 2.28 8.033c0 4.636-2.098 8.404-2.188 8.562c-.275.487-.79.75-1.312.75c-.25 0-.517-.055-.75-.188a1.521 1.521 0 0 1-.563-2.062c.027-.048 1.813-3.25 1.813-7.063c0-3.732-1.703-6.317-1.72-6.343a1.477 1.477 0 0 1 .376-2.063c.258-.177.55-.28.844-.28zm4.97 4.47c.947 1.304 1.865 2.898 2.5 4.81c-.06 1.61-.3 3.095-.595 4.345c-.547-1.905-1.207-3.07-1.438-3.438c.043-.497.063-.98.063-1.5c0-1.57-.225-2.988-.53-4.218zM6 12c.1 0 1.36.308 2.53 1.03c1.096.677 2.225 1.67 2.97 3v.157c0 .885-.114 1.724-.28 2.5c-1.053-1.588-2.4-2.653-3.564-3.437c-.858-.58-1.628-1.05-2.093-1.438C5.096 13.425 5 13.242 5 13c0-.565.435-1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GpsDevice(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 4.438l-.906 2.187l-8 19l-.907 2.125l2.157-.813L16 24.063l7.656 2.875l2.157.813l-.907-2.124l-8-19L16 4.436zm0 5.093l6.188 14.72l-5.844-2.187l-.344-.125l-.344.125l-5.844 2.188L16 9.53z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grid(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 6v20h26V6zm2 2h4v4H5zm6 0h4v4h-4zm6 0h4v4h-4zm6 0h4v4h-4zM5 14h4v4H5zm6 0h4v4h-4zm6 0h4v4h-4zm6 0h4v4h-4zM5 20h4v4H5zm6 0h4v4h-4zm6 0h4v4h-4zm6 0h4v4h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridThree(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 6v20h24V6zm2 2h5v4H6zm7 0h13v4H13zm-7 6h5v4H6zm7 0h13v4H13zm-7 6h5v4H6zm7 0h13v4H13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridTwo(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 5v22h22V5zm2 2h8v8H7zm10 0h8v8h-8zM7 17h8v8H7zm10 0h8v8h-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GroupIcon(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 7c-3.302 0-6 2.698-6 6a6 6 0 0 0 2.47 4.844C2.83 19.154 1 21.864 1 25h2c0-3.326 2.674-6 6-6s6 2.674 6 6h2c0-3.326 2.674-6 6-6s6 2.674 6 6h2c0-3.136-1.83-5.846-4.47-7.156A5.998 5.998 0 0 0 29 13c0-3.302-2.698-6-6-6s-6 2.698-6 6a6 6 0 0 0 2.47 4.844a8.048 8.048 0 0 0-3.47 3.28a8.048 8.048 0 0 0-3.47-3.28A5.998 5.998 0 0 0 15 13c0-3.302-2.698-6-6-6m0 2c2.22 0 4 1.78 4 4c0 2.22-1.78 4-4 4c-2.22 0-4-1.78-4-4c0-2.22 1.78-4 4-4m14 0c2.22 0 4 1.78 4 4c0 2.22-1.78 4-4 4c-2.22 0-4-1.78-4-4c0-2.22 1.78-4 4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hdd(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.22 6l-.19.75l-3 12l-.03.125V26h26v-7.125l-.03-.125l-3-12l-.19-.75zm1.56 2h16.44l2.5 10H5.28zM5 20h22v4H5zm19 1a1 1 0 1 0 0 2a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Header(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 5v2h1v18H6v2h8v-2h-1v-6h6v6h-1v2h8v-2h-1V7h1V5h-8v2h1v6h-6V7h1V5zm3 2h2v8h10V7h2v18h-2v-8H11v8H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 2.594l-.72.687l-13 13l1.44 1.44L5 16.437V28h9V18h4v10h9V16.437l1.28 1.282l1.44-1.44l-13-13zm0 2.844l9 9V26h-5V16h-8v10H7V14.437l9-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HospitalTwo(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 5v22h22V5zm2 2h18v18H7zm4 3v12h2v-5h6v5h2V10h-2v5h-6v-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Idea(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.813 2.406L5.405 3.812L7.5 5.906L8.906 4.5zm18.375 0L23.093 4.5L24.5 5.906l2.094-2.093zM16 3.03c-.33.004-.664.023-1 .064c-.01 0-.02-.002-.03 0c-4.056.465-7.284 3.742-7.845 7.78c-.448 3.25.892 6.197 3.125 8.095a5.238 5.238 0 0 1 1.75 3.03v6h2.28c.348.597.983 1 1.72 1s1.372-.403 1.72-1H20v-4h.094v-1.188c0-1.466.762-2.944 2-4.093C23.75 17.06 25 14.705 25 12c0-4.94-4.066-9.016-9-8.97m0 2c3.865-.054 7 3.11 7 6.97c0 2.094-.97 3.938-2.313 5.28l.032.032A7.792 7.792 0 0 0 18.279 22h-4.374c-.22-1.714-.955-3.373-2.344-4.563c-1.767-1.5-2.82-3.76-2.468-6.312c.437-3.15 2.993-5.683 6.125-6.03a6.91 6.91 0 0 1 .78-.064zM2 12v2h3v-2zm25 0v2h3v-2zM7.5 20.094l-2.094 2.093l1.407 1.407L8.905 21.5zm17 0L23.094 21.5l2.093 2.094l1.407-1.407zM14 24h4v2h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageFile(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3v26h20V9.594l-.28-.313l-6-6l-.314-.28H6zm2 2h10v6h6v16H8zm12 1.438L22.563 9H20zM21.094 14a1 1 0 1 0 0 2a1 1 0 0 0 0-2M14 15.594l-.72.687l-4 4l1.44 1.44L14 18.437l2.28 2.28l.72.688l.72-.687L19 19.437l2.28 2.28l1.44-1.437l-3-3l-.72-.686l-.72.687L17 18.563l-2.28-2.28z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Import(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 4v24h20v-9h-2v7H8V6h16v7h2V4zm11.5 7l-4.313 4.28l-.687.72l.688.72L17.5 21l1.406-1.406L16.313 17H28v-2H16.312l2.594-2.594z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Indent(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7v2h26V7zm0 4v2h19v-2zm26 0l-5 5l5 5zM3 15v2h19v-2zm0 4v2h19v-2zm0 4v2h26v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Info(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3m0 2c6.087 0 11 4.913 11 11s-4.913 11-11 11S5 22.087 5 16S9.913 5 16 5m-1 5v2h2v-2zm0 4v8h2v-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InsertTable(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 5v22h22V5zm2 2h5v5H7zm7 0h4v5h-4zm6 0h5v5h-5zM7 14h5v4H7zm7 0h4v4h-4zm6 0h5v4h-5zM7 20h5v5H7zm7 0h4v5h-4zm6 0h5v5h-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ipad(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 6C3.355 6 2 7.355 2 9v14c0 1.645 1.355 3 3 3h22c1.645 0 3-1.355 3-3V9c0-1.645-1.355-3-3-3zm0 2h22c.555 0 1 .445 1 1v14c0 .555-.445 1-1 1H5c-.555 0-1-.445-1-1V9c0-.555.445-1 1-1m1 7a1 1 0 1 0 0 2a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Iphone(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 4C9.355 4 8 5.355 8 7v18c0 1.645 1.355 3 3 3h10c1.645 0 3-1.355 3-3V7c0-1.645-1.355-3-3-3zm0 2h10c.555 0 1 .445 1 1v18c0 .555-.445 1-1 1H11c-.555 0-1-.445-1-1V7c0-.555.445-1 1-1m5 17a1 1 0 1 0 0 2a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Italic(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.75 5l-.063.938l-.187 3L11.437 10h2.001l-.876 12h-2l-.062.938l-.188 3L10.22 27h10.03l.063-.938l.187-3L20.563 22h-2.001l.875-12h2.001l.062-.938l.188-3L21.78 5zm1.875 2h6l-.063 1h-2l-.062.938l-1 14L16.437 24h2.001l-.063 1h-6l.063-1h2l.062-.938l1-14L15.563 8h-2.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JapaneseYen(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.188 5l6.875 11H10v2h5v2h-5v2h5v5h2v-5h5v-2h-5v-2h5v-2h-4.063l6.875-11h-2.375L16 15.344L9.562 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Key(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 3c-4.945 0-9 4.055-9 9c0 .52.085.978.156 1.438L3.28 21.28l-.28.314V29h7v-3h3v-3h3v-2.97c1.18.58 2.555.97 4 .97c4.945 0 9-4.055 9-9s-4.055-9-9-9m0 2c3.855 0 7 3.145 7 7s-3.145 7-7 7a7.37 7.37 0 0 1-3.406-.875l-.25-.125H14v3h-3v3H8v3H5v-4.563l7.906-7.937l.375-.344l-.092-.53c-.1-.6-.188-1.137-.188-1.626c0-3.855 3.145-7 7-7zm2 3a2 2 0 1 0-.001 3.999A2 2 0 0 0 22 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyboard(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7c-1.093 0-2 .907-2 2v14c0 1.093.907 2 2 2h26c1.093 0 2-.907 2-2V9c0-1.093-.907-2-2-2zm0 2h26v14H3zm2 2v2h2v-2zm4 0v2h2v-2zm4 0v2h2v-2zm4 0v2h2v-2zm4 0v2h2v-2zm4 0v2h2v-2zM5 15v2h4v-2zm6 0v2h2v-2zm4 0v2h2v-2zm4 0v2h2v-2zm4 0v2h4v-2zM5 19v2h4v-2zm6 0v2h10v-2zm12 0v2h4v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LastQuarter(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4C9.384 4 4 9.384 4 16s5.384 12 12 12s12-5.384 12-12S22.616 4 16 4m0 2v20c-5.535 0-10-4.465-10-10S10.465 6 16 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftArrow(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.28 6.78l-8.5 8.5l-.686.72l.687.72l8.5 8.5l1.44-1.44L7.936 17H28v-2H7.937l6.782-6.78z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftRound(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3m0 2c6.087 0 11 4.913 11 11s-4.913 11-11 11S5 22.087 5 16S9.913 5 16 5m-.72 4.594L9.595 15.28l-.72.72l.72.72l5.687 5.686L16.72 21l-4-4H23v-2H12.72l4-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftSquared(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 5v22h22V5zm2 2h18v18H7zm10.78 2.28l-6 6l-.686.72l.687.72l6 6l1.44-1.44L13.937 16l5.28-5.28z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LevelDown(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 5v2h15v17.063l-4.28-4.282l-1.44 1.44l6 6l.72.686l.72-.687l6-6l-1.44-1.44L22 24.063V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LevelUp(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21 4.094l-.72.687l-6 6l1.44 1.44L20 7.936V25H5v2h17V7.937l4.28 4.282l1.44-1.44l-6-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Library(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 3.906l-.375.156l-12 5L3 9.345V12h2v11H3v5h26v-5h-2V12h2V9.344l-.625-.28l-12-5.002zm0 2.188L25.375 10H6.625zM7 12h2v11H7zm4 0h2v11h-2zm4 0h2v11h-2zm4 0h2v11h-2zm4 0h2v11h-2zM5 25h22v1H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func List(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 5v6h6V5zm2 2h2v2H6zm6 0v2h15V7zm-8 6v6h6v-6zm2 2h2v2H6zm6 0v2h15v-2zm-8 6v6h6v-6zm2 2h2v2H6zm6 0v2h15v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3c-3.845 0-7 3.155-7 7v3H6v16h20V13h-3v-3c0-3.845-3.155-7-7-7m0 2c2.755 0 5 2.245 5 5v3H11v-3c0-2.755 2.245-5 5-5M8 15h16v12H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockTwo(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3c-3.845 0-7 3.155-7 7v2.875A9.968 9.968 0 0 0 6 20c0 5.51 4.49 10 10 10s10-4.49 10-10a9.968 9.968 0 0 0-3-7.125V10c0-3.845-3.155-7-7-7m0 2c2.755 0 5 2.245 5 5v1.375C19.525 10.515 17.826 10 16 10s-3.525.516-5 1.375V10c0-2.755 2.245-5 5-5m0 7c4.43 0 8 3.57 8 8s-3.57 8-8 8s-8-3.57-8-8s3.57-8 8-8m0 6a2 2 0 0 0-2 2c0 .74.403 1.373 1 1.72V25h2v-3.28c.597-.347 1-.98 1-1.72a2 2 0 0 0-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Male(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 4v2h7.563l-5.97 5.97A8.921 8.921 0 0 0 13 10c-4.96 0-9 4.04-9 9s4.04 9 9 9s9-4.04 9-9a8.94 8.94 0 0 0-1.97-5.594L26 7.436V15h2V4zm-4 8c3.878 0 7 3.122 7 7s-3.122 7-7 7s-7-3.122-7-7s3.122-7 7-7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mastercard(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 6C3.355 6 2 7.355 2 9v14c0 1.645 1.355 3 3 3h22c1.645 0 3-1.355 3-3V9c0-1.645-1.355-3-3-3zm0 2h22c.565 0 1 .435 1 1v14c0 .565-.435 1-1 1H5c-.565 0-1-.435-1-1V9c0-.565.435-1 1-1m7 1c-3.854 0-7 3.146-7 7s3.146 7 7 7a6.942 6.942 0 0 0 4-1.25A6.945 6.945 0 0 0 20 23c3.854 0 7-3.146 7-7s-3.146-7-7-7c-1.488 0-2.87.458-4 1.25A6.942 6.942 0 0 0 12 9m0 2c1.282 0 2.43.496 3.313 1.28a6.013 6.013 0 0 0-.97 1.72h2.22c.14.318.272.65.343 1h-2.812A5.884 5.884 0 0 0 14 16c0 .342.04.674.094 1h2.812c-.07.35-.204.682-.343 1h-2.22c.225.63.56 1.205.97 1.72C14.43 20.503 13.283 21 12 21c-2.773 0-5-2.227-5-5s2.227-5 5-5m8 0c2.774 0 5 2.226 5 5s-2.226 5-5 5a4.96 4.96 0 0 1-2.5-.688A6.926 6.926 0 0 0 19 16c0-1.63-.56-3.12-1.5-4.313A4.96 4.96 0 0 1 20 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3m0 2c6.087 0 11 4.913 11 11s-4.913 11-11 11S5 22.087 5 16S9.913 5 16 5m-6 10v2h12v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Monitor(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 6v18h13v2h-5v2h12v-2h-5v-2h13V6zm2 2h24v14H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Moon(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3c-2.922 0-5.534.98-7.688 2.594l-.687.5l.375.75C7.98 6.804 8 6.8 8 7c0 .555-.445 1-1 1c-.2 0-.195-.02-.156 0l-.75-.375l-.5.688C3.98 10.466 3 13.078 3 16c0 4.977 2.798 9.423 7.03 11.594l.658.343l.53-.53c.283-.283.515-.407.782-.407c.467 0 .954.445 1 .906l.094.782l.75.093c.687.1 1.36.22 2.156.22c6.695 0 12.15-4.983 12.906-11.47l.063-.78l-.69-.28c-.03-.013-.28-.314-.28-.47c0-.156.25-.457.28-.47l.69-.28l-.064-.78A12.946 12.946 0 0 0 16 3m-.844 2.03c.525.35 2.032 1.474 3.406 3.69l.282.467h.75a.79.79 0 0 1 .812.813c0 .267-.065.346-.406.688l-.438.437l.188.594c.08.236.19.636.313 1l.187.56l.594.095c.138.023.156.063.156.22c0 .1-.054.167-.156.217l-.532.282v.594c0 .566.094 1.08.094 1.312c0 3.118-.872 5.546-1.875 7.188l-.28.468l.25.5c.1.2.094.172.094.25c0 .224.036.188-.188.188h-.782l-.31.375c-.913 1.105-1.7 1.7-2.095 1.968c-.22-.03-.426-.06-.658-.094C14.082 25.804 13.195 25 12 25c-.59 0-1.067.23-1.5.5C7.21 23.587 5 20.024 5 16c0-2.274.754-4.256 1.938-6H7c1.645 0 3-1.355 3-3v-.063c1.526-1.035 3.23-1.745 5.156-1.906zm3.156.22c4.423.936 7.752 4.507 8.47 9.03c-.454.46-.782 1.024-.782 1.72s.328 1.26.78 1.72c-.71 4.49-4 8.03-8.374 9c.074-.082.144-.135.22-.22c1.067-.132 1.968-1.002 1.968-2.094c0-.267-.09-.394-.157-.593c1.1-1.935 1.97-4.534 1.97-7.813c0-.405-.034-.674-.063-.938c.365-.39.656-.874.656-1.468c0-.762-.49-1.38-1.125-1.78c-.02-.072-.07-.14-.094-.22c.324-.43.626-.927.626-1.594c0-1.388-1.103-2.447-2.437-2.656a17.183 17.183 0 0 0-1.657-2.094zM13.72 8.344a3.008 3.008 0 0 0-2.876 2.156c-.476 1.586.6 3.124 2.062 3.563l.594-1.907c-.537-.16-.874-.648-.75-1.062a1.01 1.01 0 0 1 1.156-.72l.375-1.968a3.077 3.077 0 0 0-.56-.062zm-5 7.437c-1.146.492-1.875 1.784-1.314 3c.007.006-.003.026 0 .032c.007.014.025.018.032.032c.527 1.114 1.843 1.804 3.062 1.28l-.78-1.843c-.152.066-.4-.045-.5-.28v-.03l-.033-.032c.02.036.078-.212.313-.313l-.78-1.844zm4.467 3.314c-1.215.912-1.67 2.688-.687 4c.912 1.216 2.688 1.703 4 .718l-1.188-1.625c-.287.216-.93.103-1.218-.28c-.216-.29-.072-.93.312-1.22l-1.22-1.593z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Music(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m27 3.78l-1.188.25l-16 3L9 7.157v13.407A3.893 3.893 0 0 0 7 20c-2.197 0-4 1.803-4 4s1.803 4 4 4s4-1.803 4-4V12.812l14-2.624v7.374A3.93 3.93 0 0 0 23 17c-2.197 0-4 1.803-4 4s1.803 4 4 4s4-1.803 4-4zm-2 2.41v2l-14 2.624v-2zM23 19c1.116 0 2 .884 2 2s-.884 2-2 2s-2-.884-2-2s.884-2 2-2M7 22c1.116 0 2 .884 2 2s-.884 2-2 2s-2-.884-2-2s.884-2 2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func News(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 5v18c0 2.21 1.79 4 4 4h18c2.21 0 4-1.79 4-4V12h-6V5zm2 2h16v16c0 .73.22 1.41.563 2H7c-1.19 0-2-.81-2-2zm2 2v5h12V9zm2 2h8v1H9zm14 3h4v9c0 1.19-.81 2-2 2s-2-.81-2-2zM7 15v2h5v-2zm7 0v2h5v-2zm-7 3v2h5v-2zm7 0v2h5v-2zm-7 3v2h5v-2zm7 0v2h5v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notebook(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 6v13.563l-2.28 2.312a2.452 2.452 0 0 0-.72 1.72A2.416 2.416 0 0 0 4.406 26h23.188A2.417 2.417 0 0 0 30 23.594c0-.64-.266-1.266-.72-1.72L27 19.564V6zm2 2h18v11H7zm-.563 13h19.125l2.313 2.28c.077.08.125.204.125.314c0 .24-.166.406-.406.406H4.406A.387.387 0 0 1 4 23.594a.47.47 0 0 1 .125-.313z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumberedList(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.97 3l-.314.344S4.9 4 4.406 4v2c.68 0 1.15-.276 1.594-.53V10h2V3zM11 6v2h17V6zm-4.5 6A2.497 2.497 0 0 0 4 14.5v.5h2v-.5c0-.217.283-.5.5-.5c.217 0 .5.283.5.5l-.03.03l-.064.064l-2.593 2.5l-.313.28V19h5v-2H7.28l.876-.875l.125-.094l-.03-.03c.502-.41.75-1.02.75-1.5A2.499 2.499 0 0 0 6.5 12m4.5 3v2h17v-2zm-7 6v2h1.375l-.25.406l-.125.22V25h1.5c.217 0 .5.283.5.5c0 .217-.283.5-.5.5H4v2h2.5C7.883 28 9 26.883 9 25.5c0-1.005-.678-1.696-1.53-2.094l.405-.687l.125-.25V21zm7 3v2h17v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumericalSortingTwelve(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.594 5l-.156.78s-.166.576-.563 1.157C7.478 7.52 6.98 8 6 8v2c1.376 0 2.32-.675 3-1.406V15h2V5zM22 5v18.688l-2.594-2.594L18 22.5l4.28 4.313l.72.687l.72-.688L28 22.5l-1.406-1.406L24 23.687V5zM8.5 17A3.512 3.512 0 0 0 5 20.5v.5h2v-.5c0-.876.624-1.5 1.5-1.5h1c.876 0 1.5.624 1.5 1.5c0 .456-.353.98-.938 1.344c-1.234.76-2.316 1.244-3.218 1.75c-.452.253-.867.496-1.22.875c-.35.377-.624.95-.624 1.53v1h8v-2H8.437c.736-.378 1.58-.756 2.688-1.438C12.14 22.928 13 21.845 13 20.5c0-1.924-1.576-3.5-3.5-3.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumericalSortingTwentyOne(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 5A3.512 3.512 0 0 0 5 8.5V9h2v-.5C7 7.624 7.624 7 8.5 7h1c.876 0 1.5.624 1.5 1.5c0 .456-.353.98-.938 1.344c-1.234.76-2.316 1.244-3.218 1.75c-.452.253-.867.496-1.22.875c-.35.377-.624.95-.624 1.53v1h8v-2H8.437c.736-.378 1.58-.756 2.688-1.438C12.14 10.928 13 9.845 13 8.5C13 6.576 11.424 5 9.5 5zM22 5v18.688l-2.594-2.594L18 22.5l4.28 4.313l.72.687l.72-.688L28 22.5l-1.406-1.406L24 23.687V5zM8.594 17l-.156.78s-.166.576-.563 1.157C7.478 19.52 6.98 20 6 20v2c1.376 0 2.32-.675 3-1.406V27h2V17z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OldTimeCamera(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.5 6l-.313.406L10 8H9V7H5v1H3v18h26V8h-7l-1.188-1.594L20.5 6zm1 2h7l1.188 1.594L21 10h6v4h-5.813c-1.042-1.784-2.98-3-5.187-3c-2.206 0-4.145 1.216-5.188 3H5v-4h6l.313-.406zM23 11v2h2v-2zm-7 2c2.22 0 4 1.78 4 4c0 2.22-1.78 4-4 4c-2.22 0-4-1.78-4-4c0-2.22 1.78-4 4-4M5 16h5.094A6.106 6.106 0 0 0 10 17c0 3.302 2.698 6 6 6s6-2.698 6-6c0-.337-.04-.678-.094-1H27v8H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OpenedFolder(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3v24.813l.78.156l12 2.5l1.22.25V28h6V15.437l1.72-1.718l.28-.314V3zm9.125 2H25v7.563l-1.72 1.718l-.28.314V26h-4v-8.906l-.28-.313L17 15.063V5.72l-.75-.19zM7 5.28l8 2v8.627l.28.313L17 17.937V28.28L7 26.188V5.282z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Organization(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 3v26h11v-4h2v4h11V3zm2 2h20v22h-7v-4h-6v4H6zm2 2v2h4V7zm6 0v2h4V7zm6 0v2h4V7zM8 11v2h4v-2zm6 0v2h4v-2zm6 0v2h4v-2zM8 15v2h4v-2zm6 0v2h4v-2zm6 0v2h4v-2zM8 19v2h4v-2zm6 0v2h4v-2zm6 0v2h4v-2zM8 23v2h4v-2zm12 0v2h4v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Outdent(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7v2h13V7zm0 4v2h20v-2zm22 0v10l5-5zM3 15v2h20v-2zm0 4v2h20v-2zm0 4v2h13v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paragraph(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5c-3.302 0-6 2.698-6 6s2.698 6 6 6h4v10h2V7h2v20h2V7h2V5zm0 2h4v8h-4c-2.22 0-4-1.78-4-4c0-2.22 1.78-4 4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParallelTasks(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5v8h3v2H5v4H2v8h8v-8H7v-2h8v2h-3v8h8v-8h-3v-2h8v2h-3v8h8v-8h-3v-4H17v-2h3V5zm2 2h4v4h-4zM4 21h4v4H4zm10 0h4v4h-4zm10 0h4v4h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paste(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 3c-1.26 0-2.152.89-2.594 2H5v23h8v2h14V14h-2V5h-7.406C17.152 3.89 16.26 3 15 3m0 2c.555 0 1 .445 1 1v1h3v2h-8V7h3V6c0-.555.445-1 1-1M7 7h2v4h12V7h2v7H13v12H7zm8 9h10v12H15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paypal(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.906 5l-.187.78l-3.69 17L4.75 24h4.875l-.594 2.78L8.75 28h6.97l.155-.813l1.03-4.812h2.69c3.762 0 7.065-2.444 7.905-6.375c.44-2.066-.026-3.816-1-5c-.79-.96-1.86-1.538-2.97-1.813c-.212-.986-.665-1.83-1.28-2.468C21.11 5.533 19.534 5 18.062 5H8.908zm1.625 2h7.534c.98 0 2.062.38 2.75 1.094c.687.715 1.09 1.734.718 3.47c-.66 3.09-2.995 4.81-5.936 4.81H11.28l-.155.782L10.095 22H7.25l3.28-15zm2.534 1.47l-.188.78l-1.03 4.625l-.283 1.22h3.532c1.777 0 3.313-1.294 3.75-3.033h.03c.007-.02-.004-.04 0-.062c.22-.876.08-1.773-.405-2.438c-.498-.68-1.333-1.093-2.22-1.093zm1.593 2h1.594c.314 0 .48.09.595.25c.116.158.2.418.093.81v.033c-.188.81-1.142 1.53-1.843 1.53h-1.03l.592-2.624zm8.938.874c.53.216 1.022.51 1.375.937c.58.71.903 1.685.56 3.284c-.66 3.092-2.995 4.812-5.936 4.812H15.28l-.155.78L14.095 26H11.25l.438-2h.03l.157-.813l1.03-4.812h2.69c3.762 0 7.065-2.444 7.905-6.375c.048-.226.067-.438.094-.656z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pdf(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3v26h20V3zm2 2h16v22H8zm7.406 5.344a1.429 1.429 0 0 0-.906.312c-.254.215-.37.482-.438.75c-.137.538-.097 1.1.032 1.72c.15.725.587 1.6.937 2.437c-.178.763-.224 1.435-.5 2.218c-.235.674-.535 1.06-.81 1.657c-.63.238-1.38.38-1.876.688c-.534.332-1.002.7-1.28 1.22c-.28.517-.25 1.25.123 1.78c.185.277.427.497.75.625c.324.128.675.133.97.03c.588-.202 1.006-.655 1.405-1.186c.37-.492.635-1.328.97-2c.504-.168.868-.38 1.405-.5c.564-.126.94-.065 1.47-.125c.227.258.42.67.655.874c.477.413 1 .743 1.625.78c.626.04 1.25-.35 1.594-.936h.032v-.032c.152-.267.257-.553.25-.875c-.006-.32-.17-.654-.375-.874c-.41-.44-.933-.55-1.5-.625c-.435-.056-1.045.1-1.562.126c-.452-.596-.903-1.046-1.313-1.812c-.22-.414-.282-.767-.468-1.188c.143-.68.43-1.436.468-2.03c.048-.72.02-1.34-.187-1.907a1.84 1.84 0 0 0-.53-.783a1.49 1.49 0 0 0-.907-.343h-.032zm.656 7.406c.18.316.402.517.594.813c-.282.05-.496-.002-.78.062c-.05.01-.078.05-.126.063c.057-.155.132-.25.188-.407c.064-.182.064-.346.124-.53zm3.688 2.03c.337.045.455.107.5.126c-.007.016.012.01 0 .032c-.123.205-.14.192-.22.187c-.066-.004-.32-.142-.56-.313c.07.006.216-.04.28-.03zm-7 1.564c-.055.08-.102.272-.156.343c-.305.406-.586.596-.656.625c-.013-.013.02 0 0-.03h-.032c-.103-.147-.072-.085 0-.22c.073-.134.31-.402.72-.656c.03-.018.093-.044.124-.062"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pencil(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.688 4.03c-.837 0-1.648.337-2.282.97l-.093.094l-.625-.594l-.688.688L5.406 20.78l-.218.22l-.063.313l-1.094 5.5l-.31 1.468l1.467-.31l5.5-1.095l.313-.063l.22-.218L26.81 11l.688-.688l-.594-.593l.063-.064l.03-.062A3.247 3.247 0 0 0 27 5a3.288 3.288 0 0 0-2.313-.97zm0 1.97c.31 0 .64.14.906.406c.533.534.533 1.248 0 1.782l-.094.093l-1.78-1.78l.093-.094c.266-.266.563-.406.875-.406zm-2.97 1.313l2.97 2.968l-1.438 1.47l-3-3l1.47-1.438zm-2.843 2.875l2.938 2.937l-10.438 10.47l-.406-1.814l-.126-.624l-.625-.125l-1.814-.405l10.47-10.438zM6.97 22.343l2.186.5l.5 2.187l-2.03.407l-1.064-1.062l.407-2.03z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.656 3c-.523 0-1.04.19-1.47.53l-.06.033l-.032.03l-3.125 3.22l.03.03c-.963.892-1.263 2.224-.844 3.376c.002.005-.002.024 0 .03c.847 2.428 3.017 7.11 7.25 11.344c4.25 4.25 8.997 6.33 11.344 7.25h.03a3.596 3.596 0 0 0 3.47-.688l.03-.03l.032-.032L28.407 25c.83-.83.83-2.264 0-3.094l-4.062-4.062l-.032-.063c-.83-.83-2.295-.83-3.125 0l-.03.032l-1.97 1.97a16.166 16.166 0 0 1-4.093-2.813c-1.636-1.563-2.474-3.36-2.78-4.064l1.998-2c.842-.842.855-2.236-.03-3.062l.03-.032l-.093-.093l-4-4.126l-.033-.03l-.062-.033A2.373 2.373 0 0 0 8.655 3zm0 2c.074 0 .147.036.22.094l4 4.093l.062.063l.03.03c-.007-.005.06.097-.062.22l-2.5 2.5l-.47.438l.22.624s1.15 3.073 3.563 5.376l.217.187C16.26 20.747 19 21.905 19 21.905l.625.282l.47-.468l2.5-2.5c.17-.17.14-.17.31 0L27 23.312c.17.17.17.11 0 .28l-3.063 3.063c-.46.396-.947.476-1.53.282c-2.268-.89-6.67-2.827-10.595-6.75C7.857 16.23 5.79 11.74 5.032 9.563c-.153-.407-.043-1.01.312-1.313l.03-.03l.032-.033l3.032-3.093A.356.356 0 0 1 8.655 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Picture(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 5v22h28V5zm2 2h24v13.906l-5.28-5.312l-.72-.72l-.72.72l-3.81 3.812l-5.75-5.812l-.72-.72l-.72.72L4 19.874zm20 2a2 2 0 1 0-.001 3.999A2 2 0 0 0 24 9m-13 6.72L20.188 25H4v-2.28zm11 2l6 6V25h-4.97l-4.155-4.188z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PieChart(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3m-1.125 2.063c.04-.005.084.003.125 0v11.343l.28.313l7.783 7.75C21.154 26.06 18.686 27 16 27C9.913 27 5 22.087 5 16c0-5.707 4.32-10.375 9.875-10.938zm2.125 0A10.956 10.956 0 0 1 26.938 15H17zM18.438 17h8.5c-.207 2.294-1.077 4.395-2.47 6.063L18.44 17z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinThree(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.47 2.563l-.72.75c-1.045 1.15-1.542 2.605-1.656 4.062l-4.063 4.063c-2.47-.44-5.127.253-7.03 2.156l-.72.72l.72.686l4.313 4.313L3 27.593V29h1.406l8.282-8.313l4.218 4.22l.688.718l.718-.72c2.096-2.094 2.69-4.973 2.063-7.56l3.563-3.564c1.737.208 3.513-.293 4.78-1.56l.688-.72l-.687-.72l-7.5-7.5l-.75-.717zm.218 3.062l5.687 5.688c-.825.466-1.716.758-2.625.53l-.563-.155l-.406.406l-4.186 4.187l-.438.44l.188.593c.587 1.762.315 3.75-.844 5.375L9.312 14.5c1.455-1 3.206-1.46 4.844-1.03l.53.155l.408-.406l4.625-4.626l.28-.28v-.408c0-.803.293-1.563.688-2.28z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3m0 2c6.087 0 11 4.913 11 11s-4.913 11-11 11S5 22.087 5 16S9.913 5 16 5m-1 5v5h-5v2h5v5h2v-5h5v-2h-5v-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Powerpoint(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3v26h20V3zm2 2h16v22H8zm5 6v2h4c1.19 0 2 .81 2 2s-.81 2-2 2s-2-.81-2-2h-2v7h2v-3.594c.594.35 1.26.594 2 .594c2.21 0 4-1.79 4-4s-1.79-4-4-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PriceTag(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 5l-.313.28L4.28 16.813l-.686.688l.687.72l9.5 9.5l.72.687l.69-.687l11.53-11.407L27 16V5zm.844 2H25v8.156L14.5 25.594L6.406 17.5zM22 9a1 1 0 1 0 0 2a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Puzzle(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 3c-2.21 0-4 1.79-4 4H7v5c-2.21 0-4 1.79-4 4s1.79 4 4 4v5h6c0 2.21 1.79 4 4 4s4-1.79 4-4h6v-7h-2c-1.19 0-2-.81-2-2s.81-2 2-2h2V7h-6c0-2.21-1.79-4-4-4m0 2c1.19 0 2 .81 2 2v2h6v3c-2.21 0-4 1.79-4 4s1.79 4 4 4v3h-6v2c0 1.19-.81 2-2 2s-2-.81-2-2v-2H9v-5H7c-1.19 0-2-.81-2-2s.81-2 2-2h2V9h6V7c0-1.19.81-2 2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QrCode(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 5v8h2v2h2v-2h4V5zm8 8v2h2v2h-4v2H5v8h8v-8h6v-2h-2v-2h4v-2h2v2h2v-2h2V5h-8v8zm12 2v2h2v-2zm0 2h-2v2h2zm0 2v2h2v-2zm0 2h-2v-2h-2v2h-5v6h2v-4h4v2h2v-2h1zm-3 4h-2v2h2zm1-8v-2h-2v2zm-12 0v-2H9v2zm-4-2H5v2h2zm8-10v4h-1v2h1v1h2V9h1V7h-1V5zM7 7h4v4H7zm14 0h4v4h-4zM8 8v2h2V8zm14 0v2h2V8zM7 21h4v4H7zm1 1v2h2v-2zm17 3v2h2v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionMark(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4c-4.33 0-8 3.056-8 7v3h6v-3c0-.147.09-.362.438-.594c.347-.23.914-.406 1.562-.406c.652 0 1.217.177 1.563.406c.345.23.437.43.437.594c0 .58-.19.98-.563 1.438c-.373.458-.955.928-1.593 1.468C14.567 14.986 13 16.496 13 19v1h6v-1c0-.34.124-.58.5-.97c.376-.388.996-.85 1.656-1.405C22.476 15.515 24 13.815 24 11c0-3.91-3.664-7-8-7m0 2c3.396 0 6 2.365 6 5c0 2.145-.976 3.1-2.156 4.094c-.59.496-1.22.98-1.782 1.562c-.347.36-.615.828-.812 1.344h-1.938c.315-1.01.922-1.783 1.844-2.563c.612-.517 1.28-1.056 1.844-1.75c.564-.693 1-1.606 1-2.687c0-.96-.572-1.76-1.313-2.25C17.948 8.26 17.006 8 16 8c-1.01 0-1.95.258-2.688.75C12.575 9.242 12 10.042 12 11v1h-2v-1c0-2.684 2.597-5 6-5m-3 16v6h6v-6zm2 2h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Recycling(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3.97a3.264 3.264 0 0 0-2.75 1.5l-2.813 4.624l1.72 1.03l2.812-4.593c.54-.805 1.563-.823 2.092 0l3.125 5l-1.5.876L23.313 15V9.687l-1.437.844l-3.125-5.06a3.264 3.264 0 0 0-2.75-1.5zm-5.594 8.124l-4.5 2.594l1.25.75l-2.562 4.218l-.032-.03c-.012.017.012.044 0 .06C3.092 21.916 4.788 25 7.5 25H13v-2H7.5c-1.267 0-1.946-1.25-1.28-2.22l.03-.03v-.03l2.625-4.25l1.53.936v-5.312zm14.656 3.562l-1.718 1.063l2.5 4c.64 1.087-.087 2.28-1.25 2.28H19v-2l-4.906 3L19 27v-2h5.594c2.637 0 4.328-3 2.968-5.313v-.03z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Refresh(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4c-5.11 0-9.383 3.16-11.125 7.625l1.844.75C8.176 8.64 11.71 6 16 6c3.24 0 6.134 1.59 7.938 4H20v2h7V5h-2v3.094A11.928 11.928 0 0 0 16 4m9.28 15.625C23.824 23.36 20.29 26 16 26c-3.276 0-6.157-1.612-7.97-4H12v-2H5v7h2v-3.094C9.19 26.386 12.395 28 16 28c5.11 0 9.383-3.16 11.125-7.625z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RemoveUser(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 4c-3.854 0-7 3.146-7 7c0 2.41 1.23 4.552 3.094 5.813C6.527 18.343 4 21.88 4 26h2c0-4.43 3.57-8 8-8c1.376 0 2.654.358 3.78.97A7.993 7.993 0 0 0 16 24c0 4.406 3.594 8 8 8c4.406 0 8-3.594 8-8c0-4.406-3.594-8-8-8a7.98 7.98 0 0 0-4.688 1.53c-.442-.277-.92-.51-1.406-.718A7.018 7.018 0 0 0 21 11c0-3.854-3.146-7-7-7m0 2c2.773 0 5 2.227 5 5s-2.227 5-5 5s-5-2.227-5-5s2.227-5 5-5m10 12c3.326 0 6 2.674 6 6s-2.674 6-6 6s-6-2.674-6-6s2.674-6 6-6m-2.28 2.28l-1.44 1.44L22.563 24l-2.28 2.28l1.437 1.44L24 25.437l2.28 2.28l1.44-1.437L25.437 24l2.28-2.28l-1.437-1.44L24 22.563l-2.28-2.28z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ResizeDiagonal(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 5v2h9.563L7 23.563V14H5v13h13v-2H8.437L25 8.437V18h2V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ResizeFourDirections(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 4v9h2V7.437L14.563 16L6 24.563V19H4v9h9v-2H7.437L16 17.437L24.563 26H19v2h9v-9h-2v5.563L17.437 16L26 7.437V13h2V4h-9v2h5.563L16 14.563L7.437 6H13V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ResizeHorizontal(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.28 6.78l-8.5 8.5l-.686.72l.687.72l8.5 8.5l1.44-1.44L3.936 17h24.125l-6.78 6.78l1.437 1.44l8.5-8.5l.686-.72l-.687-.72l-8.5-8.5l-1.44 1.44L28.063 15H3.938l6.782-6.78z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ResizeVertical(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 2.094l-.72.687l-8 8l1.44 1.44L15 5.936v20.125l-6.28-6.28l-1.44 1.437l8 8l.72.686l.72-.687l8-8l-1.44-1.44L17 26.063V5.938l6.28 6.282l1.44-1.44l-8-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightArrow(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.72 6.78l-1.44 1.44L24.063 15H4v2h20.063l-6.782 6.78l1.44 1.44l8.5-8.5l.686-.72l-.687-.72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightRound(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3m0 2c6.087 0 11 4.913 11 11s-4.913 11-11 11S5 22.087 5 16S9.913 5 16 5m.72 4.594L15.28 11l4 4H9v2h10.28l-4 4l1.44 1.406l5.686-5.687l.72-.72l-.72-.72l-5.687-5.686z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightSquared(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 5v22h22V5zm2 2h18v18H7zm7.22 2.28l-1.44 1.44L18.064 16l-5.282 5.28l1.44 1.44l6-6l.686-.72l-.687-.72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RotateLeft(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 3v2h1V3zm-2.156.438L12 3.655l-.03.03l-.033.002l-.843.312h-.03l-.033.03l-.092.032l.843 1.813h.033l.062-.03l.688-.25l.062-.032l.78-.22l-.56-1.905zM20 4v8h2V6.78c3.03 1.98 5 5.28 5 9.22c0 6.055-4.945 11-11 11S5 22.055 5 16v-1H3v1c0 7.145 5.855 13 13 13s13-5.855 13-13c0-4.06-1.795-7.615-4.656-10H28V4zM9.156 5l-.53.344l-.032.03h-.03l-.69.532l-.03.032h-.032l-.28.25l1.312 1.5l.22-.188l.06-.03l.563-.44L9.75 7l.47-.313L9.155 5zM6.094 7.594l-.22.28l-.03.032l-.032.032l-.53.687v.03l-.032.033l-.344.53l1.688 1.095l.28-.47l.064-.062l.437-.592l.03-.032l.22-.25zM3.97 11l-.033.125l-.03.03v.033L3.593 12v.03l-.03.064l-.22.844l1.906.53l.22-.78l.03-.063l.25-.688l.03-.062v-.03z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RotateRight(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3v2h1V3zm3.156.438l-.562 1.906l.78.218l.064.032l.687.25l.063.03h.03l.845-1.81l-.094-.033l-.032-.03h-.032l-.843-.313h-.032l-.03-.03zM4 4v2h3.656C4.796 8.385 3 11.94 3 16c0 7.145 5.855 13 13 13s13-5.855 13-13v-1h-2v1c0 6.055-4.945 11-11 11S5 22.055 5 16c0-3.94 1.97-7.24 5-9.22V12h2V4zm18.844 1L21.78 6.688l.47.312l.063.03l.562.44l.063.03l.218.188l1.313-1.5l-.282-.25h-.032l-.03-.032l-.69-.53h-.03l-.03-.032L22.843 5zm3.062 2.594l-1.53 1.312l.218.25l.03.03l.44.595l.06.064l.282.47l1.688-1.095l-.344-.533l-.03-.03v-.032l-.532-.688l-.032-.03l-.03-.032zM28.03 11l-1.81.844v.03l.03.063l.25.688l.03.063l.22.78l1.906-.53l-.218-.844l-.032-.063V12l-.312-.813v-.03l-.032-.032l-.03-.125z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rouble(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 6v14H7v2h2v4h2v-4h6v-2h-6v-3h7.5c3.026 0 5.5-2.474 5.5-5.5S21.526 6 18.5 6zm2 2h7.5c1.944 0 3.5 1.556 3.5 3.5S20.444 15 18.5 15H11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoundedRectangle(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 6c-2.745 0-5 2.255-5 5v10c0 2.745 2.255 5 5 5h16c2.745 0 5-2.255 5-5V11c0-2.745-2.255-5-5-5zm0 2h16c1.655 0 3 1.345 3 3v10c0 1.655-1.345 3-3 3H8c-1.655 0-3-1.345-3-3V11c0-1.655 1.345-3 3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoundedRectangleFilled(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24 6H8c-2.8 0-5 2.2-5 5v10c0 2.8 2.2 5 5 5h16c2.8 0 5-2.2 5-5V11c0-2.8-2.2-5-5-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rupee(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 5v2h4c1.704 0 3.94 1.038 4.72 3H8v2h8.97c-.31 2.61-2.996 4-4.97 4H8v2.47L18.25 27h3.125l-10.813-9H12c3.234 0 6.674-2.39 6.97-6H24v-2h-5.188C18.51 8.816 17.86 7.804 17 7h7V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.5 3C14.265 3 10 7.265 10 12.5c0 2.25.81 4.307 2.125 5.938L3.28 27.28l1.44 1.44l8.843-8.845C15.192 21.19 17.25 22 19.5 22c5.235 0 9.5-4.265 9.5-9.5S24.735 3 19.5 3m0 2c4.154 0 7.5 3.346 7.5 7.5S23.654 20 19.5 20S12 16.654 12 12.5S15.346 5 19.5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sensor(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3C8.826 3 3 8.826 3 16s5.826 13 13 13s13-5.826 13-13S23.174 3 16 3m0 1c6.633 0 12 5.367 12 12s-5.367 12-12 12S4 22.633 4 16S9.367 4 16 4m0 2C10.49 6 6 10.49 6 16s4.49 10 10 10s10-4.49 10-10S21.51 6 16 6m0 2c4.43 0 8 3.57 8 8s-3.57 8-8 8s-8-3.57-8-8s3.57-8 8-8m0 2c-3.302 0-6 2.698-6 6s2.698 6 6 6s6-2.698 6-6s-2.698-6-6-6m0 2c2.22 0 4 1.78 4 4c0 2.22-1.78 4-4 4c-2.22 0-4-1.78-4-4c0-2.22 1.78-4 4-4m0 2a2 2 0 1 0-.001 3.999A2 2 0 0 0 16 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Services(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 2.5v1.406a5.643 5.643 0 0 0-2.28.938l-1.032-.97l-1.375 1.47l1 .937a5.658 5.658 0 0 0-.907 2.22H15.5v2h1.406c.146.83.474 1.586.938 2.25l-1.063 1.03l1.44 1.44l1.03-1.064c.664.464 1.42.792 2.25.938V16.5h2v-1.406a5.657 5.657 0 0 0 2.22-.906l.936 1l1.47-1.376l-.97-1.03c.47-.67.79-1.445.938-2.282H29.5v-2h-1.406a5.609 5.609 0 0 0-.938-2.25l.938-.938l-1.407-1.406l-.937.938a5.609 5.609 0 0 0-2.25-.938V2.5zm1 3.313A3.664 3.664 0 0 1 26.188 9.5c0 2.055-1.633 3.688-3.688 3.688s-3.688-1.633-3.688-3.688s1.633-3.688 3.688-3.688zM9.53 11.718l-1.842.75l.718 1.81a6.937 6.937 0 0 0-2.344 2.314l-1.78-.72l-.75 1.845l1.78.718a6.796 6.796 0 0 0-.218 1.656c0 .57.085 1.126.218 1.656l-1.78.72l.75 1.843l1.78-.72a6.893 6.893 0 0 0 2.344 2.345l-.72 1.78l1.845.75l.72-1.78a6.76 6.76 0 0 0 1.656.218c.57 0 1.128-.085 1.656-.218l.72 1.78l1.843-.75l-.72-1.78a6.91 6.91 0 0 0 2.314-2.344l1.81.718l.75-1.843l-1.81-.72c.13-.53.218-1.087.218-1.656c0-.57-.087-1.128-.22-1.657l1.813-.718l-.75-1.845l-1.81.72a6.893 6.893 0 0 0-2.314-2.314l.72-1.81l-1.845-.75l-.717 1.81a6.86 6.86 0 0 0-1.657-.217c-.57 0-1.126.086-1.656.218l-.72-1.81zm2.376 3.592c2.663 0 4.78 2.12 4.78 4.782c.002 2.663-2.117 4.812-4.78 4.812a4.806 4.806 0 0 1-4.812-4.812c0-2.663 2.15-4.782 4.812-4.782"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Settings(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.313 4.53l-1.844.75l1.155 2.876a9.066 9.066 0 0 0-3.47 3.47L5.28 10.469l-.75 1.842l2.845 1.157A9.02 9.02 0 0 0 7 16a9 9 0 0 0 .375 2.53L4.53 19.688l.75 1.844l2.876-1.155a9.135 9.135 0 0 0 3.47 3.47l-1.157 2.874l1.842.75l1.157-2.845c.8.236 1.653.375 2.53.375a9 9 0 0 0 2.53-.375l1.157 2.844l1.844-.75l-1.155-2.876a9.066 9.066 0 0 0 3.47-3.47l2.874 1.157l.75-1.843l-2.845-1.156a9 9 0 0 0 .375-2.53a9 9 0 0 0-.375-2.53l2.844-1.158l-.75-1.843l-2.876 1.155a9.066 9.066 0 0 0-3.47-3.47l1.157-2.874l-1.843-.75l-1.156 2.845A8.912 8.912 0 0 0 16 7a9 9 0 0 0-2.53.375L12.31 4.53zM16 9c3.878 0 7 3.122 7 7s-3.122 7-7 7s-7-3.122-7-7s3.122-7 7-7m0 3c-2.197 0-4 1.803-4 4s1.803 4 4 4s4-1.803 4-4s-1.803-4-4-4m0 2c1.116 0 2 .884 2 2s-.884 2-2 2s-2-.884-2-2s.884-2 2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4c-2.144 0-3.884 1.72-3.97 3.844A9.933 9.933 0 0 0 6 17c0 .17-.008.36 0 .563c-1.185.696-2 1.972-2 3.437c0 2.197 1.803 4 4 4c.575 0 1.13-.122 1.625-.344C11.358 26.112 13.618 27 16 27s4.642-.888 6.375-2.344A3.98 3.98 0 0 0 24 25c2.197 0 4-1.803 4-4c0-1.48-.825-2.777-2.03-3.47c.015-.16.03-.32.03-.53a9.931 9.931 0 0 0-6.03-9.156C19.883 5.72 18.143 4 16 4m0 2c1.116 0 2 .884 2 2s-.884 2-2 2s-2-.884-2-2s.884-2 2-2m-3.53 3.844C13.14 11.118 14.47 12 16 12c1.53 0 2.86-.882 3.53-2.156A7.937 7.937 0 0 1 24 17c-2.197 0-4 1.803-4 4c0 .895.31 1.706.813 2.375A8.08 8.08 0 0 1 16 25a8.077 8.077 0 0 1-4.813-1.625A3.93 3.93 0 0 0 12 21c0-2.197-1.803-4-4-4c0-3.12 1.79-5.856 4.47-7.156M8 19c1.116 0 2 .884 2 2s-.884 2-2 2s-2-.884-2-2s.884-2 2-2m16 0c1.116 0 2 .884 2 2s-.884 2-2 2s-2-.884-2-2s.884-2 2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shekel(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 5v22h2V7h1c3.877 0 7 3.123 7 7v7h2v-7c0-4.96-4.04-9-9-9zm15 0v20h-1c-3.877 0-7-3.123-7-7v-7h-2v7c0 4.96 4.04 9 9 9h3V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCart(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 7a1 1 0 0 0 0 2h2.22l2.624 10.5c.223.89 1.02 1.5 1.937 1.5h11.47c.903 0 1.67-.6 1.907-1.47L27.75 10H11l.5 2h13.656l-1.906 7H11.78L9.157 8.5A1.984 1.984 0 0 0 7.22 7zm17 14c-1.645 0-3 1.355-3 3s1.355 3 3 3s3-1.355 3-3s-1.355-3-3-3m-9 0c-1.645 0-3 1.355-3 3s1.355 3 3 3s3-1.355 3-3s-1.355-3-3-3m0 2c.564 0 1 .436 1 1c0 .564-.436 1-1 1c-.564 0-1-.436-1-1c0-.564.436-1 1-1m9 0c.564 0 1 .436 1 1c0 .564-.436 1-1 1c-.564 0-1-.436-1-1c0-.564.436-1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shutdown(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 4v12h2V4zm-3 .688C7.346 6.338 4 10.788 4 16c0 6.617 5.383 12 12 12s12-5.383 12-12c0-5.213-3.346-9.662-8-11.313v2.157C23.526 8.39 26 11.91 26 16c0 5.514-4.486 10-10 10S6 21.514 6 16c0-4.09 2.474-7.61 6-9.156z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sort(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 3.594l-.72.687l-8 8L5.595 14h20.811l-1.687-1.72l-8-8l-.72-.686zm0 2.844L21.563 12H10.438zM5.594 18l1.687 1.72l8 8l.72.686l.72-.687l8-8L26.405 18zm4.843 2h11.126L16 25.563z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortDown(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.594 12l1.687 1.72l10 10l.72.686l.72-.687l10-10L28.405 12zm4.844 2h15.124L16 21.563L8.437 14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortLeft(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 4.594L18.28 6.28l-9 9l-.686.72l.687.72l9 9L20 27.405zm-2 4.843v13.126L11.437 16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortRight(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 4.594v22.812l1.72-1.687l9-9l.686-.72l-.687-.72l-9-9zm2 4.843L20.563 16L14 22.563V9.438z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortUp(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 8.594l-.72.687l-10 10L3.595 21h24.811l-1.687-1.72l-10-10l-.72-.686zm0 2.844L23.563 19H8.438L16 11.437z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spy(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.063 4c-.874 0-1.647.45-2.188 1.03c-.54.582-.935 1.31-1.28 2.095c-.534 1.21-.91 2.554-1.25 3.813c-1.087.316-2.01.71-2.75 1.187C4.727 12.683 4 13.457 4 14.5c0 .908.553 1.632 1.25 2.156c.593.446 1.324.816 2.188 1.125c.048.234.126.47.218.69c-.843.477-2.178 1.397-3.468 3.156l-.594.844l.844.593l3.28 2.25l-.624 1.25L6.374 28h19.25l-.72-1.438l-.623-1.25l3.283-2.25l.844-.593l-.593-.846c-1.29-1.76-2.626-2.68-3.47-3.156c.093-.222.17-.457.22-.69c.863-.308 1.594-.678 2.187-1.124c.698-.524 1.25-1.248 1.25-2.156c0-1.045-.724-1.817-1.593-2.375c-.742-.477-1.664-.87-2.75-1.188c-.373-1.303-.787-2.67-1.312-3.874c-.34-.775-.715-1.49-1.25-2.062c-.536-.57-1.296-1-2.157-1c-.583 0-1.024.162-1.5.28c-.475.12-.958.22-1.437.22c-.96 0-1.764-.5-2.937-.5zm0 2c.205 0 1.436.5 2.937.5c.75 0 1.42-.15 1.938-.28c.518-.132.913-.22 1-.22c.23 0 .403.072.687.375c.284.303.62.843.906 1.5c.545 1.243.96 2.94 1.408 4.5c0-.002.054-.048-.093.03c-.25.137-.772.313-1.407.408c-1.27.19-3 .187-4.437.187c-1.43 0-3.164-.02-4.437-.22c-.636-.097-1.15-.268-1.406-.405c-.078-.042-.107-.026-.125-.03a1 1 0 0 0 0-.033c.004-.01 0-.02 0-.03a1 1 0 0 0 .033-.032a1 1 0 0 0 .124-.438a1 1 0 0 0 0-.03c.36-1.328.76-2.735 1.25-3.845c.293-.667.61-1.212.906-1.53c.297-.32.5-.407.72-.407zm-4.876 7.094c.226.47.624.844 1.032 1.062c.605.325 1.306.477 2.06.594c1.51.234 3.274.25 4.72.25c1.437 0 3.206.007 4.72-.22c.754-.112 1.45-.26 2.06-.592c.41-.223.807-.618 1.03-1.094c.62.22 1.138.454 1.5.687c.58.376.69.655.69.72c0 .06-.053.25-.47.564c-.417.313-1.137.674-2.06.968c-1.853.59-4.52.97-7.47.97s-5.62-.38-7.47-.97c-.925-.293-1.645-.654-2.06-.968C6.05 14.75 6 14.56 6 14.5c0-.066.078-.315.656-.688c.363-.233.9-.49 1.532-.718zm2.594 5.47c.33.053.656.142 1 .186c.13.877.816 1.652 1.91 1.72c.84.05 1.79-.348 1.873-1.47c.15.002.288 0 .438 0c.15 0 .29.002.44 0c.082 1.122 1.032 1.52 1.874 1.47c1.092-.068 1.776-.843 1.906-1.72c.345-.044.67-.133 1-.188l-.095.625c-.308 1.643-1.044 3.17-1.97 4.22c-.923 1.05-2.01 1.61-3.155 1.593c-1.176-.017-2.24-.584-3.156-1.625c-.918-1.04-1.64-2.537-1.97-4.188l-.093-.625zM23 19.998c.372.218 1.35.86 2.47 2.094l-3.032 2.093l-.72.47l.376.78l.28.563H19.22a7.75 7.75 0 0 0 1.436-1.28c1.1-1.25 1.84-2.89 2.25-4.657c.034-.02.06-.043.094-.062zm-14.03.03c.038.024.083.042.124.064c.428 1.745 1.165 3.363 2.25 4.593c.45.51.973.95 1.53 1.313h-3.25l.282-.563l.375-.78l-.716-.47l-3.032-2.093c1.06-1.168 2.023-1.812 2.44-2.063z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Strikethrough(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 6v2h8v8h-5v2h5v8h2v-8h5v-2h-5V8h8V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stripe(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 6C3.355 6 2 7.355 2 9v14c0 1.645 1.355 3 3 3h22c1.645 0 3-1.355 3-3V9c0-1.645-1.355-3-3-3zm0 2h22c.565 0 1 .435 1 1v14c0 .565-.435 1-1 1H5c-.565 0-1-.435-1-1V9c0-.565.435-1 1-1m11.5 3.438a.834.834 0 0 0-.844.843c0 .47.37.846.844.846a.842.842 0 0 0 .844-.844a.837.837 0 0 0-.844-.842zm-5.438.75l-1.374.218l-.188 1.22l-.5.093l-.188 1.123H9.5v2.344c0 .61.15 1.022.47 1.28c.265.215.647.313 1.186.313c.416 0 .67-.078.844-.124v-1.25c-.097.026-.32.063-.47.063c-.317 0-.467-.163-.467-.533v-2.094h.843l.188-1.22h-1.03l-.002-1.437zM6.907 13.53c-.546 0-1.018.142-1.343.408c-.338.278-.5.7-.5 1.187c0 .883.53 1.244 1.406 1.563c.564.2.75.34.75.562c0 .214-.17.344-.5.344c-.41 0-1.09-.21-1.532-.47L5 18.344c.377.215 1.085.438 1.813.438c.578 0 1.05-.114 1.375-.374c.364-.286.562-.736.562-1.28c0-.904-.547-1.277-1.438-1.595c-.474-.174-.75-.285-.75-.5c0-.18.14-.31.407-.31c.485 0 .997.205 1.342.374l.187-1.22a3.692 3.692 0 0 0-1.595-.343zm13.813 0c-.47 0-.9.217-1.282.626l-.063-.53h-1.438v6.937l1.625-.282V18.69c.247.078.492.094.72.094c.402 0 .982-.1 1.437-.593c.433-.475.654-1.19.654-2.157c0-.856-.145-1.52-.47-1.936c-.285-.37-.68-.563-1.186-.563zm4.218 0c-1.37 0-2.22 1.034-2.22 2.658c0 .91.21 1.596.657 2.03c.403.39 1.004.563 1.75.563c.688 0 1.317-.137 1.72-.404l-.19-1.125a2.84 2.84 0 0 1-1.374.345c-.31 0-.53-.083-.685-.22c-.17-.142-.248-.355-.282-.686h2.657c.008-.08.03-.447.03-.564c0-.805-.186-1.458-.53-1.906c-.352-.456-.863-.69-1.532-.69zM15 13.564a.99.99 0 0 0-.97.687l-.092-.624H12.53v5.063h1.626v-3.283c.2-.247.485-.344.875-.344c.086 0 .17.013.284.032v-1.5a1.35 1.35 0 0 0-.313-.03zm.688.062v5.063h1.624v-5.064zm9.218 1c.364 0 .568.348.594 1.063h-1.188c.046-.723.224-1.064.594-1.064zm-4.75.156c.41 0 .625.42.625 1.283c.003.492-.07.896-.217 1.156c-.13.247-.328.375-.562.375a.98.98 0 0 1-.437-.094v-2.405c.273-.286.504-.313.594-.313z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Student(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 7.78l-.313.095l-12.5 4.188L.345 13L2 13.53v8.75c-.597.347-1 .98-1 1.72a2 2 0 1 0 4 0c0-.74-.403-1.373-1-1.72v-8.06l2 .655V20c0 .82.5 1.5 1.094 1.97c.594.467 1.332.797 2.218 1.093c1.774.59 4.112.937 6.688.937c2.576 0 4.914-.346 6.688-.938c.886-.295 1.624-.625 2.218-1.093C25.5 21.5 26 20.82 26 20v-5.125l2.813-.938L31.655 13l-2.843-.938l-12.5-4.187zm0 2.095L25.375 13L16 16.125L6.625 13zm-8 5.688l7.688 2.562l.312.094l.313-.095L24 15.562V20c0 .01.004.126-.313.375c-.316.25-.883.565-1.625.813C20.58 21.681 18.395 22 16 22c-2.395 0-4.58-.318-6.063-.813c-.74-.247-1.308-.563-1.624-.812C7.995 20.125 8 20.01 8 20z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Subscript(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4.156 8l1 1.53L9.313 16l-4.157 6.47l-1 1.53h6.374l.314-.47l1.656-2.56l1.656 2.56l.313.47h6.373l-1-1.53L15.687 16l4.157-6.47l1-1.53H14.47l-.314.47l-1.656 2.56l-1.656-2.56L10.53 8H4.157zm3.657 2H9.47l2.185 3.438l.844 1.312l.843-1.313L15.53 10h1.658l-3.532 5.47l-.344.53l.344.53l3.53 5.47H15.53l-2.185-3.438l-.844-1.312l-.843 1.313L9.47 22H7.81l3.532-5.47l.344-.53l-.344-.53L7.814 10zm17.156 9A2.985 2.985 0 0 0 22 21.97V22h2v-.03c0-.547.423-.97.97-.97h.06c.547 0 .97.423.97.97c0 .317-.143.6-.406.78l-2.125 1.47A3.369 3.369 0 0 0 22 27v1h6v-2h-3.47c.04-.032.023-.097.064-.125l2.125-1.47c.8-.552 1.28-1.46 1.28-2.436A2.985 2.985 0 0 0 25.03 19h-.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Superscript(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.97 3A2.987 2.987 0 0 0 22 5.97V6h2v-.03c0-.547.423-.97.97-.97h.06c.547 0 .97.423.97.97c0 .317-.143.6-.406.78L23.47 8.22A3.372 3.372 0 0 0 22 11v1h6v-2h-3.47c.04-.032.023-.097.064-.125l2.125-1.47c.8-.55 1.28-1.46 1.28-2.435A2.985 2.985 0 0 0 25.03 3zM4.155 8l1 1.53L9.313 16l-4.157 6.47l-1 1.53h6.374l.314-.47l1.656-2.56l1.656 2.56l.313.47h6.373l-1-1.53L15.687 16l4.157-6.47l1-1.53H14.47l-.314.47l-1.656 2.56l-1.656-2.56L10.53 8H4.157zm3.657 2h1.656l2.186 3.438l.844 1.312l.844-1.313L15.53 10h1.658l-3.532 5.47l-.344.53l.344.53l3.53 5.47H15.53l-2.185-3.438l-.844-1.312l-.843 1.313L9.47 22H7.81l3.532-5.47l.344-.53l-.344-.53L7.814 10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Support(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 3c-4.43 0-8 3.57-8 8c0 .235.037.553.063.844L4.5 20.406a5 5 0 0 0 0 7.094a5 5 0 0 0 7.094 0l8.562-8.563c.29.026.61.063.844.063c4.43 0 8-3.57 8-8a7.597 7.597 0 0 0-.938-3.688l-.625-1.156l-.937.938l-4.313 4.28l-1.562-1.562l4.28-4.312l.94-.938l-1.157-.625A7.618 7.618 0 0 0 21 3m0 2c.486 0 .848.21 1.28.313l-3.78 3.78l-.72.72l.72.687l3 3l.688.72l.718-.72l3.782-3.78c.103.432.312.794.312 1.28c0 3.37-2.63 6-6 6c-.4 0-.72 0-.97-.063l-.53-.156l-.406.407l-8.907 8.907c-1.23 1.23-3.05 1.23-4.28 0l-.032-.032c-1.2-1.23-1.19-3.03.03-4.25l8.907-8.906l.407-.406l-.158-.53c-.06-.25-.062-.57-.062-.97c0-3.37 2.63-6 6-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tags(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.594 4l-.313.28l-11 11l-.685.72l.687.72l9 9l.72.686l.72-.687l11-11l.28-.315V4zm.844 2H23v7.563l-10 10L5.437 16zM26 7v2h1v8.156l-9.5 9.438l-1.25-1.25l-1.406 1.406l1.937 1.97l.72.686l.69-.687L28.72 18.31L29 18V7zm-6 1a1 1 0 1 0 0 2a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tasks(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 5v6h22V5zm6 2h14v2H11zm-6 6v6h22v-6zm16 2h4v2h-4zM5 21v6h22v-6zm11 2h9v2h-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextHeight(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 6v2h7v18h2V8h7V6zm21 .625L21.5 11H24v10h-2.5l3.5 4.375L28.5 21H26V11h2.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextWidth(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 6v2h7v14h2V8h7V6zm2 15.5L5.625 25L10 28.5V26h12v2.5l4.375-3.5L22 21.5V24H10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ticket(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 7v7h1c1.19 0 2 .81 2 2s-.81 2-2 2H2v7h28v-7h-1c-1.19 0-2-.81-2-2s.81-2 2-2h1V7zm2 2h24v3.188c-1.715.45-3 1.955-3 3.812c0 1.857 1.285 3.362 3 3.813V23H4v-3.188c1.715-.45 3-1.955 3-3.812c0-1.857-1.285-3.362-3-3.813z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Timeline(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 4v24h2v-5h2.188c.418 1.156 1.52 2 2.812 2c1.292 0 2.394-.844 2.813-2H28v-2H13.812c-.418-1.156-1.52-2-2.812-2c-1.292 0-2.394.844-2.813 2H6v-4h14.188c.418 1.156 1.52 2 2.812 2c1.292 0 2.394-.844 2.813-2H28v-2h-2.188c-.418-1.156-1.52-2-2.812-2c-1.292 0-2.394.844-2.813 2H6v-4h7.188c.418 1.156 1.52 2 2.812 2c1.292 0 2.394-.844 2.813-2H28V9h-9.188c-.418-1.156-1.52-2-2.812-2c-1.292 0-2.394.844-2.813 2H6V4zm12 5c.564 0 1 .436 1 1c0 .564-.436 1-1 1c-.564 0-1-.436-1-1c0-.564.436-1 1-1m7 6c.564 0 1 .436 1 1c0 .564-.436 1-1 1c-.564 0-1-.436-1-1c0-.564.436-1 1-1m-12 6c.564 0 1 .436 1 1c0 .564-.436 1-1 1c-.564 0-1-.436-1-1c0-.564.436-1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TodoList(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.28 5.28L7 8.563l-1.28-1.28L4.28 8.72l2 2l.72.686l.72-.687l4-4zM15 7v2h13V7zm-4.72 6.28L7 16.564L5.72 15.28l-1.44 1.44l2 2l.72.686l.72-.687l4-4l-1.44-1.44zM15 15v2h13v-2zm-4.72 6.28L7 24.563l-1.28-1.28l-1.44 1.437l2 2l.72.686l.72-.687l4-4l-1.44-1.44zM15 23v2h13v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Translation(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 4v18h6v6h18V10h-6V4zm2 2h14v4.563L10.562 20H6zm5 2v1H8v2h4.938c-.13 1.15-.482 2.054-1.063 2.688a4.567 4.567 0 0 1-.906-.407c-.704-.418-.97-.86-.97-1.28H8c0 1.192.734 2.182 1.72 2.844A8.487 8.487 0 0 1 8 15v2c1.772 0 3.248-.405 4.375-1.156c.524.09 1.053.156 1.625.156v-1.875c.543-.91.833-1.973.938-3.125H16V9h-3V8zm10.438 4H26v14H12v-4.563zM20 13.844l-.938 2.844l-2 6l-.062.156V24h2v-.875l.03-.125h1.94l.03.125V24h2v-1.156l-.063-.157l-2-6L20 13.845zm0 6.28l.28.876h-.56l.28-.875z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 4c-.522 0-1.06.185-1.438.563C13.185 4.94 13 5.478 13 6v1H7v2h1v16c0 1.645 1.355 3 3 3h12c1.645 0 3-1.355 3-3V9h1V7h-6V6c0-.522-.185-1.06-.563-1.438C20.06 4.186 19.523 4 19 4zm0 2h4v1h-4zm-5 3h14v16c0 .555-.445 1-1 1H11c-.555 0-1-.445-1-1zm2 3v11h2V12zm4 0v11h2V12zm4 0v11h2V12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trophy(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 4c-.943 0-1.916.355-2.688 1.03C3.542 5.708 3 6.75 3 8c0 1.062.446 1.97 1 2.688c.554.716 1.226 1.285 1.844 1.843C7.08 13.648 8 14.565 8 16c0 .152-.15.52-.344.813c-.194.29-.375.468-.375.468l1.44 1.407s.318-.322.624-.78c.26-.394.545-.905.625-1.563c.113.322.22.657.342.968c.69 1.775 1.48 3.365 2.376 4.563c.325.436.66.824 1.03 1.156C11.653 23.18 10 24.9 10 27v1h12v-1c0-2.123-1.685-3.853-3.78-3.97a7.71 7.71 0 0 0 1.06-1.155c.9-1.194 1.682-2.79 2.376-4.563c.13-.332.253-.687.375-1.03c.07.687.357 1.22.626 1.624c.306.46.625.782.625.782l1.44-1.407s-.182-.176-.376-.468c-.195-.29-.344-.66-.344-.812c0-1.436.92-2.353 2.156-3.47c.618-.557 1.29-1.126 1.844-1.842c.554-.717 1-1.626 1-2.688c0-1.252-.54-2.293-1.313-2.97A4.122 4.122 0 0 0 25 4a3.927 3.927 0 0 0-2.625 1H9.625A3.927 3.927 0 0 0 7 4m0 2c.474 0 .778.15 1 .28c.028 2.04.38 4.354.938 6.626c-.532-.72-1.198-1.316-1.782-1.844c-.632-.57-1.21-1.096-1.593-1.593C5.178 8.97 5 8.542 5 8c0-.698.272-1.132.656-1.47C6.04 6.196 6.566 6 7 6m18 0c.434 0 .96.195 1.344.53c.384.338.656.772.656 1.47c0 .543-.18.972-.563 1.47c-.383.496-.96 1.02-1.593 1.592c-.6.54-1.274 1.163-1.813 1.907c.576-2.29.94-4.62.97-6.69c.222-.13.526-.28 1-.28zM10.125 7h11.75c-.153 2.972-.92 6.598-2.094 9.594c-.647 1.657-1.397 3.095-2.124 4.062C16.93 21.623 16.264 22 15.97 22c-.295 0-.933-.374-1.658-1.344c-.724-.97-1.48-2.403-2.124-4.062c-1.166-3-1.91-6.637-2.063-9.594M15 10v5h2v-5zm-1 15h4c.717 0 1.21.443 1.563 1h-7.125c.352-.557.845-1 1.562-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TurkishLira(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 4v5.906L8 11v2l3-1.094v2L8 15v2l3-1.094V28h1c5.76 0 10.828-3.85 12.344-9.406l.625-2.344l-1.94-.5l-.625 2.313c-1.19 4.36-4.977 7.36-9.406 7.78V15.19L19 13v-2l-6 2.188v-2L19 9V7l-6 2.188V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Umbrella(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 4v2.063C8.988 6.565 4.255 11.47 4.03 16.625c-.004.127-.03.248-.03.375h2c0-.08.088-.278.53-.53c.444-.254 1.158-.47 1.97-.47s1.526.216 1.97.47c.442.252.53.45.53.53h2c0-.012.083-.26.625-.53c.542-.272 1.406-.47 2.375-.47c.97 0 1.833.198 2.375.47c.542.27.625.518.625.53h2c0-.08.088-.278.53-.53c.444-.254 1.158-.47 1.97-.47s1.526.216 1.97.47c.442.252.53.45.53.53h2c0-.128-.01-.254-.03-.375C27.744 11.47 23.01 6.565 17 6.062V4zm1 4c4.346 0 8.06 2.954 9.438 6.28A6.582 6.582 0 0 0 23.5 14c-1.12 0-2.145.248-2.97.72a3.773 3.773 0 0 0-.56.405a4.489 4.489 0 0 0-.69-.438C18.376 14.237 17.24 14 16 14c-1.24 0-2.375.235-3.28.688c-.234.116-.45.28-.658.437c-.006-.006-.024.006-.03 0a3.858 3.858 0 0 0-.563-.406c-.826-.473-1.85-.72-2.97-.72c-.694 0-1.344.096-1.94.28C7.94 10.954 11.655 8 16 8m-1 9v8c0 .565-.435 1-1 1s-1-.435-1-1v-1h-2v1c0 1.645 1.355 3 3 3s3-1.355 3-3v-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Underline(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 4v12c0 4.43 3.57 8 8 8s8-3.57 8-8V4h-2v12c0 3.37-2.63 6-6 6s-6-2.63-6-6V4zM6 26v2h20v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Undo(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.78 5.28l-8 8l-.686.72l.687.72l8 8l1.44-1.44L7.936 15H21c2.755 0 5 2.245 5 5v7h2v-7c0-3.845-3.155-7-7-7H7.937l6.282-6.28z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnlockTwo(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3c-3.036 0-5.584 1.966-6.625 4.625l1.844.75C11.977 6.435 13.835 5 16 5c2.755 0 5 2.245 5 5v1.375C19.525 10.515 17.826 10 16 10c-5.51 0-10 4.49-10 10s4.49 10 10 10s10-4.49 10-10a9.968 9.968 0 0 0-3-7.125V10c0-3.845-3.155-7-7-7m0 9c4.43 0 8 3.57 8 8s-3.57 8-8 8s-8-3.57-8-8s3.57-8 8-8m0 6a2 2 0 0 0-2 2c0 .74.403 1.373 1 1.72V25h2v-3.28c.597-.347 1-.98 1-1.72a2 2 0 0 0-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpArrow(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 4.094l-.72.687l-8.5 8.5l1.44 1.44L15 7.936V28h2V7.937l6.78 6.782l1.44-1.44l-8.5-8.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpRound(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3m0 2c6.087 0 11 4.913 11 11s-4.913 11-11 11S5 22.087 5 16S9.913 5 16 5m0 3.875l-.72.72l-5.686 5.686L11 16.72l4-4V23h2V12.72l4 4l1.406-1.44l-5.687-5.686l-.72-.72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpSquared(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 5v22h22V5zm2 2h18v18H7zm9 4.094l-.72.687l-6 6l1.44 1.44L16 13.937l5.28 5.28l1.44-1.437l-6-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Upload(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 3.594l-.72.687l-7 7l1.44 1.44L15 7.436V24h2V7.437l5.28 5.282l1.44-1.44l-7-7zM7 26v2h18v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadTwo(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4c-4.12 0-7.36 3.13-7.813 7.125a4.895 4.895 0 0 0-3.843 3.22C1.884 15.054 0 17.248 0 20c0 3.324 2.676 6 6 6h20c3.324 0 6-2.676 6-6c0-1.76-.855-3.336-2.094-4.438c-.232-3.514-3.035-6.318-6.562-6.5C22.14 6.133 19.378 4 16 4m0 2c2.762 0 4.97 1.77 5.75 4.28l.22.72H23c2.755 0 5 2.245 5 5v.5l.406.313A4.07 4.07 0 0 1 30 20c0 2.276-1.724 4-4 4H6c-2.276 0-4-1.724-4-4c0-2.02 1.45-3.588 3.28-3.906l.657-.125l.125-.658C6.362 13.964 7.556 13 9 13h1v-1c0-3.37 2.63-6 6-6m0 5.594l-.72.687l-4 4l1.44 1.44L15 15.437V22h2v-6.563l2.28 2.282l1.44-1.44l-4-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UsDollar(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 4v2.03c-2.77.203-5 2.526-5 5.345c0 2.584 1.864 4.788 4.406 5.25l.594.125V24h-.625A3.368 3.368 0 0 1 11 20.625V19H9v1.625C9 23.575 11.426 26 14.375 26H15v2h2v-2h.625C20.573 26 23 23.574 23 20.625c0-2.584-1.863-4.788-4.406-5.25L17 15.062v-7.03a3.36 3.36 0 0 1 3 3.343V13h2v-1.625c0-2.82-2.23-5.142-5-5.344V4zm0 4.03v6.658l-.25-.032a3.314 3.314 0 0 1-2.75-3.28a3.361 3.361 0 0 1 3-3.345zm2 9.095l1.25.22a3.313 3.313 0 0 1 2.75 3.28A3.369 3.369 0 0 1 17.625 24H17z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserFemale(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.125 4c-3.304 0-6.984.562-9.72 3.594C5.673 10.626 4 15.88 4 25v1h8.656c.99.625 2.103 1 3.344 1c1.24 0 2.355-.383 3.344-1H29v-1c0-8.125-1.57-12.844-3.625-15.594c-1.81-2.42-3.892-3.094-5.438-3.25L19 4.5l-.28-.5zm-.563 2.063l.813 1.437l.28.5h.595c1.01 0 2.848.34 4.53 2.594C25.386 12.74 26.8 16.83 26.938 24h-5.375c.11-.14.21-.292.313-.438C23.233 21.625 24 19.207 24 17h-2c0 1.722-.644 3.827-1.75 5.406C19.144 23.986 17.665 25 16 25c-1.663 0-3.143-1.01-4.25-2.594C10.643 20.824 10 18.71 10 17c0-.444.085-.667.22-.844c.132-.177.364-.33.717-.468c.707-.28 1.9-.395 3.157-.5c1.258-.106 2.57-.206 3.75-.75C19.024 13.893 20 12.66 20 11h-2c0 1.044-.274 1.304-.97 1.625c-.694.32-1.882.458-3.124.563c-1.242.104-2.55.163-3.72.624c-.583.23-1.148.578-1.56 1.126C8.21 15.485 8 16.218 8 17c0 2.198.768 4.59 2.125 6.53c.11.16.225.318.344.47H6.06c.135-8.163 1.71-12.696 3.844-15.063c2.088-2.314 4.783-2.815 7.656-2.874zM13 17a1 1 0 1 0 0 2a1 1 0 0 0 0-2m6 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserMale(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 4c-3.666 0-6.446.862-8.313 2.625C7.822 8.388 7 10.958 7 14v.906C6.428 15.45 6 16.15 6 17c0 1.26.89 2.154 2 2.594c.37 1.167.773 2.393 1.22 3.437c.485 1.142.924 2.048 1.53 2.69a7.192 7.192 0 0 0 10.5 0c.606-.642 1.013-1.548 1.5-2.69c.446-1.044.88-2.27 1.25-3.436c1.11-.44 2-1.334 2-2.594c0-.846-.43-1.547-1-2.094V14c0-2.824-.643-4.834-1.78-6.156c-.966-1.12-2.255-1.58-3.532-1.72l-.782-1.56l-.28-.564zm-.594 2.063l.688 1.375c.333.666.463 1.14.468 1.343c.006.204.013.12-.03.158c-.087.073-.93.29-1.97.343c-1.04.054-2.264.028-3.406.47c-1.14.442-2.143 1.627-2.156 3.25h2c.008-.986.237-1.128.875-1.375c.638-.247 1.738-.32 2.813-.375c1.074-.055 2.183.01 3.125-.78c.47-.397.77-1.08.75-1.75c-.005-.148-.04-.29-.063-.44a3.23 3.23 0 0 1 1.188.876C22.413 10 23 11.482 23 14v1.844l.5.28c.304.177.5.496.5.876a.98.98 0 0 1-.906 1l-.688.03l-.187.657a37.548 37.548 0 0 1-1.282 3.532c-.45 1.056-.967 1.956-1.125 2.124c-2.13 2.25-5.497 2.25-7.625 0c-.16-.168-.675-1.068-1.126-2.125a37.86 37.86 0 0 1-1.28-3.532l-.188-.657l-.688-.03A.98.98 0 0 1 8 17c0-.374.193-.698.5-.875l.5-.28V14c0-2.697.684-4.636 2.063-5.938c1.28-1.21 3.37-1.9 6.343-2zM13 16a1 1 0 1 0 0 2a1 1 0 0 0 0-2m6 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoCall(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 8v16h22v-3.375l4.563 2.28l1.437.72V8.375l-1.438.72L24 11.374V8zm2 2h18v12H4zm24 1.625v8.75l-4-2v-4.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoFile(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3v26h20V9.594l-.28-.313l-6-6l-.314-.28H6zm2 2h10v6h6v16H8zm12 1.438L22.563 9H20zm-7 6.78v9.562l1.5-.936l5-3L20.938 18l-1.438-.844l-5-3zm2 3.532L17.094 18L15 19.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Visa(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 6C3.355 6 2 7.355 2 9v14c0 1.645 1.355 3 3 3h22c1.645 0 3-1.355 3-3V9c0-1.645-1.355-3-3-3zm0 2h22c.565 0 1 .435 1 1v14c0 .565-.435 1-1 1H5c-.565 0-1-.435-1-1V9c0-.565.435-1 1-1m14.406 4.375c-1.926 0-2.937.95-2.937 2.156c0 2.183 2.56 1.88 2.56 3c0 .193-.16.627-1.22.627s-1.75-.375-1.75-.375l-.31 1.44s.655.405 1.97.405c1.31 0 3.154-1.008 3.154-2.47c0-1.756-2.563-1.872-2.563-2.655c0-.4.34-.72 1.25-.72c.594 0 1.25.47 1.25.47l.314-1.53s-.87-.345-1.72-.345zm-8.218.125L9.5 17.438s-.084-.428-.125-.688c-.955-2.14-2.5-2.938-2.5-2.938l1.47 5.594h2l2.81-6.906H11.19zm2.687 0l-1.094 6.906h1.846L15.72 12.5zm9.938 0l-3 6.906h1.812l.375-.968h2.313l.187.968h1.656L25.75 12.5zm-18.97.03s3.597 1.105 4.438 3.814l-.623-3.125s-.275-.69-1-.69H4.844zM24.5 14.5l.53 2.594h-1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Won(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4 6l1.813 6H4v2h2.406L7 16H4v2h3.594L10 26h2l2-8h4l2 8h2l2.406-8H28v-2h-3l.594-2H28v-2h-1.813L28 6h-2l-1.813 6H18.5L17 6h-2l-1.5 6H7.812L6 6zm12 4l.5 2h-1zm-7.594 4H13l-.5 2H9zM15 14h2l.5 2h-3zm4 0h4.594L23 16h-3.5zm-9.406 4H12l-1.094 4.375zM20 18h2.406l-1.312 4.375z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Word(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3v26h20V3zm2 2h16v22H8zm10 7v6.5c0 .217-.283.5-.5.5c-.042 0 .02.048-.063-.063c-.082-.11-.206-.388-.28-.687C17.006 17.652 17 17 17 17v-2h-2v4.5c0 .217-.283.5-.5.5c-.217 0-.5-.283-.5-.5V13h-4v2h2v4.5c0 1.383 1.117 2.5 2.5 2.5c.984 0 1.686-.644 2.094-1.47c.302.187.52.47.906.47c1.383 0 2.5-1.117 2.5-2.5V14h2v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xls(children ...ElementRenderer) *Icons8Icon {
	return &Icons8Icon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3v7h2V5h16v5h2V3zm2 9l2 4l-2 4h2l1-2l1 2h2l-2-4l2-4h-2l-1 2l-1-2zm7 0v8h4v-2h-2v-6zm7.5 0c-1.383 0-2.5 1.117-2.5 2.5s1.117 2.5 2.5 2.5c.217 0 .5.283.5.5c0 .3-.12.5-.5.5c-.368 0-.424-.08-.438-.094c-.013-.013-.062-.08-.062-.312h-2c0 .566.163 1.2.625 1.687c.462.488 1.143.72 1.875.72c1.42 0 2.5-1.2 2.5-2.5c0-1.383-1.117-2.5-2.5-2.5c-.217 0-.5-.283-.5-.5c0-.217.283-.5.5-.5c.267 0 .348.063.406.125c.06.062.094.17.094.28h2c0-.587-.215-1.192-.656-1.655c-.442-.463-1.11-.75-1.844-.75zM6 22v7h20v-7h-2v5H8v-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
