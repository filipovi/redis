package config

import (
	"encoding/json"
	"os"
)

// Config contains the information for Redis
type Config struct {
	Redis struct {
		URL      string `json:"url"`
		Password string `json:"password"`
		DB       int    `json:"db"`
	} `json:"redis"`
}

// New returns a Config struct filled with the json file
func New(path string) (Config, error) {
	var cfg Config

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return cfg, err
	}

	if err = json.NewDecoder(file).Decode(&cfg); err != nil {
		return cfg, err
	}

	return cfg, err
}
