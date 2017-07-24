package pokeapiGo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Get(client *http.Client, url string, target interface{}) error {
	fullUrl := fmt.Sprintf("https://pokeapi.co/api/v2%v", url)
	result, err := client.Get(fullUrl)
	if err != nil {
		return err
	}
	defer result.Body.Close()
	return json.NewDecoder(result.Body).Decode(target)
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

func GetContestTypeById(id int, contestType *ContestType, client *http.Client) error {
	err := Get(client, fmt.Sprintf("/contest-type/%v", id), contestType)
	if err != nil {
		return err
	}
	return nil
}

func GetContestTypeByName(name string, contestType *ContestType, client *http.Client) error {
	err := Get(client, fmt.Sprintf("/contest-type/%v", name), contestType)
	if err != nil {
		return err
	}
	return nil
}

func GetContestEffect(id int, contestEffect *ContestEffect, client *http.Client) error {
	err := Get(client, fmt.Sprintf("/contest-effect/%v", id), contestEffect)
	if err != nil {
		return err
	}
	return nil
}

func GetSuperContestEffect(id int, superContestEffect *SuperContestEffect, client *http.Client) error {
	err := Get(client, fmt.Sprintf("/super-contest-effect/%v", id), superContestEffect)
	if err != nil {
		return err
	}
	return nil
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
