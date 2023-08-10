package server

import (
	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func init() {
	e = echo.New()
}

func StartServer() {
	// TODO

}
