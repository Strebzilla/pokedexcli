/*
pokedexcli retrieves data from Poke API
*/
package main

var commands map[string]cliCommand

func main() {
	prompt := "Pokedex > "
	mapConfig := config{
		next:     "https://pokeapi.co/api/v2/location-area/",
		previous: "",
	}

	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
			cfg:         nil,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
			cfg:         nil,
		},
		"map": {
			name:        "map",
			description: "Displays the next set of 20 locations",
			callback:    commandMap,
			cfg:         &mapConfig,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous set of 20 locations",
			callback:    commandMapb,
			cfg:         &mapConfig,
		},
		"explore": {
			name:        "explore",
			description: "Get the list of found pokemon in an area. Takes an area name as an argument",
			callback:    explore,
			cfg:         nil,
		},
		"catch": {
			name:        "catch",
			description: "Catches a pokemon. Takes a pokemon name as an argument",
			callback:    catch,
			cfg:         nil,
		},
		"inspect": {
			name:        "inspect",
			description: "Prints the name, weight, stat, and type data about a Pokemon stored in the pokedex. Takes a pokemon name as an argument.",
			callback:    inspect,
			cfg:         nil,
		},
	}

	startRepl(prompt)
}
