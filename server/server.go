package server

import (
	"Kiwi/controllerlayer"
	"Kiwi/database"
	"Kiwi/datalayer"
	"Kiwi/servicelayer"
	"log"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var e *echo.Echo

func init() {
	e = echo.New()
}

func StartServer() {
	// TODO
	db, err := database.GetConnection()
	if err != nil {
		log.Fatal(err)
	}

	// swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	userDL := datalayer.NewUserDataLayer(db)
	userSL := servicelayer.NewUserServiceLayer(userDL)
	userCtl := controllerlayer.NewUserController(userSL)
	userRoutes(e, userCtl)

	log.Fatal(e.Start(":8080"))

}
