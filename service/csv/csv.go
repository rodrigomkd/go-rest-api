package csv

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/rodrigomkd/go-rest-api/model"
)

type IFile interface {
	Open(name string) (*os.File, error)
}

type CSVService struct {
	path string
	file IFile
}

func New(path string, file IFile) *CSVService {
	return &CSVService{
		path: path,
		file: file,
	}
}

//Read CSV file
func (cs CSVService) ReadCSV() ([][]string, error) {
	csvFile, err := cs.file.Open(cs.path)
	if err != nil {
		log.Print("Error to open CSV file: ", err)
		return nil, err
	}

	log.Print("Successfully Opened CSV file: ", csvFile)
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
