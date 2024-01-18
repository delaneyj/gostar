package formkit

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = ""
	hAttr          = "1em"
	viewbox        = "0 0 0 0"
	fill           = "currentColor"
)

type FormkitIcon struct {
	*SVGSVGElement
}

type FormkitIconFn func(children ...ElementRenderer) *FormkitIcon

var IconLookup = map[string]FormkitIconFn{
	"add":           Add,
	"amex":          Amex,
	"android":       Android,
	"apple":         Apple,
	"arrowdown":     Arrowdown,
	"arrowleft":     Arrowleft,
	"arrowright":    Arrowright,
	"arrowup":       Arrowup,
	"avatarman":     Avatarman,
	"avatarwoman":   Avatarwoman,
	"bitcoin":       Bitcoin,
	"bnb":           Bnb,
	"bookmark":      Bookmark,
	"button":        Button,
	"cardano":       Cardano,
	"caretdown":     Caretdown,
	"caretleft":     Caretleft,
	"caretright":    Caretright,
	"caretup":       Caretup,
	"check":         Check,
	"checkbox":      Checkbox,
	"circle":        Circle,
	"close":         Close,
	"color":         Color,
	"compress":      Compress,
	"date":          Date,
	"datetime":      Datetime,
	"discover":      Discover,
	"dogecoin":      Dogecoin,
	"dollar":        Dollar,
	"down":          Down,
	"download":      Download,
	"downloadcloud": Downloadcloud,
	"email":         Email,
	"ethereum":      Ethereum,
	"euro":          Euro,
	"expand":        Expand,
	"export":        Export,
	"eye":           Eye,
	"eyeclosed":     Eyeclosed,
	"facebook":      Facebook,
	"fastforward":   Fastforward,
	"file":          File,
	"fileaudio":     Fileaudio,
	"filedoc":       Filedoc,
	"fileimage":     Fileimage,
	"filepdf":       Filepdf,
	"filevideo":     Filevideo,
	"flag":          Flag,
	"folder":        Folder,
	"franc":         Franc,
	"github":        Github,
	"google":        Google,
	"groupIcon":     GroupIcon,
	"happy":         Happy,
	"heart":         Heart,
	"help":          Help,
	"hidden":        Hidden,
	"info":          Info,
	"instagram":     Instagram,
	"krona":         Krona,
	"left":          Left,
	"link":          Link,
	"linkedin":      Linkedin,
	"linkexternal":  Linkexternal,
	"lira":          Lira,
	"list":          List,
	"mastercard":    Mastercard,
	"medium":        Medium,
	"megaphone":     Megaphone,
	"minimize":      Minimize,
	"month":         Month,
	"multicurrency": Multicurrency,
	"number":        Number,
	"open":          Open,
	"password":      Password,
	"pause":         Pause,
	"paypal":        Paypal,
	"people":        People,
	"peso":          Peso,
	"pinterest":     Pinterest,
	"play":          Play,
	"playcircle":    Playcircle,
	"pound":         Pound,
	"radio":         Radio,
	"rangeIcon":     RangeIcon,
	"reddit":        Reddit,
	"refresh":       Refresh,
	"reorder":       Reorder,
	"repeater":      Repeater,
	"reply":         Reply,
	"rewind":        Rewind,
	"right":         Right,
	"ruble":         Ruble,
	"rupee":         Rupee,
	"sad":           Sad,
	"search":        Search,
	"select":        Select,
	"settings":      Settings,
	"share":         Share,
	"shekel":        Shekel,
	"skype":         Skype,
	"snapchat":      Snapchat,
	"solana":        Solana,
	"spinner":       Spinner,
	"star":          Star,
	"start":         Start,
	"stepback":      Stepback,
	"stepforward":   Stepforward,
	"stop":          Stop,
	"stripe":        Stripe,
	"submit":        Submit,
	"table":         Table,
	"tag":           Tag,
	"telephone":     Telephone,
	"tether":        Tether,
	"textIcon":      TextIcon,
	"textarea":      Textarea,
	"tiktok":        Tiktok,
	"time":          Time,
	"tools":         Tools,
	"trash":         Trash,
	"twitter":       Twitter,
	"up":            Up,
	"upload":        Upload,
	"uploadcloud":   Uploadcloud,
	"url":           Url,
	"usdc":          Usdc,
	"vimeo":         Vimeo,
	"visa":          Visa,
	"volumedown":    Volumedown,
	"volumeup":      Volumeup,
	"warning":       Warning,
	"week":          Week,
	"whatsapp":      Whatsapp,
	"won":           Won,
	"wordpress":     Wordpress,
	"yen":           Yen,
	"youtube":       Youtube,
	"yuan":          Yuan,
	"zip":           Zip,
}

