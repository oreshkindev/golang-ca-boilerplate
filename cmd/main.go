package main

import (
	"context"
	"log"
	"net/http"

	"github.com/user/repository/conf"
	"github.com/user/repository/controller"
	"github.com/user/repository/router"
	"github.com/user/repository/service"
	"github.com/user/repository/store"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()

	config, err := conf.New()
	if err != nil {
		return err
	}

	store, err := store.New(ctx, &config.Postgres)
	if err != nil {
		return err
	}

	serviceManager, err := service.NewManager(ctx, store)
	if err != nil {
		return err
	}

	postsController := controller.NewPostsController(ctx, serviceManager)
	usersController := controller.NewUsersController(ctx, serviceManager)

	router, err := router.NewRouter(ctx)
	if err != nil {
		return err
	}

	// posts
	router.Get("/v1/posts", postsController.Get)
	router.Get("/v1/posts/{id}", postsController.GetBy)
	router.Post("/v1/posts", postsController.Post)
	router.Put("/v1/posts", postsController.Put)
	router.Delete("/v1/posts/{id}", postsController.Delete)

	// users
	router.Get("/v1/users", usersController.Get)
	router.Get("/v1/users/{id}", usersController.GetBy)
	router.Post("/v1/users", usersController.Post)
	router.Put("/v1/users", usersController.Put)
	router.Delete("/v1/users/{id}", usersController.Delete)

	return http.ListenAndServe(config.Port, router)
}
