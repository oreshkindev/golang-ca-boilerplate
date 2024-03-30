package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/oreshkindev/golang-ca-boilerplate/internal/controller"
)

type Router struct {
	*chi.Mux
}

func NewRouter(controller *controller.Manager) (*Router, error) {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	/*
		SetContentType is a middleware function that sets the Content-Type header
		of the response to the specified content type.
	*/

	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/v1", func(r chi.Router) {
		r.Mount("/posts", postsHandler(controller))
	})

	return &Router{r}, nil
}

func postsHandler(controller *controller.Manager) http.Handler {
	r := chi.NewRouter()
	r.Get("/", controller.Posts.Get)
	r.Post("/", controller.Posts.Post)

	return r
}
