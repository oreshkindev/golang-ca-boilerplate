package postgres

import (
	"github.com/user/repository/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func Dial(config config.Postgres) (*Database, error) {

	conn, err := gorm.Open(postgres.Open("postgres://oresh:postgres@localhost:5432/postgres"))
	if err != nil {
		return nil, err
	}
	if conn != nil {
		if err = Migrate(conn); err != nil {
			return nil, err
		}
	}

	return &Database{conn}, nil
}
