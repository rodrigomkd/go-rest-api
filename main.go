package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/rodrigomkd/go-rest-api/config"
	"github.com/rodrigomkd/go-rest-api/controller"
	"github.com/rodrigomkd/go-rest-api/route"
	"github.com/rodrigomkd/go-rest-api/service"
	"github.com/rodrigomkd/go-rest-api/service/api"
	"github.com/rodrigomkd/go-rest-api/service/csv"
)

type File struct {
	file csv.IFile
}

func (File) Open(name string) (*os.File, error) {
	return os.Open(name)
}

//main - Start Server
func main() {
	//Get config data
	config := config.ReadConfig("properties.ini")

	//create csv service
	file := &File{}
	cs := csv.New(config.DataSource, file)
	//create csv worker service
	cws := csv.NewWorker(config.DataSourceWorker)
	//create api service
	as := api.New(config.ApiUri, &http.Client{})

	//create item service
	s := service.New(*cs, *cws, *as, config.DataSource)

	//create controller
	c := controller.New(*s)

	//create router
	router := route.GetRouter(*c)
	port := config.ServerPort
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
