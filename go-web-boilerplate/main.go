package main

import (
	"os"
	"pikachu/repository"
	"pikachu/service"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	repo, err := repository.Init()
	if err != nil {
		os.Exit(1)
	}

	svc, err := service.Init(repo)
	if err != nil {
		os.Exit(1)
	}

	router.Init(e, repo, svc)

	e.Logger.Fatal(e.Start(":33333"))
}
