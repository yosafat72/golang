package main

import (
	container "go-echo/src/interface/container"
	server "go-echo/src/interface/server"
)

func main() {
	server.StartServer(container.SetupContainer())
}
