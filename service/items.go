package service

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/rodrigomkd/go-rest-api/model"
	"github.com/rodrigomkd/go-rest-api/service/api"
	"github.com/rodrigomkd/go-rest-api/service/csv"
)

//struct
type Service struct {
	cs         csv.CSVService
	api        api.ApiService
	dataSource string
}

func New(cs csv.CSVService, api api.ApiService, dataSource string) *Service {
	return &Service{
		cs:         cs,
		api:        api,
		dataSource: dataSource,
	}
}

//GetItems - Returns items from CSV
func (s Service) GetItems() ([]model.Activity, error) {
	items := s.api.GetActivities()

	s.cs.SaveActivities(s.dataSource, items)

	return items, nil
}

func (s Service) GetItemsSync() ([]model.Activity, error) {
	return s.GetItems()
}

func (s Service) GetItem(id int) (model.Activity, error) {
	items, err := s.GetItems()
	if err != nil {
		return model.Activity{}, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	for _, act := range items {
		if act.Id == id {
			log.Println("Item found: ", act)
			return act, nil
		}
	}

	return model.Activity{}, errors.New(strconv.Itoa(http.StatusNotFound))
}
