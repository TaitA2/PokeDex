package main

import (
	"encoding/json"
	"fmt"
	"github.com/Taita2/PokeDex/internal/pokeapi"
)


func commandExplore(c *config, args []string) error {
	if len(args) < 1 {
		fmt.Println("You must specify which area to explore.")
		return nil
	}
	url := "https://pokeapi.co/api/v2/location-area/"+args[0]

	data, err := pokeapi.ApiHelper(url, &Cache)
	if err != nil {
		return err
	}

	var expRes exploreResponse

	if err := json.Unmarshal(data, &expRes); err != nil {
		return err
	}

	for _, p := range expRes.PokemonEncounters{
		fmt.Println(p.Pokemon.Name)
	}

	return nil
	

}

type exploreResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
