package controller

import (
	"github.com/user/repository/usecase"
)

type Manager struct {
	Posts PostsController
}

func NewManager(usecase *usecase.Manager) (*Manager, error) {
	return &Manager{
		Posts: *NewPostsController(&usecase.Posts),
	}, nil
}
