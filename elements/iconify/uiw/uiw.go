package uiw

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "2.6.10"
	hAttr          = "1em"
	viewbox        = "0 0 20 20"
	fill           = "currentColor"
)

type UiwIcon struct {
	*SVGSVGElement
}

type UiwIconFn func(children ...ElementRenderer) *UiwIcon

var IconLookup = map[string]UiwIconFn{
	"adobe":           Adobe,
	"alipay":          Alipay,
	"aliwangwang":     Aliwangwang,
	"android":         Android,
	"androidO":        AndroidO,
	"apple":           Apple,
	"appstore":        Appstore,
	"appstoreO":       AppstoreO,
	"areaChart":       AreaChart,
	"arrowDown":       ArrowDown,
	"arrowLeft":       ArrowLeft,
	"arrowRight":      ArrowRight,
	"arrowUp":         ArrowUp,
	"arrowsAlt":       ArrowsAlt,
	"asterisk":        Asterisk,
	"backward":        Backward,
	"baidu":           Baidu,
	"barChart":        BarChart,
	"barcode":         Barcode,
	"bell":            Bell,
	"cameraO":         CameraO,
	"caretDown":       CaretDown,
	"caretLeft":       CaretLeft,
	"caretRight":      CaretRight,
	"caretUp":         CaretUp,
	"check":           Check,
	"checkSquare":     CheckSquare,
	"checkSquareO":    CheckSquareO,
	"chrome":          Chrome,
	"circleCheck":     CircleCheck,
	"circleCheckO":    CircleCheckO,
	"circleClose":     CircleClose,
	"circleCloseO":    CircleCloseO,
	"circleO":         CircleO,
	"close":           Close,
	"closeSquare":     CloseSquare,
	"closeSquareO":    CloseSquareO,
	"cloudDownload":   CloudDownload,
	"cloudDownloadO":  CloudDownloadO,
	"cloudUpload":     CloudUpload,
	"cloudUploadO":    CloudUploadO,
	"coffee":          Coffee,
	"component":       Component,
	"copy":            Copy,
	"copyright":       Copyright,
	"cssThree":        CssThree,
	"cut":             Cut,
	"darrowLeft":      DarrowLeft,
	"darrowRight":     DarrowRight,
	"dcaret":          Dcaret,
	"dashboard":       Dashboard,
	"date":            Date,
	"delete":          Delete,
	"dingding":        Dingding,
	"dislikeO":        DislikeO,
	"document":        Document,
	"dotChart":        DotChart,
	"down":            Down,
	"downCircle":      DownCircle,
	"downCircleO":     DownCircleO,
	"downSquare":      DownSquare,
	"downSquareO":     DownSquareO,
	"download":        Download,
	"edit":            Edit,
	"enter":           Enter,
	"environment":     Environment,
	"environmentO":    EnvironmentO,
	"eye":             Eye,
	"eyeO":            EyeO,
	"facebook":        Facebook,
	"fileAdd":         FileAdd,
	"fileExcel":       FileExcel,
	"fileJpg":         FileJpg,
	"filePdf":         FilePdf,
	"fileText":        FileText,
	"fileUnknown":     FileUnknown,
	"filter":          Filter,
	"firefox":         Firefox,
	"folder":          Folder,
	"folderAdd":       FolderAdd,
	"folderOpen":      FolderOpen,
	"forward":         Forward,
	"foursquare":      Foursquare,
	"frown":           Frown,
	"frownO":          FrownO,
	"github":          Github,
	"githubO":         GithubO,
	"global":          Global,
	"heartOff":        HeartOff,
	"heartOn":         HeartOn,
	"home":            Home,
	"htmlFive":        HtmlFive,
	"ie":              Ie,
	"inbox":           Inbox,
	"information":     Information,
	"informationO":    InformationO,
	"laptop":          Laptop,
	"left":            Left,
	"leftCircle":      LeftCircle,
	"leftCircleO":     LeftCircleO,
	"leftSquare":      LeftSquare,
	"leftSquareO":     LeftSquareO,
	"likeO":           LikeO,
	"link":            Link,
	"linkedin":        Linkedin,
	"linux":           Linux,
	"loading":         Loading,
	"lock":            Lock,
	"login":           Login,
	"logout":          Logout,
	"mail":            Mail,
	"mailO":           MailO,
	"man":             Man,
	"map":             Map,
	"meh":             Meh,
	"mehO":            MehO,
	"menu":            Menu,
	"menuFold":        MenuFold,
	"menuUnfold":      MenuUnfold,
	"message":         Message,
	"minus":           Minus,
	"minusCircle":     MinusCircle,
	"minusCircleO":    MinusCircleO,
	"minusSquare":     MinusSquare,
	"minusSquareO":    MinusSquareO,
	"mobile":          Mobile,
	"more":            More,
	"notification":    Notification,
	"opera":           Opera,
	"paperClip":       PaperClip,
	"pause":           Pause,
	"pauseCircle":     PauseCircle,
	"pauseCircleO":    PauseCircleO,
	"pay":             Pay,
	"payCircleO":      PayCircleO,
	"picasa":          Picasa,
	"picture":         Picture,
	"pieChart":        PieChart,
	"pinterest":       Pinterest,
	"playCircle":      PlayCircle,
	"playCircleO":     PlayCircleO,
	"plus":            Plus,
	"plusCircle":      PlusCircle,
	"plusCircleO":     PlusCircleO,
	"plusSquare":      PlusSquare,
	"plusSquareO":     PlusSquareO,
	"poweroff":        Poweroff,
	"printer":         Printer,
	"qq":              Qq,
	"qrcode":          Qrcode,
	"questionCircle":  QuestionCircle,
	"questionCircleO": QuestionCircleO,
	"reddit":          Reddit,
	"reload":          Reload,
	"right":           Right,
	"rightCircle":     RightCircle,
	"rightCircleO":    RightCircleO,
	"rightSquare":     RightSquare,
	"rightSquareO":    RightSquareO,
	"rollback":        Rollback,
	"safari":          Safari,
	"safety":          Safety,
	"save":            Save,
	"search":          Search,
	"setting":         Setting,
	"settingO":        SettingO,
	"share":           Share,
	"shoppingCart":    ShoppingCart,
	"shrink":          Shrink,
	"smile":           Smile,
	"smileO":          SmileO,
	"squareO":         SquareO,
	"starOff":         StarOff,
	"starOn":          StarOn,
	"stop":            Stop,
	"stopO":           StopO,
	"swap":            Swap,
	"swapLeft":        SwapLeft,
	"swapRight":       SwapRight,
	"table":           Table,
	"tag":             Tag,
	"tagO":            TagO,
	"tags":            Tags,
	"tagsO":           TagsO,
	"taobao":          Taobao,
	"time":            Time,
	"timeO":           TimeO,
	"twitter":         Twitter,
	"uiw":             Uiw,
	"unlock":          Unlock,
	"up":              Up,
	"upCircle":        UpCircle,
	"upCircleO":       UpCircleO,
	"upSquare":        UpSquare,
	"upSquareO":       UpSquareO,
	"upload":          Upload,
	"user":            User,
	"userAdd":         UserAdd,
	"userDelete":      UserDelete,
	"usergroupAdd":    UsergroupAdd,
	"usergroupDelete": UsergroupDelete,
	"verification":    Verification,
	"verticleLeft":    VerticleLeft,
	"verticleRight":   VerticleRight,
	"videoCamera":     VideoCamera,
	"warning":         Warning,
	"warningO":        WarningO,
	"weibo":           Weibo,
	"weixin":          Weixin,
	"wifi":            Wifi,
	"windows":         Windows,
	"woman":           Woman,
	"zoomIn":          ZoomIn,
	"zoomOut":         ZoomOut,
}

