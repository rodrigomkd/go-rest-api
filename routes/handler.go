package routes

import (
	"github.com/gorilla/mux"

	"go-rest-api/controllers"
)

const VERSION = "/api/v1"

//handler requests
func HandleRequests() *mux.Router {
	// creates a new instance of a mux router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(VERSION+"/items", controllers.GetItems).Methods("GET")
	router.HandleFunc(VERSION+"/items/{id}", controllers.GetItem).Methods("GET")

	return router
}
