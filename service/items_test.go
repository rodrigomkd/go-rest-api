package service

import (
	"errors"
	"log"
	"testing"

	"github.com/rodrigomkd/go-rest-api/model"

	"github.com/stretchr/testify/assert"
)

type mockCsvService struct {
	getResponse [][]string
	getError    error
}

type mockApiService struct {
	getResponse []model.Activity
	getError    error
}

type mockWorkerService struct {
	getResponse []model.Worker
	getError    error
}

func (m *mockCsvService) ReadCSV() ([][]string, error) {
	log.Println("Mocked Get function")
	return m.getResponse, m.getError
}

func (m *mockCsvService) SaveActivities(activities []model.Activity) error {
	log.Println("Mocked SaveActivities function")
	return m.getError
}

func (m *mockWorkerService) ReadWorkers(typ string, items int, itemsPerWork int) []model.Worker {
	log.Println("Mocked SaveActivities function")
	return m.getResponse
}

func (m *mockApiService) GetActivities() ([]model.Activity, error) {
	log.Println("Mocked SaveActivities function")
	return m.getResponse, m.getError
}

func TestItemsService(t *testing.T) {
	// test table
	tests := []struct {
		name string
		// Expected Result
		Result []model.Activity
		// Expected Error
		Error error
		//mock usecase
		useCaseCsvMock mockCsvService
		//mock api
		useCaseApiMock mockApiService
		//mock worker
		useCaseWorkerMock mockWorkerService
	}{
		{
			name: "Valid Response from CSV Service",
			Result: []model.Activity{
				{
					Id:        1,
					Title:     "Title 1",
					DueDate:   "2021-10-18T18:53:16.2413765+00:00",
					Completed: false,
				},
				{
					Id:        5,
					Title:     "Title 5",
					DueDate:   "2021-10-18T18:53:16.2413765+00:00",
					Completed: true,
				},
			},
			Error: nil,
			useCaseCsvMock: mockCsvService{
				getResponse: [][]string{
					{"1", "Title 1", "2021-10-18T18:53:16.2413765+00:00", "False"},
					{"5", "Title 5", "2021-10-18T18:53:16.2413765+00:00", "True"},
				},
			},
		},
		{
			name: "Invalid Data from CSV Service",
			Result: []model.Activity{
				{
					Id:        1,
					Title:     "Title 1",
					DueDate:   "2021-10-18T18:53:16.2413765+00:00",
					Completed: false,
				},
				{
					Id:        5,
					Title:     "Title 5",
					DueDate:   "2021-10-18T18:53:16.2413765+00:00",
					Completed: true,
				},
			},
			Error: errors.New("strconv.Atoi: parsing \"NO INTEGER\": invalid syntax"),
			useCaseCsvMock: mockCsvService{
				getResponse: [][]string{
					{"NO INTEGER", "Title 1", "2021-10-18T18:53:16.2413765+00:00", "False"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			itemsService := New(&tt.useCaseCsvMock, &tt.useCaseWorkerMock, &tt.useCaseApiMock, "test")
			activities, err := itemsService.GetItems()

			log.Println("Activities: ", activities)
			log.Println("Error: ", err)
			if tt.Error == nil {
				assert.Equal(t, tt.Result, activities)
			} else {
				assert.Equal(t, tt.Error.Error(), err.Error())
			}
		})

	}
}
