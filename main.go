package main

import (
	"time"

	"github.com/AhmadWork/pokedx/internal/pokeapi"
	"github.com/AhmadWork/pokedx/pkg/pokedx"
)



func main() {
    const cTime = time.Duration(5*time.Minute)

    cfg := pokedx.NewConfig(
         "https://pokeapi.co/api/v2/location-area/",
        pokeapi.NewClient(cTime),
        make(map[string]pokeapi.Pokemon),
    )


        pokedx.StartRepl(&cfg)
}
