package pokeapigo

type Location struct {
	Id           int
	Name         string
	Region       NamedAPIResource
	Names        []Name
	Game_Indices []GenerationGameIndex
	Areas        []NamedAPIResource
}

type LocationArea struct {
	Id                     int
	Name                   string
	Game_Index             int
	Encounter_Method_Rates []EncounterMethodRate
	Location               NamedAPIResource
	Names                  []Name
	Pokemon_Encounters     []PokemonEncounter
}

type EncounterMethodRate struct {
	Encounter_Method NamedAPIResource
	Version_Details  []EncounterVersionDetails
}

type EncounterVersionDetails struct {
	Rate    int
	Version NamedAPIResource
}

type PokemonEncounter struct {
	Pokemon         NamedAPIResource
	Version_Details []VersionEncounterDetail
}

type PalParkArea struct {
	Id                 int
	Name               string
	Names              []Name
	Pokemon_Encounters []PalParkEncounterSpecies
}

type PalParkEncounterSpecies struct {
	Base_Score      int
	Rate            int
	Pokemon_Species NamedAPIResource
}

type Region struct {
	Id              int
	Name            string
	Locations       []NamedAPIResource
	Main_Generation NamedAPIResource
	Names           []Name
	Pokedexes       []NamedAPIResource
	Version_Groups  []NamedAPIResource
}
