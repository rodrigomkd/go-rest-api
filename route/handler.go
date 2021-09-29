package route

import (
	"github.com/gorilla/mux"

	"github.com/rodrigomkd/go-rest-api/controller"
)

// HandlerRequests - Returns router that handles API paths
func GetRouter() *mux.Router {
	// creates a new instance of a mux router
	router := mux.NewRouter().StrictSlash(true)

	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/items", controller.GetItems).Methods("GET")
	api.HandleFunc("/items/{id}", controller.GetItem).Methods("GET")

	return router
}
