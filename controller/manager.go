package controller

import (
	"github.com/user/repository/usecase"
)

type Manager struct {
	Posts PostsController
}

func NewManager(store *usecase.Manager) (*Manager, error) {
	return &Manager{
		Posts: *NewPostsController(&store.Posts),
	}, nil
}
