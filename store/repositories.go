package store

import (
	"context"

	"github.com/user/repository/model"
)

type PostsRepo interface {
	Get(context.Context) ([]model.Posts, error)
	GetBy(context.Context, string) (*model.Posts, error)
	Post(context.Context, *model.Posts) (*model.Posts, error)
	Put(context.Context, *model.Posts) (*model.Posts, error)
	Delete(context.Context, string) error
}

type UsersRepo interface {
	Get(context.Context) ([]model.Users, error)
	GetBy(context.Context, string) (*model.Users, error)
	Post(context.Context, *model.Users) (*model.Users, error)
	Put(context.Context, *model.Users) (*model.Users, error)
	Delete(context.Context, string) error
}
