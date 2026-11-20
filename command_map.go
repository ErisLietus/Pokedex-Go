package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locations struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func command_map(cfg *config, _args []string) error {
	if cfg.next == "" {
		fmt.Println("you're on the last page")
		return nil
	} else {
		return fetchPage(cfg.next, cfg)
	}
}

func command_mapb(cfg *config, _args []string) error {
	if cfg.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	} else {
		return fetchPage(cfg.previous, cfg)

	}

}

func fetchPage(next string, cfg *config) error {
	var location locations
	data, found := cfg.cache.Get(next)
	if found {
		if err := json.Unmarshal(data, &location); err != nil {
			return err
		}
	} else {
		res, err := http.Get(next)
		
		if res.Status != "200 OK" {
			return fmt.Errorf("Not found")
		}
		if err != nil {
			return err
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)

		if err != nil {
			return err
		}

		if err = json.Unmarshal(body, &location); err != nil {
			return err
		}
		cfg.cache.Add(next, body)
	}

	if location.Next == nil {
		cfg.next = ""
	} else {
		cfg.next = *location.Next
	}
	if location.Previous == nil {
		cfg.previous = ""
	} else {
		cfg.previous = *location.Previous
	}
	for _, area := range location.Results {
		fmt.Println(area.Name)
	}
	return nil
}
