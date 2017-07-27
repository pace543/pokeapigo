package pokeapigo

type Generation struct {
	Id              int
	Name            string
	Abilities       []NamedAPIResource
	Names           []Name
	Main_Region     NamedAPIResource
	Moves           []NamedAPIResource
	Pokemon_Species []NamedAPIResource
	Types           []NamedAPIResource
	Version_Groups  []NamedAPIResource
}

type Pokedex struct {
	Id              int
	Name            string
	Is_Main_Series  bool
	Descriptions    []Description
	Names           []Name
	Pokemon_Entries []PokemonEntry
	Region          NamedAPIResource
	Version_Groups  []NamedAPIResource
}

type PokemonEntry struct {
	Entry_Number    int
	Pokemon_Species NamedAPIResource
}

type Version struct {
	Id            int
	Name          string
	Names         []Name
	Version_Group NamedAPIResource
}

type VersionGroup struct {
	Id                 int
	Name               string
	Order              int
	Generation         NamedAPIResource
	Move_Learn_Methods []NamedAPIResource
	Pokedexes          []NamedAPIResource
	Regions            []NamedAPIResource
	Versions           []NamedAPIResource
}
