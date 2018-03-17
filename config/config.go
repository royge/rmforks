package config

import (
	"encoding/json"
	"os"
	"time"
)

// Config type
type Config struct {
	Username    string        `json:"username"`
	AccessToken string        `json:"accesstoken"`
	Exclude     []string      `json:"exclude"`
	Timeout     time.Duration `json:"timeout"`
}

// GetConfig read and decodes Config for a file.
func GetConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(file)

	cfg := &Config{}
	err = decoder.Decode(cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
