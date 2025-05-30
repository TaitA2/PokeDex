package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"github.com/Taita2/PokeDex/internal/pokecache"
)


var supportedCommands = make(map[string]cliCommand)

var pokeDex = make(map[string]Pokemon)

var Cache = pokecache.NewCache(7*time.Second)

func main() {
	createSupportedCommands()
	config := config{}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nPokeDex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleanUserInput := cleanInput(userInput)
		userCommand := cleanUserInput[0]

		cmd, ok := supportedCommands[userCommand]
		if ok {
			cmd.callback(&config, cleanUserInput[1:])
		} else {
			fmt.Printf("Unknown command: '%s', use 'help' for a list of available commands.\n", userCommand)
		}
	}
}

type cliCommand struct {
	name		string
	description	string
	callback func(*config, []string) error
}

func createSupportedCommands() {

	supportedCommands["exit"] = cliCommand{
		name: "exit",
		description: "Exit the PokeDex.",
		callback: commandExit,
		}

	supportedCommands["help"] = cliCommand{
		name: "help",
		description: "Display this help message.",
		callback: commandHelp,
		}

	supportedCommands["map"] = cliCommand{
		name: "map",
		description: "List the NEXT 20 locations in the Pokemon world.",
		callback: commandMap,
	}

	supportedCommands["mapb"] = cliCommand{
		name: "mapb",
		description: "List the PREVIOUS 20 locations in the Pokemon world.",
		callback: commandMapb,
	}

	supportedCommands["explore"] = cliCommand{
		name: "explore",
		description: "List all pokemon in the specified location.",
		callback: commandExplore,
	}

	supportedCommands["catch"] = cliCommand{
		name: "catch",
		description: "Attempt to catch specified Pokemon.",
		callback: commandCatch,
	}

}

type config struct{
	Count		int		`json:"count"`
	Next		string	`json:"next"`
	Previous	string	`json:"previous"`
	Results		[]result`json:"results"`
}

type result struct{
	Name	string	`json:"name"`
	Url		string	`json:"url"`
}

