package pokeapiGo

import (
	"fmt"
	"net/http"
)

type ContestType struct {
	Id           int
	Name         string
	Berry_Flavor NamedAPIResource
	Names        []ContestName
}

type ContestName struct {
	Name     string
	Color    string
	Language NamedAPIResource
}

type ContestEffect struct {
	Id                  int
	Appeal              int
	Jam                 int
	Effect_Entries      []Effect
	Flavor_Text_Entries []FlavorText
}

type SuperContestEffect struct {
	Id                  int
	Appeal              int
	Flavor_Text_Entries []FlavorText
	Moves               []NamedAPIResource
}

func GetContestType(p *Param, contestType *ContestType, client *http.Client) error {
	err := Get(client, fmt.Sprintf("/contest-type/%v", p), contestType)
	if err != nil {
		return err
	}
	return nil
}

func GetContestEffect(p *Param, contestEffect *ContestEffect, client *http.Client) error {
	err := Get(client, fmt.Sprintf("/contest-effect/%v", p), contestEffect)
	if err != nil {
		return err
	}
	return nil
}

func GetSuperContestEffect(p *Param, superContestEffect *SuperContestEffect, client *http.Client) error {
	err := Get(client, fmt.Sprintf("/super-contest-effect/%v", p), superContestEffect)
	if err != nil {
		return err
	}
	return nil
}
