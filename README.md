# pokeapigo
A concurrent Go wrapper with caching based on worker goroutines for pokeapi v2.

## Installation
`go get github.com/pace543/pokeapigo`

## Usage
pokeapigo is designed around the concept of Jobs and Results.

First, create a pointer for Client by passing a &ClientParam to pokeapiGo.NewClient():

    ```
    pokeapiClient := pokeapigo.NewClient(&pokeapigo.ClientParam{})
    ```

You can specify a specific *http.Client and the number of worker goroutines in &ClientParam (defaults are a *http.Client
with a timeout of 10 seconds and 3 worker goroutines).

Next, pass any number of &Job to pokeapiClient.AddJobs():

    ```
    pokeapiClient.AddJobs(&pokeapiClient.Job{
        Endpoint: // endpoint from pokeapi.co documentation
        Id:       // out of id and name, use this because some endpoints only accept id
        Name:
    }, ...)
    ```

Then, in a for-loop, obtain the results:

    ```
    for {
        res := pokeapiClient.PullResult()
    }
    ```

You will need to use a type assertion on res.Data to obtain a struct containing the information pulled for the given
endpoint and id/name:

    ```
    res.Data.(*pokeapigo.Berry) // for berry endpoint
    ```

If in a for-loop, you could use a type switch:

    ```
    for {
        switch res.Data.(type) {
        ...
        }
    }
    ```

## Example
```
import (
	"../pokeapigo"
	"fmt"
)

func main() {
	pokeapiClient := pokeapigo.NewClient(&pokeapigo.ClientParam{})
	pokeapiClient.AddJobs(&pokeapigo.Job{
		Endpoint: "berry-firmness",
		Id: 1,
	}, &pokeapigo.Job{
		Endpoint: "berry",
		Id: 2,
	}, &pokeapigo.Job{
		Endpoint: "berry-flavor",
		Id: 3,
	}, &pokeapigo.Job{
		Endpoint: "evolution-chain",
		Id: 4,
	}, &pokeapigo.Job{
		Endpoint: "evolution-trigger",
		Id: 1,
	})

	for i := 0; i < 5; i++ {
		res := pokeapiClient.PullResult()
		switch res.Data.(type) {
		case *pokeapigo.BerryFirmness:
			fmt.Println(res.Data.(*pokeapigo.BerryFirmness))
		case *pokeapigo.Berry:
			fmt.Println(res.Data.(*pokeapigo.Berry))
		case *pokeapigo.BerryFlavor:
			fmt.Println(res.Data.(*pokeapigo.BerryFlavor))
		case *pokeapigo.EvolutionChain:
			fmt.Println(res.Data.(*pokeapigo.EvolutionChain))
		case *pokeapigo.EvolutionTrigger:
			fmt.Println(res.Data.(*pokeapigo.EvolutionTrigger))
		}
	}

    // caching - doesn't make another HTTP request.
	pokeapiClient.AddJobs(&pokeapigo.Job{
		Endpoint: "evolution-trigger",
		Id: 1,
	})
	res := pokeapiClient.PullResult()
	fmt.Println(res.Data.(*pokeapigo.EvolutionTrigger))
}
```

## Built With
*   [go-cache](https://github.com/patrickmn/go-cache)

## License
This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for more information.