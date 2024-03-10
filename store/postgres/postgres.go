package postgres

import (
	"fmt"

	"github.com/user/repository/conf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func Dial(c *conf.Postgres) (*DB, error) {

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", c.Host, c.Port, c.User, c.Name, c.Pass)

	conn, err := gorm.Open(postgres.Open(connStr))
	if err != nil {
		return nil, err
	}

	return &DB{conn}, nil
}
