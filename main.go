package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/rodrigomkd/go-rest-api/config"
	"github.com/rodrigomkd/go-rest-api/route"
)

//main - Start Server
func main() {
	router := route.GetRouter()
	port := config.ReadConfig().ServerPort

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
