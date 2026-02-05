// Package pokeapi implements functions for interactions with [Poke API]
//
// [Poke API]: https://pokeapi.co/
package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"pokedexcli/internal/pokecache"
	"time"
)

var cache *pokecache.Cache

func init() {
	cache = pokecache.NewCache(5 * time.Second)
}

func PokeApiRequest(url string) ([]byte, error) {
	jsonData, exists := cache.Get(url)
	if exists {
		return jsonData, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	jsonData, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	cache.Add(url, jsonData)
	return jsonData, nil
}

func MarshalResults[T any](jsonData []byte) (T, error) {
	var results T

	if err := json.Unmarshal(jsonData, &results); err != nil {
		return results, err
	}
	return results, nil
}
