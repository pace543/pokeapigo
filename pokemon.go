package pokeapigo

type Ability struct {
	Id                  int
	Name                string
	Is_Main_Series      bool
	Generation          NamedAPIResource
	Names               []Name
	Effect_Entries      []VerboseEffect
	Effect_Changes      []AbilityEffectChange
	Flavor_Text_Entries []AbilityFlavorText
	Pokemon             []AbilityPokemon
}

type AbilityEffectChange struct {
	Effect_Entries []Effect
	Version_Group  NamedAPIResource
}

type AbilityFlavorText struct {
	Flavor_Text   string
	Language      NamedAPIResource
	Version_Group NamedAPIResource
}

type AbilityPokemon struct {
	Is_Hidden bool
	Slot      int
	Pokemon   NamedAPIResource
}

type Characteristic struct {
	Id              int
	Gene_Modulo     int
	Possible_Values []int
	Descriptions    []Description
}

type EggGroup struct {
	Id              int
	Name            string
	Names           []Name
	Pokemon_Species []NamedAPIResource
}

type Gender struct {
	Id                      int
	Name                    string
	Pokemon_Species_Details []PokemonSpeciesGender
	Required_For_Evolution  []NamedAPIResource
}

type PokemonSpeciesGender struct {
	Rate            int
	Pokemon_Species NamedAPIResource
}

type GrowthRate struct {
	Id              int
	Name            string
	Formula         string
	Descriptions    []Description
	Levels          []GrowthRateExperienceLevel
	Pokemon_Species []NamedAPIResource
}

type GrowthRateExperienceLevel struct {
	Level      int
	Experience int
}

type Nature struct {
	Id                            int
	Name                          string
	Decreased_Stat                NamedAPIResource
	Increased_Stat                NamedAPIResource
	Hates_Flavor                  NamedAPIResource
	Likes_Flavor                  NamedAPIResource
	Pokeathlon_Stat_Changes       []NatureStatChange
	Move_Battle_Style_Preferences []MoveBattleStylePreference
	Names                         []Name
}

type NatureStatChange struct {
	Max_Change      int
	Pokeathlon_Stat NamedAPIResource
}

type MoveBattleStylePreference struct {
	Low_HP_Preference  int
	High_HP_Preference int
	Move_Battle_Style  NamedAPIResource
}

type PokeathlonStat struct {
	Id                int
	Name              string
	Names             []Name
	Affecting_Natures NaturePokeathlonStatAffectSets
}

type NaturePokeathlonStatAffectSets struct {
	Increase []NaturePokeathlonStatAffect
	Decrease []NaturePokeathlonStatAffect
}

type NaturePokeathlonStatAffect struct {
	Max_Change int
	Nature     NamedAPIResource
}

type Pokemon struct {
	Id                       int
	Name                     string
	Base_Experience          int
	Height                   int
	Is_Default               bool
	Order                    int
	Weight                   int
	Abilities                []PokemonAbility
	Forms                    []NamedAPIResource
	Game_Indices             []VersionGameIndex
	Held_Items               []PokemonHeldItem
	Location_Area_Encounters string
	Moves                    []PokemonMove
	Sprites                  PokemonSprites
	Species                  NamedAPIResource
	Stats                    []PokemonStat
	Types                    []PokemonType
}

type PokemonAbility struct {
	Is_Hidden bool
	Slot      int
	Ability   NamedAPIResource
}

type PokemonType struct {
	Slot int
	Type NamedAPIResource
}

type PokemonHeldItem struct {
	Item            NamedAPIResource
	Version_Details []PokemonHeldItemVersion
}

type PokemonHeldItemVersion struct {
	Version NamedAPIResource
	Rarity  int
}

type PokemonMove struct {
	Move                  NamedAPIResource
	Version_Group_Details []PokemonMoveVersion
}

type PokemonMoveVersion struct {
	Move_Learn_Method NamedAPIResource
	Version_Group     NamedAPIResource
	Level_Learned_At  int
}

type PokemonStat struct {
	Stat      NamedAPIResource
	Effort    int
	Base_Stat int
}

