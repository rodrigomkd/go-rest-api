package csv

import (
	"errors"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockServiceCsv struct {
	mock.Mock
}

func (m *mockServiceCsv) ReadCSV(path string) ([][]string, error) {
	log.Println("Mocked ReadCSV function")
	args := m.Called(path)

	activities := [][]string{{"1", "Activity 1", "2021-10-09T00:19:20.8445283+00:00", "false"}}
	return activities, args.Error(1)
}

func TestReadCsv_Activities(t *testing.T) {
	mockCsv := new(mockServiceCsv)
	mockCsv.On("ReadCSV", "data.csv").Return(make([][]string, 0), nil)

	csvService := mockCsv
	activities, err := csvService.ReadCSV("data.csv")

	log.Println("Activities: ", activities)
	assert.Equal(t, activities[0][0], "1")
	assert.Equal(t, activities[0][1], "Activity 1")
	assert.Equal(t, activities[0][2], "2021-10-09T00:19:20.8445283+00:00")
	assert.Equal(t, activities[0][3], "false")

	log.Println("Activities: ", activities[0][0])
	assert.NoError(t, err)
}
func TestReadCsv_Error(t *testing.T) {
	mockCsv := new(mockServiceCsv)
	mockCsv.On("ReadCSV", "error.csv").Return(make([][]string, 0), errors.New("Error to open CSV file"))

	csvService := mockCsv
	activities, err := csvService.ReadCSV("error.csv")

	log.Println("Activities: ", activities)
	log.Println("Error: ", err)
	assert.Error(t, err)
}
