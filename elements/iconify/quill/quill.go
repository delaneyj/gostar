package quill

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = ""
	hAttr          = "1em"
	viewbox        = "0 0 32 32"
	fill           = "currentColor"
)

type QuillIcon struct {
	*SVGSVGElement
}

type QuillIconFn func(children ...ElementRenderer) *QuillIcon

var IconLookup = map[string]QuillIconFn{
	"activity":         Activity,
	"add":              Add,
	"alarm":            Alarm,
	"alt":              Alt,
	"arrowDown":        ArrowDown,
	"arrowLeft":        ArrowLeft,
	"arrowRight":       ArrowRight,
	"arrowUp":          ArrowUp,
	"at":               At,
	"attachment":       Attachment,
	"breather":         Breather,
	"broadcast":        Broadcast,
	"calendar":         Calendar,
	"calendarAdd":      CalendarAdd,
	"calendarMore":     CalendarMore,
	"calendarSomeday":  CalendarSomeday,
	"catchup":          Catchup,
	"chat":             Chat,
	"checkmark":        Checkmark,
	"checkmarkDouble":  CheckmarkDouble,
	"checkmarkTodo":    CheckmarkTodo,
	"chevronDown":      ChevronDown,
	"chevronLeft":      ChevronLeft,
	"chevronRight":     ChevronRight,
	"chevronUp":        ChevronUp,
	"clock":            Clock,
	"cloudia":          Cloudia,
	"cog":              Cog,
	"cogAlt":           CogAlt,
	"collapse":         Collapse,
	"command":          Command,
	"compose":          Compose,
	"creditcard":       Creditcard,
	"desktop":          Desktop,
	"discard":          Discard,
	"download":         Download,
	"earth":            Earth,
	"escape":           Escape,
	"expand":           Expand,
	"eye":              Eye,
	"eyeClosed":        EyeClosed,
	"filter":           Filter,
	"focus":            Focus,
	"folder":           Folder,
	"folderAdd":        FolderAdd,
	"folderArchive":    FolderArchive,
	"folderDownload":   FolderDownload,
	"folderDrafts":     FolderDrafts,
	"folderList":       FolderList,
	"folderOpen":       FolderOpen,
	"folderPut":        FolderPut,
	"folderSpam":       FolderSpam,
	"folderTodo":       FolderTodo,
	"folderTrash":      FolderTrash,
	"forceBatch":       ForceBatch,
	"forcebatch":       Forcebatch,
	"formatting":       Formatting,
	"forward":          Forward,
	"fullscreen":       Fullscreen,
	"gift":             Gift,
	"hamburger":        Hamburger,
	"hamburgerSidebar": HamburgerSidebar,
	"inbox":            Inbox,
	"inboxAdd":         InboxAdd,
	"inboxDouble":      InboxDouble,
	"inboxList":        InboxList,
	"inboxNewsletter":  InboxNewsletter,
	"info":             Info,
	"inlineDown":       InlineDown,
	"inlineLeft":       InlineLeft,
	"inlineRight":      InlineRight,
	"inlineUp":         InlineUp,
	"jump":             Jump,
	"jumpAlt":          JumpAlt,
	"label":            Label,
	"labelMini":        LabelMini,
	"link":             Link,
	"linkOut":          LinkOut,
	"list":             List,
	"loadingSpin":      LoadingSpin,
	"lock":             Lock,
	"lockWindow":       LockWindow,
	"mail":             Mail,
	"mailList":         MailList,
	"mailOpen":         MailOpen,
	"mailPlus":         MailPlus,
	"mailSubbed":       MailSubbed,
	"mailUnsub":        MailUnsub,
	"markdown":         Markdown,
	"meatballsH":       MeatballsH,
	"meatballsV":       MeatballsV,
	"moon":             Moon,
	"mute":             Mute,
	"notifications":    Notifications,
	"nuclear":          Nuclear,
	"off":              Off,
	"outbox":           Outbox,
	"paper":            Paper,
	"pause":            Pause,
	"phone":            Phone,
	"pin":              Pin,
	"play":             Play,
	"printAlt":         PrintAlt,
	"printer":          Printer,
	"queue":            Queue,
	"remind":           Remind,
	"reply":            Reply,
	"replyAll":         ReplyAll,
	"search":           Search,
	"searchAlt":        SearchAlt,
	"send":             Send,
	"sendCancelled":    SendCancelled,
	"sendLater":        SendLater,
	"sendStop":         SendStop,
	"share":            Share,
	"sign":             Sign,
	"signature":        Signature,
	"skip":             Skip,
	"snoozeMonth":      SnoozeMonth,
	"snoozeTomorrow":   SnoozeTomorrow,
	"snoozeWeek":       SnoozeWeek,
	"snoozeWeekend":    SnoozeWeekend,
	"sort":             Sort,
	"sortAlt":          SortAlt,
	"sound":            Sound,
	"stack":            Stack,
	"stackAlt":         StackAlt,
	"star":             Star,
	"stopwatch":        Stopwatch,
	"sun":              Sun,
	"textCenter":       TextCenter,
	"textJustify":      TextJustify,
	"textLeft":         TextLeft,
	"textRight":        TextRight,
	"toDo":             ToDo,
	"userHappy":        UserHappy,
	"userNeutral":      UserNeutral,
	"userSad":          UserSad,
	"vip":              Vip,
	"warning":          Warning,
	"warningAlt":       WarningAlt,
}

