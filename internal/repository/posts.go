package repository

import (
	"github.com/oreshkindev/golang-ca-boilerplate/internal/entity"
	"github.com/oreshkindev/golang-ca-boilerplate/internal/repository/postgres"
)

type PostsRepository struct {
	database *postgres.Database
}

func NewPostsRepository(database *postgres.Database) *PostsRepository {
	return &PostsRepository{database: database}
}

func (repository *PostsRepository) Get() ([]entity.Posts, error) {
	entity := []entity.Posts{}

	if r := repository.database.Find(&entity); r.Error != nil {
		return nil, r.Error
	}

	return entity, nil
}

func (repository *PostsRepository) Post(entity *entity.Posts) (*entity.Posts, error) {
	if r := repository.database.Save(&entity); r.Error != nil {
		return nil, r.Error
	}
	return entity, nil
}
