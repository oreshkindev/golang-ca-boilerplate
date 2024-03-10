package store

import (
	"context"

	"github.com/user/repository/conf"
	"github.com/user/repository/store/postgres"
)

type Store struct {
	db    *postgres.DB
	Posts PostsRepo
	Users UsersRepo
}

func New(ctx context.Context, conf *conf.Postgres) (*Store, error) {
	db, err := postgres.Dial(conf)
	if err != nil {
		return nil, err
	}

	if db != nil {
		if err = Migrate(db); err != nil {
			return nil, err
		}
	}

	return &Store{
		db:    db,
		Posts: postgres.NewPostsRepo(db),
		Users: postgres.NewUsersRepo(db),
	}, nil
}
