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

func GetBerryById(id int, berry *Berry, client *http.Client) error {
	err := Get(client, fmt.Sprintf("/berry/%v", id), berry)
	if err != nil {
		return err
	}
	return nil
}

func GetBerryByName(name string, berry *Berry, client *http.Client) error {
	err := Get(client, fmt.Sprintf("/berry/%v", name), berry)
	if err != nil {
		return err
	}
	return nil
}

func GetBerryFirmnessById(id int, berryFirmness *BerryFirmness, client *http.Client) error {
	err := Get(client, fmt.Sprintf("/berry-firmness/%v", id), berryFirmness)
	if err != nil {
		return err
	}
	return nil
}

func GetBerryFirmnessByName(name string, berryFirmness *BerryFirmness, client *http.Client) error {
	err := Get(client, fmt.Sprintf("/berry-firmness/%v", name), berryFirmness)
	if err != nil {
		return err
	}
	return nil
}

func GetBerryFlavorById(id int, berryFlavor *BerryFlavor, client *http.Client) error {
	err := Get(client, fmt.Sprintf("/berry-flavor/%v", id), berryFlavor)
	if err != nil {
		return err
	}
	return nil
}

func GetBerryFlavorByName(name string, berryFlavor *BerryFlavor, client *http.Client) error {
	err := Get(client, fmt.Sprintf("/berry-flavor/%v", name), berryFlavor)
	if err != nil {
		return err
	}
	return nil
}
