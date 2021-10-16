package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

type ItemController interface {
	GetItems(w http.ResponseWriter, req *http.Request)
	GetItem(w http.ResponseWriter, req *http.Request)
	GetItemsSync(w http.ResponseWriter, req *http.Request)
	GetItemsWorkers(w http.ResponseWriter, req *http.Request)
}

// HandlerRequests - Returns router that handles API paths
func GetRouter(c ItemController) *mux.Router {
	// creates a new instance of a mux router
	router := mux.NewRouter().StrictSlash(true)

	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/items", c.GetItems).Methods(http.MethodGet)
	api.HandleFunc("/items/workers", c.GetItemsWorkers).Methods(http.MethodGet)
	api.HandleFunc("/items/{id}", c.GetItem).Methods(http.MethodGet)

	api.HandleFunc("/items/sync", c.GetItemsSync).Methods(http.MethodPost)

	return router
}
