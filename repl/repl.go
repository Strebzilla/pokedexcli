package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

func (c cliCommand) formatInfo() string {
	return fmt.Sprintf("%s: %s\n", c.name, c.description)
}

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	text = strings.TrimSpace(text)
	splitString := strings.Fields(text)
	return splitString
}

func ReadCommand(scanner *bufio.Scanner, prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	userInput := scanner.Text()
	userCleanInput := cleanInput(userInput)
	return userCleanInput[0]
}

func ExecuteCommand(command string) {
	commands[command].callback()
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Failed to exit")
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range commands {
		fmt.Print(command.formatInfo())
	}
	return nil
}
