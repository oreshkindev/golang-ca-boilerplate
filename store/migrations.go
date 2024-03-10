package store

import (
	"github.com/user/repository/model"
	"github.com/user/repository/store/postgres"
)

func Migrate(db *postgres.DB) error {
	tables := []interface{}{
		&model.Posts{},
		&model.Users{},
	}

	if err := db.Migrator().DropTable(tables...); err != nil {
		return err
	}

	if err := db.AutoMigrate(tables...); err != nil {
		return err
	}

	return nil
}
