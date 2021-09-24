package main

import (
	"go-rest-api/app"
	"go-rest-api/config"
)

func main() {
	app.StartServer(config.SERVER_PORT)
}
