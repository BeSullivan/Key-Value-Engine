package server

import (
	"Kiwi/controllerlayer"

	"github.com/labstack/echo/v4"
)

func userRoutes(e *echo.Echo, handler *controllerlayer.UserControllerImpl) {

	e.POST("/users/login", handler.LoginHandler)
	e.POST("/users/register", handler.RegisterHandler)
}
