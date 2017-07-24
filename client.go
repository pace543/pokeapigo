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
