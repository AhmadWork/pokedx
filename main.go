package main

import (
	"time"

	"github.com/AhmadWork/pokedx/internal/pokeapi"
)

type Config struct {
    next string
    prev string
    api pokeapi.PokiClient
    param string
    pokedex map[string]pokeapi.Pokemon
}


func main() {
    const cTime = time.Duration(5*time.Minute)

    cfg := Config{
        next: "https://pokeapi.co/api/v2/location-area/",
        api: pokeapi.NewClient(cTime),
        pokedex: make(map[string]pokeapi.Pokemon),
    }

        startRepl(&cfg)
}
