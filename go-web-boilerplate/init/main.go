package main

import (
	"pikachu/controller"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	controller.Controller(e)
	e.Logger.Fatal(e.Start(":33333"))
}
