package wpf

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = ""
	hAttr          = "1em"
	viewbox        = "0 0 26 26"
	fill           = "currentColor"
)

type WpfIcon struct {
	*SVGSVGElement
}

type WpfIconFn func(children ...ElementRenderer) *WpfIcon

var IconLookup = map[string]WpfIconFn{
	"addImage":              AddImage,
	"addUser":               AddUser,
	"administrator":         Administrator,
	"airplaneTakeoff":       AirplaneTakeoff,
	"alarmClock":            AlarmClock,
	"alignCenter":           AlignCenter,
	"alignJustify":          AlignJustify,
	"alignLeft":             AlignLeft,
	"alignRight":            AlignRight,
	"android":               Android,
	"androidOs":             AndroidOs,
	"appShield":             AppShield,
	"approval":              Approval,
	"archiveTwo":            ArchiveTwo,
	"askQuestion":           AskQuestion,
	"assistant":             Assistant,
	"attach":                Attach,
	"audioWave":             AudioWave,
	"autopilot":             Autopilot,
	"ballPointPen":          BallPointPen,
	"bankCards":             BankCards,
	"banknotes":             Banknotes,
	"bed":                   Bed,
	"behance":               Behance,
	"bicycle":               Bicycle,
	"birthday":              Birthday,
	"blackberry":            Blackberry,
	"bold":                  Bold,
	"books":                 Books,
	"briefcase":             Briefcase,
	"brightMoon":            BrightMoon,
	"bug":                   Bug,
	"building":              Building,
	"businessContact":       BusinessContact,
	"calendar":              Calendar,
	"callTransfer":          CallTransfer,
	"camera":                Camera,
	"carRental":             CarRental,
	"chargeBattery":         ChargeBattery,
	"chat":                  Chat,
	"checkBook":             CheckBook,
	"checkFile":             CheckFile,
	"checkmark":             Checkmark,
	"clapperboard":          Clapperboard,
	"clipboard":             Clipboard,
	"cloakroom":             Cloakroom,
	"closeUpMode":           CloseUpMode,
	"cloud":                 Cloud,
	"coins":                 Coins,
	"collaborator":          Collaborator,
	"colorDropper":          ColorDropper,
	"connected":             Connected,
	"controller":            Controller,
	"cornet":                Cornet,
	"createNew":             CreateNew,
	"cursor":                Cursor,
	"cut":                   Cut,
	"cutPaper":              CutPaper,
	"defineLocation":        DefineLocation,
	"delete":                Delete,
	"deleteShield":          DeleteShield,
	"deskLamp":              DeskLamp,
	"details":               Details,
	"diningRoom":            DiningRoom,
	"diplomaOne":            DiplomaOne,
	"disconnected":          Disconnected,
	"doctorsBag":            DoctorsBag,
	"edit":                  Edit,
	"editFile":              EditFile,
	"editImage":             EditImage,
	"eject":                 Eject,
	"emptyBattery":          EmptyBattery,
	"emptyFlag":             EmptyFlag,
	"end":                   End,
	"facialRecognitionScan": FacialRecognitionScan,
	"factory":               Factory,
	"fan":                   Fan,
	"faq":                   Faq,
	"fastForward":           FastForward,
	"filledFlag":            FilledFlag,
	"fingerprintScan":       FingerprintScan,
	"first":                 First,
	"fullBattery":           FullBattery,
	"fullTrash":             FullTrash,
	"future":                Future,
	"geoFence":              GeoFence,
	"ghost":                 Ghost,
	"globeEarth":            GlobeEarth,
	"gpsReceiving":          GpsReceiving,
	"groupIcon":             GroupIcon,
	"guitar":                Guitar,
	"happy":                 Happy,
	"headset":               Headset,
	"helicopter":            Helicopter,
	"heraldTrumpet":         HeraldTrumpet,
	"hexagon":               Hexagon,
	"highBattery":           HighBattery,
	"imageFile":             ImageFile,
	"inTransit":             InTransit,
	"inbox":                 Inbox,
	"inspection":            Inspection,
	"integratedCircuit":     IntegratedCircuit,
	"invisible":             Invisible,
	"iphone":                Iphone,
	"keepDry":               KeepDry,
	"keySecurity":           KeySecurity,
	"keyboard":              Keyboard,
	"last":                  Last,
	"like":                  Like,
	"lol":                   Lol,
	"luggageTrolley":        LuggageTrolley,
	"macOs":                 MacOs,
	"maestro":               Maestro,
	"maintenance":           Maintenance,
	"medicalDoctor":         MedicalDoctor,
	"message":               Message,
	"messageOutline":        MessageOutline,
	"microphone":            Microphone,
	"mouse":                 Mouse,
	"musicalNotes":          MusicalNotes,
	"mute":                  Mute,
	"myTopic":               MyTopic,
	"name":                  Name,
	"news":                  News,
	"next":                  Next,
	"nfcCheckpoint":         NfcCheckpoint,
	"note":                  Note,
	"online":                Online,
	"packaging":             Packaging,
	"paid":                  Paid,
	"panorama":              Panorama,
	"paperClamp":            PaperClamp,
	"paperPlane":            PaperPlane,
	"partlyCloudyDay":       PartlyCloudyDay,
	"partlyCloudyNight":     PartlyCloudyNight,
	"passwordOne":           PasswordOne,
	"past":                  Past,
	"pause":                 Pause,
	"phone":                 Phone,
	"phoneOffice":           PhoneOffice,
	"play":                  Play,
	"polyline":              Polyline,
	"previous":              Previous,
	"privacy":               Privacy,
	"qrCode":                QrCode,
	"radio":                 Radio,
	"rain":                  Rain,
	"record":                Record,
	"recurringAppointment":  RecurringAppointment,
	"refreshShield":         RefreshShield,
	"removeImage":           RemoveImage,
	"rename":                Rename,
	"renewSubscription":     RenewSubscription,
	"repeat":                Repeat,
	"restrictionShield":     RestrictionShield,
	"retroTv":               RetroTv,
	"rewind":                Rewind,
	"rfidTag":               RfidTag,
	"ruler":                 Ruler,
	"search":                Search,
	"securityChecked":       SecurityChecked,
	"sent":                  Sent,
	"settings":              Settings,
	"shoppingBag":           ShoppingBag,
	"shoppingBasket":        ShoppingBasket,
	"shoppingCart":          ShoppingCart,
	"shuffle":               Shuffle,
	"shutdown":              Shutdown,
	"signature":             Signature,
	"simCardChip":           SimCardChip,
	"skipToStart":           SkipToStart,
	"snow":                  Snow,
	"speaker":               Speaker,
	"speechBubble":          SpeechBubble,
	"stack":                 Stack,
	"stackOfPhotos":         StackOfPhotos,
	"stanleyKnife":          StanleyKnife,
	"star":                  Star,
	"statistics":            Statistics,
	"stop":                  Stop,
	"sun":                   Sun,
	"survey":                Survey,
	"swissArmyKnife":        SwissArmyKnife,
	"switchCamera":          SwitchCamera,
	"timer":                 Timer,
	"today":                 Today,
	"todoList":              TodoList,
	"twoFswipeDown":         TwoFswipeDown,
	"twoFswipeRight":        TwoFswipeRight,
	"unlock":                Unlock,
	"unlockTwo":             UnlockTwo,
	"userShield":            UserShield,
	"videoCall":             VideoCall,
	"viewFile":              ViewFile,
	"voicemail":             Voicemail,
	"volumeDown":            VolumeDown,
	"volumeUp":              VolumeUp,
	"weddingCake":           WeddingCake,
	"windRose":              WindRose,
	"worldwideLocation":     WorldwideLocation,
}

