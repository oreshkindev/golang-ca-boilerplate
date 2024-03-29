package repository

import (
	"github.com/user/repository/repository/postgres"
)

type Manager struct {
	Posts PostsRepository
}

func NewManager(postgres *postgres.Database) (*Manager, error) {
	return &Manager{
		Posts: *NewPostsRepository(postgres),
	}, nil
}
