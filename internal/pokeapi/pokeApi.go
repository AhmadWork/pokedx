package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/AhmadWork/pokedx/internal/pokecache"
)

type LocationsResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type PokemonEncountersResponse struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type PokiClient struct {
    Cache pokecache.Cache
    httpClient http.Client
}

func NewClient(cacheInt time.Duration ) PokiClient {
    return PokiClient{
        Cache: pokecache.NewCache(cacheInt),
        httpClient: http.Client{
            Timeout: time.Minute*2,
        },
    }
}

func (p *PokiClient)  GetLocations(url string) (LocationsResponse, error) {
    loc := LocationsResponse{}
    cRes,ok := p.Cache.Get(url)
    if ok {
        error := json.Unmarshal(cRes, &loc)
        if error != nil {
		log.Fatal(error)        
        }
        return loc, nil
    }
	res, err := p.httpClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
        errorMessage := fmt.Sprintf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
        var zeroLocations LocationsResponse
        return zeroLocations, errors.New(errorMessage)
	}
	if err != nil {
		log.Fatal(err)
        errorMessage := fmt.Sprintf("Response failed with status code: %d and\n", res.StatusCode)
        var zeroLocations LocationsResponse
        return zeroLocations, errors.New(errorMessage)

	}
    err = json.Unmarshal(body, &loc)

    if err != nil {
        log.Fatal(err)
        errorMessage := "Data could not be converted somthing went wrong"
        var zeroLocations LocationsResponse
        return zeroLocations, errors.New(errorMessage)

    }
    p.Cache.Add(url, body)
    return loc, nil
}

func (p *PokiClient) GetExplore(url string) (PokemonEncountersResponse, error) {
    poke := PokemonEncountersResponse{}

    cRes,ok := p.Cache.Get(url)
    if ok {
        error := json.Unmarshal(cRes, &poke)
        if error != nil {
		log.Fatal(error)        
        }
        return poke, nil
    }

    res, err := p.httpClient.Get(url)

    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode > 404 {
        errorMessage := "No area under this name please try again with a correct area"
        var zeroPoke PokemonEncountersResponse
        return zeroPoke, errors.New(errorMessage)
    }

    body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
        errorMessage := fmt.Sprintf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
        var zeroPoke PokemonEncountersResponse
        return zeroPoke, errors.New(errorMessage)
	}


	if err != nil {
        errorMessage := fmt.Sprintf("Response failed with status code: %d and\n", res.StatusCode)
        var zeroPoke PokemonEncountersResponse
        return zeroPoke, errors.New(errorMessage)

	}

    err = json.Unmarshal(body, &poke)

    if err != nil {
        log.Fatal(err)
        errorMessage := "Data could not be converted somthing went wrong"
        var zeroPoke PokemonEncountersResponse
        return zeroPoke, errors.New(errorMessage)

    }
    p.Cache.Add(url, body)
    return poke, nil
}
