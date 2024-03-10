package service

import (
	"context"

	"github.com/user/repository/model"
	"github.com/user/repository/store"
)

type PostsWebService struct {
	ctx   context.Context
	store *store.Store
}

func NewPostsWebService(ctx context.Context, store *store.Store) *PostsWebService {
	return &PostsWebService{
		ctx:   ctx,
		store: store,
	}
}

func (svc *PostsWebService) Get(ctx context.Context) ([]model.Posts, error) {
	postsDB, err := svc.store.Posts.Get(ctx)
	if err != nil {
		return nil, err
	}
	if postsDB == nil {
		return nil, err
	}

	return postsDB, nil
}

func (svc *PostsWebService) GetBy(ctx context.Context, id string) (*model.Posts, error) {
	result, err := svc.store.Posts.GetBy(ctx, id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, err
	}

	return result, nil
}

func (svc *PostsWebService) Post(ctx context.Context, mdl *model.Posts) (*model.Posts, error) {
	result, err := svc.store.Posts.Post(ctx, mdl)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, err
	}

	return result, nil
}

func (svc *PostsWebService) Put(ctx context.Context, mdl *model.Posts) (*model.Posts, error) {
	result, err := svc.store.Posts.Put(ctx, mdl)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, err
	}

	return result, nil
}

func (svc *PostsWebService) Delete(ctx context.Context, id string) error {

	err := svc.store.Posts.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
