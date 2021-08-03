package main

import (
	"fmt"
	"os"
	"pikachu/config"
	"pikachu/repository"
	"pikachu/router"
	"pikachu/service"

	"github.com/labstack/echo"
)

func main() {
	pikachu := config.Pikachu
	e := echo.New()

	repo, err := repository.Init(pikachu)
	if err != nil {
		fmt.Printf("Error when Start repository: %v\n", err)
		os.Exit(1)
	}

	svc, err := service.Init(repo)
	if err != nil {
		fmt.Printf("Error when Start service: %v\n", err)
		os.Exit(1)
	}

	router.Init(e, repo, svc)

	e.Logger.Fatal(e.Start(":33333"))
}
