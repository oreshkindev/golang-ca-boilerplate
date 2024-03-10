package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/user/repository/model"
	"github.com/user/repository/service"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type UsersController struct {
	ctx     context.Context
	service *service.Manager
}

func NewUsersController(ctx context.Context, service *service.Manager) *UsersController {
	return &UsersController{
		ctx:     ctx,
		service: service,
	}
}

func (ctr *UsersController) Get(w http.ResponseWriter, r *http.Request) {
	result, err := ctr.service.Users.Get(ctr.ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	render.JSON(w, r, result)
}

func (ctr *UsersController) GetBy(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	result, err := ctr.service.Users.GetBy(ctr.ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	render.JSON(w, r, result)
}

func (ctr *UsersController) Post(w http.ResponseWriter, r *http.Request) {
	tmp := model.Users{}

	if err := json.NewDecoder(r.Body).Decode(&tmp); err != nil {
		http.Error(w, "Получаем не то что ожидаем", http.StatusBadRequest)
		return
	}

	result, err := ctr.service.Users.Post(ctr.ctx, &tmp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	render.JSON(w, r, result)
}

func (ctr *UsersController) Put(w http.ResponseWriter, r *http.Request) {
	tmp := model.Users{}

	if err := json.NewDecoder(r.Body).Decode(&tmp); err != nil {
		http.Error(w, "Получаем не то что ожидаем", http.StatusBadRequest)
		return
	}

	result, err := ctr.service.Users.Put(ctr.ctx, &tmp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	render.JSON(w, r, result)
}

func (ctr *UsersController) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := ctr.service.Users.Delete(ctr.ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	render.JSON(w, r, nil)
}
