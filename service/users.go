package service

import (
	"context"

	"github.com/user/repository/model"
	"github.com/user/repository/store"
)

type UsersWebService struct {
	ctx   context.Context
	store *store.Store
}

func NewUsersWebService(ctx context.Context, store *store.Store) *UsersWebService {
	return &UsersWebService{
		ctx:   ctx,
		store: store,
	}
}

func (svc *UsersWebService) Get(ctx context.Context) ([]model.Users, error) {
	UsersDB, err := svc.store.Users.Get(ctx)
	if err != nil {
		return nil, err
	}
	if UsersDB == nil {
		return nil, err
	}

	return UsersDB, nil
}

func (svc *UsersWebService) GetBy(ctx context.Context, id string) (*model.Users, error) {
	result, err := svc.store.Users.GetBy(ctx, id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, err
	}

	return result, nil
}

func (svc *UsersWebService) Post(ctx context.Context, mdl *model.Users) (*model.Users, error) {
	result, err := svc.store.Users.Post(ctx, mdl)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, err
	}

	return result, nil
}

func (svc *UsersWebService) Put(ctx context.Context, mdl *model.Users) (*model.Users, error) {
	result, err := svc.store.Users.Put(ctx, mdl)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, err
	}

	return result, nil
}

func (svc *UsersWebService) Delete(ctx context.Context, id string) error {

	err := svc.store.Users.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
