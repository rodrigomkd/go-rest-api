package csv

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockFile struct {
	mock.Mock
}

func (m *mockFile) Open(name string) (*os.File, error) {
	log.Println("Mocked Open function")
	args := m.Called(name)

	return nil, args.Error(1)
}

func TestReadCsv_Activities(t *testing.T) {
	mock := new(mockFile)
	mock.On("Open", "data.csv").Return(&os.File{}, nil)

	csvService := New("data.csv", mock)
	activities, err := csvService.ReadCSV()

	assert.Equal(t, len(activities), 0)

	log.Println("Activities: ", activities)
	log.Println("Error: ", err)
}
func TestReadCsv_Error(t *testing.T) {
	mock := new(mockFile)
	mock.On("Open", "error.csv").Return(nil, errors.New("Error to open CSV file"))

	csvService := New("error.csv", mock)
	activities, err := csvService.ReadCSV()

	log.Println("Activities: ", activities)
	log.Println("Error: ", err)
	assert.Error(t, err)
}