func AddImage(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2a2 2 0 0 0-2 2c0 .738 1.156 2.057 1.156 3C1.156 7.943 0 10.921 0 13s1.156 4.991 1.156 6C1.156 20.009 0 21.262 0 22a2 2 0 0 0 2 2c.738 0 3.124-1.125 4.5-1.125S8.396 24 9.5 24s2.131-1.125 3.5-1.125c1.369 0 1.896 1.125 3 1.125s2.563-1.125 3.5-1.125c.937 0 3.762 1.125 4.5 1.125a2 2 0 0 0 2-2c0-.738-1.156-2.057-1.156-3c0-.943 1.156-4.05 1.156-6c0-1.95-1.156-4.991-1.156-6C24.844 5.991 26 4.738 26 4a2 2 0 0 0-2-2c-.738 0-3.124 1.156-4.5 1.156S17.604 2 16.5 2S14.369 3.156 13 3.156C11.631 3.156 11.104 2 10 2S7.438 3.156 6.5 3.156C5.562 3.156 2.738 2 2 2m2 3h18c.551 0 1 .449 1 1v14c0 .551-.449 1-1 1H4c-.551 0-1-.449-1-1V6c0-.551.449-1 1-1m8 2a1 1 0 0 0-1 1v3H8a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h3v3a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-3h3a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1h-3V8a1 1 0 0 0-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddUser(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5.156c-3.017 0-5.438 2.072-5.438 6.032c0 2.586 1.03 5.22 2.594 6.843c.61 1.623-.49 2.227-.718 2.313C3.781 16.502.093 18.602.093 20.688v.78c0 2.843 5.414 3.5 10.437 3.5a45.48 45.48 0 0 0 3.281-.124a7.75 7.75 0 0 1-2.124-5.344c0-1.791.61-3.432 1.624-4.75c-.15-.352-.21-.907.063-1.75c1.555-1.625 2.563-4.236 2.563-6.813c0-3.959-2.424-6.03-5.438-6.03zm9 13.031a6.312 6.312 0 1 0 0 12.625a6.312 6.312 0 0 0 0-12.625M18.625 16h1.75v2.594h2.594v1.812h-2.594V23h-1.75v-2.594H16v-1.812h2.625z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Administrator(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.563 15.9c-.159-.052-1.164-.505-.536-2.414h-.009c1.637-1.686 2.888-4.399 2.888-7.07c0-4.107-2.731-6.26-5.905-6.26c-3.176 0-5.892 2.152-5.892 6.26c0 2.682 1.244 5.406 2.891 7.088c.642 1.684-.506 2.309-.746 2.396c-3.324 1.203-7.224 3.394-7.224 5.557v.811c0 2.947 5.714 3.617 11.002 3.617c5.296 0 10.938-.67 10.938-3.617v-.811c0-2.228-3.919-4.402-7.407-5.557m-5.516 8.709c0-2.549 1.623-5.99 1.623-5.99l-1.123-.881c0-.842 1.453-1.723 1.453-1.723s1.449.895 1.449 1.723l-1.119.881s1.623 3.428 1.623 6.018c0 .406-3.906.312-3.906-.028"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirplaneTakeoff(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.25 2.031a6.52 6.52 0 0 0-2.063.563L4.908 9.406c-.453.203-.766-.03-1.188-.344c-.735-.54-1.823-1.304-2.156-1.562c-.526-.403-1.83-.154-1.47.656L2 12.437c.626 1.4 3.88 1.066 9.563-1.468c.483-.216 1.112-.507 1.656-.75l-1.438 7.656c-.185.743.005.952.25.844l.844-.375c.436-.195 1.04-.873 1.344-1.532l4.125-8.906c4.11-1.92 7.871-3.846 7.593-4.468C25.61 2.7 24.186 1.85 22.25 2.03zM7.031 3.062c-.37-.006-.688.029-.906.126l-.844.374c-.244.11-.208.39.469.75l3.406 2.063l3.781-1.688l-4.78-1.468a4.241 4.241 0 0 0-1.126-.156zM.813 24A1.001 1.001 0 0 0 1 26h24a1 1 0 1 0 0-2H1a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.094 0zM902 1469v2h26v-2zm4 5v2h18v-2zm-4 5v2h26v-2zm4 5v2h18v-2zm-4 5v2h26v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlarmClock(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 .188c-1.104 0-2 .844-2 1.812h4c0-.908-.896-1.813-2-1.813zM4.969.5C3.916.582 2.867 1.07 2 1.938C.268 3.67-.077 6.173 1.656 7.905l6.282-6.281C7.072.758 6.021.418 4.967.5zM20.25.5c-.781.105-1.538.475-2.188 1.125l6.282 6.281c1.733-1.732 1.388-4.235-.344-5.968C22.917.853 21.552.325 20.25.5M13 3.188C7.028 3.188 2.187 8.027 2.187 14c0 3.581 1.756 6.751 4.438 8.719l-.781 2.312a.684.684 0 0 0 .406.844a.695.695 0 0 0 .906-.281l1.032-1.907A10.79 10.79 0 0 0 13 24.813c1.727 0 3.364-.405 4.813-1.125l1.03 1.907c.1.199.451.48.907.281a.684.684 0 0 0 .406-.844l-.781-2.312c2.682-1.968 4.438-5.138 4.438-8.719c0-5.972-4.841-10.813-10.813-10.813m-1 2.624V6.5a1 1 0 0 0 .469.844a.754.754 0 0 0-.125.343l-.25 5.032A1.572 1.572 0 0 0 11.438 14c0 .867.695 1.563 1.562 1.563c.055 0 .103-.026.156-.032l3.063 2.719c.172.154.604.27.968-.094c.365-.365.244-.794.094-.968l-2.718-3.125V14c0-.553-.287-1.033-.72-1.313l-.187-5a.75.75 0 0 0-.125-.343A1 1 0 0 0 14 6.5v-.688c3.706.452 6.682 3.374 7.188 7.063H20.5a1.123 1.123 0 0 0-.125 0a1.127 1.127 0 1 0 .125 2.25h.688c-.5 3.648-3.415 6.562-7.063 7.063V21.5a1.123 1.123 0 0 0-1.375-1.125a1.123 1.123 0 0 0-.875 1.125v.688c-3.648-.5-6.562-3.415-7.063-7.063H5.5a1.125 1.125 0 0 0 0-2.25h-.688C5.319 9.185 8.294 6.264 12 5.812"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenter(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 2v2h26V2zm4 5v2h18V7zm-4 5v2h26v-2zm4 5v2h18v-2zm-4 5v2h26v-2zm902 1447v2h26v-2zm4 5v2h18v-2zm-4 5v2h26v-2zm4 5v2h18v-2zm-4 5v2h26v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignJustify(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 2v2h26V2zm0 5v2h26V7zm0 5v2h26v-2zm0 5v2h26v-2zm0 5v2h26v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeft(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 2v2h26V2zm0 5v2h18V7zm0 5v2h26v-2zm0 5v2h18v-2zm0 5v2h26v-2zm902 1447v2h26v-2zm4 5v2h18v-2zm-4 5v2h26v-2zm4 5v2h18v-2zm-4 5v2h26v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRight(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 2v2h26V2zm8 5v2h18V7zm-8 5v2h26v-2zm8 5v2h18v-2zm-8 5v2h26v-2zm902 1447v2h26v-2zm4 5v2h18v-2zm-4 5v2h26v-2zm4 5v2h18v-2zm-4 5v2h26v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Android(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a3 3 0 0 0-3 3v20a3 3 0 0 0 3 3h11a3 3 0 0 0 3-3V3a3 3 0 0 0-3-3zm3.5 1h4a.5.5 0 0 1 0 1h-4a.5.5 0 0 1 0-1m-4 2h12a.5.5 0 0 1 .5.5v17a.5.5 0 0 1-.5.5h-12a.5.5 0 0 1-.5-.5v-17a.5.5 0 0 1 .5-.5m4.156 3.719c-.04.029-.062.111-.031.156l.531.75c-.837.388-1.426 1.137-1.5 2h5.688c-.074-.863-.663-1.611-1.5-2l.531-.75c.031-.045.041-.128 0-.156c-.041-.028-.124-.014-.156.031l-.531.813a3.256 3.256 0 0 0-1.188-.22a3.23 3.23 0 0 0-1.188.22l-.53-.813c-.033-.045-.085-.058-.126-.031m.656 1.468c.174 0 .313.139.313.313s-.14.344-.313.344a.366.366 0 0 1-.343-.344c0-.173.17-.313.344-.313zm2.438 0a.31.31 0 0 1 .313.313c0 .174-.139.344-.313.344c-.173 0-.313-.169-.313-.344c0-.173.139-.313.313-.313M9.625 10c-.353 0-.625.298-.625.656v2.531c0 .359.272.657.625.657a.664.664 0 0 0 .656-.656v-2.532A.665.665 0 0 0 9.625 10m7.719 0c-.354 0-.625.298-.625.656v2.531c0 .359.27.657.625.657a.665.665 0 0 0 .656-.656v-2.532a.665.665 0 0 0-.656-.656m-6.688.031v4.625a.5.5 0 0 0 .5.5h.563v1.438c0 .358.302.625.656.625a.612.612 0 0 0 .625-.625v-1.438h1v1.438c0 .358.303.625.656.625a.612.612 0 0 0 .625-.625v-1.438h.594a.5.5 0 0 0 .5-.5v-4.625zM8.5 23h2a.5.5 0 0 1 0 1h-2a.5.5 0 0 1 0-1m4 0h2a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 .5-.5m4 0h2a.5.5 0 0 1 0 1h-2a.5.5 0 0 1 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AndroidOs(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M24.123 9.753v6.281c0 .886-.71 1.608-1.583 1.608s-1.579-.722-1.579-1.608V9.753c0-.886.705-1.605 1.579-1.605s1.583.719 1.583 1.605M3.456 8.147c-.872 0-1.579.72-1.579 1.605v6.281c0 .886.707 1.608 1.579 1.608c.878 0 1.583-.722 1.583-1.608v-6.28c0-.886-.705-1.606-1.583-1.606m2.537 11.519c0 .683.553 1.233 1.234 1.233h1.41v3.495c0 .886.71 1.605 1.583 1.605s1.581-.719 1.581-1.605V20.9h2.46v3.495c0 .886.712 1.605 1.583 1.605c.874 0 1.583-.719 1.583-1.605V20.9h1.407c.683 0 1.236-.55 1.236-1.233V8.196H5.993zM20.036 7.175H5.961c.184-2.134 1.634-3.961 3.703-4.922L8.363.363c-.077-.111-.06-.258.041-.328c.101-.068.246-.034.326.077l1.349 1.962c.891-.35 1.878-.548 2.921-.548s2.028.198 2.921.548L17.27.112c.08-.111.225-.145.326-.077c.101.07.118.217.041.328l-1.301 1.89c2.068.961 3.519 2.788 3.7 4.922m-9.225-2.728a.78.78 0 1 0-1.562.004a.78.78 0 0 0 1.562-.004m6.034 0a.777.777 0 0 0-.78-.777a.778.778 0 1 0 .78.777" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppShield(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.813.094c-.59.103-1.379.5-2.563 1.093c-1.921.963-4.798 3.41-7.094 3.47a1.098 1.098 0 0 0-.781.374a1.208 1.208 0 0 0-.313.844c.495 10.023 4.099 16.207 10.407 19.813a1.1 1.1 0 0 0 .531.156c.182 0 .367-.064.531-.157c6.308-3.605 9.912-9.789 10.406-19.812a1.208 1.208 0 0 0-.312-.844a1.092 1.092 0 0 0-.781-.375c-2.295-.06-5.174-2.505-7.094-3.468C14.564.595 13.777.197 13.187.093c-.06-.008-.127 0-.187 0s-.129-.01-.188 0zM8.5 7h9A1.5 1.5 0 0 1 19 8.5v7a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 7 15.5v-7A1.5 1.5 0 0 1 8.5 7M16 8v1h1V8zm-8 2v5.5c0 .275.225.5.5.5h9c.275 0 .5-.225.5-.5V10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Approval(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m25.76 13.861l-1.519-2.396l.901-2.689a.813.813 0 0 0-.365-.962l-2.458-1.417l-.452-2.8a.816.816 0 0 0-.771-.683l-2.835-.111L16.56.533a.811.811 0 0 0-.999-.246L13 1.505L10.438.286a.815.815 0 0 0-.999.247l-1.701 2.27l-2.835.111a.813.813 0 0 0-.77.682l-.452 2.8l-2.458 1.417a.813.813 0 0 0-.366.962l.901 2.689L.24 13.861a.814.814 0 0 0 .125 1.022l2.047 1.963l-.23 2.826a.814.814 0 0 0 .584.848l2.725.785L6.6 23.916a.81.81 0 0 0 .911.479l2.78-.57l2.194 1.797c.149.121.332.184.515.184s.365-.063.515-.184l2.194-1.797l2.78.57c.377.08.76-.123.911-.479l1.109-2.611l2.725-.785a.813.813 0 0 0 .584-.848l-.23-2.826l2.047-1.963a.814.814 0 0 0 .125-1.022M18.877 9.95l-5.691 8.526c-.215.318-.548.531-.879.531c-.33 0-.699-.185-.934-.421l-4.178-4.245a.752.752 0 0 1 0-1.05l1.031-1.05a.727.727 0 0 1 1.031 0l2.719 2.762l4.484-6.718a.724.724 0 0 1 1.014-.196l1.209.831c.333.23.419.693.194 1.03"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArchiveTwo(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 0C1.344 0 0 1.344 0 3v20c0 1.656 1.344 3 3 3h20c1.656 0 3-1.344 3-3V6c0-1.656-1.344-3-3-3h-9.813a1 1 0 0 0-.406 0H11c0-1.656-1.344-3-3-3zm0 5h9v1a1 1 0 1 0 2 0V5h9a1 1 0 0 1 1 1v17a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1m9.906 2.969a1 1 0 0 0-.125.031A1 1 0 0 0 12 9v2a1 1 0 1 0 2 0V9a1 1 0 0 0-1.094-1.031M12 13c-.551 0-1 .449-1 1v1c0 .34.178.632.438.813c-.677 1.528-1.282 3.418-1.282 4.28a2.844 2.844 0 0 0 5.688 0c0-.863-.605-2.752-1.281-4.28A.99.99 0 0 0 15 15v-1c0-.551-.449-1-1-1zm1 5.469c.889 0 1.625.735 1.625 1.625S13.89 21.687 13 21.687s-1.625-.704-1.625-1.593c0-.891.736-1.625 1.625-1.625"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AskQuestion(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 0C5.925 0 0 5.08 0 11.5c0 3.03 1.359 5.748 3.5 7.781a6.733 6.733 0 0 1-1.094 1.875A16.48 16.48 0 0 1 .375 23.22A1 1 0 0 0 1 25c2.215 0 3.808-.025 5.25-.406c1.29-.342 2.399-1.058 3.531-2.063c1.03.247 2.093.469 3.219.469c7.075 0 13-5.08 13-11.5S20.075 0 13 0m0 2c6.125 0 11 4.32 11 9.5S19.125 21 13 21c-1.089 0-2.22-.188-3.25-.469a1 1 0 0 0-.938.25c-1.125 1.079-1.954 1.582-3.062 1.875c-.51.135-1.494.103-2.188.157c.14-.158.271-.242.407-.407c.786-.96 1.503-1.975 1.719-3.125a1 1 0 0 0-.344-.937C3.249 16.614 2 14.189 2 11.5C2 6.32 6.875 2 13 2m-1.906 3.906a1 1 0 0 0-.469.25l-1.5 1.407l1.344 1.468l1.187-1.125h2.406L15 8.97v1.469l-2.563 1.718A1 1 0 0 0 12 13v2h2v-1.438l2.563-1.718A1 1 0 0 0 17 11V8.594a1 1 0 0 0-.25-.656l-1.5-1.688a1 1 0 0 0-.75-.344h-3.188a1 1 0 0 0-.218 0M12 16v2h2v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Assistant(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 0C6.412 0 1.248 3.477 2.219 11.75c-.173.138-.34.286-.469.469c-.503.71-.75 1.633-.75 2.625c0 1.245.734 2.203 1.5 2.906c.56.515 1.133.874 1.781 1.063c.51 1.64 1.449 3 2.563 3.937a1 1 0 1 0 1.281-1.531c-1.006-.847-1.847-1.978-2.156-3.438A1 1 0 0 0 5 17c.123 0-.646-.25-1.156-.719c-.51-.468-.844-1.055-.844-1.437c0-.464.111-.852.25-1.156A1 1 0 0 0 5 13v-.844a1 1 0 0 0 0-.094c.04-.713.134-1.368.281-1.968C6.66 9.62 7.808 10.389 9 6.875c1.74 5.173 6.246 2.214 9.344 2.063c.364.885.59 1.925.656 3.124a1 1 0 0 0 0 .094V13a1 1 0 0 0 1 1v2.469c-.434.341-.915.531-1 .531a1 1 0 0 0-.969.781c-.204.966-.658 1.779-1.218 2.469a1.01 1.01 0 1 0 1.562 1.281c.6-.739.998-1.706 1.313-2.718A4.228 4.228 0 0 0 21.313 18H22c-.245 1.281-.696 2.248-1.156 2.938c-.65.975-1.21 1.316-1.25 1.343c-.029.015-.384.21-1.407.438c-.78.173-1.903.358-3.406.437c-.201-1.115-1.369-1.968-2.781-1.968c-1.554 0-2.813 1.035-2.813 2.312s1.26 2.313 2.813 2.313c.971 0 1.838-.416 2.344-1.032c1.85-.074 3.256-.286 4.219-.5c1.186-.263 1.78-.562 1.78-.562a.8.8 0 0 0 .095-.063s.892-.605 1.718-1.843c.668-1.002 1.33-2.455 1.563-4.375c.728-.545 1.187-1.436 1.187-2.438v-2c0-1.657-1.25-3-2.906-3h-.156C21.848 3.002 17.174 0 12 0m8.656 9.875c.03.042.065.08.094.125h-.094c-.008-.041.009-.084 0-.125M8 12a2 2 0 1 0 0 4a2 2 0 0 0 0-4m8 0a2 2 0 1 0 0 4a2 2 0 0 0 0-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Attach(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.719 2.063a3.96 3.96 0 0 0-1.157.218c-1.499.505-2.785 1.66-4.062 2.938l-8.25 8.25c-.733.733-1.298 1.627-1.469 2.687a3.694 3.694 0 0 0 1.063 3.188a3.691 3.691 0 0 0 3.25 1.031c1.058-.19 1.944-.757 2.625-1.438l9.062-9.062a1 1 0 1 0-1.406-1.406l-9.063 9.062c-.43.43-1.024.779-1.562.875c-.538.096-.996.035-1.5-.468c-.525-.525-.581-.966-.5-1.47c.081-.503.397-1.084.906-1.593l8.25-8.25c1.21-1.209 2.367-2.13 3.281-2.438c.915-.307 1.571-.241 2.625.813c.788.787 1.626 1.497 1.844 2.219c.11.36.11.72-.125 1.312c-.234.592-.745 1.402-1.718 2.375c-4.148 4.15-7.332 7.332-9.063 9.063c-1.537 1.537-2.989 2.563-4.281 2.843c-1.293.281-2.52-.018-4.125-1.625c-1.607-1.607-2.169-3.163-2-4.78c.168-1.618 1.153-3.373 2.969-5.188c2.196-2.196 6.78-6.406 6.78-6.406a1 1 0 1 0-1.343-1.47S6.158 7.5 3.875 9.782C1.852 11.804.578 13.978.344 16.22c-.234 2.24.674 4.455 2.594 6.375s3.992 2.61 5.937 2.187c1.945-.422 3.63-1.755 5.281-3.406c1.731-1.73 4.915-4.913 9.063-9.063c1.1-1.1 1.812-2.083 2.187-3.03c.375-.949.39-1.884.157-2.657c-.467-1.545-1.72-2.408-2.344-3.031c-.716-.716-1.508-1.168-2.313-1.375a4.315 4.315 0 0 0-1.187-.156z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AudioWave(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.813 2.031a.95.95 0 0 0-.75.969v19a.95.95 0 1 0 1.875 0V3a.95.95 0 0 0-1.032-.969a.95.95 0 0 0-.093 0m-12 1a.95.95 0 0 0-.75.969v17a.95.95 0 1 0 1.875 0V4a.95.95 0 0 0-1.032-.969a.95.95 0 0 0-.093 0m9 3a.95.95 0 0 0-.75.969v11a.95.95 0 1 0 1.874 0V7a.95.95 0 0 0-1.03-.969a.95.95 0 0 0-.095 0zm-12 1a.95.95 0 0 0-.75.969v9a.95.95 0 1 0 1.874 0V8a.95.95 0 0 0-1.03-.969a.95.95 0 0 0-.095 0zm6 1a.95.95 0 0 0-.75.969v7a.95.95 0 1 0 1.874 0V9a.95.95 0 0 0-1.03-.969a.95.95 0 0 0-.095 0zm12 0a.95.95 0 0 0-.75.969v7a.95.95 0 1 0 1.875 0V9a.95.95 0 0 0-1.032-.969a.95.95 0 0 0-.093 0m-21 2a.95.95 0 0 0-.75.969v3a.95.95 0 1 0 1.875 0v-3a.95.95 0 0 0-1.032-.969a.95.95 0 0 0-.094 0zm12 0a.95.95 0 0 0-.75.969v3a.95.95 0 1 0 1.874 0v-3a.95.95 0 0 0-1.03-.969a.95.95 0 0 0-.095 0zm12 0a.95.95 0 0 0-.75.969v3a.95.95 0 1 0 1.875 0v-3a.95.95 0 0 0-1.032-.969a.95.95 0 0 0-.093 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Autopilot(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.6 5c-.2-.2-.5-.4-.8-.4c-2.3-.1-5.2-2.5-7.1-3.5c-1.2-.6-2-1-2.6-1.1h-.4c-.6.1-1.4.5-2.6 1.1c-1.9 1-4.8 3.4-7.1 3.5c-.2.1-.4.2-.6.4c-.2.3-.3.6-.3.9c.5 10 4.1 16.2 10.4 19.8c.2.1.3.1.5.1s.4 0 .5-.1c6.3-3.6 9.9-9.8 10.4-19.8c0-.3-.1-.6-.3-.9M19 15.5l-5-2.7V16l2 1.5v.5l-3-.5l-3 .5v-.5l2-1.5v-3.2l-5 2.7v-1.1l5-4.3V6c0-.6.4-1 1-1s1 .4 1 1v4.1l5 4.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BallPointPen(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 0s-.78.04-1.781.469c-1.002.43-2.36 1.234-3.938 2.812c-3.157 3.157-4.219 6.406-4.219 6.406a1 1 0 1 0 1.875.626s.939-2.751 3.782-5.594c1.421-1.422 2.564-2.117 3.312-2.438c.484-.207.496-.17.625-.187l1.157 1.156a61.263 61.263 0 0 0-5.782 5.063c-4.243 4.242-8.187 9.593-8.187 9.593l4.25 4.25s5.351-3.945 9.594-8.187c4.242-4.242 6.968-8.375 6.968-8.375l-4.25-4.25s-.376.287-.937.687l-1.75-1.75A1 1 0 0 0 17 0m6.656 0c-.637.07-1.387.386-2.343.844l3.843 3.844c.915-1.914 1.245-2.975.188-4.032c-.529-.528-1.05-.726-1.688-.656m-20.5 18.594c-1.325 2.21-1.715 4.043-1.812 5.093L.5 24.5a.707.707 0 0 0 1 1l.844-.844c1.054-.101 2.871-.498 5.062-1.812z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BankCards(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.688 0c-.2.008-.393.044-.594.094L2.5 3.406C.892 3.8-.114 5.422.281 7.031l1.906 7.782A2.99 2.99 0 0 0 4 16.875V15c0-2.757 2.243-5 5-5h12.594l-1.875-7.719A3.004 3.004 0 0 0 16.687 0zm1.218 4.313l.813 3.406l-3.375.812l-.844-3.375zM9 12c-1.656 0-3 1.344-3 3v8c0 1.656 1.344 3 3 3h14c1.656 0 3-1.344 3-3v-8c0-1.656-1.344-3-3-3zm0 1.594h14c.771 0 1.406.635 1.406 1.406v1H7.594v-1c0-.771.635-1.406 1.406-1.406M7.594 19h16.812v4c0 .771-.635 1.406-1.406 1.406H9A1.414 1.414 0 0 1 7.594 23z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Banknotes(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.531 1.531L14.25 4.406L5.5 9.188L.219 12.094L2.5 16.219V13A1.5 1.5 0 0 1 4 11.5h.813l.03-.125l11.595-6.344l.875.25c.024.08.05.173.093.25c.344.63 1.12.876 1.75.532c.076-.042.125-.133.188-.188l.906.281l2 3.657l-.25.906c-.079.024-.173.02-.25.062c-.293.16-.49.428-.594.719H25l-2.594-4.719zm-1.093 6.563a1.305 1.305 0 0 0-1.032 1.281c0 .716.565 1.313 1.282 1.313c.716 0 1.312-.597 1.312-1.313s-.596-1.281-1.313-1.281c-.089 0-.165-.018-.25 0zm-4.875.875a4.256 4.256 0 0 0-2.188.531c-.871.476-1.523 1.204-1.875 2h7.25a2.94 2.94 0 0 0-.313-.875c-.566-1.034-1.664-1.62-2.875-1.656zM4 13v13h22V13zm4.406 1.594h13.188l.656.656c-.017.083-.063.162-.063.25c0 .717.596 1.313 1.313 1.313c.088 0 .167-.047.25-.063l.656.656v4.188l-.656.656c-.083-.017-.162-.063-.25-.063c-.717 0-1.313.596-1.313 1.313c0 .088.046.167.063.25l-.656.656H8.406l-.656-.656c.017-.083.063-.162.063-.25c0-.717-.596-1.313-1.313-1.313c-.088 0-.167.047-.25.063l-.656-.656v-4.188l.656-.656c.083.017.162.063.25.063c.717 0 1.313-.596 1.313-1.313c0-.088-.046-.167-.063-.25zm6.594 1.5c-2.168 0-3.938 1.52-3.938 3.406c0 1.886 1.77 3.406 3.938 3.406s3.938-1.52 3.938-3.406c0-1.886-1.77-3.406-3.938-3.406m-6.5 2.093c-.716 0-1.313.597-1.313 1.313s.597 1.313 1.313 1.313s1.313-.597 1.313-1.313s-.597-1.313-1.313-1.313m13 0c-.717 0-1.313.596-1.313 1.313s.596 1.313 1.313 1.313s1.313-.596 1.313-1.313s-.596-1.313-1.313-1.313"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bed(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25 6h-1s-1 .014-1 1H3c0-.989-1-1-1-1H1c-.551 0-1 .449-1 1v18c0 .551.449 1 1 1h1s1-.014 1-1v-1c0-.551.449-1 1-1h18c.551 0 1 .449 1 1v1c0 .989 1 1 1 1h1c.551 0 1-.449 1-1V7c0-.551-.449-1-1-1M3 12c0-1.656.344-3 2-3h5c1.656 0 2 1.344 2 3zm11 0c0-1.656.344-3 2-3h5c1.656 0 2 1.344 2 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Behance(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.5 12.4c.2.3.4.6.5 1.1h-3.2c0-.1 0-.3.1-.5s.1-.3.3-.5c.1-.2.3-.3.5-.4c.2-.1.5-.2.8-.2c.4.1.8.2 1 .5m-9.2-.7c.2-.2.4-.5.4-.9c0-.2 0-.4-.1-.6c-.1-.1-.2-.3-.3-.3c-.1-.1-.3-.1-.5-.2H7.2V12h2.2c.3 0 .6-.1.9-.3m-.8 1.7H7.2v2.7h2.3c.2 0 .4 0 .6-.1c.2 0 .4-.1.5-.2c.1-.1.3-.2.4-.4c.1-.2.1-.4.1-.6c0-.5-.1-.9-.4-1.1c-.3-.2-.7-.3-1.2-.3M26 4.9v16.2c0 2.7-2.2 4.9-4.9 4.9H4.9C2.2 26 0 23.8 0 21.1V4.9C0 2.2 2.2 0 4.9 0h16.2C23.8 0 26 2.2 26 4.9m-9.6 4.7h4v-1h-4zM13.2 15c0-.6-.1-1.1-.4-1.6c-.3-.4-.7-.7-1.3-.9c.4-.2.8-.5 1-.8c.2-.3.3-.7.3-1.2s-.1-.8-.2-1.2c-.2-.3-.4-.6-.6-.7c-.3-.2-.6-.3-1-.4c-.5-.2-.9-.2-1.4-.2H5v9.8h4.8c.4 0 .9-.1 1.3-.2s.8-.3 1.1-.5c.3-.2.6-.5.8-.9c.1-.3.2-.7.2-1.2m3.6-.3h5.1c0-.6 0-1.1-.1-1.6s-.3-1-.6-1.3c-.3-.4-.7-.7-1.1-.9c-.5-.2-1-.4-1.6-.4c-.5 0-1 .1-1.5.3c-.4.2-.8.5-1.2.8c-.3.3-.6.7-.7 1.2c-.2.5-.3 1-.3 1.5s.1 1.1.3 1.5c.2.5.4.9.7 1.2c.3.3.7.6 1.1.8c.5.2 1 .3 1.5.3c.8 0 1.5-.2 2.1-.6c.6-.4 1-1 1.3-1.8h-1.7c-.1.2-.2.4-.5.6c-.3.2-.6.3-1 .3c-.5 0-1-.1-1.3-.4c-.3-.4-.5-.9-.5-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bicycle(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.5 6.719a.75.75 0 0 0-.156.031a.75.75 0 0 0-.469.313c-.707 1.018-2.781.906-2.781.906a.75.75 0 0 0-.813.906s.244 1.172.719 2.875H9.906l-.281-.688c.97-.25 1.313-.37 1.313-.718c0-.334-.445-.345-.907-.344H6.687c-.469 0-.625.191-.625.75c0 1.79.837 1.184 2.094.719l.406 1.031l-3.75 7.688A.75.75 0 0 0 5.5 21.25h6.656a.75.75 0 0 0 1.125-.531l3.813-5.563c.699 1.89 1.603 3.932 2.781 5.75a.75.75 0 1 0 1.25-.812c-1.545-2.386-2.556-5.094-3.219-7.219l-.218-.594a36.978 36.978 0 0 1-.72-2.844c.931-.013 2.232-.168 3.157-1.5A.75.75 0 0 0 19.5 6.72zm-9.031 6.531h5.969c.002.009-.003.022 0 .031l.03.094l-3.78 5.531zm-1.156 1.125l2.093 5.375H6.687zm-3.813.813A5.311 5.311 0 0 0 .187 20.5A5.311 5.311 0 0 0 5.5 25.813c2.336 0 4.288-1.531 5-3.625H8.594c-.605 1.102-1.752 1.875-3.094 1.875A3.57 3.57 0 0 1 1.937 20.5c0-1.903 1.525-3.44 3.407-3.531l.812-1.719c-.218-.027-.431-.063-.656-.063zm15 0c-.254 0-.505.027-.75.062c.224.559.486 1.127.75 1.688a3.57 3.57 0 0 1 3.563 3.562a3.57 3.57 0 0 1-3.563 3.563a3.57 3.57 0 0 1-3.563-3.563c0-.702.211-1.355.563-1.906c-.29-.602-.535-1.213-.781-1.813a5.326 5.326 0 0 0-1.532 3.719a5.311 5.311 0 0 0 5.313 5.313a5.311 5.311 0 0 0 5.313-5.313a5.311 5.311 0 0 0-5.313-5.313zm-11.281 1.53l-.844 1.72c.082.113.15.25.219.375H10.5c-.27-.791-.698-1.52-1.281-2.094zM902 1469v2h26v-2zm4 5v2h18v-2zm-4 5v2h26v-2zm4 5v2h18v-2zm-4 5v2h26v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Birthday(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 0S3.031 2.169 3.031 3.25a1.95 1.95 0 0 0 1.75 1.938a1.02 1.02 0 0 1-.812-1C3.969 3.625 5 2.5 5 2.5s1.031 1.124 1.031 1.688c0 .493-.348.905-.812 1a1.95 1.95 0 0 0 1.75-1.938C6.969 2.17 5 0 5 0m8 0s-1.969 2.169-1.969 3.25a1.95 1.95 0 0 0 1.75 1.938a1.02 1.02 0 0 1-.812-1C11.969 3.625 13 2.5 13 2.5s1.031 1.124 1.031 1.688c0 .493-.348.905-.812 1a1.95 1.95 0 0 0 1.75-1.938C14.969 2.17 13 0 13 0m8 0s-1.969 2.169-1.969 3.25a1.95 1.95 0 0 0 1.75 1.938a1.02 1.02 0 0 1-.812-1C19.969 3.625 21 2.5 21 2.5s1.031 1.124 1.031 1.688c0 .493-.348.905-.812 1a1.95 1.95 0 0 0 1.75-1.938C22.969 2.17 21 0 21 0M3.031 5.688a.558.558 0 0 0-.656.562c0 .36.656 1.011.656 2.156v5.563H6.97V8.313c0-.699.656-.952.656-1.313c0-.36-.015-.75-1.125-.75c-1.651 0-3.469-.563-3.469-.563zm8 0a.558.558 0 0 0-.656.562c0 .36.656 1.011.656 2.156v5.563h3.938V8.313c0-.699.656-.952.656-1.313c0-.36-.015-.75-1.125-.75c-1.651 0-3.469-.563-3.469-.563zm8 0a.558.558 0 0 0-.656.562c0 .36.656 1.011.656 2.156v5.563h3.938V8.313c0-.699.656-.952.656-1.313c0-.36-.015-.75-1.125-.75c-1.651 0-3.469-.563-3.469-.563zM4 15c-1.656 0-3 1.344-3 3v6a1 1 0 0 0-.094 0A1.001 1.001 0 0 0 1 26h24a1 1 0 1 0 0-2v-6c0-1.656-1.344-3-3-3zm1 4a2 2 0 0 0 4 0a2 2 0 0 0 4 0a2 2 0 0 0 4 0a2 2 0 0 0 4 0a2 2 0 0 0 2 2v3H3v-3a2 2 0 0 0 2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blackberry(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 0C6.733 0 5 1.996 5 3.652V22c0 1.656 2.422 3.813 8.5 3.813S22 23.656 22 22V3.652C22 1.996 20.267 0 13.5 0M20 10.5a.5.5 0 0 1-.5.5h-12a.5.5 0 0 1-.5-.5v-7a.5.5 0 0 1 .5-.5h12a.5.5 0 0 1 .5.5zm-5 5a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1-.5-.5v-2a.5.5 0 0 1 .5-.5h2a.5.5 0 0 1 .5.5zm-5-1.327a.52.52 0 0 1-.5.5c-1.537 0-2.5-.896-2.5-1.173a.5.5 0 0 1 .5-.5h2a.5.5 0 0 1 .5.5zm7-.673a.5.5 0 0 1 .5-.5h2a.5.5 0 0 1 .5.5c0 .276-.963 1.173-2.5 1.173a.52.52 0 0 1-.5-.5zM7.5 18a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 1 0v1a.5.5 0 0 1-.5.5m2 .5A.5.5 0 0 1 9 18v-1a.5.5 0 0 1 1 0v1a.5.5 0 0 1-.5.5m2 .5a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 1 0v1a.5.5 0 0 1-.5.5m2 0a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 1 0v1a.5.5 0 0 1-.5.5m2 0a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 1 0v1a.5.5 0 0 1-.5.5m2-.5a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 1 0v1a.5.5 0 0 1-.5.5m2-.5a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 1 0v1a.5.5 0 0 1-.5.5m-12 3a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 1 0v1a.5.5 0 0 1-.5.5m2 .5A.5.5 0 0 1 9 21v-1a.5.5 0 0 1 1 0v1a.5.5 0 0 1-.5.5m2 .5a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 1 0v1a.5.5 0 0 1-.5.5m2 0a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 1 0v1a.5.5 0 0 1-.5.5m2 0a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 1 0v1a.5.5 0 0 1-.5.5m2-.5a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 1 0v1a.5.5 0 0 1-.5.5m2-.5a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 1 0v1a.5.5 0 0 1-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bold(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.941 18.046c0 1.116-.234 2.06-.706 2.828c-.47.772-1.144 1.378-2.022 1.825c-.932.473-2.003.811-3.217 1.008c-1.215.195-1.644.293-3.288.293H2v-1.784c.32-.028 2.031-.421 2.031-1.953V6.053C4.031 4.197 2.32 3.979 2 3.929V2h11.953c3.008 0 3.639.414 4.973 1.239c1.334.829 2 2.048 2 3.66c0 3.087-2.469 4.929-3.275 5.116v.294c.806.083 5.29.951 5.29 5.737M15.434 7.824c0-.911-.344-1.623-1.031-2.132c-.687-.515-1.488-.769-2.862-.769c-.196 0-.452.007-.766.018c-.316.012-.588.022-.815.029v6.102h.806c1.675 0 2.68-.293 3.476-.877c.793-.583 1.192-1.372 1.192-2.371m.814 9.651c0-1.165-.435-2.061-1.309-2.682c-.871-.623-1.903-.934-3.539-.934c-.188 0-.438.006-.75.018l-.691.029v7.174h2.307c1.18 0 2.137-.314 2.876-.945c.737-.63 1.106-1.518 1.106-2.66"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Books(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.875 0a1 1 0 0 0-.406.156S8.204.952 6.844 1.813c-1.36.86-2.873 1.808-3.219 2a1 1 0 0 0-.063.03C2.306 4.618 2.045 5.884 2 6.594c-.003.033 0 .06 0 .095c-.011.266 0 .437 0 .437v13.063C2 22.087 4.213 23 6.313 23c.7 0 1.4-.113 2-.313c.4-.2.687-.6.687-1v-10.5c0-2.3.5-3.38 2-4.28c.4-.2 4.594-3.095 4.594-3.095c.2-.2.406-.606.406-.906v-.094c0-.4-.2-.706-.5-.906c-.3-.2-.7-.2-1 0c-.1.1-6.2 4.207-7.5 4.907c-1.3.8-2.513.993-2.813.593c-.093-.093-.174-.378-.187-.656v-.063c.001-.272.071-.784.625-1.125c.562-.313 1.957-1.204 3.313-2.062c.573-.363.644-.402 1.093-.688A1 1 0 0 0 11 2.5V1a1 1 0 0 0-1.125-1m8 3.5a1 1 0 0 0-.438.188s-5.034 3.387-5.906 3.968a1 1 0 0 0-.031.032c-.724.543-1.153 1.189-1.344 1.78A3.264 3.264 0 0 0 10 10.5v.313a1 1 0 0 0 0 .093V23c0 1.9 2.188 3 4.188 3c.9 0 1.712-.194 2.312-.594c1.2-.7 7-5.218 7-5.218c.3-.2.5-.482.5-.782v-13c0-.5-.194-.8-.594-1c-.3-.2-.793-.106-1.093.094c-1.6 1.2-5.907 4.588-6.907 5.188c-1.4.8-2.719 1-3.219.5c-.2-.2-.187-.388-.187-.688c.006-.172.025-.32.063-.438c.056-.174.17-.388.593-.718c.02-.016.01-.015.031-.031c.723-.483 2.934-1.99 4.376-2.97A1 1 0 0 0 19 6V4.5a1 1 0 0 0-1.125-1M22 10.813v2l-5 3.874v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Briefcase(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 1C8.802 1 7 2.803 7 5v1h2V5c0-1.115.884-2 2-2h4c1.116 0 2 .885 2 2v1h2V5c0-2.197-1.802-4-4-4zM4.875 7A4.874 4.874 0 0 0 0 11.875v9.25A4.874 4.874 0 0 0 4.875 26h16.25A4.874 4.874 0 0 0 26 21.125v-9.25A4.874 4.874 0 0 0 21.125 7zM13 13.906c1.161 0 2.094.933 2.094 2.094A2.088 2.088 0 0 1 13 18.094A2.088 2.088 0 0 1 10.906 16c0-1.161.933-2.094 2.094-2.094"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrightMoon(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.5.063l-.625 2.28l-2.25-.593L8.25 3.469L6.625 5.156l2.25-.593l.625 2.312l.625-2.313l2.25.594L10.75 3.47l1.625-1.719l-2.25.594zM2.469 6.5l-.5 1.813l-1.813-.47L1.47 9.189L.156 10.53l1.813-.469l.5 1.813L3 10.062l1.813.47L3.5 9.186l1.313-1.343L3 8.312zm15.625.688a8.273 8.273 0 0 1 4.937 7.593c0 4.581-3.685 8.281-8.25 8.281a8.29 8.29 0 0 1-7.594-4.968a9.39 9.39 0 0 0 9.25 7.718a9.361 9.361 0 0 0 9.375-9.375a9.396 9.396 0 0 0-7.718-9.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bug(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 .188a2.812 2.812 0 0 0-2.781 2.437C8.545 3.175 7.188 4.264 7.188 6c0 2.658 1.318 3.813 5.812 3.813S18.813 8.658 18.813 6c0-1.736-1.357-2.825-3.032-3.375A2.812 2.812 0 0 0 13 .187zM2.906 2.967A1 1 0 0 0 2.781 3A1 1 0 0 0 2 4c0 3.257 1.1 5.33 2.281 6.5c.318.314.629.55.938.75c-.02.48-.032 1.018-.032 1.75H1a1 1 0 0 0-.094 0A1.001 1.001 0 0 0 1 15h4.406c.097.49.213.973.375 1.438a8.29 8.29 0 0 0-1.25.53C2.81 17.869 1 19.787 1 23a1 1 0 1 0 2 0c0-2.515 1.19-3.582 2.469-4.25c.406-.212.79-.359 1.156-.469c1.234 2.124 3.159 3.747 5.375 4.375V15c0-3 2-3 2 0v7.656c2.216-.628 4.141-2.251 5.375-4.375c.366.11.75.257 1.156.469C21.81 19.418 23 20.485 23 23a1 1 0 1 0 2 0c0-3.214-1.81-5.132-3.531-6.031a8.328 8.328 0 0 0-1.25-.532c.162-.464.278-.948.375-1.437H25a1 1 0 1 0 0-2h-4.188c0-.713-.033-1.263-.062-1.719c.32-.203.638-.453.969-.781C22.9 9.33 24 7.257 24 4a1 1 0 0 0-1.219-1A1 1 0 0 0 22 4c0 2.819-.9 4.252-1.719 5.063c-.818.81-1.5.968-1.5.968a1 1 0 0 0-.5.188a1 1 0 0 0-.031.031c-1.163.434-2.822.938-5.25.938c-3.66 0-5.618-1.226-6.656-1.626c-.01-.004-.021.004-.032 0c-.01-.007-.02-.023-.03-.03a4.89 4.89 0 0 1-.563-.47C4.9 8.252 4 6.82 4 4a1 1 0 0 0-1.094-1.031z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Building(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 0C4.355 0 3 1.355 3 3v23h20V3c0-1.645-1.355-3-3-3zm0 2h14c.555 0 1 .445 1 1v21h-2v-5h-5v5H5V3c0-.555.445-1 1-1m1 2v3h5V4zm7 0v3h5V4zM7 9v3h5V9zm7 0v3h5V9zm-7 5v3h5v-3zm7 0v3h5v-3zm-7 5v3h5v-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BusinessContact(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 3C1.803 3 0 4.802 0 7v12c0 2.197 1.803 4 4 4h2a1 1 0 0 0 1-1v-1h1v1a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-1h1v1a1 1 0 0 0 1 1h2c2.197 0 4-1.803 4-4V7c0-2.198-1.803-4-4-4zm0 2h18c1.115 0 2 .884 2 2v12c0 1.115-.885 2-2 2h-1v-1a1 1 0 0 0-1-1h-3a1 1 0 0 0-1 1v1h-6v-1a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1v1H4c-1.115 0-2-.885-2-2V7c0-1.116.885-2 2-2m4 2.063c-1.14 0-1.969.807-1.969 2.28c0 .963.41 1.616 1 2.22c.23.603-.195.811-.281.843c-1.192.431-2.688 1.068-2.688 1.844v.406c0 1.058 2.041 1.281 3.938 1.281c1.9 0 3.938-.223 3.938-1.28v-.407c0-.8-1.5-1.43-2.75-1.844c-.059-.019-.414-.16-.188-.844c.588-.604.938-1.26.938-2.218c0-1.475-.799-2.281-1.938-2.281m10.5.03c-.608 0-1.183.088-1.719.313a4.306 4.306 0 0 0-2.343 2.344c-.227.535-.313 1.13-.313 1.75c0 .652.095 1.236.313 1.781a4.021 4.021 0 0 0 2.281 2.313a4.691 4.691 0 0 0 1.75.312c.37 0 .7.003 1-.031c.298-.033.616-.086.937-.156v-1.094c-.293.085-.62.168-.968.219c-.348.05-.673.062-.97.062c-.42 0-.834-.072-1.218-.219a2.85 2.85 0 0 1-1-.656a3.002 3.002 0 0 1-.719-1.062c-.172-.43-.25-.92-.25-1.469c0-.515.1-.981.281-1.406c.184-.425.42-.804.72-1.094c.313-.3.66-.505 1.03-.656c.372-.153.77-.25 1.188-.25c.44 0 .839.082 1.219.219c.38.136.718.34 1 .624c.279.28.501.618.656 1.032c.156.415.219.915.219 1.469c0 .313-.029.602-.094.906a3.94 3.94 0 0 1-.25.812h-.781V9.062H19.28v.313a1.668 1.668 0 0 0-.406-.281a1.295 1.295 0 0 0-.594-.125a1.85 1.85 0 0 0-.812.187a2.176 2.176 0 0 0-.688.5c-.2.233-.348.525-.468.844c-.12.32-.188.671-.188 1.063c0 .38.058.738.156 1.062c.097.325.208.604.375.813c.17.215.392.354.625.468c.233.114.481.188.75.188c.183 0 .333-.011.438-.031c.107-.022.22-.05.343-.094c.11-.036.19-.071.282-.125c.088-.054.189-.127.281-.188l.063.344h2.593a4 4 0 0 0 .532-1.156a5.24 5.24 0 0 0 .187-1.438c0-.653-.099-1.254-.313-1.781a3.965 3.965 0 0 0-.906-1.344a3.918 3.918 0 0 0-1.343-.875a4.494 4.494 0 0 0-1.688-.312zm-.063 2.876c.155 0 .315 0 .438.031c.122.031.23.086.375.156v2.563a1.765 1.765 0 0 1-.406.281a1 1 0 0 1-.469.125c-.35 0-.597-.148-.75-.406c-.153-.258-.25-.628-.25-1.156c0-.479.097-.867.281-1.157a.87.87 0 0 1 .782-.437z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 0c-.551 0-1 .449-1 1v3c0 .551.449 1 1 1c.551 0 1-.449 1-1V1c0-.551-.449-1-1-1m12 0c-.551 0-1 .449-1 1v3c0 .551.449 1 1 1c.551 0 1-.449 1-1V1c0-.551-.449-1-1-1M3 2C1.344 2 0 3.344 0 5v18c0 1.656 1.344 3 3 3h20c1.656 0 3-1.344 3-3V5c0-1.656-1.344-3-3-3h-2v2a2 2 0 0 1-4 0V2H9v2a2 2 0 0 1-4 0V2zM2 9h22v14c0 .551-.449 1-1 1H3c-.551 0-1-.449-1-1zm7 3v2.313h4.813l-3.782 7.656H13.5l3.469-8.438V12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CallTransfer(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.531 4a4.774 4.774 0 0 0-2.906 1.156C.897 6.618.006 9.265 0 13c-.011 3.733.869 6.372 2.594 7.844A4.738 4.738 0 0 0 6.063 22h.656c1.242 0 1.367-1.001 1.187-2.25L7.781 19c-.182-1.241-.564-2.24-1.812-2.25h-.75C4.675 16.747 2.99 17.513 3 13c.007-4.504 1.675-3.755 2.219-3.75h.75C7.212 9.252 7.628 8.242 7.812 7l.125-.75C8.12 5.012 8.024 4.006 6.782 4h-.687a4.707 4.707 0 0 0-.563 0zm18 0a4.774 4.774 0 0 0-2.906 1.156C18.897 6.618 18.006 9.265 18 13c-.012 3.733.869 6.372 2.594 7.844A4.736 4.736 0 0 0 24.062 22h.657c1.242 0 1.367-1.001 1.187-2.25l-.125-.75c-.182-1.241-.565-2.24-1.812-2.25h-.75c-.544-.003-2.229.763-2.219-3.75c.007-4.504 1.675-3.755 2.219-3.75h.75c1.242.002 1.657-1.008 1.843-2.25l.125-.75c.182-1.238.086-2.244-1.156-2.25h-.687a4.707 4.707 0 0 0-.563 0M12.47 9.75a.413.413 0 0 0-.219.031c-.136.072-.25.222-.25.375v1.469c-.909.152-2.445.239-2.594.344c-.22.155-.375.613-.375 1.031c0 .419.155.908.375 1.063c.15.105 1.684.16 2.594.312v1.438c0 .153.114.301.25.374c.136.073.278.086.406 0c1.656-1.113 3.129-2.864 3.188-2.937a.388.388 0 0 0 .094-.25a.477.477 0 0 0-.094-.281c-.058-.074-1.532-1.792-3.188-2.906a.332.332 0 0 0-.187-.063"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 3c-.709 0-1.18.288-1.406.688L7.813 6H6c0-.551-.449-1-1-1H3c-.551 0-1 .449-1 1v.188A2.985 2.985 0 0 0 0 9v10.125A4.874 4.874 0 0 0 4.875 24h16.25A4.874 4.874 0 0 0 26 19.125v-8.25A4.874 4.874 0 0 0 21.125 6h-.938l-1.28-2.313C18.68 3.29 18.18 3 17.5 3zM4 6.5a1.5 1.5 0 1 1-.001 3.001A1.5 1.5 0 0 1 4 6.5m10 .438a7.063 7.063 0 1 1-.001 14.126A7.063 7.063 0 0 1 14 6.937zm0 2.25a4.812 4.812 0 1 0 0 9.624a4.812 4.812 0 0 0 0-9.625z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarRental(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.5 0C7.6 0 6 1.6 6 3.5S7.6 7 9.5 7c1.36 0 2.55-.836 3.125-2H19l1-1V3l-1-1h-1l-.813.813L16.5 2h-.813l-.593.594L14.406 2h-1.781C12.049.836 10.861 0 9.5 0m-1 2.5c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1M9.094 9a4.06 4.06 0 0 0-3.906 3l-.72 3H4c-.6 0-1 .4-1 1v1c0 .6.4 1 1 1v7c0 .6.4 1 1 1h2c.6 0 1-.4 1-1v-1h10v1c0 .6.4 1 1 1h2c.6 0 1-.4 1-1v-7c.6 0 1-.4 1-1v-1c0-.6-.4-1-1-1h-.469l-.718-3c-.5-1.8-2.107-3-3.907-3zm0 2h7.812c.9 0 1.707.6 1.907 1.5l.843 3.5H6.344l.843-3.5c.2-.9 1.007-1.5 1.907-1.5M7.5 19c.8 0 1.5.7 1.5 1.5S8.3 22 7.5 22S6 21.3 6 20.5S6.7 19 7.5 19m11 0c.8 0 1.5.7 1.5 1.5s-.7 1.5-1.5 1.5s-1.5-.7-1.5-1.5s.7-1.5 1.5-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChargeBattery(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23 5H5C3.3 5 2 6.3 2 8v1H1c-.6 0-1 .4-1 1v6c0 .6.4 1 1 1h1v1c0 1.7 1.3 3 3 3h18c1.7 0 3-1.3 3-3V8c0-1.7-1.3-3-3-3m-1 9l-7-1v4l-8-5l7 1V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chat(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0C4.547 0 0 3.75 0 8.5c0 2.43 1.33 4.548 3.219 6.094a4.778 4.778 0 0 1-.969 2.25a14.4 14.4 0 0 1-.656.781a2.507 2.507 0 0 0-.313.406c-.057.093-.146.197-.187.407c-.042.209.015.553.187.812l.125.219l.25.125c.875.437 1.82.36 2.688.125c.867-.236 1.701-.64 2.5-1.063c.798-.422 1.557-.864 2.156-1.187c.084-.045.138-.056.219-.094C10.796 19.543 13.684 21 16.906 21c.031.004.06 0 .094 0c1.3 0 5.5 4.294 8 2.594c.1-.399-2.198-1.4-2.313-4.375c1.957-1.383 3.22-3.44 3.22-5.719c0-3.372-2.676-6.158-6.25-7.156C18.526 2.664 14.594 0 10 0m0 2c4.547 0 8 3.05 8 6.5S14.547 15 10 15c-.812 0-1.278.332-1.938.688c-.66.355-1.417.796-2.156 1.187c-.64.338-1.25.598-1.812.781c.547-.79 1.118-1.829 1.218-3.281l.032-.563l-.469-.343C3.093 12.22 2 10.423 2 8.5C2 5.05 5.453 2 10 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckBook(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 0C.449 0 0 .449 0 1v14c0 1.094.906 2 2 2h9.531a1 1 0 1 0 0-2H2V5h22v10h-3.5a1 1 0 1 0 0 2H24c1.094 0 2-.906 2-2V1c0-.551-.449-1-1-1zm3.219 7A1.004 1.004 0 0 0 4.5 9h8a1 1 0 1 0 0-2h-8a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.093 0a1.004 1.004 0 0 0-.094 0m13 0a1.004 1.004 0 0 0 .281 2h4a1 1 0 1 0 0-2h-4a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.093 0a1.004 1.004 0 0 0-.094 0m-13 3a1.004 1.004 0 0 0 .281 2h8a1 1 0 1 0 0-2h-8a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.093 0a1.004 1.004 0 0 0-.094 0m16.625.5a.5.5 0 0 0-.188.156l-.562.563c-1.054.137-2.905 1.312-3.375 1.781l-5.313 5.313l2.344 2.343l5.313-5.312c.393-.394 1.653-2.398 1.75-3.469l.53-.531a.5.5 0 0 0-.437-.844a.5.5 0 0 0-.062 0m-10.375 8.719l-3.844 3.843c-.48.48-.707 1.794-.063 2.438c.643.644 1.819.493 2.407-.094l3.844-3.843zm4.469.968a.8.8 0 0 0-.47.25l-4.03 4a.8.8 0 1 0 1.124 1.125l4.032-4a.8.8 0 0 0-.656-1.375"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckFile(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 0C1.8 0 0 1.8 0 4v18c0 2.2 1.8 4 4 4h11.906l-2-2H4c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h7.313c.7.2.687 1.1.687 2v3c0 .6.4 1 1 1h3c1 0 2 0 2 1v9.813l.188.187L20 16.312V8c0-1.1-.988-2.112-2.688-3.813c-.3-.2-.512-.487-.812-.687c-.2-.3-.488-.513-.688-.813C14.113.988 13.1 0 12 0zm20.125 14a.777.777 0 0 0-.438.313l-5.28 7.78l-2.407-2.5c-.3-.3-.706-.3-.906 0l-.906.907c-.3.3-.3.706 0 .906l3.718 3.688c.2.2.482.406.782.406c.3 0 .612-.2.812-.5l6.406-9.406c.2-.1.082-.582-.218-.781l-1.094-.72a.58.58 0 0 0-.469-.093"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Checkmark(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m22.567 4.73l-1.795-1.219a1.09 1.09 0 0 0-1.507.287l-8.787 12.959l-4.039-4.039a1.09 1.09 0 0 0-1.533 0l-1.533 1.536a1.084 1.084 0 0 0 0 1.534L9.582 22c.349.347.895.615 1.387.615s.988-.31 1.307-.774l10.58-15.606a1.085 1.085 0 0 0-.289-1.505"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clapperboard(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m24.219.063l-4.844 1.25l2.438 1.437l-3.844 1L15.5 2.344l-4.813 1.25l2.438 1.437l-3.875 1l-2.438-1.406l-4.843 1.25l3.656 2.156L24.969 2.97l-.75-2.906zM2 8a2 2 0 1 0 0 4a2 2 0 0 0 0-4m4 3l-3 3h5l3-3zm9 0l-3 3h5l3-3zm9 0l-3 3h5v-3zM0 15v8a3 3 0 0 0 3 3h20a3 3 0 0 0 3-3v-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clipboard(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 0c-.798 0-1.448.523-1.719 1H10c-1.094 0-2 .906-2 2H6C4.344 3 3 4.344 3 6v17c0 1.656 1.344 3 3 3h14c1.656 0 3-1.344 3-3V6c0-1.656-1.344-3-3-3h-2c0-1.094-.906-2-2-2h-1.281c-.271-.477-.92-1-1.719-1m-3 3h6v2h-6zM6 6h2.281c.349.597.986 1 1.719 1h6c.733 0 1.37-.403 1.719-1H20v17H6zm2 4v2h10v-2zm0 4v2h7v-2zm0 4v2h10v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloakroom(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 0c-1.645 0-3 1.355-3 3c0 1.257.82 2.092 1.375 2.531c.277.22.504.407.594.5c.027.028.025.03.031.032v.406c-.853.509-8.891 5.294-9.531 5.687c-.531.327-1.07.623-1.563 1.094C.413 13.72 0 14.449 0 15.281c0 .818.393 1.565.969 2.031C1.544 17.78 2.257 18 3 18h1v-2H3c-.361 0-.633-.099-.781-.219c-.149-.12-.219-.233-.219-.5c0-.292.047-.37.281-.594c.235-.223.693-.5 1.25-.843c.435-.268 8.6-5.165 9.469-5.688c.87.523 9.034 5.42 9.469 5.688c.557.342 1.015.62 1.25.844c.234.223.281.301.281.593c0 .267-.07.38-.219.5c-.148.12-.42.219-.781.219h-1v2h1c.743 0 1.456-.221 2.031-.688a2.641 2.641 0 0 0 .969-2.03c0-.833-.413-1.561-.906-2.032c-.494-.47-1.032-.767-1.563-1.094c-.64-.393-8.678-5.178-9.531-5.687V6c0-.579-.308-1.079-.594-1.375c-.285-.296-.558-.48-.781-.656C12.18 3.616 12 3.498 12 3c0-.565.435-1 1-1s1 .435 1 1a1 1 0 1 0 2 0c0-1.645-1.355-3-3-3M6 14a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V15a1 1 0 0 0-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseUpMode(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 0C11.576.997 8.951 2.519 9 7c-.448-1.025-.976-3.303-.813-5.031c0 0-4.187 1.788-4.187 6.437c0 3.483 3.02 5.993 7 6.5V26h3c4.392-.11 12.855-2.495 11.813-13.969c-4.799.71-10.605 5.749-11.813 8.938v-6.063c3.98-.506 7-3.016 7-6.5c0-4.649-4.188-6.437-4.188-6.437C16.977 3.697 16.448 5.975 16 7c.049-4.481-2.576-6.003-3.5-7M.062 12.094s-.198 4.148.844 7.5c1.218 3.913 4.722 6.375 9.094 6.375v-5.407C6.78 14.989.062 12.095.062 12.095z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.25 3C7.095 3 4.549 5.38 4.125 8.406c-2.24.598-3.938 2.56-3.938 4.969c0 2.672 2.081 4.82 4.688 5.125c.727 2.032 2.594 3.5 4.875 3.5c1.562 0 2.845-.793 3.813-1.875c.795.615 1.732 1.063 2.812 1.063c2.058 0 3.752-1.37 4.375-3.22c.032.001.062 0 .094 0c2.728 0 4.968-2.22 4.968-4.937c0-2.318-1.668-4.195-3.843-4.719c-.7-1.926-2.49-3.343-4.657-3.343c-.812 0-1.536.25-2.218.593C13.944 4.08 12.27 3 10.25 3m0 2a4.28 4.28 0 0 1 3.781 2.25l.531 1.031l.938-.719a2.993 2.993 0 0 1 1.813-.593c1.468 0 2.677 1.006 2.968 2.343l.157.72l.718.062a2.952 2.952 0 0 1 2.657 2.937c0 1.625-1.312 2.938-2.97 2.938c-.201 0-.4-.025-.593-.063l-1.094-.219l-.093 1.094c-.127 1.346-1.271 2.407-2.688 2.407c-.858 0-1.6-.39-2.094-1l-.937-1.157l-.719 1.313C12.072 19.342 10.997 20 9.75 20a3.222 3.222 0 0 1-3.188-2.594l-.156-.812h-.843c-.054 0-.075.025-.094.031c-1.823 0-3.282-1.463-3.282-3.25c0-1.663 1.27-3.007 2.907-3.188l.875-.093v-.906C5.984 6.864 7.883 5 10.25 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coins(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 .188c-4.315 0-7.813 1.929-7.813 4.312S13.686 8.813 18 8.813c4.315 0 7.813-1.93 7.813-4.313S22.314.187 18 .187zm7.813 5.593c-.002 2.383-3.498 4.313-7.813 4.313c-4.303 0-7.793-1.909-7.813-4.281V7.5c0 1.018.652 1.95 1.72 2.688c1.08.294 2.042.702 2.843 1.218c.993.252 2.085.406 3.25.406c4.315 0 7.813-1.929 7.813-4.312zm0 3c0 2.383-3.498 4.313-7.813 4.313c-.525 0-1.035-.039-1.531-.094a4.35 4.35 0 0 1 .781 1.781c.249.014.495.031.75.031c4.315 0 7.813-1.929 7.813-4.312zM8 11.187c-4.315 0-7.813 1.93-7.813 4.313S3.686 19.813 8 19.813c4.315 0 7.813-1.93 7.813-4.313S12.314 11.187 8 11.187m17.813.594c-.002 2.383-3.498 4.313-7.813 4.313c-.251 0-.505-.018-.75-.032c-.011.075-.017.175-.031.25c.05.151.093.3.093.47v1c.227.011.455.03.688.03c4.315 0 7.813-1.929 7.813-4.312zm0 3c-.002 2.383-3.498 4.313-7.813 4.313c-.251 0-.505-.018-.75-.032c-.011.075-.017.175-.031.25c.05.15.093.3.093.47v1c.227.011.455.03.688.03c4.315 0 7.813-1.929 7.813-4.312zm-10 2c-.002 2.383-3.498 4.313-7.813 4.313c-4.303 0-7.793-1.909-7.813-4.282V18.5c0 2.383 3.497 4.313 7.813 4.313s7.813-1.93 7.813-4.313zm0 3c-.002 2.383-3.498 4.313-7.813 4.313c-4.303 0-7.793-1.909-7.813-4.282V21.5c0 2.383 3.497 4.313 7.813 4.313s7.813-1.93 7.813-4.313z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Collaborator(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.002 16.417s-1.373-.31-.723-1.679c.803-1.078 1.291-2.301 1.33-2.701c0 0 1.236-.973 1.308-2.434c.077-1.457-.271-1.525-.271-1.525c.52-1.589.688-7.553-3.141-6.841c-.638-1.251-4.638-2.232-6.503 1.101C7.073 4 6.727 6.269 7.453 8.12c-.027.138-.257-.135-.341.721c-.081.825.291 2.008.695 2.556c.176.231.457.4.654.486c0 0 .242 1.47 1.389 2.855c.264 1.099-.827 1.679-.827 1.679c-3.613.695-6.934 2.701-6.934 4.871v1.681c0 2.355 5.485 2.86 10.421 2.86c4.937 0 10.385-1.419 10.385-2.86v-1.681c0-2.17-3.282-4.176-6.893-4.871M19 24h-5v-3h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ColorDropper(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.625 0a4.36 4.36 0 0 0-3.094 1.281l-3.093 3.094l-.688-.688a.972.972 0 0 0-1.375 0L11.312 5.75a.972.972 0 0 0 0 1.375l.657.656l-6.875 6.875c-1.071 1.07-1.675 1.865-2.063 2.563c-.387.697-.508 1.32-.594 1.75c-.085.43-.149.693-.406 1.156c-.257.463-.758 1.102-1.75 2.094a1 1 0 0 0 0 1.437l2.063 2.063a1 1 0 0 0 1.437 0c.992-.992 1.631-1.493 2.094-1.75c.463-.257.727-.32 1.156-.407c.43-.085 1.053-.206 1.75-.593c.698-.388 1.492-.992 2.563-2.063c1.252-1.25 5.935-5.935 6.875-6.875l.656.656c.38.38.995.38 1.375 0l2.063-2.062a.97.97 0 0 0 0-1.375l-.688-.688l3.094-3.093A4.375 4.375 0 0 0 21.625 0m-8.25 9.25l3.375 3.375l-6.844 6.844c-.992.992-1.662 1.493-2.125 1.75c-.462.257-.695.32-1.125.406c-.43.086-1.052.206-1.75.594c-.48.267-1.148.845-1.781 1.406l-.75-.75c.56-.633 1.139-1.3 1.406-1.781c.388-.698.508-1.32.594-1.75c.086-.43.15-.694.406-1.157c.257-.462.758-1.101 1.75-2.093z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Connected(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.875 0a1 1 0 0 0-.594.281l-4.656 4.657c-2.204-1.592-5.261-1.306-7.313.656c0 .076-.855.91-1.624 1.656l-.97-.969A1 1 0 0 0 8.782 6a1 1 0 0 0-.5 1.719l10 10a1.016 1.016 0 1 0 1.438-1.438l-.969-.968l1.656-1.626c1.96-1.96 2.159-5.007.625-7.28l4.688-4.688A1 1 0 0 0 24.875 0M6.906 7.969A1 1 0 0 0 6.781 8a1 1 0 0 0-.5 1.719l.969.969l-1.656 1.624c-1.96 1.96-2.159 5.008-.625 7.282L.28 24.28a1.016 1.016 0 1 0 1.44 1.44l4.656-4.657c2.204 1.592 5.261 1.306 7.313-.656c0-.076.855-.91 1.624-1.656l.97.969a1.016 1.016 0 1 0 1.437-1.438l-10-10a1 1 0 0 0-.813-.312z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Controller(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.906-.031a1 1 0 0 0-.125.031A1 1 0 0 0 23 1s-.003.723-.469 1.469C22.065 3.215 21.233 4 19 4c-3.222 0-5.096 1.148-6.031 2.406C12.034 7.665 12.03 9 12.03 9a1 1 0 1 0 2 0s-.02-.665.531-1.406C15.114 6.852 16.223 6 19 6c2.767 0 4.435-1.215 5.219-2.469C25.003 2.277 25 1 25 1a1 1 0 0 0-1.094-1.031M6.97 9.094c-2.748 0-6.75 4.952-6.75 11.187c0 1.955 1.487 4.657 3.125 4.657c3.276 0 3.896-5.032 9.656-5.032c5.76 0 6.38 5.032 9.656 5.032c1.638 0 3.125-2.702 3.125-4.657c0-6.235-4.002-11.187-6.75-11.187S15.748 12.063 13 12.063s-3.283-2.97-6.031-2.97zm11.78 3.656c.082-.017.163 0 .25 0c.697 0 1.25.551 1.25 1.25A1.24 1.24 0 0 1 19 15.25c-.697 0-1.25-.552-1.25-1.25c0-.61.425-1.132 1-1.25M6 13h2v2h2v2H8v2H6v-2H4v-2h2zm10.75 1.75c.082-.017.163 0 .25 0c.699 0 1.25.552 1.25 1.25A1.24 1.24 0 0 1 17 17.25c-.698 0-1.25-.553-1.25-1.25c0-.61.424-1.132 1-1.25m4 0c.082-.017.163 0 .25 0c.699 0 1.25.552 1.25 1.25A1.24 1.24 0 0 1 21 17.25A1.24 1.24 0 0 1 19.75 16c0-.61.424-1.132 1-1.25m-2 2c.082-.017.163 0 .25 0c.698 0 1.25.551 1.25 1.25A1.24 1.24 0 0 1 19 19.25A1.24 1.24 0 0 1 17.75 18c0-.612.424-1.132 1-1.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cornet(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.656 4.094c-.192.007-.42.163-.625.437c-.409.55-3.01 4.5-6.156 4.5H9.719C7.172 9.031 5 10.951 5 14c0 4.024 3 5 4 5v1.5a1 1 0 1 0 2 0V19h1v1.5a1 1 0 1 0 2 0V19h1v1.5a1 1 0 1 0 2 0v-1.75c2.396-.73 3.063-2.788 3.063-4.188c0-1.314-1.907-1.437-1.907-1.437s.155 2.19-1.156 3.281V12h2.563c3.013 0 4.784 3.338 5.468 4.313c.632.899.969.454.969 0V4.53c0-.303-.152-.444-.344-.437zm-16.593.718a1.194 1.194 0 0 0-.22 2.313V8h2.313v-.906a1.188 1.188 0 0 0 .313-.219a1.19 1.19 0 0 0 .375.25V8h2.312v-.906a1.188 1.188 0 0 0 .313-.219a1.19 1.19 0 0 0 .375.219V8h2.312v-.906a1.188 1.188 0 0 0-.468-2.282h-1.375a1.183 1.183 0 0 0-.126 0a1.19 1.19 0 0 0-.124 0a1.19 1.19 0 0 0-.594.282a1.188 1.188 0 0 0-.781-.282h-1.376a1.183 1.183 0 0 0-.124 0a1.19 1.19 0 0 0-.72.282a1.188 1.188 0 0 0-.78-.282H9.312a1.183 1.183 0 0 0-.124 0a1.19 1.19 0 0 0-.126 0zM.905 6.97A1 1 0 0 0 .781 7A1 1 0 0 0 0 8v5a1 1 0 1 0 2 0v-1h1.344c0-1.869 1.594-3 1.594-3H2V8A1 1 0 0 0 .906 6.969zM11 12h1v5h-1zm3 0h1v5h-1zm-5 .25V17c-.702 0-2-.288-2-2.5c0-1.33 1.108-1.956 2-2.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreateNew(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.438-.063c-.375 0-.732.17-1.032.47l-.718.687l4.218 4.218l.688-.718c.6-.6.6-1.494 0-2.094L23.5.406c-.3-.3-.688-.469-1.063-.469zM20 1.688l-1.094.907l4.5 4.5l1-1zm-1.688 1.625l-9.03 8.938a.975.975 0 0 0-.126.125l-.062.031a.975.975 0 0 0-.219.438l-1.219 4.281a.975.975 0 0 0 1.219 1.219l4.281-1.219a.975.975 0 0 0 .656-.531l8.876-8.782L21 6v.094l-1.188-1.188h.094l-1.593-1.593zM.813 4A1 1 0 0 0 0 5v20a1 1 0 0 0 1 1h20a1 1 0 0 0 1-1V14a1 1 0 1 0-2 0v10H2V6h10a1 1 0 1 0 0-2H1a1 1 0 0 0-.094 0a1 1 0 0 0-.094 0zm9.813 9.813l1.375.093l.094 1.5l-1.375.406l-.531-.53z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cursor(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7 2.014l13.162 12.328l-5.377.488l-.977.088l.406.894l3.263 7.145l-2.406 1.058l-3.113-7.222l-.391-.91l-.722.678l-3.819 3.582zm0-2a2 2 0 0 0-2 2.003l.026 18.127a2.003 2.003 0 0 0 1.999 1.998c.5 0 .991-.188 1.369-.541l2.463-2.311l2.377 5.516a2.006 2.006 0 0 0 1.836 1.208c.274 0 .549-.057.806-.169l2.406-1.059a1.998 1.998 0 0 0 1.013-2.661l-2.498-5.47l3.544-.322a1.998 1.998 0 0 0 1.187-3.451L8.367.554A1.999 1.999 0 0 0 7 .014"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cut(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.5.313s-4.218 1.01-5.5 5.093c0 0-2.352 9.353-3.219 13.344c-.935 4.307-.175 6.596 2.438 6.906c2.02.24 3.418-1.131 3.719-3.656c.298-2.527-.73-4.198-2.75-4.438c-.29-.034-.756-.209-.532-.875c.87-2.603 2.13-2.39 3.25-6.093zm-13.375 9.75c-2.545 0-4.063 1.214-4.063 3.25c0 2.63 2.145 3.67 6.532 3.25c.493-.047.935-.121 1.562-.188c.177-.76.36-1.59.563-2.438c-.578.062-.531-.352-.531-.624c0-2.036-1.52-3.25-4.063-3.25m21.813 0L16.124 11.5a12.867 12.867 0 0 1-1.656 3.25c-.224.326-.428.632-.625 1c3.673-.443 6.406-.813 6.406-.813c4.207-.791 5.688-4.874 5.688-4.874zM4.125 11.687c2.438 0 2.438 1.224 2.438 1.626c0 .401 0 1.624-2.438 1.624s-2.438-1.223-2.438-1.624c0-.402 0-1.626 2.438-1.626m6.781 2.47c.444 0 .813.368.813.812a.818.818 0 0 1-.813.812a.819.819 0 0 1-.812-.812a.82.82 0 0 1 .812-.813zm.656 5.03c.175-.017.338-.011.438 0c.399.046 1.6.204 1.313 2.625c-.288 2.42-1.508 2.265-1.907 2.22c-.399-.047-1.6-.174-1.312-2.595c.215-1.815.945-2.196 1.469-2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CutPaper(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.813 0A1 1 0 0 0 0 1v24a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-3a1 1 0 0 0 0-.094V20l-2-2v3.844A1 1 0 0 0 16 22v2H2v-8h3v-2H2V2h14v2a1 1 0 0 0 0 .094V10l1 1l1-1V4.156A1 1 0 0 0 18 4V1a1 1 0 0 0-1-1H1a1 1 0 0 0-.094 0a1 1 0 0 0-.094 0zM6 6l8 8h-1v2h3l4.156 4.156c-.082.273-.156.546-.156.844c0 1.645 1.355 3 3 3s3-1.355 3-3s-1.355-3-3-3c-.298 0-.571.074-.844.156l-3.218-3.187l3.187-3.125c.281.088.567.156.875.156c1.645 0 3-1.355 3-3s-1.355-3-3-3s-3 1.355-3 3c0 .298.074.571.156.844L16.97 13.03l-5.657-5.625C10.412 6.506 9.2 6 8 6zm17 2c.564 0 1 .436 1 1c0 .564-.436 1-1 1c-.564 0-1-.436-1-1c0-.564.436-1 1-1M7 14v2h4v-2zm16 6c.564 0 1 .436 1 1c0 .564-.436 1-1 1c-.564 0-1-.436-1-1c0-.564.436-1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DefineLocation(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.906-.031a1 1 0 0 0-.125.031A1 1 0 0 0 12 1v1.063C6.737 2.54 2.54 6.737 2.062 12H1a1 1 0 0 0-.094 0A1.001 1.001 0 0 0 1 14h1.063c.477 5.263 4.674 9.46 9.937 9.938V25a1 1 0 1 0 2 0v-1.063c5.263-.477 9.46-4.674 9.938-9.937H25a1 1 0 1 0 0-2h-1.063C23.46 6.737 19.264 2.54 14 2.062V1a1 1 0 0 0-1.094-1.031M12 5.062V6a1 1 0 1 0 2 0v-.938A7.986 7.986 0 0 1 20.938 12H20a1 1 0 0 0-.094 0A1.001 1.001 0 0 0 20 14h.938A7.986 7.986 0 0 1 14 20.938V20a1 1 0 0 0-1.219-1A1 1 0 0 0 12 20v.938A7.986 7.986 0 0 1 5.062 14H6a1 1 0 1 0 0-2h-.938A7.986 7.986 0 0 1 12 5.062"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Delete(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.5-.031c-1.958 0-3.531 1.627-3.531 3.594V4H4c-.551 0-1 .449-1 1v1H2v2h2v15c0 1.645 1.355 3 3 3h12c1.645 0 3-1.355 3-3V8h2V6h-1V5c0-.551-.449-1-1-1h-3.969v-.438c0-1.966-1.573-3.593-3.531-3.593zm0 2.062h3c.804 0 1.469.656 1.469 1.531V4H10.03v-.438c0-.875.665-1.53 1.469-1.53zM6 8h5.125c.124.013.247.031.375.031h3c.128 0 .25-.018.375-.031H20v15c0 .563-.437 1-1 1H7c-.563 0-1-.437-1-1zm2 2v12h2V10zm4 0v12h2V10zm4 0v12h2V10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeleteShield(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.633 5.028a1.074 1.074 0 0 0-.777-.366c-2.295-.06-5.199-2.514-7.119-3.477C14.551.592 13.769.201 13.18.098a1.19 1.19 0 0 0-.359.001c-.589.103-1.372.494-2.556 1.087c-1.921.962-4.825 3.417-7.121 3.476c-.295.008-.577.14-.777.366a1.167 1.167 0 0 0-.291.834c.494 10.023 4.088 16.226 10.396 19.831c.164.093.346.141.528.141s.364-.048.528-.141c6.308-3.605 9.902-9.808 10.396-19.831a1.167 1.167 0 0 0-.291-.834m-6.576 11.056l-.974.973a.46.46 0 0 1-.649 0L13 14.623l-2.434 2.434a.458.458 0 0 1-.649-.002l-.975-.971a.462.462 0 0 1 0-.65l2.434-2.433l-2.434-2.434a.46.46 0 0 1 0-.648l.974-.974a.456.456 0 0 1 .649 0L13 11.379l2.434-2.434a.455.455 0 0 1 .648 0l.975.972a.462.462 0 0 1 0 .65l-2.434 2.434l2.434 2.433a.464.464 0 0 1 0 .65"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeskLamp(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 0a2 2 0 0 0-2 2v3.063C1.669 6.123 0 8.886 0 11.5c0 0 1.275 1.5 6 1.5s6-1.5 6-1.5c0-2.613-1.669-5.378-4-6.438v-.25l6.406-.625c.365.49.936.813 1.594.813c.138 0 .276-.005.406-.031l5.688 6.406A2.01 2.01 0 0 0 22 12c0 .316.088.61.219.875L16.406 23H12a2 2 0 0 0-2 2v1h14v-1a2 2 0 0 0-2-2h-3.281l5.156-9c.042.003.083 0 .125 0a2 2 0 0 0 0-4c-.138 0-.276.005-.406.031l-5.688-6.406A2.01 2.01 0 0 0 18 3a2 2 0 0 0-3.813-.844L8 2.781V2a2 2 0 0 0-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Details(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v26h9V0zm2 2h5v3.75l-.906-.875l-1.282.625L3.22 3.594L2 4.656zm3.813.594a.602.602 0 0 0-.594.594c0 .321.272.593.593.593a.602.602 0 0 0 .594-.594a.602.602 0 0 0-.593-.593M12 4v2h6V4zm8 0v2h6V4zM2 11h5v5H2zm3 .813c-.355.053-.688.374-.688.374s-.035-.029-.093-.062c-.2-.114-.684-.34-1-.031c-.41.399-.031 1.094-.031 1.094s-.596.49-.407 1a.575.575 0 0 0 .25.28c.316.182.781.095.781.095s.21.44.532.624a.566.566 0 0 0 .312.063c.486-.027.657-.844.657-.844s.807-.09.937-.594c.115-.445-.494-.798-.625-.874c0 0 .158-.812-.281-1.063A.52.52 0 0 0 5 11.812zM12 12v2h6v-2zm8 0v2h6v-2zm-15.219.688c.028-.008.063-.018.094 0c.248.141-.344.687-.344.687s.762-.1.719.188c-.042.287-.719-.094-.719-.094s.317.704.032.75c-.286.046-.125-.75-.125-.75s-.63.537-.72.281c-.09-.254.657-.375.657-.375s-.66-.377-.438-.563c.223-.184.532.5.532.5s.115-.573.312-.624M2 19h5v5H2zm2.5.75s-.205.634-.406.719c-.201.085-.844-.219-.844-.219s.287.638.219.844c-.048.147-.719.406-.719.406s.642.23.688.406c.067.261-.188.844-.188.844s.59-.322.781-.219c.193.102.469.719.469.719s.205-.634.406-.719c.201-.085.844.219.844.219s-.29-.639-.219-.844c.054-.159.719-.406.719-.406s-.635-.247-.688-.406c-.085-.256.188-.844.188-.844s-.59.322-.781.219c-.192-.102-.469-.719-.469-.719M12 20v2h6v-2zm8 0v2h6v-2zm-15.5.688a.82.82 0 0 1 .813.812a.82.82 0 0 1-.813.813a.82.82 0 0 1-.813-.813a.82.82 0 0 1 .813-.813z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiningRoom(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 0S5 7 5 8s2.25 3 2.25 3L7 24.344C7 25.448 7.896 26 9 26s2-.552 2-1.656L10.75 11S13 8.958 13 8c0-.958-1-8-1-8h-1v6.5a.5.5 0 0 1-1 0c0-.093-.344-6.5-.344-6.5H8.344S8 6.407 8 6.5a.5.5 0 0 1-1 0V0zm10.719.063C16.13.204 16 .835 16 1.53v22.813c0 1.104.896 1.656 2 1.656s1.969-.553 1.969-1.656c0-5.087-1-7.799-1-10.344c0-1.148 2.031-2.626 2.031-7.969c0-3.268-2.896-5.968-4-5.968c-.104 0-.197-.02-.281 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiplomaOne(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 0C3.908 0 3 .908 3 2v19c0 1.092.908 2 2 2h2v3l3-2l3 2v-7.156a1.997 1.997 0 0 1-.438.187a1.923 1.923 0 0 1-1.093 1a1.922 1.922 0 0 1-.656.094c-.28 0-.56-.036-.813-.156a1.88 1.88 0 0 1-.813.156c-.224 0-.443-.017-.656-.094a1.941 1.941 0 0 1-1.094-1A2.018 2.018 0 0 1 7 18.844V21H5V2h16v19h-5a1 1 0 1 0 0 2h5c1.092 0 2-.908 2-2V2c0-1.092-.908-2-2-2zm2.813 5A1.001 1.001 0 0 0 8 7h10a1 1 0 1 0 0-2H8a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.093 0m0 4A1.001 1.001 0 0 0 8 11h7a1 1 0 1 0 0-2H8a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.093 0M10 12.656a.955.955 0 0 0-.906.656c-.041.014-.086.016-.125.032a.936.936 0 0 0-1.094.094a.95.95 0 0 0-.281 1.062c-.027.041-.038.082-.063.125a.95.95 0 0 0-.781.781a.929.929 0 0 0 .469.969c.007.063-.013.127 0 .188a.945.945 0 0 0-.094 1.062c.194.337.58.525.969.469c.05.047.101.082.156.125c.01.387.257.74.625.875c.364.133.776.005 1.031-.282c.037.002.056.032.094.032c.041 0 .084-.03.125-.032a.929.929 0 0 0 .688.313c.108 0 .206.007.312-.031a.957.957 0 0 0 .625-.875c.053-.042.105-.078.156-.125a.965.965 0 0 0 .969-.469a.945.945 0 0 0-.094-1.063c.014-.06-.009-.126 0-.187a.923.923 0 0 0 .469-.969a.949.949 0 0 0-.781-.781c-.024-.044-.036-.083-.063-.125a.948.948 0 0 0-.281-1.063a.935.935 0 0 0-1.094-.093c-.04-.016-.084-.018-.125-.031a.955.955 0 0 0-.906-.657m0 1.75c.886 0 1.594.708 1.594 1.594c0 .888-.707 1.594-1.594 1.594A1.585 1.585 0 0 1 8.406 16c0-.887.708-1.594 1.594-1.594"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Disconnected(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.875 0a1 1 0 0 0-.594.281l-1.656 1.656c-2.204-1.59-5.261-1.305-7.313.657c0 .076-.855.91-1.624 1.656l-.97-.969a1 1 0 0 0-.812-.312a1 1 0 0 0-.125.031a1 1 0 0 0-.5 1.719L13.562 7l-3.28 3.281a1.016 1.016 0 1 0 1.437 1.438L15 8.438L17.563 11l-3.282 3.281a1.016 1.016 0 1 0 1.438 1.438L19 12.438l2.281 2.28a1.016 1.016 0 1 0 1.438-1.437l-.969-.969l1.656-1.624c1.96-1.96 2.159-5.008.625-7.282L25.72 1.72A1 1 0 0 0 24.875 0M3.906 10.969a1 1 0 0 0-.125.031a1 1 0 0 0-.5 1.719l.969.969l-1.656 1.624c-1.96 1.96-2.159 5.008-.625 7.282L.28 24.28a1.016 1.016 0 1 0 1.44 1.44l1.656-1.657c2.204 1.592 5.261 1.306 7.313-.656c0-.076.855-.91 1.624-1.656l.97.969a1.016 1.016 0 1 0 1.437-1.438l-10-10a1 1 0 0 0-.813-.312z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoctorsBag(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 1C8.355 1 7 2.355 7 4v2h2V4c0-.563.437-1 1-1h6c.563 0 1 .437 1 1v2h2V4c0-1.645-1.355-3-3-3zM3 7a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h20c1.656 0 3-1.344 3-3V10a3 3 0 0 0-3-3zm10 2.906A6.089 6.089 0 0 1 19.094 16A6.089 6.089 0 0 1 13 22.094A6.089 6.089 0 0 1 6.906 16A6.089 6.089 0 0 1 13 9.906M12 13v2h-2v2h2v2h2v-2h2v-2h-2v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Edit(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.094.25a2.245 2.245 0 0 0-1.625.656l-1 1.031l6.593 6.625l1-1.03a2.319 2.319 0 0 0 0-3.282L21.75.937A2.36 2.36 0 0 0 20.094.25m-3.75 2.594l-1.563 1.5l6.875 6.875L23.25 9.75zM13.78 5.438L2.97 16.155a.975.975 0 0 0-.5.625L.156 24.625a.975.975 0 0 0 1.219 1.219l7.844-2.313a.975.975 0 0 0 .781-.656l10.656-10.563l-1.468-1.468L8.25 21.813l-4.406 1.28l-.938-.937l1.344-4.593L15.094 6.75zm2.375 2.406l-10.968 11l1.593.343l.219 1.47l11-10.97z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditFile(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 0C1.8 0 0 1.8 0 4v17c0 2.2 1.8 4 4 4h8.188v-.188l.5-1.812H4c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h6.313c.7.2.687 1.1.687 2v3c0 .6.4 1 1 1h3c1 0 2 0 2 1v8l2-2V8c0-1.1-.988-2.112-2.688-3.813c-.3-.2-.512-.487-.812-.687c-.2-.3-.488-.513-.688-.813C13.113.988 12.1 0 11 0zm19.344 14.094c-.275 0-.55.112-.75.312l-1.188 1.188l3 3l1.188-1.188c.4-.4.4-1 0-1.5l-1.5-1.5a1.07 1.07 0 0 0-.75-.312m-3.032 2.5l-4.718 5c-.1 0-.188.118-.188.218l-1.094 3.594c-.1.1-.006.306.094.407c.1.1.181.093.281.093h.126l3.593-1.093c.1 0 .088-.025.188-.125l4.906-4.875zM16 22.094l1.5.312l.313 1.594l-2 .5l-.407-.406z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditImage(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2a2 2 0 0 0-2 2c0 .738 1.156 2.057 1.156 3C1.156 7.943 0 10.921 0 13s1.156 4.991 1.156 6C1.156 20.009 0 21.262 0 22a2 2 0 0 0 2 2c.738 0 3.124-1.125 4.5-1.125S8.396 24 9.5 24s2.131-1.125 3.5-1.125c1.369 0 1.896 1.125 3 1.125s2.563-1.125 3.5-1.125c.937 0 3.762 1.125 4.5 1.125a2 2 0 0 0 2-2c0-.738-1.156-2.057-1.156-3c0-.943 1.156-4.05 1.156-6c0-1.95-1.156-4.991-1.156-6C24.844 5.991 26 4.738 26 4a2 2 0 0 0-2-2c-.738 0-3.124 1.156-4.5 1.156S17.604 2 16.5 2S14.369 3.156 13 3.156C11.631 3.156 11.104 2 10 2S7.438 3.156 6.5 3.156C5.562 3.156 2.738 2 2 2m2 3h18c.551 0 1 .449 1 1v14c0 .551-.449 1-1 1H4c-.551 0-1-.449-1-1V6c0-.551.449-1 1-1m12.281 2.125a1.044 1.044 0 0 0-.75.313l-.906.906l3.031 3.031l.907-.906a1.052 1.052 0 0 0 0-1.5L17.03 7.438a1.052 1.052 0 0 0-.75-.313zM13.563 9.25l-1.25 1.281l-3.938 3.906a.45.45 0 0 0-.25.313l-1.063 3.625a.45.45 0 0 0 .563.563l3.625-1.063a.455.455 0 0 0 .281-.219L15.72 13.5l1.031-1.063zm-4.626 5.938l1.188.656l.688 1.219l-2.032.593l-.437-.437l.594-2.032z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eject(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 4.063c-.485 0-.76.27-1.156.718L5.25 13.094c-.203.403-.17.897.063 1.281c.232.386.65.625 1.093.625h13.188c.443 0 .86-.239 1.093-.625a1.33 1.33 0 0 0 .188-.688c0-.2-.033-.408-.125-.593L14.156 4.78c-.376-.4-.671-.718-1.156-.718zM6 17c-.551 0-1 .449-1 1v3c0 .551.449 1 1 1h14c.551 0 1-.449 1-1v-3c0-.551-.449-1-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmptyBattery(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23 7c.551 0 1 .448 1 1v10c0 .552-.449 1-1 1H5c-.551 0-1-.448-1-1v-3H2v-4h2V8c0-.552.449-1 1-1zm0-2H5a3 3 0 0 0-3 3v1H1c-.551 0-1 .449-1 1v6c0 .551.449 1 1 1h1v1a3 3 0 0 0 3 3h18a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmptyFlag(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.906-.031A1 1 0 0 0 5.781 0a1 1 0 0 0-.718.656L5 .625v.188A1 1 0 0 0 5 1v24a1 1 0 1 0 2 0v-8.531L23.375 9L7 1.531V1A1 1 0 0 0 5.906-.031M7 3.719L18.563 9L7 14.281V3.72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func End(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 5c-.551 0-1 .449-1 1v4.875L7.906 5.25a1.339 1.339 0 0 0-1.281.063C6.239 5.545 6 5.963 6 6.405v13.188c0 .443.239.86.625 1.093c.208.127.449.188.688.188c.2 0 .408-.033.593-.125L15 15.125V20c0 .551.449 1 1 1h3c.551 0 1-.449 1-1V6c0-.551-.449-1-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacialRecognitionScan(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.813 0A1 1 0 0 0 0 1v3a1 1 0 1 0 2 0V2h2a1 1 0 1 0 0-2H1a1 1 0 0 0-.094 0a1 1 0 0 0-.094 0zm20.906 0A1.004 1.004 0 0 0 22 2h2v2a1 1 0 1 0 2 0V1a1 1 0 0 0-1-1h-3a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.093 0a1.004 1.004 0 0 0-.094 0M13 1C4.801 1 3.85 9.008 4.031 12.813C3.413 13.254 3 13.969 3 15H1a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.094 0A1.004 1.004 0 0 0 1 17h24a1 1 0 1 0 0-2h-2c0-.982-.376-1.746-1.031-2.219C22.142 8.964 21.176 1 13 1m-3 7c0 3.16 9 1.827 9 4v2.375l1.125-.125c.015-.001.04 0 .094 0c.71 0 .781.344.781.75h-3.094c.054-.154.094-.327.094-.5a1.5 1.5 0 0 0-3 0c0 .173.04.346.094.5h-4.188c.054-.154.094-.327.094-.5a1.5 1.5 0 0 0-3 0c0 .173.04.346.094.5H5c0-.275.003-.75.906-.75L7 14.375V12c0-2.723 1.824-1.115 3-4M5.125 18c1.94 4.639 4.594 7 7.875 7c3.282 0 5.936-2.36 7.875-7h-2.188c-1.54 3.315-3.454 5-5.687 5s-4.146-1.685-5.688-5H5.126zM.906 20.969A1 1 0 0 0 .781 21A1 1 0 0 0 0 22v3a1 1 0 0 0 1 1h3a1 1 0 1 0 0-2H2v-2a1 1 0 0 0-1.094-1.031m24 0a1 1 0 0 0-.125.031A1 1 0 0 0 24 22v2h-2a1 1 0 1 0 0 2h3a1 1 0 0 0 1-1v-3a1 1 0 0 0-1.094-1.031"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Factory(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.469.993a1 1 0 0 0-.999-.95h-2.849a1 1 0 0 0-.996.91L18 19.268V12h-2V9a1.002 1.002 0 0 0-1.554-.832L8.697 12H8V9a1.002 1.002 0 0 0-1.554-.832l-6 4A1 1 0 0 0 0 13v12a1 1 0 0 0 1 1h24.02c.553 0 .949-.352.949-.904c0-.086-1.5-24.103-1.5-24.103M5 14v3H3v-3zm5 0v3H8v-3zm5 0v3h-2v-3zM5 19v3H3v-3zm5 0v3H8v-3zm5 0v3h-2v-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fan(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 .188C7.582.188 3.187 4.58 3.187 10c0 5.417 4.396 9.811 9.813 9.813c5.416-.002 9.813-4.396 9.813-9.813c0-5.421-4.399-9.813-9.813-9.813zm0 1.125c4.794 0 8.688 3.89 8.688 8.687c0 4.796-3.892 8.686-8.688 8.688a8.687 8.687 0 1 1 0-17.375m-2.469 1.656c-1.176.115-2.349 1.63-1.625 3c.68 1.284 2.096 1.16 2.531 2.625c.405-.556 1.206-1.184 2.563-.531c.878-3.252-.975-5.332-3.469-5.094m8.031 5.5c-1.45-.055-2.077 1.229-3.562.875c.277.627.432 1.65-.813 2.5c2.377 2.385 5.085 1.81 6.126-.469c.489-1.075-.201-2.85-1.75-2.906zm-5.812.25c-.59.12-1.031.655-1.031 1.281c0 .716.565 1.281 1.281 1.281s1.281-.565 1.281-1.281S13.716 8.719 13 8.719c-.09 0-.166-.018-.25 0m-1.938 1.375c-3.254.865-4.11 3.521-2.656 5.562c.686.961 2.582 1.22 3.406-.094c.773-1.228-.05-2.39 1-3.5c-.681-.073-1.637-.466-1.75-1.968M11 21c0 2.001-4 2.344-4 4v1h12v-1c0-1.656-4-1.999-4-4c0 0-.379.313-2 .313S11 21 11 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Faq(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 0c-1.7 0-3 1.3-3 3v6c0 1.7 1.3 3 3 3h6l4 4v-4c1.7 0 3-1.3 3-3V3c0-1.7-1.3-3-3-3zm4.188 3h1.718l1.688 6h-1.5l-.407-1.5h-1.5L16.813 9H15.5zM18 4c-.1.4-.212.888-.313 1.188l-.28 1.312h1.187l-.282-1.313C18.113 4.888 18 4.4 18 4M3 10c-1.7 0-3 1.3-3 3v6c0 1.7 1.3 3 3 3v4l4-4h6c1.7 0 3-1.3 3-3v-6h-3c-1.9 0-3.406-1.3-3.906-3zm4.594 2.906c1.7 0 2.5 1.4 2.5 3c0 1.4-.481 2.288-1.281 2.688c.4.2.874.306 1.374.406l-.374 1c-.7-.2-1.426-.512-2.126-.813c-.1-.1-.275-.093-.375-.093C6.112 18.994 5 18 5 16c0-1.7.994-3.094 2.594-3.094m0 1.094c-.8 0-1.188.9-1.188 2c0 1.2.388 2 1.188 2c.8 0 1.218-.9 1.218-2s-.418-2-1.218-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastForward(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.25 5.125c-.219.01-.433.071-.625.188c-.386.232-.625.65-.625 1.093v13.188c0 .443.239.86.625 1.093c.208.127.449.188.688.188c.2 0 .408-.033.593-.125L12 15.906v3.688c0 .443.239.86.625 1.093c.208.127.448.188.688.188c.2 0 .408-.033.593-.125l8.313-6.594c.4-.376.718-.671.718-1.156c0-.485-.27-.76-.718-1.156L13.906 5.25a1.339 1.339 0 0 0-1.281.063c-.386.232-.625.65-.625 1.093v3.688L5.906 5.25a1.35 1.35 0 0 0-.656-.125"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilledFlag(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.906-.031A1 1 0 0 0 5.781 0a1 1 0 0 0-.718.656L5 .625v.188A1 1 0 0 0 5 1v24a1 1 0 1 0 2 0v-8.531L23.375 9L7 1.531V1A1 1 0 0 0 5.906-.031"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FingerprintScan(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.813 0A1 1 0 0 0 0 1v3a1 1 0 1 0 2 0V2h2a1 1 0 1 0 0-2H1a1 1 0 0 0-.094 0a1 1 0 0 0-.094 0zm20.906 0A1.004 1.004 0 0 0 22 2h2v2a1 1 0 1 0 2 0V1a1 1 0 0 0-1-1h-3a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.093 0a1.004 1.004 0 0 0-.094 0M13.5 2.313c-1.535.057-3.907.39-6.156 2.062c-3.905 2.904-4.219 7.938-4.219 7.938a1.003 1.003 0 1 0 2 .156S5.524 8.259 8.563 6c2.432-1.808 4.83-1.688 6.156-1.688a1 1 0 1 0 0-2c-.292 0-.707-.019-1.219 0zm.656 3.28c-1.187.01-2.448.363-3.719 1.063c-1.87 1.031-3.097 2.64-3.562 4.688c-.27 1.19-.226 2.452-.031 3.656H1a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.094 0A1.004 1.004 0 0 0 1 17h24a1 1 0 1 0 0-2h-5c-.783-.76-1.327-1.816-1.844-2.844c-.955-1.899-2.26-4.498-5.562-3.062c-1.072.467-1.836 1.287-2.25 2.406c-.42 1.135-.418 2.392-.188 3.5H8.875c-.21-1.075-.262-2.201-.031-3.219c.34-1.495 1.187-2.617 2.562-3.375c1.528-.842 2.906-1.045 4.156-.594c2.35.85 3.905 3.81 4.813 6.125c.202.516.8.762 1.313.563c.514-.202.765-.799.562-1.313c-1.059-2.7-2.977-6.146-6.031-7.25a5.822 5.822 0 0 0-2.063-.343zm.156 5.063c.769-.022 1.204.667 2.063 2.375c.312.622.665 1.312 1.094 1.969h-1.782c-.658-1.067-.686-1.919-.687-2a.983.983 0 0 0-.969-1c-.559-.037-1.017.417-1.031.969c-.001.077.008.911.469 2.031h-1.25c-.284-1.001-.287-2.034 0-2.813c.227-.612.61-1.03 1.187-1.28c.355-.155.65-.243.906-.25zM7.689 18c.324.751.7 1.418 1.093 1.969c1.601 2.24 3.127 3.456 5.781 4.593a.888.888 0 0 0 .376.063c.387 0 .775-.216.937-.594a1.002 1.002 0 0 0-.531-1.312c-2.33-1-3.542-1.954-4.938-3.907a7.825 7.825 0 0 1-.5-.812h-2.22zm3.718 0c1.871 3.052 4.696 4.781 7.938 4.781c.079 0 .17.002.25 0a1.001 1.001 0 0 0-.032-2h-.187c-2.135 0-4.08-.994-5.531-2.781zm4.25 0c.846.745 1.968 1.463 3.438 2.094a.93.93 0 0 0 .375.093c.388 0 .774-.214.937-.593a1.002 1.002 0 0 0-.531-1.313c-.203-.088-.378-.191-.563-.281zM.906 20.969A1 1 0 0 0 .781 21A1 1 0 0 0 0 22v3a1 1 0 0 0 1 1h3a1 1 0 1 0 0-2H2v-2a1 1 0 0 0-1.094-1.031m24 0a1 1 0 0 0-.125.031A1 1 0 0 0 24 22v2h-2a1 1 0 1 0 0 2h3a1 1 0 0 0 1-1v-3a1 1 0 0 0-1.094-1.031"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func First(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.125 0H4.875A4.874 4.874 0 0 0 0 4.875v16.25A4.874 4.874 0 0 0 4.875 26h16.25A4.874 4.874 0 0 0 26 21.125V4.875A4.874 4.874 0 0 0 21.125 0M18 17.949a.964.964 0 0 1-.479.817a.986.986 0 0 1-.952.039L11 14.382V18c0 .551-.449 1-1 1H8c-.551 0-1-.449-1-1V8c0-.551.449-1 1-1h2c.551 0 1 .449 1 1v3.618l5.569-4.425a.99.99 0 0 1 .952.04c.29.174.479.485.479.817z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FullBattery(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23 5H5C3.3 5 2 6.3 2 8v1H1c-.6 0-1 .4-1 1v6c0 .6.4 1 1 1h1v1c0 1.7 1.3 3 3 3h18c1.7 0 3-1.3 3-3V8c0-1.7-1.3-3-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FullTrash(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.875 0a1 1 0 0 0-.594.281L5.562 5H3c-.551 0-1 .449-1 1v2c0 .551.449 1 1 1h.25l2.281 13.719a.998.998 0 0 0 0 .062c.163.788.469 1.541 1.032 2.157A3.258 3.258 0 0 0 8.938 26h8.124a3.26 3.26 0 0 0 2.375-1.031c.571-.615.883-1.405 1.032-2.219a.998.998 0 0 0 0-.031L22.78 9H23c.551 0 1-.449 1-1V6c0-.551-.449-1-1-1h-1.563l-2.812-3.5a.813.813 0 0 0-.719-.313a.813.813 0 0 0-.343.125L14.688 3.25L11.717.281A1 1 0 0 0 10.876 0zM11 2.438L13.563 5H8.436L11 2.437zm6.844.656L19.375 5h-2.938l-.593-.594zM5.25 9h.688l1.187 1.188l-1.438 1.406zm2.094 0h.937l-.469.469zm2.312 0h1.688l.906.906l-2 2l-1.75-1.75zm3.125 0h.344l-.156.188L12.78 9zm1.781 0h1.688l1.156 1.156l-1.75 1.75l-2-2.031zm3.063 0h.938l-.47.469L17.626 9zm2.344 0h.812l-.437 2.688l-1.532-1.532zm-7.032 1.594l2.032 2l-2.031 2l-2-2l2-2zm-5.124.281l1.718 1.719l-2 2l-1.625-1.625l-.031-.156zm10.28 0l2 2l-1.718 1.75l-2-2.031l1.719-1.719zm-7.843 2.438l2 2l-2 2l-2-2zm5.406 0l2.031 2l-2 2l-2.03-2zm4.188 1.25l-.219 1.312l-.563-.563l.782-.75zm-13.657.093l.657.656l-.469.47zM7.532 16l2 2l-2 2.031l-.562-.562l-.407-2.5zm5.407 0l2.03 2.031l-2 2L10.939 18zm5.437 0l1.063 1.063l-.407 2.28l-.656.657l-2-2zm-8.125 2.719l2 2l-2 2.031l-2-2zm5.406 0l2 2l-2 2l-2-2zm-8.094 2.718l2 2L9 24h-.063c-.391 0-.621-.13-.874-.406a2.645 2.645 0 0 1-.594-1.188v-.031l-.125-.75l.218-.188zm5.407 0l2 2l-.563.563H11.5l-.563-.563l2.032-2zm5.406 0l.281.282l-.125.656c-.002.01.002.02 0 .031c-.095.49-.316.922-.562 1.188c-.252.27-.509.406-.907.406h-.125l-.562-.563z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Future(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 0v2C6.955 2 2 6.955 2 13s4.955 11 11 11s11-4.955 11-11c0-2.947-1.11-5.61-3.094-7.594L19.5 6.813C21.117 8.428 22 10.546 22 13c0 4.955-4.045 9-9 9s-9-4.045-9-9s4.045-9 9-9v2l5-3zm2.094 6.563l-2.5 5A1.483 1.483 0 0 0 11.5 13a1.5 1.5 0 0 0 1.5 1.5h.063l3.218 3.219l1.438-1.438l-3.219-3.218V13c0-.197-.056-.39-.125-.563l2.531-5l-1.812-.875z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GeoFence(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 0a6 6 0 0 0-6 6c0 3.314 5.219 11.543 5.219 16h1.562C13.781 17.558 19 9.044 19 6a6 6 0 0 0-6-6m0 3a3 3 0 1 1 0 6a3 3 0 0 1 0-6M6.844 17.281a1 1 0 0 0-.063.032c-1.845.324-3.409.768-4.593 1.375c-.593.303-1.104.662-1.5 1.093c-.397.432-.688.993-.688 1.625c0 .948.628 1.689 1.375 2.219c.747.53 1.69.935 2.844 1.281C6.524 25.598 9.605 26 13 26c3.395 0 6.476-.402 8.781-1.094c1.153-.346 2.097-.75 2.844-1.281c.747-.53 1.375-1.271 1.375-2.219c0-.634-.29-1.224-.688-1.656c-.398-.432-.904-.759-1.5-1.063c-1.19-.607-2.769-1.051-4.625-1.375a1 1 0 1 0-.343 1.97c1.723.3 3.155.725 4.062 1.187c.454.23.773.477.938.656c.165.179.156.25.156.281c0 .047-.064.263-.531.594c-.467.331-1.263.694-2.282 1c-2.037.611-4.955 1-8.187 1c-3.232 0-6.15-.389-8.188-1c-1.018-.306-1.814-.669-2.28-1c-.468-.331-.532-.547-.532-.594c0-.03-.007-.103.156-.281c.164-.178.487-.426.938-.656c.9-.461 2.318-.886 4.031-1.188a1.01 1.01 0 1 0-.281-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ghost(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 0c-1.937 0-3.428.508-4.5 1.344c-1.072.835-1.688 1.959-2.125 3c-.437 1.04-.717 2.035-1 2.75c-.283.715-.517.991-.719 1.062C1.904 9.13 0 11.52 0 14.656a1 1 0 0 0 .938 1c1.257.067 1.862.362 2.187.656c.325.295.474.685.563 1.22c.088.534.097 1.16.218 1.843c.061.341.144.702.375 1.063c.232.36.641.692 1.094.843c1.06.354 1.737.573 2.156.813c.42.24.669.453 1 1.156c.316.668.744 1.4 1.5 1.938c.756.537 1.78.828 3.063.812c2.101-.026 3.355-1.113 4-2.063c.322-.474.549-.887.75-1.156c.2-.269.32-.374.531-.437c1.368-.412 2.236-1.15 2.719-1.969c.482-.82.61-1.602.75-2.156c.14-.555.226-.797.531-1c.305-.204 1.031-.426 2.563-.344a1 1 0 0 0 1.062-1c0-1.379.02-2.979-.594-4.469c-.614-1.49-1.972-2.764-4.25-3.281c-.024-.006-.052.01-.156-.125c-.104-.134-.235-.4-.375-.75c-.28-.7-.581-1.721-1.031-2.781c-.45-1.06-1.091-2.198-2.157-3.063C16.372.542 14.905 0 13 0m0 2c1.538 0 2.481.396 3.188.969c.706.573 1.18 1.38 1.562 2.281c.383.9.653 1.88 1 2.75c.174.435.361.838.656 1.219c.295.38.76.75 1.313.875c1.811.41 2.443 1.122 2.843 2.094c.299.723.286 1.752.313 2.75c-1.068.05-1.962.203-2.594.624c-.863.576-1.2 1.466-1.375 2.157c-.175.69-.292 1.25-.531 1.656c-.239.405-.58.757-1.594 1.063c-.712.214-1.224.703-1.562 1.156c-.338.452-.543.898-.782 1.25c-.477.704-.822 1.136-2.374 1.156c-.983.013-1.512-.179-1.876-.438c-.363-.258-.596-.631-.843-1.156c-.447-.946-1.074-1.639-1.813-2.062c-.739-.423-1.51-.628-2.531-.969c-.056-.019-.014.027-.031 0a1.333 1.333 0 0 1-.094-.344c-.07-.396-.092-1.084-.219-1.843c-.126-.76-.427-1.655-1.187-2.344c-.565-.512-1.387-.82-2.375-1c.26-1.895 1.369-3.158 3.219-3.813c1.06-.375 1.585-1.299 1.937-2.187c.352-.889.602-1.877.969-2.75c.366-.873.832-1.612 1.531-2.157C10.449 2.393 11.403 2 13 2m-2.5 4C9.672 6 9 6.895 9 8s.672 2 1.5 2S12 9.105 12 8s-.672-2-1.5-2m4 0c-.828 0-1.5.895-1.5 2s.672 2 1.5 2S16 9.105 16 8s-.672-2-1.5-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeEarth(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.969.156a1 1 0 0 0-.438.188c-3.21 2.107-5.343 5.724-5.343 9.844c0 6.171 4.767 11.243 10.812 11.75v1.187c-2.013.231-3.469 1.063-3.469 1.063a.943.943 0 0 0 .469 1.75h8a.943.943 0 0 0 .469-1.75s-1.456-.832-3.469-1.063v-1.156a11.78 11.78 0 0 0 5.469-1.907a1 1 0 1 0-1.094-1.656A9.705 9.705 0 0 1 14.187 20a1 1 0 0 0-.374 0a9.809 9.809 0 0 1-9.626-9.813C4.188 6.753 5.96 3.75 8.625 2A1 1 0 0 0 7.969.156M14 1.188a8.71 8.71 0 0 0-4.406 1.187c-.129.058-.253.14-.375.219a2.773 2.773 0 0 0-.438.344C6.616 4.546 5.187 7.107 5.187 10c0 1.713.492 3.3 1.344 4.656a.802.802 0 0 0 .125.188c1.582 2.377 4.287 3.969 7.344 3.969a8.71 8.71 0 0 0 4.781-1.407c.046-.029.082-.062.125-.093c2.345-1.588 3.907-4.281 3.907-7.313a8.702 8.702 0 0 0-1.282-4.563a.8.8 0 0 0-.218-.343c-.009-.012-.024-.02-.032-.032A.8.8 0 0 0 21.22 5C19.623 2.714 16.986 1.187 14 1.187zm0 1.625a7.34 7.34 0 0 1 1.906.25l-1.094.718c-.36-.28-.73-.534-1.093-.75c-.117-.069-.226-.125-.344-.187A7.58 7.58 0 0 1 14 2.813m-3.094.906c.537-.01 1.218.192 2 .656c.149.088.318.207.469.313l-3.938 2.53c-.033-.18-.075-.36-.093-.53c-.157-1.448.216-2.387.781-2.75c.008-.005.024.004.031 0c.208-.127.441-.214.75-.22zm6.75.125a.8.8 0 0 0 .032 0a7.103 7.103 0 0 1 1.906 1.625l-2 1.281c-.489-.684-.989-1.327-1.532-1.875zm-3 1.937a14.29 14.29 0 0 1 1.563 1.844l-5.282 3.406a14.304 14.304 0 0 1-1.03-2.187zm-6.906.657c.008.135.016.269.031.406c.045.413.148.844.25 1.281l-1.125.719a7.112 7.112 0 0 1 .844-2.406m12.688.375c.36.732.625 1.524.718 2.375a.8.8 0 0 0-.062.062L19.5 10.281c-.27-.719-.615-1.46-1.031-2.187zm-3.344 2.156a14.15 14.15 0 0 1 1.031 2.187l-4.781 3.063a14.29 14.29 0 0 1-1.563-1.844zm-8.594.75c.274.723.642 1.457 1.063 2.187l-2 1.281a7.127 7.127 0 0 1-.72-2.406L8.5 9.72zm12.594 1.437a7.112 7.112 0 0 1-.844 2.406c-.008-.145-.015-.29-.031-.437c-.044-.402-.12-.826-.219-1.25zm-2.532 1.625c.032.175.076.366.094.531c.157 1.448-.216 2.356-.781 2.72c-.565.363-1.53.336-2.781-.407c-.149-.088-.318-.207-.469-.313zm-8.156.469a15.19 15.19 0 0 0 1.563 1.875l-1.656 1.063a7.15 7.15 0 0 1-1.907-1.625zm2.813 2.969c.349.269.71.51 1.062.718c.138.082.267.147.406.22a7.79 7.79 0 0 1-.687.03a7.348 7.348 0 0 1-1.906-.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GpsReceiving(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.813 0A1.001 1.001 0 0 0 14 2c4.442 0 8 3.513 8 8a1 1 0 1 0 2 0c0-5.555-4.492-10-10-10a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.094 0zm0 4A1.001 1.001 0 0 0 14 6c1.594 0 2.53.341 3.094.906c.564.565.906 1.505.906 3.094a1 1 0 1 0 2 0c0-1.88-.434-3.432-1.5-4.5S15.888 4 14 4a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.094 0zM6.718 5.125c-.91-.005-1.679.241-2.219.781c-1.862 1.862-4.643 7.481.75 12.875c4.933 4.934 11.012 2.612 12.875.75c1.485-1.485.808-4.667-1.438-7.812c-.422.34-.921 1.48-.687 1.875c1.111 1.874 1.431 3.538.656 4.312c-1.193 1.193-4.527-.183-7.437-3.093C6.307 11.9 4.93 8.57 6.125 7.375c.73-.731 2.251-.5 4 .469c.409.226 1.572-.26 1.906-.688c-1.91-1.303-3.797-2.022-5.312-2.031zM14 8c-.887 0-1.644.575-1.906 1.375l-2.688.313a.828.828 0 1 0 .188 1.624l2.687-.28c.17.284.403.517.688.687l-.281 2.687a.828.828 0 1 0 1.624.188l.313-2.688A2.005 2.005 0 0 0 16 10a2 2 0 0 0-2-2M7.062 22.063C5.369 22.838 4.22 24.31 4.22 26h11.687c0-1.044-.758-2.263-1.687-2.781c-.636.115-1.317.187-2 .187a10.46 10.46 0 0 1-5.156-1.343z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GroupIcon(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.531 1c-1.798 0-3.367.927-4.062 2.688c1.459 1.204 2.469 3.067 2.469 5.593c0 2.697-1.272 5.162-2.594 6.75c2.106.797 5.402 2.394 6.344 4.844c3.318-.184 6.28-.852 6.28-2.75V17.5c0-1.74-3.034-3.443-5.718-4.344c-.122-.04-.89-.226-.406-1.719c1.26-1.316 2.125-3.446 2.125-5.53C21.969 2.696 19.973 1 17.53 1zM8.97 4.094c-2.6 0-4.844 1.775-4.844 5.187c0 2.23 1.06 4.506 2.406 5.906c.525 1.399-.428 2.395-.625 2.47C3.186 18.653 0 20.452 0 22.25v.688c0 2.449 4.671 3 9 3c4.334 0 8.969-.551 8.969-3v-.688c0-1.852-3.208-3.635-6.063-4.594c-.13-.043-.951-.913-.437-2.5h-.031c1.34-1.4 2.5-3.654 2.5-5.875c0-3.412-2.371-5.187-4.97-5.187z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Guitar(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.594 0c-.256 0-.524.086-.719.281l-.406.438l-.188-.156a.497.497 0 0 0-.718 0a.497.497 0 0 0 0 .718l.156.188l-.25.25l-.188-.157a.497.497 0 0 0-.718 0a.497.497 0 0 0 0 .72l.156.187l-.188.187a.96.96 0 0 0-.25.938l-4.375 4.375c-.903-.79-2.269-1.719-3.75-1.719c-.657 0-1.332.2-2 .656c-1.482 1.013-1.198 2.548-2.125 3.438c-1.04 1-3.883.194-6.5 2.812c-2.617 2.616-1.457 6.235 1.875 9.438c1.998 2.078 4.148 3.312 6.125 3.312c1.191 0 2.328-.452 3.313-1.437c2.617-2.617 1.811-5.458 2.812-6.5c.89-.925 2.455-.611 3.469-2.094c1.487-2.173.052-4.471-1.094-5.781l4.375-4.375c.33.084.679.009.938-.25l.187-.188l.188.157a.464.464 0 0 0 .343.156a.535.535 0 0 0 .375-.157a.497.497 0 0 0 0-.718l-.156-.188l.25-.25l.188.157a.464.464 0 0 0 .343.156a.535.535 0 0 0 .375-.157a.497.497 0 0 0 0-.718l-.156-.188l.438-.406a.99.99 0 0 0 0-1.406L24.28.28a.939.939 0 0 0-.686-.28m-.063 1.188c.077 0 .16.004.219.062a.317.317 0 0 1 0 .438l-2 2a.316.316 0 0 1-.219.093a.32.32 0 0 1-.218-.094a.317.317 0 0 1 0-.437l2-2c.058-.058.142-.063.218-.063zm1 1c.077 0 .16.003.219.062a.316.316 0 0 1 0 .438l-2 2a.32.32 0 0 1-.219.093a.32.32 0 0 1-.218-.093a.317.317 0 0 1 0-.438l2-2c.058-.058.142-.063.218-.063zM12.156 7.655c.96 0 1.974.652 2.75 1.313l-3.344 3.344c-.074-.008-.142-.032-.218-.032a2.383 2.383 0 0 0-2.375 2.375a2.383 2.383 0 0 0 2.375 2.375a2.383 2.383 0 0 0 2.375-2.375c0-.076-.024-.144-.031-.219l3.374-3.374c.868 1.006 1.848 2.624.907 4c-.306.446-.691.61-1.375.843c-.602.205-1.355.455-1.969 1.094c-.674.701-.824 1.623-1 2.594c-.222 1.219-.471 2.597-1.781 3.906c-.688.688-1.446 1.031-2.313 1.031c-1.517 0-3.337-1.046-5.125-2.906a.46.46 0 0 0-.031-.031C2.53 19.822 1.484 18.01 1.469 16.5c-.01-.881.334-1.645 1.031-2.344c1.31-1.31 2.688-1.56 3.906-1.781c.971-.176 1.894-.357 2.594-1.031c.634-.608.885-1.343 1.094-1.938c.24-.682.39-1.034.844-1.344c.388-.265.79-.406 1.218-.406zM5.562 16.28a.316.316 0 0 0-.218.094L4.719 17a.313.313 0 0 0 0 .438l3.843 3.843c.118.117.32.117.438 0l.625-.625a.315.315 0 0 0 0-.437l-3.844-3.844a.316.316 0 0 0-.218-.094z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Happy(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 .188C5.924.188.187 5.923.187 13S5.925 25.813 13 25.813c7.076 0 12.813-5.737 12.813-12.813C25.813 5.924 20.075.187 13 .187zm0 2c5.962 0 10.813 4.85 10.813 10.812S18.962 23.813 13 23.813C7.038 23.813 2.187 18.962 2.187 13C2.188 7.038 7.038 2.187 13 2.187zM9 9a2 2 0 1 0 0 4a2 2 0 0 0 0-4m8 0a2 2 0 1 0 0 4a2 2 0 0 0 0-4M5.75 16.094a8.83 8.83 0 0 0 7.25 3.75a8.829 8.829 0 0 0 7.25-3.75A12.374 12.374 0 0 1 13 18.437c-2.707 0-5.208-.874-7.25-2.343"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Headset(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 .031c-5.203 0-9.588 3.606-11 8.531c-1.262.76-2 2.224-2 3.938C0 14.985 1.515 17 4 17a1 1 0 0 0 1-1V9c0-.46-.327-.821-.75-.938C5.71 4.48 9.08 1.97 13 1.97s7.29 2.51 8.75 6.093C21.327 8.18 21 8.54 21 9v7a1 1 0 0 0 1 1c.233 0 .441-.028.656-.063c-.323 1.866-1.107 3.257-1.906 4.22a7.343 7.343 0 0 1-1.406 1.312a3.439 3.439 0 0 1-.469.281c-.058.028-.106.054-.125.063h-1.875C16.515 21.767 15.369 21 14 21c-1.657 0-3 1.12-3 2.5s1.343 2.5 3 2.5c1.369 0 2.515-.768 2.875-1.813h1.969c.338 0 .424-.09.625-.187c.2-.098.437-.231.687-.406c.5-.35 1.105-.86 1.688-1.563c1.054-1.27 2.036-3.155 2.281-5.687C25.305 15.56 26 14.152 26 12.5c0-1.714-.738-3.177-2-3.938c-1.412-4.925-5.797-8.53-11-8.53z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Helicopter(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.813 7.531a.95.95 0 0 0-.75.969v.063L5 8.063A.93.93 0 0 0 4.062 9A.93.93 0 0 0 5 9.938l9.063-.5v2.093A4.41 4.41 0 0 0 12.28 13H2.5c-.5 0-.619.844-.063.938c.153.025 5.561 1.345 8.75 2.124c-.057.447-.093.9-.093 1.375c0 3.525 3.079 3.532 6.875 3.532s6.843-.006 6.843-3.532c0-3.522-4.609-6.375-8.406-6.375c-.162 0-.315.022-.468.032V9.438l9.062.5A.93.93 0 0 0 25.938 9A.93.93 0 0 0 25 8.062l-9.063.5V8.5a.95.95 0 0 0-1.03-.969a.95.95 0 0 0-.095 0zM2.5 11.188a2.312 2.312 0 1 0 0 4.624a2.28 2.28 0 0 0 1.813-.906l-1.22-.344c-.177.1-.376.188-.593.188c-.686 0-1.25-.565-1.25-1.25s.565-1.25 1.25-1.25c.332 0 .622.13.844.344h1.281A2.293 2.293 0 0 0 2.5 11.188m15.625 1.874c1.78 0 4.844 1.663 4.844 3.75c0 .782-.81 1.157-3.563 1.157c-1.806 0-2.312-.84-2.312-1.907c0-1.613.198-3 1.031-3m6.688 7.97a.95.95 0 0 0-.688.593c.01-.01.035-.017-.063.031c-.277.139-1.15.407-3.062.407H10a.95.95 0 1 0 0 1.875h11c2.071 0 3.18-.233 3.906-.594c.363-.181.641-.404.813-.657c.086-.126.152-.257.187-.375c.036-.117.032-.28.032-.28a.95.95 0 0 0-1.032-1a.95.95 0 0 0-.093 0M902 1469v2h26v-2zm4 5v2h18v-2zm-4 5v2h26v-2zm4 5v2h18v-2zm-4 5v2h26v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeraldTrumpet(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.875 0c-.215-.023-.34.285-.188.813c.15.518 1.319 3.775-.437 5.53l-5.125 5.095l-.938-.938l.157-.125a.5.5 0 0 0-.313-.906a.5.5 0 0 0-.125.031a.5.5 0 0 0-.281.156l-1 1a.5.5 0 1 0 .719.688l.125-.125l.969.969l-1.313 1.25l-.938-.938l.157-.125a.5.5 0 0 0-.313-.906a.5.5 0 0 0-.125.031a.5.5 0 0 0-.281.156l-1 1a.5.5 0 1 0 .719.688l.125-.125l.969.969l-1.313 1.25l-.938-.938l.157-.125a.5.5 0 0 0-.313-.906a.5.5 0 0 0-.125.031a.5.5 0 0 0-.281.156l-1 1a.5.5 0 1 0 .719.688l.125-.125l.968.969L4.095 20.5l1.437 1.406L9 18.438c.088.601.353 1.197.844 1.687c.878.878 2.076 1.044 3.094.688c1.017-.357 1.917-1.106 2.812-2l.031-.032c.88-.883 1.617-1.807 1.969-2.812c.357-1.018.191-2.215-.688-3.094a2.931 2.931 0 0 0-1.5-.813c-.051-.01-.103-.021-.156-.03l4.25-4.25c1.681-1.683 4.7-.537 5.625-.376c.857.15.817-.278.563-.531l-6.75-6.75C19.009.04 18.947.008 18.875 0m-3.969 13.438c.132-.023.252-.024.375 0c.248.047.479.197.75.468c.544.544.57.947.344 1.594c-.227.647-.836 1.461-1.656 2.281c-.82.82-1.635 1.43-2.281 1.657c-.647.226-1.02.199-1.563-.344c-.544-.544-.601-.948-.375-1.594c.17-.485.557-1.05 1.094-1.656c.179-.202.388-.42.594-.625c.102-.103.21-.186.312-.281c.139-.14.173-.142.313-.282c.605-.528 1.17-.923 1.656-1.094c.161-.056.305-.102.437-.124m-12 8.28l-.875.876l-.312-.313A1 1 0 0 0 .78 22a1 1 0 0 0-.5 1.719l2 2a1.016 1.016 0 1 0 1.44-1.439L3.438 24l.874-.875l-1.406-1.406z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hexagon(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1a2 2 0 0 0-2 2c0 .183.016.363.063.531L1.968 11a2 2 0 0 0 0 4l4.093 7.469A2.02 2.02 0 0 0 6 23a2 2 0 0 0 2 2a1.98 1.98 0 0 0 1.719-1h6.562c.346.597.98 1 1.719 1a2 2 0 0 0 2-2a2.02 2.02 0 0 0-.063-.531L24.032 15a2 2 0 0 0 0-4l-4.093-7.469A2.02 2.02 0 0 0 20 3a2 2 0 0 0-2-2a1.98 1.98 0 0 0-1.719 1H9.72A1.98 1.98 0 0 0 8 1m1.719 3h6.562c.346.597.98 1 1.719 1c.148 0 .298 0 .438-.031l3.875 7A1.987 1.987 0 0 0 22 13c0 .38.129.729.313 1.031l-3.875 7c-.14-.03-.29-.031-.438-.031a1.98 1.98 0 0 0-1.719 1H9.72A1.98 1.98 0 0 0 8 21c-.148 0-.298 0-.438.031l-3.875-7C3.872 13.73 4 13.38 4 13c0-.38-.129-.729-.313-1.031l3.876-7c.139.03.289.031.437.031a1.98 1.98 0 0 0 1.719-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HighBattery(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 5C3.3 5 2 6.3 2 8v1H1c-.6 0-1 .4-1 1v6c0 .6.4 1 1 1h1v1c0 1.7 1.3 3 3 3h18c1.7 0 3-1.3 3-3V8c0-1.7-1.3-3-3-3zm0 2h1v12H5c-.6 0-1-.4-1-1v-3H2v-4h2V8c0-.6.4-1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageFile(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 0C4.795 0 3 1.795 3 4v18c0 2.205 1.795 4 4 4h12c2.205 0 4-1.795 4-4V8c0-1.062-.973-2.07-2.719-3.781c-.244-.24-.504-.502-.75-.75c-.248-.246-.51-.475-.75-.719C17.07 1.003 16.063 0 15 0zm0 2h7.281c.721.184.719 1.05.719 1.938V7c0 .551.449 1 1 1h3c.998 0 2 .005 2 1v13a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2m3 7a2 2 0 1 0 0 4a2 2 0 0 0 0-4m5.969 3.156c-1.187 0-1.607 4.906-3.25 4.906c-1.217 0-1.69-2-2.719-2S7.125 19.25 7.125 19.25c-.38.422.186.75.625.75h10.563c.357 0 .81-.453.53-1.063c0 0-1.687-6.78-2.874-6.78z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InTransit(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 4v1h8v1H2v1H0v1h7v1H2v1H0v1h6v1H2v1H0v1h5v1H2v4c0 .6.4 1 1 1h1c.2-1.7 1.7-3.094 3.5-3.094S10.8 18.2 11 20h3c.6 0 1-.4 1-1V5c0-.6-.4-1-1-1zm17 4c-.6 0-1 .4-1 1v10c0 .5.4 1 1 1c.2-1.7 1.7-3.094 3.5-3.094S23.8 18.2 24 20h1c.6 0 1-.4 1-1v-5c0-.9-.813-2-.813-2L23 9c-.5-.6-.8-1-1.5-1zm2 2h2.406c.2 0 .407.188.407.188l2.093 2.906c.2.3.407.612.407.812H19zM7.5 18.188A2.321 2.321 0 0 0 5.187 20.5A2.321 2.321 0 0 0 7.5 22.813A2.321 2.321 0 0 0 9.813 20.5A2.321 2.321 0 0 0 7.5 18.187zm13 0a2.321 2.321 0 0 0-2.313 2.312a2.321 2.321 0 0 0 2.313 2.313a2.321 2.321 0 0 0 2.313-2.313a2.321 2.321 0 0 0-2.313-2.313z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inbox(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.406 3c-1.3 0-2.396.87-2.843 2.063v.03l-2.5 7.595l-.063.156V19c0 2.2 1.8 4 4 4h18c2.2 0 4-1.8 4-4v-6.156l-.063-.156l-2.593-7.626C22.922 3.938 21.844 3 20.5 3zm0 2H20.5c.456 0 .79.275.969.75l2.5 7.25H16c0 1.7-1.3 3-3 3s-3-1.3-3-3H2.031l2.406-7.25c.153-.407.67-.75.97-.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inspection(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 0c-.7 0-1.206.294-1.406.594c-.1.1-.181.306-.281.406H10c-1.1 0-2 .9-2 2H6C4.3 3 3 4.3 3 6v17c0 1.7 1.3 3 3 3h14c1.7 0 3-1.3 3-3V6c0-1.7-1.3-3-3-3h-2c0-1.1-.9-2-2-2h-1.313c-.1-.1-.18-.306-.28-.406C14.206.294 13.7 0 13 0m-3 3h6v2h-6zM6 6h2.313c.3.6.987 1 1.687 1h6c.7 0 1.387-.4 1.688-1H20v17H6zm10.563 4.688c-.188 0-.37.068-.47.218l-4.405 4.688L10 14c-.3-.3-.8-.3-1 0l-.688.688c-.3.3-.3.8 0 1l2.782 2.624c.4.4 1.006.275 1.406-.125l5.406-5.78c.2-.2.207-.607-.093-.907l-.72-.594a.747.747 0 0 0-.53-.219z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IntegratedCircuit(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 2a2 2 0 0 0-2 2v18a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm3 1.938c1.132 0 2.063.93 2.063 2.062A2.073 2.073 0 0 1 10 8.063A2.072 2.072 0 0 1 7.937 6c0-1.132.931-2.063 2.063-2.063zM2.437 4A.433.433 0 0 0 2 4.438V5H.437A.433.433 0 0 0 0 5.438v1.125C0 6.809.193 7 .438 7H2v.563c0 .245.192.437.438.437h1.124A.433.433 0 0 0 4 7.562V4.438A.433.433 0 0 0 3.562 4H2.438zm20 0a.433.433 0 0 0-.437.438v3.125c0 .245.192.437.438.437h1.125A.433.433 0 0 0 24 7.562V7h1.563A.433.433 0 0 0 26 6.562V5.438A.433.433 0 0 0 25.562 5H24v-.563A.433.433 0 0 0 23.562 4zm-20 7a.433.433 0 0 0-.437.438V12H.437a.433.433 0 0 0-.437.438v1.124c0 .247.193.438.438.438H2v.563c0 .244.192.437.438.437h1.124A.433.433 0 0 0 4 14.562v-3.124A.433.433 0 0 0 3.562 11H2.438zm20 0a.433.433 0 0 0-.437.438v3.124c0 .245.192.438.438.438h1.125a.433.433 0 0 0 .437-.438V14h1.563a.433.433 0 0 0 .437-.438v-1.124a.433.433 0 0 0-.438-.438H24v-.563a.433.433 0 0 0-.438-.437zm-20 7a.433.433 0 0 0-.437.438V19H.437a.433.433 0 0 0-.437.438v1.125c0 .245.193.437.438.437H2v.563c0 .245.192.437.438.437h1.124A.433.433 0 0 0 4 21.562v-3.125A.433.433 0 0 0 3.562 18H2.438zm20 0a.433.433 0 0 0-.437.438v3.125c0 .245.192.437.438.437h1.125a.433.433 0 0 0 .437-.438V21h1.563a.433.433 0 0 0 .437-.438v-1.125a.433.433 0 0 0-.438-.437H24v-.563a.433.433 0 0 0-.438-.437z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Invisible(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 6.156c-5.02 0-9.712 2.494-12.531 6.657a1.68 1.68 0 0 0 .469 2.343c.292.2.608.313.937.313c.546 0 1.108-.265 1.438-.75c.104-.155.232-.29.343-.438A9.873 9.873 0 0 0 13 20.937a9.872 9.872 0 0 0 9.344-6.656c.11.149.239.283.343.438c.525.778 1.595.963 2.375.437a1.681 1.681 0 0 0 .47-2.344A15.13 15.13 0 0 0 13 6.157zm0 3.625A3.227 3.227 0 0 1 16.219 13A3.227 3.227 0 0 1 13 16.219A3.227 3.227 0 0 1 9.781 13A3.227 3.227 0 0 1 13 9.781m3.969.469c1.67.6 3.157 1.589 4.406 2.875c-.933 3.758-4.332 6.563-8.375 6.563s-7.442-2.805-8.375-6.563a11.69 11.69 0 0 1 4.406-2.844A4.812 4.812 0 0 0 13 17.812A4.812 4.812 0 0 0 17.813 13a4.787 4.787 0 0 0-.844-2.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Iphone(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a3 3 0 0 0-3 3v20a3 3 0 0 0 3 3h11a3 3 0 0 0 3-3V3a3 3 0 0 0-3-3zm4.5 1h2a.5.5 0 0 1 0 1h-2a.5.5 0 0 1 0-1m-5 2h12a.5.5 0 0 1 .5.5v17a.5.5 0 0 1-.5.5h-12a.5.5 0 0 1-.5-.5v-17a.5.5 0 0 1 .5-.5m7.438 5c-.413.027-.875.282-1.157.625c-.256.31-.489.769-.406 1.219c.451.012.918-.246 1.188-.594c.252-.324.444-.782.374-1.25zm-2.782 1.781c-.59 0-1.227.386-1.625 1c-.562.862-.47 2.459.438 3.844c.326.494.776 1.057 1.344 1.063c.503.005.627-.309 1.312-.313c.685-.004.807.318 1.313.313c.566-.006 1.05-.6 1.374-1.094a5.59 5.59 0 0 0 .5-.938c-1.315-.497-1.55-2.382-.25-3.094c-.396-.494-.943-.78-1.468-.78c-.695 0-.986.343-1.469.343c-.498 0-.867-.344-1.469-.344M13.5 22.22c.706 0 1.281.575 1.281 1.281s-.575 1.281-1.281 1.281a1.283 1.283 0 0 1-1.281-1.281c0-.706.575-1.281 1.281-1.281"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeepDry(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 0s-1 1.26-1 2a1 1 0 0 0 2 0c0-.74-1-2-1-2M5 1S4 2.26 4 3a1 1 0 0 0 2 0c0-.74-1-2-1-2m16 0s-1 1.26-1 2a1 1 0 0 0 2 0c0-.74-1-2-1-2M9 2S8 3.26 8 4a1 1 0 0 0 2 0c0-.74-1-2-1-2m8 0s-1 1.26-1 2a1 1 0 0 0 2 0c0-.74-1-2-1-2m-4.188 2.188a.8.8 0 0 0-.624.812v1.031C2.68 6.493 2 15 2 15s1.406-1.125 3.5-1.125C7.594 13.875 9 15 9 15s1.409-.841 3.188-1.063V23c0 .673-.515 1.188-1.188 1.188A1.166 1.166 0 0 1 9.812 23a.813.813 0 0 0-1.624 0A2.834 2.834 0 0 0 11 25.813A2.834 2.834 0 0 0 13.813 23v-9.063A9.125 9.125 0 0 1 17 15s1.368-1.125 3.5-1.125C22.632 13.875 24 15 24 15s-.68-8.507-10.188-8.969V5a.8.8 0 0 0-.906-.813a.8.8 0 0 0-.094 0zM2 5S1 6.26 1 7a1 1 0 0 0 2 0c0-.74-1-2-1-2m22 0s-1 1.26-1 2a1 1 0 0 0 2 0c0-.74-1-2-1-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeySecurity(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 .188a7.813 7.813 0 1 0 0 15.625A7.813 7.813 0 0 0 8 .187zM5.5 2.905A2.589 2.589 0 0 1 8.094 5.5A2.587 2.587 0 0 1 5.5 8.094A2.587 2.587 0 0 1 2.906 5.5A2.589 2.589 0 0 1 5.5 2.906zm11.094 8.719a9.232 9.232 0 0 1-1.032 1.813l7.813 7.812a.89.89 0 0 1 0 1.25a.892.892 0 0 1-1.25 0l-7.719-7.75A9.396 9.396 0 0 1 11 16.813v1.375c0 .44.371.812.813.812H14v2.188c0 .44.371.812.813.812H17v2.188c0 .44.372.812.813.812h2.218l.563.563c.342.341 3.768.512 4.625-.344c.857-.858.684-4.284.343-4.625z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyboard(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23 6H3a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h20a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3m-3 5h-2V9h2zm3 0h-2V9h2zm-9 0h-2V9h2zm3 0h-2V9h2zm-6 0H9V9h2zm-6 0H3V9h2zm3 0H6V9h2zm12 8H6v-2h14zm1-4h-2v-2h2zm-6 0h-2v-2h2zm3 0h-2v-2h2zm-6 0h-2v-2h2zm-6 0H4v-2h2zm3 0H7v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Last(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.125 0H4.875A4.874 4.874 0 0 0 0 4.875v16.25A4.874 4.874 0 0 0 4.875 26h16.25A4.874 4.874 0 0 0 26 21.125V4.875A4.874 4.874 0 0 0 21.125 0M19 18c0 .551-.449 1-1 1h-2c-.551 0-1-.449-1-1v-3.618l-5.569 4.425a.984.984 0 0 1-.951-.04a.969.969 0 0 1-.48-.817V8.051c0-.334.19-.644.479-.817a.986.986 0 0 1 .952-.039L15 11.618V8c0-.551.449-1 1-1h2c.551 0 1 .449 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Like(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.869 3.889c-2.096 0-3.887 1.494-4.871 2.524c-.984-1.03-2.771-2.524-4.866-2.524C4.521 3.889 2 6.406 2 10.009c0 3.97 3.131 6.536 6.16 9.018c1.43 1.173 2.91 2.385 4.045 3.729c.191.225.471.355.765.355h.058c.295 0 .574-.131.764-.355c1.137-1.344 2.616-2.557 4.047-3.729C20.867 16.546 24 13.98 24 10.009c0-3.603-2.521-6.12-6.131-6.12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lol(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 .188C5.924.188.187 5.923.187 13S5.925 25.813 13 25.813c7.076 0 12.813-5.737 12.813-12.813C25.813 5.924 20.075.187 13 .187zm0 2c5.962 0 10.813 4.85 10.813 10.812S18.962 23.813 13 23.813C7.038 23.813 2.187 18.962 2.187 13C2.188 7.038 7.038 2.187 13 2.187zM9 9c-1.104 0-3 1.375-3 2s1.896 0 3 0s2 .703 2 0s-.896-2-2-2m8 0c-1.104 0-2 1.297-2 2c0 .703.896 0 2 0s3 .625 3 0s-1.896-2-3-2M6.156 15c0 3.227 2.305 6.844 6.844 6.844s6.844-3.617 6.844-6.844c0 0-2.203 1.156-6.844 1.156S6.156 15 6.156 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LuggageTrolley(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 3c-1.093 0-2 .907-2 2v1H8a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h17a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1h-5V5c0-1.093-.907-2-2-2zM.812 4A1.001 1.001 0 0 0 1 6h1c.505 0 .66.108.781.25c.122.142.219.413.219.75v11c0 1.124.248 2.182.969 2.938C4.689 21.692 5.769 22 7 22h18a1 1 0 1 0 0-2H7c-.905 0-1.318-.212-1.563-.469C5.194 19.275 5 18.83 5 18V7c0-.69-.173-1.43-.688-2.031C3.798 4.368 2.95 4 2 4H1a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.094 0M23 22a2 2 0 1 0 0 4a2 2 0 0 0 0-4M7 22a2 2 0 1 0 0 4a2 2 0 0 0 0-4m8-17h3v1h-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MacOs(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.934 18.947c-.598 1.324-.884 1.916-1.652 3.086c-1.073 1.634-2.588 3.673-4.461 3.687c-1.666.014-2.096-1.087-4.357-1.069c-2.261.011-2.732 1.089-4.4 1.072c-1.873-.017-3.307-1.854-4.381-3.485c-3.003-4.575-3.32-9.937-1.464-12.79C4.532 7.425 6.61 6.237 8.561 6.237c1.987 0 3.236 1.092 4.879 1.092c1.594 0 2.565-1.095 4.863-1.095c1.738 0 3.576.947 4.889 2.581c-4.296 2.354-3.598 8.49.742 10.132M16.559 4.408c.836-1.073 1.47-2.587 1.24-4.131c-1.364.093-2.959.964-3.891 2.092c-.844 1.027-1.544 2.553-1.271 4.029c1.488.048 3.028-.839 3.922-1.99"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Maestro(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 5.188A7.812 7.812 0 0 0 .187 13A7.812 7.812 0 0 0 13 19a7.812 7.812 0 1 0 0-12a7.772 7.772 0 0 0-5-1.812m10 2A5.818 5.818 0 0 1 23.813 13A5.818 5.818 0 0 1 18 18.813c-1.378 0-2.628-.51-3.625-1.313a7.79 7.79 0 0 0 1.438-4.5a7.79 7.79 0 0 0-1.438-4.5c.997-.804 2.247-1.313 3.625-1.313z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Maintenance(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.313 0L0 1.313l2.313 4l1.5-.22l9.156 9.157l-.781.75c-.4.4-.4 1.006 0 1.406l.406.407c.4.4 1.012.4 1.312 0L15.094 18c-.1.6 0 1.313.5 1.813L21 25.188c1.1 1.1 2.9 1.1 4 0c1.3-1.2 1.288-2.994.188-4.094l-5.375-5.407c-.5-.5-1.213-.7-1.813-.5L16.687 14c.3-.4.3-1.012 0-1.313l-.375-.374a.974.974 0 0 0-1.406 0l-.656.656l-9.156-9.156l.218-1.5l-4-2.313zm19.5.031C18.84-.133 16.224 1.175 15 2.312c-1.506 1.506-1.26 3.475-.063 5.376l-2.124 2.125l1.5 1.687c.8-.7 1.98-.7 2.78 0l.407.406l.094.094l.875-.875c1.808 1.063 3.69 1.216 5.125-.219c1.4-1.3 2.918-4.506 2.218-6.406L23 7.406c-.4.4-1.006.4-1.406 0L18.687 4.5a.974.974 0 0 1 0-1.406L21.595.188c-.25-.088-.5-.133-.782-.157m-11 12.469l-3.626 3.625A5.26 5.26 0 0 0 5 16c-2.8 0-5 2.2-5 5s2.2 5 5 5s5-2.2 5-5c0-.513-.081-1.006-.219-1.469l2.125-2.125l-.312-.406c-.8-.8-.794-2.012-.094-2.813L9.812 12.5zm7.75 4.563c.125 0 .243.024.343.125l5.907 5.906c.2.2.2.518 0 .718c-.2.2-.52.2-.72 0l-5.905-5.906c-.2-.2-.2-.518 0-.718c.1-.1.25-.125.375-.125M5.688 18.405l1.906 1.907l-.688 2.593l-2.593.688l-1.907-1.907l.688-2.593z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MedicalDoctor(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 0c-2.1 0-3.357.32-4.156.719c-.4.2-.684.41-.875.625a2 2 0 0 0-.313.531a1 1 0 0 0-.062.25l-.532 6.844c-.007.042.007.11 0 .156L7 9.813A1.003 1.003 0 0 0 7 10c0 3.3 2.7 6 6 6s6-2.7 6-6v-.156a1.003 1.003 0 0 0 0-.031l-.594-7.688a1 1 0 0 0-.062-.25s-.121-.316-.313-.531c-.191-.216-.475-.426-.875-.625C16.357.319 15.1 0 13 0m0 16c-6.6 0-12 3.106-12 5.906V26h24v-4.094c0-2.66-4.882-5.59-11.031-5.875A14.294 14.294 0 0 0 13 16m0-14c1.9 0 2.849.3 3.25.5c.134.067.15.093.188.125l.406 5.125C15.924 7.806 14.67 8 13 8s-2.923-.194-3.844-.25l.406-5.125c.037-.032.054-.058.188-.125c.401-.2 1.35-.5 3.25-.5m-.813 1v1.188H11v1.625h1.188V7h1.624V5.812H15V4.188h-1.188V3zM10 18.25L12.563 24H3v-2.094c0-.745 2.55-2.927 7-3.656m6 0c4.45.73 7 2.911 7 3.656V24h-9.563z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Message(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23 4H3C1.3 4 0 5.3 0 7v12c0 1.7 1.3 3 3 3h20c1.7 0 3-1.3 3-3V7c0-1.7-1.3-3-3-3m.8 15.4L16 13.8l-3 2l-3.1-2l-7.7 5.6l6.3-6.5l-7.7-6L13 13.5L25.1 7l-7.6 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageOutline(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 4C1.344 4 0 5.344 0 7v12c0 1.656 1.344 3 3 3h20c1.656 0 3-1.344 3-3V7c0-1.656-1.344-3-3-3zm0 2h20c.551 0 1 .449 1 1v.5l-11 5.938L2 7.5V7c0-.551.449-1 1-1M2 7.781l6.531 5.094l-6.406 6.563l7.813-5.563L13 15.844l3.063-1.969l7.812 5.563l-6.406-6.563L24 7.781V19a.95.95 0 0 1-.125.438c-.165.325-.486.562-.875.562H3c-.389 0-.71-.237-.875-.563A.95.95 0 0 1 2 19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microphone(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 0a5 5 0 0 0-5 5v8a5 5 0 0 0 5 5c.173 0 .332-.014.5-.031a5.049 5.049 0 0 0 3.656-2.188c.54-.798.844-1.745.844-2.781V5a5 5 0 0 0-5-5m0 2c1.3 0 2.396.842 2.813 2H15c-.551 0-1 .448-1 1c0 .965 1 1 1 1h1v1h-1c-.551 0-1 .448-1 1c0 .965 1 1 1 1h1v1h-6V9h1s1-.035 1-1c0-.552-.449-1-1-1h-1V6h1s1-.035 1-1c0-.552-.449-1-1-1h-.813c.417-1.158 1.514-2 2.813-2m-8 9c-.891 0-1 1-1 1v1c0 4.275 2.998 7.837 7 8.75V24H8c-.959 0-1 1-1 1v1h12v-1c0-.551-.448-1-1-1h-3v-2.25c4.002-.913 7-4.475 7-8.75v-1s-.094-1-1-1s-1 1-1 1v1c0 3.859-3.141 7-7 7s-7-3.141-7-7v-1s-.109-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mouse(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.625 3.25c-.6.033-1.192.186-1.719.438c-1.405.67-2.576 1.857-3.468 2.75a.8.8 0 1 0 1.124 1.125c.88-.88 1.99-1.91 3.032-2.407c.521-.248.992-.358 1.437-.312c.445.046.903.215 1.407.718c.849.85.897 1.566.53 2.563c-.365.997-1.277 2.185-2.374 3.281c-1.05 1.05-1.834 2.028-2.188 3.063a2.884 2.884 0 0 0 .688 3c.822.822 2.011 1.01 3.093.593c1.083-.416 2.163-1.287 3.375-2.5a.8.8 0 1 0-1.125-1.124c-1.139 1.139-2.12 1.858-2.812 2.124c-.692.267-.967.221-1.406-.218c-.45-.45-.494-.752-.282-1.375c.213-.623.818-1.474 1.782-2.438c1.175-1.174 2.239-2.45 2.75-3.844c.51-1.392.328-3.015-.907-4.25c-.714-.715-1.524-1.103-2.343-1.187a3.678 3.678 0 0 0-.594 0m-10.688 3A6.472 6.472 0 0 0 5.72 8.156L4.656 9.22l4.032 4.031l1.093-1.094a1.714 1.714 0 0 1 .313-1.969l1.094-1.093a1.713 1.713 0 0 1 1.968-.313l1.156-1.156A6.56 6.56 0 0 0 9.939 6.25zM15.5 8.781L14.281 10a1.696 1.696 0 0 1-.375 1.813l-1.094 1.093a1.696 1.696 0 0 1-1.812.375l-1.156 1.156l4.062 4.032l1.063-1.063c2.356-2.356 2.522-6.059.531-8.625m-3.094.813a.725.725 0 0 0-.5.219l-1.094 1.093a.715.715 0 0 0 0 1l.282.281a.719.719 0 0 0 .5.22a.723.723 0 0 0 .5-.22l1.094-1.093a.72.72 0 0 0 .218-.5a.723.723 0 0 0-.219-.5l-.28-.281a.713.713 0 0 0-.5-.22zm-8.719.594l-1.5 1.468c-2.555 2.556-2.555 6.725 0 9.281a6.54 6.54 0 0 0 9.25 0l1.5-1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicalNotes(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.563.17l-11.11 3.399c-1.346.384-2.437 1.788-2.437 3.134v11.719s-.805-.543-2.598-.289C2.785 18.507.65 20.528.65 22.648s2.135 3.419 4.768 3.045c2.635-.372 4.566-2.331 4.566-4.452V11.235c0-.94 1.13-1.343 1.13-1.343l9.823-3.079s1.087-.365 1.087.641v8.037s-1.001-.576-2.794-.358c-2.633.319-4.768 2.298-4.768 4.417c0 2.121 2.135 3.463 4.768 3.143c2.635-.319 4.77-2.297 4.77-4.418V1.912C24 .566 22.908-.214 21.563.17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mute(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 2A1 1 0 0 0 3 3.719l20 20a1 1 0 1 0 1.406-1.407L17 14.907V3.312c0-1.265-1.105-1.582-1.969-.718L9.812 7.719L4.407 2.312A1 1 0 0 0 3.594 2A1 1 0 0 0 3.5 2M5 9.063c-.551 0-1 .448-1 1v6c0 .55.449 1 1 1h3.438L15 23.468c1 1 2 .488 2-.875V20.03L6.031 9.063z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MyTopic(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 .188C6 .188.187 5.252.187 11.624c0 2.998 1.407 5.633 3.5 7.656c-.081 1.159-.663 1.852-1.468 2.532c-.433.364-.912.705-1.344 1.125c-.432.42-.875.99-.875 1.78a1 1 0 0 0 .75.97c1.878.503 3.9.132 5.75-.532c1.666-.598 3.072-1.483 4.188-2.343c.754.124 1.514.218 2.312.218c7 0 12.813-5.034 12.813-11.406C25.813 5.253 20 .187 13 .187zm0 2c6.046 0 10.813 4.29 10.813 9.437c0 5.148-4.767 9.406-10.813 9.406c-.785 0-1.537-.047-2.281-.187a1 1 0 0 0-.844.187c-.903.759-2.418 1.671-4.031 2.25c-.994.357-2 .52-2.938.563c.21-.17.377-.29.625-.5c1.007-.85 2.219-2.216 2.219-4.282a1 1 0 0 0-.344-.75c-2-1.726-3.218-4.078-3.218-6.687c0-5.149 4.766-9.438 10.812-9.438zm0 2.874c-1.465 0-2.719 1.01-2.719 2.907c0 1.237.584 2.475 1.344 3.25c.296.777-.232 1.084-.344 1.125c-1.533.554-3.344 1.563-3.344 2.562v.375c0 1.36 2.623 1.656 5.063 1.656c2.444 0 5.063-.296 5.063-1.656v-.375c0-1.027-1.797-2.03-3.407-2.562c-.074-.024-.539-.244-.25-1.125c.755-.778 1.313-2.017 1.313-3.25c0-1.896-1.254-2.907-2.719-2.907"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Name(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.563 15.9c-.159-.052-1.164-.505-.536-2.414h-.009c1.637-1.686 2.888-4.399 2.888-7.07c0-4.107-2.731-6.26-5.905-6.26c-3.176 0-5.892 2.152-5.892 6.26c0 2.682 1.244 5.406 2.891 7.088c.642 1.684-.506 2.309-.746 2.397c-3.324 1.202-7.224 3.393-7.224 5.556v.811c0 2.947 5.714 3.617 11.002 3.617c5.296 0 10.938-.67 10.938-3.617v-.811c0-2.228-3.919-4.402-7.407-5.557"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func News(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.906 1.969a1 1 0 0 0-.125.031H1a1 1 0 0 0-.094 0A1 1 0 0 0 0 3v15.969c0 2.365 1.319 3.818 2.563 4.437C3.806 24.025 5 24 5 24h16.094a1 1 0 0 0 .031 0c.175 0 1.176-.026 2.313-.594c1.186-.593 2.427-1.959 2.53-4.125a1 1 0 0 0 0-.062c.002-.022.031-.04.032-.063A1 1 0 0 0 26 19a1 1 0 0 0 0-.094V8.031A1 1 0 0 0 26 8s-.004-.702-.375-1.438c-.337-.667-1.147-1.4-2.281-1.53A1 1 0 0 0 23 4.968h-5V3.187A1 1 0 0 0 18 3a1 1 0 0 0 0-.125a1 1 0 0 0 0-.063a1 1 0 0 0 0-.03a1 1 0 0 0-.031-.063a1 1 0 0 0 0-.031a1 1 0 0 0-.032-.094a1 1 0 0 0-.062-.094a1 1 0 0 0-.094-.156a1 1 0 0 0-.093-.063a1 1 0 0 0-.125-.125a1 1 0 0 0-.032 0A1 1 0 0 0 17.188 2a1 1 0 0 0-.282-.031M2 4h14v1.75a1.001 1.001 0 0 0 0 .406V18.97c0 1.28.406 2.275.969 3.031H5s-.454.002-1-.156a3.389 3.389 0 0 1-.563-.219C2.681 21.249 2 20.68 2 18.969zm2 2c-.551 0-1 .449-1 1v2c0 .551.449 1 1 1h10c.551 0 1-.449 1-1V7c0-.551-.449-1-1-1zm14 .969h2.188c-.187.548-.188 1-.188 1A1 1 0 0 0 20 8v10a1 1 0 1 0 2 0V8c.002-.028.025-.275.156-.531c.139-.27.209-.469.844-.469c.63 0 .703.19.844.469c.133.265.154.534.156.562V19a1 1 0 0 0 0 .094c-.027 1.588-.692 2.127-1.438 2.5C21.802 21.974 21 22 21 22s-.8-.021-1.563-.406c-.762-.385-1.437-.951-1.437-2.625zM3.719 12A1.004 1.004 0 0 0 4 14h3a1 1 0 1 0 0-2H4a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.094 0a1.004 1.004 0 0 0-.093 0m6 0A1.004 1.004 0 0 0 10 14h4a1 1 0 1 0 0-2h-4a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.094 0a1.004 1.004 0 0 0-.093 0m-6 3A1.004 1.004 0 0 0 4 17h3a1 1 0 1 0 0-2H4a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.094 0a1.004 1.004 0 0 0-.093 0m6 0A1.004 1.004 0 0 0 10 17h4a1 1 0 1 0 0-2h-4a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.094 0a1.004 1.004 0 0 0-.093 0m-5.5 3a1.004 1.004 0 0 0 .281 2H7a1 1 0 1 0 0-2H4.5a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.093 0a1.004 1.004 0 0 0-.094 0m5.5 0A1.004 1.004 0 0 0 10 20h4a1 1 0 1 0 0-2h-4a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.094 0a1.004 1.004 0 0 0-.093 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Next(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.125 0H4.875A4.874 4.874 0 0 0 0 4.875v16.25A4.874 4.874 0 0 0 4.875 26h16.25A4.874 4.874 0 0 0 26 21.125V4.875A4.874 4.874 0 0 0 21.125 0M17.66 13.857l-6.229 4.949a.984.984 0 0 1-.951-.04a.966.966 0 0 1-.48-.816V8.051c0-.334.19-.644.479-.817a.986.986 0 0 1 .952-.039l6.229 4.948c.336.297.537.494.537.857s-.236.575-.537.857"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NfcCheckpoint(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 0a11.503 11.503 0 0 0-8.125 3.344A1.005 1.005 0 1 0 6.313 4.75A9.41 9.41 0 0 1 13 2c2.618 0 4.97 1.04 6.688 2.75a1.005 1.005 0 1 0 1.437-1.406A11.51 11.51 0 0 0 13 0m0 4a6.472 6.472 0 0 0-4.813 2.125a1 1 0 1 0 1.47 1.344A4.482 4.482 0 0 1 13 6c1.327 0 2.518.558 3.344 1.469a1 1 0 1 0 1.468-1.344A6.478 6.478 0 0 0 13 4M5 7v16a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V7l-2 2v11.5a.5.5 0 0 1-.5.5h-11a.5.5 0 0 1-.5-.5V9zm8 1c-1.108 0-2 .893-2 2s.892 2 2 2a2 2 0 1 0 0-4m.5 14.219c.706 0 1.281.575 1.281 1.281s-.575 1.281-1.281 1.281a1.283 1.283 0 0 1-1.281-1.281c0-.706.575-1.281 1.281-1.281"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Note(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 0C1.795 0 0 1.795 0 4v18c0 2.205 1.795 4 4 4h13c1.063 0 2.539-1.535 4.25-3.281c.24-.244.47-.473.719-.719c.246-.248.506-.51.75-.75C24.466 19.538 26 18.063 26 17V4c0-2.205-1.795-4-4-4zm0 2h18a2 2 0 0 1 2 2v13c0 .995-1.002 1-2 1h-3c-.551 0-1 .449-1 1v3.063c0 .887.002 1.753-.719 1.937H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2m2.813 6A1.001 1.001 0 0 0 7 10h12a1 1 0 1 0 0-2H7a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.093 0m0 5A1.001 1.001 0 0 0 7 15h10a1 1 0 1 0 0-2H7a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.093 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Online(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.281 4.063a1.5 1.5 0 0 0-.906 2.593A8.941 8.941 0 0 1 22 13a8.953 8.953 0 0 1-2.625 6.344a1.503 1.503 0 1 0 2.125 2.125a11.995 11.995 0 0 0 0-16.938a1.5 1.5 0 0 0-1.063-.469a1.5 1.5 0 0 0-.156 0zm-14.906.03a1.5 1.5 0 0 0-.875.438a11.995 11.995 0 0 0 0 16.938a1.503 1.503 0 1 0 2.125-2.125A8.941 8.941 0 0 1 4 13a8.953 8.953 0 0 1 2.625-6.344a1.5 1.5 0 0 0-1.25-2.562zm3.813 3.313a1.5 1.5 0 0 0-.876.407A7.014 7.014 0 0 0 6 13c0 2.048.87 3.91 2.281 5.188a1.504 1.504 0 1 0 2.031-2.22A3.975 3.975 0 0 1 9 13a3.98 3.98 0 0 1 1.313-2.969a1.5 1.5 0 0 0-1.126-2.625zm7.406 0a1.5 1.5 0 0 0-.907 2.625A3.975 3.975 0 0 1 17 13a3.98 3.98 0 0 1-1.313 2.969a1.5 1.5 0 1 0 2 2.219A7.014 7.014 0 0 0 20 13a6.98 6.98 0 0 0-2.281-5.188a1.5 1.5 0 0 0-1.125-.406M13 11.188A1.812 1.812 0 1 0 14.813 13A1.81 1.81 0 0 0 13 11.187z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Packaging(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.656.656l-1.312.688l1 2l-2-1l-.688 1.312L11.344 5H2a1 1 0 0 0-.094 0a1 1 0 0 0-.093 0A1 1 0 0 0 1 6v4a1 1 0 0 0 1 1v13a1 1 0 0 0 1 1h20a1 1 0 0 0 1-1V11a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1h-9.344l2.688-1.344l-.688-1.312l-2 1l1-2l-1.312-.688L13 3.344zM3 7h9v2H3.187A1 1 0 0 0 3 8.969zm11 0h9v2h-9zM4 11h8v12H4zm10 0h8v12h-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paid(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.875 2a1 1 0 0 0-.594.313L5.687 9H1c-.6 0-1 .4-1 1v2c0 .6.4 1 1 1h.094l2.718 9.313c.1.4.5.687 1 .687h16.5c.4 0 .8-.288 1-.688L25 13c.6 0 1-.4 1-1v-2c0-.6-.4-1-1-1h-4.594L13.72 2.281A1 1 0 0 0 12.875 2M13 4.438L17.594 9H8.5zm4.563 7.343c.2 0 .38.069.53.219l.72.688c.3.3.3.824 0 1.124l-6 6c-.3.3-.825.3-1.126 0l-3.374-3.406c-.3-.3-.3-.793 0-1.093l.687-.72c.3-.3.794-.3 1.094 0l2.219 2.22L17 12a.798.798 0 0 1 .563-.219"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Panorama(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.875 3A1 1 0 0 0 0 4v18a1 1 0 0 0 1.375.938S6.245 21 13 21s11.625 1.938 11.625 1.938A1 1 0 0 0 26 22V4a1 1 0 0 0-1.438-.906S20.826 5 13 5C5.174 5 1.437 3.094 1.437 3.094A1 1 0 0 0 .875 3M2 5.375c.785.301 2.063.719 4 1.063v13.187c-1.869.35-3.222.767-4 1.031zm22 0v15.281c-.782-.265-2.13-.68-4-1.031V6.437c1.937-.343 3.216-.761 4-1.062M8 6.719C9.407 6.878 11.03 7 13 7c1.97 0 3.593-.122 5-.281v12.593A39.245 39.245 0 0 0 13 19c-1.898 0-3.55.128-5 .313z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperClamp(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 .063C10.182.063 8.062 1.321 8.062 3c0 .761.408 1.447 1.188 1.969L7.781 12h1.907l1.562-7.344a.913.913 0 0 0-.531-1.031c-.576-.261-.781-.541-.781-.625c0-.271 1.056-1.063 3.062-1.063S16.063 2.73 16.063 3c0 .084-.205.364-.782.625a.914.914 0 0 0-.531 1.031L16.313 12h1.906L16.75 4.969c.78-.522 1.188-1.209 1.188-1.969c0-1.678-2.12-2.938-4.938-2.938zM18.219 12l2.531 12h-1.938l-2.5-12H9.688l-2.5 12H5.25l2.531-12H.75c-.562 0-.875.418-.688.938L3.22 24H3a1 1 0 0 0-.094 0A1.001 1.001 0 0 0 3 26h20a1 1 0 1 0 0-2h-.219l3.157-11.063c.187-.518-.126-.937-.688-.937z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperPlane(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.906 0a.99.99 0 0 0-.375.125l-24 13a.99.99 0 0 0 .188 1.813l6.375 1.906c.149 1.179.813 6.285.937 7.281c.124.992.798 1.164 1.469.25c.454-.619 3.124-4.375 3.125-4.375l5.688 5.688a.99.99 0 0 0 1.656-.438l6-24A.99.99 0 0 0 24.906 0M23.47 2.938l-5.032 20.125l-5.656-5.657L21 6L8.219 15.125L3.563 13.75L23.468 2.937z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PartlyCloudyDay(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.875.375a.6.6 0 0 0-.063.031a.6.6 0 0 0-.406.594v2.656a.6.6 0 0 0 .094.375a4.888 4.888 0 0 0-1.531.438a.6.6 0 0 0-.125-.375L5.53 1.78a.6.6 0 0 0-.03-.06a.6.6 0 0 0-.031-.032a.6.6 0 0 0-.032-.062a.6.6 0 0 0-.062-.031a.6.6 0 0 0-.188-.125a.6.6 0 0 0-.25-.032a.6.6 0 0 0-.093.032a.6.6 0 0 0-.375.906l1.343 2.313a.6.6 0 0 0 .22.28c-.402.297-.734.695-1.032 1.095a.6.6 0 0 0-.313-.25L2.376 4.468a.6.6 0 0 0-.094-.063a.6.6 0 0 0-.25-.031a.6.6 0 0 0-.125.031a.6.6 0 0 0-.125 1.125l2.313 1.313a.6.6 0 0 0 .343.125c-.213.468-.316.969-.375 1.5a.6.6 0 0 0-.093-.031a.6.6 0 0 0-.313-.032H1a.6.6 0 0 0-.125 0a.6.6 0 0 0-.063 0A.601.601 0 0 0 1 9.594h2.656a.6.6 0 0 0 .25 0a.6.6 0 0 0 .125-.094c.054.542.192 1.053.407 1.531a.6.6 0 0 0-.063 0a.6.6 0 0 0-.094.031a.6.6 0 0 0-.187.094L1.78 12.47a.609.609 0 1 0 .594 1.062l2.313-1.344a.6.6 0 0 0 .406-.53l1.437-.97A2.987 2.987 0 0 1 9 6a3 3 0 0 1 2.625 1.532l1.438-.812a.6.6 0 0 0 .687.218a.6.6 0 0 0 .156-.093L16.22 5.53a.604.604 0 0 0 .156-1.062a.604.604 0 0 0-.156-.094a.604.604 0 0 0-.344 0a.604.604 0 0 0-.094.031a.6.6 0 0 0-.156.063l-2.313 1.343a.6.6 0 0 0-.28.22a4.918 4.918 0 0 0-1.063-1.063a.6.6 0 0 0 .125-.125a.6.6 0 0 0 .031-.063a.6.6 0 0 0 .031-.031a.6.6 0 0 0 .031-.063l1.344-2.312a.6.6 0 0 0 .031-.625a.6.6 0 0 0-.312-.281a.6.6 0 0 0-.25-.032a.6.6 0 0 0-.125.032a.6.6 0 0 0-.406.312l-1.313 2.313a.6.6 0 0 0-.031.062a.6.6 0 0 0-.094.313A5.133 5.133 0 0 0 9.5 4.062c.01.001.022-.032.031-.03a.6.6 0 0 0 .031-.126a.6.6 0 0 0 .032-.156a.6.6 0 0 0 0-.094V1a.6.6 0 0 0-.031-.188A.6.6 0 0 0 9.53.72a.6.6 0 0 0-.06-.126a.6.6 0 0 0-.031-.031a.6.6 0 0 0-.25-.157a.6.6 0 0 0-.313-.031zM17 7c-2.262 0-4.17 1.3-5.188 3.156C11.54 10.105 11.299 10 11 10c-2.407 0-4.345 1.757-4.813 4.031C6.12 14.027 6.07 14 6 14c-2.197 0-4 1.803-4 4s1.803 4 4 4h15c2.75 0 5-2.25 5-5c0-2.06-1.305-3.77-3.094-4.531C22.626 9.423 20.116 7 17 7m0 2c2.222 0 4 1.778 4 4c0-.137.005-.095 0 .125l-.031.781l.781.188A2.985 2.985 0 0 1 24 17c0 1.668-1.332 3-3 3H6c-1.115 0-2-.885-2-2a1.987 1.987 0 0 1 2.688-1.875l1.437.531l-.094-1.531C8.027 15.061 8.005 15.03 8 15c0-.026-.002-.023 0-.031v-.031A2.987 2.987 0 0 1 11 12c.334 0 .662.068 1 .188l.938.312l.343-.906C13.851 10.078 15.28 9 17 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PartlyCloudyNight(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.688.188c1.46.946 2.406 2.885 2.406 5.218c0 2.144-2.759 4.688-4.688 4.688c-2.333 0-3.942-.566-5.218-2.406c.43 2.397 2.206 4.307 4.5 5c.097-.024.183-.076.28-.094A6.493 6.493 0 0 1 11 8.5h.031a7.606 7.606 0 0 1 1.938-1.813c.002-.058 0-.097 0-.156A6.428 6.428 0 0 0 7.688.188M20.468.5l-.5 1.813l-1.812-.47l1.313 1.345l-1.313 1.343l1.813-.468l.5 1.812L21 4.062l1.813.47L21.5 3.186l1.313-1.343L21 2.312L20.469.5zM17 7c-2.262 0-4.17 1.3-5.188 3.156C11.54 10.105 11.299 10 11 10c-2.407 0-4.345 1.757-4.813 4.031C6.12 14.027 6.07 14 6 14c-2.197 0-4 1.803-4 4s1.803 4 4 4h15c2.75 0 5-2.25 5-5c0-2.06-1.305-3.77-3.094-4.531C22.626 9.423 20.116 7 17 7m0 2c2.222 0 4 1.778 4 4c0-.137.005-.095 0 .125l-.031.781l.781.188A2.985 2.985 0 0 1 24 17c0 1.668-1.332 3-3 3H6c-1.115 0-2-.885-2-2a1.987 1.987 0 0 1 2.688-1.875l1.437.531l-.094-1.531C8.027 15.061 8.005 15.03 8 15c0-.026-.002-.023 0-.031v-.031A2.987 2.987 0 0 1 11 12c.334 0 .662.068 1 .188l.938.312l.343-.906C13.851 10.078 15.28 9 17 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PasswordOne(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.625.188c-1.678 0-3.242.68-4.531 1.968C11.63 3.62 10.96 5.42 11.156 7.344c.147 1.428.799 2.868 1.844 4.156l-7.25 7.25l-1.313-1.313c-.64-.643-1.171-.11-1.812.532L2 18.625c-.642.639-1.172 1.14-.531 1.781L2.78 21.72l-1.03 1.03a1.63 1.63 0 0 0 0 2.313c.64.639 1.671.639 2.313 0l11.375-11.375c1.232.772 2.57 1.187 3.874 1.187c1.677 0 3.244-.68 4.532-1.969c1.465-1.463 2.134-3.263 1.937-5.187c-.176-1.731-1.095-3.473-2.562-4.938C21.557 1.117 19.572.187 17.625.187zm.063 2.062c1.393 0 2.84.713 4.093 1.969c1.114 1.11 1.809 2.402 1.938 3.656c.133 1.304-.326 2.546-1.344 3.563c-.912.91-1.975 1.374-3.125 1.374c-1.394 0-2.838-.714-4.094-1.968c-1.113-1.113-1.808-2.402-1.937-3.656c-.132-1.304.325-2.545 1.344-3.563c.912-.91 1.975-1.375 3.124-1.375z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Past(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 0L8 3l5 3V4c4.955 0 9 4.045 9 9s-4.045 9-9 9s-9-4.045-9-9c0-2.453.883-4.57 2.5-6.188L5.094 5.407C3.11 7.39 2 10.053 2 13c0 6.045 4.955 11 11 11s11-4.955 11-11S19.045 2 13 2zm-2.094 6.563l-1.812.875l2.531 5A1.539 1.539 0 0 0 11.5 13v.063L8.281 16.28l1.439 1.44l3.219-3.219H13a1.5 1.5 0 0 0 1.5-1.5c0-.69-.459-1.263-1.094-1.438z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 5c-.551 0-1 .449-1 1v14c0 .551.449 1 1 1h3c.551 0 1-.449 1-1V6c0-.551-.449-1-1-1zm9 0c-.551 0-1 .449-1 1v14c0 .551.449 1 1 1h3c.551 0 1-.449 1-1V6c0-.551-.449-1-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.386 18.026c-1.548-1.324-3.119-2.126-4.648-.804l-.913.799c-.668.58-1.91 3.29-6.712-2.234c-4.801-5.517-1.944-6.376-1.275-6.951l.918-.8c1.521-1.325.947-2.993-.15-4.71l-.662-1.04C7.842.573 6.642-.552 5.117.771l-.824.72c-.674.491-2.558 2.087-3.015 5.119c-.55 3.638 1.185 7.804 5.16 12.375c3.97 4.573 7.857 6.87 11.539 6.83c3.06-.033 4.908-1.675 5.486-2.272l.827-.721c1.521-1.322.576-2.668-.973-3.995z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneOffice(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 0a3 3 0 0 0-3 3v15a3 3 0 0 0 3 3v.406c0 .825-.203 1.51-.531 1.938C5.14 23.77 4.725 24 4 24c-.725 0-1.14-.23-1.469-.656C2.203 22.917 2 22.23 2 21.406v-4.812a1 1 0 0 0-1.219-1a1 1 0 0 0-.781 1v4.812c0 1.164.295 2.28.969 3.157C1.643 25.438 2.735 26 4 26c1.265 0 2.357-.562 3.031-1.438c.674-.875.969-1.992.969-3.156V21a3 3 0 0 0 3-3V3a3 3 0 0 0-3-3zm9 0a3 3 0 0 0-3 3v15a3 3 0 0 0 3 3h7c1.657 0 3.906-1.313 3.906-10.5S23.657 0 22 0zm.5 3h7a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5h-7a.5.5 0 0 1-.5-.5v-3a.5.5 0 0 1 .5-.5m0 6h1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 .5-.5m3 0h1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 .5-.5m3 0h1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 .5-.5m-6 3h1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 .5-.5m3 0h1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 .5-.5m3 0h1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 .5-.5m-6 3h1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 .5-.5m3 0h1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 .5-.5m3 0h1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 .5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.208 11.857L6.902 5.26a1.312 1.312 0 0 0-1.268.052a1.272 1.272 0 0 0-.619 1.09V19.6c0 .443.233.856.619 1.089a1.316 1.316 0 0 0 1.269.052l13.306-6.599c.438-.218.716-.658.716-1.143s-.279-.924-.717-1.142"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Polyline(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24 .188A1.812 1.812 0 0 0 22.187 2c0 .193.039.388.094.563l-2.75 2.718A1.862 1.862 0 0 0 19 5.188c-.5 0-.953.203-1.281.53L11.78 3.782A1.793 1.793 0 0 0 10 2.187A1.812 1.812 0 0 0 8.187 4c0 .918.678 1.66 1.563 1.781l3.031 9.875c-.107.098-.2.224-.281.344l-4.781-.563A1.794 1.794 0 0 0 6 14.188A1.812 1.812 0 0 0 4.187 16c0 .406.15.76.375 1.063L2 22.188A1.812 1.812 0 1 0 3.813 24c0-.415-.141-.788-.376-1.094L6 17.813c.61 0 1.14-.32 1.469-.782l4.843.563A1.775 1.775 0 0 0 14 18.812A1.812 1.812 0 0 0 15.813 17c0-.907-.662-1.648-1.532-1.781L11.25 5.313V5.28h.031l5.939 1.97c.12.885.863 1.563 1.781 1.563A1.812 1.812 0 0 0 20.813 7a1.88 1.88 0 0 0-.125-.625l2.718-2.688c.185.064.388.126.594.126a1.812 1.812 0 1 0 0-3.626z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Previous(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.125 0H4.875A4.874 4.874 0 0 0 0 4.875v16.25A4.874 4.874 0 0 0 4.875 26h16.25A4.874 4.874 0 0 0 26 21.125V4.875A4.874 4.874 0 0 0 21.125 0M16 17.949a.964.964 0 0 1-.479.817a.986.986 0 0 1-.952.039L8.34 13.857c-.337-.296-.537-.494-.537-.857s.237-.575.537-.857l6.229-4.949a.99.99 0 0 1 .952.04c.29.173.479.484.479.816z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Privacy(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 0C9.155 0 6 3.155 6 7v1H5c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V10c0-1.1-.9-2-2-2h-1V7c0-3.845-3.155-7-7-7m0 2c2.755 0 5 2.245 5 5v1H8V7c0-2.755 2.245-5 5-5m0 9c2.8 0 5 2.2 5 5s-2.2 5-5 5s-5-2.2-5-5s2.2-5 5-5m0 2c-1.7 0-3 1.3-3 3s1.3 3 3 3s3-1.3 3-3c0-.3-.094-.606-.094-.906c-.3.5-.806.906-1.406.906c-.8 0-1.5-.7-1.5-1.5c0-.6.406-1.106.906-1.406c-.3 0-.606-.094-.906-.094"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QrCode(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v9h9V0zm11 0v2h2V0zm2 2v2h-2v4h2V6h2V2zm0 6v3H4v2H2v2h4v-2h2v2h5v-2h2v2h3v-2h2v-2h-5V8zm7 5v2h6v-2h-2v-2h-2v2zM2 13v-2H0v2zM17 0v9h9V0zM2 2h5v5H2zm17 0h5v5h-5zM3 3v3h3V3zm17 0v3h3V3zm-9 13v2h2v-2zm2 2v2h-2v4h2v-2h4v-2h2v-2h2v2h-2v2h-2v4h2v-2h2v-2h1v2h-1v2h2v-2h1v-2h2v-4h-2v-2h-7v2zm11 6v2h2v-2zm-11 0v2h2v-2zM0 17v9h9v-9zm2 2h5v5H2zm1 1v3h3v-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Radio(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.5.063c-.794 0-1.438.643-1.438 1.437L4.657 6.313c-.2-.076-.43-.125-.656-.125A1.812 1.812 0 0 0 2.187 8c0 .041-.002.084 0 .125C.933 8.484 0 9.63 0 11v11c0 1.656 1.344 3 3 3h20c1.656 0 3-1.344 3-3V11c0-1.656-1.344-3-3-3H5.812c0-.277-.076-.546-.187-.781l16.656-5c.25.428.688.719 1.219.719a1.437 1.437 0 1 0 0-2.876zm-6 10.75a5.696 5.696 0 0 1 5.688 5.687a5.696 5.696 0 0 1-5.688 5.688a5.697 5.697 0 0 1-5.688-5.688a5.697 5.697 0 0 1 5.688-5.688zm-13 .093c.877 0 1.594.717 1.594 1.594c0 .877-.717 1.594-1.594 1.594A1.597 1.597 0 0 1 2.906 12.5c0-.877.716-1.594 1.594-1.594m13 1.281c-.937 0-1.793.306-2.5.813h5a4.263 4.263 0 0 0-2.5-.813M14 14a4.298 4.298 0 0 0-.531 1h8.062A4.375 4.375 0 0 0 21 14zm-9.5 1.906c.877 0 1.594.717 1.594 1.594c0 .877-.717 1.594-1.594 1.594A1.597 1.597 0 0 1 2.906 17.5c0-.877.716-1.594 1.594-1.594m8.75.094c-.02.166-.063.328-.063.5s.043.334.063.5h8.5c.02-.166.063-.328.063-.5s-.044-.334-.063-.5zm.219 2c.134.358.312.693.531 1h7c.219-.307.398-.642.531-1H13.47zM15 20a4.266 4.266 0 0 0 2.5.813c.938 0 1.793-.306 2.5-.813z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rain(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 0c-2.29 0-4.188 1.293-5.344 3.094C9.94 3.064 9.735 3 9.5 3C6.998 3 5 4.735 4.344 7.031C1.946 7.12 0 9.081 0 11.5C0 13.973 2.027 16 4.5 16h16c3.026 0 5.5-2.474 5.5-5.5c0-2.575-1.822-4.662-4.219-5.25C21.183 2.279 18.64 0 15.5 0m0 2a4.47 4.47 0 0 1 4.469 4.125l.093.844l.813.062A3.479 3.479 0 0 1 24 10.5c0 1.944-1.556 3.5-3.5 3.5h-16A2.485 2.485 0 0 1 2 11.5a2.485 2.485 0 0 1 2.875-2.469l1.063.188l.093-1.063A3.48 3.48 0 0 1 9.5 5c.26 0 .53.027.813.094l.78.187l.345-.718A4.492 4.492 0 0 1 15.5 2m5.469 14.844c-.95.492-3.124 1.286-3.688 2.187a1.92 1.92 0 0 0 .594 2.657c.899.565 2.021.227 2.656-.625c.565-.758.437-2.18.438-4.22zm-14 1c-.95.492-3.124 1.286-3.688 2.187a1.92 1.92 0 0 0 .594 2.657c.899.565 2.021.227 2.656-.625c.565-.758.437-2.18.438-4.22zm7 3c-.95.492-3.124 1.286-3.688 2.187a1.92 1.92 0 0 0 .594 2.657c.899.565 2.021.227 2.656-.625c.565-.758.437-2.18.438-4.22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Record(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 6.188a6.812 6.812 0 1 0 0 13.625a6.812 6.812 0 1 0 0-13.625"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RecurringAppointment(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 0C5.838 0 0 5.838 0 13c0 7.163 5.838 13 13 13c6.555 0 12-4.888 12.875-11.219a1.5 1.5 0 1 0-2.969-.406C22.233 19.245 18.071 23 13 23C7.46 23 3 18.541 3 13C3 7.46 7.46 3 13 3c3.236 0 6.109 1.535 7.938 3.906l-1.782 1.782c-.15.15-.189.391-.125.593c.06.202.229.332.438.375c2.697.53 5.807.296 5.937.281a.623.623 0 0 0 .344-.187c.09-.088.17-.18.188-.313c.015-.13.25-3.267-.282-5.968a.55.55 0 0 0-.375-.438c-.204-.062-.444-.025-.593.125l-1.625 1.625C20.678 1.87 17.047 0 13 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RefreshShield(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.633 5.028a1.074 1.074 0 0 0-.777-.366c-2.295-.06-5.199-2.514-7.119-3.477C14.551.592 13.769.201 13.18.098a1.19 1.19 0 0 0-.359.001c-.589.103-1.372.494-2.556 1.087c-1.921.962-4.825 3.417-7.121 3.476c-.295.008-.577.14-.777.366a1.167 1.167 0 0 0-.291.834c.494 10.023 4.088 16.226 10.396 19.831c.164.093.346.141.528.141s.364-.048.528-.141c6.308-3.605 9.902-9.808 10.396-19.831a1.167 1.167 0 0 0-.291-.834m-5.925 8.183c-.287.787-1.887 4.454-6.47 3.633c-.136.477-.464 1.514-.565 1.834c-.019.062-.049.066-.086.008c-.391-.585-2.644-3.984-2.688-4.054c-.052-.074.029-.109.029-.109l4.329-1.984c.135-.063.099.088.099.088l-.543 1.975s.643.386 1.908.358c1.866-.038 3.489-1.391 3.925-1.786c.071-.07.097-.055.062.037m.326-3.695c-.119.06-3.651 1.643-4.371 2.003c-.135.069-.099-.062-.099-.062l.561-2.033s-.652-.392-1.936-.366c-1.831.039-3.428 1.32-3.932 1.77c-.107.093-.142.083-.097-.038c.331-.879 1.977-4.461 6.545-3.642c.133-.46.441-1.433.559-1.806c.046-.127.135-.021.135-.021l2.68 4.045c-.001 0 .044.106-.045.15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RemoveImage(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2a2 2 0 0 0-2 2c0 .738 1.156 2.057 1.156 3C1.156 7.943 0 10.921 0 13s1.156 4.991 1.156 6C1.156 20.009 0 21.262 0 22a2 2 0 0 0 2 2c.738 0 3.124-1.125 4.5-1.125S8.396 24 9.5 24s2.131-1.125 3.5-1.125c1.369 0 1.896 1.125 3 1.125s2.563-1.125 3.5-1.125c.937 0 3.762 1.125 4.5 1.125a2 2 0 0 0 2-2c0-.738-1.156-2.057-1.156-3c0-.943 1.156-4.05 1.156-6c0-1.95-1.156-4.991-1.156-6C24.844 5.991 26 4.738 26 4a2 2 0 0 0-2-2c-.738 0-3.124 1.156-4.5 1.156S17.604 2 16.5 2S14.369 3.156 13 3.156C11.631 3.156 11.104 2 10 2S7.438 3.156 6.5 3.156C5.562 3.156 2.738 2 2 2m2 3h18c.551 0 1 .449 1 1v14c0 .551-.449 1-1 1H4c-.551 0-1-.449-1-1V6c0-.551.449-1 1-1m6.031 3.063a.751.751 0 0 0-.562.218L8.28 9.47a.812.812 0 0 0 0 1.156L10.656 13l-2.375 2.375a.81.81 0 0 0 0 1.156L9.47 17.72a.812.812 0 0 0 1.156 0L13 15.344l2.375 2.375a.814.814 0 0 0 1.156 0l1.188-1.188a.814.814 0 0 0 0-1.156L15.344 13l2.375-2.375a.814.814 0 0 0 0-1.156L16.53 8.28a.814.814 0 0 0-1.156 0L13 10.656l-2.375-2.375c-.16-.162-.382-.218-.594-.219z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rename(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.188 0a1.502 1.502 0 0 0 .312 3c2.164 0 2.895.586 3.219.969c.309.366.283.516.281.531v3.188a1.001 1.001 0 0 0 0 .593v9.531a1 1 0 0 0 0 .407V21.5s.025.163-.281.531c-.306.369-1.007.969-3.219.969a1.5 1.5 0 1 0 0 3c2.367 0 3.981-.654 5-1.531c1.019.877 2.633 1.531 5 1.531a1.5 1.5 0 1 0 0-3c-2.212 0-2.913-.6-3.219-.969c-.306-.368-.281-.531-.281-.531V19h5a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1h-5V4.5a1.5 1.5 0 0 0 0-.031c-.002-.043.009-.177.281-.5C20.605 3.586 21.336 3 23.5 3a1.5 1.5 0 1 0 0-3c-2.36 0-3.981.676-5 1.563C17.481.675 15.86 0 13.5 0a1.5 1.5 0 0 0-.156 0a1.502 1.502 0 0 0-.156 0M.812 7A1 1 0 0 0 0 8v10a1 1 0 0 0 1 1h14v-2H2V9h13V7H1a1 1 0 0 0-.094 0a1 1 0 0 0-.094 0M20 9h4v8h-4zM4 11v4h9v-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RenewSubscription(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 0v19c0 .555-.445 1-1 1c-.555 0-1-.445-1-1V7h1V5H0v14c0 1.645 1.355 3 3 3h10c-.2-.6-.313-1.3-.313-2H5.814c.114-.316.187-.647.187-1V2h16v11c.7.2 1.4.5 2 1V0zm4 4v4h12V4zm0 6v2h5v-2zm7 0v2h5v-2zm-7 3v2h5v-2zm12 1c-3.324 0-6 2.676-6 6s2.676 6 6 6v-2c-2.276 0-4-1.724-4-4s1.724-4 4-4s4 1.724 4 4c0 .868-.247 1.67-.688 2.313L22 21l-.5 4.5L26 25l-1.25-1.25C25.581 22.706 26 21.377 26 20c0-3.324-2.676-6-6-6M8 16v2h5v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Repeat(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.156 0a.86.86 0 0 0-.375.094l-4.25 4.187c-.254.24-.469.411-.469.719c0 .307.185.467.47.719l4.25 4.187c.253.13.569.116.812-.031A.811.811 0 0 0 17 9.187V7h2c1.682 0 3 1.318 3 3v6c0 1.682-1.318 3-3 3a2 2 0 1 0 0 4c3.842 0 7-3.158 7-7v-6c0-3.842-3.158-7-7-7h-2V.812a.81.81 0 0 0-.406-.687a.853.853 0 0 0-.438-.125M7 3c-3.842 0-7 3.158-7 7v6c0 3.842 3.158 7 7 7h2v2.188c0 .28.16.541.406.687a.843.843 0 0 0 .438.125a.86.86 0 0 0 .375-.094l4.25-4.187c.254-.24.469-.411.469-.719c0-.306-.185-.467-.47-.719l-4.25-4.187a.847.847 0 0 0-.812.031a.811.811 0 0 0-.406.688V19H7c-1.682 0-3-1.318-3-3v-6c0-1.682 1.318-3 3-3a2 2 0 1 0 0-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RestrictionShield(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.633 5.028a1.074 1.074 0 0 0-.777-.366c-2.295-.06-5.199-2.514-7.119-3.477C14.551.592 13.768.201 13.18.098a1.225 1.225 0 0 0-.36.001c-.588.103-1.371.494-2.556 1.087c-1.92.962-4.824 3.417-7.119 3.476a1.08 1.08 0 0 0-.778.366a1.167 1.167 0 0 0-.291.834c.493 10.023 4.088 16.226 10.396 19.831c.164.093.346.141.527.141s.363-.048.528-.141c6.308-3.605 9.902-9.808 10.396-19.831a1.161 1.161 0 0 0-.29-.834M13 18.057a6.057 6.057 0 1 1 0-12.114a6.057 6.057 0 0 1 0 12.114m2.48-9.548L9.509 14.48A4.261 4.261 0 0 1 8.707 12A4.299 4.299 0 0 1 13 7.707c.926 0 1.777.301 2.48.802m1.01 1.011c.501.702.803 1.555.803 2.48A4.299 4.299 0 0 1 13 16.293a4.259 4.259 0 0 1-2.48-.802z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RetroTv(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.875 0a1 1 0 0 0-.594.281L13 6.563L9.719 3.28A1.016 1.016 0 1 0 8.28 4.72L11.562 8H3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h20a3 3 0 0 0 3-3V11a3 3 0 0 0-3-3h-8.563l6.282-6.281A1 1 0 0 0 19.875 0M10.5 10c7.445 0 8.5.021 8.5 7c0 6.98-1.128 7-8.5 7c-7.406 0-8.5-.074-8.5-7c0-6.924 1.094-7 8.5-7m12 2.938c.866 0 1.563.696 1.563 1.562a1.56 1.56 0 0 1-1.563 1.563a1.559 1.559 0 0 1-1.563-1.563a1.56 1.56 0 0 1 1.563-1.563zm0 4c.866 0 1.563.696 1.563 1.562a1.56 1.56 0 0 1-1.563 1.563a1.559 1.559 0 0 1-1.563-1.563a1.56 1.56 0 0 1 1.563-1.563z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rewind(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.688 5.125c-.202 0-.41.033-.594.125L3.78 11.844c-.4.376-.719.671-.719 1.156c0 .485.271.76.72 1.156l8.312 6.594c.403.203.897.17 1.281-.063c.386-.232.625-.65.625-1.093v-3.688l6.094 4.844c.403.203.897.17 1.281-.063c.386-.232.625-.65.625-1.093V6.406c0-.443-.239-.86-.625-1.093a1.33 1.33 0 0 0-.688-.188c-.2 0-.408.033-.593.125L14 10.094V6.406c0-.443-.239-.86-.625-1.093a1.33 1.33 0 0 0-.688-.188z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RfidTag(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.813 0a1 1 0 0 0-.532.281l-6 6A1 1 0 0 0 0 7v17c0 1.093.907 2 2 2h22c1.093 0 2-.907 2-2V2c0-1.093-.907-2-2-2H7a1 1 0 0 0-.094 0a1 1 0 0 0-.093 0m.625 2H24v22H2V7.437zm1.375 1a1 1 0 0 0-.532.281l-5 5A1 1 0 0 0 3 9v13a1 1 0 0 0 1 1h18a1 1 0 1 0 0-2H5V9.437L9.438 5H21v13H8v-7.563L10.438 8H18v5h-2v-1.75a.25.25 0 0 0-.25-.25h-5.5a.25.25 0 0 0-.25.25v5.5c0 .138.112.25.25.25h5.5a.25.25 0 0 0 .25-.25V15h3a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1h-9a1 1 0 0 0-.719.281l-3 3A1 1 0 0 0 6 10v9a1 1 0 0 0 1 1h15a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H9a1 1 0 0 0-.094 0a1 1 0 0 0-.094 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ruler(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 .063a.8.8 0 0 0-.188.03a.8.8 0 0 0-.375.22L.313 18.438a.8.8 0 0 0 0 1.093l6.157 6.157a.8.8 0 0 0 .656.218a.8.8 0 0 0 .438-.218L25.688 7.563a.8.8 0 0 0 0-1.094L19.53.312a.8.8 0 0 0-.06-.062a.8.8 0 0 0-.032-.031a.8.8 0 0 0-.28-.125A.8.8 0 0 0 19 .062zm-.031 1.968l5 4.969L7 23.969L2 19l1.219-1.219l1.531 1.532a.8.8 0 0 0 .781.28a.8.8 0 0 0 .157-.062a.8.8 0 0 0 .062-.031a.8.8 0 0 0 .156-.094a.8.8 0 0 0 .032-.031a.8.8 0 0 0 .125-.125a.8.8 0 0 0-.188-1.063l-1.531-1.53l1-1l1.531 1.53a.8.8 0 0 0 .781.282a.8.8 0 0 0 .157-.063a.8.8 0 0 0 .062-.031a.8.8 0 0 0 .156-.094a.8.8 0 0 0 .031-.031a.8.8 0 0 0 .126-.125A.8.8 0 0 0 8 16.062l-1.531-1.53l1-1l2.906 2.937a.8.8 0 0 0 .188.156a.8.8 0 0 0 .062.031a.8.8 0 0 0 .156.063a.8.8 0 0 0 .063.031a.8.8 0 0 0 .312 0a.8.8 0 0 0 .156-.063a.8.8 0 0 0 .063-.03a.8.8 0 0 0 .156-.095a.8.8 0 0 0 .031-.03a.8.8 0 0 0 .126-.126a.8.8 0 0 0-.188-1.062l-2.906-2.938l.969-.969l1.53 1.5a.807.807 0 0 0 .813.22a.807.807 0 0 0 .406-.282a.807.807 0 0 0-.062-1.063l-1.531-1.53l.969-.97l1.53 1.5a.807.807 0 0 0 .813.22a.807.807 0 0 0 .406-.282a.807.807 0 0 0-.062-1.063l-1.531-1.53l.969-.97l2.937 2.907a.8.8 0 0 0 .406.25a.8.8 0 0 0 .063.031a.8.8 0 0 0 .312 0a.8.8 0 0 0 .157-.063a.8.8 0 0 0 .062-.03a.8.8 0 0 0 .156-.095a.8.8 0 0 0 .032-.03a.8.8 0 0 0 .125-.126a.8.8 0 0 0-.188-1.062l-2.938-2.906l1-1l1.532 1.53a.807.807 0 0 0 .75.22a.807.807 0 0 0 .156-.063a.807.807 0 0 0 .063-.031a.807.807 0 0 0 .125-.094a.807.807 0 0 0 .03-1.188l-1.53-1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 .188A9.812 9.812 0 0 0 .187 10A9.812 9.812 0 0 0 10 19.813c2.29 0 4.393-.811 6.063-2.125l.875.875a1.845 1.845 0 0 0 .343 2.156l4.594 4.625c.713.714 1.88.714 2.594 0l.875-.875a1.84 1.84 0 0 0 0-2.594l-4.625-4.594a1.824 1.824 0 0 0-2.157-.312l-.875-.875A9.812 9.812 0 0 0 10 .188M10 2a8 8 0 1 1 0 16a8 8 0 0 1 0-16M4.937 7.469a5.446 5.446 0 0 0-.812 2.875a5.46 5.46 0 0 0 5.469 5.469a5.516 5.516 0 0 0 3.156-1a7.166 7.166 0 0 1-.75.03a7.045 7.045 0 0 1-7.063-7.062c0-.104-.005-.208 0-.312"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SecurityChecked(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.633 5.028a1.074 1.074 0 0 0-.777-.366c-2.295-.06-5.199-2.514-7.119-3.477C14.551.592 13.768.201 13.18.098a1.225 1.225 0 0 0-.36.001c-.588.103-1.371.494-2.556 1.087c-1.92.962-4.824 3.416-7.119 3.476a1.08 1.08 0 0 0-.778.366a1.167 1.167 0 0 0-.291.834c.493 10.023 4.088 16.226 10.396 19.831c.164.093.346.141.527.141s.363-.048.528-.141c6.308-3.605 9.902-9.808 10.396-19.831a1.161 1.161 0 0 0-.29-.834M18.617 8.97l-5.323 7.855c-.191.282-.491.469-.788.469c-.298 0-.629-.163-.838-.372l-3.752-3.753a.656.656 0 0 1 0-.926l.927-.929a.658.658 0 0 1 .926 0l2.44 2.44l4.239-6.257a.657.657 0 0 1 .91-.173l1.085.736a.657.657 0 0 1 .174.91"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sent(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 2v8.5L15 13L0 15.5V24l26-11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Settings(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.469.969l-.563 3.562A8.704 8.704 0 0 0 8.5 5.5L5.562 3.406L3.438 5.531L5.5 8.47a8.767 8.767 0 0 0-1 2.406l-3.531.594v3l3.531.625a8.683 8.683 0 0 0 1 2.406l-2.094 2.938l2.125 2.125L8.47 20.5a8.717 8.717 0 0 0 2.406.969l.594 3.562h3l.656-3.562a8.627 8.627 0 0 0 2.375-1l2.969 2.093l2.125-2.125L20.47 17.5c.438-.73.79-1.526 1-2.375l3.562-.656v-3l-3.562-.594a8.746 8.746 0 0 0-1-2.375l2.093-2.969l-2.125-2.125L17.5 5.531a8.767 8.767 0 0 0-2.406-1L14.469.97zM13 6.469A6.535 6.535 0 0 1 19.531 13A6.535 6.535 0 0 1 13 19.531A6.536 6.536 0 0 1 6.469 13A6.536 6.536 0 0 1 13 6.469m0 1.593A4.949 4.949 0 0 0 8.062 13A4.949 4.949 0 0 0 13 17.938A4.949 4.949 0 0 0 17.938 13A4.949 4.949 0 0 0 13 8.062m-.031 2.876c1.146 0 2.094.915 2.094 2.062c0 1.146-.948 2.063-2.094 2.063A2.054 2.054 0 0 1 10.906 13c0-1.147.917-2.063 2.063-2.063z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBag(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 0c-2.206 0-3.683.992-4.344 2.281C7.996 3.571 8 4.983 8 6a1 1 0 1 0 2 0c0-.983.073-2.07.438-2.781C10.802 2.508 11.306 2 13 2c1.685 0 2.194.522 2.563 1.25C15.93 3.978 16 5.068 16 6a1 1 0 1 0 2 0c0-.968 0-2.36-.656-3.656C16.687 1.047 15.215 0 13 0M4.187 8c-.4 0-.775.313-.875.813L.188 24.687C.087 25.387.4 26 1 26h24c.6 0 .913-.613.813-1.313L22.688 8.813c-.1-.5-.476-.812-.875-.812H4.188zM8.5 9c.8 0 1.5.7 1.5 1.5S9.3 12 8.5 12S7 11.3 7 10.5S7.7 9 8.5 9m9 0c.8 0 1.5.7 1.5 1.5s-.7 1.5-1.5 1.5s-1.5-.7-1.5-1.5s.7-1.5 1.5-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBasket(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.094 4.094a1.02 1.02 0 0 0-.719.281a.998.998 0 0 0 0 1.406l6.844 6.844a.947.947 0 0 0 .687.281c.253 0 .526-.088.719-.281a.997.997 0 0 0 0-1.406L2.781 4.375a.95.95 0 0 0-.687-.281m21.812 0a.95.95 0 0 0-.687.281l-6.844 6.844a.997.997 0 0 0 0 1.406c.193.193.466.281.719.281a.947.947 0 0 0 .687-.281l6.844-6.844a.998.998 0 0 0 0-1.406a1.02 1.02 0 0 0-.719-.281M1 11c-.551 0-1 .449-1 1v3c0 .551.449 1 1 1h1v1c0 .089 2 8 2 8c.297.531.547 1 1 1h16c.551 0 .766-.438 1-1c0 0 2-7.911 2-8v-1h1c.551 0 1-.449 1-1v-3c0-.551-.449-1-1-1h-4.094l-2.562 2.563c-.39.39-.918.624-1.469.624c-.55 0-1.047-.233-1.438-.624a2.043 2.043 0 0 1-.28-2.563h-4.313a2.043 2.043 0 0 1-.281 2.563c-.39.39-.888.624-1.438.624a2.088 2.088 0 0 1-1.469-.624L5.094 11zm8 5c.551 0 1 .449 1 1v6c0 .551-.449 1-1 1c-.551 0-1-.449-1-1v-6c0-.551.449-1 1-1m4 0c.551 0 1 .449 1 1v6c0 .551-.449 1-1 1c-.551 0-1-.449-1-1v-6c0-.551.449-1 1-1m4 0c.551 0 1 .449 1 1v6c0 .551-.449 1-1 1c-.551 0-1-.449-1-1v-6c0-.551.449-1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCart(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.719 0A1.004 1.004 0 0 0 1 2h3.469c.574 0 .757.115.906.25c.149.135.335.407.5.969l4.688 15.937a1 1 0 0 0 .03.094S11.704 22 14.5 22h7.844a1 1 0 1 0 0-2H14.5c-1.475 0-2.063-1.5-2.063-1.5l-.156-.5h8.125c1.03 0 1.496-.864 1.594-1.188l3.594-7.718C26.15 7.76 25.859 6 23.812 6h-14.5c-.166 0-.351.042-.53.094c-.36-1.223-.928-3.19-1-3.438c-.21-.713-.49-1.387-1.063-1.906C6.146.231 5.349 0 4.469 0H1a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.094 0A1.004 1.004 0 0 0 .72 0zm8.906 8.063h4.219v1.78h-3.656l-.47-1.5c-.032-.11-.08-.204-.093-.28m5.531 0h2.688v1.78h-2.688zm4 0h4.625a1.17 1.17 0 0 1-.062.218L23 9.844h-3.844zm-8.562 3.093h3.25v1.688h-2.719zm4.562 0h2.688v1.688h-2.688zm4 0h3.219l-.781 1.688h-2.438zm-7.625 3h2.313v1.781h-1.656c-.04 0-.062-.013-.094-.03l-.563-1.75zm3.625 0h2.688v1.781h-2.688v-1.78zm4 0H21l-.844 1.781h-1v-1.78zM14 22.187a1.812 1.812 0 1 0 0 3.625a1.812 1.812 0 0 0 0-3.625m6 0a1.812 1.812 0 1 0 0 3.625a1.812 1.812 0 0 0 0-3.625"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shuffle(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.813 0a.883.883 0 0 0-.407.125a.811.811 0 0 0-.406.688V3h-1c-.672 0-1.48.096-2.219.5c-.739.404-1.301 1.008-1.812 1.688c-.766 1.018-1.706 2.61-2.938 4.687a2 2 0 0 0-.25-1.969C9.76 6.21 9.126 5.143 8.25 4.281C7.374 3.42 6.02 3 5 3H2a2 2 0 0 0-.188 0A2.002 2.002 0 1 0 2 7h3c.5 0 .161-.116.438.156c.276.272.942 1.161 1.937 2.813a2 2 0 0 0 3.344.406c-.33.555-.542.937-.938 1.594c-2.113 3.506-3.292 5.538-3.969 6.437c-.338.45-.488.556-.5.563c-.011.006.017.031-.312.031H2a2 2 0 1 0 0 4h3c.671 0 1.48-.096 2.219-.5c.738-.404 1.301-1.009 1.812-1.688c1.023-1.357 2.098-3.314 4.188-6.78c2.113-3.507 3.262-5.54 3.937-6.438c.338-.45.52-.557.532-.563c.01-.006-.02-.031.312-.031h1v2.188c0 .28.16.541.406.687a.843.843 0 0 0 .438.125a.86.86 0 0 0 .375-.094l5.25-4.187c.254-.24.468-.411.468-.719c.002-.306-.184-.467-.468-.719L20.219.094A.816.816 0 0 0 19.812 0zm0 16a.883.883 0 0 0-.407.125a.811.811 0 0 0-.406.688V19h-1c-.45 0-.21.077-.375-.063c-.164-.139-.665-.74-1.406-1.937a2 2 0 0 0-1.844-.969a2 2 0 0 0-1.563 3.063c.799 1.288 1.372 2.157 2.22 2.875C15.877 22.687 17.086 23 18 23h1v2.188c0 .28.16.541.406.687a.843.843 0 0 0 .438.125a.86.86 0 0 0 .375-.094l5.25-4.187c.254-.24.468-.411.468-.719c.002-.306-.184-.467-.468-.719l-5.25-4.187a.816.816 0 0 0-.407-.094z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shutdown(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 0a2 2 0 0 0-2 2v10a2 2 0 0 0 4 0V2a2 2 0 0 0-2-2M5.844 2.594a1.5 1.5 0 0 0-.782.344A12.796 12.796 0 0 0 .188 13C.188 20.058 5.94 25.813 13 25.813S25.813 20.058 25.813 13c0-4.073-1.902-7.717-4.875-10.063a1.5 1.5 0 1 0-1.875 2.344A9.804 9.804 0 0 1 22.813 13A9.789 9.789 0 0 1 13 22.813A9.789 9.789 0 0 1 3.187 13a9.808 9.808 0 0 1 3.75-7.719a1.5 1.5 0 0 0-1.093-2.687"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Signature(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.594 1c-1.419 0-2.903.917-4.219 2.625C1.059 5.333.041 7.907.094 11.313c0 1.626.519 3.43 1.5 4.937c.144.222.393.355.562.563c-.688.831-1.875 2.406-1.875 2.406a1 1 0 1 0 1.438 1.375s1.203-1.61 1.843-2.375c.738.469 1.582.781 2.532.781c1.708 0 3.312-.817 4.687-1.938c.195-.158.313-.391.5-.562c.441.309.96.5 1.531.5c.959 0 1.773-.457 2.47-1c.696-.543 1.293-1.197 1.843-1.844c.369-.434.639-.794.938-1.156c-.037.471-.067.935-.063 1.375c.002.23.012.437.094.719c.04.14.076.334.25.531s.528.375.843.375c.817 0 1.425-.458 1.875-.906c.45-.449.811-.965 1.22-1.438C23.096 12.711 23.88 12 25 12a1 1 0 1 0 0-2c-1.98 0-3.31 1.29-4.219 2.344c-.282.327-.494.55-.718.812c.002-.036-.003-.057 0-.094c.038-.508.1-1.003.062-1.5a2.19 2.19 0 0 0-.188-.812c-.14-.306-.608-.75-1.125-.75c-.541 0-.846.238-1.093.438s-.463.432-.688.687c-.45.51-.925 1.116-1.437 1.719c-.513.603-1.056 1.199-1.563 1.594c-.506.394-.927.562-1.219.562c-.066 0-.1-.113-.156-.125c.527-.633 1.175-1.177 1.594-1.844c.913-1.456 1.548-2.897 1.75-4.156c.101-.63.137-1.242-.156-1.844C15.55 6.43 14.784 6 14.094 6c-1.345 0-2.476.87-3.25 2.031c-.775 1.162-1.25 2.686-1.25 4.375c0 .877.275 1.614.562 2.344c-.238.23-.41.55-.656.75c-1.156.942-2.365 1.5-3.406 1.5c-.453 0-.857-.198-1.25-.406c.467-.622.503-.54 1-1.282C7.892 12.253 10 8.343 10 4.595c0-.445-.03-1.25-.5-2.063C9.03 1.718 7.982 1 6.594 1m0 2c.812 0 .976.22 1.156.531c.18.312.25.808.25 1.063c0 2.952-1.892 6.709-3.844 9.625c-.399.596-.399.514-.781 1.031c-.027-.04-.067-.053-.094-.094c-.744-1.143-1.187-2.67-1.187-3.844a1 1 0 0 0 0-.03c-.047-2.995.84-5.065 1.875-6.407C5.003 3.533 6.312 3 6.594 3m7.437 5.031c.014.074.052.206 0 .531c-.129.804-.657 2.113-1.469 3.407c-.223.355-.583.62-.843.969c-.024-.208-.125-.304-.125-.532c0-1.31.38-2.493.906-3.281c.502-.754 1.083-1.065 1.531-1.094M0 23v2h26v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SimCardChip(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.813 2C2.647 2 0 4.648 0 7.813v10.375C0 21.352 2.648 24 5.813 24h14.375C23.352 24 26 21.352 26 18.187V7.813C26 4.649 23.352 2 20.187 2zm0 2h14.375C22.223 4 24 5.777 24 7.813V9h-6c-.555 0-1-.445-1-1c0-.555.445-1 1-1a1 1 0 1 0 0-2c-1.645 0-3 1.355-3 3c0 1.292.844 2.394 2 2.813v4.968c-1.198.814-2 2.18-2 3.719c0 .923.293 1.781.781 2.5H10.22a4.439 4.439 0 0 0 .78-2.5c0-1.538-.802-2.905-2-3.719v-4.969c1.156-.418 2-1.52 2-2.812c0-1.645-1.355-3-3-3H6a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.093 0A1.004 1.004 0 0 0 6 7h2c.555 0 1 .445 1 1c0 .555-.445 1-1 1H2V7.812C2 5.777 3.777 4 5.813 4M2 11h5v4H2zm17 0h5v4h-5zM2 17h4.5C7.839 17 9 18.161 9 19.5S7.839 22 6.5 22h-.688C3.777 22 2 20.223 2 18.187zm17.5 0H24v1.188C24 20.223 22.223 22 20.187 22H19.5c-1.339 0-2.5-1.161-2.5-2.5s1.161-2.5 2.5-2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkipToStart(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 5c-.551 0-1 .449-1 1v14c0 .551.449 1 1 1h3c.551 0 1-.449 1-1v-4.875l7.094 5.625c.403.203.897.17 1.281-.063c.386-.232.625-.65.625-1.093V6.406c0-.443-.239-.86-.625-1.093a1.33 1.33 0 0 0-.688-.188c-.2 0-.408.033-.593.125L11 10.875V6c0-.551-.449-1-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Snow(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 0c-2.29 0-4.188 1.293-5.344 3.094C9.94 3.064 9.735 3 9.5 3C6.998 3 5 4.735 4.344 7.031C1.946 7.12 0 9.081 0 11.5C0 13.973 2.027 16 4.5 16h16c3.026 0 5.5-2.474 5.5-5.5c0-2.575-1.822-4.662-4.219-5.25C21.183 2.279 18.64 0 15.5 0m0 2a4.47 4.47 0 0 1 4.469 4.125l.093.844l.813.062A3.479 3.479 0 0 1 24 10.5c0 1.944-1.556 3.5-3.5 3.5h-16A2.485 2.485 0 0 1 2 11.5a2.485 2.485 0 0 1 2.875-2.469l1.063.188l.093-1.063A3.48 3.48 0 0 1 9.5 5c.26 0 .53.027.813.094l.78.187l.345-.718A4.492 4.492 0 0 1 15.5 2m3.688 14.844a.518.518 0 0 0-.188.906l1.219 1.219H18.5a.518.518 0 0 0-.063 0a.532.532 0 1 0 .063 1.062h1.719L19 21.25a.53.53 0 1 0 .75.75l1.219-1.219V22.5a.531.531 0 0 0 1.062 0v-1.719L23.25 22a.53.53 0 1 0 .75-.75l-1.219-1.219H24.5a.531.531 0 0 0 0-1.062h-1.719L24 17.75a.518.518 0 0 0-.375-.906a.518.518 0 0 0-.375.156l-1.219 1.219V17.5a.518.518 0 0 0-.593-.531a.518.518 0 0 0-.063 0a.518.518 0 0 0-.406.531v.719L19.75 17a.518.518 0 0 0-.438-.156a.518.518 0 0 0-.062 0a.518.518 0 0 0-.063 0zm-13.875.125a.518.518 0 0 0-.344.531v1.719L3.75 18a.518.518 0 0 0-.438-.156a.518.518 0 0 0-.062 0a.518.518 0 0 0-.25.906l1.219 1.219H2.5a.518.518 0 0 0-.063 0A.532.532 0 1 0 2.5 21.03h1.719L3 22.25a.53.53 0 1 0 .75.75l1.219-1.219V23.5a.531.531 0 0 0 1.062 0v-1.719L7.25 23a.53.53 0 1 0 .75-.75l-1.219-1.219H8.5a.531.531 0 0 0 0-1.062H6.781L8 18.75a.518.518 0 0 0-.375-.906a.518.518 0 0 0-.375.156l-1.219 1.219V17.5a.518.518 0 0 0-.593-.531a.518.518 0 0 0-.063 0a.518.518 0 0 0-.063 0zm8 2a.518.518 0 0 0-.344.531v1.719L11.75 20a.518.518 0 0 0-.438-.156a.518.518 0 0 0-.312.906l1.219 1.219H10.5a.518.518 0 0 0-.063 0a.532.532 0 1 0 .063 1.062h1.719L11 24.25a.53.53 0 1 0 .75.75l1.219-1.219V25.5a.531.531 0 0 0 1.062 0v-1.719L15.25 25a.53.53 0 1 0 .75-.75l-1.219-1.219H16.5a.531.531 0 0 0 0-1.062h-1.719L16 20.75a.518.518 0 0 0-.375-.906a.518.518 0 0 0-.375.156l-1.219 1.219V19.5a.518.518 0 0 0-.594-.531a.518.518 0 0 0-.062 0a.518.518 0 0 0-.063 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Speaker(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.031 1.063c-.321.001-.676.145-1 .468L5.312 8H1c-.551 0-1 .449-1 1v8c0 .551.449 1 1 1h4.313L11 24.438c1 1 2 .488 2-.875V2.28c0-.791-.433-1.222-.969-1.219zm7.25 2a1 1 0 0 0-.218 1.906A8.963 8.963 0 0 1 24 13c0 3.524-2 6.55-4.938 8.031a1 1 0 1 0 .875 1.782C23.53 21 26 17.288 26 13s-2.471-8-6.063-9.813a1 1 0 0 0-.562-.124a1 1 0 0 0-.094 0m-2.375 3.874a1 1 0 0 0-.406 1.875C18.043 9.771 19 11.29 19 13c0 1.722-.972 3.261-2.531 4.219a1 1 0 1 0 1.062 1.687C19.601 17.636 21 15.476 21 13c0-2.461-1.387-4.633-3.438-5.906A1 1 0 0 0 17 6.937a1 1 0 0 0-.094 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpeechBubble(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 .188C5.924.188.187 5.252.187 11.5c0 3.177 1.488 6.039 3.876 8.094c-.521 3.009-3.887 4.234-3.657 5.062c3.01 1.245 8.971-1.645 9.875-2.093c.874.166 1.789.25 2.719.25c7.076 0 12.813-5.065 12.813-11.313S20.075.187 13 .187z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stack(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 0C3.344 0 2.625 1.422 2 3L.062 8.25C.346 8.094.654 8 1 8h24c.346 0 .654.094.938.25L24 3c-.703-1.531-1.344-3-3-3zM1 9c-.551 0-1 .449-1 1v3c0 .551.449 1 1 1h24c.551 0 1-.449 1-1v-3c0-.551-.449-1-1-1zm21.5 1a1.5 1.5 0 1 1-.001 3.001A1.5 1.5 0 0 1 22.5 10M1 15c-.551 0-1 .449-1 1v3c0 .551.449 1 1 1h24c.551 0 1-.449 1-1v-3c0-.551-.449-1-1-1zm21.5 1a1.5 1.5 0 1 1-.001 3.001A1.5 1.5 0 0 1 22.5 16M1 21c-.551 0-1 .449-1 1v3c0 .551.449 1 1 1h24c.551 0 1-.449 1-1v-3c0-.551-.449-1-1-1zm21.5 1a1.5 1.5 0 1 1-.001 3.001A1.5 1.5 0 0 1 22.5 22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StackOfPhotos(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 0C.449 0 0 .449 0 1v16c0 .551.449 1 1 1h16c.551 0 1-.449 1-1V1c0-.551-.449-1-1-1zm1 2h14v12H2zm17 .906v2.031l1.813.313L19 15.75V17c0 1.104-.897 2-2 2H6.406l12.688 2.188a1 1 0 0 0 1.156-.813l2.688-15.781a.999.999 0 0 0-.813-1.157zM9 3.937c-1.151 0-2.125.792-2.125 2.282c0 .974.434 1.952 1.031 2.562c.234.61-.164.842-.25.875c-1.206.436-2.625 1.245-2.625 2.031v1.282h7.938v-1.281c0-.81-1.422-1.614-2.688-2.032c-.058-.019-.417-.18-.187-.875c.595-.61 1.062-1.593 1.062-2.562c0-1.49-1.005-2.282-2.156-2.282m14.406 3.97l-.343 1.968l.718.156l-2.75 11.688l-.406-.094a1.954 1.954 0 0 1-1.719.531L5.063 19.781L4.78 20.97a1.023 1.023 0 0 0 .75 1.218l15.563 3.657a1.023 1.023 0 0 0 1.218-.75L25.938 9.53c.127-.536-.18-1.091-.718-1.219l-1.813-.406z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StanleyKnife(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.344 0c-1.317 0-2.602.53-3.657 1.438c-.006.005-.024-.006-.03 0a1 1 0 0 0-.313.218l-12.969 13a1.931 1.931 0 0 0-.531 1.219c-.015.361.06.635.094.875l.25 1.688l-2 2l1.093 5.468L6.47 21.72l1.687.25c.205.032.483.107.844.093a1.877 1.877 0 0 0 1.25-.53L19.781 12c.464-.465 1.198-.767 2.032-1.063c.833-.295 1.778-.528 2.53-1.28c2.199-2.199 2.199-5.803 0-8a5.656 5.656 0 0 0-4-1.657zm-2 2.594l5.062 5.031c-.137.21-.283.408-.468.594c-.222.22-.912.503-1.782.812c-.87.31-1.925.705-2.781 1.563l-9.438 9.469h-.03c-.091 0-.247-.033-.438-.063a1 1 0 0 0-.031 0l-3.094-.469l-.438-3.062c-.032-.236-.06-.389-.062-.469L17.625 3.187a1 1 0 0 0 .156-.124c.18-.18.362-.336.563-.47zM16.906 5.75l-9.625 9.594l2.282 2.281L19.155 8zm-2.093 3.563l.78.75l-2.28 2.28l-.75-.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.326 10.137a1.001 1.001 0 0 0-.807-.68l-7.34-1.066l-3.283-6.651c-.337-.683-1.456-.683-1.793 0L8.82 8.391L1.48 9.457a1 1 0 0 0-.554 1.705l5.312 5.178l-1.254 7.31a1.001 1.001 0 0 0 1.451 1.054L13 21.252l6.564 3.451a1 1 0 0 0 1.451-1.054l-1.254-7.31l5.312-5.178a.998.998 0 0 0 .253-1.024"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Statistics(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.906-.031a1 1 0 0 0-.125.031A1 1 0 0 0 12 1v1H3a3 3 0 0 0-3 3v13c0 1.656 1.344 3 3 3h9v.375l-5.438 2.719a1.006 1.006 0 0 0 .875 1.812L12 23.625V24a1 1 0 1 0 2 0v-.375l4.563 2.281a1.006 1.006 0 0 0 .875-1.812L14 21.375V21h9c1.656 0 3-1.344 3-3V5a3 3 0 0 0-3-3h-9V1a1 1 0 0 0-1.094-1.031M2 5h22v13H2zm18.875 1a1 1 0 0 0-.594.281L17 9.563L14.719 7.28a1 1 0 0 0-1.594.219l-2.969 5.188l-1.219-3.063a1 1 0 0 0-1.656-.344l-3 3a1.016 1.016 0 1 0 1.439 1.44l1.906-1.906l1.438 3.562a1 1 0 0 0 1.812.125l3.344-5.844l2.062 2.063a1 1 0 0 0 1.438 0l4-4A1 1 0 0 0 20.875 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stop(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 20c0 .551-.449 1-1 1H6c-.551 0-1-.449-1-1V6c0-.551.449-1 1-1h14c.551 0 1 .449 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sun(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.906-.031a1 1 0 0 0-.125.031A1 1 0 0 0 12 1v3a1 1 0 1 0 2 0V1a1 1 0 0 0-1.094-1.031M6.844 1.594a1 1 0 0 0-.719 1.5l1.5 2.625a1.008 1.008 0 0 0 1.75-1l-1.5-2.625a1 1 0 0 0-1.031-.5m11.968 0a1 1 0 0 0-.687.5l-1.5 2.625a1.008 1.008 0 0 0 1.75 1l1.5-2.625a1 1 0 0 0-.969-1.5a1 1 0 0 0-.093 0zm4.532 4.375A1 1 0 0 0 23.25 6a1 1 0 0 0-.344.125l-2.625 1.5a1.008 1.008 0 0 0 1 1.75l2.625-1.5a1 1 0 0 0-.562-1.906M2.437 6a1 1 0 0 0-.343 1.875l2.625 1.5a1.008 1.008 0 0 0 1-1.75l-2.625-1.5A1 1 0 0 0 2.438 6zM13 6a7 7 0 1 0 0 14a7 7 0 0 0 0-14m0 2c2.757 0 5 2.243 5 5s-2.243 5-5 5s-5-2.243-5-5s2.243-5 5-5M.719 12A1.004 1.004 0 0 0 1 14h3a1 1 0 1 0 0-2H1a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.094 0a1.004 1.004 0 0 0-.093 0m21 0A1.004 1.004 0 0 0 22 14h3a1 1 0 1 0 0-2h-3a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.093 0a1.004 1.004 0 0 0-.094 0M5.156 16.469a1 1 0 0 0-.093.031a1 1 0 0 0-.344.125l-2.625 1.5a1.008 1.008 0 0 0 1 1.75l2.625-1.5a1 1 0 0 0-.563-1.906m15.469.031a1 1 0 0 0-.344 1.875l2.625 1.5a1.008 1.008 0 0 0 1-1.75l-2.625-1.5a1 1 0 0 0-.656-.125M8.312 19.781a1 1 0 0 0-.687.5l-1.5 2.625a1.008 1.008 0 0 0 1.75 1l1.5-2.625a1 1 0 0 0-.969-1.5a1 1 0 0 0-.094 0m9.032 0a1 1 0 0 0-.719 1.5l1.5 2.625a1.008 1.008 0 0 0 1.75-1l-1.5-2.625a1 1 0 0 0-1.031-.5m-4.438 1.188a1 1 0 0 0-.125.031A1 1 0 0 0 12 22v3a1 1 0 1 0 2 0v-3a1 1 0 0 0-1.094-1.031"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Survey(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.813 0A1 1 0 0 0 8 1v1H4.406C3.606 2 3 2.606 3 3.406V24.5c0 .9.606 1.5 1.406 1.5H21.5c.8 0 1.406-.606 1.406-1.406V3.406c.1-.8-.512-1.406-1.312-1.406H18V1a1 1 0 0 0-1-1H9a1 1 0 0 0-.094 0a1 1 0 0 0-.094 0zM10 2h6v2h-6zM5 4h3v1a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V4h3v20H5zm2 5v4h4V9zm1 1h2v2H8zm5 0v2h6v-2zm-6 5v4h4v-4zm6 1v2h6v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwissArmyKnife(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 0A4.5 4.5 0 0 0 8 4.5v17a4.5 4.5 0 0 0 9 0v-17A4.5 4.5 0 0 0 12.5 0M.062.406c-.164 1.052-.09 3.636.813 4.875L7 13.781V7.47l-1.125-.157l.219-1.343l-1.344-.22l.219-1.313l-1.344-.218l.219-1.344L2.5 2.687l.219-1.343l-1.344-.219C.983 1.034.31.541.062.406M12.5 3a1.5 1.5 0 1 1-.001 3.001A1.5 1.5 0 0 1 12.5 3M18 4v6.531s1.802 2.182 2.938 3.563c1.006 1.224 3.015 2.743 4.968 2.094c-.83-1.948-1.6-3.413-2.25-4.875C21.591 8.466 18 4 18 4M1 12c-.747.072-1 .806-1 2.188C0 17.907 7 22 7 22v-2.844s-2-.726-2-3.062c0-1.318-.309-2.148-1-2.719c-.219-.181-.9-.438-1.219-.219c.271.477 1.044 2.183.625 2.438C2.83 15.947.832 15.274 1 12m11 8h1v1h1v1h-1v1h-1v-1h-1v-1h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwitchCamera(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.125 5H19.2l-1.304-2.301C17.67 2.3 17.18 2 16.5 2h-7c-.709 0-1.169.3-1.396.699L6.8 5H4.875A4.874 4.874 0 0 0 0 9.875v8.25A4.874 4.874 0 0 0 4.875 23h16.25A4.874 4.874 0 0 0 26 18.125v-8.25A4.874 4.874 0 0 0 21.125 5M13 21a7 7 0 0 1-7-7H4l3-4l3 4H8c0 2.757 2.243 5 5 5c1.268 0 2.415-.49 3.297-1.271l1.214 1.619A6.968 6.968 0 0 1 13 21m6-3l-3-4h2c0-2.757-2.243-5-5-5c-1.268 0-2.415.49-3.297 1.271L8.489 8.652A6.968 6.968 0 0 1 13 7a7 7 0 0 1 7 7h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Timer(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 0a2 2 0 1 0-.001 3.999A2 2 0 0 0 13 0m5.156 1.25c-.293.016-.634.159-.875.563c-.688 1.153.392 1.852.719 2.03C21.244 5.62 23.469 9.048 23.469 13c0 5.77-4.699 10.469-10.469 10.469S2.531 18.77 2.531 13c0-1.326.265-2.584.719-3.75c.059-.151.616-1.173-.5-1.875c-.764-.48-1.506.064-1.656.406A12.94 12.94 0 0 0 0 13c0 7.18 5.82 13 13 13s13-5.82 13-13c0-5.113-2.956-9.538-7.25-11.656c-.045-.023-.3-.11-.594-.094M3.375 3.344c-.314.314 7.727 10.476 7.969 10.781c.008.012.022.019.031.031l.094.125c.366.44.914.718 1.531.719a2.003 2.003 0 0 0 2-2a1.984 1.984 0 0 0-.719-1.531S3.694 3.025 3.375 3.344"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Today(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 0c-.551 0-1 .449-1 1v3c0 .551.449 1 1 1c.551 0 1-.449 1-1V1c0-.551-.449-1-1-1m12 0c-.551 0-1 .449-1 1v3c0 .551.449 1 1 1c.551 0 1-.449 1-1V1c0-.551-.449-1-1-1M3 2C1.344 2 0 3.344 0 5v18c0 1.656 1.344 3 3 3h20c1.656 0 3-1.344 3-3V5c0-1.656-1.344-3-3-3h-2v2a2 2 0 0 1-4 0V2H9v2a2 2 0 0 1-4 0V2zM2 9h22v14c0 .551-.449 1-1 1H3c-.551 0-1-.449-1-1zm14.906 2.156a.575.575 0 0 0-.375.25l-4.906 7.25l-2.281-2.25a.595.595 0 0 0-.844 0l-.875.844a.627.627 0 0 0 0 .875l3.469 3.469c.195.193.505.343.781.343s.572-.176.75-.437l5.906-8.719a.607.607 0 0 0-.156-.844l-1-.687a.64.64 0 0 0-.469-.094M902 1469v2h26v-2zm4 5v2h18v-2zm-4 5v2h26v-2zm4 5v2h18v-2zm-4 5v2h26v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TodoList(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.875 0a1 1 0 0 0-.656.375L3.812 4.656l-2.25-1.5A1.014 1.014 0 1 0 .438 4.844l3 2a1 1 0 0 0 1.344-.219l4-5A1 1 0 0 0 7.875 0M12 3v2h14V3zM7.875 9a1 1 0 0 0-.656.375l-3.407 4.281l-2.25-1.5a1.014 1.014 0 1 0-1.125 1.688l3 2a1 1 0 0 0 1.344-.219l4-5A1 1 0 0 0 7.875 9M12 12v2h14v-2zm-4.125 6a1 1 0 0 0-.656.375l-3.407 4.281l-2.25-1.5a1.014 1.014 0 1 0-1.125 1.688l3 2a1 1 0 0 0 1.344-.219l4-5A1 1 0 0 0 7.875 18M12 21v2h14v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwoFswipeDown(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 6a4 4 0 0 0-4 4c0 1.863 1.276 3.4 3 3.844v5.719L5.719 18.28A1 1 0 0 0 4.78 18a1 1 0 0 0-.5 1.719l3 3a1 1 0 0 0 1.438 0l3-3a1.016 1.016 0 1 0-1.438-1.438L9 19.563v-5.72c1.724-.444 3-1.98 3-3.843a4 4 0 0 0-4-4m10 0a4 4 0 0 0-4 4c0 1.863 1.276 3.4 3 3.844v5.719l-1.281-1.282A1 1 0 0 0 14.78 18a1 1 0 0 0-.5 1.719l3 3a1 1 0 0 0 1.438 0l3-3a1.016 1.016 0 1 0-1.438-1.438L19 19.563v-5.72c1.724-.444 3-1.98 3-3.843a4 4 0 0 0-4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwoFswipeRight(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.906 3.969a1 1 0 0 0-.125.031a1 1 0 0 0-.5 1.719L16.562 7h-5.718C10.399 5.276 8.864 4 7 4a4 4 0 1 0 0 8c1.863 0 3.4-1.276 3.844-3h5.719l-1.282 1.281a1.016 1.016 0 1 0 1.438 1.438l3-3a1 1 0 0 0 0-1.438l-3-3a1 1 0 0 0-.813-.312m4 10a1 1 0 0 0-.125.031a1 1 0 0 0-.5 1.719L20.563 17h-5.72c-.444-1.724-1.98-3-3.843-3a4 4 0 1 0 0 8c1.863 0 3.4-1.276 3.844-3h5.719l-1.282 1.281a1.016 1.016 0 1 0 1.438 1.438l3-3a1 1 0 0 0 0-1.438l-3-3a1 1 0 0 0-.813-.312"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlock(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 0C4.79 0 2.878.917 1.687 2.406C.498 3.896 0 5.826 0 7.906V11h3V7.906c0-1.58.389-2.82 1.031-3.625C4.674 3.477 5.541 3 7 3c1.463 0 2.328.45 2.969 1.25c.64.8 1.031 2.06 1.031 3.656V9h3V7.906c0-2.092-.527-4.044-1.719-5.531C11.09.888 9.206 0 7 0m2 10c-1.656 0-3 1.344-3 3v10c0 1.656 1.344 3 3 3h14c1.656 0 3-1.344 3-3V13c0-1.656-1.344-3-3-3zm7 5a2 2 0 0 1 2 2c0 .738-.404 1.372-1 1.719V21c0 .551-.449 1-1 1c-.551 0-1-.449-1-1v-2.281c-.596-.347-1-.98-1-1.719a2 2 0 0 1 2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnlockTwo(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 0C9.676 0 7 2.676 7 6v4c-2.2 0-4 1.8-4 4v8c0 2.2 1.8 4 4 4h12c2.2 0 4-1.8 4-4v-8c0-2.2-1.8-4-4-4H9V6c0-2.276 1.724-4 4-4s4 1.724 4 4v1h2V6c0-3.324-2.676-6-6-6m0 15c1.1 0 2 .9 2 2c0 .7-.4 1.387-1 1.688V21c0 .6-.4 1-1 1s-1-.4-1-1v-2.313c-.6-.3-1-.987-1-1.687c0-1.1.9-2 2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserShield(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.633 5.028a1.074 1.074 0 0 0-.777-.366c-2.295-.06-5.199-2.514-7.119-3.477C14.551.592 13.768.201 13.18.098a1.225 1.225 0 0 0-.36.001c-.588.103-1.371.494-2.556 1.087c-1.92.962-4.824 3.417-7.119 3.476a1.08 1.08 0 0 0-.778.366a1.167 1.167 0 0 0-.291.834c.493 10.023 4.088 16.226 10.396 19.831c.164.093.346.141.527.141s.363-.048.528-.141c6.308-3.605 9.902-9.808 10.396-19.831a1.161 1.161 0 0 0-.29-.834m-5.57 10.249c0 1.36-2.604 1.67-5.048 1.67c-2.44 0-5.077-.31-5.077-1.67v-.374c0-.999 1.8-2.011 3.333-2.564c.111-.041.641-.329.345-1.106c-.76-.775-1.334-2.034-1.334-3.271c0-1.896 1.254-2.889 2.719-2.889s2.726.993 2.726 2.889c0 1.232-.577 2.485-1.332 3.264h.003c-.289.881.174 1.09.248 1.114c1.61.532 3.418 1.536 3.418 2.564z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoCall(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 6C2.5 6 0 6.813 0 7.469V18.5c0 .641 2.5 1.5 8.5 1.5s8.5-.813 8.5-1.469V7.5C17 6.859 14.5 6 8.5 6m16.875.031a.885.885 0 0 0-.469.157l-6.5 4.187c-.25.188-.406.498-.406.813v3.624c0 .315.156.624.406.813l6.5 4.188c.176.133 1.032.55 1.032-.813V7c0-.787-.271-.981-.563-.969"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewFile(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 0C1.8 0 0 1.8 0 4v17c0 2.2 1.8 4 4 4h11c.4 0 .7-.094 1-.094c-1.4-.3-2.594-1.006-3.594-1.906H4c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h6.313c.7.2.687 1.1.687 2v3c0 .6.4 1 1 1h3c1 0 2 0 2 1v1h.5c.5 0 1 .088 1.5.188V8c0-1.1-.988-2.112-2.688-3.813c-.3-.2-.512-.487-.812-.687c-.2-.3-.488-.513-.688-.813C13.113.988 12.1 0 11 0zm13.5 12c-3 0-5.5 2.5-5.5 5.5s2.5 5.5 5.5 5.5c1.273 0 2.435-.471 3.375-1.219l.313.313a.955.955 0 0 0 .125 1.218l2.5 2.5c.4.4.975.4 1.375 0l.5-.5c.4-.4.4-1.006 0-1.406l-2.5-2.5a.935.935 0 0 0-1.157-.156l-.281-.313c.773-.948 1.25-2.14 1.25-3.437c0-3-2.5-5.5-5.5-5.5m0 1.5c2.2 0 4 1.8 4 4s-1.8 4-4 4s-4-1.8-4-4s1.8-4 4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Voicemail(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 7c-3.302 0-6 2.698-6 6s2.698 6 6 6h14c3.302 0 6-2.698 6-6s-2.698-6-6-6s-6 2.698-6 6a5.98 5.98 0 0 0 1.531 4H10.47A5.98 5.98 0 0 0 12 13c0-3.302-2.698-6-6-6m0 2c2.221 0 4 1.779 4 4s-1.779 4-4 4s-4-1.779-4-4s1.779-4 4-4m14 0c2.221 0 4 1.779 4 4s-1.779 4-4 4s-4-1.779-4-4s1.779-4 4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeDown(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.031 2.063c-.321.001-.676.145-1 .468L8.437 9H5c-.551 0-1 .449-1 1v6c0 .551.449 1 1 1h3.438L15 23.438c1 1 2 .488 2-.875V3.28c0-.791-.433-1.222-.969-1.219zm2.5 7.03a1 1 0 0 0-.031 1.97c.857.223 1.5.999 1.5 1.937c0 .938-.643 1.714-1.5 1.938a1 1 0 1 0 .5 1.937c1.721-.45 3-2.025 3-3.875s-1.279-3.425-3-3.875a1 1 0 0 0-.375-.031a1 1 0 0 0-.094 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeUp(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.031 2.063c-.321.001-.676.145-1 .468L4.437 9H1c-.551 0-1 .449-1 1v6c0 .551.449 1 1 1h3.438L11 23.438c1 1 2 .488 2-.875V3.28c0-.791-.433-1.222-.969-1.219zm7.25 1a1 1 0 0 0-.218 1.906A8.963 8.963 0 0 1 24 13c0 3.524-2 6.55-4.938 8.031a1 1 0 1 0 .875 1.782C23.53 21 26 17.288 26 13s-2.471-8-6.063-9.813a1 1 0 0 0-.562-.124a1 1 0 0 0-.094 0m-2.156 2.75a1 1 0 0 0-.281 1.906A5.983 5.983 0 0 1 20 13a5.984 5.984 0 0 1-3.125 5.281a1 1 0 1 0 .938 1.75A8 8 0 0 0 22 13c0-3.038-1.712-5.679-4.219-7.031a1 1 0 0 0-.468-.157a1 1 0 0 0-.188 0zm-2.594 3.28a1 1 0 0 0-.031 1.97c.857.223 1.5.999 1.5 1.937c0 .938-.643 1.714-1.5 1.938a1 1 0 1 0 .5 1.937c1.721-.45 3-2.025 3-3.875s-1.279-3.425-3-3.875a1 1 0 0 0-.375-.031a1 1 0 0 0-.094 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WeddingCake(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.844.094c-1.613 0-2.75 1.142-2.75 2.75c0 1.77 1.398 2.921 2.75 4.031c.442.363.874.728 1.281 1.125H5c-1.316 0-2.409.851-2.813 2.031c0 0 .204 1.969 3.313 1.969c2.89 0 2.997-1.47 3-1.531A.495.495 0 0 1 9 10a.517.517 0 0 1 .5.5c.004.051.145 1.5 3.5 1.5s3.495-1.47 3.5-1.531a.512.512 0 0 1 .531-.469a.49.49 0 0 1 .469.5c.002.053.09 1.5 3 1.5c3.203 0 3.313-1.969 3.313-1.969C23.409 8.852 22.314 8 21 8h-7.125a19.66 19.66 0 0 1 1.281-1.125c1.351-1.11 2.75-2.262 2.75-4.031c0-1.608-1.136-2.75-2.75-2.75c-.93 0-1.717.666-2.156 1.125c-.437-.46-1.227-1.125-2.156-1.125M9.03 11.719C8.542 12.355 7.542 13 5.5 13c-2.005 0-3.001-.625-3.5-1.25V16h22v-4.25c-.499.62-1.498 1.25-3.5 1.25c-2.05 0-3.044-.647-3.531-1.281C16.415 12.352 15.285 13 13 13c-2.293 0-3.418-.645-3.969-1.281zM3 17c-1.316 0-2.409.851-2.813 2.031C.188 19.031.391 21 3.5 21c2.89 0 3.997-1.47 4-1.531A.495.495 0 0 1 8 19a.517.517 0 0 1 .5.5c.004.051 1.145 1.5 4.5 1.5s4.495-1.47 4.5-1.531a.512.512 0 0 1 .531-.469a.49.49 0 0 1 .469.5c.002.053 1.09 1.5 4 1.5c3.203 0 3.313-1.969 3.313-1.969C25.409 17.852 24.314 17 23 17zm5.031 3.719C7.542 21.355 5.542 22 3.5 22c-2.005 0-3.001-.625-3.5-1.25V24c0 1.656 1.344 2 3 2h20c1.656 0 3-.344 3-2v-3.25c-.499.62-1.498 1.25-3.5 1.25c-2.05 0-4.044-.647-4.531-1.281C17.415 21.352 15.285 22 13 22c-2.293 0-4.418-.645-4.969-1.281"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindRose(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.875 0a1 1 0 0 0-.844.719l-1.937 6.375l-4-2.094c-.3-.2-.7-.112-1 .188c-.3.3-.294.606-.094.906l2.094 4L.719 12.03a1 1 0 0 0 0 1.938l6.343 1.906L5 19.906c-.2.3-.112.7.188 1c.2.2.425.282.625.282c.1 0 .275.006.375-.094l3.968-2.063l1.875 6.25a1 1 0 0 0 1.938 0l1.906-6.343L19.813 21c.1.1.274.094.375.094c.2 0 .425-.082.625-.282c.3-.3.287-.7.187-1l-2.094-3.906l6.375-1.937a1 1 0 0 0 0-1.938L19 10.125l2-4.031c.2-.3.113-.7-.188-1c-.3-.3-.7-.288-1-.188L15.845 6.97L13.969.719A1 1 0 0 0 12.875 0M13 4.469V13l2.094-2.094a1 1 0 0 0 .406.25L21.531 13H13l2.094 2.094a1 1 0 0 0-.25.406L13 21.531V13l-2.094 2.094a1 1 0 0 0-.406-.25L4.469 13H13l-2.094-2.094a1 1 0 0 0 .25-.406z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WorldwideLocation(children ...ElementRenderer) *WpfIcon {
	return &WpfIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 0c-2.209 0-4 1.793-4 4c0 2.665 3.5 6.306 3.5 9.969h1C19.501 10.306 23 6.77 23 4c0-2.207-1.791-4-4-4m0 1.938a2.063 2.063 0 1 1-.002 4.126A2.063 2.063 0 0 1 19 1.937zM13 2C6.383 2 1 7.383 1 14s5.383 12 12 12s12-5.383 12-12c0-2.01-.498-3.896-1.375-5.563c-.133.264-.267.515-.406.782c-.263.508-.54 1.021-.782 1.531c.25.715.422 1.471.5 2.25h-1.312c-.072.335-.125.65-.125.969V15h1.438a9.88 9.88 0 0 1-1.282 4H18.72c.22-.938.347-1.975.437-3.031h-1.97a18.72 18.72 0 0 1-.53 3.031H9.343a19.24 19.24 0 0 1-.594-4h7.75v-1.031a4.68 4.68 0 0 0-.125-.969H8.75a19.24 19.24 0 0 1 .594-4h5.344a31.369 31.369 0 0 1-.97-2h-3.687c.827-1.887 1.92-3 2.969-3c0-.695.125-1.348.344-1.969C13.229 2.027 13.116 2 13 2M8.875 4.906a12.364 12.364 0 0 0-1 2.094h-2a9.991 9.991 0 0 1 3-2.094M4.344 9H7.28a21.423 21.423 0 0 0-.531 4H3.062a9.865 9.865 0 0 1 1.282-4m-1.282 6H6.75c.059 1.425.243 2.77.531 4H4.344a9.895 9.895 0 0 1-1.282-4m2.813 6h2c.288.776.625 1.473 1 2.094a9.991 9.991 0 0 1-3-2.094m4.156 0h5.938c-.827 1.887-1.92 3-2.969 3s-2.143-1.113-2.969-3m8.094 0h2a9.991 9.991 0 0 1-3 2.094c.375-.62.712-1.317 1-2.094"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
