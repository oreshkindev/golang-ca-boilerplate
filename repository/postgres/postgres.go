package postgres

import (
	"sync"

	"github.com/user/repository/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var (
	instance *Database
	once     sync.Once
)

func Dial(config config.Postgres) (*Database, error) {
	once.Do(func() {
		conn, err := gorm.Open(postgres.Open(config.Url))
		if err != nil {
			return
		}

		instance = &Database{conn}
	})

	if instance != nil {
		if err := Migrate(instance.DB); err != nil {
			return nil, err
		}
	}

	return instance, nil
}