func Add(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 15c-3.86 0-7-3.14-7-7s3.14-7 7-7s7 3.14 7 7s-3.14 7-7 7M8 2C4.69 2 2 4.69 2 8s2.69 6 6 6s6-2.69 6-6s-2.69-6-6-6"/><path fill="currentColor" d="M8 11.5c-.28 0-.5-.22-.5-.5V5c0-.28.22-.5.5-.5s.5.22.5.5v6c0 .28-.22.5-.5.5"/><path fill="currentColor" d="M11 8.5H5c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h6c.28 0 .5.22.5.5s-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Amex(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 13h-13C.67 13 0 12.33 0 11.5v-9C0 1.67.67 1 1.5 1h13c.83 0 1.5.67 1.5 1.5v9c0 .83-.67 1.5-1.5 1.5M1.5 2c-.28 0-.5.22-.5.5v9c0 .28.22.5.5.5h13c.28 0 .5-.22.5-.5v-9c0-.28-.22-.5-.5-.5z"/><path fill="currentColor" fill-rule="evenodd" d="M8 10.05V7.23h2.39v.65H8.77v.44h1.58v.64H8.77v.43h1.62v.66z"/><path fill="currentColor" fill-rule="evenodd" d="m10.38 10.05l1.32-1.41l-1.32-1.41h1.02l.81.89l.81-.89H14v.02l-1.29 1.39L14 10.01v.04h-.99l-.82-.9l-.81.9zM8.24 4.01L7 6.79h.85l.23-.56h1.26l.23.56h.87L9.21 4.01h-.98Zm.11 1.6l.37-.89l.37.89z"/><path fill="currentColor" fill-rule="evenodd" d="M10.43 6.79V4.01h1.19l.61 1.71l.62-1.71H14v2.78h-.74v-1.9l-.7 1.9h-.67l-.71-1.91v1.91z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Android(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.14 6.25c-.47 0-.86.39-.86.87v3.5c0 .48.39.88.86.88s.86-.39.86-.88v-3.5c0-.48-.39-.87-.86-.87m-10.29 0c-.47 0-.86.39-.86.87v3.5c0 .48.39.88.86.88s.86-.39.86-.88v-3.5c0-.48-.39-.87-.86-.87m1.29 4.81c0 .35.14.68.38.93s.57.38.91.38v1.75c0 .48.39.88.86.88s.86-.39.86-.88v-1.75h1.71v1.75c0 .48.39.88.86.88s.86-.39.86-.88v-1.75c.34 0 .67-.14.91-.38c.24-.25.38-.58.38-.93V6.25H4.14zm7.69-5.69c-.06-.58-.25-1.14-.55-1.63c-.3-.5-.7-.92-1.18-1.23l.43-.88c.05-.1.06-.22.02-.33a.418.418 0 0 0-.22-.25a.42.42 0 0 0-.33-.02c-.11.04-.2.12-.25.22l-.43.88l-.11-.05c-.79-.27-1.64-.27-2.43 0l-.11.05l-.43-.88a.452.452 0 0 0-.25-.22a.42.42 0 0 0-.33.02c-.1.05-.18.14-.22.25a.42.42 0 0 0 .02.33l.43.88c-.48.32-.88.74-1.18 1.23c-.3.5-.49 1.05-.55 1.63v.44h7.69v-.44zM6.71 4.5c-.11 0-.22-.05-.3-.13a.439.439 0 0 1-.13-.31c0-.12.04-.23.13-.31c.08-.08.19-.13.3-.13c.11 0 .22.05.3.13c.08.08.13.19.13.31s-.05.23-.13.31c-.08.08-.19.13-.3.13m2.57 0c-.11 0-.22-.05-.3-.13a.439.439 0 0 1-.13-.31c0-.12.04-.23.13-.31c.08-.08.19-.13.3-.13c.11 0 .22.05.3.13c.08.08.13.19.13.31s-.05.23-.13.31c-.08.08-.19.13-.3.13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Apple(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.74 1.01s-1.08.01-2 1.03c-.92 1.01-.78 2.17-.76 2.2c.02.03 1.31.08 2.13-1.1c.82-1.18.66-2.09.63-2.13m2.86 10.27c-.04-.08-2-1.08-1.82-2.99c.18-1.92 1.44-2.44 1.46-2.5c.02-.06-.51-.69-1.08-1.01a3.08 3.08 0 0 0-1.35-.38c-.09 0-.42-.08-1.08.1c-.44.12-1.42.52-1.7.53c-.27.02-1.08-.46-1.95-.58c-.56-.11-1.15.11-1.57.29c-.42.17-1.23.66-1.79 1.96c-.56 1.3-.27 3.35-.06 3.99c.21.64.54 1.68 1.1 2.45c.5.86 1.15 1.46 1.43 1.66c.27.2 1.05.34 1.59.06c.43-.27 1.21-.42 1.52-.41c.31.01.91.13 1.54.47c.49.17.96.1 1.42-.09c.47-.19 1.14-.93 1.93-2.41c.3-.69.44-1.06.41-1.12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arrowdown(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 13c-.28 0-.5-.22-.5-.5v-9c0-.28.22-.5.5-.5s.5.22.5.5v9c0 .28-.22.5-.5.5"/><path fill="currentColor" d="M4.5 14a.47.47 0 0 1-.35-.15l-3.5-3.5c-.2-.2-.2-.51 0-.71c.2-.2.51-.2.71 0l3.15 3.15l3.15-3.15c.2-.2.51-.2.71 0c.2.2.2.51 0 .71l-3.5 3.5c-.1.1-.23.15-.35.15Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arrowleft(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 5h-9c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h9c.28 0 .5.22.5.5s-.22.5-.5.5"/><path fill="currentColor" d="M6 8.5a.47.47 0 0 1-.35-.15l-3.5-3.5c-.2-.2-.2-.51 0-.71L5.65.65c.2-.2.51-.2.71 0c.2.2.2.51 0 .71L3.21 4.51l3.15 3.15c.2.2.2.51 0 .71c-.1.1-.23.15-.35.15Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arrowright(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 5h-9c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h9c.28 0 .5.22.5.5s-.22.5-.5.5"/><path fill="currentColor" d="M10 8.5a.47.47 0 0 1-.35-.15c-.2-.2-.2-.51 0-.71l3.15-3.15l-3.15-3.15c-.2-.2-.2-.51 0-.71c.2-.2.51-.2.71 0l3.5 3.5c.2.2.2.51 0 .71l-3.5 3.5c-.1.1-.23.15-.35.15Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arrowup(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 14c-.28 0-.5-.22-.5-.5v-9c0-.28.22-.5.5-.5s.5.22.5.5v9c0 .28-.22.5-.5.5"/><path fill="currentColor" d="M8 7.5a.47.47 0 0 1-.35-.15L4.5 4.2L1.35 7.35c-.2.2-.51.2-.71 0c-.2-.2-.2-.51 0-.71l3.5-3.5c.2-.2.51-.2.71 0l3.5 3.5c.2.2.2.51 0 .71c-.1.1-.23.15-.35.15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Avatarman(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 10h5c1.66 0 3 1.34 3 3v2H1v-2c0-1.66 1.34-3 3-3m0-6h5v2.5a2.5 2.5 0 0 1-5 0zm5 0c.55 0 1 .45 1 1s-.45 1-1 1zM4 6c-.55 0-1-.45-1-1s.45-1 1-1z"/><path fill="currentColor" d="M4.12 4.12h-.5C2.87 3.5 3 2.55 3 1.75s.5-.26 1-.26s1 .5 1 .5c-.88 0-1-1-1-1h3c1.1 0 2 .9 2 2h.5s.25.75-.12 1.25l-5.25-.12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Avatarwoman(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 10h5c1.66 0 3 1.34 3 3v2H1v-2c0-1.66 1.34-3 3-3m0-6h5v2.5a2.5 2.5 0 0 1-5 0z"/><path fill="currentColor" d="M9.5 3.93V3.5c0-1.1-.9-2-2-2L7 1H5.5c-1.1 0-2 .9-2 2v.93c0 1.33-.53 2.6-1.46 3.54l-.54.54l.74.37c.76.38 1.67.23 2.26-.37h4c.6.6 1.51.75 2.26.37l.74-.37l-.54-.54A4.994 4.994 0 0 1 9.5 3.93"/><path fill="currentColor" d="M7 1s.5 0 .5.5H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bitcoin(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.19 7.47c.76-.39 1.24-1.07 1.13-2.21c-.15-1.56-1.49-2.08-3.19-2.23V1H6.82v2H5.77V1H4.46v2c-.29 0-.56.01-.84.01H2v1.43s.77-.02.75 0c.53 0 .71.31.76.58v2.46h.14h-.14v3.44c-.02.17-.12.44-.49.44c.02.01-.75 0-.75 0l-.26 1.63h2.44v2h1.31v-2h1.05v2h1.31v-2.05c2.21-.13 3.75-.68 3.95-2.76c.16-1.67-.63-2.42-1.88-2.72Zm-4.4-2.94h.27c.92-.02 2.8-.07 2.8 1.32C8.86 7.21 6.9 7.18 6 7.16h-.22V4.54Zm3.68 5.39c0 1.5-2.35 1.46-3.43 1.45h-.26V8.48h.32c1.1-.03 3.36-.08 3.36 1.45Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bnb(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.67 7.04L8 4.71l2.33 2.33l1.36-1.36L8 1.99L4.31 5.68zM2 8l1.36-1.36L4.72 8L3.36 9.36zm6 3.29L5.67 8.96l-1.36 1.35L8 14l3.69-3.69l-1.36-1.36L8 11.28ZM11.29 8l1.36-1.36L14.01 8l-1.36 1.36zM8 6.62L9.38 8L8 9.38L6.62 8.01l.24-.25l.12-.12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.7 14.01c-.23 0-.45-.06-.66-.18l-3.29-1.94a.515.515 0 0 0-.51 0l-3.29 1.94c-.41.24-.89.24-1.3.01c-.41-.23-.65-.66-.65-1.13V3.49c0-.83.67-1.5 1.5-1.5h8c.83 0 1.5.67 1.5 1.5v9.22a1.293 1.293 0 0 1-1.29 1.3Zm-4.2-3.2c.26 0 .53.07.76.21l3.29 1.94c.13.08.25.03.3 0c.05-.03.15-.1.15-.26V3.49c0-.28-.22-.5-.5-.5h-8c-.28 0-.5.22-.5.5v9.22c0 .16.1.23.15.26c.05.03.16.08.3 0l3.29-1.94c.24-.14.5-.21.76-.21Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Button(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 3h36a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1M0 4a4 4 0 0 1 4-4h36a4 4 0 0 1 4 4v16a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4zm12 6.5a1.5 1.5 0 0 0 0 3h20a1.5 1.5 0 0 0 0-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cardano(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.34 2.03c.16-.31-.21-.68-.52-.52c-.25.1-.31.46-.12.65c.18.19.54.12.63-.13Zm-3.4.14c.02-.2-.22-.38-.4-.3c-.26.07-.29.48-.04.58c.19.11.46-.06.44-.28m6.32.29c-.29-.04-.34-.5-.06-.59c.22-.11.41.09.45.3c-.05.18-.19.34-.39.29M5.99 3.75c.05-.36-.39-.66-.7-.46c-.29.15-.33.6-.07.79c.27.25.76.04.77-.33m4.03-.25c.14-.36.7-.37.86-.02c.16.27-.04.61-.31.7c-.37.07-.72-.33-.54-.68ZM8 3.6c-.29.02-.57.25-.57.56c-.02.25.16.46.37.56h.06c.12.01.24.02.35-.04c.28-.12.42-.5.26-.77c-.08-.2-.29-.28-.49-.32ZM2.55 4.61c.28-.18.66.15.56.46c-.06.27-.42.38-.62.21c-.22-.16-.18-.55.06-.66Zm11.09.33c.03-.31-.37-.52-.61-.33c-.25.15-.22.56.04.68c.24.14.58-.07.56-.35Zm-4.84.22c.4-.14.87.03 1.09.39c.3.44.14 1.1-.32 1.35c-.48.3-1.18.05-1.36-.5c-.2-.48.09-1.09.59-1.24m-1.08.47c-.21-.45-.8-.63-1.24-.42c-.47.21-.68.83-.44 1.29c.22.47.85.67 1.29.41c.45-.23.64-.84.39-1.29Zm-3.57.38c.03-.29.3-.47.57-.49c.29.04.52.26.55.56c-.02.29-.24.58-.55.58c-.34.03-.64-.31-.57-.65m7.69.07c.02-.4-.47-.7-.82-.5c-.37.18-.4.77-.04.98c.35.25.88-.05.86-.48M5.53 7.09c.41-.11.87.08 1.07.46c.22.38.13.91-.21 1.19c-.44.41-1.24.24-1.47-.31c-.27-.51.06-1.2.61-1.33Zm5.48.4c-.21-.36-.68-.52-1.07-.4c-.56.15-.85.87-.56 1.37c.24.49.92.65 1.36.34c.42-.27.55-.89.27-1.31m-8.1.07c.35-.13.74.26.6.61c-.09.32-.53.44-.77.2c-.27-.22-.17-.72.17-.81m10.03-.06c-.3.02-.51.28-.5.57c.06.14.13.28.27.35c.29.16.71-.07.69-.42c.02-.26-.22-.46-.46-.5m-11.76.23c.21-.09.47.09.42.33c-.01.27-.4.37-.55.15c-.13-.16-.05-.4.12-.48Zm13.78.14c-.06-.18-.3-.25-.45-.15c-.25.13-.15.6.15.57c.21.03.4-.23.3-.42M6.67 9.01c.56-.15 1.15.32 1.16.9c.04.59-.55 1.1-1.12.97a.947.947 0 0 1-.77-.95c0-.43.32-.83.74-.92Zm3.38.91c0-.59-.61-1.07-1.17-.91c-.33.08-.62.36-.69.71c-.15.51.24 1.1.77 1.17c.56.11 1.13-.39 1.1-.96Zm-5.42-.55c.39-.07.76.37.6.75c-.12.4-.69.51-.95.18c-.3-.31-.08-.89.35-.93m7.21.56c.03-.38-.4-.69-.75-.54c-.42.14-.49.78-.11 1c.34.25.87-.04.86-.47Zm1.06 1.29c-.14-.26.11-.6.39-.55c.13 0 .23.1.31.19v.02l.02.03c.02.15.04.31-.08.43c-.17.22-.55.16-.65-.1Zm-9.77-.08c.06-.31-.32-.58-.58-.4c-.25.13-.26.53-.02.68c.23.16.57 0 .6-.28m4.63.18c.35-.15.79.15.75.53c.03.42-.52.72-.86.46c-.38-.22-.31-.86.1-1Zm-1.75 1.02c.03-.33-.37-.6-.66-.45c-.21.08-.29.31-.3.52c.07.19.22.39.44.4c.27.04.53-.19.52-.47m4.23-.44c.3-.19.73.08.7.44c0 .37-.48.61-.76.37c-.27-.19-.24-.65.06-.8Zm1.09 1.62c-.22.05-.38.29-.24.5c.14.24.54.15.57-.13c.04-.21-.15-.34-.32-.38Zm-7 .32c.05-.16.18-.33.37-.28c.28.02.36.44.12.57c-.22.15-.47-.05-.49-.29m3.83-.04c-.2-.13-.51-.01-.57.23c-.1.24.14.54.39.51c.14 0 .25-.1.33-.2c.02-.07.03-.13.05-.2c-.03-.13-.07-.28-.21-.34Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Caretdown(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.18 7.62L3.03 2.79c-.6-.7-.1-1.79.82-1.79h8.29c.93 0 1.42 1.09.82 1.79L8.82 7.62c-.43.5-1.21.5-1.64 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Caretleft(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1.38 7.18l4.83-4.15c.7-.6 1.79-.1 1.79.82v8.29c0 .93-1.09 1.42-1.79.82L1.38 8.82c-.5-.43-.5-1.21 0-1.64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Caretright(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.62 7.18L2.79 3.03c-.7-.6-1.79-.1-1.79.82v8.29c0 .93 1.09 1.42 1.79.82l4.83-4.14c.5-.43.5-1.21 0-1.64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Caretup(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.18 1.38L3.03 6.21c-.6.7-.1 1.79.82 1.79h8.29c.93 0 1.42-1.09.82-1.79L8.82 1.38c-.43-.5-1.21-.5-1.64 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M26.99 0L10.13 17.17l-5.44-5.54L0 16.41L10.4 27l4.65-4.73l.04.04L32 5.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Checkbox(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<rect width="41" height="41" x="1.5" y="1.5" fill="none" stroke="currentColor" stroke-width="3" rx="2.5"/><path fill="currentColor" d="m17.758 26.254l13.198-13.211l2.36 2.358l-15.557 15.557l-7.071-7.071l2.358-2.358z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circle(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="16" cy="16" r="16" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Close(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 12.5a.47.47 0 0 1-.35-.15l-8-8c-.2-.2-.2-.51 0-.71c.2-.2.51-.2.71 0l7.99 8.01c.2.2.2.51 0 .71c-.1.1-.23.15-.35.15Z"/><path fill="currentColor" d="M2 12.5a.47.47 0 0 1-.35-.15c-.2-.2-.2-.51 0-.71l8-7.99c.2-.2.51-.2.71 0c.2.2.2.51 0 .71l-8.01 7.99c-.1.1-.23.15-.35.15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Color(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 14h-3c-.28 0-.5-.22-.5-.5v-3c0-.13.05-.26.15-.35l5-5c.2-.2.51-.2.71 0l3 3c.2.2.2.51 0 .71l-5 5a.51.51 0 0 1-.35.15ZM3 13h2.29l4.5-4.5L7.5 6.21L3 10.71zm10.37-7.38L11.99 7l-3-3l1.38-1.38c.42-.42.96-.62 1.5-.62s1.08.2 1.5.62c.83.83.83 2.17 0 3"/><path fill="currentColor" d="M12.5 8a.47.47 0 0 1-.35-.15l-4-4c-.2-.2-.2-.51 0-.71c.2-.2.51-.2.71 0l4 4c.2.2.2.51 0 .71c-.1.1-.23.15-.35.15Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compress(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 5.09c-.38 0-.77-.15-1.06-.44l-2.79-2.8c-.2-.2-.2-.51 0-.71s.51-.2.71 0l2.79 2.79c.19.19.51.19.71 0l2.79-2.79c.2-.2.51-.2.71 0c.2.2.2.51 0 .71L9.07 4.64c-.29.29-.68.44-1.06.44ZM11.5 14a.47.47 0 0 1-.35-.15l-2.79-2.79a.513.513 0 0 0-.71 0l-2.79 2.79c-.2.2-.51.2-.71 0s-.2-.51 0-.71l2.79-2.79c.57-.57 1.55-.57 2.12 0l2.79 2.79c.2.2.2.51 0 .71c-.1.1-.23.15-.35.15m2-6h-11c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h11c.28 0 .5.22.5.5s-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Date(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 16h-13C.67 16 0 15.33 0 14.5v-12C0 1.67.67 1 1.5 1h13c.83 0 1.5.67 1.5 1.5v12c0 .83-.67 1.5-1.5 1.5M1.5 2c-.28 0-.5.22-.5.5v12c0 .28.22.5.5.5h13c.28 0 .5-.22.5-.5v-12c0-.28-.22-.5-.5-.5z"/><path fill="currentColor" d="M4.5 4c-.28 0-.5-.22-.5-.5v-3c0-.28.22-.5.5-.5s.5.22.5.5v3c0 .28-.22.5-.5.5m7 0c-.28 0-.5-.22-.5-.5v-3c0-.28.22-.5.5-.5s.5.22.5.5v3c0 .28-.22.5-.5.5m4 2H.5C.22 6 0 5.78 0 5.5S.22 5 .5 5h15c.28 0 .5.22.5.5s-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Datetime(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.5 14h-8C.67 14 0 13.33 0 12.5V2.38C0 1.55.67.88 1.5.88h11c.83 0 1.5.67 1.5 1.5v7.25c0 .28-.22.5-.5.5s-.5-.22-.5-.5V2.38c0-.28-.22-.5-.5-.5h-11c-.28 0-.5.22-.5.5V12.5c0 .28.22.5.5.5h8c.28 0 .5.22.5.5s-.22.5-.5.5"/><path fill="currentColor" d="M4 3.62c-.28 0-.5-.22-.5-.5V.5c0-.28.22-.5.5-.5s.5.22.5.5v2.62c0 .28-.22.5-.5.5m6.12 0c-.28 0-.5-.22-.5-.5V.5c0-.28.22-.5.5-.5s.5.22.5.5v2.62c0 .28-.22.5-.5.5M13.5 6H.5C.22 6 0 5.78 0 5.5S.22 5 .5 5h13c.28 0 .5.22.5.5s-.22.5-.5.5m-1 10C10.57 16 9 14.43 9 12.5S10.57 9 12.5 9s3.5 1.57 3.5 3.5s-1.57 3.5-3.5 3.5m0-6a2.5 2.5 0 0 0 0 5a2.5 2.5 0 0 0 0-5"/><path fill="currentColor" d="M13.5 14a.47.47 0 0 1-.35-.15l-1-1a.51.51 0 0 1-.15-.35V11c0-.28.22-.5.5-.5s.5.22.5.5v1.29l.85.85c.2.2.2.51 0 .71c-.1.1-.23.15-.35.15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Discover(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 13h-13C.67 13 0 12.33 0 11.5v-9C0 1.67.67 1 1.5 1h13c.83 0 1.5.67 1.5 1.5v9c0 .83-.67 1.5-1.5 1.5M1.5 2c-.28 0-.5.22-.5.5v9c0 .28.22.5.5.5h13c.28 0 .5-.22.5-.5v-9c0-.28-.22-.5-.5-.5z"/><path fill="currentColor" d="M9.37 6.99c0-.26-.1-.52-.29-.71c-.2-.2-.45-.29-.71-.29c-.56 0-1.01.44-1.01 1.01s.45 1.01 1.01 1.01S9.38 7.56 9.38 7Zm-2.2.44c-.15.15-.29.21-.44.21c-.36 0-.63-.27-.63-.66c0-.19.07-.35.19-.49c.11-.1.25-.17.41-.17c.17 0 .31.06.46.21v-.45a.93.93 0 0 0-.46-.12c-.25.02-.48.12-.66.27c-.05.05-.1.1-.14.16c-.13.17-.21.37-.21.6c0 .56.45.99 1.01.99h.02c.15 0 .29-.04.46-.12v-.45s0 .01-.01.02v-.02Zm3.47-1.18l-.44 1.08l-.5-1.28h-.41L10.08 8h.21l.83-1.95h-.41l-.08.2Zm.65.65v1.05h1.06v-.31h-.68v-.52h.66v-.33h-.66v-.41h.68v-.33h-1.06zm-8.75-.85H2v1.9h.54c.29 0 .5-.08.68-.21c.21-.17.35-.45.35-.74c0-.56-.41-.95-1.03-.95m.45 1.43c-.12.1-.27.15-.5.15h-.12V6.37h.1c.23 0 .39.04.5.15c.14.12.21.29.21.46a.58.58 0 0 1-.19.48Zm1.12-1.43h-.37v1.9h.37zm.93.74c-.23-.08-.29-.14-.29-.23c0-.12.12-.21.27-.21c.12 0 .21.04.31.15l.19-.25a.812.812 0 0 0-.56-.21c-.33 0-.6.23-.6.54c0 .27.12.39.46.52c.15.06.21.08.25.12c.08.04.12.12.12.19c0 .15-.12.27-.29.27s-.33-.1-.41-.25l-.23.23c.17.25.39.37.66.37c.39 0 .66-.25.66-.62c0-.33-.12-.46-.54-.62m8.36.36c.29-.06.45-.25.45-.54c0-.35-.25-.56-.68-.56h-.56v1.9h.37V7.2h.06l.52.75h.45l-.6-.79Zm-.31-.23h-.12v-.58h.12c.23 0 .37.1.37.29s-.14.29-.37.29"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dogecoin(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.72 1.5H4V6H2l-.04 2H4v4.5h3.91c1.68 0 2.08-.2 2.89-.59c.06-.03.13-.06.2-.1c.95-.46 1.69-1.1 2.21-1.92c.53-.83.79-1.79.79-2.89s-.26-2.06-.79-2.89c-.53-.83-1.26-1.47-2.21-1.93c-.95-.46-2.04-.68-3.28-.68M6.75 8v2.41h.83c1.14 0 2.05-.31 2.73-.92c.68-.61 1.02-1.44 1.02-2.49s-.34-1.88-1.02-2.49c-.68-.61-1.59-.92-2.73-.92h-.83V6h2.58v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dollar(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 8.5h-.07c-1.14-.15-2.51-.68-2.51-2.38C5.42 5.1 6.1 4 7.99 4s2.51 1.36 2.57 2.08c.02.28-.18.52-.46.54a.508.508 0 0 1-.54-.45c-.02-.19-.19-1.16-1.58-1.16c-1.04 0-1.57.38-1.57 1.12c0 .38 0 1.18 1.64 1.39a.503.503 0 0 1-.07 1Z"/><path fill="currentColor" d="M8 12c-1.9 0-2.51-1.36-2.57-2.08c-.02-.28.18-.52.46-.54c.27-.02.51.18.54.45c.02.19.19 1.16 1.58 1.16c1.04 0 1.57-.38 1.57-1.12c0-.38 0-1.18-1.64-1.39a.503.503 0 0 1-.43-.56c.04-.27.29-.47.56-.43c1.14.15 2.51.68 2.51 2.38c0 1.02-.68 2.12-2.57 2.12Z"/><path fill="currentColor" d="M8 15c-3.86 0-7-3.14-7-7s3.14-7 7-7s7 3.14 7 7s-3.14 7-7 7M8 2C4.69 2 2 4.69 2 8s2.69 6 6 6s6-2.69 6-6s-2.69-6-6-6"/><path fill="currentColor" d="M8 13c-.28 0-.5-.22-.5-.5v-9c0-.28.22-.5.5-.5s.5.22.5.5v9c0 .28-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Down(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 6.5a.47.47 0 0 1-.35-.15l-4.5-4.5c-.2-.2-.2-.51 0-.71c.2-.2.51-.2.71 0l4.15 4.15l4.14-4.14c.2-.2.51-.2.71 0c.2.2.2.51 0 .71l-4.5 4.5c-.1.1-.23.15-.35.15Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 11c-.28 0-.5-.22-.5-.5v-9c0-.28.22-.5.5-.5s.5.22.5.5v9c0 .28-.22.5-.5.5m6 4h-12c-.28 0-.5-.22-.5-.5v-2c0-.28.22-.5.5-.5s.5.22.5.5V14h11v-1.5c0-.28.22-.5.5-.5s.5.22.5.5v2c0 .28-.22.5-.5.5"/><path fill="currentColor" d="M7.5 12a.47.47 0 0 1-.35-.15l-3.5-3.5c-.2-.2-.2-.51 0-.71s.51-.2.71 0l3.15 3.15l3.15-3.15c.2-.2.51-.2.71 0s.2.51 0 .71l-3.5 3.5c-.1.1-.23.15-.35.15Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Downloadcloud(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 11c-.28 0-.5-.22-.5-.5s.22-.5.5-.5c1.65 0 3-1.35 3-3s-1.35-3-3-3h-1.05c-.18 0-.34-.09-.43-.25C9.88 2.65 8.76 2 7.51 2c-1.93 0-3.5 1.57-3.5 3.5c0 .28-.22.5-.5.5h-.5c-1.1 0-2 .9-2 2s.9 2 2 2c.28 0 .5.22.5.5s-.22.5-.5.5c-1.65 0-3-1.35-3-3s1.35-3 3-3h.03c.25-2.25 2.16-4 4.47-4c1.49 0 2.89.76 3.72 2h.78c2.21 0 4 1.79 4 4s-1.79 4-4 4Z"/><path fill="currentColor" d="M7.5 13a.47.47 0 0 1-.35-.15l-2.5-2.5c-.2-.2-.2-.51 0-.71c.2-.2.51-.2.71 0l2.15 2.15l2.15-2.15c.2-.2.51-.2.71 0s.2.51 0 .71l-2.5 2.5c-.1.1-.23.15-.35.15Z"/><path fill="currentColor" d="M7.5 12.75c-.28 0-.5-.22-.5-.5v-6c0-.28.22-.5.5-.5s.5.22.5.5v6c0 .28-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Email(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 13h-13C.67 13 0 12.33 0 11.5v-9C0 1.67.67 1 1.5 1h13c.83 0 1.5.67 1.5 1.5v9c0 .83-.67 1.5-1.5 1.5M1.5 2c-.28 0-.5.22-.5.5v9c0 .28.22.5.5.5h13c.28 0 .5-.22.5-.5v-9c0-.28-.22-.5-.5-.5z"/><path fill="currentColor" d="M8 8.96c-.7 0-1.34-.28-1.82-.79L.93 2.59c-.19-.2-.18-.52.02-.71c.2-.19.52-.18.71.02l5.25 5.58c.57.61 1.61.61 2.18 0l5.25-5.57c.19-.2.51-.21.71-.02c.2.19.21.51.02.71L9.82 8.18c-.48.51-1.12.79-1.82.79Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ethereum(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5 1l-.09.31v8.88l.09.09l4-2.44z"/><path fill="currentColor" d="M5 1L1 7.84l4 2.44zm0 10.62l-.05.06v3.17L5 15l4-5.81l-4 2.44Z"/><path fill="currentColor" d="M5 15v-3.38L1 9.18l4 5.81Zm0-4.72l4-2.44l-4-1.87zM1 7.84l4 2.44V5.97z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Euro(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 15c-3.86 0-7-3.14-7-7s3.14-7 7-7s7 3.14 7 7s-3.14 7-7 7M8 2C4.69 2 2 4.69 2 8s2.69 6 6 6s6-2.69 6-6s-2.69-6-6-6"/><path fill="currentColor" d="M7.89 12c-1.35 0-2.44-1.1-2.44-2.44V6.45c0-1.35 1.1-2.44 2.44-2.44s2.44 1.1 2.44 2.44c0 .28-.22.5-.5.5s-.5-.22-.5-.5c0-.8-.65-1.44-1.44-1.44s-1.44.65-1.44 1.44v3.11c0 .8.65 1.44 1.44 1.44s1.44-.65 1.44-1.44c0-.28.22-.5.5-.5s.5.22.5.5c0 1.35-1.1 2.44-2.44 2.44"/><path fill="currentColor" d="M7.5 7.5H5.17c-.28 0-.5-.22-.5-.5s.22-.5.5-.5H7.5c.28 0 .5.22.5.5s-.22.5-.5.5m0 2H5.17c-.28 0-.5-.22-.5-.5s.22-.5.5-.5H7.5c.28 0 .5.22.5.5s-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Expand(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 7a.47.47 0 0 1-.35-.15l-4.5-4.5c-.2-.2-.2-.51 0-.71s.51-.2.71 0l4.49 4.51c.2.2.2.51 0 .71c-.1.1-.23.15-.35.15Zm-4.44 7.44a.47.47 0 0 1-.35-.15c-.2-.2-.2-.51 0-.71l4.44-4.44c.2-.2.51-.2.71 0s.2.51 0 .71l-4.44 4.44c-.1.1-.23.15-.35.15Zm12.75-.13a.47.47 0 0 1-.35-.15L9.65 9.85c-.2-.2-.2-.51 0-.71s.51-.2.71 0l4.31 4.31c.2.2.2.51 0 .71c-.1.1-.23.15-.35.15Z"/><path fill="currentColor" d="M1 6.33c-.28 0-.5-.22-.5-.5V2.5C.5 1.67 1.17 1 2 1h3.67c.28 0 .5.22.5.5s-.22.5-.5.5H2c-.28 0-.5.22-.5.5v3.33c0 .28-.22.5-.5.5M5.33 15H2c-.83 0-1.5-.67-1.5-1.5V9.83c0-.28.22-.5.5-.5s.5.22.5.5v3.67c0 .28.22.5.5.5h3.33c.28 0 .5.22.5.5s-.22.5-.5.5M14 15h-3.67c-.28 0-.5-.22-.5-.5s.22-.5.5-.5H14c.28 0 .5-.22.5-.5v-3.33c0-.28.22-.5.5-.5s.5.22.5.5v3.33c0 .83-.67 1.5-1.5 1.5m1-8.33c-.28 0-.5-.22-.5-.5V2.5c0-.28-.22-.5-.5-.5h-3.33c-.28 0-.5-.22-.5-.5s.22-.5.5-.5H14c.83 0 1.5.67 1.5 1.5v3.67c0 .28-.22.5-.5.5"/><path fill="currentColor" d="M10 7a.47.47 0 0 1-.35-.15c-.2-.2-.2-.51 0-.71l4.5-4.49c.2-.2.51-.2.71 0s.2.51 0 .71l-4.5 4.5c-.1.1-.23.15-.35.15Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Export(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 14c-.28 0-.5-.22-.5-.5v-8c0-.28.22-.5.5-.5s.5.22.5.5v8c0 .28-.22.5-.5.5"/><path fill="currentColor" d="M10.5 9a.47.47 0 0 1-.35-.15L7.5 6.2L4.85 8.85c-.2.2-.51.2-.71 0s-.2-.51 0-.71l3-3c.2-.2.51-.2.71 0l3 3c.2.2.2.51 0 .71c-.1.1-.23.15-.35.15"/><path fill="currentColor" d="M10.07 12.94c-.19 0-.37-.11-.45-.29c-.11-.25 0-.55.25-.66c1.91-.87 3.14-2.74 3.14-4.77c0-2.91-2.47-5.28-5.5-5.28S2 4.31 2 7.22c0 2.02 1.23 3.9 3.14 4.77c.25.11.36.41.25.66c-.11.25-.41.36-.66.25c-2.26-1.03-3.72-3.26-3.72-5.68C1 3.75 3.92.94 7.5.94S14 3.76 14 7.22c0 2.41-1.46 4.64-3.72 5.68c-.07.03-.14.05-.21.05Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 11.5c-1.56 0-3.07-.61-4.5-1.8c-1.17-.99-1.99-2.13-2.37-2.73a.868.868 0 0 1 0-.93c.38-.6 1.2-1.75 2.37-2.73c2.84-2.39 6.15-2.39 8.99 0c1.17.99 1.99 2.13 2.37 2.73c.18.29.18.64 0 .93c-.38.6-1.2 1.75-2.37 2.73c-1.42 1.2-2.93 1.8-4.5 1.8Zm-5.97-5c.37.57 1.1 1.57 2.12 2.43c2.47 2.08 5.23 2.08 7.7 0c1.02-.86 1.75-1.87 2.12-2.43c-.37-.57-1.1-1.57-2.12-2.43c-2.47-2.08-5.23-2.08-7.7 0c-1.02.86-1.75 1.87-2.12 2.43"/><path fill="currentColor" d="M8 9a2.5 2.5 0 0 1 0-5a2.5 2.5 0 0 1 0 5m0-4c-.83 0-1.5.67-1.5 1.5S7.17 8 8 8s1.5-.67 1.5-1.5S8.83 5 8 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eyeclosed(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 11c-1.65 0-3-1.35-3-3s1.35-3 3-3s3 1.35 3 3s-1.35 3-3 3m0-5c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2"/><path fill="currentColor" d="M8 13c-3.19 0-5.99-1.94-6.97-4.84a.442.442 0 0 1 0-.32C2.01 4.95 4.82 3 8 3s5.99 1.94 6.97 4.84c.04.1.04.22 0 .32C13.99 11.05 11.18 13 8 13M2.03 8c.89 2.4 3.27 4 5.97 4s5.07-1.6 5.97-4C13.08 5.6 10.7 4 8 4S2.93 5.6 2.03 8"/><path fill="currentColor" d="M14 14.5a.47.47 0 0 1-.35-.15l-12-12c-.2-.2-.2-.51 0-.71c.2-.2.51-.2.71 0l11.99 12.01c.2.2.2.51 0 .71c-.1.1-.23.15-.35.15Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Facebook(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1C4.13 1 1 4.15 1 8.04c0 3.51 2.56 6.43 5.91 6.96v-4.92H5.13V8.04h1.78V6.49c0-1.77 1.05-2.74 2.64-2.74c.77 0 1.57.14 1.57.14v1.73h-.88c-.87 0-1.14.54-1.14 1.1v1.32h1.94l-.31 2.04H9.1V15c3.35-.53 5.91-3.44 5.91-6.96c0-3.89-3.13-7.04-7-7.04Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fastforward(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.09 11.01c-.18 0-.36-.05-.53-.14c-.35-.19-.57-.56-.57-.96V2.09c0-.4.22-.77.57-.96c.35-.19.78-.18 1.12.03l3.03 1.92c.23.15.3.46.15.69a.5.5 0 0 1-.69.15L3.14 2s-.07-.02-.1 0s-.05.05-.05.09v7.82s.02.07.05.09c.03.02.07.02.1 0l3.03-1.92a.5.5 0 0 1 .69.15a.5.5 0 0 1-.15.69l-3.03 1.92c-.18.11-.38.17-.59.17"/><path fill="currentColor" d="M7.04 10.98c-.18 0-.36-.04-.52-.13c-.36-.19-.58-.56-.58-.97V2.11c0-.41.22-.78.58-.97c.36-.19.79-.17 1.13.05l5.87 3.89c.31.2.49.55.49.92s-.18.71-.49.92l-5.87 3.89c-.18.12-.39.18-.61.18Zm0-8.97s-.03 0-.05.01c-.03.02-.05.05-.05.09v7.77s.02.07.05.09c.03.02.07.02.1 0l5.87-3.89l.28-.58l-.28.42l-5.87-3.89s-.04-.02-.05-.02"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func File(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 5.88c-.28 0-.5-.22-.5-.5V1.5c0-.28-.22-.5-.5-.5h-9c-.28 0-.5.22-.5.5v2c0 .28-.22.5-.5.5S2 3.78 2 3.5v-2C2 .67 2.67 0 3.5 0h9c.83 0 1.5.67 1.5 1.5v3.88c0 .28-.22.5-.5.5"/><path fill="currentColor" d="M14.5 16h-13C.67 16 0 15.33 0 14.5v-10C0 3.67.67 3 1.5 3h4.75c.16 0 .31.07.4.2L8 5h6.5c.83 0 1.5.67 1.5 1.5v8c0 .83-.67 1.5-1.5 1.5M1.5 4c-.28 0-.5.22-.5.5v10c0 .28.22.5.5.5h13c.28 0 .5-.22.5-.5v-8c0-.28-.22-.5-.5-.5H7.75a.48.48 0 0 1-.4-.2L6 4z"/><path fill="currentColor" d="M5.5 13h-2c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h2c.28 0 .5.22.5.5s-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fileaudio(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 16h-10c-.83 0-1.5-.67-1.5-1.5v-13C1 .67 1.67 0 2.5 0h7.09c.4 0 .78.16 1.06.44l2.91 2.91c.28.28.44.66.44 1.06V14.5c0 .83-.67 1.5-1.5 1.5M2.5 1c-.28 0-.5.22-.5.5v13c0 .28.22.5.5.5h10c.28 0 .5-.22.5-.5V4.41a.47.47 0 0 0-.15-.35L9.94 1.15A.51.51 0 0 0 9.59 1z"/><path fill="currentColor" d="M13.38 5h-2.91C9.66 5 9 4.34 9 3.53V.62c0-.28.22-.5.5-.5s.5.22.5.5v2.91c0 .26.21.47.47.47h2.91c.28 0 .5.22.5.5s-.22.5-.5.5M5 6h5v2H5z"/><path fill="currentColor" d="M9.5 11c-.28 0-.5-.22-.5-.5v-3c0-.28.22-.5.5-.5s.5.22.5.5v3c0 .28-.22.5-.5.5m-4 1c-.28 0-.5-.22-.5-.5v-4c0-.28.22-.5.5-.5s.5.22.5.5v4c0 .28-.22.5-.5.5"/><path fill="currentColor" d="M8 10h2v1c0 .55-.45 1-1 1H8c-.55 0-1-.45-1-1s.45-1 1-1m-4 1h2v1c0 .55-.45 1-1 1H4c-.55 0-1-.45-1-1s.45-1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filedoc(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 16h-10c-.83 0-1.5-.67-1.5-1.5v-13C1 .67 1.67 0 2.5 0h7.09c.4 0 .78.16 1.06.44l2.91 2.91c.28.28.44.66.44 1.06V14.5c0 .83-.67 1.5-1.5 1.5M2.5 1c-.28 0-.5.22-.5.5v13c0 .28.22.5.5.5h10c.28 0 .5-.22.5-.5V4.41a.47.47 0 0 0-.15-.35L9.94 1.15A.51.51 0 0 0 9.59 1z"/><path fill="currentColor" d="M13.38 5h-2.91C9.66 5 9 4.34 9 3.53V.62c0-.28.22-.5.5-.5s.5.22.5.5v2.91c0 .26.21.47.47.47h2.91c.28 0 .5.22.5.5s-.22.5-.5.5M10 13H5c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h5c.28 0 .5.22.5.5s-.22.5-.5.5m0-3H5c-.28 0-.5-.22-.5-.5S4.72 9 5 9h5c.28 0 .5.22.5.5s-.22.5-.5.5M7 7H5c-.28 0-.5-.22-.5-.5S4.72 6 5 6h2c.28 0 .5.22.5.5S7.28 7 7 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fileimage(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.22 8.14l-3.17 4.28c-.18.24 0 .58.29.58h6.32c.3 0 .47-.34.29-.58L6.8 8.14c-.14-.2-.44-.2-.58 0"/><path fill="currentColor" d="m9.72 10.09l-1.9 2.57c-.11.14 0 .35.17.35h3.79c.18 0 .28-.2.17-.35l-1.89-2.57a.215.215 0 0 0-.35 0Z"/><circle cx="10" cy="8" r="1" fill="currentColor"/><path fill="currentColor" d="M12.5 16h-10c-.83 0-1.5-.67-1.5-1.5v-13C1 .67 1.67 0 2.5 0h7.09c.4 0 .78.16 1.06.44l2.91 2.91c.28.28.44.66.44 1.06V14.5c0 .83-.67 1.5-1.5 1.5M2.5 1c-.28 0-.5.22-.5.5v13c0 .28.22.5.5.5h10c.28 0 .5-.22.5-.5V4.41a.47.47 0 0 0-.15-.35L9.94 1.15A.51.51 0 0 0 9.59 1z"/><path fill="currentColor" d="M13.38 5h-2.91C9.66 5 9 4.34 9 3.53V.62c0-.28.22-.5.5-.5s.5.22.5.5v2.91c0 .26.21.47.47.47h2.91c.28 0 .5.22.5.5s-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filepdf(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 13h.86v-.9h.39c.62 0 1.14-.45 1.14-1.06s-.5-1.05-1.14-1.05H3v3Zm.86-1.59v-.72h.3c.2 0 .37.13.37.35s-.16.36-.37.36h-.3ZM6.19 13h1.19c1 0 1.62-.59 1.62-1.52C9 10.61 8.38 10 7.38 10H6.19zm.86-.71V10.7h.29c.33 0 .78.16.78.78c0 .65-.45.81-.78.81zM10 13h.86v-1.07h1.06v-.69h-1.06v-.54h1.21v-.69h-2.06v3Z"/><path fill="currentColor" d="M12.5 16h-10c-.83 0-1.5-.67-1.5-1.5v-13C1 .67 1.67 0 2.5 0h7.09c.4 0 .78.16 1.06.44l2.91 2.91c.28.28.44.66.44 1.06V14.5c0 .83-.67 1.5-1.5 1.5M2.5 1c-.28 0-.5.22-.5.5v13c0 .28.22.5.5.5h10c.28 0 .5-.22.5-.5V4.41a.47.47 0 0 0-.15-.35L9.94 1.15A.51.51 0 0 0 9.59 1z"/><path fill="currentColor" d="M13.38 5h-2.91C9.66 5 9 4.34 9 3.53V.62c0-.28.22-.5.5-.5s.5.22.5.5v2.91c0 .26.21.47.47.47h2.91c.28 0 .5.22.5.5s-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filevideo(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.65 11.94l4.14-2.56c.28-.17.28-.59 0-.76L5.65 6.06c-.29-.18-.65.04-.65.38v5.11c0 .34.36.56.65.38Z"/><path fill="currentColor" d="M12.5 16h-10c-.83 0-1.5-.67-1.5-1.5v-13C1 .67 1.67 0 2.5 0h7.09c.4 0 .78.16 1.06.44l2.91 2.91c.28.28.44.66.44 1.06V14.5c0 .83-.67 1.5-1.5 1.5M2.5 1c-.28 0-.5.22-.5.5v13c0 .28.22.5.5.5h10c.28 0 .5-.22.5-.5V4.41a.47.47 0 0 0-.15-.35L9.94 1.15A.51.51 0 0 0 9.59 1z"/><path fill="currentColor" d="M13.38 5h-2.91C9.66 5 9 4.34 9 3.53V.62c0-.28.22-.5.5-.5s.5.22.5.5v2.91c0 .26.21.47.47.47h2.91c.28 0 .5.22.5.5s-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.36 10H1.5c-.28 0-.5-.22-.5-.5v-7c0-.28.22-.5.5-.5h9.86c.19 0 .42.14.51.31c.1.19.08.42-.04.59L9.62 5.99l2.21 3.09c.11.15.13.43.04.59a.57.57 0 0 1-.51.31ZM2 9h8.53l-1.9-2.67c-.12-.17-.12-.49 0-.67l1.9-2.67H2v6Z"/><path fill="currentColor" d="M1.5 14c-.28 0-.5-.22-.5-.5V4.12c0-.28.22-.5.5-.5s.5.22.5.5v9.38c0 .28-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 5.88c-.28 0-.5-.22-.5-.5V1.5c0-.28-.22-.5-.5-.5h-9c-.28 0-.5.22-.5.5v2c0 .28-.22.5-.5.5S2 3.78 2 3.5v-2C2 .67 2.67 0 3.5 0h9c.83 0 1.5.67 1.5 1.5v3.88c0 .28-.22.5-.5.5"/><path fill="currentColor" d="M14.5 16h-13C.67 16 0 15.33 0 14.5v-10C0 3.67.67 3 1.5 3h4.75c.16 0 .31.07.4.2L8 5h6.5c.83 0 1.5.67 1.5 1.5v8c0 .83-.67 1.5-1.5 1.5M1.5 4c-.28 0-.5.22-.5.5v10c0 .28.22.5.5.5h13c.28 0 .5-.22.5-.5v-8c0-.28-.22-.5-.5-.5H7.75a.48.48 0 0 1-.4-.2L6 4z"/><path fill="currentColor" d="M5.5 13h-2c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h2c.28 0 .5.22.5.5s-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Franc(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 15c-3.86 0-7-3.14-7-7s3.14-7 7-7s7 3.14 7 7s-3.14 7-7 7M8 2C4.69 2 2 4.69 2 8s2.69 6 6 6s6-2.69 6-6s-2.69-6-6-6"/><path fill="currentColor" d="M6.5 12.12c-.28 0-.5-.22-.5-.5V4.5c0-.28.22-.5.5-.5h4c.28 0 .5.22.5.5s-.22.5-.5.5H7v6.62c0 .28-.22.5-.5.5"/><path fill="currentColor" d="M8.62 11H5.37c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h3.25c.28 0 .5.22.5.5s-.22.5-.5.5m.88-3H6.62c-.28 0-.5-.22-.5-.5s.22-.5.5-.5H9.5c.28 0 .5.22.5.5s-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Github(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1C4.13 1 1 4.21 1 8.18c0 3.14 1.96 5.81 4.69 6.79c.37.09.31-.17.31-.36v-1.25c-2.12.26-2.21-1.19-2.35-1.43c-.29-.5-.97-.63-.76-.87c.48-.26.98.06 1.55.93c.41.63 1.22.52 1.63.42c.09-.38.28-.71.54-.98c-2.2-.4-3.12-1.78-3.12-3.42c0-.8.25-1.53.76-2.12c-.32-.97.03-1.8.08-1.93c.91-.08 1.85.67 1.93.73c.52-.14 1.11-.22 1.77-.22s1.25.08 1.78.22c.18-.14 1.05-.78 1.9-.71c.05.12.39.94.09 1.9c.51.59.76 1.33.76 2.13c0 1.64-.92 3.02-3.13 3.42a2 2 0 0 1 .44.67c.1.25.15.52.15.79v1.81c.01.14 0 .29.23.29c2.77-.96 4.76-3.65 4.76-6.81c0-3.97-3.13-7.19-7-7.19Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Google(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.14 8.33v-2.6h6.69c.1.44.18.85.18 1.44c0 4-2.74 6.84-6.86 6.84S0 10.86 0 7s3.2-7 7.14-7c1.93 0 3.54.69 4.78 1.83L9.89 3.76c-.51-.48-1.41-1.04-2.75-1.04c-2.36 0-4.29 1.93-4.29 4.28s1.93 4.28 4.29 4.28c2.74 0 3.74-1.86 3.93-2.95H7.13Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GroupIcon(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 8.5c-.08 0-.15-.02-.22-.05L1.28 5.2a.5.5 0 0 1 0-.9l6.5-3.25c.14-.07.31-.07.45 0l6.5 3.25a.5.5 0 0 1 0 .9l-6.5 3.25c-.07.04-.15.05-.22.05ZM2.62 4.75L8 7.44l5.38-2.69L8 2.06z"/><path fill="currentColor" d="M8 11.75c-.08 0-.15-.02-.22-.05l-6.5-3.25a.5.5 0 0 1 0-.9l3.25-1.62c.25-.12.55-.02.67.22c.12.25.02.55-.22.67L2.62 8L8 10.69L13.38 8l-2.36-1.18a.488.488 0 0 1-.22-.67c.12-.25.42-.35.67-.22l3.25 1.62a.5.5 0 0 1 0 .9l-6.5 3.25c-.07.04-.15.05-.22.05"/><path fill="currentColor" d="M8 15c-.08 0-.15-.02-.22-.05l-6.5-3.25a.5.5 0 0 1 0-.9l3.25-1.62c.25-.12.55-.02.67.22c.12.25.02.55-.22.67l-2.36 1.18L8 13.94l5.38-2.69l-2.36-1.18a.488.488 0 0 1-.22-.67c.12-.25.42-.35.67-.22l3.25 1.62a.5.5 0 0 1 0 .9l-6.5 3.25c-.07.04-.15.05-.22.05"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Happy(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="6" cy="6" r="1" fill="currentColor"/><circle cx="10" cy="6" r="1" fill="currentColor"/><path fill="currentColor" d="M8 15c-3.86 0-7-3.14-7-7s3.14-7 7-7s7 3.14 7 7s-3.14 7-7 7M8 2C4.69 2 2 4.69 2 8s2.69 6 6 6s6-2.69 6-6s-2.69-6-6-6"/><path fill="currentColor" d="M8 11.5c-1.48 0-2.81-.94-3.3-2.33a.501.501 0 0 1 .94-.34a2.502 2.502 0 0 0 4.72 0c.09-.26.38-.4.64-.3c.26.09.4.38.3.64A3.494 3.494 0 0 1 8 11.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 14c-.2 0-.4-.08-.56-.23c-1.06-1.04-4.58-4.59-5.49-6.34c-.63-1.2-.59-2.7.09-3.83c.61-1.01 1.67-1.59 2.9-1.59c1.56 0 2.53.81 3.06 1.63c.53-.82 1.5-1.63 3.06-1.63c1.23 0 2.29.58 2.9 1.59c.68 1.13.72 2.63.09 3.83c-.92 1.76-4.43 5.3-5.49 6.34c-.16.16-.36.23-.56.23M4.44 3c-.88 0-1.61.39-2.04 1.11c-.51.83-.53 1.95-.06 2.85c.66 1.26 3.07 3.88 5.17 5.96c2.09-2.08 4.51-4.69 5.17-5.96c.47-.9.44-2.02-.06-2.85C12.19 3.39 11.46 3 10.58 3C8.46 3 8.03 4.9 8.01 4.98h-.98C7.01 4.9 6.56 3 4.46 3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Help(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 15c-3.86 0-7-3.14-7-7s3.14-7 7-7s7 3.14 7 7s-3.14 7-7 7M8 2C4.69 2 2 4.69 2 8s2.69 6 6 6s6-2.69 6-6s-2.69-6-6-6"/><path fill="currentColor" d="M8 4.5c-1.11 0-2 .89-2 2h1c0-.55.45-1 1-1s1 .45 1 1c0 1-1.5.88-1.5 2.5h1c0-1.12 1.5-1.25 1.5-2.5c0-1.11-.89-2-2-2"/><circle cx="8" cy="11" r=".62" fill="currentColor"/><circle cx="6.5" cy="6.5" r=".5" fill="currentColor"/><circle cx="8" cy="9" r=".5" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hidden(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 11c-1.65 0-3-1.35-3-3s1.35-3 3-3s3 1.35 3 3s-1.35 3-3 3m0-5c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2"/><path fill="currentColor" d="M8 13c-3.19 0-5.99-1.94-6.97-4.84a.442.442 0 0 1 0-.32C2.01 4.95 4.82 3 8 3s5.99 1.94 6.97 4.84c.04.1.04.22 0 .32C13.99 11.05 11.18 13 8 13M2.03 8c.89 2.4 3.27 4 5.97 4s5.07-1.6 5.97-4C13.08 5.6 10.7 4 8 4S2.93 5.6 2.03 8"/><path fill="currentColor" d="M14 14.5a.47.47 0 0 1-.35-.15l-12-12c-.2-.2-.2-.51 0-.71c.2-.2.51-.2.71 0l11.99 12.01c.2.2.2.51 0 .71c-.1.1-.23.15-.35.15Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Info(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 15c-3.86 0-7-3.14-7-7s3.14-7 7-7s7 3.14 7 7s-3.14 7-7 7M8 2C4.69 2 2 4.69 2 8s2.69 6 6 6s6-2.69 6-6s-2.69-6-6-6"/><circle cx="8" cy="6" r=".75" fill="currentColor"/><path fill="currentColor" d="M8 12c-.28 0-.5-.22-.5-.5v-3c0-.28.22-.5.5-.5s.5.22.5.5v3c0 .28-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Instagram(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 5.67C6.71 5.67 5.67 6.72 5.67 8S6.72 10.33 8 10.33S10.33 9.28 10.33 8S9.28 5.67 8 5.67M15 8c0-.97 0-1.92-.05-2.89c-.05-1.12-.31-2.12-1.13-2.93c-.82-.82-1.81-1.08-2.93-1.13C9.92 1 8.97 1 8 1s-1.92 0-2.89.05c-1.12.05-2.12.31-2.93 1.13C1.36 3 1.1 3.99 1.05 5.11C1 6.08 1 7.03 1 8s0 1.92.05 2.89c.05 1.12.31 2.12 1.13 2.93c.82.82 1.81 1.08 2.93 1.13C6.08 15 7.03 15 8 15s1.92 0 2.89-.05c1.12-.05 2.12-.31 2.93-1.13c.82-.82 1.08-1.81 1.13-2.93c.06-.96.05-1.92.05-2.89m-7 3.59c-1.99 0-3.59-1.6-3.59-3.59S6.01 4.41 8 4.41s3.59 1.6 3.59 3.59s-1.6 3.59-3.59 3.59m3.74-6.49c-.46 0-.84-.37-.84-.84s.37-.84.84-.84s.84.37.84.84a.8.8 0 0 1-.24.59a.8.8 0 0 1-.59.24Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Krona(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 15c-3.86 0-7-3.14-7-7s3.14-7 7-7s7 3.14 7 7s-3.14 7-7 7M8 2C4.69 2 2 4.69 2 8s2.69 6 6 6s6-2.69 6-6s-2.69-6-6-6"/><path fill="currentColor" d="M5 4h1v8H5zm5 2.6h1V12h-1z"/><path fill="currentColor" d="M11 9h-1c0-1.1.9-2 2-2h.62v1H12c-.55 0-1 .45-1 1m-1.4 3L7.17 8.53L9.71 6H8.29L5.15 9.15l.7.7l.6-.6L8.38 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Left(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 13a.47.47 0 0 1-.35-.15l-4.5-4.5c-.2-.2-.2-.51 0-.71l4.5-4.49c.2-.2.51-.2.71 0c.2.2.2.51 0 .71L1.71 8l4.15 4.15c.2.2.2.51 0 .71c-.1.1-.23.15-.35.15Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.12 11c-.84 0-1.62-.32-2.2-.91a3.111 3.111 0 0 1 0-4.4l2.77-2.77a3.12 3.12 0 0 1 4.4 0a3.102 3.102 0 0 1 0 4.4c-.2.2-.51.2-.71 0c-.2-.2-.2-.51 0-.71c.82-.82.82-2.16 0-2.99c-.83-.83-2.17-.83-2.99 0L2.62 6.39c-.82.82-.83 2.16 0 2.99c.4.4.92.63 1.5.62c.56 0 1.09-.22 1.49-.62c.2-.2.51-.2.71 0c.2.2.2.51 0 .71c-.59.59-1.37.91-2.2.91"/><path fill="currentColor" d="M9.14 9.97c-.83 0-1.62-.32-2.2-.91a3.111 3.111 0 0 1 0-4.4c.2-.2.51-.19.71 0c.2.2.19.51 0 .71c-.82.82-.83 2.16 0 2.99c.4.4.95.63 1.5.62c.56 0 1.09-.22 1.49-.62l2.75-2.75c.4-.4.62-.93.62-1.49s-.22-1.1-.62-1.49c-.83-.83-2.17-.83-2.99 0c-.2.2-.51.19-.71 0c-.2-.2-.19-.51 0-.71a3.12 3.12 0 0 1 4.4 0c.59.59.91 1.37.91 2.2s-.32 1.61-.91 2.2l-2.75 2.75c-.59.59-1.37.91-2.2.91Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linkedin(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.44 4.89c.8 0 1.44-.65 1.44-1.44s-.65-1.44-1.44-1.44S2 2.66 2 3.45s.65 1.44 1.44 1.44m2.81 1.09V14h2.48v-3.96c0-1.05.2-2.06 1.49-2.06s1.29 1.2 1.29 2.12V14H14V9.6c0-2.16-.46-3.82-2.98-3.82c-1.21 0-2.02.66-2.35 1.29h-.03v-1.1H6.26Zm-4.05 0h2.49V14H2.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linkexternal(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.01 10.49a.47.47 0 0 1-.35-.15c-.2-.2-.2-.51 0-.71l8.49-8.48c.2-.2.51-.2.71 0c.2.2.2.51 0 .71l-8.5 8.48c-.1.1-.23.15-.35.15"/><path fill="currentColor" d="M14.5 7c-.28 0-.5-.22-.5-.5V2H9.5c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h5c.28 0 .5.22.5.5v5c0 .28-.22.5-.5.5m-3 8H2.49C1.67 15 1 14.33 1 13.51V4.49C1 3.67 1.67 3 2.49 3H7.5c.28 0 .5.22.5.5s-.22.5-.5.5H2.49a.49.49 0 0 0-.49.49v9.02c0 .27.22.49.49.49h9.01c.27 0 .49-.22.49-.49V8.5c0-.28.22-.5.5-.5s.5.22.5.5v5.01c0 .82-.67 1.49-1.49 1.49"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lira(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 15.5C3.86 15.5.5 12.14.5 8S3.86.5 8 .5s7.5 3.36 7.5 7.5s-3.36 7.5-7.5 7.5m0-14C4.42 1.5 1.5 4.42 1.5 8s2.92 6.5 6.5 6.5s6.5-2.92 6.5-6.5S11.58 1.5 8 1.5"/><path fill="currentColor" d="M7.5 13h-1c-.28 0-.5-.22-.5-.5v-9c0-.28.22-.5.5-.5s.5.22.5.5V12h.5A2.5 2.5 0 0 0 10 9.5c0-.28.22-.5.5-.5s.5.22.5.5c0 1.93-1.57 3.5-3.5 3.5"/><path fill="currentColor" d="M5 8.5c-.21 0-.4-.13-.47-.34c-.09-.26.05-.55.32-.63l4.5-1.5c.26-.09.55.05.63.32c.09.26-.05.55-.32.63l-4.5 1.5c-.05.02-.11.03-.16.03Zm0-2c-.21 0-.4-.13-.47-.34c-.09-.26.05-.55.32-.63l4.5-1.5c.26-.09.55.05.63.32c.09.26-.05.55-.32.63l-4.5 1.5c-.05.02-.11.03-.16.03Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func List(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="3" cy="2.5" r="1" fill="currentColor"/><circle cx="3" cy="5.5" r="1" fill="currentColor"/><circle cx="3" cy="8.5" r="1" fill="currentColor"/><path fill="currentColor" d="M13.5 3h-7c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h7c.28 0 .5.22.5.5s-.22.5-.5.5m0 3h-7c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h7c.28 0 .5.22.5.5s-.22.5-.5.5m0 3h-7c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h7c.28 0 .5.22.5.5s-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mastercard(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 13h-13C.67 13 0 12.33 0 11.5v-9C0 1.67.67 1 1.5 1h13c.83 0 1.5.67 1.5 1.5v9c0 .83-.67 1.5-1.5 1.5M1.5 2c-.28 0-.5.22-.5.5v9c0 .28.22.5.5.5h13c.28 0 .5-.22.5-.5v-9c0-.28-.22-.5-.5-.5z"/><path fill="currentColor" d="M9.35 4.57h-2.7v4.86h2.7z"/><path fill="currentColor" d="M6.82 7c0-.47.11-.93.31-1.35c.2-.42.5-.79.87-1.08c-.46-.36-1-.58-1.58-.64c-.58-.06-1.16.04-1.68.29c-.52.25-.96.65-1.27 1.14C3.16 5.85 3 6.42 3 7s.16 1.15.47 1.64c.31.49.75.89 1.27 1.14s1.1.35 1.68.29C7 10.01 7.54 9.79 8 9.43c-.37-.29-.67-.66-.87-1.08c-.2-.42-.31-.88-.31-1.35"/><path fill="currentColor" d="M13 7c0 .58-.16 1.15-.47 1.64c-.31.49-.75.89-1.27 1.14c-.52.25-1.1.35-1.68.29C9 10.01 8.46 9.79 8 9.43c.37-.29.67-.66.87-1.08c.2-.42.31-.88.31-1.35s-.11-.93-.31-1.35c-.2-.42-.5-.79-.87-1.08c.46-.36 1-.58 1.58-.64c.58-.06 1.16.04 1.68.29c.52.25.96.65 1.27 1.14S13 6.42 13 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Medium(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.66 3.04c0-.08 0-.15-.03-.22a.522.522 0 0 0-.14-.19L1.18 1.21V1h4.06l3.14 6.19L11.13 1H15v.21l-1.12.96s-.08.08-.11.13c-.02.05-.03.1-.02.16v7.08c0 .05 0 .11.02.16s.06.09.11.13l1.09.96V11H9.48v-.21l1.13-.99c.11-.1.11-.13.11-.28V3.79l-3.14 7.18h-.42L3.5 3.79v4.82c-.03.2.04.41.2.55l1.47 1.61v.21H1v-.21l1.47-1.61c.08-.07.14-.16.17-.26s.04-.2.02-.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Megaphone(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 12.46c-.15 0-.3-.02-.45-.07l-6-1.87C1.42 10.32 1 9.75 1 9.09V5.56c0-.66.42-1.23 1.05-1.43l6-1.87c.46-.14.95-.06 1.34.22c.39.29.61.73.61 1.21v7.28a1.492 1.492 0 0 1-1.49 1.5Zm0-9.28c-.05 0-.1 0-.15.02l-6 1.88c-.21.07-.35.26-.35.48v3.53c0 .22.14.41.35.48l6 1.87c.15.05.32.02.45-.07c.13-.1.2-.24.2-.4V3.68a.48.48 0 0 0-.2-.4a.534.534 0 0 0-.29-.1Zm5 4.82h-2c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h2c.28 0 .5.22.5.5s-.22.5-.5.5M11 6.5a.47.47 0 0 1-.35-.15c-.2-.2-.2-.51 0-.71l1.5-1.5c.2-.2.51-.2.71 0c.2.2.2.51 0 .71l-1.5 1.5c-.1.1-.23.15-.35.15Zm1.5 4.5a.47.47 0 0 1-.35-.15l-1.5-1.5c-.2-.2-.2-.51 0-.71c.2-.2.51-.2.71 0l1.5 1.5c.2.2.2.51 0 .71c-.1.1-.23.15-.35.15Z"/><path fill="currentColor" d="M4 13.82c-1.1 0-2-.9-2-2V10.2h1v1.62c0 .55.45 1 1 1s1-.45 1-1v-.4h1v.4c0 1.1-.9 2-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minimize(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 10a.47.47 0 0 1-.35-.15c-.2-.2-.2-.51 0-.71l6.5-6.49c.2-.2.51-.2.71 0c.2.2.2.51 0 .71l-6.5 6.5c-.1.1-.23.15-.35.15Z"/><path fill="currentColor" d="M9.5 10h-4c-.28 0-.5-.22-.5-.5v-4c0-.28.22-.5.5-.5s.5.22.5.5V9h3.5c.28 0 .5.22.5.5s-.22.5-.5.5"/><path fill="currentColor" d="M11.5 14h-9c-.83 0-1.5-.67-1.5-1.5v-9C1 2.67 1.67 2 2.5 2h5c.28 0 .5.22.5.5s-.22.5-.5.5h-5c-.28 0-.5.22-.5.5v9c0 .28.22.5.5.5h9c.28 0 .5-.22.5-.5v-5c0-.28.22-.5.5-.5s.5.22.5.5v5c0 .83-.67 1.5-1.5 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Month(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 16h-13C.67 16 0 15.33 0 14.5v-12C0 1.67.67 1 1.5 1h13c.83 0 1.5.67 1.5 1.5v12c0 .83-.67 1.5-1.5 1.5M1.5 2c-.28 0-.5.22-.5.5v12c0 .28.22.5.5.5h13c.28 0 .5-.22.5-.5v-12c0-.28-.22-.5-.5-.5z"/><path fill="currentColor" d="M4.5 4c-.28 0-.5-.22-.5-.5v-3c0-.28.22-.5.5-.5s.5.22.5.5v3c0 .28-.22.5-.5.5m7 0c-.28 0-.5-.22-.5-.5v-3c0-.28.22-.5.5-.5s.5.22.5.5v3c0 .28-.22.5-.5.5m4 2H.5C.22 6 0 5.78 0 5.5S.22 5 .5 5h15c.28 0 .5.22.5.5s-.22.5-.5.5M3.23 12h.88l.38-1h1.25l.38 1H7L5.5 8h-.73l-1.52 4Zm1.38-1.62l.5-1.5l.5 1.5zM8.19 12c.41 0 .7-.21.82-.43h.04V12h.69V9h-.72v1.68c0 .37-.25.65-.6.65s-.57-.24-.57-.6V9h-.72v1.91c0 .62.41 1.09 1.06 1.09m3.43 1.02c.92 0 1.38-.54 1.38-1.27V9.06h-.66v.35h-.04a.978.978 0 0 0-.83-.42c-.72 0-1.25.56-1.25 1.41s.54 1.4 1.27 1.4c.48 0 .74-.29.81-.41h.04v.35c0 .46-.24.72-.72.72c-.37 0-.58-.15-.63-.4h-.62c.05.5.47.95 1.25.95Zm0-1.81c-.43 0-.73-.31-.73-.8s.29-.8.72-.8s.73.3.73.8c0 .46-.28.8-.73.8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Multicurrency(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.73 16c-2.82 0-5.12-2.29-5.12-5.12a5.1 5.1 0 0 1 2.72-4.52c.24-.13.55-.04.68.21c.13.24.04.55-.21.68a4.116 4.116 0 0 0-2.19 3.64c0 2.27 1.85 4.12 4.12 4.12s4.12-1.85 4.12-4.12c0-.83-.25-1.63-.71-2.31c-.16-.23-.1-.54.13-.69c.23-.16.54-.1.69.13c.58.85.89 1.85.89 2.88c0 2.82-2.29 5.12-5.12 5.12Z"/><path fill="currentColor" d="M11.5 14H7.95c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h3.55c.28 0 .5.22.5.5s-.22.5-.5.5m-1.41-2H7.96c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h2.13c.28 0 .5.22.5.5s-.22.5-.5.5"/><path fill="currentColor" d="M8.67 13.87c-.28 0-.5-.22-.5-.5V9.91c0-.77.43-1.45 1.11-1.8c.25-.12.55-.02.67.22c.12.25.02.55-.22.67c-.34.17-.56.52-.56.9v3.46c0 .28-.22.5-.5.5Z"/><path fill="currentColor" d="M11.32 9.37c-2.58 0-4.68-2.1-4.68-4.68S8.74 0 11.32 0S16 2.1 16 4.68s-2.1 4.68-4.68 4.68Zm0-8.37C9.29 1 7.64 2.65 7.64 4.68s1.65 3.68 3.68 3.68S15 6.71 15 4.68S13.35 1 11.32 1"/><path fill="currentColor" d="M11.25 7.43c-.97 0-1.75-.79-1.75-1.75v-2c0-.97.79-1.75 1.75-1.75S13 2.72 13 3.68c0 .28-.22.5-.5.5s-.5-.22-.5-.5c0-.41-.34-.75-.75-.75s-.75.34-.75.75v2c0 .41.34.75.75.75s.75-.34.75-.75c0-.28.22-.5.5-.5s.5.22.5.5c0 .97-.79 1.75-1.75 1.75"/><path fill="currentColor" d="M11 4.46H9.5c-.28 0-.5-.22-.5-.5s.22-.5.5-.5H11c.28 0 .5.22.5.5s-.22.5-.5.5M11 6H9.5c-.28 0-.5-.22-.5-.5s.22-.5.5-.5H11c.28 0 .5.22.5.5s-.22.5-.5.5m-5.6 7.4C2.42 13.4 0 10.98 0 8a5.403 5.403 0 0 1 7.47-4.99c.26.11.38.4.27.65c-.11.26-.4.38-.65.27A4.406 4.406 0 0 0 1.01 8c0 2.42 1.98 4.4 4.4 4.4c.28 0 .5.22.5.5s-.22.5-.5.5Z"/><path fill="currentColor" d="M5.47 9.5c-.28 0-.5-.22-.5-.5V4.85c0-.28.22-.5.5-.5s.5.22.5.5V9c0 .28-.22.5-.5.5"/><path fill="currentColor" d="M5.3 8.5a1.73 1.73 0 0 1 0-3.46h.35c.95 0 1.73.77 1.73 1.73v.18c0 .28-.22.5-.5.5s-.5-.22-.5-.5v-.18c0-.4-.33-.73-.73-.73H5.3c-.4 0-.73.33-.73.73s.33.73.73.73c.28 0 .5.22.5.5s-.22.5-.5.5M5 10.92s-.08 0-.12-.02c-.77-.19-1.31-.88-1.31-1.67v-.18c0-.28.22-.5.5-.5s.5.22.5.5v.18c0 .33.23.62.55.7c.27.07.43.34.36.61c-.06.23-.26.38-.48.38"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Number(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 16h-12C.67 16 0 15.33 0 14.5v-13C0 .67.67 0 1.5 0h12c.83 0 1.5.67 1.5 1.5v13c0 .83-.67 1.5-1.5 1.5M1.5 1c-.28 0-.5.22-.5.5v13c0 .28.22.5.5.5h12c.28 0 .5-.22.5-.5v-13c0-.28-.22-.5-.5-.5z"/><path fill="currentColor" d="M7.41 12C9.07 12 10 10.8 10 9.46S8.93 7.1 7.61 7.06v-.07L9.7 5V4H5.25v1H8.5v.12L6.5 7v.88l.89-.03c.96 0 1.69.65 1.69 1.64s-.68 1.66-1.67 1.66c-.91 0-1.54-.63-1.57-1.45H5c0 1.02.81 2.3 2.41 2.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Open(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 3h-9c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h9c.28 0 .5.22.5.5s-.22.5-.5.5m0 3h-9c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h9c.28 0 .5.22.5.5s-.22.5-.5.5m0 3h-9c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h9c.28 0 .5.22.5.5s-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Password(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.35 16H2.65C1.74 16 1 15.26 1 14.35v-7.7C1 5.74 1.74 5 2.65 5h10.7c.91 0 1.65.74 1.65 1.65v7.69c0 .91-.74 1.65-1.65 1.65ZM2.65 6c-.36 0-.65.29-.65.65v7.69c0 .36.29.65.65.65h10.7c.36 0 .65-.29.65-.65V6.65c0-.36-.29-.65-.65-.65z"/><path fill="currentColor" d="M12.54 6H3.46V4.54C3.46 2.04 5.5 0 8 0s4.54 2.04 4.54 4.54zM4.46 5h7.08v-.46C11.54 2.59 9.95 1 8 1S4.46 2.59 4.46 4.54z"/><circle cx="12" cy="10.5" r="1" fill="currentColor"/><circle cx="8" cy="10.5" r="1" fill="currentColor"/><circle cx="4" cy="10.5" r="1" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.62 13c-.28 0-.5-.22-.5-.5v-9c0-.28.22-.5.5-.5s.5.22.5.5v9c0 .28-.22.5-.5.5m6 0c-.28 0-.5-.22-.5-.5v-9c0-.28.22-.5.5-.5s.5.22.5.5v9c0 .28-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paypal(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 13h-13C.67 13 0 12.33 0 11.5v-9C0 1.67.67 1 1.5 1h13c.83 0 1.5.67 1.5 1.5v9c0 .83-.67 1.5-1.5 1.5M1.5 2c-.28 0-.5.22-.5.5v9c0 .28.22.5.5.5h13c.28 0 .5-.22.5-.5v-9c0-.28-.22-.5-.5-.5z"/><path fill="currentColor" d="M10.06 5.51s0 .06-.03.1c-.26 1.28-1.09 1.7-2.18 1.7H7.3c-.13 0-.26.1-.26.22l-.29 1.8l-.06.51c0 .1.06.16.13.16h.99c.13 0 .22-.1.22-.19v-.06l.19-1.16v-.06c.03-.13.13-.19.22-.19h.16c.96 0 1.7-.39 1.89-1.51c.1-.48.03-.87-.19-1.12a1.46 1.46 0 0 0-.26-.19Z"/><path fill="currentColor" d="M9.81 5.41s-.06-.03-.13-.03c-.03 0-.1-.03-.13-.03c-.16-.03-.32-.03-.48-.03H7.59s-.06 0-.1.03c-.06.03-.13.1-.13.16L7.04 7.5v.06c.03-.13.13-.22.26-.22h.55c1.09 0 1.93-.45 2.18-1.7c0-.03 0-.06.03-.1l-.19-.1z"/><path fill="currentColor" d="M7.37 5.51c0-.06.06-.13.13-.16c.03 0 .06-.03.1-.03h1.48c.16 0 .35 0 .48.03c.03 0 .1 0 .13.03c.03 0 .06.03.13.03l.06.03l.19.1c.06-.48 0-.8-.26-1.09c-.29-.32-.8-.45-1.44-.45H6.48c-.13 0-.26.1-.26.22l-.77 4.94c0 .1.06.19.16.19h1.16l.29-1.86l.32-1.99Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func People(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 7a2.5 2.5 0 0 1 0-5a2.5 2.5 0 0 1 0 5m0-4C6.67 3 6 3.67 6 4.5S6.67 6 7.5 6S9 5.33 9 4.5S8.33 3 7.5 3"/><path fill="currentColor" d="M13.5 11c-.28 0-.5-.22-.5-.5s.22-.5.5-.5s.5-.22.5-.5A2.5 2.5 0 0 0 11.5 7h-1c-.28 0-.5-.22-.5-.5s.22-.5.5-.5c.83 0 1.5-.67 1.5-1.5S11.33 3 10.5 3c-.28 0-.5-.22-.5-.5s.22-.5.5-.5A2.5 2.5 0 0 1 13 4.5c0 .62-.22 1.18-.6 1.62c1.49.4 2.6 1.76 2.6 3.38c0 .83-.67 1.5-1.5 1.5m-12 0C.67 11 0 10.33 0 9.5c0-1.62 1.1-2.98 2.6-3.38c-.37-.44-.6-1-.6-1.62A2.5 2.5 0 0 1 4.5 2c.28 0 .5.22.5.5s-.22.5-.5.5C3.67 3 3 3.67 3 4.5S3.67 6 4.5 6c.28 0 .5.22.5.5s-.22.5-.5.5h-1A2.5 2.5 0 0 0 1 9.5c0 .28.22.5.5.5s.5.22.5.5s-.22.5-.5.5m9 3h-6c-.83 0-1.5-.67-1.5-1.5v-1C3 9.57 4.57 8 6.5 8h2c1.93 0 3.5 1.57 3.5 3.5v1c0 .83-.67 1.5-1.5 1.5m-4-5A2.5 2.5 0 0 0 4 11.5v1c0 .28.22.5.5.5h6c.28 0 .5-.22.5-.5v-1A2.5 2.5 0 0 0 8.5 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Peso(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.5 7h-6c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h6c.28 0 .5.22.5.5s-.22.5-.5.5m0 2h-6c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h6c.28 0 .5.22.5.5s-.22.5-.5.5"/><path fill="currentColor" d="M8 15c-3.86 0-7-3.14-7-7s3.14-7 7-7s7 3.14 7 7s-3.14 7-7 7M8 2C4.69 2 2 4.69 2 8s2.69 6 6 6s6-2.69 6-6s-2.69-6-6-6"/><path fill="currentColor" d="M6.5 13c-.28 0-.5-.22-.5-.5v-7C6 4.67 6.67 4 7.5 4h1A2.5 2.5 0 0 1 11 6.5v2A2.5 2.5 0 0 1 8.5 11H7v1.5c0 .28-.22.5-.5.5m.5-3h1.5c.83 0 1.5-.67 1.5-1.5v-2C10 5.67 9.33 5 8.5 5h-1c-.28 0-.5.22-.5.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pinterest(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.34 1C5.46 1 2.62 2.92 2.62 6.02c0 1.97 1.11 3.1 1.78 3.1c.28 0 .44-.77.44-.99c0-.26-.66-.82-.66-1.9c0-2.26 1.72-3.85 3.94-3.85c1.91 0 3.32 1.09 3.32 3.08c0 1.49-.6 4.28-2.53 4.28c-.7 0-1.3-.5-1.3-1.23c0-1.06.74-2.09.74-3.18c0-1.86-2.63-1.52-2.63.72c0 .47.06.99.27 1.42c-.39 1.67-1.18 4.15-1.18 5.87c0 .53.08 1.05.13 1.58c.1.11.05.1.19.04c1.41-1.94 1.36-2.31 2-4.85c.35.66 1.24 1.01 1.94 1.01c2.98 0 4.32-2.9 4.32-5.52c0-2.79-2.41-4.6-5.05-4.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 13.62c-.24 0-.47-.06-.69-.17c-.5-.26-.81-.77-.81-1.33V3.98c0-.56.31-1.07.81-1.33s1.1-.22 1.56.11l5.77 4.07c.4.28.63.74.63 1.23s-.24.94-.63 1.23l-5.77 4.06c-.26.18-.56.28-.86.28Zm0-10.14c-.08 0-.16.02-.23.06c-.17.09-.27.25-.27.44v8.14c0 .19.1.36.27.44c.17.09.36.07.52-.04l5.77-4.07a.507.507 0 0 0 0-.82L2.79 3.57a.521.521 0 0 0-.29-.09"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Playcircle(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.89 7.79l-3-2A.25.25 0 0 0 6.5 6v4c0 .2.22.32.39.21l3-2c.15-.1.15-.31 0-.41Z"/><path fill="currentColor" d="M8 15c-3.86 0-7-3.14-7-7s3.14-7 7-7s7 3.14 7 7s-3.14 7-7 7M8 2C4.69 2 2 4.69 2 8s2.69 6 6 6s6-2.69 6-6s-2.69-6-6-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pound(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 15c-3.86 0-7-3.14-7-7s3.14-7 7-7s7 3.14 7 7s-3.14 7-7 7M8 2C4.69 2 2 4.69 2 8s2.69 6 6 6s6-2.69 6-6s-2.69-6-6-6"/><path fill="currentColor" d="M10.5 12h-5c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h5c.28 0 .5.22.5.5s-.22.5-.5.5m-2-3h-3c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h3c.28 0 .5.22.5.5s-.22.5-.5.5"/><path fill="currentColor" d="M6.5 12c-.28 0-.5-.22-.5-.5V6.62C6 5.17 7.18 4 8.62 4c.7 0 1.36.27 1.85.77l.38.38c.2.2.2.51 0 .71c-.2.2-.51.2-.71 0l-.38-.38c-.31-.31-.71-.47-1.15-.47c-.89 0-1.62.73-1.62 1.62v4.88c0 .28-.22.5-.5.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Radio(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 15c-3.86 0-7-3.14-7-7s3.14-7 7-7s7 3.14 7 7s-3.14 7-7 7M8 2C4.69 2 2 4.69 2 8s2.69 6 6 6s6-2.69 6-6s-2.69-6-6-6"/><circle cx="8" cy="8" r="4" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RangeIcon(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.31 2.5H.5C.22 2.5 0 2.28 0 2s.22-.5.5-.5h1.81c-.04.16-.07.33-.07.5s.03.34.07.5m11.19-1H6.17c.04.16.07.33.07.5s-.03.34-.07.5h7.33c.28 0 .5-.22.5-.5s-.22-.5-.5-.5m-7.33 1c-.22.86-1 1.5-1.93 1.5s-1.71-.64-1.93-1.5c-.04-.16-.07-.33-.07-.5s.03-.34.07-.5C2.53.64 3.31 0 4.24 0s1.71.64 1.93 1.5c.04.16.07.33.07.5s-.03.34-.07.5M5.24 2c0-.55-.45-1-1-1s-1 .45-1 1s.45 1 1 1s1-.45 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reddit(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 8c0 3.87-3.13 7-7 7s-7-3.13-7-7s3.13-7 7-7s7 3.13 7 7m-2.52-.55c.11.16.18.35.18.55c0 .19-.05.38-.15.55c-.1.16-.25.3-.42.38v.31c0 1.57-1.83 2.84-4.08 2.84s-4.08-1.27-4.08-2.84v-.31c-.13-.06-.25-.15-.35-.27a.907.907 0 0 1-.2-.39a.932.932 0 0 1-.02-.44c.02-.15.08-.28.16-.41a.92.92 0 0 1 .32-.3c.13-.08.27-.13.41-.14c.15-.02.29 0 .43.04s.27.12.38.22c.81-.55 1.76-.85 2.73-.86l.52-2.43s.02-.05.03-.08l.06-.06c.02-.02.05-.03.08-.03h.09l1.72.34a.75.75 0 0 1 .37-.31c.16-.06.33-.05.48 0s.28.17.36.32a.696.696 0 0 1-.65 1.03c-.17 0-.32-.08-.45-.19a.717.717 0 0 1-.22-.43l-1.5-.31l-.46 2.18c.96.02 1.9.32 2.69.86c.11-.1.23-.18.37-.23c.14-.05.29-.07.43-.05c.15.01.29.06.42.13s.24.18.32.3Zm-6.7.86a.6.6 0 0 0-.1.25c-.02.09-.02.18 0 .27a.696.696 0 0 0 .83.55c.09-.02.18-.05.25-.1a.687.687 0 0 0 .29-.72a.696.696 0 0 0-.83-.55c-.18.04-.34.14-.45.3ZM8 11.16c.62.03 1.23-.14 1.73-.51c.04-.04.06-.09.06-.14c0-.05-.02-.1-.06-.14c-.02-.02-.04-.03-.06-.04a.185.185 0 0 0-.07-.02c-.05 0-.1.02-.14.06c-.42.3-.94.46-1.46.43c-.52.02-1.04-.13-1.46-.44a.201.201 0 0 0-.13-.04c-.05 0-.09.02-.12.06c-.03.03-.05.08-.06.12c0 .05.01.09.04.13c.5.37 1.11.56 1.73.54Zm1.21-1.85c.12.08.24.15.38.15c.09 0 .19-.02.28-.05c.09-.04.17-.09.23-.16c.07-.07.12-.15.15-.24s.05-.18.05-.28c0-.12-.03-.24-.09-.34a.777.777 0 0 0-.25-.26c-.1-.06-.22-.1-.34-.1c-.12 0-.24.03-.35.08a.67.67 0 0 0-.26.24c-.06.1-.1.22-.11.34c0 .12.02.24.07.35c.05.11.14.2.24.27"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Refresh(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 6.68H8.82c-.28 0-.5-.22-.5-.5s.22-.5.5-.5H12V2.49c0-.28.22-.5.5-.5s.5.22.5.5v3.68c0 .28-.22.5-.5.5Z"/><path fill="currentColor" d="M7 14.01c-3.31 0-6-2.69-6-6s2.69-6 6-6c2.34 0 4.48 1.37 5.46 3.5c.12.25 0 .55-.25.66c-.25.12-.55 0-.66-.25A5.022 5.022 0 0 0 7 3C4.24 3 2 5.24 2 8s2.24 5 5 5s5-2.24 5-5c0-.28.22-.5.5-.5s.5.22.5.5c0 3.31-2.69 6-6 6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reorder(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="3.5" cy="3.5" r="1.5" fill="currentColor"/><circle cx="3.5" cy="8.5" r="1.5" fill="currentColor"/><circle cx="3.5" cy="13.5" r="1.5" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Repeater(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 5h-11C1.67 5 1 4.33 1 3.5v-1C1 1.67 1.67 1 2.5 1h11c.83 0 1.5.67 1.5 1.5v1c0 .83-.67 1.5-1.5 1.5m-11-3c-.28 0-.5.22-.5.5v1c0 .28.22.5.5.5h11c.28 0 .5-.22.5-.5v-1c0-.28-.22-.5-.5-.5zm11 8h-11C1.67 10 1 9.33 1 8.5v-1C1 6.67 1.67 6 2.5 6h11c.83 0 1.5.67 1.5 1.5v1c0 .83-.67 1.5-1.5 1.5m-11-3c-.28 0-.5.22-.5.5v1c0 .28.22.5.5.5h11c.28 0 .5-.22.5-.5v-1c0-.28-.22-.5-.5-.5zm11 8h-11c-.83 0-1.5-.67-1.5-1.5v-1c0-.83.67-1.5 1.5-1.5h11c.83 0 1.5.67 1.5 1.5v1c0 .83-.67 1.5-1.5 1.5m-11-3c-.28 0-.5.22-.5.5v1c0 .28.22.5.5.5h11c.28 0 .5-.22.5-.5v-1c0-.28-.22-.5-.5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reply(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 12.5a.47.47 0 0 1-.35-.15l-4.5-4.5C1.06 7.76 1 7.63 1 7.5s.05-.26.15-.35l4.5-4.5c.2-.2.51-.2.71 0c.2.2.2.51 0 .71L2.21 7.5l4.15 4.15c.2.2.2.51 0 .71c-.1.1-.23.15-.35.15Z"/><path fill="currentColor" d="M13.5 14c-.28 0-.5-.22-.5-.5v-3A2.5 2.5 0 0 0 10.5 8H2.7c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h7.8c1.93 0 3.5 1.57 3.5 3.5v3c0 .28-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rewind(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.91 11.01c-.2 0-.41-.06-.59-.17L9.29 8.92a.5.5 0 0 1-.15-.69a.5.5 0 0 1 .69-.15L12.86 10s.07.02.1 0s.05-.05.05-.09V2.09s-.02-.07-.05-.09a.09.09 0 0 0-.1 0L9.83 3.92a.5.5 0 0 1-.69-.15a.5.5 0 0 1 .15-.69l3.03-1.92c.34-.21.77-.23 1.12-.03c.35.19.57.56.57.96v7.82c0 .4-.22.77-.57.96c-.17.09-.35.14-.53.14"/><path fill="currentColor" d="M8.96 10.98c-.21 0-.42-.06-.61-.18L2.49 6.92C2.18 6.72 2 6.37 2 6s.18-.71.49-.92L8.36 1.2c.34-.22.77-.24 1.13-.05c.36.19.58.56.58.97v7.77c0 .41-.22.78-.58.97c-.16.09-.34.13-.52.13ZM2.76 5.5l.28.42S3 5.98 3 6l5.92 3.97s.07.02.1 0s.05-.05.05-.09V2.11s-.02-.07-.05-.09a.09.09 0 0 0-.1 0l-5.88 3.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Right(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 13a.47.47 0 0 1-.35-.15c-.2-.2-.2-.51 0-.71L5.3 7.99L1.15 3.85c-.2-.2-.2-.51 0-.71c.2-.2.51-.2.71 0l4.49 4.51c.2.2.2.51 0 .71l-4.5 4.49c-.1.1-.23.15-.35.15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ruble(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 15c-3.86 0-7-3.14-7-7s3.14-7 7-7s7 3.14 7 7s-3.14 7-7 7M8 2C4.69 2 2 4.69 2 8s2.69 6 6 6s6-2.69 6-6s-2.69-6-6-6"/><path fill="currentColor" d="M6.5 13c-.28 0-.5-.22-.5-.5V9H4.5c-.28 0-.5-.22-.5-.5s.22-.5.5-.5H6V4.5c0-.28.22-.5.5-.5h3a2.5 2.5 0 0 1 0 5H7v3.5c0 .28-.22.5-.5.5M7 8h2.5c.83 0 1.5-.67 1.5-1.5S10.33 5 9.5 5H7z"/><path fill="currentColor" d="M9.5 11h-5c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h5c.28 0 .5.22.5.5s-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rupee(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 15c-3.86 0-7-3.14-7-7s3.14-7 7-7s7 3.14 7 7s-3.14 7-7 7M8 2C4.69 2 2 4.69 2 8s2.69 6 6 6s6-2.69 6-6s-2.69-6-6-6"/><path fill="currentColor" d="M6.5 9h-1c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h1C7.33 8 8 7.33 8 6.5S7.33 5 6.5 5h-1c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h1a2.5 2.5 0 0 1 0 5"/><path fill="currentColor" d="M10.5 7h-5c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h5c.28 0 .5.22.5.5s-.22.5-.5.5m0-2h-5c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h5c.28 0 .5.22.5.5s-.22.5-.5.5m-3 7c-.16 0-.32-.08-.42-.22l-2-3a.498.498 0 0 1 .83-.55l2 3c.15.23.09.54-.14.69c-.09.06-.18.08-.28.08Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sad(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="6" cy="6" r="1" fill="currentColor"/><circle cx="10" cy="6" r="1" fill="currentColor"/><path fill="currentColor" d="M8 15c-3.86 0-7-3.14-7-7s3.14-7 7-7s7 3.14 7 7s-3.14 7-7 7M8 2C4.69 2 2 4.69 2 8s2.69 6 6 6s6-2.69 6-6s-2.69-6-6-6"/><path fill="currentColor" d="M10.83 11.5c-.21 0-.4-.13-.47-.33a2.502 2.502 0 0 0-4.72 0c-.09.26-.38.4-.64.3a.493.493 0 0 1-.3-.64C5.19 9.43 6.52 8.5 8 8.5s2.81.94 3.3 2.33a.501.501 0 0 1-.47.67"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.5 13.02a5.49 5.49 0 0 1-3.89-1.61C1.57 10.37 1 8.99 1 7.52s.57-2.85 1.61-3.89c2.14-2.14 5.63-2.14 7.78 0C11.43 4.67 12 6.05 12 7.52s-.57 2.85-1.61 3.89a5.49 5.49 0 0 1-3.89 1.61m0-10c-1.15 0-2.3.44-3.18 1.32C2.47 5.19 2 6.32 2 7.52s.47 2.33 1.32 3.18a4.509 4.509 0 0 0 6.36 0C10.53 9.85 11 8.72 11 7.52s-.47-2.33-1.32-3.18A4.483 4.483 0 0 0 6.5 3.02"/><path fill="currentColor" d="M13.5 15a.47.47 0 0 1-.35-.15l-3.38-3.38c-.2-.2-.2-.51 0-.71c.2-.2.51-.2.71 0l3.38 3.38c.2.2.2.51 0 .71c-.1.1-.23.15-.35.15Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Select(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M40 3H4a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h36a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1M4 0a4 4 0 0 0-4 4v16a4 4 0 0 0 4 4h36a4 4 0 0 0 4-4V4a4 4 0 0 0-4-4zm4 11.5c0-.828.38-1.5.848-1.5h11.304c.468 0 .848.672.848 1.5s-.38 1.5-.848 1.5H8.848C8.38 13 8 12.328 8 11.5m23.707-1.207L31 9.586L29.586 11l.707.707l2 2l.707.707l.707-.707l2-2l.707-.707L35 9.586l-.707.707L33 11.586z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Settings(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 10.5a2.5 2.5 0 0 1 0-5a2.5 2.5 0 0 1 0 5m0-4c-.83 0-1.5.67-1.5 1.5S7.17 9.5 8 9.5S9.5 8.83 9.5 8S8.83 6.5 8 6.5"/><path fill="currentColor" d="M8.85 15h-1.7c-.32 0-.6-.22-.67-.53l-.41-1.79l-1.56.98c-.27.17-.62.13-.85-.1l-1.2-1.2a.683.683 0 0 1-.1-.85l.98-1.56l-1.79-.41a.686.686 0 0 1-.53-.67v-1.7c0-.32.22-.6.53-.67l1.79-.41l-.98-1.56a.683.683 0 0 1 .1-.85l1.2-1.2c.23-.23.58-.27.85-.1l1.56.98l.41-1.79c.07-.31.35-.53.67-.53h1.7c.32 0 .6.22.67.53l.41 1.79l1.56-.98c.27-.17.62-.13.85.1l1.2 1.2c.23.23.27.58.1.85l-.98 1.56l1.79.41c.31.07.53.35.53.67v1.7c0 .32-.22.6-.53.67l-1.79.41l.98 1.56c.17.27.13.62-.1.85l-1.2 1.2c-.23.23-.58.27-.85.1l-1.56-.98l-.41 1.79c-.07.31-.35.53-.67.53M7.4 14h1.2l.67-2.92l2.54 1.59l.85-.85l-1.59-2.54l2.92-.67v-1.2l-2.92-.67l1.59-2.54l-.85-.85l-2.54 1.59l-.67-2.92H7.4l-.67 2.92l-2.54-1.59l-.85.85l1.59 2.54l-2.92.67v1.2l2.92.67l-1.59 2.54l.85.85l2.54-1.59zm6.84-6.55"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 10c-.28 0-.5-.22-.5-.5v-8c0-.28.22-.5.5-.5s.5.22.5.5v8c0 .28-.22.5-.5.5"/><path fill="currentColor" d="M9.5 4a.47.47 0 0 1-.35-.15L7.5 2.2L5.85 3.85c-.2.2-.51.2-.71 0c-.2-.2-.2-.51 0-.71l2.01-1.99c.2-.2.51-.2.71 0l2 2c.2.2.2.51 0 .71c-.1.1-.23.15-.35.15Zm3 11h-10c-.83 0-1.5-.67-1.5-1.5v-7C1 5.67 1.67 5 2.5 5h2c.28 0 .5.22.5.5s-.22.5-.5.5h-2c-.28 0-.5.22-.5.5v7c0 .28.22.5.5.5h10c.28 0 .5-.22.5-.5v-7c0-.28-.22-.5-.5-.5h-2c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h2c.83 0 1.5.67 1.5 1.5v7c0 .83-.67 1.5-1.5 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shekel(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 15.5C3.86 15.5.5 12.14.5 8S3.86.5 8 .5s7.5 3.36 7.5 7.5s-3.36 7.5-7.5 7.5m0-14C4.42 1.5 1.5 4.42 1.5 8s2.92 6.5 6.5 6.5s6.5-2.92 6.5-6.5S11.58 1.5 8 1.5"/><path fill="currentColor" d="M5 12c-.28 0-.5-.22-.5-.5v-7c0-.28.22-.5.5-.5h2a2.5 2.5 0 0 1 2.5 2.5v3c0 .28-.22.5-.5.5s-.5-.22-.5-.5v-3C8.5 5.67 7.83 5 7 5H5.5v6.5c0 .28-.22.5-.5.5"/><path fill="currentColor" d="M11 12H9a2.5 2.5 0 0 1-2.5-2.5v-3c0-.28.22-.5.5-.5s.5.22.5.5v3c0 .83.67 1.5 1.5 1.5h1.5V4.5c0-.28.22-.5.5-.5s.5.22.5.5v7c0 .28-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Skype(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.45 9.02a6.525 6.525 0 0 0-1.83-5.63A6.566 6.566 0 0 0 8 1.48c-.34 0-.68.03-1.02.08c-.76-.45-1.64-.64-2.51-.53c-.87.11-1.69.5-2.31 1.12a3.96 3.96 0 0 0-1.13 2.3c-.11.87.08 1.75.53 2.51c-.06.34-.09.69-.1 1.03c0 1.73.69 3.39 1.91 4.61a6.566 6.566 0 0 0 4.62 1.91c.34 0 .68-.03 1.02-.08c.76.45 1.64.64 2.51.53c.87-.11 1.69-.5 2.31-1.12a3.96 3.96 0 0 0 1.13-2.3c.11-.87-.08-1.75-.53-2.51l.02-.02Zm-3.35 1.95c-.31.42-.73.75-1.21.94c-.6.24-1.24.35-1.89.34c-.75.03-1.5-.12-2.17-.45c-.39-.21-.73-.51-.99-.88c-.23-.31-.36-.69-.38-1.08c0-.1.02-.2.06-.29c.04-.09.1-.18.17-.25c.16-.14.38-.21.59-.2c.18 0 .35.06.49.17c.15.14.27.32.34.51c.09.23.21.44.36.64c.14.18.33.33.54.42c.29.12.6.18.92.17c.45.02.89-.1 1.27-.34c.15-.08.27-.21.35-.35a.974.974 0 0 0 .07-.82a.722.722 0 0 0-.2-.29c-.19-.18-.43-.31-.68-.38c-.28-.09-.66-.18-1.13-.29c-.54-.13-1.07-.3-1.58-.53c-.4-.16-.75-.42-1.02-.76a1.9 1.9 0 0 1-.38-1.19c0-.44.14-.86.4-1.21c.3-.37.7-.66 1.16-.81a5.18 5.18 0 0 1 1.79-.26c.47 0 .95.06 1.4.19c.36.11.69.27.99.5c.23.18.43.4.57.65c.12.21.18.44.18.68c0 .1-.02.2-.06.29s-.1.18-.17.25c-.07.08-.16.14-.26.19c-.1.04-.21.06-.32.06a.69.69 0 0 1-.48-.15c-.14-.14-.26-.3-.35-.47c-.13-.28-.33-.52-.57-.71c-.32-.2-.69-.28-1.06-.25c-.38-.02-.76.08-1.08.27c-.12.06-.22.16-.29.27c-.07.11-.11.24-.12.38c0 .15.04.29.14.41c.11.12.24.23.39.3c.16.08.33.15.5.19c.17.05.46.12.85.21c.5.11.94.23 1.35.36c.36.11.71.27 1.03.48c.28.18.51.42.67.71c.16.33.24.69.22 1.06c0 .48-.14.95-.42 1.34Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Snapchat(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.4 3.7c-.35.77-.21 2.16-.16 3.13c-.38.21-.86-.16-1.14-.16s-.63.19-.68.46c-.04.2.05.49.7.74c.25.1.85.21.99.53c.19.45-1 2.53-2.87 2.84c-.07.01-.13.05-.18.1c-.04.06-.07.12-.06.2c.03.56 1.31.78 1.87.87c.06.08.1.4.18.65c.03.11.12.24.34.24c.29 0 .77-.22 1.6-.08c.82.13 1.58 1.27 3.05 1.27c1.37 0 2.18-1.15 2.97-1.27c.45-.07.84-.05 1.28.03c.3.06.57.09.66-.2c.08-.25.12-.57.18-.65c.56-.09 1.84-.3 1.87-.87a.36.36 0 0 0-.06-.2a.253.253 0 0 0-.18-.1c-1.84-.3-3.07-2.37-2.87-2.84c.14-.32.73-.43.99-.53c.47-.18.71-.41.71-.67c0-.34-.42-.54-.72-.54s-.75.36-1.11.16c.06-.98.19-2.36-.16-3.13c-.66-1.46-2.14-2.2-3.61-2.2s-2.92.73-3.59 2.2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Solana(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.45 6.76h9.59c.12 0 .23.05.32.14l1.52 1.56c.28.29.08.78-.32.78H3.97c-.12 0-.23-.05-.32-.14L2.13 7.54c-.28-.29-.08-.78.32-.78m-.32-2.07l1.52-1.56c.08-.09.2-.14.32-.14h9.58c.4 0 .6.49.32.78l-1.51 1.56c-.08.09-.2.14-.32.14H2.45c-.4 0-.6-.49-.32-.78m11.74 6.61l-1.52 1.56c-.09.09-.2.14-.32.14H2.45c-.4 0-.6-.49-.32-.78l1.52-1.56c.08-.09.2-.14.32-.14h9.58c.4 0 .6.49.32.78"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spinner(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.56 13.88c-.28 0-.5-.22-.5-.5s.22-.5.5-.5c2.96 0 5.38-2.41 5.38-5.38s-2.41-5.38-5.38-5.38c-.28 0-.5-.22-.5-.5s.22-.5.5-.5c3.52 0 6.38 2.86 6.38 6.38s-2.86 6.38-6.38 6.38"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.41 8.41l1.14-.93l1.14-.93c.48-.39.37-.74-.25-.77l-1.58-.09l-2.5-.14l-.41-1.05l-.53-1.38l-.53-1.38c-.22-.58-.59-.58-.81 0L6.01 4.49L5.6 5.54l-2.5.14l-1.58.09c-.62.03-.73.38-.25.77l1.14.93l1.14.93l.87.71l-.57 2.15l-.47 1.79c-.16.6.14.81.66.48l2.48-1.6l.94-.61l.94.61l1.24.8l1.24.8c.52.33.82.12.66-.48l-.47-1.79l-.57-2.15l.87-.71Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Start(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 10c-.28 0-.5-.22-.5-.5v-7c0-.28.22-.5.5-.5s.5.22.5.5v7c0 .28-.22.5-.5.5"/><path fill="currentColor" d="M7.5 15C3.92 15 1 12.18 1 8.72c0-2.41 1.46-4.64 3.72-5.68c.25-.11.55 0 .66.25c.11.25 0 .55-.25.66c-1.91.87-3.14 2.74-3.14 4.77c0 2.91 2.47 5.28 5.5 5.28s5.5-2.37 5.5-5.28c0-2.02-1.23-3.9-3.14-4.77a.502.502 0 0 1-.25-.66c.11-.25.41-.36.66-.25c2.26 1.03 3.72 3.26 3.72 5.68c0 3.46-2.92 6.28-6.5 6.28Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stepback(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 12.7c-.27 0-.55-.08-.79-.23L4.6 9.27c-.44-.28-.71-.75-.71-1.27s.26-1 .71-1.27l.26.42l-.26-.42l5.11-3.2c.47-.29 1.04-.31 1.52-.04s.77.76.77 1.31v6.39c0 .55-.29 1.04-.77 1.31c-.23.13-.48.19-.73.19Zm0-8.4c-.08 0-.17.02-.26.08l-5.11 3.2c-.21.13-.24.34-.24.42s.02.29.24.42l5.11 3.2c.23.14.43.06.51.01c.08-.04.26-.17.26-.44V4.8a.496.496 0 0 0-.51-.5m-9 8.2c-.28 0-.5-.22-.5-.5V4c0-.28.22-.5.5-.5s.5.22.5.5v8c0 .28-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stepforward(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 12.7c-.25 0-.5-.06-.73-.19A1.49 1.49 0 0 1 1 11.2V4.8c0-.55.29-1.04.77-1.31c.48-.27 1.05-.25 1.52.04l5.11 3.2c.44.28.71.75.71 1.27s-.26 1-.71 1.27l-5.11 3.2c-.25.15-.52.23-.79.23m0-8.4c-.11 0-.2.04-.25.06c-.08.04-.26.17-.26.44v6.39c0 .27.18.39.26.44c.08.04.28.13.51-.01l5.11-3.2c.21-.13.24-.34.24-.42s-.02-.29-.24-.42l-5.11-3.2a.461.461 0 0 0-.26-.08m9 8.2c-.28 0-.5-.22-.5-.5V4c0-.28.22-.5.5-.5s.5.22.5.5v8c0 .28-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stop(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 14h-9c-.83 0-1.5-.67-1.5-1.5v-9C2 2.67 2.67 2 3.5 2h9c.83 0 1.5.67 1.5 1.5v9c0 .83-.67 1.5-1.5 1.5m-9-11c-.28 0-.5.22-.5.5v9c0 .28.22.5.5.5h9c.28 0 .5-.22.5-.5v-9c0-.28-.22-.5-.5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stripe(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 14h-13C.67 14 0 13.33 0 12.5v-9C0 2.67.67 2 1.5 2h13c.83 0 1.5.67 1.5 1.5v9c0 .83-.67 1.5-1.5 1.5M1.5 3c-.28 0-.5.22-.5.5v9c0 .28.22.5.5.5h13c.28 0 .5-.22.5-.5v-9c0-.28-.22-.5-.5-.5z"/><path fill="currentColor" fill-rule="evenodd" d="m8.5 6.18l-.84.18v-.68l.84-.18zm1.74.38c-.33 0-.54.15-.65.26l-.04-.21h-.73v3.88l.83-.18v-.94c.12.09.3.21.59.21c.6 0 1.14-.48 1.14-1.53c0-.96-.55-1.49-1.14-1.49m-.2 2.29c-.2 0-.31-.07-.39-.16V7.46c.08-.1.2-.16.39-.16c.3 0 .51.34.51.77s-.21.78-.51.78M14 8.08c0-.85-.41-1.52-1.2-1.52s-1.27.67-1.27 1.52c0 1 .57 1.51 1.38 1.51c.4 0 .7-.09.92-.22V8.7c-.23.11-.49.18-.82.18s-.61-.11-.65-.51h1.63v-.3Zm-1.65-.32c0-.38.23-.53.44-.53s.42.16.42.53zM7.66 6.61h.84v2.91h-.84zm-.95 0l.05.25c.2-.36.59-.29.69-.25v.76c-.1-.04-.44-.08-.63.17v1.97h-.83V6.6h.72ZM5.1 5.89l-.81.17v2.66c0 .49.37.85.86.85c.27 0 .47-.05.58-.11v-.68c-.11.04-.63.2-.63-.3V7.3h.63v-.71H5.1v-.72ZM3.13 7.28c-.18 0-.28.05-.28.18c0 .14.18.2.41.28c.37.13.86.29.86.91c0 .6-.48.94-1.17.94c-.29 0-.6-.06-.91-.19v-.79c.28.15.63.27.91.27c.19 0 .32-.05.32-.2s-.2-.23-.44-.32c-.37-.13-.83-.3-.83-.85c0-.59.45-.94 1.13-.94c.28 0 .55.04.83.15v.78c-.25-.14-.57-.21-.83-.21Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Submit(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.49 7.14L3.44 2.27c-.76-.41-1.64.3-1.4 1.13l1.24 4.34c.05.18.05.36 0 .54l-1.24 4.34c-.24.83.64 1.54 1.4 1.13l9.05-4.87a.98.98 0 0 0 0-1.72Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Table(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 16h-12C.67 16 0 15.33 0 14.5v-13C0 .67.67 0 1.5 0h12c.83 0 1.5.67 1.5 1.5v13c0 .83-.67 1.5-1.5 1.5M1.5 1c-.28 0-.5.22-.5.5v13c0 .28.22.5.5.5h12c.28 0 .5-.22.5-.5v-13c0-.28-.22-.5-.5-.5z"/><path fill="currentColor" d="M7 .62h1v14.75H7z"/><path fill="currentColor" d="M.38 5H14.5v1H.38zM.5 9.99h14v1H.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.07 15.03c-.4 0-.77-.16-1.05-.44l-5.61-5.6c-.28-.28-.44-.66-.44-1.05s.16-.77.44-1.05l4.45-4.46c.28-.28.66-.44 1.05-.44h4.2c.77 0 1.5.3 2.04.85c.55.55.85 1.27.85 2.04v4.2c0 .4-.16.77-.44 1.05l-4.45 4.45c-.28.28-.66.44-1.05.44ZM6.92 2.99c-.13 0-.25.05-.35.14L2.12 7.59c-.09.09-.14.22-.14.35s.05.25.14.35l5.6 5.6c.19.19.5.19.69 0l4.45-4.45a.5.5 0 0 0 .14-.35v-4.2c0-.51-.2-.98-.55-1.34c-.36-.36-.83-.55-1.34-.55H6.92Z"/><circle cx="10" cy="6" r=".99" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Telephone(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.38 1.52c-.44.05-.79.41-.85.85l-.77 3.27c-.11.46.12.94.55 1.14l.59.27C9.46 8.2 8.49 8.99 8.49 8.99s-.8.97-1.94 1.41l-.27-.59a.998.998 0 0 0-1.14-.55l-3.27.77c-.44.06-.8.41-.85.85c-.1.82-.07 2.1.85 2.78c0 0 4.15 2.92 9.19-2.12c5.04-5.04 2.12-9.19 2.12-9.19c-.69-.92-1.97-.94-2.78-.85Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tether(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 2h10v2.38H3zm6.5 6.76v5.25h-3V8.76z"/><path fill="currentColor" d="M8 5.75c-3.66 0-6.62.73-6.62 1.62S4.35 8.99 8 8.99s6.62-.73 6.62-1.62S11.65 5.75 8 5.75m0 2.5c-3.38 0-6.12-.5-6.12-1.12S4.62 6.01 8 6.01s6.12.5 6.12 1.12S11.38 8.25 8 8.25"/><path fill="currentColor" d="M8 7.88c.52 0 1.02-.01 1.5-.04V3.75h-3v4.09c.48.02.98.04 1.5.04"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextIcon(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.5 4h-8c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h8c.28 0 .5.22.5.5s-.22.5-.5.5m4 3h-6c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h6c.28 0 .5.22.5.5s-.22.5-.5.5"/><path fill="currentColor" d="M5.5 14c-.28 0-.5-.22-.5-.5V3.55c0-.28.22-.5.5-.5s.5.22.5.5v9.95c0 .28-.22.5-.5.5m5 0c-.28 0-.5-.22-.5-.5V6.55c0-.28.22-.5.5-.5s.5.22.5.5v6.95c0 .28-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Textarea(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 2h-10v4.5h-1V2c0-.55.45-1 1-1h10c.55 0 1 .45 1 1v4.5h-1zm-10 12h10V9.5h1V14c0 .55-.45 1-1 1h-10c-.55 0-1-.45-1-1V9.5h1z"/><path fill="currentColor" d="M10.5 5h-6c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h6c.28 0 .5.22.5.5s-.22.5-.5.5"/><path fill="currentColor" d="M7.5 12c-.28 0-.5-.22-.5-.5v-7c0-.28.22-.5.5-.5s.5.22.5.5v7c0 .28-.22.5-.5.5M2 10c-1.1 0-2-.9-2-2s.9-2 2-2s2 .9 2 2s-.9 2-2 2m0-3c-.55 0-1 .45-1 1s.45 1 1 1s1-.45 1-1s-.45-1-1-1m11 3c-1.1 0-2-.9-2-2s.9-2 2-2s2 .9 2 2s-.9 2-2 2m0-3c-.55 0-1 .45-1 1s.45 1 1 1s1-.45 1-1s-.45-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tiktok(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.3 1.01c.75-.01 1.5 0 2.25-.01c.05.89.36 1.8 1.01 2.43c.64.65 1.55.94 2.44 1.04v2.35c-.83-.03-1.66-.2-2.42-.56c-.33-.15-.63-.34-.93-.54c0 1.7 0 3.41-.01 5.1c-.04.82-.31 1.63-.78 2.3c-.75 1.12-2.06 1.85-3.4 1.87c-.82.05-1.65-.18-2.35-.6c-1.16-.69-1.98-1.97-2.1-3.33c-.02-.29-.02-.58 0-.87c.1-1.11.65-2.17 1.49-2.89c.95-.84 2.29-1.24 3.54-1c.01.86-.02 1.73-.02 2.59c-.57-.19-1.24-.13-1.74.22c-.37.24-.64.6-.79 1.02c-.12.3-.09.62-.08.94c.14.96 1.05 1.76 2.01 1.67c.64 0 1.26-.39 1.59-.94c.11-.19.23-.39.24-.62c.06-1.04.03-2.08.04-3.13c0-2.35 0-4.7.01-7.04"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Time(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 15c-3.86 0-7-3.14-7-7s3.14-7 7-7s7 3.14 7 7s-3.14 7-7 7M8 2C4.69 2 2 4.69 2 8s2.69 6 6 6s6-2.69 6-6s-2.69-6-6-6"/><path fill="currentColor" d="M10 10.5c-.09 0-.18-.02-.26-.07l-2.5-1.5A.495.495 0 0 1 7 8.5v-4c0-.28.22-.5.5-.5s.5.22.5.5v3.72l2.26 1.35a.502.502 0 0 1-.26.93"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tools(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.09 15.07H2.5c-.83 0-1.5-.67-1.5-1.5v-2.59c0-.4.16-.78.44-1.06l2.21-2.21c.2-.2.51-.2.71 0s.2.51 0 .71l-2.21 2.21a.51.51 0 0 0-.15.35v2.59c0 .28.22.5.5.5h2.59c.13 0 .26-.05.35-.15l2.21-2.21c.2-.2.51-.2.71 0c.2.2.2.51 0 .71l-2.21 2.21c-.28.28-.66.44-1.06.44m7.03-6.62a.47.47 0 0 1-.35-.15c-.2-.2-.2-.51 0-.71l2.1-2.1c.19-.19.19-.51 0-.71l-2.59-2.59a.513.513 0 0 0-.71 0L8.41 4.35c-.2.2-.51.2-.71 0c-.2-.2-.2-.51 0-.71l2.16-2.16c.57-.57 1.55-.57 2.12 0l2.59 2.59c.58.58.58 1.54 0 2.12l-2.1 2.1c-.1.1-.23.15-.35.15Z"/><path fill="currentColor" d="M10.9 15.06c-.4 0-.78-.16-1.06-.44l-2.19-2.19c-.2-.2-.2-.51 0-.71c.2-.2.51-.2.71 0l2.19 2.19c.19.19.52.19.71 0l2.59-2.59c.19-.19.19-.51 0-.71l-2.19-2.19c-.2-.2-.2-.51 0-.71c.2-.2.51-.2.71 0l2.19 2.19c.58.58.58 1.54 0 2.12l-2.59 2.59c-.28.28-.66.44-1.06.44Z"/><path fill="currentColor" d="M8.25 12.82a.47.47 0 0 1-.35-.15L1.36 6.14c-.58-.58-.58-1.54 0-2.12l2.59-2.59c.57-.57 1.55-.57 2.12 0l6.54 6.54c.2.2.2.51 0 .71c-.2.2-.51.2-.71 0L5.36 2.14a.513.513 0 0 0-.71 0L2.06 4.73c-.19.19-.19.51 0 .71l6.54 6.54c.2.2.2.51 0 .71c-.1.1-.23.15-.35.15Z"/><path fill="currentColor" d="M5.56 6.01a.47.47 0 0 1-.35-.15c-.2-.2-.2-.51 0-.71l1.94-1.94c.2-.2.51-.2.71 0c.2.2.2.51 0 .71L5.92 5.86c-.1.1-.23.15-.35.15Zm2.03 2.03a.47.47 0 0 1-.35-.15c-.2-.2-.2-.51 0-.71l1.94-1.94c.2-.2.51-.2.71 0c.2.2.2.51 0 .71L7.95 7.89c-.1.1-.23.15-.35.15Zm2.03 2.03a.47.47 0 0 1-.35-.15c-.2-.2-.2-.51 0-.71l1.94-1.94c.2-.2.51-.2.71 0c.2.2.2.51 0 .71L9.98 9.92c-.1.1-.23.15-.35.15Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 4h-13c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h13c.28 0 .5.22.5.5s-.22.5-.5.5"/><path fill="currentColor" d="m11.02 3.81l-.44-1.46a.504.504 0 0 0-.48-.36H5.9a.5.5 0 0 0-.48.36l-.44 1.46l-.96-.29l.44-1.46C4.65 1.42 5.23.99 5.9.99h4.2c.67 0 1.24.43 1.44 1.07l.44 1.46z"/><path fill="currentColor" d="M11.53 15H4.47c-.81 0-1.47-.64-1.5-1.45l-.34-9.87l1-.03l.34 9.87c0 .27.23.48.5.48h7.07c.27 0 .49-.21.5-.48l.34-9.87l1 .03l-.34 9.87c-.03.81-.69 1.45-1.5 1.45Z"/><path fill="currentColor" d="M6.5 11.62c-.28 0-.5-.22-.5-.5v-4c0-.28.22-.5.5-.5s.5.22.5.5v4c0 .28-.22.5-.5.5m3 0c-.28 0-.5-.22-.5-.5v-4c0-.28.22-.5.5-.5s.5.22.5.5v4c0 .28-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitter(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 0c3.87 0 7 3.13 7 7s-3.13 7-7 7s-7-3.13-7-7s3.13-7 7-7M5.72 10.69c3.1 0 4.8-2.57 4.8-4.8v-.22c.33-.24.62-.54.84-.88c-.3.13-.63.22-.97.27c.35-.21.62-.54.74-.93c-.33.19-.69.33-1.07.41c-.31-.33-.75-.53-1.23-.53c-.93 0-1.69.76-1.69 1.69c0 .13.01.26.05.38c-1.4-.07-2.65-.74-3.48-1.76c-.14.25-.23.54-.23.85c0 .58.3 1.1.75 1.4c-.28 0-.54-.08-.76-.21v.02c0 .82.58 1.5 1.35 1.66c-.14.04-.29.06-.44.06c-.11 0-.21-.01-.32-.03c.21.67.84 1.16 1.57 1.17c-.58.45-1.31.72-2.1.72c-.14 0-.27 0-.4-.02c.74.48 1.63.76 2.58.76" class="cls-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Up(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 6a.47.47 0 0 1-.35-.15L8 1.71L3.85 5.85c-.2.2-.51.2-.71 0c-.2-.2-.2-.51 0-.71L7.65.65c.2-.2.51-.2.71 0l4.5 4.5c.2.2.2.51 0 .71c-.1.1-.23.15-.35.15Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Upload(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 12c-.28 0-.5-.22-.5-.5v-9c0-.28.22-.5.5-.5s.5.22.5.5v9c0 .28-.22.5-.5.5m6 3h-12c-.28 0-.5-.22-.5-.5v-2c0-.28.22-.5.5-.5s.5.22.5.5V14h11v-1.5c0-.28.22-.5.5-.5s.5.22.5.5v2c0 .28-.22.5-.5.5"/><path fill="currentColor" d="M11 5.5a.47.47 0 0 1-.35-.15L7.5 2.2L4.35 5.35c-.2.2-.51.2-.71 0s-.2-.51 0-.71l3.51-3.49c.2-.2.51-.2.71 0l3.5 3.5c.2.2.2.51 0 .71c-.1.1-.23.15-.35.15Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Uploadcloud(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 11c-.28 0-.5-.22-.5-.5s.22-.5.5-.5c1.65 0 3-1.35 3-3s-1.35-3-3-3h-1.05c-.18 0-.34-.09-.43-.25C9.88 2.65 8.76 2 7.51 2c-1.93 0-3.5 1.57-3.5 3.5c0 .28-.22.5-.5.5h-.5c-1.1 0-2 .9-2 2s.9 2 2 2c.28 0 .5.22.5.5s-.22.5-.5.5c-1.65 0-3-1.35-3-3s1.35-3 3-3h.03c.25-2.25 2.16-4 4.47-4c1.49 0 2.89.76 3.72 2h.78c2.21 0 4 1.79 4 4s-1.79 4-4 4Z"/><path fill="currentColor" d="M10 9.25a.47.47 0 0 1-.35-.15L7.5 6.95L5.35 9.1c-.2.2-.51.2-.71 0s-.2-.51 0-.71l2.5-2.5c.2-.2.51-.2.71 0l2.5 2.5c.2.2.2.51 0 .71c-.1.1-.23.15-.35.15"/><path fill="currentColor" d="M7.5 13c-.28 0-.5-.22-.5-.5v-6c0-.28.22-.5.5-.5s.5.22.5.5v6c0 .28-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Url(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="2" cy="10" r="1" fill="currentColor"/><circle cx="2" cy="6" r="1" fill="currentColor"/><path fill="currentColor" d="M4.5 14c-.06 0-.11 0-.17-.03a.501.501 0 0 1-.3-.64l4-11a.501.501 0 0 1 .94.34l-4 11c-.07.2-.27.33-.47.33m3 0c-.06 0-.11 0-.17-.03a.501.501 0 0 1-.3-.64l4-11a.501.501 0 0 1 .94.34l-4 11c-.07.2-.27.33-.47.33"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Usdc(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.5 13.47c0 .19-.15.29-.33.24a5.995 5.995 0 0 1 0-11.42c.18-.06.33.05.33.24v.46c0 .13-.1.27-.22.31C3.37 4 2 5.84 2 7.99s1.37 3.99 3.28 4.69c.12.04.22.19.22.31v.47Z"/><path fill="currentColor" fill-rule="evenodd" d="M7.5 11.75c0 .14-.11.25-.25.25h-.5c-.14 0-.25-.11-.25-.25v-.79c-1.09-.15-1.62-.76-1.77-1.59c-.03-.14.09-.27.23-.27h.57c.12 0 .22.08.24.2c.11.5.39.87 1.27.87c.65 0 1.1-.36 1.1-.9s-.27-.74-1.22-.9c-1.4-.19-2.06-.61-2.06-1.71c0-.85.64-1.5 1.63-1.65v-.77c0-.14.11-.25.25-.25h.5c.14 0 .25.11.25.25v.8c.81.14 1.32.6 1.48 1.36c.03.14-.08.28-.23.28h-.53c-.11 0-.21-.08-.24-.18c-.14-.48-.49-.69-1.08-.69c-.66 0-1 .32-1 .77c0 .47.19.71 1.21.86c1.37.19 2.08.58 2.08 1.75c0 .89-.66 1.61-1.69 1.77v.79Z"/><path fill="currentColor" fill-rule="evenodd" d="M8.83 13.71c-.18.06-.33-.05-.33-.24v-.46c0-.14.08-.27.22-.31C10.63 12 12 10.16 12 8.01s-1.37-3.99-3.28-4.69a.361.361 0 0 1-.22-.31v-.46c0-.19.15-.3.33-.24C11.25 3.08 13 5.35 13 8.02s-1.75 4.93-4.17 5.71Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vimeo(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.99 5.31c-.05 1.13-.87 2.67-2.45 4.62c-1.63 2.04-3.01 3.06-4.14 3.06c-.7 0-1.29-.62-1.77-1.87c-.32-1.14-.64-2.28-.97-3.42c-.36-1.24-.74-1.87-1.16-1.87c-.09 0-.4.18-.94.54l-.57-.7c.59-.5 1.18-1 1.75-1.51c.79-.66 1.38-1 1.78-1.04c.93-.08 1.51.53 1.73 1.85c.23 1.42.39 2.31.49 2.65c.27 1.18.57 1.77.89 1.77c.25 0 .63-.38 1.13-1.15c.5-.77.77-1.35.81-1.75c.07-.66-.2-.99-.81-.99c-.29 0-.58.06-.89.19c.59-1.86 1.72-2.77 3.38-2.71c1.23.03 1.81.8 1.74 2.31Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Visa(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 14h-13C.67 14 0 13.33 0 12.5v-9C0 2.67.67 2 1.5 2h13c.83 0 1.5.67 1.5 1.5v9c0 .83-.67 1.5-1.5 1.5M1.5 3c-.28 0-.5.22-.5.5v9c0 .28.22.5.5.5h13c.28 0 .5-.22.5-.5v-9c0-.28-.22-.5-.5-.5z"/><path fill="currentColor" d="M13.19 6.43h-.75c-.23 0-.4.07-.51.3l-1.44 3.28h1.02s.18-.44.21-.54h1.25c.04.12.12.53.12.53H14l-.81-3.56ZM12 8.73c.09-.21.39-1 .39-1c0 .02.09-.21.12-.33l.07.32l.23 1.04H12v-.02Zm-1.44.1c0 .74-.67 1.23-1.7 1.23c-.44 0-.86-.09-1.09-.19l.14-.81l.12.05c.32.14.53.19.91.19c.28 0 .58-.11.58-.35c0-.16-.12-.26-.51-.44c-.37-.18-.86-.46-.86-.96c0-.7.68-1.18 1.65-1.18c.37 0 .68.07.88.16l-.14.77l-.07-.07c-.18-.07-.4-.14-.74-.14c-.37.02-.54.18-.54.32c0 .16.21.28.54.44c.56.26.82.56.82.98ZM2 6.46l.02-.07h1.51c.21 0 .37.07.42.3l.33 1.58C3.95 7.43 3.17 6.74 2 6.46"/><path fill="currentColor" d="M6.4 6.43L4.87 9.99H3.83l-.88-2.98c.63.4 1.16 1.04 1.35 1.47l.11.37l.95-2.44H6.4zm.41-.02h.96l-.61 3.58H6.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Volumedown(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 11.46c-.33 0-.65-.11-.92-.32L3.83 9H2.5C1.67 9 1 8.33 1 7.5v-2C1 4.67 1.67 4 2.5 4h1.33l2.75-2.14c.46-.35 1.06-.42 1.58-.16c.52.25.84.77.84 1.35v6.91a1.495 1.495 0 0 1-1.5 1.5M2.5 5c-.28 0-.5.22-.5.5v2c0 .28.22.5.5.5H4c.11 0 .22.04.31.11l2.89 2.24c.15.12.35.14.53.05c.18-.09.28-.25.28-.45V3.04c0-.2-.1-.36-.28-.45a.487.487 0 0 0-.53.05L4.31 4.88a.51.51 0 0 1-.31.11H2.5Zm12 2h-4c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h4c.28 0 .5.22.5.5s-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Volumeup(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 11.5c-.33 0-.65-.11-.92-.32L3.83 9.04H2.5c-.83 0-1.5-.67-1.5-1.5v-2c0-.83.67-1.5 1.5-1.5h1.33L6.58 1.9c.46-.35 1.06-.42 1.58-.16c.52.25.84.77.84 1.35V10a1.495 1.495 0 0 1-1.5 1.5m-5-6.46c-.28 0-.5.22-.5.5v2c0 .28.22.5.5.5H4c.11 0 .22.04.31.11l2.89 2.24c.15.12.35.14.53.05c.18-.09.28-.25.28-.45v-6.9c0-.2-.1-.36-.28-.45a.487.487 0 0 0-.53.05L4.31 4.93a.51.51 0 0 1-.31.11zm10 4c-.28 0-.5-.22-.5-.5v-4c0-.28.22-.5.5-.5s.5.22.5.5v4c0 .28-.22.5-.5.5"/><path fill="currentColor" d="M14.5 7.04h-4c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h4c.28 0 .5.22.5.5s-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Warning(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.38 15H2.62c-.56 0-1.07-.28-1.36-.76s-.32-1.06-.07-1.56l.45.22l-.45-.22L6.57 1.89C6.84 1.34 7.39 1 8 1s1.16.34 1.43.89l5.38 10.8c.25.5.22 1.08-.07 1.56c-.29.48-.8.76-1.36.76ZM8 2c-.23 0-.43.12-.54.33l-5.38 10.8c-.1.19-.09.4.03.59c.11.18.3.29.51.29h10.76c.21 0 .4-.1.51-.29c.11-.18.12-.39.03-.59L8.54 2.33A.587.587 0 0 0 8 2"/><circle cx="8" cy="11" r=".62" fill="currentColor"/><path fill="currentColor" d="M8 9.12c-.28 0-.5-.22-.5-.5v-3c0-.28.22-.5.5-.5s.5.22.5.5v3c0 .28-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Week(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 16h-13C.67 16 0 15.33 0 14.5v-12C0 1.67.67 1 1.5 1h13c.83 0 1.5.67 1.5 1.5v12c0 .83-.67 1.5-1.5 1.5M1.5 2c-.28 0-.5.22-.5.5v12c0 .28.22.5.5.5h13c.28 0 .5-.22.5-.5v-12c0-.28-.22-.5-.5-.5z"/><path fill="currentColor" d="M4.5 4c-.28 0-.5-.22-.5-.5v-3c0-.28.22-.5.5-.5s.5.22.5.5v3c0 .28-.22.5-.5.5m7 0c-.28 0-.5-.22-.5-.5v-3c0-.28.22-.5.5-.5s.5.22.5.5v3c0 .28-.22.5-.5.5m4 2H.5C.22 6 0 5.78 0 5.5S.22 5 .5 5h15c.28 0 .5.22.5.5s-.22.5-.5.5"/><circle cx="4" cy="9" r="1" fill="currentColor"/><circle cx="8" cy="9" r="1" fill="currentColor"/><circle cx="12" cy="9" r="1" fill="currentColor"/><circle cx="4" cy="12" r="1" fill="currentColor"/><circle cx="8" cy="12" r="1" fill="currentColor"/><circle cx="12" cy="12" r="1" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Whatsapp(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.95 4.24C11.86 1 7.58.04 4.27 2.05C1.04 4.06 0 8.44 2.09 11.67l.17.26l-.7 2.62l2.62-.7l.26.17c1.13.61 2.36.96 3.58.96c1.31 0 2.62-.35 3.75-1.05c3.23-2.1 4.19-6.39 2.18-9.71Zm-1.83 6.74c-.35.52-.79.87-1.4.96c-.35 0-.79.17-2.53-.52c-1.48-.7-2.71-1.84-3.58-3.15c-.52-.61-.79-1.4-.87-2.19c0-.7.26-1.31.7-1.75c.17-.17.35-.26.52-.26h.44c.17 0 .35 0 .44.35c.17.44.61 1.49.61 1.58c.09.09.05.76-.35 1.14c-.22.25-.26.26-.17.44c.35.52.79 1.05 1.22 1.49c.52.44 1.05.79 1.66 1.05c.17.09.35.09.44-.09c.09-.17.52-.61.7-.79c.17-.17.26-.17.44-.09l1.4.7c.17.09.35.17.44.26c.09.26.09.61-.09.87Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Won(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 15c-3.86 0-7-3.14-7-7s3.14-7 7-7s7 3.14 7 7s-3.14 7-7 7M8 2C4.69 2 2 4.69 2 8s2.69 6 6 6s6-2.69 6-6s-2.69-6-6-6"/><path fill="currentColor" d="M10 12.5H9c-.25 0-.46-.18-.49-.42L8 8.79l-.51 3.29c-.04.24-.25.42-.49.42H6a.51.51 0 0 1-.49-.39l-1.5-6.5c-.06-.27.11-.54.37-.6c.27-.06.54.11.6.37l1.41 6.11h.17l.93-6.08a.495.495 0 0 1 .98 0l.93 6.08h.17l1.41-6.11a.49.49 0 0 1 .6-.37c.27.06.44.33.37.6l-1.5 6.5c-.05.23-.25.39-.49.39Z"/><path fill="currentColor" d="M12.5 9h-9c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h9c.28 0 .5.22.5.5s-.22.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wordpress(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.06 1C5.3 1 2.71 2.62 1.52 5.2h3.35v.37c-1.03-.12-.94.41-.79.97l2.2 5.48l1.61-3.77l-.74-1.77c-.27-.47-.27-1-1.38-.95l.02-.34h4.91v.39c-.57.03-1.27-.14-.9 1.03l1.99 5.19l1.09-3.03c.36-1.27.02-2.49-.4-3.69c-.23-.8-.13-1.47.93-1.74C11.8 1.71 9.91.99 8.08.99Zm6.68 4.79l-3.18 8.26c.88-.76 1.78-1.03 2.88-3.13c.73-1.74.66-3.36.3-5.13m-13.55.13c-.31 2.44-.28 3.15.49 5.14c1 1.91 1.78 2.31 2.88 3.06L1.19 5.93Zm7 3.09l-2.36 5.66c1.66.47 3.14.41 4.49-.11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yen(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 15c-3.86 0-7-3.14-7-7s3.14-7 7-7s7 3.14 7 7s-3.14 7-7 7M8 2C4.69 2 2 4.69 2 8s2.69 6 6 6s6-2.69 6-6s-2.69-6-6-6"/><path fill="currentColor" d="M9.62 11H6.37c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h3.25c.28 0 .5.22.5.5s-.22.5-.5.5m0-2H6.37c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h3.25c.28 0 .5.22.5.5s-.22.5-.5.5"/><path fill="currentColor" d="M8 12.5c-.28 0-.5-.22-.5-.5V8.5c0-.28.22-.5.5-.5s.5.22.5.5V12c0 .28-.22.5-.5.5"/><path fill="currentColor" d="M8 9a.42.42 0 0 1-.26-.08a.501.501 0 0 1-.16-.69l2.5-4a.501.501 0 0 1 .85.53l-2.5 4c-.09.15-.26.24-.42.24Z"/><path fill="currentColor" d="M8 9a.47.47 0 0 1-.42-.24l-2.5-4a.495.495 0 0 1 .16-.69c.23-.15.54-.08.69.16l2.5 4c.15.23.08.54-.16.69c-.08.05-.17.08-.26.08Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Youtube(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1.5c-6.88 0-7 .62-7 5.5s.12 5.5 7 5.5s7-.62 7-5.5s-.12-5.5-7-5.5m2.24 5.74L7.1 8.74c-.28.13-.5-.02-.5-.33V5.59c0-.31.23-.46.5-.33l3.14 1.5c.28.13.28.35 0 .48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yuan(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 15c-3.86 0-7-3.14-7-7s3.14-7 7-7s7 3.14 7 7s-3.14 7-7 7M8 2C4.69 2 2 4.69 2 8s2.69 6 6 6s6-2.69 6-6s-2.69-6-6-6"/><path fill="currentColor" d="M9.62 11H6.37c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h3.25c.28 0 .5.22.5.5s-.22.5-.5.5m0-2H6.37c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h3.25c.28 0 .5.22.5.5s-.22.5-.5.5"/><path fill="currentColor" d="M8 12.5c-.28 0-.5-.22-.5-.5V8.5c0-.28.22-.5.5-.5s.5.22.5.5V12c0 .28-.22.5-.5.5"/><path fill="currentColor" d="M8 9a.42.42 0 0 1-.26-.08a.501.501 0 0 1-.16-.69l2.5-4a.501.501 0 0 1 .85.53l-2.5 4c-.09.15-.26.24-.42.24Z"/><path fill="currentColor" d="M8 9a.47.47 0 0 1-.42-.24l-2.5-4a.495.495 0 0 1 .16-.69c.23-.15.54-.08.69.16l2.5 4c.15.23.08.54-.16.69c-.08.05-.17.08-.26.08Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Zip(children ...ElementRenderer) *FormkitIcon {
	return &FormkitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 12.25V13H3.75v-.81l1.18-1.46v-.03H3.75V10H6v.7l-1.27 1.55zm1.12.75h.75v-3h-.75zM9 13h.75v-.9h.4c.6 0 1.1-.45 1.1-1.06s-.48-1.05-1.1-1.05H9v3Zm.75-1.59v-.72h.31c.2 0 .35.13.35.35s-.15.36-.35.36h-.31ZM5.5 2c-.28 0-.5-.22-.5-.5V1c0-.28.22-.5.5-.5s.5.22.5.5v.5c0 .28-.22.5-.5.5m0 3c-.28 0-.5-.22-.5-.5v-1c0-.28.22-.5.5-.5s.5.22.5.5v1c0 .28-.22.5-.5.5"/><path fill="currentColor" d="M12.5 16h-10c-.83 0-1.5-.67-1.5-1.5v-13C1 .67 1.67 0 2.5 0h7.09c.4 0 .78.16 1.06.44l2.91 2.91c.28.28.44.66.44 1.06V14.5c0 .83-.67 1.5-1.5 1.5M2.5 1c-.28 0-.5.22-.5.5v13c0 .28.22.5.5.5h10c.28 0 .5-.22.5-.5V4.41a.47.47 0 0 0-.15-.35L9.94 1.15A.51.51 0 0 0 9.59 1z"/><path fill="currentColor" d="M13.38 5h-2.91C9.66 5 9 4.34 9 3.53V.62c0-.28.22-.5.5-.5s.5.22.5.5v2.91c0 .26.21.47.47.47h2.91c.28 0 .5.22.5.5s-.22.5-.5.5m-9.63 5.38h1.22v.38H3.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
