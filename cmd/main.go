package main

import (
	"log"
	"net/http"

	"github.com/user/repository/config"
	"github.com/user/repository/controller"
	"github.com/user/repository/repository"
	"github.com/user/repository/repository/postgres"
	"github.com/user/repository/router"
	"github.com/user/repository/usecase"
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
