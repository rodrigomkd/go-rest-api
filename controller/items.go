package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/rodrigomkd/go-rest-api/model"

	"github.com/gorilla/mux"
)

type ItemService interface {
	GetItems() ([]model.Activity, error)
	GetItemsSync() ([]model.Activity, error)
	GetItem(taskID int) (model.Activity, error)
	GetItemsWorker(typ string, items int, itemsPerWork int) ([]model.Worker, error)
}

type ItemController struct {
	is ItemService
}

func New(is ItemService) *ItemController {
	return &ItemController{
		is: is,
	}
}

func (c ItemController) GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items, err := c.is.GetItems()
	if err != nil {
		errorResponse(w, "Some Error Occurred", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}

func (c ItemController) GetItemsSync(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items, err := c.is.GetItems()
	if err != nil {
		errorResponse(w, "Some Error Occurred", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}

func (c ItemController) GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		errorResponse(w, "ID path parameter must be an Integer", http.StatusBadRequest)
		return
	}

	item, err := c.is.GetItem(taskID)
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

func (c ItemController) GetItemsWorkers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	log.Println("params: ", query)
	err, sc := handleQueryParams(query)
	if err != nil {
		errorResponse(w, err.Error(), sc)
		return
	}

	items, err := strconv.Atoi(r.URL.Query().Get("items"))
	itemsPerWork, err := strconv.Atoi(r.URL.Query().Get("items_per_workers"))
	workers, err := c.is.GetItemsWorker(r.URL.Query().Get("type"), items, itemsPerWork)
	if err != nil {
		errorResponse(w, "Some Error Occurred", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(workers)
}

func handleQueryParams(params url.Values) (error, int) {
	if !params.Has("type") {
		return errors.New("type not found"), http.StatusBadRequest
	} else {
		typeParam := params.Get("type")
		if strings.ToLower(typeParam) != "odd" && strings.ToLower(typeParam) != "even" {
			return errors.New("type not valid, valid options: odd and even"), http.StatusBadRequest
		}
	}

	if !params.Has("items") {
		return errors.New("items not found"), http.StatusBadRequest
	} else {
		items, err := strconv.Atoi(params.Get("items"))
		log.Println("items value: ", items)
		if err != nil {
			return errors.New("items param not valid, must be an integer: " + params.Get("items")), http.StatusBadRequest
		}
	}

	if !params.Has("items_per_workers") {
		return errors.New("items_per_workers not found"), http.StatusBadRequest
	} else {
		itemsPerWork, err := strconv.Atoi(params.Get("items_per_workers"))
		log.Println("items_per_workers value: ", itemsPerWork)
		if err != nil {
			return errors.New("items_per_workers param not valid, must be an integer: " + params.Get("items_per_workers")), http.StatusBadRequest
		}
	}

	items, err := strconv.Atoi(params.Get("items"))
	itemsPerWork, err := strconv.Atoi(params.Get("items_per_workers"))

	if err != nil {
		log.Println("ERROR: ", err)
		return errors.New("Some Error Occurred"), http.StatusInternalServerError
	}

	if itemsPerWork > items {
		return errors.New("items_per_workers (" + params.Get("items_per_workers") + ") is higher than items (" + params.Get("items") + ")"), http.StatusBadRequest
	}

	return nil, 0
}

func errorResponse(w http.ResponseWriter, errorMessage string, statusCode int) {
	w.WriteHeader(statusCode)
	resp := make(map[string]string)
	resp["message"] = errorMessage
	json.NewEncoder(w).Encode(resp)
}
