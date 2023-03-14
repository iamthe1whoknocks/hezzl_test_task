package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App        `yaml:"app"`
		HTTP       `yaml:"http"`
		Log        `yaml:"logger"`
		PG         `yaml:"postgres"`
		ClickHouse `yaml:"clickhouse"`
		Redis      `yaml:"redis"`
		Nats       `yaml:"nats"`
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
		URL     string `yaml:"pg_url" env:"PG_URL"`
	}
	// clickhouse
	ClickHouse struct {
		Host     string `yaml:"ch_host"`
		Port     string `yaml:"ch_port"`
		DbName   string `yaml:"db_name"`
		Username string `yaml:"username"`
		Password string `yaml:"password" env:"CLICKHOUSE_PASSWORD"`
		Engine   string `yaml:"engine"`
	}

	// redis
	Redis struct {
		Host     string `yaml:"redis_host"`
		Port     string `yaml:"redis_port"`
		Password string `yaml:"redis_password" env:"REDIS_PASSWORD"`
		DB       int    `yaml:"redis_db"`
		TTL      int    `yaml:"redis_ttl"`
	}

	// nats
	Nats struct {
		Host       string `yaml:"host"`
		Port       string `yaml:"port"`
		Topic      string `yaml:"topic"`
		BatchCount int    `yaml:"batch_count"`
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
