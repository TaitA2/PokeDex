package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


var supportedCommands = make(map[string]cliCommand)

func main() {
	createSupportedCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("PokeDex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleanUserInput := cleanInput(userInput)
		userCommand := cleanUserInput[0]

		cmd, ok := supportedCommands[userCommand]
		if ok {
			cmd.callback()
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
	callback func() error
}

func createSupportedCommands() {

	supportedCommands["exit"] = cliCommand{
		name: "exit",
		description: "Exits the PokeDex",
		callback: commandExit,
		}

	supportedCommands["help"] = cliCommand{
		name: "help",
		description: "Display this help message",
		callback: commandHelp,
		}

}

func commandExit() error {
	fmt.Println("PokeDex exiting...")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("\nWelcome to the PokeDex!!")
	fmt.Print("Available commands are as follows:\n\n")
	for _, cmd := range supportedCommands{
		fmt.Printf("%s \t\t- Usage: %s\n", cmd.name,cmd.description)
		
	}
	fmt.Println()
	return nil
}