func Activity(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 16h6l4-11l6 22l4-11h6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Add(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 25V7m-9 9h18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alarm(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M26.607 6.392a15 15 0 0 0-6.725-3.882m-7.764 0a15 15 0 0 0-6.725 3.882M16 9v8l-4 4m15-4c0 6.075-4.925 11-11 11S5 23.075 5 17S9.925 6 16 6s11 4.925 11 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alt(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M5 9h5.485a1 1 0 0 1 .814.419l9.402 13.162a1 1 0 0 0 .814.419H27M21 9h6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDown(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 26V5M8 19l8 8l8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 16h21M13 8l-8 8l8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 16h21m-7-8l8 8l-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUp(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 27V6m8 7l-8-8l-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func At(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m23.363 5.286l.567-.824zm-9.228-2.151l.143.99zM5.897 7.819l-.777-.63zm-2.87 9.031l-.997.066zm22.34 9.554a1 1 0 1 0-1.337-1.486zM22 10.5a1 1 0 1 0-2 0zm-.744 9.896l.8-.6zM24.5 22v1zM20 16a5 5 0 0 1-5 5v2a7 7 0 0 0 7-7zm-5 5a5 5 0 0 1-5-5H8a7 7 0 0 0 7 7zm-5-5a5 5 0 0 1 5-5V9a7 7 0 0 0-7 7zm5-5a5 5 0 0 1 5 5h2a7 7 0 0 0-7-7zm14.657 1.919c-.735-3.386-2.827-6.463-5.727-8.457l-1.133 1.649c2.491 1.712 4.281 4.356 4.905 7.232zM23.93 4.462a14 14 0 0 0-9.939-2.317l.287 1.98a12 12 0 0 1 8.519 1.986zm-9.94-2.317A14 14 0 0 0 5.12 7.19l1.554 1.258a12 12 0 0 1 7.604-4.324zM5.12 7.19a14 14 0 0 0-3.09 9.726l1.996-.131a12 12 0 0 1 2.648-8.337zm-3.09 9.726a14 14 0 0 0 4.333 9.24l1.377-1.451a12 12 0 0 1-3.714-7.92zm4.333 9.24a14 14 0 0 0 9.454 3.843l.026-2a12 12 0 0 1-8.103-3.294zm9.454 3.843a14 14 0 0 0 9.55-3.595l-1.337-1.486a12 12 0 0 1-8.187 3.081zM20 10.5v9.028h2V10.5zm.456 10.496c.54.72 1.823 2.004 4.044 2.004v-2a2.97 2.97 0 0 1-2.444-1.204zM20 19.528c0 .449.1.992.456 1.468l1.6-1.2c-.016-.022-.056-.096-.056-.268zM24.5 23c1.357 0 2.465-.44 3.314-1.2c.83-.744 1.354-1.737 1.673-2.764c.632-2.034.535-4.434.17-6.117l-1.955.424c.314 1.447.378 3.482-.125 5.1c-.248.8-.616 1.436-1.097 1.866c-.462.414-1.087.691-1.98.691z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Attachment(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10.49 19.182l7.778-7.778c.976-.976 2.382-1.153 3.359-.177c.976.976.8 2.383-.177 3.359l-8.132 8.132c-1.414 1.414-4.243 1.414-6.01-.354c-1.768-1.768-1.768-4.596-.354-6.01l8.132-8.132a6.5 6.5 0 0 1 9.192 9.192l-4.596 4.596"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Breather(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M25.192 25.192A13 13 0 1 0 6.808 6.808a13 13 0 0 0 18.384 18.384"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Broadcast(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 20.66a10 10 0 0 0 0-17.32M11 20.66a10 10 0 0 1 0-17.32m8 13.856a6 6 0 0 0 0-10.392m-6 10.392a6 6 0 0 1 0-10.392M13.09 26L16 18l2.91 8m-5.82 0L12 29m1.09-3h5.82M20 29l-1.09-3M18 12a2 2 0 1 0-4 0a2 2 0 0 0 4 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.5 16.5h2V23M5 12h22m-6-4V4M11 8V4M7 28h18a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v18a2 2 0 0 0 2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarAdd(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 23.5V20m0 0v-3.5m0 3.5h-3.5m3.5 0h3.5M5 12h22m-6-4V4M11 8V4M7 28h18a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v18a2 2 0 0 0 2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarMore(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<mask id="quillCalendarMore0" fill="#fff"><path fill-rule="evenodd" d="M12.5 19.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0m5 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0M21 21a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3" clip-rule="evenodd"/></mask><g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h22m-6-4V4M11 8V4M7 28h18a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v18a2 2 0 0 0 2 2"/><path fill="currentColor" fill-rule="evenodd" d="M12.5 19.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0m5 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0M21 21a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3" clip-rule="evenodd"/><path fill="currentColor" d="M11 22a2.5 2.5 0 0 0 2.5-2.5h-2a.5.5 0 0 1-.5.5zm-2.5-2.5A2.5 2.5 0 0 0 11 22v-2a.5.5 0 0 1-.5-.5zM11 17a2.5 2.5 0 0 0-2.5 2.5h2a.5.5 0 0 1 .5-.5zm2.5 2.5A2.5 2.5 0 0 0 11 17v2a.5.5 0 0 1 .5.5zM16 22a2.5 2.5 0 0 0 2.5-2.5h-2a.5.5 0 0 1-.5.5zm-2.5-2.5A2.5 2.5 0 0 0 16 22v-2a.5.5 0 0 1-.5-.5zM16 17a2.5 2.5 0 0 0-2.5 2.5h2a.5.5 0 0 1 .5-.5zm2.5 2.5A2.5 2.5 0 0 0 16 17v2a.5.5 0 0 1 .5.5zm3 0a.5.5 0 0 1-.5.5v2a2.5 2.5 0 0 0 2.5-2.5zM21 19a.5.5 0 0 1 .5.5h2A2.5 2.5 0 0 0 21 17zm-.5.5a.5.5 0 0 1 .5-.5v-2a2.5 2.5 0 0 0-2.5 2.5zm.5.5a.5.5 0 0 1-.5-.5h-2A2.5 2.5 0 0 0 21 22z" mask="url(#quillCalendarMore0)"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarSomeday(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path fill="currentColor" d="M16.75 23.75a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Z"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 17.462a2 2 0 1 1 2 2V20.5M5 12h22m-6-4V4M11 8V4M7 28h18a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v18a2 2 0 0 0 2 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Catchup(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M23 27H7a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v6m0 14h2a2 2 0 0 0 2-2V15a2 2 0 0 0-2-2h-2m0 14V13M9 10h4m-4 7h10M9 21h6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chat(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M25 5H7a4 4 0 0 0-4 4v10a4 4 0 0 0 4 4h11l6 4v-4h1a4 4 0 0 0 4-4V9a4 4 0 0 0-4-4"/><path d="M10 15a1 1 0 1 1-2 0a1 1 0 0 1 2 0m6 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0m6 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Checkmark(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6.5 17l6 6l13-13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckmarkDouble(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 17l5 5l12-12m-5 10l2 2l12-12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckmarkTodo(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M26.207 9.293a1 1 0 0 1 0 1.414l-1.625 1.625a1 1 0 0 1-1.414-1.414l1.625-1.625a1 1 0 0 1 1.414 0m-4.875 4.875a1 1 0 0 1 0 1.414l-3.25 3.25a1 1 0 0 1-1.414-1.414l3.25-3.25a1 1 0 0 1 1.414 0m-15.54 2.125a1 1 0 0 1 1.415 0l1.5 1.5a1 1 0 1 1-1.414 1.414l-1.5-1.5a1 1 0 0 1 0-1.414m9.04 4.375a1 1 0 0 1 0 1.414l-1.625 1.625a1 1 0 0 1-1.414 0l-1.5-1.5a1 1 0 1 1 1.414-1.414l.793.793l.918-.918a1 1 0 0 1 1.414 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDown(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 12l10 10l10-10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeft(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 6L10 16l10 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRight(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 26l10-10L12 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUp(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M26 20L16 10L6 20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 8v8l4 4m9-4c0 7.18-5.82 13-13 13S3 23.18 3 16S8.82 3 16 3s13 5.82 13 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloudia(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path fill="currentColor" d="M11.5 18a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm10 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 21h-2m12.708-4.717a9 9 0 1 0-17.66-3.218A7.001 7.001 0 0 0 10 27h13a6 6 0 0 0 3.708-10.717"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cog(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13.905 3.379A.5.5 0 0 1 14.39 3h3.22a.5.5 0 0 1 .485.379l.689 2.757a.515.515 0 0 0 .341.362c.383.126.755.274 1.115.443a.515.515 0 0 0 .449-.003l2.767-1.383a.5.5 0 0 1 .577.093l2.319 2.319a.5.5 0 0 1 .093.577l-1.383 2.767a.515.515 0 0 0-.003.449c.127.271.243.549.346.833c.053.148.17.265.319.315l2.934.978a.5.5 0 0 1 .342.474v3.28a.5.5 0 0 1-.342.474l-2.934.978a.515.515 0 0 0-.32.315a9.937 9.937 0 0 1-.345.833a.515.515 0 0 0 .003.449l1.383 2.767a.5.5 0 0 1-.093.577l-2.319 2.319a.5.5 0 0 1-.577.093l-2.767-1.383a.515.515 0 0 0-.449-.003a9.994 9.994 0 0 1-.833.346a.515.515 0 0 0-.315.319l-.978 2.934a.5.5 0 0 1-.474.342h-3.28a.5.5 0 0 1-.474-.342l-.978-2.934a.515.515 0 0 0-.315-.32a9.95 9.95 0 0 1-1.101-.475a.515.515 0 0 0-.498.014l-2.437 1.463a.5.5 0 0 1-.611-.075l-2.277-2.277a.5.5 0 0 1-.075-.61l1.463-2.438a.515.515 0 0 0 .014-.498a9.938 9.938 0 0 1-.573-1.383a.515.515 0 0 0-.362-.341l-2.757-.69A.5.5 0 0 1 3 17.61v-3.22a.5.5 0 0 1 .379-.485l2.757-.689a.515.515 0 0 0 .362-.341c.157-.478.35-.94.573-1.383a.515.515 0 0 0-.014-.498L5.594 8.557a.5.5 0 0 1 .075-.611l2.277-2.277a.5.5 0 0 1 .61-.075l2.438 1.463c.152.091.34.094.498.014a9.938 9.938 0 0 1 1.382-.573a.515.515 0 0 0 .342-.362z"/><path d="M21 16a5 5 0 1 1-10 0a5 5 0 0 1 10 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CogAlt(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M19.19 3.757A1 1 0 0 0 18.22 3h-4.44a1 1 0 0 0-.97.757l-.66 2.644a1 1 0 0 1-.47.623l-1.291.747a1 1 0 0 1-.776.095l-2.62-.75a1 1 0 0 0-1.142.462l-2.219 3.844a1 1 0 0 0 .171 1.219l1.96 1.895a1 1 0 0 1 .305.719v1.49a1 1 0 0 1-.305.72l-1.96 1.894a1 1 0 0 0-.17 1.22l2.218 3.843a1 1 0 0 0 1.141.461l2.61-.746a1 1 0 0 1 .793.106l.963.584a1 1 0 0 1 .43.54l.984 2.95a1 1 0 0 0 .949.683h4.558a1 1 0 0 0 .949-.684l.982-2.947a1 1 0 0 1 .435-.542l.982-.587a1 1 0 0 1 .79-.103l2.59.745a1 1 0 0 0 1.142-.461l2.222-3.848a1 1 0 0 0-.166-1.214l-1.933-1.896a1 1 0 0 1-.3-.713v-1.5a1 1 0 0 1 .3-.713l1.933-1.896a1 1 0 0 0 .166-1.214l-2.222-3.848a1 1 0 0 0-1.142-.46l-2.6.747a1 1 0 0 1-.774-.093l-1.31-.75a1 1 0 0 1-.474-.625z"/><path d="M21 16a5 5 0 1 1-10 0a5 5 0 0 1 10 0Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Collapse(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m23 26l-7-7l-7 7M9 6l7 7l7-7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Command(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="2" d="M19 19v4a4 4 0 1 0 4-4zm0 0v-6m0 6h-6m6-6V9a4 4 0 1 1 4 4zm0 0h-6m-4 0h-.007m0 0A4 4 0 1 1 13 9v4m-4.007 0H13m0 0v6m0 0v4a4 4 0 1 1-4-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compose(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M27 5L17 15m0-10H8a3 3 0 0 0-3 3v16a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3v-9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Creditcard(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 20h8M3 13h26m-6 7h2M5 7C4 7 3 8 3 9v14c0 1 1 2 2 2h22c1 0 2-1 2-2V9c0-1-1-2-2-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Desktop(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 27h4m-4 0l.5-4h3l.5 4m-4 0h-3m7 0h3M5 5h22a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Discard(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9.5 23l13-13.5M29 16c0 7.18-5.82 13-13 13S3 23.18 3 16S8.82 3 16 3s13 5.82 13 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 22V5M9 16l7 7l7-7M9 27h14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Earth(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="2" d="M21.5 4.5c0 1.167-1.735 1.5-3 1.5c-5.217 0-10.705 3.48-11.421 8.004C6.992 14.549 6.552 15 6 15H3m1 5.5c.785.262 2.126 1.285 3.44 1.517c.57.101 1.153.464 1.299 1.023c.303 1.16.548 1.992-.239 3.96m8.725-1.707c-1 0-5-2.2-5-9c0-5.5 3.5-7 8.5-7c3 0 5 1.5 5 5s-3.5 2.707-5 4.207s0 6.793-3.5 6.793ZM29 16c0 7.18-5.82 13-13 13S3 23.18 3 16S8.82 3 16 3s13 5.82 13 13Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Escape(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 23L23 9m0 14L9 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Expand(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 20l7 7l7-7m0-8l-7-7l-7 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M29 16c0 3-5.82 9-13 9S3 19 3 16s5.82-9 13-9s13 6 13 9Z"/><path d="M21 16a5 5 0 1 1-10 0a5 5 0 0 1 10 0Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeClosed(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 16a5 5 0 0 1-5 5m-5-5a5 5 0 0 1 5-5m-3 13.654A13.38 13.38 0 0 0 16 25c7.18 0 13-6 13-9c0-1.336-1.155-3.268-3.071-5M19.5 7.47A13.49 13.49 0 0 0 16 7C8.82 7 3 13 3 16c0 1.32 1.127 3.22 3 4.935M7 25L25 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filter(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 7h22M7 13h18M9 19h14m-12 6h10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Focus(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 16a2 2 0 1 1-4 0a2 2 0 0 1 4 0m9 0a2 2 0 1 1-4 0a2 2 0 0 1 4 0"/><path fill="currentColor" d="M9.5 16a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M28 11v13a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h6c3 0 3 3 5 3h9.003C27.108 9 28 9.895 28 11Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderAdd(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M16 21v-8m-4 4h8m8 7V11c0-1.105-.892-2-1.997-2H17c-2 0-2-3-5-3H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderArchive(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 11h20m-14 5h8M5 7v4h2v14.002C7 26.106 7.895 27 9 27h14c1.105 0 2-.894 2-1.998V11h2V7a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderDownload(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 21l-.707.707a1 1 0 0 0 1.414 0zm1-8a1 1 0 1 0-2 0zm3.207 5.207a1 1 0 0 0-1.414-1.414zm-7-1.414a1 1 0 0 0-1.414 1.414zM29 24V11h-2v13zM26.003 8H17v2h9.003zM17 8c-.208 0-.36-.063-.552-.228c-.245-.21-.442-.477-.805-.912c-.327-.393-.755-.872-1.35-1.24C13.678 5.236 12.933 5 12 5v2c.568 0 .947.138 1.238.318c.311.194.571.465.869.822c.262.315.627.797 1.04 1.15c.463.398 1.06.71 1.853.71zm-5-3H6v2h6zM3 8v16h2V8zm3 19h20v-2H6zM6 5a3 3 0 0 0-3 3h2a1 1 0 0 1 1-1zm23 6c0-1.655-1.338-3-2.997-3v2c.55 0 .997.446.997 1zM3 24a3 3 0 0 0 3 3v-2a1 1 0 0 1-1-1zm24 0a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3zm-10-3v-8h-2v8zm1.793-4.207l-3.5 3.5l1.414 1.414l3.5-3.5zm-2.086 3.5l-3.5-3.5l-1.414 1.414l3.5 3.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderDrafts(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 8h8m-8 4h22M5 16h6m-6 4h12m-3-4h13m-6 4h6M5 24h2m4 0h4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderList(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M12 15h8m-8 4h8m8 5V11c0-1.105-.892-2-1.997-2H17c-2 0-2-3-5-3H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpen(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M4 26V8a2 2 0 0 1 2-2h6c3 0 3 3 5 3h7a2 2 0 0 1 2 2v2M4 26l3.783-12.294A1 1 0 0 1 8.739 13H26M4 26h19.523a2 2 0 0 0 1.911-1.412l3.168-10.294A1 1 0 0 0 27.646 13H26"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderPut(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.168 9.555a1 1 0 0 0 1.664-1.11zM25 8a1 1 0 1 0 0 2zm-4 10l-.707.707a1 1 0 0 0 1.414 0zm1-12a1 1 0 1 0-2 0zm3.207 9.207a1 1 0 0 0-1.414-1.414zm-7-1.414a1 1 0 0 0-1.414 1.414zM29 24V11h-2v13zM17 9l.832-.555v-.001l-.002-.002a.164.164 0 0 1-.01-.014a3.047 3.047 0 0 0-.107-.15a9.757 9.757 0 0 0-1.437-1.537C15.313 5.915 13.838 5 12 5v2c1.162 0 2.187.585 2.974 1.26a7.753 7.753 0 0 1 1.13 1.205a3.88 3.88 0 0 1 .065.091c0-.001-.001-.001.831-.556m-5-4H6v2h6zM3 8v16h2V8zm3 19h20v-2H6zM26 8h-1v2h1zM6 5a3 3 0 0 0-3 3h2a1 1 0 0 1 1-1zm23 6a3 3 0 0 0-3-3v2a1 1 0 0 1 1 1zM3 24a3 3 0 0 0 3 3v-2a1 1 0 0 1-1-1zm24 0a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3zm-5-6V6h-2v12zm1.793-4.207l-3.5 3.5l1.414 1.414l3.5-3.5zm-2.086 3.5l-3.5-3.5l-1.414 1.414l3.5 3.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderSpam(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.5 12.5c1.5 0 5.5 0 5.5 7M11 5c4 0 8.5 2.8 8.5 6c0 5-9.5 3-9.5 8c0 5.5 17-5 17 4c0 4-8.5 4-13 4c-12 0-11-10.5-3-10.5C15.4 12.1 12.5 7 11 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderTodo(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.707 16.293a1 1 0 0 0-1.414 1.414zM15 20l-.707.707a1 1 0 0 0 1.414 0zm6.207-4.793a1 1 0 0 0-1.414-1.414zm-9.914 2.5l3 3l1.414-1.414l-3-3zm4.414 3l5.5-5.5l-1.414-1.414l-5.5 5.5zM29 24V11h-2v13zM26.003 8H17v2h9.003zM17 8c-.208 0-.36-.063-.552-.228c-.245-.21-.442-.477-.805-.912c-.327-.393-.755-.872-1.35-1.24C13.678 5.236 12.933 5 12 5v2c.568 0 .947.138 1.238.318c.311.194.571.465.869.822c.262.315.627.797 1.04 1.15c.463.398 1.06.71 1.853.71zm-5-3H6v2h6zM3 8v16h2V8zm3 19h20v-2H6zM6 5a3 3 0 0 0-3 3h2a1 1 0 0 1 1-1zm23 6c0-1.655-1.338-3-2.997-3v2c.55 0 .997.446.997 1zM3 24a3 3 0 0 0 3 3v-2a1 1 0 0 1-1-1zm24 0a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderTrash(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M27 8c0 1.657-4.925 3-11 3S5 9.657 5 8m22 0c0-1.657-4.925-3-11-3S5 6.343 5 8m22 0l-3 18s-1 2-8 2s-8-2-8-2L5 8m13.5 8.5l-5 5m0-5l5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForceBatch(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M24 15h1a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h1m6 4h4m-2-4V3m0 12l4-4m-4 4l-4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Forcebatch(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24 14a1 1 0 1 0 0 2v-2ZM8 16a1 1 0 1 0 0-2v2Zm6 2a1 1 0 1 0 0 2v-2Zm4 2a1 1 0 1 0 0-2v2Zm-2-5l-.707.707a1 1 0 0 0 1.414 0L16 15Zm1-12a1 1 0 1 0-2 0h2Zm3.707 8.707a1 1 0 0 0-1.414-1.414l1.414 1.414Zm-8-1.414a1 1 0 1 0-1.414 1.414l1.414-1.414ZM24 16h1v-2h-1v2Zm2 1v8h2v-8h-2Zm-1 9H7v2h18v-2ZM6 25v-8H4v8h2Zm1-9h1v-2H7v2Zm-1 1a1 1 0 0 1 1-1v-2a3 3 0 0 0-3 3h2Zm1 9a1 1 0 0 1-1-1H4a3 3 0 0 0 3 3v-2Zm19-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3h-2Zm-1-9a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3v2Zm-11 4h4v-2h-4v2Zm3-5V3h-2v12h2Zm-.293.707l4-4l-1.414-1.414l-4 4l1.414 1.414Zm0-1.414l-4-4l-1.414 1.414l4 4l1.414-1.414Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Formatting(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.75 18L10 8L6.25 18m7.5 0L16 24m-2.25-6h-7.5M4 24l2.25-6M28 16c-.333.833-1.8 2-5 2s-3.5 2-3.5 3s.7 3 3.5 3s5-2.5 5-4.5m0-3.5v3.5m0-3.5c0-1.333-.8-4-4-4c-2.4 0-3.667 1.333-4 2m8 5.5V24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Forward(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 13S3 13 3 16s14 3 14 3v3.453c0 1.74 2.069 2.65 3.351 1.475l7.04-6.454a2 2 0 0 0 0-2.948l-7.04-6.454C19.07 6.896 17 7.806 17 9.546z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fullscreen(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 21v3.5a.5.5 0 0 0 .5.5H7M3 11V7.5a.5.5 0 0 1 .5-.5H7m18 0h3.5a.5.5 0 0 1 .5.5V11m0 10v3.5a.5.5 0 0 1-.5.5H25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gift(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 15H5v-4c0-1 1-2 2-2h18c1 0 2 1 2 2v4h-2M7 15v10c0 1 1 2 2 2h14c1 0 2-1 2-2V15M7 15h18m-9-6v18m0-18c-1.333-2.833-4.1-7.4-6.5-5c-2.4 2.4 3 5 6.5 5m0 0c0-4.5 4.5-7.5 6.5-5.5C25 6 20 9 16 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hamburger(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 8h22M5 16h22M5 24h22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HamburgerSidebar(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8h15M5 16h22M5 24h22M5 11l3-3l-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inbox(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 16v8c0 1.5 1.5 3 3 3h16c1.5 0 3-1.5 3-3v-8M5 16h5.5s1.5 3.5 5.5 3.5s5.5-3.5 5.5-3.5H27M5 16L9 5h14l4 11M5 16v3.5M27 16v3.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxAdd(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 27H8c-1.5 0-3-1.5-3-3v-8m0 0h5.5s1 3.5 5.5 3.5s5.5-3.5 5.5-3.5H27L23 5H9zm20 12v-8m-4 4h8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxDouble(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 11v-1a1 1 0 0 0-.962.725zm5.5 0h1a.999.999 0 0 0-.106-.447zM5 11l-.894-.447A1 1 0 0 0 4 11zm5.5 0l.961-.275A1 1 0 0 0 10.5 10zM5 19H4a1 1 0 0 0 1 1zm22 0v1a1 1 0 0 0 1-1zM8 5V4a1 1 0 0 0-.894.553zm16 0l.894-.447A1 1 0 0 0 24 4zm2 9.5a1 1 0 1 0 2 0zM4 14a1 1 0 1 0 2 0zm0-1v11h2V13zm0 11c0 1.075.528 2.067 1.23 2.77C5.933 27.473 6.925 28 8 28v-2c-.425 0-.933-.223-1.355-.645C6.222 24.933 6 24.425 6 24zm4 4h16v-2H8zm16 0c1.075 0 2.067-.527 2.77-1.23c.703-.703 1.23-1.695 1.23-2.77h-2c0 .425-.223.933-.645 1.355c-.422.422-.93.645-1.355.645zm4-4V13h-2v11zm-12-8.5c2.547 0 4.182-1.005 5.17-2.07c.484-.52.8-1.041.998-1.436a5.28 5.28 0 0 0 .285-.691a1.307 1.307 0 0 0 .007-.023v-.003l.001-.001L21.5 11a87.01 87.01 0 0 1-.961-.276v-.003l.001-.002v-.002l-.002.01a3.259 3.259 0 0 1-.159.373c-.13.261-.346.615-.674.97c-.637.685-1.752 1.43-3.705 1.43zm5.5-3.5H27v-2h-5.5zM5 12h5.5v-2H5zm5.5-1a78.07 78.07 0 0 0-.961.276v.004l.003.006a1.573 1.573 0 0 0 .02.066c.013.038.03.09.054.152c.047.124.117.292.216.49a6.2 6.2 0 0 0 .998 1.437c.988 1.064 2.623 2.069 5.17 2.069v-2c-1.953 0-3.068-.745-3.705-1.43a4.196 4.196 0 0 1-.674-.97a3.259 3.259 0 0 1-.162-.383v.002l.001.002v.002l.001.001zM16 21.5c-2.236 0-3.323-.768-3.866-1.4a3.104 3.104 0 0 1-.543-.91a2.405 2.405 0 0 1-.102-.338v-.006v.011s0 .002-.989.143a68.748 68.748 0 0 0-.99.144a.23.23 0 0 0 0 .004l.002.007a1.565 1.565 0 0 0 .012.071a4.387 4.387 0 0 0 .197.676c.158.414.43.957.895 1.499c.957 1.117 2.62 2.099 5.384 2.099zM10.5 18H5v2h5.5zM27 18h-5.5v2H27zm-5.5 1a76.41 76.41 0 0 1-.99-.144v-.002l.001-.004a.1.1 0 0 1 0-.004v.006l-.014.06c-.014.06-.042.158-.088.279a3.104 3.104 0 0 1-.543.908c-.543.633-1.63 1.401-3.866 1.401v2c2.764 0 4.427-.982 5.384-2.1a5.101 5.101 0 0 0 .894-1.497a4.36 4.36 0 0 0 .207-.73a.946.946 0 0 0 .003-.018l.001-.008v-.004s.001-.002-.989-.143m4.5-8v8h2v-8zM6 19v-8H4v8zM8 6h16V4H8zm15.106-.553l3 6l1.788-.894l-3-6zM26 11v3.5h2V11zM7.106 4.553l-3 6l1.788.894l3-6zM4 11v3h2v-3zm6.5 9h11v-2h-11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxList(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 28a1 1 0 1 0 0-2zM5 16l-.94-.342A1 1 0 0 0 4 16zm5.5 0l.961-.275A1 1 0 0 0 10.5 15zm11 0v-1a1 1 0 0 0-.962.725zm5.5 0v1a1 1 0 0 0 .94-1.342zM23 5l.94-.342A1 1 0 0 0 23 4zM9 5V4a1 1 0 0 0-.94.658zm12 17a1 1 0 1 0 0 2zm6 2a1 1 0 1 0 0-2zm-6 2a1 1 0 1 0 0 2zm6 2a1 1 0 1 0 0-2zm-10-2H8v2h9zm-9 0c-.425 0-.933-.223-1.355-.645C6.222 24.933 6 24.425 6 24H4c0 1.075.528 2.067 1.23 2.77C5.933 27.473 6.925 28 8 28zm-2-2v-8H4v8zm-1-7h5.5v-2H5zm5.5-1a78.07 78.07 0 0 0-.961.276v.004l.003.006a1.573 1.573 0 0 0 .02.066c.013.038.03.09.054.152c.047.124.117.292.216.49a6.2 6.2 0 0 0 .998 1.436c.988 1.065 2.623 2.07 5.17 2.07v-2c-1.953 0-3.068-.745-3.705-1.43a4.196 4.196 0 0 1-.674-.97a3.259 3.259 0 0 1-.162-.383v.002l.001.002v.002l.001.001zm5.5 4.5c2.547 0 4.182-1.005 5.17-2.07c.484-.52.8-1.041.998-1.436a5.28 5.28 0 0 0 .285-.691a1.307 1.307 0 0 0 .007-.023v-.003l.001-.001L21.5 16a87.01 87.01 0 0 1-.961-.276v-.003l.001-.002v-.002l-.002.01a3.259 3.259 0 0 1-.159.372c-.13.262-.346.616-.674.97c-.637.686-1.752 1.431-3.705 1.431zm5.5-3.5H27v-2h-5.5zm6.44-1.342l-4-11l-1.88.684l4 11zM23 4H9v2h14zm-14.94.658l-4 11l1.88.684l4-11zM21 24h6v-2h-6zm0 4h6v-2h-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxNewsletter(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 16v8c0 1.5 1.5 3 3 3h16c1.5 0 3-1.5 3-3v-8M5 16h5.5s1 3.5 5.5 3.5s5.5-3.5 5.5-3.5H27M5 16v3.5M5 16l1-5m21 5l-1-5M13.5 9h5m-5 4h5m-9 0V5h13v8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Info(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 14h1v9h1m12-7a13 13 0 1 1-26 0a13 13 0 0 1 26 0"/><path fill="currentColor" d="M17 9.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InlineDown(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 13l8 8l8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InlineLeft(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 8l-8 8l8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InlineRight(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13 24l8-8l-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InlineUp(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m24 19l-8-8l-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Jump(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M26.622 5.17c.096-.043.185.072.119.154l-7.355 9.193a.5.5 0 0 0 .167.76l3.563 1.781a1 1 0 0 1-.037 1.806L5.38 26.83c-.096.043-.185-.072-.12-.154l7.355-9.193a.5.5 0 0 0-.167-.76l-3.563-1.781a1 1 0 0 1 .037-1.806z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JumpAlt(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m24.95 19l-6 6l-6-6M19 24.5V9a2 2 0 0 0-2-2H9m0-2v4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Label(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 19l5.5-5.5M13 22l2.5-2.5M4.414 16.586l11-11A2 2 0 0 1 16.828 5H25a2 2 0 0 1 2 2v8.172a2 2 0 0 1-.586 1.414l-11 11a2 2 0 0 1-2.828 0l-8.172-8.172a2 2 0 0 1 0-2.828"/><path fill="currentColor" d="M23 10a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LabelMini(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M24 14.172V10a2 2 0 0 0-2-2h-4.172a2 2 0 0 0-1.414.586l-8 8a2 2 0 0 0 0 2.828l4.172 4.172a2 2 0 0 0 2.828 0l8-8A2 2 0 0 0 24 14.172"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14.9 19.142l-.708-.707c-2.12-2.121-2.12-4.95 0-7.071l2.829-2.828c2.121-2.122 4.95-2.122 7.07 0c2.122 2.12 2.122 4.95 0 7.07m-6.363-2.12l.707.706c2.121 2.122 2.121 4.95 0 7.072l-2.828 2.828c-2.122 2.121-4.95 2.121-7.071 0c-2.122-2.121-2.122-4.95 0-7.071"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkOut(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 3h7v7m-1.5-5.5L20 12m-3-7H8a3 3 0 0 0-3 3v16a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3v-9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func List(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M12 8h15m-15 8h9m-9 8h15M7 24a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm0-8a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm0-8a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LoadingSpin(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22.5 4.742A13 13 0 1 0 29 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 24v-4m5-5V8a5 5 0 0 0-10 0v7M6 27V17a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockWindow(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><circle cx="22" cy="23" r=".5" fill="currentColor"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M24 19v-2.917c0-1.15-.895-2.083-2-2.083s-2 .933-2 2.083V19m-7 8H8a3 3 0 0 1-3-3V8a3 3 0 0 1 3-3h16a3 3 0 0 1 3 3v4M17 25v-4a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-6a2 2 0 0 1-2-2"/><path fill="currentColor" d="M9.5 9a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm4 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm4 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mail(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M29 9v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9m26 0a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2m26 0l-11.862 8.212a2 2 0 0 1-2.276 0L3 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailList(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 7h14m4 0h4M5 13h10m8 0h4M5 19h14m4 0h4M5 25h12m6 0h4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailOpen(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 8l9 7h6l9-7M15.118 3.271L3.706 6.783A1 1 0 0 0 3 7.739V23a2 2 0 0 0 2 2h22a2 2 0 0 0 2-2V7.739a1 1 0 0 0-.706-.956L16.882 3.27a3 3 0 0 0-1.764 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailPlus(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M25 28v-8m-4 4h8m0-15v7m0-7a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2m26 0l-11.862 8.212a2 2 0 0 1-2.276 0L3 9m0 0v14a2 2 0 0 0 2 2h10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailSubbed(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M29 9v7m0-7c0-1-1-2-2-2H5C4 7 3 8 3 9m26 0l-11.862 8.212a2 2 0 0 1-2.276 0L3 9m13 16H5c-1 0-2-1-2-2V9m16.5 14l3 3l6-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailUnsub(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M29 9v7m0-7c0-1-1-2-2-2H5C4 7 3 8 3 9m26 0l-11.862 8.212a2 2 0 0 1-2.276 0L3 9m13 16H5c-1 0-2-1-2-2V9m18 17l6-6m2 3a5 5 0 1 1-10 0a5 5 0 0 1 10 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Markdown(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 20v-8l-4 4l-4-4v8m12-3.5l3.5 3.5m0 0l3.5-3.5M22.5 20v-9M5 7h22a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MeatballsH(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="2" d="M8 17a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm8 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm8 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MeatballsV(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="2" d="M16 24a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0-7a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0-7a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Moon(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.294 5A11.19 11.19 0 1 0 27 18.706s-5.723 2.19-10.81-2.897C11.105 10.723 13.295 5 13.295 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mute(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.5 21H8a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1h7l6.586-6.586C22.846 3.154 25 4.047 25 5.828V6m0 8.5v11.672c0 1.781-2.154 2.674-3.414 1.414L17 23M7 28L29 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notifications(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M27 25c0 1.657-4.925 3-11 3S5 26.657 5 25m22 0c0-1.47-3.88-2.694-9-2.95M27 25c0-2-.2-6.2-3-9c-2.986-2.986-.513-7.427-5-8.668M5 25c0-1.47 3.88-2.694 9-2.95M5 25c0-2 .2-6.2 3-9c2.986-2.986.513-7.427 5-8.668m1 14.717a40.015 40.015 0 0 1 4 0m-4 0V23a2 2 0 1 0 4 0v-.95M13 7.331C13.773 7.12 14.751 7 16 7s2.227.119 3 .332m-6 0V7a3 3 0 0 1 6 0v.332"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Nuclear(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><path fill="currentColor" d="M11.1 17a4.996 4.996 0 0 0 1.585 2.743l-2.024 3.504A8.991 8.991 0 0 1 7.056 17zm8.216 2.743A4.996 4.996 0 0 0 20.9 17h4.045a8.992 8.992 0 0 1-3.607 6.247zM16 11c-.553 0-1.086.09-1.584.256l-2.023-3.504A8.966 8.966 0 0 1 16 7c1.284 0 2.504.268 3.607.752l-2.023 3.504A4.994 4.994 0 0 0 16 11Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M15 16a1 1 0 1 1 2 0a1 1 0 0 1-2 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Off(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22.5 4.742a13 13 0 1 1-13 0M16 3v10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Outbox(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 16v8c0 1.5 1.5 3 3 3h16c1.5 0 3-1.5 3-3v-8M5 16h5.5s1 3.5 5.5 3.5s5.5-3.5 5.5-3.5H27M5 16v3.5M5 16l2.5-6h2M27 16l-2.5-6h-2m4.5 6v3.5M16 15V4m4 3l-4-4l-4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paper(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 9h4m-4 7h12m-12 4h12m-12 4h4m-6 5h16a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v22a2 2 0 0 0 2 2"/><circle cx="22" cy="9" r=".5" fill="currentColor"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.5 7v18m9-18v18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 25h2M11 3h10a2 2 0 0 1 2 2v22a2 2 0 0 1-2 2H11a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pin(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 18l-8 8M20.667 4L28 11.333l-6.38 6.076a2 2 0 0 0-.62 1.448v3.729c0 .89-1.077 1.337-1.707.707L8.707 12.707c-.63-.63-.184-1.707.707-1.707h3.729a2 2 0 0 0 1.448-.62z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 24.414V7.586c0-1.746 2.081-2.653 3.36-1.465l9.062 8.413a2 2 0 0 1 0 2.932l-9.061 8.413C13.08 27.067 11 26.16 11 24.414"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PrintAlt(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 22h-3.621c-.252 0-.504-.022-.752-.067l-.92-.164c-3.383-.604-5.244-3.076-5.62-5.769c-.484-3.466 1.494-7.3 5.861-7.992a.114.114 0 0 0 .08-.053l.636-1.049C13.44 3.982 16.583 3.56 19 4.687c2.378 1.11 4.055 3.722 3.037 6.93a.11.11 0 0 0 .084.142c2.952.586 4.372 3.143 4.173 5.567a5.188 5.188 0 0 1-.915 2.566c-1.353 1.92-4.028 2.1-6.377 2.115c-1 .006-2-.007-3.002-.007"/><path d="m14.5 22l-.5 6h4l-.5-6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Printer(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 24h3a2 2 0 0 0 2-2V12a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h3m0-14V5h12v5M9 17h14m-13 0v10h12V17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Queue(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M23 5v8m-4-4h8M7 9.731V24.27c0 1.508 1.753 2.293 2.83 1.266l7.632-7.269a1.758 1.758 0 0 0 0-2.532L9.83 8.465C8.753 7.44 7 8.223 7 9.731Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Remind(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M15 8c-2 2-6.4 5.6-8 4c-3.5-3.5 4.667-3.333 8-4Zm0 0c2.833 0 12-1 8.5 2.5C21 13 17.833 9.667 15 8Zm-2.172 17.828L7.5 20.5a2.121 2.121 0 1 1 3-3l1.83 1.83c.247.247.67.072.67-.277V7a2 2 0 1 1 4 0v7.997c0 .333.419.48.626.22a1.759 1.759 0 0 1 2.747 0l.398.497a.453.453 0 0 0 .557.122l.666-.333a2.249 2.249 0 0 1 2.012 0c.609.304.994.927.994 1.609V23a4 4 0 0 1-4 4h-5.343a4 4 0 0 1-2.829-1.172Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reply(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5.608 12.526l7.04-6.454C13.932 4.896 16 5.806 16 7.546V11c13 0 11 16 11 16s-4-10-11-10v3.453c0 1.74-2.069 2.65-3.351 1.475l-7.04-6.454a2 2 0 0 1 0-2.948"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReplyAll(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 19.141v1.313c0 1.74-2.069 2.65-3.351 1.474l-7.04-6.454a2 2 0 0 1 0-2.948l7.04-6.454C11.93 4.896 14 5.806 14 7.546V8.86m3.04-2.787L10 12.526a2 2 0 0 0 0 2.948l7.04 6.454c1.283 1.176 3.352.266 3.352-1.475V17c5 0 8.5 10 8.5 10s1.5-16-8.5-16V7.546c0-1.74-2.069-2.65-3.352-1.474"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 27l7.5-7.5M28 13a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchAlt(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m28 27l-7.5-7.5M5 13a9 9 0 1 0 18 0a9 9 0 0 0-18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Send(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M29 3L3 15l12 2.5M29 3L19 29l-4-11.5M29 3L15 17.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SendCancelled(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m20 19l4 4m0 0l4 4m-4-4l-4 4m4-4l4-4m1-16L3 15l12 2.5M29 3L15 17.5M29 3l-4.375 11.375M15 17.5l.625 2.875"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SendLater(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M24 21v2l-1.5 1.5M29 3L3 15l12 2.5M29 3L15 17.5M29 3l-4.375 11.375M15 17.5l.625 2.875M29 23a5 5 0 1 1-10 0a5 5 0 0 1 10 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SendStop(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21 26l6-6m2-17L3 15l12 2.5M29 3L15 17.5M29 3l-4.375 11.375M15 17.5l.625 2.875M29 23a5 5 0 1 1-10 0a5 5 0 0 1 10 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17v8a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2v-8m-11 3V3.5M22 9l-6-6l-6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sign(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 25h-5L7 15c2-.667 6-4 6-12h6c0 7.6 4 11.167 6 12l-4 10zm0 0v-6m0-6V3M9 25h14v4H9zm9-9a2 2 0 1 1-4 0a2 2 0 0 1 4 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Signature(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 27C6 19.333 13.5 5 18.5 5C23 5 10 26 15 26c3.5 0 6-10.5 8.5-10.5s-1 9 1 9c2.5 0 4-4 4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Skip(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M23 5v22M8 7.586v16.828c0 1.746 2.081 2.653 3.36 1.465l9.062-8.413a2 2 0 0 0 0-2.932L11.36 6.121C10.08 4.933 8 5.84 8 7.586"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnoozeMonth(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16.5 17l3 3l-3 3M12 17l3 3l-3 3M5 12h22m-6-4V4M11 8V4M7 28h18a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v18a2 2 0 0 0 2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnoozeTomorrow(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 21h22M7 26h18M10 16a6 6 0 0 1 12 0M6.5 16h-2M9 9L7.5 7.5M16 6V4m7 5l1.485-1.485M28 16h-2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnoozeWeek(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 17l3 3l-3 3M5 12h22m-6-4V4M11 8V4M7 28h18a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v18a2 2 0 0 0 2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnoozeWeekend(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.5 14v-4a3 3 0 0 1 3-3h15a3 3 0 0 1 3 3v4m-21 0A2.5 2.5 0 0 0 3 16.5V25h5M5.5 14A2.5 2.5 0 0 1 8 16.5V20h16v-3.5a2.5 2.5 0 0 1 2.5-2.5m0 0a2.5 2.5 0 0 1 2.5 2.5V25h-5m0 0H8m16 0v2M8 25v2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sort(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><circle cx="18.078" cy="8.286" r="1" transform="rotate(-9 18.078 8.286)"/><path d="M8.045 8.862a1 1 0 0 0 .313 1.976zm4.263 1.35a1 1 0 0 0-.312-1.976zm-3.325 4.576a1 1 0 1 0 .313 1.976zm10.19.411a1 1 0 1 0-.313-1.975zM9.61 18.74a1 1 0 0 0 .313 1.976zm8.215.724a1 1 0 1 0-.313-1.975zM8.73 25.967l-.157-.988zm-2.29-1.664l-.988.157zm17.778-2.815l-.988.156zm-1.663 2.288l.157.987zM19.271 3.034l-.156-.987zm2.288 1.663l.988-.157zm-16.115.527l.156.988zM3.78 7.513l.988-.157zM26 10v17h2V10zm-1 18H11v2h14zm-14 0a1 1 0 0 1-1-1H8a3 3 0 0 0 3 3zm15-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3zM25 9a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3zM10 27v-2H8v2zM25 7h-3v2h3zM8.358 10.838l3.95-.626l-.312-1.976l-3.951.626zm.938 5.926l9.877-1.565l-.313-1.975l-9.877 1.564zm.626 3.95l7.901-1.251l-.312-1.975l-7.902 1.251zM5.6 6.213l13.828-2.19l-.313-1.975l-13.828 2.19zm14.972-1.359l2.66 16.79l1.975-.312l-2.66-16.79zM22.4 22.788l-13.828 2.19l.313 1.975l13.828-2.19zM7.428 24.147l-2.66-16.79l-1.975.312l2.66 16.79zm1.144.831a1 1 0 0 1-1.144-.831l-1.975.313a3 3 0 0 0 3.432 2.493zm14.66-3.334a1 1 0 0 1-.832 1.144l.313 1.975a3 3 0 0 0 2.494-3.432zM19.427 4.022a1 1 0 0 1 1.144.831l1.975-.313a3 3 0 0 0-3.432-2.493zm-14.14.215a3 3 0 0 0-2.495 3.432l1.976-.313A1 1 0 0 1 5.6 6.212z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAlt(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 5.015v-1a1 1 0 0 0-1 1zm8 1a1 1 0 1 0 0-2zM4 13a1 1 0 1 0 2 0zm23-7.985h1a1 1 0 0 0-1-1zM26 13a1 1 0 1 0 2 0zm-7-8.985a1 1 0 1 0 0 2zm-13.293.278a1 1 0 0 0-1.414 1.414zm22 1.414a1 1 0 0 0-1.414-1.414zM11 26a1 1 0 1 0 0 2zm10 2a1 1 0 1 0 0-2zM5 6.015h.015v-2H5zm.015 0H13v-2H5.015zM4 5.015V13h2V5.015zm22 0V13h2V5.015zm1-1h-.015v2H27zm-.015 0H19v2h7.985zM16.707 15.293L5.722 4.308L4.308 5.722l10.985 10.985zM5.722 4.308l-.015-.015l-1.414 1.414l.015.015zm10.985 12.4L27.692 5.721l-1.414-1.414l-10.985 10.985zM27.692 5.721l.015-.015l-1.414-1.414l-.015.015zM17 27V16h-2v11zm-6 1h5v-2h-5zm5 0h5v-2h-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sound(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 12v8a1 1 0 0 0 1 1h7l6.586 6.586c1.26 1.26 3.414.367 3.414-1.414V5.828c0-1.781-2.154-2.674-3.414-1.414L15 11H8a1 1 0 0 0-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stack(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><circle cx="19" cy="9" r=".5" stroke="currentColor"/><path d="M9 8a1 1 0 0 0 0 2zm4 2a1 1 0 1 0 0-2zm-4 4a1 1 0 1 0 0 2zm10 2a1 1 0 1 0 0-2zM9 18a1 1 0 1 0 0 2zm8 2a1 1 0 1 0 0-2zm9-10v17h2V10zm-1 18H11v2h14zm-14 0a1 1 0 0 1-1-1H8a3 3 0 0 0 3 3zm15-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3zM25 9a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3zM10 27v-2H8v2zM25 7h-2v2h2zM9 10h4V8H9zm0 6h10v-2H9zm0 4h8v-2H9zM7 5h14V3H7zm15 1v17h2V6zm-1 18H7v2h14zM6 23V6H4v17zm1 1a1 1 0 0 1-1-1H4a3 3 0 0 0 3 3zm15-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3zM21 5a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3zM7 3a3 3 0 0 0-3 3h2a1 1 0 0 1 1-1z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StackAlt(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 15.5l13 6l13-6M3 20l13 6l13-6M3 11l13 6l13-6l-13-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.546 4.984a.5.5 0 0 1 .908 0l3.097 6.714a.5.5 0 0 0 .395.287l7.341.87a.5.5 0 0 1 .28.864l-5.427 5.02a.5.5 0 0 0-.15.464l1.44 7.251a.5.5 0 0 1-.735.534l-6.45-3.611a.5.5 0 0 0-.49 0l-6.45 3.61a.5.5 0 0 1-.735-.533l1.44-7.251a.5.5 0 0 0-.15-.465l-5.428-5.02a.5.5 0 0 1 .28-.863l7.342-.87a.5.5 0 0 0 .396-.287z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stopwatch(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M27.9 7.869a15 15 0 0 0-2.769-2.77m-7.434-.988a13 13 0 0 0-3.394 0m12.304 2.282l-2.829 2.829M16 17l5 5m6-5c0 6.075-4.925 11-11 11S5 23.075 5 17S9.925 6 16 6s11 4.925 11 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sun(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 6V3M8.929 8.929L6.808 6.808M6 16H3m13 13v-3m9.192-.808l-2.121-2.12M29 16h-3M8.929 23.071l-2.121 2.121M25.192 6.808l-2.12 2.121M22 16a6 6 0 1 1-12 0a6 6 0 0 1 12 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextCenter(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 7h22M9 13h14M5 19h22M9 25h14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextJustify(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 7h22M5 13h22M5 19h22M5 25h22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextLeft(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 7h22M5 13h14M5 19h22M5 25h14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextRight(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 7h22m-14 6h14M5 19h22m-14 6h14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToDo(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11.5 16l3.5 3.5l6-6m8 2.5a13 13 0 1 1-26 0a13 13 0 0 1 26 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserHappy(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.5 23s1 1 2.5 1s2.5-1 2.5-1m4.5-5a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-12 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0m18-2c0 7.18-5.82 13-13 13S3 23.18 3 16S8.82 3 16 3s13 5.82 13 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserNeutral(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.5 23h5m4.5-5a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-12 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0m18-2c0 7.18-5.82 13-13 13S3 23.18 3 16S8.82 3 16 3s13 5.82 13 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserSad(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.5 24s1-1 2.5-1s2.5 1 2.5 1M29 16c0 7.18-5.82 13-13 13S3 23.18 3 16S8.82 3 16 3s13 5.82 13 13m-6 2a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-12 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vip(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 13l3 6.5l3-6.5m3.5 0v6.5m4-2.225h1.48c.651 0 1.277-.275 1.721-.75a2.338 2.338 0 0 0 .215-2.932a1.394 1.394 0 0 0-1.14-.593H20.5zm0 0V19.5M5 7h22a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Warning(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-width="2" d="M16 18v-6M6.358 27h19.284c1.516 0 2.48-1.62 1.759-2.953l-9.642-17.8c-.757-1.397-2.761-1.397-3.518 0L4.6 24.047C3.877 25.38 4.842 27 6.358 27Z"/><path fill="currentColor" d="M17 21.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WarningAlt(children ...ElementRenderer) *QuillIcon {
	return &QuillIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 17.5v-8M29 16a13 13 0 1 1-26 0a13 13 0 0 1 26 0"/><path fill="currentColor" d="M17 22a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
