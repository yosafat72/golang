package server

import (
	"fmt"

	container "go-echo/src/interface/container"

	transport "go-echo/src/transport"

	"github.com/labstack/echo/v4"
)

func StartServer(cons container.Container) {
	//Init echo
	app := echo.New()

	transport := transport.SetupTransport(&cons)
	setupRouter(transport, app)

	//Run echo apps
	fmt.Println(app.Start(":1323"))
}
