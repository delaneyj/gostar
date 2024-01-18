package covid

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = ""
	hAttr          = "1em"
	viewbox        = "0 0 24 24"
	fill           = "currentColor"
)

type CovidIcon struct {
	*SVGSVGElement
}

type CovidIconFn func(children ...ElementRenderer) *CovidIcon

var IconLookup = map[string]CovidIconFn{
	"covidCarrierBloodOne":                         CovidCarrierBloodOne,
	"covidCarrierBloodTwo":                         CovidCarrierBloodTwo,
	"covidCarrierHuman":                            CovidCarrierHuman,
	"covidCarrierPackages":                         CovidCarrierPackages,
	"covidNineteenVirusBat":                        CovidNineteenVirusBat,
	"covidNineteenVirusFour":                       CovidNineteenVirusFour,
	"covidNineteenVirusHealOne":                    CovidNineteenVirusHealOne,
	"covidNineteenVirusHealTwo":                    CovidNineteenVirusHealTwo,
	"covidNineteenVirusLifelongOne":                CovidNineteenVirusLifelongOne,
	"covidNineteenVirusLifelongTwo":                CovidNineteenVirusLifelongTwo,
	"covidNineteenVirusOne":                        CovidNineteenVirusOne,
	"covidNineteenVirusPandemicOne":                CovidNineteenVirusPandemicOne,
	"covidNineteenVirusPandemicThree":              CovidNineteenVirusPandemicThree,
	"covidNineteenVirusPandemicTwo":                CovidNineteenVirusPandemicTwo,
	"covidNineteenVirusPatientOne":                 CovidNineteenVirusPatientOne,
	"covidNineteenVirusPatientTwo":                 CovidNineteenVirusPatientTwo,
	"covidNineteenVirusReinfected":                 CovidNineteenVirusReinfected,
	"covidNineteenVirusThree":                      CovidNineteenVirusThree,
	"covidNineteenVirusTwo":                        CovidNineteenVirusTwo,
	"covidNineteenVirusWarningOne":                 CovidNineteenVirusWarningOne,
	"covidNineteenVirusWarningThree":               CovidNineteenVirusWarningThree,
	"covidNineteenVirusWarningTwo":                 CovidNineteenVirusWarningTwo,
	"graphCuredDecreasing":                         GraphCuredDecreasing,
	"graphCuredIncreasing":                         GraphCuredIncreasing,
	"graphCuredStable":                             GraphCuredStable,
	"graphDeathRateDecreasing":                     GraphDeathRateDecreasing,
	"graphDeathRateIncreasing":                     GraphDeathRateIncreasing,
	"graphDeathRateStable":                         GraphDeathRateStable,
	"graphDocumentInfectedReport":                  GraphDocumentInfectedReport,
	"graphInfectedDecreasing":                      GraphInfectedDecreasing,
	"graphInfectedIncreasing":                      GraphInfectedIncreasing,
	"graphInfectedStable":                          GraphInfectedStable,
	"mutationOne":                                  MutationOne,
	"mutationStronger":                             MutationStronger,
	"mutationTemperatureChange":                    MutationTemperatureChange,
	"mutationTwo":                                  MutationTwo,
	"personalHygieneCleanBottleShield":             PersonalHygieneCleanBottleShield,
	"personalHygieneCleanBottleVirus":              PersonalHygieneCleanBottleVirus,
	"personalHygieneCleanBottleVirusBlock":         PersonalHygieneCleanBottleVirusBlock,
	"personalHygieneCleanGel":                      PersonalHygieneCleanGel,
	"personalHygieneCleanToothpaste":               PersonalHygieneCleanToothpaste,
	"personalHygieneHandLiquidSoap":                PersonalHygieneHandLiquidSoap,
	"personalHygieneHandSanitizerLiquidDrop":       PersonalHygieneHandSanitizerLiquidDrop,
	"personalHygieneHandSanitizerLiquidOne":        PersonalHygieneHandSanitizerLiquidOne,
	"personalHygieneHandSanitizerLiquidThree":      PersonalHygieneHandSanitizerLiquidThree,
	"personalHygieneHandSanitizerLiquidTwo":        PersonalHygieneHandSanitizerLiquidTwo,
	"personalHygieneHandSanitizerLiquidVirusBlock": PersonalHygieneHandSanitizerLiquidVirusBlock,
	"personalHygieneHandSanitizerSpray":            PersonalHygieneHandSanitizerSpray,
	"personalHygieneHandSanitizerSprayVirusBlock":  PersonalHygieneHandSanitizerSprayVirusBlock,
	"personalHygieneHandSanitizerVirusBlock":       PersonalHygieneHandSanitizerVirusBlock,
	"personalHygieneHandSanitizerVirusShield":      PersonalHygieneHandSanitizerVirusShield,
	"personalHygieneHandSoapOne":                   PersonalHygieneHandSoapOne,
	"personalHygieneHandSoapTwo":                   PersonalHygieneHandSoapTwo,
	"personalHygieneHandWash":                      PersonalHygieneHandWash,
	"personalHygieneHandWipePaperFour":             PersonalHygieneHandWipePaperFour,
	"personalHygieneHandWipePaperOne":              PersonalHygieneHandWipePaperOne,
	"personalHygieneHandWipePaperThree":            PersonalHygieneHandWipePaperThree,
	"personalHygieneHandWipePaperTwo":              PersonalHygieneHandWipePaperTwo,
	"quarantinePlaceBed":                           QuarantinePlaceBed,
	"quarantinePlaceHospital":                      QuarantinePlaceHospital,
	"quarantinePlaceHouseOne":                      QuarantinePlaceHouseOne,
	"quarantinePlaceHouseShield":                   QuarantinePlaceHouseShield,
	"quarantinePlaceHouseTwo":                      QuarantinePlaceHouseTwo,
	"quarantinePlaceSelfLockdownOne":               QuarantinePlaceSelfLockdownOne,
	"quarantinePlaceSelfLockdownTwo":               QuarantinePlaceSelfLockdownTwo,
	"quarantinePlaceTimeCalendarDay":               QuarantinePlaceTimeCalendarDay,
	"quarantinePlaceTimeCalendarOne":               QuarantinePlaceTimeCalendarOne,
	"quarantinePlaceTimeCalendarTwo":               QuarantinePlaceTimeCalendarTwo,
	"quarantinePlaceTimeDateDay":                   QuarantinePlaceTimeDateDay,
	"socialDistancingAttention":                    SocialDistancingAttention,
	"socialDistancingCorrectFive":                  SocialDistancingCorrectFive,
	"socialDistancingCorrectFour":                  SocialDistancingCorrectFour,
	"socialDistancingCorrectOne":                   SocialDistancingCorrectOne,
	"socialDistancingCorrectSix":                   SocialDistancingCorrectSix,
	"socialDistancingCorrectThree":                 SocialDistancingCorrectThree,
	"socialDistancingCorrectTwo":                   SocialDistancingCorrectTwo,
	"socialDistancingDoNotCloseFour":               SocialDistancingDoNotCloseFour,
	"socialDistancingDoNotCloseOne":                SocialDistancingDoNotCloseOne,
	"socialDistancingDoNotCloseThree":              SocialDistancingDoNotCloseThree,
	"socialDistancingDoNotCloseTwo":                SocialDistancingDoNotCloseTwo,
	"socialDistancingDoNotTouch":                   SocialDistancingDoNotTouch,
	"socialDistancingMan":                          SocialDistancingMan,
	"socialDistancingNotAllowedSpaceMan":           SocialDistancingNotAllowedSpaceMan,
	"socialDistancingNotAllowedSpaceWoman":         SocialDistancingNotAllowedSpaceWoman,
	"socialDistancingOne":                          SocialDistancingOne,
	"socialDistancingProtectShieldOne":             SocialDistancingProtectShieldOne,
	"socialDistancingProtectShieldTwo":             SocialDistancingProtectShieldTwo,
	"socialDistancingTwo":                          SocialDistancingTwo,
	"socialDistancingVirus":                        SocialDistancingVirus,
	"socialDistancingWoman":                        SocialDistancingWoman,
	"symptomsColdFever":                            SymptomsColdFever,
	"symptomsFever":                                SymptomsFever,
	"symptomsNausea":                               SymptomsNausea,
	"symptomsVirusDiarrheaOne":                     SymptomsVirusDiarrheaOne,
	"symptomsVirusDiarrheaTwo":                     SymptomsVirusDiarrheaTwo,
	"symptomsVirusHeadacheOne":                     SymptomsVirusHeadacheOne,
	"symptomsVirusHeadacheTwo":                     SymptomsVirusHeadacheTwo,
	"symptomsVirusLossSmellOne":                    SymptomsVirusLossSmellOne,
	"symptomsVirusLossSmellTwo":                    SymptomsVirusLossSmellTwo,
	"symptomsVirusLungDamage":                      SymptomsVirusLungDamage,
	"symptomsVirusStomach":                         SymptomsVirusStomach,
	"transmissionVirusAirplaneFlight":              TransmissionVirusAirplaneFlight,
	"transmissionVirusBriefcase":                   TransmissionVirusBriefcase,
	"transmissionVirusCough":                       TransmissionVirusCough,
	"transmissionVirusCrowd":                       TransmissionVirusCrowd,
	"transmissionVirusExpand":                      TransmissionVirusExpand,
	"transmissionVirusFloatWind":                   TransmissionVirusFloatWind,
	"transmissionVirusHandshake":                   TransmissionVirusHandshake,
	"transmissionVirusHumanInfected":               TransmissionVirusHumanInfected,
	"transmissionVirusHumanTransmitOne":            TransmissionVirusHumanTransmitOne,
	"transmissionVirusHumanTransmitTwo":            TransmissionVirusHumanTransmitTwo,
	"transmissionVirusIncreaseRate":                TransmissionVirusIncreaseRate,
	"transmissionVirusInhale":                      TransmissionVirusInhale,
	"transmissionVirusMobileApplication":           TransmissionVirusMobileApplication,
	"transmissionVirusRatMouse":                    TransmissionVirusRatMouse,
	"transmissionVirusTouchFinger":                 TransmissionVirusTouchFinger,
	"transmissionVirusTouchHandOne":                TransmissionVirusTouchHandOne,
	"transmissionVirusTouchHandTwo":                TransmissionVirusTouchHandTwo,
	"transmissionVirusTransmit":                    TransmissionVirusTransmit,
	"transmissionVirusTransportation":              TransmissionVirusTransportation,
	"transmissionVirusTrashBin":                    TransmissionVirusTrashBin,
	"transmissionVirusVisible":                     TransmissionVirusVisible,
	"transmissionVirusWindBreath":                  TransmissionVirusWindBreath,
	"vaccineProtectionFaceMaskOne":                 VaccineProtectionFaceMaskOne,
	"vaccineProtectionFaceMaskThree":               VaccineProtectionFaceMaskThree,
	"vaccineProtectionFaceMaskTwo":                 VaccineProtectionFaceMaskTwo,
	"vaccineProtectionFaceShieldOne":               VaccineProtectionFaceShieldOne,
	"vaccineProtectionFaceShieldTwo":               VaccineProtectionFaceShieldTwo,
	"vaccineProtectionInfraredThermometerGun":      VaccineProtectionInfraredThermometerGun,
	"vaccineProtectionMedicinePill":                VaccineProtectionMedicinePill,
	"vaccineProtectionPeopleShield":                VaccineProtectionPeopleShield,
	"vaccineProtectionSanitizerSpray":              VaccineProtectionSanitizerSpray,
	"vaccineProtectionShield":                      VaccineProtectionShield,
	"vaccineProtectionSyringe":                     VaccineProtectionSyringe,
	"vaccineProtectionVirus":                       VaccineProtectionVirus,
	"vaccineProtectionWashHands":                   VaccineProtectionWashHands,
	"virusLabResearchMagnifierOne":                 VirusLabResearchMagnifierOne,
	"virusLabResearchMagnifierTwo":                 VirusLabResearchMagnifierTwo,
	"virusLabResearchMedicinePill":                 VirusLabResearchMedicinePill,
	"virusLabResearchMicroscope":                   VirusLabResearchMicroscope,
	"virusLabResearchSyringe":                      VirusLabResearchSyringe,
	"virusLabResearchTestTube":                     VirusLabResearchTestTube,
}

