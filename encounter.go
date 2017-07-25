package pokeapiGo

import (
	"fmt"
	"net/http"
)

type EncounterMethod struct {
	id    int
	name  string
	order int
	names []Name
}

type EncounterCondition struct {
	id     int
	name   string
	names  []Name
	values []NamedAPIResource
}

type EncounterConditionValue struct {
	id        int
	name      string
	condition []NamedAPIResource
	names     []Name
}

func GetEncounterMethod(p *Param, encounterMethod *EncounterMethod, client *http.Client) error {
	err := Get(client, fmt.Sprintf("/encounter-method/%v", p), encounterMethod)
	if err != nil {
		return err
	}
	return nil
}

func GetEncounterCondition(p *Param, encounterCondition *EncounterCondition, client *http.Client) error {
	err := Get(client, fmt.Sprintf("/encounter-condition/%v", p), encounterCondition)
	if err != nil {
		return err
	}
	return nil
}

func GetEncounterConditionValue(p *Param, encounterConditionValue *EncounterConditionValue, client *http.Client) error {
	err := Get(client, fmt.Sprintf("/encounter-condition-value/%v", p), encounterConditionValue)
	if err != nil {
		return err
	}
	return nil
}
