package repository

import (
	"github.com/user/repository/repository/postgres"
)

type Store struct {
	db    *postgres.Database
	Posts PostsRepository
}

func NewManager(db *postgres.Database) (*Store, error) {
	return &Store{
		db:    db,
		Posts: *NewPostsRepository(db),
	}, nil
}
