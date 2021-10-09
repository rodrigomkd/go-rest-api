package route

import (
	"github.com/rodrigomkd/go-rest-api/controller"

	"github.com/gorilla/mux"
)

// HandlerRequests - Returns router that handles API paths
func GetRouter(c controller.Controller) *mux.Router {
	// creates a new instance of a mux router
	router := mux.NewRouter().StrictSlash(true)

	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/items", c.GetItems).Methods("GET")
	api.HandleFunc("/items/{id}", c.GetItem).Methods("GET")
	api.HandleFunc("/items/sync", c.GetItemsSync).Methods("POST")

	return router
}
