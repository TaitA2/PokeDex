package main

import "fmt"

func commandHelp(c *config, args []string) error {
	fmt.Println("\nWelcome to the PokeDex!!")
	fmt.Print("Available commands are as follows:\n\n")
	for _, cmd := range supportedCommands{
		fmt.Printf("%s\t\t\t- Usage: %s\n", cmd.name,cmd.description)
		
	}
	fmt.Println()
	return nil
}
