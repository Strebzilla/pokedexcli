package pokeapi

import (
	"encoding/json"
	"net/http"
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

func GetPokeJson(url string) (Results, error) {
	res, err := http.Get(url)
	if err != nil {
		return Results{}, err
	}
	defer res.Body.Close()

	var responseJson Results
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&responseJson); err != nil {
		return Results{}, err
	}
	return responseJson, nil
}
