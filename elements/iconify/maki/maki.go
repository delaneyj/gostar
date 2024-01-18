package maki

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "8.0.1"
	hAttr          = "1em"
	viewbox        = "0 0 15 15"
	fill           = "currentColor"
)

type MakiIcon struct {
	*SVGSVGElement
}

type MakiIconFn func(children ...ElementRenderer) *MakiIcon

var IconLookup = map[string]MakiIconFn{
	"aerialway":                  Aerialway,
	"aerialwayEleven":            AerialwayEleven,
	"airfield":                   Airfield,
	"airfieldEleven":             AirfieldEleven,
	"airport":                    Airport,
	"airportEleven":              AirportEleven,
	"alcoholShop":                AlcoholShop,
	"alcoholShopEleven":          AlcoholShopEleven,
	"americanFootball":           AmericanFootball,
	"americanFootballEleven":     AmericanFootballEleven,
	"amusementPark":              AmusementPark,
	"amusementParkEleven":        AmusementParkEleven,
	"animalShelter":              AnimalShelter,
	"aquarium":                   Aquarium,
	"aquariumEleven":             AquariumEleven,
	"arrow":                      Arrow,
	"artGallery":                 ArtGallery,
	"artGalleryEleven":           ArtGalleryEleven,
	"attraction":                 Attraction,
	"attractionEleven":           AttractionEleven,
	"bakery":                     Bakery,
	"bakeryEleven":               BakeryEleven,
	"bakeryFifteen":              BakeryFifteen,
	"bank":                       Bank,
	"bankEleven":                 BankEleven,
	"bankJp":                     BankJp,
	"bankJpEleven":               BankJpEleven,
	"bar":                        Bar,
	"barEleven":                  BarEleven,
	"barrier":                    Barrier,
	"barrierEleven":              BarrierEleven,
	"baseball":                   Baseball,
	"baseballEleven":             BaseballEleven,
	"basketball":                 Basketball,
	"basketballEleven":           BasketballEleven,
	"bbq":                        Bbq,
	"bbqEleven":                  BbqEleven,
	"beach":                      Beach,
	"beachEleven":                BeachEleven,
	"beer":                       Beer,
	"beerEleven":                 BeerEleven,
	"bicycle":                    Bicycle,
	"bicycleEleven":              BicycleEleven,
	"bicycleShare":               BicycleShare,
	"bicycleShareEleven":         BicycleShareEleven,
	"bloodBank":                  BloodBank,
	"bloodBankEleven":            BloodBankEleven,
	"bowlingAlley":               BowlingAlley,
	"bowlingAlleyEleven":         BowlingAlleyEleven,
	"bridge":                     Bridge,
	"bridgeEleven":               BridgeEleven,
	"building":                   Building,
	"buildingAltOne":             BuildingAltOne,
	"buildingAltOneEleven":       BuildingAltOneEleven,
	"buildingEleven":             BuildingEleven,
	"bus":                        Bus,
	"busEleven":                  BusEleven,
	"cafe":                       Cafe,
	"cafeEleven":                 CafeEleven,
	"campsite":                   Campsite,
	"campsiteEleven":             CampsiteEleven,
	"car":                        Car,
	"carEleven":                  CarEleven,
	"carFifteen":                 CarFifteen,
	"carRental":                  CarRental,
	"carRentalEleven":            CarRentalEleven,
	"carRentalFifteen":           CarRentalFifteen,
	"carRepair":                  CarRepair,
	"carRepairEleven":            CarRepairEleven,
	"carRepairFifteen":           CarRepairFifteen,
	"casino":                     Casino,
	"casinoEleven":               CasinoEleven,
	"castle":                     Castle,
	"castleEleven":               CastleEleven,
	"castleJp":                   CastleJp,
	"castleJpEleven":             CastleJpEleven,
	"caution":                    Caution,
	"cemetery":                   Cemetery,
	"cemeteryEleven":             CemeteryEleven,
	"cemeteryJp":                 CemeteryJp,
	"cemeteryJpEleven":           CemeteryJpEleven,
	"chargingStation":            ChargingStation,
	"chargingStationEleven":      ChargingStationEleven,
	"cinema":                     Cinema,
	"cinemaEleven":               CinemaEleven,
	"circle":                     Circle,
	"circleEleven":               CircleEleven,
	"circleStroked":              CircleStroked,
	"circleStrokedEleven":        CircleStrokedEleven,
	"city":                       City,
	"cityEleven":                 CityEleven,
	"clothingStore":              ClothingStore,
	"clothingStoreEleven":        ClothingStoreEleven,
	"college":                    College,
	"collegeEleven":              CollegeEleven,
	"collegeJp":                  CollegeJp,
	"collegeJpEleven":            CollegeJpEleven,
	"commercial":                 Commercial,
	"commercialEleven":           CommercialEleven,
	"communicationsTower":        CommunicationsTower,
	"communicationsTowerEleven":  CommunicationsTowerEleven,
	"confectionery":              Confectionery,
	"confectioneryEleven":        ConfectioneryEleven,
	"construction":               Construction,
	"convenience":                Convenience,
	"convenienceEleven":          ConvenienceEleven,
	"cricket":                    Cricket,
	"cricketEleven":              CricketEleven,
	"cross":                      Cross,
	"crossEleven":                CrossEleven,
	"dam":                        Dam,
	"damEleven":                  DamEleven,
	"danger":                     Danger,
	"dangerEleven":               DangerEleven,
	"defibrillator":              Defibrillator,
	"defibrillatorEleven":        DefibrillatorEleven,
	"dentist":                    Dentist,
	"dentistEleven":              DentistEleven,
	"diamond":                    Diamond,
	"doctor":                     Doctor,
	"doctorEleven":               DoctorEleven,
	"dogPark":                    DogPark,
	"dogParkEleven":              DogParkEleven,
	"drinkingWater":              DrinkingWater,
	"drinkingWaterEleven":        DrinkingWaterEleven,
	"drinkingWaterFifteen":       DrinkingWaterFifteen,
	"elevator":                   Elevator,
	"embassy":                    Embassy,
	"embassyEleven":              EmbassyEleven,
	"emergencyPhone":             EmergencyPhone,
	"emergencyPhoneEleven":       EmergencyPhoneEleven,
	"entrance":                   Entrance,
	"entranceAltOne":             EntranceAltOne,
	"entranceAltOneEleven":       EntranceAltOneEleven,
	"entranceEleven":             EntranceEleven,
	"entranceFifteen":            EntranceFifteen,
	"farm":                       Farm,
	"farmEleven":                 FarmEleven,
	"fastFood":                   FastFood,
	"fastFoodEleven":             FastFoodEleven,
	"fence":                      Fence,
	"fenceEleven":                FenceEleven,
	"ferry":                      Ferry,
	"ferryEleven":                FerryEleven,
	"ferryJp":                    FerryJp,
	"fireStation":                FireStation,
	"fireStationEleven":          FireStationEleven,
	"fireStationJp":              FireStationJp,
	"fireStationJpEleven":        FireStationJpEleven,
	"fitnessCentre":              FitnessCentre,
	"fitnessCentreEleven":        FitnessCentreEleven,
	"florist":                    Florist,
	"floristEleven":              FloristEleven,
	"fuel":                       Fuel,
	"fuelEleven":                 FuelEleven,
	"furniture":                  Furniture,
	"furnitureEleven":            FurnitureEleven,
	"furnitureFifteen":           FurnitureFifteen,
	"gaming":                     Gaming,
	"gamingEleven":               GamingEleven,
	"garden":                     Garden,
	"gardenCentre":               GardenCentre,
	"gardenCentreEleven":         GardenCentreEleven,
	"gardenEleven":               GardenEleven,
	"gate":                       Gate,
	"gift":                       Gift,
	"giftEleven":                 GiftEleven,
	"globe":                      Globe,
	"globeEleven":                GlobeEleven,
	"globeFifteen":               GlobeFifteen,
	"golf":                       Golf,
	"golfEleven":                 GolfEleven,
	"grocery":                    Grocery,
	"groceryEleven":              GroceryEleven,
	"hairdresser":                Hairdresser,
	"hairdresserEleven":          HairdresserEleven,
	"harbor":                     Harbor,
	"harborEleven":               HarborEleven,
	"hardware":                   Hardware,
	"hardwareEleven":             HardwareEleven,
	"heart":                      Heart,
	"heartEleven":                HeartEleven,
	"heliport":                   Heliport,
	"heliportEleven":             HeliportEleven,
	"highwayRestArea":            HighwayRestArea,
	"historic":                   Historic,
	"home":                       Home,
	"homeEleven":                 HomeEleven,
	"horseRiding":                HorseRiding,
	"horseRidingEleven":          HorseRidingEleven,
	"hospital":                   Hospital,
	"hospitalEleven":             HospitalEleven,
	"hospitalJp":                 HospitalJp,
	"hospitalJpEleven":           HospitalJpEleven,
	"hotSpring":                  HotSpring,
	"iceCream":                   IceCream,
	"iceCreamEleven":             IceCreamEleven,
	"iceCreamFifteen":            IceCreamFifteen,
	"industry":                   Industry,
	"industryEleven":             IndustryEleven,
	"information":                Information,
	"informationEleven":          InformationEleven,
	"jewelryStore":               JewelryStore,
	"jewelryStoreEleven":         JewelryStoreEleven,
	"karaoke":                    Karaoke,
	"karaokeEleven":              KaraokeEleven,
	"karaokeFifteen":             KaraokeFifteen,
	"landmark":                   Landmark,
	"landmarkEleven":             LandmarkEleven,
	"landmarkJp":                 LandmarkJp,
	"landmarkJpEleven":           LandmarkJpEleven,
	"landuse":                    Landuse,
	"landuseEleven":              LanduseEleven,
	"laundry":                    Laundry,
	"laundryEleven":              LaundryEleven,
	"library":                    Library,
	"libraryEleven":              LibraryEleven,
	"libraryFifteen":             LibraryFifteen,
	"liftGate":                   LiftGate,
	"lighthouse":                 Lighthouse,
	"lighthouseEleven":           LighthouseEleven,
	"lighthouseJp":               LighthouseJp,
	"lodging":                    Lodging,
	"lodgingEleven":              LodgingEleven,
	"logging":                    Logging,
	"loggingEleven":              LoggingEleven,
	"marker":                     Marker,
	"markerEleven":               MarkerEleven,
	"markerStroked":              MarkerStroked,
	"markerStrokedEleven":        MarkerStrokedEleven,
	"mobilePhone":                MobilePhone,
	"mobilePhoneEleven":          MobilePhoneEleven,
	"monument":                   Monument,
	"monumentEleven":             MonumentEleven,
	"monumentJp":                 MonumentJp,
	"mountain":                   Mountain,
	"mountainEleven":             MountainEleven,
	"museum":                     Museum,
	"museumEleven":               MuseumEleven,
	"museumFifteen":              MuseumFifteen,
	"music":                      Music,
	"musicEleven":                MusicEleven,
	"natural":                    Natural,
	"naturalEleven":              NaturalEleven,
	"observationTower":           ObservationTower,
	"optician":                   Optician,
	"opticianEleven":             OpticianEleven,
	"paint":                      Paint,
	"paintEleven":                PaintEleven,
	"park":                       Park,
	"parkAltOne":                 ParkAltOne,
	"parkAltOneEleven":           ParkAltOneEleven,
	"parkEleven":                 ParkEleven,
	"parking":                    Parking,
	"parkingEleven":              ParkingEleven,
	"parkingGarage":              ParkingGarage,
	"parkingGarageEleven":        ParkingGarageEleven,
	"parkingPaid":                ParkingPaid,
	"pharmacy":                   Pharmacy,
	"pharmacyEleven":             PharmacyEleven,
	"picnicSite":                 PicnicSite,
	"picnicSiteEleven":           PicnicSiteEleven,
	"pitch":                      Pitch,
	"pitchEleven":                PitchEleven,
	"placeOfWorship":             PlaceOfWorship,
	"placeOfWorshipEleven":       PlaceOfWorshipEleven,
	"playground":                 Playground,
	"playgroundEleven":           PlaygroundEleven,
	"police":                     Police,
	"policeEleven":               PoliceEleven,
	"policeJp":                   PoliceJp,
	"policeJpEleven":             PoliceJpEleven,
	"post":                       Post,
	"postEleven":                 PostEleven,
	"postJp":                     PostJp,
	"postJpEleven":               PostJpEleven,
	"prison":                     Prison,
	"prisonEleven":               PrisonEleven,
	"racetrack":                  Racetrack,
	"racetrackBoat":              RacetrackBoat,
	"racetrackCycling":           RacetrackCycling,
	"racetrackHorse":             RacetrackHorse,
	"rail":                       Rail,
	"railEleven":                 RailEleven,
	"railLight":                  RailLight,
	"railLightEleven":            RailLightEleven,
	"railMetro":                  RailMetro,
	"railMetroEleven":            RailMetroEleven,
	"rangerStation":              RangerStation,
	"rangerStationEleven":        RangerStationEleven,
	"recycling":                  Recycling,
	"recyclingEleven":            RecyclingEleven,
	"religiousBuddhist":          ReligiousBuddhist,
	"religiousBuddhistEleven":    ReligiousBuddhistEleven,
	"religiousChristian":         ReligiousChristian,
	"religiousChristianEleven":   ReligiousChristianEleven,
	"religiousJewish":            ReligiousJewish,
	"religiousJewishEleven":      ReligiousJewishEleven,
	"religiousMuslim":            ReligiousMuslim,
	"religiousMuslimEleven":      ReligiousMuslimEleven,
	"religiousShinto":            ReligiousShinto,
	"religiousShintoEleven":      ReligiousShintoEleven,
	"residentialCommunity":       ResidentialCommunity,
	"residentialCommunityEleven": ResidentialCommunityEleven,
	"restaurant":                 Restaurant,
	"restaurantBbq":              RestaurantBbq,
	"restaurantEleven":           RestaurantEleven,
	"restaurantNoodle":           RestaurantNoodle,
	"restaurantNoodleEleven":     RestaurantNoodleEleven,
	"restaurantPizza":            RestaurantPizza,
	"restaurantPizzaEleven":      RestaurantPizzaEleven,
	"restaurantSeafood":          RestaurantSeafood,
	"restaurantSeafoodEleven":    RestaurantSeafoodEleven,
	"restaurantSushi":            RestaurantSushi,
	"roadAccident":               RoadAccident,
	"roadblock":                  Roadblock,
	"roadblockEleven":            RoadblockEleven,
	"rocket":                     Rocket,
	"rocketEleven":               RocketEleven,
	"school":                     School,
	"schoolEleven":               SchoolEleven,
	"schoolJp":                   SchoolJp,
	"schoolJpEleven":             SchoolJpEleven,
	"scooter":                    Scooter,
	"scooterEleven":              ScooterEleven,
	"shelter":                    Shelter,
	"shelterEleven":              ShelterEleven,
	"shoe":                       Shoe,
	"shoeEleven":                 ShoeEleven,
	"shoeFifteen":                ShoeFifteen,
	"shop":                       Shop,
	"shopEleven":                 ShopEleven,
	"skateboard":                 Skateboard,
	"skateboardEleven":           SkateboardEleven,
	"skiing":                     Skiing,
	"skiingEleven":               SkiingEleven,
	"slaughterhouse":             Slaughterhouse,
	"slaughterhouseEleven":       SlaughterhouseEleven,
	"slipway":                    Slipway,
	"slipwayEleven":              SlipwayEleven,
	"snowmobile":                 Snowmobile,
	"snowmobileEleven":           SnowmobileEleven,
	"soccer":                     Soccer,
	"soccerEleven":               SoccerEleven,
	"square":                     Square,
	"squareEleven":               SquareEleven,
	"squareStroked":              SquareStroked,
	"squareStrokedEleven":        SquareStrokedEleven,
	"stadium":                    Stadium,
	"stadiumEleven":              StadiumEleven,
	"star":                       Star,
	"starEleven":                 StarEleven,
	"starStroked":                StarStroked,
	"starStrokedEleven":          StarStrokedEleven,
	"suitcase":                   Suitcase,
	"suitcaseEleven":             SuitcaseEleven,
	"sushiEleven":                SushiEleven,
	"swimming":                   Swimming,
	"swimmingEleven":             SwimmingEleven,
	"tableTennis":                TableTennis,
	"tableTennisEleven":          TableTennisEleven,
	"teahouse":                   Teahouse,
	"teahouseEleven":             TeahouseEleven,
	"telephone":                  Telephone,
	"telephoneEleven":            TelephoneEleven,
	"tennis":                     Tennis,
	"tennisEleven":               TennisEleven,
	"terminal":                   Terminal,
	"theatre":                    Theatre,
	"theatreEleven":              TheatreEleven,
	"toilet":                     Toilet,
	"toiletEleven":               ToiletEleven,
	"toll":                       Toll,
	"town":                       Town,
	"townEleven":                 TownEleven,
	"townHall":                   TownHall,
	"townHallEleven":             TownHallEleven,
	"townHallFifteen":            TownHallFifteen,
	"triangle":                   Triangle,
	"triangleEleven":             TriangleEleven,
	"triangleStroked":            TriangleStroked,
	"triangleStrokedEleven":      TriangleStrokedEleven,
	"tunnel":                     Tunnel,
	"veterinary":                 Veterinary,
	"veterinaryEleven":           VeterinaryEleven,
	"veterinaryFifteen":          VeterinaryFifteen,
	"viewpoint":                  Viewpoint,
	"viewpointEleven":            ViewpointEleven,
	"village":                    Village,
	"villageEleven":              VillageEleven,
	"volcano":                    Volcano,
	"volcanoEleven":              VolcanoEleven,
	"volleyball":                 Volleyball,
	"volleyballEleven":           VolleyballEleven,
	"warehouse":                  Warehouse,
	"warehouseEleven":            WarehouseEleven,
	"wasteBasket":                WasteBasket,
	"wasteBasketEleven":          WasteBasketEleven,
	"watch":                      Watch,
	"watchEleven":                WatchEleven,
	"watchFifteen":               WatchFifteen,
	"water":                      Water,
	"waterEleven":                WaterEleven,
	"waterfall":                  Waterfall,
	"waterfallEleven":            WaterfallEleven,
	"watermill":                  Watermill,
	"watermillEleven":            WatermillEleven,
	"wetland":                    Wetland,
	"wetlandEleven":              WetlandEleven,
	"wheelchair":                 Wheelchair,
	"wheelchairEleven":           WheelchairEleven,
	"windmill":                   Windmill,
	"windmillEleven":             WindmillEleven,
	"zoo":                        Zoo,
	"zooEleven":                  ZooEleven,
}

