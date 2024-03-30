package main

import (
	"log"
	"net/http"

	"github.com/oreshkindev/golang-ca-boilerplate/config"
	"github.com/oreshkindev/golang-ca-boilerplate/internal/controller"
	"github.com/oreshkindev/golang-ca-boilerplate/internal/controller/router"
	"github.com/oreshkindev/golang-ca-boilerplate/internal/repository"
	"github.com/oreshkindev/golang-ca-boilerplate/internal/repository/postgres"
	"github.com/oreshkindev/golang-ca-boilerplate/internal/usecase"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	config, err := config.Load()
	if err != nil {
		return err
	}

	postgres, err := postgres.Dial(config.Postgres)
	if err != nil {
		return err
	}

	repository, err := repository.NewManager(postgres)
	if err != nil {
		return err
	}

	usecase, err := usecase.NewManager(repository)
	if err != nil {
		return err
	}

	controller, err := controller.NewManager(usecase)
	if err != nil {
		return err
	}

	router, err := router.NewRouter(controller)
	if err != nil {
		return err
	}

	return http.ListenAndServe(config.Port, router)
}
