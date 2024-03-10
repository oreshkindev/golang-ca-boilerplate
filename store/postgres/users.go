package postgres

import (
	"context"

	"github.com/user/repository/model"
	"gorm.io/gorm"
)

type UsersRepo struct {
	db *DB
}

func NewUsersRepo(db *DB) *UsersRepo {
	return &UsersRepo{db: db}
}

func (repo *UsersRepo) Get(ctx context.Context) ([]model.Users, error) {
	mdl := []model.Users{}

	if r := repo.db.Find(&mdl); r.Error != nil {
		return nil, r.Error
	}

	return mdl, nil
}

func (repo *UsersRepo) GetBy(ctx context.Context, id string) (*model.Users, error) {
	mdl := &model.Users{}

	if r := repo.db.Find(&mdl, "id = ?", id); r.Error != nil {
		return nil, r.Error
	}

	return mdl, nil
}

func (repo *UsersRepo) Post(ctx context.Context, tmp *model.Users) (*model.Users, error) {
	if r := repo.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&tmp); r.Error != nil {
		return nil, r.Error
	}
	return tmp, nil
}

func (repo *UsersRepo) Put(ctx context.Context, tmp *model.Users) (*model.Users, error) {

	if r := repo.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&tmp); r.Error != nil {
		return nil, r.Error
	}
	return tmp, nil
}

func (repo *UsersRepo) Delete(ctx context.Context, id string) error {
	if r := repo.db.Delete(&model.Users{}, id); r.Error != nil {
		return r.Error
	}
	return nil
}
