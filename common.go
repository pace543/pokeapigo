package pokeapigo

import (
	"strconv"
	"strings"
)

type APIResource struct {
	Url string
}

type APIResourceList struct {
	Count    int
	Next     string
	Previous bool
	Results  []APIResource
}

type NamedAPIResource struct {
	Name string
	Url  string
}

type NamedAPIResourceList struct {
	Count    int
	Next     string
	Previous bool
	Results  []NamedAPIResource
}

// GetJob is called on a *NamedAPIResource and returns a *Job which can then be submitted to the client to receive
// a result containing the data which the *NamedAPIResource refers to.
func (namedResource *NamedAPIResource) GetJob() *Job {
	stringSlice := strings.Split(namedResource.Url, "/")
	num, _ := strconv.Atoi(stringSlice[len(stringSlice)-1])
	return &Job{
		Endpoint: stringSlice[len(stringSlice)-2],
		Id:       num,
	}
}

// GetJob is called on a *APIResource and returns a *Job which can then be submitted to the client to receive
// a result containing the data which the *APIResource refers to.
func (apires *APIResource) GetJob() *Job {
	stringSlice := strings.Split(apires.Url, "/")
	num, _ := strconv.Atoi(stringSlice[len(stringSlice)-1])
	return &Job{
		Endpoint: stringSlice[len(stringSlice)-2],
		Id:       num,
	}
}

type Description struct {
	Description string
	Language    NamedAPIResource
}

type Effect struct {
	Effect   string
	Language NamedAPIResource
}

type Encounter struct {
	Min_Level        int
	Max_Level        int
	Condition_Values []NamedAPIResource
	Chance           int
	Method           NamedAPIResource
}

type FlavorText struct {
	Flavor_Text string
	Language    NamedAPIResource
}

type GenerationGameIndex struct {
	Game_Index int
	Generation NamedAPIResource
}

type MachineVersionDetail struct {
	Machine       APIResource
	Version_Group NamedAPIResource
}

type Name struct {
	Name     string
	Language NamedAPIResource
}

type VerboseEffect struct {
	Effect       string
	Short_Effect string
	Language     NamedAPIResource
}

type VersionEncounterDetail struct {
	Version           NamedAPIResource
	Max_Chance        int
	Encounter_Details []Encounter
}

type VersionGameIndex struct {
	Game_Index int
	Version    NamedAPIResource
}

type VersionGroupFlavorText struct {
	Text          string
	Language      NamedAPIResource
	Version_Group NamedAPIResource
}

type Language struct {
	Id       int
	Name     string
	Official bool
	Iso639   string
	Iso3166  string
	Names    []Name
}
