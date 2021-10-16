package csv

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/rodrigomkd/go-rest-api/model"
)

type ICSVService interface {
	ReadCSV(path string) ([][]string, error)
	SaveCSV(path string, records [][]string)
	SaveActivities(path string, activities []model.Activity)
}

type CSVService struct {
	csv ICSVService
}

func New() *CSVService {
	return &CSVService{}
}

//Read CSV file
func (cs CSVService) ReadCSV(path string) ([][]string, error) {
	csvFile, err := os.Open(path)
	if err != nil {
		log.Print("Error to open CSV file: ", err)
	}

	log.Print("Successfully Opened CSV file")
	defer csvFile.Close()

	return csv.NewReader(csvFile).ReadAll()
}

func (cs CSVService) SaveCSV(path string, records [][]string) {
	if _, err := os.Stat(path); os.IsExist(err) {
		os.Remove(path)
	}

	file, err := os.Create(path)
	checkError("Cannot create file: ", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(records)
	checkError("Cannot write to file", err)
}

func (cs CSVService) SaveActivities(path string, activities []model.Activity) {
	if _, err := os.Stat(path); os.IsExist(err) {
		os.Remove(path)
	}

	file, err := os.Create(path)
	checkError("Cannot create file: ", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	records := convertActivitiesCsv(activities)
	err = writer.WriteAll(records)
	checkError("Cannot write to file", err)
}

func checkError(message string, err error) {
	if err != nil {
		log.Println("ERROR: ", message, err)
	}
}

func convertActivitiesCsv(activities []model.Activity) [][]string {
	var records = [][]string{}

	for i := 0; i < len(activities); i++ {
		activity := activities[i]
		var record = []string{}
		record = append(record, strconv.Itoa(activity.Id))
		record = append(record, activity.Title)
		record = append(record, activity.DueDate)
		record = append(record, strconv.FormatBool(activity.Completed))

		records = append(records, record)
	}

	return records
}
