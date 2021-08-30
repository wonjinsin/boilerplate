package router

import (
	ct "pikachu/controller"
	"pikachu/repository"
	"pikachu/service"

	"github.com/labstack/echo"
)

// Init ...
func Init(e *echo.Echo, repo *repository.Repository, svc *service.Service) {
	// Default Group
	api := e.Group("/api")
	ver := api.Group("/v1")

	// User Controller
	user := ver.Group("/user")
	userCt := ct.NewUserController(svc.User, repo.User)
	user.POST("", userCt.NewUser)
	user.GET("/:uid", userCt.GetUser)
}
