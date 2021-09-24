package app

import (
	"log"
	"net/http"
	"strconv"

	"go-rest-api/routes"
)

//Start Server
func StartServer(port int) {
	router := routes.HandleRequests()

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
