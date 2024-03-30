package usecase

import (
	"github.com/oreshkindev/golang-ca-boilerplate/internal/entity"
)

type PostsUseCase struct {
	repository entity.PostsRepository
}

func NewPostsUseCase(repository entity.PostsRepository) *PostsUseCase {
	return &PostsUseCase{
		repository: repository,
	}
}

func (usecase *PostsUseCase) Get() ([]entity.Posts, error) {
	result, err := usecase.repository.Get()
	if err != nil || result == nil {
		return nil, err
	}

	return result, nil
}

func (usecase *PostsUseCase) Post(entity *entity.Posts) (*entity.Posts, error) {
	result, err := usecase.repository.Post(entity)
	if err != nil {
		return nil, err
	}
	return result, nil
}
