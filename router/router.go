package router

import (
	"context"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

type Router struct {
	*chi.Mux
}

func NewRouter(ctx context.Context) (*Router, error) {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Signature"},
		ExposedHeaders: []string{"Signature"},
	}))

	/*
		SetContentType is a middleware function that sets the Content-Type header
		of the response to the specified content type.
	*/

	r.Use(render.SetContentType(render.ContentTypeJSON))

	return &Router{r}, nil
}
