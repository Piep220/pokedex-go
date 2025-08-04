package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	scan := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scan.Scan()
		input := scan.Text()
		cleaned := cleanInput(input)
		
		if len(cleaned) == 0 {
			fmt.Println("Please enter a command.")
			continue
		}

		found := false
		for name, cmd := range commands {
			if name == cleaned[0] {
				cmd.callback()
				found = true
				break
			}
		}

		if !found {
			fmt.Println("Unknown command")
		}


	}
}


func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)

	return words
}