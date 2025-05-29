package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)


var supportedCommands = make(map[string]cliCommand)

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
			cmd.callback(&config)
		} else {
			fmt.Printf("Unknown command: '%s', use 'help' for a list of available commands.\n", userCommand)
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Split(text, " ")
	var cleanWords []string
	for _, word := range words {
		if word != "" {
			cleanWords = append(cleanWords, strings.ToLower(word))
	}
}
	return cleanWords
}

type cliCommand struct {
	name		string
	description	string
	callback func(*config) error
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

func commandExit(c *config) error {
	fmt.Println("PokeDex exiting...")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println("\nWelcome to the PokeDex!!")
	fmt.Print("Available commands are as follows:\n\n")
	for _, cmd := range supportedCommands{
		fmt.Printf("%s \t\t- Usage: %s\n", cmd.name,cmd.description)
		
	}
	fmt.Println()
	return nil
}



func commandMap(c *config) error {
	url := c.Next
	mapHelper(url, c)
	return nil
}

func commandMapb(c *config) error {
	url := c.Previous
	mapHelper(url, c)
	return nil
}


func mapHelper(url string, c *config) error {
	if !strings.Contains(url, "location-area"){
		url = "https://pokeapi.co/api/v2/location-area/"
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&c); err != nil {
		return err
	}

	for _, l := range c.Results {
		fmt.Println(l.Name)
	}

	return nil
	
}
