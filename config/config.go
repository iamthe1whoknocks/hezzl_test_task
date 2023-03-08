package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		PG   `yaml:"postgres"`
	}

	// app
	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}

	// http
	HTTP struct {
		Port string `yaml:"port"`
	}

	// log
	Log struct {
		Level string `yaml:"log_level"`
	}

	// postgres
	PG struct {
		PoolMax int    `yaml:"pool_max"`
		URL     string `yaml:"pg_url"`
	}
)

// get config from specified yml file
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("../../config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
