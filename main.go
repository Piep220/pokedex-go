package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex-go/internal/pokeapi"
	"strings"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeApiClient: pokeClient,
		areaMapNext: "default",
		areaMapPrevous: "default",
	}
	repl(cfg)
}

func repl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		input := reader.Text()
		cleaned := cleanInput(input)

		if len(cleaned) == 0 {
			fmt.Println("Please enter a command.")
			continue
		}

		cmd, exists := getCommands()[cleaned[0]]
		if exists {
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)

	return words
}