package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
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

func GetLocations(url string) (LocationsResponse, error) {
	res, err := http.Get(url)
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
    loc := LocationsResponse{}
    err = json.Unmarshal(body, &loc)

    if err != nil {
        log.Fatal(err)
        errorMessage := "Data could not be converted somthing went wrong"
        var zeroLocations LocationsResponse
        return zeroLocations, errors.New(errorMessage)

    }

    return loc, nil
}
