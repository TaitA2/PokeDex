package main

import "strings"

func commandExplore(c *config, args ...string) error {
	url := c.Previous
	
	if !strings.Contains(url, "location-area"){
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	url += args[0]

	return mapHelper(url, c)

}
