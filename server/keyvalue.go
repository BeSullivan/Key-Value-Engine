package server

import (
	"Kiwi/controllerlayer"

	"github.com/labstack/echo/v4"
)

func kiwiRoutes(e *echo.Echo, handler *controllerlayer.KiwiControllerImpl) {
	e.POST("/keyvalue", handler.AddValueHandler)
	e.GET("/keyvalue/:id", handler.GetValueHandler)
	e.GET("/keyvalue", handler.ListHandler)

}
