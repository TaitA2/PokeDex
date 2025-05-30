package main

import "fmt"

func commandInspect(c *config, args []string) error {
	if len(args) < 1 {
		fmt.Println("You must specify which pokemon to inspect.")
		return nil
	}
	name := args[0]
	pokemon, ok := pokeDex[name]
	if !ok {
		fmt.Println("You have not caught that pokemon.")
		return nil
	}
	fmt.Println("Name:", name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %v\n",stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %s\n",t.Type.Name)
	}
	return nil
}
