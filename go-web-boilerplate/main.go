package main

import (
	"fmt"
	"log"
	"os"
	"pikachu/config"
	"pikachu/repository"
	"pikachu/router"
	"pikachu/service"
	"pikachu/util"

	"github.com/dimiro1/banner"
	"github.com/labstack/echo"
)

var zlog *util.Logger

func init() {
	zlog, err := util.NewLogger()
	if err != nil {
		log.Fatalf("InitLog module[main] err[%s]", err.Error())
		os.Exit(1)
	}

	zlog.Infow("logger started")
	bannerInit()
}

func main() {
	pikachu := config.Pikachu
	e := echo.New()
	e.HideBanner = true

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

func bannerInit() {
	isEnabled := true
	isColorEnabled := true
	in, err := os.Open("banner.txt")
	if in == nil || err != nil {
		os.Exit(1)
	}

	banner.Init(os.Stdout, isEnabled, isColorEnabled, in)
}