type PokemonSprites struct {
	Front_Default      string
	Front_Shiny        string
	Front_Female       string
	Front_Shiny_Female string
	Back_Default       string
	Back_Shiny         string
	Back_Female        string
	Back_Shiny_Female  string
}

type LocationAreaEncounter struct {
	Location_Area   NamedAPIResource
	Version_Details []VersionEncounterDetail
}

type PokemonColor struct {
	Id              int
	Name            string
	Names           []Name
	Pokemon_Species []NamedAPIResource
}

type PokemonForm struct {
	Id             int
	Name           string
	Order          int
	Form_Order     int
	Is_Default     bool
	Is_Battle_Only bool
	Is_Mega        bool
	Form_Name      string
	Pokemon        NamedAPIResource
	Sprites        PokemonFormSprites
	Version_Group  NamedAPIResource
	Names          []Name
	Form_Names     []Name
}

type PokemonFormSprites struct {
	Front_Default string
	Front_Shiny   string
	Back_Default  string
	Back_Shiny    string
}

type PokemonHabitat struct {
	Id              int
	Name            string
	Names           []Name
	Pokemon_Species []NamedAPIResource
}

type PokemonShape struct {
	Id              int
	Name            string
	Awesome_Names   []AwesomeName
	Names           []Name
	Pokemon_Species []NamedAPIResource
}

type AwesomeName struct {
	Awesome_Name string
	Language     NamedAPIResource
}

type PokemonSpecies struct {
	Id                     int
	Name                   string
	Order                  int
	Gender_Rate            int
	Capture_Rate           int
	Base_Happiness         int
	Is_Baby                bool
	Hatch_Counter          int
	Has_Gender_Differences bool
	Forms_Switchable       bool
	Growth_Rate            NamedAPIResource
	Pokedex_Numbers        []PokemonSpeciesDexEntry
	Egg_Groups             []NamedAPIResource
	Color                  NamedAPIResource
	Shape                  NamedAPIResource
	Evolves_From_Species   NamedAPIResource
	Evolution_Chain        APIResource
	Habitat                NamedAPIResource
	Generation             NamedAPIResource
	Names                  []Name
	Pal_Park_Encounters    []PalParkEncounterArea
	Flavor_Text_Entries    []FlavorText
	Form_Descriptions      []Description
	Genera                 []Genus
	Varieties              []PokemonSpeciesVariety
}

type Genus struct {
	Genus    string
	Language NamedAPIResource
}

type PokemonSpeciesDexEntry struct {
	Entry_Number int
	Pokedex      NamedAPIResource
}

type PalParkEncounterArea struct {
	Base_Score int
	Rate       int
	Area       NamedAPIResource
}

type PokemonSpeciesVariety struct {
	Is_Default bool
	Pokemon    NamedAPIResource
}

type Stat struct {
	Id                int
	Name              string
	Game_Index        int
	Is_Battle_Only    bool
	Affecting_Moves   MoveStatAffectSets
	Affecting_Natures NatureStatAffectSets
	Characteristics   []APIResource
	Move_Damage_Class NamedAPIResource
	Names             []Name
}

type MoveStatAffectSets struct {
	Increase []MoveStatAffect
	Decrease []MoveStatAffect
}

type MoveStatAffect struct {
	Change int
	Move   NamedAPIResource
}

type NatureStatAffectSets struct {
	Increase []NamedAPIResource
	Decrease []NamedAPIResource
}

type Type struct {
	Id                int
	Name              string
	Damage_Relations  TypeRelations
	Game_Indices      []GenerationGameIndex
	Generation        NamedAPIResource
	Move_Damage_Class NamedAPIResource
	Names             []Name
	Pokemon           []TypePokemon
	Moves             []NamedAPIResource
}

type TypePokemon struct {
	Slot    int
	Pokemon NamedAPIResource
}

type TypeRelations struct {
	No_Damage_To       []NamedAPIResource
	Half_Damage_To     []NamedAPIResource
	Double_Damage_To   []NamedAPIResource
	No_Damage_From     []NamedAPIResource
	Half_Damage_From   []NamedAPIResource
	Double_Damage_From []NamedAPIResource
}
