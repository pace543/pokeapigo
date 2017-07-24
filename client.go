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

func GetWithUrl(client *http.Client, url string, target interface{}) error {
	result, err := client.Get(url)
	if err != nil {
		return err
	}
	defer result.Body.Close()
	return json.NewDecoder(result.Body).Decode(target)
}

func GetEndpointName(endpoint string, limit int, offset int, client *http.Client, apiList *NamedAPIResourceList) {
	err := GetWithUrl(client, fmt.Sprintf("https://pokeapi.co/api/v2/%v/?limit=%v&offset=%v", endpoint, limit, offset), apiList)
	if err != nil {
		return err
	}
	return nil
}

func GetEndpointNoName(endpoint string, limit int, offset int, client *http.Client, apiList *APIResourceList) {
	err := GetWithUrl(client, fmt.Sprintf("https://pokeapi.co/api/v2/%v/?limit=%v&offset=%v", endpoint, limit, offset), apiList)
	if err != nil {
		return err
	}
	return nil
}
