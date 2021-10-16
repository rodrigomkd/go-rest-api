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
	csw        csv.CSVWorkerService
	api        api.ApiService
	dataSource string
}

func New(cs csv.CSVService, csw csv.CSVWorkerService, api api.ApiService, dataSource string) *Service {
	return &Service{
		cs:         cs,
		csw:        csw,
		api:        api,
		dataSource: dataSource,
	}
}

//GetItems - Returns items from CSV
func (s Service) GetItems() ([]model.Activity, error) {
	items := []model.Activity{}

	csvLines, err := s.cs.ReadCSV()
	if err != nil {
		log.Println("Error reading CSV: ", err)
		return items, err
	}
	for _, line := range csvLines {
		id, err := strconv.Atoi(line[0])
		if err != nil {
			log.Println("Error ID is not an Integer: ", err)
			return items, err
		}

		completed, err := strconv.ParseBool(line[3])
		if err != nil {
			log.Println("Error ID is not an Integer: ", err)
			return items, err
		}

		temp := model.Activity{
			Id:        id,
			Title:     line[1],
			DueDate:   line[2],
			Completed: completed,
		}
		items = append(items, temp)
		log.Println("Activities: ", items)
	}

	return items, nil
}

//GetItems - Read API and save into CSV
func (s Service) GetItemsSync() ([]model.Activity, error) {
	//read api
	items := s.api.GetActivities()
	//save content
	s.cs.SaveActivities(items)

	return items, nil
}

//GetItemsWorker - Read from CSV using a worker pool
func (s Service) GetItemsWorker(typ string, items int, itemsPerWork int) ([]model.Worker, error) {
	return s.csw.ReadWorkers(typ, items, itemsPerWork), nil
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
