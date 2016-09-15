package config

import (
	"encoding/json"
	"os"
	"time"
)

type Config struct {
	Username    string
	AccessToken string
	Exclude     []string
	Timeout     time.Duration
}

func GetConfig(configFile string) (*Config, error) {
	file, err := os.Open(configFile)
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
