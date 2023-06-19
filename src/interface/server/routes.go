package server

import (
	transport "go-echo/src/transport"
	"net/http"

	"github.com/labstack/echo/v4"
)

func setupRouter(transport *transport.Tp, app *echo.Echo) {

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Yosafat Hermawan S")
	})

	app.POST("/allUsers", transport.UserTransport.FindAllUsers())
	app.POST("/", transport.UserTransport.FindByIdUser())
	app.POST("/save", transport.UserTransport.SaveUser())

	// app.Use(middleware.FindByIdUserValidator)

}
