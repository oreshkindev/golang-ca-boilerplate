package postgres

import (
	"context"

	"github.com/user/repository/model"
	"gorm.io/gorm"
)

type PostsRepo struct {
	db *DB
}

func NewPostsRepo(db *DB) *PostsRepo {
	return &PostsRepo{db: db}
}

func (repo *PostsRepo) Get(ctx context.Context) ([]model.Posts, error) {
	mdl := []model.Posts{}

	if r := repo.db.Find(&mdl); r.Error != nil {
		return nil, r.Error
	}

	return mdl, nil
}

func (repo *PostsRepo) GetBy(ctx context.Context, id string) (*model.Posts, error) {
	mdl := &model.Posts{}

	if r := repo.db.Find(&mdl, "id = ?", id); r.Error != nil {
		return nil, r.Error
	}

	return mdl, nil
}

func (repo *PostsRepo) Post(ctx context.Context, tmp *model.Posts) (*model.Posts, error) {
	if r := repo.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&tmp); r.Error != nil {
		return nil, r.Error
	}
	return tmp, nil
}

func (repo *PostsRepo) Put(ctx context.Context, tmp *model.Posts) (*model.Posts, error) {

	if r := repo.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&tmp); r.Error != nil {
		return nil, r.Error
	}
	return tmp, nil
}

func (repo *PostsRepo) Delete(ctx context.Context, id string) error {
	if r := repo.db.Delete(&model.Posts{}, id); r.Error != nil {
		return r.Error
	}
	return nil
}
