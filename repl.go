package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"strings"
)

type cliCommand struct {
	cfg         *config
	name        string
	description string
	callback    func([]string) error
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
		command, parameters := ReadCommand(scanner, prompt)
		ExecuteCommand(command, parameters)
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	text = strings.TrimSpace(text)
	splitString := strings.Fields(text)
	return splitString
}

func ReadCommand(scanner *bufio.Scanner, prompt string) (command string, parameters []string) {

	fmt.Print(prompt)
	scanner.Scan()
	userInput := scanner.Text()
	userCleanInput := cleanInput(userInput)
	cleanInputLength := len(userCleanInput)

	if cleanInputLength < 1 {
		return "", parameters
	}

	command = userCleanInput[0]
	if cleanInputLength == 1 {
		return command, parameters
	}

	parameters = userCleanInput[1:]
	return command, parameters
}

func ExecuteCommand(command string, parameters []string) {
	if command == "" {
		return
	}
	_, exists := commands[command]
	if !exists {
		fmt.Println("Invalid command")
		return
	}

	err := commands[command].callback(parameters)
	if err != nil {
		fmt.Println(fmt.Errorf("Error Executing callback: %w", err))
	}
}

func commandExit(parameters []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Failed to exit")
}

func commandHelp(parameters []string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range commands {
		fmt.Print(command.formatInfo())
	}
	return nil
}

func commandMap(parameters []string) error {
	if commands["map"].cfg.next == "" {
		return nil
	}
	jsonData, err := pokeapi.PokeApiRequest(commands["map"].cfg.next)
	if err != nil {
		return err
	}
	locations, err := pokeapi.MarshalResults[pokeapi.Locations](jsonData)

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	commands["map"].cfg.next = locations.Next
	commands["map"].cfg.previous = locations.Previous
	return nil
}

func commandMapb(parameters []string) error {
	if commands["map"].cfg.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	jsonData, err := pokeapi.PokeApiRequest(commands["map"].cfg.previous)
	if err != nil {
		return err
	}
	locations, err := pokeapi.MarshalResults[pokeapi.Locations](jsonData)

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	commands["map"].cfg.next = locations.Next
	commands["map"].cfg.previous = locations.Previous
	return nil
}

func explore(parameters []string) error {
	if len(parameters) < 1 {
		return errors.New("Not enough arguments. Usage: explore <area-name>")
	}
	url := "https://pokeapi.co/api/v2/location-area/" + parameters[0]
	jsonData, err := pokeapi.PokeApiRequest(url)
	if string(jsonData) == "Not Found" {
		fmt.Println("Area not found")
		return nil
	}
	if err != nil {
		return err
	}
	location, err := pokeapi.MarshalResults[pokeapi.Location](jsonData)

	if len(location.PokemonEncounters) == 0 {
		fmt.Println("No pokemon in this area")
		return nil
	}
	for _, pokemon := range location.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
