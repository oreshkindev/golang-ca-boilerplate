package conf

import (
	"log"
	"os"
	"sync"
)

type Conf struct {
	Host     string
	Port     string
	Source   string
	Postgres Postgres
}

type Postgres struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}

var (
	config Conf
	once   sync.Once
)

// Initializes a new Conf instance.
func New() (*Conf, error) {
	once.Do(func() {
		config = Conf{
			Host:   env("APP_HOST"),
			Port:   env("APP_PORT"),
			Source: env("APP_SOURCE"),
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
