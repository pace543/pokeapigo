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

func GetEncounterMethodById(id int, encounterMethod *EncounterMethod, client *http.Client) error {
	err := Get(client, fmt.Sprintf("/encounter-method/%v", id), encounterMethod)
	if err != nil {
		return err
	}
	return nil
}

func GetEncounterMethodByName(name string, encounterMethod *EncounterMethod, client *http.Client) error {
	err := Get(client, fmt.Sprintf("/encounter-method/%v", name), encounterMethod)
	if err != nil {
		return err
	}
	return nil
}

func GetEncounterConditionById(id int, encounterCondition *EncounterCondition, client *http.Client) error {
	err := Get(client, fmt.Sprintf("/encounter-condition/%v", id), encounterCondition)
	if err != nil {
		return err
	}
	return nil
}

func GetEncounterConditionByName(name string, encounterCondition *EncounterCondition, client *http.Client) error {
	err := Get(client, fmt.Sprintf("/encounter-condition/%v", name), encounterCondition)
	if err != nil {
		return err
	}
	return nil
}

func GetEncounterConditionValueById(id int, encounterConditionValue *EncounterConditionValue, client *http.Client) error {
	err := Get(client, fmt.Sprintf("/encounter-condition-value/%v", id), encounterConditionValue)
	if err != nil {
		return err
	}
	return nil
}

func GetEncounterConditionValueByName(name string, encounterConditionValue *EncounterConditionValue, client *http.Client) error {
	err := Get(client, fmt.Sprintf("/encounter-condition-value/%v", name), encounterConditionValue)
	if err != nil {
		return err
	}
	return nil
}
