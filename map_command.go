package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func commandMap(c *config, args ...string) error {
	url := c.Next

	if !strings.Contains(url, "location-area"){
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	return mapHelper(url, c)
}

func commandMapb(c *config, args ...string) error {
	url := c.Previous

	if !strings.Contains(url, "location-area"){
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	return mapHelper(url, c)
}

func mapHelper(url string, c *config) error {

	data, ok := Cache.Get(url)

	if !ok {
		res, err := http.Get(url)
			if err != nil {
				return err
		}
		defer res.Body.Close()
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
	}

	Cache.Add(url, data)

	if err := json.Unmarshal(data, &c); err != nil {
		return err
	}

	for _, l := range c.Results {
		fmt.Println(l.Name)
	}

	return nil
	
}
