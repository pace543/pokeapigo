package pokeapiGo

type Move struct {
	Id                   int
	Name                 string
	Accuracy             int
	Effect_Chance        int
	Pp                   int
	Priority             int
	Power                int
	Contest_Combos       ContestComboSets
	Contest_Type         NamedAPIResource
	Contest_Effect       APIResource
	Damage_Class         NamedAPIResource
	Effect_Entries       []VerboseEffect
	Effect_Changes       []AbilityEffectChange
	Flavor_Text_Entries  []MoveFlavorText
	Generation           NamedAPIResource
	Machines             []MachineVersionDetail
	Meta                 MoveMetaData
	Names                []Name
	Past_Values          []PastMoveStatValues
	Stat_Changes         []MoveStatChange
	Super_Contest_Effect APIResource
	Target               NamedAPIResource
	Type                 NamedAPIResource
}

type ContestComboSets struct {
	Normal ContestComboDetail
	Super  ContestComboDetail
}

type ContestComboDetail struct {
	Use_Before []NamedAPIResource
	Use_After  []NamedAPIResource
}

type MoveFlavorText struct {
	Flavor_Text   string
	Language      NamedAPIResource
	Version_Group NamedAPIResource
}

type MoveMetaData struct {
	Ailment        NamedAPIResource
	Category       NamedAPIResource
	Min_Hits       int
	Max_Hits       int
	Min_Turns      int
	Max_Turns      int
	Drain          int
	Healing        int
	Crit_Rate      int
	Ailment_Chance int
	Flinch_Chance  int
	Stat_Chance    int
}

type MoveStatChange struct {
	Change int
	Stat   NamedAPIResource
}

type PastMoveStatValues struct {
	Accuracy       int
	Effect_Chance  int
	Power          int
	Pp             int
	Effect_Entries []VerboseEffect
	Type           NamedAPIResource
	Version_Group  NamedAPIResource
}

type MoveAilment struct {
	Id    int
	Name  string
	Moves []NamedAPIResource
	Names []Name
}

type MoveBattleStyle struct {
	Id    int
	Name  string
	Names []Name
}

type MoveCategory struct {
	Id           int
	Name         string
	Moves        []NamedAPIResource
	Descriptions []Description
}

type MoveDamageClass struct {
	Id           int
	Name         string
	Descriptions []Description
	Moves        []NamedAPIResource
	Names        []Name
}

type MoveLearnMethod struct {
	Id             int
	Name           string
	Descriptions   []Description
	Names          []Name
	Version_Groups []NamedAPIResource
}

type MoveTarget struct {
	Id           int
	Name         string
	Descriptions []Description
	Moves        []NamedAPIResource
	Names        []Name
}
