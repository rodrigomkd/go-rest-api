package services

import (
	"errors"
	"log"
	"strconv"

	"go-rest-api/common"
	"go-rest-api/config"
)

// Types
type item struct {
	ID          int    `json:"ID"`
	Description string `json:"Description"`
}

//Get Items
func GetItems() ([]item, error) {
	items := []item{}

	csvLines, err := common.ReadCSV(config.DATA_SOURCE)
	if err != nil {
		return items, err
	}
	for _, line := range csvLines {
		id, err := strconv.Atoi(line[0])
		if err != nil {
			log.Println("Error ID is not an Integer: ", err)
			return items, err
		}

		emp := item{
			ID:          id,
			Description: line[1],
		}
		items = append(items, emp)
		log.Println("items: ", items)
	}

	return items, nil
}

func GetItem(id int) (item, error) {
	items, err := GetItems()
	if err != nil {
		return item{}, errors.New("500")
	}

	for _, task := range items {
		if task.ID == id {
			log.Println("Item found: ", task)
			return task, nil
		}
	}

	return item{}, errors.New("404")
}
