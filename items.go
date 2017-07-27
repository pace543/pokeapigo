package pokeapigo

type Item struct {
	Id                  int
	Name                string
	Cost                int
	Fling_Power         int
	Fling_Effect        NamedAPIResource
	Attributes          []NamedAPIResource
	Category            ItemCategory
	Effect_Entries      []VerboseEffect
	Flavor_Text_Entries []VersionGroupFlavorText
	Game_Indices        []GenerationGameIndex
	Names               []Name
	Sprites             ItemSprites
	Held_By_Pokemon     []ItemHolderPokemon
	Baby_Trigger_For    APIResource
	Machines            []MachineVersionDetail
}

type ItemSprites struct {
	Default string
}

type ItemHolderPokemon struct {
	Pokemon         string
	Version_Details []ItemHolderPokemonVersionDetail
}

type ItemHolderPokemonVersionDetail struct {
	Rarity  string
	Version NamedAPIResource
}

type ItemAttribute struct {
	Id           int
	Name         string
	Items        []NamedAPIResource
	Names        []Name
	Descriptions []Description
}

type ItemCategory struct {
	Id     int
	Name   string
	Items  []NamedAPIResource
	Names  []Name
	Pocket []NamedAPIResource
}

type ItemFlingEffect struct {
	Id             int
	Name           string
	Effect_Entries []Effect
	Items          []NamedAPIResource
}

type ItemPocket struct {
	Id         int
	Name       string
	Categories []NamedAPIResource
	Names      []Name
}
