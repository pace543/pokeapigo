package pokeapiGo

type Berry struct {
	Id 			int
	Name			string
	Growth_Time		int
	Max_Harvest		int
	Natural_Gift_Power	int
	Size 			int
	Smoothness		int
	Soil_Dryness		int
	Firmness 		NamedAPIResource
	Flavors 		[]BerryFlavorMap
	Item 			NamedAPIResource
	Natural_Gift_Type 	NamedAPIResource
}

type BerryFlavorMap struct {
	Potency 	int
	Flavor 		NamedAPIResource
}

type BerryFirmness struct {
	Id 		int
	Name 		string
	Berries		[]NamedAPIResource
	Names 		[]Name
}

type BerryFlavor struct {
	Id 		int
	Name 		string
	Berries  	[]FlavorBerryMap
	Contest_Type 	NamedAPIResource
	Names 		[]Name
}

type FlavorBerryMap struct {
	Potency 	int
	Berry 		NamedAPIResource
}