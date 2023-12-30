package zmdi

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = ""
	hAttr          = "1em"
	viewbox        = "0 0 432 384"
	fill           = "currentColor"
)

type ZmdiIcon struct {
	*SVGSVGElement
}

func Account(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M170.5 192q-35.5 0-60.5-25t-25-60.5T110 46t60.5-25T231 46t25 60.5t-25 60.5t-60.5 25m0 43q31.5 0 69.5 9t69.5 29.5T341 320v43H0v-43q0-26 31.5-46.5T101 244t69.5-9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccountAdd(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M298.5 192q-35.5 0-60.5-25t-25-60.5T238 46t60.5-25T359 46t25 60.5t-25 60.5t-60.5 25M107 149h64v43h-64v64H64v-64H0v-43h64V85h43zm191.5 86q31.5 0 69.5 9t69.5 29.5T469 320v43H128v-43q0-26 31.5-46.5T229 244t69.5-9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccountBox(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 43q0-18 12.5-30.5T43 0h298q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341zm256 85q0-27-18.5-45.5T192 64t-45.5 18.5T128 128t18.5 45.5T192 192t45.5-18.5T256 128M64 299v21h256v-21q0-20-23.5-36T244 240t-52-7t-52 7t-52.5 23T64 299"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccountBoxMail(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m448 107l-64 42l-64-42V85l64 43l64-43zM469 0q18 0 30.5 12.5T512 43v298q0 18-12.5 30.5T469 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zM170.5 64q-26.5 0-45 18.5T107 128t18.5 45.5t45 18.5t45.5-18.5t19-45.5t-19-45.5T170.5 64M299 320v-21q0-20-24-36t-52.5-23t-52-7t-52 7t-52 23T43 299v21zm170-128V64H299v128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccountBoxO(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 197q-20 0-34-14t-14-34t14-34t34-14t34 14t14 34t-14 34t-34 14m96 86v16H96v-16q0-22 33-35t63-13t63 13t33 35M341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zm0 341V43H43v298z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccountBoxPhone(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 0q18 0 30.5 12.5T512 43v298q0 18-12.5 30.5T469 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zM170.5 64q-26.5 0-45 18.5T107 128t18.5 45.5t45 18.5t45.5-18.5t19-45.5t-19-45.5T170.5 64M299 320v-21q0-20-24-36t-52.5-23t-52-7t-52 7t-52 23T43 299v21zm82-85q-8-22-8-43t8-43h35l32-42l-42-43q-44 33-59 85q-6 22-6 43t6 43q15 52 59 85l42-43l-32-42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccountCalendar(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 48q18 0 30.5 12.5T384 91v298q0 18-12.5 30.5T341 432H43q-18 0-30.5-12.5T0 389V91q0-18 12.5-30.5T43 48h21V5h43v43h170V5h43v43zm-149 64q-27 0-45.5 18.5T128 176t18.5 45.5T192 240t45.5-18.5T256 176t-18.5-45.5T192 112m128 256v-21q0-20-23.5-36T244 288t-52-7t-52 7t-52.5 23T64 347v21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccountCircle(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3m0 64Q187 67 168 85.5t-19 45t19 45.5t45.5 19t45-19t18.5-45.5t-18.5-45t-45-18.5m0 303q39.5 0 73-18.5T341 301q0-20-23.5-35.5t-52-23t-52-7.5t-52 7.5t-52 23T85 301q21 32 55 50.5t73.5 18.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccountO(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M170.5 62Q152 62 139 75t-13 31.5t13 31.5t31.5 13t31.5-13t13-31.5T202 75t-31.5-13m.5 192q-44 0-87 16.5T41 299v23h260v-23q0-12-43-28.5T171 254m-.5-233Q206 21 231 46t25 60.5t-25 60.5t-60.5 25t-60.5-25t-25-60.5T110 46t60.5-25m0 192q31.5 0 69.5 9t69.5 29.5T341 299v64H0v-64q0-27 31.5-47.5T101 222t69.5-9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Accounts(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M320 171q-27 0-45.5-19T256 106.5t18.5-45T320 43t45.5 18.5t18.5 45t-18.5 45.5t-45.5 19m-170.5 0q-26.5 0-45.5-19t-19-45.5t19-45T149.5 43t45 18.5t18.5 45t-18.5 45.5t-45 19m0 42q27.5 0 60.5 8t61 26t28 41v53H0v-53q0-23 27.5-41t61-26t61-8m170.5 0q28 0 61 8t60.5 26t27.5 41v53H341v-53q0-43-42-74q13-1 21-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccountsAdd(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M171 149v43h-64v64H64v-64H0v-43h64V85h43v64zm213 22q-10 0-19-3q19-28 19-61q0-34-19-61q9-3 19-3q27 0 45.5 18.5t18.5 45t-18.5 45.5t-45.5 19m-106.5 0q-26.5 0-45.5-19t-19-45.5t19-45T277.5 43t45 18.5t18.5 45t-18.5 45.5t-45 19M419 217q37 6 65 22t28 38v43h-64v-43q0-34-29-60m-142-4q40 0 84 18t44 46v43H149v-43q0-28 44-46t84-18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccountsAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M309 192q-22 0-37.5-15.5T256 139t15.5-38T309 85t37.5 16t15.5 38t-15.5 37.5T309 192m-159.5-21q-26.5 0-45.5-19t-19-45.5t19-45T149.5 43t45 18.5t18.5 45t-18.5 45.5t-45 19m160 64q36.5 0 77 16t40.5 42v48H192v-48q0-26 40.5-42t77-16M149 213q22 0 51 6q-51 28-51 74v48H0v-53q0-23 27.5-41t61-26t60.5-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccountsList(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 43h42v298h-42zm-86 298V43h43v298zM277 43q9 0 15.5 6t6.5 15v256q0 9-6.5 15t-15.5 6H21q-8 0-14.5-6T0 320V64q0-9 6.5-15T21 43zm-128 58q-20 0-34 14t-14 34t14 34t34 14t34-14t14-34t-14-34t-34-14m96 198v-16q0-22-33-35t-63-13t-63 13t-33 35v16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccountsListAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 0v43H43V0zM43 512v-43h341v43zM384 85q18 0 30.5 12.5T427 128v256q0 18-12.5 30.5T384 427H43q-18 0-30.5-12.5T0 384V128q0-18 12.5-30.5T43 85zm-171 59q-20 0-34 14t-14 34t14 34t34 14t34-14t14-34t-14-34t-34-14m107 219v-32q0-24-36.5-39t-70-15t-70 15t-36.5 39v32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccountsOutline(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M331 213q25 0 56 7.5t56.5 24T469 283v58H0v-58q0-22 25.5-38.5t56.5-24t57-7.5q50 0 96 22q46-22 96-22m-86 96v-26q0-10-35-24t-71.5-14T67 259t-35 24v26zm192 0v-26q0-10-35-24t-71-14q-32 0-65 12q11 12 11 26v26zM139 192q-31 0-53-22t-22-53t22-52.5T139 43t52.5 21.5T213 117t-21.5 53t-52.5 22m-.5-117q-17.5 0-30 12.5t-12.5 30t12.5 30t30 12.5t30-12.5t12.5-30t-12.5-30t-30-12.5M331 192q-31 0-53-22t-22-53t22-52.5T331 43t52.5 21.5T405 117t-21.5 53t-52.5 22m-.5-117q-17.5 0-30 12.5t-12.5 30t12.5 30t30 12.5t30-12.5t12.5-30t-12.5-30t-30-12.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirlineSeatFlat(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 171v42H149V85h192q36 0 61 25t25 61M0 235h427v42H299v43H128v-43H0zm109.5-41q-18.5 19-45 19.5T19 195T0 150t18.5-45t45-19.5t45.5 18t19 45t-18.5 45.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirlineSeatFlatAngled(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m443 241l-15 40l-264-95l45-121l182 66q34 12 49 44t3 66M0 195l15-40l405 146l-14 40l-97-34v34H139v-96zm124-41.5q-24 11.5-49 3T38.5 124t-3-49T68 38.5t49-3T153.5 68t3 49t-32.5 36.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirlineSeatIndividualSuite(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 213q-26 0-45-18.5t-19-45T83 104t45-19t45 19t19 45.5t-19 45t-45 18.5M384 85q35 0 60 25t25 61v128H0V85h43v150h170V85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirlineSeatLegroomExtra(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 192q0 27 18.5 45.5T107 256h128v43H107q-44 0-75.5-31.5T0 192V0h43zm401 112q7 12 2.5 25T429 348l-79 36l-73-149H128q-27 0-45.5-19T64 171V0h128v128h75q26 0 38 24l73 149l23-11q12-5 24.5-1.5T444 304"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirlineSeatLegroomNormal(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 192q0 27 18.5 45.5T107 256h128v43H107q-44 0-75.5-31.5T0 192V0h43zm330 128q14 0 23 9.5t9 22.5t-9 22.5t-23 9.5h-96V235H128q-26 0-45-19t-19-45V0h128v128h107q17 0 29.5 12.5T341 171v149z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirlineSeatLegroomReduced(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M362 346q3 15-6.5 26.5T331 384h-96v-64l21-85H128q-26 0-45-19t-19-45V0h128v128h107q17 0 29.5 12.5T341 171l-42 149h30q12 0 21.5 7t11.5 19M43 192q0 27 18.5 45.5T107 256h85v43h-85q-44 0-75.5-31.5T0 192V0h43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirlineSeatReclineExtra(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M71.5 80Q57 70 54 52.5t7-32T88.5 3t32 7.5T138 38t-7 31.5T103.5 87t-32-7M299 365v43H148q-39 0-69-25.5T42 318L0 109h43l42 202q3 24 21 39t42 15zm5-85l123 96l-32 32l-82-64H167q-23 0-40.5-14.5T104 292L75 166q-3-20 8.5-36.5T115 110q10-2 21 1q10 3 16 8l35 27q47 37 100 27v46q-48 8-110-26l22 87z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirlineSeatReclineNormal(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M76.5 75.5Q64 63 64 45.5t12.5-30t30-12.5t30 12.5t12.5 30t-12.5 30t-30 12.5t-30-12.5M43 301q0 27 18.5 45.5T107 365h128v43H107q-44 0-75.5-31.5T0 301V109h43zm298 87l-30 31l-75-75H128q-27 0-45.5-18.5T64 280V157q0-20 14-34t34-14h1q10 0 20 5q9 4 15 11l30 33q17 19 45 31.5t54 11.5v47q-29 0-61-13.5T160 201v79h73z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Airplane(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m405 301l-170-53v117l42 32v32l-74-21l-75 21v-32l43-32V248L0 301v-42l171-107V35q0-14 9-23t22.5-9t23 9t9.5 23v117l170 107z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirplaneOff(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m235 152l170 107v42l-67-21l-167-167V35q0-14 9-23t22.5-9t23 9t9.5 23zM21 72l27-27l336 336l-27 27l-122-122v79l42 32v32l-74-21l-75 21v-32l43-32V248L0 301v-42l128-80z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Airplay(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m107 405l128-128l128 128zM427 0q17 0 29.5 12.5T469 43v256q0 17-12.5 29.5T427 341h-86v-42h86V43H43v256h85v42H43q-18 0-30.5-12.5T0 299V43q0-18 12.5-30.5T43 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alarm(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m427 82l-28 33l-98-83l28-32zM125 32l-97 82L0 82L98 0zm99 99v112l85 50l-16 26l-101-60V131zm-10.5-86Q293 45 349 101.5t56 136T349 373t-135.5 56t-136-56T21 237.5t56.5-136t136-56.5m-.5 342q62 0 106-44t44-106t-44-105.5T213 88t-105.5 43.5T64 237t43.5 106T213 387"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlarmCheck(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m427 82l-28 33l-98-83l28-32zM125 32l-97 82L0 82L98 0zm88.5 13Q293 45 349 101.5t56 136T349 373t-135.5 56t-136-56T21 237.5t56.5-136t136-56.5m-.5 342q62 0 106-44t44-106t-44-105.5T213 88t-105.5 43.5T64 237t43.5 106T213 387m-31-117l105-106l23 23l-128 128l-68-68l23-22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlarmOff(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M221 88q-26 0-51 9l-33-32q40-20 84-20q79 0 135.5 56.5T413 237q0 44-20 84l-32-32q9-26 9-52q0-62-43.5-105.5T221 88m213-6l-27 33l-99-83l28-32zM27 9l21 21l372 372l-27 27l-47-47q-54 47-125 47q-80 0-136-56T29 237q0-71 47-125L59 95l-24 20L5 84l23-19L0 36zm289 343L106 142q-35 42-35 95q0 62 44 106t106 44q54 0 95-35M136 30l-18 15l-31-30l19-15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlarmPlus(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m125 32l-97 82L0 82L98 0zm302 50l-28 33l-98-83l28-32zM213.5 45Q293 45 349 101.5t56 136T349 373t-135.5 56t-136-56T21 237.5t56.5-136t136-56.5m-.5 342q62 0 106-44t44-106t-44-105.5T213 88t-105.5 43.5T64 237t43.5 106T213 387m22-235v64h64v43h-64v64h-43v-64h-64v-43h64v-64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlarmSnooze(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m125 32l-97 82L0 82L98 0zm302 50l-28 33l-98-83l28-32zM213.5 45Q293 45 349 101.5t56 136T349 373t-135.5 56t-136-56T21 237.5t56.5-136t136-56.5m-.5 342q62 0 106-44t44-106t-44-105.5T213 88t-105.5 43.5T64 237t43.5 106T213 387m-64-192v-43h128v38l-77 90h77v43H149v-39l78-89z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Album(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3m-.5 309q40 0 68-28t28-68t-28-68t-68-28t-68 28t-28 68t28 68t68 28m.5-117q8.5 0 15 6t6.5 15t-6.5 15t-15 6t-15-6t-6.5-15t6.5-15t15-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertCircle(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3M235 323v-43h-43v43zm0-86V109h-43v128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertCircleO(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 280h43v43h-43zm0-171h43v128h-43zM213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3m0 384q70.5 0 120.5-50t50-121t-50-121t-120.5-50T93 95T43 216t50 121t120.5 50"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertOctagon(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m272 0l112 112v160L272 384H112L0 272V112L112 0zm-80 305q12 0 20-8t8-19.5t-8-19.5t-20-8t-20 8t-8 19.5t8 19.5t20 8m21-92V85h-42v128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertPolygon(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m469 224l-52 59l8 79l-77 17l-41 68l-72-31l-73 31l-40-67l-77-18l7-79l-52-59l52-60l-7-78l77-17l40-68l73 31l72-31l41 68l77 17l-8 79zM256 331v-43h-43v43zm0-86V117h-43v128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertTriangle(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 408L235 3l234 405zm256-64v-43h-43v43zm0-85v-86h-43v86z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Amazon(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M379 297q6-2 9 2.5t-2 8.5q-34 25-81 39t-92 14Q91 361 2 280q-3-3-1-5.5t6-.5q96 56 211 56q83 0 161-33m46-26q5 6-2.5 31.5T399 342q-3 3-5 2t-1-5q18-45 12-53t-54-2q-4 0-4.5-2t2.5-4q18-13 46-13.5t30 6.5M237 113v-6q0-22-6-30q-7-11-23-11q-28 0-33 25q-2 8-8 8l-40-4q-8-2-6-9q6-34 32.5-49T214 22q41 0 63 21q3 3 5.5 6t4.5 7.5t3.5 7t2 8t1.5 8t1 9V180q0 17 16 38q5 7 0 12q-16 12-32 27q-5 4-10 1q-11-9-24-28q-17 18-32 24.5t-37 6.5q-27 0-44.5-17T114 196q0-49 44-69q17-7 79-14m-8 87q8-14 8-45v-9q-62 0-62 42q0 14 6.5 22.5T200 219q18 0 29-19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Android(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M85 392V179h256v213q0 9-6 15t-15 6h-21v75q0 13-9.5 22.5t-23 9.5t-22.5-9.5t-9-22.5v-75h-43v75q0 13-9.5 22.5T160 520t-22.5-9.5T128 488v-75h-21q-9 0-15.5-6T85 392M32 179q13 0 22.5 9t9.5 23v149q0 13-9.5 22.5T32 392t-22.5-9.5T0 360V211q0-14 9.5-23t22.5-9m362.5 0q13.5 0 23 9t9.5 23v149q0 13-9.5 22.5t-23 9.5t-22.5-9.5t-9-22.5V211q0-14 9-23t22.5-9M289 54q52 38 52 103H85q0-64 53-103l-28-28q-8-7-.5-14.5T125 11l32 32q26-14 56-14t57 14l31-32q8-7 15.5.5T316 26zm-118 61V93h-22v22zm106 0V93h-21v22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AndroidAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 317v-85h299v85q0 62-44 106t-106 44t-105.5-44T0 317M237 69q29 21 45.5 52.5T299 189v22H0v-22q0-36 16.5-67.5T61 69L17 24L34 7l49 49q32-16 66-16t66 16l50-49l17 17zM85.5 168q8.5 0 15-6.5t6.5-15t-6.5-15t-15-6.5t-15 6.5t-6.5 15t6.5 15t15 6.5m128 0q8.5 0 15-6.5t6.5-15t-6.5-15t-15-6.5t-15 6.5t-6.5 15t6.5 15t15 6.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Apple(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M353 146q-21 7-35 32.5T304 229q0 31 16 57.5t43 33.5q-8 27-26.5 55.5T299 418q-16 11-40 11q-16 0-37-8q-18-9-31-9q-10 0-40 12q-18 5-26 5q-24 0-49-20q-36-34-56-81T0 230q0-53 30.5-93.5T108 96q26 0 48 11q17 11 34 11q16 0 31-6q39-16 52-16q35 0 61 23q12 12 19 27M179 99q0-32 25-63q25-27 61-33q0 38-24 67q-27 29-62 29"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Apps(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 107V21h85v86zm128 256v-86h85v86zM0 363v-86h85v86zm0-128v-86h85v86zm128 0v-86h85v86zM256 21h85v86h-85zm-128 86V21h85v86zm128 128v-86h85v86zm0 128v-86h85v86z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Archive(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M374 48q10 11 10 27v266q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V75q0-16 10-27l29-36Q49 0 64 0h256q15 0 25 12zM192 309l117-117h-74v-43h-86v43H75zM45 43h294l-20-22H63z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 171v42H82l119 120l-30 30L0 192L171 21l30 30L82 171z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftBottom(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M320 51L73 299h140v42H0V128h43v141L290 21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowMerge(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m243 371l-72-72l30-30l72 72zM41 107l96-96l96 96h-75v136L30 371L0 341l115-115V107z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowMissed(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m354 85l30 30l-192 192L43 158v98H0V85h171v43H73l119 119z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m171 21l170 171l-170 171l-30-30l119-120H0v-42h260L141 51z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightTop(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M107 43h213v213h-43V115L30 363L0 333L247 85H107z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSplit(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213 21h128v128l-49-49l-61 62l-30-30l61-62zm-85 0L79 70l113 113v180h-43V201L49 100L0 149V21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arrows(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 176v-64h-64L235 5l106 107h-64v64zm-21 21v86h-64v64L0 240l107-107v64zm298 43L363 347v-64h-64v-86h64v-64zm-192 64v64h64L235 475L128 368h64v-64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AspectRatio(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 149v43h-43v-43zm0 86v42h-43v-42zm-171-86v43H85v-43zm85 0v43h-42v-43zM384 21q18 0 30.5 12.5T427 64v256q0 18-12.5 30.5T384 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21zm0 299V64H43v256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AspectRatioAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 192v107H277v-43h64v-64zm-256-64v64H85V85h107v43zM427 0q17 0 29.5 12.5T469 43v298q0 18-12.5 30.5T427 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zm0 342V42H43v300z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Assignment(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 45q18 0 30.5 12.5T384 88v299q0 17-12.5 29.5T341 429H43q-18 0-30.5-12.5T0 387V88q0-18 12.5-30.5T43 45h89q7-19 23.5-30.5T192 3t36.5 11.5T252 45zm-149 0q-9 0-15 6.5t-6 15t6 15t15 6.5t15-6.5t6-15t-6-15t-15-6.5m43 299v-43H85v43zm64-85v-43H85v43zm0-86v-42H85v42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AssignmentAccount(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 48q18 0 30.5 12.5T384 91v298q0 18-12.5 30.5T341 432H43q-18 0-30.5-12.5T0 389V91q0-18 12.5-30.5T43 48h89q7-19 23.5-31T192 5t36.5 12T252 48zm-149 0q-9 0-15 6.5t-6 15t6 15t15 6.5t15-6.5t6-15t-6-15t-15-6.5m0 85q-27 0-45.5 19T128 197.5t18.5 45T192 261t45.5-18.5t18.5-45t-18.5-45.5t-45.5-19m128 256v-30q0-19-23.5-35T244 300.5t-52-7.5t-52 7.5T87.5 324T64 359v30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AssignmentAlert(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 48q18 0 30.5 12.5T384 91v298q0 18-12.5 30.5T341 432H43q-18 0-30.5-12.5T0 389V91q0-18 12.5-30.5T43 48h89q7-19 23.5-31T192 5t36.5 12T252 48zM213 368v-43h-42v43zm0-85V155h-42v128zM192 91q9 0 15-6.5t6-15t-6-15t-15-6.5t-15 6.5t-6 15t6 15t15 6.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AssignmentCheck(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 48q18 0 30.5 12.5T384 91v298q0 18-12.5 30.5T341 432H43q-18 0-30.5-12.5T0 389V91q0-18 12.5-30.5T43 48h89q7-19 23.5-31T192 5t36.5 12T252 48zm-149 0q-9 0-15 6.5t-6 15t6 15t15 6.5t15-6.5t6-15t-6-15t-15-6.5m-43 299l171-171l-30-30l-141 140l-55-55l-30 30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AssignmentO(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 43q18 0 30.5 12.5T384 85v342q0 17-12.5 29.5T341 469H43q-18 0-30.5-12.5T0 427V85q0-17 12.5-29.5T43 43h89q7-19 23.5-31T192 0t36.5 12T252 43zm-149 0q-9 0-15 6t-6 15t6 15t15 6t15-6t6-15t-6-15t-15-6m149 384V85h-42v64H85V85H43v342z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AssignmentReturn(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 48q18 0 30.5 12.5T384 91v298q0 18-12.5 30.5T341 432H43q-18 0-30.5-12.5T0 389V91q0-18 12.5-30.5T43 48h89q7-19 23.5-31T192 5t36.5 12T252 48zm-149 0q-9 0-15 6.5t-6 15t6 15t15 6.5t15-6.5t6-15t-6-15t-15-6.5m85 256v-85h-85v-64L85 261l107 107v-64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AssignmentReturned(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 48q18 0 30.5 12.5T384 91v298q0 18-12.5 30.5T341 432H43q-18 0-30.5-12.5T0 389V91q0-18 12.5-30.5T43 48h89q7-19 23.5-31T192 5t36.5 12T252 48zm-149 0q-9 0-15 6.5t-6 15t6 15t15 6.5t15-6.5t6-15t-6-15t-15-6.5m0 320l107-107h-64v-85h-86v85H85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Attachment(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M117 309q-48 0-82.5-34T0 192t34.5-83T117 75h224q36 0 61 25t25 60t-25 60t-61 25H160q-22 0-37.5-15.5T107 192t15.5-37.5T160 139h160v32H160q-9 0-15 6t-6 15t6 15t15 6h181q22 0 38-15.5t16-37.5t-16-37.5t-38-15.5H117q-35 0-60 25t-25 60t25 60t60 25h203v32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AttachmentAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M203 112h32v245q0 49-34.5 83.5t-83 34.5t-83-34.5T0 357V91q0-36 25-61T85.5 5T146 30t25 61v224q0 22-16 37.5T117 368t-37.5-15.5T64 315V112h32v203q0 8 6.5 14.5t15 6.5t15-6.5T139 315V91q0-22-16-38T85 37T47.5 53T32 91v266q0 36 25 61t60.5 25t60.5-25t25-61z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Audio(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 0h149v64h-85v235h-1q-4 36-31 60.5T96 384q-40 0-68-28T0 288t28-68t68-28q15 0 32 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BadgeCheck(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 5q18 0 30.5 12.5T384 48v276q0 23-19 35L192 475L19 359Q0 347 0 324V48q0-18 12.5-30.5T43 5zM149 325l192-192l-30-30l-162 162l-76-76l-30 30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Balance(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 197h64v150H43zm128 0h64v150h-64zM0 453v-64h405v64zm299-256h64v150h-64zM203 5l202 107v43H0v-43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BalanceWallet(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 320v21q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0h298q18 0 30.5 12.5T384 43v21H192q-18 0-30.5 12.5T149 107v170q0 18 12.5 30.5T192 320zm-192-43V107h213v170zm85.5-53q13.5 0 22.5-9.5t9-22.5t-9-22.5t-22.5-9.5t-23 9.5T245 192t9.5 22.5t23 9.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Battery(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M185 45q12 0 20 8.5t8 20.5v327q0 12-8 20t-20 8H28q-11 0-19.5-8T0 401V74q0-12 8.5-20.5T28 45h36V3h85v42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryAlert(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M185 45q12 0 20 8.5t8 20.5v327q0 12-8 20t-20 8H28q-11 0-19.5-8T0 401V74q0-12 8.5-20.5T28 45h36V3h85v42zm-57 299v-43H85v43zm0-85V152H85v107z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFlash(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M185 45q12 0 20 8.5t8 20.5v327q0 12-8 20t-20 8H28q-11 0-19.5-8T0 401V74q0-12 8.5-20.5T28 45h36V3h85v42zM85 387l86-160h-43V109L43 269h42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryUnknown(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M185 45q12 0 20 8.5t8 20.5v327q0 12-8 20t-20 8H28q-11 0-19.5-8T0 401V74q0-12 8.5-20.5T28 45h36V3h85v42zm-58 298v-41H86v41zm29-112q15-15 15-36q0-27-19-45.5T106.5 131t-45 18.5T43 195h32q0-14 9-23t22.5-9t23 9t9.5 22.5t-10 22.5l-20 20q-19 21-19 43h34q0-16 17-34z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Behance(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 103H277V77h107zM208 203.5q12 17.5 12 42.5q0 20-8 35q-7 14-21 23q-12 9-30 14q-14 4-34 4H0V56h124q12 0 34 5q13 3 26 12q11 7 18 20q6 13 6 31q0 20-9.5 33.5T172 179q24 7 36 24.5M55 163h61q17 0 26-6q10-7 10-23q0-9-3.5-15t-9.5-9q-6-4-12-5q-9-2-15-2H55zm107 80q0-20-11-29q-11-8-30-8H55v73h64q7 0 17-2q8-2 13.5-5.5T159 260q3-6 3-17m264-3H289q0 24 13 37q12 11 34 11q15 0 27-8q12-9 14-18h46q-10 35-34 50q-24 16-55 16q-22 0-40-7t-31-21q-13-13-19-32q-7-18-7-40t7-40.5t20-32.5q13-13 30-21q18-8 40-8q24 0 42 9.5t30 25.5q11 15 17 37q5 21 3 42m-52-34q-2-18-12-30q-9-10-29-10q-13 0-21 4.5T298.5 181t-6.5 13q-3 7-3 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bike(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M330.5 85q-17.5 0-30-12.5t-12.5-30t12.5-30t30-12.5t30 12.5t12.5 30t-12.5 30t-30 12.5m-224 139q44.5 0 75.5 31t31 75.5t-31 75.5t-75.5 31T31 406T0 330.5T31 255t75.5-31m0 181q30.5 0 52.5-22t22-52.5t-22-52.5t-52.5-22T54 278t-22 52.5T54 383t52.5 22M230 192l47 49v132h-42V267l-69-60q-12-10-12-30q0-17 12-30l60-60q10-12 30-12q18 0 34 12l41 41q32 32 76 32v43q-64 0-108-45l-17-17zm175.5 32q44.5 0 75.5 31t31 75.5t-31 75.5t-75.5 31t-75.5-31t-31-75.5t31-75.5t75.5-31m0 181q30.5 0 52.5-22t22-52.5t-22-52.5t-52.5-22t-52.5 22t-22 52.5t22 52.5t52.5 22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Block(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3M43 216q0 59 36 105L318 81q-46-36-105-36q-70 0-120 50T43 216m170 171q71 0 121-50t50-121q0-59-36-105L109 351q46 36 104 36"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BlockAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3m-.5 384q58 0 105-36L79 111q-36 47-36 105q0 71 50 121t120 50m135-66q36-47 36-105q0-71-50-121T213 45q-58 0-104 36z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blogger(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M397 165q30 0 30 31v103q0 53-39.5 91.5T295 429H124q-48 0-86-38.5T0 300V138q0-57 39-96t96-39h90q44 0 84.5 39.5T350 128v11q0 11 7.5 18.5T378 165zm-262-51q-10 0-17.5 7.5t-7.5 18t7.5 18T135 165h78q10 0 17-8t7-18t-7-17.5t-17-7.5zm154 204q10 0 17.5-6.5T314 295t-7.5-17t-17.5-7H135q-10 0-17.5 7t-7.5 17t7.5 16.5T135 318z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bluetooth(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m271 124l-92 92l92 92l-122 121h-21V267l-98 98l-30-30l119-119L0 97l30-30l98 98V3h21zM171 84v81l40-41zm40 224l-40-41v81z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BluetoothConnected(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m85 216l-42 43l-43-43l43-43zm229-92l-92 92l92 92l-122 121h-21V267l-98 98l-30-30l119-119L43 97l30-30l98 98V3h21zM213 84v81l40-41zm40 224l-40-41v81zm88-135l43 43l-43 43l-42-43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BluetoothOff(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 84v69l-43-43V3h22l121 121l-64 65l-30-30l34-35zM30 45l311 312l-30 30l-49-49l-91 91h-22V267l-98 98l-30-30l120-119L0 75zm162 303l40-40l-40-41z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BluetoothSearch(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m240 216l49-49q10 24 10 49q0 26-10 50zm113-113q31 51 31 111q0 61-33 113l-25-25q21-41 21-86q0-46-21-86zm-82 21l-92 92l92 92l-122 121h-21V267l-98 98l-30-30l119-119L0 97l30-30l98 98V3h21zM171 84v81l40-41zm40 224l-40-41v81z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BluetoothSetting(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 512v-43h43v43zm-85 0v-43h42v43zm170 0v-43h43v43zm58-390l-92 91l92 92l-122 122h-21V265l-98 98l-30-30l119-120L0 94l30-30l98 98V0h21zM171 82v80l40-40zm40 223l-40-40v80z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blur(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M74.5 229q8.5 0 15 6.5t6.5 15t-6.5 15t-15 6.5t-15-6.5t-6.5-15t6.5-15t15-6.5m0 86q8.5 0 15 6t6.5 15t-6.5 15t-15 6t-15-6t-6.5-15t6.5-15t15-6m0-171q8.5 0 15 6.5t6.5 15t-6.5 15t-15 6.5t-15-6.5t-6.5-15t6.5-15t15-6.5m-64 11Q21 155 21 165.5T10.5 176T0 165.5T10.5 155m64-96q8.5 0 15 6T96 80t-6.5 15t-15 6t-15-6T53 80t6.5-15t15-6m320 117q-10.5 0-10.5-10.5t10.5-10.5t10.5 10.5t-10.5 10.5m-149-75q-8.5 0-15-6T224 80t6.5-15t15-6t15 6t6.5 15t-6.5 15t-15 6m0-74Q235 27 235 16t10.5-11T256 16t-10.5 11m-235 213Q21 240 21 250.5T10.5 261T0 250.5T10.5 240M160 389q11 0 11 11t-11 11t-11-11t11-11m0-362q-11 0-11-11t11-11t11 11t-11 11m0 74q-9 0-15-6t-6-15t6-15t15-6t15 6t6 15t-6 15t-15 6m0 118q13 0 22.5 9t9.5 22.5t-9.5 23T160 283t-22.5-9.5t-9.5-23t9.5-22.5t22.5-9m170.5 10q8.5 0 15 6.5t6.5 15t-6.5 15t-15 6.5t-15-6.5t-6.5-15t6.5-15t15-6.5m0 86q8.5 0 15 6t6.5 15t-6.5 15t-15 6t-15-6t-6.5-15t6.5-15t15-6m0-171q8.5 0 15 6.5t6.5 15t-6.5 15t-15 6.5t-15-6.5t-6.5-15t6.5-15t15-6.5m0-85q8.5 0 15 6t6.5 15t-6.5 15t-15 6t-15-6t-6.5-15t6.5-15t15-6m64 181q10.5 0 10.5 10.5T394.5 261T384 250.5t10.5-10.5m-149 75q8.5 0 15 6t6.5 15t-6.5 15t-15 6t-15-6t-6.5-15t6.5-15t15-6m0 74q10.5 0 10.5 11t-10.5 11t-10.5-11t10.5-11M160 133q13 0 22.5 9.5t9.5 23t-9.5 22.5t-22.5 9t-22.5-9t-9.5-22.5t9.5-23T160 133m0 182q9 0 15 6t6 15t-6 15t-15 6t-15-6t-6-15t6-15t15-6m85.5-96q13.5 0 22.5 9t9 22.5t-9 23t-22.5 9.5t-23-9.5t-9.5-23t9.5-22.5t23-9m0-86q13.5 0 22.5 9.5t9 23t-9 22.5t-22.5 9t-23-9t-9.5-22.5t9.5-23t23-9.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BlurCircular(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M170.5 152q8.5 0 15 6.5t6.5 15t-6.5 15t-15 6.5t-15-6.5t-6.5-15t6.5-15t15-6.5m0 85q8.5 0 15 6.5t6.5 15t-6.5 15t-15 6.5t-15-6.5t-6.5-15t6.5-15t15-6.5m-64-74q10.5 0 10.5 10.5T106.5 184T96 173.5t10.5-10.5m64 149q10.5 0 10.5 10.5T170.5 333T160 322.5t10.5-10.5m-64-64q10.5 0 10.5 10.5T106.5 269T96 258.5t10.5-10.5m64-128q-10.5 0-10.5-10.5T170.5 99t10.5 10.5t-10.5 10.5m85.5 32q9 0 15 6.5t6 15t-6 15t-15 6.5t-15-6.5t-6-15t6-15t15-6.5m0-32q-11 0-11-10.5T256 99t11 10.5t-11 10.5m64 128q11 0 11 10.5T320 269t-11-10.5t11-10.5m0-85q11 0 11 10.5T320 184t-11-10.5t11-10.5M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3m0 384q70.5 0 120.5-50t50-121t-50-121t-120.5-50T93 95T43 216t50 121t120.5 50m42.5-75q11 0 11 10.5T256 333t-11-10.5t11-10.5m0-75q9 0 15 6.5t6 15t-6 15t-15 6.5t-15-6.5t-6-15t6-15t15-6.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BlurLinear(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M42.5 309Q29 309 20 300t-9-22.5t9-23t22.5-9.5t23 9.5t9.5 23t-9.5 22.5t-23 9m85.5-96q-9 0-15-6t-6-15t6-15t15-6t15 6t6 15t-6 15t-15 6m0-85q-9 0-15-6.5t-6-15t6-15t15-6.5t15 6.5t6 15t-6 15t-15 6.5M0 384v-43h384v43zm42.5-245Q29 139 20 129.5t-9-23T20 84t22.5-9t23 9t9.5 22.5t-9.5 23t-23 9.5m0 85Q29 224 20 214.5T11 192t9-22.5t22.5-9.5t23 9.5T75 192t-9.5 22.5t-23 9.5m85.5 75q-9 0-15-6.5t-6-15t6-15t15-6.5t15 6.5t6 15t-6 15t-15 6.5m170.5-11q-10.5 0-10.5-10.5t10.5-10.5t10.5 10.5t-10.5 10.5M0 0h384v43H0zm298.5 117q-10.5 0-10.5-10.5T298.5 96t10.5 10.5t-10.5 10.5m0 86q-10.5 0-10.5-11t10.5-11t10.5 11t-10.5 11m-85-75q-8.5 0-15-6.5t-6.5-15t6.5-15t15-6.5t15 6.5t6.5 15t-6.5 15t-15 6.5m0 85q-8.5 0-15-6t-6.5-15t6.5-15t15-6t15 6t6.5 15t-6.5 15t-15 6m0 86q-8.5 0-15-6.5t-6.5-15t6.5-15t15-6.5t15 6.5t6.5 15t-6.5 15t-15 6.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BlurOff(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M245.5 101q-8.5 0-15-6T224 80t6.5-15t15-6t15 6t6.5 15t-6.5 15t-15 6m-4.5 96q-11-2-18.5-9.5T214 170l-1-5q0-13 9.5-22.5t23-9.5t22.5 9.5t9 23t-9 22.5t-23 9zm4.5-170Q235 27 235 16t10.5-11T256 16t-10.5 11M160 27q-11 0-11-11t11-11t11 11t-11 11m234.5 149q-10.5 0-10.5-10.5t10.5-10.5t10.5 10.5t-10.5 10.5M160 101q-9 0-15-6t-6-15t6-15t15-6t15 6t6 15t-6 15t-15 6m170.5 171q-8.5 0-15-6.5t-6.5-15t6.5-15t15-6.5t15 6.5t6.5 15t-6.5 15t-15 6.5m0-85q-8.5 0-15-6.5t-6.5-15t6.5-15t15-6.5t15 6.5t6.5 15t-6.5 15t-15 6.5m0-86q-8.5 0-15-6T309 80t6.5-15t15-6t15 6t6.5 15t-6.5 15t-15 6m-85 288q10.5 0 10.5 11t-10.5 11t-10.5-11t10.5-11M0 64l27-27l346 347l-27 27l-81-81q2 4 2 6q0 9-6.5 15t-15 6t-15-6t-6.5-15t6.5-15t14.5-6q2 0 6 1l-60-60q-1 11-10 19t-21 8q-13 0-22.5-9.5T128 251q0-12 7.5-21t19.5-11l-60-60q1 4 1 6q0 9-6.5 15.5t-15 6.5t-15-6.5t-6.5-15t6.5-15T75 144l6 1zm160 251q9 0 15 6t6 15t-6 15t-15 6t-15-6t-6-15t6-15t15-6m234.5-75q10.5 0 10.5 10.5T394.5 261T384 250.5t10.5-10.5m-320-11q8.5 0 15 6.5t6.5 15t-6.5 15t-15 6.5t-15-6.5t-6.5-15t6.5-15t15-6.5m-64-74Q21 155 21 165.5T10.5 176T0 165.5T10.5 155M160 389q11 0 11 11t-11 11t-11-11t11-11m-85.5-74q8.5 0 15 6t6.5 15t-6.5 15t-15 6t-15-6t-6.5-15t6.5-15t15-6m-64-75Q21 240 21 250.5T10.5 261T0 250.5T10.5 240"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Boat(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 432h43v43h-43q-44 0-85-21q-41 20-86 20t-85-20q-42 21-85 21H0v-43h43q45 0 85-28q39 27 85.5 27t85.5-27q40 28 85 28M42 389L1 247q-3-8 1-17q4-8 13-10l28-9v-99q0-18 12.5-30.5T85 69h64V5h128v64h64q18 0 30.5 12.5T384 112v99l27 9q9 2 13 10t1 17l-40 142h-1q-48 0-85-42q-38 42-86 42t-85-42q-37 42-85 42zm43-277v85l128-42l128 42v-85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Book(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 3q17 0 29.5 12.5T341 45v342q0 17-12.5 29.5T299 429H43q-18 0-30.5-12.5T0 387V45q0-17 12.5-29.5T43 3zM43 45v171l53-32l53 32V45z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookImage(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 3q17 0 29.5 12.5T341 45v342q0 17-12.5 29.5T299 429H43q-18 0-30.5-12.5T0 387V45q0-17 12.5-29.5T43 3zM43 45v171l53-32l53 32V45zm0 320h256l-83-109l-64 82l-45-55z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0q18 0 30.5 12.5T299 43v341l-150-64L0 384V43q0-18 12.5-30.5T43 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkOutline(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0q18 0 30.5 12.5T299 43v341l-150-64L0 384V43q0-18 12.5-30.5T43 0zm0 320V43H43v277l106-47z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderAll(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h384v384H0zm171 341V213H43v128zm0-170V43H43v128zm170 170V213H213v128zm0-170V43H213v128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderBottom(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 171v42H85v-42zm85 85v43h-42v-43zM128 0v43H85V0zm85 171v42h-42v-42zM43 0v43H0V0zm170 85v43h-42V85zm86 86v42h-43v-42zM213 0v43h-42V0zm86 0v43h-43V0zm42 213v-42h43v42zm0 86v-43h43v43zM43 85v43H0V85zM341 0h43v43h-43zm0 128V85h43v43zM43 171v42H0v-42zM0 384v-43h384v43zm43-128v43H0v-43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderClear(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M85 43V0h43v43zm0 170v-42h43v42zm0 171v-43h43v43zm86-85v-43h42v43zm0 85v-43h42v43zM0 384v-43h43v43zm0-85v-43h43v43zm0-86v-42h43v42zm0-85V85h43v43zm0-85V0h43v43zm171 170v-42h42v42zm170 86v-43h43v43zm0-86v-42h43v42zm0 171v-43h43v43zm0-256V85h43v43zm-170 0V85h42v43zM341 0h43v43h-43zM171 43V0h42v43zm85 341v-43h43v43zm0-171v-42h43v42zm0-170V0h43v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderColor(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M379 149L165 363H85v-80L299 69zm63-63l-42 42l-80-80l42-42q6-6 15-6t15 6l50 50q6 6 6 15t-6 15M0 427h512v85H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderHorizontal(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 384v-43h43v43zM43 85v43H0V85zM0 299v-43h43v43zm85 85v-43h43v43zM43 0v43H0V0zm85 0v43H85V0zm171 0v43h-43V0zm-86 85v43h-42V85zm0-85v43h-42V0zm128 299v-43h43v43zm-170 85v-43h42v43zM0 213v-42h384v42zM341 0h43v43h-43zm0 128V85h43v43zM171 299v-43h42v43zm85 85v-43h43v43zm85 0v-43h43v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderInner(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 384v-43h43v43zm85 0v-43h43v43zM43 85v43H0V85zM0 299v-43h43v43zM128 0v43H85V0zM43 0v43H0V0zm256 0v43h-43V0zm42 128V85h43v43zm0-128h43v43h-43zm-85 384v-43h43v43zM213 0v171h171v42H213v171h-42V213H0v-42h171V0zm128 384v-43h43v43zm0-85v-43h43v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderLeft(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M171 384v-43h42v43zm0-85v-43h42v43zm0-256V0h42v43zm0 85V85h42v43zm0 85v-42h42v42zM85 384v-43h43v43zm0-341V0h43v43zm0 170v-42h43v42zM0 384V0h43v384zm341-256V85h43v43zm-85 256v-43h43v43zm85-85v-43h43v43zm0-299h43v43h-43zm0 213v-42h43v42zm0 171v-43h43v43zm-85-171v-42h43v42zm0-170V0h43v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderOuter(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213 85v43h-42V85zm0 86v42h-42v-42zm86 0v42h-43v-42zM0 0h384v384H0zm341 341V43H43v298zm-128-85v43h-42v-43zm-85-85v42H85v-42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderRight(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M85 384v-43h43v43zM0 43V0h43v43zm85 0V0h43v43zm0 170v-42h43v42zM0 384v-43h43v43zm171 0v-43h42v43zM0 213v-42h43v42zm0 86v-43h43v43zm0-171V85h43v43zm171 171v-43h42v43zm85-86v-42h43v42zM341 0h43v384h-43zm-85 384v-43h43v43zm0-341V0h43v43zm-85 170v-42h42v42zm0-170V0h42v43zm0 85V85h42v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderStyle(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 384v-43h43v43zm85 0v-43h43v43zm-256 0v-43h43v43zm86 0v-43h42v43zm170-85v-43h43v43zm0-86v-42h43v42zM0 0h384v43H43v341H0zm341 128V85h43v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderTop(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M85 384v-43h43v43zm0-171v-42h43v42zm86 0v-42h42v42zm0 171v-43h42v43zM0 299v-43h43v43zm0 85v-43h43v43zm0-171v-42h43v42zm0-85V85h43v43zm171 171v-43h42v43zm170-171V85h43v43zm0 85v-42h43v42zM0 0h384v43H0zm341 299v-43h43v43zm-85 85v-43h43v43zm-85-256V85h42v43zm170 256v-43h43v43zm-85-171v-42h43v42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderVertical(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 128V85h43v43zm0-85V0h43v43zm85 341v-43h43v43zm0-171v-42h43v42zm-85 0v-42h43v42zm0 171v-43h43v43zm0-85v-43h43v43zM85 43V0h43v43zm256 256v-43h43v43zm-170 85V0h42v384zm170 0v-43h43v43zm0-171v-42h43v42zm0-213h43v43h-43zm0 128V85h43v43zm-85-85V0h43v43zm0 341v-43h43v43zm0-171v-42h43v42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrightnessAuto(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m217 262l24-78l25 78zm195-85l71 71l-71 71v100H312l-71 70l-70-70H71V319L0 248l71-71V77h100l70-70l71 70h100zM290 333h41l-68-192h-43l-68 192h40l15-42h68z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrightnessFive(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M412 319v100H312l-71 70l-70-70H71V319L0 248l71-71V77h100l70-70l71 70h100v100l71 71zm-171 57q53 0 90.5-37.5T369 248t-37.5-90.5T241 120t-90.5 37.5T113 248t37.5 90.5T241 376"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrightnessFour(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m412 177l71 71l-71 71v100H312l-71 70l-70-70H71V319L0 248l71-71V77h100l70-70l71 70h100zM241 376q53 0 90.5-37.5T369 248t-37.5-90.5T241 120q-28 0-53 12q33 15 54 46.5t21 69.5t-21 69.5t-54 46.5q25 12 53 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrightnessSetting(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 0q17 0 29.5 12.5T469 43v298q0 18-12.5 30.5T427 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zm0 342V42H43v300zm-278-65v-53l-32-32l32-32v-53h54l32-32l32 32h53v53l32 32l-32 32v53h-53l-32 32l-32-32zm86-149v128q26 0 45-18.5t19-45.5t-19-45.5t-45-18.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrightnessSeven(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m412 177l71 71l-71 71v100H312l-71 70l-70-70H71V319L0 248l71-71V77h100l70-70l71 70h100zM241 376q53 0 90.5-37.5T369 248t-37.5-90.5T241 120t-90.5 37.5T113 248t37.5 90.5T241 376m.5-213q35.5 0 60.5 25t25 60t-25 60t-60.5 25t-60.5-25t-25-60t25-60t60.5-25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrightnessSix(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M412 319v100H312l-71 70l-70-70H71V319L0 248l71-71V77h100l70-70l71 70h100v100l71 71zm-171 57q53 0 90.5-37.5T369 248t-37.5-90.5T241 120z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrightnessThree(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 3q88 0 150.5 62.5T277 216t-62.5 150.5T64 429q-33 0-64-9q66-21 107.5-77T149 216T107.5 89T0 12q31-9 64-9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrightnessTwo(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M107 3q88 0 150.5 62.5T320 216t-62.5 150.5T107 429q-57 0-107-28q49-29 78-78t29-107t-29-107T0 31Q50 3 107 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrokenImage(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 43v140l-64-64l-85 86l-86-86l-85 86l-64-65V43q0-18 12.5-30.5T43 0h298q18 0 30.5 12.5T384 43m-64 137l64 64v97q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V201l64 64l85-86l86 86z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Brush(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M106.5 235q26.5 0 45.5 18.5t19 45.5q0 35-25 60t-61 25q-24 0-47-11.5T0 341q15 0 29-11.5T43 299q0-27 18.5-45.5t45-18.5M399 35q6 6 6 15t-6 15L208 256l-59-59L340 6q7-6 15.5-6T371 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bug(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 107v42h-44q2 13 2 22v21h42v43h-42v21q0 9-2 21h44v43h-60q-17 29-46 46.5T171 384t-64.5-17.5T60 320H0v-43h45q-2-12-2-21v-21H0v-43h43v-21q0-9 2-22H0v-42h60q15-26 39-42L64 30L94 0l47 46q14-3 29.5-3t30.5 3l46-46l30 30l-34 35q24 16 38 42zM213 277v-42h-85v42zm0-85v-43h-85v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bus(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 301V88q0-27 12.5-44.5t38-26t53-11.5t67-3t67 3t53 11.5t38 26T341 88v213q0 28-21 48v38q0 8-6.5 14.5T299 408h-22q-8 0-14.5-6.5T256 387v-22H85v22q0 8-6 14.5T64 408H43q-9 0-15.5-6.5T21 387v-38Q0 329 0 301m74.5 22q13.5 0 23-9.5t9.5-23t-9.5-22.5t-23-9t-22.5 9t-9 22.5t9 23t22.5 9.5m192 0q13.5 0 23-9.5t9.5-23t-9.5-22.5t-23-9t-22.5 9t-9 22.5t9 23t22.5 9.5M299 195V88H43v107z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cake(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 128q-18 0-30.5-12.5T149 85q0-12 7-22l36-63l36 63q7 10 7 22q0 18-12.5 30.5T192 128m98 213q22 22 52 22q23 0 42-13v98q0 9-6.5 15t-14.5 6H21q-8 0-14.5-6T0 448v-98q19 13 42 13q30 0 52-22l23-23l23 23q21 21 52 21t52-21l23-23zm30-149q27 0 45.5 18.5T384 256v33q0 17-12.5 29.5T342 331t-29-12l-46-46l-46 46q-11 11-29 11t-30-11l-45-46l-46 46q-12 12-29 12t-29.5-12.5T0 289v-33q0-27 18.5-45.5T64 192h107v-43h42v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 240v107H192V240zM277 5h43v43h21q18 0 30.5 12.5T384 91v298q0 18-12.5 30.5T341 432H43q-18 0-30.5-12.5T0 389V91q0-18 12.5-30.5T43 48h21V5h43v43h170zm64 384V155H43v234z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 48q18 0 30.5 12.5T384 91v298q0 18-12.5 30.5T341 432H43q-18 0-30.5-12.5T0 389V91q0-18 12.5-30.5T43 48h21V5h43v43h170V5h43v43zm0 341V155H43v234zM85 197h107v107H85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarCheck(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M289 220L162 347l-68-68l23-23l45 45l104-104zm52-172q18 0 30.5 12.5T384 91v298q0 18-12.5 30.5T341 432H43q-18 0-30.5-12.5T0 389V91q0-18 12.5-30.5T43 48h21V5h43v43h170V5h43v43zm0 341V155H43v234z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarClose(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m135 347l-23-23l52-52l-52-52l23-23l52 52l52-52l22 23l-52 52l52 52l-22 23l-52-52zM341 48q18 0 30.5 12.5T384 91v298q0 18-12.5 30.5T341 432H43q-18 0-30.5-12.5T0 389V91q0-18 12.5-30.5T43 48h21V5h43v43h170V5h43v43zm0 341V155H43v234z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarNote(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 197v43H85v-43zm42-149q18 0 30.5 12.5T384 91v298q0 18-12.5 30.5T341 432H43q-18 0-30.5-12.5T0 389V91q0-18 12.5-30.5T43 48h21V5h43v43h170V5h43v43zm0 341V155H43v234zM235 283v42H85v-42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M145 216q0-68 68.5-68t68.5 68t-68.5 68t-68.5-68m4-213h128l39 42h68q18 0 30.5 12.5T427 88v256q0 18-12.5 30.5T384 387H43q-18 0-30.5-12.5T0 344V88q0-18 12.5-30.5T43 45h67zm64 320q44 0 75.5-31.5T320 216t-31.5-75.5T213 109t-75 31.5t-31 75.5t31 75.5t75 31.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraAdd(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M149 0h128l39 43h68q18 0 30.5 12.5T427 85v256q0 18-12.5 30.5T384 384H43q-18 0-30.5-12.5T0 341V85q0-17 12.5-29.5T43 43h67zm64 320q44 0 75.5-31.5T320 213t-31.5-75t-75.5-31t-75 31t-31 75t31 75.5t75 31.5m0-21l-26-59l-59-27l59-26l26-59l27 59l59 26l-59 27z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m158 184l-1 2L79 51Q137 3 213 3q23 0 47 5zm259-32H211l78-135q46 17 79.5 52.5T417 152m5 21q5 22 5 43q0 83-57 144L269 184l-6-11zm-282 43l24 43H4q-4-22-4-43q0-82 56-144zM10 280h206l-78 135q-46-17-79.5-52.5T10 280m240 0l20-34l78 135q-59 48-135 48q-22 0-46-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraBw(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 43q18 0 30.5 12.5T427 85v256q0 18-12.5 30.5T384 384H43q-18 0-30.5-12.5T0 341V85q0-17 12.5-29.5T43 43h68l38-43h128l39 43zm0 298V85H213v22q-44 0-75 31t-31 75.5t31 75.5t75 31v21zm-64-127.5q0 44.5-31 75.5t-76 31v-38q29 0 49-20t20-48.5t-20-48.5t-49-20v-38q45 0 76 31t31 75.5m-175 0q0-28.5 20-48.5t48-20v137q-28 0-48-20t-20-48.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraFront(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M107 427v-43l64 64l-64 64v-43H0v-42zm85 0h107v42H192zm-42.5-256q-17.5 0-30-12.5T107 128t12.5-30.5t30-12.5t30 12.5T192 128t-12.5 30.5t-30 12.5M256 0q18 0 30.5 12.5T299 43v298q0 18-12.5 30.5T256 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zM43 43v224q0-24 36.5-39t70-15t70 15t36.5 39V43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraMic(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 43q18 0 30.5 12.5T427 85v256q0 18-12.5 30.5T384 384H235v-45q45-7 75.5-43t30.5-83h-42q0 36-25 61t-60.5 25t-60.5-25t-25-61H85q0 47 30.5 83t76.5 43v45H43q-18 0-30.5-12.5T0 341V85q0-17 12.5-29.5T43 43h67l39-43h128l39 43zM256 213v-85q0-18-12.5-30.5t-30-12.5t-30 12.5T171 128v85q0 18 12.5 30.5t30 12.5t30-12.5T256 213"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraPartyMode(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 45q18 0 30.5 12.5T427 88v256q0 18-12.5 30.5T384 387H43q-18 0-30.5-12.5T0 344V88q0-18 12.5-30.5T43 45h67l39-42h128l39 42zm-171 64q-44 0-75 31.5T107 216q0 10 2 21h44q-4-10-4-21q0-27 19-45.5t45-18.5h85q-32-43-85-43m0 214q44 0 75.5-31.5T320 216q0-12-2-21h-45q4 10 4 21q0 27-18.5 45.5T213 280h-85q33 43 85 43"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraRear(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M107 427v-43l64 64l-64 64v-43H0v-42zm85 0h107v42H192zM256 0q18 0 30.5 12.5T299 43v298q0 18-12.5 30.5T256 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zM149.5 128q17.5 0 30-12.5t12.5-30t-12.5-30t-30-12.5t-30 12.5t-12.5 30t12.5 30t30 12.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraRoll(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 91h171v320H256q0 17-12.5 29.5T213 453H43q-18 0-30.5-12.5T0 411V91q0-18 12.5-30.5T43 48h21V27q0-9 6.5-15.5T85 5h86q8 0 14.5 6.5T192 27v21h21q18 0 30.5 12.5T256 91m-43 277v-43h-42v43zm0-192v-43h-42v43zm86 192v-43h-43v43zm0-192v-43h-43v43zm85 192v-43h-43v43zm0-192v-43h-43v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraSwitch(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 45q18 0 30.5 12.5T427 88v256q0 18-12.5 30.5T384 387H43q-18 0-30.5-12.5T0 344V88q0-18 12.5-30.5T43 45h67l39-42h128l39 42zM277 291l75-75l-75-75v54H149v-54l-74 75l74 75v-54h128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Car(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m340 64l44 128v171q0 8-6.5 14.5T363 384h-22q-8 0-14.5-6.5T320 363v-22H64v22q0 8-6.5 14.5T43 384H21q-8 0-14.5-6.5T0 363V192L44 64q8-21 31-21h234q23 0 31 21M74.5 277q13.5 0 23-9t9.5-22.5t-9.5-23t-23-9.5t-22.5 9.5t-9 23t9 22.5t22.5 9m235 0q13.5 0 22.5-9t9-22.5t-9-23t-22.5-9.5t-23 9.5t-9.5 23t9.5 22.5t23 9M43 171h298l-32-96H75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarTaxi(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m340 64l44 128v171q0 8-6.5 14.5T363 384h-22q-8 0-14.5-6.5T320 363v-22H64v22q0 8-6.5 14.5T43 384H21q-8 0-14.5-6.5T0 363V192L44 64q8-21 31-21h53V0h128v43h53q23 0 31 21M74.5 277q13.5 0 23-9t9.5-22.5t-9.5-23t-23-9.5t-22.5 9.5t-9 23t9 22.5t22.5 9m235 0q13.5 0 22.5-9t9-22.5t-9-23t-22.5-9.5t-23 9.5t-9.5 23t9.5 22.5t23 9M43 171h298l-32-96H75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarWash(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M298.5 91Q285 91 276 81.5T267 59q0-10 8-24.5T291 11l8-10q32 36 32 58q0 13-9.5 22.5t-23 9.5M192 91q-13 0-22.5-9.5T160 59q0-10 8-24.5T184 11l8-10q32 36 32 58q0 13-9.5 22.5T192 91M85.5 91q-13.5 0-23-9.5T53 59q0-10 8-24.5T77 11l8-10q32 36 32 58q0 13-9 22.5T85.5 91M340 155l44 128v170q0 9-6.5 15.5T363 475h-22q-8 0-14.5-6.5T320 453v-21H64v21q0 9-6.5 15.5T43 475H21q-8 0-14.5-6.5T0 453V283l44-128q8-22 31-22h234q23 0 31 22M74.5 368q13.5 0 23-9.5T107 336t-9.5-22.5t-23-9.5t-22.5 9.5t-9 22.5t9 22.5t22.5 9.5m235 0q13.5 0 22.5-9.5t9-22.5t-9-22.5t-22.5-9.5t-23 9.5T277 336t9.5 22.5t23 9.5M43 261h298l-32-96H75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Card(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 21q18 0 30.5 12.5T427 64v256q0 18-12.5 30.5T384 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21zm0 299V192H43v128zm0-213V64H43v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardAlert(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 3q17 0 29.5 12.5T341 45v342q0 17-12.5 29.5T299 429H43q-18 0-30.5-12.5T0 387V131L128 3zM192 323v-43h-43v43zm0-86V131h-43v106z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardGiftcard(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 88q18 0 30.5 12.5T427 131v234q0 18-12.5 30.5T384 408H43q-18 0-30.5-12.5T0 365V131q0-18 12.5-30.5T43 88h46q-4-11-4-21q0-27 19-45.5T149 3q34 0 54 28l10 15l11-15q19-28 53-28q27 0 45.5 18.5T341 67q0 10-4 21zM277.5 45q-8.5 0-15 6.5t-6.5 15t6.5 15t15 6.5t15-6.5t6.5-15t-6.5-15t-15-6.5m-128 0q-8.5 0-15 6.5t-6.5 15t6.5 15t15 6.5t15-6.5t6.5-15t-6.5-15t-15-6.5M384 365v-42H43v42zm0-106V131H276l44 60l-35 25l-50-69l-22-29l-21 29l-51 69l-34-25l44-60H43v128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardMembership(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 3q18 0 30.5 12.5T427 45v235q0 18-12.5 30.5T384 323h-85v106l-86-42l-85 42V323H43q-18 0-30.5-12.5T0 280V45q0-17 12.5-29.5T43 3zm0 277v-43H43v43zm0-107V45H43v128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardOff(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m354 43l1 249L113 50l50-50h149q17 0 29.5 12.5T354 43M27 19l373 372l-27 28l-40-41q-10 6-21 6H99q-18 0-30.5-12.5T56 341V102L0 46z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardSd(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 3q17 0 29.5 12.5T341 45v342q0 17-12.5 29.5T299 429H43q-18 0-30.5-12.5T0 387V131L128 3zM171 131V45h-43v86zm64 0V45h-43v86zm64 0V45h-43v86z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardSim(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 45v342q0 17-12.5 29.5T299 429H43q-18 0-30.5-12.5T0 387V131L128 3h171q17 0 29.5 12.5T341 45M107 365v-42H64v42zm170 0v-42h-42v42zm-170-85v-85H64v85zm85 85v-85h-43v85zm0-128v-42h-43v42zm85 43v-85h-42v85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardTravel(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 88q18 0 30.5 12.5T427 131v234q0 18-12.5 30.5T384 408H43q-18 0-30.5-12.5T0 365V131q0-18 12.5-30.5T43 88h64V45q0-17 12.5-29.5T149 3h128q18 0 30.5 12.5T320 45v43zM149 45v43h128V45zm235 320v-42H43v42zm0-106V131h-64v42h-43v-42H149v42h-42v-42H43v128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretDown(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 139h213L107 245z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretDownCircle(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3M107 173l106 107l107-107z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretLeft(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M107 85v214L0 192z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretLeftCircle(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3M256 109L149 216l107 107z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretRight(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 299V85l107 107z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretRightCircle(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3M171 109v214l106-107z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretUp(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213 245H0l107-106z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretUpCircle(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3m-.5 145L107 254h213z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Case(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 88q18 0 30.5 12.5T427 131v234q0 18-12.5 30.5T384 408H43q-18 0-30.5-12.5T0 365V131q0-18 12.5-30.5T43 88h85V45q0-17 12.5-29.5T171 3h85q18 0 30.5 12.5T299 45v43zm-128 0V45h-85v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaseCheck(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 91q18 0 30.5 12.5T427 133v235q0 18-12.5 30.5T384 411H43q-18 0-30.5-12.5T0 368V133q0-17 12.5-29.5T43 91h85V48l43-43h85l43 43v43zM171 48v43h85V48zm10 288l141-141l-30-30l-111 111l-44-45l-30 30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaseDownload(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 91q18 0 30.5 12.5T427 133v235q0 18-12.5 30.5T384 411H43q-18 0-30.5-12.5T0 368V133q0-17 12.5-29.5T43 91h85V48l43-43h85l43 43v43zM171 48v43h85V48zm42 320l107-107h-64v-85h-85v85h-64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CasePlay(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 88h128v277q0 18-12.5 30.5T384 408H43q-18 0-30.5-12.5T0 365V88h128V45q0-17 12.5-29.5T171 3h85q18 0 30.5 12.5T299 45zM171 45v43h85V45zm-22 299l160-107l-160-85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cast(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 0q17 0 29.5 12.5T469 43v298q0 18-12.5 30.5T427 384H277v-43h150V43H43v64H0V43q0-18 12.5-30.5T43 0zM0 320q27 0 45.5 18.5T64 384H0zm0-85q62 0 105.5 43.5T149 384h-42q0-44-31.5-75.5T0 277zm0-86q97 0 166 69t69 166h-43q0-80-56-136T0 192z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CastConnected(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 320q27 0 45.5 18.5T64 384H0zm0-85q62 0 105.5 43.5T149 384h-42q0-44-31.5-75.5T0 277zM384 85v214H264q-21-64-68-111T85 120V85zM0 149q97 0 166 69t69 166h-43q0-80-56-136T0 192zM427 0q17 0 29.5 12.5T469 43v298q0 18-12.5 30.5T427 384H277v-43h150V43H43v64H0V43q0-18 12.5-30.5T43 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CenterFocusStrong(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 107q35 0 60 25t25 60t-25 60t-60 25t-60-25t-25-60t25-60t60-25M43 256v85h85v43H43q-18 0-30.5-12.5T0 341v-85zm0-213v85H0V43q0-18 12.5-30.5T43 0h85v43zM341 0q18 0 30.5 12.5T384 43v85h-43V43h-85V0zm0 341v-85h43v85q0 18-12.5 30.5T341 384h-85v-43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CenterFocusWeak(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 256v85h85v43H43q-18 0-30.5-12.5T0 341v-85zm0-213v85H0V43q0-18 12.5-30.5T43 0h85v43zM341 0q18 0 30.5 12.5T384 43v85h-43V43h-85V0zm0 341v-85h43v85q0 18-12.5 30.5T341 384h-85v-43zM192 107q35 0 60 25t25 60t-25 60t-60 25t-60-25t-25-60t25-60t60-25m0 128q18 0 30.5-12.5T235 192t-12.5-30.5T192 149t-30.5 12.5T149 192t12.5 30.5T192 235"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chart(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zM128 299V149H85v150zm85 0V85h-42v214zm86 0v-86h-43v86z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartDonut(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M235 4q81 8 136.5 68.5T427 216q0 45-19 87l-56-33q11-27 11-54q0-56-37-98t-91-50zm-22 361q72 0 117-56l55 33q-30 41-75 64t-97 23q-88 0-150.5-62.5T0 216q0-83 55.5-143.5T192 4v64q-55 8-91.5 50T64 216q0 62 43.5 105.5T213 365"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M119 282L346 55l29 30l-256 256L0 222l30-30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckAll(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M375 85L240 221l-30-30L345 55zm91-30l30 30l-256 256l-119-119l30-30l89 89zM0 222l30-30l119 119l-30 30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircle(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3M171 323l192-192l-30-31l-162 162l-77-76l-30 30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircleU(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213 3q88 0 151 62.5T427 216t-63 150.5T213 429T62.5 366.5T0 216T62.5 65.5T213 3m107 341v-43H107v43zm-143-85l143-143l-30-30l-113 113l-40-41l-30 30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckSquare(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zM149 299l192-192l-30-31l-162 162l-76-76l-30 30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDown(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m226 119l30 30l-128 128L0 149l30-30l98 98z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeft(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m158 94l-98 98l98 98l-30 30L0 192L128 64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRight(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m30 64l128 128L30 320L0 290l98-98L0 94z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUp(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m128 107l128 128l-30 30l-98-98l-98 98l-30-30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circle(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleO(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3m0 384q70.5 0 120.5-50t50-121t-50-121t-120.5-50T93 95T43 216t50 121t120.5 50"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func City(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 195h128v213H0V109h128V67l64-64l64 64zM85 365v-42H43v42zm0-85v-43H43v43zm0-85v-43H43v43zm128 170v-42h-42v42zm0-85v-43h-42v43zm0-85v-43h-42v43zm0-86V67h-42v42zm128 256v-42h-42v42zm0-85v-43h-42v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CityAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213 85h214v299H0V0h213zM85 341v-42H43v42zm0-85v-43H43v43zm0-85v-43H43v43zm0-86V43H43v42zm86 256v-42h-43v42zm0-85v-43h-43v43zm0-85v-43h-43v43zm0-86V43h-43v42zm213 256V128H213v43h43v42h-43v43h43v43h-43v42zm-43-170v42h-42v-42zm0 85v43h-42v-43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Close(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 73L179 192l120 119l-30 30l-120-119L30 341L0 311l119-119L0 73l30-30l119 119L269 43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseCircle(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3M320 293l-77-77l77-77l-30-30l-77 77l-76-77l-30 30l76 77l-76 77l30 30l76-77l77 77z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseCircleO(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m269 131l30 30l-56 55l56 55l-30 30l-56-55l-55 55l-30-30l55-55l-55-55l30-30l55 55zM213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3m0 384q70.5 0 120.5-50t50-121t-50-121t-120.5-50T93 95T43 216t50 121t120.5 50"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClosedCaption(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 21q18 0 30.5 12.5T384 64v256q0 18-12.5 30.5T341 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21zM171 171v-22q0-8-6.5-14.5T149 128H85q-8 0-14.5 6.5T64 149v86q0 8 6.5 14.5T85 256h64q9 0 15.5-6.5T171 235v-22h-32v11H96v-64h43v11zm149 0v-22q0-8-6.5-14.5T299 128h-64q-9 0-15.5 6.5T213 149v86q0 8 6.5 14.5T235 256h64q8 0 14.5-6.5T320 235v-22h-32v11h-43v-64h43v11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M413 150q42 3 70.5 33.5T512 256q0 44-31.5 75.5T405 363H128q-53 0-90.5-37.5T0 235q0-50 33-86t81-41q20-40 58-63.5T256 21q58 0 102 37t55 92"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudBox(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M171 277q-27 0-45.5-18.5T107 213q0-24 16.5-42.5T164 150h3q9-20 27.5-31.5T235 107q28 0 48.5 18t24.5 46h1q22 0 38 15.5t16 37.5t-16 37.5t-38 15.5zM427 0q17 0 29.5 12.5T469 43v298q0 18-12.5 30.5T427 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zm0 342V42H43v300z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudCircle(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3M309 301q22 0 38-15.5t16-37.5t-16-37.5t-38-15.5h-10q0-36-25-61t-61-25q-29 0-52 18.5T131 174l-3-1q-27 0-45.5 19T64 237.5t18.5 45T128 301z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDone(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M413 150q42 3 70.5 33.5T512 256q0 44-31.5 75.5T405 363H128q-53 0-90.5-37.5T0 235q0-50 33-86t81-41q20-40 58-63.5T256 21q58 0 102 37t55 92M213 299l141-141l-30-30l-111 110l-44-44l-30 30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDownload(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M413 150q42 3 70.5 33.5T512 256q0 44-31.5 75.5T405 363H128q-53 0-90.5-37.5T0 235q0-50 33-86t81-41q20-40 58-63.5T256 21q58 0 102 37t55 92m-50 63h-64v-85h-86v85h-64l107 107z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudOff(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M413 150q42 3 70.5 33.5T512 256q0 55-45 87l-31-31q33-19 33-56q0-27-18.5-45.5T405 192h-32v-11q0-48-34-82.5T256 64q-29 0-54 13l-32-31q40-25 86-25q58 0 102 37t55 92M64 48l27-27l357 357l-27 27l-43-42H128q-53 0-90.5-37.5T0 235q0-52 35.5-89t87.5-39zm101 101h-37q-35 0-60 25t-25 60.5T68 295t60 25h208z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudOutline(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M413 150q42 3 70.5 33.5T512 256q0 44-31.5 75.5T405 363H128q-53 0-90.5-37.5T0 235q0-50 33-86t81-41q20-40 58-63.5T256 21q58 0 102 37t55 92m-8 170q27 0 45.5-18.5T469 256t-18.5-45.5T405 192h-32v-11q0-48-34-82.5T256 64q-40 0-71 24t-42 61h-15q-35 0-60 25t-25 60.5T68 295t60 25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudOutlineAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M413 150q42 3 70.5 33.5T512 256q0 44-31.5 75.5T405 363H128q-53 0-90.5-37.5T0 235q0-50 33-86t81-41q21-40 59-63.5T256 21q58 0 102 37t55 92m-8 170q27 0 45.5-19t18.5-45t-18.5-45t-45.5-19h-32v-11q0-48-34.5-82.5T256 64q-58 0-94 47q41 12 67.5 46t26.5 78h-43q0-36-25-61t-60-25t-60 25t-25 60.5T68 295t60 25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUpload(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M413 150q42 3 70.5 33.5T512 256q0 44-31.5 75.5T405 363H128q-53 0-90.5-37.5T0 235q0-50 33-86t81-41q20-40 58-63.5T256 21q58 0 102 37t55 92m-114 63h64L256 107L149 213h64v86h86z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cocktail(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M171 213L0 43V0h384v43L213 213v128h107v43H64v-43h107zM96 85h192l43-42H53z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Code(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m158 290l-30 30L0 192L128 64l30 30l-98 98zm111 0l98-98l-98-98l30-30l128 128l-128 128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeSetting(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M148 80L55 192l93 112l-33 27L0 192L115 53zm-16 133v-42h43v42zm213-42v42h-42v-42zm-128 42v-42h43v42zM362 53l115 139l-115 139l-33-27l93-112l-93-112z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeSmartphone(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 91v42H21V48q0-18 12.5-30.5T64 5l213 1q18 0 30.5 12T320 48v85h-43V91zm179 247l-30-30l68-68l-68-68l30-30l98 98zm-115-30l-30 30l-98-98l98-98l30 30l-68 68zm149 81v-42h43v85q0 18-12.5 30.5T277 475H64q-18 0-30.5-12.5T21 432v-85h43v42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Codepen(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M390 247v-62l-46 31zM232 377l143-96l-64-43l-79 53zm-19-118l65-43l-65-43l-65 43zm-18 118v-86l-80-53l-64 43zM37 185v62l46-31zM195 55L51 151l64 43l80-53zm37 0v86l79 53l64-43zm195 94v135q0 1-1 2v1q-1 0-1 1v1l-1 1v1l-.5.5l-.5.5q0 1-1 1l-1 1l-1 1l-1 1l-195 130q-5 3-10.5 3t-10.5-3L8 296H7v-1l-1-.5l-1-.5v-1H4v-1l-1-1v-1H2v-1l-1-1v-2q-1-1-1-2V148q0-1 1-2v-2l1-1v-1l1-1l.5-.5l.5-.5v-1q1 0 1-.5v-.5h.5l.5-1h1l1-1L203 6q10-7 21 0l195 130l1 1h1v1q1 0 1 1q1 0 1 .5v.5l1 1v1q1 0 1 1v1l1 1v1q1 1 1 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coffee(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 0q18 0 30.5 12.5T427 43v64q0 17-12.5 29.5T384 149h-43v64q0 36-25 61t-60 25H128q-35 0-60-25t-25-61V0zm0 107V43h-43v64zM0 384v-43h384v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollectionBookmark(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 88v299h298v42H43q-18 0-30.5-12.5T0 387V88zM384 3q18 0 30.5 12.5T427 45v256q0 18-12.5 30.5T384 344H128q-18 0-30.5-12.5T85 301V45q0-17 12.5-29.5T128 3zm0 213V45H277v171l54-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollectionCasePlay(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 176v235h341q0 17-12.5 29.5T341 453H43q-18 0-30.5-12.5T0 411V176zm320-85h106v234q0 18-12.5 30.5T427 368H128q-18 0-30.5-12.5T85 325V91h107V48q0-18 12.5-30.5T235 5h85q18 0 30.5 12.5T363 48zM235 48v43h85V48zm0 256l117-85l-117-64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollectionFolderImage(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 88v299h384v42H43q-18 0-30.5-12.5T0 387V88zm426-43q18 0 30.5 12.5T512 88v213q0 18-12.5 30.5T469 344H128q-18 0-30.5-12.5T85 301l1-256q0-17 12-29.5T128 3h128l43 42zM149 280h299l-75-96l-53 64l-75-96z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollectionImage(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 301q0 18-12.5 30.5T384 344H128q-18 0-30.5-12.5T85 301V45q0-17 12.5-29.5T128 3h256q18 0 30.5 12.5T427 45zm-235-85l-64 85h256l-85-106l-64 79zM0 88h43v299h298v42H43q-18 0-30.5-12.5T0 387z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollectionImageO(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m319 204l76 100H160l59-75l41 50zM43 91v341h341v43H43q-18 0-30.5-12.5T0 432V91zM427 5q17 0 29.5 12.5T469 48v299q0 17-12.5 29.5T427 389H128q-18 0-30.5-12.5T85 347V48q0-18 12.5-30.5T128 5zm0 342V48H128v299z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollectionItem(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 91v341h341v43H43q-18 0-30.5-12.5T0 432V91zM427 5q17 0 29.5 12.5T469 48v299q0 17-12.5 29.5T427 389H128q-18 0-30.5-12.5T85 347V48q0-18 12.5-30.5T128 5zm0 342V48H128v299z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollectionItemEight(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 91v341h341v43H43q-18 0-30.5-12.5T0 432V91zM427 5q17 0 29.5 12.5T469 48v299q0 17-12.5 29.5T427 389H128q-18 0-30.5-12.5T85 347V48q0-18 12.5-30.5T128 5zm0 342V48H128v299zm-171-43q-18 0-30.5-12.5T213 261v-32q0-13 9.5-22.5T245 197q-13 0-22.5-9t-9.5-23v-32q0-17 12.5-29.5T256 91h43q17 0 29.5 12.5T341 133v32q0 14-9 23t-23 9q14 0 23 9.5t9 22.5v32q0 18-12.5 30.5T299 304zm0-171v43h43v-43zm0 86v42h43v-42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollectionItemFive(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 5q17 0 29.5 12.5T469 48v299q0 17-12.5 29.5T427 389H128q-18 0-30.5-12.5T85 347V48q0-18 12.5-30.5T128 5zm0 342V48H128v299zM43 91v341h341v43H43q-18 0-30.5-12.5T0 432V91zm298 170q0 18-12.5 30.5T299 304h-86v-43h86v-42h-86V91h128v42h-85v43h43q17 0 29.5 12.5T341 219z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollectionItemFour(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 91v341h341v43H43q-18 0-30.5-12.5T0 432V91zm256 213v-85h-86V91h43v85h43V91h42v213zM427 5q17 0 29.5 12.5T469 48v299q0 17-12.5 29.5T427 389H128q-18 0-30.5-12.5T85 347V48q0-18 12.5-30.5T128 5zm0 342V48H128v299z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollectionItemNine(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 91v341h341v43H43q-18 0-30.5-12.5T0 432V91zM427 5q17 0 29.5 12.5T469 48v299q0 17-12.5 29.5T427 389H128q-18 0-30.5-12.5T85 347V48q0-18 12.5-30.5T128 5zm0 342V48H128v299zM299 91q17 0 29.5 12.5T341 133v128q0 18-12.5 30.5T299 304h-86v-43h86v-42h-43q-18 0-30.5-12.5T213 176v-43q0-17 12.5-29.5T256 91zm0 85v-43h-43v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollectionItemNinePlus(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 91v341h341v43H43q-18 0-30.5-12.5T0 432V91zm234 149q0 18-12.5 30.5T235 283h-64v-43h64v-21h-22q-17 0-29.5-12.5T171 176v-21q0-18 12.5-30.5T213 112h22q17 0 29.5 12.5T277 155zm-64-64h22v-21h-22zM427 5q17 0 29.5 12.5T469 48v299q0 17-12.5 29.5T427 389H128q-18 0-30.5-12.5T85 347V48q0-18 12.5-30.5T128 5zm0 171V48H128v299h299V219h-43v42h-43v-42h-42v-43h42v-43h43v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollectionItemOne(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 91v341h341v43H43q-18 0-30.5-12.5T0 432V91zm234 213V133h-42V91h85v213zM427 5q17 0 29.5 12.5T469 48v299q0 17-12.5 29.5T427 389H128q-18 0-30.5-12.5T85 347V48q0-18 12.5-30.5T128 5zm0 342V48H128v299z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollectionItemSeven(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 91v341h341v43H43q-18 0-30.5-12.5T0 432V91zM427 5q17 0 29.5 12.5T469 48v299q0 17-12.5 29.5T427 389H128q-18 0-30.5-12.5T85 347V48q0-18 12.5-30.5T128 5zm0 342V48H128v299zm-171-43h-43l86-171h-86V91h128v42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollectionItemSix(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 91v341h341v43H43q-18 0-30.5-12.5T0 432V91zM427 5q17 0 29.5 12.5T469 48v299q0 17-12.5 29.5T427 389H128q-18 0-30.5-12.5T85 347V48q0-18 12.5-30.5T128 5zm0 342V48H128v299zm-171-43q-18 0-30.5-12.5T213 261V133q0-17 12.5-29.5T256 91h85v42h-85v43h43q17 0 29.5 12.5T341 219v42q0 18-12.5 30.5T299 304zm0-85v42h43v-42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollectionItemThree(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 5q17 0 29.5 12.5T469 48v299q0 17-12.5 29.5T427 389H128q-18 0-30.5-12.5T85 347V48q0-18 12.5-30.5T128 5zm0 342V48H128v299zM43 91v341h341v43H43q-18 0-30.5-12.5T0 432V91zm298 170q0 18-12.5 30.5T299 304h-86v-43h86v-42h-43v-43h43v-43h-86V91h86q17 0 29.5 12.5T341 133v32q0 14-9 23t-23 9q14 0 23 9.5t9 22.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollectionItemTwo(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 91v341h341v43H43q-18 0-30.5-12.5T0 432V91zM427 5q17 0 29.5 12.5T469 48v299q0 17-12.5 29.5T427 389H128q-18 0-30.5-12.5T85 347V48q0-18 12.5-30.5T128 5zm0 342V48H128v299zm-86-86v43H213v-85q0-18 12.5-30.5T256 176h43v-43h-86V91h86q17 0 29.5 12.5T341 133v43q0 18-12.5 30.5T299 219h-43v42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollectionMusic(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 3q18 0 30.5 12.5T427 45v256q0 18-12.5 30.5T384 344H128q-18 0-30.5-12.5T85 301V45q0-17 12.5-29.5T128 3zm-43 106V67h-85v117q-14-11-32-11q-22 0-37.5 16T171 227t15.5 37.5T224 280t37.5-15.5T277 227V109zM43 88v299h298v42H43q-18 0-30.5-12.5T0 387V88z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollectionPdf(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 3q18 0 30.5 12.5T427 45v256q0 18-12.5 30.5T384 344H128q-18 0-30.5-12.5T85 301V45q0-17 12.5-29.5T128 3zM203 163v-22q0-13-9.5-22.5T171 109h-54v128h32v-42h22q13 0 22.5-9.5T203 163m106 42v-64q0-13-9-22.5t-23-9.5h-53v128h53q14 0 23-9t9-23m86-64v-32h-64v128h32v-42h32v-32h-32v-22zm-246 22v-22h22v22zM43 88v299h298v42H43q-18 0-30.5-12.5T0 387V88zm213 117v-64h21v64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollectionPlus(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 88v299h298v42H43q-18 0-30.5-12.5T0 387V88zM384 3q18 0 30.5 12.5T427 45v256q0 18-12.5 30.5T384 344H128q-18 0-30.5-12.5T85 301V45q0-17 12.5-29.5T128 3zm-21 192v-43h-86V67h-42v85h-86v43h86v85h42v-85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollectionSpeaker(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M303 5q16 0 27 11.5T341 44v307q0 16-11 27t-27 11H124q-16 0-27.5-11T85 351V44q0-16 11.5-27.5T124 5zm-89.5 43q-17.5 0-30 12.5t-12.5 30t12.5 30t30 12.5t30-12.5t12.5-30t-12.5-30t-30-12.5m0 288q35.5 0 60.5-25t25-60.5t-25-60.5t-60.5-25t-60.5 25t-25 60.5t25 60.5t60.5 25M160 250.5q0-53.5 53.5-53.5t53.5 53.5t-53.5 53.5t-53.5-53.5M43 91v341h213v43H43q-18 0-30.5-12.5T0 432V91z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollectionText(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 88v299h298v42H43q-18 0-30.5-12.5T0 387V88zM384 3q18 0 30.5 12.5T427 45v256q0 18-12.5 30.5T384 344H128q-18 0-30.5-12.5T85 301V45q0-17 12.5-29.5T128 3zm-21 192v-43H149v43zm-86 85v-43H149v43zm86-171V67H149v42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollectionVideo(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 88v299h298v42H43q-18 0-30.5-12.5T0 387V88zM384 3q18 0 30.5 12.5T427 45v256q0 18-12.5 30.5T384 344H128q-18 0-30.5-12.5T85 301V45q0-17 12.5-29.5T128 3zM213 269l128-96l-128-96z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comment(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 3q18 0 30.5 12.5T427 45v256q0 18-12.5 30.5T384 344H85L0 429V45q0-17 12.5-29.5T43 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAlert(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 3q18 0 30.5 12.5T427 45v256q0 18-12.5 30.5T384 344H85L0 429V45q0-17 12.5-29.5T43 3zM235 259v-43h-43v43zm0-86V88h-43v85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m426 45l1 384l-86-85H43q-18 0-30.5-12.5T0 301V45q0-17 12.5-29.5T43 3h341q18 0 30 12.5T426 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAltText(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 3q18 0 30.5 12.5T427 45v384l-86-85H43q-18 0-30.5-12.5T0 301V45q0-17 12.5-29.5T43 3zm-43 256v-43H85v43zm0-64v-43H85v43zm0-64V88H85v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentEdit(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 3q18 0 30.5 12.5T427 45v256q0 18-12.5 30.5T384 344H85L0 429V45q0-17 12.5-29.5T43 3zM85 259h53l147-147q8-7 0-15l-38-38q-7-7-15 0L85 206zm256 0v-43H224l-43 43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentImage(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 3q18 0 30.5 12.5T427 45v256q0 18-12.5 30.5T384 344H85L0 429V45q0-17 12.5-29.5T43 3zM64 259h299l-96-128l-75 96l-53-64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentList(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 3q18 0 30.5 12.5T427 45v256q0 18-12.5 30.5T384 344H85L0 429V45q0-17 12.5-29.5T43 3zM128 259v-43H85v43zm0-64v-43H85v43zm0-64V88H85v43zm149 128v-43H171v43zm64-64v-43H171v43zm0-64V88H171v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentMore(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 3q18 0 30.5 12.5T427 45v256q0 18-12.5 30.5T384 344H85L0 429V45q0-17 12.5-29.5T43 3zM149 195v-43h-42v43zm86 0v-43h-43v43zm85 0v-43h-43v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentOutline(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 3q18 0 30.5 12.5T427 45v256q0 18-12.5 30.5T384 344H85L0 429V45q0-17 12.5-29.5T43 3zm0 298V45H43v299l42-43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentText(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 3q18 0 30.5 12.5T427 45v256q0 18-12.5 30.5T384 344H85L0 429V45q0-17 12.5-29.5T43 3zm-43 256v-43H85v43zm0-64v-43H85v43zm0-64V88H85v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentTextAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 3q18 0 30.5 12.5T427 45v256q0 18-12.5 30.5T384 344H85L0 429V45q0-17 12.5-29.5T43 3zM85 152v43h256v-43zm171 107v-43H85v43zm85-128V88H85v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentVideo(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 3q18 0 30.5 12.5T427 45v256q0 18-12.5 30.5T384 344H85L0 429V45q0-17 12.5-29.5T43 3zm-43 256V88l-85 68V88H85v171h171v-69z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comments(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M405 88q9 0 15.5 6.5T427 109v320l-86-85H107q-9 0-15.5-6.5T85 323v-43h278V88zm-85 128q0 9-6.5 15t-14.5 6H85L0 323V24q0-9 6.5-15T21 3h278q8 0 14.5 6t6.5 15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compare(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M149 48V5h43v470h-43v-43H43q-18 0-30.5-12.5T0 389V91q0-18 12.5-30.5T43 48zm0 320V240L43 368zM341 48q18 0 30.5 12.5T384 91v298q0 18-12.5 30.5T341 432H235V240l106 128V91H235V48z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compass(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 193q9.5 0 16.5 6.5t7 16.5t-7 16.5t-16.5 6.5t-16.5-6.5t-7-16.5t7-16.5t16.5-6.5m0-190q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3M260 263l81-175l-174 81l-82 175z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ConfirmationNumber(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 149q-18 0-30.5 12.5T384 192t12.5 30.5T427 235v85q0 18-12.5 30.5T384 363H43q-18 0-30.5-12.5T0 320v-85q18 0 30.5-12.5T43 192t-12.5-30.5T0 149V64q0-18 12.5-30.5T43 21h341q18 0 30.5 12.5T427 64zM235 309v-42h-43v42zm0-96v-42h-43v42zm0-96V75h-43v42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copy(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 5v43H43v299H0V48q0-18 12.5-30.5T43 5zm64 86q17 0 29.5 12.5T405 133v299q0 18-12.5 30.5T363 475H128q-18 0-30.5-12.5T85 432V133q0-17 12.5-29.5T128 91zm0 341V133H128v299z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crop(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 304V133H171V91h170q18 0 30.5 12.5T384 133v171zm-213 43h341v42h-85v86h-43v-86H128q-18 0-30.5-12.5T85 347V133H0V91h85V5h43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CropDin(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zm0 341V43H43v298z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CropFiveFour(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 43q18 0 30.5 12.5T384 85v214q0 17-12.5 29.5T341 341H43q-18 0-30.5-12.5T0 299V85q0-17 12.5-29.5T43 43zm0 256V85H43v214z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CropFree(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 43q0-18 12.5-30.5T43 0h85v43H43v85H0zm43 213v85h85v43H43q-18 0-30.5-12.5T0 341v-85zm298 85v-85h43v85q0 18-12.5 30.5T341 384h-85v-43zm0-341q18 0 30.5 12.5T384 43v85h-43V43h-85V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CropLandscape(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 43q18 0 30.5 12.5T384 85v214q0 17-12.5 29.5T341 341H43q-18 0-30.5-12.5T0 299V85q0-17 12.5-29.5T43 43zm0 256V85H43v214z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CropPortrait(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0q18 0 30.5 12.5T299 43v298q0 18-12.5 30.5T256 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zm0 341V43H43v298z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CropSevenFive(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 85q18 0 30.5 12.5T384 128v128q0 18-12.5 30.5T341 299H43q-18 0-30.5-12.5T0 256V128q0-18 12.5-30.5T43 85zm0 171V128H43v128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CropSixteenNine(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 64q18 0 30.5 12.5T384 107v170q0 18-12.5 30.5T341 320H43q-18 0-30.5-12.5T0 277V107q0-18 12.5-30.5T43 64zm0 213V107H43v170z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CropSquare(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 21q17 0 29.5 12.5T341 64v256q0 18-12.5 30.5T299 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21zm0 299V64H43v256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CropThreeTwo(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 21q18 0 30.5 12.5T384 64v256q0 18-12.5 30.5T341 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21zm0 299V64H43v256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cutlery(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m114 221l-89-90Q0 106 0 71t25-60l150 149zm145-39l-31 31l146 147l-30 30l-146-147L51 390l-31-30l209-208q-12-24-4-56t33-57q31-30 69-35t61.5 18.5T407 84t-36 69q-25 25-56.5 33t-55.5-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Delete(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 341V85h256v256q0 18-12.5 30.5T235 384H64q-18 0-30.5-12.5T21 341M299 21v43H0V21h75L96 0h107l21 21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Delicious(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 216H213v213H0V216h213V3h214z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DesktopMac(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 3q17 0 29.5 12.5T469 45v256q0 18-12.5 30.5T427 344H277l43 64v21H149v-21l43-64H43q-18 0-30.5-12.5T0 301V45q0-17 12.5-29.5T43 3zm0 256V45H43v214z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DesktopWindows(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 3q17 0 29.5 12.5T469 45v256q0 18-12.5 30.5T427 344H277v43h43v42H149v-42h43v-43H43q-18 0-30.5-12.5T0 301V45q0-17 12.5-29.5T43 3zm0 298V45H43v256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeveloperBoard(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 128h-43v43h43v42h-43v43h43v43h-43v42q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0h298q18 0 30.5 12.5T384 43v42h43zm-86 213V43H43v298zM85 213h107v86H85zm128-42h86v64h-86zM85 85h107v107H85zm128 192h86v128h-86z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeviceHub(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 277h85v107H277v-65l-85-90l-85 90v65H0V277h85l86-85v-68q-19-7-31-23.5T128 64q0-27 18.5-45.5T192 0t45.5 18.5T256 64q0 20-12 36.5T213 124v68z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Devices(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M85 64v235h214v64H0v-64h43V64q0-18 12.5-30.5T85 21h384v43zm406 43q8 0 14.5 6t6.5 15v213q0 9-6.5 15.5T491 363H363q-9 0-15.5-6.5T341 341V128q0-9 6.5-15t15.5-6zm-22 192V149h-85v150z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DevicesOff(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 96H188l-43-43h324zM41 3l42 42l372 373l-27 27l-50-50H0v-64h43V96q0-15 10-27L14 30zm44 99v229h229zm406 37q8 0 14.5 6t6.5 15v213q0 9-6.5 15.5T491 395h-4l-64-64h46V181h-85v111l-43-43v-89q0-9 6.5-15t15.5-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dialpad(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M170.5 389q17.5 0 30 12.5T213 432t-12.5 30.5t-30 12.5t-30-12.5T128 432t12.5-30.5t30-12.5M42.5 5q17.5 0 30 12.5T85 48T72.5 78.5T42.5 91t-30-12.5T0 48t12.5-30.5T42.5 5m0 128q17.5 0 30 12.5T85 176t-12.5 30.5t-30 12.5t-30-12.5T0 176t12.5-30.5t30-12.5m0 128q17.5 0 30 12.5T85 304t-12.5 30.5t-30 12.5t-30-12.5T0 304t12.5-30.5t30-12.5m256-170q-17.5 0-30-12.5T256 48t12.5-30.5t30-12.5t30 12.5T341 48t-12.5 30.5t-30 12.5m-128 170q17.5 0 30 12.5T213 304t-12.5 30.5t-30 12.5t-30-12.5T128 304t12.5-30.5t30-12.5m128 0q17.5 0 30 12.5T341 304t-12.5 30.5t-30 12.5t-30-12.5T256 304t12.5-30.5t30-12.5m0-128q17.5 0 30 12.5T341 176t-12.5 30.5t-30 12.5t-30-12.5T256 176t12.5-30.5t30-12.5m-128 0q17.5 0 30 12.5T213 176t-12.5 30.5t-30 12.5t-30-12.5T128 176t12.5-30.5t30-12.5m0-128q17.5 0 30 12.5T213 48t-12.5 30.5t-30 12.5t-30-12.5T128 48t12.5-30.5t30-12.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiscFull(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 277v-42h43v42zm0-192h43v107h-43zM170.5 21Q241 21 291 71t50 121t-50 121t-120.5 50T50 313T0 192T50 71t120.5-50m0 214q17.5 0 30-12.5T213 192t-12.5-30.5t-30-12.5t-30 12.5T128 192t12.5 30.5t30 12.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Disqus(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M219 157q26 0 42 13t16 38t-16 38t-42 13h-23V157zm2-157q85 0 145.5 61T427 208t-60.5 147T221 416q-75 0-133-49L0 379l34-85q-18-41-18-86q0-86 60-147T221 0m112 207q0-46-30.5-74T219 105h-78v206h76q54 0 85-29t31-75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dns(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M363 213q8 0 14.5 6.5T384 235v128q0 8-6.5 14.5T363 384H21q-8 0-14.5-6.5T0 363V235q0-9 6.5-15.5T21 213zM85.5 341q17.5 0 30-12.5t12.5-30t-12.5-30t-30-12.5t-30 12.5t-12.5 30t12.5 30t30 12.5M363 0q8 0 14.5 6.5T384 21v128q0 9-6.5 15.5T363 171H21q-8 0-14.5-6.5T0 149V21q0-8 6.5-14.5T21 0zM85.5 128q17.5 0 30-12.5t12.5-30t-12.5-30t-30-12.5t-30 12.5t-12.5 30t12.5 30t30 12.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dock(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 475v-43h170v43zM213 6q18 0 30.5 12T256 48v299q0 17-12.5 29.5T213 389H43q-18 0-30.5-12.5T0 347V48q0-18 12.5-30.5T43 5zm0 298V91H43v213z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotCircle(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213 109q44 0 75.5 31.5T320 216t-31.5 75.5T213 323t-75-31.5t-31-75.5t31-75.5t75-31.5m.5-106q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3m0 384q70.5 0 120.5-50t50-121t-50-121t-120.5-50T93 95T43 216t50 121t120.5 50"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotCircleAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213 3q88 0 151 62.5T427 216t-63 150.5T213 429T62.5 366.5T0 216T62.5 65.5T213 3m.5 384q70.5 0 120.5-50t50-121t-50-121t-120.5-50T93 95T43 216t50 121t120.5 50M277 216q0 27-18.5 45.5t-45 18.5t-45.5-18.5t-19-45.5t19-45.5t45.5-18.5t45 18.5T277 216"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 128L149 277L0 128h85V0h128v128zM0 320h299v43H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dribbble(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M308 358q56-40 69-107q-35-8-66-8q-17 0-34 3q19 57 31 112m-95 29q31 0 59-11q-12-63-32-121q-49 16-87 52q-23 22-39 47q44 33 99 33M47 221q0 60 39 106q19-28 46-53q42-38 94-55q-4-10-10-22q-67 21-151 22q-13 0-18-1zm93-150q-33 16-56 45t-32 64q3 1 13 1h3q70 0 131-19q-29-54-59-91m73-17q-16 0-35 4q32 42 57 91q53-23 82-58q-45-37-104-37m131 64q-36 41-92 66q4 8 11 25q24-4 48-4q33 0 69 8q-3-53-36-95M213.5 7q88.5 0 151 62.5t62.5 151t-62.5 151t-151 62.5t-151-62.5T0 220.5t62.5-151T213.5 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drink(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 3h384l-43 389q-2 16-14 26.5T299 429H85q-16 0-28-10.5T43 392zm192 362q27 0 45.5-18.5T256 301q0-19-16-47.5T208 205l-16-19q-7 8-17.5 21.5t-28.5 44t-18 49.5q0 27 18.5 45.5T192 365m135-234l9-86H48l9 86z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dropbox(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m126 3l87 72l88-72l126 81l-87 69l87 69l-126 82l-88-73l-87 73L0 222l87-69L0 84zm87 72L87 153l126 78l127-78zm0 172l89 73l37-25v27l-126 75l-125-75v-27l38 25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Edit(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 304L236 68l80 80L80 384H0zM378 86l-39 39l-80-80l39-39q6-6 15-6t15 6l50 50q6 6 6 15t-6 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EightTracks(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M117 195h193q31 0 53-22t22-53t-22-53.5T309.5 44T256 66.5T234 120v33h-42v-33q0-49 34.5-83.5t83-34.5t83 34.5t34.5 83t-34.5 83T310 237H117q-31 0-53 22t-22 53t22 53t53 22t53-22t22-53v-34h42v34q0 48-34 82.5T117.5 429t-83-34.5t-34.5-83T34.5 229t82.5-34"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eject(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 299h299v42H0zM149 85l143 171H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EjectAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m128 115l-98 98l-30-30L128 55l128 128l-30 30zM0 320v-43h256v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Email(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 21q18 0 30.5 12.5T427 64v256q0 18-12.5 30.5T384 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21zm0 86V64L213 171L43 64v43l170 106z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmailOpen(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m426 155l1 213q0 18-12.5 30.5T384 411H43q-18 0-30.5-12.5T0 368V155q0-24 20-37L213 5l193 113q20 13 20 37M213 261l177-110L213 48L37 151z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Equalizer(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 363V21h85v342zM0 363V192h85v171zm256-235h85v235h-85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Evernote(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M249 208h50q3-10-8-21q-12-12-27-3.5T249 208m83-148q11 14 17.5 34.5t8 32T362 165q4 39 3.5 88.5T355 341q-9 61-49 80.5t-95 4.5q-22-6-32-27t-7-44q4-21 24-31.5t43-10.5v21q2 7-1 9.5t-8.5 2t-11.5.5q-8 5-9 16.5t8.5 21T245 393q33-1 40-12t5-48q2-19-14-32t-36-14q-37 3-65-43q-1 2-1 10.5V286q-1 15-15 23.5T128 321q-60 5-84-19q-34-36-43-120q-7-48 22-69h81q4-2 10.5-9.5T122 95V52q1-4 .5-14.5t1-17T130 9q22-11 47-4t38 28h55q43 6 62 27M87 95H18l86-88v70z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Explicit(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zm-85 128V85H128v214h128v-43h-85v-43h85v-42h-85v-43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Exposure(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M277 323h-42v-43h42v-43h43v43h43v43h-43v42h-43zM384 3q18 0 30.5 12.5T427 45v342q0 17-12.5 29.5T384 429H43q-18 0-30.5-12.5T0 387V45q0-17 12.5-29.5T43 3zM64 67v42h128V67zm320 320V45L43 387z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExposureAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zM53 96v32h43v43h32v-43h43V96h-43V53H96v43zm288 245V43L43 341zm-42-42H192v-32h107z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M235 32q79 0 142.5 44.5T469 192q-28 71-91.5 115.5T235 352T92 307.5T0 192q28-71 92-115.5T235 32m0 267q44 0 75-31.5t31-75.5t-31-75.5T235 85t-75.5 31.5T128 192t31.5 75.5T235 299m-.5-171q26.5 0 45.5 18.5t19 45.5t-19 45.5t-45.5 18.5t-45-18.5T171 192t18.5-45.5t45-18.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeOff(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M235 85q-20 0-39 8l-46-46q41-15 84-15q79 0 143 44.5T469 192q-23 60-73 101l-62-62q7-19 7-39q0-44-31-75.5T235 85M21 27L48 0l379 378l-27 27l-63-62l-9-9q-45 18-93 18q-79 0-143-44.5T0 192q25-64 80-106L70 76zm118 118q-11 23-11 47q0 44 31.5 75.5T235 299q24 0 47-12l-33-33q-8 2-14 2q-27 0-45.5-18.5T171 192q0-7 1-14zm92-17h4q26 0 45 19t19 45l-1 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eyedropper(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M378 56q6 6 6 15t-6 15l-67 67l41 41l-30 30l-30-30l-191 190H0V283L190 92l-30-30l30-30l41 41l67-67q6-6 15-6t15 6zM84 341l172-172l-41-41L43 300z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Face(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M149 211q11 0 19 7.5t8 18.5t-8 19t-19 8t-18.5-8t-7.5-19t7.5-18.5T149 211m128 0q11 0 19 7.5t8 18.5t-8 19t-19 8t-18.5-8t-7.5-19t7.5-18.5T277 211M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3m0 384q70.5 0 120.5-50t50-121q0-24-7-48q-24 5-48 5q-53 0-99-24t-75-66q-33 80-111 115q-1 10-1 18q0 71 50 121t120.5 50"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Facebook(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M145 429H66V235H0v-76h66v-56q0-48 27-74t72-26q36 0 59 3v67l-41 1q-22 0-30 9t-8 27v49h76l-10 76h-66z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacebookBox(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M363 88h-54q-31 0-52.5 22T235 163v53h-43v64h43v149h64V280h64v-64h-64v-43q0-8 6-14.5t15-6.5h43zM0 3h427v426H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastForward(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 320V64l181 128zM192 64l181 128l-181 128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastRewind(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M181 320L0 192L181 64zm11-128L373 64v256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Favorite(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m213 391l-31-28q-42-38-62-56.5T72 258t-40.5-49t-22-43.5T0 117q0-49 34-83t83-34q58 0 96 45q38-45 96-45q50 0 84 34t34 83q0 24-10 48.5T395 209t-40.5 49t-48 48.5T244 364z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FavoriteOutline(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M309 0q50 0 84 34t34 83q0 24-10 48.5T395 209t-40.5 49t-48 48.5T244 363l-31 28l-31-27q-42-39-62-57.5T71.5 258T31 209T9.5 165.5T0 117q0-49 34-83t83-34q58 0 96 45q38-45 96-45m-94 332q49-44 71.5-65.5T336 215t37.5-52.5T384 117q0-32-21.5-53T309 43q-24 0-45.5 14T233 93h-40q-8-22-29.5-36T117 43q-32 0-53 21t-21 53q0 23 10 45.5T90.5 215t50 51.5T211 332l2 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Female(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 429H64V301H0l54-162q4-14 15.5-22t25.5-8h2q14 0 25 8t16 22l54 162h-64zM96 88q-18 0-30.5-12.5T53 45.5t12.5-30T96 3t30.5 12.5t12.5 30t-12.5 30T96 88"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func File(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 3h170l128 128v256q0 17-12.5 29.5T299 429H42q-17 0-29.5-12.5T0 387V45q0-17 12.5-29.5T43 3m149 149h117L192 35z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePlus(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m213 3l128 128v256q0 17-12.5 29.5T299 429H42q-17 0-29.5-12.5T0 387V45q0-17 12.5-29.5T43 3zm43 298v-42h-64v-64h-43v64H85v42h64v64h43v-64zm-64-149h117L192 35z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileText(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m213 3l128 128v256q0 17-12.5 29.5T299 429H42q-17 0-29.5-12.5T0 387V45q0-17 12.5-29.5T43 3zm43 341v-43H85v43zm0-85v-43H85v43zm-64-107h117L192 35z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterBandW(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zm0 341V43H192v128L43 341h149V171z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterCenterFocus(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 256v85h85v43H43q-18 0-30.5-12.5T0 341v-85zm0-213v85H0V43q0-18 12.5-30.5T43 0h85v43zM341 0q18 0 30.5 12.5T384 43v85h-43V43h-85V0zm0 341v-85h43v85q0 18-12.5 30.5T341 384h-85v-43zM192 128q27 0 45.5 18.5T256 192t-18.5 45.5T192 256t-45.5-18.5T128 192t18.5-45.5T192 128"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterFrames(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 85q18 0 30.5 12.5T427 128v299q0 17-12.5 29.5T384 469H43q-18 0-30.5-12.5T0 427V128q0-18 12.5-30.5T43 85h85l85-85l86 85zm0 342V128h-96l-74-75l-75 75H43v299zm-43-256H85v213h256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterList(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M149 320v-43h86v43zM0 64h384v43H0zm64 149v-42h256v42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterTiltShift(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M191 47q-46 5-83 34L77 51q50-41 114-47zm156 4l-30 30q-38-29-83-34V4q63 6 113 47m34 144q-5-46-34-84l30-30q41 50 48 114zM78 111q-29 38-35 84H0q6-64 47-114zM43 237q6 46 35 83l-31 31Q6 301 0 237zm233-21q0 27-18.5 45.5t-45 18.5t-45.5-18.5t-19-45.5t19-45.5t45.5-18.5t45 18.5T276 216m71 105q29-38 34-83h44q-7 63-48 113zm-113 64q46-6 83-34l30 30q-50 41-113 47zm-157-4l31-30q37 29 83 34v43q-64-6-114-47"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fire(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M203 6q64 52 101 126t37 159q0 70-50 120t-120.5 50T50 411T0 291q0-108 69-190l-1 8q0 33 22.5 56t55.5 23q32 0 52-23t20-56q0-21-3.5-46.5T207 22zm-39 391q43 0 73-30t30-72q0-45-13-86q-30 41-98 55q-29 6-44.5 23.5T96 330q0 28 20 47.5t48 19.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FiveHundredPx(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M281 247q-12 13-12 14q-14 13-21 17q-26 19-59 12q-30-6-45-36q-1-2-2-3q-8 12-8 13q-21 28-57 29q-23 1-40-6q-36-16-37-56h40l1 4q5 20 21 24q21 5 35-9q9-10 10.5-25t-6.5-26q-8-12-23-14.5T52 191q-4 3-6 5q-3 5-11 5H6q9-53 20-110h111v33H58q-3 0-4 3q-1 4-7 39v3q21-19 52-14q27 5 42 34q1-1 1-3q2-2 2-3q21-43 69-36q26 3 46 22q1 1 23 24l2-2q22-23 24-25q17-16 38-19q23-3 43 5.5t30 30.5q16 38-2 74q-16 33-55 35q-29 1-51-17q-3-2-25-22q0-1-5-6m-76 5q12 0 24-7q8-4 28-23q1-2 0-4q-2-1-9.5-8T236 199q-9-8-21-12q-14-4-25 2t-15 20q-1 4-1 8q-1 16 8 25.5t23 9.5m99-32q22 20 24 21q13 12 30 11q25 0 30-24q1-7 0-15q-2-13-11.5-21t-22.5-6q-15 1-28 13q-1 1-22 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M201 64h119v213H171l-9-42H43v149H0V21h192z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flare(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 219v42H0v-42zm46-69l-30 30l-45-46l30-30zM256 5v128h-43V5zm114 129l-45 46l-30-30l45-46zm-29 85h128v42H341zm-106.5-43q26.5 0 45.5 18.5t19 45.5t-19 45.5t-45.5 18.5t-45-18.5T171 240t18.5-45.5t45-18.5M295 330l30-30l45 46l-30 30zM99 346l45-46l30 30l-45 46zm114 129V347h43v128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flash(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 3h213l-85 170h85L64 429V237H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlashAuto(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 3h213l-85 192h85L64 451V259H0zm341 0l69 192h-41l-15-43h-68l-15 43h-41L299 3zm-46 120h50l-25-78z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlashOff(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m27 24l336 336l-27 27l-89-89l-76 131V237h-64v-79L0 51zm293 149l-33 57L107 49V3h213l-85 170z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flattr(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M426 263q0 46-13 81q-25 66-96 81q-20 4-43 4H31v-1l45-45l124-124l.5-.5l1.5-.5q4-4 7-3q4 2 4 8v67q0 1 2 1q47-1 55-1q8-1 19-5q28-9 35-42q3-16 3-33V136q0-3 3-6l93-93q0-1 2-4l1 1h1v4q0 24-1 225m-326 33q0 3-3 6L3 396l-3 3V170q0-45 12-79q25-69 99-84q19-4 42-4h243q-1 0-2 2q0 1-1 1q-27 28-82.5 83.5L227 173q-3 2-3 3q-4 3-7 1q-1-2-4-6v-69q0-1-1-1q-50 1-59 2q-4 0-13 3q-31 9-37 44q-3 15-3 34q-1 25 0 112"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flickr(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M97 289q-40 0-68.5-28.5T0 192.5T28.5 125T97 97t68 28t28 67.5t-28 68T97 289m233 0q-40 0-68.5-28.5t-28.5-68t28.5-67.5T330 97t68.5 28t28.5 67.5t-28.5 68T330 289"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlightLand(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 368h405v43H0zm153-122L40 215l-34-9V96l31 8l20 50l106 28V5l41 11l59 193l113 30q13 3 19.5 14.5t3 24.5t-15 19.5T359 301l-113-30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlightTakeoff(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 335h405v43H14zm417.5-199.5q3.5 12.5-3 24T409 175l-114 30l-92 25l-114 30l-34 10l-16-29l-39-67l31-9l42 33l106-28L91 17l41-11l147 137l113-30q13-4 24.5 3t15 19.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flip(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 432v-43h43v43zm85-256v-43h43v43zM0 91q0-18 12.5-30.5T43 48h85v43H43v298h85v43H43q-18 0-30.5-12.5T0 389zm341-43q18 0 30.5 12.5T384 91h-43zM171 475V5h42v470zm170-128v-43h43v43zM256 91V48h43v43zm85 170v-42h43v42zm0 171v-43h43q0 18-12.5 30.5T341 432"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlipToBack(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 85v43H85V85zm0 86v42H85v-42zm0-171v43H85q0-18 12.5-30.5T128 0m85 256v43h-42v-43zM341 0q18 0 30.5 12.5T384 43h-43zM213 0v43h-42V0zm-85 299q-18 0-30.5-12.5T85 256h43zm213-86v-42h43v42zm0-85V85h43v43zm0 171v-43h43q0 18-12.5 30.5T341 299M43 85v256h256v43H43q-18 0-30.5-12.5T0 341V85zm213-42V0h43v43zm0 256v-43h43v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlipToFront(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 213v-42h43v42zm0 86v-43h43v43zm43 85q-18 0-30.5-12.5T0 341h43zM0 128V85h43v43zm256 256v-43h43v43zM341 0q18 0 30.5 12.5T384 43v213q0 18-12.5 30.5T341 299H128q-18 0-30.5-12.5T85 256V43q0-18 12.5-30.5T128 0zm0 256V43H128v213zM171 384v-43h42v43zm-86 0v-43h43v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Floppy(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m299 0l85 85v256q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zM192 341q27 0 45.5-18.5t18.5-45t-18.5-45.5t-45.5-19t-45.5 19t-18.5 45.5t18.5 45T192 341m64-213V43H43v85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flower(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M350 249q30 17 47 47t17 63q-29 17-63 17.5T286 359q-9-5-17-11q2 10 2 20q0 35-17.5 64.5T207 479q-29-17-46.5-46.5T143 368q0-10 2-20q-9 7-17 11q-31 17-65 17T0 359q0-34 17-63.5T64 248q8-4 18-8q-10-4-18-9q-30-17-47-47T0 121q29-17 63-17.5t65 17.5q8 4 17 11q-2-10-2-20q0-35 17.5-64.5T207 1q29 17 46.5 46.5T271 112q0 10-2 20q9-7 17-11q31-18 65-17.5t63 17.5q0 33-17 63t-47 47q-8 5-18 9q10 4 18 9m-143 76q35 0 60-25t25-60t-25-60t-60-25t-60 25t-25 60t25 60t60 25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlowerAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 453q0-79 56-135.5T384 261q0 80-56 136t-136 56M55 203q0-34 31-48q-31-15-31-48q0-22 16-38t38-16q17 0 30 10v-4q0-22 15.5-38T192 5t37.5 16T245 59v4q14-10 30-10q22 0 38 16t16 38q0 33-31 48q31 14 31 48q0 22-16 37.5T275 256q-17 0-30-9v4q0 22-15.5 37.5T192 304t-37.5-15.5T139 251v-4q-14 9-30 9q-22 0-38-15.5T55 203m137-102q-22 0-37.5 16T139 155t15.5 37.5T192 208t37.5-15.5T245 155t-15.5-38t-37.5-16M0 261q80 0 136 56.5T192 453q-80 0-136-56T0 261"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m171 21l42 43h171q18 0 30.5 12.5T427 107v213q0 18-12.5 30.5T384 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOutline(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 64q18 0 30.5 12.5T427 107v213q0 18-12.5 30.5T384 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21h128l42 43zm0 256V107H43v213z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderPerson(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 64q18 0 30.5 12.5T427 107v213q0 18-12.5 30.5T384 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21h128l42 43zm-106.5 64q-17.5 0-30 12.5t-12.5 30t12.5 30t30 12.5t30-12.5t12.5-30t-12.5-30t-30-12.5M363 299v-22q0-19-29.5-30.5t-56-11.5t-56 11.5T192 277v22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderStar(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 21h128l42 43h171q18 0 30.5 12.5T427 107v213q0 18-12.5 30.5T384 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21m233 107l-25 60l-65 5l49 43l-15 63l56-33l56 33l-14-63l49-43l-65-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderStarAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 64q18 0 30.5 12.5T427 107v213q0 18-12.5 30.5T384 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21h128l42 43zM247 320l-20-87l67-58l-89-8l-34-82l-35 82l-89 8l68 58l-21 87l77-45z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Font(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m169 248l44-118l44 118zM384 3q18 0 30.5 12.5T427 45v342q0 17-12.5 29.5T384 429H43q-18 0-30.5-12.5T0 387V45q0-17 12.5-29.5T43 3zm-86 352h44L233 77h-40L84 355h45l24-64h120z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatAlignCenter(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M85 256h214v43H85zM0 384v-43h384v43zm0-171v-42h384v42zM85 85h214v43H85zM0 0h384v43H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatAlignJustify(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 384v-43h384v43zm0-85v-43h384v43zm0-86v-42h384v42zm0-85V85h384v43zM0 0h384v43H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatAlignLeft(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 256v43H0v-43zm0-171v43H0V85zM0 213v-42h384v42zm0 171v-43h384v43zM0 0h384v43H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatAlignRight(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 384v-43h384v43zm128-85v-43h256v43zM0 213v-42h384v42zm128-85V85h256v43zM0 0h384v43H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatBold(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M183 166q21 10 33.5 29.5T229 239q0 34-23 57.5T150 320H0V21h133q36 0 61 25t25 61q0 35-36 59M64 75v64h64q13 0 22.5-9.5t9.5-23t-9.5-22.5t-22.5-9zm75 192q13 0 22.5-9.5t9.5-23t-9.5-22.5t-22.5-9H64v64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatClear(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m27 43l6 5l308 309l-27 27l-121-121l-33 78H96l53-123L0 70zm58 0h299v64H260l-34 80l-45-44l16-36h-52L85 47z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatClearAll(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 213v-42h298v42zM0 299v-43h299v43zM85 85h299v43H85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatColorFill(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M353 191q10 9 10 22.5T353 236L236 353q-9 10-22.5 10T191 353L73 236q-9-9-9-22.5t9-22.5L183 81l-51-51l31-30zm-242 22h205L213 111zm294 32q43 47 43 75q0 18-12.5 30.5t-30 12.5t-30-12.5T363 320q0-13 10.5-31.5T395 258zM0 427h512v85H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatColorReset(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 235q0 14-3 28L112 79q14-19 28.5-37.5T163 14l8-10q5 6 13.5 16.5t30.5 40t39 56.5t31 60.5t14 57.5m-19 66l58 59l-27 27l-56-56q-36 32-84 32q-53 0-90.5-37.5T43 235q0-35 28-88L0 76l27-28l154 155z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatColorText(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 363h512v85H0zM235 0h42l117 299h-48l-23-64H189l-24 64h-48zm-30 192h102L256 57z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatIndentDecrease(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M171 299v-43h213v43zM0 192l85-85v170zm0 192v-43h384v43zM0 0h384v43H0zm171 128V85h213v43zm0 85v-42h213v42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatIndentIncrease(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 384v-43h384v43zm0-277l85 85l-85 85zm171 192v-43h213v43zM0 0h384v43H0zm171 128V85h213v43zm0 85v-42h213v42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatItalic(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M85 21h171v64h-60l-72 171h47v64H0v-64h60l72-171H85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatLineSpacing(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M96 85v214h53l-74 74l-75-74h53V85H0l75-74l74 74zm85-42h256v42H181zm0 298v-42h256v42zm0-128v-42h256v42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatListBulleted(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M32 160q13 0 22.5 9.5T64 192t-9.5 22.5T32 224t-22.5-9.5T0 192t9.5-22.5T32 160m0-128q13 0 22.5 9.5T64 64t-9.5 22.5T32 96T9.5 86.5T0 64t9.5-22.5T32 32m0 260q12 0 20 8t8 20t-8 20t-20 8t-20-8t-8-20t8-20t20-8m64 49v-42h299v42zm0-128v-42h299v42zm0-170h299v42H96z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatListNumbered(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 299v-22h64v86H0v-22h43v-10H21v-22h22v-10zm21-192V43H0V21h43v86zM0 171v-22h64v20l-38 44h38v22H0v-20l38-44zM107 43h298v42H107zm0 298v-42h298v42zm0-128v-42h298v42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatLtr(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M85 173q-35 0-60-25T0 88t25-60T85 3h171v42h-43v235h-42V45h-43v235H85zm256 171l-85 85v-64H0v-42h256v-64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatRtl(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 173q-35 0-60-25T43 88t25-60t60-25h171v42h-43v235h-43V45h-42v235h-43zM85 323h256v42H85v64L0 344l85-85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatSize(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 21h277v64H299v256h-64V85H128zM0 192v-64h192v64h-64v149H64V192z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatStrikethrough(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M149 341v-64h86v64zM43 21h298v64H235v64h-86V85H43zM0 235v-43h384v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatStrikethroughS(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M105 173q-5-4-7-8q-11-22-11-47t13-47q8-18 30-36q19-14 49-24q26-8 62-8q40 0 66 10q25 6 49 26q20 16 30 40q11 25 11 52h-86q0-11-4-24q-3-13-13-19q-10-10-21-13q-17-4-30-4t-30 4q-8 2-21 11q-10 7-13 15q-4 13-4 19q0 22 21 34q14 9 43 19zm364 43v43h-91q1 1 1.5 2t1 3t1.5 3q8 20 8 47q0 24-10 49q-8 18-30 36q-21 18-47 24q-26 8-62 8q-15 0-40-4q-13-2-39-10q-13-7-34-20q-14-8-28-25q-13-17-19-34q-6-20-6-45h85q0 21 6 34q5 8 17 21q10 10 26 13q21 4 34 4t30-4q3-2 10-5t9-6q10-6 13-15q4-12 4-19q0-13-2-19q-3-11-13-17q-17-12-25-15q-2-1-7.5-3t-7.5-3H0v-43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatSubject(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213 299v42H0v-42zm128-171v43H0v-43zM0 256v-43h341v43zM0 43h341v42H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatUnderlined(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M149 299q-53 0-90.5-37.5T21 171V0h54v171q0 31 21.5 52.5T149 245t53-21.5t22-52.5V0h53v171q0 53-37.5 90.5T149 299M0 341h299v43H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatValignBottom(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m256 213l-85 86l-86-86h64V0h43v213zM0 341h341v43H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatValignCenter(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m85 389l86-85l85 85h-64v86h-43v-86zM256 91l-85 85l-86-85h64V5h43v86zM0 219h341v42H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatValignTop(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m85 171l86-86l85 86h-64v213h-43V171zM0 0h341v43H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Forward(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M171 107V21l170 171l-170 171v-86H0V107z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardFive(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 261.5Q0 191 50 141t121-50V5l106 107l-106 107v-86q-53 0-90.5 38T43 261.5t37.5 90t90 37.5t90.5-37.5t38-90.5h42q0 71-50 121t-120.5 50T50 382T0 261.5M143 281l4-47h51v15h-36l-2 19q2 0 2-2q0-1 1-1t1-2h9q8 0 10 3q2 1 5 3t4 3q2 2 6 11q3 4 3 12.5t-3 10.5q0 1-2 4.5t-4 6.5q-2 2-11 6q-4 2-12.5 2t-10.5-2q-1-1-5-2t-6-2q-3-1-6-9q-2-4-2-10h17q0 4 4 8q2 2 9 2q4 0 6-2l4-4q2-4 2-6v-13l-2-4l-4-5q-4-2-6-2h-5q-2 0-4 2q-1 1-1.5 1t-.5 1l-2 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardTen(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 261.5Q0 191 50 141t121-50V5l106 107l-106 107v-86q-53 0-90.5 38T43 261.5t37.5 90t90 37.5t90.5-37.5t38-90.5h42q0 71-50 121t-120.5 50T50 382T0 261.5M145 325h-17v-70l-21 6v-15l38-12h2v91zm94-38q0 13-2 17l-7 13q-6 6-10 6q-2 0-6.5 1t-6.5 1q-9 0-13-2q-2-1-5-3t-6-3q-2-1-6-13q-2-6-2-17v-15q0-13 2-17l6-13q7-6 11-6q2 0 6.5-1t6.5-1q8 0 13 2q2 1 5 3t5 3q3 1 7 13q2 6 2 17zm-17-17v-11q-2-4-2-6l-5-4q-2-3-6-3t-6 3l-5 4q-2 4-2 6v43q2 4 2 6t2 3t3 2q2 2 6 2t6-2l5-5q2-4 2-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardThirty(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M119 272h9q6 0 10.5-4.5t4.5-8.5v-4q-2-2-2-4t-4-2h-11q-2 2-4.5 2t-2.5 4v4H98q0-6 2-10.5t6.5-8.5t8.5-4q1 0 5.5-1t5.5-1q8 0 13 2q2 1 5 2t5 2q3 1 7 9q2 4 2 10v7q-2 4-2 6q0 4-5 4q-2 0-6 5q9 4 11 8q4 9 4 13q0 8-2 11q-1 1-3 4t-4 4q-4 4-10 4q-2 0-6.5 1t-6.5 1q-9 0-11-2q-1-1-5-2t-5-2q-3-1-7-8q-2-5-2-13h17v4q2 2 2 4t5 2h10q2-2 4.5-2t2.5-4v-11q-2-2-2-4t-5-2h-13zm122 15q0 13-2 17l-6 13q-7 6-11 6q-2 0-6.5 1t-6.5 1q-8 0-13-2q-2-1-5-3t-5-3q-3-1-7-13q-2-6-2-17v-15q0-13 2-17l7-13q6-6 10-6q2 0 6.5-1t6.5-1q9 0 13 2q2 1 5 3t6 3q2 1 6 13q2 6 2 17zm-19-17v-11q-2-4-2-6l-5-4q-2-3-6-3t-6 3l-5 4q-2 4-2 6v43q2 4 2 6l5 5q2 2 6 2t6-2l5-5q2-4 2-6zM0 261.5Q0 191 50 141t121-50V5l106 107l-106 107v-86q-53 0-90.5 38T43 261.5t37.5 90t90 37.5t90.5-37.5t38-90.5h42q0 71-50 121t-120.5 50T50 382T0 261.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fullscreen(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 235v64h64v42H0V235zM0 149V43h107v42H43v64zm256 150v-64h43v106H192v-42zM192 43h107v106h-43V85h-64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FullscreenAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m235 53l42 54h-85zm128 96l53 43l-53 43zm-256 0v86l-54-43zm170 128l-42 54l-43-54zM427 0q17 0 29.5 12.5T469 43v298q0 18-12.5 30.5T427 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zm0 342V42H43v300z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FullscreenExit(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 277v-42h107v106H64v-64zm64-170V43h43v106H0v-42zm128 234V235h107v42h-64v64zm43-234h64v42H192V43h43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Functions(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 21v64H107l106 107l-106 107h149v64H0v-43l139-128L0 64V21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gamepad(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m277 120l-64 64l-64-64V3h128zm-160 32l64 64l-64 64H0V152zm32 160l64-64l64 64v117H149zm160-160h118v128H309l-64-64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GasStation(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 90q16 16 16 38v203q0 22-15.5 37.5T299 384t-38-15.5t-16-37.5V224h-32v160H0V43q0-18 12.5-30.5T43 0h128q17 0 29.5 12.5T213 43v149h22q17 0 29.5 12.5T277 235v96q0 8 6.5 14.5t15 6.5t15-6.5T320 331V177q-11 4-21 4q-22 0-38-15.5T245 128q0-17 9.5-30.5T280 78l-45-45l22-22l80 79zm-165 59V43H43v106zm127.5 0q8.5 0 15-6t6.5-15t-6.5-15t-15-6t-15 6t-6.5 15t6.5 15t15 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gesture(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M37 83L0 47q5-6 19-20Q45 0 77 0q18 0 35.5 15T130 61q0 20-6 34t-21 36q-29 43-40 75q-5 18-2.5 29.5T71 247q9 0 24-18q16-17 48-58q18-22 46-41t60-19q42 0 62.5 27.5T335 200h52v53h-52q-6 69-36.5 100T235 384q-28 0-48.5-19.5T166 318q0-33 30-69.5t85-45.5v-3q-1-8-2.5-12.5t-5-10.5t-11-9t-18.5-3q-18 0-39 20t-48 53q-16 19-23.5 28T114 284.5T91 297q-30 10-56-9q-29-22-29-64q0-14 6-32.5T27 156t16.5-30T59 101.5T67 89q18-28 7-32q-8-3-37 26m199 249q14 0 27.5-18t17.5-57q-30 8-45.5 27T220 316q0 7 5 11.5t11 4.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gif(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M139 128h32v128h-32zm-54 0q10 0 16 6.5t6 14.5v11H32v64h43v-32h32v43q0 8-6 14.5T85 256H21q-9 0-15-6.5T0 235v-86q0-8 6-14.5t15-6.5zm214 32h-64v21h42v32h-42v43h-32V128h96z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Github(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 0q88.5 0 151 62.5T427 213q0 70-41 125.5T281 416q-14 2-14-11v-58q0-27-15-40q44-5 70.5-27t26.5-77q0-34-22-58q11-26-2-57q-18-5-58 22q-26-7-54-7t-53 7q-18-12-32.5-17.5T107 88h-6q-12 31-2 57q-22 24-22 58q0 55 27 77t70 27q-11 10-13 29q-42 18-62-18q-12-20-33-22q-2 0-4.5.5t-5 3.5t8.5 9q14 7 23 31q1 2 2 4.5t6.5 9.5t13 10.5T130 371t30-2v36q0 13-14 11q-64-22-105-77.5T0 213q0-88 62.5-150.5T213.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GithubAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 151q0 38-10.5 65T343 257.5t-40 21t-47 9.5q19 16 19 51v90H136v-65q-16 3-29.5 3t-23-2.5t-17-6.5t-12-8.5T47 341t-4-7l-1-3q-6-14-13.5-24T16 294l-5-3q-11-9-11-12.5t7-4.5h6q12 1 23 8t15 14l5 6q27 47 81 23q3-24 18-37q-27-3-47-9.5t-39.5-21T38 216t-11-65q0-43 29-74q-13-33 3-74q3 1 8-.5t25 6T136 32q33-9 70-10q36 1 70 10q23-16 42.5-23T345 2l7 1q17 41 3 74q29 31 29 74M32 286.5q1-2.5-2.5-4t-4.5 1t2.5 4t4.5-1M43.5 299q2.5-2-1-5.5t-6-1.5t1 5.5t6 1.5M54 315q3-2 0-6.5t-6-2.5t0 6.5t6 2.5m15.5 16q2.5-3-1.5-7.5t-7-1.5t1.5 7.5t7 1.5m20.5 8.5q1-3.5-4.5-5.5t-6.5 2t4.5 5.5t6.5-2m17 5.5q6 0 6-4t-6-4t-6 4t6 4m22-2q3-1 4.5-2.5t.5-2.5q0-5-6-4q-3 1-4.5 2.5t-.5 3.5q0 4 6 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GithubBox(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 3h341q18 0 30.5 12.5T427 45v342q0 17-12.5 29.5T384 429H274q-7-1-7-21v-58q0-27-15-40q44-5 70.5-27t26.5-78q0-33-22-57q11-26-2-57q-18-6-58 22q-26-7-54-7t-53 7q-18-12-32.5-17.5T107 91h-6q-12 31-2 57q-22 24-22 57q0 55 27 77.5t70 27.5q-11 10-13 29q-42 18-62-18q-12-20-33-22q-2 0-4.5.5T56 303t8 9q15 7 24 31q1 2 2 4.5t6.5 9.5t13 10.5T130 374t30-2v36q0 20-8 21H43q-18 0-30.5-12.5T0 387V45q0-17 12.5-29.5T43 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globe(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3M192 385v-41q-18 0-30.5-12.5T149 301v-21L47 178q-4 20-4 38q0 65 42.5 113T192 385m147-54q45-49 45-115q0-53-29.5-96T277 58v9q0 17-12.5 29.5T235 109h-43v43q0 9-6.5 15t-14.5 6h-43v43h128q9 0 15 6.5t6 14.5v64h22q14 0 25 8.5t15 21.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3M361 131q-32-56-92-76q19 35 29 76zM213 46q-27 39-40 85h81q-14-46-41-85M48 259h72q-3-25-3-43t3-43H48q-5 23-5 43t5 43m18 42q32 56 92 76q-19-35-29-76zm63-170q10-41 29-76q-60 20-92 76zm84 255q27-39 41-85h-81q13 46 40 85m50-127q4-25 4-43t-4-43H163q-3 25-3 43t3 43zm6 118q60-20 92-76h-63q-10 41-29 76m37-118h72q6-23 6-43t-6-43h-72q3 25 3 43t-3 43"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeLock(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 69q9 0 15 6.5t6 15.5v85q0 9-6 15t-15 6H341q-8 0-14.5-6t-6.5-15V91q0-9 6.5-15.5T341 69V59q0-22 16-38t38-16t37.5 16T448 59zm-17 0V59q0-15-10.5-26T395 22t-26 11t-11 26v10zm-49 171h44q1 12 1 21q0 89-62.5 151.5t-151 62.5t-151-62.5T0 261.5t62.5-151T213 48q33 0 64 10v54q0 18-12.5 30.5T235 155h-43v42q0 9-6.5 15.5T171 219h-43v42h128q9 0 15 6.5t6 15.5v64h22q14 0 25 8t15 21q45-49 45-115q0-7-2-21M192 431v-42q-18 0-30.5-12.5T149 347v-22L47 223q-4 20-4 38q0 65 42.5 113.5T192 431"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Google(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M214 186v-1h201q3 12 3 36q0 93-56.5 150.5T213 429q-88 0-150.5-62T0 216T62 65T213 3q87 0 144 57l-57 56q-33-33-86-33q-54 0-92.5 39.5t-38.5 95t38.5 94.5t92.5 39q31 0 55-9.5t37.5-24.5t20.5-29.5t10-27.5H214z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleDrive(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m140 35l73 128L73 408L0 280zm43 245h280l-73 128H110zm268-21H305L158 3h146z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleEarth(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M222 121q114 108 165 114q1-11 1-19q0-25-7-50q-4 9-11 10t-15.5-5.5T339 156t-14.5-18.5t-10-15t-3.5-6.5q-47-66-163-62q-32 13-56 36q65-30 130 31m143 182q11-20 16-39q-33-3-85.5-29.5T208 183l-35-25q-74-58-127 9q-8 24-8 49q0 38 16 73q9-26 25-26q15 0 40.5 13.5T161 295q10 3 31 10l31.5 10.5L250 322l30 3q12 0 22-1.5t20-4.5t15.5-4.5t15.5-6t12-5.5m-152 88q76 0 128-56q-45 13-83.5 13t-62.5-7l-25-8q-26-8-31 6t7 38q32 14 67 14m0-388q88 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleGlass(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M235 195h181v21q0 89-58 151t-145 62q-88 0-150.5-62.5T0 216T62.5 65.5T213 3q89 0 148 65l-38 38q-43-50-110-50q-66 0-113 47T53 216t47 113t113 47q56 0 96.5-36t50.5-92H235z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleMaps(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 5q44 0 75.5 31.5T448 112q0 22-26.5 67.5t-52 92.5t-22.5 75q0 5-5.5 5t-5.5-5q2-28-23-75t-51.5-92.5T235 112q0-44 31-75.5T341 5m.5 64q-17.5 0-30 12.5T299 112t12.5 30.5t30 12.5t30-12.5T384 112t-12.5-30.5t-30-12.5M43 48h185q-20 32-20 69q0 26 32 83L1 439l-1-7V91q0-18 12.5-30.5T43 48m267 275l-51-51l14-15q24 39 37 66m61 152H56l157-158zm56-248v205l-1 7l-72-72q3-9 7-18.5t9-20t9.5-19t12-21.5t11-19.5T415 247zm-327 24q-17 0-27-7t-10-19q0-14 18-21q10-3 22-3h5q13 10 18 15t5 12q0 9-9 16t-22 7M75 129q0-10 5.5-15.5T93 108q13 0 20.5 12t7.5 25q0 11-6.5 15.5T101 165q-11 0-18.5-11.5T75 129m52 62l-7-6q-6-5-6-9q0-7 7-12q17-13 17-29q0-14-14-26h12l9-9h-43q-21 0-32.5 11.5T58 139q0 13 9 23t25 10h5l-2 8q0 7 6 14q-24 1-40 11q-16 9-16 25q0 13 11.5 21.5T90 260q25 0 39.5-12t14.5-27q0-16-17-30"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleOld(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M147 403q36 0 59-17.5t23-41.5q0-20-12-33.5T169 271h-14q-33 0-59 9q-48 17-48 57q0 30 27 48t72 18M81 79q0 36 19 66.5t50 30.5q17 0 34.5-12.5T202 121q0-33-20-66t-52-33q-21 0-35 14.5T81 79m139 165q22 19 33.5 36t11.5 43q0 43-38.5 74.5T119 429q-58 0-88.5-23.5T0 348q0-43 42-67q39-24 107-29q-17-19-17-36q0-6 7-23h-15q-41 0-65.5-26.5T34 106q0-44 31.5-73.5T154 3h113l-23 22h-32q37 32 37 71q0 19-7.5 34.5T226 154t-23 20q-18 14-18 29q0 13 15 26z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePages(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 43q0-18 12.5-30.5T43 0h128v107L85 85l22 86H0zm107 170l-22 86l86-22v107H43q-18 0-30.5-12.5T0 341V213zm192 86l-22-86h107v128q0 18-12.5 30.5T341 384H213V277zM341 0q18 0 30.5 12.5T384 43v128H277l22-86l-86 22V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePlay(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 397V35Q0 14 18 6l210 210L18 426q-18-9-18-29m295-114L65 415l181-181zm71-92q13 10 13 25t-12 25l-49 28l-54-53l54-53zM65 17l230 132l-49 49z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePlus(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M137 167h128q7 37-3 72q-10 34-35 57q-23 21-56 29q-36 8-70-1q-27-7-49-25q-24-19-37-45q-22-42-12-89q3-18 12-34q24-50 77-68q46-16 92 1q24 9 44 27q-2 3-7 7.5t-6 6.5q-4 3-12.5 11.5T190 130q-13-13-30-18q-20-6-40-1q-24 5-41 22q-13 14-20 33q-9 26 0 53q9 26 32 42q14 10 30 13q15 3 33 0q17-3 30-12q23-15 27-42h-74zm290 3v34h-47v46h-34v-46h-47v-34h47v-47h34v47z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePlusBox(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 3h340q17 0 30.5 17T427 56v330q0 18-12.5 30.5T384 429H43q-18 0-30.5-12.5T0 387V46q0-18 12.5-30.5T43 3m108 104q-45 0-76.5 32T43 216t31.5 77t76.5 32q47 0 75.5-29.5T255 219q0-13-2-19H151v38h62q-3 17-18 31.5T151 284q-28 0-47.5-20T84 216t19.5-48t47.5-20q27 0 43 16l30-28q-29-29-73-29m171 62v31h-31v31h31v31h31v-31h30l1-31h-31v-31z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gps(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M425 219h44v42h-44q-7 67-54.5 114.5T256 431v44h-43v-44q-66-8-114-55.5T44 261H0v-42h44q7-67 55-114.5T213 49V5h43v44q67 8 114.5 55.5T425 219M235 389q62 0 105.5-43.5T384 240t-43.5-105.5T235 91t-106 43.5T85 240t44 105.5T235 389"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GpsDot(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M234.5 155q35.5 0 60.5 25t25 60t-25 60t-60.5 25t-60.5-25t-25-60t25-60t60.5-25M425 219h44v42h-44q-7 67-54.5 114.5T256 431v44h-43v-44q-66-8-114-55.5T44 261H0v-42h44q7-67 55-114.5T213 49V5h43v44q67 8 114.5 55.5T425 219M235 389q62 0 105.5-43.5T384 240t-43.5-105.5T235 91t-106 43.5T85 240t44 105.5T235 389"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GpsOff(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M425 219h44v42h-43q-4 36-21 68l-32-32q11-28 11-57q0-62-43.5-105.5T235 91q-30 0-57 11l-32-32q32-17 67-21V5h43v44q67 8 114.5 55.5T425 219M43 75l27-27l357 357l-27 27l-44-44q-44 36-100 43v44h-43v-44q-66-8-114-55.5T44 261H0v-42h44q6-56 42-100zm283 283L116 149q-31 40-31 91q0 62 44 105.5T235 389q50 0 91-31"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gradient(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M171 128h42v43h-42zm-43 43h43v42h-43zm85 0h43v42h-43zm43-43h43v43h-43zm-171 0h43v43H85zM341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zM128 320v-43H85v43zm85 0v-43h-42v43zm86 0v-43h-43v43zm42-149V43H43v128h42v42H43v43h42v-43h43v43h43v-43h42v43h43v-43h43v43h42v-43h-42v-42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraduationCap(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m85 217l150 82l149-82v86l-149 81l-150-81zM235 0l234 128v171h-42V151L235 256L0 128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grain(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 192q18 0 30.5 12.5t12.5 30t-12.5 30T128 277t-30.5-12.5t-12.5-30t12.5-30T128 192m-85.5-85q17.5 0 30 12.5t12.5 30t-12.5 30t-30 12.5t-30-12.5t-12.5-30t12.5-30t30-12.5m0 170q17.5 0 30 12.5T85 320t-12.5 30.5t-30 12.5t-30-12.5T0 320t12.5-30.5t30-12.5m256-170q-17.5 0-30-12.5T256 64t12.5-30.5t30-12.5t30 12.5T341 64t-12.5 30.5t-30 12.5m-85 170q17.5 0 30 12.5T256 320t-12.5 30.5t-30 12.5t-30-12.5T171 320t12.5-30.5t30-12.5m85-85q17.5 0 30 12.5t12.5 30t-12.5 30t-30 12.5t-30-12.5t-12.5-30t12.5-30t30-12.5m-85-85q17.5 0 30 12.5t12.5 30t-12.5 30t-30 12.5t-30-12.5t-12.5-30t12.5-30t30-12.5M128 21q18 0 30.5 12.5T171 64t-12.5 30.5T128 107T97.5 94.5T85 64t12.5-30.5T128 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraphicEq(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M85 344V88h43v256zm86 85V3h42v426zM0 259v-86h43v86zm256 85V88h43v256zm85-171h43v86h-43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grid(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 3q18 0 30.5 12.5T427 45v342q0 17-12.5 29.5T384 429H43q-18 0-30.5-12.5T0 387V45q0-17 12.5-29.5T43 3zM128 387v-86H43v86zm0-128v-86H43v86zm0-128V45H43v86zm128 256v-86h-85v86zm0-128v-86h-85v86zm0-128V45h-85v86zm128 256v-86h-85v86zm0-128v-86h-85v86zm0-128V45h-85v86z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridOff(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M171 61h-31L97 19h330q17 0 29.5 12.5T469 61v330l-42-43v-31h-31l-43-42h74v-86h-86v74l-42-43v-31h-31l-43-42h74V61h-86v74l-42-43zm170 0v86h86V61zM27 3l458 458l-27 27l-43-43H85q-17 0-29.5-12.5T43 403V73L0 30zm186 241v31h31zM85 116v31h31zm86 287v-86H85v86zm0-128v-74l-12-12H85v86zm128 128v-74l-12-12h-74v86zm42 0h31l-31-31z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GroupIcon(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M149 155v85h107v-85zM0 5h85v22h299V5h85v86h-21v298h21v86h-85v-22H85v22H0v-86h21V91H0zm85 384v22h299v-22h21V91h-21V69H85v22H64v298zm22-277h192v85h64v171H149v-85h-42zm192 171H192v42h128v-85h-21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GroupWork(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3M128 333q22 0 37.5-15.5T181 280t-15.5-37.5T128 227t-37.5 15.5T75 280t15.5 37.5T128 333m32-202q0 22 15.5 37.5T213 184t38-15.5t16-37.5t-16-38t-38-16t-37.5 16t-15.5 38m139 202q22 0 37.5-15.5T352 280t-15.5-37.5T299 227t-38 15.5t-16 37.5t16 37.5t38 15.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hd(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zM171 256V128h-32v53H96v-53H64v128h32v-43h43v43zm42-128v128h86q8 0 14.5-6.5T320 235v-86q0-8-6.5-14.5T299 128zm32 96v-64h43v64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hdr(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 181q0 19-19 30l19 45h-32l-19-43h-24v43h-32V128h75q13 0 22.5 9.5T384 160zm-32 0v-21h-43v21zM75 171v-43h32v128H75v-53H32v53H0V128h32v43zm138-43q13 0 22.5 9.5T245 160v64q0 13-9.5 22.5T213 256h-74V128zm0 96v-64h-42v64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HdrOff(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M343 272h-8l-24-23V144h75q13 0 22.5 9.5T418 176v21q0 10-5.5 18.5T399 227l19 45h-32l-19-43h-24zm0-96v21h43v-21zm-96 0h-8l-32-32h40q13 0 22.5 9.5T279 176v41l-32-32zm-74-21l258 256l-24 23l-162-162h-72v-73l-32-32v105h-32v-53H66v53H34V144h32v43h43v-43h8L0 27L23 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HdrStrong(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 64q53 0 90.5 37.5T469 192t-37.5 90.5T341 320t-90.5-37.5T213 192t37.5-90.5T341 64M85.5 107q35.5 0 60.5 25t25 60t-25 60t-60.5 25T25 252T0 192t25-60t60.5-25m0 128q17.5 0 30-12.5T128 192t-12.5-30.5t-30-12.5t-30 12.5T43 192t12.5 30.5t30 12.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HdrWeak(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M85.5 107q35.5 0 60.5 25t25 60t-25 60t-60.5 25T25 252T0 192t25-60t60.5-25M341 64q53 0 90.5 37.5T469 192t-37.5 90.5T341 320t-90.5-37.5T213 192t37.5-90.5T341 64m.5 213q35.5 0 60.5-25t25-60t-25-60t-60.5-25t-60.5 25t-25 60t25 60t60.5 25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Headset(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 5q80 0 136 56.5T384 197v150q0 26-18.5 45T320 411h-64V240h85v-43q0-62-43.5-105.5T192 48T86.5 91.5T43 197v43h85v171H64q-27 0-45.5-19T0 347V197q0-79 56-135.5T192 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadsetMic(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 5q80 0 136 56.5T384 197v214q0 26-18.5 45T320 475H192v-43h149v-21h-85V240h85v-43q0-62-43.5-105.5T192 48T86.5 91.5T43 197v43h85v171H64q-27 0-45.5-19T0 347V197q0-79 56-135.5T192 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hearing(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 403q17 0 29.5-12.5T341 360h43q0 35-25 60t-60 25q-19 0-35-7q-41-21-59-76q-4-14-12-22.5T169 318q-41-31-61-67q-23-41-23-83q0-63 43.5-106T235 19t106 43t43 106h-43q0-45-31-76t-75.5-31T159 92t-31 76q0 31 17 63q16 27 50 54q13 10 20 16t16.5 19t14.5 29q13 38 36 50q8 4 17 4M99 32Q43 88 43 168q0 79 56 136l-30 30Q0 265 0 168T69 2zm82 136q0-22 16-37.5t38-15.5t37.5 15.5T288 168t-15.5 37.5T235 221t-38-15.5t-16-37.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Help(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3M235 365v-42h-43v42zm44-165q20-20 20-48q0-35-25-60t-60.5-25T153 92t-25 60h43q0-18 12.5-30.5t30-12.5t30 12.5T256 152t-13 30l-26 27q-25 25-25 60v11h43q0-22 5.5-34.5T260 220z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HelpOutline(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 344v-43h43v43zM213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3m0 384q70.5 0 120.5-50t50-121t-50-121t-120.5-50T93 95T43 216t50 121t120.5 50m0-299q35.5 0 60.5 25t25 60q0 18-10 32.5t-22 23t-22 22t-10 29.5h-43q0-23 10-39.5t22-24t22-18.5t10-25q0-17-12.5-29.5t-30-12.5t-30 12.5T171 173h-43q0-35 25-60t60.5-25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M171 363H64V192H0L213 0l214 192h-64v171H256V235h-85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hospital(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zm-21 235v-86h-85V64h-86v85H64v86h85v85h86v-85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HospitalAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 91v42l-43 128l43 128v43H0v-43l43-128L0 133V91h271l31-86l50 19l-24 67zM277 283v-43h-64v-64h-42v64h-64v43h64v64h42v-64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hotel(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 213q-27 0-45.5-18.5t-18.5-45T82.5 104T128 85t45.5 19t18.5 45.5t-18.5 45T128 213M384 85q35 0 60 25t25 61v192h-42v-64H43v64H0V43h43v192h170V85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hourglass(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 3h256v128l-85 85l85 85v128H0V301l85-85l-85-85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HourglassAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 3h256v128l-85 85l85 85v128H0V301l85-85l-85-85zm213 309l-85-85l-85 85zM43 120h170V45H43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HourglassOutline(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 3h256v128l-85 85l85 85v128H0V301l85-85l-85-85zm213 309l-85-85l-85 85v75h170zm-85-107l85-85V45H43v75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hq(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 21q18 0 30.5 12.5T384 64v256q0 18-12.5 30.5T341 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21zM171 256V128h-32v53H96v-53H64v128h32v-43h43v43zm149-21v-86q0-8-6.5-14.5T299 128h-64q-9 0-15.5 6.5T213 149v86q0 8 6.5 14.5T235 256h16v32h32v-32h16q8 0 14.5-6.5T320 235m-75-11v-64h43v64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Http(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M75 171v-43h32v128H75v-53H32v53H0V128h32v43zm53-11v-32h96v32h-32v96h-32v-96zm117 0v-32h96v32h-32v96h-32v-96zm192-32q13 0 22.5 9.5T469 160v21q0 13-9.5 22.5T437 213h-42v43h-32V128zm0 53v-21h-42v21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Image(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 341q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0h298q18 0 30.5 12.5T384 43zM117 224l-74 96h298l-96-128l-74 96z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 320q0 18-12.5 30.5T427 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21h384q17 0 29.5 12.5T469 64zM160 203l-75 96h299l-96-128l-75 96z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageO(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zm0 341V43H43v298zM234 198l75 101H75l58-76l42 51z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inbox(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H42q-17 0-29.5-12.5T0 341V43q0-18 12.5-30.5T42 0zm0 256V43H42v213h86q0 27 18.5 45.5T192 320t45.5-18.5T256 256zm-64-107l-85 86l-85-86h42V85h86v64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Info(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3M235 323V195h-43v128zm0-171v-43h-43v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoOutline(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 323V195h43v128zM213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3m0 384q70.5 0 120.5-50t50-121t-50-121t-120.5-50T93 95T43 216t50 121t120.5 50M192 152v-43h43v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InputAntenna(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M235 91q62 0 105.5 43.5T384 240h-43q0-44-31-75.5T235 133t-75.5 31.5T128 240H85q0-62 44-105.5T235 91m21 198v70l73 73l-30 30l-64-64l-64 64l-30-30l72-73v-70q-14-6-23-19.5t-9-29.5q0-22 16-37.5t38-15.5t37.5 15.5T288 240q0 35-32 49M235 5q97 0 165.5 69T469 240h-42q0-80-56.5-136t-136-56T99 104T43 240H0q0-97 69-166T235 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InputComposite(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M85 27v85h43v128H0V112h43V27q0-9 6-15.5T64 5t15 6.5T85 27m86 298v-42h128v42q0 21-12 37.5T256 385v90h-43v-90q-19-6-30.5-22.5T171 325M0 325v-42h128v42q0 21-12 37.5T85 385v90H43v-90q-19-6-31-22.5T0 325m427-213h42v128H341V112h43V27q0-9 6.5-15.5t15-6.5t15 6.5T427 27zM256 27v85h43v128H171V112h42V27q0-9 6.5-15.5t15-6.5t15 6.5T256 27m85 298v-42h128v42q0 21-11.5 37.5T427 385v90h-43v-90q-19-6-31-22.5T341 325"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InputHdmi(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M277 109h22v128l-64 128v64H64v-64L0 237V109h21V45q0-17 12.5-29.5T64 3h171q17 0 29.5 12.5T277 45zM64 45v64h43V67h21v42h43V67h21v42h43V45z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InputPower(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M214 85q15 0 28.5 13.5T256 128v117l-75 75v64H75v-64L0 245V128q0-16 13.5-29.5T42 85h1V0h42v85h86V0h42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InputSvideo(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M149 229.5q0 13.5-9 22.5t-22.5 9t-23-9t-9.5-22.5t9.5-23t23-9.5t22.5 9.5t9 23m150-107q0 13.5-9.5 23T267 155h-64q-14 0-23-9.5t-9-23t9-22.5t23-9h64q13 0 22.5 9t9.5 22.5M160 304q13 0 22.5 9.5T192 336t-9.5 22.5T160 368t-22.5-9.5T128 336t9.5-22.5T160 304M235 5q97 0 165.5 69T469 240t-68.5 166T235 475T69 406T0 240T69 74T235 5m-.5 427q79.5 0 136-56.5T427 240t-56.5-135.5t-136-56.5T99 104.5T43 240t56 135.5T234.5 432M352 197q13 0 22.5 9.5t9.5 23t-9.5 22.5t-22.5 9t-22.5-9t-9.5-22.5t9.5-23T352 197m-42.5 107q13.5 0 22.5 9.5t9 22.5t-9 22.5t-22.5 9.5t-23-9.5T277 336t9.5-22.5t23-9.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Instagram(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 99V56q0-11-11-11h-42q-11 0-11 11v43q0 10 11 10h42q11 0 11-10M53 387h320q11 0 11-11V195h-45q2 12 2 21q0 53-37.5 90.5T213 344t-90.5-37.5T85 216q0-11 2-21H43v181q0 11 10 11m160.5-256q-35.5 0-60.5 25t-25 60t25 60t60.5 25t60.5-25t25-60t-25-60t-60.5-25M384 3q18 0 30.5 12.5T427 45v342q0 17-12.5 29.5T384 429H43q-18 0-30.5-12.5T0 387V45q0-17 12.5-29.5T43 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InvertColors(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M291 121q50 50 50 121t-50 120.5T170.5 412T50 362.5T0 242t50-121L171 0zM171 370V61l-91 90q-37 38-37 91t37 90q37 38 91 38"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InvertColorsOff(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m377 397l7 8l-27 27l-58-58q-46 38-107 38q-71 0-121-50q-46-46-49.5-112T59 134L0 75l27-27l59 59l30 30l76 76l134 134zm-185-27V267L90 165q-26 34-26 77q0 53 38 90q37 38 90 38m0-309l-49 48l-30-30l79-79l121 121q38 39 47 92.5T347 313L192 159z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Iridescent(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M31 301V173h299v128zM159 4h43v63h-43zm171 53l31 30l-39 38l-30-30zM202 471h-43v-63h43zm159-83l-31 30l-38-39l30-30zM0 87l30-30l38 38l-30 30zm30 331L0 387l38-38l30 30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Key(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M249 149h220v86h-42v85h-86v-85h-92q-14 37-47 61t-74 24q-53 0-90.5-37.5T0 192t37.5-90.5T128 64q41 0 74 24t47 61m-121 86q18 0 30.5-12.5T171 192t-12.5-30.5T128 149t-30.5 12.5T85 192t12.5 30.5T128 235"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyboard(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 43q18 0 30.5 12.5T427 85v214q0 17-12.5 29.5T384 341H43q-18 0-30.5-12.5T0 299V85q0-17 12.5-29.5T43 43zm-192 64v42h43v-42zm0 64v42h43v-42zm-64-64v42h43v-42zm0 64v42h43v-42zm-21 42v-42H64v42zm0-64v-42H64v42zm192 150v-43H128v43zm0-86v-42h-43v42zm0-64v-42h-43v42zm64 64v-42h-43v42zm0-64v-42h-43v42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyboardHide(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 0q18 0 30.5 12.5T427 43v213q0 18-12.5 30.5T384 299H43q-18 0-30.5-12.5T0 256V43q0-18 12.5-30.5T43 0zM192 64v43h43V64zm0 64v43h43v-43zm-64-64v43h43V64zm0 64v43h43v-43zm-21 43v-43H64v43zm0-64V64H64v43zm192 149v-43H128v43zm0-85v-43h-43v43zm0-64V64h-43v43zm64 64v-43h-43v43zm0-64V64h-43v43zM213 427l-85-86h171z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Label(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M414 207q13 13 13 30.5T414 267L265 417q-13 12-30.5 12T205 417L13 225Q0 212 0 195V45q0-17 12.5-29.5T43 3h149q18 0 30 12zM74.5 109q13.5 0 23-9t9.5-22.5t-9.5-23t-23-9.5T52 54.5t-9 23t9 22.5t22.5 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LabelAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m312 61l93 131l-93 131q-13 18-35 18H43q-18 0-30.5-12.5T0 299V85q0-17 12.5-29.5T43 43h234q22 0 35 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LabelAltOutline(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m312 61l93 131l-93 131q-13 18-35 18H43q-18 0-30.5-12.5T0 299V85q0-17 12.5-29.5T43 43h234q22 0 35 18m-35 238l76-107l-76-107H43v214z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LabelHeart(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M414 207q13 13 13 30.5T414 267L265 417q-13 12-30.5 12T205 417L13 225Q0 212 0 195V45q0-17 12.5-29.5T43 3h149q18 0 30 12zM74.5 109q13.5 0 23-9t9.5-22.5t-9.5-23t-23-9.5T52 54.5t-9 23t9 22.5t22.5 9M326 286q15-16 15-38t-15.5-37.5T288 195t-38 15l-15 16l-16-16q-15-15-38-15q-22 0-37.5 15.5T128 248t16 38l91 91z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Labels(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M26 363q-16-7-22.5-23T3 308l52-125v192zm416-79q7 16 0 32.5T419 340l-157 65q-8 3-16 3q-29 0-39-26L101 126q-4-8-3-17q0-28 26-38L281 6q8-3 17-3q28 0 39 26zM140.5 131q8.5 0 15-6.5t6.5-15t-6.5-15t-15-6.5t-15 6.5t-6.5 15t6.5 15t15 6.5M98 365V230l73 178h-31q-17 0-29.5-12.5T98 365"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lamp(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m54 364l39-39l30 30l-39 39zm159 83v-63h43v63zM64 192v43H0v-43zm235-89q29 17 46.5 46t17.5 64q0 53-37.5 90.5T235 341t-90.5-37.5T107 213q0-35 17-64t47-46V0h128zm106 89h64v43h-64zm-59 163l30-29l39 38l-30 30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Landscape(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m277 64l192 256H0l128-171l96 128l34-25l-61-81z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LanguageCssThree(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M59 30h342l-31 156l-5 25l-24 121l-183 61L0 332l16-80h67l-6 33l95 36l111-36l15-77H24l13-67h274l9-44H46z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LanguageHtmlFive(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m179 334l91-25l13-138H120l-4-45h171l4-45H66l13 135h155l-5 58l-50 14l-50-14l-4-37H80l7 72zM0 6h357l-32 365l-146 51l-147-51z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LanguageJavascript(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h384v384H0zm101 321q15 33 54 33q25 0 39.5-13.5T209 300V176h-36v123q0 23-19 23q-13 0-24-19zm127-4q19 37 66 37q27 0 43.5-13.5T354 304q0-22-11.5-34T306 247l-9-4q-12-5-17-9.5t-5-12.5q0-6 4.5-10.5T292 206q15 0 24 15l27-18q-16-29-51-29q-24 0-38.5 13.5T239 222t11 33t33 21l9 4q10 5 14.5 7t8 6.5T318 304q0 8-6.5 13t-17.5 5q-23 0-36-22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LanguagePython(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M366 112q25 0 43 18t18 43v81q0 25-18 43t-43 18H213q0 6 5 13t10 7h92v36q0 25-18 43t-43 18h-91q-26 0-43.5-18T107 371v-80q0-25 17.5-43t43.5-18h112q25 0 42.5-18t17.5-43v-57zm-92 252q-15 0-15 19q0 15 15 15q7 0 11-4.5t4-10.5q0-19-15-19M61 325q-25 0-43-17.5T0 264v-80q0-26 18-43.5T61 123h152q0-7-4.5-14t-10.5-7h-91V66q0-25 17.5-43T168 5h91q25 0 43 18t18 43v80q0 26-18 43.5T259 207H147q-25 0-43 18t-18 43v57zm91-251q16 0 16-19q0-15-16-15q-15 0-15 15q0 19 15 19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LanguagePythonAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 90q45-29 82-35.5t60 5.5t39 35.5t23 48t8 49.5q3 37-18.5 72.5T136 313t-83-16v120L0 383zm51 38v121q41 25 65.5 21t35-24.5T162 189q0-47-17-68t-41.5-17.5T51 128m248-72q-4 78 0 155q3 21 14.5 30.5t26.5 8t30-6t25-10.5l10-5V73l53 6v207q0 28-8 50.5t-20 36t-27 23t-30.5 13.5t-27.5 6t-20 2h-8l-18-51q35 0 59-8.5t33-20t13.5-23.5t3.5-20l-1-8q-42 16-73.5 17.5T286 296t-25.5-20.5T249 255l-2-10V90z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laptop(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 320h85v43H0v-43h85q-17 0-29.5-12.5T43 277V64q0-18 12.5-30.5T85 21h342q17 0 29.5 12.5T469 64v213q0 18-12.5 30.5T427 320M85 64v213h342V64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LaptopChromebook(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 320h43v43H0v-43h43V0h426zm-170 0v-21h-86v21zm128-64V43H85v213z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LaptopMac(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 320h85q0 18-12.5 30.5T469 363H43q-18 0-30.5-12.5T0 320h85q-17 0-29.5-12.5T43 277V43q0-18 12.5-30.5T85 0h342q17 0 29.5 12.5T469 43v234q0 18-12.5 30.5T427 320M85 43v234h342V43zm171 298q9 0 15-6t6-15t-6-15t-15-6t-15 6t-6 15t6 15t15 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lastfm(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M325 312q-58 0-87-22.5T196 225l-16-49q-11-32-25-48t-44-16q-25 0-42.5 20T51 194q0 35 16 56t42 21q17 0 33-7t23-14l8-7l15 43q-3 3-9 7t-27 11.5t-45 7.5q-52 0-79.5-30T0 196q0-59 28.5-91.5T110 72q49 0 76 20t42 68l16 50q10 30 28.5 46t53.5 16q51 0 51-26q0-23-33-30l-34-8q-56-14-56-65q0-38 24.5-54.5T341 72q78 0 84 63l-49 6q-3-30-38-30t-35 26q0 23 28 29l31 7q65 15 65 71q0 68-102 68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Layers(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m192 356l157-123l35 27l-192 149L0 260l35-27zm0-55L35 179L0 152L192 3l192 149l-35 27z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayersOff(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m380 304l-31-31l26-19l30 30zm-10-101l-51 40L151 75l62-48l192 149zM27 5l400 400l-27 27l-81-81l-106 82L21 284l35-27l157 123l76-59l-31-30l-45 34L56 203l-35-27l69-54L0 32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Leak(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 0q0 27-18.5 45.5T0 64V0zm171 0q0 97-69 166T0 235v-43q80 0 136-56T192 0zm-86 0q0 62-43.5 105.5T0 149v-42q44 0 75.5-31.5T107 0zm0 384q0-97 69-166t166-69v43q-80 0-136 56t-56 136zm171 0q0-27 18.5-45.5T384 320v64zm-85 0q0-62 43.5-105.5T384 235v42q-44 0-75.5 31.5T277 384z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeakOff(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M149 0q0 30-11 57l-34-34q3-11 3-23zM0 27L27 0l357 357l-27 27l-61-61q-19 28-19 61h-42q0-51 31-91l-31-30q-43 52-43 121h-43q0-86 56-152l-53-53Q86 235 0 235v-43q68 0 122-43l-31-31q-40 31-91 31v-42q33 0 61-19zM235 0q0 64-34 120l-31-31q22-42 22-89zm126 280l-34-34q28-11 57-11v42q-12 0-23 3m-97-97q56-34 120-34v43q-47 0-89 22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Library(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 206q81-75 192-75v234q-110 0-192 76q-81-76-192-76V131q111 0 192 75m0-75q-27 0-45.5-19T128 66.5t18.5-45T192 3t45.5 18.5t18.5 45t-18.5 45.5t-45.5 19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M41 192q0 27 19 46.5t47 19.5h85v41h-85q-44 0-75.5-31.5T0 192t31.5-75.5T107 85h85v41h-85q-28 0-47 19.5T41 192m87 21v-42h171v42zM320 85q44 0 75.5 31.5T427 192t-31.5 75.5T320 299h-85v-41h85q27 0 46.5-19.5T386 192t-19.5-46.5T320 126h-85V85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linkedin(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M319 221.5q-8-10.5-30-10.5q-27 0-38 16t-11 45v146q0 5-3 8t-8 3h-76q-4 0-7.5-3t-3.5-8V148q0-4 3.5-7.5t7.5-3.5h74q4 0 6.5 2t3.5 6v5q1 2 1 7q28-27 76-27q53 0 83 27t30 79v182q0 5-3.5 8t-7.5 3h-78q-4 0-7.5-3t-3.5-8V254q0-22-8-32.5M88 91.5Q73 107 51.5 107T15 91.5t-15-37T15 18T51.5 3T88 18t15 36.5t-15 37m13 56.5v270q0 5-3.5 8t-7.5 3H14q-5 0-8-3t-3-8V148q0-4 3-7.5t8-3.5h76q4 0 7.5 3.5t3.5 7.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkedinBox(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M363 365V244q0-31-22-53t-53-22q-15 0-30 8.5T235 199v-26h-64v192h64V252q0-13 9-22.5t22.5-9.5t23 9.5T299 252v113zM96 137q16 0 27.5-11T135 99t-11.5-27.5T96 60T68.5 71.5T57 99t11.5 27T96 137m32 228V173H64v192zM384 3q18 0 30.5 12.5T427 45v342q0 17-12.5 29.5T384 429H43q-18 0-30.5-12.5T0 387V45q0-17 12.5-29.5T43 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 155q17 0 29.5 12.5T341 197v214q0 17-12.5 29.5T299 453H43q-18 0-30.5-12.5T0 411V197q0-17 12.5-29.5T43 155h21v-43q0-44 31.5-75.5T171 5t75 31.5t31 75.5v43zM170.5 347q17.5 0 30-12.5T213 304t-12.5-30.5t-30-12.5t-30 12.5T128 304t12.5 30.5t30 12.5M237 155v-43q0-27-19.5-46.5t-47-19.5T124 65.5T105 112v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOpen(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M170.5 347q-17.5 0-30-12.5T128 304t12.5-30.5t30-12.5t30 12.5T213 304t-12.5 30.5t-30 12.5M299 155q17 0 29.5 12.5T341 197v214q0 17-12.5 29.5T299 453H43q-18 0-30.5-12.5T0 411V197q0-17 12.5-29.5T43 155h194v-43q0-27-19.5-46.5t-47-19.5T124 65.5T105 112H64q0-44 31.5-75.5T171 5t75 31.5t31 75.5v43zm0 256V197H43v214z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOutline(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 155q17 0 29.5 12.5T341 197v214q0 17-12.5 29.5T299 453H43q-18 0-30.5-12.5T0 411V197q0-17 12.5-29.5T43 155h21v-43q0-44 31.5-75.5T171 5t75 31.5t31 75.5v43zM170.5 46Q143 46 124 65.5T105 112h2v43h130v-43q0-27-19.5-46.5t-47-19.5M299 411V197H43v214zm-128.5-64q-17.5 0-30-12.5T128 304t12.5-30.5t30-12.5t30 12.5T213 304t-12.5 30.5t-30 12.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LongArrowDown(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M149 0v302l77-76l30 30l-128 128L0 256l30-30l77 76V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LongArrowLeft(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 171v42H82l76 77l-30 30L0 192L128 64l30 30l-76 77z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LongArrowReturn(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M363 85h42v128H82l76 77l-30 30L0 192L128 64l30 30l-76 77h281z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LongArrowRight(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 171h302l-76-77l30-30l128 128l-128 128l-30-30l76-77H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LongArrowTab(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m226 94l30-30l128 128l-128 128l-30-30l76-77H0v-42h302zm179-30h43v256h-43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LongArrowUp(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M107 384V82l-77 76l-30-30L128 0l128 128l-30 30l-77-76v302z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Looks(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M234.5 149q61.5 0 105.5 44t44 106h-43q0-44-31-75.5T235 192t-75.5 31.5T128 299H85q0-62 44-106t105.5-44m.5-85q97 0 165.5 69T469 299h-42q0-80-56.5-136t-136-56T99 163T43 299H0q0-97 69-166t166-69"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Loupe(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M235 109v86h85v42h-85v86h-43v-86h-85v-42h85v-86zM213 3q88 0 151 62.5T427 216v171q0 17-12.5 29.5T384 429H213q-88 0-150.5-62.5T0 216T62.5 65.5T213 3m.5 384q70.5 0 120.5-50t50-121t-50-121t-120.5-50T93 95T43 216t50 121t120.5 50"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailReply(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M149 128q54 8 96.5 30.5T315 214t43.5 69.5T384 363q-78-109-235-109v87L0 192L149 43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailReplyAll(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m149 107l-85 85l85 85v64L0 192L149 43zm128 21q54 8 96.5 30.5T443 214t43.5 69.5T512 363q-78-109-235-109v87L128 192L277 43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailSend(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 384V235l320-43L0 149V0l448 192z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Male(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 3q18 0 30.5 12.5t12.5 30t-12.5 30T192 88t-30.5-12.5t-12.5-30t12.5-30T192 3m192 149H256v277h-43V301h-42v128h-43V152H0v-43h384z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MaleAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M32 429V269H0V152q0-18 12.5-30.5T43 109h64q17 0 29.5 12.5T149 152v117h-32v160zM74.5 88q-17.5 0-30-12.5T32 45.5t12.5-30T74.5 3t30 12.5t12.5 30t-12.5 30t-30 12.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MaleFemale(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M32 429V269H0V152q0-18 12.5-30.5T43 109h64q17 0 29.5 12.5T149 152v117h-32v160zm267 0h-64V301h-64l54-162q4-14 15.5-22t24.5-8h3q14 0 25 8t15 22l55 162h-64zM74.5 88q-17.5 0-30-12.5T32 45.5t12.5-30T74.5 3t30 12.5t12.5 30t-12.5 30t-30 12.5m192 0q-17.5 0-30-12.5t-12.5-30t12.5-30t30-12.5t30 12.5t12.5 30t-12.5 30t-30 12.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mall(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 112q18 0 30.5 12.5T384 155v256q0 17-12.5 29.5T341 453H43q-18 0-30.5-12.5T0 411V155q0-18 12.5-30.5T43 112h42q0-44 31.5-75.5T192 5t75.5 31.5T299 112zM192 48q-27 0-45.5 18.5T128 112h128q0-27-18.5-45.5T192 48m0 213q44 0 75.5-31t31.5-75h-43q0 26-18.5 45T192 219t-45.5-19t-18.5-45H85q0 44 31.5 75t75.5 31"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Map(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M373 0q11 0 11 11v322q0 8-8 10l-120 41l-128-45l-114 44l-3 1q-11 0-11-11V51q0-8 8-10L128 0l128 45L370 1zM256 341V88L128 43v253z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarkunreadMailbox(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 128q18 0 30.5 12.5T427 171v256q0 17-12.5 29.5T384 469H43q-18 0-30.5-12.5T0 427V171q0-18 12.5-30.5T43 128h42V0h171v85H128v171h43V128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Memory(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 128v128H128V128zm-43 85v-42h-42v42zm171-42h-43v42h43v43h-43v43q0 17-12.5 29.5T299 341h-43v43h-43v-43h-42v43h-43v-43H85q-17 0-29.5-12.5T43 299v-43H0v-43h43v-42H0v-43h43V85q0-17 12.5-29.5T85 43h43V0h43v43h42V0h43v43h43q17 0 29.5 12.5T341 85v43h43zm-85 128V85H85v214z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Menu(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 336v-48h432v48zm0-120v-48h432v48zM0 48h432v48H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mic(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M149.5 259q-26.5 0-45.5-19t-19-45V67q0-27 19-45.5T149.5 3t45 18.5T213 67v128q0 26-18.5 45t-45 19M262 195h37q0 54-37.5 94.5T171 338v70h-43v-70q-53-8-90.5-49T0 195h36q0 46 34 77t79.5 31t79-31t33.5-77"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicOff(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 195q0 36-19 70l-26-27q9-21 9-43zm-85 3L128 71v-4q0-27 18.5-45.5T192 3t45.5 18.5T256 67zM27 24l357 357l-27 27l-89-89q-26 15-55 19v70h-42v-70q-54-8-91-49t-37-94h36q0 46 33.5 77t79.5 31q25 0 49-11l-35-35q-7 2-14 2q-27 0-45.5-19T128 195v-16L0 51z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicOutline(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M149.5 261q-26.5 0-45.5-18.5T85 197V69q0-26 19-45t45.5-19t45 19T213 69v128q0 27-18.5 45.5t-45 18.5M124 67v132q0 11 7.5 18.5t18 7.5t18-7.5T175 199V67q0-10-7.5-17.5t-18-7.5t-18 7.5T124 67m138 130h37q0 54-37.5 95T171 341v70h-43v-70q-53-8-90.5-49T0 197h36q0 47 34 78t79.5 31t79-31t33.5-78"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicSetting(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 496v-43h42v43zm106.5-235q-26.5 0-45.5-18.5T85 197V69q0-26 19-45t45.5-19t45 19T213 69v128q0 27-18.5 45.5t-45 18.5M128 496v-43h43v43zm85 0v-43h43v43zm86-299q0 54-37.5 95T171 341v70h-43v-70q-53-8-90.5-49T0 197h36q0 47 34 78t79.5 31t79-31t33.5-78z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 213H0v-42h299z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircle(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3M320 237v-42H107v42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircleOutline(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M107 195h213v42H107zM213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3m0 384q70.5 0 120.5-50t50-121t-50-121t-120.5-50T93 95T43 216t50 121t120.5 50"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSquare(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zm-42 213v-42H85v42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Money(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M117 169q46 11 73 32t27 61q0 32-20.5 51T143 338v46H79v-46q-34-7-55.5-28T0 256h47q4 45 64 45q31 0 44-12t13-26q0-17-13.5-30T104 211Q4 187 4 123q0-29 21-49.5T79 46V0h64v47q32 8 49.5 30t18.5 51h-47q-2-45-53-45q-27 0-42.5 11T53 123q0 15 14 25.5t50 20.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneyBox(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 299v-22h-43v-42h86v-22h-64q-9 0-15.5-6t-6.5-15v-64q0-9 6.5-15t15.5-6h21V85h43v22h42v42h-85v22h64q9 0 15 6t6 15v64q0 9-6 15t-15 6h-21v22zM384 21q18 0 30.5 12.5T427 64v256q0 18-12.5 30.5T384 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21zm0 299V64H43v256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneyOff(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M180 83q-18 0-32 6l-32-31q15-8 32-12V0h64v47q32 8 49.5 30t19.5 51h-48q-2-45-53-45M27 23l312 312l-27 27l-48-48q-19 18-52 24v46h-64v-46q-33-7-55-28t-23-54h46q5 45 64 45q38 0 52-20l-75-74q-84-25-84-84L0 50z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mood(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3m0 384q70.5 0 120.5-50t50-121t-50-121t-120.5-50T93 95T43 216t50 121t120.5 50M288 195q-13 0-22.5-9.5t-9.5-23t9.5-22.5t22.5-9t22.5 9t9.5 22.5t-9.5 23T288 195m-149.5 0q-13.5 0-22.5-9.5t-9-23t9-22.5t22.5-9t23 9t9.5 22.5t-9.5 23t-23 9.5m75 138q-36.5 0-66.5-20.5T104 259h218q-13 33-42.5 53.5t-66 20.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoodBad(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3m0 384q70.5 0 120.5-50t50-121t-50-121t-120.5-50T93 95T43 216t50 121t120.5 50M288 195q-13 0-22.5-9.5t-9.5-23t9.5-22.5t22.5-9t22.5 9t9.5 22.5t-9.5 23T288 195m-149.5 0q-13.5 0-22.5-9.5t-9-23t9-22.5t22.5-9t23 9t9.5 22.5t-9.5 23t-23 9.5m75 64q36.5 0 66 20.5T322 333H104q13-33 43-53.5t66.5-20.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func More(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M42.5 149q17.5 0 30 12.5T85 192t-12.5 30.5t-30 12.5t-30-12.5T0 192t12.5-30.5t30-12.5m256 0q17.5 0 30 12.5T341 192t-12.5 30.5t-30 12.5t-30-12.5T256 192t12.5-30.5t30-12.5m-128 0q17.5 0 30 12.5T213 192t-12.5 30.5t-30 12.5t-30-12.5T128 192t12.5-30.5t30-12.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreVert(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M42.5 107q-17.5 0-30-12.5T0 64t12.5-30.5t30-12.5t30 12.5T85 64T72.5 94.5t-30 12.5m0 42q17.5 0 30 12.5T85 192t-12.5 30.5t-30 12.5t-30-12.5T0 192t12.5-30.5t30-12.5m0 128q17.5 0 30 12.5T85 320t-12.5 30.5t-30 12.5t-30-12.5T0 320t12.5-30.5t30-12.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mouse(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 7q64 8 106.5 56T341 176H192zM0 304v-85h341v85q0 71-50 121t-120.5 50T50 425T0 304M149 7v169H0q0-65 43-113T149 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Movie(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 21h86v299q0 18-12.5 30.5T384 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21h21l43 86h64l-43-86h43l42 86h64l-42-86h42l43 86h64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MovieAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 0h42v384h-42v-43h-43v43H85v-43H43v43H0V0h43v43h42V0h171v43h43zM85 299v-43H43v43zm0-86v-42H43v42zm0-85V85H43v43zm214 171v-43h-43v43zm0-86v-42h-43v42zm0-85V85h-43v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NfiveSquare(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zm-85 128V85H128v128h85v43h-85v43h85q18 0 30.5-12.5T256 256v-43q0-17-12.5-29.5T213 171h-42v-43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NfourSquare(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zm-85 299V85h-43v86h-42V85h-43v128h85v86z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoneSquare(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zM235 299V85h-86v43h43v171z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NsixSquare(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M171 256v-43h42v43zM341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zm-85 128V85h-85q-18 0-30.5 12.5T128 128v128q0 18 12.5 30.5T171 299h42q18 0 30.5-12.5T256 256v-43q0-17-12.5-29.5T213 171h-42v-43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NthreeSquare(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zm-85 160v-32q0-18-12.5-30.5T213 85h-85v43h85v43h-42v42h42v43h-85v43h85q18 0 30.5-12.5T256 256v-32q0-13-9.5-22.5T224 192q13 0 22.5-9.5T256 160"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NtwoSquare(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zm-85 171v-43q0-18-12.5-30.5T213 85h-85v43h85v43h-42q-18 0-30.5 12.5T128 213v86h128v-43h-85v-43h42q18 0 30.5-12.5T256 171"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Nature(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M171 304v83h128v42H0v-42h128v-84q-53-9-88.5-50.5T4 156Q4 94 47.5 50T153 6t105.5 44T302 156q0 57-37.5 99T171 304"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NaturePeople(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M430 156q0 57-37.5 99T299 304v83h64v42H21V323H0v-86q0-8 6.5-14.5T21 216h64q9 0 15.5 6.5T107 237v86H85v64h171v-84q-53-9-88.5-50.5T132 156q0-62 43.5-106T281 6t105.5 44T430 156M53.5 195q-13.5 0-23-9.5t-9.5-23t9.5-22.5t23-9t22.5 9t9 22.5t-9 23t-22.5 9.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Navigation(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m160 3l160 390l-15 15l-145-64l-145 64l-15-15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NegOne(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 171h171v42H0zm320 149h-43V93l-64 22V79l101-36h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NegTwo(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M278 284h127v36H221v-32l89-97q10-11 19-22q7-8 12-18q4-7 6-14q2-8 2-14q0-9-3-18q-3-8-8-13q-5-7-12.5-10T308 79q-12 0-20 4q-9 4-15 10q-6 8-8 16q-3 9-3 19h-46q1-17 6-32q6-16 18-28t29-19q18-6 40-6q20 0 36 5q17 6 27 15q11 10 17 24t6 31q0 13-4 25q-5 12-12 25q-8 13-17 25q-13 15-23 25zM0 171h171v42H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Network(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 429L427 3v426z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NetworkAlert(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 344V173h43v171zm0 85v-42h43v42zM0 429L427 3v128h-86v298z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NetworkLocked(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M395 197q-40 0-68 28.5T299 293v6q-22 19-22 48v64H0L405 5v193q-9-1-10-1m53 128q9 0 15 6.5t6 15.5v85q0 9-6 15t-15 6H341q-8 0-14.5-6t-6.5-15v-85q0-9 6.5-15.5T341 325v-32q0-22 16-37.5t38-15.5t37.5 15.5T448 293zm-21 0v-32q0-13-9.5-22.5t-23-9.5t-22.5 9.5t-9 22.5v32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NetworkOff(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 5v367L243 189zM80 80l368 368l-27 27l-43-43H0l189-189L53 107z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NetworkOutline(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 106L103 387h281zM427 3v426H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NetworkSetting(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M405 245q-66 0-113 47t-47 113q0 9 2 22H0L427 0l-1 247q-12-2-21-2m79 171l23 17q3 3 1 7l-21 37q-2 4-7 3l-26-11q-8 6-18 10l-4 29q-1 4-6 4h-42q-5 0-6-4l-4-29q-9-3-18-10l-26 11q-5 1-7-3l-21-37q-2-4 1-7l23-17q-1-5-1-10.5t1-10.5l-23-18q-3-3-1-7l21-37q3-3 7-2l26 11q8-6 18-11l4-28q1-4 6-4h42q5 0 6 4l4 28q9 4 18 11l26-11q5-1 7 2l21 37q2 4-1 7l-23 18q1 4 1 10q0 4-1 11m-79 21q13 0 22.5-9t9.5-22.5t-9.5-23T405 373t-22.5 9.5t-9.5 23t9.5 22.5t22.5 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Nfc(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 3q18 0 30.5 12.5T427 45v342q0 17-12.5 29.5T384 429H43q-18 0-30.5-12.5T0 387V45q0-17 12.5-29.5T43 3zm0 384V45H43v342zM341 88v256H85V88h86v43h-43v170h171V131h-64v48q21 12 21 37q0 18-12.5 30.5t-30 12.5t-30-12.5T171 216q0-24 21-37v-48q0-18 12.5-30.5T235 88z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notifications(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M181.5 429q-17.5 0-30-12.5T139 387h85q0 17-12.5 29.5t-30 12.5M320 301l43 43v21H0v-21l43-43V184q0-49 30-86.5T149 49V35q0-14 9.5-23t23-9t22.5 9t9 23v14q47 11 77 48.5t30 86.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationsActive(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M118 36Q85 60 65 96t-22 77H0q2-50 25.5-94T88 6zm286 137q-2-41-22.5-77T328 36l31-30q39 29 62 73t26 94zm-42 11v117l43 43v21H42v-21l43-43V184q0-49 30-86.5T191 49V35q0-14 9.5-23t23-9t22.5 9t9 23v14q47 11 77 48.5t30 86.5M223 429q-17 0-29.5-12.5T181 387h85q0 8-3 16q-9 21-31 25q-4 1-9 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationsAdd(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M150 408h84q0 18-12 30.5T192 451t-30-12.5t-12-30.5m189-89l45 45v23H0v-23l45-45V195q0-52 32-91.5T158 52V37q0-14 10-24t24-10t24 10t10 24v15q49 12 81 51.5t32 91.5zm-62-81v-43h-64v-64h-42v64h-64v43h64v64h42v-64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationsNone(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M181.5 429q-17.5 0-30-12.5T139 387h85q0 17-12.5 29.5t-30 12.5M320 301l43 43v21H0v-21l43-43V184q0-49 30-86.5T149 49V35q0-14 9.5-23t23-9t22.5 9t9 23v14q47 11 77 48.5t30 86.5zm-43 22V184q0-40-28-68t-68-28t-68 28t-28 68v139z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationsOff(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M181.5 429q-17.5 0-30-12.5T139 387h85q0 17-12.5 29.5t-30 12.5M320 184v79L118 61q17-8 31-12V35q0-14 9.5-23t23-9t22.5 9t9 23v14q47 11 77 48.5t30 86.5m-6 181H0v-21l43-43V184q0-38 19-71L0 51l27-27l357 357l-27 27z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationsPaused(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M181.5 429q-17.5 0-30-12.5T139 387h85q0 17-12.5 29.5t-30 12.5M320 301l43 43v21H0v-21l43-43V184q0-49 30-86.5T149 49V35q0-14 9.5-23t23-9t22.5 9t9 23v14q47 11 77 48.5t30 86.5zm-85-132v-38H128v38h60l-60 73v38h107v-38h-60z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Odnoklassniki(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M136 219q-45 0-76.5-32T28 110.5t31.5-76T136 3t76.5 31.5t31.5 76t-31.5 76.5t-76.5 32m0-161q-22 0-37.5 15.5T83 111t15.5 37.5T136 164t37.5-15.5T189 111t-15.5-37.5T136 58m124 174q8 15 1 24.5T232 281q-27 17-75 22l81 81q7 7 7 17.5t-7 17.5l-3 3q-8 7-18 7t-17-7q-12-11-64-64l-63 64q-7 7-17.5 7T38 422l-3-3q-7-7-7-17.5t7-17.5l63-63l18-18q-48-4-76-22q-22-15-29-24.5t1-24.5q5-11 16-13.5t29 8.5q14 11 33.5 17t32.5 6l13 1q49 0 79-24q18-11 29-8.5t16 13.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OpenInBrowser(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 21q18 0 30.5 12.5T384 64v256q0 18-12.5 30.5T341 363h-85v-43h85V107H43v213h85v43H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21zM192 149l85 86h-64v128h-42V235h-64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OpenInNew(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 341V192h43v149q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0h149v43H43v298zM235 0h149v149h-43V73L132 282l-30-30L311 43h-76z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Outlook(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M267 93h145q4 0 9.5 5t5.5 12l-127 85h-4l-29-18zm0 115l27 18q2 1 4 1h3l1-1q-2 1 29-19.5t64-41.5l32-21v153q0 12-6.5 18t-16.5 6H267zm-139-39q13 0 20.5 12.5T156 216t-7.5 34t-21.5 12q-13 0-21-12.5T98 216t8-34t22-13M0 51L251 3v426L0 377zm168 218q16-21 16-54t-15.5-53.5T128 141q-26 0-42 21t-16 56q0 32 16 52t41 20t41-21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Palette(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 0q80 0 136 50t56 121q0 44-31.5 75T277 277h-37q-14 0-23 9.5t-9 22.5q0 12 8 21q8 10 8 22q0 13-9.5 22.5T192 384q-80 0-136-56T0 192T56 56T192 0M74.5 192q13.5 0 23-9.5T107 160t-9.5-22.5t-23-9.5t-22.5 9.5t-9 22.5t9 22.5t22.5 9.5m64-85q13.5 0 23-9.5t9.5-23t-9.5-22.5t-23-9t-22.5 9t-9 22.5t9 23t22.5 9.5m107 0q13.5 0 22.5-9.5t9-23t-9-22.5t-22.5-9t-23 9t-9.5 22.5t9.5 23t23 9.5m64 85q13.5 0 22.5-9.5t9-22.5t-9-22.5t-22.5-9.5t-23 9.5T277 160t9.5 22.5t23 9.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanoramaHorizontal(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 76q-84 25-171 24q-87 0-170-24v232q83-24 170-24t171 24zm31-55q12 0 12 14v314q0 14-12 14q-4 0-7-2q-94-35-195-35T19 361q-4 2-7 2q-12 0-12-14V35q0-14 12-14q3 0 7 2q94 35 194 35q101 0 195-35q3-2 7-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanoramaVertical(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m340 411l2 6q0 12-14 12H13q-13 0-13-12q0-3 1-6q35-95 35-195T1 21q-1-3-1-6Q0 3 13 3h315q13 0 13 12q0 3-1 6q-35 95-35 195q0 101 35 195M54 387h233q-25-84-25-171t25-171H54q25 84 25 171T54 387"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanoramaWideAngle(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 64Q138 64 58 78q-15 57-15 114t15 114q80 14 155.5 14T369 306q15-57 15-114T369 78q-80-14-155.5-14m-.5-43q83 0 170 16l20 3l5 19q19 67 19 133t-19 133l-5 19l-20 3q-87 16-170 16T44 347l-20-3l-5-19Q0 258 0 192T19 59l5-19l20-3q87-16 169-16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Parking(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M149 0q53 0 90.5 37.5T277 128t-37.5 90.5T149 256H85v128H0V0zm5 171q17 0 29.5-12.5T196 128t-12.5-30.5T154 85H85v86z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 341V43h85v298zM171 43h85v298h-85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseCircle(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3M192 301V131h-43v170zm85 0V131h-42v170z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseCircleOutline(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M149 301V131h43v170zM213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3m0 384q70.5 0 120.5-50t50-121t-50-121t-120.5-50T93 95T43 216t50 121t120.5 50m21.5-86V131h42v170z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paypal(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M237 3q79 0 112 39q30 35 20 99q-23 146-175 146h-49q-8 0-14 5t-7 13l-17 106q-1 8-7 13t-14 5H13q-6 0-10-4.5T0 415L62 21q2-8 7.5-13T83 3zm18 144q4-29-8-43q-6-8-18-11.5t-21.5-4T180 88h-11q-11 0-12 11l-17 103h23q17 0 25.5-.5t22-3.5t21-8.5t14-16.5t9.5-26"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaypalAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M225 108q0-34-52-34h-15q-7 0-13 4.5T138 90l-14 60v3q0 5 3.5 8.5t8.5 3.5h12q15 0 28-3t24.5-9t18-17.5T225 108m131 33q0 58-48 93q-47 35-133 35h-13q-7 0-13 4.5t-7 11.5l-16 69q-2 7-9 12.5t-15 5.5H56q-7 0-11.5-4T40 357q0-2 3-14h32q8 0 14.5-5t7.5-12l16-69q2-7 8.5-12t13.5-5h13q85 0 132-35t47-92q0-28-11-44q40 20 40 72m-40-40q0 57-48 93q-47 35-133 35h-13q-7 0-13 4.5t-7 11.5l-16 68q-2 8-8.5 13.5T62 332H16q-7 0-11.5-4T0 317v-4L66 30q1-7 8-12.5T89 12h97q14 0 26.5.5t26.5 3t24.5 6.5t21 11t17 16T312 71.5t4 29.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M77 166q47 93 141 141l47-47q9-10 22-5q36 12 76 12q8 0 14.5 6t6.5 15v75q0 8-6.5 14.5T363 384q-99 0-182.5-48.5t-132-132T0 21q0-8 6.5-14.5T21 0h75q9 0 15 6.5t6 14.5q0 40 12 76q4 13-5 22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneBluetooth(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m250 187l-15-15l59-60l-59-60l15-15l49 49V5h10l61 61l-46 46l46 46l-61 61h-10v-81zm70-141v40l20-20zm0 92v40l20-20zm43 177q8 0 14.5 6t6.5 15v75q0 8-6.5 14.5T363 432q-99 0-182.5-48.5t-132-132T0 69q0-8 6.5-14.5T21 48h75q9 0 15 6.5t6 14.5q0 40 12 76q4 13-5 22l-47 47q47 93 141 141l47-47q9-9 22-5q36 12 76 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneEnd(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 128q-51 0-98 15v66q0 14-12 20q-31 15-57 39q-6 6-15 6t-15-6L6 215q-6-6-6-15t6-15Q111 85 256 85t250 100q6 6 6 15t-6 15l-53 53q-6 6-15 6t-15-6q-25-23-57-39q-12-6-12-19v-66q-47-16-98-16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneForwarded(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M320 219v-64h-85V69h85V5l107 107zm43 96q8 0 14.5 6t6.5 15v75q0 8-6.5 14.5T363 432q-99 0-182.5-48.5t-132-132T0 69q0-8 6.5-14.5T21 48h75q9 0 15 6.5t6 14.5q0 40 12 76q4 13-5 22l-47 47q47 93 141 141l47-47q9-9 22-5q36 12 76 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneInTalk(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M363 267q8 0 14.5 6t6.5 15v75q0 8-6.5 14.5T363 384q-99 0-182.5-48.5t-132-132T0 21q0-8 6.5-14.5T21 0h75q9 0 15 6.5t6 14.5q0 40 12 76q4 13-5 22l-47 47q47 93 141 141l47-47q9-9 22-5q36 12 76 12m-22-75q0-62-43.5-105.5T192 43V0q80 0 136 56t56 136zm-85 0q0-27-18.5-45.5T192 128V85q44 0 75.5 31.5T299 192z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneLocked(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M363 315q8 0 14.5 6t6.5 15v75q0 8-6.5 14.5T363 432q-99 0-182.5-48.5t-132-132T0 69q0-8 6.5-14.5T21 48h75q9 0 15 6.5t6 14.5q0 40 12 76q4 13-5 22l-47 47q47 93 141 141l47-47q9-9 22-5q36 12 76 12m0-246q8 0 14.5 6.5T384 91v85q0 9-6.5 15t-14.5 6H256q-9 0-15-6t-6-15V91q0-9 6-15.5t15-6.5V59q0-22 15.5-38T309 5t38 16t16 38zm-17 0V59q0-15-11-26t-26-11t-25.5 11T273 59v10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneMissed(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M139 53v75h-32V0h128v32h-75l96 96L384 0l21 21l-149 150zm367 239q6 6 6 15t-6 15l-53 53q-6 6-15 6t-15-6q-27-24-57-40q-12-5-12-19v-66q-47-15-98-15t-98 15v66q0 14-12 20q-32 16-57 39q-6 6-15 6t-15-6L6 322q-6-6-6-15t6-15q105-100 250-100t250 100"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneMsg(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M363 267q8 0 14.5 6t6.5 15v75q0 8-6.5 14.5T363 384q-99 0-182.5-48.5t-132-132T0 21q0-8 6.5-14.5T21 0h75q9 0 15 6.5t6 14.5q0 40 12 76q4 13-5 22l-47 47q47 93 141 141l47-47q9-9 22-5q36 12 76 12M192 0h192v149H256l-64 64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhonePaused(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 0v149h-43V0zm64 267q8 0 14.5 6t6.5 15v75q0 8-6.5 14.5T363 384q-99 0-182.5-48.5t-132-132T0 21q0-8 6.5-14.5T21 0h75q9 0 15 6.5t6 14.5q0 40 12 76q4 13-5 22l-47 47q47 93 141 141l47-47q9-9 22-5q36 12 76 12M341 0h43v149h-43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneRing(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M506 316q6 6 6 15t-6 15l-53 53q-6 6-15 6t-15-6q-26-24-57-40q-12-5-12-19v-66q-47-15-98-15t-98 15v66q0 14-12 20q-32 16-57 39q-6 6-15 6t-15-6L6 346q-6-6-6-15t6-15q105-100 250-100t250 100M451 94l-76 75l-30-30l76-76zM277 3v106h-42V3zM137 169Q63 94 61 94l30-31l76 76z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneSetting(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213 128v43h-42v-43zm86 0v43h-43v-43zm64 139q8 0 14.5 6t6.5 15v75q0 8-6.5 14.5T363 384q-99 0-182.5-48.5t-132-132T0 21q0-8 6.5-14.5T21 0h75q9 0 15 6.5t6 14.5q0 40 12 76q4 13-5 22l-47 47q47 93 141 141l47-47q9-9 22-5q36 12 76 12m-22-139h43v43h-43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneSip(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 0v107h-22V0zm-43 43v64h-64V85h43V64h-43V0h64v21h-43v22zm64-43h64v64h-43v43h-21zm43 43V21h-22v22zm0 224q8 0 14.5 6t6.5 15v75q0 8-6.5 14.5T363 384q-99 0-182.5-48.5t-132-132T0 21q0-8 6.5-14.5T21 0h75q9 0 15 6.5t6 14.5q0 40 12 76q4 13-5 22l-47 47q48 93 141 141l47-47q9-9 22-5q36 12 76 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhotoSizeSelectLarge(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 256h42v43h-42zm0-85h42v42h-42zm42 170q0 16-13 29.5T427 384v-43zM256 0h43v43h-43zm171 85h42v43h-42zm0-85q16 0 29 13.5T469 43h-42zM0 85h43v43H0zM341 0h43v43h-43zm0 341h43v43h-43zM43 0v43H0q0-16 13.5-29.5T43 0m128 0h42v43h-42zM85 0h43v43H85zM0 171h299v213H43q-18 0-30.5-12.5T0 341zm43 170h213l-68-91l-54 69l-38-46z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhotoSizeSelectSmall(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 256v43h-42v-43zm0-85v42h-42v-42zm0 170q0 16-13 29.5T427 384v-43zM299 0v43h-43V0zm170 85v43h-42V85zM427 0q16 0 29 13.5T469 43h-42zM43 384q-18 0-30.5-12.5T0 341v-85h213v128zm0-299v43H0V85zm256 256v43h-43v-43zM384 0v43h-43V0zm0 341v43h-43v-43zM43 0v43H0q0-16 13.5-29.5T43 0m0 171v42H0v-42zM213 0v43h-42V0zm-85 0v43H85V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PictureInPicture(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 85v128H213V85zm43-85q17 0 29.5 12.5T469 43v298q0 18-12.5 30.5T427 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zm0 342V42H43v300z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pin(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M149 3q62 0 106 43.5T299 152q0 31-15.5 71.5t-37.5 75t-44 65t-37 48.5l-16 17q-6-6-16-18t-35.5-46.5t-45.5-67T16 224T0 152Q0 90 43.5 46.5T149 3m0 202q22 0 38-15.5t16-37.5t-16-37.5T149 99t-37.5 15.5T96 152t15.5 37.5T149 205"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinAccount(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 3q18 0 30.5 12.5T384 45v299q0 18-12.5 30.5T341 387h-85l-64 64l-64-64H43q-18 0-30.5-12.5T0 344V45q0-17 12.5-29.5T43 3zM192 73q-24 0-41 17t-17 41t17 40.5t41 16.5t41-16.5t17-40.5t-17-41t-41-17m128 228v-19q0-20-23.5-35.5t-52.5-23t-52-7.5t-52 7.5t-52.5 23T64 282v19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinAssistant(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 3q18 0 30.5 12.5T384 45v299q0 18-12.5 30.5T341 387h-85l-64 64l-64-64H43q-18 0-30.5-12.5T0 344V45q0-17 12.5-29.5T43 3zM232 235l88-40l-88-40l-40-88l-40 88l-88 40l88 40l40 88z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinDrop(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M277 131q0 27-13 61t-32 63t-37.5 55t-31.5 40l-14 15q-5-5-13.5-15T105 311t-39-56.5t-31-62T21 131q0-53 37.5-90.5T149 3t90.5 37.5T277 131m-170-.5q0 17.5 12.5 30t30 12.5t30-12.5t12.5-30t-12.5-30t-30-12.5t-30 12.5t-12.5 30M0 387h299v42H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinHelp(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 3q18 0 30.5 12.5T384 45v299q0 18-12.5 30.5T341 387h-85l-64 64l-64-64H43q-18 0-30.5-12.5T0 344V45q0-17 12.5-29.5T43 3zM213 344v-43h-42v43zm44-165q20-20 20-48q0-36-25-61t-60-25t-60 25t-25 61h42q0-18 12.5-30.5T192 88t30.5 12.5T235 131q0 17-13 30l-26 27q-25 25-25 60v11h42q0-22 6-34.5t19-26.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinOff(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213 99q-23 0-39 18l-68-68Q150 3 213 3q62 0 106 43.5T363 152q0 48-37 117l-77-78q18-16 18-39q0-22-16-37.5T213 99m94 204l77 78l-27 27l-72-71q-16 23-34 46.5T223 418l-10 11q-6-6-16-18t-35.5-46.5t-45.5-67T80 224t-16-72q0-16 4-33L0 51l27-27l178 178l3 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pinterest(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M99 166q0-27 14-46t34-19q17 0 25.5 11t8.5 27q0 10-3 25q-4 14-10 34q-6 19-9 31q-5 20 7.5 34.5T199 278q35 0 57.5-39.5T279 143q0-43-27.5-70T174 46q-56 0-90.5 35.5T49 167q0 29 17 50q6 6 4 14q-2 5-5 20q-2 5-5.5 6.5t-7.5.5q-26-11-39-37T0 161q0-22 7-44t22-42.5T65 38t51-25.5T181 3t65.5 12t51 32t32 46.5T341 148q0 75-38 124t-98 49q-20 0-37.5-9T143 290q-15 58-18 69q-8 30-36 70H72q-6-51 2-84l33-138q-8-17-8-41"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinterestBox(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M235 306q53 0 82-35t29-82q0-52-39-89.5T213.5 62T120 99.5T81 189q0 34 18 63q6 11 18 11q9 0 15.5-6.5T139 242q0-5-4-11q-11-20-11-42q0-35 26-59.5t63-24.5t63.5 24.5T303 189q0 30-16.5 51.5T235 262q-12 0-20-8.5t-8-20.5q0-9 9.5-28.5T226 169q0-28-31-28q-14 0-24.5 11.5T160 189q0 8 1 16t2 12l1 3l-39 119l-1 4v3q0 10 6.5 17t16.5 7q14 0 20-12l1 1l1-4l20-69q19 20 46 20M384 3q18 0 30.5 12.5T427 45v342q0 17-12.5 29.5T384 429H43q-18 0-30.5-12.5T0 387V45q0-17 12.5-29.5T43 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pizza(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 3q56 0 105.5 22.5T384 88L192 429L0 88q36-40 86-62.5T192 3M85 109.5q0 17.5 12.5 30t30 12.5t30-12.5t12.5-30t-12.5-30t-30-12.5t-30 12.5t-12.5 30M191.5 280q17.5 0 30-12.5t12.5-30t-12.5-30t-30-12.5t-30 12.5t-12.5 30t12.5 30t30 12.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plaster(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m336 216l86 85q6 7 6 15.5t-6 14.5l-93 93q-6 6-15 6t-15-6l-85-85l-85 85q-6 6-15 6t-15-6L6 331q-6-6-6-14.5T6 301l85-85l-85-84q-6-7-6-15.5T6 101L99 9q6-6 14.5-6T129 9l85 85l85-85q6-6 15-6t15 6l92 92q7 7 7 15.5t-7 15.5zm-122-64q-9 0-15 6.5t-6 15t6 15t15 6.5t15.5-6.5t6.5-15t-6.5-15T214 152m-100 42l77-78l-77-77l-78 78zm57.5 43q8.5 0 15-6t6.5-15t-6.5-15t-15-6t-15 6t-6.5 15t6.5 15t15 6m42.5 43q9 0 15.5-6.5t6.5-15t-6.5-15T214 237t-15 6.5t-6 15t6 15t15 6.5m43-85q-9 0-15 6t-6 15t6 15t15 6t15-6t6-15t-6-15t-15-6m57 199l77-78l-77-77l-78 78z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 43l235 149L0 341z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayCircle(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3M171 312l128-96l-128-96z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayCircleOutline(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M171 312V120l128 96zM213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3m0 384q70.5 0 120.5-50t50-121t-50-121t-120.5-50T93 95T43 216t50 121t120.5 50"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayForWork(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M107 43h42v119h75l-96 96l-96-96h75zM0 235h43q0 35 25 60t60 25t60-25t25-60h43q0 53-37.5 90.5T128 363t-90.5-37.5T0 235"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaylistAudio(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 64v43H0V64zm0 85v43H0v-43zM0 277v-42h171v42zM299 64h106v43h-64v192q0 26-18.5 45t-45 19t-45.5-19t-19-45.5t19-45t45-18.5q11 0 22 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaylistPlus(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 149v43H0v-43zm0-85v43H0V64zm85 171h86v42h-86v86h-42v-86h-86v-42h86v-86h42zM0 277v-42h171v42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Playstation(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M426 263q2 9-6.5 18T392 296l-2 1l-64-20l17-6q21-7 21-13q-2-10-37-4l-36 12l-61 21v22l96-32l64 20l-99 34l-61 21v-1v1l-69-22v-39v19q-40 14-84 6q-3 0-11-1.5t-12-2t-11-1.5t-11.5-2.5t-10-3t-8.5-3t-6.5-4t-5-5T0 288q-2-25 34-37l59 18l-15 6q-15 5-6 13q9 9 25 4l64-22v-44l-27-8l27-9V32l88 23q91 24 90 95q-1 90-82 67V100q0-6-7-9t-13.5-1t-6.5 9v148l6-2q58-20 104-17q80 6 86 35M34 251h2l98-33l27 8v19l-68 24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 213H171v128h-43V213H0v-42h128V43h43v128h128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircle(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3M320 237v-42h-85v-86h-43v86h-85v42h85v86h43v-86z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircleO(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M235 109v86h85v42h-85v86h-43v-86h-85v-42h85v-86zM213 3q88 0 151 62.5T427 216t-63 150.5T213 429T62.5 366.5T0 216T62.5 65.5T213 3m.5 384q70.5 0 120.5-50t50-121t-50-121t-120.5-50T93 95T43 216t50 121t120.5 50"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircleOduplicate(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 107v64h64v42h-64v64h-42v-64h-64v-42h64v-64zM43 192q0 44 23.5 80.5T128 327v46q-56-20-92-69.5T0 192T36 80.5T128 11v46q-38 18-61.5 54.5T43 192M320 0q79 0 135.5 56.5T512 192t-56.5 135.5T320 384t-135.5-56.5T128 192t56.5-135.5T320 0m0 341q62 0 105.5-43.5T469 192T425.5 86.5T320 43T214.5 86.5T171 192t43.5 105.5T320 341"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusOne(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 85v86h85v42h-85v86H85v-86H0v-42h85V85zm213 235h-42V93l-64 22V79l100-36h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSquare(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zm-42 213v-42h-86V85h-42v86H85v42h86v86h42v-86z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusTwo(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M300 284h127v36H243v-32l89-97q10-11 18-22q7-8 12-18q4-7 6-14q2-8 2-14q0-9-3-18q-3-8-8-13q-5-7-12-10t-17-3q-12 0-21 4t-14 10q-6 8-9 16q-2 9-3 19h-45q0-17 6-32q6-16 17.5-28T291 49q17-6 39-6q20 0 37 5q16 6 27 15q10 10 16 24t6 31q0 13-4 25t-12 25q-7 13-17 25q-13 15-22 25zM128 85v86h85v42h-85v86H85v-86H0v-42h85V85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pocket(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 52v112q0 39-10 69q-18 60-68 102q-53 44-121 48q-70 5-129-32q-54-35-80-93q-15-33-18-66q-1-18-1-75V42q0-14 7.5-25T28 2q8-2 16-2h169q25 0 75.5.5T364 1q27 0 35 2q14 4 22 17q6 9 6 32m-85 97q5-15-6-27q-10-13-27-10q-5 0-9.5 3t-7 5t-8 7.5L278 134q-56 55-64 62q-2-1-56-53q-7-7-15-14q-11-11-14-13q-13-9-27-2q-15 6-17.5 21.5T93 162q1 0 58 56l28 26q1 2 5.5 6.5t7 6.5t7 5t8.5 4q15 3 27-8q4-4 9-8.5t11-10.5l9-9q52-50 58-55l5.5-5.5l6.5-6.5l5-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Polymer(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m395 21l96 171l-96 171h-86l96-171l-55-99l-169 270H96L0 192L96 21h85L85 192l56 99L309 21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PortableWifi(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 171q17.5 0 30 12.5t12.5 30t-12.5 30t-30 12.5t-30-12.5t-12.5-30t12.5-30t30-12.5M341 213q0 35-17 64.5T277 324l-21-37q43-25 43-74q0-35-25-60t-60.5-25t-60.5 25t-25 60q0 49 43 74l-22 37q-29-17-46.5-46.5T85 213q0-53 37.5-90.5T213 85t90.5 37.5T341 213M213.5 0q88.5 0 151 62.5T427 213q0 59-29 108t-78 77l-21-37q39-23 62-62t23-86q0-70-50-120T213.5 43T93 93T43 213q0 47 23 86t62 62l-22 37q-48-28-77-77T0 213q0-88 62.5-150.5T213.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PortableWifiChanges(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M364 65q63 63 63 151t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213 3h22v176q21 12 21 37q0 18-12.5 30.5t-30 12.5t-30-12.5T171 216q0-24 21-37v-45q-28 7-46 30t-18 52q0 35 25 60t60.5 25t60.5-25t25-60t-25-60l30-30q37 37 37 90t-37.5 90.5T213 344t-90.5-37.5T85 216q0-47 30.5-82.5T192 90V47q-64 8-106.5 56T43 216q0 71 50 121t120.5 50T334 337t50-121q0-70-50-121z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PortableWifiOff(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m332 264l-34-35q1-7 1-13q0-35-25-60t-61-25q-4 0-13 1l-34-35q23-9 47-9q53 0 90.5 37.5T341 216q0 25-9 48M213 45q-42 0-80 20l-31-31Q154 3 213 3q89 0 151.5 62.5T427 216q0 60-32 111l-31-31q20-38 20-80q0-71-50-121T213 45M27 13l21 22l357 357l-27 27l-160-161l-5 1q-17 0-29.5-12.5T171 216v-4l-34-34q-9 19-9 38q0 49 43 74l-22 37q-29-17-46.5-46.5T85 216q0-38 21-69l-31-31q-32 44-32 100q0 47 23 86t62 62l-22 37q-48-29-77-78T0 216q0-73 45-131L0 40z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Power(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213 0v213h-42V0zm103 46q68 58 68 146q0 80-56 136t-136 56t-136-56T0 192q0-88 68-146l30 30q-55 45-55 116q0 62 43.5 105.5T192 341t105.5-43.5T341 192q0-71-55-115z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PowerInput(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 128h405v43H0zm0 128v-43h107v43zm149 0v-43h107v43zm150 0v-43h106v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PowerSetting(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 472v-43h43v43zm85 0v-43h43v43zM192 3v213h-43V3zm76 52q34 23 53.5 60t19.5 80q0 70-50 120t-120.5 50T50 315T0 195q0-43 19.5-80T73 55l31 30q-28 18-44.5 47T43 195q0 53 37.5 90.5T171 323t90.5-37.5T299 195q0-34-17-63t-45-46zm-33 417v-43h42v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PresentToAll(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 0q17 0 29.5 12.5T469 43v298q0 18-12.5 30.5T427 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zm0 342V42H43v300zM192 192h-43l86-85l85 85h-43v85h-85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Print(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M363 107q26 0 45 18.5t19 45.5v128h-86v85H85v-85H0V171q0-27 18.5-45.5T64 107zm-64 234V235H128v106zm63.5-149q8.5 0 15-6.5t6.5-15t-6.5-15t-15-6.5t-15 6.5t-6.5 15t6.5 15t15 6.5M341 0v85H85V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PuzzlePiece(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M395 219q22 0 37.5 15.5T448 272t-15.5 37.5T395 325h-32v86q0 17-12.5 29.5T320 453h-81v-32q0-24-17-40.5T181 364t-40.5 16.5T124 421v32H43q-18 0-30.5-12.5T0 411v-81h32q24 0 41-17t17-41t-17-41t-41-17H0v-81q0-17 12.5-29.5T43 91h85V59q0-22 15.5-38T181 5t38 16t16 38v32h85q18 0 30.5 12.5T363 133v86z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Quote(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21 299l43-86H0V85h128v128l-43 86zm171 0l43-86h-64V85h128v128l-43 86z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Radio(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M26 115L296 5l15 36l-177 71h250q18 0 30.5 12.5T427 155v256q0 17-12.5 29.5T384 453H43q-18 0-30.5-12.5T0 411V155q0-14 7.5-24.5T26 115m80.5 296q26.5 0 45.5-19t19-45.5t-19-45t-45.5-18.5t-45 18.5t-18.5 45T61.5 392t45 19M384 240v-85H43v85h256v-43h42v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Railway(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 315V91q0-27 12.5-45t38-26.5t53-11.5t67-3t67 3t53 11.5t38 26.5T341 91v224q0 31-21.5 52.5T267 389l32 32v11H43v-11l32-32q-31 0-53-21.5T0 315m170.5 32q17.5 0 30-12.5T213 304t-12.5-30.5t-30-12.5t-30 12.5T128 304t12.5 30.5t30 12.5M299 197V91H43v106z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reader(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 192h149v32H256zm0-21h149zm0 106h149zM427 21q17 0 29.5 12.5T469 64v277q0 18-12.5 30.5T427 384H43q-18 0-30.5-12.5T0 341V64q0-18 12.5-30.5T43 21zm0 320V64H235v277z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Receipt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M320 323v-43H64v43zm0-86v-42H64v42zm0-85v-43H64v43zM0 429V3l32 32L64 3l32 32l32-32l32 32l32-32l32 32l32-32l32 32l32-32l32 32l32-32v426l-32-32l-32 32l-32-32l-32 32l-32-32l-32 32l-32-32l-32 32l-32-32l-32 32l-32-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reddit(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 189q0 28-27 39q2 9 2 19q0 51-55.5 87.5t-134 36.5t-134-36.5T23 247q0-10 2-20q-25-11-25-38q0-18 12.5-30.5T42 146q19 0 32 15q52-36 129-39l35-104q3-7 10-5l83 20q1 0 3 1q8-20 30-20q13 0 23 10t10 23.5T387 71t-23 10q-14 0-23.5-9.5T331 48q-2 1-3 0l-77-18l-31 92q79 2 132 40q13-16 33-16q17 0 29.5 12.5T427 189m-311 33.5q0 12.5 9 21.5t21.5 9t21.5-9t9-21.5t-9-22t-21.5-9.5t-21.5 9.5t-9 22M282 307q4-3 .5-6.5t-7.5-.5q-18 19-62 19t-62-19q-3-3-6.5.5t-.5 6.5q21 22 70 22q47 0 68-22m-1.5-54q12.5 0 21.5-9t9-21.5t-9-22t-21.5-9.5t-22 9.5t-9.5 22t9.5 21.5t22 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Redo(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m360 162l76-77v192H244l78-77q-48-40-110-40q-56 0-100.5 33T50 277L0 261q22-68 80.5-111T212 107q84 0 148 55"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Refresh(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M171 112q-53 0-90.5 37.5T43 240q0 32 15 60l-32 31Q0 289 0 240q0-71 50-121t121-50V5l85 86l-85 85zm144 37q26 42 26 91q0 71-50 121t-120 50v64l-86-86l86-85v64q53 0 90.5-37.5T299 240q0-31-15-60z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RefreshAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m384 107l85 85h-64q0 71-50 121t-120 50q-49 0-91-27l31-31q27 15 60 15q53 0 90.5-37.5T363 192h-64zm-277 85h64l-86 85l-85-85h64q0-71 50-121t121-50q49 0 91 27l-32 31q-27-15-59-15q-53 0-90.5 37.5T107 192"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RefreshSync(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M171 69q70 0 120 50t50 121q0 49-26 91l-31-31q15-28 15-60q0-53-37.5-90.5T171 112v64L85 91l86-86zm0 299v-64l85 85l-85 86v-64q-71 0-121-50T0 240q0-49 26-91l32 31q-15 28-15 60q0 53 37.5 90.5T171 368"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RefreshSyncAlert(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 192q0-59 36-105t92-60v44q-38 14-61.5 47T43 192q0 53 37 90l48-47v128H0l50-51Q0 262 0 192m171 107v-43h42v43zM384 21l-50 51q50 50 50 120q0 59-36 105t-92 60v-44q38-14 61.5-47t23.5-74q0-53-37-90l-48 47V21zM171 213V85h42v128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RefreshSyncOff(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M152 71q-5 2-16 8l-31-32q22-14 47-20zM0 51l27-27l335 336l-27 27l-50-50q-22 14-48 20v-44q7-3 17-8L81 133q-14 28-14 59q0 53 38 90l47-47v128H24l51-51q-51-50-51-120q0-49 26-90zm366-30l-51 51q51 50 51 120q0 49-26 90l-32-31q15-28 15-59q0-53-38-90l-47 47V21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RemoteControl(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M230 192q9 0 15 6.5t6 14.5v256q0 9-6 15.5t-15 6.5H102q-9 0-15-6.5T81 469V213q0-8 6-14.5t15-6.5zm-64 128q18 0 30.5-12.5t12.5-30t-12.5-30T166 235t-30.5 12.5t-12.5 30t12.5 30T166 320M60 129q44-44 106-44t106 44l-31 30q-31-31-75-31t-76 31zM166 0q98 0 166 69l-30 30q-56-56-136-56q-79 0-136 56L0 69Q69 0 166 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RemoteControlAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M59 135q44-44 106.5-44T272 135l-31 30q-31-31-75.5-31T89 165zM165.5 5Q262 5 331 74l-30 30q-56-56-135.5-56T30 104L0 74Q69 5 165.5 5M226 198q10 0 17.5 7t6.5 17v207q0 10-7 17t-17 7H104q-10 0-17-7t-7-17V222q0-10 7-17.5t17-7.5zm3 213V240H101v171z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Repeat(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M85 109v86H43V67h256V3l85 85l-85 85v-64zm214 214v-86h42v128H85v64L0 344l85-85v64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RepeatOne(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M85 109v86H43V67h256V3l85 85l-85 85v-64zm214 214v-86h42v128H85v64L0 344l85-85v64zm-86-43h-32v-85h-32v-22l43-21h21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Replay(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M171 91q70 0 120 50t50 120.5T291 382t-120.5 50T50 382T0 261h43q0 53 37.5 90.5T171 389t90.5-37.5T299 261t-37.5-90.5T171 133v86L64 112L171 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReplayFive(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m142 277l5-46h51v15h-36l-2 20q6-4 13-4q13 0 20.5 8t7.5 23q0 8-4 15t-10.5 11t-16.5 4q-8 0-15-3.5t-11-9.5t-4-13h18q0 5 3.5 8t8.5 3q6 0 9.5-4t3.5-12t-4-12t-11-4q-6 0-10 3l-2 2zm29-189q70 0 120 50t50 120.5T291 379t-120.5 50T50 379T0 259h43q0 52 37.5 90t90 38t90.5-38t38-90.5t-38-90t-90-37.5v85L64 109L171 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReplayTen(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M239 284q0 20-8 30t-23 10t-23-10t-8-29v-17q0-19 8-29t23-10t23 10t8 28zm-18-18q0-12-3-17t-10-5t-10 5t-3 15v23q0 11 3 16.5t10 5.5t10-5t3-16zm-74 57h-19v-71l-22 7v-15l39-14h2zm24-235q70 0 120 50t50 120.5T291 379t-120.5 50T50 379T0 259h43q0 52 37.5 90t90 38t90.5-38t38-90.5t-38-90t-90-37.5v85L64 109L171 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReplayThirty(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M239 285q0 20-8 30t-23.5 10t-23.5-10t-8-29v-17q0-20 8-30t23.5-10t23.5 10t8 29zm-18-19q0-11-3.5-16.5t-10-5.5t-9.5 5t-3 16v23q0 11 3 16.5t10 5.5t10-5t3-16zm-101 3h10q7 0 10-3.5t3-9.5t-3-9t-9-3t-9.5 3t-3.5 8h-18q0-8 4-13.5t11-9t15-3.5q15 0 23.5 7t8.5 20q0 6-4 11.5t-10 8.5q8 3 11.5 8.5T163 298q0 12-9 19.5t-24 7.5q-14 0-23-7t-9-20h19q0 6 4 9t10 3t10-3.5t4-8.5q0-14-16-14h-9zm51-181q70 0 120 50t50 120.5T291 379t-120.5 50T50 379T0 259h43q0 52 37.5 90t90 38t90.5-38t38-90.5t-38-90t-90-37.5v85L64 109L171 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Roller(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 45h64v171H192v192q0 9-6.5 15t-14.5 6h-43q-9 0-15-6t-6-15V173h213V88h-21v21q0 9-6.5 15.5T277 131H21q-8 0-14.5-6.5T0 109V24q0-9 6.5-15T21 3h256q9 0 15.5 6t6.5 15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RotateCcw(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m138 126l139 138l-139 139L0 264zM60 264l78 78l78-78l-78-78zm334-133q57 56 57 135.5T394 402q-56 56-135 56q-49 0-93-24l32-31q29 13 61 13q62 0 105.5-44T408 266.5T364.5 161T259 117v69l-91-90l91-90v69q79 0 135 56"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RotateCw(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m312 126l139 138l-139 139l-138-139zm78 138l-78-78l-78 78l78 78zM56 131q56-56 136-56V6l90 90l-90 90v-69q-62 0-105.5 44T43 266.5T86.5 372T192 416q31 0 60-13l32 31q-43 24-92 24q-80 0-136-56T0 266.5T56 131"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RotateLeft(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M65 166q-17 24-22 53H0q6-46 35-83zm-22 95q5 28 22 53l-30 30Q6 307 0 261zm22 114l30-31q24 17 53 22v43q-46-5-83-34M191 71q63 8 106 56t43 113t-43 113t-106 56v-43q45-8 75.5-43.5T297 240t-30.5-82.5T191 114v83l-98-95l98-97z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RotateRight(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m246 102l-97 95v-83q-45 8-75.5 43.5T43 240t30.5 82.5T149 366v43q-63-8-106-56T0 240t43-113t106-56V5zm94 117h-43q-5-29-22-53l30-30q29 37 35 83M192 366q28-5 52-22l31 31q-37 28-83 34zm83-52q17-24 22-53h43q-6 46-35 83z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Router(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M367 62q-40-36-90-36t-89 36l-17-17Q215 0 277 0t107 45zm-19 17l-17 17q-22-21-54-21t-53 21l-17-17q30-30 70.5-30T348 79m-7 134q18 0 30.5 12.5T384 256v85q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341v-85q0-18 12.5-30.5T43 213h213v-85h43v85zM107 320v-43H64v43zm74 0v-43h-42v43zm75 0v-43h-43v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rss(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M47 269q19 0 33 13.5t14 33T80 349t-33 14t-33-14t-14-33.5t14-33T47 269M0 146q90 0 153.5 63.5T217 363h-62q0-64-45.5-109.5T0 208zM0 21q93 0 171.5 46t124 124.5T341 363h-62q0-116-81.5-198T0 83z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ruler(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 64q17 0 29.5 12.5T469 107v170q0 18-12.5 30.5T427 320H43q-18 0-30.5-12.5T0 277V107q0-18 12.5-30.5T43 64zm0 213V107h-43v85h-43v-85h-42v85h-43v-85h-43v85h-42v-85h-43v85H85v-85H43v170z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Run(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M226.5 85Q209 85 196 72.5t-13-30t13-30T226.5 0t30 12.5t12.5 30t-12.5 30t-30 12.5M149 381L0 352l9-43l104 21l34-173l-38 15v73H66V145l111-47q3 0 8.5-1t8.5-1q22 0 36 21l22 34q13 23 37.5 37t53.5 14v43q-71 0-117-53l-13 64l45 42v160h-43V330l-44-42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Satellite(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zM43 42v65q26 0 45-19t19-46zm0 150q62 0 105.5-44T192 42h-43q0 45-31 76t-75 31zm0 128h298l-96-128l-74 96l-54-64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scanner(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M358 164q11 3 18.5 14.5T384 203v117q0 18-12.5 30.5T341 363H43q-18 0-30.5-12.5T0 320v-85q0-18 12.5-30.5T43 192h268L11 83l15-40zM85 299v-43H43v43zm256 0v-43H128v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scissors(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m163 123l264 264v21h-64L213 259l-50 50q8 17 8 35q0 35-25 60t-60.5 25T25 404T0 344t25-60t60-25q19 0 35 7l51-50l-51-50q-16 7-35 7q-35 0-60-25T0 88t25-60T85.5 3T146 28t25 60q0 18-8 35m-77.5 8q17.5 0 30-12.5T128 88t-12.5-30.5t-30-12.5t-30 12.5T43 88t12.5 30.5t30 12.5m0 256q17.5 0 30-12.5T128 344t-12.5-30.5t-30-12.5t-30 12.5T43 344t12.5 30.5t30 12.5m128-160q10.5 0 10.5-11t-10.5-11t-10.5 11t10.5 11M363 24h64v21L277 195l-42-43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScreenRotation(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m351 54l-29 28l-81-81l14-1q100 0 173.5 68T510 235h-32q-6-60-40.5-108T351 54M217 37l257 257q9 9 9 22.5t-9 22.5L338 475q-9 9-22.5 9t-22.5-9L36 218q-9-9-9-22.5t9-22.5L172 37q9-9 22.5-9t22.5 9m98 415l136-136L195 60L59 196zm-156 6l29-28l81 81l-14 1q-100 0-173.5-68T0 277h32q6 60 40 108t87 73"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScreenRotationLock(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M475 272q9 10 9 23t-9 23L339 453q-9 10-22.5 10T294 453L37 197q-9-9-9-22.5t9-22.5L173 16q9-9 22.5-9t22.5 9l53 52l-31 30l-44-44L75 174l241 242l121-121l-47-47l30-30zM159 437l29-28l81 81l-14 1q-100 0-173.5-68T0 256h32q6 60 40 108t87 73m161-245q-9 0-15-6.5t-6-14.5V85q0-8 6-14.5t15-6.5V53q0-22 15.5-37.5T373 0t38 15.5T427 53v11q8 0 14.5 6.5T448 85v86q0 8-6.5 14.5T427 192zm17-139v11h73V53q0-15-11-25.5T373 17t-25.5 10.5T337 53"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m267 235l106 106l-32 32l-106-106v-17l-6-6q-39 33-90 33q-58 0-98.5-40.5T0 138.5t40.5-98t98-40.5t98 40.5T277 139q0 51-33 90l6 6zm-128 0q40 0 68-28t28-68t-28-68t-68-28t-68 28t-28 68t28 68t68 28"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchFor(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m320 235l107 106l-32 32l-107-107v-16l-6-6q-39 33-90 33q-38 0-70-19l31-31q19 8 39 8q40 0 68-28.5t28-68T260 71t-68-28t-68 28t-28 68h75l-89 85l-82-85h54q0-57 40.5-98T192 0q58 0 98.5 40.5T331 139q0 51-34 90l6 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchInFile(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m341 378l-81-82q17-27 17-59q0-44-31-75t-75-31t-75.5 31T64 237t31.5 75.5T171 344q31 0 59-18l94 95q-12 8-25 8H42q-17 0-29.5-12.5T0 387V45q0-17 12.5-29.5T43 3h170l128 128zM107 237.5q0-26.5 18.5-45.5t45-19t45.5 19t19 45.5t-19 45t-45.5 18.5t-45-18.5t-18.5-45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchInPage(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M203 128q22 0 37.5 15.5T256 181t-15.5 38t-37.5 16t-38-16t-16-38t16-37.5t38-15.5M384 21q18 0 30.5 12.5T427 64v256q0 18-12.5 30.5T384 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21zm-68 303l30-30l-62-62q15-23 15-51q0-40-28-68t-68-28t-68 28t-28 68t28 68t67 28q28 0 51-15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchReplace(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M149 64q-38 0-67.5 24.5T45 149H2q8-54 49.5-91T149 21q62 0 106 44l44-44v128H171l54-54q-32-31-76-31m121 195l103 104l-32 31l-103-103q-40 29-89 29q-62 0-105-44L0 320V192h128l-54 54q31 31 75 31q39 0 68-24t37-61h43q-5 37-27 67"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Seat(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 320v-64h341v128h-64v-64H107v64H43zm320-171h64v64h-64zM0 149h64v64H0zm320 64H107V43q0-18 12.5-30.5T149 0h128q18 0 30.5 12.5T320 43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SecTen(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 101l101-37h6v256H64V115L0 137zm507 142q5 8 5 21t-5 23q-6 11-15 18q-10 7-24 11q-13 4-30 4q-20 0-34.5-5T379 301q-9-9-14.5-20t-5.5-22h41q0 8 3 14q4 6 9 9q5 4 12 5q6 2 14 2q16 0 24.5-6t8.5-17q0-4-1-8q-2-4-6-7q-5-4-12-6q-8-3-20-6q-16-3-28-8t-20-11q-9-6-14-15t-5-21t5-21q5-11 14-18.5t23-12.5q13-4 29-4q18 0 32 5t23 12q10 8 15 19t5 23h-42q0-4-2-10t-6-9q-5-4-10-6q-7-3-14.5-3t-13.5 2t-10 5q-3 3-6 8q-1 4-1 8.5t1.5 8t5.5 6.5t12 5q8 3 19 5q15 4 28 8q12 5 22 12q9 6 13 16M295 86q11 13 16 34q6 21 6 51v41q0 30-6 51q-5 21-16 34q-11 14-26 19q-15 6-34 6q-18 0-34-6q-15-5-26-19q-11-12-17-34q-6-20-6-51v-41q0-29 6-51q6-21 16.5-34t26-18.5T235 62t34 5.5T295 86m-20 132v-54q0-18-2-32q-3-13-8-21t-13-11.5t-17-3.5q-10 0-18 3q-7 4-12.5 12t-8.5 21q-2 13-2 33v53q0 20 2 33q3 13 9 21q5 9 12.5 12.5T235 288t17.5-3.5t12.5-12t7.5-22T275 218"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SecThree(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M158 213q5 7 8 16q2 9 2 18q0 18-6.5 32T144 302q-12 10-27 15q-16 5-34 5q-16 0-31-4.5T25 304T7 281q-7-13-7-31h42q0 9 3 16t8.5 12t13.5 7q7 3 17 3q20 0 31-10q11-11 11-31q0-10-3-18t-10-13q-5-5-14-8q-9-2-20-2H54v-33h25q11 0 19-3t13-8q6-5 8.5-12t2.5-16q0-18-9-28q-10-10-29-10q-9 0-16 2.5T56 105q-5 5-8 11.5T45 131H2q0-15 7-27q5-12 16-22t26-15q14-5 32-5t32.5 4T141 79.5t17 22.5q6 14 6 32q0 8-2 15t-8 15q-4 7-12 14q-6 6-17 11q12 4 20 10t13 14m198 30q4 8 5 21q0 13-5 23q-6 11-16 18q-9 7-23 11t-31 4q-19 0-33.5-5T228 301t-15-20t-5-22h40q0 8 4 14t9 9q5 4 12 5q6 2 13 2q17 0 25.5-6t8.5-17q0-4-2-8q-1-4-5-7q-5-4-13-6q-8-3-20-6q-15-3-27-8q-13-5-21-11q-9-7-14-15q-4-10-4-21q0-12 4-21q6-11 15-18.5t22-12.5q13-4 30-4q18 0 31 5q14 5 23 12q10 8 15 19t5 23h-41q0-4-2-10q-3-6-7-9q-4-4-10-6q-6-3-14-3t-13.5 2t-9.5 5t-6 8q-2 4-2 8.5t1.5 8t6 6.5t11.5 5q9 3 19 5q15 4 29 8q12 5 21 12q9 6 14 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SelectAll(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 43q0-18 12.5-30.5T43 0v43zm0 170v-42h43v42zm85 171v-43h43v43zM0 128V85h43v43zM213 0v43h-42V0zm128 0q18 0 30.5 12.5T384 43h-43zM43 384q-18 0-30.5-12.5T0 341h43zM0 299v-43h43v43zM128 0v43H85V0zm43 384v-43h42v43zm170-171v-42h43v42zm0 171v-43h43q0 18-12.5 30.5T341 384m0-256V85h43v43zm0 171v-43h43v43zm-85 85v-43h43v43zm0-341V0h43v43zM85 299V85h214v214zm43-171v128h128V128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Settings(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m366 237l45 35q7 6 3 14l-43 74q-4 8-13 4l-53-21q-18 13-36 21l-8 56q-1 9-11 9h-85q-9 0-11-9l-8-56q-19-8-36-21l-53 21q-9 3-13-4L1 286q-4-8 3-14l45-35q-1-12-1-21t1-21L4 160q-7-6-3-14l43-74q5-8 13-4l53 21q18-13 36-21l8-56q2-9 11-9h85q10 0 11 9l8 56q19 8 36 21l53-21q9-3 13 4l43 74q4 8-3 14l-45 35q2 12 2 21t-2 21m-158.5 54q30.5 0 52.5-22t22-53t-22-53t-52.5-22t-52.5 22t-22 53t22 53t52.5 22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SettingsSquare(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 149q18 0 30.5 12.5T235 192t-12.5 30.5T192 235t-30.5-12.5T149 192t12.5-30.5T192 149M341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zm-37 192q0-7-1-15l32-24q4-5 1-10l-30-52q-3-5-9-3l-37 15q-12-9-25-15l-6-39q-1-6-7-6h-60q-6 0-7 6l-6 40q-14 5-25 14L87 88q-6-2-9 4l-30 51q-3 6 1 10l32 24q-1 8-1 15t1 15l-32 24q-4 5-1 10l30 52q3 5 9 3l37-15q12 9 25 15l6 39q1 6 7 6h60q6 0 7-6l6-40q14-5 25-14l37 15q6 2 9-4l30-51q3-6-1-10l-32-24q1-8 1-15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shape(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M79 0h164q16-2 29.5 9T287 37v111l36-59q30 47 84 135.5T474 335H233q-29 30-52 39q-42 18-88 4t-71-53q-27-39-21-89t41-82V48q-2-18 8.5-33T79 0m12 53v72q40-13 82 1t65 49V53zm40 114q-30 1-51 18.5T49 233q-8 34 13 65.5t56 36.5q34 6 64-17t32-58q5-36-21-65t-62-28m192 16l-62 102h125z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M320 303q26 0 44 18.5t18 44t-18 44t-44 18.5t-44-18.5t-18-44.5q0-6 1-14l-151-88q-19 17-44 17q-27 0-45.5-18.5T0 216t18.5-45.5T64 152q25 0 44 17l150-87q-2-9-2-15q0-27 18.5-45.5T320 3t45.5 18.5t18.5 45t-18.5 45.5t-45.5 19q-25 0-44-18l-150 88q2 9 2 15t-2 15l152 88q18-16 42-16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldCheck(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m192 5l192 86v128q0 89-55 162.5T192 475q-82-20-137-93.5T0 219V91zm-43 342l171-171l-30-30l-141 140l-55-55l-30 30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldSecurity(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m192 5l192 86v128q0 89-55 162.5T192 475q-82-20-137-93.5T0 219V91zm0 235V52L43 118v122zv191q59-19 100-72t49-119z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBasket(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M346 152h102q9 0 15 6.5t6 14.5v6l-54 198q-4 13-15.5 22t-26.5 9H96q-15 0-26-9t-15-22L1 179q-1-2-1-6q0-8 6.5-14.5T21 152h103l93-140q6-9 17.5-9t17.5 9zm-175 0h128l-64-94zm63.5 171q17.5 0 30-12.5T277 280t-12.5-30.5t-30-12.5t-30 12.5T192 280t12.5 30.5t30 12.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCart(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 344q18 0 30.5 12.5t12.5 30t-12.5 30T128 429t-30-12.5t-12-30t12-30t30-12.5M0 3h70l20 42h315q9 0 15.5 6.5T427 67q0 5-3 10l-76 138q-12 22-38 22H151l-19 35v3q0 5 5 5h247v43H128q-18 0-30.5-12.5T85 280q0-11 6-20l28-53L43 45H0zm341.5 341q17.5 0 30 12.5t12.5 30t-12.5 30t-30 12.5t-30-12.5t-12.5-30t12.5-30t30-12.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCartPlus(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213 176v-64h-64V69h64V5h43v64h64v43h-64v64zm-85 192q18 0 30.5 12.5t12.5 30t-12.5 30T128 453t-30-12.5t-12-30t12-30t30-12.5m213.5 0q17.5 0 30 12.5t12.5 30t-12.5 30t-30 12.5t-30-12.5t-12.5-30t12.5-30t30-12.5M132 299q0 5 5 5h247v43H128q-18 0-30.5-12.5T85 304q0-11 6-20l28-53L43 69H0V27h70l20 42l20 43l48 101l3 6h149l59-107l24-43l37 21l-82 149q-12 22-38 22H151l-19 35z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shuffle(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m141 132l-31 30L0 51l30-30zm83-111h117v118l-43-44L30 363L0 333L268 65zm7 201l67 67l43-44v118H224l44-44l-67-67z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignIn(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m151 269l55-56H0v-42h206l-55-56l30-30l107 107l-107 107zM341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341v-85h43v85h298V43H43v85H0V43q0-18 12.5-30.5T43 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkipNext(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 320V64l181 128zM213 64h43v256h-43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkipPrevious(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 64h43v256H0zm75 128L256 64v256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Skype(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M414 261q13 25 13 52q0 48-34.5 82T309 429q-30 0-56-13q-19 3-37 3q-85 0-144.5-59T12 218q0-20 4-40q-16-27-16-59q0-48 34.5-82T118 3q34 0 63 18q17-3 35-3q84 0 143.5 58.5T419 218q0 22-5 43m-95.5 52.5Q333 294 333 269q0-21-8.5-35.5T301 210q-14-10-34-16q-21-6-45-12q-20-4-29-7q-8-2-16-6t-12-9t-4-12q0-11 12-19q14-8 36-8q23 0 34 7q10 8 18 23q6 11 12 16t18 5t20.5-8.5T320 144t-6.5-22.5t-20-22t-33.5-17t-47-6.5q-35 0-60 10q-26 9-39.5 27T100 153q0 24 13 41q13 16 35 25q21 9 53 16q23 4 37 9q14 4 22 11q8 8 8 20q0 14-15 25q-16 10-41 10q-18 0-29.5-5T165 292t-11-21q-5-11-12-17q-8-6-18-6q-13 0-21.5 8T94 275q0 18 13 36t34 29q28 15 72 15q37 0 64-11t41.5-30.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Slideshare(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M153 132q21 0 35.5 14t14.5 33.5t-14.5 33T153 226t-36-13.5t-15-33t15-33.5t36-14m119 0q21 0 35.5 13.5t14.5 33t-14.5 33.5t-35.5 14t-36-14t-15-33.5t15-33t36-13.5m131 74q10-7 15-.5t-1 15.5q-29 36-88 60q26 89-22 131q-32 27-64 14q-27-10-26-42q0 1-.5-24.5T216 306l-4-1l-7-2v81q1 36-32 44q-36 9-65-23q-40-43-16-124q-60-25-89-60q-6-9-1-15.5t14 .5l4 3V44q0-17 12.5-29T61 3h300q16 0 26 12t10 29v165q2 0 6-3m-27 16V63q0-22-6.5-30.5T345 24H79q-20 0-26.5 8.5T46 63v160q23 14 51 19.5t46 4.5t34 0q15-1 22 6q1 0 1.5 1l.5 1q9 8 15 12q1-22 27-20q16-1 34 0t46-5t53-20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Slideshow(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m149 107l107 85l-107 85zM341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zm0 341V43H43v298z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Smartphone(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 6q18 0 30.5 12T299 48v384q0 18-12.5 30.5T256 475H43q-18 0-30.5-12.5T0 432V48q0-18 12.5-30.5T43 5zm0 383V91H43v298z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmartphoneAndroid(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M235 5q26 0 45 19t19 45v342q0 26-19 45t-45 19H64q-27 0-45.5-19T0 411V69q0-26 18.5-45T64 5zm-43 427v-21h-85v21zm69-64V69H37v299z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmartphoneDownload(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 6q18 0 30.5 12T299 48v384q0 18-12.5 30.5T256 475H43q-18 0-30.5-12.5T0 432V48q0-18 12.5-30.5T43 5zm0 383V91H43v298zm-21-128l-86 86l-85-86h64V155h43v106z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmartphoneErase(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m213 159l-85 85l85 86l-21 21l-85-85l-86 85l-21-21l85-86l-85-85l21-21l86 85l85-85zM341 5q18 0 30.5 12.5T384 48v384q0 18-12.5 30.5T341 475H128q-18 0-30.5-12.5T85 432v-64h43v43h213V69H128v43H85V48q0-18 12.5-30.5T128 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmartphoneInfo(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M171 133v43h-43v-43zm0 86v128h-43V219zM256 6q18 0 30.5 12T299 48v384q0 18-12.5 30.5T256 475H43q-18 0-30.5-12.5T0 432V48q0-18 12.5-30.5T43 5zm0 383V91H43v298z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmartphoneIphone(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M224 5q22 0 37.5 16T277 59v362q0 22-15.5 38T224 475H53q-22 0-37.5-16T0 421V59q0-22 15.5-38T53 5zm-85.5 448q13.5 0 23-9t9.5-22.5t-9.5-23t-23-9.5t-22.5 9.5t-9 23t9 22.5t22.5 9m96.5-85V69H43v299z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmartphoneLandscape(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 85q0-17 12.5-29.5T43 43h384q17 0 29.5 12.5T469 85v214q0 17-12.5 29.5T427 341H43q-18 0-30.5-12.5T0 299zm384 0H85v214h299z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmartphoneLandscapeLock(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 43q17 0 29.5 12.5T469 85v214q0 17-12.5 29.5T427 341H43q-18 0-30.5-12.5T0 299V85q0-17 12.5-29.5T43 43zm-43 256V85H85v214zm-192-22q-9 0-15-6t-6-15v-64q0-9 6-15t15-6v-22q0-17 12.5-29.5t30-12.5t30 12.5T277 149v22q9 0 15.5 6t6.5 15v64q0 9-6.5 15t-15.5 6zm17-128v22h51v-22q0-10-7.5-17.5t-18-7.5t-18 7.5T209 149"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmartphoneLock(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M320 5q18 0 30.5 12.5T363 48v384q0 18-12.5 30.5T320 475H107q-18 0-30.5-12.5T64 432v-64h43v43h213V69H107v43H64V48q0-18 12.5-30.5T107 5zM145 219q10 0 18 8t8 19v75q0 10-8.5 18t-19.5 8H26q-10 0-18-8.5T0 319v-75q0-9 8-17t18-8v-32q0-22 18-38t41-16t41.5 16t18.5 38zm-28 0v-32q0-13-9-20.5T85.5 159t-23 7.5T53 187v32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmartphonePortraitLock(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M107 325q-9 0-15.5-6T85 304v-64q0-9 6.5-15t15.5-6v-22q0-17 12.5-29.5t30-12.5t30 12.5T192 197v22q9 0 15 6t6 15v64q0 9-6 15t-15 6zm17-128v22h51v-22q0-10-7.5-17.5t-18-7.5t-18 7.5T124 197M256 5q18 0 30.5 12.5T299 48v384q0 18-12.5 30.5T256 475H43q-18 0-30.5-12.5T0 432V48q0-18 12.5-30.5T43 5zm0 384V91H43v298z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmartphoneRing(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M386 148q40 39 40 92t-40 90l-21-22q29-30 29-70t-29-68zm-45 45q20 21 20 47t-20 45l-21-22q18-24 0-49zM256 5q18 0 30.5 12.5T299 48v384q0 18-12.5 30.5T256 475H43q-18 0-30.5-12.5T0 432V48q0-18 12.5-30.5T43 5zm0 406V69H43v342z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmartphoneSetting(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 512v-43h43v43zm86 0v-43h42v43zm85 0v-43h43v43zM213 0q18 0 30.5 12.5T256 43v341q0 18-12.5 30.5T213 427H43q-18 0-30.5-12.5T0 384V43q0-18 12.5-30.5T43 0zm0 341V85H43v256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmartphoneSetup(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m189 251l23 19q4 4 2 6l-21 37q-2 2-7 2l-27-11q-13 9-19 11l-5 27q-4 5-6 5H86q-2 0-3.5-1.5T82 342l-4-27q-7-2-20-11l-29 9q-3 2-7-3L1 274q0-4 2-8l23-17v-22L3 210q-4-4-2-6l21-37q2-2 7-2l27 11q13-9 20-11l4-27q4-5 6-5h43q6 0 6 5l5 27q6 2 19 11l27-9q3-2 7 3l21 36q0 4-2 6l-23 17zm-81.5 32q17.5 0 30-12.5T150 240t-12.5-30.5t-30-12.5t-30 12.5T65 240t12.5 30.5t30 12.5M342 5q18 0 30.5 12.5T385 48v384q0 18-12.5 30.5T342 475H129q-18 0-30.5-12.5T86 432v-64h43v43h213V69H129v43H86V48q0-18 12.5-30.5T129 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAmountAsc(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 64h128v43H0zm0 256v-43h384v43zm0-149h256v42H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAmountDesc(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 320v-43h128v43zM0 64h384v43H0zm0 149v-42h256v42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAsc(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M282 76H135l74-73zM135 420h147l-73 73zM70 283h81l-41-111zm23-158h35l93 246h-38l-20-53H57l-19 53H0zm192 212h132v34H233v-28l128-183H234v-35h179v27z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortDesc(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M282 76H135l74-73zM135 420h147l-73 73zm131-137h81l-40-111zm24-158h34l93 246h-38l-19-53H254l-20 53h-38zM52 337h132v34H0v-28l128-183H1v-35h179v27z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Soundcloud(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 260H0v-51h9zm21 17h-9v-81h9zm17 4h-9v-94h9zm17 4h-9v-94h9zm21 0h-8V162h8zm17 0h-8V145h8zm22 0h-9V136h9zm17 0h-9V132h9zm21 0h-8V136h8zm17 0h-8V140h8zm17 0h-8V123h8zm22 0h-9V110h9zm156-1H228q-6 0-6-6V111q0-4 5-6q17-6 34-6q36 0 62.5 24.5T354 183q9-4 20-4q22 0 37.5 15.5T427 232t-15.5 37t-37.5 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpaceBar(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 128h42v128H0V128h43v85h256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Speaker(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 3q18 0 30.5 12.5T299 45v342q0 17-12.5 29.5T256 429H43q-18 0-30.5-12.5T0 387V45q0-17 12.5-29.5T43 3zM149.5 45q-17.5 0-30 12.5T107 88t12.5 30.5t30 12.5t30-12.5T192 88t-12.5-30.5t-30-12.5m-.5 342q44 0 75.5-31.5T256 280t-31.5-75.5T149 173t-75 31.5T43 280t31 75.5t75 31.5m.5-171q26.5 0 45 18.5T213 280t-18.5 45.5t-45 18.5t-45.5-18.5T85 280t19-45.5t45.5-18.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spellcheck(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m213 277l-24-64H68l-23 64H0L109 0h40l109 277zM85 171h88L129 53zm323 12l30 30l-202 203l-109-109l30-30l79 79z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spinner(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M212.5 7Q231 7 244 20.5T257 52t-13 31.5T212.5 97T181 83.5T168 52t13-31.5T212.5 7M337 54q22 0 37.5 16t15.5 37.5t-15.5 37.5t-37.5 16t-38-16t-16-37.5T299 70t38-16M91.5 64q16.5 0 28 12t11.5 28.5t-11.5 28.5t-28 12T63 133t-12-28.5T63 76t28.5-12M34 198q14 0 24 10t10 24t-10 24t-24 10t-24-10t-10-24t10-24t24-10m355 0q14 0 24 10t10 24t-10 24t-24 10t-24-10t-10-24t10-24t24-10M85 325q14 0 24 10t10 24t-10 24t-24 10t-24-10t-10-24t10-24t24-10m255 0q14 0 24 10t10 24t-10 24t-24 10t-24-10t-10-24t10-24t24-10m-128 47q14 0 24.5 10.5T247 407t-10.5 24t-24.5 10t-24-10t-10-24t10-24.5t24-10.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareDown(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m235 288l-86-85h64V11h43v192h64zM427 11q17 0 29.5 12.5T469 53v299q0 18-12.5 30.5T427 395H43q-18 0-30.5-12.5T0 352V53q0-17 12.5-29.5T43 11h128v42H43v299h384V53H299V11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareO(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 43H43v298h298zm0-43q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareRight(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 0q17 0 29.5 12.5T469 43v299q0 17-12.5 29.5T427 384H43q-18 0-30.5-12.5T0 342v-86h43v86h384V42H43v86H0V43q0-18 12.5-30.5T43 0zM213 277v-64H0v-42h213v-64l86 85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stackoverflow(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M245 395V250h34v179H0V251l32 1l-1 143zM52 335h167v35H52zm5-63l168 16l-4 36l-168-16zm15-73l163 46l-10 35l-163-46zm40-82l144 87l-19 32l-144-87zm150 81L164 61l30-21l98 137zM272 9l36-6l28 166l-36 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213 328L81 408l35-150L0 157l153-13L213 3l60 141l154 13l-117 101l35 150z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarCircle(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3M304 344l-24-103l79-69l-105-9l-41-96l-41 97l-105 8l80 69l-24 103l90-54z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarHalf(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 160L310 261l35 150l-132-80l-132 80l35-150L0 160l153-13L213 5l60 142zM213 291l81 49l-22-91l71-62l-93-8l-37-86z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarOutline(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 157L310 258l35 150l-132-80l-132 80l35-150L0 157l153-13L213 3l60 141zM213 289l81 48l-22-91l71-62l-93-8l-37-86l-36 86l-93 8l70 62l-21 91z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Steam(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M372 119q0 26-18 44.5T310 182t-44.5-18.5T247 119t18.5-44.5T310 56t44 18.5t18 44.5M0 319V209l65 26q20-12 45-12h9l73-105q0-48 34.5-82T309 2q49 0 83.5 34.5t34.5 83t-34.5 83T309 237l-112 82q-3 34-28 56.5T110 398q-32 0-56-19.5T24 329zM309.5 40Q277 40 254 63.5t-23 56t23 55.5t55.5 23t55.5-23t23-55.5t-23-56T309.5 40M110 246q-7 0-14 2l27 10q19 8 27.5 27.5t.5 39.5t-27.5 28t-39.5 1q-6-3-16.5-7.5T53 341q18 34 57 34q26 0 45-19t19-45.5t-19-45.5t-45-19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SteamSquare(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M355 153q0-23-16.5-39t-39-16t-39 16t-16.5 39t16.5 39t39 16t39-16t16.5-39M181 321q0 24-17 40t-40 16q-16 0-29.5-8T74 347q15 6 28 12q17 6 34-1t25-25q6-17-1-34t-25-24l-23-9q6-2 12-2q23 0 40 16.5t17 40.5M439 87v274q0 34-24 58t-58 24H82q-34 0-58-24T0 361v-44l49 20q6 26 27 43.5t48 17.5q30 0 52-20t25-50l98-72q43 0 73-30t30-73q0-42-30-72.5T299 50q-42 0-72 30t-31 72l-64 92h-8q-21 0-39 11L0 221V87q0-34 24-58T82 5h275q34 0 58 24t24 58m-71 66q0 29-20 49t-48.5 20t-49-20t-20.5-48.5t20.5-49T299 84q29 0 49 20.5t20 48.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stop(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 64h256v256H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Storage(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 363v-86h427v86zm43-64v42h42v-42zM0 21h427v86H0zm85 64V43H43v42zM0 235v-86h427v86zm43-64v42h42v-42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Store(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M363 21v43H21V21zm21 214h-21v128h-43V235h-85v128H21V235H0v-43L21 85h342l21 107zm-192 85v-85H64v85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StoreTwentyFour(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M363 85h64v278H256v-86h-85v86H0V85h64V21h299zm-171 64V85h-64v22h43v21h-43v64h64v-21h-43v-22zm107 43V85h-22v43h-21V85h-21v64h42v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Subway(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M170.5 3q39.5 0 67 3t53 11.5t38 26T341 88v203q0 31-21.5 52.5T267 365l32 32v11H43v-11l32-32q-31 0-53-21.5T0 291V88q0-27 12.5-44.5t38-26t53-11.5t67-3m-96 320q13.5 0 23-9.5t9.5-23t-9.5-22.5t-23-9t-22.5 9t-9 22.5t9 23t22.5 9.5M149 195V88H43v107zm117.5 128q13.5 0 23-9.5t9.5-23t-9.5-22.5t-23-9t-22.5 9t-9 22.5t9 23t22.5 9.5M299 195V88H192v107z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sun(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m123 95l-30 30l-39-38l30-30zM64 216v43H0v-43zM256 4v63h-43V4zm159 83l-38 38l-30-30l38-38zm-69 292l30-29l39 38l-30 30zm59-163h64v43h-64zM235 109q53 0 90.5 37.5T363 237t-37.5 90.5T235 365t-90.5-37.5T107 237t37.5-90.5T235 109m-22 362v-63h43v63zM54 388l39-39l30 30l-39 39z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SurroundSound(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 21q18 0 30.5 12.5T427 64v256q0 18-12.5 30.5T384 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21zM123 282q-38-37-38-90t38-91L93 71q-50 50-50 121t50 121zm90.5-5q35.5 0 60.5-25t25-60t-25-60t-60.5-25t-60.5 25t-25 60t25 60t60.5 25M334 313q50-50 50-121T334 71l-30 31q37 37 37 90t-37 91zM213.5 149q17.5 0 30 12.5T256 192t-12.5 30.5t-30 12.5t-30-12.5T171 192t12.5-30.5t30-12.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Swap(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M85 171v64h150v42H85v64L0 256zm299-43l-85 85v-64H149v-42h150V43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwapAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m341 21l86 86h-64v149q0 35-25 60t-60.5 25t-60.5-25t-25-60V107q0-18-12.5-30.5t-30-12.5t-30 12.5T107 107v149h64l-86 85l-85-85h64V107q0-36 25-61t60.5-25T210 46t25 61v149q0 18 12.5 30.5t30 12.5t30-12.5T320 256V107h-64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwapVertical(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M235 299h64l-86 85l-85-85h64V149h43zM85 0l86 85h-64v150H64V85H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwapVerticalCircle(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3M96 152h53v85h43v-85h53l-74-75zm235 128h-54v-85h-42v85h-54l75 75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tab(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 0q17 0 29.5 12.5T469 43v298q0 18-12.5 30.5T427 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zm0 341V128H256V43H43v298z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TabUnselected(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 128V85h43v43zm0 85v-42h43v42zM0 43q0-18 12.5-30.5T43 0v43zm171 341v-43h42v43zM0 299v-43h43v43zm43 85q-18 0-30.5-12.5T0 341h43zM427 0q17 0 29.5 12.5T469 43v85H256V0zm0 299v-43h42v43zM171 43V0h42v43zM85 384v-43h43v43zm0-341V0h43v43zm342 341v-43h42q0 18-12.5 30.5T427 384m0-171v-42h42v42zM256 384v-43h43v43zm85 0v-43h43v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tablet(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 21q17 0 29.5 12.5T469 64v256q0 18-12.5 30.5T427 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21zm-43 299V64H85v256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TabletAndroid(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M320 0q27 0 45.5 18.5T384 64v384q0 27-18.5 45.5T320 512H64q-27 0-45.5-18.5T0 448V64q0-27 18.5-45.5T64 0zm-85 469v-21h-86v21zm112-64V64H37v341z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TabletMac(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M352 0q22 0 37.5 15.5T405 53v406q0 22-15.5 37.5T352 512H53q-22 0-37.5-15.5T0 459V53q0-22 15.5-37.5T53 0zM202.5 491q13.5 0 23-9.5t9.5-23t-9.5-22.5t-23-9t-22.5 9t-9 22.5t9 23t22.5 9.5M363 405V64H43v341z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 0q18 0 30.5 12.5T512 43v298q0 18-12.5 30.5T469 384H149q-21 0-34-19L0 192L115 19q13-19 34-19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagClose(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 0q18 0 30.5 12.5T512 43v298q0 18-12.5 30.5T469 384H149q-21 0-34-19L0 192L115 19q13-19 34-19zm-64 269l-76-77l76-77l-30-30l-76 77l-77-77l-30 30l77 77l-77 77l30 30l77-77l76 77z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagMore(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 0q18 0 30.5 12.5T512 43v298q0 18-12.5 30.5T469 384H151q-23 0-36-19L0 192L115 19q13-19 34-19zM192 224q13 0 22.5-9.5T224 192t-9.5-22.5T192 160t-22.5 9.5T160 192t9.5 22.5T192 224m106.5 0q13.5 0 23-9.5T331 192t-9.5-22.5t-23-9.5t-22.5 9.5t-9 22.5t9 22.5t22.5 9.5m107 0q13.5 0 22.5-9.5t9-22.5t-9-22.5t-22.5-9.5t-23 9.5T373 192t9.5 22.5t23 9.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TapAndPlay(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 325q62 0 105.5 44T149 475h-42q0-44-31.5-75.5T0 368zm0 86q27 0 45.5 18.5T64 475H0zm0-171q97 0 166 68.5T235 475h-43q0-80-56-136T0 283zM320 6q18 0 30.5 12T363 48v363q0 17-12.5 29.5T320 453h-44q-4-45-21-85h65V91H107v128q-20-8-43-14V48q0-18 12.5-30.5T107 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextFormat(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 299h299v42H0zm96-90l-19 47H32L133 21h32l102 235h-45l-19-47zm53-145l-40 107h80z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Texture(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M352 2q24 6 31 30L32 382q-11-3-19-11T2 352zM189 0h61L0 250v-61zM43 0h42L0 85V43q0-18 12.5-30.5T43 0m298 384h-42l85-85v42q0 18-13 30q-12 13-30 13m-207 0l250-250v61L195 384z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThreeDrotation(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m159 458l29-28l81 81l-14 1q-100 0-173.5-68T0 277h32q6 60 40 108t87 73m19-139q14 0 21-7t7-20q0-7-2-12t-6-8q-4-4-9.5-5.5T175 265h-16v-22h16q8 0 13-2t8-5q4-3 6-8t2-10q0-12-7-19q-6-6-19-6q-5 0-10 2q-4 1-8 4q-3 3-5 8q-2 4-2 9h-28q0-10 4-18t11-14t17-10q9-3 21-3q11 0 22 3q10 3 16 9q7 6 11 15t4 20q0 5-2 10q-1 5-4 10q-4 5-8 9q-5 4-11 7q7 3 13 7q5 4 8 9q3 4 5 11q2 5 2 12q0 11-5 20q-4 9-11.5 15.5T200 338t-22 3q-11 0-21-3q-9-3-17-9t-12-14.5t-4-20.5h27q0 6 2 10.5t6 7.5q3 3 8 5t11 2m182.5-126.5Q371 203 377 218q5 16 5 34v8q0 19-5 34q-6 15-16 25q-10 11-25 17q-14 5-32 5h-49V171h50q18 0 31.5 5.5t24 16M352 260v-8q0-28-12-43q-12-14-35-14h-20v123h19q12 0 21-4t15-11q6-8 9-19t3-24M255 0q100 0 173.5 68T510 234h-32q-6-59-40.5-107T351 54l-29 28l-81-81z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbDown(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 0q17 0 29.5 12.5T341 43v213q0 18-12 30L188 427l-22-23q-10-9-10-22l1-7l20-98H43q-18 0-30.5-12.5T0 235v-43q0-8 3-16L67 26Q78 0 107 0zm85 0h85v256h-85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbUp(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 432V176h85v256zm469-235v43q0 8-3 16l-64 150q-11 26-39 26H171q-18 0-30.5-12.5T128 389V176q0-18 13-30L281 5l23 23q9 9 9 22l-1 7l-20 98h135q17 0 29.5 12.5T469 197"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbUpDown(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 128v27q0 6-2 11l-49 113q-8 20-29 20H32q-13 0-22.5-9.5T0 267V128q0-13 9-23L115 0l17 17q7 7 7 17l-1 5l-14 68h111q8 0 14.5 6t6.5 15m224 85q13 0 22.5 9.5T512 245v139q0 13-9 23L397 512l-17-17q-7-7-7-17l1-5l14-68H277q-8 0-14.5-6t-6.5-15v-27q0-6 2-11l49-113q8-20 29-20z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TicketStar(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 192q0 18 12.5 30.5T427 235v85q0 18-12.5 30.5T384 363H43q-18 0-30.5-12.5T0 320v-85q18 0 30.5-12.5T43 192t-12.5-30.5T0 149V64q0-18 12.5-30.5T43 21h341q18 0 30.5 12.5T427 64v85q-18 0-30.5 12.5T384 192m-94 102l-24-87l71-58l-91-5l-33-84l-33 84l-90 5l70 58l-23 87l76-49z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Time(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3m0 384q70.5 0 120.5-50t50-121t-50-121t-120.5-50T93 95T43 216t50 121t120.5 50M224 109v112l96 57l-16 27l-112-68V109z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeCountdown(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M171 298.5q0-8.5 6-15t15-6.5t15 6.5t6 15t-6 15t-15 6.5t-15-6.5t-6-15M171 0h21q80 0 136 56t56 136t-56 136t-136 56t-136-56T0 192q0-46 20.5-86.5T77 39v-1l145 145l-30 30L76 98q-33 41-33 94q0 62 43.5 105.5T192 341t105.5-43.5T341 192q0-56-36.5-98T213 44v41h-42zm149 192q0 9-6.5 15t-15 6t-15-6t-6.5-15t6.5-15t15-6t15 6t6.5 15m-256 0q0-9 6.5-15t15-6t15 6t6.5 15t-6.5 15t-15 6t-15-6t-6.5-15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeInterval(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M303.5 125.5Q341 163 341 216t-37.5 90.5T213 344t-90-38l90-90V88q53 0 90.5 37.5M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3m0 384q70.5 0 120.5-50t50-121t-50-121t-120.5-50T93 95T43 216t50 121t120.5 50"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeRestore(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0q80 0 136 56t56 136t-56 136t-136 56q-79 0-136-56l31-31q43 44 105 44t105.5-43.5T405 192T361.5 86.5T256 43T150.5 86.5T107 192h64l-87 86l-1-3l-83-83h64q0-80 56-136T256 0m-21 107h32v90l74 45l-15 26l-91-55z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeRestoreSetting(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 192q0 18-12.5 30.5T256 235t-30.5-12.5T213 192t12.5-30.5T256 149t30.5 12.5T299 192M256 0q80 0 136 56t56 136t-56 136t-136 56q-65 0-117-40l30-30q40 27 87 27q62 0 105.5-43.5T405 192T361.5 86.5T256 43T150.5 86.5T107 192h64l-86 85l-85-85h64q0-80 56-136T256 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Timer(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 5v43H128V5zm-85 278V155h42v128zm171-141q42 52 42 119q0 80-56 136t-136 56t-136-56T0 261.5t56-136T192 69q67 0 120 43l30-31q16 13 30 30zM192 411q62 0 105.5-44T341 261t-43.5-105.5T192 112T86.5 155.5T43 261t43.5 106T192 411"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimerOff(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m369 81l30 30l-30 31q42 52 42 119q0 58-32 106l-31-31q20-35 20-75q0-62-43.5-105.5T219 112q-40 0-75 20l-31-31q48-32 106-32q67 0 120 42zM283 5v43H155V5zm-86 180v-30h43v73zM27 69l214 214l164 165l-27 27l-53-54q-48 32-106 32q-80 0-136-56T27 261q0-58 32-106L0 96zm192 342q40 0 75-21L90 186q-21 35-21 75q0 62 44 106t106 44"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Toll(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M298.5 21Q369 21 419 71t50 121t-50 121t-120.5 50T178 313t-50-121t50-121t120.5-50m.5 299q53 0 90.5-37.5T427 192t-37.5-90.5T299 64t-90.5 37.5T171 192t37.5 90.5T299 320M43 192q0 41 23.5 74t61.5 47v44q-56-14-92-60T0 192T36 87t92-60v44q-38 14-61.5 47T43 192"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tonality(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3M192 385V47q-64 8-106.5 56T43 216t43 113t106 56m43-338v20h61q-29-16-61-20m0 62v22h126q-7-12-15-22zm0 64v22h148q-2-9-5-22zm0 212q32-4 61-20h-61zm111-62q8-10 15-22H235v22zm32-64q3-13 5-22H235v22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Toys(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M235 240q0-48 34.5-82.5T352 123t82.5 34.5T469 240zm0 0q0 48-35 82.5T117 357t-82.5-34.5T0 240zm0 0q-48 0-83-34.5T117 123t35-83t83-35zm0 0q48 0 82.5 34.5T352 357t-34.5 83t-82.5 35z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Traffic(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 149q0 30-18 52.5T277 232v24h64q0 29-18 52t-46 30v25q0 8-6 14.5t-15 6.5H85q-8 0-14.5-6.5T64 363v-25q-28-7-46-30T0 256h64v-24q-28-8-46-30.5T0 149h64v-24q-28-7-46-30T0 43h64V21q0-8 6.5-14.5T85 0h171q9 0 15 6.5t6 14.5v22h64q0 29-18 52t-46 30v24zM170.5 341q17.5 0 30-12.5t12.5-30t-12.5-30t-30-12.5t-30 12.5t-12.5 30t12.5 30t30 12.5m0-106q17.5 0 30-12.5T213 192t-12.5-30.5t-30-12.5t-30 12.5T128 192t12.5 30.5t30 12.5m0-107q17.5 0 30-12.5t12.5-30t-12.5-30t-30-12.5t-30 12.5t-12.5 30t12.5 30t30 12.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Transform(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 368h-86v43h43l-64 64l-64-64h43v-43H128q-18 0-30.5-12.5T85 325V155H0v-43h85V69H43l64-64l64 64h-43v256h299zM171 155v-43h128q17 0 29.5 12.5T341 155v128h-42V155z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Translate(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m253 281l-16 44l-66-66L64 365l-30-30l108-107q-40-44-63-97h42q20 39 50 71q45-50 67-114H0V45h149V3h43v42h149v43h-62q-24 78-79 139l-1 1zm120-108l96 256h-42l-24-64H301l-24 64h-42l96-256zm-56 150h70l-35-93z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrendingDown(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m299 320l49-49l-105-104l-85 85L0 94l30-30l128 128l85-85l135 134l49-49v128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrendingFlat(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m405 192l-85 85v-64H0v-42h320v-64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrendingUp(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 64h128v128l-49-49l-135 134l-85-85L30 320L0 290l158-158l85 85l105-104z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleDown(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213 306L350 87H77zm0 80L0 45h427z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleUp(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213 102L77 320h273zm0-81l214 342H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Truck(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m405 107l64 85v107h-42q0 26-19 45t-45.5 19t-45-19t-18.5-45H171q0 26-19 45t-45.5 19t-45-19T43 299H0V64q0-18 12.5-30.5T43 21h298v86zM106.5 331q13.5 0 23-9.5t9.5-23t-9.5-22.5t-23-9t-22.5 9t-9 22.5t9 23t22.5 9.5M395 139h-54v53h95zm-32.5 192q13.5 0 23-9.5t9.5-23t-9.5-22.5t-23-9t-22.5 9t-9 22.5t9 23t22.5 9.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tumblr(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 177v-60q25-8 43-23.5T72 57T87 3h61v108h102v66H148v110q0 37 4 47.5t15 16.5q14 9 33 9q32 0 65-21v67q-28 13-50.5 18t-48.5 5q-29 0-51.5-7T76 401t-22.5-29.5T47 327V178H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tune(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 299h128v42H0zM0 43h213v42H0zm213 341h-42V256h42v43h171v42H213zM85 128h43v128H85v-43H0v-42h85zm299 85H171v-42h213zm-128-85V0h43v43h85v42h-85v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TurningSign(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M420 201q7 6 7 15t-7 15L228 423q-6 6-15 6t-15-6L6 231q-6-6-6-15t6-15L198 9q6-6 15-6t15 6zm-164 68l75-74l-75-75v53H149q-9 0-15 6.5t-6 15.5v85h43v-64h85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tv(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 0q17 0 29.5 12.5T469 43v256q0 17-12.5 29.5T427 341H320v43H149v-43H43q-18 0-30.5-12.5T0 299V43q0-18 12.5-30.5T43 0zm0 299V43H43v256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TvAltPlay(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 88q17 0 29.5 12.5T469 131v256q0 17-12.5 29.5T427 429H43q-18 0-30.5-12.5T0 387V131q0-18 12.5-30.5T43 88h162l-71-70l15-15l86 85l85-85l15 15l-70 70zm0 299V131H43v256zM171 173l149 86l-149 85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TvList(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 0q17 0 29.5 12.5T469 43v256q0 17-12.5 29.5T427 341H320v43H149v-43H43q-18 0-30.5-12.5T0 299V43q0-18 12.5-30.5T43 0zm0 299V43H43v256zm-43-192v42H149v-42zm0 85v43H149v-43zm-256-85v42H85v-42zm0 85v43H85v-43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TvPlay(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 0q17 0 29.5 12.5T469 43v256q0 17-12.5 29.5T427 341H320v43H149v-43H43q-18 0-30.5-12.5T0 299V43q0-18 12.5-30.5T43 0zm0 299V43H43v256zM320 171l-149 85V85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitch(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M296 374h-83l-56 55h-55v-55H0V77L28 3h379v259zm74-130V40H65v269h83v55l56-55h102zM269 114h37v111h-37zM167 225V114h37v111z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitter(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M383 105v11q0 45-16.5 88.5t-47 79.5t-79 58.5T134 365q-73 0-134-39q10 1 21 1q61 0 109-37q-29-1-51.5-18T48 229q8 2 16 2q12 0 23-4q-30-6-50-30t-20-55v-1q19 10 40 11q-39-27-39-73q0-24 12-44q33 40 79.5 64T210 126q-2-10-2-20q0-36 25.5-61.5T295 19q38 0 64 27q30-6 56-21q-10 31-39 48q27-3 51-13q-18 26-44 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwitterBox(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M335 159q22-18 28-30q-13 6-31 9q18-13 24-32q-20 11-37 14q-12-14-31-16.5t-35.5 5t-26.5 25t-5 38.5q-67-4-118-59q-11 20-4.5 43.5T120 189q-11-1-24-7q1 43 44 57q-12 3-24 1q12 36 53 40q-15 13-39 19.5T85 303q45 28 92 26q70-3 113.5-49.5T335 159M384 3q18 0 30.5 12.5T427 45v342q0 17-12.5 29.5T384 429H43q-18 0-30.5-12.5T0 387V45q0-17 12.5-29.5T43 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Undo(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M224 107q73 0 131.5 43T437 261l-51 16q-17-51-61.5-84T224 160q-61 0-109 40l77 77H0V85l77 77q64-55 147-55"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnfoldLess(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 333l98-98l98 98l-30 30l-68-68l-68 68zM196 51l-98 98L0 51l30-30l68 68l68-68z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnfoldMore(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m98 60l-68 68L0 98L98 0l98 98l-30 30zm0 264l68-68l30 30l-98 98l-98-98l30-30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ungroup(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 3h85v21h150V3h85v85h-21v64h42v-21h86v85h-22v128h22v85h-86v-21H213v21h-85v-85h21v-43H85v22H0v-86h21V88H0zm341 213v-21h-42v42h21v86h-85v-22h-43v43h21v21h128v-21h22V216zM235 88V67H85v21H64v149h21v22h64v-43h-21v-85h85v21h43V88zm-22 128h-21v43h43v-22h21v-42h-43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Upload(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 21h299v43H0zm0 214L149 85l150 150h-86v128H85V235z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Usb(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M218 133h85v86h-21v42q0 18-12.5 30.5T239 304h-64v65q26 13 26 42q0 19-14 33t-33.5 14t-33-14t-13.5-33q0-29 25-42v-65H68q-17 0-29.5-12.5T26 261v-44Q0 204 0 176q0-19 14-33t33-14t33 14t14 33q0 28-26 41v44h64V91H90l64-86l64 86h-43v170h64v-42h-21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vibration(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 256V128h43v128zm64 43V85h43v214zm405-171h43v128h-43zm-64 171V85h43v214zM352 0q13 0 22.5 9.5T384 32v320q0 13-9.5 22.5T352 384H160q-13 0-22.5-9.5T128 352V32q0-13 9.5-22.5T160 0zm-11 341V43H171v298z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Videocam(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m299 160l85-85v234l-85-85v75q0 8-6.5 14.5T277 320H21q-8 0-14.5-6.5T0 299V85q0-8 6.5-14.5T21 64h256q9 0 15.5 6.5T299 85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideocamOff(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M405 99v228L167 88h132q8 0 14.5 6.5T320 109v75zM27 3l378 378l-27 27l-68-68q-6 4-11 4H43q-9 0-15.5-6.5T21 323V109q0-8 6.5-14.5T43 88h15L0 30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideocamSwitch(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m341 139l86-86v278l-86-86v75q0 9-6 15t-15 6H21q-8 0-14.5-6T0 320V64q0-9 6.5-15T21 43h299q9 0 15 6t6 15zM235 267l74-75l-74-75v54H107v-54l-75 75l75 75v-54h128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewAgenda(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 213q9 0 15 6.5t6 15.5v128q0 8-6 14.5t-15 6.5H21q-8 0-14.5-6.5T0 363V235q0-9 6.5-15.5T21 213zm0-213q9 0 15 6.5t6 14.5v128q0 9-6 15.5t-15 6.5H21q-8 0-14.5-6.5T0 149V21q0-8 6.5-14.5T21 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewArray(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 320V43h64v277zM299 43h64v277h-64zM85 320V43h192v277z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewCarousel(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M107 341V21h213v320zM0 299V64h85v235zM341 64h86v235h-86z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewColumn(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 320V43h107v277zM0 320V43h107v277zM256 43h107v277H256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewComfy(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 128V43h85v85zm0 107v-86h85v86zm107 0v-86h85v86zm106 0v-86h86v86zM107 128V43h85v85zm106-85h86v85h-86zm107 192v-86h85v86zM0 341v-85h85v85zm107 0v-85h85v85zm106 0v-85h86v85zm107 0v-85h85v85zm0-298h85v85h-85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewCompact(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 341V192h128v149zm149 0V192h256v149zM0 43h405v128H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewDashboard(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 213V0h171v213zm0 171V256h171v128zm213 0V171h171v213zm0-384h171v128H213z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewDay(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 384v-64h405v64zm384-277q9 0 15 6t6 15v128q0 9-6 15t-15 6H21q-8 0-14.5-6T0 256V128q0-9 6.5-15t14.5-6zM0 0h405v64H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewHeadline(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 256v-43h341v43zm0 85v-42h341v42zm0-170v-43h341v43zM0 43h341v42H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewList(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 235v-86h85v86zm0 106v-85h85v85zm0-213V43h85v85zm107 107v-86h256v86zm0 106v-85h256v85zm0-298h256v85H107z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewListAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 213v-42h43v42zm0 86v-43h43v43zm0-171V85h43v43zm85 85v-42h299v42zm0 86v-43h299v43zm0-214h299v43H85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewModule(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 171V43h107v128zm0 149V192h107v128zm128 0V192h107v128zm128 0V192h107v128zM128 171V43h107v128zM256 43h107v128H256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewQuilt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 320V192h107v128zM0 320V43h107v277zm256 0V192h107v128zM128 43h235v128H128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewStream(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 320V192h363v128zM0 43h363v128H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewSubtitles(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 21q18 0 30.5 12.5T427 64v256q0 18-12.5 30.5T384 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21zM43 192v43h85v-43zm213 128v-43H43v43zm128 0v-43h-85v43zm0-85v-43H171v43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewToc(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 128V85h299v43zm0 85v-42h299v42zm0 86v-43h299v43zm341 0v-43h43v43zm0-214h43v43h-43zm0 128v-42h43v42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewWeb(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 21q18 0 30.5 12.5T427 64v256q0 18-12.5 30.5T384 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21zM277 320v-85H43v85zm0-107v-85H43v85zm107 107V128h-85v192z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewWeek(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M85 43q9 0 15.5 6t6.5 15v256q0 9-6.5 15T85 341H21q-8 0-14.5-6T0 320V64q0-9 6.5-15T21 43zm299 0q9 0 15 6t6 15v256q0 9-6 15t-15 6h-64q-9 0-15-6t-6-15V64q0-9 6-15t15-6zm-149 0q8 0 14.5 6t6.5 15v256q0 9-6.5 15t-14.5 6h-64q-9 0-15.5-6t-6.5-15V64q0-9 6.5-15t15.5-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vignette(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 0q17 0 29.5 12.5T469 43v298q0 18-12.5 30.5T427 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0zM234.5 320q70.5 0 120.5-37.5t50-90.5t-50-90.5T234.5 64T114 101.5T64 192t50 90.5T234.5 320"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vimeo(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M276 100q-15 0-32 7Q276 4 364 6q66 2 62 86q-2 63-87 172q-87 114-147 114q-37 0-63-70q-18-66-34-127q-19-69-41-69q-5 0-34 20L0 106q33-29 62-56q42-36 63-38q50-5 62 68q12 80 17 99q14 65 32 65q13 0 40-42.5t29-64.5q3-37-29-37"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vk(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M548 85q7 18-43 84q-7 9-18 24q-23 28-26 37q-5 12 4 23q5 6 23 24h1v1q41 37 55 63l2 4l2 7.5v9.5l-7 7.5l-17 3.5l-73 2q-7 1-16.5-2t-14.5-6l-6-4q-9-6-20-18t-19.5-22t-17.5-16.5t-16-4.5q-1 0-2.5 1t-5 4.5t-6 8.5t-4.5 14.5t-2 22.5q0 4-1 7.5t-2 5.5l-1 1q-6 6-16 6h-32q-21 2-42.5-4T189 353.5t-29-19t-20-16.5l-7-7q-3-2-8-8t-20.5-26T74 234t-35-60.5T2 96q-2-5-2-8t1-5l1-1q4-6 16-6h79q3 0 6 1.5l5 2.5l1 1q5 3 7 9q6 14 13.5 29.5T141 143l4 8q9 17 16.5 29.5T175 200t12 11t10 4t8-1l1-1.5l3.5-6.5l4-13l2.5-23v-36q-1-11-3-20.5t-4-13.5l-1-3q-7-10-25-13q-3 0 2-7q5-5 11-8q15-8 68-7q23 0 39 4q5 1 9 3.5t6 7t3 9t1 13V114q-1 8-1 20v24q0 3-.5 12t-.5 14t1 11.5t3.5 11t6.5 6.5q2 1 4.5 1.5t7.5-3t11-10t15-19.5t19-30q17-30 31-65q1-2 2.5-4.5T425 79h1l1-1l4-1h6l82-1q11-1 18.5 1t8.5 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Voicemail(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M394.5 64q48.5 0 83 34.5t34.5 83t-34.5 83T395 299H117q-48 0-82.5-34.5T0 181.5t34.5-83t83-34.5t83 34.5T235 181q0 43-27 75h96q-27-32-27-75q0-48 34.5-82.5t83-34.5M117 256q31 0 53-22t22-53t-22-52.5t-53-21.5t-52.5 21.5T43 181t21.5 53t52.5 22m278 0q31 0 52.5-22t21.5-53t-21.5-52.5T395 107t-53 21.5t-22 52.5t22 53t53 22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeDown(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M288 192q0 28-14.5 51T235 278V106q24 12 38.5 35t14.5 51M0 128h85L192 21v342L85 256H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeMute(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 128h85L192 21v342L85 256H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeOff(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M288 192q0 6-1 13l-52-52v-47q24 12 38.5 35t14.5 51m53 0q0-50-30-89.5T235 49V5q64 15 106.5 67T384 192q0 47-22 89l-32-33q11-27 11-56M27 0l165 165l192 192l-27 27l-44-44q-35 29-78 39v-44q25-8 48-25l-91-91v144L85 256H0V128h101L0 27zm165 21v90l-45-45z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeUp(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 128h85L192 21v342L85 256H0zm288 64q0 28-14.5 51T235 278V106q24 12 38.5 35t14.5 51M235 5q64 15 106.5 67T384 192t-42.5 120T235 379v-44q46-14 76-53.5t30-89.5t-30-89.5T235 49z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Walk(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M160 85q-18 0-30.5-12.5t-12.5-30t12.5-30T160 0t30.5 12.5t12.5 30t-12.5 30T160 85m-79 73L21 459h45l39-171l44 43v128h43V299l-45-43l13-64q46 53 117 53v-42q-29 0-53.5-14.5T186 151l-22-34q-14-21-36-21q-3 0-8.5 1t-8.5 1L0 145v100h43v-72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wallpaper(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 45v150H0V45q0-17 12.5-29.5T43 3h149v42zm128 192l63 79l43-57l64 85H85zm149-95.5q0 13.5-9.5 22.5t-22.5 9t-22.5-9t-9.5-22.5t9.5-23T288 109t22.5 9.5t9.5 23M384 3q18 0 30.5 12.5T427 45v150h-43V45H235V3zm0 384V237h43v150q0 17-12.5 29.5T384 429H235v-42zM43 237v150h149v42H43q-18 0-30.5-12.5T0 387V237z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WashingMachine(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m110 319l121-121q25 25 25 60.5T231 319t-60.5 25t-60.5-25M299 3q17 0 29.5 12.5T341 45v342q0 17-12.5 29.5T299 429H43q-18 0-30.5-12.5T0 387V45q0-17 12.5-29.5T43 3zM128 45q-9 0-15 6.5t-6 15t6 15t15 6.5t15-6.5t6-15t-6-15t-15-6.5m-64 0q-9 0-15 6.5t-6 15t6 15T64 88t15-6.5t6-15t-6-15T64 45m107 342q53 0 90.5-37.5T299 259t-37.5-90.5T171 131t-90.5 37.5T43 259t37.5 90.5T171 387"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Watch(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 256q0 40-17 75t-48 59l-20 122H85L65 390Q0 339 0 256t65-134L85 0h171l20 122q31 24 48 59t17 75m-298 0q0 53 37.5 90.5T171 384t90.5-37.5T299 256t-37.5-90.5T171 128t-90.5 37.5T43 256"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WbAuto(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m146 206l25-78l24 78zM469 85h39l-44 192h-37l-32-130l-32 130h-38l-2-9q-21 43-62 69t-90 26q-71 0-121-50T0 192T50 71t121-50q81 0 133 64h16l26 135l32-135h34l32 135zM220 277h40L192 85h-43L81 277h41l15-42h68z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Whatsapp(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M364.5 65Q427 127 427 214.5T364.5 364T214 426q-54 0-101-26L0 429l30-109Q2 271 2 214q0-87 62-149T214 3t150.5 62M214 390q73 0 125-51.5T391 214T339 89.5T214 38T89.5 89.5T38 214q0 51 27 94l4 6l-18 65l67-17l6 3q42 25 90 25m97-132q9 5 10 7q4 6-3 25q-3 8-15 15.5t-21 9.5q-18 2-33-2q-17-6-30-11q-8-4-15.5-8.5t-14.5-9t-13-9.5t-11.5-10t-10.5-10.5t-8.5-9.5t-7-8.5t-5.5-7t-3.5-5L128 222q-22-29-22-55q0-24 19-44q6-7 14-7q6 0 10 1q8 0 12 9q2 3 6 13l7 17.5l3 8.5q3 5 1 9q-3 7-5 9l-3 3l-3 3.5l-2 2.5q-6 6-3 11q13 22 30 37q13 11 43 26q7 3 11-1q12-15 17-21q4-6 12-3q6 3 36 17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Widgets(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213 245h171v171H213zM0 416V245h171v171zM0 32h171v171H0zM291 4l121 121l-121 120l-120-120z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wifi(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m249 394l-1 1v-1L0 85Q113 0 248.5 0T497 85zM68 170q82-63 180.5-63T429 170L249 394l-1 1v-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiAlt(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 128q64-64 149.5-86.5t171 0T469 128l-42 43q-80-80-192.5-80T43 171zm171 171q26-27 63.5-27t64.5 27l-64 64zm-86-86q62-61 149.5-61T384 213l-43 43q-44-44-106.5-44T128 256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiAltTwo(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M249 42Q142 41 59 97l43 53q49-29 107-38q108-15 187 38l42-53q-83-56-189-55m-.5-39Q387 3 497 88L249 397L0 88Q110 3 248.5 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiInfo(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0q136 0 256 91L256 405L0 90Q119 0 256 0m21 277V149h-42v128zm-42-170h42V64h-42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiLock(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M482 277q8 0 14.5 7t6.5 15v85q0 8-6.5 14.5T482 405H375q-8 0-14.5-6.5T354 384v-85q0-8 6.5-15t14.5-7v-32q0-22 15.5-37.5t38-15.5t38 15.5T482 245zm-21 0v-32q0-12-9.5-22T429 213t-22.5 10t-9.5 22v32zm-139-32v56l-75 94L0 85Q114 0 247.5 0T495 85l-45 56q-6-2-21-2q-45 0-76 31t-31 75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiOff(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M497 125L380 270L160 50q44-10 88-10q136 0 249 85M356 301l74 74l-27 27l-71-71l-83 103l-1 1v-1L0 125q35-27 79-47L35 34L62 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiOutline(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M249 42Q144 42 59 97l190 237L438 97q-84-55-189-55m-.5-39q49.5 0 96 11T425 43.5t47 27T497 88L249 397L0 88q12-9 25-17.5t47.5-27T153 14t95.5-11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wikipedia(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M40 104Q23 75 2 67l-2-1V51h110v15q-13 1-21.5 7T83 90q14 33 40.5 94t38.5 89l46-87q-7-14-23-51.5T158 76q-7-10-36-11V51h102l1 14q-6 1-10 2t-7 4.5t-2 8.5l29 64q28-60 28-61q3-11-5-14.5T237 65l-1-14h92v14q-24 2-33 15q-14 20-46 89q23 53 43 95l78-180q-6-13-29-19l-1-14l87 1v14q-6 1-11 3q-11 5-18 17L291 333h-18l-52-120l-62 120h-18q-16-33-48-111T40 104"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowMaximize(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 21h341v342H0zm43 86v213h256V107z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowMinimize(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 363H0v-86h341z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowRestore(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 107h85V21h256v256h-85v86H0zm256 0v128h43V64H128v43zM43 192v128h170V192z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Windows(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 192V80l128-28v138zM363 0v187l-214 3V47zM0 213l128 2v146L0 336zm363 6v186l-214-40V215z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WrapText(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 341v-42h128v42zM341 43v42H0V43zm-64 128q36 0 61 25t25 60t-25 60t-61 25h-42v43l-64-64l64-64v43h48q17 0 29.5-12.5T325 256t-12.5-30.5T283 213H0v-42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wrench(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 389q6 5 6 14.5t-8 15.5l-49 49q-7 7-15.5 7t-14.5-7L189 274q-37 15-77.5 6.5T41 242q-31-32-39-75.5T14 84l94 92l64-64l-92-92q38-18 82-10.5T238 48q30 30 38.5 70.5T270 195z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xbox(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213 85q-24-18-47-27.5T127.5 47t-28 0T81 51l-6 3Q134 3 213 3t139 51q-3-1-7-3t-17.5-4t-28.5 0t-38.5 11T213 85m-56 41q-39 40-65 78t-34.5 63.5t-12 44.5t-1.5 28l3 9Q0 291 0 216q0-84 57-145q38 16 100 55m270 90q0 75-47 133q1-3 2.5-9t-1.5-27.5t-12-45.5t-34.5-62.5T269 126q28-17 53-31t36-19l11-5q58 61 58 145m-215-44q38 27 67.5 57t45 53t26 42t13.5 29l3 10q-62 66-153.5 66T59 363q2-4 5-11.5t15-30t28-44.5t44-51t61-54"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yahoo(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M414 113q-5 0-49 10q-10 3-62.5 45.5T246 224q-2 10-2 27l-1 15q0 9 4 39q4 1 32 1t32 1l-1 20q-6-1-105-1q-6 0-44 1t-49 1l4-19h15.5l27-2l15.5-6q1-1 1.5-2t1-2.5t.5-3V252q0-17-1-27q-3-10-51.5-69.5T59 83q-3-1-28.5-4T1 75L0 57q2-1 17.5-1t35.5.5t44-.5q23 0 61 .5t45 .5l-3 16q-4 1-30.5 2.5T138 79q16 24 50 68.5t39 51.5q2-3 41.5-36t40.5-43q-38-7-54-7l-3-20h89q72 0 86 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Youtube(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M170 184q18 0 57.5 1t59.5 2q15 0 26 3q28 6 34 40q5 35 5 59q0 39-3 87q-1 12-7 25q-11 24-43 25q-103 3-152 3q-17 0-47.5-1T59 427t-22-4q-20-5-29-26q-6-17-8-52q-1-41 2-94q1-15 5-31q9-31 42-33q31 0 121-3m32 191q4 3 13 9q9 5 17.5 1t10.5-15q2-9 2-14v-60q0-8-3-15q-3-13-12-16.5t-20 4.5q-2 1-4.5 3.5L202 276v-50h-21v158h21q-1-4 0-9m-40 9V266h-22v83q0 8-6 12q-4 5-9 3q-3-1-3-7v-91h-22v99q0 3 1 9q4 16 20 11q4-1 13-7q2-1 6-6v12zm161-55q0-4 .5-11t.5-12.5t-1-10.5q-1-14-9-22t-21-9q-14-1-23.5 6.5T259 292q-3 33 0 67q2 15 14 22.5t28 3.5q13-3 19.5-15.5T323 343h-22q0 10-1 14q-1 9-9 9t-9-8q-1-9-2-30q16 1 43 1m-220-80v-23H29v23h24v135h25V249zM247 44v93q0 6 4 7q3 1 7-2q7-5 7-14V44h22v119h-22v-12l-5 5q-8 6-11 7q-8 4-14.5.5T226 152q-1-4-1-7V45q7-1 22-1M66 3h22q2 0 4 4q10 34 14 51q0 1 2 3q4-18 9-32q1-3 3-10.5T123 7q1-3 4-4h22q0 1-1 4q-1 1-1 2q-4 14-13 42t-13 42q-2 6-2 10v59H96q0-3-.5-8.5T95 144t1-8q2-36-13-78q-9-27-17-55m142 101v11l-.5 13l-1.5 11q-1 11-10 18.5t-21 7.5t-20-7.5t-11-18.5q-1-7-1-21q0-38 1-49q4-27 31-27q26 0 31 27q0 1 .5 2.5t.5 2.5q0 5 .5 15t.5 15m-41-1v28q0 6 1 8q4 5 8 5t8-5q1-2 1-8V70q-1-7-9-7q-7 0-9 7q-1 2-1 6t.5 13t.5 14m39 222v-31q0-9 11-9q8 0 8 7v67q-1 7-8 7q-11-1-11-9q-1-16 0-32m93-14h-20q0-4 .5-10.5t.5-10.5q1-6 9-6q7 0 8 6q2 10 2 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YoutubePlay(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M422 107q5 35 5 69v32l-5 69q-4 29-17 42q-14 14-42 18q-27 2-64.5 3t-61.5 1h-24q-111-1-145-4l-8-1l-13-2l-12.5-5l-13-10l-10-16.5L6 284l-2-7q-4-35-4-69v-32l4-69q4-29 17-42q14-15 43-18q27-2 64-3t61-1h24q90 0 150 4q28 3 42 18q4 4 7 9.5t5 11t3 10.5t2 8zm-151 88l14-7l-115-60v120z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Zero(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M177 203q0 32-6.5 54t-18 36t-28 20.5T89 320q-20 0-37-6q-16-7-27-21q-12-14-19-36q-6-22-6-54v-44q0-32 6.5-54t18-36t28-20t36-6t36 6t28 20t18 36t6.5 54zm-45-51q0-19-3-34q-3-14-8-23q-6-8-14-12t-18-4q-11 0-19 4T57 95q-6 9-9 22.5T45 152v57q0 20 3 35q3 13 9 23q5 9 13 13t19 4t19-4t13-13t8-23t3-35z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomIn(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m267 235l106 106l-32 32l-106-106v-17l-6-6q-39 33-90 33q-58 0-98.5-40.5T0 138.5t40.5-98t98-40.5t98 40.5T277 139q0 51-33 90l6 6zm-128 0q40 0 68-28t28-68t-28-68t-68-28t-68 28t-28 68t28 68t68 28m53-86h-43v43h-21v-43H85v-21h43V85h21v43h43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomOut(children ...ElementRenderer) *ZmdiIcon {
	return &ZmdiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m267 235l106 106l-32 32l-106-106v-17l-6-6q-39 33-90 33q-58 0-98.5-40.5T0 138.5t40.5-98t98-40.5t98 40.5T277 139q0 51-33 90l6 6zm-128 0q40 0 68-28t28-68t-28-68t-68-28t-68 28t-28 68t28 68t68 28M85 128h107v21H85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
