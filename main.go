package main

import (
	"bufio"
	"os"
	"pokedexcli/repl"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	prompt := "Pokedex > "

	for {
		command := repl.ReadCommand(scanner, prompt)
		repl.ExecuteCommand(command)
	}
}
