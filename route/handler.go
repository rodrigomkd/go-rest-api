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
func GetRouter(ic ItemController) *mux.Router {
	// creates a new instance of a mux router
	router := mux.NewRouter().StrictSlash(true)

	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/items", ic.GetItems).Methods(http.MethodGet)
	api.HandleFunc("/items/workers", ic.GetItemsWorkers).Methods(http.MethodGet)
	api.HandleFunc("/items/{id}", ic.GetItem).Methods(http.MethodGet)

	api.HandleFunc("/items/sync", ic.GetItemsSync).Methods(http.MethodPost)

	return router
}
