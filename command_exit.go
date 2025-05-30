package main

import (
	"fmt"
	"os"
)

func commandExit(c *config, args []string) error {
	fmt.Println("PokeDex exiting...")
	os.Exit(0)
	return nil
}
