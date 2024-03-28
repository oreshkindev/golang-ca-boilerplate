package usecase

import "github.com/user/repository/repository"

type Manager struct {
	Posts PostsUseCase
}

func NewManager(store *repository.Store) (*Manager, error) {
	return &Manager{
		Posts: *NewPostsUseCase(&store.Posts),
	}, nil
}
