package main

import(
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
} 

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for name, cmd := range commands {
		fmt.Printf("%s: %s\n", name, cmd.description)
	}
	return nil
}

var commands map[string]cliCommand

func init() {
    commands = map[string]cliCommand{
        "help": {
            name:        "help",
            description: "Displays a help message", 
            callback:    commandHelp,
        },
        "exit": {
            name:        "exit",
            description: "Exit the Pokedex",
            callback:    commandExit,
        },
    }
}