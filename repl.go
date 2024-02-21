package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

        command.callback()
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
