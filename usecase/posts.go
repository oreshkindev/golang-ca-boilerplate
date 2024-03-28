package usecase

import (
	"github.com/user/repository/entity"
)

type PostsUseCase struct {
	repository entity.PostsRepository
}

func NewPostsUseCase(r entity.PostsRepository) *PostsUseCase {
	return &PostsUseCase{
		repository: r,
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