func CovidCarrierBloodOne(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.182 15.067A8.019 8.019 0 0 1 12 23.25a8.02 8.02 0 0 1-8.182-8.182C3.818 8.144 12 .75 12 .75s8.182 7.395 8.182 14.318"/><path d="M12 17.66a2.795 2.795 0 1 0 0-5.59a2.795 2.795 0 0 0 0 5.59m-.466-7.686h.932m-.466 0v2.096m3.129-.993l.659.659m-.329-.33l-1.483 1.483m2.915 1.51v.932m0-.466h-2.096m.993 3.129l-.659.659m.33-.329l-1.483-1.483m-1.51 2.915h-.932m.466 0V17.66m-3.129.993l-.659-.659m.329.33l1.483-1.483m-2.915-1.51v-.932m0 .466h2.096m-.993-3.129l.659-.659m-.33.329l1.483 1.483"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CovidCarrierBloodTwo(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.881 10.851a3.517 3.517 0 1 0 0-7.034a3.517 3.517 0 0 0 0 7.034m-.586-9.672h1.172m-.586 0v2.638m3.937-1.25l.829.829m-.414-.414l-1.865 1.865m3.668 1.901V7.92m0-.586h-2.638m1.249 3.937l-.829.829m.415-.414l-1.865-1.865m-1.901 3.667h-1.172m.586 0v-2.637M12.943 12.1l-.829-.829m.415.415l1.865-1.865M10.726 7.92V6.748m0 .586h2.638m-1.25-3.938l.829-.829m-.414.415l1.865 1.865m.442 11.467a6.936 6.936 0 1 1-13.872 0c0-5 4.571-12.548 6.292-15.211a.765.765 0 0 1 1.289 0l.761 1.288"/><path d="M7.9 20.167a3.853 3.853 0 0 1-3.853-3.853"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CovidCarrierHuman(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M11.857 20.628a3.163 3.163 0 1 0 0-6.326a3.163 3.163 0 0 0 0 6.326m-.527-8.698h1.054m-.527 0v2.372m3.541-1.124l.746.746m-.373-.373l-1.677 1.677m3.298 1.71v1.054m0-.527H15.02m1.124 3.541l-.746.746m.373-.373l-1.677-1.678M12.384 23H11.33m.527 0v-2.372m-3.541 1.124l-.745-.746m.372.373l1.678-1.678m-3.299-1.709v-1.054m0 .527h2.372m-1.123-3.541l.745-.746m-.373.373l1.678 1.677"/><path d="M9.357 9.004a4.529 4.529 0 1 1 5.286 0a8.41 8.41 0 0 1 5.769 7.98v6.219"/><path d="M3.588 23.203v-6.218a8.41 8.41 0 0 1 5.769-7.981"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CovidCarrierPackages(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 17.357a2.857 2.857 0 1 0 0-5.714a2.857 2.857 0 0 0 0 5.714M11.524 9.5h.952M12 9.5v2.143m3.199-1.015l.673.673m-.336-.337L14.02 12.48M17 14.024v.952m0-.476h-2.143m1.015 3.199l-.673.673m.337-.336L14.02 16.52m-1.544 2.98h-.952m.476 0v-2.143m-3.199 1.015l-.673-.673m.336.337L9.98 16.52M7 14.976v-.952m0 .476h2.143m-1.015-3.199l.673-.673m-.337.336L9.98 12.48"/><path d="M21.167 6.5H2.833C1.821 6.5 1 7.32 1 8.333v12.834C1 22.179 1.82 23 2.833 23h18.334C22.179 23 23 22.18 23 21.167V8.333c0-1.012-.82-1.833-1.833-1.833"/><path d="m22.408 6.985l-3.447-5.169A1.834 1.834 0 0 0 17.435 1H6.565a1.834 1.834 0 0 0-1.526.816L1.592 6.985M12 1v5.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CovidNineteenVirusBat(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.211 20.568a3.429 3.429 0 1 0 0-6.858a3.429 3.429 0 0 0 0 6.858m-.571-9.429h1.143m-.572 0v2.571m3.839-1.218l.808.808m-.404-.404l-1.819 1.819m3.576 1.853v1.142m0-.571H20.64m1.218 3.839l-.808.808m.404-.404l-1.819-1.819m-1.852 3.576H16.64m.571 0v-2.571m-3.838 1.218l-.809-.808m.404.404l1.819-1.819m-3.576-1.853v-1.142m0 .571h2.572M12.564 13.3l.809-.808m-.405.404l1.819 1.819M8.813.861v4.012L5.665 3.824A3.7 3.7 0 0 0 .789 7.339v2.719s2.485-1.043 4.076 2.619c2.038-2.038 4.075 0 4.075 2.038M14.163.861v4.012l3.147-1.049a3.705 3.705 0 0 1 4.877 3.515"/><path d="M8.813 3.536A4.667 4.667 0 0 1 11.488 2.2c1.014.14 1.954.61 2.675 1.338"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CovidNineteenVirusFour(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.25 19.651a3.714 3.714 0 1 0 0-7.428a3.714 3.714 0 0 0 0 7.428M6.631 9.437h1.238m-.619 0v2.786m4.158-1.32l.876.875m-.438-.437l-1.97 1.969m3.874 2.008v1.238m0-.619h-2.786m1.32 4.158l-.876.876m.438-.438l-1.97-1.97m-2.007 3.874H6.631m.619 0v-2.786m-4.158 1.32l-.876-.876m.438.438l1.97-1.97M.75 16.556v-1.238m0 .619h2.786m-1.32-4.159l.876-.875m-.438.438l1.97 1.969M18.5 9.027a2.714 2.714 0 1 0 0-5.428a2.714 2.714 0 0 0 0 5.428m-.452-7.464h.904m-.452 0v2.036m3.039-.964l.64.639m-.32-.32l-1.44 1.44m2.831 1.467v.905m0-.453h-2.036m.965 3.039l-.64.64m.32-.32l-1.44-1.44m-1.467 2.831h-.904m.452 0V9.027m-3.039.965l-.64-.64m.32.32l1.44-1.44M13.75 6.766v-.905m0 .452h2.036m-.965-3.039l.64-.639m-.32.319l1.44 1.44"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CovidNineteenVirusHealOne(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M23.25 17.25a1 1 0 0 0-1-1h-2.5v-2.5a1 1 0 0 0-1-1h-1.5a1 1 0 0 0-1 1v2.5h-2.5a1 1 0 0 0-1 1v1.5a1 1 0 0 0 1 1h2.5v2.5a1 1 0 0 0 1 1h1.5a1 1 0 0 0 1-1v-2.5h2.5a1 1 0 0 0 1-1zM4.183 12.754H2.162a1.413 1.413 0 0 1 0-2.825h2.021m5.746-5.746V2.162a1.413 1.413 0 0 1 2.825 0v2.021"/><path d="M15.405 5.28a7.3 7.3 0 0 0-8.128 0M18.5 9.929a7.25 7.25 0 0 0-1.1-2.652M7.277 17.4a7.271 7.271 0 0 0 2.478 1.061M5.28 7.277a7.3 7.3 0 0 0 0 8.128"/><path d="M6.848 9.511c0-.426-.17-.835-.471-1.137L4.518 6.515a1.414 1.414 0 0 1 2-2l1.856 1.862c.302.301.71.47 1.137.471m0 8.987a1.606 1.606 0 0 0-1.137.471l-1.859 1.859a1.414 1.414 0 0 1-2-2l1.859-1.86c.301-.301.47-.71.471-1.136m6.327-6.321c.426 0 .835-.17 1.136-.471l1.86-1.859a1.414 1.414 0 1 1 2 2l-1.862 1.856a1.606 1.606 0 0 0-.471 1.137"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CovidNineteenVirusHealTwo(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12.017.75H8.983m1.517 0v2.6m-5.822-.817L3.606 3.606L2.533 4.678m1.073-1.072l1.838 1.838M.75 8.983v3.034m0-1.517h2.6m-.817 5.822l1.073 1.072l1.072 1.073m-1.072-1.073l1.838-1.838M18.467 4.678l-1.073-1.072l-1.072-1.073m1.072 1.073l-1.838 1.838M8.983 9.85a1.517 1.517 0 1 0 0-3.034a1.517 1.517 0 0 0 0 3.034"/><path stroke-linecap="round" stroke-linejoin="round" d="M17.61 9.75a7.15 7.15 0 1 0-7.982 7.842m13.622-.342a1 1 0 0 0-1-1h-2.5v-2.5a1 1 0 0 0-1-1h-1.5a1 1 0 0 0-1 1v2.5h-2.5a1 1 0 0 0-1 1v1.5a1 1 0 0 0 1 1h2.5v2.5a1 1 0 0 0 1 1h1.5a1 1 0 0 0 1-1v-2.5h2.5a1 1 0 0 0 1-1z"/><path d="M10.208 14.434a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CovidNineteenVirusLifelongOne(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M11.167 15.286a3.286 3.286 0 1 0 0-6.572a3.286 3.286 0 0 0 0 6.572m-.548-9.036h1.095m-.547 0v2.464m3.678-1.167l.775.774m-.387-.387L13.49 9.677m3.427 1.775v1.096m0-.548h-2.465m1.168 3.679l-.775.774m.388-.387l-1.743-1.743m-1.776 3.427h-1.095m.548 0v-2.464m-3.679 1.167l-.774-.774m.387.387l1.742-1.743m-3.426-1.775v-1.096m0 .548h2.464M6.714 8.321l.774-.774m-.387.387l1.742 1.743"/><path d="M19.673 17.598a10.267 10.267 0 1 1 1.66-5.6v.934"/><path d="m19.467 11.066l1.867 1.866l1.866-1.866"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CovidNineteenVirusLifelongTwo(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12.017.767H8.983m1.517 0v2.6M4.678 2.55L3.606 3.622L2.533 4.695m1.073-1.073l1.838 1.839M.75 9v3.033m0-1.516h2.6m-.817 5.821l1.073 1.073l1.072 1.072m-1.072-1.072l1.838-1.839M18.467 4.695l-1.073-1.073l-1.072-1.072m1.072 1.072l-1.838 1.839M8.983 9.867a1.517 1.517 0 1 0 0-3.034a1.517 1.517 0 0 0 0 3.034"/><path stroke-linecap="round" stroke-linejoin="round" d="M17.356 8.483a7.15 7.15 0 1 0-8.889 8.884"/><path stroke-linecap="round" stroke-linejoin="round" d="M17.25 23.233a6 6 0 1 0 0-12a6 6 0 0 0 0 12"/><path stroke-linecap="round" stroke-linejoin="round" d="M19.902 17.233H17.25v-2.651"/><path d="M9.417 13.367a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CovidNineteenVirusOne(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 20.25a8.25 8.25 0 1 0 0-16.5a8.25 8.25 0 0 0 0 16.5M13.75.75h-3.5m1.75 0v3m-6.718-.942L4.045 4.045L2.808 5.282m1.237-1.237l2.121 2.121M.75 10.25v3.5m0-1.75h3m-.942 6.718l1.237 1.237l1.237 1.237m-1.237-1.237l2.121-2.121m4.084 5.416h3.5m-1.75 0v-3m6.718.942l1.237-1.237l1.237-1.237m-1.237 1.237l-2.121-2.121m5.416-4.084v-3.5m0 1.75h-3m.942-6.718l-1.237-1.237l-1.237-1.237m1.237 1.237l-2.121 2.121"/><path stroke-linecap="round" stroke-linejoin="round" d="M9.5 11.25a1.75 1.75 0 1 0 0-3.5a1.75 1.75 0 0 0 0 3.5"/><path d="M12.75 17a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75m4-3.625a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CovidNineteenVirusPandemicOne(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.25 20.679a3.429 3.429 0 1 0 0-6.858a3.429 3.429 0 0 0 0 6.858m-.571-9.429h1.142m-.571 0v2.571m3.839-1.218l.808.808m-.404-.404l-1.819 1.819m3.576 1.853v1.142m0-.571h-2.571m1.218 3.839l-.808.808m.404-.404l-1.819-1.819m-1.853 3.576h-1.142m.571 0v-2.571m-3.839 1.218l-.808-.808m.404.404l1.819-1.819m-3.576-1.853v-1.142m0 .571h2.571m-1.218-3.839l.808-.808m-.404.404l1.819 1.819"/><path d="M9.639 22.194A10.808 10.808 0 1 1 22.272 10.14"/><path d="m4.823 20.01l.97-4.851H6.83a1.442 1.442 0 0 0 1.4-1.79l-.721-2.882a1.44 1.44 0 0 0-1.4-1.087H.967M20.2 5.073h-3.916a1.442 1.442 0 0 0-1.4 1.091l-.72 2.881"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CovidNineteenVirusPandemicThree(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.851 8.875a2.94 2.94 0 1 0 0-5.88a2.94 2.94 0 0 0 0 5.88M17.361.79h.98m-.49 0v2.205m3.292-1.045l.693.693m-.347-.346L19.93 3.856m3.066 1.589v.98m0-.49h-2.205m1.045 3.292l-.693.693m.346-.347L19.93 8.014m-1.589 3.066h-.98m.49 0V8.875M14.56 9.92l-.693-.693m.346.346l1.559-1.559m-3.066-1.589v-.98m0 .49h2.205m-1.044-3.292l.693-.693m-.347.347l1.559 1.559m4.515 9.806a9.57 9.57 0 0 1-.606 3.264m-1.809 3.024a9.637 9.637 0 0 1-2.588 2.083m-3.351 1.094c-1.106.143-2.227.1-3.319-.127m-3.253-1.358a9.812 9.812 0 0 1-2.418-2.274m-1.571-3.155a9.216 9.216 0 0 1-.341-3.3m.88-3.413a9.68 9.68 0 0 1 1.9-2.726m2.901-2.002a9.806 9.806 0 0 1 3.22-.808m6.499 9.814A5.789 5.789 0 1 1 10 7.822"/><path d="m8.772 8.1l.355 1.447A1.287 1.287 0 0 1 8.84 10.7l-.653.738a1.284 1.284 0 0 0-.307.65L7.671 13.4a1.285 1.285 0 0 1-1.094 1.071l-1.612.223m11.223.547l-2.073-.885a1.287 1.287 0 0 0-1.764.921l-.27 1.293c-.052.248-.029.507.066.742l.657 1.634"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CovidNineteenVirusPandemicTwo(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.885 6.648a5.352 5.352 0 0 1 0 10.7m-.892-14.713h1.784m-.892 0v4.013m5.991-1.901l1.262 1.261m-.631-.63l-2.838 2.838m5.581 2.892v1.784m0-.892h-4.014m1.902 5.992l-1.262 1.261m.631-.631l-2.838-2.838m-2.892 5.581h-1.784m.892 0v-4.013M10.25 21.5a9.5 9.5 0 0 1 0-19z"/><path d="M10.25 2.5A10.924 10.924 0 0 0 5.5 12a10.924 10.924 0 0 0 4.75 9.5m0-12.571H1.665M10 15.071H1.415"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CovidNineteenVirusPatientOne(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.25 20.679a3.429 3.429 0 1 0 0-6.858a3.429 3.429 0 0 0 0 6.858m-.571-9.429h1.142m-.571 0v2.571m3.839-1.218l.808.808m-.404-.404l-1.819 1.819m3.576 1.853v1.142m0-.571h-2.571m1.218 3.839l-.808.808m.404-.404l-1.819-1.819m-1.853 3.576h-1.142m.571 0v-2.571m-3.839 1.218l-.808-.808m.404.404l1.819-1.819m-3.576-1.853v-1.142m0 .571h2.571m-1.218-3.839l.808-.808m-.404.404l1.819 1.819M7.5 9a4.125 4.125 0 1 0 0-8.25A4.125 4.125 0 0 0 7.5 9M.75 17.25a6.753 6.753 0 0 1 9.4-6.208"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CovidNineteenVirusPatientTwo(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.25 20.679a3.429 3.429 0 1 0 0-6.858a3.429 3.429 0 0 0 0 6.858m-.571-9.429h1.142m-.571 0v2.571m3.839-1.218l.808.808m-.404-.404l-1.819 1.819m3.576 1.853v1.142m0-.571h-2.571m1.218 3.839l-.808.808m.404-.404l-1.819-1.819m-1.853 3.576h-1.142m.571 0v-2.571m-3.839 1.218l-.808-.808m.404.404l1.819-1.819m-3.576-1.853v-1.142m0 .571h2.571m-1.218-3.839l.808-.808m-.404.404l1.819 1.819M6 6.75a3 3 0 1 0 0-6a3 3 0 0 0 0 6m4.555 4.138A5.251 5.251 0 0 0 .75 13.5v2.25H3l.75 7.5h4.5l.323-3.232"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CovidNineteenVirusReinfected(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><path id="covidCovid19VirusReinfected0" d="m15.664 13.654l-2.914 2.914l2.914 2.914"/></defs><g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12.017.75H8.983m1.517 0v2.6m-5.822-.817L3.606 3.606L2.533 4.678m1.073-1.072l1.838 1.838M.75 8.983v3.034m0-1.517h2.6m-.817 5.822l1.073 1.072l1.072 1.073m-1.072-1.073l1.838-1.838M18.467 4.678l-1.073-1.072l-1.072-1.073m1.072 1.073l-1.838 1.838M8.983 9.85a1.517 1.517 0 1 0 0-3.034a1.517 1.517 0 0 0 0 3.034"/><path stroke-linecap="round" stroke-linejoin="round" d="M17.61 9.75a7.15 7.15 0 0 0-1.485-3.668c-2.593-3.305-7.64-3.631-10.637-.687c-2.971 2.919-2.764 7.993.5 10.647a7.15 7.15 0 0 0 3.64 1.55m3.122-1.024h7.159a3.341 3.341 0 1 1 0 6.682h-6.2"/><use href="#covidCovid19VirusReinfected0" stroke-linecap="round" stroke-linejoin="round"/><path stroke-linecap="round" stroke-linejoin="round" d="M12.75 16.568h7.159a3.341 3.341 0 1 1 0 6.682h-6.2"/><use href="#covidCovid19VirusReinfected0" stroke-linecap="round" stroke-linejoin="round"/><path d="M10.15 14.492a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CovidNineteenVirusThree(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M7.851 4.958L6.215 3.322a2.046 2.046 0 0 0-2.893 2.893l1.6 1.694m11.227 11.133l1.636 1.636a2.046 2.046 0 0 0 2.893-2.893l-1.636-1.636m-14.084 0l-1.636 1.636a2.046 2.046 0 0 0 2.893 2.893l1.636-1.637M19.075 7.909l1.589-1.694a2.047 2.047 0 1 0-2.893-2.893L16.14 4.953"/><path stroke-linecap="round" stroke-linejoin="round" d="M21.205 9.955h-1.291a8.181 8.181 0 0 0-5.869-5.869V2.8a2.045 2.045 0 0 0-4.09 0v1.286a8.181 8.181 0 0 0-5.869 5.869H2.8a2.045 2.045 0 0 0 0 4.09h1.286a8.181 8.181 0 0 0 5.869 5.869v1.291a2.045 2.045 0 0 0 4.09 0v-1.291a8.182 8.182 0 0 0 5.869-5.869h1.291a2.045 2.045 0 1 0 0-4.09"/><path stroke-linecap="round" stroke-linejoin="round" d="M13.534 13.023a2.557 2.557 0 1 0 0-5.114a2.557 2.557 0 0 0 0 5.114"/><path d="M9.443 15.955a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CovidNineteenVirusTwo(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.934 7.263L3.269 5.6A1.648 1.648 0 1 1 5.6 3.269l1.663 1.665m0 14.132L5.6 20.731A1.647 1.647 0 1 1 3.269 18.4l1.665-1.666m14.132.003l1.665 1.663a1.648 1.648 0 0 1-2.331 2.331l-1.666-1.665m.003-14.132L18.4 3.269A1.648 1.648 0 0 1 20.731 5.6l-1.665 1.663"/><path d="M20.347 10.354a8.514 8.514 0 0 0-6.7-6.7m-.001 16.693a8.515 8.515 0 0 0 6.7-6.7m-16.693-.001a8.514 8.514 0 0 0 6.7 6.7m.001-16.693a8.514 8.514 0 0 0-6.7 6.7"/><path d="M9.805 6.787c.351-.352.549-.828.549-1.325V2.4a1.646 1.646 0 1 1 3.292 0v3.062c.002.498.201.974.554 1.325M6.787 14.2a1.877 1.877 0 0 0-1.325-.549H2.4a1.646 1.646 0 1 1 0-3.292h3.062c.497 0 .973-.198 1.325-.549m7.413 7.403a1.877 1.877 0 0 0-.549 1.325V21.6a1.646 1.646 0 1 1-3.292 0v-3.062c0-.497-.198-.973-.549-1.325m7.403-7.408c.352.351.828.549 1.325.549H21.6a1.646 1.646 0 0 1 0 3.292h-3.062c-.497 0-.973.198-1.325.549"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CovidNineteenVirusWarningOne(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12.017.75H8.983m1.517 0v2.6m-5.822-.817L3.606 3.606L2.533 4.678m1.073-1.072l1.838 1.838M.75 8.983v3.034m0-1.517h2.6m-.817 5.822l1.073 1.072l1.072 1.073m-1.072-1.073l1.838-1.838M18.467 4.678l-1.073-1.072l-1.072-1.073m1.072 1.073l-1.838 1.838M8.983 9.85a1.517 1.517 0 1 0 0-3.034a1.517 1.517 0 0 0 0 3.034"/><path stroke-linecap="round" stroke-linejoin="round" d="M17.284 8.252a7.148 7.148 0 1 0-7.656 9.34M17.25 18v-3m5.813 5.682a1.773 1.773 0 0 1-1.587 2.568h-8.452a1.773 1.773 0 0 1-1.587-2.568l4.226-8.451a1.774 1.774 0 0 1 3.174 0z"/><path d="M10.092 14.492a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75M17.25 21a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CovidNineteenVirusWarningThree(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M10 12.583V6.667m5.472 3.346L11.333 1.58a1.485 1.485 0 0 0-2.666 0L.878 17.447A1.252 1.252 0 0 0 2 19.25h7.75M18 21a3 3 0 1 0 0-6a3 3 0 0 0 0 6m-.5-8.25h1m-.5 0V15m3.359-1.066l.707.707m-.354-.353l-1.591 1.591M23.25 17.5v1m0-.5H21m1.066 3.359l-.707.707m.353-.354l-1.591-1.591M18.5 23.25h-1m.5 0V21m-3.359 1.066l-.707-.707m.354.353l1.591-1.591M12.75 18.5v-1m0 .5H15m-1.066-3.359l.707-.707m-.353.354l1.591 1.591"/><path d="M10 15.983a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CovidNineteenVirusWarningTwo(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M17.25 18v-3m5.813 5.682a1.773 1.773 0 0 1-1.587 2.568h-8.452a1.773 1.773 0 0 1-1.587-2.568l4.226-8.451a1.774 1.774 0 0 1 3.174 0zm-18.88-7.928H2.162a1.413 1.413 0 0 1 0-2.825h2.021m5.746-5.746V2.162a1.413 1.413 0 0 1 2.825 0v2.021"/><path stroke-linecap="round" stroke-linejoin="round" d="M15.405 5.28a7.3 7.3 0 0 0-8.128 0M18.288 9.1a7.255 7.255 0 0 0-.888-1.823M7.277 17.4a7.282 7.282 0 0 0 1.959.926M5.28 7.277a7.3 7.3 0 0 0 0 8.128"/><path stroke-linecap="round" stroke-linejoin="round" d="M6.848 9.511c0-.426-.17-.835-.471-1.137L4.518 6.515a1.414 1.414 0 0 1 2-2l1.856 1.862c.302.301.71.47 1.137.471m0 8.987a1.606 1.606 0 0 0-1.137.471l-1.859 1.859a1.414 1.414 0 0 1-2-2l1.859-1.86c.301-.301.47-.71.471-1.136m6.327-6.321c.426 0 .835-.17 1.136-.471l1.86-1.859a1.414 1.414 0 1 1 2 2l-1.862 1.856"/><path d="M17.25 21a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraphCuredDecreasing(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.777 22.5h-21v-21m21 4.5a1 1 0 0 0-1-1h-2.5V2.5a1 1 0 0 0-1-1h-1.5a1 1 0 0 0-1 1V5h-2.5a1 1 0 0 0-1 1v1.5a1 1 0 0 0 1 1h2.5V11a1 1 0 0 0 1 1h1.5a1 1 0 0 0 1-1V8.5h2.5a1 1 0 0 0 1-1z"/><path d="M.777 9.5h1.4a10.826 10.826 0 0 1 8.6 4.25a10.824 10.824 0 0 0 8.6 4.25h3.846"/><path d="M20.833 15.609L23.223 18l-2.39 2.391"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraphCuredIncreasing(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.755 22.5h-21v-21"/><path d="M15.755 6a1 1 0 0 0-1-1h-2.5V2.5a1 1 0 0 0-1-1h-1.5a1 1 0 0 0-1 1V5h-2.5a1 1 0 0 0-1 1v1.5a1 1 0 0 0 1 1h2.5V11a1 1 0 0 0 1 1h1.5a1 1 0 0 0 1-1V8.5h2.5a1 1 0 0 0 1-1zm-15 13.512h2.7c8.9 0 16.065-3.387 17.955-12.762"/><path d="m18.571 8.582l2.842-1.832l1.832 2.842"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraphCuredStable(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.787 22.5h-21v-21"/><path d="M16.537 6a1 1 0 0 0-1-1h-2.5V2.5a1 1 0 0 0-1-1h-1.5a1 1 0 0 0-1 1V5h-2.5a1 1 0 0 0-1 1v1.5a1 1 0 0 0 1 1h2.5V11a1 1 0 0 0 1 1h1.5a1 1 0 0 0 1-1V8.5h2.5a1 1 0 0 0 1-1zm4.285 9.234l2.391 2.391l-2.391 2.391"/><path d="M.787 17.625h.5a10.6 10.6 0 0 0 4.737-1.118L8.446 15.3a6.352 6.352 0 0 1 5.683 0l2.422 1.211a10.6 10.6 0 0 0 4.736 1.118h1.926"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraphDeathRateDecreasing(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M21.777 6.375a4.875 4.875 0 1 0-7.877 3.816v.309a1.5 1.5 0 0 0 1.5 1.5h3a1.5 1.5 0 0 0 1.5-1.5v-.309a4.847 4.847 0 0 0 1.877-3.816M16.902 10.5V12m4.875 10.5h-21v-21"/><path stroke-linecap="round" stroke-linejoin="round" d="M.777 9.5h1.4a10.826 10.826 0 0 1 8.6 4.25a10.824 10.824 0 0 0 8.6 4.25h3.846"/><path stroke-linecap="round" stroke-linejoin="round" d="M20.833 15.609L23.223 18l-2.39 2.391"/><path d="M18.777 7.113a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75m-3.752.75a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraphDeathRateIncreasing(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M21.753 22.5h-21v-21"/><path stroke-linecap="round" stroke-linejoin="round" d="M15.378 6.375A4.876 4.876 0 1 0 7.5 10.191v.309A1.5 1.5 0 0 0 9 12h3a1.5 1.5 0 0 0 1.5-1.5v-.309a4.848 4.848 0 0 0 1.878-3.816M10.503 10.5V12M.758 19.512h2.7c8.894 0 16.064-3.387 17.954-12.762"/><path stroke-linecap="round" stroke-linejoin="round" d="m18.574 8.582l2.841-1.832l1.832 2.842"/><path d="M12.378 7.125a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75m-3.752.75a.375.375 0 1 1 0-.75m0 .75a.375.375 0 1 0 0-.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraphDeathRateStable(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M21.787 22.5h-21v-21"/><path stroke-linecap="round" stroke-linejoin="round" d="M16.162 6.375a4.876 4.876 0 1 0-7.875 3.816v.309a1.5 1.5 0 0 0 1.5 1.5h3a1.5 1.5 0 0 0 1.5-1.5v-.309a4.846 4.846 0 0 0 1.875-3.816M11.287 10.5V12m9.535 3.234l2.391 2.391l-2.391 2.391"/><path stroke-linecap="round" stroke-linejoin="round" d="M.787 17.625h.5a10.6 10.6 0 0 0 4.737-1.118L8.446 15.3a6.352 6.352 0 0 1 5.683 0l2.422 1.211a10.6 10.6 0 0 0 4.736 1.118h1.926"/><path d="M13.162 7.125a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75m-3.752.75a.375.375 0 1 1 0-.75m0 .75a.375.375 0 1 0 0-.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraphDocumentInfectedReport(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.75 20.532a3.137 3.137 0 1 0 0-6.274a3.137 3.137 0 0 0 0 6.274m-.523-8.627h1.046m-.523 0v2.353m3.512-1.114l.74.739m-.37-.37l-1.664 1.664m3.272 1.695v1.046m0-.523h-2.353m1.115 3.513l-.74.739m.37-.37l-1.664-1.663m-1.695 3.271h-1.046m.523 0v-2.353m-3.512 1.115l-.74-.739m.37.369l1.664-1.663m-3.272-1.696v-1.046m0 .523h2.353m-1.115-3.512l.74-.739m-.37.369l1.664 1.664M3.76 7.115v9h4.5"/><path d="m3.76 13.115l2.689-2.689a1.5 1.5 0 0 1 2.122 0l.581.58a1.5 1.5 0 0 0 2.347-.289l.357-.6"/><path d="M9.01 20.615H2.26a1.5 1.5 0 0 1-1.5-1.5v-16.5a1.5 1.5 0 0 1 1.5-1.5h10.629a1.5 1.5 0 0 1 1.06.439l2.872 2.872a1.5 1.5 0 0 1 .439 1.06v2.379"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraphInfectedDecreasing(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.777 22.5h-21v-21m15.75 8.25a3 3 0 1 0 0-6a3 3 0 0 0 0 6m-.5-8.25h1m-.5 0v2.25m3.359-1.066l.707.707m-.354-.353l-1.591 1.591m3.129 1.621v1m0-.5h-2.25m1.066 3.359l-.707.707m.353-.354l-1.591-1.591M17.027 12h-1m.5 0V9.75m-3.359 1.066l-.707-.707m.353.353l1.591-1.591M11.277 7.25v-1m0 .5h2.25m-1.066-3.359l.707-.707m-.354.354l1.591 1.591"/><path d="M.777 8.875h1.4a10.826 10.826 0 0 1 8.6 4.25a10.824 10.824 0 0 0 8.6 4.25h3.846"/><path d="m20.833 14.984l2.39 2.391l-2.39 2.391"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraphInfectedIncreasing(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.755 22.5h-21v-21"/><path d="M.755 19.5h2.7c8.9 0 16.065-3.387 17.955-12.762M10.505 9.75a3 3 0 1 0 0-6a3 3 0 0 0 0 6m-.5-8.25h1m-.5 0v2.25m3.359-1.066l.707.707m-.354-.353l-1.591 1.591m3.129 1.621v1m0-.5h-2.25m1.066 3.359l-.707.707m.353-.354l-1.591-1.591M11.005 12h-1m.5 0V9.75m-3.359 1.066l-.707-.707m.354.353l1.591-1.591M5.255 7.25v-1m0 .5h2.25M6.439 3.391l.707-.707m-.353.354l1.591 1.591"/><path d="m18.571 8.57l2.842-1.832l1.832 2.842"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraphInfectedStable(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.787 22.5h-21v-21m20.035 13.734l2.391 2.391l-2.391 2.391M11.287 9.75a3 3 0 1 0 0-6a3 3 0 0 0 0 6m-.5-8.25h1m-.5 0v2.25m3.359-1.066l.707.707M15 3.038l-1.591 1.591m3.128 1.621v1m0-.5h-2.25m1.066 3.359l-.707.707m.354-.354l-1.591-1.591M11.787 12h-1m.5 0V9.75m-3.358 1.066l-.707-.707m.353.353l1.591-1.591M6.037 7.25v-1m0 .5h2.25M7.222 3.391l.707-.707m-.354.354l1.591 1.591"/><path d="M.787 17.625h.5a10.6 10.6 0 0 0 4.737-1.118L8.446 15.3a6.352 6.352 0 0 1 5.683 0l2.422 1.211a10.6 10.6 0 0 0 4.736 1.118h1.926"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MutationOne(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12.408 1.208H9.375m1.517 0v2.6M5.07 2.992L3.997 4.064L2.925 5.136m1.072-1.072l1.839 1.839M1.142 9.442v3.033m0-1.517h2.6m-.817 5.822l1.072 1.073l1.073 1.072m-1.073-1.072l1.839-1.839M18.858 5.136l-1.072-1.072l-1.072-1.072m1.072 1.072l-1.839 1.839m-6.572 4.406a1.517 1.517 0 1 0 0-3.034a1.517 1.517 0 0 0 0 3.034"/><path stroke-linecap="round" stroke-linejoin="round" d="M17.3 7.791a7.15 7.15 0 1 0-8.442 10.018m10.462 3.137H14.7a1.847 1.847 0 0 1-1.842-1.846v-4.616m3.542-1.846h4.615a1.844 1.844 0 0 1 1.846 1.846V19.1"/><path stroke-linecap="round" stroke-linejoin="round" d="m17.474 19.099l1.846 1.847l-1.846 1.846m.769-8.308l-1.846-1.846l1.846-1.846"/><path d="M9.367 14.6a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MutationStronger(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12.408.966H9.375m1.517 0v2.6M5.07 2.749L3.997 3.822L2.925 4.894m1.072-1.072L5.836 5.66M1.142 9.2v3.033m0-1.517h2.6m-.817 5.822l1.072 1.073l1.073 1.072m-1.073-1.072l1.839-1.839M18.858 4.894l-1.072-1.072l-1.073-1.073m1.073 1.073L15.947 5.66m-6.572 4.407a1.517 1.517 0 1 0 0-3.034a1.517 1.517 0 0 0 0 3.034"/><path stroke-linecap="round" stroke-linejoin="round" d="M17.748 8.683a7.15 7.15 0 1 0-8.89 8.884m3 5.467h8.458a2.533 2.533 0 0 0 2.542-2.523v-6.606a2.189 2.189 0 0 0-1.9-2.211a2.1 2.1 0 0 0-1.152 3.969a5.994 5.994 0 0 0 0 4.428a4.674 4.674 0 0 0-7.338-1.265"/><path d="M11.633 13.625a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MutationTemperatureChange(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.25 10.179a3.429 3.429 0 1 0 0-6.858a3.429 3.429 0 0 0 0 6.858M16.679.75h1.142m-.571 0v2.571m3.839-1.218l.808.808m-.404-.404l-1.819 1.819m3.576 1.853v1.142m0-.571h-2.571m1.218 3.839l-.808.808m.404-.404l-1.819-1.819m-1.853 3.576h-1.142m.571 0v-2.571m-3.839 1.218l-.808-.808m.404.404l1.819-1.819M11.25 7.321V6.179m0 .571h2.571m-1.218-3.839l.808-.808m-.404.404l1.819 1.819M8.25 15.418V3.75a3 3 0 1 0-6 0v11.668a4.493 4.493 0 1 0 7.5 3.332a4.472 4.472 0 0 0-1.5-3.332m-3-10.168v12"/><path d="M5.25 20.25a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MutationTwo(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.528 19.777a3.714 3.714 0 1 0 0-7.428a3.714 3.714 0 0 0 0 7.428M6.909 9.563h1.238m-.619 0v2.786m4.159-1.32l.875.876m-.438-.438l-1.969 1.97m3.873 2.007v1.238m0-.619h-2.785m1.319 4.159l-.875.875m.437-.438l-1.969-1.969m-2.008 3.873H6.909m.619 0v-2.786m-4.158 1.32l-.876-.875m.438.437l1.97-1.969m-3.874-2.008v-1.238m0 .619h2.786m-1.32-4.158l.876-.876m-.438.438l1.97 1.97m13.446-4.283a2.714 2.714 0 1 0 0-5.428a2.714 2.714 0 0 0 0 5.428m-.452-7.464h.905m-.453 0v2.035m3.039-.964l.64.64m-.32-.32l-1.44 1.439m2.831 1.467v.905m0-.452h-2.036m.965 3.038l-.64.64m.32-.32l-1.44-1.439m-1.466 2.831h-.905m.452 0V9.154m-3.039.964l-.64-.64m.32.32l1.44-1.439m-2.831-1.467v-.905m0 .453h2.036m-.965-3.039l.64-.64m-.32.32l1.44 1.439M2.748 7.9V3.283a1.846 1.846 0 0 1 1.846-1.846H9.21"/><path d="M4.594 6.052L2.748 7.898L.902 6.052m20.154 8.385v6.154a1.847 1.847 0 0 1-1.846 1.846h-2.616"/><path d="m19.21 16.283l1.846-1.846l1.846 1.846"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonalHygieneCleanBottleShield(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M23 15.749a7.67 7.67 0 0 1-6 7.5a7.67 7.67 0 0 1-6-7.5V13.5a1.5 1.5 0 0 1 1.5-1.5h9a1.5 1.5 0 0 1 1.5 1.5zm-5.999-.75v4.5m-2.251-2.25h4.501"/><path d="M8.5 20.253H4a3 3 0 0 1-3-3v-7.5a3 3 0 0 1 3-3h6a3 3 0 0 1 2.925 2.331M1 2.25l.622-.621A3 3 0 0 1 3.741.75h7.009M10 6.751H4v-1.5a1.5 1.5 0 0 1 1.5-1.5h3a1.5 1.5 0 0 1 1.5 1.5zM6.999 3.75v-3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonalHygieneCleanBottleVirus(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.608 1.272h1.329m-.664 0v2.956m4.463-1.382l.94.94m-.47-.47L19.092 5.43m4.158 2.155v1.329m0-.664h-2.99m1.416 4.464l-.94.939m.47-.47l-2.114-2.114m-2.155 4.158h-1.329m.665 0v-2.99M9.749 9.229H3.75a3 3 0 0 0-3 3v7.499a3 3 0 0 0 3 3h5.999a3 3 0 0 0 3-3v-7.499a3 3 0 0 0-3-3M.75 4.729l.621-.621a3 3 0 0 1 2.121-.879H10.5m-.751 6h-6v-1.5a1.5 1.5 0 0 1 1.5-1.5h3a1.5 1.5 0 0 1 1.5 1.5zm-3-3v-3m0 9.998v6m-2.999-3h5.999m3.856-10.999a4 4 0 0 1 6.387 1.586c.23.605.308 1.257.23 1.9a4 4 0 0 1-3.971 3.514"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonalHygieneCleanBottleVirusBlock(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15.28 19.72a3.145 3.145 0 0 1-.029-4.473a3.143 3.143 0 0 1 4.474.03m.92 2.222a3.143 3.143 0 0 1-3.143 3.143m-.524-8.643h1.047m-.523 0v2.356m5.5 2.62v1.047m0-.523h-2.357m1.117 3.519l-.742.74m.37-.37l-1.666-1.668m-1.698 3.279h-1.049m.525 0v-2.357m-5.5-2.619v-1.049m0 .525h2.356m-1.115-3.519l.74-.74m-.37.37l1.667 1.667m7.722-3.278l-11 11m-3.503-2.751h-4.5a3 3 0 0 1-3-3V9.749a3 3 0 0 1 3-3h6a3 3 0 0 1 3 3v.75M1 2.252l.621-.621A3 3 0 0 1 3.742.75h7.009"/><path d="M9.999 6.75h-6v-1.5a1.5 1.5 0 0 1 1.5-1.5h3a1.5 1.5 0 0 1 1.5 1.5zM7 3.752v-3M7 9.75v5.999M4 12.75h5.998"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonalHygieneCleanGel(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.022.794h.665m-.665 0V3.75m4.464-1.382l.94.94m-.47-.47l-2.114 2.114M23 7.107v1.33m0-.665h-2.991m1.417 4.464l-.94.94m.47-.47l-2.114-2.114m-2.155 4.158h-1.33m.665 0v-2.991M10 20.25H4A108.39 108.39 0 0 1 1 1.536A.751.751 0 0 1 1.75.75h10.5a.752.752 0 0 1 .75.786a108.395 108.395 0 0 1-3 18.714m0 0v2.25a.75.75 0 0 1-.75.75h-4.5A.75.75 0 0 1 4 22.5v-2.25m12-16.5a4 4 0 1 1-1.221 7.81M7 4.75v6m-3-3h6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonalHygieneCleanToothpaste(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.972 20.25h-7.5v1.5a1.5 1.5 0 0 0 1.5 1.5h4.5a1.5 1.5 0 0 0 1.5-1.5zm0 0h-7.5L1.738 2.478A1.5 1.5 0 0 1 3.221.75h10a1.5 1.5 0 0 1 1.482 1.728zm8.308 3v-11l2-3v-8m-4 1h4m-4 3h4m-4 3h4m-14-2v6m-3-3h6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonalHygieneHandLiquidSoap(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22.012 14.74a3.504 3.504 0 0 1-7.008 0c0-2.628 3.5-7.009 3.5-7.009s3.508 4.381 3.508 7.009M9.998 9.233H3.99a2.002 2.002 0 0 0-2.002 2.002v10.013c0 1.106.896 2.002 2.002 2.002h6.008A2.002 2.002 0 0 0 12 21.248V11.235a2.002 2.002 0 0 0-2.002-2.002M4.766 6.23h4.456a.776.776 0 0 1 .778.775v2.228H3.99V7.005a.776.776 0 0 1 .776-.775M14 2.752l-.447-.895A2 2 0 0 0 11.764.75H2.989m4.005 13.489v4.005m-2.002-2.002h4.004M6.994.75v5.48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonalHygieneHandSanitizerLiquidDrop(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18.251 21.664h3a2 2 0 0 0 2-2v-10a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v6m6-8.001h-4V5.438a.775.775 0 0 1 .775-.775h2.45a.775.775 0 0 1 .775.775zm2 4.001h-8m8 6h-5M.749 18.836l2.937 2.35a5.247 5.247 0 0 0 3.279 1.15H13a1.75 1.75 0 0 0 1.75-1.75v0a1.75 1.75 0 0 0-1.75-1.75H8.478"/><path d="M.908 13.586h3.5c1.037 0 2.051.308 2.914.883l2 1.333a1.633 1.633 0 0 1 .518 2.307v0a1.634 1.634 0 0 1-2.089.555l-2.627-1.578m7.127-7.922a2.5 2.5 0 0 1-5 0c0-1.875 2.5-5 2.5-5s2.5 3.125 2.5 5m7-7.5v3m-6-1l.447-.895a2 2 0 0 1 1.789-1.105h6.764"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonalHygieneHandSanitizerLiquidOne(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m1.121 19.788l2.905 2.324a5.187 5.187 0 0 0 3.243 1.138h5.967a1.731 1.731 0 0 0 1.731-1.731v0a1.731 1.731 0 0 0-1.731-1.731H8.765"/><path d="M1.278 14.6H4.74c1.026 0 2.029.304 2.882.873L9.6 16.788a1.615 1.615 0 0 1 .513 2.281v0a1.614 1.614 0 0 1-2.066.549l-2.6-1.56m17.432-7.418a4.945 4.945 0 1 1-9.89 0c0-3.709 4.945-9.89 4.945-9.89s4.945 6.181 4.945 9.89m-6.923-.989h3.956m-1.978-1.978v3.956"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonalHygieneHandSanitizerLiquidThree(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.125 20.076a6.505 6.505 0 0 1-8-6.326c0-4.875 6.5-13 6.5-13a46.389 46.389 0 0 1 5.306 8.5"/><path d="M22.875 15.75a7.669 7.669 0 0 1-6 7.5a7.669 7.669 0 0 1-6-7.5V13.5a1.5 1.5 0 0 1 1.5-1.5h9a1.5 1.5 0 0 1 1.5 1.5zm-6-.75v4.5m-2.25-2.25h4.5m-11.5 0a3.5 3.5 0 0 1-3.5-3.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonalHygieneHandSanitizerLiquidTwo(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.031 10.75a5 5 0 1 1-10 0c0-3.75 5-10 5-10s5 6.25 5 10m-7-1h4m-2-2v4m-8.312 11.5l-1.846-2.308a4.124 4.124 0 0 1-.9-2.576v-4.741a1.375 1.375 0 0 1 1.371-1.375v0a1.376 1.376 0 0 1 1.375 1.375v3.552"/><path d="M7.844 23.125v-2.75a4.13 4.13 0 0 0-.694-2.29L6.1 16.516a1.284 1.284 0 0 0-1.813-.408v0a1.285 1.285 0 0 0-.436 1.642l1.24 2.063m15.19 3.437l1.846-2.308c.584-.732.901-1.64.9-2.576v-4.741a1.375 1.375 0 0 0-1.375-1.375v0a1.376 1.376 0 0 0-1.375 1.375v3.552"/><path d="M16.156 23.125v-2.75c0-.815.242-1.612.694-2.29l1.05-1.569a1.285 1.285 0 0 1 1.813-.408v0a1.286 1.286 0 0 1 .436 1.642l-1.24 2.063"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonalHygieneHandSanitizerLiquidVirusBlock(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15.4 19.85a3.142 3.142 0 1 1 4.444-4.444m.924 2.219a3.144 3.144 0 0 1-3.143 3.143m-.524-8.643h1.048m-.524 0v2.357m5.5 2.619v1.048m0-.524h-2.357m1.116 3.519l-.74.74m.37-.37l-1.667-1.667m-1.698 3.278h-1.048m.524 0v-2.357m-5.5-2.619v-1.048m0 .524h2.357m-1.116-3.519l.74-.74m-.37.37l1.667 1.667m7.722-3.278l-11 11m-2.188-3.503A6.459 6.459 0 0 1 1 13.661C1 8.82 7.454.75 7.454.75a43.358 43.358 0 0 1 5.682 9.435"/><path d="M7.454 17.137a3.476 3.476 0 0 1-3.476-3.476"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonalHygieneHandSanitizerSpray(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M10.456 8.843H4.281a2.058 2.058 0 0 0-2.058 2.058v10.291c0 1.137.921 2.058 2.058 2.058h6.175a2.058 2.058 0 0 0 2.058-2.058V10.901a2.058 2.058 0 0 0-2.058-2.058M5.079 1.871h4.58a.8.8 0 0 1 .8.8v6.172H4.282V2.668a.8.8 0 0 1 .8-.8zm2.29 12.117v4.116m-2.058-2.058h4.116"/><path d="M21.356 5.71a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75m-1.217 5.052a.375.375 0 1 1 0-.75m0 .75a.375.375 0 1 0 0-.75m0-7.67a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75M16.21 4.118a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75m0 4.118a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonalHygieneHandSanitizerSprayVirusBlock(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.673 19.577a3.49 3.49 0 0 1 4.937-4.937m1.022 2.469a3.491 3.491 0 0 1-3.491 3.49M16.56 11h1.163m-.582 0v2.618m6.109 2.909v1.163m0-.581h-2.618m1.24 3.908l-.823.822m.412-.411l-1.851-1.851m-1.887 3.64H16.56m.581 0v-2.618m-6.108-2.909v-1.163m0 .582h2.618m-1.24-3.908l.822-.823m-.411.411l1.851 1.851M23.25 11L11.033 23.217m-2.508-.982h-4.85A2.925 2.925 0 0 1 .75 19.31v-8.539A2.925 2.925 0 0 1 1.607 8.7l2.068-2.066h5.851L11.6 8.7c.36.36.62.809.753 1.3M5.626.783h1.95a1.95 1.95 0 0 1 1.95 1.95v3.9H3.675v-3.9A1.95 1.95 0 0 1 5.626.783m9.28 1.144l2.945-.798m-2.945 5.583l2.945.798m-2.945-3.19h2.945"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonalHygieneHandSanitizerVirusBlock(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.408 21.443H2.962a2.177 2.177 0 0 1-2.178-2.178V8.374A2.178 2.178 0 0 1 2.962 6.2h7.757A2.178 2.178 0 0 1 12.9 8.374v.956"/><path d="M4.052.75H9.63a1.089 1.089 0 0 1 1.089 1.089V6.2H2.962V1.839A1.089 1.089 0 0 1 4.052.75m10.66 18.891a3.461 3.461 0 0 1 4.894-4.895m1.014 2.447a3.46 3.46 0 0 1-3.461 3.461m-.577-9.517h1.154m-.577 0v2.596m6.057 2.884v1.153m0-.577H20.62m1.23 3.875l-.816.816m.408-.408l-1.836-1.835m-1.87 3.609h-1.154m.577 0v-2.596m-6.056-2.884v-1.153m0 .576h2.595m-1.229-3.874l.815-.816m-.407.408l1.835 1.835m8.504-3.609L11.103 23.25M.784 10.552h8.58m-8.58 6.535h6.535"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonalHygieneHandSanitizerVirusShield(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.319 21.66H2.912A2.162 2.162 0 0 1 .75 19.5V8.685a2.162 2.162 0 0 1 2.162-2.162h7.9a2.162 2.162 0 0 1 2.162 2.162v.747m-8.98-8.315h5.741A1.08 1.08 0 0 1 10.816 2.2v4.323h-7.9V2.2a1.081 1.081 0 0 1 1.078-1.083M.75 10.848h8.315M.75 17.335h6.487"/><path d="M23.25 15.546a7.5 7.5 0 0 1-5.87 7.337a7.5 7.5 0 0 1-5.869-7.337v-2.2a1.467 1.467 0 0 1 1.467-1.467h8.805a1.467 1.467 0 0 1 1.467 1.467zm-5.87-.733v4.402m-2.201-2.201h4.403"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonalHygieneHandSoapOne(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M.75 21.252h1.9c.233 0 .463.054.671.158l1.733.867a4.507 4.507 0 0 0 2.012.475H18.75a1.5 1.5 0 1 0 0-3H21a1.5 1.5 0 1 0 0-3h.75a1.5 1.5 0 1 0 0-3h-1.5a1.5 1.5 0 1 0 0-3H15a1.5 1.5 0 1 0 0-3H9.158a4.5 4.5 0 0 0-3.744 2L4.2 11.584a1.5 1.5 0 0 1-1.248.668H.75m12-1.5H15m1.5 6H21m-4.5-3h3.75m-4.5 6h3M23 8.248a1.992 1.992 0 0 0-2.146-1.985c.093-.33.142-.672.146-1.015a3.99 3.99 0 0 0-7.2-2.383a2.986 2.986 0 0 0-3.8.147"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonalHygieneHandSoapTwo(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.928 7.272H4.89a3.358 3.358 0 0 0-3.358 3.358v9.236a3.358 3.358 0 0 0 3.358 3.358h5.038a3.358 3.358 0 0 0 3.359-3.358V10.63a3.358 3.358 0 0 0-3.359-3.358"/><path d="M16.488 14.572a2.992 2.992 0 0 0 3.085-4.962a4.967 4.967 0 0 0-7.215-6.698a2.983 2.983 0 0 0-5.844.863m13.961 18.452a1.994 1.994 0 1 0 0-3.987a1.994 1.994 0 0 0 0 3.987M7.514 13.254v3.988M5.52 15.249h3.988"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonalHygieneHandWash(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m1.1 19.616l3.049 2.44a5.452 5.452 0 0 0 3.404 1.194h6.264a1.817 1.817 0 0 0 1.817-1.817v0a1.818 1.818 0 0 0-1.817-1.817H9.124"/><path d="M1.264 14.166H4.9c1.077 0 2.13.319 3.025.917L10 16.466a1.7 1.7 0 0 1 .539 2.395v0a1.7 1.7 0 0 1-2.17.576L5.641 17.8M21.863 5.941a2.6 2.6 0 0 1-5.191 0c0-1.947 2.595-5.191 2.595-5.191s2.596 3.244 2.596 5.191M22.9 18.119a1.917 1.917 0 0 0-1.916-1.919c-.048 0-.093.01-.14.014a3.816 3.816 0 0 0-6.756-3.256a2.85 2.85 0 0 0-3.645.141"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonalHygieneHandWipePaperFour(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m6.5 23.25l1.5-.431a1.25 1.25 0 0 1 .567-.028l1.567.291a3.693 3.693 0 0 0 1.7-.082l9.239-2.658a1.234 1.234 0 0 0-.682-2.371l1.778-.512a1.233 1.233 0 1 0-.682-2.371l.593-.171a1.234 1.234 0 1 0-.683-2.371l-1.185.341a1.234 1.234 0 0 0-.682-2.371l-4.15 1.194a1.234 1.234 0 1 0-.683-2.372l-4.617 1.329a3.7 3.7 0 0 0-2.5 2.435l-.56 1.723a1.234 1.234 0 0 1-.835.811l-1.737.5m9.144-3.915l1.779-.511m2.55 4.401l3.557-1.023m-4.239-1.348l2.964-.853m-2.193 5.766l2.372-.682"/><path d="m22.619 8.092l.335-1.34l-.5-2.5l-2.5.5l-.5-2.5l-2.5.5l-1.5-2c-1 2-2.432 3.092-6.778 3.865c-4.271.757-6.763 2.317-7.729 4.636l2.5.5l-.5 2.5l1.5.3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonalHygieneHandWipePaperOne(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M23.25 15.3h-1.611a1.1 1.1 0 0 1-.916-.49l-.893-1.341A3.3 3.3 0 0 0 17.084 12H12.75v3h-3.8a1.1 1.1 0 0 0-1.1 1.1l1.1 5.8a1.1 1.1 0 0 0 1.1 1.1h8.571a3.3 3.3 0 0 0 1.476-.348l1.271-.636a1.1 1.1 0 0 1 .492-.116h1.39M2.75 9c1.105 0 2-1.79 2-4s-.895-4-2-4s-2 1.79-2 4s.895 4 2 4"/><path d="M2.75 9h10c1.1 0 2-1.791 2-4s-.9-4-2-4h-10m3 19l-2-1l-3 1V5m12 4v3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonalHygieneHandWipePaperThree(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.5.75v5a2 2 0 0 0 2 2h1m18-7v5a2 2 0 0 1-2 2h-1M9.786 18.893l1.714.857l2-1l2 1l2-1l2 1V4.25h-15v10.5m-3 8.5v-6.824a5 5 0 0 1 2.226-4.16l.774-.516"/><path d="m7.5 21.75l2.714-3.393a1.887 1.887 0 0 0-.427-2.749v0a1.886 1.886 0 0 0-2.381.236L5.5 17.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonalHygieneHandWipePaperTwo(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M23.083 7.916a3.583 3.583 0 0 1-7.166 0C15.917 5.229 19.5.75 19.5.75s3.583 4.479 3.583 7.166m-5.264 6.335c.859 2.035 2.006 3.466-.084 6h-4.818m-6-10.201c-3.314 0-6-1.074-6-2.4v13.2c0 1.326 2.686 2.4 6 2.4s6-1.074 6-2.4V7.65c0 1.326-2.687 2.4-6 2.4"/><path d="M6.917 10.05c3.314 0 6-1.075 6-2.4s-2.686-2.4-6-2.4s-6 1.075-6 2.4s2.686 2.4 6 2.4m-.6-2.4h1.2m-.6-2.4h6.47"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuarantinePlaceBed(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.312 8.607a2.857 2.857 0 1 0 0-5.714a2.857 2.857 0 0 0 0 5.714M9.836.75h.953m-.477 0v2.143m3.199-1.015l.674.673m-.337-.337L12.333 3.73m2.979 1.544v.952m0-.476H13.17m1.015 3.199l-.674.673m.337-.336L12.333 7.77m-1.544 2.98h-.953m.476 0V8.607M7.114 9.622l-.674-.673m.337.337L8.292 7.77m-2.98-1.544v-.952m0 .476h2.143M6.44 2.551l.674-.673m-.337.336L8.292 3.73m-1.48 15.02a2.25 2.25 0 1 0 0-4.5a2.25 2.25 0 0 0 0 4.5m4.5 0v-4.5h9.876a2 2 0 0 1 2 2v2.5m-.001 0H.812v3h22.375zM.812 23.25v-9m22.376 9v-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuarantinePlaceHospital(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18.24 20.729a2.857 2.857 0 1 0 0-5.714a2.857 2.857 0 0 0 0 5.714m-.476-7.857h.952m-.476 0v2.143M21.439 14l.673.673m-.337-.337l-1.515 1.516m2.98 1.544v.952m0-.476h-2.143m1.015 3.199l-.673.673m.336-.337l-1.515-1.515m-1.544 2.98h-.952m.476 0v-2.143m-3.199 1.015l-.673-.673m.336.336l1.516-1.515m-2.98-1.544v-.952m0 .476h2.143m-1.015-3.199l.673-.673m-.337.336l1.516 1.516M3.227 10.259v7.4a1.233 1.233 0 0 0 1.233 1.233h5.8m9-9.866l-7.885-7.36a2 2 0 0 0-2.729 0L.76 9.026m9.25-1.047v6.006m-3.003-3.003h6.006"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuarantinePlaceHouseOne(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 17.239a2.857 2.857 0 1 0 0-5.714a2.857 2.857 0 0 0 0 5.714m-.476-7.857h.952m-.476 0v2.143m3.199-1.015l.673.673m-.336-.336l-1.516 1.515M17 13.906v.952m0-.476h-2.143m1.015 3.199l-.673.673m.337-.336l-1.516-1.516m-1.544 2.98h-.952m.476 0v-2.143m-3.199 1.015l-.673-.673m.336.337l1.516-1.516M7 14.858v-.952m0 .476h2.143m-1.015-3.199l.673-.673m-.337.337l1.516 1.515"/><path d="M3.75 12.382v9a1.5 1.5 0 0 0 1.5 1.5h13.5a1.5 1.5 0 0 0 1.5-1.5v-9m3-1.5l-9.885-9.226a2 2 0 0 0-2.73 0L.75 10.882"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuarantinePlaceHouseShield(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18.344 21.148a2.804 2.804 0 1 0 0-5.608a2.804 2.804 0 0 0 0 5.608m-.468-7.711h.935m-.467 0v2.103m3.139-.996l.66.661m-.33-.33l-1.487 1.486m2.924 1.515v.935m0-.467h-2.103m.996 3.139l-.66.66m.33-.33l-1.487-1.487m-1.515 2.924h-.935m.468 0v-2.103m-3.139.996l-.661-.66m.331.33l1.486-1.487m-2.924-1.515v-.935m0 .468h2.103m-.996-3.139l.661-.661m-.33.331l1.486 1.486M19.25 9.5V3.76a1.411 1.411 0 0 0-.823-1.292A20.6 20.6 0 0 0 10 .75a20.6 20.6 0 0 0-8.427 1.718A1.411 1.411 0 0 0 .75 3.76v7.224c0 3.532 1.546 8.352 8.228 10.922a2.847 2.847 0 0 0 2.044 0m3.79-11.052V7.068"/><path d="M5.062 7.068v6.286a1.625 1.625 0 0 0 1.625 1.625h3.75"/><path d="m3.437 8.49l4.9-4.283a2.439 2.439 0 0 1 3.211 0l4.889 4.283m-4.875 6.489h-3.25v-3.25a1.625 1.625 0 1 1 3.25 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuarantinePlaceHouseTwo(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12.017.768H8.983m1.517 0v2.6m-5.822-.817L3.606 3.623L2.533 4.696m1.073-1.073l1.838 1.839M.75 9.001v3.033m0-1.516h2.6m-.817 5.821l1.073 1.073l1.072 1.072m-1.072-1.072l1.838-1.839M18.467 4.696l-1.073-1.073l-1.072-1.072m1.072 1.072l-1.838 1.839M8.983 9.868a1.517 1.517 0 1 0 0-3.034a1.517 1.517 0 0 0 0 3.034"/><path stroke-linecap="round" stroke-linejoin="round" d="M16.229 6.24a7.149 7.149 0 1 0-7.762 11.128m9.141-7.611a7.128 7.128 0 0 0-.253-1.272v0c-.24-.81-.62-1.57-1.126-2.245m-3.479 9.69v5.8a1.5 1.5 0 0 0 1.5 1.5h6a1.5 1.5 0 0 0 1.5-1.5v-5.8"/><path stroke-linecap="round" stroke-linejoin="round" d="m11.25 17.243l4.518-3.954a2.252 2.252 0 0 1 2.964 0l4.518 3.954m-4.5 5.989h-3v-3a1.5 1.5 0 1 1 3 0z"/><path d="M9.358 13.368a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuarantinePlaceSelfLockdownOne(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 11.25a2.25 2.25 0 1 0 0-4.5a2.25 2.25 0 0 0 0 4.5m1.5 12l.75-4.5h1.5V16.5a3.75 3.75 0 0 0-7.5 0v2.25h1.5l.75 4.5z"/><path d="M22.272 23.25a.981.981 0 0 0 .978-.978V9.75a1.186 1.186 0 0 0-.377-.8L12 .75L1.127 8.949c-.224.208-.36.495-.377.8v12.523a.981.981 0 0 0 .978.978z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuarantinePlaceSelfLockdownTwo(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 11.25a2.25 2.25 0 1 0 0-4.5a2.25 2.25 0 0 0 0 4.5m1.5 12l.75-4.5h1.5V16.5a3.75 3.75 0 0 0-7.5 0v2.25h1.5l.75 4.5z"/><path d="M22.873 8.949L20.5 7.161V1.25H16v2.518L12 .75L1.127 8.949c-.224.208-.36.495-.377.8v12.522a.98.98 0 0 0 .978.979h20.544a.98.98 0 0 0 .978-.979V9.75a1.186 1.186 0 0 0-.377-.801"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuarantinePlaceTimeCalendarDay(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M6.75 12.007h2v8m8.5-3h-5.5c0-3 4-5 4-5v8"/><path stroke-linecap="round" d="M21.75 3.007H2.25a1.5 1.5 0 0 0-1.5 1.5v17a1.5 1.5 0 0 0 1.5 1.5h19.5a1.5 1.5 0 0 0 1.5-1.5v-17a1.5 1.5 0 0 0-1.5-1.5"/><path d="M.75 8.507h22.5"/><path stroke-linecap="round" d="M7.25 4.993v-4m9 4v-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuarantinePlaceTimeCalendarOne(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M17.25 20.643a3.429 3.429 0 1 0 0-6.858a3.429 3.429 0 0 0 0 6.858m-.571-9.429h1.142m-.571 0v2.572m3.839-1.218l.808.808m-.404-.404l-1.819 1.818m3.576 1.853v1.143m0-.572h-2.571m1.218 3.839l-.808.808m.404-.404l-1.819-1.818m-1.853 3.575h-1.142m.571 0v-2.571m-3.839 1.218l-.808-.808m.404.404l1.819-1.818m-3.576-1.853v-1.143m0 .571h2.571m-1.218-3.838l.808-.808m-.404.404l1.819 1.818m-6.861 3.049H2.3a1.553 1.553 0 0 1-1.55-1.551V3.889A1.552 1.552 0 0 1 2.3 2.337h13.977a1.552 1.552 0 0 1 1.552 1.552v4.364"/><path d="M.757 6.992h17.079"/><path stroke-linecap="round" d="M5.408 3.889V.786m7.763 3.103V.786"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuarantinePlaceTimeCalendarTwo(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M12 18.373a2.714 2.714 0 1 0 0-5.428a2.714 2.714 0 0 0 0 5.428m-.452-7.464h.904m-.452 0v2.036m3.039-.965l.64.64m-.32-.32l-1.44 1.44m2.831 1.467v.904m0-.452h-2.036m.965 3.039l-.64.64m.32-.32l-1.44-1.44m-1.467 2.831h-.904m.452 0v-2.036m-3.039.965l-.64-.64m.32.32l1.44-1.44M7.25 16.111v-.904m0 .452h2.036m-.965-3.039l.64-.64m-.32.32l1.44 1.44"/><path stroke-linecap="round" d="M21.75 2.82H2.25a1.5 1.5 0 0 0-1.5 1.5v17.375a1.5 1.5 0 0 0 1.5 1.5h19.5a1.5 1.5 0 0 0 1.5-1.5V4.32a1.5 1.5 0 0 0-1.5-1.5"/><path d="M.75 8.32h22.5"/><path stroke-linecap="round" d="M7.25 4.805v-4m9 4v-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuarantinePlaceTimeDateDay(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12.75 14.875h2v8m8.5-3h-5.5c0-3 4-5 4-5v8"/><path stroke-linecap="round" stroke-linejoin="round" d="M10.125 17.375a6.87 6.87 0 1 1 6.737-5.5m-5.279-10.75H8.667m1.458 0v2.5M4.527 2.84L3.496 3.871L2.465 4.902m1.031-1.031l1.768 1.768M.75 9.042v2.916m0-1.458h2.5m-.785 5.598l1.031 1.031l1.031 1.031m-1.031-1.031l1.768-1.768m3.403 4.514h2.916m-1.458 0v-2.5m9.375-5.417V9.042m0 1.458H17m.785-5.598l-1.031-1.031l-1.031-1.031m1.031 1.031l-1.768 1.768"/><path stroke-linecap="round" stroke-linejoin="round" d="M8.042 9.875a1.458 1.458 0 1 0 0-2.916a1.458 1.458 0 0 0 0 2.916"/><path d="M9.65 14.208a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDistancingAttention(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M.75 17v6m3.5-1l-2-2l2-2m19-1v6m-3.5-1l2-2l-2-2m2 2H2.25M12 9.275V5.887m6.269 9.613a.98.98 0 0 0 .881-1.413L13.045 1.65a1.164 1.164 0 0 0-2.09 0L4.85 14.087a.981.981 0 0 0 .881 1.413z"/><path d="M12 12.625a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDistancingCorrectFive(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18.5 19.375a3.375 3.375 0 1 0 0-6.75a3.375 3.375 0 0 0 0 6.75m4.75 3.75a6.027 6.027 0 0 0-9.5 0"/><path d="M21.875 16.111a6.762 6.762 0 0 1-6.443-1.511M5.5 19.375a3.375 3.375 0 1 0 0-6.75a3.375 3.375 0 0 0 0 6.75m4.749 3.75a6.026 6.026 0 0 0-9.5 0"/><path d="M8.875 16.111A6.762 6.762 0 0 1 2.432 14.6M.75 6.875h3.5m-1.5 2l-2-2l2-2m20.5 2h-3.5m1.5 2l2-2l-2-2m-9.25 6a5 5 0 1 0 0-10a5 5 0 0 0 0 10"/><path d="m10 6.375l1.083 1.083a.5.5 0 0 0 .76-.063L14 4.375"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDistancingCorrectFour(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M5.089 19.288a2.625 2.625 0 1 0 0-5.25a2.625 2.625 0 0 0 0 5.25m4.224 3.75a4.474 4.474 0 0 0-8.449 0m18.026-17a2.625 2.625 0 1 0 0-5.25a2.625 2.625 0 0 0 0 5.25m4.224 3.75a4.474 4.474 0 0 0-8.448 0"/><path d="M11.614 19.913a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75m3-2.25a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75m3-2.25a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75"/><path stroke-linecap="round" stroke-linejoin="round" d="M5.386 9.962a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9"/><path stroke-linecap="round" stroke-linejoin="round" d="m3.586 5.912l.974.974a.449.449 0 0 0 .684-.056l1.942-2.718"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDistancingCorrectOne(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18.5 19.375a3.375 3.375 0 1 0 0-6.75a3.375 3.375 0 0 0 0 6.75m4.75 3.75a6.027 6.027 0 0 0-9.5 0"/><path d="M21.875 16.111a6.762 6.762 0 0 1-6.443-1.511M5.5 19.375a3.375 3.375 0 1 0 0-6.75a3.375 3.375 0 0 0 0 6.75m4.749 3.75a6.026 6.026 0 0 0-9.5 0"/><path d="M8.875 16.111A6.762 6.762 0 0 1 2.432 14.6M5 7.875v2m14-2v2m-14-1h14m-10-6l2 2l4-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDistancingCorrectSix(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.974 19.375a2.625 2.625 0 1 0 0-5.25a2.625 2.625 0 0 0 0 5.25m4.226 3.75a4.473 4.473 0 0 0-8.449 0m18.275-3.75a2.625 2.625 0 1 0 0-5.25a2.625 2.625 0 0 0 0 5.25m4.224 3.75a4.473 4.473 0 0 0-8.449 0M.75 6.875h3.5m-1.5 2l-2-2l2-2m20.5 2h-3.5m1.5 2l2-2l-2-2m-9.25 6a5 5 0 1 0 0-10a5 5 0 0 0 0 10"/><path d="m10 6.375l1.083 1.083a.5.5 0 0 0 .76-.063L14 4.375"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDistancingCorrectThree(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M18.5 7.538a3.375 3.375 0 1 0 0-6.75a3.375 3.375 0 0 0 0 6.75m4.75 3.75a6.027 6.027 0 0 0-9.5 0"/><path stroke-linecap="round" stroke-linejoin="round" d="M21.875 4.275A6.78 6.78 0 0 1 20 4.538a6.73 6.73 0 0 1-4.568-1.78M5.5 19.288a3.375 3.375 0 1 0 0-6.75a3.375 3.375 0 0 0 0 6.75m4.749 3.75a6.026 6.026 0 0 0-9.5 0"/><path stroke-linecap="round" stroke-linejoin="round" d="M8.875 16.024a6.762 6.762 0 0 1-6.443-1.516m2.839-4.546a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9"/><path stroke-linecap="round" stroke-linejoin="round" d="m3.471 5.912l.975.974a.45.45 0 0 0 .684-.056l1.941-2.718"/><path d="M12.5 19.788a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75m3-2.125a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75m3-2.125a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDistancingCorrectTwo(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.974 19.375a2.625 2.625 0 1 0 0-5.25a2.625 2.625 0 0 0 0 5.25m4.226 3.75a4.473 4.473 0 0 0-8.449 0m18.275-3.75a2.625 2.625 0 1 0 0-5.25a2.625 2.625 0 0 0 0 5.25m4.224 3.75a4.473 4.473 0 0 0-8.449 0M5 8.875v2m14-2v2m-14-1h14m-10-7l2 2l4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDistancingDoNotCloseFour(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m13.524 22.009l.226 1.241h3.5l1-5.5H20v-3a4.5 4.5 0 0 0-2.484-4.024M8.5 9.25a2.75 2.75 0 1 0 0-5.5a2.75 2.75 0 0 0 0 5.5m9.668-3.418a2.735 2.735 0 0 0-4.345-1.5M.75 23.25L23.25.75M12.035 11.965A4.5 4.5 0 0 0 4 14.75v3h2.25m.229 4.013l.271 1.487h3.5l1-5.5H13v-2.507"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDistancingDoNotCloseOne(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.856 6.524a2.887 2.887 0 1 0 0-5.774a2.887 2.887 0 0 0 0 5.774M4.833 9.521a1.619 1.619 0 0 1-2.133.819l-.99-.44a1.618 1.618 0 0 1 1.318-2.956l.986.439a1.619 1.619 0 0 1 .819 2.138m2.306 12.11v-5.605m3.237-3.568v9.173a1.619 1.619 0 1 1-3.237 0"/><path d="M7.139 21.631a1.619 1.619 0 1 1-3.238 0v-8.173m-2.701-.25v2.421a1.617 1.617 0 0 0 1.618 1.618H3.9M15.516 6.524a2.887 2.887 0 1 0 0-5.774a2.887 2.887 0 0 0 0 5.774"/><path d="M21.084 10.968a5.4 5.4 0 0 0-5.312-4.445a12.107 12.107 0 0 0-1.835.053m-.863 5.343a5.4 5.4 0 0 0-5.4-5.4a7.485 7.485 0 0 0-2.844.3a1.74 1.74 0 0 0-.82.564m14.24 15.842a5 5 0 1 0 0-10a5 5 0 0 0 0 10m3.535-8.535l-7.07 7.07"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDistancingDoNotCloseThree(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.218 10.25a4.75 4.75 0 1 0 0-9.5a4.75 4.75 0 0 0 0 9.5M17.021.946a4.75 4.75 0 1 1 .642 9.251M6.471 17.724a2.125 2.125 0 0 1-3 0l-2.1-2.1a2.125 2.125 0 0 1 3.005-3l2.1 2.1a2.125 2.125 0 0 1-.005 3m-3.498 3.034v2.492"/><path d="M4.251 12.508A6.235 6.235 0 0 1 9.06 10.25h3.222a6.3 6.3 0 0 1 2.815.656m2.913 12.319a5 5 0 1 0 0-10a5 5 0 0 0 0 10m3.535-8.535l-7.07 7.07"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDistancingDoNotCloseTwo(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.083 7.936a5.007 5.007 0 0 1 3.551 2.439M6.225 6.525a2.825 2.825 0 1 0 0-5.65a2.825 2.825 0 0 0 0 5.65m4.828 8.1a1.5 1.5 0 1 1 0-3h2.5m-7.25 10.25v-5.194m2.997.436v4.508a1.5 1.5 0 0 1-3 0M10.809.996a2.825 2.825 0 1 1-.05 5.392"/><path d="M.8 12.875a5 5 0 0 1 5-5h1a5 5 0 0 1 4.843 3.75"/><path d="M6.3 21.625a1.5 1.5 0 0 1-3 0v-6.812h-1a1.5 1.5 0 0 1-1.5-1.5v-.438m17.396 10.107a5 5 0 1 0 0-10a5 5 0 0 0 0 10m3.535-8.534l-7.07 7.07"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDistancingDoNotTouch(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m20.847 10.46l-1.079-5.394a1.613 1.613 0 0 0-3.185.14l-.362 3.255l-7.228-7.229A1.572 1.572 0 0 0 6.77 3.456L5.658 2.345a1.572 1.572 0 1 0-2.224 2.223L4.546 5.68A1.573 1.573 0 0 0 2.322 7.9l1.112 1.116A1.573 1.573 0 0 0 1.21 11.24l2.078 2.126m5.88-7.511L6.767 3.454m-.686 3.761L4.544 5.678"/><path d="M12.382 12.47C8.271 8.726 6.458 9.717 5.9 10.28l-3.715 4.392a3.752 3.752 0 0 0-.9 2.111m9.039 1.208l-1.246-.013l2.261-2.581m-6.918 6.759l.714-.892c.78.169 1.59.14 2.355-.083l3.259-.715m7.507 2.762a4.994 4.994 0 1 0 0-9.988a4.994 4.994 0 0 0 0 9.988m3.53-8.525l-7.061 7.062"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDistancingMan(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 5.8A2.4 2.4 0 1 0 12 1a2.4 2.4 0 0 0 0 4.8m4.2 5.4a4.2 4.2 0 1 0-8.4 0V13h1.8l.6 6h3.6l.6-6h1.8zM1 6v4m4-2H1m22-2v4m-4-2h4"/><path d="M5 15.914c-2.443.734-4 1.844-4 3.086c0 2.209 4.925 4 11 4s11-1.791 11-4c0-1.242-1.557-2.352-4-3.086"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDistancingNotAllowedSpaceMan(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 6.75a3 3 0 1 0 0-6a3 3 0 0 0 0 6m5.25 6.75a5.25 5.25 0 1 0-10.5 0v2.25H9l.75 7.5h4.5l.75-7.5h2.25zM1 18.991l4 4m0-4l-4 4m18-4l4 4m0-4l-4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDistancingNotAllowedSpaceWoman(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m1 19.05l4 4m0-4l-4 4m18-4l4 4m0-4l-4 4m-7-16.3a3 3 0 1 0 0-6a3 3 0 0 0 0 6M9.114 9.114A5.246 5.246 0 0 0 6.75 13.5v2.25H9l.75 7.5h4.5l.75-7.5h2.25V13.5a5.246 5.246 0 0 0-2.364-4.386L12 13.5z"/><path d="M9 3.75S9 7.5 6.75 7.5M15 3.75s0 3.75 2.25 3.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDistancingOne(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M18.5 17.375a3.375 3.375 0 1 0 0-6.75a3.375 3.375 0 0 0 0 6.75m4.75 3.75a6.027 6.027 0 0 0-9.5 0"/><path stroke-linecap="round" stroke-linejoin="round" d="M21.875 14.111a6.762 6.762 0 0 1-6.443-1.511M5.5 17.375a3.375 3.375 0 1 0 0-6.75a3.375 3.375 0 0 0 0 6.75m4.749 3.75a6.026 6.026 0 0 0-9.5 0"/><path stroke-linecap="round" stroke-linejoin="round" d="M8.875 14.111A6.762 6.762 0 0 1 2.432 12.6"/><path d="M5.875 3.75a.375.375 0 1 1 0-.75m0 .75a.375.375 0 1 0 0-.75M12 3.75A.375.375 0 0 1 12 3m0 .75A.375.375 0 0 0 12 3m6.125.75a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDistancingProtectShieldOne(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 12.85a1.95 1.95 0 1 0 0-3.9a1.95 1.95 0 0 0 0 3.9m1.3 10.4l.65-3.9h1.3V17.4a3.25 3.25 0 0 0-6.5 0v1.95h1.3l.65 3.9zM20 12.85a1.95 1.95 0 1 0 0-3.9a1.95 1.95 0 0 0 0 3.9m-1.3 10.4l-.65-3.9h-1.3V17.4a3.25 3.25 0 0 1 6.5 0v1.95h-1.3l-.65 3.9zM4 4.75h4.65M4 5.75v-2m16 1h-4.651m4.651 1v-2m-8-3a6.82 6.82 0 0 1-3.045.975a.5.5 0 0 0-.455.5v1.22A5.776 5.776 0 0 0 12 8.75a5.776 5.776 0 0 0 3.5-5.308v-1.22a.5.5 0 0 0-.455-.5A6.82 6.82 0 0 1 12 .75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDistancingProtectShieldTwo(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M.75 9.25v4m22.5-4v4m-22.5-2h5.5m0-4.013v4.49a7.022 7.022 0 0 0 4.5 6.554l.614.236c.409.157.861.157 1.27 0l.614-.236a7.022 7.022 0 0 0 4.5-6.554v-4.49a.877.877 0 0 0-.512-.8A12.8 12.8 0 0 0 12 5.365a12.8 12.8 0 0 0-5.238 1.068a.877.877 0 0 0-.512.804m11.5 4.013h5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDistancingTwo(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M4.974 17.375a2.625 2.625 0 1 0 0-5.25a2.625 2.625 0 0 0 0 5.25m4.226 3.75a4.473 4.473 0 0 0-8.449 0m18.275-3.75a2.625 2.625 0 1 0 0-5.25a2.625 2.625 0 0 0 0 5.25m4.224 3.75a4.473 4.473 0 0 0-8.449 0"/><path d="M6 3.75A.375.375 0 1 1 6 3m0 .75A.375.375 0 1 0 6 3m6 .75A.375.375 0 0 1 12 3m0 .75A.375.375 0 0 0 12 3m6 .75A.375.375 0 0 1 18 3m0 .75A.375.375 0 0 0 18 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDistancingVirus(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 18.675a1.95 1.95 0 1 0 0-3.9a1.95 1.95 0 0 0 0 3.9m3.25 4.55a3.25 3.25 0 1 0-6.5 0M20 18.675a1.95 1.95 0 1 0 0-3.9a1.95 1.95 0 0 0 0 3.9m-3.25 4.55a3.25 3.25 0 1 1 6.5 0M12 9.025a3 3 0 1 0 0-6a3 3 0 0 0 0 6m-.5-8.25h1m-.5 0v2.25m3.359-1.066l.707.707m-.354-.353l-1.591 1.591m3.129 1.621v1m0-.5H15m1.066 3.359l-.707.707m.353-.354l-1.591-1.591M12.5 11.275h-1m.5 0v-2.25m-3.359 1.066l-.707-.707m.354.353l1.591-1.591M6.75 6.525v-1m0 .5H9M7.934 2.666l.707-.707m-.353.354l1.591 1.591m-.629 9.871l-1 1l1 1m6.5-1h-7.5m6.5-1l1 1l-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDistancingWoman(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 6.009a2.567 2.567 0 1 0 0-5.134a2.567 2.567 0 0 0 0 5.134m1.925 14.116l.642-6.417h1.925v-1.925a4.488 4.488 0 0 0-2.023-3.752L12 11.783L9.531 8.031a4.488 4.488 0 0 0-2.023 3.752v1.925h1.925l.642 6.417z"/><path d="M9.433 3.442s0 3.208-1.925 3.208m7.059-3.208s0 3.208 1.925 3.208M1 6.125v4m4-2H1m22-2v4m-4-2h4M5 16.039c-2.443.734-4 1.844-4 3.086c0 2.209 4.925 4 11 4s11-1.791 11-4c0-1.242-1.557-2.352-4-3.086"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymptomsColdFever(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.966 20.602a2.714 2.714 0 1 0 0-5.428a2.714 2.714 0 0 0 0 5.428m-.452-7.464h.905m-.453 0v2.036m3.039-.965l.64.64m-.32-.32l-1.439 1.44m2.83 1.467v.904m0-.452h-2.035m.964 3.039l-.64.64m.32-.32l-1.439-1.44m-1.467 2.831h-.905m.452 0v-2.036m-3.038.965l-.64-.64m.32.32l1.439-1.44m-2.831-1.467v-.904m0 .452h2.036m-.964-3.039l.64-.64m-.32.32l1.439 1.44M11.411 6.75a3 3 0 1 0 0-6a3 3 0 0 0 0 6m4.555 4.138a5.251 5.251 0 0 0-9.8 2.612v2.25h2.25l.75 7.5h2.805M2.966 3.888L1.524 5.571a1 1 0 0 0 0 1.3L2.409 7.9a1 1 0 0 1 0 1.3l-.885 1.031a1 1 0 0 0 0 1.3l.885 1.032a1 1 0 0 1 0 1.3L1.524 14.9a1 1 0 0 0 0 1.3l1.442 1.682m16.558-8.676a1 1 0 0 1 0-1.3l.885-1.032a1 1 0 0 0 0-1.3l-1.443-1.686"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymptomsFever(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.01 16.607a3.495 3.495 0 1 0 0-6.99a3.495 3.495 0 0 0 0 6.99m5.242 6.616a5.243 5.243 0 0 0-10.485 0m20.968-7.821V3.751a2.996 2.996 0 1 0-5.991 0v11.651a4.493 4.493 0 1 0 5.991 0M18.74 5.249v11.983"/><path d="M18.74 20.227a1.498 1.498 0 1 0 0-2.995a1.498 1.498 0 0 0 0 2.995M2.014 7.745a2.47 2.47 0 0 1 0-3.495a2.472 2.472 0 0 0 0-3.495m3.721 6.99a2.47 2.47 0 0 1 0-3.495a2.472 2.472 0 0 0 0-3.495m3.995 6.99a2.47 2.47 0 0 1 0-3.495a2.473 2.473 0 0 0 0-3.495"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymptomsNausea(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m17.222 3.832l2.927-1.947v1.947l2.927-1.947m-5.719 15.22l3.231 1.383l-1.628 1.066l3.231 1.383m-5.105-9.755l3.354-1.051L19.896 12l3.354-1.052M4.834 12A4.087 4.087 0 0 1 .769 7.515A4.2 4.2 0 0 1 5.01 3.832h.845a1.021 1.021 0 0 0 1.021-1.021V.769M4.834 12a4.087 4.087 0 0 0-4.065 4.485a4.2 4.2 0 0 0 4.241 3.683h.845a1.021 1.021 0 0 1 1.021 1.021v2.042m5.105 0v-3.063A4.083 4.083 0 0 0 7.9 16.084h2.039a4.084 4.084 0 0 0 0-8.168H7.9a4.083 4.083 0 0 0 4.084-4.084V.769M4.834 12h2.042m1.021-4.084H6.876m1.021 8.168H6.876"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymptomsVirusDiarrheaOne(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.886 10.392a3.506 3.506 0 1 0 0-7.012a3.506 3.506 0 0 0 0 7.012M6.302.75h1.169m-.585 0v2.63m3.926-1.246l.827.827m-.414-.414l-1.859 1.86m3.657 1.895v1.169m0-.585h-2.63m1.246 3.926l-.827.827m.413-.414L9.366 9.366m-1.895 3.657H6.302m.584 0v-2.63m-3.925 1.246l-.827-.827m.413.413l1.86-1.859M.75 7.471V6.302m0 .584h2.63M2.134 2.961l.827-.827m-.414.413l1.86 1.86m14.207 10.729h3.477a1.159 1.159 0 0 0 1.159-1.159V7.023a1.159 1.159 0 0 0-1.159-1.159h-2.318a1.159 1.159 0 0 0-1.159 1.159z"/><path d="M20.509 17.335a2.32 2.32 0 0 0 1.582-2.2H9.341a6.374 6.374 0 0 0 3.477 5.68l-.579 2.434h7.534v-4.8a1.152 1.152 0 0 1 .736-1.114"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymptomsVirusDiarrheaTwo(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m14.513 6.193l2.222-2.63l.531 1.832l2.222-2.63m-2.405 6.151l3.07-1.557l-.223 1.894L23 7.695M9.12 6.001a2.623 2.623 0 1 0 0-5.246a2.623 2.623 0 0 0 0 5.246m-2.25 7.498A1.499 1.499 0 0 0 5.373 12H2.374a1.499 1.499 0 0 0-1.499 1.499v9.743m8.994-.003a4.497 4.497 0 0 0 4.496-4.497M.875 16.494H6.81"/><path d="M10.01 15.744h5.854a1.499 1.499 0 0 1 1.341.83l3.336 6.671"/><path d="m17.18 23.229l-2.242-4.487H8.37a1.5 1.5 0 0 1-1.493-1.635V9.313a1.913 1.913 0 0 1 3.302-1.37l5.247 5.246a1.499 1.499 0 1 1-2.12 2.12l-2.984-2.983"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymptomsVirusHeadacheOne(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.326 16.343a3.454 3.454 0 1 0 0-6.908a3.454 3.454 0 0 0 0 6.908m5.181 6.907a5.18 5.18 0 0 0-10.361 0m17.021-4.969a2.678 2.678 0 1 0 0-5.356a2.678 2.678 0 0 0 0 5.356m-.446-7.366h.893m-.447 0v2.009m2.999-.951l.631.631m-.315-.316l-1.421 1.421m2.793 1.447v.893m0-.446h-2.008m.951 2.998l-.631.632m.316-.316l-1.421-1.421m-1.447 2.794h-.893m.446 0v-2.009m-2.998.952l-.632-.632m.316.316l1.42-1.421m-2.793-1.447v-.893m0 .447h2.009m-.952-2.999l.632-.631m-.316.315l1.42 1.421M2.83 7.462L1.56 4.311l1.842.381l-1.27-3.151m4.259 4.934l.056-3.396l1.549 1.068L8.053.75m2.36 6.712L11.655 4.3l1.077 1.542l1.242-3.162"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymptomsVirusHeadacheTwo(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15.326 23.25v-3.99h2.763a1.33 1.33 0 0 0 1.33-1.33v-2.867h1.536a1.022 1.022 0 0 0 .933-1.44c-1.005-2.237-2.473-4.987-2.473-4.987V6.457a2.04 2.04 0 0 0-.182-.849a7.3 7.3 0 0 0-4.976-4.527C8.279-.487 2.067 3.7 2.022 9.876a9.185 9.185 0 0 0 3.07 6.934v6.44"/><path d="M10.984 12.475a2.776 2.776 0 1 0 0-5.552a2.776 2.776 0 0 0 0 5.552m-.462-7.634h.925m-.463 0v2.082m3.108-.986l.654.654m-.327-.327l-1.472 1.472m2.895 1.5v.926m0-.463H13.76m.986 3.108l-.654.654m.327-.327l-1.472-1.472m-1.5 2.895h-.925m.462 0v-2.082m-3.108.986l-.654-.654m.327.327l1.472-1.472m-2.895-1.5v-.926m0 .463h2.082m-.986-3.108l.654-.654m-.327.327l1.472 1.472"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymptomsVirusLossSmellOne(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.886 10.392a3.506 3.506 0 1 0 0-7.012a3.506 3.506 0 0 0 0 7.012M6.302.75h1.169m-.585 0v2.63m3.926-1.246l.827.827m-.414-.414l-1.859 1.86m3.657 1.895v1.169m0-.585h-2.63m1.246 3.926l-.827.827m.413-.414L9.366 9.366m-1.895 3.657H6.302m.584 0v-2.63m-3.925 1.246l-.827-.827m.413.413l1.86-1.859M.75 7.471V6.302m0 .584h2.63M2.134 2.961l.827-.827m-.414.413l1.86 1.86M22.468.75s-2.82 5.83-5.475 7.821c-3.128 2.346-1.564 5.474.782 5.474h5.475"/><path d="M19.578 10.977c.767-1.022 1.534-1.022 3.068-1.022m-8.166 7.67a6.498 6.498 0 0 0 1.007-1.534m-5.114 4.091a8.386 8.386 0 0 0 2.046-.842m7.765 1.865a3.98 3.98 0 0 0 .416 2.045m1.023-6.136c-.341.681-.682 1.363-.947 2.045m-5.189 3.068c-.549.56-1.269.92-2.046 1.023m4.091-4.09a8.37 8.37 0 0 1-.758 1.533"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymptomsVirusLossSmellTwo(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15.409 17.5a6.498 6.498 0 0 0 1.007-1.534M11.3 20.058a8.379 8.379 0 0 0 2.046-.842m7.767 1.864a3.98 3.98 0 0 0 .417 2.046m1.022-6.137c-.341.682-.682 1.364-.947 2.046M16.416 22.1c-.549.56-1.269.92-2.046 1.023m4.091-4.088a8.37 8.37 0 0 1-.758 1.534M22.468.874S19.648 6.7 16.993 8.7c-3.128 2.346-1.564 5.475.782 5.475h5.475"/><path d="M19.578 11.1c.767-1.022 1.534-1.022 3.068-1.022M6.375 12.124a5.625 5.625 0 1 0 0-11.25a5.625 5.625 0 0 0 0 11.25m-3.977-1.648l7.954-7.954"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymptomsVirusLungDamage(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18.058 9.159a2.967 2.967 0 1 0 0-5.934a2.967 2.967 0 0 0 0 5.934M17.563 1h.989m-.494 0v2.225m3.322-1.054l.699.699m-.35-.349l-1.573 1.573m3.094 1.604v.989m0-.495h-2.225m1.054 3.322l-.699.699m.349-.349L20.156 8.29m-1.604 3.095h-.989m.495 0V9.159m-3.322 1.054l-.699-.699m.349.35L15.96 8.29m-3.095-1.603v-.989m0 .494h2.226M14.037 2.87l.699-.699m-.35.35l1.574 1.573m-6.062.609v7.929M6.849 9.078a1.219 1.219 0 0 0-2.122-.82l-.805.884A12.2 12.2 0 0 0 .75 17.35v3.21a2.44 2.44 0 0 0 3.35 2.266l1.986-.8a1.22 1.22 0 0 0 .767-1.133zm-2.44 7.823l5.489-4.269m8.834 1.961c.21.903.315 1.828.315 2.755v3.212a2.44 2.44 0 0 1-3.347 2.266l-1.986-.8a1.22 1.22 0 0 1-.767-1.133v-9.267m2.44 5.275l-5.489-4.269"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymptomsVirusStomach(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.906 13.498a2.776 2.776 0 1 0 0-5.552a2.776 2.776 0 0 0 0 5.552m-.462-7.634h.925m-.463 0v2.082m3.108-.987l.655.655m-.328-.328l-1.472 1.473m2.895 1.5v.925m0-.462h-2.082m.987 3.108l-.655.654m.327-.327l-1.472-1.472m-1.5 2.895h-.925m.462 0v-2.082m-3.108.986l-.654-.654m.327.327l1.472-1.472m-2.895-1.501v-.925m0 .463h2.082m-.986-3.108l.654-.655m-.327.327l1.472 1.473"/><path d="M8.083 23.25a11.1 11.1 0 0 1 .826-3.731c4.13.685 11.261-1.615 13.4-5.892C25.17 7.9 20.184.781 13.6 1.878a10.621 10.621 0 0 0-5.3 2.551C7.051 3.248 7 2.8 5.98.75M.867.75a20.462 20.462 0 0 0 4.356 7.8C3.827 11.48 3.6 14.7 4.7 16.828a14.343 14.343 0 0 0-1.73 6.422"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransmissionVirusAirplaneFlight(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.231 20.429a3.429 3.429 0 1 0 0-6.858a3.429 3.429 0 0 0 0 6.858M16.66 11h1.142m-.571 0v2.571m3.839-1.218l.808.808m-.404-.404l-1.819 1.819m3.576 1.853v1.142m0-.571H20.66m1.218 3.839l-.808.808m.404-.404l-1.819-1.819M17.802 23H16.66m.571 0v-2.571m-3.839 1.218l-.808-.808m.404.404l1.819-1.819m-3.576-1.853v-1.142m0 .571h2.571m-1.218-3.839l.808-.808m-.404.404l1.819 1.819m-.234-5.282L15 8.206l3.2-1.076a2.413 2.413 0 0 0-1.54-4.573L5.228 6.4L3.7 4.878l-2.271.763A.964.964 0 0 0 1.015 7.2l3.073 3.44a1.931 1.931 0 0 0 2.054.544l3.366-1.134l-.682 4.864"/><path d="M13.459 3.633L9.883 1.17a.969.969 0 0 0-.855-.12l-1.721.579a.964.964 0 0 0-.373 1.6L8.886 5.17"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransmissionVirusBriefcase(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.369 11.107a2.857 2.857 0 1 0 0-5.714a2.857 2.857 0 0 0 0 5.714m-.476-7.857h.952m-.476 0v2.143m3.199-1.015l.673.673m-.336-.337L14.389 6.23m2.98 1.544v.952m0-.476h-2.143m1.015 3.199l-.673.673m.337-.336l-1.516-1.516m-1.544 2.98h-.952m.476 0v-2.143M9.17 12.122l-.673-.673m.337.337l1.515-1.516m-2.98-1.544v-.952m0 .476h2.143M8.497 5.051l.673-.673m-.336.336l1.515 1.516M6.75 2.25V.75m10.5 1.5V.75m-9.75 18a1.5 1.5 0 0 0 1.5 1.5h6a1.5 1.5 0 0 0 1.5-1.5"/><path d="M22.25 15.75H1.75a1 1 0 0 0-1 1v5.5a1 1 0 0 0 1 1h20.5a1 1 0 0 0 1-1v-5.5a1 1 0 0 0-1-1m-6.5-13.5h3.75a1.5 1.5 0 0 1 1.5 1.5v12m-18 0v-12a1.5 1.5 0 0 1 1.5-1.5h3.94"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransmissionVirusCough(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.069 23v-2.482h1.72a.828.828 0 0 0 .828-.828v-1.784h.956a.634.634 0 0 0 .58-.9c-.625-1.392-1.538-3.1-1.538-3.1v-1.355c0-.182-.038-.362-.113-.528A4.54 4.54 0 0 0 8.4 9.206a5.966 5.966 0 0 0-7.609 5.473A5.72 5.72 0 0 0 2.7 18.993V23M16.959 8.464a2.714 2.714 0 1 0 0-5.428a2.714 2.714 0 0 0 0 5.428M16.507 1h.904m-.452 0v2.036m3.039-.965l.64.64m-.32-.32l-1.44 1.44m2.831 1.467v.904m0-.452h-2.036m.965 3.039l-.64.64m.32-.32l-1.44-1.44M17.411 10.5h-.904m.452 0V8.464m-3.039.965l-.64-.64m.32.32l1.44-1.44m-2.831-1.467v-.904m0 .452h2.036m-.965-3.039l.64-.64m-.32.32l1.44 1.44M16.209 15l3.5-2v2l3.5-2m-7 6l3.923.885l-.846 1.23l3.923.885"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransmissionVirusCrowd(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.107 18.251a2.063 2.063 0 1 0 0-4.126a2.063 2.063 0 0 0 0 4.126m3.656 4.874a3.656 3.656 0 0 0-7.312 0m10.902-7.481a2.017 2.017 0 1 0 0-4.034a2.017 2.017 0 0 0 0 4.034m3.576 4.767A3.575 3.575 0 0 0 17 17.719M4.647 15.644a2.017 2.017 0 1 0 0-4.034a2.017 2.017 0 0 0 0 4.034m-3.576 4.767A3.575 3.575 0 0 1 7 17.719m5-8.594a3 3 0 1 0 0-6a3 3 0 0 0 0 6m-.5-8.25h1m-.5 0v2.25m3.359-1.066l.707.707m-.354-.353l-1.591 1.591m3.129 1.621v1m0-.5H15m1.066 3.359l-.707.707m.353-.354l-1.591-1.591M12.5 11.375h-1m.5 0v-2.25m-3.359 1.066l-.707-.707m.354.353l1.591-1.591M6.75 6.625v-1m0 .5H9M7.934 2.766l.707-.707m-.353.354l1.591 1.591"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransmissionVirusExpand(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 14.857a2.857 2.857 0 1 0 0-5.714a2.857 2.857 0 0 0 0 5.714M11.524 7h.952M12 7v2.143m3.199-1.015l.673.673m-.336-.337L14.02 9.98M17 11.524v.952M17 12h-2.143m1.015 3.199l-.673.673m.337-.336L14.02 14.02M12.476 17h-.952M12 17v-2.143m-3.199 1.015l-.673-.673m.336.337L9.98 14.02M7 12.476v-.952M7 12h2.143M8.128 8.801l.673-.673m-.337.336L9.98 9.98M1 1l5 5M1 5V1h4m18 0l-5 5m1-5h4v4m0 18l-5-5m5 1v4h-4M1 23l5-5m-1 5H1v-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransmissionVirusFloatWind(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.926 19.117a4.857 4.857 0 1 0 0-9.714a4.857 4.857 0 0 0 0 9.714m-.81-13.357h1.619m-.809 0v3.642m7.126-8.162l1.374.857m-.687-.428L16.811 4.76m3.618 2.851l1.145 1.145m-.572-.572l-2.576 2.576m-.917 8.938l-1.145 1.144m.572-.572l-2.576-2.576m-2.625 5.066h-1.619m.81 0v-3.643m-5.438 1.725l-1.145-1.144m.572.572l2.576-2.576m-5.065-2.625V13.45m0 .81h3.643M4.343 8.822l1.145-1.145m-.573.572l2.576 2.576"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransmissionVirusHandshake(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18.25 8.685a2.857 2.857 0 1 0 0-5.714a2.857 2.857 0 0 0 0 5.714M17.774.828h.952m-.476 0v2.143m3.199-1.015l.673.673m-.336-.337L20.27 3.807m2.98 1.545v.952m0-.476h-2.143m1.015 3.199l-.673.673m.337-.337L20.27 7.848m-1.544 2.98h-.952m.476 0V8.685M15.051 9.7l-.673-.673m.336.336l1.516-1.515m-2.98-1.544v-.952m0 .476h2.143m-1.015-3.199l.673-.673m-.337.336l1.516 1.515m.003 16.153l-4.269.854m.422-4.843l-2.014.851a1.27 1.27 0 0 1-1-.014a1.323 1.323 0 0 1-.7-.741a1.36 1.36 0 0 1 .63-1.693l2.014-1.006a1.754 1.754 0 0 1 1.422-.084l4.138 1.631"/><path d="M4.143 20.848h1.425l2.719 2.074A.739.739 0 0 0 9.312 23l3.8-3.138a.751.751 0 0 0 .115-1.038L11.133 16.5m-2.048-1.984l-.22-.181a1.593 1.593 0 0 0-1.538-.2l-3.185 1.359"/><path d="M.75 21.517h2.372a1.018 1.018 0 0 0 1.07-.956v-4.778a1.018 1.018 0 0 0-1.07-.955H.75m18.925 6.689H17.3a1.018 1.018 0 0 1-1.07-.956v-4.778a1.018 1.018 0 0 1 1.07-.955h2.372m-9.396-9h-9m3-3l-3 3l3 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransmissionVirusHumanInfected(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.25 19.491a3.429 3.429 0 1 0 0-6.858a3.429 3.429 0 0 0 0 6.858m-.571-9.428h1.142m-.571 0v2.571m3.839-1.218l.808.808m-.404-.404l-1.819 1.818m3.576 1.853v1.143m0-.572h-2.571m1.218 3.839l-.808.808m.404-.404l-1.819-1.818m-1.853 3.575l-1.142.001m.571-.001v-2.571m-3.839 1.218l-.808-.808m.404.404l1.819-1.818m-3.576-1.853v-1.143m0 .572h2.571m-1.218-3.839l.808-.808m-.404.404l1.819 1.818M4.5 6.188a2.625 2.625 0 1 0 0-5.25a2.625 2.625 0 0 0 0 5.25M2.25 20.812a2.25 2.25 0 0 0 4.5 0v-3.75a1.5 1.5 0 0 0 1.5-1.5v-4.5a3 3 0 0 0-3-3h-1.5a3 3 0 0 0-3 3v4.5a1.5 1.5 0 0 0 1.5 1.5zM16.5 1.313l-2.25 2.25l2.25 2.25m-2.25-2.25h6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransmissionVirusHumanTransmitOne(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.25 20.55a3.429 3.429 0 1 0 0-6.858a3.429 3.429 0 0 0 0 6.858m-.571-9.429h1.142m-.571 0v2.572m3.839-1.218l.808.808m-.404-.404l-1.819 1.818m3.576 1.853v1.143m0-.572h-2.571m1.218 3.839l-.808.808m.404-.404l-1.819-1.818m-1.853 3.575h-1.142m.571 0V20.55m-3.839 1.218l-.808-.808m.404.404l1.819-1.818m-3.576-1.853V16.55m0 .571h2.571m-1.218-3.838l.808-.808m-.404.404l1.819 1.818M6.167 10.5a3.311 3.311 0 1 0 0-6.622a3.311 3.311 0 0 0 0 6.622M.75 17.121a5.42 5.42 0 0 1 7.55-4.982m5.707-4.26l.243-.97a4 4 0 0 1 3.88-3.03h4.877"/><path d="m20.007 6.879l3-3l-3-3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransmissionVirusHumanTransmitTwo(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.25 19.491a3.429 3.429 0 1 0 0-6.858a3.429 3.429 0 0 0 0 6.858m-.571-9.428h1.142m-.571 0v2.571m3.839-1.218l.808.808m-.404-.404l-1.819 1.818m3.576 1.853v1.143m0-.572h-2.571m1.218 3.839l-.808.808m.404-.404l-1.819-1.818m-1.853 3.575l-1.142.001m.571-.001v-2.571m-3.839 1.218l-.808-.808m.404.404l1.819-1.818m-3.576-1.853v-1.143m0 .572h2.571m-1.218-3.839l.808-.808m-.404.404l1.819 1.818M4.5 6.188a2.625 2.625 0 1 0 0-5.25a2.625 2.625 0 0 0 0 5.25M2.25 20.812a2.25 2.25 0 0 0 4.5 0v-3.75a1.5 1.5 0 0 0 1.5-1.5v-4.5a3 3 0 0 0-3-3h-1.5a3 3 0 0 0-3 3v4.5a1.5 1.5 0 0 0 1.5 1.5zm16-19.499l2.25 2.25l-2.25 2.25m2.25-2.25l-6-.001"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransmissionVirusIncreaseRate(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 19.571a4.571 4.571 0 1 0 0-9.142a4.571 4.571 0 0 0 0 9.142M8.238 7h1.524M9 7v3.429m5.118-1.625l1.078 1.078m-.539-.539l-2.425 2.425M17 14.238v1.524M17 15h-3.429m1.625 5.118l-1.078 1.078m.539-.539l-2.425-2.425M9.762 23H8.238M9 23v-3.429m-5.118 1.625l-1.078-1.078m.539.539l2.425-2.425M1 15.762v-1.524M1 15h3.429M2.804 9.882l1.078-1.078m-.539.539l2.425 2.425M17 4l3-3l3 3m-3-3v7m0 4v3m0 4v3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransmissionVirusInhale(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.028 23.25v-2.482h1.72a.827.827 0 0 0 .827-.828v-1.784h.956a.634.634 0 0 0 .58-.9c-.625-1.392-1.538-3.1-1.538-3.1V12.8c0-.182-.038-.362-.113-.528a4.541 4.541 0 0 0-3.1-2.817a5.966 5.966 0 0 0-7.61 5.474a5.719 5.719 0 0 0 1.91 4.314v4.007M18 9a3 3 0 1 0 0-6a3 3 0 0 0 0 6M17.5.75h1m-.5 0V3m3.359-1.066l.707.707m-.354-.353l-1.591 1.591M23.25 5.5v1m0-.5H21m1.066 3.359l-.707.707m.353-.354l-1.591-1.591M18.5 11.25h-1m.5 0V9m-3.359 1.066l-.707-.707m.354.353l1.591-1.591M12.75 6.5v-1m0 .5H15m-1.066-3.359l.707-.707m-.353.354l1.591 1.591m5.875 10.371v4a3 3 0 0 1-3 3h-4"/><path d="m16.754 23.25l-2-2l2-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransmissionVirusMobileApplication(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.75 9.893a3.143 3.143 0 1 0 0-6.286a3.143 3.143 0 0 0 0 6.286m-.524-8.643h1.048m-.524 0v2.357m3.519-1.116l.74.74m-.37-.37l-1.667 1.667m3.278 1.698v1.048m0-.524h-2.357m1.116 3.519l-.74.74m.37-.37l-1.667-1.667m-1.698 3.278h-1.048m.524 0V9.893m-3.519 1.116l-.74-.74m.37.37l1.667-1.667M12.25 7.274V6.226m0 .524h2.357m-1.116-3.519l.74-.74m-.37.37l1.667 1.667M7.5 15l1.673-1.255a1.568 1.568 0 0 1 2.247.385v0a1.569 1.569 0 0 1 0 1.74l-1.645 2.467A4.5 4.5 0 0 1 7.7 20.019L5.25 21m3-11.25l-3.595 1.541a4.5 4.5 0 0 0-2.364 2.363L1.114 16.4a4.5 4.5 0 0 0-.364 1.774v5.076"/><path d="M23.25 14.25V21A2.25 2.25 0 0 1 21 23.25H10.5A2.25 2.25 0 0 1 8.25 21v-1.245m0-5.318V3A2.25 2.25 0 0 1 10.5.75h1.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransmissionVirusRatMouse(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 20.429a3.429 3.429 0 1 0 0-6.858a3.429 3.429 0 0 0 0 6.858M6.429 11h1.142M7 11v2.571m3.839-1.218l.808.808m-.404-.404l-1.819 1.819M13 16.429v1.142M13 17h-2.571m1.218 3.839l-.808.808m.404-.404l-1.819-1.819M7.571 23H6.429M7 23v-2.571m-3.839 1.218l-.808-.808m.404.404l1.819-1.819M1 17.571v-1.142M1 17h2.571m-1.218-3.839l.808-.808m-.404.404l1.819 1.819M15.75 13h6a1.5 1.5 0 0 0 0-3c-1 0-.131.495-4.3-2.844A4.1 4.1 0 0 0 13.875 1A4.183 4.183 0 0 0 9.75 5.125a4.116 4.116 0 0 0 1.5 3.179M8.336 1.127A7.126 7.126 0 0 0 7 1a6 6 0 0 0-5.2 9"/><path d="M16.125 9.25a.375.375 0 0 1 .375.375m-.75 0a.375.375 0 0 1 .375-.375m0 .75a.375.375 0 0 1-.375-.375m.75 0a.375.375 0 0 1-.375.375"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransmissionVirusTouchFinger(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 10.179a3.429 3.429 0 1 0 0-6.858a3.429 3.429 0 0 0 0 6.858M6.429.75h1.142M7 .75v2.571m3.839-1.218l.808.808m-.404-.404L9.424 4.326M13 6.179v1.142m0-.571h-2.571m1.218 3.839l-.808.808m.404-.404l-1.819-1.82M7.071 12.75H5.929m1.071 0v-2.571m-3.839 1.218l-.808-.808m.404.404l1.819-1.82M1 7.321V6.179m0 .571h2.571M2.353 2.911l.808-.808m-.404.404l1.819 1.819M12.75 23.25l-2.764-3.11a1.557 1.557 0 1 1 2.327-2.069l1.937 2.179v-9a1.5 1.5 0 1 1 3 0V18h1.5a4.5 4.5 0 0 1 4.5 4.5v.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransmissionVirusTouchHandOne(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17 20.429a3.429 3.429 0 1 0 0-6.858a3.429 3.429 0 0 0 0 6.858M16.429 11h1.142M17 11v2.571m3.839-1.218l.808.808m-.404-.404l-1.819 1.819M23 16.429v1.142M23 17h-2.571m1.218 3.839l-.808.808m.404-.404l-1.819-1.819M17.571 23h-1.142M17 23v-2.571m-3.839 1.218l-.808-.808m.404.404l1.819-1.819M11 17.571v-1.142M11 17h2.571m-1.218-3.839l.808-.808m-.404.404l1.819 1.819M4 11.252V3.786c0-.396.158-.776.44-1.056a1.503 1.503 0 0 1 2.12 0c.282.28.44.66.44 1.056"/><path d="M4 9.012a1.49 1.49 0 0 0-.44-1.056a1.503 1.503 0 0 0-2.12 0c-.282.28-.44.66-.44 1.056v5.815a4.463 4.463 0 0 0 2 3.727l1.828 1.214A1.495 1.495 0 0 1 5.5 21.01v2.19M7 7.52V2.292c0-.396.158-.776.44-1.056a1.503 1.503 0 0 1 2.12 0c.282.28.44.66.44 1.056V7.52m0-4.48c0-.396.158-.776.44-1.056a1.504 1.504 0 0 1 2.12 0c.282.28.44.66.44 1.056v2.24m0 2.986V5.279c0-.396.158-.775.44-1.055a1.504 1.504 0 0 1 2.12 0c.282.28.44.66.44 1.055v2.987"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransmissionVirusTouchHandTwo(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.214 18.804a3.429 3.429 0 1 0 0-6.858a3.429 3.429 0 0 0 0 6.858m-.571-9.429h1.143m-.572 0v2.571m3.839-1.218l.808.808m-.404-.404l-1.818 1.819m3.575 1.853v1.142m0-.571h-2.571m1.218 3.839l-.808.808m.404-.404l-1.818-1.819m-1.853 3.576H6.643m.571 0v-2.571m-3.838 1.218l-.808-.808m.404.404l1.818-1.819m-3.576-1.853v-1.142m0 .571h2.572m-1.218-3.839l.808-.808m-.404.404l1.818 1.819m17.996-6.826l-2.937-2.35a5.255 5.255 0 0 0-3.28-1.15h-6.033a1.749 1.749 0 0 0-1.237 2.987a1.75 1.75 0 0 0 1.237.513h4.52"/><path d="M22.626 11.375h-3.5a5.25 5.25 0 0 1-2.914-.883l-2-1.333a1.634 1.634 0 0 1-.512-2.306v0a1.633 1.633 0 0 1 2.09-.555l2.625 1.577"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransmissionVirusTransmit(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12.267 2.142H9.233m1.517 0v2.6m-5.822-.817L3.856 4.997L2.783 6.07m1.073-1.073l1.838 1.839M1 10.375v3.033m0-1.516h2.6m-.817 5.821l1.073 1.073l1.072 1.072m-1.072-1.072l1.838-1.839M18.717 6.07l-1.073-1.073l-1.072-1.072m1.072 1.072l-1.838 1.839m-6.573 4.406a1.517 1.517 0 1 0 0-3.034a1.517 1.517 0 0 0 0 3.034"/><path stroke-linecap="round" stroke-linejoin="round" d="M17.606 9.858a7.15 7.15 0 1 0-8.889 8.884M12 21.858l.3-1.185a4.89 4.89 0 0 1 4.743-3.7H23"/><path stroke-linecap="round" stroke-linejoin="round" d="M19.333 20.636L23 16.969l-3.667-3.666"/><path d="M11.491 14.742a.375.375 0 1 1 0-.75m0 .75a.375.375 0 1 0 0-.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransmissionVirusTransportation(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.25 10.514a3.429 3.429 0 1 0 0-6.858a3.429 3.429 0 0 0 0 6.858m-.571-9.429h1.142m-.571 0v2.572m3.839-1.218l.808.808m-.404-.404l-1.819 1.818m3.576 1.853v1.143m0-.572h-2.571m1.218 3.839l-.808.808m.404-.404L19.674 9.51m-1.853 3.575h-1.142m.571 0v-2.571m-3.839 1.218l-.808-.808m.404.404l1.819-1.818M11.25 7.657V6.514m0 .571h2.571m-1.218-3.838l.808-.808m-.404.404l1.819 1.818M9.75 10.915h-1.5a1.5 1.5 0 0 0-1.5 1.5V19.2m11.212 2.215h.788a1.5 1.5 0 0 0 1.5-1.5v-4.83m-13.5-2.67h-3a3 3 0 0 0-3 3v4.5a1.5 1.5 0 0 0 1.5 1.5h2.287"/><path d="M6.375 22.915a1.875 1.875 0 1 0 0-3.75a1.875 1.875 0 0 0 0 3.75m9.75 0a1.875 1.875 0 1 0 0-3.75a1.875 1.875 0 0 0 0 3.75m-7.913-1.5h6.076M.75 17.016h6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransmissionVirusTrashBin(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 20.429a3.429 3.429 0 1 0 0-6.858a3.429 3.429 0 0 0 0 6.858M16.429 11h1.142M17 11v2.571m3.839-1.218l.808.808m-.404-.404l-1.819 1.819M23 16.429v1.142M23 17h-2.571m1.218 3.839l-.808.808m.404-.404l-1.819-1.819M17.571 23h-1.142M17 23v-2.571m-3.839 1.218l-.808-.808m.404.404l1.819-1.819M11 17.571v-1.142M11 17h2.571m-1.218-3.839l.808-.808m-.404.404l1.819 1.819M6.715 16.75a3 3 0 0 1-2.985-2.7L2.5 1h12l-.525 6M1 1h15M6.25 1L7 13.75M10.75 1l-.43 7.499"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransmissionVirusVisible(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.824 20.36a3.429 3.429 0 1 0 0-6.858a3.429 3.429 0 0 0 0 6.858m-.572-9.429h1.143m-.571 0v2.571m3.838-1.218l.808.808m-.404-.404l-1.818 1.818m3.576 1.853v1.143m0-.571h-2.572m1.218 3.838l-.808.808m.404-.404l-1.818-1.818m-1.853 3.576h-1.143m.572 0v-2.572m-3.839 1.218l-.808-.808m.404.404l1.818-1.818m-3.575-1.853v-1.143m0 .572h2.571m-1.218-3.839l.808-.808m-.404.404l1.818 1.818m7.648-4.889c.254-.253.498-.508.727-.762a1.667 1.667 0 0 0 0-2.226C20.163 3.746 16.013.924 11.999.995C7.986.925 3.835 3.746 1.226 6.63a1.667 1.667 0 0 0 0 2.226a19.308 19.308 0 0 0 6.287 4.647"/><path d="M9.256 10.376a3.751 3.751 0 1 1 6.466-2.091"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransmissionVirusWindBreath(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.75 20.679a3.429 3.429 0 1 0 0-6.858a3.429 3.429 0 0 0 0 6.858m-.571-9.429h1.142m-.571 0v2.571m3.839-1.218l.808.808m-.404-.404l-1.819 1.819m3.576 1.853v1.142m0-.571h-2.571m1.218 3.839l-.808.808m.404-.404l-1.819-1.819M7.321 23.25H6.179m.571 0v-2.571m-3.839 1.218l-.808-.808m.404.404l1.819-1.819M.75 17.821v-1.142m0 .571h2.571m-1.218-3.839l.808-.808m-.404.404l1.819 1.819M18.25 3.25a2.5 2.5 0 1 1 2.5 2.5H1.25m15.5 8.5a2.5 2.5 0 1 0 2.5-2.5h-5m-13-3h16.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VaccineProtectionFaceMaskOne(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 4.563a3 3 0 0 0-1.8.6c-.779.584-2.726 2.15-3.7 2.15L5 7.438v9l3.285 1.971A7.223 7.223 0 0 0 12 19.438m-7-3l-.357-.168A6.782 6.782 0 0 1 .75 10.135A2.981 2.981 0 0 1 5 7.438m7-2.875a3 3 0 0 1 1.8.6c.779.584 2.726 2.15 3.7 2.15l1.5.125v9l-3.285 1.971A7.223 7.223 0 0 1 12 19.438m7-3l.357-.168a6.782 6.782 0 0 0 3.893-6.135A2.981 2.981 0 0 0 19 7.438m-9.5 4.75h5m-2.5 2.5v-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VaccineProtectionFaceMaskThree(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 11.127a38.599 38.599 0 0 1 5-.5m-5 4.125c1.65.31 3.322.477 5 .5"/><path d="M12 19.5c2.743.074 5.47-.437 8-1.5V9s-4.9-4.5-8-4.5C8.9 4.5 4 9 4 9v9a19.313 19.313 0 0 0 8 1.5"/><path d="M17 11.127a38.598 38.598 0 0 0-5-.5m5 4.125c-1.65.31-3.322.477-5 .5M4 18a6.828 6.828 0 0 1-3.25-5.814V8.054a1.082 1.082 0 0 1 1.166-1.08A1.467 1.467 0 0 1 3 7.59L4 9m16 9a6.828 6.828 0 0 0 3.25-5.816v-4.13a1.083 1.083 0 0 0-1.166-1.08A1.467 1.467 0 0 0 21 7.59L20 9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VaccineProtectionFaceMaskTwo(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 20.582a19.224 19.224 0 0 1-7.964-1.493V10.13s5.973-3.11 7.963-3.11"/><path d="M7.022 12.868A38.424 38.424 0 0 1 12 12.37m-4.978 3.982a29.11 29.11 0 0 0 4.977.497"/><path d="M12 20.582c2.73.074 5.445-.435 7.963-1.493V10.13s-5.973-3.11-7.964-3.11"/><path d="M16.977 12.868a38.424 38.424 0 0 0-4.978-.498m4.978 3.982a29.11 29.11 0 0 1-4.978.497M4.036 10.13l-.749-.398A4.727 4.727 0 0 1 .801 5.564a2.154 2.154 0 0 1 2.153-2.153c.386 0 .768.086 1.116.254a2.104 2.104 0 0 1 .712 3.23m-.746 12.194A6.797 6.797 0 0 1 .8 13.299a8.87 8.87 0 0 1 .127-1.61m19.036-1.559l.748-.398A4.728 4.728 0 0 0 23.2 5.569a2.154 2.154 0 0 0-2.153-2.153c-.386 0-.768.086-1.116.254a2.104 2.104 0 0 0-.713 3.23m.745 12.189a6.797 6.797 0 0 0 3.235-5.79a8.869 8.869 0 0 0-.127-1.61"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VaccineProtectionFaceShieldOne(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 .81c-3.559.543-6 1.645-6 2.92c0 1.811 4.925 3.28 11 3.28s11-1.469 11-3.28c0-1.275-2.441-2.377-6-2.92"/><path d="M1 3.73v3.3c0 1.812 4.925 3.281 11 3.281s11-1.465 11-3.277V3.73"/><path d="m4 9.315l-.9 9.794a1.98 1.98 0 0 0 .848 1.823A14.277 14.277 0 0 0 12 23.19c2.85.068 5.655-.718 8.054-2.258a1.98 1.98 0 0 0 .848-1.823L20 9.315"/><path d="M13 20.19a7.638 7.638 0 0 0 5-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VaccineProtectionFaceShieldTwo(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18.167 17.107V15.06H19.7a1.02 1.02 0 0 0 .935-1.435c-1-2.239-2.476-4.994-2.476-4.994V6.45a2.024 2.024 0 0 0-.18-.847c-1.111-2.409-2.617-4.074-5.5-4.652C6.611-.224.807 3.887.763 9.871a9.19 9.19 0 0 0 3.071 6.937v6.442m10.238 0v-1.34M.777 9.433h10.418"/><path d="M11.2 4.628h8.042a4 4 0 0 1 4 4v11.577H18.2a7 7 0 0 1-7-7zm-.005 0H2.528"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VaccineProtectionInfraredThermometerGun(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.394 20.75a2.34 2.34 0 0 0 2.3-2.8l-2.3-11.483a4 4 0 0 0-3.923-3.217H4.62a1.5 1.5 0 0 0-1.464 1.825L4.4 10.684a2 2 0 0 0 1.954 1.566h4.566a4 4 0 0 1 3.772 2.669l1.292 3.66a3.254 3.254 0 0 0 3.069 2.171z"/><path d="M11.769 12.34a6.23 6.23 0 0 1-.569 3.56a.6.6 0 0 0 .552.854h3.588M19.594 4.75h2.156a1.5 1.5 0 0 1 1.5 1.5v1.5a1.5 1.5 0 0 1-1.5 1.5h-.8m-16.532 1.5H2.25a1.5 1.5 0 0 1-1.5-1.5v-3a1.5 1.5 0 0 1 1.5-1.5h.87m11.13 3a1.5 1.5 0 0 1-1.5 1.5h-3.5a1.5 1.5 0 0 1 0-3h3.5a1.5 1.5 0 0 1 1.5 1.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VaccineProtectionMedicinePill(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.6 16.4a4.6 4.6 0 0 1-6.5-6.5l7.8-7.8a4.6 4.6 0 0 1 6.5 6.5M5.999 5.999l4.25 4.25"/><path d="M23.25 15.75a7.669 7.669 0 0 1-6 7.5a7.669 7.669 0 0 1-6-7.5V13.5a1.5 1.5 0 0 1 1.5-1.5h9a1.5 1.5 0 0 1 1.5 1.5zm-6-.75v4.5M15 17.25h4.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VaccineProtectionPeopleShield(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.75 20.893a3.143 3.143 0 1 0 0-6.286a3.143 3.143 0 0 0 0 6.286m-.524-8.643h1.048m-.524 0v2.357m3.519-1.116l.74.74m-.37-.37l-1.667 1.667m3.278 1.698v1.048m0-.524h-2.357m1.116 3.519l-.74.74m.37-.37l-1.667-1.667m-1.698 3.278h-1.048m.524 0v-2.357m-3.519 1.116l-.74-.74m.37.37l1.667-1.667m-3.278-1.698v-1.048m0 .524h2.357m-1.116-3.519l.74-.74m-.37.37l1.667 1.667"/><path d="M19.25 8.25V3.76a1.411 1.411 0 0 0-.823-1.292A20.6 20.6 0 0 0 10 .75a20.6 20.6 0 0 0-8.427 1.718A1.411 1.411 0 0 0 .75 3.76v7.224c0 3.532 1.546 8.352 8.228 10.922a2.847 2.847 0 0 0 2.044 0"/><path d="M10.147 9.996a2.791 2.791 0 1 0 0-5.582a2.791 2.791 0 0 0 0 5.582m1.72 1.74a4.571 4.571 0 0 0-6.288 4.233"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VaccineProtectionSanitizerSpray(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19.9 15.23a3.429 3.429 0 1 1-6.142 1.393m9.367.056v1.142m0-.571h-2.571m1.218 3.839l-.808.808m.404-.404l-1.819-1.819m-1.853 3.576h-1.142m.571 0v-2.571m-3.839 1.218l-.808-.808m.404.404l1.819-1.819m-3.576-1.853v-1.142m0 .571h2.571m-8.159-7.037a1.5 1.5 0 0 1-2.121-.013L1.308 8.065a1.5 1.5 0 0 1 .013-2.122l3.116-3.076A1.5 1.5 0 0 1 5.5 2.434l1.5.01l2.108 2.134l-.008 1.5a1.5 1.5 0 0 1-.446 1.058z"/><path d="m9.837 1.047l.703.711a1 1 0 0 1-.009 1.415L9.108 4.578L7 2.443l1.423-1.405a1 1 0 0 1 1.414.009m1.616 6.423l1.265 2.13a2.09 2.09 0 1 0 2.34 3.442a2.783 2.783 0 0 0 3.911-3.862a2.09 2.09 0 1 0-3.411-2.384l-2.117-1.288M4.876 20.436a2.127 2.127 0 1 0 0-4.254a2.127 2.127 0 0 0 0 4.254m0-5.85v1.596m2.633-.505L6.38 16.805m2.219 1.504H7.003m.506 2.633L6.38 19.813m-1.504 2.219v-1.596m-2.632.506l1.128-1.129m-2.219-1.504h1.596m-.505-2.632l1.128 1.128"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VaccineProtectionShield(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.125 14.339a3.714 3.714 0 1 0 0-7.428a3.714 3.714 0 0 0 0 7.428m-.619-10.214h1.238m-.619 0v2.786m4.158-1.321l.876.877m-.438-.438l-1.97 1.97m3.874 2.007v1.238m0-.619h-2.786m1.32 4.158l-.876.876m.438-.438l-1.97-1.97m-2.007 3.874h-1.238m.619 0v-2.786m-4.158 1.32l-.876-.876m.438.438l1.97-1.97m-3.874-2.007v-1.238m0 .619h2.786m-1.32-4.158l.876-.876m-.438.438l1.97 1.97"/><path d="M2.25 3.923v7.614c0 3.723 1.629 8.8 8.673 11.513a3 3 0 0 0 2.154 0c7.041-2.708 8.673-7.822 8.673-11.513V3.923a1.486 1.486 0 0 0-.868-1.362A21.7 21.7 0 0 0 12 .75a21.7 21.7 0 0 0-8.882 1.811a1.486 1.486 0 0 0-.868 1.362"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VaccineProtectionSyringe(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m17.123 9.03l-4.816-4.816L4.121 12.4a1.362 1.362 0 0 0 0 1.926L6.689 16.9a1.5 1.5 0 0 0 2.124 0v0"/><path d="m4.602 14.808l-2.408 2.408l1.926 1.926l2.408-2.408zM.75 20.587l2.408-2.408M10.863 2.77l1.444 1.444m1.445 1.445l3.371-3.371m1.926 1.926l-3.371 3.371m0-6.742l4.816 4.816m2.756 9.998a7.67 7.67 0 0 1-6 7.5a7.67 7.67 0 0 1-6-7.5v-2.25a1.5 1.5 0 0 1 1.5-1.5h9a1.5 1.5 0 0 1 1.5 1.5zm-6-.75v4.5M15 17.157h4.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VaccineProtectionVirus(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.267.75H9.233m1.517 0v2.6m-5.822-.817L3.856 3.606L2.783 4.678m1.073-1.072l1.838 1.838M1 8.983v3.034M1 10.5h2.6m-.817 5.822l1.073 1.072l1.072 1.073m-1.072-1.073l1.838-1.838M18.717 4.678l-1.073-1.072l-1.072-1.073m1.072 1.073l-1.838 1.838M9.233 9.85a1.517 1.517 0 1 0 0-3.034a1.517 1.517 0 0 0 0 3.034"/><path d="M17.538 8.25A7.15 7.15 0 1 0 8 17.1"/><path d="M23 15.75a7.669 7.669 0 0 1-6 7.5a7.669 7.669 0 0 1-6-7.5V13.5a1.5 1.5 0 0 1 1.5-1.5h9a1.5 1.5 0 0 1 1.5 1.5zM17 15v4.5m-2.25-2.25h4.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VaccineProtectionWashHands(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 4.5h-3a1.5 1.5 0 0 0 0-3H8.158a4.5 4.5 0 0 0-3.744 2L3.2 5.332A1.5 1.5 0 0 1 1.947 6H.75m11-1.5H14M.75 21.75h1.9c.233 0 .463.054.671.158l1.733.867a4.507 4.507 0 0 0 2.012.475H18.75a1.5 1.5 0 1 0 0-3H21a1.5 1.5 0 1 0 0-3h.75a1.5 1.5 0 1 0 0-3h-1.5a1.5 1.5 0 1 0 0-3H15a1.5 1.5 0 1 0 0-3H9.158a4.5 4.5 0 0 0-3.744 2L4.2 12.082a1.5 1.5 0 0 1-1.248.668H.75m12-1.5H15m1.5 6H21m-4.5-3h3.75m-4.5 6h3M21.007.75a.455.455 0 0 1 .423.289c.414 1.063 1.82 4.749 1.82 5.789a2.21 2.21 0 1 1-4.421 0c0-1.062 1.465-4.878 1.845-5.851a.357.357 0 0 1 .333-.227"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VirusLabResearchMagnifierOne(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.017.75H8.983m1.517 0v2.6m-5.822-.817L3.606 3.606L2.533 4.678m1.073-1.072l1.838 1.838M.75 8.983v3.034m0-1.517h2.6m-.817 5.822l1.073 1.072l1.072 1.073m-1.072-1.073l1.838-1.838M18.467 4.678l-1.073-1.072l-1.072-1.073m1.072 1.073l-1.838 1.838M8.983 9.85a1.517 1.517 0 1 0 0-3.034a1.517 1.517 0 0 0 0 3.034"/><path d="M17.435 8.755A7.15 7.15 0 1 0 8.5 17.361"/><path d="M16.264 21.154a4.89 4.89 0 1 0 0-9.78a4.89 4.89 0 0 0 0 9.78m6.986 2.096l-3.528-3.528"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VirusLabResearchMagnifierTwo(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.617 13.837a3.22 3.22 0 1 0 0-6.44a3.22 3.22 0 0 0 0 6.44m-.537-8.855h1.074m-.537 0v2.415m3.605-1.144l.759.76m-.38-.38L12.894 8.34m3.358 1.74v1.074m0-.537h-2.415m1.144 3.605l-.759.759m.379-.38l-1.707-1.707m-1.74 3.358H10.08m.537 0v-2.415m-3.605 1.144l-.759-.759m.38.379l1.707-1.707m-3.358-1.74V10.08m0 .537h2.415M6.253 7.012l.759-.759m-.379.38L8.34 8.34"/><path d="M10.617 20.484c5.45 0 9.867-4.418 9.867-9.867c0-5.45-4.418-9.867-9.867-9.867C5.167.75.75 5.168.75 10.617c0 5.45 4.418 9.867 9.867 9.867M23.25 23.25l-5.656-5.656"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VirusLabResearchMedicinePill(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.25 20.679a3.429 3.429 0 1 0 0-6.858a3.429 3.429 0 0 0 0 6.858m-.571-9.429h1.142m-.571 0v2.571m3.839-1.218l.808.808m-.404-.404l-1.819 1.819m3.576 1.853v1.142m0-.571h-2.571m1.218 3.839l-.808.808m.404-.404l-1.819-1.819m-1.853 3.576h-1.142m.571 0v-2.571m-3.839 1.218l-.808-.808m.404.404l1.819-1.819m-3.576-1.853v-1.142m0 .571h2.571m-1.218-3.839l.808-.808m-.404.404l1.819 1.819M8.6 16.4a4.6 4.6 0 0 1-6.5-6.5l7.8-7.8a4.6 4.6 0 0 1 6.5 6.5M5.999 5.999l5.138 5.138"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VirusLabResearchMicroscope(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.047 9.375a3 3 0 1 0 0-6a3 3 0 0 0 0 6m-.5-8.25h1m-.5 0v2.25m3.359-1.066l.707.707m-.354-.353L8.168 4.254m3.129 1.621v1m0-.5h-2.25m1.066 3.359l-.707.707m.353-.354L8.168 8.496m-1.621 3.129h-1m.5 0v-2.25m-3.359 1.066l-.707-.707m.354.353l1.591-1.591M.797 6.875v-1m0 .5h2.25M1.981 3.016l.707-.707m-.353.354l1.591 1.591M18.146 9.1a6.134 6.134 0 0 1 .315 9.07a6.4 6.4 0 0 1-8.941 0a6.17 6.17 0 0 1-1.558-2.52"/><path d="M16.151 11.054a.913.913 0 0 1-1.277 0L13.6 9.8a.875.875 0 0 1 0-1.254l3.832-3.764a.915.915 0 0 1 1.277 0l1.277 1.255a.875.875 0 0 1 0 1.254z"/><path d="m19.168 5.236l2.59-2.592l-1.445-1.445m2.89 2.89l-1.445-1.445M5.797 15.65h5.058m-4.27 7.225h13.728m-6.503-2.89v2.89"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VirusLabResearchSyringe(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.964 20.675a3.429 3.429 0 1 0 0-6.858a3.429 3.429 0 0 0 0 6.858m-.571-9.429h1.142m-.571 0v2.571m3.839-1.218l.808.808m-.404-.404l-1.819 1.818m3.576 1.853v1.143m0-.571h-2.571m1.218 3.838l-.808.809m.404-.405l-1.819-1.818m-1.853 3.576h-1.142m.571 0v-2.572m-3.839 1.219l-.808-.809m.404.404l1.819-1.818m-3.576-1.853v-1.143m0 .572h2.571m-1.218-3.839l.808-.808m-.404.404l1.819 1.818M7.3 17.127l-2.893-2.889a1.363 1.363 0 0 1 0-1.927l8.186-8.186l4.816 4.815M4.888 14.719l-2.407 2.408l1.926 1.926l2.407-2.408zm-3.852 5.779l2.408-2.408m7.705-15.41l5.815 5.816"/><path d="m17.409 2.199l-3.371 3.37l1.926 1.927l3.371-3.371zM15.964.754L20.78 5.57"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VirusLabResearchTestTube(children ...ElementRenderer) *CovidIcon {
	return &CovidIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17 20.679a3.429 3.429 0 1 0 0-6.858a3.429 3.429 0 0 0 0 6.858m-.571-9.429h1.142m-.571 0v2.571m3.839-1.218l.808.808m-.404-.404l-1.819 1.819M23 16.679v1.142m0-.571h-2.571m1.218 3.839l-.808.808m.404-.404l-1.819-1.819m-1.853 3.576h-1.142m.571 0v-2.571m-3.839 1.218l-.808-.808m.404.404l1.819-1.819M11 17.821v-1.142m0 .571h2.571m-1.218-3.839l.808-.808m-.404.404l1.819 1.819M9.549 22.25a3.75 3.75 0 0 1-6.3-2.75V.75h7.5v9.5M13 .75H1"/><path d="M8.125 8.25a.375.375 0 0 1 .375.375m-.75 0a.375.375 0 0 1 .375-.375m0 .75a.375.375 0 0 1-.375-.375m.75 0A.375.375 0 0 1 8.125 9m-2.25 3.75a.375.375 0 0 1 .375.375m-.75 0a.375.375 0 0 1 .375-.375m0 .75a.375.375 0 0 1-.375-.375m.75 0a.375.375 0 0 1-.375.375m1.5 3.75a.375.375 0 0 1 .375.375m-.75 0a.375.375 0 0 1 .375-.375m0 .75A.375.375 0 0 1 7 17.625m.75 0a.375.375 0 0 1-.375.375M10.75 5.25h-7.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
