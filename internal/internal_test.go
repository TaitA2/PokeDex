package pokecache

import (
	"fmt"
	"testing"
	"time"
	"github.com/Taita2/PokeDex/internal/pokecache"
)

func TestAddGet(t *testing.T) {
	const delay = 7 * time.Second

	cases := []struct{
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://differenturl.com",
			val: []byte("more data"),
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := pokecache.NewCache(delay)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("Could not find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("Could not find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 7 * time.Millisecond
	const waitTime = baseTime + 7 * time.Millisecond
	cache := pokecache.NewCache(baseTime)
	url := "https://example.com"
	cache.Add(url, []byte("testdata"))

	_, ok := cache.Get(url)
	if !ok {
		t.Errorf("Could not get test key from cache")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get(url)
	if ok {
		t.Errorf("Got test key from cache when it should be reaped")
		return
	}
}
