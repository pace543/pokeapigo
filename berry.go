package pokeapiGo

import (
	"fmt"
	"net/http"
)

type Berry struct {
	Id                 int
	Name               string
	Growth_Time        int
	Max_Harvest        int
	Natural_Gift_Power int
	Size               int
	Smoothness         int
	Soil_Dryness       int
	Firmness           NamedAPIResource
	Flavors            []BerryFlavorMap
	Item               NamedAPIResource
	Natural_Gift_Type  NamedAPIResource
}

type BerryFlavorMap struct {
	Potency int
	Flavor  NamedAPIResource
}

type BerryFirmness struct {
	Id      int
	Name    string
	Berries []NamedAPIResource
	Names   []Name
}

type BerryFlavor struct {
	Id           int
	Name         string
	Berries      []FlavorBerryMap
	Contest_Type NamedAPIResource
	Names        []Name
}

type FlavorBerryMap struct {
	Potency int
	Berry   NamedAPIResource
}

func GetBerry(p *Param, berry *Berry, client *http.Client) error {
	err := Get(client, fmt.Sprintf("/berry/%v", p), berry)
	if err != nil {
		return err
	}
	return nil
}

func GetBerryFirmness(p *Param, berryFirmness *BerryFirmness, client *http.Client) error {
	err := Get(client, fmt.Sprintf("/berry-firmness/%v", p), berryFirmness)
	if err != nil {
		return err
	}
	return nil
}

func GetBerryFlavor(p *Param, berryFlavor *BerryFlavor, client *http.Client) error {
	err := Get(client, fmt.Sprintf("/berry-flavor/%v", p), berryFlavor)
	if err != nil {
		return err
	}
	return nil
}
