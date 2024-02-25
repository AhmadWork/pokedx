package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	pokiapi "github.com/AhmadWork/pokedx/internal/pokeApi"
)

type CliCommand struct {
    name string
    desc string
    callback func() error
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
}

type Map struct {
    next string
    prev string
}

var mapStatus Map = Map{
    next: "https://pokeapi.co/api/v2/location-area/",
}

func startRepl() {
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

        err := command.callback()
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

func exitFunc() error {
    os.Exit(0)
    return nil
}

func helpFunc() error {
    fmt.Println("Welcome to the help menu of Pokidex")
    fmt.Println("the available commands are:")
    fmt.Println("- exit")
    fmt.Println("- help")
    fmt.Println("")
    return nil
}

func mapFunc() error {
    res, err :=  pokiapi.GetLocations(mapStatus.next)
    
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
        mapStatus.prev = "https://pokeapi.co/api/v2/location-area/"
    } else {
        mapStatus.prev = mapPrev
    }
    mapStatus.next = res.Next
    return nil
}

func mapbFunc() error {
    res, err :=  pokiapi.GetLocations(mapStatus.prev)
    
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
        mapStatus.prev = "https://pokeapi.co/api/v2/location-area/"
    } else {
        mapStatus.prev = mapPrev
    }
    mapStatus.next = res.Next

    return nil
}

