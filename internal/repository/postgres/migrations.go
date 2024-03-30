package postgres

import (
	"github.com/oreshkindev/golang-ca-boilerplate/internal/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	tables := []interface{}{
		&entity.Posts{},
	}

	if err := db.Migrator().DropTable(tables...); err != nil {
		return err
	}

	if err := db.AutoMigrate(tables...); err != nil {
		return err
	}

	return nil
}
