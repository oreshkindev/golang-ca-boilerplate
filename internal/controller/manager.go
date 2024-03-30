package controller

import (
	"github.com/oreshkindev/golang-ca-boilerplate/internal/usecase"
)

type Manager struct {
	Posts PostsController
}

func NewManager(usecase *usecase.Manager) (*Manager, error) {
	return &Manager{
		Posts: *NewPostsController(&usecase.Posts),
	}, nil
}
