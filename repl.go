package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

type CliCommand struct {
    name string
    desc string
    callback func(*Config) error
    withParam bool
}

 var commands map[string]CliCommand= map[string]CliCommand{
     "exit":{
        name: "exit",
        desc: "to exit the program",
        callback: exitFunc,
    },
     "help":{
        name: "help",
        desc: "show the help menu of pokidex",
        callback: helpFunc,
    },
     "map":{
        name: "map",
        desc: "show the next 20 locations of the pokimon world",
        callback: mapFunc,
    },
     "mapb":{
        name: "mapb",
        desc: "show the previous 20 locations of the pokimon world",
        callback: mapbFunc,
    },
    "explore":{
        name: "explore",
        desc: "Explore the Pokemons of each locations in  the pokimon world",
        callback: exploreFunc,
        withParam: true,
    },
    "catch":{
        name: "catch",
        desc: "Gonna Catch'em all right?",
        callback: catchFunc,
        withParam: true,
    },
    "inspect":{
        name: "catch",
        desc: "Gonna Inspect'em all right?",
        callback: inspectFunc,
        withParam: true,
    },
    "pokedex":{
        name: "catch",
        desc: "Gonna Inspect'em all right?",
        callback: pokedexFunc,
    },

}


func startRepl(cfg *Config) {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("What pokimon you are looking for ?")
    for{
        fmt.Print(">")
         scanner.Scan()
        text := scanner.Text()
        cleaned := cleanInput(text)
        if len(cleaned) == 0 {
            continue
        }

        cReq := cleaned[0]
        command, ok := commands[cReq]

        if !ok {
            fmt.Println("Invalid command")
            fmt.Println("use help to find the available commands")
            continue
        }
        if command.withParam {
            if len(cleaned) < 2{
                fmt.Println("Command param is missing please add the area to explore it")
                continue
            }
            cfg.param = cleaned[1]
        }
        err := command.callback(cfg)
            if err != nil {
                log.Fatal(err)
        }
    }
}

func cleanInput (str string) []string {
    lowerd := strings.ToLower(str) 
    words := strings.Fields(lowerd)
    return words
}

func exitFunc(cfg *Config) error {
    os.Exit(0)
    return nil
}

func helpFunc(cfg *Config) error {
    fmt.Println("Welcome to the help menu of Pokidex")
    fmt.Println("the available commands are:")
    fmt.Println("- exit")
    fmt.Println("- help")
    fmt.Println("")
    return nil
}

func mapFunc(cfg *Config) error {
    res, err :=  cfg.api.GetLocations(cfg.next)
    
    if err != nil {
        return err
    }
    fmt.Println("Areas of the Poki worlds:")
    fmt.Println("-------------------------")
    for _, loc := range res.Results {
        fmt.Println(loc.Name)
    }
   mapPrev, ok := res.Previous.(string)
    if !ok {
        cfg.prev = "https://pokeapi.co/api/v2/location-area/"
    } else {
        cfg.prev = mapPrev
    }
    cfg.next = res.Next
    return nil
}

func mapbFunc(cfg *Config) error {
    if len(cfg.prev) == 0 {
        fmt.Println("you can't go back before starting")
        return nil
    }
    res, err :=  cfg.api.GetLocations(cfg.prev)
    
    if err != nil {
        return err
    }
    fmt.Println("Areas of the Poki worlds:")
    fmt.Println("-------------------------")
    for _, loc := range res.Results {
        fmt.Println(loc.Name)
    }
   mapPrev, ok := res.Previous.(string)
    if !ok {
        cfg.prev = "https://pokeapi.co/api/v2/location-area/"
    } else {
        cfg.prev = mapPrev
    }
    cfg.next = res.Next

    return nil
}

func exploreFunc(cfg *Config) error {
    url := "https://pokeapi.co/api/v2/location-area/" + cfg.param
    res, err :=  cfg.api.GetExplore(url)
    
    if err != nil {
        fmt.Println(err.Error())
        return nil
    }
    fmt.Printf("Pokemons encounters in %v : \n", cfg.param)
    fmt.Println("-------------------------")
    for _, poke := range res.PokemonEncounters {
        fmt.Println(poke.Pokemon.Name)
    }
    return nil
}

func catchFunc(cfg *Config) error {
    if _, ok := cfg.pokedex[cfg.param]; ok {
        fmt.Printf("you already have %v in your Pokedx \n", cfg.param)
        return nil
    }
    url := "https://pokeapi.co/api/v2/pokemon/" + cfg.param
    res, err :=  cfg.api.GetPokeData(url)
    
    if err != nil {
        fmt.Println(err.Error())
        return nil
    }
    fmt.Printf("Throwing a Pokeball at  %v : \n", cfg.param)
    fmt.Println("-------------------------")
    isCatched := catchPokemon(res.BaseExperience)
    if isCatched {
        fmt.Printf("%v has been catched \n", cfg.param)
        fmt.Println("Let's add to our PokeDex")
        cfg.pokedex[cfg.param] = res
    }else {
        fmt.Printf("%v has escaped \n", cfg.param)
        fmt.Println("Try again!")
    }
    return nil
}

func catchPokemon(rate int) bool {
	weight := 1.0 - float64(rate) / float64(690)
	return rand.Float64() < weight
}

func inspectFunc (cfg *Config) error {
    poke, ok := cfg.pokedex[cfg.param] 
    if !ok {
        fmt.Printf("You have  to catch %v first \n", cfg.param)
        return nil
    }

    fmt.Printf("Name: %v \n", cfg.param)
    fmt.Printf("Height: %v \n", poke.Height)
    fmt.Printf("Weight: %v \n", poke.Weight)
    fmt.Println("Stats:")
    for _,s := range poke.Stats {
        fmt.Printf("-%v: %v \n", s.Stat.Name, s.BaseStat)
    }
    fmt.Println("Types:")
    for _,t := range poke.Types {
        fmt.Printf("-%v\n", t.Type.Name)
    }

    return nil
}

func pokedexFunc (cfg *Config) error {

    fmt.Println("Your Pokedex:")
    for k := range cfg.pokedex {
        fmt.Printf("-%v \n", k)
    }
    return nil
}
