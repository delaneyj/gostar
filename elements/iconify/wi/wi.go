package wi

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "2.0.12"
	hAttr          = "1em"
	viewbox        = "0 0 30 30"
	fill           = "currentColor"
)

type WiIcon struct {
	*SVGSVGElement
}

type WiIconFn func(children ...ElementRenderer) *WiIcon

var IconLookup = map[string]WiIconFn{
	"alien":                      Alien,
	"aliens":                     Aliens,
	"barometer":                  Barometer,
	"celsius":                    Celsius,
	"cloud":                      Cloud,
	"cloudDown":                  CloudDown,
	"cloudRefresh":               CloudRefresh,
	"cloudUp":                    CloudUp,
	"cloudy":                     Cloudy,
	"cloudyGusts":                CloudyGusts,
	"cloudyWindy":                CloudyWindy,
	"dayCloudy":                  DayCloudy,
	"dayCloudyGusts":             DayCloudyGusts,
	"dayCloudyHigh":              DayCloudyHigh,
	"dayCloudyWindy":             DayCloudyWindy,
	"dayFog":                     DayFog,
	"dayHail":                    DayHail,
	"dayHaze":                    DayHaze,
	"dayLightWind":               DayLightWind,
	"dayLightning":               DayLightning,
	"dayRain":                    DayRain,
	"dayRainMix":                 DayRainMix,
	"dayRainWind":                DayRainWind,
	"dayShowers":                 DayShowers,
	"daySleet":                   DaySleet,
	"daySleetStorm":              DaySleetStorm,
	"daySnow":                    DaySnow,
	"daySnowThunderstorm":        DaySnowThunderstorm,
	"daySnowWind":                DaySnowWind,
	"daySprinkle":                DaySprinkle,
	"dayStormShowers":            DayStormShowers,
	"daySunny":                   DaySunny,
	"daySunnyOvercast":           DaySunnyOvercast,
	"dayThunderstorm":            DayThunderstorm,
	"dayWindy":                   DayWindy,
	"degrees":                    Degrees,
	"directionDown":              DirectionDown,
	"directionDownLeft":          DirectionDownLeft,
	"directionDownRight":         DirectionDownRight,
	"directionLeft":              DirectionLeft,
	"directionRight":             DirectionRight,
	"directionUp":                DirectionUp,
	"directionUpLeft":            DirectionUpLeft,
	"directionUpRight":           DirectionUpRight,
	"dust":                       Dust,
	"earthquake":                 Earthquake,
	"fahrenheit":                 Fahrenheit,
	"fire":                       Fire,
	"flood":                      Flood,
	"fog":                        Fog,
	"galeWarning":                GaleWarning,
	"hail":                       Hail,
	"horizon":                    Horizon,
	"horizonAlt":                 HorizonAlt,
	"hot":                        Hot,
	"humidity":                   Humidity,
	"hurricane":                  Hurricane,
	"hurricaneWarning":           HurricaneWarning,
	"lightning":                  Lightning,
	"lunarEclipse":               LunarEclipse,
	"meteor":                     Meteor,
	"moonAltFirstQuarter":        MoonAltFirstQuarter,
	"moonAltFull":                MoonAltFull,
	"moonAltNew":                 MoonAltNew,
	"moonAltThirdQuarter":        MoonAltThirdQuarter,
	"moonAltWaningCrescentFive":  MoonAltWaningCrescentFive,
	"moonAltWaningCrescentFour":  MoonAltWaningCrescentFour,
	"moonAltWaningCrescentOne":   MoonAltWaningCrescentOne,
	"moonAltWaningCrescentSix":   MoonAltWaningCrescentSix,
	"moonAltWaningCrescentThree": MoonAltWaningCrescentThree,
	"moonAltWaningCrescentTwo":   MoonAltWaningCrescentTwo,
	"moonAltWaningGibbousFive":   MoonAltWaningGibbousFive,
	"moonAltWaningGibbousFour":   MoonAltWaningGibbousFour,
	"moonAltWaningGibbousOne":    MoonAltWaningGibbousOne,
	"moonAltWaningGibbousSix":    MoonAltWaningGibbousSix,
	"moonAltWaningGibbousThree":  MoonAltWaningGibbousThree,
	"moonAltWaningGibbousTwo":    MoonAltWaningGibbousTwo,
	"moonAltWaxingCrescentFive":  MoonAltWaxingCrescentFive,
	"moonAltWaxingCrescentFour":  MoonAltWaxingCrescentFour,
	"moonAltWaxingCrescentOne":   MoonAltWaxingCrescentOne,
	"moonAltWaxingCrescentSix":   MoonAltWaxingCrescentSix,
	"moonAltWaxingCrescentThree": MoonAltWaxingCrescentThree,
	"moonAltWaxingCrescentTwo":   MoonAltWaxingCrescentTwo,
	"moonAltWaxingGibbousFive":   MoonAltWaxingGibbousFive,
	"moonAltWaxingGibbousFour":   MoonAltWaxingGibbousFour,
	"moonAltWaxingGibbousOne":    MoonAltWaxingGibbousOne,
	"moonAltWaxingGibbousSix":    MoonAltWaxingGibbousSix,
	"moonAltWaxingGibbousThree":  MoonAltWaxingGibbousThree,
	"moonAltWaxingGibbousTwo":    MoonAltWaxingGibbousTwo,
	"moonFirstQuarter":           MoonFirstQuarter,
	"moonFull":                   MoonFull,
	"moonNew":                    MoonNew,
	"moonThirdQuarter":           MoonThirdQuarter,
	"moonWaningCrescentFive":     MoonWaningCrescentFive,
	"moonWaningCrescentFour":     MoonWaningCrescentFour,
	"moonWaningCrescentOne":      MoonWaningCrescentOne,
	"moonWaningCrescentSix":      MoonWaningCrescentSix,
	"moonWaningCrescentThree":    MoonWaningCrescentThree,
	"moonWaningCrescentTwo":      MoonWaningCrescentTwo,
	"moonWaningGibbousFive":      MoonWaningGibbousFive,
	"moonWaningGibbousFour":      MoonWaningGibbousFour,
	"moonWaningGibbousOne":       MoonWaningGibbousOne,
	"moonWaningGibbousSix":       MoonWaningGibbousSix,
	"moonWaningGibbousThree":     MoonWaningGibbousThree,
	"moonWaningGibbousTwo":       MoonWaningGibbousTwo,
	"moonWaxingCrescentFive":     MoonWaxingCrescentFive,
	"moonWaxingCrescentFour":     MoonWaxingCrescentFour,
	"moonWaxingCrescentOne":      MoonWaxingCrescentOne,
	"moonWaxingCrescentSix":      MoonWaxingCrescentSix,
	"moonWaxingCrescentThree":    MoonWaxingCrescentThree,
	"moonWaxingCrescentTwo":      MoonWaxingCrescentTwo,
	"moonWaxingGibbousFive":      MoonWaxingGibbousFive,
	"moonWaxingGibbousFour":      MoonWaxingGibbousFour,
	"moonWaxingGibbousOne":       MoonWaxingGibbousOne,
	"moonWaxingGibbousSix":       MoonWaxingGibbousSix,
	"moonWaxingGibbousThree":     MoonWaxingGibbousThree,
	"moonWaxingGibbousTwo":       MoonWaxingGibbousTwo,
	"moonWaxingSix":              MoonWaxingSix,
	"moonrise":                   Moonrise,
	"moonset":                    Moonset,
	"na":                         Na,
	"nightAltCloudy":             NightAltCloudy,
	"nightAltCloudyGusts":        NightAltCloudyGusts,
	"nightAltCloudyHigh":         NightAltCloudyHigh,
	"nightAltCloudyWindy":        NightAltCloudyWindy,
	"nightAltHail":               NightAltHail,
	"nightAltLightning":          NightAltLightning,
	"nightAltPartlyCloudy":       NightAltPartlyCloudy,
	"nightAltRain":               NightAltRain,
	"nightAltRainMix":            NightAltRainMix,
	"nightAltRainWind":           NightAltRainWind,
	"nightAltShowers":            NightAltShowers,
	"nightAltSleet":              NightAltSleet,
	"nightAltSleetStorm":         NightAltSleetStorm,
	"nightAltSnow":               NightAltSnow,
	"nightAltSnowThunderstorm":   NightAltSnowThunderstorm,
	"nightAltSnowWind":           NightAltSnowWind,
	"nightAltSprinkle":           NightAltSprinkle,
	"nightAltStormShowers":       NightAltStormShowers,
	"nightAltThunderstorm":       NightAltThunderstorm,
	"nightClear":                 NightClear,
	"nightCloudy":                NightCloudy,
	"nightCloudyGusts":           NightCloudyGusts,
	"nightCloudyHigh":            NightCloudyHigh,
	"nightCloudyWindy":           NightCloudyWindy,
	"nightFog":                   NightFog,
	"nightHail":                  NightHail,
	"nightLightning":             NightLightning,
	"nightPartlyCloudy":          NightPartlyCloudy,
	"nightRain":                  NightRain,
	"nightRainMix":               NightRainMix,
	"nightRainWind":              NightRainWind,
	"nightShowers":               NightShowers,
	"nightSleet":                 NightSleet,
	"nightSleetStorm":            NightSleetStorm,
	"nightSnow":                  NightSnow,
	"nightSnowThunderstorm":      NightSnowThunderstorm,
	"nightSnowWind":              NightSnowWind,
	"nightSprinkle":              NightSprinkle,
	"nightStormShowers":          NightStormShowers,
	"nightThunderstorm":          NightThunderstorm,
	"rain":                       Rain,
	"rainMix":                    RainMix,
	"rainWind":                   RainWind,
	"raindrop":                   Raindrop,
	"raindrops":                  Raindrops,
	"refresh":                    Refresh,
	"refreshAlt":                 RefreshAlt,
	"sandstorm":                  Sandstorm,
	"showers":                    Showers,
	"sleet":                      Sleet,
	"smallCraftAdvisory":         SmallCraftAdvisory,
	"smog":                       Smog,
	"smoke":                      Smoke,
	"snow":                       Snow,
	"snowWind":                   SnowWind,
	"snowflakeCold":              SnowflakeCold,
	"solarEclipse":               SolarEclipse,
	"sprinkle":                   Sprinkle,
	"stars":                      Stars,
	"stormShowers":               StormShowers,
	"stormWarning":               StormWarning,
	"strongWind":                 StrongWind,
	"sunrise":                    Sunrise,
	"sunset":                     Sunset,
	"thermometer":                Thermometer,
	"thermometerExterior":        ThermometerExterior,
	"thermometerInternal":        ThermometerInternal,
	"thunderstorm":               Thunderstorm,
	"timeEight":                  TimeEight,
	"timeEleven":                 TimeEleven,
	"timeFive":                   TimeFive,
	"timeFour":                   TimeFour,
	"timeNine":                   TimeNine,
	"timeOne":                    TimeOne,
	"timeSeven":                  TimeSeven,
	"timeSix":                    TimeSix,
	"timeTen":                    TimeTen,
	"timeThree":                  TimeThree,
	"timeTwelve":                 TimeTwelve,
	"timeTwo":                    TimeTwo,
	"tornado":                    Tornado,
	"train":                      Train,
	"tsunami":                    Tsunami,
	"umbrella":                   Umbrella,
	"volcano":                    Volcano,
	"windBeaufortEight":          WindBeaufortEight,
	"windBeaufortEleven":         WindBeaufortEleven,
	"windBeaufortFive":           WindBeaufortFive,
	"windBeaufortFour":           WindBeaufortFour,
	"windBeaufortNine":           WindBeaufortNine,
	"windBeaufortOne":            WindBeaufortOne,
	"windBeaufortSeven":          WindBeaufortSeven,
	"windBeaufortSix":            WindBeaufortSix,
	"windBeaufortTen":            WindBeaufortTen,
	"windBeaufortThree":          WindBeaufortThree,
	"windBeaufortTwelve":         WindBeaufortTwelve,
	"windBeaufortTwo":            WindBeaufortTwo,
	"windBeaufortZero":           WindBeaufortZero,
	"windDeg":                    WindDeg,
	"windDirection":              WindDirection,
	"windDirectionE":             WindDirectionE,
	"windDirectionN":             WindDirectionN,
	"windDirectionNe":            WindDirectionNe,
	"windDirectionNw":            WindDirectionNw,
	"windDirectionS":             WindDirectionS,
	"windDirectionSe":            WindDirectionSe,
	"windDirectionSw":            WindDirectionSw,
	"windDirectionW":             WindDirectionW,
	"windy":                      Windy,
}

func Alien(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.75 15.54c-1.12-2.4-.95-4.66.52-6.79c1.03-1.48 2.6-2.39 4.73-2.72c.16-.04.34-.07.54-.08h.63c2.91.09 5.05 1.38 6.4 3.88c.64 1.18.8 2.48.48 3.91c-.26 1.13-.68 2.19-1.28 3.17c-1.29 2.01-2.63 3.64-4 4.88c-.07.07-.17.16-.3.26c-.46.35-.89.53-1.28.54s-.83-.14-1.31-.45a3.37 3.37 0 0 1-.74-.59c-1.96-2-3.43-4-4.39-6.01m.11-2.21c.02.11.05.25.09.44s.07.32.09.4c.28 1.26.86 2.23 1.73 2.93c.88.7 1.96 1.11 3.26 1.23c.29.03.46.02.51-.03s.08-.23.09-.52c-.01-.08-.03-.21-.05-.39c-.02-.18-.04-.31-.06-.39c-.25-1.34-.88-2.32-1.9-2.93c-.18-.11-.39-.22-.62-.34s-.44-.2-.61-.27c-.17-.07-.4-.16-.69-.27c-.29-.11-.5-.19-.63-.25c-.16-.06-.42-.1-.8-.11c-.32 0-.46.17-.41.5m6.8 4.4c-.02.31 0 .49.06.56c.07.07.25.08.55.04c.38-.04.78-.12 1.2-.22c1.07-.27 1.94-.84 2.62-1.71c.34-.41.6-.86.77-1.34s.34-1.05.47-1.72c.05-.23.04-.38-.03-.46c-.07-.08-.22-.11-.44-.08c-.59.1-1.12.23-1.59.4c-1.15.43-2.02 1.01-2.62 1.74c-.6.74-.93 1.66-.99 2.79"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Aliens(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M27 372q-61-129 18-243q56-79 169-97l20-3h23q155 5 227 139q35 62 17 140q-12 61-45 112q-65 104-143 175q-4 5-12 9q-44 35-91 3q-18-12-26-20Q78 477 27 372zm5-79q1 10 5 30q15 68 62.5 105T216 471q16 2 18.5-.5T237 451l-3-27q-13-72-69-105q-11-7-90-40q-12-4-29-4t-14 18zm241 156q-1 17 3 20t20 2q14-2 42-8q56-13 94-60q29-35 44-110q2-11-1.5-15t-14.5-3q-32 3-58 13q-121 45-129 161zm28 586q-60-128 19-243q56-79 169-97l19-3h23q155 5 227 139q35 63 18 140q-12 59-46 112q-63 102-143 174q-8 8-11 10q-45 35-92 3q-15-10-26-21q-107-111-157-214zm5-79q1 6 3.5 16.5T312 985q15 68 62 105.5t116 43.5q16 1 19-1.5t3-19.5l-4-26q-13-72-68-105q-9-6-91-40q-15-5-28-5q-18 0-15 19zm241 156q-1 17 3 20t20 2q14-2 42-8q55-13 94-61q29-36 44-109q2-11-1.5-15.5T734 937q-28 3-58 14q-121 45-129 161zm74-740q-60-128 19-243q56-79 169-97l19-3h23q155 5 227 139q35 63 18 140q-12 59-46 112q-65 104-143 175q-6 6-11 9q-45 35-92 3q-18-12-26-20q-106-110-157-215zm5-79l6 30q15 68 62 105t116 43q16 2 19-1t3-19l-4-27q-13-72-68-105q-9-6-91-40q-12-4-28-4q-18 0-15 18zm241 156q-1 17 3 20t20 2q14-2 42-8q56-13 94-60q29-35 44-110q2-11-1.5-15t-14.5-3q-32 3-58 13q-121 45-129 161z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Barometer(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.69 13.2c0-.99.19-1.94.58-2.85c.39-.91.91-1.68 1.57-2.33s1.44-1.17 2.34-1.56c.9-.39 1.85-.58 2.84-.58c.99 0 1.94.19 2.85.58c.9.39 1.68.91 2.33 1.56c.65.65 1.17 1.43 1.56 2.33s.58 1.85.58 2.85c0 1.62-.48 3.06-1.44 4.34a7.247 7.247 0 0 1-3.71 2.61v3.29h-4.24v-3.25c-1.54-.45-2.81-1.32-3.79-2.61s-1.47-2.75-1.47-4.38m1.61 0c0 1.55.56 2.88 1.69 3.99c1.11 1.12 2.45 1.68 4.02 1.68c1.03 0 1.99-.25 2.86-.76a5.76 5.76 0 0 0 2.09-2.07c.51-.87.77-1.82.77-2.85c0-.77-.15-1.5-.45-2.21s-.71-1.31-1.22-1.82s-1.12-.92-1.83-1.22a5.6 5.6 0 0 0-2.21-.45c-.77 0-1.5.15-2.21.45a5.651 5.651 0 0 0-3.04 3.04c-.32.72-.47 1.45-.47 2.22m.58.36v-.72h2.17v.72zm1.09-3.54l.52-.52l1.52 1.52l-.52.53zm2.53 4.93c0-.42.15-.78.44-1.09c.29-.31.65-.47 1.06-.48l2.73-4.49l.66.35l-2.02 4.83c.18.25.26.54.26.88c0 .44-.15.81-.46 1.11c-.31.3-.68.45-1.12.45c-.43 0-.8-.15-1.1-.45c-.3-.3-.45-.67-.45-1.11m1.31-4.67V8.12h.69v2.17h-.69zm2.94 3.27v-.74h2.17v.74z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Celsius(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.75 10.98c0-.5.18-.93.53-1.28c.36-.36.78-.53 1.28-.53c.49 0 .92.18 1.27.53c.35.36.53.78.53 1.28s-.18.93-.53 1.28c-.35.36-.78.53-1.27.53c-.5 0-.93-.18-1.28-.53s-.53-.78-.53-1.28m.88 0c0 .26.09.48.27.67c.19.19.41.28.67.28c.26 0 .48-.09.67-.28s.28-.41.28-.67c0-.26-.09-.48-.28-.67s-.41-.28-.67-.28c-.26 0-.48.09-.67.28a.92.92 0 0 0-.27.67m3.89 4.42c0 .77.21 1.45.64 2.05c.22.31.53.56.93.75c.39.18.84.28 1.34.28c1.46 0 2.38-.56 2.75-1.67c.04-.14.02-.28-.06-.41a.485.485 0 0 0-.33-.23a.443.443 0 0 0-.4.07c-.12.08-.2.19-.23.34c0 .01 0 .02-.01.05l-.02.07c-.11.19-.26.34-.45.45c-.31.19-.72.28-1.23.28c-.31 0-.59-.05-.83-.16c-.4-.17-.68-.47-.85-.89c-.11-.27-.17-.6-.17-.97v-3.22c0-.15.01-.3.03-.45c.04-.38.19-.73.45-1.04c.29-.35.75-.52 1.38-.52c.52 0 .93.09 1.23.27c.2.12.35.27.45.45c.01.02.01.05.02.08s.01.05.01.06c.04.14.12.24.23.3c.12.07.25.08.4.05c.14-.03.25-.11.33-.23c.08-.12.1-.25.06-.4v-.01l-.08-.23c-.05-.11-.14-.26-.28-.43c-.13-.18-.29-.32-.45-.44c-.21-.15-.48-.27-.82-.38c-.34-.1-.71-.15-1.11-.15c-.51 0-.95.09-1.35.27c-.39.18-.7.42-.91.73c-.43.59-.65 1.28-.65 2.07v3.21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.61 16.88c0-1.15.36-2.17 1.08-3.07c.72-.9 1.63-1.48 2.74-1.73c.31-1.37 1.02-2.49 2.11-3.37s2.35-1.32 3.76-1.32c1.38 0 2.61.43 3.69 1.28s1.78 1.95 2.1 3.29h.33c.9 0 1.73.22 2.49.65s1.37 1.03 1.81 1.79c.44.76.67 1.58.67 2.48a4.94 4.94 0 0 1-2.36 4.25c-.73.45-1.54.69-2.41.72H9.41c-1.34-.06-2.47-.57-3.4-1.53c-.93-.95-1.4-2.1-1.4-3.44m1.71 0c0 .87.3 1.62.9 2.26s1.33.98 2.19 1.03H20.6c.86-.04 1.59-.39 2.19-1.03c.61-.64.91-1.4.91-2.26c0-.88-.33-1.63-.98-2.27c-.65-.64-1.42-.96-2.32-.96h-1.6c-.11 0-.17-.06-.17-.18l-.07-.57c-.11-1.08-.58-1.99-1.4-2.72c-.82-.73-1.77-1.1-2.86-1.1c-1.09 0-2.05.37-2.85 1.1c-.81.73-1.27 1.64-1.37 2.72l-.08.57c0 .12-.07.18-.2.18h-.53c-.84.1-1.54.46-2.1 1.07s-.85 1.33-.85 2.16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDown(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.61 16.88c0 1.34.47 2.48 1.4 3.44c.93.96 2.07 1.47 3.4 1.53c.11 0 .17-.06.17-.17v-1.34c0-.12-.06-.18-.17-.18c-.86-.04-1.59-.38-2.19-1.02c-.6-.64-.9-1.39-.9-2.26c0-.83.28-1.55.85-2.17c.57-.62 1.27-.97 2.1-1.07l.53-.04c.13 0 .2-.06.2-.17l.08-.55c.1-1.08.55-1.99 1.36-2.71c.81-.73 1.76-1.09 2.86-1.09c1.09 0 2.04.36 2.86 1.09a4.11 4.11 0 0 1 1.4 2.71l.07.58c0 .11.06.17.17.17h1.62c.89 0 1.66.32 2.31.97c.65.64.98 1.4.98 2.28c0 .87-.3 1.62-.91 2.26c-.61.64-1.34.98-2.19 1.02c-.13 0-.19.06-.19.18v1.34c0 .11.06.17.19.17c.88-.02 1.68-.26 2.41-.72a4.94 4.94 0 0 0 2.36-4.25c0-.9-.22-1.73-.67-2.48c-.44-.76-1.05-1.35-1.81-1.79s-1.59-.65-2.49-.65h-.33c-.33-1.34-1.03-2.43-2.1-3.29s-2.31-1.28-3.69-1.28c-1.41 0-2.67.44-3.76 1.31s-1.8 2-2.11 3.37c-1.1.26-2.01.84-2.73 1.74s-1.08 1.92-1.08 3.07m6.97 1.52c0 .24.08.44.24.6l2.59 2.61c.12.16.32.23.57.23c.28 0 .48-.08.61-.23l2.6-2.61c.16-.17.24-.38.24-.6c0-.23-.08-.43-.24-.58s-.36-.23-.6-.23s-.44.08-.62.23l-1.12 1.11v-3.98c0-.24-.08-.43-.25-.59c-.17-.16-.37-.23-.61-.23s-.43.08-.59.23c-.16.16-.23.35-.23.59v3.98l-1.1-1.11a.893.893 0 0 0-.63-.23c-.25 0-.45.08-.61.23c-.17.15-.25.35-.25.58"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudRefresh(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.63 16.91c0 .66.12 1.28.38 1.88c.25.6.59 1.11 1.02 1.55c.43.43.94.79 1.53 1.05c.59.27 1.21.42 1.87.45c.11 0 .17-.06.17-.17v-1.33c0-.12-.06-.19-.17-.19c-.87-.06-1.6-.41-2.19-1.03s-.89-1.36-.89-2.21c0-.84.28-1.57.85-2.19c.57-.62 1.26-.97 2.1-1.04l.52-.07c.13 0 .2-.06.2-.17l.07-.52A4.173 4.173 0 0 1 12.3 9.6c.62-.33 1.29-.49 2.01-.49c1.09 0 2.05.36 2.86 1.08a4.01 4.01 0 0 1 1.39 2.7l.06.57c0 .12.06.18.19.18h1.61c.9 0 1.67.32 2.32.97c.64.64.97 1.41.97 2.3c0 .85-.3 1.59-.89 2.21c-.59.62-1.32.97-2.19 1.03c-.13 0-.2.06-.2.19v1.33c0 .11.07.17.2.17c1.34-.06 2.47-.57 3.39-1.51s1.38-2.09 1.38-3.42c0-.89-.22-1.72-.67-2.48a4.96 4.96 0 0 0-1.81-1.8c-.76-.44-1.59-.67-2.48-.67h-.32c-.33-1.33-1.04-2.42-2.11-3.28c-1.11-.86-2.34-1.28-3.71-1.28c-1.41 0-2.66.44-3.75 1.33s-1.8 2.01-2.11 3.38a4.77 4.77 0 0 0-2.73 1.74c-.72.89-1.08 1.91-1.08 3.06m6.23 1.27c0 .74.19 1.43.56 2.07s.88 1.14 1.51 1.51c.63.38 1.32.56 2.06.56c1.15 0 2.13-.41 2.95-1.22c.82-.82 1.23-1.79 1.23-2.92c0-.23-.08-.43-.25-.6a.822.822 0 0 0-.61-.25c-.24 0-.44.08-.61.25s-.26.37-.26.6c0 .67-.24 1.24-.72 1.73c-.48.48-1.05.72-1.73.72c-.66 0-1.23-.24-1.71-.72s-.72-1.06-.72-1.73c0-.6.18-1.13.53-1.6c.36-.47.79-.73 1.31-.77l-.41.39c-.15.15-.23.34-.23.57c0 .25.07.47.23.66c.14.15.31.23.53.23c.22.01.45-.07.7-.23l1.82-1.87c.17-.17.25-.36.25-.58c0-.25-.08-.45-.25-.61l-1.82-1.83a.863.863 0 0 0-.62-.26c-.23 0-.43.08-.59.25c-.16.17-.24.37-.24.61s.07.43.23.58l.35.36c-1 .17-1.83.63-2.49 1.4c-.67.76-1 1.66-1 2.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUp(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.64 16.88c0 1.33.46 2.48 1.39 3.43c.93.96 2.06 1.47 3.4 1.53c.11 0 .17-.06.17-.17v-1.33c0-.12-.06-.19-.17-.19c-.86-.04-1.58-.38-2.18-1.02s-.9-1.39-.9-2.25c0-.82.28-1.54.84-2.16c.56-.61 1.26-.97 2.1-1.07h.53c.13 0 .2-.06.2-.18l.06-.57c.11-1.08.57-1.99 1.38-2.72s1.77-1.1 2.86-1.1c1.08 0 2.03.37 2.85 1.1c.82.73 1.28 1.64 1.39 2.72l.08.57c0 .12.06.18.18.18h1.61c.89 0 1.66.32 2.31.96s.98 1.4.98 2.26c0 .86-.3 1.61-.9 2.25c-.6.64-1.33.98-2.18 1.02c-.13 0-.2.06-.2.19v1.33c0 .11.07.17.2.17c.87-.02 1.67-.26 2.4-.71c.73-.45 1.31-1.05 1.73-1.8c.42-.75.63-1.57.63-2.44c0-.67-.13-1.31-.39-1.91a4.94 4.94 0 0 0-1.06-1.57a4.96 4.96 0 0 0-1.58-1.05c-.61-.26-1.25-.39-1.92-.39h-.32a5.885 5.885 0 0 0-2.11-3.29a5.763 5.763 0 0 0-3.68-1.28c-1.41 0-2.67.44-3.76 1.32s-1.79 2-2.1 3.36c-1.11.26-2.02.83-2.73 1.73c-.76.91-1.11 1.93-1.11 3.08m6.94.63c0 .25.08.46.24.64c.15.15.35.23.61.23c.24 0 .45-.08.62-.23l1.11-1.14v3.98c0 .24.08.44.23.61c.16.17.35.25.59.25c.23 0 .43-.08.6-.25c.17-.17.25-.37.25-.61v-3.94l1.12 1.11c.4.31.81.31 1.22 0c.16-.15.24-.36.24-.62c0-.24-.08-.44-.24-.62l-2.59-2.57a.814.814 0 0 0-.6-.24c-.24 0-.44.08-.59.24l-2.58 2.57c-.15.16-.23.35-.23.59"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloudy(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.89 17.6c0-.99.31-1.88.93-2.65s1.41-1.27 2.38-1.49c.26-1.17.85-2.14 1.78-2.88c.93-.75 2-1.12 3.22-1.12c1.18 0 2.24.36 3.16 1.09c.93.73 1.53 1.66 1.8 2.8h.27c1.18 0 2.18.41 3.01 1.24s1.25 1.83 1.25 3c0 1.18-.42 2.18-1.25 3.01s-1.83 1.25-3.01 1.25H8.16c-.58 0-1.13-.11-1.65-.34s-.99-.51-1.37-.89c-.38-.38-.68-.84-.91-1.36s-.34-1.09-.34-1.66m1.45 0c0 .76.28 1.42.82 1.96s1.21.82 1.99.82h9.28c.77 0 1.44-.27 1.99-.82c.55-.55.83-1.2.83-1.96s-.27-1.42-.83-1.96c-.55-.54-1.21-.82-1.99-.82h-1.39c-.1 0-.15-.05-.15-.15l-.07-.49c-.1-.94-.5-1.73-1.19-2.35s-1.51-.93-2.45-.93c-.94 0-1.76.31-2.46.94c-.7.62-1.09 1.41-1.18 2.34l-.07.42c0 .1-.05.15-.16.15l-.45.07c-.72.06-1.32.36-1.81.89c-.46.53-.71 1.16-.71 1.89m8.85-8.72c-.1.09-.08.16.07.21c.43.19.79.37 1.08.55c.11.03.19.02.22-.03c.61-.57 1.31-.86 2.12-.86s1.5.27 2.1.81c.59.54.92 1.21.99 2l.09.64h1.42c.65 0 1.21.23 1.68.7s.7 1.02.7 1.66c0 .6-.21 1.12-.62 1.57s-.92.7-1.53.77c-.1 0-.15.05-.15.16v1.13c0 .11.05.16.15.16c1.01-.06 1.86-.46 2.55-1.19s1.04-1.6 1.04-2.6c0-1.06-.37-1.96-1.12-2.7c-.75-.75-1.65-1.12-2.7-1.12h-.15c-.26-1-.81-1.82-1.65-2.47c-.83-.65-1.77-.97-2.8-.97c-1.4-.01-2.57.52-3.49 1.58"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudyGusts(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.62 21.01c0-.25.08-.46.25-.63c.17-.16.37-.24.6-.24h5.42c.74 0 1.37.26 1.89.79c.52.53.78 1.16.78 1.9s-.26 1.38-.78 1.9s-1.15.78-1.89.78s-1.38-.26-1.9-.79a.806.806 0 0 1-.23-.6c0-.24.08-.45.23-.6c.15-.16.35-.24.6-.24c.23 0 .43.08.61.24c.2.19.43.29.69.29s.49-.1.68-.29a.95.95 0 0 0 .29-.7c0-.26-.1-.49-.29-.68s-.42-.29-.68-.29H4.47c-.23 0-.43-.08-.6-.25s-.25-.35-.25-.59m0-3.04c0-.24.08-.45.25-.62c.17-.16.37-.24.6-.24h10.55c.26 0 .49-.1.68-.29c.19-.19.29-.43.29-.69s-.1-.5-.29-.69a.939.939 0 0 0-.68-.29c-.28 0-.5.09-.68.28c-.18.15-.39.23-.64.23c-.24 0-.44-.08-.6-.23a.814.814 0 0 1-.23-.6c0-.25.07-.45.23-.61c.51-.51 1.15-.76 1.92-.76c.74 0 1.38.26 1.9.78s.78 1.15.78 1.88s-.26 1.37-.78 1.89c-.52.52-1.15.78-1.9.78H4.47c-.24 0-.44-.08-.6-.24a.743.743 0 0 1-.25-.58m2.15-2.36c0 .08.05.12.16.12h1.44c.08 0 .15-.05.22-.15c.22-.54.58-.99 1.05-1.35c.48-.36 1.01-.56 1.59-.6l.53-.08c.13 0 .2-.06.2-.17l.07-.52c.11-1.08.56-1.99 1.37-2.72s1.76-1.1 2.86-1.1c1.11 0 2.07.36 2.88 1.09c.81.73 1.27 1.64 1.39 2.73l.07.59c0 .11.06.17.17.17h1.62c.91 0 1.68.32 2.33.96c.65.64.97 1.4.97 2.3c0 .89-.32 1.66-.97 2.3c-.65.64-1.42.96-2.33.96h-6.91c-.12 0-.19.06-.19.17v1.39c0 .11.06.17.19.17h6.91c.91 0 1.74-.22 2.51-.67c.77-.44 1.37-1.05 1.82-1.81c.45-.77.67-1.6.67-2.5c0-.91-.22-1.74-.67-2.5a4.938 4.938 0 0 0-1.82-1.81c-.77-.44-1.6-.67-2.51-.67h-.31c-.31-1.33-1.01-2.42-2.1-3.27c-1.08-.85-2.33-1.27-3.73-1.27c-1.41 0-2.66.44-3.75 1.32s-1.78 2-2.07 3.37c-.86.2-1.62.61-2.28 1.23s-1.12 1.36-1.38 2.21v.04z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudyWindy(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.1 21.04c0-.24.08-.45.25-.61s.38-.24.63-.24h8.97c.24 0 .43.08.59.24c.16.16.23.36.23.61c0 .24-.08.44-.24.6c-.16.16-.35.24-.59.24H3.98c-.25 0-.46-.08-.63-.24s-.25-.37-.25-.6m2.63-3.06c0-.24.09-.44.27-.6c.14-.15.34-.23.59-.23h9c.23 0 .42.08.58.23s.23.35.23.59s-.08.44-.23.61c-.15.17-.35.25-.58.25h-9c-.23 0-.43-.09-.6-.26s-.26-.36-.26-.59m.62-2.33c0 .09.06.14.17.14h1.43c.09 0 .17-.05.23-.14c.23-.54.57-.99 1.04-1.35s.99-.56 1.58-.6l.54-.07c.11 0 .17-.06.17-.18l.07-.52c.11-1.09.58-1.99 1.39-2.72c.82-.73 1.77-1.09 2.87-1.09c1.09 0 2.03.36 2.83 1.07c.8.72 1.27 1.62 1.41 2.7l.07.59c0 .11.06.16.18.16h1.6c.91 0 1.68.32 2.32.96c.64.64.96 1.41.96 2.32c0 .88-.33 1.64-.97 2.28c-.65.65-1.42.97-2.31.97h-6.89c-.12 0-.18.06-.18.17v1.34c0 .12.06.18.18.18h6.89c.68 0 1.32-.13 1.94-.39s1.14-.61 1.58-1.05s.79-.97 1.05-1.58s.39-1.25.39-1.92c0-.9-.22-1.73-.66-2.49a4.858 4.858 0 0 0-1.8-1.8c-.76-.44-1.6-.66-2.5-.66h-.31a5.948 5.948 0 0 0-2.1-3.3c-1.08-.85-2.3-1.28-3.68-1.28c-1.42 0-2.67.44-3.76 1.33c-1.09.88-1.78 2.01-2.08 3.39c-.86.19-1.62.6-2.27 1.21s-1.1 1.35-1.36 2.22v.02c-.01.04-.02.06-.02.09m1.15 8.48c0-.24.09-.44.26-.6c.15-.16.35-.23.59-.23h8.99c.24 0 .45.08.61.24c.17.16.25.36.25.6s-.08.44-.25.61c-.17.17-.37.25-.61.25H8.35c-.23 0-.43-.08-.6-.25a.86.86 0 0 1-.25-.62"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DayCloudy(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.56 16.9c0 .9.22 1.73.66 2.49s1.04 1.36 1.8 1.8c.76.44 1.58.66 2.47.66h10.83c.89 0 1.72-.22 2.48-.66c.76-.44 1.37-1.04 1.81-1.8c.44-.76.67-1.59.67-2.49c0-.66-.14-1.33-.42-2c.76-.92 1.14-2.03 1.14-3.3c0-.71-.14-1.39-.41-2.04c-.27-.65-.65-1.2-1.12-1.67c-.47-.47-1.02-.85-1.67-1.12c-.65-.28-1.33-.41-2.04-.41c-1.48 0-2.77.58-3.88 1.74c-.77-.44-1.67-.66-2.7-.66c-1.41 0-2.65.44-3.73 1.31a5.8 5.8 0 0 0-2.08 3.35c-1.12.26-2.03.83-2.74 1.73s-1.07 1.92-1.07 3.07m1.71 0c0-.84.28-1.56.84-2.17c.56-.61 1.26-.96 2.1-1.06l.5-.03c.12 0 .19-.06.19-.18l.07-.54c.14-1.08.61-1.99 1.41-2.71c.8-.73 1.74-1.09 2.81-1.09c1.1 0 2.06.37 2.87 1.1a3.99 3.99 0 0 1 1.37 2.71l.07.58c.02.11.09.17.21.17h1.61c.88 0 1.64.32 2.28.96c.64.64.96 1.39.96 2.27c0 .91-.32 1.68-.95 2.32c-.63.64-1.4.96-2.28.96H6.49c-.88 0-1.63-.32-2.27-.97c-.63-.65-.95-1.42-.95-2.32m6.7-12.27c0 .24.08.45.24.63l.66.64c.25.19.46.27.64.25c.21 0 .39-.09.55-.26s.24-.38.24-.62s-.09-.44-.26-.59l-.59-.66a.888.888 0 0 0-.61-.24c-.24 0-.45.08-.62.25c-.17.16-.25.36-.25.6m5.34 4.43c.69-.67 1.51-1 2.45-1c.99 0 1.83.34 2.52 1.03s1.04 1.52 1.04 2.51c0 .62-.17 1.24-.51 1.84c-.97-.96-2.13-1.44-3.49-1.44H17c-.25-1.09-.81-2.07-1.69-2.94m1.63-5.28c0 .26.08.46.23.62s.35.23.59.23c.26 0 .46-.08.62-.23c.16-.16.23-.36.23-.62V1.73c0-.24-.08-.43-.24-.59s-.36-.23-.61-.23c-.24 0-.43.08-.59.23s-.23.35-.23.59zm5.52 2.29c0 .26.07.46.22.62c.21.16.42.24.62.24c.18 0 .38-.08.59-.24l1.43-1.43c.16-.18.24-.39.24-.64c0-.24-.08-.44-.24-.6a.807.807 0 0 0-.59-.24c-.24 0-.43.08-.58.24l-1.47 1.43c-.15.19-.22.39-.22.62m.79 11.84c0 .24.08.45.25.63l.65.63c.15.16.34.24.58.24s.44-.08.6-.25a.86.86 0 0 0 .24-.62c0-.22-.08-.42-.24-.58l-.65-.65a.779.779 0 0 0-.57-.24c-.24 0-.44.08-.6.24c-.17.16-.26.36-.26.6m1.47-6.31c0 .23.09.42.26.58c.16.16.37.24.61.24h2.04c.23 0 .42-.08.58-.23s.23-.35.23-.59s-.08-.44-.23-.6s-.35-.25-.58-.25h-2.04c-.24 0-.44.08-.61.25a.79.79 0 0 0-.26.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DayCloudyGusts(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.35 21.01c0 .24.09.44.27.6c.17.17.37.25.61.25h5.88c.26 0 .48.09.68.28a.932.932 0 0 1 0 1.37c-.2.19-.42.29-.68.29c-.26 0-.48-.1-.68-.3a.888.888 0 0 0-.61-.24c-.24 0-.44.08-.6.24c-.16.16-.24.36-.24.6s.08.44.24.6c.53.53 1.16.8 1.89.8c.74 0 1.37-.26 1.88-.78c.52-.52.78-1.15.78-1.89s-.26-1.37-.78-1.89c-.52-.53-1.15-.79-1.88-.79H1.23c-.24 0-.44.09-.62.26s-.26.37-.26.6m0-3.01c0 .22.09.41.27.57c.17.17.37.25.61.25H12.2c.74 0 1.37-.26 1.89-.78c.52-.52.78-1.15.78-1.89s-.26-1.36-.78-1.88c-.52-.51-1.15-.77-1.89-.77c-.76 0-1.38.25-1.88.76c-.16.16-.23.37-.23.61s.08.44.23.6c.15.15.35.23.6.23c.24 0 .44-.08.62-.23c.19-.19.41-.28.67-.28c.26 0 .49.09.68.28c.19.19.29.42.29.68c0 .27-.1.5-.29.69c-.19.19-.42.29-.68.29H1.23a.87.87 0 0 0-.62.26c-.18.18-.26.38-.26.61M3 15.65c0 .08.06.12.17.12H4.6c.07 0 .13-.05.2-.14c.22-.54.57-.99 1.05-1.35c.47-.35 1-.55 1.6-.6l.53-.07c.12 0 .19-.06.19-.17l.07-.52c.11-1.08.56-1.98 1.37-2.71s1.76-1.09 2.85-1.09c1.1 0 2.05.36 2.86 1.08s1.27 1.63 1.38 2.71l.07.58c0 .12.06.18.18.18h1.63c.9 0 1.67.31 2.3.94c.63.63.95 1.39.95 2.27c0 .89-.32 1.66-.95 2.29c-.63.63-1.4.95-2.3.95h-6.9c-.13 0-.19.06-.19.18v1.37c0 .11.06.17.19.17h6.9c.89 0 1.72-.22 2.48-.67c.76-.44 1.36-1.05 1.8-1.81c.44-.76.66-1.59.66-2.48c0-.74-.14-1.41-.42-2.03c.76-.99 1.13-2.1 1.13-3.31c0-.94-.24-1.81-.71-2.62S22.41 7.47 21.6 7c-.81-.47-1.68-.71-2.62-.71c-1.54 0-2.83.58-3.86 1.73c-.8-.41-1.69-.61-2.67-.61c-1.41 0-2.65.44-3.73 1.31s-1.77 1.99-2.06 3.34c-.85.2-1.6.61-2.25 1.23c-.65.62-1.11 1.35-1.36 2.19v.04c-.04.07-.05.1-.05.13m8.18-11.03c0 .23.09.43.27.6l.61.65c.16.16.37.24.61.24c.25 0 .45-.08.61-.23c.16-.15.24-.35.24-.6c0-.24-.07-.44-.23-.6l-.66-.65c-.16-.17-.35-.26-.59-.26s-.44.08-.61.25c-.17.17-.25.37-.25.6m5.32 4.41c.72-.68 1.54-1.02 2.48-1.02c.97 0 1.8.35 2.51 1.05c.7.7 1.05 1.54 1.05 2.51c0 .65-.17 1.26-.52 1.83c-.96-.96-2.11-1.43-3.46-1.43h-.34c-.23-1.09-.81-2.07-1.72-2.94m1.62-5.24c0 .23.08.43.25.59c.17.16.37.24.6.24c.24 0 .44-.08.61-.24c.17-.16.25-.35.25-.59V1.74c0-.24-.08-.44-.25-.61a.822.822 0 0 0-.61-.25c-.24 0-.44.08-.6.25s-.25.37-.25.61zm5.55 2.27c0 .24.08.44.23.6c.17.17.37.25.61.26s.43-.08.57-.26l1.46-1.43c.17-.17.25-.37.25-.61c0-.23-.08-.43-.25-.6c-.17-.17-.37-.25-.61-.25s-.44.07-.6.23L23.9 5.47c-.15.16-.23.35-.23.59m.77 11.83c0 .23.08.43.25.6l.64.65c.2.16.41.24.62.24c.19 0 .39-.08.59-.24c.17-.17.25-.37.25-.6a.89.89 0 0 0-.25-.61l-.64-.65a.801.801 0 0 0-.58-.24c-.25 0-.46.08-.63.25c-.17.16-.25.36-.25.6m1.51-6.32c0 .24.08.43.25.59c.15.18.34.26.57.26h2.02c.24 0 .44-.08.61-.25c.17-.17.25-.37.25-.6c0-.23-.09-.43-.26-.6a.826.826 0 0 0-.6-.26h-2.02c-.24 0-.43.08-.59.25a.89.89 0 0 0-.23.61"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DayCloudyHigh(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.95 13.05c0-.93.29-1.75.87-2.48s1.31-1.2 2.19-1.4c.26-1.1.82-2 1.7-2.71s1.88-1.06 3.01-1.06c1.1 0 2.08.35 2.95 1.04s1.43 1.57 1.68 2.65h.26c1.1 0 2.04.39 2.82 1.16c.78.77 1.17 1.71 1.17 2.81v.1c.75.8 1.12 1.75 1.12 2.85c0 .76-.19 1.46-.57 2.1c-.38.65-.89 1.16-1.53 1.53c-.64.38-1.34.56-2.09.56c-.96 0-1.82-.3-2.56-.89s-1.24-1.35-1.48-2.26h-5.7c-1.07-.05-1.98-.46-2.72-1.23s-1.12-1.7-1.12-2.77m1.36 0c0 .7.24 1.31.73 1.82s1.07.79 1.75.82h8.99c.68-.03 1.27-.3 1.75-.82c.49-.52.73-1.12.73-1.82c0-.71-.26-1.32-.79-1.83c-.53-.52-1.14-.77-1.86-.77h-1.29c-.09 0-.14-.05-.14-.14l-.07-.47c-.09-.87-.46-1.6-1.12-2.19s-1.42-.89-2.28-.89c-.89 0-1.66.29-2.31.88s-1 1.32-1.09 2.19l-.06.47c0 .09-.05.14-.16.14h-.4c-.67.08-1.24.36-1.69.86a2.48 2.48 0 0 0-.69 1.75m6.2 9.01c-.25-.33-.25-.65 0-.98l1.13-1.15c.14-.12.31-.18.52-.18c.19 0 .34.06.46.18s.18.28.18.47c0 .2-.06.36-.18.48l-1.14 1.18a.68.68 0 0 1-.49.19c-.2 0-.36-.07-.48-.19m3.39-5.02c.21.54.56.97 1.04 1.3c.48.33 1.01.5 1.6.5c.77 0 1.43-.28 1.97-.83a2.8 2.8 0 0 0 .81-2.02c0-.39-.06-.74-.19-1.05c-.33.61-.8 1.11-1.39 1.49c-.6.38-1.25.58-1.96.61zm1.95 5.19c0-.19.07-.34.2-.47c.13-.12.3-.19.48-.19s.35.07.5.21c.12.12.19.27.19.45v1.64c0 .19-.07.35-.2.49s-.3.21-.48.21s-.35-.07-.48-.21c-.13-.14-.2-.3-.2-.49v-1.64zm4.41-1.83c0-.18.06-.33.19-.46c.13-.12.29-.19.47-.19c.19 0 .35.06.47.18l1.18 1.15c.13.14.2.3.2.48c0 .19-.07.35-.2.48s-.3.2-.49.2c-.21 0-.37-.06-.5-.19l-1.13-1.18a.674.674 0 0 1-.19-.47m0-8.81c0-.19.06-.35.19-.47l1.13-1.18c.14-.12.3-.19.5-.19c.19 0 .35.06.5.19c.13.15.2.32.2.51c0 .18-.07.33-.2.48l-1.18 1.15c-.12.12-.28.19-.47.19s-.35-.06-.47-.19c-.14-.14-.2-.3-.2-.49m1.82 4.4c0-.19.06-.35.19-.48c.12-.13.28-.2.47-.2h1.62c.19 0 .36.07.5.2s.21.29.21.48s-.07.36-.21.49c-.14.13-.3.2-.5.2h-1.62c-.19 0-.34-.07-.47-.2a.655.655 0 0 1-.19-.49"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DayCloudyWindy(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.45 20.97c0 .24.08.45.24.61c.44.18.73.27.88.27h7.88c.24 0 .44-.09.6-.26c.17-.17.25-.38.25-.61c0-.23-.08-.43-.25-.59a.847.847 0 0 0-.6-.24H1.57c-.26 0-.52.08-.76.24c-.24.16-.36.36-.36.58m1.39-3c0 .24.08.43.25.59c.15.17.34.26.58.26h9.4c.24 0 .44-.08.61-.25c.17-.17.25-.37.25-.6c0-.24-.08-.44-.25-.61a.822.822 0 0 0-.61-.25h-9.4c-.23 0-.43.08-.59.25c-.16.17-.24.37-.24.61m1.05-2.37c0 .09.06.13.17.13h1.39c.12 0 .19-.04.22-.13c.26-.53.62-.97 1.09-1.32c.47-.35 1-.55 1.58-.62h.54c.11 0 .16-.06.16-.19l.07-.56c.07-.71.3-1.36.69-1.95c.39-.58.9-1.04 1.53-1.37s1.3-.5 2.02-.5c1.09 0 2.04.37 2.85 1.1s1.27 1.64 1.39 2.72l.07.56c0 .12.06.19.18.19h1.6c.89 0 1.65.32 2.3.96c.65.64.97 1.39.97 2.27c0 .9-.32 1.67-.97 2.31c-.64.64-1.41.96-2.31.96h-6.89c-.11 0-.17.06-.17.19v1.33c0 .12.06.18.17.18h6.89c.9 0 1.73-.22 2.49-.67c.76-.44 1.37-1.05 1.81-1.81c.44-.76.67-1.59.67-2.49c0-.73-.14-1.39-.43-2.01c.78-.96 1.16-2.06 1.16-3.28c0-.94-.24-1.81-.71-2.62a5.201 5.201 0 0 0-1.92-1.92c-.81-.47-1.69-.71-2.63-.71c-.73 0-1.43.15-2.1.45c-.67.3-1.25.71-1.74 1.25c-.83-.43-1.73-.65-2.7-.65c-1.41 0-2.67.44-3.76 1.31s-1.79 1.99-2.09 3.36c-.85.21-1.6.63-2.25 1.25s-1.1 1.36-1.35 2.21c.02.02.01.04.01.07m.71 8.56c0 .24.09.43.26.59c.16.17.36.25.59.25h9.42c.23 0 .43-.08.59-.25s.24-.36.24-.6c0-.25-.08-.46-.24-.62s-.36-.25-.6-.25H4.45c-.24 0-.44.08-.6.25s-.25.38-.25.63m7.49-19.51c0 .25.08.45.24.6l.64.66c.16.16.36.24.6.24c.26 0 .46-.08.62-.24c.16-.16.24-.36.24-.61c0-.23-.08-.43-.24-.59l-.65-.65a.783.783 0 0 0-.57-.25c-.25 0-.46.08-.63.25s-.25.36-.25.59m5.36 4.38c.66-.63 1.48-.95 2.45-.95c.97 0 1.8.34 2.49 1.03c.68.68 1.03 1.51 1.03 2.49c0 .67-.15 1.27-.46 1.81c-.94-.95-2.11-1.43-3.5-1.43h-.3a6.084 6.084 0 0 0-1.71-2.95m1.6-5.22c0 .24.08.43.25.59s.36.23.6.23c.25 0 .45-.08.6-.23c.15-.15.23-.35.23-.6V1.76c0-.24-.08-.45-.23-.61a.785.785 0 0 0-.6-.25c-.23 0-.43.08-.6.25s-.25.37-.25.61zm5.52 2.28c0 .24.08.44.25.6c.12.16.33.24.6.24c.27 0 .47-.08.59-.24l1.46-1.44c.16-.15.24-.36.24-.62c0-.23-.08-.43-.25-.6c-.17-.17-.37-.25-.6-.25s-.44.08-.61.23L23.83 5.5c-.17.17-.26.36-.26.59m.8 11.86c0 .24.08.44.23.6l.66.63c.24.18.45.27.61.27c.16 0 .37-.09.61-.27c.16-.16.24-.36.24-.6c0-.23-.08-.43-.24-.61l-.64-.61a.94.94 0 0 0-.65-.26c-.24 0-.43.08-.59.24a.89.89 0 0 0-.23.61m1.44-6.32c0 .24.09.45.27.61c.18.17.38.25.6.25h2.03c.23 0 .42-.08.59-.25c.17-.17.25-.37.25-.61c0-.22-.08-.41-.24-.57a.84.84 0 0 0-.59-.23h-2.03c-.24 0-.45.08-.62.23c-.17.16-.26.35-.26.57"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DayFog(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.32 21.06c0 .23.08.43.25.59c.17.16.38.24.63.24h18.71c.24 0 .44-.08.61-.24c.17-.16.25-.35.25-.59s-.08-.44-.25-.6a.822.822 0 0 0-.61-.25H1.2c-.25 0-.46.08-.63.25s-.25.36-.25.6m2.62-3.14c0 .23.08.43.25.58a.8.8 0 0 0 .6.27h18.72c.23 0 .43-.08.59-.25c.16-.17.24-.37.24-.6a.84.84 0 0 0-.23-.59a.791.791 0 0 0-.59-.24H3.8c-.24 0-.44.08-.6.24c-.17.17-.26.36-.26.59m.13-2.4c0 .09.05.13.16.13h1.43c.07 0 .14-.05.21-.16c.24-.52.59-.94 1.06-1.27c.47-.33.99-.52 1.56-.56l.54-.07c.11 0 .17-.06.17-.18l.07-.51c.11-1.08.56-1.98 1.37-2.71c.81-.72 1.76-1.09 2.86-1.09c1.08 0 2.03.36 2.84 1.08c.81.72 1.27 1.61 1.38 2.68l.07.58c0 .11.06.17.19.17h1.61c.64 0 1.23.17 1.76.52c.53.34.92.8 1.18 1.37c.07.11.13.16.2.16h1.44c.13 0 .18-.07.13-.23l-.2-.55c.76-.94 1.13-2.04 1.13-3.31c0-.71-.14-1.38-.41-2.03s-.64-1.2-1.11-1.67c-.46-.47-1.02-.84-1.67-1.12s-1.32-.4-2.04-.4c-1.54 0-2.82.56-3.82 1.68c-.85-.42-1.74-.63-2.68-.63c-1.4 0-2.65.44-3.74 1.32s-1.79 2-2.1 3.37c-1.78.47-2.98 1.58-3.58 3.35c-.01.01-.01.04-.01.08m1.62 8.61c0 .24.09.44.27.6c.16.17.35.26.59.26h18.74c.23 0 .43-.08.6-.25c.17-.17.25-.37.25-.61c0-.23-.08-.42-.25-.58a.882.882 0 0 0-.6-.23H5.55a.89.89 0 0 0-.61.23c-.17.16-.25.35-.25.58m6.57-19.47c0 .24.08.43.23.59l.65.64c.17.18.36.27.58.27c.22 0 .42-.08.6-.25c.17-.17.26-.37.26-.61s-.08-.45-.25-.63l-.64-.61a.794.794 0 0 0-.6-.26c-.24 0-.44.08-.6.25c-.15.16-.23.37-.23.61m5.32 4.38c.67-.68 1.48-1.01 2.43-1.01c.98 0 1.82.35 2.51 1.04c.69.69 1.04 1.53 1.04 2.5c0 .65-.16 1.25-.49 1.8c-.95-.95-2.11-1.42-3.47-1.42h-.34c-.29-1.18-.85-2.15-1.68-2.91m1.6-5.23c0 .23.08.43.24.59c.16.16.35.24.59.24c.25 0 .46-.08.63-.24c.17-.16.25-.35.25-.59V1.76c0-.23-.09-.43-.26-.6A.884.884 0 0 0 19 .91c-.23 0-.43.08-.59.25c-.16.17-.24.37-.24.6v2.05zm5.49 2.27c0 .22.08.43.24.6c.37.36.78.36 1.23 0l1.43-1.43c.16-.18.24-.39.24-.64c0-.23-.08-.43-.24-.59a.807.807 0 0 0-.59-.24c-.24 0-.44.08-.6.24l-1.46 1.47c-.17.18-.25.38-.25.59m.81 11.8c0 .24.09.44.26.6l.64.65c.16.16.36.24.58.24c.21 0 .41-.08.61-.24c.16-.17.24-.39.24-.64c0-.22-.08-.41-.24-.56l-.65-.66a.882.882 0 0 0-.6-.24c-.24 0-.44.08-.6.25c-.16.16-.24.36-.24.6m1.48-6.31c0 .24.09.44.26.6c.15.17.35.25.59.25h2.05c.23 0 .43-.08.59-.25c.16-.17.24-.37.24-.6c0-.24-.08-.44-.24-.6a.791.791 0 0 0-.59-.24h-2.05c-.24 0-.44.08-.6.25c-.17.16-.25.36-.25.59"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DayHail(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.48 16.88c0 1.34.47 2.49 1.4 3.45s2.07 1.47 3.4 1.53c.12 0 .18-.06.18-.17v-1.34c0-.12-.06-.18-.18-.18c-.86-.04-1.59-.39-2.19-1.03s-.9-1.4-.9-2.26c0-.83.28-1.55.85-2.17s1.27-.97 2.1-1.07l.53-.04c.13 0 .2-.06.2-.18l.07-.55c.11-1.08.56-1.99 1.37-2.71c.81-.73 1.76-1.09 2.86-1.09c1.09 0 2.04.36 2.86 1.09c.82.73 1.29 1.63 1.4 2.71l.07.58c0 .12.06.19.17.19h1.62c.89 0 1.67.32 2.32.96c.65.64.98 1.4.98 2.27c0 .87-.3 1.62-.91 2.26c-.61.64-1.34.98-2.19 1.03c-.13 0-.2.06-.2.18v1.34c0 .11.07.17.2.17c1.34-.04 2.47-.55 3.39-1.51c.93-.97 1.39-2.12 1.39-3.46c0-.74-.14-1.41-.41-2.01c.79-.96 1.18-2.07 1.18-3.33c0-.94-.24-1.82-.71-2.63a5.409 5.409 0 0 0-1.93-1.93c-.81-.47-1.69-.71-2.63-.71c-1.56 0-2.86.58-3.9 1.75c-.8-.44-1.7-.66-2.71-.66c-1.41 0-2.67.44-3.76 1.31s-1.8 2-2.11 3.37c-1.11.26-2.02.84-2.74 1.74c-.71.92-1.07 1.95-1.07 3.1m5.34 7.06c.1.22.25.37.46.45c.2.1.41.11.63.02c.22-.08.37-.23.45-.45c.1-.22.11-.43.02-.65a.728.728 0 0 0-.45-.44a.679.679 0 0 0-.62-.03c-.22.09-.37.24-.47.47c-.1.18-.11.39-.02.63m.64-2.84c0 .14.03.27.09.38c.19.31.49.46.89.46c.32 0 .55-.22.69-.65l1.04-3.22c.08-.24.06-.47-.07-.67s-.31-.33-.55-.37a.892.892 0 0 0-.62.07a.85.85 0 0 0-.41.5l-1.03 3.22c-.02.1-.03.2-.03.28m1.87 5.62c0 .13.02.23.05.29c.09.22.24.37.45.45c.09.05.21.07.35.07c.06 0 .16-.02.3-.06c.22-.08.38-.23.47-.45s.1-.44 0-.66c-.1-.22-.25-.37-.45-.45s-.41-.08-.62 0c-.19.07-.33.19-.42.35a.95.95 0 0 0-.13.46m.61-22.15c0 .25.08.45.24.6l.65.65c.16.16.34.25.54.27c.21.03.41-.05.61-.23s.3-.4.3-.64s-.08-.44-.24-.6l-.64-.64a.895.895 0 0 0-.62-.25c-.24 0-.45.08-.61.24c-.15.17-.23.37-.23.6m.12 19.46c0 .16.05.32.16.48s.27.27.48.33c.11.02.19.04.24.04c.15 0 .28-.03.38-.08c.2-.08.34-.27.43-.57l1.8-6.14a.81.81 0 0 0-.06-.65a.762.762 0 0 0-.5-.39a.834.834 0 0 0-.66.06c-.2.11-.34.27-.41.51l-1.84 6.19c-.01.11-.02.19-.02.22m3.45-.39c0 .13.02.23.07.31c.09.21.24.35.45.44c.11.05.22.08.35.08c.06 0 .16-.02.3-.06c.23-.09.38-.23.46-.44c.08-.22.08-.43 0-.63a.839.839 0 0 0-.42-.45a.767.767 0 0 0-.66-.03c-.21.09-.37.24-.48.47a.57.57 0 0 0-.07.31m.72-2.56c0 .16.05.31.15.45c.1.15.26.25.46.31c.09.02.17.03.25.03c.39 0 .65-.2.79-.61l1.03-3.18c.08-.23.05-.45-.07-.65s-.29-.33-.52-.39c-.24-.07-.45-.05-.64.06s-.32.27-.4.51l-1.02 3.2c-.01.13-.03.22-.03.27M15.3 9c.67-.64 1.5-.97 2.48-.97s1.81.34 2.5 1.03c.69.68 1.04 1.51 1.04 2.49c0 .62-.17 1.24-.52 1.85c-.96-.96-2.12-1.44-3.51-1.44H17c-.3-1.16-.86-2.15-1.7-2.96m1.62-5.27c0 .24.08.44.25.61c.17.17.37.25.61.25c.23 0 .43-.08.59-.25c.16-.17.24-.37.24-.61V1.67c0-.24-.08-.44-.24-.61a.764.764 0 0 0-.59-.25c-.24 0-.44.08-.61.25c-.17.17-.25.37-.25.61zm5.55 2.29c0 .24.08.44.25.6c.15.17.34.26.58.26c.23 0 .44-.09.6-.26l1.44-1.44c.18-.15.27-.35.27-.6c0-.24-.09-.44-.26-.61a.838.838 0 0 0-.61-.25c-.22 0-.41.09-.57.27l-1.45 1.43c-.16.15-.25.36-.25.6m.81 11.9c0 .23.08.43.24.6l.66.63c.14.18.34.27.6.27c.24 0 .43-.09.57-.27a.77.77 0 0 0 .27-.6c0-.24-.09-.44-.27-.61l-.65-.62c-.16-.18-.35-.26-.58-.26s-.43.08-.6.25a.87.87 0 0 0-.24.61m1.46-6.37c0 .24.09.44.26.6c.18.18.38.26.62.26h2.03c.24 0 .44-.08.61-.25c.17-.17.26-.37.26-.61c0-.23-.08-.43-.25-.59a.877.877 0 0 0-.62-.24h-2.03c-.25 0-.46.08-.63.24c-.16.16-.25.36-.25.59"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DayHaze(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.37 15.62c0-.24.08-.44.25-.61c.18-.17.38-.25.6-.25h2.04c.23 0 .42.08.58.25c.15.17.23.37.23.61s-.08.44-.23.6c-.15.17-.35.25-.58.25H5.23c-.23 0-.43-.08-.6-.25a.826.826 0 0 1-.26-.6m2.86-6.91c0-.23.08-.43.23-.61c.2-.17.41-.25.64-.25c.21 0 .41.08.59.25l1.43 1.46c.16.15.24.35.24.59s-.08.44-.24.6c-.16.16-.36.24-.6.24s-.44-.08-.59-.24L7.47 9.32a.853.853 0 0 1-.24-.61m.16 9.31c0-.22.08-.41.23-.55c.16-.14.37-.22.64-.22h5.71c.27 0 .48.07.64.22c.16.14.24.33.24.55c0 .27-.08.48-.23.64c-.16.16-.37.24-.65.24H8.26c-.27 0-.48-.08-.64-.23s-.23-.38-.23-.65m.99 2.95c0-.22.09-.42.28-.6a.849.849 0 0 1 1.23-.01c.16.18.24.38.24.61c0 .28-.08.49-.23.65s-.37.23-.64.23a.87.87 0 0 1-.61-.24a.824.824 0 0 1-.27-.64m1.41-5.06v-.07c.03-1.26.47-2.35 1.31-3.28c.84-.93 1.87-1.49 3.1-1.69h.05c.19-.04.45-.06.76-.06s.57.02.76.06h.04c1.22.19 2.26.76 3.1 1.69c.84.93 1.28 2.02 1.31 3.28v.07c0 .16-.08.24-.23.24h-1.13c-.12 0-.2-.03-.25-.09a.274.274 0 0 1-.07-.18c-.04-.93-.4-1.72-1.08-2.37c-.68-.65-1.49-.97-2.44-.97s-1.76.32-2.44.97c-.68.65-1.04 1.44-1.08 2.37c0 .06-.03.12-.08.18s-.14.09-.26.09h-1.13c-.16 0-.24-.08-.24-.24m1.01 5.06c0-.23.08-.43.24-.61c.16-.17.37-.26.63-.26h3.83c.22 0 .42.09.6.27c.18.18.28.38.28.6c0 .26-.09.48-.27.64s-.38.24-.61.24h-3.83c-.27 0-.48-.08-.64-.23c-.16-.16-.23-.37-.23-.65m3.34-13.08V5.86c0-.24.08-.44.25-.61S14.76 5 15 5s.43.08.6.25c.17.17.25.37.25.61v2.03c0 .23-.08.43-.25.58s-.37.24-.6.24s-.44-.08-.6-.23s-.26-.36-.26-.59m1.36 10.13c0-.21.09-.39.27-.54s.38-.23.61-.23s.43.08.61.23s.27.33.27.54c0 .26-.09.48-.27.64c-.18.16-.38.24-.61.24s-.44-.08-.61-.24a.838.838 0 0 1-.27-.64m1.54 2.95c0-.23.08-.43.24-.61c.16-.17.38-.26.64-.26h1.86c.26 0 .47.09.63.26c.16.18.24.38.24.61c0 .28-.08.49-.23.65c-.15.16-.37.23-.64.23h-1.86c-.27 0-.48-.08-.64-.23c-.16-.16-.24-.37-.24-.65m.88-2.95c0-.21.09-.39.27-.54s.38-.23.6-.23h3.07c.22 0 .4.07.54.22s.22.33.22.55c0 .27-.07.48-.21.64c-.14.16-.32.24-.55.24h-3.07c-.23 0-.43-.08-.61-.24a.867.867 0 0 1-.26-.64m1.74-7.87c0-.25.08-.45.23-.59l1.42-1.46c.18-.17.38-.25.59-.25c.23 0 .43.08.6.25c.17.17.25.37.25.61s-.08.45-.24.61l-1.46 1.42c-.18.16-.38.24-.6.24c-.23 0-.41-.08-.56-.24a.807.807 0 0 1-.23-.59m2.26 5.47c0-.23.08-.43.24-.61c.17-.17.36-.25.57-.25h2.02c.23 0 .43.09.6.26c.17.17.26.37.26.6c0 .23-.09.43-.26.6c-.17.17-.37.25-.6.25h-2.02c-.23 0-.43-.08-.58-.25s-.23-.36-.23-.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DayLightWind(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.32 14.85c0 .24.09.44.26.6c.16.17.36.25.6.25h9.42c.23 0 .43-.08.59-.25c.16-.17.24-.37.24-.6a.84.84 0 0 0-.23-.59a.791.791 0 0 0-.59-.24H3.18c-.24 0-.44.08-.61.24c-.17.16-.25.36-.25.59M2.65 21c0 .24.08.44.25.6c.16.17.36.25.6.25h9.43c.24 0 .44-.08.61-.25c.17-.17.25-.37.25-.6s-.08-.43-.25-.59s-.37-.24-.61-.24H3.51c-.24 0-.44.08-.6.24c-.17.16-.26.36-.26.59m1.37-3.1c0 .24.08.43.25.59c.17.16.38.23.63.23h9.4c.24 0 .43-.08.59-.23c.16-.16.23-.35.23-.59c0-.25-.08-.45-.23-.61a.791.791 0 0 0-.59-.24H4.9c-.25 0-.46.08-.63.24c-.16.16-.25.36-.25.61m2.43-6.35c0-.24.09-.44.26-.62c.17-.16.38-.24.6-.24h2.03c.23 0 .42.08.58.25c.16.17.23.37.23.61s-.08.44-.23.6c-.16.17-.35.25-.58.25H7.31c-.24 0-.44-.08-.61-.25a.816.816 0 0 1-.25-.6m2.86-6.92c0-.22.08-.43.24-.61c.19-.16.4-.24.63-.24c.22 0 .42.08.59.24l1.42 1.47c.16.15.24.35.24.59s-.08.44-.24.6c-.16.16-.36.24-.6.24s-.44-.08-.59-.24L9.55 5.25a.894.894 0 0 1-.24-.62m2.55 6.8v-.07c.02-.91.27-1.75.74-2.53c.47-.77 1.11-1.38 1.9-1.83s1.65-.67 2.57-.67c.7 0 1.37.14 2.02.42c.64.28 1.2.65 1.66 1.12c.47.47.84 1.02 1.11 1.66c.27.64.41 1.32.41 2.02c0 .94-.23 1.8-.69 2.6s-1.09 1.43-1.88 1.89a5.1 5.1 0 0 1-2.6.71h-.2c-.07 0-.13-.02-.18-.07a.25.25 0 0 1-.08-.18v-1.22c0-.13.07-.2.22-.2h.24c.96-.01 1.79-.35 2.47-1.05c.68-.69 1.03-1.52 1.03-2.49c0-.96-.35-1.78-1.04-2.47c-.69-.68-1.52-1.02-2.5-1.02c-.94 0-1.76.32-2.44.98c-.68.65-1.04 1.44-1.08 2.37c0 .06-.02.11-.07.17s-.13.09-.25.09h-1.14c-.15.01-.22-.07-.22-.23m4.37 9.88v-1.99c0-.24.08-.44.24-.6s.36-.24.6-.24s.45.08.61.24s.24.36.24.6v1.99c0 .24-.08.45-.25.62c-.17.17-.37.25-.6.25c-.24 0-.44-.08-.6-.25s-.24-.37-.24-.62m0-17.48V1.78c0-.24.08-.44.25-.6s.36-.25.6-.25c.23 0 .43.08.6.25s.25.37.25.6v2.04c0 .23-.08.42-.25.58c-.17.15-.37.23-.6.23c-.24 0-.44-.08-.6-.23a.723.723 0 0 1-.25-.57m5.51 13.18c0-.23.07-.42.23-.56a.74.74 0 0 1 .57-.23c.24 0 .44.08.6.23l1.45 1.42c.16.17.24.38.24.61c0 .23-.08.43-.24.59c-.4.31-.8.31-1.2 0l-1.42-1.43a.904.904 0 0 1-.23-.63m0-10.93c0-.25.07-.45.23-.59l1.42-1.47c.18-.16.37-.24.59-.24c.24 0 .44.08.6.25c.17.17.25.37.25.6c0 .25-.08.46-.24.62l-1.45 1.43c-.18.16-.38.24-.6.24c-.23 0-.41-.08-.57-.24s-.23-.36-.23-.6M24 11.55c0-.23.08-.44.25-.62c.16-.16.35-.24.56-.24h2.03c.23 0 .43.09.61.26c.17.17.26.37.26.6c0 .23-.09.43-.26.6c-.18.17-.38.25-.61.25h-2.03c-.23 0-.42-.08-.58-.25a.847.847 0 0 1-.23-.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DayLightning(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.56 16.9c0 1.33.46 2.47 1.39 3.43c.93.96 2.06 1.47 3.4 1.53c.11 0 .17-.06.17-.17v-1.37c0-.12-.06-.18-.17-.18c-.87-.07-1.6-.41-2.19-1.04c-.59-.62-.89-1.36-.89-2.21c0-.84.28-1.57.85-2.19s1.26-.97 2.1-1.04l.52-.08c.13 0 .2-.06.2-.17L7 12.9c.11-1.08.56-1.99 1.37-2.71c.81-.73 1.76-1.09 2.86-1.09c1.09 0 2.04.36 2.85 1.09c.81.72 1.27 1.63 1.39 2.72l.07.58c0 .11.06.17.19.17h1.6c.91 0 1.68.32 2.32.95c.64.63.96 1.39.96 2.28c0 .85-.3 1.59-.89 2.21c-.59.62-1.32.97-2.19 1.04c-.13 0-.2.06-.2.18v1.37c0 .11.07.17.2.17c1.33-.04 2.46-.55 3.38-1.51s1.38-2.11 1.38-3.45c0-.68-.16-1.37-.47-2.07c.78-.94 1.18-2.03 1.18-3.27c0-.71-.14-1.39-.42-2.04c-.28-.65-.65-1.2-1.12-1.67s-1.03-.84-1.67-1.12c-.65-.27-1.32-.41-2.03-.41c-1.54 0-2.84.58-3.88 1.73c-.86-.41-1.74-.62-2.65-.62c-1.42 0-2.67.43-3.76 1.3s-1.79 1.99-2.1 3.37c-1.1.26-2.01.83-2.73 1.73s-1.08 1.92-1.08 3.07m7.5 11.27h.3l5.32-7.85c.04-.04.05-.09.02-.14s-.07-.07-.14-.07h-2.18l2.3-4.21c.07-.14.03-.22-.14-.22H11.6c-.08 0-.15.05-.22.14l-2.16 5.72c-.02.14.02.21.14.21h2.17zm.88-23.58c0 .25.08.46.24.62l.66.65c.42.32.82.32 1.21 0c.16-.18.24-.39.24-.64c0-.23-.08-.43-.24-.59l-.64-.65a.895.895 0 0 0-.62-.25c-.23 0-.43.08-.6.25c-.17.17-.25.38-.25.61m5.34 4.43C15.96 8.34 16.79 8 17.76 8c.98 0 1.81.35 2.49 1.04c.69.69 1.03 1.53 1.03 2.52c0 .61-.17 1.21-.51 1.81c-.98-.94-2.13-1.41-3.46-1.41H17c-.26-1.14-.84-2.12-1.72-2.94m1.62-5.29c0 .25.08.46.24.62c.16.16.36.24.61.24c.24 0 .43-.08.59-.24c.16-.16.23-.37.23-.62V1.69c0-.24-.08-.43-.23-.59a.784.784 0 0 0-.59-.23c-.24 0-.44.08-.6.23s-.25.35-.25.59zm5.52 2.32c0 .23.09.43.27.6c.18.17.37.25.55.25c.16 0 .37-.08.62-.25l1.44-1.43c.17-.18.25-.39.25-.63s-.08-.45-.25-.61a.847.847 0 0 0-.6-.24c-.22 0-.41.08-.58.25l-1.43 1.43c-.18.19-.27.4-.27.63m.81 11.82c0 .23.08.43.24.61l.65.66c.16.16.36.24.58.24c.24 0 .44-.08.6-.25c.17-.17.25-.38.25-.63c0-.23-.08-.42-.25-.57l-.62-.66a.92.92 0 0 0-.62-.23c-.23 0-.43.08-.59.24c-.16.16-.24.36-.24.59m1.45-6.31c0 .22.09.41.26.57c.17.17.37.25.6.25h2.04c.24 0 .44-.08.61-.24c.17-.16.25-.35.25-.59s-.09-.44-.26-.61s-.37-.25-.6-.25h-2.04c-.23 0-.43.08-.6.25a.87.87 0 0 0-.26.62"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DayRain(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.51 16.89c0 1.33.46 2.48 1.39 3.44s2.06 1.47 3.41 1.53c.11 0 .17-.06.17-.17v-1.34c0-.11-.06-.17-.17-.17c-.86-.04-1.59-.39-2.19-1.03s-.9-1.4-.9-2.26c0-.82.28-1.54.85-2.16s1.27-.97 2.1-1.07l.53-.05c.13 0 .2-.06.2-.17l.07-.54c.11-1.08.56-1.99 1.37-2.72s1.76-1.1 2.85-1.1s2.04.37 2.86 1.1s1.28 1.64 1.4 2.72l.07.57c0 .12.06.19.17.19h1.62c.89 0 1.65.32 2.3.96c.65.64.97 1.39.97 2.27c0 .87-.3 1.62-.9 2.26c-.6.64-1.33.98-2.18 1.03c-.12 0-.19.06-.19.17v1.34c0 .11.06.17.19.17c1.33-.04 2.46-.55 3.39-1.51c.93-.97 1.39-2.12 1.39-3.45c0-.72-.14-1.39-.42-2.01c.78-.97 1.17-2.07 1.17-3.31c0-.95-.24-1.83-.71-2.64s-1.11-1.45-1.92-1.92s-1.68-.7-2.62-.7c-1.56 0-2.85.58-3.88 1.74c-.82-.44-1.72-.66-2.71-.66c-1.41 0-2.67.44-3.76 1.32s-1.79 2-2.1 3.36c-1.11.26-2.02.84-2.74 1.74s-1.08 1.92-1.08 3.07m5.4 6.86c0 .17.05.33.16.49c.11.16.27.27.49.33c.11.02.2.04.27.04c.39 0 .65-.21.77-.64l1.58-5.88c.07-.24.04-.46-.08-.67a.767.767 0 0 0-.53-.38a.737.737 0 0 0-.63.07c-.2.11-.34.28-.41.51l-1.58 5.91c-.02.13-.04.2-.04.22m2.61 3.08c0 .19.05.36.15.52c.1.16.27.26.52.3c.11.02.2.04.26.04c.16 0 .31-.06.45-.17c.14-.12.23-.28.27-.48l2.4-8.93c.06-.23.04-.45-.07-.64s-.28-.33-.5-.4a.777.777 0 0 0-.65.07c-.2.11-.34.28-.4.51l-2.4 8.93c-.02.15-.03.24-.03.25M9.94 4.6c0 .24.08.44.25.61l.65.66c.19.15.4.22.62.22c.21 0 .41-.08.58-.23c.17-.16.26-.35.26-.59s-.08-.46-.24-.64l-.64-.65a.854.854 0 0 0-1.22.01c-.17.17-.26.38-.26.61m3.73 19.17a.822.822 0 0 0 .61.79c.11.02.2.04.25.04c.17 0 .34-.05.49-.15c.15-.1.25-.26.3-.49l1.58-5.88a.835.835 0 0 0-.57-1.04a.773.773 0 0 0-.65.07a.78.78 0 0 0-.39.51l-1.58 5.91c-.02.15-.04.23-.04.24M15.3 9.03c.71-.67 1.53-1 2.48-1c.98 0 1.82.35 2.5 1.04c.69.69 1.03 1.53 1.03 2.52c0 .62-.17 1.24-.52 1.85c-.97-.97-2.13-1.45-3.49-1.45h-.33a5.763 5.763 0 0 0-1.67-2.96m1.62-5.25c0 .23.08.43.25.59c.17.16.37.24.61.24c.23 0 .43-.08.59-.23c.16-.16.24-.35.24-.59V1.73c0-.26-.08-.47-.23-.63a.791.791 0 0 0-.59-.24c-.25 0-.46.08-.62.25s-.24.37-.24.62v2.05zm5.53 2.28c0 .24.09.44.27.59c.14.16.32.24.55.26c.23.02.44-.07.62-.26l1.44-1.43c.18-.17.26-.38.26-.63c0-.24-.08-.45-.25-.61a.853.853 0 0 0-.61-.24c-.21 0-.4.08-.58.25l-1.43 1.44c-.18.17-.27.38-.27.63m.81 11.85c0 .24.08.45.24.63l.65.63c.18.14.38.21.6.21l.02.02c.23 0 .42-.08.58-.24c.16-.16.24-.37.24-.61s-.09-.43-.26-.58l-.62-.66c-.18-.16-.39-.24-.62-.24s-.43.08-.59.25s-.24.35-.24.59m1.46-6.33c0 .24.09.43.26.59c.18.18.38.26.62.26h2.03c.24 0 .44-.08.61-.25c.17-.17.25-.37.25-.6c0-.24-.08-.44-.25-.61s-.37-.26-.61-.26H25.6a.87.87 0 0 0-.62.26c-.18.17-.26.37-.26.61"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DayRainMix(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.48 16.95c0 1.32.46 2.46 1.37 3.4c.91.94 2.04 1.45 3.38 1.51c.12 0 .18-.06.18-.17v-1.33c0-.12-.06-.18-.18-.18c-.86-.04-1.58-.38-2.17-1s-.88-1.37-.88-2.24c0-.84.28-1.58.84-2.19c.56-.62 1.26-.96 2.1-1.03l.53-.08c.11 0 .16-.05.16-.14l.08-.55c.12-1.09.59-2 1.38-2.72S10 9.16 11.1 9.16s2.05.36 2.86 1.08c.82.72 1.28 1.62 1.38 2.69l.07.58c.02.11.1.17.22.17h1.6c.89 0 1.65.32 2.29.96c.64.64.96 1.41.96 2.31c0 .87-.29 1.61-.88 2.24s-1.31.95-2.17 1c-.13 0-.2.06-.2.18v1.33c0 .11.07.17.2.17c1.33-.04 2.45-.54 3.37-1.49c.91-.95 1.37-2.09 1.37-3.42c0-.68-.13-1.34-.38-2.01c.78-.96 1.16-2.08 1.16-3.35c0-.71-.14-1.38-.41-2.03c-.27-.65-.65-1.2-1.12-1.67s-1.03-.84-1.67-1.12s-1.33-.42-2.04-.42c-1.54 0-2.83.58-3.86 1.74c-.89-.44-1.81-.66-2.74-.66c-1.41 0-2.66.44-3.74 1.31s-1.77 2-2.08 3.38c-1.12.26-2.04.83-2.75 1.73s-1.06 1.93-1.06 3.09m5.35 7.03c0 .17.05.34.16.51c.11.17.27.28.47.35c.23.07.45.06.64-.04c.2-.09.33-.28.4-.56l.14-.61c.05-.23.02-.44-.1-.63a.875.875 0 0 0-.52-.4c-.23-.07-.44-.04-.64.08s-.32.29-.38.52l-.14.59c-.02.06-.03.12-.03.19m.77-2.9c0 .22.08.41.24.57c.16.18.35.26.56.26c.24 0 .44-.08.6-.24c.17-.16.25-.35.25-.59c0-.23-.08-.43-.25-.59a.847.847 0 0 0-.6-.24c-.23 0-.42.08-.58.23s-.22.37-.22.6m.61-2.27c-.01.16.03.31.13.45c.1.15.26.25.48.32a.8.8 0 0 0 .62-.07c.21-.11.35-.28.42-.51l.27-.9c.07-.24.05-.46-.07-.65a.913.913 0 0 0-.54-.39a.737.737 0 0 0-.63.07a.85.85 0 0 0-.41.5l-.24.92c-.02.16-.03.25-.03.26m1.15 8.29c0 .17.05.33.16.49c.11.16.27.27.49.33c.09.02.17.03.24.03c.43 0 .7-.2.8-.61l.13-.59c.06-.26.03-.48-.08-.68s-.29-.32-.52-.37a.744.744 0 0 0-.63.07c-.21.12-.34.29-.41.51l-.14.6c-.03.11-.04.19-.04.22m.56-22.44c0 .24.08.44.24.6l.66.64c.14.16.32.24.54.26c.22.02.43-.07.62-.26c.16-.16.24-.36.24-.59c0-.24-.08-.44-.24-.6l-.63-.65a.869.869 0 0 0-.58-.26c-.23 0-.43.08-.6.25c-.16.17-.25.37-.25.61m.23 19.54c0 .23.08.42.24.58c.16.16.36.24.58.24c.24 0 .43-.08.59-.23c.16-.16.23-.35.23-.59a.784.784 0 0 0-.82-.81c-.24 0-.43.08-.59.23s-.23.35-.23.58m.62-2.27c-.01.15.03.31.14.47c.1.16.25.26.45.3c.23.06.44.04.64-.06s.33-.29.41-.56l.26-.9c.07-.22.05-.43-.07-.63a.867.867 0 0 0-.53-.4c-.22-.07-.43-.04-.64.08s-.34.3-.41.53l-.22.9c-.02.08-.03.17-.03.27m2.76 2.15c0 .17.05.33.15.48c.1.15.25.26.46.32c.03 0 .08.01.14.02c.06.01.11.02.14.02c.41 0 .66-.22.76-.66l.14-.6c.07-.21.05-.42-.07-.63a.863.863 0 0 0-.51-.41c-.25-.06-.48-.04-.68.08s-.34.29-.41.53l-.09.59c0 .01 0 .05-.01.11c-.01.07-.02.11-.02.15m.74-2.96c0 .23.08.42.24.57c.15.16.34.24.58.24s.43-.08.59-.23c.16-.16.23-.35.23-.58c0-.24-.08-.43-.23-.59c-.16-.16-.35-.23-.59-.23s-.43.08-.59.23s-.23.35-.23.59m.61-2.31c0 .17.05.33.16.48c.11.15.27.26.49.32c.02 0 .06.01.12.02s.11.02.14.02c.11 0 .23-.03.37-.09c.21-.11.34-.28.4-.52l.24-.9c.06-.23.04-.44-.07-.63s-.28-.33-.5-.4a.787.787 0 0 0-.64.06c-.2.11-.33.27-.4.51l-.28.91c0 .02 0 .05-.01.11a.36.36 0 0 0-.02.11m.35-9.72c.66-.66 1.48-.99 2.47-.99c.99 0 1.83.34 2.52 1.02s1.04 1.5 1.04 2.48c0 .66-.18 1.29-.53 1.88c-.98-.98-2.15-1.47-3.5-1.47h-.31c-.28-1.1-.85-2.07-1.69-2.92m1.65-5.26c0 .23.08.42.23.58c.15.15.35.23.59.23s.45-.08.62-.23c.17-.15.25-.35.25-.58V1.76c0-.23-.09-.43-.26-.6a.838.838 0 0 0-.61-.25c-.23 0-.43.08-.58.25c-.16.17-.23.37-.23.6v2.07zm5.52 2.26c0 .25.08.45.23.6c.36.36.76.36 1.21 0l1.43-1.43c.17-.17.25-.38.25-.63c0-.24-.08-.44-.25-.6a.816.816 0 0 0-.6-.25c-.23 0-.43.08-.61.24L22.63 5.5a.79.79 0 0 0-.23.59m.78 11.85c0 .23.09.43.27.59l.61.63c.2.16.4.24.61.24c.2 0 .4-.08.6-.24c.17-.16.25-.35.25-.59c0-.23-.08-.43-.25-.62l-.65-.61c-.15-.17-.34-.25-.57-.25s-.43.08-.6.25a.8.8 0 0 0-.27.6m1.48-6.34c0 .24.09.43.26.59c.18.18.39.27.62.27h2.03c.23 0 .43-.08.59-.25c.16-.17.24-.37.24-.61s-.08-.44-.24-.6a.764.764 0 0 0-.59-.25h-2.03c-.24 0-.44.08-.62.25s-.26.37-.26.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DayRainWind(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.48 16.91c0 1.12.33 2.13 1 3.01c.67.88 1.53 1.47 2.58 1.77c.09.01.17-.01.24-.08l1.17-1.41c-.89 0-1.66-.32-2.3-.97s-.96-1.42-.96-2.33c0-.83.28-1.55.85-2.17c.57-.61 1.27-.97 2.11-1.07l.49-.03c.13 0 .2-.06.2-.19l.07-.54c.11-1.08.57-1.99 1.38-2.72s1.77-1.1 2.86-1.1c1.1 0 2.06.37 2.88 1.1s1.28 1.64 1.39 2.72l.07.59c.04.11.12.17.22.17h1.64c.88 0 1.64.32 2.28.95c.64.63.96 1.4.96 2.28c0 .84-.28 1.58-.85 2.21c-.57.63-1.25.98-2.07 1.05c-.49.06-.79.15-.9.28l-2.19 2.82a.84.84 0 0 0-.16.63c.03.23.13.42.31.56c.15.17.35.24.6.21s.45-.15.6-.36l1.92-2.46c.82-.08 1.57-.35 2.25-.82c.68-.46 1.21-1.06 1.61-1.79c.39-.73.59-1.51.59-2.34c0-.72-.14-1.39-.42-2.01c.79-.96 1.18-2.07 1.18-3.33c0-.71-.14-1.38-.42-2.03a5.34 5.34 0 0 0-1.13-1.68c-.47-.47-1.03-.84-1.68-1.12c-.65-.28-1.33-.42-2.04-.42c-1.5 0-2.79.58-3.88 1.74c-.79-.44-1.7-.66-2.72-.66c-1.41 0-2.66.44-3.75 1.32s-1.78 2-2.07 3.37c-1.13.26-2.06.83-2.79 1.73s-1.12 1.97-1.12 3.12m3.9 8.24c.07.22.23.38.46.48c.22.1.44.1.67.01s.38-.25.46-.47a.71.71 0 0 0 0-.62a.879.879 0 0 0-.45-.46a.8.8 0 0 0-.66-.02c-.22.09-.37.24-.45.46c-.09.25-.1.46-.03.62m1.81-2.74v.11c.02.22.12.4.3.55c.18.15.38.22.63.2c.24-.02.43-.12.57-.29l2.18-2.82c.14-.18.21-.39.19-.63s-.12-.43-.29-.57a.835.835 0 0 0-.63-.19c-.24.02-.44.13-.6.31l-2.15 2.8c-.14.16-.2.34-.2.53m.51 4.86c0 .06.02.16.06.3c.09.22.24.37.46.46c.11.04.23.07.36.07c.09 0 .2-.02.31-.05c.21-.08.36-.23.44-.44c.1-.22.11-.44.02-.67a.712.712 0 0 0-.45-.46c-.22-.1-.44-.1-.67-.01s-.38.24-.45.45c-.05.11-.08.23-.08.35m1.67-2.37v.1c.02.23.13.42.33.58c.15.16.35.23.6.2a.84.84 0 0 0 .6-.34l4.14-5.23c.14-.2.19-.41.17-.64s-.13-.42-.31-.59a.81.81 0 0 0-.62-.17c-.23.03-.42.13-.56.31l-4.16 5.26c-.13.15-.19.33-.19.52m.58-20.29c0 .24.09.44.27.59l.61.66c.16.16.34.24.54.26c.21.03.41-.04.61-.22s.3-.39.3-.64c0-.22-.08-.43-.24-.6l-.64-.65a.754.754 0 0 0-.58-.24c-.24 0-.45.08-.62.24c-.16.17-.25.37-.25.6m2.95 21.43c0 .12.02.22.07.31c.09.22.24.37.45.46c.11.05.22.08.35.08c.06 0 .16-.02.3-.06c.22-.09.38-.24.46-.46a.8.8 0 0 0 .02-.66a.719.719 0 0 0-.45-.45a.823.823 0 0 0-.67-.02a.76.76 0 0 0-.46.45c-.05.11-.07.23-.07.35m2.4-17.01c.71-.64 1.53-.97 2.48-.97c.99 0 1.83.34 2.53 1.03s1.05 1.51 1.05 2.49c0 .66-.17 1.28-.52 1.86c-.97-.97-2.13-1.45-3.48-1.45h-.34a5.726 5.726 0 0 0-1.72-2.96m1.63-5.26c0 .23.08.43.25.59c.17.16.37.24.6.24c.24 0 .45-.08.62-.24c.17-.16.25-.36.25-.59V1.7c0-.23-.09-.43-.26-.6a.838.838 0 0 0-.61-.25c-.23 0-.43.08-.6.25s-.25.37-.25.6zm5.57 2.28c0 .24.07.44.22.6c.18.17.38.25.61.26c.23 0 .42-.08.57-.26l1.5-1.45c.16-.16.24-.37.24-.61s-.09-.44-.26-.61a.826.826 0 0 0-.6-.26c-.22 0-.42.08-.6.25l-1.45 1.47a.89.89 0 0 0-.23.61m.78 11.9c0 .24.08.44.24.6l.66.63c.16.16.34.25.54.27h.06c.18 0 .38-.09.62-.27c.16-.16.24-.36.24-.58c0-.24-.08-.45-.24-.63l-.66-.62a.779.779 0 0 0-.59-.25c-.24 0-.45.08-.62.25a.84.84 0 0 0-.25.6m1.5-6.37c0 .24.08.44.24.6c.16.18.36.27.6.27h2.04c.24 0 .44-.09.61-.26c.17-.17.25-.38.25-.61c0-.23-.08-.43-.25-.59a.853.853 0 0 0-.61-.24h-2.04c-.24 0-.44.08-.6.24c-.16.16-.24.36-.24.59"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DayShowers(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.56 16.88c0 1.33.46 2.47 1.39 3.43s2.06 1.46 3.4 1.53c.11 0 .17-.06.17-.17v-1.34c0-.12-.06-.18-.17-.18c-.85-.04-1.57-.38-2.17-1.02s-.89-1.39-.89-2.25c0-.84.28-1.56.84-2.17s1.26-.96 2.1-1.06l.5-.04c.13 0 .2-.06.2-.18l.06-.53c.11-1.08.56-1.99 1.37-2.71c.81-.73 1.76-1.09 2.86-1.09c1.09 0 2.05.36 2.86 1.09c.81.73 1.27 1.63 1.37 2.71l.07.57c0 .12.06.18.18.18h1.67c.88 0 1.63.32 2.27.96c.64.64.96 1.39.96 2.27c0 .85-.3 1.59-.9 2.22s-1.32.98-2.16 1.05c-.11 0-.17.06-.17.18v1.34c0 .11.06.17.17.17c.88-.02 1.67-.26 2.4-.72s1.3-1.05 1.71-1.8c.42-.75.62-1.56.62-2.44c0-.71-.14-1.37-.41-1.96c.76-.94 1.13-2.03 1.13-3.28c0-.71-.14-1.39-.41-2.04c-.27-.65-.65-1.2-1.12-1.67a5.04 5.04 0 0 0-1.66-1.11c-.65-.28-1.33-.41-2.04-.41c-1.51 0-2.78.55-3.81 1.66c-.79-.43-1.7-.64-2.73-.64c-1.41 0-2.66.44-3.75 1.31S5.7 10.73 5.4 12.09c-1.12.26-2.05.83-2.77 1.72c-.71.89-1.07 1.92-1.07 3.07m5.41 6.7c0 .18.05.36.16.53c.11.18.26.29.45.36s.4.05.61-.06c.22-.11.36-.29.44-.55l.25-1.05c.07-.21.05-.41-.07-.62a.88.88 0 0 0-.51-.42c-.25-.06-.47-.03-.67.08s-.32.3-.37.53l-.28.99c0 .05-.01.12-.01.21m1.31-4.72c0 .38.21.64.64.79c.22.08.43.06.64-.05c.21-.11.34-.29.41-.53l.24-1.03c.07-.21.05-.41-.07-.62a.824.824 0 0 0-.51-.42c-.24-.06-.47-.04-.67.08s-.32.29-.37.52l-.3 1.02c0 .08-.01.16-.01.24m1.22 7.89c0 .16.06.33.17.5c.11.17.28.29.49.36c.01 0 .04 0 .1.01c.06.01.11.01.15.01c.14 0 .26-.02.37-.07c.19-.08.33-.27.41-.58l.27-.99c.07-.23.05-.45-.07-.65a.847.847 0 0 0-.51-.4a.777.777 0 0 0-.65.07c-.2.12-.34.29-.4.51l-.28 1.02c-.04.09-.05.16-.05.21m.46-22.07c0 .25.08.46.25.62l.66.65c.34.34.73.34 1.17 0a.87.87 0 0 0 .24-.61c0-.23-.08-.43-.24-.61l-.63-.66a.814.814 0 0 0-.6-.24c-.23 0-.43.08-.6.25c-.17.16-.25.36-.25.6m.89 17.28c0 .17.05.34.16.51c.11.17.26.28.47.35c.23.07.44.05.64-.05c.19-.1.33-.29.4-.56l.24-1.01c.07-.23.05-.45-.06-.65a.85.85 0 0 0-1.18-.33c-.2.12-.33.3-.37.53l-.28 1.03c-.02.03-.02.09-.02.18m2.78 1.72c.02.38.23.65.63.83l.25.04c.16 0 .32-.05.47-.16c.15-.11.26-.27.32-.5l.29-1.01a.84.84 0 0 0-.09-.66c-.12-.2-.3-.33-.53-.37c-.21-.07-.41-.05-.62.07s-.34.29-.41.51l-.27 1.02c-.01.02-.01.05-.02.08s-.01.06-.02.08s0 .05 0 .07m1.4-4.76c0 .16.05.32.15.48c.1.16.25.27.45.32l.25.03c.19 0 .37-.06.52-.18s.24-.28.28-.47l.27-.99c.07-.24.05-.45-.07-.65a.78.78 0 0 0-.51-.39a.787.787 0 0 0-.64.06c-.2.11-.33.28-.39.5l-.3 1.06c0 .08-.01.16-.01.23m.33-9.86c.66-.64 1.46-.96 2.4-.96c.98 0 1.82.35 2.51 1.04c.69.69 1.04 1.53 1.04 2.51c0 .56-.16 1.15-.47 1.76c-.96-.96-2.11-1.43-3.47-1.43h-.34a5.706 5.706 0 0 0-1.67-2.92m1.54-5.23a.821.821 0 0 0 .85.85a.821.821 0 0 0 .85-.85V1.81c0-.25-.08-.46-.24-.62c-.16-.16-.36-.24-.61-.24s-.45.08-.61.24c-.16.16-.24.37-.24.62zm5.55 2.29c0 .25.08.45.23.61c.21.17.41.25.62.25c.19 0 .38-.08.59-.25l1.43-1.43c.16-.18.24-.39.24-.63s-.08-.44-.24-.6c-.16-.16-.36-.24-.59-.24s-.43.08-.61.24L22.68 5.5c-.15.18-.23.38-.23.62m.79 11.83c0 .23.09.44.26.63l.62.64c.21.21.41.31.62.31c.19 0 .39-.1.58-.31c.18-.18.27-.39.26-.61a.867.867 0 0 0-.26-.6l-.65-.66a.779.779 0 0 0-.57-.24c-.24 0-.44.08-.61.25c-.16.16-.25.36-.25.59m1.47-6.31c0 .22.08.42.24.58c.16.16.36.24.58.24h2.04c.26 0 .47-.08.63-.23c.16-.16.24-.35.24-.59c0-.25-.08-.46-.25-.62a.86.86 0 0 0-.62-.24h-2.04c-.23 0-.43.08-.59.25c-.15.17-.23.38-.23.61"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DaySleet(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.49 16.92c0-1.17.36-2.19 1.08-3.09s1.64-1.48 2.74-1.74c.31-1.37 1.01-2.49 2.1-3.38c1.1-.88 2.35-1.32 3.77-1.32c.99 0 1.9.22 2.72.66c.5-.53 1.09-.95 1.76-1.25c.67-.3 1.37-.45 2.09-.45c.95 0 1.83.24 2.64.71s1.45 1.11 1.92 1.92c.47.81.71 1.69.71 2.64c0 1.23-.38 2.33-1.14 3.29c.29.61.43 1.28.43 2.02c0 .88-.21 1.7-.64 2.45c-.42.75-1 1.36-1.74 1.81c-.73.46-1.54.7-2.42.72c-.13 0-.2-.06-.2-.18v-1.34c0-.12.07-.18.2-.18c.86-.04 1.58-.39 2.18-1.03c.6-.64.9-1.4.9-2.26c0-.89-.32-1.65-.97-2.29s-1.41-.96-2.31-.96H15.7c-.12 0-.18-.06-.18-.17l-.08-.59c-.11-1.08-.58-1.99-1.4-2.72c-.82-.73-1.78-1.1-2.86-1.1c-1.1 0-2.05.37-2.86 1.1c-.81.73-1.27 1.64-1.37 2.72l-.07.59c-.03.09-.11.13-.22.13l-.51.04c-.84.1-1.55.45-2.11 1.06s-.84 1.34-.84 2.18v.05h.03c.01.98.38 1.78 1.11 2.43c.22.19.47.36.74.49v.02c.41.19.82.29 1.21.31c.11 0 .17.06.17.17v1.34c0 .11-.06.17-.17.17c-.52-.03-1.01-.13-1.48-.3v.02c-.83-.29-1.54-.77-2.11-1.43s-.95-1.44-1.11-2.31v-.04c-.01-.01-.01-.02-.01-.03c-.07-.26-.09-.55-.09-.88m5.5 7.17c0-.03.01-.07.02-.13c.01-.05.02-.09.02-.12l.09-.59c.07-.24.2-.41.41-.53c.2-.12.43-.14.68-.08c.23.07.39.2.51.41c.11.2.13.41.07.62l-.14.6c-.1.44-.35.66-.76.66c-.03 0-.08 0-.15-.01s-.11-.01-.13-.01a.761.761 0 0 1-.46-.33a.82.82 0 0 1-.16-.49m.74-2.94c0-.24.08-.43.23-.59c.16-.16.35-.23.59-.23s.43.08.59.23c.16.16.23.35.23.59c0 .23-.08.42-.23.58s-.35.23-.59.23c-.23 0-.42-.08-.57-.24a.752.752 0 0 1-.25-.57m1.64 5.98c0-.04.01-.11.04-.23l.13-.58c.07-.23.21-.4.41-.51c.21-.12.42-.14.63-.07c.23.04.41.17.53.37c.12.2.15.43.08.68l-.13.59c-.1.41-.37.61-.8.61c-.05 0-.13-.01-.24-.04a.831.831 0 0 1-.49-.33a.842.842 0 0 1-.16-.49M9.9 4.62c0-.24.08-.44.25-.6c.17-.16.38-.24.63-.24c.24 0 .44.08.6.24l.63.66c.17.17.25.37.25.6c0 .24-.1.46-.3.64c-.2.18-.4.26-.61.23a.875.875 0 0 1-.55-.27l-.65-.66a.857.857 0 0 1-.25-.6m.26 19.61a.784.784 0 0 1 .82-.81c.24 0 .43.08.59.23c.16.16.23.35.23.58c0 .24-.08.43-.23.59c-.16.16-.35.23-.59.23c-.22 0-.41-.08-.58-.25a.779.779 0 0 1-.24-.57m.62-2.27c0-.09.01-.18.03-.26l.23-.9c.07-.23.21-.41.41-.53c.21-.12.42-.15.64-.08c.24.07.41.2.53.4c.12.2.14.4.07.62l-.26.9c-.08.27-.22.46-.41.57s-.41.12-.64.06a.657.657 0 0 1-.45-.3c-.12-.18-.17-.33-.15-.48m2.76 2.14c0-.03 0-.07.01-.13s.01-.09.01-.11l.09-.59c.07-.24.2-.41.41-.53c.2-.12.43-.14.68-.08c.23.07.4.21.51.41c.12.21.14.42.07.63l-.14.6c-.1.43-.35.65-.76.65c-.03 0-.08 0-.15-.01c-.07-.01-.11-.01-.13-.01a.768.768 0 0 1-.45-.33a.965.965 0 0 1-.15-.5m.74-2.94c0-.24.08-.43.23-.59c.16-.16.35-.23.59-.23s.43.08.59.23c.16.16.23.35.23.59a.784.784 0 0 1-.82.81c-.23 0-.43-.08-.58-.25a.741.741 0 0 1-.24-.56m1.03-12.11c.84.76 1.4 1.74 1.7 2.93h.31c1.38 0 2.55.48 3.52 1.45c.31-.55.47-1.16.47-1.82c0-.98-.35-1.81-1.04-2.5a3.458 3.458 0 0 0-2.51-1.03c-.96 0-1.78.32-2.45.97m1.6-5.27V1.73c0-.24.08-.44.25-.61c.17-.17.37-.26.6-.26c.24 0 .44.08.6.25c.16.17.24.38.24.62v2.05c0 .24-.08.45-.24.62c-.16.17-.36.25-.6.25c-.23 0-.43-.09-.6-.26a.822.822 0 0 1-.25-.61m5.58 2.29c0-.24.08-.44.23-.6l1.44-1.45a.73.73 0 0 1 .58-.25c.24 0 .44.08.6.25c.18.16.26.36.26.6s-.09.44-.26.6L23.9 6.68a.81.81 0 0 1-.63.26a.778.778 0 0 1-.55-.26c-.16-.16-.23-.36-.23-.61m.77 11.91c0-.24.08-.44.25-.61c.17-.17.37-.25.6-.25c.23 0 .43.09.61.26l.62.63c.18.17.26.38.26.61c0 .24-.09.44-.26.6c-.14.17-.32.26-.54.26l-.02-.02c-.24 0-.44-.08-.62-.24l-.65-.64a.857.857 0 0 1-.25-.6m1.47-6.37c0-.24.08-.44.25-.6c.17-.16.37-.24.61-.24h2.06a.837.837 0 0 1 .86.84c0 .24-.08.44-.25.61c-.17.17-.37.25-.61.25h-2.06a.77.77 0 0 1-.6-.27a.727.727 0 0 1-.26-.59"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DaySleetStorm(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.49 16.88c0 1.12.33 2.12 1 3s1.53 1.47 2.58 1.76l-.66 1.7c-.05.14 0 .22.14.22h2.13l-1.43 4.21h.29l4.36-5.66c.04-.04.04-.09.02-.14c-.02-.05-.07-.07-.14-.07H7.59l2.49-4.65c.07-.14.03-.22-.14-.22H6.98c-.09 0-.17.05-.23.15l-1.07 2.88A3.14 3.14 0 0 1 3.9 18.9c-.47-.59-.7-1.26-.7-2.02c0-.84.28-1.57.84-2.18c.56-.61 1.27-.97 2.11-1.07l.51-.03c.12 0 .19-.05.22-.14l.07-.59c.11-1.08.56-1.99 1.37-2.72s1.76-1.1 2.86-1.1c1.09 0 2.04.37 2.86 1.1s1.29 1.64 1.4 2.72l.08.59c0 .11.06.17.18.17h1.61c.89 0 1.66.32 2.31.96s.97 1.4.97 2.29c0 .87-.3 1.62-.9 2.26s-1.32.98-2.18 1.03c-.13 0-.2.06-.2.18v1.34c0 .11.07.17.2.17c.88-.02 1.69-.26 2.42-.72c.73-.45 1.31-1.06 1.74-1.81s.64-1.57.64-2.45c0-.73-.14-1.4-.43-2.02c.76-.96 1.14-2.06 1.14-3.29c0-.95-.24-1.83-.71-2.64a5.201 5.201 0 0 0-1.92-1.92c-.81-.47-1.69-.71-2.64-.71c-.72 0-1.42.15-2.1.45c-.68.3-1.26.72-1.76 1.25c-.81-.43-1.71-.65-2.72-.65c-1.42 0-2.68.44-3.77 1.32s-1.8 2-2.1 3.37c-1.11.26-2.02.84-2.74 1.74c-.71.92-1.07 1.95-1.07 3.1M9.37 27.1c0 .17.05.33.16.49c.11.16.27.27.49.33c.09.02.17.03.24.03c.43 0 .7-.2.8-.61l.13-.59a.92.92 0 0 0-.08-.68a.762.762 0 0 0-.53-.37a.744.744 0 0 0-.63.07c-.21.12-.34.29-.41.51l-.13.59c-.03.12-.04.2-.04.23M9.9 4.59c0 .23.08.43.25.6l.65.66c.16.16.34.24.55.26c.21.03.41-.04.61-.23c.2-.18.3-.39.3-.64c0-.23-.08-.43-.25-.6l-.63-.66a.814.814 0 0 0-.6-.24c-.25 0-.46.08-.63.24c-.16.18-.25.38-.25.61m.26 19.61c0 .23.08.42.24.58c.16.16.36.24.58.24c.24 0 .43-.08.59-.23c.16-.16.23-.35.23-.59a.784.784 0 0 0-.82-.81c-.24 0-.43.08-.59.23a.8.8 0 0 0-.23.58m.62-2.27c-.01.15.03.31.14.47c.1.16.25.26.45.3c.23.06.44.04.64-.06s.33-.29.41-.56l.26-.9c.07-.22.05-.43-.07-.63a.867.867 0 0 0-.53-.4c-.22-.07-.43-.04-.64.08s-.34.3-.41.53l-.23.9c-.01.08-.02.17-.02.27m2.76 2.15c0 .17.05.33.15.48c.1.15.25.26.46.32c.03 0 .08.01.14.02c.06.01.11.02.14.02c.41 0 .66-.22.76-.66l.14-.6c.07-.21.05-.42-.07-.63a.863.863 0 0 0-.51-.41c-.25-.06-.48-.04-.68.08c-.2.12-.34.29-.41.53l-.09.59c0 .01 0 .05-.01.11c-.01.07-.02.11-.02.15m.74-2.96c0 .23.08.42.24.57c.15.16.34.24.58.24c.24 0 .43-.08.59-.23c.16-.16.23-.35.23-.58c0-.24-.08-.43-.23-.59a.784.784 0 0 0-.59-.23c-.24 0-.43.08-.59.23c-.15.16-.23.35-.23.59m1.03-12.1c.67-.64 1.48-.97 2.45-.97c.98 0 1.82.34 2.51 1.03c.69.68 1.04 1.52 1.04 2.5c0 .66-.16 1.26-.47 1.81c-.96-.96-2.13-1.44-3.52-1.44h-.31c-.29-1.19-.86-2.17-1.7-2.93m1.6-5.27c0 .24.08.44.25.61s.37.25.6.25c.24 0 .44-.08.6-.25c.16-.17.24-.37.24-.61V1.69c0-.24-.08-.45-.24-.61a.794.794 0 0 0-.6-.26c-.24 0-.44.08-.6.25s-.25.37-.25.61zm5.58 2.29c0 .24.08.44.23.6c.14.16.32.24.55.26c.23.02.44-.07.63-.26l1.44-1.44c.18-.16.26-.36.26-.6s-.09-.44-.26-.6a.76.76 0 0 0-.6-.26c-.23 0-.42.09-.58.26l-1.44 1.44c-.16.15-.23.35-.23.6m.77 11.91c0 .23.08.43.25.6l.65.63c.18.17.39.25.62.25l.02.02c.22 0 .4-.09.54-.27a.76.76 0 0 0 .26-.6c0-.23-.09-.43-.26-.61l-.62-.62a.836.836 0 0 0-.61-.27c-.24 0-.44.09-.6.26a.88.88 0 0 0-.25.61m1.47-6.37c0 .24.09.44.26.59c.16.18.36.26.6.26h2.06c.24 0 .44-.08.61-.25c.17-.17.25-.37.25-.6s-.08-.44-.25-.6a.853.853 0 0 0-.61-.24h-2.06c-.24 0-.45.08-.61.24c-.17.16-.25.36-.25.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DaySnow(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.58 16.93c0 .86.21 1.67.64 2.41c.42.74 1 1.34 1.74 1.79c.73.45 1.54.69 2.4.71c.11 0 .17-.06.17-.17v-1.33c0-.12-.06-.19-.17-.19c-.85-.04-1.58-.38-2.18-1.02s-.9-1.37-.9-2.21c0-.82.28-1.54.85-2.16c.57-.61 1.26-.97 2.1-1.07l.53-.06c.12 0 .18-.06.18-.19l.08-.51c.11-1.09.56-2 1.36-2.73s1.75-1.09 2.85-1.09c1.09 0 2.04.36 2.85 1.09c.82.73 1.28 1.63 1.38 2.7l.07.58c0 .11.06.17.17.17h1.61c.9 0 1.67.32 2.31.96c.64.64.96 1.4.96 2.29c0 .84-.3 1.57-.9 2.21c-.6.63-1.33.97-2.17 1.02c-.12 0-.19.06-.19.19v1.33c0 .11.06.17.19.17c1.33-.04 2.45-.54 3.38-1.5c.93-.96 1.39-2.09 1.39-3.41c0-.76-.14-1.43-.43-2.03c.75-.93 1.15-2.03 1.15-3.28c0-.94-.23-1.81-.7-2.61c-.47-.8-1.11-1.44-1.91-1.91s-1.68-.7-2.62-.7c-1.54 0-2.83.58-3.87 1.73c-.81-.44-1.71-.66-2.69-.66c-1.41 0-2.65.44-3.74 1.31s-1.78 1.99-2.09 3.34c-1.12.28-2.03.86-2.74 1.75c-.71.9-1.06 1.92-1.06 3.08m6.34 4.05c0 .24.08.44.24.61c.16.17.35.25.59.25c.23 0 .43-.08.59-.25c.16-.17.24-.37.24-.61c0-.23-.08-.42-.24-.58s-.35-.24-.59-.24c-.23 0-.43.08-.59.24s-.24.36-.24.58m0 3.63c0 .21.08.4.24.57c.18.16.37.24.58.24c.24 0 .43-.08.59-.23c.16-.16.23-.35.23-.58c0-.24-.08-.43-.24-.59a.784.784 0 0 0-.59-.23a.84.84 0 0 0-.59.23c-.14.15-.22.35-.22.59M9.97 4.68c0 .24.08.44.24.59l.66.66c.16.16.34.25.53.25c.21.03.41-.04.61-.22s.3-.39.3-.63s-.08-.46-.24-.64l-.64-.61a.73.73 0 0 0-.58-.25c-.25 0-.46.08-.63.25c-.17.16-.25.36-.25.6M11.1 22.9c0 .22.08.42.24.6c.16.16.36.24.58.24c.24 0 .44-.08.6-.24s.25-.36.25-.6c0-.23-.08-.43-.25-.6s-.37-.25-.6-.25c-.23 0-.42.08-.58.25s-.24.37-.24.6m0-3.6c0 .23.08.42.24.58s.36.24.58.24c.24 0 .44-.08.6-.24c.17-.16.25-.35.25-.59c0-.23-.08-.43-.25-.59s-.37-.24-.6-.24c-.23 0-.42.08-.58.24s-.24.37-.24.6m0 7.26c0 .22.08.41.24.57c.17.17.36.25.58.25c.24 0 .44-.08.6-.23c.17-.16.25-.35.25-.59s-.08-.44-.25-.6a.816.816 0 0 0-.6-.25c-.22 0-.41.08-.58.25c-.16.17-.24.37-.24.6m3.22-5.58c0 .24.08.44.24.61c.16.17.36.25.59.25s.43-.08.59-.25c.16-.17.24-.37.24-.61c0-.23-.08-.42-.24-.58s-.35-.24-.59-.24s-.43.08-.59.24s-.24.36-.24.58m0 3.63c0 .21.08.4.23.57c.18.16.38.24.6.24c.24 0 .43-.08.59-.23c.16-.16.23-.35.23-.58c0-.24-.08-.43-.24-.59a.784.784 0 0 0-.59-.23c-.24 0-.44.08-.6.24c-.14.15-.22.35-.22.58m.98-15.55c.69-.66 1.51-.99 2.47-.99c.97 0 1.8.35 2.48 1.04c.69.69 1.03 1.53 1.03 2.49c0 .62-.17 1.24-.51 1.84c-.95-.96-2.11-1.44-3.47-1.44h-.32c-.3-1.17-.86-2.15-1.68-2.94m1.6-5.22c0 .23.08.43.25.58s.37.23.61.23s.43-.08.59-.23c.16-.16.23-.35.23-.58V1.8c0-.24-.08-.44-.24-.61S18 .94 17.77.94s-.43.09-.6.26c-.17.17-.26.37-.26.6v2.04zm5.52 2.27c0 .23.08.43.25.59c.15.16.34.24.56.26s.43-.07.62-.26l1.43-1.43c.18-.18.26-.38.26-.61c0-.24-.09-.44-.26-.61a.816.816 0 0 0-.6-.25c-.22 0-.41.08-.58.25l-1.43 1.46c-.17.16-.25.36-.25.6m.8 11.8c0 .25.08.46.24.62l.64.63c.24.16.46.24.64.24c.21 0 .39-.09.56-.26c.17-.17.25-.38.25-.61c0-.23-.09-.42-.26-.58l-.62-.65c-.18-.16-.38-.24-.61-.24s-.43.08-.59.25c-.17.16-.25.36-.25.6m1.45-6.31c0 .24.09.43.26.59c.17.18.38.27.62.27h2.02c.23 0 .43-.08.6-.25s.25-.37.25-.61s-.08-.44-.25-.6s-.37-.25-.6-.25h-2.02c-.24 0-.44.08-.62.25s-.26.37-.26.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DaySnowThunderstorm(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.49 16.88c0 1.12.33 2.12 1 3s1.53 1.47 2.58 1.76l-.66 1.7c-.05.14 0 .22.14.22h2.13l-1.43 4.21h.29l4.36-5.66c.04-.04.04-.09.02-.14c-.02-.05-.07-.07-.14-.07H7.59l2.49-4.65c.07-.14.03-.22-.14-.22H6.98c-.09 0-.17.05-.23.15l-1.07 2.88A3.14 3.14 0 0 1 3.9 18.9c-.47-.59-.7-1.26-.7-2.02c0-.84.28-1.57.84-2.18c.56-.61 1.27-.97 2.11-1.07l.51-.03c.12 0 .19-.05.22-.14l.07-.59c.11-1.08.56-1.99 1.37-2.72s1.76-1.1 2.86-1.1c1.09 0 2.04.37 2.86 1.1s1.29 1.64 1.4 2.72l.08.59c0 .11.06.17.18.17h1.61c.89 0 1.66.32 2.31.96s.97 1.4.97 2.29c0 .87-.3 1.62-.9 2.26s-1.32.98-2.18 1.03c-.13 0-.2.06-.2.18v1.34c0 .11.07.17.2.17c.88-.02 1.69-.26 2.42-.72c.73-.45 1.31-1.06 1.74-1.81s.64-1.57.64-2.45c0-.73-.14-1.4-.43-2.02c.76-.96 1.14-2.06 1.14-3.29c0-.95-.24-1.83-.71-2.64a5.201 5.201 0 0 0-1.92-1.92c-.81-.47-1.69-.71-2.64-.71c-.72 0-1.42.15-2.1.45c-.68.3-1.26.72-1.76 1.25c-.81-.43-1.71-.65-2.72-.65c-1.42 0-2.68.44-3.77 1.32s-1.8 2-2.1 3.37c-1.11.26-2.02.84-2.74 1.74c-.71.92-1.07 1.95-1.07 3.1M9.9 4.59c0 .23.08.43.25.6l.65.66c.16.16.34.24.55.26c.21.03.41-.04.61-.23c.2-.18.3-.39.3-.64c0-.23-.08-.43-.25-.6l-.63-.66a.814.814 0 0 0-.6-.24c-.25 0-.46.08-.63.24c-.16.18-.25.38-.25.61m1.18 18.37c0 .24.08.44.24.59c.16.16.36.24.58.24c.24 0 .44-.08.61-.24s.25-.36.25-.59c0-.24-.08-.44-.25-.61s-.37-.26-.61-.26c-.22 0-.41.09-.58.26s-.24.37-.24.61m0-3.64c0 .24.08.43.24.58c.16.16.36.24.58.24c.24 0 .45-.08.61-.23s.25-.35.25-.59c0-.23-.08-.43-.25-.6s-.37-.25-.61-.25c-.23 0-.42.08-.58.25s-.24.37-.24.6m0 7.31c0 .22.08.41.24.57c.17.17.36.25.58.25c.24 0 .44-.08.61-.24c.17-.16.25-.35.25-.59s-.08-.44-.25-.61a.832.832 0 0 0-.61-.26c-.22 0-.41.09-.58.26c-.16.18-.24.39-.24.62m3.23-5.61c0 .24.08.44.25.6s.36.25.6.25c.23 0 .43-.08.59-.25c.16-.17.24-.37.24-.6c0-.22-.08-.42-.24-.58a.791.791 0 0 0-.59-.24c-.23 0-.43.08-.6.24s-.25.35-.25.58m0 3.64c0 .23.08.42.24.58c.16.16.36.24.6.24s.43-.08.59-.24c.16-.16.23-.35.23-.59s-.08-.43-.23-.59a.784.784 0 0 0-.59-.23c-.24 0-.44.08-.6.23c-.16.16-.24.36-.24.6m1-15.64c.67-.64 1.48-.97 2.45-.97c.98 0 1.82.34 2.51 1.03c.69.68 1.04 1.52 1.04 2.5c0 .66-.16 1.26-.47 1.81c-.96-.96-2.13-1.44-3.52-1.44h-.31c-.29-1.19-.86-2.17-1.7-2.93m1.6-5.27c0 .24.08.44.25.61s.37.25.6.25c.24 0 .44-.08.6-.25c.16-.17.24-.37.24-.61V1.69c0-.24-.08-.45-.24-.61a.794.794 0 0 0-.6-.26c-.24 0-.44.08-.6.25s-.25.37-.25.61zm5.58 2.29c0 .24.08.44.23.6c.14.16.32.24.55.26c.23.02.44-.07.63-.26l1.44-1.44c.18-.16.26-.36.26-.6s-.09-.44-.26-.6a.76.76 0 0 0-.6-.26c-.23 0-.42.09-.58.26l-1.44 1.44c-.16.15-.23.35-.23.6m.77 11.91c0 .23.08.43.25.6l.65.63c.18.17.39.25.62.25l.02.02c.22 0 .4-.09.54-.27a.76.76 0 0 0 .26-.6c0-.23-.09-.43-.26-.61l-.62-.62a.836.836 0 0 0-.61-.27c-.24 0-.44.09-.6.26a.88.88 0 0 0-.25.61m1.47-6.37c0 .24.09.44.26.59c.16.18.36.26.6.26h2.06c.24 0 .44-.08.61-.25c.17-.17.25-.37.25-.6s-.08-.44-.25-.6a.853.853 0 0 0-.61-.24h-2.06c-.24 0-.45.08-.61.24c-.17.16-.25.36-.25.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DaySnowWind(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.58 16.93c0 .86.21 1.67.64 2.41c.42.74 1 1.34 1.74 1.79c.73.45 1.54.69 2.4.71c.11 0 .17-.06.17-.17v-1.33c0-.12-.06-.19-.17-.19c-.85-.04-1.58-.38-2.18-1.02s-.9-1.37-.9-2.21c0-.82.28-1.54.85-2.16c.57-.61 1.26-.97 2.1-1.07l.53-.06c.12 0 .18-.06.18-.19l.08-.51c.11-1.09.56-2 1.36-2.73s1.75-1.09 2.85-1.09c1.09 0 2.04.36 2.85 1.09c.82.73 1.28 1.63 1.38 2.7l.07.58c0 .11.06.17.17.17h1.61c.9 0 1.67.32 2.31.96c.64.64.96 1.4.96 2.29c0 .84-.3 1.57-.9 2.21c-.6.63-1.33.97-2.17 1.02c-.12 0-.19.06-.19.19v1.33c0 .11.06.17.19.17c1.33-.04 2.45-.54 3.38-1.5c.93-.96 1.39-2.09 1.39-3.41c0-.76-.14-1.43-.43-2.03c.75-.93 1.15-2.03 1.15-3.28c0-.94-.23-1.81-.7-2.61c-.47-.8-1.11-1.44-1.91-1.91s-1.68-.7-2.62-.7c-1.54 0-2.83.58-3.87 1.73c-.81-.44-1.71-.66-2.69-.66c-1.41 0-2.65.44-3.74 1.31s-1.78 1.99-2.09 3.34c-1.12.28-2.03.86-2.74 1.75c-.71.9-1.06 1.92-1.06 3.08m5.48 7.68c0 .21.08.4.24.57c.18.16.37.24.58.24c.24 0 .43-.08.59-.23c.16-.16.23-.35.23-.58c0-.24-.08-.43-.24-.59c-.16-.16-.35-.23-.59-.23s-.43.08-.59.23c-.14.15-.22.35-.22.59m.86-3.63c0 .24.08.44.24.61c.16.17.35.25.59.25c.23 0 .43-.08.59-.25c.16-.17.24-.37.24-.61c0-.23-.08-.42-.24-.58s-.35-.24-.59-.24c-.23 0-.43.08-.59.24s-.24.36-.24.58m1.9 5.58c0 .22.08.41.24.57c.17.17.36.25.58.25c.24 0 .44-.08.6-.23c.17-.16.25-.35.25-.59s-.08-.44-.25-.6a.816.816 0 0 0-.6-.25c-.22 0-.41.08-.58.25c-.16.17-.24.37-.24.6m.15-21.88c0 .24.08.44.24.59l.66.66c.16.16.34.25.53.25c.21.03.41-.04.61-.22s.3-.39.3-.63s-.08-.46-.24-.64l-.64-.61a.73.73 0 0 0-.58-.25c-.25 0-.46.08-.63.25c-.17.16-.25.36-.25.6m.7 18.22c0 .22.08.42.25.6c.16.16.35.24.57.24c.24 0 .44-.08.61-.24c.17-.16.25-.36.25-.6c0-.23-.08-.43-.25-.6a.822.822 0 0 0-.61-.25c-.22 0-.42.08-.58.25c-.16.17-.24.37-.24.6m.43-3.6c0 .23.08.42.24.58s.36.24.58.24c.24 0 .44-.08.6-.24c.17-.16.25-.35.25-.59c0-.23-.08-.43-.25-.59s-.37-.24-.6-.24c-.23 0-.42.08-.58.24s-.24.37-.24.6m2.37 5.31c0 .21.08.4.23.57c.17.16.38.24.6.24c.24 0 .43-.08.59-.23c.16-.16.23-.35.23-.58c0-.24-.08-.43-.23-.59a.784.784 0 0 0-.59-.23c-.24 0-.44.08-.6.24s-.23.35-.23.58m.85-3.63c0 .24.08.44.24.61c.16.17.36.25.59.25s.43-.08.59-.25c.16-.17.24-.37.24-.61c0-.23-.08-.42-.24-.58s-.35-.24-.59-.24s-.43.08-.59.24s-.24.36-.24.58m.98-11.92c.69-.66 1.51-.99 2.47-.99c.97 0 1.8.35 2.48 1.04c.69.69 1.03 1.53 1.03 2.49c0 .62-.17 1.24-.51 1.84c-.95-.96-2.11-1.44-3.47-1.44h-.32c-.3-1.17-.86-2.15-1.68-2.94m1.6-5.22c0 .23.08.43.25.58s.37.23.61.23s.43-.08.59-.23c.16-.16.23-.35.23-.58V1.8c0-.24-.08-.44-.24-.61S18 .94 17.77.94s-.43.09-.6.26c-.17.17-.26.37-.26.6v2.04zm5.52 2.27c0 .23.08.43.25.59c.15.16.34.24.56.26s.43-.07.62-.26l1.43-1.43c.18-.18.26-.38.26-.61c0-.24-.09-.44-.26-.61a.816.816 0 0 0-.6-.25c-.22 0-.41.08-.58.25l-1.43 1.46c-.17.16-.25.36-.25.6m.8 11.8c0 .25.08.46.24.62l.64.63c.24.16.46.24.64.24c.21 0 .39-.09.56-.26c.17-.17.25-.38.25-.61c0-.23-.09-.42-.26-.58l-.62-.65c-.18-.16-.38-.24-.61-.24s-.43.08-.59.25c-.17.16-.25.36-.25.6m1.45-6.31c0 .24.09.43.26.59c.17.18.38.27.62.27h2.02c.23 0 .43-.08.6-.25s.25-.37.25-.61s-.08-.44-.25-.6s-.37-.25-.6-.25h-2.02c-.24 0-.44.08-.62.25s-.26.37-.26.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DaySprinkle(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.58 16.89c0 .87.21 1.68.64 2.42c.42.75 1 1.35 1.73 1.8c.73.45 1.53.69 2.4.73c.12 0 .18-.06.18-.17v-1.33c0-.12-.06-.19-.18-.19c-.85-.04-1.58-.38-2.18-1.02s-.9-1.38-.9-2.25c0-.82.28-1.54.84-2.15s1.26-.96 2.09-1.06l.52-.03c.12 0 .19-.06.19-.18l.08-.54c.11-1.08.56-1.98 1.36-2.71s1.75-1.09 2.85-1.09c1.07 0 2.02.36 2.84 1.09a4.11 4.11 0 0 1 1.4 2.7l.07.58c0 .11.06.17.17.17h1.62c.88 0 1.65.32 2.29.96c.65.64.97 1.39.97 2.26c0 .86-.3 1.61-.9 2.25c-.6.63-1.33.97-2.18 1.02c-.12 0-.18.06-.18.19v1.33c0 .11.06.17.18.17c.87-.02 1.67-.26 2.4-.72s1.31-1.05 1.73-1.8s.63-1.56.63-2.43c0-.74-.14-1.4-.41-2c.78-.95 1.16-2.05 1.16-3.3c0-.94-.23-1.82-.7-2.62s-1.1-1.44-1.9-1.9c-.8-.47-1.67-.7-2.61-.7c-1.55 0-2.84.58-3.87 1.73c-.82-.44-1.72-.66-2.69-.66c-1.41 0-2.65.44-3.74 1.31s-1.78 1.99-2.09 3.34c-1.1.26-2.01.84-2.72 1.73c-.74.91-1.09 1.93-1.09 3.07m5.9.88c0 .38.14.71.42.98s.62.4 1.01.4c.4 0 .73-.13 1-.4s.4-.59.4-.98c0-.24-.12-.58-.35-1c-.23-.43-.45-.76-.65-.99c-.21-.22-.35-.37-.4-.42l-.36.4c-.27.29-.52.63-.74 1.03s-.33.73-.33.98M9.97 4.66c0 .26.08.46.24.61l.65.66c.42.31.82.31 1.21 0a.98.98 0 0 0 .24-.64c0-.23-.08-.43-.24-.59l-.64-.65a.89.89 0 0 0-.61-.25c-.24 0-.45.08-.61.25c-.16.17-.24.37-.24.61m.48 17.07c0 .66.23 1.21.68 1.65s1 .67 1.65.67c.65 0 1.2-.23 1.66-.68c.46-.46.68-1 .68-1.64c0-.54-.27-1.19-.81-1.97c-.46-.61-.89-1.1-1.28-1.49c-.08-.06-.17-.13-.26-.23l-.23.23c-.36.32-.78.82-1.27 1.47c-.54.76-.82 1.42-.82 1.99m1.48-6.62c0 .25.1.47.29.65c.19.18.42.27.69.27c.26 0 .48-.09.66-.27c.18-.18.27-.4.27-.65c0-.41-.31-.95-.93-1.6l-.24.25c-.18.2-.35.43-.5.7c-.17.27-.24.48-.24.65m3.37-6.05c.66-.66 1.48-.99 2.47-.99c.98 0 1.8.34 2.49 1.03c.68.69 1.03 1.52 1.03 2.5c0 .59-.17 1.2-.52 1.84c-.97-.96-2.13-1.44-3.47-1.44h-.33a6.086 6.086 0 0 0-1.67-2.94m1.6-5.26c0 .24.09.44.26.61c.17.17.37.25.6.25s.43-.08.58-.25c.16-.17.23-.37.23-.61V1.76c0-.24-.08-.43-.23-.59S18 .94 17.77.94s-.44.08-.61.23s-.26.35-.26.59zm5.52 2.31c0 .24.08.44.25.59c.17.16.36.24.56.24c.17 0 .38-.08.61-.24l1.43-1.43c.18-.18.27-.39.27-.62c0-.24-.08-.45-.25-.61a.853.853 0 0 0-.61-.24c-.22 0-.41.08-.58.25l-1.43 1.43c-.17.17-.25.38-.25.63m.8 11.78c0 .25.08.46.23.64l.65.58c.14.18.34.27.59.27c.24 0 .44-.09.58-.27c.18-.16.27-.36.27-.58c0-.22-.09-.41-.27-.58l-.61-.65a.888.888 0 0 0-.61-.24c-.24 0-.44.08-.6.24s-.23.36-.23.59m1.45-6.29c0 .22.08.42.25.58c.17.16.38.24.63.24h2.02c.24 0 .44-.08.6-.24c.17-.16.25-.35.25-.59s-.08-.44-.25-.6s-.37-.25-.6-.25h-2.02c-.24 0-.44.08-.62.25s-.26.38-.26.61"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DayStormShowers(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.49 16.88c0 1.12.33 2.12 1 3s1.53 1.47 2.58 1.76l-.66 1.7c-.05.14 0 .22.14.22h2.13l-1.43 4.21h.29l4.36-5.66c.04-.04.04-.09.02-.14c-.02-.05-.07-.07-.14-.07H7.59l2.49-4.65c.07-.14.03-.22-.14-.22H6.98c-.09 0-.17.05-.23.15l-1.07 2.88A3.14 3.14 0 0 1 3.9 18.9c-.47-.59-.7-1.26-.7-2.02c0-.84.28-1.57.84-2.18c.56-.61 1.27-.97 2.11-1.07l.51-.03c.12 0 .19-.05.22-.14l.07-.59c.11-1.08.56-1.99 1.37-2.72s1.76-1.1 2.86-1.1c1.09 0 2.04.37 2.86 1.1s1.29 1.64 1.4 2.72l.08.59c0 .11.06.17.18.17h1.61c.89 0 1.66.32 2.31.96s.97 1.4.97 2.29c0 .87-.3 1.62-.9 2.26s-1.32.98-2.18 1.03c-.13 0-.2.06-.2.18v1.34c0 .11.07.17.2.17c.88-.02 1.69-.26 2.42-.72c.73-.45 1.31-1.06 1.74-1.81s.64-1.57.64-2.45c0-.73-.14-1.4-.43-2.02c.76-.96 1.14-2.06 1.14-3.29c0-.95-.24-1.83-.71-2.64a5.201 5.201 0 0 0-1.92-1.92c-.81-.47-1.69-.71-2.64-.71c-.72 0-1.42.15-2.1.45c-.68.3-1.26.72-1.76 1.25c-.81-.43-1.71-.65-2.72-.65c-1.42 0-2.68.44-3.77 1.32s-1.8 2-2.1 3.37c-1.11.26-2.02.84-2.74 1.74c-.71.92-1.07 1.95-1.07 3.1m8.18 9.92c0 .15.05.31.16.47a.89.89 0 0 0 .71.38c.15 0 .28-.03.38-.08c.21-.08.36-.27.43-.57l.27-1.03c.06-.25.03-.47-.08-.67s-.3-.32-.53-.37c-.21-.07-.41-.04-.62.07a.87.87 0 0 0-.42.52l-.25 1.04c-.03.14-.05.22-.05.24M9.9 4.59c0 .23.08.43.25.6l.65.66c.16.16.34.24.55.26c.21.03.41-.04.61-.23c.2-.18.3-.39.3-.64c0-.23-.08-.43-.25-.6l-.63-.66a.814.814 0 0 0-.6-.24c-.25 0-.46.08-.63.24c-.16.18-.25.38-.25.61M11.01 22c-.01.16.04.32.14.47c.1.15.26.26.48.32c.21.07.42.05.62-.06s.34-.3.42-.56l.3-1.03c.07-.22.04-.43-.08-.63s-.3-.34-.54-.41a.755.755 0 0 0-.64.07c-.2.12-.34.29-.41.53l-.24 1.05c-.03.15-.05.23-.05.25m2.83 1.68c0 .14.03.28.1.39c.13.21.31.36.54.43c.11.04.21.06.28.06c.13 0 .23-.02.31-.08c.2-.07.35-.27.45-.6l.25-1.01c.07-.24.05-.45-.07-.65a.78.78 0 0 0-.51-.39a.777.777 0 0 0-.65.07c-.2.11-.34.28-.41.51l-.28 1.04c0 .08-.01.16-.01.23m1.37-4.82a.85.85 0 0 0 .65.83c.17.06.37.04.61-.05c.2-.09.34-.28.43-.57l.27-1c.06-.25.04-.47-.08-.67s-.29-.32-.53-.37a.79.79 0 0 0-.64.06c-.2.11-.33.28-.4.5l-.29 1.06c-.01.14-.02.21-.02.21m.1-9.84c.67-.64 1.48-.97 2.45-.97c.98 0 1.82.34 2.51 1.03c.69.68 1.04 1.52 1.04 2.5c0 .66-.16 1.26-.47 1.81c-.96-.96-2.13-1.44-3.52-1.44h-.31c-.29-1.19-.86-2.17-1.7-2.93m1.6-5.27c0 .24.08.44.25.61s.37.25.6.25c.24 0 .44-.08.6-.25c.16-.17.24-.37.24-.61V1.69c0-.24-.08-.45-.24-.61a.794.794 0 0 0-.6-.26c-.24 0-.44.08-.6.25s-.25.37-.25.61zm5.58 2.29c0 .24.08.44.23.6c.14.16.32.24.55.26c.23.02.44-.07.63-.26l1.44-1.44c.18-.16.26-.36.26-.6s-.09-.44-.26-.6a.76.76 0 0 0-.6-.26c-.23 0-.42.09-.58.26l-1.44 1.44c-.16.15-.23.35-.23.6m.77 11.91c0 .23.08.43.25.6l.65.63c.18.17.39.25.62.25l.02.02c.22 0 .4-.09.54-.27a.76.76 0 0 0 .26-.6c0-.23-.09-.43-.26-.61l-.62-.62a.836.836 0 0 0-.61-.27c-.24 0-.44.09-.6.26a.88.88 0 0 0-.25.61m1.47-6.37c0 .24.09.44.26.59c.16.18.36.26.6.26h2.06c.24 0 .44-.08.61-.25c.17-.17.25-.37.25-.6s-.08-.44-.25-.6a.853.853 0 0 0-.61-.24h-2.06c-.24 0-.45.08-.61.24c-.17.16-.25.36-.25.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DaySunny(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.37 14.62c0-.24.08-.45.25-.62c.17-.16.38-.24.6-.24h2.04c.23 0 .42.08.58.25c.15.17.23.37.23.61s-.07.44-.22.61c-.15.17-.35.25-.58.25H5.23c-.23 0-.43-.08-.6-.25a.832.832 0 0 1-.26-.61m2.86 6.93c0-.23.08-.43.23-.61l1.47-1.43c.15-.16.35-.23.59-.23s.44.08.6.23s.24.34.24.57c0 .24-.08.46-.24.64L8.7 22.14c-.41.32-.82.32-1.23 0a.807.807 0 0 1-.24-.59m0-13.84c0-.23.08-.43.23-.61c.2-.17.41-.25.64-.25c.22 0 .42.08.59.24l1.43 1.47c.16.15.24.35.24.59s-.08.44-.24.6s-.36.24-.6.24s-.44-.08-.59-.24L7.47 8.32a.837.837 0 0 1-.24-.61m2.55 6.91c0-.93.23-1.8.7-2.6s1.1-1.44 1.91-1.91s1.67-.7 2.6-.7c.7 0 1.37.14 2.02.42c.64.28 1.2.65 1.66 1.12c.47.47.84 1.02 1.11 1.66c.27.64.41 1.32.41 2.02c0 .94-.23 1.81-.7 2.61c-.47.8-1.1 1.43-1.9 1.9c-.8.47-1.67.7-2.61.7s-1.81-.23-2.61-.7c-.8-.47-1.43-1.1-1.9-1.9c-.45-.81-.69-1.68-.69-2.62m1.7 0c0 .98.34 1.81 1.03 2.5c.68.69 1.51 1.04 2.49 1.04s1.81-.35 2.5-1.04s1.04-1.52 1.04-2.5c0-.96-.35-1.78-1.04-2.47c-.69-.68-1.52-1.02-2.5-1.02c-.97 0-1.8.34-2.48 1.02c-.7.69-1.04 1.51-1.04 2.47m2.66 7.78c0-.24.08-.44.25-.6s.37-.24.6-.24c.24 0 .45.08.61.24s.24.36.24.6v1.99c0 .24-.08.45-.25.62c-.17.17-.37.25-.6.25s-.44-.08-.6-.25a.845.845 0 0 1-.25-.62zm0-15.5V4.86c0-.23.08-.43.25-.6c.17-.17.37-.26.61-.26s.43.08.6.25c.17.17.25.37.25.6V6.9c0 .23-.08.42-.25.58s-.37.23-.6.23s-.44-.08-.6-.23s-.26-.35-.26-.58m5.52 13.18c0-.23.08-.42.23-.56c.15-.16.34-.23.56-.23c.24 0 .44.08.6.23l1.46 1.43c.16.17.24.38.24.61c0 .23-.08.43-.24.59c-.4.31-.8.31-1.2 0l-1.42-1.42a.974.974 0 0 1-.23-.65m0-10.92c0-.25.08-.45.23-.59l1.42-1.47a.84.84 0 0 1 .59-.24c.24 0 .44.08.6.25c.17.17.25.37.25.6c0 .25-.08.46-.24.62l-1.46 1.43c-.18.16-.38.24-.6.24c-.23 0-.41-.08-.56-.24s-.23-.36-.23-.6m2.26 5.46c0-.24.08-.44.24-.62c.16-.16.35-.24.57-.24h2.02c.23 0 .43.09.6.26c.17.17.26.37.26.6s-.09.43-.26.6c-.17.17-.37.25-.6.25h-2.02c-.23 0-.43-.08-.58-.25s-.23-.36-.23-.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DaySunnyOvercast(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.09 13.43c0 .27.09.49.28.67s.43.26.72.26h1.69c.27 0 .5-.09.69-.27s.29-.4.29-.66c0-.29-.09-.53-.28-.71s-.42-.28-.69-.28h-1.7c-.29 0-.53.09-.72.28c-.19.17-.28.41-.28.71m1.77 6.18c0 .97.35 1.81 1.06 2.52c.71.71 1.54 1.06 2.51 1.06h6.86c.97 0 1.8-.35 2.49-1.05c.69-.7 1.04-1.54 1.04-2.53c0-.48-.07-.89-.21-1.23c.83-.53 1.49-1.24 1.97-2.12c.48-.88.73-1.83.73-2.84c0-.81-.16-1.59-.48-2.33a6.1 6.1 0 0 0-1.28-1.91a6.1 6.1 0 0 0-1.91-1.28c-.74-.32-1.51-.48-2.32-.48c-1.09 0-2.1.27-3.02.81s-1.65 1.27-2.18 2.18c-.53.92-.79 1.92-.79 3.01v.34c-1 .57-1.66 1.4-1.98 2.49c-.76.24-1.36.66-1.81 1.27a3.44 3.44 0 0 0-.68 2.09M6.42 5.45c0 .28.09.51.26.67l1.61 1.69c.47.36.94.36 1.41 0c.19-.19.28-.42.28-.7c0-.27-.09-.5-.28-.7L8.05 4.76c-.21-.19-.45-.29-.7-.29c-.28 0-.51.09-.68.28c-.17.19-.25.42-.25.7m.42 14.16c0-.42.13-.78.4-1.08c.27-.3.61-.47 1.02-.51l.62-.08c.13 0 .2-.08.2-.23l.09-.56c.07-.58.31-1.06.73-1.44c.42-.39.91-.58 1.48-.58c.58 0 1.09.19 1.51.58c.43.39.68.87.75 1.44l.08.65c.06.15.14.23.24.23h1.32c.43 0 .8.16 1.12.47c.32.31.47.68.47 1.12c0 .45-.16.83-.47 1.15s-.69.48-1.12.48H8.43c-.45 0-.83-.16-1.13-.48c-.31-.32-.46-.71-.46-1.16m4.42-6.39c.07-1.09.49-2.01 1.27-2.76c.77-.74 1.71-1.12 2.79-1.12c1.11 0 2.06.4 2.84 1.19c.78.79 1.17 1.76 1.17 2.89c0 .7-.17 1.35-.51 1.95c-.34.6-.8 1.08-1.38 1.45c-.59-.49-1.27-.73-2.03-.73c-.29-.88-.81-1.57-1.54-2.09c-.73-.52-1.56-.78-2.48-.78zm3.09-8.75c0 .27.1.51.29.7c.19.19.42.29.69.29c.28 0 .51-.1.7-.29a.95.95 0 0 0 .29-.7V2.13c0-.26-.1-.48-.29-.66a.974.974 0 0 0-.7-.27c-.27 0-.5.09-.69.27c-.19.18-.29.4-.29.66zm6.32 15.23c0 .27.09.5.27.7l1.64 1.62c.42.42.89.42 1.41 0c.18-.17.26-.39.26-.68c0-.27-.09-.49-.26-.67L22.3 19a.953.953 0 0 0-.68-.25c-.28 0-.5.09-.68.27c-.18.18-.27.41-.27.68m0-12.61c0 .28.09.52.27.72c.18.18.41.27.68.27c.27 0 .5-.09.68-.27l1.69-1.69c.18-.17.26-.39.26-.67s-.1-.51-.29-.69s-.42-.28-.7-.28c-.26 0-.49.1-.68.29L20.94 6.4c-.18.18-.27.41-.27.69m2.58 6.34c0 .27.09.49.28.67s.43.26.72.26h1.69c.27 0 .5-.09.69-.27s.29-.4.29-.66c0-.29-.09-.53-.28-.71s-.42-.28-.69-.28h-1.69c-.29 0-.53.09-.72.28c-.19.17-.29.41-.29.71"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DayThunderstorm(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.52 16.9c0 1.11.33 2.09.98 2.96s1.51 1.46 2.57 1.78l-.64 1.7c-.04.14 0 .21.14.21H6.7L5.45 27.5h.29l4.17-5.39c.04-.04.04-.09.01-.14c-.02-.05-.07-.07-.14-.07H7.61l2.47-4.63c.07-.14.02-.22-.14-.22H7c-.09 0-.17.05-.23.14L5.7 20.07c-.71-.18-1.3-.57-1.77-1.16s-.7-1.26-.7-2.01c0-.83.28-1.55.85-2.17s1.27-.97 2.1-1.07l.52-.06c.13 0 .2-.06.2-.18l.06-.51c.11-1.08.57-1.99 1.38-2.72a4.15 4.15 0 0 1 2.86-1.1c1.09 0 2.04.37 2.85 1.1s1.28 1.64 1.4 2.72l.06.58c0 .11.06.17.18.17h1.61c.91 0 1.68.32 2.32.95c.64.63.96 1.39.96 2.29c0 .85-.3 1.59-.89 2.21c-.59.62-1.32.97-2.19 1.04c-.13 0-.2.06-.2.18v1.37c0 .11.07.17.2.17c1.33-.04 2.46-.55 3.39-1.51c.93-.96 1.39-2.11 1.39-3.45c0-.74-.14-1.41-.43-2.01c.79-.96 1.18-2.06 1.18-3.32c0-.94-.24-1.81-.71-2.62a5.201 5.201 0 0 0-1.92-1.92c-.81-.47-1.68-.71-2.62-.71c-1.54 0-2.84.58-3.88 1.73c-.81-.43-1.71-.65-2.7-.65c-1.41 0-2.67.44-3.76 1.31s-1.79 1.99-2.1 3.36c-1.11.26-2.02.83-2.73 1.73s-1.09 1.94-1.09 3.09m8.09 9.58c-.01.15.03.3.14.44s.26.25.46.33c.07.02.14.03.21.03c.17 0 .34-.05.51-.15s.28-.26.34-.47l2.29-8.57a.83.83 0 0 0-.07-.64a.815.815 0 0 0-.49-.4a.777.777 0 0 0-.65.07c-.2.11-.34.28-.4.51l-2.31 8.6c-.02.07-.03.16-.03.25m.33-21.85c0 .24.08.43.25.59l.64.66c.17.17.37.25.61.26c.24 0 .43-.08.57-.26c.19-.15.28-.35.28-.6c0-.24-.08-.43-.25-.59l-.63-.66a.87.87 0 0 0-.61-.24c-.25 0-.46.08-.62.24c-.16.16-.24.36-.24.6m3.83 18.8c0 .12.04.24.11.38c.13.2.29.34.5.43c.07.03.17.05.3.05c.15 0 .26-.02.33-.06c.2-.08.34-.28.41-.58l1.49-5.55c.06-.24.04-.45-.07-.65a.848.848 0 0 0-.51-.39a.751.751 0 0 0-.64.07a.78.78 0 0 0-.39.51l-1.5 5.56c0 .02-.01.06-.02.11c-.01.06-.01.09-.01.12M15.3 9.04c.67-.64 1.49-.97 2.48-.97c.97 0 1.81.34 2.5 1.02c.69.68 1.04 1.51 1.04 2.48c0 .62-.17 1.24-.52 1.85c-.99-.98-2.16-1.47-3.5-1.47h-.31c-.31-1.17-.88-2.14-1.69-2.91m1.61-5.25c0 .23.09.43.26.6s.37.26.6.26c.24 0 .43-.08.59-.25c.16-.17.23-.37.23-.61V1.73c0-.24-.08-.44-.23-.61s-.35-.25-.59-.25c-.23 0-.43.08-.6.25s-.26.37-.26.61zm5.53 2.28c0 .24.09.44.26.6c.14.17.33.25.57.25s.44-.08.6-.25l1.44-1.45c.17-.16.26-.35.26-.59s-.08-.44-.25-.61a.822.822 0 0 0-.61-.25c-.22 0-.41.09-.57.26L22.7 5.47c-.17.16-.26.36-.26.6m.81 11.86c0 .22.08.42.24.6l.66.63c.12.14.31.23.54.24l.01.01h.1c.19 0 .36-.09.53-.26c.17-.16.26-.36.26-.6c0-.23-.09-.43-.26-.61l-.65-.61a.759.759 0 0 0-.58-.27c-.23 0-.43.08-.6.25c-.17.18-.25.39-.25.62m1.45-6.35c0 .23.09.43.27.6c.18.18.38.27.61.27h2.03c.23 0 .43-.09.6-.26s.26-.38.26-.61c0-.23-.08-.43-.25-.59a.853.853 0 0 0-.61-.24h-2.03c-.25 0-.46.08-.63.24c-.17.16-.25.36-.25.59"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DayWindy(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.48 21.1c0 .24.09.44.27.6c.17.17.37.25.61.25h5.88c.26 0 .48.09.68.28a.932.932 0 0 1 0 1.37c-.2.19-.42.29-.68.29c-.26 0-.48-.1-.68-.3a.888.888 0 0 0-.61-.24c-.24 0-.44.08-.6.24c-.16.16-.24.36-.24.6s.08.44.24.6c.53.53 1.16.8 1.89.8c.74 0 1.37-.26 1.88-.78s.78-1.15.78-1.89s-.26-1.37-.78-1.89c-.52-.52-1.15-.78-1.88-.78H2.36a.86.86 0 0 0-.62.25c-.17.17-.26.37-.26.6m0-3.01c0 .23.09.42.27.58c.16.16.37.24.61.24h10.97c.74 0 1.37-.26 1.89-.78c.52-.52.78-1.15.78-1.89s-.26-1.36-.78-1.88c-.52-.51-1.15-.77-1.89-.77c-.76 0-1.38.25-1.88.76c-.16.16-.23.37-.23.61s.08.44.23.6c.15.15.35.23.6.23c.24 0 .44-.08.62-.23c.19-.19.41-.28.68-.28s.49.09.68.28s.29.42.29.68c0 .27-.1.5-.29.69c-.19.19-.42.29-.68.29H2.36a.87.87 0 0 0-.62.26c-.17.18-.26.38-.26.61m5.79-6.54c0-.24.09-.44.26-.62c.18-.16.38-.24.6-.24h2.03c.23 0 .42.08.58.25c.16.17.23.37.23.61s-.08.44-.23.6c-.16.17-.35.25-.58.25H8.13c-.24 0-.44-.08-.61-.25a.816.816 0 0 1-.25-.6m2.85-6.92c0-.23.08-.43.23-.61a.98.98 0 0 1 .64-.24c.22 0 .42.08.59.24l1.43 1.47c.16.15.24.35.24.59s-.08.44-.24.6s-.36.24-.59.24c-.24 0-.44-.08-.59-.24l-1.47-1.43a.86.86 0 0 1-.24-.62m2.56 6.8v-.07c.02-.91.27-1.75.74-2.53A5.29 5.29 0 0 1 15.32 7a5.11 5.11 0 0 1 2.57-.67c.7 0 1.37.14 2.02.42c.64.28 1.2.65 1.66 1.12c.47.47.84 1.02 1.11 1.66s.41 1.32.41 2.02c0 .94-.23 1.8-.69 2.6s-1.09 1.43-1.88 1.89a5.1 5.1 0 0 1-2.6.71h-.21c-.07 0-.13-.02-.17-.07a.241.241 0 0 1-.07-.18v-1.22c0-.13.07-.2.22-.2h.24c.96-.01 1.79-.35 2.47-1.05c.68-.69 1.03-1.52 1.03-2.49c0-.96-.35-1.78-1.04-2.47c-.69-.68-1.52-1.02-2.5-1.02c-.94 0-1.76.32-2.44.98c-.68.65-1.04 1.44-1.08 2.37c0 .06-.03.11-.08.17s-.14.09-.26.09H12.9c-.15.01-.22-.07-.22-.23m4.35 9.88v-1.99c0-.24.08-.44.25-.6s.37-.24.6-.24a.821.821 0 0 1 .85.84v1.99c0 .24-.08.45-.25.62c-.17.17-.37.25-.6.25c-.24 0-.44-.08-.6-.25a.886.886 0 0 1-.25-.62m0-17.48V1.78c0-.23.08-.43.25-.6s.37-.25.6-.25c.24 0 .44.08.6.25s.25.37.25.6v2.04c0 .23-.08.42-.25.58c-.17.15-.37.23-.6.23c-.24 0-.44-.08-.6-.23a.756.756 0 0 1-.25-.57m5.53 13.18c0-.23.08-.42.23-.56a.74.74 0 0 1 .57-.23c.24 0 .44.08.6.23l1.45 1.42c.16.17.24.38.24.61c0 .23-.08.43-.24.59c-.4.31-.8.31-1.2 0l-1.42-1.43c-.16-.16-.23-.38-.23-.63m0-10.93c0-.25.08-.45.23-.59l1.42-1.47c.18-.16.37-.24.59-.24c.23 0 .43.08.6.25c.17.17.25.37.25.6c0 .25-.08.46-.24.62l-1.45 1.43c-.18.16-.38.24-.6.24c-.23 0-.41-.08-.57-.24a.806.806 0 0 1-.23-.6m2.26 5.47c0-.24.08-.44.24-.62c.16-.16.35-.24.57-.24h2.02c.23 0 .43.09.61.26s.26.37.26.6c0 .23-.09.43-.26.6c-.18.17-.38.25-.61.25h-2.02c-.23 0-.42-.08-.58-.25a.806.806 0 0 1-.23-.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Degrees(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.19 9.21c0-.5.18-.93.53-1.28c.36-.36.78-.53 1.28-.53c.49 0 .92.18 1.27.53c.35.36.53.78.53 1.28s-.18.93-.53 1.29c-.35.36-.78.54-1.27.54s-.92-.18-1.28-.54s-.53-.79-.53-1.29m.88 0c0 .26.09.48.27.67c.19.19.41.28.67.28c.26 0 .48-.09.67-.28s.28-.41.28-.67a.87.87 0 0 0-.28-.66a.947.947 0 0 0-.67-.28c-.26 0-.48.09-.67.27c-.18.18-.27.4-.27.67"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DirectionDown(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.77 16.47c0 .22.08.4.25.55l2.4 2.45c.16.16.35.23.58.23c.24 0 .43-.08.59-.23l2.4-2.45c.16-.14.24-.33.24-.55c0-.22-.08-.41-.23-.57s-.34-.23-.56-.23s-.42.08-.57.23l-1.06 1.05v-6.59c0-.22-.08-.41-.24-.56c-.15-.14-.34-.22-.57-.22s-.42.07-.58.22c-.16.15-.24.34-.24.56v6.59l-1.06-1.05a.74.74 0 0 0-.55-.23c-.22 0-.42.08-.57.23s-.23.35-.23.57"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DirectionDownLeft(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.83 16.77a.648.648 0 0 0 .66.67h2.87c.19 0 .35-.06.47-.19c.13-.13.19-.29.19-.48a.59.59 0 0 0-.19-.46a.674.674 0 0 0-.47-.18h-1.24L18 12.24c.12-.14.18-.3.18-.5a.65.65 0 0 0-.18-.46c-.12-.12-.29-.18-.5-.18c-.2 0-.36.06-.48.18l-3.86 3.87V13.9a.648.648 0 0 0-.67-.67c-.19 0-.35.07-.47.2c-.13.13-.19.29-.19.48z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DirectionDownRight(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.04 10.08c0-.3.09-.55.26-.73c.2-.19.46-.28.79-.28c.3 0 .55.09.73.28l6.05 6.05v-1.95c0-.3.1-.55.3-.75s.45-.3.75-.3c.29 0 .54.1.74.31s.3.45.3.75v4.48c0 .3-.1.55-.3.75s-.45.3-.74.3h-4.48c-.29 0-.54-.1-.74-.3s-.3-.45-.3-.75c0-.29.1-.54.3-.73s.45-.29.74-.29h1.93l-6.08-6.06a1.32 1.32 0 0 1-.25-.78"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DirectionLeft(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.09 14.96c0 .37.12.68.37.92l3.84 3.75c.22.25.51.38.85.38c.35 0 .65-.12.89-.35s.37-.53.37-.88s-.12-.65-.37-.89l-1.64-1.64h10.3c.35 0 .64-.12.87-.37s.34-.55.34-.9s-.11-.65-.34-.9s-.52-.38-.87-.39H11.4l1.64-1.66c.24-.24.37-.53.37-.86c0-.35-.12-.65-.37-.89s-.54-.38-.9-.38c-.32 0-.61.14-.85.41l-3.84 3.75c-.24.25-.36.54-.36.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DirectionRight(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.94 14.36c0 .22.08.42.23.57s.34.22.56.2h6.58l-1.03 1.08c-.16.16-.24.35-.24.55c0 .22.08.42.24.57c.16.16.35.23.58.23c.21-.01.39-.1.53-.27l2.45-2.41c.16-.16.23-.35.23-.58a.832.832 0 0 0-.24-.58l-2.47-2.39a.672.672 0 0 0-.54-.23c-.23 0-.42.07-.57.22c-.16.15-.23.34-.23.56c0 .23.08.42.23.57l1.06 1.08h-6.59a.77.77 0 0 0-.56.25c-.15.17-.22.36-.22.58"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DirectionUp(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.95 10.87c-.01.35.1.65.34.9s.53.37.89.36c.34.02.63-.1.88-.37l1.66-1.64v10.3c-.01.35.11.64.36.88s.55.35.92.34c.34.02.64-.09.89-.32s.39-.53.4-.88v-10.3l1.64 1.64c.23.24.53.37.88.37c.36 0 .66-.12.9-.36s.36-.53.36-.89c-.02-.36-.15-.64-.4-.85l-3.74-3.84c-.24-.23-.55-.37-.92-.4c-.37.02-.68.16-.92.41l-3.75 3.81c-.26.22-.39.5-.39.84"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DirectionUpLeft(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.03 14.31V9.84c0-.3.1-.55.3-.75s.45-.3.74-.3h4.48c.29 0 .54.1.74.3s.3.45.3.75c0 .29-.1.53-.3.73s-.45.29-.74.29H13.6l6.06 6.06c.18.21.26.46.26.78c0 .29-.09.53-.26.72c-.2.19-.46.28-.79.28c-.3 0-.55-.09-.73-.28l-6.02-6.05v1.95c0 .3-.1.55-.3.75c-.2.2-.45.3-.75.3c-.29 0-.54-.1-.74-.31s-.3-.46-.3-.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DirectionUpRight(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.05 17.55c0 .3.09.55.26.73c.2.19.46.28.79.28c.3 0 .55-.09.73-.28l6.04-6.05v1.95c0 .3.1.55.3.75c.2.2.45.3.75.3c.29 0 .54-.1.74-.31s.3-.45.3-.75V9.7c0-.3-.1-.55-.3-.75s-.45-.3-.74-.3h-4.5c-.29 0-.54.1-.73.3s-.29.44-.29.75c0 .29.1.54.29.73s.44.29.73.29h1.95l-6.06 6.06c-.17.21-.26.47-.26.77"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dust(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.33 16.58c0-.23.08-.41.23-.56c.16-.15.37-.22.64-.22h5.71c.27 0 .48.07.64.22c.16.15.23.33.23.56c0 .27-.08.49-.23.64s-.37.23-.64.23H8.2c-.27 0-.48-.08-.64-.23s-.23-.36-.23-.64m0-5.91c0-.22.08-.41.23-.55c.16-.15.37-.22.64-.22h2.96c.27 0 .48.07.64.22c.16.14.24.33.24.55c0 .27-.08.48-.24.64c-.16.16-.37.24-.64.24H8.2c-.27 0-.48-.08-.64-.23s-.23-.38-.23-.65m.99 8.87c0-.22.09-.42.28-.6c.18-.18.39-.28.6-.28c.26 0 .46.09.62.27s.24.38.24.61c0 .27-.08.49-.23.65c-.15.16-.36.23-.63.23a.87.87 0 0 1-.61-.24c-.19-.17-.27-.38-.27-.64m1.42-5.93c0-.23.07-.44.22-.61c.15-.18.33-.27.54-.27c.26 0 .48.09.64.27c.16.18.24.38.24.61c0 .27-.08.49-.23.65c-.16.16-.37.23-.65.23c-.23 0-.41-.08-.55-.24s-.21-.37-.21-.64m.99 5.93a.87.87 0 0 1 .88-.88h3.83l.88.88c0 .26-.09.47-.27.64s-.38.24-.61.24h-3.83c-.27 0-.49-.08-.65-.24s-.23-.37-.23-.64m1.32-5.93c0-.22.09-.42.28-.6c.18-.18.39-.28.6-.28h3.83c.26 0 .47.09.63.27c.16.18.24.38.24.61c0 .27-.08.49-.23.65c-.16.16-.37.23-.64.23h-3.83a.87.87 0 0 1-.61-.24c-.18-.16-.27-.37-.27-.64m.76-2.94c0-.22.08-.41.24-.55c.16-.14.37-.22.64-.22h5.71c.23 0 .43.08.61.23c.18.15.27.33.27.54c0 .26-.09.48-.27.64c-.18.16-.38.24-.61.24h-5.71c-.27 0-.49-.08-.65-.24c-.16-.16-.23-.37-.23-.64m2.63 5.91c0-.21.09-.4.27-.55a.926.926 0 0 1 1.22 0a.7.7 0 0 1 .27.55c0 .26-.09.47-.27.63c-.18.16-.38.24-.61.24a.87.87 0 0 1-.61-.24c-.18-.15-.27-.36-.27-.63m1.54 2.96c0-.23.08-.44.24-.61c.16-.18.37-.27.63-.27h1.87c.26 0 .47.09.63.26c.16.17.24.38.24.62c0 .27-.08.49-.23.65c-.15.16-.37.23-.64.23h-1.87c-.27 0-.48-.08-.64-.23c-.15-.16-.23-.38-.23-.65m.87-2.96c0-.21.09-.4.27-.55c.18-.15.38-.23.61-.23h3.07c.22 0 .4.08.54.23c.14.15.22.33.22.55c0 .27-.07.48-.21.64c-.14.16-.32.23-.55.23h-3.07a.87.87 0 0 1-.61-.24c-.18-.15-.27-.36-.27-.63m.44-2.97c0-.22.09-.42.28-.6c.18-.18.39-.28.6-.28h1.96c.21 0 .39.09.54.27c.15.18.23.38.23.61c0 .27-.07.48-.22.64c-.14.16-.33.24-.55.24h-1.96a.87.87 0 0 1-.61-.24c-.18-.16-.27-.37-.27-.64m2.74-2.94c0-.22.07-.4.22-.55c.15-.15.33-.22.55-.22c.27 0 .48.07.64.22c.16.14.24.33.24.55c0 .27-.08.48-.24.64c-.16.16-.37.24-.64.24c-.23 0-.41-.08-.55-.24c-.15-.16-.22-.37-.22-.64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Earthquake(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.25 15.3c0 .16.06.29.17.4c.11.11.25.16.4.16H8.8c.14 0 .27-.04.38-.13c.11-.09.17-.2.2-.34l.9-5.27L11.9 23.3c.02.14.09.26.19.36c.1.09.22.14.36.14c.15 0 .28-.05.38-.14s.17-.21.2-.36l1.25-9.67l1.04 2.8c.04.11.1.2.2.27s.2.1.32.1h.05c.12-.01.23-.05.32-.13c.1-.08.16-.18.19-.31l1.53-6.86l.71 13.18c.01.14.06.27.15.37c.09.1.21.16.36.17c.14.01.27-.02.38-.1a.59.59 0 0 0 .22-.33l1.65-6.94h2.77c.16 0 .29-.05.4-.16c.11-.11.17-.24.17-.4c0-.16-.06-.29-.17-.4a.56.56 0 0 0-.4-.17h-3.23c-.13 0-.25.04-.35.12s-.17.18-.2.31l-.83 3.54l-.72-13.36a.568.568 0 0 0-.16-.37a.5.5 0 0 0-.36-.16c-.14-.01-.27.02-.39.11s-.19.2-.22.34l-2 8.97l-1.16-3.16c-.04-.12-.12-.21-.24-.28s-.24-.1-.36-.08c-.13.01-.24.07-.33.16c-.09.09-.15.21-.17.34l-.98 7.51l-1.53-12.56a.635.635 0 0 0-.19-.35a.585.585 0 0 0-.36-.15a.527.527 0 0 0-.38.12c-.11.09-.18.2-.2.35l-1.48 8.61H5.82c-.16 0-.29.06-.4.17c-.11.11-.17.24-.17.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fahrenheit(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.67 11.01c0-.5.18-.93.53-1.28s.78-.53 1.28-.53c.49 0 .92.18 1.27.53c.35.36.53.78.53 1.28s-.18.93-.53 1.29c-.35.36-.78.54-1.27.54s-.92-.18-1.28-.54a1.79 1.79 0 0 1-.53-1.29m.88 0c0 .26.09.48.27.67c.19.19.41.28.67.28s.48-.09.67-.28s.28-.41.28-.67a.87.87 0 0 0-.28-.66a.947.947 0 0 0-.67-.28c-.26 0-.48.09-.67.27c-.18.18-.27.4-.27.67m4.41 6.89a.514.514 0 0 0 .52.52a.514.514 0 0 0 .52-.52v-3.79h2.86c.14 0 .27-.05.37-.16s.15-.23.15-.38s-.05-.27-.15-.38a.52.52 0 0 0-.38-.15h-2.86v-2.73h3.82c.14 0 .26-.05.36-.15s.14-.23.14-.38s-.05-.27-.14-.38s-.21-.15-.36-.15h-4.77c-.07 0-.1.04-.1.11v8.54z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fire(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.38 21.83c0-.3.1-.55.29-.76c.19-.21.43-.31.7-.31v-.02l13.16.04c.06-.01.1-.02.1-.02c.27.01.51.12.7.33c.19.21.29.47.28.77c0 .3-.1.55-.29.76c-.19.21-.43.31-.7.31v.01L8.59 22.9c-.1.01-.17.02-.22.02a.942.942 0 0 1-.7-.32c-.2-.21-.3-.47-.29-.77m.82-5.46c-.01-.43.04-.93.16-1.52c.06-.3.2-.76.44-1.37c.02-.05.07-.14.13-.28c.01.02.03.03.04.05s.02.02.02.03c.11.44.27.84.49 1.2c.21.32.48.56.82.69c.26.11.63.17 1.1.18h.16c-.33-.33-.59-.67-.79-1c-.3-.52-.49-1.12-.57-1.81c-.06-.54-.03-1.19.09-1.96c.02-.15.12-.49.29-1.01c.15-.47.36-.9.64-1.28c.32-.49.78-.99 1.39-1.51c.37-.31.89-.67 1.56-1.07c.07-.04.18-.11.35-.19v.11c-.24.57-.41 1.15-.49 1.73c-.06.53.02 1.02.24 1.48c.17.36.48.75.92 1.15c.09.09.29.29.6.58c.3.29.54.52.7.68l.25.25c.26-.38.41-.83.44-1.35c.04-.55 0-1.15-.14-1.8c0-.01 0-.04.01-.11c.02.02.13.1.3.24c.56.5.98.95 1.28 1.34c.48.62.83 1.21 1.06 1.74c.19.46.31.92.38 1.4c.06.42.08.77.07 1.05c-.01.78-.1 1.43-.25 1.96c-.07.21-.13.38-.19.52c.25-.07.47-.16.65-.26c.25-.16.45-.37.6-.66c.16-.29.29-.62.38-.98c0-.01.01-.03.03-.05c.01.02.02.05.05.09c.02.04.04.07.05.1c.13.31.22.63.27.97c.08.38.1.75.08 1.13c-.02.29-.07.56-.16.81c-.08.24-.16.43-.22.58c-.19.38-.39.71-.62.98c-.06.07-.11.13-.14.16h-11c-.01-.01-.03-.03-.07-.06s-.06-.05-.08-.07c-.26-.25-.54-.63-.82-1.13c-.08-.15-.18-.38-.29-.69c-.12-.31-.19-.66-.21-1.04"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flood(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.72 20.76c0-.05.01-.12.02-.21v-4.76c.28.49.66.89 1.15 1.19c.49.3 1.03.45 1.61.45c.59 0 1.13-.15 1.62-.45c.49-.3.87-.69 1.15-1.19c.27.49.66.89 1.15 1.19c.49.3 1.03.45 1.62.45c.58 0 1.12-.15 1.61-.45c.49-.3.87-.69 1.15-1.19c.28.49.67.89 1.15 1.19a3 3 0 0 0 1.6.45c.58 0 1.12-.15 1.61-.45c.49-.3.88-.69 1.15-1.19v5.12c-.03.23-.13.43-.3.6c-.17.17-.37.27-.6.3c-.02 0-.05 0-.08.01s-.06.01-.08.01c-.01 0-.04 0-.07-.01c-.03-.01-.06-.01-.08-.01H7.94c-.02 0-.04 0-.08.01c-.03.01-.06.01-.07.01c-.02 0-.05 0-.08-.01s-.06-.01-.07-.01a.993.993 0 0 1-.58-.28a1.13 1.13 0 0 1-.32-.56c-.01-.09-.02-.16-.02-.21m5.51-11.09c0-.16.06-.3.17-.42l2.21-2.22l.03-.02c.01 0 .01 0 .01-.01c.01 0 .01 0 .01-.01c.01 0 .01 0 .01-.01h.01c.01 0 .01 0 .01-.01s0-.01.01-.02h.02l.01-.01h.01l.01-.01h.01l.01-.01h.01c.01-.01.01-.01.02-.01h.01c0-.01.01-.01.02-.01c.01-.01.01-.01.02-.01l.04-.02h.01c.01 0 .01 0 .01-.01h.07l.01-.01h.12c.01 0 .01 0 .02.01h.06c0 .01 0 .01.01.01h.02c.01.01.02.02.03.02l.02.01h.02l.01.01h.01l.01.01c.01 0 .01 0 .01.01h.04c.01.01.01.01.02.01c0 .01 0 .02.01.02l.01.01h.02l.01.01l-.03.02l.01.01l.02.02l2.22 2.22c.12.12.18.26.18.42c0 .16-.06.3-.18.41c-.11.12-.25.18-.41.18c-.16 0-.3-.06-.41-.18l-1.23-1.22v6.9c0 .16-.06.29-.17.4c-.11.11-.25.17-.41.17c-.16 0-.3-.06-.42-.17a.523.523 0 0 1-.17-.4v-6.9l-1.22 1.22c-.12.12-.26.18-.42.18c-.16 0-.3-.06-.41-.18a.6.6 0 0 1-.16-.41"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fog(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.62 21.05c0-.24.08-.45.25-.61c.17-.16.38-.24.63-.24h18.67a.821.821 0 0 1 .85.85c0 .23-.08.43-.25.58c-.17.16-.37.23-.6.23H3.5c-.25 0-.46-.08-.63-.23a.758.758 0 0 1-.25-.58m2.62-3.14c0-.24.09-.44.26-.6c.15-.15.35-.23.59-.23h18.67c.23 0 .42.08.58.24c.16.16.23.35.23.59s-.08.44-.23.6c-.16.17-.35.25-.58.25H6.09c-.24 0-.44-.08-.6-.25a.816.816 0 0 1-.25-.6m.13-2.39c0 .09.05.13.15.13h1.43c.06 0 .13-.05.2-.16c.24-.52.59-.94 1.06-1.27c.47-.33.99-.52 1.55-.56l.55-.07c.11 0 .17-.06.17-.18l.07-.5c.11-1.08.56-1.98 1.37-2.7c.81-.72 1.76-1.08 2.85-1.08c1.08 0 2.02.36 2.83 1.07c.8.71 1.26 1.61 1.37 2.68l.08.57c0 .11.07.17.2.17h1.59c.64 0 1.23.17 1.76.52s.92.8 1.18 1.37c.07.11.14.16.21.16h1.43c.12 0 .17-.07.14-.23c-.29-1.02-.88-1.86-1.74-2.51c-.87-.65-1.86-.97-2.97-.97h-.32c-.33-1.33-1.03-2.42-2.1-3.27s-2.28-1.27-3.65-1.27c-1.4 0-2.64.44-3.73 1.32s-1.78 2-2.09 3.36c-.85.2-1.6.6-2.24 1.21c-.64.61-1.09 1.33-1.34 2.18v-.04c-.01 0-.01.03-.01.07m1.61 8.59c0-.24.09-.43.26-.59c.15-.15.35-.23.6-.23h18.68c.24 0 .44.08.6.23c.17.16.25.35.25.58c0 .24-.08.44-.25.61c-.17.17-.37.25-.6.25H7.84c-.23 0-.43-.09-.6-.26a.773.773 0 0 1-.26-.59"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GaleWarning(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.67 24.6V7.45h1.03V24.6zm1.73-2.16v-7.41l8.65 3.69zm0-7.58V7.45l8.65 3.69z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hail(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.64 16.9c0 1.33.46 2.47 1.39 3.43c.93.96 2.06 1.47 3.4 1.53c.11 0 .17-.06.17-.17v-1.34c0-.11-.06-.17-.17-.17c-.86-.04-1.58-.38-2.18-1.02c-.6-.64-.9-1.39-.9-2.26c0-.83.28-1.54.84-2.16c.56-.61 1.26-.97 2.09-1.07l.53-.03c.13 0 .2-.06.2-.19l.06-.53c.11-1.08.56-1.99 1.37-2.71c.81-.73 1.76-1.09 2.85-1.09s2.04.36 2.85 1.09c.81.73 1.27 1.63 1.39 2.71l.08.58c0 .11.06.17.18.17h1.61c.89 0 1.66.32 2.31.96c.65.64.98 1.39.98 2.27c0 .87-.3 1.62-.9 2.26c-.6.64-1.33.98-2.18 1.02c-.13 0-.2.06-.2.17v1.34c0 .11.07.17.2.17c.87-.02 1.67-.26 2.4-.71c.73-.45 1.31-1.05 1.73-1.8c.42-.75.63-1.57.63-2.44c0-.89-.22-1.72-.67-2.47c-.44-.75-1.05-1.35-1.81-1.78S21.29 12 20.4 12h-.32c-.32-1.34-1.03-2.43-2.1-3.28s-2.3-1.28-3.68-1.28c-1.41 0-2.66.44-3.75 1.31a5.83 5.83 0 0 0-2.1 3.35c-1.11.26-2.02.83-2.73 1.73s-1.08 1.92-1.08 3.07m5.45 7.2c.09.21.25.37.46.46c.2.1.41.11.62.02c.22-.09.36-.24.45-.45c.1-.22.11-.43.02-.64c-.08-.21-.24-.35-.45-.44c-.2-.11-.4-.12-.61-.03a.85.85 0 0 0-.46.47c-.11.17-.11.37-.03.61m.63-2.82c0 .16.05.31.15.45c.1.15.26.25.46.32c.19.11.4.12.62.01c.22-.1.37-.3.44-.6l.9-3.38c.06-.25.04-.47-.08-.67a.722.722 0 0 0-.53-.36a.843.843 0 0 0-.71.12c-.15.1-.26.25-.32.44L10.77 21c-.04.16-.05.25-.05.28m1.86 5.59c0 .12.02.22.06.29c.09.22.24.37.45.45c.09.05.2.08.33.08c.06 0 .16-.02.3-.06c.22-.08.38-.23.47-.45c.1-.22.1-.44 0-.66a.88.88 0 0 0-.45-.46c-.2-.09-.4-.09-.62 0a.81.81 0 0 0-.41.36c-.09.16-.13.31-.13.45m.73-2.61c0 .37.21.61.63.73a.855.855 0 0 0 .62-.04c.21-.08.35-.27.42-.57l1.67-6.29c.06-.24.04-.45-.06-.65a.76.76 0 0 0-.49-.38c-.08-.02-.17-.03-.27-.03a.9.9 0 0 0-.48.15c-.16.1-.26.25-.3.44l-1.71 6.34c-.02.14-.03.24-.03.3m3.43-.46c0 .12.02.23.08.32c.08.19.23.34.44.44c.11.04.23.07.35.07c.06 0 .16-.02.3-.06c.21-.08.37-.23.46-.44c.07-.22.07-.43-.01-.63a.839.839 0 0 0-.42-.45c-.23-.11-.44-.12-.65-.03a.85.85 0 0 0-.46.47c-.06.1-.09.2-.09.31m.73-2.57c0 .14.05.29.16.45c.11.16.26.27.45.33c.16.03.25.05.27.05c.09 0 .22-.03.37-.1c.2-.09.33-.27.4-.52l.9-3.34l.03-.26c0-.16-.05-.31-.15-.46a.784.784 0 0 0-.45-.31c-.09-.02-.18-.03-.26-.03c-.16 0-.32.05-.47.15s-.25.25-.31.45l-.9 3.36z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Horizon(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.93 20.97c0-.26.09-.47.28-.62c.14-.16.35-.23.63-.23h18.34c.25 0 .46.08.64.24c.18.16.26.37.26.61s-.09.45-.27.62s-.39.27-.63.27H5.84a.87.87 0 0 1-.64-.27a.826.826 0 0 1-.27-.62m1.97-8.29c0-.26.08-.47.23-.63c.17-.18.38-.26.65-.26c.23 0 .43.09.6.26l1.5 1.5a.858.858 0 0 1 0 1.25c-.15.18-.35.27-.6.27s-.47-.09-.64-.27l-1.5-1.5c-.16-.15-.24-.35-.24-.62m2.93 5.59c-.04.16.01.23.15.23h1.49c.07 0 .14-.06.22-.17c.3-.64.74-1.14 1.33-1.52s1.24-.56 1.96-.56c.73 0 1.39.19 1.99.56s1.05.88 1.35 1.52c.08.11.16.17.23.17h1.48c.13 0 .18-.08.15-.23a5.225 5.225 0 0 0-1.95-2.76c-.96-.71-2.04-1.06-3.25-1.06c-1.2 0-2.28.35-3.23 1.06c-.93.71-1.58 1.63-1.92 2.76m4.31-6.46V9.68c0-.25.08-.46.24-.64c.16-.18.37-.26.61-.26a.874.874 0 0 1 .88.9v2.14a.852.852 0 0 1-.88.89c-.24 0-.45-.09-.61-.26s-.24-.39-.24-.64m5.72 2.37c0-.24.08-.45.25-.63l1.54-1.5c.16-.18.36-.26.62-.26c.24 0 .44.08.6.25s.23.38.23.64s-.08.47-.23.62l-1.48 1.5c-.17.17-.36.26-.56.28c-.23.02-.44-.06-.65-.24s-.32-.41-.32-.66"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HorizonAlt(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.11 15.9c0-.26.09-.47.27-.62c.14-.15.35-.23.62-.23h2.08c.24 0 .45.08.6.24c.16.16.24.36.24.6s-.08.44-.24.61c-.16.17-.36.25-.6.25H5c-.24 0-.45-.08-.63-.25a.76.76 0 0 1-.26-.6m.92 5.08c0-.23.09-.43.26-.61c.16-.16.37-.23.61-.23h18.21c.24 0 .45.08.62.24c.17.16.25.36.25.6s-.09.44-.26.61s-.38.26-.61.26H5.91c-.24 0-.44-.09-.61-.26s-.27-.37-.27-.61M7.08 8.81c0-.26.08-.45.23-.59c.17-.18.38-.27.62-.27s.44.09.61.27l1.44 1.46c.18.16.26.36.26.6a.821.821 0 0 1-.84.85c-.22 0-.42-.08-.6-.24L7.31 9.42c-.15-.14-.23-.34-.23-.61m2.6 7.09c0 .87.18 1.65.53 2.33c.03.09.11.14.24.14h1.67c.07 0 .12-.02.14-.06c.02-.04-.01-.1-.07-.16c-.53-.65-.8-1.4-.8-2.25c0-.99.36-1.84 1.07-2.54c.71-.7 1.56-1.05 2.55-1.05c.99 0 1.84.35 2.55 1.05s1.05 1.55 1.05 2.54c0 .86-.27 1.61-.8 2.25c-.04.06-.06.1-.06.12c-.01.03 0 .06.03.07c.02.02.06.03.1.03h1.7c.09 0 .16-.05.21-.14c.38-.71.56-1.48.56-2.33c0-.96-.24-1.85-.72-2.67c-.48-.82-1.13-1.47-1.95-1.95s-1.71-.72-2.67-.72s-1.85.24-2.67.72c-.82.48-1.47 1.13-1.95 1.95c-.48.82-.71 1.71-.71 2.67m4.47-7.93V5.88c0-.24.08-.44.25-.62c.17-.18.37-.26.61-.26s.44.09.62.26c.17.17.26.38.26.62v2.09c0 .24-.09.44-.26.62c-.18.18-.38.26-.62.26s-.44-.09-.61-.26a.86.86 0 0 1-.25-.62m5.62 2.31c0-.24.08-.44.24-.6l1.44-1.46c.17-.18.38-.27.62-.27c.25 0 .46.08.62.25c.17.17.25.37.25.61c0 .26-.08.46-.23.61l-1.51 1.47c-.16.15-.36.22-.59.22a.807.807 0 0 1-.61-.22a.838.838 0 0 1-.23-.61m2.33 5.62c0-.27.08-.47.24-.62c.14-.15.34-.23.59-.23h2.09c.24 0 .45.08.62.24c.17.16.26.36.26.6s-.09.44-.26.61a.86.86 0 0 1-.62.25h-2.09c-.23 0-.43-.08-.59-.25a.814.814 0 0 1-.24-.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hot(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.14 14.76c0-.24.09-.44.27-.61c.17-.17.38-.25.62-.25h2.06c.24 0 .44.08.6.25c.17.17.25.37.25.61c0 .25-.08.46-.25.63c-.17.17-.37.25-.6.25H5.03c-.25 0-.46-.08-.63-.25a.823.823 0 0 1-.26-.63m2.92-7.02c0-.23.08-.44.24-.62c.2-.16.41-.25.66-.25c.21 0 .41.08.59.25L10 8.6c.17.16.25.36.25.6s-.08.44-.25.6a.8.8 0 0 1-.6.25c-.26 0-.46-.08-.61-.25l-1.5-1.44a.966.966 0 0 1-.23-.62m2.61 7.02c0-.71.14-1.39.43-2.04c.28-.65.67-1.22 1.14-1.69c.48-.47 1.05-.85 1.7-1.13s1.34-.43 2.06-.43c.96 0 1.84.24 2.66.72c.82.48 1.47 1.12 1.94 1.94c.47.81.71 1.69.71 2.63c0 .15-.01.29-.03.42c-.28-.17-.6-.25-.97-.25c-.24 0-.48.05-.72.15c.01-.07.01-.18.01-.32c0-.98-.35-1.81-1.06-2.5c-.71-.69-1.56-1.04-2.55-1.04c-.99 0-1.83.35-2.54 1.04s-1.05 1.52-1.05 2.5v.18c-.29.02-.57.1-.84.25c-.01.01-.04.03-.1.07s-.12.09-.19.14s-.14.11-.23.19s-.18.16-.26.24c-.08-.32-.11-.68-.11-1.07m.03 3.65v-.15c-.01-.09 0-.2.02-.33c.05-.36.19-.7.42-1.02c.13-.16.22-.27.27-.31a.21.21 0 0 1 .04-.06l.04-.04a.44.44 0 0 0 .12-.1l.08-.08c.03-.03.06-.04.07-.06c.05-.05.1-.08.14-.1l.17-.11c.14-.09.31-.14.5-.14h.03c.1 0 .19.01.26.03c.03.01.07.03.13.07v.01a.798.798 0 0 1 .47.71c0 .17-.05.31-.14.42c-.06.09-.14.17-.22.23l-.06.03l-.04.02l-.06.04c-.04.03-.07.06-.1.08s-.06.06-.11.11c-.04.05-.08.09-.11.14c-.03.04-.06.1-.09.15c-.03.06-.05.12-.05.17v.15c.04.15.08.27.11.36c.07.14.18.28.34.44c.01.02.09.1.24.25c.86.78 1.27 1.62 1.21 2.5c-.02.3-.09.59-.21.87s-.26.51-.43.7c-.16.19-.29.33-.39.43c-.1.09-.18.16-.25.21c-.01.01-.03.02-.06.04s-.06.04-.07.04c-.08.04-.15.06-.22.07c-.09.01-.15.02-.2.02c-.3 0-.54-.1-.71-.3a.788.788 0 0 1-.18-.59c.02-.22.13-.4.33-.53l.01-.03c.01-.01.03-.03.05-.04c.02-.02.04-.04.07-.06c.03-.02.06-.05.08-.08a.7.7 0 0 0 .08-.1c.03-.04.06-.08.08-.12c.03-.04.06-.09.08-.14c.03-.05.05-.1.07-.15c.02-.05.03-.1.05-.16c.01-.06.02-.12.02-.17c.02-.2-.03-.4-.15-.6c-.05-.11-.12-.22-.22-.33c-.07-.08-.12-.13-.15-.16c-.09-.11-.14-.17-.15-.18c-.02-.01-.04-.03-.07-.06s-.05-.04-.06-.05c-.15-.14-.26-.26-.34-.36c-.12-.16-.21-.26-.24-.32c-.19-.26-.32-.52-.39-.78c-.04-.14-.07-.25-.08-.32c0-.02-.01-.05-.02-.08s-.01-.05-.01-.08m3.88-.33c0-.06 0-.1.01-.14c.02-.22.09-.43.2-.64c.11-.21.22-.39.35-.53c.13-.14.25-.27.38-.38s.23-.19.31-.25l.12-.07c.15-.09.32-.14.5-.14c.11 0 .21.01.3.03c.01 0 .02.01.05.02c.03.02.05.03.06.04c.01 0 .02 0 .03.01c.01 0 .03.02.07.05c.2.12.32.3.38.54v.12c0 .03-.01.07-.02.12s-.02.08-.02.09c-.07.23-.21.39-.42.5c-.33.22-.51.45-.53.69a.81.81 0 0 0 0 .22c.02.12.08.25.17.39c.11.16.19.27.24.32c.16.16.25.25.28.27c.12.11.28.28.47.51c.54.65.79 1.32.74 2c-.02.3-.09.59-.21.87s-.26.51-.43.7c-.16.18-.3.33-.4.43c-.11.1-.19.17-.25.21l-.12.08c-.11.04-.17.06-.21.07c-.11.01-.18.02-.2.02h-.03a.92.92 0 0 1-.19-.02c-.02 0-.05 0-.08-.01s-.06-.01-.07-.01c-.01 0-.02 0-.03-.01c-.01-.01-.02-.01-.04-.02c-.01-.01-.02-.01-.03-.01c-.15-.11-.24-.17-.26-.21a.744.744 0 0 1 .15-1.13l.03-.04c.02-.02.05-.05.09-.08l.12-.12l.13-.16l.12-.19l.09-.22l.04-.24c.01-.4-.22-.82-.69-1.27c-.19-.17-.33-.31-.44-.43c-.55-.69-.8-1.34-.76-1.98m.54-11.16V4.85c0-.24.09-.45.26-.62c.17-.17.38-.25.61-.25c.24 0 .45.08.62.25c.17.17.25.38.25.62v2.07c0 .24-.08.43-.25.59s-.37.23-.61.23a.89.89 0 0 1-.61-.23a.75.75 0 0 1-.27-.59m3.36 11.01c.02-.22.09-.43.2-.64c.11-.21.22-.39.35-.53c.13-.14.25-.27.38-.38c.12-.11.22-.19.3-.25l.13-.07c.02-.02.06-.04.1-.08c.11-.04.24-.07.38-.07c.34 0 .59.13.77.38c.05.07.08.14.1.23c.01.02.02.05.02.08v.11c0 .31-.15.55-.45.7c-.32.21-.49.44-.52.69c-.04.34.19.74.68 1.2c.88.77 1.28 1.61 1.23 2.5c-.02.3-.09.59-.21.87s-.27.51-.43.7c-.17.19-.3.33-.39.43s-.18.16-.25.21c-.16.1-.3.15-.41.16c-.03.01-.08.01-.15.01c-.3 0-.53-.1-.69-.3c-.15-.17-.21-.37-.19-.59s.13-.4.32-.53c.03-.01.09-.06.18-.14s.18-.21.29-.38c.1-.18.16-.35.17-.53c.02-.4-.22-.82-.7-1.27c-.41-.36-.73-.75-.94-1.16c-.24-.42-.33-.88-.27-1.35m2.29-8.72c0-.25.08-.45.23-.6l1.46-1.48c.18-.17.38-.25.61-.25c.24 0 .44.09.61.26s.26.38.26.61c0 .25-.09.46-.26.62L21.2 9.81c-.18.17-.38.25-.61.25c-.23 0-.43-.08-.58-.25a.83.83 0 0 1-.24-.6m2.3 5.55c0-.22.09-.42.26-.61c.16-.17.35-.25.58-.25h2.06c.24 0 .45.09.62.26s.27.37.27.6c0 .24-.09.45-.26.62c-.18.17-.38.26-.63.26h-2.06a.77.77 0 0 1-.6-.25a.884.884 0 0 1-.24-.63"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Humidity(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.56 17.19c0-.88.24-1.89.72-3.03s1.1-2.25 1.86-3.31c1.56-2.06 2.92-3.62 4.06-4.67l.75-.72c.25.26.53.5.83.72c.41.42 1.04 1.11 1.88 2.09s1.57 1.85 2.17 2.65c.71 1.01 1.32 2.1 1.81 3.25s.74 2.16.74 3.03c0 1-.19 1.95-.58 2.86c-.39.91-.91 1.7-1.57 2.36c-.66.66-1.45 1.19-2.37 1.58c-.92.39-1.89.59-2.91.59c-1 0-1.95-.19-2.86-.57c-.91-.38-1.7-.89-2.36-1.55c-.66-.65-1.19-1.44-1.58-2.35s-.59-1.89-.59-2.93m2.26-2.93c0 .83.17 1.49.52 1.99c.35.49.88.74 1.59.74c.72 0 1.25-.25 1.61-.74c.35-.49.53-1.15.54-1.99c-.01-.84-.19-1.5-.54-2c-.35-.49-.89-.74-1.61-.74c-.71 0-1.24.25-1.59.74c-.35.5-.52 1.16-.52 2m1.57 0v-.35c0-.08.01-.19.02-.33s.02-.25.05-.32s.05-.16.09-.24c.04-.08.09-.15.15-.18c.07-.04.14-.06.23-.06c.14 0 .25.04.33.12s.14.21.17.38c.03.18.05.32.06.45s.01.3.01.52c0 .23 0 .4-.01.52s-.03.27-.06.45c-.03.17-.09.3-.17.38s-.19.12-.33.12c-.09 0-.16-.02-.23-.06a.335.335 0 0 1-.15-.18c-.04-.08-.07-.17-.09-.24c-.02-.08-.04-.19-.05-.32c-.01-.14-.02-.25-.02-.32zm.59 7.75h1.32l4.99-10.74h-1.35zm4.3-2.99c.01.84.2 1.5.55 2c.35.49.89.74 1.6.74c.72 0 1.25-.25 1.6-.74c.35-.49.52-1.16.53-2c-.01-.84-.18-1.5-.53-1.99c-.35-.49-.88-.74-1.6-.74c-.71 0-1.25.25-1.6.74c-.36.49-.54 1.15-.55 1.99m1.57 0c0-.23 0-.4.01-.52s.03-.27.06-.45s.09-.3.17-.38s.19-.12.33-.12c.09 0 .17.02.24.06c.07.04.12.1.16.19c.04.09.07.17.1.24s.04.18.05.32l.01.32v.69l-.01.32l-.05.32l-.1.24l-.16.19l-.24.06c-.14 0-.25-.04-.33-.12s-.14-.21-.17-.38c-.03-.18-.05-.33-.06-.45s-.01-.3-.01-.53"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hurricane(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.08 14.53v-.02c-.01-.08 0-.2 0-.37c.01-.16.04-.43.1-.81c.06-.38.14-.76.25-1.15s.28-.84.51-1.35c.23-.51.5-.99.82-1.44c.32-.45.74-.92 1.26-1.39c.52-.47 1.1-.89 1.73-1.24c.16-.09.32-.11.49-.06s.3.15.38.31c.09.16.11.32.06.5c-.05.17-.15.31-.3.39a7.42 7.42 0 0 0-3.03 3.05c.54-.25 1.08-.38 1.63-.38c1.07 0 2 .38 2.77 1.15c.77.77 1.15 1.69 1.15 2.76v.24c0 .08-.02.25-.04.52s-.06.52-.11.77s-.13.56-.23.93c-.11.37-.23.73-.38 1.06c-.15.33-.34.7-.58 1.1s-.51.77-.81 1.12c-.3.35-.66.7-1.08 1.05c-.43.35-.89.67-1.39.95c-.09.06-.2.08-.31.08c-.26 0-.45-.12-.58-.35a.606.606 0 0 1-.06-.49c.05-.17.15-.3.31-.38C14.98 20.33 16 19.3 16.7 18c-.54.26-1.11.38-1.71.38c-.69 0-1.34-.17-1.94-.52c-.6-.34-1.07-.81-1.43-1.41c-.35-.58-.53-1.22-.54-1.92m1.7-.05c0 .61.22 1.13.65 1.57c.43.43.95.65 1.56.65c.57 0 1.06-.19 1.49-.57c.42-.38.66-.85.73-1.41l.01-.23c0-.02 0-.04.01-.05c-.01-.6-.23-1.11-.66-1.52c-.43-.42-.96-.62-1.57-.62c-.56 0-1.04.18-1.46.54s-.66.82-.73 1.36l-.02.25v.03z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HurricaneWarning(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.8 24.6V7.45h1.13V24.6zm1.93-3.08v-6.6h8.55v6.6zm0-7.47v-6.6h8.55v6.6zm2.36 5.47h3.81v-2.51h-3.81zm0-7.47h3.81v-2.5h-3.81z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lightning(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.96 24.51h.39l6.88-10.18c.09-.18.04-.27-.15-.27h-2.84l2.99-5.45c.09-.18.02-.27-.2-.27h-3.81c-.11 0-.2.06-.29.18l-2.78 7.4c-.02.18.04.27.19.27h2.75zm8.5-6.33h.27l5.22-7.67c.05-.08.06-.15.04-.2s-.08-.07-.17-.07h-2.1l2.18-4.03c.12-.2.06-.3-.18-.3h-2.74c-.13 0-.23.06-.3.19l-2.08 5.48c-.03.09-.03.16.01.21c.04.05.1.07.19.07h2.04z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LunarEclipse(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.8 14.62c0-.93.23-1.8.7-2.6s1.1-1.44 1.9-1.91s1.67-.7 2.6-.7c.94 0 1.81.23 2.61.7c.8.47 1.43 1.1 1.9 1.9c.47.8.7 1.67.7 2.61s-.23 1.81-.7 2.61c-.47.8-1.1 1.43-1.9 1.9c-.8.47-1.67.7-2.61.7s-1.8-.23-2.6-.7s-1.43-1.1-1.9-1.9c-.47-.8-.7-1.67-.7-2.61m4.45-3.4c.87.11 1.6.49 2.19 1.15c.59.66.89 1.44.89 2.33c0 .83-.26 1.56-.78 2.2s-1.18 1.04-1.98 1.21c.2.02.34.04.43.04c.98 0 1.81-.35 2.5-1.04c.69-.69 1.04-1.52 1.04-2.5c0-.96-.35-1.78-1.04-2.47a3.442 3.442 0 0 0-2.5-1.02c-.26.02-.51.05-.75.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Meteor(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.09 19.39c0-.13 0-.23.01-.29v-.3c0-.03.01-.07.02-.12c.01-.05.02-.08.02-.09v-.12l.08-.37c0-.01 0-.01.01-.02v-.02l.04-.14c.01-.01.01-.01.01-.02c.01-.01.01-.02.01-.03v-.03c.04-.12.07-.22.1-.28c0-.01.01-.02.02-.03c.01-.01.02-.06.05-.15c.17-.38.38-.74.63-1.08l.06-.07c.01-.01.02-.02.03-.04c.01-.02.02-.03.03-.04c.01-.01.03-.03.07-.06a.21.21 0 0 1 .04-.06c.02-.02.03-.04.04-.06c.04-.02.06-.05.07-.07c.01-.01.03-.02.07-.06l.07-.07l7.6-8.33l-.38 2.2l6.82-7.29l-4.18 8.14l4.18-3.16l-3.79 7.6l2.71-1.87l-4.68 8.33c0 .01-.01.02-.02.04s-.02.04-.03.05c-.01.01-.01.02-.02.04s-.01.03-.02.05c-.01.01-.01.02-.02.05c-.01.02-.02.04-.02.05a5.09 5.09 0 0 1-1.86 2.02c-.81.51-1.7.76-2.67.76c-.92 0-1.77-.23-2.55-.68c-.78-.46-1.4-1.07-1.86-1.86s-.69-1.6-.69-2.52m1.2 0c0 1.08.38 1.99 1.14 2.75c.76.76 1.68 1.14 2.75 1.14c.82 0 1.56-.24 2.22-.71c.66-.47 1.13-1.09 1.41-1.84c.17-.43.25-.87.25-1.34c0-1.07-.38-1.99-1.13-2.75c-.76-.76-1.67-1.13-2.75-1.13c-1 0-1.87.33-2.6 1c-.41.36-.72.78-.95 1.28c-.23.51-.34 1.04-.34 1.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltFirstQuarter(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0-1.52.3-2.98.89-4.37s1.4-2.58 2.4-3.59s2.2-1.81 3.59-2.4s2.84-.89 4.37-.89s2.98.3 4.37.89s2.59 1.4 3.6 2.4s1.81 2.2 2.4 3.59s.89 2.84.89 4.37s-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.6s-2.2 1.81-3.6 2.4s-2.85.89-4.37.89s-2.98-.3-4.37-.89s-2.58-1.4-3.59-2.4s-1.81-2.2-2.4-3.6s-.89-2.84-.89-4.37M14.8 24.51h.19c1.37 0 2.67-.27 3.91-.8s2.31-1.25 3.21-2.15s1.61-1.97 2.15-3.21s.8-2.54.8-3.91c0-1.36-.27-2.66-.8-3.9s-1.25-2.31-2.15-3.21s-1.97-1.61-3.21-2.15s-2.54-.8-3.91-.8h-.19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltFull(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0-1.53.3-3 .89-4.39s1.4-2.59 2.4-3.6s2.2-1.81 3.6-2.4s2.85-.89 4.37-.89c1.53 0 3 .3 4.39.89s2.59 1.4 3.6 2.4s1.81 2.2 2.4 3.6s.89 2.85.89 4.39c0 1.52-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.6s-2.2 1.81-3.6 2.4s-2.85.89-4.39.89c-1.52 0-2.98-.3-4.37-.89s-2.59-1.4-3.6-2.4s-1.81-2.2-2.4-3.6s-.89-2.84-.89-4.37m1.2 0c0 1.36.27 2.66.8 3.9s1.25 2.32 2.15 3.22s1.97 1.61 3.22 2.15s2.55.8 3.9.8c1.37 0 2.67-.27 3.91-.8s2.31-1.25 3.22-2.15s1.62-1.97 2.15-3.22s.8-2.55.8-3.9c0-1.82-.45-3.5-1.35-5.05s-2.13-2.77-3.68-3.68s-3.23-1.35-5.05-1.35c-1.36 0-2.66.27-3.9.8S8.79 6.41 7.89 7.32S6.28 9.3 5.74 10.54s-.8 2.54-.8 3.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltNew(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0 2.04.5 3.93 1.51 5.65s2.37 3.1 4.1 4.1S12.96 25.7 15 25.7s3.92-.5 5.65-1.51s3.09-2.37 4.09-4.1s1.51-3.61 1.51-5.65s-.5-3.92-1.51-5.65s-2.37-3.09-4.09-4.09S17.04 3.19 15 3.19s-3.92.51-5.65 1.51s-3.1 2.37-4.1 4.09s-1.51 3.61-1.51 5.65"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltThirdQuarter(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0-1.52.3-2.98.89-4.37s1.4-2.58 2.4-3.59s2.2-1.81 3.59-2.4s2.84-.89 4.37-.89s2.98.3 4.37.89s2.59 1.4 3.6 2.4s1.81 2.2 2.4 3.59s.89 2.84.89 4.37s-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.6s-2.2 1.81-3.6 2.4s-2.85.89-4.37.89s-2.98-.3-4.37-.89s-2.58-1.4-3.59-2.4s-1.81-2.2-2.4-3.6s-.89-2.84-.89-4.37m1.2 0c0 1.37.27 2.67.8 3.91s1.25 2.31 2.15 3.21s1.97 1.61 3.21 2.15s2.54.8 3.9.8h.21V4.39H15c-1.36 0-2.66.27-3.9.8S8.79 6.44 7.89 7.34s-1.61 1.97-2.15 3.21s-.8 2.54-.8 3.89"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltWaningCrescentFive(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0-1.52.3-2.98.89-4.37s1.4-2.58 2.4-3.59s2.2-1.81 3.59-2.4s2.84-.89 4.37-.89s2.98.3 4.37.89s2.59 1.4 3.6 2.4s1.81 2.2 2.4 3.59s.89 2.84.89 4.37s-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.6s-2.2 1.81-3.6 2.4s-2.85.89-4.37.89s-2.98-.3-4.37-.89s-2.58-1.4-3.59-2.4s-1.81-2.2-2.4-3.6s-.89-2.84-.89-4.37m1.2 0c0 1.27.23 2.49.7 3.66s1.09 2.2 1.89 3.08s1.75 1.61 2.85 2.19s2.28.94 3.52 1.08c-1.75-1.04-2.98-2.39-3.7-4.07s-1.08-3.66-1.08-5.93c0-2.07.44-4 1.32-5.78s2.1-3.2 3.66-4.24c-1.26.11-2.46.45-3.59 1.02s-2.1 1.3-2.92 2.19s-1.46 1.92-1.93 3.11s-.72 2.4-.72 3.69"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltWaningCrescentFour(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0-1.52.3-2.98.89-4.37s1.4-2.58 2.4-3.59s2.2-1.81 3.59-2.4s2.84-.89 4.37-.89s2.98.3 4.37.89s2.59 1.4 3.6 2.4s1.81 2.2 2.4 3.59s.89 2.84.89 4.37s-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.6s-2.2 1.81-3.6 2.4s-2.85.89-4.37.89s-2.98-.3-4.37-.89s-2.58-1.4-3.59-2.4s-1.81-2.2-2.4-3.6s-.89-2.84-.89-4.37m1.2 0c0 1.33.25 2.6.75 3.81s1.18 2.26 2.03 3.15s1.87 1.61 3.05 2.17s2.43.87 3.74.94c-1.24-1.19-2.11-2.63-2.61-4.31s-.75-3.6-.75-5.76c0-1.93.31-3.78.92-5.54s1.47-3.26 2.56-4.5c-1.77.07-3.39.56-4.88 1.47S7.09 8 6.23 9.51s-1.29 3.17-1.29 4.93"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltWaningCrescentOne(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0-1.52.3-2.98.89-4.37s1.4-2.58 2.4-3.59s2.2-1.81 3.59-2.4s2.84-.89 4.37-.89s2.98.3 4.37.89s2.59 1.4 3.6 2.4s1.81 2.2 2.4 3.59s.89 2.84.89 4.37s-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.6s-2.2 1.81-3.6 2.4s-2.85.89-4.37.89s-2.98-.3-4.37-.89s-2.58-1.4-3.59-2.4s-1.81-2.2-2.4-3.6s-.89-2.84-.89-4.37m1.2 0c0 1.37.27 2.67.8 3.91s1.25 2.31 2.15 3.21s1.97 1.61 3.21 2.15s2.54.8 3.9.8h.07c-.59-2.67-.88-6.02-.88-10.06c0-3.39.3-6.74.91-10.05H15c-1.36 0-2.66.27-3.9.8S8.79 6.44 7.89 7.34s-1.61 1.97-2.15 3.21s-.8 2.54-.8 3.89"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltWaningCrescentSix(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0-1.52.3-2.98.89-4.37s1.4-2.58 2.4-3.59s2.2-1.81 3.59-2.4s2.84-.89 4.37-.89s2.98.3 4.37.89s2.59 1.4 3.6 2.4s1.81 2.2 2.4 3.59s.89 2.84.89 4.37s-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.6s-2.2 1.81-3.6 2.4s-2.85.89-4.37.89s-2.98-.3-4.37-.89s-2.58-1.4-3.59-2.4s-1.81-2.2-2.4-3.6s-.89-2.84-.89-4.37m1.2 0c0 2.48.8 4.66 2.41 6.53s3.62 3.01 6.03 3.41c-1.01-.5-1.86-1.1-2.56-1.82s-1.25-1.5-1.63-2.37s-.66-1.77-.83-2.7s-.26-1.95-.26-3.06c0-2.11.5-4.06 1.49-5.84s2.37-3.16 4.12-4.12c-1.63.21-3.11.77-4.45 1.7S6.87 8.3 6.1 9.76s-1.16 3.01-1.16 4.68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltWaningCrescentThree(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0-1.52.3-2.98.89-4.37s1.4-2.58 2.4-3.59s2.2-1.81 3.59-2.4s2.84-.89 4.37-.89s2.98.3 4.37.89s2.59 1.4 3.6 2.4s1.81 2.2 2.4 3.59s.89 2.84.89 4.37s-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.6s-2.2 1.81-3.6 2.4s-2.85.89-4.37.89s-2.98-.3-4.37-.89s-2.58-1.4-3.59-2.4s-1.81-2.2-2.4-3.6s-.89-2.84-.89-4.37m1.2 0c0 1.34.26 2.62.77 3.85s1.21 2.29 2.08 3.19s1.92 1.62 3.13 2.16s2.48.83 3.81.87c-1.71-2.32-2.56-5.68-2.56-10.06c0-1.87.23-3.67.69-5.41s1.11-3.29 1.95-4.64c-1.8.03-3.45.5-4.96 1.41s-2.7 2.13-3.58 3.65s-1.33 3.19-1.33 4.98"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltWaningCrescentTwo(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0-1.52.3-2.98.89-4.37s1.4-2.58 2.4-3.59s2.2-1.81 3.59-2.4s2.84-.89 4.37-.89s2.98.3 4.37.89s2.59 1.4 3.6 2.4s1.81 2.2 2.4 3.59s.89 2.84.89 4.37s-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.6s-2.2 1.81-3.6 2.4s-2.85.89-4.37.89s-2.98-.3-4.37-.89s-2.58-1.4-3.59-2.4s-1.81-2.2-2.4-3.6s-.89-2.84-.89-4.37m1.2 0c0 1.36.26 2.65.79 3.89s1.24 2.3 2.12 3.2s1.95 1.62 3.17 2.15s2.52.81 3.87.82c-1.16-2.47-1.74-5.83-1.74-10.06c0-3.61.6-6.96 1.8-10.05c-1.36 0-2.65.27-3.89.81s-2.3 1.25-3.19 2.15s-1.61 1.97-2.14 3.2s-.79 2.54-.79 3.89"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltWaningGibbousFive(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0-1.52.3-2.98.89-4.37s1.4-2.58 2.4-3.59s2.2-1.81 3.59-2.4s2.84-.89 4.37-.89s2.98.3 4.37.89s2.59 1.4 3.6 2.4s1.81 2.2 2.4 3.59s.89 2.84.89 4.37s-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.6s-2.2 1.81-3.6 2.4s-2.85.89-4.37.89s-2.98-.3-4.37-.89s-2.58-1.4-3.59-2.4s-1.81-2.2-2.4-3.6s-.89-2.84-.89-4.37m1.2 0c0 1.37.27 2.67.8 3.91s1.25 2.31 2.15 3.21s1.97 1.61 3.21 2.15s2.54.8 3.9.8c.26 0 .46 0 .59-.01c.77-1.33 1.3-2.83 1.57-4.5s.42-3.52.42-5.55c0-4.01-.68-7.36-2.04-10.03c-.11-.01-.29-.01-.54-.01c-1.36 0-2.66.27-3.9.8S8.79 6.44 7.89 7.34s-1.61 1.97-2.15 3.21s-.8 2.54-.8 3.89"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltWaningGibbousFour(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0-1.52.3-2.98.89-4.37s1.4-2.58 2.4-3.59s2.2-1.81 3.59-2.4s2.84-.89 4.37-.89s2.98.3 4.37.89s2.59 1.4 3.6 2.4s1.81 2.2 2.4 3.59s.89 2.84.89 4.37s-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.6s-2.2 1.81-3.6 2.4s-2.85.89-4.37.89s-2.98-.3-4.37-.89s-2.58-1.4-3.59-2.4s-1.81-2.2-2.4-3.6s-.89-2.84-.89-4.37m1.2 0c0 1.37.27 2.67.8 3.91s1.25 2.31 2.15 3.21s1.97 1.61 3.21 2.15s2.54.8 3.9.8h.38c.6-.64 1.1-1.37 1.5-2.19s.71-1.67.9-2.58s.33-1.77.41-2.59s.12-1.73.12-2.7c0-1.88-.24-3.7-.73-5.46s-1.25-3.28-2.3-4.59H15c-1.36 0-2.66.27-3.9.8S8.79 6.44 7.89 7.34s-1.61 1.97-2.15 3.21s-.8 2.54-.8 3.89"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltWaningGibbousOne(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.43c0-2.04.51-3.92 1.52-5.65S7.64 5.7 9.37 4.69s3.61-1.5 5.65-1.5s3.92.5 5.65 1.5s3.09 2.36 4.09 4.09s1.5 3.61 1.5 5.65s-.5 3.93-1.5 5.65s-2.36 3.1-4.09 4.11s-3.61 1.52-5.65 1.52s-3.93-.51-5.65-1.52s-3.1-2.38-4.11-4.11s-1.52-3.61-1.52-5.65m1.2 0c0 1.36.27 2.66.81 3.9S7 20.65 7.9 21.55s1.97 1.62 3.22 2.15s2.55.81 3.9.81c.86 0 1.62-.09 2.29-.28c.83-.44 1.55-.96 2.17-1.57s1.1-1.22 1.46-1.85s.64-1.33.86-2.09s.36-1.48.43-2.14s.1-1.37.1-2.15c0-.93-.1-1.84-.3-2.74s-.51-1.79-.93-2.67s-.99-1.72-1.72-2.5s-1.57-1.45-2.54-1.99c-.4-.09-1.01-.13-1.82-.13c-1.36 0-2.66.26-3.9.79S8.8 6.43 7.9 7.32s-1.62 1.97-2.15 3.2s-.81 2.54-.81 3.91"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltWaningGibbousSix(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0-1.52.3-2.98.89-4.37s1.4-2.58 2.4-3.59s2.2-1.81 3.59-2.4s2.84-.89 4.37-.89s2.98.3 4.37.89s2.59 1.4 3.6 2.4s1.81 2.2 2.4 3.59s.89 2.84.89 4.37s-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.6s-2.2 1.81-3.6 2.4s-2.85.89-4.37.89s-2.98-.3-4.37-.89s-2.58-1.4-3.59-2.4s-1.81-2.2-2.4-3.6s-.89-2.84-.89-4.37m1.2 0c0 1.37.27 2.67.8 3.91s1.25 2.31 2.15 3.21s1.97 1.61 3.21 2.15s2.54.8 3.9.8h.38c.67-2.49 1.01-5.84 1.01-10.06c0-3.82-.34-7.17-1.03-10.05h-.37c-1.36 0-2.66.27-3.9.8s-2.3 1.24-3.2 2.14s-1.61 1.97-2.15 3.21s-.8 2.54-.8 3.89"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltWaningGibbousThree(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0-1.52.3-2.98.89-4.37s1.4-2.58 2.4-3.59s2.2-1.81 3.59-2.4s2.84-.89 4.37-.89s2.98.3 4.37.89s2.59 1.4 3.6 2.4s1.81 2.2 2.4 3.59s.89 2.84.89 4.37s-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.6s-2.2 1.81-3.6 2.4s-2.85.89-4.37.89s-2.98-.3-4.37-.89s-2.58-1.4-3.59-2.4s-1.81-2.2-2.4-3.6s-.89-2.84-.89-4.37m1.2 0c0 1.37.27 2.67.8 3.91s1.25 2.31 2.15 3.21s1.97 1.61 3.21 2.15s2.54.8 3.9.8c.33 0 .58 0 .73-.01c.78-.61 1.44-1.31 1.96-2.11s.92-1.66 1.18-2.57s.44-1.79.54-2.63s.15-1.75.15-2.74c0-1.91-.32-3.76-.97-5.54s-1.65-3.28-3.02-4.49c-.13-.01-.32-.01-.59-.01c-1.36 0-2.66.27-3.9.8S8.79 6.44 7.89 7.34s-1.61 1.97-2.15 3.21s-.8 2.54-.8 3.89"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltWaningGibbousTwo(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0-1.52.3-2.98.89-4.37s1.4-2.58 2.4-3.59s2.2-1.81 3.59-2.4s2.84-.89 4.37-.89s2.98.3 4.37.89s2.59 1.4 3.6 2.4s1.81 2.2 2.4 3.59s.89 2.84.89 4.37s-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.6s-2.2 1.81-3.6 2.4s-2.85.89-4.37.89s-2.98-.3-4.37-.89s-2.58-1.4-3.59-2.4s-1.81-2.2-2.4-3.6s-.89-2.84-.89-4.37m1.2 0c0 1.37.27 2.67.8 3.91s1.25 2.31 2.15 3.21s1.97 1.61 3.21 2.15s2.54.8 3.9.8c.36 0 .76-.02 1.2-.07c.93-.57 1.71-1.24 2.35-2.03s1.12-1.64 1.43-2.56s.53-1.8.65-2.65s.18-1.77.18-2.75c0-1.25-.15-2.46-.46-3.64s-.84-2.34-1.59-3.49s-1.69-2.11-2.81-2.89c-.41-.02-.73-.03-.95-.03c-1.36 0-2.66.27-3.9.8S8.79 6.44 7.89 7.34s-1.61 1.97-2.15 3.21s-.8 2.54-.8 3.89"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltWaxingCrescentFive(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0-1.52.3-2.98.89-4.37s1.4-2.58 2.4-3.59s2.2-1.81 3.59-2.4s2.84-.89 4.37-.89s2.98.3 4.37.89s2.59 1.4 3.6 2.4s1.81 2.2 2.4 3.59s.89 2.84.89 4.37s-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.6s-2.2 1.81-3.6 2.4s-2.85.89-4.37.89s-2.98-.3-4.37-.89s-2.58-1.4-3.59-2.4s-1.81-2.2-2.4-3.6s-.89-2.84-.89-4.37M15.15 4.39c1.83 3.22 2.74 6.57 2.74 10.05c0 4.16-.88 7.51-2.65 10.06c1.34-.03 2.61-.32 3.82-.86s2.25-1.27 3.13-2.16s1.57-1.95 2.09-3.18s.78-2.51.78-3.86c0-1.8-.44-3.46-1.33-5s-2.08-2.75-3.6-3.65s-3.18-1.37-4.98-1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltWaxingCrescentFour(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.75 14.44c0-1.52.3-2.98.89-4.37s1.4-2.58 2.4-3.59s2.2-1.81 3.59-2.4s2.85-.89 4.37-.89s2.98.3 4.37.89s2.59 1.4 3.6 2.4s1.81 2.2 2.4 3.59s.89 2.84.89 4.37s-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.6s-2.2 1.81-3.6 2.4s-2.85.89-4.37.89s-2.98-.3-4.37-.89s-2.58-1.4-3.59-2.4s-1.81-2.2-2.4-3.6s-.89-2.84-.89-4.37m11.6-10.05c1.05 1.27 1.91 2.75 2.57 4.44s.99 3.56.99 5.61c0 4.39-1.14 7.75-3.43 10.06c1.31-.06 2.55-.37 3.74-.92s2.2-1.28 3.05-2.18s1.53-1.95 2.04-3.16s.75-2.48.75-3.81c0-1.78-.43-3.43-1.3-4.94s-2.04-2.73-3.53-3.65s-3.11-1.39-4.88-1.45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltWaxingCrescentOne(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0-1.52.3-2.98.89-4.37s1.4-2.58 2.4-3.59s2.2-1.81 3.59-2.4s2.84-.89 4.37-.89s2.98.3 4.37.89s2.59 1.4 3.6 2.4s1.81 2.2 2.4 3.59s.89 2.84.89 4.37s-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.6s-2.2 1.81-3.6 2.4s-2.85.89-4.37.89s-2.98-.3-4.37-.89s-2.58-1.4-3.59-2.4s-1.81-2.2-2.4-3.6s-.89-2.84-.89-4.37m12.64-9.96c1.64.9 2.99 2.2 4.04 3.9s1.57 3.72 1.57 6.06c0 4.9-1.72 8.2-5.16 9.89c2.35-.44 4.31-1.59 5.87-3.44s2.34-4 2.34-6.45c0-1.66-.38-3.21-1.14-4.66s-1.8-2.63-3.13-3.57s-2.77-1.5-4.39-1.73"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltWaxingCrescentSix(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0-1.52.3-2.98.89-4.37s1.4-2.58 2.4-3.59s2.2-1.81 3.59-2.4s2.84-.89 4.37-.89s2.98.3 4.37.89s2.59 1.4 3.6 2.4s1.81 2.2 2.4 3.59s.89 2.84.89 4.37s-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.6s-2.2 1.81-3.6 2.4s-2.85.89-4.37.89s-2.98-.3-4.37-.89s-2.58-1.4-3.59-2.4s-1.81-2.2-2.4-3.6s-.89-2.84-.89-4.37M15.01 4.39c1.23 3.54 1.85 6.89 1.85 10.05c0 3.93-.59 7.28-1.77 10.06c1.35-.01 2.64-.28 3.87-.81s2.3-1.25 3.19-2.15s1.6-1.97 2.12-3.21s.79-2.54.79-3.9s-.27-2.66-.8-3.9s-1.25-2.31-2.15-3.21s-1.97-1.61-3.21-2.15s-2.54-.78-3.89-.78"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltWaxingCrescentThree(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0-1.52.3-2.98.89-4.37s1.4-2.58 2.4-3.59s2.2-1.81 3.59-2.4s2.84-.89 4.37-.89s2.98.3 4.37.89s2.59 1.4 3.6 2.4s1.81 2.2 2.4 3.59s.89 2.84.89 4.37s-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.6s-2.2 1.81-3.6 2.4s-2.85.89-4.37.89s-2.98-.3-4.37-.89s-2.58-1.4-3.59-2.4s-1.81-2.2-2.4-3.6s-.89-2.84-.89-4.37M15.58 4.41c1.28 1.1 2.32 2.51 3.13 4.22s1.22 3.65 1.22 5.82c0 4.64-1.39 7.99-4.16 10.05c1.28-.1 2.49-.43 3.63-1s2.13-1.29 2.96-2.18s1.49-1.93 1.97-3.13s.73-2.44.73-3.74c0-1.75-.42-3.38-1.26-4.89s-1.99-2.72-3.44-3.64s-3.05-1.42-4.78-1.51"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltWaxingCrescentTwo(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0-1.52.3-2.98.89-4.37s1.4-2.58 2.4-3.59s2.2-1.81 3.59-2.4s2.84-.89 4.37-.89s2.98.3 4.37.89s2.59 1.4 3.6 2.4s1.81 2.2 2.4 3.59s.89 2.84.89 4.37s-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.6s-2.2 1.81-3.6 2.4s-2.85.89-4.37.89s-2.98-.3-4.37-.89s-2.58-1.4-3.59-2.4s-1.81-2.2-2.4-3.6s-.89-2.84-.89-4.37M15.9 4.42c1.48.99 2.7 2.34 3.65 4.05s1.42 3.7 1.42 5.97c0 4.8-1.6 8.13-4.79 9.99c1.65-.2 3.15-.76 4.5-1.68s2.42-2.12 3.2-3.58s1.17-3.03 1.17-4.72c0-1.72-.41-3.32-1.22-4.8s-1.91-2.69-3.31-3.61s-2.93-1.47-4.62-1.62"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltWaxingGibbousFive(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0-1.53.3-3 .89-4.39s1.4-2.59 2.4-3.6s2.2-1.81 3.6-2.4s2.85-.89 4.37-.89c1.53 0 3 .3 4.39.89s2.59 1.4 3.6 2.4s1.81 2.2 2.4 3.6s.89 2.85.89 4.39c0 1.52-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.6s-2.2 1.81-3.6 2.4s-2.85.89-4.39.89c-1.52 0-2.98-.3-4.37-.89s-2.59-1.4-3.6-2.4s-1.81-2.2-2.4-3.6s-.89-2.84-.89-4.37m3.97 0c0 4.76 1.66 8.02 4.97 9.79c.73.19 1.51.28 2.33.28c1.37 0 2.67-.27 3.91-.8s2.31-1.25 3.22-2.15s1.62-1.97 2.15-3.22s.8-2.55.8-3.9c0-1.82-.45-3.5-1.35-5.05s-2.13-2.77-3.68-3.68s-3.23-1.35-5.05-1.35c-.68 0-1.3.05-1.85.16c-1.63.94-2.95 2.27-3.95 3.99s-1.5 3.71-1.5 5.93"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltWaxingGibbousFour(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0-1.53.3-3 .89-4.39s1.4-2.59 2.4-3.6s2.2-1.81 3.6-2.4s2.85-.89 4.37-.89c1.53 0 3 .3 4.39.89s2.59 1.4 3.6 2.4s1.81 2.2 2.4 3.6s.89 2.85.89 4.39c0 1.52-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.6s-2.2 1.81-3.6 2.4s-2.85.89-4.39.89c-1.52 0-2.98-.3-4.37-.89s-2.59-1.4-3.6-2.4s-1.81-2.2-2.4-3.6s-.89-2.84-.89-4.37m5.39-.01c0 4.68 1.48 8 4.45 9.98c.44.07.91.1 1.42.1c1.37 0 2.67-.27 3.91-.8s2.31-1.25 3.22-2.15s1.62-1.97 2.15-3.22s.8-2.55.8-3.9c0-1.82-.45-3.5-1.35-5.05s-2.13-2.77-3.68-3.68S16.82 4.36 15 4.36c-.45 0-.84.02-1.19.06c-1.4 1.06-2.53 2.46-3.39 4.2s-1.29 3.67-1.29 5.81"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltWaxingGibbousOne(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0-1.52.3-2.98.89-4.37s1.4-2.58 2.4-3.59s2.2-1.81 3.59-2.4s2.84-.89 4.37-.89s2.98.3 4.37.89s2.59 1.4 3.6 2.4s1.81 2.2 2.4 3.59s.89 2.84.89 4.37s-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.6s-2.2 1.81-3.6 2.4s-2.85.89-4.37.89s-2.98-.3-4.37-.89s-2.58-1.4-3.59-2.4s-1.81-2.2-2.4-3.6s-.89-2.84-.89-4.37m9.64 0c0 3.81.41 7.16 1.23 10.06h.38c1.37 0 2.67-.27 3.91-.8s2.31-1.25 3.21-2.15s1.61-1.97 2.15-3.21s.8-2.54.8-3.91c0-1.36-.27-2.66-.8-3.9s-1.25-2.31-2.15-3.21s-1.97-1.61-3.21-2.15s-2.54-.8-3.91-.8h-.34c-.84 3.62-1.27 6.97-1.27 10.07"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltWaxingGibbousSix(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0-1.53.3-3 .89-4.39s1.4-2.59 2.4-3.6s2.2-1.81 3.6-2.4s2.85-.89 4.37-.89c1.53 0 3 .3 4.39.89s2.59 1.4 3.6 2.4s1.81 2.2 2.4 3.6s.89 2.85.89 4.39c0 1.52-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.6s-2.2 1.81-3.6 2.4s-2.85.89-4.39.89c-1.52 0-2.98-.3-4.37-.89s-2.59-1.4-3.6-2.4s-1.81-2.2-2.4-3.6s-.89-2.84-.89-4.37m2.68 0c0 .78.05 1.5.15 2.15c.1.65.29 1.35.57 2.09s.66 1.41 1.13 2.03s1.09 1.24 1.88 1.86s1.7 1.2 2.77 1.71c.67.15 1.37.22 2.09.22c1.37 0 2.67-.27 3.91-.8s2.31-1.25 3.22-2.15s1.62-1.97 2.15-3.22s.8-2.55.8-3.9c0-1.82-.45-3.5-1.35-5.05s-2.13-2.77-3.68-3.68s-3.23-1.35-5.05-1.35c-.85 0-1.62.09-2.31.26c-.89.44-1.67.9-2.35 1.39S9.1 7.01 8.66 7.54s-.82 1.08-1.13 1.62s-.54 1.12-.69 1.74s-.26 1.2-.32 1.75s-.1 1.15-.1 1.79"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltWaxingGibbousThree(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0-2.03.5-3.91 1.51-5.64s2.37-3.1 4.1-4.1s3.61-1.51 5.64-1.51c1.52 0 2.98.3 4.37.89s2.59 1.4 3.6 2.4s1.81 2.2 2.4 3.59s.89 2.84.89 4.37s-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.6s-2.2 1.81-3.6 2.4s-2.85.89-4.37.89s-2.98-.3-4.37-.89s-2.59-1.4-3.59-2.4s-1.8-2.2-2.4-3.6s-.89-2.84-.89-4.37m6.8 0c0 4.44 1.17 7.78 3.5 10.02c.07 0 .17 0 .3.01s.25.02.35.02s.2.01.3.01c1.36 0 2.66-.27 3.9-.8s2.32-1.25 3.22-2.15s1.61-1.97 2.15-3.21s.8-2.54.8-3.91c0-1.36-.27-2.66-.8-3.9s-1.25-2.31-2.15-3.21s-1.97-1.61-3.22-2.15s-2.55-.8-3.9-.8c-.36 0-.63.01-.81.03c-1.08 1.22-1.96 2.69-2.64 4.42s-1 3.61-1 5.62"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonAltWaxingGibbousTwo(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0-1.52.3-2.98.89-4.37s1.4-2.58 2.4-3.59s2.2-1.81 3.59-2.4s2.84-.89 4.37-.89s2.98.3 4.37.89s2.59 1.4 3.6 2.4s1.81 2.2 2.4 3.59s.89 2.84.89 4.37s-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.6s-2.2 1.81-3.6 2.4s-2.85.89-4.37.89s-2.98-.3-4.37-.89s-2.58-1.4-3.59-2.4s-1.81-2.2-2.4-3.6s-.89-2.84-.89-4.37m8.22 0c0 4.15.81 7.5 2.42 10.05c.15.01.35.01.62.01c1.37 0 2.67-.27 3.91-.8s2.31-1.25 3.21-2.15s1.61-1.97 2.15-3.21s.8-2.54.8-3.91c0-1.36-.27-2.66-.8-3.9s-1.25-2.31-2.15-3.21s-1.97-1.61-3.21-2.15s-2.54-.8-3.91-.8c-.23 0-.42 0-.54.01c-1.67 3.17-2.5 6.52-2.5 10.06"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonFirstQuarter(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.01 25.71c2.04 0 3.92-.5 5.65-1.51s3.09-2.37 4.09-4.1s1.51-3.61 1.51-5.65s-.5-3.92-1.51-5.65s-2.37-3.09-4.09-4.09s-3.61-1.51-5.65-1.51z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonFull(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0 2.04.5 3.93 1.51 5.65s2.37 3.1 4.1 4.1S12.96 25.7 15 25.7s3.92-.5 5.65-1.51s3.09-2.37 4.09-4.1s1.51-3.61 1.51-5.65s-.5-3.92-1.51-5.65s-2.37-3.09-4.09-4.09S17.04 3.19 15 3.19s-3.92.51-5.65 1.51s-3.1 2.37-4.1 4.09s-1.51 3.61-1.51 5.65"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonNew(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0-1.53.3-3 .89-4.39s1.4-2.59 2.4-3.6s2.2-1.81 3.6-2.4s2.85-.89 4.37-.89c1.53 0 3 .3 4.39.89s2.59 1.4 3.6 2.4s1.81 2.2 2.4 3.6s.89 2.85.89 4.39c0 1.52-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.6s-2.2 1.81-3.6 2.4s-2.85.89-4.39.89c-1.52 0-2.98-.3-4.37-.89s-2.59-1.4-3.6-2.4s-1.81-2.2-2.4-3.6s-.89-2.84-.89-4.37m1.2 0c0 1.36.27 2.66.8 3.9s1.25 2.32 2.15 3.22s1.97 1.61 3.22 2.15s2.55.8 3.9.8c1.37 0 2.67-.27 3.91-.8s2.31-1.25 3.22-2.15s1.62-1.97 2.15-3.22s.8-2.55.8-3.9c0-1.82-.45-3.5-1.35-5.05s-2.13-2.77-3.68-3.68s-3.23-1.35-5.05-1.35c-1.36 0-2.66.27-3.9.8S8.79 6.41 7.89 7.32S6.28 9.3 5.74 10.54s-.8 2.54-.8 3.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonThirdQuarter(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0 2.04.5 3.93 1.51 5.65s2.37 3.1 4.09 4.1s3.61 1.51 5.65 1.51V3.19c-2.04 0-3.92.5-5.65 1.51S6.26 7.07 5.25 8.8s-1.51 3.6-1.51 5.64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonWaningCrescentFive(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0 2.04.5 3.93 1.51 5.65s2.37 3.1 4.09 4.1s3.61 1.51 5.65 1.51c-2.59-.79-4.48-2.13-5.69-4.02s-1.81-4.3-1.81-7.24c0-1.39.2-2.7.6-3.95c.4-1.25.94-2.34 1.63-3.27s1.48-1.75 2.37-2.44s1.86-1.22 2.89-1.59c-2.04 0-3.92.5-5.65 1.51S6.26 7.07 5.25 8.8s-1.51 3.6-1.51 5.64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonWaningCrescentFour(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0 2.04.5 3.93 1.51 5.65s2.37 3.1 4.1 4.1S12.96 25.7 15 25.7c-2.07-1.01-3.59-2.45-4.56-4.33S9 17.19 9 14.44c0-2.53.56-4.78 1.69-6.75s2.57-3.47 4.31-4.5c-2.04 0-3.93.5-5.65 1.51s-3.1 2.37-4.1 4.09s-1.51 3.61-1.51 5.65"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonWaningCrescentOne(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0 2.04.5 3.93 1.51 5.65s2.37 3.1 4.09 4.1s3.61 1.51 5.65 1.51c-1-3.14-1.49-6.9-1.49-11.26c0-3.43.5-7.18 1.49-11.25c-2.04 0-3.92.5-5.65 1.51S6.26 7.07 5.25 8.8s-1.51 3.6-1.51 5.64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonWaningCrescentSix(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0 2.04.5 3.93 1.51 5.65s2.37 3.1 4.09 4.1s3.61 1.51 5.65 1.51c-1.46-.56-2.72-1.18-3.79-1.88s-1.93-1.39-2.57-2.1s-1.15-1.49-1.53-2.34s-.64-1.66-.77-2.42s-.2-1.6-.2-2.52c0-.65.03-1.26.08-1.81s.16-1.14.32-1.77s.38-1.21.64-1.75s.63-1.09 1.08-1.66s.98-1.1 1.59-1.57s1.34-.95 2.21-1.42s1.85-.89 2.93-1.27c-2.04 0-3.92.5-5.65 1.51S6.26 7.07 5.25 8.8s-1.51 3.6-1.51 5.64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonWaningCrescentThree(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0 2.04.5 3.93 1.51 5.65s2.37 3.1 4.09 4.1s3.61 1.51 5.65 1.51c-2.99-2.33-4.48-6.09-4.48-11.26c0-2.32.42-4.46 1.25-6.4s1.91-3.56 3.23-4.85c-2.04 0-3.92.5-5.65 1.51S6.26 7.07 5.25 8.8s-1.51 3.6-1.51 5.64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonWaningCrescentTwo(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.44c0 2.04.5 3.93 1.51 5.65s2.37 3.1 4.1 4.1S12.96 25.7 15 25.7c-2.01-2.74-3.02-6.5-3.02-11.26c0-3.98 1.01-7.73 3.02-11.25c-2.04 0-3.93.5-5.65 1.51s-3.1 2.37-4.1 4.09s-1.51 3.61-1.51 5.65"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonWaningGibbousFive(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.47c0 2.03.5 3.91 1.51 5.63s2.37 3.09 4.09 4.09s3.6 1.51 5.63 1.51c2.17-2.75 3.25-6.5 3.25-11.24c0-3.96-1.08-7.71-3.25-11.24c-2.03 0-3.91.5-5.63 1.5S6.26 7.1 5.25 8.83s-1.51 3.61-1.51 5.64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonWaningGibbousFour(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.47c0 1.52.3 2.98.89 4.37s1.39 2.58 2.4 3.59s2.2 1.8 3.59 2.4s2.84.89 4.37.89c2.89-2.39 4.34-6.14 4.34-11.24c0-2.34-.41-4.47-1.22-6.36S16.26 4.6 15 3.25c-2.03 0-3.91.5-5.64 1.51S6.25 7.12 5.24 8.84s-1.5 3.6-1.5 5.63"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonWaningGibbousOne(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.49c0 1.22.19 2.4.56 3.54s.91 2.17 1.6 3.09s1.5 1.72 2.42 2.42s1.95 1.23 3.09 1.6s2.32.56 3.54.56c5.03-1.4 7.54-5.14 7.54-11.22c0-1.18-.14-2.3-.42-3.37s-.65-2.01-1.13-2.83s-1.04-1.57-1.68-2.24s-1.34-1.24-2.06-1.68s-1.47-.81-2.26-1.07c-1.52 0-2.98.3-4.37.89S8.02 5.57 7.02 6.57s-1.8 2.19-2.39 3.57s-.89 2.83-.89 4.35"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonWaningGibbousSix(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.46c0 2.04.5 3.92 1.51 5.65s2.37 3.09 4.09 4.09s3.61 1.51 5.65 1.51c1.44-3.08 2.15-6.83 2.15-11.25c0-3.46-.72-7.2-2.15-11.24c-1.52 0-2.98.3-4.37.89S8.03 5.5 7.03 6.5s-1.8 2.2-2.4 3.59s-.89 2.84-.89 4.37"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonWaningGibbousThree(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.49c0 1.22.19 2.4.56 3.54s.91 2.17 1.6 3.09s1.5 1.72 2.42 2.42s1.95 1.23 3.09 1.6s2.32.56 3.54.56c3.61-2.07 5.42-5.81 5.42-11.22c0-1.31-.15-2.56-.45-3.74s-.71-2.24-1.23-3.17s-1.1-1.75-1.72-2.46s-1.3-1.33-2.02-1.86c-1.52 0-2.98.3-4.37.89S8 5.53 7 6.54s-1.8 2.2-2.39 3.59s-.87 2.83-.87 4.36"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonWaningGibbousTwo(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.49c0 1.22.19 2.4.56 3.54s.91 2.17 1.6 3.09s1.5 1.72 2.42 2.42s1.95 1.23 3.09 1.6s2.32.56 3.54.56c4.33-1.73 6.49-5.47 6.49-11.22c0-1.39-.18-2.7-.54-3.93s-.85-2.31-1.47-3.23s-1.31-1.71-2.06-2.39s-1.56-1.23-2.42-1.66c-2.03 0-3.91.5-5.63 1.5S6.25 7.14 5.24 8.86s-1.5 3.6-1.5 5.63"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonWaxingCrescentFive(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.99 25.71c2.04 0 3.93-.5 5.65-1.51s3.1-2.37 4.1-4.1s1.51-3.61 1.51-5.65s-.5-3.92-1.51-5.65s-2.37-3.09-4.1-4.09s-3.61-1.51-5.65-1.51c1.67 2.9 2.5 6.65 2.5 11.25c0 2.33-.17 4.43-.52 6.3s-1 3.51-1.98 4.96"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonWaxingCrescentFour(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.01 25.71c2.04 0 3.92-.5 5.65-1.51s3.09-2.37 4.09-4.1s1.51-3.61 1.51-5.65s-.5-3.92-1.51-5.65s-2.37-3.09-4.09-4.09s-3.61-1.51-5.65-1.51c1.29 1.39 2.24 3.07 2.84 5.05s.91 4.05.91 6.2c0 .88-.03 1.69-.08 2.44s-.16 1.55-.32 2.41s-.38 1.65-.64 2.37s-.63 1.43-1.09 2.15s-1.01 1.33-1.62 1.89"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonWaxingCrescentOne(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.01 25.71c2.04 0 3.92-.5 5.65-1.51s3.09-2.37 4.09-4.1s1.51-3.61 1.51-5.65s-.5-3.92-1.51-5.65s-2.37-3.09-4.09-4.09s-3.61-1.51-5.65-1.51c1.32.52 2.48 1.2 3.47 2.06s1.78 1.79 2.35 2.82s.99 2.07 1.27 3.13s.41 2.14.41 3.24c0 .64-.02 1.26-.07 1.84c-.05.58-.15 1.2-.29 1.87s-.33 1.28-.56 1.86s-.54 1.15-.92 1.74s-.83 1.11-1.35 1.58s-1.14.92-1.87 1.33s-1.55.75-2.44 1.04"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonWaxingCrescentSix(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M0 1537q209 0 386-103t280-280t103-386t-103-385.5T386 103T0 0q171 297 171 768q0 239-35.5 430T0 1537z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonWaxingCrescentThree(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.01 25.71c2.04 0 3.92-.5 5.65-1.51s3.09-2.37 4.09-4.1s1.51-3.61 1.51-5.65s-.5-3.92-1.51-5.65s-2.37-3.09-4.09-4.09s-3.61-1.51-5.65-1.51c1.71 1.26 2.97 2.9 3.78 4.91S20 12.24 20 14.44c0 .9-.03 1.73-.1 2.5s-.21 1.59-.43 2.47s-.51 1.68-.86 2.4s-.83 1.42-1.45 2.12s-1.33 1.28-2.15 1.78"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonWaxingCrescentTwo(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.01 25.71c2.04 0 3.92-.5 5.65-1.51s3.09-2.37 4.09-4.1s1.51-3.61 1.51-5.65s-.5-3.92-1.51-5.65s-2.37-3.09-4.09-4.09s-3.61-1.51-5.65-1.51c1.1.59 2.07 1.32 2.89 2.19s1.47 1.82 1.95 2.83s.83 2.03 1.05 3.07s.34 2.09.34 3.16c0 .91-.04 1.76-.13 2.54s-.27 1.63-.53 2.53s-.62 1.71-1.06 2.43s-1.04 1.42-1.82 2.09s-1.67 1.22-2.69 1.67"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonWaxingGibbousFive(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.65 14.43c0 1.24.08 2.38.25 3.41s.44 2.05.83 3.04s.95 1.89 1.67 2.71s1.6 1.52 2.62 2.12c1.52 0 2.98-.3 4.37-.89s2.59-1.4 3.6-2.4s1.81-2.2 2.4-3.6s.89-2.85.89-4.39s-.3-2.99-.89-4.38s-1.4-2.58-2.4-3.59s-2.2-1.81-3.6-2.4s-2.85-.89-4.37-.89c-1.67 1.14-2.98 2.72-3.94 4.74s-1.43 4.18-1.43 6.52"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonWaxingGibbousFour(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.73 14.43c0 1.19.07 2.29.2 3.3s.35 2 .67 2.99s.76 1.9 1.33 2.75s1.27 1.59 2.09 2.25c1.53 0 2.99-.3 4.38-.89s2.58-1.4 3.59-2.4s1.81-2.2 2.4-3.6s.89-2.85.89-4.39c0-2.04-.5-3.93-1.51-5.65s-2.37-3.1-4.1-4.1s-3.61-1.51-5.65-1.51c-1.35 1.3-2.4 2.94-3.16 4.93s-1.13 4.08-1.13 6.32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonWaxingGibbousOne(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.93 14.44c-.02 4.53.33 8.29 1.04 11.27c2.04.01 3.93-.49 5.65-1.49s3.1-2.36 4.11-4.08s1.52-3.61 1.53-5.65c.01-2.04-.49-3.93-1.49-5.65c-1-1.73-2.36-3.1-4.08-4.11s-3.6-1.52-5.64-1.53c-.73 3.71-1.11 7.46-1.12 11.24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonWaxingGibbousSix(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.58 14.43c0 1.03.06 1.97.18 2.83s.32 1.73.62 2.59s.69 1.65 1.16 2.34s1.1 1.35 1.85 1.96s1.63 1.12 2.63 1.55c1.53 0 2.99-.3 4.38-.89s2.58-1.4 3.59-2.4s1.81-2.2 2.4-3.6s.89-2.85.89-4.39c0-2.04-.5-3.93-1.51-5.65s-2.37-3.1-4.1-4.1s-3.61-1.51-5.65-1.51c-1.99 1-3.56 2.51-4.72 4.55s-1.72 4.28-1.72 6.72"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonWaxingGibbousThree(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.8 14.43c0 2.39.24 4.52.71 6.39s1.31 3.5 2.51 4.89c1.52 0 2.98-.3 4.37-.89s2.59-1.4 3.6-2.4s1.81-2.2 2.4-3.6s.89-2.85.89-4.39s-.3-2.99-.89-4.38s-1.4-2.58-2.4-3.59s-2.2-1.81-3.6-2.4s-2.85-.89-4.37-.89C14 4.63 13.21 6.33 12.65 8.3s-.85 4-.85 6.13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonWaxingGibbousTwo(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.85 14.44c0 4.77.71 8.52 2.14 11.26c2.04 0 3.93-.5 5.65-1.51s3.1-2.37 4.1-4.1s1.51-3.61 1.51-5.65s-.5-3.92-1.51-5.65s-2.37-3.09-4.1-4.09s-3.61-1.51-5.65-1.51c-1.42 3.42-2.14 7.17-2.14 11.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonWaxingSix(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.99 25.71c2.04 0 3.93-.5 5.65-1.51s3.1-2.37 4.1-4.1s1.51-3.61 1.51-5.65s-.5-3.92-1.51-5.65s-2.37-3.09-4.1-4.09s-3.61-1.51-5.65-1.51c1.67 2.9 2.5 6.65 2.5 11.25c0 2.33-.17 4.43-.52 6.3s-1 3.51-1.98 4.96"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Moonrise(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.8 14.86c0-.98.19-1.92.58-2.82c.38-.9.9-1.67 1.55-2.32c.65-.65 1.43-1.17 2.32-1.56s1.84-.58 2.83-.58h1.17c.16.04.24.14.24.28l.04.9c.04 1.3.51 2.41 1.41 3.33s2 1.41 3.28 1.46l.85.07c.16 0 .23.08.23.23v1.01c.01 1.03-.19 2-.58 2.92h-2.05c.51-.74.83-1.59.97-2.53c-1.67-.35-3.02-1.07-4.03-2.16s-1.59-2.37-1.75-3.83c-.97.05-1.85.35-2.66.9c-.8.55-1.43 1.24-1.87 2.08c-.44.84-.66 1.72-.66 2.63c0 1.07.28 2.04.83 2.92H8.4c-.4-.94-.6-1.91-.6-2.93m.29 6.01c0-.29.09-.52.28-.68c.18-.18.41-.26.69-.26h2.63L14.8 17c.1-.08.22-.08.35 0l3.16 2.92h2.77c.27 0 .5.09.69.28a.9.9 0 0 1 .29.67c0 .27-.1.5-.29.69c-.19.19-.42.29-.69.29H17.7c-.1 0-.2-.02-.29-.07L15 19.51l-2.39 2.27c-.08.05-.17.07-.28.07H9.06c-.27 0-.5-.1-.69-.29s-.28-.42-.28-.69"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Moonset(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.74 14.86c0-.98.19-1.92.58-2.82c.38-.9.9-1.67 1.55-2.32c.65-.65 1.43-1.17 2.32-1.56c.9-.39 1.84-.58 2.83-.58h1.17c.16.04.24.14.24.28l.05.9c.02.64.15 1.25.4 1.83s.58 1.08 1 1.5c.42.43.91.77 1.48 1.03s1.17.4 1.8.43l.85.07c.16 0 .24.08.24.23v1.01c.01 1.01-.19 1.98-.59 2.92h-2.05c.51-.74.83-1.59.97-2.53c-1.68-.35-3.02-1.07-4.03-2.16s-1.59-2.37-1.75-3.83c-.97.05-1.85.35-2.66.9c-.8.55-1.42 1.24-1.87 2.08c-.44.84-.66 1.72-.66 2.63c0 1.07.28 2.04.83 2.92h-2.1c-.4-.94-.6-1.91-.6-2.93m.25 6.03c0-.26.1-.49.3-.69c.18-.18.41-.27.68-.27h3.22c.11 0 .2.02.28.08l2.35 2.22L17.21 20c.07-.05.17-.08.29-.08h3.3c.27 0 .5.09.69.28c.19.19.29.42.29.68c0 .27-.1.5-.29.69c-.19.19-.42.29-.69.29h-2.68l-3.13 2.84c-.12.09-.24.09-.34 0l-3.08-2.84h-2.6c-.27 0-.5-.1-.69-.29a.906.906 0 0 1-.29-.68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Na(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.87 18.23h1.94v-3.64h.02l2.05 3.64h1.99v-6.66h-1.94v3.55h-.02l-1.94-3.55h-2.1zm6.52.15h1.43l2.61-6.97h-1.42zm2.87-.15h2.07l.29-.95h2.12l.28.95h2.13l-2.43-6.66h-2.01zm2.81-2.39l.64-2.04h.03l.6 2.04z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightAltCloudy(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.14 16.9c0-1.16.35-2.18 1.06-3.08s1.62-1.47 2.74-1.72c.23-1.03.7-1.93 1.4-2.7a5.93 5.93 0 0 1 2.53-1.65c.62-.21 1.26-.32 1.93-.32c.81 0 1.6.16 2.35.48c.28-.47.61-.88.99-1.22c.38-.34.77-.61 1.17-.79c.4-.18.8-.32 1.18-.41s.76-.13 1.12-.13c.38 0 .79.05 1.23.16l.82.25c.14.06.18.13.14.22l-.14.6c-.07.31-.1.6-.1.86c0 .31.05.63.15.95c.1.32.24.63.44.94c.19.31.46.58.8.83s.72.44 1.15.57l.62.22c.1.03.15.08.15.16c0 .02-.01.04-.02.07l-.18.67c-.27 1.08-.78 1.93-1.5 2.57c.4.7.62 1.45.65 2.24c.01.05.01.12.01.23c0 .89-.22 1.72-.67 2.48c-.44.76-1.05 1.36-1.8 1.8c-.76.44-1.59.67-2.48.67H9.07c-.89 0-1.72-.22-2.48-.67s-1.35-1.05-1.79-1.8s-.66-1.58-.66-2.48m1.71 0c0 .89.32 1.66.96 2.31c.64.65 1.39.98 2.26.98h10.81c.89 0 1.65-.32 2.28-.97s.95-1.42.95-2.32c0-.88-.32-1.63-.96-2.26c-.64-.63-1.4-.95-2.28-.95h-1.78l-.1-.75c-.1-1.01-.52-1.88-1.26-2.59s-1.62-1.11-2.63-1.2c-.03 0-.08 0-.15-.01c-.07-.01-.11-.01-.15-.01c-.51 0-1.02.1-1.54.29V9.4c-.73.28-1.35.74-1.84 1.37c-.5.63-.8 1.35-.9 2.17l-.07.72l-.68.03c-.84.1-1.54.45-2.1 1.06s-.82 1.32-.82 2.15M17.6 8.79c1.06.91 1.72 1.97 1.97 3.18h.32c1.24 0 2.3.39 3.17 1.18c.33-.31.58-.67.76-1.07a4.95 4.95 0 0 1-2.16-1.97c-.52-.88-.79-1.81-.79-2.78v-.24c-.05-.01-.13-.01-.24-.01c-.58-.01-1.15.13-1.7.44c-.55.3-1 .72-1.33 1.27"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightAltCloudyGusts(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.98 21.02c0 .23.09.43.27.6c.17.17.37.25.61.25H9.6c.26 0 .49.1.69.29c.2.19.3.42.3.68c0 .27-.1.5-.3.69s-.43.29-.69.29c-.26 0-.48-.1-.68-.3a.888.888 0 0 0-.61-.24c-.24 0-.44.08-.6.24c-.16.16-.24.36-.24.6c0 .22.08.42.24.6c.52.53 1.16.79 1.89.79s1.37-.26 1.89-.78c.52-.52.78-1.15.78-1.89s-.26-1.37-.78-1.89s-1.15-.79-1.89-.79H3.86a.87.87 0 0 0-.62.26c-.17.17-.26.37-.26.6m0-3.02c0 .22.09.41.27.58c.17.16.38.24.61.24h10.85c.74 0 1.37-.26 1.89-.78c.52-.52.78-1.15.78-1.88c0-.74-.26-1.36-.78-1.88c-.52-.52-1.15-.77-1.89-.77c-.76 0-1.39.25-1.89.76c-.15.16-.23.36-.23.61s.08.45.23.6c.15.15.35.23.59.23s.44-.07.62-.23c.19-.19.42-.28.68-.28c.26 0 .48.09.67.28a.939.939 0 0 1 0 1.36c-.19.19-.42.29-.67.29H3.86a.87.87 0 0 0-.62.26c-.17.18-.26.38-.26.61m2.52-2.35c0 .09.05.13.16.13H7.1c.08 0 .15-.05.22-.15c.22-.54.57-.99 1.05-1.35c.47-.35 1-.55 1.58-.6l.52-.07c.12 0 .19-.06.19-.17l.08-.52c.11-1.08.57-1.99 1.38-2.71c.81-.73 1.77-1.09 2.86-1.09s2.05.36 2.85 1.09c.81.72 1.27 1.63 1.38 2.72l.07.58c0 .12.06.18.19.18h1.62c.91 0 1.68.32 2.32.95c.64.63.96 1.39.96 2.28c0 .89-.32 1.65-.96 2.29c-.64.64-1.41.96-2.31.96h-6.91c-.11 0-.17.06-.17.18v1.37c0 .11.06.17.17.17h6.91c.89 0 1.72-.22 2.48-.67s1.36-1.05 1.8-1.81s.67-1.59.67-2.48c0-.88-.23-1.71-.68-2.48c.73-.71 1.23-1.57 1.51-2.58l.12-.69a.09.09 0 0 0 .03-.07c0-.04-.05-.1-.14-.16l-.6-.21c-.84-.26-1.48-.71-1.92-1.36c-.44-.65-.66-1.32-.66-2.02c0-.24.03-.51.09-.79l.13-.62c.02-.1-.02-.18-.13-.22l-.8-.24c-.44-.11-.85-.16-1.25-.16c-.37 0-.74.04-1.12.13c-.38.09-.77.22-1.18.41c-.4.18-.8.45-1.18.8c-.38.35-.72.75-1 1.22c-.71-.3-1.48-.45-2.32-.45c-1.41 0-2.66.44-3.75 1.31s-1.77 1.99-2.07 3.35c-.86.2-1.61.61-2.27 1.23c-.66.62-1.11 1.35-1.36 2.2v.03c.01.02 0 .05 0 .09m13.23-6.89c.31-.55.74-.97 1.29-1.26c.55-.29 1.12-.44 1.71-.44c.14 0 .24 0 .31.01c-.01.09-.02.21-.02.36c0 .94.26 1.85.79 2.71a4.92 4.92 0 0 0 2.17 1.94c-.16.38-.41.72-.75 1.03c-.89-.76-1.94-1.14-3.16-1.14h-.33c-.26-1.26-.93-2.33-2.01-3.21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightAltCloudyHigh(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.57 13.43c0-1.15.36-2.18 1.08-3.08s1.63-1.48 2.73-1.74C7.7 7.24 8.4 6.12 9.5 5.24s2.35-1.31 3.76-1.31c1.38 0 2.61.43 3.69 1.28s1.78 1.95 2.1 3.29h.33c.9 0 1.73.22 2.49.65c.76.43 1.37 1.03 1.81 1.79c.44.76.67 1.58.67 2.48c0 1.15-.35 2.18-1.06 3.08c.64.55 1.4.84 2.26.87l.66.06c.12 0 .18.06.18.19v.77c.01 1.01-.24 1.95-.73 2.8c-.49.86-1.17 1.53-2.02 2.03c-.85.5-1.78.75-2.79.75c-.77 0-1.5-.15-2.19-.44c-.69-.29-1.28-.69-1.78-1.19c-.49-.5-.89-1.09-1.18-1.78c-.29-.69-.44-1.41-.44-2.17H8.37c-1.34-.06-2.47-.57-3.4-1.53s-1.4-2.1-1.4-3.43m1.71 0c0 .87.3 1.62.9 2.26s1.33.98 2.19 1.02h11.19c.86-.04 1.59-.38 2.19-1.02c.6-.64.9-1.39.9-2.26c0-.88-.32-1.63-.97-2.28c-.65-.64-1.42-.97-2.31-.97h-1.62c-.11 0-.17-.06-.17-.17l-.07-.58c-.11-1.08-.58-1.99-1.4-2.71s-1.77-1.09-2.86-1.09c-1.1 0-2.05.36-2.86 1.09S9.13 8.35 9.03 9.43l-.07.58c0 .11-.07.17-.2.17h-.53c-.84.1-1.54.46-2.1 1.07s-.85 1.34-.85 2.18m11.43 4.96c0 .79.2 1.52.6 2.17c.4.65.91 1.15 1.54 1.5a3.98 3.98 0 0 0 3.81.07c.6-.3 1.12-.75 1.58-1.36s.75-1.31.86-2.1c-1.08-.22-1.98-.65-2.72-1.3c-.84.65-1.78.99-2.82 1.01h-2.85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightAltCloudyWindy(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.35 21.05c0 .24.08.43.25.59c.17.16.38.23.63.23h9.4c.24 0 .43-.08.59-.23c.16-.16.23-.35.23-.59c0-.25-.08-.45-.23-.61a.791.791 0 0 0-.59-.24h-9.4c-.25 0-.46.08-.63.24s-.25.36-.25.61M4.98 18c0 .24.09.44.26.6c.16.17.36.25.6.25h9.42c.23 0 .43-.08.59-.25c.16-.17.24-.37.24-.6c0-.23-.08-.43-.23-.59s-.35-.24-.59-.24H5.85c-.24 0-.44.08-.61.24c-.17.16-.26.36-.26.59m1.04-2.34c0 .09.06.14.18.14h1.43c.09 0 .16-.05.22-.14c.23-.54.57-.99 1.04-1.35c.47-.36.99-.56 1.58-.6l.55-.07c.12 0 .18-.06.18-.17l.07-.52c.11-1.09.57-2 1.38-2.72c.82-.73 1.77-1.09 2.87-1.09c1.09 0 2.04.36 2.84 1.08c.8.72 1.27 1.62 1.41 2.71l.08.58c0 .11.06.17.18.17h1.61c.91 0 1.68.32 2.32.96c.64.64.96 1.41.96 2.31c0 .88-.32 1.65-.97 2.29c-.65.65-1.41.97-2.3.97h-6.91c-.11 0-.17.06-.17.17v1.34c0 .11.06.17.17.17h6.91c.9 0 1.73-.22 2.49-.66c.76-.44 1.37-1.04 1.81-1.8c.44-.76.67-1.59.67-2.49s-.22-1.72-.65-2.47c.72-.64 1.22-1.5 1.51-2.58l.18-.68c.01-.01.01-.03.01-.07c0-.08-.05-.13-.15-.16l-.62-.22c-.44-.13-.83-.32-1.16-.57c-.34-.25-.6-.53-.8-.84c-.2-.31-.34-.62-.44-.94c-.1-.32-.15-.63-.15-.95c0-.24.04-.53.11-.87l.13-.61c.04-.09 0-.17-.13-.23l-.84-.25c-.44-.11-.85-.16-1.25-.16c-.38 0-.75.04-1.13.13s-.77.22-1.18.41c-.41.18-.8.45-1.18.8c-.38.35-.71.75-.99 1.22c-.77-.32-1.57-.48-2.37-.48c-1.41 0-2.66.44-3.75 1.32s-1.78 2-2.08 3.38c-.87.2-1.63.61-2.28 1.22a4.9 4.9 0 0 0-1.36 2.21v.03c-.02.01-.03.04-.03.08m.73 8.49c0 .24.08.44.25.6c.16.17.36.25.6.25h9.43c.24 0 .44-.08.61-.25c.17-.17.25-.37.25-.6c0-.23-.08-.43-.25-.59a.853.853 0 0 0-.61-.24H7.6c-.24 0-.44.08-.6.24c-.17.16-.25.36-.25.59M19.33 8.78c.34-.55.79-.98 1.35-1.28c.55-.3 1.12-.45 1.7-.44c.11 0 .2 0 .25.01v.24c0 .97.26 1.9.79 2.79c.53.88 1.25 1.55 2.17 1.98c-.17.4-.43.76-.76 1.07c-.88-.79-1.95-1.18-3.2-1.18h-.32c-.25-1.2-.91-2.26-1.98-3.19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightAltHail(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.1 16.91c0 1.33.46 2.48 1.39 3.43s2.06 1.47 3.4 1.53c.12 0 .18-.06.18-.17v-1.34c0-.11-.06-.17-.18-.17c-.86-.04-1.58-.38-2.18-1.02s-.9-1.39-.9-2.26c0-.83.28-1.55.84-2.17c.56-.61 1.26-.97 2.1-1.07l.53-.03c.13 0 .2-.06.2-.18l.07-.54c.11-1.08.56-1.99 1.37-2.72c.81-.73 1.76-1.1 2.85-1.1c1.08 0 2.03.37 2.85 1.1c.82.73 1.28 1.64 1.4 2.72l.08.58c0 .11.06.17.17.17h1.61c.89 0 1.66.32 2.31.96c.65.64.98 1.4.98 2.27c0 .87-.3 1.62-.9 2.26c-.6.64-1.33.98-2.18 1.02c-.13 0-.2.06-.2.17v1.34c0 .11.07.17.2.17c1.33-.04 2.46-.55 3.38-1.51c.93-.96 1.39-2.11 1.39-3.45c0-.86-.22-1.66-.65-2.41c.79-.74 1.3-1.62 1.55-2.62l.13-.68c.02-.01.03-.03.03-.07c0-.07-.05-.13-.16-.16l-.56-.17c-.57-.17-1.05-.45-1.46-.85c-.4-.4-.69-.81-.86-1.25c-.17-.43-.25-.87-.25-1.32c-.01-.24.02-.51.08-.79l.14-.58c.03-.09-.02-.16-.14-.22l-.8-.25c-.42-.12-.86-.19-1.31-.19c-.35 0-.71.04-1.08.13s-.76.22-1.17.4c-.41.18-.8.45-1.19.8c-.38.35-.72.75-1 1.22c-.75-.32-1.54-.49-2.37-.49c-1.41 0-2.67.44-3.76 1.31s-1.79 1.99-2.1 3.36c-1.11.26-2.02.83-2.74 1.73S4.1 15.76 4.1 16.91m5.48 7.03c.09.21.24.36.46.45c.19.1.4.11.62.02c.22-.08.37-.23.45-.45c.1-.22.11-.43.02-.65a.728.728 0 0 0-.45-.44a.71.71 0 0 0-.62-.02a.88.88 0 0 0-.47.46c-.09.17-.1.38-.01.63m.62-2.83c0 .15.05.3.16.45s.26.26.46.32c.26.1.48.1.67 0c.19-.1.32-.29.4-.57l.88-3.21c.07-.25.04-.47-.08-.67c-.12-.2-.3-.32-.54-.37a.737.737 0 0 0-.63.07c-.2.11-.33.28-.4.51l-.88 3.22c0 .02-.01.06-.02.12c-.01.05-.02.1-.02.13m1.87 5.6c0 .12.02.22.06.29c.09.22.24.37.45.45c.09.05.2.07.33.07c.06 0 .16-.02.3-.06c.23-.08.39-.23.48-.45c.1-.22.1-.44 0-.66a.88.88 0 0 0-.45-.46c-.2-.09-.4-.09-.61 0c-.19.08-.33.2-.42.36c-.1.17-.14.32-.14.46m.74-2.65c0 .38.21.64.64.78c.09.03.17.05.23.05c.11 0 .23-.03.35-.08c.23-.08.39-.27.47-.57l1.65-6.12c.06-.24.04-.45-.07-.65a.856.856 0 0 0-.5-.39a.777.777 0 0 0-.65.07c-.2.11-.34.28-.4.51l-1.68 6.17a.74.74 0 0 0-.04.23m3.44-.42c0 .13.02.23.07.31c.08.2.23.35.44.44c.12.05.23.08.35.08c.06 0 .16-.02.3-.06c.22-.09.37-.23.45-.44c.08-.22.08-.43 0-.63a.839.839 0 0 0-.42-.45a.779.779 0 0 0-.65-.02c-.22.08-.37.24-.47.46a.57.57 0 0 0-.07.31m.72-2.56c0 .16.05.32.15.46c.1.14.25.25.45.31c.17.02.26.03.27.03c.41 0 .66-.2.77-.61l.87-3.17c.06-.24.04-.45-.07-.65a.856.856 0 0 0-.5-.39a.751.751 0 0 0-.64.07c-.2.11-.33.28-.4.51L17 20.81c-.02.09-.03.18-.03.27m.65-12.25c.31-.57.75-1.01 1.3-1.31c.55-.3 1.14-.45 1.76-.44c.11 0 .2.01.25.02v.31c0 .98.26 1.89.78 2.75c.52.86 1.25 1.51 2.17 1.95c-.19.44-.44.79-.75 1.07c-.88-.79-1.96-1.18-3.25-1.18h-.32c-.26-1.25-.9-2.31-1.94-3.17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightAltLightning(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.12 16.91c0 1.33.46 2.48 1.39 3.43s2.06 1.47 3.4 1.53c.11 0 .17-.06.17-.17v-1.37c0-.12-.06-.18-.17-.18c-.87-.07-1.6-.41-2.19-1.04c-.59-.62-.89-1.36-.89-2.21c0-.84.28-1.57.85-2.19c.57-.62 1.26-.97 2.1-1.04l.53-.07c.13 0 .2-.06.2-.17l.07-.52c.11-1.08.56-1.99 1.37-2.71c.81-.73 1.76-1.09 2.85-1.09s2.04.36 2.85 1.09c.81.72 1.27 1.63 1.39 2.72l.07.58c0 .12.06.18.18.18h1.61c.91 0 1.68.32 2.32.95c.64.63.96 1.39.96 2.28c0 .85-.3 1.59-.89 2.21c-.59.62-1.32.97-2.19 1.04c-.13 0-.2.06-.2.18v1.37c0 .11.07.17.2.17c1.33-.04 2.46-.55 3.38-1.51s1.38-2.11 1.38-3.45c0-.89-.23-1.72-.68-2.48c.8-.72 1.32-1.58 1.55-2.58l.15-.72c.01-.01.01-.03.01-.07c0-.07-.05-.13-.16-.16l-.58-.17c-.57-.16-1.05-.44-1.45-.82c-.4-.39-.68-.8-.85-1.23a3.555 3.555 0 0 1-.16-2.11l.14-.62c.03-.09-.02-.17-.14-.22l-.79-.24a4.944 4.944 0 0 0-2.37-.03c-.38.09-.78.22-1.19.41c-.41.18-.81.45-1.2.8c-.39.35-.72.75-1 1.22c-.71-.3-1.48-.45-2.33-.45c-1.41 0-2.66.44-3.75 1.31a5.83 5.83 0 0 0-2.1 3.35c-1.1.26-2.01.84-2.73 1.74c-.75.89-1.11 1.91-1.11 3.06m7.67 4.65c-.05.14 0 .22.14.22h2.16l-1.31 4.14h.3l4.17-5.59c.04-.04.05-.09.03-.14c-.02-.05-.06-.07-.13-.07h-2.2l2.31-4.21c.07-.14.02-.22-.14-.22h-2.94c-.08 0-.15.05-.22.14zM17.6 8.81c.33-.57.77-1 1.33-1.3c.55-.3 1.13-.45 1.72-.45c.13 0 .22.01.27.02v.31c0 .96.26 1.87.78 2.73a5.03 5.03 0 0 0 2.17 1.96c-.16.37-.41.73-.75 1.07c-.92-.79-1.99-1.18-3.22-1.18h-.32c-.29-1.26-.95-2.31-1.98-3.16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightAltPartlyCloudy(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.77 19.65c0-.79.23-1.48.68-2.09c.45-.61 1.06-1.03 1.81-1.27c.32-1.09.98-1.92 1.99-2.49v-.35c0-1.46.46-2.74 1.38-3.85s2.09-1.8 3.5-2.06c.36-.06.72-.09 1.08-.09h.03c.21 0 .44.02.7.05c.26.02.5.06.73.11l.91.28c.13.07.18.16.16.26l-.13.7c-.07.33-.11.65-.11.97c0 .35.05.71.16 1.07c.11.37.27.72.5 1.08s.52.68.91.97c.38.29.83.51 1.33.66l.71.21c.11.03.17.08.17.18c0 .04 0 .06-.01.07l-.18.68c-.06.24-.13.49-.22.73c-.15.44-.38.89-.7 1.37c0-.01 0-.01-.01-.01c-.44.63-.98 1.16-1.64 1.58c.14.34.21.75.21 1.24c0 .99-.35 1.83-1.04 2.53c-.69.7-1.52 1.05-2.49 1.05h-6.85c-.97 0-1.81-.35-2.52-1.06c-.7-.71-1.06-1.55-1.06-2.52m1.98 0c0 .45.15.83.46 1.15s.69.47 1.14.47h6.85c.43 0 .8-.16 1.12-.48c.32-.32.47-.7.47-1.14c0-.43-.16-.8-.47-1.12s-.69-.47-1.12-.47h-1.3c-.11 0-.19-.07-.25-.23l-.08-.64c-.07-.58-.32-1.06-.75-1.44s-.93-.58-1.51-.58c-.57 0-1.06.19-1.48.58c-.42.39-.66.87-.73 1.44l-.1.55c0 .15-.06.22-.19.22l-.63.08c-.41.04-.75.21-1.02.51c-.27.32-.41.68-.41 1.1m4.43-6.4h.12c.93 0 1.75.26 2.49.78c.73.52 1.25 1.22 1.54 2.1c.77 0 1.45.24 2.03.72c.69-.43 1.2-1.02 1.53-1.75c-1.04-.52-1.85-1.27-2.43-2.25s-.88-2.01-.88-3.11v-.35h-.24c-.61 0-1.2.13-1.77.39c-.57.26-1.05.64-1.44 1.12l-.03-.02c-.55.68-.85 1.47-.92 2.37"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightAltRain(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.07 16.9c0 1.33.47 2.48 1.4 3.44s2.07 1.47 3.4 1.53c.12 0 .18-.06.18-.17v-1.34c0-.11-.06-.17-.18-.17a3.15 3.15 0 0 1-2.19-1.03c-.6-.64-.9-1.39-.9-2.26c0-.83.28-1.55.85-2.17c.57-.62 1.27-.97 2.1-1.07l.53-.04c.13 0 .2-.06.2-.17l.07-.54c.11-1.08.57-1.99 1.38-2.72a4.15 4.15 0 0 1 2.86-1.1c1.09 0 2.04.37 2.86 1.1c.82.73 1.28 1.64 1.4 2.72l.08.57c0 .12.06.18.17.18h1.62c.89 0 1.67.32 2.32.96c.65.64.98 1.4.98 2.28c0 .87-.3 1.62-.9 2.26c-.6.64-1.33.98-2.19 1.03c-.14 0-.21.06-.21.17v1.34c0 .11.07.17.21.17c1.33-.04 2.46-.55 3.38-1.51c.93-.97 1.39-2.12 1.39-3.45c0-.88-.23-1.7-.68-2.46c.81-.73 1.33-1.6 1.58-2.62l.14-.72c.01-.01.02-.03.02-.07c0-.07-.05-.13-.16-.16l-.56-.18c-.57-.16-1.06-.44-1.46-.83c-.41-.39-.7-.8-.87-1.23c-.17-.43-.26-.86-.26-1.28c-.02-.22.01-.5.08-.82l.14-.61c.04-.1 0-.18-.14-.24l-.79-.24c-.45-.1-.87-.15-1.27-.15c-.38 0-.76.04-1.14.13c-.39.09-.79.22-1.2.41c-.41.18-.81.45-1.2.8c-.39.35-.72.75-1 1.22c-.82-.3-1.6-.45-2.33-.45c-1.41 0-2.67.44-3.76 1.32s-1.8 2-2.11 3.37c-1.11.26-2.02.83-2.74 1.73c-.74.89-1.1 1.92-1.1 3.07m5.56 6.84c0 .17.05.33.16.49c.11.16.27.27.49.33c.23.07.45.05.64-.04c.2-.1.33-.28.4-.56l1.43-5.87c.06-.25.03-.48-.08-.67a.765.765 0 0 0-.52-.37a.737.737 0 0 0-.63.07c-.2.11-.34.28-.41.51l-1.44 5.9c0 .01-.01.04-.02.09c-.01.05-.02.09-.02.12m2.61 3.07c0 .16.05.31.15.46c.1.15.25.25.45.31c.11.03.19.04.24.04c.44 0 .71-.2.82-.59l2.25-8.93c.06-.24.04-.46-.07-.65a.856.856 0 0 0-.5-.39c-.23-.07-.45-.05-.66.07s-.34.28-.39.5l-2.26 8.92c0 .01 0 .05-.01.12c-.02.06-.02.11-.02.14m4.16-2.99c0 .36.21.6.63.74c.14.04.24.06.3.06c.11 0 .23-.02.35-.07c.21-.08.34-.28.39-.58l1.43-5.87c.06-.24.04-.45-.08-.65a.848.848 0 0 0-.51-.39a.804.804 0 0 0-.66.07c-.21.11-.33.28-.38.51l-1.43 5.9c-.02.16-.04.26-.04.28m1.18-15.05c.32-.58.75-1.02 1.31-1.33c.55-.3 1.14-.45 1.76-.44c.12 0 .21 0 .27.01v.3c-.01.97.24 1.88.77 2.75c.52.86 1.26 1.52 2.21 1.97c-.22.46-.49.81-.79 1.07c-.92-.76-1.99-1.13-3.23-1.13h-.31c-.27-1.27-.93-2.33-1.99-3.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightAltRainMix(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.07 16.93c0 1.33.47 2.47 1.4 3.43s2.07 1.47 3.4 1.51c.12 0 .18-.06.18-.17v-1.34c0-.11-.06-.17-.18-.17c-.85-.04-1.58-.39-2.18-1.03s-.91-1.39-.91-2.23c0-.85.28-1.59.85-2.21c.57-.62 1.27-.97 2.1-1.04l.53-.07c.13 0 .2-.06.2-.18l.07-.51c.11-1.1.56-2.02 1.37-2.75s1.76-1.1 2.86-1.1c1.09 0 2.04.37 2.86 1.1c.82.73 1.29 1.64 1.4 2.72l.07.57c0 .12.06.19.17.19h1.62c.91 0 1.68.32 2.33.97c.65.64.97 1.41.97 2.31c0 .55-.14 1.07-.41 1.56s-.65.89-1.13 1.2c-.48.31-1 .48-1.56.51c-.13 0-.2.06-.2.17v1.34c0 .11.07.17.2.17c.88-.02 1.69-.26 2.42-.71c.73-.45 1.31-1.05 1.73-1.8c.42-.75.63-1.56.63-2.43c0-.88-.23-1.72-.68-2.51c.83-.74 1.36-1.62 1.58-2.62l.14-.68a.09.09 0 0 0 .03-.07c0-.06-.06-.11-.17-.16l-.55-.18c-.57-.17-1.07-.45-1.47-.85c-.41-.4-.7-.81-.87-1.25c-.17-.43-.26-.86-.26-1.29c-.02-.21.01-.49.09-.82l.13-.58c.04-.1 0-.18-.13-.23l-.8-.24a4.958 4.958 0 0 0-2.41-.04c-.38.09-.78.22-1.19.41s-.81.46-1.2.81c-.39.35-.72.76-1 1.23c-.81-.31-1.6-.46-2.35-.46c-1.41 0-2.67.44-3.76 1.32s-1.8 2-2.11 3.37c-1.12.29-2.04.88-2.75 1.77c-.71.87-1.06 1.9-1.06 3.06m5.41 7.05c0 .17.05.34.16.51c.11.17.27.28.47.35c.23.07.44.06.64-.04s.32-.28.39-.56l.14-.61c.05-.23.02-.44-.09-.63a.82.82 0 0 0-.52-.4c-.23-.07-.44-.04-.64.08s-.34.3-.4.53l-.13.58c-.02.04-.02.1-.02.19m.76-2.9c0 .21.08.4.25.57c.16.17.34.25.56.25c.24 0 .44-.08.6-.24c.16-.16.24-.35.24-.59c0-.23-.08-.43-.24-.59a.814.814 0 0 0-.6-.24c-.23 0-.43.08-.58.23s-.23.38-.23.61m.61-2.27c-.01.16.03.31.14.45c.1.15.26.25.48.32a.8.8 0 0 0 .62-.07c.21-.11.34-.28.41-.51l.28-.9c.07-.24.05-.46-.07-.65a.913.913 0 0 0-.54-.39a.737.737 0 0 0-.63.07a.85.85 0 0 0-.41.5l-.24.92c0 .02-.01.06-.02.12c-.02.05-.02.1-.02.14m1.16 8.29c0 .18.05.34.15.5c.1.16.26.27.48.33c.08.02.17.03.25.03c.43 0 .69-.2.79-.61l.14-.59c.06-.26.03-.48-.08-.68s-.29-.32-.52-.37a.744.744 0 0 0-.63.07c-.21.12-.34.29-.41.51l-.14.59c-.02.09-.03.16-.03.22m.78-2.9c0 .22.08.41.25.58c.16.16.35.24.57.24c.24 0 .43-.08.59-.23c.16-.16.23-.35.23-.59c0-.23-.08-.42-.23-.58s-.35-.23-.59-.23s-.43.08-.59.23a.8.8 0 0 0-.23.58m.63-2.27c-.01.15.03.31.13.47s.25.26.45.3c.23.06.44.04.64-.06c.19-.1.33-.29.41-.56l.27-.9c.07-.22.05-.43-.07-.63a.867.867 0 0 0-.53-.4c-.22-.07-.43-.04-.64.08s-.34.3-.41.53l-.22.9c-.02.08-.03.17-.03.27m2.75 2.15c0 .16.05.32.15.48s.26.27.46.33c.03 0 .08.01.14.02c.06.01.1.02.14.02c.41 0 .66-.22.76-.66l.14-.6c.07-.21.05-.42-.07-.63a.809.809 0 0 0-.51-.41c-.25-.06-.48-.04-.68.08s-.34.29-.41.53l-.09.59c0 .02-.01.07-.02.12c0 .05-.01.09-.01.13m.74-2.96c0 .22.08.42.25.57c.15.16.34.24.57.24c.24 0 .43-.08.59-.23s.23-.35.23-.58c0-.24-.08-.43-.23-.59s-.35-.23-.59-.23s-.43.08-.59.23s-.23.35-.23.59m.61-2.31c0 .17.05.33.16.48s.27.26.49.32c.02 0 .06.01.12.02s.11.02.14.02c.1 0 .22-.03.36-.09c.21-.11.35-.29.41-.52l.24-.9c.06-.23.04-.44-.08-.63a.827.827 0 0 0-.51-.4c-.23-.07-.44-.05-.64.06s-.32.27-.39.51l-.28.91c0 .02-.01.06-.02.12c.01.03 0 .07 0 .1m.07-10.05c.32-.58.76-1.02 1.31-1.34c.56-.32 1.13-.47 1.73-.46c.09 0 .19.01.3.03v.31c-.01.98.25 1.9.77 2.76c.53.86 1.27 1.5 2.22 1.94c-.19.41-.46.78-.8 1.11c-.92-.76-2-1.14-3.23-1.14h-.31c-.31-1.31-.97-2.38-1.99-3.21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightAltRainWind(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.06 16.93c0 1.12.33 2.12 1 3c.67.88 1.52 1.47 2.57 1.77c.09.02.17-.01.24-.08L9 20.22c-.88 0-1.63-.32-2.27-.97c-.64-.65-.96-1.42-.96-2.32c0-.84.28-1.56.84-2.17s1.27-.95 2.11-1.03l.5-.07c.12 0 .19-.06.19-.19l.08-.53c.12-1.09.59-2 1.41-2.73a4.15 4.15 0 0 1 2.86-1.1c1.09 0 2.04.37 2.86 1.1s1.29 1.64 1.41 2.72l.07.58c0 .11.06.17.18.17h1.62c.88 0 1.64.32 2.28.96s.96 1.4.96 2.28c0 .85-.28 1.59-.84 2.22s-1.25.98-2.07 1.05c-.45.06-.74.15-.86.28l-2.33 2.91c-.16.17-.22.38-.19.63c.02.24.13.43.31.59c.18.16.37.23.57.23c.23 0 .44-.12.64-.38l2.04-2.59c.62-.06 1.2-.24 1.76-.52c.55-.28 1.03-.65 1.42-1.08c.39-.44.71-.95.94-1.53c.23-.58.35-1.18.35-1.81c0-.87-.23-1.68-.68-2.44c.81-.74 1.34-1.61 1.58-2.62v-.09l.2-.77l-.76-.26c-.57-.17-1.06-.45-1.47-.83s-.69-.8-.86-1.23c-.17-.43-.26-.87-.26-1.31c0-.26.03-.52.08-.8l.19-.78l-.83-.23c-.01 0-.02 0-.03-.01s-.02-.02-.04-.02s-.03-.01-.04-.02c-.05-.01-.06-.02-.06-.02c-.44-.11-.85-.16-1.25-.16c-.38.01-.76.05-1.15.14s-.78.22-1.2.41c-.42.19-.82.46-1.2.81s-.72.76-1 1.24c-.75-.33-1.53-.49-2.34-.49c-1.41 0-2.67.44-3.76 1.31s-1.8 1.99-2.11 3.36c-1.13.27-2.05.86-2.76 1.75s-1.07 1.91-1.07 3.07m3.71 7.99c0 .13.02.23.07.31c.09.22.23.37.43.46c.22.1.44.11.67.03a.72.72 0 0 0 .46-.44c.1-.22.1-.44.01-.67c-.09-.23-.24-.38-.45-.45a.8.8 0 0 0-.66-.02c-.22.08-.37.24-.45.45c-.06.08-.08.2-.08.33m1.84-2.45v.11c.02.23.13.41.33.55c.13.15.31.22.54.22c.23-.01.45-.11.66-.32l2.33-2.92c.14-.17.19-.38.17-.62a.847.847 0 0 0-.3-.58a.859.859 0 0 0-.63-.18c-.24.02-.43.14-.57.34l-2.32 2.86c-.14.16-.21.34-.21.54m.58 5.21c.09.21.24.36.46.45a.8.8 0 0 0 .33.08c.06 0 .16-.02.3-.06c.21-.09.36-.23.44-.44c.08-.22.08-.43.01-.65a.826.826 0 0 0-.44-.48c-.22-.08-.43-.08-.63 0s-.35.23-.45.44c-.11.2-.12.41-.02.66m1.59-2.66v.08c.02.22.13.42.32.58c.19.16.38.24.56.24c.22 0 .42-.11.6-.34l4.31-5.36a.84.84 0 0 0 .19-.62a.817.817 0 0 0-.29-.58c-.2-.14-.42-.2-.66-.18a.77.77 0 0 0-.57.3l-4.27 5.36c-.13.15-.19.33-.19.52m3.51 1.11c0 .11.02.22.07.33c.08.23.24.38.47.47c.23.09.43.09.61.02c.22-.09.37-.24.46-.46c.1-.22.11-.43.03-.64a.813.813 0 0 0-.45-.46a.91.91 0 0 0-.65 0c-.22.08-.37.22-.47.42c-.05.11-.07.22-.07.32m2.28-17.32c.31-.57.75-1.01 1.3-1.32c.55-.3 1.14-.45 1.76-.44c.12 0 .21 0 .26.01v.3c0 .97.27 1.89.8 2.75c.53.87 1.26 1.52 2.19 1.96c-.25.47-.51.84-.79 1.12c-.89-.79-1.96-1.18-3.22-1.18h-.32c-.29-1.27-.95-2.34-1.98-3.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightAltShowers(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.07 16.91c0 1.33.46 2.48 1.39 3.43s2.06 1.47 3.4 1.53c.11 0 .17-.06.17-.17v-1.37c0-.12-.06-.18-.17-.18c-.87-.07-1.6-.41-2.19-1.04c-.59-.62-.89-1.36-.89-2.21c0-.83.28-1.54.84-2.16s1.26-.97 2.1-1.07l.53-.07c.13 0 .2-.06.2-.17l.07-.52c.11-1.08.56-1.99 1.37-2.71c.81-.73 1.76-1.09 2.85-1.09s2.04.36 2.85 1.09c.81.72 1.28 1.63 1.39 2.72l.08.58c0 .12.06.18.18.18h1.61c.9 0 1.67.32 2.32.95c.64.63.97 1.39.97 2.28c0 .85-.3 1.59-.89 2.21c-.59.62-1.33.97-2.19 1.04c-.13 0-.2.06-.2.18v1.37c0 .11.07.17.2.17c1.33-.04 2.46-.55 3.38-1.51c.92-.96 1.38-2.11 1.38-3.45c0-.87-.22-1.68-.65-2.43c.81-.73 1.34-1.6 1.58-2.62v-.13l.19-.79l-.76-.21c-.81-.24-1.44-.7-1.89-1.35c-.45-.66-.67-1.34-.67-2.03c0-.26.03-.52.08-.78l.2-.8l-.85-.25l-.15-.04a6.72 6.72 0 0 0-1.25-.14c-.38 0-.76.04-1.14.13c-.39.09-.79.22-1.2.41c-.42.19-.82.45-1.2.8c-.38.35-.72.76-1 1.23c-.74-.33-1.53-.49-2.36-.49c-1.41 0-2.66.44-3.75 1.31s-1.77 1.99-2.07 3.36c-1.12.26-2.05.83-2.77 1.72c-.73.91-1.09 1.94-1.09 3.09m5.4 6.77c0 .15.05.3.15.45c.1.15.25.26.45.33c.22.07.43.06.64-.05s.34-.28.41-.51l.28-1.06c.07-.21.05-.41-.07-.62a.863.863 0 0 0-.51-.41c-.23-.06-.45-.03-.65.08s-.34.3-.42.53l-.23.99c-.03.16-.05.25-.05.27m1.3-4.73c0 .11.03.23.1.36c.07.17.25.3.53.38c.24.06.46.04.66-.06c.19-.1.33-.28.4-.52l.28-1.03c.07-.23.05-.45-.07-.64a.831.831 0 0 0-.51-.39a.905.905 0 0 0-.67.07a.82.82 0 0 0-.4.52l-.27 1.01c-.03.13-.05.23-.05.3m1.25 7.85c0 .17.05.33.15.49c.1.16.25.27.45.33c.11.03.18.05.23.05c.09 0 .21-.03.38-.1c.2-.08.34-.27.43-.55l.3-1.05c.07-.21.05-.42-.07-.63a.861.861 0 0 0-1.18-.33c-.2.12-.34.29-.41.53l-.25 1.01c-.02.08-.03.17-.03.25m1.33-4.77a.85.85 0 0 0 .61.78c.22.07.44.05.64-.06s.33-.28.4-.52l.27-1.04c.07-.21.05-.42-.06-.62a.831.831 0 0 0-.49-.41c-.24-.06-.47-.03-.68.08s-.35.3-.42.53l-.24 1zm2.81 1.76c0 .38.21.62.64.75c.09.02.17.03.24.03c.15 0 .27-.02.37-.07c.21-.08.36-.27.44-.57l.27-1.02c.06-.25.04-.47-.08-.67s-.29-.32-.53-.37c-.23-.07-.45-.05-.64.07s-.33.29-.4.51l-.27 1.04c0 .02-.01.07-.02.15c-.02.07-.02.12-.02.15m1.39-4.81a.822.822 0 0 0 .61.78c.14.03.22.05.23.05c.09 0 .21-.03.38-.1c.21-.08.35-.27.44-.55l.28-1.04c.06-.22.03-.43-.08-.63s-.3-.33-.53-.4c-.22-.07-.43-.05-.63.07s-.33.29-.4.52l-.26 1.06c-.03.09-.04.18-.04.24m.03-10.17c.32-.56.76-1 1.33-1.31a3.73 3.73 0 0 1 1.81-.47h.21c-.01.09-.01.21-.01.38c0 .95.26 1.85.78 2.71c.52.86 1.25 1.51 2.17 1.96c-.22.43-.48.8-.78 1.1c-.93-.81-2.02-1.21-3.25-1.21h-.32c-.25-1.19-.89-2.24-1.94-3.16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightAltSleet(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.09 17.13v-.37c.04-1.12.42-2.11 1.13-2.97c.71-.86 1.61-1.42 2.68-1.67c.31-1.36 1.02-2.48 2.11-3.36s2.35-1.31 3.76-1.31c.74 0 1.53.15 2.38.46c.28-.46.61-.86.99-1.2c.38-.34.77-.6 1.18-.79c.41-.18.8-.32 1.19-.4c.38-.08.76-.12 1.13-.12c.39 0 .8.04 1.24.13l.8.25c.12.06.17.13.14.23l-.12.62c-.07.29-.1.57-.1.84c0 .31.05.62.15.95c.1.32.24.64.44.95s.46.59.8.85s.72.45 1.16.58l.62.17c.1.03.15.08.15.16c0 .04 0 .06-.01.07l-.19.68a5.274 5.274 0 0 1-1.51 2.62c.44.76.66 1.58.66 2.45c0 1.34-.46 2.49-1.39 3.45c-.93.97-2.06 1.47-3.39 1.51c-.12 0-.19-.06-.19-.19v-1.36c0-.12.06-.18.19-.18c.87-.06 1.6-.41 2.2-1.03c.6-.62.9-1.36.9-2.21c0-.89-.32-1.66-.97-2.29s-1.42-.95-2.33-.95h-1.62c-.11 0-.17-.06-.17-.18l-.07-.58c-.11-1.07-.58-1.98-1.4-2.71s-1.77-1.1-2.86-1.1c-1.09 0-2.05.37-2.85 1.1c-.81.73-1.27 1.64-1.37 2.72l-.07.52c0 .11-.07.17-.2.17l-.53.07c-.84.1-1.54.46-2.1 1.07c-.57.62-.85 1.34-.85 2.17v.02l.01.01h.02c0 .52.12 1.02.36 1.48c.24.46.56.84.97 1.14h.01l.01.01l.01.01c.22.16.42.27.62.35c.34.15.69.23 1.07.24c.11 0 .17.06.17.17v1.34c0 .12-.06.18-.17.18c-.93-.04-1.78-.32-2.55-.82h-.01c-.66-.43-1.19-.99-1.59-1.69c-.4-.7-.61-1.45-.64-2.26m5.5 6.97c0-.03.01-.07.02-.13s.02-.09.02-.12l.09-.59c.07-.24.2-.41.41-.53s.43-.14.68-.08c.23.07.39.21.51.41c.11.21.13.42.07.63l-.14.6c-.1.43-.35.65-.76.65c-.03 0-.08 0-.15-.01c-.07-.01-.11-.01-.13-.01a.761.761 0 0 1-.46-.33a.842.842 0 0 1-.16-.49m.74-2.94c0-.24.08-.43.23-.59s.35-.23.59-.23s.43.08.59.23s.23.35.23.59c0 .23-.08.43-.23.58s-.35.23-.59.23c-.23 0-.42-.08-.57-.25a.75.75 0 0 1-.25-.56m1.64 6c0-.04.01-.11.04-.22l.13-.59c.07-.23.21-.4.41-.51c.21-.12.42-.14.63-.07c.23.04.41.17.53.37c.12.2.15.43.08.68l-.13.59c-.1.41-.37.61-.8.61c-.07 0-.16-.01-.24-.03a.831.831 0 0 1-.49-.33c-.1-.17-.16-.33-.16-.5m.79-2.91a.784.784 0 0 1 .82-.81c.24 0 .43.08.59.23c.16.16.23.35.23.58c0 .24-.08.43-.23.59c-.16.16-.35.23-.59.23c-.23 0-.42-.08-.58-.24a.785.785 0 0 1-.24-.58m.62-2.27c0-.09.01-.18.03-.27l.23-.9c.07-.23.21-.41.41-.53c.21-.12.42-.15.64-.08c.24.07.41.2.53.4c.12.2.14.41.07.63l-.26.9c-.08.28-.22.46-.41.56c-.19.1-.41.12-.64.06a.657.657 0 0 1-.45-.3a.69.69 0 0 1-.15-.47m2.76 2.14c0-.03 0-.08.01-.13s.01-.09.01-.11l.09-.59c.07-.24.2-.41.41-.53s.43-.14.68-.08c.23.07.4.21.51.41c.12.21.14.42.07.63l-.14.6c-.1.44-.35.66-.76.66c-.03 0-.08-.01-.14-.02c-.06-.01-.11-.02-.14-.02c-.2-.06-.35-.17-.45-.33s-.15-.32-.15-.49m.74-2.94c0-.24.08-.43.23-.59s.35-.23.59-.23s.43.08.59.23c.16.16.23.35.23.59a.784.784 0 0 1-.82.81c-.24 0-.43-.08-.58-.24a.785.785 0 0 1-.24-.57m.71-12.36c1.04.85 1.7 1.9 1.98 3.16h.33c1.23 0 2.3.39 3.22 1.18c.34-.31.59-.65.76-1.04c-.62-.3-1.15-.7-1.61-1.21c-.45-.51-.79-1.06-1.02-1.66c-.23-.6-.34-1.22-.34-1.86v-.3h-.22c-.62 0-1.21.15-1.77.45c-.56.3-1 .73-1.33 1.28"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightAltSleetStorm(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.09 16.89c0 1.11.33 2.1.99 2.97c.66.87 1.52 1.47 2.58 1.79l-.65 1.7c-.04.14 0 .21.14.21h2.12l-1.29 4.18h.28l4.23-5.62c.04-.04.04-.09.02-.14c-.03-.05-.07-.07-.14-.07h-2.18l2.47-4.64c.07-.14.03-.22-.13-.22H9.57c-.09 0-.16.05-.22.15l-1.07 2.88c-.71-.18-1.3-.57-1.78-1.17s-.71-1.27-.71-2.01c0-.83.28-1.55.85-2.17c.57-.61 1.27-.97 2.1-1.07l.53-.07c.13 0 .2-.06.2-.18l.07-.51c.11-1.08.56-1.99 1.37-2.72c.81-.73 1.76-1.1 2.85-1.1s2.04.37 2.86 1.1c.82.73 1.28 1.64 1.4 2.71l.07.57c0 .12.06.19.17.19h1.62c.91 0 1.68.32 2.33.95s.97 1.4.97 2.28c0 .85-.3 1.59-.9 2.21c-.6.62-1.33.97-2.2 1.03c-.12 0-.19.06-.19.19v1.36c0 .12.06.18.19.18c1.33-.04 2.46-.55 3.39-1.51c.93-.97 1.39-2.12 1.39-3.45c0-.87-.22-1.68-.66-2.45c.76-.74 1.27-1.61 1.51-2.62l.19-.68c.01-.01.01-.03.01-.07c0-.08-.05-.13-.15-.16l-.62-.17c-.57-.17-1.06-.45-1.46-.84c-.4-.39-.68-.8-.85-1.22s-.25-.84-.24-1.26c0-.28.03-.56.1-.85l.11-.61c.02-.1-.02-.18-.14-.23l-.8-.24c-.47-.09-.88-.14-1.24-.14c-.37-.01-.75.03-1.13.12c-.38.08-.78.22-1.19.4c-.41.18-.8.45-1.18.79c-.38.34-.71.74-.99 1.2c-.83-.31-1.62-.46-2.36-.46c-1.41 0-2.67.44-3.76 1.32s-1.8 2-2.11 3.36c-1.11.26-2.02.84-2.74 1.74c-.71.89-1.07 1.92-1.07 3.07m7.88 10.21c0 .17.05.33.16.49c.11.16.27.27.49.33c.09.02.17.03.24.03c.43 0 .7-.2.8-.61l.13-.59a.92.92 0 0 0-.08-.68a.762.762 0 0 0-.53-.37a.744.744 0 0 0-.63.07c-.21.12-.34.29-.41.51l-.13.59c-.03.12-.04.2-.04.23m.79-2.9c0 .23.08.42.24.58c.16.16.36.24.58.24c.24 0 .43-.08.59-.23c.16-.16.23-.35.23-.59a.784.784 0 0 0-.82-.81c-.24 0-.43.08-.59.23a.8.8 0 0 0-.23.58m.62-2.27c-.01.15.03.31.14.47c.1.16.25.26.45.3c.23.06.44.04.64-.06s.33-.29.41-.56l.26-.9c.07-.22.05-.43-.07-.63a.867.867 0 0 0-.53-.4c-.22-.07-.43-.04-.64.08s-.34.3-.41.53l-.23.9c-.01.08-.02.17-.02.27m2.76 2.15c0 .17.05.33.15.48c.1.15.25.26.46.32c.03 0 .08.01.14.02c.06.01.11.02.14.02c.41 0 .66-.22.76-.66l.14-.6c.07-.21.05-.42-.07-.63a.863.863 0 0 0-.51-.41c-.25-.06-.48-.04-.68.08s-.34.29-.41.53l-.09.59c0 .01 0 .05-.01.11c-.01.07-.02.11-.02.15m.74-2.96c0 .23.08.42.24.57c.15.16.34.24.58.24c.24 0 .43-.08.59-.23c.16-.16.23-.35.23-.58c0-.24-.08-.43-.23-.59a.784.784 0 0 0-.59-.23c-.24 0-.43.08-.59.23s-.23.35-.23.59m.71-12.35c.33-.56.78-.99 1.34-1.29s1.15-.45 1.76-.45h.22v.32c0 .64.11 1.26.34 1.86c.23.6.56 1.15 1.02 1.66c.45.51.99.91 1.61 1.21c-.17.38-.42.72-.76 1.03c-.91-.78-1.98-1.17-3.22-1.17h-.33c-.29-1.26-.95-2.32-1.98-3.17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightAltSnow(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.07 16.93c0 1.33.47 2.47 1.4 3.43s2.07 1.47 3.4 1.51c.12 0 .18-.06.18-.17v-1.34c0-.11-.06-.17-.18-.17c-.85-.04-1.58-.39-2.18-1.03c-.61-.64-.91-1.39-.91-2.24c0-.85.28-1.58.85-2.2c.57-.62 1.27-.96 2.1-1.03l.53-.07c.13 0 .2-.06.2-.17l.07-.52c.11-1.09.56-2.01 1.37-2.75s1.76-1.11 2.86-1.11c1.09 0 2.04.37 2.86 1.1c.82.73 1.28 1.64 1.4 2.72l.08.57c0 .12.06.18.17.18h1.62c.91 0 1.68.32 2.33.97c.65.64.97 1.41.97 2.31c0 .85-.3 1.6-.91 2.24c-.61.64-1.33.98-2.18 1.03c-.14 0-.21.06-.21.17v1.34c0 .11.07.17.21.17a4.71 4.71 0 0 0 2.41-.71c.73-.45 1.31-1.05 1.73-1.8s.63-1.56.63-2.43c0-.91-.22-1.74-.65-2.48c.74-.66 1.24-1.52 1.52-2.58l.17-.72c.01-.01.02-.04.02-.08c0-.07-.05-.13-.16-.16l-.61-.17c-.44-.13-.83-.32-1.17-.57s-.61-.53-.81-.84c-.2-.31-.34-.62-.44-.95c-.1-.32-.15-.64-.15-.95c0-.27.03-.56.1-.86l.11-.62c.02-.09-.02-.17-.14-.22l-.8-.24c-.44-.11-.85-.16-1.25-.16c-.37 0-.74.04-1.12.13c-.38.09-.77.22-1.18.41c-.41.19-.8.45-1.18.8c-.38.35-.71.75-.99 1.22c-.81-.33-1.6-.5-2.38-.5c-1.41 0-2.67.44-3.76 1.32s-1.8 2-2.11 3.37c-1.12.28-2.04.87-2.75 1.76c-.71.9-1.07 1.93-1.07 3.09m6.39 4.09c0 .24.08.44.24.6c.16.17.35.25.59.25s.44-.08.6-.25s.24-.37.24-.6c0-.22-.08-.42-.24-.58s-.36-.24-.6-.24c-.23 0-.43.08-.59.24c-.16.16-.24.35-.24.58m0 3.64c0 .23.08.42.24.58c.16.16.36.24.58.24c.24 0 .44-.08.6-.23c.16-.16.24-.35.24-.59s-.08-.43-.24-.59a.806.806 0 0 0-.6-.23c-.24 0-.43.08-.59.23c-.15.15-.23.35-.23.59m3.2-1.7c0 .24.08.44.24.59c.16.16.36.24.58.24c.24 0 .44-.08.61-.24s.25-.36.25-.59c0-.24-.08-.44-.25-.61s-.37-.26-.61-.26c-.22 0-.41.09-.58.26s-.24.37-.24.61m0-3.64c0 .24.08.43.24.58c.16.16.36.24.58.24c.24 0 .45-.08.61-.23s.25-.35.25-.59c0-.23-.08-.43-.25-.6s-.37-.25-.61-.25c-.22 0-.42.08-.58.25s-.24.37-.24.6m0 7.31c0 .22.08.41.24.57c.17.17.36.25.58.25c.24 0 .44-.08.61-.24c.17-.16.25-.35.25-.59s-.08-.44-.25-.61a.832.832 0 0 0-.61-.26c-.22 0-.41.09-.58.26c-.15.18-.24.39-.24.62m3.24-5.61c0 .24.08.44.25.6s.36.25.6.25s.43-.08.59-.25s.24-.37.24-.6c0-.22-.08-.42-.24-.58s-.35-.24-.59-.24s-.43.08-.6.24s-.25.35-.25.58m0 3.64c0 .23.08.42.24.58c.16.16.36.24.6.24s.43-.08.59-.24c.16-.16.23-.35.23-.59s-.08-.43-.23-.59c-.16-.16-.35-.23-.59-.23s-.44.08-.6.23c-.16.16-.24.36-.24.6m.68-15.89c.31-.54.75-.96 1.3-1.26s1.12-.45 1.71-.46c.15 0 .26.01.33.02v.31c0 .97.26 1.88.78 2.74s1.25 1.51 2.17 1.96c-.16.36-.41.72-.76 1.07c-.89-.79-1.96-1.18-3.23-1.18h-.31c-.27-1.23-.93-2.29-1.99-3.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightAltSnowThunderstorm(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.09 16.89c0 1.11.33 2.1.99 2.97c.66.87 1.52 1.47 2.58 1.79l-.65 1.7c-.04.14 0 .21.14.21h2.12l-1.29 4.18h.28l4.23-5.62c.04-.04.04-.09.02-.14c-.03-.05-.07-.07-.14-.07h-2.18l2.47-4.64c.07-.14.03-.22-.13-.22H9.57c-.09 0-.16.05-.22.15l-1.07 2.88c-.71-.18-1.3-.57-1.78-1.17s-.71-1.27-.71-2.01c0-.83.28-1.55.85-2.17c.57-.61 1.27-.97 2.1-1.07l.53-.07c.13 0 .2-.06.2-.18l.07-.51c.11-1.08.56-1.99 1.37-2.72c.81-.73 1.76-1.1 2.85-1.1s2.04.37 2.86 1.1c.82.73 1.28 1.64 1.4 2.71l.07.57c0 .12.06.19.17.19h1.62c.91 0 1.68.32 2.33.95s.97 1.4.97 2.28c0 .85-.3 1.59-.9 2.21c-.6.62-1.33.97-2.2 1.03c-.12 0-.19.06-.19.19v1.36c0 .12.06.18.19.18c1.33-.04 2.46-.55 3.39-1.51c.93-.97 1.39-2.12 1.39-3.45c0-.87-.22-1.68-.66-2.45c.76-.74 1.27-1.61 1.51-2.62l.19-.68c.01-.01.01-.03.01-.07c0-.08-.05-.13-.15-.16l-.62-.17c-.57-.17-1.06-.45-1.46-.84c-.4-.39-.68-.8-.85-1.22s-.25-.84-.24-1.26c0-.28.03-.56.1-.85l.11-.61c.02-.1-.02-.18-.14-.23l-.8-.24c-.47-.09-.88-.14-1.24-.14c-.37-.01-.75.03-1.13.12c-.38.08-.78.22-1.19.4c-.41.18-.8.45-1.18.79c-.38.34-.71.74-.99 1.2c-.83-.31-1.62-.46-2.36-.46c-1.41 0-2.67.44-3.76 1.32s-1.8 2-2.11 3.36c-1.11.26-2.02.84-2.74 1.74c-.71.89-1.07 1.92-1.07 3.07m9.59 6.07c0 .24.08.44.24.59c.16.16.36.24.58.24c.24 0 .44-.08.61-.24s.25-.36.25-.59c0-.24-.08-.44-.25-.61s-.37-.26-.61-.26c-.22 0-.41.09-.58.26s-.24.37-.24.61m0-3.64c0 .24.08.43.24.58c.16.16.36.24.58.24c.24 0 .45-.08.61-.23s.25-.35.25-.59c0-.23-.08-.43-.25-.6s-.37-.25-.61-.25c-.23 0-.42.08-.58.25s-.24.37-.24.6m0 7.31c0 .22.08.41.24.57c.17.17.36.25.58.25c.24 0 .44-.08.61-.24c.17-.16.25-.35.25-.59s-.08-.44-.25-.61a.832.832 0 0 0-.61-.26c-.22 0-.41.09-.58.26c-.16.18-.24.39-.24.62m3.23-5.61c0 .24.08.44.25.6s.36.25.6.25c.23 0 .43-.08.59-.25c.16-.17.24-.37.24-.6c0-.22-.08-.42-.24-.58a.791.791 0 0 0-.59-.24c-.23 0-.43.08-.6.24s-.25.35-.25.58m0 3.64c0 .23.08.42.24.58c.16.16.36.24.6.24s.43-.08.59-.24c.16-.16.23-.35.23-.59s-.08-.43-.23-.59a.784.784 0 0 0-.59-.23c-.24 0-.44.08-.6.23c-.16.16-.24.36-.24.6m.68-15.89c.33-.56.78-.99 1.34-1.29s1.15-.45 1.76-.45h.22v.32c0 .64.11 1.26.34 1.86c.23.6.56 1.15 1.02 1.66c.45.51.99.91 1.61 1.21c-.17.38-.42.72-.76 1.03c-.91-.78-1.98-1.17-3.22-1.17h-.33c-.29-1.26-.95-2.32-1.98-3.17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightAltSnowWind(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.07 16.93c0 1.33.47 2.47 1.4 3.43s2.07 1.47 3.4 1.51c.12 0 .18-.06.18-.17v-1.34c0-.11-.06-.17-.18-.17c-.85-.04-1.58-.39-2.18-1.03c-.61-.64-.91-1.39-.91-2.24c0-.85.28-1.58.85-2.2c.57-.62 1.27-.96 2.1-1.03l.53-.07c.13 0 .2-.06.2-.17l.07-.52c.11-1.09.56-2.01 1.37-2.75s1.76-1.11 2.86-1.11c1.09 0 2.04.37 2.86 1.1c.82.73 1.28 1.64 1.4 2.72l.08.57c0 .12.06.18.17.18h1.62c.91 0 1.68.32 2.33.97c.65.64.97 1.41.97 2.31c0 .85-.3 1.6-.91 2.24c-.61.64-1.33.98-2.18 1.03c-.14 0-.21.06-.21.17v1.34c0 .11.07.17.21.17a4.71 4.71 0 0 0 2.41-.71c.73-.45 1.31-1.05 1.73-1.8s.63-1.56.63-2.43c0-.91-.22-1.74-.65-2.48c.74-.66 1.24-1.52 1.52-2.58l.17-.72c.01-.01.02-.04.02-.08c0-.07-.05-.13-.16-.16l-.61-.17c-.44-.13-.83-.32-1.17-.57s-.61-.53-.81-.84c-.2-.31-.34-.62-.44-.95c-.1-.32-.15-.64-.15-.95c0-.27.03-.56.1-.86l.11-.62c.02-.09-.02-.17-.14-.22l-.8-.24c-.44-.11-.85-.16-1.25-.16c-.37 0-.74.04-1.12.13c-.38.09-.77.22-1.18.41c-.41.19-.8.45-1.18.8c-.38.35-.71.75-.99 1.22c-.81-.33-1.6-.5-2.38-.5c-1.41 0-2.67.44-3.76 1.32s-1.8 2-2.11 3.37c-1.12.28-2.04.87-2.75 1.76c-.71.9-1.07 1.93-1.07 3.09m5.53 7.73c0 .23.08.42.24.58c.16.16.36.24.58.24c.24 0 .44-.08.6-.23c.16-.16.24-.35.24-.59s-.08-.43-.24-.59c-.16-.16-.36-.23-.6-.23s-.42.07-.58.23c-.16.15-.24.35-.24.59m.43-3.64c0 .24.08.44.24.6s.35.25.59.25s.44-.08.6-.25c.16-.17.24-.37.24-.6c0-.22-.08-.42-.24-.58a.814.814 0 0 0-.6-.24c-.23 0-.43.08-.59.24s-.24.35-.24.58m2.35 5.61c0 .22.08.41.24.57c.17.17.36.25.58.25c.24 0 .44-.08.61-.24c.17-.16.25-.35.25-.59s-.08-.44-.25-.61a.832.832 0 0 0-.61-.26c-.22 0-.41.09-.58.26c-.16.18-.24.39-.24.62m.64-3.67c0 .24.08.43.25.59c.16.16.35.24.57.24c.24 0 .44-.08.61-.24c.17-.16.26-.36.26-.59c0-.24-.09-.44-.26-.61a.848.848 0 0 0-.61-.26c-.22 0-.41.09-.58.26s-.24.37-.24.61m.64-3.64c0 .24.08.43.24.58c.16.16.36.24.58.24c.24 0 .45-.08.61-.23s.25-.35.25-.59c0-.23-.08-.43-.25-.6s-.37-.25-.61-.25c-.22 0-.42.08-.58.25s-.24.37-.24.6m2.39 5.34c0 .23.08.42.24.58c.16.16.36.24.6.24c.23 0 .43-.08.59-.24c.16-.16.24-.35.24-.59s-.08-.43-.23-.59a.784.784 0 0 0-.59-.23c-.24 0-.44.08-.6.23c-.17.16-.25.36-.25.6m.42-3.64c0 .23.08.43.25.6c.17.17.37.25.6.25s.43-.08.59-.25s.24-.37.24-.6c0-.22-.08-.42-.24-.58s-.35-.24-.59-.24s-.43.08-.6.25c-.17.16-.25.35-.25.57m1.11-12.25c.31-.54.75-.96 1.3-1.26s1.12-.45 1.71-.46c.15 0 .26.01.33.02v.31c0 .97.26 1.88.78 2.74s1.25 1.51 2.17 1.96c-.16.36-.41.72-.76 1.07c-.89-.79-1.96-1.18-3.23-1.18h-.31c-.27-1.23-.93-2.29-1.99-3.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightAltSprinkle(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.11 16.89c0 1.33.46 2.48 1.39 3.43c.93.96 2.06 1.47 3.4 1.54c.12 0 .18-.06.18-.18v-1.33c0-.12-.06-.18-.18-.18c-.86-.04-1.58-.38-2.18-1.02c-.6-.64-.9-1.39-.9-2.26c0-.83.28-1.55.84-2.17c.56-.61 1.26-.97 2.1-1.07l.52-.04c.13 0 .2-.06.2-.18l.07-.54c.11-1.08.56-1.99 1.37-2.71c.81-.73 1.76-1.09 2.85-1.09s2.04.37 2.86 1.1a4.11 4.11 0 0 1 1.4 2.71l.07.57c0 .12.06.19.18.19h1.62c.89 0 1.65.32 2.3.96s.97 1.4.97 2.27c0 .87-.3 1.62-.9 2.26c-.6.64-1.33.98-2.18 1.02c-.12 0-.19.06-.19.18v1.33c0 .12.06.18.19.18c.88-.03 1.68-.27 2.41-.72s1.31-1.05 1.73-1.8s.63-1.57.63-2.44c0-.87-.23-1.68-.68-2.45c.78-.74 1.29-1.6 1.54-2.58l.14-.73c.01-.01.02-.03.02-.07c0-.07-.05-.13-.16-.16l-.57-.17c-.57-.16-1.06-.44-1.46-.82c-.41-.38-.7-.79-.87-1.21c-.17-.43-.26-.85-.26-1.28c0-.29.04-.57.11-.85l.13-.61c.02-.1-.02-.18-.13-.23l-.8-.24c-.45-.1-.87-.15-1.27-.15c-.36 0-.73.04-1.12.13c-.38.09-.78.22-1.19.4s-.81.44-1.2.79s-.72.74-1 1.2c-.81-.31-1.59-.46-2.33-.46c-1.41 0-2.67.44-3.76 1.32s-1.8 2-2.11 3.36a4.77 4.77 0 0 0-2.73 1.74c-.69.88-1.05 1.91-1.05 3.06m5.94.88c0 .38.14.71.42.98c.28.27.62.4 1.02.4s.73-.13 1-.4s.4-.59.4-.98c0-.26-.12-.6-.35-1.02c-.23-.42-.45-.75-.65-.98c-.11-.12-.24-.26-.41-.43l-.35.41c-.27.29-.52.64-.75 1.04s-.33.72-.33.98m2.99 3.99c0 .66.23 1.21.68 1.66c.46.45 1.01.67 1.65.67c.66 0 1.21-.23 1.66-.68c.45-.46.68-1.01.68-1.65c0-.55-.27-1.22-.8-2c-.44-.58-.87-1.08-1.28-1.49c-.08-.06-.17-.13-.26-.23l-.23.23c-.39.36-.82.86-1.28 1.48c-.24.33-.43.68-.59 1.04c-.16.37-.23.69-.23.97m1.47-6.67c0 .26.1.47.29.66s.42.27.7.27c.26 0 .47-.09.66-.27c.18-.18.27-.4.27-.66c0-.43-.31-.97-.93-1.62l-.25.27c-.18.2-.35.43-.5.7c-.17.27-.24.49-.24.65m3.05-6.32c.35-.57.8-1 1.34-1.29c.54-.29 1.12-.44 1.72-.44c.12 0 .21.01.27.02v.3c0 .96.26 1.87.79 2.74s1.25 1.52 2.18 1.97c-.16.38-.41.72-.75 1.03a4.909 4.909 0 0 0-3.21-1.14h-.33c-.3-1.31-.97-2.37-2.01-3.19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightAltStormShowers(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.09 16.89c0 1.11.33 2.1.99 2.97c.66.87 1.52 1.47 2.58 1.79l-.65 1.7c-.04.14 0 .21.14.21h2.12l-1.29 4.18h.28l4.23-5.62c.04-.04.04-.09.02-.14c-.03-.05-.07-.07-.14-.07h-2.18l2.47-4.64c.07-.14.03-.22-.13-.22H9.57c-.09 0-.16.05-.22.15l-1.07 2.88c-.71-.18-1.3-.57-1.78-1.17s-.71-1.27-.71-2.01c0-.83.28-1.55.85-2.17c.57-.61 1.27-.97 2.1-1.07l.53-.07c.13 0 .2-.06.2-.18l.07-.51c.11-1.08.56-1.99 1.37-2.72c.81-.73 1.76-1.1 2.85-1.1s2.04.37 2.86 1.1c.82.73 1.28 1.64 1.4 2.71l.07.57c0 .12.06.19.17.19h1.62c.91 0 1.68.32 2.33.95s.97 1.4.97 2.28c0 .85-.3 1.59-.9 2.21c-.6.62-1.33.97-2.2 1.03c-.12 0-.19.06-.19.19v1.36c0 .12.06.18.19.18c1.33-.04 2.46-.55 3.39-1.51c.93-.97 1.39-2.12 1.39-3.45c0-.87-.22-1.68-.66-2.45c.76-.74 1.27-1.61 1.51-2.62l.19-.68c.01-.01.01-.03.01-.07c0-.08-.05-.13-.15-.16l-.62-.17c-.57-.17-1.06-.45-1.46-.84c-.4-.39-.68-.8-.85-1.22s-.25-.84-.24-1.26c0-.28.03-.56.1-.85l.11-.61c.02-.1-.02-.18-.14-.23l-.8-.24c-.47-.09-.88-.14-1.24-.14c-.37-.01-.75.03-1.13.12c-.38.08-.78.22-1.19.4c-.41.18-.8.45-1.18.79c-.38.34-.71.74-.99 1.2c-.83-.31-1.62-.46-2.36-.46c-1.41 0-2.67.44-3.76 1.32s-1.8 2-2.11 3.36c-1.11.26-2.02.84-2.74 1.74c-.71.89-1.07 1.92-1.07 3.07m8.17 9.87c0 .16.05.31.15.47c.1.16.25.27.45.33c.16.03.25.05.27.05c.09 0 .22-.03.37-.1c.21-.1.35-.27.42-.53l.28-1.05c.06-.22.04-.43-.08-.63s-.29-.34-.53-.41a.817.817 0 0 0-.63.08c-.2.12-.34.3-.41.53l-.27 1zM13.6 22c0 .43.2.68.61.75c.14.03.23.05.27.05c.38 0 .63-.21.77-.63l.3-1.02c.06-.22.03-.43-.08-.63s-.3-.34-.53-.41a.759.759 0 0 0-.64.07c-.2.12-.34.29-.41.53l-.25 1.01c-.03.09-.04.18-.04.28m2.81 1.67c.01.17.07.33.18.48s.27.27.48.34c.16.04.27.06.33.06c.34 0 .58-.23.71-.68l.24-1.02c.07-.23.05-.45-.06-.66a.816.816 0 0 0-.5-.41a.926.926 0 0 0-.68.08c-.2.12-.33.3-.37.53l-.29 1.03c0 .02-.01.06-.02.12c-.02.07-.02.11-.02.13m1.18-14.9c.33-.56.78-.99 1.34-1.29s1.15-.45 1.76-.45h.22v.32c0 .64.11 1.26.34 1.86c.23.6.56 1.15 1.02 1.66c.45.51.99.91 1.61 1.21c-.17.38-.42.72-.76 1.03c-.91-.78-1.98-1.17-3.22-1.17h-.33c-.29-1.26-.95-2.32-1.98-3.17m.19 10.1c0 .43.22.71.65.82c.14.02.24.04.3.04c.36 0 .61-.22.74-.65l.28-1.04c.01-.05.01-.12.01-.22a.754.754 0 0 0-.14-.49a.831.831 0 0 0-.49-.33c-.01 0-.05 0-.1-.01s-.1-.01-.13-.01c-.16 0-.32.05-.48.15s-.27.26-.33.47l-.29 1.02c0 .01 0 .04-.01.1c0 .07-.01.12-.01.15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightAltThunderstorm(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.08 16.88c0 1.11.33 2.1.99 2.98s1.52 1.47 2.58 1.79l-.66 1.69c-.03.14.02.21.15.21h2.12l-.97 3.51h.29l3.91-4.94c.04-.05.04-.1.01-.15c-.03-.05-.08-.07-.15-.07h-2.18l2.48-4.63c.07-.14.03-.22-.13-.22H9.56c-.09 0-.16.05-.23.14l-1.07 2.88c-.72-.18-1.31-.57-1.78-1.17s-.7-1.27-.7-2.01c0-.83.28-1.55.85-2.17c.57-.62 1.27-.97 2.1-1.07l.52-.08c.13 0 .2-.06.2-.17l.07-.52c.11-1.08.56-1.99 1.37-2.72s1.76-1.1 2.85-1.1c1.08 0 2.03.37 2.85 1.1s1.29 1.64 1.41 2.71l.07.59c0 .11.06.17.18.17h1.62c.91 0 1.68.32 2.33.95s.97 1.4.97 2.29c0 .85-.3 1.59-.9 2.21c-.6.62-1.33.97-2.2 1.04c-.12 0-.19.06-.19.17v1.38c0 .12.06.18.19.18c.88-.03 1.68-.27 2.41-.72a4.94 4.94 0 0 0 2.36-4.25c0-.87-.22-1.68-.66-2.45c.79-.76 1.31-1.63 1.56-2.62l.14-.72c.01-.01.02-.04.02-.07c0-.07-.05-.12-.16-.15l-.56-.18c-.57-.16-1.06-.44-1.46-.82a3.4 3.4 0 0 1-.87-1.23c-.17-.44-.26-.88-.26-1.32c0-.26.03-.53.08-.8l.14-.61c.04-.1 0-.18-.14-.23c-.21-.09-.51-.17-.9-.26s-.77-.13-1.15-.13c-.36 0-.73.04-1.12.13c-.38.09-.78.22-1.19.41c-.41.18-.81.45-1.2.8c-.39.35-.72.75-1 1.22c-.82-.3-1.62-.45-2.38-.45c-1.41 0-2.67.44-3.76 1.31s-1.8 1.99-2.11 3.36c-1.11.26-2.02.84-2.74 1.74a4.91 4.91 0 0 0-1.04 3.07m8.1 9.82c0 .16.05.32.15.46c.1.15.25.25.45.3c.11.02.21.03.3.03c.41 0 .66-.21.76-.63l2.32-8.79c.06-.24.04-.45-.07-.65a.786.786 0 0 0-.5-.39a.814.814 0 0 0-.65.06c-.2.11-.34.27-.4.49l-2.32 8.84c-.03.1-.04.19-.04.28m4.17-3.02c0 .16.05.32.15.46c.1.14.25.25.46.31c.03 0 .08 0 .15.01c.07.01.13.01.16.01c.38 0 .62-.21.72-.63l1.5-5.77c.06-.24.04-.46-.08-.66a.799.799 0 0 0-.51-.38a.814.814 0 0 0-.65.06c-.2.11-.33.27-.39.5l-1.5 5.82c0 .1-.01.19-.01.27m1.24-14.93c.33-.57.77-1 1.33-1.3c.55-.3 1.14-.45 1.76-.45c.12 0 .22 0 .27.01v.32c0 .96.26 1.87.78 2.73s1.25 1.51 2.17 1.97c-.18.42-.44.77-.79 1.07c-.92-.79-1.99-1.18-3.22-1.18h-.32c-.28-1.26-.94-2.31-1.98-3.17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightClear(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.91 14.48c0-.96.19-1.87.56-2.75s.88-1.63 1.51-2.26c.63-.63 1.39-1.14 2.27-1.52c.88-.38 1.8-.57 2.75-.57h1.14c.16.04.23.14.23.28l.05.88c.04 1.27.49 2.35 1.37 3.24c.88.89 1.94 1.37 3.19 1.42l.82.07c.16 0 .24.08.24.23v.98c.01 1.28-.3 2.47-.93 3.56a7.017 7.017 0 0 1-2.57 2.59c-1.08.63-2.27.95-3.55.95c-.97 0-1.9-.19-2.78-.56s-1.63-.88-2.26-1.51a7.084 7.084 0 0 1-1.5-2.26c-.35-.88-.54-1.8-.54-2.77m1.83 0c0 .76.15 1.48.45 2.16c.3.67.7 1.24 1.19 1.7c.49.46 1.05.82 1.69 1.08a4.823 4.823 0 0 0 3.7.06a6.24 6.24 0 0 0 1.65-.96c.51-.41.94-.93 1.31-1.57c.37-.64.6-1.33.71-2.09c-1.63-.34-2.94-1.04-3.92-2.1s-1.55-2.3-1.7-3.74c-.96.06-1.82.35-2.61.88c-.78.53-1.39 1.2-1.82 2.02c-.43.82-.65 1.67-.65 2.56"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightCloudy(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.3 16.89c0 .89.22 1.72.66 2.48s1.03 1.36 1.79 1.8s1.58.67 2.48.67h10.81c.89 0 1.72-.22 2.48-.67s1.36-1.05 1.8-1.8c.44-.76.67-1.59.67-2.48c0-.64-.14-1.3-.42-2c.76-.93 1.13-2.03 1.13-3.3a5.195 5.195 0 0 0-2.59-4.54c-.8-.47-1.68-.7-2.63-.7c-1.49 0-2.78.58-3.87 1.74c-.76-.43-1.66-.65-2.69-.65c-1.41 0-2.65.43-3.73 1.3s-1.77 1.98-2.08 3.35c-1.12.25-2.03.82-2.74 1.72a4.85 4.85 0 0 0-1.07 3.08m1.71 0c0-.83.28-1.55.83-2.16c.56-.61 1.26-.96 2.1-1.06l.68-.03l.07-.72c.14-1.08.61-1.99 1.41-2.71c.8-.73 1.74-1.09 2.81-1.09c1.09 0 2.05.37 2.86 1.1s1.27 1.63 1.38 2.71l.1.75h1.78c.88 0 1.64.32 2.28.95s.96 1.39.96 2.26c0 .9-.32 1.67-.95 2.32s-1.4.97-2.28.97H9.23c-.87 0-1.62-.32-2.26-.97c-.64-.66-.96-1.43-.96-2.32m12.03-7.83c.69-.66 1.5-.99 2.44-.99c.99 0 1.83.34 2.52 1.03c.69.68 1.04 1.52 1.04 2.51c0 .62-.17 1.23-.52 1.84c-.96-.97-2.12-1.45-3.49-1.45h-.31c-.27-1.11-.83-2.09-1.68-2.94"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightCloudyGusts(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.06 20.98c0 .24.09.44.27.6c.18.18.38.27.61.27h5.88c.26 0 .49.09.69.28c.2.19.3.42.3.68c0 .26-.1.48-.3.68s-.43.3-.68.3c-.25 0-.48-.1-.68-.3c-.19-.17-.39-.26-.62-.26c-.23 0-.43.08-.59.25c-.16.17-.24.37-.24.61s.08.44.24.6c.52.52 1.15.78 1.88.78c.74 0 1.38-.26 1.89-.77c.52-.52.78-1.14.78-1.88s-.26-1.38-.78-1.9s-1.15-.79-1.89-.79H3.94a.86.86 0 0 0-.62.25a.79.79 0 0 0-.26.6m0-3.01c0 .23.09.42.27.58c.15.16.35.24.61.24h10.99c.74 0 1.37-.26 1.89-.78c.52-.52.78-1.15.78-1.88s-.26-1.36-.78-1.88c-.52-.52-1.15-.77-1.89-.77c-.76 0-1.39.25-1.89.75c-.15.17-.23.38-.23.63c0 .24.08.43.23.59s.35.23.6.23s.45-.07.61-.23c.19-.19.42-.28.68-.28c.26 0 .48.09.68.28c.19.19.29.42.29.68s-.1.5-.29.69c-.19.19-.42.29-.68.29H3.94a.86.86 0 0 0-.62.25c-.18.18-.26.38-.26.61m2.65-2.34c0 .08.06.12.17.12h1.43c.08 0 .15-.05.22-.14c.23-.54.57-.99 1.05-1.35c.47-.36 1-.56 1.59-.6l.52-.07c.12 0 .19-.06.19-.19l.07-.5c.11-1.08.57-1.99 1.38-2.71c.81-.73 1.77-1.09 2.86-1.09s2.04.36 2.85 1.08c.81.72 1.27 1.63 1.39 2.72l.07.57c0 .12.06.18.18.18h1.63c.9 0 1.67.32 2.31.95c.64.63.96 1.39.96 2.28c0 .89-.32 1.66-.96 2.29c-.64.63-1.41.95-2.31.95H14.4c-.11 0-.17.06-.17.18v1.37c0 .12.06.18.17.18h6.91c.89 0 1.72-.22 2.48-.67c.76-.44 1.36-1.05 1.8-1.81c.44-.76.66-1.59.66-2.49c0-.74-.14-1.42-.42-2.02c.76-1 1.14-2.11 1.14-3.33c0-.71-.14-1.39-.42-2.04s-.65-1.2-1.12-1.67c-.47-.47-1.03-.84-1.67-1.11s-1.34-.41-2.05-.41c-1.54 0-2.84.58-3.88 1.73c-.78-.41-1.67-.61-2.65-.61c-1.41 0-2.66.44-3.75 1.31s-1.77 1.99-2.07 3.35c-.85.2-1.6.61-2.26 1.23s-1.11 1.35-1.37 2.18v.04c-.01.05-.02.09-.02.1M19.24 9c.72-.68 1.54-1.02 2.48-1.02c.98 0 1.81.35 2.51 1.05s1.05 1.53 1.05 2.5c0 .61-.17 1.22-.51 1.85c-.96-.96-2.11-1.43-3.47-1.43h-.33c-.24-1.07-.81-2.06-1.73-2.95"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightCloudyHigh(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.58 13.45c0-1.15.36-2.18 1.08-3.07c.72-.9 1.63-1.48 2.74-1.74c.31-1.37 1.02-2.49 2.11-3.37s2.35-1.32 3.76-1.32c1.38 0 2.61.43 3.69 1.28s1.78 1.95 2.1 3.29h.33c.9 0 1.73.22 2.49.65s1.37 1.03 1.81 1.79c.44.76.67 1.58.67 2.48c0 .2-.01.4-.03.61c.65.51 1.16 1.15 1.54 1.91s.56 1.57.56 2.43c0 .77-.15 1.5-.45 2.19c-.3.69-.7 1.28-1.2 1.78c-.5.49-1.1.89-1.79 1.18a5.537 5.537 0 0 1-4.33 0c-.69-.29-1.28-.69-1.78-1.19s-.89-1.09-1.19-1.78s-.45-1.41-.45-2.16H8.38c-1.34-.06-2.47-.57-3.4-1.53c-.93-.94-1.4-2.09-1.4-3.43m1.71 0c0 .87.3 1.62.9 2.26c.6.64 1.33.98 2.19 1.03h11.19c.86-.04 1.59-.39 2.19-1.03c.61-.64.91-1.4.91-2.26c0-.88-.33-1.63-.98-2.27s-1.42-.96-2.32-.96h-1.62c-.11 0-.17-.06-.17-.17l-.07-.58c-.11-1.08-.58-1.99-1.4-2.72s-1.77-1.1-2.86-1.1c-1.09 0-2.05.37-2.85 1.1S9.14 8.39 9.04 9.47l-.08.58c0 .11-.07.17-.2.17h-.52c-.84.1-1.54.46-2.1 1.07c-.57.61-.85 1.33-.85 2.16m11.26 5.11c.06 1.12.52 2.07 1.37 2.83c.85.76 1.82 1.14 2.91 1.14c.6 0 1.17-.12 1.7-.35s.98-.55 1.34-.93c.36-.39.65-.83.85-1.33c.21-.5.31-1 .31-1.52c0-.49-.1-.98-.3-1.47s-.48-.94-.85-1.35c-.39.82-.97 1.5-1.74 2.02s-1.63.79-2.57.83h-3.03c0 .01 0 .04.01.13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightCloudyWindy(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.43 21c0 .25.09.45.27.6c.17.17.37.26.61.26h9.54c.23 0 .43-.08.59-.25c.16-.17.24-.37.24-.61s-.08-.44-.24-.61a.764.764 0 0 0-.59-.25H3.31a.87.87 0 0 0-.62.26c-.17.17-.26.37-.26.6m2.64-3.03c0 .23.09.42.27.58c.16.16.36.24.6.24h9.55c.24 0 .43-.08.59-.23c.16-.16.23-.35.23-.59s-.08-.44-.24-.6a.764.764 0 0 0-.59-.25H5.94c-.24 0-.44.08-.61.25a.79.79 0 0 0-.26.6m1.14-2.33c0 .07.07.11.2.11h1.42c.09 0 .16-.05.23-.14c.22-.54.57-.99 1.05-1.35c.47-.36 1-.56 1.58-.6l.54-.07c.12 0 .18-.06.18-.18l.07-.51c.11-1.08.57-1.99 1.38-2.72c.81-.73 1.77-1.1 2.87-1.1s2.06.36 2.87 1.09c.81.72 1.28 1.63 1.39 2.73l.08.57c0 .12.06.19.17.19h1.62c.91 0 1.69.32 2.33.95c.64.63.96 1.39.96 2.29c0 .89-.32 1.65-.96 2.29c-.64.64-1.42.96-2.33.96h-6.91c-.11 0-.17.06-.17.17v1.38c0 .12.06.18.17.18h6.91c.91 0 1.74-.22 2.51-.67c.77-.44 1.37-1.05 1.82-1.81c.45-.76.67-1.59.67-2.49c0-.71-.15-1.37-.44-2.01c.77-1.01 1.15-2.1 1.15-3.29c0-.95-.24-1.83-.71-2.64s-1.11-1.45-1.92-1.92c-.81-.47-1.69-.7-2.64-.7c-1.52 0-2.81.56-3.84 1.67a6.09 6.09 0 0 0-2.72-.62c-.93 0-1.81.2-2.63.59s-1.51.95-2.07 1.66a5.78 5.78 0 0 0-1.13 2.43c-.88.2-1.64.6-2.29 1.2c-.65.6-1.11 1.33-1.36 2.17zm.62 8.45c0 .23.09.43.26.58c.16.16.36.24.6.24h9.56c.24 0 .44-.08.6-.23s.25-.35.25-.59s-.08-.44-.25-.61a.816.816 0 0 0-.6-.25H7.69c-.23 0-.43.09-.6.26c-.17.17-.26.37-.26.6m13-15.07c.67-.65 1.5-.98 2.47-.98c.99 0 1.83.35 2.52 1.04c.69.69 1.04 1.53 1.04 2.52c0 .63-.16 1.22-.49 1.77c-.98-.96-2.15-1.43-3.52-1.43h-.32c-.23-1.1-.8-2.07-1.7-2.92"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightFog(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.66 20.92c0 .23.08.42.25.57c.17.16.38.23.62.23h18.61c.24 0 .44-.08.6-.23c.17-.16.25-.35.25-.57c0-.24-.08-.45-.24-.61a.807.807 0 0 0-.61-.25H3.53c-.24 0-.44.08-.61.25c-.17.17-.26.38-.26.61m2.61-3.11c0 .24.09.43.26.59c.14.18.33.27.59.27h18.61c.23 0 .42-.08.58-.25s.23-.37.23-.61a.8.8 0 0 0-.23-.58a.8.8 0 0 0-.58-.23H6.12c-.24 0-.44.08-.6.23c-.17.16-.25.35-.25.58m.15-2.42v-.05c-.04.15 0 .22.12.22h1.44c.06 0 .12-.05.19-.15c.24-.52.59-.94 1.06-1.27c.47-.33.99-.52 1.55-.56l.53-.08c.12 0 .19-.06.19-.18l.06-.5c.11-1.08.56-1.97 1.36-2.7c.8-.72 1.75-1.08 2.84-1.08c1.07 0 2.02.36 2.82 1.07s1.27 1.6 1.38 2.67l.07.57c0 .12.07.18.21.18h1.58c.64 0 1.23.17 1.75.52c.52.34.92.8 1.17 1.36c.07.1.14.15.22.15h1.42c.12 0 .17-.07.15-.22c-.22-.56-.37-.91-.46-1.06c.72-.65 1.23-1.51 1.5-2.57l.17-.66a.15.15 0 0 0-.01-.16c-.03-.04-.07-.07-.12-.07l-.62-.22c-.89-.26-1.57-.78-2.04-1.58s-.59-1.65-.37-2.56l.13-.58c.05-.09.01-.17-.13-.23l-.84-.23a5.026 5.026 0 0 0-3.22.26a5.21 5.21 0 0 0-2.47 2.12c-.79-.31-1.56-.46-2.29-.46c-1.39 0-2.62.44-3.71 1.31S9.27 10.64 8.95 12c-.84.2-1.58.6-2.22 1.21s-1.06 1.34-1.31 2.18M7 23.97c0 .24.09.43.26.59c.17.18.37.27.59.27H26.5c.23 0 .43-.08.59-.25c.16-.17.24-.37.24-.61c0-.23-.08-.42-.24-.58s-.36-.23-.59-.23H7.86c-.24 0-.44.08-.6.23c-.17.16-.26.35-.26.58M18.51 8.7c.35-.57.82-1.02 1.41-1.33c.59-.31 1.21-.44 1.87-.38c-.07 1.04.17 2.02.7 2.93c.54.91 1.28 1.58 2.22 2.02c-.15.35-.4.71-.75 1.07a4.791 4.791 0 0 0-3.14-1.13h-.32c-.32-1.31-.98-2.37-1.99-3.18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightHail(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.25 16.89c0 1.33.46 2.48 1.39 3.43c.93.96 2.06 1.47 3.4 1.54c.12 0 .18-.06.18-.18v-1.33c0-.12-.06-.18-.18-.18c-.86-.04-1.58-.38-2.18-1.02c-.6-.64-.9-1.39-.9-2.26c0-.83.28-1.55.84-2.17c.56-.61 1.26-.97 2.1-1.07l.52-.07c.13 0 .2-.06.2-.18l.08-.51c.11-1.08.56-1.99 1.37-2.71c.81-.73 1.76-1.09 2.85-1.09s2.04.37 2.86 1.1s1.28 1.63 1.4 2.71l.07.58c0 .12.06.18.18.18h1.62c.91 0 1.68.32 2.32.95c.64.63.96 1.39.96 2.29c0 .87-.3 1.62-.9 2.26c-.6.64-1.33.98-2.18 1.02c-.13 0-.19.06-.19.18v1.33c0 .12.06.18.19.18c.88-.03 1.68-.27 2.41-.72c.73-.45 1.31-1.05 1.73-1.8s.63-1.57.63-2.44c0-.74-.14-1.41-.43-2.01c.79-.96 1.18-2.06 1.18-3.32c0-.94-.24-1.82-.71-2.62a5.201 5.201 0 0 0-1.92-1.92c-.81-.47-1.68-.71-2.62-.71c-1.56 0-2.85.58-3.88 1.73c-.82-.44-1.72-.66-2.71-.66c-1.41 0-2.67.44-3.76 1.32s-1.8 2-2.11 3.36c-1.1.26-2.01.84-2.73 1.74c-.72.89-1.08 1.92-1.08 3.07m5.37 6.98c.09.22.24.37.46.46c.2.1.41.11.62.02c.22-.08.36-.23.45-.45c.09-.22.09-.44.01-.65a.814.814 0 0 0-.44-.47c-.2-.08-.4-.08-.61.01a.88.88 0 0 0-.46.43c-.11.17-.12.39-.03.65m.63-2.83c0 .16.05.31.15.46c.1.15.26.25.46.31c.22.07.44.05.65-.06s.35-.29.43-.55l.98-3.11c.07-.24.05-.47-.08-.67a.813.813 0 0 0-.55-.38a.744.744 0 0 0-.62.06c-.2.11-.33.28-.4.5l-1 3.18zm1.86 5.6c0 .07.02.17.06.29c.09.22.25.38.46.45c.08.05.19.08.33.08c.06 0 .16-.02.3-.06c.22-.08.38-.23.47-.45c.1-.22.1-.44 0-.66c-.1-.22-.25-.37-.45-.46s-.41-.09-.62-.01c-.19.08-.33.2-.42.36c-.09.16-.13.31-.13.46m.74-2.67c0 .18.05.33.15.48c.1.14.26.24.48.28c.02 0 .06.01.11.02s.1.02.13.02c.43-.03.7-.24.81-.62l1.76-6.07c.07-.24.05-.46-.06-.65a.856.856 0 0 0-.5-.39a.814.814 0 0 0-.65.06c-.2.11-.33.28-.4.5l-1.8 6.07c0 .02 0 .06-.01.1c-.01.04-.01.08-.01.11c-.01.04-.01.07-.01.09m3.44-.4c0 .1.02.21.05.32c.08.21.23.36.46.44c.09.04.21.07.36.07c.12 0 .22-.02.29-.06c.23-.09.38-.23.46-.43a.91.91 0 0 0 0-.65a.898.898 0 0 0-.42-.48a.86.86 0 0 0-.65.01c-.22.09-.37.23-.47.43c-.05.11-.08.22-.08.35m.72-2.54c0 .36.21.61.62.75c.14.04.24.06.3.06c.12 0 .23-.03.34-.08c.17-.09.31-.27.4-.55l.98-3.11a.75.75 0 0 0-.06-.64a.885.885 0 0 0-.51-.4a.79.79 0 0 0-.64.06c-.19.11-.33.28-.4.5l-.98 3.13c-.04.16-.05.26-.05.28m1.03-12.01c.69-.66 1.51-.99 2.48-.99s1.81.35 2.5 1.04c.69.69 1.04 1.53 1.04 2.5c0 .62-.17 1.23-.52 1.84c-.98-.98-2.14-1.47-3.49-1.47h-.33c-.31-1.16-.87-2.13-1.68-2.92"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightLightning(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.29 16.93c0 .66.12 1.28.38 1.88s.59 1.11 1.02 1.55s.94.79 1.52 1.05s1.21.42 1.87.45c.12 0 .18-.06.18-.17v-1.34c0-.11-.06-.17-.18-.17c-.87-.06-1.6-.41-2.19-1.03A3.13 3.13 0 0 1 6 16.93c0-.84.28-1.57.85-2.18c.57-.62 1.26-.97 2.1-1.04l.52-.06c.12 0 .19-.06.19-.18l.08-.52c.07-.71.3-1.36.69-1.94s.9-1.04 1.52-1.36s1.29-.49 2.02-.49c1.09 0 2.04.36 2.85 1.08c.81.72 1.27 1.62 1.39 2.69l.07.58c0 .11.06.17.19.17h1.6c.9 0 1.67.32 2.32.96c.64.64.97 1.4.97 2.29c0 .86-.3 1.6-.89 2.22c-.59.62-1.33.97-2.19 1.03c-.13 0-.2.06-.2.17v1.34c0 .11.07.17.2.17c1.34-.06 2.47-.57 3.38-1.51c.88-.95 1.34-2.09 1.34-3.42c0-.64-.16-1.32-.47-2.06c.79-.99 1.19-2.08 1.19-3.27c0-.95-.24-1.83-.71-2.63c-.47-.81-1.11-1.44-1.91-1.91s-1.68-.7-2.62-.7c-1.59 0-2.88.58-3.87 1.73c-.81-.43-1.7-.64-2.66-.64c-1.41 0-2.66.44-3.75 1.32s-1.79 2-2.1 3.37c-1.12.26-2.03.83-2.74 1.72c-.72.89-1.07 1.91-1.07 3.07m7.65 4.62c-.02.14.02.21.14.21h2.17l-1.32 4.17h.29l4.18-5.58c.04-.04.05-.09.02-.14s-.07-.07-.14-.07H15.1l2.3-4.24c.07-.14.03-.22-.14-.22h-2.94c-.08 0-.14.05-.21.14zM18 9.05c.67-.66 1.49-.99 2.47-.99c.98 0 1.81.34 2.5 1.03c.69.68 1.03 1.52 1.03 2.51c0 .59-.17 1.19-.52 1.8c-.97-.93-2.12-1.4-3.45-1.4h-.31c-.28-1.19-.86-2.17-1.72-2.95"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightPartlyCloudy(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.77 19.61c0 .97.35 1.81 1.06 2.52c.71.71 1.55 1.06 2.52 1.06h6.85c.97 0 1.8-.35 2.49-1.05c.69-.7 1.04-1.54 1.04-2.53c0-.48-.07-.89-.21-1.23c.83-.53 1.49-1.24 1.97-2.12c.48-.88.73-1.83.73-2.84c0-.81-.16-1.59-.47-2.33c-.32-.74-.74-1.38-1.28-1.91s-1.17-.96-1.91-1.28c-.74-.32-1.51-.48-2.32-.48c-1.09 0-2.1.27-3.02.81s-1.65 1.27-2.18 2.18c-.53.92-.79 1.92-.79 3.01v.34c-1.01.57-1.67 1.41-1.99 2.49c-.76.24-1.36.66-1.81 1.27a3.44 3.44 0 0 0-.68 2.09m1.98 0c0-.42.13-.78.4-1.08c.27-.3.61-.47 1.02-.51l.63-.08c.12 0 .19-.08.19-.23l.1-.56c.07-.58.31-1.06.73-1.44c.42-.39.91-.58 1.48-.58c.58 0 1.09.19 1.51.58c.43.39.68.87.75 1.44l.08.65c.06.15.15.23.25.23h1.31c.43 0 .8.16 1.12.47c.32.31.47.68.47 1.12c0 .45-.16.83-.47 1.15s-.69.48-1.12.48h-6.85c-.45 0-.83-.16-1.14-.48s-.46-.71-.46-1.16m4.43-6.39c.07-1.09.49-2.01 1.27-2.76c.78-.74 1.71-1.12 2.8-1.12c1.11 0 2.06.4 2.84 1.19c.78.79 1.17 1.76 1.17 2.89c0 .7-.17 1.35-.51 1.95c-.34.6-.8 1.08-1.38 1.45c-.59-.49-1.27-.73-2.03-.73c-.29-.88-.81-1.57-1.54-2.09c-.73-.52-1.56-.78-2.49-.78z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightRain(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.26 16.93c0 .66.12 1.28.38 1.88s.59 1.11 1.02 1.55c.43.43.94.79 1.53 1.05s1.21.42 1.87.45c.11 0 .17-.06.17-.17v-1.34c0-.11-.06-.17-.17-.17c-.87-.06-1.6-.41-2.19-1.03a3.13 3.13 0 0 1-.89-2.22c0-.84.28-1.57.85-2.19c.57-.62 1.26-.97 2.1-1.04l.53-.06c.13 0 .2-.06.2-.19l.06-.51c.11-1.09.56-1.99 1.37-2.71s1.76-1.08 2.86-1.08c1.09 0 2.05.36 2.85 1.07c.81.72 1.27 1.61 1.38 2.69l.07.58c0 .12.06.18.19.18h1.6c.9 0 1.67.32 2.32.96s.97 1.4.97 2.29c0 .86-.3 1.6-.89 2.22c-.59.62-1.33.97-2.19 1.03c-.13 0-.2.06-.2.17v1.34c0 .11.07.17.2.17c1.34-.06 2.47-.57 3.39-1.51c.92-.95 1.38-2.09 1.38-3.42c0-.72-.14-1.38-.42-1.99c.78-.94 1.17-2.06 1.17-3.36c0-.94-.23-1.81-.7-2.62s-1.11-1.45-1.91-1.92s-1.68-.71-2.62-.71c-1.56 0-2.85.58-3.88 1.73c-.88-.43-1.78-.65-2.7-.65c-1.41 0-2.66.44-3.75 1.32s-1.79 2-2.1 3.38c-1.1.26-2.01.83-2.73 1.73c-.76.93-1.12 1.95-1.12 3.1m5.49 6.68c0 .4.22.66.65.78c.21.07.42.05.63-.06c.21-.11.35-.28.41-.5l1.5-5.73a.78.78 0 0 0-.09-.63c-.12-.2-.3-.33-.54-.4c-.22-.07-.43-.05-.63.07s-.33.29-.39.52l-1.5 5.7c-.03.11-.04.19-.04.25m2.59 3.05c0 .12.03.24.08.37c.1.21.27.35.51.43c.02 0 .06 0 .1.01s.08.01.11.01h.09c.43 0 .68-.22.76-.66l2.3-8.74c.07-.22.05-.43-.06-.63a.835.835 0 0 0-.5-.4c-.24-.07-.47-.05-.68.07s-.33.29-.38.52l-2.31 8.75c-.01.1-.02.19-.02.27m4.16-3.06c0 .16.05.31.16.47a.819.819 0 0 0 .68.35c.17 0 .33-.05.49-.14c.16-.09.26-.24.32-.45l1.5-5.73c.08-.21.05-.41-.07-.62a.854.854 0 0 0-.52-.41a.796.796 0 0 0-.66.07c-.2.12-.32.29-.36.52l-1.5 5.7c-.03.11-.04.19-.04.24m1.53-14.52c.67-.67 1.49-1 2.48-1c.98 0 1.81.34 2.49 1.02c.69.68 1.03 1.51 1.03 2.48c0 .63-.17 1.24-.52 1.85c-.95-.95-2.11-1.43-3.49-1.43h-.31c-.28-1.17-.84-2.14-1.68-2.92"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightRainMix(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.19 16.91c0 .87.21 1.68.64 2.43c.42.75 1.01 1.35 1.74 1.8c.73.46 1.54.7 2.43.72c.12 0 .19-.06.19-.17v-1.34c0-.12-.06-.18-.19-.18c-.86-.04-1.59-.39-2.19-1.03s-.91-1.39-.91-2.24c0-.85.28-1.59.85-2.21c.57-.62 1.27-.97 2.11-1.04l.52-.07c.13 0 .2-.06.2-.17l.07-.52c.11-1.1.57-2.02 1.38-2.76s1.77-1.11 2.87-1.11c1.09 0 2.04.37 2.86 1.1c.82.73 1.28 1.65 1.4 2.73l.08.58c0 .11.06.17.18.17h1.62c.9 0 1.67.32 2.32.97c.65.65.97 1.42.97 2.32c0 .85-.3 1.6-.91 2.24c-.61.64-1.33.98-2.18 1.03c-.13 0-.2.06-.2.18v1.34c0 .11.07.17.2.17c.88-.02 1.69-.26 2.42-.72c.73-.45 1.31-1.05 1.73-1.8s.63-1.56.63-2.43c0-.73-.14-1.4-.42-2.02c.81-.99 1.21-2.12 1.21-3.37c0-.96-.24-1.84-.71-2.65s-1.11-1.45-1.92-1.92c-.81-.47-1.69-.71-2.64-.71c-.74 0-1.46.15-2.15.46c-.68.34-1.27.76-1.77 1.31c-.88-.43-1.78-.65-2.71-.65c-1.42 0-2.68.44-3.78 1.32s-1.81 2-2.12 3.37c-1.12.29-2.04.88-2.76 1.78c-.71.9-1.06 1.93-1.06 3.09m5.33 7.07c0 .17.05.34.16.51c.11.17.27.28.47.35c.23.07.44.06.64-.04c.19-.09.32-.28.39-.56l.14-.61c.05-.23.02-.44-.09-.63a.82.82 0 0 0-.52-.4c-.22-.07-.44-.04-.64.08s-.34.3-.4.53l-.14.59c0 .03-.01.09-.01.18m.76-2.9c0 .21.08.4.25.57c.16.17.34.25.56.25c.24 0 .44-.08.6-.24c.16-.16.24-.35.24-.59c0-.23-.08-.43-.24-.59a.814.814 0 0 0-.6-.24c-.23 0-.42.08-.58.23c-.15.18-.23.38-.23.61m.61-2.27c-.01.16.03.31.14.45c.1.15.26.25.48.32a.8.8 0 0 0 .62-.07c.21-.11.34-.28.41-.51l.28-.9c.07-.24.05-.46-.07-.65a.913.913 0 0 0-.54-.39a.737.737 0 0 0-.63.07a.85.85 0 0 0-.41.5l-.24.92c0 .02-.01.06-.02.12c-.01.05-.02.1-.02.14m1.16 8.29c0 .18.05.34.15.5c.1.16.26.27.48.33c.08.02.17.03.25.03c.43 0 .69-.2.79-.61l.14-.59c.06-.26.03-.48-.08-.68s-.29-.32-.52-.37a.744.744 0 0 0-.63.07c-.21.12-.34.29-.41.51l-.14.59c-.02.09-.03.16-.03.22m.78-2.9c0 .22.08.41.25.58c.16.16.35.24.57.24c.24 0 .43-.08.59-.23c.16-.16.23-.35.23-.59a.784.784 0 0 0-.82-.81c-.24 0-.43.08-.59.23a.8.8 0 0 0-.23.58m.63-2.27c-.01.15.03.31.13.47s.25.26.45.3c.23.06.44.04.64-.06s.33-.29.41-.56l.27-.9c.07-.22.05-.43-.07-.63a.867.867 0 0 0-.53-.4a.766.766 0 0 0-.64.08c-.21.12-.34.3-.41.53l-.23.9c-.01.08-.02.17-.02.27m2.75 2.15c0 .16.05.32.15.48s.26.27.46.33c.03 0 .08.01.14.02s.1.02.14.02c.41 0 .66-.22.76-.66l.14-.6c.07-.21.05-.42-.07-.63a.809.809 0 0 0-.51-.41c-.25-.06-.48-.04-.68.08c-.2.12-.34.29-.41.53l-.09.59c0 .02-.01.07-.02.12s-.01.09-.01.13m.74-2.96c0 .22.08.42.25.57c.15.16.34.24.57.24c.24 0 .43-.08.59-.23c.16-.16.23-.35.23-.58c0-.24-.08-.43-.23-.59a.784.784 0 0 0-.59-.23c-.24 0-.43.08-.59.23s-.23.35-.23.59m.61-2.31c0 .17.05.33.16.48s.27.26.49.32c.02 0 .06.01.12.02s.11.02.14.02c.1 0 .22-.03.36-.09c.21-.11.35-.29.41-.52l.24-.9c.06-.23.04-.44-.07-.63a.827.827 0 0 0-.51-.4a.79.79 0 0 0-.64.06a.82.82 0 0 0-.39.51l-.28.91c0 .02-.01.06-.02.12c0 .03-.01.07-.01.1m.47-9.8c.69-.69 1.53-1.04 2.51-1.04c.98 0 1.82.35 2.51 1.05c.69.7 1.04 1.54 1.04 2.52c0 .67-.17 1.28-.51 1.85c-.96-.96-2.14-1.44-3.54-1.44h-.32c-.28-1.18-.84-2.16-1.69-2.94"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightRainWind(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.11 17.02c0 1.13.33 2.13 1 3.01c.67.88 1.54 1.48 2.62 1.8c.1.01.18-.01.25-.08l1.13-1.46c-.89 0-1.66-.32-2.31-.96s-.97-1.41-.97-2.31c0-.86.29-1.61.86-2.24s1.29-.98 2.14-1.05l.52-.07c.11 0 .16-.05.16-.14l.07-.56c.12-1.1.59-2.02 1.41-2.76c.82-.74 1.78-1.11 2.88-1.11c1.11 0 2.08.37 2.91 1.1c.83.73 1.3 1.64 1.4 2.74l.07.59c.02.11.09.17.21.17h1.63c.9 0 1.67.33 2.32.98c.65.66.98 1.44.98 2.35c0 .84-.28 1.58-.85 2.21c-.57.63-1.27.98-2.1 1.06c-.48 0-.78.09-.91.28l-2.18 2.4a.75.75 0 0 0-.19.62c.03.23.13.45.31.63c.13.17.33.25.59.23s.46-.15.6-.38L20.64 22c.82-.08 1.58-.35 2.28-.82c.69-.47 1.24-1.07 1.65-1.8s.6-1.52.6-2.36c0-.63-.14-1.32-.43-2.08c.77-.98 1.15-2.08 1.15-3.32c0-.98-.24-1.87-.71-2.69c-.48-.82-1.12-1.46-1.94-1.93s-1.71-.7-2.68-.7c-1.57 0-2.87.57-3.9 1.71c-.87-.43-1.79-.65-2.77-.65c-1.43 0-2.7.44-3.79 1.33s-1.8 2.03-2.11 3.43c-1.14.26-2.07.84-2.79 1.75s-1.09 1.96-1.09 3.15m3.8 7.5c0 .14.02.25.05.32c.08.21.23.36.44.44c.23.1.45.11.68.02c.23-.08.38-.24.45-.45c.1-.22.11-.44.02-.67a.75.75 0 0 0-.46-.46a.937.937 0 0 0-.66 0a.83.83 0 0 0-.47.45c-.03.08-.05.2-.05.35m1.84-2.44c0 .23.11.45.32.67c.43.36.84.31 1.26-.15l2.19-2.44c.15-.17.21-.38.18-.61a.857.857 0 0 0-.31-.57a.818.818 0 0 0-.63-.16c-.24.03-.43.13-.59.29l-2.2 2.38c-.15.21-.22.41-.22.59m.64 4.93c0 .12.03.23.08.32c.09.23.22.38.41.46c.12.05.24.07.37.07c.07 0 .18-.02.32-.06a.78.78 0 0 0 .44-.45c.1-.2.11-.41.02-.64a.712.712 0 0 0-.45-.46a.766.766 0 0 0-.66-.03c-.21.09-.38.25-.49.48c-.02.1-.04.2-.04.31m1.67-2.39v.13c.02.24.12.44.32.6c.14.18.34.26.6.24c.25-.02.45-.15.6-.38l4.22-4.91c.16-.18.22-.39.2-.64a.736.736 0 0 0-.35-.57a.82.82 0 0 0-.6-.19c-.23.02-.42.12-.58.3l-4.22 4.92c-.13.13-.19.29-.19.5m3.63.83a.94.94 0 0 0 0 .62c.09.22.24.38.45.49c.11.05.23.07.36.07c.06 0 .16-.02.3-.06a.75.75 0 0 0 .46-.46c.1-.23.11-.44.03-.66a.737.737 0 0 0-.44-.44a.768.768 0 0 0-.66-.03c-.22.09-.39.25-.5.47m2.4-16.42c.68-.68 1.51-1.02 2.48-1.02c1.01 0 1.86.35 2.56 1.05s1.05 1.56 1.05 2.56c0 .62-.17 1.23-.52 1.82c-.97-.98-2.16-1.46-3.55-1.46h-.31c-.25-1.14-.82-2.12-1.71-2.95"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightShowers(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.25 16.92c0 .66.12 1.28.38 1.88c.25.6.59 1.11 1.02 1.55c.43.44.94.79 1.53 1.06s1.22.42 1.88.45c.12 0 .18-.06.18-.17v-1.34c0-.11-.06-.17-.18-.17c-.85-.04-1.58-.39-2.18-1.03s-.9-1.39-.9-2.24s.28-1.58.84-2.2s1.26-.96 2.1-1.03l.53-.07c.1 0 .15-.05.15-.15l.08-.53c.11-1.09.58-2 1.4-2.73s1.78-1.09 2.88-1.09c1.09 0 2.04.36 2.85 1.08c.82.72 1.28 1.62 1.4 2.7l.07.57c0 .12.06.19.18.19h1.62c.89 0 1.65.32 2.29.96c.64.64.96 1.41.96 2.31c0 .85-.3 1.6-.89 2.24s-1.32.98-2.17 1.03c-.13 0-.19.06-.19.17v1.34c0 .11.06.17.19.17c1.34-.06 2.47-.57 3.39-1.52c.93-.95 1.39-2.09 1.39-3.42c0-.64-.14-1.31-.41-2c.76-.99 1.13-2.09 1.13-3.29c0-.94-.24-1.82-.71-2.63s-1.11-1.45-1.92-1.92c-.81-.47-1.68-.71-2.62-.71c-1.53 0-2.82.56-3.86 1.67a6.03 6.03 0 0 0-2.7-.65c-1.42 0-2.68.44-3.77 1.32s-1.79 2.01-2.1 3.38c-1.13.26-2.05.84-2.76 1.72s-1.08 1.92-1.08 3.1m5.39 6.75c0 .17.05.33.15.48c.1.15.26.27.48.34c.11.06.24.08.37.07c.13-.01.27-.08.4-.2a.96.96 0 0 0 .28-.48l.28-1.01c.06-.25.04-.48-.08-.67a.762.762 0 0 0-.53-.37c-.21-.07-.42-.05-.63.07s-.34.28-.41.5l-.28 1.04c-.02.09-.03.17-.03.23m1.31-4.77c.01.19.06.37.17.52s.27.25.48.28c.18.03.27.05.3.05c.38 0 .63-.22.76-.66l.28-1c.06-.23.04-.45-.08-.66a.862.862 0 0 0-.53-.42c-.22-.06-.44-.04-.64.08s-.33.29-.4.52l-.3 1.05c-.02.12-.04.2-.04.24m1.24 7.9c0 .18.05.34.15.5c.1.16.26.27.48.33c.14.03.23.05.28.05c.09 0 .21-.03.38-.1c.17-.08.3-.27.38-.55l.3-1.01c.06-.25.03-.48-.08-.68a.736.736 0 0 0-.52-.37a.763.763 0 0 0-.64.07a.88.88 0 0 0-.42.51l-.28 1.04c-.02.08-.03.15-.03.21m1.33-4.8c0 .17.05.33.16.49c.11.16.27.27.49.33c.22.07.44.05.63-.05c.2-.1.33-.29.41-.56l.28-1.01c.07-.25.05-.47-.07-.67a.78.78 0 0 0-.53-.36c-.22-.08-.43-.05-.64.06s-.34.29-.41.51l-.28 1.04c-.02.1-.04.17-.04.22m2.85 1.75c-.01.16.03.31.14.46s.26.26.46.33l.25.03c.11.01.24-.02.38-.07c.21-.08.35-.26.42-.54l.28-1.05c.07-.23.05-.45-.07-.64a.885.885 0 0 0-.51-.4c-.25-.06-.47-.03-.67.08s-.32.3-.36.53l-.29 1c-.02.17-.03.26-.03.27m1.35-4.8c0 .17.05.34.16.5c.11.16.27.26.48.3l.25.03c.43 0 .7-.21.81-.62l.28-1.03c.06-.25.03-.48-.08-.68s-.3-.32-.53-.37a.751.751 0 0 0-.64.07a.82.82 0 0 0-.39.53l-.3 1.02c-.02.12-.04.2-.04.25m.34-9.9c.67-.64 1.48-.97 2.45-.97c.98 0 1.81.35 2.49 1.05c.69.7 1.03 1.53 1.03 2.51c0 .64-.16 1.23-.48 1.77c-.96-.96-2.12-1.44-3.49-1.44h-.32c-.24-1.1-.8-2.07-1.68-2.92"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightSleet(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.23 16.99v-.02c0-1.16.36-2.19 1.08-3.09s1.64-1.49 2.74-1.74c.31-1.37 1.01-2.49 2.1-3.37s2.35-1.32 3.77-1.32c.99 0 1.9.22 2.72.66c.5-.53 1.09-.95 1.76-1.25c.67-.3 1.37-.45 2.09-.45c.95 0 1.83.24 2.64.71s1.45 1.11 1.92 1.92s.71 1.69.71 2.64c0 1.23-.38 2.33-1.14 3.29c.29.61.43 1.28.43 2.02c0 .88-.21 1.7-.64 2.45c-.42.75-1 1.36-1.74 1.81c-.73.45-1.54.69-2.42.72c-.13 0-.2-.06-.2-.17v-1.34c0-.12.07-.18.2-.18c.86-.04 1.58-.39 2.18-1.03s.9-1.4.9-2.26c0-.89-.32-1.65-.97-2.29a3.15 3.15 0 0 0-2.31-.96h-1.61c-.12 0-.18-.06-.18-.18l-.08-.59c-.11-1.08-.58-1.99-1.4-2.72c-.82-.73-1.78-1.1-2.86-1.1c-1.1 0-2.05.37-2.86 1.1c-.81.73-1.27 1.64-1.37 2.72l-.08.59c-.03.09-.11.14-.22.14l-.51.04c-.84.1-1.55.45-2.11 1.06s-.84 1.34-.84 2.18v.04c.01.01.02.02.03.02c.01.85.31 1.59.9 2.22c.28.29.59.52.92.67v.02c.38.19.79.31 1.24.34c.11 0 .17.06.17.17v1.34c0 .11-.06.17-.17.17c-.49-.02-.97-.12-1.43-.29c-.76-.27-1.42-.68-1.97-1.26s-.95-1.24-1.18-2a4.95 4.95 0 0 1-.21-1.43m5.5 7.17c0-.03.01-.07.02-.13s.02-.1.02-.12l.09-.58c.07-.24.21-.42.41-.53c.21-.12.43-.15.68-.08c.23.07.39.21.51.41c.11.21.13.42.07.63l-.14.6c-.1.44-.35.66-.76.66c-.03 0-.08 0-.15-.01c-.07-.01-.11-.01-.13-.01a.761.761 0 0 1-.46-.33a.95.95 0 0 1-.16-.51m.74-2.95c0-.24.08-.43.23-.59s.35-.23.59-.23s.43.08.59.23s.23.35.23.59c0 .23-.08.43-.23.58s-.35.23-.59.23c-.23 0-.42-.08-.57-.24a.789.789 0 0 1-.25-.57m1.64 5.98c0-.04.01-.11.04-.23l.13-.59a.87.87 0 0 1 .41-.51c.21-.11.42-.13.63-.07c.23.04.41.17.53.37c.12.2.15.43.08.68l-.13.58c-.11.41-.37.62-.8.62c-.05 0-.13-.01-.24-.04c-.22-.06-.38-.17-.49-.33s-.16-.31-.16-.48m.79-2.91c0-.23.08-.42.23-.58c.16-.15.35-.23.59-.23s.43.08.59.23c.16.15.23.35.23.58c0 .24-.08.43-.23.59c-.16.16-.35.23-.59.23c-.23 0-.42-.08-.58-.24a.785.785 0 0 1-.24-.58m.62-2.27c0-.09.01-.18.03-.26l.23-.9c.07-.24.21-.42.41-.53s.42-.15.64-.08c.24.07.41.2.53.4s.14.41.07.63l-.26.9c-.08.27-.22.46-.41.56c-.19.1-.41.12-.64.06a.657.657 0 0 1-.45-.3a.73.73 0 0 1-.15-.48m2.76 2.15c0-.03 0-.08.01-.14c.01-.06.01-.1.01-.11l.09-.58c.07-.24.21-.42.41-.53c.21-.12.43-.15.67-.08c.23.07.4.21.51.41c.12.21.14.42.07.63l-.14.6c-.1.44-.35.66-.76.66c-.03 0-.08 0-.15-.01c-.07-.01-.11-.01-.13-.01c-.2-.06-.35-.17-.45-.33s-.14-.34-.14-.51m.74-2.95c0-.24.08-.43.23-.59s.35-.23.59-.23s.43.08.59.23c.16.16.23.35.23.59a.784.784 0 0 1-.82.81c-.24 0-.43-.08-.58-.24a.785.785 0 0 1-.24-.57M18.06 9.1c.84.76 1.4 1.74 1.7 2.93h.31c1.38 0 2.55.48 3.52 1.44c.31-.55.47-1.15.47-1.81c0-.98-.35-1.81-1.04-2.5a3.458 3.458 0 0 0-2.51-1.03c-.97.01-1.78.33-2.45.97"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightSleetStorm(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.19 16.88c0 1.12.33 2.12 1 3s1.53 1.47 2.58 1.76l-.66 1.7c-.05.14 0 .22.14.22h2.13l-1.43 4.21h.29l4.36-5.66c.04-.04.04-.09.02-.14c-.02-.05-.07-.07-.14-.07h-2.19l2.49-4.65c.07-.14.03-.22-.14-.22H9.68c-.09 0-.17.05-.23.15l-1.07 2.88c-.72-.18-1.31-.56-1.78-1.16c-.47-.59-.7-1.26-.7-2.02c0-.84.28-1.57.84-2.18c.56-.61 1.26-.97 2.11-1.07l.51-.03c.12 0 .19-.05.22-.14l.07-.59c.11-1.08.56-1.99 1.37-2.72s1.76-1.1 2.86-1.1c1.09 0 2.04.37 2.86 1.1c.82.73 1.29 1.64 1.4 2.72l.08.59c0 .11.06.17.18.17h1.61c.89 0 1.66.32 2.31.96s.97 1.4.97 2.29c0 .87-.3 1.62-.9 2.26s-1.32.98-2.18 1.03c-.13 0-.2.06-.2.18v1.34c0 .11.07.17.2.17c.88-.02 1.69-.26 2.42-.72c.73-.45 1.31-1.06 1.74-1.81s.64-1.57.64-2.45c0-.73-.14-1.4-.43-2.02c.76-.96 1.14-2.06 1.14-3.29c0-.95-.24-1.83-.71-2.64a5.201 5.201 0 0 0-1.92-1.92c-.81-.47-1.69-.71-2.64-.71c-.72 0-1.42.15-2.1.45c-.68.3-1.26.72-1.76 1.25c-.81-.43-1.71-.65-2.72-.65c-1.42 0-2.68.44-3.77 1.32s-1.8 2-2.1 3.37c-1.11.26-2.02.84-2.74 1.74c-.71.92-1.07 1.95-1.07 3.1m7.88 10.22c0 .17.05.33.16.49c.11.16.27.27.49.33c.09.02.17.03.24.03c.43 0 .7-.2.8-.61l.13-.59a.92.92 0 0 0-.08-.68a.762.762 0 0 0-.53-.37a.744.744 0 0 0-.63.07c-.21.12-.34.29-.41.51l-.13.59c-.03.12-.04.2-.04.23m.79-2.9c0 .23.08.42.24.58c.16.16.36.24.58.24c.24 0 .43-.08.59-.23c.16-.16.23-.35.23-.59a.784.784 0 0 0-.82-.81c-.24 0-.43.08-.59.23s-.23.35-.23.58m.62-2.27c-.01.15.03.31.14.47c.1.16.25.26.45.3c.23.06.44.04.64-.06s.33-.29.41-.56l.26-.9c.07-.22.05-.43-.07-.63a.867.867 0 0 0-.53-.4c-.22-.07-.43-.04-.64.08s-.34.3-.41.53l-.22.9c-.02.08-.03.17-.03.27m2.76 2.15c0 .17.05.33.15.48c.1.15.25.26.46.32c.03 0 .08.01.14.02c.06.01.11.02.14.02c.41 0 .66-.22.76-.66l.14-.6c.07-.21.05-.42-.07-.63a.863.863 0 0 0-.51-.41c-.25-.06-.48-.04-.68.08s-.34.29-.41.53l-.09.59c0 .01 0 .05-.01.11c-.01.07-.02.11-.02.15m.74-2.96c0 .23.08.42.24.57c.15.16.34.24.58.24c.24 0 .43-.08.59-.23c.16-.16.23-.35.23-.58c0-.24-.08-.43-.23-.59a.784.784 0 0 0-.59-.23c-.24 0-.43.08-.59.23s-.23.35-.23.59m1.04-12.1c.67-.64 1.48-.97 2.45-.97c.98 0 1.82.34 2.51 1.03c.69.68 1.04 1.52 1.04 2.5c0 .66-.16 1.26-.47 1.81c-.96-.96-2.13-1.44-3.52-1.44h-.31c-.3-1.19-.87-2.17-1.7-2.93"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightSnow(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.23 16.89c0 1.33.47 2.48 1.4 3.44s2.07 1.47 3.4 1.53c.12 0 .18-.06.18-.17v-1.34c0-.11-.06-.17-.18-.17a3.15 3.15 0 0 1-2.19-1.03c-.6-.64-.9-1.39-.9-2.26c0-.83.28-1.55.85-2.17c.57-.61 1.27-.97 2.1-1.07l.53-.04c.13 0 .2-.06.2-.18l.07-.54c.11-1.08.56-1.99 1.37-2.72c.81-.73 1.76-1.1 2.85-1.1s2.04.37 2.86 1.1c.82.73 1.28 1.64 1.4 2.72l.08.57c0 .12.06.19.17.19h1.62c.91 0 1.68.32 2.33.95c.64.63.97 1.4.97 2.28c0 .86-.3 1.61-.91 2.25c-.61.64-1.34.99-2.19 1.04c-.12 0-.19.06-.19.17v1.34c0 .11.06.17.19.17c1.34-.04 2.47-.55 3.4-1.51c.93-.97 1.39-2.12 1.39-3.45c0-.71-.14-1.38-.43-2.01c.79-.96 1.18-2.07 1.18-3.32c0-.95-.24-1.83-.71-2.64S23.96 7.47 23.15 7s-1.68-.7-2.62-.7c-1.55 0-2.85.58-3.89 1.73c-.81-.43-1.71-.65-2.71-.65c-1.41 0-2.67.44-3.76 1.32s-1.8 2-2.11 3.36c-1.11.26-2.02.84-2.74 1.74c-.73.91-1.09 1.94-1.09 3.09m6.39 4.12c0 .22.08.41.24.57c.17.17.36.25.58.25c.23 0 .43-.08.59-.23s.24-.35.24-.59s-.08-.43-.23-.59a.784.784 0 0 0-.59-.23c-.24 0-.43.08-.59.23c-.16.15-.24.35-.24.59m0 3.63c0 .24.08.43.24.58c.16.16.36.24.58.24c.24 0 .43-.08.59-.23c.16-.16.23-.35.23-.59s-.08-.44-.24-.6s-.35-.25-.59-.25c-.23 0-.43.08-.59.25s-.22.36-.22.6m3.19-1.74c0 .23.08.44.25.61s.36.27.58.27c.23 0 .43-.09.6-.26c.17-.17.25-.38.25-.62c0-.22-.08-.41-.25-.58a.816.816 0 0 0-.6-.25c-.22 0-.41.08-.58.25s-.25.36-.25.58m0-3.59c0 .23.08.42.24.58s.36.24.59.24c.24 0 .44-.08.6-.24s.25-.35.25-.59s-.08-.43-.25-.6s-.37-.25-.6-.25c-.23 0-.42.08-.59.25s-.24.38-.24.61m0 7.27c0 .23.08.43.25.6c.17.17.36.25.59.25c.23 0 .43-.08.6-.25c.17-.17.25-.37.25-.6c0-.22-.08-.41-.25-.58a.816.816 0 0 0-.6-.25c-.22 0-.41.08-.58.25c-.17.17-.26.36-.26.58m3.24-5.57c0 .21.08.4.24.57c.17.17.37.25.6.25c.23 0 .43-.08.59-.23c.16-.16.24-.35.24-.59s-.08-.43-.23-.59a.784.784 0 0 0-.59-.23c-.24 0-.44.08-.6.23a.75.75 0 0 0-.25.59m0 3.63c0 .22.08.42.24.58s.36.24.6.24s.43-.08.59-.23c.16-.16.23-.35.23-.59s-.08-.44-.24-.6a.764.764 0 0 0-.59-.25c-.23 0-.43.08-.6.25s-.23.36-.23.6m.98-15.62c.68-.68 1.5-1.01 2.48-1.01s1.81.35 2.5 1.04s1.03 1.53 1.03 2.52c0 .59-.17 1.2-.51 1.84c-.96-.96-2.13-1.44-3.5-1.44h-.31a5.877 5.877 0 0 0-1.69-2.95"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightSnowThunderstorm(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.23 16.88c0 1.12.33 2.12 1 3s1.53 1.47 2.58 1.76l-.66 1.7c-.05.14 0 .22.14.22h2.13L8 27.77h.29l4.36-5.66c.04-.04.04-.09.02-.14c-.02-.05-.07-.07-.14-.07h-2.19l2.49-4.65c.07-.14.03-.22-.14-.22H9.72c-.09 0-.17.05-.23.15l-1.07 2.88c-.72-.18-1.31-.56-1.78-1.16c-.47-.59-.7-1.26-.7-2.02c0-.84.28-1.57.84-2.18c.56-.61 1.27-.97 2.11-1.07l.51-.03c.12 0 .19-.05.22-.14l.08-.59c.11-1.08.56-1.99 1.37-2.72s1.76-1.1 2.86-1.1c1.09 0 2.04.37 2.86 1.1s1.29 1.64 1.4 2.72l.08.59c0 .11.06.17.18.17h1.61c.89 0 1.66.32 2.31.96s.97 1.4.97 2.29c0 .87-.3 1.62-.9 2.26s-1.32.98-2.18 1.03c-.13 0-.2.06-.2.18v1.34c0 .11.07.17.2.17c.88-.02 1.69-.26 2.42-.72c.73-.45 1.31-1.06 1.74-1.81s.64-1.57.64-2.45c0-.73-.14-1.4-.43-2.02c.76-.96 1.14-2.06 1.14-3.29c0-.95-.24-1.83-.71-2.64a5.201 5.201 0 0 0-1.92-1.92c-.81-.47-1.69-.71-2.64-.71c-.72 0-1.42.15-2.1.45c-.68.3-1.26.72-1.76 1.25c-.81-.43-1.71-.65-2.72-.65c-1.42 0-2.68.44-3.77 1.32s-1.8 2-2.1 3.37c-1.11.26-2.02.84-2.74 1.74c-.72.92-1.08 1.95-1.08 3.1m9.59 6.08c0 .24.08.44.24.59c.16.16.36.24.58.24c.24 0 .44-.08.61-.24s.25-.36.25-.59c0-.24-.08-.44-.25-.61s-.37-.26-.61-.26c-.22 0-.41.09-.58.26s-.24.37-.24.61m0-3.64c0 .24.08.43.24.58c.16.16.36.24.58.24c.24 0 .45-.08.61-.23s.25-.35.25-.59c0-.23-.08-.43-.25-.6s-.37-.25-.61-.25c-.23 0-.42.08-.58.25s-.24.37-.24.6m0 7.31c0 .22.08.41.24.57c.17.17.36.25.58.25c.24 0 .44-.08.61-.24c.17-.16.25-.35.25-.59s-.08-.44-.25-.61a.832.832 0 0 0-.61-.26c-.22 0-.41.09-.58.26c-.16.18-.24.39-.24.62m3.23-5.61c0 .24.08.44.25.6s.36.25.6.25c.23 0 .43-.08.59-.25c.16-.17.24-.37.24-.6c0-.22-.08-.42-.24-.58a.791.791 0 0 0-.59-.24c-.23 0-.43.08-.6.24s-.25.35-.25.58m0 3.64c0 .23.08.42.24.58c.16.16.36.24.6.24s.43-.08.59-.24c.16-.16.23-.35.23-.59s-.08-.43-.23-.59a.784.784 0 0 0-.59-.23c-.24 0-.44.08-.6.23c-.16.16-.24.36-.24.6m1.01-15.64c.67-.64 1.48-.97 2.45-.97c.98 0 1.82.34 2.51 1.03c.69.68 1.04 1.52 1.04 2.5c0 .66-.16 1.26-.47 1.81c-.96-.96-2.13-1.44-3.52-1.44h-.31c-.3-1.19-.87-2.17-1.7-2.93"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightSnowWind(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.23 16.89c0 1.33.47 2.48 1.4 3.44s2.07 1.47 3.4 1.53c.12 0 .18-.06.18-.17v-1.34c0-.11-.06-.17-.18-.17a3.15 3.15 0 0 1-2.19-1.03c-.6-.64-.9-1.39-.9-2.26c0-.83.28-1.55.85-2.17c.57-.61 1.27-.97 2.1-1.07l.53-.04c.13 0 .2-.06.2-.18l.07-.54c.11-1.08.56-1.99 1.37-2.72c.81-.73 1.76-1.1 2.85-1.1s2.04.37 2.86 1.1c.82.73 1.28 1.64 1.4 2.72l.08.57c0 .12.06.19.17.19h1.62c.91 0 1.68.32 2.33.95c.64.63.97 1.4.97 2.28c0 .86-.3 1.61-.91 2.25c-.61.64-1.34.99-2.19 1.04c-.12 0-.19.06-.19.17v1.34c0 .11.06.17.19.17c1.34-.04 2.47-.55 3.4-1.51c.93-.97 1.39-2.12 1.39-3.45c0-.71-.14-1.38-.43-2.01c.79-.96 1.18-2.07 1.18-3.32c0-.95-.24-1.83-.71-2.64S23.96 7.47 23.15 7s-1.68-.7-2.62-.7c-1.55 0-2.85.58-3.89 1.73c-.81-.43-1.71-.65-2.71-.65c-1.41 0-2.67.44-3.76 1.32s-1.8 2-2.11 3.36c-1.11.26-2.02.84-2.74 1.74c-.73.91-1.09 1.94-1.09 3.09m5.49 7.72c0 .21.08.4.24.57c.18.16.37.24.58.24c.24 0 .43-.08.59-.23c.16-.16.23-.35.23-.58c0-.24-.08-.43-.24-.59a.784.784 0 0 0-.59-.23a.84.84 0 0 0-.59.23c-.14.15-.22.35-.22.59m.86-3.63c0 .24.08.44.24.61c.16.17.35.25.59.25c.23 0 .43-.08.59-.25c.16-.17.24-.37.24-.61c0-.23-.08-.42-.24-.58a.791.791 0 0 0-.59-.24c-.23 0-.43.08-.59.24c-.16.16-.24.36-.24.58m1.89 5.58c0 .22.08.41.24.57c.17.17.36.25.58.25c.24 0 .44-.08.6-.23c.17-.16.25-.35.25-.59s-.08-.44-.25-.6a.816.816 0 0 0-.6-.25c-.22 0-.41.08-.58.25c-.15.17-.24.37-.24.6m.86-3.66c0 .22.08.42.25.6c.16.16.35.24.57.24c.24 0 .44-.08.61-.24s.25-.36.25-.6c0-.23-.08-.43-.25-.6s-.37-.25-.61-.25c-.23 0-.42.08-.58.25s-.24.37-.24.6m.43-3.6c0 .23.08.42.24.58s.36.24.58.24c.24 0 .44-.08.6-.24c.17-.16.25-.35.25-.59c0-.23-.08-.43-.25-.59s-.37-.24-.6-.24c-.22 0-.42.08-.58.24s-.24.37-.24.6m2.37 5.31c0 .21.08.4.23.57c.17.16.38.24.6.24c.24 0 .43-.08.59-.23c.16-.16.23-.35.23-.58c0-.24-.08-.43-.23-.59a.784.784 0 0 0-.59-.23a.8.8 0 0 0-.83.82m.85-3.63c0 .24.08.44.24.61c.16.17.36.25.59.25c.23 0 .43-.08.59-.25c.16-.17.24-.37.24-.61c0-.23-.08-.42-.24-.58a.791.791 0 0 0-.59-.24c-.23 0-.43.08-.59.24s-.24.36-.24.58m1.05-11.96c.68-.68 1.5-1.01 2.48-1.01s1.81.35 2.5 1.04s1.03 1.53 1.03 2.52c0 .59-.17 1.2-.51 1.84c-.96-.96-2.13-1.44-3.5-1.44h-.31a5.877 5.877 0 0 0-1.69-2.95"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightSprinkle(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.22 16.89c0 1.33.46 2.48 1.39 3.44s2.06 1.47 3.41 1.53c.11 0 .17-.06.17-.17v-1.37c0-.13-.06-.19-.17-.19c-.88-.06-1.61-.41-2.21-1.03c-.6-.62-.9-1.36-.9-2.21c0-.84.28-1.58.85-2.2c.57-.62 1.27-.97 2.11-1.04l.52-.07c.12 0 .19-.06.19-.19l.07-.5c.11-1.08.57-1.99 1.38-2.72s1.77-1.1 2.87-1.1c1.09 0 2.05.36 2.86 1.09c.81.73 1.28 1.64 1.4 2.72l.06.57c0 .12.06.18.19.18h1.6c.91 0 1.68.32 2.32.95c.64.63.97 1.4.97 2.28c0 .85-.3 1.59-.89 2.21c-.59.62-1.33.97-2.19 1.03c-.14 0-.21.06-.21.19v1.37c0 .11.07.17.21.17c1.33-.04 2.46-.55 3.39-1.51c.93-.97 1.39-2.12 1.39-3.45c0-.74-.14-1.41-.43-2.01c.79-.96 1.18-2.07 1.18-3.36c0-.94-.24-1.82-.71-2.63s-1.11-1.45-1.92-1.92c-.81-.47-1.68-.71-2.62-.71c-1.52 0-2.83.58-3.93 1.75c-.83-.38-1.72-.59-2.67-.59c-1.41 0-2.67.44-3.76 1.31s-1.8 2-2.1 3.37c-1.11.26-2.02.84-2.74 1.74c-.72.89-1.08 1.92-1.08 3.07m5.93.88c0 .38.14.7.43.98c.28.27.62.41 1.02.41s.73-.13 1-.4s.41-.6.41-.98c0-.26-.12-.6-.35-1.02c-.23-.42-.45-.76-.66-1c-.02-.02-.08-.09-.18-.2c-.1-.11-.17-.19-.21-.24l-.36.4c-.28.3-.53.65-.75 1.05c-.23.4-.35.74-.35 1m2.99 3.99c0 .63.23 1.18.69 1.64c.46.46 1.01.69 1.65.69c.64 0 1.2-.23 1.66-.69c.46-.46.69-1.01.69-1.64c0-.27-.08-.59-.23-.97c-.16-.38-.34-.72-.56-1.04c-.46-.59-.89-1.09-1.29-1.49c-.06-.04-.14-.13-.26-.24l-.59.58c-.44.42-.85.95-1.21 1.56c-.37.62-.55 1.15-.55 1.6m1.48-6.7c0 .27.09.49.28.67s.43.27.71.27c.26 0 .48-.09.66-.27s.27-.4.27-.67c0-.41-.31-.94-.93-1.61l-.25.26c-.19.2-.36.43-.51.7c-.16.26-.23.48-.23.65m3.39-6.04c.67-.66 1.5-.99 2.48-.99s1.81.34 2.49 1.02s1.03 1.51 1.03 2.48c0 .63-.17 1.25-.51 1.85c-1-.96-2.17-1.44-3.51-1.44h-.29c-.28-1.18-.85-2.15-1.69-2.92"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightStormShowers(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.25 16.86c0 1.1.33 2.09 1 2.98c.67.88 1.52 1.48 2.57 1.8l-.65 1.66c-.04.14 0 .21.14.21h2.12L8.29 27.6h.29l4.08-5.49c.04-.04.04-.09.01-.14c-.03-.05-.08-.07-.15-.07h-2.17l2.47-4.67c.07-.14.03-.22-.13-.22H9.73c-.09 0-.16.05-.19.14l-1.11 2.93c-.71-.18-1.3-.57-1.78-1.17c-.47-.6-.71-1.27-.71-2.02c0-.84.28-1.57.85-2.19s1.27-.97 2.1-1.05l.52-.07c.13 0 .2-.06.2-.17l.07-.52c.11-1.09.56-2 1.37-2.72c.81-.73 1.76-1.09 2.86-1.09c1.09 0 2.05.36 2.86 1.09c.81.73 1.28 1.64 1.4 2.72l.07.58c0 .11.06.17.18.17h1.62c.91 0 1.68.32 2.33.95c.64.63.97 1.4.97 2.28c0 .86-.3 1.6-.9 2.23c-.6.63-1.33.97-2.2 1.04c-.12 0-.19.06-.19.18v1.38c0 .11.06.17.19.17c1.33-.04 2.46-.55 3.39-1.52c.93-.97 1.39-2.12 1.39-3.47c0-.73-.14-1.39-.41-2c.76-1 1.14-2.1 1.14-3.29c0-.71-.14-1.39-.42-2.04c-.28-.65-.66-1.2-1.12-1.67s-1.03-.84-1.68-1.12a5.175 5.175 0 0 0-4.15.03c-.67.3-1.26.72-1.74 1.26a5.62 5.62 0 0 0-2.7-.66c-1.42 0-2.68.44-3.77 1.31s-1.8 2-2.11 3.37c-1.11.26-2.02.83-2.74 1.73s-1.07 1.91-1.07 3.06m8.17 9.87c0 .18.05.35.16.51c.11.17.26.27.46.3c.02 0 .05 0 .08.01s.07.01.09.01h.08c.43-.03.69-.23.8-.61l.29-1.06a.78.78 0 0 0-.09-.63c-.12-.2-.3-.34-.53-.41a.817.817 0 0 0-.63.08c-.2.12-.34.3-.41.53l-.25 1c-.03.18-.05.26-.05.27m1.34-4.77c0 .15.05.3.15.45c.1.15.26.26.46.34c.22.08.43.06.63-.05s.33-.29.4-.53l.3-1.04c.06-.25.04-.48-.08-.68s-.29-.32-.53-.37c-.22-.07-.44-.05-.64.07s-.34.29-.42.53l-.25 1.02c0 .02 0 .05-.01.08s-.01.07-.01.09zm2.83 1.62c0 .19.05.36.16.52c.11.16.26.27.47.32c.16.03.25.05.27.05c.39 0 .65-.2.77-.6l.24-1.06c.07-.22.05-.43-.06-.63a.822.822 0 0 0-.5-.41a.882.882 0 0 0-.68.09c-.21.12-.33.3-.38.53l-.28.99c0 .05-.01.12-.01.2m1.36-4.75c0 .16.05.32.16.47c.11.15.27.27.49.34l.26.03c.12 0 .24-.03.38-.08c.19-.1.33-.27.39-.52l.29-1.04a.78.78 0 0 0-.09-.63c-.12-.2-.3-.34-.53-.41c-.23-.07-.44-.05-.64.07s-.33.29-.4.53l-.28 1.02c-.02.09-.03.16-.03.22m.11-9.85c.66-.64 1.48-.96 2.45-.96c.98 0 1.82.35 2.5 1.04c.69.69 1.03 1.53 1.03 2.51c0 .63-.16 1.22-.49 1.78c-.99-.96-2.15-1.44-3.49-1.44h-.32c-.3-1.16-.86-2.14-1.68-2.93"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightThunderstorm(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.28 16.89c0 1.11.33 2.11.99 2.98s1.52 1.46 2.56 1.75l-.64 1.68c-.05.14 0 .22.14.22h2.12l-1.04 4.19h.28l3.97-5.62c.04-.04.04-.09.01-.14c-.03-.05-.08-.07-.15-.07h-2.17l2.47-4.61c.07-.14.02-.22-.14-.22H9.74c-.09 0-.16.05-.23.14l-1.07 2.87c-.71-.17-1.3-.56-1.77-1.14s-.7-1.26-.7-2.02c0-.83.28-1.55.84-2.16s1.26-.96 2.1-1.06l.53-.04c.12 0 .18-.06.18-.18l.07-.53c.07-.71.3-1.35.69-1.94c.39-.58.9-1.04 1.52-1.37s1.29-.5 2.01-.5c1.08 0 2.03.37 2.84 1.1c.81.73 1.27 1.63 1.39 2.71l.08.56c0 .12.06.19.17.19h1.62c.89 0 1.65.32 2.3.96s.97 1.39.97 2.27c0 .86-.3 1.61-.9 2.25s-1.33.97-2.18 1.02c-.13 0-.2.06-.2.18v1.34c0 .11.07.17.2.17a4.62 4.62 0 0 0 2.4-.72c.73-.45 1.31-1.05 1.72-1.8s.63-1.56.63-2.43c0-.73-.14-1.4-.42-2.01c.78-.93 1.17-2.03 1.17-3.31c0-.71-.14-1.38-.42-2.02c-.28-.64-.65-1.2-1.12-1.67a5.27 5.27 0 0 0-1.67-1.12c-.64-.28-1.32-.42-2.02-.42c-1.54 0-2.83.58-3.86 1.73c-.81-.43-1.71-.65-2.7-.65c-1.41 0-2.66.44-3.75 1.31s-1.79 1.99-2.1 3.35c-1.1.26-2.01.83-2.73 1.73s-1.05 1.9-1.05 3.05m7.93 9.88c0 .16.05.32.15.46s.25.25.45.31l.25.03c.42 0 .68-.2.8-.6l2.43-8.89a.83.83 0 0 0-.07-.64a.815.815 0 0 0-.49-.4a.814.814 0 0 0-.65.06c-.2.11-.34.28-.4.5l-2.45 8.9c-.01.17-.02.26-.02.27m4.14-3.03c0 .4.21.67.62.8l.26.03c.11 0 .23-.02.35-.08c.2-.09.34-.27.42-.55l1.64-5.85a.79.79 0 0 0-.08-.64a.827.827 0 0 0-.51-.4a.814.814 0 0 0-1.04.56l-1.62 5.89c-.02.14-.04.22-.04.24m1.67-14.7c.68-.64 1.5-.96 2.48-.96c.97 0 1.8.34 2.48 1.02c.69.68 1.03 1.51 1.03 2.48c0 .63-.17 1.25-.51 1.85c-.96-.96-2.12-1.44-3.48-1.44h-.32a6.033 6.033 0 0 0-1.68-2.95"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rain(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.64 16.91c0-1.15.36-2.17 1.08-3.07a4.82 4.82 0 0 1 2.73-1.73c.31-1.36 1.02-2.48 2.11-3.36s2.34-1.31 3.75-1.31c1.38 0 2.6.43 3.68 1.28c1.08.85 1.78 1.95 2.1 3.29h.32c.89 0 1.72.22 2.48.65s1.37 1.03 1.81 1.78c.44.75.67 1.58.67 2.47c0 .88-.21 1.69-.63 2.44c-.42.75-1 1.35-1.73 1.8c-.73.45-1.53.69-2.4.71c-.13 0-.2-.06-.2-.17v-1.33c0-.12.07-.18.2-.18c.85-.04 1.58-.38 2.18-1.02s.9-1.39.9-2.26s-.33-1.62-.98-2.26s-1.42-.96-2.31-.96h-1.61c-.12 0-.18-.06-.18-.17l-.08-.58a4.076 4.076 0 0 0-1.39-2.71c-.82-.73-1.76-1.09-2.85-1.09s-2.05.36-2.85 1.09a4.02 4.02 0 0 0-1.36 2.71l-.07.53c0 .12-.07.19-.2.19l-.53.03c-.83.1-1.53.46-2.1 1.07s-.85 1.33-.85 2.16c0 .87.3 1.62.9 2.26s1.33.98 2.18 1.02c.11 0 .17.06.17.18v1.33c0 .11-.06.17-.17.17c-1.34-.06-2.47-.57-3.4-1.53s-1.37-2.1-1.37-3.43m5.35 6.69c0-.04.01-.11.04-.2l1.63-5.77a.837.837 0 0 1 1.02-.56c.24.04.42.17.54.37c.12.2.15.42.08.67l-1.63 5.73c-.12.43-.4.64-.82.64c-.04 0-.07-.01-.11-.02c-.06-.02-.09-.03-.1-.03a.831.831 0 0 1-.49-.33a.895.895 0 0 1-.16-.5m2.62 2.81l2.44-8.77c.04-.19.14-.34.3-.44c.16-.1.32-.15.49-.15c.09 0 .18.01.27.03c.22.06.38.19.49.39c.11.2.13.41.07.64l-2.43 8.78c-.04.17-.13.31-.29.43c-.16.12-.32.18-.51.18c-.09 0-.18-.02-.25-.05c-.2-.05-.37-.18-.52-.39c-.11-.18-.13-.39-.06-.65m4.13-2.79c0-.04.01-.11.04-.23l1.63-5.77a.83.83 0 0 1 .3-.44c.15-.1.3-.15.46-.15c.08 0 .17.01.26.03c.21.06.36.16.46.31c.1.15.15.31.15.47c0 .03-.01.08-.02.14s-.02.1-.02.12l-1.63 5.73c-.04.19-.13.35-.28.46s-.32.17-.51.17l-.24-.05a.809.809 0 0 1-.46-.32a.916.916 0 0 1-.14-.47"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RainMix(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.65 16.96c0 1.32.47 2.46 1.4 3.41c.93.96 2.06 1.46 3.38 1.5c.12 0 .18-.06.18-.17v-1.33c0-.12-.06-.18-.18-.18c-.84-.04-1.57-.38-2.17-1.02s-.91-1.37-.91-2.22c0-.84.28-1.57.85-2.19c.57-.62 1.26-.97 2.1-1.04l.53-.07c.12 0 .19-.06.19-.18l.07-.5c.1-1.09.55-2.01 1.36-2.75s1.76-1.11 2.86-1.11c1.08 0 2.03.37 2.84 1.1c.81.73 1.28 1.63 1.4 2.71l.07.58c0 .12.06.18.19.18h1.6c.9 0 1.67.32 2.32.97c.64.64.97 1.41.97 2.3c0 .84-.3 1.58-.9 2.22c-.6.63-1.33.97-2.18 1.02c-.13 0-.2.06-.2.18v1.33c0 .11.07.17.2.17c1.33-.04 2.46-.54 3.38-1.5s1.38-2.09 1.38-3.42c0-.89-.22-1.72-.67-2.48a4.884 4.884 0 0 0-1.81-1.8c-.76-.44-1.59-.66-2.48-.66h-.31A5.885 5.885 0 0 0 18 8.72a5.763 5.763 0 0 0-3.68-1.28c-1.41 0-2.66.44-3.75 1.31s-1.79 1.99-2.1 3.35c-1.13.29-2.04.88-2.75 1.77s-1.07 1.93-1.07 3.09m5.4 7.02c0 .17.05.34.16.51c.11.17.27.28.47.35c.23.07.44.06.64-.04c.19-.09.33-.28.39-.56l.14-.61a.853.853 0 0 0-.61-1.03c-.22-.07-.44-.04-.64.08s-.34.3-.4.53l-.14.59c0 .03-.01.09-.01.18m.76-2.9c0 .21.08.4.25.57c.16.17.34.25.56.25c.24 0 .44-.08.6-.24c.16-.16.24-.35.24-.59c0-.23-.08-.43-.24-.59a.814.814 0 0 0-.6-.24c-.23 0-.42.08-.58.23c-.15.18-.23.38-.23.61m.61-2.27c-.01.16.03.31.14.45c.1.15.26.25.48.32c.21.06.41.04.62-.07s.34-.28.41-.51l.28-.9c.07-.24.05-.46-.07-.65a.913.913 0 0 0-.54-.39a.737.737 0 0 0-.63.07a.85.85 0 0 0-.41.5l-.24.92c0 .02-.01.06-.02.12c-.01.05-.02.1-.02.14m1.17 8.29c0 .18.05.34.15.5c.1.16.26.27.48.33c.08.02.17.03.25.03c.43 0 .69-.2.79-.61l.14-.59a.92.92 0 0 0-.08-.68a.765.765 0 0 0-.52-.37a.744.744 0 0 0-.63.07c-.21.12-.34.29-.41.51l-.14.59c-.02.09-.03.16-.03.22m.77-2.9c0 .22.08.41.25.58c.16.16.35.24.57.24c.24 0 .43-.08.59-.23c.16-.16.23-.35.23-.59a.784.784 0 0 0-.82-.81c-.24 0-.43.08-.59.23s-.23.35-.23.58m.63-2.27c-.01.15.03.31.13.47c.1.16.25.26.45.3c.23.06.44.04.64-.06s.33-.29.41-.56l.27-.9c.07-.22.05-.43-.07-.63a.867.867 0 0 0-.53-.4a.766.766 0 0 0-.64.08c-.21.12-.34.3-.41.53l-.23.9c-.01.08-.02.17-.02.27m2.76 2.15c0 .16.05.32.15.48c.1.16.26.27.46.33c.03 0 .08.01.14.02c.06.01.1.02.14.02c.41 0 .66-.22.76-.66l.14-.6c.07-.21.05-.42-.07-.63a.809.809 0 0 0-.51-.41c-.25-.06-.48-.04-.68.08s-.34.29-.41.53l-.09.59c0 .02-.01.07-.02.12s-.01.09-.01.13m.74-2.96c0 .22.08.42.25.57c.15.16.34.24.57.24c.24 0 .43-.08.59-.23s.23-.35.23-.58c0-.24-.08-.43-.23-.59s-.35-.23-.59-.23s-.43.08-.59.23c-.15.16-.23.35-.23.59m.61-2.31c0 .17.05.33.16.48c.11.15.27.26.49.32c.02 0 .06.01.12.02s.11.02.14.02c.1 0 .22-.03.36-.09c.21-.11.35-.29.41-.52l.24-.9c.06-.23.04-.44-.08-.63a.827.827 0 0 0-.51-.4a.79.79 0 0 0-.64.06c-.19.11-.33.27-.39.51l-.28.91c0 .02-.01.06-.02.12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RainWind(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.63 16.93c0 1.12.33 2.11.98 2.99c.65.87 1.5 1.47 2.55 1.79c.09.02.17-.01.24-.08l1.16-1.43c-.89 0-1.65-.32-2.28-.96c-.63-.64-.95-1.41-.95-2.31c0-.84.28-1.58.84-2.2s1.26-.97 2.1-1.04l.53-.07c.11 0 .16-.04.16-.13l.08-.55c.12-1.1.59-2.01 1.39-2.73s1.75-1.08 2.85-1.08s2.06.36 2.87 1.09c.82.73 1.27 1.64 1.37 2.72l.07.58c.02.11.1.17.22.17h1.62c.9 0 1.67.32 2.3.95s.95 1.39.95 2.29c0 .83-.28 1.56-.84 2.18s-1.25.98-2.07 1.08c-.12 0-.28.02-.49.06c-.19.02-.33.09-.41.23l-2.36 2.79a.78.78 0 0 0-.16.63c.03.24.14.43.31.57c.11.12.29.19.56.19c.26 0 .47-.12.61-.35l2.12-2.44c1.24-.13 2.29-.66 3.15-1.61c.86-.95 1.28-2.06 1.28-3.33c0-.67-.13-1.32-.39-1.93a4.96 4.96 0 0 0-1.05-1.58a4.96 4.96 0 0 0-1.58-1.05a4.86 4.86 0 0 0-1.93-.39h-.32c-.33-1.32-1.04-2.41-2.12-3.26s-2.32-1.27-3.72-1.27c-.93 0-1.81.2-2.63.6c-.82.4-1.51.95-2.08 1.66s-.94 1.52-1.13 2.42c-1.12.25-2.04.82-2.75 1.72c-.7.89-1.05 1.92-1.05 3.08m3.38 8.02c0 .06.02.16.06.3c.09.21.23.36.44.44c.22.1.44.11.67.02a.76.76 0 0 0 .46-.45c.1-.22.11-.43.02-.65a.728.728 0 0 0-.46-.43a.745.745 0 0 0-.65-.03a.85.85 0 0 0-.46.47a.73.73 0 0 0-.08.33m1.85-2.44v.1c.02.23.12.41.3.56c.23.13.43.19.62.19c.22 0 .43-.11.61-.33l2.32-2.77c.14-.17.21-.39.2-.66a.795.795 0 0 0-.28-.53c-.16-.14-.33-.22-.52-.22c-.06 0-.1 0-.14.01c-.23.04-.42.15-.56.33l-2.36 2.77c-.13.16-.19.34-.19.55m.77 4.72c0 .12.03.23.08.32c.08.21.23.37.44.47c.11.05.22.07.33.07c.12 0 .23-.02.31-.07c.23-.09.39-.23.47-.41c.1-.22.11-.44.02-.67a.712.712 0 0 0-.45-.46a.823.823 0 0 0-.67-.02c-.23.09-.38.24-.45.45c-.06.09-.08.2-.08.32m1.67-2.35v.11c.02.22.13.4.31.55c.18.15.37.22.55.22c.23 0 .43-.11.63-.33l4.35-5.24c.11-.12.17-.3.17-.52v-.12c-.02-.23-.12-.4-.27-.53s-.33-.2-.52-.2h-.13c-.23.01-.42.12-.55.31l-4.35 5.2c-.14.18-.19.36-.19.55m3.51 1.15c0 .09.02.19.06.3c.09.22.24.38.46.47c.14.04.24.06.31.06c.14 0 .26-.03.34-.08c.22-.09.38-.23.46-.42c.1-.17.11-.39.02-.67a.763.763 0 0 0-.44-.44l-.36-.09c-.09.02-.19.04-.32.07c-.22.08-.37.23-.45.44c-.05.13-.08.25-.08.36"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Raindrop(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.81 15.25c0 .92.23 1.78.7 2.57s1.1 1.43 1.9 1.9c.8.47 1.66.71 2.59.71c.93 0 1.8-.24 2.61-.71a5.3 5.3 0 0 0 1.92-1.9c.47-.8.71-1.65.71-2.57c0-.6-.17-1.31-.52-2.14c-.35-.83-.77-1.6-1.26-2.3c-.44-.57-.96-1.2-1.56-1.88c-.6-.68-1.65-1.73-1.89-1.97l-1.28 1.29c-.62.6-1.22 1.29-1.79 2.08c-.57.79-1.07 1.64-1.49 2.55c-.44.91-.64 1.7-.64 2.37"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Raindrops(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.01 12.23c0-.26.13-.59.38-1.01c.25-.42.5-.77.73-1.04c.06-.07.14-.17.23-.28s.15-.17.16-.18l.37.43c.28.31.53.66.76 1.07c.23.41.35.74.35 1.01c0 .41-.14.77-.43 1.06c-.28.29-.63.44-1.05.44c-.41 0-.77-.15-1.06-.44c-.29-.3-.44-.65-.44-1.06m3.12 4.15c0-.29.08-.62.24-1.01c.16-.38.36-.74.6-1.06c.46-.61.89-1.12 1.31-1.53c.04-.03.13-.11.26-.24l.25.24c.39.37.83.88 1.32 1.52c.26.34.46.7.62 1.08s.24.71.24 1c0 .69-.23 1.26-.7 1.73s-1.05.7-1.73.7c-.68 0-1.25-.24-1.72-.71s-.69-1.05-.69-1.72m1.52-6.9c0-.43.33-1 1-1.7l.25.28c.19.22.36.46.51.74c.15.27.23.5.23.68c0 .28-.1.5-.29.69c-.19.18-.42.28-.7.28c-.29 0-.53-.09-.72-.28a.976.976 0 0 1-.28-.69"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Refresh(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.77 15.53c0 .94.24 1.82.71 2.62s1.11 1.44 1.91 1.9c.8.47 1.67.7 2.61.7c.96 0 1.83-.23 2.63-.69c.8-.46 1.43-1.09 1.89-1.89c.46-.8.69-1.68.69-2.63c0-.24-.08-.44-.24-.61a.764.764 0 0 0-.59-.25c-.23 0-.43.08-.6.25c-.17.17-.25.37-.25.61c0 .98-.35 1.82-1.04 2.51c-.69.69-1.53 1.04-2.51 1.04c-.97 0-1.79-.35-2.47-1.04c-.68-.69-1.02-1.53-1.02-2.51c0-.85.26-1.62.79-2.31s1.14-1.06 1.84-1.1l-.38.37c-.16.18-.24.37-.24.58c0 .22.08.42.24.6c.36.35.77.35 1.21 0l1.84-1.82c.16-.12.24-.33.24-.62c0-.26-.08-.45-.24-.57L14.97 8.8c-.18-.16-.37-.24-.57-.24c-.25 0-.46.08-.63.25c-.17.17-.25.37-.25.6c0 .24.08.45.24.61l.38.36c-1.25.22-2.29.82-3.12 1.8s-1.25 2.09-1.25 3.35"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RefreshAlt(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.78 14.91c0 .79.19 1.51.57 2.17c.38.66.9 1.19 1.57 1.57c.67.38 1.39.58 2.18.58c.19 0 .35-.07.48-.22c.13-.14.2-.31.2-.51c0-.19-.07-.35-.2-.48s-.29-.19-.49-.19c-.81 0-1.5-.28-2.07-.85c-.57-.57-.85-1.26-.85-2.07c0-.78.27-1.45.8-2.02s1.16-.86 1.88-.86l-.33.32c-.15.15-.22.31-.21.49c0 .18.07.34.2.48c.13.14.29.21.49.21s.37-.07.51-.21l1.51-1.5c.13-.11.2-.27.2-.51c0-.22-.07-.38-.2-.47l-1.51-1.53c-.13-.14-.29-.21-.49-.21s-.36.07-.5.21s-.21.3-.21.5c0 .21.07.38.22.51l.3.28c-1.15.08-2.11.53-2.89 1.35c-.77.82-1.16 1.81-1.16 2.96"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sandstorm(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.1 16.97c0 .24.09.45.28.62c.16.19.37.28.63.28H18.7c.29 0 .53.1.73.3c.2.2.3.45.3.74s-.1.53-.3.72c-.2.19-.44.29-.74.29c-.29 0-.54-.1-.73-.29a.76.76 0 0 0-.6-.26c-.25 0-.46.09-.64.26s-.27.38-.27.61c0 .25.09.46.28.63c.56.55 1.22.83 1.96.83c.78 0 1.45-.27 2.01-.81c.56-.54.83-1.19.83-1.97s-.28-1.44-.84-2c-.56-.56-1.23-.84-2-.84H4.01a.9.9 0 0 0-.64.26c-.18.17-.27.38-.27.63m0-3.28c0 .23.09.43.28.61c.17.18.38.26.63.26h20.04c.78 0 1.45-.27 2.01-.82c.56-.54.84-1.2.84-1.97s-.28-1.44-.84-1.99s-1.23-.83-2.01-.83c-.77 0-1.42.27-1.95.8c-.18.16-.27.38-.27.67c0 .26.09.47.26.63c.17.16.38.24.63.24c.24 0 .45-.08.63-.24c.19-.21.42-.31.7-.31c.29 0 .53.1.73.3c.2.2.3.44.3.73s-.1.53-.3.72c-.2.19-.44.29-.73.29H4.01a.908.908 0 0 0-.91.91m1.57 6.92c0 .24.08.43.24.58c.16.16.36.24.58.24c.24 0 .45-.08.62-.23s.25-.35.25-.59c0-.23-.09-.43-.26-.6a.838.838 0 0 0-.61-.25c-.22 0-.42.08-.58.25c-.16.17-.24.37-.24.6m.95-9.93c0 .24.08.43.24.58c.16.16.36.24.58.24c.24 0 .45-.08.61-.23c.17-.16.25-.35.25-.59c0-.23-.08-.43-.25-.6a.84.84 0 0 0-.6-.25c-.22 0-.42.08-.58.25c-.17.17-.25.37-.25.6m3.03-2.31c0 .24.08.43.24.58c.16.17.36.25.58.25c.24 0 .45-.08.62-.23c.17-.16.25-.35.25-.59c0-.23-.09-.43-.26-.6a.807.807 0 0 0-.61-.25c-.22 0-.42.08-.58.25c-.16.16-.24.36-.24.59m.09 11.6c0 .23.08.43.25.58c.16.16.35.24.57.24c.24 0 .45-.08.62-.23c.17-.16.25-.35.25-.59c0-.23-.09-.43-.26-.6a.838.838 0 0 0-.61-.25c-.22 0-.42.08-.58.25c-.16.17-.24.37-.24.6m4.18-9.83c0 .24.08.43.24.58c.16.16.36.24.58.24c.24 0 .45-.08.62-.23s.25-.35.25-.59c0-.23-.09-.43-.26-.6a.838.838 0 0 0-.61-.25c-.23 0-.42.08-.58.25c-.16.17-.24.37-.24.6m.32 10.99c0 .23.08.42.25.58c.16.16.35.24.57.24c.24 0 .45-.08.62-.23c.17-.16.25-.35.25-.59c0-.23-.09-.43-.26-.6a.838.838 0 0 0-.61-.25c-.22 0-.42.08-.58.25s-.24.37-.24.6m4.48-11.15c0 .23.08.42.24.57c.17.17.36.25.58.25c.24 0 .45-.08.62-.23c.17-.16.25-.35.25-.59c0-.23-.09-.43-.26-.6a.838.838 0 0 0-.61-.25c-.22 0-.41.08-.58.25c-.15.17-.24.37-.24.6m5.09 7.06c0 .24.08.43.24.58c.16.16.36.24.58.24c.24 0 .45-.08.61-.23c.17-.16.25-.35.25-.59c0-.23-.08-.43-.25-.6a.822.822 0 0 0-.61-.25c-.23 0-.42.08-.58.25c-.15.17-.24.37-.24.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Showers(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.6 16.93c0-1.16.36-2.18 1.09-3.08c.72-.9 1.65-1.48 2.78-1.73c.29-1.38.98-2.5 2.07-3.39S12.88 7.4 14.3 7.4c1.39 0 2.63.43 3.72 1.28c1.08.85 1.79 1.95 2.12 3.3h.34c.9 0 1.73.22 2.48.66c.76.44 1.35 1.04 1.79 1.8c.43.76.65 1.59.65 2.49c0 1.34-.46 2.48-1.37 3.44c-.92.96-2.04 1.46-3.37 1.5c-.12 0-.18-.06-.18-.17v-1.34c0-.11.06-.17.18-.17c.84-.07 1.57-.42 2.17-1.05s.9-1.37.9-2.22c0-.89-.32-1.66-.96-2.31c-.64-.64-1.4-.97-2.29-.97h-1.63c-.12 0-.19-.06-.22-.18l-.07-.57c-.07-.71-.3-1.36-.7-1.94a4.257 4.257 0 0 0-3.55-1.85c-1.1 0-2.05.36-2.86 1.09c-.81.73-1.27 1.64-1.37 2.72l-.07.54c0 .09-.05.14-.16.14l-.54.11c-.84.07-1.55.41-2.11 1.03c-.57.62-.85 1.35-.85 2.2c0 .87.3 1.62.89 2.25c.59.63 1.31.97 2.17 1.02c.12 0 .18.06.18.17v1.34c0 .11-.06.17-.18.17c-.66-.03-1.28-.18-1.88-.45S6.42 20.8 6 20.36c-.43-.44-.77-.95-1.02-1.55s-.38-1.22-.38-1.88m5.42 6.77c0-.03.01-.08.02-.13s.02-.09.02-.11l.27-1.03c.07-.22.2-.4.4-.51c.2-.12.41-.14.64-.07c.23.07.4.2.52.4s.14.41.07.64l-.24 1.01c-.13.44-.38.66-.76.66h-.09c-.03 0-.07-.01-.11-.01c-.04-.01-.07-.01-.1-.01c-.21-.06-.37-.18-.48-.34s-.16-.34-.16-.5m1.32-4.82c0-.02 0-.06.01-.11s.01-.08.01-.09l.3-1.05c.06-.19.17-.34.32-.45c.15-.1.31-.15.47-.15h.08c.03 0 .06.01.09.01c.03.01.06.01.08.01c.23.07.4.2.51.4c.12.2.14.41.07.64l-.24 1c-.07.28-.2.47-.4.59s-.42.12-.65.02c-.22-.06-.38-.17-.49-.34s-.16-.32-.16-.48m1.23 7.95c0-.03.01-.07.02-.13s.02-.09.02-.12l.29-.99c.06-.24.2-.42.4-.54c.2-.12.42-.15.65-.08c.23.07.39.2.51.41s.13.42.07.65l-.25 1.04c-.11.41-.37.61-.8.61c-.05 0-.13-.01-.24-.04a.718.718 0 0 1-.49-.3a.836.836 0 0 1-.18-.51m1.34-4.77c0-.06.01-.14.04-.25l.27-1.03c.07-.23.2-.4.41-.51c.2-.12.42-.14.65-.07a.805.805 0 0 1 .57 1.04l-.24.99c-.13.45-.37.68-.72.68c-.04 0-.15-.02-.31-.06a.718.718 0 0 1-.49-.3a.782.782 0 0 1-.18-.49m2.82 1.68c0-.07.01-.15.03-.24l.28-.99c.07-.24.2-.42.41-.54s.41-.15.63-.09c.23.07.41.2.53.41c.12.2.15.41.09.63l-.29 1.06c-.1.41-.36.61-.79.61c-.09 0-.18-.01-.26-.03a.7.7 0 0 1-.46-.3c-.1-.18-.16-.35-.17-.52m1.38-4.76c0-.03.02-.12.05-.26l.3-1.03c.04-.21.13-.37.29-.47c.16-.1.32-.15.49-.14c.04-.01.13 0 .24.03c.22.05.39.18.52.38c.12.17.14.38.07.65l-.24 1.03c-.13.43-.38.65-.76.65c-.06 0-.17-.02-.34-.06a.823.823 0 0 1-.62-.78"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sleet(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.09 16.95c0-1.16.36-2.19 1.08-3.09s1.64-1.49 2.74-1.74c.31-1.37 1.01-2.49 2.1-3.37s2.35-1.32 3.77-1.32c.81 0 1.55.13 2.2.4c0 .01 0 .01.01.02c.84.32 1.58.84 2.21 1.55h.03c.68.73 1.13 1.6 1.37 2.62h.31c1.08 0 2.02.29 2.83.86v-.01c.53.36.98.8 1.34 1.33c.36.53.6 1.11.73 1.74c.04.21.06.38.08.52v.06c0 .01 0 .06.01.17s.01.19.01.24v.03c0 .88-.21 1.7-.64 2.45c-.42.75-1 1.36-1.74 1.81c-.73.45-1.54.69-2.42.72c-.13 0-.2-.06-.2-.17v-1.34c0-.13.07-.19.2-.19c.86-.04 1.58-.38 2.18-1.02c.6-.64.9-1.39.9-2.26c0-.89-.32-1.65-.97-2.29a3.15 3.15 0 0 0-2.31-.96H18.3c-.12 0-.18-.06-.18-.17l-.08-.59c-.1-1-.52-1.86-1.27-2.59c-.01-.01-.01-.02-.02-.03s-.02-.02-.02-.03c-.01-.01-.02-.02-.04-.02c0-.02-.01-.03-.02-.03a4.24 4.24 0 0 0-2.56-1.12c-.07-.01-.18-.01-.34-.01c-1.1 0-2.05.37-2.86 1.1s-1.27 1.64-1.37 2.72l-.08.59c-.03.09-.11.14-.22.14l-.49.03c-.84.1-1.55.46-2.11 1.07s-.84 1.34-.84 2.18v.04h.03c.01.48.11.93.3 1.35c.2.43.46.79.8 1.09c.21.18.45.34.74.48v.01c.4.19.8.3 1.21.32c.11 0 .17.06.17.18v1.34c0 .11-.06.17-.17.17c-.52-.03-1.01-.13-1.48-.3v.01c-.83-.29-1.54-.77-2.11-1.43s-.95-1.44-1.11-2.31v-.03c-.01-.01-.01-.02-.01-.04c-.06-.26-.08-.56-.08-.88m5.5 7.17c0-.03.01-.07.02-.13c.01-.05.02-.09.02-.12l.09-.59c.07-.24.2-.41.41-.53s.43-.14.68-.08c.23.07.39.21.51.41c.11.21.13.42.07.63l-.14.6c-.1.44-.35.66-.76.66c-.03 0-.08-.01-.14-.02c-.06-.01-.1-.02-.14-.02a.761.761 0 0 1-.46-.33a.791.791 0 0 1-.16-.48m.74-2.94c0-.24.08-.43.23-.59s.35-.23.59-.23s.43.08.59.23s.23.35.23.59c0 .23-.08.42-.23.58s-.35.23-.59.23c-.23 0-.42-.08-.57-.24a.789.789 0 0 1-.25-.57m1.64 5.99c0-.04.01-.11.04-.23l.13-.59c.07-.23.21-.4.41-.51c.21-.12.42-.14.63-.07c.23.04.41.17.53.37c.12.2.15.43.08.68l-.13.59c-.1.41-.37.61-.8.61c-.07 0-.16-.01-.24-.03a.831.831 0 0 1-.49-.33a.92.92 0 0 1-.16-.49m.79-2.91a.784.784 0 0 1 .82-.81c.24 0 .43.08.59.23c.16.16.23.35.23.58c0 .24-.08.43-.23.59c-.16.16-.35.23-.59.23c-.23 0-.42-.08-.58-.24a.785.785 0 0 1-.24-.58m.62-2.27c0-.1.01-.19.03-.27l.23-.9c.07-.23.21-.41.41-.53c.21-.12.42-.15.64-.08c.24.07.41.2.53.4s.14.41.07.63l-.26.9c-.08.28-.22.46-.41.56c-.19.1-.41.12-.64.06a.657.657 0 0 1-.45-.3a.69.69 0 0 1-.15-.47m2.76 2.14c0-.03 0-.08.01-.13s.01-.09.01-.11l.09-.59c.07-.24.2-.41.41-.53s.43-.14.68-.08c.23.07.4.21.51.41c.12.21.14.42.07.63l-.14.6c-.1.44-.35.66-.76.66c-.03 0-.08-.01-.14-.02c-.06-.01-.11-.02-.14-.02a.768.768 0 0 1-.45-.33a1 1 0 0 1-.15-.49m.74-2.94c0-.24.08-.43.23-.59c.16-.16.35-.23.59-.23s.43.08.59.23c.16.16.23.35.23.59a.783.783 0 0 1-.81.81c-.24 0-.43-.08-.58-.24a.756.756 0 0 1-.25-.57"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmallCraftAdvisory(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.81 24.6V7.45h1.03V24.6zm1.73-9.74V7.45l8.65 3.69z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Smog(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.35 12.46c0-.94.3-1.77.9-2.5s1.37-1.21 2.31-1.43c.23-1.11.79-2.03 1.7-2.75a4.89 4.89 0 0 1 3.12-1.08c1.12 0 2.13.35 3 1.04c.88.69 1.45 1.59 1.72 2.7h.28c.76 0 1.46.16 2.12.49s1.18.77 1.57 1.34c.39.57.59 1.18.59 1.84c0 1.12-.43 2.08-1.29 2.86c0 .35-.11.75-.32 1.2c-.22.45-.5.86-.87 1.23c-.36.37-.73.59-1.1.68c-.12.62-.41 1.14-.86 1.57s-.99.71-1.63.85c.3.3.45.65.45 1.06c0 .49-.17.91-.52 1.26s-.77.52-1.27.52c-.49 0-.91-.17-1.26-.52s-.53-.77-.53-1.26c0-.06.01-.14.04-.26s.04-.21.04-.27h-.08c-.59 0-1.09-.21-1.51-.63a2.07 2.07 0 0 1-.63-1.51c0-.23.12-.58.37-1.06c-.49-.26-.88-.67-1.17-1.26h-1.25a4.296 4.296 0 0 1-2.78-1.3a3.888 3.888 0 0 1-1.14-2.81"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Smoke(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.34 12.48c0-.94.3-1.78.89-2.52s1.34-1.21 2.25-1.41c.25-1.12.82-2.05 1.72-2.77s1.92-1.08 3.08-1.08c1.12 0 2.13.35 3.02 1.05c.89.7 1.46 1.6 1.73 2.69h.27c1.12 0 2.08.39 2.88 1.18c.79.78 1.19 1.74 1.19 2.85c0 .6-.12 1.17-.37 1.7s-.59.99-1.03 1.37v.03c0 .59-.19 1.12-.56 1.59c-.37.47-.84.76-1.4.89c-.14.62-.45 1.15-.91 1.58c-.46.43-1.01.7-1.63.8c.29.34.43.72.43 1.13c0 .48-.17.89-.51 1.24c-.34.34-.75.52-1.23.52s-.89-.17-1.23-.52c-.34-.34-.51-.76-.51-1.24c0-.19.03-.38.1-.57h-.1c-.58 0-1.08-.21-1.5-.63c-.42-.42-.63-.92-.63-1.5c0-.4.1-.76.3-1.07c-.52-.29-.89-.7-1.12-1.25h-1.28v-.01c-1.07-.07-1.98-.49-2.73-1.27s-1.12-1.7-1.12-2.78m1.4-.25c0 .8.28 1.48.84 2.04s1.24.84 2.03.84c.49 0 .95-.11 1.37-.34c.12.74.47 1.36 1.04 1.86s1.25.74 2.02.74c.87 0 1.61-.31 2.22-.92c.41.48.92.71 1.54.71c.57 0 1.05-.2 1.46-.6c.4-.4.6-.89.6-1.46c.4-.27.72-.61.95-1.04c.23-.42.35-.88.35-1.37c0-.79-.28-1.47-.85-2.02s-1.25-.83-2.05-.83c-.56 0-1.07.15-1.53.44c.06-.24.08-.51.08-.79c0-.96-.34-1.78-1.03-2.46c-.69-.68-1.52-1.01-2.49-1.01c-.94 0-1.75.33-2.43.97s-1.04 1.44-1.07 2.37h-.17c-.79 0-1.46.28-2.03.84s-.85 1.25-.85 2.03"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Snow(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.64 16.95c0-1.16.35-2.18 1.06-3.08s1.62-1.48 2.74-1.76c.31-1.36 1.01-2.48 2.1-3.36s2.34-1.31 3.75-1.31c1.38 0 2.6.43 3.68 1.28c1.08.85 1.78 1.95 2.1 3.29h.32c.89 0 1.72.22 2.48.66c.76.44 1.37 1.04 1.81 1.8c.44.76.67 1.59.67 2.48c0 1.32-.46 2.47-1.39 3.42c-.92.96-2.05 1.46-3.38 1.5c-.13 0-.2-.06-.2-.17v-1.33c0-.12.07-.18.2-.18c.85-.04 1.58-.38 2.18-1.02s.9-1.38.9-2.23c0-.89-.32-1.65-.97-2.3s-1.42-.97-2.32-.97h-1.61c-.12 0-.18-.06-.18-.17l-.08-.58c-.11-1.08-.58-1.99-1.39-2.72c-.82-.73-1.76-1.1-2.85-1.1c-1.1 0-2.05.37-2.86 1.11c-.81.74-1.27 1.65-1.37 2.75l-.06.5c0 .12-.07.19-.2.19l-.53.07c-.83.07-1.53.41-2.1 1.04s-.85 1.35-.85 2.19c0 .85.3 1.59.9 2.23s1.33.97 2.18 1.02c.11 0 .17.06.17.18v1.33c0 .11-.06.17-.17.17c-1.34-.04-2.47-.54-3.4-1.5c-.87-.96-1.33-2.11-1.33-3.43M11 21.02c0-.22.08-.42.24-.58c.16-.16.35-.24.59-.24c.23 0 .43.08.59.24c.16.16.24.36.24.58c0 .24-.08.44-.24.6c-.16.17-.35.25-.59.25c-.23 0-.43-.08-.59-.25a.814.814 0 0 1-.24-.6m0 3.63c0-.24.08-.44.24-.6c.16-.15.35-.23.58-.23c.23 0 .43.08.59.23c.16.16.24.35.24.59s-.08.43-.24.59c-.16.16-.35.23-.59.23a.84.84 0 0 1-.59-.23a.8.8 0 0 1-.23-.58m3.19-1.7c0-.23.08-.44.25-.62c.16-.16.35-.24.57-.24c.23 0 .43.09.6.26c.17.17.26.37.26.6c0 .23-.08.43-.25.6c-.17.17-.37.25-.61.25c-.23 0-.42-.08-.58-.25s-.24-.37-.24-.6m0-3.62c0-.23.08-.43.25-.6c.18-.16.37-.24.57-.24c.24 0 .44.08.61.25a.8.8 0 0 1 .25.6c0 .23-.08.43-.25.59c-.17.16-.37.24-.61.24c-.23 0-.42-.08-.58-.24a.847.847 0 0 1-.24-.6m0 7.28c0-.23.08-.43.25-.61c.16-.16.35-.24.57-.24c.24 0 .44.08.61.25c.17.17.25.37.25.6s-.08.43-.25.59c-.17.16-.37.24-.61.24a.824.824 0 0 1-.82-.83m3.22-5.59c0-.22.08-.41.25-.58c.17-.17.37-.25.6-.25c.23 0 .43.08.59.24c.16.16.24.36.24.58c0 .24-.08.44-.24.6c-.16.17-.35.25-.59.25s-.44-.08-.6-.25a.816.816 0 0 1-.25-.59m0 3.63c0-.22.08-.42.25-.6c.16-.15.36-.23.6-.23s.43.08.59.23s.23.35.23.59s-.08.43-.23.59c-.16.16-.35.23-.59.23s-.44-.08-.6-.24a.756.756 0 0 1-.25-.57"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnowWind(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.64 16.95c0-1.16.35-2.18 1.06-3.08s1.62-1.48 2.74-1.76c.31-1.36 1.01-2.48 2.1-3.36s2.34-1.31 3.75-1.31c1.38 0 2.6.43 3.68 1.28c1.08.85 1.78 1.95 2.1 3.29h.32c.89 0 1.72.22 2.48.66c.76.44 1.37 1.04 1.81 1.8c.44.76.67 1.59.67 2.48c0 1.32-.46 2.47-1.39 3.42c-.92.96-2.05 1.46-3.38 1.5c-.13 0-.2-.06-.2-.17v-1.33c0-.12.07-.18.2-.18c.85-.04 1.58-.38 2.18-1.02s.9-1.38.9-2.23c0-.89-.32-1.65-.97-2.3s-1.42-.97-2.32-.97h-1.61c-.12 0-.18-.06-.18-.17l-.08-.58c-.11-1.08-.58-1.99-1.39-2.72c-.82-.73-1.76-1.1-2.85-1.1c-1.1 0-2.05.37-2.86 1.11c-.81.74-1.27 1.65-1.37 2.75l-.06.5c0 .12-.07.19-.2.19l-.53.07c-.83.07-1.53.41-2.1 1.04s-.85 1.35-.85 2.19c0 .85.3 1.59.9 2.23s1.33.97 2.18 1.02c.11 0 .17.06.17.18v1.33c0 .11-.06.17-.17.17c-1.34-.04-2.47-.54-3.4-1.5c-.87-.96-1.33-2.11-1.33-3.43m5.5 7.7a.816.816 0 0 1 .82-.83c.23 0 .43.08.59.23s.24.35.24.59s-.08.43-.24.59c-.16.16-.35.23-.59.23a.84.84 0 0 1-.59-.23a.8.8 0 0 1-.23-.58m.86-3.63c0-.22.08-.42.24-.58c.16-.16.35-.24.59-.24c.23 0 .43.08.59.24c.16.16.24.36.24.58c0 .24-.08.44-.24.6c-.16.17-.35.25-.59.25c-.23 0-.43-.08-.59-.25a.814.814 0 0 1-.24-.6m1.9 5.59c0-.23.08-.43.25-.61c.16-.16.35-.24.57-.24c.24 0 .44.08.61.25c.17.17.25.37.25.6s-.08.43-.25.59c-.17.16-.37.24-.61.24c-.23 0-.42-.08-.58-.24a.867.867 0 0 1-.24-.59m.87-3.66c0-.24.08-.44.24-.62c.16-.16.36-.24.58-.24c.23 0 .43.08.6.25c.17.17.25.37.25.61c0 .23-.08.43-.25.6s-.37.25-.6.25c-.23 0-.42-.08-.58-.25a.847.847 0 0 1-.24-.6m.42-3.62c0-.23.08-.43.25-.6c.18-.16.37-.24.57-.24c.24 0 .44.08.61.25a.8.8 0 0 1 .25.6c0 .23-.08.43-.25.59c-.17.16-.37.24-.61.24c-.23 0-.42-.08-.58-.24a.847.847 0 0 1-.24-.6m2.37 5.32c0-.23.08-.43.24-.6c.16-.15.36-.23.6-.23s.43.08.59.23c.16.16.23.35.23.59s-.08.43-.23.59c-.16.16-.35.23-.59.23s-.44-.08-.6-.24a.748.748 0 0 1-.24-.57m.85-3.63c0-.22.08-.41.25-.58c.17-.17.37-.25.6-.25c.23 0 .43.08.59.24c.16.16.24.36.24.58c0 .24-.08.44-.24.6c-.16.17-.35.25-.59.25s-.44-.08-.6-.25a.816.816 0 0 1-.25-.59"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnowflakeCold(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.46 14.41c0-.19.07-.36.2-.5a.63.63 0 0 1 .48-.21c.19 0 .36.07.49.21c.13.14.2.3.2.5c0 .19-.07.35-.2.48a.673.673 0 0 1-.96 0a.624.624 0 0 1-.21-.48m1.98 0c0-.19.07-.36.2-.5a.63.63 0 0 1 .48-.21h3.23l-2.28-2.28a.594.594 0 0 1-.21-.47c0-.19.07-.35.21-.49c.14-.14.3-.21.49-.21s.35.07.49.21l2.27 2.27V9.52c0-.19.07-.36.21-.5s.3-.21.5-.21c.19 0 .35.07.48.21c.13.14.2.3.2.5v3.23l2.29-2.3c.14-.14.3-.21.48-.21s.35.07.49.21c.12.14.18.3.18.49s-.06.35-.18.47l-2.28 2.28h3.23c.18 0 .34.07.47.21c.13.14.2.3.2.5c0 .18-.07.34-.2.47a.65.65 0 0 1-.47.2h-3.23l2.29 2.29c.12.12.18.28.18.47s-.06.35-.18.49c-.14.14-.31.21-.49.21s-.35-.07-.48-.21l-2.29-2.3v3.24c0 .19-.07.36-.2.49s-.29.2-.48.2a.72.72 0 0 1-.5-.2a.65.65 0 0 1-.21-.49v-3.22l-2.27 2.27c-.14.14-.3.21-.49.21s-.35-.07-.49-.21s-.21-.3-.21-.49s.07-.34.21-.47l2.3-2.29h-3.24c-.19 0-.35-.07-.48-.2a.512.512 0 0 1-.22-.45m.01 4.84c0-.18.07-.35.21-.48c.12-.14.28-.21.47-.21c.19 0 .35.07.49.21c.14.14.21.3.21.48s-.07.35-.21.48c-.14.14-.3.21-.49.21s-.35-.07-.47-.21a.63.63 0 0 1-.21-.48m0-9.71c0-.18.07-.35.21-.48c.12-.14.28-.21.47-.21c.19 0 .35.07.49.21c.14.14.21.3.21.48s-.07.35-.21.49c-.14.14-.3.21-.49.21s-.35-.07-.47-.21a.683.683 0 0 1-.21-.49m4.85 11.71c0-.18.07-.34.21-.47c.14-.13.3-.2.5-.2c.19 0 .35.07.48.2s.2.29.2.47c0 .19-.07.36-.2.49s-.29.2-.48.2a.72.72 0 0 1-.5-.2a.635.635 0 0 1-.21-.49m0-13.71c0-.19.07-.36.21-.49s.3-.2.5-.2c.19 0 .35.07.48.2a.68.68 0 0 1-.48 1.16a.72.72 0 0 1-.5-.2a.624.624 0 0 1-.21-.47m4.87 11.71c0-.18.07-.35.21-.48c.12-.14.28-.21.47-.21c.19 0 .35.07.49.21c.14.14.21.3.21.48s-.07.35-.21.48c-.14.14-.3.21-.49.21s-.35-.07-.47-.21a.63.63 0 0 1-.21-.48m0-9.71c0-.18.07-.35.21-.48c.12-.14.28-.21.47-.21c.19 0 .35.07.49.21c.14.14.21.3.21.48s-.07.35-.21.49c-.14.14-.3.21-.49.21s-.35-.07-.47-.21a.683.683 0 0 1-.21-.49m2 4.87c0-.19.07-.36.2-.5s.3-.21.49-.21c.18 0 .34.07.47.21c.13.14.2.3.2.5c0 .18-.07.34-.2.47a.68.68 0 0 1-.96 0a.65.65 0 0 1-.2-.47"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SolarEclipse(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.37 14.62c0-.24.08-.45.25-.62c.17-.16.38-.24.6-.24h2.04c.23 0 .42.08.58.25c.15.17.23.37.23.61s-.07.44-.22.61c-.15.17-.35.25-.58.25H5.23c-.23 0-.43-.08-.6-.25a.832.832 0 0 1-.26-.61m2.86 6.93c0-.23.08-.43.23-.61l1.47-1.43c.15-.16.35-.23.59-.23s.44.08.6.23s.24.34.24.57c0 .24-.08.46-.24.64L8.7 22.14c-.41.32-.82.32-1.23 0a.807.807 0 0 1-.24-.59m0-13.84c0-.23.08-.43.23-.61c.2-.17.41-.25.64-.25c.22 0 .42.08.59.24l1.43 1.47c.16.15.24.35.24.59s-.08.44-.24.6s-.36.24-.6.24s-.44-.08-.59-.24L7.47 8.32a.837.837 0 0 1-.24-.61m2.55 6.91c0-.93.23-1.8.7-2.6s1.1-1.44 1.91-1.91s1.67-.7 2.6-.7c.7 0 1.37.14 2.02.42c.64.28 1.2.65 1.66 1.12c.47.47.84 1.02 1.11 1.66c.27.64.41 1.32.41 2.02c0 .94-.23 1.81-.7 2.61c-.47.8-1.1 1.43-1.9 1.9c-.8.47-1.67.7-2.61.7s-1.81-.23-2.61-.7c-.8-.47-1.43-1.1-1.9-1.9c-.45-.81-.69-1.68-.69-2.62m4.36 7.78c0-.24.08-.44.25-.6s.37-.24.6-.24c.24 0 .45.08.61.24s.24.36.24.6v1.99c0 .24-.08.45-.25.62c-.17.17-.37.25-.6.25s-.44-.08-.6-.25a.845.845 0 0 1-.25-.62zm0-15.5V4.86c0-.23.08-.43.25-.6c.17-.17.37-.26.61-.26s.43.08.6.25c.17.17.25.37.25.6V6.9c0 .23-.08.42-.25.58s-.37.23-.6.23s-.44-.08-.6-.23s-.26-.35-.26-.58m.11 4.32c.87.11 1.6.49 2.19 1.15c.59.66.89 1.44.89 2.33c0 .83-.26 1.56-.78 2.2c-.52.63-1.18 1.04-1.98 1.21c.2.02.35.04.44.04c.97 0 1.81-.35 2.5-1.04s1.04-1.52 1.04-2.5c0-.96-.35-1.78-1.04-2.47c-.69-.68-1.52-1.02-2.5-1.02c-.27.02-.52.05-.76.1m5.41 8.86c0-.23.08-.42.23-.56c.15-.16.34-.23.56-.23c.24 0 .44.08.6.23l1.46 1.43c.16.17.24.38.24.61c0 .23-.08.43-.24.59c-.4.31-.8.31-1.2 0l-1.42-1.42a.974.974 0 0 1-.23-.65m0-10.92c0-.25.08-.45.23-.59l1.42-1.47a.84.84 0 0 1 .59-.24c.24 0 .44.08.6.25c.17.17.25.37.25.6c0 .25-.08.46-.24.62l-1.46 1.43c-.18.16-.38.24-.6.24c-.23 0-.41-.08-.56-.24s-.23-.36-.23-.6m2.26 5.46c0-.24.08-.44.24-.62c.16-.16.35-.24.57-.24h2.02c.23 0 .43.09.6.26c.17.17.26.37.26.6s-.09.43-.26.6c-.17.17-.37.25-.6.25h-2.02c-.23 0-.43-.08-.58-.25s-.23-.36-.23-.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sprinkle(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.64 16.91c0-1.15.36-2.17 1.08-3.07a4.82 4.82 0 0 1 2.73-1.73c.31-1.36 1.01-2.48 2.1-3.35s2.35-1.31 3.76-1.31c1.38 0 2.6.43 3.68 1.27A5.88 5.88 0 0 1 20.1 12h.31c.89 0 1.72.22 2.48.65s1.37 1.03 1.81 1.78c.44.75.67 1.58.67 2.47c0 1.34-.46 2.49-1.38 3.45s-2.05 1.47-3.38 1.51c-.13 0-.2-.06-.2-.17v-1.33c0-.12.07-.18.2-.18c.86-.04 1.58-.38 2.18-1.02s.9-1.39.9-2.26s-.32-1.62-.98-2.26c-.65-.64-1.42-.96-2.31-.96h-1.6c-.12 0-.19-.06-.19-.17l-.07-.58a4.108 4.108 0 0 0-1.38-2.71c-.82-.73-1.77-1.1-2.85-1.1c-1.09 0-2.05.36-2.86 1.09c-.81.73-1.27 1.63-1.38 2.71l-.06.54c0 .12-.07.18-.2.18l-.53.03c-.82.04-1.51.37-2.09 1s-.86 1.37-.86 2.22c0 .87.3 1.62.9 2.26s1.33.98 2.18 1.02c.11 0 .17.06.17.18v1.33c0 .11-.06.17-.17.17c-1.34-.06-2.47-.57-3.4-1.53s-1.37-2.08-1.37-3.41m5.93.88c0-.24.12-.57.37-.99c.24-.42.47-.75.68-1.01c.21-.24.34-.38.38-.42l.36.4c.26.28.5.61.72 1.02c.22.4.33.74.33 1c0 .39-.13.72-.4.98s-.6.39-1 .39c-.39 0-.73-.13-1.01-.4c-.29-.26-.43-.59-.43-.97m2.98 3.99c0-.28.08-.59.24-.96s.35-.7.59-1.02c.18-.26.4-.54.67-.84c.26-.3.46-.52.6-.65c.07-.06.15-.14.24-.23l.24.23c.38.33.8.82 1.27 1.46c.24.33.43.68.59 1.04s.23.68.23.97c0 .64-.23 1.19-.68 1.65s-1.01.68-1.66.68c-.64 0-1.19-.23-1.65-.67c-.46-.46-.68-1.01-.68-1.66m1.47-6.66c0-.42.32-.95.97-1.6l.24.25c.18.21.33.45.48.71c.14.26.22.47.22.64c0 .26-.09.48-.28.66c-.18.18-.4.28-.66.28c-.27 0-.5-.09-.69-.28a.87.87 0 0 1-.28-.66"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stars(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.37 16.18c.65-.03 1.2-.28 1.65-.75c.45-.47.68-1.03.68-1.68c0 .65.22 1.21.67 1.68c.45.47 1 .72 1.65.75c-.65.03-1.2.28-1.65.75a2.34 2.34 0 0 0-.67 1.68c0-.65-.22-1.21-.68-1.68c-.45-.47-1-.72-1.65-.75m2.33-7.2c1.26-.06 2.33-.55 3.21-1.47c.88-.92 1.32-2.01 1.32-3.28c0 1.27.44 2.36 1.32 3.28s1.95 1.4 3.22 1.47c-.83.04-1.59.27-2.29.71c-.69.43-1.24 1.01-1.65 1.73c-.4.72-.6 1.49-.6 2.33c0-1.27-.44-2.37-1.32-3.29c-.88-.93-1.95-1.42-3.21-1.48m3.32 10.77c.95-.04 1.76-.41 2.42-1.1c.66-.69.99-1.51.99-2.47c0 .96.33 1.78.99 2.47c.66.69 1.46 1.06 2.41 1.1c-.95.04-1.75.41-2.41 1.1c-.66.69-.99 1.51-.99 2.47c0-.96-.33-1.78-.99-2.47a3.503 3.503 0 0 0-2.42-1.1m6.81-4.74c.95-.04 1.75-.41 2.41-1.1c.66-.69.98-1.51.98-2.48c0 .96.33 1.78.99 2.47s1.47 1.06 2.42 1.1c-.95.04-1.76.41-2.42 1.1c-.66.69-.99 1.51-.99 2.47c0-.96-.33-1.78-.98-2.47c-.66-.68-1.46-1.05-2.41-1.09"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StormShowers(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.63 16.91c0 1.11.33 2.1.99 2.97s1.51 1.47 2.56 1.79l-.65 1.68c-.03.14.02.22.14.22H9.8l-1.04 3.78h.28l3.97-5.22c.04-.04.04-.09.02-.14s-.07-.07-.14-.07h-2.18l2.48-4.64c.06-.14.02-.21-.14-.21h-2.94a.25.25 0 0 0-.22.14L8.8 20.08c-.71-.18-1.3-.57-1.77-1.16s-.7-1.26-.7-2.01c0-.83.28-1.55.84-2.16s1.26-.96 2.1-1.06l.52-.07c.13 0 .2-.06.2-.17l.07-.52c.1-1.08.55-1.99 1.36-2.72c.81-.73 1.76-1.1 2.85-1.1s2.04.37 2.85 1.1c.82.73 1.28 1.64 1.4 2.72l.06.58c0 .12.06.18.19.18h1.61c.91 0 1.68.32 2.32.95c.64.63.96 1.39.96 2.28c0 .85-.3 1.59-.89 2.21c-.59.62-1.32.96-2.19 1.03c-.13 0-.2.06-.2.19v1.37c0 .11.07.17.2.17c.65-.02 1.27-.17 1.86-.44c.59-.27 1.1-.63 1.52-1.07c.42-.44.76-.96 1.01-1.57c.25-.6.38-1.23.38-1.88c0-.9-.22-1.73-.67-2.49a4.96 4.96 0 0 0-4.29-2.46h-.32c-.33-1.33-1.03-2.42-2.11-3.26c-1.08-.84-2.3-1.27-3.68-1.27c-1.41 0-2.67.44-3.76 1.31s-1.79 1.99-2.1 3.36c-1.1.26-2.01.83-2.73 1.73s-1.06 1.91-1.06 3.06m8.16 9.86a.822.822 0 0 0 .61.78c.14.03.22.05.25.05c.09 0 .21-.03.38-.1c.21-.09.35-.27.42-.52l.28-1.05c.06-.22.04-.43-.08-.63s-.29-.33-.53-.4a.776.776 0 0 0-.63.08c-.2.12-.34.29-.41.53l-.27 1c-.01.17-.02.26-.02.26M14.13 22c0 .14.05.29.15.44c.1.15.25.26.45.33c.22.07.44.06.64-.05s.33-.28.4-.52l.3-1.04c.06-.22.03-.43-.08-.63c-.12-.2-.3-.34-.53-.41c-.23-.06-.44-.04-.65.08s-.34.29-.41.52l-.24 1.01c-.02.17-.03.26-.03.27m2.82 1.65c0 .17.05.34.16.51c.11.17.27.28.47.35c.02 0 .06.01.12.02c.05.01.09.02.12.02c.13 0 .26-.02.38-.08c.19-.07.33-.26.41-.57l.25-1.01c.07-.23.05-.45-.06-.66a.832.832 0 0 0-.5-.42a.926.926 0 0 0-.68.08c-.2.12-.33.3-.37.53l-.27 1.03c-.02.06-.03.13-.03.2m1.36-4.79c-.01.16.04.31.15.47c.11.16.27.28.49.38c.08.04.16.06.26.06c.11 0 .22-.03.34-.08c.21-.1.35-.29.44-.57l.29-1.03c.02-.13.03-.2.03-.22c0-.17-.05-.33-.16-.49s-.27-.27-.49-.33c-.02 0-.06-.01-.11-.02c-.06-.03-.1-.03-.13-.03c-.17 0-.33.05-.49.15c-.16.1-.27.26-.33.48l-.27 1.01c-.01.08-.02.15-.02.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StormWarning(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.76 24.6V7.45h1.13V24.6zm1.94-10.55v-6.6h8.55v6.6zm2.36-2h3.81v-2.5h-3.81z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StrongWind(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.1 16.97c0 .24.09.45.28.62c.16.19.37.28.63.28H18.7c.29 0 .53.1.73.3c.2.2.3.45.3.74s-.1.53-.3.72c-.2.19-.44.29-.74.29c-.29 0-.54-.1-.73-.29a.76.76 0 0 0-.6-.26c-.25 0-.46.09-.64.26s-.27.38-.27.61c0 .25.09.46.28.63c.56.55 1.22.83 1.96.83c.78 0 1.45-.27 2.01-.81c.56-.54.83-1.19.83-1.97s-.28-1.44-.84-2c-.56-.56-1.23-.84-2-.84H4.01a.9.9 0 0 0-.64.26c-.18.17-.27.38-.27.63m0-3.28c0 .23.09.43.28.61c.17.18.38.26.63.26h20.04c.78 0 1.45-.27 2.01-.82c.56-.54.84-1.2.84-1.97s-.28-1.44-.84-1.99s-1.23-.83-2.01-.83c-.77 0-1.42.27-1.95.8c-.18.16-.27.38-.27.67c0 .26.09.47.26.63c.17.16.38.24.63.24c.24 0 .45-.08.63-.24c.19-.21.42-.31.7-.31c.29 0 .53.1.73.3c.2.2.3.44.3.73s-.1.53-.3.72c-.2.19-.44.29-.73.29H4.01a.908.908 0 0 0-.91.91"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sunrise(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.75 15.36c0-.25.1-.48.3-.69c.22-.19.46-.29.7-.29h2.33c.27 0 .49.1.67.29c.18.19.27.43.27.69c0 .29-.09.53-.27.72a.9.9 0 0 1-.67.29H3.75c-.27 0-.5-.1-.7-.3c-.2-.21-.3-.45-.3-.71m3.33-7.98c0-.27.09-.5.26-.68c.23-.2.46-.3.71-.3c.26 0 .49.1.68.29l1.64 1.65c.19.22.28.45.28.69c0 .28-.09.52-.27.7s-.4.28-.66.28c-.24 0-.48-.1-.7-.29L6.34 8.11c-.17-.21-.26-.46-.26-.73m2 13.5c0-.28.1-.51.29-.68c.18-.17.4-.26.68-.26h2.63l3.11-2.92c.1-.08.21-.08.34 0l3.16 2.92h2.77c.27 0 .5.09.69.28a.9.9 0 0 1 .29.67c0 .27-.1.5-.29.69c-.19.19-.42.29-.69.29h-3.38c-.1 0-.2-.02-.29-.07l-2.41-2.27l-2.39 2.27c-.08.05-.17.07-.28.07H9.05a.974.974 0 0 1-.97-.99M9 15.36c0 .97.21 1.85.62 2.64c.02.12.11.18.25.18h1.88c.07 0 .12-.03.15-.08c.03-.06.02-.12-.02-.19c-.64-.77-.96-1.62-.96-2.55c0-1.12.4-2.08 1.2-2.87c.8-.79 1.76-1.18 2.89-1.18c1.12 0 2.07.39 2.86 1.18c.79.79 1.19 1.74 1.19 2.87c0 .94-.32 1.79-.95 2.55c-.04.07-.05.13-.03.19s.07.08.15.08h1.9c.13 0 .21-.06.23-.18c.44-.77.64-1.65.64-2.64c0-.81-.16-1.59-.48-2.32c-.32-.74-.75-1.37-1.28-1.91a6.1 6.1 0 0 0-1.91-1.28c-.74-.32-1.51-.47-2.32-.47c-.81 0-1.59.16-2.33.47c-.74.32-1.38.74-1.92 1.28A5.96 5.96 0 0 0 9 15.36m5.03-8.96V4.1c0-.29.09-.52.28-.71s.43-.28.71-.28c.28 0 .51.09.7.28s.28.44.28.72v2.3c0 .29-.09.52-.28.71c-.18.18-.42.28-.7.28a.95.95 0 0 1-.71-.28a.972.972 0 0 1-.28-.72m6.35 2.64c0-.25.09-.48.27-.69l1.62-1.65c.19-.19.43-.29.7-.29c.27 0 .51.1.69.29c.19.19.28.42.28.69c0 .29-.09.53-.26.73L22 9.73c-.21.19-.45.29-.7.29c-.27 0-.49-.09-.66-.28s-.26-.42-.26-.7m2.61 6.32c0-.27.09-.5.27-.69c.18-.19.4-.29.66-.29h2.35c.27 0 .5.1.69.29a.99.99 0 0 1 0 1.4c-.19.2-.42.3-.69.3h-2.35c-.27 0-.49-.1-.67-.29c-.17-.2-.26-.44-.26-.72"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sunset(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.88 15.47c0-.28.1-.5.3-.68c.17-.18.4-.26.68-.26h2.31a.904.904 0 0 1 .93.94c0 .28-.09.52-.27.71c-.18.19-.4.29-.66.29h-2.3c-.27 0-.5-.1-.69-.3c-.2-.2-.3-.43-.3-.7m3.29-7.86c0-.28.08-.51.25-.68c.2-.2.43-.3.7-.3c.29 0 .51.1.68.3l1.62 1.63c.46.44.46.89 0 1.35a.88.88 0 0 1-.65.28c-.22 0-.44-.09-.68-.28L6.43 8.29a.96.96 0 0 1-.26-.68m1.97 13.28c0-.26.1-.49.3-.69c.18-.18.41-.27.68-.27h3.22c.11 0 .2.02.28.08l2.35 2.22L17.36 20c.07-.05.17-.08.29-.08h3.3c.27 0 .5.09.69.28c.19.19.29.42.29.68c0 .27-.1.5-.29.69c-.19.19-.42.29-.69.29h-2.68l-3.14 2.84c-.12.09-.23.09-.33 0l-3.08-2.84h-2.6c-.27 0-.5-.1-.69-.29a.906.906 0 0 1-.29-.68m.94-5.42c0 .99.19 1.87.58 2.62c.06.11.15.16.27.16h1.87c.08 0 .13-.02.15-.07c.02-.05-.01-.11-.07-.18c-.59-.74-.89-1.59-.89-2.53c0-1.1.39-2.04 1.18-2.81c.79-.77 1.74-1.16 2.85-1.16c1.1 0 2.04.39 2.83 1.16c.78.78 1.18 1.71 1.18 2.8c0 .94-.3 1.79-.89 2.53c-.07.07-.09.13-.07.18c.02.05.07.07.15.07h1.88c.13 0 .21-.05.24-.16c.41-.78.62-1.66.62-2.62c0-.79-.16-1.56-.47-2.29s-.74-1.37-1.27-1.9s-1.16-.95-1.89-1.27a5.67 5.67 0 0 0-2.3-.47c-.8 0-1.57.16-2.3.47c-.73.32-1.36.74-1.89 1.27s-.95 1.16-1.27 1.9s-.49 1.51-.49 2.3m4.96-8.81V4.33c0-.27.1-.5.29-.69s.42-.29.69-.29c.27 0 .5.1.69.29s.29.42.29.69v2.32c0 .27-.1.5-.29.69c-.19.19-.42.29-.69.29c-.27 0-.5-.1-.69-.29a.886.886 0 0 1-.29-.68m6.27 2.58c0-.28.09-.51.26-.67l1.63-1.63c.16-.2.39-.3.68-.3c.27 0 .5.1.68.29c.18.19.27.42.27.69c0 .28-.08.51-.25.68l-1.66 1.63c-.23.19-.46.28-.69.28a.87.87 0 0 1-.66-.28c-.17-.19-.26-.42-.26-.69m2.59 6.23c0-.27.09-.49.26-.67c.17-.18.4-.27.67-.27h2.32c.27 0 .5.09.69.27c.19.18.29.4.29.67s-.1.5-.29.7c-.19.2-.42.3-.69.3h-2.32c-.26 0-.48-.1-.66-.29a.998.998 0 0 1-.27-.71"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Thermometer(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.91 19.56a5.101 5.101 0 0 1 2.24-4.22V5.42c0-.8.27-1.48.82-2.03s1.23-.84 2.03-.84c.81 0 1.49.28 2.04.83c.55.56.83 1.23.83 2.03v9.92c.71.49 1.25 1.11 1.64 1.84s.58 1.53.58 2.38c0 .92-.23 1.78-.68 2.56s-1.07 1.4-1.85 1.85s-1.63.68-2.56.68c-.92 0-1.77-.23-2.55-.68s-1.4-1.07-1.86-1.85s-.68-1.63-.68-2.55m1.76 0c0 .93.33 1.73.98 2.39c.65.66 1.44.99 2.36.99c.93 0 1.73-.33 2.4-1s1.01-1.46 1.01-2.37c0-.62-.16-1.2-.48-1.73c-.32-.53-.76-.94-1.32-1.23l-.28-.14c-.1-.04-.15-.14-.15-.29V5.42c0-.32-.11-.59-.34-.81c-.23-.21-.51-.32-.85-.32c-.32 0-.6.11-.83.32c-.23.21-.34.48-.34.81v10.74c0 .15-.05.25-.14.29l-.27.14c-.55.29-.98.7-1.29 1.23c-.31.53-.46 1.1-.46 1.74m.78 0c0 .71.24 1.32.73 1.82s1.07.75 1.76.75s1.28-.25 1.79-.75s.76-1.11.76-1.81c0-.63-.22-1.19-.65-1.67c-.43-.48-.96-.77-1.58-.85V9.69c0-.06-.03-.13-.1-.19a.299.299 0 0 0-.22-.1c-.09 0-.16.03-.21.08c-.05.06-.08.12-.08.21v7.34c-.61.09-1.13.37-1.56.85c-.43.49-.64 1.04-.64 1.68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThermometerExterior(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.91 19.56a5.101 5.101 0 0 1 2.24-4.22V5.42c0-.8.27-1.48.82-2.03s1.23-.84 2.03-.84c.81 0 1.49.28 2.04.83c.55.56.83 1.23.83 2.03v9.92c.71.49 1.25 1.11 1.64 1.84s.58 1.53.58 2.38c0 .92-.23 1.78-.68 2.56s-1.07 1.4-1.85 1.85s-1.63.68-2.56.68c-.92 0-1.77-.23-2.55-.68s-1.4-1.07-1.86-1.85s-.68-1.63-.68-2.55m1.76 0c0 .93.33 1.73.98 2.39c.65.66 1.44.99 2.36.99c.93 0 1.73-.33 2.4-1s1.01-1.46 1.01-2.37c0-.62-.16-1.2-.48-1.73c-.32-.53-.76-.94-1.32-1.23l-.28-.14c-.1-.04-.15-.14-.15-.29V5.42c0-.32-.11-.59-.34-.81c-.23-.21-.51-.32-.85-.32c-.32 0-.6.11-.83.32c-.23.21-.34.48-.34.81v10.74c0 .15-.05.25-.14.29l-.27.14c-.55.29-.98.7-1.29 1.23c-.31.53-.46 1.1-.46 1.74"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThermometerInternal(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.48 19.56c0 .71.24 1.32.73 1.82s1.07.75 1.76.75s1.28-.25 1.79-.75s.76-1.11.76-1.81c0-.63-.22-1.19-.65-1.67c-.43-.48-.96-.77-1.57-.85V9.69c0-.06-.03-.13-.1-.19a.299.299 0 0 0-.22-.1c-.09 0-.16.03-.21.08c-.05.06-.08.12-.08.21v7.34c-.61.09-1.13.37-1.56.85c-.44.49-.65 1.04-.65 1.68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Thunderstorm(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.63 16.91c0 1.11.33 2.1.99 2.97s1.52 1.47 2.58 1.79l-.66 1.68c-.03.14.02.22.14.22h2.13l-.98 4.3h.28l3.92-5.75c.04-.04.04-.09.01-.14c-.03-.05-.08-.07-.15-.07h-2.18l2.48-4.64c.07-.14.02-.22-.14-.22h-2.94c-.09 0-.17.05-.23.15l-1.07 2.87c-.71-.18-1.3-.57-1.77-1.16s-.7-1.26-.7-2.01c0-.83.28-1.55.85-2.17c.57-.61 1.27-.97 2.1-1.07l.53-.07c.13 0 .2-.06.2-.18l.07-.51c.11-1.08.56-1.99 1.37-2.72c.81-.73 1.76-1.1 2.85-1.1s2.04.37 2.85 1.1c.82.73 1.28 1.64 1.4 2.72l.07.58c0 .11.06.17.18.17h1.6c.91 0 1.68.32 2.32.95c.64.63.97 1.4.97 2.28c0 .85-.3 1.59-.89 2.21c-.59.62-1.33.97-2.2 1.04c-.13 0-.2.06-.2.18v1.37c0 .11.07.17.2.17c1.33-.04 2.46-.55 3.39-1.51s1.39-2.11 1.39-3.45c0-.9-.22-1.73-.67-2.49a4.884 4.884 0 0 0-1.81-1.8c-.77-.44-1.6-.66-2.5-.66h-.31c-.33-1.33-1.04-2.42-2.11-3.26s-2.3-1.27-3.68-1.27c-1.41 0-2.67.44-3.76 1.31s-1.79 1.99-2.1 3.36c-1.11.26-2.02.83-2.74 1.73s-1.08 1.95-1.08 3.1m8.14 9.71c0 .39.19.65.58.77c.01 0 .05 0 .11.01s.11.01.14.01c.17 0 .33-.05.49-.15c.16-.1.27-.26.32-.48l2.25-8.69c.06-.24.04-.45-.07-.65a.827.827 0 0 0-.5-.39l-.26-.03c-.16 0-.32.05-.47.15a.74.74 0 0 0-.31.45l-2.26 8.72c-.01.1-.02.19-.02.28m4.16-3.06c0 .13.03.26.1.38c.14.22.31.37.51.44c.11.03.21.05.3.05s.2-.02.32-.08c.21-.09.35-.28.42-.57l1.44-5.67c.03-.14.05-.23.05-.27c0-.15-.05-.3-.16-.45s-.26-.26-.46-.32l-.26-.03c-.17 0-.33.05-.47.15a.82.82 0 0 0-.3.45l-1.46 5.7c0 .02 0 .05-.01.11c-.02.05-.02.08-.02.11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeEight(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.47c0-2.04.51-3.93 1.52-5.66s2.38-3.1 4.11-4.11s3.61-1.51 5.64-1.51c1.52 0 2.98.3 4.37.89s2.58 1.4 3.59 2.4s1.81 2.2 2.4 3.6s.89 2.85.89 4.39c0 1.52-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.59s-2.2 1.8-3.59 2.39s-2.84.89-4.37.89c-1.53 0-3-.3-4.39-.89s-2.59-1.4-3.6-2.4s-1.8-2.2-2.4-3.58s-.88-2.84-.88-4.37m2.48 0c0 2.37.86 4.43 2.59 6.18c1.73 1.73 3.79 2.59 6.2 2.59c1.58 0 3.05-.39 4.39-1.18s2.42-1.85 3.21-3.2s1.19-2.81 1.19-4.39s-.4-3.05-1.19-4.4s-1.86-2.42-3.21-3.21s-2.81-1.18-4.39-1.18s-3.05.39-4.39 1.18S8.2 8.72 7.4 10.07s-1.18 2.82-1.18 4.4m3.95 2.09c.06-.22.19-.39.38-.51l3.59-2.09V7.81c0-.23.08-.43.24-.59s.36-.24.59-.24s.43.08.59.24s.24.36.24.59v6.67c0 .19-.06.35-.17.5s-.25.24-.42.29l-3.84 2.23c-.12.08-.25.12-.41.12c-.32 0-.56-.14-.72-.42a.893.893 0 0 1-.07-.64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeEleven(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.47c0-2.04.51-3.93 1.52-5.66s2.38-3.1 4.11-4.11s3.61-1.51 5.64-1.51c1.52 0 2.98.3 4.37.89s2.58 1.4 3.59 2.4s1.81 2.2 2.4 3.6s.89 2.85.89 4.39c0 1.52-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.59s-2.2 1.8-3.59 2.39s-2.84.89-4.37.89c-1.53 0-3-.3-4.39-.89s-2.59-1.4-3.6-2.4s-1.8-2.2-2.4-3.58s-.88-2.84-.88-4.37m2.48 0c0 2.37.86 4.43 2.59 6.18c1.73 1.73 3.79 2.59 6.2 2.59c1.58 0 3.05-.39 4.39-1.18s2.42-1.85 3.21-3.2s1.19-2.81 1.19-4.39s-.4-3.05-1.19-4.4s-1.86-2.42-3.21-3.21c-1.35-.79-2.81-1.18-4.39-1.18s-3.05.39-4.39 1.18c-1.35.79-2.42 1.86-3.22 3.21s-1.18 2.82-1.18 4.4m5.62-4.23c.06-.22.19-.39.38-.5c.2-.12.41-.15.64-.09s.4.19.51.38l.78 1.3V7.81c0-.23.08-.43.24-.59s.36-.24.59-.24s.43.08.59.24s.24.36.24.59v6.67c0 .23-.08.43-.24.59s-.36.24-.59.24c-.4 0-.66-.18-.79-.53l-2.26-3.91a.78.78 0 0 1-.09-.63"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeFive(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.47c0-2.04.51-3.93 1.52-5.66s2.38-3.1 4.11-4.11s3.61-1.51 5.64-1.51c1.52 0 2.98.3 4.37.89s2.58 1.4 3.59 2.4s1.81 2.2 2.4 3.6s.89 2.85.89 4.39c0 1.52-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.59s-2.2 1.8-3.59 2.39s-2.84.89-4.37.89c-1.53 0-3-.3-4.39-.89s-2.59-1.4-3.6-2.4s-1.8-2.2-2.4-3.58s-.88-2.84-.88-4.37m2.48 0c0 2.37.86 4.43 2.59 6.18c1.73 1.73 3.79 2.59 6.2 2.59c1.58 0 3.05-.39 4.39-1.18s2.42-1.85 3.21-3.2s1.19-2.81 1.19-4.39s-.4-3.05-1.19-4.4s-1.86-2.42-3.21-3.21s-2.81-1.18-4.39-1.18s-3.05.39-4.39 1.18S8.2 8.72 7.4 10.07s-1.18 2.82-1.18 4.4m7.92 0V7.81c0-.23.08-.43.24-.59s.36-.24.59-.24s.43.08.59.24s.24.36.24.59v6.42l2.15 3.84c.12.21.14.43.08.65s-.19.39-.39.51a.83.83 0 0 1-.41.09c-.33 0-.58-.14-.73-.41l-2.2-3.9a.98.98 0 0 1-.16-.54"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeFour(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.47c0-2.04.51-3.93 1.52-5.66s2.38-3.1 4.11-4.11s3.61-1.51 5.64-1.51c1.52 0 2.98.3 4.37.89s2.58 1.4 3.59 2.4s1.81 2.2 2.4 3.6s.89 2.85.89 4.39c0 1.52-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.59s-2.2 1.8-3.59 2.39s-2.84.89-4.37.89c-1.53 0-3-.3-4.39-.89s-2.59-1.4-3.6-2.4s-1.8-2.2-2.4-3.58s-.88-2.84-.88-4.37m2.48 0c0 2.37.86 4.43 2.59 6.18c1.73 1.73 3.79 2.59 6.2 2.59c1.58 0 3.05-.39 4.39-1.18s2.42-1.85 3.21-3.2s1.19-2.81 1.19-4.39s-.4-3.05-1.19-4.4s-1.86-2.42-3.21-3.21s-2.81-1.18-4.39-1.18s-3.05.39-4.39 1.18S8.2 8.72 7.4 10.07s-1.18 2.82-1.18 4.4m7.92 0V7.81c0-.23.08-.43.24-.59s.36-.24.59-.24s.43.08.59.24s.24.36.24.59v6.15l3.59 2.09c.2.12.32.29.38.51s.03.43-.09.62c-.16.28-.4.42-.72.42c-.17 0-.31-.04-.42-.12l-3.82-2.23c-.17-.05-.31-.15-.42-.29s-.16-.3-.16-.49"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeNine(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.47c0-2.04.51-3.93 1.52-5.66s2.38-3.1 4.11-4.11s3.61-1.51 5.64-1.51c1.52 0 2.98.3 4.37.89s2.58 1.4 3.59 2.4s1.81 2.2 2.4 3.6s.89 2.85.89 4.39c0 1.52-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.59s-2.2 1.8-3.59 2.39s-2.84.89-4.37.89c-1.53 0-3-.3-4.39-.89s-2.59-1.4-3.6-2.4s-1.8-2.2-2.4-3.58s-.88-2.84-.88-4.37m2.48 0c0 2.37.86 4.43 2.59 6.18c1.73 1.73 3.79 2.59 6.2 2.59c1.58 0 3.05-.39 4.39-1.18s2.42-1.85 3.21-3.2s1.19-2.81 1.19-4.39s-.4-3.05-1.19-4.4s-1.86-2.42-3.21-3.21s-2.81-1.18-4.39-1.18s-3.05.39-4.39 1.18S8.2 8.72 7.4 10.07s-1.18 2.82-1.18 4.4m3.29-.01c0-.23.08-.43.24-.59s.36-.24.59-.24h3.79V7.81c0-.23.08-.43.24-.59s.36-.24.59-.24s.43.08.59.24s.24.36.24.59v6.67c0 .23-.08.43-.24.59s-.36.24-.59.24c-.1 0-.16 0-.19-.01h-4.44c-.23 0-.43-.08-.59-.25a.876.876 0 0 1-.23-.59"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeOne(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.47c0-2.04.51-3.93 1.52-5.66s2.38-3.1 4.11-4.11s3.61-1.51 5.64-1.51c1.52 0 2.98.3 4.37.89s2.58 1.4 3.59 2.4s1.81 2.2 2.4 3.6s.89 2.85.89 4.39c0 1.52-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.59s-2.2 1.8-3.59 2.39s-2.84.89-4.37.89c-1.53 0-3-.3-4.39-.89s-2.59-1.4-3.6-2.4s-1.8-2.2-2.4-3.58s-.88-2.84-.88-4.37m2.48 0c0 2.37.86 4.43 2.59 6.18c1.73 1.73 3.79 2.59 6.2 2.59c1.58 0 3.05-.39 4.39-1.18s2.42-1.85 3.21-3.2s1.19-2.81 1.19-4.39s-.4-3.05-1.19-4.4s-1.86-2.42-3.21-3.21s-2.81-1.18-4.39-1.18s-3.05.39-4.39 1.18S8.2 8.72 7.4 10.07s-1.18 2.82-1.18 4.4m7.92 0V7.81c0-.23.08-.43.24-.59s.36-.24.59-.24c.22 0 .42.08.59.24s.25.36.25.59v3.53l.75-1.3c.12-.2.29-.32.52-.38s.44-.03.64.09c.2.11.32.27.39.5s.04.43-.08.63l-2.29 3.91c-.13.35-.38.53-.76.53c-.23 0-.43-.08-.59-.24s-.25-.37-.25-.61"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeSeven(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.47c0-2.04.51-3.93 1.52-5.66s2.38-3.1 4.11-4.11s3.61-1.51 5.64-1.51c1.52 0 2.98.3 4.37.89s2.58 1.4 3.59 2.4s1.81 2.2 2.4 3.6s.89 2.85.89 4.39c0 1.52-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.59s-2.2 1.8-3.59 2.39s-2.84.89-4.37.89c-1.53 0-3-.3-4.39-.89s-2.59-1.4-3.6-2.4s-1.8-2.2-2.4-3.58s-.88-2.84-.88-4.37m2.48 0c0 2.37.86 4.43 2.59 6.18c1.73 1.73 3.79 2.59 6.2 2.59c1.58 0 3.05-.39 4.39-1.18s2.42-1.85 3.21-3.2s1.19-2.81 1.19-4.39s-.4-3.05-1.19-4.4s-1.86-2.42-3.21-3.21s-2.81-1.18-4.39-1.18s-3.05.39-4.39 1.18S8.2 8.72 7.4 10.07s-1.18 2.82-1.18 4.4m5.67 4.24a.828.828 0 0 1 .08-.65l2.17-3.84V7.81c0-.23.08-.43.24-.59s.36-.24.59-.24s.43.08.59.24s.24.36.24.59v6.67c0 .2-.06.37-.19.53l-2.18 3.9c-.16.27-.41.41-.75.41c-.16 0-.29-.03-.4-.09a.967.967 0 0 1-.39-.52"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeSix(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.47c0-2.04.51-3.93 1.52-5.66s2.38-3.1 4.11-4.11s3.61-1.51 5.64-1.51c1.52 0 2.98.3 4.37.89s2.58 1.4 3.59 2.4s1.81 2.2 2.4 3.6s.89 2.85.89 4.39c0 1.52-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.59s-2.2 1.8-3.59 2.39s-2.84.89-4.37.89c-1.53 0-3-.3-4.39-.89s-2.59-1.4-3.6-2.4s-1.8-2.2-2.4-3.58s-.88-2.84-.88-4.37m2.48 0c0 2.37.86 4.43 2.59 6.18c1.73 1.73 3.79 2.59 6.2 2.59c1.58 0 3.05-.39 4.39-1.18s2.42-1.85 3.21-3.2s1.19-2.81 1.19-4.39s-.4-3.05-1.19-4.4s-1.86-2.42-3.21-3.21s-2.81-1.18-4.39-1.18s-3.05.39-4.39 1.18S8.2 8.72 7.4 10.07s-1.18 2.82-1.18 4.4m7.92 4.6V7.81c0-.23.08-.43.24-.59s.36-.24.59-.24s.43.08.59.24s.24.36.24.59v11.26c0 .23-.08.43-.24.6s-.36.25-.59.25s-.43-.08-.59-.25s-.24-.36-.24-.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeTen(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.47c0-2.04.51-3.93 1.52-5.66s2.38-3.1 4.11-4.11s3.61-1.51 5.64-1.51c1.52 0 2.98.3 4.37.89s2.58 1.4 3.59 2.4s1.81 2.2 2.4 3.6s.89 2.85.89 4.39c0 1.52-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.59s-2.2 1.8-3.59 2.39s-2.84.89-4.37.89c-1.53 0-3-.3-4.39-.89s-2.59-1.4-3.6-2.4s-1.8-2.2-2.4-3.58s-.88-2.84-.88-4.37m2.48 0c0 2.37.86 4.43 2.59 6.18c1.73 1.73 3.79 2.59 6.2 2.59c1.58 0 3.05-.39 4.39-1.18s2.42-1.85 3.21-3.2s1.19-2.81 1.19-4.39s-.4-3.05-1.19-4.4s-1.86-2.42-3.21-3.21s-2.81-1.18-4.39-1.18s-3.05.39-4.39 1.18S8.2 8.72 7.4 10.07s-1.18 2.82-1.18 4.4m3.92-2.06a.773.773 0 0 1 .07-.63c.11-.2.28-.33.51-.4s.44-.04.64.07l2.78 1.57V7.81c0-.23.08-.43.24-.59s.36-.24.59-.24s.43.08.59.24s.24.36.24.59v6.67c0 .23-.08.43-.24.59s-.36.24-.59.24a.81.81 0 0 1-.56-.22l-3.88-2.17a.91.91 0 0 1-.39-.51"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeThree(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.47c0-2.04.51-3.93 1.52-5.66s2.38-3.1 4.11-4.11s3.61-1.51 5.64-1.51c1.52 0 2.98.3 4.37.89s2.58 1.4 3.59 2.4s1.81 2.2 2.4 3.6s.89 2.85.89 4.39c0 1.52-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.59s-2.2 1.8-3.59 2.39s-2.84.89-4.37.89c-1.53 0-3-.3-4.39-.89s-2.59-1.4-3.6-2.4s-1.8-2.2-2.4-3.58s-.88-2.84-.88-4.37m2.48 0c0 2.37.86 4.43 2.59 6.18c1.73 1.73 3.79 2.59 6.2 2.59c1.58 0 3.05-.39 4.39-1.18s2.42-1.85 3.21-3.2s1.19-2.81 1.19-4.39s-.4-3.05-1.19-4.4s-1.86-2.42-3.21-3.21s-2.81-1.18-4.39-1.18s-3.05.39-4.39 1.18S8.2 8.72 7.4 10.07s-1.18 2.82-1.18 4.4m7.92 0V7.81c0-.23.08-.43.24-.59s.36-.24.59-.24s.43.08.59.24s.24.36.24.59v5.82h3.78c.23 0 .43.08.59.24s.24.36.24.59c0 .22-.08.42-.24.59c-.16.17-.36.25-.59.25h-4.44c-.03.01-.09.01-.18.01a.806.806 0 0 1-.82-.84"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeTwelve(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.47c0-2.04.51-3.93 1.52-5.66s2.38-3.1 4.11-4.11s3.61-1.51 5.64-1.51c1.52 0 2.98.3 4.37.89s2.58 1.4 3.59 2.4s1.81 2.2 2.4 3.6s.89 2.85.89 4.39c0 1.52-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.59s-2.2 1.8-3.59 2.39s-2.84.89-4.37.89c-1.53 0-3-.3-4.39-.89s-2.59-1.4-3.6-2.4s-1.8-2.2-2.4-3.58s-.88-2.84-.88-4.37m2.48 0c0 2.37.86 4.43 2.59 6.18c1.73 1.73 3.79 2.59 6.2 2.59c1.58 0 3.05-.39 4.39-1.18s2.42-1.85 3.21-3.2s1.19-2.81 1.19-4.39s-.4-3.05-1.19-4.4s-1.86-2.42-3.21-3.21s-2.81-1.18-4.39-1.18s-3.05.39-4.39 1.18S8.2 8.72 7.4 10.07s-1.18 2.82-1.18 4.4m7.92 0c0 .22.08.42.24.59c.16.17.36.25.59.25c.22 0 .42-.08.59-.25c.17-.17.25-.36.25-.59V7.81c0-.23-.08-.43-.25-.59s-.36-.24-.59-.24c-.23 0-.43.08-.59.24s-.24.36-.24.59z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeTwo(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.47c0-2.04.51-3.93 1.52-5.66s2.38-3.1 4.11-4.11s3.61-1.51 5.64-1.51c1.52 0 2.98.3 4.37.89s2.58 1.4 3.59 2.4s1.81 2.2 2.4 3.6s.89 2.85.89 4.39c0 1.52-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.59s-2.2 1.8-3.59 2.39s-2.84.89-4.37.89c-1.53 0-3-.3-4.39-.89s-2.59-1.4-3.6-2.4s-1.8-2.2-2.4-3.58s-.88-2.84-.88-4.37m2.48 0c0 2.37.86 4.43 2.59 6.18c1.73 1.73 3.79 2.59 6.2 2.59c1.58 0 3.05-.39 4.39-1.18s2.42-1.85 3.21-3.2s1.19-2.81 1.19-4.39s-.4-3.05-1.19-4.4s-1.86-2.42-3.21-3.21s-2.81-1.18-4.39-1.18s-3.05.39-4.39 1.18S8.2 8.72 7.4 10.07s-1.18 2.82-1.18 4.4m7.92 0V7.81c0-.23.08-.43.24-.59s.36-.24.59-.24s.43.08.59.24s.24.36.24.59v5.21l2.78-1.57c.2-.12.41-.15.63-.09s.39.2.5.41c.12.2.14.41.08.63s-.19.4-.39.51l-3.88 2.17a.81.81 0 0 1-.56.22a.807.807 0 0 1-.82-.83"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tornado(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.13 15.19c0 .69.36 1.28 1.08 1.77c1.32.93 3.31 1.39 5.98 1.39c1.2 0 2.31-.1 3.34-.31c1.08-.23 1.97-.6 2.65-1.1s1.03-1.08 1.03-1.76c0-.21-.04-.41-.12-.62c1.39-.34 2.48-.8 3.27-1.38s1.19-1.25 1.19-2c0-.19-.03-.39-.09-.6c2.29-.81 3.43-1.9 3.43-3.28c0-.88-.5-1.66-1.49-2.34c-1.95-1.3-4.81-1.95-8.58-1.95c-1.78 0-3.39.16-4.83.47c-1.57.32-2.83.82-3.79 1.5S5.76 6.44 5.76 7.31c0 .52.16.99.48 1.42c-1.18.67-1.77 1.49-1.77 2.46c0 .75.37 1.41 1.1 1.98c-.96.56-1.44 1.23-1.44 2.02m.6 4.5c0 .73.45 1.31 1.35 1.72s2.04.62 3.41.62c1.39 0 2.53-.21 3.44-.62s1.36-.99 1.36-1.72c0-.27-.09-.5-.26-.69s-.4-.28-.67-.28c-.22 0-.42.08-.6.23s-.29.35-.34.57c-.2.16-.56.3-1.1.43s-1.15.2-1.83.2c-1.1 0-2-.16-2.68-.47c.16-.16.24-.36.26-.6s-.04-.45-.15-.62c-.16-.21-.36-.35-.61-.4s-.48 0-.7.13c-.59.41-.88.9-.88 1.5m1.28-4.5c0-.01.06-.07.19-.18c.09-.09.28-.2.56-.34s.61-.25.96-.35l.12-.06c1.62.54 3.51.81 5.67.81c.95 0 1.81-.05 2.58-.16l.26.23c-.09.16-.3.32-.63.5c-.4.21-1.02.41-1.86.57s-1.73.25-2.67.25s-1.83-.08-2.67-.25s-1.47-.36-1.88-.57c-.34-.14-.55-.29-.63-.45m.11 8.42c0 .63.36 1.12 1.08 1.46s1.61.51 2.67.51c1.08 0 1.99-.17 2.72-.51s1.1-.83 1.1-1.46c0-.25-.09-.48-.28-.67s-.41-.29-.66-.29c-.47 0-.78.24-.92.72c-.39.24-1.04.37-1.96.37c-.8 0-1.44-.12-1.92-.37c-.15-.48-.45-.72-.92-.72c-.25 0-.47.09-.64.28s-.27.41-.27.68m.21-12.42c0-.08.05-.17.15-.28c.24-.3.72-.6 1.42-.88c1.92 1.03 4.56 1.54 7.91 1.54c1.71 0 3.32-.16 4.82-.47v.09c0 .15-.09.3-.28.45c-.41.36-1.17.7-2.29 1.03c-1.21.36-2.73.54-4.56.54c-1.84 0-3.36-.18-4.57-.54c-1.16-.32-1.93-.66-2.32-1.02c-.19-.15-.28-.3-.28-.46m1.3-3.88c0-.18.12-.37.35-.59c.45-.42 1.35-.82 2.68-1.21c1.43-.42 3.14-.63 5.14-.63c2.01 0 3.74.21 5.19.63c1.35.39 2.24.8 2.68 1.22c.22.22.34.42.34.59s-.11.35-.34.56c-.44.42-1.33.83-2.68 1.23c-1.45.42-3.17.63-5.19.63c-2 0-3.72-.21-5.14-.63c-1.34-.4-2.24-.81-2.68-1.24c-.23-.21-.35-.39-.35-.56"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Train(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.25 12.68v-.32c0-.1.03-.18.1-.25s.15-.1.25-.1h7.58c.1 0 .18.03.25.1s.1.15.1.25v.32c0 .1-.03.18-.1.25s-.15.1-.25.1h-.44v1.65h2.12c.02-.28.14-.52.35-.71c.21-.19.46-.29.75-.29s.53.1.74.29c.21.19.32.43.35.71h1.32v-3.39a.52.52 0 0 1-.35-.16c-.1-.1-.15-.23-.15-.37v-.31c0-.14.05-.27.16-.38s.24-.16.39-.16h1.99c.15 0 .28.05.38.16s.15.23.15.38v.31c0 .14-.05.27-.14.37c-.09.1-.2.16-.34.16v3.39h1.56c.27 0 .51.1.71.3s.3.44.3.71v2.93l3.73 4.87h-4.74v-3.04h-.71c.11.26.16.54.16.83c0 .61-.21 1.12-.64 1.56c-.43.43-.95.65-1.55.65c-.61 0-1.12-.22-1.56-.65a2.13 2.13 0 0 1-.65-1.56c0-.29.05-.57.16-.83h-1c.11.27.17.55.17.83c0 .61-.22 1.12-.65 1.56s-.95.65-1.56.65c-.61 0-1.12-.22-1.55-.65s-.64-.95-.64-1.56c0-.29.05-.57.16-.83H9.97c.12.29.18.57.18.83c0 .61-.22 1.12-.65 1.56s-.95.65-1.56.65s-1.12-.22-1.56-.65s-.65-.95-.65-1.56c0-.29.06-.57.17-.84c-.24-.04-.45-.15-.61-.34s-.24-.41-.24-.66v-.86h-.02v-5.55H4.6c-.1 0-.18-.03-.25-.1a.332.332 0 0 1-.1-.25m2.05 3.94c0 .21.07.39.22.54c.15.15.33.22.54.22H8.5c.21 0 .39-.07.53-.22s.22-.33.22-.54v-2.3a.71.71 0 0 0-.22-.53a.71.71 0 0 0-.53-.22H7.07c-.21 0-.39.07-.54.23c-.15.15-.22.32-.22.52v2.3zm9.48-11.19c0 .41.16.76.47 1.04c0 .2.09.43.26.68s.36.4.56.44c.04.22.15.41.31.57c.16.15.36.25.59.3c-.11.11-.16.24-.16.39c0 .18.06.33.18.45s.27.18.45.18s.33-.06.46-.19c.13-.12.19-.28.19-.45c0-.02 0-.05-.01-.09c-.01-.04-.01-.08-.01-.1h.03c.21 0 .39-.08.54-.23c.15-.15.23-.34.23-.55c0-.1-.04-.22-.12-.38c.17-.09.31-.25.41-.47h.45c.39-.02.73-.17 1-.45c.28-.28.42-.61.42-1.01c0-.34-.11-.64-.33-.9c-.22-.26-.5-.43-.83-.52c-.08-.4-.29-.73-.62-.99s-.71-.39-1.12-.39c-.41 0-.77.13-1.08.38c-.31.25-.52.58-.62.97h-.11c-.41 0-.77.13-1.08.39c-.31.25-.46.57-.46.93"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tsunami(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.07 21.24c0-.16.06-.3.17-.42c.12-.12.25-.18.41-.18h.4c-.21-.66-.39-1.35-.53-2.07c-.21-1.1-.32-2.1-.32-2.99c0-1.71.3-3.32.91-4.81c.62-1.46 1.48-2.71 2.59-3.76c1.12-1.06 2.42-1.87 3.9-2.42c1.51-.57 3.14-.86 4.91-.86c1.06 0 2.06.09 3 .28c.94.22 1.85.56 2.73 1.03l1.7.91l-1.88.39c-.58.13-.98.39-1.2.78c-.16.32-.15.69.03 1.11l.41.95l-1.02.05c-.43.03-.83.12-1.18.27c-.33.16-.52.32-.58.5c-.11.23.01.56.36 1l.81.96l-1.26.18c-1.55.23-2.82.55-3.81.96s-1.77.94-2.35 1.59c-.56.62-.98 1.42-1.25 2.37c-.27.96-.42 2.15-.45 3.59h5.26v-2.78l-.38.23a.55.55 0 0 1-.45.07a.532.532 0 0 1-.37-.28c-.09-.14-.11-.29-.08-.45s.12-.29.27-.38l3.82-2.38l.02-.02c.01 0 .01 0 .01-.01h.02c.01 0 .02 0 .03-.01c.07-.02.14-.05.23-.07h.06c.01.01.02.02.03.02h.07c0 .01.01.01.02.01h.03l.02.01h.02c.01.01.02.02.02.03h.02c.01 0 .01 0 .01.01c.02 0 .03 0 .03.01c.01 0 .02 0 .03.01l.02.01l3.82 2.35c.14.09.23.22.27.38c.03.17.01.32-.08.46c-.08.14-.2.23-.37.26s-.32.01-.45-.08l-.31-.19v2.77h.96c.16 0 .29.06.4.18c.11.12.16.26.16.42c.01.17-.05.31-.16.43s-.25.18-.4.18H5.65a.58.58 0 0 1-.58-.6m1.55-5.66c0 .71.1 1.62.3 2.73c.15.81.33 1.52.54 2.12h2.69c.05-1.45.2-2.65.45-3.59c.35-1.27.88-2.31 1.6-3.09c.73-.82 1.69-1.47 2.89-1.96c.82-.34 1.86-.63 3.11-.87l-.08-.25c-.1-.46-.07-.87.09-1.23c.22-.51.65-.92 1.28-1.21c.07-.03.13-.06.19-.07c-.86-.2-1.73-.25-2.6-.14c-.99.12-1.92.41-2.78.85c-1.11.58-2.11 1.41-3.01 2.48c-.1.12-.23.18-.38.18a.48.48 0 0 1-.31-.1c-.1-.09-.16-.2-.17-.34s.02-.26.11-.37c1-1.19 2.11-2.1 3.34-2.73c.98-.49 2.03-.81 3.14-.95a7.19 7.19 0 0 1 1.67-.02c.54.06.92.12 1.14.17s.37.09.45.12l.08.03c.05-.31.13-.59.24-.84c.16-.29.37-.56.64-.8c-.3-.09-.65-.19-1.04-.28c-.8-.18-1.7-.26-2.69-.26c-1.58 0-3.05.26-4.42.77c-1.34.51-2.48 1.22-3.42 2.14c-.98.91-1.73 2-2.23 3.26c-.55 1.32-.82 2.74-.82 4.25m11.43 5.06h3.88v-3.52l-1.98-1.21l-1.9 1.19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Umbrella(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.64 14.65c.01-.34.18-.86.5-1.58c.32-.72.76-1.48 1.33-2.3c1.86-2.61 4.49-3.98 7.88-4.13v-.43c0-.21.07-.37.21-.5c.14-.13.3-.19.48-.19c.19 0 .35.06.5.19c.14.13.22.3.22.5v.44c.98.04 1.9.19 2.75.45c.85.26 1.59.59 2.22 1c.63.41 1.17.83 1.61 1.27c.45.43.85.9 1.2 1.41c.41.59.77 1.23 1.06 1.9c.29.67.5 1.21.61 1.61c.11.4.17.6.18.61v.19c0 .18-.07.32-.21.44s-.3.17-.49.17c-.31 0-.51-.09-.6-.26c-.78-.88-1.63-1.31-2.55-1.31c-.34.02-.69.1-1.03.23c-.34.13-.62.27-.82.42c-.21.14-.4.29-.58.44c-.18.15-.27.22-.28.23c-.19.17-.37.26-.53.26c-.23 0-.4-.06-.52-.18c-.73-.73-1.39-1.17-2.01-1.32v7.74l-.01.21l-.04.23l-.06.25l-.09.26l-.13.27l-.17.26l-.21.25c-.51.59-1.23.88-2.18.88c-1.01 0-1.77-.29-2.28-.88c-.12-.12-.22-.25-.31-.38c-.09-.14-.16-.27-.21-.41c-.05-.13-.09-.26-.12-.38s-.05-.24-.06-.36c-.02-.15-.03-.25-.03-.33v-.23c0-.07.01-.12.01-.13c0-.18.08-.34.23-.47a.64.64 0 0 1 .55-.14c.18 0 .32.08.44.23s.18.34.18.56c-.06.41.02.76.25 1.05c.21.29.65.44 1.32.44c.52 0 .9-.12 1.13-.36c.13-.13.23-.29.29-.48c.06-.19.09-.34.08-.47l-.01-.19v-7.36c-.73.18-1.38.56-1.93 1.14c-.04.08-.12.16-.23.23s-.21.11-.3.11c-.18 0-.38-.11-.6-.34c-.8-.89-1.65-1.33-2.55-1.31c-.4.01-.78.07-1.12.2c-.35.13-.61.26-.79.39c-.17.13-.36.28-.54.45s-.29.27-.32.29c-.21.14-.38.22-.51.22s-.3-.06-.48-.17c-.16-.1-.26-.21-.3-.32c-.03-.12-.04-.29-.03-.51m2.09-1.42c.68-.36 1.32-.53 1.92-.53h.08c1.15 0 2.2.44 3.15 1.33c.38-.33.84-.62 1.39-.88c.54-.26 1.13-.41 1.77-.45h.08c1.15 0 2.2.44 3.15 1.33c.38-.33.84-.62 1.39-.88c.54-.26 1.13-.41 1.77-.45h.09c.56 0 1.15.15 1.75.44c-.44-.86-.74-1.41-.88-1.66c-1.79-2.34-4.27-3.51-7.43-3.51c-1.58 0-2.99.3-4.24.9c-1.24.6-2.26 1.47-3.05 2.61c-.23.34-.55.93-.94 1.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Volcano(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.39 22.61c-.12-.27-.09-.54.09-.81l1.4-2.67c.01-.04.05-.09.11-.15c.04-.04.17-.14.38-.29c.02-.01.25-.18.68-.5c.48-.32 1.03-.72 1.68-1.19l1.8-2.98c.17-.27.41-.41.72-.41h.7c-.16.19-.31.39-.45.6c-.14.21-.27.5-.38.85c-.12.36-.18.71-.18 1.07c0 .36.09.77.28 1.25c.19.47.48.94.88 1.39c.27.31.44.62.5.93s.02.58-.1.83c-.12.25-.32.5-.59.74c-.27.24-.56.45-.88.63c-.32.18-.68.35-1.07.52c-.39.17-.75.3-1.05.41c-.31.1-.62.2-.93.29H8.16c-.16 0-.32-.05-.46-.14a.703.703 0 0 1-.31-.37m2.75-14.87c0-.46.15-.88.45-1.24c.3-.37.69-.6 1.16-.72c.11-.56.4-1.02.85-1.38s.98-.54 1.56-.54c.56 0 1.06.17 1.5.52s.73.8.86 1.35h.14c.57 0 1.07.18 1.5.54c.42.36.64.79.64 1.3c0 .56-.22 1.03-.65 1.43c0 .27-.12.59-.36.93c-.24.35-.5.55-.78.61c-.06.31-.21.57-.43.78c-.23.22-.5.36-.82.43c.15.16.22.34.22.54c0 .25-.09.46-.26.63a.89.89 0 0 1-.64.25c-.24 0-.45-.09-.63-.26a.817.817 0 0 1-.26-.62c0-.03.01-.08.02-.14s.02-.11.02-.13h-.03c-.29 0-.54-.11-.75-.32a1.03 1.03 0 0 1-.32-.75c0-.12.06-.3.18-.53c-.24-.12-.43-.33-.57-.63h-.63a2.21 2.21 0 0 1-1.39-.65c-.38-.38-.58-.85-.58-1.4m4.62 7.74c0-.16.02-.34.07-.54s.11-.35.16-.47s.12-.27.21-.45s.15-.31.19-.41h.38c.28 0 .49.11.66.32l.07.1l1.31 2.48l4.65 5.23l.04.03c.21.27.24.56.08.88c-.15.31-.4.46-.75.46H16.2c.17-.16.32-.29.44-.39c.12-.11.27-.27.45-.49s.33-.43.42-.61s.17-.42.23-.69c.06-.27.07-.53.01-.79c-.06-.25-.18-.53-.38-.84c-.19-.31-.46-.61-.81-.91c-.34-.3-.64-.59-.88-.88c-.24-.28-.43-.54-.56-.76c-.13-.22-.23-.45-.29-.68c-.04-.23-.07-.43-.07-.59"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindBeaufortEight(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.99 13.5c0 .18.06.31.17.4c.12.11.26.17.42.17h10.4c.18 0 .33.06.46.19c.13.12.2.28.2.46s-.07.34-.2.47a.629.629 0 0 1-.93-.01a.603.603 0 0 0-.4-.16c-.16 0-.3.05-.41.16c-.11.11-.16.24-.16.39c0 .16.06.3.17.41c.36.36.78.53 1.28.53c.49 0 .91-.17 1.26-.52s.52-.77.52-1.27c0-.49-.17-.92-.52-1.27s-.77-.53-1.26-.53H5.58c-.16 0-.3.06-.42.17c-.11.12-.17.25-.17.41m0-2.02c0 .17.06.3.17.39c.12.11.26.16.42.16h13.81c.49 0 .92-.18 1.27-.52c.35-.35.52-.77.52-1.27c0-.49-.17-.91-.52-1.26s-.77-.52-1.27-.52c-.49 0-.91.17-1.27.51c-.11.12-.16.27-.16.42c0 .16.05.3.16.4c.11.1.24.15.4.15c.15 0 .29-.05.41-.16a.61.61 0 0 1 .45-.18c.17 0 .33.06.46.18s.2.27.2.45s-.07.34-.2.47c-.13.13-.28.2-.46.2H5.58c-.16 0-.3.06-.42.17c-.11.11-.17.25-.17.41m13 8.19c0 .73.29 1.29.86 1.66c.57.38 1.34.57 2.31.57c.59 0 1.12-.06 1.57-.18c.46-.12.81-.27 1.07-.44s.46-.38.62-.62c.16-.24.26-.46.31-.66c.05-.2.08-.4.08-.61c0-.41-.12-.77-.36-1.06c-.24-.3-.55-.49-.94-.57l.02-.03v.01c.45-.06.82-.26 1.12-.6c.29-.33.44-.73.44-1.19c0-.38-.09-.71-.26-.98s-.41-.48-.71-.61c-.3-.14-.61-.24-.92-.3c-.31-.06-.65-.09-1.01-.09c-.48 0-.9.05-1.28.14c-.38.09-.69.22-.93.37c-.24.15-.43.33-.59.53s-.27.4-.33.6c-.06.2-.09.41-.09.62c0 .34.09.64.27.9c.18.26.43.43.75.53v.03c-.56.06-1.04.27-1.42.61c-.39.37-.58.82-.58 1.37m2.11-.23c0-.35.14-.61.42-.77s.62-.24 1.01-.24c.41 0 .7.09.89.28c.18.18.28.38.28.6v.13c0 .28-.13.49-.38.64c-.25.14-.58.22-.97.22l.03-.01c-.14 0-.27-.01-.4-.03s-.27-.06-.41-.11a.797.797 0 0 1-.34-.26a.847.847 0 0 1-.13-.45m.76-3.07c0-.32.12-.55.37-.69s.55-.22.9-.22c.3 0 .55.07.76.2s.31.35.31.63c0 .07-.02.15-.05.23c-.03.08-.09.17-.17.27c-.08.1-.21.18-.39.24c-.18.06-.4.09-.66.09c-.4 0-.68-.08-.84-.23a.698.698 0 0 1-.23-.52"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindBeaufortEleven(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.68 13.5c0 .18.06.31.17.4c.12.11.26.17.42.17h10.4c.17 0 .33.06.46.19c.13.12.2.28.2.46s-.07.34-.2.47a.629.629 0 0 1-.93-.01a.603.603 0 0 0-.4-.16c-.16 0-.3.05-.41.16c-.11.11-.16.24-.16.39c0 .16.06.3.17.41c.36.36.78.53 1.28.53s.91-.17 1.26-.52s.52-.77.52-1.27c0-.49-.17-.92-.52-1.27s-.77-.53-1.26-.53H5.27c-.16 0-.3.06-.42.17c-.11.12-.17.25-.17.41m0-2.02c0 .17.06.3.17.39c.12.11.26.16.42.16h13.81c.49 0 .92-.18 1.27-.52c.35-.35.52-.77.52-1.27c0-.49-.17-.91-.52-1.26s-.77-.52-1.27-.52c-.49 0-.91.17-1.27.51c-.11.12-.16.27-.16.42c0 .16.05.3.16.4c.11.1.24.15.4.15c.15 0 .29-.05.41-.16a.61.61 0 0 1 .45-.18c.17 0 .33.06.46.18s.2.27.2.45s-.07.34-.2.47c-.13.13-.28.2-.46.2H5.27c-.16 0-.3.06-.42.17c-.11.11-.17.25-.17.41M17.57 21.9h2.47l1.65-7.99h-2.47zm3.73 0h2.46l1.65-7.99h-2.45z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindBeaufortFive(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.97 13.5c0 .18.06.31.17.4c.12.11.26.17.42.17h10.4c.18 0 .33.06.46.19c.13.12.2.28.2.46s-.07.34-.2.47a.629.629 0 0 1-.93-.01a.603.603 0 0 0-.4-.16c-.16 0-.3.05-.41.16c-.11.11-.16.24-.16.39c0 .16.06.3.17.41c.36.36.78.53 1.28.53c.49 0 .91-.17 1.26-.52s.52-.77.52-1.27c0-.49-.17-.92-.52-1.27s-.77-.53-1.26-.53H5.56c-.16 0-.3.06-.42.17c-.11.12-.17.25-.17.41m0-2.02c0 .17.06.3.17.39c.12.11.26.16.42.16h13.81c.49 0 .92-.18 1.27-.52c.35-.35.52-.77.52-1.27c0-.49-.17-.91-.52-1.26s-.77-.52-1.27-.52c-.49 0-.91.17-1.27.51a.64.64 0 0 0-.16.41c0 .16.05.3.16.4c.11.1.24.15.4.15c.15 0 .29-.05.41-.16a.61.61 0 0 1 .45-.18c.17 0 .33.06.46.18s.2.27.2.45s-.07.34-.2.47c-.13.13-.28.2-.46.2H5.56c-.16 0-.3.06-.42.17c-.11.12-.17.26-.17.42m13.07 7.9c-.02.32.01.62.12.91c.1.29.27.56.5.81c.23.25.55.44.98.59c.42.15.92.22 1.49.22c.58 0 1.09-.08 1.53-.23s.8-.34 1.05-.57c.25-.22.45-.49.61-.79c.16-.3.27-.57.32-.82c.05-.25.08-.49.08-.74c0-.67-.21-1.21-.64-1.61s-.98-.61-1.65-.61c-.69 0-1.18.14-1.45.43h-.02l.35-1.02h3.45l.39-1.88h-5.24l-1.45 4.46h2c.16-.34.53-.51 1.11-.51c.32 0 .58.08.77.25c.19.17.29.41.29.75s-.12.61-.35.82c-.23.21-.57.31-1 .31c-.31 0-.56-.06-.73-.17c-.21-.11-.33-.31-.36-.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindBeaufortFour(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.98 13.5c0 .18.06.31.17.4c.12.11.26.17.42.17h10.4c.17 0 .33.06.46.19c.13.12.2.28.2.46s-.07.34-.2.47a.629.629 0 0 1-.93-.01a.584.584 0 0 0-.4-.16c-.16 0-.3.05-.41.16c-.11.11-.16.24-.16.39c0 .16.06.3.17.41c.36.36.78.53 1.27.53s.91-.17 1.26-.52s.52-.77.52-1.27c0-.49-.17-.92-.52-1.27s-.77-.53-1.26-.53H5.57c-.16 0-.3.06-.42.17c-.11.12-.17.25-.17.41m0-2.02c0 .17.06.3.17.39c.12.11.26.16.42.16h13.81c.49 0 .92-.18 1.27-.52c.35-.35.52-.77.52-1.27c0-.49-.17-.91-.52-1.26s-.77-.52-1.27-.52c-.49 0-.91.17-1.27.51c-.11.12-.16.27-.16.42c0 .16.05.3.16.4c.11.1.24.15.4.15c.15 0 .29-.05.41-.16a.61.61 0 0 1 .45-.18c.17 0 .33.06.46.18s.2.27.2.45s-.07.34-.2.47c-.13.13-.28.2-.46.2H5.57c-.16 0-.3.06-.42.17c-.11.11-.17.25-.17.41m13 8.87h3.57l-.32 1.55h2.2l.36-1.55h1.01l.36-1.9h-1l.9-4.34h-2.22l-4.43 4.16zm2.25-1.9l2.24-2.21h.03l-.49 2.21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindBeaufortNine(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.09 13.5c0 .18.06.31.17.4c.12.11.26.17.42.17h10.4c.18 0 .33.06.46.19c.13.12.2.28.2.46s-.07.34-.2.47a.629.629 0 0 1-.93-.01a.603.603 0 0 0-.4-.16c-.16 0-.3.05-.41.16c-.11.11-.16.24-.16.39c0 .16.06.3.17.41c.36.36.78.53 1.27.53s.91-.17 1.26-.52s.52-.77.52-1.27c0-.49-.17-.92-.52-1.27s-.77-.53-1.26-.53H5.68c-.16 0-.3.06-.42.17c-.12.12-.17.25-.17.41m0-2.02c0 .17.06.3.17.39c.12.11.26.16.42.16h13.81c.49 0 .92-.18 1.27-.52c.35-.35.52-.77.52-1.27c0-.49-.17-.91-.52-1.26s-.77-.52-1.27-.52c-.49 0-.91.17-1.27.51c-.11.12-.16.27-.16.42c0 .16.05.3.16.4c.11.1.24.15.4.15c.15 0 .29-.05.41-.16a.61.61 0 0 1 .45-.18c.17 0 .33.06.46.18s.2.27.2.45s-.07.34-.2.47c-.13.13-.28.2-.46.2H5.68c-.16 0-.3.06-.42.17a.53.53 0 0 0-.17.41m13.09 8.28c.02.39.11.73.28 1.02c.17.29.39.51.67.67c.28.16.58.27.9.34s.67.11 1.04.11c.57 0 1.09-.11 1.55-.32c.47-.21.84-.48 1.13-.81c.29-.33.53-.7.73-1.13s.33-.84.41-1.23s.12-.78.12-1.15c0-1.06-.27-1.87-.81-2.43c-.54-.57-1.26-.85-2.17-.85c-.93 0-1.72.28-2.36.85c-.64.57-.97 1.32-.97 2.24c0 .66.21 1.2.63 1.62c.42.42.96.63 1.63.63c.36 0 .7-.07 1.05-.22c.34-.14.58-.33.72-.54h.03c-.12.48-.31.88-.58 1.22c-.27.34-.62.51-1.06.51c-.29 0-.48-.03-.59-.1a.743.743 0 0 1-.24-.42h-2.11zm2.57-2.88c0-.31.1-.58.29-.81c.19-.23.48-.34.86-.34c.34 0 .6.09.77.26c.18.17.27.43.27.76c0 .09-.02.2-.06.31s-.1.23-.18.36c-.08.12-.2.23-.37.31s-.35.12-.56.12s-.39-.04-.54-.11a.7.7 0 0 1-.32-.29c-.07-.12-.11-.22-.14-.31s-.02-.18-.02-.26"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindBeaufortOne(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.76 13.5c0 .18.06.31.17.4c.12.11.26.17.42.17h10.4c.17 0 .33.06.46.19c.13.12.2.28.2.46s-.07.34-.2.47a.629.629 0 0 1-.93-.01a.584.584 0 0 0-.4-.16c-.16 0-.3.05-.41.16c-.11.11-.16.24-.16.39c0 .16.06.3.17.41c.36.36.78.53 1.27.53s.91-.17 1.26-.52s.52-.77.52-1.27c0-.49-.17-.92-.52-1.27s-.77-.53-1.26-.53H6.35c-.16 0-.3.06-.42.17c-.12.12-.17.25-.17.41m0-2.02c0 .17.06.3.17.39c.12.11.26.16.42.16h13.81c.49 0 .92-.18 1.27-.52c.35-.35.52-.77.52-1.27c0-.49-.17-.91-.52-1.26s-.77-.52-1.27-.52c-.49 0-.91.17-1.27.51c-.11.12-.16.27-.16.42c0 .16.05.3.16.4c.11.1.24.15.4.15c.15 0 .29-.05.41-.16a.61.61 0 0 1 .45-.18c.17 0 .33.06.46.18s.2.27.2.45s-.07.34-.2.47c-.13.13-.28.2-.46.2H6.35c-.16 0-.3.06-.42.17a.53.53 0 0 0-.17.41m12.89 10.37h2.47l1.65-7.98H20.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindBeaufortSeven(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.83 13.5c0 .18.06.31.17.4c.12.11.26.17.42.17h10.4c.17 0 .33.06.46.19c.13.12.2.28.2.46s-.07.34-.2.47a.629.629 0 0 1-.93-.01a.584.584 0 0 0-.4-.16c-.16 0-.3.05-.41.16c-.11.11-.16.24-.16.39c0 .16.06.3.17.41c.36.36.78.53 1.27.53s.91-.17 1.26-.52s.52-.77.52-1.27c0-.49-.17-.92-.52-1.27s-.77-.53-1.26-.53H5.42c-.16 0-.3.06-.42.17c-.11.12-.17.25-.17.41m0-2.02c0 .17.06.3.17.39c.12.11.26.16.42.16h13.81c.49 0 .92-.18 1.27-.52c.35-.35.52-.77.52-1.27c0-.49-.17-.91-.52-1.26s-.77-.52-1.27-.52c-.49 0-.91.17-1.27.51c-.11.12-.16.27-.16.42c0 .16.05.3.16.4c.11.1.24.15.4.15c.15 0 .29-.05.41-.16a.61.61 0 0 1 .45-.18c.17 0 .33.06.46.18s.2.27.2.45s-.07.34-.2.47c-.13.13-.28.2-.46.2H5.42c-.16 0-.3.06-.42.17c-.11.11-.17.25-.17.41M18.85 21.9h2.47c.26-1.29.73-2.45 1.39-3.47c.67-1.02 1.39-1.84 2.16-2.44l.38-1.87h-5.96l-.41 1.89h3.49c-1.98 1.99-3.15 3.95-3.52 5.89"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindBeaufortSix(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.92 13.5c0 .18.06.31.17.4c.12.11.26.17.42.17h10.4c.18 0 .33.06.46.19c.13.12.2.28.2.46s-.07.34-.2.47a.629.629 0 0 1-.93-.01a.603.603 0 0 0-.4-.16c-.16 0-.3.05-.41.16c-.11.11-.16.24-.16.39c0 .16.06.3.17.41c.36.36.78.53 1.28.53c.49 0 .91-.17 1.26-.52s.52-.77.52-1.27c0-.49-.17-.92-.52-1.27s-.77-.53-1.26-.53H5.51c-.16 0-.3.06-.42.17c-.11.12-.17.25-.17.41m0-2.02c0 .17.06.3.17.39c.12.11.26.16.42.16h13.81c.49 0 .92-.18 1.27-.52c.35-.35.52-.77.52-1.27c0-.49-.17-.91-.52-1.26s-.77-.52-1.27-.52c-.49 0-.91.17-1.27.51c-.11.12-.16.27-.16.42c0 .16.05.3.16.4c.11.1.24.15.4.15c.15 0 .29-.05.41-.16a.61.61 0 0 1 .45-.18c.17 0 .33.06.46.18s.2.27.2.45s-.07.34-.2.47c-.13.13-.28.2-.46.2H5.51c-.16 0-.3.06-.42.17c-.11.11-.17.25-.17.41m13.41 7.24c0 .96.25 1.73.75 2.31c.5.58 1.26.87 2.29.87c.95 0 1.73-.29 2.35-.87c.62-.58.92-1.34.92-2.28c0-.64-.22-1.17-.67-1.57s-.99-.6-1.65-.6c-.73 0-1.3.25-1.72.75h-.02c.33-1.16.88-1.74 1.65-1.74c.25 0 .44.05.58.14c.12.09.2.22.23.41h2.11a2.25 2.25 0 0 0-.19-.84c-.12-.25-.26-.46-.44-.62s-.39-.3-.63-.4c-.24-.11-.49-.18-.75-.23c-.26-.04-.52-.07-.8-.07c-.62 0-1.18.12-1.68.36s-.88.54-1.17.89c-.28.35-.52.75-.71 1.2c-.19.45-.31.86-.38 1.23c-.04.38-.07.73-.07 1.06m2.05.45c0-.37.12-.65.37-.84c.24-.19.52-.29.82-.29c.19 0 .35.03.48.08s.23.12.3.19c.07.07.12.16.15.27c.04.11.06.19.07.25c.01.06.01.12.01.19c0 .31-.11.58-.32.79s-.5.32-.85.32c-.31 0-.55-.09-.75-.27a.91.91 0 0 1-.28-.69"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindBeaufortTen(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.15 13.5c0 .18.06.31.17.4c.12.11.26.17.42.17h10.4c.18 0 .33.06.46.19c.13.12.2.28.2.46s-.07.34-.2.47a.629.629 0 0 1-.93-.01a.584.584 0 0 0-.4-.16c-.16 0-.3.05-.41.16c-.11.11-.16.24-.16.39c0 .16.06.3.17.41c.36.36.78.53 1.27.53s.91-.17 1.26-.52s.52-.77.52-1.27c0-.49-.17-.92-.52-1.27s-.77-.53-1.26-.53H3.75c-.16 0-.3.06-.42.17c-.12.12-.18.25-.18.41m0-2.02c0 .17.06.3.17.39c.12.11.26.16.42.16h13.81c.49 0 .92-.18 1.27-.52c.35-.35.52-.77.52-1.27c0-.49-.17-.91-.52-1.26s-.77-.52-1.27-.52c-.49 0-.91.17-1.27.51c-.11.12-.16.27-.16.42c0 .16.05.3.16.4c.11.1.24.15.4.15c.15 0 .29-.05.41-.16a.61.61 0 0 1 .45-.18c.17 0 .33.06.46.18s.2.27.2.45s-.07.34-.2.47c-.13.13-.28.2-.46.2H3.75c-.16 0-.3.06-.42.17c-.12.11-.18.25-.18.41M15.97 21.8h2.46l1.64-7.94h-2.45zm4.19-2.92c0 .52.08.98.24 1.38s.38.72.66.95c.27.23.58.4.9.52s.68.17 1.05.17c.61 0 1.16-.12 1.64-.38c.48-.25.86-.56 1.13-.93c.27-.37.5-.79.68-1.25c.18-.47.3-.89.37-1.27c.06-.38.09-.73.09-1.05c0-.97-.27-1.72-.8-2.25s-1.24-.8-2.13-.8c-1.03 0-1.93.46-2.7 1.37c-.75.92-1.13 2.1-1.13 3.54m2.05.1c0-.16.01-.35.04-.59c.03-.23.08-.51.16-.84c.08-.32.18-.62.3-.9c.12-.27.29-.5.52-.69c.22-.19.47-.29.75-.29c.27 0 .48.09.65.27c.16.18.24.44.24.79c0 .96-.17 1.78-.5 2.45s-.75 1.01-1.23 1.01c-.62 0-.93-.4-.93-1.21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindBeaufortThree(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.03 13.5c0 .18.06.31.17.4c.12.11.26.17.42.17h10.4c.17 0 .33.06.46.19c.13.12.2.28.2.46s-.07.34-.2.47a.629.629 0 0 1-.93-.01a.584.584 0 0 0-.4-.16c-.16 0-.3.05-.41.16c-.11.11-.16.24-.16.39c0 .16.06.3.17.41c.36.36.78.53 1.27.53s.91-.17 1.26-.52s.52-.77.52-1.27c0-.49-.17-.92-.52-1.27s-.77-.53-1.26-.53H5.62c-.16 0-.3.06-.42.17c-.11.12-.17.25-.17.41m0-2.02c0 .17.06.3.17.39c.12.11.26.16.42.16h13.81c.49 0 .92-.18 1.27-.52c.35-.35.52-.77.52-1.27c0-.49-.17-.91-.52-1.26s-.77-.52-1.27-.52c-.49 0-.91.17-1.27.51a.64.64 0 0 0-.16.41c0 .16.05.3.16.4c.11.1.24.15.4.15c.15 0 .29-.05.41-.16a.61.61 0 0 1 .45-.18c.17 0 .33.06.46.18s.2.27.2.45s-.07.34-.2.47c-.13.13-.28.2-.46.2H5.62c-.16 0-.3.06-.42.17c-.11.12-.17.26-.17.42m13.09 8.04c0 .27.05.53.16.79s.27.5.5.75c.23.24.55.43.96.58s.9.22 1.46.22c1.21 0 2.08-.24 2.63-.72c.55-.48.82-1.13.82-1.95c0-.36-.1-.69-.3-.99c-.2-.3-.47-.47-.79-.51v-.02c.43-.08.79-.27 1.07-.58c.28-.31.43-.69.43-1.12c0-.31-.06-.58-.17-.82c-.11-.24-.26-.43-.44-.58c-.18-.15-.39-.27-.64-.37c-.25-.1-.5-.16-.75-.2c-.25-.04-.52-.06-.8-.06c-.92 0-1.68.22-2.28.67c-.59.45-.96 1.12-1.1 2.01h2.03c.04-.31.17-.55.38-.72c.21-.17.47-.26.78-.26c.29 0 .51.06.68.18s.25.29.25.5c0 .47-.42.7-1.27.7h-.47l-.29 1.4h.44c.68 0 1.02.23 1.02.7c0 .31-.11.55-.34.72c-.23.17-.5.25-.83.25c-.38 0-.66-.11-.83-.34c-.17-.21-.24-.51-.21-.89h-2.07c-.02.2-.03.41-.03.66"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindBeaufortTwelve(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.07 13.5c0 .18.06.31.17.4c.12.11.26.17.42.17h10.4c.17 0 .33.06.46.19c.13.12.2.28.2.46s-.07.34-.2.47a.629.629 0 0 1-.93-.01a.603.603 0 0 0-.4-.16c-.16 0-.3.05-.41.16c-.11.11-.16.24-.16.39c0 .16.06.3.17.41c.36.36.78.53 1.28.53s.91-.17 1.26-.52s.52-.77.52-1.27c0-.49-.17-.92-.52-1.27s-.77-.53-1.26-.53H3.66c-.16 0-.3.06-.42.17c-.11.12-.17.25-.17.41m0-2.02c0 .17.06.3.17.39c.12.11.26.16.42.16h13.81c.49 0 .92-.18 1.27-.52c.35-.35.52-.77.52-1.27c0-.49-.17-.91-.52-1.26s-.77-.52-1.27-.52c-.49 0-.91.17-1.27.51c-.11.12-.16.27-.16.42c0 .16.05.3.16.4c.11.1.24.15.4.15c.15 0 .29-.05.41-.16a.61.61 0 0 1 .45-.18c.17 0 .33.06.46.18s.2.27.2.45s-.07.34-.2.47c-.13.13-.28.2-.46.2H3.66c-.16 0-.3.06-.42.17c-.11.11-.17.25-.17.41M15.96 21.9h2.47l1.65-7.99h-2.47zm3.55 0h6.62l.4-1.9h-3.67v-.02c.2-.09.49-.22.86-.37c.38-.15.69-.28.95-.38s.54-.25.86-.44c.32-.19.58-.38.77-.58s.36-.45.5-.75s.21-.64.21-1c0-.56-.14-1.02-.43-1.4c-.29-.38-.65-.64-1.08-.8c-.43-.16-.92-.23-1.45-.23c-.97 0-1.76.26-2.37.78c-.61.52-.98 1.29-1.1 2.31h2.07c0-.38.11-.69.33-.95c.22-.26.53-.38.93-.38c.3 0 .52.08.67.24c.15.16.22.34.22.55c0 .32-.11.58-.33.76c-.22.18-.63.42-1.25.72c-.04.01-.07.02-.08.04c-.89.43-1.44.7-1.65.83c-.79.47-1.34 1.06-1.65 1.74c-.16.36-.27.77-.33 1.23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindBeaufortTwo(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.94 13.5c0 .18.06.31.17.4c.12.11.26.17.42.17h10.4c.17 0 .33.06.46.19c.13.12.2.28.2.46s-.07.34-.2.47a.629.629 0 0 1-.93-.01a.584.584 0 0 0-.4-.16c-.16 0-.3.05-.41.16c-.11.11-.16.24-.16.39c0 .16.06.3.17.41c.36.36.78.53 1.27.53s.91-.17 1.26-.52s.52-.77.52-1.27c0-.49-.17-.92-.52-1.27s-.77-.53-1.26-.53H5.53c-.16 0-.3.06-.42.17c-.12.12-.17.25-.17.41m0-2.02c0 .17.06.3.17.39c.12.11.26.16.42.16h13.81c.49 0 .92-.18 1.27-.52c.35-.35.52-.77.52-1.27c0-.49-.17-.91-.52-1.26s-.77-.52-1.27-.52c-.49 0-.91.17-1.27.51c-.11.12-.16.27-.16.42c0 .16.05.3.16.4c.11.1.24.15.4.15c.15 0 .29-.05.41-.16c.12-.12.27-.18.46-.18c.17 0 .33.06.46.18s.2.27.2.45s-.07.34-.2.47c-.13.13-.28.2-.46.2H5.53c-.16 0-.3.06-.42.17a.53.53 0 0 0-.17.41m12.72 10.37h6.62l.4-1.89H21v-.03c.2-.09.49-.22.86-.37c.38-.15.69-.28.95-.38s.55-.25.87-.44s.57-.38.77-.57c.19-.19.36-.44.5-.75s.21-.64.21-1c0-.56-.14-1.02-.43-1.4s-.65-.65-1.08-.81a4.16 4.16 0 0 0-1.45-.24c-.97 0-1.76.26-2.38.78c-.62.52-.98 1.29-1.1 2.31h2.09c0-.37.11-.68.32-.94c.22-.26.52-.38.91-.38c.3 0 .52.08.67.24s.23.34.23.54c0 .12-.01.23-.03.32s-.07.19-.15.28s-.15.16-.21.22s-.17.13-.34.23c-.17.09-.3.17-.4.22a39.469 39.469 0 0 0-2.17 1.08c-.8.48-1.35 1.07-1.66 1.78c-.16.36-.27.76-.32 1.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindBeaufortZero(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.01 13.5c0 .18.06.31.17.4c.12.11.26.17.42.17H16c.17 0 .33.06.46.19c.13.12.2.28.2.46s-.07.34-.2.47a.629.629 0 0 1-.93-.01a.603.603 0 0 0-.4-.16c-.16 0-.3.05-.41.16c-.11.11-.16.24-.16.39c0 .16.06.3.17.41c.36.36.78.53 1.28.53s.91-.17 1.26-.52s.52-.77.52-1.27c0-.49-.17-.92-.52-1.27s-.77-.53-1.26-.53H5.6c-.16 0-.3.06-.42.17c-.11.12-.17.25-.17.41m0-2.02c0 .17.06.3.17.39c.12.11.26.16.42.16h13.81c.49 0 .92-.18 1.27-.52c.35-.35.52-.77.52-1.27c0-.49-.17-.91-.52-1.26s-.77-.52-1.27-.52c-.49 0-.91.17-1.27.51c-.11.12-.16.27-.16.42c0 .16.05.3.16.4c.11.1.24.15.4.15c.15 0 .29-.05.41-.16a.61.61 0 0 1 .45-.18c.17 0 .33.06.46.18s.2.27.2.45s-.07.34-.2.47c-.13.13-.28.2-.46.2H5.6c-.16 0-.3.06-.42.17c-.11.11-.17.25-.17.41m13.26 7.42c0 .52.08.98.24 1.37s.38.71.66.94c.28.23.58.4.91.52c.33.11.68.17 1.05.17c.51 0 .98-.09 1.41-.26c.43-.17.77-.4 1.05-.69c.27-.29.51-.61.71-.95c.2-.34.35-.7.45-1.08s.18-.72.23-1.03s.07-.6.07-.86c0-.97-.27-1.72-.8-2.25s-1.24-.8-2.12-.8c-.49 0-.97.12-1.43.35s-.87.56-1.23.98c-.36.42-.65.94-.86 1.56c-.23.62-.34 1.3-.34 2.03m2.05.06c0-.15.01-.34.04-.58c.03-.23.08-.51.16-.83c.08-.32.18-.62.3-.89s.29-.5.52-.69c.22-.19.47-.29.75-.29c.27 0 .49.09.65.26c.16.17.23.44.23.79c0 .96-.17 1.78-.5 2.45s-.74 1.01-1.23 1.01c-.61 0-.92-.41-.92-1.23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindDeg(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.74 14.5c0-2.04.51-3.93 1.52-5.66s2.38-3.1 4.11-4.11s3.61-1.51 5.64-1.51c1.52 0 2.98.3 4.37.89s2.58 1.4 3.59 2.4s1.81 2.2 2.4 3.6s.89 2.85.89 4.39c0 1.52-.3 2.98-.89 4.37s-1.4 2.59-2.4 3.59s-2.2 1.8-3.59 2.39s-2.84.89-4.37.89c-1.53 0-3-.3-4.39-.89s-2.59-1.4-3.6-2.4s-1.8-2.2-2.4-3.58s-.88-2.84-.88-4.37m2.48 0c0 2.37.86 4.43 2.59 6.18c1.73 1.73 3.79 2.59 6.2 2.59c1.58 0 3.05-.39 4.39-1.18s2.42-1.85 3.21-3.2s1.19-2.81 1.19-4.39s-.4-3.05-1.19-4.4s-1.86-2.42-3.21-3.21s-2.81-1.18-4.39-1.18s-3.05.39-4.39 1.18S8.2 8.75 7.4 10.1s-1.18 2.82-1.18 4.4m4.89 5.85l3.75-13.11c.01-.1.06-.15.15-.15s.14.05.15.15l3.74 13.11c.04.11.03.19-.02.25s-.13.06-.24 0l-3.47-1.3c-.1-.04-.2-.04-.29 0l-3.5 1.3c-.1.06-.17.06-.21 0s-.08-.15-.06-.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindDirection(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M0 770q0-209 103.5-386.5T384 103T769 0q156 0 298 61t245 164t164 245.5t61 299.5q0 156-61 298.5t-164 245t-245 163t-298 60.5q-157 0-299.5-61T224 1312T60.5 1067.5T0 770zm169 0q0 243 177 422q177 177 423 177q162 0 300-80.5t219-218.5t81-300t-81-300.5t-219-219T769 170t-300 80.5t-219 219T169 770zm334 399l256-895q1-10 10-10t10 10l255 895q4 11-1.5 17t-16.5 0l-237-89q-10-4-20 0l-239 89q-10 6-14.5 0t-2.5-17z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindDirectionE(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M0 767q0-156 60.5-298.5t163-244.5t244-163T765 0q157 0 299 61t245 163.5T1473 469t61 297q0 156-61 298.5T1309 1310t-245 163.5t-298 60.5q-155 0-297-60.5t-244.5-163T61 1066T0 767zm168 0q0 162 80.5 299.5t218.5 217t299 79.5q121 0 232-47.5t191.5-127.5t128-190.5T1365 766q0-244-178-422q-178-176-422-176q-121 0-231.5 48T343 344T215.5 535T168 767zm323 248l90-239q5-10 0-21l-90-235q-5-10 1-16t17-1l598 252q11 2 11 11q0 8-11 10l-598 257q-11 4-17-1.5t-1-16.5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindDirectionN(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M0 770q0-209 103.5-386.5T384 103T769 0q156 0 298 61t245 164t164 245.5t61 299.5q0 156-61 298.5t-164 245t-245 163t-298 60.5q-157 0-299.5-61T224 1312T60.5 1067.5T0 770zm169 0q0 243 177 422q177 177 423 177q162 0 300-80.5t219-218.5t81-300t-81-300.5t-219-219T769 170t-300 80.5t-219 219T169 770zm334 258l256-600q2-11 10-11t10 11l255 600q4 11-1.5 17t-16.5 0l-237-89q-10-4-20 0l-239 89q-10 6-16 0t-1-17z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindDirectionNe(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M0 767q0-156 60.5-298.5t163-244.5t244-163T765 0q157 0 299 61t245 163.5T1473 469t61 297q0 156-61 298.5T1309 1310t-245 163.5t-298 60.5q-155 0-297-60.5t-244.5-163T61 1066T0 767zm168 0q0 162 80.5 299.5t218.5 217t299 79.5q121 0 232-47.5t191.5-127.5t128-190.5T1365 766q0-244-178-422q-178-176-422-176q-121 0-231.5 48T343 344T215.5 535T168 767zm267 8q0-7 12-10l604-250q11-2 16.5 3.5t1.5 16.5l-248 604q-2 11-11 11q-11 0-16-11L691 909q-9-14-14-15L447 790q-12-5-12-15z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindDirectionNw(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M0 767q0-156 60.5-298.5t163-244.5t244-163T765 0q157 0 299 61t245 163.5T1473 469t61 297q0 156-61 298.5T1309 1310t-245 163.5t-298 60.5q-155 0-297-60.5t-244.5-163T61 1066T0 767zm168 0q0 162 80.5 299.5t218.5 217t299 79.5q121 0 232-47.5t191.5-127.5t128-190.5T1365 766q0-244-178-422q-178-176-422-176q-121 0-231.5 48T343 344T215.5 535T168 767zm350-229q-2-11 3.5-17t15.5-2l602 249q12 2 12 10q0 9-12 14L910 896q-10 4-15 15l-104 230q-5 11-14 11q-8 0-10-11z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindDirectionS(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M0 768q0-156 61.5-298.5T226 224T471 60.5T769 0t298.5 60.5T1313 224t164 245t61 299t-61 299t-164 244.5t-245.5 163T769 1535t-298.5-61T225 1310T61 1065T0 768zm170 0q0 245 177 422q176 176 422 176q163 0 301.5-80.5t219-218T1370 768q0-121-47.5-232T1194 344.5t-192-128T769 169q-121 0-231.5 47.5t-191 128T218 536t-48 232zm334-259q-5-11 1-16.5t16-.5l238 89q10 4 23 0l235-89q10-5 16 .5t2 16.5l-253 599q-3 10-13 10q-7 0-10-10z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindDirectionSe(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M0 767q0-156 60.5-298.5t163-244.5t244-163T765 0q157 0 299 61t245 163.5T1473 469t61 297q0 156-61 298.5T1309 1310t-245 163.5t-298 60.5q-155 0-297-60.5t-244.5-163T61 1066T0 767zm168 0q0 162 80.5 299.5t218.5 217t299 79.5q121 0 232-47.5t191.5-127.5t128-190.5T1365 766q0-244-178-422q-178-176-422-176q-121 0-231.5 48T343 344T215.5 535T168 767zm207-8q0-10 10-14l231-104q10-5 15-15l103-229q9-12 16-12q8 0 11 12l248 605q2 11-2 15t-16 2L385 771q-10-3-10-12z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindDirectionSw(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M0 767q0-156 60.5-298.5t163-244.5t244-163T765 0q157 0 299 61t245 163.5T1473 469t61 297q0 156-61 298.5T1309 1310t-245 163.5t-298 60.5q-155 0-297-60.5t-244.5-163T61 1066T0 767zm168 0q0 162 80.5 299.5t218.5 217t299 79.5q121 0 232-47.5t191.5-127.5t128-190.5T1365 766q0-244-178-422q-178-176-422-176q-121 0-231.5 48T343 344T215.5 535T168 767zm357 233l250-605q1-10 10-10q10 0 14 10l104 231q2 9 16 14l230 104q11 4 11 14q0 9-11 12l-605 248q-10 4-15-1t-4-17z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindDirectionW(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M0 772q0-157 61.5-300t165-246.5T472.5 61T771 0t299 61t246 164.5T1480.5 472t61.5 300t-61.5 300t-165 245.5t-246 163.5t-298.5 61q-157 0-300-61t-246-164.5t-164-246T0 772zm170 0q0 245 179 424q177 177 422 177q163 0 302-81t220-219.5t81-300.5t-81-300.5T1073 252t-302-81q-162 0-300.5 81T251 471.5T170 772zm248 0q0-9 12-10l601-258q11-5 16.5 1.5t.5 16.5l-89 240q-7 10 0 20l89 238q5 10-.5 16t-16.5 1L430 782q-12-1-12-10z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Windy(children ...ElementRenderer) *WiIcon {
	return &WiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.65 15.5c0-.22.08-.41.23-.56a.8.8 0 0 1 .57-.22h12.08c.22 0 .4.07.54.22s.22.34.22.57c0 .22-.07.4-.22.54c-.14.14-.32.22-.54.22H5.45c-.22 0-.42-.07-.57-.22a.711.711 0 0 1-.23-.55m2.41-2.9c0-.22.08-.4.23-.55a.77.77 0 0 1 .56-.23h12.09c.21 0 .39.08.54.23c.15.15.22.33.22.55c0 .22-.07.4-.22.56c-.15.15-.33.23-.54.23H7.86c-.22 0-.41-.08-.56-.23s-.24-.34-.24-.56m1.62 5.74c0-.21.08-.39.24-.54c.14-.14.32-.22.54-.22h12.1c.22 0 .41.07.56.22c.15.14.22.32.22.54s-.08.41-.23.56s-.34.23-.56.23H9.46c-.22 0-.4-.08-.56-.23s-.22-.34-.22-.56m10.58-2.84c0-.23.07-.42.22-.57c.15-.15.34-.22.57-.22h4.52c.23 0 .42.07.57.22c.15.15.22.34.22.56c0 .22-.07.4-.22.54s-.34.22-.56.22h-4.52c-.23 0-.42-.07-.57-.22c-.16-.13-.23-.31-.23-.53"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
