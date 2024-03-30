package repository

import (
	"github.com/oreshkindev/golang-ca-boilerplate/internal/repository/postgres"
)

type Manager struct {
	Posts PostsRepository
}

func NewManager(postgres *postgres.Database) (*Manager, error) {
	return &Manager{
		Posts: *NewPostsRepository(postgres),
	}, nil
}