func Aerialway(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 5H8V2.6a1 1 0 0 0 .42-.46l5.08-.64a.5.5 0 0 0 0-1l-5.22.65a1 1 0 0 0-.78-.4a1 1 0 0 0-.92.62L1.5 2a.5.5 0 0 0 0 1l5.22-.65c.077.1.172.185.28.25V5H2a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1m-6 6H3V7h4zm5 0H8V7h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AerialwayEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9 4.5H6V3.1a1 1 0 0 0 .43-.52L9.5 2a.5.5 0 0 0 0-1l-3.25.61a1 1 0 0 0-1.69.32L1.5 2.5a.5.5 0 0 0 0 1l3.25-.61A1 1 0 0 0 5 3.1v1.4H2a1 1 0 0 0-1 1V9a1 1 0 0 0 1 1h7a1 1 0 0 0 1-1V5.5a1 1 0 0 0-1-1zm-4 4H2.5v-3H5v3zm3.5 0H6v-3h2.5v3z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Airfield(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.818.682H4.773C4.09.682 4.09 0 4.773 0h5.454c.682 0 .682.682 0 .682H8.182S9 1.272 9 2.636V4h6v2L9 8l-.5 5l2.5 1.318V15H4v-.682L6.5 13L6 8L0 6V4h6V2.636c0-1.363.818-1.954.818-1.954"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirfieldEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5 .5H3.5C3 .5 3 0 3.5 0h4c.5 0 .5.5 0 .5H6s.5.5.5 1.5v1H11v1.5l-4.5 2L6 10l1.5.5v.5h-4v-.5L5 10l-.5-3.5l-4.5-2V3h4.5V2C4.5 1 5 .5 5 .5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Airport(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 6.818V8.5l-6.5-1l-.318 4.773L11 14v1l-3.5-.682L4 15v-1l2.818-1.727L6.5 7.5L0 8.5V6.818L6.5 4.5v-3s0-1.5 1-1.5s1 1.5 1 1.5v2.818z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirportEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M6.5 6.4V6l4.5.5V5L6.5 3.2V1.5C6.5.5 6 0 5.5 0s-1 .5-1 1.5v1.7L0 5v1.4L4.5 6v3.3L3 10v1l2.5-.5L8 11v-1l-1.5-.8V6.4z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlcoholShop(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 4h-4v3.5a2 2 0 0 0 1.5 1.93V13H11a.5.5 0 0 0 0 1h2a.5.5 0 0 0 0-1h-.5V9.43A2 2 0 0 0 14 7.5v-.06zm-1 3.5a1 1 0 1 1-2 0V5h2zm-7.5-5V2a.5.5 0 0 0 0-1V.5A.5.5 0 0 0 5 0H4a.5.5 0 0 0-.5.5V1a.5.5 0 0 0 0 1v.5C3.5 3.93 1 5.57 1 7v6a1 1 0 0 0 1 1h5a1.1 1.1 0 0 0 1-1V7c0-1.35-2.5-3.15-2.5-4.5m-1 9.5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlcoholShopEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M7 4v2.5a1.5 1.5 0 0 0 1 1.41V10h-.5a.5.5 0 0 0 0 1h2a.5.5 0 0 0 0-1H9V7.91a1.5 1.5 0 0 0 1-1.41V4H7zm2.5 2.5a1 1 0 1 1-2 0v-2h2v2zM4.21 2.85V2.5a.355.355 0 1 0 0-.71v-.35a.35.35 0 0 0-.36-.35h-.71a.35.35 0 0 0-.36.34v.36a.355.355 0 1 0 0 .71v.35C2.79 3.87 1 5 1 6v4.25a.7.7 0 0 0 .71.71h3.58a.79.79 0 0 0 .71-.67V6c0-.91-1.79-2.19-1.79-3.15zM3.5 9a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AmericanFootball(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.53 3C3.09 3 1 7.5 1 7.5S3.09 12 7.53 12S14 7.5 14 7.5S12 3 7.53 3M11 7v1.5a.5.5 0 0 1-1 0V8H8v.5a.5.5 0 0 1-1 0V8H5v.5a.5.5 0 0 1-1 0v-2a.5.5 0 0 1 1 0V7h2v-.5a.5.5 0 0 1 1 0V7h2v-.5a.5.5 0 0 1 1 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AmericanFootballEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.53 2C2.47 2 1 5.5 1 5.5S2.47 9 5.53 9S10 5.5 10 5.5S8.6 2 5.53 2zM7 6H4a.5.5 0 0 1 0-1h3a.5.5 0 0 1 0 1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AmusementPark(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 7a5.992 5.992 0 0 1-2.546 4.907L12 14H3l1.046-2.093A6 6 0 1 1 13.5 7m-1.202.125H8.995c-.027.32-.154.612-.35.844l2.336 2.336a4.783 4.783 0 0 0 1.317-3.18m-1.493 3.356L8.469 8.145a1.498 1.498 0 0 1-.469.27v3.36a4.784 4.784 0 0 0 2.805-1.294m-4.8-3.356H2.702a4.783 4.783 0 0 0 1.317 3.18l2.336-2.336a1.494 1.494 0 0 1-.35-.844m-1.81 3.356A4.784 4.784 0 0 0 7 11.774v-3.36a1.498 1.498 0 0 1-.469-.269zm8.103-3.606a4.783 4.783 0 0 0-1.317-3.18L8.645 6.031c.196.232.323.524.35.844zm-1.493-3.356a4.783 4.783 0 0 0-3.18-1.317v3.303c.32.027.612.154.844.35zm-3.43 1.986V2.202a4.783 4.783 0 0 0-3.18 1.317l2.336 2.336c.232-.196.524-.323.844-.35m-1.02.526L4.019 3.695a4.783 4.783 0 0 0-1.317 3.18h3.303c.027-.32.154-.612.35-.844"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AmusementParkEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.5 1C3.02 1 1 3.02 1 5.5c0 1.792 1.063 3.332 2.584 4.055L2.5 11h6L7.416 9.555C8.937 8.832 10 7.292 10 5.5C10 3.02 7.98 1 5.5 1zm-.125 1.012v1.994a1.5 1.5 0 0 0-.844.35L3.12 2.942a3.475 3.475 0 0 1 2.256-.931zm.25 0c.874.03 1.66.377 2.256.931L6.469 4.356l-.002-.002a1.5 1.5 0 0 0-.842-.348V2.012zM2.943 3.119l1.412 1.412s0 .002-.002.002a1.5 1.5 0 0 0-.347.842H2.012c.03-.874.377-1.66.931-2.256zm5.114 0c.554.596.9 1.382.931 2.256H6.994a1.5 1.5 0 0 0-.35-.844L8.058 3.12zM2.012 5.625h1.994a1.5 1.5 0 0 0 .35.844L2.942 7.88a3.475 3.475 0 0 1-.931-2.256zm4.982 0h1.994a3.474 3.474 0 0 1-.931 2.256L6.644 6.469l.002-.002a1.5 1.5 0 0 0 .348-.842zm-2.463 1.02s.002 0 .002.002A1.5 1.5 0 0 0 5 6.911v2.04a3.47 3.47 0 0 1-1.88-.895L4.53 6.644zm1.938 0L7.88 8.056A3.47 3.47 0 0 1 6 8.95V6.912a1.5 1.5 0 0 0 .469-.268z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AnimalShelter(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.528.895L9 2h1L11.473.896a.33.33 0 0 1 .527.263V4.5c0 .722-.522 1.184-1 1.573V10l.8 2.4l.706.353a.894.894 0 0 1 .494.8c0 .247-.2.447-.447.447H4.25C3.017 14 2 13 2 11.713c0-2.17.453-3.783.745-4.868c.146-.543.29-.95.393-1.231c.058-.157.115-.303.192-.449c.449-.855 1.744-.21 1.342.668a6.311 6.311 0 0 0-.133.317c-.091.232-.214.593-.346 1.084c-.246.915-.519 2.287-.635 4.133c.368-2.06 2.422-4.128 4.077-5.609c-.345-.326-.634-.72-.634-1.258V1.158c0-.27.31-.425.527-.263"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Aquarium(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.135 8.577C1.989 6.147 4.721 2 10 2c2.652 0 4 1 4 1s-2.313 0-3.38.445a.379.379 0 0 0-.187.169c-.19.346-.68 1.433-.11 2.581c.094.19.313.27.515.208c.51-.157 1.492-.506 2.238-1.046c.221-.16.539-.1.604.166c.26 1.067.26 2.887 0 3.954c-.065.265-.383.326-.604.166c-.746-.54-1.729-.889-2.238-1.046a.427.427 0 0 0-.515.208c-.57 1.148-.08 2.235.11 2.581a.379.379 0 0 0 .187.169C11.687 12 14 12 14 12c-.5.333-2.2 1-5 1c-4.483 0-6.973-2.657-7.74-3.64a.821.821 0 0 1-.125-.783M5 9a1.25 1.25 0 1 0 0-2.5A1.25 1.25 0 0 0 5 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AquariumEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M8 1c-.876 0-1.85.092-3.004.527C3.843 1.962 2.848 2.657 2 3.5c-.852.847-2 2.5-2 3s1.135 1.943 2.678 2.621c1.542.678 2.39.798 3.283.895c.797.086 1.942-.027 2.885-.233C9.592 9.621 10.994 9.31 11 9c0 0-2.756-.063-3-.5c-.249-.445-.25-1.586 0-2c.258-.428 2.5 1 2.5 1c.644.258.644-4.258 0-4c0 0-2.277 1.447-2.5 1c-.25-.5-.25-1.5 0-2c.223-.447 3-.5 3-.5c0-.5-2.124-1-3-1zM3.514 4.502a1.014 1.014 0 1 1 0 2.028a1.014 1.014 0 0 1 0-2.028z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arrow(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.293 2.293a1 1 0 0 1 1.414 0l4.5 4.5a1 1 0 0 1 0 1.414l-4.5 4.5a1 1 0 0 1-1.414-1.414L11 8.5H1.5a1 1 0 0 1 0-2H11L8.293 3.707a1 1 0 0 1 0-1.414"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArtGallery(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.71 3L7.85.15a.5.5 0 0 0-.707-.003L7.14.15L4.29 3H1.5a.5.5 0 0 0-.5.5v9a.5.5 0 0 0 .5.5h12a.5.5 0 0 0 .5-.5v-9a.5.5 0 0 0-.5-.5zM7.5 1.21L9.29 3H5.71zM13 12H2V4h11zM5 7a1 1 0 1 1 0-2a1 1 0 0 1 0 2m7 4H4.5L6 8l1.25 2.5L9.5 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArtGalleryEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M8.21 3L5.85.65a.5.5 0 0 0-.707-.003L5.14.65L2.79 3H1.5a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5H8.21zM5.5 1.71L6.79 3H4.21L5.5 1.71zM9 9H2V4h7v5zM4.5 5.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0zM8 8H4l.75-1.5l.5 1L6.5 5L8 8z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Attraction(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 2c-.554 0-.752.505-1 1l-.5 1h-2C1.669 4 1 4.669 1 5.5v5c0 .831.669 1.5 1.5 1.5h10c.831 0 1.5-.669 1.5-1.5v-5c0-.831-.669-1.5-1.5-1.5h-2L10 3c-.25-.5-.446-1-1-1zM2.5 5a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1m5 0a3 3 0 1 1 0 6a3 3 0 0 1 0-6m0 1.5a1.5 1.5 0 0 0 0 3a1.5 1.5 0 0 0 0-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AttractionEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M4.5 1.5s-.5 0-.7.5l-.3.5H1c-.6 0-1 .4-1 1v5c0 .6.4 1 1 1h9c.6 0 1-.4 1-1v-5c0-.6-.4-1-1-1H7.5L7.2 2c-.2-.5-.7-.5-.7-.5h-2zm1 2C6.9 3.5 8 4.6 8 6S6.9 8.5 5.5 8.5S3 7.4 3 6s1.1-2.5 2.5-2.5zm0 1.5c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bakery(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.294 4.382L6 9.5a.979.979 0 0 0 1 1h1a.979.979 0 0 0 1-1l.706-5.118C9.706 3 7.5 3 7.5 3S5.291 3 5.294 4.382M3.5 5C2 5 2 6 2 6l1 4h1.5a.793.793 0 0 0 .794-.765L4.5 5Zm-2 2.5a1.533 1.533 0 0 0-1.059.412A1.366 1.366 0 0 0 0 8.794V11h.882A1.02 1.02 0 0 0 2 10Zm10-2.5C13 5 13 6 13 6l-1 4h-1.5a.793.793 0 0 1-.794-.765L10.5 5Zm2 2.5a1.533 1.533 0 0 1 1.059.412a1.366 1.366 0 0 1 .441.882V11h-.882A1.02 1.02 0 0 1 13 10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BakeryEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M4.5 2c-1 0-1 1-1 1L5 7.5s0 .5.5.5s.5-.5.5-.5L7.5 3s0-1-1-1h-2zM9 3.5l-2 4h1.5l1 1h.5c1 0 1-.9 1-.9V6.3L9 3.5zM0 6.3v1.2s.03 1.01 1 1c.97-.01.5 0 .5 0l1-1H4l-2-4l-2 2.8z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BakeryFifteen(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.294 4.382L6 9.5s0 1 1 1h1c1 0 1-1 1-1l.706-5.118C9.706 3 7.5 3 7.5 3S5.291 3 5.294 4.382zM3.5 5C2 5 2 6 2 6l1 4h1.5c.755 0 .794-.765.794-.765L4.5 5h-1zm-2 2.5s-.618-.03-1.059.412C0 8.352 0 8.794 0 8.794V11h.882C2 11 2 10 2 10l-.5-2.5z" fill="currentColor"/><path d="M11.5 5C13 5 13 6 13 6l-1 4h-1.5c-.755 0-.794-.765-.794-.765L10.5 5h1zm2 2.5s.618-.03 1.059.412c.441.44.441.882.441.882V11h-.882C13 11 13 10 13 10l.5-2.5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bank(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 3c-.554 0-1 .446-1 1v7c0 .554.446 1 1 1h13c.554 0 1-.446 1-1V4c0-.554-.446-1-1-1zm0 1h1.5a.5.5 0 1 1-.5.5l-.5.5a.5.5 0 1 1-.5.5zm6.5 0C8.88 4 10 5.567 10 7.5S8.88 11 7.5 11S5 9.433 5 7.5S6.12 4 7.5 4m5 0H14v1.5a.5.5 0 1 1-.5-.5l-.5-.5a.5.5 0 1 1-.5-.5m-5 1.5c-.323 0-.534.109-.682.25h1.364c-.148-.141-.359-.25-.682-.25m-.875.5c-.045.091-.062.171-.08.25h1.91c-.018-.079-.034-.159-.08-.25zm-.125.5v.25h2V6.5zm0 .5v.25h2V7zm0 .5v.25h2V7.5zm0 .5l-.25.25h2L8.5 8zm-.5.5s.035.102.102.25h2.273L8 8.5zM1.5 9a.5.5 0 0 1 0 1l.5.5a.5.5 0 1 1 .5.5H1V9.5a.5.5 0 0 1 .5-.5m4.738 0c.046.086.076.159.137.25h2.268c.066-.138.107-.25.107-.25zM13.5 9a.5.5 0 0 1 .5.5V11h-1.5a.5.5 0 1 1 .5-.5l.5-.5a.5.5 0 0 1 0-1m-6.934.5c.079.091.165.176.26.25h1.42c.1-.077.188-.162.254-.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BankEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M1 2C0 2 0 3 0 3v5c0 1 1 1 1 1h9c1 0 1-1 1-1V3s0-1-1-1H1zm0 1h1.5a.5.5 0 1 1-.5.5l-.5.5a.5.5 0 1 1-.5.5V3zm4.5 0c1.105 0 2 1.12 2 2.5S6.605 8 5.5 8s-2-1.12-2-2.5s.895-2.5 2-2.5zm3 0H10v1.5a.5.5 0 1 1-.5-.5L9 3.5a.5.5 0 1 1-.5-.5zm-7 3a.5.5 0 0 1 0 1l.5.5a.5.5 0 1 1 .5.5H1V6.5a.5.5 0 0 1 .5-.5zm8 0a.5.5 0 0 1 .5.5V8H8.5a.5.5 0 1 1 .5-.5l.5-.5a.5.5 0 0 1 0-1z" fill="currentColor"/><path d="M4.902 4.25a1.157 1.157 0 0 0-.2.25h1.597a1.158 1.158 0 0 0-.201-.25H4.902zm-.316.5c-.029.08-.05.164-.065.25h1.96a1.479 1.479 0 0 0-.065-.25h-1.83zm-.086.5c0 .084.006.168.02.25h1.959c.013-.082.02-.166.021-.25h-2zm.086.5c.031.089.07.173.117.25H6.3c.046-.077.084-.161.115-.25H4.586zm.316.5l.034.031L4.5 6.5h2l-.436-.219l.034-.031H4.902zm-.32.5c.029.079.056.16.11.25h1.626c.053-.091.078-.171.106-.25H4.582zm.31.5c.147.141.341.25.608.25c.274 0 .47-.108.617-.25H4.893z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BankJp(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.28 3.376l-2.9 3.625h1.87a.75.75 0 0 1 0 1.5H8.5v1h2.75a.75.75 0 0 1 0 1.5H8.5v1.25a1 1 0 0 1-2 0v-1.25H3.75a.75.75 0 1 1 0-1.5H6.5v-1H3.75a.75.75 0 0 1 0-1.5h1.87l-2.9-3.625a1 1 0 0 1 1.557-1.254l.004.004L7.5 6.15l3.22-4.024a1 1 0 0 1 1.565 1.245Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BankJpEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M8.826 2.48L6.892 4.75H8a.5.5 0 0 1 0 1H6.25v1H8a.5.5 0 0 1 0 1H6.25V9a.75.75 0 0 1-1.5 0V7.75H3a.5.5 0 0 1 0-1h1.75v-.999H3a.5.5 0 0 1 0-1h1.107l-1.933-2.27a.75.75 0 0 1 1.152-.961L5.5 4.078L7.674 1.52a.75.75 0 1 1 1.152.96z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bar(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 1c-2 0-7 .25-6.5.75L7 8v4c0 1-3 .5-3 2h7c0-1.5-3-1-3-2V8l6-6.25c.5-.5-4.5-.75-6.5-.75m0 1c2.5 0 4.75.25 4.75.25L11.5 3h-8l-.75-.75S5 2 7.5 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.488 1C4.976 1 .5 1 1 1.5L5 6v2.5C5 9 2.5 9 2.5 10h6C8.5 9 6 9 6 8.5V6l4-4.5C10.5 1 6 1 5.488 1zM2.5 2h6l-1 1h-4l-1-1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Barrier(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 2H2a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h1v2.5a.5.5 0 0 0 1 0V12h7v.5a.5.5 0 0 0 1 0V10h1a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1m0 1v2l-2-2zM9.5 3L13 6.5V9L7 3zm-4 6L2 5.5V3l6 6zM2 9V7l2 2zm9 2H4v-1h7zm-.207-2H9.5l-6-6h2l6 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarrierEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.5 2h-8a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5H2v2.5a.5.5 0 0 0 1 0V9h5v.5a.5.5 0 0 0 1 0V7h.5a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5zM2 3h1.5l3 3h-2L2 3.5zm0 3V4.5L3.5 6zm1 2V7h5v1zm6-2H7.5l-3-3h2L9 5.5zm0-1.5L7.5 3H9z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Baseball(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 3.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0M6 .28A.28.28 0 0 0 5.72 0a.49.49 0 0 0-.25.16L3 4.59a.48.48 0 0 0 0 .13c0 .155.125.28.28.28a.49.49 0 0 0 .26-.16L6 .41a.472.472 0 0 0 0-.13m5.9 13.92L9 6.39A.49.49 0 0 0 8.52 6h-5a.5.5 0 0 0 0 1H6l1.45 2.51l-4.27 4.61a.49.49 0 0 0-.18.38a.5.5 0 0 0 .5.5a.49.49 0 0 0 .33-.13l4.45-4.15l2.76 4a.51.51 0 0 0 .44.26c.28 0 .51-.22.52-.5a.5.5 0 0 0-.1-.28"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BaseballEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M7 3a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm2.85 7.24l-3-4.85a.86.86 0 0 0-.5-.39H2.5a.5.5 0 0 0 0 1H5l.92 1.09l-2.73 3a.49.49 0 0 0-.19.41a.5.5 0 0 0 .5.5a.49.49 0 0 0 .33-.13l3-2.71L9 10.81a.49.49 0 0 0 .38.19a.5.5 0 0 0 .5-.5a.49.49 0 0 0-.03-.26zM4 .28A.28.28 0 0 0 3.72 0a.49.49 0 0 0-.25.16L2 4.59a.48.48 0 0 0 0 .13c0 .155.125.28.28.28a.49.49 0 0 0 .26-.16L4 .41a.472.472 0 0 0 0-.13z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Basketball(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.66 7H2.03c.08-.98.43-1.947 1.01-2.74A8.9 8.9 0 0 1 3.66 7m1.075 0H7V2c-1.21.106-2.275.613-3.165 1.44a9.442 9.442 0 0 1 .9 3.56M8 2v5h2.265a9.442 9.442 0 0 1 .9-3.56C10.275 2.613 9.21 2.106 8 2m3.96 2.26A8.9 8.9 0 0 0 11.34 7h1.63a5.425 5.425 0 0 0-1.01-2.74M14 8.5a.5.5 0 0 1-.5.498h-1.78l-.76 1.52v2.38a.393.393 0 0 1-.24.35a.586.586 0 0 1-.14.02a.327.327 0 0 1-.26-.11l-1.02-1.01l-1.59 1.06a.398.398 0 0 1-.42 0l-1.59-1.06l-1.02 1.01a.371.371 0 0 1-.4.09a.393.393 0 0 1-.24-.35v-2.38l-.76-1.52H1.5A.5.5 0 1 1 1.5 8h12a.5.5 0 0 1 .5.5m-7.14.498H4.12l.63 1.27a.317.317 0 0 1 .04.16v1.56l.56-.55zm1.98 2.55L7.5 9.388l-1.34 2.16l1.34.9zm2.05-2.55H8.14l1.51 2.44l.56.55v-1.56a.317.317 0 0 1 .04-.16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BasketballEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M2.63 3.51a5.155 5.155 0 0 1 .34 1.48h-.93a3.259 3.259 0 0 1 .59-1.48zm.55-.63a6.202 6.202 0 0 1 .55 2.11h1.39V2.01a3.426 3.426 0 0 0-1.94.87zm2.7-.87v2.98h1.39a6.202 6.202 0 0 1 .55-2.11a3.426 3.426 0 0 0-1.94-.87zm2.49 1.5a5.155 5.155 0 0 0-.34 1.48h.93a3.259 3.259 0 0 0-.59-1.48zM9.9 6.5a.495.495 0 0 1-.49.5h-.97l-.53 1.05v1.7a.218.218 0 0 1-.14.2a.236.236 0 0 1-.08.02a.199.199 0 0 1-.16-.07l-.75-.74l-1.16.77a.21.21 0 0 1-.24 0l-1.16-.77l-.75.74a.222.222 0 0 1-.38-.15v-1.7L2.57 7H1.6a.5.5 0 1 1 0-1h7.8a.495.495 0 0 1 .5.49zM3.53 8v1.22l.5-.5c.01 0 .01-.01.02-.01L5.1 7H3.06l.45.9a.356.356 0 0 1 .02.1zm2.97.82l-1-1.63l-1 1.63l1 .67zM7.94 7H5.9l1.05 1.71c.01 0 .01.01.02.01l.5.5V8a.356.356 0 0 1 .02-.1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bbq(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.32 7.655A4.018 4.018 0 0 0 12 4H3a4.01 4.01 0 0 0 2.63 3.631l-.707 1.413A2.494 2.494 0 1 0 6.49 13h4.223l.335.717a.5.5 0 0 0 .905-.428zM4.5 12.75a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5M6.95 12a2.49 2.49 0 0 0-1.091-2.594L6.6 7.923a4.91 4.91 0 0 0 1.745 0l1.9 4.07zM4.75 1a.25.25 0 0 0-.25.25c0 .5-.5.5-.5.5c-1 0-1 1-1 1a.25.25 0 1 0 .5 0c0-.5.5-.5.5-.5c1 0 1-1 1-1A.25.25 0 0 0 4.75 1m2 0a.25.25 0 0 0-.25.25c0 .5-.5.5-.5.5c-1 0-1 1-1 1a.25.25 0 1 0 .5 0c0-.5.5-.5.5-.5c1 0 1-1 1-1A.25.25 0 0 0 6.75 1m2 0a.25.25 0 0 0-.25.25c0 .5-.5.5-.5.5c-1 0-1 1-1 1a.25.25 0 1 0 .5 0c0-.5.5-.5.5-.5c1 0 1-1 1-1A.25.25 0 0 0 8.75 1m2 0a.25.25 0 0 0-.25.25c0 .5-.5.5-.5.5c-1 0-1 1-1 1a.25.25 0 1 0 .5 0c0-.5.5-.5.5-.5c1 0 1-1 1-1a.25.25 0 0 0-.25-.25m2 0a.25.25 0 0 0-.25.25c0 .5-.5.5-.5.5c-1 0-1 1-1 1a.25.25 0 1 0 .5 0c0-.5.5-.5.5-.5c1 0 1-1 1-1a.25.25 0 0 0-.25-.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BbqEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M4 1.75s0-1 1-1c0 0 .5 0 .5-.5a.25.25 0 1 1 .5 0s0 1-1 1c0 0-.5 0-.5.5a.25.25 0 1 1-.5 0zM2.25 2a.25.25 0 0 0 .25-.25c0-.5.5-.5.5-.5c1 0 1-1 1-1a.25.25 0 1 0-.5 0c0 .5-.5.5-.5.5c-1 0-1 1-1 1c0 .138.112.25.25.25zm4 0a.25.25 0 0 0 .25-.25c0-.5.5-.5.5-.5c1 0 1-1 1-1a.25.25 0 1 0-.5 0c0 .5-.5.5-.5.5c-1 0-1 1-1 1c0 .138.112.25.25.25zm3.5-2a.25.25 0 0 0-.25.25c0 .5-.5.5-.5.5c-1 0-1 1-1 1a.25.25 0 1 0 .5 0c0-.5.5-.5.5-.5c1 0 1-1 1-1A.25.25 0 0 0 9.75 0zM6.675 5.865l-.001.001l2.3 4.782v.009a.242.242 0 0 1-.12.32a.247.247 0 0 1-.328-.12L7.845 9.5H4a1.5 1.5 0 1 1-.91-1.379c.053.02.105.045.155.072L4.278 5.85A3.038 3.038 0 0 1 2 3h7s-.002 2.282-2.325 2.865zM3.25 9.5a.75.75 0 1 0-1.5 0a.75.75 0 0 0 1.5 0zm2.923-3.544C5.963 5.983 5.742 6 5.5 6c-.285 0-.713-.047-.713-.047l-1.144 2.58c.123.136.218.295.279.467h3.673L6.173 5.956z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Beach(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.36 1.67l-.01 4.02a4.452 4.452 0 0 0-1.1-.11c-.37.1-.74.63-1.1.76a4.202 4.202 0 0 1 2.21-4.67m2.41-.64L9.8 4.48a3.183 3.183 0 0 1 .84-.61c.36-.1.94.17 1.34.11a4.202 4.202 0 0 0-4.21-2.95M1 13h13c-.66-.66-2.64-1.11-4.34-1.33l-1.87-7c.52-.05 1.15.03 1.53 0l-2.11-3.6H7.2a6.174 6.174 0 0 0-.7.14a4.38 4.38 0 0 0-.64.22l-.01 4.15c.35-.17.84-.54 1.3-.74l1.8 6.74c-.58-.05-1.09-.08-1.45-.08C6.03 11.5 2 12 1 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BeachEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M3.5 1.65v2.82a1.703 1.703 0 0 0-.58-.04c-.28.07-.56.47-.83.57A3.126 3.126 0 0 1 3.5 1.65zm2.31-.62l1.45 2.42a1.216 1.216 0 0 1 .45-.31c.27-.07.7.13 1 .09a3.106 3.106 0 0 0-2.9-2.2zM.984 10h8.99A4.84 4.84 0 0 0 6.9 8.68L5.57 3.74c.46-.04 1.02.06 1.27-.01L5.23 1.04a3.525 3.525 0 0 0-.62.11a2.96 2.96 0 0 0-.61.23v3.1c.25-.08.68-.42 1.08-.6l1.27 4.75q-.39-.03-.84-.03C1.99 8.6.984 10 .984 10z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Beer(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5V2s-1-1-4.5-1S3 2 3 2v3a9.27 9.27 0 0 0 1 4a5.63 5.63 0 0 1 0 4.5s0 1 3.5 1s3.5-1 3.5-1A5.63 5.63 0 0 1 11 9a9.27 9.27 0 0 0 1-4m-4.5 8.5a7.368 7.368 0 0 1-2.36-.28c.203-.722.304-1.47.3-2.22h4.12c-.004.75.097 1.498.3 2.22a7.368 7.368 0 0 1-2.36.28m0-8.5A10.65 10.65 0 0 1 4 4.5v-2A10.74 10.74 0 0 1 7.5 2a10.74 10.74 0 0 1 3.5.5v2c-1.13.36-2.314.53-3.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BeerEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.5 1c-2.3 0-3 .66-3 .66v2a6.6 6.6 0 0 0 .66 2.65a4.13 4.13 0 0 1 0 3s0 .66 2.32.66s2.32-.66 2.32-.66a4.13 4.13 0 0 1 0-3a6.6 6.6 0 0 0 .66-2.65v-2S7.8 1 5.5 1zm0 8.28a4.77 4.77 0 0 1-1.56-.18c.133-.479.2-.973.2-1.47h2.72c-.014.22-.014.44 0 .66c.034.274.087.544.16.81a4.77 4.77 0 0 1-1.52.19v-.01zm2.32-6a8.24 8.24 0 0 1-4.63 0L3.18 2a8.28 8.28 0 0 1 4.64 0s.03 1.33 0 1.33v-.05z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bicycle(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 2c-.676-.01-.676 1.01 0 1H9v1.266L6.197 6.6L5.223 4H5.5c.676.01.676-1.01 0-1h-2c-.676-.01-.676 1.01 0 1h.652l.891 2.375A3.45 3.45 0 0 0 3.5 6C1.573 6 0 7.573 0 9.5S1.573 13 3.5 13S7 11.427 7 9.5c0-.67-.2-1.291-.53-1.824l2.821-2.35l.463 1.16C8.71 7.094 8 8.211 8 9.5c0 1.927 1.573 3.5 3.5 3.5S15 11.427 15 9.5S13.427 6 11.5 6c-.283 0-.554.043-.818.107L10 4.402V2.5a.5.5 0 0 0-.5-.5zm-4 5a2.48 2.48 0 0 1 1.555.553L3.18 9.115c-.511.427.128 1.195.64.77l1.875-1.563c.188.352.305.75.305 1.178C6 10.887 4.887 12 3.5 12S1 10.887 1 9.5S2.113 7 3.5 7m8 0C12.887 7 14 8.113 14 9.5S12.887 12 11.5 12S9 10.887 9 9.5a2.49 2.49 0 0 1 1.125-2.088l.91 2.274c.246.623 1.18.25.93-.372l-.908-2.27C11.2 7.02 11.348 7 11.5 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BicycleEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M6.5 1.5c-.676-.01-.676 1.01 0 1H7v.711l-2.748 1.57L3.738 3.5h.715c.676.01.676-1.01 0-1H2.426c-.676-.01-.676 1.01 0 1h.234l.432 1.078A2.461 2.461 0 0 0 2.5 4.5C1.125 4.5 0 5.625 0 7s1.125 2.5 2.5 2.5S5 8.375 5 7c0-.471-.14-.908-.37-1.285l2.472-1.412l.402.408A2.507 2.507 0 0 0 6 7.001c0 1.374 1.125 2.5 2.5 2.5S11 8.374 11 7c0-1.297-1.003-2.358-2.27-2.478L8 3.794V2a.5.5 0 0 0-.5-.5h-1zm-4 4C3.334 5.5 4 6.166 4 7s-.666 1.5-1.5 1.5S1 7.834 1 7s.666-1.5 1.5-1.5zm5.955.004h.002a.45.45 0 0 0 .09 0A1.49 1.49 0 0 1 10 7c0 .834-.666 1.5-1.5 1.5S7 7.835 7 7a1.49 1.49 0 0 1 1.455-1.496z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BicycleShare(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 1a1 1 0 1 0 0 2a1 1 0 0 0 0-2M8.145 2.994a.5.5 0 0 0-.348.143l-2.64 2.5a.5.5 0 0 0 .042.763L7 7.75v2.75c-.01.676 1.01.676 1 0v-3a.5.5 0 0 0-.2-.4l-.767-.577l1.818-1.72l.749.998A.5.5 0 0 0 10 6h1.5c.676.01.676-1.01 0-1h-1.25L9.5 4l-.6-.8a.5.5 0 0 0-.384-.206zM3 7a3 3 0 1 0 0 6a3 3 0 0 0 0-6m9 0a3 3 0 1 0 0 6a3 3 0 0 0 0-6M3 8a2 2 0 1 1 0 4a2 2 0 0 1 0-4m9 0a2 2 0 1 1 0 4a2 2 0 0 1 0-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BicycleShareEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M7 1a1 1 0 1 1 1 1a1 1 0 0 1-1-1zM1.973 7.1a1.5 1.5 0 0 1 1.654.408a.5.5 0 0 0 .749-.663A2.519 2.519 0 0 0 2.363 6a2.5 2.5 0 1 0 2.008 4.158a.5.5 0 0 0-.748-.658a1.5 1.5 0 1 1-1.65-2.4zM7.1 4.8a.5.5 0 0 0 .4.2h2a.5.5 0 1 0 .014-1H7.75L6.4 2.2a.5.5 0 0 0-.386-.2a.506.506 0 0 0-.368.147l-2 2a.5.5 0 0 0 0 .707L5 6.207V8.5a.5.5 0 0 0 1 .015V6a.505.505 0 0 0-.144-.353L5.1 4.895l1.165-1.2zm2.287 1.362A2.526 2.526 0 0 0 8.643 6a2.525 2.525 0 0 0-2.014.838a.5.5 0 0 0 .75.664a1.5 1.5 0 1 1 0 1.992a.5.5 0 0 0-.78.627a.596.596 0 0 0 .034.037a2.5 2.5 0 1 0 2.752-4z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BloodBank(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.2 6.1L7.5 1L3.8 6.1c-.5.7-.8 1.6-.8 2.5C3 11 5 13 7.5 13S12 11 12 8.6c0-.9-.3-1.8-.8-2.5M10 9H8v2H7V9H5V8h2V6h1v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BloodBankEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M8.405 4.644h.001L5.5 1L2.593 4.644h.002A3.37 3.37 0 0 0 2 6.56a3.464 3.464 0 0 0 3.5 3.43A3.464 3.464 0 0 0 9 6.558a3.37 3.37 0 0 0-.595-1.915zM8 7H6v2H5V7H3V6h2V4h1v2h2v1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BowlingAlley(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.922 3.166c0 1.006-.335 1.857-.693 2.77c-.414 1.052-.86 2.185-.86 3.73c0 2.166.687 4.333 1.409 4.333h1.444c.723 0 1.409-2.166 1.409-4.333c0-1.545-.446-2.678-.86-3.73c-.358-.912-.693-1.764-.693-2.77c0-.722.144-.723.144-1.445A.743.743 0 0 0 11.5.999a.743.743 0 0 0-.722.722c0 .723.144.723.144 1.445m-2.552 6.5c0-1.234.254-2.236.55-3.091a4.5 4.5 0 1 0-.17 6.038c-.237-.865-.38-1.91-.38-2.947M5.24 7.12a.62.62 0 1 1-1.24 0a.62.62 0 0 1 1.24 0m-1 2a.62.62 0 1 1-1.24 0a.62.62 0 0 1 1.24 0m2 0a.62.62 0 1 1-1.24 0a.62.62 0 0 1 1.24 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BowlingAlleyEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M6.02 3.94c.04-.1.13-.4.17-.49A2.825 2.825 0 0 0 4.5 3a3.5 3.5 0 0 0 0 7a2.88 2.88 0 0 0 1.54-.4a8.597 8.597 0 0 1-.4-2.61a12.759 12.759 0 0 1 .38-3.05zM2.5 6.98a.48.48 0 1 1 .48-.48a.487.487 0 0 1-.48.48zm1-2a.48.48 0 1 1 .48-.48a.487.487 0 0 1-.48.48zm1 2a.48.48 0 1 1 .48-.48a.487.487 0 0 1-.48.48zM8.098 2.5c0-.5-.1-.5-.1-1c0-.782.5-.7.5-.7s.5-.082.5.7c0 .5-.1.5-.1 1c0 1.5 1.075 2.503 1.075 4.5c0 1.5-.475 3-.975 3h-1c-.5 0-.975-1.5-.975-3c0-1.997 1.075-3 1.075-4.5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bridge(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 6.842V11h2v-1a2.07 2.07 0 0 1 2-2a2.07 2.07 0 0 1 2 2v1h3v-1a2.07 2.07 0 0 1 2-2a2.07 2.07 0 0 1 2 2v1h2V6.842a9.245 9.245 0 0 0-15 0m9.5-3.108a8.802 8.802 0 0 1 1.5.5V7H9.5zm-.5-.1V7H7.5V3.506c.503 0 1.005.043 1.5.128m-5.5.843c.48-.248.982-.451 1.5-.606V7H3.5zM5.5 7V3.73A8.732 8.732 0 0 1 7 3.515V7zM3 4.744V7H.5A8.777 8.777 0 0 1 3 4.744M11.5 7V4.5l-.053-.053A8.674 8.674 0 0 1 14.5 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BridgeEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M0 5.835V9h1c0-1.105.896-2.006 2-2.006c1.103 0 1.998.903 2 2.006h1c0-1.105.895-2.006 1.999-2.006A2.01 2.01 0 0 1 10 9h1V5.835c-2.917-4.029-8.611-3.523-11 0zM2 6H.5A6.276 6.276 0 0 1 2 4.582V6zm2 0H2.5V4.263A6.593 6.593 0 0 1 4 3.682V6zm2 0H4.5V3.578c.496-.081 1-.102 1.5-.062V6zm2 0H6.5V3.579c.517.084 1.02.232 1.5.441V6zm.5 0V4.264A6.247 6.247 0 0 1 10.5 6h-2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Building(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 2v11h5v-3h3v3h1V2zm4 10H4v-2h3zm0-3H4V7h3zm0-3H4V4h3zm4 3H8V7h3zm0-3H8V4h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingAltOne(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 13.5v-9c0-.3-.2-.5-.5-.5H9V1L5 2.1v11.4H2v.5h11v-.5zm-4 0V3h1v10.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingAltOneEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M8 9.5v-5c0-.3-.2-.5-.5-.5H6V1L3 2v7.5H2v.5h7v-.5H8zm-3 0H4V3h1v6.5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M2 1v9h4V6h2v4h1V1H2zm3 8H3V6h2v3zm0-4H3V3h2v2zm3 0H6V3h2v2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bus(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 3c0-1.1.9-2 2-2h7c1.1 0 2 .9 2 2v8c0 1-1 1-1 1v1c0 .55-.45 1-1 1s-1-.45-1-1v-1H5v1c0 .55-.45 1-1 1s-1-.45-1-1v-1c-1 0-1-1-1-1zm1.5 1c-.28 0-.5.22-.5.5v3c0 .28.22.5.5.5h8c.28 0 .5-.22.5-.5v-3c0-.28-.22-.5-.5-.5zM4 9c-.55 0-1 .45-1 1s.45 1 1 1s1-.45 1-1s-.45-1-1-1m7 0c-.55 0-1 .45-1 1s.45 1 1 1s1-.45 1-1s-.45-1-1-1M4 2.5c0 .28.22.5.5.5h6c.28 0 .5-.22.5-.5s-.22-.5-.5-.5h-6c-.28 0-.5.22-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BusEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M3 0C2 0 1 .531 1 2v7.5s0 .5.5.5l.5.016v.484s0 .5.5.5H3s.5 0 .5-.5v-.484l4-.016v.5s0 .5.5.5h.5c.5 0 .5-.5.5-.5v-.484L9.5 10s.5 0 .5-.5V2c0-1.5-1-2-2-2H3zm.75 1h3.5a.25.25 0 1 1 0 .5h-3.5a.25.25 0 1 1 0-.5zM3 2h5c1 0 1 1 1 1v2s0 1-1 1H3C2 6 2 5 2 5V3s0-1 1-1zm-.25 5.5a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5zm5.5 0a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cafe(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5h-2V3H2v4a4 4 0 0 0 7.45 2H12a2 2 0 1 0 0-4m0 3H9.86A4 4 0 0 0 10 7V6h2a1 1 0 1 1 0 2m-2 4.5a.5.5 0 0 1-.5.5h-7a.5.5 0 0 1 0-1h7a.5.5 0 0 1 .5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CafeEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M7 9.5a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1 0-1h4a.5.5 0 0 1 .5.5zM8 3H7V2H2v4a2.5 2.5 0 0 0 4.79 1H8a2 2 0 1 0 0-4zm0 3H7V4h1a1 1 0 1 1 0 2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Campsite(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 10v1a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1v-1a1 1 0 0 1 1-1h.25l4.782-7.742a.55.55 0 0 1 .936 0L12.75 9H13a1 1 0 0 1 1 1m-3.5-1l-3-5l-3 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CampsiteEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.92 2.19a.5.5 0 0 0-.85 0L1.5 8h-1a.42.42 0 0 0-.5.39V9.5a.49.49 0 0 0 .5.5h10a.49.49 0 0 0 .5-.5V8.39a.42.42 0 0 0-.5-.39h-1L5.92 2.19zM5.5 3l3 5h-6l3-5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Car(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.84 6.852L12.6 5.7l-1.1-2.2a1.05 1.05 0 0 0-.9-.5H4.4a1.05 1.05 0 0 0-.9.5L2.4 5.7L1.16 6.852a.5.5 0 0 0-.16.367V11.5a.5.5 0 0 0 .5.5h2c.2 0 .5-.2.5-.4V11h7v.5c0 .2.2.5.4.5h2.1a.5.5 0 0 0 .5-.5V7.219a.5.5 0 0 0-.16-.367M4.5 4h6l1 2h-8ZM5 8.6c0 .2-.3.4-.5.4H2.4c-.2 0-.4-.3-.4-.5V7.4c.1-.3.3-.5.6-.4l2 .4c.2 0 .4.3.4.5Zm8-.1c0 .2-.2.5-.4.5h-2.1c-.2 0-.5-.2-.5-.4v-.7c0-.2.2-.5.4-.5l2-.4c.3-.1.5.1.6.4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9 4l-.89-2.66A.5.5 0 0 0 7.64 1H3.36a.5.5 0 0 0-.47.34L2 4a1 1 0 0 0-1 1v3h1v1a1 1 0 1 0 2 0V8h3v1a1 1 0 1 0 2 0V8h1V5a1 1 0 0 0-1-1zM3 7a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm0-3l.62-2h3.76L8 4H3zm5 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarFifteen(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M14 7a1.5 1.5 0 0 0-1.15-1.45l-1.39-3.24A.5.5 0 0 0 11 2H4a.5.5 0 0 0-.44.28L2.15 5.54A1.5 1.5 0 0 0 1 7v3.5h1v1a1 1 0 1 0 2 0v-1h7v1a1 1 0 1 0 2 0v-1h1V7zM4.3 3h6.4l1.05 2.5h-8.5L4.3 3zM3 9a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm9 0a1 1 0 1 1 0-2a1 1 0 0 1 0 2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarRental(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 1a1 1 0 0 0-1 1H2v1l1 1l1-1l1 1l1-1l1 1h2.5a1 1 0 0 0 1 1h2a.5.5 0 0 0 .5-.5v-3a.5.5 0 0 0-.5-.5zm.5 1.5a.5.5 0 0 1 1 0v1a.5.5 0 0 1-1 0zM2.146 9.354A.5.5 0 0 0 2 9.707V13.5a.5.5 0 0 0 .5.5H4a.5.5 0 0 0 .5-.5V13h6v.5a.5.5 0 0 0 .5.5h1.5a.5.5 0 0 0 .5-.5V9.707a.5.5 0 0 0-.146-.353L12 8.5l-1.354-2.257a.5.5 0 0 0-.43-.243H4.784a.5.5 0 0 0-.429.243L3 8.5zM11.134 9H3.866l1.2-2h4.868zM5.5 10.828v.372a.3.3 0 0 1-.3.3H3.3a.3.3 0 0 1-.3-.3v-.834a.3.3 0 0 1 .359-.294l1.82.364a.4.4 0 0 1 .321.392m6.5-.34v.712a.3.3 0 0 1-.3.3H9.8a.3.3 0 0 1-.3-.3v-.454a.3.3 0 0 1 .241-.294l1.78-.356a.4.4 0 0 1 .479.392"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarRentalEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.025 7H8.84l-.278-1.67a.988.988 0 0 0-.979-.83H3.417c-.488 0-.9.35-.979.83L2.16 7H1.98A.98.98 0 0 0 1 7.98V10h1a1 1 0 0 0 2 0h3a1 1 0 0 0 2 0h1V7.975A.975.975 0 0 0 9.025 7zM3.25 9a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5zm-.08-2l.247-1.5l4.158-.007L7.826 7H3.171zm4.58 2a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5z" fill="currentColor"/><path d="M9.497 0h-1.5A.997.997 0 0 0 7 .997V1H2.5l-1 1l1 1l1-1l1 1l1-1l1 1H7c.003.553.45 1 1.004 1h1.494A.502.502 0 0 0 10 3.498V.503A.503.503 0 0 0 9.497 0zM9 2.5a.5.5 0 0 1-1 0v-1a.5.5 0 1 1 1 0v1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarRentalFifteen(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M11.504 1H10c-.738 0-1.376.405-1.723 1H3.5l-1 1l1 1l1-1l1 1l1-1l1 1h.777c.347.595.985 1 1.723 1h1.498A.502.502 0 0 0 12 4.498V1.496A.496.496 0 0 0 11.504 1zM11 3.5a.5.5 0 0 1-1 0v-1a.5.5 0 1 1 1 0v1z" fill="currentColor"/><path d="M10.925 9.01l-.363-2.18A.988.988 0 0 0 9.583 6H5.417a.989.989 0 0 0-.978.83l-.364 2.182A1.126 1.126 0 0 0 3 10.132V13h1a1 1 0 0 0 2 0h3a1 1 0 0 0 2 0h1v-2.874a1.12 1.12 0 0 0-1.075-1.116zM4.75 12a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5zm.339-3l.328-2l4.158-.007L9.91 9H5.089zm5.161 3a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarRepair(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 5a2 2 0 0 0 1.732-1H12a1 1 0 1 0 0-2H4.732a2 2 0 0 0-3.464 0H3v2H1.268A2 2 0 0 0 3 5m-.854 4.354A.5.5 0 0 0 2 9.707V13.5a.5.5 0 0 0 .5.5H4a.5.5 0 0 0 .5-.5V13h6v.5a.5.5 0 0 0 .5.5h1.5a.5.5 0 0 0 .5-.5V9.707a.5.5 0 0 0-.146-.353L12 8.5l-1.354-2.257a.5.5 0 0 0-.43-.243H4.784a.5.5 0 0 0-.429.243L3 8.5zM11.134 9H3.866l1.2-2h4.868zM5.5 10.828v.372a.3.3 0 0 1-.3.3H3.3a.3.3 0 0 1-.3-.3v-.834a.3.3 0 0 1 .359-.294l1.82.364a.4.4 0 0 1 .321.392m6.5-.34v.712a.3.3 0 0 1-.3.3H9.8a.3.3 0 0 1-.3-.3v-.454a.3.3 0 0 1 .241-.294l1.78-.356a.4.4 0 0 1 .479.392"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarRepairEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.5 1.25H4.352A1.991 1.991 0 0 0 .778 1H2.5L3 2l-.5 1H.777a1.991 1.991 0 0 0 3.575-.25H9.5a.75.75 0 0 0 0-1.5z" fill="currentColor"/><path d="M8.31 7.034L8.061 5.83A.988.988 0 0 0 7.083 5H3.917c-.488 0-.9.35-.975.81L2.69 7.036a.853.853 0 0 0-.689.83V10h.5a1 1 0 0 0 2 0h2a1 1 0 0 0 2 0H9V7.86c0-.415-.3-.744-.69-.826zM3.75 9a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5zm-.037-2l.204-1l3.162.012l.204.988h-3.57zM7.25 9a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarRepairFifteen(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M10.925 9.01l-.363-2.18A.988.988 0 0 0 9.583 6H5.417a.989.989 0 0 0-.978.83l-.364 2.182A1.126 1.126 0 0 0 3 10.132V13h1a1 1 0 0 0 2 0h3a1 1 0 0 0 2 0h1v-2.874a1.12 1.12 0 0 0-1.075-1.116zM4.75 12a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5zm.339-3l.328-2l4.158-.007L9.91 9H5.089zm5.161 3a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5z" fill="currentColor"/><path d="M12 2H5.789A2.498 2.498 0 0 0 1.21 2H3.5L4 3l-.5 1H1.211A2.498 2.498 0 0 0 5.79 4H12a1 1 0 1 0 0-2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Casino(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 10A2.5 2.5 0 0 0 13 8.5c0-.564-.194-1.079-.509-1.497L12.5 7l-5-6l-5 6l.009.003A2.478 2.478 0 0 0 2 8.5A2.5 2.5 0 0 0 6.5 10l.5-.666V11.5C7 13 4.5 13 4.5 13a.5.5 0 1 0 0 1h6a.5.5 0 0 0 0-1S8 13 8 11.5V9.334z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CasinoEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M6.136 6.85c.29.395.753.65 1.273.65C8.288 7.5 9 6.772 9 5.875c0-.367-.123-.701-.324-.973l.006-.002L5.5 1L2.318 4.9l.006.002c-.2.272-.324.606-.324.973C2 6.772 2.712 7.5 3.59 7.5c.521 0 .983-.255 1.274-.65l.261-.356V8.5c0 .5-1.75.75-1.75.75a.375.375 0 0 0 0 .75h4.25a.375.375 0 0 0 0-.75S5.875 9 5.875 8.5V6.494l.261.356z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Castle(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 4H4a1 1 0 0 1-1-1V.5a.5.5 0 0 1 1 0V2h1V1a1 1 0 1 1 2 0v1h1V1a1 1 0 1 1 2 0v1h1V.5a.5.5 0 0 1 1 0V3a1 1 0 0 1-1 1m3 10.5a.5.5 0 0 1-.5.5h-12a.5.5 0 0 1 0-1H2a1 1 0 0 0 1-1s1-6 1-7a1 1 0 0 1 1-1h5a1 1 0 0 1 1 1c0 1 1 7 1 7a1 1 0 0 0 1 1h.5a.5.5 0 0 1 .5.49zm-5-4a1.5 1.5 0 0 0-3 0V14h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CastleEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M8.67.81v1.48A.69.69 0 0 1 8 3H3.09a.69.69 0 0 1-.7-.7V.81a.345.345 0 0 1 .69 0v.69H4V1a.5.5 0 0 1 1 0v.5h1V1a.5.5 0 0 1 1 0v.5h1V.81a.34.34 0 1 1 .67 0zm1.39 8.82a.35.35 0 0 1-.35.35H1.35a.35.35 0 0 1 0-.7h.35a.68.68 0 0 0 .7-.7s.7-3.2.7-3.89A.68.68 0 0 1 3.79 4h3.48a.68.68 0 0 1 .7.7c0 .7.7 3.89.7 3.89a.68.68 0 0 0 .7.7h.34a.35.35 0 0 1 .35.35v-.01zM6.5 7.5a1 1 0 1 0-2 0v2h2v-2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CastleJp(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 12.75a1 1 0 0 1-1-1v-4.5H10a1 1 0 0 1-1-1v-2H6v2a1 1 0 0 1-1 1H3.5v4.5a1 1 0 0 1-2 0v-5.5a1 1 0 0 1 1-1H4v-2a1 1 0 0 1 1-1h5a1 1 0 0 1 1 1v2h1.5a1 1 0 0 1 1 1v5.5a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CastleJpEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M1.75 9.503a.75.75 0 0 1-.75-.75V4.748a.75.75 0 0 1 .75-.75l1.251-.001v-1.75a.75.75 0 0 1 .75-.75h3.5a.75.75 0 0 1 .75.75v1.75l1.25.001a.75.75 0 0 1 .75.75v4a.75.75 0 0 1-1.5 0v-3.25l-1.25-.001a.75.75 0 0 1-.75-.75v-1.75h-2v1.75a.75.75 0 0 1-.75.75L2.5 5.498v3.255a.75.75 0 0 1-.75.75z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Caution(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.093 11.892L6.84 1.391a.752.752 0 0 1 1.32 0l5.747 10.501a.75.75 0 0 1-.66 1.11H1.753a.75.75 0 0 1-.66-1.11M8.3 8l.403-2.418A.5.5 0 0 0 8.21 5H6.79a.5.5 0 0 0-.493.582L6.7 8zm.3 1.9a1.1 1.1 0 1 0-2.2 0a1.1 1.1 0 0 0 2.2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cemetery(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.46 12h-.68L12 3.55a.52.52 0 0 0-.56-.55h-1.18c0-.92-1.23-2-2.75-2S4.77 2.08 4.77 3H3.54a.52.52 0 0 0-.54.55L4.2 12h-.65a.53.53 0 0 0-.55.5V14h9v-1.51a.52.52 0 0 0-.54-.49M4.5 5h6v1h-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CemeteryEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M8.65 8H8l1-5.61A.36.36 0 0 0 8.63 2H7.16c0-.65-.7-1-1.67-1s-1.83.35-1.83 1H2.35a.36.36 0 0 0-.35.39L3 8h-.65a.35.35 0 0 0-.35.36V10h7V8.36A.35.35 0 0 0 8.66 8h-.01zM7 5H4V4h3v1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CemeteryJp(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.5 11h-3V2.5a.945.945 0 0 0-1-1a.945.945 0 0 0-1 1V11h-3a1 1 0 0 0 0 2h8a.945.945 0 0 0 1-1a1.002 1.002 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CemeteryJpEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M8.205 7.905h-2v-5.6a.736.736 0 0 0-.8-.8a.713.713 0 0 0-.794.62a.734.734 0 0 0-.006.08v5.7h-1.8a.736.736 0 0 0-.8.8a.714.714 0 0 0 .623.795a.701.701 0 0 0 .077.005h5.5a.736.736 0 0 0 .8-.666a.75.75 0 0 0 0-.134a.768.768 0 0 0-.735-.8h-.065z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChargingStation(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.646 7.801L7.752.838a.826.826 0 0 1 1.463.705L8.086 5.684A.25.25 0 0 0 8.327 6h3.42a.753.753 0 0 1 .607 1.199l-5.106 6.964a.826.826 0 0 1-1.463-.706l1.129-4.141A.25.25 0 0 0 6.673 9h-3.42a.753.753 0 0 1-.607-1.199"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChargingStationEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.5 3H9V1.5a.5.5 0 0 0-1 0V3a1 1 0 0 0 1 1v4.25a.25.25 0 1 1-.5 0V6.5A1.5 1.5 0 0 0 7 5V2a1 1 0 0 0-1-1H2a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V6a.5.5 0 0 1 .5.5v1.75a1.25 1.25 0 1 0 2.5 0V3.5a.5.5 0 0 0-.5-.5zm-6 5.75H3L4 6H1.75L4.5 2.25H5L4 5h2.25L3.5 8.75z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cinema(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 7.5v2a.5.5 0 0 1-1 0s.06-.5-1-.5h-1v2.5a.5.5 0 0 1-.5.5h-8a.5.5 0 0 1-.5-.5v-4a.5.5 0 0 1 .5-.5h8a.5.5 0 0 1 .5.5V8h1c1.06 0 1-.5 1-.5a.5.5 0 0 1 1 0M4 3a2 2 0 1 0 0 4a2 2 0 0 0 0-4m0 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2m4.5-4a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5m0 4a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CinemaEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M10 5.5v2a.5.5 0 0 1-1 0a.66.66 0 0 0-.51-.5H8v1.63a.37.37 0 0 1-.37.37H1.37A.37.37 0 0 1 1 8.63V5.37A.37.37 0 0 1 1.37 5h6.26a.37.37 0 0 1 .37.37V6h.49A.66.66 0 0 0 9 5.5a.5.5 0 0 1 1 0zM2.5 2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3zm0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zM6 1a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm0 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circle(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 7.5a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M10 5.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleStroked(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 12.3a4.8 4.8 0 1 1 0-9.6a4.8 4.8 0 0 1 0 9.6m0 1.7a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleStrokedEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.5 0a5.5 5.5 0 1 1 0 11a5.5 5.5 0 0 1 0-11zm0 1.222a4.278 4.278 0 1 0 0 8.556a4.278 4.278 0 0 0 0-8.556z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func City(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.637 4h-1.639V2.36a.36.36 0 0 0-.36-.36h-.279a.36.36 0 0 0-.36.36V4H8.36a.36.36 0 0 0-.36.36v5.638H4.361c-.2 0-.363.163-.363.363v2.275c0 .2.163.362.363.362h8.275c.2 0 .36-.162.36-.361V4.36a.36.36 0 0 0-.36-.36m-6.638 7.998h-1v-1h1zm2 0h-1v-1h1zm2 0h-1v-1h1zm0-2h-1v-1h1zm0-2h-1v-1h1zm0-2h-1v-1h1zm2 6h-1v-1h1zm0-2h-1v-1h1zm0-2h-1v-1h1zm0-2h-1v-1h1zm-5-3.637A.36.36 0 0 0 6.638 2H4.36a.36.36 0 0 0-.36.36V4H2.36a.36.36 0 0 0-.36.36v8.287c0 .194.157.35.35.35H3V9h3.999zM4 7.998H3v-1h1zm0-2H3v-1h1zm2 2H5v-1h1zm0-2H5v-1h1zm0-2H5V3h1v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CityEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.751 3h-.75V1.25a.25.25 0 0 0-.25-.25h-.5a.25.25 0 0 0-.25.25V3h-.75a.25.25 0 0 0-.25.25v3.751h-1.75a.251.251 0 0 0-.25.251v2.5a.251.251 0 0 0 .25.249h4.501a.25.25 0 0 0 .25-.25V3.25a.25.25 0 0 0-.25-.25zm-3.75 5.001h1v1H6v-1zm3 1H8v-1h1v1zm0-2H8v-1h1v1zm0-2H8v-1h1v1zm-3-3.75A.25.25 0 0 0 5.75 1H3.25a.25.25 0 0 0-.25.25V2h-.75a.25.25 0 0 0-.25.25V3h-.75a.25.25 0 0 0-.25.25v6.501c0 .138.112.25.25.25H4v-4h2V1.25zM3 9H2v-1h1v1zm0-2H2v-1h1v1zm0-2H2v-1h1v1zm2 0H4v-1h1v1zm0-2H4V2h1v1.001z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClothingStore(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2L1 4.5V7h3v6h7V7h3V4.5L11 2H9.5l-2 4l-2-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClothingStoreEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M2.5 1l-2 2v2h2v5h6V5h2V3l-2-2H7L5.5 4L4 1H2.5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func College(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 1L0 4.5l2 .9v1.7c-.6.2-1 .8-1 1.4s.4 1.2 1 1.4v.1l-.9 2.1C.8 13 1 14 2.5 14s1.7-1 1.4-1.9L3 10c.6-.3 1-.8 1-1.5s-.4-1.2-1-1.4V5.9L7.5 8L15 4.5zm4.4 6.5l-4.5 2L5 8.4v.1c0 .7-.3 1.3-.8 1.8l.6 1.4v.1c.1.4.2.8.1 1.2c.7.3 1.5.5 2.5.5c3.3 0 4.5-2 4.5-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollegeEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M2 7.3c.3-.2.5-.5.5-.8c0-.4-.2-.7-.6-.9V4.4L5.5 6L11 3.5L5.5 1L0 3.5l1.2.5v1.6c-.4.2-.6.5-.6.9c0 .3.2.6.5.8L.6 9c-.3 1 .5 1 .5 1h1s.8 0 .5-1L2 7.3z" fill="currentColor"/><path d="M3.5 6.2v.3c0 .4-.2.8-.4 1.1c.2.4.4.8.4 1.4v.6c.5.2 1.2.4 2 .4C8 10 9 8.5 9 8.5v-3L5.5 7.1l-2-.9z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollegeJp(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 3h-4V2a1 1 0 0 0-2 0v1h-4a1 1 0 0 0 0 2h.67a10.817 10.817 0 0 0 1.98 3.81c.333.432.66.825.979 1.194a14.383 14.383 0 0 1-2.608 2.117a1 1 0 1 0 .955 1.758A15.65 15.65 0 0 0 7.5 11.461a15.65 15.65 0 0 0 3.024 2.418a1 1 0 1 0 .955-1.758a14.379 14.379 0 0 1-2.607-2.117c.319-.369.645-.762.978-1.193a10.818 10.818 0 0 0 1.98-3.812h.67a1 1 0 0 0 0-2M8.267 7.587c-.261.338-.515.641-.767.937a25.017 25.017 0 0 1-.767-.937a9.589 9.589 0 0 1-1.465-2.589h4.464a9.579 9.579 0 0 1-1.465 2.589"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollegeJpEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.25 2.25h-3V1.5a.75.75 0 0 0-1.5 0v.75h-3a.75.75 0 0 0 0 1.5h.403a9.139 9.139 0 0 0 1.66 2.966c.218.274.43.52.636.754a6.088 6.088 0 0 1-1.9 1.307a.75.75 0 0 0 .391 1.448a6.728 6.728 0 0 0 2.558-1.672a6.722 6.722 0 0 0 2.563 1.672a.736.736 0 0 0 .188.024a.746.746 0 0 0 .717-.55a.756.756 0 0 0-.515-.923A6.024 6.024 0 0 1 6.546 7.47c.207-.234.417-.48.636-.754a9.193 9.193 0 0 0 1.663-2.967h.405a.75.75 0 0 0 0-1.5zM6.008 5.781c-.177.223-.343.41-.51.601c-.168-.191-.334-.378-.511-.6A8.51 8.51 0 0 1 3.75 3.748h3.497a8.574 8.574 0 0 1-1.24 2.033z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Commercial(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 6H1a11.431 11.431 0 0 1 1-4h11a11.429 11.429 0 0 1 1 4M3 7h9v6h-1V8H8v5H3zm1 3h3V8H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommercialEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M0 5a4.71 4.71 0 0 1 1-3h9a4.71 4.71 0 0 1 1 3H0zm2 1v5h4V7h2v4h1V6H2zm3 3H3V7h2v2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommunicationsTower(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.855 6.434l-.414-.282a4.762 4.762 0 0 0 .282-4.878l-.084-.153l.436-.246l.09.167a5.262 5.262 0 0 1-.31 5.392m1.152 7.131V14h-11v-.435h2.488L6.044 6.4a.5.5 0 0 1 .489-.394h.725V4.6a1.14 1.14 0 0 1-.882-1.1a1.157 1.157 0 1 1 2.313 0a1.14 1.14 0 0 1-.93 1.111v1.395h.721a.5.5 0 0 1 .49.394l1.547 7.165ZM8.454 8.751H6.56l-.323 1.493h2.541Zm-2.326 1.993l-.323 1.495h3.403l-.323-1.495Zm.808-3.738l-.27 1.245h1.68l-.269-1.245Zm-1.418 6.56h3.977l-.179-.827h-3.62ZM5.21 5.013a2.752 2.752 0 0 1 .016-3.052l-.414-.28a3.25 3.25 0 0 0-.019 3.607ZM10.757 3.5a3.243 3.243 0 0 0-.534-1.786l-.418.275a2.752 2.752 0 0 1-.018 3.05l.414.278a3.234 3.234 0 0 0 .556-1.817M3.534 6.118a4.764 4.764 0 0 1-.153-4.988L2.948.88a5.264 5.264 0 0 0 .17 5.514Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommunicationsTowerEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M7.655 3.564h-.4V1.448h.4zM3.702 1.448h-.4v2.116h.4zm6-.285h-.4v3.278h.4zm-8.023 0h-.4v3.278h.4zm4.078 3.005h.437a.354.354 0 0 1 .346.282l1.04 4.987H9V10H2v-.563h1.434l1.04-4.987a.355.355 0 0 1 .344-.281h.44V2.912a.915.915 0 0 1-.668-.875a.927.927 0 1 1 1.167.888zm-.96 2.188h1.419l-.309-1.481h-.801zm-.414 1.989H6.63l-.351-1.688H4.735zm2.475 1.092l-.165-.791H4.32l-.165.79z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Confectionery(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 4a1 1 0 0 0-1-1a1 1 0 0 0-2 0v2.067A3.48 3.48 0 0 0 4.358 9H2a1 1 0 0 0 0 2a1 1 0 0 0 1 1a1 1 0 0 0 2 0V9.933A3.48 3.48 0 0 0 10.642 6H13a1 1 0 0 0 0-2M7.5 9.993a2.484 2.484 0 0 1-1.182-4.674a1.08 1.08 0 0 1 .546.307a2.124 2.124 0 0 1 .386 1.386l.001.956a2.582 2.582 0 0 0 .509 1.735a1.546 1.546 0 0 0 .26.233a2.486 2.486 0 0 1-.52.056m1.181-.312a1.081 1.081 0 0 1-.545-.307a2.124 2.124 0 0 1-.386-1.386l-.001-.956a2.582 2.582 0 0 0-.509-1.735a1.556 1.556 0 0 0-.26-.233A2.484 2.484 0 0 1 8.682 9.68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ConfectioneryEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.24 2.513a.746.746 0 0 0-.745-.746a.746.746 0 1 0-1.493 0v1.746A2.489 2.489 0 0 0 3.512 7H1.747a.746.746 0 0 0 0 1.492a.746.746 0 0 0 .747.746a.746.746 0 0 0 1.492 0v-1.76a2.489 2.489 0 0 0 3.506-3.473h1.748a.746.746 0 0 0 0-1.492zM5.5 7.284a1.782 1.782 0 0 1-.72-3.413a.816.816 0 0 1 .244.177a1.648 1.648 0 0 1 .301 1.076v.737a1.972 1.972 0 0 0 .388 1.322a1.155 1.155 0 0 0 .096.07a1.776 1.776 0 0 1-.309.031zm.72-.155a.816.816 0 0 1-.244-.177a1.648 1.648 0 0 1-.301-1.076v-.737a1.972 1.972 0 0 0-.388-1.322a1.155 1.155 0 0 0-.096-.07A1.78 1.78 0 0 1 6.22 7.13z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Construction(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 11h-1.8L8.2.5C8-.2 7-.2 6.8.5L3.3 11H1.5c-.3 0-.5.2-.5.5v1c0 .3.2.5.5.5h12c.3 0 .5-.2.5-.5v-1c0-.3-.2-.5-.5-.5M7 3h1l.7 2H6.4zM5.7 7h3.6l.7 2H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Convenience(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.998 3.069s.188-1.062 1-1.062h1.006v-.28a.5.5 0 0 1 .5-.5h3.984a.5.5 0 0 1 .5.5v.28h1.01a.832.832 0 0 1 .638.293a1.759 1.759 0 0 1 .362.769s.311 1.572.592 2.918h-.938L11 3l-7.018-.004l-.648 2.99h-.952Zm11 4.914a1 1 0 0 1-.999 1l-1 3.973s-.188 1.042-1 1.042h-7c-.813 0-1-1.042-1-1.042L1.981 8.98a.999.999 0 0 1 .016-1.997h11a1 1 0 0 1 1 1M5.033 12.13l-.003-1.125l-.39-1.536a.537.537 0 1 0-1.066.144l.206 1.395l.187 1.267a.537.537 0 1 0 1.066-.145m2-2.597A.533.533 0 0 0 6.499 9h-.01a.5.5 0 0 0-.358.15a.529.529 0 0 0-.165.384v2.68a.533.533 0 1 0 1.067 0Zm1.998 0a.531.531 0 0 0-.337-.493a.5.5 0 0 0-.393 0a.531.531 0 0 0-.337.493v2.68a.533.533 0 1 0 1.067 0Zm2.385.08a.537.537 0 1 0-1.065-.144l-.391 1.525l-.002 1.136a.537.537 0 1 0 1.065.145Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ConvenienceEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M8.821 4.269h-.62c-.155-.6-.28-1.066-.283-1.074l-.01-.05c0-.001-.12-.548-.41-.548H3.518c-.289 0-.41.547-.41.553c-.008.033-.136.511-.293 1.118h-.62c.176-.68.323-1.23.323-1.23s.187-1.04 1-1.04h.496a.493.493 0 0 1 .483-.416h2a.493.493 0 0 1 .483.415h.517c.812 0 1 1.042 1 1.042s.147.549.323 1.23zM2.92 9.995a.494.494 0 0 1-.483-.368c-.23-.873-.8-3.053-.938-3.63a.5.5 0 1 1 0-1h8a.5.5 0 0 1 0 1c-.137.577-.715 2.76-.948 3.632a.495.495 0 0 1-.483.366zM4.013 6.7a.513.513 0 0 0-1.024 0V8.51a.513.513 0 0 0 1.024 0zm1.997 0a.513.513 0 0 0-1.024 0V8.51a.513.513 0 0 0 1.024 0zm2.006 0a.513.513 0 0 0-1.024 0V8.51a.513.513 0 0 0 1.024 0z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cricket(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0m2.88 5.68l-2-2a.5.5 0 0 0-.4-.18H3.75a.74.74 0 0 0-.75.64l-1 7.7a.483.483 0 0 0 0 .14a.5.5 0 0 0 .5.5a.49.49 0 0 0 .5-.34l1.2-3.89l.26-.83l.4.44L6 10.6v2.9a.5.5 0 0 0 1 0v-3a.48.48 0 0 0-.08-.22L5.48 8.5l1-2.5h1.71l2 1.84a.49.49 0 0 0 .37.17a.5.5 0 0 0 .44-.51a.49.49 0 0 0-.12-.32M14 11.27a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5m-3.24-1.53V9a.25.25 0 1 0-.5 0v.74a.49.49 0 0 0-.25.42v3.34a.5.5 0 0 0 1 0v-3.34a.49.49 0 0 0-.25-.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CricketEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M6 1a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm2.88 4.18l-1-2A.5.5 0 0 0 7.5 3H2.77a.74.74 0 0 0-.77.65l-1 6.71a.483.483 0 0 0 0 .14a.5.5 0 0 0 .5.5a.49.49 0 0 0 .5-.36l1.22-3.89l.21-.83l.4.44L5 7.6v2.9a.5.5 0 0 0 1 0v-3a.48.48 0 0 0-.1-.28L4.45 5.5L5.5 4h1.71l.92 1.84A.49.49 0 0 0 8.5 6a.5.5 0 0 0 .5-.49a.88.88 0 0 0-.12-.33zM10.5 8a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1zM8.69 8v-.75a.25.25 0 1 0-.5 0V8a.49.49 0 0 0-.25.42v2.08a.5.5 0 0 0 1 0V8.41A.49.49 0 0 0 8.69 8z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cross(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.64 2.27L7.5 6.13l3.84-3.84A.92.92 0 0 1 12 2a1 1 0 0 1 1 1a.9.9 0 0 1-.27.66L8.84 7.5l3.89 3.89A.9.9 0 0 1 13 12a1 1 0 0 1-1 1a.92.92 0 0 1-.69-.27L7.5 8.87l-3.85 3.85A.92.92 0 0 1 3 13a1 1 0 0 1-1-1a.9.9 0 0 1 .27-.66L6.16 7.5L2.27 3.61A.9.9 0 0 1 2 3a1 1 0 0 1 1-1c.24.003.47.1.64.27"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CrossEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M2.2 1.19l3.3 3.3L8.8 1.2a.67.67 0 0 1 .5-.2a.75.75 0 0 1 .7.7a.66.66 0 0 1-.2.48L6.49 5.5L9.8 8.82c.13.126.202.3.2.48a.75.75 0 0 1-.7.7a.67.67 0 0 1-.5-.2L5.5 6.51L2.21 9.8a.67.67 0 0 1-.5.2a.75.75 0 0 1-.71-.71a.66.66 0 0 1 .2-.48L4.51 5.5L1.19 2.18A.66.66 0 0 1 1 1.7a.75.75 0 0 1 .7-.7a.67.67 0 0 1 .5.19z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dam(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.94 9.5a.5.5 0 0 1-.5.5a1 1 0 0 0-.67.34a2.14 2.14 0 0 1-1 .6a1.89 1.89 0 0 1-1.76-.37l-.39-.35a.77.77 0 0 0-1.089-.011l-.011.011c-.14.12-.27.25-.42.37a1.91 1.91 0 0 1-2.46-.07l-.34-.33l-.1-.06l.3 1.22l.49 2a.5.5 0 0 1-.49.65h-4a.5.5 0 0 1-.5-.5v-12a.5.5 0 0 1 .5-.5h1.1a.5.5 0 0 1 .49.39L4 5.06v.14a1.88 1.88 0 0 1 2 .24c.16.13.31.28.47.41a.75.75 0 0 0 1 0c.16-.13.31-.28.47-.41a1.9 1.9 0 0 1 2.45 0c.15.13.29.27.44.39a.75.75 0 0 0 1 0l.47-.41A1.78 1.78 0 0 1 13.43 5a.5.5 0 0 1 0 1a1 1 0 0 0-.67.34a2.14 2.14 0 0 1-1 .6A1.89 1.89 0 0 1 10 6.57l-.33-.34a.77.77 0 0 0-1.089-.011l-.011.011c-.14.12-.27.25-.42.37a1.91 1.91 0 0 1-2.46-.07l-.39-.34a.75.75 0 0 0-1 0l-.06.05L4.93 9A1.9 1.9 0 0 1 6 9.43c.16.13.31.28.47.41a.75.75 0 0 0 1 0c.16-.13.31-.28.47-.41a1.9 1.9 0 0 1 2.45 0c.15.13.29.27.44.39a.75.75 0 0 0 1 0l.47-.41A1.78 1.78 0 0 1 13.43 9a.5.5 0 0 1 .51.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DamEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M10 7.51a.5.5 0 0 1-.5.5a1 1 0 0 0-.67.34a2.14 2.14 0 0 1-1 .6a1.89 1.89 0 0 1-1.76-.37l-.34-.35a.77.77 0 0 0-1.08 0L5 9.38a.5.5 0 0 1-.5.62h-3a.5.5 0 0 1-.5-.5v-8a.5.5 0 0 1 .49-.5h.7a.5.5 0 0 1 .49.39l.73 2.51l.1-.06c.16-.12.31-.27.49-.4a1.9 1.9 0 0 1 2.45 0c.15.13.29.27.44.39a.75.75 0 0 0 1 0l.41-.4A1.78 1.78 0 0 1 9.49 3a.5.5 0 0 1 0 1a1 1 0 0 0-.67.34a2.14 2.14 0 0 1-1 .6a1.89 1.89 0 0 1-1.76-.37l-.33-.34a.77.77 0 0 0-1.089-.011l-.011.011c-.14.12-.27.25-.42.37a2 2 0 0 1-.52.29l.67 2.32a1.88 1.88 0 0 1 2.06.24c.15.13.29.27.44.39a.75.75 0 0 0 1 0l.44-.42A1.78 1.78 0 0 1 9.49 7a.5.5 0 0 1 .51.51z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Danger(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.94 14.68a.5.5 0 0 1-.47.32c-.06.01-.12.01-.18 0L7.5 12.56L1.7 15a.5.5 0 0 1-.648-.284l-.002-.006a.5.5 0 0 1 .29-.71l4.85-2l-4.85-2a.5.5 0 1 1 .36-.93l5.8 2.41l5.8-2.41a.5.5 0 1 1 .36.93L8.8 12l4.85 2a.5.5 0 0 1 .29.68M12 4.23v.45a1 1 0 0 1-.2.59A17.718 17.718 0 0 1 10 7v1.16a.5.5 0 0 1-.32.47l-2.16.87h-.07l-2.17-.87A.5.5 0 0 1 5 8.16V7a17.563 17.563 0 0 1-1.82-1.73A1 1 0 0 1 3 4.68v-.45A4.6 4.6 0 0 1 7.11 0h.75A4.59 4.59 0 0 1 12 4.23M6 4a1 1 0 1 0-2 0a1 1 0 0 0 2 0m1 3a.5.5 0 0 0-1 0v.5a.5.5 0 0 0 1 0zm2 0a.5.5 0 0 0-1 0v.5a.5.5 0 0 0 1 0zm2-3a1 1 0 1 0-2 0a1 1 0 0 0 2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DangerEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M6.62 8.5l3.11 1.55l-.45.89L5.5 9.06l-3.78 1.89l-.45-.89L4.38 8.5l-3.1-1.55l.45-.89L5.5 7.94l3.78-1.89l.44.9l-3.1 1.55zM8.5 3.21v.29l-1 1v1l-2 1l-2-1v-1l-1-1V3a2.9 2.9 0 0 1 3-3a3.09 3.09 0 0 1 3 3.21zm-3.79-.5a.79.79 0 1 0-.79.79a.79.79 0 0 0 .79-.79zM5 4.5h-.5v1H5v-1zm1.5 0H6v1h.5v-1zm1.36-1.79a.79.79 0 1 0-.79.79a.79.79 0 0 0 .79-.79z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Defibrillator(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.55 7.338C-.837 2.742 5.18-.322 7.502 4.274c2.321-4.597 8.339-1.532 5.952 3.064c-.087.167-.203.346-.311.521h-1.808L9.52 5.138a.625.625 0 0 0-1.1.114l-1.648 4.12l-1.33-1.33A.625.625 0 0 0 5 7.86H1.86c-.109-.175-.224-.354-.311-.52M11 9.11a.626.626 0 0 1-.52-.278L9.139 6.82L7.58 10.717a.625.625 0 0 1-1.022.21L4.741 9.109H2.736a42.67 42.67 0 0 0 4.46 4.674a.464.464 0 0 0 .622 0a43.255 43.255 0 0 0 4.45-4.674z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DefibrillatorEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M1.42 4.795C-.213 1.53 3.933-.65 5.512 2.615c1.58-3.266 5.725-1.086 4.092 2.18c-.023.038-.048.074-.071.111H8.387l-.828-1.654a.65.65 0 0 0-1.118 0L5 6.134l-.441-.882A.624.624 0 0 0 4 4.906H1.491l-.071-.11zM8 6.156a.624.624 0 0 1-.559-.345L7 4.929L5.559 7.81a.624.624 0 0 1-1.118 0l-.828-1.654H2.301a24.227 24.227 0 0 0 2.897 3.45l.015.015a.44.44 0 0 0 .622-.016a24.229 24.229 0 0 0 2.89-3.449H8z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dentist(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.36 14c-1 0-.56-2.67-.86-5c-.1-.76-1-1.49-1.12-2.06C2 5 1.39 1.44 3.66 1S6 3 7.54 3S9.11.64 11.39 1s1.59 3.9 1.29 5.9c-.1.45-1.1 1.48-1.18 2.06c-.33 2.4.32 5-.8 5c-.93 0-1.32-2.72-2-4.5C8.43 8.63 8.06 8 7.54 8C6 8 5.75 14 4.36 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DentistEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.48 6A2 2 0 0 0 4 7c-.46 1.21-.14 3-.82 3S2.7 8.49 2.5 7a16.259 16.259 0 0 0-.76-1.89C1.53 3.7 1 1.28 2.67 1S4.35 2.52 5.5 2.52S6.67.72 8.33 1s1.14 2.7.93 4.11A16.259 16.259 0 0 0 8.5 7c-.2 1.49 0 3-.68 3S7.46 8.21 7 7a2 2 0 0 0-1.48-1h-.04z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Diamond(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.037 1.201L2.125 6.7a.482.482 0 0 0-.124.3c-.006.11.028.21.096.3l4.91 6.46c.055.074.13.135.215.177a.637.637 0 0 0 .556 0a.595.595 0 0 0 .216-.177l4.91-6.46a.453.453 0 0 0 .095-.3a.482.482 0 0 0-.124-.3L7.963 1.201a.604.604 0 0 0-.208-.148a.64.64 0 0 0-.718.148"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Doctor(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 7A2.5 2.5 0 0 1 3 4.5v-2a.5.5 0 0 1 .5-.5H4a.5.5 0 0 0 0-1h-.5A1.5 1.5 0 0 0 2 2.5v2a3.49 3.49 0 0 0 1.51 2.87A4.41 4.41 0 0 1 5 10.5a3.5 3.5 0 1 0 7 0v-.57a2 2 0 1 0-1 0v.57a2.5 2.5 0 0 1-5 0a4.41 4.41 0 0 1 1.5-3.13A3.49 3.49 0 0 0 9 4.5v-2A1.5 1.5 0 0 0 7.5 1H7a.5.5 0 0 0 0 1h.5a.5.5 0 0 1 .5.5v2A2.5 2.5 0 0 1 5.5 7m6 2a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoctorEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.5 5.87A1.48 1.48 0 1 0 7.65 7.3v.42a1.855 1.855 0 1 1-3.71 0A3.27 3.27 0 0 1 5.05 5.4H5a2.58 2.58 0 0 0 1.17-2.12V1.79A1.11 1.11 0 0 0 5.06.68H4.5a.37.37 0 0 0 0 .74h.55a.37.37 0 0 1 .37.37v1.49a1.85 1.85 0 0 1-1.84 1.85v.27v-.27a1.85 1.85 0 0 1-1.86-1.84v-1.5a.37.37 0 0 1 .37-.37h.52a.37.37 0 0 0 0-.74h-.52A1.11 1.11 0 0 0 1 1.79v1.49A2.58 2.58 0 0 0 2.1 5.4a3.27 3.27 0 0 1 1.1 2.32a2.595 2.595 0 1 0 5.19 0V7.3A1.48 1.48 0 0 0 9.5 5.87zM8 6.61a.74.74 0 1 1 .74-.74a.74.74 0 0 1-.74.74z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DogPark(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.3 1.22c-.386 0-.7-.02-.8.78l-.4 2.521L11.5 6.5h2c1.4 0 1.5-.959 1.5-.959l-1.9-2.219c-.7-.7-1.4-.822-2.1-.822V2s.067-.78-.7-.78m-5.55.28s-.95.02-1.55.22c-.6.2-1.2.901-1.2 1.901v3.9C2 9.221 1.3 9.5 1 9.5c0 0-1 .021-1 1.021v2.2s0 .8.8.8H2v-.699c0-.4-.2-.6-.5-.7v-1.1c1 0 1.1-.201 1.5-.301l.55 2.197c.1.3.2.5.5.6h1L6 13.5v-.8c0-.677-1-.7-1-.7V9.5h3.5l.7 2.621c.4 1.4 1.3 1.379 1.3 1.379H12v-.8c0-.713-1-.7-1-.7l.1-4.078L8 5.5H3.5v-2c0-.4.276-.495.5-.5c.494-.012.75 0 .75 0a.75.75 0 0 0 0-1.5m7 2.5a.25.25 0 1 1 0 .5a.25.25 0 0 1 0-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DogParkEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M7.5 1s-.346.037-.5.5l-.5 2l2 1.5H10c1 0 1-1 1-1L9.5 2.5C9 2 8.5 2 8 2v-.5S8 1 7.5 1zm-5 1s-.353-.007-.723.178S1 2.833 1 3.5v2C1 6.5 1 7 .5 7c0 0-.5 0-.5.5v2s0 .5.5.5s.5-.5.5-.5V8c.354 0 .69-.137 1-.299V9.5s0 .5.5.5s.5-.5.5-.5V7h3l.664 1.992C7 10 7.5 10 7.5 10H8s.5 0 .5-.5S8 9 8 9V6.5c0-.89-.632-1.255-1-1.5L5.498 4H2v-.5c0-.333.092-.362.223-.428C2.353 3.007 2.5 3 2.5 3c.676.01.676-1.01 0-1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DrinkingWater(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 1a2 2 0 0 0-2 2v3.5a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 1 .5-.5H14V1Zm1 14H4a.5.5 0 0 1-.48-.38L2 8.62a.5.5 0 0 1 .365-.606A.558.558 0 0 1 2.5 8h6a.5.5 0 0 1 .514.485A.47.47 0 0 1 9 8.62l-1.5 6A.5.5 0 0 1 7 15m-3.35-4h3.71l.5-2H3.14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DrinkingWaterEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5 11H3a.51.51 0 0 1-.5-.4L1 5.6a.5.5 0 0 1 .5-.6h5a.5.5 0 0 1 .5.6l-1.49 5A.51.51 0 0 1 5 11zM2.76 8h2.46l.67-2H2.11z" fill="currentColor"/><path d="M4.5 0A1.5 1.5 0 0 0 3 1.51v2a.5.5 0 0 0 .5.5h1A.5.5 0 0 0 5 3.5v-1a.5.5 0 0 1 .5-.5H10V0H4.5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DrinkingWaterFifteen(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M6 1a2 2 0 0 0-2 2v3.5a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 1 .5-.5H14V1H6z" fill="currentColor"/><path d="M7 15H4a.5.5 0 0 1-.48-.38L2 8.62A.5.5 0 0 1 2.5 8h6a.5.5 0 0 1 .5.62l-1.5 6A.5.5 0 0 1 7 15zm-3.35-4h3.71l.5-2H3.14z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Elevator(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 1H4a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h7a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1M7.5 12.5l-2-4h4Zm-2-6l2-4l2 4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Embassy(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.65 2C5.43 2 4.48 3.38 4.11 3.82a.49.49 0 0 0-.11.32v4.4a.44.44 0 0 0 .72.36a3 3 0 0 1 1.93-1.17C8.06 7.73 8.6 9 10.07 9a5.28 5.28 0 0 0 2.73-1.09a.49.49 0 0 0 .2-.4V2.45a.44.44 0 0 0-.62-.45a5.75 5.75 0 0 1-2.31 1.06C8.6 3.08 8.12 2 6.65 2M2.5 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2M3 4v9.48a.5.5 0 0 1-1 0V4a.5.5 0 0 1 1 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmbassyEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.5 2a4.07 4.07 0 0 0-2.41 1.06a.36.36 0 0 0-.09.24v3.32a.34.34 0 0 0 .57.27a3.23 3.23 0 0 1 1.93-.67C6.61 6.22 6.85 7 8 7a3.08 3.08 0 0 0 1.83-.84a.36.36 0 0 0 .17-.32V2.37A.35.35 0 0 0 9.5 2a3.13 3.13 0 0 1-1.5.67C6.85 2.67 6.65 2 5.5 2zm-4-.5a1 1 0 1 1 0 2a1 1 0 0 1 0-2zm.5 3v6a.5.5 0 0 1-1 0v-6a.5.5 0 0 1 1 0z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmergencyPhone(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.875 11.49a.51.51 0 0 0 .72 0l.72-.72l2.18 2.16l-.37.37a2.24 2.24 0 0 1-1.44.7H8.24a2.24 2.24 0 0 1-1.45-.7L1.72 8.23A2.24 2.24 0 0 1 1 6.78V5.33a2.24 2.24 0 0 1 .72-1.45l.36-.36l2.18 2.17l-.73.73a.51.51 0 0 0 0 .72Zm4.72.38a1 1 0 0 0 .036-1.414l-.036-.036l-.72-.72a1 1 0 0 0-1.414-.036l-.036.036Zm-7.28-7.25a1 1 0 0 0 .036-1.414l-.756-.756a1 1 0 0 0-1.414-.036l-.041.036ZM10 2v2H8v1h2v2h1V5h2V4h-2V2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmergencyPhoneEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M8.87 8.53a.73.73 0 0 0 0-1l-.74-.74a.73.73 0 0 0-1 0zm-4.6-4.56a.73.73 0 0 0 0-1l-.71-.71a.73.73 0 0 0-1 0zM3.04 5.65l2.31 2.31a.37.37 0 0 0 .52 0l.44-.43l1.76 1.74a2.27 2.27 0 0 1-1.34.73h-1a1.345 1.345 0 0 1-1-.52L1.52 6.27a1.345 1.345 0 0 1-.52-1v-1a2.27 2.27 0 0 1 .73-1.34l1.74 1.76l-.43.44a.37.37 0 0 0 0 .52M8 2.5H6.5v1H8V5h1V3.5h1.5v-1H9V1H8z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Entrance(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 6.5v-1a1 1 0 0 0-2 0v3zM6 3a1 1 0 1 1 2 0a1 1 0 0 1-2 0m9 3a1 1 0 0 1-1 1h-1.58a1 1 0 0 0-.71.29l-5.42 5.42a1 1 0 0 1-.7.29H3a1 1 0 0 1 0-2h1.59a1 1 0 0 0 .7-.29l5.42-5.42a1 1 0 0 1 .71-.29H14a1 1 0 0 1 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EntranceAltOne(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.554 9.639a.5.5 0 0 0 .707.707l2.667-2.677a.25.25 0 0 0 0-.354L7.261 4.639a.5.5 0 0 0-.707.707L8.2 7H1.5a.5.5 0 0 0 0 1h6.7ZM12 1H5.5a.5.5 0 0 0 0 1h6a.5.5 0 0 1 .5.5v10a.5.5 0 0 1-.5.5H5.25a.5.5 0 0 0 0 1H12a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EntranceAltOneEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M4.793 7.263a.5.5 0 0 0 .707.707l2.243-2.293a.25.25 0 0 0 0-.354L5.489 3.042a.5.5 0 0 0-.707.707L6 5H1.5a.5.5 0 0 0 0 1H6zM9 1H4.5a.5.5 0 0 0 0 1h4a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-.5.5h-4a.5.5 0 0 0 0 1H9a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EntranceEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M4 2.25a.75.75 0 1 1 1.5 0a.75.75 0 0 1-1.5 0zM9.27 4H7.88a.73.73 0 0 0-.52.21l-4 4a.73.73 0 0 1-.51.21H1.73a.73.73 0 1 0 0 1.46h1.89a.73.73 0 0 0 .51-.21l4-4a.73.73 0 0 1 .48-.17h.66a.73.73 0 0 0 .73-.73a.73.73 0 0 0-.73-.77zm-4.52-.5a.75.75 0 0 0-.75.75V6l1.5-1.5v-.25a.75.75 0 0 0-.75-.75z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EntranceFifteen(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M7 6.5v-1a1 1 0 1 0-2 0v3l2-2zm-2.35 4.06L5 3a1 1 0 1 1 2 0a1 1 0 0 1-2 0l-.35 7.56z" fill="currentColor"/><path d="M14 6a1 1 0 0 1-1 1h-1.58a1 1 0 0 0-.71.29l-5.42 5.42a1 1 0 0 1-.7.29H2a1 1 0 1 1 0-2h1.59a1 1 0 0 0 .7-.29l5.42-5.42a1 1 0 0 1 .71-.29H13a1 1 0 0 1 1 1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Farm(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 6L5 4L2 6L1 8v4h2v-2h4v2h2V8zM6 8H4V6h2zm8 4h-3V2.5a1.5 1.5 0 1 1 3 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FarmEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M6 5L4 4L2 5L1 7v3h1.5V8h3v2H7V7zM5 7H3V5.5h2V7z" fill="currentColor"/><path d="M10 2a1 1 0 0 0-2 0v8h2V2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastFood(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 8a1 1 0 0 1-1 1H2a1 1 0 1 1 0-2h11a1 1 0 0 1 1 1M3.5 10H2a3 3 0 0 0 3 3h5a3 3 0 0 0 3-3zM3 6H2V4a2 2 0 0 1 2-2h7a2 2 0 0 1 2 2v2zm8-1.5a.5.5 0 1 0 1 0a.5.5 0 0 0-1 0m-2-1a.5.5 0 1 0 1 0a.5.5 0 0 0-1 0m-2 1a.5.5 0 1 0 1 0a.5.5 0 0 0-1 0m-2-1a.5.5 0 1 0 1 0a.5.5 0 0 0-1 0m-2 1a.5.5 0 1 0 1 0a.5.5 0 0 0-1 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastFoodEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M10 8a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2h9zm0-3H1a1 1 0 1 0 0 2h9a1 1 0 1 0 0-2zM8.55 1H2.46A1.46 1.46 0 0 0 1 2.46V4h9V2.47A1.46 1.46 0 0 0 8.55 1zm-5 2A.5.5 0 1 1 4 2.5a.5.5 0 0 1-.5.5h.05zm4 0A.5.5 0 1 1 8 2.5a.5.5 0 0 1-.5.5h.05z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fence(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 9H13V6h.5a.5.5 0 0 0 0-1H13V3l-.286-.573a.249.249 0 0 0-.424-.006L12 3v2h-1V3l-.286-.573a.249.249 0 0 0-.424-.006L10 3v2H9V3l-.286-.573a.25.25 0 0 0-.424-.006L8 3v2H7V3l-.286-.573a.25.25 0 0 0-.424-.006L6 3v2H5V3l-.286-.573a.25.25 0 0 0-.424-.006L4 3v2H3V3l-.286-.573a.25.25 0 0 0-.424-.006L2 3v2h-.5a.5.5 0 1 0 0 1H2v3h-.5a.5.5 0 1 0 0 1H2v1.5a.5.5 0 0 0 1 0V10h1v1.5a.5.5 0 0 0 1 0V10h1v1.5a.5.5 0 0 0 1 0V10h1v1.5a.5.5 0 0 0 1 0V10h1v1.5a.5.5 0 0 0 1 0V10h1v1.5a.5.5 0 0 0 1 0V10h.5a.5.5 0 0 0 0-1M3 9V6h1v3zm2 0V6h1v3zm2 0V6h1v3zm2 0V6h1v3zm2 0V6h1v3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FenceEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.5 7H9V5h.5a.5.5 0 0 0 0-1H9V3l-.278-.555a.254.254 0 0 0-.443 0L8 3v1H7V3l-.278-.555a.254.254 0 0 0-.443 0L6 3v1H5V3l-.278-.555a.254.254 0 0 0-.443 0L4 3v1H3V3l-.278-.555a.254.254 0 0 0-.443 0L2 3v1h-.5a.5.5 0 0 0 0 1H2v2h-.5a.5.5 0 0 0 0 1H2v.5a.5.5 0 0 0 1 0V8h1v.5a.5.5 0 0 0 1 0V8h1v.5a.5.5 0 0 0 1 0V8h1v.5a.5.5 0 0 0 1 0V8h.5a.5.5 0 0 0 0-1zM3 7V5h1v2zm2 0V5h1v2zm2 0V5h1v2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ferry(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.33 11a3 3 0 0 1 1.85.67l.26.23h.05l.31-.28a2.89 2.89 0 0 1 3.44-.18L13.5 7L12 6V2.45A1.54 1.54 0 0 0 10.5 1H10a.94.94 0 0 0-1-1H6a.94.94 0 0 0-1 1h-.5A1.54 1.54 0 0 0 3 2.45V6L1.5 7l2.25 4.53A2.93 2.93 0 0 1 5.33 11M4 2.45A.47.47 0 0 1 4.5 2h6a.47.47 0 0 1 .5.5v3l-3.5-2L4 5.45zM14 13v1a1 1 0 0 0-.68.34a2.15 2.15 0 0 1-1 .6a1.91 1.91 0 0 1-1.77-.37l-.39-.35a.78.78 0 0 0-1.1 0c-.14.12-.27.25-.42.37a1.92 1.92 0 0 1-2.48-.07l-.39-.35a.76.76 0 0 0-1 0c-.19.15-.36.32-.55.47a1.91 1.91 0 0 1-2.35-.06l-.31-.27A.94.94 0 0 0 1 14v-1c.259-.032.52.017.75.14a6.2 6.2 0 0 1 .79.58a.84.84 0 0 0 .81.25a.93.93 0 0 0 .31-.16c.16-.12.29-.26.45-.39a1.92 1.92 0 0 1 2.45 0c.16.13.31.28.47.41a.76.76 0 0 0 1 0c.16-.13.31-.28.47-.41a1.92 1.92 0 0 1 2.46 0c.15.13.29.27.44.39a.76.76 0 0 0 1 0l.47-.41c.32-.26.719-.4 1.13-.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FerryEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M11 9.45v.77a.83.83 0 0 0-.57.26a1.86 1.86 0 0 1-.81.46a1.74 1.74 0 0 1-1.49-.29l-.33-.27a.71.71 0 0 0-.93 0c-.12.09-.23.2-.35.28a1.75 1.75 0 0 1-2.1-.05l-.33-.27a.69.69 0 0 0-.83 0c-.16.11-.3.25-.46.36a1.75 1.75 0 0 1-2 0l-.26-.21a.82.82 0 0 0-.54-.27v-.77a1.16 1.16 0 0 1 .63.11c.234.13.458.276.67.44a.76.76 0 0 0 .7.2a.82.82 0 0 0 .26-.12c.13-.09.25-.2.38-.3a1.75 1.75 0 0 1 2.07 0l.4.31a.69.69 0 0 0 .81 0l.4-.31a1.75 1.75 0 0 1 2.08 0l.38.3a.69.69 0 0 0 .82 0l.4-.31a1.61 1.61 0 0 1 1-.32zM2.61 7.61L1 5l1-.91V1.15A1.18 1.18 0 0 1 3.19 0h4.62A1.18 1.18 0 0 1 9 1.15v2.94L10 5L8.39 7.62a2.22 2.22 0 0 0-2.65.14L5.5 8l-.2-.18a2.22 2.22 0 0 0-2.69-.21zm.2-4L5.5 2l2.69 1.63V1.15a.36.36 0 0 0-.38-.38H3.19a.36.36 0 0 0-.38.38v2.48v-.02z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FerryJp(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.5 12.5h-4V8c0-1.088.58-2.552 1.354-3.862c.22-.371.44-.708.646-.993c.205.285.427.622.646.993C8.92 5.448 9.5 6.912 9.5 8zM7.5 1C7 1 4 5 4 8v4.5A1.5 1.5 0 0 0 5.5 14h4a1.5 1.5 0 0 0 1.5-1.5V8c0-3-3-7-3.5-7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FireStation(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 14c3.59 0 6.5-3 6.5-6.5c0-3-2.5-5.5-2.5-5.5l-1 3.5l-3-4.5l-3 4.5l-1-3.5S1 4.5 1 7.5C1 11 3.91 14 7.5 14m0-1.5A2.5 2.5 0 0 1 5 10c0-1.38 2.5-4.5 2.5-4.5S10 8.62 10 10a2.5 2.5 0 0 1-2.5 2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FireStationEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.5 0l-2 4L2 2c-.405.712-2 2.167-2 4c0 2.7 2.8 5 5.5 5S11 8.7 11 6c0-1.833-1.595-3.288-2-4L7.5 4l-2-4zm0 5.5s2 1.585 2 3c0 .611-.778 1.278-2 1.278s-2-.667-2-1.278c0-1.366 2-3 2-3z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FireStationJp(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 8.149v3.601a1 1 0 0 1-2 0V8.149a5.008 5.008 0 0 1-4-4.899a1 1 0 0 1 2 0a3 3 0 0 0 6 0a1 1 0 0 1 2 0a5.008 5.008 0 0 1-4 4.899"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FireStationJpEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M6.375 5.887v2.738a.875.875 0 0 1-1.75 0V5.887A3.75 3.75 0 0 1 1.75 2.25a.75.75 0 0 1 1.5 0a2.25 2.25 0 0 0 4.5 0a.75.75 0 0 1 1.5 0a3.75 3.75 0 0 1-2.875 3.637z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FitnessCentre(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 7v1h-1v2h-1v1H11V8H4v3H2.5v-1h-1V8h-1V7h1V5h1V4H4v3h7V4h1.5v1h1v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FitnessCentreEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M7 6H4V5h3zM2 3v1H1v1a.5.5 0 0 0 0 1v1h1v1h1.5V3zm8 2V4H9V3H7.5v5H9V7h1V6a.5.5 0 0 0 0-1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Florist(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 1A2.7 2.7 0 0 1 4 3l.5-3L5 3a2.7 2.7 0 0 1 2.5-2c-.2.03-1 .26-1 2v.083A1.959 1.959 0 0 1 4.5 5h-.083A1.959 1.959 0 0 1 2.5 3c0-1.739-.8-1.97-1-2m10.512 10l3-.5l-3-.5A2.686 2.686 0 0 0 14 7.5c-.03.2-.248 1-1.988 1a2 2 0 0 0 0 4c1.739 0 1.958.8 1.988 1a2.686 2.686 0 0 0-1.988-2.5M9.688 5.548a1 1 0 0 1 0-2a1 1 0 0 1 0-2a.986.986 0 0 1 .852.507l.023-.012a.978.978 0 0 1-.116-.444a1 1 0 1 1 2 0a.978.978 0 0 1-.116.444l.019.01a.986.986 0 0 1 .85-.5a1 1 0 0 1 .025 2a1 1 0 0 1-.025 2a.986.986 0 0 1-.85-.5l-.016.009a.978.978 0 0 1 .113.44a1 1 0 0 1-2 0a.978.978 0 0 1 .113-.44l-.02-.011a.986.986 0 0 1-.852.507Zm.71-1.995A1.051 1.051 0 1 0 11.449 2.5A1.051 1.051 0 0 0 10.4 3.553Zm-5.452 7.891l-.516-.515l4.462-4.454a1.746 1.746 0 0 1-.452-.255l-4.364 4.355l-.519-.518A16.051 16.051 0 0 0 4.912 6a3.373 3.373 0 0 1-.412.035c-.041 0-.073-.008-.112-.01a16.953 16.953 0 0 1-1.257 3.606L2.76 9.26a.246.246 0 0 0-.4.079L.231 14.445a.287.287 0 0 0-.016.089a.25.25 0 0 0 .25.25a.289.289 0 0 0 .1-.019l5.1-2.124a.246.246 0 0 0 .079-.4l-.372-.372a16.874 16.874 0 0 1 3.612-1.256c0-.059-.015-.106-.015-.166A3.349 3.349 0 0 1 9 10.089a16.076 16.076 0 0 0-4.054 1.355"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FloristEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M1 2.5a3.1 3.1 0 0 0-1-2A3.9 3.9 0 0 1 2 2l.526-2l.437 2a3.9 3.9 0 0 1 2-1.5a3.1 3.1 0 0 0-1 2A1.307 1.307 0 0 1 2.662 4h-.253A1.36 1.36 0 0 1 1 2.5zm8 6.482l2-.437l-2-.527a3.9 3.9 0 0 0 1.5-2a3.1 3.1 0 0 1-2 1A1.36 1.36 0 0 0 7 8.427v.253a1.307 1.307 0 0 0 1.5 1.3a3.1 3.1 0 0 1 2 1A3.9 3.9 0 0 0 9 8.982zM3.4 7.9l2.713-2.719a1.382 1.382 0 0 1-.436-.271L3.044 7.543l-.279-.279A11.045 11.045 0 0 0 3 4.5h-.5a15.272 15.272 0 0 1-.161 2.338l-.088-.088a.247.247 0 0 0-.4.071L.059 10.657a.27.27 0 0 0-.026.108a.25.25 0 0 0 .25.25a.27.27 0 0 0 .089-.021L.378 11l3.8-1.85a.247.247 0 0 0 .068-.4l-.063-.063A12.768 12.768 0 0 1 6.5 8.5V8a10.381 10.381 0 0 0-2.724.275zm3.843-5.378a.75.75 0 0 1 .018-1.5a.739.739 0 0 1 .561.266a.75.75 0 1 1 1.383 0a.739.739 0 0 1 .561-.266a.75.75 0 0 1 .014 1.5a.748.748 0 1 1-.561 1.26a.728.728 0 0 1 .044.218a.75.75 0 1 1-1.5 0a.737.737 0 0 1 .05-.238a.739.739 0 0 1-.558.26a.75.75 0 0 1-.012-1.5zm.518-.022a.75.75 0 1 0 .75-.75a.75.75 0 0 0-.75.75z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fuel(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 6v5.5a.5.5 0 0 1-1 0v-2A1.5 1.5 0 0 0 11.5 8H10V2a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V9h1.5a.5.5 0 0 1 .5.5v2a1.5 1.5 0 0 0 3 0V5a1 1 0 0 0-1-1V2.49a.5.5 0 0 0-.5-.49a.51.51 0 0 0-.5.55V5a1 1 0 1 0 1-1zm-5 .5a.5.5 0 0 1-.5.5h-5a.5.5 0 0 1-.5-.5v-3a.5.5 0 0 1 .5-.5h5a.5.5 0 0 1 .5.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FuelEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.5 3H9V1.5a.5.5 0 0 0-1 0V3a1 1 0 0 0 1 1v4.25a.25.25 0 1 1-.5 0V6.5A1.5 1.5 0 0 0 7 5V2a1 1 0 0 0-1-1H2a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V6a.5.5 0 0 1 .5.5v1.75a1.25 1.25 0 1 0 2.5 0V3.5a.5.5 0 0 0-.5-.5zM6 4.5a.49.49 0 0 1-.48.5h-3A.51.51 0 0 1 2 4.5V3a.5.5 0 0 1 .5-.5h3A.5.5 0 0 1 6 3v1.5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Furniture(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.691 10.142V8.5a.5.5 0 0 0-.5-.5h-1a.5.5 0 0 0-.5.5v1.641a3.99 3.99 0 0 0-2.957 3.272a.507.507 0 0 0 .5.586h6.922a.507.507 0 0 0 .5-.586a3.99 3.99 0 0 0-2.965-3.271m4.639-3.863l-2.5-5A.5.5 0 0 0 10.383 1H4.999a.5.5 0 0 0-.446.276l-2.5 5A.5.5 0 0 0 2.497 7h8.194v1.5a.5.5 0 0 0 1 0V7h1.194a.5.5 0 0 0 .445-.721"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FurnitureEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M6 7.555V6a.5.5 0 1 0-1 0v1.555a2.5 2.5 0 0 0-1.923 1.827a.502.502 0 0 0 .489.618h3.868a.502.502 0 0 0 .49-.618A2.5 2.5 0 0 0 6 7.555z" fill="currentColor"/><path d="M9.135 4.27L7.64 1.279A.505.505 0 0 0 7.188 1H3.812a.505.505 0 0 0-.451.279l-1.496 2.99A.505.505 0 0 0 2.317 5H7v.75a.25.25 0 1 0 .5 0V5h1.183c.376 0 .62-.395.452-.73z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FurnitureFifteen(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9 10.142v-1.64A.501.501 0 0 0 8.499 8H7.5A.501.501 0 0 0 7 8.501v1.64a3.993 3.993 0 0 0-2.957 3.273a.507.507 0 0 0 .496.586h6.922c.31 0 .541-.28.496-.586A3.993 3.993 0 0 0 9 10.142z" fill="currentColor"/><path d="M13.64 6.279l-2.502-5.004A.498.498 0 0 0 10.692 1H5.308a.498.498 0 0 0-.446.276L2.361 6.279A.498.498 0 0 0 2.806 7H11v1.5a.5.5 0 0 0 1 0V7h1.194c.37 0 .611-.39.445-.721z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gaming(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.1 11.5c-.6.3-1.4.1-1.8-.5l-1.1-1.4H4.8L3.7 11c-.5.7-1.4.8-2.1.3c-.5-.4-.7-1-.6-1.5l.7-3.7C1.9 4.9 3 4 4.2 4H7V2.5C7 1.7 7.6 1 8.4 1h3.1c.3 0 .5.2.5.5s-.2.5-.5.5h-3c-.3 0-.5.2-.5.4V4h2.8c1.2 0 2.3.9 2.5 2.1l.7 3.7c.1.7-.2 1.4-.9 1.7M6 6.5C6 5.7 5.3 5 4.5 5S3 5.7 3 6.5S3.7 8 4.5 8S6 7.3 6 6.5m6 0c0-.3-.2-.5-.5-.5H11v-.5c0-.3-.2-.5-.5-.5s-.5.2-.5.5V6h-.5c-.3 0-.5.2-.5.5s.2.5.5.5h.5v.5c0 .3.2.5.5.5s.5-.2.5-.5V7h.5c.3 0 .5-.2.5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GamingEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.715 5.8a2.046 2.046 0 0 0-2-1.8h-1.7V2.5c0-.2.2-.5.4-.5h2.1a.472.472 0 0 0 .5-.5a.472.472 0 0 0-.5-.5h-2a1.453 1.453 0 0 0-1.5 1.4V4h-1.8a2.046 2.046 0 0 0-2 1.8l-.2 2.8a.991.991 0 0 0 .8 1.1a1.613 1.613 0 0 0 .9-.3l1.4-1.4h2.8l1.4 1.4a1.071 1.071 0 0 0 1.4.1a1.613 1.613 0 0 0 .3-.9zM4.015 7h-2V6a.945.945 0 0 1 1-1h1zm5 0h-2V5h1a1 1 0 0 1 1 1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Garden(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 8c0 3.31-2.19 6-5.5 6S2 11.31 2 8a5.33 5.33 0 0 1 5 3.61V7H4.5A1.5 1.5 0 0 1 3 5.5v-3a.5.5 0 0 1 .9-.3l1.53 2l1.65-3a.5.5 0 0 1 .84 0l1.65 3l1.53-2a.5.5 0 0 1 .9.3v3A1.5 1.5 0 0 1 10.5 7H8v4.61A5.33 5.33 0 0 1 13 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GardenCentre(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 5v-.5a2.5 2.5 0 0 1 5 0v5.793l2.365-2.365l-.347-1.295l-.001-.006h-.001a.5.5 0 0 1 .838-.481l2 2a.5.5 0 0 1-.479.838l-.01-.003l-1.293-.346L9 11.707V12a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-.464L1.732 9.268a2.503 2.503 0 0 1 0-3.536A2.493 2.493 0 0 1 3.5 5zm0 1h-.5a1.5 1.5 0 0 0-1.061 2.561L4 10.121zm4-1v-.5a1.5 1.5 0 0 0-3 0V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GardenCentreEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M10.875 5.164l-.007-.008l-.029-.029l-.971-.97A.5.5 0 0 0 9 4.494v1.148L7 7.646V3a2 2 0 0 0-4 0a2 2 0 0 0 0 4v1a1 1 0 0 0 1 1h2a.984.984 0 0 0 .833-.48L9.36 5.99h1.061a.5.5 0 0 0 .453-.826zM1.5 5A1.5 1.5 0 0 1 3 3.5v3A1.5 1.5 0 0 1 1.5 5zm2-2a1.5 1.5 0 0 1 3 0h-3z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GardenEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M10 6a4 4 0 0 1-4.25 4A4 4 0 0 1 1.5 6a4 4 0 0 1 3.9 2.26V5h-2a.89.89 0 0 1-.9-.88V1.84a.35.35 0 0 1 .64-.21L4.28 3L5.45.67a.35.35 0 0 1 .6 0L7.22 3l1.13-1.38a.35.35 0 0 1 .65.21v2.28a.89.89 0 0 1-.89.89H6.1v3.26A4 4 0 0 1 10 6z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gate(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 2.75A.75.75 0 0 1 1.75 2h1.114c.09 0 .18.016.263.048l3.386 1.27A.75.75 0 0 1 7 4.02v5.96a.75.75 0 0 1-.487.702L3 12v.25a.75.75 0 0 1-.75.75h-.5a.75.75 0 0 1-.75-.75zm3 5.736v-4.87a.25.25 0 0 0-.162-.234l-.5-.187A.25.25 0 0 0 3 3.429v5.224c0 .17.167.291.33.237l.5-.166A.25.25 0 0 0 4 8.486M5 4.18v3.807a.25.25 0 0 0 .33.238l.5-.167A.25.25 0 0 0 6 7.82V4.366a.25.25 0 0 0-.162-.234l-.5-.187A.25.25 0 0 0 5 4.179Zm9-1.43a.75.75 0 0 0-.75-.75h-1.114a.75.75 0 0 0-.263.048l-3.386 1.27A.75.75 0 0 0 8 4.02v5.96a.75.75 0 0 0 .487.702L12 12v.25c0 .414.336.75.75.75h.5a.75.75 0 0 0 .75-.75zm-2.83 5.974a.25.25 0 0 1-.17-.238v-4.87a.25.25 0 0 1 .162-.234l.5-.187a.25.25 0 0 1 .338.234v5.224a.25.25 0 0 1-.33.237zM10 4.179v3.807a.25.25 0 0 1-.33.238l-.5-.167A.25.25 0 0 1 9 7.82V4.366a.25.25 0 0 1 .162-.234l.5-.187a.25.25 0 0 1 .338.234"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gift(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 1.002a3 3 0 0 0-3 2.25a3 3 0 0 0-3-2.25C2.768.947 2.235 2.797 2.818 4H1.5a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5h12a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5h-1.306c.606-1.185.098-2.999-1.654-2.999zM10.683 4H8.5a2 2 0 0 1 2-1.998c1.175 0 1.383 1.872.183 1.998M6.5 4H4.292c-1.035-.117-1.096-1.894.04-1.998a.921.921 0 0 1 .168 0A2 2 0 0 1 6.5 4M2 7.001v4.5a1.5 1.5 0 0 0 1.5 1.5h3v-6zm6.5 0v6h3a1.5 1.5 0 0 0 1.5-1.5v-4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GiftEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M0 5h4.5v1H0V5zm1 4.79c0 .668.542 1.21 1.21 1.21H4.52V7H1v2.79zM7.64 4H3.36A1.26 1.26 0 0 1 2 2.5A1.41 1.41 0 0 1 3.42 1a2.12 2.12 0 0 1 2.1 1.69A2.13 2.13 0 0 1 7.62 1A1.43 1.43 0 0 1 9 2.52A1.27 1.27 0 0 1 7.64 4zm-2.82-.75a1.43 1.43 0 0 0-1.4-1.5a.68.68 0 0 0-.7.75a.68.68 0 0 0 .606.747c.031.003.063.004.094.003h1.4zm2.8 0a.68.68 0 0 0 .7-.75a.71.71 0 0 0-.7-.75a1.43 1.43 0 0 0-1.4 1.5h1.4zM6.52 5v1H11V5H6.52zm0 6h2.29A1.21 1.21 0 0 0 10 9.79V7H6.52v4z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globe(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.981 9.481l-.721-.721a4.979 4.979 0 1 1-7.02-7.02l-.721-.721A5.991 5.991 0 0 0 7 11.475V13h-.5a.5.5 0 0 0 0 1h2a.5.5 0 0 0 0-1H8v-1.527a5.983 5.983 0 0 0 3.981-1.992M7.5 9A3.5 3.5 0 1 0 4 5.5A3.5 3.5 0 0 0 7.5 9m1-5l.364-.592a2.5 2.5 0 0 1 .823.9L9.5 5h-1Zm-1.844-.844L7.5 4v1L8 6h1.94a2.494 2.494 0 0 1-1.5 1.814L8 7H6.5L5.05 5a2.5 2.5 0 0 1 1.606-1.844"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M6.5 10H6V8.95a4.478 4.478 0 0 0 2.682-1.268L8.15 7.15a3.739 3.739 0 0 1-2.65 1.1A3.754 3.754 0 0 1 1.75 4.5c0-1.034.42-1.971 1.1-2.65l-.532-.532A4.486 4.486 0 0 0 1 4.5c0 2.314 1.753 4.198 4 4.45V10h-.5a.5.5 0 1 0 0 1h2a.5.5 0 0 0 0-1z" fill="currentColor"/><path d="M5.5 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5zm-.545-4.415l.493.616v.837l.613 1.111h1.322c-.17.492-.522.892-.98 1.126L5.5 5.5h-1l-.968-1.32a1.991 1.991 0 0 1 1.423-1.595z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeFifteen(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M11.98 9.48l-.72-.72c-.918 1.057-2.254 1.74-3.76 1.74c-2.757 0-5-2.243-5-5c0-1.506.683-2.842 1.74-3.76l-.72-.72A5.978 5.978 0 0 0 1.5 5.5c0 3.145 2.42 5.72 5.5 5.975V13h-.5a.5.5 0 1 0 0 1h2a.5.5 0 0 0 0-1H8v-1.527a5.982 5.982 0 0 0 3.98-1.992z" fill="currentColor"/><path d="M7.5 9a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7zm1-5l.364-.592a2.5 2.5 0 0 1 .823.905L9.5 5h-1V4zm-1.844-.844L7.5 4v1L8 6h1.94a2.495 2.495 0 0 1-1.502 1.814L8 7H6.5L5.05 5a2.496 2.496 0 0 1 1.606-1.844z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Golf(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.4 1.1v.2c0 .4.3.7.7.7c.3 0 .5-.2.6-.5l.2-.5l5.6 2.3L6.6 6c-.4.3-.4.7-.3 1.1l.9 2.1l-1.3 3.9c-.2.5.2.9.6.9c.3 0 .5-.1.6-.5l1.4-4l.1.3v3.5s0 .7.7.7s.7-.7.7-.7V10c0-.2 0-.3-.1-.5L8.5 6.1l2.7-1.9c.2-.2.4-.3.4-.6s-.2-.5-.4-.6L4 .1c-.088 0-.118.018-.2.1zM5.5 3C4.7 3 4 3.7 4 4.5S4.7 6 5.5 6S7 5.3 7 4.5S6.3 3 5.5 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GolfEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M4.05.638c-.518.543.223 1.175.678.675L5.224.77l2.332 1.393L3.306 4.5c-.293.162-.341.44-.234.721l.889 2.317l-.935 2.804c-.129.385.171.649.474.658a.468.468 0 0 0 .475-.341l.92-2.764l.853-.283l.252.505V10.5s0 .5.5.5s.5-.5.5-.5V8.117a.84.84 0 0 0-.064-.362l-1.281-3.34l2.552-1.403c.187-.08.29-.254.29-.512c0-.226-.217-.413-.456-.556L4.905.071c-.16-.096-.275-.043-.323.007l-.533.56zM3 2a1 1 0 1 0 0 2a1 1 0 0 0 0-2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grocery(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.2 1.5s-1.391-.041-1.946.5c-.534.52-.754.918-.754 2H1.2l1.394 4.814c.003.008.01.015.013.022c.235.657.848 1.13 1.579 1.158l.013.006h6.5v.2s.001.3-.199.7c-.2.4-.3.6-1.1.6H2.9c-1 0-1 1.5 0 1.5h6.4c1.2 0 2.1-.7 2.4-1.4c.3-.7.3-1.3.3-1.3V4c0-.524.229-1 .7-1h.55a.75.75 0 0 0 0-1.5zM9.2 13c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m-5 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GroceryEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.75 1.5c-.004 0-.318-.006-.63.188C8.803 1.881 8.5 2.3 8.5 3H1l.75 3.5c.107.5.75.5.75.5h6s-.003.247-.152.496S7.93 8 7.25 8H2c-.338-.005-.338.505 0 .5h5.25c.82 0 1.302-.37 1.527-.746C9.003 7.378 9 7 9 7V3c0-.567.196-.772.38-.887C9.567 2 9.747 2 9.747 2h.504c.338.005.338-.505 0-.5h-.5zm-2.5 7a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5zm-4 0a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hairdresser(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 3s-2-.6-3.5.5l-4.3 3c-.8-.6-2-1.3-3.2-1.7V4c0-1.1-.9-2-2-2s-2 .9-2 2v1.5c0 .5.5.5.5.5h2C4.5 6 6 7.5 6 7.5S4.5 9 2.5 9h-2S0 9 0 9.5V11c0 1.1.9 2 2 2s2-.9 2-2v-.8c1.2-.4 2.4-1.1 3.2-1.7l4.3 3c1.5 1.1 3.5.5 3.5.5L8.5 7.5zM3 5H1V4c0-.6.4-1 1-1s1 .4 1 1zm0 6c0 .6-.4 1-1 1s-1-.4-1-1v-1h2zm4.25-3a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HairdresserEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M1.5 2A1.5 1.5 0 0 0 0 3.5v1c0 .5.5.5.5.5h1C3 5 4 5.5 4 5.5S3 6 1.5 6h-1S0 6 0 6.5v1a1.5 1.5 0 0 0 3 0v-.615c.808-.158 1.587-.453 2.225-.742L8.5 8C10 8.75 11 8 11 8L6.5 5.5L11 3s-1-.75-2.5 0L5.225 4.857C4.587 4.568 3.808 4.273 3 4.115v-.611V3.5A1.5 1.5 0 0 0 1.5 2zm0 1a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1zm3.732 2.25h.018a.25.25 0 1 1-.018 0zM1.5 7a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Harbor(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 0C5.5 0 4 1.567 4 3.5a3.492 3.492 0 0 0 2.5 3.338v6.039c-.93-.165-1.875-.55-2.648-1.27c-1.053-.98-1.85-2.54-1.85-5.109a1 1 0 1 0-2.002 0c0 3.003 1.012 5.196 2.49 6.572C3.97 14.447 5.838 15 7.5 15c1.666 0 3.535-.56 5.012-1.94s2.486-3.573 2.486-6.562c.065-1.395-2.063-1.395-1.998 0c0 2.553-.8 4.115-1.853 5.1c-.774.722-1.718 1.11-2.647 1.277V6.842A3.494 3.494 0 0 0 11 3.5C11 1.567 9.5 0 7.5 0m0 2a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HarborEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.5 0A2.5 2.5 0 0 0 3 2.5a2.493 2.493 0 0 0 1.75 2.371v4.545c-.659-.115-1.338-.375-1.893-.857C2.09 7.89 1.5 6.829 1.5 5A.75.75 0 1 0 0 5c0 2.17.773 3.735 1.873 4.691S4.333 11 5.5 11s2.527-.352 3.627-1.309S11 7.171 11 5c.014-1.014-1.514-1.014-1.5 0c0 1.83-.589 2.89-1.357 3.559c-.555.482-1.234.742-1.893.857V4.875A2.494 2.494 0 0 0 5.5 0zm0 1.5a1 1 0 1 1 0 2a1 1 0 0 1 0-2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hardware(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.792 3.262s-1.676 1.675-2.116 2.12a.296.296 0 0 1-.282.097c-.149-.023-.298-.038-.447-.056l-1.229-.146c-.069-.592-.14-1.172-.2-1.753c-.006-.055.033-.13.074-.172c.508-.513 1.927-1.928 2.134-2.136c-.45-.213-1.224-.297-1.8-.203a3.545 3.545 0 0 0-2.82 4.625c.037.115.013.182-.069.264c-1.848 1.843-3.693 3.69-5.539 5.535a2.569 2.569 0 0 0-.178.192a1.457 1.457 0 0 0 .298 2.102c.604.42 1.365.344 1.905-.195c1.853-1.852 3.706-3.705 5.556-5.56c.092-.092.165-.123.299-.078a3.39 3.39 0 0 0 1.888.089a3.544 3.544 0 0 0 2.655-4.304c-.031-.13-.078-.259-.13-.421"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HardwareEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M1.217 8.365a1.008 1.008 0 0 0 .206 1.453c.418.29.944.238 1.317-.134c1.281-1.28 2.562-2.561 3.841-3.844c.064-.064.114-.085.207-.054c.428.146.866.163 1.306.061a2.45 2.45 0 0 0 1.835-2.975c-.023-.094-.057-.185-.094-.305c-.035.044-.997 1.013-1.459 1.48a.205.205 0 0 1-.195.067C8.08 4.098 7.287 4.034 7 4c-.048-.41-.073-.836-.115-1.238c-.004-.038.023-.09.051-.119c.35-.354 1.332-1.333 1.476-1.476c-.312-.147-.847-.206-1.245-.14a2.45 2.45 0 0 0-1.95 3.197c.026.08.01.126-.047.182l-3.953 3.96z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.91 7.75c-1.17 2.25-4.3 5.31-6.07 6.94a.5.5 0 0 1-.67 0C5.39 13.06 2.26 10 1.09 7.75C-1.48 2.8 5-.5 7.5 4.45C10-.5 16.48 2.8 13.91 7.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M10.06 4.76a27 27 0 0 1-4.2 5.36a.49.49 0 0 1-.71 0A27.003 27.003 0 0 1 .94 4.76C-.88 1.12 3.74-1.31 5.5 2.33c1.76-3.64 6.38-1.21 4.56 2.43z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heliport(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2C3 2 3 3 4 3h4v1c-.277 0-.5.223-.5.5V5H3.932A1.998 1.998 0 0 0 0 5.5a2 2 0 0 0 2 2a2 2 0 0 0 1.053-.303L5.5 10.5C6.507 11.95 8.318 12 9 12h5s1 0 1-1v-.994c0-.733-.126-1.132-.5-1.506l-3-3s-.592-.5-1.273-.5H9.5v-.5c0-.277-.223-.5-.5-.5V3h4c1 0 1-1 0-1zM2 4.5a1 1 0 1 1 0 2a1 1 0 0 1 0-2M10 6c.5 0 .79.323 1 .5L13.5 9H10S9 9 9 8V7s0-1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeliportEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M3 1c-.277 0-.5.223-.5.5s.224.482.5.5h3v2H2.912a1.5 1.5 0 1 0-.39 1.596L4.5 8.5c.681 1 1.5 1 2 1h3.535s.965 0 .965-1v-1c0-.5-.5-1-.5-1l-2-2S8 4 7.5 4H7V2h3a.499.499 0 1 0 0-1H3zM1.5 4a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1zm6.25 1s.25 0 .75.5L10 7H7.5S7 7 7 6.5v-1s0-.5.5-.5h.25z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HighwayRestArea(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 13h-3V9h3a.5.5 0 0 0 .41-.787L11.66 5h.84a.5.5 0 0 0 .384-.82l-2.5-3a.516.516 0 0 0-.768 0l-2.5 3A.5.5 0 0 0 7.5 5h.84L6.09 8.213A.5.5 0 0 0 6.5 9h3v4H4v-2h1.5a.5.5 0 0 0 0-1h-4a.5.5 0 0 0 0 1H3v2H1.5a.5.5 0 0 0 0 1h12a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Historic(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 1c-1 0-1 1-2 1H2c-.545 0-1 .455-1 1v7c0 .545.455 1 1 1h5v2.5s0 .5.5.5s.5-.5.5-.5V11h5c.545 0 1-.455 1-1V3c0-.545-.455-1-1-1H9.5c-1 0-1-1-2-1M3 5V4h9v1zm0 1h4v1H3zm0 2h7v1H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 13.748c0 .138.112.25.25.25h3.749v-3h3v3h3.749a.25.25 0 0 0 .25-.25v-5.75H2zm11.93-7.17l-.932-.82V2a1 1 0 1 0-2 0v2L7.681 1.09a.25.25 0 0 0-.353-.011l-.011.011l-6.25 5.463a.25.25 0 0 0 .18.42L3 7h10.747a.25.25 0 0 0 .183-.421"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M10.002 4.75a.25.25 0 0 1-.25.25H1.25A.25.25 0 0 1 1 4.75c0-.07.028-.14.08-.19l4.238-3.454l.016-.016a.248.248 0 0 1 .35.016L8 2.996V1.5a.5.5 0 0 1 1 0v2.31l.92.75c.052.05.081.119.08.19zM2 9.752a.248.248 0 0 0 .246.25h2.755v-2h1v2h2.752a.248.248 0 0 0 .248-.248V6.001H2v3.75z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HorseRiding(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 1a1 1 0 1 1 1 1a1 1 0 0 1-1-1m2 2.5a.484.484 0 0 0 0-.058A.472.472 0 0 0 7.5 3h-1a.484.484 0 0 0-.058 0A.472.472 0 0 0 6 3.5V7h2Zm6.85 3.644L12.8 4.8l.085-.509a.478.478 0 0 0 .008-.063a.25.25 0 0 0-.25-.25a.346.346 0 0 0-.158.056L9 7H8l1 1v1.5a.5.5 0 0 1-1 0v-1L6 7H4a1.5 1.5 0 0 0-1.243.661A1.466 1.466 0 0 0 1.563 7H1.5A1.449 1.449 0 0 0 0 8.4v.086A3.781 3.781 0 0 0 .559 10.4a.278.278 0 0 0 .191.1a.25.25 0 0 0 .25-.25V9s-.02-.924.753-1c.5-.048.747.253.747.5V11L2 13v1.75a.25.25 0 0 0 .25.25a.254.254 0 0 0 .25-.234V13L4 11v1l.5 2.8a.255.255 0 0 0 .246.2a.25.25 0 0 0 .254-.246L4.855 12.3L5.5 11H8v1l.508 2.813A.257.257 0 0 0 8.75 15a.25.25 0 0 0 .25-.25V12l.367-1a1.155 1.155 0 0 0 .543-.6l2.161-3.767a.863.863 0 0 0 1.023.4l1.066.818a.591.591 0 0 0 .35.135a.487.487 0 0 0 .475-.4a.552.552 0 0 0-.135-.442"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HorseRidingEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M4 1a1 1 0 1 1 1 1a1 1 0 0 1-1-1zm6.88 3.5L9 3v-.5L7 5H6l1 1v1.014a.5.5 0 1 1-1 0V6.5L4 5H3a1 1 0 0 0-.8.446A1.189 1.189 0 0 0 1.247 5A1.076 1.076 0 0 0 0 5.988C0 7.3.635 7.471.635 7.471a.33.33 0 0 0 .115.023A.253.253 0 0 0 1 7.25V6a.49.49 0 0 1 .48-.5h.02A.5.5 0 0 1 2 6v2.014l-.3.6a1.609 1.609 0 0 0-.2.6v1.542a.244.244 0 0 0 .244.244a.255.255 0 0 0 .256-.244V9.514a.367.367 0 0 1 .1-.3l.9-1.2V9l.467 1.816a.256.256 0 0 0 .242.184a.25.25 0 0 0 .25-.25v-.016l-.242-1.61a.6.6 0 0 1 .025-.236l.058-.174L4 8h2v1l.466 1.816a.256.256 0 0 0 .243.184a.25.25 0 0 0 .25-.25l-.241-1.626a.6.6 0 0 1 .025-.236L7 8a.877.877 0 0 0 .934-.661L8.5 4.5a.6.6 0 0 0 .71.454a.562.562 0 0 0 .143-.07l.9.116a.721.721 0 0 0 .392.1a.331.331 0 0 0 .355-.316a.406.406 0 0 0-.12-.284zM6 3.5a.51.51 0 0 0-.5-.5h-1a.482.482 0 0 0-.5.5V5h2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hospital(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 1c-.6 0-1 .4-1 1v4H2c-.6 0-1 .4-1 1v1c0 .6.4 1 1 1h4v4c0 .6.4 1 1 1h1c.6 0 1-.4 1-1V9h4c.6 0 1-.4 1-1V7c0-.6-.4-1-1-1H9V2c0-.6-.4-1-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HospitalEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M10 4H7V1a1.08 1.08 0 0 0-1-1H5a1.08 1.08 0 0 0-1 1v3H1a1.08 1.08 0 0 0-1 1v1a1.08 1.08 0 0 0 1 1h3v3a1.08 1.08 0 0 0 1 1h1a1.08 1.08 0 0 0 1-1V7h3a1.08 1.08 0 0 0 1-1V5a1 1 0 0 0-1-1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HospitalJp(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.504 3H2.496A1.005 1.005 0 0 0 1.5 4.013v5.58a1.015 1.015 0 0 0 .55.907l4.78 2.84a1.474 1.474 0 0 0 1.339 0l4.78-2.84a1.015 1.015 0 0 0 .551-.906v-5.58A1.002 1.002 0 0 0 12.504 3M10.5 8.5h-2v2h-2v-2h-2v-2h2v-2h2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HospitalJpEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.496 1.75H1.504A.505.505 0 0 0 1 2.257v4.53a.507.507 0 0 0 .3.463l3.995 1.956a.501.501 0 0 0 .41 0L9.7 7.213A.507.507 0 0 0 10 6.75V2.257a.505.505 0 0 0-.504-.507zM7.75 6h-1.5v1.5h-1.5V6h-1.5V4.5h1.5V3h1.5v1.5h1.5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HotSpring(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.682 5.759c-.505-.725-.781-1.121-.63-2.376c.139-1.156.678-1.616 1.198-2.059c.404-.345.797-.68.981-1.324c.301.722-.036 1.47-.349 2.164c-.311.691-.599 1.329-.211 1.836c.971 1.271 1.963 2.88.971 4.586c-.214.393-.539.681-.861.968c-.422.375-.841.747-1.004 1.348c-.269-.677.008-1.366.285-2.052c.306-.759.61-1.516.172-2.251c-.203-.34-.389-.607-.551-.84Zm3.371-1.174a2.641 2.641 0 0 1 .368-1.809c.367-.45.701-.642.991-.81c.345-.199.627-.362.83-.877c.289.684 0 1.188-.304 1.718c-.119.207-.24.418-.329.645c-.191.61.101.893.467 1.247c.358.346.786.76.899 1.617A2.375 2.375 0 0 1 12.603 8c-.356.436-.681.655-.965.846c-.356.24-.648.437-.856.965c-.292-.688.031-1.259.353-1.828c.123-.217.246-.434.334-.658c.073-.558-.234-.867-.581-1.217c-.357-.36-.757-.763-.834-1.523Zm-8.018 0c-.098-.632.005-1.3.368-1.809c.376-.528.729-.717 1.035-.881c.323-.173.594-.318.786-.806c.289.683-.014 1.208-.325 1.749c-.122.212-.246.427-.336.655c-.152.385.131 1.014.501 1.242c1.002.801 1.227 2.286.547 3.353c-.366.448-.704.648-.999.823c-.353.209-.645.382-.849.898c-.278-.656.008-1.18.3-1.714c.131-.24.263-.481.345-.737c.179-.557-.134-.884-.51-1.277c-.349-.365-.752-.785-.863-1.498ZM13 8.999c.7 2.616-3.474 3-5.5 3s-6.2-.484-5.5-3c-.633.316-1 .961-1 1.5c0 1.519 2.91 2.9 6.5 2.9s6.5-1.381 6.5-2.9c0-.539-.367-1.184-1-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IceCream(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.44 8.17a3.473 3.473 0 0 0 2-.63a3.5 3.5 0 0 0 1.56.6h.44L8 13.7a.5.5 0 0 1-.92 0Zm6-3.473a2 2 0 0 1-4 0a2 2 0 1 1-2-2h.12a2 2 0 1 1 3.75 0h.13A2 2 0 0 1 11.439 4.7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IceCreamEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M4 6a2.49 2.49 0 0 0 1.5-.5c.432.325.959.5 1.5.5l-1 4.69a.5.5 0 0 1-.92 0L4 6zm3-4h-.09a1.5 1.5 0 1 0-2.82 0H4a1.5 1.5 0 1 0 1.5 1.5A1.5 1.5 0 1 0 7 2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IceCreamFifteen(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.44 8.17a3.48 3.48 0 0 0 2-.63c.464.323 1 .53 1.56.6h.44L8 13.7a.5.5 0 0 1-.92 0L5.44 8.17z" fill="currentColor"/><path d="M11.44 4.67a2 2 0 1 1-4 0a2 2 0 1 1-2-2h.12a2 2 0 1 1 3.75 0h.13a2 2 0 0 1 2 2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Industry(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 1v12H1V8.72a.5.5 0 0 1 .17-.37l3-3.22a.5.5 0 0 1 .83.38v3l3.16-3.37a.5.5 0 0 1 .84.37V11h3V1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IndustryEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M10 1v8H1V6l2.11-1.78c.32-.32.89-.31.89.14V6l2.13-1.86A.5.5 0 0 1 7 4.5V8h2V1h1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Information(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 1C6.7 1 6 1.7 6 2.5S6.7 4 7.5 4S9 3.3 9 2.5S8.3 1 7.5 1M4 5v1s2 0 2 2v2c0 2-2 2-2 2v1h7v-1s-2 0-2-2V6c0-.5-.5-1-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InformationEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.599.94c-.6 0-1.1.5-1.1 1.1s.5 1.1 1.1 1.1s1.1-.5 1.1-1.1s-.5-1.1-1.1-1.1zM3 4l-.001.74S4.5 4.634 4.5 6v1.5c0 1.5-1.501 1.74-1.501 1.74L3 10h5.2l-.001-.76s-1.2 0-1.2-1.5L7 5s0-1-1-1H3z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JewelryStore(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 8.5a4.5 4.5 0 1 1-9 0a4.498 4.498 0 0 1 2.71-4.125l.176.137l.774.601A3.498 3.498 0 0 0 4 8.5C4 10.43 5.57 12 7.5 12S11 10.43 11 8.5a3.498 3.498 0 0 0-2.66-3.387l.95-.738A4.498 4.498 0 0 1 12 8.5m-4.5-4L10 2.555L9 1H6L5 2.555l1.5 1.167z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JewelryStoreEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M7.574 3.694l-.851.639A2.492 2.492 0 0 1 8 6.5C8 7.878 6.878 9 5.5 9S3 7.878 3 6.5c0-.932.519-1.737 1.277-2.167l-.851-.639A3.485 3.485 0 0 0 2 6.5a3.5 3.5 0 1 0 7 0a3.485 3.485 0 0 0-1.426-2.806zM7.5 2.5L6.5 1h-2l-1 1.5l2 1.5l2-1.5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Karaoke(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.1 2.952a2.988 2.988 0 0 0-5.11 1.965l3.142 3.142A2.988 2.988 0 0 0 12.1 2.952m-7.428 5.3L2.55 10.377a1 1 0 0 0 0 1.414l.707.707a1 1 0 0 0 1.414 0l2.121-2.121l2.122-2.122l-2.121-2.121Zm.741 2.087l-.707-.707l2.087-2.084l.707.707Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KaraokeEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M4.5 4.4l-2 2l-1 1c-.5.4-.6 1.1-.3 1.7l.6.7c.6.4 1.3.2 1.7-.3l1-1l2-2l-2-2.1zM3.1 8.5l-.6-.7l2-2l.7.7l-2.1 2zM5 2l1-1h3l1 1v3L9 6H8L5 3V2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KaraokeFifteen(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M12.1 2.952a2.988 2.988 0 0 0-5.11 1.965l3.142 3.142A2.988 2.988 0 0 0 12.1 2.952z" fill="currentColor"/><path d="M4.672 8.255L2.55 10.377a1 1 0 0 0 0 1.414l.707.707a1 1 0 0 0 1.414 0l2.121-2.121l2.122-2.122l-2.121-2.121zm.741 2.087l-.707-.707l2.087-2.087l.707.707z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Landmark(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 12H12v-.5c0-.3-.2-.5-.5-.5H11V6h1l1-2c-1 .1-2 .1-3 0c-.8-.6-1.4-1.2-2-2v-.5c0-.3-.2-.5-.5-.5s-.5.2-.5.5V2c-.6.8-1.2 1.4-2 2c-1 .1-2 .1-3 0l1 2h1v5h-.5c-.3 0-.5.2-.5.5v.5h-.5c-.3 0-.5.2-.5.5v.5h11v-.5c0-.3-.2-.5-.5-.5M7 11H5V6h2zm3 0H8V6h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LandmarkEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.5 9H8V5h1l1-2c-.7.1-1.3.1-2 0c-.7-.3-1.4-.6-2-1v-.5c0-.3-.2-.5-.5-.5s-.5.2-.5.5V2c-.6.4-1.3.7-2 1c-.7.1-1.3.1-2 0l1 2h1v4H1.5c-.3 0-.5.2-.5.5s.2.5.5.5h8c.3 0 .5-.2.5-.5S9.8 9 9.5 9zM7 9H4V5h3v4z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LandmarkJp(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.5 4a2 2 0 1 1-4 0a2 2 0 0 1 4 0m-6 4a2 2 0 1 0 0 4a2 2 0 0 0 0-4m8 0a2 2 0 1 0 0 4a2 2 0 0 0 0-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LandmarkJpEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M6.75 3.5A1.25 1.25 0 1 1 5.5 2.25A1.25 1.25 0 0 1 6.75 3.5zm-4 2.75A1.25 1.25 0 1 0 4 7.5a1.25 1.25 0 0 0-1.25-1.25zm5.5 0A1.25 1.25 0 1 0 9.5 7.5a1.25 1.25 0 0 0-1.25-1.25z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Landuse(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.911 2.854a.248.248 0 0 1 .089.19V6.75a.25.25 0 0 1-.25.25h-2.5A.25.25 0 0 1 6 6.75V5H5v1.75a.25.25 0 0 1-.25.25h-2.5A.25.25 0 0 1 2 6.75v-3.7a.248.248 0 0 1 .089-.19L5.343.132a.245.245 0 0 1 .315 0zM7.752 8.5a.248.248 0 0 0-.138.042L5 10.5H4v-2a.5.5 0 1 0-1 0v2H2v-2a.5.5 0 1 0-1 0v4.25a.25.25 0 0 0 .25.25h6.5a.25.25 0 0 0 .25-.25v-4a.248.248 0 0 0-.248-.25M14 5.245v6.5a.253.253 0 0 1-.253.253h-3.494a.253.253 0 0 1-.253-.251V5.25a.248.248 0 0 1 .154-.231A.249.249 0 0 1 10.25 5H11v-.751A.249.249 0 0 1 11.249 4h1.5a.253.253 0 0 1 .251.253V5h.755a.245.245 0 0 1 .245.245M13 10h-2v1h2zm0-2h-2v1h2zm0-2h-2v1h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LanduseEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M1 4.749L.995 2.057a.251.251 0 0 1 .1-.2l2.249-1.8a.251.251 0 0 1 .313-.005l2.249 1.8a.251.251 0 0 1 .094.2v2.7A.251.251 0 0 1 5.751 5h-1.5A.251.251 0 0 1 4 4.749V3H3v1.752A.251.251 0 0 1 2.746 5h-1.5A.247.247 0 0 1 1 4.749zm4.753 2.6a.248.248 0 0 0-.173.072L4 9V6.5a.5.5 0 0 0-1 0V9H2V6.5a.5.5 0 0 0-1 0v4.25a.25.25 0 0 0 .25.25h4.5a.249.249 0 0 0 .25-.248V7.6a.25.25 0 0 0-.247-.253zM11 3.253v6.5a.247.247 0 0 1-.247.247H7.247A.247.247 0 0 1 7 9.753v-6.5A.252.252 0 0 1 7.252 3H8v-.752A.248.248 0 0 1 8.248 2h1.506a.246.246 0 0 1 .246.246V3h.747a.253.253 0 0 1 .253.253zM10 8H8v1h2zm0-2H8v1h2zm0-2H8v1h2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laundry(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1L6 3H3S2 3 2 4v9c0 1 1 1 1 1h9s1 0 1-1V2c0-1-1-1-1-1zm.5 1h2a.499.499 0 1 1 0 1h-2a.499.499 0 1 1 0-1m-1 4a3 3 0 1 1 0 6a3 3 0 0 1 0-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LaundryEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5 0L4 2H2S1 2 1 3v7s0 1 1 1h7c1 0 1-1 1-1V1c0-1-1-1-1-1H5zm1 1h3v1H6V1zm-.5 3a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Library(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.082 9.939C.987 9.867 1 9.748 1 9.748L1.526 3.5s.008-.069.039-.104c.019-.022.043-.054.09-.073C2.182 3.107 5.053 1.552 6.5 3c.24.278.5.688.5 1v5.288s.006.112-.095.18a.252.252 0 0 1-.242.003c-2.226-1.114-4.711.154-5.338.487a.22.22 0 0 1-.243-.02m12.593.019c-.627-.333-3.112-1.6-5.338-.487a.252.252 0 0 1-.242-.003C7.994 9.4 8 9.288 8 9.288V4c0-.312.26-.722.5-1c1.446-1.448 4.3.107 4.827.323c.046.019.07.051.09.073c.03.035.039.104.039.104l.543 6.248s.014.119-.08.19a.22.22 0 0 1-.244.02m-4.81 2.728a.24.24 0 0 0 .118-.077a.214.214 0 0 0 .042-.109c.05-.938 1.624-1.812 4.648-.03c.077.044.166.04.242-.015c.086-.063.085-.17.085-.17v-.553s0-.077-.027-.119a.227.227 0 0 0-.093-.085c-2.025-1.315-4.586-1.898-5.885-.16a.197.197 0 0 1-.073.09c-.057.045-.126.042-.126.042h-.585s-.07.003-.126-.041a.197.197 0 0 1-.073-.09c-1.3-1.739-3.86-1.184-5.885.131a.227.227 0 0 0-.093.086c-.027.042-.027.118-.027.118v.554s-.001.107.085.17a.22.22 0 0 0 .243.015c3.023-1.782 4.598-.88 4.647.057a.214.214 0 0 0 .043.109a.24.24 0 0 0 .118.077c.721.18 1.768.25 2.722 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LibraryEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M0 1v7c3.26 0 5.49 2 5.49 2S7.76 8 11 8V1C7 1 5.49 2.725 5.49 2.725S4 1 0 1zm1 1c1.195-.031 3.053.402 4 1.348L5.5 4l.5-.652c.962-.93 2.8-1.36 4-1.348v5c-2 0-3.354.856-4.51 1.781C4.35 7.853 3 7 1 7V2zm1 1.303v.181c.823.17 2.1.687 3 1.176v-.228c-.917-.479-2.176-.972-3-1.13zm7 0c-.824.157-2.083.65-3 1.129v.228c.9-.489 2.177-1.007 3-1.176v-.181zm-7 .92v.181c.822.154 2.099.659 3 1.133v-.195c-.917-.476-2.176-.967-3-1.12zm7 0c-.824.152-2.083.643-3 1.119v.195c.901-.474 2.178-.979 3-1.133v-.181zm-7 .894V5.3c.822.155 2.098.659 3 1.133v-.196c-.918-.475-2.176-.965-3-1.119zm7 0c-.824.154-2.082.644-3 1.12v.195c.902-.474 2.178-.979 3-1.133v-.182zM2 6v.182c.82.14 2.097.629 3 1.087v-.15C4.082 6.645 2.824 6.154 2 6zm7 0c-.824.154-2.082.644-3 1.12v.15c.903-.459 2.18-.948 3-1.088V6z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LibraryFifteen(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M7.47 4.92S5.7 3 1 3v8c4.7 0 6.47 2 6.47 2S9.3 11 14 11V3C9.3 3 7.47 4.92 7.47 4.92zM13 10a10.84 10.84 0 0 0-5.53 1.68A10.56 10.56 0 0 0 2 10V4c3.4.26 4.73 1.6 4.75 1.61l.73.74l.72-.75S9.6 4.26 13 4v6zm-5 .24l-.1-.17A14.519 14.519 0 0 1 12 8.89v.2a14.27 14.27 0 0 0-4 1.18v-.03zm0-1l-.1-.17A14.51 14.51 0 0 1 12 7.9v.2a14.29 14.29 0 0 0-4 1.19v-.05zm0-1l-.1-.17A14.491 14.491 0 0 1 12 6.9v.2a14.24 14.24 0 0 0-4 1.19v-.05zm0-1l-.1-.17A14.45 14.45 0 0 1 12 5.9v.2a14.28 14.28 0 0 0-4 1.19v-.05zm-1.1 3A13.9 13.9 0 0 0 3 9.08v-.2c1.387.203 2.736.614 4 1.22l-.1.14zm0-1A13.899 13.899 0 0 0 3 8.1v-.2a14.12 14.12 0 0 1 4 1.21l-.1.13zm0-1A13.86 13.86 0 0 0 3 7.1v-.2c1.387.199 2.735.607 4 1.21l-.1.13zm0-1A13.899 13.899 0 0 0 3 6.1v-.2a14.1 14.1 0 0 1 4 1.21l-.1.13z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LiftGate(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 4a2 2 0 1 1 4 0v8a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1zm3 0a1 1 0 1 0-2 0a1 1 0 0 0 2 0m4.5 2L7 3H6v3zM10 6h1.5L10 3H8.5zm3.25 0H13l-1.5-3h1.75a.75.75 0 0 1 .75.75v1.5a.75.75 0 0 1-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lighthouse(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 6L0 7v-.5l4.5-1zm0-2.5L0 2.5V3l4.5 1zm6 0V4L15 3v-.5zm0 2.5L15 7v-.5l-4.5-1zM8 7V2h2.5a.51.51 0 0 0 .2-1l-3-1a.5.5 0 0 0-.41 0l-3 1a.51.51 0 0 0 .2 1H7v5H5l-2 7h9l-2-7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LighthouseEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M7 6l1 5H3l1-5h1.2V2h-.92a.32.32 0 0 1-.12-.6L5.38 1a.29.29 0 0 1 .24 0l1.22.4a.32.32 0 0 1-.12.6h-.91v4H7zm1-3v.5l3-.5v-.5L8 3zm0 2.5l3 .5v-.5L8 5v.5zM3 3l-3-.5V3l3 .5V3zm0 2l-3 .5V6l3-.5V5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LighthouseJp(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.5 7.4c0 1.1-.9 2-2 2s-2-.9-2-2s.9-2 2-2s2 .9 2 2m4.5.1c0 .3-.2.5-.5.5h-1.6c-.1.9-.311 1.582-.911 2.282l1.064 1.064c.2.2.196.504-.004.704a.478.478 0 0 1-.704.004L10.3 11c-.6.5-1.4.8-2.3.9v1.6c0 .3-.2.5-.5.5s-.5-.2-.5-.5v-1.6c-.9-.1-1.7-.4-2.3-.9l-1.1 1.1c-.207.207-.48.213-.696-.004c-.2-.2-.207-.5-.007-.7L4 10.3c-.5-.6-.9-1.4-1-2.3H1.5c-.3 0-.5-.2-.5-.5s.2-.5.5-.5H3c.1-.9.4-1.7 1-2.4l-1-1c-.2-.2-.2-.5 0-.7s.5-.2.7 0l1 1c.6-.5 1.4-.8 2.3-.9V1.5c0-.3.2-.5.5-.5s.5.2.5.5V3c.9.1 1.7.5 2.3 1l1.1-1.1c.2-.2.5-.2.7 0s.2.5 0 .7L11 4.7c.5.7.8 1.4.9 2.3h1.6c.3 0 .5.2.5.5m-2.9-.1c0-2-1.6-3.6-3.6-3.6S3.8 5.5 3.8 7.4S5.4 11 7.4 11s3.7-1.6 3.7-3.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lodging(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.5 2.5c-.3 0-.5.2-.5.5v9.5c0 .3.2.5.5.5s.5-.2.5-.5V11h13v1.5c0 .3.2.5.5.5s.5-.2.5-.5v-2c0-.3-.2-.5-.5-.5H1V3c0-.3-.2-.5-.5-.5m3 .5C2.7 3 2 3.7 2 4.5S2.7 6 3.5 6S5 5.3 5 4.5S4.3 3 3.5 3M7 4C5.5 4 5.5 5.5 5.5 5.5V7h-3c-.3 0-.5.2-.5.5v1c0 .3.2.5.5.5H15V6.5C15 4 12.5 4 12.5 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LodgingEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M1.5 2a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 1 0V8h7v.5a.5.5 0 0 0 1 0v-1a.5.5 0 0 0-.5-.5H2V2.5a.5.5 0 0 0-.5-.5zm2 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2zM6 3a1 1 0 0 0-1 1v1H3a.5.5 0 0 0 0 1h7V5a2 2 0 0 0-2-2H6z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Logging(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.91 3.41L13.5 3l.2-.28c.02-.03.04-.05.06-.08a.969.969 0 0 0 .24-.65a1 1 0 0 0-1.62-.77c-.03.02-.06.05-.09.07L12 1.5l-.386-.386a.354.354 0 0 0-.525.475l.417.417l-.5.5l-.392-.392a.354.354 0 0 0-.525.475L10.5 3l-.5.5l-.381-.381a.354.354 0 1 0-.528.472l.028.028L9.5 4l-.5.5l-.395-.395a.354.354 0 0 0-.517.483L8.5 5l-.5.5l-.365-.365a.354.354 0 0 0-.556.439C7.1 5.6 7.508 6 7.508 6L7 6.5l-.394-.392a.392.392 0 0 0-.5-.028a.384.384 0 0 0-.028.5l.022.02l.4.4l-.5.5l-.4-.4a.37.37 0 0 0-.5 0l1.148 1.144L12.09 2.41L12.5 2h.5v.5l-.41.41l-5.839 5.839L7.9 9.9a.355.355 0 0 0 0-.5l-.008-.009L7.5 9l.5-.5l.384.384a.354.354 0 0 0 .528-.472L8.5 8l.5-.5l.38.38a.354.354 0 0 0 .528-.472l-.02-.02L9.5 7l.5-.5l.377.377A.35.35 0 0 0 10.64 7a.347.347 0 0 0 .252-.6L10.5 6l.5-.5l.38.38a.354.354 0 0 0 .528-.472l-.028-.028L11.5 5l.5-.5l.381.381a.354.354 0 0 0 .528-.472l-.025-.025L12.5 4l.5-.5l.382.382a.354.354 0 0 0 .528-.472M4.39 7.916C3.893 7.419 3.641 7 3 7H1.5a.5.5 0 0 0-.5.5v3a2.19 2.19 0 0 0 .5 1.5l.815.811A2.251 2.251 0 0 0 5.493 13L7.5 11ZM3.5 10l-1.238 1.238A1.3 1.3 0 0 1 2 10.5V8h1a.545.545 0 0 1 .335.194a.455.455 0 0 1 .165.418Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LoggingEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.25 1a.409.409 0 0 0-.11.01a.668.668 0 0 0-.44.24L8 2h-.737A.25.25 0 0 0 7 2.236a.243.243 0 0 0 0 .025V3h-.737a.256.256 0 0 0-.19.073a.248.248 0 0 0-.073.179V4h-.763a.275.275 0 0 0-.164.073a.248.248 0 0 0-.073.179V5h-.751a.257.257 0 0 0-.176.075l.178.178l.5.5L8.5 2H9v.5L5.25 6.25l.5.5l.177.177A.249.249 0 0 0 6 6.749V6h.751a.25.25 0 0 0 .176-.073A.266.266 0 0 0 7 5.759V5h.748A.248.248 0 0 0 8 4.756V4h.751a.267.267 0 0 0 .176-.073A.258.258 0 0 0 9 3.733V3l.78-.72a.734.734 0 0 0 .21-.42l.01-.11A.755.755 0 0 0 9.25 1zM2 8l-.431-1.2a1.422 1.422 0 0 1-.07-.465V6.25a.25.25 0 0 1 .25-.25h1a.25.25 0 0 0 0-.5H1.5A.5.5 0 0 0 1 6v.415a2 2 0 0 0 .106.642L1.57 8.43a1.841 1.841 0 0 0 .1.24a1.551 1.551 0 0 0 .624.643a1.342 1.342 0 0 0 1.315 0C3.869 9.122 5 8 5 8L3.5 6.5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Marker(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 1C5.423 1 3 2.288 3 5.568C3 7.793 6.462 12.712 7.5 14c.923-1.288 4.5-6.09 4.5-8.432C12 2.288 9.577 1 7.5 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarkerEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.5-.018c-1.787 0-3.871 1.092-3.871 3.872C1.629 5.739 4.607 9.908 5.5 11c.794-1.092 3.871-5.161 3.871-7.147c0-2.779-2.084-3.87-3.871-3.87z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarkerStroked(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.117 2.818C5.77 2.266 6.652 2 7.5 2c.848 0 1.73.266 2.383.818C10.508 3.346 11 4.2 11 5.568c0 .372-.154.956-.49 1.723c-.324.743-.775 1.56-1.267 2.367a47.732 47.732 0 0 1-1.761 2.652a39.975 39.975 0 0 1-1.763-2.687c-.485-.814-.926-1.636-1.243-2.375C4.147 6.482 4 5.912 4 5.568c0-1.368.492-2.222 1.117-2.75m2.99 10.337C9.46 11.263 12 7.542 12 5.568C12 2.288 9.577 1 7.5 1S3 2.288 3 5.568c0 1.881 2.475 5.69 3.868 7.6c.254.348.472.633.632.832l.197-.273z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarkerStrokedEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.486 11l-.365-.446c-.7-.858-3.544-4.739-3.544-6.638A3.726 3.726 0 0 1 5.132.026q.167-.008.333 0a3.726 3.726 0 0 1 3.888 3.557q.007.166 0 .331c0 1.6-1.806 4.268-3.38 6.415zM5.465.916a2.817 2.817 0 0 0-3 3c0 1.268 1.883 4.161 2.987 5.62c.935-1.282 3.011-4.217 3.011-5.62a2.817 2.817 0 0 0-3-3z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobilePhone(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 2v-.5a.5.5 0 0 0-1 0V2H5a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1M6 13H5v-1h1Zm0-2H5v-1h1Zm0-2H5V8h1Zm2 4H7v-1h1Zm0-2H7v-1h1Zm0-2H7V8h1Zm2 4H9v-1h1Zm0-2H9v-1h1Zm0-2H9V8h1Zm0-2.5a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5v-3a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobilePhoneEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M7 2v-.5a.5.5 0 0 0-1 0V2H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1zM5 9H4V8h1zm0-2H4V6h1zm2 2H6V8h1zm0-2H6V6h1zm0-2H4V3h3z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Monument(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 0L6 2.5v7h3v-7zM3 11.5V15h9v-3.5L10.5 10h-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MonumentEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.5 0L4 2v4.5h3V2L5.5 0zM3 7L2 8v3h7V8L8 7H3z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MonumentJp(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 13H4c-.5 0-1-.5-1-1V5c0-2 .8-4 4-4c3.1 0 4 2 4 4v6h1c.5 0 1 .5 1 1s-.5 1-1 1m-7-2h4V5c0-1 0-2-2-2S5 4 5 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mountain(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 1c-.3 0-.4.2-.6.4l-5.8 9.5c-.1.1-.1.3-.1.4c0 .5.4.7.7.7h11.6c.4 0 .7-.2.7-.7c0-.2 0-.2-.1-.4L8.2 1.4C8 1.2 7.8 1 7.5 1m0 1.5L10.8 8H10L8.5 6.5L7.5 8l-1-1.5L5 8h-.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MountainEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.518 1.232a.556.556 0 0 0-.495.268l-4 6.66c-.222.37.045.84.477.84h8c.432 0 .7-.47.477-.84l-4-6.66a.556.556 0 0 0-.46-.268zm.002.922L8.431 7H7.76L6.416 5.773L5.519 7l-.894-1.227L3.281 7H2.61l2.91-4.846z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Museum(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 0L1 3.5V4h13v-.5zM2 5v5l-1 1.6V13h13v-1.4L13 10V5zm2 1h1v5.5H4zm3 0h1v5.5H7zm3 0h1v5.5h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MuseumEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.5 0L1 2v1h9V2L5.5 0zM2 4v4L1 9v1h9V9L9 8V4H2zm1.49 1a.5.5 0 0 1 .36.15L5.5 6.79l1.65-1.64A.5.5 0 0 1 8 5.5v3a.5.5 0 0 1-1 0V6.71L5.85 7.85a.5.5 0 0 1-.707.003L5.14 7.85L4 6.71V8.5a.5.5 0 0 1-1 0v-3a.5.5 0 0 1 .49-.5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MuseumFifteen(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M7.5 0L1 3.445V4h13v-.555L7.5 0zM2 5v5l-1 1.555V13h13v-1.445L13 10V5H2zm2.615 1a.625.625 0 0 1 .451.184L7.5 8.617l2.434-2.433A.625.625 0 0 1 11 6.625v4.242a.625.625 0 1 1-1.25 0V8.133L7.941 9.94a.625.625 0 0 1-.882 0L5.25 8.133v2.734a.625.625 0 1 1-1.25 0V6.625c0-.341.274-.62.615-.625z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Music(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 0a.49.49 0 0 0-.23.06L3.5 2.5A.5.5 0 0 0 3 3v6.28A2 2 0 0 0 2 9a2 2 0 1 0 2 2V6.36l8-2.22v3.64a2 2 0 0 0-1-.28a2 2 0 1 0 2 2v-9a.5.5 0 0 0-.5-.5M12 3.14L4 5.36v-2l8-2.22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.5.5a.489.489 0 0 0-.22.05L3.5 2a.5.5 0 0 0-.5.5v4.59A1.5 1.5 0 1 0 4 8.5V5.38l5-1.25v1.46A1.5 1.5 0 1 0 10 7V1a.5.5 0 0 0-.5-.5zM4 4.38v-1.5l5-1.25v1.5L4 4.38z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Natural(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.753 6.278a.5.5 0 0 1-.51 0A4.109 4.109 0 0 1 6.5 3.5a2.779 2.779 0 0 0-.59-1.506l-.019-.027a.257.257 0 0 1-.056-.144a.237.237 0 0 1 .25-.25a.264.264 0 0 1 .057.011A3.523 3.523 0 0 1 7.5 2.5L8.28.94a.246.246 0 0 1 .44 0L9.5 2.5a3.4 3.4 0 0 1 1.339-.907a.283.283 0 0 1 .1-.021c.175.009.212.119.221.249a.35.35 0 0 1-.043.141A5.2 5.2 0 0 0 10.5 3.5a4.113 4.113 0 0 1-1.747 2.778M5 8h1.289a.25.25 0 0 0 .25-.25a.241.241 0 0 0-.061-.15L4.7 5.235a.255.255 0 0 0-.391 0L2.518 7.589a.274.274 0 0 0-.062.161a.25.25 0 0 0 .25.25H4l-2.665 2.6a.273.273 0 0 0-.058.152a.25.25 0 0 0 .252.248H3l-1.565 1.565a.255.255 0 0 0 .18.435H4v1h1v-1h2.385a.255.255 0 0 0 .18-.435L6 11h1.471a.25.25 0 0 0 .25-.25a.233.233 0 0 0-.058-.149Zm9.345 3.748a.252.252 0 0 1-.252.252H8.908a.252.252 0 0 1-.226-.365l2.588-5.184a.252.252 0 0 1 .451 0l2.6 5.184a.251.251 0 0 1 .024.113M12.75 10L11.5 7.5L10.25 10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NaturalEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M4.579 9.579L3 8h1.4a.25.25 0 0 0 .25-.25a.246.246 0 0 0-.079-.179L3 6h.736a.25.25 0 0 0 .25-.25a.246.246 0 0 0-.078-.179l-1.2-1.253a.253.253 0 0 0-.4-.015c-.02.023-1.21 1.266-1.21 1.266a.245.245 0 0 0-.082.18a.25.25 0 0 0 .25.25H2L.425 7.575A.249.249 0 0 0 .6 8H2L.434 9.566a.244.244 0 0 0-.082.18A.25.25 0 0 0 .6 10H2v1h1v-1h1.4a.247.247 0 0 0 .179-.421zm6.242-.938l-2.1-4.2a.248.248 0 0 0-.443 0l-2.1 4.2A.248.248 0 0 0 6.4 9h4.2a.248.248 0 0 0 .221-.359zM7.5 7l1-2l1 2zm.279-5.921a.266.266 0 0 1-.049.148A3.513 3.513 0 0 0 7 3a2.141 2.141 0 0 1-1.291 1.911a.475.475 0 0 1-.419 0A2.141 2.141 0 0 1 4 3a3.5 3.5 0 0 0-.726-1.769a.271.271 0 0 1-.046-.148a.25.25 0 0 1 .25-.25a.27.27 0 0 1 .067.009A2.939 2.939 0 0 1 5 2L5.25.224A.25.25 0 0 1 5.744.2L6 2A2.957 2.957 0 0 1 7.453.841a.272.272 0 0 1 .076-.012a.25.25 0 0 1 .25.25z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObservationTower(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 12.5h-.5L10 9V7c.995 0 .995-.75.995-.75L11.5 3s0-.5-.5-.5h-1s0-.5-.5-.5h-1v-.5s0-1-1-1s-1 1-1 1V2h-1c-.5 0-.5.5-.5.5H4c-.5 0-.5.5-.5.5l.505 3.25S4.005 7 5 7v2l-1.5 3.5H3s-1 0-1 .75S3 14 3 14h9s1 0 1-.75s-1-.75-1-.75m-1.75-9L10 5H5l-.25-1.503zM8.5 7v1h-2V7zm-2 2.497h2l1 3.003h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Optician(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.785 3.812a1.469 1.469 0 0 1 2.352-.382l.717.716a.5.5 0 1 1-.708.708l-.716-.717a.469.469 0 0 0-.75.122L1.308 7h3.939c.817 0 1.5.566 1.693 1.325a2.41 2.41 0 0 1 .56-.075c.2 0 .385.034.559.075A1.751 1.751 0 0 1 9.752 7h3.939l-1.37-2.74a.469.469 0 0 0-.751-.123l-.716.717a.5.5 0 1 1-.707-.708l.716-.716a1.47 1.47 0 0 1 2.352.382l1.726 3.45v.001c.037.07.059.151.059.237v1a.5.5 0 0 1-.5.5H14v.5c0 1.589-1.002 2.5-2.75 2.5h-.5C9.002 12 8 11.089 8 9.5v-.159a1.709 1.709 0 0 0-.5-.091c-.17 0-.35.043-.5.091V9.5C7 11.089 5.998 12 4.25 12h-.5C2.002 12 1 11.089 1 9.5V9H.5a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 .06-.237l-.001-.001zM9.752 8A.753.753 0 0 0 9 8.752V9.5c0 .698.2 1.5 1.75 1.5h.5C12.8 11 13 10.198 13 9.5v-.747A.754.754 0 0 0 12.247 8zm-7 0A.754.754 0 0 0 2 8.752V9.5c0 .698.199 1.5 1.75 1.5h.5C5.801 11 6 10.198 6 9.5v-.748A.753.753 0 0 0 5.247 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OpticianEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M10.5 4c-.025 0-.046.01-.07.014C10.28 3.35 9.816 3 9 3H7.5c-.885 0-1.355.414-1.457 1.197c-.156-.047-.344-.089-.543-.089s-.387.042-.543.089C4.855 3.414 4.385 3 3.5 3H2C1.185 3 .72 3.35.57 4.014C.545 4.011.524 4 .5 4a.5.5 0 1 0 0 1c0 1.5 1 2 2.25 2S5 6.5 5 5v-.027c.115-.047.317-.115.5-.115s.385.068.5.115V5c0 1.5 1 2 2.25 2s2.25-.5 2.25-2a.5.5 0 0 0 0-1zM4.25 5c0 .65-.182 1.25-1.5 1.25s-1.5-.6-1.5-1.25v-.5c0-.589.161-.75.75-.75h1.5c.589 0 .75.161.75.75V5zm5.5 0c0 .65-.182 1.25-1.5 1.25s-1.5-.6-1.5-1.25v-.5c0-.589.161-.75.75-.75H9c.589 0 .75.161.75.75V5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paint(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 3.5h-1v-.997A.503.503 0 0 0 11.997 2H1.503A.503.503 0 0 0 1 2.503v2.994c0 .278.225.503.503.503h10.494a.503.503 0 0 0 .503-.503V4.5h.5v3.102l-6.112 1.41A.5.5 0 0 0 6.5 9.5v1.507a.533.533 0 0 0-.5.53v2.926c0 .297.24.537.537.537h.926c.297 0 .537-.24.537-.537v-2.926a.533.533 0 0 0-.5-.53v-1.11l6.112-1.41A.5.5 0 0 0 14 8V4a.5.5 0 0 0-.5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaintEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M10 2H9v-.637C9 1.163 8.837 1 8.637 1H1.363c-.2 0-.363.163-.363.363v2.274c0 .2.163.363.363.363h7.274c.2 0 .363-.163.363-.363V3h.5v2.018l-4.604.974a.5.5 0 0 0-.396.489v1.096h-.041A.459.459 0 0 0 4 8.036v2.005c0 .253.205.459.459.459h1.082A.459.459 0 0 0 6 10.041V8.036a.459.459 0 0 0-.459-.46H5.5v-.69l4.604-.974a.5.5 0 0 0 .396-.49V2.5A.5.5 0 0 0 10 2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Park(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 5.75a1.75 1.75 0 0 0-1-1.61A1.24 1.24 0 0 0 11.75 3a1.197 1.197 0 0 0-.29.06a1.2 1.2 0 0 0-2-.78a1.25 1.25 0 1 0-2.5 0s.04.02.04.05a1.22 1.22 0 0 0-2 .74A1.491 1.491 0 0 0 4.51 3a1.5 1.5 0 0 0-1.4 2.07a1.49 1.49 0 0 0 0 2.87A1.49 1.49 0 0 0 6 7.71c.183.16.41.26.65.29v5L5 14h5l-1.6-1v-2a6.8 6.8 0 0 1 2.77-2A1.49 1.49 0 0 0 13 7.52a1.472 1.472 0 0 0 0-.16a1.75 1.75 0 0 0 1-1.61m-5.6 4.51V6.82c.27.48.778.779 1.33.78h.28c.016.44.224.849.57 1.12a7.25 7.25 0 0 0-2.18 1.54"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParkAltOne(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.428 10.429a.269.269 0 0 1-.074-.18A.25.25 0 0 1 9.6 10h1.447a.25.25 0 0 0 .25-.25a.258.258 0 0 0-.079-.179L9.07 7.419a.3.3 0 0 1-.063-.17a.249.249 0 0 1 .25-.249H10.4a.251.251 0 0 0 .25-.251a.247.247 0 0 0-.077-.178L8.432 4.434l-.014-.014a.262.262 0 0 1-.066-.17A.25.25 0 0 1 8.6 4h.866a.25.25 0 0 0 .25-.25a.246.246 0 0 0-.068-.164h.006L7.7 1.238a.253.253 0 0 0-.042-.044a.249.249 0 0 0-.158-.055a.249.249 0 0 0-.158.055a.253.253 0 0 0-.042.044L5.352 3.586a.246.246 0 0 0-.068.164a.25.25 0 0 0 .25.25H6.4a.241.241 0 0 1 .184.42l-.014.014l-2.139 2.137a.247.247 0 0 0-.077.178A.251.251 0 0 0 4.6 7h1.145a.249.249 0 0 1 .25.249a.3.3 0 0 1-.063.17l-2.15 2.152a.258.258 0 0 0-.082.179a.25.25 0 0 0 .25.25H5.4a.25.25 0 0 1 .248.249a.269.269 0 0 1-.074.18l-2.14 2.132l-.009.009a.248.248 0 0 0 0 .351a.256.256 0 0 0 .18.079H7v1l1-.008V13h3.391a.263.263 0 0 0 .26-.254a.248.248 0 0 0-.071-.177Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParkAltOneEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.571 8.563L7.425 6.424a.255.255 0 0 1-.069-.174A.25.25 0 0 1 7.6 6h.8a.25.25 0 0 0 .25-.25a.246.246 0 0 0-.068-.165l-.008-.008L6.4 3.4a.27.27 0 0 1-.051-.149A.25.25 0 0 1 6.6 3h.88a.257.257 0 0 0 .25-.258a.234.234 0 0 0-.05-.142L5.694.224a.223.223 0 0 0-.351-.038C5.327.2 3.352 2.578 3.352 2.578a.246.246 0 0 0-.068.164a.257.257 0 0 0 .25.258H4.4a.25.25 0 0 1 .25.25a.262.262 0 0 1-.066.17L2.431 5.571a.247.247 0 0 0-.077.178A.251.251 0 0 0 2.6 6h.789a.249.249 0 0 1 .25.249a.3.3 0 0 1-.063.17L1.43 8.563A.253.253 0 0 0 1.6 9H5v1h1V9h3.4a.255.255 0 0 0 .249-.255a.248.248 0 0 0-.078-.182z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParkEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M11 3.75A1.75 1.75 0 0 0 9.25 2a1.73 1.73 0 0 0-.8.2A1.24 1.24 0 0 0 7.21 1a1.19 1.19 0 0 0-.21.05A1.23 1.23 0 0 0 5.75 0a1.25 1.25 0 0 0-1.13.73A1.21 1.21 0 0 0 4 .52a1.23 1.23 0 0 0-1 .55A1.491 1.491 0 0 0 2.51 1a1.51 1.51 0 0 0-1.4 2.08A1.49 1.49 0 0 0 1.08 6a1.49 1.49 0 0 0 2.55.52h.12c.321 0 .63-.126.86-.35V10L3 11h5l-1.6-1V9a8.42 8.42 0 0 1 2.38-2c.26-.05.5-.167.7-.34A1.49 1.49 0 0 0 10 5.5a1.46 1.46 0 0 0 0-.17c.61-.29 1-.904 1-1.58zm-4.64 4.5V6h.18a1.52 1.52 0 0 0 .53-.1a1.5 1.5 0 0 0 .89 1a8.821 8.821 0 0 0-1.6 1.35z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Parking(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2v11h2V9h2.5a3.5 3.5 0 1 0 0-7zm2 5V4h2.5a1.5 1.5 0 1 1 0 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParkingEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M8.1 6.1a3.53 3.53 0 0 1-2.29.66H3.9V10H2V1h3.93a3.2 3.2 0 0 1 2.16.69a2.69 2.69 0 0 1 .81 2.16a2.76 2.76 0 0 1-.8 2.25zM6.64 2.86a1.56 1.56 0 0 0-1-.3H3.9v2.65h1.72a1.48 1.48 0 0 0 1-.32a1.31 1.31 0 0 0 .36-1a1.23 1.23 0 0 0-.34-1.03z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParkingGarage(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 10.14a3.53 3.53 0 0 1-2.29.66h-1.9V14h-1.9V5h3.92a3.2 3.2 0 0 1 2.16.69a2.69 2.69 0 0 1 .81 2.15a2.76 2.76 0 0 1-.8 2.3M9 6.9a1.56 1.56 0 0 0-1-.3H6.31v2.65H8a1.48 1.48 0 0 0 1-.32a1.31 1.31 0 0 0 .36-1A1.23 1.23 0 0 0 9 6.9m5.41-2.69a.5.5 0 0 0-.24-.66L7.5.45L.79 3.55a.501.501 0 1 0 .42.91L7.5 1.55l6.29 2.9a.5.5 0 0 0 .66-.24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParkingGarageEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M7.25 7.44a2.35 2.35 0 0 1-1.53.44H4.45V10H3.19V4H5.8a2.13 2.13 0 0 1 1.44.46c.385.372.583.897.54 1.43a1.84 1.84 0 0 1-.53 1.55zm-1-2.16a1 1 0 0 0-.68-.2H4.45v1.76H5.6a1 1 0 0 0 .68-.22a.87.87 0 0 0 .24-.68a.82.82 0 0 0-.24-.66h-.03zm4.16-2a.5.5 0 0 0-.19-.68L5.72.1a.5.5 0 0 0-.49 0L.73 2.6a.5.5 0 0 0 .49.87l4.28-2.4l4.26 2.37a.5.5 0 0 0 .679-.198l.001-.002l-.03.04z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParkingPaid(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 13V2h4.5a3.5 3.5 0 0 1 3.21 2.103a3.36 3.36 0 0 0-.504.676c-.23.41-.232.837-.174 1.154l.3 1.625A3.495 3.495 0 0 1 5.5 9H3v4zm4.5-6a1.5 1.5 0 1 0 0-3H3v3zm3.516-1.248a.716.716 0 0 1 .062-.484C9.315 4.848 9.978 4 11.5 4c1.521 0 2.185.847 2.421 1.268a.716.716 0 0 1 .064.484L13.019 11H12v2h-1v-2H9.981zM12.75 7L13 5.655C12.83 5.409 12.42 5 11.5 5s-1.33.41-1.5.655L10.25 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pharmacy(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.5 4l1.07-1.54c.06.005.12.005.18 0A1.25 1.25 0 1 0 9.5 1.25v.1L7 4zM12 6V5H3v1l1.5 3.5L3 13v1h9v-1l-1-3.5zm-2 4H8v2H7v-2H5V9h2V7h1v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PharmacyEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M6 4l2-2a.732.732 0 0 1 0-.21a.75.75 0 1 1 .75.75h-.11L7.5 4H6zm3.48 1.83L8.65 7.5l.83 1.67V10h-8v-.83l.84-1.67l-.84-1.67V5h8v.83zM7.5 7H6V5.5H5V7H3.5v1H5v1.5h1V8h1.5V7z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PicnicSite(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 3c-.554 0-1 .446-1 1s.446 1 1 1h1.297l-.645 2H2.5c-.554 0-1 .446-1 1s.446 1 1 1h1.51l-.969 3.01a.75.75 0 1 0 1.426.465l.002-.004L5.586 9h3.828l1.117 3.47a.75.75 0 1 0 1.428-.46L10.99 9h1.51c.554 0 1-.446 1-1s-.446-1-1-1h-2.152l-.645-2H11c.554 0 1-.446 1-1s-.446-1-1-1zm2.873 2h1.254l.645 2H6.229z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PicnicSiteEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M2.75 2c-.416 0-.75.334-.75.75s.334.75.75.75H4L3.54 5H1.75c-.415 0-.75.335-.75.75s.335.75.75.75h1.326L2 10h1l1.076-3.5h2.848L8 10h1L7.924 6.5H9.25c.415 0 .75-.335.75-.75S9.665 5 9.25 5H7.46L7 3.5h1.25c.415 0 .75-.334.75-.75S8.665 2 8.25 2h-5.5zM5 3.5h1L6.46 5H4.54L5 3.5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pitch(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2m7.5 7H10L9 7L8 5.25L9 5l2.3 1a.539.539 0 0 0 .4-1L9 4H7L5 5L4 6H2.5a.5.5 0 0 0 0 1H5l1-1l1 2l-2 2v3.5a.5.5 0 0 0 1 0v-3.11L8 9l1 2h3.5a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PitchEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M4 2a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm6.5 6H9L8 5L7 3.25L8 3l2.3 1a.531.531 0 0 0 .36-1L8 2H6L4 3L3 4H1.47a.5.5 0 0 0 0 1H4l1-1l1 2l-2 1v3.5a.5.5 0 0 0 1 0V7.39L7 7l1 2h2.5a.5.5 0 0 0 0-1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaceOfWorship(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.5 0l-2 2v2h4V2zm-2 4.5L4 6h7L9.5 4.5zM2 6.5a1 1 0 0 0-1 1V13h2V7.5a1 1 0 0 0-1-1m2 0V13h7V6.5zm9 0a1 1 0 0 0-1 1V13h2V7.5a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaceOfWorshipEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.52 0L4 2v1h3V2L5.52 0zM4 4L2 5v5h7V5L7 4H4zm7 1.5V10h-1V5.5a.5.5 0 0 1 1 0zm-10 0V10H0V5.5a.5.5 0 0 1 1 0z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Playground(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 1.15a1.5 1.5 0 1 1 2.903.757A1.5 1.5 0 0 1 3 1.15m11 11.17a1 1 0 0 1-.796 1.17H13.2a1 1 0 0 1-1.07-.49l-1.68-3.37L9 9.92l-.22.08h-.06v2.15l.62-.15h.14a.52.52 0 0 1 .19 1l-5 1a.51.51 0 0 1-.17 0a.52.52 0 0 1-.2-1l4.15-.83V10l-3.22.58a1 1 0 0 1-1.21-.68H4L3 5.83a1 1 0 0 1 0-.43a1 1 0 0 1 .8-.75l4.7-.52V0h.22v4.1h.06L9 4.08L9.4 4h.21a.5.5 0 0 1 .37.6a.49.49 0 0 1-.49.4L9 5.08h-.28v2.86h.06L9 7.88l1.81-.36a1 1 0 0 1 1 .6l2 3.94a.999.999 0 0 1 .19.26M8.5 5.13L6 5.4l.74 2.94L8.5 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaygroundEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M1 .79a1 1 0 1 1-.005.02L1 .79zm9.85 7.37l-2-2A.49.49 0 0 0 8.5 6h-.13l-1.87.55V4a.49.49 0 0 0 .5-.41a.5.5 0 0 0-.38-.59H6.5V0h-.22v3l-4.46.55A1 1 0 0 0 1 4.3a1 1 0 0 0 0 .43l.81 3.13a1 1 0 0 0 .71.68c.151.036.309.036.46 0H3l3.29-.89v1.62l-3 .74a.524.524 0 0 0 .31 1a.946.946 0 0 0 .17 0l4-1a.522.522 0 0 0-.3-1L7.27 9l-.77.22V7.58l1.83-.51l1.81 1.78a.5.5 0 1 0 .72-.69h-.01zM6.28 6.61l-2.07.46l-.71-2.74L6.28 4v2.61z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Police(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 1L6 2h5l.5-1zM6 2.5v1.25S6 6.5 8.5 6.5S11 3.75 11 3.75V2.5zM1.984 3.986A1 1 0 0 0 1 5v4a1 1 0 0 0 1.217.977L5 9.357V14l5.879-6.93a1.58 1.58 0 0 0-.438-.07H6.5L3 7.754V5a1 1 0 0 0-1.016-1.014m9.764 3.725L6.412 14H12V8.559c0-.314-.094-.604-.252-.848"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PoliceEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M4.5.5l.5 1h3l.5-1h-4zM5 2v1a1.5 1.5 0 0 0 3 0V2H5zM1.75 3C1 3 1 3.75 1 3.75v3a.75.75 0 0 0 .975.715L4 6.826V10l4-5H5s-.195 0-.424.072L2.5 5.727V3.75S2.5 3 1.75 3zm7.021 2.389L5 10h4V6a.98.98 0 0 0-.229-.611z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PoliceJp(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 1A6.5 6.5 0 1 0 14 7.5A6.5 6.5 0 0 0 7.5 1m2.6 2.84l-2.6 2.6l-2.6-2.6a4.443 4.443 0 0 1 5.2 0M3.84 4.9l2.6 2.6l-2.6 2.6a4.443 4.443 0 0 1 0-5.2m1.06 6.26l2.6-2.6l2.6 2.6a4.443 4.443 0 0 1-5.2 0m6.26-1.06l-2.6-2.6l2.6-2.6a4.443 4.443 0 0 1 0 5.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PoliceJpEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.5.5a5 5 0 1 0 5 5a5 5 0 0 0-5-5zm2.057 2.182L5.5 4.739L3.443 2.682a3.442 3.442 0 0 1 4.114 0zm-4.875.761L4.739 5.5L2.682 7.557a3.442 3.442 0 0 1 0-4.114zm.761 4.875L5.5 6.261l2.057 2.057a3.442 3.442 0 0 1-4.114 0zm4.875-.761L6.261 5.5l2.057-2.057a3.442 3.442 0 0 1 0 4.114z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Post(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 3.651a.651.651 0 0 1-.29.542L7.5 8L1.79 4.193A.651.651 0 0 1 2.151 3H12.85c.36 0 .651.292.651.651m0 2.316V11a1 1 0 0 1-1 1h-10a1 1 0 0 1-1-1V5.967c0-.2.223-.319.389-.208L7.5 9.5l5.611-3.74a.25.25 0 0 1 .389.207"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PostEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M10 5.5V9a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1V5.5a.5.5 0 0 1 .5-.5a.49.49 0 0 1 .21 0L5.5 7l3.8-2a.488.488 0 0 1 .2 0a.5.5 0 0 1 .5.5zM1.25 2.92l.08.08L5.5 5l4.19-2h.06a.49.49 0 0 0 .25-.5a.5.5 0 0 0-.5-.5h-8a.5.5 0 0 0-.5.5a.49.49 0 0 0 .25.42z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PostJp(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 4a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2h-10a1 1 0 0 1-1-1m11 2.5h-10a1 1 0 0 0 0 2h4V13a1 1 0 1 0 2 0V8.5h4a1 1 0 1 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PostJpEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M1 2.25a.75.75 0 0 1 .75-.75h7.5a.75.75 0 0 1 0 1.5h-7.5A.75.75 0 0 1 1 2.25zm8.25 1.749h-7.5a.75.75 0 0 0 0 1.5H4.5v3a1 1 0 0 0 2 0v-3h2.75a.75.75 0 1 0 0-1.5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Prison(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 1v13H12V1zm6 1H11v3.5H9.5zm-5 .055H6V7H4.5zm2.5 0h1.5V7H7zM10.25 6.5a.75.75 0 0 1 0 1.5a.75.75 0 0 1 0-1.5M7 8h1.473l.027 5H7.027zm-2.5.166H6V13H4.5zM9.5 9H11v4H9.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PrisonEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M2 1v9h7V1H2zm1 1h1v3H3V2zm2 0h1v3H5V2zm2 0h1v2H7V2zm.5 3a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1zM3 6h1v3H3V6zm2 0h1v3H5V6zm2 1h1v2H7V7z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Racetrack(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.993 1.582a.5.5 0 1 0-.986-.164l-2 12a.5.5 0 1 0 .986.164l.67-4.02c.806.118 1.677.157 2.363.638a3.3 3.3 0 0 0 1.432.583c.966.146 1.83-.385 2.784-.234l1.289.194c.26.04.53-.16.569-.42l.884-5.934l.004-.004a.518.518 0 0 0-.427-.564l-1.289-.194c-.963-.143-1.829.373-2.783.23A2.8 2.8 0 0 1 7.3 3.38c-.739-.517-1.619-.603-2.486-.725zm-.59 3.538l.33-1.972c.599.082 1.233.129 1.788.369l-.295 1.965c-.57-.233-1.213-.278-1.822-.362m-.658 3.95l.33-1.974c.62.086 1.277.13 1.858.368l.3-1.976c.658.27 1.159.733 1.893.841l.3-1.98c.738.111 1.349-.177 2.058-.234l-.3 1.966c-.71.06-1.324.36-2.06.25l-.286 1.978c-.736-.11-1.238-.575-1.899-.844l-.3 1.976c-.595-.239-1.263-.281-1.894-.371m4.094-.76c.734.11 1.351-.192 2.061-.251l.284-1.978c.655-.06 1.325.111 1.968.209l-.28 1.976c-.644-.097-1.316-.269-1.971-.207l-.3 1.976c-.709.048-1.335.36-2.062.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RacetrackBoat(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.064 11.216c.47-.513.936-1.022 1.594-1.242a2.004 2.004 0 0 1 1.867.28c.724.532 1.39-.22.591-.806a3.023 3.023 0 0 0-2.774-.423c-.72.24-1.213.767-1.705 1.294c-.4.427-.8.855-1.322 1.129c-.62.324-1.065-.027-1.065-.636V10.8c0-.798-.408-1.3-1.25-1.3c-.49 0-.992.258-1.5.5a16.15 16.15 0 0 0-3.317 2.113a.505.505 0 0 0-.067.707c.25.3.58.162.83-.038c.413-.333 1.78-1.391 2.933-1.94c.477-.227.862-.342 1.121-.342c.25 0 .25.102.25.34c-.002 1.326 1.29 2.142 2.529 1.494c.504-.265.896-.692 1.285-1.118M2.5 9c1.5-1 3.255-2 5-2c.75 0 .75.375.75.75s0 .75.75.75c2.527 0 5.363-3.214 5.5-5.5L12 4l-2.342-.974c-.612-.204-.928.744-.316.948L10.5 4.5L9 5L5 3.75c-.236-.079-.389.027-.5.25L3 7L1 8zM5 2a1 1 0 1 1 2 0a1 1 0 0 1-2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RacetrackCycling(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8 1l3 3h1c2 0 2-3 0-3zm.463 2a.5.5 0 0 0-.317.146l-2.5 2.5a.5.5 0 0 0 .042.745l2.158 1.726l-1.276 2.125a.501.501 0 1 0 .86.516l1.5-2.5a.5.5 0 0 0-.117-.649L7.304 6.402L9 4.707l1.146 1.147A.5.5 0 0 0 10.5 6H13a.5.5 0 0 0 0-1h-2.293L8.854 3.146A.5.5 0 0 0 8.463 3M2.949 7a3 3 0 1 0 .1 6a3 3 0 0 0-.099-6zm9 0a3 3 0 1 0 .1 6a3 3 0 0 0-.099-6zm.1 1a2 2 0 1 1-.098 3.998A2 2 0 0 1 12.05 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RacetrackHorse(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 1a1 1 0 1 0 0 2a1 1 0 0 0 0-2M5.5 3c-.25 0-1 0-.5.75L6.25 6H4.5c-.277.003-.517.1-.75.25c-.806-1.129-1.508-.698-2.24-.25c-.403.247-.816.499-1.26.5a.258.258 0 0 0-.25.25a.258.258 0 0 0 .25.25c.376.02.807-.093 1.238-.207c.656-.173 1.31-.345 1.766-.043C3.16 6.937 3 7.5 3 8l-.002 1.51c-.206.204-.379.123-.555.04c-.14-.065-.283-.13-.443-.05l-1 2.95c0 .25.393.302.5.05l.75-1.756C2.998 10.744 4.5 10 5 9l4 1l2.5 2.5c.197.197.62-.138.46-.385L10 9l2-2.5c.248.5.495.5.986.5H13l.6.75c.388.485 1.22-.007.9-.5l-1.666-2.46L13 4.5c.062-.187-.363-.576-.5-.5l-2 1.5l-1.61-2.313A.5.5 0 0 0 8.5 3zM9 5.25l.5.75l-2.723 1.916a.5.5 0 1 1-.554-.832z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rail(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 1a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1zm2.75.5h3.51a.25.25 0 1 1 0 .5H5.75a.25.25 0 1 1 0-.5M3.5 3H7v4H3.5a.5.5 0 0 1-.5-.5v-3a.5.5 0 0 1 .5-.5M8 3h3.5a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5H8zM5 8a1 1 0 1 1 0 2a1 1 0 0 1 0-2m5 0a1 1 0 1 1 0 2a1 1 0 0 1 0-2m.445 3.994a.504.504 0 0 0-.425.676l.17.33H4.81l.13-.27a.5.5 0 0 0-.91-.41l-1 2a.487.487 0 0 0-.03.18a.5.5 0 0 0 .5.5a.49.49 0 0 0 .43-.26v-.05H4l.31-.69h6.38l.31.69v.05a.49.49 0 0 0 .43.26a.5.5 0 0 0 .5-.5a.49.49 0 0 0 0-.24l-1-2a.5.5 0 0 0-.485-.266"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RailEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9 10.5a.5.5 0 0 1-.5.5a.49.49 0 0 1-.43-.27L7.69 10H3.31l-.37.74A.5.5 0 0 1 2 10.5a.49.49 0 0 1 .07-.24l1-2A.49.49 0 0 1 3.5 8a.5.5 0 0 1 .5.7l-.19.3h3.38L7 8.7a.5.5 0 0 1 .9-.43l1 2a.49.49 0 0 1 .1.23zM8 0H3a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1V1a1 1 0 0 0-1-1zM3.5 6a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zm0-2a.5.5 0 0 1-.5-.5v-2a.5.5 0 0 1 .5-.5h1.79v3H3.5zm4 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zM8 3.5a.5.5 0 0 1-.5.5H5.69V1H7.5a.5.5 0 0 1 .5.5v2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RailLight(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 0C5 0 5 .5 5 .5v1a.499.499 0 1 0 1 0V1h1v2H6S4 3 4 5v3c0 3 3 3 3 3h1s3 0 3-3V5c0-2-2-2-2-2H8V1h1v.5a.499.499 0 1 0 1 0v-1c0-.5-.5-.5-.5-.5zm2 4l2.045.773L10 6.5c.132.5-.5.5-.5.5h-4s-.632 0-.5-.5l.455-1.727zm0 4a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1m-3.375 4L3 15h1.5l.375-1h5.25l.375 1H12l-1.125-3h-1.5l.375 1h-4.5l.375-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RailLightEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M4 0c-.5 0-.5.5-.5.5s0 .5.5.5h1v.97L4 2S3 2 3 3v3c0 2 2 2 2 2h1s2 0 2-2V3c0-1-1-1-1-1H6V1h1c.5 0 .5-.5.5-.5S7.5 0 7 0H4zm1.5 3l1.5.5V5H4V3.5L5.5 3zm0 3a.499.499 0 1 1 0 1a.499.499 0 1 1 0-1zM2.834 8.5L2 11h1.5l.334-1h3.332l.334 1H9l-.834-2.5H6.668l.166.5H4.166l.166-.5H2.834z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RailMetro(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 0s-.75 0-1 1L3 6.5V10c0 1 1 1 1 1h7s1 0 1-1V6.5L10.5 1c-.273-1-1-1-1-1zm1 1.5h2s.536 0 .75 1L10 6c.215 1.002-1 1-1 1H6s-1.215.002-1-1l.75-3.5c.214-1 .75-1 .75-1M5 8a1 1 0 1 1 0 2a1 1 0 0 1 0-2m1.75 0h1.5a.25.25 0 1 1 0 .5h-1.5a.25.25 0 1 1 0-.5M10 8a1 1 0 1 1 0 2a1 1 0 0 1 0-2m-5.875 4L3 15h1.5l.375-1h5.25l.375 1H12l-1.125-3h-1.5l.375 1h-4.5l.375-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RailMetroEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M3.5 0C3 0 3 .5 3 .5L2 5v1c0 1.024 1 1 1 1h5s1 0 1-1V5L8 .5S8 0 7.5 0h-4zM4 1h3l.5 3h-4L4 1zm-.5 4a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1zm1.75 0h.5a.25.25 0 1 1 0 .5h-.5a.25.25 0 1 1 0-.5zM7.5 5a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1zM3 8l-1 3h1.5l.334-1h3.332l.334 1H9L8 8H6.5l.334 1H4.166L4.5 8H3z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RangerStation(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9 .5l-2 1v3.773L2 8v6h4v-4h3v4h4V8L8 5.273V4l1-.5l2 1l2-1v-3l-2 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RangerStationEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M6.334 0L5 1v3L1 6.055V10h3V7h3v3h3V6.055L6 4V2.25L6.334 2l1.332 1L9 2V0L7.666 1L6.334 0z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Recycling(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.456 8.613c-.338.598-.955 1.69.137 2.418c.343.227.728.384 1.131.462c.307.045.323.518-.038.507c-.385-.02-2.26-.193-2.561-1.6c-.156-.82.02-1.557.504-2.355l.697-1.233l-1.306-.743L4.5 4v4l-1.306-.694zM6.7 2.034c1.155-.628 1.823.43 2.191 1.007l.806 1.263l-1.266.808L12 6.986l-.197-4.026l-1.264.807l-.76-1.189c-.522-.746-.904-1.297-1.835-1.545C6.307.72 5.301 2.619 5.311 2.607c-.164.287.216.54.451.21c.258-.32.577-.586.938-.783m6.594 6.187c-.088-.19-.549-.141-.419.267c.131.39.185.8.157 1.21C12.939 11.01 11.684 11 11 11H9.5V9.5l-3.5 2l3.488 2.025L9.493 12H11c.89.015 1.6-.176 2.2-.713c1.2-1.061.094-3.066.094-3.066"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RecyclingEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.5 8.5c-.547.519-1.66.435-2.5.5v1L4.5 8.5L7 7v1h1c.634 0 .83 0 1.054-.127a.81.81 0 0 0 .413-.958c-.085-.358.4-.425.485-.115A1.724 1.724 0 0 1 9.5 8.5zM1.948 5.092l-.001.002c-.4.743-1.065 1.637-.921 2.377a1.724 1.724 0 0 0 1.192 1.293c.372.12.466-.395.161-.472a.81.81 0 0 1-.586-.862c.012-.258.122-.42.461-.956l.537-.843V5.63l.844.537l.078-2.914l-2.608 1.302l.843.537zm2.936-3.178a.81.81 0 0 1 1.025-.189c.231.115.336.281.689.807c.14.208.32.479.555.832l-.831.555l2.636 1.246l-.141-2.912l-.833.556c-.52-.662-1.067-1.63-1.796-1.795a1.724 1.724 0 0 0-1.671.562c-.264.338.199.546.367.338z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReligiousBuddhist(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.498 6.999h-.533a5.463 5.463 0 0 0-1.248-3.01l.378-.379a.5.5 0 0 0-.695-.719l-.012.012l-.378.378a5.464 5.464 0 0 0-3.011-1.248V1.5a.5.5 0 0 0-1 0v.533a5.464 5.464 0 0 0-3.01 1.249L3.61 2.9a.5.5 0 0 0-.707.707l.378.377A5.463 5.463 0 0 0 2.034 7H1.5a.5.5 0 0 0 0 1h.533a5.463 5.463 0 0 0 1.247 3.01l-.378.378a.5.5 0 0 0 .695.72l.012-.013l.378-.378a5.462 5.462 0 0 0 3.011 1.249v.533a.5.5 0 0 0 1 0v-.533a5.463 5.463 0 0 0 3.01-1.249l.379.378a.5.5 0 0 0 .719-.695l-.012-.012l-.378-.378A5.464 5.464 0 0 0 12.965 8h.533a.5.5 0 0 0 0-1m-5.5-3.948A4.45 4.45 0 0 1 10.291 4l-2.15 2.15a1.484 1.484 0 0 0-.143-.058zm-1 0v3.04a1.484 1.484 0 0 0-.142.058L4.706 4a4.45 4.45 0 0 1 2.293-.948zm-3 1.655l2.15 2.15A1.484 1.484 0 0 0 6.092 7H3.05a4.45 4.45 0 0 1 .948-2.293M3.05 8h3.041c.017.048.036.096.058.143l-2.15 2.15A4.448 4.448 0 0 1 3.05 8M7 11.947a4.45 4.45 0 0 1-2.293-.949l2.15-2.15c.047.022.095.041.143.058zm1 0v-3.04c.049-.017.096-.037.143-.058l2.15 2.15A4.448 4.448 0 0 1 8 11.947m3-1.656l-2.15-2.15c.022-.046.041-.094.058-.143h3.041a4.448 4.448 0 0 1-.95 2.293zM8.907 7a1.484 1.484 0 0 0-.058-.143l2.15-2.15a4.45 4.45 0 0 1 .948 2.292z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReligiousBuddhistEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M11.002 5.8v-.5H9.989a4.474 4.474 0 0 0-1.138-2.746l.827-.827l-.354-.354l-.823.827a4.475 4.475 0 0 0-2.75-1.138V0h-.5v1.062a4.472 4.472 0 0 0-2.7 1.1l-.876-.886l-.355.352l.87.872a4.478 4.478 0 0 0-1.178 2.8H0v.5h1.012A4.475 4.475 0 0 0 2.15 8.547l-.827.827l.354.354l.823-.827a4.475 4.475 0 0 0 2.75 1.139v.962h.5v-.963A4.476 4.476 0 0 0 8.5 8.896l.82.831l.355-.352l-.82-.834a4.474 4.474 0 0 0 1.134-2.74h1.013zm-1.44-.5H6.101l2.448-2.444a4.05 4.05 0 0 1 1.013 2.445zM8.195 2.5L5.751 4.947V1.489c.905.054 1.766.41 2.444 1.011zM5.251 1.49v3.4L2.85 2.46a4.046 4.046 0 0 1 2.4-.972zM2.49 2.81l2.46 2.49H1.44a4.054 4.054 0 0 1 1.052-2.49zM1.44 5.8h3.462L2.452 8.246A4.05 4.05 0 0 1 1.44 5.801zm1.367 2.8l2.445-2.446V9.61a4.05 4.05 0 0 1-2.445-1.01zm2.945 1.01v-3.5L8.2 8.593a4.05 4.05 0 0 1-2.45 1.02V9.61zm2.8-1.373L6.148 5.8h3.414a4.049 4.049 0 0 1-1.007 2.436h-.004z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReligiousChristian(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 .955V4H3v3h3v8h3V7h3V4H9V1c0-1-.978-1-.978-1H6.99S6 0 6 .955"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReligiousChristianEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M4.5 0v3H2v2h2.5v6h2V5H9V3H6.5V0h-2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReligiousJewish(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5.5L5.5 4H1l2.5 3.5L1 11h4.5l2 3.5l2-3.5H14l-2.5-3.5L14 4H9.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReligiousJewishEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M11 8H7.1l-1.6 3l-1.6-3H0l1.95-2.5L0 3h3.9l1.6-3l1.6 3H11L9.05 5.5L11 8z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReligiousMuslim(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.096 12.096a6.5 6.5 0 1 1 .48-8.658a5.147 5.147 0 1 0 0 8.124a6.556 6.556 0 0 1-.48.534M10.1 6.5H8L9.7 8L9 10.5L11 9l2 1.5l-.7-2.5L14 6.5h-2.1l-.9-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReligiousMuslimEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M4.9 0C2.2 0 0 2.2 0 4.9s2.2 4.9 4.9 4.9c.9 0 1.7-.2 2.4-.6h-.6c-2.4 0-4.3-1.9-4.3-4.3S4.3.6 6.7.6h.6C6.5.2 5.8 0 4.9 0zm3.6 2l-.816 1.7L6 4l1 1.5L6.5 7l2-1l2 1l-.5-1.5L11 4l-1.5-.3l-1-1.7z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReligiousShinto(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 2.5h-10a.945.945 0 0 0-1 1a.945.945 0 0 0 1 1H3V12a.945.945 0 0 0 1 1a.945.945 0 0 0 1-1V8h5v4a.945.945 0 0 0 1 1a.945.945 0 0 0 1-1V4.5h.5a.945.945 0 0 0 1-1a.945.945 0 0 0-1-1M10 6H8V4.5h2ZM7 6H5V4.5h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReligiousShintoEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.25 2.009h-7.5a.75.75 0 0 0 0 1.5H2v5.25a.75.75 0 0 0 1.5 0v-2.75h4v2.75a.75.75 0 0 0 1.5 0v-5.25h.25a.75.75 0 0 0 0-1.5zM7.5 4.759H6v-1.25h1.5zm-4-1.25H5v1.25H3.5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ResidentialCommunity(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.8 12.5V10c.7-.1 1.2-.7 1.2-1.5S12.3 7 11.5 7S10 7.7 10 8.5c0 .7.5 1.4 1.2 1.5v2.5H7V1H2v11.5H1v.5h13v-.5zM6 10H3V9h3zm0-2H3V7h3zm0-2H3V5h3zm0-2H3V3h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ResidentialCommunityEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M8.8 8c.7-.1 1.2-.8 1.2-1.5C10 5.7 9.3 5 8.5 5S7 5.7 7 6.5c0 .7.5 1.4 1.2 1.5v1.5H6V1H2v8.5H1v.5h9v-.5H8.8V8zM3 2h2v1H3V2zm0 2h2v1H3V4zm0 2h2v1H3V6z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Restaurant(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.5 0l-1 5.5c-.146.805 1.782 1.181 1.75 2L4 14c-.038 1 1 1 1 1s1.038 0 1-1l-.25-6.5c-.031-.818 1.733-1.18 1.75-2L6.5 0H6l.25 4l-.75.5L5.25 0h-.5L4.5 4.5L3.75 4L4 0zM12 0c-.736 0-1.964.655-2.455 1.637C9.135 2.373 9 4.018 9 5v2.5c0 .818 1.09 1 1.5 1L10 14c-.09.996 1 1 1 1s1 0 1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RestaurantBbq(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.017 12.5H14.5L12 15v-2.483a1.619 1.619 0 0 1-1.137.48a1.6 1.6 0 0 1-.993-.345a6.784 6.784 0 0 1 .008-10.31a1.618 1.618 0 0 1 2.122.12V0l2.5 2.5h-2.464a1.593 1.593 0 0 1 .123.129a1.624 1.624 0 0 1-.288 2.28a3.515 3.515 0 0 0-1.121 2.586a3.558 3.558 0 0 0 1.147 2.616a1.635 1.635 0 0 1 .232 2.279c-.034.041-.075.072-.112.11M2.504 0l-1 5.5c-.146.805 1.781 1.181 1.75 2l-.25 6.5a.963.963 0 0 0 1 1a.963.963 0 0 0 1-1l-.25-6.5c-.031-.818 1.733-1.18 1.75-2l-1-5.5h-.5l.25 4l-.75.5l-.25-4.5h-.5l-.25 4.5l-.75-.5l.25-4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RestaurantEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M2.25 0l-.5 3.5c-.07.495 1.24 1.4 1.25 2V11h1V5.5c0-.6 1.32-1.505 1.25-2L4.75 0h-.5l.25 2.75l-.75.5V0h-.5v3.25l-.75-.5L2.75 0h-.5zM9 0C7.5 0 7.006 1.724 7 3v3h1v5h1V0z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RestaurantNoodle(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.457 11.99L1 8V7h13v1l-3.496 3.99ZM3.988 2.5a.5.5 0 0 0-1 0v.567l-1.797.368a.25.25 0 1 0 .094.49l1.703-.277v.566l-1.75.036a.25.25 0 0 0 0 .5l1.75.036v1.212h1Zm9.5 1.5l-7.5.263V2.995l7.594-1.074a.5.5 0 0 0-.188-.982L5.98 2.455a.496.496 0 0 0-.99.045v.228l-.494.1v.352l.493-.08v1.197l-.493.01v.461L13.488 5a.5.5 0 0 0 0-1M10 13H5v.576h5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RestaurantNoodleEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M3.469 8.81L1 6.625V6h9v.625l-2.484 2.2zm-.476-7.096a.332.332 0 0 0-.664 0v.88l-1.098.24a.25.25 0 0 0 .092.491l1.006-.177v.432l-1.079.022a.25.25 0 1 0 0 .5l1.079.022v1.2h.664zM9.5 3.352l-.491.193l-4.177.046v-.926l4.055-.788l.519.114a.5.5 0 1 0-.186-.982l-.445.271l-3.943.847v-.413a.332.332 0 0 0-.664 0v.557l-.626.137v.479l.626-.11v.82l-.626.013v.486l5.456.04l.502.216a.5.5 0 1 0 0-1zM7.531 9.743H3.453v.245h4.078z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RestaurantPizza(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.393 7.931L.025 7.315A14.05 14.05 0 0 1 7.764-.001l.027.07l.511 1.331l-.003.001a12.613 12.613 0 0 0-6.906 6.53m10.963 4.96L2.42 8.414a11.604 11.604 0 0 1 6.354-6.009l.003-.001l.817 2.122c-.028-.002-.052-.016-.081-.016a1.062 1.062 0 1 0 .764 1.793l2.413 6.273a.249.249 0 0 1-.334.316ZM7.426 7.52a.925.925 0 1 0-1.849 0a.925.925 0 0 0 1.849 0m3.017 2.03a.925.925 0 1 0-.07.353a.924.924 0 0 0 .07-.354Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RestaurantPizzaEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M1.937 6.507L1.022 6.1A10.038 10.038 0 0 1 6.1 1.022l.406.915a9.033 9.033 0 0 0-4.57 4.57zm7.683 3.44L2.897 6.96a8.03 8.03 0 0 1 4.062-4.062l.391.88a.748.748 0 0 0-.628.732a.758.758 0 0 0 .757.758a.742.742 0 0 0 .458-.17l2.01 4.522a.248.248 0 0 1-.327.328zM6.252 6.496a.75.75 0 1 0-.75.75a.75.75 0 0 0 .75-.75zM8.236 7.5a.74.74 0 1 0-.74.74a.74.74 0 0 0 .74-.74z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RestaurantSeafood(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.433 10.089h2.49v.75h-2.07a4.81 4.81 0 0 1 1.622 2.301l-.71.242a4.01 4.01 0 0 0-1.74-2.197a1.778 1.778 0 0 1-1.035.454l-1.553 1.346H6.52l-1.535-1.362a2.022 2.022 0 0 1-1.093-.44a4.01 4.01 0 0 0-1.743 2.199l-.71-.242a4.815 4.815 0 0 1 1.622-2.301H.99v-.75h2.494l.863-.677l-1.014-1.655a2.274 2.274 0 0 1-.27.029a2.02 2.02 0 0 1-2.052-2.08c0-2.08 1.979-4.16 1.979-4.16l-.737 4.447l1.726-.001L4.99 2.985v2.996a2.094 2.094 0 0 1-.699 1.412l.854 1.394l.084-.066a3.867 3.867 0 0 1 4.478 0l.06.047l.87-1.415a2.118 2.118 0 0 1-.643-1.372V2.985l.983 3.007l1.755.001l-.737-4.447s1.982 2.08 1.982 4.16a2.021 2.021 0 0 1-2.056 2.08a2.313 2.313 0 0 1-.344-.037l-1.01 1.646l.84.659c.013.01.014.024.025.035"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RestaurantSeafoodEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M8.478 7.186h1.494v.6H8.647A3.386 3.386 0 0 1 9.697 9.33l-.568.193a2.758 2.758 0 0 0-1.26-1.548L6.372 8.991L4.53 8.974l-1.492-.963a2.748 2.748 0 0 0-1.206 1.513l-.568-.193a3.386 3.386 0 0 1 1.051-1.546H.99v-.6h1.498a.539.539 0 0 1 .073-.08l.712-.64l-.7-1.143c-.014 0-.024.005-.04.005c0 0-1.533 0-1.533-1.444A4.964 4.964 0 0 1 2.533.994L2 3v1h1l.992-2.006v2a1.423 1.423 0 0 1-.703 1.159l.514.838l.132-.118a2.552 2.552 0 0 1 3.11 0l.116.104l.514-.838a1.426 1.426 0 0 1-.668-1.145v-2L8 4h1V3L8.466.994a4.968 4.968 0 0 1 1.528 2.889c0 1.444-1.528 1.444-1.528 1.444c-.03 0-.05-.008-.078-.01L7.69 6.454l.673.605a.551.551 0 0 1 .114.128z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RestaurantSushi(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 3a.5.5 0 0 0-.5.5H2.75A2.744 2.744 0 0 0 .943 8.313A1.998 1.998 0 0 0 0 10c0 1.1.9 2 2 2h8c1.1 0 2-.9 2-2c0-.488-.186-.93-.48-1.277l2.177.963a.75.75 0 1 0 .606-1.372l-1.81-.802l1.63-.272a.75.75 0 1 0-.246-1.48l-2.13.355C11.258 4.952 10.038 3.784 8 3.55V3.5a.5.5 0 0 0-.5-.5zM6 4h1v7H6V9zm3.5 4c.259 0 .464.197.49.45l.076-.134a.5.5 0 1 1 .866.5l-.258.448A.986.986 0 0 1 11 10c0 .563-.437 1-1 1H8V9h1v-.5c0-.277.223-.5.5-.5m-7.06.07a.498.498 0 0 1 .494.246l.2.348A.493.493 0 0 1 3.5 8.5c.277 0 .5.223.5.5h1v2H2c-.563 0-1-.437-1-1s.437-1 1-1h.174l-.106-.184a.497.497 0 0 1 .371-.746"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoadAccident(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.6 8.7l-1.1-2.2a1.05 1.05 0 0 0-.9-.5H4.4a1.05 1.05 0 0 0-.9.5L2.4 8.7L1.16 9.852a.5.5 0 0 0-.16.366V14.5a.5.5 0 0 0 .5.5h2c.2 0 .5-.2.5-.4V14h7v.5c0 .2.2.5.4.5h2.1a.5.5 0 0 0 .5-.5v-4.282a.5.5 0 0 0-.16-.366ZM4.5 7h6l1 2h-8Zm.5 4.6c0 .2-.3.4-.5.4H2.4c-.2 0-.4-.3-.4-.5v-1.1c.1-.3.3-.5.6-.4l2 .4c.2 0 .4.3.4.5Zm8-.1c0 .2-.2.5-.4.5h-2.1c-.2 0-.5-.2-.5-.4v-.7c0-.2.2-.5.4-.5l2-.4c.3-.1.5.1.6.4ZM8 4H7L6 0h3ZM1.1 1.8L3.7.3l1.1 4l-.8.4ZM11.3.3l2.6 1.5l-2.9 3l-.9-.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Roadblock(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 14a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13M12 6v3H3V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoadblockEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.5 0a5.5 5.5 0 1 0 0 11a5.5 5.5 0 0 0 0-11zM2 4.5h7v2H2v-2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rocket(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.555 2C9.41 2 6.534 3.471 4.602 6H3C1.843 6 1.18 6.864.723 7.777L.11 9H3l1.5 1.5L6 12v2.889l1.223-.612C8.136 13.821 9 13.157 9 12v-1.602c2.529-1.932 4-4.809 4-6.953V2zM9 5a1 1 0 1 1 0 2a1 1 0 0 1 0-2m-6.5 6l-.5.5c-.722.722-1 2.5-1 2.5s1.698-.198 2.5-1l.5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RocketEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9 1C7.488 1 5.408 2.146 4.049 4H3c-.801 0-1.184.367-1.5 1L1 6h2l1 1l1 1v2l1-.5c.633-.316 1-.699 1-1.5V6.951C8.854 5.592 10 3.512 10 2V1H9zM7.5 3a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1zM2.75 7.25l-.25.25C2 8 2 9 2 9s.945.055 1.5-.5l.25-.25l-1-1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func School(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.542 3.647L3.106 3l.443-1.63a.505.505 0 0 1 .618-.352l1.46.392a.5.5 0 0 1 .355.613zm-4.52 7.356a.496.496 0 0 1-.005-.276l1.819-6.726l2.435.647l-1.819 6.726a.499.499 0 0 1-.143.237l-1.457 1.347a.152.152 0 0 1-.247-.066zM10 5c-2.25 0-3-.75-3-3c2.25 0 3 .75 3 3m-1.4 7.984c-1.37.21-3.126-1.706-3.52-3.8L5.969 5.9c.399-.35.903-.533 1.419-.533a2.71 2.71 0 0 1 1.564.489a.964.964 0 0 0 1.089-.01a2.438 2.438 0 0 1 1.46-.479c.77 0 1.643.489 2.05 1.201c1.536 2.696-1.194 6.709-3.144 6.417a.867.867 0 0 1-.255-.093a1.427 1.427 0 0 0-1.302 0a.866.866 0 0 1-.25.092"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SchoolEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M8.5 9V8H10V6H7.5V5H10V3H8.5V2H10V1H6v9h4V9H8.5zM4 7H1V1h3v6zm0 1l-1.5 2L1 8h3z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SchoolJp(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 3h-4V2a1 1 0 0 0-2 0v1h-4a1 1 0 0 0 0 2h.67a10.817 10.817 0 0 0 1.98 3.81c.333.432.66.825.979 1.194a14.383 14.383 0 0 1-2.608 2.117a1 1 0 1 0 .955 1.758A15.65 15.65 0 0 0 7.5 11.461a15.65 15.65 0 0 0 3.024 2.418a1 1 0 1 0 .955-1.758a14.379 14.379 0 0 1-2.607-2.117c.319-.369.645-.762.978-1.193a10.818 10.818 0 0 0 1.98-3.812h.67a1 1 0 0 0 0-2M8.267 7.587c-.261.338-.515.641-.767.937a25.017 25.017 0 0 1-.767-.937a9.589 9.589 0 0 1-1.465-2.589h4.464a9.579 9.579 0 0 1-1.465 2.589"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SchoolJpEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.25 2.25h-3V1.5a.75.75 0 1 0-1.5 0v.75h-3a.75.75 0 0 0 0 1.5h.403a9.139 9.139 0 0 0 1.66 2.967c.218.274.43.52.636.754a6.09 6.09 0 0 1-1.9 1.306a.75.75 0 0 0 .391 1.448a6.727 6.727 0 0 0 2.558-1.672a6.722 6.722 0 0 0 2.563 1.672a.736.736 0 0 0 .188.025a.746.746 0 0 0 .717-.551a.756.756 0 0 0-.515-.922a6.024 6.024 0 0 1-1.905-1.306c.207-.234.417-.48.636-.754A9.193 9.193 0 0 0 8.845 3.75h.405a.75.75 0 0 0 0-1.5zM6.008 5.783c-.177.222-.343.41-.51.601c-.168-.192-.334-.379-.511-.601A8.514 8.514 0 0 1 3.75 3.75h3.497a8.574 8.574 0 0 1-1.24 2.033z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scooter(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.908 12a1.5 1.5 0 1 1-2.816 0Zm8.65-6H13V3h.351a.282.282 0 0 0 .223-.148l.268-.536a.334.334 0 0 0 .009-.066A.25.25 0 0 0 13.6 2H13v-.3a.215.215 0 0 0-.2-.2H9.25a.25.25 0 0 0 0 .5H12v4.6L7.6 10H6V7.5a.5.5 0 0 0-.5-.5H5V5h1.75a.25.25 0 0 0 0-.5l-4.484-.466c-.006 0-.01-.007-.016-.007a.25.25 0 0 0-.25.25v.473a.25.25 0 0 0 .25.25H3v2h-.5A1.538 1.538 0 0 0 1 8.5v2a.472.472 0 0 0 .442.5H7.5l2.5-1h3.5a.472.472 0 0 0 .5-.442V6.5a.472.472 0 0 0-.442-.5M12.5 11a1.5 1.5 0 1 0 1.5 1.5a1.538 1.538 0 0 0-1.5-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScooterEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M1 9h2a.979.979 0 0 1-1 1a.979.979 0 0 1-1-1zm8.753-5H9V3h.351a.282.282 0 0 0 .223-.148l.268-.536a.333.333 0 0 0 .009-.066A.25.25 0 0 0 9.6 2H9v-.5H6.25a.25.25 0 0 0 0 .5H8v2.5L5 7H4V5.5a.5.5 0 0 0-.5-.5H3V4h1.75A.25.25 0 0 0 5 3.75a.245.245 0 0 0-.223-.239V3.5L1.25 3a.25.25 0 0 0-.25.25v.5a.25.25 0 0 0 .25.25H2v1h-.5a.5.5 0 0 0-.5.5V8h5.172a1 1 0 0 0 .709-.294l.419-.414A1 1 0 0 1 8 7h1.752A.248.248 0 0 0 10 6.752v-2.5A.247.247 0 0 0 9.753 4zM9 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shelter(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 2L1 6v2l1-.333V13h10v-2H4V7l9-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShelterEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M4 5v3h6v2H2V5.67L1 6V4l9-3v2L4 5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shoe(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.5 7a9.97 9.97 0 0 1-1.315-.948L6.01 3.221a.558.558 0 0 0-1 .279H5V5H3.209a.5.5 0 0 1-.357-.148S2.5 4 2 4h-.5a.5.5 0 0 0-.5.5V9h5.5c1.5 0 2 1 3.5 1h4v-.5C14 8 10.547 7.594 9.5 7m0 4a3.131 3.131 0 0 1-1.526-.447A4.1 4.1 0 0 0 6 10H1v1.5a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V11a3.134 3.134 0 0 1 1.526.447A4.1 4.1 0 0 0 9.5 12h4a.5.5 0 0 0 .5-.5V11Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoeEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M7.5 5L5.805 3.983L4.8 2.676a.446.446 0 0 0-.627-.083a.442.442 0 0 0-.164.392h-.008V4H2v-.5a.5.5 0 1 0-1 0V6h3.5C5.5 6 6 7 7 7h3l.004-.5C10.004 5.946 7.5 5 7.5 5z" fill="currentColor"/><path d="M5.527 7.584C5.117 7.31 4.651 7 4 7H1v1.47c0 .293.237.53.53.53h1.94A.53.53 0 0 0 4 8.47V8c.349 0 .638.192.973.416c.41.274.876.584 1.527.584h3a.5.5 0 0 0 .5-.5V8H6.5c-.349 0-.638-.193-.973-.416z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoeFifteen(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.5 7c-.414-.235-.89-.596-1.315-.948L6.01 3.222a.56.56 0 0 0-1 .278H5V5H3.209a.504.504 0 0 1-.357-.148S2.5 4 2 4h-.5a.5.5 0 0 0-.5.5V9h5.5c1.5 0 2 1 3.5 1h4v-.5C14 8 10.547 7.594 9.5 7z" fill="currentColor"/><path d="M9.5 11c-.632 0-1.047-.207-1.526-.447C7.456 10.293 6.868 10 6 10H1v1.5a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V11c.632 0 1.046.207 1.526.447c.519.26 1.106.553 1.974.553h4a.5.5 0 0 0 .5-.5V11H9.5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shop(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.33 5H11.5l-.39-2.33A2 2 0 0 0 9.7 1.18A3.76 3.76 0 0 0 8.62 1H6.38a3.76 3.76 0 0 0-1.08.18a2 2 0 0 0-1.41 1.49L3.5 5H1.67a.5.5 0 0 0-.48.65l1.88 6.3A1.5 1.5 0 0 0 4.5 13h6a1.5 1.5 0 0 0 1.42-1.05l1.88-6.3a.5.5 0 0 0-.47-.65M4.52 5l.36-2.17a.91.91 0 0 1 .74-.7c.246-.078.502-.121.76-.13h2.24c.261.008.52.052.77.13a.91.91 0 0 1 .74.7L10.48 5h-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShopEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.6 4H8.49L8.2 2.27a1.47 1.47 0 0 0-1.06-1.11A2.799 2.799 0 0 0 6.33 1H4.67a2.798 2.798 0 0 0-.8.14A1.47 1.47 0 0 0 2.8 2.27L2.51 4H1.34a.37.37 0 0 0-.34.49l1.21 4.7c.135.47.56.799 1.05.81h4.48a1.12 1.12 0 0 0 1.06-.81l1.2-4.7A.37.37 0 0 0 9.6 4zM3.27 4l.27-1.61a.68.68 0 0 1 .55-.52c.185-.06.376-.093.57-.1h1.67c.194.007.385.04.57.1a.68.68 0 0 1 .55.52L7.73 4H3.27z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Skateboard(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.25 5.25V5h1.5v.25c0 .69.56 1.25 1.25 1.25h7c.69 0 1.25-.56 1.25-1.25V5h1.5v.25A2.75 2.75 0 0 1 11 8H4a2.75 2.75 0 0 1-2.75-2.75M5 10a1 1 0 1 1-2 0a1 1 0 0 1 2 0m7 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkateboardEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M8 6H3a2.002 2.002 0 0 1-2-2h1a1.001 1.001 0 0 0 1 1h5a1.001 1.001 0 0 0 1-1h1a2.002 2.002 0 0 1-2 2zm0 1a1 1 0 1 0 1 1a1 1 0 0 0-1-1zM3 7a1 1 0 1 0 1 1a1 1 0 0 0-1-1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Skiing(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 4.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0m-1.28 7.39a.25.25 0 0 0-.33-.11a2.69 2.69 0 0 1-2.28.25L4.78 8.86L7.5 7.5v-3l1.5.75v3l1.5.75l1.5-.75l-1.5-.75V3l-3-1.5l-1.5.75v4.5L3.28 8.11L.61 6.78a.25.25 0 1 0-.22.45l10.5 5.25c.316.135.657.2 1 .19a3.84 3.84 0 0 0 1.72-.44a.25.25 0 0 0 .113-.335z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkiingEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M11 3.68a1.09 1.09 0 1 1-2.18-.008A1.09 1.09 0 0 1 11 3.68zM10.17 9a.25.25 0 0 0-.33-.11a1.9 1.9 0 0 1-1.59.18L3.69 6.81l1.9-1V3.68l1.09.55v2.18L7.77 7l1.09-.55l-1.09-.59V2.59L5.59 1.5L4.5 2v3.32l-1.9 1l-2-1a.25.25 0 1 0-.22.45L8 9.54c.24.101.5.149.76.14a2.85 2.85 0 0 0 1.28-.33a.25.25 0 0 0 .13-.35z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Slaughterhouse(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 6.1c0 1.376-1.363.694-1.363.694L12.5 7H11a5.31 5.31 0 0 1-1.5 2a3.272 3.272 0 0 1-.523 1.125c-.077.091-.06.2-.068.33l.091 2.3a.264.264 0 0 1-.266.25a.242.242 0 0 1-.234-.205L8 9.5H4v3.253a.247.247 0 0 1-.247.247a.25.25 0 0 1-.24-.2L3 10v-.5a8.2 8.2 0 0 1-1.426-.1A1.886 1.886 0 0 1 .9 10.826c-.128.083-.148.211-.133.386l.19 1.538a.25.25 0 0 1-.25.25a.238.238 0 0 1-.23-.174l-.427-1.7a.35.35 0 0 1 .055-.3c.437-.68-.049-2.55-.049-2.55A1.354 1.354 0 0 1 0 7.922V5.5a2.027 2.027 0 0 1 2.736-1.914s.1.03.142.049a15.15 15.15 0 0 0 3.814.038l.18-.062a1.842 1.842 0 0 1 1.26 0c.078.02.154.05.226.089c.1.049.196.106.287.171A1.8 1.8 0 0 0 9.5 4h1V3l.5.5c.5-1.5 2.5-1 2.5-1a1.687 1.687 0 0 0-1.5 1l2.5 2a.612.612 0 0 1 .186-.069a.318.318 0 0 1 .314.321z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlaughterhouseEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M11 5.75v.241C11 7 10 6.5 10 6.5H8A3.081 3.081 0 0 1 7 8a2.848 2.848 0 0 1-.416.9a.26.26 0 0 0-.05.242l.458 1.55A.227.227 0 0 1 7 10.75a.25.25 0 0 1-.25.25a.259.259 0 0 1-.244-.173L6 9v-.5H2.75l-.21.42a.25.25 0 0 0-.02.168L3 10.75a.25.25 0 0 1-.25.25a.26.26 0 0 1-.237-.172L2 9v-.5a4.013 4.013 0 0 1-.843-.139a1.383 1.383 0 0 1-.5 1.045a.242.242 0 0 0-.094.282l.414.99a.213.213 0 0 1 .011.072a.25.25 0 0 1-.25.25a.253.253 0 0 1-.228-.148L0 9.5a6.031 6.031 0 0 0 0-2v-2a1.486 1.486 0 0 1 2-1.4s.768.132 1.1.154a2.457 2.457 0 0 0 .9-.167A1.666 1.666 0 0 1 4.5 4a1.648 1.648 0 0 1 .844.257a3.166 3.166 0 0 0 .9.23L7 4.5L8 4v-.5l.5.5a1.474 1.474 0 0 1 1.5-.5a.914.914 0 0 0-1 .5l1 1l.68.529a.425.425 0 0 1 .07-.029a.238.238 0 0 1 .25.25z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Slipway(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2 10l12 1.495V12H2zm12-4l-1 1v.583L5.196 4.332l.063-.125L6.61 2.845h.831a.35.35 0 0 0 0-.7h-.976a.352.352 0 0 0-.248.103L4.723 3.753a.371.371 0 0 0-.066.09l-.109.219L2 3c0 2-.03 3.958 2.86 4.5C6.28 7.765 13 9 13 9l2-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlipwayEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M1.042 8l8 1.38V10h-8zm8-4l-1 1v.625L3.735 4.01l.077-.285l.86-.845H5.5a.35.35 0 0 0 0-.7h-.971a.348.348 0 0 0-.245.1l-1.03 1.011a.357.357 0 0 0-.092.158l-.086.314L1.042 3C1.095 4 1 5.51 3 5.884C3.983 6.07 8.095 7 8.095 7l1.947-2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Snowmobile(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 10a1 1 0 0 1-1 1H9.414a1 1 0 0 1-.707-.293L7.146 9.146A.499.499 0 0 0 6.793 9H3.334a1 1 0 0 1-.601-.2L1.386 7.788A1.013 1.013 0 0 1 1 7a.995.995 0 0 1 .472-.844L5 4l1.038-1.692A.5.5 0 0 1 7 2.5V6l1 1h1l5.412-.992a.5.5 0 0 1 .355.915L13 8l1.631 1.224A.99.99 0 0 1 15 10m-9.5 0H3.225c-.504 0-.999-.127-1.44-.368L.74 9.062a.5.5 0 0 0-.48.877l1.269.693c.441.242.937.368 1.44.368H5.5a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnowmobileEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M4 8.5a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1 0-1h3a.5.5 0 0 1 .5.5zm7-3a.5.5 0 0 0-.5-.5a.929.929 0 0 0-.097.01L6 6h-.5a.5.5 0 0 1-.5-.5v-3a.5.5 0 0 0-.931-.253L3 4L.312 5.038A.499.499 0 0 0 0 5.5a.52.52 0 0 0 .086.28L1 7h2.411a.488.488 0 0 1 .314.116L5.723 8.77A.996.996 0 0 0 6.36 9h3.136a.504.504 0 0 0 .451-.73L9 7l1.78-1.085A.503.503 0 0 0 11 5.5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Soccer(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 1.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0m0 9.5a1 1 0 1 0 0 2a1 1 0 0 0 0-2m1.84-4.91l-1.91-1.91a.48.48 0 0 0-.37-.18H3.5a.5.5 0 0 0 0 1h2.7L3 11.3a.488.488 0 0 0 0 .2a.511.511 0 0 0 1 .21L5 10h2l-1.93 4.24a.49.49 0 0 0-.07.26a.51.51 0 0 0 1 .2l4.7-9.38l1.44 1.48a.5.5 0 0 0 .7-.71"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoccerEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9 1.25a1.25 1.25 0 1 1-2.5-.001A1.25 1.25 0 0 1 9 1.25zm0 7.23a1 1 0 1 0 0 2a1 1 0 0 0 0-2zm1.81-3.39L8.94 3.18A.48.48 0 0 0 8.56 3H1.51a.5.5 0 0 0 0 1H5L2.07 8.3a.488.488 0 0 0 0 .2a.511.511 0 0 0 1 .21H3L4.16 7H6l-1.93 3.24a.49.49 0 0 0-.07.26a.51.51 0 0 0 1 .2l3.67-6.38l1.48 1.48a.5.5 0 1 0 .7-.71h-.04z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Square(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 13H3a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9 10H2a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareStroked(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.7 3.3v8.4H3.3V3.3zM12 2H3a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareStrokedEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.19 1H1.81a.81.81 0 0 0-.81.81v7.38c0 .447.363.81.81.81h7.38a.81.81 0 0 0 .81-.81V1.81A.81.81 0 0 0 9.19 1zM2 2h7v7H2V2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stadium(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 1v4.01c-2.83.094-4.998.957-5 1.99v4.5c0 1.105 2.462 2 5.5 2s5.5-.895 5.5-2V7c-.002-1.033-2.17-1.896-5-1.99v-.947l3-1.313zM3 8.147c.515.267 1.201.484 2 .632v2.967c-1.205-.269-2-.726-2-1.246zm9 .001V10.5c0 .52-.795.977-2 1.246V8.781c.799-.148 1.485-.366 2-.633m-6 .774a14.68 14.68 0 0 0 3 .002v2.984c-.471.056-.971.092-1.5.092s-1.029-.036-1.5-.092z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StadiumEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5 0v3.012c-2.208.097-3.912.806-4 1.67v3.681C1 9.267 3.014 10 5.5 10S10 9.267 10 8.363V4.682c-.088-.864-1.79-1.574-4-1.67v-.44L8.5 1.5L5 0zM1.818 5.752c.319.178.72.332 1.182.453v2.46c-.75-.237-1.179-.568-1.182-.915V5.752zm7.364.004V7.75c-.002.348-.43.68-1.182.916V6.203c.461-.12.862-.271 1.182-.447zM4 6.398a11.066 11.066 0 0 0 3 0v2.493C6.528 8.962 6.017 9 5.5 9S4.472 8.962 4 8.89V6.399z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.5 0l-2 5h-5l4 3.5l-2 6l5-3.5l5 3.5l-2-6l4-3.5h-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.4 0L4 3.5H0l3 3L1.5 11l3.9-2.6L9.5 11L8 6.5l3-3H7L5.4 0z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarStroked(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.5 2.69l1.07 2.68l.25.63h3l-2 1.75l-.5.44l.23.63l1 3.13l-2.48-1.77l-.57-.4l-.57.4l-2.52 1.77l1-3.13l.21-.63l-.5-.44l-2-1.75h3l.25-.63zM7.5 0l-2 5h-5l4 3.5l-2 6l5-3.5l5 3.5l-2-6l4-3.5h-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarStrokedEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.5 2.69l.59 1.47l.25.63h1.81l-1 .9l-.5.44l.18.63l.56 1.69l-1.32-.92l-.57-.41l-.57.4l-1.31.92l.55-1.68l.21-.63l-.5-.43l-1-.9h1.78l.25-.63l.59-1.48M5.5 0L4 3.79H.19l3 2.66L1.71 11L5.5 8.34L9.29 11L7.78 6.45l3-2.66H7L5.5 0z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Suitcase(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 4V2a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v2H12a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1zm1-2v2h4V2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SuitcaseEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M8 3V1.578L7.36 1H3.64L3 1.748V3H1.5l-.5.5v6l.5.5h8l.5-.5v-6L9.5 3H8zM4 2h3v1H4V2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SushiEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M3.5 2.5c-.3 0-.5.1-.5.4h-.9c-1.1 0-2 .8-2 1.9c0 .5.3 1 .7 1.4c-.5.3-.8.7-.8 1.3C0 8.3.7 9 1.6 9h5.8C8.3 9 9 8.3 9 7.5c0-.2 0-.4-.1-.6l1 .5c.4.2.8 0 1-.3c.2-.4-.053-.794-.343-.944L9.8 5.8l.5-.1c.4-.1.7-.4.6-.8c-.1-.4-.4-.6-.8-.6H10l-1.6.3C8 3.9 7.3 3.1 6 3c0 0 0-.5-.5-.5h-2zm.5 1h1v4.4H4V3.5zm3 2.4c.2 0 .3.1.4.3l.1-.1c.1-.2.3-.2.5-.1s.2.3.1.5l-.3.5c.1.1.2.2.2.4c0 .3-.2.5-.6.5H6v-1h.7v-.6c0-.2.1-.4.3-.4zM1.9 6c.1 0 .3 0 .4.2l.1.2c.1-.1.2-.1.3-.1c.2 0 .3.2.3.3v1.3H1.6c-.4 0-.6-.2-.6-.5s.2-.5.6-.5h.3l-.2-.4c-.1-.2 0-.4.1-.5h.1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Swimming(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.111 2c-.112 0-.435.147-.435.147l-3.322 1.68c-.443.175-.618.881-.352 1.234l.97 1.408l-3.97 2.029L5 9.998l2.502-1.5l2.5 1.5l1.002-1.002l-3-4l2.557-1.53c.528-.266.443-.704.443-.97C11 2.286 10.644 2 10.111 2m2.141 3a1.75 1.75 0 1 0-.003 3.501A1.75 1.75 0 0 0 12.252 5M2.5 10L0 11.5V13l2.5-1.5L5 13l2.502-1.5l2.5 1.5L12 11.5l3 1.5v-1.5L12 10l-1.998 1.5l-2.5-1.5L5 11.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwimmingEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M8.004.494l-.7.502l-2.7 1.602c-.3.099-.4.598-.202.898l.6 1.002l-2.5 2l1 1.002l2-1.002L7.504 7.5l1-1.002l-3.002-3.502l3.002-1.5V.494h-.5zm1.002 2.502A1 1 0 1 0 9.003 5a1 1 0 0 0 .003-2.003zM2 7.998L0 9v1l2-1l1.5 1l2-1l2 1L9 9l2 1V9L9 7.998L7.5 9l-2-1.002L3.5 9L2 7.998z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTennis(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.968 10.228a3.812 3.812 0 0 1-1.913.983L3.768 4.934a4.028 4.028 0 0 1 1.005-1.902C7.03.774 9.98.224 12.378 2.622s1.848 5.348-.41 7.605m-6.988 1.61a3.842 3.842 0 0 1 1.168-.559A4.533 4.533 0 0 1 8 11.445L3.546 7a4.413 4.413 0 0 1 .157 1.922a3.664 3.664 0 0 1-.521 1.116C2.11 11.301 1.05 11.765 1.05 12.226a1.838 1.838 0 0 0 1.724 1.724c.46 0 .918-1.013 2.207-2.113"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTennisEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M8.593 7.232a2.605 2.605 0 0 1-1.145.642L3.115 3.541a2.604 2.604 0 0 1 .642-1.144C5.319.845 7.33.439 8.946 2.054c1.605 1.605 1.209 3.627-.353 5.178zM3 5.253a2.6 2.6 0 0 1-.013 1.156a2.732 2.732 0 0 1-.364.77c-.739.867-1.573 1.337-1.573 1.648A2.386 2.386 0 0 0 2.173 9.95c.31 0 .77-.802 1.659-1.562a2.518 2.518 0 0 1 .802-.396A2.674 2.674 0 0 1 5.725 8z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Teahouse(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 5s2-.1 2-1.6c0-.934-.822-1.191-1.524-1.41c-.521-.164-.976-.306-.976-.69C7.5.387 8 0 8 0S7 .532 7 1.5c0 .71.54.962 1.06 1.203c.479.223.94.437.94.997C9 4.7 8 5 8 5m3.837 1.674C11.376 8.621 10.785 11.107 9 12v2H6v-2c-1.785-.893-2.375-3.38-2.837-5.326c-.056-.233-.11-.46-.163-.674h9c-.054.215-.107.44-.163.674M7 4c0 1-1.5 1-1.5 1s.793-.1.793-.7c0-.275-.277-.445-.574-.628C5.373 3.459 5 3.23 5 2.8c0-.834 1-1.3 1-1.3s-.5.273-.5 1c0 .388.353.541.717.698c.386.167.783.339.783.802"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TeahouseEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M6.2 4C7.4 3.7 8 3.7 8 2.5c0-.6-.4-.8-1.6-1.3C5.6.9 5.3.9 5.3.1c-.3.9-.1 1.2.8 1.7c2 1 .1 2.2.1 2.2zM3.9 4c.7-.2 1.1-.2 1.1-.9c0-.4-.3-.5-1-.9c-.5-.2-.7-.7-.7-1.2c-.2.6-.1 1.3.5 1.6c1.2.6.1 1.4.1 1.4zM9 5H2l1 3c.3.4.6.7 1 1v2h3V9c.4-.3.7-.6 1-1l1-3z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Telephone(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.51 8.88a.51.51 0 0 0 0 .72l.72.72l-2.16 2.18l-.37-.37a2.24 2.24 0 0 1-.7-1.44V9.24a2.24 2.24 0 0 1 .7-1.45l5.07-5.07A2.24 2.24 0 0 1 9.22 2h1.45a2.24 2.24 0 0 1 1.45.72l.36.36l-2.17 2.18l-.73-.73a.51.51 0 0 0-.72 0Zm-.38 4.72a1 1 0 0 0 1.414.036l.036-.036l.72-.72a1 1 0 0 0 .036-1.414L6.3 11.43Zm7.25-7.28a1 1 0 0 0 1.414.036l.756-.756a1 1 0 0 0 .036-1.414l-.036-.036Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelephoneEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M3 9.4a.73.73 0 0 0 1 0l.74-.74a.73.73 0 0 0 0-1zm4.56-4.6a.73.73 0 0 0 1 0l.71-.71a.73.73 0 0 0 0-1zM5.88 3.57L3.57 5.88a.37.37 0 0 0 0 .52l.43.44L2.26 8.6a2.27 2.27 0 0 1-.73-1.34v-1a1.345 1.345 0 0 1 .52-1l3.21-3.21a1.345 1.345 0 0 1 1-.52h1a2.27 2.27 0 0 1 1.34.73L6.84 4l-.44-.43a.37.37 0 0 0-.52 0" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tennis(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.128 1.87c-1.541-1.54-4.553-.776-6.28 1.05a6.828 6.828 0 0 0-1.88 4.4a5.728 5.728 0 0 1-.57 2.72a.681.681 0 0 0-.67.17l-2.49 2.5a.694.694 0 0 0 0 .98l.07.07a.705.705 0 0 0 .98 0l2.5-2.49a.683.683 0 0 0 .18-.66a8.57 8.57 0 0 1 3.12-.58a6.549 6.549 0 0 0 3.99-1.87c2.03-2.03 2.501-4.84 1.05-6.29m-2.2-.04a2.307 2.307 0 0 1 1.64.61a2.548 2.548 0 0 1 .5 2.44l-2.95-2.96a3.865 3.865 0 0 1 .81-.09m-4.57 6.81a2.649 2.649 0 0 1-.48-2.55l3.04 3.04a2.852 2.852 0 0 1-2.56-.49m2.88.39l-3.26-3.27a5.162 5.162 0 0 1 .49-1.08l3.85 3.85a5.51 5.51 0 0 1-1.08.5m1.34-.66l-3.94-3.95a5.203 5.203 0 0 1 .74-.9l4.1 4.1a5.687 5.687 0 0 1-.9.75m1.11-.96l-4.1-4.1a5.571 5.571 0 0 1 .84-.65l3.92 3.92a5.719 5.719 0 0 1-.66.83m.82-1.09L8.698 2.5a4.17 4.17 0 0 1 1.09-.48l3.2 3.2a5.566 5.566 0 0 1-.48 1.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TennisEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.38 1.63c-.823-.853-2.768-.822-4.14.61A4.26 4.26 0 0 0 4.05 5a4.15 4.15 0 0 1-.5 2.07a.62.62 0 0 0-.67.12L1.19 8.88a.61.61 0 0 0 0 .85l.09.09a.594.594 0 0 0 .84.01l.01-.01l1.69-1.69a.6.6 0 0 0 .12-.67a5.824 5.824 0 0 1 2.21-.49a3.966 3.966 0 0 0 2.62-1.19c1.32-1.32 1.572-3.152.61-4.15zM6.15 6.47a1.659 1.659 0 0 1-1.17-.44a1.923 1.923 0 0 1-.31-1.98l2.28 2.28a2.543 2.543 0 0 1-.8.14zm1.06-.23L4.77 3.8a3.72 3.72 0 0 1 .7-1.07l2.81 2.82a3.556 3.556 0 0 1-1.07.69zm1.25-.87L5.64 2.54a3.924 3.924 0 0 1 1.04-.71h.01l2.5 2.5a3.578 3.578 0 0 1-.73 1.04zm.83-1.29L6.94 1.72a2.76 2.76 0 0 1 .92-.17a1.616 1.616 0 0 1 1.17.44a1.716 1.716 0 0 1 .41 1.52a3 3 0 0 1-.15.57z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Terminal(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.3 2.1a.5.5 0 0 1 .409-.017l5.926 2.372l3.514-1.457a1.35 1.35 0 1 1 .993 2.513L4.709 9.058a2 2 0 0 1-2.12-.46L.195 6.195a.2.2 0 0 1 .052-.32l.507-.253a.5.5 0 0 1 .48.018L3.5 7l3.048-1.264l-3.804-3.04a.2.2 0 0 1 .036-.336zM.5 12a.5.5 0 0 0 0 1h14a.5.5 0 0 0 0-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Theatre(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 1S1 1 1 2v5.158C1 8.888 1.354 11 4.5 11H5V8L2.5 9s0-2.5 2.5-2.5V5c0-.708.087-1.32.5-1.775c.381-.42 1.005-1.258 2.656-.471L9 3.303V2s0-1-1-1c-.708 0-1.978 1-3 1S2.787 1 2 1m1 2a1 1 0 1 1 0 2a1 1 0 0 1 0-2m4 1S6 4 6 5v5c0 2 1 4 4 4s4-2 4-4V5c0-1-1-1-1-1c-.708 0-1.978 1-3 1S7.787 4 7 4m1 2a1 1 0 1 1 0 2a1 1 0 0 1 0-2m4 0a1 1 0 1 1 0 2a1 1 0 0 1 0-2m-4.5 4h5s0 2.5-2.5 2.5S7.5 10 7.5 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TheatreEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M.606.7C.303.7 0 1.001 0 1.304v3.734C0 6.452.578 8 3 8h.5V6.184l-2.29-.008s.304-1.137 1.817-1.137c.303 0 .404 0 .606.102V3.727c0-.606.202-1.008.605-1.31C4.541 2.112 4.79 2 5.394 2H6v-.605c0-.303-.197-.659-.5-.659c-.606 0-1.16.569-2.473.569c-1.009 0-1.816-.606-2.421-.606zM1.75 2.5a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5zM5.5 3c-.303 0-.555.424-.555.727V7.46c0 1.413.606 3.027 3.028 3.027S11 8.773 11 7.461V3.727c0-.303-.197-.727-.5-.727c-.605 0-1.592.5-2.5.5S6.106 3 5.5 3zm1.25 2a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5zm2.5 0a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5zM6.156 8.066h3.735s-.303 1.211-1.817 1.211s-1.918-1.21-1.918-1.21z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Toilet(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 1.5a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0-3 0M11.5 0a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3M3.29 4a1 1 0 0 0-.868.504L.566 7.752a.5.5 0 1 0 .868.496l1.412-2.472A345.048 345.048 0 0 0 1 11h2v2.5a.5.5 0 0 0 1 0V11h1v2.5a.5.5 0 0 0 1 0V11h2L6.103 5.687l1.463 2.561a.5.5 0 1 0 .868-.496L6.578 4.504A1 1 0 0 0 5.71 4zM9 4.5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5v5a.5.5 0 0 1-1 0v4a.5.5 0 0 1-1 0v-4h-1v4a.5.5 0 0 1-1 0v-4a.5.5 0 0 1-1 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToiletEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M3.33 2.19a1.1 1.1 0 1 1 1.1-1.1a1.1 1.1 0 0 1-1.1 1.1zm6.94-1.1a1.1 1.1 0 1 0 0 .01v-.01zM6.51 4.93L4.7 3.12A.37.37 0 0 0 4.43 3H2.22a.37.37 0 0 0-.25.1H2L.14 4.93a.38.38 0 1 0 .53.53l1.58-1.58L.77 8h1.46v2.51a.38.38 0 0 0 .75.11H3V8h.69v2.63a.38.38 0 0 0 .75-.11V8h1.44L4.41 3.88L6 5.46a.37.37 0 0 0 .28.12c.21 0 .38-.17.38-.38a.37.37 0 0 0-.15-.27zM8.62 7v3.63a.37.37 0 1 0 .73 0V7H11V3.37a.37.37 0 0 0-.37-.37H7.71a.37.37 0 0 0-.37.37V7h1.28z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Toll(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 1a.5.5 0 0 0-.5.5V13h-.5a.5.5 0 0 0 0 1h7a.5.5 0 0 0 0-1H7v-2.901l6.776-4.553a.5.5 0 0 0 .136-.694l-.279-.415a.5.5 0 0 0-.693-.136L7 8.29V2h.5a.5.5 0 0 0 0-1zM3 3.5a.5.5 0 0 1 .5-.5h2a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1-.5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Town(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.651 6.121a.25.25 0 0 0-.314 0L8.092 7.929A.247.247 0 0 0 8 8.122v4.625a.253.253 0 0 0 .253.253h1.494a.253.253 0 0 0 .253-.253V11h1v1.747a.253.253 0 0 0 .253.253h1.494a.253.253 0 0 0 .253-.253V8.12a.25.25 0 0 0-.094-.2zM10 10H9V9h1zm2 0h-1V9h1zM5.71.815a.252.252 0 0 0-.42 0L2.042 4.936a.252.252 0 0 0-.042.14v7.671a.252.252 0 0 0 .251.253h2.5A.252.252 0 0 0 5 12.748V11h1v1.748a.252.252 0 0 0 .252.252H7V7a.5.5 0 0 1 .188-.391L9 5C9 4.95 5.71.815 5.71.815M4 9H3V8h1zm0-3H3V5h1zm2 3H5V8h1zm0-3H5V5h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TownEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M3.695 1.1a.256.256 0 0 0-.4 0l-2.24 2.831A.254.254 0 0 0 1 4.088V9.75a.25.25 0 0 0 .25.25h1.5A.25.25 0 0 0 3 9.75V8h1v1.75a.25.25 0 0 0 .25.25H5V5.5a.615.615 0 0 1 .147-.4L6 4zM3 7H2V6h1zm0-2H2V4h1zm5.194-1.258a.248.248 0 0 0-.387 0l-1.753 2.19A.249.249 0 0 0 6 6.087v3.665a.248.248 0 0 0 .248.248h3.5A.248.248 0 0 0 10 9.756V6.087a.249.249 0 0 0-.054-.155zM7 6h1v1H7zm2 3H8V8h1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TownHall(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 4H9V1L7.5 0L6 1v3H2L1 5v1h13V5zM7.5 1.5c.4 0 .7.3.7.8s-.3.7-.7.7s-.8-.3-.8-.8c0-.4.4-.7.8-.7M13 7H2v4l-1 1.5V14h13v-1.5L13 11zm-8 5.5H4V8h1zm3 0H7V8h1zm3 0h-1V8h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TownHallEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.5 0L1 2v1h9V2L5.5 0zM2 4v4L1 9v1h9V9L9 8V4H2zm1 1h1v3H3V5zm2 0h1v3H5V5zm2 0h1v3H7V5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TownHallFifteen(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M7.5 0L1 3.445V4h13v-.555L7.5 0zM2 5v5l-1 1.555V13h13v-1.445L13 10V5H2zm2 1h1v5.5H4V6zm3 0h1v5.5H7V6zm3 0h1v5.5h-1V6z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Triangle(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.539 1c-.295 0-.489.177-.616.385l-5.846 9.538C1 11 1 11.153 1 11.308c0 .538.385.692.692.692h11.616c.384 0 .692-.154.692-.692c0-.154 0-.231-.077-.385l-5.77-9.538C8.029 1.177 7.789 1 7.54 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.517 1.232a.556.556 0 0 0-.493.268l-4 6.66c-.223.37.044.84.476.84h8c.432 0 .699-.47.476-.84l-4-6.66a.556.556 0 0 0-.459-.268z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleStroked(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.524.5a.77.77 0 0 0-.69.395l-5.5 9.87C1.022 11.307 1.395 12 2 12h11c.605 0 .978-.692.666-1.235l-5.5-9.87A.773.773 0 0 0 7.524.5M7.5 2.9l4.127 7.47H3.373z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleStrokedEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.517 1.232a.556.556 0 0 0-.493.268l-4 6.66c-.223.37.044.84.476.84h8c.432 0 .699-.47.476-.84l-4-6.66a.556.556 0 0 0-.459-.268zM5.5 2.861l3.02 5.03H2.48l3.02-5.03z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tunnel(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2C2.343 2 1 3.327 1 4.964v7.048c0 .546.448.988 1 .988h1V8.554c0-2.456 2.015-4.446 4.5-4.446S12 6.098 12 8.554V13h1c.552 0 1-.442 1-.988V4.964C14 3.327 12.657 2 11 2zm7 8v-.952c0-2.183-1.567-3.952-3.5-3.952S4 6.866 4 9.048V13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Veterinary(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 6c-2.5 0-3 2.28-3 3.47a6.149 6.149 0 0 0-1.7.89a2 2 0 0 0-.4 2.78a2.06 2.06 0 0 0 2.8.4a4 4 0 0 1 2.3-.69a4 4 0 0 1 2.3.69a2 2 0 0 0 2.8-.3a1.929 1.929 0 0 0-.226-2.72a1.386 1.386 0 0 0-.074-.06l-.1-.1a8.992 8.992 0 0 0-1.7-.89C10.5 8.29 10 6 7.5 6M2.08 4.3a1.58 1.58 0 0 0-.76 2a1.52 1.52 0 0 0 1.61 1.4a1.58 1.58 0 0 0 .76-2a1.52 1.52 0 0 0-1.61-1.4m10.85 0a1.58 1.58 0 0 1 .76 2a1.52 1.52 0 0 1-1.61 1.4a1.58 1.58 0 0 1-.76-2a1.52 1.52 0 0 1 1.61-1.4m-7.85-3c-.68.09-1 .94-.76 1.87A1.77 1.77 0 0 0 5.93 4.7c.68-.09 1-.94.76-1.87A1.77 1.77 0 0 0 5.08 1.3m4.85 0c.68.09 1 .94.76 1.87A1.77 1.77 0 0 1 9.07 4.7c-.68-.08-1-.94-.76-1.87A1.769 1.769 0 0 1 9.93 1.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VeterinaryEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.5 4.32a2.2 2.2 0 0 0-2.14 2.46a4.4 4.4 0 0 0-1.21.63a1.42 1.42 0 0 0-.29 2a1.48 1.48 0 0 0 2 .28a2.89 2.89 0 0 1 1.64-.52a2.89 2.89 0 0 1 1.64.49c.6.495 1.49.41 1.985-.191l.015-.019a1.35 1.35 0 0 0-.21-2l-.08-.04a6.4 6.4 0 0 0-1.21-.63A2.2 2.2 0 0 0 5.5 4.32zm-3.87-1.2a1.11 1.11 0 0 0-.55 1.44a1.08 1.08 0 0 0 1.15 1a1.11 1.11 0 0 0 .55-1.44a1.08 1.08 0 0 0-1.15-1zm7.74 0c.541.254.784.89.55 1.44a1.08 1.08 0 0 1-1.15 1a1.11 1.11 0 0 1-.55-1.44a1.08 1.08 0 0 1 1.15-1zM3.77 1a1 1 0 0 0-.55 1.32a1.26 1.26 0 0 0 1.16 1.09a1 1 0 0 0 .558-1.3l-.008-.02A1.26 1.26 0 0 0 3.77 1zm3.46 0a1 1 0 0 1 .558 1.3l-.008.02a1.26 1.26 0 0 1-1.16 1.09a1 1 0 0 1-.559-1.3l.009-.02A1.26 1.26 0 0 1 7.23 1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VeterinaryFifteen(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M7.5 6c-2.5 0-3 2.28-3 3.47a6.15 6.15 0 0 0-1.7.89a2 2 0 0 0-.4 2.78a2.06 2.06 0 0 0 2.8.4a4 4 0 0 1 2.3-.69a4 4 0 0 1 2.3.69a2 2 0 0 0 2.8-.3a1.93 1.93 0 0 0-.3-2.78l-.1-.1a8.996 8.996 0 0 0-1.7-.89C10.5 8.29 10 6 7.5 6z" fill="currentColor"/><path d="M2.08 4.3a1.58 1.58 0 0 0-.76 2a1.52 1.52 0 0 0 1.61 1.4a1.58 1.58 0 0 0 .76-2a1.52 1.52 0 0 0-1.61-1.4z" fill="currentColor"/><path d="M12.93 4.3a1.58 1.58 0 0 1 .76 2a1.52 1.52 0 0 1-1.61 1.4a1.58 1.58 0 0 1-.76-2a1.52 1.52 0 0 1 1.61-1.4z" fill="currentColor"/><path d="M5.08 1.3c-.68.09-1 .94-.76 1.87A1.77 1.77 0 0 0 5.93 4.7c.68-.09 1-.94.76-1.87A1.77 1.77 0 0 0 5.08 1.3z" fill="currentColor"/><path d="M9.93 1.3c.68.09 1 .94.76 1.87A1.77 1.77 0 0 1 9.07 4.7c-.68-.08-1-.94-.76-1.87A1.77 1.77 0 0 1 9.93 1.3z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Viewpoint(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.02 8.425a2.386 2.386 0 0 0-.46.44l-4.55-3.5a7.998 7.998 0 0 1 1.51-1.51Zm6.46-4.56l-3.5 4.55a2.397 2.397 0 0 1 .45.45l4.56-3.5a7.945 7.945 0 0 0-1.51-1.5m-5.176 6.148a1.5 1.5 0 1 0 1.683 1.291a1.5 1.5 0 0 0-1.683-1.291M6.43 2.235a7.933 7.933 0 0 0-2.06.55l2.2 5.32a2.044 2.044 0 0 1 .61-.17Zm2.14.01l-.75 5.69a2.49 2.49 0 0 1 .61.16l2.2-5.3a7.213 7.213 0 0 0-2.06-.55"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewpointEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M4.403 5.777a.852.852 0 0 0-.183.215L.75 3.322a5.78 5.78 0 0 1 1.1-1.11zM9.15 2.222l-2.544 3.56a1.379 1.379 0 0 1 .184.2l3.47-2.66a5.78 5.78 0 0 0-1.11-1.1zM5.37 7.009a1 1 0 1 0 1.122.86a1 1 0 0 0-1.122-.86zm-.65-5.987a5.774 5.774 0 0 0-1.52.41l1.768 4.003a.815.815 0 0 1 .238-.061zm1.56.01l-.454 4.353a.761.761 0 0 1 .206.044L7.79 1.431a5.519 5.519 0 0 0-1.51-.4z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Village(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.176 1.176a.249.249 0 0 0-.352 0l-4.4 4.4A.25.25 0 0 0 1.6 6H3v6.751a.25.25 0 0 0 .249.249h3.5A.248.248 0 0 0 7 12.753v-7.43c0-.066.026-.13.073-.176L8.5 3.5zM6 11H5v-1h1zm0-2H5V8h1zm0-3v1H5V6zm6.75-3h-.5a.25.25 0 0 0-.25.25V5l-1.324-1.824a.249.249 0 0 0-.352 0L8.056 5.932A.246.246 0 0 0 8 6.088v6.66a.249.249 0 0 0 .246.252h1.5a.253.253 0 0 0 .254-.252V11h1v1.747a.253.253 0 0 0 .253.253h1.5a.25.25 0 0 0 .247-.249V3.25a.25.25 0 0 0-.25-.25M10 8H9V7h1zm2 0h-1V7h1zm-2 2H9V9h1zm2 0h-1V9h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VillageEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M2.777 2.3L.3 5.6a.272.272 0 0 0-.05.15A.25.25 0 0 0 .5 6H1v3.745a.255.255 0 0 0 .255.255h2.49A.255.255 0 0 0 4 9.745V5.467a.253.253 0 0 1 .064-.167L5 4.5L3.2 2.316a.255.255 0 0 0-.423-.016zM3 7H2V6h1zm6.75-4h-.5a.25.25 0 0 0-.25.25V5.2L7.658 4.126a.253.253 0 0 0-.316 0L5.1 5.926a.253.253 0 0 0-.095.2v3.621a.253.253 0 0 0 .248.253h1.494A.253.253 0 0 0 7 9.747V8h1v1.747a.253.253 0 0 0 .253.253h1.494A.253.253 0 0 0 10 9.747V3.25A.25.25 0 0 0 9.75 3zM7 7H6V6h1zm2 0H8V6h1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Volcano(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.484 1a.502.502 0 0 0-.375.188L6.45 3.262L4.844 1.739c-.436-.402-1.084.154-.754.646l1.666 2.4a.5.5 0 0 0 .803.026c.166-.21.441-.678.941-.678s.793.49.941.678a.5.5 0 0 0 .713.072l1.668-1.4c.441-.375 0-1.074-.527-.838L9 3.227V1.5a.5.5 0 0 0-.516-.5M5 6l-2.924 5.924C2 12.001 2 12.155 2 12.309c0 .538.384.691.691.691h9.618c.384 0 .691-.153.691-.691c0-.154 0-.231-.076-.385L10 6c-.5 0-1 .5-1 1v.5a.499.499 0 1 1-1 0V7a.5.5 0 1 0-1 0v2a.499.499 0 1 1-1 0V7c0-.5-.5-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolcanoEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M3 1l1.5 3h2L8 3V2L6 3V1h-.5L5 2.5L3.5 1H3zm.522 4L1.023 9.16c-.223.37.044.84.476.84h8c.432 0 .7-.47.477-.84L7.479 5H7v.5a.499.499 0 1 1-1 0a.5.5 0 0 0-1 0v2a.499.499 0 1 1-1 0V5h-.478z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Volleyball(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.91 7.94a5.909 5.909 0 0 1-.13.94a8.822 8.822 0 0 0-2.17-1.49a8.89 8.89 0 0 0-1.4-.56a8.614 8.614 0 0 0-1.82-.35a.221.221 0 0 0-.08-.01a7.32 7.32 0 0 1-.81-1.88c-.01-.01-.01-.01 0-.02a7.757 7.757 0 0 1 1.52.15a7.099 7.099 0 0 1 1.56.5a6.859 6.859 0 0 1 1.4.78a7.417 7.417 0 0 1 1.93 1.94m-1.87-5a6.223 6.223 0 0 0-1.85-1.27a4.273 4.273 0 0 0-.98-.36a5.641 5.641 0 0 0-1.55-.23a1.288 1.288 0 0 0-.32 0a7.426 7.426 0 0 0-.13 1.42c0 .2.01.39.02.58a2.28 2.28 0 0 1 .26-.01a2.644 2.644 0 0 1 .28.01a8.287 8.287 0 0 1 1.49.16a8.988 8.988 0 0 1 1.35.37a8.792 8.792 0 0 1 1.53.7a10.18 10.18 0 0 1 1.39.97a6.459 6.459 0 0 0-1.49-2.34M4.01 8.76a9.408 9.408 0 0 1-.62-1.37a9.104 9.104 0 0 1-.37-1.41a9.175 9.175 0 0 1-.16-1.67v-.03a8.455 8.455 0 0 1 .1-1.34A6.28 6.28 0 0 0 1.47 5.3a6.222 6.222 0 0 0-.4 2.2c0 .15.01.3.02.44a5.909 5.909 0 0 0 .13.94a5.734 5.734 0 0 0 .41 1.23c.22-.05.44-.11.65-.17a8.019 8.019 0 0 0 1.9-.88c-.06-.1-.12-.2-.17-.3m2.6-2.28a9 9 0 0 1-.63-1.76a9.135 9.135 0 0 1-.24-1.48c-.02-.25-.03-.49-.03-.74a9.378 9.378 0 0 1 .08-1.19a4.273 4.273 0 0 0-.98.36a7.704 7.704 0 0 0-.42 1.94c-.02.22-.03.44-.03.67a7.266 7.266 0 0 0 .06.93a7.426 7.426 0 0 0 .37 1.63a6.47 6.47 0 0 0 .59 1.28a8.017 8.017 0 0 0 1.27-1.57a.252.252 0 0 1-.04-.07m6.09 3.45s0-.01-.01 0a7.013 7.013 0 0 0-1.7-1.16c-.06.1-.11.19-.17.29a8.783 8.783 0 0 1-.88 1.22a8.913 8.913 0 0 1-1.06 1.07a10.418 10.418 0 0 1-1.38.98a8.993 8.993 0 0 1-2.64 1.02a6.299 6.299 0 0 0 5.28 0a5.967 5.967 0 0 0 2.42-1.89a4.826 4.826 0 0 0 .65-1.03a6.032 6.032 0 0 0-.51-.5m-5.17-2c-.01.01-.02.03-.03.04a9.26 9.26 0 0 1-1.25 1.4a9.593 9.593 0 0 1-1.19.91a9.268 9.268 0 0 1-2.62 1.18a6.099 6.099 0 0 0 .53.61a7.735 7.735 0 0 0 3.15-.72a7.384 7.384 0 0 0 1.38-.82a7.643 7.643 0 0 0 1.25-1.16a7.744 7.744 0 0 0 .82-1.15a7.341 7.341 0 0 0-2.04-.29"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolleyballEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.99 5.73a3.657 3.657 0 0 1-.1.74a6.322 6.322 0 0 0-2.45-1.42a5.954 5.954 0 0 0-1.32-.27c-.03 0-.07-.01-.11-.01a4.605 4.605 0 0 1-.51-1.13a1.056 1.056 0 0 1-.06-.24h.12a4.878 4.878 0 0 1 1.01.11a5.039 5.039 0 0 1 1.14.36a4.834 4.834 0 0 1 .94.53a5.136 5.136 0 0 1 1.34 1.33zm-.66 1.58a5.196 5.196 0 0 0-1.36-.97c-.05.1-.11.19-.17.29a5.864 5.864 0 0 1-.59.82a7.18 7.18 0 0 1-.78.78a6.72 6.72 0 0 1-.93.65a6.364 6.364 0 0 1-1.86.72a4.524 4.524 0 0 0 3.72 0a4.158 4.158 0 0 0 1.25-.86a3.198 3.198 0 0 0 .43-.47a3.747 3.747 0 0 0 .48-.76a1.192 1.192 0 0 0-.19-.2zm-.65-4.99a4.508 4.508 0 0 0-1.24-.88a3.979 3.979 0 0 0-.75-.28a4.342 4.342 0 0 0-1.03-.15a1.279 1.279 0 0 0-.32 0a5.363 5.363 0 0 0-.09.99c0 .14.01.28.02.41a1.617 1.617 0 0 1 .22-.01a1.924 1.924 0 0 1 .24.01a6.65 6.65 0 0 1 1 .11a5.695 5.695 0 0 1 1 .28a6.781 6.781 0 0 1 1.02.47a6.665 6.665 0 0 1 .98.69a4.566 4.566 0 0 0-1.05-1.64zM3.03 6.34a5.808 5.808 0 0 1-.4-.92a5.785 5.785 0 0 1-.28-1.03a6.86 6.86 0 0 1-.1-1.12v-.02a5.656 5.656 0 0 1 .07-.93a4.394 4.394 0 0 0-1.05 1.65A4.32 4.32 0 0 0 1 5.5a1.767 1.767 0 0 0 .01.23a4.416 4.416 0 0 0 .09.73a4.075 4.075 0 0 0 .31.91a1.069 1.069 0 0 0 .27-.07a5.068 5.068 0 0 0 1.52-.67c-.06-.1-.12-.19-.17-.29zm3.88-.4a5.5 5.5 0 0 0-1.35-.19c-.02.03-.04.05-.06.08a6.247 6.247 0 0 1-.91 1.01a5.833 5.833 0 0 1-.8.61a6.573 6.573 0 0 1-1.83.82a3.198 3.198 0 0 0 .43.47a5.172 5.172 0 0 0 2.18-.51a5.415 5.415 0 0 0 .93-.55a5.596 5.596 0 0 0 .91-.84A4.805 4.805 0 0 0 7 6.01a.052.052 0 0 0 .02-.04a.228.228 0 0 0-.11-.03zM4.88 4.78a5.64 5.64 0 0 1-.45-1.27a6.386 6.386 0 0 1-.16-.99c-.01-.17-.02-.35-.02-.52a5.83 5.83 0 0 1 .06-.84a3.982 3.982 0 0 0-.75.28a5.58 5.58 0 0 0-.29 1.36c-.01.15-.02.3-.02.45a4.215 4.215 0 0 0 .04.62a4.71 4.71 0 0 0 .27 1.18a5.098 5.098 0 0 0 .42.92a.052.052 0 0 0 .02.04a.178.178 0 0 0 .08-.07h.01a5.682 5.682 0 0 0 .84-1.06a.313.313 0 0 1-.05-.1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Warehouse(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 5a.5.5 0 0 1-.22-.05L7.5 2L1.72 4.93A.514.514 0 0 1 1.28 4L7.5.92L13.72 4a.512.512 0 0 1-.22 1M5 10H2v3h3zm4 0H6v3h3zm4 0h-3v3h3zm-2-4H8v3h3zM7 6H4v3h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WarehouseEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M3 11H0V8h3v3zm4-3H4v3h3V8zm4 0H8v3h3V8zM5 4H2v3h3V4zm4 0H6v3h3V4zm1.44-.76a.5.5 0 0 0-.19-.68L5.75.06a.5.5 0 0 0-.49 0l-4.5 2.5a.5.5 0 0 0 .49.87L5.5 1.07l4.26 2.37a.5.5 0 0 0 .679-.198l.001-.002z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WasteBasket(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.41 5.58l-1.34 8a.5.5 0 0 1-.49.41H4.42a.5.5 0 0 1-.49-.41l-1.34-8A.5.5 0 0 1 3.08 5h8.83a.5.5 0 0 1 .5.58M13 3.5a.5.5 0 0 1-.5.5h-10a.5.5 0 0 1 0-1H5V1.5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5V3h2.5a.5.5 0 0 1 .5.5M9 3V2H6v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WasteBasketEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9 4l-1.17 7H3.17L2 4h7zm.25-1.75A.25.25 0 0 1 9 2.5H2A.25.25 0 1 1 2 2h2V0h3v2h2a.25.25 0 0 1 .25.25zM6.5 2V.5h-2V2h2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Watch(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 7H12a.426.426 0 0 0-.049.01A4.484 4.484 0 0 0 10 3.76V1.5a.5.5 0 0 0-.5-.5h-4a.5.5 0 0 0-.5.5v2.26a4.5 4.5 0 0 0 0 7.48v2.26a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-2.26a4.486 4.486 0 0 0 1.951-3.25A.426.426 0 0 0 12 8h.5a.5.5 0 0 0 0-1m-5 4A3.5 3.5 0 1 1 11 7.5A3.5 3.5 0 0 1 7.5 11M9 7H8V5.5a.5.5 0 0 0-1 0v2a.5.5 0 0 0 .5.5H9a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WatchEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.077 5.154h-.112A3.483 3.483 0 0 0 7 2.35V1H4v1.35c-1.18.563-2 1.756-2 3.15s.82 2.587 2 3.15V10h3V8.65a3.482 3.482 0 0 0 1.965-2.804h.112a.346.346 0 1 0 0-.692zM5.5 8a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5z" fill="currentColor"/><path d="M6.5 5H6V4a.5.5 0 1 0-1 0v1.5a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 0-1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WatchFifteen(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M12.5 7H12c-.018 0-.032.008-.05.01A4.485 4.485 0 0 0 10 3.76V1.498A.498.498 0 0 0 9.502 1H5.498A.498.498 0 0 0 5 1.498V3.76c-1.205.807-2 2.18-2 3.74s.795 2.933 2 3.74v2.262c0 .275.223.498.498.498h4.004a.498.498 0 0 0 .498-.498V11.24a4.485 4.485 0 0 0 1.95-3.25c.018.002.033.01.05.01h.5a.5.5 0 0 0 0-1zm-5 4a3.5 3.5 0 1 1 0-7a3.5 3.5 0 0 1 0 7z" fill="currentColor"/><path d="M9 7H8V5.5a.5.5 0 1 0-1 0v2a.5.5 0 0 0 .5.5H9a.5.5 0 0 0 0-1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Water(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 14c2.077 0 4.5-1.288 4.5-4.568C12 7.207 8.538 2.288 7.5 1C6.577 2.288 3 7.09 3 9.432C3 12.712 5.423 14 7.5 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WaterEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M5.5 11C3.59 11 2 9 2 7s2.61-5.81 3.5-7C6.39 1.19 9 5 9 7s-1.59 4-3.5 4z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Waterfall(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 1H5a3 3 0 0 0-3 3v4.88a2.25 2.25 0 0 0 2.5 3.742a2.247 2.247 0 0 0 2.664-.122h.353A3.25 3.25 0 1 0 10.5 6.75V5a2 2 0 0 1 2-2H14zm-2.5 8.75a2.25 2.25 0 0 1-3.664 1.75H6.75a1.248 1.248 0 0 1-2 0h-.5A1.25 1.25 0 1 1 3 9.525V5.75a.75.75 0 0 1 1.5 0V9a.5.5 0 0 0 1 0V6.75a.75.75 0 0 1 1.5 0V9a.5.5 0 0 0 1 0V5.75a.75.75 0 0 1 1.5 0v1.764a2.25 2.25 0 0 1 2 2.236"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WaterfallEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.9 2H11V1H4C2.4 1 1 2.2 1 3.9v3.4C0 7.8-.3 9 .3 10c.6 1 1.8 1.3 2.7.7c.7.4 1.7.3 2.3-.2c1.4.9 3.2.6 4.2-.8c.9-1.4.6-3.2-.8-4.2c-.2-.1-.5-.2-.7-.3V4c0-1.1.8-2 1.9-2zM9 8c0 1.1-.9 2-2 2c-1 0-1.1-.3-1.3-.5h-.8c-.2.2-.4.5-.9.5s-.7-.2-.8-.5h-.4c-.1.3-.4.5-.8.5c-.6 0-1-.4-1-1s.4-1 1-1V5s0-.5.5-.5s.5.5.5.5v2.5s0 .5.5.5s.5-.5.5-.5V6s0-.5.5-.5s.5.5.5.5v1.5s0 .5.5.5s.5-.5.5-.5V5s0-.5.5-.5s.5.5.5.5v1c1.1 0 2 .9 2 2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Watermill(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 10.2L4 9l2.5 1.5l1.466-.879a3.5 3.5 0 0 1-.843-.567L8.536 7.64c.231.196.52.32.839.347v1.237l.25.15V7.987c.32-.026.608-.151.839-.347l1.413 1.414a3.48 3.48 0 0 1-1.453.8l1.076.646l.698-.42A4.478 4.478 0 0 0 14 6.5A4.5 4.5 0 0 0 9.5 2a4.474 4.474 0 0 0-3.298 1.459L4.5 2L1 5h1zm10.054-1.323L10.64 7.464c.196-.231.32-.52.347-.839h2a3.474 3.474 0 0 1-.933 2.252m.933-2.502h-2a1.484 1.484 0 0 0-.346-.839l1.413-1.413c.555.596.902 1.382.933 2.252M9.625 3.013c.87.031 1.656.378 2.252.933L10.464 5.36a1.484 1.484 0 0 0-.839-.346zm-.25 0v2a1.48 1.48 0 0 0-.839.346L7.123 3.946a3.475 3.475 0 0 1 2.252-.933m-2.429 1.11L8.36 5.536c-.196.231-.32.52-.347.839h-2c.031-.87.378-1.656.933-2.252m1.067 2.502c.026.32.151.608.347.839L6.946 8.877a3.475 3.475 0 0 1-.933-2.252zM9.002 10L6.5 11.5L4 10l-2.5 1.5V13L4 11.5L6.5 13l2.502-1.5L11.5 13l2.5-1.5V10l-2.5 1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WatermillEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9.987 4.375a3.487 3.487 0 0 0-3.362-3.362V1h-.25v.013A3.476 3.476 0 0 0 3.7 2.43L2.506 1L0 4h1v4l1.8-1l1.8 1l1.12-.622A2.98 2.98 0 0 1 4.476 6.7l1.776-1.775a.482.482 0 0 0 .123.05v2.04L6.401 7l.224.124v-2.15a.482.482 0 0 0 .123-.05L8.524 6.7c-.381.35-.85.602-1.371.718l.593.33a3.477 3.477 0 0 0 2.241-3.123H10v-.25h-.013zM4.3 6.524a2.977 2.977 0 0 1-.787-1.899h2.512a.482.482 0 0 0 .05.123L4.3 6.524zm1.725-2.149H3.513c.03-.732.322-1.393.787-1.899l1.775 1.776a.482.482 0 0 0-.05.123zm.35-.35a.482.482 0 0 0-.123.05L4.476 2.3a2.977 2.977 0 0 1 1.899-.787v2.512zm.25 0V1.513c.732.03 1.393.322 1.899.787L6.748 4.075a.482.482 0 0 0-.123-.05zM8.7 6.523L6.925 4.748a.482.482 0 0 0 .05-.123h2.512A2.977 2.977 0 0 1 8.7 6.524zM6.975 4.376a.482.482 0 0 0-.05-.123L8.7 2.476c.465.506.757 1.167.787 1.899H6.975zM6.4 8L4.6 9l-.792-.44L2.8 8l-1.008.56L1 9v1l1.8-1l1.8 1l1.801-1L8.2 10L10 9V8L8.2 9L6.401 8z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wetland(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.48 4.5A2.18 2.18 0 0 1 3 4a2.36 2.36 0 0 1 2.5 2l.78 4.68a2.3 2.3 0 0 0-2 .05L3.48 6a2.21 2.21 0 0 0-2-1.5m6 6.74a3.41 3.41 0 0 1 1.3-.65L10 3a2.21 2.21 0 0 1 2-1.5a2.18 2.18 0 0 0-1.5-.5A2.36 2.36 0 0 0 8 3l-1.3 7.79c.279.115.54.267.78.45m4.32-.5c.157-.128.324-.241.5-.34L13 6a2.21 2.21 0 0 1 2-1.5a2.18 2.18 0 0 0-1.5-.5A2.36 2.36 0 0 0 11 6l-.67 4a3.41 3.41 0 0 1 1.47.74M14 12a1.78 1.78 0 0 0-1.19.42l-.47.41a.75.75 0 0 1-1 0c-.15-.12-.29-.26-.44-.39a1.9 1.9 0 0 0-2.45 0c-.16.13-.31.28-.47.41a.75.75 0 0 1-1 0c-.16-.13-.31-.28-.47-.41a1.9 1.9 0 0 0-2.44 0c-.15.13-.29.27-.44.39a.92.92 0 0 1-.3.16a.84.84 0 0 1-.8-.25a6.167 6.167 0 0 0-.79-.58a1.23 1.23 0 0 0-.68-.16H1a.5.5 0 0 0 0 1a.93.93 0 0 1 .64.31l.36.26a1.9 1.9 0 0 0 2.33.06c.19-.14.36-.32.55-.47a.75.75 0 0 1 1 0l.39.35a1.91 1.91 0 0 0 2.46.07c.15-.11.27-.25.42-.37a.77.77 0 0 1 1.089-.011l.011.011l.39.35a1.89 1.89 0 0 0 1.76.37a2.14 2.14 0 0 0 1-.6A1 1 0 0 1 14 13a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WetlandEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M1.11 3.33A1.58 1.58 0 0 1 2.19 3A1.71 1.71 0 0 1 4 4.41l.57 3.39a1.66 1.66 0 0 0-1.44 0l-.57-3.39a1.6 1.6 0 0 0-1.45-1.08zm4.34 4.88a2.47 2.47 0 0 1 .94-.47l.9-5.5a1.6 1.6 0 0 1 1.45-1.08A1.58 1.58 0 0 0 7.65.8a1.71 1.71 0 0 0-1.81 1.44L4.9 7.88c.197.086.382.197.55.33zm3.13-.36c.113-.094.233-.177.36-.25l.5-3.21a1.6 1.6 0 0 1 1.45-1.09a1.58 1.58 0 0 0-1.09-.36A1.71 1.71 0 0 0 8 4.39l-.48 3c.381.069.74.226 1.05.46h.01zm1.56.9a1.29 1.29 0 0 0-.86.3l-.34.3a.54.54 0 0 1-.7 0l-.33-.28a1.38 1.38 0 0 0-1.77 0l-.34.3a.54.54 0 0 1-.69 0l-.34-.32A1.38 1.38 0 0 0 3 9.06c-.11.09-.21.19-.32.28a.67.67 0 0 1-.22.11a.61.61 0 0 1-.58-.18a4.465 4.465 0 0 0-.57-.42a.89.89 0 0 0-.51-.1a.36.36 0 0 0 0 .72a.67.67 0 0 1 .46.23l.22.2a1.37 1.37 0 0 0 1.69 0c.14-.1.26-.23.4-.34a.54.54 0 0 1 .71 0l.25.29a1.38 1.38 0 0 0 1.78 0c.11-.08.2-.18.3-.27a.56.56 0 0 1 .79 0l.29.3c.36.311.856.414 1.31.27a1.55 1.55 0 0 0 .69-.43a.68.68 0 0 1 .49-.25a.366.366 0 1 0-.05-.73l.01.01z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wheelchair(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0M2.82 4.87l1.74-1.71l1.85 1.29l-.74 1.25a5.49 5.49 0 0 1 2.71 1.88l.88-1.5a.86.86 0 0 0-.34-1.17l-.62-.37L4.79 2.1a.5.5 0 0 0-.64.05l-2 2a.5.5 0 0 0 .67.72m5.31 6.07a4.11 4.11 0 0 1-1.54 3.16a4 4 0 0 1-2.52.9A4.07 4.07 0 0 1 0 10.94a4 4 0 0 1 .91-2.53a4.17 4.17 0 0 1 .63-.63a4 4 0 0 1 2.53-.91a4.07 4.07 0 0 1 4.06 4.07M6 12.21A2.29 2.29 0 0 0 2.79 9a1.72 1.72 0 0 0-.63.63a2.29 2.29 0 0 0 3.17 3.17a3.14 3.14 0 0 0 .67-.59m8.2 1l-2.49-5a.75.75 0 0 0-.71-.46H8.51c.33.458.586.964.76 1.5h1.26l2.29 4.58a.754.754 0 1 0 1.35-.67z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WheelchairEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M9 2a1 1 0 1 1-2 0a1 1 0 0 1 2 0zM2.5 4a.5.5 0 0 0 .18 0l1.79-.83l.46.48l-.54.85a3.74 3.74 0 0 1 2.07 1.06l.44-.8a.64.64 0 0 0-.2-.89L4.85 2.15A.49.49 0 0 0 4.3 2l-2 1a.51.51 0 0 0 .2 1zm4.09 4.2a2.74 2.74 0 0 1-.51 1.57a2.42 2.42 0 0 1-.71.71a2.74 2.74 0 0 1-1.58.52A2.8 2.8 0 0 1 1 8.2a2.74 2.74 0 0 1 .52-1.58c.206-.266.445-.504.71-.71a2.74 2.74 0 0 1 1.56-.51a2.8 2.8 0 0 1 2.8 2.8zm-1.4.69a1.54 1.54 0 0 0 .18-.69a1.58 1.58 0 0 0-1.58-1.57a1.54 1.54 0 0 0-.69.18a2.14 2.14 0 0 0-.71.7a1.54 1.54 0 0 0-.18.7a1.58 1.58 0 0 0 1.58 1.57a1.54 1.54 0 0 0 .7-.18a1.9 1.9 0 0 0 .7-.71zM9 6.4a.49.49 0 0 0-.53-.4H6.82c.223.305.398.642.52 1H8l1 2.59a.5.5 0 0 0 .49.4h.1a.5.5 0 0 0 .413-.574L10 9.4l-1-3z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Windmill(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 5L3.5 7.5l1 1L7 6v-.5l.5-.5l.5.5V6l2.5 2.5l1-1L9 5h-.5L8 4.5l.5-.5H9l2.5-2.5l-1-1L8 3v.5l-.5.5l-.5-.5V3L4.5.5l-1 1L6 4h.5l.5.5l-.5.5zm5.5 8h-1l-1-4l-2-2l-2 2l-1 4h-1a.5.5 0 1 0 0 1h8a.5.5 0 0 0 0-1M8 13H7v-1.502c0-.275.223-.498.498-.498c.277 0 .502.225.502.502z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindmillEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M6.375 3.938L5.937 3.5l.438-.438h.438L9 .876L8.125 0L5.937 2.188v.437l-.437.438l-.438-.438v-.438L2.876 0L2 .875l2.188 2.188h.437l.438.437l-.438.438h-.438L2 6.124L2.875 7l2.188-2.188v-.437l.437-.438l.438.438v.438L8.124 7L9 6.125L6.812 3.937h-.437zM8.5 10H8L7 7.5L5.5 6L4 7.5L3 10h-.5a.5.5 0 1 0 0 1h6a.5.5 0 0 0 0-1zM6 9.972a.028.028 0 0 1-.028.028h-.944A.028.028 0 0 1 5 9.972V9a.5.5 0 1 1 1 0v.972z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Zoo(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.388 3.893S7.682 5 7 5H4.5c-.682 0-1.298.409-1.773.864L.5 8c-.269.258-.493.768-.5 1c-.016.5 0 1 0 1s1 0 1-1v-.5l1-1h.5l.158.29s-.953 1.824-.953 3.528c0 .682.681.682.681.682h.682s.341 0 0-.34l-.34-.342C2.727 10.636 3.5 9.667 4 9c0 0 .017 1.158 0 2c-.008.41.272 1 .682 1h.682s.34 0 0-.34l-.341-.342C4.71 10.75 5.5 8.5 5.5 8.5c1.272 0 1.5.5 3 .5l.364 2.318c.111.71.636.682.636.682H11c.34 0 .961-.311.34-.774L11 11V9c1.214-.172 2-1 2-2h1c.321 0 1 0 1-.5v-1l-1.704-1.682C12.544 3.078 12 2.5 11 2.5c-1.453 0-2.136.569-2.612 1.393"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZooEleven(children ...ElementRenderer) *MakiIcon {
	return &MakiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M8 2c-.9 0-1.5.5-1.8 1.1c0 0-.7.9-1.2.9H3c-.5 0-1 .5-1 .5l-2 2V7h.5l1.2-1.2l.2.2S1 7.7 1 9c0 .5.5.5.5.5H2s.2 0 0-.2l-.2-.2c0-.5.8-1.4 1.2-1.9v1.4c0 .3.111.9.411.9h.5s.2 0 0-.2l-.2-.2c-.2-.4.589-1.6.589-1.6h1.8L6.6 9c.2.5.5.5.5.5h1c.2 0 .7-.2.2-.5l-.2-.2V7c1 0 1.1-1.3 1.5-1.7l.7-.1c.2 0 .8-.2.8-.8V4L9.9 3C9.4 2.4 8.7 2 8 2z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
