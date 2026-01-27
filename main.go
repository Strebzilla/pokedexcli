package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		userCleanInput := cleanInput(userInput)
		fmt.Println(fmt.Sprintf("Your command was: %s", userCleanInput[0]))
	}
}
