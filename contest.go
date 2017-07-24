package pokeapiGo

type ContestType struct {
	Id 		int
	Name 		string
	Berry_Flavor	NamedAPIResource
	Names		[]ContestName
}

type ContestName struct {
	Name 		string
	Color 		string
	Language 	NamedAPIResource
}

type ContestEffect struct {
	Id 			int
	Appeal 			int
	Jam 			int
	Effect_Entries		[]Effect
	Flavor_Text_Entries 	[]FlavorText
}

type SuperContestEffect struct {
	Id 			int
	Appeal 			int
	Flavor_Text_Entries 	[]FlavorText
	Moves			[]NamedAPIResource
}