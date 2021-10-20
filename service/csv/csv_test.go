package csv

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/rodrigomkd/github.com/stretchr/testify/assert"
)

//temp file folder
const TEMP_FILE = ""

type mockFile struct {
	getFile  *os.File
	getError error
}

func (m *mockFile) Open(name string) (*os.File, error) {
	return m.getFile, m.getError
}

func newTempFile(path string, name string) (f *os.File) {
	f, err := os.CreateTemp(path, name+"*.csv")
	if err != nil {
		log.Fatal("TempFile: ", err)
	}
	f.WriteString("1,Activity 1,2021-10-14T01:07:35.3812533+00:00,false")

	return f
}

func TestCsvService(t *testing.T) {
	//create temp file
	f := newTempFile(TEMP_FILE, "temp_")
	tests := []struct {
		name string
		// Expected Result
		Result [][]string
		// Expected Error
		Error error
		//mock usecase
		useCaseMock mockFile
	}{
		{
			name:   "Open File",
			Result: [][]string{{"1", "Activity 1", "2021-10-14T01:07:35.3812533+00:00", "false"}},
			Error:  nil,
			useCaseMock: mockFile{
				getFile: &os.File{},
			},
		},
		{
			name:   "Invalid Response from API",
			Result: make([][]string, 0),
			Error:  errors.New("Ocurred an error."),
			useCaseMock: mockFile{
				getFile:  &os.File{},
				getError: errors.New("Ocurred an error."),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			csv := New(TEMP_FILE+f.Name(), &tt.useCaseMock)
			output, err := csv.ReadCSV()

			log.Println("CSV Output: ", output)
			log.Println("Error: ", err)
			if tt.Error == nil {
				assert.Equal(t, tt.Result, output)
			} else {
				assert.Equal(t, tt.Error, err)
			}
		})
	}

	//remove temp file
	defer os.RemoveAll(f.Name())
}
