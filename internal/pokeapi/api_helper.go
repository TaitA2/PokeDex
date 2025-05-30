package pokeapi

import (
	"io"
	"net/http"

	"github.com/Taita2/PokeDex/internal/pokecache"
)

func ApiHelper(url string, Cache *pokecache.Cache) ([]byte, error) {
	data, ok := Cache.Get(url)

	if !ok {
		res, err := http.Get(url)
			if err != nil {
				return nil, err
		}
		defer res.Body.Close()
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
	}

	Cache.Add(url, data)
	return data, nil
}
