package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Mode string `default:"production"`
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load env: %w", err)
	}

	var config Config
	err = envconfig.Process("myapp", &config)
	if err != nil {
		return nil, fmt.Errorf("failed to process config: %w", err)
	}
	return &config, nil
}
