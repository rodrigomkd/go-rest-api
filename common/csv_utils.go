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

func SaveCSV(path string, records [][]string) {
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

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
