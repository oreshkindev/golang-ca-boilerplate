package usecase

import "github.com/user/repository/repository"

type Manager struct {
	Posts PostsUseCase
}

func NewManager(repository *repository.Manager) (*Manager, error) {
	return &Manager{
		Posts: *NewPostsUseCase(&repository.Posts),
	}, nil
}
