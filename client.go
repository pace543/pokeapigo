package pokeapiGo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type EndpointCall struct {
	Endpoint string
	Limit    int
	Offset   int
}

// Certain endpoints only use Id, so Id is preferred.
type Param struct {
	Id   int
	Name string
}

func (p *Param) String() string {
	if p == nil {
		return ""
	} else if p.Id != 0 {
		return fmt.Sprintf("%v", p.Id)
	} else if p.Name != "" {
		return fmt.Sprintf("%v", p.Name)
	} else {
		return ""
	}
}

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

func GetEndpointName(endpointCall *EndpointCall, client *http.Client, apiList *NamedAPIResourceList) error {
	err := GetWithUrl(client, fmt.Sprintf("https://pokeapi.co/api/v2/%v/?limit=%v&offset=%v", endpointCall.Endpoint,
		endpointCall.Limit, endpointCall.Offset), apiList)
	if err != nil {
		return err
	}
	return nil
}

func GetEndpointNoName(endpointCall *EndpointCall, client *http.Client, apiList *APIResourceList) error {
	err := GetWithUrl(client, fmt.Sprintf("https://pokeapi.co/api/v2/%v/?limit=%v&offset=%v", endpointCall.Endpoint,
		endpointCall.Limit, endpointCall.Offset), apiList)
	if err != nil {
		return err
	}
	return nil
}
