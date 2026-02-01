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

type Results struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var cache *pokecache.Cache

func init() {
	cache = pokecache.NewCache(5 * time.Second)
}

func GetPokeJson(url string) (Results, error) {
	var responseJson Results

	jsonData, exists := cache.Get(url)
	if exists {
		if err := json.Unmarshal(jsonData, &responseJson); err != nil {
			return Results{}, err
		}
		return responseJson, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return Results{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Results{}, err
	}
	cache.Add(url, body)
	if err := json.Unmarshal(body, &responseJson); err != nil {
		return Results{}, err
	}
	return responseJson, nil
}
