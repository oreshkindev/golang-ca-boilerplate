package service

import (
	"context"
	"errors"

	"github.com/user/repository/store"
)

type Manager struct {
	Posts PostsService
	Users UsersService
}

func NewManager(ctx context.Context, store *store.Store) (*Manager, error) {
	if store == nil {
		return nil, errors.New("NO STORE PROVIDED")
	}
	return &Manager{
		Posts: NewPostsWebService(ctx, store),
		Users: NewUsersWebService(ctx, store),
	}, nil
}
