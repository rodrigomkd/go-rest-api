package csv

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/rodrigomkd/go-rest-api/model"
)

type ICSVService interface {
	ReadCSV() ([][]string, error)
	SaveCSV(records [][]string)
	SaveActivities(activities []model.Activity)
}

type CSVService struct {
	csv        ICSVService
	path       string
	activities []model.Activity
}

func New(path string) *CSVService {
	return &CSVService{
		path: path,
	}
}

//Read CSV file
func (cs CSVService) ReadCSV() ([][]string, error) {
	csvFile, err := os.Open(cs.path)
	if err != nil {
		log.Print("Error to open CSV file: ", err)
		return nil, err
	}

	log.Print("Successfully Opened CSV file")
	defer csvFile.Close()

	return csv.NewReader(csvFile).ReadAll()
}

func (cs CSVService) SaveCSV(records [][]string) error {
	if _, err := os.Stat(cs.path); os.IsExist(err) {
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}
		log.Println("File was found, removing file... ", cs.path)
		os.Remove(cs.path)
	}

	file, err := os.Create(cs.path)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(records)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	return nil
}

func (cs CSVService) SaveActivities(activities []model.Activity) error {
	if _, err := os.Stat(cs.path); os.IsExist(err) {
		log.Println("Removing existing path...", cs.path)
		os.Remove(cs.path)
	}

	file, err := os.Create(cs.path)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	records := convertActivitiesCsv(activities)
	err = writer.WriteAll(records)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	return nil
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
