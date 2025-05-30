package main

import "fmt"

func commandPokedex(c *config, args[]string) error {
	fmt.Println("Your PokeDex:")
	for k,_ := range pokeDex {
		fmt.Println(" -", k)
	}
	return nil
}
