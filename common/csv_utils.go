package common

import (
	"encoding/csv"
	"log"
	"os"
)

//Read CSV file
func ReadCSV(path string) ([][]string, error) {
	csvFile, err := os.Open(path)
	if err != nil {
		log.Print("Error to open CSV file: ", err)
	}

	log.Print("Successfully Opened CSV file")
	defer csvFile.Close()

	return csv.NewReader(csvFile).ReadAll()
}
