package repository

import (
	"github.com/user/repository/entity"
	"github.com/user/repository/repository/postgres"
)

type PostsRepository struct {
	database *postgres.Database
}

func NewPostsRepository(database *postgres.Database) *PostsRepository {
	return &PostsRepository{database: database}
}

func (repository *PostsRepository) Get() ([]entity.Posts, error) {
	mdl := []entity.Posts{}

	if r := repository.database.Find(&mdl); r.Error != nil {
		return nil, r.Error
	}

	return mdl, nil
}

func (repository *PostsRepository) Post(tmp *entity.Posts) (*entity.Posts, error) {
	if r := repository.database.Save(&tmp); r.Error != nil {
		return nil, r.Error
	}
	return tmp, nil
}
