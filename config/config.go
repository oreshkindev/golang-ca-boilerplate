package config

import (
	"log"
	"os"
	"sync"
)

type Config struct {
	Host, Port string
	Postgres   Postgres
}

type Postgres struct {
	User, Pass, Host, Port, Name string
}

var (
	config Config
	once   sync.Once
)

// Initializes a new Config instance.
func Load() (*Config, error) {
	once.Do(func() {
		config = Config{
			Host: env("APP_HOST"),
			Port: env("APP_PORT"),
			Postgres: Postgres{
				User: env("DB_USER"),
				Pass: env("DB_PASS"),
				Host: env("DB_HOST"),
				Port: env("DB_PORT"),
				Name: env("DB_NAME"),
			},
		}
	})
	return &config, nil
}

func env(key string) string {
	value, ok := os.LookupEnv(key)
	// If the variable was not set, return error.
	if !ok {
		log.Fatal("Could not load environment variable")
	}
	// Return the value of the variable.
	return value
}
