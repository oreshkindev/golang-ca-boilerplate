package controller

import (
	"encoding/json"
	"net/http"

	"github.com/user/repository/entity"

	"github.com/go-chi/render"
)

type PostsController struct {
	usecase entity.PostsUseCase
}

func NewPostsController(usecase entity.PostsUseCase) *PostsController {
	return &PostsController{
		usecase: usecase,
	}
}

func (controller *PostsController) Get(w http.ResponseWriter, r *http.Request) {
	result, err := controller.usecase.Get()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	render.JSON(w, r, result)
}

func (controller *PostsController) Post(w http.ResponseWriter, r *http.Request) {
	entity := entity.Posts{}

	err := json.NewDecoder(r.Body).Decode(&entity)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	result, err := controller.usecase.Post(&entity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	render.JSON(w, r, result)
}
