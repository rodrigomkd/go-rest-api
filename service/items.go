package service

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/rodrigomkd/go-rest-api/common"
	"github.com/rodrigomkd/go-rest-api/config"
)

type item struct {
	ID          int    `json:"ID"`
	Description string `json:"Description"`
}

//GetItems - Returns items from CSV
func GetItems() ([]item, error) {
	items := []item{}

	csvLines, err := common.ReadCSV(config.ReadConfig().DataSource)
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

		temp := item{
			ID:          id,
			Description: line[1],
		}
		items = append(items, temp)
		log.Println("items: ", items)
	}

	return items, nil
}

func GetItem(id int) (item, error) {
	items, err := GetItems()
	if err != nil {
		return item{}, errors.New(strconv.Itoa(http.StatusInternalServerError))
	}

	for _, task := range items {
		if task.ID == id {
			log.Println("Item found: ", task)
			return task, nil
		}
	}

	return item{}, errors.New(strconv.Itoa(http.StatusNotFound))
}
