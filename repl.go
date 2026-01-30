package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"strings"
)

type cliCommand struct {
	cfg         *config
	name        string
	description string
	callback    func() error
}

type config struct {
	next     string
	previous string
}

func (c cliCommand) formatInfo() string {
	return fmt.Sprintf("%s: %s\n", c.name, c.description)
}

func startRepl(prompt string) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		command := ReadCommand(scanner, prompt)
		ExecuteCommand(command)
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
	err := commands[command].callback()
	if err != nil {
		fmt.Println(fmt.Errorf("Error Executing callback: %w", err))
	}
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

func commandMap() error {
	if commands["map"].cfg.next == "" {
		return nil
	}
	json, err := pokeapi.GetPokeJson(commands["map"].cfg.next)
	if err != nil {
		return err
	}
	for _, location := range json.Results {
		fmt.Println(location.Name)
	}
	commands["map"].cfg.next = json.Next
	commands["map"].cfg.previous = json.Previous
	return nil
}

func commandMapb() error {
	if commands["map"].cfg.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	json, err := pokeapi.GetPokeJson(commands["map"].cfg.previous)
	if err != nil {
		return err
	}
	for _, location := range json.Results {
		fmt.Println(location.Name)
	}
	commands["map"].cfg.next = json.Next
	commands["map"].cfg.previous = json.Previous
	return nil
}