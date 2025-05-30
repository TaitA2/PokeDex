package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"github.com/Taita2/PokeDex/internal/pokeapi"
)

func commandMap(c *config, args []string) error {
	url := c.Next

	if !strings.Contains(url, "location-area"){
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	return mapHelper(url, c)
}

func commandMapb(c *config, args []string) error {
	url := c.Previous

	if !strings.Contains(url, "location-area"){
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	return mapHelper(url, c)
}

func mapHelper(url string, c *config) error {

	data, err := pokeapi.ApiHelper(url, &Cache)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &c); err != nil {
		return err
	}

	for _, l := range c.Results {
		fmt.Println(l.Name)
	}

	return nil
	
}
