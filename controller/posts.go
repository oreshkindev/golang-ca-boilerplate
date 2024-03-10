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

type PostsController struct {
	ctx     context.Context
	service *service.Manager
}

func NewPostsController(ctx context.Context, service *service.Manager) *PostsController {
	return &PostsController{
		ctx:     ctx,
		service: service,
	}
}

func (ctr *PostsController) Get(w http.ResponseWriter, r *http.Request) {
	result, err := ctr.service.Posts.Get(ctr.ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	render.JSON(w, r, result)
}

func (ctr *PostsController) GetBy(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	result, err := ctr.service.Posts.GetBy(ctr.ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	render.JSON(w, r, result)
}

func (ctr *PostsController) Post(w http.ResponseWriter, r *http.Request) {
	tmp := model.Posts{}

	if err := json.NewDecoder(r.Body).Decode(&tmp); err != nil {
		http.Error(w, "Получаем не то что ожидаем", http.StatusBadRequest)
		return
	}

	result, err := ctr.service.Posts.Post(ctr.ctx, &tmp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	render.JSON(w, r, result)
}

func (ctr *PostsController) Put(w http.ResponseWriter, r *http.Request) {
	tmp := model.Posts{}

	if err := json.NewDecoder(r.Body).Decode(&tmp); err != nil {
		http.Error(w, "Получаем не то что ожидаем", http.StatusBadRequest)
		return
	}

	result, err := ctr.service.Posts.Put(ctr.ctx, &tmp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	render.JSON(w, r, result)
}

func (ctr *PostsController) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := ctr.service.Posts.Delete(ctr.ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	render.JSON(w, r, nil)
}
