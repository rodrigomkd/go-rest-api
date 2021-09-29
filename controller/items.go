package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rodrigomkd/go-rest-api/service"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items, err := service.GetItems()
	if err != nil {
		errorResponse(w, "Some Error Occurred", 500)
		return
	}

	json.NewEncoder(w).Encode(items)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		errorResponse(w, "ID path parameter must be an Integer", http.StatusBadRequest)
		return
	}

	item, err := service.GetItem(taskID)
	if err != nil {
		log.Println("log error: ", err.Error())
		if err.Error() == strconv.Itoa(http.StatusNotFound) {
			errorResponse(w, "Resource Not Found", http.StatusNotFound)
		} else {
			errorResponse(w, "Some Error Ocurred", http.StatusInternalServerError)
		}

		return
	}

	json.NewEncoder(w).Encode(item)
}

func errorResponse(w http.ResponseWriter, errorMessage string, statusCode int) {
	w.WriteHeader(statusCode)
	resp := make(map[string]string)
	resp["message"] = errorMessage
	json.NewEncoder(w).Encode(resp)
}
