package raphael

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = ""
	hAttr          = "1em"
	viewbox        = "0 0 32 32"
	fill           = "currentColor"
)

type RaphaelIcon struct {
	*SVGSVGElement
}

type RaphaelIconFn func(children ...ElementRenderer) *RaphaelIcon

var IconLookup = map[string]RaphaelIconFn{
	"acw":                  Acw,
	"alarm":                Alarm,
	"android":              Android,
	"anonymous":            Anonymous,
	"apple":                Apple,
	"apps":                 Apps,
	"arrowalt":             Arrowalt,
	"arrowdown":            Arrowdown,
	"arrowleft":            Arrowleft,
	"arrowleftTwo":         ArrowleftTwo,
	"arrowleftalt":         Arrowleftalt,
	"arrowright":           Arrowright,
	"arrowrightTwo":        ArrowrightTwo,
	"arrowup":              Arrowup,
	"aumade":               Aumade,
	"barchart":             Barchart,
	"bell":                 Bell,
	"biohazard":            Biohazard,
	"bolt":                 Bolt,
	"book":                 Book,
	"bookmark":             Bookmark,
	"books":                Books,
	"brush":                Brush,
	"bubble":               Bubble,
	"bucket":               Bucket,
	"bug":                  Bug,
	"bus":                  Bus,
	"calendar":             Calendar,
	"calendarTwo":          CalendarTwo,
	"car":                  Car,
	"cart":                 Cart,
	"ccw":                  Ccw,
	"charging":             Charging,
	"chat":                 Chat,
	"check":                Check,
	"checkbox":             Checkbox,
	"checked":              Checked,
	"chrome":               Chrome,
	"clip":                 Clip,
	"clock":                Clock,
	"cloud":                Cloud,
	"cloudTwo":             CloudTwo,
	"clouddown":            Clouddown,
	"cloudup":              Cloudup,
	"cloudy":               Cloudy,
	"code":                 Code,
	"codetalk":             Codetalk,
	"coffee":               Coffee,
	"commandline":          Commandline,
	"connect":              Connect,
	"contract":             Contract,
	"crop":                 Crop,
	"cross":                Cross,
	"crown":                Crown,
	"cube":                 Cube,
	"customer":             Customer,
	"db":                   Db,
	"detour":               Detour,
	"diagram":              Diagram,
	"disconnect":           Disconnect,
	"dockbottom":           Dockbottom,
	"dockleft":             Dockleft,
	"dockright":            Dockright,
	"docktop":              Docktop,
	"dollar":               Dollar,
	"download":             Download,
	"dribbble":             Dribbble,
	"dry":                  Dry,
	"edit":                 Edit,
	"employee":             Employee,
	"end":                  End,
	"ethernet":             Ethernet,
	"exchange":             Exchange,
	"exclamation":          Exclamation,
	"expand":               Expand,
	"export":               Export,
	"facebook":             Facebook,
	"fave":                 Fave,
	"feed":                 Feed,
	"ff":                   Ff,
	"filter":               Filter,
	"firefox":              Firefox,
	"fitocracy":            Fitocracy,
	"fiveHundredPx":        FiveHundredPx,
	"flag":                 Flag,
	"flagAlt":              FlagAlt,
	"flickr":               Flickr,
	"folder":               Folder,
	"font":                 Font,
	"fork":                 Fork,
	"forkAlt":              ForkAlt,
	"fullBattery":          FullBattery,
	"fullcube":             Fullcube,
	"future":               Future,
	"gear":                 Gear,
	"github":               Github,
	"githubalt":            Githubalt,
	"glasses":              Glasses,
	"globe":                Globe,
	"globealt":             Globealt,
	"globealtTwo":          GlobealtTwo,
	"gplus":                Gplus,
	"graphael":             Graphael,
	"green":                Green,
	"hail":                 Hail,
	"hammer":               Hammer,
	"hammerandscrewdriver": Hammerandscrewdriver,
	"hangup":               Hangup,
	"help":                 Help,
	"history":              History,
	"home":                 Home,
	"hp":                   Hp,
	"icons":                Icons,
	"ie":                   Ie,
	"ieNine":               IeNine,
	"imac":                 Imac,
	"import":               Import,
	"inbox":                Inbox,
	"info":                 Info,
	"inkscape":             Inkscape,
	"instagram":            Instagram,
	"ios":                  Ios,
	"ipad":                 Ipad,
	"iphone":               Iphone,
	"jigsaw":               Jigsaw,
	"jquery":               Jquery,
	"js":                   Js,
	"key":                  Key,
	"lab":                  Lab,
	"lamp":                 Lamp,
	"lampAlt":              LampAlt,
	"landing":              Landing,
	"landscapeOne":         LandscapeOne,
	"landscapeTwo":         LandscapeTwo,
	"linechart":            Linechart,
	"link":                 Link,
	"linkedin":             Linkedin,
	"linux":                Linux,
	"list":                 List,
	"loactionTwo":          LoactionTwo,
	"location":             Location,
	"lock":                 Lock,
	"locked":               Locked,
	"lowBattery":           LowBattery,
	"magic":                Magic,
	"magnet":               Magnet,
	"mail":                 Mail,
	"man":                  Man,
	"merge":                Merge,
	"mic":                  Mic,
	"micmute":              Micmute,
	"minus":                Minus,
	"music":                Music,
	"mute":                 Mute,
	"newwindow":            Newwindow,
	"no":                   No,
	"nodejs":               Nodejs,
	"nomagnet":             Nomagnet,
	"notebook":             Notebook,
	"noview":               Noview,
	"opensource":           Opensource,
	"opera":                Opera,
	"package":              Package,
	"page":                 Page,
	"pageTwo":              PageTwo,
	"paint":                Paint,
	"pallete":              Pallete,
	"palm":                 Palm,
	"paper":                Paper,
	"parent":               Parent,
	"pc":                   Pc,
	"pen":                  Pen,
	"pensil":               Pensil,
	"people":               People,
	"phone":                Phone,
	"photo":                Photo,
	"picker":               Picker,
	"picture":              Picture,
	"piechart":             Piechart,
	"plane":                Plane,
	"play":                 Play,
	"plugin":               Plugin,
	"plus":                 Plus,
	"power":                Power,
	"ppt":                  Ppt,
	"printer":              Printer,
	"question":             Question,
	"questionTwo":          QuestionTwo,
	"quote":                Quote,
	"rain":                 Rain,
	"raphael":              Raphael,
	"reflecth":             Reflecth,
	"reflectv":             Reflectv,
	"refresh":              Refresh,
	"resizeTwo":            ResizeTwo,
	"roadmap":              Roadmap,
	"rotate":               Rotate,
	"ruler":                Ruler,
	"run":                  Run,
	"rw":                   Rw,
	"safari":               Safari,
	"scissors":             Scissors,
	"screwdriver":          Screwdriver,
	"search":               Search,
	"sencha":               Sencha,
	"settings":             Settings,
	"settingsalt":          Settingsalt,
	"shuffle":              Shuffle,
	"skull":                Skull,
	"skype":                Skype,
	"slideshare":           Slideshare,
	"smallgear":            Smallgear,
	"smile":                Smile,
	"smileTwo":             SmileTwo,
	"snow":                 Snow,
	"speaker":              Speaker,
	"split":                Split,
	"star":                 Star,
	"starThree":            StarThree,
	"starThreeOff":         StarThreeOff,
	"starTwo":              StarTwo,
	"starTwoOff":           StarTwoOff,
	"staroff":              Staroff,
	"start":                Start,
	"sticker":              Sticker,
	"stop":                 Stop,
	"stopsign":             Stopsign,
	"stopwatch":            Stopwatch,
	"sun":                  Sun,
	"svg":                  Svg,
	"tag":                  Tag,
	"takeoff":              Takeoff,
	"talke":                Talke,
	"talkq":                Talkq,
	"taxi":                 Taxi,
	"temp":                 Temp,
	"terminal":             Terminal,
	"thunder":              Thunder,
	"ticket":               Ticket,
	"train":                Train,
	"trash":                Trash,
	"tshirt":               Tshirt,
	"twitter":              Twitter,
	"twitterbird":          Twitterbird,
	"umbrella":             Umbrella,
	"undo":                 Undo,
	"unlock":               Unlock,
	"usb":                  Usb,
	"user":                 User,
	"users":                Users,
	"video":                Video,
	"view":                 View,
	"vim":                  Vim,
	"volumeOne":            VolumeOne,
	"volumeThree":          VolumeThree,
	"volumeTwo":            VolumeTwo,
	"volumeZero":           VolumeZero,
	"warning":              Warning,
	"wheelchair":           Wheelchair,
	"windows":              Windows,
	"woman":                Woman,
	"wrench":               Wrench,
	"wrenchThree":          WrenchThree,
	"wrenchTwo":            WrenchTwo,
	"zoomin":               Zoomin,
	"zoomout":              Zoomout,
}

func Acw(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.275 3.85l1.695 8.56l1.875-1.643c2.31 3.59 1.72 8.415-1.584 11.317a8.74 8.74 0 0 1-7.874 1.908l-.84 3.396c3.75.93 7.89.066 11.02-2.672c4.768-4.173 5.52-11.22 1.94-16.28l2.028-1.774l-8.26-2.813zM8.155 20.23c-2.313-3.59-1.722-8.416 1.58-11.317a8.745 8.745 0 0 1 7.876-1.91l.843-3.397A12.249 12.249 0 0 0 7.43 6.28c-4.764 4.174-5.518 11.223-1.938 16.283l-2.026 1.772l8.26 2.812l-1.693-8.56l-1.88 1.645z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alarm(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.18 20.115a.998.998 0 0 0-.366 1.366c.275.48.89.643 1.365.367a1.001 1.001 0 0 0-1.001-1.732zm1-7.876a.999.999 0 1 0-1.001 1.731a.998.998 0 1 0 1-1.731zm-.555 4.803a1 1 0 1 0-2 0a1 1 0 0 0 2 0m7.687-13.086V3.03h1a.5.5 0 0 0 .5-.5v-.5c0-.274-.225-.5-.5-.5h-3.625c-.275 0-.5.226-.5.5v.5a.5.5 0 0 0 .5.5h1v.927c-6.868.424-12.31 6.11-12.313 13.085c.002 7.25 5.877 13.124 13.126 13.127c7.25-.004 13.124-5.88 13.125-13.128c0-6.975-5.444-12.662-12.313-13.085m-.812 23.21A10.14 10.14 0 0 1 5.375 17.041A10.14 10.14 0 0 1 15.5 6.916a10.139 10.139 0 0 1 10.124 10.125A10.137 10.137 0 0 1 15.5 27.165zm-3.438-4.17a1.001 1.001 0 0 0-1 1.733a1 1 0 0 0 1-1.732zm0-11.91a1 1 0 1 0-1-1.73a1 1 0 0 0 1 1.731zm10.76 2.884a1 1 0 1 0-1.366-.365c.276.477.888.64 1.366.365m-7.32 9.95a1 1 0 1 0-.002 1.998a1 1 0 0 0 .002-1.998m4.436-14.565a1.001 1.001 0 1 0-.996 1.734a1.001 1.001 0 0 0 .995-1.733zm3.44 6.687a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-.555 4.073a1.004 1.004 0 0 0-1.367.365a1 1 0 1 0 1.367-.364zM15.5 8.167c-.55 0-1 .448-1 1l-.465 7.343l-3.004 1.96a1.001 1.001 0 0 0 1.001 1.732l3.306-1.676c.055.006.11.017.166.017a1.5 1.5 0 0 0 1.5-1.5l-.5-7.876c0-.553-.45-1-1-1zm3.44 14.83a1 1 0 1 0 .999 1.733a1.001 1.001 0 0 0-1.001-1.732zM11.196 3.594c-.836-1.04-2.103-1.718-3.54-1.718a4.562 4.562 0 0 0-4.563 4.562c0 .957.297 1.843.8 2.576a14.178 14.178 0 0 1 7.303-5.42m15.91 5.42c.502-.732.8-1.618.8-2.575a4.563 4.563 0 0 0-4.563-4.562c-1.438 0-2.704.678-3.54 1.717a14.18 14.18 0 0 1 7.302 5.42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Android(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.13 11.126c-.894 0-1.624.73-1.624 1.624v6.496c0 .895.73 1.624 1.624 1.624c.893 0 1.624-.73 1.624-1.624V12.75c0-.894-.73-1.624-1.624-1.624M19.516 4.96l1.32-2.035a.245.245 0 0 0-.07-.338a.247.247 0 0 0-.34.072L19.06 4.764a8.48 8.48 0 0 0-3.06-.56c-1.094 0-2.13.2-3.058.56L11.574 2.66a.244.244 0 1 0-.41.265l1.32 2.035C10.3 5.97 8.812 7.888 8.77 10.098h14.464c-.043-2.21-1.53-4.13-3.716-5.138zm-6.618 3.102a.832.832 0 1 1 0-1.664a.832.832 0 0 1 0 1.664m6.204 0a.832.832 0 1 1 0-1.664a.832.832 0 0 1 0 1.664m6.768 3.064c-.894 0-1.624.73-1.624 1.624v6.496c0 .895.73 1.624 1.624 1.624s1.624-.73 1.624-1.624V12.75c0-.894-.73-1.624-1.624-1.624M8.756 22.904c0 .723.59 1.312 1.314 1.312h1.363v3.61c0 .896.73 1.624 1.625 1.624c.893 0 1.624-.73 1.624-1.624v-3.61h2.636v3.61c0 .895.73 1.623 1.625 1.623c.894 0 1.623-.73 1.623-1.625v-3.61h1.363c.722 0 1.31-.59 1.31-1.312V11.07H8.757z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Anonymous(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28.523 23.813c-.518-.51-6.795-2.938-7.934-3.396c-1.133-.45-1.585-1.697-1.585-1.697s-.51.282-.51-.51c0-.793.51.51 1.02-2.548c0 0 1.415-.397 1.134-3.68h-.34s.85-3.51 0-4.698c-.853-1.188-1.187-1.98-3.06-2.548c-1.87-.567-1.19-.454-2.548-.396c-1.36.057-2.492.793-2.492 1.188c0 0-.85.057-1.188.397c-.34.34-.906 1.924-.906 2.32s.283 3.06.566 3.624l-.337.11c-.283 3.284 1.132 3.682 1.132 3.682c.51 3.058 1.02 1.755 1.02 2.548c0 .792-.51.51-.51.51s-.453 1.246-1.585 1.697c-1.132.453-7.416 2.887-7.927 3.396c-.51.52-.453 2.896-.453 2.896h26.954s.063-2.378-.453-2.897zm-11.905-10.12c-.398-.25-.783-1.21-.783-1.64v-.236c-.105-.106-.574-.096-.67 0v.236c0 .43-.385 1.39-.783 1.64c-.4.25-1.61.237-2.084-.236c-.473-.473-.524-1.663-.643-1.78c-.118-.12-.185-.185-.185-.185l.03-.414s.84-.207 1.698-.207s1.803.503 1.803.503c.232-.074.785-.083.997 0c0 0 .945-.502 1.803-.502s1.7.208 1.7.208l.028.414l-.185.185c-.118.118-.17 1.308-.643 1.78c-.47.473-1.682.487-2.082.236z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Apple(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.32 10.85c-1.743 1.233-2.615 2.72-2.615 4.455c0 2.08 1.078 3.673 3.232 4.786c-.578 1.678-1.416 3.135-2.514 4.376c-1.097 1.24-2.098 1.862-3.004 1.862c-.428 0-1.01-.143-1.75-.423l-.353-.138c-.725-.28-1.363-.423-1.92-.423c-.525 0-1.1.11-1.725.33l-.445.16l-.56.23c-.44.176-.888.264-1.337.264c-1.06 0-2.228-.872-3.507-2.616c-1.843-2.498-2.764-5.22-2.764-8.167c0-2.095.573-3.78 1.724-5.06c1.15-1.28 2.673-1.92 4.568-1.92c.71 0 1.37.13 1.988.388l.423.172l.445.183c.396.167.716.25.96.25c.31 0 .658-.07 1.04-.216l.58-.23l.436-.16c.693-.25 1.46-.376 2.297-.376c1.992 0 3.59.758 4.8 2.274zm-4.705-7.563c.02.267.033.473.033.617c0 1.317-.48 2.473-1.438 3.467s-2.075 1.49-3.347 1.49a5.807 5.807 0 0 1-.058-.638c0-1.12.445-2.17 1.337-3.153c.89-.983 1.922-1.56 3.096-1.726c.082-.015.21-.033.377-.057"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Apps(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m24.36 18.424l-2.327 1.215c.708 1.173 1.384 2.28 1.844 3.032l2.043-1.066c-.382-.784-.954-1.954-1.56-3.182zm-5.217-3.736c.445.84 1.342 2.367 2.274 3.926l2.414-1.26c-.87-1.77-1.72-3.46-2.086-4.123c-.896-1.62-1.982-3.107-3.454-5.416c-1.673-2.625-3.462-5.492-4.052-4.947c-1.194.384 1.237 4.094 1.876 5.715c.616 1.565 1.877 3.93 3.03 6.106zm7.314 7.985l-1.96 1.022l1.98 4.598s.812.684 1.92.213c1.105-.47.81-1.706.81-1.706zM24.35 15.71l2.924 5.93h1.983v-5.93zm-6.01-.006h-4.726L10.19 21.64h11.66c-.29-.48-3.08-5.16-3.51-5.936m-15.11 5.91l3.44-5.904H2.083v5.93h1.133l.015-.027zm11.818-11.47a1.684 1.684 0 0 0-3.192-.748l2.976 1.572c.138-.243.216-.524.216-.823zm-.705 1.916l-3.188-1.684l-1.535 2.636l3.197 1.69zM3.193 26.886l-.385 1.108v.3l.298-.13l.725-.895l2.998-2.355l-3.137-1.65l-.498 3.62zM9.02 14.044l-4.757 8.17l3.23 1.706l4.728-8.186l-3.2-1.69z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arrowalt(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 1.466C7.973 1.466 1.466 7.973 1.466 16c0 8.027 6.507 14.534 14.534 14.534c8.027 0 14.534-6.507 14.534-14.534c0-8.027-6.507-14.534-14.534-14.534m-2.335 24.26l-3.536-3.54L16.315 16L10.13 9.81l3.535-3.536L23.39 16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arrowdown(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.037 11.166L14.5 22.36c.825 1.43 2.175 1.43 3 0l6.463-11.195c.826-1.43.15-2.598-1.5-2.598H9.537c-1.65 0-2.326 1.17-1.5 2.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arrowleft(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.834 8.037L9.64 14.5c-1.43.824-1.43 2.175 0 3l11.194 6.463c1.43.826 2.598.15 2.598-1.5V9.537c0-1.65-1.17-2.326-2.598-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowleftTwo(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.87 9.814L15.685 16l6.187 6.188l-3.535 3.537L8.612 16l9.723-9.724z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arrowleftalt(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 30.534c8.027 0 14.534-6.507 14.534-14.534c0-8.027-6.507-14.534-14.534-14.534C7.973 1.466 1.466 7.973 1.466 16c0 8.027 6.507 14.534 14.534 14.534m2.335-24.258l3.536 3.538L15.685 16l6.187 6.188l-3.535 3.537L8.612 16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arrowright(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.166 23.963L22.36 17.5c1.43-.824 1.43-2.175 0-3L11.165 8.037c-1.43-.826-2.598-.15-2.598 1.5v12.926c0 1.65 1.17 2.326 2.598 1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowrightTwo(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.13 22.186L16.315 16L10.13 9.81l3.535-3.536L23.39 16l-9.725 9.725z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arrowup(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.963 20.834L17.5 9.64c-.825-1.43-2.175-1.43-3 0L8.037 20.834c-.825 1.43-.15 2.598 1.5 2.598h12.926c1.65 0 2.325-1.17 1.5-2.598"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Aumade(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.513 24.52a.724.724 0 0 1-.022-.71l1.988-3.845a.727.727 0 1 1 1.293.67l-1.8 3.48l2.228 3.7h12.27L15.67 2.183L6.435 18.178c.434.242.91.48 1.39.654c.572.21 1.15.342 1.66.342c.275 0 .578-.078.915-.238c.337-.158.7-.396 1.073-.688c.75-.582 1.528-1.354 2.335-2.02c.54-.443 1.09-.845 1.706-1.1c.353-.145.73-.24 1.13-.24c.62 0 1.173.215 1.62.5c.45.288.814.647 1.11.996c.59.697.903 1.36.925 1.394c.18.36.02.8-.34.977a.741.741 0 0 1-.98-.34c0-.002 0-.004-.003-.007a.094.094 0 0 0-.018-.034l-.078-.146a6.519 6.519 0 0 0-.312-.496c-.27-.39-.668-.845-1.092-1.104c-.28-.178-.56-.272-.844-.272c-.216 0-.48.07-.788.23c-.31.152-.653.395-1.016.687c-.727.584-1.51 1.362-2.35 2.033c-.563.445-1.15.853-1.81 1.103a3.367 3.367 0 0 1-1.195.228c-.75 0-1.48-.18-2.165-.433a11.14 11.14 0 0 1-1.613-.764L.86 27.817h15.63l-1.977-3.296zm3.7-2.278a9.93 9.93 0 0 1 .964-1.848c.427-.627.957-1.232 1.646-1.646c.38-.23.812-.39 1.282-.438l-.604-.934a.73.73 0 0 1 .216-1.01a.735.735 0 0 1 1.013.217l1.544 2.39a.73.73 0 0 1-.962 1.039c-.354-.19-.646-.258-.9-.258c-.292 0-.563.084-.846.25a3.197 3.197 0 0 0-.813.72c-.52.606-.937 1.42-1.185 2.054a.736.736 0 0 1-.68.466a.717.717 0 0 1-.27-.056a.73.73 0 0 1-.406-.946z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Barchart(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.25 8.375V28h6.5V8.375zM12.25 28h6.5V4.125h-6.5zm-9 0h6.5V12.625h-6.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.264 20.958c-2.484-4.226-2.168-13.2-6.143-15.486A2.525 2.525 0 0 0 16 1.58a2.526 2.526 0 0 0-2.528 2.526c0 .505.15.973.406 1.367c-3.973 2.287-3.657 11.26-6.142 15.485c-2.15.565-3.486 1.353-3.486 2.222v1.125c0 1.604 3.877 2.938 9.077 3.283c-.003.048-.015.096-.015.145a2.689 2.689 0 0 0 5.376.001c0-.05-.012-.097-.015-.145c5.2-.35 9.077-1.688 9.077-3.283V23.18c0-.87-1.335-1.657-3.486-2.222M14.472 4.105a1.53 1.53 0 0 1 1.527-1.527a1.531 1.531 0 0 1 1.527 1.527a1.558 1.558 0 0 1-.36.974c-.36-.097-.746-.15-1.167-.15s-.807.054-1.167.15a1.524 1.524 0 0 1-.36-.975z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Biohazard(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M26.154 13.988a8.219 8.219 0 0 0-3.02-1.032a8.216 8.216 0 0 0 .617-3.13c0-4.4-3.437-8.002-7.77-8.265a6.318 6.318 0 0 1 5.802 6.293c0 3.373-2.653 6.123-5.983 6.294v1.292a1.908 1.908 0 0 1 1.486 2.542l1.18.683c1.827-2.758 5.51-3.658 8.41-1.98a6.312 6.312 0 0 1 2.507 8.253c2.003-3.9.61-8.734-3.23-10.95zm-22.032 2.6c2.92-1.687 6.628-.765 8.442 2.033l1.14-.656c-.07-.2-.108-.417-.108-.642c0-.91.638-1.67 1.49-1.86v-1.318a6.309 6.309 0 0 1-5.92-6.292a6.31 6.31 0 0 1 5.756-6.286c-4.312.285-7.73 3.875-7.73 8.258a8.24 8.24 0 0 0 .582 3.05c-1.004.147-2 .48-2.93 1.02c-3.813 2.2-5.21 6.985-3.265 10.87a6.31 6.31 0 0 1 2.542-8.177zm7.028-5.136a5.457 5.457 0 0 0 .759.758c.1.085.21.16.317.237a5.9 5.9 0 0 1 3.244-.97c1.202 0 2.313.358 3.243.97c.107-.077.217-.152.318-.236a5.672 5.672 0 0 0 1.083-1.188a7.832 7.832 0 0 0-4.643-1.53h-.184c-1.666.04-3.2.606-4.462 1.53c.105.148.213.292.326.43zm-1.88 5.335a5.509 5.509 0 0 0-1.038-.277a5.314 5.314 0 0 0-.535-.065a7.835 7.835 0 0 0 .998 4.786a7.836 7.836 0 0 0 3.645 3.26a5.537 5.537 0 0 0 .538-1.927a5.894 5.894 0 0 1-2.46-2.325a5.89 5.89 0 0 1-.784-3.294c-.12-.055-.24-.11-.364-.157zm8.852 5.87c.014.132.024.263.046.394c.03.178.067.353.113.525a6.037 6.037 0 0 0 .378 1.008a7.838 7.838 0 0 0 3.644-3.257a7.838 7.838 0 0 0 1-4.788a5.762 5.762 0 0 0-1.572.34c-.124.047-.24.105-.362.16a5.885 5.885 0 0 1-.784 3.292a5.887 5.887 0 0 1-2.462 2.325zm2.447 4.954a6.307 6.307 0 0 1-2.46-8.328l-1.193-.69c-.35.39-.854.635-1.417.635a1.91 1.91 0 0 1-1.436-.653l-1.146.666c1.475 2.96.414 6.598-2.488 8.272a6.312 6.312 0 0 1-8.386-1.935c2.38 3.667 7.25 4.87 11.08 2.658a8.228 8.228 0 0 0 2.338-2.018a8.282 8.282 0 0 0 13.502-.56a6.31 6.31 0 0 1-8.395 1.953z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bolt(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.727 18.242L4.792 27.208l8.966-8.966l-4.483-4.484l17.933-8.966l-8.966 8.966z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Book(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.754 4.626a.876.876 0 0 0-.802-.097L12.16 9.41c-.557.212-1.253.315-1.968.315c-.997.002-2.03-.202-2.747-.48a3.441 3.441 0 0 1-.624-.302c.057-.024.12-.05.194-.075L18.648 4.43l1.733.654V3.172a.869.869 0 0 0-.373-.714a.877.877 0 0 0-.802-.097L6.415 7.24c-.396.143-.733.313-1.02.565c-.284.244-.527.645-.523 1.09c0 .013.004.032.004.032v17.186c0 .008-.003.015-.003.02c0 .007.003.01.003.017v.017h.002c.028.6.37.983.7 1.255c1.033.803 2.768 1.252 4.613 1.274c.875 0 1.762-.116 2.584-.427l12.796-4.882a.863.863 0 0 0 .558-.81V5.342a.87.87 0 0 0-.374-.714zm-20.082 7.11a.865.865 0 0 1 .07.273l.003.053c.016.264.13.406.363.61c.783.627 2.382 1.08 4.083 1.094a6.77 6.77 0 0 0 1.932-.264v1.79c-.647.144-1.3.207-1.942.207c-1.674-.026-3.266-.353-4.51-1.053zm4.51 12.852c-1.675-.028-3.267-.354-4.51-1.055V20.82a.802.802 0 0 1 .07.276l.003.053c.018.266.13.407.364.612c.782.625 2.38 1.08 4.082 1.09c.67 0 1.327-.08 1.932-.26v1.788a8.907 8.907 0 0 1-1.943.208z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.396 1.84L6.076 25.987l7.34-4.566l1.187 8.564l11.32-24.146l-8.527-3.997zm1.735 7.394a1.125 1.125 0 0 1 .956-2.037a1.124 1.124 0 1 1-.955 2.036z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Books(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M26.68 7.858a.595.595 0 0 0-.607-.083l-9.66 4.183c-.42.183-.946.27-1.486.27c-.753.003-1.532-.172-2.075-.41a2.523 2.523 0 0 1-.47-.26c.04-.02.09-.042.145-.064l8.786-3.804l1.31.56V6.613a.775.775 0 0 0-.283-.612a.592.592 0 0 0-.605-.083l-9.66 4.183c-.298.12-.554.268-.77.483a1.327 1.327 0 0 0-.395.934c0 .01.003.027.003.027v14.73l-.002.02c0 .004.003.006.003.01v.016h.002c.02.515.28.843.528 1.075c.78.688 2.09 1.073 3.484 1.093c.66 0 1.33-.1 1.95-.366l9.663-4.184c.255-.11.422-.383.422-.692V8.47a.781.781 0 0 0-.283-.612m-6.127-2.8c-.017-.22-.108-.43-.27-.556a.595.595 0 0 0-.607-.083L10.016 8.6c-.42.182-.947.27-1.486.27c-.753.002-1.532-.173-2.075-.412a2.448 2.448 0 0 1-.47-.258c.04-.02.09-.042.145-.064l8.787-3.804l1.31.56V3.257a.776.776 0 0 0-.284-.612a.594.594 0 0 0-.606-.083l-9.66 4.184c-.298.12-.553.267-.77.483a1.327 1.327 0 0 0-.394.934c0 .012.003.028.003.028v14.777h.002c.02.515.28.843.528 1.075c.78.688 2.09 1.072 3.485 1.092a5.57 5.57 0 0 0 1.127-.122V11.544c-.01-.7.27-1.372.762-1.856a3.476 3.476 0 0 1 1.19-.756z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Brush(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.125 29.178l1.31-1.5l1.316 1.5l1.312-1.5l1.31 1.5l1.316-1.5l1.31 1.5l1.313-1.5l1.315 1.5l1.312-1.5l1.312 1.5l1.313-1.5l1.312 1.5v-8.52H8.125zm15.25-12.022c-.354 0-5.833-.166-5.833-2.917s.75-8.835.75-8.835s.25-2.583-2.292-2.583s-2.292 2.583-2.292 2.583s.75 6.083.75 8.834s-5.48 2.916-5.833 2.916c-.5 0-.5 1.166-.5 1.166v1.27h15.75v-1.27s0-1.166-.5-1.166zM16 8.03c-.62 0-1.125-2.19-1.125-2.81S15.38 4.093 16 4.093s1.125.504 1.125 1.125S16.62 8.03 16 8.03"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bubble(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 5.333c-7.732 0-14 4.7-14 10.5c0 1.982.74 3.833 2.016 5.414L2 25.667l5.613-1.44c2.34 1.316 5.237 2.106 8.387 2.106c7.732 0 14-4.7 14-10.5s-6.268-10.5-14-10.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bucket(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.59 6.787c-.25-.152-.505-.3-.76-.445c-3.833-2.154-8.235-3.31-9.47-1.32c-1.225 2.104 2.314 5.8 6.293 8.223c.082.05.167.098.25.146c5.463-.4 9.887.205 9.99 1.403a.591.591 0 0 1-.084.318c.017-.024.04-.044.056-.07c1.28-2.097-2.28-5.818-6.277-8.253zm-11.253.38l-.722 1.52c-4.34 2.746-6.542 6.192-5.484 8.624c.19.44.483.813.847 1.138l.456-.96a2.212 2.212 0 0 1-.384-.576c-.733-1.682.765-4.188 3.706-6.417L5.433 17.49c1.492 1.688 5.748 1.747 10.276.153c-.038-.354.03-.722.23-1.05a1.69 1.69 0 0 1 2.882 1.758a1.689 1.689 0 0 1-2.32.562a1.688 1.688 0 0 1-.393-.344c-4.355 1.56-8.373 1.643-10.554.314a4.087 4.087 0 0 1-.58-.436l-.124.26c-1.088 1.785 1.883 4.916 5.23 6.957c3.347 2.04 7.493 3.246 8.552 1.502l7.77-10.204c-2.48.385-6.154-.962-9.272-2.863c-3.118-1.9-6-4.55-6.795-6.93z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bug(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m28.59 10.903l-5.83 1.612a10.241 10.241 0 0 0-2.31-3.628l3.082-5.44a1 1 0 1 0-1.74-.988l-2.924 5.16a8.104 8.104 0 0 0-1.9-.895a3.577 3.577 0 0 0-3.058-2.562l-.003-.012c-.06-.24-.093-.46-.098-.65a1.65 1.65 0 0 1 .045-.478c.037-.13.08-.235.125-.317c.146-.26.34-.43.577-.51c.023.282.142.483.352.602a.838.838 0 0 0 .545.086c.21-.03.376-.152.496-.363a.757.757 0 0 0 .065-.607a.837.837 0 0 0-.414-.507a.957.957 0 0 0-.48-.13a1.3 1.3 0 0 0-.912.388c-.13.128-.24.27-.33.425a2.12 2.12 0 0 0-.25.782a3.938 3.938 0 0 0-.008.84c.017.16.06.3.094.45c-.375-.022-.758 0-1.14.107c-.482.132-.913.36-1.28.652c-.052-.172-.098-.344-.18-.518a3.875 3.875 0 0 0-.438-.716a2.074 2.074 0 0 0-.618-.543a1.983 1.983 0 0 0-.5-.194a1.285 1.285 0 0 0-.979.138a.947.947 0 0 0-.348.357a.845.845 0 0 0-.094.647a.76.76 0 0 0 .367.49c.21.12.415.138.61.056a.836.836 0 0 0 .426-.355c.118-.21.117-.443-.008-.695c.244-.055.497-.008.757.14c.08.046.17.116.27.21c.096.09.192.22.285.387c.094.166.18.368.25.608c.014.044.024.098.036.146a3.562 3.562 0 0 0-1.097 3.356a7.912 7.912 0 0 0-1.457 2.035l-5.16-2.925a1.004 1.004 0 0 0-1.364.377a1.007 1.007 0 0 0 .378 1.367l5.44 3.08a10.313 10.313 0 0 0-.116 4.298L1.926 18.28a1 1 0 0 0-.698 1.232a.998.998 0 0 0 1.23.694l5.88-1.627c.503 1.056 1.363 2.28 2.37 3.442l-3.193 5.64a.998.998 0 0 0 1.136 1.456c.25-.068.47-.23.605-.47l2.895-5.11c2.7 2.594 5.684 4.123 5.778 1.053c1.598 2.56 3.45-.337 4.502-3.975l5.203 2.947c.24.138.514.162.762.094a1 1 0 0 0 .226-1.833l-5.7-3.23c.29-1.504.422-2.982.32-4.137l5.885-1.627a1 1 0 0 0-.534-1.927z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bus(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M30.17 7.314c-.024-.274-.214-.498-.42-.498s-.375-.17-.375-.375s-.222-.336-.493-.29l-1.472.25c-.296-1.793-.74-2.914-1.41-3.584c-2-2-18-2-20 0c-.67.67-1.114 1.79-1.41 3.583l-1.472-.25c-.27-.046-.493.085-.493.29s-.17.376-.375.376s-.396.225-.42.498l-.47 5.066a.447.447 0 0 0 .453.498h1.062c.275 0 .52-.224.546-.498l.394-4.232a.459.459 0 0 1 .54-.415l.054.01C4.008 11.395 4 29.682 4 29.682c0 .554.14 1 .313 1h2.562c.173 0 .313-.446.313-1V27.79c4.643.7 12.982.7 17.625 0v1.89c0 .553.14 1 .312 1h2.562c.172 0 .312-.447.312-1c0 0-.008-18.283-.408-21.937l.054-.01a.461.461 0 0 1 .54.416l.393 4.23c.024.275.27.5.545.5h1.062a.45.45 0 0 0 .454-.5l-.47-5.066zM7.126 23.37a1.811 1.811 0 1 1-.002-3.62a1.811 1.811 0 0 1 .002 3.619zm-2.083-7.393c.1-7.435.45-11.238 1.665-12.454c.487-.486 3.777-1.207 9.293-1.207c5.516 0 8.806.72 9.293 1.207c1.217 1.216 1.564 5.02 1.665 12.455c-1.175.473-4.904 1.025-10.958 1.025c-6.05 0-9.778-.55-10.958-1.026m18.02 5.582a1.813 1.813 0 0 1 3.624-.001a1.811 1.811 0 0 1-3.624 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 4.582h-2v3.335h2zm3.416 1.166H23v3.17h-4v-3.17h-6v3.168H9.002V5.748h-2.42v21.555h18.834zm-.998 20.555H7.584V13.988h16.833v12.315zM12 4.583h-2v3.334h2V4.582zm7.428 19.38h1.568v-7.79H19.72c0 .068-.022.173-.062.313c-.066.232-.168.42-.3.56a1.42 1.42 0 0 1-.75.407c-.192.043-.53.075-1.013.097v1.042h1.832v5.37zm-5.755-1.054c-.49 0-.827-.19-1.013-.565c-.1-.203-.15-.46-.15-.773h-1.504c.025.62.15 1.12.376 1.504c.43.72 1.194 1.08 2.296 1.08c.895 0 1.57-.25 2.026-.75c.455-.5.684-1.078.684-1.736c0-.627-.195-1.12-.586-1.482c-.26-.24-.46-.36-.602-.36c.187-.07.365-.206.537-.403c.272-.314.408-.7.408-1.16c0-.647-.228-1.164-.684-1.55c-.455-.385-1.055-.577-1.8-.577c-.4 0-.737.05-1.013.146a2.035 2.035 0 0 0-.714.42c-.27.257-.465.538-.59.842a3.88 3.88 0 0 0-.2 1.102h1.43c-.007-.384.074-.69.244-.92c.17-.228.435-.343.795-.343c.314 0 .56.094.73.28c.175.186.26.427.26.725c0 .458-.168.763-.507.913c-.196.09-.543.138-1.04.145v1.096c.508 0 .88.05 1.115.146c.414.172.62.514.62 1.026c0 .387-.11.683-.334.89c-.222.2-.483.303-.783.303z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarTwo(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 4.582h-2v3.335h2zm-10 0h-2v3.335h2zm13.416 1.166H23v3.17h-4v-3.17h-6v3.168H9.002V5.748h-2.42v21.555h18.834zM11.033 26.303h-3.45v-3.44h3.45zm0-4.44h-3.45v-3.435h3.45v3.434zm0-4.434h-3.45v-3.442h3.45v3.44zm4.468 8.873h-3.467v-3.44H15.5zm0-4.44h-3.467v-3.435H15.5v3.434zm0-4.434h-3.467v-3.442H15.5v3.44zm4.47 8.873H16.5v-3.44h3.47v3.44zm0-4.44H16.5v-3.435h3.47v3.434zm0-4.434H16.5v-3.442h3.47v3.44zm4.448 8.873H20.97v-3.44h3.448v3.44zm0-4.44H20.97v-3.435h3.448v3.434zm0-4.434H20.97v-3.442h3.448v3.44z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Car(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28.59 10.78h-2.242a.496.496 0 0 0-.333.134c-.716-1.143-1.457-2.058-2.032-2.633c-2-2-14-2-16 0c-.573.574-1.31 1.483-2.022 2.618a.483.483 0 0 0-.308-.117H3.41c-.275 0-.5.226-.5.5v1.01c0 .274.22.54.49.593l1.36.26c-1.174 2.618-1.866 5.876-.778 9.14v1.937c0 .553.14 1 .313 1h2.562c.173 0 .313-.447.313-1v-1.584c2.298.22 5.55.46 8.812.46c3.232 0 6.52-.236 8.814-.454v1.578c0 .553.14 1 .312 1h2.562c.172 0 .312-.447.312-1l-.002-1.938c1.087-3.26.397-6.516-.775-9.134l1.392-.265a.63.63 0 0 0 .49-.594v-1.01a.498.498 0 0 0-.497-.5zM7.107 18.907a1.813 1.813 0 0 1 0-3.624a1.813 1.813 0 0 1-.001 3.624zm-1.524-5.19C6.543 11.52 7.88 9.8 8.69 8.988c.584-.585 3.34-1.207 7.292-1.207c3.953 0 6.708.623 7.293 1.208c.81.81 2.146 2.53 3.106 4.728c-2.132.236-6.285-.31-10.398-.31s-8.266.546-10.4.31m19.274 5.19a1.813 1.813 0 0 1 0-3.624a1.813 1.813 0 0 1-.001 3.624z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cart(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M29.02 11.754L8.416 9.474L7.16 4.715a.758.758 0 0 0-.727-.558H3.34a1.214 1.214 0 0 0-.963-.49a1.24 1.24 0 1 0 0 2.483c.4 0 .738-.2.965-.492h2.512l5.23 19.8a3.282 3.282 0 0 0-.89 2.242a3.29 3.29 0 0 0 3.292 3.293a3.296 3.296 0 0 0 3.297-3.293a3.19 3.19 0 0 0-.093-.743h5.533a3.25 3.25 0 0 0-.092.743c0 1.82 1.476 3.293 3.296 3.293S28.72 29.52 28.72 27.7a3.296 3.296 0 0 0-3.294-3.297c-.95 0-1.8.41-2.402 1.053h-7.136a3.276 3.276 0 0 0-2.402-1.053c-.38 0-.738.078-1.077.196l-.182-.686H26.81a2.5 2.5 0 0 0 2.39-1.96l1.575-7.798a2.17 2.17 0 0 0 .04-.414a1.995 1.995 0 0 0-1.795-1.988zm-3.592 16.24a.298.298 0 0 1-.297-.295c.003-.166.135-.298.298-.298s.295.132.297.297a.298.298 0 0 1-.297.294zm1.78-7.495l.948-.95l-.318 1.58zm-14.453-9.037L13.79 12.5l-1.29 1.29l-1.293-1.29l1.087-1.088zm4.498.498l.538.54l-1.29 1.29l-1.293-1.29l.688-.69l1.358.15zM9.63 14.076l.87-.868l1.29 1.292l-1.29 1.29l-.565-.563l-.304-1.152zm-.295-1.12l-.328-1.24l.785.785zM21.79 16.5l-1.29 1.29l-1.293-1.29l1.292-1.293zm-.583-2l1.292-1.292l1.29 1.292l-1.29 1.292zM18.5 15.79l-1.293-1.29l1.292-1.293l1.29 1.292zm-.71.71l-1.29 1.29l-1.292-1.29l1.292-1.293zm-3.29-.71l-1.293-1.29l1.292-1.293l1.29 1.292zm-.71.71l-1.29 1.29l-1.293-1.29l1.292-1.293zm-3.29.707l1.29 1.292l-.784.783l-.54-2.044l.033-.033zm.802 3.197l1.197-1.197l1.29 1.292l-1.29 1.29l-1.13-1.13zm1.906-1.905l1.29-1.293l1.293 1.292l-1.29 1.29l-1.292-1.29zm3.292.707l1.292 1.292l-1.292 1.29l-1.292-1.29zm.708-.708l1.292-1.293l1.29 1.292l-1.29 1.29zm3.29.707l1.293 1.292l-1.29 1.29l-1.292-1.292zm.71-.708l1.29-1.293l1.293 1.292l-1.29 1.29zm2-2l1.29-1.293l1.293 1.292l-1.29 1.29zm2-2l1.29-1.293L27.79 14.5l-1.29 1.292l-1.293-1.293zm-.71-.708l-1.155-1.156l2.082.23zM21.792 12.5l-1.29 1.292l-1.293-1.292l.29-.29l2.253.25zm-7.29-.71l-.152-.15l.273.03l-.12.12zm-4 .002l-.65-.65l1.17.13zm4 9.415l1.205 1.205h-2.41zm4 0l1.205 1.206h-2.412zm4 0l1.207 1.207h-2.414zm.707-.708l1.292-1.293l1.29 1.292l-1.29 1.29zm2-2l1.292-1.292l1.29 1.29l-1.29 1.293l-1.293-1.29zm3.292-.71l-1.292-1.29l1.29-1.292l.445.444l-.43 2.124l-.014.015zm.5-4.5l-.5.5l-.66-.657l1.017.112c.054.008.1.026.144.044zM13.488 27.993a.297.297 0 0 1 0-.593a.296.296 0 0 1 0 .591zm13.323-5.58h-1.517l1.207-1.207l.93.93c-.187.17-.423.29-.62.277"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ccw(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.25 15.5a8.766 8.766 0 0 1-8.75 8.75a8.7 8.7 0 0 1-6.366-2.764l2.068-1.442l-7.9-3.703l.743 8.695l2.193-1.53a12.217 12.217 0 0 0 9.26 4.243c6.767 0 12.25-5.482 12.25-12.25h-3.5zM15.5 6.75a8.704 8.704 0 0 1 6.366 2.764l-2.068 1.443l7.9 3.7l-.745-8.692l-2.192 1.53A12.22 12.22 0 0 0 15.5 3.248C8.733 3.25 3.25 8.733 3.25 15.5h3.5a8.76 8.76 0 0 1 8.75-8.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Charging(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27.27 13.5h-1.103v-1.416a1 1 0 0 0-1-1H5.25a1 1 0 0 0-1 1v7.832a1 1 0 0 0 1 1h19.917a1 1 0 0 0 1-1V18.5h1.104c.266 0 .48-.448.48-1v-3c0-.552-.214-1-.48-1zm-8.745 3.342l-6.733 3.366l3.366-3.366l-1.683-1.684l6.733-3.366l-3.366 3.366z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chat(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.985 5.972c-7.563 0-13.695 4.077-13.695 9.106c0 2.877 2.013 5.44 5.147 7.108c-.446 1.48-1.336 3.117-3.056 4.566c0 0 4.016-.266 6.852-3.143c.163.04.332.07.497.106a4.49 4.49 0 0 1-.247-1.443c0-3.393 3.776-6.05 8.6-6.05c3.463 0 6.378 1.376 7.75 3.406c1.168-1.34 1.847-2.893 1.847-4.553c0-5.028-6.132-9.105-13.695-9.105zM27.68 22.274c0-2.79-3.4-5.053-7.6-5.053c-4.195 0-7.598 2.264-7.598 5.054c0 2.79 3.403 5.053 7.6 5.053c.928 0 1.813-.116 2.636-.32c1.573 1.598 3.8 1.745 3.8 1.745c-.953-.804-1.446-1.713-1.694-2.534c1.738-.925 2.856-2.347 2.856-3.944z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.38 14.73l2.828-2.83l7.75 7.748l12.92-12.915l2.83 2.828l-15.75 15.748z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Checkbox(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M26 27.5H6A1.5 1.5 0 0 1 4.5 26V6c0-.83.67-1.5 1.5-1.5h20c.828 0 1.5.67 1.5 1.5v20a1.5 1.5 0 0 1-1.5 1.5m-18.5-3h17v-17h-17z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Checked(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M29.548 3.043a2.5 2.5 0 0 0-3.513.4L16 16.067l-3.508-4.414a2.5 2.5 0 0 0-3.915 3.112l5.465 6.875c.474.597 1.195.943 1.957.943s1.482-.35 1.957-.944L29.95 6.555c.86-1.08.68-2.654-.402-3.513zM24.5 24.5h-17v-17h12.756l2.385-3H6c-.83 0-1.5.67-1.5 1.5v20c0 .828.67 1.5 1.5 1.5h20a1.5 1.5 0 0 0 1.5-1.5V12.85l-3 3.774z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chrome(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.318 7.677a.583.583 0 0 1 .23-.046h11.948a14.211 14.211 0 0 0-11.95-6.505A14.209 14.209 0 0 0 3.588 7.648l4.29 7.43a7.675 7.675 0 0 1 7.44-7.4zM28.196 8.84h-8.58a7.666 7.666 0 0 1 3.606 6.506c0 1.32-.334 2.564-.92 3.65a.675.675 0 0 1-.074.208L16.255 29.55c7.526-.367 13.514-6.586 13.514-14.204c0-2.344-.57-4.555-1.574-6.506zm-12.65 14.182a7.669 7.669 0 0 1-6.532-3.646a.63.63 0 0 1-.15-.17L2.89 8.855a14.155 14.155 0 0 0-1.565 6.49c0 7.625 6 13.847 13.534 14.206l4.286-7.425a7.651 7.651 0 0 1-3.6.895zM9.08 15.347c0 1.788.723 3.4 1.894 4.573a6.44 6.44 0 0 0 4.573 1.895c1.788 0 3.4-.723 4.573-1.895s1.895-2.785 1.895-4.573c0-1.788-.723-3.4-1.895-4.573a6.447 6.447 0 0 0-4.573-1.894c-1.788 0-3.4.723-4.573 1.894a6.452 6.452 0 0 0-1.894 4.573"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clip(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.898 6.135a3.501 3.501 0 0 0-4.884.808l-8.832 12.33a2.5 2.5 0 1 0 4.064 2.91l6.236-8.705l-.813-.583l-6.235 8.707a1.502 1.502 0 0 1-2.092.347a1.502 1.502 0 0 1-.345-2.094l8.83-12.33h.002l-.002-.002a2.506 2.506 0 0 1 3.49-.576a2.506 2.506 0 0 1 .576 3.49v-.002l-9.68 13.516a3.506 3.506 0 0 1-4.884.81a3.507 3.507 0 0 1-.807-4.886l7.035-9.822l-.813-.582l-7.035 9.822a4.497 4.497 0 0 0 1.04 6.277a4.496 4.496 0 0 0 6.277-1.036l9.68-13.516a3.501 3.501 0 0 0-.81-4.883z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 2.374C8.25 2.374 2.376 8.25 2.374 15.5C2.376 22.748 8.25 28.623 15.5 28.627c7.25-.004 13.124-5.88 13.125-13.127c0-7.25-5.876-13.125-13.125-13.126m0 23.25C9.91 25.613 5.385 21.09 5.375 15.5A10.141 10.141 0 0 1 15.5 5.374a10.139 10.139 0 0 1 10.124 10.125c-.01 5.59-4.533 10.115-10.124 10.123zM8.625 15.5a1 1 0 1 0-2 0a1 1 0 0 0 2 0m-.446 3.072a1 1 0 1 0 1.365.364a1.004 1.004 0 0 0-1.367-.365zm1-7.876a1 1 0 1 0-.998 1.732a1 1 0 0 0 .997-1.732zm13.642 1.732a1 1 0 1 0-.998-1.731a1 1 0 0 0 .998 1.731m-10.76 9.027a1.001 1.001 0 0 0-1 1.732a1 1 0 0 0 1-1.732m0-11.91a.999.999 0 0 0 .366-1.366a.999.999 0 1 0-.366 1.364zm10.76 9.027a1.003 1.003 0 0 0-1.366.365a1.001 1.001 0 0 0 1.732 1.002a1.003 1.003 0 0 0-.365-1.368zM19.94 7.812a1.001 1.001 0 0 0-.999 1.733a1 1 0 1 0 .998-1.733zm3.44 6.688a1 1 0 1 0 0 2a1 1 0 0 0 0-2M15.5 6.624c-.55 0-1 .448-1 1l-.465 7.343l-3.004 1.96a1 1 0 0 0 1.001 1.732l3.306-1.677c.054.007.108.017.165.017c.83 0 1.5-.67 1.5-1.5l-.5-7.876a1 1 0 0 0-1-1zm0 15.753c-.55 0-1 .447-1 1s.45 1 1 1s1-.447 1-1s-.447-1-1-1m3.44-.922a1 1 0 1 0 1.364.367a1.002 1.002 0 0 0-1.366-.367z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.345 13.904a6.25 6.25 0 0 0-6.22-6.84a6.244 6.244 0 0 0-5.847 4.058a3.205 3.205 0 0 0-2.372-1.06a3.22 3.22 0 0 0-3.22 3.22c0 .21.024.415.063.613c-2.373.39-4.188 2.436-4.188 4.918a5 5 0 0 0 5 5h15.875a4.998 4.998 0 0 0 .908-9.91z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudTwo(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.562 24.812a6 6 0 0 1-6-6a5.984 5.984 0 0 1 4.127-5.684a4.22 4.22 0 0 1 4.215-4.066c.73 0 1.415.19 2.01.517a7.236 7.236 0 0 1 6.208-3.518a7.252 7.252 0 0 1 7.248 7.08a5.99 5.99 0 0 1 4.065 5.67a6 6 0 0 1-6 6H7.562zm16.6-9.925a1.001 1.001 0 0 1-.814-1.08c.017-.17.027-.335.027-.496a5.257 5.257 0 0 0-5.25-5.248a5.246 5.246 0 0 0-4.912 3.41a.997.997 0 0 1-1.675.32a2.199 2.199 0 0 0-1.633-.73a2.224 2.224 0 0 0-2.22 2.218c0 .136.017.276.045.424a1.005 1.005 0 0 1-.82 1.176c-1.9.313-3.352 1.95-3.35 3.93a4.008 4.008 0 0 0 4 4.002h15.875a4.008 4.008 0 0 0 4-4c.003-1.958-1.41-3.58-3.272-3.925"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clouddown(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.345 13.904a6.25 6.25 0 0 0-6.22-6.84a6.244 6.244 0 0 0-5.847 4.058a3.205 3.205 0 0 0-2.372-1.06a3.22 3.22 0 0 0-3.22 3.22c0 .21.024.415.063.613c-2.373.39-4.188 2.436-4.188 4.918a5 5 0 0 0 5 5h3.404l-.707-.707c-.378-.377-.586-.88-.586-1.413s.208-1.035.585-1.412l.555-.556a2.106 2.106 0 0 1 1.55-.626v-.472c0-1.104.9-2.002 2-2.002h3.267c1.103 0 2 .898 2 2.002v.472c.027-.002.054-.002.08-.002c.534 0 1.07.23 1.47.63l.558.552c.78.78.78 2.05 0 2.828l-.706.707h2.403a4.998 4.998 0 0 0 .908-9.91zm-3.312 7.082l-.556-.555c-.39-.388-.964-.45-1.276-.136c-.31.312-.567.118-.567-.432v-1.238c0-.55-.45-1-1-1h-3.265c-.55 0-1 .45-1 1v1.238c0 .55-.256.744-.57.432c-.31-.313-.886-.252-1.275.137l-.556.556a1 1 0 0 0 0 1.413l4.327 4.33c.194.194.45.29.707.29s.513-.096.708-.29l4.327-4.33a1 1 0 0 0-.002-1.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloudup(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.345 13.904a6.25 6.25 0 0 0-6.22-6.84a6.244 6.244 0 0 0-5.847 4.058a3.205 3.205 0 0 0-2.372-1.06a3.22 3.22 0 0 0-3.22 3.22c0 .21.024.415.063.613c-2.373.39-4.188 2.436-4.188 4.918a5 5 0 0 0 5 5h2.312a1.976 1.976 0 0 1 .387-2.274l4.326-4.33a1.979 1.979 0 0 1 1.413-.585c.54 0 1.04.21 1.416.587l4.323 4.33a1.976 1.976 0 0 1 .384 2.272h1.312a5 5 0 0 0 5-5a4.992 4.992 0 0 0-4.09-4.91zm-7.64 4.012c-.192-.195-.45-.29-.705-.29s-.512.095-.707.29l-4.327 4.33a1 1 0 0 0 0 1.414l.557.555c.39.39.964.45 1.276.137s.567-.12.567.432v1.238c0 .55.45 1 1 1h3.265c.55 0 1-.45 1-1v-1.238c0-.55.256-.744.57-.432c.31.312.886.252 1.275-.137l.556-.555a1 1 0 0 0 0-1.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloudy(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.378 6.78c.41.99 1.938.347 1.524-.647l-.582-1.4c-.29-.694-.875-3.232-2.042-2.746c-1.03.433-.128 1.846.142 2.494l.958 2.3m6.422.443c1.094.453 1.538-1.55 1.813-2.216c.28-.677 1.478-2.565.357-3.03c-1.092-.452-1.537 1.55-1.813 2.217c-.28.678-1.477 2.563-.357 3.03m-2.663-.532c1.183 0 .83-2.02.83-2.742c0-.732.382-2.935-.83-2.935c-1.183 0-.828 2.02-.828 2.742c0 .733-.384 2.935.827 2.935m4.92 2.037c.853.85 2.143-.973 2.66-1.49c.512-.514 2.187-1.688 1.352-2.525c-.835-.836-2.014.843-2.523 1.353c-.52.517-2.35 1.806-1.49 2.66m1.508 2.258c.448 1.09 2.183-.01 2.85-.286c.675-.28 2.857-.77 2.393-1.89c-.455-1.09-2.18.008-2.85.285c-.677.282-2.856.77-2.393 1.89m-12.53-2.243c.753.75 1.933-.415 1.17-1.173c-.252-.348-.645-.646-.948-.947c-.54-.54-2.162-2.8-3.068-1.89c-.79.792.585 1.756 1.082 2.25zm17.33 8.654c-.77-.317-1.535-.635-2.303-.952c-.646-.268-2.07-1.17-2.495-.135c-.48 1.168 2.054 1.747 2.75 2.035l1.368.565c1.014.42 1.693-1.094.68-1.513m.576-4.58h-2.49c-.7 0-2.357-.288-2.355.83c0 1.262 2.567.827 3.32.827h1.478c1.095 0 1.148-1.66.047-1.657m-5.708 5.75c-.673-.672-1.773.19-1.28 1.008a5.216 5.216 0 0 0-.962-.69c3.894-2.866 3.328-9.006-1.02-11.107a6.573 6.573 0 0 0-6.37.394a6.57 6.57 0 0 0-2.118 2.236a6.447 6.447 0 0 0-.784 2.235c-.083.534-.11.553-.553.87c-.182-.956-1.64-.674-2.326-.673c-.814 0-1.962-.217-2.75.046c-.868.29-.653 1.615.262 1.613c.324.05.7-.002 1.028-.002l2.713-.003c-.308.352-.496.97-.94.77a3.645 3.645 0 0 0-1.49-.32a3.7 3.7 0 0 0-2.56 1.037a3.702 3.702 0 0 0-1.124 2.516a5.27 5.27 0 0 0-2.422 1.658c-2.756 3.354-.265 8.554 4.058 8.554v-.002h10.792c1.34 0 2.843.167 4.168-.113c3.652-.772 5.36-5.21 3.133-8.23l1.643 1.642c.183.183.364.424.575.574c.552.553 1.524.067 1.403-.712c-.096-.622-1.04-1.267-1.447-1.673l-1.63-1.627m-6.096-9.78c4.56.008 6.576 5.978 2.912 8.733c-.638-3.504-4.162-5.823-7.63-5.03a4.869 4.869 0 0 1 4.717-3.703m4.758 15.293c-.633 3.346-4.15 2.88-6.68 2.88h-9.05c-.766 0-1.62.083-2.372-.096c-2.274-.538-3.416-3.242-2.172-5.235c.68-1.088 1.57-1.19 2.627-1.67c.604-.274.456-.808.456-1.332c.002-.597.284-1.17.756-1.533c.786-.608 1.942-.497 2.61.234c1.098 1.204 1.96-1.347 2.507-1.894c2.025-2.025 5.475-1.708 7.068.684c.344.516.58 1.102.693 1.712c.097.53-.115 1.34.188 1.796c.29.47.943.463 1.397.68a3.464 3.464 0 0 1 1.972 3.772M6.905 9.917l2.645 1.09c.353.147.707.293 1.06.438c.997.412 1.637-1.12.642-1.526c-.782-.48-1.796-.743-2.643-1.092L7.55 8.39c-.996-.41-1.638 1.115-.644 1.527"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Code(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.982 7.107L.322 15.77l8.66 8.662l3.15-3.15l-5.51-5.512l5.51-5.51zm12.675 0l-3.148 3.15l5.51 5.512l-5.51 5.51l3.147 3.15l8.662-8.662l-8.663-8.66z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Codetalk(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4.938c-7.732 0-14 4.7-14 10.5c0 1.98.74 3.833 2.016 5.414L2 25.272l5.613-1.44c2.34 1.316 5.237 2.106 8.387 2.106c7.732 0 14-4.7 14-10.5s-6.268-10.5-14-10.5M13.704 19.47l-2.338 2.336l-6.43-6.43l6.43-6.433l2.338 2.34l-4.09 4.092zm7.07 2.333l-2.336-2.34l4.092-4.09l-4.092-4.09l2.337-2.34l6.43 6.426l-6.43 6.433z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coffee(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.845 9.983L9.88 27.336c0 .977 2.74 1.77 6.12 1.77s6.12-.793 6.12-1.77L24.5 9.85c-2.455 1.024-6.812 1.134-8.498 1.134c-1.61 0-5.655-.1-8.155-1zm16.285-4.23l-.376-1.68c0-.65-3.472-1.178-7.754-1.178s-7.754.53-7.754 1.18L7.87 5.752c-.714.284-1.12.608-1.12.953V7.99c0 1.1 4.142 1.994 9.25 1.994s9.25-.894 9.25-1.995V6.704c0-.345-.406-.67-1.12-.953z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Commandline(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.02 9.748v-.002zm.002-.002l5.77 5.773l-5.77 5.77l2.12 2.123l7.895-7.895l-7.894-7.895zM12.248 23.27h14.42v-3h-14.42zm4.335-6.25h10.084v-3H16.583zm-4.335-9.25v3h14.42v-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Connect(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.06 13.72c-.944-5.173-5.46-9.095-10.903-9.095v4a7.104 7.104 0 0 1 7.094 7.094a7.104 7.104 0 0 1-7.093 7.092v4.002c5.442-.004 9.96-3.926 10.903-9.096h4.69v-4h-4.69zm-4.685 2a6.216 6.216 0 0 0-12.103-2.002H1.438v4h6.834a6.216 6.216 0 0 0 12.104-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Contract(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m25.083 18.895l-8.428-2.26l2.258 8.43l1.838-1.838l7.054 7.053l2.476-2.476l-7.053-7.053l1.856-1.855zM5.543 11.73l8.427 2.26l-2.258-8.43L9.874 7.4L3.196.72L.72 3.196l6.678 6.678l-1.856 1.857zm2.046 9.205l-6.87 6.87l2.475 2.475l6.87-6.87l1.857 1.858l2.258-8.428l-8.428 2.258l1.837 1.837zm15.822-10.87l6.867-6.87L27.802.717l-6.868 6.87L19.08 5.73l-2.26 8.43l8.43-2.26z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crop(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.303 21.707V8.275l4.48-4.42l-2.02-2.05l-4.127 4.07H8.76V2.084H5.883v3.793H1.8v2.877h4.083v15.832h15.542v4.61h2.878v-4.61H29.2v-2.878h-4.897zM19.72 8.753L8.76 19.565V8.753zm-9.032 12.953l10.735-10.59v10.59H10.69z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cross(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m24.778 21.42l-5.502-5.503l5.5-5.502l-2.827-2.83l-5.503 5.502l-5.502-5.502l-2.828 2.83l5.5 5.502l-5.5 5.502l2.83 2.828l5.5-5.502l5.5 5.502z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crown(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 8.087c1.007 0 1.826-.82 1.826-1.826c0-1.006-.82-1.825-1.826-1.825s-1.826.82-1.826 1.826c0 1.008.82 1.827 1.826 1.827zm0-3.044a1.217 1.217 0 1 1 0 2.437a1.218 1.218 0 0 1 0-2.437m-11.077 8.25c.812 2.61 1.132 6.645 1.257 9.41c.027 0 .053-.008.08-.008h19.48c.027 0 .054.008.08.01c.125-2.767.445-6.8 1.257-9.413c-.45-.23-.813-.59-1.05-1.038c-2.4 2.048-5.157 6.18-5.157 6.18s-3.026-6.058-4.305-9.812c-.182.045-.37.074-.565.074s-.383-.03-.565-.073c-1.278 3.754-4.305 9.812-4.305 9.812s-2.756-4.132-5.157-6.18c-.235.446-.6.808-1.05 1.037zm.73-2.163c0-1.007-.82-1.826-1.827-1.826S2 10.124 2 11.13s.82 1.826 1.826 1.826s1.826-.82 1.826-1.826zm-1.827 1.218a1.217 1.217 0 1 1 0-2.437a1.217 1.217 0 0 1 0 2.439zM25.74 23.913H6.26a1.826 1.826 0 0 0 .001 3.651h19.48a1.827 1.827 0 0 0-.001-3.653zm2.434-14.61c-1.007 0-1.826.82-1.826 1.827s.82 1.826 1.826 1.826S30 12.136 30 11.13s-.82-1.826-1.826-1.826zm0 3.045a1.217 1.217 0 1 1-.001-2.439a1.217 1.217 0 0 1 .001 2.441z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cube(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 3.03L4.7 9.263v12.47l10.8 6.237l10.8-6.234V9.266L15.5 3.028zm9.488 7.57L16 15.79v10.377c0 .275-.225.5-.5.5s-.5-.225-.5-.5v-10.38l-8.987-5.19a.499.499 0 1 1 .5-.865l8.988 5.19l8.99-5.19a.501.501 0 0 1 .683.184a.502.502 0 0 1-.185.683z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Customer(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28.523 23.813c-.518-.51-6.795-2.938-7.934-3.396c-1.133-.45-1.585-1.697-1.585-1.697s-.51.282-.51-.51c0-.793.51.51 1.02-2.548c0 0 1.415-.397 1.134-3.68h-.34s.85-3.51 0-4.698c-.853-1.188-1.187-1.98-3.06-2.548c-1.87-.567-1.19-.454-2.548-.396c-1.36.057-2.492.793-2.492 1.188c0 0-.85.057-1.188.397c-.34.34-.906 1.924-.906 2.32s.283 3.06.566 3.624l-.337.11c-.283 3.284 1.132 3.682 1.132 3.682c.51 3.058 1.02 1.755 1.02 2.548c0 .792-.51.51-.51.51s-.453 1.246-1.585 1.697c-1.132.453-7.416 2.887-7.927 3.396c-.51.52-.453 2.896-.453 2.896h12.036l.878-3.46l-.78-.78l1.343-1.345l1.343 1.344l-.78.78l.878 3.46h12.036s.063-2.378-.453-2.897z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Db(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 23.438c-3.847 0-7.71-.987-9.535-3.117c-.054.237-.09.48-.09.738v3.877c0 3.435 4.988 4.998 9.625 4.998s9.625-1.563 9.625-4.998v-3.877c0-.258-.036-.5-.09-.737c-1.826 2.13-5.688 3.118-9.536 3.118zm0-7.495c-3.847 0-7.71-.987-9.534-3.117a3.28 3.28 0 0 0-.09.736v3.877c0 3.434 4.987 4.997 9.624 4.997s9.625-1.563 9.625-4.998V13.56c0-.257-.036-.5-.09-.737c-1.826 2.13-5.688 3.118-9.536 3.118zm0-14.877c-4.637 0-9.625 1.565-9.625 5v3.877c0 3.435 4.988 4.998 9.625 4.998s9.625-1.562 9.625-4.997V6.067c0-3.435-4.988-5-9.625-5zm0 8c-4.21 0-7.625-1.343-7.625-3c0-1.656 3.414-3 7.625-3s7.625 1.344 7.625 3c0 1.658-3.414 3-7.625 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Detour(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m29.342 15.5l-7.556-4.363v2.614H18.75c-1.44-.003-2.423 1.003-2.875 1.785c-.735 1.222-1.056 2.56-1.44 3.522c-.136.36-.28.655-.377.817H11.29c-.213-.398-.57-1.557-.923-2.692c-.237-.676-.5-1.38-1.013-2.07c-.476-.682-1.464-1.386-2.604-1.362H2.812v3.5h3.772l.03.045c.29.4.633 1.663 1.03 2.888c.218.623.455 1.262.92 1.897c.417.614 1.32 1.293 2.383 1.293h3.427c.696.01 1.37-.286 1.81-.657c1.438-1.338 1.607-2.886 2.13-4.127c.217-.61.452-1.116.604-1.315l.018-.025h2.85v2.614zm-19.17-.96c.57.76.875 1.558 1.138 2.31l.125.4h2.58c.246-.698.553-1.48 1.005-2.23c.252-.437.62-.886 1.08-1.27H9.43c.305.253.56.527.743.79z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Diagram(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.812 17.202l7.396-3.665v-2.164h-.834c-.414 0-.808-.084-1.167-.237v1.16L4.81 15.96v2.912h2V17.2zm19.75 1.673v-2.913l-7.397-3.666v-1.158a2.98 2.98 0 0 1-1.166.236h-.833v2.164l7.395 3.666v1.672h2zm-9.874 0v-7.5h-2v7.5zm11.187 1H23.25a2 2 0 0 0-2 2V26.5a2 2 0 0 0 2 2h4.625a2 2 0 0 0 2-2v-4.625a2 2 0 0 0-2-2m-19.75 0H3.5a2 2 0 0 0-2 2V26.5a2 2 0 0 0 2 2h4.625a2 2 0 0 0 2-2v-4.625a2 2 0 0 0-2-2m5.25-9.5H18a2 2 0 0 0 2-2V3.75a2 2 0 0 0-2-2h-4.625a2 2 0 0 0-2 2v4.625a2 2 0 0 0 2 2m4.625 9.5h-4.625a2 2 0 0 0-2 2V26.5a2 2 0 0 0 2 2H18a2 2 0 0 0 2-2v-4.625a2 2 0 0 0-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Disconnect(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.22 9.5a6.219 6.219 0 0 0-5.885 4.218H1.438v4h1.897a6.216 6.216 0 0 0 12.102-2A6.217 6.217 0 0 0 9.218 9.5zm18.465 4.22c-.944-5.173-5.46-9.095-10.903-9.095v4a7.104 7.104 0 0 1 7.094 7.094a7.106 7.106 0 0 1-7.094 7.092v4.002c5.442-.004 9.96-3.926 10.903-9.096h2.065v-4h-2.065z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dockbottom(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.083 7.333v16.334h24.833V7.333zm21.832 9.5H6.083v-6.5h18.833z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dockleft(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.084 7.333v16.334h24.832V7.333zm8.583 3h13.25v10.335h-13.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dockright(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.083 7.333v16.334h24.833V7.333zm16.25 13.335H6.083V10.332h13.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Docktop(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27.916 23.667V7.333H3.083v16.334zm-3-3H6.082v-6.5h18.833v6.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dollar(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 1.466C7.973 1.466 1.466 7.973 1.466 16c0 8.027 6.507 14.534 14.534 14.534c8.027 0 14.534-6.507 14.534-14.534c0-8.027-6.507-14.534-14.534-14.534m1.255 22.414v2.047h-1.958v-2.024c-3.213-.44-4.62-3.08-4.62-3.08l2-1.673s1.277 2.223 3.587 2.223c1.276 0 2.244-.683 2.244-1.85c0-2.728-7.35-2.397-7.35-7.458c0-2.2 1.74-3.785 4.138-4.16V5.86h1.958v2.045c1.672.22 3.652 1.1 3.652 2.993v1.452H18.31v-.704c0-.726-.925-1.21-1.96-1.21c-1.32 0-2.287.66-2.287 1.584c0 2.794 7.35 2.112 7.35 7.415c0 2.18-1.628 4.07-4.158 4.445"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 1.466C7.973 1.466 1.466 7.973 1.466 16c0 8.027 6.507 14.534 14.534 14.534c8.027 0 14.534-6.507 14.534-14.534c0-8.027-6.507-14.534-14.534-14.534m0 27.326a2.806 2.806 0 1 1 0-5.611a2.806 2.806 0 0 1 0 5.611m0-7.705l-7.858-6.562h3.47V5.747h8.778v8.778h3.468z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dribbble(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 1.667C8.084 1.667 1.667 8.084 1.667 16S8.084 30.333 16 30.333S30.333 23.916 30.333 16S23.916 1.667 16 1.667m9.534 6.565a12.262 12.262 0 0 1 2.767 7.618c-3.23-.656-6.14-.783-8.776-.38a57.828 57.828 0 0 0-1.125-2.412c4.03-1.516 6.097-3.277 7.134-4.826m-1.58-1.608c-.876 1.658-3.19 3.232-6.563 4.476c-1.532-2.86-3.114-5.3-4.387-7.033c.96-.24 1.964-.37 2.997-.37c3.03 0 5.81 1.103 7.955 2.927zM10.847 4.83c1.09 1.442 2.75 3.91 4.415 6.968c-3.686 1.073-8.03 1.677-11.28 1.585a12.346 12.346 0 0 1 6.866-8.552zM3.696 16c0-.147.006-.293.01-.44c.313.014.633.02.96.02c3.465 0 7.755-.646 11.485-1.765l.14-.044c.355.705.698 1.432 1.034 2.176c-.516.146-1.022.314-1.518.507c-3.547 1.375-6.512 3.895-9.03 7.678a12.25 12.25 0 0 1-3.08-8.13zm4.692 9.656c2.31-3.574 5.002-5.924 8.21-7.167c.515-.2 1.048-.37 1.593-.514c1.206 2.996 2.167 6.205 2.56 9.373a12.271 12.271 0 0 1-4.75.952a12.22 12.22 0 0 1-7.612-2.646zm14.414.59c-.446-2.94-1.32-5.895-2.413-8.686c2.315-.28 4.88-.117 7.74.485a12.319 12.319 0 0 1-5.328 8.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dry(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.784 26.99c0 1.24-1.33 1.697-1.835 1.697c-.505 0-1.537-.413-1.65-1.812a.645.645 0 0 0-.645-.642a.64.64 0 0 0-.64.642c.044 1.056.755 3.052 2.934 3.052c2.43 0 3.165-1.882 3.165-3.144v-8.176l-1.328-.024l-.003 8.408zm.8-17.186c-6.807 0-11.084 4.86-11.587 8.326a3.516 3.516 0 0 1 2.89-1.51c1.197 0 2.22.582 2.855 1.495a3.519 3.519 0 0 1 2.88-1.495c1.2 0 2.26.6 2.896 1.517c.635-.917 1.83-1.517 3.03-1.517c1.19 0 2.24.59 2.88 1.495a3.436 3.436 0 0 1 2.854-1.495c1.197 0 2.254.597 2.89 1.51c-.503-3.467-4.78-8.326-11.588-8.326m-.85-2.68v2.082h1.322v-2.08a.661.661 0 0 0-1.323-.002zM7.56 6.015a1.124 1.124 0 0 0 1.69-1.013V3.09l-1.662.957a1.125 1.125 0 0 0-.028 1.968m-1.99 5.703c.11-.19.158-.398.15-.602V9.203l-1.66.957a1.126 1.126 0 0 0-.029 1.968a1.128 1.128 0 0 0 1.54-.41zm4.95-2.362a1.13 1.13 0 0 0 1.692-1.015V6.43l-1.662.956a1.12 1.12 0 0 0-.44.43a1.127 1.127 0 0 0 .41 1.54m4.73-5.357a1.128 1.128 0 0 0 1.69-1.015l-.002-1.91l-1.66.955a1.126 1.126 0 0 0-.029 1.97zm4.096 4.916a1.126 1.126 0 0 0 1.69-1.014V5.988l-1.662.957c-.178.096-.332.24-.44.43a1.126 1.126 0 0 0 .412 1.54m4.285-3.718a1.13 1.13 0 0 0 1.692-1.015h-.002V2.27l-1.662.957a1.13 1.13 0 0 0-.027 1.97zm3.895 3.314l-1.66.956a1.126 1.126 0 0 0-.029 1.97a1.127 1.127 0 0 0 1.69-1.015z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Edit(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27.87 7.863L23.024 4.82l-7.89 12.566l4.843 3.04zM14.395 21.25l-.107 2.855l2.527-1.337l2.35-1.24l-4.673-2.936zM29.163 3.24L26.63 1.647a1.364 1.364 0 0 0-1.88.43l-1 1.588l4.843 3.042l1-1.586c.4-.64.21-1.483-.43-1.883zm-3.965 23.82c0 .275-.225.5-.5.5h-19a.5.5 0 0 1-.5-.5v-19a.5.5 0 0 1 .5-.5h13.244l1.884-3H5.698c-1.93 0-3.5 1.57-3.5 3.5v19c0 1.93 1.57 3.5 3.5 3.5h19c1.93 0 3.5-1.57 3.5-3.5V11.097l-3 4.776v11.19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Employee(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28.523 23.813c-.518-.51-6.795-2.938-7.934-3.396c-1.133-.45-1.585-1.697-1.585-1.697s-.51.282-.51-.51c0-.793.51.51 1.02-2.548c0 0 1.415-.397 1.134-3.68h-.34s.85-3.51 0-4.698c-.853-1.188-1.187-1.98-3.06-2.548c-1.87-.567-1.19-.454-2.548-.396c-1.36.057-2.492.793-2.492 1.188c0 0-.85.057-1.188.397c-.34.34-.906 1.924-.906 2.32s.283 3.06.566 3.624l-.337.11c-.283 3.284 1.132 3.682 1.132 3.682c.51 3.058 1.02 1.755 1.02 2.548c0 .792-.51.51-.51.51s-.453 1.246-1.585 1.697c-1.132.453-7.416 2.887-7.927 3.396c-.51.52-.453 2.896-.453 2.896h26.954s.063-2.378-.453-2.897zm-6.335 2.25h-4.562v-1.25h4.562z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func End(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.167 5.5v8.18L6.684 5.32v20.364l14.483-8.364v8.18H25.5v-20z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ethernet(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.5 8.5v3.168l3.832 3.832l-3.832 3.832V22.5l7-7zm-14 14v-3.168L4.667 15.5L8.5 11.668V8.5l-7 7zm7-8.4a1.68 1.68 0 1 0 0 3.356a1.68 1.68 0 0 0 0-3.354zm-5.04 0a1.68 1.68 0 1 0 0 3.362a1.681 1.681 0 0 0 0-3.36zm10.08 0a1.68 1.68 0 1 0 0 3.357a1.68 1.68 0 0 0 0-3.355z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Exchange(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.786 12.876l7.556-4.363l-7.556-4.363v2.598H2.813v3.5h18.973zm-11.418 5.248l-7.556 4.362l7.556 4.362V24.25h18.974v-3.5H10.368z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Exclamation(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M26.71 14.086L16.915 4.29a2.007 2.007 0 0 0-2.83 0L4.29 14.086a2.007 2.007 0 0 0 0 2.83l9.796 9.795c.778.778 2.05.778 2.83 0l9.796-9.796a2.009 2.009 0 0 0 0-2.828zM14.703 8.98c.22-.237.5-.356.844-.356s.624.118.844.353c.22.235.33.53.33.885c0 .306-.1 1.333-.303 3.082c-.2 1.75-.38 3.44-.53 5.072h-.717c-.135-1.633-.3-3.323-.5-5.072c-.198-1.75-.298-2.776-.298-3.082c0-.35.11-.642.33-.88zm1.73 12.82c-.248.24-.543.36-.886.36s-.638-.12-.885-.36c-.247-.242-.37-.534-.37-.877s.123-.638.37-.885c.248-.248.543-.372.886-.372s.638.124.885.372c.25.247.373.542.373.885s-.124.635-.372.876z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Expand(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m25.545 23.328l-7.627-7.705l7.616-7.616l1.857 1.857l2.26-8.428l-8.428 2.258l1.836 1.836l-7.603 7.604l-7.513-7.59L9.81 3.695L1.392 1.394l2.215 8.44l1.848-1.83l7.524 7.604l-7.515 7.515l-1.856-1.855l-2.26 8.427l8.43-2.257L7.94 25.6l7.503-7.502l7.614 7.693l-1.867 1.848l8.416 2.3l-2.213-8.438z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Export(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.086 20.904c-1.805 3.113-5.163 5.212-9.023 5.22A10.45 10.45 0 0 1 4.625 15.688A10.45 10.45 0 0 1 15.063 5.25c3.86.007 7.216 2.105 9.022 5.218l3.962 2.284l.143.082C26.88 6.784 21.504 2.25 15.063 2.248C7.64 2.25 1.625 8.265 1.623 15.688c.003 7.42 6.018 13.435 13.44 13.437c6.442-.002 11.82-4.538 13.127-10.59l-.14.082zm4.314-5.216l-7.15-4.13v2.298H10.275v3.66H21.25v2.298z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Facebook(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.566 2.433H6.433c-2.2 0-4 1.8-4 4v19.135c0 2.2 1.8 4 4 4h19.135c2.2 0 4-1.8 4-4V6.433c-.002-2.2-1.8-4-4.002-4m-.257 14.483h-3.22v11.65h-4.818v-11.65h-2.41V12.9h2.41v-2.41c0-3.276 1.36-5.225 5.23-5.225h3.217V9.28h-2.012c-1.504 0-1.604.563-1.604 1.61l-.013 2.01h3.645l-.426 4.016z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fave(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.132 7.97c-2.203-2.204-5.916-2.097-8.25.236l-.382.382l-.382-.382c-2.334-2.333-6.047-2.44-8.25-.235c-2.204 2.204-2.098 5.917.235 8.25l8.396 8.396l8.395-8.396c2.334-2.333 2.44-6.046.237-8.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Feed(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.135 16.762c3.078 0 5.972 1.205 8.146 3.39a11.543 11.543 0 0 1 3.378 8.203h4.745c0-9.008-7.3-16.335-16.27-16.335zm.006-8.408c10.974 0 19.9 8.975 19.9 20.006h4.742c0-13.646-11.054-24.75-24.642-24.75zm6.56 16.69a3.286 3.286 0 1 1-6.572.002a3.286 3.286 0 0 1 6.572-.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ff(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.5 15.5L15.2 9.552v5.6l-9.7-5.6v11.895l9.7-5.6v5.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filter(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M26.834 6.958c0-2.094-4.852-3.79-10.834-3.79c-5.983 0-10.833 1.696-10.833 3.79c0 .43.213.84.588 1.224l8.662 15.002v4.9c0 .413.71.75 1.583.75c.875 0 1.584-.337 1.584-.75v-4.817L26.3 8.174h-.046c.37-.382.58-.79.58-1.216M16 9.75c-6.363 0-9.833-1.845-9.833-2.792s3.47-2.79 9.833-2.79c6.363 0 9.834 1.843 9.834 2.79S22.364 9.75 16 9.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Firefox(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28.4 22.47c.48-.965.85-1.992 1.095-3.067c.953-3.66.666-6.854.666-6.854l-.326 2.103s-.47-3.896-1.044-5.353c-.88-2.23-1.273-2.214-1.274-2.21c.542 1.38.494 2.17.483 2.288l-.028-.047c-.13-.323-.797-1.818-2.225-2.877a13.803 13.803 0 0 0-9.745-4.015A13.811 13.811 0 0 0 5.764 6.963c-.32-.463-.58-1.025-.605-1.646l-.007.005l-.003-.03S3.54 6.537 3.714 9.9c-.3.575-.56 1.173-.777 1.792c-.375.817-.75 2.004-1.06 3.746c0 0 .134-.422.4-.988c-.064.482-.103.97-.116 1.467c-.09.844-.117 1.864-.038 3.087c0 0 .032-.406.136-1.02C3.092 24.84 8.925 30.15 16 30.15c1.86 0 3.637-.37 5.257-1.035c3.68-1.344 5.86-3.92 7.143-6.646zM16.002 3.355c2.446 0 4.73.68 6.68 1.86c-2.274-.528-3.433-.26-3.423-.248c.012.015 3.383.59 3.98 1.41c0 0-1.43 0-2.856.41c-.065.02 5.242.664 6.327 5.967c0 0-.58-1.213-1.3-1.42c.473 1.44.35 4.17-.1 5.528c-.058.174-.118-.755-1.004-1.155c.284 2.037-.018 5.268-1.432 6.158c-.11.07.887-3.19.2-1.93c-4.092 6.276-8.958 2.54-10.933 1.208c1.586.388 3.268.108 4.243-.56c.982-.67 1.564-1.16 2.087-1.046c.522.117.87-.407.464-.872c-.405-.466-1.392-1.105-2.725-.757c-.94.246-2.108 1.286-3.887.232c-1.518-.9-1.507-1.63-1.507-2.095c0-.366.257-.88.734-1.028c.58.06 1.044.213 1.537.465a8.113 8.113 0 0 0 0-.52c.038-.076.014-.31-.048-.595a4.384 4.384 0 0 0-.19-.85c.01-.003.016-.008.02-.022c.076-.344 2.147-1.544 2.3-1.66c.152-.113.55-.377.505-1.182c-.015-.265-.058-.294-2.232-.286c-.917.003-1.425-.894-1.59-1.245c.223-1.23.864-2.11 1.92-2.704c.02-.01.015-.02-.008-.026c.22-.127-2.524-.006-3.76 1.604c-.33.047-.786-.048-1.295-.048c-.638 0-1.14.07-1.603.187a.62.62 0 0 1-.208 0a5.251 5.251 0 0 1-.535-.465A12.879 12.879 0 0 1 16 3.354z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fitocracy(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 10c1.93 0 3.5-1.57 3.5-3.5S17.93 3 16 3s-3.5 1.57-3.5 3.5S14.07 10 16 10m7.884-2.228s-2.067-1.923-2.52-.05c-.59 2.442-2.746 4.28-5.364 4.28s-4.774-1.838-5.364-4.28c-.452-1.873-2.52.05-2.52.05a12.034 12.034 0 0 0-4.2 9.145C3.917 23.59 9.327 29 16 29s12.083-5.41 12.083-12.083c0-3.658-1.63-6.93-4.2-9.145zm-8.924 8.502l.807-1.635l.807 1.634l1.805.263l-1.307 1.27l.31 1.798l-1.615-.854l-1.614.855l.31-1.797l-1.307-1.27l1.803-.264zm-7.122 5.638a9.502 9.502 0 0 1 0-9.99c1.908.862 3.246 2.765 3.246 4.995s-1.338 4.133-3.246 4.995M16 26.5a9.508 9.508 0 0 1-4.995-1.42c.862-1.91 2.765-3.247 4.995-3.247s4.133 1.338 4.995 3.246A9.508 9.508 0 0 1 16 26.499zm8.162-4.588c-1.908-.862-3.246-2.766-3.246-4.995c0-2.23 1.338-4.133 3.247-4.995a9.502 9.502 0 0 1 0 9.99z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FiveHundredPx(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.902 13.033c-2.57 0-4.65 2.998-4.65 2.998s-2.14-2.997-4.772-2.997c-2.63 0-3.793 2.69-3.793 2.69s-1.04-2.14-3.365-2.14c-2.095 0-2.35.842-2.38 1.01l.423-2.784h5.077V9.73H4.53l-1.393 6.864l2.432.05s.916-1.347 1.895-1.347s1.958.917 1.958 2.264c0 1.35-1.04 2.556-2.02 2.556s-1.957-1.023-1.957-1.76H3v.308c0 2.08 1.835 3.608 4.344 3.608c2.508 0 4.343-1.896 4.343-2.69c0 0 .795 2.69 3.916 2.69c3.122 0 4.588-2.69 4.588-2.69s2.143 2.69 4.712 2.69c2.568 0 4.1-2.143 4.1-4.77c-.002-2.634-1.898-4.47-4.1-4.47zm-9.39 6.545c-1.286 0-1.96-1.284-1.96-1.957s.43-1.957 1.897-1.957c1.468 0 2.997 1.958 2.997 1.958s-1.652 1.958-2.936 1.958zm9.267 0c-1.286 0-2.94-1.957-2.94-1.957s1.532-1.957 3-1.957c1.47 0 1.895 1.285 1.895 1.958s-.67 1.958-1.956 1.958z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.5 3v10c8 0 8 4 16 4V7c-8 0-8-4-16-4m-3 26h2V3h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagAlt(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.562 10.75C21.74 8.572 25.5 7 25.5 7c-8 0-8-4-16-4v10c8 0 8 4 16 4c0 0-3.75-3-5.938-6.25M6.5 29h2V3h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flickr(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.77 8.895a7.09 7.09 0 0 0-5.77 2.97a7.09 7.09 0 0 0-5.77-2.97A7.104 7.104 0 0 0 3.125 16a7.104 7.104 0 0 0 7.105 7.105c2.38 0 4.48-1.175 5.77-2.97a7.092 7.092 0 0 0 5.77 2.97a7.105 7.105 0 1 0 0-14.21m0 12.927a5.826 5.826 0 0 1-5.822-5.82a5.827 5.827 0 0 1 5.82-5.825a5.828 5.828 0 0 1 5.825 5.824a5.828 5.828 0 0 1-5.824 5.822z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28.625 26.75h-26.5V8.375H3.25c1.75 0 .747-3.125 3-3.125h5.125c2.25 0 1.25 3.125 3 3.125h14.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Font(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m22.255 19.327l-1.017.13c-.61.082-1.067.21-1.375.383c-.52.293-.78.76-.78 1.398c0 .484.18.867.533 1.146c.354.28.774.42 1.262.42c.593 0 1.164-.137 1.72-.41c.938-.454 1.4-1.19 1.4-2.23V18.81c-.205.13-.47.23-.792.328c-.323.09-.642.152-.95.19zm-14.22-1.054h4.31l-2.113-6.063zM28.168 7.75h-25c-.55 0-1 .448-1 1v16.583c0 .553.45 1 1 1h25c.554 0 1-.447 1-1V8.75a1 1 0 0 0-1-1M14.305 23.896l-1.433-4.11H7.488L6 23.897H4.094L9.262 10.17h2.1l4.98 13.727zm12.487.047c-.263.074-.46.12-.6.14a4.11 4.11 0 0 1-.56.028c-.58 0-1-.203-1.262-.614c-.138-.22-.23-.525-.29-.926c-.344.45-.834.84-1.477 1.17a4.619 4.619 0 0 1-2.12.492c-.93 0-1.69-.28-2.274-.844c-.59-.562-.885-1.27-.885-2.113c0-.928.29-1.646.868-2.155c.578-.51 1.34-.824 2.28-.942l2.68-.336c.39-.05.647-.21.776-.484c.062-.146.103-.354.103-.646c0-.575-.203-.993-.604-1.252c-.41-.26-.99-.39-1.75-.39c-.876 0-1.5.24-1.864.714c-.205.263-.34.654-.4 1.174H17.85c.03-1.237.438-2.097 1.2-2.582c.77-.484 1.658-.726 2.673-.726c1.176 0 2.13.225 2.864.673c.73.448 1.093 1.146 1.093 2.093v5.766c0 .176.035.313.106.422c.07.104.223.156.452.156c.076 0 .16-.005.254-.015c.093-.01.19-.02.3-.04z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fork(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.74 10.25h8.046v2.626l7.556-4.363l-7.556-4.363v2.598H9.826c1.543.864 2.79 2.174 3.915 3.5zm8.046 10.404c-.618-.195-1.407-.703-2.29-1.587c-1.79-1.756-3.713-4.675-5.732-7.227c-2.05-2.486-4.16-4.972-7.45-5.09h-3.5v3.5h3.5c.655-.028 1.682.485 2.878 1.682c1.788 1.753 3.712 4.674 5.73 7.226c1.922 2.33 3.908 4.64 6.864 5.016v2.702l7.556-4.362l-7.556-4.362z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForkAlt(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.786 12.873l7.556-4.36l-7.556-4.363v2.7c-2.93.375-4.905 2.64-6.81 4.953c.546.703 1.08 1.418 1.605 2.127l.574.77c.802-1.043 1.584-2 2.34-2.74c.885-.885 1.674-1.393 2.292-1.588v2.5zm-4.125 4.133c-1.892-2.37-3.814-5.354-6.008-7.537c-1.46-1.43-3.155-2.665-5.34-2.694h-3.5v3.5h3.5c.97-.12 2.845 1.333 4.712 3.77c1.895 2.372 3.815 5.355 6.01 7.538c1.327 1.297 2.848 2.426 4.752 2.645v2.645l7.556-4.363l-7.556-4.362v2.535c-1.04-.34-2.58-1.663-4.125-3.68z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FullBattery(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27.27 13.5h-1.103v-1.416a1 1 0 0 0-1-1H5.25a1 1 0 0 0-1 1v7.832a1 1 0 0 0 1 1h19.917a1 1 0 0 0 1-1V18.5h1.104c.266 0 .48-.448.48-1v-3c0-.552-.214-1-.48-1zm-3.103 5.416H6.25v-5.832h17.917zm-1-4.832H7.25v3.832h15.917z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fullcube(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 3.03L4.7 9.263v12.47l10.8 6.237l10.8-6.234V9.266L15.5 3.028zm0 4l6.327 3.65l-6.327 3.654l-6.326-3.652zm9.488 3.57L16 15.79v10.377c0 .275-.225.5-.5.5s-.5-.225-.5-.5v-10.38l-8.987-5.19a.499.499 0 1 1 .5-.865l8.988 5.19l8.99-5.19a.501.501 0 0 1 .683.184a.502.502 0 0 1-.185.683z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Future(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m17 15.5l-.5-7.876a1 1 0 1 0-2 0l-.465 7.343l-3.004 1.96a1 1 0 0 0 1.001 1.732l3.306-1.678c.054.008.108.018.165.018c.83 0 1.5-.67 1.5-1.5zm1.94 5.955a.999.999 0 1 0 1.363.367a1 1 0 0 0-1.365-.367zm.998-13.642a1 1 0 1 0-.997 1.735a1 1 0 0 0 .997-1.736zm1.885 12.492a1 1 0 1 0 .998-1.734a1 1 0 0 0-.998 1.734m1-7.877a1 1 0 1 0-1.001-1.733a1 1 0 0 0 1 1.733zm1.555 3.072c0-.55-.448-1-1-1a1 1 0 1 0 1 1M9.546 12.062a1 1 0 1 0-1.73-1a1 1 0 0 0 1.729 1zM6.624 15.5a1.001 1.001 0 1 0 2.002-.002a1.001 1.001 0 0 0-2.002.002m2.555 4.805a1.001 1.001 0 1 0-.002 0zm2.882-10.76a.999.999 0 0 0 .366-1.366a.999.999 0 1 0-.366 1.364zm2.44 13.832a1 1 0 1 0 1.998.002a1 1 0 0 0-1.998-.002m-3.806-1.555c-.275.48-.11 1.09.366 1.365a1 1 0 1 0-.366-1.365m17.978-7.735l-3.27-1.186c.29 1.106.41 2.275.31 3.48c-.493 5.638-5.45 9.808-11.092 9.332a10.26 10.26 0 0 1-9.333-11.09c.493-5.64 5.448-9.812 11.088-9.335a10.169 10.169 0 0 1 6.194 2.833l-1.637 1.377l7.03 2.548l-1.308-7.364l-1.77 1.493c-2.134-2.15-4.997-3.597-8.25-3.877c-7.29-.626-13.71 4.776-14.34 12.068c-.625 7.29 4.777 13.71 12.066 14.34c7.293.624 13.713-4.777 14.342-12.067c.074-.866.057-1.718-.03-2.55z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gear(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m26.974 16.514l3.765-1.99a14.746 14.746 0 0 0-.397-2.158l-4.182-.58c-.36-.87-.84-1.68-1.4-2.422l1.593-3.92a14.652 14.652 0 0 0-1.686-1.407l-3.55 2.23c-.834-.44-1.73-.77-2.673-.985L17.16 1.306c-.363-.027-.727-.056-1.098-.056s-.734.028-1.1.056l-1.27 3.94c-.967.208-1.884.544-2.738.987L7.458 4.037c-.595.43-1.16.895-1.685 1.406l1.55 3.812c-.604.775-1.11 1.63-1.49 2.55l-4.05.56a14.623 14.623 0 0 0-.395 2.157l3.635 1.923c.04 1.013.21 1.994.506 2.918l-2.743 3.032c.32.66.674 1.303 1.085 1.905l4.037-.867c.66.72 1.415 1.35 2.247 1.873l-.153 4.13c.663.3 1.352.55 2.062.75l2.554-3.283c.453.058.912.097 1.38.097c.507 0 1.003-.046 1.49-.113l2.568 3.3c.71-.2 1.4-.45 2.062-.748l-.156-4.206a11.036 11.036 0 0 0 2.146-1.82l4.142.888a14.75 14.75 0 0 0 1.085-1.905l-2.83-3.13c.27-.876.423-1.8.467-2.753zm-6.257 4.783l-1.785 1.162l-1.098-1.688a5.08 5.08 0 0 1-1.834.353a5.126 5.126 0 1 1 3.625-1.504z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Github(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28.436 15.1c-1.2-.203-2.45-.336-3.466-.372l-.18-.006c.042-.09.073-.15.083-.16c.022-.018.04-.094.042-.168c0-.04.018-.174.046-.35c.276.01.64.018 1.04.02c1.536.013 3.144.137 4.247.332c.657.116.874.112.39-.006c-.492-.12-1.948-.294-3.108-.37a35.645 35.645 0 0 0-2.555-.062c.02-.114.04-.24.064-.37c.092-.504.123-1.01.125-2.017c.002-1.56-.082-1.99-.59-3.024a5.011 5.011 0 0 0-.725-1.104c.247-.73.24-1.858-.015-2.848c-.21-.812-.285-.864-1.02-.708c-.625.133-1.125.314-1.766.637c-.304.153-.722.39-1.025.578c-.79-.277-1.607-.46-2.48-.56c-.883-.1-3.05-.044-3.82.098c-.75.14-1.428.31-2.04.51c-.307-.188-.75-.443-1.068-.603c-.644-.324-1.144-.504-1.77-.637c-.734-.157-.81-.104-1.02.708c-.26 1.003-.262 2.15-.005 2.878c.03.083.048.142.055.188c-1.042 1.312-1.382 2.78-1.156 4.83c.06.533.15 1.023.278 1.472c-.665-.004-1.61.02-2.294.064c-1.162.077-2.618.25-3.11.37c-.483.117-.268.12.39.006c1.103-.194 2.712-.32 4.248-.33c.29-.002.56-.008.794-.014c.07.237.15.463.24.678H7.26c-1.015.036-2.264.17-3.465.37c-.9.152-2.23.454-2.386.54c-.164.092-.03.072.667-.105c1.273-.322 2.928-.57 4.978-.74l.23-.02c.44 1.02 1.117 1.8 2.075 2.41c.586.372 1.525.755 1.998.815c.13.016.508.094.84.172c.333.078.984.195 1.446.262h.01l-.024.016c-.56.29-.924.744-1.17 1.457c-.11.033-.246.078-.394.13c-.53.18-.735.216-1.27.22c-.557.004-.69-.02-1.02-.176a2.968 2.968 0 0 1-1.234-1.134c-.5-.826-1.367-1.41-2.09-1.41c-.616 0-.733.25-.287.615c.672.55 1.174 1.11 1.38 1.538l.397.824c.11.227.342.535.564.748c.522.498 1.026.736 1.778.848c.504.074.628.074 1.223-.002c.287-.035.53-.076.746-.127v.854c0 1.766-.02 2.334-.09 2.5c-.133.316-.43.64-.717.787c-.287.147-.376.308-.255.456c.068.08.197.094.63.066c.822-.05 1.403-.355 1.7-.89c.094-.173.116-.52.146-2.32c.032-1.952.046-2.14.173-2.42c.076-.165.187-.345.25-.394c.103-.086.11.084.11 2.42c0 2.578-.027 2.89-.285 3.385c-.058.113-.168.26-.245.33c-.135.123-.192.438-.098.533c.155.154.932-.088 1.356-.422c.722-.572.808-1.045.814-4.46l.003-2.005l.22.02l.218.02l.036 2.62c.04 2.952.046 2.995.548 3.565c.285.322.572.5 1.04.64c.624.187.812-.103.392-.606c-.457-.548-.48-.757-.454-3.995c.017-2.076.017-2.076.15-1.955c.283.256.337.676.337 2.623c0 2.418.07 2.648.923 3.07c.4.195.51.22 1.022.22c.544.003.577-.005.597-.147c.017-.115-.05-.193-.304-.348a1.741 1.741 0 0 1-.708-.797c-.055-.126-.092-.958-.117-2.67c-.036-2.394-.044-2.503-.193-2.878c-.2-.504-.508-.902-.897-1.166c-.1-.066-.202-.12-.333-.162c.16-.016.317-.033.468-.055c1.572-.21 2.403-.384 3.07-.642c1.41-.543 2.365-1.445 2.882-2.724c.046-.114.092-.222.13-.31l.4.034c2.05.173 3.705.42 4.978.743c.698.178.83.2.668.106c-.152-.085-1.482-.386-2.382-.537zm-6.014-.032c-.233.512-.883 1.17-1.408 1.428c-.518.256-1.33.45-2.25.544c-.63.064-4.137.083-4.716.026c-1.917-.188-2.99-.557-3.783-1.296c-.75-.702-1.1-1.655-1.04-2.828c.04-.734.217-1.195.68-1.755c.42-.51.864-.825 1.386-.985c.438-.134 1.78-.146 3.582-.03c.797.05 1.456.05 2.252 0c1.886-.12 3.145-.106 3.61.038c.73.226 1.397.834 1.797 1.644c.18.362.216.516.242 1.075c.036.77-.097 1.587-.35 2.138zm-9.51-3.306c-1.073-.188-1.686 1.65-.863 2.587c.39.444.737.517 1.17.247c.403-.25.62-.72.62-1.328c0-.812-.368-1.408-.928-1.508zm6.513.11c-1.073-.188-1.687 1.647-.864 2.586c.393.445.74.52 1.174.247c.4-.25.62-.72.62-1.328c0-.808-.37-1.406-.93-1.505m-2.886 3.612c-.024.074-.136.184-.25.243c-.285.147-.49.096-.793-.18c-.187-.168-.272-.257-.33-.08c-.052.164.28.493.538.594c.236.095.405.098.66-.01c.255-.105.477-.39.477-.575c0-.172-.247-.164-.303.01zm-.318-.575c.163-.145.2-.44.044-.598s-.473-.133-.597.043c-.145.206-.068.363.035.53c.16.124.375.15.517.024z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Githubalt(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m23.356 17.485l-.01.02l.162.006c.107.005.218.01.33.017c-.046-.004-.09-.01-.136-.013zM15.5 1.25C7.63 1.25 1.25 7.63 1.25 15.5S7.63 29.75 15.5 29.75s14.25-6.38 14.25-14.25S23.37 1.25 15.5 1.25M3.77 17.092a28.99 28.99 0 0 1 2.792-.156c.262 0 .507-.006.717-.012c.062.213.135.42.218.613h-.006a24.658 24.658 0 0 0-3.574.415c-.06-.283-.108-.57-.147-.86m8.426 5.13c-.1.03-.224.07-.357.118c-.48.17-.666.207-1.15.207c-.503.015-.622-.02-.922-.17a2.694 2.694 0 0 1-1.117-1.024c-.452-.746-1.235-1.275-1.888-1.275c-.56 0-.664.227-.26.557c.607.496 1.06.998 1.247 1.385c.104.215.265.546.357.744c.1.206.31.474.51.676c.473.44.93.66 1.61.772c.454.06.566.06 1.104-.004c.26-.03.48-.067.676-.118v.77c0 1.05-.008 1.63-.03 1.946a11.847 11.847 0 0 1-8.044-8.782a35.03 35.03 0 0 1 3.373-.43l.208-.018c.398.925 1.01 1.63 1.876 2.18c.53.336 1.38.684 1.807.732c.118.02.46.09.76.16c.302.066.89.172 1.31.236h.008c-.007.018-.014.02-.022.02c-.506.265-.835.675-1.057 1.32zm1.536 4.984a13.363 13.363 0 0 1-.5-.087c.024-.286.038-.785.054-1.723c.028-1.767.04-1.94.156-2.19c.07-.15.17-.32.226-.356c.095-.077.1.077.1 2.19c0 1.103-.005 1.746-.036 2.167zm1.768.132c-.148 0-.296-.007-.443-.013c.086-.562.104-1.428.106-2.87l.003-1.82l.197.018l.2.02l.03 2.365c.018 1.21.028 1.878.076 2.296c-.057 0-.112.003-.17.003zm1.506-.1c-.04-.485-.037-1.243-.027-2.553c.018-1.866.018-1.866.13-1.77c.246.247.305.624.305 2.374c0 .93.01 1.498.082 1.877c-.163.03-.327.053-.49.073zm10.083-9.313a25.549 25.549 0 0 0-3.251-.4c1.25.108 2.326.248 3.245.418a11.807 11.807 0 0 1-3.214 5.928c-1.4 1.4-3.15 2.448-5.105 3.008c-.034-.334-.058-1.047-.066-2.21c-.03-2.168-.04-2.264-.17-2.603a2.346 2.346 0 0 0-.812-1.055a1.259 1.259 0 0 0-.3-.14c.144-.02.28-.02.426-.057c1.418-.188 2.168-.357 2.772-.584c1.263-.49 2.13-1.3 2.606-2.467c.044-.103.088-.2.123-.28l.01.002a.96.96 0 0 1 .065-.125c.02-.017.036-.085.038-.15c0-.038.017-.158.04-.318c.25.01.58.018.94.02c.958.008 1.944.064 2.793.156c-.036.288-.082.576-.14.86zm-1.267-1.057a32.65 32.65 0 0 0-2.31-.057c.02-.103.036-.218.058-.336c.084-.454.112-.912.114-1.823c.002-1.413-.074-1.8-.534-2.735a4.586 4.586 0 0 0-.655-1c.225-.658.207-1.68-.02-2.574c-.19-.734-.258-.78-.924-.64a6.182 6.182 0 0 0-1.597.576c-.274.138-.652.354-.923.522c-.715-.25-1.45-.42-2.242-.508c-.8-.092-2.76-.04-3.454.09c-.68.125-1.293.28-1.848.46c-.276-.17-.678-.4-.964-.546a6.055 6.055 0 0 0-1.597-.573c-.664-.144-.732-.095-.922.64c-.235.907-.237 1.945-.004 2.603c.025.075.042.13.05.17c-.943 1.187-1.25 2.515-1.047 4.367c.053.482.136.926.25 1.333c-.6-.004-1.456.018-2.073.057c-.454.03-.957.076-1.418.13a11.794 11.794 0 0 1 3.367-9.897a11.791 11.791 0 0 1 8.37-3.467c3.273 0 6.226 1.323 8.37 3.467a11.794 11.794 0 0 1 3.37 9.872c-.46-.056-.964-.103-1.417-.132z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Glasses(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.075 9.53S11.37 8.094 8.917 8.094c-2.453 0-4.862.593-4.862.593L3.97 9.87l.53.53c.337.335.485 3.74 1.837 5.093c1.353 1.354 4.82 1.396 5.963.676c1.14-.72 2.24-3.467 2.24-4.694V10.8c.275-.275 1.616-.303 1.918 0v.676c0 1.227 1.1 3.975 2.24 4.693c1.145.72 4.612.677 5.964-.677c1.355-1.353 1.5-4.757 1.84-5.094l.527-.53l-.085-1.184s-2.408-.593-4.862-.593c-2.453 0-5.158 1.438-5.158 1.438c-.606-.238-2.188-.21-2.85 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globe(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 1.466C7.973 1.466 1.466 7.973 1.466 16c0 8.027 6.507 14.534 14.534 14.534c8.027 0 14.534-6.507 14.534-14.534c0-8.027-6.507-14.534-14.534-14.534m3.158 21.803c-.08.063-.183.13-.105.206c.078.078-.09.13-.09.17s.104.246.052.336c-.052.092-.09.223-.13.3c-.04.078-.13.156-.104.273c.026.116-.103.077-.103.194c0 .115.116.064.09.207c-.025.144-.09.183-.09.285c0 .104.064.247.064.286s-.063.17-.154.272c-.092.104-.155.17-.144.233c.014.065.104.144.09.184c-.012.037-.128.168-.115.26c.014.09.13.052.155.115c.026.065-.155.118-.078.183c.078.064.183.05.156.208c-.02.112.064.163.126.198c-.89.222-1.818.353-2.777.353C9.64 27.533 4.465 22.36 4.465 16c0-2.073.557-4.015 1.518-5.697c.08-.042.137-.07.17-.062c.066.014.08.105.184.13c.104.027.195-.077.26-.116c.064-.04.116-.195.05-.182c-.064.013-.233 0-.233 0s.183-.104.183-.17s.025-.168.13-.207c.102-.04.102-.014.037.142s.13.09.208.026s.13-.013.272-.104c.143-.092.143-.144.22-.144c.08 0 .222.143.3.09c.077-.05.3.066.43.066c.128 0 .544.17.623.17c.078 0 .312.09.325.258c.013.17.09.156.168.156s.26.065.26.13s-.05.325.08.39c.128.064.246.17.298.143c.052-.027 0-.234-.064-.26c-.065-.027-.027-.118-.052-.17c-.026-.05.078-.05.117.04c.038.09.142.26.207.26c.064 0 .208.155.168.246c-.04.09.04.22.156.22c.117 0 .26.183.313.196c.052.013.117.078.117.117c0 .04.065.26.065.35c0 .09-.04.454-.053.597s.104.39.234.52s.246.377.324.43c.08.05.13.194.247.18c.117-.012.195.08.3.26c.103.183.207.48.285.507c.078.026.208.117.142.182c-.064.064-.168.208-.05.208c.116 0 .155-.065.246.053c.09.116.208.18.194.26c-.013.077.104.103.156.116c.052.013.17.247.286.143c.117-.104-.155-.26-.234-.326c-.078-.064 0-.207-.182-.35s-.156-.247-.286-.35c-.13-.105-.233-.196-.104-.287c.13-.09.143.09.195.208c.052.117.324.352.44.455c.118.104.327.468.39.468s.248.208.248.208s.103.17.064.22c-.04.053.053.248.144.3c.09.052.456.22.508.247c.052.027.155.22.3.22c.14 0 .246.015.285.054c.04.04.155.195.234.105c.078-.092.09-.13.208-.13c.117 0 .168.09.233.155s.247.235.338.222c.09-.013.208.104.273.064s.168.025.22.052c.053.026.233.118.22.272c-.012.157.104.195.183.234c.078.04.182.13.248.195c.065.062.207.076.247.115c.04.04.065.117.182.052c.115-.064.09-.18.09-.18s.13-.027.195.025c.064.05.104.22.144.246c.037.026.114.22.062.362c-.05.145-.038.286-.09.286c-.053 0-.117.17-.196.21c-.076.038-.285.22-.272.285c.013.064.13.26.104.35c-.025.092-.194.196-.154.34c.038.143.312.182.323.31c.014.132.21.418.235.547c.025.13.246.272.246.453c0 .184.312.3.377.312c.063.013.182.13.272.17s.17.116.233.22s.053.262.053.3c0 .04-.04.44-.078.674c-.04.235.05.417-.027.483zm-8.392-12.082c.04.013.117.09.156.09c.04 0 .234.157.286.21c.053.05.053.194-.013.207s-.104-.143-.117-.208c-.013-.065-.143-.065-.208-.104c-.065-.04-.21-.23-.104-.195M27.51 16.41c-.144.182-.13.272-.195.286c-.064.013.065.065.09.194c.022.112-.065.224.063.327c-.486 4.62-3.71 8.434-8.016 9.787c-.007-.01-.02-.025-.02-.034c-.028-.078-.028-.233.063-.285c.09-.053.312-.233.363-.272c.052-.04.13-.22.09-.247c-.037-.026-.23 0-.26-.04c-.025-.038-.025-.09.105-.18c.13-.092.195-.223.247-.26c.052-.04.155-.118.195-.21c.038-.09-.04-.04-.118-.04S20 25.297 20 25.23s.195.026.34.052c.142.024.076-.065.063-.142c-.013-.078.026-.21.105-.17c.076.04.48-.013.53-.026c.053-.013.195-.246.247-.312c.053-.065.064-.13 0-.168c-.065-.04-.143-.184-.168-.22a.725.725 0 0 1-.013-.34c.025-.064 0 .376.18.43c.184.05.287.077.456-.08c.17-.154.298-.26.312-.362c.013-.104.052-.21.117-.246c.064-.04.103.103.18-.065c.08-.17.157-.156.235-.298c.077-.144-.13-.325.024-.43c.156-.103.43-.233.43-.233s.077-.04.233-.08c.155-.037.324-.013.376-.09c.05-.078.103-.246.18-.337c.08-.092.17-.234.13-.3c-.038-.064.105-.35.092-.428c-.013-.078-.13-.26.065-.416s.402-.392.416-.455c.012-.065.17-.338.154-.47c-.012-.128-.154-.284-.245-.324c-.092-.036-.286-.05-.364-.153s-.3-.208-.377-.182c-.076.026-.207.05-.31-.015c-.105-.063-.273-.143-.338-.194c-.066-.05-.234-.09-.312-.09s-.065-.053-.182.103c-.117.156 0 .208-.208.18c-.21-.023.025-.037.144-.193c.115-.155-.014-.247-.144-.207c-.13.04-.04.117-.247.156c-.207.038-.207-.092-.077-.117c.13-.027.363-.144.363-.195c0-.053-.027-.196-.13-.196s-.08-.13-.234-.298c-.156-.17-.35-.274-.508-.25c-.154.027-.272.066-.35-.075c-.078-.144-.17-.17-.222-.247c-.05-.078-.182 0-.22-.04s-.04-.038-.04-.038s-.17.04-.077-.078c.09-.117.128-.338.09-.325c-.04.012-.105.195-.17.182c-.063-.013-.013-.04-.143-.117c-.13-.078-.337-.013-.337.052s-.065.117-.065.117s-.04-.038-.078-.117c-.04-.078-.22-.09-.312-.013c-.09.078-.142-.196-.207-.196s-.194.065-.26.184c-.064.116-.038.285-.092.272c-.05-.013-.063-.233-.05-.312c.012-.08.155-.208.05-.234c-.103-.026-.26.13-.323.143a.616.616 0 0 0-.273.21c-.077.103-.116.168-.195.207c-.077.04-.193 0-.167-.04c.025-.038-.222-.18-.26-.13c-.04.053-.156.092-.273.145c-.117.052-.222-.065-.247-.117s-.08-.064-.09-.234c-.014-.167.026-.35.064-.453c.038-.104-.195-.312-.286-.3c-.09.015-.182.105-.272.09c-.092-.01-.052-.037-.195-.037c-.144 0-.027-.025 0-.143c.024-.116-.053-.273.09-.377c.143-.104.092-.35 0-.363c-.09-.014-.26.04-.376.026a.156.156 0 0 0-.17.207c.04.117-.064.195-.103.183c-.038-.013-.09-.078-.233.026c-.142.102-.194.063-.337-.053c-.143-.118-.3-.234-.325-.416c-.026-.18-.04-.363.013-.467c.05-.104.05-.285-.026-.312c-.078-.024.09-.154.18-.18a.579.579 0 0 0 .26-.195c.027-.052.157-.04.3-.04c.142 0 .168 0 .31.078c.144.078.17-.04.17-.078c0-.04.052-.117.208-.104c.156.014.376-.05.416-.012s.116.195.194.143c.08-.05.104-.142.13.015c.026.155.09.39.21.43c.115.038.05.193.167.206c.115.014.17-.245.13-.336c-.04-.09-.117-.363-.182-.428c-.064-.065-.064-.234.064-.286c.13-.052.44-.312.53-.39c.093-.078.34-.143.262-.247c-.078-.104-.104-.168-.104-.247s.078-.052.117 0s.194-.078.155-.143c-.038-.064-.026-.155.065-.143c.09.013.116-.065.078-.117c-.04-.052.09-.117.182-.09c.092.025.325-.014.364-.066c.037-.052-.08-.104-.08-.208c0-.105.156-.196.248-.21c.09-.012.207 0 .22-.038c.013-.04.144-.143.156-.052c.014.09 0 .247.104.247s.232-.117.272-.13c.038-.012.286-.064.338-.077c.052-.013.363-.04.325-.13c-.04-.09-.078-.18-.118-.22c-.04-.04-.077.013-.13.078c-.05.065-.143.065-.168.013c-.026-.05.012-.206-.078-.155c-.093.052-.105.104-.158.078c-.052-.026-.103-.117-.103-.117s.128-.064.037-.182c-.09-.117-.22-.09-.35-.025c-.13.063-.118.05-.273.09s-.234.077-.234.077s.21-.13.3-.208c.09-.08.208-.118.285-.196c.078-.078.285.04.285.04s.105-.105.105-.04s-.027.234.05.234c.08 0 .3-.104.21-.13c-.092-.027.13 0 .22-.066c.092-.065.194-.065.247-.09c.052-.026.092-.143.182-.143c.092 0 .13.116 0 .194s-.143.273-.208.325c-.064.05-.026.116.078.103c.104-.013.194.013.286-.013s.143.026.168.065c.026.04.104-.04.104-.04s.17-.038.22.027c.054.064.093-.04.054-.104c-.04-.064-.092-.13-.13-.208c-.04-.078-.09-.104-.194-.078c-.104.026-.13-.026-.195-.064c-.065-.04-.118.052-.065-.04c.053-.09.078-.117.117-.195c.038-.078.208-.22.038-.26c-.17-.04-.222-.064-.247-.142c-.025-.077-.22-.22-.27-.22c-.054 0-.234 0-.248-.065c-.013-.065-.143-.208-.208-.273c-.064-.065-.312-.35-.35-.377c-.04-.026-.092-.013-.21.143c-.115.157-.22.183-.31.144c-.092-.04-.105-.026-.194-.13c-.093-.104.09-.117.05-.182c-.04-.064-.246-.09-.376-.104c-.13-.013-.22-.156-.416-.17c-.194-.012-.428.027-.493.027c-.064 0-.064.09-.09.234c-.027.143.09.182-.027.208c-.116.026-.17.04-.052.09c.117.053.273.26.273.26s0 .118-.093.183c-.09.065-.182.13-.233.053c-.053-.078-.195-.064-.155.014c.037.078.115.117.115.195c0 .076.117.27.04.336c-.08.065-.17.014-.234.026s-.13-.104-.077-.13c.05-.026-.014-.22-.014-.22s-.156.22-.144.103c.014-.117-.064-.13-.064-.22c0-.092-.08-.13-.194-.105c-.118.026-.26-.04-.482-.08a1.433 1.433 0 0 1-.493-.155c-.182-.09-.247-.026-.338-.013c-.09.013-.052-.182-.17-.207c-.115-.027-.18.025-.206-.144c-.027-.167.038-.207.323-.39c.286-.18.247-.26.468-.285c.22-.026.326.026.326-.04s.052-.324.052-.194s.262.14.144.23c-.116.09-.052.104.04.104c.09 0 .26-.09.26-.09s.207-.092.26-.014c.052.078.052.156.143.156s.285-.104.116-.195c-.167-.09-.27-.077-.375-.18s-.078-.066-.195-.04c-.116.026-.116-.04-.156-.04s-.104.027-.13-.025c-.025-.052.014-.065.145-.065c.128 0 .284.04.284.04l.194-.066c.04-.013.247-.04.208-.155c-.04-.118-.17-.118-.208-.157s.078-.09.143-.117c.066-.026.248 0 .248 0s.117.013.117-.04s-.028-.116.05-.077s0 .155.118.13c.116-.027.143 0 .207.038c.065.04-.013.195-.077.22c-.065.026-.17.078-.026.09c.144.015.246.015.246.015s.092-.09.13-.17c.04-.077.105-.025.156 0c.05.026.246.066.064.118c-.183.053-.22.118-.26.183c-.038.065-.053.104-.22.065c-.17-.04-.26-.026-.3.04c-.04.063-.013.272.053.246c.063-.026.13-.026.207-.052c.078-.026.39.026.467.013c.08-.013.21.13.25.104c.038-.027.116.05.193.103c.078.052.052-.117.194-.013c.144.104.065.104.144.104c.077 0 .247.013.247.013s.014-.13.144-.104c.13.026.246.17.233.064c-.012-.103.013-.18-.09-.26c-.104-.077-.272-.13-.3-.168c-.025-.04-.05-.092-.012-.118c.04-.025.222.013.325.08c.104.064.195.13.273.077c.077-.053.17-.08.208-.118c.038-.04.13-.156.13-.156s-.39-.05-.44-.117c-.054-.066-.236-.157-.288-.157s-.194.09-.246-.04s-.052-.285-.105-.298c-.05-.013-.597-.09-.674-.13c-.078-.04-.39-.13-.507-.195s-.286-.156-.39-.156c-.103 0-.532.052-.61.04c-.078-.014-.312.025-.403.038c-.09.013.117.182-.077.22c-.195.04-.17.066-.13-.13c.038-.194-.13-.246-.3-.168c-.168.078-.44.13-.376.22c.065.092-.012.158.117.248c.13.09.183.117.35.104c.17-.013.34.025.34.025s0 .158-.065.183c-.065.026-.17.026-.196.104c-.025.078-.155.117-.155.078s.066-.17-.025-.234c-.09-.064-.117-.077-.22-.012c-.105.065-.117.09-.17-.013c-.052-.104-.208-.195-.208-.195s-.104-.13-.182-.194c-.077-.065-.22-.052-.234.013c-.012.063.027.128.08.246c.05.117.103.337.012.35c-.09.014-.104.027-.195.053c-.09.025-.13-.04-.13-.144s-.04-.195-.013-.234c.026-.04-.104.027-.234 0c-.13-.025-.233.052-.104.092c.13.04.157.194.04.233c-.118.04-.56 0-.703 0s-.35.04-.39-.04c-.04-.077.118-.128.208-.128c.092 0 .364.01.468-.13c.104-.144-.13-.17-.233-.17c-.104 0-.183-.04-.3-.155c-.117-.117.08-.195.053-.247c-.026-.052-.156-.014-.272-.014c-.117 0-.3-.09-.3.014s.144.402.053.337c-.092-.064-.08-.156-.144-.234c-.065-.077-.168-.064-.3-.05c-.128.012-.35.05-.414.038c-.064-.013-.013-.013-.156-.078c-.142-.065-.208-.052-.312-.117c-.103-.067-.012-.092-.194-.105c-.18-.013-.168.09-.35.065c-.182-.026-.234.013-.416 0c-.182-.013-.272-.026-.3.065c-.024.09-.077.247-.155.247s-.17.09.077.104c.247.012.105.128.325.116c.22-.013.416-.013.468-.117c.052-.105.09-.105.117-.066c.025.04.22.272.22.272s.13.104.183.13c.05.026-.052.143-.156.078c-.104-.065-.3-.05-.377-.116c-.078-.065-.43-.065-.52-.052c-.09.013-.247-.04-.3-.04c-.05 0-.22.13-.22.13s-.065-.104-.103-.156c-.04-.052-.104.052-.156.065c-.052.013-.208-.104-.364-.052s-.104.104-.325.09s-.273-.05-.35-.038l-.08.013s.013.052-.077.104s-.157.117-.222.04c-.063-.08-.102-.08-.14-.092c-.04-.013-.12-.078-.287-.04c-.17.04-.43-.05-.48-.05c-.04 0-.313-.016-.49-.034C9.63 5.914 12.645 4.466 16 4.466c2.128 0 4.117.59 5.83 1.598a.49.49 0 0 1-.388.023c-.078-.043-.158-.078-.475-.06c-.317.017-.665.12-.595.225c.072.104-.142.165-.197.113c-.055-.052-.31.06-.293.165c.016.104-.04.225-.175.2c-.134-.028-.23.06-.237.145c-.007.087-.31.147-.332.147c-.024 0-.412-.008-.27.095c.097.07.15.027.27.052c.12.025.214.216.277.242c.062.026.15 0 .19-.052c.04-.052.094-.234.094-.234s0 .173.096.208c.095.035.33-.026.395-.017c.065.008.438.06.54.112c.103.05.355.086.427.198c.07.113.08.503.12.546c.04.043.173-.14.204-.182c.032-.044.2-.018.255.042c.056.06.182.208.175.27c-.01.06-.033.155-.08.12c-.048-.033-.127-.024-.096-.094c.032-.07.048-.217-.015-.217h-.12s-.12-.035-.2.095s-.014.26.04.26s.185 0 .185.034c0 .035-.136.14-.128.2c.01.06.11.268.144.312c.032.043.198.086.245.096c.05.008-.11.017-.07.077c.04.06.102.208.19.243c.086.035.332.19.362.26c.032.07.222-.052.262-.06c.04-.01.03.18.142.19c.11.008.15-.018.245-.096s.072-.182.08-.26c.008-.078 0-.138.103-.113c.104.027.158-.017.15-.103c-.008-.087-.095-.19.07-.217c.167-.026.254-.138.357-.138c.103 0 .39.043.42 0c.03-.043.166-.243.253-.25c.067-.008.224-.022.385-.043a11.472 11.472 0 0 1 2.673 6.904c-.118.16-.012.305.02.408c.002.03.006.058.006.088c0 .136-.016.27-.02.404l-.004.006zm-9.716-4.326c-.064.013-.17-.052-.17-.143s-.09.17-.04.248c.054.078-.103.17-.154.17s-.09-.117-.078-.234c.014-.117-.077-.22-.22-.208c-.144.014-.21.13-.26.17c-.053.038-.053.258-.04.31s.013.236-.116.222c-.118-.013-.092-.233-.08-.312c.015-.078-.038-.273.015-.376c.054-.104.208-.143.313-.156s.324.065.363.052c.04-.014.222-.014.312 0c.09.01.22.245.156.258zm.233.04c.04.025.31-.04.364.025c.052.064-.053.077-.182.13c-.13.05-.17.038-.22.103s-.222.09-.3.168c-.078.08-.217.126-.246.066c-.04-.078.013-.04.025-.078c.013-.04.245-.13.245-.13s.276-.31.315-.285zm-1.04-.456c-.037.013-.18-.026-.3-.026c-.115 0-.09-.078-.142-.064c-.05.013-.168.04-.247.078c-.078.04-.208.03-.208-.04c0-.104.052-.078.22-.143c.17-.065.353-.247.43-.17c.078.08.22.17.312.183c.092.014-.026.17-.064.182zm-1.328-4.03c-.08.025-.348.138-.322.198c.01.023.078.07.19.052c.113-.018.276-.035.355-.043c.078-.01.095-.14.01-.147c-.088-.01-.157-.087-.234-.06zm-.962.103c-.06.027-.243-.042-.338.02c-.06.037-.026.163.07.17c.095.01.26-.06.276-.007c.018.052.078.286.234.208c.156-.077.147-.146.19-.155c.043-.01-.008-.2-.078-.243c-.07-.043-.294-.017-.354.01zm-.313-.735c.017.044-.008.078.113.095c.12.018.173.035.243.035s.043-.113-.017-.19c-.06-.078-.043-.07-.2-.113a1.237 1.237 0 0 0-.415-.035c-.104.01-.217-.017-.243.104c-.013.063.07.113.174.113s.328-.05.345-.008zm.226.476c.044.096.044.052.166.062c.12.01.215-.12.224-.164c.01-.044-.06-.13-.225-.113s-.667-.026-.736.034c-.067.057 0 .232-.027.25c-.026.017.01.095.077.078c.07-.017.104-.182.157-.182c.052 0 .32-.06.364.035m-1.628-.354c.052.043.183.008.173-.035c-.008-.042.053-.216-.05-.224c-.105-.008-.25.096-.408.148c-.1.033-.078.13-.01.13s.244-.06.295-.018zm.738.156c-.087.043-.114.07-.19.052c-.08-.017-.08-.156-.218-.13c-.138.026-.164.104-.207.14s-.14.06-.173.042c-.034-.017-.234-.13-.234-.13s-.416-.017-.433-.07c-.017-.05-.086-.137-.277-.12s-.52.13-.572.13c-.052 0 .062.104-.01.104c-.068 0-.154-.008-.18.07c-.018.052.078.05.19.05c.11 0 .294 0 .346-.025c.052-.026.312-.087.303-.01c-.01.08.104.2.164.183c.06-.017.182-.13.242-.086c.06.043.07.146.13.173c.06.025.226.025.304 0c.077-.027.294-.027.39-.01c.094.02.372.07.398.02c.026-.054.104-.062.112-.114s.05-.215.05-.215s-.05-.096-.136-.053zm4.385 8.957c-.12.02-.09.25.052.21c.143-.043.066-.233-.052-.21m1.13-.31c-.104-.027-.22 0-.3.012c-.077.013-.298.208-.298.208s.143.026.233.026c.092 0 .144.05.22.09c.08.04.222-.052.273-.052c.053 0 .118.156.13-.013c.015-.168-.154-.245-.258-.27zm-3.62-8.423c-.042-.104-.258-.14-.303-.035c-.038.09.347.14.304.035zm2.478 7.785c.143-.026.064-.144-.053-.13c-.118.013-.09.156.053.13m1.727.803c-.12.022-.092.253.05.21c.145-.04.067-.23-.05-.21m-1.573-.387c-.09.013-.285-.09-.39-.182c-.103-.09-.298-.09-.376-.09c-.076 0-.39.09-.39.09c-.012.13.118.09.274.09s.43-.025.48.04c.052.064.285.168.35.22c.065.053.273.066.286.014s-.142-.195-.233-.182zm-1.55-8.296c-.13-.01-.294-.01-.398 0c-.105.008-.183-.07-.26-.113c-.078-.044-.252-.183-.355-.2c-.104-.017-.086-.017-.303-.07c-.11-.026-.294-.06-.294-.085c0-.026-.052.12.043.165c.096.043.252.12.364.164c.114.043.33.052.4.14c.068.085.302.155.302.155l.277.025s.19-.043.39-.026c.2.016.493.042.66.034c.162-.008.188-.06.207-.095c.016-.036-.304-.105-.383-.096c-.076.006-.52.006-.65-.002zm.485 1.992c-.034.04.157.095.19.043c.035-.05-.1-.138-.19-.043m.582.18c.086-.034.043-.138-.08-.103c-.137.04-.008.14.08.105z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globealt(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 1.466C7.973 1.466 1.466 7.973 1.466 16c0 8.027 6.507 14.534 14.534 14.534c8.027 0 14.534-6.507 14.534-14.534c0-8.027-6.507-14.534-14.534-14.534M27.436 17.39c0 .002.004.002.005.004c-.02.187-.053.37-.084.554l-.047-.036c-.104-.09-.255-.128-.32-.115c-.156.032.23.305.268.342c.01.008.03.03.062.057c-1.03 5.312-5.71 9.338-11.32 9.338c-4.122 0-7.735-2.18-9.775-5.44c.123-.017.24-.017.28-.077c.05-.077.102-.24.178-.33c.077-.09.165-.23.127-.293c-.04-.064.1-.344.088-.42c-.013-.075-.127-.255.064-.406s.394-.383.407-.445c.01-.063.165-.33.15-.458c-.01-.127-.15-.28-.24-.318c-.09-.038-.28-.05-.355-.152c-.077-.103-.292-.203-.368-.178c-.076.025-.204.05-.305-.015c-.102-.062-.267-.14-.33-.19a.67.67 0 0 0-.305-.087c-.077 0-.065-.052-.178.1c-.115.154 0 .205-.205.178c-.204-.023.025-.036.14-.19c.114-.15-.012-.24-.14-.202c-.126.038-.038.115-.24.153c-.204.036-.204-.09-.077-.115s.355-.14.355-.19c0-.05-.025-.19-.127-.19s-.077-.127-.23-.292a.774.774 0 0 0-.298-.204c-.09-.58-.15-1.167-.15-1.77c0-2.845 1.04-5.447 2.75-7.46a.148.148 0 0 1 .07-.035c.084-.008.31-.024.51-.058c.202-.034.418-.16.49-.153c.068.007.24.007.185-.043c-.056-.057-.134-.175-.25-.175s-.008-.186.023-.186c.03 0 .186.025.178.11c-.007.085.047.194.2.16c.157-.033.133-.16.103-.194c-.032-.034-.055-.12-.085-.178c-.03-.06-.023-.136.14-.153c.162-.015.386.068.378-.067s0-.28.046-.356c.047-.075.148-.127.093-.152c-.054-.025-.124.034-.18-.093c-.053-.127.048-.22.11-.27c.02-.016.066-.06.123-.113a11.452 11.452 0 0 1 6.647-2.124c3.213 0 6.122 1.323 8.214 3.45c-.008.022-.01.052-.03.056c-.078.013-.167.063-.18-.05c-.013-.115-.013-.332-.102-.204c-.09.127-.127.127-.127.19c0 .064.076.128.05.242c-.025.114-.025.19.015.19c.036 0 .126-.114.24-.14c.115-.025.318-.088.33.026c.014.115.14.152.014.203c-.128.05-.267.025-.293-.052c-.024-.077-.113-.077-.202-.013c-.088.063-.28.292-.28.292s-.305.14-.342.114c-.04-.024.1-.164.203-.227c.1-.064.177-.204.14-.242c-.04-.037-.09-.278-.064-.342c.025-.063.14-.152.013-.216c-.128-.063-.218-.14-.32-.178s-.215.152-.304.204c-.09.05-.076.114-.19.127c-.115.013-.19.165 0 .254c.19.088.254.15.203.203c-.052.05-.268-.025-.268-.025s-.165-.076-.268-.076c-.1 0-.23-.063-.33-.076c-.102-.013-.306-.013-.355.038c-.052.05-.18.203-.28.152c-.102-.05-.102-.102-.242-.05c-.14.05-.28-.04-.355.037c-.077.076-.013.076-.255 0c-.24-.076-.19.05-.42.09s-.367-.04-.43.037c-.065.077-.154.217-.19.127c-.04-.088.125-.24.06-.292c-.06-.05-.33-.025-.366.013c-.04.038-.014.178.01.23c.027.05.065.253-.01.215c-.077-.038-.064-.166-.14-.152a.63.63 0 0 0-.204.078c-.038.025-.19.025-.23.076c-.036.05.015.19-.05.203c-.063.013-.114.064-.254-.025c-.14-.09-.14-.038-.178-.012c-.038.025-.216.127-.23.012c-.012-.114.026-.152-.088-.23c-.115-.075-.026-.075.127-.024c.152.05.343.075.622-.013c.28-.09.395-.127.28-.178c-.115-.05-.23-.1-.406-.127c-.178-.025-.42-.025-.7-.127c-.278-.102-.342-.14-.456-.165c-.115-.026-.813-.14-1.132-.09c-.317.052-1.193.28-1.245.32s-.128.19-.292.317c-.165.127-.47.42-.712.47c-.24.05-.52.254-.52.305c0 .05.1.242.075.28c-.025.038.05.23.19.28c.14.05.382.038.394-.04c.014-.075.204-.24.217-.126c.012.115.14.292.113.368c-.025.077 0 .153.09.14c.088-.012.56-.114.56-.114s.152-.064.126-.166c-.027-.1.165-.24.202-.28c.038-.037.178-.19.014-.24c-.167-.05-.293-.064-.115-.216s.293 0 .522-.23c.23-.228-.05-.29.19-.304c.242-.013.497-.025.445.05c-.05.077-.342.243-.508.32c-.167.076-.14.215-.077.29c.063.077.09.255.204.23c.113-.025.254-.114.38-.1c.128.01.383-.014.42-.014c.04 0 .216.178.114.203c-.1.026-.228.014-.444.026c-.215.013-.456.013-.456.05c0 .04.292.128.19.192c-.102.063-.203-.013-.33-.026c-.128-.012-.204.166-.242.267c-.04.1.063.28-.127.215c-.19-.063-.33-.063-.38-.038c-.05.025-.203.076-.33.114c-.127.04-.077-.062-.243-.062h-.164l-.103.013s-.1-.062-.114-.164c-.013-.102.05-.216-.013-.24c-.064-.027-.292.01-.33.087c-.038.076-.077.216-.026.28c.052.063.204.19.064.152c-.14-.038-.316-.05-.418.026c-.1.077-.28.242-.28.242s-.317.025-.317.102c0 .077 0 .178-.114.19c-.114.014-.267.05-.42.077c-.152.026-.138.09-.316.103s-.204.09-.038.114c.165.025.418.127.43.24a.593.593 0 0 1-.075.357c-.042.08-.304.026-.457.026c-.152 0-.456-.05-.584 0c-.128.05-.103.305-.065.42c.04.113-.012.177-.063.215c-.05.038-.065.152 0 .204c.063.05.114.164.166.177c.05.013.215-.038.28.025c.063.064.126.216.164.178c.04-.038.09-.203.153-.166c.064.04.216-.012.33-.025s.178-.14.293-.205c.114-.063.05-.063.013-.14c-.037-.076.115-.165.205-.254c.088-.088.253-.012.292-.114c.038-.102.05-.28.15-.267c.104.012.244.075.332.075c.088 0 .278-.14.33-.165c.05-.025.242-.013.268.102c.025.114.24.254.292.28c.05.024.38.126.433.164c.05.038.126.153.152.254c.025.102.114.102.128.013c.012-.09-.065-.254.025-.242c.088.014.19-.025.19-.025s-.242-.165-.33-.203c-.088-.038-.255-.114-.33-.24c-.077-.128-.268-.154-.255-.28c.013-.127.19-.05.292.05c.102.103.356.242.445.33c.088.09.23.128.267.243c.04.114.152.24.19.292c.038.05.165.33.204.394c.04.064.166-.01.23-.062c.063-.05.18-.076.19-.178c.014-.102-.152-.178-.202-.216c-.05-.038.127-.076.19-.127c.064-.05.178-.14.23-.063c.05.077.025.38.05.432c.025.05.28.127.33.19c.05.064.268.09.305.052c.04-.038.242.026.294.038c.05.013.202-.025.304-.05c.104-.025.205-.102.192.063c-.013.164-.05.418-.18.545c-.126.127-.075.19-.2.19c-.06 0-.114 0-.157.022c-.04-.065-.098-.117-.175-.097c-.152.038-.344.038-.47.19c-.128.153-.178.165-.204.114c-.025-.05.37-.267.317-.33c-.05-.064-.355-.04-.52-.04c-.167 0-.306-.1-.434-.126c-.127-.025-.293.127-.42.254s-.215.038-.33.038c-.115 0-.33-.165-.33-.165s-.217-.09-.306-.09c-.088 0-.267-.164-.318-.164c-.05 0-.19-.114-.088-.165c.1-.05.202.05.1-.23c-.1-.278-.33-.215-.418-.177c-.088.04-.724.025-.775.025c-.05 0-.42.127-.533.178c-.116.052-.318.116-.37.14c-.05.026-.317-.05-.432.014c-.15.084-.29.216-.33.216c-.038 0-.153.09-.23.28c-.076.19.014.355-.127.42c-.14.062-.394.203-.495.304c-.102.1-.23.458-.355.623c-.127.165 0 .317.025.42c.025.1.114.29-.025.47c-.14.178-.127.266-.19.28c-.064.012.062.062.087.19c.025.127-.114.254.128.368c.24.113.355.217.418.367c.064.153.382.407.382.407s.23.204.344.292c.114.09.152.038.177-.05c.024-.09.177-.104.354-.104c.178 0 .305.04.483.014c.178-.025.356-.14.42-.166c.063-.025.28-.164.443-.063c.166.103.14.24.23.332c.088.088.24.037.355-.05c.114-.09.064-.053.203.024c.14.074.204.15.077.266c-.128.113-.05.293-.128.47c-.076.178-.063.203.077.278c.14.076.394.548.47.638c.077.088-.025.342.064.495c.09.15.178.254.077.33c-.103.076-.28.217-.292.47s.05.432.102.522s.177.33.24.418c.065.09.14.305.153.445c.013.14-.024.306.04.38c.063.077.1.192.215.293c.115.103.152.318.152.318s.04.09.05.23c.013.14.026.227.153.29c.126.064.215.077.28.014c.063-.063.38-.077.546-.063c.165.013.355-.075.52-.19s.408-.42.497-.508c.09-.09.293-.255.27-.356c-.026-.1-.078-.202.023-.253c.102-.052.344-.152.356-.23c.012-.076-.09-.394-.116-.456c-.024-.063.064-.18.165-.305c.102-.128.42-.216.47-.267c.05-.053.19-.267.217-.433c.024-.167-.05-.37 0-.457c.05-.09.013-.165-.103-.268c-.114-.102-.09-.407-.127-.457c-.037-.05-.013-.32.063-.345c.076-.023.242-.28.344-.393c.102-.114.394-.47.534-.496c.14-.026.355-.23.368-.344c.013-.115.38-.547.394-.635c.013-.09.166-.42.102-.496c-.062-.076-.56.115-.622.14c-.064.026-.24.128-.446.114c-.202-.013-.114-.177-.127-.254c-.012-.076-.228-.368-.28-.38c-.05-.013-.202-.167-.266-.318c-.062-.153-.15-.343-.253-.458c-.102-.114-.165-.38-.268-.56c-.102-.177-.19-.406-.28-.57c-.02-.042-.045-.08-.067-.118c.118-.03.29-.082.31-.01c.024.09.165.28.19.42s.165.09.178.216c.015.128.14.433.19.47c.053.038.28.242.32.318c.037.076.088.178.126.37c.038.19.076.443.18.48c.1.04.443-.063.507-.1s.482-.243.635-.256c.153-.012.18-.115.368-.152c.19-.038.33-.177.458-.28c.127-.1.28-.355.33-.444c.052-.088.18-.152.115-.253c-.064-.104-.332-.255-.434-.27c-.102-.01-.09-.177-.152-.177s-.05.088-.178.153c-.127.063-.255.19-.344.165s.026-.088-.113-.202s-.193-.14-.193-.228c0-.09-.278-.255-.304-.382c-.026-.127.19-.305.254-.19c.063.114.115.292.28.368c.164.076.317.204.394.23c.077.024.268-.14.33-.115c.064.025.192.254.307.293c.113.038.495.05.56.05s.33.014.38-.062c.05-.075.09-.075.153-.075c.062 0 .177.23.267.254c.09.024.254.012.24.178c-.01.164.077.305.166.317c.09.012.293-.19.293-.19s0 .317-.012.432c-.014.113.14.534.14.534s.19.394.24.483s.267.355.267.47c0 .115.026.293.104.293c.076 0 .152-.203.24-.33c.09-.127.116-.306.153-.433c.038-.127.038-.356.038-.444c0-.09.075-.167.255-.243c.178-.076.304-.292.456-.407c.153-.115.14-.305.446-.305c.305 0 .278 0 .355-.077c.076-.076.15-.127.19.013c.038.14.254.344.292.395a.57.57 0 0 1 .102.344c-.013.15.012.33.075.33s.19-.217.19-.217s.28-.19.268.013c-.014.203.025.42.025.545c0 .054.042.136.088.21c-.005.06-.004.12-.01.18c-.09.023-.09.197-.042.26zm-7.054-5.326c.076.05.102.127.152.203c.052.076.14.05.203.114c.063.065-.178.14-.075.217c.1.077.15.38.165.458c.013.076-.28.114-.37.102c-.088-.013-.353-.102-.444-.127c-.09-.026-.14-.343-.025-.33c.116.012.14-.026.267-.14c.128-.115-.19-.166-.278-.19c-.09-.026-.268-.306-.33-.395c-.063-.09-.015-.228.14-.33c.076-.052.28.062.38 0c.102-.064.204-.14.242-.166c.038-.026.292.037.33.113c.038.076.19.19.14.23c-.052.037-.28.075-.356 0c-.075-.078-.255.01-.268.15c-.014.142.05.04.126.09zm-3.507.216c-.077-.025.025-.178.102-.23c.075-.05.164-.177.24-.304c.077-.127.18-.14.242-.127c.062.012.202.24.24.317c.038.076.165-.026.217-.05c.05-.026.127-.103.14-.166s.127-.1.254-.1s.014.1-.075.126c-.09.025-.038.077.113.127c.153.05.293.19.46.28c.164.088.19.266.087.29c-.1.026-.406.052-.52.04c-.115-.014-.255-.128-.42-.154c-.165-.026-.37-.014-.433.076s-.292.05-.395.05c-.102 0-.228.127-.253.077c0 0 .077-.227 0-.253zm.432-2.822c.063-.178.42.038.355.127c-.063.09-.398-.006-.355-.127m.495 9.126c.063.102-.14.43-.254.407c-.113-.026-.076-.317-.038-.38c.038-.065.26-.08.292-.026zm-4.613-5.91c.024-.05-.04-.153-.128-.013c-.03.05.102.065.127.014zm7.623-4.64c.14.077.34.108.433.014c.076-.076.013-.204-.05-.216c-.064-.013-.104-.115.062-.203c.165-.09.343-.205.534-.23c.19-.025.622-.038.774 0c.152.04.382-.166.445-.254s-.202-.152-.278-.05c-.077.1-.444.075-.52.05c-.077-.025-.687.102-.813.102c-.13 0-.18.152-.357.23c-.18.075-.42.19-.51.228c-.087.038-.176.19-.1.216c.076.025.24.037.38.113m-6.67 4.64c.063-.09-.052-.217-.115-.217c-.12 0-.178.19-.103.254c.077.066.153.053.217-.036zm.57.343c.065.025.115.102.166.114c.05.014.216 0 .166-.126s-.167-.127-.204-.127c-.038 0-.203-.038-.267 0c-.047.028.075.115.14.14zm-3.405-2.06c.1.014.217-.062.305-.1c.088-.038.216-.114.216-.23c0-.113-.026-.215-.078-.266c-.05-.05-.14-.063-.216-.05c-.115.02-.127.14-.203.14c-.076 0-.165.025-.14.114s.077.152 0 .19c-.076.04.012.192.115.204zm.623-.545c.128.05.395.102.293.153c-.102.05-.28.19-.305.267s.216.153.216.153s-.077.09-.013.114c.064.024.103-.09.204-.09c.1 0 .304.063.406.063c.103 0 .267-.14.254-.23c-.013-.088-.14-.228-.254-.28c-.113-.05-.24-.28-.317-.33c-.076-.05.076-.178-.013-.267c-.09-.09-.153-.076-.255-.14c-.102-.063-.19.013-.254.09c-.062.075-.14-.014-.216.01c-.102.036-.063.167-.012.23c.052.065.142.205.27.256zm12.8 6.786c-.084.037-.154.47 0 .52c.15.053.24-.2.19-.266c-.052-.062-.077-.305-.19-.254m-4.616 3.266c-.16-.045-.177.166-.304.306c-.128.14-.267.254-.317.24c-.052-.012-.33.09-.242.28c.09.19.077.382-.012.472c-.09.088.076.342.052.482c-.026.14.037.23.215.23s.242-.065.318-.23c.076-.166.088-.33.164-.47c.077-.14.14-.434.18-.51c.037-.075.113-.316.1-.457c-.013-.138-.063-.318-.153-.343zM10.39 8.802c-.068-.06-.228-.102-.305-.11c-.076-.008-.152.06-.32.06c-.17 0-.28.067-.348 0c-.068-.068-.35-.102-.375-.06c-.034.057-.1.06-.034.178c.07.118.12.186.18.178s.26-.017.287.05c.025.07.094.128.237.086c.145-.042.263-.068.296-.12c.033-.05.263-.058.263-.058s.188-.145.12-.204z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobealtTwo(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 1.466C7.973 1.466 1.466 7.973 1.466 16c0 8.027 6.507 14.534 14.534 14.534c8.027 0 14.534-6.507 14.534-14.534c0-8.027-6.507-14.534-14.534-14.534M8.25 7.48c.123.055.256.104.28.137c.04.05.09.206.027.244c-.065.04-.167.027-.18-.09c-.014-.114-.027-.127-.18-.18l-.052-.014l.106-.096zm-.517.494c.03.087.113.125 0 .17a.336.336 0 0 1-.174.02c.057-.062.112-.13.173-.19zM16 27.534C9.64 27.533 4.466 22.36 4.466 16c0-.085.01-.168.013-.254l.01-.01c.13-.1.284-.358.335-.45c.052-.088.18-.153.116-.255c-.058-.095-.29-.23-.406-.26c.01-.1.032-.195.045-.294c.062.077.136.17.207.194c.115.038.5.052.566.052c.063 0 .334.014.386-.064c.05-.077.09-.077.153-.077c.064 0 .18.232.27.258c.09.026.258.013.245.18c-.012.167.077.31.167.322c.09.013.296-.194.296-.194s0 .322-.01.438c-.014.114.14.54.14.54s.193.397.244.488c.052.09.27.36.27.476c0 .117.026.297.104.297s.155-.206.244-.335c.09-.128.117-.31.155-.438c.04-.13.04-.36.04-.45c0-.09.075-.168.256-.245c.18-.077.31-.296.463-.412c.155-.116.142-.31.452-.31c.308 0 .282 0 .36-.077c.077-.077.154-.128.192.013c.04.143.257.348.296.4c.04.052.116.193.104.348c-.013.153.012.334.077.334c.064 0 .193-.22.193-.22s.283-.19.27.015c-.014.205.025.425.025.552c0 .13.232.437.232.36c0-.078.103-.295.103-.412c0-.114.064-.063.23.05c.168.117.284.35.284.35s.168.154.193.22c.026.063.206-.026.244-.105c.04-.076.066-.115.168-.14c.104-.027.23-.027.27-.17c.04-.14.155-.307 0-.5c-.153-.194-.23-.322-.346-.413c-.117-.09-.206-.322-.206-.322s.244-.218.32-.296a.122.122 0 0 1 .208.064c.013.09-.115.168-.14.36c-.027.193.153.258.205.193c.05-.066.18-.22.18-.258c0-.04-.09-.026-.103-.167c-.013-.142.166-.245.23-.207c.066.04.477-.05.67-.154s.308-.322.425-.412c.116-.09.515-.386.49-.527c-.027-.142.01-.334-.09-.515c-.104-.18-.233-.295-.284-.373c-.05-.077.22-.09.347-.206c.13-.115 0-.218-.064-.205c-.063.013-.23.052-.295.04c-.064-.014-.103-.078-.206-.156s.025-.193.09-.18c.064.013.23-.13.308-.193c.077-.064.193-.115.154-.05c-.038.063-.128.295-.026.308c.104.013.348-.193.388-.18c.038.013.102.18.064.257c-.04.077-.04.206.013.193c.052-.013.155-.13.18-.09c.028.04.155.116.09.257c-.062.142-.192.193-.038.284c.154.09.206.012.322-.052c.115-.064.193-.347.128-.438c-.064-.09-.218-.27-.218-.334s.257-.064.257-.167s.09-.18.18-.22c.092-.038.207-.205.245-.153c.04.052.27.116.334.04c.063-.078.4-.36.604-.516a.904.904 0 0 0 .336-.515c.05-.18.128-.295.102-.436c.077.18.09.31.077.45c-.014.142 0 .438.025.476c.025.04.13.128.192.103c.064-.025-.025-.283-.025-.334c0-.053.09-.13.142-.143c.052-.013 0-.23-.065-.322c-.064-.09-.155-.142-.103-.154c.05-.012.115-.115.077-.14c-.04-.026-.014-.117-.103-.09c-.064.018-.24-.016-.234.094c-.037-.11-.116-.183-.216-.172c-.116.013-.18.077-.296.077s-.025-.18-.077-.18c-.05 0-.168.166-.23.076c-.065-.09.18-.206.372-.27c.19-.064.513-.438.643-.45c.128-.014.45.025.733.012c.283-.013.373-.13.463-.064s.282.142.398.13c.116-.015.064 0 .244-.13s.348-.193.438-.296c.09-.103.335-.18.348-.077c.014.104-.026.207.077.207s.258-.103.386-.154c.13-.05.232-.116.232-.116s-.527.36-.655.438c-.13.077-.438.13-.567.283c-.128.155-.205.206-.192.374c.014.167.23.386.128.54c-.103.154-.14.373-.14.373s.153-.22.372-.36s.348-.334.425-.412s.31-.09.31-.18s.063-.207.103-.31c.038-.103-.077-.078 0-.206c.076-.13.064-.232.45-.232s.257.026.566.013c.31-.013.424-.167.72-.245c.296-.077.527-.128.618-.09c.09.04.232.013.14-.077c-.088-.09-.294-.22-.192-.245c.104-.025.207-.038.246-.14c.04-.104-.142-.284-.04-.387c.105-.103-.076-.23-.206-.257c-.128-.025-.63.026-.73-.025c-.105-.05-.272-.115-.323-.077c-.052.04-.168.245-.168.245s-.09.025-.168-.09c-.077-.116-.5-.103-.63-.103s-.27.025-.413.04c-.14.012-.22.05-.322-.04c-.102-.09-.243-.13-.296-.167c-.052-.04-.335-.04-.554-.012c-.218.025-.438.025-.438.025s-.104-.038-.257-.128c-.153-.09-.308-.154-.36-.154c-.05 0-.45.064-.54 0c-.09-.064-.18-.103-.244-.103s-.115-.103-.038-.103s.437-.103.437-.103s-.103-.143-.23-.143c-.13 0-.36-.064-.425-.064s-.014.064-.142.04c-.13-.027-.258-.08-.335-.027c-.076.05-.258.128-.064.18c.192.052.372 0 .424.078c.052.077 0 .115 0 .167s-.103.194-.167.22c-.064.025-.143-.04-.27.025c-.13.064-.45.013-.49.052c-.038.04-.115-.103-.18-.077c-.064.025-.232.193-.322.18c-.09-.013-.206-.103-.206-.206s-.04-.232-.078-.258c-.038-.025-.322-.04-.425-.025c-.103.014-.424.04-.477.09c-.052.053-.193.09-.283.09s-.167-.09-.36-.115c-.192-.026-.617-.04-.67-.026s-.217-.026-.154-.078c.065-.05.257-.22.143-.295c-.117-.078-.375-.078-.49-.09c-.116-.013-.23-.04-.412-.013c-.18.026-.22.116-.296.04c-.077-.078.193.038-.077-.078c-.27-.116-.398-.103-.476-.064c-.077.04.013.025-.192.103a4.76 4.76 0 0 1-.374.13c-.05.01-.372-.066-.41-.092c-.04-.025-.182.013-.31.064s-.223.128-.35.103a.93.93 0 0 0-.465.025c-.13.05-.22.05-.347.14c-.13.09-.373.143-.373.143L8.448 7.3A11.476 11.476 0 0 1 16 4.467c6.36 0 11.534 5.173 11.534 11.534S22.36 27.533 16 27.533zm-1.112-7.614s.207-.026.207-.117c0-.09-.207-.205-.282-.102c-.078.103-.22.206-.207.297c.02.14.282-.077.282-.077zm-.013-2.897c-.18.233-.167.182-.296.128c-.13-.05-.335.117-.297.183c.04.064.322-.014.386.102c.064.116.064.13.192.104c.128-.027.257-.206.22-.296c-.04-.093-.207-.22-.207-.22zm-.038 1.222c-.05 0-.412.064-.45.08c-.04.012-.27-.026-.27-.026c-.09.088-.027.178.115.165s.438-.052.502-.052c.065 0 .154-.168.103-.168zm-.553-3.577c-.19.03-.308.438-.155.425c.154-.012.32-.45.154-.425zm.45 2.29c-.052-.063-.18-.27-.323-.218c-.04.017-.152.245-.01.245c.14 0 .386.038.334-.026zm.116-.153c.232-.013.167-.245-.013-.257c-.05-.004-.22.27.013.257m2.74 2.123c-.192-.04-.243-.102-.45-.205c-.206-.103-.67-.103-.68-.04c-.015.065 0 0-.156-.05c-.153-.053-.27 0-.31-.09c-.037-.09-.127-.117-.243-.002c-.096.098-.14.105.08.144c.217.04.282.04.19.14c-.09.105-.153.234-.076.245c.077.014.31-.05.334 0c.026.053-.05.063.207.104c.258.037.31.128.36.178c.05.052.205.22.103.22c-.104 0-.22.128-.142.143c.077.013.31-.04.32 0c.015.037.144.283.272.27c.13-.012.206-.243.27-.31c.065-.063.322-.103.35.013a.456.456 0 0 0 .256.31c.154.077.335.155.348.09c.013-.064-.077-.31-.18-.346c-.104-.04-.283-.26-.283-.348c0-.09-.156-.117-.233-.182c-.078-.065-.143-.245-.336-.284zm-9.548-1.905c-.084.037-.155.476 0 .527c.154.052.244-.205.193-.27c-.052-.062-.077-.307-.193-.257m7.387 1.094c-.12-.05-.336.424-.182.463c.155.04.27-.424.18-.463zm.257-4.414c.077 0 .18-.05.18-.193c0-.142.18 0 .27-.013s.14-.103.18-.206l.01-.026c-.004.024-.002.093.094.117c.155.04.206-.063.206-.102s.283-.103.336-.142c.05-.038.258-.103.27-.154c.013-.05 0-.348.064-.373c.064-.027.154-.027.052-.207c-.104-.18-.104-.348-.232-.27c-.095.056-.038.283-.115.437s-.14.296-.192.296s-.32.103-.4.18c-.075.077-.45-.064-.5 0c-.052.064-.154.14-.22.193c-.064.05-.244.012-.206.166c.036.156.127.297.205.297zm1.762-1.647c.18-.013.347-.064.347-.064s.27.013.232-.116c-.04-.128-.323-.14-.376-.128c-.05.013-.142-.142-.244-.116c-.096.023-.128.155-.128.193c0 .04-.36.115-.245.22c.116.1.233.024.412.01zm-3.54 7.002c.104.064.296-.22.35-.13c.05.09-.014.13.075.246c.09.114.258.102.258.102s-.013-.31-.155-.387c-.142-.078-.232-.167-.064-.142c.167.026.257-.04.22-.114c-.04-.078-.284-.04-.362-.026s-.193-.052-.193-.052c-.078.024-.064.09-.09.22c-.027.126-.143.216-.04.282zm7.014 2.56c-.23-.052-.077.04 0 .154c.077.116.232.176.258.05c.01-.063-.027-.152-.258-.204m.99 3.126c-.076.064 0 .09-.218.22c-.22.13-.49.27-.54.386c-.053.116.05.18.257.192c.206.013.154.053.296-.103s.27-.245.437-.374c.168-.128.168-.322.168-.322s-.18-.178-.193-.14c-.022.06-.13.077-.207.14zm-3.41-3.126c.014-.116-.22-.116-.334-.207c-.116-.088-.128-.358-.193-.514c-.064-.153-.192-.257-.322-.397c-.128-.144-.192-.466-.23-.44c-.04.026-.154.4-.064.516c.09.116-.038.348-.102.503c-.065.152-.22-.027-.35-.105c-.128-.078-.307-.128-.397-.22c-.09-.09.156-.334.092-.425c-.065-.09-.412-.013-.45-.013c-.04 0-.116-.128-.194-.128c-.077 0-.064.257-.064.257s-.078-.09-.193-.207c-.116-.115.013.077-.102.193c-.117.117-.08.078-.13.206s-.166.076-.282-.053c-.116-.128-.18-.037-.258 0c-.077.04-.14.26-.18.31c-.04.05-.31.116-.374.18c-.064.063-.09.27-.09.323c0 .05-.27.023-.36.09c-.09.063-.23.024-.322.024c-.09 0-.4.244-.502.308c-.103.066-.103.298-.05.362c.05.063.153.22.09.244c-.065.026-.105.206.05.36c.154.154.103.193.115.27c.014.077.078.104.18.232c.103.128-.18.23-.218.31c-.04.076.09.192.167.257c.077.063.27.026.386-.013c.117-.04.245-.143.32-.155c.08-.014.44-.027.44-.027s.128-.192.218-.296c.09-.102.372-.013.372-.013s.117-.076.426-.14c.31-.066.18.063.296.103c.115.037.27.062.36.128c.09.064 0 .218-.013.283c-.015.064.218.038.23-.026c.013-.064.076-.128.206-.205c.128-.078.025.114.076.23c.052.117.13-.156.13-.025c0 .04.038.078.05.116c.014.04.18.052.18.18c0 .13 0 .207.04.23c.038.027.244 0 .335.156c.09.154.154.013.205-.052c.052-.064.23.026.283.078c.052.05.193-.104.387-.155c.192-.05.167-.04.22-.115c.05-.078.09-.283.204-.438c.115-.153.27-.424.27-.63c0-.207-.013-.683-.154-.9c-.14-.22-.41-.44-.398-.554zm-.475 3.152c-.066-.013-.208-.062-.208-.062c-.142.14.142.14.104.283c-.04.142.193.09.257.065c.063-.027.22-.323.193-.4c-.025-.078-.283.128-.347.115zm4.773-.592c-.052 0-.077.064-.192 0c-.116-.063-.09-.037-.167-.167c-.077-.126-.09-.295-.22-.23c-.05.026 0 .17.052.22c.053.05.077.23.064.282c-.013.052-.232.116-.13.18c.104.064.297 0 .27.078c-.024.077-.128.18-.012.205c.115.025.154-.09.207-.178c.05-.093.09-.17.18-.22c.09-.053 0-.17-.052-.17m.012-2.998c.168.064.464-.23.347-.27c-.115-.04-.347.27-.347.27m-8.892-1.274c-.03.115.193.167.206.04c.01-.13-.18-.143-.207-.04zm.347-3.436c-.064.065-.257.193-.283.31c-.025.115.31-.182.4-.296c.09-.117.27-.052.307-.117l.04-.063s-.142-.025-.257-.063c-.117-.038-.258.103-.193-.103c.064-.206.257-.167.22-.322a.246.246 0 0 0-.208-.193c-.09 0 .013.14-.116.23c-.128.09-.27.13-.193.284c.077.154.347.27.283.334zm-1.016 3.28c.013-.075-.142-.19-.206-.19c-.065 0-.386-.078-.386-.078c-.058.023-.135.045-.158.077c-.007-.012-.022-.025-.05-.04c-.14-.075-.308 0-.36-.102c-.05-.104-.127-.104-.18-.04c-.093.118.026.207.064.232c.038.024.18.052.31.04c.08-.01.18-.028.21-.06c.003.014.015.027.034.044c.103.092.167.13.32.116c.157-.01.39.08.402 0zm-.373-.758c.066 0 .194 0 .284.026c.09.025.386.05.373-.064c-.013-.115-.038-.297.09-.41c.13-.118.256-.18.192-.35c-.064-.166-.194-.27-.104-.348c.09-.076.192-.102.192-.166c0-.065-.217.18-.244-.246c-.005-.09-.206.025-.22.116c-.01.09.143.167-.102.167s-.257.194-.31.232c-.05.038-.102.05-.206.075c-.102.026-.127.13-.153.194c-.025.062-.206-.117-.257-.065c-.05.052-.013.296.077.5c.092.208.323.337.388.337zm-.9-.064c.064-.077.037-.192-.064-.18c-.103.013-.193-.168-.36-.283c-.168-.114-.296-.194-.45-.36c-.155-.167-.348-.27-.45-.36c-.105-.09-.258-.13-.323-.115c-.16.032.23.31.27.346c.04.04.388.335.388.478s.232.476.297.527c.064.053.385.245.437.35c.052.102.167.13.167-.014c0-.142.026-.31.09-.388zM11 17.474c.064.232.193.464.244.555c.052.088.27.216.348.28c.077.064.192-.024.143-.102c-.052-.078-.155-.192-.167-.283c-.013-.09-.078-.233-.18-.387c-.103-.153-.193-.192-.258-.295c-.064-.104-.296-.297-.296-.297c-.102.013-.102.205-.05.27c.048.064.152.027.216.258z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gplus(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.208 22.73a9.45 9.45 0 0 1-.78-1.23c-.266-.423-.398-.934-.398-1.527c0-.354.05-.654.152-.893c.09-.25.17-.482.248-.693c-.46.053-.893.078-1.3.078c-1.927-.023-3.444-.568-4.55-1.636v7.21c.99-.454 2.013-.772 3.073-.956c1.425-.21 2.61-.328 3.555-.354zm1.103 1.142a9.118 9.118 0 0 0-.907-.04c-.207-.025-.734 0-1.584.08a17.6 17.6 0 0 0-2.57.57c-.205.08-.495.197-.868.354c-.374.17-.753.414-1.14.73a4.335 4.335 0 0 0-.663.77v1.104a1.87 1.87 0 0 0 1.868 1.865h10.188l.002-.068a3.97 3.97 0 0 0-1.004-2.688c-.71-.738-1.817-1.625-3.323-2.678zm-3.56-7.278c.657.514 1.405.77 2.248.77c1.065-.038 1.952-.42 2.662-1.145a4.211 4.211 0 0 0 .66-1.582c.06-.54.088-.995.088-1.364c0-1.594-.408-3.202-1.224-4.822c-.382-.777-.886-1.41-1.51-1.897a3.868 3.868 0 0 0-2.198-.73c-1.096.023-2.008.464-2.734 1.32a4.994 4.994 0 0 0-.88 3.005c0 1.463.43 2.985 1.285 4.566a6.3 6.3 0 0 0 1.604 1.878zM27.554 2.707H4.447c-1.03 0-1.87.838-1.87 1.87v2.506c.12-.125.24-.25.368-.372c1.15-.946 2.345-1.564 3.584-1.854c1.226-.25 2.375-.375 3.45-.375h8.086l-2.5 1.46h-2.494c.254.157.54.387.857.69c.303.314.602.702.894 1.163a7.7 7.7 0 0 1 .743 1.538c.177.593.267 1.283.267 2.072c-.024 1.447-.343 2.604-.958 3.472c-.302.42-.62.81-.958 1.164c-.374.354-.77.718-1.193 1.085c-.24.25-.464.533-.67.848c-.24.33-.36.712-.36 1.147c0 .42.124.77.37 1.046c.21.264.415.493.613.688l1.372 1.126c.853.688 1.6 1.467 2.243 2.31c.604.854.92 1.972.943 3.354c0 .562-.07 1.106-.223 1.646H27.54c1.03 0 1.87-.838 1.87-1.87V4.576a1.854 1.854 0 0 0-1.857-1.87zm1.47 8.246H24.78v4.244h-2.06v-4.244h-4.24V8.897h4.242V4.654h2.057v4.243h4.243v2.056z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Graphael(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28.833 15.71a5.48 5.48 0 0 0-1.617-3.905l-7.002-7h-.003a5.535 5.535 0 0 0-3.9-1.62a5.54 5.54 0 0 0-3.9 1.622v-.002l-7 7c-1.034 1.03-1.62 2.445-1.62 3.904s.586 2.872 1.62 3.9l6.31 6.313c.254.285.52.556.8.8a5.584 5.584 0 0 0 3.91 1.51a5.492 5.492 0 0 0 3.784-1.604l6.812-6.813c.02-.02.035-.03.062-.062l.143-.146c.272-.272.485-.563.726-.86l-.012-.003c.572-.897.89-1.942.89-3.038zM18.77 25.17a3.52 3.52 0 0 1-4.27.514v-.002a7.441 7.441 0 0 1-.61-.47c-.015-.017-7.04-7.042-7.04-7.042c-1.34-1.34-1.34-3.584 0-4.92l7-6.998c1.122-1.12 2.91-1.338 4.26-.512v.002c.213.14.414.3.604.467c.02.015 7.053 7.042 7.053 7.042c.396.388.655.852.818 1.348l-2.607.006a7.745 7.745 0 1 0-7.666 8.85a7.741 7.741 0 0 0 7.67-6.688l2.637-.02c-.16.52-.44 1.02-.85 1.41zm-2.458-8.38h.004l5.476-.02c-.5 2.56-2.76 4.517-5.48 4.52a5.59 5.59 0 0 1-5.584-5.582a5.59 5.59 0 0 1 5.584-5.584a5.59 5.59 0 0 1 5.472 4.484l-5.476.018a1.08 1.08 0 0 0-1.076 1.084c0 .598.483 1.08 1.08 1.08"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Green(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.485 2c0 8-18 4-18 20c0 6 2 8 2 8h2s-3-2-3-8c0-4 9-8 9-8s-7.98 4.328-7.98 8.436C21.238 24.43 28.287 9.606 24.484 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hail(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.372 6.912a7.254 7.254 0 0 0-7.248-7.08a7.238 7.238 0 0 0-6.208 3.518a4.164 4.164 0 0 0-2.01-.518A4.221 4.221 0 0 0 5.69 6.898a5.988 5.988 0 0 0-4.127 5.686c0 3.312 2.687 6 6 6v-.002h15.875c3.312 0 6-2.688 6-6a5.997 5.997 0 0 0-4.066-5.67m-1.936 9.672H7.562a4.007 4.007 0 0 1-4-4c-.003-1.983 1.45-3.62 3.35-3.933a.99.99 0 0 0 .656-.413a1.01 1.01 0 0 0 .163-.762a2.386 2.386 0 0 1-.044-.424a2.224 2.224 0 0 1 2.22-2.218c.647 0 1.217.278 1.633.73a.995.995 0 0 0 .926.31c.342-.065.626-.307.748-.63c.75-1.992 2.662-3.412 4.91-3.41a5.258 5.258 0 0 1 5.252 5.25c0 .16-.01.325-.026.496c-.05.518.305.984.814 1.08c1.86.344 3.273 1.965 3.27 3.922a4.008 4.008 0 0 1-3.998 4.004zM11.503 23.71a1.418 1.418 0 1 1 1.416-1.417a1.422 1.422 0 0 1-1.417 1.416zm7.5 0a1.422 1.422 0 0 1-1.42-1.417a1.417 1.417 0 1 1 1.419 1.416zm-11.5 5.06c-.783 0-1.417-.636-1.417-1.417s.634-1.414 1.417-1.416a1.418 1.418 0 0 1 0 2.834zm7.498 0c-.78 0-1.416-.636-1.416-1.417s.634-1.414 1.417-1.416a1.417 1.417 0 0 1 .001 2.834zm7.5 0a1.415 1.415 0 0 1-1.416-1.415c0-.785.634-1.418 1.416-1.42c.78.002 1.414.635 1.418 1.42a1.42 1.42 0 0 1-1.418 1.416z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hammer(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.83 29.354c.686.353 1.62 1.178 2.345.876c.475-.195.753-1.3 1.048-1.883c2.22-4.376 4.635-9.353 6.392-13.61c0-.19.1-.338-.05-.596c.984-1.6 1.65-3.357 2.725-5.137c.34-.566.686-1.35 1.163-1.577l.88-.368c1.12-.288 1.94-.278 2.72.473c.396.384.578 1.016.96 1.396c.26.26 1.247.9 1.614.8c.285-.077.52-.364.72-.728l.696-1.286c.195-.366.306-.718.215-1c-.117-.36-1.192-.84-1.552-.914c-.528-.113-1.154.08-1.692-.04c-1.057-.244-1.513-.923-1.883-2.02c-2.607-1.534-6.118-2.53-10.206-1.245c-1.11.35-2.172.614-2.9 1.323c-.147.412.142.494.445.49c-.238.215-.62.34-.4.847c2.495-1.146 7.34-1.542 7.67.804c.07.522-.396 1.24-.683 1.835c-.905 1.874-2.01 3.394-2.813 5.09c-.298.018-.366.18-.525.288c-2.605 3.8-5.452 8.54-7.9 12.794c-.327.566-1.1 1.402-1.003 1.906c.144.77 1.33 1.13 2.014 1.484z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hammerandscrewdriver(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28.537 9.86c-.473-.26-1.127-.253-1.61-.524c-.942-.534-1.185-1.316-1.225-2.475c-2.06-2.214-5.138-4.175-9.424-4.113c-1.162.017-2.256-.035-3.158.435c-.258.354-.004.516.288.6c-.29.137-.692.146-.626.696c2.72-.383 7.475.624 7.116 2.966c-.08.52-.735 1.076-1.18 1.563c-1.262 1.382-2.598 2.45-3.76 3.667l.336.336c.742-.52 1.446-.785 2.104-.785c.707 0 1.12.297 1.276.433c.575-.618 1.166-1.244 1.84-1.853c.487-.444 1.046-1.1 1.565-1.178l.948-.1c1.156.046 1.937.29 2.47 1.23c.27.482.263 1.14.522 1.614c.175.325.937 1.22 1.316 1.23c.294.008.603-.2.9-.49l1.032-1.035c.29-.294.5-.6.492-.896c-.006-.38-.9-1.145-1.223-1.32zM13.02 15.352l-.74-.74C9.14 17.254 5.76 20.35 2.75 23.2c-.474.444-1.453 1.022-1.507 1.54c-.083.78.95 1.465 1.506 2c.555.533 1.21 1.602 1.993 1.51c.51-.043 1.095-1.03 1.544-1.502a256.924 256.924 0 0 0 6.883-7.51c-.312-.37-.498-.595-.498-.595c-.137-.192-.893-1.37.35-3.29zm7.62.29c-.366-.318-1.466.143-1.777-.122c-.31-.265.17-1.258-.06-1.454l-.77-.646s-.863-.83-2.813.928l-7.78-7.78c-.395-.395-.237-1.822-.237-1.822l-3.686-2.1l-.894.895l2.1 3.687s1.428-.158 1.824.237l7.792 7.793c-1.55 1.83-.896 2.752-.896 2.752s.238.287.646.77c.196.23 1.188-.25 1.455.06c.264.313-.196 1.41.12 1.778c2.666 3.064 6.926 7.736 8.125 7.736c.892 0 2.02-.724 2.948-1.64c.925-.917 1.64-2.055 1.64-2.947c0-1.2-4.674-5.458-7.738-8.124z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hangup(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28.563 10.494c-7.35-7.35-19.265-7.348-26.612 0c-1.795 1.797-.246 6.84-.246 6.84a.917.917 0 0 0 1.067.615l6.9-1.605c.45-.105.82-.57.82-1.033v-3.685c0-.463.38-.842.842-.842h8.285c.464 0 .843.38.843.842v3.47c0 .463.374.908.83.987l7.634 1.316a.787.787 0 0 0 .926-.692s.51-4.418-1.287-6.214zm-11.3 3.578h-3.5v4.39h-2.625l4.363 7.556l4.364-7.556h-2.6v-4.39z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Help(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.834 4.834c-5.89 5.89-5.89 15.442 0 21.333s15.44 5.888 21.33-.002c5.89-5.89 5.893-15.44.002-21.33c-5.89-5.89-15.44-5.89-21.332 0zm20.625.708a14.164 14.164 0 0 1 2.103 2.726l-4.08 4.08A8.528 8.528 0 0 0 21.57 9.43a8.52 8.52 0 0 0-2.92-1.913l4.08-4.08a14.15 14.15 0 0 1 2.73 2.105zm-15.32 15.32a7.588 7.588 0 0 1-.002-10.725a7.592 7.592 0 0 1 10.725 0a7.595 7.595 0 0 1 0 10.724a7.59 7.59 0 0 1-10.724.002zM5.54 25.46a14.19 14.19 0 0 1-2.105-2.73l4.08-4.08a8.583 8.583 0 0 0 4.832 4.832l-4.082 4.08c-.97-.58-1.89-1.27-2.726-2.103zM8.268 3.434l4.082 4.082a8.558 8.558 0 0 0-4.832 4.831l-4.082-4.08c.58-.97 1.27-1.89 2.105-2.728a14.209 14.209 0 0 1 2.728-2.105zm14.464 24.128L18.65 23.48a8.549 8.549 0 0 0 4.832-4.83l4.082 4.08c-.58.97-1.27 1.892-2.105 2.73a14.197 14.197 0 0 1-2.728 2.103z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func History(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.666 18.292c.275.48.89.644 1.365.367l3.306-1.678c.054.008.108.018.165.018c.83 0 1.5-.67 1.5-1.5l-.5-7.876a1 1 0 1 0-2 0l-.465 7.343l-3.004 1.96a1 1 0 0 0-.364 1.365zm1.396-8.747a.999.999 0 0 0 .366-1.366a.999.999 0 1 0-.366 1.364zM8.18 18.572a1 1 0 1 0 1.365.364a1.004 1.004 0 0 0-1.367-.365zm1-7.876a1 1 0 1 0-.998 1.732a1 1 0 0 0 .997-1.732zM6.624 15.5a1.001 1.001 0 1 0 2.002-.002a1.001 1.001 0 0 0-2.002.002m7.877 7.877a1 1 0 1 0 2.002.002a1 1 0 0 0-2.002-.002m-3.804-1.555c-.275.48-.11 1.09.366 1.365a1 1 0 1 0-.366-1.365m11.126-11.126a1 1 0 1 0 .998 1.732a1 1 0 0 0-.998-1.732m-.366 8.242a1 1 0 1 0 1.732 1a1 1 0 0 0-1.732-1m2.922-3.438c0-.55-.448-1-1-1a1 1 0 1 0 1 1m-5.805 7.322a.999.999 0 1 0 1.73-1a1 1 0 1 0-1.732 1zm.366-13.276a.999.999 0 1 0 .997-1.733a1 1 0 0 0-.998 1.732zm9.763 4.818c-.63-7.292-7.05-12.694-14.34-12.07a13.18 13.18 0 0 0-8.25 3.878l-1.772-1.49l-1.308 7.363l7.03-2.548L8.428 8.12a10.177 10.177 0 0 1 6.194-2.833a10.26 10.26 0 0 1 11.09 9.335a10.26 10.26 0 0 1-9.333 11.09c-5.643.476-10.6-3.694-11.092-9.333c-.102-1.205.02-2.374.31-3.48l-3.27 1.187c-.09.832-.106 1.684-.03 2.55c.628 7.29 7.047 12.69 14.34 12.066c7.29-.63 12.693-7.048 12.068-14.34z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m27.812 16l-3.062-3.062V5.625h-2.625v4.688L16 4.188L4.188 16L7 15.933v11.942h17.875V16zM16 26.167h-5.833v-7H16zm5.667-3h-3.833v-4.042h3.833z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hp(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M29.936 14.968a13.976 13.976 0 0 0-2.3-6.714C26.327 6.3 24.527 4.67 22.43 3.597c-1.617-.853-3.414-1.353-5.23-1.514c-.41-.087-.823.006-1.237.01c-.98 2.81-2 5.605-2.99 8.412c.72.013 1.44-.01 2.16.01c.524.022 1.075.242 1.365.702c.333.52.312 1.185.104 1.748l-2.938 8.253c-.863.013-1.727 0-2.59.007c-.12-.002-.245.012-.36-.033c1.096-3.02 2.167-6.06 3.25-9.083c-.517-.01-1.032.002-1.547-.007c-.08.165-.137.34-.196.513c-.994 2.794-1.997 5.588-2.987 8.384c-.047.066-.035.23-.147.23c-.822.003-1.645-.005-2.467 0c-.148 0-.3.022-.443-.02c2.255-6.256 4.46-12.532 6.716-18.79c-4.407.944-8.2 4.243-9.838 8.432c-.672 1.66-.965 3.45-1.002 5.233c.046.854.106 1.715.273 2.562a13.871 13.871 0 0 0 4.027 7.408a13.97 13.97 0 0 0 6.17 3.452c2.247-6.276 4.48-12.562 6.717-18.85c.02-.046.064-.138.09-.184c1.38.014 2.766.002 4.146.006c.453.007.918-.043 1.354.094c.506.137.98.502 1.14 1.018c.11.407.03.84-.104 1.23c-.887 2.573-1.812 5.124-2.73 7.694c-.228.67-.903 1.09-1.598 1.107c-1.076.02-2.147-.012-3.23.014c-.982 2.762-1.96 5.52-2.944 8.272c.256.103.528.02.794.035c.4.027.802-.043 1.198-.065a13.853 13.853 0 0 0 7.242-2.885a14.1 14.1 0 0 0 3.953-4.875c.812-1.64 1.27-3.438 1.396-5.244c-.06-.642-.044-1.277-.014-1.908zm-9.514 5.063c.916-2.658 1.89-5.296 2.818-7.95c-.398-.008-.8.002-1.197-.005c-.115.017-.268-.042-.354.055c-.944 2.636-1.88 5.276-2.823 7.912c.518.003 1.04.024 1.555-.01z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Icons(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.083 14H14V4.083H4.083zM17 4.083V14h9.917V4.083zm0 22.834h9.917V17H17zm-12.917 0H14V17H4.083z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ie(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27.998 2.266c-2.12-1.91-6.925.382-9.575 1.93a15.3 15.3 0 0 0-2.388-.185c-3.35 0-6.052.986-8.106 2.844c-2.337 2.14-3.632 4.94-3.632 8.177v.085c3.288-5.15 8.343-7.79 9.683-8.487c.213-.1.34.155.142.253c-.015.043-.015 0 0 0c-2.254 1.35-6.434 5.26-9.146 10.887l-.003-.007c-1.717 3.547-3.167 8.53-.267 10.358C6.903 29.5 10.836 27.87 14 25.8c.765.108 1.568.165 2.416.165c5.84 0 9.937-3.223 11.4-7.924l-8.023-.013c-.337 1.66-1.464 2.548-3.223 2.548c-2.21 0-3.73-1.21-3.828-4.012l15.228-.014c.028-.58-.042-.986-.042-1.437c0-5.25-3.143-9.355-8.255-10.663c2.08-1.294 5.974-3.21 7.848-1.68c1.408 1.14.634 3.532.296 4.517c-.056.254.24.296.296.057c.702-1.77.914-4.15-.114-5.078zm-14.726 23.41c-2.47 1.475-5.873 2.54-7.54 1.29c-1.242-.936-.695-3.47.4-5.94a10.91 10.91 0 0 0 2.472 2.63c1.322.995 2.875 1.668 4.668 2.02m-.558-12.63c.042-2.435 1.787-3.49 3.617-3.49c1.93 0 3.49 1.112 3.49 3.49z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IeNine(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27.75 17.887a10.795 10.795 0 0 0-1.064-6.178c1.133-2.886 1.155-5.37-.2-6.778c-1.757-1.822-5.392-1.406-9.434.72l-.206-.002c-6.07 0-10.988 4.888-10.988 10.917c0 .183.005.354.014.53c-2.688 4.07-3.49 7.966-1.688 9.837c1.557 1.613 4.69 1.344 8.2-.392c1.363.604 2.873.938 4.462.938a10.99 10.99 0 0 0 10.37-7.3H21.26c-.814 1.484-2.438 2.505-4.307 2.505c-2.688 0-4.867-2.104-4.867-4.688c0-.036.002-.07.003-.106h15.66zM26.338 6.1c.903.936.806 2.683-.087 4.817a11 11 0 0 0-5.546-4.576c2.54-1.123 4.62-1.293 5.633-.24zm-9.42 4.272c2.522 0 4.585 1.99 4.748 4.51H12.17c.163-2.52 2.226-4.51 4.747-4.51zm-11.23 16.13c-1.103-1.147-.712-3.503.8-6.3a10.958 10.958 0 0 0 5.09 5.94c-2.657 1.226-4.844 1.445-5.89.36"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Imac(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28.936 2.1H2.046a.922.922 0 0 0-.92.92v21.096c0 .506.414.92.92.92h17.062v-.004h9.828c.506 0 .92-.413.92-.92V3.018a.925.925 0 0 0-.92-.92zm-.374 17.962c0 .412-.338.75-.75.75H3.062a.752.752 0 0 1-.75-.75v-16c0-.413.337-.75.75-.75h24.75c.412 0 .75.337.75.75zM20.518 28.4c-.033-.035-.062-.055-.068-.062l-.01-.004l-.008-.004s-.046-.02-.12-.062c-.107-.056-.282-.144-.444-.237a2.435 2.435 0 0 1-.393-.27a.935.935 0 0 1-.058-.17c-.083-.32-.16-.95-.22-1.54h-7.5c-.023.23-.048.468-.076.692a9.512 9.512 0 0 1-.112.716a1.793 1.793 0 0 1-.083.292l-.007.013c-.094.096-.34.246-.553.36c-.107.062-.21.11-.283.146l-.12.06l-.006.005l-.008.004c-.01.008-.038.02-.07.06a.306.306 0 0 0-.067.186c.003.002-.003.037-.005.088c0 .043.007.118.068.185c.06.06.143.08.217.08h9.716a.312.312 0 0 0 .215-.08a.26.26 0 0 0 .07-.186c-.002-.05-.01-.086-.008-.088a.306.306 0 0 0-.064-.186z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Import(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.067 2.25c-5.98 0-11.035 3.91-12.778 9.31H5.5c1.602-3.706 5.27-6.302 9.565-6.31a10.45 10.45 0 0 1 10.437 10.437a10.45 10.45 0 0 1-10.437 10.438c-4.294-.007-7.964-2.605-9.566-6.31H2.29c1.743 5.398 6.798 9.31 12.778 9.31c7.42 0 13.437-6.015 13.438-13.437c-.002-7.423-6.02-13.436-13.44-13.438zm-4.15 17.563l7.15-4.126l-7.15-4.13v2.298H-.056v3.66h10.975v2.298z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inbox(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24 5H8L4 19v8h24v-8zm3 21H5v-7h8.75a2.25 2.25 0 0 0 4.5 0H27z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Info(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 1.466C7.973 1.466 1.466 7.973 1.466 16c0 8.027 6.507 14.534 14.534 14.534c8.027 0 14.534-6.507 14.534-14.534c0-8.027-6.507-14.534-14.534-14.534M14.757 8h2.42v2.574h-2.42zm4.005 15.622H16.1c-1.034 0-1.475-.44-1.475-1.496V15.26c0-.33-.176-.483-.484-.483h-.88V12.4h2.663c1.035 0 1.474.462 1.474 1.496v6.887c0 .31.176.484.484.484h.88z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inkscape(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.402 17.626c.84-.772 2.468-.38 5.98-1.853c1.715-.72 1.57-1.265 1.565-1.928v-.04h-.005a2.56 2.56 0 0 0-.846-1.846L17.05 2.683a3.01 3.01 0 0 0-2.043-.782v-.004c-.738 0-1.475.26-2.04.783l-10.05 9.277a2.555 2.555 0 0 0-.843 1.844H2.07l.004.012c0 .012-.004.024-.004.034h.017c.193.676 5.164 1.536 5.718 2.05c.838.773-3.21 1.338-2.374 2.113c.84.772 5.063 1.495 5.9 2.27c.837.77-1.712 1.596-.875 2.366c.837.773 3.65-.19 3.142 1.822c1.13 1.045 3.49.547 5.07-.498c.838-.77-1.606-.703-.77-1.477c.837-.774 2.95-.777 4.73-2.627c-.714-1.028-3.06-1.468-2.225-2.24zm-9.43-.758l-.42-.504c1.498.358 3.163.827 4.5.837l.057.555c-1.146-.11-3.566-.618-4.137-.887zm7.19-8.288l-1.397-.74l-2.235 1.754l-1.067-3.192l-1.177 2.545l-3.288.303l.036-1.352c0-.324 1.895-2.596 3.05-3.136l2.112-1.4c.312-.187.53-.262.727-.258c.327.01.593.24 1.112.55l4.748 3.25c.357.215.62.522.626.898l-2.814-1.254l-.435 2.032zm8.272 11.014c-.313-.07-1.688-.69-2.035.165c.967.98 2.644 2.18 3.315 1.48c.676-.697-.613-1.495-1.28-1.646zm-21.77 1.972c-.316.074-1.843.115-1.72 1.02c1.35.452 3.438.793 3.684-.112c.245-.908-1.298-1.06-1.965-.908zm13.213 4.83c-.232.21-1.53.953-.93 1.68c1.414-.235 3.404-.913 3.12-1.81c-.28-.897-1.697-.313-2.19.13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Instagram(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 9.904a6.097 6.097 0 0 0 0 12.192a6.097 6.097 0 0 0 0-12.192m0 10.576a4.48 4.48 0 1 1 0-8.96a4.48 4.48 0 0 1 0 8.96m8.576-18.16H7.423a5.157 5.157 0 0 0-5.156 5.158V24.52c0 2.85 2.31 5.156 5.156 5.156h17.153c2.85 0 5.156-2.31 5.156-5.155V7.48c0-2.85-2.307-5.16-5.156-5.16m-2.34 3.074c0-.605.49-1.097 1.097-1.097h3.22c.604 0 1.096.49 1.096 1.097v3.22c0 .604-.49 1.096-1.098 1.096h-3.22c-.603 0-1.096-.49-1.096-1.097zm6.4 19.127c0 2.238-1.82 4.06-4.06 4.06H7.423a4.064 4.064 0 0 1-4.06-4.06V10.88h.827c.102-.5.544-.878 1.075-.878h3.438c.53 0 .974.377 1.074.877h1.498C12.517 9.732 14.175 9.03 16 9.03s3.48.7 4.723 1.847h7.912V24.52h.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ios(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.065 21.797H6.3v-8.784H4.066zm1.12-12.445c-.74 0-1.246.523-1.246 1.208c0 .667.486 1.19 1.226 1.19c.775 0 1.263-.523 1.263-1.19c-.02-.685-.49-1.208-1.246-1.208zm8.73.09c-3.427 0-5.77 2.633-5.77 6.367c0 3.57 2.163 6.185 5.59 6.185c3.355 0 5.826-2.326 5.826-6.402c0-3.463-2.092-6.15-5.645-6.15zM13.86 20.21c-2.128 0-3.373-1.966-3.373-4.437c0-2.507 1.172-4.545 3.373-4.545c2.218 0 3.354 2.164 3.354 4.455c0 2.544-1.207 4.527-3.354 4.527m11.835-5.483c-1.622-.63-2.326-1.064-2.326-2.002c0-.703.612-1.46 2.02-1.46c1.137 0 1.98.342 2.416.576l.54-1.784c-.642-.325-1.603-.613-2.932-.613c-2.633 0-4.293 1.515-4.293 3.5c0 1.75 1.28 2.812 3.283 3.534c1.552.56 2.162 1.1 2.162 2.02c0 .992-.797 1.66-2.227 1.66c-1.137 0-2.22-.36-2.938-.776l-.484 1.84c.668.397 2.002.76 3.28.76c3.14 0 4.616-1.69 4.616-3.645c.023-1.75-1.005-2.814-3.117-3.608z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ipad(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.22 1.417H6.11c-.865 0-1.566.702-1.566 1.566v25.313c0 .865.7 1.565 1.566 1.565h19.11c.866 0 1.566-.7 1.566-1.564V2.984c0-.865-.7-1.567-1.565-1.567zM15.667 29.3a.626.626 0 1 1 0-1.25a.626.626 0 0 1 0 1.249zm8.71-2.445a.313.313 0 0 1-.313.312H7.27a.314.314 0 0 1-.313-.312V4.3c0-.173.14-.313.313-.313h16.792c.172 0 .312.14.312.313z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Iphone(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.755 1H10.62c-1.136 0-2.058.92-2.058 2.058v24.385c0 1.136.92 2.058 2.058 2.058h10.135c1.136 0 2.058-.92 2.058-2.057V3.058A2.058 2.058 0 0 0 20.755 1M14.66 3.264h2.056c.1 0 .183.08.183.18c0 .1-.083.18-.184.18H14.66c-.1 0-.182-.08-.182-.18c0-.1.08-.18.18-.18zm-1.435-.206a.36.36 0 1 1 0 .72a.36.36 0 0 1 0-.72m2.463 25.415a1.44 1.44 0 1 1 0-2.88a1.44 1.44 0 0 1 0 2.88m6.353-4.118a.31.31 0 0 1-.308.31H9.642a.31.31 0 0 1-.308-.31V6.042c0-.17.138-.31.308-.31h12.09c.17 0 .31.14.31.31z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Jigsaw(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 13.62s3.515-4.67 5.59-3.643c2.078 1.027-.413 2.795 1.6 3.72c2.01.923 5.047-.23 4.375-2.9c-.672-2.67-1.866-.776-2.798-2.208C11.573 7.156 17.093 4 17.093 4s3.36 6.65 4.316 4.91c1.156-2.105 3.192-4.265 5.304-1.025c0 0 1.814 2.412.246 3.434s-2.917.442-3.506 1.552c-.586 1.112 3.784 4.093 3.784 4.093s-2.987 4.81-4.926 3.548c-1.94-1.262.356-3.364-2.6-3.99c-1.287-.23-3.437.54-3.817 2.34c-.13 2.71 1.604 2.017 2.797 3.476c1.19 1.456-4.484 4.52-4.484 4.52s-1.584-3.922-3.81-4.656c-2.228-.735-.894 2.135-2.918 2.53c-2.024.397-4.816-2.398-3.46-4.788c1.358-2.39 3.275-.044 3.44-1.95c.17-1.91-3.72-4.377-3.72-4.377z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Jquery(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.322 23.04C4.58 18.724 2.777 11.07 5.494 4.584c-.254.29-.502.59-.74.904C-.42 12.374.748 21.992 7.37 26.97c6.62 4.978 16.184 3.43 21.362-3.456c.237-.314.454-.635.663-.96c-5.478 4.41-13.33 4.803-19.07.487zm3.34-4.442c4.765 3.582 11.604 2.564 15.567-2.198c-3.61 2.64-9.09 2.475-13.362-.736S9.916 7.23 11.45 3.03c-3.474 5.13-2.553 11.985 2.212 15.568m4.98-6.622c3.254 2.447 8.146 1.438 10.967-2.242c-2.605 1.92-6.342 1.955-9.158-.164c-2.82-2.118-3.826-5.718-2.7-8.754c-2.754 3.733-2.365 8.712.89 11.16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Js(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 1v30h30V1zm16.326 23.398c0 2.92-1.712 4.248-4.21 4.248c-2.254 0-3.563-1.168-4.228-2.576l2.296-1.39c.443.776.846 1.44 1.812 1.44c.926 0 1.51-.353 1.51-1.77v-9.58h2.82zm6.666 4.248c-2.618 0-4.31-1.248-5.135-2.88l2.295-1.327c.604.978 1.39 1.71 2.78 1.71c1.167 0 1.903-.584 1.903-1.396c0-.966-.766-1.31-2.054-1.865l-.7-.31c-2.034-.865-3.383-1.953-3.383-4.25c0-2.113 1.604-3.725 4.128-3.725c1.792 0 3.08.625 4.008 2.254l-2.19 1.405c-.48-.86-1.006-1.21-1.812-1.21c-.824 0-1.352.522-1.352 1.21c0 .852.52 1.188 1.73 1.71l.703.31c2.397 1.02 3.747 2.07 3.747 4.43c.007 2.544-1.986 3.93-4.665 3.93z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Key(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.386 16.01l.01-.007l-.58-.912c1.653-2.225 1.875-5.318.3-7.8a6.897 6.897 0 0 0-11.638 7.4a6.9 6.9 0 0 0 6.943 3.102l.424.67l.206.044l.78-.447l-.306 1.376l2.483.552l-.296 1.325l1.903.424l-.68 3.06l1.406.313l-.424 1.906l4.134.918l.758-3.392l-5.426-8.533zm-7.39-7.066a1.471 1.471 0 0 1-1.578-2.482a1.47 1.47 0 0 1 1.578 2.483z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lab(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m22.12 24.438l-3.36-7.847c-.33-.768-.6-2.08-.6-2.916s.513-1.52 1.14-1.52s1.14-.514 1.14-1.14s-.684-1.14-1.52-1.14h-6.84c-.836 0-1.52.512-1.52 1.14s.513 1.14 1.14 1.14s1.14.684 1.14 1.52s-.27 2.148-.6 2.917L8.88 24.44c-.33.768-.6 1.74-.6 2.157s.342 1.103.76 1.52s1.444.76 2.28.76h8.36c.835 0 1.86-.34 2.28-.76s.76-1.102.76-1.52s-.27-1.39-.6-2.157zM16.583 7.625a1.082 1.082 0 1 0 2.166 0a1.083 1.083 0 1 0-2.166.001zm-2.915.167a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1m1.917-2.5a1.584 1.584 0 1 0-.002-3.164a1.584 1.584 0 0 0 .002 3.162z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lamp(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 2.833a7 7 0 0 0-7 7c0 3.86 3.945 4.937 4.223 9.5h5.553c.278-4.563 4.224-5.64 4.224-9.5a7 7 0 0 0-7-7m0 25.333c1.894 0 2.483-1.027 2.667-1.666h-5.334c.184.64.773 1.666 2.667 1.666m-2.75-2.668h5.5v-5.164h-5.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LampAlt(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.75 25.498h5.5v-5.164h-5.5zm2.75 2.668c1.894 0 2.483-1.027 2.667-1.666h-5.334c.184.64.773 1.666 2.667 1.666m0-25.333a7 7 0 0 0-7 7c0 3.86 3.945 4.937 4.223 9.5h1.27c-.008-.026-.023-.05-.028-.08l-2-10.997a.502.502 0 0 1 .94-.313l.553 1.106l.553-1.107a.5.5 0 0 1 .893 0l.553 1.106l.553-1.107a.498.498 0 0 1 .893 0l.554 1.106l.553-1.107a.5.5 0 0 1 .937.313l-2 10.998c-.004.03-.02.053-.028.078h1.355c.278-4.562 4.224-5.64 4.224-9.5c0-3.864-3.134-7-7-7zm1.958 7.833a.496.496 0 0 1-.446-.275l-.554-1.105l-.553 1.106a.503.503 0 0 1-.896.001l-.553-1.105l-.553 1.106a.496.496 0 0 1-.446.276h-.037l1.455 8h1.167l1.454-8h-.038z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Landing(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.322 19.49s1.903.343.3-1.868c-1.354-1.866-5.262-3.104-5.262-3.104l-4.213-8.23l-2.47-.393l.973 5.45l-3.41-1.235l-.468-2.837l-1.764-.97s-.496 2.74-.15 5.27c0 0 7.107 6.426 16.464 7.918zm-20.07 3.616v1.998h24.497v-1.998H3.25zM14 17.94a.75.75 0 1 0 1.5 0a.75.75 0 0 0-1.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LandscapeOne(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.883 5.71H2.746c-.163 0-.32.07-.435.188a.611.611 0 0 0-.18.435v18.364c0 .164.064.318.18.436c.124.117.288.18.436.18h25.75a.624.624 0 0 0 .438-.18a.606.606 0 0 0 .182-.436V14.55c-.002-.1-.01-.187-.02-.27c-.187-1.543-1.544-3.424-3.237-5.168c-1.82-1.802-3.99-3.36-5.977-3.402m7.03 6.604c-.007-.005-.018-.007-.028-.01c-1.092-.293-2.33-.355-3.2-.355c-.16 0-.31 0-.444.003a17.29 17.29 0 0 0-.355-2.625c-.11-.46-.246-.94-.433-1.42c.857.54 1.748 1.264 2.535 2.068a14.91 14.91 0 0 1 1.927 2.338zM3.366 6.947h16.517c.058 0 .12 0 .183.004c.694.106 1.307 1.222 1.616 2.647c.336 1.484.355 2.997.355 3l.007.656l.65-.05s.4-.028.99-.026c.81 0 1.978.062 2.872.312c.94.274 1.352.634 1.326 1.05h.002v9.542H3.365V6.947z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LandscapeTwo(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.883 5.71H2.746c-.163 0-.32.07-.435.188a.611.611 0 0 0-.18.435v18.364c0 .164.064.318.18.436c.124.117.288.18.436.18h25.75a.624.624 0 0 0 .438-.18a.606.606 0 0 0 .182-.436V14.55c-.002-.1-.01-.187-.02-.27c-.187-1.543-1.544-3.424-3.237-5.168c-1.82-1.802-3.99-3.36-5.977-3.402M3.365 6.947h16.517c.058 0 .12 0 .183.004c.694.106 1.307 1.222 1.616 2.647c.336 1.484.355 2.997.355 3l.007.656l.65-.05s.4-.028.99-.026c.81 0 1.978.062 2.872.312c.94.274 1.352.634 1.326 1.05h.002v9.542H3.365z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linechart(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.625 25.062a1 1 0 0 1-.77-1.187L6.51 6.585l2.267 9.258l1.923-5.188l3.58 3.74l3.884-13.102l2.934 11.734l1.96-1.51l5.27 11.74a1 1 0 1 1-1.826.817l-4.23-9.428l-2.374 1.826l-1.896-7.596l-2.783 9.393l-3.755-3.924l-3.08 8.314l-1.73-7.083l-1.843 8.71a1 1 0 0 1-1.187.775z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16.45 18.085l-2.47 2.47a3.729 3.729 0 0 1-1.078 2.847A3.755 3.755 0 0 1 7.6 23.4a3.754 3.754 0 0 1 0-5.3a3.731 3.731 0 0 1 2.846-1.08l2.47-2.468a6.753 6.753 0 0 0-7.44 1.426a6.75 6.75 0 0 0 9.546 9.545v.002a6.76 6.76 0 0 0 1.428-7.44m-1.898-5.17l2.467-2.47a3.735 3.735 0 0 1 1.077-2.847a3.756 3.756 0 0 1 5.303.002a3.75 3.75 0 0 1 0 5.3a3.72 3.72 0 0 1-2.846 1.08l-2.47 2.468a6.76 6.76 0 0 0 7.44-1.424a6.75 6.75 0 1 0-9.546-9.546a6.75 6.75 0 0 0-1.426 7.437zm3.6-2.188l-7.424 7.426a1.498 1.498 0 1 0 2.12 2.12l7.426-7.425a1.5 1.5 0 1 0-2.122-2.12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linkedin(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27.25 3.125h-22a2 2 0 0 0-2 2v22a2 2 0 0 0 2 2h22a2 2 0 0 0 2-2v-22a2 2 0 0 0-2-2M11.22 26.78h-4v-14h4zm-2-15.5c-1.384 0-2.5-1.118-2.5-2.5s1.116-2.5 2.5-2.5s2.5 1.12 2.5 2.5s-1.118 2.5-2.5 2.5m16 15.5h-4v-8.5c0-.4-.404-1.054-.688-1.212c-.375-.21-1.26-.23-1.665-.034l-1.648.793v8.954h-4v-14h4v.615c1.582-.723 3.78-.652 5.27.184c1.58.885 2.73 2.863 2.73 4.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linux(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.79 25.23a.988.988 0 0 0 .68-1.667c-.397-.392-2.597-2.022-3.17-2.664c-.267-.3-.627-.446-.78-.78c-.352-.77-.598-1.87-.15-2.658c.08-.144.132-.078.07.22c-.35 1.684.746 3.06.986 2.354c.167-.487.013-1.358.102-2.05c.158-1.227 1.273-3.578 1.763-3.713c-.754-1.398.887-2.494.867-3.723c-.014-.798.7.982 1.42 1.36c.8.42 1.683-.795 2.935-1.41c.354-.177.81-.377.776-.525c-.146-.718-1.644.886-2.98.94c-.61.023-.836-.12-1.07-.348c-.713-.69.072-.115 1.13-.307c.472-.085.63-.163 1.13-.365c.5-.2 1.068-.5 1.635-.654c.395-.105.36-.4.208-.49c-.087-.05-.218-.046-.32.133c-.244.42-1.383.66-1.74.77c-.457.14-.962.272-1.634.244c-1.02-.042-.782-.51-1.513-.928c-.213-.123-.156-.445.13-.73c.147-.148.556-.232.76-.572c.027-.047.288-.32.493-.46c.07-.05.076-1.296-.562-1.32c-.543-.022-.697.397-.675.817c.02.42.244.765.392.764c.285-.005.02.31-.138.36c-.237.078-.562-.934-.525-1.418c.038-.506.302-1.4.94-1.383c.577.015.994.736.974 1.982c-.003.21.935-.1 1.247.23c.225.235-.766-2.208 1.44-2.376c.58.11 1.14.305 1.37 1.64c-.086.14.146 1.07-.215 1.183c-.438.134-.707-.02-.453-.44c.172-.417.004-1.482-.882-1.42c-.887.065-.77 1.638-.526 1.67c.243.03.854.464 1.282.548c1.4.27.37 1.075.555 2.048c.206 1.1.93.81 1.58 3.717c.136.177.675.345 1.198 2.58c.473 2.01-.195 3.472.938 3.352c.256-.026.63-.1.792-.668c.425-1.49-.213-3.263-.855-4.46c-.375-.698-.73-1.174-.916-1.337c.738.436 1.683 1.83 1.898 2.862c.286 1.358.49 1.934.06 3.37c.25.125.87.39.87.685c-.647-.53-2.63-.625-2.68.646c-.338.008-.594.034-.81.293c-.798.944-.06 2.842-.14 3.86c-.07.895-.318 1.782-.46 2.682c-.474-.02-.428-.364-.274-.852c.136-.43.352-.968.366-1.484c.012-.467-.04-.76-.156-.83c-.118-.073-.303.073-.56.484c-.542.875-1.72 1.26-2.82 1.397c-1.1.138-2.123.028-2.664-.578c-.186-.207-.492.058-.53.11c-.048.075.18.22.353.534c.25.46.49 1.158-.106 1.478c.006-1.624-.508-1.717-1.033-2.71zm-.392-.042c.395.62 1.783 3.232-.652 3.57c-.814.115-2.125-.473-3.396-.783c-1.142-.28-2.3-.444-2.95-.627c-.39-.108-.553-.25-.587-.414c-.09-.434.474-1.04.503-1.555c.028-.515-.188-.78-.364-1.2c-.177-.42-.224-.734-.08-.914c.108-.14.333-.2.697-.164c.46.047 1.02-.05 1.318-.23c.505-.31.742-.94.516-1.7c0 .745-.244 1.026-.855 1.367c-.578.318-1.468.06-1.876.415c-.492.427.175 1.528.12 2.338c-.042.622-.69 1.322-.4 1.946c.29.626 1.647.694 3.063.99c2.012.42 3.184 1.152 4.113 1.187c1.356.05 1.564-1.342 3.693-1.36a45.87 45.87 0 0 1 1.835-.06c.688-.01 1.375-.003 2.08.014c1.416.035.93.774 1.85 1.247c.774.397 2.17.24 2.504-.077c.45-.43 1.662-1.467 2.592-1.935c1.156-.583 3.876-1.588 1.902-2.812c-.46-.285-1.547-.588-1.64-2.676c-.41.366-.364 2.312.785 2.697c1.284.43 2.086 1.152-.3 1.97c-1.58.54-1.85.705-3.1 1.746c-1.266 1.054-3.144.636-2.814-1.582c.17-1.155.27-2.11-.02-3.114c-.14-.49-.21-1.12-.113-1.562c.187-.858.65-1.117 1.106-.293c.284.518.384 1.12 1.407 1.17c1.607.077 1.926-1.553 2.44-1.627c.342-.05.685-1.02.424-2.59c-.28-1.68-1.268-4.33-2.535-5.676C21.612 11.78 20.948 10.8 20.53 9.4c-.352-1.175-.547-2.318-.475-3.412c.094-1.417-.69-3.39-1.943-4.316c-.782-.58-2.01-.893-3.122-.88c-.623.007-1.21.1-1.66.343c-1.856 1.008-2.114 2.445-2.087 4.088c.025 1.543.078 3.303.254 4.977c-.208.77-1.288 2.227-1.98 3.114c-.927.92-1.396 2.696-1.997 4.247c-.32.83-.862 1.2-.908 2.266c-.012.296-.002 1.065.282.846c1.086-.843 2.45 1.278 4.504 4.516zm5.646-22.235c-.06.176-.3.32-.146.443c.152.123.24-.17.55-.28c.08-.03.448.01.518-.165c.03-.076-.19-.163-.32-.29c-.134-.125-.263-.236-.387-.23c-.322.02-.164.368-.216.523zm1.89 6.397c.115-.12.174.207.483.402c.244.154.48.04.545.354c.044.225-.097.467-.284.436c-.328-.056-1.082-.837-.744-1.192m-5.102-1.975c-.508-.037-.543.33-.375.324c.172-.007.066-.292.375-.325zm-.872-.94c.06-.012.146.09.12.234c-.038.198-.022.323.116.324c.022 0 .048-.005.056-.057c.066-.396-.14-.688-.225-.71c-.193-.05-.17.23-.067.21zm3.703-.167c.13.04.253.262.28.504c.002.02.168-.035.17-.088c.01-.39-.32-.57-.408-.562c-.2.017-.143.116-.042.146m-1.898 1.155c.463-.214.625.118.465.17c-.164.055-.165-.248-.465-.17m-5.587 7.88c-.22-.025.063-.19.184-.396c.13-.227.105-.51.244-.47s.06.2-.033.462c-.082.22-.315.413-.395.404"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func List(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.082 4.083v3h22.835v-3zm0 16.223h22.835v-3H4.082zm0-6.612h22.835v-3H4.082zm0 13.223h22.835v-3H4.082z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LoactionTwo(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.834 29.084V16.166H2.917l26.166-13.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Location(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3.5A7.5 7.5 0 0 0 8.5 11c0 4.143 7.5 18.12 7.5 18.12S23.5 15.144 23.5 11A7.5 7.5 0 0 0 16 3.5m0 11.084a3.583 3.583 0 1 1 0-7.168a3.583 3.583 0 1 1 0 7.168"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.875 15.334v-4.876c0-4.894-3.98-8.875-8.875-8.875s-8.875 3.98-8.875 8.875v4.876H5.042v15.083h21.916V15.334zm-14.25-4.876c0-2.964 2.41-5.375 5.375-5.375s5.375 2.41 5.375 5.375v4.876h-10.75zm7.647 16.498h-4.545l1.222-3.667a2.37 2.37 0 0 1-1.325-2.12a2.375 2.375 0 1 1 4.75 0c0 .932-.542 1.73-1.324 2.12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Locked(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M26.71 14.085L16.915 4.29a2.007 2.007 0 0 0-2.83 0L4.29 14.086a2.007 2.007 0 0 0 0 2.83l9.796 9.795c.778.778 2.05.778 2.83 0l9.796-9.796a2.011 2.011 0 0 0 0-2.83zm-18.49 2.34c-.4-.154-.687-.534-.687-.988s.288-.834.688-.987v1.974zm.75 0V14.45c.4.152.688.533.688.987s-.287.835-.687.987zm4.532 2.763l1.203-3.61a1.963 1.963 0 0 1-1.172-1.796a1.97 1.97 0 0 1 3.94 0c0 .803-.483 1.49-1.173 1.797l1.203 3.608h-4zm8.53-2.64c-.4-.15-.687-.532-.687-.985s.287-.834.687-.987v1.973zm.75 0v-1.972c.4.152.69.533.69.987s-.288.834-.69.986z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LowBattery(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27.27 13.5h-1.103v-1.416a1 1 0 0 0-1-1H5.25a1 1 0 0 0-1 1v7.832a1 1 0 0 0 1 1h19.917a1 1 0 0 0 1-1V18.5h1.104c.266 0 .48-.448.48-1v-3c0-.552-.214-1-.48-1zm-3.103 5.416H6.25v-5.832h17.917zm-15-4.832H7.25v3.832h1.917z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Magic(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m23.043 4.65l-.404-2.313l-1.59 1.727l-2.324-.33l1.15 2.045l-1.03 2.107l2.3-.463l1.687 1.633l.27-2.332l2.075-1.1l-2.135-.976zm3.174 13.548l-.182-1.25l-.882.905l-1.245-.214l.588 1.117l-.588 1.118l1.245-.214l.882.906l.182-1.25l1.133-.56zM4.92 7.672l.948-.372l.844.57l-.062-1.017l.802-.627l-.985-.256l-.35-.957l-.546.86l-1.017.035l.647.786l-.28.978zm5.52 2.833l1.02-1.096l1.48.218l-.726-1.31l.667-1.34l-1.47.286l-1.068-1.048l-.182 1.487l-1.328.693l1.358.632l.25 1.477zm6.794 2.216c-.588-.367-1.172-.617-1.692-.728c-.492-.09-1.04-.15-1.425.374L2.562 30.788h6.68l9.67-15.416c.302-.576.01-1.04-.284-1.447c-.325-.417-.806-.835-1.394-1.204m-3.62 9.216c-.255-.396-.74-.857-1.374-1.254c-.632-.396-1.258-.634-1.726-.69l4.42-7.052c.065-.013.263-.02.544.066c.346.092.785.285 1.225.562c.504.313.908.677 1.133.97a1.12 1.12 0 0 1 .2.35z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Magnet(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.812 19.5h5.002v-6.867c-.028-1.706-.61-3.807-2.172-5.84c-1.54-2.015-4.315-3.72-7.94-3.688c-3.626-.032-6.402 1.674-7.94 3.687c-1.562 2.034-2.145 4.136-2.174 5.842V19.5h5v-6.866c-.027-.377.303-1.79 1.1-2.748c.818-.98 1.847-1.747 4.013-1.778c2.166.032 3.196.8 4.014 1.778c.798.96 1.126 2.372 1.1 2.748V19.5zm5.002 6.08V20.5h-5.002v5.08h5.004zm-20.226 0h5V20.5h-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mail(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28.516 7.167H3.482L16 14.275zM16.74 17.303a1.494 1.494 0 0 1-1.48 0L2.5 10.06v14.773h27V10.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Man(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.02 16.35c-.61-1.105-1.358-2-2.108-2.624a5.233 5.233 0 0 1-3.103 1.03a5.236 5.236 0 0 1-3.106-1.03c-.75.625-1.498 1.52-2.11 2.623c-1.423 2.562-1.58 5.19-.35 5.873c.55.307 1.126.078 1.722-.496a10.585 10.585 0 0 0-.166 1.873c0 2.932 1.14 5.307 2.543 5.307c.846 0 1.265-.865 1.466-2.19c.2 1.325.62 2.19 1.462 2.19c1.406 0 2.545-2.375 2.545-5.307c0-.66-.06-1.29-.168-1.873c.597.574 1.173.803 1.724.496c1.228-.682 1.07-3.31-.353-5.874zm-5.212-2.593a4.28 4.28 0 1 0-.003-8.56a4.28 4.28 0 0 0 .003 8.561z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Merge(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m29.342 15.5l-7.556-4.363v2.613h-1.41c-.79-.01-1.332-.24-2.02-.743c-1.02-.745-2.094-2.18-3.55-3.568c-1.44-1.38-3.515-2.71-6.306-2.69H2.812v3.5H8.5c2.23.01 3.44 1.184 5.07 2.933c.697.753 1.428 1.58 2.324 2.323c-1.396 1.165-2.412 2.516-3.484 3.5c-1.183 1.082-2.202 1.724-3.912 1.742H2.813v3.5H8.53c3.75 0 6.034-2.32 7.618-4.066c.817-.895 1.537-1.69 2.21-2.19c.685-.503 1.23-.733 2.016-.743h1.412v2.613l7.556-4.363z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mic(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 21.125c2.682 0 4.875-2.25 4.875-5V5.875c0-2.75-2.193-5-4.875-5s-4.875 2.25-4.875 5v10.25c0 2.75 2.193 5 4.875 5M21.376 11v5.125c0 3.308-2.636 6-5.876 6s-5.875-2.69-5.875-6V11h-3v5.125c0 4.443 3.195 8.132 7.373 8.86v2.14h-3.372v3h9.75v-3h-3.377v-2.14c4.18-.726 7.374-4.417 7.374-8.86V11h-2.998z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Micmute(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.12 18.53a6.09 6.09 0 0 1-.495-2.405V11h-3v5.125c0 1.693.467 3.275 1.273 4.627l2.223-2.223zM20.376 8.275v-2.4c0-2.75-2.193-5-4.875-5s-4.875 2.25-4.875 5v10.25c0 .567.113 1.104.285 1.614zm1 4.655v3.196c0 3.308-2.635 6-5.875 6c-.958 0-1.86-.24-2.66-.657l-2.18 2.178a8.758 8.758 0 0 0 3.338 1.34v2.14h-3.372v3h9.75v-3H17v-2.14c4.18-.727 7.374-4.418 7.374-8.86V11h-1.068zm-1 3.195V13.93l-6.788 6.79c.588.26 1.234.404 1.913.404c2.682 0 4.875-2.25 4.875-5zm5.167-11.603L4.855 25.21l1.415 1.415L26.956 5.937l-1.414-1.415z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.98 12.896h-20v6.666h20z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Music(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.75 8.042v15.903c-.676-.45-1.635-.736-2.707-.736c-2.048 0-3.708 1.024-3.708 2.29s1.66 2.292 3.708 2.292s3.708-1.026 3.708-2.29V13.785l10.916-3.24v9.565c-.678-.45-1.635-.735-2.708-.735c-2.048 0-3.708 1.025-3.708 2.292c0 1.265 1.66 2.29 3.708 2.29s3.708-1.025 3.708-2.29V4.207z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mute(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.328 8.956c-.605-1.545-1.4-2.81-2.572-3.484a2.525 2.525 0 0 0-2.122-3.892a2.525 2.525 0 0 0-2.121 3.893c-3.973 2.287-3.657 11.26-6.142 15.485c-2.15.565-3.485 1.353-3.485 2.222v1.125c0 .27.117.525.322.774zm-4.693-6.378a1.53 1.53 0 0 1 1.527 1.527c0 .372-.14.708-.36.974a4.432 4.432 0 0 0-1.165-.15c-.422 0-.808.054-1.167.15a1.526 1.526 0 0 1-.36-.975c0-.842.682-1.525 1.525-1.527zm8.263 18.38c-1.125-1.914-1.678-4.802-2.312-7.602L9.066 26.878c1.394.338 3.06.587 4.895.71c-.002.048-.014.096-.014.145a2.689 2.689 0 0 0 5.376.001c0-.05-.012-.097-.016-.145c5.188-.35 9.062-1.688 9.062-3.283V23.18c.017-.87-1.318-1.657-3.47-2.222zm1.78-15.937L3.614 28.084L5.03 29.5L28.09 6.435L26.678 5.02z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Newwindow(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.896 5.333V21.25h23.417V5.333zM26.312 18.25H8.896V8.334h17.417v9.916zM4.896 9.542h-3.21V25.46h23.418v-3.21H4.896z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func No(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3.667C9.19 3.667 3.667 9.187 3.667 16S9.19 28.333 16 28.333c6.812 0 12.333-5.52 12.333-12.333S22.813 3.667 16 3.667m0 3c1.85 0 3.572.548 5.024 1.48L8.147 21.024A9.263 9.263 0 0 1 6.667 16c0-5.146 4.187-9.333 9.333-9.333m0 18.666a9.271 9.271 0 0 1-5.024-1.48l12.876-12.877A9.263 9.263 0 0 1 25.332 16c0 5.146-4.186 9.333-9.332 9.333"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Nodejs(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.783 4.458L2.59 3.196a.264.264 0 0 0-.12-.035h-.023a.266.266 0 0 0-.12.036L.132 4.458a.265.265 0 0 0-.132.23l.005 3.398c0 .047.024.092.066.114a.132.132 0 0 0 .133 0l1.303-.746a.267.267 0 0 0 .132-.23V5.638c0-.095.05-.183.132-.23l.555-.32a.269.269 0 0 1 .264.001l.554.32a.265.265 0 0 1 .133.23v1.587c0 .094.05.182.132.23L4.71 8.2a.132.132 0 0 0 .198-.114l.004-3.398a.267.267 0 0 0-.13-.23zM17.93.745l-1.305-.73a.14.14 0 0 0-.132.002a.133.133 0 0 0-.065.114v3.366a.093.093 0 0 1-.139.08l-.55-.315a.272.272 0 0 0-.265 0l-2.193 1.267a.265.265 0 0 0-.133.23V7.29c0 .095.05.182.13.23l2.195 1.266a.262.262 0 0 0 .265 0l2.194-1.267a.27.27 0 0 0 .133-.23V.977a.266.266 0 0 0-.137-.232zm-1.51 5.713a.067.067 0 0 1-.032.057l-.753.435a.065.065 0 0 1-.066 0l-.754-.435a.068.068 0 0 1-.033-.057v-.87a.07.07 0 0 1 .033-.058l.753-.435a.063.063 0 0 1 .065 0l.753.435c.02.012.033.034.033.058v.87zm8.053-1.965l-2.18-1.266a.264.264 0 0 0-.265 0l-2.193 1.266a.265.265 0 0 0-.132.23v2.53c0 .097.05.185.133.23l2.18 1.243c.08.045.18.046.26 0l1.318-.73a.134.134 0 0 0 .002-.232L21.39 6.498a.13.13 0 0 1-.067-.115V5.59c0-.047.025-.09.065-.115l.688-.396a.126.126 0 0 1 .132 0l.688.395c.04.023.066.067.066.115v.625a.132.132 0 0 0 .198.114l1.314-.765a.269.269 0 0 0 .132-.23V4.72c0-.094-.05-.18-.133-.23zm-13.11-.013L9.17 3.214a.269.269 0 0 0-.266 0L6.71 4.48a.267.267 0 0 0-.13.23v2.533c0 .095.05.182.132.23l2.193 1.266a.267.267 0 0 0 .265 0l2.193-1.268a.266.266 0 0 0 .132-.23V4.71a.265.265 0 0 0-.132-.23m19.656-.098l-2.07-1.195a.454.454 0 0 0-.435 0l-2.068 1.195a.435.435 0 0 0-.218.377v2.385c0 .156.082.3.217.378l.542.312c.262.13.355.13.476.13c.39 0 .612-.236.612-.646V4.96a.06.06 0 0 0-.06-.06h-.263a.06.06 0 0 0-.06.06v2.356c0 .182-.19.363-.496.21l-.567-.327a.062.062 0 0 1-.033-.056V4.76c0-.024.013-.046.033-.058L28.7 3.51a.06.06 0 0 1 .063 0L30.83 4.7c.02.012.032.034.032.057v2.385a.068.068 0 0 1-.032.057l-2.068 1.193a.06.06 0 0 1-.063 0l-.53-.314c-.018-.01-.037-.012-.053-.003a1.447 1.447 0 0 1-.312.143c-.034.012-.084.03.02.09l.69.408c.065.038.14.06.217.06s.15-.022.218-.06l2.068-1.194a.437.437 0 0 0 .217-.378V4.76c0-.156-.083-.3-.216-.378zm-1.65 2.386c-.547 0-.667-.138-.707-.41a.06.06 0 0 0-.06-.05h-.268a.059.059 0 0 0-.06.06c0 .348.19.764 1.095.764c.655 0 1.03-.26 1.03-.71c0-.446-.3-.565-.937-.65c-.643-.084-.708-.127-.708-.278c0-.125.056-.29.53-.29c.426 0 .582.09.647.378a.06.06 0 0 0 .06.047h.268a.062.062 0 0 0 .044-.02a.072.072 0 0 0 .016-.046c-.042-.493-.37-.723-1.032-.723c-.59 0-.94.25-.94.667c0 .453.35.578.915.634c.677.067.73.167.73.3c0 .23-.186.328-.622.328zM22.13 5.446l-.42.243a.047.047 0 0 0-.025.043v.486c0 .018.01.034.025.043l.42.243c.016.01.035.01.052 0l.42-.243a.049.049 0 0 0 .026-.044V5.73a.052.052 0 0 0-.025-.044l-.42-.244a.06.06 0 0 0-.052 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Nomagnet(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.59 17.857v-5.225c-.027-.376.303-1.79 1.1-2.748c.818-.98 1.848-1.748 4.013-1.778c1.704.026 2.7.508 3.447 1.19l3.54-3.54c-1.617-1.526-4.01-2.68-6.987-2.652c-3.626-.03-6.403 1.675-7.94 3.69c-1.563 2.032-2.146 4.134-2.173 5.84V19.5h3.357zm-5 2.643v2.357L7.947 20.5zm15.222-7.21v6.21h5.002v-6.866a9.314 9.314 0 0 0-.803-3.542zm4.527-8.768L4.65 25.21l1.415 1.415L26.753 5.937L25.34 4.522zM20.81 25.58h5.002V20.5H20.81zm-10.222 0v-2.064L8.525 25.58h2.065z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notebook(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.875 1.375H8a1.995 1.995 0 0 0-1.98 1.792h1.605c1.102 0 2 .898 2 2c0 1.102-.898 2-2 2H6v1h1.625c1.104 0 2.002.897 2.002 2a2.004 2.004 0 0 1-2.002 2.002H6v.996h1.625c1.102 0 2 .898 2 2a2.005 2.005 0 0 1-2 2.004H6v.994h1.625c1.102 0 2 .898 2 2.002s-.898 2.002-2 2.002H6v.997h1.624c1.104 0 2.002.897 2.002 2a2.004 2.004 0 0 1-2.002 2.003h-1.62A1.998 1.998 0 0 0 8 29.124h16.875a2 2 0 0 0 2-2V3.375a2 2 0 0 0-2-2m.375 7a1 1 0 0 1-1 1H14a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h10.25a1 1 0 0 1 1 1zM8.625 25.165c0-.553-.45-1-1-1h-3.25a1 1 0 1 0 0 2h3.25c.55 0 1-.447 1-1m-4.25-19h3.25a1 1 0 1 0 0-1.998h-3.25a1.001 1.001 0 0 0 0 2zm0 5.002h3.25a1 1 0 1 0 0-2h-3.25a1 1 0 1 0 0 2m0 5h3.25a1 1 0 0 0 0-2h-3.25c-.553 0-1 .446-1 1s.447 1 1 1m-1 3.998a1 1 0 0 0 1 1.002h3.25a1 1 0 0 0 0-2.002h-3.25a1 1 0 0 0-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Noview(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.478 17.568A4.752 4.752 0 0 1 11.193 16A4.814 4.814 0 0 1 16 11.193c.552 0 1.074.113 1.568.285l2.283-2.283c-1.31-.548-2.623-.91-3.85-.91C8.454 8.286 2.5 16 2.5 16s2.167 2.79 5.53 5.017l3.448-3.45zm12.04-6.383l-3.056 3.056c.217.547.345 1.14.345 1.76A4.815 4.815 0 0 1 16 20.809a4.757 4.757 0 0 1-1.76-.345l-2.47 2.47c1.328.48 2.746.783 4.23.783c5.77 0 13.5-7.715 13.5-7.715s-2.64-2.626-5.982-4.815zm2.024-6.268L4.855 25.604L6.27 27.02L26.956 6.332z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Opensource(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 1.125C7.278 1.125.59 7.815.59 16.035c0 6.263 3.88 11.635 9.36 13.84l3.64-9.076a5.131 5.131 0 1 1 3.818-.001l3.64 9.075c5.48-2.206 9.36-7.578 9.36-13.84c.002-8.22-6.687-14.91-14.908-14.91z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Opera(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.954 2.046c-7.49 0-12.872 5.432-12.872 13.58c0 7.25 5.234 13.836 12.873 13.836c7.712 0 12.974-6.583 12.974-13.835c0-8.214-5.555-13.58-12.976-13.58zm-.002 24.502c-2.29 0-3.49-1.61-4.12-3.796c-.285-1.037-.46-2.185-.564-3.34c-.114-1.375-.13-2.774-.13-4.03c0-.992.02-1.978.075-2.925c.124-1.728.386-3.43.89-4.833c.694-1.718 1.87-2.822 3.85-2.822c2.5 0 3.762 1.782 4.384 4.322c.43 1.894.56 4.124.56 6.274c0 2.3-.103 5.153-.763 7.442c-.66 2.14-1.892 3.708-4.182 3.708"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Package(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m17.078 22.004l-1.758-4.13l-2.007 4.753l-7.52-3.29l.175 3.906l9.437 4.374l10.91-5.365l-.15-4.99zM29.454 6.62L18.52 3.382l-3.005 2.67l-3.09-2.358L1.544 8.2l3.796 3.047l-3.43 5.303l10.88 4.756l2.53-5.998l2.256 5.308l11.393-5.942l-3.105-4.71l3.592-3.345zm-14.177 7.96l-9.06-3.83l9.276-4.102L25.1 9.903z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Page(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.024 5.673C21.28 3.98 19.4 2.623 17.856 2.437a1.835 1.835 0 0 0-.263-.02H7.438a.606.606 0 0 0-.436.18a.624.624 0 0 0-.18.436v25.75a.615.615 0 0 0 .616.615h18.364a.61.61 0 0 0 .434-.18a.613.613 0 0 0 .182-.435V11.648c-.036-1.99-1.594-4.158-3.394-5.975m2.16 22.49H8.052V3.647h9.542v.002c.416-.025.775.386 1.05 1.326c.25.895.313 2.062.312 2.87c.002.594-.027.992-.027.992l-.05.652l.656.007c.003 0 1.516.018 3 .355c1.426.308 2.54.922 2.645 1.617c.005.062.006.124.005.182z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageTwo(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.024 5.673C21.28 3.98 19.4 2.623 17.856 2.437a1.835 1.835 0 0 0-.263-.02H7.438a.606.606 0 0 0-.436.18a.624.624 0 0 0-.18.436v25.75a.615.615 0 0 0 .616.615h18.364a.61.61 0 0 0 .434-.18a.613.613 0 0 0 .182-.435V11.648c-.036-1.99-1.594-4.158-3.394-5.975m-.867.872c.805.786 1.53 1.676 2.07 2.534a10.39 10.39 0 0 0-1.42-.432a17.073 17.073 0 0 0-2.626-.357c.004-.132.005-.282.005-.445c0-.87-.055-2.108-.356-3.2l-.01-.03a14.622 14.622 0 0 1 2.337 1.93m3.027 21.62H8.052V3.644h9.542v.003c.416-.025.775.386 1.05 1.326c.25.895.313 2.062.312 2.87c.002.594-.027.992-.027.992l-.05.652l.656.007c.003 0 1.516.018 3 .355c1.426.308 2.54.922 2.645 1.617c.005.062.006.124.005.182v16.514z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paint(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.54 5.772V2.208H3.293v8.083h22.25V6.77c1.21.114 2.166 1.175 2.166 2.48v1.374c0 1.398-1.164 2.784-2.54 3.025l-7.883 1.38c-1.857.326-3.37 2.125-3.37 4.01v.386c-.74.366-1.29 1.507-1.29 2.865v4.5c0 1.65.806 3 1.79 3s1.793-1.35 1.793-3v-4.5c0-1.357-.55-2.498-1.292-2.864v-.385c0-1.397 1.164-2.783 2.54-3.024l7.883-1.38c1.856-.327 3.368-2.126 3.368-4.01V9.25c0-1.856-1.404-3.364-3.167-3.478z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pallete(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.653 7.25c-3.417 0-8.577.983-8.577 3.282c0 1.91 2.704 3.23 1.69 3.89c-1.02.665-2.683-1.85-4.047-1.85c-1.654 0-2.816 1.435-2.816 2.927c0 4.557 6.326 8.25 13.75 8.25c7.423 0 13.442-3.693 13.442-8.25c0-4.556-6.02-8.25-13.443-8.25zm-5.345 6.27c0-.644.887-1.165 1.98-1.165s1.98.52 1.98 1.166c0 .645-.887 1.167-1.98 1.167s-1.98-.523-1.98-1.166zm3.98 8.78c-1.057 0-1.913-.68-1.913-1.52s.856-1.517 1.914-1.517c1.056 0 1.913.68 1.913 1.518s-.857 1.52-1.914 1.52zm5.323-.53c-1.056 0-1.912-.68-1.912-1.518c0-.84.856-1.52 1.913-1.52c1.06 0 1.915.68 1.915 1.52s-.855 1.52-1.914 1.52zm.465-11.11c0-.838.856-1.518 1.914-1.518s1.912.68 1.912 1.518c0 .84-.855 1.518-1.913 1.518c-1.056 0-1.915-.68-1.915-1.518zm4.2 8.822c-1.057 0-1.914-.68-1.914-1.52s.858-1.517 1.915-1.517c1.06 0 1.914.68 1.914 1.518s-.856 1.52-1.915 1.52zm1.01-4.007c-1.057 0-1.913-.68-1.913-1.52c0-.837.856-1.517 1.914-1.517c1.057 0 1.913.68 1.913 1.518c0 .84-.857 1.52-1.914 1.52z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Palm(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.296 27.885v-2.013s-.402-1.408-1.073-2.013c-.67-.605-1.274-1.275-1.41-1.61c0 0-.267.134-.736-.336s-1.812-2.616-1.812-2.616l-.67-.872s-.47-.67-1.276-1.342s-.94-.067-1.477-.738s.604-1.275 1.006-1.41c.402-.133 1.945.135 2.683.873l.738.737l1.074 1.14l.537.202l.672-1.074l-.27-2.28s-.603-2.55-.736-4.765c-.135-2.213-.47-5.702 1.006-5.836s1.007 2.55 1.073 3.49c.067.937.806 5.23 1.208 5.567c.402.336.67.067.67.067l.403-7.514s-.48-2.438 1.073-2.55c.938-.066.87 1.544.87 2.148c0 .605.27 7.515.27 7.515l.537.135s.402-2.214.604-3.153s.605-2.416.538-3.087c-.067-.67-.135-2.348 1.006-2.348s.872 1.812.94 2.415s-.135 3.153-.135 3.757c0 .604-.74 3.623-.538 3.824s2.08-2.817 2.35-3.958c.267-1.14.2-3.02 1.407-2.885c1.207.134.47 2.817.4 3.086c-.065.27-.67 2.35-.87 2.953s-.806 1.476-1.007 2.013s.402 2.35 0 4.63s-1.61 5.165-1.61 5.165l.604 2.08s-1.744.672-3.824.806c-2.08.135-4.227-.2-4.227-.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paper(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28.916 8.01L15.953 1.887a.92.92 0 0 0-1.288.638c-.003.01-1.04 4.83-2.58 9.635c-.525 1.647-1.113 3.275-1.727 4.705l1.665.786c2-4.642 3.584-11.05 4.18-13.613L27.47 9.353c-.346 1.513-1.233 5.223-2.42 8.927c-.767 2.4-1.665 4.797-2.585 6.532c-.89 1.79-1.958 2.67-2.197 2.552c-1.42.03-2.418-1.262-3.09-2.918a13.7 13.7 0 0 1-.657-2.246c-.128-.618-.167-1.006-.17-1.006a.906.906 0 0 0-.52-.73l-12.96-6.12a.924.924 0 0 0-.926.08a.92.92 0 0 0-.38.847c.008.046.195 1.85.947 3.737c.522 1.32 1.407 2.818 2.846 3.575l12.956 6.13l.006-.012c.562.295 1.2.487 1.947.496c1.797-.117 2.777-1.668 3.818-3.525c3-5.69 5.32-16.6 5.338-16.64a.91.91 0 0 0-.504-1.02z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Parent(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.423 12.17a5.233 5.233 0 0 1-3.102 1.03c-1.163 0-2.23-.39-3.103-1.03c-.75.625-1.498 1.52-2.11 2.623c-1.423 2.563-1.58 5.192-.35 5.874c.548.312 1.126.078 1.722-.496a10.51 10.51 0 0 0-.167 1.874c0 2.938 1.14 5.312 2.543 5.312c.846 0 1.265-.865 1.466-2.188c.2 1.314.62 2.188 1.46 2.188c1.397 0 2.546-2.375 2.546-5.312c0-.66-.062-1.29-.168-1.873c.6.575 1.176.813 1.726.497c1.227-.682 1.068-3.31-.354-5.874c-.61-1.104-1.36-1.998-2.11-2.623zm-3.103.03a4.279 4.279 0 1 0-.003-8.561a4.279 4.279 0 0 0 .003 8.563zm10.667 5.47c1.508 0 2.732-1.224 2.732-2.734S23.493 12.2 21.986 12.2a2.737 2.737 0 0 0-2.736 2.736a2.737 2.737 0 0 0 2.737 2.735zm3.33 1.657c-.39-.705-.868-1.277-1.348-1.677c-.56.41-1.24.66-1.983.66c-.744 0-1.426-.25-1.983-.66c-.48.4-.958.972-1.35 1.677c-.91 1.638-1.01 3.318-.224 3.754c.35.2.72.05 1.1-.316a6.874 6.874 0 0 0-.104 1.197c0 1.88.728 3.397 1.625 3.397c.54 0 .81-.553.938-1.398c.128.84.396 1.397.934 1.397c.893 0 1.627-1.518 1.627-3.396c0-.42-.04-.824-.108-1.196c.383.367.752.52 1.104.317c.782-.434.68-2.115-.228-3.753z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pc(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M29.25 3.14h-9.19l-.458.46v18.224l.33 2.39h-.362v.244h-.307v-.306h-.61v.244h-.312v-.367h-.485v.306H16.75l-2.02-.367v-.92h.857l.302-1.47h2.727a.34.34 0 0 0 .34-.34V7.828a.337.337 0 0 0-.34-.338H1.59a.339.339 0 0 0-.338.338V21.24c0 .187.152.34.34.34h3.015l.2 1.47h1.408l-3.4 3.4l-.705 1.5s2.94 1.103 6.678 1.103c3.737 0 9.68-.857 10.476-.857h4.84V26.97l-.137-1.067h1.744c-.2.106-.32.244-.32.396v.978c0 .34.603.613 1.35.613c.743 0 1.35-.27 1.35-.612v-.98c0-.338-.605-.61-1.35-.61c-.187 0-.363.02-.524.05v-.17h-2.29l-.055-.433h5.383zM2.477 20.17V8.714h15.07V20.17z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pen(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.587 12.074a.765.765 0 0 0-.188-.202a.75.75 0 1 0-.881 1.214c.302.22.707.167.965-.09L6.4 22.762L4.195 30.76l6.917-4.577L26.865 4.468l-4.716-3.42l-1.52 2.096a1.588 1.588 0 0 0-.597-.907a1.63 1.63 0 0 0-2.28.363c-3.032 4.182-1.35 5.296-4.166 9.474zm-3.47 13.074L6.56 27.503l1.133-4.117zM14.31 11.86c2.182-3.224 1.974-4.098 3.842-6.96c.31.21.664.286 1.012.268z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pensil(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.31 2.872L21.926.745a1.826 1.826 0 0 0-2.517.576l-1.335 2.124L24.55 7.51l1.334-2.122c.536-.855.28-1.98-.574-2.516M6.555 21.786l6.474 4.066L23.58 9.054l-6.476-4.067l-10.55 16.8zm-.99 5.166l-.142 3.82l3.38-1.788l3.14-1.658L5.695 23.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func People(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.066 20.667c1.227-.682 1.068-3.31-.354-5.874c-.61-1.104-1.36-1.998-2.11-2.623a5.229 5.229 0 0 1-3.1 1.03a5.23 5.23 0 0 1-3.105-1.03c-.75.625-1.498 1.52-2.11 2.623c-1.423 2.563-1.58 5.192-.35 5.874c.548.312 1.126.078 1.722-.496a10.51 10.51 0 0 0-.167 1.874c0 2.938 1.14 5.312 2.543 5.312c.846 0 1.265-.865 1.466-2.188c.2 1.314.62 2.188 1.46 2.188c1.397 0 2.546-2.375 2.546-5.312c0-.66-.062-1.29-.168-1.873c.6.575 1.176.813 1.726.497zM15.5 12.2a4.279 4.279 0 1 0-.003-8.557A4.279 4.279 0 0 0 15.5 12.2m8.594 2.714a3.514 3.514 0 0 0 0-7.025a3.513 3.513 0 1 0 .001 7.027zm4.28 2.13c-.502-.908-1.116-1.642-1.732-2.155a4.3 4.3 0 0 1-2.546.845c-.756 0-1.46-.207-2.076-.55c.496 1.093.803 2.2.86 3.19c.094 1.516-.38 2.64-1.328 3.165a2.017 2.017 0 0 1-.653.224c-.057.392-.096.8-.096 1.23c0 2.413.935 4.362 2.088 4.362c.694 0 1.04-.71 1.204-1.796c.163 1.08.508 1.796 1.2 1.796c1.145 0 2.09-1.95 2.09-4.36c0-.543-.053-1.06-.14-1.54c.492.473.966.668 1.418.408c1.007-.56.877-2.718-.29-4.82zm-21.468-2.13a3.512 3.512 0 1 0-3.514-3.512a3.515 3.515 0 0 0 3.514 3.514zm2.535 6.622c-1.592-.885-1.738-3.524-.456-6.354a4.242 4.242 0 0 1-2.078.553c-.956 0-1.832-.32-2.55-.846c-.615.512-1.228 1.246-1.732 2.153c-1.167 2.104-1.295 4.262-.287 4.82c.45.258.925.065 1.414-.406a8.83 8.83 0 0 0-.135 1.538c0 2.412.935 4.36 2.088 4.36c.694 0 1.04-.71 1.204-1.795c.165 1.08.51 1.796 1.2 1.796c1.147 0 2.09-1.95 2.09-4.36c0-.433-.04-.842-.097-1.234a2.02 2.02 0 0 1-.66-.226z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.065 18.53c-.467-.29-1.167-.21-1.556.18l-3.094 3.09a1 1 0 0 1-1.414 0L9.05 14.85a1 1 0 0 1 0-1.414l2.913-2.912c.39-.39.447-1.075.13-1.524l-5.3-7.513a.936.936 0 0 0-1.36-.195s-4.134 3.28-4.134 6.295c0 12.335 10 22.334 22.333 22.334c3.015 0 5.948-5.534 5.948-5.534a1.087 1.087 0 0 0-.38-1.412l-7.135-4.444z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Photo(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.25 10.25H20.5v-1.5h-9.375v1.5h-3.75a2 2 0 0 0-2 2v10.375a2 2 0 0 0 2 2H24.25a2 2 0 0 0 2-2V12.25a2 2 0 0 0-2-2M15.812 23.5c-3.342 0-6.06-2.72-6.06-6.062s2.718-6.062 6.06-6.062s6.062 2.72 6.062 6.062s-2.72 6.06-6.062 6.06zm0-10.125a4.062 4.062 0 1 0 .001 8.127a4.062 4.062 0 0 0-.001-8.128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Picker(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.22 10.853c-.11-.414-.26-.412.222-1.54l1.66-3.518c.02-.05.2-.412.192-.946c.015-.53-.313-1.29-1.12-1.643c-1.17-.555-1.17-.557-2.343-1.107c-.783-.396-1.58-.17-1.978.18c-.42.332-.584.7-.61.75L16.58 6.544c-.564 1.084-.655.97-1.048 1.147c-.47.13-1.244.558-1.785 1.815l-1.108 2.346l-.277.586l1.17.552l-3.6 7.623c-.38.828-.165 1.436-.165 2.032c.01.627-.077 1.51-.876 3.21l-.276.586l3.517 1.66l.276-.584c.807-1.7 1.43-2.327 1.92-2.718c.46-.38 1.067-.6 1.466-1.42l3.6-7.618l1.17.554l.28-.59l1.105-2.344c.628-1.218.47-2.083.27-2.53zM14.624 22.83c-.156.353-.413.44-1.09.955c-.578.448-1.265 1.172-2.01 2.6l-1.19-.562c.627-1.48.75-2.474.73-3.203c-.032-.85-.13-1.104.044-1.45l3.6-7.62l3.516 1.662z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Picture(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 4.833v22.334h27V4.833zM25.25 25.25H6.75V6.75h18.5zM11.25 14a2.584 2.584 0 1 0-.001-5.167A2.584 2.584 0 0 0 11.25 14m13 2.25l-4.916-4.917l-6.917 6.917l-1.917-1.917l-2.752 2.752v5.165H24.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Piechart(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.583 15.917l1.648-10.78A11.294 11.294 0 0 0 15.584 5C9.553 5 4.666 9.888 4.666 15.917c0 6.03 4.888 10.917 10.917 10.917S26.5 21.946 26.5 15.917c0-.256-.02-.507-.038-.76l-10.88.76zm3.854-12.79l-1.648 10.78l10.878-.76a10.908 10.908 0 0 0-9.23-10.02"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plane(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.67 8.11l-2.776 2.777l-3.837-.86c.362-.506.916-1.684.464-2.136c-.517-.516-1.978.28-2.304.605l-.913.913l-2.69-.604l-2.02 2.02l2.232 1.062l-.082.082l1.7 1.7l.69-.686l3.163 1.504L9.57 18.21H6.414l-1.137 1.138l3.6.948l1.83 1.83l.947 3.598l1.137-1.137V21.43l3.725-3.725l1.504 3.164l-.688.686l1.702 1.7l.08-.08l1.063 2.23l2.02-2.02l-.604-2.688l.912-.912c.326-.326 1.12-1.79.604-2.306c-.453-.452-1.63.1-2.136.464l-.86-3.838l2.776-2.777c.947-.948 3.6-4.863 2.62-5.84c-.977-.978-4.892 1.673-5.84 2.62z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.684 25.682L24.316 15.5L6.684 5.318z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plugin(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m26.33 15.836l-3.893-1.545l3.136-7.9a1.374 1.374 0 1 0-2.556-1.014l-3.136 7.9L15 11.34l3.136-7.9a1.375 1.375 0 0 0-2.555-1.016l-3.135 7.9l-3.89-1.543l-1.615 4.067l2.15.854l-2.537 6.392a3.002 3.002 0 0 0 1.683 3.895l1.626.646l-.877 2.207a2 2 0 0 0 1.122 2.596l.93.37a2.001 2.001 0 0 0 2.596-1.122l.877-2.207l1.858.737a3.002 3.002 0 0 0 3.896-1.682l2.535-6.39l1.917.76l1.613-4.066z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.98 12.896h-6.67V6.23h-6.665v6.666H5.98v6.666h6.667v6.667h6.665V19.56h6.667z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Power(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.542 8.354c-1.47-1.766-2.896-2.617-3.025-2.695a1.99 1.99 0 0 0-2.74.723c-.555.96-.24 2.194.706 2.763h.002l.003.003h.004c.102.062 1.124.73 2.08 1.925c1.003 1.26 1.933 3.017 1.937 5.438a9.122 9.122 0 0 1-2.64 6.438c-1.638 1.653-3.878 2.67-6.37 2.67c-2.49 0-4.73-1.017-6.368-2.67a9.105 9.105 0 0 1-2.64-6.437c.005-2.5.995-4.292 2.035-5.558a9.678 9.678 0 0 1 1.425-1.4c.19-.153.346-.264.445-.33l.104-.07l.013-.007l.004-.002a2.036 2.036 0 0 0 .705-2.763a1.991 1.991 0 0 0-2.74-.724c-.127.078-1.554.93-3.023 2.695c-1.463 1.75-2.975 4.51-2.97 8.157c0 7.263 5.825 13.152 13.01 13.155c7.186-.003 13.01-5.892 13.012-13.155c.004-3.648-1.507-6.407-2.97-8.158zM15.5 17.524a2.015 2.015 0 0 0 2.002-2.024V3.357c0-1.118-.897-2.024-2.002-2.024s-2.002.906-2.002 2.024V15.5c0 1.116.897 2.023 2.002 2.023z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ppt(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.604 1.914a1.04 1.04 0 1 0-2.082 0v1.04h2.082zm0 20.803h-2.082v3.207c0 .574-4.225 4.03-4.225 4.03l2.468-.002l2.807-2.673l3.013 2.692l2.272-.04l-4.254-4.01zM28.566 7.113a1.56 1.56 0 0 0 0-3.12H2.56a1.56 1.56 0 0 0 0 3.12h1.584v12.505l-.932-.022c-.86 0-1.213.467-1.213 1.04c0 .576.35 1.04 1.212 1.04H27.81c.86 0 1.298-.464 1.298-1.04c0-1.094-1.3-1.04-1.3-1.04l-.803.11V7.112h1.56zm-17.13 10.403c-3.772 0-4.195-4.19-4.195-4.19c0-4.097 4.163-4.162 4.163-4.162v4.16h4.193c0 4.192-4.16 4.192-4.16 4.192zm7.28-4.128h-1.07v-1.073h1.07zm0-3.12h-1.07V9.193h1.07v1.073zm4.598 3.12H20.26a.536.536 0 0 1 0-1.073h3.057a.537.537 0 0 1-.003 1.073m0-3.12H20.26a.536.536 0 0 1 0-1.073h3.057a.535.535 0 1 1-.003 1.072z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Printer(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.57 12.125h-2.12c-.208-1.34-1.248-2.76-2.445-3.967c-1.277-1.24-2.654-2.234-3.784-2.37a1.615 1.615 0 0 0-.197-.015h-7.43a.449.449 0 0 0-.45.452v5.9H6.07a2 2 0 0 0-2 2V23h4.073v2.08a.45.45 0 0 0 .452.447h13.444c.117 0 .23-.045.317-.13s.138-.2.138-.318V23h4.074v-8.875a2 2 0 0 0-2-2zm-2.98 12.5H9.042V21.5H21.59v3.126zm0-10.704c0-.03 0-.062-.004-.095a.45.45 0 0 0-.124-.2H9.042v-6.95h6.988c.305-.018.567.283.77.972c.182.655.228 1.51.228 2.102c0 .432-.02.724-.02.724l-.036.478l.48.005c.002 0 1.11.014 2.196.26c1.044.226 1.86.675 1.938 1.184c.003.046.003.092.003.134v1.387z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Question(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 1.466C7.973 1.466 1.466 7.973 1.466 16c0 8.027 6.507 14.534 14.534 14.534c8.027 0 14.534-6.507 14.534-14.534c0-8.027-6.507-14.534-14.534-14.534m1.328 22.905H14.62v-2.595h2.708v2.596zm0-5.367v.858H14.62v-1.056c0-3.19 3.63-3.696 3.63-5.963c0-1.033-.923-1.825-2.133-1.825c-1.254 0-2.354.924-2.354.924l-1.54-1.916S13.74 8.44 16.358 8.44c2.486 0 4.795 1.54 4.795 4.136c0 3.632-3.827 4.05-3.827 6.427z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionTwo(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M26.71 14.086L16.915 4.29a2.007 2.007 0 0 0-2.83 0L4.29 14.086a2.007 2.007 0 0 0 0 2.83l9.796 9.795c.778.778 2.05.778 2.83 0l9.796-9.796a2.009 2.009 0 0 0 0-2.828zM16.43 21.8c-.247.24-.542.36-.884.36a1.23 1.23 0 0 1-.886-.36c-.247-.242-.37-.534-.37-.877s.123-.638.37-.885a1.21 1.21 0 0 1 .886-.372c.342 0 .637.124.885.372c.25.247.372.542.372.885s-.123.635-.37.876zm2.48-6.602c-.72.716-1.71 1.147-2.97 1.294v2.027h-.845v-3.477c.386-.03.768-.093 1.146-.188c.38-.095.72-.25 1.02-.464c.312-.226.555-.5.73-.82c.173-.323.26-.77.26-1.347c0-.918-.194-1.623-.582-2.113c-.39-.49-.956-.735-1.7-.735c-.282 0-.528.042-.74.124s-.365.16-.463.233c.03.146.072.357.124.633c.05.275.077.486.077.633c0 .225-.098.432-.294.618c-.195.187-.48.28-.853.28c-.33 0-.565-.113-.706-.34s-.21-.488-.21-.788c0-.244.066-.484.2-.72c.135-.235.346-.463.633-.684c.245-.195.577-.364.995-.504a4.1 4.1 0 0 1 1.308-.21c.647 0 1.223.102 1.724.307a3.48 3.48 0 0 1 1.238.82c.337.356.586.756.748 1.2c.162.443.243.925.243 1.445c0 1.133-.36 2.057-1.082 2.773z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Quote(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.505 5.873C10.568 8.393 8.6 11.43 8.6 14.98c0 1.105.193 1.657.577 1.657l.396-.107c.312-.12.563-.18.756-.18c1.127 0 2.07.41 2.825 1.23c.756.82 1.134 1.83 1.134 3.036c0 1.157-.41 2.14-1.225 2.947c-.816.807-1.8 1.21-2.952 1.21c-1.608 0-2.935-.66-3.98-1.983c-1.043-1.32-1.564-2.98-1.564-4.977c0-2.26.442-4.327 1.33-6.203c.89-1.875 2.244-3.57 4.068-5.085c1.824-1.514 2.988-2.272 3.492-2.272c.336 0 .612.162.828.486c.216.323.324.605.324.845l-.107.288zm12.96 0c-3.937 2.52-5.904 5.556-5.904 9.108c0 1.105.193 1.657.577 1.657l.396-.107c.312-.12.563-.18.756-.18c1.103 0 2.04.41 2.807 1.23c.77.82 1.152 1.83 1.152 3.036c0 1.157-.41 2.14-1.225 2.947c-.816.807-1.8 1.21-2.952 1.21c-1.608 0-2.935-.66-3.98-1.983c-1.043-1.32-1.564-2.98-1.564-4.977c0-2.284.448-4.37 1.35-6.256c.9-1.887 2.255-3.577 4.067-5.067C24.76 5 25.917 4.254 26.42 4.254c.337 0 .613.162.83.486c.215.324.323.606.323.846z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rain(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.37 7.306a7.252 7.252 0 0 0-7.247-7.08a7.24 7.24 0 0 0-6.208 3.518a4.163 4.163 0 0 0-2.01-.517a4.22 4.22 0 0 0-4.216 4.066a5.987 5.987 0 0 0-4.128 5.686c0 3.31 2.687 6 6 6v-.003h15.874c3.312 0 6-2.688 6-6a5.993 5.993 0 0 0-4.065-5.67m-1.934 9.673H7.56a4.01 4.01 0 0 1-4-4.002c-.002-1.982 1.45-3.618 3.35-3.93a.998.998 0 0 0 .657-.415c.155-.223.212-.497.163-.763a2.284 2.284 0 0 1-.045-.424a2.224 2.224 0 0 1 2.22-2.218c.647 0 1.217.278 1.633.73a.997.997 0 0 0 1.675-.32c.75-1.992 2.662-3.412 4.91-3.41a5.258 5.258 0 0 1 5.252 5.25c0 .16-.01.325-.027.496c-.05.517.305.984.815 1.08c1.86.343 3.274 1.965 3.27 3.922a4.005 4.005 0 0 1-3.997 4.003zM9.03 26.68c0-1.115.02-5.425.02-5.432a1.001 1.001 0 0 0-1.728-.692c-.005.008-1.036 1.098-2.08 2.342a25.656 25.656 0 0 0-1.463 1.896c-.4.648-.754 1.066-.812 1.885a3.037 3.037 0 0 0 3.032 3.034a3.036 3.036 0 0 0 3.03-3.032zm-4.06.045c.092-.35 1.082-1.72 1.994-2.764l.076-.09c-.005 1.125-.01 2.295-.01 2.81c0 .566-.46 1.027-1.03 1.03a1.038 1.038 0 0 1-1.03-.986m11.455-.045c0-1.115.02-5.424.02-5.43a1 1 0 0 0-1.727-.692c-.006.008-1.035 1.094-2.08 2.342a25.344 25.344 0 0 0-1.463 1.894c-.4.65-.753 1.068-.81 1.888a3.031 3.031 0 0 0 6.06-.002m-4.06.047c.092-.35 1.08-1.72 1.993-2.766l.075-.09c-.005 1.124-.01 2.295-.01 2.808a1.035 1.035 0 0 1-1.03 1.03c-.553-.003-1-.44-1.028-.983zm10.906-6.413a.996.996 0 0 0-1.098.24a56.757 56.757 0 0 0-2.08 2.342c-.523.624-1.05 1.284-1.462 1.895c-.402.65-.754 1.067-.812 1.886a3.029 3.029 0 1 0 6.057 0c0-1.114.022-5.424.022-5.43a.999.999 0 0 0-.626-.933zm-1.39 6.364a1.036 1.036 0 0 1-1.032 1.03a1.033 1.033 0 0 1-1.028-.982c.092-.35 1.08-1.72 1.993-2.765c.025-.028.05-.06.074-.088c-.004 1.123-.008 2.292-.008 2.806z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Raphael(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27.777 18.94c.584-.88.896-1.913.896-2.997a5.405 5.405 0 0 0-1.598-3.854l-6.91-6.912l-.003.002a5.427 5.427 0 0 0-3.85-1.6a5.421 5.421 0 0 0-3.85 1.6h-.002l-6.91 6.91a5.412 5.412 0 0 0-1.6 3.855c0 1.457.568 2.826 1.6 3.854l6.23 6.228c.25.28.512.544.79.785a5.413 5.413 0 0 0 3.742 1.49a5.4 5.4 0 0 0 3.854-1.598l6.723-6.725l.05-.05l.14-.14c.26-.26.487-.54.688-.837c.004-.008.01-.015.014-.02l-.005.008zm-1.12-2.994a3.41 3.41 0 0 1-.56 1.88a6.147 6.147 0 0 1-.684.783l-.013.015c-1.105 1.052-2.354 1.35-3.414 1.35c-.584 0-1.11-.09-1.523-.195c-2.422-.608-5.056-2.692-6.26-5.732c.648.274 1.36.426 2.11.426c2.81 0 5.128-2.14 5.414-4.877l3.924 3.925a3.392 3.392 0 0 1 1.008 2.424zM16.313 5.6a3.43 3.43 0 0 1 3.426 3.427a3.43 3.43 0 0 1-3.426 3.427a3.43 3.43 0 0 1-3.426-3.427A3.43 3.43 0 0 1 16.313 5.6M6.974 18.375a3.409 3.409 0 0 1-1.007-2.43c0-.916.357-1.78 1.007-2.427l2.655-2.656c-.694 2.36-.992 4.842-.832 7.22c.057.855.175 1.678.345 2.46l-2.17-2.167zm4.54-6.783c.583 4.562 4.195 9.066 8.455 10.143a8.07 8.07 0 0 0 2.032.265h.027l-3.29 3.29a3.426 3.426 0 0 1-4.208.505l.002-.002a7.007 7.007 0 0 1-.603-.46c-.014-.02-.03-.027-.045-.044l-.665-.665c-1.367-1.567-2.227-3.903-2.412-6.67c-.14-2.1.113-4.282.706-6.363"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reflecth(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.57 20.273h.854v-1.705h-.854zm0 3.413h.854V21.98h-.854zm0 3.41h.854V25.39h-.854zm0 2.593h.854v-.89h-.854zm0-12.825h.854V15.16h-.854v1.705zm0-13.64h.854V1.52h-.854zm0 3.41h.854V4.93h-.854zm0 3.41h.854V8.34h-.854zm0 3.41h.854V11.75h-.854zm2.84-10.128V25.46h12.015zm.854 3.353l9.73 17.93h-9.73zm-5.73 18.78V3.327L1.522 25.46h12.015z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reflectv(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.643 16.008v-.854h-1.705v.854zm3.41 0v-.854h-1.705v.854zm3.41 0v-.854h-1.705v.854zm2.596 0v-.854h-.892v.854h.89zm-12.828 0v-.854h-1.71v.854zm-13.64 0v-.854H1.89v.854h1.705zm3.41 0v-.854H5.3v.854h1.705zm3.412 0v-.854H8.71v.854h1.704zm3.41 0v-.854H12.12v.854zm-10.13-2.84h22.134V1.15zm3.354-.854l17.93-9.73v9.73zm18.78 5.728H3.694l22.134 12.015z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Refresh(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.083 15.5a8.596 8.596 0 0 1-8.583 8.583A8.597 8.597 0 0 1 6.915 15.5A8.596 8.596 0 0 1 15.5 6.915c1.913 0 3.665.63 5.09 1.686l-1.782 1.784l8.43 2.256l-2.26-8.427l-1.89 1.89A12.036 12.036 0 0 0 15.5 3.415C8.826 3.418 3.418 8.825 3.416 15.5c.002 6.675 5.41 12.083 12.084 12.083S27.583 22.175 27.583 15.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ResizeTwo(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2.332v26.5h26.5V2.33H2zm24.5 24.5H4V12.5h8.167V4.332H26.5zM15.63 17.65l5.47 5.468l-1.21 1.206l5.483 1.47l-1.47-5.482l-1.195 1.195l-5.467-5.466l1.21-1.207l-5.482-1.47l1.468 5.48l1.195-1.194z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Roadmap(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.188 3.735a1.766 1.766 0 0 0-3.532-.001c0 .975 1.766 4.267 1.766 4.267s1.766-3.292 1.766-4.267zm-2.61 0a.844.844 0 1 1 1.687-.001a.844.844 0 0 1-1.687.001m4.703 14.76c-.56 0-1.097.047-1.59.123L11.1 13.976c.2-.18.312-.38.312-.59a.663.663 0 0 0-.088-.315l8.41-2.238c.46.137 1.023.22 1.646.22c1.52 0 2.75-.484 2.75-1.082c0-.6-1.23-1.083-2.75-1.083s-2.75.485-2.75 1.083c0 .07.02.137.054.202L9.896 12.2a8.075 8.075 0 0 0-2.265-.303c-2.087 0-3.78.667-3.78 1.49s1.693 1.49 3.78 1.49c.574 0 1.11-.055 1.598-.145l11.99 4.866c-.19.192-.306.4-.306.623c0 .19.096.364.236.533L8.695 25.415c-.158-.005-.316-.01-.477-.01c-3.24 0-5.87 1.036-5.87 2.31c0 1.277 2.63 2.313 5.87 2.313s5.87-1.034 5.87-2.312c0-.22-.083-.432-.23-.633l10.266-5.214c.37.04.753.065 1.155.065c2.413 0 4.37-.77 4.37-1.723c0-.944-1.957-1.716-4.37-1.716z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rotate(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 5.27a8.5 8.5 0 0 1 5.09 1.686L18.807 8.74l8.428 2.255l-2.26-8.427l-1.89 1.89A12.01 12.01 0 0 0 15.5 1.77C8.827 1.773 3.418 7.18 3.417 13.855c0 4.063 2.012 7.647 5.084 9.838v-4.887a8.55 8.55 0 0 1-1.584-4.952a8.594 8.594 0 0 1 8.584-8.584zm-6 23.96h12V12.355h-12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ruler(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.63 21.796l-5.122 5.12H27.25V1.177zM18.702 10.48a.47.47 0 0 1 .664 0l1.16 1.16a.466.466 0 1 1-.66.661l-1.164-1.16a.468.468 0 0 1 0-.662zm-1.6 1.604a.465.465 0 0 1 .66 0l2.157 2.154a.465.465 0 0 1 0 .66a.463.463 0 0 1-.663.003L17.1 12.748a.469.469 0 0 1 0-.663zm-1.605 1.6a.472.472 0 0 1 .664 0l1.16 1.162a.467.467 0 0 1-.33.8a.469.469 0 0 1-.333-.138l-1.16-1.16a.471.471 0 0 1 0-.665zm-1.6 1.604a.47.47 0 0 1 .663.002l1.158 1.16a.468.468 0 1 1-.664.662l-1.158-1.16a.47.47 0 0 1 0-.664zm-1.604 1.604a.467.467 0 0 1 .663 0l2.154 2.153a.47.47 0 1 1-.664.665l-2.153-2.155a.469.469 0 0 1 0-.663m-1.99 7.623a.468.468 0 0 1-.663.002l-2.154-2.153a.468.468 0 1 1 .662-.663l2.154 2.154a.465.465 0 0 1 0 .66zm.61-2.597a.476.476 0 0 1-.334.14a.46.46 0 0 1-.33-.138l-1.163-1.16a.464.464 0 0 1 0-.66a.467.467 0 0 1 .664-.004l1.162 1.162a.465.465 0 0 1 0 .66zm1.6-1.602a.465.465 0 0 1-.662 0l-1.16-1.16a.468.468 0 1 1 .664-.662l1.16 1.16a.47.47 0 0 1 0 .662zm9.737 1.6h-8.67l8.67-8.67zM22.13 10.7a.46.46 0 0 1-.33.138a.473.473 0 0 1-.334-.138l-1.16-1.16a.468.468 0 1 1 .662-.662l1.16 1.16c.184.183.185.48.002.662m2.596-.608a.469.469 0 0 1-.662-.001L21.91 7.938a.468.468 0 1 1 .664-.662l2.154 2.154c.183.183.18.48-.002.662"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Run(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m17.41 20.395l-.778-2.723c.228-.2.442-.414.644-.643l2.72.777a9.75 9.75 0 0 0 .756-1.323l-2.025-1.96c.097-.288.18-.58.24-.883l2.73-.683c.02-.252.04-.505.04-.763s-.02-.51-.04-.762l-2.73-.684a7.273 7.273 0 0 0-.24-.883l2.026-1.96c-.222-.46-.47-.905-.756-1.323l-2.72.777a7.127 7.127 0 0 0-.645-.644l.778-2.722a9.972 9.972 0 0 0-1.324-.755l-1.96 2.025a7.002 7.002 0 0 0-.883-.24l-.683-2.73c-.253-.02-.505-.04-.763-.04s-.51.02-.762.04l-.684 2.73c-.3.06-.594.143-.882.24L7.508 3.24c-.46.223-.904.47-1.322.757l.778 2.722c-.228.2-.443.414-.643.64L3.6 6.584a9.76 9.76 0 0 0-.756 1.324l2.026 1.96c-.096.288-.18.58-.24.883l-2.73.684c-.02.252-.04.505-.04.762s.02.51.04.763l2.73.683c.06.302.144.595.24.883l-2.026 1.96c.22.46.468.905.755 1.323l2.72-.78c.2.23.416.443.644.644l-.778 2.723c.418.286.863.533 1.323.755l1.96-2.026c.287.097.58.18.882.24l.684 2.73c.252.02.505.04.763.04s.51-.02.762-.04l.683-2.73c.302-.06.596-.144.883-.24l1.96 2.026c.46-.224.904-.47 1.322-.757zm-5.612-4.8a3.4 3.4 0 1 1-.001-6.797a3.4 3.4 0 0 1 .001 6.797m15.492 7.104a4.987 4.987 0 0 0-.23-1.655l1.244-1.773c-.188-.35-.4-.682-.64-.984l-2.123.445a4.986 4.986 0 0 0-1.435-.85l-.61-2.08a6.87 6.87 0 0 0-1.174-.106l-.973 1.936c-.28.054-.558.128-.832.233c-.257.098-.497.22-.727.353l-2.006-.82a6.743 6.743 0 0 0-.813.852l.906 1.968a4.982 4.982 0 0 0-.52 1.585l-1.89 1.06c.02.388.076.776.164 1.165l2.104.52c.23.523.54.992.916 1.392l-.352 2.138c.32.23.66.428 1.013.6l1.716-1.32c.536.14 1.097.195 1.662.15l1.452 1.607c.2-.057.4-.118.596-.193a6.38 6.38 0 0 0 .505-.223l.038-2.164c.455-.34.843-.747 1.152-1.206l2.16-.134c.153-.36.28-.732.37-1.115l-1.67-1.38zm-4.163 2.006a2.328 2.328 0 1 1-1.658-4.355a2.328 2.328 0 0 1 1.658 4.357z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rw(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.5 15.5l10.3 5.947v-5.6l9.7 5.6V9.552l-9.7 5.6v-5.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Safari(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.154 5.135c-.504 0-1 .03-1.488.09l-.036-.18a.836.836 0 0 0-.112-.284a1.677 1.677 0 1 0-2.731-1.305c0 .745.485 1.376 1.157 1.595a.904.904 0 0 0 0 .328l.033.167c-5.332 1.405-9.265 6.26-9.265 12.033c0 6.87 5.57 12.44 12.442 12.44c6.87 0 12.44-5.57 12.44-12.44c.002-6.872-5.57-12.443-12.44-12.443zM16.37 8.1c4.454 0 8.182 3.116 9.122 7.287l-.576.234a1.516 1.516 0 0 0-2.978.546l-2.77-.027a3.077 3.077 0 0 0-.183-.334l3.4-5.11l.054-.083l-4.767 4.06a2.92 2.92 0 0 0-1.148-.257l.086-.018l-1.176-2.585a1.515 1.515 0 0 0-.625-2.957l.058-.634c.49-.082.99-.123 1.5-.123zm-4.224-4.645a1.319 1.319 0 0 1 2.636 0c0 .425-.203.802-.516 1.043a.859.859 0 0 0-1.183.22a1.324 1.324 0 0 1-.938-1.263zm-5.13 13.997c0-4.443 3.1-8.163 7.253-9.116l.296.573a1.515 1.515 0 0 0 .664 2.942l-.058 2.845l.052-.01a2.963 2.963 0 0 0-1.116.894l.093-.146l-1.573-.603l1.172 1.24l.025-.042c-.19.37-.306.788-.324 1.23l-.002-.017l-2.623 1.21a1.515 1.515 0 0 0-2.94.665l-.783-.075a9.541 9.541 0 0 1-.133-1.588zm9.353 9.353a9.355 9.355 0 0 1-9.107-7.21l.69-.354a1.516 1.516 0 0 0 2.968-.622l2.858.03c.045.095.096.187.15.276l-3.45 5.277l.227-.194l4.53-3.92c.335.153.704.248 1.093.266l-.02.004l1.227 2.627a1.516 1.516 0 0 0 .374 2.984c.075 0 .15-.007.224-.018l.004.688a9.36 9.36 0 0 1-1.77.17zm2.292-.284l-.39-.6a1.514 1.514 0 0 0-.557-2.971v-2.86l-.025.003c.41-.185.77-.46 1.056-.798l1.516.66l-1.104-1.305c.158-.335.256-.704.278-1.095l2.552-1.164a1.514 1.514 0 0 0 2.948-.651l.65.12a9.348 9.348 0 0 1-6.922 10.66z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scissors(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.108 10.27c1.083-1.875.16-4.442-2.06-5.724a5.248 5.248 0 0 0-2.615-.72c-1.46 0-2.72.672-3.373 1.8c-.545.944-.608 2.077-.176 3.188C3.287 9.85 4.08 10.75 5.12 11.35a5.242 5.242 0 0 0 2.616.723c.673 0 1.3-.15 1.85-.414l2.4 1.385c1.582.914.56 3.86 5.918 6.955C23.26 23.094 29.4 21.535 29.4 21.535L10.75 10.767c.132-.156.255-.318.358-.496zm-1.733-1c-.506.88-2.033 1.056-3.255.348c-.646-.373-1.134-.916-1.37-1.528c-.21-.535-.194-1.055.042-1.464c.293-.51.892-.8 1.64-.8c.543 0 1.102.156 1.616.453c1.243.716 1.85 2.086 1.327 2.99zm7.87 6.522c0 .483-.39.875-.874.875c-.036 0-.067-.017-.103-.02l.667-1.512a.86.86 0 0 1 .312.657zm-.874-.875c.038 0 .07.017.105.02l-.666 1.51a.851.851 0 0 1-.313-.655c0-.483.39-.875.874-.875m13.03-4.45s-6.14-1.56-11.496 1.535c-.537.31-.995.618-1.415.924l4.325 2.497zm-16.23 6.63c-.35.85-.574 1.508-1.186 1.86l-2.4 1.385a4.235 4.235 0 0 0-1.85-.414c-.893 0-1.797.25-2.615.72c-2.22 1.283-3.144 3.852-2.06 5.727c.65 1.127 1.91 1.8 3.372 1.8c.894 0 1.8-.25 2.616-.72c1.04-.602 1.833-1.502 2.236-2.537c.432-1.112.368-2.245-.178-3.19a3.436 3.436 0 0 0-.356-.493l3.982-2.3a9.7 9.7 0 0 1-1.56-1.838zm-3.75 7.095c-.238.612-.725 1.155-1.37 1.528c-1.222.706-2.75.532-3.258-.347c-.522-.903.086-2.274 1.328-2.992a3.254 3.254 0 0 1 1.615-.452c.75 0 1.346.29 1.64.8c.237.41.253.93.045 1.464z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Screwdriver(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.387 14.373c2.12-2.62 5.322-6.77 5.15-7.75c-.13-.73-.883-1.547-1.764-2.17c-.883-.626-1.916-1.045-2.645-.916c-.98.173-3.786 4.603-5.52 7.49c-.21.344.327 1.177.155 1.468c-.172.292-1.052.042-1.18.26c-.263.452-.417.723-.417.723s-.553.823 1.163 2.163l-5.234 7.474c-.267.38-1.456.46-1.456.46l-1.184 3.31l.86.603l2.707-2.246S9.69 24.1 9.955 23.72l5.242-7.49c1.72 1 2.377.337 2.377.337l.536-.64c.16-.193-.374-.935-.16-1.196c.22-.263 1.183-.045 1.437-.357z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m29.772 26.433l-7.126-7.126a10.43 10.43 0 0 0 1.524-5.42c0-5.794-4.692-10.486-10.482-10.488c-5.79 0-10.484 4.693-10.484 10.485c0 5.79 4.693 10.48 10.484 10.48c1.987 0 3.84-.562 5.422-1.522l7.128 7.127l3.534-3.537zM7.202 13.885a6.496 6.496 0 0 1 6.485-6.486a6.493 6.493 0 0 1 6.484 6.485a6.494 6.494 0 0 1-6.483 6.484a6.496 6.496 0 0 1-6.484-6.485z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sencha(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.265 22.734a4.09 4.09 0 0 1 .133 7.297l1.922-.98a9.63 9.63 0 0 0 5.332-8.616c0-3.72-2.11-6.945-5.195-8.547l-6.272-3.144a4.097 4.097 0 0 1-2.308-3.684c0-1.566.88-2.927 2.175-3.613l-1.922.98a9.63 9.63 0 0 0-.137 17.163z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Settings(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.015 12.03a3.902 3.902 0 1 0 0 7.806c.494 0 .962-.102 1.397-.27l.836 1.285l1.36-.884l-.832-1.276a3.902 3.902 0 0 0-2.761-6.66zM16 1.466C7.973 1.466 1.466 7.973 1.466 16c0 8.027 6.507 14.534 14.534 14.534c8.027 0 14.534-6.507 14.534-14.534c0-8.027-6.507-14.534-14.534-14.534M26.174 20.81c-.24.503-.513.99-.826 1.45l-3.158-.68c-.48.526-1.03.994-1.634 1.385l.12 3.202c-.508.23-1.03.42-1.57.57l-1.955-2.514c-.37.05-.75.086-1.135.086c-.356 0-.706-.03-1.05-.075l-1.946 2.5a10.793 10.793 0 0 1-1.57-.57l.116-3.145a8.413 8.413 0 0 1-1.712-1.427l-3.074.658a10.892 10.892 0 0 1-.826-1.447l2.088-2.31a8.343 8.343 0 0 1-.385-2.222L4.89 14.808c.054-.563.164-1.107.3-1.643l3.084-.427a8.334 8.334 0 0 1 1.135-1.942l-1.183-2.9c.4-.39.83-.745 1.283-1.07l2.663 1.67a8.375 8.375 0 0 1 2.085-.75l.968-3c.278-.02.555-.042.837-.042c.282 0 .56.022.837.042l.976 3.028c.72.163 1.4.416 2.036.75l2.704-1.697c.455.326.887.68 1.285 1.07l-1.215 2.986c.428.564.793 1.18 1.068 1.845l3.185.44c.135.536.247 1.082.302 1.644l-2.867 1.517a8.53 8.53 0 0 1-.355 2.1l2.156 2.382z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Settingsalt(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 1.466C7.973 1.466 1.466 7.973 1.466 16c0 8.027 6.507 14.534 14.534 14.534c8.027 0 14.534-6.507 14.534-14.534c0-8.027-6.507-14.534-14.534-14.534m8.386 13.502a5.36 5.36 0 0 1-5.685 1.586l-7.187 8.266a2.113 2.113 0 0 1-3.187-2.775l7.198-8.274a5.348 5.348 0 0 1 6.137-7.497l-2.755 3.212l.9 2.62l2.723.53l2.76-3.22a5.339 5.339 0 0 1-.902 5.553z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shuffle(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.786 20.654c-.618-.195-1.407-.703-2.29-1.587c-.758-.742-1.54-1.698-2.34-2.74c-.192.255-.383.51-.575.77c-.523.708-1.058 1.423-1.603 2.126c1.904 2.31 3.88 4.578 6.81 4.952v2.7l7.555-4.36l-7.556-4.363v2.502zm-12.594-8.72c.756.74 1.538 1.696 2.34 2.738c.194-.262.39-.52.586-.788a113.2 113.2 0 0 1 1.592-2.11C11.678 9.31 9.577 6.867 6.314 6.75h-3.5v3.5h3.5c.655-.027 1.682.485 2.878 1.683zm12.594-1.593v2.536l7.556-4.363l-7.556-4.363v2.647c-1.904.22-3.425 1.348-4.75 2.644c-2.197 2.184-4.117 5.168-6.012 7.54c-1.867 2.437-3.74 3.887-4.712 3.77h-3.5v3.5h3.5c2.185-.03 3.88-1.266 5.34-2.693c2.194-2.184 4.116-5.167 6.01-7.538c1.543-2.017 3.084-3.34 4.124-3.68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Skull(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.947 11.14c0-5.174-3.98-9.406-10.613-9.406c-6.633 0-10.282 4.232-10.282 9.406s1.46 4.51 1.46 7.43c0 1.095-1.062.564-1.062 2.92c0 2.586 3.615 2.222 4.677 3.282c1.06 1.062.96 3.02.96 3.02s.2.795.565.562c0 0 .232.564.498.232c0 0 .265.563.53.1c0 0 .266.63.697.166c0 0 .43.63.93.133c0 0 .563.53 1.193.133c.63.397 1.194-.133 1.194-.133c.497.497.93-.133.93-.133c.43.465.694-.166.694-.166c.268.464.53-.1.53-.1c.267.332.5-.232.5-.232c.364.232.563-.563.563-.563s-.1-1.956.962-3.018c1.062-1.06 4.676-.696 4.676-3.283c0-2.355-1.06-1.825-1.06-2.92c0-2.92 1.46-2.256 1.46-7.43zm-15.614 9.852c-1.783.285-2.59-.215-2.785-1.492c-.508-3.328 2.555-3.866 4.08-3.683c.73.088 1.99.862 1.99 1.825c0 2.587-1.626 3.085-3.285 3.35m6.128 4.31c-.33 0-.86-.43-.894-1.226c-.033.796-.63 1.227-.96 1.227c-.333 0-.83-.33-.864-1.127c-.033-.796 1.028-4.013 1.792-4.013c.762 0 1.824 3.217 1.79 4.013s-.53 1.127-.863 1.127zm6.9-5.802c-.194 1.277-1.003 1.777-2.786 1.492c-1.658-.266-3.283-.763-3.283-3.35c0-.963 1.26-1.737 1.99-1.825c1.525-.183 4.59.355 4.08 3.683z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Skype(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28.777 18.438c.21-.948.318-1.934.318-2.944c0-7.578-6.144-13.722-13.724-13.722c-.798 0-1.583.07-2.345.2a7.957 7.957 0 0 0-11.097 10.76a13.65 13.65 0 0 0-.278 2.76c0 7.58 6.144 13.723 13.722 13.723c.86 0 1.7-.078 2.515-.23a7.957 7.957 0 0 0 10.888-10.547zm-6.365 3.613c-.635.9-1.573 1.61-2.79 2.116c-1.202.5-2.645.754-4.286.754c-1.97 0-3.624-.346-4.914-1.03a6.472 6.472 0 0 1-2.26-2.005c-.582-.842-.878-1.676-.878-2.48c0-.502.192-.938.573-1.295c.375-.354.857-.532 1.432-.532c.47 0 .877.14 1.208.422c.315.27.586.662.805 1.174c.242.558.508 1.027.788 1.397c.27.356.657.657 1.152.89c.497.236 1.168.355 1.992.355c1.135 0 2.064-.24 2.764-.72c.684-.466 1.016-1.026 1.016-1.712c0-.543-.173-.97-.53-1.303c-.372-.348-.864-.62-1.464-.807c-.623-.195-1.47-.404-2.518-.623c-1.424-.306-2.634-.668-3.596-1.076c-.984-.42-1.777-1-2.357-1.727c-.59-.736-.89-1.662-.89-2.75c0-1.036.314-1.97.933-2.776c.613-.8 1.51-1.422 2.663-1.848c1.14-.422 2.494-.635 4.027-.635c1.225 0 2.303.14 3.2.42c.905.283 1.67.663 2.268 1.13c.605.473 1.055.978 1.336 1.5c.284.53.43 1.058.43 1.566c0 .49-.19.937-.563 1.324c-.375.39-.85.59-1.408.59c-.51 0-.905-.125-1.183-.37c-.258-.227-.523-.58-.82-1.09c-.34-.65-.755-1.162-1.228-1.523c-.463-.35-1.232-.53-2.292-.53c-.984 0-1.784.198-2.38.59c-.57.374-.85.804-.85 1.313c0 .312.09.574.274.8c.195.237.47.446.818.62c.36.182.732.326 1.104.43c.382.105 1.02.262 1.9.465c1.11.238 2.13.506 3.033.793c.914.293 1.704.654 2.35 1.072c.655.43 1.177.98 1.546 1.635c.37.658.558 1.47.558 2.416a5.232 5.232 0 0 1-.962 3.063z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Slideshare(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28.952 12.795c-.956 1.062-5.073 2.41-5.604 2.41h-4.513c-.75 0-1.877.146-2.408.483c.06.054.122.108.18.163c.41.38 1.363.914 2.207.914c.397 0 .723-.115 1-.354c1.178-1.007 1.79-1.125 2.145-1.125c.42 0 .782.193.995.53c.4.627.106 1.446-.194 2.088c-.717 1.524-3.057 3.17-5.594 3.17h-.004c-.354 0-.7-.032-1.033-.098v3.25c0 .743 1.032 2.534 4.166 2.534s3.955-3.702 3.955-4.34v-4.51c2.23-1.17 4.513-1.806 5.605-3.896c1.027-1.964.052-2.28-.903-1.22zm-7.01 4.726c.796-1.698-.053-1.698-1.54-.424s-3.665.105-4.408-.585c-.743-.688-1.486-1.22-2.814-1.166c-1.328.053-4.46-.16-6.267-.585c-1.805-.426-4.895-3-5.15-2.336c-.266.69.21 1.168 1.168 2.335c.956 1.168 5.076 2.777 5.076 2.777v4.886c0 1.435 2.973 3.61 4.512 3.61s2.708-1.062 2.708-1.806v-4.512c2.55 1.33 5.92-.494 6.716-2.194zm-1.6-3.79a3.16 3.16 0 1 0 0-6.32a3.16 3.16 0 0 0 0 6.32m-8.323 0a3.16 3.16 0 1 0-.002-6.318a3.16 3.16 0 0 0 .001 6.318z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Smallgear(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M31.23 17.736c.063-.57.103-1.148.103-1.736s-.04-1.166-.104-1.737l-4.378-1.557c-.218-.716-.504-1.4-.85-2.05l1.992-4.192a15.582 15.582 0 0 0-2.458-2.46L21.343 6a11.162 11.162 0 0 0-2.05-.85L17.737.772C17.166.708 16.588.667 16 .667s-1.166.04-1.737.105L12.707 5.15c-.716.217-1.4.502-2.05.85L6.464 4.004c-.91.725-1.734 1.55-2.46 2.46L6 10.654c-.348.65-.633 1.335-.85 2.05l-4.378 1.56c-.064.57-.105 1.15-.105 1.737s.04 1.165.105 1.736l4.378 1.558c.217.715.502 1.4.85 2.05l-1.995 4.192c.725.91 1.55 1.733 2.46 2.458L10.654 26c.65.348 1.335.634 2.05.852l1.558 4.377c.57.063 1.148.103 1.737.103c.588 0 1.165-.04 1.736-.104l1.558-4.378c.715-.218 1.4-.504 2.05-.85l4.192 1.992c.91-.725 1.733-1.55 2.458-2.458L26 21.343c.348-.647.634-1.334.852-2.05l4.377-1.557zM16 20.87a4.871 4.871 0 1 1-.002-9.742A4.871 4.871 0 0 1 16 20.87"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Smile(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 1.466C7.973 1.466 1.466 7.973 1.466 16c0 8.027 6.507 14.534 14.534 14.534c8.027 0 14.534-6.507 14.534-14.534c0-8.027-6.507-14.534-14.534-14.534m4.73 5.91c.933 0 1.687 1.482 1.687 3.31S21.66 14 20.73 14c-.933 0-1.69-1.483-1.69-3.312s.758-3.313 1.69-3.313zm-9.626 0c.932 0 1.688 1.482 1.688 3.31S12.037 14 11.104 14s-1.688-1.483-1.688-3.312s.756-3.313 1.688-3.313zM16.02 26c-2.872 0-5.562-1.757-7.878-4.81c2.397 1.563 5.02 2.435 7.774 2.435c2.923 0 5.7-.98 8.215-2.734c-2.364 3.242-5.14 5.11-8.11 5.11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmileTwo(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 1.466C7.973 1.466 1.466 7.973 1.466 16c0 8.027 6.507 14.534 14.534 14.534c8.027 0 14.534-6.507 14.534-14.534c0-8.027-6.507-14.534-14.534-14.534m0 28.068C8.54 29.534 2.466 23.462 2.466 16C2.466 8.54 8.54 2.466 16 2.466c7.462 0 13.535 6.072 13.535 13.533c0 7.462-6.073 13.534-13.535 13.534zM11.104 14c.932 0 1.688-1.483 1.688-3.312s-.755-3.312-1.688-3.312s-1.688 1.483-1.688 3.312S10.172 14 11.104 14m9.625 0c.933 0 1.687-1.483 1.687-3.312s-.756-3.312-1.688-3.312c-.933 0-1.69 1.483-1.69 3.312S19.8 14 20.73 14zM8.142 21.19C10.458 24.242 13.148 26 16.02 26c2.97 0 5.746-1.868 8.11-5.11c-2.514 1.755-5.29 2.735-8.214 2.735c-2.752 0-5.376-.87-7.773-2.436z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Snow(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.372 6.912a7.254 7.254 0 0 0-7.248-7.08a7.238 7.238 0 0 0-6.208 3.518a4.164 4.164 0 0 0-2.01-.518A4.221 4.221 0 0 0 5.69 6.898a5.988 5.988 0 0 0-4.127 5.686c0 3.312 2.687 6 6 6v-.002h15.875c3.312 0 6-2.688 6-6a5.997 5.997 0 0 0-4.066-5.67m-1.936 9.672H7.562a4.007 4.007 0 0 1-4-4c-.003-1.983 1.45-3.62 3.35-3.933a.99.99 0 0 0 .656-.413a1.01 1.01 0 0 0 .163-.762a2.386 2.386 0 0 1-.044-.424a2.224 2.224 0 0 1 2.22-2.218c.647 0 1.217.278 1.633.73a.995.995 0 0 0 .926.31c.342-.065.626-.307.748-.63c.75-1.992 2.662-3.412 4.91-3.41a5.258 5.258 0 0 1 5.252 5.25c0 .16-.01.325-.026.496c-.05.518.305.984.814 1.08c1.86.344 3.273 1.965 3.27 3.922a4.008 4.008 0 0 1-3.998 4.004zm-6.77 7.506l1.12-1.12c.39-.39.39-1.024 0-1.415a1.002 1.002 0 0 0-1.415 0l-1.118 1.12l-1.12-1.12a1 1 0 0 0-1.414 1.416l1.118 1.118l-1.12 1.12a1.002 1.002 0 0 0 0 1.415c.195.188.45.293.707.293c.256 0 .512-.104.708-.293l1.12-1.12l1.12 1.12c.196.188.452.293.71.293a1 1 0 0 0 .706-1.707l-1.12-1.12zm8.453-2.273a1 1 0 0 0-1.416 0l-1.12 1.12l-1.12-1.12a.997.997 0 0 0-1.414 0a1.003 1.003 0 0 0 0 1.416l1.12 1.12l-1.12 1.118a1 1 0 0 0 1.414 1.414l1.12-1.118l1.12 1.118c.194.195.45.294.707.294c.257 0 .514-.1.71-.294a1.003 1.003 0 0 0 0-1.413L24 24.353l1.12-1.12a1.007 1.007 0 0 0 0-1.415zM9.333 23.953l1.12-1.12a1.003 1.003 0 0 0 0-1.413a.996.996 0 0 0-1.416 0l-1.12 1.12l-1.12-1.122a1 1 0 1 0-1.414 1.418l1.12 1.117l-1.12 1.118a.998.998 0 0 0 .707 1.708a1 1 0 0 0 .708-.293l1.12-1.12l1.12 1.122a1.001 1.001 0 0 0 1.416-1.416l-1.12-1.117z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Speaker(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.558 15.254a4.278 4.278 0 0 0 4.277-4.28a4.277 4.277 0 1 0-8.557.002a4.278 4.278 0 0 0 4.28 4.278m3.104-.03a5.233 5.233 0 0 1-3.103 1.03a5.236 5.236 0 0 1-3.106-1.03c-.75.625-1.498 1.52-2.11 2.623c-1.423 2.563-1.58 5.192-.35 5.874c.55.313 1.126.08 1.722-.495a10.585 10.585 0 0 0-.166 1.873c0 2.938 1.14 5.312 2.543 5.312c.846 0 1.265-.865 1.466-2.188c.2 1.31.62 2.188 1.46 2.188c1.397 0 2.545-2.375 2.545-5.312c0-.66-.062-1.29-.167-1.873c.598.574 1.174.812 1.725.496c1.228-.68 1.07-3.31-.353-5.873c-.61-1.105-1.358-1.998-2.108-2.623zm4.16-11.513l-1.415 1.415a8.263 8.263 0 0 1 0 11.704l1.413 1.41a10.249 10.249 0 0 0 3.015-7.264c0-2.834-1.152-5.404-3.014-7.265zm-3.534 10.997L17.7 16.12c1.32-1.317 2.136-3.137 2.136-5.144S19.02 7.15 17.702 5.83l-1.414 1.415a5.269 5.269 0 0 1 0 7.462zM21.94 1.59l-1.41 1.414a11.243 11.243 0 0 1 3.307 7.97c0 3.11-1.265 5.93-3.308 7.973l1.413 1.414a13.233 13.233 0 0 0 3.895-9.385c0-3.66-1.49-6.98-3.894-9.385z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Split(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.786 20.698c-1.792-.237-2.912-1.33-4.358-2.886c-.697-.75-1.428-1.577-2.324-2.32c1.396-1.164 2.41-2.518 3.483-3.502c1.01-.92 1.9-1.52 3.2-1.688v2.574l7.555-4.363l-7.556-4.363v2.652c-3.34.266-5.45 2.378-6.934 4.013c-.82.896-1.537 1.692-2.212 2.192c-.685.5-1.227.73-2.013.742H2.812v3.5h7.814c.786.01 1.33.24 2.017.742c1.02.743 2.095 2.18 3.552 3.568c1.312 1.258 3.162 2.46 5.592 2.65v2.663l7.556-4.36l-7.556-4.36z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 22.375L7.116 28.83l3.396-10.438l-8.883-6.458l10.978.002L16.002 1.5l3.39 10.434h10.982l-8.886 6.457l3.396 10.44z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarThree(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.44 28.18c-.418 0-.834-.13-1.188-.39l-5.75-4.248L9.75 27.79a2.003 2.003 0 0 1-3.087-2.243l2.26-6.783l-5.815-4.158a2.001 2.001 0 0 1 1.164-3.628h.014l7.15.056l2.157-6.816a1.999 1.999 0 0 1 3.811.001l2.155 6.816l7.15-.056h.015c.867 0 1.636.556 1.903 1.382c.27.83-.028 1.737-.74 2.246l-5.814 4.158l2.263 6.783a2.003 2.003 0 0 1-1.896 2.634z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarThreeOff(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28.63 12.36a1.996 1.996 0 0 0-1.902-1.383h-.015l-7.15.056l-2.155-6.816a2 2 0 0 0-3.812 0l-2.157 6.816l-7.15-.056h-.017a1.999 1.999 0 0 0-1.164 3.627l5.814 4.158l-2.26 6.783c-.276.828.017 1.74.723 2.25c.35.256.763.384 1.175.384c.418 0 .834-.132 1.19-.392l5.75-4.247l5.75 4.248a1.997 1.997 0 0 0 2.367.008a2.003 2.003 0 0 0 .72-2.25l-2.263-6.783l5.815-4.158a2 2 0 0 0 .74-2.246zm-8.918 5.636l2.73 8.184l-6.94-5.125L8.56 26.18l2.73-8.184l-7.02-5.018l8.627.066L15.5 4.82l2.603 8.225l8.627-.066l-7.018 5.016z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarTwo(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.615 4.928c.487-.986 1.284-.986 1.77 0l2.25 4.554c.486.986 1.775 1.923 2.864 2.08l5.023.73c1.09.16 1.335.917.547 1.685l-3.636 3.544c-.788.77-1.28 2.284-1.095 3.37l.858 5.003c.186 1.085-.46 1.553-1.433 1.04l-4.495-2.362c-.974-.51-2.567-.51-3.54 0l-4.496 2.364c-.974.512-1.618.044-1.432-1.04l.858-5.005c.186-1.086-.307-2.6-1.094-3.37L3.93 13.978c-.788-.768-.542-1.525.547-1.684l5.026-.73c1.088-.158 2.377-1.095 2.864-2.08l2.248-4.555z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarTwoOff(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m26.522 12.293l-5.024-.73c-1.09-.158-2.378-1.095-2.864-2.08l-2.25-4.555c-.486-.986-1.283-.986-1.77 0l-2.247 4.554c-.487.986-1.776 1.923-2.864 2.08l-5.026.73c-1.088.16-1.334.917-.547 1.685l3.637 3.544c.788.77 1.28 2.284 1.094 3.37l-.857 5.003c-.186 1.085.458 1.553 1.432 1.04l4.495-2.362c.974-.51 2.566-.51 3.54 0l4.496 2.364c.974.512 1.618.044 1.433-1.04l-.86-5.005c-.186-1.086.307-2.6 1.095-3.37l3.636-3.543c.787-.768.54-1.525-.548-1.684zm-4.485 3.796c-1.266 1.23-1.966 3.393-1.67 5.136l.514 2.984l-2.678-1.41c-.757-.395-1.715-.61-2.702-.61s-1.945.215-2.7.61l-2.68 1.408l.512-2.982c.297-1.743-.404-3.905-1.67-5.137l-2.167-2.113l2.995-.435c1.754-.255 3.592-1.59 4.373-3.175L15.5 7.652l1.342 2.716c.78 1.583 2.617 2.92 4.37 3.173l2.99.436l-2.165 2.113z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Staroff(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 22.375L7.116 28.83l3.396-10.438l-8.883-6.458l10.978.002L16.002 1.5l3.39 10.434h10.982l-8.886 6.457l3.396 10.44zm6.98 3.834l-2.665-8.206l6.98-5.062h-8.628L16 4.73l-2.666 8.205H4.708l6.98 5.07l-2.667 8.203L16 21.146z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Start(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.316 5.318L9.833 13.682V5.5H5.5v20h4.333v-8.182l14.483 8.364z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sticker(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 2c-1.042 0-1.916.376-2.57 1.087L2.895 13.137C2.302 13.785 2 14.58 2 15.5C2 22.943 8.056 29 15.5 29S29 22.943 29 15.5S22.943 2 15.5 2m0 26C8.596 28 3 22.404 3 15.5c0-3.452 5.24-2.737 7.5-5c2.262-2.26 1.548-7.5 5-7.5C22.404 3 28 8.597 28 15.5S22.404 28 15.5 28"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stop(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 5.5h20v20h-20z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stopsign(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.833 2.625H10.167l-7.542 7.542v10.666l7.542 7.542h10.666l7.542-7.542V10.167zm2.927 14.52c-.46.432-.496-.02-1.15.385c-.656.404-1.068.844-1.068.844l-.578.564s-1.2 1.422-1.59 1.717c-.392.29-.572.166-.572.166c-.133.23-.64.657-1.193 1.026c-.55.37-.99 1.354-.99 1.354l-.21 1.465s-1.596.02-3.095-.293c-1.5-.314-2.7-.982-2.7-.982l.656-1.45s-.58-2.226-.635-3.927c-.056-1.703.572-2.958.48-3.37c-.09-.412-.437-1.11-.522-1.57a67.574 67.574 0 0 1-.327-2.237c-.02-.202-.28-2.232.614-2.204c.89.027.648 1.388.725 2.246c.077.857 1.13 3.25 1.297 3.123c.167-.124-.056-2.397.006-2.837c.063-.44.182-2.315.293-2.747c.113-.433.106-1.778.936-1.66c.83.118.606 1.332.488 1.813c-.118.48.02 1.596.07 2.3c.048.705.112 2.357.112 2.357l.404-.042s.913-5 .976-5.44c.062-.438.182-1.617.858-1.47c1.117.24.516 1.966.516 1.966l-.486 5.51s.166.224.492.02c.33-.2 1.312-3.25 1.46-3.926c.145-.676.083-2.678 1.144-2.428c1.06.252.453 2.755.124 4.353c-.327 1.596-1.03 3.39-1.03 3.39l-.433 1.63l.377.852l.412-.092l.897-.717l.614-.46c.614-.46 1.765-.496 2.045-.356c.278.137 1.046.696.585 1.128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stopwatch(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27.216 18.533a11.69 11.69 0 0 0-4.253-9.032l.733-.997l.482.354a.45.45 0 0 0 .628-.097l.442-.604a.45.45 0 0 0-.097-.628L23.1 6.022a.453.453 0 0 0-.63.097l-.44.604a.45.45 0 0 0 .096.628l.483.354l-.737 1.003a11.595 11.595 0 0 0-4.296-1.7a3.468 3.468 0 0 0 1.403-2.778h-.002a3.48 3.48 0 1 0-6.96 0c0 1.14.557 2.144 1.402 2.778c-1.55.277-2.99.857-4.265 1.68l-.73-1.006l.483-.353a.45.45 0 0 0 .098-.627l-.44-.604a.454.454 0 0 0-.63-.1L5.882 7.5a.452.452 0 0 0-.098.628l.442.604c.145.2.428.244.627.1l.484-.355l.73 1a11.685 11.685 0 0 0-4.283 9.056c0 6.47 5.246 11.716 11.72 11.716c6.47 0 11.715-5.244 11.717-11.717zM12.918 4.23a2.582 2.582 0 0 1 5.162 0c0 1.193-.81 2.185-1.907 2.483v-1.77h.6c.246 0 .45-.203.45-.45v-.747a.453.453 0 0 0-.45-.45h-2.545a.452.452 0 0 0-.45.45v.75c0 .245.203.448.45.448h.6v1.77c-1.1-.3-1.91-1.29-1.91-2.483zM15.5 27.555a9.03 9.03 0 0 1-9.022-9.02A9.034 9.034 0 0 1 15.5 9.512a9.034 9.034 0 0 1 9.02 9.023a9.031 9.031 0 0 1-9.02 9.02m0-15.416a.86.86 0 1 0 .001-1.719a.86.86 0 0 0-.001 1.719m0 12.79a.86.86 0 1 0-.001 1.719a.86.86 0 0 0 .001-1.719m-2.882-13.11a.86.86 0 1 0-1.491.862l2.867 6.722a1.741 1.741 0 0 0 2.375.637a1.742 1.742 0 0 0 .634-2.375l-4.386-5.847zm6.08 12.252a.86.86 0 1 0 .86 1.491a.86.86 0 1 0-.86-1.49zm-9.91-8.42a.86.86 0 1 0 .861-1.491a.86.86 0 1 0-.863 1.49zm13.427 5.763a.862.862 0 0 0-.863 1.491a.86.86 0 1 0 .861-1.491zM9.107 18.53a.863.863 0 0 0-1.723 0a.862.862 0 0 0 1.723.002zm12.79 0c0 .478.383.863.858.86c.476.003.862-.38.862-.858s-.387-.86-.862-.862a.86.86 0 0 0-.86.86zm-13.11 2.883a.861.861 0 1 0 .862 1.491a.862.862 0 0 0-.861-1.492zm12.565-7.256a.861.861 0 1 0 1.178.315a.865.865 0 0 0-1.178-.315m-9.048 9.91a.863.863 0 0 0-.864 1.493a.862.862 0 0 0 .864-1.493m6.394-11.075a.86.86 0 1 0 .858-1.489a.86.86 0 0 0-.857 1.488z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sun(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.502 7.504a7.87 7.87 0 0 0-7.873 7.873a7.87 7.87 0 1 0 15.742-.001a7.868 7.868 0 0 0-7.87-7.873zm0 13.746a5.884 5.884 0 0 1-5.874-5.872a5.884 5.884 0 0 1 5.874-5.874a5.881 5.881 0 0 1 5.87 5.874a5.88 5.88 0 0 1-5.87 5.872m0-14.273a1 1 0 0 0 1-1V1.124a1.001 1.001 0 0 0-2 .002v4.85c0 .55.447 1 1 1zm3.213.638a.995.995 0 0 0 1.307-.541l1.856-4.483a.999.999 0 0 0-.54-1.305a1.005 1.005 0 0 0-1.31.54L18.175 6.31a.999.999 0 0 0 .54 1.305m2.725 1.82a.997.997 0 0 0 1.414 0l3.43-3.433a1 1 0 1 0-1.415-1.413l-3.43 3.43a1.003 1.003 0 0 0 0 1.416zm1.823 2.725a.996.996 0 0 0 1.306.539l4.48-1.858a1 1 0 1 0-.767-1.847l-4.48 1.857c-.51.212-.752.798-.54 1.308zm6.49 2.21l-4.852.002c-.55 0-1 .448-.997 1c0 .554.447 1 .998 1l4.853-.002a1.001 1.001 0 0 0 0-2zm-.7 5.53l-4.48-1.855a1 1 0 0 0-1.308.54c-.21.512.03 1.097.54 1.31l4.483 1.853a1.001 1.001 0 1 0 .767-1.849zm-6.193 1.41a1 1 0 0 0-1.414 1.414l3.434 3.43a.998.998 0 0 0 1.414-.001c.39-.392.39-1.025 0-1.415zm-2.83 2.363a1 1 0 0 0-1.848.767l1.86 4.48a.999.999 0 0 0 1.308.54c.51-.21.752-.797.54-1.31zm-4.518.103a1 1 0 0 0-1 1l.004 4.85a1 1 0 0 0 1 1c.553 0 1-.448.998-1l-.003-4.852c0-.55-.448-1-.998-.998zm-3.216-.636a1 1 0 0 0-1.306.543l-1.852 4.483a1.003 1.003 0 0 0 .926 1.383a1 1 0 0 0 .924-.618l1.853-4.485a1 1 0 0 0-.544-1.305zM9.57 21.325a1 1 0 0 0-1.415.002L4.73 24.76a1 1 0 1 0 1.415 1.414l3.427-3.434a1 1 0 0 0-.002-1.415m-1.824-2.72a1 1 0 0 0-1.307-.54l-4.48 1.86a1.001 1.001 0 0 0 .767 1.847l4.48-1.86c.508-.215.75-.8.54-1.31zM7.1 15.39a1 1 0 0 0-1-1l-4.85.007a1 1 0 1 0 0 1.998l4.852-.006c.552 0 1-.45.998-1zm-5.156-4.52l4.485 1.85a1 1 0 1 0 .761-1.849L2.71 9.02a1.002 1.002 0 0 0-.764 1.849zm6.193-1.42a.998.998 0 0 0 1.415-.003a.998.998 0 0 0-.005-1.414L6.113 4.61a.999.999 0 0 0-1.414.002a.999.999 0 0 0 0 1.414zm2.827-2.366a1 1 0 0 0 1.847-.77l-1.863-4.48a1 1 0 0 0-1.848.77z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Svg(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M31.274 15.99a4.48 4.48 0 0 0-4.478-4.48A4.477 4.477 0 0 0 15.984.7a4.478 4.478 0 0 0-4.48 4.477a4.479 4.479 0 1 0-6.331 15.291a4.48 4.48 0 1 0 15.289 6.333A4.478 4.478 0 0 0 31.274 15.99m-7.69 5.22h.046a2.424 2.424 0 1 1-2.424 2.424v-.047l-3.54-3.542v5.01a2.425 2.425 0 1 1-4.107 1.746c0-.686.286-1.305.743-1.745v-5.01l-3.54 3.543v.047a2.424 2.424 0 1 1-2.425-2.426h.045l3.542-3.54H6.92a2.423 2.423 0 0 1-4.171-1.681a2.424 2.424 0 0 1 4.169-1.681h5.007l-3.542-3.543H8.34a2.423 2.423 0 1 1 0-4.845a2.423 2.423 0 0 1 2.424 2.424v.046l3.54 3.542V6.924a2.426 2.426 0 1 1 3.364 0v5.008l3.54-3.542v-.046a2.425 2.425 0 1 1 2.425 2.424h-.046l-3.54 3.54h5.007a2.423 2.423 0 1 1 .002 3.361h-5.007l3.54 3.54z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.263 2.826h-6.36l-5.2 5.202v6.36L18.404 30.09l11.56-11.562zM6.495 8.86a1.581 1.581 0 0 1 0-2.24c.62-.62 1.622-.62 2.24 0c.62.62.62 1.62 0 2.24c-.618.62-1.62.62-2.24 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Takeoff(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.27 19.267s9.375-1.98 16.074-8.68c0 0 1.395-1.34-1.338-1.34c-2.305 0-5.6 2.438-5.6 2.438l-9.137-1.42l-1.77 1.77l4.983 2.41l-3 2.035l-2.572-1.285l-1.82.857s1.93 2.01 4.18 3.215m-7.02 3.84v1.997h24.5v-1.998H3.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Talke(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4.938c-7.732 0-14 4.7-14 10.5c0 1.98.74 3.833 2.016 5.414L2 25.272l5.613-1.44c2.34 1.316 5.237 2.106 8.387 2.106c7.732 0 14-4.7 14-10.5s-6.268-10.5-14-10.5m.982 16.437h-1.97v-1.89h1.97zm0-3.906v.624h-1.97v-.77c0-2.32 2.642-2.688 2.642-4.336c0-.752-.672-1.33-1.553-1.33c-.91 0-1.712.673-1.712.673l-1.12-1.392s1.104-1.153 3.01-1.153c1.81 0 3.49 1.12 3.49 3.01c0 2.642-2.786 2.946-2.786 4.674z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Talkq(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4.938c-7.732 0-14 4.7-14 10.5c0 1.98.74 3.833 2.016 5.414L2 25.272l5.613-1.44c2.34 1.316 5.237 2.106 8.387 2.106c7.732 0 14-4.7 14-10.5s-6.268-10.5-14-10.5m.868 16.437h-1.97v-1.89h1.97zm-.096-3.28h-1.777l-.176-8.084h2.112l-.16 8.084z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Taxi(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28.59 10.78h-2.242a.496.496 0 0 0-.333.134c-.716-1.143-1.457-2.058-2.032-2.633c-.575-.574-1.98-.983-3.732-1.228V5.03c-1.54-.198-3.098-.25-4.267-.25c-1.16 0-2.703.05-4.233.246v2.022c-1.77.244-3.188.654-3.768 1.233c-.572.574-1.308 1.483-2.02 2.618a.489.489 0 0 0-.31-.117H3.41c-.275 0-.5.226-.5.5v1.01c0 .274.22.54.49.593l1.36.26c-1.174 2.618-1.866 5.876-.778 9.14v1.937c0 .553.14 1 .313 1h2.562c.173 0 .313-.447.313-1v-1.584c2.298.22 5.55.46 8.812.46c3.232 0 6.52-.236 8.814-.454v1.578c0 .553.14 1 .312 1h2.562c.172 0 .312-.447.312-1l-.002-1.938c1.087-3.26.397-6.516-.775-9.134l1.392-.265a.63.63 0 0 0 .49-.594v-1.01a.498.498 0 0 0-.497-.5zM7.107 18.907a1.813 1.813 0 0 1 0-3.624a1.813 1.813 0 0 1-.001 3.624zm-1.524-5.19C6.543 11.52 7.88 9.8 8.69 8.988c.584-.585 3.34-1.207 7.292-1.207c3.953 0 6.708.623 7.293 1.208c.81.81 2.146 2.53 3.106 4.728c-2.132.236-6.285-.31-10.398-.31s-8.266.546-10.4.31m19.274 5.19a1.813 1.813 0 0 1 0-3.624a1.813 1.813 0 0 1-.001 3.624z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Temp(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.5 19.508V8.626h-4v10.88c-1.403.728-2.374 2.18-2.374 3.87a4.376 4.376 0 0 0 8.751.001c0-1.69-.97-3.142-2.376-3.868zm3-14.26c0-2.756-2.244-5-5-5s-5 2.245-5 5v12.727a7.313 7.313 0 0 0-2.375 5.4c0 4.066 3.31 7.377 7.376 7.377s7.375-3.31 7.375-7.377c0-2.086-.878-4.03-2.375-5.402V5.25zm.375 18.13A5.379 5.379 0 0 1 15.5 28.75a5.38 5.38 0 0 1-5.373-5.373c0-1.795.896-3.443 2.376-4.438V5.25c0-1.653 1.343-3 2.997-3s3 1.346 3 3v13.69a5.332 5.332 0 0 1 2.375 4.437zm1.21-14.752l4.5 2.598V6.03z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Terminal(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.25 6.47v19.06h25.5V6.47zm7.095 5.043l-4.33 1.926v-1l3.123-1.288v-.018L6.014 9.848v-1l4.33 1.927zM16.04 14.6h-5.05v-.88h5.05z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Thunder(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.37 7.306a7.252 7.252 0 0 0-7.247-7.08a7.24 7.24 0 0 0-6.208 3.518a4.163 4.163 0 0 0-2.01-.517a4.22 4.22 0 0 0-4.216 4.066a5.987 5.987 0 0 0-4.128 5.686c0 3.31 2.687 6 6 6v-.003h5.27l-2.165 3.398l1.977-.41L10 30.874l9.138-10.102L17 21l2.167-2.023h4.27c3.31 0 6-2.688 6-6a5.995 5.995 0 0 0-4.066-5.67zm-1.934 9.673H7.56a4.01 4.01 0 0 1-4-4.002c-.002-1.982 1.45-3.618 3.35-3.93a.998.998 0 0 0 .657-.415c.155-.223.212-.497.163-.763a2.284 2.284 0 0 1-.045-.424a2.224 2.224 0 0 1 2.22-2.218c.647 0 1.217.278 1.633.73a.997.997 0 0 0 1.675-.32c.75-1.992 2.662-3.412 4.91-3.41a5.258 5.258 0 0 1 5.252 5.25c0 .16-.01.325-.027.496c-.05.517.305.984.815 1.08c1.86.343 3.274 1.965 3.27 3.922a4.005 4.005 0 0 1-3.997 4.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ticket(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.338 6.713a2 2 0 0 1-.967-2.658L16.58 2.75L6.21 24.948l2.793 1.305c.468-1 1.658-1.434 2.66-.966s1.433 1.657.965 2.658l2.793 1.305L25.79 7.052l-2.794-1.305a2 2 0 0 1-2.658.966"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Train(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.78 21.382c.69-.69.608-4.31.608-4.31c0-2.21-.663-10.335-1.602-12.104c-.94-1.768-6.19-3.26-8.787-3.26S8.15 3.2 7.212 4.968C6.273 6.738 5.61 14.86 5.61 17.07c0 0-.083 3.62.607 4.312c.455.455 2.205 1.58 3.663 2.423L6.136 30.29h2.31l.673-1.165h13.758l.673 1.166h2.312l-3.744-6.485c1.458-.842 3.208-1.968 3.663-2.423zm-1.96-4.587a.857.857 0 1 1 .854-.857a.855.855 0 0 1-.856.855zm-3.4-9.836h3.77c.325 1.564.615 3.98.806 6.133h-4.574V6.96zm-6.003-3.835h3.166v1.333h-3.166zM7.812 6.96h3.768v6.133H7.005c.19-2.154.48-4.57.807-6.134zm.37 9.835a.856.856 0 1 1-.002-1.708a.856.856 0 0 1 .002 1.706zm1.16-.857a.855.855 0 1 1 1.713-.001a.855.855 0 0 1-1.714.001zm.933 11.187l.99-1.715h9.47l.99 1.715zm10.67-11.187a.857.857 0 1 1 1.711.001a.857.857 0 0 1-1.71-.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.826 5.75l.396 1.188c1.54.575 2.59 1.44 2.59 2.626c0 2.405-4.31 3.498-8.313 3.498c-4.004 0-8.312-1.093-8.312-3.498c0-1.272 1.21-2.174 2.938-2.746l.388-1.165C8.07 6.3 6.187 7.53 6.187 9.563v2.264c0 1.224.685 2.155 1.76 2.845l.395 9.265c0 1.38 3.274 2.5 7.312 2.5s7.313-1.12 7.313-2.5l.405-9.493c.885-.664 1.438-1.52 1.438-2.617V9.562c.002-1.937-1.71-3.142-3.984-3.812m-9.733 18.377c-.476-.286-1.022-.846-1.166-1.237c-1.007-2.76-.73-4.92-.53-7.51c.748.28 1.58.492 2.45.643c-.215 2.658-.43 4.923.004 7.828c.066.428-.283.56-.757.277zm6.126.202c-.02.444-.692.855-1.518.855c-.828 0-1.498-.413-1.517-.858c-.126-2.996-.032-5.322.068-8.04c.418.023.835.038 1.246.038c.542 0 1.096-.02 1.65-.06c.1 2.73.196 5.06.07 8.064zm4.256-1.438c-.143.392-.69.95-1.165 1.235c-.473.284-.816.15-.753-.276c.437-2.93.214-5.208-.005-7.896c.88-.174 1.708-.417 2.44-.73c.202 2.66.51 4.852-.516 7.668zM11.338 9.512a1.006 1.006 0 0 0 1.268-.633h-.002l.77-2.317h4.56l.772 2.316a1.003 1.003 0 0 0 1.265.632a1 1 0 0 0 .634-1.265l-1.002-3a1 1 0 0 0-.945-.684h-6.002c-.428 0-.812.275-.948.683l-1 3c-.175.524.108 1.09.63 1.266z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tshirt(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.1 4.04a4.576 4.576 0 0 1-4.24 2.86a4.576 4.576 0 0 1-4.24-2.86L1.238 8.44l2.92 6.884l3.21-1.36V28h17.098V14.015l3.093 1.312l2.92-6.884z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitter(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.295 22.567h-7.213c-2.125 0-4.103-2.215-4.103-4.736v-1.828h11.23a3.287 3.287 0 0 0 3.292-3.28a3.287 3.287 0 0 0-3.29-3.283H11.978V6.197c0-1.835-1.376-3.323-3.193-3.323c-1.816 0-3.29 1.488-3.29 3.323V17.83c0 6.23 4.685 11.275 10.476 11.275h7.21c1.82 0 3.32-1.463 3.32-3.298s-1.39-3.24-3.208-3.24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitterbird(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M26.492 9.493a8.57 8.57 0 0 1-2.473.678a4.351 4.351 0 0 0 1.892-2.38a8.647 8.647 0 0 1-2.734 1.043a4.308 4.308 0 0 0-7.336 3.927a12.227 12.227 0 0 1-8.874-4.5a4.29 4.29 0 0 0-.583 2.165a4.3 4.3 0 0 0 1.915 3.583a4.286 4.286 0 0 1-1.95-.538v.053a4.305 4.305 0 0 0 3.454 4.222c-.36.1-.74.147-1.134.147c-.278 0-.547-.023-.81-.076a4.311 4.311 0 0 0 4.022 2.99a8.692 8.692 0 0 1-6.374 1.78a12.158 12.158 0 0 0 6.6 1.938c7.92 0 12.248-6.562 12.248-12.25c0-.187-.002-.372-.01-.557a8.63 8.63 0 0 0 2.146-2.224"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Umbrella(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.784 26.99c0 1.24-1.33 1.697-1.835 1.697c-.505 0-1.537-.413-1.65-1.812a.645.645 0 0 0-.645-.642a.64.64 0 0 0-.64.642c.044 1.056.755 3.052 2.934 3.052c2.43 0 3.165-1.882 3.165-3.144v-8.176l-1.328-.024l-.003 8.408zm.8-17.186c-6.807 0-11.084 4.86-11.587 8.326a3.516 3.516 0 0 1 2.89-1.51c1.197 0 2.22.582 2.855 1.495a3.519 3.519 0 0 1 2.88-1.495c1.2 0 2.26.6 2.896 1.517c.635-.917 1.83-1.517 3.03-1.517c1.19 0 2.24.59 2.88 1.495a3.436 3.436 0 0 1 2.854-1.495c1.197 0 2.254.597 2.89 1.51c-.503-3.467-4.78-8.326-11.588-8.326m-.85-2.68v2.082h1.322v-2.08a.661.661 0 0 0-1.323-.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Undo(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.98 9.073V6.817L.876 13.807l12.106 6.99v-2.422c3.286-.002 9.053.28 9.053 2.27c0 2.78-6.023 4.262-6.023 4.262v2.132s13.53.462 13.53-9.824c0-8.082-11.588-8.385-16.56-8.143z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlock(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.875 15.334v-4.876c0-4.894-3.98-8.875-8.875-8.875s-8.875 3.98-8.875 8.875v.375h3.5v-.375c0-2.964 2.41-5.375 5.375-5.375s5.375 2.41 5.375 5.375v4.876H5.042v15.083h21.916V15.334zm-6.603 11.622h-4.545l1.222-3.667a2.37 2.37 0 0 1-1.325-2.12a2.375 2.375 0 1 1 4.75 0c0 .932-.542 1.73-1.324 2.12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Usb(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 1.667L13.745 4.74h1.252v15.71L11.8 17.39a1.63 1.63 0 0 1-.358-.95v-2.57c.59-.21 1.02-.773 1.02-1.443a1.53 1.53 0 0 0-1.522-1.538c-.84 0-1.52.688-1.52 1.537c0 .67.426 1.234 1.018 1.444v2.54c0 .688.372 1.41.81 1.868c-.012-.013-.026-.025 0 0c.012.01 3.393 3.245 3.393 3.245c.206.26.35.6.358.95v1.777a2.562 2.562 0 0 0-2.036 2.517c0 1.418 1.137 2.566 2.54 2.566c1.402 0 2.54-1.148 2.54-2.566c0-1.244-.876-2.28-2.04-2.517v-5.62c.01-.35.153-.69.36-.95c0 0 3.38-3.234 3.39-3.245c.028-.026.013-.013 0 0c.44-.46.812-1.18.812-1.87V10.12h1.02V7.046h-3.04v3.075h1.017v2.477a1.65 1.65 0 0 1-.358.952l-3.198 3.06V4.74h1.254z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func User(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.77 12.364s.85-3.51 0-4.7c-.85-1.188-1.188-1.98-3.057-2.547s-1.188-.454-2.547-.396c-1.36.058-2.492.793-2.492 1.19c0 0-.85.056-1.188.396c-.34.34-.906 1.924-.906 2.32s.283 3.06.566 3.625l-.337.114c-.284 3.283 1.13 3.68 1.13 3.68c.51 3.058 1.02 1.756 1.02 2.548s-.51.51-.51.51s-.452 1.245-1.584 1.698c-1.132.452-7.416 2.886-7.927 3.396c-.512.51-.454 2.888-.454 2.888h26.947s.06-2.377-.452-2.888c-.51-.51-6.795-2.944-7.927-3.396c-1.132-.453-1.584-1.698-1.584-1.698s-.51.282-.51-.51s.51.51 1.02-2.548c0 0 1.413-.397 1.13-3.68h-.34z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Users(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.053 20.8c-1.132-.453-1.584-1.698-1.584-1.698s-.51.282-.51-.51s.51.51 1.02-2.548c0 0 1.413-.397 1.13-3.68h-.34s.85-3.51 0-4.7c-.85-1.188-1.188-1.98-3.057-2.547s-1.188-.454-2.547-.396c-1.36.058-2.492.793-2.492 1.19c0 0-.85.056-1.188.396c-.34.34-.906 1.924-.906 2.32s.283 3.06.566 3.625l-.337.114c-.284 3.283 1.13 3.68 1.13 3.68c.51 3.058 1.02 1.756 1.02 2.548s-.51.51-.51.51s-.452 1.245-1.584 1.698c-1.132.452-7.416 2.886-7.927 3.396c-.512.51-.454 2.888-.454 2.888H29.43s.06-2.377-.452-2.888c-.51-.51-6.795-2.944-7.927-3.396zm-12.47-.172c-.1-.18-.148-.31-.148-.31s-.432.24-.432-.432s.432.432.864-2.16c0 0 1.2-.335.96-3.118h-.29s.144-.59.238-1.334a10.01 10.01 0 0 1 .037-.996l.038-.426c-.02-.492-.107-.94-.312-1.226c-.72-1.007-1.008-1.68-2.59-2.16c-1.584-.48-1.01-.384-2.16-.335c-1.152.05-2.112.672-2.112 1.01c0 0-.72.047-1.008.335c-.27.27-.705 1.462-.757 1.885v.28c.048.654.26 2.45.47 2.873l-.286.096c-.24 2.782.96 3.118.96 3.118c.43 2.59.863 1.488.863 2.16s-.432.43-.432.43s-.383 1.058-1.343 1.44l-.232.092v5.234h.575c-.03-1.278.077-2.927.746-3.594c.357-.355 1.524-.94 6.353-2.862zm22.33-9.056c-.04-.378-.127-.715-.292-.946c-.718-1.008-1.007-1.68-2.59-2.16c-1.583-.48-1.007-.384-2.16-.335c-1.15.05-2.11.672-2.11 1.01c0 0-.72.047-1.008.335c-.27.272-.71 1.472-.758 1.89h.033l.08.914c.02.23.022.435.027.644c.09.666.21 1.35.33 1.59l-.286.095c-.24 2.782.96 3.118.96 3.118c.432 2.59.863 1.488.863 2.16s-.43.43-.43.43s-.054.143-.164.34c4.77 1.9 5.927 2.48 6.28 2.833c.67.668.774 2.316.745 3.595h.48V21.78l-.05-.022c-.96-.383-1.344-1.44-1.344-1.44s-.433.24-.433-.43s.433.43.864-2.16c0 0 .804-.23.963-1.84V14.66c0-.018 0-.033-.003-.05h-.29s.216-.89.293-1.862z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Video(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27.188 4.875V5.97h-4.5V4.874H8.062V5.97h-4.5V4.874h-1v21.25h1V25.03h4.5v1.095h14.625V25.03h4.5v1.095h1.25V4.875h-1.25zM8.062 23.72h-4.5v-3.126h4.5v3.125zm0-4.44h-4.5v-3.124h4.5zm0-4.436h-4.5V11.72h4.5zm0-4.438h-4.5V7.28h4.5zm3.185 10.184V9.754l9.382 5.418zm15.94 3.13h-4.5v-3.126h4.5v3.125zm0-4.44h-4.5v-3.124h4.5zm0-4.436h-4.5V11.72h4.5zm0-4.438h-4.5V7.28h4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func View(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 8.286C8.454 8.286 2.5 16 2.5 16s5.954 7.715 13.5 7.715c5.77 0 13.5-7.715 13.5-7.715S21.77 8.286 16 8.286m0 12.52c-2.65 0-4.807-2.156-4.807-4.806S13.35 11.193 16 11.193S20.807 13.35 20.807 16S18.65 20.807 16 20.807zm0-7.612a2.806 2.806 0 1 0 0 5.611a2.806 2.806 0 0 0 0-5.611"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vim(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m25.012 10.44l4.57-4.645a.607.607 0 0 0 .174-.423V3.134a.605.605 0 0 0-.177-.427l-.605-.602a.592.592 0 0 0-.42-.177l-9.646-.086a.597.597 0 0 0-.5.257l-.603.86a.54.54 0 0 0-.076.154L15.588.958L13.94 2.624l-.446-.497a.6.6 0 0 0-.438-.2l-9.732-.17a.6.6 0 0 0-.437.175l-.603.603a.605.605 0 0 0-.177.427v2.325c0 .164.066.32.183.434l.657.636c.11.105.253.165.405.17l.285.006l.007 6.512L1.117 15.6l2.533 2.534l.008 8.084c0 .16.065.314.177.427l.86.86c.113.112.27.177.428.177h2.67c.16 0 .317-.064.43-.18l2.378-2.418l4.9 4.9l14.47-14.558l-4.958-4.986zM9.747 24.232L7.54 26.474H5.37l-.51-.51l-.006-6.624l-.008-7.515l-.006-5.882a.603.603 0 0 0-.588-.603l-.637-.014l-.304-.295V3.21l.245-.244l9.215.163l.32.353l.125.14v1.42l-.352.362h-.606c-.33 0-.6.266-.603.597l-.076 7.203c0 .244.142.463.366.56a.6.6 0 0 0 .657-.12l7.495-7.235a.612.612 0 0 0 .14-.66a.6.6 0 0 0-.56-.377h-.48l-.296-.38V3.497l.312-.445l9.083.082l.252.252v1.743l-4.39 4.458l-14.41 14.645z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeOne(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.998 12.127v7.896h4.495l6.73 5.526l.003-18.95l-6.73 5.527zm13.808-.908a1.002 1.002 0 0 0-1.415 0c-.39.392-.39 1.025.003 1.417v-.002A4.741 4.741 0 0 1 18.789 16c0 1.317-.53 2.498-1.393 3.362a.996.996 0 0 0-.002 1.415a.998.998 0 0 0 1.415 0a6.74 6.74 0 0 0 1.98-4.776c0-1.864-.76-3.56-1.982-4.78z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeThree(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.998 12.127v7.896h4.495l6.73 5.526l.003-18.95l-6.73 5.527zm13.808-.908a1.002 1.002 0 0 0-1.415 0c-.39.392-.39 1.025.003 1.417v-.002A4.741 4.741 0 0 1 18.789 16c0 1.317-.53 2.498-1.393 3.362a.996.996 0 0 0-.002 1.415a.998.998 0 0 0 1.415 0a6.74 6.74 0 0 0 1.98-4.776c0-1.864-.76-3.56-1.982-4.78zM21.1 8.924a.998.998 0 1 0-1.412 1.414a7.97 7.97 0 0 1 2.344 5.66a7.984 7.984 0 0 1-2.342 5.66a.996.996 0 0 0 0 1.413a.997.997 0 0 0 1.415 0a9.973 9.973 0 0 0 2.927-7.073c0-2.76-1.12-5.268-2.93-7.075zm2.18-2.18a1 1 0 0 0-1.412 1.416h-.002A11.037 11.037 0 0 1 25.114 16c0 3.063-1.24 5.828-3.246 7.838a1 1 0 1 0 1.416 1.413a13.06 13.06 0 0 0 3.83-9.25c0-3.61-1.467-6.89-3.834-9.254z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeTwo(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.998 12.127v7.896h4.495l6.73 5.526l.003-18.95l-6.73 5.527zm13.808-.908a1.002 1.002 0 0 0-1.415 0c-.39.392-.39 1.025.003 1.417v-.002A4.741 4.741 0 0 1 18.789 16c0 1.317-.53 2.498-1.393 3.362a.996.996 0 0 0-.002 1.415a.998.998 0 0 0 1.415 0a6.74 6.74 0 0 0 1.98-4.776c0-1.864-.76-3.56-1.982-4.78zM21.1 8.924a.998.998 0 1 0-1.412 1.414a7.97 7.97 0 0 1 2.344 5.66a7.984 7.984 0 0 1-2.342 5.66a.996.996 0 0 0 0 1.413a.997.997 0 0 0 1.415 0a9.973 9.973 0 0 0 2.927-7.073c0-2.76-1.12-5.268-2.93-7.075z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeZero(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.998 12.127v7.896h4.495l6.73 5.526l.003-18.95l-6.73 5.527z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Warning(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m29.225 23.567l-3.778-6.542c-1.14-1.972-3.002-5.2-4.14-7.172l-3.78-6.542c-1.14-1.972-3.002-1.972-4.14 0L9.608 9.854l-4.143 7.172l-3.777 6.542c-1.14 1.974-.207 3.587 2.07 3.587H27.15c2.28 0 3.21-1.613 2.073-3.587zm-12.69 1.013h-2.24v-2.15h2.24zm-.107-3.736h-2.023l-.2-9.204h2.406l-.182 9.204z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wheelchair(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.373 19.85c0 4.08-3.318 7.397-7.398 7.397S5.577 23.93 5.577 19.85a7.39 7.39 0 0 1 3.073-5.997l-.25-2.21c-2.876 1.61-4.826 4.684-4.826 8.207c0 5.184 4.217 9.4 9.4 9.4c4.396 0 8.094-3.03 9.118-7.11l-1.722-2.41zM11.768 6.534a2.392 2.392 0 1 0 0-4.784a2.392 2.392 0 0 0 0 4.784m15.42 16.143l-5.367-7.505c-.28-.393-.748-.58-1.225-.538c-.035-.003-.07-.006-.106-.006h-6.133l-.152-1.335h4.557a.96.96 0 0 0 0-1.918h-4.776l-.25-2.192a2.335 2.335 0 1 0-4.641.527L9.8 15.9a2.331 2.331 0 0 0 2.33 2.07l.013.002h8.023l4.603 6.436c.438.615 1.337.727 2.006.248c.666-.478.852-1.364.412-1.98z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Windows(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.023 17.484c-1.732-.205-3.022-.908-4.212-1.7l-.558.278l-2.578 8.924c1.217.805 2.905 1.707 4.682 1.914c2.686.312 5.56-.744 6.39-1.195l2.618-9.06l-.56-.28s-2.61 1.492-5.78 1.12zm-5.6-2.66c-1.266-.87-2.577-1.65-4.374-1.815a9.897 9.897 0 0 0-.926-.043c-3.01 0-4.948 1.347-4.948 1.347L1.61 23.19l.527.282a10.08 10.08 0 0 1 5.09-.984c1.665.113 2.92.78 4.117 1.53l.507-.26l2.574-8.933zm-4.222-2.73c1.665.114 2.922.78 4.118 1.533l.51-.26L17.4 4.43c-1.27-.87-2.58-1.652-4.377-1.815a9.816 9.816 0 0 0-.924-.042c-3.012 0-4.95 1.347-4.95 1.347l-2.566 8.878l.526.282a10.114 10.114 0 0 1 5.09-.986zM28.78 5.97c0 .002-2.61 1.493-5.78 1.12c-1.734-.204-3.023-.907-4.213-1.7l-.56.28l-2.576 8.923c1.216.803 2.907 1.71 4.68 1.915c2.688.312 5.56-.745 6.393-1.197l2.615-9.058l-.56-.28z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Woman(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.022 16.35c-.61-1.105-1.36-2-2.11-2.624a5.231 5.231 0 0 1-3.103 1.03a5.236 5.236 0 0 1-3.106-1.03c-.75.625-1.498 1.52-2.11 2.623c-1.423 2.562-1.58 5.19-.352 5.873c.55.307 1.127.078 1.723-.496a10.583 10.583 0 0 0-.167 1.873c0 2.932 1.14 5.307 2.543 5.307c.847 0 1.266-.865 1.467-2.19c.2 1.325.62 2.19 1.464 2.19c1.407 0 2.546-2.375 2.546-5.307c0-.66-.06-1.29-.168-1.873c.598.574 1.174.803 1.725.496c1.23-.682 1.07-3.31-.35-5.874zm-5.214-2.593a4.279 4.279 0 1 0-.001-8.557a4.28 4.28 0 1 0 0 8.558zm2.923-8.783c1.236.455.493-.725.493-1.53s.762-1.793-.492-1.392c-1.315.422-2.382.654-2.382 1.46s1.067.977 2.383 1.462zM15.817 4.4c.782 0 .345-.396.345-.884s.44-.883-.344-.883s-.374.396-.374.883c0 .49-.408.884.374.884zm-2.932.574c1.316-.484 2.383-.654 2.383-1.46S14.2 2.473 12.884 2.05c-1.254-.402-.492.584-.492 1.39s-.744 1.986.492 1.532z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wrench(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M26.834 14.693a7.062 7.062 0 0 0 1.193-7.334l-3.646 4.25l-3.593-.698l-1.19-3.462l3.636-4.242a7.063 7.063 0 0 0-8.106 9.903L5.625 24.04a2.79 2.79 0 0 0 4.209 3.661l9.493-10.917a7.072 7.072 0 0 0 7.508-2.09z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WrenchThree(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m27.84 6.775l-3.198 3.24l-2.872-.77L21 6.373l3.187-3.23a5.727 5.727 0 0 0-5.848 1.39a5.738 5.738 0 0 0-1.28 6.172l-9.64 9.64a3.971 3.971 0 0 0-.62-.062a3.932 3.932 0 0 0-3.934 3.933a3.932 3.932 0 1 0 7.865 0c0-.24-.03-.473-.07-.7l9.59-9.59a5.75 5.75 0 0 0 6.203-1.272a5.733 5.733 0 0 0 1.384-5.878zM6.8 25.145a.934.934 0 0 1-.936-.932a.932.932 0 1 1 .935.933z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WrenchTwo(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m24.946 9.72l-2.872-.767l-.77-2.874l3.187-3.232a5.73 5.73 0 0 0-5.847 1.39a5.747 5.747 0 0 0-1.28 6.173l-3.475 3.48l-3.478 3.477a5.745 5.745 0 0 0-6.173 1.277a5.731 5.731 0 0 0-1.39 5.85l3.23-3.19l2.875.77l.77 2.873l-3.24 3.197a5.743 5.743 0 0 0 7.146-7.586l3.464-3.464l3.464-3.463c2.07.83 4.523.407 6.202-1.27a5.732 5.732 0 0 0 1.384-5.877z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Zoomin(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.646 19.307a10.43 10.43 0 0 0 1.524-5.42c0-5.794-4.692-10.486-10.482-10.488c-5.79 0-10.484 4.693-10.484 10.485c0 5.79 4.693 10.48 10.484 10.48c1.987 0 3.84-.562 5.422-1.522l7.128 7.127l3.535-3.537l-7.127-7.126zm-8.958 1.062a6.495 6.495 0 0 1-6.484-6.485a6.494 6.494 0 0 1 6.484-6.486a6.493 6.493 0 0 1 6.484 6.485a6.496 6.496 0 0 1-6.484 6.484zm2-11.32h-4v2.834H8.853v4h2.833v2.834h4v-2.834h2.832v-4h-2.834V9.052z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Zoomout(children ...ElementRenderer) *RaphaelIcon {
	return &RaphaelIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.646 19.307a10.43 10.43 0 0 0 1.524-5.42c0-5.794-4.692-10.486-10.482-10.488c-5.79 0-10.484 4.693-10.484 10.485c0 5.79 4.693 10.48 10.484 10.48c1.987 0 3.84-.562 5.422-1.522l7.128 7.127l3.535-3.537l-7.127-7.126zm-8.958 1.062a6.495 6.495 0 0 1-6.484-6.485a6.494 6.494 0 0 1 6.484-6.486a6.493 6.493 0 0 1 6.484 6.485a6.496 6.496 0 0 1-6.484 6.484zm-4.834-8.486v4h9.665v-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
