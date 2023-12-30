package ls

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = ""
	hAttr          = "1em"
	viewbox        = "0 0 717 717"
	fill           = "currentColor"
)

type LsIcon struct {
	*SVGSVGElement
}

func Aicon(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M491 222v-94h72v526h-72v-74c-50 55-124 91-206 91C133 671 0 551 0 401s133-270 285-270c82 0 156 36 206 91m0 185v-12c-4-110-95-198-207-198c-115 0-215 91-215 204s100 204 215 204c112 0 203-88 207-198"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AupperCase(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M547 762h-81l-89-242H170L82 762H0L273 0zM273 236l-77 214h155z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Addstar(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M78 621h561c43 0 78-35 78-78V150c0-43-35-78-78-78H78c-43 0-78 35-78 78v393c0 43 35 78 78 78m-36-78V150c0-19 17-35 36-35h561c19 0 36 16 36 35v393c0 19-17 35-36 35H78c-19 0-36-16-36-35m94-19l114-60l114 60l-22-127l92-89l-127-20l-57-114l-56 114l-128 20l92 89zm426-177h80v45h-80v81h-46v-81h-81v-45h81v-81h46z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Adjust(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M358 0c198 0 359 161 359 359S556 717 358 717S0 557 0 359S160 0 358 0m0 63v592c164 0 296-132 296-296S522 63 358 63"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Aim(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M447 51c25 25 37 54 37 88c0 29-8 55-26 77c-17 22-40 37-67 44c33 4 64-2 93-17l21-13l58 133l-18 10c-45 25-85 40-119 45c57 45 103 104 136 178l7 16l-133 79l-11-19c-22-50-65-92-128-123c-25 34-50 61-74 78c-38 29-86 50-145 62l-19 5L0 533l45-7c12-4 24-8 37-15c38-18 65-46 83-83c2-3 3-12 4-24c7-89 38-152 91-190c-17-22-25-47-25-75c0-34 12-63 37-88c24-25 53-37 88-37c34 0 63 12 87 37m61 294l-24-55c-46 19-90 22-132 8l-9-3l-4-10c-1-4-2-10-3-18v-5c-17-4-30-9-43-18c-26 15-45 35-57 60c-13 25-21 60-25 104c-2 20-4 33-8 39c-29 61-77 101-145 119l28 78c47-11 87-30 117-55c19-15 39-38 63-69l10-13c3-4 8-6 14-6c3 0 7 1 9 3c72 31 124 73 154 128l62-36c-30-58-68-107-116-145c-15-11-30-22-46-30l-21-11c-7-3-10-9-9-18c2-14 8-21 20-21l13 2l62 3c26-3 56-13 90-31M360 60c-22 0-41 7-57 23s-24 34-24 56s8 41 24 56c16 16 35 24 57 24s40-8 56-24c16-15 24-34 24-56s-8-40-24-56s-34-23-56-23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Album(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M81 0h434c45 0 81 37 81 81v475c0 43-35 79-78 81v-54c13-2 24-14 24-27V81c0-14-13-26-27-26H81c-14 0-27 12-27 26c5-1 12-2 18-2c5 0 9 0 14 1l325 49c44 8 76 48 76 92v423c0 42-30 73-71 73c-4 0-10 0-15-1L77 667c-43-8-77-48-77-91V81C0 37 37 0 81 0m16 255l285 42l7-49l-285-42zm60 43l-4 30l158 24l4-30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alignadjust(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M29 50h609c15 0 28 13 28 27v41c0 16-13 29-28 29H29c-16 0-29-13-29-29V77c0-15 13-27 29-27m0 166h609c15 0 28 12 28 26v42c0 15-13 28-28 28H29c-16 0-29-13-29-28v-42c0-14 13-26 29-26m0 166h609c15 0 28 12 28 26v42c0 15-13 28-28 28H29c-16 0-29-13-29-28v-42c0-14 13-26 29-26m0 164h609c15 0 28 14 28 28v41c0 16-13 29-28 29H29c-16 0-29-13-29-29v-41c0-15 13-28 29-28"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Aligncenter(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M500 147H167c-16 0-29-14-29-28V77c0-14 13-27 29-27h333c14 0 27 13 27 27v42c0 15-13 28-27 28m110 162H55c-15 0-27-12-27-28v-41c0-16 12-27 27-27h555c15 0 28 11 28 27v41c0 16-13 28-28 28m-83 163H139c-16 0-28-13-28-28v-42c0-15 12-28 28-28h388c14 0 28 13 28 28v42c0 15-14 28-28 28M29 546h609c15 0 28 14 28 28v41c0 16-13 29-28 29H29c-16 0-29-13-29-29v-41c0-15 13-28 29-28"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alignleft(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 77v42c0 14 13 28 29 28h331c15 0 28-13 28-28V77c0-14-13-27-28-27H29C13 50 0 63 0 77m0 163v41c0 16 13 28 29 28h554c15 0 27-12 27-28v-41c0-15-12-27-27-27H29c-16 0-29 12-29 27m0 162v42c0 15 13 28 29 28h388c14 0 26-13 26-28v-42c0-15-12-28-26-28H29c-16 0-29 13-29 28m0 172v41c0 16 13 29 29 29h609c15 0 28-13 28-29v-41c0-14-13-28-28-28H29c-16 0-29 13-29 28"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alignright(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M638 147H305c-15 0-27-14-27-28V77c0-14 12-27 27-27h333c15 0 28 13 28 27v42c0 15-13 28-28 28m0 162H83c-15 0-28-12-28-28v-41c0-16 13-27 28-27h555c15 0 28 11 28 27v41c0 16-13 28-28 28m0 163H250c-15 0-28-13-28-28v-42c0-15 13-28 28-28h388c15 0 28 13 28 28v42c0 15-13 28-28 28M29 546h609c15 0 28 14 28 28v41c0 16-13 29-28 29H29c-16 0-29-13-29-29v-41c0-15 13-28 29-28"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Amazon(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M525.008 221v155c0 47 20 68 38 94c6 9 7 19-1 26c-19 17-57 49-76 67l-1-1c-6 6-15 7-23 3c-33-28-40-41-58-67c-54 55-92 73-164 73c-84 0-150-53-150-157c0-81 45-136 107-163c54-24 130-27 188-34v-14c0-24 2-51-12-71c-12-18-37-27-57-27c-38 0-71 19-79 60c-2 9-9 18-18 19l-97-10c-8-2-17-10-14-22c22-118 128-152 223-152c49 0 112 13 150 49c49 46 44 106 44 172m-140 97v-21c-73 0-148 15-148 101c0 43 22 72 60 72c29 0 54-16 70-45c19-35 18-67 18-107m296 191c12 11-2 68-8 93c-2 7 5 9 10 3c38-41 39-116 29-125c-10-10-85-10-126 28c-6 6-4 13 3 11c26-6 81-21 92-10m-28 42c12-11-1-24-13-17c-83 49-175 84-261 96c-129 17-259-13-369-58c-9-4-14 6-6 11c104 72 235 121 372 103c98-13 206-66 277-135"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ampersand(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m606 490l-49-52c-23 20-53 44-93 73l-73-73l-112-118c51-43 86-79 103-107s26-53 26-78c0-24-7-45-20-67c-13-20-33-37-59-49c-25-12-52-19-83-19c-48 0-89 13-122 42c-28 24-42 54-42 91c0 20 5 43 17 66c11 24 37 60 77 108c-71 56-117 103-141 140c-23 37-35 73-35 110c0 48 17 89 52 122c34 32 79 48 136 48c42 0 83-7 121-22c39-14 90-47 150-96l98 100h102c-18-22-39-45-65-70c-33-33-59-57-79-75c29-20 59-45 91-74M234 266c-31-34-53-66-69-94c-7-12-11-24-11-37c0-17 9-33 25-46c16-12 37-18 63-18c28 0 51 6 67 18s25 27 25 44c0 14-7 32-20 51c-17 27-45 54-80 82m58 171l116 119c-51 40-96 70-136 86c-25 11-50 17-77 17c-34 0-63-11-87-33s-35-46-35-74c0-24 10-50 29-78s61-64 123-111c25 28 48 52 67 74"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Android(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m466 12l-39 55c61 29 105 83 110 146H121c5-63 48-117 109-146l-39-55c-2-4-1-8 2-10s7-1 10 2l40 58c26-10 55-16 86-16s59 6 86 16l40-58c2-3 6-4 9-2s4 6 2 10M241 155c13 0 23-10 23-23c0-12-10-23-23-23c-12 0-23 11-23 23c0 13 11 23 23 23m178 0c13 0 23-10 23-23c0-12-10-23-23-23c-12 0-23 11-23 23c0 13 11 23 23 23M0 474V289c0-26 20-48 47-48c26 0 46 22 46 48v185c0 27-20 48-46 48c-27 0-47-21-47-48m563 0V289c0-26 21-48 47-48s47 22 47 48v185c0 27-21 48-47 48s-47-21-47-48M122 581V242h416v339c0 20-17 37-37 37h-41v105c0 26-21 47-47 47s-47-21-47-47V618h-73v105c0 26-20 47-46 47c-27 0-47-21-47-47V618h-42c-19 0-36-17-36-37"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func App(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m577 429l-3 2c-15 8-25 14-29 14c-3 1-5 2-7 3c-7-13-27-52-57-104c-36-66-139-309-142-319c-2-10 0-20 7-24c1-1 3-1 5-1c6 0 13 4 17 7c6 5 123 184 186 298c38 69 51 94 56 103c-2 1-4 2-6 4c-4 5-18 13-24 16h-1zM353 129l13 7c22 12 27 38 16 60l-87-50c8-14 23-22 37-22c8 0 15 1 21 5m15 91l1 1c3 1 3 6 2 8s-3 3-5 3h-4l-1-1l-3 5h1c3 2 4 5 3 8c-2 2-3 3-5 3c-1 0-2 0-3-1l-1-1l-6 11l1 1c2 1 3 5 2 7s-3 3-5 3c-1 0-2 0-3-1h-1l-3 5l1 1c3 2 4 5 2 8c-1 2-2 3-4 3c-1 0-2-1-3-1l-1-1l-3 4l-4 7l-87-50l4-7l3-4l-1-1c-3-1-4-6-2-8c1-2 2-3 4-3c1 0 2 0 3 1h1l3-4l-1-1c-3-2-4-5-2-8c1-2 3-3 5-3c1 0 2 1 3 1l1 1l6-11h-1c-3-2-4-5-2-8c1-2 2-3 4-3c1 0 2 1 3 1l1 1l3-6l-1-1c-3-1-4-5-2-7c1-2 3-3 5-3h4l1 1l4-7l87 50zM56 558l172-298l87 51l-172 297l-4 7c-15-6-30-14-44-22c-14-9-29-19-42-29l1-3zM0 514V365h133L46 514zm521-32l14 21l5 5l3 6H232l87-149h139c34 61 55 101 55 102c2 3 3 5 5 6c1 3 2 7 3 9m136-48l-9-21c-2-3-3-5-5-7c-1-3-2-6-3-8l-8-16c-3-5-6-10-10-17h146v149h-75c0-1-1-3-1-5l-11-23c-1-2-2-5-4-7c0-3-1-5-2-7l-10-23c-1-2-4-4-5-6c-1-3-2-7-3-9m-4 65l10 20l1 3s-15 13-28 20s-31 13-31 13l-3-4l-11-18v-8l-8-6l-15-21v-7l-7-5l-8-13l-5-8c7-1 19-6 37-16c1 0 1-1 2-1s1-1 2-1c15-8 26-15 31-22l4 9l6 13v11l7 4l10 23l-1 10zm17 45l84 126c11 20 1 44-19 55c-7 4-15 6-22 6c-14 0-27-6-35-19l-61-139c7-3 19-8 29-13c8-5 18-11 24-16M32 715l-5 9c-1 2-4 3-7 3h-1c-1 0-2 0-2-1c-1 0-1-1-2-1c-2-2-2-6-1-8l5-9l29-122c12 9 25 18 37 25s25 13 38 18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Apple(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M426 78c8-22 13-41 13-59c0-3-1-5-1-8s-1-7-2-11c-50 11-85 33-107 65c-22 31-33 68-34 111c18-2 33-5 44-8c15-5 31-15 47-30c18-18 32-40 40-60m-57 100c-30 10-51 14-62 14c-9 0-28-4-59-12c-30-9-55-14-77-14c-49 0-90 21-123 63c-32 42-48 95-48 161c0 71 21 143 63 217c43 74 87 110 130 110c14 0 32-5 57-14c23-9 43-14 61-14s40 4 65 13c26 9 46 13 61 13c37 0 74-28 111-83c12-18 23-37 32-55c8-17 15-35 21-52c-27-8-50-27-69-55c-19-29-29-61-29-96c0-33 9-62 27-89c11-14 28-32 50-51c-7-9-15-18-22-25c-8-7-15-13-22-18c-28-18-59-28-93-28c-20 0-45 6-74 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arrowdown(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M466 0H251c-26 0-45 20-45 46v250H36c-16 0-27 7-34 22c-1 5-2 9-2 12c0 10 3 19 10 26l324 323c12 15 36 14 50 0l323-323c22-20 7-60-26-60H510V46c0-26-18-46-44-46"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arrowleft(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M689.13 466V251c0-26-19-45-45-45h-250V36c0-16-8-27-22-34c-5-1-10-2-14-2c-10 0-18 3-24 10l-324 324c-14 12-13 36 0 50l324 324c20 22 60 7 60-26V510h250c26 0 45-18 45-44"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arrowright(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 251v215c0 26 19 45 46 45h250v172c0 16 7 26 22 33c4 1 9 2 13 2c10 0 18-3 25-10l323-325c14-12 14-36 0-50L356 11c-20-22-60-8-60 25v171H46c-27 0-46 18-46 44"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arrowup(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M251.378 694h215c26 0 45-20 45-46V398h170c16 0 28-8 34-23c2-5 2-8 2-12c0-10-3-18-10-25l-324-323c-11-15-36-14-50 0l-323 323c-22 20-7 60 26 60h171v250c0 26 18 46 44 46"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Asciicircum(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 516L230 11l6-11h68l6 11l231 505h-79L270 97L79 516z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Asciitilde(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M508 325v-82c-42 27-70 43-84 49c-14 5-29 8-43 8c-11 0-25-3-39-7s-42-17-81-35c-57-27-101-40-134-40c-15 0-31 2-47 8c-16 5-43 19-80 39v81c37-23 65-37 83-43c17-7 34-10 48-10c10 0 22 2 34 6s41 16 85 36c56 26 100 37 132 37c14 0 30-2 46-8c15-5 43-19 80-39"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Asterisk(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m311 59l37 62l-158 58l158 60l-37 62l-129-108l28 167h-73l30-167L36 301L0 239l160-60L0 121l36-62l131 108L137 0h73l-28 167z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func At(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M514 199h53s-39 221-59 332c-17 96 108 14 145-29c36-41 118-223-27-368C495 4 297 40 211 90C110 149 19 287 49 446c33 178 158 259 233 278c163 41 270 2 372-103c17 0 32-1 52-1c-69 105-186 162-284 164c-86 1-191-2-299-102C14 581-22 430 13 293C48 157 151 46 292 12s327 2 421 168c80 141 14 287-38 345c-80 90-247 163-223 30c1-5 2-12 3-18c-39 36-88 59-139 59c-101 0-176-88-159-197s121-197 222-197c53 0 96 25 124 64c7-40 11-67 11-67M322 547c63 0 124-46 150-109c5-24 10-49 14-74c-4-66-50-113-116-113c-76 0-134 66-147 148s23 148 99 148"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bicon(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M72 0v324c53-58 132-96 216-96c152 0 275 120 275 270S440 768 288 768c-84 0-163-38-216-96v82H0V0zm215 702c115 0 207-91 207-204s-92-204-207-204S72 385 72 498s100 204 215 204"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BupperCase(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h175c58 0 109 27 141 71c23 30 35 66 35 105s-12 75-35 104c-3 5-7 10-10 15c34 12 63 31 88 58c38 40 61 96 61 155c0 60-23 114-61 155c-41 44-100 72-166 72H0zm72 280h108c55-2 99-48 99-104S235 74 180 71H72zm0 73v310h158c85-1 153-70 153-155s-68-155-153-155z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Back(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M489 65v563L0 347z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Backslash(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h82l477 854h-83z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Backspace(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M253 94h452c35 0 63 28 63 63v379c0 35-28 64-63 64H253L0 347zm313 408l57-56l-100-99l100-99l-57-57l-99 99l-98-99l-57 56l99 100l-99 98l57 56l99-98z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bad(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M505 122V85h109V15H243l-59 45v319s82 61 118 136c36 76 34 164 34 164h94s4-72 0-129c-5-58-24-129-24-129h260v-82H505v-37h161v-73H505v-37h130v-70zM0 64h142v308H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bag(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M680 221v120c-95 35-213 55-340 55S95 376 0 341V221c0-22 18-40 40-40h165v-30c0-46 38-85 85-85h100c47 0 85 39 85 85v30h165c22 0 40 18 40 40m-425-70v30h170v-30c0-19-16-35-35-35H290c-19 0-35 16-35 35m35 183v20c0 6 4 10 10 10h80c6 0 10-4 10-10v-20c0-6-4-10-10-10h-80c-6 0-10 4-10 10M0 602V384c96 33 213 52 340 52s244-19 340-52v218c0 22-18 40-40 40H40c-22 0-40-18-40-40m390-103v-20c0-6-4-10-10-10h-80c-6 0-10 4-10 10v20c0 6 4 10 10 10h80c6 0 10-4 10-10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ban(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M651.22 154c98 139 85 334-40 459s-318 137-458 40c-16-12-34-26-49-40c-15-15-28-32-39-49c-98-139-86-334 39-459s319-137 459-40c16 12 33 26 48 40c15 15 29 32 40 49m-522 345l370-370c-104-63-242-50-331 39c-90 90-102 228-39 331m458-280l-370 369c104 63 242 50 331-39c90-90 102-227 39-330"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bar(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h72v954H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Barcode(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 644h25V50H0zm75 0h49V50H75zm98 0h50V50h-50zm99 0h100V50H272zm149 0h75V50h-75zm124 0h50V50h-50zm99 0h25V50h-25zm75 0h49V50h-49z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M282 675H77c-40 0-75-29-77-70c60-45 98-101 123-163c24-62 33-130 38-202c4-49 36-93 80-126c25-18 52-33 80-42c0-3-2-6-2-8c0-36 29-64 65-64s65 28 65 64c0 2-2 6-2 8c28 9 56 24 81 41c44 32 79 76 79 124c4 71 15 142 40 204c25 63 63 120 121 164c-2 41-37 70-77 70H486c-10 48-51 83-101 83s-92-35-103-83m-179-70h562c-45-50-76-107-94-168c-19-60-28-124-32-190c-2-32-27-61-58-83c-30-22-66-34-96-34c-28 0-65 14-97 34c-31 20-57 50-58 80c-3 65-13 130-31 191c-20 62-52 120-96 170m282 121c5 0 7-3 7-8c0-6-2-9-7-9c-32 0-56-25-56-56c0-6-3-9-8-9c-6 0-10 3-10 9c0 40 33 73 74 73"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bicycle(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m573 226l56 60c32-20 68-33 108-33c101 0 185 83 185 185c0 101-84 184-185 184c-91 0-167-65-182-153l-83 5h-1c-7 0-12-4-17-8L284 255l-10 21c56 32 96 92 96 162c0 101-84 184-186 184C83 622 0 539 0 438c0-102 83-185 184-185c16 0 31 2 46 5l42-110l-115-31c-11-3-19-15-16-26c4-12 16-20 28-17l138 39c6 1 13 6 15 12c2 4 3 11 1 17l-10 23l221 42l29-87l-40 3h-1c-12 0-22-12-22-23s9-22 22-23l130-6c12 0 23 10 23 22c1 13-10 22-21 23l-44 2zM463 405l55-157l-215-39zm40 23l50-3c2-41 18-78 44-108l-41-44zm-337 1l48-128c-9-2-20-3-30-3c-77 0-139 62-139 140c0 77 62 139 139 139s141-62 141-139c0-51-27-96-67-120l-50 127c-5 11-18 16-30 12c-11-5-16-17-12-28m571 30l-137 8c14 63 69 110 137 110c77 0 140-62 140-139c0-78-63-140-140-140c-29 0-55 7-77 22l93 103c5 7 8 16 4 23c-3 8-11 13-20 13m-139-36l91-5l-62-68c-16 22-27 45-29 73"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bing(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M340 180c188 0 340 101 340 226S528 633 340 633C153 633 1 532 0 407V75h98v172c62-42 147-67 242-67m0 397c131 0 238-76 238-171c0-94-107-171-238-171s-238 77-238 171c0 95 107 171 238 171"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blogger(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M326 306h-73c-12 0-23-11-23-24c0-12 11-23 23-23h73c13 0 23 11 23 23c0 13-10 24-23 24m-73 98h145c13 0 23 10 23 23s-10 23-23 23H253c-12 0-23-10-23-23s11-23 23-23M98 29h454c54 0 98 44 98 98v454c0 54-44 98-98 98H98c-54 0-98-44-98-98V127c0-54 44-98 98-98m428 403v-96c0-18-5-30-27-30h-17c-18 0-28-11-28-24v-11c0-57-62-117-116-117h-85c-71 0-127 56-127 127v152c0 67 58 121 117 121h161c65 0 122-55 122-122"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bold(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 653v43c34-1 62-2 82-3c20 0 35-1 42-1c53-4 98-7 137-6l90 2c78 0 136-10 184-32c23-11 43-27 63-49c15-15 26-33 33-53c8-27 12-51 12-73c0-80-55-148-163-173c17-8 33-16 44-22s20-10 25-14c38-28 57-61 57-105c0-19-3-38-9-55c-12-35-39-64-77-82c-17-10-33-15-44-17c-30-8-59-12-87-12h-35c-7 0-13 0-17-1h-16c-2 0-5 0-8 1h-20L120 6L1 9l2 38c28 4 45 6 52 6c14 0 25 3 31 7c3 0 5 2 6 4c2 7 4 22 5 50c2 50 2 90 2 120c1 31 1 52 3 64v225c0 39-1 68-5 88c-1 7-4 14-9 22c-14 6-32 12-56 15c-12 2-23 4-32 5m247-360v-79c2-55 0-95-1-127c-2-14-2-27-2-34c24-5 43-7 59-7c52 0 91 11 117 35c26 22 39 50 39 85c0 95-56 131-162 131c-18 0-35-1-50-4m0 196V345c10-2 24-4 46-4c51-1 90 4 114 14c49 18 83 70 83 139c0 33-7 59-18 82c-12 23-29 39-55 51c-53 25-119 23-164 4c-3-8-4-14-4-19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Book(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 601V169s54-81 196-81c143 0 188 77 188 77s45-77 188-77c142 0 196 81 196 81v432s-95-70-198-70c-102 0-186 75-186 75s-84-75-186-75C95 531 0 601 0 601m65-415v314s41-31 139-31c97 0 148 58 148 58V200s-69-50-156-50c-86 0-131 36-131 36m351 14v327s48-58 148-58c99 0 139 31 139 31V186s-45-36-131-36c-87 0-156 50-156 50"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M39 14h495c40 0 71 34 71 75v506c0 50-31 85-77 85H43c-31 0-43-24-43-38V51c0-18 14-37 39-37m95 624h394c22 0 35-16 35-43V89c0-15-12-33-29-33h-47v383l-138-44l-134 44V56h-81zm144-422v37h55v56h37v-56h55v-37h-55v-56h-37v56z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Braceleft(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M289 0h-28c-29 0-51 2-63 5c-22 6-40 17-54 31c-13 14-24 32-31 56c-6 24-10 64-10 118c-1 55-1 91-2 105c-2 32-7 56-15 73c-8 16-21 30-37 40c-12 7-29 11-49 12v81c25 1 44 8 59 19c16 12 28 28 34 51c6 22 10 63 10 124c0 60 2 100 5 122c4 32 13 57 26 75c14 18 31 32 54 40c15 6 39 8 73 8h28v-77c-33 0-55-3-65-7s-17-9-21-15c-7-10-12-24-14-42c-1-11-2-49-2-114c0-72-8-122-24-154s-43-56-80-71c33-15 57-33 71-56c15-22 26-53 30-94c2-22 3-71 3-149c0-43 7-71 18-84c11-12 33-18 68-18h16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Braceright(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M289 521v-81c-25-1-45-7-61-19c-15-11-27-28-33-50c-7-22-10-64-10-125c-1-59-2-101-5-123c-4-31-13-56-27-74c-13-19-30-33-53-40c-15-6-40-9-73-9H0v79h15c30 0 50 3 60 10s17 20 21 37c3 12 4 55 4 129c0 71 9 120 24 151s42 56 81 74c-29 12-53 30-69 53s-27 55-32 95c-2 22-4 72-4 152c0 43-5 71-16 83s-34 20-69 20H0v77h27c29 0 50-1 63-5c22-6 40-16 53-30c14-13 25-33 32-57c6-24 10-63 10-118c1-55 1-90 2-105c1-32 7-56 16-73c8-16 20-30 36-40c12-7 30-11 50-11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bracketleft(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M168 0v73H72v809h96v72H0V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bracketright(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 73V0h168v954H0v-72h96V73z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Brush(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m680 28l11 12s54 54 5 104c-42 41-89 43-145 79c-67 41-103 100-95 135c15 65 92 118 25 186c-5 4-15 16-19 20c0 0 3 21-20 21c0 0 5 22-20 22L270 455l-5-5l-151-151c0-24 21-21 21-21c0-16 14-19 20-19c0 0 14-15 20-21c67-68 122 11 185 25c36 9 96-27 137-94c35-57 36-104 78-146c50-50 105 5 105 5m-86 29c-19 19-19 51 0 70s52 20 71 1s19-53 0-72s-52-18-71 1M3 453l-3-40l82-82l305 307l-81 83l-41-5l-4-48l-48-4l-3-51l-49-2l-4-49l-48-3l-5-49l-49-4l-3-49z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Buffalo(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m524 134l-41 149c-11 42-50 62-70 64c29 6 46 29 36 66l-37 138c-13 53-54 83-109 83H0L154 60h314c40 0 66 37 56 74m192 397L842 60h149L860 549c-13 49-60 85-114 85H522c-39 0-65-37-55-75L601 60h149L623 537c-3 10 4 16 10 16h54c11 0 25-7 29-22m407-196l43-159c6-26 28-36 43-36h189c9-35 24-63 57-80h-321c-38 0-93 25-109 85L895 634h148l59-219h177l43-80zm430 0l43-159c6-26 28-36 43-36h212c8-35 24-63 58-80h-345c-38 0-93 25-109 85l-130 489h149l58-219h177l43-80zm74 299l331-516c25-39 66-58 106-58h102l-57 574h-143l13-139h-167l-89 139zm659-489l-111 414c-11 40 19 75 54 75h296c-27-20-41-41-41-81h-143c-6 0-14-6-11-16l128-477h-61c-46 0-96 29-111 85m246 400l108-400c16-63 72-85 110-85h226c42 0 66 38 56 73l-111 416c-13 49-58 85-113 85h-209c-45 0-79-43-67-89M334 284l35-129c2-7-2-15-12-15h-75l-45 166h69c11 0 24-7 28-22m2450-123l-100 376c-3 10 4 16 10 16h54c11 0 25-7 29-22l100-376c2-7-2-15-11-15h-55c-13 0-24 10-27 21m-774 27l-146 227h124zM268 531l35-129c2-7-2-15-12-15h-76l-44 166h69c11 0 24-7 28-22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Building(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M41 0h525c23 0 41 18 41 41v592c0 23-18 41-41 41H431V541H176v133H41c-23 0-41-18-41-41V41C0 18 18 0 41 0m49 134h427V84H90zm0 115h427v-50H90zm0 115h427v-50H90zm0 115h427v-50H90zm200 217h-87V568h87zm114 0h-87V568h87z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bus(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 574V140c0-14 8-32 19-43c3-3 7-5 10-7h1c48-37 111-64 183-79c38-7 78-11 120-11s82 4 120 11c72 15 135 42 183 79h1c3 2 6 4 9 7c11 11 20 29 20 43v434c0 32-22 57-51 64v58c0 27-22 50-49 50c-28 0-50-23-50-50v-56H150v56c0 27-23 50-50 50c-28 0-50-23-50-50v-58c-29-7-50-32-50-64M500 75H166c-18 0-33 15-33 33c0 17 15 32 33 32h334c18 0 32-15 32-32c0-18-14-33-32-33M100 407h465c18 0 33-16 33-34V208c0-18-15-34-33-34H100c-18 0-32 16-32 34v165c0 18 14 34 32 34m36 167c29 0 52-24 52-53c0-30-23-53-52-53c-30 0-53 23-53 53c0 29 23 53 53 53m393 0c30 0 54-24 54-53c0-30-24-53-54-53c-29 0-52 23-52 53c0 29 23 53 52 53"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func C(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m507 242l-56 39c-38-51-98-84-167-84c-70 0-134 34-174 85c-26 34-41 75-41 119s15 85 41 119c40 51 104 85 174 85c69 0 129-33 167-84l56 39c-50 68-131 111-222 111c-92 0-178-44-230-112C20 514 0 459 0 401s20-113 55-158c52-68 138-112 230-112c91 0 172 43 222 111"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CupperCase(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m721 160l-63 44C600 125 505 73 397 73s-204 52-262 131c-38 51-60 115-60 184s22 132 60 183c58 79 154 130 262 130s203-51 261-130l60 41c-58 80-148 139-254 157c-155 28-305-37-391-155c-33-45-57-99-67-158c-19-106 10-211 70-294C135 83 225 25 330 7c156-28 306 36 391 153"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 155v476c0 26 22 48 49 48h552c27 0 49-22 49-48V155c0-26-22-48-49-48h-56v-5c0-17-5-36-17-50c-13-13-33-23-62-23c-30 0-50 10-62 23c-13 14-18 33-18 50v5H264v-5c0-17-5-36-17-50c-13-13-33-23-63-23c-29 0-49 10-61 23c-13 14-18 33-18 50v5H49c-27 0-49 22-49 48m66 459V510h117v104zm0-120V390h117v104zm0-120V269h117v105zm87-188v-84c0-16 10-24 31-24c22 0 32 8 32 24v84c0 16-10 24-32 24c-21 0-31-8-31-24m46 428V510h118v104zm0-120V390h118v104zm0-120V269h118v105zm134 240V510h118v104zm0-120V390h118v104zm0-120V269h118v105zm134 240V510h117v104zm0-120V390h117v104zm0-120V269h117v105zm-33-188v-84c0-16 10-24 32-24c20 0 31 8 31 24v84c0 16-10 24-31 24c-22 0-32-8-32-24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M585 114h68c46 0 84 37 84 84v364c0 46-38 84-84 84H84c-46 0-84-38-84-84V198c0-47 38-84 84-84h69c17-27 48-66 76-66h279c28 0 60 39 77 66M366 540c103 0 187-84 187-187c0-104-84-188-187-188c-104 0-189 84-189 188c0 103 85 187 189 187m3-300c62 0 112 51 112 113c0 61-50 112-112 112s-113-51-113-112c0-62 51-113 113-113"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Car(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m638 65l114 229c2 3 3 7 4 10c7 18 12 43 12 59v109c0 34-23 62-55 70v64c0 31-25 55-55 55c-31 0-55-24-55-55v-62H165v62c0 31-24 55-55 55c-30 0-55-24-55-55v-64c-32-8-55-36-55-70V363c0-16 5-41 12-59c1-3 2-7 4-10L130 65c9-17 33-33 53-33h402c20 0 44 16 53 33m-461 69l-61 139c-3 5-3 11-3 15h542c0-4 0-10-3-15l-61-139c-8-18-32-34-51-34H228c-19 0-43 16-51 34m-27 338c32 0 57-26 57-59c0-32-25-57-57-57s-59 25-59 57c0 33 27 59 59 59m468 0c32 0 59-26 59-59c0-32-27-57-59-57s-57 25-57 57c0 33 25 59 57 59"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Category(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M83 0L28 56h496L469 0zm468 276H0V83h551zm-165-55v-55H166v55h27v-28h165v28zm165 276H0V304h551zm-165-55v-56H166v56h27v-28h165v28zm165 275H0V525h551zm-165-55v-55H166v55h27v-28h165v28z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M620 0L316 436L106 280L0 427l358 273l410-588z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Checkbox(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m517 480l74-106v303H0V87h476l-51 74H74v442l443 1zm66-449L347 372L182 250L99 365l280 214l321-461z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Checkboxempty(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 677V87h591v590zm517-516H74v443h443z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chrome(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M680 200H358c-31 0-59 10-84 25c-43 27-72 75-74 129L76 139C142 55 243 0 358 0c141 0 264 82 322 200M60 159l160 278c1 1 1 2 2 3c14 25 36 45 62 59c23 11 47 19 74 19s52-8 75-19L308 714C134 689 0 540 0 359c0-74 22-143 60-200m382 66h249c16 41 26 87 26 134c0 198-161 358-359 358c-8 0-15 0-23-1l160-276c1-1 1-2 2-3c13-24 20-50 20-78v-5c-2-54-32-102-75-129M239 359c0-67 52-120 119-120s120 53 120 120s-53 120-120 120s-119-53-119-120"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cinnamon(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m793 342l26 26l-26 17l24 37l-41 4c2 0 16 32 16 32l-34-5c0 4-2 52-2 52l-19-14l2 26l-26-14l2 35l-10-9c2 14 8 42 10 43c-5 0-31-9-31-9c-2 8-6 43-6 43l-21-11l-4 46l-13-11c4 30 7 77-22 80c-28 3-54-32-61-59c-2-3-3-5-3-5c-7 10-31 31-31 31l-9-28l-16 20c0 24-6 51-26 56c-39 10-63-41-70-57v-2c-16 3-68 13-123 15c-2 19-15 37-39 37c-28 0-50-24-54-46c-4-2-8-3-12-5c-5 22-16 45-28 50c-19 8-65-7-68-31c-1-5-1-12-1-19l-37 13s4-27 7-33L7 627s13-13 23-20c2-2 4-3 6-5L0 584s22-16 32-21c0 0-25-12-32-16c0 0 25-18 34-22c1 0-33-36-33-36s14-6 24-9c-5 0-16-23-23-30c0 0 20-5 34-6l-21-33s33-1 36-1l-14-32s28 3 36 7l-12-36l29 2l-7-38s39 14 43 16c0 0-7-33 0-25c5 6 33 34 41 40c7 4 85 36 103 41c11 4 41-10 60-21c-23-6-59-19-60-25c5-9 53-11 79-52c26-42 26-85 26-85l-14-15l20-11c-5-2-28-29-31-91c-3-22-19-64-19-64l14 7l3-23l18 19c7-5 13-20 13-20s0 12 4 23c12 4 73 28 83 55s24 36 24 44c2-7 16-18 16-18l-3 19c4 0 87 10 87 10l3-20l11 10s37-63 47-85c10-23 35-42 35-42l-16 33c12-7 45-26 45-26c-8 6-22 20-20 26c13 24 33 137 15 163c32 0 9 8 9 8l-11 7c0 6 6 40 19 55c15 19 26 32 33 36c10 3 59 20 59 20zm-51 140l-2-43c12 3 26 6 29 4c-1-11-11-19-10-26c0 0 31-8 34-6c0 0-22-26-22-29c0 0 21-8 22-13c2-3-20-18-25-26c0 0 19-13 24-16c0 0-13-3-24-10c-26-7-37-31-41-39c-3-8-34-49-35-69c10-32 22-115-13-165c-5-8-15 3-15 3c-5 4-51 87-51 97c-4 0-10 0-10 5c-10-2-100-10-113 3c-5-11-35-74-70-94c-36-20-63-16-63-16c2 12 3 99 39 135c-17 5-20 11-20 11l13 9c-3 4-11 83-40 109c-20 18-52 37-52 37c3 4 22 14 39 18c9-5 15-10 15-10l-7 11c7 1 14 1 18-1c-20 14-30 24-30 24c2 2 27 8 27 8c-4 4-17 23-17 23c0 1 23 6 23 6c-2 2-10 17-10 17s23 2 34 14c1 1 0 1-1 1c6 4 11 7 14 9c-8-2-51-14-60-13c4-7 9-21 9-21l-28-8l21-24l-30-8c-23 9-53 17-77 3c-31-18-93-33-93-35c-17-16-42-20-42-20l3 32c-5 0-23 1-26-2c-2 3 9 28 9 28c2 2-32-3-32-3c0 3 16 27 21 30c-10-1-39 0-39 0c0 2 23 31 23 31c-8 0-33 6-33 6l28 29c-5 2-29 7-29 7c0 2 24 30 31 30c-17 4-31 22-31 22s28 13 31 13c-6 4-30 24-30 24l33 13c-11 8-28 24-28 24s28 13 36 10l-11 27l20-10c1-15 4-28 8-29c2-1 3 0 6 5c9 22-13 66 29 76c31 8 37-39 42-71c0 0 3-3 6-1c4 0 8 1 10 6c1 2 1 7 0 14c49 17 180 14 217 1c-3-11-4-21-3-30c3-8 14 4 14 4c-4 35 19 100 52 94c28-6 26-52 25-86l-34 21c2-1-10-18-9-36c0 0 11 15 18 21c6-5 14-9 26-21c8 12 16 30 16 30l18-20l11 23l26-25l18 23s18-26 18-29c0 0 11 18 12 19c4 0 19-35 19-35s15 12 19 11c1-7 6-36 6-34l18 7s9-36 11-41c5 2 12 2 17 5c0 0-3-2-4-32l-20-15l-25 41l-27-28s-28 34-37 36c-6-5-4-22-4-22c-7 9-33 27-33 27c0-5-6-22-6-22l-32-5c3 0 37-9 37-9l8 21l30-26s5 18 5 21c0 0 34-31 34-35c0 0 20 30 25 28c5-4 22-42 22-43c0 0 23 16 26 23c0-1 7-28 7-28s26 18 26 15l-7-33zm-58-303l-19-3s7-6 12-11c-3-2-15-5-24-4c2-6 23-12 19-16c-26 4-32 7-35 16c-1 1-1 3-2 4c-4-6-8-12-8-14c0-8 30-59 41-75c14-20 35 103 16 120c-5-4-10-7-10-7s11-6 10-10m-263-15l-14-3c-25-26-27-92-27-92s21 6 34 10c33 19 51 79 52 83l-3 2c-7-4-28-16-37-19c15 14-7 0-7 0l19 17c-1 0-26-11-26-11c2 2 9 13 9 13m212 78v-1c-4-11-3-22 1-28c3-4 7-6 11-6c6 0 14 3 22 17c10 20 4 39 0 47c-3 6-9 7-12 7c-10 0-14-12-22-36m-108-32c10 0 22 9 23 24c3 31-5 41-12 43c-2 0-4 1-6 1c-12 0-15-15-18-30c-1-5-2-10-4-14c-1-6-1-13 3-17c3-5 8-7 14-7m15 29c1-3 0-10-3-14c0 0-1 0-1-1c-13-12-3 21 1 19c1-1 2-2 3-4m195 107c7 0 15 6 16 13c1 4-2 8-6 11c-7 9-18 14-31 14c-7 0-12-1-13-1c-2-1-4-2-5-4c0-1-2-4-4-9c-2 2-4 3-6 4c-4 4-9 7-15 9c-10 4-22 6-34 6c-25 0-53-7-62-16c-3-3-8-8-6-15c2-4 7-7 11-7c3 0 8 3 17 6c9 4 21 8 29 9c4 1 8 1 12 1c14 0 29-3 38-8c3-2 5-6 7-10v-7c0-4-1-8-1-11c-4-2-8-4-11-7c-4-4-11-7-17-10c-9-5-17-10-17-18c0-2 0-7 5-11c11-8 40-8 43-8c6 0 17 0 23 3c15 6 22 8 23 16c2 7-3 13-9 19c-5 4-10 9-16 12v5c0 9 1 11 1 11c7 11 12 16 14 17l1-1c1-9 8-13 13-13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circle(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M441 0c153 0 276 124 276 276c0 143-111 262-251 274c-32 94-121 163-226 163C108 713 0 606 0 474c0-108 71-198 168-229C183 108 300 0 441 0M78 157c-43 0-78-35-78-78S35 1 78 1s77 35 77 78s-34 78-77 78m0-119c-23 0-41 18-41 41s18 41 41 41s41-18 41-41s-18-41-41-41m400 438c95-17 167-101 167-200c0-113-91-204-204-204c-98 0-180 70-198 163c130 1 235 108 235 239zm-75 0v-2c0-90-72-163-162-164c14 84 79 151 162 166M240 637c65 0 119-37 147-91c-111-23-197-110-216-221c-56 26-95 83-95 149c0 91 73 163 164 163"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clear(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M438 167V40c0-23-17-40-39-40h-30c-22 0-39 17-39 40v127c0 22 17 40 39 40h30c22 0 39-18 39-40m-169 24l-90-89c-15-15-41-15-56 0l-21 21c-15 15-15 41 0 56l89 90c15 15 42 15 57 0l21-21c15-15 15-42 0-57m308 78l89-90c15-15 15-41 0-56l-21-21c-15-15-41-15-56 0l-90 89c-15 15-15 42 0 57l21 21c15 15 42 15 57 0M40 438h127c22 0 40-17 40-39v-31c0-22-18-38-40-38H40c-23 0-40 16-40 38v31c0 22 17 39 40 39m561 0h127c23 0 40-17 40-39v-31c0-22-17-38-40-38H601c-22 0-40 16-40 38v31c0 22 18 39 40 39M179 665l90-89c15-15 15-41 0-56l-21-22c-15-15-42-15-57 0l-89 91c-15 15-15 41 0 56l21 20c15 15 41 15 56 0m487-76l-89-91c-15-15-42-15-57 0l-21 22c-15 15-15 41 0 56l90 89c15 15 41 15 56 0l21-20c15-15 15-41 0-56M438 728V601c0-22-17-40-39-40h-30c-22 0-39 18-39 40v127c0 23 17 40 39 40h30c22 0 39-17 39-40"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clip(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m175 397l280-280c3-3 27-27 65-40c53-17 107-3 150 40c44 43 57 96 40 149c-12 38-37 64-40 67L369 634c-52 52-182 119-302 0c-119-119-51-250 0-301L390 9c12-12 31-12 43 0s12 31 0 43L111 376c-5 4-106 109 0 215c103 103 204 11 215 0l301-302s17-17 25-41c10-32 3-60-25-88c-60-60-118-10-129 0L218 439c-10 10-17 27 0 44s33 9 43 0l194-194c12-11 32-11 43 0c12 12 12 32 0 44L304 526c-27 26-82 47-129 0c-48-47-27-103 0-129"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M343 620H171C77 620 0 543 0 449c0-71 43-130 103-156c0-5-1-10-1-15c0-56 47-102 103-102c34 0 64 17 83 42c31-84 112-145 207-145c122 0 220 99 222 220c60 27 102 86 102 156c0 94-77 171-171 171H481V520h81c12 0 16-8 7-17L428 352c-4-4-10-7-16-7s-13 3-17 7L254 503c-8 9-6 17 7 17h82z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Code(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m159 345l172 173l-83 80L0 350L254 96l79 78zm450 4L437 176l83-80l248 248l-254 254l-79-78z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coffee(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M397 344c-3-67-7-103-47-147c-51-57-70-112-28-197c11 57 15 108 55 139s73 115 20 205m-75 0c-2-36-4-54-26-79c-27-30-38-59-15-105c6 31 7 58 29 74s41 62 12 110m181 316h182c11 0 15 8 8 17c0 0-71 91-169 91H173L4 677c-7-9-4-17 7-17h180c-25-13-62-70-62-141V369h438v12c14-12 31-18 49-18c25 0 72 26 72 71c0 65-43 89-68 106c0 0-47 37-80 80c-11 19-24 34-37 40m113-270c-16 0-36 7-49 39v83c0 11-1 24-3 35c22-22 48-37 48-37c16-11 48-32 48-73c1-29-26-47-44-47m-411 12h-28v41h28zm0 65h-28v67s8 90 65 117c-35-41-37-88-37-117z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func College(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m390 86l390 156l-390 155L99 281c-18 15-31 39-33 66h7c8 0 14 6 14 14v37c0 8-6 14-14 14h-5l19 194c1 8-5 15-13 15H39c-8 0-14-7-14-15l20-194h-5c-8 0-15-6-15-14v-37c0-8 7-14 15-14h11c2-29 15-54 33-72L0 242zm-11 339l11 4l11-4l214-86l16 177c1 12-7 17-18 12l-84-43c-11-6-29-5-39 1l-81 46c-10 6-28 6-38 0l-81-46c-10-6-28-7-38-1l-85 43c-11 6-19 0-18-12l16-177z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Colon(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 145h77v77H0zm0 427h77v77H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comma(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m84 538l70 33l-94 196l-60-29z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comment(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M80 9h557c44 0 80 36 80 80v338c0 44-36 80-80 80H531l5 137c0 33-18 41-42 18L337 507H80c-44 0-80-36-80-80V89C0 45 36 9 80 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comments(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M227 507L97 594s27-72 29-124C49 429 0 362 0 287C0 162 139 60 310 60c142 0 262 71 299 166c-14-2-28-3-42-3h-14c-70 3-135 26-185 65c-55 44-85 106-83 168c1 19 5 40 12 58c-24 0-47-3-70-7m481 127l-85-46c-17 4-37 7-55 8h-12c-109 0-197-63-200-143c-4-83 86-154 200-158h11c109 0 198 61 201 142c2 49-28 93-77 123c2 35 17 74 17 74"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compass(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M358 0c198 0 359 161 359 359S556 717 358 717S0 557 0 359S160 0 358 0m0 634c152 0 276-123 276-275S510 83 358 83S83 207 83 359s123 275 275 275m59-217l-214 97l97-214l214-97zm-175 59l155-78l-78-78z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cookpad(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M358 0c198 0 359 161 359 359S556 717 358 717S0 557 0 359S160 0 358 0m0 680c177 0 322-144 322-321S535 37 358 37S37 182 37 359s144 321 321 321m104-482c-16-39-55-69-105-69c-40 0-79 27-93 52c-79-36-161 28-161 103c0 52 31 87 75 109c-5 45-13 97-18 141c-2 15-6 33 4 45c25 6 52 3 80 3h244c17-8 14-29 16-43c5-45 12-95 18-138c73 2 123-46 123-114c0-89-100-147-183-89m-29 29c0 9-6 11-6 18c9 6 20 14 32 18c10-30 44-57 85-51c40 5 66 46 63 85c-3 33-34 66-82 67c3-20 4-41-15-42c-20 0-22 31-25 58c-4 29-5 57-8 74H279c-28 0-57-4-53 24c3 16 29 13 53 13c62 0 135-2 193 1c-2 18-4 37-7 53H197c8-60 15-124 23-184c-45-2-84-36-79-87c8-69 106-88 141-32c2-38 30-71 65-75c41-5 76 21 86 60"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copy(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m572 22l105 109c10 11 19 34 19 49v323c0 15-12 29-28 29H481v-49h157c5 0 9-4 9-9V202c0-5-4-9-9-9H522c-5 0-10-4-10-9V60c0-5-4-9-9-9H313c-5 0-10 4-10 9v69c-10-4-23-6-33-6h-15V30c0-16 13-28 28-28h242c15 0 37 9 47 20m-2 122h42c5 0 6-3 3-7l-48-49c-3-4-6-3-6 2v45c0 5 4 9 9 9m-252 38l104 110c11 11 19 33 19 48v323c0 16-12 29-27 29H29c-16 0-29-13-29-29V190c0-15 13-29 29-29h241c15 0 37 10 48 21M58 643h325c5 0 10-4 10-9V362c0-5-5-9-10-9H267c-5 0-9-4-9-9V220c0-5-4-9-9-9H58c-5 0-9 4-9 9v414c0 5 4 9 9 9m302-345l-47-51c-4-4-7-1-7 4v44c0 5 5 9 10 9h41c5 0 7-3 3-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crop(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M591 526h126v65H591v126h-66V591H126V192H0v-66h126V0h65v126h374l74-73l35 35l-83 83zm-400-26l309-308H191zm334-263L237 526h288z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crown(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 185c-44 0-80-36-80-80c0-45 36-80 80-80s80 35 80 80c0 44-36 80-80 80M65 152c35 0 63 29 63 64c0 36-28 64-63 64c-36 0-65-28-65-64c0-35 29-64 65-64m640 0c35 0 63 29 63 64c0 36-28 64-63 64c-36 0-65-28-65-64c0-35 29-64 65-64M135 594L82 304c16-3 31-12 43-22c33 30 75 59 115 59c48 0 87-74 114-136c9 3 20 6 30 6s21-3 30-6c27 62 66 136 114 136c40 0 82-29 115-59c12 10 27 19 43 22l-53 290zm-11 25h520v49H124z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cursor(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M298 312v42h42v42H213v85h42v85h43v85h-85v-85h-43v-85h-42v-42H85v42H43v43H0V57h43v42h42v43h43v42h42v43h43v42h42v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cut(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M184 396h-4c-42 0-87 20-121 53c-35 32-59 75-59 118c0 49 35 79 83 79c41 0 87-21 122-53c34-33 58-75 58-118c0-5 0-12-1-17l32-28l314 161c5 2 10 4 15 4c7 0 20-4 34-6c13-3 27-6 35-7c14-2 25-15 25-29c0-10-5-19-14-25L442 347l261-181c9-6 14-15 14-25c0-14-11-28-25-30c-4 0-9-1-15-3c-12-3-28-5-40-7s-21-2-29 2c-2 1-8 4-19 9c-10 5-23 12-38 20s-32 16-50 26c-18 9-39 19-57 29c-38 19-75 38-103 52c-14 7-26 14-35 18c-8 4-12 7-12 7l-32-29c1-5 1-11 1-16c0-43-24-86-58-118c-35-32-81-53-122-53c-48 0-83 30-83 79c0 42 24 85 60 118c36 34 82 56 124 53c4 6 13 14 22 22s17 17 23 23v8c-6 6-13 13-23 22c-9 8-17 17-22 23M74 127v-3c3-1 6-1 9-1c24 0 50 15 67 29c8 7 18 19 25 31c8 12 15 26 15 37v2c-5 1-8 1-10 1c-24 0-48-15-66-30c-7-7-18-17-26-29s-14-26-14-37m145 276l41-40l-2-48l365-186l64 12l-380 262l-17-9l-40 36c-7-11-18-20-31-27m13-100l-13-14c13-5 24-14 31-25l16 14c-15 7-28 11-34 25m46 44c0 16 13 29 29 29c17 0 30-13 30-29c0-17-13-30-29-30c-18 0-30 13-30 30m60 71l77-53l272 188l-64 12zM75 567c0-22 15-47 37-65s48-32 68-32c3 0 6 1 8 1v2c1 0 1 1 1 2c0 22-15 46-37 64s-47 32-69 32h-4c-1-1-2-1-4-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func D(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M490 324V0h72v754h-72v-82c-53 58-132 96-216 96C122 768 0 648 0 498s122-270 274-270c84 0 163 38 216 96M275 702c115 0 215-91 215-204S390 294 275 294c-114 0-207 91-207 204s93 204 207 204"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DupperCase(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h254c82 0 157 26 218 71c90 67 149 174 149 296c0 121-59 228-149 296c-61 45-136 72-218 72H0zm71 663h189c161-4 290-135 290-296c0-162-129-293-290-296H71z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dailycalendar(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M558 94h57c28 0 51 22 51 49v486c0 28-23 51-51 51H50c-28 0-50-23-50-51V143c0-27 22-49 50-49h58v-6c0-17 5-36 18-50c12-13 32-24 62-24c31 0 52 11 65 24c12 14 17 34 17 50v6h125v-6c0-17 6-36 19-50c12-13 32-24 63-24c30 0 51 11 64 24c12 14 17 34 17 50zm-337 81V88c0-16-10-24-33-24c-22 0-31 8-31 24v87c0 16 9 23 31 23c23 0 33-7 33-23m288 0V88c0-16-12-24-32-24c-23 0-33 8-33 24v87c0 16 10 23 33 23c22 0 32-7 32-23M68 612h530V259H68zm93-220h50c0-13 3-24 10-33c6-8 17-13 30-13c9 0 18 3 25 9s12 15 12 26c0 7-2 13-6 17c-3 4-9 8-14 10s-11 4-17 4c-7 1-14 1-19 0v37c6 0 14 0 21 1c6 0 13 1 18 4c6 2 12 6 16 12c3 5 5 12 5 21c0 13-4 23-12 30c-8 8-18 11-30 11c-7 0-15-1-20-4s-10-6-14-10c-3-5-6-11-8-17s-3-13-3-20h-50c0 15 2 29 7 40c4 12 10 21 19 29c8 8 19 15 30 19c12 4 26 6 40 6c12 0 25-2 36-6c11-3 21-8 30-15s17-17 22-27c5-11 8-23 8-37s-4-27-12-38c-8-10-20-17-34-20v-1c12-3 21-10 27-20c6-9 9-21 9-33c0-11-2-21-7-30s-13-17-21-23s-17-10-27-13s-21-5-31-5c-13 0-26 2-37 6c-10 5-20 11-28 18c-7 8-13 18-18 29c-4 10-6 23-7 36m284 7v167h53V308h-42c-1 10-5 18-10 24c-4 7-10 13-16 17c-7 4-14 7-23 8c-8 2-17 2-26 2v40z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dark(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M312.297 99v7c0 27-12 40-37 40s-36-13-36-40v-7c0-26 11-39 36-39s37 13 37 39m74 32l5-6c15-22 33-25 52-11c20 14 23 33 8 55l-6 5c-15 22-33 25-52 10c-20-14-22-31-7-53m-226-6l4 6c15 22 13 39-6 53c-20 15-39 12-54-10l-4-5c-15-22-13-41 6-55c20-14 39-9 54 11m115 66c85 0 155 70 155 156c0 85-70 154-155 154c-86 0-155-69-155-154c0-86 69-156 155-156m-224 44l7 2c26 8 34 23 27 47c-8 24-24 31-49 23l-7-2c-25-8-34-23-27-47c8-24 24-31 49-23m441 2l7-2c26-8 42-1 49 23s-1 39-27 47l-6 2c-26 8-42 1-49-23c-8-24 1-39 26-47m-217 196c47 0 86-39 86-86c0-48-39-87-86-87c-48 0-87 39-87 87c0 47 39 86 87 86m-246-45l7-2c25-8 41-2 49 22c7 24-1 39-27 47l-7 2c-25 8-41 2-49-22c-7-24 2-39 27-47m486-2l6 2c26 8 34 23 27 47s-23 30-49 22l-7-2c-25-8-34-23-26-47c7-24 23-30 49-22m-415 139l4-6c15-22 34-24 54-10c19 14 21 31 6 53l-4 6c-15 22-34 24-54 10c-19-14-21-31-6-53m345-6l6 6c15 22 12 39-8 53c-19 14-37 12-52-10l-5-6c-15-22-13-39 7-53c19-14 37-12 52 10m-133 67v8c0 27-12 40-37 40s-36-13-36-40v-8c0-27 11-40 36-40s37 13 37 40"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dashboard(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M105 646h506c9 0 18-4 23-11c27-30 47-68 61-107c14-40 22-82 22-122c0-98-40-187-105-253C546 87 457 48 358 48S170 87 104 153C39 219 0 307 0 406c0 40 8 82 22 122c14 39 33 77 61 107c5 7 12 11 22 11m208-493c0-26 20-45 45-45s45 19 45 45c0 25-20 44-45 44s-45-19-45-44m-179 74c0-25 20-44 45-44s45 19 45 44s-20 45-45 45s-45-20-45-45m359 0c0-25 20-44 45-44c26 0 45 19 45 44s-19 44-45 44c-25 0-45-19-45-44M355 452l52-186c3-9 11-16 22-16s22 10 22 21c0 2-2 7-4 16c-3 10-5 22-9 34c-3 13-7 28-11 42c-9 31-18 60-23 80c-3 10-5 19-6 22c20 14 35 36 35 61c0 42-33 75-75 75s-74-33-74-75c0-19 7-38 20-51s32-22 51-23M60 406c0-25 19-44 44-44s46 19 46 44s-21 45-46 45s-44-20-44-45m507 0c0-25 20-44 45-44s44 19 44 44s-19 45-44 45s-45-20-45-45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Delicious(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M303 322H0V19h303zm352 0H352V19h303zM272 50H31v242h241zm31 624H0V371h303zm49 0V371h303v303zm47-272h-16v16zm41 0l-57 57v41l98-98zm82 0L383 541v41l180-180zM383 644h19l223-222v-20h-21L383 623zm101 0l141-141v-40L443 644zm82 0l59-59v-41L525 644zm41 0h18v-18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Digg(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M132 410H90V297h42zm0-191H0v270h222V104h-90zm131-37h90v-78h-90zm0 307h90V219h-90zm264-79h-42V297h42zM395 525v79h222V219H395v270h132v36zm395-115h-42V297h42zM658 525v79h222V219H658v270h132v36z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dollar(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M263 148v256c65 47 107 89 131 125c24 37 35 73 36 110c2 60-44 175-167 195v119h-71V836C82 822 28 733 0 682l63-37c38 71 81 109 129 119V445c-56-43-93-75-107-94c-26-35-39-73-39-113c0-65 41-155 146-170V0h71v69c79 15 125 77 155 117l-60 45c-22-29-51-68-95-83m-71 202V145c-54 14-70 63-70 90c0 46 29 78 47 95c3 4 12 11 23 20m71 152v257c65-22 88-83 87-118c0-25-8-48-24-72c-14-19-35-42-63-67"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Down(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m356 357l246-246l115 119l-355 354L0 222l112-113z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M586 271H457V60c0-18-15-33-33-33H295c-19 0-34 15-34 33v211H132c-17 0-23 11-10 24l214 213c6 6 14 9 23 9s18-3 24-9l213-213c12-13 8-24-10-24M0 422v228c0 10 5 16 16 16h684c11 0 17-6 17-16V422c0-10-8-17-17-17h-65c-9 0-17 8-17 17v145H98V422c0-10-7-17-16-17H16c-9 0-16 8-16 17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dribbble(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M358 0c198 0 359 161 359 359S556 717 358 717S0 557 0 359S160 0 358 0m0 53c-25 0-48 3-71 9c4 5 61 82 113 178c116-43 160-110 160-111c-53-48-124-76-202-76M227 82C142 122 78 201 59 295c7 1 136 2 282-37c-52-94-110-171-114-176m368 82c-1 1-50 72-171 121c7 15 14 30 20 45c2 5 6 10 8 15c104-13 208 10 213 11c-1-73-27-140-70-192M52 359c0 79 30 151 79 205c0 0 83-150 249-203c4-1 8-2 11-3c-7-17-15-35-24-52c-158 47-309 44-315 44zm477 254l1-1c69-46 117-120 130-205c-6-2-92-28-190-13c40 111 56 201 59 219m-52 27h1c-5-27-22-120-65-232c-1 0-1 1-2 1c-174 60-236 181-240 190l6 7c51 37 113 59 181 59c42 0 83-9 119-25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dropbox(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M682 151L552 263l130 113l-217 135l-124-111l-124 111L0 376l130-113L0 151L216 14l125 112L466 14zM140 263l201 127l201-127l-201-127zm201 181l122 108l97-61v65L341 694L122 556v-65l97 61z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dropdown(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m602.442 200l-253 317c-24 29-61 29-84 0l-253-317c-24-30-12-53 25-53h540c38 0 49 23 25 53"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func E(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M559 399H69v2c0 113 100 204 215 204c82 0 152-47 186-115l61 29c-44 90-137 152-246 152C133 671 0 551 0 401c0-42 11-82 29-117c47-91 147-153 256-153c130 0 238 89 267 208c4 19 7 40 7 60M80 339h400c-27-82-104-142-196-142c-82 0-158 47-194 116c-4 8-7 17-10 26"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EupperCase(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M455 0v71H72v209h383v72H72v311h383v72H0V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Edit(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M558 554V425l80-80v235c0 47-39 86-85 86H85c-47 0-85-39-85-86V113c0-47 38-86 85-86h475v1l-80 79H112c-17 0-32 15-32 33v414c0 18 15 32 32 32h414c17 0 32-14 32-32m76-488l85 85l39-39l-85-85zM335 366l84 85l271-271l-84-85zm-60 144l116-31l-85-85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eight(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M343 349c67 40 112 113 112 196c0 126-102 227-227 227C102 772 0 671 0 545c0-83 44-156 111-196c-47-35-78-91-78-154C33 87 120 0 228 0s193 87 193 195c0 63-30 119-78 154M105 195c0 68 55 123 123 123s122-55 122-123S296 72 228 72s-123 55-123 123m123 507c86 0 155-71 155-157s-69-156-155-156S71 459 71 545s71 157 157 157"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eject(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.483 380l274-343c26-31 66-31 91 0l274 343c26 32 14 57-27 57h-585c-40 0-53-25-27-57m37 308h565c28 0 51-22 51-51v-67c0-28-23-50-51-50h-565c-28 0-50 22-50 50v67c0 29 22 51 50 51"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Emdash(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M643 248v73H0v-73z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Emphasis(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M525 594s34 4 35 4c-4 43-27 110-40 132c-51-1-143-2-222-2H173c-46 0-101 1-173 2v-29l53-5c44-6 49-11 49-87V301c0-75-5-80-49-88l-30-4v-29c45 2 101 3 147 3h203c62 0 116-1 129-3c1 13 7 72 12 127l-36 4c-10-29-23-53-36-65c-18-15-47-25-94-25h-77c-31 0-31 1-31 32v144c0 24 1 25 27 25h66c50 0 61-6 71-45l5-21s35 0 35 1c-2 26-3 56-3 87c0 32 1 62 3 88c0 1-36 1-35 1l-5-21c-10-39-21-46-71-46h-66c-26 0-27 1-27 26v98c0 38 4 64 15 78c13 15 31 22 101 22c83 0 116-3 169-96M340 70c0-39-31-70-70-70c-38 0-70 31-70 70c0 38 32 70 70 70c39 0 70-32 70-70"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Endash(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 248v73H0v-73z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Equal(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M506 149v72H0v-72zm0 219v73H0v-73z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eraser(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M68 276L322 22c23-23 58-23 81 0l246 246c23 23 23 59 0 82L395 604c-90 90-236 90-327 0c-90-90-90-238 0-328m286 287L109 317c-68 68-68 178 0 246s177 68 245 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Etc(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M176 347c0-48-40-88-88-88c-49 0-88 40-88 88s39 88 88 88c48 0 88-40 88-88m270 0c0-48-40-88-88-88s-88 40-88 88s40 88 88 88s88-40 88-88m271 0c0-48-39-88-88-88c-48 0-88 40-88 88s40 88 88 88c49 0 88-40 88-88"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Evernote(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M65 164h69c15 0 28-13 28-28c0 0-1-14-1-31V61c0-13 3-24 9-32c8-12 23-20 37-24c16-5 82-9 125 8c16 6 31 27 34 47c25-1 65-1 100 3c44 5 78 11 95 17c16 5 34 20 39 49c9 50 21 250 15 317c-10 104-25 167-32 187c-33 103-62 107-143 107c-102 0-134-16-134-99c0-91 43-93 116-91c10 1-1 11-1 29c0 17 7 24-3 24c-20 0-49-3-49 26c0 35 17 34 61 34c55 0 66-8 66-58c0-82-23-94-56-98c-35-4-70-9-87-15c-42-15-39-72-39-87c0-3-4-2-4 0c0 23-1 53-13 89c-3 10-6 16-6 16c-13 29-36 18-72 14c-36-5-111-20-145-34c-14-7-20-14-28-30C27 419 6 285 3 266c-3-25-3-36-3-36c0-16 2-31 11-44c4-6 11-12 20-16c8-3 20-6 34-6m66-20H61c-16 0-31 3-43 8c0 1-4 3-7 4l1-1L146 22h1l-2 4c-5 9-8 22-8 35c0 17 1 75 1 75c0 4-3 8-7 8m329 206c18 0 36 5 49 11c0-20-4-53-37-54c-30-1-39 24-41 44c9-1 19-2 29-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Exchange(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M696.033 490h81c24 0 28 14 11 30l-118 119c-10 11-29 11-40 0l-119-119c-16-16-12-30 12-30h81V196c-4-41-38-72-80-72c-43 0-78 33-80 75v296c-3 62-40 117-92 145c-24 12-50 18-79 18s-54-6-78-18c-53-28-91-83-93-146V204h-80c-24 0-29-13-12-29l118-119c11-11 30-11 40 0l119 119c17 16 12 29-12 29h-80v291c2 42 35 76 78 76s79-36 79-79V199c3-63 39-116 92-144c24-12 51-19 80-19s55 7 80 19c51 27 87 79 92 140z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Exclam(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 0h73v570H2zM0 652h77v77H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func External(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M501 107L167 443l70 71l335-335l62 61c17 17 32 12 32-14V45c0-16-14-31-31-31H454c-26 0-32 15-15 32zm37 465V345h76v253c0 45-37 82-82 82H82c-45 0-82-37-82-82V147c0-45 37-82 82-81h277v76H108c-16 0-31 15-31 31v399c0 16 15 31 31 31h399c16 0 31-15 31-31"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func F(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M255 5v78c-11-7-25-11-39-11c-39 0-70 31-71 70v86h110v60H145v466H73V288H0v-60h73v-84c0-53 29-99 72-124c22-12 45-20 71-20c14 0 27 2 39 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FupperCase(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M455 0v71H72v209h383v72H72v383H0V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Facebook(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M38 14h590c20 0 38 18 38 38v590c0 20-18 38-38 38H38c-20 0-38-18-38-38V52c0-20 18-38 38-38m435 374h92l4-88h-96v-65c0-25 5-39 37-39h56l2-82s-25-4-61-4c-88 0-127 55-127 114v76h-65v88h65v244h93z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Female(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M245.116 76c0-42-34-76-76-76c-41 0-75 34-75 76s34 76 75 76c42 0 76-34 76-76m42 124l50 157c5 13-3 27-16 31c-13 5-27-3-32-16l-27-83c-4-13-16-24-27-24c-10 0-15 11-11 24l58 179c4 13-4 24-18 24h-25v203c0 14-12 25-25 25c-14 0-26-11-26-25V492h-38v203c0 14-11 25-25 25s-25-11-25-25V492h-25c-14 0-22-11-18-24l58-179c4-13-1-24-12-24c-10 0-22 11-27 24l-27 83c-4 13-18 21-31 16c-14-4-21-18-17-31l51-157c4-13 19-24 33-24h169c14 0 28 11 33 24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func File(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m430 28l141 148c14 15 26 43 26 65v438c0 20-16 38-38 38H38c-20 0-38-18-38-38V38C0 18 18 0 38 0h328c22 0 50 13 64 28M79 651h439c7 0 13-6 13-13V272c0-7-6-13-13-13H361c-7 0-13-6-13-13V79c0-7-5-12-12-12H79c-7 0-12 5-12 12v559c0 7 5 13 12 13m408-467l-63-68c-5-5-9-3-9 4v61c0 6 6 12 12 12h56c7 0 9-4 4-9m-331-58h124c13 0 25 12 25 26v15c0 13-12 25-25 25H156c-14 0-26-12-26-25v-15c0-14 12-26 26-26m0 134h124c13 0 25 12 25 26v15c0 13-12 25-25 25H156c-14 0-26-12-26-25v-15c0-14 12-26 26-26m283 200H156c-14 0-26-11-26-25v-15c0-14 12-25 26-25h283c13 0 26 11 26 25v15c0 14-13 25-26 25m0 137H156c-14 0-26-11-26-24v-16c0-14 12-25 26-25h283c13 0 26 11 26 25v16c0 13-13 24-26 24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Firefox(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M444.255 752h-1c-1 1-1 2-2 2c-3 6-10 10-17 10c-17-1-37-2-55-5c-2 0-5-1-7-1c-3 3-8 5-13 5h-3c-23-4-48-12-72-22c-19-8-37-16-53-26c-57-30-107-74-144-128c-27-37-46-80-57-124c-9-1-16-9-16-17c-1-13-2-27-2-40c0-19 1-39 3-56c-4-4-6-10-5-16c4-35 15-70 29-102c5-13 11-25 18-36c-1-8-1-15-1-23c0-42 14-82 42-114c3-4 8-6 14-6h6c7 2 12 8 13 15c2 14 7 28 15 39c67-67 159-107 260-107c60 0 117 14 167 40c43 16 82 41 114 74c1-2 2-5 4-6c3-4 8-6 13-6c4 0 9 1 12 3c22 16 40 40 52 69c13 29 21 61 23 96c0 4 1 8 1 13c6 2 12 8 13 15c3 15 4 32 4 47c0 50-11 98-34 142c-2 6-4 13-6 18c4 4 7 9 7 15v3c-6 48-48 101-106 137c-6 4-12 8-18 11c1 6-1 12-5 16c-31 34-89 57-157 64c-12 1-25 1-36 1m130-249l-21 28c-4 3-9 5-13 7c-1 0-2 0-3-1c-3-3-2-10 1-17c-4 3-10 7-15 11c-10 8-18 17-24 25c-3 1-7 2-10 2c3-3 6-6 8-10c-18 9-43 15-67 15h-4c-6-1-12-1-17-2c-19-3-39-9-56-18c0 0-1 0-2-1c5-1 10-1 15-1s9 0 13 1c-9-4-21-8-33-8h-5c-1 0-1 0-2 1c-47-28-83-72-88-119c-3-27 0-50 12-73c-7-1-14-2-22-2c-12 0-25 2-36 5c-5 1-8 2-11 3c-1-3-4-6-5-9c16-6 35-9 52-9c10 0 20 1 29 2c1 1 3 1 4 1l1 1h5c0 1 0 1 1 1s2 0 3 1h1c1 0 2 1 3 1h1c1 1 2 1 3 1c0 1 1 1 1 1c1 0 2 1 3 1c0 0 1 0 1 1c1 0 2 0 3 1h1c2 1 3 1 4 2c3-7 4-15 4-23v-3c0-8-2-16-5-23c-2-4-5-9-7-13c-12-19-18-42-18-68c0-47 23-90 60-116c-5 0-10-1-14-1c-37 0-69 17-89 45c-12-3-34-5-45-5c-48 0-91 22-120 58c-10 13-19 28-25 43c-13 30-24 63-28 98c4-12 10-24 16-34c-8 28-12 63-12 103c0 12 0 26 1 38c1-16 4-32 6-45c11 131 89 243 202 299c14 9 33 18 51 26c24 9 47 16 67 19c-10-4-21-9-31-14c5 1 11 2 17 3c11 3 24 6 37 8c18 3 36 4 52 4c-5-2-11-4-17-7c12-1 24-2 35-3h2c11 0 23-1 34-2c65-6 119-29 146-58c-14 7-32 13-50 17c16-9 34-20 48-31c9-5 19-10 28-16c55-34 92-82 98-124c-8 14-21 29-38 42c15-26 29-53 38-82c20-41 31-87 31-135c0-15-1-30-3-44v6c0 22-4 41-11 58c-1-15-2-32-4-46c2-15 3-32 2-48c-4-67-31-121-69-150c19 20 32 51 32 86c0 3 0 5-1 8c-38-85-116-150-210-169c53 27 98 70 127 122c-19-18-43-34-71-41c46 42 76 101 78 169c0 8 0 17-1 25c-7-18-20-34-33-47c4 22 8 44 8 69c1 35-3 68-10 99c-3-15-7-30-12-42c0 4 0 8-1 11c-2 36-12 68-26 92m-472-431c-24 28-37 62-37 101v2l4-4s1-1 1-2c2-1 3-2 4-3l1-1s2-2 4-3v-1c2-1 3-2 5-4c2-1 3-3 5-4v-1c5-3 11-7 16-10c8-4 17-8 26-12c-3-2-5-5-8-8c0-1 0-1-1-2c0 0 0-1-1-1c0-1-1-1-1-2l-1-1c0-1-1-2-1-2c0-1-1-2-1-3l-1-1c0-1-1-1-1-2s-1-2-1-2c0-1-1-2-1-2c0-1-1-1-1-2s0-1-1-2v-1c-1-1-1-2-1-2c-1-1-1-1-1-2s-1-1-1-2v-1c-1-1-2-3-2-4c0 0-1 0-1-1v-2c-1-1-1-1-1-2v-2l-1-1v-3s0-1-1-1zm288 144c-1-15-47-11-65-11c-9 0-19-5-26-10c-1 6-2 14-2 20c0 20 5 41 15 58c3 4 5 8 7 12c16-11 40-23 48-29c13-10 24-24 23-40m36 191c-47 0-47 21-90 21c-22 0-51-11-68-38c0 7 0 16 1 24c4 32 26 65 58 90c20-1 42-6 53-10c30-10 53-33 70-39c7-3 21 5 28-9c6-11-20-39-52-39"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Five(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m203 72l-36 192c20-6 42-9 65-9c137 0 248 112 248 249S369 752 232 752C115 752 27 672 0 562h73c25 71 81 122 160 122c99 0 179-80 179-179c0-98-80-180-179-180c-26 0-53 6-81 14c-30 8-51 14-78 27L144 0h317v72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M108 699V134c23-14 35-36 35-62c0-39-32-72-71-72C32 0 0 33 0 72c0 26 14 48 37 62v565c0 10 7 18 16 18h37c10 0 18-8 18-18m35-533v324c0 19 10 27 22 18c23-17 43-30 61-39c37-17 63-25 87-26c30 0 53 6 76 17c22 10 42 26 65 41c13 8 29 13 47 14c30 2 72-5 121-45c11-9 20-32 20-49V97c0-19-8-26-20-17c-25 19-48 32-68 38c-41 11-73 9-100-7c-23-15-43-31-65-41c-23-11-46-17-76-17c-11 0-25 3-40 7c-27 8-62 25-108 57c-12 9-22 31-22 49"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flickr(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M84 14h498c47 0 84 37 84 84v498c0 25-10 44-25 59c-14 14-36 25-59 25H84c-47 0-84-37-84-84V98c0-23 11-45 25-59c15-15 34-25 59-25m111 443c62 0 113-50 113-112s-51-113-113-113c-61 0-113 51-113 113s52 112 113 112m276 0c61 0 113-50 113-112s-52-113-113-113c-62 0-114 51-114 113s52 112 114 112"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M380 146h305c17 0 32 15 32 32v423c0 17-15 31-32 31H32c-17 0-32-14-32-31V178c0-17 15-32 32-32h34l16-53c0-17 15-32 32-32h218c17 0 32 15 32 32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Forward(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M494 38v124C-36 162 1 656 1 656c104-345 493-245 493-245v127l274-254z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Four(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M398 0v494h89v72h-89v209h-71V566H0zM138 494h189V227z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Foursquare(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m356 538l156-155c11-11 12-29 3-42l151-127v466H0V14h554L400 226l-44-44c-12-12-33-12-45 0l-9 9c-25-40-70-66-120-66c-78 0-140 62-140 140c0 76 60 137 135 140l134 133c12 13 33 13 45 0m-30-88l-86-68c24-12 44-32 57-56c15 20 26 34 27 35c1 2 4 4 7 4s6-1 9-5L592 14h74v160L337 451c-2 1-4 2-5 2c-4 0-6-3-6-3M182 155c60 0 110 49 110 110c0 60-50 109-110 109S73 325 73 265c0-61 49-110 109-110"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Friend(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M552.664 491c41 23 64 61 60 100c-2 24-1 24-32 28c-19 2-140 3-263 3c-140 0-286-1-297-4c-41-11-17-87 37-123c43-27 130-69 154-74c33-7 37-28 0-91c-8-15-17-58-18-104c-1-74 12-124 77-149c12-4 26-6 39-6c43 0 83 24 100 59c23 47 12 170-12 215c-29 51-26 67 6 75c20 5 86 36 149 71m219-63c32 17 50 48 47 78c-1 18-1 19-25 22c-12 2-75 3-149 3c-13-31-38-59-72-77c-39-22-85-45-120-60c23-11 43-20 52-22c25-6 28-20 0-70c-7-12-15-46-16-82c-1-58 12-98 62-117c10-3 21-5 31-5c33 0 64 18 77 46c17 37 9 133-9 169c-22 40-20 52 5 58c16 4 67 29 117 57"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Frustrate(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M358 0c198 0 359 155 359 349S556 698 358 698S0 543 0 349S160 0 358 0M120 311l8 21c3 7 11 12 19 9l175-63c7-3 9-8 10-21c2-13 0-19-7-23l-160-92c-7-4-17-2-21 5l-11 20c-4 7-2 16 5 20l108 62l-117 42c-7 3-11 12-9 20m469 21l8-21c2-8-2-17-9-20l-117-42l108-62c7-4 9-13 5-20l-12-20c-4-7-13-9-20-5l-160 92c-7 4-9 10-7 23c1 13 3 18 10 21l174 63c8 3 17-2 20-9m-59 225c20-7 31-32 24-52c-32-81-109-133-196-133s-163 52-195 133c-7 20 3 45 23 52c20 8 325 8 344 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Full(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m338 272l70 69c4 4 8 6 14 6c5 0 9-2 13-6l119-118l72 72c23 22 40 13 40-18V51c0-19-16-37-38-37H402c-32 0-40 17-17 40l72 72l-119 118c-3 4-5 9-5 14c0 6 2 10 5 14M0 416v227c0 19 16 37 38 37h225c32 0 41-17 18-40l-72-72l119-119c3-4 5-8 5-14c0-5-2-9-5-13l-70-70c-4-3-8-5-14-5c-5 0-10 2-14 5L112 471l-72-72c-23-22-40-14-40 17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func G(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 697h75c36 70 111 118 195 118c103 0 189-74 205-172c1-4 1-30 1-68c-50 57-124 93-207 93C132 668 0 548 0 398s132-270 284-270c82 0 156 36 207 92c-1-54-1-92-1-92h72v484c0 14-1 29-3 42c-22 129-136 227-273 227c-122 0-231-76-271-184m268-95c114 0 207-91 207-204s-93-204-207-204c-115 0-215 91-215 204s100 204 215 204"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GupperCase(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M397.018 475v-73h397c-1 25-5 49-11 73c-36 145-157 265-319 294c-155 28-305-37-391-155c-33-45-57-99-67-158c-38-210 108-411 324-449c156-28 306 36 391 153l-63 44c-58-79-153-131-261-131c-178 0-322 142-322 315c0 69 22 131 60 183c58 79 154 130 262 130c147 0 272-95 311-226z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Game(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M143 96V24h72v72zm359-72h71v72h-71zm143 215V96h72v287h-72v71h-72v72h72v72h72v71H573v-71h-71v-72H215v72h-72v71H0v-71h72v-72h71v-72H72v-71H0V96h72v143h71v-71h72V96h72v72h143V96h72v72h71v71zm-358 72v-72h-72v72zm215 0v-72h-72v72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gear(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m649 143l-68 69c18 28 31 60 37 94h99v106h-99c-6 34-19 65-37 93l68 70l-75 75l-69-69c-28 18-59 32-93 38v98H305v-98c-34-6-65-20-93-38l-70 69l-74-75l68-70c-18-28-32-59-38-93H0V306h98c6-34 20-66 38-94l-68-69l74-75l70 69c28-18 59-32 93-38V0h107v99c34 6 65 20 93 38l69-69zM358 507c82 0 149-66 149-148s-67-149-149-149s-147 67-147 149s65 148 147 148"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Geo(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M358 0c198 0 359 161 359 359S556 717 358 717S0 557 0 359S160 0 358 0m25 511v115c129-11 232-115 243-243H512v-47h114C616 206 512 103 383 92v115h-47V92C206 102 101 206 91 336h116v47H91c11 129 115 233 245 244V511z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Github(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m754 490l-2 8c-3 8-7 16-10 24c0 2-2 6-3 8c-32 71-92 105-185 122c-5 1-14 3-23 4c-90 11-158 13-248 5c-3 0-7-1-9-1c-125-14-199-59-238-131c-1-1-3-4-4-8c-4-7-7-16-10-23c-2-4-3-6-4-9C5 452 0 409 0 361c0-84 26-117 61-162C34 103 72 37 72 37s57-12 165 65c58-25 213-27 287-5c45-30 128-72 162-60c8 14 27 56 11 149c11 20 71 64 71 187c-1 45-6 84-14 117M99 461c0 111 78 160 187 169c7 1 14 1 22 2c51 1 115 2 178 0c14-1 28-2 41-4c104-15 147-75 147-159c0-61-28-141-140-141c-8 0-74 8-137 8c-71 0-144-8-155-8c-115 0-143 91-143 133m132-51c22 0 39 26 39 58s-17 58-39 58s-38-26-38-58s16-58 38-58m323 116c-23 0-40-26-40-58s17-58 40-58c22 0 39 26 39 58s-17 58-39 58m-166 28c-10 0-18-9-18-19c0-9 8-17 18-17s19 8 19 17c0 10-9 19-19 19m1 60c-29 0-52-25-52-40c0-6 4-11 10-11s10 5 10 11c1 3 14 20 32 20c14 0 27-12 27-20c0-6 5-11 10-11c6 0 11 5 11 11c0 18-20 40-48 40"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globe(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M610 106c66 66 107 156 107 253c0 98-41 186-107 252S455 717 358 717c-98 0-186-40-252-106S0 457 0 359c0-97 40-187 106-253S260 0 358 0c97 0 186 40 252 106M399 561h2c3 0 6 1 8 4c2 2 4 4 8 4c1 0 2-1 2-3c0-4 1-8 1-9c1-3 3-3 5-3c1 0 1 0 1-2l-2-2c-1 0-2 1-2 2c-1 1-2 2-3 0l-6 4l-7-2l-12-18l4-26c0-1-1-3-4-4s-3-3-2-4c-4-2-9-4-15-4c-1 0-4 1-9 2c-5 0-8 0-8-2c0-3 3-9 5-16c1-3 2-8 3-10c0-2 1-4 1-4l5-12l-1-2l-7-1c-2 0-9 4-14 8c-5 6-9 11-9 15l-20 4c-4 0-8-1-11-5c-1-5-4-11-8-17c-4-5-5-10-5-15c0-6 1-14 3-19c2-6 2-13-2-19c1 0 4-2 2-4l5-6l1-1l1 1c4-3 11-4 17-3c7 1 11-2 13-7l7 7c2 1 6-6 6-8l-3-2l18-4l1 2h8l10 6c1 0 2-1 3-2c2-2 3-2 4-1l9 10c-1 2-1 4-1 6c1 1 2 2 2 3c0 4 9 22 14 22c3 0 4-2 3-6v-8c0-5-1-11-3-17c-3-5-5-12-7-18v-2c0-3 2-6 6-8c3-2 5-4 5-4c3-2 6-5 10-7c3-2 6-7 8-10l3-7v-4h3s2-1 2-3c0 0-1-1-2-1c-1-1-4-3-6-5l2-2c3-3 4-10 2-13l8-4c-2 4 4 5 6 4l4-9c-2-5-3 0 0-6c3-1 6-2 9-4c3-1 5-2 8-2c1 0 2 1 2 1c2 0 3-1 3-1c0-3-1-5-2-5l5-11c4 0 8-2 10-6l8-1c2 0 3-1 3-3v-2l15-4l2-5l-4-6s1-1 1-2s-1-2-2-3c-1 0-2-1-3-1c-1-1-2-2-3-2h8s4-1 4-4c0-4-2-5-6-5c-11 0-28 4-34 16l-6 3l8-8l1-2c0-1-1-2-3-3h-3c10 0 16-3 20-6c3-3 6-6 13-8c7 1 14 1 21 1c6-1 12-1 19-1c3-2 7-6 9-10l10-1c2 2 9-3 9-6s-1-4-4-5s-6-3-6-6v-3c1-1 0-2-1-2c-4 0-14 5-17 7c-2 2-3-1-3-3l1 1l4-2l10-4l1-1c0-4-6-5-8-5c-1 0-3 1-5 2c-2 0-4 0-4-2l1-2c-7-5-13-15-13-18v-4c1-1 0-1-2-1h-3c-1 1-2 0-2-2s-4-13-9-13l-3 3c0 2-1 3-3 4s-3 2-3 3h-1l-10 6c0-2-1-2-2-1c-1 0-1 1-3 1h-2c5 0 4-7 0-7l-8 1c-2 0-2 0 0-2c1-2 3-4 1-6c0-2-1-2-2-1l2-3c2-1-1-5-3-7l-9-1l-6-6c-1 1-2 0-4-1c-1-1-3-3-4-3l-7 3l-18-4c-2 0-5 1-5 4c0 1 0 2 1 2c2 1 2 2 2 3s1 5 2 10c0 5 0 7-3 6l-4 6c0 2 2 3 3 3c3 2 6 5 8 8c2 2 2 5 1 8l-19 15v1c0 3 0 5 1 7c2 2 2 5 4 8c5 2 5 4 1 6c-2 0-4 1-6 2l-2 2c0 1-1 1-2 1h-5s0-1 1-1c0 0 1-1 1-2l-9-5l-4-7c1-3 1-5 0-6s-1-3-1-5c0-4-2-7-6-7c-5 0-9 1-14 2c2-1-2-6-3-6c-5 0-11-2-18-6s-12-6-16-6c-3 0-9 1-13 3c2-1 2-2 3-5l-6-10l-1-1c-2 0-4 1-6 3c-2 1-5 0-5-4c0-1 1-2 1-2c1-1 1-1 0-2c0-3 1-6 4-9c2-2 3-6 4-9c2-1 2-2 2-3c0-2 1-2 2-2c2 0 5-1 7-2c2-2 4-3 7-3l1-3c0-1-3-2-9-3c-7-1-10-1-10-3l1-1c6 2 11 3 14 3c6 1 9-1 15-4c2-1 8-3 13-6c0-1-1-1-4-2h4l7-4v-4l-2-2l10-2c-1-2 1-3 2-3c2 0 4 1 6 3s3 2 4 2l9-3c0-1 1-1 3-1c2-1 3-2 2-3l-7-7c-2 0-3-3-1-3c3 0 5-1 3-4c-2-1-5-3-8-4c-3-2-6-3-9-3c-2 0-6 1-6 4c0 2 1 2 2 2c2 0 3 1 3 2c2 1 1 2-1 2h-2c-2 0-5 2-7 6c-2 3-5 6-8 7c-1 0-1 0-1-1c1-1 0-1 0-2c-1-1-2-2-3-2c-2-1-4-1-4-2c0-2 2-3 4-6c1-3-2-4-7-4c-2 0-4 1-5 3c-1 1-2 3-3 5l-10-11l-8-1c0-2 1-4 2-6c2-4-6-9-9-12c-1 0-3-1-5-1c-1 0-7 5-11 7c-2 1-3 2-3 3c0 2 2 2 5 2h-1c-2 0-3 1-3 4c0 1 8 3 10 3s4 0 4 1c1 1 1 1 3 1l3-1v2l-1 2l1 3l-10 5l-1 1c-1 0-2 0-2 1s0 3 2 5c1 2 0 3-5 3l-3-2c0-2-2-4-7-6c-9-3-25-4-39-4c-7-1-12-1-16-1l-15 5l4 8c-1 0-3 1-3 1c-4-3-10-9-12-9l-7-1c-22 1-48 22-73 47s-47 53-58 70h2c2 0 3-1 6-2c2-1 3 1 3 4c0 2 0 3-1 5s1 3 5 3c1 0 2-1 2-3c1-2 1-2 1 0l2 7v1c0 1-1 2-2 3c-3 1 0 4 3 4h5l2-1l1-1c0 4 4 5 7 5h1s-1 2-2 3c-2 0-2 1 0 2l10 2v1l6 13c0 2-2 7-4 7s-2-1-2-2c1 0 1-1 1-3c0-1-1-2-1-3c-1-1-3-1-7-1c-1 0-4 0-2 3l5 11l2 2c-6 0-7 20-7 26l1 8l1 3v1l-1 9l12 20h3c0 1 0 2-1 3s-1 2-1 4l4 3c0 6 2 7 7 12c0 3 4 7 10 11c7 3 11 5 13 6c4 14 11 29 21 38l1 6l-2 2c-2 0-2 1-1 2l8 4c2-2 3 0 5 3c2 4 3 6 4 7v3l5 7l2 2l2-5c-2-5-9-16-15-25c-3-5-6-10-9-14c-2-3-3-5-3-6s-1-12-2-14c1 1 4 2 6 3c3 1 6 2 8 3c1 8 4 14 9 19c4 4 8 9 11 15c-2 2 2 2 4 2c1 2 1 3 1 7c9 9 27 30 27 39v1l-2 6c4 9 13 16 22 19h2c5 3 12 6 19 9c6 4 12 7 18 9l10-6c5 1 11 6 18 12s16 13 28 15c4-3 6-2 6 2v2l11 13l2 6c6 4 13 10 16 17M270 58l-1-2c-2-1-7-3-9-3c-1 0-3 1-7 2c-7 4-18 7-26 11c-4 2-7 3-9 3c4 0 8-1 12-3c3-2 7-3 11-3l3 2c2 1 3 1 5 1c1 0 2 0 4 2c2-1 4-1 7 1l3-2v-4l-1-2l2 1c2 0 4-1 6-4m151 586l1 6c0 2-1 2-2 4c-3 2-8 5-8 11c0 1 1 2 1 4c1 1 2 2 2 4c0 1-1 2-1 3c69-12 131-45 179-96h-1c-2 1-5 1-6 0l-3 1l-5-1h-3c-2-3-4-5-7-7l-1-1c-1 0-2 0-2 2c0-9-7-17-15-17c0 1-1 1-1 1h-2l-2-1h3l2-6l-5-3l-1 1c-4-1-7-3-9-7l-3-1v1l-2 1c-4 1-7 2-9 3c-4-1-7-3-8-4l-12 1c0-4-2-9-7-9c-6 0-14 0-16 7c0 2 1 5 2 6v4l-1 2l-1 1h-1l-3-7l3-4c-1-2-1-3-1-5s-1-3-1-5l-1-1h-3l-6 4h-4l-1 2c-1 0-1 1-1 2c0 0-1 0-1 1l-1-1h-4c-2 2-3 5-4 8l2 2l-5 2l-1 2l-3 2v1h-1v5h-1c-1-4-4-8-10-11h-3v2c0 2 2 3 4 4c2 2 3 4 4 5c-1-1-3 0-3 1v2l6 9v17l2 4c-2 5-4 10-7 13v-1l-2 1l-1 1l-1 4l1 1v1l-2-2l-1 6l-5 2c-2 1-3 2-2 5c0 2-1 3-3 4l1 2l-2 3c0 1 0 2-1 2v4l1 7l2 1l2-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Good(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M505 572v37h109v70H243l-59-45V315s82-61 118-136c36-76 34-164 34-164h94s4 72 0 129c-5 58-24 129-24 129h260v82H505v37h161v73H505v36h130v71zM0 629h142V321H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Google(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M82 19h491c45 0 82 37 82 82v82H553V81h-51v102H399v51h103v103h51V234h102v359c0 45-37 81-82 81H377c1-6 1-12 1-18c0-68-14-102-86-155c-20-15-66-46-66-68c0-25 7-37 45-67s66-70 66-119c0-53-23-100-64-124h59l50-53H160C97 70 38 97 0 137v-36c0-45 37-82 82-82m166 233c10 78-25 129-84 127c-60-2-117-57-127-136s30-138 89-137c60 2 112 67 122 146M0 543V366c29 29 70 46 121 46c7 0 14 0 21-1c-7 13-12 29-12 44c0 26 15 41 33 58h-41c-46 0-88 12-122 30m309 131H82c-39 0-72-26-80-63c27-42 84-73 150-72c20 0 40 5 57 10c48 34 88 54 97 93c2 7 3 15 3 23z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grab(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M298 14h85v42h-85zM128 56h85v43h-85zm85 43h42V56h43v213h-43v-85h-42zm212-43h85v43h-85v170h-42V56zM85 184V99h43v85zm468-85v85h42v42h-42v85h-43V99zm42 85v-43h43v43zm43 42v-42h42v170h-42zm-468 43h-42v-85h42zm-127 0h85v42H43zM0 396v-85h43v85zm128-42v-43h42v-42h43v127h-43v-42zm467 85v-85h43v127h-43zM85 396v43H43v-43zm43 128H85v-85h43zm425 42v-85h42v85zm-383-42v42h-42v-42zm43 42v43h-43v-43zm297 86v-86h43v128h-43zm-255 42h-42v-85h42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Graph(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 648h717v-56H56V46H0zm98-110h616V96L584 215l-121-44l-128 146l-131-47L98 380z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grayscale(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 646V48h768v598zm383-114v63h64v-63h65v-61h-65v-63h65v-61h-65v-62h65v-61h-65v-63h65V99h-65v62h-64V99h-65v62h-62V99H51v496h267v-63zm0-371v63h-65v-63zM256 285v-61h62v61zm191 0h-64v-61h64zm-129 0h65v62h-65zm-62 123v-61h62v61zm191 0h-64v-61h64zm-129 63v-63h65v63zm-62 61v-61h62v61zm127-61h64v61h-64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gree(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 532l309 172l309-172V176L309 4L0 176z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GroupIcon(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M669.775 499c42 23 66 64 63 105c-2 25-3 26-35 30c-20 3-146 4-276 4c-147 0-301-3-312-6c-44-11-20-91 38-129c45-29 137-73 162-78c34-7 38-28 0-96c-9-15-19-61-20-109c-1-78 14-131 82-157c14-5 28-7 41-7c45 0 87 25 104 62c25 49 15 180-11 227c-30 53-27 70 6 79c22 6 90 38 158 75m-544-30c-16 11-31 25-43 39c-38 0-67-1-70-2c-25-6-11-52 22-72c25-16 76-41 90-44c18-4 22-15 0-53c-5-8-11-34-11-61c-1-44 8-74 46-87c32-12 68 3 81 30c13 28 6 100-8 126c-16 31-13 39 5 44c4 1 12 5 23 10c-44 18-104 51-135 70m658-38c24 13 37 35 35 58c-1 14-1 15-19 17c-5 1-28 2-58 2c-13-17-31-33-53-44c-41-24-91-50-128-66c10-4 19-7 23-8c19-4 22-15 0-53c-4-8-11-34-11-61c-1-44 7-74 45-87c33-12 69 3 81 30c13 28 8 100-6 126c-16 31-14 39 3 44c12 4 50 22 88 42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Guillemotleft(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M516 556L11 324l-11-6v-67l11-7L516 14v79L98 284l418 192z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Guillemotright(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 14l505 230l11 7v67l-11 6L0 556v-80l418-192L0 93z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Guilsinglleft(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M516 556L11 324l-11-6v-67l11-7L516 14v79L98 284l418 192z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Guilsinglright(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 14l505 230l11 7v67l-11 6L0 556v-80l418-192L0 93z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gumroad(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M85 566V184c0-20 23-43 42-43h310c14 25 41 43 73 43c47 0 85-39 85-85c0-48-38-85-85-85c-32 0-59 17-73 42H127c-31 0-63 15-88 39c-25 25-39 57-39 89v382c0 32 14 64 39 89c25 24 57 39 88 39h341c31 0 63-15 88-39c25-25 39-57 39-89V354c0-35-14-68-40-92c-24-23-56-36-87-36H298c-34 0-66 14-91 40c-23 24-37 56-37 88v42c0 32 14 64 39 88c25 25 57 40 89 40h11c15 25 42 42 74 42c47 0 85-38 85-85s-38-85-85-85c-32 0-59 17-74 43h-11c-20 0-43-24-43-43v-42c0-20 19-43 43-43h170c19 0 42 19 42 43v212c0 19-23 43-42 43H127c-19 0-42-24-42-43"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func H(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M72 0v292c42-40 109-64 171-64s110 24 152 64c44 42 71 102 71 168v294h-71V456c-3-87-64-157-152-157c-87 0-169 70-171 157v298H0V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HupperCase(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M455 280V0h72v735h-72V352H72v383H0V0h72v280z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hatena(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M397 46c-17 6-91 33-94 34c-2-1-76-28-93-34c17-6 91-32 93-33c3 1 77 27 94 33m212 0c-17 6-91 33-93 34c-3-1-77-28-94-34c17-6 91-32 94-33c2 1 76 27 93 33m-314 81c-17 6-92 31-94 32c-2-1-76-26-93-32c17-6 91-32 93-33c2 1 77 27 94 33m115 32c-2-1-77-26-94-32c17-6 92-32 94-33c3 1 77 27 94 33c-17 6-91 31-94 32m208 0c-2-1-77-26-94-32c17-6 92-32 94-33c2 1 77 27 94 33c-17 6-92 31-94 32m-208-58l-48 16s-14 6-26 10c13 5 71 24 74 26c4-2 60-21 74-26c-15-5-70-25-74-26M93 175c3 0 77 27 94 33c-17 6-91 31-94 32c-2-1-76-26-93-32c17-6 91-33 93-33m211 0c3 0 77 27 94 33c-17 6-91 31-94 32c-2-1-76-26-93-32c17-6 91-33 93-33m213 0c2 0 77 27 94 33c-17 6-92 31-94 32c-2-1-76-26-93-32c17-6 91-33 93-33m209 0c2 0 76 27 93 33c-17 6-91 31-93 32c-3-1-77-26-94-32c17-6 91-33 94-33M19 208s70 24 74 25c4-1 52-17 75-25c-23-8-71-25-75-27c-4 2-73 27-73 27zm211 0c23 8 71 24 74 25c4-1 52-17 75-25c-23-8-71-25-75-27c-3 2-51 19-74 27m421 0c23 8 71 24 75 25c4-1 72-24 74-25c-2-1-70-25-74-27c-4 2-52 19-75 27m-548 82c17-6 92-31 95-32c2 1 76 26 93 32c-17 6-91 32-93 33c-3-1-78-27-95-33m398 0c-17 6-92 32-94 33c-2-1-77-27-94-33c17-6 92-31 94-32c2 1 77 26 94 32m27 0c17-6 92-31 94-32c3 1 77 26 94 32c-17 6-91 32-94 33c-2-1-77-27-94-33m-330 26c3-1 50-18 73-26c-23-8-70-24-73-25c-4 1-51 17-74 25c23 8 70 25 74 26m209-51c-3 1-51 17-74 25c23 8 71 25 74 26c3-1 50-18 73-26c-23-8-70-24-73-25m215 51c3-1 50-18 73-26c-23-8-70-24-73-25c-3 1-51 17-74 25c23 8 71 25 74 26m-306 81c17-6 92-33 94-33s76 27 93 33c-17 6-91 32-93 33c-2-1-77-27-94-33m51 110h-32v47h-18v26h18v101h32V580h18v-26h-18zm-326 2H7v172h34v-76h71v76h32V509h-32v64H41zm538 44h-31v128h31v-51c0-12 1-21 1-24c2-8 5-14 10-19c4-4 9-6 15-6c5 0 9 2 12 4s6 7 8 12c1 5 2 14 2 29v55h30v-82c0-14-3-26-10-33c-9-8-20-13-32-13c-6 0-12 2-17 4c-6 2-12 6-19 12zm-314 13v-12h32v127h-32v-14c-10 8-24 14-38 14c-35 0-63-29-63-64c0-36 28-63 63-63c14 0 28 4 38 12m229 76h27c-5 13-16 26-25 30s-20 9-32 9c-18 0-34-7-46-19c-12-11-18-27-18-45s6-33 18-45s27-18 45-18s34 6 46 18c12 11 17 28 17 48v5h-97c1 9 5 16 11 21s15 8 24 8c11 0 22-4 30-12m285-76v-12h33v127h-33v-14c-10 8-24 14-38 14c-35 0-63-29-63-64c0-36 28-63 63-63c14 0 28 4 38 12m-349 36h67c-2-6-7-13-13-17c-5-4-12-6-20-6s-16 2-23 7c-4 3-7 9-11 16m-203 50c19 0 35-16 35-35s-16-36-35-36s-34 17-34 36s15 35 34 35m514 0c19 0 35-16 35-35s-16-36-35-36s-34 17-34 36s15 35 34 35"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hatenabookmark(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M293 390c8 6 12 16 12 31c0 13-4 22-14 28c-9 6-25 9-48 9h-33v-77h35c23 0 39 3 48 9m-16-79c-8 5-24 7-48 7h-19v-70h21c23 0 39 3 47 8s13 15 13 28s-5 22-14 27M80 29h490c44 0 80 36 80 80v490c0 44-36 80-80 80H80c-44 0-80-36-80-80V109c0-44 36-80 80-80m453 154h-85v230h85zM384 478c8-14 12-30 12-49c0-26-7-47-21-63c-14-15-34-24-58-26c22-6 38-15 48-26c10-12 15-28 15-48c0-15-3-29-10-41s-17-22-30-29c-11-6-24-11-40-13c-16-3-43-4-83-4h-96v350h99c40 0 69-1 86-4c18-3 32-7 44-14c15-8 26-19 34-33m139 38c9-9 13-19 13-32c0-12-4-23-13-32c-9-8-20-12-33-12s-24 4-32 13c-9 8-13 19-13 31c0 13 4 23 13 32s19 13 32 13s24-4 33-13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M582 30c166 0 221 192 163 297c-89 163-361 336-361 336S111 490 23 327C-35 222 19 30 186 30c154 0 193 133 198 153c5-20 44-153 198-153"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heartempty(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384.022 668c5 0 11-1 15-4c11-7 259-165 344-320c36-66 32-155-9-224c-36-61-96-95-168-95c-99 0-153 53-182 98c-29-45-84-98-182-98c-73 0-132 34-169 95c-41 69-44 158-8 224c84 155 333 313 343 320c5 3 10 4 16 4m-182-583c118 0 148 98 152 117c4 13 16 24 30 24c13 0 26-11 29-24c5-19 35-117 153-117c65 0 100 35 117 66c30 49 33 118 8 164c-67 120-252 250-307 288c-56-38-241-168-307-288c-25-46-22-115 8-164c17-31 52-66 117-66"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Help(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M358 0c198 0 359 161 359 359S556 717 358 717S0 557 0 359S160 0 358 0m-27 470h24c3-61 17-81 74-115c35-22 46-29 59-43c18-19 27-43 27-71c0-71-57-114-150-114c-32 0-57 6-82 15c-50 19-81 59-81 103c0 22 15 33 46 33c26 0 40-11 39-29v-24c0-15 6-41 14-54c10-16 31-27 55-27c45 0 74 36 74 89c0 20-3 39-11 57c-7 14-15 25-39 51c-38 40-49 69-49 114zm13 145c29 0 52-23 52-52s-22-52-52-52c-29 0-51 23-51 52s22 52 51 52"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heteml(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1132 146c-24-4-48 6-63 24c21 25 31 57 34 103c7 84 5 150-71 212c-27 22-56 37-87 49c3 20 3 42 0 63c-6 30-57 27-83 28c-41 2-95 1-111-9c-9-6-2-31 3-36c3-4 5-10 5-16c-27 0-54 0-82-1c-40-1-81 1-120 1c0 11-1 23-3 34c-4 30-48 27-70 27c-35 2-81 0-93-10c-9-7-4-28 2-36c7-10 8-32 9-53c-23-17-42-42-56-77c-2-6-5-13-7-19c-28 23-60 44-94 66c-75 49-150 31-198-5C-2 456-8 371 8 345c15-26 53-38 117-69c60-29 59-69 81-108c4-6 10-12 18-17c-7-6-13-14-17-25c17-5 35 3 45 12c-1-12 9-30 28-41c5 11 5 24 1 35c29-5 58-5 73-3c23-14 57-25 117-32c120-16 453-18 525 18c23 12 42 24 57 37c22-17 53-23 79-6M201 262c7 0 13-6 13-13c0-6-6-12-13-12s-12 6-12 12c0 7 5 13 12 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m623.352 238l116 106c14 13 10 23-9 23h-87v277c0 19-16 36-35 36h-151V486c0-19-17-36-36-36h-95c-19 0-36 17-36 36v194h-151c-19 0-36-17-36-36V367h-86c-19 0-23-10-9-23l339-308c14-13 38-13 53 0l77 69V50c0-19 17-36 36-36h75c19 0 35 17 35 36z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Horizontal(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M330 129H43v179c-15 4-31 11-43 21V68C0 31 31 0 68 0h237c37 0 68 31 68 68v238h-43zm141 84h69c-3-28-15-54-35-74c-27-27-62-38-96-35V53c47-3 96 13 132 50c31 31 47 69 50 110h67l-92 100zM68 344h552c37 0 67 31 67 68v238c0 37-30 67-67 67H68c-37 0-68-30-68-67V412c0-37 31-68 68-68m66 329h424V387H134zM78 568c20 0 38-17 38-37s-18-37-38-37s-37 17-37 37s17 37 37 37"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hot(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M630 379c0-125-65-190-190-190c125 0 190-64 190-189c0 125 64 189 189 189c-125 0-189 65-189 190M345 58c0 227 119 346 346 346c-227 0-346 119-346 346c0-227-118-346-345-346c227 0 345-119 345-346m250 478c0 59 31 91 90 91h1c-60 0-91 31-91 90c0-59-31-90-90-90c59 0 90-32 90-91"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hyphen(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M274 248v73H0v-73z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Iicon(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M82 82H0V0h82zm-5 608H5V164h72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IupperCase(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 735V0h72v735z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Image(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 640V54h717v586zm644-194V127H73v271l156-93l204 111l83-49zM510 298c-32 0-57-26-57-58s25-58 57-58s57 26 57 58s-25 58-57 58"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Info(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M358 0c198 0 359 161 359 359S556 717 358 717S0 557 0 359S160 0 358 0m-4 132c-10 12-15 26-15 41c0 12 4 22 10 29c7 7 17 11 27 11c14 0 27-6 37-18c9-11 14-27 14-44c0-10-3-19-10-26s-17-11-27-11c-13 0-26 6-36 18m-91 207l12 13c9-10 18-18 24-23c6-4 11-6 15-6c3 0 6 1 8 3c1 3 2 6 2 11c0 25-3 41-16 101s-20 104-20 133c0 11 2 19 6 24c3 5 9 9 17 9c13 0 32-10 54-27c22-18 44-43 68-75l-13-11c-8 10-16 17-22 22c-6 4-11 8-15 8c-3 0-6-3-8-5c-1-3-2-7-2-13c0-4 1-17 5-40c3-23 2-23 8-57c1-10 4-24 7-41c8-51 13-82 13-92c0-9-3-17-6-22c-4-5-10-7-17-7c-12 0-28 9-50 25c-22 17-44 40-70 70"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ink(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M500 351c-6-10-18-26-33-48c-16-22-34-46-52-74c-18-29-38-59-56-93c-17-34-32-68-43-104c-3-11-9-20-17-25s-18-7-28-6c-9-1-18 1-27 6c-8 5-15 14-18 25c-11 36-26 70-43 104s-36 64-54 93c-19 28-37 52-52 74c-16 22-27 38-34 48c-13 20-25 44-32 68c-7 25-11 49-11 77c0 38 8 73 22 106c13 34 32 62 57 87c26 25 54 44 87 58s68 21 105 21c38 0 73-7 107-21c33-14 61-33 86-58s44-53 57-87c14-33 22-68 22-106c0-28-5-52-12-77c-7-24-18-48-31-68M186 624c-18 0-35-6-47-19c-13-13-20-30-20-48c0-14 4-26 11-37c2-2 4-6 8-11s10-12 14-19c5-7 9-15 13-23s8-17 10-26c2-7 6-10 11-8c6-1 11 1 12 8c3 8 6 18 11 26c4 8 9 16 13 23c5 7 10 14 14 19s6 9 8 11c7 10 11 23 11 37c0 18-8 35-21 48c-12 13-29 19-48 19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Instagram(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M680 289V82c0-37-31-68-68-68H68C31 14 0 45 0 82v207h165c26-71 95-122 175-122s149 51 175 122zm0 60H527v5c0 103-84 187-187 187s-187-84-187-187v-5H0v277c0 37 31 68 68 68h544c37 0 68-31 68-68zM59 74h34v156H59zm68 0h34v156h-34zm68 0h34v98c-11 8-23 17-34 26zm145 408c71 0 127-57 127-128s-56-127-127-127s-128 56-128 127s57 128 128 128M459 74h161v156H512c-14-21-33-38-53-52z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Instapaper(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M129 14h408c72 0 129 57 129 129v307c-3 30-18 58-34 67c-18 10-54-3-81 13c-28 16-31 52-49 66c-22 16-40 21-72 24H114C52 620 0 570 0 508V143C0 71 57 14 129 14m256 140h29v-38H254v38h27c10 0 19 9 19 20v285c0 11-9 20-19 20h-27v38h160v-38h-29c-10 0-19-9-19-20V174c0-11 9-20 19-20m167 466H431c114 0 235-72 235-170v58c0 62-52 112-114 112m-438 15h438c62 0 114-52 114-114v13c0 62-52 114-114 114H114C52 648 0 596 0 534v-13c0 62 52 114 114 114m0 26h438c62 0 114-51 114-113v3c0 72-57 129-129 129H129C57 680 0 623 0 551v-3c0 62 52 113 114 113"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Internetexplorer(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M537 489h183c-44 132-170 227-317 227c-144 0-267-92-314-220c-41 85-50 154-18 185c38 39 96 35 166 1c-89 51-165 65-212 18c-55-55-18-253 119-426c0 0 0-1 1-1c5-7 12-15 18-22c0 0 1 0 1-1c2-3 5-6 8-9c1-1 2-2 2-3c3-2 5-5 7-7l3-3c2-2 4-6 6-8c1-1 3-2 4-3l10-10l15-15c1-1 2-1 3-2c5-5 11-10 15-15c2-1 3-2 4-3l15-12c1-2 2-3 3-4l12-9c2-2 4-3 5-4c4-4 10-8 14-11c2-2 4-3 6-5l12-9c3-2 6-4 8-5c2-2 5-3 7-5l15-9l4-4c10-6 21-12 31-18c-85 36-183 111-268 213c-11 13-22 27-31 40C99 170 236 47 403 47c8 0 16 0 24 1C546-7 652-15 695 29c37 36 38 89 13 152c14-48 9-88-20-117c-32-32-91-27-165 6c125 48 214 170 214 311c0 14-1 28-2 41H261c3 26 12 48 22 67c27 49 71 76 120 76c54 0 107-17 134-76M263 329h287c-16-93-80-131-147-131c-66 0-123 48-140 131"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Invert(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 48h768v598H0zm384 547v-71l177-177l-177-177V99H51v496zm0-71L207 347l177-177z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Iphone(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M76 0h265c42 0 76 34 76 76v616c0 42-34 76-76 76H76c-42 0-76-34-76-76V76C0 34 34 0 76 0M49 617h319V144H49zm159 105c24 0 42-18 42-41c0-24-18-42-42-42c-23 0-41 18-41 42c0 23 18 41 41 41"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Italic(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8 643l-8 36c3 0 8 0 16-1c6 0 16-1 27-2c35-4 62-6 79-5h87c19 2 35 4 46 6c10 1 17 1 19 1c3 2 6 2 8 2h11c3 0 7 0 12-1h5c1-1 3-1 5-1c2-3 3-7 3-10s2-5 3-7c0-4 1-7 1-12c0-3 1-8 1-13c-8-1-15-3-20-4s-8-1-9-1c-17-2-33-5-53-9c-1-3-1-5-1-7v-5l6-19l18-102l18-68l27-136c5-31 16-74 29-132c1-9 3-21 6-35c4-16 7-28 11-38c11-4 26-9 45-14c17-4 35-8 49-13c2-9 4-17 6-22c0-3 1-6 1-9c1-2 1-5 1-8c-3 0-9 0-17 1c-8 0-17 1-29 2c-45 3-76 4-96 4h-11c-6 0-13-1-22-1l-141-6l-9 45h8c2 1 7 1 10 1c31 2 51 6 64 12v18l-4 20l-9 59l-7 28l-14 68c0 1-3 11-8 27c-5 17-11 42-17 77l-5 28l-26 116l-12 59c-3 18-10 34-18 45c-12 5-28 12-52 17c-15 4-28 6-34 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func J(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M188 82h-82V0h82zm-77 694V164h73v610c0 53-30 99-73 124c-20 12-45 19-72 19c-13 0-27-2-39-5v-78c11 7 25 12 39 12c39 0 71-32 72-70"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JupperCase(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M317 529V0h72v527c0 66-28 125-72 166c-41 38-96 61-156 61c-62 0-120-25-161-66l52-51c28 29 66 46 109 46c85 0 155-69 156-154"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Jpa(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M162 268c-3 0-4-1-6-2c-1-2-1-4-1-7s0-6 1-7c2-2 3-3 6-3c2 0 4 1 5 3c1 1 2 4 2 7s-1 5-2 7c-1 1-3 2-5 2m107-114c-1 0-3 0-4-1c-2-1-3-4-3-7c0-2 1-4 1-6c1-2 3-4 6-4s4 1 5 3s2 4 2 6c0 3-1 5-2 7s-3 2-5 2m-35 115c-1 0-2 0-3-1c0 0-1-1-1-3c0-1 1-2 2-3s2-1 4-1l2-1h2l1-1v3c0 3-1 5-2 6s-3 1-5 1M73 258H62l5-16zm164-111h1c1 0 2-1 3-1h1v2c0 3-1 5-2 6c-2 1-3 1-5 1h-3c0-1-1-2-1-3c0-2 1-3 2-4c1 0 2-1 4-1m35 47c-1 1-2 1-4 1h-8v-12h8c2 0 3 1 4 2c2 0 2 2 2 4s0 4-2 5m27-1c2 0 3 0 5 1c1 1 2 3 2 5h-14c0-2 1-4 2-5s3-1 5-1m0 56c2 0 4 1 5 3c1 1 2 4 2 7s-1 5-2 7c-1 1-3 2-5 2c-3 0-5-1-6-2c-1-2-2-4-2-7s1-6 2-7c1-2 3-3 6-3m75-137h868c169 0 305 120 305 267v325H425c-168 0-305-119-305-266v-98C50 314 0 252 0 179C0 82 87 4 194 4c82 0 151 45 180 108m-49 18h-7v30h8v-21c1-2 3-3 6-3c2 0 3 1 4 2s1 2 1 4v18h8v-20c0-4-1-7-3-8c-2-2-5-3-8-3c-2 0-4 1-6 2l-3 3zm11 122c1 1 1 2 1 3v19h8v-21c0-3-1-6-3-8c-2-1-5-2-8-2c-2 0-4 0-6 2c-1 0-2 2-3 3v-4h-7v30h8v-21c1-2 3-4 6-4c2 0 3 1 4 3m1-76v41h8v-41zm-13 11h-7v30h8v-14c0-3 0-5 1-6c1-2 3-3 6-3h2v-8h-1c-2 0-4 1-6 2c0 1-2 2-3 4zm-21-51c1 1 2 2 2 3s-1 2-2 2s-2 1-3 1h-3c-3 1-5 1-7 2c-3 2-4 4-4 8c0 3 1 5 3 7c1 1 4 2 6 2c3 0 5 0 6-1c2-1 3-2 4-3v2s0 1 1 1h8v-1s-1 0-1-1v-19c0-4-2-6-4-8c-3-1-6-2-9-2c-5 0-9 2-11 4c-1 2-2 4-2 7h8c0-1 0-2 1-3c0-1 2-1 4-1zm-8 75c-2-2-3-4-3-7h22v-6c-1-2-1-4-3-6c-1-2-3-4-5-5s-4-1-7-1c-4 0-8 1-11 4s-4 7-4 12c0 6 2 10 5 12c3 3 6 4 10 4c5 0 9-1 12-4l3-6h-8c-1 1-1 2-2 2c-1 1-3 2-4 2c-2 0-4-1-5-1m19 48c0-5-1-8-4-12c-2-3-6-4-11-4c-6 0-10 1-12 4c-3 4-4 7-4 12c0 4 1 8 4 11c2 3 6 5 12 5c5 0 9-2 11-5c3-3 4-7 4-11m-52-129h-7v42h8v-16c0 2 1 3 2 3c2 1 4 2 6 2c4 0 7-1 10-4c2-3 3-7 3-12s-1-9-4-12c-2-2-5-4-9-4c-2 0-4 1-6 2l-3 3zm7 46h-17v41h8v-15h9c4 0 8-1 10-3s4-5 4-10c0-4-2-8-4-10s-6-3-10-3m2 57v7h8v-7zm8 41v-30h-8v30zm-15-6c-1-1-1-1-1-3v-15h5v-6h-5v-9h-8v9h-4v6h4v18c0 2 1 3 2 4c1 2 3 2 7 2h4v-6zm-23-132c1 1 1 2 1 3s0 2-2 2c0 0-1 1-3 1h-2c-4 1-6 1-7 2c-3 2-5 4-5 8c0 3 1 5 3 7c2 1 4 2 7 2c2 0 4 0 6-1l3-3s0 1 1 2v1h9v-1c-1 0-1 0-1-1c-1 0-1-1-1-2v-17c0-4-1-6-4-8c-2-1-5-2-9-2c-5 0-8 2-11 4c-1 2-2 4-2 7h8c0-1 0-2 1-3s2-1 4-1zm-14 122c-3 1-5 4-5 8c0 3 1 5 3 6c2 2 4 3 7 3c2 0 4-1 6-2c1-1 2-2 4-3v4h9v-1l-1-1c-1-1-1-1-1-2v-18c0-3-1-6-4-7c-2-1-5-2-9-2c-5 0-8 1-10 4c-2 2-2 4-3 6h8c0-1 1-2 1-2c1-1 2-2 4-2s3 0 4 1c1 0 1 1 1 3c0 1 0 1-1 2h-4l-2 1c-3 0-6 1-7 2m-25-113h-9v1c0 4 1 8 3 11s5 4 11 4c5 0 8-2 10-5c2-2 2-5 2-8v-29h-8v29c0 2-1 3-1 4c-1 2-2 2-4 2s-3 0-3-2c-1-1-1-3-1-6zm9 88v7h8v-7zm8 41v-30h-8v30zm-36-27c-2 3-4 7-4 12s2 9 4 12c2 2 6 4 11 4c4 0 8-2 11-5c1-2 2-5 2-7h-8c0 1-1 3-1 4c-1 1-2 1-4 1c-3 0-5-1-6-4v-10c1-3 3-4 6-4c2 0 3 0 4 1s1 2 1 3h8c0-4-2-7-4-9c-2-1-6-2-9-2c-5 0-8 1-11 4m-33 0c-2 4-4 7-4 12c0 4 2 8 4 11c3 3 7 5 12 5s9-2 11-5c3-3 4-7 4-11c0-5-1-8-4-12c-2-3-6-4-11-4s-9 1-12 4m-75 18l2 9h10l-15-41h-9l-15 41h9l3-9zm37 7c2-2 3-4 3-7s-1-4-2-6c-2-1-4-3-7-3c-5-1-8-2-9-3c-1 0-1-1-1-2c0 0 0-1 1-1c1-1 2-1 4-1s3 0 4 1s1 2 1 3h8c0-4-2-6-4-8s-5-2-9-2s-7 1-10 3c-2 2-3 4-3 7c0 2 1 4 2 5c2 2 4 3 8 4s7 2 8 2s1 1 1 2s-1 2-1 2c-1 1-3 1-4 1c-3 0-5-1-6-2c0 0-1-1-1-3h-8c0 3 1 6 4 8c2 2 5 3 10 3s8-1 11-3m29 0c2-2 4-4 4-7s-1-4-3-6c-1-1-4-3-7-3c-5-1-7-2-8-3c-1 0-1-1-1-2c0 0 0-1 1-1c1-1 2-1 3-1c3 0 4 0 5 1c0 1 1 2 1 3h8c-1-4-2-6-4-8c-3-2-6-2-10-2s-7 1-9 3s-3 4-3 7c0 2 1 4 2 5c1 2 4 3 7 4c5 1 8 2 8 2c1 0 2 1 2 2s-1 2-2 2c-1 1-2 1-4 1s-4-1-5-2c-1 0-1-1-1-3h-8c0 3 1 6 3 8c3 2 6 3 11 3s8-1 10-3m284 407h1097V379c0-134-125-242-280-242H383c4 13 6 27 6 42c0 96-87 175-195 175c-17 0-33-2-49-6v90c0 133 126 241 280 241M299 147h2c1 0 2-1 2-1h2v2c0 3-1 5-3 6c-1 1-3 1-4 1h-3c-1-1-1-2-1-3c0-2 0-3 2-4c0 0 2-1 3-1m854 32h69c139 0 251 93 251 207v254h-64V396c0-86-91-156-202-156h-54v166h192v64h-192v170h-64V179zM459 579h55V179h64v461H444c-138 0-251-93-251-208v-22h67v17c0 84 89 152 199 152m245-400h144c57 0 107 25 141 64c25 30 41 68 41 110s-16 80-41 110c-34 39-84 64-141 64h-80v-64h88c58-6 103-53 103-110s-45-104-103-110H704v397h-64V179z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func K(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M411 228L137 466l319 288H349L81 513l-9 9v232H0V0h72v426l229-198z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KupperCase(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M467 0L117 331l374 404h-94L71 384v351H0V0h71v280L368 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Key(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m318 483l-39-40s-111 43-223-68c-95-95-53-244 9-305c66-66 207-106 310-3s59 225 59 225l283 282l-8 147l-144 9l-8-72l-75-9l-7-75l-74-10l-11-72zM104 110c-44 44-56 103-27 131c29 29 87 17 131-27s57-103 28-131c-28-29-88-17-132 27m579 520l2-42l-269-269l-22 22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyboard(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M30 101h708c16 0 30 14 30 30v432c0 16-14 30-30 30H30c-16 0-30-14-30-30V131c0-16 14-30 30-30m21 440h666V152H51zm72-348h61v62h-61zm116 0h61v62h-61zm114 0h62v62h-62zm116 0h60v62h-60zm115 0h61v62h-61zM181 316h62v62h-62zm116 62v-62h61v62zm115 0v-62h61v62zm177 0h-62v-62h62zm-60 122H239v-61h290z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Kudakurage(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m613.132 194l5 8c17 33 42 79 22 131c-7 18-22 31-36 42c-6 5-12 10-17 15c13 50 34 90 61 117c2 2 6 5 10 8c12 7 28 18 24 38c0 2-2 4-5 4c-23 0-43-14-66-43c-20-29-39-65-54-113c-5 2-11 3-17 5c-3 0-7 1-10 2c4 24 8 46 11 70c6 52 12 101 36 139c3 5 7 9 11 14c10 10 21 23 13 43c-1 2-2 3-3 3c-2 1-3 0-4 0c-60-33-71-114-81-193c-3-24-6-49-10-72c-4 0-7 1-11 1c-11 2-24 4-36 4c3 67 11 132 18 190c1 16 3 33 5 48c1 8 4 21 8 35v2c0 1 1 3 1 5c3 7 9 18 4 27c-2 3-6 6-10 8h-5c-20 0-25-38-27-62l-1-6c-2-17-5-37-7-55c-8-61-16-124-19-189c-13 2-28 3-41 4c-9 0-20 1-30 1c18 45 11 102 5 156c-7 60-13 118 10 159c3 5 7 8 12 12c9 7 20 16 12 38c-1 3-4 4-6 3c-72-24-63-113-55-199c5-55 10-112-5-153c-1-6-1-10 1-13c-13 1-27 2-42 2c-8 36-15 75-21 118c-2 12-4 25-6 37c-8 56-17 114-16 179c1 5 2 11 3 16c3 13 5 29-10 38c-1 1-2 1-3 1h-1c-18-2-17-28-17-42v-5c0-74 13-150 24-225c6-40 11-77 15-114c-9-1-22-1-32 1c-5 0-10 1-14 1c-8 78-28 152-60 224c-1 4-3 8-4 12c-7 18-16 40-36 48c-2 0-5 0-6-3c-11-18 2-37 11-51c4-5 7-11 9-15c30-63 48-136 56-216c-5-2-15-2-19-2c-3 0-7 1-10 1h-13c-4 68-28 128-69 170c-13 14-28 28-43 26c-2 0-4-1-5-2c-4-6-5-13-4-18c3-8 12-13 20-17c3-2 7-4 9-6c36-38 59-93 63-153c0-1-4-1-7 0h-9c-1-1-2-1-3-2c-6-5-11-9-17-14c-32-26-62-50-68-102c-6-55 13-95 30-133c3-5 5-11 8-16c24-53 54-94 89-119c44-31 101-47 166-47c37 0 76 6 114 17c90 26 155 82 205 177m-38 168c41-24 50-77 21-130c-41-77-93-137-140-164c-31-17-105-38-161-38c-6 0-12 1-17 1c-87 8-150 43-191 106c-11 19-50 85-55 136c-4 41 3 70 22 92c26 28 71 41 143 41c37 0 76-4 114-7c20-2 40-4 58-5c11-1 23-2 35-2c65-3 131-6 171-30m-118-169c-2 3-5 6-8 7c-1 1-4 2-5 2h-4s-2-1-3-1c-5-2-10-6-12-11c-1 0-1 0-2-1c-1-4-1-8-1-12c0-7 2-13 5-17c2-4 5-6 9-7c1-1 2-2 3-2h4c1 0 2 0 3 1c6 2 11 5 14 10c0 1 0 1 1 1c1 4 2 8 1 12c0 6-1 13-5 18m-257 24c-12 0-22-7-24-18c-1-7 0-15 4-20c8-9 25-9 33 0c4 4 8 13 6 27c0 0 0 1-1 2c-5 6-11 9-18 9m328 42c1 1 1 2 1 3c-3 14-17 30-32 30c-5 0-10-1-13-4c-2-1-3-3-2-5c-6 5-12 9-19 12h-2c-2 0-4-1-5-3c0 0 0-1-1-1c-2-1-4-3-3-6c2-16 16-29 31-29h7c1 1 4 2 4 3c1 1 1 3 0 4c0 2-1 3-2 5c9-9 22-14 33-12c1 1 2 1 3 3m31 4c2 1 3 3 3 4c0 5-1 10-4 13c-4 4-10 7-16 7s-13-3-16-6c-1-1-1-3-1-4c0-9 10-17 22-17c5 0 9 1 12 3m-455 50v-2l-1-1c-1 0-2-1-3-2c-1-2-1-3-1-4c3-13 17-26 30-26c2 0 4 0 5 1c3 0 5 3 4 6c-2 15-12 26-27 30h-1c-2 0-5-1-6-2m80-31v4c-6 12-14 27-31 33c-1 1-1 1-2 1c-2 0-5-1-6-3v-2c-2 0-3-1-4-2c0-1-1-4 0-5c4-15 16-30 32-30c2 0 5 0 7 1c2 0 3 1 4 3m34 1c2 0 3 0 4 1c1 2 2 3 1 5c-2 19-17 27-28 32c-1 1-2 1-2 1c-2 0-4-1-5-2c-5-6-6-13-3-19c5-11 21-18 33-18m174 14c4 0 8 1 12 4c2 1 4 3 3 5c-5 31-44 44-78 44c-37 0-64-15-70-40c0-1 0-4 2-5c4-3 9-5 14-5c7 0 15 5 20 10c4 3 8 7 12 8c15 5 32 4 49-2c5-2 10-6 14-9c7-5 14-10 22-10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func L(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 754V0h72v754z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LupperCase(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M71 0v663h383v72H0V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laugh(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M358 0c198 0 359 155 359 349S556 698 358 698S0 543 0 349S160 0 358 0M138 290l3 1c5 0 10-3 12-8c0-1 11-49 55-49c45 0 55 47 56 49c2 5 6 8 11 8l32-10c3-1 5-3 6-6c2-3 2-6 1-8c-1-4-20-85-106-85c-91 0-105 81-106 84s-1 6 1 9c1 3 4 5 7 6zm437 1l32-10c3-1 5-3 6-6c2-3 2-6 1-8c-1-4-20-85-105-85c-91 0-106 81-107 84s-1 6 1 9c1 3 4 5 7 6l28 9l3 1c5 0 10-3 12-8c0-1 11-49 56-49c44 0 55 47 55 49c2 5 6 8 11 8m-21 157c7-20-4-43-24-51s-325-8-344 0c-20 8-30 31-23 51c32 81 108 134 195 134s164-53 196-134"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Left(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m227 357l246 246l-119 114L0 363L362 0l113 112z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Light(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M380.297 39v79c0 27-12 40-37 40s-37-13-37-40V39c0-26 12-39 37-39s37 13 37 39m75 104l46-65c15-20 34-24 53-10s21 32 6 54l-46 64c-15 22-33 24-53 10c-19-14-21-31-6-53m-270-64l47 64c15 22 13 39-7 53c-19 15-37 12-52-10l-47-64c-15-22-13-39 7-53c19-15 37-12 52 10m158 124c86 0 156 70 156 156c0 85-70 154-156 154s-156-69-156-154c0-86 70-156 156-156m-292 22l76 24c25 8 34 23 26 47c-7 24-24 31-50 23l-74-24c-25-8-34-24-27-48c8-23 24-30 49-22m509 24l75-24c26-8 42-1 49 22c8 24-1 40-27 47l-74 25c-25 8-41 1-49-23c-7-24 1-39 26-47m-217 196c48 0 87-39 87-86c0-48-39-87-87-87s-86 39-86 87c0 47 38 86 86 86m-314-23l75-24c25-8 42-2 49 22c8 24-1 39-26 47l-76 25c-25 8-41 1-49-23c-7-24 2-39 27-47m554-24l75 24c25 8 34 23 26 47c-7 24-23 31-49 23l-75-25c-26-8-33-23-26-47c8-24 24-30 49-22m-457 196l47-63c15-22 33-25 52-10c20 14 22 31 7 53l-47 63c-15 22-33 26-53 12c-19-14-21-33-6-55m388-63l46 63c15 22 13 41-6 55s-38 9-53-11l-46-64c-15-22-13-39 6-53c20-14 38-12 53 10m-134 67v79c0 27-12 40-37 40s-37-13-37-40v-79c0-27 12-40 37-40s37 13 37 40"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Line(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 25c190 0 343 124 343 278c0 75-37 141-95 193c-94 93-256 187-273 180c-28-11 21-71 1-97c-3-4-14-3-34-6C118 551 0 442 0 302C0 148 152 25 341 25M146 395h58c10 0 16-8 16-18c0-7-5-16-18-16h-47c-4 0-4-4-4-18V242c0-9-6-18-17-18c-10 0-17 8-17 18v122c0 27 13 31 29 31m131-14V242c0-9-7-18-18-18c-9 0-17 8-17 18v139c0 9 7 16 18 16c9 0 17-7 17-16m164-19V240c0-9-7-16-18-16c-8 0-17 5-17 16v88l-58-81c-8-14-18-23-30-23c-15 0-16 12-16 29v128c0 9 7 16 17 16c9 0 17-6 17-16v-91l58 81c13 17 16 26 32 26c11 0 15-10 15-35m112-69h-51v-29c0-3 1-7 6-7h45c11 0 18-6 18-17c0-13-10-17-19-17h-58c-18 0-27 12-27 28v113c0 20 8 31 27 31h59c13 0 18-8 18-18c0-13-10-17-18-17h-44c-4 0-7-3-7-6v-28h53c13 0 15-10 15-17c0-10-8-16-17-16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m333 276l80-81c7-7 15-16 23-26c8-9 17-17 26-24c17-15 37-28 62-28c22 0 41 10 55 23s22 32 22 54c0 9-1 19-4 26c-7 14-13 24-19 32c-3 4-6 8-6 12c0 3 0 6 2 7c11 27 19 50 24 78c3 11 9 16 19 16c4 0 8-1 12-4c8-5 15-13 22-21c4-4 7-8 9-10c36-35 57-85 57-136c0-54-22-102-57-137c-35-34-84-57-136-57S422 21 387 58L250 194c-36 37-56 85-56 136c0 15 6 46 16 75c10 28 23 53 40 53c9 0 28-15 44-32s32-36 32-43c0-5-4-11-8-20c-5-9-7-20-7-33c0-20 8-40 22-54m-3 384l137-136c36-35 56-86 56-137c0-15-6-45-16-74c-9-28-23-54-40-54c-7 0-28 16-43 33c-17 17-33 35-33 42c0 5 3 13 8 21c5 9 9 19 9 32c-1 20-9 40-23 56l-81 80c-7 8-15 16-23 25l-25 25c-17 16-38 27-63 27c-42 0-76-33-76-76c0-10 2-19 5-26c6-14 12-24 19-32c3-4 4-8 4-11c0-2-1-4-2-8c-12-27-19-50-24-78c-2-5-3-9-6-11c-3-4-8-5-13-5c-4 0-7 1-11 4c-8 5-16 13-23 21c-3 4-7 7-9 9c-36 36-57 86-57 137c0 53 21 102 57 137c35 35 83 56 136 56c52 0 102-20 137-57"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linkedin(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 98v498c0 47 37 84 84 84h498c23 0 45-11 59-25c15-15 25-34 25-59V98c0-47-37-84-84-84H84c-25 0-44 10-59 25C11 53 0 75 0 98m90 66c0-32 26-60 58-60c33 0 60 28 60 60c0 33-27 59-60 59c-32 0-58-26-58-59m161 411V270c0-7 7-13 12-13h85c12 0 12 14 12 23c24-24 55-30 87-30c78 0 128 37 128 119v206c0 7-6 13-12 13h-88c-7 0-12-7-12-13V389c0-31-9-48-44-48c-44 0-55 29-55 68v166c0 7-7 13-14 13h-87c-5 0-12-7-12-13m-159 0V270c0-7 7-13 12-13h87c8 0 13 5 13 13v305c0 7-6 13-13 13h-87c-6 0-12-7-12-13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func List(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M129 118c0-36-28-64-64-64S0 82 0 118s29 65 65 65s64-29 64-65m75 53h564V67H204zm-75 176c0-36-28-65-64-65S0 311 0 347s29 64 65 64s64-28 64-64m75 52h564V295H204zm-75 176c0-36-28-64-64-64S0 539 0 575s29 65 65 65s64-29 64-65m75 51h564V523H204z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Location(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M333 702L666 0L0 356h333z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M468 332h67c9 0 16 8 16 17v352c0 9-7 16-16 16H16c-9 0-16-7-16-16V349c0-9 7-17 16-17h65V168C81 76 140 0 231 0h86c92 0 151 76 151 168zm-302 0h218V176c0-51-42-94-94-94h-31c-51 0-93 43-93 94zm63 283h93l-23-101c16-8 27-25 27-44c0-29-23-51-51-51s-50 22-50 51c0 19 11 36 27 44z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Login(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M391 672h195c37 0 68-13 92-38c24-24 39-56 39-92V151c0-36-15-68-39-92s-56-38-92-38H391v82h195c28 0 49 21 49 48v391c0 27-21 49-49 49H391zM0 269v156c0 18 15 33 33 33h182v123c0 11 5 20 15 25c4 1 9 1 11 1c7 0 13-2 18-7l235-235c9-9 8-27 0-37L259 94c-8-8-18-9-29-6c-10 5-15 13-15 24v124H33c-18 0-33 15-33 33"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Logout(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 151v391c0 36 15 68 39 93c24 24 55 37 91 37h196v-81H130c-27 0-48-22-48-49V151c0-27 21-48 48-48h196V21H130c-36 0-67 14-91 38c-24 25-39 56-39 92m215 118v156c0 18 16 33 34 33h181v123c0 11 6 20 16 25c4 1 8 1 10 1c7 0 13-2 18-7l235-235c11-9 10-27 0-37L474 94c-14-15-44-6-44 18v124H249c-18 0-34 15-34 33"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ltthon(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M757 351c3 49-14 78-27 114c-9 25-20 50-32 73c-5 2-9-8-16-12c-6-5-11-10-23-9c13-18 26-40 35-68c8-27 16-53 24-81c6-22 16-39 14-57c-2-20-18-31-28-45c-4-21-12-39-22-56c-19-34-44-71-73-98c-48-43-96-79-183-73c-20 2-40 6-60 13c-6 1-11 5-16 6c-12 0-24-8-37-6c-12 1-22 8-28 14c-12 11-11 32-29 40c-10-13-18-29-28-42c-4-4-10-6-12-12C285 25 343-6 441 1c94 6 157 57 212 121c49 57 96 135 104 229M121 210c-14-7-50-39-45-63c3-13 25-19 32-29c14-17 16-41 29-48c15-10 45-10 61 0c1 1 2 3 4 4c12 8 37 48 34 65c-2 10-7 8-16 18c-8 4-10 13-18 16c0 8 0 15-2 20c-23 13-50 32-79 17m178 36c-7 0-14 1-16-4c8-44 39-65 52-103c-16-16-39-12-40-35c-1-11 8-24 22-26c18-3 65 26 75 36c6 6 13 19 12 25c-1 8-7 7-12 14c-31-3-49 6-51 32c16 5 23-7 47-4c16 2 33 13 43 27c13 18 16 36 6 62c-30 18-68 27-112 31c-4-8-1-22-2-33c-8-7-24-7-24-22m67 16c-1-4 2-12-4-10c-6 5-18 3-19 12c8 8 14 0 23-2m-12-61c-13 4-37 16-35 25c9-1 15-10 26-12s22 2 31 0c1-16-14-15-22-13m239-52c13 2 18 19 24 22c6 10 17 14 16 30c-5 6-16 6-24 9c-12 18-39 22-41 50c-8 8-17 14-26 20c-2 11-8 17-12 25c2 23 3 50-6 67c-20 10-28-13-39-29c-3-5-7-12-12-18c-7-10-16-12-16-26c-1-14 5-52 10-63c8-17 28-24 30-39c-10-2-24 3-36 0c-16-3-32-15-32-28c-1-26 28-19 50-14c9 2 19 5 26 6c22 3 46 0 61-4c12-3 16-10 27-8M66 256c-11-2-22-13-22-24c-1-11 9-27 20-29c12-1 15 7 26 13c52 15 90-13 144-8c16 7 33 15 34 36c2 16-8 33-18 45c22 2 37 10 45 26c-5 21-7 49 4 65c35 0 44-26 67-37c10-20 16-44 46-44c17 19-3 33-16 50c-9 12-13 25-22 35c-17 18-48 25-71 36c-10-3-24-2-37-2c-9-8-19-12-24-26c-5-15 0-36-10-49c-12-3-9 9-14 12c-5 5-12 7-18 11c-4 10-17 17-25 32c-2 4-2 11-4 16c-8 22-18 53-56 41c-1-15-3-29 2-43c2-8 11-13 14-20c4-12 4-27 8-39c9-31 29-63 40-97c-17-2-28-13-52-12c-23 1-41 16-61 12m55 327c0-1 0-2 1-3l5 6l-1 1h-3c0-1-1-1-1-2c-1 0-1-1-1-2m-22-47c0 1 1 1 1 2c1 0 1 1 1 2v1l-2 2l-1 1c-1 1-1 1-2 1s-1 1-2 1s-1-1-2-1l-1-1c-1-1-1-1-1-2v-2c1 0 1-1 2-1c0-1 1-1 1-2c1 0 2 0 2-1zM40 358c1 0 1 1 1 1c1 1 1 1 1 2v2c0 1-1 1-1 2c-1 0-1 0-2 1h-4c-1 0-2 0-2-1c-1 0-1 0-2-1v-4l1-1c1-1 1-1 2-1s1-1 2-1h2c1 1 2 1 2 1m4 56h-4l-1-1c-1-1-1-1-1-2c-1-1-1-2 0-2c0-1 0-1 1-2l1-1c1 0 2 0 2-1h2c1 0 2 0 2 1c1 0 1 0 2 1c0 0 1 1 1 2v2l-1 1c-1 1-1 1-2 1c-1 1-1 1-2 1m19 61c0-1 0-1-1-1v-1l1-1h3s1 0 1 1c1 0 1 1 2 2v3h-1l-1 1c-1 0-1-1-1-1h-1c0-1-1-1-1-1v-1zm53 86c1 0 1 1 2 2c0 0 0 1 1 1c0 1 0 2-1 2c0 1 0 1-1 2c0 1-1 1-1 1l-2 2h-4l-1-1c-1-1-1-2-1-2v-2c0-1 1-1 1-2c1 0 1-1 2-1c0-1 1-1 1-1c1-1 2-1 2-1zm28 47c0-1 0-2 1-3l6 5l-1 1h-3s-1 0-2-1s-1-2-1-2m56 42c1 0 1 0 2 1l4 2l-4 6l-4-2l-1-1c-1 0-1-1-1-1c-1-1-1-1-1-2c0 0 0-1 1-2l1-1zm281-3c19 21 70 29 106 16c-12 13-31 10-45 14c-7 3-13 7-20 11c-44 20-111 32-175 32c-47 0-100-10-143-32c-2-1-4-5-6-6c-7-4-16-6-23-11c-4-2-7-7-12-10c-4-3-10-3-14-6c-8-5-17-13-24-20c-22-19-37-40-55-67C33 512-6 453 1 353c2-31 13-64 24-95c15 23 56 32 85 16c-6 23-11 52-10 83c1 8 3 17 2 25c-1 10-9 18-12 28c-8 33-5 72 33 73c27 85 77 147 160 176c11 14 26 42 50 45c28 3 53-11 57-41c17 7 30 20 55 12c8-2 14-7 20-12s9-13 16-16M55 393s0-1-1-1c-1-1-1-1-2-1h-2l-19 3l1 5l18-3h2v3l4-1v-2s0-2-1-3m1 31l-1-4l-18 4l1 5zm-21-9l3 3c1 0 2 1 3 1h4c1-1 3-1 4-2c1 0 2-1 2-2c1-1 1-2 2-3c0-2 0-3-1-4c0-2 0-3-1-4s-1-2-2-2c-1-1-2-1-4-2c-1 0-2 0-4 1c-1 0-2 0-3 1s-2 1-3 2c0 1-1 2-1 4v4c0 1 1 2 1 3m11-55c0-2-1-3-1-4c-1-1-2-1-3-2s-2-1-3-1c-1-1-2-1-3-1c-1 1-3 1-4 1c-1 1-2 1-3 2c0 1-1 2-2 3v3c0 2 0 3 1 3c0 1 1 2 2 2v1h-2v4l24-2v-4h-9c1-1 2-1 2-2c1-1 1-2 1-3m-4 19h5v-6l-6 1zm-16-53v3s0 2 1 3c0 1 1 2 2 2c-1 1-2 1-3 2v3c0 2 0 3 1 4l2 2h-3v5l18-1v-5l-10 1c-1 0-2-1-2-1c-1 0-1 0-2-1v-4h1l1-1h12v-5H31c0-1-1-1-1-1c0-1-1-1-1-2s1-1 1-2l1-1h13v-5H32c-1 0-2 0-3 1c-1 0-2 0-2 1c-1 1-1 1-1 2m1 59l-4 1l1 5l4-1zm-2 10v5l4-1v-5zm23-8v-5l-18 3l1 5zm-17 38l1 5l4-1l-2-5zm14 18l-9 3l1 5l24-7l-1-5l-9 3h-4s-2-1-2-2c-1-2-1-3 0-3c1-1 2-1 3-2l10-2l-2-5l-10 3c-1 0-2 0-3 1c-1 0-1 1-2 1c0 1-1 2-1 3s0 2 1 3c0 1 0 2 1 3c1 0 2 1 3 1m4 21c1 1 2 1 3 2h7c1-1 2-1 3-2s2-2 2-3c1-1 1-2 1-3v-4c-1-2-2-4-4-5c-1-1-4-1-6-1l2 5h3c0 1 1 1 1 2c1 1 1 2 0 2v2c-1 1-1 1-2 1c0 1-1 1-2 1c0 0-1 1-2 1h-2c0-1-1-1-1-1l-2-2c0-2 0-3 2-4l-1-5c-1 0-2 1-3 2c-1 0-1 1-1 2c-1 1-1 2-1 3s1 2 1 3c0 2 1 3 2 4m5 8c-1 1-1 1-1 2v3l1 2c0 1 1 2 1 3c1 1 1 2 2 2c1 1 2 1 3 1c0 1 1 0 3 0l-2-5c-1 1-2 1-3 0c0 0-1-1-1-2v-3h1c1-1 1-1 1 0c1 0 1 0 1 1c1 0 1 1 1 1c1 1 1 2 1 2l2 2c0 1 1 1 1 2h5s1-1 2-1c0-1 1-2 1-2v-2c0-1 0-2-1-3c0-1 0-2-1-3c-1 0-2-1-3-2h1l1-1l-2-4l-1 1c-1 0-2 1-3 1l-8 3c-1 1-2 1-2 2m5 36l3 5l22-10l-2-5l-10 4l-4-9l10-4l-3-5l-22 10l2 5l9-4l4 9zm30 15l-9 9l3 5l11-14h1c0-1 1-1 1-1c1 0 1 1 2 1v1c1 1 1 1 1 2l4-2c-1-1-1-1-1-2l-1-1l-3-3c-1 0-2-1-3 0l-21 3l2 4zm3 25l-8 5l3 4l21-13l-3-4l-2 1c1-1 1-2 1-3c0-2-1-3-1-4c-1 0-2-1-2-2c-1 0-2-1-3-1h-3c-2 1-3 1-4 2c-2 1-3 2-4 3c0 1-1 2-1 3v3s0 2 1 2c0 1 1 2 2 3c1 0 2 1 3 1m20 28c1 0 3 0 4-1c1 0 2-1 3-2c1 0 2-1 2-2c1-1 1-2 2-3v-3c0-2-1-3-2-4c0-1-1-1-2-2h-3l2-2l-3-4l-20 15l3 4l7-5v3s1 2 2 3c0 1 1 2 2 2c1 1 2 1 3 1m10 14c1 1 2 1 3 1s3 0 4-1c1 0 2-1 3-1c1-1 2-2 2-3c1-2 1-3 1-4c1-1 0-2 0-3s-1-2-2-4c-1-1-3-2-4-3c-2 0-4 0-6 1l3 3h2l2 2s1 2 1 3s0 2-1 3l-9-10l-3 3c0 1-1 2-1 4v3s1 3 2 4s2 2 3 2m11 8l-2 3l3 3l13-12l-3-4l-6 6c-1 0-1 1-2 1c0 0-1 1-2 1c0 0-1 0-2-1c0 0-1 0-2-1v-1h-1l-3 3s1 0 1 1l1 1h2c0 1 1 1 2 1zm13 17h3s3 0 4-1c1 0 2-1 3-2s1-2 2-3c0-1 1-2 1-4c0-1-1-2-1-3c-1-1-1-2-3-3c-1-1-3-2-4-3c-2 0-4 1-6 2l3 3c1-1 2-1 2-1c1 0 2 1 3 2c1 0 1 1 1 3c1 1 0 2-1 3l-9-9l-3 3s-1 2-1 4v3c1 2 2 3 3 4s2 1 3 2m22 19l4 3l7-17l-4-3l-10 7l5-11l-4-4l-15 11l3 3l11-7l-6 11l4 3l10-8zm16 11c2 1 3 1 4 0c1 0 2 0 3-1l3-3s1-2 2-3v-4c-1-1-1-2-2-3c0-1-1-2-2-3c-2-1-3-1-4-1c-1-1-2-1-3-1c-2 1-3 1-4 2s-2 1-2 3c-1 1-2 2-2 3v4s0 2 1 3l3 3c1 0 2 1 3 1m11 15l9 6l13-22l-5-2l-4 7l-5-3c-2 0-3-1-4-1h-3s-2 1-3 2c-1 0-1 1-2 2c0 1-1 2-1 3v3l2 2c0 1 1 2 3 3m-4-27v3c0 1-1 2-1 2l-2 2l-1 1h-4c-1-1-1-1-1-2c-1-1-1-1-1-2c0 0 0-1 1-2c0 0 0-1 1-2l1-1c1-1 1-1 2-1c0-1 1-1 1-1c1 0 2 1 2 1c1 1 2 1 2 2m475-249c-8-12-3-32-5-49c-1-15-1-32 9-42c7 2 11-3 16-2c20 2 24 32 14 52c-3 6-9 9-12 14c-4 9-3 18-6 23c-7 0-9 4-16 4m-59-31c-1 13 6 25 2 37c-24 15-26-9-35-31c-5-14-14-27-12-40c4-26 29-21 51-8c4 13-6 27-6 42m50 189c16-1 19 16 13 36c-18 9-31 29-51 43c-37 25-88 32-128 4c-8-18-11-41-14-65c-27 5-10 54-22 77c-10 18-39 19-51 0c4-38 19-97 22-146c2-18 2-35 7-49c4-12 31-47 44-50c23-6 33 8 41 22c7 45-24 52-27 87c10 8 22 4 31 6c9 3 22 8 28 17c17 22-4 63 24 72c18-4 44-11 59-28c6-7 13-26 24-26M396 394c19 13 2 34-10 53c-25 37-43 76-67 115c19 9 58-1 69 16c0 8 0 15-2 20c-14 9-42 4-51 19c-4 29 24 26 29 46c-3 9-14 13-21 12c-9 0-33-14-38-28c-6-15 1-29 0-44c-11-10-30 5-45-3c-5-5-5-18-2-26c8-10 32-3 39-14c-38-2-61-19-79-41c-5-32-20-73 8-91c29-2 38 17 57 25c10 24 17 53 28 77c8-6 10-17 14-27c4-8 11-14 14-24c7-17 7-40 15-57c1-3 13-21 18-24c7-4 15-4 24-4m150 71v-20h4v16h10v4zm23-17v17h-5v-17h-6v-3h17v3zm6 7h8v4h-8zm18-7v17h-4v-17h-6v-3h16v3zm26-3v20h-5v-9h-8v9h-4v-20h4v7h8v-7zm9 0c1-1 2-1 4-1c1 0 3 0 4 1c1 0 2 1 3 2s2 2 2 4c0 1 1 2 1 4c0 1-1 3-1 4s-1 2-2 3s-2 2-3 2c-1 1-3 1-4 1c-2 0-3 0-4-1c-2 0-3-1-4-2c0-1-1-2-2-3v-8c1-2 2-3 2-4c1-1 2-2 4-2m8 5c0-1-1-1-2-2h-5l-2 2v2c-1 1-1 2-1 3s0 1 1 2v2c1 1 1 2 2 2s2 1 3 1s2-1 2-1c1 0 2-1 2-2c0 0 1-1 1-2v-5s-1-1-1-2m22 8v-13h4v20h-5l-8-14v14h-4v-20h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func M(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M71 128v64c42-40 109-64 171-64s110 24 152 64c13 13 26 28 36 44c41-66 124-108 207-108c62 0 108 24 150 64c45 42 73 102 73 168v294h-73V356c-2-87-62-157-150-157s-170 71-171 159v296h-72V356c-3-87-64-157-152-157c-87 0-169 70-171 157v298H0V128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MupperCase(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 0l310 443L620 0v753h-73V229L310 568L73 229v524H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Magic(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m304.188 151l-29 93l-28-93l-93-29l93-28l28-94l29 94l93 28zm162-51l-14 47l-16-47l-46-15l46-14l16-46l14 46l46 14zm-457 511l573-556c4-5 10-8 17-8c9-1 18 2 25 9c3 2 7 6 13 11c11 10 23 22 32 34c4 6 6 11 6 16c1 10-2 19-8 25l-573 558c-5 5-11 7-18 8c-9 1-18-2-25-9c-3-3-7-6-13-12c-11-9-23-21-31-34c-4-6-6-12-7-17c-1-10 2-18 9-25m460-404l42 43l133-129l-42-44zm-279 57l-13 46l-14-46l-46-15l46-14l14-47l13 47l48 14zm466-35l15 47l46 14l-46 15l-15 46l-15-46l-46-15l46-14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mail(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 103L384 432L0 103V84h768zm0 464L514 384l254-220zM0 164l254 220L0 567zm384 333l105-92l279 201v4H0v-4l279-201z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Male(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M224 76c0-42-34-76-76-76c-41 0-75 34-75 76s34 76 75 76c42 0 76-34 76-76M25 177h253c14 0 25 11 25 25v164c0 14-11 26-25 26s-25-12-25-26v-75c0-14-7-26-16-26s-16 12-16 26v404c0 14-11 25-25 25s-25-11-25-25V518c0-14-9-25-19-25c-11 0-19 11-19 25v177c0 14-12 25-26 25s-25-11-25-25V291c0-14-7-26-16-26c-8 0-15 12-15 26v75c0 14-12 26-26 26S0 380 0 366V202c0-14 11-25 25-25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Map(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 637l220-83V57L0 140zm469 0l-221-83V57l221 83zm28 0l220-83V57l-220 83z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Meal(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M201 248V14h67v268c0 37-30 68-67 68v296c0 18-15 34-34 34h-67c-18 0-33-16-33-34V350c-37 0-67-31-67-68V14h67v234h33V14h67v234zm168 102V114c0-55 45-100 100-100s100 45 100 100v532c0 18-14 34-32 34h-68c-18 0-34-16-34-34V350z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Memo(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M158 0h401c63 0 116 53 116 116v485c0 63-53 116-116 116H158c-63 0-116-53-116-116v-11H22c-12 0-22-9-22-20s10-22 22-22h20V380H22c-12 0-22-10-22-21s10-22 22-22h20V169H22c-12 0-22-10-22-21s10-21 22-21h20v-11C42 53 95 0 158 0m453 601V116c0-29-23-52-52-52H158c-29 0-53 23-53 52v11h42c11 0 22 10 22 21s-11 21-22 21h-42v168h42c11 0 22 11 22 22s-11 21-22 21h-42v168h42c11 0 22 11 22 22s-11 20-22 20h-42v11c0 29 24 53 53 53h401c29 0 52-24 52-53M253 127h253c12 0 21 10 21 21s-9 21-21 21H253c-11 0-22-10-22-21s11-21 22-21m0 84h253c12 0 21 10 21 21c0 12-9 21-21 21H253c-11 0-22-9-22-21c0-11 11-21 22-21m0 84h253c12 0 21 10 21 22c0 11-9 20-21 20H253c-11 0-22-9-22-20c0-12 11-22 22-22m0 85h253c12 0 21 10 21 21c0 12-9 21-21 21H253c-11 0-22-9-22-21c0-11 11-21 22-21m0 84h253c12 0 21 10 21 22c0 11-9 20-21 20H253c-11 0-22-9-22-20c0-12 11-22 22-22m0 84h253c12 0 21 11 21 22s-9 20-21 20H253c-11 0-22-9-22-20s11-22 22-22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Menu(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M44 128h578c25 0 44-19 44-44c0-24-19-43-44-43H44C19 41 0 60 0 84c0 25 19 44 44 44m0 175h578c25 0 44-19 44-44c0-24-19-43-44-43H44c-25 0-44 19-44 43c0 25 19 44 44 44m0 175h578c25 0 44-19 44-43c0-25-19-44-44-44H44c-25 0-44 19-44 44c0 24 19 43 44 43m0 175h578c25 0 44-19 44-43c0-25-19-44-44-44H44c-25 0-44 19-44 44c0 24 19 43 44 43"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 268h758v158H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mixi(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M352.45 17c200-13 376 119 363 319c-12 193-177 293-335 354v-71c-209 15-399-86-379-301c15-166 181-289 351-301m162 449h51V308c0-41-3-71-20-91c-18-20-31-31-75-31c-33 0-67 17-96 53c-19-35-39-53-83-53c-28 0-69 16-91 51v-46h-52v275h52V324c0-43 20-67 36-79c11-8 27-13 39-13c20 0 36 6 44 19c8 12 13 23 13 67v148h51V312c0-43 37-80 71-80c33 0 42 11 52 23c7 8 8 29 8 53z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mobage(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M654 252C769 110 632 0 632 0s34 107-55 147c-61-57-144-91-234-91C154 56 0 209 0 397c0 190 154 343 343 343c190 0 344-153 344-343c0-51-12-101-33-145m-98 206c-27 92-111 159-213 159c-121 0-220-98-220-220c0-99 66-182 157-210c-6 20-9 40-9 62c0 121 98 219 220 219c22 0 44-3 65-10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Move(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m224 593l116 116c5 5 11 8 18 7c8 0 15-2 20-7l115-116c16-16 11-28-12-28h-78V401h162v80c0 23 14 27 30 11l115-115c5-5 7-11 7-19c1-7-2-14-7-20L595 224c-16-16-30-12-29 11v78H403V151h78c23 0 28-13 12-29L378 8c-10-10-28-11-38 0L224 122c-16 16-11 29 12 29h78v161H153v-77c0-23-14-27-30-11L8 338c-5 5-7 12-8 20c0 8 3 14 8 19l115 115c16 16 29 12 29-11v-79h162v163h-78c-23 0-28 12-12 28"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Music(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M182 156v349c-15-4-31-7-45-7c-32 0-65 11-92 29S0 574 0 608c0 36 18 63 45 81s60 28 92 28c30 0 64-10 91-28s44-45 44-81V242l305-87v241c-14-4-31-6-46-6c-31 0-65 9-92 27s-44 47-44 81c0 36 17 64 44 81c27 18 61 29 92 29c30 0 65-11 92-29c28-17 45-45 45-81V37c0-22-15-37-37-37c-2 0-14 4-33 9c-40 11-100 28-168 48c-35 9-66 19-97 27c-30 9-56 18-79 24c-24 7-39 11-45 13c-16 5-27 19-27 35"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Myspace(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M436 247c-59 0-106-47-106-106S377 34 436 34s107 48 107 107s-48 106-107 106m-194-18c-43 0-78-35-78-78s35-78 78-78c44 0 79 35 79 78s-35 78-79 78m340 183c0 3 0 7-1 10h1v252H291V542H136v-66H0V319c0-48 39-87 87-87c34 0 62 18 77 46c20-21 47-35 78-35c43 0 80 25 97 62c26-24 60-38 97-38c81 0 146 65 146 145M87 220c-35 0-64-28-64-64c0-35 29-64 64-64c36 0 64 29 64 64c0 36-28 64-64 64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func N(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M72 128v64c42-40 109-64 171-64s109 24 151 64c44 42 72 102 72 168v294h-72V356c-2-87-64-157-151-157c-88 0-169 70-171 157v298H0V128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NupperCase(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M455 571V20h72v756L72 206v550H0V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Next(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 628V65l489 282z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Nine(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m82 743l186-292c-13 2-27 4-41 4C101 455 0 353 0 228C0 102 101 0 227 0c105 0 194 73 220 171c4 17 8 38 8 57c0 43-13 82-33 116l-7 11l-273 427zm266-417l19-29c11-22 16-44 16-69c0-86-70-157-156-157S71 142 71 228s70 155 156 155c49 0 92-22 121-57"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notify(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 14h679v500H435l-96 166l-95-166H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Numbersign(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M555 232v60H435l-24 161h114v60H402l-36 246h-60l35-246H180l-37 246H82l37-246H0v-60h128l24-161H31v-60h130L195 0h60l-33 232h161L418 0h61l-35 232zM189 453h162l23-161H212z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func O(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M285 113c152 0 275 121 275 271S437 654 285 654S0 534 0 384s133-271 285-271m-1 474c114 0 207-90 207-203s-93-204-207-204c-115 0-215 91-215 204s100 203 215 203"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OupperCase(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M788.018 318c37 211-107 412-324 450c-216 38-421-101-458-312c-38-210 108-411 324-449s421 101 458 311m-391 383c178 0 323-140 323-313s-145-315-323-315s-322 142-322 315s144 313 322 313"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Off(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M459 83c118 50 201 167 201 304c0 181-148 330-330 330C149 717 0 568 0 387C0 250 83 133 202 83v111C140 236 98 306 98 387c0 128 104 232 232 232s232-104 232-232c0-81-41-151-103-193zm-74 302V0H275v385z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func One(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h175v735h-73V71H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Opera(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M330 0c197 0 337 145 337 357c0 197-140 360-337 360S0 554 0 357C0 149 133 0 330 0m1 636c104 0 126-120 126-274c0-185-28-283-126-283c-113 0-122 132-122 297c0 154 29 260 122 260"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ordble(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M356 0c96 0 179 35 249 105c70 71 105 156 105 256c0 99-35 184-104 253S454 717 357 717c-102 0-188-35-256-106S0 456 0 359c0-66 15-126 47-180s75-98 130-130c5-3 11-7 17-9c0 0 58 35 57 94c-51-25-105 28-105 28s58-16 93 9c-85 50-130 153-103 253c32 121 156 193 277 160s192-156 159-276c-30-113-138-182-250-166c8-39 50-81 50-81s-44-6-79 57c-11-69-98-78-99-78C244 13 298 0 356 0M222 238c23-23 54-29 69-14c14 15 8 45-16 68c-23 24-53 30-68 15c-14-14-9-45 15-69"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Picon(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M72 113v97c53-58 132-97 216-97c152 0 275 121 275 271S440 654 288 654c-84 0-163-38-216-96v309H0V113zm215 474c115 0 207-90 207-203s-92-204-207-204S72 271 72 384s100 203 215 203"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PupperCase(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h279c58 0 109 27 141 71c22 30 35 66 35 105s-13 75-35 104c-32 44-83 72-141 72H72v383H0zm72 280h210c56-1 101-47 101-104c0-56-45-103-101-105H72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paint(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M326 337h82V28h-82zm131-87l85 85c14 14 14 36 0 50L275 651c-13 14-36 14-49 0L10 435c-13-13-13-36 0-49l266-267v259h181zm74 24s67 27 67 159c0 82-9 95-9 95s128 14 128-111c0-126-186-143-186-143"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Palette(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M767.074 148c15-50-112-130-227-135c-223-9-515 108-539 365c-18 206 194 312 350 312c157 0 306-121 322-187c16-68-143-49-132-157c16-153 204-125 226-198m-236-35c0 26-20 46-45 46c-26 0-46-20-46-46c0-25 20-45 46-45c25 0 45 20 45 45m-124 24c0 28-22 50-49 50s-49-22-49-50c0-26 22-48 49-48s49 22 49 48m-119 65c0 29-22 52-51 52c-28 0-51-23-51-52c0-27 23-49 51-49c29 0 51 22 51 49m-100 114c0 29-24 53-53 53c-31 0-55-24-55-53c0-31 24-55 55-55c29 0 53 24 53 55m275 192c0 38-32 69-71 69c-38 0-69-31-69-69c0-39 31-72 69-72c39 0 71 33 71 72"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paperboy(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M576 288C576 129 447 0 288 0S0 129 0 288c0 118 73 223 181 266c13 34 52 129 52 129c9 22 31 37 55 37s46-15 55-37c0 0 39-95 53-129c107-43 180-148 180-266M288 601c-18-45-42-101-42-101c-2-7-7-12-14-13c-88-25-150-107-150-199c0-114 92-206 206-206s206 92 206 206c0 92-61 174-150 199c-6 1-12 6-14 13c0 0-23 56-42 101"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paramater(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M414 74v37h303v60H414v37c0 11-10 20-21 20h-69c-11 0-21-9-21-20v-37H0v-60h303V74c0-11 10-20 21-20h69c11 0 21 9 21 20M125 260h71c11 0 19 9 19 20v37h502v59H215v37c0 11-8 21-19 21h-71c-11 0-21-10-21-21v-37H0v-59h104v-37c0-11 10-20 21-20m389 206h71c11 0 20 9 20 20v37h112v59H605v37c0 11-9 21-20 21h-71c-11 0-20-10-20-21v-37H0v-59h494v-37c0-11 9-20 20-20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Parenleft(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M182 0h-81C68 76 43 158 26 247S0 427 0 520c0 87 8 171 22 252c15 80 34 150 59 209h77c-26-61-47-134-62-215c-14-82-21-167-21-255c0-95 9-188 28-277c20-89 46-166 79-234"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Parenright(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 981h81c33-76 59-158 76-247s25-180 25-273c0-87-7-170-21-251S126 59 101 0H25c26 62 46 134 61 216c14 81 22 166 22 254c0 95-10 187-29 276S33 914 0 981"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 45v604c0 17 14 31 31 31h212c16 0 29-14 29-31V45c0-17-13-31-29-31H31C14 14 0 28 0 45m393 0v604c0 17 14 31 30 31h212c17 0 31-14 31-31V45c0-17-14-31-31-31H423c-16 0-30 14-30 31"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pc(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M59 4h598c33 0 60 26 60 59v382c0 33-27 61-60 61H424v99h170c25 0 27 82 30 85H93s0-85 30-85h170v-99H59c-33 0-59-28-59-61V63C0 30 26 4 59 4m0 441h598V63H59z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pencil(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M587 0L478 98s36 22 69 53c32 31 55 66 55 66l93-111s-1-20-45-63C605 0 587 0 587 0M232 549l-5 48l346-360s-23-32-51-59c-29-28-62-50-62-50L114 488l54-3l6 58zM0 702l86-176l48-4l4 55l58 8l-6 44l-174 88z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Percent(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M62 819L596 0h77L137 819zM174 15c95 0 173 78 173 173c0 96-78 174-173 174C78 362 0 284 0 188C0 93 78 15 174 15m0 285c61 0 111-51 111-112S235 78 174 78S63 127 63 188s50 112 111 112m402 142c96 0 174 78 174 174c0 95-78 174-174 174c-95 0-173-79-173-174c0-96 78-174 173-174m0 285c61 0 111-50 111-111s-50-111-111-111s-111 50-111 111s50 111 111 111"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Period(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 577h77v77H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Periodcentered(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 311h77v77H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 179c0 30 11 63 25 98s32 68 46 93c38 66 87 131 143 185s122 99 196 126c43 17 86 36 134 36c24 0 58-9 89-22c30-13 57-33 67-57c4-9 9-21 12-35c3-13 5-26 5-35c0-5 0-10-1-11c-3-6-10-11-18-16c-18-9-21-12-48-28c-28-16-57-34-82-47c-12-7-20-10-23-10c-16 0-36 22-53 46c-18 24-38 46-52 46c-7 0-14-3-22-8s-16-9-21-12c-88-50-147-109-197-196c-6-10-20-31-20-42c0-13 18-28 36-44c17-15 35-33 35-55c0-3-2-12-6-25c-9-26-21-56-31-85c-5-14-9-24-10-29c-2-3-3-7-4-12s-3-9-4-13c-3-9-8-16-13-20c-5-2-17-5-29-6c-12 0-26-1-32-1c-3 0-7 0-11 1H98C66 15 43 43 26 76C10 109 0 146 0 179"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Photo(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M522 686V114H0v572zM47 551V161h428v390zM666 8v571H548V445h70V55H190v38h-47V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Picasa(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M503 271L231 24C271 9 313 0 358 0c51 0 101 12 145 31zM329 163L12 453c-8-30-12-61-12-94C0 221 78 100 193 41zm213 314V51c104 62 175 177 175 308c0 41-8 81-20 118zM171 360v304C105 624 55 564 26 492zm38 155h472c-58 119-181 202-323 202c-53 0-104-13-149-33z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pin(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M217 0c116 0 209 95 209 209c0 115-213 559-213 559S0 324 0 209C0 95 93 0 209 0zm-4 287c52 0 95-44 95-97c0-52-43-95-95-95c-53 0-96 43-96 95c0 53 43 97 96 97"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pinterest(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m298 607l-5 18c-7 29-23 56-37 78c34 9 67 14 102 14c99 0 188-39 254-104c65-65 105-155 105-254s-40-189-105-254C546 40 457 0 358 0S170 40 105 105C39 170 0 260 0 359c0 74 21 142 60 198c38 56 92 102 154 129c-1-25 1-53 7-80c1-8 4-20 8-36c3-14 8-34 14-60c6-25 14-58 23-100c-3-6-5-12-6-19c-4-13-6-25-5-38c0-52 30-93 70-93c33 0 48 25 48 54c0 16-5 35-12 58c-8 24-15 45-20 69c-5 19 0 36 11 49c11 12 28 20 46 20c34 0 62-21 82-56s32-83 32-135c0-39-13-72-39-97s-63-39-111-39c-54 0-98 19-129 51s-48 75-48 121c0 29 9 51 25 69c5 8 7 11 5 20c-1 3-3 8-4 14c-3 13-4 30-22 24c-50-20-73-76-73-138c0-51 21-106 64-149c43-44 108-74 192-74c69 0 126 24 166 63c40 38 61 90 61 142c0 71-19 132-53 176s-83 70-140 70c-39 0-75-20-88-45c-6 22-10 38-13 50c-3 13-5 24-7 30"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plane(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m392 123l-15 15l128 11l77-77c12-12 34-30 66-50s53-26 61-18c22 22 0 65-66 130l-76 76l11 128l15-14c12-13 25-13 35-3c12 12 7 29-17 53l-24 25l12 86l17-18c11-12 23-12 34-2c12 12 10 27-6 43l-36 36c14 96 16 150 5 161c-5 4-12 8-22 10c-9 1-17 0-22-4c-4-5-10-23-16-54c-7-32-24-86-54-162c-29-77-50-120-60-131c-4-4-10-4-17 0c-7 3-32 24-72 63c-40 40-77 72-112 97c-1 6 2 26 7 59s7 55 6 63c-2 8-10 19-23 32c-7 7-12 11-17 12c-13-13-33-55-60-124c-75-33-116-54-124-62c0-4 4-10 11-17c13-13 24-21 33-22c8-2 29 1 62 7c33 5 53 7 60 7c24-36 69-89 139-159c20-22 26-36 19-43c-10-11-54-31-130-60c-77-29-130-48-162-54c-32-7-50-12-54-17c-5-4-6-12-4-22c1-10 5-17 10-22c10-10 64-9 161 5l35-36c16-15 31-17 43-5c11 10 10 22-1 33l-18 17l86 12l24-23c24-24 41-29 53-17c11 11 10 22-2 35"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 49v605c0 11 5 20 15 26c4 3 11 4 16 4s9-1 14-4l524-302c10-7 16-16 16-27c0-10-6-19-16-26L45 23C27 10 0 26 0 49"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Playmedia(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 359c0 99 39 189 105 254c65 65 154 104 253 104s188-39 254-104c65-65 105-155 105-254s-40-189-105-254C546 40 457 0 358 0S170 40 105 105C39 170 0 260 0 359m108 0c0-70 28-132 73-177s107-74 177-74s133 29 178 74c46 45 73 107 73 177s-27 132-73 177c-45 45-107 74-178 74c-70 0-132-29-177-74s-73-107-73-177m153-140v281c0 5 2 9 7 12c2 1 5 2 7 2s6-1 8-2l244-141c4-2 8-6 8-12s-4-10-8-12L283 205c-3-2-9-2-15 0c-5 3-7 9-7 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M438 279h279v159H438v277H279V438H0V279h279V0h159z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pointer(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M298 14v42h-85V14zm0 170V56h42v128h170v42h-42v85h-43v-85h-85v85h-42zm-170 85v42H43v-42zm425 42h-43v-85h85v43h-42zm42 0v-42h43v212h-43zM0 311h43v85H0zm170 85v-42h-42v-43h42V56h43v340zM43 439v-43h42v43zm255 42v-85h42v170h-42zm85 0v-85h42v170h-42zM85 439h43v85H85zm425 0v127h-42V396h42zm43 127v-85h42v85zm-425 0v-42h42v42zm42 43v-43h43v43zm340 43v-86h43v128h-85v-42zm-85 0v-43h43v43zm-127 0h127v42H213v-85h42v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Present(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M454 194h121c29 0 53 23 53 53v182H340V217c-10-3-19-7-25-13c-1 0-1-1-2-1c0 0 0 1-1 1c-7 6-15 10-24 13v212H0V247c0-30 23-53 52-53h122c-14-7-27-17-38-26c-48-46-57-110-19-146c16-14 37-22 60-22c33 0 67 13 96 40c16 15 31 36 40 56c10-20 25-41 41-56c28-27 64-40 96-40c24 0 45 8 60 22c38 36 29 100-20 146c-10 9-23 19-36 26m-178-29c4-9-6-55-39-86c-17-16-40-25-60-25c-6 0-16 1-23 7c-2 2-8 6-8 17c0 13 8 33 27 50c24 23 62 39 90 39c9 0 13-2 13-2m76 0s5 2 13 2c28 0 66-16 90-39c19-17 25-37 25-50c0-11-5-15-6-17c-7-6-18-7-24-7c-19 0-43 9-60 25c-32 31-43 76-38 87zm-64 552H52c-29 0-52-23-52-52V482h288zm287 0H340V482h288v183c0 29-24 52-53 52"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Print(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m181 57l-11 86h377l-11-86c0-26-21-46-47-46H227c-26 0-46 20-46 46M46 190h625c26 0 46 20 46 46v274c0 26-20 46-46 46h-85v127H131V556H46c-26 0-46-20-46-46V236c0-26 20-46 46-46m494 447V340H177v297zm-54-211H226c-6 0-11-6-11-12v-14c0-6 5-11 11-11h260c6 0 12 5 12 11v14c0 6-6 12-12 12m0 83H226c-6 0-11-6-11-12v-14c0-6 5-11 11-11h260c6 0 12 5 12 11v14c0 6-6 12-12 12m0 83H226c-6 0-11-6-11-12v-14c0-6 5-11 11-11h260c6 0 12 5 12 11v14c0 6-6 12-12 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Qicon(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M490 210v-97h72v754h-72V558c-53 58-132 96-216 96C122 654 0 534 0 384s122-271 274-271c84 0 163 39 216 97M275 587c115 0 215-90 215-203S390 180 275 180c-114 0-207 91-207 204s93 203 207 203"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QupperCase(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m676.018 813l-76-95c-41 25-86 42-136 50c-216 38-421-101-458-312c-38-210 108-411 324-449s421 101 458 312c25 137-30 271-130 357l110 137zm-187-350l124 157c66-57 107-141 107-233c0-173-145-313-323-313s-322 140-322 313s144 315 322 315c57 0 111-16 158-40l-158-199z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Question(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M86 281H12c-7-23-12-46-12-71C0 84 81 0 207 0c117 0 191 64 204 178c1 8 1 17 1 26c0 53-14 98-50 135c-35 35-59 58-97 91c-41 36-41 98-41 144h-77c0-108 24-153 59-188c34-34 83-64 109-97c22-26 28-57 28-85c0-87-49-136-136-136c-88 0-138 55-138 142c0 26 6 49 17 71m61 373h77v77h-77z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Quote(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m306 106l-26-40C100 187 0 334 0 455c0 117 86 172 159 172c92 0 157-78 157-160c0-69-44-128-103-150c-17-6-33-11-33-40c0-37 27-92 126-171m397 0l-26-40C499 187 397 334 397 455c0 117 88 172 161 172c93 0 159-78 159-160c0-69-45-128-106-150c-17-6-32-11-32-40c0-37 28-92 124-171"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Quotedbl(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h95v113L39 231L3 213L52 95H0zm167 0h95v113l-56 118l-37-18l50-118h-52z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Quotesingle(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h94v113L38 231L2 213L52 95H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func R(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M71 128v64c42-40 99-64 161-64c19 0 38 3 55 7v75c-17-6-36-11-55-11c-87 0-159 70-161 157v298H0V128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RupperCase(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M278 352H174l297 383h-91L83 352H71v383H0V0h278c58 0 110 27 142 71c22 30 34 66 34 105s-12 75-34 104c-32 44-84 72-142 72m3-281H71v209h210c56-1 102-47 102-104c0-56-46-103-102-105"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Readability(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M210 353v-1c-2-16-7-33-15-47c-7-12-17-23-28-31C181 152 285 58 410 58s228 94 242 216c-10 8-20 19-27 31c-8 14-14 31-16 47v1l-8 77c-57-36-123-55-191-55s-135 19-192 55zm-60 243L91 458c-50-1-91-43-91-93c0-52 42-94 94-94c34 0 64 18 80 46c7 11 11 24 13 38v1l29 269c0 5-4 11-9 11h-48c-5 0-9-6-9-11zm482-240v-1c2-14 6-27 13-38c16-28 46-46 80-46c52 0 94 42 94 94c0 50-41 92-91 93l-58 138v29c0 5-5 11-10 11h-48c-5 0-9-6-9-11zM232 566l-11-111c54-36 119-56 189-56s134 20 188 56l-11 111z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Record(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M85 207v-95C85 51 140 0 210 0c69 0 125 51 125 112v95zm337 6v191c-41 112-162 129-162 129v127h133v57H27v-57h134V533S41 516 0 404V213l49-23v174s15 117 161 117s163-117 163-117V190zm-87 27v86c0 61-56 112-125 112c-70 0-125-51-125-112v-86z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Refresh(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M536.25 399h94c0 82-31 165-92 226c-123 123-323 123-446 0s-123-323 0-446c56-56 129-86 203-91V0l241 139l-241 140v-97c-50 4-98 26-136 64c-87 86-87 226 0 312c86 87 226 87 312 0c44-43 66-102 65-159"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Refreshbutton(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M358 0c198 0 359 161 359 359S556 717 358 717S0 557 0 359S160 0 358 0M198 368h-69c0 60 23 120 67 165c90 89 235 89 325 0c89-90 89-236 0-325c-42-41-94-63-148-66V65L198 179l175 113v-82c37 3 71 18 99 46c63 62 63 166 0 228c-62 63-165 63-227 0c-33-31-48-74-47-116"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Remove(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M358 0c198 0 359 161 359 359S556 717 358 717S0 557 0 359S160 0 358 0m2 411l114 115c14 14 37 14 51 0c13-13 13-36 0-50L411 361l114-115c13-14 13-37 0-50c-14-14-37-14-51 0L360 311L246 196c-14-14-37-14-51 0c-13 13-13 36 0 50l114 115l-114 115c-13 14-13 37 0 50c14 14 37 14 51 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Repeat(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M453 28v95H45c-16 0-30 14-30 31v264l107-95c5-6 10-8 15-11v-67h316v95c0 28 17 36 40 16l161-148c7-6 12-14 12-24c0-9-5-18-12-24L493 12c-23-20-40-12-40 16m197 288l-106 96c-5 4-12 7-16 10v67H213v-96c0-28-17-35-40-15L11 526c-7 6-11 14-11 24c0 9 4 18 11 24l162 147c23 20 40 13 40-15v-95h407c17 0 30-14 30-31z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reply(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 284l274 254V411s389-100 493 245c0 0 37-494-493-494V38z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Right(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M248 361L2 115L121 0l354 355l-363 362L0 606z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rss(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 126V27c39-8 80-13 122-13c300 0 544 244 544 544c0 42-6 83-14 122h-98c11-39 16-80 16-122c0-248-200-449-448-449c-42 0-83 6-122 17m0 228V248c38-15 79-23 122-23c183 0 333 150 333 333c0 43-9 84-24 122H326c22-36 33-78 33-122c0-131-106-238-237-238c-45 0-86 12-122 34m122 82c68 0 122 54 122 122s-54 122-122 122S0 626 0 558s54-122 122-122"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sicon(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m280 220l45-47c-50-49-101-74-151-74c-43 0-78 14-107 41c-28 28-42 62-42 103c0 32 9 59 26 85c17 25 52 49 101 75c45 24 75 42 88 57c13 16 19 35 19 55c0 24-11 45-30 63c-19 17-43 27-71 27c-40 0-78-21-114-61L0 595c18 24 42 43 71 56s58 20 90 20c47 0 85-15 117-46c32-32 48-69 48-113c0-32-10-60-28-86c-18-25-55-51-107-78c-43-22-71-42-84-58s-20-33-20-50c0-20 9-38 25-53c16-14 35-22 58-22c36 0 73 18 110 55"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SupperCase(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M63 580L0 617c31 56 93 155 222 155c154 0 211-130 208-197c-1-37-11-74-35-111s-68-79-131-126c-54-40-86-63-95-73c-17-17-46-50-46-96c0-31 21-93 100-93c69 0 108 53 136 90l59-46C384 74 328 0 225 0C95 0 46 102 46 172c0 40 13 78 39 113c15 20 54 54 115 100s103 85 126 118c17 24 24 48 25 73c1 43-36 125-136 125c-57 0-108-40-152-121"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Safari(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M358 61c198 0 359 161 359 359S556 778 358 778S0 618 0 420C0 253 113 113 267 72c0-1-1-2-1-4c-19-4-34-18-34-34c0-19 21-34 47-34s47 15 47 34c0 11-7 21-17 27v3c16-2 33-3 49-3m-93-3c3-10 12-17 22-17c8 0 14 4 18 10c5-5 9-12 9-18c0-14-16-26-35-26s-35 12-35 26c0 11 9 21 21 25m93 661c165 0 299-134 299-299c0-166-134-299-299-299S59 254 59 420c0 165 134 299 299 299m2-583c154 0 279 125 279 280c0 154-125 280-279 280S82 570 82 416c0-155 124-280 278-280m0 541c142 0 260-118 260-261c0-144-118-260-260-260S101 272 101 416c0 143 117 261 259 261m-52-435l-8-40c-4-2-5-3-6-4h-3v-2l17-4l36 29h1l-5-25c-1-2-2-4-3-5s-3-1-6-1l-1-2l18-4v2c-2 1-4 2-5 3c0 1-1 3 0 5l10 49l-2 1l-46-37l-1 1l7 32c0 2 1 4 2 5s5 2 8 1l1 3l-20 4v-2c3-1 5-3 5-4c1-1 1-3 1-5m227 139l-91 41c-1 9-3 17-7 25l35 36l-47-16c-6 8-16 15-25 20l-4 92l-39-84c-11 0-21-3-30-8L196 600l99-141c-2-3-4-5-5-8l-96-3l87-41c1-11 4-21 9-30l-33-37l46 18c7-7 16-13 24-17l5-100l41 91c11 2 21 6 30 11l123-105l-94 132c1 2 2 3 3 5zm-197 17l60-50c-7-4-14-6-22-8c-4-1-9-1-14-1c-13 0-24 3-35 9c-6 3-13 7-17 12c-7 6-12 14-16 22c-3 7-5 14-6 22c0 3-1 6-1 10c0 13 4 26 10 37c1 1 1 2 2 3zm239-7v-1c4-2 7-4 8-7c2-2 2-7 0-13c0-2-1-3-1-4c-1 0-2-1-3 0l-17 3c0 4 2 8 4 9c1 1 5 1 10 1v2l-25 5v-2c4-1 6-3 8-5c1-1 1-5 0-9l-17 3c-1 0-2 1-3 2c0 1 0 3 1 6c0 4 2 8 4 10c3 3 7 5 12 6v2h-15l-9-42l2-1c1 2 1 4 2 5s2 1 4 1l36-7c2-1 3-2 3-3c1-1 1-3 0-5h2l9 42zm-194 49l-51 43c7 3 14 5 22 6h8c14 0 28-3 38-10c7-4 14-9 19-14c6-7 10-15 14-23c2-6 3-11 4-17c1-4 1-8 1-11c0-14-4-26-10-37l-45 62zm-58 39l52-45l-3-3c4-3 6-9 6-14c0-10-8-17-18-17c-6 0-11 3-14 7l-4-3l-41 56l-3 7l-61 85l80-69zm37-75c7 0 14 6 14 13s-7 13-14 13s-12-6-12-13s5-13 12-13m-165 54v2l-42 26c-2 1-3 2-4 3v3h-3l-4-21l2-1c0 2 1 3 1 4c1 0 2 1 4 0h1l1-1l24-15l-24-4s-1 0-3 2l-4 2l-2 2v5h-1l-5-24h2c0 2 1 3 1 4c1 1 1 1 2 1s1-1 2-1c0 0 1-1 2-1l22-13l-25-3h-4s-2 1-2 2v4h-2l-2-14h1c1 1 1 2 2 2c1 1 2 2 4 2l49 6v2l-32 19v1zm205 145l7 2c3 1 5 2 7 3c3 2 6 5 6 10c1 4 1 8-2 12c-2 4-7 8-13 9c-2 0-5 1-6 1c-2 0-4 0-5-1h-4c-1 0-1 1-1 1c-1 1-1 1-1 2l-2 1l-4-21h2c3 5 6 10 9 12s8 3 11 2c4 0 6-3 7-5s2-4 1-6c0-2-1-4-3-5s-5-2-8-3l-6-1c-5-2-9-4-12-6c-2-2-4-4-5-8c0-4 0-8 2-12c2-3 6-7 12-8c2-1 5-1 8 0h6l1-1c1-1 1-2 1-2l2-1l4 18l-3 1c-1-4-4-8-7-10c-3-3-8-4-11-3s-5 2-6 4c-1 1-2 4-1 6c0 3 1 5 3 5c2 1 5 3 11 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Save(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M390 0v149h78L312 337L156 149h82V0zm132 241l91 144c5 12 11 35 11 49v277c0 15-12 26-26 26H27c-15 0-27-11-27-26V434c0-14 5-37 10-49l91-144c5-13 23-24 36-24h41l42 51h-84L55 416c-1 0-1 1-1 2c0 2 0 3-1 4h517v-4c0-1-1-2-1-3l-81-147h-84l42-51h40c14 0 31 11 36 24M239 530h147c13 0 25-12 25-26c0-13-12-25-25-25H239c-14 0-26 12-26 25c0 14 12 26 26 26"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m541 517l176 176l-72 73l-184-185c-45 25-97 39-151 39C139 620 0 481 0 310S139 0 310 0s311 139 311 310c0 80-31 153-80 207M103 310c0 115 92 207 207 207s207-92 207-207s-92-207-207-207s-207 92-207 207"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Semicolon(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M67 145h77v77H67zm17 393l70 33l-93 196l-61-29z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sepia(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M49.064 397c-60 42-24 95 26 98c50 4 92 38 80 79c-22 70-8 128 106 70c58-29 97-101 172 30c44 77 156 26 141-62c-7-43 17-69 101-94c153-45 88-172 12-204s-87-87-60-143c26-54-48-141-113-83c-52 46-117 4-118-12s-57-133-90-34c-22 63 13 172-181 59c-123-72-183 113-49 178c63 31 33 77-27 118"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Server(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M36 160h508c19 0 36 17 36 36v84c0 20-17 36-36 36H36c-20 0-36-16-36-36v-84c0-19 16-36 36-36m500 86v-16c0-6-6-12-12-12h-64c-7 0-12 6-12 12v16c0 7 5 12 12 12h64c6 0 12-5 12-12M36 344h508c19 0 36 16 36 36v84c0 20-17 36-36 36H36c-20 0-36-16-36-36v-84c0-20 16-36 36-36m500 86v-16c0-6-6-12-12-12h-64c-7 0-12 6-12 12v16c0 7 5 12 12 12h64c6 0 12-5 12-12M36 528h508c19 0 36 16 36 36v84c0 20-17 36-36 36H36c-20 0-36-16-36-36v-84c0-20 16-36 36-36m500 86v-16c0-7-6-12-12-12h-64c-7 0-12 5-12 12v16c0 7 5 12 12 12h64c6 0 12-5 12-12m8-481H36c-6 0-12 1-17 3l61-83c12-16 38-29 58-29h304c19 0 45 13 57 29l62 83c-6-2-12-3-17-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Seven(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 71V0h467L73 754L9 721L349 71z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m623 474l75-74v265H0V133h281l-84 72H76v388h547zm-73-121v110l238-222L550 28v109c-460 0-427 427-427 427c90-298 427-211 427-211"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shopping(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M644 480v46H203L126 73H0V28h165l77 452zm73-453l-69 400H287L217 27zM334 150V73h-62l14 77zm46 0h63V73h-63zm109-77v77h64V73zm110 77h50l14-77h-64zm-265 46h-41l11 64h30zm109 64v-64h-63v64zm46 0h64v-64h-64zm152-64h-42v64h31zM334 306h-22l13 75h9zm46 75h63v-75h-63zm173-75h-64v75h64zm46 0v75h10l14-75zM254 553c32 0 56 25 56 56c0 32-24 57-56 57c-31 0-56-25-56-57c0-31 25-56 56-56m343 0c32 0 56 25 56 56c0 32-24 57-56 57s-56-25-56-57c0-31 24-56 56-56"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shuffle(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 497v64c0 9 7 16 16 16h92c54 0 99-25 136-63c38-38 69-89 100-138c25-40 50-78 76-107c25-30 50-48 80-48h85v73c0 23 16 27 35 12l139-116c12-10 12-29 0-39L620 36c-19-16-35-12-35 12v77h-85c-55 0-98 24-136 63c-37 38-71 89-101 137c-26 41-50 80-76 109c-25 29-50 47-79 47H16c-9 0-16 7-16 16m0-357v64c0 9 7 16 16 16h92c45 0 81 44 120 101c1-3 2-5 3-7c1-1 4-3 5-5c15-25 32-49 48-75c-24-31-47-59-77-78c-29-20-61-31-99-31H16c-11 0-16 5-16 15m585 339h-85c-44 0-82-43-120-99c-2 3-3 5-4 7c-2 2-3 4-5 7c-7 12-16 25-23 37c-8 11-16 24-23 35c23 31 47 58 76 77c28 20 61 33 99 33h85v70c0 23 16 28 35 12l139-117c12-10 12-28 0-38L620 388c-19-15-35-12-35 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sitemap(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M604 338v142h76v181H499V480h75V368H355v112h76v181H249V480h76V368H106v112h75v181H0V480h76V338h249V228h-98V47h226v181h-98v110z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Six(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M373 39L186 331c14-2 28-4 42-4c125 0 227 102 227 227c0 126-102 227-227 227c-106 0-196-73-221-171c-5-17-7-37-7-56c0-43 13-83 33-117l6-10L311 0zM106 456l-19 28c-10 22-16 45-16 70c0 86 71 157 157 157s155-71 155-157s-69-156-155-156c-49 0-93 23-122 58"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Skype(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M696 441c13 27 21 56 21 88c0 110-89 198-199 198c-34 0-65-8-93-23c-20 4-41 6-63 6c-189 0-342-154-342-342c0-24 3-47 8-69c-17-30-28-64-28-100C0 89 89 0 199 0c38 0 74 12 105 31c18-3 38-5 58-5c188 0 343 153 343 342c0 25-4 49-9 73m-158 90c15-23 23-48 23-76c0-24-4-44-13-60s-23-30-39-41c-16-10-35-20-58-27s-48-14-76-20c-22-5-38-9-47-11c-9-3-19-7-28-11c-8-4-15-9-20-15c-4-6-7-13-7-20c0-13 7-24 21-33c15-10 34-15 59-15c27 0 47 5 58 13c11 9 21 22 30 38c7 13 14 22 21 28c6 6 16 9 28 9c14 0 27-5 36-15c9-9 14-21 14-33s-4-26-11-39s-18-26-32-37c-15-12-35-22-58-29c-22-7-49-10-79-10c-39 0-72 6-100 16c-30 11-51 27-66 46c-15 20-24 43-24 69c0 27 8 51 23 69c14 17 33 33 58 43c24 10 54 19 90 27c26 5 47 10 62 15c14 5 27 11 36 20c9 8 14 19 14 32c0 17-9 32-25 43c-17 12-41 19-70 19c-20 0-37-4-49-10s-22-13-28-22c-7-9-14-21-20-35c-6-12-12-23-20-29c-8-7-19-11-30-11c-14 0-27 5-36 13c-9 9-14 21-14 33c0 19 7 42 21 62s34 38 57 50c33 17 73 25 122 25c41 0 77-6 107-18c31-13 55-30 70-53"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Slash(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 854L306 0h77L77 854z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sleipnir(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M406 606H176s-8-10-21-23c0-140 80-205 117-257c-36 19-68 30-86 31c-11 17-45 58-55 65c-12 26-31 32-51 32c-70 0-83-73-79-83s90-140 105-172s4-23 20-50c16-28 51-54 58-59c-15-11-48-52-52-71c-6-36 17-10 44 8c16 11 49 31 67 39c138 0 219 89 219 195c0 83-33 172-33 285v37zm43-40v-20c0-59 9-111 17-161c8-43 15-84 15-124c0-59-22-112-63-152c-43-40-101-63-171-64h-1c10-13 28-26 34-27c9-2 22-2 30-4c9-2 15-8 38-8c18 0 42 7 42 12c0 7-26 3-42 20c-4 6 9-1 16-1c6-1 17-2 33 2c11 3 24 6 46 1c7-2 22 0 4 16c-12 10-18 14-22 17c-6 4 27-13 62 19c19 16 46 36 28 44c-10 5-24 0-34 3c-8 3 18 2 25 7c4 4 6 16 13 23c8 7 16 8 22 18c6 9 19 37 0 27c-13-7-24-12-37-11c-10 0 20 8 32 28c7 13 0 18 4 29c4 9 13 16 16 24c5 13-3 61-9 74c-5 12-10 20-16 13c-5-6-15-6-19-6c7 11 2 18 6 27c5 11 13 17 13 28c0 10-6 35-15 35c-11 0-24-16-30-12c-8 4 20 24 23 37c4 16-10 28-13 38s0 29-9 38c-7 7-13 10-38 10m-264 67h212l74 82l-10 55H122l-7-58z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Slideshare(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M634 66v264c3-2 5-4 8-6c20-14 34 5 21 24c-24 29-70 66-141 95c26 88 12 151-17 190c-17 24-40 38-63 43v1c-50 13-100-11-98-69c0-1-1-71-1-124c-5-1-11-3-18-5v129c4 130-253 90-179-165c-70-29-116-66-141-95c-12-19 1-38 22-24c2 2 5 4 8 6V66C35 30 62 0 96 0h477c34 0 61 30 61 66m-33 283V96c0-43-14-61-54-61H125c-42 0-54 15-54 61v256c90 47 167 38 209 37c18-1 30 3 36 10c2 1 3 2 4 4c8 7 16 13 23 19c2-21 14-35 45-33c43 1 122 10 213-40M244 204c45 0 81 35 81 77s-36 77-81 77c-46 0-82-35-82-77s36-77 82-77m190 0c45 0 82 35 82 77s-37 77-82 77s-82-35-82-77s37-77 82-77"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Small(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M333 83v226c0 10 4 19 11 27c6 7 16 11 27 11h225c32 0 40-17 17-40l-71-73l118-118c3-4 6-9 6-14s-3-9-6-13l-69-69c-4-4-8-6-14-6c-5 0-10 2-14 6L444 139l-71-72c-23-23-40-15-40 16M5 605l70 69c4 3 8 6 14 6c5 0 9-3 13-6l119-119l72 71c23 23 40 16 40-15V385c0-19-16-38-38-38H70c-32 0-41 17-18 40l72 72L5 577c-3 4-5 10-5 15s2 9 5 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Smile(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M358 0c198 0 359 155 359 349S556 698 358 698S0 543 0 349S160 0 358 0M258 171h-41c-11 0-19 9-19 20v100c0 11 8 21 19 21h41c10 0 20-10 20-21V191c0-11-10-20-20-20m242 0h-41c-10 0-20 9-20 20v100c0 11 10 21 20 21h41c11 0 19-10 19-21V191c0-11-8-20-19-20M228 413l-39 13c-4 1-7 4-8 7c-2 3-2 7-1 11s36 105 178 104c142 1 178-100 179-104s1-8-1-11c-1-3-5-6-9-7l-34-12l-5-1c-6 0-11 5-13 10c0 1-22 61-117 61c-94 0-115-58-116-61c-2-5-8-10-14-10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sns(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m271 392l173 97c25-29 62-47 104-47c76 0 137 62 137 138s-61 137-137 137c-77 0-138-61-138-137c0-8 1-17 2-25l-178-99c-25 25-59 40-97 40C61 496 0 435 0 359s61-137 137-137c39 0 75 16 100 42l177-96c-2-10-4-20-4-30C410 62 471 0 548 0c76 0 137 62 137 138s-61 137-137 137c-39 0-74-15-99-42l-178 95c2 10 3 21 3 31c0 11-1 23-3 33"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sort(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 255L200 17c17-22 45-22 62 0l191 238c17 22 9 40-20 40H29c-29 0-37-18-20-40m444 188L262 681c-17 22-45 22-62 0L9 443c-17-22-9-40 20-40h404c29 0 37 18 20 40"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Soundcloud(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M900 333c-8-69-36-116-77-153c-38-35-93-65-168-61c-17 1-33 3-49 7c-10 3-33 9-37 17c-3 5-2 16-2 26v364c0 16-3 45 5 52c5 7 18 5 32 5h335c42 0 70-9 94-28c23-17 40-42 48-73c9-36 1-73-14-99c-29-49-92-84-167-57M539 151c-16 1-14 16-14 34c-2 29-1 59-2 85c-2 54 0 114-3 169c-1 38 2 71 4 110c1 19-1 41 17 40c16-1 14-17 15-34c2-31 5-63 5-92c1-50-3-103-3-153c0-37-1-77-3-115c-1-22 3-46-16-44m-47 26c-14 2-11 16-11 33c-1 26-2 56-2 79c0 35 0 67-2 105c-2 54 0 103 2 156c1 19 0 39 14 40c19 1 16-22 16-41c3-52 6-103 4-156c-3-51-1-106-4-159c0-10 2-44-2-51c-4-6-12-7-15-6m-136 31c-11 2-9 18-9 33c0 55-3 105-5 159c-1 23-2 45-2 64c1 31 5 69 7 101c0 11 0 25 11 25c14 0 12-23 13-37c1-17 2-33 3-48c5-64-1-128-2-195c0-26-2-47-2-74c0-16-1-30-14-28m44 6c-11 3-9 20-9 35c0 33-2 60-3 93c-1 40-3 86-3 122c0 30 4 70 6 102c1 13 1 23 12 23c13 1 12-13 13-25c3-34 6-67 6-99c1-19 0-41-1-62c-2-52-3-102-5-156c-1-19 2-37-16-33m-88 3c-11 2-9 14-9 27c-1 24-2 51-2 72c-2 62-7 128-2 190c0 13 2 28 3 45c1 15-2 39 12 38c10 0 10-13 11-24c2-34 6-67 7-99c1-20-1-41-2-62c-2-50-3-102-5-154c-1-15 1-35-13-33m133 5c-11 3-9 15-10 29c-1 87-7 174-2 259c0 13 1 29 2 46c1 16-1 35 15 34c9-1 12-10 13-23c1-35 5-71 5-103c0-38-2-80-3-121c-1-28-2-62-3-91c-1-11 1-23-4-27c-3-4-11-5-13-3m-178 15c-8 1-7 9-7 20c-2 59-4 104-6 162c-1 15-2 31-2 45c1 32 4 57 6 89c1 13-1 37 12 37c10-1 9-16 9-26c3-34 8-67 8-100c0-17-1-37-2-56c-2-48-4-95-6-143c0-12 1-31-12-28m-43 41c-10 2-8 27-9 40c-2 45-3 75-5 120c-2 42 2 79 5 119c1 13 0 32 10 32c10 1 9-16 10-26c1-13 2-23 2-33c2-23 6-45 6-65c0-12-2-30-3-46c-2-41-4-77-6-116c-1-13 1-26-10-25m-81 64c-2-1-7-3-8 0c-5 2-4 20-5 30c-1 11-2 22-3 31c-1 21-5 41-5 63c0 20 3 42 5 63c1 9 2 20 3 31c1 8 0 27 9 26c7 0 7-17 7-26c1-12 2-22 3-32c2-20 6-43 7-62c0-11-2-21-3-32c-2-21-4-42-6-62c-1-11-1-21-4-30m-50 3c-5 2-5 22-5 30c-1 11-2 21-3 31c-2 20-6 39-5 61c0 28 5 60 8 90c1 7-1 26 8 25c6-1 6-17 7-26c1-12 2-22 3-31c2-20 6-39 6-60c0-22-4-40-6-61c-1-10-2-19-3-31c-1-9 0-32-10-28m87 3c-8 2-7 17-8 27c0 11-1 22-2 31c-1 21-5 38-5 60c0 20 3 41 5 63c1 10 2 20 2 33c1 10 1 25 8 26c10 1 9-13 11-25c2-25 3-43 5-65c1-11 3-21 3-33c0-19-4-38-6-57c-1-9-1-19-2-30s0-32-11-30M49 365c-3 16-4 33-6 50s-6 34-6 51c1 16 4 35 6 52c1 7 2 15 3 25c1 6 1 22 7 22c3 0 4-2 5-9c3-23 5-43 8-64c1-9 3-17 3-26s-2-18-3-27c-3-18-4-33-7-52c-1-8 0-29-10-22m-34 37c-2-1-6-2-6 1c-2 8-4 20-5 31c-1 10-4 22-4 32c0 9 3 22 5 33c0 4 1 10 2 17c1 5 1 13 6 13c6-1 7-23 8-30c2-11 5-24 5-33c0-11-3-22-5-33s-2-22-6-31"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spa(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m528 349l-14 19c46 16 77 46 77 90c0 81-117 142-271 142S49 539 49 458c0-44 31-74 77-90l-14-19C45 365 0 411 0 475c0 103 134 187 320 187c184 0 320-83 320-187c0-64-45-110-112-126M148 492l20 13c57-36 71-77 71-128c0-26-11-63-31-99c-21-36-34-73-34-101c0-29 13-60 47-88l-20-13c-58 37-71 82-71 127c0 27 11 64 32 99c20 36 34 73 34 101s-13 61-48 89m127 31l21 11c58-39 76-88 76-145c0-30-12-71-33-112s-33-81-33-114c0-34 16-73 52-105l-22-12c-61 42-76 94-76 146c0 30 12 72 33 112c21 41 34 80 34 114s-17 72-52 105m139-31l20 13c57-36 70-77 70-128c0-26-10-63-31-99c-20-36-34-73-34-101c0-29 14-61 48-89l-20-12c-59 37-71 82-71 127c0 27 10 64 31 99c21 36 34 73 34 101s-13 61-47 89"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sqale(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M178 55C92 55 10 84 10 191c0 55 30 103 113 163c85 61 110 85 110 128c0 60-50 76-105 76c-26 0-64-2-95-8c-23-4-33 16-33 30c0 13 4 21 11 26c23 14 88 18 117 18c110 0 183-46 183-149c0-63-32-104-117-163c-70-48-105-86-105-126c0-53 34-66 103-66c38 0 59 3 82 5c26 3 34-11 34-27c0-12-4-21-17-28c-18-10-53-15-113-15m960 446l20 83c6 27 18 35 38 35c38 0 47-13 47-32c0-7-2-20-9-45c-25-99-80-313-114-434c-12-41-35-53-81-53c-41 0-67 15-77 47c-33 108-96 343-122 440c-6 20-9 38-9 45c0 23 20 32 45 32c15 0 25-5 32-30l21-88c8-28 17-34 35-34h135c24 0 33 8 39 34m-119-351c5-16 9-20 19-20c9 0 14 4 19 20c17 55 55 220 55 235c0 16-8 19-30 19h-92c-20 0-29-3-29-19s42-180 58-235m381 378V91c0-26-18-31-40-31c-17 0-39 5-39 31v463c0 41 23 62 74 62h149c19 0 24-15 24-34c0-16-5-33-24-33h-121c-19 0-23-4-23-21m339-398h157c18 0 23-16 23-34c0-17-5-33-23-33h-186c-50 0-74 23-74 62v429c0 41 24 62 74 62h199c18 0 23-15 23-34c0-16-5-33-23-33h-170c-20 0-23-4-23-21V384c0-18 3-22 23-22h145c19 0 24-13 24-32c0-16-5-30-24-30h-145c-20 0-23-4-23-21V165c0-28 3-35 23-35M746 345c33 0 41 14 40 35c-1 2-1 4-1 7c-8 102-34 160-76 195c-7 6-6 12-3 15c18 20 37 41 43 46c12 14 14 32-2 48c-19 17-36 9-47-2c-7-7-37-42-59-67c-3-3-10-8-20-8c-24 3-54 3-76 0c-80-11-137-57-157-171c-3-16-5-30-6-46c0-3 0-6-1-10v-7c-1-20 7-35 40-35c32 0 42 15 42 29c0 3 0 6 1 10v16c1 7 2 32 6 50c14 83 50 100 116 100c86 0 115-23 120-167v-9c0-14 8-29 40-29m47-86c2 22-1 43-41 43c-16 0-93-23-166-23c-80 0-155 23-169 23c-44 0-42-28-40-41c2-15 4-30 7-43c20-103 71-153 156-167c11-2 15-9 16-13c1-9 4-18 8-25c5-10 15-13 27-13c10 0 21 4 25 13c4 8 6 17 8 26c0 4 2 11 15 13c94 18 142 85 154 207"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m148 739l241-125l241 125l-46-267l194-189l-269-39L389 0L269 244L0 283l195 189z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Starempty(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M630 739L389 614L148 739l47-267L0 283l269-39L389 0l120 244l269 39l-194 189zM278 300l-155 23l111 108c13 13 20 32 17 50l-26 154l137-73c8-4 18-6 27-6s19 2 27 6l137 73l-26-154c-3-18 4-37 17-50l111-108l-154-23c-18-3-35-15-43-32l-69-139l-68 139c-8 17-25 29-43 32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stop(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 45v604c0 17 14 31 31 31h604c17 0 31-14 31-31V45c0-17-14-31-31-31H31C14 14 0 28 0 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Strike(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 323h700v62H0zm370-64l64 34H179c-18-27-29-60-29-97C149 80 239 1 387 1c33 0 75 4 116 16c13 4 23 6 35 8c8 23 18 78 18 146c0 0-39 3-40 3c-17-52-61-126-134-126c-68 0-104 44-104 95c0 47 41 89 92 116m73 298c0-66-50-111-102-142h223c11 24 18 52 18 85c0 126-103 207-258 207c-93 0-152-26-173-38c-20-23-33-96-33-166l39-3c20 66 89 160 176 160c75 0 110-52 110-103"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Surprise(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M358 0c198 0 359 155 359 349S556 698 358 698S0 543 0 349S160 0 358 0M119 236c0 58 47 106 105 106s106-48 106-106s-48-105-106-105s-105 47-105 105m374 106c58 0 105-48 105-106s-47-105-105-105s-106 47-106 105s48 106 106 106M271 236c0 26-21 47-47 47s-47-21-47-47s21-47 47-47s47 21 47 47m222-47c26 0 47 21 47 47s-21 47-47 47s-48-21-48-47s22-47 48-47M316 579h84c51 0 94-42 94-94s-43-94-94-94h-84c-51 0-93 42-93 94s42 94 93 94"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sync(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M563 347h65c0-58-22-113-56-155h-1c-25-30-55-53-91-68c-1 0-2-1-2-1c-6-2-12-6-18-8c-2 0-3-1-5-1c-5-2-12-3-17-4c-2-1-5-2-7-2c-6-1-11-2-17-3h-6c-7-1-15-1-22-1c-50-1-99 14-141 43c-17 12-43 8-54-9c-12-17-7-41 10-53c54-38 117-58 182-58c1 0 1-1 1-1c2 0 3 1 4 1c9 0 18 0 26 1c4 0 7 1 11 1c4 1 9 1 14 2c2 0 4 1 6 1c3 1 8 3 12 4c6 1 13 3 19 5c3 1 6 2 8 3c8 2 16 5 23 8c2 0 3 1 4 1c46 20 87 50 120 90c0 1 1 1 1 2c5 6 10 13 14 19c1 0 1 1 1 2c36 52 58 113 58 181h63L666 500zM0 347l102-154l103 154h-65c0 57 21 109 55 151c1 1 1 2 2 3c5 6 10 12 15 17v1c16 16 35 30 55 41c1 0 1 1 1 1c6 3 13 6 19 8c1 1 3 2 4 2c5 2 10 4 16 6c2 1 4 2 7 2c4 2 10 3 15 4c3 1 6 1 9 2c5 1 10 2 15 2c3 1 5 1 8 1c6 1 14 1 20 1c51 1 100-14 142-44c17-11 43-6 54 11c12 17 7 41-10 53c-54 38-117 58-182 58c-11 0-22-1-33-2c-2 0-4-1-6-1c-6-1-11-1-17-2c-2-1-4-1-5-1c-4-1-10-2-14-3c-6-2-12-3-18-5c-3-1-6-2-10-3c-6-2-14-5-21-8c-2 0-4-1-5-2c-8-4-17-8-24-12c-1 0-2-1-2-1c-27-14-50-32-72-54c0-1-1-1-1-2c-6-6-13-13-19-20c-2-1-3-4-4-6l-12-15c0-1-1-1-1-2c-36-51-58-113-58-181z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func T(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M164 228h91v60h-91v466H92V288H0v-60h92V0h72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TupperCase(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h455v71H263v664h-71V71H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tabezou(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1087 158l15-15c7-25-7-19 26-51c4-4-10 10 32-33c28-27 79-7 67 4c-14 15-68 69-117 118l69 69c6 6 6 16 0 22s-17 6-23 0l-19-18c-4 5-8 6-8 13c0 9-4 68-6 85c-4 28-42 84-61 108c-19 23-67 44-105 44c-38-1-82-33-110-53c-28-22-63-62-63-62v284H656V521h-28v152H493V521h-75v156H290V521h-26v152H135V295L0 416s95-193 128-230c30-32 52-44 81-44h461s38-11 54-10c16 2 59 14 75 23c15 9 52 52 63 78c12 26 27 114 38 129c10 15 39 43 57 47s57-11 57-11s23-20 35-49c7-18 8-47 3-71c-3-10-13-17-18-22l-16 16c-6 6-17 6-23 0s-6-16 0-22l69-69l-32-33c-17 13-36 15-58-7l-46-45c-4-5-4-12 0-17c5-5 12-5 17 0l44 44c3 3 8 3 11 0s3-9 0-12l-44-44c-5-5-5-12 0-17c5-4 13-4 18 0v1l44 43c3 3 8 3 11 0s3-8 0-11l-44-44c-5-4-5-13 0-17c4-5 12-5 16 0l47 45c22 22 19 42 7 59zM601 455h81v-29h-81c-22 0-46-1-63-18c-18-18-34-51-34-65l-2-142h-30l2 142c0 25 19 63 43 87c26 26 59 25 84 25m526-211l-40-41c-17 16-33 32-45 44c12-1 32 3 45 3c14 0 30-5 40-6m-410 51c14 0 26-13 26-27s-12-27-26-27s-27 13-27 27s13 27 27 27"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Table(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 93v508c0 25 20 45 45 45h628c25 0 44-20 44-45V93c0-25-19-45-44-45H45C20 48 0 68 0 93m60 173V137h179v129zm209 0V137h180v129zm209 0V137h179v129zM60 426V297h179v129zm209 0V297h180v129zm209 0V297h179v129zM60 585V455h179v130zm209 0V455h180v130zm209 0V455h179v130z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tabs(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M430 96H108c-19 0-36 17-36 36v286c0 19 17 36 36 36h35v72h-35C49 526 0 477 0 418V132C0 73 49 24 108 24h322c59 0 108 49 108 108h-72c0-19-17-36-36-36m-143 72h322c59 0 108 48 108 107v287c0 59-49 107-108 107H287c-59 0-108-48-108-107V275c0-59 49-107 108-107m0 430h322c19 0 36-17 36-36V275c0-19-17-36-36-36H287c-19 0-36 17-36 36v287c0 19 17 36 36 36"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 2h267l379 379c44 44 0 89 0 89S512 603 468 648c-45 44-89 0-89 0L0 269zm183 183c22-23 22-57 0-79s-57-22-79 0s-22 57 0 79s57 22 79 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Terminal(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 614V94h680v520zm59-343l121-47v-27L59 149v30l86 31l-86 32zm132 35h119v-30H191z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Three(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M105 192H33C35 87 122 0 228 0s193 87 193 195c0 63-31 119-78 154c67 40 112 113 112 196c0 126-102 227-227 227C105 772 5 676 0 555h72c5 81 73 147 156 147c86 0 155-71 155-157s-69-156-155-156v-71c68 0 122-55 122-123S296 72 228 72s-122 53-123 120"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tile(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 311h297V14H0zm369 0h297V14H369zM0 680h297V383H0zm369 0h297V383H369z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tilemenu(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 204h190V14H0zm238 0h190V14H238zm237 0h191V14H475zM0 442h190V252H0zm238 0h190V252H238zm237 0h191V252H475zM0 680h190V489H0zm238 0h190V489H238zm237 0h191V489H475z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Time(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M358 717C160 717 0 557 0 359S160 0 358 0s359 161 359 359s-161 358-359 358m0-628C210 89 90 211 90 359s120 268 268 268s270-120 270-268S506 89 358 89m97 407l-127-93c-9-7-18-22-18-34V166c0-11 10-20 22-20h42c11 0 20 9 20 20v160c0 12 9 27 18 34l93 67c9 7 11 21 4 30l-25 34c-7 9-20 11-29 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M441 99h99c18 0 32 14 32 32v23c0 15-11 29-25 32v475c0 31-26 56-58 56H83c-32 0-57-25-57-56V186c-14-3-26-17-26-32v-23c0-18 15-32 33-32h98V57c0-31 25-57 57-57h196c32 0 57 26 57 57zm-57-49H188c-5 0-8 3-8 7v42h212V57c0-4-3-7-8-7M99 653h375c4 0 7-5 7-9V187H91v457c0 4 3 9 8 9m77-37h-6c-18 0-33-15-33-33V256c0-17 15-31 33-31h6c17 0 32 14 32 31v327c0 18-15 33-32 33m113 0h-5c-18 0-33-15-33-33V256c0-17 15-31 33-31h5c18 0 33 14 33 31v327c0 18-15 33-33 33m113 0h-5c-18 0-32-15-32-33V256c0-17 14-31 32-31h5c18 0 33 14 33 31v327c0 18-15 33-33 33"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trouble(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M358 0c198 0 359 155 359 349S556 698 358 698S0 543 0 349S160 0 358 0M198 191v100c0 11 8 21 19 21h41c10 0 20-10 20-21V191c0-11-10-20-20-20h-41c-11 0-19 9-19 20m241 0v100c0 11 10 21 20 21h41c11 0 19-10 19-21V191c0-11-8-20-19-20h-41c-10 0-20 9-20 20m145 321c9-9 17-19 6-31c-9-9-20-19-29-29c-9-9-20-18-32-25c-20-12-46-15-70-9c-18 5-34 18-46 31c-13 13-25 34-43 40c-20 6-36-6-48-21c-12-14-24-29-40-39c-19-12-43-16-65-13c-23 4-40 16-55 31c-8 8-16 17-24 24c-8 8-24 20-12 32c7 7 15 19 25 24c12 5 24-13 31-21c18-17 39-45 68-30c16 9 28 29 41 42c14 14 33 24 53 27c23 4 47-1 66-13c16-11 28-26 40-40c14-16 34-27 54-15c8 5 14 13 20 20c9 8 18 18 27 26c12 12 25-3 33-11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tumblr(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M84 14h498c47 0 84 37 84 84v498c0 25-10 44-25 59c-14 14-36 25-59 25H84c-47 0-84-37-84-84V98c0-23 11-45 25-59c15-15 34-25 59-25m141 275v186c0 25 2 43 7 55s15 25 28 36s29 21 48 27c18 6 40 9 64 9c22 0 41-2 59-7c18-4 39-12 62-22v-84c-27 17-53 27-80 27c-15 0-29-5-40-12c-9-5-16-11-19-20c-3-8-5-28-5-59V289h126v-82H349V72h-75c-3 27-9 50-18 67c-9 18-21 33-36 46c-14 12-31 22-53 29v75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitter(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M768 108c-22 32-48 59-79 82c1 7 1 14 1 21c0 208-158 448-448 448c-89 0-172-27-242-71c12 1 25 3 38 3c74 0 141-26 195-68c-70-1-127-48-147-110c9 2 20 3 30 3c14 0 29-2 42-6c-73-14-127-77-127-153v-2c22 11 45 18 71 19c-42-29-70-77-70-131c0-29 8-56 21-80c78 95 194 159 325 165c-2-11-4-24-4-36c0-87 70-158 157-158c46 0 87 20 116 51c36-7 69-21 99-39c-11 37-36 69-68 88c32-4 61-13 90-26"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Two(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M423 368L156 682h325v72H0l255-299s117-135 128-152s19-47 19-75c0-87-70-158-158-158c-87 0-158 71-158 158c0 26 6 49 17 71H29c-8-23-13-45-13-71C16 103 119 0 244 0c117 0 213 88 226 202c1 8 1 18 1 26c0 53-18 102-48 140"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Uicon(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M394 654v-47c-42 40-110 64-172 64s-108-24-150-64C28 564 0 505 0 439V128h72v315c2 87 63 157 150 157c88 0 169-70 172-157V128h72v526z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UupperCase(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M527 0v491c0 71-28 134-72 181c-48 51-115 82-191 82s-144-31-192-82C28 625 0 561 0 491V0h72v492c1 105 87 191 192 191s190-86 191-190V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ubuntu(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M527 135c-34-19-45-62-26-96c13-23 37-35 61-35c12 0 24 3 35 9c34 20 45 63 26 96c-13 23-37 35-61 35c-12 0-24-3-35-9m175 199H586c-4-54-31-102-72-133c-7-5-14-10-22-15c-26-14-56-22-88-22c-33 0-63 8-90 23l-64-97c45-27 97-42 154-42c22 0 44 2 66 7c-8 39 9 80 45 101c14 8 30 12 47 12c23 0 45-8 62-23c46 50 75 116 78 189M228 105l64 97c-40 31-67 78-70 132c-1 4-1 8-1 13c0 4 0 9 1 13c4 57 34 106 78 137l-61 99c-58-38-102-96-122-164c28-16 48-47 48-82s-20-67-50-83c19-66 59-123 113-162M70 280c39 0 71 31 71 70s-32 71-71 71s-70-32-70-71s31-70 70-70m192 330l61-100c24 12 52 19 81 19c32 0 62-8 89-23c8-4 15-9 22-15c40-30 67-77 71-131h116c-3 76-35 145-85 196c-15-11-34-17-53-17c-17 0-33 4-47 12c-32 19-49 53-47 87c-22 5-44 8-66 8c-52 0-100-13-142-36m302-47c24 0 48 12 61 35c19 34 7 78-26 97c-11 6-23 9-35 9c-25 0-48-12-61-35c-20-34-8-77 26-97c11-6 23-9 35-9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Umbrella(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M288 589V347c10-3 20-4 26-4s16 1 26 4v242c0 42-28 105-105 105c-76 0-104-63-104-105c0-14 11-26 26-26c14 0 26 12 26 26c0 13 5 53 52 53c50 0 53-44 53-53m52-549v28c161 13 288 148 288 312h-4c-12-45-52-69-101-69s-89 24-101 69h-7c-10-36-38-55-75-65c-8-2-17-3-26-3s-18 1-26 3c-37 10-66 29-75 65h-8c-11-45-52-69-100-69c-49 0-90 24-101 69H0C0 216 127 81 288 68V40c0-14 11-26 26-26c14 0 26 12 26 26"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Underline(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 16l2 37c4 1 10 2 18 2c16 0 35 1 49 5s18 12 20 27c2 16 2 36 2 53v158c0 53 3 100 19 139c17 39 47 71 102 91c39 14 82 20 124 20s90-7 133-22c42-15 80-42 104-81c10-18 17-44 21-71c3-26 4-50 4-72c0-7 0-31-2-59c0-14-1-29-1-45c-2-32-3-60-4-79c-1-8-2-14-2-16c-2-8-2-15-2-23s3-14 8-17c11-7 29-9 46-10c8-1 17-2 23-4c2-4 2-8 2-12c0-7-1-15-3-22h-9c-11 0-22 0-33 2c-11 1-23 2-34 2c-25 0-48-1-72-3s-47-3-72-1v35c4 2 14 2 25 2h15c5-1 9-1 11-1c13 0 24 5 33 14c2 2 5 10 7 20c2 11 4 24 6 39c3 30 7 67 8 100c2 34 3 62 3 73c0 67-10 117-39 151c-28 34-73 51-146 51c-91 0-137-57-142-142c-3-79-6-156-6-235c0-23 2-40 4-48c5-18 18-21 43-19c11 1 28 0 46-3v-3c0-5 0-12-1-17c-1-6-1-11 0-16c-56 2-116 4-173 4c-19 0-40-1-59-2c-19-2-41-3-61-3H8c-2 0-5 1-8 1m0 665h666v-68H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Underscore(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 737h524v60H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Undo(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M248 270v66c0 14-11 20-22 12L9 192c-12-8-12-23 0-31L226 4c11-8 22-3 22 11v65c49 0 150 6 216 49c151 99 238 423-256 587c0 0 298-149 246-354c-13-55-74-104-206-92"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlock(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M468 332h67c9 0 16 8 16 17v352c0 9-7 16-16 16H16c-9 0-16-7-16-16V349c0-9 7-17 16-17h368V176c0-51-42-94-94-94h-31c-51 0-93 43-93 94v59H81v-67C81 76 140 0 231 0h86c92 0 151 76 151 168zM228 615h94l-23-101c16-8 27-25 27-44c0-29-23-51-51-51s-50 22-50 51c0 19 11 36 27 44z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Up(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M360 337L115 582L0 464l354-355l363 363l-112 112z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Upload(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M132 273h129v211c0 18 15 33 33 33h129c18 0 34-15 34-33V273h129c17 0 22-9 10-22L382 38c-6-7-15-11-24-11s-17 4-23 11L122 251c-12 13-8 22 10 22M0 422v228c0 10 5 16 16 16h684c11 0 17-6 17-16V422c0-10-8-18-17-18h-65c-9 0-17 9-17 18v145H98V422c0-10-7-18-16-18H16c-9 0-16 9-16 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func User(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M646 516c47 26 74 70 71 115c-3 28-3 29-39 33c-23 3-163 5-307 5c-163 0-333-2-346-5c-48-13-22-102 43-143c50-33 152-82 179-88c38-8 43-31 1-106c-10-17-22-68-23-122c-2-86 16-145 91-173c14-5 30-8 45-8c50 0 97 28 116 69c27 55 16 199-13 251c-32 60-29 78 8 88c24 6 99 43 174 84"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ustream(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M373 534c-71 0-113-55-113-163V29h229v342c0 111-41 163-116 163m257-172V39c12 9 20 24 20 40v550c0 28-22 50-50 50H50c-28 0-50-22-50-50V79c0-28 22-50 50-50h68v331c0 199 92 288 251 288c163 0 261-92 261-286"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func V(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m249 495l169-367h79L249 667L0 128h79z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VupperCase(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M548 0L273 762L0 0h82l191 526L466 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vertical(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M68 3h237c37 0 68 30 68 68v551c0 38-31 69-68 69H68c-37 0-68-31-68-69V71C0 33 31 3 68 3m597 279h-52c3-35-7-69-34-96c-20-20-47-33-75-35v69l-100-96l100-92v68c41 2 79 19 110 50c37 36 54 85 51 132M44 556h286V133H44zm368-239h237c37 0 68 31 68 69v236c0 38-31 69-68 69H388c10-12 18-28 22-44h178V361H412zM186 650c20 0 37-18 37-38s-17-37-37-37s-36 17-36 37s16 38 36 38"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Video(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m496 258l219-126c29-17 53-5 53 30v372c0 34-24 48-53 31L496 438v94c0 35-28 63-63 63H61c-34 0-61-28-61-63V161c0-35 27-62 61-62h372c35 0 63 27 63 62z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func View(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M397.25 278c38 0 69 31 69 69s-31 68-69 68s-68-30-68-68s30-69 68-69m0-170c226 0 389 212 389 212c11 14 11 39 0 53c0 0-163 212-389 212s-389-212-389-212c-11-14-11-39 0-53c0 0 163-212 389-212m0 410c94 0 171-77 171-171s-77-171-171-171s-171 77-171 171s77 171 171 171"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vimeo(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M130 29h390c72 0 130 58 130 130v390c0 72-58 130-130 130H130C58 679 0 621 0 549V159C0 87 58 29 130 29m408 216c4-20 4-41-9-57c-17-21-54-22-79-18c-20 3-89 33-113 106c42-3 64 3 60 49c-2 20-12 41-23 61c-12 23-36 69-66 36c-28-30-26-86-32-124c-4-21-8-48-14-70c-6-18-20-41-37-46c-18-5-40 3-53 11c-42 24-74 59-110 88v3c7 6 9 18 20 19c25 4 48-23 65 5c10 17 13 35 19 54c9 24 16 51 23 79c11 47 26 119 67 136c21 9 53-3 69-13c43-25 77-62 105-100c66-89 103-190 108-219"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vk(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M530 339c-16-6-16-29-17-45c-4-57 16-144-8-181c-17-22-99-20-149-17c-14 2-31 5-43 11s-24 16-24 25c0 13 30 11 41 28c12 18 12 57 12 89c0 37-6 86-21 88c-23 1-36-22-48-39c-24-33-48-74-66-114c-9-21-14-44-27-53c-20-14-56-15-91-14c-32 1-78-3-87 16c-7 21 8 41 16 58c41 89 85 167 139 241c50 69 97 124 189 153c26 8 140 31 163 0c8-12 6-39 10-59s9-40 28-41c16-1 25 13 35 23c11 11 20 20 28 30c19 19 39 45 63 55c33 14 84 10 132 8c39-1 67-9 70-32c2-18-18-44-30-59c-30-37-44-48-78-82c-15-15-34-31-34-49c-1-11 8-21 16-32c35-52 70-89 102-143c9-16 30-53 22-71c-9-20-59-14-90-14c-40 0-92-3-102 5c-19 13-27 34-36 54c-20 46-47 93-75 128c-10 12-29 37-40 33"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Volume(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 272v150c0 9 7 15 15 15h84l169 129c27 28 47 18 47-20V147c0-39-20-47-47-19L99 257H15c-8 0-15 6-15 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Volumedown(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 272v150c0 9 7 15 15 15h84l169 129c27 28 47 19 47-20V147c0-38-20-47-47-19L99 257H15c-8 0-15 7-15 15m380 237c5 3 10 5 15 5c12 0 22-6 28-17c28-47 41-97 41-150c0-52-13-103-41-151c-9-15-28-19-43-11s-20 28-12 43c23 38 32 78 32 119s-9 81-32 119c-8 15-3 35 12 43"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Volumeup(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M582 659c5 4 11 5 17 5c10 0 22-5 27-14c60-91 91-193 91-301c0-109-31-211-91-303c-9-13-29-16-44-8c-15 9-19 29-10 44c53 81 81 172 81 267c0 94-29 185-81 267c-9 13-5 34 10 43m-101-75c4 4 11 5 17 5c10 0 21-5 26-14c44-70 66-145 66-226c0-82-22-157-65-227c-10-15-29-20-44-10c-15 8-19 28-10 43c75 119 75 267 0 386c-9 15-5 35 10 43M0 273v151c0 9 7 15 15 15h84l169 129c27 27 47 19 47-19V148c0-38-20-46-47-19L99 258H15c-8 0-15 7-15 15m380 238c5 3 10 5 15 5c12 0 22-7 28-17c28-48 41-98 41-150c0-53-13-104-41-152c-9-15-28-19-43-11s-20 28-12 43c23 39 32 78 32 120c0 41-9 80-32 119c-8 15-3 35 12 43"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func W(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m590 487l156-362h79L591 667L412 261L234 667L0 125h79l156 362L412 82z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WupperCase(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M665 550L806 24h82L676 786L444 206L212 786L0 24h82l140 526L444 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Walking(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M304 73c0-40-33-73-74-73c-40 0-72 33-72 73c0 41 32 73 72 73c41 0 74-32 74-73m-79 644h-65s-3-185 2-207c4-22 47-68 49-86c2-19-17-97-17-97s-43 39-57 46s-124 25-124 25L0 342s105-23 120-35c15-11 55-108 85-120c20-9 74-5 101-5c28 0 120 53 129 65c10 11 65 119 65 119l-50 28l-40-86s-31-26-43-30c-13-4-30-9-33 2c-4 10 23 92 30 115c6 24 39 133 51 156s83 125 83 125l-61 40s-87-118-98-135c-12-17-46-108-46-108s-58 71-63 93c-4 20-5 151-5 151"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Web(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 342v-1C9 225 51 138 124 83C196 27 274 0 359 0c89-1 172 29 246 89c75 59 112 147 112 261v7c0 115-37 201-112 260c-75 60-158 90-247 90h-4c-87-1-167-31-241-91C39 557 1 466 0 342m337-157V27h-24c-5 6-9 13-13 19s-8 12-12 19c-4 6-9 14-12 20c-4 6-7 12-10 19c-6 10-11 20-15 29c-5 9-9 18-12 26c6 4 13 9 21 12c9 3 18 6 27 8c10 2 19 3 27 4c9 1 17 2 23 2m42-158v157c5 0 11 1 17 0c6 0 13-1 19-2c11-2 24-5 34-8c11-4 21-9 28-15c-12-29-25-52-38-74s-28-41-44-57v-1zm-108 9v-1c-7 3-14 8-22 11c-7 3-16 7-23 11c-14 7-28 15-42 24c-13 9-26 19-37 29c4 4 9 8 13 11s9 7 14 11c4 2 9 5 14 9c5 3 11 7 17 11c8-18 16-37 26-54c10-18 21-35 32-50c1-2 2-4 4-6c1-1 2-4 4-6m306 74v-1c-25-19-47-35-69-45s-44-20-64-28c16 16 30 36 39 55c10 20 20 41 28 61c4-2 9-4 15-8c5-3 12-6 18-9c6-4 11-8 17-12c6-5 12-9 16-13m-17 223h115c0-37-8-72-23-108c-14-35-34-64-59-89v-1c-5 8-12 15-19 20c-8 5-15 9-23 13c-5 3-10 7-15 9c-5 3-12 5-17 8c4 8 8 17 12 25c3 9 8 19 10 28c6 16 10 33 14 50c3 15 5 31 5 45M198 185v-1c-8-4-17-7-24-11c-7-5-14-10-20-14c-6-3-12-7-17-11s-9-8-13-12c-26 25-45 53-57 86c-13 33-22 70-26 110h123c0-28 4-55 11-82c6-28 15-49 23-65m139 147V217c-8 1-18 1-27 0c-10-1-20-4-29-7c-8-2-18-4-26-7c-9-2-18-6-25-10c-6 10-11 22-15 35c-5 12-8 25-11 37c-2 12-3 25-4 36s-2 22-2 31zm42-113v113h148c0-6 0-13-1-21c-2-7-3-16-4-24c-3-12-6-25-10-37c-4-11-8-22-13-31c-2-5-4-11-6-15c-3-5-6-8-8-11c-11 8-26 13-44 17c-17 4-37 8-55 9zM164 365H41c0 17 3 36 8 59c5 24 15 46 26 70c5 12 11 24 17 36c7 11 15 23 23 34c6-4 12-7 17-10c6-3 13-7 19-10c7-3 14-6 22-10c7-3 16-7 25-11c-8-25-17-49-23-76c-7-26-11-52-11-81zm173 117V365H198c0 8 1 20 3 32c1 12 4 25 6 38c4 14 7 28 11 40s8 22 12 31c18-7 36-12 48-15c13-4 25-7 37-8h11c4-1 8-1 11-1m42-117v116c7 1 15 2 23 3c9 1 20 3 29 5l15 3c5 2 11 3 16 4c6 2 12 3 18 5c5 2 10 3 14 5c13-33 21-60 26-83c5-22 7-41 7-57v-1zm296 1v-1H560v6c-1 19-4 41-8 64c-4 22-13 48-25 80c16 8 31 16 43 24s23 17 31 25c17-17 32-40 45-70c13-29 22-59 27-91c1-6 1-12 2-18zM337 670V514c-25 4-46 9-62 13s-28 8-36 12c6 15 12 28 18 40c6 11 13 23 19 33c2 4 6 8 9 13c3 4 5 9 8 13c3 5 6 11 9 17c4 5 7 11 11 15zm42 0h23c9-7 18-17 25-28c8-11 17-24 23-35c7-13 14-25 20-37s11-23 15-31c-11-4-26-8-41-12s-36-8-65-12zm198-82v-1c-2-3-6-7-9-10s-8-7-13-10c-4-3-9-6-15-9s-14-7-21-11c-4 8-11 22-21 41c-9 19-25 42-45 67c25-4 47-13 67-25c22-11 41-26 57-42m-372-32v-1c-7 4-18 9-29 14c-12 5-25 11-37 19c7 6 14 12 20 16c6 5 13 9 19 13c11 7 24 14 37 20s29 12 49 18c-7-8-12-18-17-26s-11-16-16-24s-9-17-13-25c-5-8-9-16-13-24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wifi(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M486 476c-43-42-110-42-153 0c-42 42-42 110 0 153c43 42 110 42 153 0c42-43 42-111 0-153m-226-73c-28 28-46 61-55 97l-67-65c14-34 35-64 63-91c115-116 302-116 417 0c27 27 49 57 63 91l-67 65c-9-36-27-69-55-97c-83-83-217-83-299 0M127 270c-26 27-49 56-66 87L0 296c19-30 42-58 68-85c188-188 496-188 684 0c26 27 48 55 67 85l-61 61c-17-31-40-60-66-87c-156-156-409-156-565 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Windows(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m444 82l-73 253s-83-58-144-52c-60 6-141 36-141 36l74-255s153-82 284 18m-34 277l73-251c22 15 61 42 124 50c63 7 161-33 161-33l-73 253c-11 4-57 30-141 35c-85 5-144-54-144-54M0 618l73-255s82-38 151-33c68 4 123 42 134 51l-73 255s-84-55-142-50S34 601 0 618m324 40l73-253s78 51 129 49c51-1 86-4 155-31h1c-5 13-73 251-73 251s-80 38-154 34c-73-3-117-45-131-50"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wink(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M358 0c198 0 359 155 359 349S556 698 358 698S0 543 0 349S160 0 358 0m229 332l8-21c3-8-1-17-9-20l-117-42l108-62c7-4 9-13 5-20l-12-20c-4-7-13-9-19-5l-161 92c-7 4-8 10-7 23s3 18 10 21l175 63c7 3 16-2 19-9M258 171h-41c-11 0-19 9-19 20v100c0 11 8 21 19 21h41c10 0 20-10 20-21V191c0-11-10-20-20-20m-30 242l-39 13c-4 1-7 4-8 7c-2 3-2 7-1 11s36 105 178 104c142 1 178-100 179-104s1-8-1-11c-1-3-5-6-9-7l-34-12l-5-1c-6 0-11 5-13 10c0 1-22 61-117 61c-94 0-115-58-116-61c-2-5-8-10-14-10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wordpress(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M50 354c0-44 9-85 26-122l143 392C119 576 50 473 50 354m503-15c0 26-10 55-23 97l-30 100l-109-323c18-1 35-3 35-3c16-2 14-25-2-24c0 0-49 3-80 3c-30 0-79-3-79-3c-16-1-18 23-2 24c0 0 15 2 31 3l47 128l-66 197l-109-325c18-1 34-3 34-3c17-2 15-25-2-24c0 0-48 3-80 3H99c54-81 146-135 251-135c78 0 149 30 203 79c-2-1-3-1-4-1c-30 0-51 26-51 54c0 24 15 45 30 70c11 20 25 46 25 83M265 642l90-262l93 253c0 2 1 3 2 4c-31 11-65 17-100 17c-29 0-58-4-85-12m348-432c24 43 37 92 37 144c0 111-60 207-149 260l92-266c17-42 22-77 22-107c0-11 0-21-2-31M350 4c193 0 350 157 350 350S543 704 350 704S0 547 0 354S157 4 350 4m0 684c184 0 334-150 334-334S534 20 350 20S16 170 16 354s150 334 334 334"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wrench(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m247 245l10-85L122 23c29-14 62-23 97-23c121 0 219 97 219 219c0 18-3 38-7 56l249 248c45 46 49 116 8 157s-112 38-158-8L287 427c-22 7-44 12-68 12C98 439 0 340 0 219c0-34 7-64 20-93l139 141z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func X(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M519 128L315 380l222 274h-93L269 437L93 654H0l223-274L19 128h93l157 194l156-194z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func XupperCase(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M535 0L316 357l232 378h-84L274 427L85 735H0l231-378L12 0h84l178 289L451 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Y(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m249 494l169-366h78L149 881H69l140-301L0 128h79z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YupperCase(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M262 289L439 0h85L298 368v367h-72V367L0 0h85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yahoo(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m741 210l22-33c-25-2-112-2-152-2c-48 0-143-3-161 0l7 35c10 0 76 10 96 13c-4 31-132 119-147 141c-26-39-108-138-160-215c20-5 97-7 113-11l3-27c-15-2-136-3-189-2c-81 3-164-2-173 2l1 31c18 4 93 11 105 15c41 32 202 232 208 252c3 19 3 39 3 49v23c0 47 2 56-5 65c-16 16-88 12-105 13l-6 34c32 1 134-3 166-3c63 0 173 0 188 1l2-35c-16-3-102-1-114-3c-3-18-9-50-9-70l2-27c0-18 1-36 4-48c12-35 194-175 213-180c19-4 74-17 88-18m-33 256l53 24c9-13 115-186 144-223l35-46c-13-5-17-9-54-26c-62-29-60-27-74-35l-13 45c-7 23-76 235-91 261m-52 110l29 12l30 13l31-58l-33-13l-35-18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yapcasia(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M65 701h6v-45h12v-6H53v6h12zm105-25c0-15-12-27-26-27s-26 12-26 27s11 26 26 26c14 0 26-11 26-26m-6 0c0 11-9 20-20 20s-20-9-20-20s8-21 20-21s20 10 20 21m81-26h-8l-21 22v-22h-7v51h7v-21l2-3l19 24h8l-22-28zm40 0h-8l17 29v22h6v-22l16-29h-7l-12 22zm118 26c0-15-12-27-26-27c-15 0-27 12-27 27s12 26 27 26c14 0 26-11 26-26m-6 0c0 11-10 20-20 20c-11 0-20-9-20-20s8-21 20-21c11 0 20 10 20 21m109 25v-6h-20l17-20c2-3 3-7 3-10c0-9-7-16-15-16c-9 0-15 7-16 15h6c0-5 5-9 10-9s9 5 9 10c0 3-2 6-4 9l-22 27zm90-25c0-15-11-27-26-27c-14 0-26 12-26 27s12 26 26 26c15 0 26-11 26-26m-6 0c0 11-9 20-20 20s-20-9-20-20s8-21 20-21s20 10 20 21m49-26l-4 6h9v45h6v-51zm92 51v-6h-20l17-20c2-3 3-7 3-10c0-9-7-16-16-16s-14 7-15 15h6c0-5 4-9 9-9c6 0 10 5 10 10c0 3-2 6-4 9l-22 27zM497 557h-22l11-33zm298 19c4 4 5 9 5 14c0 4-1 10-6 14c-4 4-9 5-14 5h-23c-5-1-10-1-14-5c-2-1-3-2-3-3l-3 3c-5 4-10 5-14 5h-20c-4 0-8-1-13-4c-4 3-8 4-12 4h-45c-5 0-10-1-15-5c-1-1-1-2-2-3c-9 7-20 10-33 10c-5 0-10 0-14-1c-3 0-5-1-7-1c-4 2-7 2-10 2c-4 0-9-1-14-5c-4 2-8 3-12 3h-23c-4-1-9-1-14-5l-3-3l-3 3c-4 4-9 5-14 5h-20c-5-1-10-1-14-5c-3-3-5-6-6-10c-3 1-6 2-9 2c-7 0-13-3-18-7h-1c-5 4-12 7-18 7h-1c-5 0-11-2-15-5c-1 1-2 3-4 5c-3 3-7 6-12 8c-9 5-20 7-31 7h-1c-15 0-29-6-39-17c-7-7-12-17-14-27c-7 5-16 8-26 9c5 4 6 9 6 14c0 4-1 10-6 14c-4 4-9 5-14 5h-26c-4 0-8-1-13-4c-4 3-8 4-12 4h-23c-5-1-10-1-14-5c-2-1-3-2-3-3l-3 3c-5 4-10 5-15 5H98c-5-1-9-1-14-4c-4 3-8 4-13 4H31c-4 0-9-1-14-5c-4-4-5-10-5-14c0-5 1-10 5-14c5-4 10-5 14-5v-4l-18-33c-3-1-6-2-8-5c-4-4-5-9-5-13c0-5 1-10 6-14c4-4 9-5 14-5h20c4 0 8 1 12 4c4-3 9-4 13-4h17c5 0 10 1 14 5h2c4-4 9-5 14-5h29c4 0 8 1 12 4c3 3 6 6 7 10c0 0 0 1 1 1l21 60c2 1 4 1 5 3c1-1 2-1 3-2v-41c-2 0-4-2-5-3c-4-4-5-9-5-13c0-5 1-10 6-14c4-4 9-5 14-5h32c13 0 25 2 34 9c4 3 7 7 9 11c2-2 3-4 5-6c10-11 24-17 39-17c5 0 9 1 14 2c2 0 4 1 6 2v-1c3-2 7-3 10-3c4 0 9 2 13 6c3 4 4 7 4 11c5-4 11-6 17-6s13 2 18 7h1c5-5 12-7 18-7c7 0 14 2 19 7c2 2 3 4 5 7v-3c0-5 1-10 5-14c5-4 10-5 15-5h28c4 0 9 1 12 4c4 3 6 6 8 10v1l21 58c0-3 1-7 5-11c1-1 3-2 4-3c-5-6-7-14-7-22c0-11 5-22 13-29c9-7 20-11 32-11c5 0 9 1 13 2c2 0 3 1 5 1c4-2 7-3 10-3c4 1 8 1 13 6c4-2 8-3 12-3h45c5 0 10 1 15 5c3 2 4 6 5 9c0-4 2-7 5-9c5-4 10-5 14-5h29c4 0 9 1 12 4c4 3 6 6 7 10l1 1l21 60c3 1 6 2 8 4M79 596c1-1 2-3 2-6s-1-5-2-7c-2-1-4-1-8-1H61v-19l21-39h1c3 0 5-1 6-2c2-1 3-4 3-6c0-3-1-6-3-7c-1-1-3-2-7-2H65c-4 0-6 1-7 2s-2 4-2 7c0 2 1 5 2 6s3 2 5 2h1l-12 25l-12-25h1c3 0 4-1 6-2c1-1 1-4 1-6c0-3 0-6-2-7c-1-1-3-2-6-2H20c-4 0-6 1-8 2c-1 1-2 4-2 7c0 2 1 5 2 6c2 1 4 2 6 2h2l22 41v17H31c-3 0-6 0-7 1c-1 2-2 4-2 7s1 5 2 6c2 2 4 2 7 2h40c4 0 6 0 8-2m-7-31v6c4 0 8 1 12 4l6-3l14-39c-2 0-4-1-6-3l-1-1c-3 3-6 4-8 5zm110 31c2-1 3-3 3-6s-1-5-2-6c-2-2-4-2-7-2h-1l-24-67c-2-5-6-8-10-8h-29c-3 0-6 1-7 2c-2 1-3 4-3 7s1 5 3 6c1 1 4 2 7 2h6l-20 58h-2c-2 0-5 0-6 2c-1 1-2 3-2 6s1 5 2 6c2 2 4 2 8 2h19c4 0 6 0 8-2c1-1 2-3 2-6s-1-5-2-7c-2-1-4-1-8-1h-2l3-9h33l3 9h-2c-4 0-6 0-8 1c-1 2-2 4-2 7s1 5 2 6c2 2 4 2 8 2h23c3 0 6 0 7-2m78-38c7-5 10-12 10-22s-3-17-10-22c-6-4-16-7-28-7h-32c-3 0-6 1-7 2c-2 2-3 4-3 7c0 2 1 5 3 6c1 1 3 2 6 2h1v58h-1c-3 0-5 0-7 2c-1 1-2 3-2 6s1 5 3 6c1 2 3 2 7 2h26c4 0 6 0 8-2c1-1 2-3 2-6s-1-5-2-7c-2-1-4-1-8-1h-7v-16h13c12 0 22-3 28-8m149-39c-3 3-5 6-5 11c0 4 2 8 5 10c3 3 7 4 11 4c5 0 9-1 12-4c3-2 4-6 4-10c0-5-1-8-4-11c-3-2-7-4-12-4s-8 2-11 4m-38 0c-3 3-4 6-4 11c0 4 1 8 4 10c3 3 7 4 12 4s8-1 11-4c3-2 5-6 5-10c0-5-2-8-5-11c-3-2-6-4-11-4s-9 2-12 4m-24 76c8-4 12-8 12-13c0-2-1-5-2-6c-2-2-3-3-5-3s-6 2-13 6c-6 3-12 5-17 5c-8 0-14-3-19-8c-4-6-6-13-6-23s2-18 6-23c4-6 10-8 18-8c6 0 10 1 13 3c2 2 5 5 6 10c1 3 2 5 3 6s3 1 5 1c3 0 5 0 6-1c1-2 2-4 2-8v-20c0-3-1-5-2-6c-1-2-3-2-5-2s-3 0-4 1s-2 2-3 4c-4-2-8-3-12-4c-3-1-7-1-11-1c-12 0-23 4-32 13c-8 9-12 21-12 35s4 26 12 34c9 9 19 14 33 14c10 0 19-2 27-6m17-42c1-1 2-2 4-3c-2 0-3-1-4-2c0-1-1-1-1-2c-1 1-1 1-1 2c-5 4-10 5-14 5s-8-1-12-4c-4-4-5-7-6-11s-2-5-3-5c0 0-2-1-6-1c-5 0-7 1-10 4c-2 3-4 8-4 17c0 8 2 14 4 16c3 3 5 5 11 5c3 0 7-1 12-4c4-2 7-4 9-5c3-1 4-2 9-2c2 0 4 0 6 1c1-4 3-8 6-11m30 28c3-2 5-6 5-10c0-5-2-8-5-11c-3-2-6-4-11-4s-9 2-12 4c-3 3-4 7-4 11s1 7 4 10s7 4 12 4s8-1 11-4m8-28l3-3c-1 0-2-1-3-2h-1c-1 1-2 2-3 2l3 3zm30 28c3-2 4-6 4-10c0-5-1-8-4-11c-3-2-7-4-12-4c-4 0-8 2-11 4c-3 3-5 7-5 11s2 7 5 10s7 4 11 4c5 0 9-1 12-4m3-31c2 1 3 2 4 3c3 2 5 6 6 10l11-30c-3 0-5-1-7-3c-1-1-2-3-3-4v4c0 6-2 13-7 18c-1 1-2 2-4 2m99 46c1-1 2-3 2-6s-1-5-2-6c-1-2-4-2-6-2h-2l-24-67c-2-5-5-8-10-8h-28c-4 0-6 1-8 2c-1 1-2 4-2 7s1 5 2 6c2 1 4 2 8 2h6l-21 58h-1c-3 0-5 0-7 2c-1 1-2 3-2 6s1 5 3 6c1 2 4 2 7 2h20c3 0 6 0 7-2c2-1 3-3 3-6s-1-5-3-7c-1-1-4-1-7-1h-3l4-9h33l3 9h-3c-3 0-6 0-7 1c-2 2-3 4-3 7s1 5 3 6c1 2 4 2 7 2h23c4 0 6 0 8-2m76-3c7-5 10-12 10-22c0-7-2-13-7-18c-4-4-11-7-21-10c-1 0-3 0-5-1c-13-2-20-6-20-11c0-3 2-6 4-8c3-2 7-3 12-3c3 0 6 1 9 2c3 2 5 4 7 7v2c2 4 5 6 8 6s5-1 7-2c1-1 1-3 1-7v-15c0-3-1-6-1-7c-1-1-3-1-6-1c-2 0-4 0-5 1c-1 0-1 1-2 3c-4-2-7-3-10-3c-4-1-7-1-11-1c-10 0-19 2-25 8c-7 6-10 13-10 21c0 7 2 13 6 17s11 7 20 9c3 1 5 2 9 2c12 3 18 7 18 13c0 3-2 6-4 8c-3 2-7 3-11 3c-5 0-10-1-13-3c-4-2-7-5-9-9c0 0 0-1-1-2c-2-5-5-8-8-8s-6 1-7 2c-1 2-2 4-2 7v19c1 4 1 6 3 7c1 1 3 2 6 2c2 0 4-1 5-1l3-3c3 1 7 2 10 3c4 0 8 1 13 1c12 0 21-3 27-8m-39-22l3 3c2 1 5 2 9 2c1 0 3 0 3-1c-1 0-4-2-8-3c-3 0-5-1-7-1m17-39l-1-1c-1-1-2-1-4-1h-3c3 1 5 2 8 2m30 12c1 1 2 1 2 2c7 7 10 16 10 25h5v-37h-2c-3 0-5 0-8-1c0 3-2 6-4 9c-1 1-2 1-3 2m68 52c1-1 2-3 2-6s-1-5-2-7c-2-1-4-1-8-1h-13v-58h13c4 0 6-1 8-2c1-1 2-3 2-6s-1-5-2-7c-2-1-4-2-8-2h-45c-4 0-6 1-8 2c-1 2-2 4-2 7s1 5 2 6c2 1 4 2 8 2h13v58h-13c-4 0-6 0-8 1c-1 2-2 4-2 7s1 5 2 6c2 2 4 2 8 2h45c4 0 6 0 8-2m-10-25h2c4 0 8 1 13 3c1-1 3-1 4-2l14-39c-2 0-4-1-6-3c-3-3-5-6-5-9c-1 3-2 6-5 9c-5 4-10 4-15 4h-2zm112 25c1-1 2-3 2-6s-1-5-2-6c-2-2-4-2-7-2h-1l-24-67c-2-5-5-8-10-8h-29c-3 0-6 1-7 2c-2 1-3 4-3 7s1 5 3 6c1 1 4 2 7 2h6l-20 58h-2c-2 0-4 0-6 2c-1 1-2 3-2 6s1 5 2 6c2 2 4 2 8 2h20c3 0 6 0 7-2c2-1 2-3 2-6s0-5-2-7c-1-1-4-1-7-1h-3l3-9h33l4 9h-3c-4 0-6 0-8 1c-1 2-2 4-2 7s1 5 2 6c2 2 4 2 8 2h23c4 0 6 0 8-2m-37-39h-22l11-33zm-508-30c3 2 5 5 5 9s-2 7-5 9s-7 3-14 3h-10v-24h10c6 0 11 1 14 3m-6 10v-1c0-1-3-2-8-2v4c5 0 8-1 8-1m-91 20h-23l12-33zm427-309c0 63-34 119-85 149l10 40c1 5-1 11-4 16c-4 4-9 7-15 7h-43c-2 0-4 0-6-1c-2 1-5 1-7 1h-46c-2 0-4 0-7-1c-2 1-4 1-6 1h-43c-6 0-11-3-15-7c-3-5-5-11-3-16l9-40c-51-30-85-86-85-149c0-86 62-158 145-171l8-55c2-9 10-16 19-16c10 0 18 7 19 16l9 55c83 12 146 84 146 171m-92-51H319l29 51h104zM407 75l-8-50l-8 50l-4 27h24zm17 67l14-27h-76l14 27l-6 42h60zM321 441h43l2-23l13-156h-17l-33 144zm58-21l-2 21h46l-2-21l-13-158h-16zm57 21h43l-8-35l-33-144h-17l13 156z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yapcasialogo(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M612 426c0 112-61 209-150 263l16 70c2 10 0 20-7 28c-6 8-16 13-26 13h-75c-4 0-8-1-12-2c-4 1-8 2-12 2h-81c-4 0-8-1-11-2c-4 1-8 2-12 2h-75c-11 0-20-5-27-13c-6-8-8-18-6-28l16-70C60 635 0 538 0 426c0-152 111-278 256-302l15-96c3-16 17-28 33-28c17 0 31 12 33 28l16 96c146 23 259 149 259 302m-164-89H164l50 90h183zM318 121l-14-87l-13 87l-8 48h43zm30 118l25-47H238l25 47l-10 75h106l-10-75zM167 766h75l3-40l24-274h-31l-57 253zm101-37l-3 37h81l-3-37l-24-277h-27zm102 37h75l-14-61l-58-253h-30l23 274zm459-97h15V562h28v-14h-72v14h29zm254-61c0-34-29-63-64-63c-34 0-63 29-63 63c0 36 28 64 63 64s64-28 64-64m-15 0c0 28-22 49-49 49c-26 0-48-21-48-49c0-26 20-48 48-48c29 0 49 22 49 48m196-60h-20l-51 52v-52h-15v121h15v-51l6-5l46 56h20l-55-67zm95 0h-17l40 68v53h15v-53l39-68h-17l-30 52zm286 60c0-34-29-63-63-63c-35 0-64 29-64 63c0 36 29 64 64 64s63-28 63-64m-15 0c0 28-22 49-48 49s-49-21-49-49c0-26 20-48 49-48c28 0 48 22 48 48m266 61v-14h-49l40-48c5-7 9-16 9-25c0-21-18-37-38-37c-22 0-36 15-39 36h15c1-12 11-22 24-22c12 0 23 11 23 23c0 8-5 16-10 22l-53 65zm217-61c0-34-28-63-63-63s-63 29-63 63c0 36 28 64 63 64s63-28 63-64m-15 0c0 28-22 49-48 49s-48-21-48-49c0-26 20-48 48-48s48 22 48 48m117-60l-8 14h21v107h14V548zm224 121v-14h-49l40-48c6-7 9-16 9-25c0-21-17-37-38-37c-22 0-36 15-38 36h15c1-12 11-22 23-22c13 0 24 11 24 23c0 8-6 16-11 22l-53 65zM1235 305c8 5 12 13 12 23c0 11-4 19-12 24c-7 5-19 8-35 8h-28v-63h28c15 0 27 3 35 8m-14 25v-3h-1c-1-1-8-4-20-4h-2v10h2c13 1 20-2 21-3m1437 102c10 10 13 23 13 34c0 12-3 27-15 37s-25 11-37 12h-59c-12-1-25-2-37-12c-3-3-6-6-8-9c-1 3-4 6-7 9c-12 10-25 11-37 12h-51c-10-1-21-2-32-8c-10 6-21 8-31 8h-118c-12-1-24-2-36-12c-3-2-5-4-6-7c-23 18-53 25-86 25c-13 0-26-1-37-3c-5-1-11-2-16-3h-2c-8 5-16 6-24 6c-11 0-23-1-35-12l-1-1c-10 6-21 7-30 7h-60c-12-1-25-2-37-12c-3-3-5-6-7-9c-2 3-5 6-8 9c-12 10-25 11-37 12h-51c-12-1-24-2-37-12c-8-8-12-17-13-26c-8 3-16 4-24 4h-1c-17 0-34-6-47-17v-1s-1 0-1 1c-13 12-30 17-48 17c-15 0-29-4-41-13c-3 5-6 9-9 13c-9 9-19 15-31 22c-24 12-52 18-81 18h-1c-40 0-76-15-102-43c-19-20-30-44-36-72c-18 14-41 21-66 24h1c11 11 14 25 14 36c0 12-3 26-14 36c-12 11-25 13-37 13h-68c-10 0-22-2-32-9c-11 7-22 9-32 9h-60c-12-1-25-2-37-12c-3-3-5-6-7-9c-2 3-5 6-8 9c-12 10-25 11-36 12h-52c-11-1-22-2-34-10c-11 8-22 10-33 10H690c-12 0-25-2-37-12c-11-11-14-25-14-37c0-11 2-25 14-36c12-10 25-11 36-11v-10l-46-87c-8-2-15-6-21-12c-10-10-13-24-13-35s2-25 14-36c12-10 25-12 37-12h51c10 0 22 2 32 10c11-9 22-10 32-10h45c12 0 25 1 37 12l2 2c0-1 1-2 2-2c12-10 25-12 37-12h74c10 0 22 3 31 10s15 16 18 26c1 1 1 2 1 3l57 155c4 1 9 3 13 6l6-3V318c-5-2-9-5-12-8c-10-10-13-24-13-35s3-26 14-36c12-10 25-12 37-12h82c35 0 66 6 89 24c10 7 18 17 24 27c3-4 7-9 11-14c26-29 63-44 102-44c12 0 23 2 34 4c5 1 10 3 15 5c1 0 1-1 1-1c8-5 17-8 26-8c10 0 25 4 34 16c7 9 9 18 10 27c12-10 28-15 43-15c17 0 35 5 48 17c0 1 0 1 1 1v-1c13-12 31-17 48-17s35 5 47 17c6 5 10 11 14 18c-1-3-1-5-1-8c0-12 3-26 15-36s25-12 37-12h74c10 0 22 3 31 10s15 16 18 26c1 1 1 2 1 3l55 150c1-9 3-19 12-28c3-4 7-7 11-8c-13-16-19-36-19-56c0-30 13-57 35-76s51-28 83-28c11 0 22 2 32 4c4 1 9 2 13 3c10-6 18-6 26-7c10 1 22 1 34 14c10-6 20-6 29-7h118c12 1 25 1 37 12c8 7 12 15 13 24c2-9 6-17 14-24c12-10 24-12 36-12h75c10 0 22 3 31 10s15 16 18 26c1 1 1 2 1 3l57 155c6 2 13 5 19 11M812 483c4-3 6-9 6-17c0-7-2-13-6-16s-10-5-19-5h-26v-49l53-99h3c7 0 13-2 16-5c4-4 6-10 6-17c0-8-2-13-6-17c-4-3-10-5-19-5h-45c-8 0-13 2-16 5c-4 4-5 9-5 17c0 7 1 13 4 17c3 3 8 5 14 5h1l-30 65l-31-65h3c6 0 11-2 14-6c3-3 4-9 4-16c0-8-1-13-5-17c-3-3-9-5-17-5h-51c-9 0-16 2-19 5c-4 4-6 9-6 17c0 7 2 13 5 17c4 3 9 5 16 5h4l56 105v43h-26c-9 0-16 2-19 5c-4 3-6 9-6 16c0 8 2 14 6 17c4 4 10 5 19 5h103c8 0 15-1 19-5m-19-80v16c11 0 22 1 33 9c5-3 10-6 15-7l36-100c-5-2-11-5-16-9c-1-1-2-2-2-3c-1 0-1 1-1 1c-6 6-14 10-21 12zm286 80c4-3 6-9 6-17c0-7-2-12-6-16c-3-3-9-5-16-5h-3l-63-172v-1c-5-12-13-19-25-19h-74c-9 0-16 2-20 5c-4 4-6 9-6 17s2 14 6 17s11 5 20 5h16l-54 148h-3c-7 0-13 2-16 5c-4 4-6 9-6 16c0 8 2 14 6 17c4 4 10 5 19 5h52c8 0 15-1 19-5c4-3 6-9 6-17c0-7-2-13-6-16s-11-5-19-5h-7l8-23h86l8 23h-7c-9 0-16 2-19 5c-4 3-6 9-6 16c0 8 2 14 6 17c3 4 10 5 19 5h60c9 0 15-1 19-5m200-98c17-13 25-33 25-58c0-24-8-43-25-55c-16-13-41-19-73-19h-82c-9 0-15 2-19 6c-4 3-6 9-6 16s2 13 6 16c4 4 9 6 15 6h4v148h-4c-6 0-12 2-15 5c-4 4-6 9-6 16c0 8 2 14 6 17c4 4 10 5 19 5h68c9 0 15-1 19-5c3-3 5-9 5-17c0-7-2-13-5-16c-4-3-10-5-19-5h-20v-41h34c32 0 57-6 73-19m383-100c-8 7-11 15-11 26c0 12 3 21 11 28s18 10 30 10s22-3 30-10c7-7 11-16 11-28c0-11-4-19-11-26c-8-7-18-10-30-10s-22 3-30 10m-96 0c-8 7-12 15-12 26c0 12 4 21 12 28c7 7 17 10 29 10c13 0 23-3 30-10c8-7 11-16 11-28c0-11-3-19-11-26c-7-7-17-10-30-10c-12 0-22 3-29 10m-63 194c20-10 30-21 30-33c0-6-1-12-5-16s-8-7-13-7c-4 0-15 5-32 14c-16 9-32 14-46 14c-20 0-35-7-47-21c-11-14-16-33-16-59s5-46 16-60s27-21 47-21c13 0 24 3 31 9c8 5 13 14 17 27c2 7 4 12 7 14c3 3 8 4 13 4c8 0 14-1 16-4c3-4 5-10 5-19v-51c0-9-2-15-5-18c-2-4-7-5-13-5c-5 0-8 1-11 3s-5 5-8 10c-10-5-20-8-29-10s-18-3-28-3c-33 0-61 11-82 35c-22 24-33 53-33 89c0 37 11 66 32 89c22 23 50 35 84 35c26 0 49-6 70-16m45-108c3-2 5-4 8-6c-3-2-6-4-8-7c-2-1-3-3-4-5c-1 2-2 3-3 5c-12 11-25 12-36 13c-9 0-21-3-30-11c-9-9-12-17-16-27c-2-10-5-12-6-13s-6-3-16-3c-15 0-20 3-26 10s-11 21-11 44c0 22 5 36 11 43s12 10 26 11c8 0 20-3 34-11c8-5 16-8 22-11c7-3 11-5 22-6c5 0 10 1 15 3c3-11 9-22 18-29m77 73c8-7 12-16 12-27s-4-20-12-27c-7-6-17-10-30-10c-12 0-22 4-29 11c-8 6-12 15-12 26s4 20 12 27c7 7 17 11 29 11c13 0 23-4 30-11m19-72c0-1 0-1 1-1c2-2 5-4 8-6c-3-2-6-4-9-7v-1l-1 1c-3 2-5 5-8 7c3 1 5 4 8 6c0 0 1 0 1 1m78 72c8-7 11-16 11-27s-3-20-11-27c-8-6-18-10-30-10s-22 4-30 11c-8 6-11 15-11 26s3 20 11 27s18 11 30 11s22-4 30-11m9-79c3 1 6 3 8 6c8 6 14 15 17 25l27-75c-5-2-11-5-16-9c-4-4-7-7-9-11c1 3 1 7 1 10c1 17-6 35-19 47c-3 2-6 5-9 7m254 118c4-3 6-9 6-17c0-7-2-12-6-16c-4-3-9-5-16-5h-3l-63-172v-1c-5-12-13-19-25-19h-74c-9 0-16 2-20 5c-4 4-6 9-6 17s2 14 6 17s11 5 20 5h15l-53 148h-3c-7 0-13 2-16 5c-4 4-6 9-6 16c0 8 2 14 6 17c4 4 10 5 19 5h52c8 0 15-1 19-5c4-3 6-9 6-17c0-7-2-13-6-16s-11-5-19-5h-7l8-23h86l8 23h-7c-9 0-16 2-20 5c-3 3-5 9-5 16c0 8 2 14 5 17c4 4 11 5 20 5h60c9 0 15-1 19-5m197-8c16-13 25-32 25-56c0-20-6-36-18-47c-11-11-29-20-55-25c-3-1-7-2-13-3c-34-7-50-17-50-29c0-8 3-15 11-21c7-5 16-8 28-8c10 0 18 2 25 6s13 10 17 19c0 1 1 2 1 3c5 10 12 15 21 15c8 0 13-1 16-4c3-4 5-9 5-17v-2l-1-38c0-9-1-14-4-17c-2-3-7-4-14-4c-6 0-9 0-12 2c-2 2-5 5-7 9c-8-4-17-7-26-9c-8-1-17-2-26-2c-27 0-49 7-66 21c-17 15-25 33-25 56c0 17 5 31 14 41c10 10 28 19 54 25c5 2 12 4 21 6c31 7 46 18 46 32c0 9-3 16-10 21s-16 8-28 8c-14 0-25-3-34-8s-16-12-22-23c0-1-1-4-3-7c-6-12-13-18-21-18s-14 1-17 5c-4 3-5 9-5 18v2l1 47c0 8 1 14 5 17c3 3 9 5 17 5c5 0 9-1 12-3c2-1 5-4 7-7c9 3 18 6 28 7c9 2 20 3 32 3c31 0 54-7 71-20m-101-58c2 5 5 7 9 9s11 4 21 4c4 0 7 0 9-1c-4-2-12-5-23-8c-6-1-12-2-16-4m44-99c-1-2-2-3-3-3c-3-2-6-3-12-3h-5c5 2 12 4 20 6m77 30c2 2 4 4 6 5c17 17 25 41 25 66h13v-96h-7c-6 0-13 0-19-2c-1 7-4 15-11 22c-2 2-4 4-7 5m174 135c4-3 6-9 6-17c0-7-2-13-6-16c-3-3-10-5-19-5h-33V297h33c9 0 16-2 19-5c4-3 6-9 6-17s-2-13-5-16c-4-4-11-5-20-5h-118c-9 0-15 1-19 5c-4 3-6 8-6 16s2 14 6 17s10 5 19 5h34v148h-34c-9 0-15 2-19 5s-6 9-6 16c0 8 2 14 6 17c4 4 10 5 19 5h118c9 0 15-1 19-5m-26-64h7c10 0 21 1 32 8c4-3 8-5 12-6l36-100c-6-2-11-5-16-9c-8-7-12-16-14-24c-1 8-5 17-13 23c-12 11-25 12-37 12h-7zm289 64c4-3 6-9 6-17c0-7-2-12-6-16c-4-3-9-5-16-5h-4l-62-172v-1c-5-12-13-19-25-19h-75c-9 0-15 2-19 5c-4 4-6 9-6 17s2 14 6 17s10 5 19 5h16l-53 148h-3c-7 0-13 2-16 5c-4 4-6 9-6 16c0 8 2 14 6 17c4 4 10 5 19 5h51c9 0 16-1 20-5c4-3 6-9 6-17c0-7-2-13-6-16s-11-5-20-5h-7l9-23h86l8 23h-7c-9 0-16 2-20 5s-5 9-5 16c0 8 1 14 5 17c4 4 11 5 20 5h59c9 0 16-1 20-5m-95-102h-57l29-84zm-654 0h-57l29-84zm-906 0h-57l29-84z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yapcasialogomark(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M550 383c0 101-54 189-134 237l14 63c2 9 0 18-6 26c-6 7-14 11-24 11h-67c-4 0-7-1-11-2c-3 1-7 2-10 2h-73c-4 0-7-1-11-2c-3 1-7 2-10 2h-68c-9 0-18-4-24-11c-5-8-7-17-5-26l14-63C54 572 0 484 0 383c0-136 100-250 230-271l14-86c2-15 15-26 30-26s27 11 30 26l13 86c132 20 233 134 233 271m-147-80H147l46 81h165zM286 109l-12-79l-12 79l-7 43h38zm27 106l23-42H214l23 42l-9 68h95l-9-68zM150 690h68l3-37l21-246h-28l-51 227zm92-34l-3 34h73l-3-34l-22-249h-24zm91 34h67l-12-56l-52-227h-27l21 246z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yelp(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M386 217c-1-24-5-47-7-70c-2-25 1-52-7-73c-11-30-46-40-74-53c-27-13-40-23-77-20c-22 2-53 11-80 20c-38 13-77 27-88 56c-12 30 4 51 16 70c37 57 75 115 111 170c7 9 17 16 16 28c-22-8-40-14-60-21c-19-7-38-20-60-21c-26-2-50 13-61 32c-16 30-20 116-7 154c18 49 83 49 123 78c-10 14-24 25-27 41c-8 42 25 68 50 80c24 11 60 23 94 35s71 31 103 27c42-6 51-51 49-101c37 10 67 44 109 49c45 6 72-33 90-55c19-22 41-51 45-80c5-44-19-59-45-71c-21-9-42-18-62-26c-10-4-23-7-28-16c61-19 137-20 127-92c-2-22-19-51-35-73c-26-37-55-58-96-76c-25-12-54-30-85-18c-14 5-19 20-34 26M149 56c29-10 51-18 85-19c9-1 27-1 33 10c5 8 5 27 6 39c6 68 10 126 16 201c2 31-1 56-23 58c-26 2-42-33-55-53c-38-57-68-102-107-160c-9-13-19-23-18-35c2-23 40-33 63-41m353 216c16 24 29 42 30 73c-15 18-50 23-84 32c-13 4-30 7-46 11c-25 6-61 20-67-13c-6-26 28-55 44-78c20-28 30-51 52-71c24-22 62 35 71 46M152 469c-34 12-55 21-92 25c-27-10-29-118-11-145c4-7 17-13 23-11c36 11 82 31 123 46c21 9 54 15 39 50c-7 15-65 30-82 35m192 8c1-13 11-23 25-25c0 0 1-3 3-2c34 9 87 26 128 39c13 4 32 7 37 18c12 27-20 59-32 74c-10 14-33 51-60 48c-21-2-46-54-57-71c-8-13-18-29-25-39c-8-12-21-26-19-42m-124 39c9-11 20-29 35-32c10-3 27 0 34 10c9 16 4 66 5 91c1 28 4 78-5 88c-16 19-61 1-84-7c-24-8-61-18-64-40c-2-17 11-26 24-41c19-25 35-47 55-69"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Youtube(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M98 206v125h52V206L220 0h-55l-41 134L83 0H26zm140 8c0 46 4 63 10 81c11 29 31 43 62 43s52-14 63-43c6-18 10-35 10-81s-4-71-10-89c-11-29-32-43-63-43s-51 14-62 43c-6 19-10 43-10 89m316 117V89h-48v145c0 15-1 27-3 33c-3 12-10 17-19 17s-15-5-18-15c-1-6-2-17-2-31V89h-46v184c0 43 15 65 46 65c18 0 34-11 43-34v27zM334 209v18c0 29 0 29-1 33c-3 20-11 31-23 31c-13 0-20-12-23-36v-53c0-25 0-41 1-46c3-18 10-26 22-26s20 11 23 30c1 6 1 22 1 49M102 365h405c56 0 102 47 102 103v197c0 56-46 103-102 103H102C46 768 0 721 0 665V468c0-56 46-103 102-103m36 108h42v-42H59v42h42v214h37zm218 214v-21c6 16 17 26 31 26c17 0 29-12 34-37c3-12 5-33 5-63c0-31-2-53-4-66c-5-25-18-37-35-37c-13 0-23 8-29 22v-80h-33v256zm126-90h68v-15c0-36-4-59-10-71c-8-14-22-22-40-22c-19 0-33 8-41 21c-5 8-9 20-10 36c-1 7-2 21-2 39c0 23 1 38 2 46c1 18 5 33 11 43c8 12 22 18 39 18s30-5 38-15s12-26 12-49v-8h-33c0 26-5 38-16 38c-8 0-14-6-16-18c-2-6-2-17-2-31zm-221 90h34V494h-35v123c0 11 0 19-2 24c-2 8-6 12-13 12c-6 0-11-3-13-11c-2-4-2-12-2-22V494h-32v151c0 31 9 47 32 47c13 0 24-9 31-25zm221-126v-3c-1-11 0-21 3-27c3-5 8-8 14-8c7 0 12 4 14 11c2 5 2 13 2 24v8h-33zm-106-35c8 0 13 10 15 27v38s0 35-1 38c-1 17-6 26-14 26c-9 0-16-7-18-21c-1-5-1-17-1-35c0-25 0-39 1-43c2-19 8-30 18-30"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Z(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28 128h492L142 594h352v60H0l378-466H28z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZupperCase(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M34 0h469L124 663h365v72H0L379 71H34z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Zero(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M247 0c205 0 255 224 255 386s-50 386-255 386S0 546 0 386S41 0 247 0m3 706c157 0 183-201 183-319c0-117-26-318-183-318C88 69 70 269 70 387s18 319 180 319"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Zoomin(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m541 517l176 176l-72 73l-184-185c-45 25-97 39-151 39C139 620 0 481 0 310S139 0 310 0s311 139 311 310c0 80-31 153-80 207M86 310c0 124 100 223 224 223c123 0 224-99 224-223S433 87 310 87C186 87 86 186 86 310m260-35h124v71H346v123h-71V346H152v-71h123V151h71z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Zoomout(children ...ElementRenderer) *LsIcon {
	return &LsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m541 517l176 176l-72 73l-184-185c-45 25-97 39-151 39C139 620 0 481 0 310S139 0 310 0s311 139 311 310c0 80-31 153-80 207M86 310c0 124 100 223 224 223c123 0 224-99 224-223S433 87 310 87C186 87 86 186 86 310m56-35h337v71H142z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