func Adobe(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M20 1v18.001L12.607 1zM7.399 1L0 19.001V1zm2.604 6.265L14.713 19h-3.086l-1.41-3.419H6.77z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alipay(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 13.692V3.846A3.845 3.845 0 0 0 16.154 0H3.846A3.845 3.845 0 0 0 0 3.846v12.308A3.845 3.845 0 0 0 3.846 20h12.308a3.848 3.848 0 0 0 3.787-3.167c-1.02-.439-5.44-2.347-7.742-3.45c-1.755 2.122-3.589 3.396-6.356 3.396s-4.613-1.703-4.388-3.79c.145-1.368 1.084-3.605 5.161-3.22c2.148.201 3.132.604 4.886 1.182c.45-.83.83-1.745 1.114-2.72H4.847v-.77H8.69V6.077H4V5.23h4.69V3.236s.045-.315.389-.315H11V5.23h5.002v.847H11v1.384h4.078a15.726 15.726 0 0 1-1.654 4.154c1.182.43 6.575 2.077 6.575 2.077M5.538 15.46c-2.925 0-3.384-1.846-3.23-2.617c.154-.77 1.002-1.768 2.625-1.768c1.87 0 3.541.477 5.547 1.454c-1.407 1.837-3.144 2.93-4.942 2.93"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Aliwangwang(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.23 6.572a9.611 9.611 0 0 0-2.104-3.073a9.875 9.875 0 0 0-6.94-2.83a9.882 9.882 0 0 0-3.226.533C4.191-.093.9-.002.756.003a.769.769 0 0 0-.7 1.052l1.727 4.277a9.513 9.513 0 0 0-1.415 5.002a9.48 9.48 0 0 0 .771 3.763a9.611 9.611 0 0 0 2.103 3.073a9.875 9.875 0 0 0 6.94 2.83a9.891 9.891 0 0 0 3.824-.76a9.776 9.776 0 0 0 5.223-5.143A9.472 9.472 0 0 0 20 10.334c0-1.304-.259-2.57-.77-3.762m-8.584 1.855c0 .442-.363.8-.812.8a.806.806 0 0 1-.812-.8V6.973c0-.442.363-.8.812-.8c.449 0 .812.358.812.8zm4.721 0c0 .442-.363.8-.812.8a.806.806 0 0 1-.813-.8V6.973c0-.442.364-.8.813-.8c.449 0 .812.358.812.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Android(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.735 1.828L13.6.254a.172.172 0 0 0-.3-.164l-.874 1.59a5.987 5.987 0 0 0-2.428-.508c-.869 0-1.69.181-2.424.506L6.698.09a.171.171 0 0 0-.232-.069a.169.169 0 0 0-.069.232l.866 1.572c-1.7.874-2.85 2.537-2.849 4.447l11.166-.001c0-1.91-1.146-3.57-2.845-4.444m-5.278 2.43a.467.467 0 1 1 .001-.934a.467.467 0 0 1 0 .935M15.542 6.7l.002 8.012c0 .481-.262.897-.648 1.127a1.31 1.31 0 0 1-.675.19l-.904.001v2.734a1.237 1.237 0 0 1-1.489 1.21a1.238 1.238 0 0 1-.99-1.209V16.03H9.163v2.735A1.237 1.237 0 0 1 7.925 20a1.239 1.239 0 0 1-1.238-1.235V16.03h-.901c-.521 0-.967-.3-1.182-.736a1.298 1.298 0 0 1-.141-.581l-.002-8.01zM2.74 6.47c.684 0 1.24.553 1.24 1.234v5.17a1.238 1.238 0 0 1-1.922 1.03a1.231 1.231 0 0 1-.557-1.03L1.5 7.703c0-.68.555-1.234 1.24-1.234m14.52-.001c.684 0 1.24.552 1.24 1.234v5.169c0 .683-.555 1.235-1.239 1.235c-.685 0-1.24-.552-1.24-1.235V7.702c0-.682.554-1.235 1.238-1.234M12.54 3.325a.467.467 0 1 1 0 .933a.467.467 0 1 1 0-.933"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AndroidO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.074 6.445a.29.29 0 0 0-.296.285v5.986a.29.29 0 0 0 .296.285a.29.29 0 0 0 .296-.285V6.73a.29.29 0 0 0-.296-.285m.089-1.318c.918 0 1.663.718 1.663 1.603v5.986c0 .885-.745 1.603-1.663 1.603S2.5 13.6 2.5 12.716V6.73c0-.885.745-1.603 1.663-1.603m1.74 1.501v8.13h8.201v-8.13zM4.534 5.31H15.47v9.887c0 .485-.408.879-.911.879H5.446c-.503 0-.91-.394-.91-.879zm11.372 1.12a.29.29 0 0 0-.296.284V12.7a.29.29 0 0 0 .296.285c.164 0 .296-.189.296-.285V6.714a.29.29 0 0 0-.296-.285m-6.89 8.775c.378 0 .684.295.684.659v2.32c0 1.109-.596 1.816-1.66 1.816c-1.046 0-1.662-.684-1.724-1.816v-2.32c0-.364.306-.66.684-.66c.377 0 .683.296.683.66v2.285c.024.446.102.533.358.533c.238 0 .292-.064.292-.498v-2.32c0-.364.306-.66.684-.66m4.542 0c.377 0 .683.295.683.659v2.32c0 1.109-.597 1.816-1.66 1.816c-1.046 0-1.662-.684-1.724-1.816v-2.32c0-.364.306-.66.684-.66c.377 0 .683.296.683.66v2.285c.024.446.102.533.358.533c.238 0 .292-.064.292-.498v-2.32c0-.364.306-.66.684-.66M12.664.27a.7.7 0 0 1 .955-.145a.644.644 0 0 1 .15.92l-.68.903a5.501 5.501 0 0 1 2.382 3.202c.117-.026.24-.04.366-.04c.918 0 1.663.718 1.663 1.603V12.7c0 .885-.745 1.603-1.663 1.603s-1.663-.718-1.663-1.603V6.714c0-.318.096-.615.263-.865l-.203.032c-.33-2.013-2.138-3.517-4.276-3.517c-2.211 0-4.061 1.607-4.302 3.709l-1.359-.145c.187-1.63 1.109-3.034 2.43-3.912l-.732-.97a.644.644 0 0 1 .15-.92a.7.7 0 0 1 .955.145l.85 1.127a5.872 5.872 0 0 1 2.008-.352c.66 0 1.295.109 1.887.31ZM8.636 3.337c.252 0 .456.197.456.44a.448.448 0 0 1-.456.439a.448.448 0 0 1-.456-.44c0-.242.204-.439.456-.439m2.734 0c.252 0 .456.197.456.44a.448.448 0 0 1-.456.439a.448.448 0 0 1-.455-.44c0-.242.204-.439.455-.439"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Apple(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.122 4.682c1.35 0 2.781.743 3.8 2.028c-3.34 1.851-2.797 6.674.578 7.963c-.465 1.04-.687 1.505-1.285 2.426c-.835 1.284-2.01 2.884-3.469 2.898c-1.295.012-1.628-.853-3.386-.843c-1.758.01-2.125.858-3.42.846c-1.458-.014-2.573-1.458-3.408-2.743C1.198 13.665.954 9.45 2.394 7.21C3.417 5.616 5.03 4.683 6.548 4.683c1.545 0 2.516.857 3.794.857c1.24 0 1.994-.858 3.78-.858M13.73 0c.18 1.215-.314 2.405-.963 3.247c-.695.902-1.892 1.601-3.05 1.565c-.21-1.163.332-2.36.99-3.167C11.43.755 12.67.074 13.73 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Appstore(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.07 10.91a2.02 2.02 0 0 1 2.02 2.02v5.05A2.02 2.02 0 0 1 7.07 20H2.02A2.02 2.02 0 0 1 0 17.98v-5.05a2.02 2.02 0 0 1 2.02-2.02zm10.91 0A2.02 2.02 0 0 1 20 12.93v5.05A2.02 2.02 0 0 1 17.98 20h-5.05a2.02 2.02 0 0 1-2.02-2.02v-5.05a2.02 2.02 0 0 1 2.02-2.02zM7.07 0a2.02 2.02 0 0 1 2.02 2.02v5.05a2.02 2.02 0 0 1-2.02 2.02H2.02A2.02 2.02 0 0 1 0 7.07V2.02A2.02 2.02 0 0 1 2.02 0zm10.91 0A2.02 2.02 0 0 1 20 2.02v5.05a2.02 2.02 0 0 1-2.02 2.02h-5.05a2.02 2.02 0 0 1-2.02-2.02V2.02A2.02 2.02 0 0 1 12.93 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppstoreO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.02 1.515a.505.505 0 0 0-.505.505v5.05c0 .28.226.506.505.506h5.05a.505.505 0 0 0 .506-.505V2.02a.505.505 0 0 0-.505-.506zm5.05 9.394a2.02 2.02 0 0 1 2.02 2.02v5.05A2.02 2.02 0 0 1 7.07 20H2.02A2.02 2.02 0 0 1 0 17.98v-5.05a2.02 2.02 0 0 1 2.02-2.02zm10.91 0A2.02 2.02 0 0 1 20 12.93v5.05A2.02 2.02 0 0 1 17.98 20h-5.05a2.02 2.02 0 0 1-2.02-2.02v-5.05a2.02 2.02 0 0 1 2.02-2.02zM7.07 12.424H2.02a.505.505 0 0 0-.505.505v5.05c0 .28.226.506.505.506h5.05a.505.505 0 0 0 .506-.505v-5.05a.505.505 0 0 0-.505-.506m10.91 0h-5.05a.505.505 0 0 0-.506.505v5.05c0 .28.226.506.505.506h5.05a.505.505 0 0 0 .506-.505v-5.05a.505.505 0 0 0-.505-.506M7.07 0a2.02 2.02 0 0 1 2.02 2.02v5.05a2.02 2.02 0 0 1-2.02 2.02H2.02A2.02 2.02 0 0 1 0 7.07V2.02A2.02 2.02 0 0 1 2.02 0zm10.91 0A2.02 2.02 0 0 1 20 2.02v5.05a2.02 2.02 0 0 1-2.02 2.02h-5.05a2.02 2.02 0 0 1-2.02-2.02V2.02A2.02 2.02 0 0 1 12.93 0zm0 1.515h-5.05a.505.505 0 0 0-.506.505v5.05c0 .28.226.506.505.506h5.05a.505.505 0 0 0 .506-.505V2.02a.505.505 0 0 0-.505-.506"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AreaChart(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.087 2.054c.198 0 .357.16.357.357c0 .018-.007.034-.01.052h.01v7.755c0 .197.16.357.357.357h7.78c.196 0 .356.16.356.357l-.002.01l.002.085c0 4.955-4.015 8.973-8.969 8.973C4.015 20 0 15.982 0 11.027C0 6.07 4.015 2.054 8.968 2.054l.096.002c.008.002.015-.002.023-.002M11.13 0c4.877.05 8.82 3.996 8.871 8.876c0 .197-.16.357-.357.357c-.018 0-.034-.008-.052-.01v.01H11.13a.357.357 0 0 1-.357-.357V.41h.01c-.002-.018-.01-.034-.01-.052c0-.197.16-.357.357-.357"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDown(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.364.765c-.037 5.419-.078 10.936-.12 16.553a398914.9 398914.9 0 0 0-6.454-6.361c-.385-.312-.802-.221-1.067.035c-.248.24-.353.7.007 1.09c2.368 2.326 4.975 4.892 7.82 7.697a.794.794 0 0 0 .554.222a.745.745 0 0 0 .539-.222c2.726-2.75 5.287-5.335 7.683-7.755a.754.754 0 0 0-.055-1.032c-.371-.374-.885-.229-1.093 0a3545.93 3545.93 0 0 1-6.386 6.437c.041-5.609.08-11.163.117-16.664c0-.26-.212-.765-.767-.765s-.778.469-.778.765"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.397 2.21A.726.726 0 0 1 8.41 2.2a.69.69 0 0 1 .012.99L2.373 9.243l16.931.105a.7.7 0 0 1-.008 1.4l-16.809-.104l6.303 6.16a.69.69 0 0 1 .076.903l-.075.087a.727.727 0 0 1-1.012 0L.21 10.4a.69.69 0 0 1-.005-.985Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M.696 9.349c5.518-.03 11.165-.062 16.942-.096c-2.807-2.81-4.835-4.842-6.084-6.095c-.143-.179-.158-.554.112-.847c.27-.293.752-.281.89-.14c2.362 2.372 4.772 4.782 7.23 7.23a.65.65 0 0 1 .215.503a.645.645 0 0 1-.215.502a8382.495 8382.495 0 0 1-7.6 7.421a.742.742 0 0 1-1.014-.063c-.263-.287-.29-.588.061-.982c2.002-1.96 4.097-4.004 6.287-6.13c-5.713.038-11.321.07-16.824.097a.701.701 0 0 1-.696-.72c0-.507.388-.68.696-.68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUp(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m9.551.222l-7.82 7.696c-.306.416-.309.78-.008 1.091c.301.312.665.314 1.093.007l6.428-6.334V19.32c.041.453.319.68.833.681c.514.001.792-.226.833-.68l-.111-16.64l6.38 6.328c.415.321.78.321 1.092 0c.313-.32.334-.663.062-1.027l-7.7-7.76A.776.776 0 0 0 10.078 0a.674.674 0 0 0-.526.222"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsAlt(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.163 10.762a.7.7 0 0 1 .008.99l-6.737 6.854l4.266-.001a.7.7 0 0 1 0 1.4h-6a.7.7 0 0 1-.7-.7v-6a.7.7 0 0 1 1.4 0l.001 4.354l6.772-6.888a.7.7 0 0 1 .99-.009M19.302 0a.7.7 0 0 1 .7.7v6a.7.7 0 0 1-1.4 0L18.6 2.346l-6.772 6.888a.7.7 0 0 1-.998-.981l6.737-6.854l-4.266.001a.7.7 0 0 1 0-1.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Asterisk(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c.412 0 .746.344.746.768l-.001 7.725l6.07-4.54a.73.73 0 0 1 .971.082l.072.087a.783.783 0 0 1-.166 1.073l-6.424 4.804l6.424 4.806c.334.25.408.73.166 1.073a.732.732 0 0 1-1.043.17l-6.07-4.542l.001 7.726A.757.757 0 0 1 10 20a.757.757 0 0 1-.746-.768v-7.726l-6.07 4.542a.73.73 0 0 1-.97-.083l-.072-.087a.783.783 0 0 1 .166-1.073L8.73 10L2.308 5.195a.783.783 0 0 1-.166-1.073a.732.732 0 0 1 1.043-.17l6.069 4.541V.768C9.254.344 9.588 0 10 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Backward(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.903 2.931l-.001 5.108l6.615-5.593c.62-.526 1.58-.323 1.58.485v14.14c0 .805-.96 1.009-1.58.483l-6.615-5.593v5.11c0 .805-.96 1.009-1.58.483l-8.085-6.836a.936.936 0 0 1 0-1.434l8.086-6.838c.62-.526 1.58-.323 1.58.485"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Baidu(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.723 10.54c2.132-.47 1.838-3.07 1.777-3.639c-.104-.875-1.115-2.406-2.486-2.283c-1.723.155-1.974 2.7-1.974 2.7c-.235 1.174.557 3.687 2.683 3.22M7.68 6.172c1.176 0 2.127-1.383 2.127-3.09C9.807 1.382 8.859 0 7.683 0C6.507 0 5.551 1.375 5.551 3.083c0 1.708.956 3.09 2.132 3.09m5.069.205c1.576.217 2.582-1.5 2.786-2.8c.204-1.29-.817-2.798-1.927-3.056c-1.12-.264-2.5 1.56-2.638 2.749c-.147 1.458.204 2.907 1.772 3.113m6.24 2.184c0-.621-.5-2.493-2.376-2.493c-1.875 0-2.132 1.766-2.132 3.015c0 1.192.097 2.85 2.438 2.8c2.332-.059 2.077-2.7 2.077-3.324m-2.375 5.447s-2.438-1.924-3.86-3.999c-1.927-3.063-4.667-1.816-5.58-.263c-.915 1.569-2.336 2.551-2.537 2.813c-.204.259-2.94 1.766-2.33 4.516c.612 2.749 2.744 2.699 2.744 2.699s1.568.158 3.397-.258c1.83-.417 3.397.1 3.397.1s4.253 1.457 5.43-1.342c1.163-2.807-.662-4.257-.662-4.257"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarChart(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 6a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H1a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1zm5.365 6a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1v-6a1 1 0 0 1 1-1zm5.332-12a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1V1a1 1 0 0 1 1-1zM19 0a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1V1a1 1 0 0 1 1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Barcode(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.438 2.5h.889v12.73h-.889zm-2.22 0H7.55v12.73H6.218zm-.888 0h.444v12.73H5.33zm-3.11 0h.445v12.73H2.22zm.889 0H4.44v12.73H3.11zm12.45 0h1.332v12.73H15.56zm-2.665 0h1.332v12.73h-1.332zm1.776 0h.445v12.73h-.445zm-1.776 13.638h4.885V17.5h-4.885zM0 2.5h1.332v15H0zm10.673 13.638h1.333V17.5h-1.333zM17.335 2.5h.444v12.73h-.444zm1.333 0H20v15h-1.332zm-7.106 0h.888v12.73h-.888zM2.22 16.138h1.332V17.5H2.221zm2.22 0h5.33V17.5H4.44zM10.23 2.5h.444v12.73h-.444z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.985 0c1.089 0 1.971.898 1.971 2.006l-.009.163c.868.352 1.707.936 2.451 1.71c.862.893 1.366 2.077 1.521 3.596v5.478l1.191 2.098c.4.666.528 1.224.216 1.707c-.286.441-.797.595-1.49.583h-2.67C12.854 18.86 11.532 20 9.95 20c-1.584 0-2.905-1.14-3.216-2.658H3.778l-.056-.003c-.627-.054-1.094-.357-1.199-.94c-.071-.397.023-.823.268-1.331l1.225-2.18l.003-5.473c.107-1.21.56-2.337 1.348-3.371c.667-.875 1.62-1.519 2.654-1.89a1.752 1.752 0 0 1-.006-.148C8.015.898 8.897 0 9.985 0m1.818 17.342H8.097c.275.77 1 1.32 1.853 1.32c.852 0 1.578-.55 1.853-1.32M10.082 3.124c-1.354 0-2.843.645-3.677 1.74c-.638.836-.994 1.722-1.075 2.61v5.59c0 .117-.03.232-.087.333l-1.291 2.296a1.71 1.71 0 0 0-.12.311h12.014c.121.002.213-.003.276-.005a2.615 2.615 0 0 0-.141-.265l-1.287-2.267a.678.678 0 0 1-.088-.335l.003-5.586c-.121-1.162-.506-2.064-1.149-2.732c-1.04-1.08-2.262-1.69-3.378-1.69m-.097-1.787a.66.66 0 0 0-.635.497c.246-.031.49-.047.732-.047c.177 0 .356.01.535.032a.66.66 0 0 0-.632-.482"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.818 5.174a.449.449 0 0 0-.454.443V16.23c0 .244.203.442.454.442h16.364a.449.449 0 0 0 .454-.442V5.617a.449.449 0 0 0-.454-.443zm0-1.326h16.364c1.004 0 1.818.792 1.818 1.769V16.23c0 .977-.814 1.769-1.818 1.769H1.818C.814 18 0 17.208 0 16.231V5.617c0-.977.814-1.77 1.818-1.77m-.79 5.528a.673.673 0 0 1-.681-.664c0-.366.305-.663.682-.663h4.5c.376 0 .681.297.681.663a.673.673 0 0 1-.682.664zm13.109 0a.673.673 0 0 1-.682-.664c0-.366.305-.663.682-.663h4.975c.376 0 .682.297.682.663a.673.673 0 0 1-.682.664zm-4.3 6.52c-2.887 0-5.227-2.277-5.227-5.086c0-2.81 2.34-5.086 5.227-5.086c2.887 0 5.228 2.277 5.228 5.086s-2.34 5.086-5.228 5.086m0-1.327c2.134 0 3.864-1.683 3.864-3.76c0-2.075-1.73-3.759-3.864-3.759c-2.133 0-3.863 1.684-3.863 3.76c0 2.076 1.73 3.759 3.863 3.759M6.852 4.673l-1.2-.63c.292-.526.623-.976.994-1.35A2.263 2.263 0 0 1 8.29 2l6.371.002a2.49 2.49 0 0 1 1.657.777c.355.374.668.79.94 1.246l-1.178.666a5.773 5.773 0 0 0-.765-1.014a1.089 1.089 0 0 0-.709-.35H8.29c-.255 0-.463.089-.663.29c-.28.282-.54.633-.774 1.056M2.5 3.181c0-.367.305-.664.682-.664c.376 0 .682.297.682.664V4.29a.673.673 0 0 1-.682.663a.673.673 0 0 1-.682-.663z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretDown(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m17.704 6.58l-6.972 8.086a.967.967 0 0 1-1.463 0L2.296 6.58C1.76 5.96 1.967 5 2.791 5h14.42c.821 0 1.029.96.493 1.58"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretLeft(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m13.42 2.296l-8.047 6.94A.923.923 0 0 0 5 10c0 .315.137.58.41.797l8.01 6.907c.32.268.664.357 1.032.267c.3-.087.482-.29.548-.607v-14.7a.758.758 0 0 0-.526-.628c-.383-.091-.734-.005-1.054.26"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretRight(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 2.643v14.765c.092.32.299.511.619.572c.32.061.633-.024.94-.255l8.107-6.993A.944.944 0 0 0 15 10a.94.94 0 0 0-.334-.73L6.58 2.295c-.232-.197-.639-.383-1.061-.253c-.282.087-.455.287-.519.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretUp(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m10.685 5.284l7.019 8.136c.24.277.339.595.296.954c-.043.359-.28.568-.71.626H2.744c-.351-.025-.582-.199-.693-.522c-.11-.324-.054-.651.17-.981l7.097-8.213a.962.962 0 0 1 1.367 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.635 4.243a.795.795 0 0 1 1.12-.023a.788.788 0 0 1 .024 1.117L8.787 16.757a.795.795 0 0 1-1.137.008L.228 9.262a.788.788 0 0 1 .008-1.117a.795.795 0 0 1 1.122.008l6.849 6.924z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckSquare(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.182 0C19.186 0 20 .814 20 1.818v16.364A1.818 1.818 0 0 1 18.182 20H1.818A1.818 1.818 0 0 1 0 18.182V1.818C0 .814.814 0 1.818 0zm-3.26 6.564a.682.682 0 0 0-.963.053L9.547 11.55L6.715 8.52a.682.682 0 0 0-.996.931l3.34 3.575a.682.682 0 0 0 1.007-.01l4.91-5.489a.682.682 0 0 0-.054-.962"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckSquareO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.818 1.364a.455.455 0 0 0-.454.454v16.364c0 .25.203.454.454.454h16.364a.455.455 0 0 0 .454-.454V1.818a.455.455 0 0 0-.454-.454zM18.182 0C19.186 0 20 .814 20 1.818v16.364A1.818 1.818 0 0 1 18.182 20H1.818A1.818 1.818 0 0 1 0 18.182V1.818C0 .814.814 0 1.818 0zm-3.26 6.564a.682.682 0 0 0-.963.053L9.547 11.55L6.715 8.52a.682.682 0 0 0-.996.931l3.34 3.575a.682.682 0 0 0 1.007-.01l4.91-5.489a.682.682 0 0 0-.054-.962"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chrome(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.39 11.918c1.06 2.086 3.303 3.358 5.624 2.923l-2.566 5.033A9.985 9.985 0 0 1 0 9.998c0-2.02.603-3.905 1.63-5.479zm9.608-10.579a10.021 10.021 0 0 1 3.928 4.151l-8.28-.435C8.302 4.921 6.07 6.238 5.3 8.447L2.22 3.716A9.991 9.991 0 0 1 9.965 0a9.928 9.928 0 0 1 5.033 1.34M13.368 10a3.367 3.367 0 0 1-3.37 3.37A3.367 3.367 0 0 1 6.628 10a3.367 3.367 0 0 1 3.37-3.37a3.367 3.367 0 0 1 3.37 3.37m1.63 8.658a9.976 9.976 0 0 1-5.557 1.328l4.52-6.951c1.283-1.976 1.26-4.553-.269-6.339l5.636-.29a9.99 9.99 0 0 1-4.33 12.252"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleCheck(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m4.922 6.583a.682.682 0 0 0-.963.053l-4.412 4.932l-2.832-3.03a.682.682 0 1 0-.996.931l3.34 3.575a.682.682 0 0 0 1.007-.01l4.91-5.489a.682.682 0 0 0-.054-.962"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleCheckO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21m4.922 5.188c.28.25.304.682.053.962l-4.909 5.488a.682.682 0 0 1-1.006.011L5.72 9.47a.682.682 0 0 1 .995-.931l2.832 3.03l4.412-4.932a.682.682 0 0 1 .963-.053"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleClose(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m2.207 6.837L10.01 9.03L7.815 6.837a.682.682 0 0 0-.88-.072l-.084.072a.682.682 0 0 0 0 .964l2.195 2.193l-2.195 2.193a.682.682 0 1 0 .964.965l2.195-2.195l2.197 2.195c.24.24.613.263.88.071l.084-.072a.682.682 0 0 0 0-.964l-2.196-2.193l2.195-2.193a.682.682 0 0 0-.963-.964"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleCloseO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21m2.207 5.442a.682.682 0 0 1 .963.964l-2.195 2.193l2.195 2.193a.682.682 0 0 1-.963.965l-2.197-2.195l-2.195 2.195a.682.682 0 0 1-.88.071l-.084-.072a.682.682 0 0 1 0-.964l2.195-2.193l-2.195-2.193a.682.682 0 1 1 .964-.964L10.01 9.03Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Close(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.219 2.154l6.778 6.773l6.706-6.705c.457-.407.93-.164 1.119.04a.777.777 0 0 1-.044 1.035l-6.707 6.704l6.707 6.702c.298.25.298.74.059 1.014c-.24.273-.68.431-1.095.107l-6.745-6.749l-6.753 6.752c-.296.265-.784.211-1.025-.052c-.242-.264-.334-.72-.025-1.042l6.729-6.732l-6.701-6.704c-.245-.27-.33-.764 0-1.075c.33-.311.822-.268.997-.068"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseSquare(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.182 0C19.186 0 20 .814 20 1.818v16.364A1.818 1.818 0 0 1 18.182 20H1.818A1.818 1.818 0 0 1 0 18.182V1.818C0 .814.814 0 1.818 0zm-5.975 6.818L10.01 9.01L7.815 6.818a.682.682 0 0 0-.88-.072l-.084.072a.682.682 0 0 0 0 .964l2.195 2.193l-2.195 2.193a.682.682 0 1 0 .964.965l2.195-2.195l2.197 2.195c.24.24.613.263.88.071l.084-.072a.682.682 0 0 0 0-.964l-2.196-2.193l2.195-2.193a.682.682 0 0 0-.963-.964"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseSquareO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.818 1.364a.455.455 0 0 0-.454.454v16.364c0 .25.203.454.454.454h16.364a.455.455 0 0 0 .454-.454V1.818a.455.455 0 0 0-.454-.454zM18.182 0C19.186 0 20 .814 20 1.818v16.364A1.818 1.818 0 0 1 18.182 20H1.818A1.818 1.818 0 0 1 0 18.182V1.818C0 .814.814 0 1.818 0zm-5.975 6.818L10.01 9.01L7.815 6.818a.682.682 0 0 0-.88-.072l-.084.072a.682.682 0 0 0 0 .964l2.195 2.193l-2.195 2.193a.682.682 0 1 0 .964.965l2.195-2.195l2.197 2.195c.24.24.613.263.88.071l.084-.072a.682.682 0 0 0 0-.964l-2.196-2.193l2.195-2.193a.682.682 0 0 0-.963-.964"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDownload(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 2.5a4.499 4.499 0 0 1 4.492 4.77H16l-.001 1.227a4 4 0 0 1 3.996 3.799l.005.2a4 4 0 0 1-3.8 3.992l-.2.005h-.001L16 16.5H4a4.01 4.01 0 0 1-4-4.005a4 4 0 0 1 3.198-3.918a3 3 0 0 1 4.313-3.664A4.498 4.498 0 0 1 11.5 2.5m-.046 4h-2v4.922L7 11.42l3.454 4.076l3.464-4.073h-2.464z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDownloadO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.668 8.668a.712.712 0 0 1-.611.737c-1.282.155-2.201.552-2.78 1.17c-.576.612-.874 1.41-.899 2.35c.09 1.007.422 1.762.997 2.302c.572.537 1.431.822 2.603.83l10.67.001c.996-.067 1.73-.444 2.251-1.139c.659-.876.797-1.665.682-2.547c-.086-.662-.495-1.502-1.157-2.073c-.505-.436-1.3-.701-2.401-.77c-.374-.023-.661-.353-.65-.743a4.835 4.835 0 0 0-.912-3.034c-.872-1.23-2.223-1.857-3.551-1.806c-1.299.05-2.502.663-3.371 1.822c-.612.816-.9 1.772-.871 2.9m8.903-3.775a6.288 6.288 0 0 1 1.17 3.257c1.069.146 1.925.486 2.563 1.036c.944.815 1.516 1.987 1.646 2.991c.161 1.237-.045 2.41-.964 3.633c-.774 1.03-1.875 1.596-3.295 1.69H4.973c-1.489-.01-2.672-.402-3.522-1.2C.605 15.506.122 14.412 0 12.975c.03-1.376.46-2.525 1.29-3.41c.717-.765 1.725-1.26 3.009-1.501c.076-1.195.46-2.263 1.153-3.185C6.577 3.377 8.163 2.57 9.859 2.505c1.768-.068 3.559.763 4.712 2.389m-4.415 2.128c-.382 0-.691.323-.691.721l-.001 4.143l-.787-.81a.682.682 0 0 0-.897-.082l-.086.075a.733.733 0 0 0 .007 1.014l1.98 2.041a.68.68 0 0 0 .982.007l1.954-2.014a.733.733 0 0 0-.006-1.014a.68.68 0 0 0-.984-.007l-.781.805V7.742a.712.712 0 0 0-.597-.715Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUpload(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 3a4.499 4.499 0 0 1 4.492 4.77H16l-.001 1.227a4 4 0 0 1 3.996 3.799l.005.2a4 4 0 0 1-3.8 3.992l-.2.005h-.001L16 17h-5.001v-3.923h2.464L10 9.003l-3.454 4.075H9L8.999 17H4a4.01 4.01 0 0 1-4-4.005a4 4 0 0 1 3.198-3.918a3 3 0 0 1 4.313-3.664A4.498 4.498 0 0 1 11.5 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUploadO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.668 8.668a.712.712 0 0 1-.611.737c-1.282.155-2.201.552-2.78 1.17c-.576.612-.874 1.41-.899 2.35c.09 1.007.422 1.762.997 2.302c.572.537 1.431.822 2.603.83l10.67.001c.996-.067 1.73-.444 2.251-1.139c.659-.876.797-1.665.682-2.547c-.086-.662-.495-1.502-1.157-2.073c-.505-.436-1.3-.701-2.401-.77c-.374-.023-.661-.353-.65-.743a4.835 4.835 0 0 0-.912-3.034c-.872-1.23-2.223-1.857-3.551-1.806c-1.299.05-2.502.663-3.371 1.822c-.612.816-.9 1.772-.871 2.9m8.903-3.775a6.288 6.288 0 0 1 1.17 3.257c1.069.146 1.925.486 2.563 1.036c.944.815 1.516 1.987 1.646 2.991c.161 1.237-.045 2.41-.964 3.633c-.774 1.03-1.875 1.596-3.295 1.69H4.973c-1.489-.01-2.672-.402-3.522-1.2C.605 15.506.122 14.412 0 12.975c.03-1.376.46-2.525 1.29-3.41c.717-.765 1.725-1.26 3.009-1.501c.076-1.195.46-2.263 1.153-3.185C6.577 3.377 8.163 2.57 9.859 2.505c1.768-.068 3.559.763 4.712 2.389M9.68 7.135L7.7 9.177a.733.733 0 0 0-.007 1.014a.68.68 0 0 0 .983-.007l.787-.811v4.239c0 .365.26.667.598.715l.094.007c.381 0 .69-.324.69-.722V9.358l.781.806a.68.68 0 0 0 .984-.007a.733.733 0 0 0 .006-1.014L10.663 7.13a.68.68 0 0 0-.983.006"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coffee(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.382 8.505v5.058a5.057 5.057 0 0 0 5.057 5.058h3.677a5.057 5.057 0 0 0 5.057-5.058V8.506zM11.887.16a.69.69 0 0 1 .086.972c-.642.765-.784 1.287-.586 1.637c.062.109.593.948.715 1.207c.276.585.312 1.152.074 1.822a4.622 4.622 0 0 1-.751 1.328h3.881c.437.016.754.127.95.335c.11.114.188.258.237.432l.06-.002a3.448 3.448 0 0 1 0 6.897l-.116-.004A6.438 6.438 0 0 1 10.117 20H6.438a6.436 6.436 0 0 1-6.436-6.437V8.337c-.02-.433.062-.74.244-.92c.183-.18.453-.277.809-.29h2.953a.689.689 0 0 1 .144-.17C4.762 6.44 5.16 5.9 5.36 5.337c.114-.32.101-.51-.022-.771c-.078-.166-.569-.942-.667-1.116c-.539-.952-.242-2.044.728-3.202a.69.69 0 1 1 1.057.886c-.642.765-.783 1.287-.585 1.637c.061.109.593.948.715 1.207c.275.585.312 1.152.073 1.822a4.622 4.622 0 0 1-.75 1.328h.858a.689.689 0 0 1 .144-.17C7.52 6.44 7.918 5.9 8.118 5.337c.114-.32.102-.51-.022-.771c-.078-.166-.569-.942-.667-1.116c-.539-.952-.242-2.044.729-3.202a.69.69 0 1 1 1.056.886c-.641.765-.783 1.287-.585 1.637c.062.109.593.948.715 1.207c.276.585.312 1.152.073 1.822a4.622 4.622 0 0 1-.75 1.328h.859a.689.689 0 0 1 .143-.17c.61-.518 1.007-1.058 1.207-1.621c.114-.32.102-.51-.022-.771c-.078-.166-.568-.942-.667-1.116c-.538-.952-.242-2.044.729-3.202a.69.69 0 0 1 .971-.086m4.665 9.11v4.138a2.069 2.069 0 0 0 0-4.138"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Component(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.62 3.448h-3.448a3.448 3.448 0 0 0-6.896 0H4.828a1.38 1.38 0 0 0-1.38 1.38v3.448a3.448 3.448 0 1 0 0 6.896v3.449A1.38 1.38 0 0 0 4.828 20h4.827v-1.38a2.069 2.069 0 1 1 4.138 0V20h4.828A1.38 1.38 0 0 0 20 18.62v-4.827h-1.38a2.069 2.069 0 1 1 0-4.138H20V4.828a1.38 1.38 0 0 0-1.38-1.38m-3.448 8.276a3.448 3.448 0 0 0 3.449 3.448v3.449h-3.449a3.448 3.448 0 1 0-6.896 0H4.828v-4.828h-1.38a2.069 2.069 0 1 1 0-4.138h1.38V4.828h4.827v-1.38a2.069 2.069 0 1 1 4.138 0v1.38h4.828v3.448a3.448 3.448 0 0 0-3.449 3.448"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copy(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.644 2.983a.252.252 0 0 0-.253.252c0 .139.113.251.253.251h3.713c.14 0 .253-.112.253-.251a.252.252 0 0 0-.253-.252zm3.713-1.342c.734 0 1.353.49 1.544 1.16l2.175.001c.621.004 1.122.205 1.432.638c.266.372.372.85.345 1.387L15.85 17.84c.042.552-.062 1.04-.328 1.445c-.312.473-.821.71-1.452.716H3.14c-.76-.03-1.323-.209-1.675-.609c-.327-.371-.47-.88-.464-1.5V4.84c-.013-.6.154-1.106.518-1.48c.376-.384.932-.554 1.647-.559h1.935c.19-.67.809-1.16 1.543-1.16zm0 3.187H6.644c-.546 0-1.027-.27-1.317-.684H3.17c-.383.002-.602.07-.682.152c-.091.093-.144.252-.138.531v13.07c-.003.325.052.522.13.61c.054.061.286.135.685.151h10.9c.2-.002.28-.04.326-.109c.091-.138.133-.334.11-.658l.001-13.096c.014-.293-.027-.482-.096-.578c-.026-.035-.116-.072-.336-.073h-2.397c-.29.414-.771.684-1.317.684M17.2 0c.994 0 1.8.801 1.8 1.79v14.082c0 .988-.806 1.79-1.8 1.79h-1.958v-1.343h1.957c.249 0 .45-.2.45-.447V1.789a.449.449 0 0 0-.45-.447H9.643c-.248 0-.45.2-.45.447v.157h-1.35v-.157C7.843.801 8.649 0 9.643 0zM8.196 11.751c.373 0 .675.3.675.671c0 .37-.302.671-.675.671H4.145a.673.673 0 0 1-.676-.67c0-.371.303-.672.676-.672zm4.052-2.684c.372 0 .675.3.675.671c0 .37-.303.671-.675.671H4.145a.673.673 0 0 1-.676-.67c0-.371.303-.672.676-.672zm0-2.684c.372 0 .675.3.675.671a.673.673 0 0 1-.675.671H4.145a.673.673 0 0 1-.676-.67c0-.371.303-.672.676-.672z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copyright(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21m.16 3.256c1.08 0 2.117.303 3.003.864a.698.698 0 0 1-.745 1.179a4.211 4.211 0 0 0-2.257-.647c-2.278 0-4.114 1.775-4.114 3.953s1.836 3.953 4.114 3.953c.801 0 1.567-.22 2.224-.627a.698.698 0 0 1 .735 1.187a5.608 5.608 0 0 1-2.96.836c-3.037 0-5.509-2.39-5.509-5.349c0-2.96 2.472-5.349 5.51-5.349"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CssThree(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m1.5 0l1.594 18.056L10 20l6.906-1.944L18.5 0zm13.577 5.852L9.994 8.125l-.013.005h4.917l-.563 6.762l-4.334 1.323l-.007-.003v.003l-4.358-1.348l-.28-3.405h2.162l.14 1.764l2.316.611l.02-.006v.003l2.397-.706l.164-2.842l-2.561-.008l-4.78-.016l-.163-2.132l4.943-2.153l.288-.125H4.864l-.259-2.18h10.683z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cut(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.394.22a.7.7 0 0 1 .986-.033a.69.69 0 0 1 .034.98l-7.473 7.948l3.269 3.476a3.956 3.956 0 0 1 1.837-.448c2.183 0 3.953 1.759 3.953 3.928C20 18.241 18.23 20 16.047 20c-2.184 0-3.954-1.759-3.954-3.929a3.9 3.9 0 0 1 1.011-2.624l-3.117-3.316l-3.104 3.303a3.9 3.9 0 0 1 1.024 2.637c0 2.17-1.77 3.929-3.954 3.929C1.77 20 0 18.241 0 16.071s1.77-3.928 3.953-3.928c.657 0 1.276.158 1.82.44l3.259-3.468l-7.47-7.948a.69.69 0 0 1 .033-.98a.7.7 0 0 1 .986.033L9.987 8.1ZM3.954 13.53a2.55 2.55 0 0 0-2.559 2.541a2.55 2.55 0 0 0 2.558 2.542a2.55 2.55 0 0 0 2.559-2.542a2.55 2.55 0 0 0-2.559-2.542m12.093 0a2.55 2.55 0 0 0-2.559 2.541a2.55 2.55 0 0 0 2.559 2.542a2.55 2.55 0 0 0 2.558-2.542a2.55 2.55 0 0 0-2.558-2.542"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DarrowLeft(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.397 2.21A.726.726 0 0 1 9.41 2.2a.69.69 0 0 1 .012.99l-6.7 6.707l7.07 6.908a.69.69 0 0 1 0 .99a.727.727 0 0 1-1.012 0L1.21 10.4a.69.69 0 0 1-.005-.985Zm9 0a.726.726 0 0 1 1.012-.01a.69.69 0 0 1 .012.99l-6.7 6.707l7.07 6.908a.69.69 0 0 1 0 .99a.727.727 0 0 1-1.012 0L10.21 10.4a.69.69 0 0 1-.005-.985Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DarrowRight(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m2.542 2.154l7.254 7.26c.136.14.204.302.204.483a.73.73 0 0 1-.204.5l-7.575 7.398c-.383.317-.724.317-1.022 0c-.299-.317-.299-.643 0-.98l7.08-6.918l-6.754-6.763c-.237-.343-.215-.654.066-.935c.281-.28.598-.295.951-.045m9 0l7.254 7.26c.136.14.204.302.204.483a.73.73 0 0 1-.204.5l-7.575 7.398c-.383.317-.724.317-1.022 0c-.299-.317-.299-.643 0-.98l7.08-6.918l-6.754-6.763c-.237-.343-.215-.654.066-.935c.281-.28.598-.295.951-.045"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dcaret(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.407 11.333c.618 0 .773.64.37 1.054l-5.23 5.39a.786.786 0 0 1-1.096 0l-5.229-5.39c-.402-.414-.246-1.054.37-1.054Zm-4.86-9.11l5.23 5.39c.403.414.248 1.054-.37 1.054H4.592c-.616 0-.772-.64-.37-1.054l5.229-5.39a.786.786 0 0 1 1.097 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dashboard(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m.667 1.359v1.035a.667.667 0 0 1-1.334 0V1.359A8.614 8.614 0 0 0 5.637 2.51l.522.584a.667.667 0 0 1-.995.888l-.63-.707a8.714 8.714 0 0 0-1.776 1.962l.843.506a.667.667 0 0 1-.686 1.143l-.803-.481a8.607 8.607 0 0 0-.709 2.491h.907a.667.667 0 1 1 0 1.334l-.973-.001v.031a8.627 8.627 0 0 0 .742 3.263l.836-.559a.667.667 0 0 1 .741 1.109l-.939.627A8.66 8.66 0 0 0 10 18.667a8.662 8.662 0 0 0 7.447-4.23l-1.132-.757a.667.667 0 0 1 .74-1.109l.989.661a8.633 8.633 0 0 0 .62-3.003H17.58a.667.667 0 0 1 0-1.333h1.017a8.608 8.608 0 0 0-.57-2.168l-.95.492a.667.667 0 1 1-.612-1.184l.965-.5a8.71 8.71 0 0 0-1.839-2.158l-.602.789a.667.667 0 1 1-1.06-.81l.58-.76a8.615 8.615 0 0 0-3.842-1.238m3.248 5.46a.667.667 0 0 1-.104.937l-2.04 1.631l.007.12c0 .692-.529 1.262-1.205 1.326l-.129.006a1.333 1.333 0 1 1 .558-2.544l1.976-1.58a.667.667 0 0 1 .937.104"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Date(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.673 0a.7.7 0 0 1 .7.7v1.309h7.517v-1.3a.7.7 0 0 1 1.4 0v1.3H18a2 2 0 0 1 2 1.999v13.993A2 2 0 0 1 18 20H2a2 2 0 0 1-2-1.999V4.008a2 2 0 0 1 2-1.999h2.973V.699a.7.7 0 0 1 .7-.699M1.4 7.742v10.259a.6.6 0 0 0 .6.6h16a.6.6 0 0 0 .6-.6V7.756zm5.267 6.877v1.666H5v-1.666zm4.166 0v1.666H9.167v-1.666zm4.167 0v1.666h-1.667v-1.666zm-8.333-3.977v1.666H5v-1.666zm4.166 0v1.666H9.167v-1.666zm4.167 0v1.666h-1.667v-1.666zM4.973 3.408H2a.6.6 0 0 0-.6.6v2.335l17.2.014V4.008a.6.6 0 0 0-.6-.6h-2.71v.929a.7.7 0 0 1-1.4 0v-.929H6.373v.92a.7.7 0 0 1-1.4 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Delete(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.129 0l1.974.005c.778.094 1.46.46 2.022 1.078c.459.504.7 1.09.714 1.728h5.475a.69.69 0 0 1 .686.693a.689.689 0 0 1-.686.692l-1.836-.001v11.627c0 2.543-.949 4.178-3.041 4.178H5.419c-2.092 0-3.026-1.626-3.026-4.178V4.195H.686A.689.689 0 0 1 0 3.505c0-.383.307-.692.686-.692h5.47c.014-.514.205-1.035.554-1.55C7.23.495 8.042.074 9.129 0m6.977 4.195H3.764v11.627c0 1.888.52 2.794 1.655 2.794h9.018c1.139 0 1.67-.914 1.67-2.794zM6.716 6.34c.378 0 .685.31.685.692v8.05a.689.689 0 0 1-.686.692a.689.689 0 0 1-.685-.692v-8.05c0-.382.307-.692.685-.692m2.726 0c.38 0 .686.31.686.692v8.05a.689.689 0 0 1-.686.692a.689.689 0 0 1-.685-.692v-8.05c0-.382.307-.692.685-.692m2.728 0c.378 0 .685.31.685.692v8.05a.689.689 0 0 1-.685.692a.689.689 0 0 1-.686-.692v-8.05a.69.69 0 0 1 .686-.692M9.176 1.382c-.642.045-1.065.264-1.334.662c-.198.291-.297.543-.313.768l4.938-.001c-.014-.291-.129-.547-.352-.792c-.346-.38-.73-.586-1.093-.635z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dingding(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.466 7.475c-.034.147-.115.36-.228.618h.003l-.013.025c-.666 1.497-2.405 4.433-2.405 4.433s-.002-.007-.008-.018l-.509.929h2.45L12.075 20l1.062-4.447H11.21l.67-2.94c-.543.139-1.183.326-1.94.583c0 0-1.026.632-2.957-1.214c0 0-1.302-1.206-.547-1.505c.322-.127 1.558-.29 2.533-.428c1.315-.188 2.124-.286 2.124-.286s-4.056.065-5.018-.096C5.113 9.508 3.89 7.82 3.632 6.339c0 0-.403-.815.864-.43c1.266.387 6.51 1.501 6.51 1.501S4.184 5.213 3.731 4.678C3.278 4.142 2.395 1.753 2.51.285c0 0 .05-.366.407-.268c0 0 5.041 2.42 8.488 3.745c3.45 1.324 6.449 1.998 6.061 3.713"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DislikeO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-opacity=".89" fill-rule="evenodd" d="M15.807.939c.396.367.655 1.133.706 1.595c.59.366 1.288 1.104 1.349 2.494c1.053.731 1.853 2.083.854 4.06c.58.54 1.227 2.188.395 3.516c-.969 1.552-3.075 1.66-5.174 1.383c.56 1.565.77 3.009-.116 4.488C12.94 19.787 11.724 20 11.308 20c-1.138 0-1.918-.979-2.234-2.283c-.115-.364-.246-1.224-.297-1.45c-.265-1.44-1.156-2.592-2.67-3.453a11.392 11.392 0 0 0-2.123-.946h-2.34c-.521 0-1.144-.527-1.144-1.165V3.067c.074-.722.475-1.082 1.202-1.082h3.11c1.364-.31 2.724-.642 4.079-.995C10.2.637 10.487.52 11.403.268c2.053-.532 3.478-.24 4.404.67m-2.382.424c-.819 0-1.856.252-2.316.399c-.162.051-.446.135-.745.221l-.3.087l-.288.082l-.56.158s-1.41.378-4.173 1.02c-.103.012-.158.02-.166.022v7.38c1.511.582 2.688 1.288 3.53 2.118c1.264 1.244 1.615 2.368 1.822 3.807c.118.723.309 1.306.597 1.705a.65.65 0 0 0 .342.251c.147.047.35.05.783-.184c.433-.236.99-.853 1.095-1.772c.07-.893-.17-1.667-.492-2.481a15.705 15.705 0 0 0-.357-.822c-.218-.413.11-1.099.786-.95c.906.255 3.154.6 4.422 0c.737-.427.92-1.116.547-2.066a1.74 1.74 0 0 0-.495-.553c-.17-.102-.502-.544-.103-1.045c.396-.635.975-1.928-.49-2.734a.656.656 0 0 1-.34-.57c-.02-.274.024-1.29-.73-1.744c-.18-.097-.397-.177-.52-.41c-.078-.154-.103-.528-.103-.528c-.103-.632-.245-1.222-1.746-1.391M3.519 3.348H1.861v7.157h1.658z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Document(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.226 0C18.206 0 19 .814 19 1.818v16.364C19 19.186 18.206 20 17.226 20H4.542c-.98 0-1.774-.814-1.774-1.818l-.001-1.686H1.665a.674.674 0 0 1-.665-.68c0-.377.298-.682.665-.682h1.102v-2.419H1.665A.674.674 0 0 1 1 12.033c0-.377.298-.682.665-.682l1.102-.001V8.919H1.665A.674.674 0 0 1 1 8.239c0-.377.298-.682.665-.682l1.102-.001V5.122H1.665A.674.674 0 0 1 1 4.441c0-.377.298-.682.665-.682l1.102-.001l.001-1.94C2.768.814 3.562 0 4.542 0zm-3.248 1.364H4.466a.344.344 0 0 0-.246.118a.428.428 0 0 0-.12.268v2.008h.844a.668.668 0 0 1 .665.683a.674.674 0 0 1-.665.681H4.1v2.431h.873l.045.007l-.074-.004a.65.65 0 0 1 .313.08l.02.011a.53.53 0 0 1 .124.101l-.055-.053a.684.684 0 0 1 .261.509l-.007-.08a.676.676 0 0 1-.596.792l-.03.002l-.016.001H4.1v2.431h.844a.65.65 0 0 1 .308.078c.062.03.111.066.15.111l-.03-.029a.687.687 0 0 1 .216.696l-.009.03a.682.682 0 0 1-.286.378l-.01.005a.644.644 0 0 1-.339.095H4.1v2.419h.873c.008 0 .016.002.023.004l-.052-.004c.367 0 .665.305.665.682c0 .222-.104.42-.265.544l-.013.01a.524.524 0 0 1-.062.04l-.008.004a.628.628 0 0 1-.275.082h-.013l-.015.002h-.014l-.844-.001v1.764c.006.067.03.13.072.19l.048.058c.073.077.163.12.27.13h9.488zm3.264-.002h-1.938v17.274h1.974c.091 0 .176-.042.256-.13a.473.473 0 0 0 .134-.267V1.794a.486.486 0 0 0-.134-.298a.415.415 0 0 0-.292-.134"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotChart(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.02 13.997a1 1 0 0 1 1 1V17a1 1 0 0 1-1 1.001H1a1 1 0 0 1-1-1v-2.002a1 1 0 0 1 1-1.001zM19 8.007a1 1 0 0 1 1 1.002v2.001a1 1 0 0 1-1 1H1a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1.001zM11.03 2a1 1 0 0 1 1 1v2.002a1 1 0 0 1-1 1.001H1.001a1 1 0 0 1-1-1V3A1 1 0 0 1 1 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Down(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.103 12.778L16.81 6.08a.69.69 0 0 1 .99.012a.726.726 0 0 1-.012 1.012l-7.203 7.193a.69.69 0 0 1-.985-.006L2.205 6.72a.727.727 0 0 1 0-1.01a.69.69 0 0 1 .99 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownCircle(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0M6.586 7.636a.682.682 0 1 0-.962.967l3.895 3.87c.265.263.69.265.957.005l3.971-3.87a.682.682 0 1 0-.951-.977l-3.491 3.402Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownCircleO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21M6.586 7.636l3.419 3.397l3.49-3.402a.682.682 0 0 1 .952.977l-3.971 3.87a.682.682 0 0 1-.957-.004l-3.895-3.87a.682.682 0 1 1 .962-.968"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownSquare(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.182 0C19.186 0 20 .814 20 1.818v16.364A1.818 1.818 0 0 1 18.182 20H1.818A1.818 1.818 0 0 1 0 18.182V1.818C0 .814.814 0 1.818 0zM14.46 7.625a.682.682 0 0 0-.964-.013l-3.491 3.402l-3.42-3.397a.682.682 0 0 0-.96.967l3.894 3.87c.265.263.69.265.957.005l3.971-3.87a.682.682 0 0 0 .013-.964"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownSquareO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.818 1.364a.455.455 0 0 0-.454.454v16.364c0 .25.203.454.454.454h16.364a.455.455 0 0 0 .454-.454V1.818a.455.455 0 0 0-.454-.454zM18.182 0C19.186 0 20 .814 20 1.818v16.364A1.818 1.818 0 0 1 18.182 20H1.818A1.818 1.818 0 0 1 0 18.182V1.818C0 .814.814 0 1.818 0zM14.46 7.625a.682.682 0 0 0-.964-.013l-3.491 3.402l-3.42-3.397a.682.682 0 0 0-.96.967l3.894 3.87c.265.263.69.265.957.005l3.971-3.87a.682.682 0 0 0 .013-.964"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.31 12.051c.381 0 .69.314.69.7v4.918c-.006.67-.127 1.2-.399 1.594c-.328.476-.908.692-1.747.737l-15.903-.002c-.646-.046-1.168-.302-1.507-.777c-.302-.423-.446-.95-.444-1.558V12.75c0-.386.309-.7.69-.7c.38 0 .688.314.688.7v4.913c0 .333.065.572.182.736c.081.114.224.184.44.201l15.817.001c.42-.023.627-.1.655-.14c.084-.123.146-.393.15-.8V12.75c0-.386.308-.7.689-.7M9.99 0c.38 0 .69.313.69.7l-.001 10.869l3.062-3.079a.682.682 0 0 1 .975.004a.707.707 0 0 1-.004.99l-4.356 4.38a.682.682 0 0 1-.973-.003l-4.296-4.38a.707.707 0 0 1 .002-.99a.682.682 0 0 1 .975.002L9.3 11.794V.699C9.3.313 9.61 0 9.99 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Edit(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.292 13.36l4.523 4.756L.5 20zM12.705 2.412l4.522 4.755L7.266 17.64l-4.523-4.754zM16.142.348l2.976 3.129c.807.848.086 1.613.086 1.613l-1.521 1.6l-4.524-4.757L14.68.334l.02-.019c.119-.112.776-.668 1.443.033"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Enter(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.3 0a.7.7 0 0 1 .7.7v8.278a6.7 6.7 0 0 1-6.699 6.698l-10.996-.001l3.131 3.13a.7.7 0 0 1-.99.99l-4.24-4.241a.7.7 0 0 1 0-.99l4.241-4.241a.7.7 0 1 1 .99.99l-2.965 2.963h10.83A5.299 5.299 0 0 0 18.6 8.978V.7a.7.7 0 0 1 .7-.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Environment(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.37.805c1.14.574 2.172 1.551 3.126 2.954a7.134 7.134 0 0 1 .996 3.343c.061 1.225-.233 2.517-.873 3.874l-.057.1l-6.023 8.642a.645.645 0 0 1-1.067 0l-5.858-8.4a7.453 7.453 0 0 1-.972-2.314c-.387-1.726 0-4.106 1.368-5.974C5.346 1.206 7.21.315 9.094.069c1.438-.189 2.808-.004 4.277.736m-3.339 1.822c-2.681 0-4.838 2.16-4.838 4.804c0 2.643 2.157 4.804 4.838 4.804s4.838-2.16 4.838-4.804c0-2.643-2.157-4.804-4.838-4.804m0 2.58c1.241 0 2.24 1 2.24 2.224a2.232 2.232 0 0 1-2.24 2.224c-1.24 0-2.24-1-2.24-2.224c0-1.224 1-2.224 2.24-2.224"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvironmentO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.37.805c1.14.574 2.172 1.551 3.126 2.954a7.134 7.134 0 0 1 .996 3.343c.061 1.225-.233 2.517-.873 3.874l-.057.1l-6.023 8.642a.645.645 0 0 1-1.067 0l-5.858-8.4a7.453 7.453 0 0 1-.972-2.314c-.387-1.726 0-4.106 1.368-5.974C5.346 1.206 7.21.315 9.094.069c1.438-.189 2.808-.004 4.277.736m-4.111.61c-1.564.204-3.109.943-4.2 2.433c-1.12 1.529-1.44 3.497-1.137 4.85c.14.628.406 1.256.778 1.856l5.305 7.606l5.456-7.829c.531-1.142.767-2.194.718-3.159c-.05-.99-.316-1.886-.78-2.667c-.812-1.192-1.682-2.015-2.605-2.48c-1.223-.617-2.338-.767-3.535-.61m.772 2.435c1.959 0 3.554 1.599 3.554 3.58c0 1.983-1.595 3.582-3.554 3.582c-1.958 0-3.554-1.599-3.554-3.581s1.596-3.58 3.554-3.58m0 1.357c-1.24 0-2.24 1-2.24 2.224c0 1.223 1 2.224 2.24 2.224c1.241 0 2.24-1 2.24-2.224a2.232 2.232 0 0 0-2.24-2.224"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.963 2.5c2.043 0 4.223.775 6.184 2.434c1.449 1.226 2.703 2.802 3.763 4.722c.11.198.12.439.027.645c-.988 2.2-2.295 3.882-3.921 5.032C13.94 16.8 11.749 17.5 9.999 17.5c-1.973 0-3.734-.525-5.74-2.094c-1.577-1.232-2.964-2.901-4.164-5a.724.724 0 0 1-.021-.678c.734-1.493 1.851-2.95 3.347-4.377C5.438 3.428 7.833 2.5 9.963 2.5m.036 3.406c-2.148 0-3.89 1.792-3.89 4.003c0 2.21 1.742 4.002 3.89 4.002c2.149 0 3.89-1.792 3.89-4.002S12.148 5.906 10 5.906m0 1.413c1.39 0 2.517 1.16 2.517 2.59s-1.127 2.59-2.517 2.59s-2.517-1.16-2.517-2.59S8.61 7.319 10 7.319"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.09 14.781c1.749 1.368 3.219 1.806 4.91 1.806c1.471 0 3.391-.613 5.238-1.919c1.332-.942 2.433-2.315 3.3-4.13c-.94-1.632-2.028-2.968-3.263-4.013c-1.71-1.448-3.582-2.112-5.312-2.112c-1.79 0-3.85.798-5.608 2.474c-1.266 1.206-2.225 2.42-2.88 3.638c1.065 1.789 2.27 3.206 3.614 4.256M10 18c-1.974 0-3.735-.525-5.741-2.094c-1.577-1.232-2.964-2.901-4.164-5a.724.724 0 0 1-.021-.678c.734-1.493 1.851-2.95 3.347-4.377C5.438 3.928 7.833 3 9.963 3c2.043 0 4.223.775 6.184 2.434c1.449 1.226 2.703 2.802 3.763 4.722c.11.198.12.439.027.645c-.988 2.2-2.295 3.882-3.921 5.032C13.94 17.3 11.749 18 9.999 18m.234-3.6a3.7 3.7 0 1 1 0-7.4a3.7 3.7 0 0 1 0 7.4m0-1.4a2.3 2.3 0 1 0 0-4.6a2.3 2.3 0 0 0 0 4.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Facebook(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M18.896 0H1.104C.494 0 0 .494 0 1.104v17.792C0 19.506.494 20 1.104 20h9.578v-7.745H8.076V9.237h2.606V7.01c0-2.584 1.578-3.99 3.883-3.99c1.104 0 2.052.082 2.329.119v2.7h-1.598c-1.254 0-1.496.596-1.496 1.47v1.927h2.989l-.39 3.018h-2.6V20h5.097c.61 0 1.104-.494 1.104-1.104V1.104C20 .494 19.506 0 18.896 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileAdd(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.839 15h-2.865v-3.046a.7.7 0 0 0-.193-.498a.619.619 0 0 0-.468-.205a.619.619 0 0 0-.469.205a.7.7 0 0 0-.193.498v3.047H9.933a.619.619 0 0 0-.468.205a.7.7 0 0 0-.193.498a.7.7 0 0 0 .193.498a.619.619 0 0 0 .468.205h2.718v2.89a.7.7 0 0 0 .193.498a.619.619 0 0 0 .469.205c.183 0 .34-.068.468-.205a.7.7 0 0 0 .193-.498v-2.89h2.865a.62.62 0 0 0 .468-.205a.7.7 0 0 0 .193-.498a.7.7 0 0 0-.193-.498a.619.619 0 0 0-.468-.205m.516-9.973L13.33.481l-.44-.479H3.78c-.17 0-.36-.01-.526.017A.827.827 0 0 0 2.742.3c-.281.318-.24.717-.24 1.108v17.185c0 .391-.045.814.24 1.11c.285.296.476.296.843.296H8.61a.62.62 0 0 0 .469-.205a.7.7 0 0 0 .193-.498a.7.7 0 0 0-.193-.498a.619.619 0 0 0-.469-.205H3.845L3.839 1.364h7.275v3.593c0 .391.152.743.391.996c.24.253.57.41.937.41h3.735v4.185a.7.7 0 0 0 .193.498a.619.619 0 0 0 .468.205c.184 0 .34-.069.469-.205a.7.7 0 0 0 .192-.498V5.475l.001-.279zm-4.837-.172c-.041-.053-.066-.146-.075-.28V1.52l3.059 3.437H12.76c-.12-.016-.2-.05-.242-.103"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileExcel(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.534 1.36L14.309 0H4.662c-.696 0-.965.516-.965.919v3.63H5.05V1.653c0-.154.13-.284.28-.284h6.903c.152 0 .228.027.228.152v4.82h4.913c.193 0 .268.1.268.246v11.77c0 .246-.1.283-.25.283H5.33a.287.287 0 0 1-.28-.284V17.28H3.706v1.695c-.018.6.302 1.025.956 1.025H18.06c.7 0 .939-.507.939-.969V5.187l-.35-.38zm-1.698.16l.387.434l2.596 2.853l.143.173h-2.653c-.2 0-.327-.033-.38-.1c-.053-.065-.084-.17-.093-.313zm-1.09 9.147h4.577v1.334h-4.578zm0-2.666h4.577v1.333h-4.578zm0 5.333h4.577v1.334h-4.578zM1 5.626v10.667h10.465V5.626zm5.233 6.204l-.64.978h.64V14H3.016l2.334-3.51l-2.068-3.156H5.01L6.234 9.17l1.223-1.836h1.727L7.112 10.49L9.449 14H7.656z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileJpg(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.928 9.097a2.07 2.07 0 0 0-.526.053v1.728c.109.027.242.035.426.035c.675 0 1.093-.356 1.093-.955c0-.54-.36-.86-.993-.86m8.996-1.943h-.514l.027-1.89a.464.464 0 0 0-.12-.298L12.901.134A.393.393 0 0 0 12.618 0h-9.24a.8.8 0 0 0-.787.784v6.37h-.515c-.285 0-.56.118-.76.328A1.14 1.14 0 0 0 1 8.275v5.83c0 .618.482 1.12 1.076 1.12h.515v3.99A.8.8 0 0 0 3.38 20h13.278c.415 0 .78-.352.78-.784v-3.99h.487c.594 0 1.076-.503 1.076-1.122v-5.83c0-.296-.113-.582-.315-.792a1.054 1.054 0 0 0-.76-.328M3.95 1.378h6.956v4.577a.4.4 0 0 0 .11.277a.37.37 0 0 0 .267.115h4.759v.807H3.95zm7.23 8.545c0 .573-.184 1.06-.517 1.39c-.434.426-1.076.617-1.827.617c-.167 0-.317-.009-.434-.026v2.094h-1.17V8.22c.393-.07.854-.122 1.63-.122c.783 0 1.342.157 1.717.47c.359.295.6.781.6 1.355m-7.792 2.963c.175.06.4.104.65.104c.534 0 .867-.252.867-1.164V8.142h1.268v3.7c0 1.67-.768 2.251-2.002 2.251a3.17 3.17 0 0 1-.925-.139zm.562 5.736v-3.397h12.092v3.397zm12.734-4.894a6.061 6.061 0 0 1-1.877.33c-1.025 0-1.767-.269-2.284-.79c-.517-.504-.801-1.268-.793-2.129c.009-1.946 1.368-3.058 3.211-3.058c.726 0 1.285.147 1.56.286l-.267 1.06c-.309-.139-.693-.252-1.31-.252c-1.059 0-1.86.626-1.86 1.895c0 1.207.726 1.92 1.769 1.92c.292 0 .525-.035.625-.087v-1.225h-.867v-1.034h2.093zM12.291 1.52l.385.434l2.58 2.853l.143.173h-2.637c-.2 0-.325-.033-.378-.1c-.053-.065-.084-.17-.093-.313z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePdf(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.924 7.154h-.514l.027-1.89a.464.464 0 0 0-.12-.298L12.901.134A.393.393 0 0 0 12.618 0h-9.24a.8.8 0 0 0-.787.784v6.37h-.515c-.285 0-.56.118-.76.328A1.14 1.14 0 0 0 1 8.275v5.83c0 .618.482 1.12 1.076 1.12h.515v3.99A.8.8 0 0 0 3.38 20h13.278c.415 0 .78-.352.78-.784v-3.99h.487c.594 0 1.076-.503 1.076-1.122v-5.83c0-.296-.113-.582-.315-.792a1.054 1.054 0 0 0-.76-.328M3.95 1.378h6.956v4.577a.4.4 0 0 0 .11.277a.37.37 0 0 0 .267.115h4.759v.807H3.95zm0 17.244v-3.397h12.092v3.397zM12.291 1.52l.385.434l2.58 2.853l.143.173h-2.637c-.2 0-.325-.033-.378-.1c-.053-.065-.084-.17-.093-.313zM3 14.232v-6h1.918c.726 0 1.2.03 1.42.09c.34.09.624.286.853.588c.228.301.343.69.343 1.168c0 .368-.066.678-.198.93c-.132.25-.3.447-.503.59a1.72 1.72 0 0 1-.62.285c-.285.057-.698.086-1.239.086h-.779v2.263zm1.195-4.985v1.703h.654c.471 0 .786-.032.945-.094a.786.786 0 0 0 .508-.762a.781.781 0 0 0-.19-.54a.823.823 0 0 0-.48-.266c-.142-.027-.429-.04-.86-.04zm4.04-1.015h2.184c.493 0 .868.038 1.127.115c.347.103.644.288.892.552c.247.265.436.589.565.972c.13.384.194.856.194 1.418c0 .494-.06.92-.182 1.277c-.148.437-.36.79-.634 1.06c-.207.205-.487.365-.84.48c-.263.084-.616.126-1.057.126H8.235zM9.43 9.247v3.974h.892c.334 0 .575-.019.723-.057c.194-.05.355-.132.482-.25c.128-.117.233-.31.313-.579c.081-.269.121-.635.121-1.099c0-.464-.04-.82-.12-1.068a1.377 1.377 0 0 0-.34-.581a1.132 1.132 0 0 0-.553-.283c-.167-.038-.494-.057-.98-.057zm4.513 4.985v-6H18v1.015h-2.862v1.42h2.47v1.015h-2.47v2.55z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileText(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.89 0l4.61 5.192V19.09a.896.896 0 0 1-.882.909H3.382a.896.896 0 0 1-.882-.91V.91c0-.503.395-.91.882-.91zm-.584 1.364H3.824v17.272h12.352V6.342h-2.829c-.705.028-1.25-.086-1.64-.394c-.454-.358-.62-.968-.598-1.781v-2.8h1.201zm-1.382 13.181c.365 0 .662.306.662.682a.672.672 0 0 1-.662.682H5.809a.672.672 0 0 1-.662-.682c0-.376.296-.682.662-.682zm3.055-3.636c.365 0 .662.305.662.682a.672.672 0 0 1-.662.682h-8.17a.672.672 0 0 1-.662-.682c0-.377.296-.682.662-.682zm-3.055-3.636c.365 0 .662.305.662.682a.672.672 0 0 1-.662.681H5.809a.672.672 0 0 1-.662-.681c0-.377.296-.682.662-.682zm1.46-5.823l.048.059v2.676c-.012.439.047.656.08.68c.095.076.358.131.81.114h2.193l.375.424l.286.319z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileUnknown(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.89 0l4.61 5.192V19.09a.896.896 0 0 1-.882.909H3.382a.896.896 0 0 1-.882-.91V.91c0-.503.395-.91.882-.91zm-.584 1.364H3.824v17.272h12.352V6.342h-2.829c-.705.028-1.25-.086-1.64-.394c-.454-.358-.62-.968-.598-1.781v-2.8h1.201zm-3.29 13.615c.389 0 .705.313.705.7a.703.703 0 0 1-.706.699a.703.703 0 0 1-.706-.7c0-.386.316-.699.706-.699m2.167-6.497c.577.578.773 1.277.699 1.998c-.055.54-.404 1.071-.887 1.57l-.049.05l-.105.106c-.077.077-.17.168-.307.3c-.145.14-.245.238-.322.315l-.1.1l-.082.086c-.204.218-.267.31-.301.411c-.034.101-.08.321-.127.641a.665.665 0 0 1-.754.571a.68.68 0 0 1-.554-.777c.059-.396.117-.683.184-.882c.11-.326.274-.562.599-.91c.124-.133.242-.249.547-.545c.23-.223.334-.325.433-.428c.283-.292.495-.615.509-.751a1.01 1.01 0 0 0-.306-.877c-.345-.346-.768-.508-1.195-.475c-.49.038-.878.22-1.196.58a2.093 2.093 0 0 0-.472.941a.661.661 0 0 1-.795.508a.683.683 0 0 1-.493-.82a3.47 3.47 0 0 1 .782-1.548c.554-.627 1.26-.957 2.074-1.02c.81-.064 1.604.241 2.218.856m1.201-7.032l.048.059v2.676c-.012.439.047.656.08.68c.095.076.358.131.81.114h2.193l.375.424l.286.319z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filter(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.398 14.605l1.323 1.143c.29.251.323.691.075.984a.688.688 0 0 1-.976.075l-1.565-1.352a.7.7 0 0 1-.242-.53V7.938L1.171 1.155C.78.703 1.1 0 1.694 0h16.612c.594 0 .912.704.523 1.155l-5.85 6.784v11.363c0 .386-.31.698-.692.698a.695.695 0 0 1-.692-.698V7.678a.7.7 0 0 1 .17-.458l5.023-5.825H3.21L8.228 7.22a.7.7 0 0 1 .17.458z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Firefox(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.971 6.676l-.231 1.422s-.332-2.633-.737-3.617c-.622-1.508-.899-1.496-.9-1.494c.416 1.013.34 1.558.34 1.558s-.737-1.925-2.689-2.538c-2.162-.678-3.331-.492-3.466-.458h-.06l.048.004l-.002.001c.009.01 2.389.398 2.81.954c0 0-1.01 0-2.016.277c-.046.013 3.701.448 4.467 4.031c0 0-.41-.82-.919-.959c.334.973.249 2.818-.07 3.735c-.04.118-.082-.51-.71-.78c.202 1.376-.011 3.56-1.01 4.161c-.077.047.626-2.155.143-1.304c-2.788 4.09-6.082 1.887-7.564.918c.76.158 2.2-.024 2.838-.479l.002-.001c.693-.453 1.103-.785 1.472-.706c.368.079.614-.275.327-.59c-.287-.314-.983-.747-1.924-.51c-.664.166-1.487.869-2.743.157c-.964-.547-1.055-1.001-1.064-1.316c.024-.11.054-.215.09-.31c.11-.297.447-.387.634-.457c.318.052.591.147.878.288a4.641 4.641 0 0 0 0-.35c.028-.053.01-.21-.034-.404a2.827 2.827 0 0 0-.132-.574l.004-.002a.011.011 0 0 0 .004-.003v-.001a.023.023 0 0 0 .005-.009c.02-.086.235-.253.502-.432c.24-.16.522-.33.744-.463c.196-.116.346-.203.378-.226l.042-.03l.009-.007l.006-.004a.744.744 0 0 0 .296-.553v-.002l.003-.029l.001-.02l.001-.017l.001-.038v-.003a1.349 1.349 0 0 0-.008-.149v-.002l-.001-.005l-.003-.008l-.003-.01c-.035-.077-.163-.105-.693-.114h-.001a49.458 49.458 0 0 0-.87-.003c-.649.003-1.008-.607-1.122-.843c.157-.831.61-1.423 1.356-1.825c.014-.008.011-.014-.005-.018c.146-.085-1.763-.003-2.64 1.066c-.78-.186-1.458-.173-2.043-.042a2.519 2.519 0 0 1-.42-.049c-.388-.338-.946-.96-.975-1.705l-.005.004l-.002-.022s-1.186.873-1.008 3.25l-.002.11c-.321.417-.48.767-.492.844C.523 6.53.235 7.363 0 8.63c0 0 .164-.498.494-1.063c-.242.711-.433 1.816-.321 3.474c0 0 .03-.367.134-.897c.082 1.028.44 2.298 1.345 3.79c1.737 2.866 4.407 4.313 7.359 4.535a10.5 10.5 0 0 0 1.59.004l.148-.011c.605-.04 1.214-.128 1.821-.269c8.304-1.92 7.401-11.517 7.401-11.517"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.566 5.838a1.357 1.357 0 0 1-1.347-1.135L7.984 3.22a.452.452 0 0 0-.45-.378H1.818a.45.45 0 0 0-.454.447v13.422a.45.45 0 0 0 .454.447h16.364c.25 0 .454-.2.454-.447V6.285a.451.451 0 0 0-.454-.447zm0-1.342h8.616c1.004 0 1.818.8 1.818 1.79V16.71c0 .988-.814 1.789-1.818 1.789H1.818C.814 18.5 0 17.699 0 16.71V3.29C0 2.3.814 1.5 1.818 1.5h5.716a1.81 1.81 0 0 1 1.797 1.514z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderAdd(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.511 9.409c.377 0 .682.299.682.668v2.964h3.032c.378 0 .684.3.684.668a.676.676 0 0 1-.684.667l-3.032-.001v2.965c0 .37-.305.669-.682.669a.675.675 0 0 1-.682-.669l-.001-2.965h-3.031a.676.676 0 0 1-.684-.666c0-.369.306-.667.684-.667l3.031-.001v-2.964c0-.369.306-.668.683-.668M1.417 2l6.116.002c.43.028.783.14 1.047.37c.226.198.392.462.518.806l.38 1.17a.656.656 0 0 1 .028.126h9.81c.378 0 .684.299.684.667v3.58a.676.676 0 0 1-.684.667a.676.676 0 0 1-.683-.667V5.808H1.367v10.368c.004.21.002.303.05.373c.037.056.067.067.15.07h8.957c.377 0 .683.298.683.667a.676.676 0 0 1-.683.667H1.577l-.117-.01c-.536-.09-.94-.29-1.177-.654c-.19-.293-.276-.658-.283-1.103V5.137l.003-1.822c.036-.355.155-.66.376-.895c.244-.261.585-.39 1.038-.42m.047 1.333c-.084.005-.09.024-.096.047l-.001 1.094h6.717l-.282-.87c-.046-.126-.084-.199-.134-.237c-.049-.038-.074-.024-.182-.033z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpen(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.282 3.828c-.35.063-.566.189-.691.374a1.33 1.33 0 0 0-.225.841v11.718L.002 15.82V5.072A2.6 2.6 0 0 1 .45 3.474c.357-.53.929-.85 1.666-.966l.108-.008h3.56c.742.038 1.355.35 1.783.917c.387.51.472.91.484 1.639c.002.088.003.15.006.201c.046 0 .109-.003.193-.01h5.636l.101.008a2.6 2.6 0 0 1 1.489.724c.435.431.632 1.028.614 1.734v1.799h-1.364V7.698c.01-.39-.073-.639-.224-.789a1.21 1.21 0 0 0-.671-.333l-5.516-.003c-.854.08-1.504-.252-1.601-1.018a3.953 3.953 0 0 1-.027-.478c-.008-.475-.043-.642-.218-.873a.908.908 0 0 0-.72-.376zm15.772 5.504c.831.014 1.46.304 1.77.908c.294.574.21 1.226-.18 1.91l-.054.08l-2.843 3.626c-.34.448-.73.818-1.172 1.106c-.483.316-1.088.49-1.851.538H2.016c-.633 0-1.13-.141-1.472-.477c-.42-.414-.589-.85-.531-1.394c.035-.338.169-.668.415-1.022l3.445-3.833c.538-.558 1.02-.94 1.465-1.147c.404-.187.888-.282 1.444-.294zM7.012 10.66c-.476-.009-.839.051-1.088.166c-.265.123-.624.409-1.013.807l-3.376 3.749a.858.858 0 0 0-.165.383c-.015.139.003.186.143.323c.04.039.2.084.503.084l11.663.002c.494-.032.87-.14 1.136-.314c.308-.201.585-.464.84-.801l2.82-3.594c.17-.311.191-.508.128-.631c-.05-.098-.202-.168-.56-.174z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Forward(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.097 8.038v-5.19c.038-.343.227-.563.568-.662c.341-.1.678-.013 1.012.26l8.086 6.838a.924.924 0 0 1 .334.716c0 .28-.111.52-.334.718l-8.14 6.884c-.333.23-.66.299-.98.205c-.32-.093-.502-.298-.546-.613v-5.238l-6.614 5.598c-.33.264-.68.349-1.046.253c-.366-.095-.544-.3-.534-.613V2.93c-.01-.329.14-.566.451-.712c.311-.146.668-.089 1.07.172z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Foursquare(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.786 19.744c-.432.506-1.286.212-1.286-.443l.001-17.784c.031-.541.233-.976.652-1.232C3.487.081 3.93 0 4.51 0h11.586c.541-.002.99.188 1.23.615c.188.333.217.713.119 1.086l-.648 2.913c-.163.554-.41.959-.789 1.197c-.353.222-.824.304-1.413.29H9.21c-.194 0-.29.023-.288.02c.046-.034.048-.038.046.028v1.24c.004.11.006.116-.067.065c-.017-.012.16.018.468.018h4.75c.606-.016 1.111.155 1.449.548c.32.373.426.85.337 1.422l-.605 2.823c-.124.498-.362.877-.752 1.094c-.303.168-.676.255-1.157.277l-3.852-.001c-.217-.013-.367.002-.444.03a.986.986 0 0 0-.292.213zm.165-2.39l3.765-4.4c.281-.284.565-.486.867-.595c.294-.107.64-.14 1-.118h3.772c.232-.01.389-.047.467-.09c-.008.004.027-.052.062-.19l.592-2.761c.026-.169.008-.25-.027-.291c-.019-.022-.093-.047-.31-.041H9.37c-.595 0-1.013-.07-1.319-.283c-.362-.252-.515-.65-.533-1.17V6.131c.011-.464.168-.852.512-1.11c.3-.226.693-.316 1.181-.316h5.401c.34.008.548-.028.604-.063c.03-.02.11-.15.174-.365l.64-2.883H4.51c-.302 0-.496.031-.563.06c.017-.005.01.022.004.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Frown(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m0 11.399a3.637 3.637 0 0 0-3.61 3.2a.682.682 0 0 0 1.354.162a2.273 2.273 0 0 1 4.51-.015a.682.682 0 1 0 1.353-.172A3.637 3.637 0 0 0 10 11.4M5.814 6.28a1.395 1.395 0 1 0 0 2.79a1.395 1.395 0 0 0 0-2.79m8.372 0a1.395 1.395 0 1 0 0 2.79a1.395 1.395 0 0 0 0-2.79"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FrownO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21M10 11.4a3.637 3.637 0 0 1 3.607 3.175a.682.682 0 1 1-1.352.172a2.273 2.273 0 0 0-4.511.015a.682.682 0 1 1-1.354-.163a3.637 3.637 0 0 1 3.61-3.2m-4.186-5.12a1.395 1.395 0 1 1 0 2.79a1.395 1.395 0 0 1 0-2.79m8.372 0a1.395 1.395 0 1 1 0 2.79a1.395 1.395 0 0 1 0-2.79"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Github(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 10.25c0 2.234-.636 4.243-1.908 6.027c-1.271 1.784-2.914 3.018-4.928 3.703c-.234.045-.406.014-.514-.093a.539.539 0 0 1-.163-.4V16.67c0-.863-.226-1.495-.677-1.895a8.72 8.72 0 0 0 1.335-.24c.394-.107.802-.28 1.223-.52a3.66 3.66 0 0 0 1.055-.888c.282-.352.512-.819.69-1.402c.178-.583.267-1.252.267-2.008c0-1.077-.343-1.994-1.028-2.75c.32-.81.286-1.717-.105-2.723c-.243-.08-.594-.03-1.054.147a6.94 6.94 0 0 0-1.198.587l-.495.32a9.03 9.03 0 0 0-2.5-.346a9.03 9.03 0 0 0-2.5.347a11.52 11.52 0 0 0-.553-.36c-.23-.143-.593-.314-1.088-.514c-.494-.2-.868-.26-1.12-.18c-.381 1.005-.412 1.912-.09 2.722c-.686.756-1.03 1.673-1.03 2.75c0 .756.09 1.423.268 2.002c.178.578.406 1.045.683 1.401a3.53 3.53 0 0 0 1.048.894c.421.24.83.414 1.224.52c.395.108.84.188 1.335.241c-.347.32-.56.779-.638 1.375a2.539 2.539 0 0 1-.586.2a3.597 3.597 0 0 1-.742.067c-.287 0-.57-.096-.853-.287c-.282-.192-.523-.47-.723-.834a2.133 2.133 0 0 0-.631-.694c-.256-.178-.471-.285-.645-.32l-.26-.04c-.182 0-.308.02-.378.06c-.07.04-.09.09-.065.153a.738.738 0 0 0 .117.187a.961.961 0 0 0 .17.16l.09.066c.192.09.38.259.567.508c.187.249.324.476.41.68l.13.307c.113.338.304.612.574.821c.269.21.56.343.872.4c.312.058.614.09.905.094c.29.004.532-.011.723-.047l.299-.053c0 .338.002.734.007 1.188l.006.72c0 .16-.056.294-.17.4c-.112.108-.286.139-.52.094c-2.014-.685-3.657-1.92-4.928-3.703C.636 14.493 0 12.484 0 10.25c0-1.86.447-3.574 1.341-5.145a10.083 10.083 0 0 1 3.64-3.73A9.6 9.6 0 0 1 10 0a9.6 9.6 0 0 1 5.02 1.375a10.083 10.083 0 0 1 3.639 3.73C19.553 6.675 20 8.391 20 10.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GithubO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.773 16.16c-.82-1.06-1.063-1.406-1.25-1.745l-.039-.073a10.42 10.42 0 0 1-.122-.226a3.352 3.352 0 0 1-.14-.3c-.174-.432-.207-.813.07-1.165c.485-.62 1.273-.445 1.937.123c.428.365 1.066 1.02 1.455 1.486c.255.305.72.621 1.06.734a.83.83 0 0 0 .1.027l.178-.56a7.352 7.352 0 0 1-1.869-.965c-1.3-.936-2.43-2.773-2.333-4.174c.06-.885.36-1.753.89-2.596a12.72 12.72 0 0 1-.073-1.238c0-.489.16-1.02.467-1.603a.661.661 0 0 1 .566-.352a5.853 5.853 0 0 1 1.598.17c.455.115.936.317 1.445.605c.762-.158 1.5-.238 2.216-.238c.717 0 1.463.08 2.239.239c.662-.371 1.244-.623 1.752-.755c.66-.171 1.254-.12 1.743.182a.659.659 0 0 1 .175.158c.3.388.466.903.52 1.538c.042.484.027.965-.045 1.442c.433.651.715 1.483.854 2.49c.226 1.644-.525 3.337-1.718 4.312c-.556.454-1.32.81-2.294 1.08c.062.216.106.431.132.646c.044.354.075 1.304.096 2.89c1.102-.574 1.957-1.14 2.565-1.695c1.035-.943 1.815-2.281 2.323-3.835c.425-1.3.507-2.397.297-3.606c-.257-1.47-.627-2.669-1.184-3.631c-.652-1.126-1.692-2.19-2.936-2.968c-1.19-.746-3.12-1.245-4.64-1.245c-1.596 0-3.382.517-4.737 1.473c-1.318.93-2.41 2.33-3.043 3.703c-.501 1.088-.707 2.247-.707 3.958c0 1.386.504 3.07 1.098 4.124c.469.832 1.135 1.62 2.114 2.476c.493.432 1.19.863 2.09 1.289v-.423c-1.205-.28-2.163-.864-2.85-1.752m-.357-2.922a.197.197 0 0 0 0-.005zm1.405 2.122c.57.739 1.411 1.192 2.56 1.36a.658.658 0 0 1 .563.65v1.973a.66.66 0 0 1-.913.606c-1.475-.606-2.597-1.24-3.371-1.917c-1.086-.95-1.847-1.852-2.394-2.82C.568 13.971 0 12.072 0 10.445C0 8.56.238 7.22.827 5.942c.724-1.572 1.959-3.154 3.48-4.227C5.894.595 7.947 0 9.806 0c1.76 0 3.933.562 5.345 1.446c1.424.892 2.616 2.111 3.376 3.424c.649 1.119 1.061 2.453 1.341 4.062c.248 1.421.15 2.734-.341 4.235c-.574 1.753-1.467 3.285-2.686 4.397c-.882.804-2.156 1.591-3.827 2.373a.66.66 0 0 1-.942-.588c-.02-2.149-.053-3.434-.097-3.788a3.274 3.274 0 0 0-.297-.989a.655.655 0 0 1 .46-.919c1.144-.243 1.968-.58 2.47-.99c.855-.699 1.41-1.95 1.249-3.12c-.128-.926-.387-1.626-.764-2.108a.653.653 0 0 1-.127-.53c.09-.45.115-.903.075-1.362c-.028-.326-.092-.572-.183-.742c-.145-.048-.341-.046-.606.023c-.43.111-.98.36-1.64.745a.664.664 0 0 1-.482.075a10.036 10.036 0 0 0-2.202-.261c-.693 0-1.42.086-2.179.26a.664.664 0 0 1-.492-.079c-.49-.296-.93-.492-1.314-.588c-.27-.068-.55-.11-.837-.126c-.101.251-.15.464-.15.638c0 .342.03.786.09 1.328a.653.653 0 0 1-.108.437c-.488.724-.752 1.441-.801 2.159c-.062.885.826 2.327 1.79 3.021c.595.429 1.292.75 2.093.962a.656.656 0 0 1 .46.831l-.503 1.587a.658.658 0 0 1-.394.416c-.401.152-.83.152-1.26.01c-.572-.19-1.244-.647-1.657-1.141a10.733 10.733 0 0 0-.354-.4c.136.18.302.396.508.662"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Global(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21m2.048 2.093a6.702 6.702 0 0 1 1.802 1.024c.125.096.245.198.363.307c.022.019.041.039.062.06c.328.303.626.644.89 1.01c.037.051.076.098.11.147c.064.094.124.192.184.29a6.024 6.024 0 0 1 .285.51c.053.1.1.209.149.315c.033.073.07.148.1.226c.073.182.138.365.197.552c.019.062.034.128.052.19a6.726 6.726 0 0 1 .215 1.087c.007.066.017.136.022.201c.02.215.033.432.033.65a7.186 7.186 0 0 1-.117 1.27c-.014.073-.026.147-.04.217c-.04.19-.086.376-.14.56c-.509-.233-1.107-.576-1.263-.953c-.284-.68-1.04-1.02-1.348-1.896c-.507-1.45.166-1.412.26-2.312c.044-.422-.26-.51-.661-.338c-.936.393-1.253.242-1.442-.463c-.189-.703 0-.899 0-.899c-.638.07-.662-.707-.331-.903c.23-.132.425-.537.618-.852M9.374 7.797c.59-.27 1.135-.367 1.063-.831c-.07-.459-.236-.801-1.158-.801c-.922 0-.52 1.265-1.276.51c-.756-.75.165-.556.543-.727c.379-.172.757-.877.095-.927c-.661-.047-.52.292-1.04.1c-.52-.196-.756.679-1.088.557c-.218-.082-.803-.532-1.191-.975a7 7 0 0 0-1.834 2.51c.113 1.307.804 1.993.804 1.993s.355.851 2.483 1.897c0 0 .4.024-.072-.461c-.472-.487-.993-1.095-.402-1.41c.59-.319.757-.292.899.293c.141.584.615.24.661-.319c.048-.557.922-1.14 1.513-1.41m-.45 2.94c1.018 0 .923.317 1.727 1.025c.803.704.378 1.409-.025 1.945c-.401.534-.756 1.14-.945 2.238c-.19 1.094-.686.314-.85.047c-.166-.269-.426-.511-.354-1.63c.07-1.118-.687-.46-.946-1.92c-.26-1.458.378-1.704 1.394-1.704m4.977.964c.271-.173.92.278.78.753c-.143.475-.591.207-.816 0c-.225-.206-.237-.583.036-.753"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartOff(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.682 1.682c2.007-.48 3.667.026 4.849 1.094c1.032.934 1.628 2.559 1.432 4.39c-.147 1.374-.893 2.769-2.242 4.22l-7.274 6.923a.7.7 0 0 1-.963-.002L2.266 11.36C1.198 10.227.516 9.15.233 8.118C-.18 6.62-.04 5.107.61 3.875c.77-1.461 2.306-2.329 4.532-2.329c1.522 0 3.12.753 4.805 2.221c1.148-1.066 2.394-1.765 3.735-2.085M9.97 16.883l6.76-6.433c1.138-1.226 1.75-2.37 1.863-3.425c.151-1.415-.29-2.619-.997-3.258c-.859-.777-2.056-1.142-3.588-.776c-1.257.3-2.434 1.02-3.538 2.174a.7.7 0 0 1-.981.025C7.852 3.644 6.397 2.895 5.142 2.895c-1.722 0-2.78.597-3.307 1.598c-.489.926-.597 2.094-.272 3.275c.218.795.786 1.692 1.692 2.655z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartOn(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.67 1.658c2.103-.494 4.165.222 5.135 1.434c.896 1.119 1.303 2.282 1.17 3.998c-.095 1.237-.5 2.154-1.298 3.163c-.147.186-.308.377-.511.609l-.548.618c-.8.903-3.155 3.163-7.103 6.82a.764.764 0 0 1-1.036-.004C6.09 15.12 3.896 13.028 2.89 12.014C1.033 10.14.326 9.05.07 7.293c-.27-1.84.245-3.349 1.398-4.521c1.082-1.1 3.017-1.508 4.928-1.12c1.119.228 2.31.96 3.6 2.181c1.101-1.129 2.327-1.86 3.675-2.175"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.178 11.373a.7.7 0 0 1 .7.7v5.874c.027.812-.071 1.345-.434 1.68c-.338.311-.828.4-1.463.366H3.144C2.5 19.961 2 19.7 1.768 19.173c-.154-.347-.226-.757-.226-1.228v-5.873a.7.7 0 0 1 1.4 0v5.873c0 .232.026.42.07.562l.036.098l-.003-.01c.001-.013.03-.008.132-.002h13.84c.245.014.401 0 .456-.001l.004-.001c-.013-.053.012-.27 0-.622v-5.897a.7.7 0 0 1 .701-.7M10.434 0c.264 0 .5.104.722.297l8.625 8.139a.7.7 0 1 1-.962 1.017l-8.417-7.944l-9.244 7.965a.7.7 0 0 1-.915-1.06L9.689.277l.086-.064c.214-.134.428-.212.66-.212"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HtmlFive(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m1.5 0l1.547 18l6.943 2l6.961-2.002L18.5 0zm13.773 4.272l-.097 1.119l-.043.496H6.99l.195 2.26h7.755l-.052.594l-.5 5.812l-.032.373L10 16.178l-.01.004l-4.36-1.256l-.297-3.467h2.136l.151 1.762l2.37.663h.002l2.374-.665l.247-2.863H5.237l-.523-6.084l-.05-.593h10.66z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ie(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 10.466c0 .398-.022.796-.078 1.182H7.076c0 2.261 1.953 3.897 4.096 3.897c1.45 0 2.846-.727 3.594-2.011h4.72c-1.276 3.648-4.666 6.086-8.47 6.09a8.933 8.933 0 0 1-3.974-.943C5.792 19.33 4.04 20 2.645 20C.77 20 0 18.83 0 17.012c0-1.057.224-2.114.502-3.125c.179-.66.893-2 1.217-2.603C3.09 8.751 4.899 6.32 7.02 4.398c-1.708.75-3.56 2.636-4.766 4.022c.939-4.157 4.573-7.102 8.762-7.102c.167 0 .335 0 .502.011C12.9.681 14.833 0 16.35 0c1.808 0 3.359.705 3.359 2.784c0 1.09-.413 2.273-.837 3.25A9.298 9.298 0 0 1 20 10.466m-.781-7.273c0-1.272-.893-2.056-2.121-2.056c-.938 0-1.998.386-2.835.795a9.105 9.105 0 0 1 4.386 3.716c.28-.75.57-1.659.57-2.455M1.429 17.25c0 1.319.77 2.034 2.042 2.034c.993 0 2.098-.454 2.968-.943a9.12 9.12 0 0 1-3.917-4.863c-.513 1.09-1.093 2.545-1.093 3.772m5.625-8.102h8.125c-.078-2.194-1.976-3.773-4.063-3.773c-2.099 0-3.985 1.58-4.063 3.773z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inbox(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.157 0c.378 0 .842.372 1.035.83L20 7.439V18a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V7.438L2.808.831C3 .372 3.465 0 3.843 0ZM6.741 8.838H1.4V18a.6.6 0 0 0 .6.6h16a.6.6 0 0 0 .6-.6V8.838h-4.341a3.902 3.902 0 0 1-7.518 0M15.913 1.4H4.087L1.52 7.438h6.505a2.5 2.5 0 1 0 4.95 0h5.505z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Information(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m-.145 7.21a.698.698 0 0 0-.698.697v7.558a.698.698 0 0 0 1.395 0V7.907a.698.698 0 0 0-.697-.698m.028-2.791a.93.93 0 1 0 0 1.86a.93.93 0 0 0 0-1.86"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InformationO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21M9.855 7.21c.385 0 .697.313.697.698v7.558a.698.698 0 0 1-1.395 0V7.907c0-.385.312-.698.698-.698m.028-2.79a.93.93 0 1 1 0 1.86a.93.93 0 0 1 0-1.86"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laptop(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.364 2.35v10.794h17.272V2.35zM19.09 1c.502 0 .909.403.909.9v11.694c0 .497-.407.9-.91.9l-8.53-.001v3.157h3.437c.377 0 .682.303.682.675a.678.678 0 0 1-.682.675H5.981a.678.678 0 0 1-.682-.675c0-.372.306-.674.682-.674l3.216-.001v-3.157H.909A.904.904 0 0 1 0 13.595V1.9c0-.497.407-.9.91-.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Left(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.222 9.897c2.3-2.307 4.548-4.559 6.744-6.754a.65.65 0 0 0 0-.896c-.311-.346-.803-.316-1.027-.08c-2.276 2.282-4.657 4.667-7.143 7.156c-.197.162-.296.354-.296.574c0 .221.099.418.296.592l7.483 7.306a.749.749 0 0 0 1.044-.029c.358-.359.22-.713.058-.881a3407.721 3407.721 0 0 1-7.16-6.988"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftCircle(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m.788 5.646l-3.87 3.971a.682.682 0 0 0 .005.957l3.87 3.895a.682.682 0 1 0 .967-.961l-3.397-3.42l3.402-3.49a.682.682 0 1 0-.977-.952"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftCircleO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21m.788 4.25a.682.682 0 0 1 .977.952l-3.402 3.491l3.397 3.42a.682.682 0 0 1-.967.96l-3.87-3.894a.682.682 0 0 1-.005-.957Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftSquare(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.182 0C19.186 0 20 .814 20 1.818v16.364A1.818 1.818 0 0 1 18.182 20H1.818A1.818 1.818 0 0 1 0 18.182V1.818C0 .814.814 0 1.818 0zm-7.394 5.627l-3.87 3.971a.682.682 0 0 0 .005.957l3.87 3.895a.682.682 0 1 0 .967-.962L8.363 10.07l3.402-3.49a.682.682 0 1 0-.977-.952"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftSquareO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.818 1.364a.455.455 0 0 0-.454.454v16.364c0 .25.203.454.454.454h16.364a.455.455 0 0 0 .454-.454V1.818a.455.455 0 0 0-.454-.454zM18.182 0C19.186 0 20 .814 20 1.818v16.364A1.818 1.818 0 0 1 18.182 20H1.818A1.818 1.818 0 0 1 0 18.182V1.818C0 .814.814 0 1.818 0zm-7.394 5.627l-3.87 3.971a.682.682 0 0 0 .005.957l3.87 3.895a.682.682 0 1 0 .967-.962L8.363 10.07l3.402-3.49a.682.682 0 1 0-.977-.952"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LikeO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.36 9.495v7.157h3.03l.153.018c2.813.656 4.677 1.129 5.606 1.422c1.234.389 1.694.484 2.531.54c.626.043 1.337-.198 1.661-.528c.179-.182.313-.556.366-1.136a.681.681 0 0 1 .406-.563c.249-.108.456-.284.629-.54c.16-.234.264-.67.283-1.301a.682.682 0 0 1 .339-.57c.582-.337.87-.717.93-1.163c.066-.493-.094-1.048-.513-1.68a.683.683 0 0 1 .176-.936c.401-.282.621-.674.676-1.23c.088-.886-.477-1.541-1.756-1.672a9.42 9.42 0 0 0-3.394.283a.68.68 0 0 1-.786-.95c.5-1.058.778-1.931.843-2.607c.085-.897-.122-1.547-.606-2.083c-.367-.406-.954-.638-1.174-.59c-.29.062-.479.23-.725.818c-.145.348-.215.644-.335 1.335c-.115.656-.178.952-.309 1.34c-.395 1.176-1.364 2.395-2.665 3.236a11.877 11.877 0 0 1-2.937 1.37a.676.676 0 0 1-.2.03zm-.042 8.52c-.323.009-.613-.063-.856-.233c-.31-.217-.456-.559-.459-.953l.003-7.323c-.034-.39.081-.748.353-1.014c.255-.25.588-.368.94-.36h2.185A10.505 10.505 0 0 0 5.99 6.95c1.048-.678 1.82-1.65 2.115-2.526c.101-.302.155-.552.257-1.14c.138-.789.224-1.156.422-1.628c.41-.982.948-1.462 1.69-1.623c.73-.158 1.793.263 2.465 1.007c.745.824 1.074 1.855.952 3.129c-.052.548-.204 1.161-.454 1.844a10.509 10.509 0 0 1 2.578-.056c2.007.205 3.134 1.512 2.97 3.164c-.072.712-.33 1.317-.769 1.792c.369.711.516 1.414.424 2.1c-.106.79-.546 1.448-1.278 1.959c-.057.693-.216 1.246-.498 1.66a2.87 2.87 0 0 1-.851.834c-.108.684-.335 1.219-.706 1.595c-.615.626-1.714.999-2.718.931c-.953-.064-1.517-.18-2.847-.6c-.877-.277-2.693-.737-5.43-1.377zm1.701-8.831a.68.68 0 0 1 .68-.682a.68.68 0 0 1 .678.682v7.678a.68.68 0 0 1-.679.681a.68.68 0 0 1-.679-.681z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.598 7.761a.789.789 0 0 1 0 1.116l-4.057 4.056c-.668.74-.977 1.469-.963 2.211c.02 1 .277 1.555.872 2.23c.51.578 1.204.96 2.156 1.039c.682.057 1.388-.16 2.136-.678l4.357-4.357a.789.789 0 1 1 1.116 1.116L7.809 18.9l-.101.085c-1.056.75-2.14 1.092-3.234 1c-1.379-.115-2.441-.697-3.208-1.567c-.824-.936-1.239-1.83-1.265-3.244c-.022-1.17.446-2.277 1.396-3.327L5.482 7.76a.789.789 0 0 1 1.116 0M15.174 0c1.414.026 2.308.441 3.244 1.265c.87.767 1.452 1.83 1.567 3.208c.092 1.093-.25 2.178-1 3.234l-.085.101l-4.406 4.406a.789.789 0 1 1-1.116-1.116l4.357-4.357c.518-.748.735-1.454.678-2.136c-.08-.952-.46-1.647-1.038-2.156c-.676-.595-1.231-.853-2.231-.872c-.742-.014-1.471.295-2.21.963L8.876 6.598A.789.789 0 0 1 7.76 5.482l4.086-4.085c1.05-.95 2.157-1.418 3.327-1.396M11.95 7.698a.789.789 0 0 1 0 1.116l-2.924 2.924a.789.789 0 1 1-1.116-1.115l2.925-2.925a.789.789 0 0 1 1.115 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linkedin(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.04 17.043h-2.962v-4.64c0-1.107-.023-2.531-1.544-2.531c-1.544 0-1.78 1.204-1.78 2.449v4.722H7.793V7.5h2.844v1.3h.039c.397-.75 1.364-1.54 2.808-1.54c3.001 0 3.556 1.974 3.556 4.545zM4.447 6.194c-.954 0-1.72-.771-1.72-1.72s.767-1.72 1.72-1.72a1.72 1.72 0 0 1 0 3.44m1.484 10.85h-2.97V7.5h2.97zM18.522 0H1.476C.66 0 0 .645 0 1.44v17.12C0 19.355.66 20 1.476 20h17.042c.815 0 1.482-.644 1.482-1.44V1.44C20 .646 19.333 0 18.518 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linux(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.73 4.706c-.031.07-.061.106-.09.106c-.034.007-.052-.012-.052-.056c0-.09.067-.145.2-.168h.105a.186.186 0 0 0-.163.118m.957-.034a.167.167 0 0 0-.184-.05c.168-.082.28-.075.336.022c.021.045.01.078-.031.1c-.029.008-.069-.016-.121-.072M6.054 9.543a.927.927 0 0 0-.047.14a.676.676 0 0 1-.058.15a1.114 1.114 0 0 1-.105.146c-.049.074-.052.119-.01.134c.028.007.072-.02.131-.079a.561.561 0 0 0 .153-.279a.383.383 0 0 1 .021-.067a.248.248 0 0 0 .021-.095v-.034l-.01-.027c-.007-.019-.018-.027-.032-.023c-.028-.007-.05.005-.064.034m8.474 3.504c.028-.111.054-.214.079-.307c.024-.092.041-.189.052-.29c.01-.1.021-.18.032-.24a1.14 1.14 0 0 0 .005-.25a4.177 4.177 0 0 1-.01-.218c0-.037-.013-.119-.037-.245a6.692 6.692 0 0 1-.042-.23a8.873 8.873 0 0 0-.053-.278l-.057-.296c-.07-.357-.235-.74-.495-1.15c-.26-.409-.511-.688-.757-.836c.169.148.368.457.6.926c.61 1.205.799 2.24.567 3.102c-.077.298-.252.454-.525.469c-.218.03-.353-.039-.405-.206c-.053-.168-.08-.478-.084-.932a5.697 5.697 0 0 0-.121-1.194a7.686 7.686 0 0 0-.205-.77a3.218 3.218 0 0 0-.205-.508a3.396 3.396 0 0 0-.163-.274a1.111 1.111 0 0 0-.136-.167a5.309 5.309 0 0 1-.079-.078a7.534 7.534 0 0 0-.326-1.15c-.119-.305-.222-.513-.31-.625a3.082 3.082 0 0 1-.247-.368a1.347 1.347 0 0 1-.158-.446c-.028-.156-.007-.355.063-.597s.086-.426.048-.553c-.039-.126-.195-.22-.468-.279a1.997 1.997 0 0 1-.468-.2c-.207-.112-.331-.172-.373-.18c-.056-.007-.095-.103-.116-.29a1.166 1.166 0 0 1 .084-.568c.077-.194.204-.294.379-.302c.259-.022.438.09.536.335c.098.246.112.462.042.648c-.077.141-.084.24-.021.295c.063.056.168.058.315.006c.091-.03.137-.164.137-.402v-.413a2.596 2.596 0 0 0-.142-.557a.897.897 0 0 0-.468-.508a1.428 1.428 0 0 0-.284-.084c-.75.06-1.061.558-.935 1.495c0 .112-.004.168-.01.168c-.064-.067-.167-.106-.31-.117a1.463 1.463 0 0 0-.348.005c-.087.015-.141-.003-.162-.055c.007-.424-.05-.76-.169-1.005c-.119-.245-.276-.372-.473-.38c-.189-.007-.334.095-.436.307c-.102.212-.16.434-.174.664c-.007.112.006.25.037.413c.032.164.077.303.137.419c.06.115.114.165.163.15a.292.292 0 0 0 .168-.156c.028-.067.003-.097-.074-.09c-.05 0-.103-.053-.163-.161a.854.854 0 0 1-.1-.374a.704.704 0 0 1 .095-.413c.07-.112.189-.164.357-.156c.12 0 .214.078.284.234c.07.156.104.301.1.435a1.8 1.8 0 0 1-.015.246c-.155.111-.263.22-.326.323a.892.892 0 0 1-.29.262a3.555 3.555 0 0 0-.215.14a.647.647 0 0 0-.163.301c-.017.097.009.164.079.201c.098.06.186.132.263.218c.076.085.133.156.168.212c.035.056.1.104.194.145a1.1 1.1 0 0 0 .373.072c.33.015.687-.04 1.073-.167a3.36 3.36 0 0 1 .242-.078c.147-.045.268-.084.362-.117c.095-.034.198-.082.31-.145a.538.538 0 0 0 .221-.196c.063-.104.133-.134.21-.089a.154.154 0 0 1 .069.095c.01.04 0 .086-.032.134c-.031.048-.09.084-.173.106c-.14.045-.339.125-.595.24l-.478.217c-.308.142-.554.227-.736.257c-.175.037-.452.03-.83-.022c-.07-.015-.102-.008-.095.022c.007.03.066.1.178.212c.176.171.41.253.705.246a1.48 1.48 0 0 0 .378-.079a2.94 2.94 0 0 0 .379-.156c.119-.06.236-.125.352-.195c.116-.071.221-.134.316-.19c.094-.056.18-.1.257-.134c.078-.033.139-.043.184-.028c.046.015.076.056.09.123a.145.145 0 0 1-.01.05a.145.145 0 0 1-.043.056a.78.78 0 0 1-.152.106a1.34 1.34 0 0 1-.095.05c-.024.012-.06.03-.105.056a.811.811 0 0 1-.1.05a4.73 4.73 0 0 0-.71.492c-.276.223-.51.383-.698.48c-.19.097-.361.1-.515.01c-.148-.081-.368-.353-.663-.814c-.154-.23-.242-.312-.263-.245a.417.417 0 0 0-.01.111c0 .186-.053.396-.158.63a7.004 7.004 0 0 1-.31.62a2.026 2.026 0 0 0-.22.648a1 1 0 0 0 .12.703c-.161.044-.38.38-.657 1.004c-.277.625-.443 1.15-.5 1.574a8.2 8.2 0 0 0-.015.77c.004.379-.016.598-.058.658c-.056.179-.158.19-.305.034c-.224-.231-.35-.58-.378-1.05c-.014-.208 0-.416.042-.624c.028-.142.024-.209-.01-.201l-.043.056c-.252.483-.217 1.1.105 1.852c.035.09.123.193.263.312c.14.12.224.194.252.223c.14.172.505.508 1.094 1.01c.589.502.914.787.977.854a.631.631 0 0 1 .185.424a.72.72 0 0 1-.148.48a.644.644 0 0 1-.478.257c.056.111.157.277.305.497c.147.22.245.42.294.602c.049.182.074.444.074.787c.322-.179.347-.521.073-1.027a.989.989 0 0 0-.11-.179a5.023 5.023 0 0 1-.1-.133c-.021-.03-.028-.053-.021-.068c.02-.037.066-.072.136-.106c.07-.033.14-.024.21.028c.323.387.905.521 1.746.402c.932-.112 1.552-.435 1.86-.971c.162-.283.28-.394.358-.335c.084.045.12.238.105.58c-.007.186-.088.528-.242 1.027c-.063.171-.084.31-.063.419c.02.107.105.165.252.172c.021-.141.072-.427.153-.859c.08-.431.127-.766.142-1.004c.014-.157-.01-.43-.069-.82a6.298 6.298 0 0 1-.078-1.083c.007-.33.087-.593.241-.787c.106-.134.284-.2.537-.2c.007-.276.128-.473.362-.592a1.24 1.24 0 0 1 .763-.117c.273.04.483.124.63.25c0-.133-.193-.29-.578-.468m-6.05-9.095c-.039-.097-.08-.153-.121-.167c-.063-.015-.095.01-.095.078c.014.037.032.06.053.067c.07 0 .094.056.073.167c-.02.15.007.224.085.224c.02 0 .031-.012.031-.034a.622.622 0 0 0-.026-.335m4.363 2.405a.3.3 0 0 0-.137-.056a.381.381 0 0 1-.152-.06a.453.453 0 0 1-.1-.09a2.92 2.92 0 0 1-.132-.162a.232.232 0 0 0-.042-.045c-.007-.003-.02.002-.042.017c-.098.12-.074.281.074.485c.147.205.284.322.41.352c.063.007.114-.023.152-.09a.323.323 0 0 0 .037-.223a.23.23 0 0 0-.068-.128M11.038 4.11a.44.44 0 0 0-.053-.218a.413.413 0 0 0-.115-.14c-.043-.03-.074-.04-.095-.033c-.098.007-.123.033-.074.078l.042.023c.099.03.162.145.19.346c0 .022.028.014.084-.023zm.541-2.656a.416.416 0 0 0-.094-.079a4.735 4.735 0 0 1-.1-.067c-.106-.111-.19-.167-.253-.167c-.063.007-.103.035-.12.083a.257.257 0 0 0-.01.145a.267.267 0 0 1-.006.14a.361.361 0 0 1-.064.117a.547.547 0 0 0-.063.1c-.007.02.004.05.032.096c.028.022.056.022.084 0a3.49 3.49 0 0 0 .116-.1a.524.524 0 0 1 .157-.101c.007-.008.039-.011.095-.011a.538.538 0 0 0 .158-.023c.049-.015.08-.04.094-.078c0-.015-.008-.033-.026-.055m6.293 15.294a.492.492 0 0 1 .126.268a.678.678 0 0 1-.027.25a.585.585 0 0 1-.162.246a3.169 3.169 0 0 1-.247.218c-.08.063-.186.132-.316.206a13.611 13.611 0 0 1-.667.357c-.134.067-.228.115-.284.145a5.54 5.54 0 0 0-.9.625a10.43 10.43 0 0 0-.793.714c-.119.12-.357.192-.715.218c-.357.026-.67-.028-.935-.162a.915.915 0 0 1-.31-.262a1.461 1.461 0 0 1-.174-.285c-.035-.082-.112-.154-.231-.217a1.155 1.155 0 0 0-.494-.106a62.12 62.12 0 0 0-1.367-.011c-.134 0-.333.005-.6.016c-.266.012-.469.02-.61.028a2.286 2.286 0 0 0-.835.168a1.887 1.887 0 0 0-.563.335a2.029 2.029 0 0 1-.457.318a1.105 1.105 0 0 1-.562.128c-.203-.007-.592-.123-1.167-.346a10.323 10.323 0 0 0-2.072-.586c-.224-.041-.4-.074-.525-.1a5.04 5.04 0 0 1-.415-.106a1.123 1.123 0 0 1-.353-.162a.538.538 0 0 1-.178-.218c-.07-.17-.046-.418.074-.742c.119-.323.182-.526.189-.608a1.839 1.839 0 0 0-.042-.446a8.855 8.855 0 0 0-.105-.475a1.495 1.495 0 0 1-.048-.407c.004-.134.04-.234.11-.301c.1-.09.299-.142.6-.157c.301-.015.511-.06.63-.134c.211-.134.358-.264.443-.39c.084-.127.126-.317.126-.57c.147.543.035.938-.337 1.183c-.224.149-.515.205-.872.168c-.239-.023-.39.014-.453.111c-.09.112-.073.324.053.636c.014.045.042.112.084.201c.042.09.072.156.09.201a.99.99 0 0 1 .047.19a.968.968 0 0 1 .01.245c0 .112-.06.294-.178.547c-.12.253-.168.431-.148.536c.021.126.151.223.39.29c.14.044.436.113.888.206c.452.093.8.169 1.046.229c.168.045.427.126.778.245c.35.119.64.205.867.257c.228.052.422.067.583.045c.302-.045.528-.15.679-.313a.847.847 0 0 0 .242-.536a2.146 2.146 0 0 0-.08-.652a3.257 3.257 0 0 0-.199-.58a8.313 8.313 0 0 0-.21-.408c-.848-1.414-1.44-2.314-1.777-2.701c-.476-.55-.872-.7-1.188-.446c-.077.067-.13.01-.158-.168a1.98 1.98 0 0 1-.02-.424a1.82 1.82 0 0 1 .104-.58c.063-.171.147-.346.253-.525c.105-.178.182-.334.23-.468c.057-.157.15-.424.28-.804c.129-.38.232-.67.31-.87c.077-.2.182-.428.315-.68c.133-.254.27-.455.41-.603c.77-1.064 1.205-1.79 1.304-2.177c-.084-.833-.14-1.986-.169-3.46c-.014-.669.07-1.232.253-1.69c.182-.457.554-.846 1.114-1.166C8.83.079 9.195.001 9.65.001c.372-.008.743.043 1.115.15c.371.108.683.263.935.464c.4.312.72.764.962 1.356c.242.59.345 1.14.31 1.645c-.035.707.07 1.503.316 2.389c.238.84.704 1.651 1.398 2.432c.385.44.734 1.046 1.046 1.82c.312.773.52 1.484.625 2.131c.056.365.074.679.053.943a2.07 2.07 0 0 1-.126.62c-.063.148-.133.23-.21.245c-.07.015-.153.085-.248.212c-.094.126-.189.258-.284.396a1.315 1.315 0 0 1-.425.374c-.19.111-.403.164-.642.156a1.373 1.373 0 0 1-.33-.056a.52.52 0 0 1-.237-.15a2.05 2.05 0 0 1-.142-.173a1.543 1.543 0 0 1-.12-.229a3.231 3.231 0 0 1-.095-.217c-.155-.276-.298-.387-.431-.335c-.134.051-.232.234-.295.546c-.063.313-.038.674.074 1.083c.14.52.144 1.246.01 2.176c-.07.483-.007.857.19 1.122c.196.264.452.387.767.368c.315-.02.613-.151.894-.396c.413-.365.727-.612.94-.742c.214-.13.577-.289 1.089-.475c.371-.134.64-.27.809-.407c.168-.138.233-.266.194-.385a.63.63 0 0 0-.262-.318a2.474 2.474 0 0 0-.542-.262c-.23-.082-.404-.26-.52-.536a1.88 1.88 0 0 1-.158-.81c.01-.263.065-.44.163-.53a2.862 2.862 0 0 0 .237 1.084c.052.111.124.217.215.318c.091.1.165.171.22.212c.057.04.132.09.227.145c.095.056.153.091.174.106c.14.09.248.18.326.274"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Loading(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 16a2 2 0 1 1 0 4a2 2 0 0 1 0-4m-6.259-3a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5m11.578.5a2 2 0 1 1 0 4a2 2 0 0 1 0-4M18.5 9.319a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3M2.5 6a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5m15.286-.793a1 1 0 1 1 0 2a1 1 0 0 1 0-2M8 0a3 3 0 1 1 0 6a3 3 0 0 1 0-6m7.5 3a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0a8 8 0 0 1 8 8v11a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V8a8 8 0 0 1 8-8m6.6 10.245H3.4V18.6h13.2zm-6.239 1.378a1.5 1.5 0 0 1 .695 2.83v2.169a.7.7 0 0 1-1.4 0v-2.175a1.5 1.5 0 0 1 .705-2.824M10 1.4A6.6 6.6 0 0 0 3.4 8v.845h13.2V8A6.6 6.6 0 0 0 10 1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Login(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.76 0C15.417 0 20 4.477 20 10S15.416 20 9.76 20c-3.191 0-6.142-1.437-8.07-3.846a.644.644 0 0 1 .115-.918a.68.68 0 0 1 .94.113a8.96 8.96 0 0 0 7.016 3.343c4.915 0 8.9-3.892 8.9-8.692c0-4.8-3.985-8.692-8.9-8.692a8.961 8.961 0 0 0-6.944 3.255a.68.68 0 0 1-.942.101a.644.644 0 0 1-.103-.92C3.703 1.394 6.615 0 9.761 0m.545 6.862l2.707 2.707c.262.262.267.68.011.936L10.38 13.15a.662.662 0 0 1-.937-.011a.662.662 0 0 1-.01-.937l1.547-1.548l-10.31.001A.662.662 0 0 1 0 10c0-.361.3-.654.67-.654h10.268L9.38 7.787a.662.662 0 0 1-.01-.937a.662.662 0 0 1 .935.011"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Logout(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.24 0c3.145 0 6.057 1.395 7.988 3.744a.644.644 0 0 1-.103.92a.68.68 0 0 1-.942-.1a8.961 8.961 0 0 0-6.944-3.256c-4.915 0-8.9 3.892-8.9 8.692c0 4.8 3.985 8.692 8.9 8.692a8.962 8.962 0 0 0 7.016-3.343a.68.68 0 0 1 .94-.113a.644.644 0 0 1 .115.918C16.382 18.564 13.431 20 10.24 20C4.583 20 0 15.523 0 10S4.584 0 10.24 0m6.858 7.16l2.706 2.707c.262.261.267.68.012.936l-2.644 2.643a.662.662 0 0 1-.936-.01a.662.662 0 0 1-.011-.937l1.547-1.547H7.462a.662.662 0 0 1-.67-.654c0-.362.3-.655.67-.655h10.269l-1.558-1.558a.662.662 0 0 1-.011-.936a.662.662 0 0 1 .936.011"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mail(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7.172 11.334l2.83 1.935l2.728-1.882l6.115 6.033c-.161.052-.333.08-.512.08H1.667c-.22 0-.43-.043-.623-.12zM20 6.376v9.457c0 .247-.054.481-.15.692l-5.994-5.914zM0 6.429l6.042 4.132l-5.936 5.858A1.663 1.663 0 0 1 0 15.833zM18.333 2.5c.92 0 1.667.746 1.667 1.667v.586L9.998 11.648L0 4.81v-.643C0 3.247.746 2.5 1.667 2.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M18.333 2.5c.92 0 1.667.746 1.667 1.667v11.666c0 .92-.746 1.667-1.667 1.667H1.667C.747 17.5 0 16.754 0 15.833V4.167C0 3.247.746 2.5 1.667 2.5zM7.168 11.328l-4.91 4.852h15.325l-4.857-4.802L10 13.265zM18.64 7.292l-4.796 3.316l4.796 4.736zm-17.279.061v7.836l4.686-4.631zm16.956-3.532H1.698a.358.358 0 0 0-.25.086a.26.26 0 0 0-.085.222v1.62L10 11.656l8.644-5.965V4.199c.001-.134-.03-.231-.092-.292a.306.306 0 0 0-.234-.086"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Man(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.309 0c.381 0 .691.31.691.695v5.163c0 .383-.31.694-.691.694a.693.693 0 0 1-.692-.694V2.372l-5.048 5.245a7.629 7.629 0 0 1 1.643 4.743c0 4.22-3.405 7.64-7.606 7.64C3.406 20 0 16.58 0 12.36s3.405-7.64 7.606-7.64c1.919 0 3.671.714 5.01 1.891l5.022-5.222h-3.745A.693.693 0 0 1 13.2.695c0-.384.31-.695.692-.695zM7.606 6.11c-3.437 0-6.223 2.798-6.223 6.25s2.786 6.251 6.223 6.251s6.223-2.799 6.223-6.25c0-3.453-2.786-6.251-6.223-6.251"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Map(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.219 1.943c.653.512 1.103 1.339 1.287 2.205a.474.474 0 0 1 .065.026l2.045.946a.659.659 0 0 1 .384.597v12.367a.665.665 0 0 1-.85.634l-5.669-1.6l-6.74 1.858a.674.674 0 0 1-.371-.004L.474 17.217a.66.66 0 0 1-.474-.63V3.998c0-.44.428-.756.855-.632l5.702 1.661l2.898-.887a.734.734 0 0 1 .122-.025c.112-.656.425-1.286.95-1.9c.623-.73 1.716-1.158 2.781-1.209c1.105-.053 1.949.183 2.91.936M1.333 4.881v11.215l4.87 1.449V6.298zm8.209.614l-2.006.613v11.279l5.065-1.394v-3.295c0-.364.299-.659.667-.659c.368 0 .666.295.666.66v3.177l4.733 1.335V6.136l-1.12-.52c-.019.11-.043.218-.073.323A6.134 6.134 0 0 1 16.4 8.05l-2.477 3.093a.67.67 0 0 1-1.073-.037l-2.315-3.353c-.382-.534-.65-1.01-.801-1.436a3.744 3.744 0 0 1-.192-.822m3.83-3.171c-.726.035-1.472.327-1.827.742c-.427.5-.637.968-.679 1.442c-.05.571-.016.974.126 1.373c.105.295.314.669.637 1.12l1.811 2.622l1.91-2.385a4.812 4.812 0 0 0 .841-1.657c.24-.84-.122-2.074-.8-2.604c-.695-.545-1.22-.692-2.018-.653m.138.697c1.104 0 2 .885 2 1.977a1.988 1.988 0 0 1-2 1.977c-1.104 0-2-.885-2-1.977s.896-1.977 2-1.977m0 1.318a.663.663 0 0 0-.667.659c0 .364.299.659.667.659a.663.663 0 0 0 .666-.66a.663.663 0 0 0-.666-.658"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Meh(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m2.529 13.674H6.947a.698.698 0 0 0 0 1.396h5.582a.698.698 0 1 0 0-1.396M5.814 6.28a1.395 1.395 0 1 0 0 2.79a1.395 1.395 0 0 0 0-2.79m8.372 0a1.395 1.395 0 1 0 0 2.79a1.395 1.395 0 0 0 0-2.79"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MehO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21m2.529 12.28a.698.698 0 1 1 0 1.395H6.947a.698.698 0 0 1 0-1.396zM5.814 6.278a1.395 1.395 0 1 1 0 2.79a1.395 1.395 0 0 1 0-2.79m8.372 0a1.395 1.395 0 1 1 0 2.79a1.395 1.395 0 0 1 0-2.79"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Menu(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.497 15.602a.7.7 0 1 1 0 1.398H.7a.7.7 0 1 1 0-1.398zm15.803 0a.7.7 0 1 1 0 1.398H5.529a.7.7 0 1 1 0-1.398zM3.497 9.334a.7.7 0 1 1 0 1.399H.7a.7.7 0 1 1 0-1.399zm15.803 0a.7.7 0 1 1 0 1.399H5.528a.7.7 0 1 1 0-1.399zM3.497 3a.7.7 0 1 1 0 1.398H.7A.7.7 0 1 1 .7 3zM19.3 3a.7.7 0 1 1 0 1.398H5.528a.7.7 0 1 1 0-1.398z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuFold(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.296 16.608c.386 0 .7.312.7.696a.698.698 0 0 1-.7.696H.7a.698.698 0 0 1-.7-.696c0-.384.313-.696.7-.696zM4.033 6.625a.698.698 0 0 1-.003.987L1.697 9.944l2.32 2.32a.698.698 0 0 1-.003.987a.698.698 0 0 1-.987.003L.216 10.442a.698.698 0 0 1 .002-.987l2.828-2.827a.698.698 0 0 1 .987-.003M19.3 9.131c.387 0 .7.312.7.696a.698.698 0 0 1-.7.696H5.183a.698.698 0 0 1-.7-.696c0-.384.313-.696.7-.696zM19.296 2c.386 0 .7.312.7.696a.698.698 0 0 1-.7.696H.7a.698.698 0 0 1-.7-.696C0 2.312.313 2 .7 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuUnfold(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.3 16.608c.386 0 .699.312.699.696a.698.698 0 0 1-.7.696H.703a.698.698 0 0 1-.7-.696c0-.384.314-.696.7-.696zM15.98 6.492a.698.698 0 0 1 .988.003l2.827 2.827a.698.698 0 0 1 .003.987l-2.812 2.812a.698.698 0 0 1-.987-.003a.698.698 0 0 1-.003-.987l2.32-2.32l-2.332-2.332a.698.698 0 0 1-.003-.987m-1.163 2.639c.387 0 .7.312.7.696a.698.698 0 0 1-.7.696H.7a.698.698 0 0 1-.7-.696c0-.384.313-.696.7-.696zM19.3 2c.387 0 .7.312.7.696a.698.698 0 0 1-.7.696H.703a.698.698 0 0 1-.7-.696c0-.384.314-.696.7-.696z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Message(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.342 0 10 4.41 10 9.5c0 5.004-4.553 8.942-10 8.942a11.01 11.01 0 0 1-3.43-.546c-.464.45-.623.603-1.608 1.553c-.71.536-1.378.718-1.975.38c-.602-.34-.783-1.002-.66-1.874l.4-2.319C.99 14.002 0 11.842 0 9.5C0 4.41 4.657 0 10 0m0 1.4c-4.586 0-8.6 3.8-8.6 8.1c0 2.045.912 3.928 2.52 5.33l.02.017l.297.258l-.067.39l-.138.804l-.037.214l-.285 1.658a2.79 2.79 0 0 0-.03.337v.095c0 .005-.001.007-.002.008c.007-.01.143-.053.376-.223l2.17-2.106l.414.156a9.589 9.589 0 0 0 3.362.605c4.716 0 8.6-3.36 8.6-7.543c0-4.299-4.014-8.1-8.6-8.1M5.227 7.813a1.5 1.5 0 1 1 0 2.998a1.5 1.5 0 0 1 0-2.998m4.998 0a1.5 1.5 0 1 1 0 2.998a1.5 1.5 0 0 1 0-2.998m4.997 0a1.5 1.5 0 1 1 0 2.998a1.5 1.5 0 0 1 0-2.998"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.7 10.4A.7.7 0 0 1 .7 9h18.606a.7.7 0 0 1 0 1.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircle(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m4.455 9.09H6.273a.818.818 0 0 0 0 1.637h8.182a.818.818 0 0 0 0-1.636"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircleO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21m4.588 7.923a.682.682 0 1 1 0 1.364H6.143a.682.682 0 0 1 0-1.364z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSquare(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.182 0C19.186 0 20 .814 20 1.818v16.364A1.818 1.818 0 0 1 18.182 20H1.818A1.818 1.818 0 0 1 0 18.182V1.818C0 .814.814 0 1.818 0zm-3.496 9.3H5.91a.682.682 0 0 0 0 1.363h8.777a.682.682 0 1 0 0-1.364"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSquareO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.818 1.364a.455.455 0 0 0-.454.454v16.364c0 .25.203.454.454.454h16.364a.455.455 0 0 0 .454-.454V1.818a.455.455 0 0 0-.454-.454zM18.182 0C19.186 0 20 .814 20 1.818v16.364A1.818 1.818 0 0 1 18.182 20H1.818A1.818 1.818 0 0 1 0 18.182V1.818C0 .814.814 0 1.818 0zm-3.496 9.3H5.91a.682.682 0 0 0 0 1.363h8.777a.682.682 0 1 0 0-1.364"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mobile(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 0a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2zm.6 15.388H4.4V18a.6.6 0 0 0 .6.6h10a.6.6 0 0 0 .6-.6zM10 16a1 1 0 1 1 0 2a1 1 0 0 1 0-2m5-14.6H5a.6.6 0 0 0-.6.6v11.988h11.2V2a.6.6 0 0 0-.6-.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func More(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 7.5a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5m15 0a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5m-7.274 0a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notification(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.026 18.64a.111.111 0 0 0 .011 0zm-.197-.124l.003-17.038l-.07.07L4.94 6.316a.697.697 0 0 1-.49.2H1.968a.646.646 0 0 0-.365.242c-.13.156-.207.388-.22.69l-.001 5.411c.01.19.074.334.203.463c.123.122.343.209.622.238h2.304c.183 0 .359.071.488.199zM10.207.019c.583.097.953.522 1.005 1.165V18.84l-.005.085c-.078.603-.462 1.032-1.067 1.074c-.451.03-.871-.137-1.252-.484L4.224 14.92l-2.082-.002c-.644-.06-1.166-.267-1.54-.64A2.035 2.035 0 0 1 0 12.896V7.42c.027-.606.2-1.12.532-1.522a2.01 2.01 0 0 1 1.27-.736l.105-.008h2.258L8.77.6c.428-.444.913-.668 1.437-.58m6.21 2.227C18.602 3.618 20 6.576 20 9.862c0 3.285-1.398 6.243-3.582 7.615a.698.698 0 0 1-.955-.208a.675.675 0 0 1 .211-.94c1.754-1.102 2.943-3.618 2.943-6.467c0-2.85-1.189-5.366-2.943-6.468a.675.675 0 0 1-.211-.94a.698.698 0 0 1 .954-.208m-2.301 2.686c1.36 1.007 2.197 2.88 2.197 4.93c0 2.165-.935 4.128-2.42 5.084a.698.698 0 0 1-.957-.198a.675.675 0 0 1 .2-.943c1.068-.686 1.794-2.212 1.794-3.943c0-1.644-.654-3.108-1.645-3.841a.674.674 0 0 1-.137-.954a.7.7 0 0 1 .968-.135"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Opera(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.71 4.365c-1.107 1.305-1.822 3.236-1.872 5.4v.47c.05 2.165.765 4.093 1.871 5.4c1.434 1.862 3.566 3.044 5.95 3.044a7.208 7.208 0 0 0 4.005-1.226a9.94 9.94 0 0 1-7.139 2.535A9.998 9.998 0 0 1 0 10C0 4.476 4.478 0 10 0h.037a9.97 9.97 0 0 1 6.628 2.546a7.239 7.239 0 0 0-4.008-1.226c-2.382 0-4.514 1.183-5.95 3.045zM20 10a9.969 9.969 0 0 1-3.335 7.454c-2.565 1.25-4.955.376-5.747-.17c2.52-.554 4.423-3.6 4.423-7.284c0-3.685-1.903-6.73-4.423-7.283c.791-.545 3.182-1.42 5.747-.171A9.967 9.967 0 0 1 20 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperClip(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.264 8.579a.683.683 0 0 1-.975 0a.704.704 0 0 1 0-.987L8.32 1.5C9.68.444 11.048-.063 12.41.006c1.716.088 3.052.742 4.186 1.815C17.752 2.915 18.5 4.476 18.5 6.368c0 1.452-.422 2.73-1.313 3.864l-8.503 8.76c-.86.705-1.816 1.046-2.84 1.005c-1.3-.054-2.267-.474-2.986-1.185c-.842-.831-1.358-1.852-1.358-3.225c0-1.092.377-2.1 1.155-3.046L10.139 4.9c.6-.64 1.187-1.02 1.787-1.112a2.486 2.486 0 0 1 2.2.755c.532.563.76 1.265.68 2.064c-.055.545-.278 1.047-.688 1.528l-6.88 7.048a.683.683 0 0 1-.974.006a.704.704 0 0 1-.006-.987l6.847-7.012c.2-.235.305-.472.33-.724c.04-.4-.056-.695-.305-.958a1.118 1.118 0 0 0-1-.34c-.243.037-.583.258-1.002.704l-7.453 7.607c-.537.655-.797 1.35-.797 2.109c0 .954.345 1.637.942 2.226c.475.47 1.12.75 2.08.79c.68.027 1.31-.198 1.858-.642l8.397-8.65c.645-.827.967-1.8.967-2.943c0-1.482-.577-2.684-1.468-3.528c-.91-.862-1.95-1.37-3.313-1.44c-1.008-.052-2.065.34-3.117 1.146z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.78 2.044a.73.73 0 0 1 .72.741v14.473a.73.73 0 0 1-.72.742a.731.731 0 0 1-.72-.742V2.785a.73.73 0 0 1 .72-.741M5.22 2a.73.73 0 0 1 .72.742v14.473a.73.73 0 0 1-.72.741a.731.731 0 0 1-.72-.741V2.742A.73.73 0 0 1 5.22 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseCircle(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0M8.056 5.455a.682.682 0 0 0-.682.681v7.272a.682.682 0 0 0 1.364 0V6.136a.682.682 0 0 0-.682-.681m4.546 0a.682.682 0 0 0-.682.681v7.272a.682.682 0 0 0 1.363 0V6.136a.682.682 0 0 0-.681-.681"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseCircleO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21m-1.944 4.06c.377 0 .682.305.682.681v7.272a.682.682 0 0 1-1.364 0V6.136c0-.376.306-.681.682-.681m4.546 0c.376 0 .681.305.681.681v7.272a.682.682 0 0 1-1.363 0V6.136c0-.376.305-.681.682-.681"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pay(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.67 8.537a.3.3 0 0 0-.302.296v2.212a.3.3 0 0 0 .303.296h6.663a.3.3 0 0 0 .303-.296V8.833a.3.3 0 0 0-.303-.296zm4.086-7.036c.922.044 1.585.226 2.005.612c.415.382.628.935.67 1.667v2.097a.674.674 0 0 1-.681.666a.674.674 0 0 1-.682-.666l.001-2.059c-.022-.38-.113-.616-.243-.736c-.126-.116-.51-.22-1.103-.25H2.647c-.537.02-.886.122-1.055.267c-.13.111-.228.417-.229.946l-.003 11.77c.05.514.163.857.308 1.028c.11.13.451.26.953.324h13.116c.614.012.976-.08 1.098-.203c.135-.137.233-.497.233-1.086v-2.045c0-.367.305-.666.682-.666c.376 0 .681.299.681.666v2.045c0 .9-.184 1.573-.615 2.01c-.444.45-1.15.63-2.093.61L2.54 18.495c-.897-.104-1.54-.35-1.923-.803c-.347-.41-.54-.995-.617-1.813V4.044c.002-.876.212-1.535.694-1.947c.442-.38 1.08-.565 1.927-.597zm2.578 5.704c.92 0 1.666.729 1.666 1.628v2.212c0 .899-.746 1.628-1.666 1.628h-6.663c-.92 0-1.666-.73-1.666-1.628V8.833c0-.899.746-1.628 1.666-1.628zm-4.997 1.94c-.46 0-.833.36-.833.803c0 .444.373.803.833.803c.46 0 .833-.36.833-.803c0-.444-.373-.804-.833-.804"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PayCircleO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21M7.113 5.112a.728.728 0 0 1 1.023-.004l1.86 1.842l1.848-1.836a.723.723 0 0 1 1.019-.002a.704.704 0 0 1-.002 1.006l-2.366 2.338a.06.06 0 0 1-.042.02h3.342c.377 0 .705.34.705.713v.08c0 .374-.328.656-.705.656h-3.078v1.27h2.438c.376 0 .705.324.705.697v.08c0 .372-.329.652-.705.652h-2.438v2.962c0 .189-.075.37-.21.505a.724.724 0 0 1-1.23-.505v-2.962H6.825a.65.65 0 0 1-.67-.653v-.08c0-.372.294-.696.67-.696h2.452v-1.27h-3.09a.679.679 0 0 1-.687-.67v-.066a.69.69 0 0 1 .687-.685h3.332c-.014 0-.028-.02-.042-.034L7.11 6.126a.714.714 0 0 1 .002-1.014"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Picasa(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M18.983 14.377c-1.533 3.166-4.617 5.3-8.07 5.623H9.075a9.967 9.967 0 0 1-3.174-.843v-4.78ZM4.853 9.964v8.615A10.045 10.045 0 0 1 .72 13.726l1.889-1.72l.378-.344zm10.324-8.491C18.152 3.287 20 6.539 20 10.02c0 1.122-.195 2.232-.567 3.3h-4.256Zm-9.734-.361l1.59 1.446l.563.512l1.63 1.484L4.71 8.668l-.488.444l-3.859 3.512A10.036 10.036 0 0 1 0 10.021a9.973 9.973 0 0 1 5.443-8.91M9.999 0c1.433 0 2.82.304 4.124.898V7.58L6.507.653A9.91 9.91 0 0 1 10 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Picture(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.09 2a.9.9 0 0 1 .91.889V17.11a.9.9 0 0 1-.91.889H.91A.9.9 0 0 1 0 17.11V2.89A.9.9 0 0 1 .91 2zM5.416 8.417l-4.06 4.042v4.217H18.64v-1.433l-3.2-3.12l-2.777 2.333c-.166.117-.326.168-.48.155a.792.792 0 0 1-.439-.189zm13.22-5.086H1.362v7.23L4.968 6.97a.718.718 0 0 1 .44-.156c.155 0 .291.047.41.14l6.431 6.088l2.805-2.35a.704.704 0 0 1 .421-.146a.69.69 0 0 1 .418.145l2.742 2.665zM15.273 5.23c.753 0 1.363.597 1.363 1.333s-.61 1.333-1.363 1.333c-.754 0-1.364-.597-1.364-1.333s.61-1.333 1.364-1.333"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PieChart(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><path id="uiwPieChart0" d="M1.37 18.615h17.945c.379 0 .685.31.685.693a.689.689 0 0 1-.685.692H.685A.689.689 0 0 1 0 19.308V.692C0 .31.306 0 .685 0c.378 0 .684.31.684.692zM2.836 17.4l2.753-5.57l2.754 1.392l3.442-4.179l2.754 2.09l4.131-6.268V17.4z"/></defs><use fill="currentColor" href="#uiwPieChart0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pinterest(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.015 0C4.484 0 0 4.473 0 9.99c0 4.232 2.638 7.847 6.364 9.3c-.088-.79-.166-2.002.034-2.865c.183-.78 1.175-4.964 1.175-4.964s-.3-.6-.3-1.484c0-1.386.808-2.426 1.811-2.426c.855 0 1.268.64 1.268 1.406c0 .858-.545 2.14-.829 3.327c-.238.994.502 1.804 1.483 1.804c1.778 0 3.148-1.87 3.148-4.572c0-2.384-1.723-4.058-4.184-4.058c-2.848 0-4.518 2.135-4.518 4.333c0 .86.329 1.786.742 2.284c.083.1.094.188.071.288c-.075.312-.244.999-.279 1.135c-.044.188-.143.226-.335.138c-1.249-.575-2.032-2.398-2.032-3.872c0-3.146 2.296-6.043 6.616-6.043c3.474 0 6.175 2.472 6.175 5.769c0 3.446-2.178 6.218-5.207 6.218c-1.014 0-1.966-.524-2.304-1.149l-.625 2.374c-.225.87-.84 1.96-1.252 2.621A10.07 10.07 0 0 0 9.988 20C15.508 20 20 15.53 20 10.01C20 4.493 15.507.023 9.988.023z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayCircle(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0M8.182 5.904v8.193c0 .467.436.585.718.28l3.675-3.961a.644.644 0 0 0 0-.831L8.9 5.623c-.282-.305-.718-.187-.718.281"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayCircleO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21M8.182 5.904c0-.468.436-.586.718-.281l3.675 3.962a.644.644 0 0 1 0 .831L8.9 14.377c-.282.305-.718.187-.718-.28Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c.423 0 .766.343.766.766v8.467h8.468a.766.766 0 1 1 0 1.533h-8.468v8.468a.766.766 0 1 1-1.532 0l-.001-8.468H.766a.766.766 0 0 1 0-1.532l8.467-.001V.766A.768.768 0 0 1 10 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircle(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m0 5.475a.682.682 0 0 0-.682.681V9.28H6.195a.682.682 0 0 0-.674.582l-.008.1c0 .377.305.682.682.682h3.123v3.123c0 .343.252.626.581.675l.101.007a.682.682 0 0 0 .682-.682l-.001-3.123h3.124a.682.682 0 0 0 .674-.58l.008-.102a.682.682 0 0 0-.682-.681l-3.124-.001V6.156a.682.682 0 0 0-.58-.674Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircleO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21m0 4.08c.377 0 .682.305.682.681L10.68 9.28h3.124a.682.682 0 1 1 0 1.364H10.68v3.123a.682.682 0 1 1-1.363 0v-3.123H6.195a.682.682 0 1 1 0-1.363l3.123-.001V6.156c0-.376.305-.681.682-.681"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSquare(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.182 0C19.186 0 20 .814 20 1.818v16.364A1.818 1.818 0 0 1 18.182 20H1.818A1.818 1.818 0 0 1 0 18.182V1.818C0 .814.814 0 1.818 0zm-7.884 4.91a.682.682 0 0 0-.682.682L9.615 9.3H5.909a.682.682 0 0 0-.674.581l-.008.1c0 .378.306.683.682.683l3.706-.001v3.707c0 .343.253.626.582.675l.1.007a.682.682 0 0 0 .682-.682v-3.707h3.707a.682.682 0 0 0 .675-.58l.007-.101a.682.682 0 0 0-.682-.682H10.98V5.592a.682.682 0 0 0-.58-.674Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSquareO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.818 1.364a.455.455 0 0 0-.454.454v16.364c0 .25.203.454.454.454h16.364a.455.455 0 0 0 .454-.454V1.818a.455.455 0 0 0-.454-.454zM18.182 0C19.186 0 20 .814 20 1.818v16.364A1.818 1.818 0 0 1 18.182 20H1.818A1.818 1.818 0 0 1 0 18.182V1.818C0 .814.814 0 1.818 0zm-7.884 4.91a.682.682 0 0 0-.682.682L9.615 9.3H5.909a.682.682 0 0 0-.674.581l-.008.1c0 .378.306.683.682.683l3.706-.001v3.707c0 .343.253.626.582.675l.1.007a.682.682 0 0 0 .682-.682v-3.707h3.707a.682.682 0 0 0 .675-.58l.007-.101a.682.682 0 0 0-.682-.682H10.98V5.592a.682.682 0 0 0-.58-.674Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Poweroff(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.074 2.328a.663.663 0 0 1-.143.927c-2.24 1.642-3.603 4.27-3.603 7.03c0 4.546 4.011 8.389 8.679 8.389c4.665 0 8.665-3.84 8.665-8.39c0-2.73-1.33-5.331-3.523-6.976a.663.663 0 1 1 .797-1.06C18.47 4.14 20 7.131 20 10.283C20 15.578 15.393 20 10.007 20C4.618 20 0 15.576 0 10.284c0-3.187 1.569-6.21 4.146-8.098a.664.664 0 0 1 .928.142M10 0a.7.7 0 0 1 .7.7v7.8a.7.7 0 0 1-1.4 0V.7A.7.7 0 0 1 10 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Printer(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.398.02h9.831c.628.022 1.113.204 1.407.61c.234.322.355.728.38 1.241l-.001 2.093h1.199c.627.022 1.112.203 1.406.609c.234.322.355.729.38 1.242v8.097c-.002.626-.046 1.094-.219 1.418c-.268.5-.764.723-1.59.789l-1.175-.001v1.949c0 .457-.082.859-.257 1.2c-.246.478-.715.719-1.317.733H4.725c-.594 0-1.082-.217-1.386-.662c-.243-.357-.37-.779-.385-1.271v-1.949h-1.02c-.737 0-1.3-.22-1.599-.732c-.223-.382-.327-.871-.333-1.476V6.025C-.02 5.389.11 4.868.427 4.482c.383-.468 1.047-.594 1.956-.518l.633-.001l.001-1.882c-.022-.636.108-1.156.425-1.543C3.825.071 4.49-.055 5.398.02m10.279 12.131H4.293v5.895c.007.227.06.4.152.536c.033.048.1.078.28.078h10.701a2.03 2.03 0 0 0 .13-.006h.012c.07-.135.11-.33.11-.587zm-1.378 4.534a.67.67 0 0 1 0 1.34h-8.58a.67.67 0 0 1 0-1.34zm0-1.897a.67.67 0 0 1 0 1.34h-8.58a.67.67 0 0 1 0-1.34zM2.329 5.302c-.569.002-.7.002-.82.14c-.12.14-.18.198-.168.56v7.902c.004.384.016.568.103.708s.18.167.49.167l1.02-.001v-2.627H2.82a.67.67 0 0 1 0-1.34h14.404a.67.67 0 0 1 0 1.34h-.208v2.627l1.12.003c.208-.017.327.005.435-.108c.108-.114.089-.39.09-.763V5.848c-.011-.238-.033-.318-.107-.405c-.074-.088-.13-.13-.363-.139zm11.97 7.59a.67.67 0 0 1 0 1.34h-8.58a.67.67 0 0 1 0-1.34zm1.785-6.978a2.01 2.01 0 1 1-.002 4.02a2.01 2.01 0 0 1 .002-4.02m0 1.34a.67.67 0 1 0 0 1.34a.67.67 0 0 0 0-1.34M4.478 1.388c-.082.1-.134.309-.122.67l-.001 1.905l11.321.001l.001-2.06c-.012-.237-.059-.397-.125-.487c-.005-.008-.112-.048-.346-.056l-9.862-.003c-.557-.045-.851.011-.866.03"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Qq(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.496 13.607c-.134-1.931-1.372-3.55-2.088-4.387c.1-.243.341-1.653-.593-2.615c.002-.023.002-.046.002-.068C15.817 2.743 13.237.012 10 0C6.763.013 4.183 2.743 4.183 6.537c0 .023 0 .046.002.068c-.934.962-.692 2.372-.593 2.615c-.715.837-1.953 2.456-2.088 4.387c-.024.508.051 1.248.288 1.577c.289.4 1.081-.081 1.648-1.362c.158.594.521 1.5 1.345 2.649c-1.378.33-1.771 1.752-1.307 2.53c.327.548 1.075.999 2.365.999c2.296 0 3.31-.645 3.763-1.095c.091-.099.225-.146.394-.146c.17 0 .303.047.394.146c.453.45 1.467 1.095 3.762 1.095c1.29 0 2.039-.45 2.366-.999c.464-.778.07-2.2-1.307-2.53c.824-1.15 1.188-2.055 1.345-2.649c.567 1.281 1.36 1.763 1.648 1.362c.237-.33.312-1.07.288-1.577"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Qrcode(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.363 1.364v6.363h6.363V1.364zm7.726 9.545V20H0v-9.09zm-1.363 1.364H1.363v6.363h6.363zm-1.732 1.873v2.727H3.267v-2.727zM9.089 0v9.09H0V0zM5.994 3.237H3.267v2.727h2.727zm13.999 13.695v3.067h-4.61v-1.364h3.246v-1.703zm-6.15 1.701v1.364h-2.937v-1.364zm-1.575-7.73v5.999h-1.363v-5.999zm3.27 3.243v1.362h1.366v1.363h-2.729v-2.725zM20 13.85v1.529h-1.363v-1.53zm-.01-2.907v1.364h-4.604v-1.364zM19.996 0v9.09h-9.089V0zm-1.363 1.364h-6.362v6.363h6.362zm-1.732 1.873v2.727h-2.727V3.237z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionCircle(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m0 13.636a.91.91 0 1 0 0 1.819a.91.91 0 0 0 0-1.819m2.68-8.306c-.74-.745-1.843-1.083-2.768-.963c-1.002.13-1.553.33-2.312.997C6.967 5.92 6.571 6.7 6.406 7.68a.682.682 0 1 0 1.345.227c.115-.686.367-1.183.75-1.519c.522-.458.83-.57 1.587-.668c.519-.068 1.196.14 1.625.572c.407.409.52.814.455 1.285c-.049.348-.294.784-.607 1.066c-.111.096-.563.487-.678.59a8.904 8.904 0 0 0-.615.593a3.94 3.94 0 0 0-.82 1.38c-.113.316-.196.716-.254 1.209a.682.682 0 1 0 1.354.158c.047-.397.11-.701.183-.905c.131-.364.3-.659.538-.915l.102-.11l.06-.062c.131-.133.304-.283 1.003-.89c.565-.49.988-1.234 1.084-1.925c.122-.872-.112-1.706-.838-2.436"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionCircleO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21m0 12.241a.91.91 0 1 1 0 1.819a.91.91 0 0 1 0-1.819m2.68-8.306c.726.73.96 1.564.838 2.436c-.096.691-.52 1.435-1.084 1.926c-.7.606-.872.756-1.004.889l-.06.062l-.101.11a2.595 2.595 0 0 0-.538.915c-.073.204-.136.508-.183.905a.682.682 0 0 1-1.354-.158c.058-.493.14-.893.255-1.21a3.94 3.94 0 0 1 .82-1.379c.17-.184.365-.37.614-.593c.115-.103.567-.494.678-.59c.313-.282.558-.718.607-1.066c.066-.471-.048-.876-.455-1.285c-.43-.432-1.106-.64-1.625-.572c-.758.098-1.065.21-1.588.668c-.382.336-.634.833-.75 1.519a.682.682 0 0 1-1.344-.227c.164-.98.561-1.76 1.194-2.316c.759-.667 1.31-.867 2.312-.997c.925-.12 2.027.218 2.768.963"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reddit(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M20 9.787c0-1.205-1.034-2.184-2.304-2.184c-.55 0-1.065.19-1.48.524l-.041.024c-1.499-.91-3.486-1.49-5.683-1.568l1.311-3.537l3.275.742c.002.982.844 1.78 1.88 1.78c1.038 0 1.882-.8 1.882-1.784S17.996 2 16.96 2c-.789 0-1.462.462-1.741 1.114l-3.88-.879l-1.61 4.34c-2.293.04-4.37.628-5.924 1.568a2.35 2.35 0 0 0-1.5-.54C1.032 7.602 0 8.582 0 9.786c0 .745.405 1.433 1.063 1.834a3.64 3.64 0 0 0-.066.661C.997 15.435 5.028 18 9.982 18c4.955 0 8.987-2.565 8.987-5.718c0-.217-.023-.431-.061-.642c.674-.398 1.092-1.096 1.092-1.853m-6.962 2.828c-.782 0-1.415-.6-1.415-1.342c0-.741.633-1.342 1.415-1.342s1.416.6 1.416 1.342c0 .742-.635 1.342-1.416 1.342m.281 2.289c-.042.04-1.057 1.02-3.353 1.02c-2.308 0-3.23-.992-3.27-1.035a.331.331 0 0 1 .04-.484a.378.378 0 0 1 .51.035c.02.022.79.8 2.72.8c1.963 0 2.824-.805 2.833-.813a.376.376 0 0 1 .51-.006a.33.33 0 0 1 .01.483m-7.607-3.631c0-.741.633-1.342 1.417-1.342c.78 0 1.415.6 1.415 1.342c0 .742-.635 1.342-1.415 1.342c-.783 0-1.417-.6-1.417-1.342"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reload(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.295 12a.704.704 0 0 1 .705.709v3.204a.704.704 0 0 1-.7.709a.704.704 0 0 1-.7-.709v-1.125C16.779 17.844 13.399 20 9.757 20c-4.41 0-8.106-2.721-9.709-6.915a.712.712 0 0 1 .4-.917c.36-.141.766.04.906.405c1.4 3.662 4.588 6.01 8.403 6.01c3.371 0 6.52-2.182 7.987-5.154l-1.471.01a.704.704 0 0 1-.705-.704a.705.705 0 0 1 .695-.714zm-9.05-12c4.408 0 8.105 2.721 9.708 6.915a.712.712 0 0 1-.4.917a.697.697 0 0 1-.906-.405c-1.4-3.662-4.588-6.01-8.403-6.01c-3.371 0-6.52 2.182-7.987 5.154l1.471-.01a.704.704 0 0 1 .705.704a.705.705 0 0 1-.695.714L.705 8A.704.704 0 0 1 0 7.291V4.087c0-.392.313-.709.7-.709c.386 0 .7.317.7.709v1.125C3.221 2.156 6.601 0 10.243 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Right(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7.053 2.158l7.243 7.256a.66.66 0 0 1 .204.483a.705.705 0 0 1-.204.497c-2.62 2.556-5.145 5.023-7.575 7.401c-.125.117-.625.408-1.011-.024c-.386-.433-.152-.81 0-.966l7.068-6.908l-6.747-6.759c-.246-.339-.226-.652.06-.939c.286-.287.607-.3.962-.04"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightCircle(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0M8.749 5.646a.682.682 0 1 0-.977.951l3.402 3.491l-3.397 3.42a.682.682 0 1 0 .967.96l3.87-3.894a.682.682 0 0 0 .005-.957Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightCircleO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21m-1.251 4.25l3.87 3.972a.682.682 0 0 1-.005.957l-3.87 3.895a.682.682 0 1 1-.967-.961l3.397-3.42l-3.402-3.49a.682.682 0 1 1 .977-.952"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightSquare(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.182 0C19.186 0 20 .814 20 1.818v16.364A1.818 1.818 0 0 1 18.182 20H1.818A1.818 1.818 0 0 1 0 18.182V1.818C0 .814.814 0 1.818 0zM8.749 5.627a.682.682 0 1 0-.977.951l3.402 3.491l-3.397 3.42a.682.682 0 0 0 .967.96l3.87-3.894a.682.682 0 0 0 .005-.957Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightSquareO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.818 1.364a.455.455 0 0 0-.454.454v16.364c0 .25.203.454.454.454h16.364a.455.455 0 0 0 .454-.454V1.818a.455.455 0 0 0-.454-.454zM18.182 0C19.186 0 20 .814 20 1.818v16.364A1.818 1.818 0 0 1 18.182 20H1.818A1.818 1.818 0 0 1 0 18.182V1.818C0 .814.814 0 1.818 0zM8.749 5.627a.682.682 0 1 0-.977.951l3.402 3.491l-3.397 3.42a.682.682 0 0 0 .967.96l3.87-3.894a.682.682 0 0 0 .005-.957Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rollback(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.31 4.324h11.12c1.502.01 2.889.52 4.158 1.535C19.492 7.379 20 9.77 20 10.644V19.3c-.085.467-.318.7-.7.7c-.382 0-.615-.233-.7-.7c0-3.132-.008-6.017-.024-8.656c0-1.298-1.29-4.754-5.147-4.922H2.473l2.964 2.966c.157.153.354.557.06.919a.7.7 0 0 1-1.05.07C4.266 9.494 2.84 8.067.168 5.401C.056 5.29 0 5.137 0 4.94s.068-.361.205-.495L4.491.158c.312-.231.627-.216.946.047c.318.262.318.592 0 .99z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Safari(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.591 9.966a.742.742 0 0 1-.184.503a.586.586 0 0 1-.463.212a.742.742 0 0 1-.503-.185a.586.586 0 0 1-.212-.463a.73.73 0 0 1 .19-.502a.602.602 0 0 1 .47-.212c.185 0 .349.062.49.184a.584.584 0 0 1 .212.463m.168.648l3.906-6.485c-.067.06-.318.292-.753.698c-.435.405-.902.839-1.4 1.3c-.5.461-1.007.934-1.524 1.417c-.518.484-.953.895-1.306 1.233c-.353.339-.541.53-.564.575l-3.895 6.474c.052-.052.301-.283.748-.692c.446-.41.915-.843 1.406-1.3l1.518-1.418c.52-.487.956-.9 1.306-1.239c.35-.338.535-.526.558-.563M17.98 10a7.787 7.787 0 0 1-1.16 4.14a18.429 18.429 0 0 0-.486-.306a.508.508 0 0 0-.185-.084c-.096 0-.145.048-.145.145c0 .075.22.238.66.491a7.989 7.989 0 0 1-2.06 2.127a7.817 7.817 0 0 1-2.696 1.233l-.179-.748c-.007-.074-.063-.112-.167-.112c-.037 0-.067.02-.09.062c-.022.04-.03.076-.022.106l.179.759a7.965 7.965 0 0 1-1.63.167a7.86 7.86 0 0 1-4.152-1.172c.008-.015.056-.091.145-.229c.09-.137.17-.262.24-.373c.071-.112.107-.183.107-.212c0-.097-.049-.145-.145-.145c-.045 0-.108.054-.19.162a3.6 3.6 0 0 0-.252.385a8.227 8.227 0 0 1-.15.256a8.015 8.015 0 0 1-2.143-2.092a7.813 7.813 0 0 1-1.228-2.729l.77-.167c.075-.023.112-.079.112-.168c0-.037-.02-.067-.061-.09a.17.17 0 0 0-.118-.022l-.758.168A8.127 8.127 0 0 1 2.019 10a7.78 7.78 0 0 1 1.217-4.23c.015.007.083.052.206.133c.122.082.234.153.335.212c.1.06.166.09.195.09c.097 0 .145-.045.145-.134c0-.045-.046-.102-.139-.173a4.02 4.02 0 0 0-.362-.24l-.224-.134a8.123 8.123 0 0 1 2.11-2.11A7.83 7.83 0 0 1 8.225 2.22l.167.748c.015.074.071.112.168.112c.037 0 .067-.02.09-.062c.022-.04.03-.08.022-.117l-.168-.737A8.327 8.327 0 0 1 10 2.02c1.518 0 2.928.405 4.23 1.216c-.29.417-.435.659-.435.726c0 .097.045.145.134.145c.081 0 .26-.238.535-.714a7.875 7.875 0 0 1 2.093 2.075a7.968 7.968 0 0 1 1.2 2.69l-.625.134c-.075.015-.112.074-.112.179c0 .037.02.067.061.089c.041.023.076.03.107.023l.636-.145A8.18 8.18 0 0 1 17.98 10m.949 0c0-1.213-.236-2.37-.709-3.471a8.988 8.988 0 0 0-1.903-2.846a8.988 8.988 0 0 0-2.846-1.904a8.713 8.713 0 0 0-3.472-.708a8.71 8.71 0 0 0-3.47.708a8.988 8.988 0 0 0-2.847 1.904a8.988 8.988 0 0 0-1.903 2.846A8.713 8.713 0 0 0 1.07 10c0 1.213.236 2.37.709 3.471a8.988 8.988 0 0 0 1.903 2.846a8.988 8.988 0 0 0 2.846 1.904a8.713 8.713 0 0 0 3.471.708c1.213 0 2.37-.236 3.472-.708a8.988 8.988 0 0 0 2.846-1.904a8.988 8.988 0 0 0 1.903-2.846A8.713 8.713 0 0 0 18.929 10M20 10a9.773 9.773 0 0 1-.793 3.884c-.528 1.235-1.238 2.299-2.131 3.192c-.893.893-1.957 1.603-3.192 2.131A9.773 9.773 0 0 1 10 20a9.773 9.773 0 0 1-3.884-.793c-1.235-.528-2.299-1.238-3.192-2.131c-.893-.893-1.603-1.957-2.132-3.192A9.773 9.773 0 0 1 0 10a9.78 9.78 0 0 1 .792-3.884c.529-1.235 1.24-2.299 2.132-3.192c.893-.893 1.957-1.603 3.192-2.132A9.773 9.773 0 0 1 10 0a9.78 9.78 0 0 1 3.884.792c1.235.529 2.299 1.24 3.192 2.132c.893.893 1.603 1.957 2.131 3.192A9.773 9.773 0 0 1 20 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Safety(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.842 3.865v7.434c0 4.047 3.205 7.327 7.158 7.327c3.953 0 7.158-3.28 7.158-7.327v-7.43L9.957 1.446zM9.955 0L18.5 2.874v8.425C18.5 16.104 14.694 20 10 20s-8.5-3.896-8.5-8.701V2.874zm4.419 6.485a.66.66 0 0 0-.948.057L9.11 11.511L6.322 8.426a.66.66 0 0 0-.948-.038a.698.698 0 0 0-.037.971l3.29 3.64a.66.66 0 0 0 .995-.01l4.807-5.534a.698.698 0 0 0-.055-.97"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Save(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.727 1.364c-.753 0-1.363.61-1.363 1.363v14.546c0 .753.61 1.363 1.363 1.363h14.546c.753 0 1.363-.61 1.363-1.363V2.727c0-.753-.61-1.363-1.363-1.363zM17.273 0A2.727 2.727 0 0 1 20 2.727v14.546A2.727 2.727 0 0 1 17.273 20H2.727A2.727 2.727 0 0 1 0 17.273V2.727A2.727 2.727 0 0 1 2.727 0zm-3.318 2.435a.682.682 0 0 0-.682.682V5a.682.682 0 0 0 1.364 0V3.117a.682.682 0 0 0-.682-.682M2.744.94l1.363.006l-.021 4.446c-.03.406.1.734.415 1.037c.313.301.744.438 1.384.4l8.625.004c.482-.055.821-.213 1.05-.469c.228-.257.346-.603.345-1.073V.942h1.364v4.347c.001.789-.227 1.46-.692 1.982c-.465.522-1.114.825-1.99.92h-8.66c-.959.059-1.765-.197-2.371-.78c-.604-.58-.89-1.304-.832-2.072z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.195 0c4.527 0 8.196 3.62 8.196 8.084a7.989 7.989 0 0 1-1.977 5.267l5.388 5.473a.686.686 0 0 1-.015.98a.71.71 0 0 1-.993-.014l-5.383-5.47a8.23 8.23 0 0 1-5.216 1.849C3.67 16.169 0 12.549 0 8.084C0 3.62 3.67 0 8.195 0m0 1.386c-3.75 0-6.79 2.999-6.79 6.698c0 3.7 3.04 6.699 6.79 6.699s6.791-3 6.791-6.699c0-3.7-3.04-6.698-6.79-6.698"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Setting(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.078 0c.294 0 .557.183.656.457l.706 1.957c.253.063.47.126.654.192c.201.072.46.181.78.33l1.644-.87a.702.702 0 0 1 .832.131l1.446 1.495c.192.199.246.49.138.744l-.771 1.807c.128.235.23.436.308.604c.084.183.188.435.312.76l1.797.77c.27.115.437.385.419.674l-.132 2.075a.69.69 0 0 1-.46.605l-1.702.605c-.049.235-.1.436-.154.606a8.79 8.79 0 0 1-.298.774l.855 1.89a.683.683 0 0 1-.168.793l-1.626 1.452a.703.703 0 0 1-.796.096l-1.676-.888a7.23 7.23 0 0 1-.81.367l-.732.274l-.65 1.8a.696.696 0 0 1-.64.457L9.11 20a.697.697 0 0 1-.669-.447l-.766-2.027a14.625 14.625 0 0 1-.776-.29a9.987 9.987 0 0 1-.618-.293l-1.9.812a.702.702 0 0 1-.755-.133L2.22 16.303a.683.683 0 0 1-.155-.783l.817-1.78a9.517 9.517 0 0 1-.302-.644a14.395 14.395 0 0 1-.3-.811L.49 11.74a.69.69 0 0 1-.49-.683l.07-1.921a.688.688 0 0 1 .392-.594L2.34 7.64c.087-.319.163-.567.23-.748a8.99 8.99 0 0 1 .314-.712L2.07 4.46a.683.683 0 0 1 .15-.79l1.404-1.326a.702.702 0 0 1 .75-.138l1.898.784c.21-.14.4-.253.572-.344c.205-.109.479-.223.824-.346l.66-1.841A.696.696 0 0 1 8.984 0zm-1.054 7.019c-1.667 0-3.018 1.335-3.018 2.983c0 1.648 1.351 2.984 3.018 2.984c1.666 0 3.017-1.336 3.017-2.984s-1.35-2.983-3.017-2.983"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SettingO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.078 0c.294 0 .557.183.656.457l.706 1.957c.253.063.47.126.654.192c.201.072.46.181.78.33l1.644-.87a.702.702 0 0 1 .832.131l1.446 1.495c.192.199.246.49.138.744l-.771 1.807c.128.235.23.436.308.604c.084.183.188.435.312.76l1.797.77c.27.115.437.385.419.674l-.132 2.075a.69.69 0 0 1-.46.605l-1.702.605c-.049.235-.1.436-.154.606a8.79 8.79 0 0 1-.298.774l.855 1.89a.683.683 0 0 1-.168.793l-1.626 1.452a.703.703 0 0 1-.796.096l-1.676-.888a7.23 7.23 0 0 1-.81.367l-.732.274l-.65 1.8a.696.696 0 0 1-.64.457L9.11 20a.697.697 0 0 1-.669-.447l-.766-2.027a14.625 14.625 0 0 1-.776-.29a9.987 9.987 0 0 1-.618-.293l-1.9.812a.702.702 0 0 1-.755-.133L2.22 16.303a.683.683 0 0 1-.155-.783l.817-1.78a9.517 9.517 0 0 1-.302-.644a14.395 14.395 0 0 1-.3-.811L.49 11.74a.69.69 0 0 1-.49-.683l.07-1.921a.688.688 0 0 1 .392-.594L2.34 7.64c.087-.319.163-.567.23-.748a8.99 8.99 0 0 1 .314-.712L2.07 4.46a.683.683 0 0 1 .15-.79l1.404-1.326a.702.702 0 0 1 .75-.138l1.898.784c.21-.14.4-.253.572-.344c.205-.109.479-.223.824-.346l.66-1.841A.696.696 0 0 1 8.984 0zm-.49 1.377H9.475L8.87 3.071a.693.693 0 0 1-.434.423c-.436.145-.751.27-.935.367c-.195.103-.444.26-.74.47a.703.703 0 0 1-.673.074l-1.83-.755l-.713.674l.743 1.57a.68.68 0 0 1-.006.597c-.2.401-.335.697-.403.879a10.31 10.31 0 0 0-.27.922a.69.69 0 0 1-.37.45l-1.79.859l-.036.98l1.62.492c.215.065.385.23.456.442c.16.48.288.834.38 1.056a10 10 0 0 0 .404.827a.68.68 0 0 1 .019.606l-.751 1.638l.711.668l1.782-.762a.703.703 0 0 1 .603.024c.365.192.637.325.809.398c.175.073.51.195.996.361a.693.693 0 0 1 .424.41l.708 1.871l.926-.02l.597-1.654a.692.692 0 0 1 .409-.413l1.037-.388c.262-.097.58-.25.951-.46a.703.703 0 0 1 .674-.008l1.577.835l.887-.791L15.856 14a.681.681 0 0 1-.001-.56c.182-.407.305-.714.367-.91c.061-.192.124-.469.185-.825a.69.69 0 0 1 .451-.533l1.648-.585l.072-1.14l-1.62-.694a.692.692 0 0 1-.377-.394a15.337 15.337 0 0 0-.378-.944a11.01 11.01 0 0 0-.42-.794a.682.682 0 0 1-.035-.606l.725-1.7l-.764-.79l-1.488.788a.703.703 0 0 1-.633.013a11.296 11.296 0 0 0-.968-.426a7.185 7.185 0 0 0-.857-.23a.694.694 0 0 1-.508-.441zm-.564 4.264c2.435 0 4.41 1.953 4.41 4.361c0 2.408-1.975 4.36-4.41 4.36c-2.436 0-4.41-1.952-4.41-4.36c0-2.408 1.974-4.36 4.41-4.36m0 1.378c-1.667 0-3.018 1.335-3.018 2.983c0 1.648 1.351 2.984 3.018 2.984c1.666 0 3.017-1.336 3.017-2.984s-1.35-2.983-3.017-2.983"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.815 0C17.575 0 19 1.45 19 3.235s-1.424 3.234-3.185 3.234a3.155 3.155 0 0 1-2.378-1.084l-6.284 3.44c.14.364.216.76.216 1.175c0 .542-.13 1.052-.363 1.501l6.008 3.725a3.177 3.177 0 0 1 2.801-1.695c1.76 0 3.185 1.45 3.185 3.234C19 18.55 17.576 20 15.815 20c-1.76 0-3.184-1.45-3.184-3.235l.003-.138l-6.53-4.046a3.138 3.138 0 0 1-1.92.654C2.425 13.235 1 11.785 1 10s1.424-3.235 3.185-3.235c.852 0 1.626.34 2.197.893l6.382-3.493a3.282 3.282 0 0 1-.133-.93C12.63 1.45 14.055 0 15.815 0m0 14.926c-.992 0-1.8.822-1.8 1.84c0 1.017.808 1.839 1.8 1.839c.993 0 1.8-.822 1.8-1.84c0-1.017-.807-1.839-1.8-1.839M4.185 8.161c-.993 0-1.8.822-1.8 1.839s.807 1.84 1.8 1.84c.992 0 1.8-.823 1.8-1.84c0-1.017-.808-1.84-1.8-1.84m11.63-6.766c-.992 0-1.8.822-1.8 1.84c0 1.017.808 1.839 1.8 1.839c.993 0 1.8-.822 1.8-1.84c0-1.017-.807-1.839-1.8-1.839"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCart(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.36 17.23c.752 0 1.362.62 1.362 1.386c0 .765-.61 1.386-1.361 1.386c-.752 0-1.361-.62-1.361-1.386c0-.765.61-1.385 1.36-1.385m15.423 0c.752 0 1.36.62 1.36 1.386c0 .765-.608 1.386-1.36 1.386c-.751 0-1.36-.62-1.36-1.386c0-.765.609-1.385 1.36-1.385M19.972.498a.695.695 0 0 1-.46.86c-.714.215-1.178.486-1.405.783c-.239.313-.37.696-.393 1.131v9.561c-.04 1.114-.297 1.992-.798 2.62c-.521.656-1.321 1.025-2.424 1.132H3.55c-.977 0-1.784-.255-2.396-.782c-.612-.527-.99-1.29-1.145-2.26L0 13.432v-5.94c.046-1.03.369-1.88.977-2.516c.619-.645 1.483-.975 2.573-1.005h9.518c.376 0 .68.31.68.692a.687.687 0 0 1-.68.693h-9.5c-.74.02-1.266.221-1.617.588c-.36.376-.56.901-.59 1.579v5.85c.11.634.336 1.081.672 1.37c.345.298.842.455 1.517.455l10.875.003c.701-.07 1.166-.284 1.434-.622c.29-.364.464-.96.494-1.773v-9.57c.038-.747.264-1.403.68-1.948c.429-.56 1.131-.97 2.094-1.26a.679.679 0 0 1 .845.468m-6.904 11.04c.376 0 .68.31.68.692a.687.687 0 0 1-.68.693H4.759a.687.687 0 0 1-.68-.693c0-.383.305-.693.68-.693zm0-3.696c.376 0 .68.31.68.693a.687.687 0 0 1-.68.693H4.759a.687.687 0 0 1-.68-.693c0-.382.305-.693.68-.693z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shrink(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.672 10.561a.7.7 0 0 1 .7.7v6a.7.7 0 1 1-1.4 0l-.001-4.354l-6.772 6.889a.7.7 0 1 1-.998-.982l6.737-6.854l-4.266.001a.7.7 0 1 1 0-1.4zM19.792.201a.7.7 0 0 1 .009.99l-6.737 6.854l4.266-.001a.7.7 0 0 1 0 1.4h-6a.7.7 0 0 1-.7-.7v-6a.7.7 0 1 1 1.4 0v4.354L18.804.209a.7.7 0 0 1 .99-.008"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Smile(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0M7.086 11.399l-.1.004a.682.682 0 0 0-.596.759a3.637 3.637 0 0 0 7.217.024a.682.682 0 0 0-1.352-.172A2.273 2.273 0 0 1 7.744 12a.682.682 0 0 0-.759-.596Zm-1.272-5.12a1.395 1.395 0 1 0 0 2.79a1.395 1.395 0 0 0 0-2.79m8.372 0a1.395 1.395 0 1 0 0 2.79a1.395 1.395 0 0 0 0-2.79"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmileO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21M7.744 12a2.273 2.273 0 0 0 4.51.015a.682.682 0 1 1 1.353.172a3.637 3.637 0 0 1-7.217-.024A.682.682 0 1 1 7.744 12m-1.93-5.72a1.395 1.395 0 1 1 0 2.79a1.395 1.395 0 0 1 0-2.79m8.372 0a1.395 1.395 0 1 1 0 2.79a1.395 1.395 0 0 1 0-2.79"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.818 1.364a.455.455 0 0 0-.454.454v16.364c0 .25.203.454.454.454h16.364a.455.455 0 0 0 .454-.454V1.818a.455.455 0 0 0-.454-.454zm0-1.364h16.364C19.186 0 20 .814 20 1.818v16.364A1.818 1.818 0 0 1 18.182 20H1.818A1.818 1.818 0 0 1 0 18.182V1.818C0 .814.814 0 1.818 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarOff(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.902 12.937L.429 9.127l-.058-.074C0 8.513-.107 7.943.116 7.406c.226-.545.737-.87 1.472-1.014l4.543-.715L8.675.752C8.965.279 9.39 0 9.908 0c.51 0 .936.27 1.242.73l.05.088l2.285 4.863l5.168.827c.705.127 1.198.508 1.32 1.138c.11.563-.118 1.113-.622 1.667l-3.593 3.637l.86 5.277c.086.638-.084 1.181-.549 1.523c-.477.35-1.087.318-1.863-.008l-4.38-2.355l-4.318 2.315c-.607.333-1.22.354-1.74.004s-.73-.925-.644-1.627zm1.454-.441l-.049.323l-.824 5.446c-.025.206-.002.268.051.304c.055.036.126.034.32-.072l4.973-2.667l.325.174l4.643 2.5c.317.133.462.14.462.14c.013-.009.026-.052.004-.215l-.969-5.946l.245-.248l3.818-3.864c.22-.244.284-.397.271-.464c0 0-.018-.014-.203-.047l-5.872-.94l-2.57-5.47c-.049-.069-.065-.079-.073-.079c0 0-.007.004-.039.053l-2.838 5.5l-.344.054l-4.86.764c-.302.06-.421.136-.444.19c-.025.06-.011.146.095.31z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarOn(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.89 17.514l-4.21 2.257l-.099.044c-.715.27-1.39.216-1.903-.242c-.482-.43-.657-1.046-.557-1.755l.704-4.86l-3.18-3.342c-.55-.56-.765-1.248-.58-1.968c.205-.799.88-1.258 1.851-1.412l4.227-.638l2.213-4.585C8.7.366 9.236-.017 9.911.001c.66.017 1.183.422 1.593 1.143l2.14 4.486l4.74.658c.753.13 1.308.522 1.53 1.176c.22.653.01 1.313-.557 1.987l-3.44 3.51l.772 4.856c.122.84-.025 1.505-.586 1.9c-.506.357-1.139.357-1.867.107l-.12-.053z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stop(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0C4.478 0 0 4.478 0 10s4.478 10 10 10s10-4.478 10-10S15.522 0 10 0m5.241 16.44L3.561 4.759c.356-.44.758-.842 1.198-1.199l11.68 11.681a8.335 8.335 0 0 1-1.198 1.199"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0C4.478 0 0 4.478 0 10s4.478 10 10 10s10-4.478 10-10S15.522 0 10 0m0 18.304A8.305 8.305 0 0 1 3.56 4.759l11.681 11.68A8.266 8.266 0 0 1 10 18.305m6.44-3.063L4.759 3.561a8.305 8.305 0 0 1 11.68 11.68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Swap(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.299 13.098a.7.7 0 0 1 .498 1.191l-5.427 5.503a.7.7 0 1 1-.996-.984l4.252-4.31H.703a.7.7 0 0 1 0-1.4zM6.619.202a.7.7 0 0 1 .007.99l-4.252 4.31h16.923a.7.7 0 0 1 0 1.4H.7a.7.7 0 0 1-.498-1.191L5.63.208a.7.7 0 0 1 .99-.006"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwapLeft(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.64 11.917h16.591a.78.78 0 0 1 .769.792a.78.78 0 0 1-.769.791H.771c-.688 0-1.03-.857-.541-1.354L5.549 6.73a.754.754 0 0 1 1.087.006a.808.808 0 0 1-.005 1.119z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwapRight(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M17.625 12.08H.7c-.176 0-.7.092-.7.71s.524.71.7.71h18.599c.406 0 .701-.287.701-.71a.665.665 0 0 0-.164-.453a4.428 4.428 0 0 0-.173-.187L14.34 6.68c-.348-.278-.668-.27-.96.025c-.292.294-.311.61-.058.95z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Table(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.364 5.138v12.02h17.272V5.138zM.909 1.5h18.182c.502 0 .909.4.909.895v15.21a.902.902 0 0 1-.91.895H.91c-.503 0-.91-.4-.91-.895V2.395C0 1.9.407 1.5.91 1.5m5.227 1.759c0-.37.306-.671.682-.671c.377 0 .682.3.682.671v13.899c0 .37-.305.67-.682.67a.676.676 0 0 1-.682-.67zm6.96-.64c.377 0 .682.3.682.67v4.995h4.91c.377 0 .683.301.683.672c0 .37-.306.671-.682.671l-4.911-.001v3.062h5.002c.377 0 .682.3.682.671c0 .37-.305.671-.682.671h-5.002v3.158a.676.676 0 0 1-.682.671a.676.676 0 0 1-.681-.67l-.001-3.159H1.001a.676.676 0 0 1-.682-.67c0-.371.305-.672.682-.672h11.413V9.626L.909 9.627a.676.676 0 0 1-.682-.671c0-.37.306-.671.682-.671l11.505-.001V3.289c0-.37.306-.67.682-.67"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.368 0a.68.68 0 0 1 .484.201l8.35 8.39c.558.586.836 1.249.794 1.958c-.04.664-.304 1.274-.807 1.846l-7.074 7.037l-.115.092c-.638.406-1.26.564-1.85.428c-.515-.118-1.054-.436-1.677-.978l-8.16-8.219a.68.68 0 0 1-.199-.47L0 1.472C.007 1.044.126.681.392.413C.666.138 1.055.023 1.588 0zM6.473 4.574a1.59 1.59 0 1 0 0 3.18a1.59 1.59 0 1 0 0-3.18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.368 0a.68.68 0 0 1 .484.201l8.35 8.39c.558.586.836 1.249.794 1.958c-.04.664-.304 1.274-.807 1.846l-7.074 7.037l-.115.092c-.638.406-1.26.564-1.85.428c-.515-.118-1.054-.436-1.677-.978l-8.16-8.219a.68.68 0 0 1-.199-.47L0 1.472C.007 1.044.126.681.392.413C.666.138 1.055.023 1.588 0zm-.284 1.363H1.618c-.13.005-.21.018-.243.02l-.006-.001v.015l-.004.078l.111 8.516l7.929 7.988c.431.374.789.585 1.052.645c.178.041.424-.017.75-.213l6.986-6.948c.278-.316.419-.642.44-.995c.018-.307-.109-.61-.41-.927zM6.474 3.21a2.956 2.956 0 0 1 2.958 2.953a2.956 2.956 0 0 1-2.959 2.954a2.956 2.956 0 0 1-2.958-2.954A2.956 2.956 0 0 1 6.473 3.21m0 1.363a1.59 1.59 0 1 0-.001 3.18a1.59 1.59 0 1 0 0-3.18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tags(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.258 2.006c.18 0 .353.07.48.194l9.015 8.798c.187.22.273.492.24.775c-.027.234-.124.449-.322.673l-5.442 5.236a1.123 1.123 0 0 1-.782.318c-.276 0-.507-.098-.794-.318l-.17-.135l.941-.942l4.437-4.265c.168-.194.251-.426.251-.697a.834.834 0 0 0-.275-.652l-8.21-7.976c-.396-.37-.805-.631-1.226-.783a4.629 4.629 0 0 0-.724-.193l-.131-.023c-.035-.007-.036-.009.007-.01zM6.817 2c.18 0 .353.07.48.194l9.015 8.798c.186.22.273.491.24.774c-.027.234-.124.45-.322.673l-5.443 5.236a1.123 1.123 0 0 1-.782.319c-.275 0-.532-.098-.818-.318L.205 9.166A.648.648 0 0 1 0 8.696V3.358c-.01-.334.076-.639.27-.89C.5 2.172.848 2.024 1.299 2zM4.49 4.648c-.88 0-1.593.696-1.593 1.555c0 .858.713 1.554 1.593 1.554s1.593-.696 1.593-1.554c0-.859-.713-1.555-1.593-1.555"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagsO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.205 9.166A.648.648 0 0 1 0 8.696V3.358c-.01-.334.076-.639.27-.89C.5 2.172.848 2.024 1.299 2h5.518c.18 0 .353.07.48.194l9.015 8.798c.186.22.273.491.24.774c-.027.234-.124.45-.322.673l-5.443 5.236a1.123 1.123 0 0 1-.782.319c-.275 0-.532-.098-.818-.318zm1.15-.747l8.628 8.18l5.122-4.925l-8.57-8.36h-5.18zM9.976 3.32h-2.99a.667.667 0 0 1-.677-.657c0-.363.304-.658.678-.658h3.271c.18 0 .353.07.48.194l9.015 8.798c.187.22.273.492.24.775c-.027.234-.124.449-.322.673l-5.442 5.236a1.123 1.123 0 0 1-.782.318c-.276 0-.532-.097-.819-.318l-1.341-1.271a.643.643 0 0 1-.012-.93a.692.692 0 0 1 .958-.01l1.191 1.134l5.123-4.924zm-5.48 4.845c-1.122 0-2.032-.882-2.032-1.97c0-1.09.91-1.972 2.032-1.972c1.122 0 2.031.882 2.031 1.971s-.91 1.971-2.031 1.971m0-1.314a.667.667 0 0 0 .677-.657a.667.667 0 0 0-.677-.657a.667.667 0 0 0-.678.657c0 .363.304.657.678.657"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Taobao(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.977 6.56L.841 8.262L2.937 9.53s1.394.696.726 2.004C3.045 12.77 0 15.478 0 15.478l2.726 1.66c1.89-3.996 1.762-3.463 2.235-4.901c.487-1.46.595-2.581-.23-3.397C3.67 7.8 3.55 7.704 1.976 6.56M7.111 2.5l2.515.693s-.204.473-.638 1.207c10.08-2.787 10.674 1.713 10.674 1.713l.01.045c.077.398.614 3.41.125 7.949c-.53 4.88-7.023 3.062-7.023 3.062l.35-1.397l1.505.316c2.78.168 2.51-2.201 2.51-2.201V6.981c.02-2.622-2.526-2.898-7.095-1.296l1.068.283c-.093.302-.434.782-.876 1.307h6.105v1.207h-3.433v1.512h3.421v1.211h-3.42v2.53c.51-.168.99-.392 1.397-.697l-.3-1.1l1.613-.495l1.348 3.219l-1.989.816l-.353-1.29c-.891.667-2.742 1.636-5.967 1.55c-3.452.086-2.56-3.8-2.56-3.8l.087-.048h2.423c-.011.503-.222 1.312.062 1.759c.238.365.837.428 1.224.447l.135.003v-2.894h-3.51V9.998h3.51V8.486H9.13c-.787.816-1.512 1.497-1.512 1.497l-1.06-.909C7.31 8.3 8.059 7.07 8.528 6.251a38.27 38.27 0 0 0-1.168.484c-.384.484-.837.984-1.34 1.475c.02.03-1.739-.96-1.739-.96c1.7-1.415 2.7-4.352 2.818-4.713zm-3.644.283c.99 0 1.793.7 1.793 1.572c0 .864-.802 1.564-1.793 1.564c-.998 0-1.793-.704-1.797-1.564c0-.868.799-1.572 1.797-1.572"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Time(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m-.93 5.581a.698.698 0 0 0-.698.698v5.581c0 .386.312.698.698.698h5.581a.698.698 0 1 0 0-1.395H9.767V6.279a.698.698 0 0 0-.697-.698"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21m-.93 4.186c.385 0 .697.313.697.698v4.884h4.884a.698.698 0 0 1 0 1.395H9.07a.698.698 0 0 1-.698-.698V6.28c0-.386.312-.699.698-.699"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitter(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M20 3.894a8.299 8.299 0 0 1-2.357.636a4.062 4.062 0 0 0 1.804-2.234a8.298 8.298 0 0 1-2.605.98A4.13 4.13 0 0 0 13.847 2c-2.266 0-4.103 1.808-4.103 4.04c0 .316.036.624.106.92a11.71 11.71 0 0 1-8.457-4.22a3.974 3.974 0 0 0-.556 2.03a4.02 4.02 0 0 0 1.826 3.362a4.143 4.143 0 0 1-1.859-.505v.05c0 1.957 1.414 3.59 3.29 3.961a4.189 4.189 0 0 1-1.852.07c.522 1.604 2.037 2.772 3.833 2.805a8.317 8.317 0 0 1-5.096 1.73A8.42 8.42 0 0 1 0 16.185A11.747 11.747 0 0 0 6.29 18c7.547 0 11.674-6.155 11.674-11.492c0-.175-.004-.35-.012-.523A8.249 8.249 0 0 0 20 3.895"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Uiw(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.976 0l4.475 3.418l-1.71 5.531H3.21L1.5 3.42zm0 20L1.5 16.582l1.71-5.531h5.532l1.709 5.53zM18.5 12.968l-5.261 1.797l-3.252-4.705l3.252-4.705L18.5 7.152z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlock(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.615 2.046a9.878 9.878 0 0 1 2.322 3.076a.7.7 0 1 1-1.262.606a8.479 8.479 0 0 0-1.998-2.643c-1.304-1.177-2.786-1.75-4.43-1.679c-2.22.097-3.644.63-5.159 2.225C3.947 4.833 3.391 6.196 3.4 7.76v1.245h14.607v10a1 1 0 0 1-1 1h-14a1 1 0 0 1-1-1v-9.25A.706.706 0 0 1 2 9.66V7.763c-.011-1.926.686-3.636 2.073-5.096C5.865.781 7.631.118 10.186.007c2.024-.088 3.86.623 5.43 2.039m.992 8.359h-13.2v8.2h13.2zm-6.245 1.223a.7.7 0 0 1 .7.7v2.169a1.5 1.5 0 1 1-1.4.006v-2.175a.7.7 0 0 1 .7-.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Up(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.103 7.222c3.447 3.468 5.744 5.764 6.89 6.887c.198.185.539.56 1.046.07c.339-.327.325-.685-.039-1.073l-7.444-7.43a.638.638 0 0 0-.455-.176a.702.702 0 0 0-.472.176l-7.453 7.635c-.241.388-.231.715.03.98c.26.265.577.28.95.043z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpCircle(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m-.48 7.636l-3.896 3.87a.682.682 0 0 0 .962.968l3.419-3.398l3.49 3.402a.682.682 0 0 0 .952-.976l-3.971-3.87a.682.682 0 0 0-.957.004"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpCircleO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0m0 1.395a8.605 8.605 0 1 0 0 17.21a8.605 8.605 0 0 0 0-17.21m-.48 6.241a.682.682 0 0 1 .956-.005l3.971 3.87a.682.682 0 1 1-.951.977l-3.491-3.402l-3.42 3.398a.682.682 0 1 1-.96-.968Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpSquare(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.182 0C19.186 0 20 .814 20 1.818v16.364A1.818 1.818 0 0 1 18.182 20H1.818A1.818 1.818 0 0 1 0 18.182V1.818C0 .814.814 0 1.818 0zM9.519 7.617l-3.895 3.87a.682.682 0 0 0 .962.968l3.419-3.398l3.49 3.402a.682.682 0 1 0 .952-.976l-3.971-3.87a.682.682 0 0 0-.957.004"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpSquareO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.818 1.364a.455.455 0 0 0-.454.454v16.364c0 .25.203.454.454.454h16.364a.455.455 0 0 0 .454-.454V1.818a.455.455 0 0 0-.454-.454zM18.182 0C19.186 0 20 .814 20 1.818v16.364A1.818 1.818 0 0 1 18.182 20H1.818A1.818 1.818 0 0 1 0 18.182V1.818C0 .814.814 0 1.818 0zM9.519 7.617l-3.895 3.87a.682.682 0 0 0 .962.968l3.419-3.398l3.49 3.402a.682.682 0 1 0 .952-.976l-3.971-3.87a.682.682 0 0 0-.957.004"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Upload(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.31 12.051c.381 0 .69.314.69.7v4.918c-.006.67-.127 1.2-.399 1.594c-.328.476-.908.692-1.747.737l-15.903-.002c-.646-.046-1.168-.302-1.507-.777c-.302-.423-.446-.95-.444-1.558V12.75c0-.386.309-.7.69-.7c.38 0 .688.314.688.7v4.913c0 .333.065.572.182.736c.081.114.224.184.44.201l15.817.001c.42-.023.627-.1.655-.14c.084-.123.146-.393.15-.8V12.75c0-.386.308-.7.689-.7M9.99 0c.224 0 .423.108.549.276l4.281 4.308c.27.272.272.715.004.99a.682.682 0 0 1-.974.003l-3.171-3.189v9.643c0 .387-.308.7-.689.7a.694.694 0 0 1-.69-.7V2.383L6.172 5.574a.682.682 0 0 1-.89.076l-.085-.074a.707.707 0 0 1-.002-.989L9.49.207a.682.682 0 0 1 .404-.202h.011A.462.462 0 0 1 9.977 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func User(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.993 10.573a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9M10 0a6 6 0 0 1 3.04 11.174c3.688 1.11 6.458 4.218 6.955 8.078c.047.367-.226.7-.61.745c-.383.045-.733-.215-.78-.582c-.54-4.19-4.169-7.345-8.57-7.345c-4.425 0-8.101 3.161-8.64 7.345c-.047.367-.397.627-.78.582c-.384-.045-.657-.378-.61-.745c.496-3.844 3.281-6.948 6.975-8.068A6 6 0 0 1 10 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserAdd(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.467 0c3.184 0 5.765 2.566 5.765 5.732a5.706 5.706 0 0 1-2.02 4.358c.27.134.516.268.738.403c.478.29.974.65 1.49 1.079a.685.685 0 0 1 .086.969a.694.694 0 0 1-.975.086a11.05 11.05 0 0 0-1.322-.96a11.403 11.403 0 0 0-1.405-.703a5.76 5.76 0 0 1-2.357.5a5.767 5.767 0 0 1-2.58-.605l-.042.02c-1.95.756-3.373 1.874-4.292 3.358c-.922 1.489-1.299 3.153-1.13 5.014a.689.689 0 0 1-.628.746a.69.69 0 0 1-.75-.623c-.195-2.152.249-4.113 1.33-5.858c.95-1.536 2.347-2.73 4.174-3.582a5.694 5.694 0 0 1-1.846-4.202C3.703 2.566 6.283 0 9.467 0m7.401 12.693c.38 0 .688.31.688.691v1.752h1.752a.69.69 0 0 1 .692.689a.69.69 0 0 1-.692.687h-1.752v1.753a.69.69 0 0 1-.688.691a.69.69 0 0 1-.688-.691v-1.753h-1.752a.69.69 0 0 1-.692-.687c0-.38.31-.688.692-.688l1.752-.001v-1.752a.69.69 0 0 1 .688-.691m-7.4-11.317c-2.42 0-4.382 1.95-4.382 4.356c0 2.406 1.962 4.357 4.381 4.357c2.42 0 4.381-1.95 4.381-4.357c0-2.406-1.961-4.356-4.38-4.356"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserDelete(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.596 0c3.226 0 5.842 2.566 5.842 5.732a5.676 5.676 0 0 1-2.047 4.359c.274.133.523.267.748.402c.485.29.988.65 1.51 1.079a.679.679 0 0 1 .088.969a.71.71 0 0 1-.988.086a11.197 11.197 0 0 0-1.34-.96a11.62 11.62 0 0 0-1.424-.703a5.91 5.91 0 0 1-2.39.5a5.91 5.91 0 0 1-2.615-.605l-.042.02c-1.977.756-3.42 1.874-4.35 3.358c-.935 1.489-1.317 3.153-1.146 5.014a.692.692 0 0 1-.636.746a.697.697 0 0 1-.761-.623c-.197-2.152.253-4.113 1.348-5.858c.964-1.536 2.38-2.73 4.23-3.581a5.664 5.664 0 0 1-1.87-4.203C4.753 2.566 7.37 0 10.596 0m7.223 14.24a.698.698 0 0 1 .978-.003c.27.266.27.698.002.965l-1.192 1.179l1.192 1.18a.675.675 0 0 1-.002.964a.698.698 0 0 1-.978-.001l-1.187-1.177l-1.187 1.177a.698.698 0 0 1-.891.073l-.086-.072a.675.675 0 0 1-.002-.964l1.19-1.18l-1.19-1.18a.675.675 0 0 1 .002-.964a.698.698 0 0 1 .977.002l1.187 1.176ZM10.596 1.375c-2.453 0-4.44 1.95-4.44 4.356c0 2.406 1.987 4.357 4.44 4.357c2.452 0 4.44-1.95 4.44-4.357c0-2.406-1.988-4.356-4.44-4.356"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UsergroupAdd(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.346 11.101c.88 0 1.741.146 2.555.428a.681.681 0 0 1 .426.873a.7.7 0 0 1-.89.418a6.372 6.372 0 0 0-2.09-.35c-3.467 0-6.276 2.757-6.276 6.159c0 .208.01.414.03.619a.688.688 0 0 1-.624.749a.694.694 0 0 1-.763-.614a7.469 7.469 0 0 1-.038-.754c0-4.157 3.434-7.528 7.67-7.528m4.146 1.369a.707.707 0 0 1 .7.694v2.259l2.114-.005a.707.707 0 0 1 .686.599l.008.102a.675.675 0 0 1-.688.68l-2.12.005v2.37a.674.674 0 0 1-.68.688a.707.707 0 0 1-.7-.694l-.001-2.361l-2.51.007a.707.707 0 0 1-.685-.598l-.008-.103a.675.675 0 0 1 .687-.68l2.516-.007v-2.269a.675.675 0 0 1 .681-.687M7.671 8.22a.69.69 0 0 1 .697.684a.69.69 0 0 1-.697.685c-3.467 0-6.276 2.757-6.276 6.159c0 .207.01.414.03.618a.688.688 0 0 1-.624.75a.694.694 0 0 1-.763-.614A7.469 7.469 0 0 1 0 15.748C0 11.59 3.434 8.22 7.67 8.22m4.938-5.254c2.183 0 3.952 1.737 3.952 3.878c0 2.142-1.77 3.878-3.952 3.878s-3.95-1.736-3.95-3.878c0-2.141 1.769-3.878 3.951-3.878m0 1.369c-1.412 0-2.557 1.123-2.557 2.51c0 1.385 1.145 2.508 2.557 2.508s2.557-1.123 2.557-2.509s-1.145-2.509-2.557-2.509M8.025 0c1.245 0 2.395.57 3.138 1.52c.234.3.176.73-.13.96a.706.706 0 0 1-.977-.127a2.568 2.568 0 0 0-2.031-.984c-1.412 0-2.557 1.123-2.557 2.509c0 1.212.882 2.245 2.081 2.466c.378.07.628.427.557.799a.697.697 0 0 1-.815.546c-1.855-.342-3.218-1.938-3.218-3.811C4.073 1.736 5.843 0 8.025 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UsergroupDelete(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.442 11.101c.869 0 1.718.146 2.521.428a.683.683 0 0 1 .42.873a.689.689 0 0 1-.878.418a6.212 6.212 0 0 0-2.063-.35c-3.42 0-6.193 2.757-6.193 6.159c0 .208.01.414.031.619c.038.376-.198.71-.604.75a.682.682 0 0 1-.765-.615a7.568 7.568 0 0 1-.038-.754c0-4.157 3.389-7.528 7.569-7.528m6.384 2.685a.69.69 0 0 1 .973.005a.682.682 0 0 1-.005.968l-1.634 1.608l1.634 1.609c.27.266.272.7.005.968a.69.69 0 0 1-.973.004l-1.644-1.618l-1.641 1.618a.691.691 0 0 1-.889.068l-.084-.072a.682.682 0 0 1 .004-.968l1.633-1.609l-1.633-1.608a.682.682 0 0 1-.004-.968a.69.69 0 0 1 .973-.005l1.642 1.618ZM7.569 8.22c.38 0 .688.306.688.684a.686.686 0 0 1-.688.685c-3.42 0-6.193 2.757-6.193 6.159c0 .207.01.414.031.618a.685.685 0 0 1-.617.75a.687.687 0 0 1-.752-.614A7.568 7.568 0 0 1 0 15.748C0 11.59 3.389 8.22 7.569 8.22M12.7 2.966c2.154 0 3.9 1.737 3.9 3.878c0 2.142-1.746 3.878-3.9 3.878c-2.153 0-3.899-1.736-3.899-3.878c0-2.141 1.746-3.878 3.9-3.878m0 1.369a2.516 2.516 0 0 0-2.522 2.51a2.516 2.516 0 0 0 2.522 2.508a2.516 2.516 0 0 0 2.523-2.509a2.516 2.516 0 0 0-2.523-2.509M7.918 0a3.9 3.9 0 0 1 3.097 1.52c.23.3.174.73-.128.96a.69.69 0 0 1-.965-.127a2.522 2.522 0 0 0-2.004-.984a2.516 2.516 0 0 0-2.523 2.509c0 1.212.871 2.245 2.054 2.466a.685.685 0 1 1-.254 1.345a3.884 3.884 0 0 1-3.176-3.811C4.02 1.736 5.765 0 7.92 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Verification(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.727 3.333c-.753 0-1.363.597-1.363 1.334v10.666c0 .737.61 1.334 1.363 1.334h14.546c.753 0 1.363-.597 1.363-1.334V4.667c0-.737-.61-1.334-1.363-1.334zM17.273 2C18.779 2 20 3.194 20 4.667v10.666C20 16.806 18.779 18 17.273 18H2.727C1.221 18 0 16.806 0 15.333V4.667C0 3.194 1.221 2 2.727 2zM8.167 13.133H4a.674.674 0 0 0-.682.667c0 .368.305.667.682.667h4.167a.674.674 0 0 0 .681-.667a.674.674 0 0 0-.681-.667m5.583-7.8c-1.524 0-2.765 1.191-2.765 2.667c0 .773.34 1.468.884 1.955c-1.03.608-1.717 1.703-1.717 2.954c0 .357.056.708.165 1.043c.115.35.499.544.857.432a.664.664 0 0 0 .442-.838a2.043 2.043 0 0 1-.1-.637c0-1.175.997-2.133 2.234-2.133s2.235.958 2.235 2.133c0 .173-.021.342-.063.506a.666.666 0 0 0 .497.808a.683.683 0 0 0 .826-.486c.069-.27.103-.547.103-.828c0-1.251-.687-2.346-1.717-2.956c.544-.485.884-1.18.884-1.953c0-1.476-1.24-2.667-2.765-2.667m-5.583 4.3H4a.674.674 0 0 0-.682.667c0 .368.305.667.682.667h4.167a.674.674 0 0 0 .681-.667a.674.674 0 0 0-.681-.667m5.583-2.966c.777 0 1.402.6 1.402 1.333c0 .734-.625 1.333-1.402 1.333c-.777 0-1.402-.6-1.402-1.333c0-.734.625-1.333 1.402-1.333m-5.583-.534H4a.674.674 0 0 0-.682.667c0 .368.305.667.682.667h4.167a.674.674 0 0 0 .681-.667a.674.674 0 0 0-.681-.667"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerticleLeft(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.886 2.203a.715.715 0 0 1 1-.012a.69.69 0 0 1 .012.986L8.272 9.855l6.992 6.876a.69.69 0 0 1 0 .986a.715.715 0 0 1-1 0L6.893 10.47v6.839a.695.695 0 0 1-.696.692a.695.695 0 0 1-.697-.692V2.692c0-.382.312-.692.697-.692c.385 0 .697.31.697.692L6.893 9.25Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerticleRight(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.606 9.256L6.614 2.203c-.367-.292-.704-.292-1.012 0c-.308.291-.308.616 0 .974l6.626 6.678l-6.992 6.876c-.313.304-.321.625-.023.966c.299.34.64.347 1.023.02l7.37-7.245v6.836c0 .461.233.692.697.692c.465 0 .697-.23.697-.692V2.692c0-.466-.232-.697-.697-.692c-.464.005-.697.236-.697.692z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoCamera(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.823 10.146a.457.457 0 0 0-.456.458v7.566c0 .253.204.458.456.458H14.99a.457.457 0 0 0 .455-.458v-7.566a.457.457 0 0 0-.455-.458zM14.99 8.774c1.006 0 1.822.819 1.822 1.83v7.566c0 1.01-.816 1.83-1.822 1.83H1.823A1.826 1.826 0 0 1 0 18.17v-7.566c0-1.011.816-1.83 1.823-1.83zM20 12.69v3.482c0 .435-.397.76-.821.672l-1.433-.296v-1.402l.887.184v-1.792l-.887.189v-1.403l1.428-.305a.684.684 0 0 1 .826.67m-14.987.447c-1.007 0-1.823.82-1.823 1.83c0 1.01.816 1.83 1.823 1.83c1.007 0 1.823-.82 1.823-1.83c0-1.01-.816-1.83-1.823-1.83m0 .915c.503 0 .911.41.911.915a.913.913 0 0 1-.911.915a.913.913 0 0 1-.912-.915c0-.506.408-.915.912-.915M12.699.1a3.881 3.881 0 0 1 3.874 3.888a3.881 3.881 0 0 1-3.874 3.889a3.881 3.881 0 0 1-3.873-3.889A3.881 3.881 0 0 1 12.699.1M4.101 0a3.881 3.881 0 0 1 3.874 3.889A3.881 3.881 0 0 1 4.1 7.777A3.881 3.881 0 0 1 .228 3.89A3.881 3.881 0 0 1 4.1 0m8.6 1.472a2.511 2.511 0 0 0-2.506 2.516A2.511 2.511 0 0 0 12.7 6.504a2.511 2.511 0 0 0 2.507-2.516A2.511 2.511 0 0 0 12.7 1.472m-8.598-.1A2.511 2.511 0 0 0 1.595 3.89a2.511 2.511 0 0 0 2.506 2.516A2.511 2.511 0 0 0 6.608 3.89A2.511 2.511 0 0 0 4.1 1.372"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Warning(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.59 15.86L12.007 1.924C11.515 1.011 10.779.5 9.989.5c-.79 0-1.515.521-2.016 1.434L.409 15.861c-.49.901-.544 1.825-.138 2.53c.405.707 1.216 1.109 2.219 1.109h15.02c1.003 0 1.814-.402 2.22-1.108c.405-.706.351-1.619-.14-2.531M10 4.857c.395 0 .715.326.715.728v6.583c0 .402-.32.728-.715.728a.721.721 0 0 1-.715-.728V5.584c0-.391.32-.728.715-.728m0 11.624c-.619 0-1.11-.51-1.11-1.14c0-.63.502-1.141 1.11-1.141c.619 0 1.11.51 1.11 1.14c0 .63-.502 1.141-1.11 1.141"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WarningO(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.51 19.5H2.49c-1.007 0-1.816-.405-2.22-1.111c-.404-.707-.355-1.63.136-2.53l7.577-13.93C8.477 1.02 9.211.5 9.995.5c.785 0 1.518.52 2.012 1.427l7.586 13.933c.49.901.54 1.823.136 2.53c-.404.706-1.213 1.11-2.22 1.11M9.995 1.962c-.25 0-.535.254-.767.679L1.652 16.57c-.238.438-.293.83-.151 1.078c.142.248.503.39.99.39H17.51c.488 0 .85-.142.991-.39c.141-.247.086-.64-.152-1.076L10.763 2.64c-.23-.425-.518-.679-.767-.679M10 12.904a.722.722 0 0 1-.712-.731v-6.58c0-.404.319-.732.712-.732c.394 0 .713.328.713.731v6.58c0 .405-.32.732-.713.732m-.005 3.574c.614 0 1.11-.511 1.11-1.142c0-.63-.496-1.14-1.11-1.14c-.613 0-1.11.51-1.11 1.14c0 .63.497 1.142 1.11 1.142"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Weibo(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.812 9.801c-.778-.141-.4-.534-.4-.534s.761-1.178-.15-2.034c-1.13-1.061-3.877.135-3.877.135c-1.05.306-.77-.14-.622-.897c0-.892-.326-2.402-3.12-1.51C3.853 5.858 1.455 9 1.455 9C-.212 11.087.01 12.7.01 12.7c.416 3.562 4.448 4.54 7.584 4.771c3.299.243 7.752-1.067 9.102-3.76c1.35-2.696-1.104-3.763-1.884-3.91m-1.044 2.549c0 2.051-2.653 3.977-5.93 4.117c-3.276.144-5.923-1.398-5.923-3.45c0-2.054 2.647-3.7 5.923-3.842c3.277-.142 5.93 1.126 5.93 3.175m-6.584-1.823c-3.293.362-2.913 3.259-2.913 3.259s-.034.917.883 1.384c1.927.98 3.912.387 4.915-.829c1.003-1.216.415-4.173-2.885-3.814m.281 3.075c0 .48-.498.925-1.112.99c-.614.068-1.11-.265-1.11-.747s.44-.985 1.055-1.045c.707-.064 1.167.318 1.167.802m1.003-1.15c.11.174.031.437-.173.588c-.208.146-.464.126-.574-.05c-.115-.17-.072-.445.139-.588c.244-.171.498-.122.608.05m4.86-9.806c.335-.06 1.532-.281 2.696-.025c2.083.456 4.941 2.346 3.655 6.368c-.094.575-.398.62-.76.62c-.432 0-.781-.255-.781-.662c0-.352.155-.71.155-.71c.046-.148.411-1.07-.241-2.448c-1.198-1.887-3.609-1.915-3.893-1.807a3.48 3.48 0 0 1-.591.141l-.106.016l-.014.002c-.437 0-.786-.333-.786-.737a.75.75 0 0 1 .573-.715s.007-.011.018-.014c.024-.004.049-.027.075-.029m.66 2.611s3.367-.584 2.964 2.811a.21.21 0 0 1-.007.054c-.037.241-.264.426-.529.426c-.3 0-.543-.225-.543-.507c0 0 .534-2.269-1.885-1.768c-.299 0-.538-.227-.538-.505c0-.283.24-.51.538-.51"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Weixin(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 7.3c3.2 0 6 2.258 6 5.007c0 1.472-1.1 2.846-2.5 3.926L18 18l-1.9-1.08c-.7.196-1.4.393-2.1.393c-3.4 0-6-2.258-6-5.006C8 9.558 10.7 7.3 14 7.3M7 2c3.5 0 6.5 2.061 7.3 4.81h-.7c-3.4 0-5.999 2.454-5.999 5.497c0 .49.1.981.2 1.472h-.7c-.9 0-1.6-.196-2.5-.393l-2.5 1.178l.699-2.06C1.1 11.324 0 9.753 0 7.89C0 4.552 3.1 2 7 2m5.1 8.049c-.3 0-.7.393-.7.687c0 .392.3.687.7.687c.5 0 .9-.392.9-.687c0-.393-.4-.687-.9-.687m3.8 0c-.3 0-.7.393-.7.687c0 .392.4.687.7.687c.6 0 .9-.392.9-.687c0-.393-.4-.687-.9-.687M4.8 4.846c-.6 0-1.1.393-1.1.884c0 .589.6.884 1.1.884c.5 0 .8-.295.9-.884c0-.59-.4-.884-.9-.884m4.9 0c-.6 0-1.1.393-1.1.884c0 .589.6.884 1.1.884c.5 0 .9-.295.9-.884c0-.59-.4-.884-.9-.884"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wifi(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.001 14a2 2 0 1 1 0 4a2 2 0 0 1 0-4m0-3.853a5.695 5.695 0 0 1 4.8 2.626a.7.7 0 0 1-1.178.756a4.295 4.295 0 0 0-3.622-1.982a4.295 4.295 0 0 0-3.624 1.985a.7.7 0 0 1-1.179-.755a5.695 5.695 0 0 1 4.803-2.63m0-3.944c3.06 0 5.863 1.381 7.676 3.724a.7.7 0 0 1-1.107.857C15.02 8.783 12.626 7.603 10 7.603a8.283 8.283 0 0 0-6.537 3.184a.7.7 0 1 1-1.102-.863a9.682 9.682 0 0 1 7.64-3.721m0-4.203c3.905 0 7.243 1.5 9.819 4.32a.7.7 0 1 1-1.034.945C16.473 4.732 13.51 3.4 10.001 3.4c-3.422 0-6.453 1.372-8.79 3.87A.7.7 0 1 1 .19 6.316C2.785 3.538 6.184 2 10 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Windows(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M20 10.873V20L8.479 18.537l.001-7.664zm-13.12 0l-.001 7.461L0 17.461v-6.588zM20 9.273H8.48l-.001-7.81L20 0zM6.879 1.666l.001 7.607H0V2.539z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Woman(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.3 12.535a.7.7 0 1 1 0 1.399l-6.257-.001V19.3a.7.7 0 0 1-1.4 0v-5.367l-6.943.001a.7.7 0 1 1 0-1.4zM10.343 0c3.148 0 5.7 2.55 5.7 5.696a5.698 5.698 0 0 1-5.7 5.697a5.698 5.698 0 1 1 0-11.393m0 1.4a4.298 4.298 0 1 0 4.3 4.296a4.298 4.298 0 0 0-4.3-4.297"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomIn(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.195 0c4.527 0 8.196 3.62 8.196 8.084a7.989 7.989 0 0 1-1.977 5.267l5.388 5.473a.686.686 0 0 1-.015.98a.71.71 0 0 1-.993-.014l-5.383-5.47a8.23 8.23 0 0 1-5.216 1.849C3.67 16.169 0 12.549 0 8.084C0 3.62 3.67 0 8.195 0m0 1.386c-3.75 0-6.79 2.999-6.79 6.698c0 3.7 3.04 6.699 6.79 6.699s6.791-3 6.791-6.699c0-3.7-3.04-6.698-6.79-6.698m.11 2.19c.383 0 .693.314.693.702v2.976h2.976c.388 0 .703.31.703.693a.698.698 0 0 1-.703.693l-2.976-.001v2.977c0 .388-.31.703-.693.703a.698.698 0 0 1-.693-.703V8.64H4.636a.698.698 0 0 1-.702-.692c0-.383.314-.693.702-.693h2.976V4.278c0-.388.31-.703.693-.703"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomOut(children ...ElementRenderer) *UiwIcon {
	return &UiwIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.195 0c4.527 0 8.196 3.62 8.196 8.084a7.989 7.989 0 0 1-1.977 5.267l5.388 5.473a.686.686 0 0 1-.015.98a.71.71 0 0 1-.993-.014l-5.383-5.47a8.23 8.23 0 0 1-5.216 1.849C3.67 16.169 0 12.549 0 8.084C0 3.62 3.67 0 8.195 0m0 1.386c-3.75 0-6.79 2.999-6.79 6.698c0 3.7 3.04 6.699 6.79 6.699s6.791-3 6.791-6.699c0-3.7-3.04-6.698-6.79-6.698m3.78 5.868c.387 0 .702.31.702.693a.698.698 0 0 1-.703.693H4.636a.698.698 0 0 1-.702-.693c0-.383.314-.693.702-.693z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
