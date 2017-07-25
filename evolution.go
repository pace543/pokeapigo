package pokeapiGo

import (
	"fmt"
	"net/http"
)

type EvolutionChain struct {
	Id                int
	Baby_Trigger_Item NamedAPIResource
	Chain             ChainLink
}

type ChainLink struct {
	Is_Baby           bool
	Species           NamedAPIResource
	Evolution_Details []EvolutionDetail
	Evolves_To        []ChainLink
}

type EvolutionDetail struct {
	Item                    NamedAPIResource
	Trigger                 NamedAPIResource
	Gender                  int
	Held_Item               NamedAPIResource
	Known_Move              NamedAPIResource
	Known_Move_Type         NamedAPIResource
	Location                NamedAPIResource
	Min_Level               int
	Min_Happiness           int
	Min_Beauty              int
	Min_Affection           int
	Needs_Overworld_Rain    bool
	Party_Species           NamedAPIResource
	Party_Type              NamedAPIResource
	Relative_Physical_Stats int
	Time_Of_Day             string
	Trade_Species           NamedAPIResource
	Turn_Upside_Down        bool
}

type EvolutionTrigger struct {
	Id              int
	Name            string
	Names           []Name
	Pokemon_Species []NamedAPIResource
}

func GetEvolutionChain(p *Param, evolutionChain *EvolutionChain, client *http.Client) error {
	err := Get(client, fmt.Sprintf("/evolution-chain/%v", p), evolutionChain)
	if err != nil {
		return err
	}
	return nil
}

func GetEvolutionTrigger(p *Param, evolutionTrigger *EvolutionTrigger, client *http.Client) error {
	err := Get(client, fmt.Sprintf("/evolution-trigger/%v", p), evolutionTrigger)
	if err != nil {
		return err
	}
	return nil
}
