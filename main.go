package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/rodrigomkd/go-rest-api/config"
	"github.com/rodrigomkd/go-rest-api/controller"
	"github.com/rodrigomkd/go-rest-api/route"
	"github.com/rodrigomkd/go-rest-api/service"
	"github.com/rodrigomkd/go-rest-api/service/api"
	"github.com/rodrigomkd/go-rest-api/service/csv"
)

//main - Start Server
func main() {
	//Get config data
	config := config.ReadConfig("properties.ini")

	//create csv service
	cs := csv.New()
	//create api service
	as := api.New(config.ApiUri)

	//create item service
	s := service.New(*cs, *as, config.DataSource)

	//create controller
	c := controller.New(*s)

	//create router
	router := route.GetRouter(*c)
	port := config.ServerPort
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
