package service

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/rodrigomkd/go-rest-api/model"
	"github.com/rodrigomkd/go-rest-api/service/api"
	"github.com/rodrigomkd/go-rest-api/service/csv"

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

type mockCsvService struct {
	mock.Mock
}

type mockCsvWorkerService struct {
	mock.Mock
}

type mockApiService struct {
	mock.Mock
}

func (s mockCsvService) ReadCSV() ([][]string, error) {
	log.Println("Mocked GetItems function")

	activities := []model.Activity{}
	act := model.Activity{}
	act.Id = 1
	act.Title = "Activity 1"
	act.DueDate = "2021-10-09T00:19:20.8445283+00:00"
	act.Completed = false
	activities = append(activities, act)

	return nil, nil
}

func (s mockCsvService) GetItems() ([]model.Activity, error) {
	log.Println("Mocked GetItems function")

	activities := []model.Activity{}
	act := model.Activity{}
	act.Id = 1
	act.Title = "Activity 1"
	act.DueDate = "2021-10-09T00:19:20.8445283+00:00"
	act.Completed = false
	activities = append(activities, act)

	return activities, nil
}

func TestService_GetItems(t *testing.T) {
	csvService := csv.New("test", &mockFile{})
	worker := csv.NewWorker("test")
	api := api.New("activities", &http.Client{})
	mockCsv := new(mockCsvService)
	mockCsv.On("ReadCSV").Return()

	service := New(*csvService, *worker, *api, "test")
	activities, err := service.GetItems()

	log.Println("Activities: ", activities)
	log.Println("err: ", err)
	assert.Equal(t, activities[0].Id, 1)
	assert.Equal(t, activities[0].Title, "Activity 1")
	assert.Equal(t, activities[0].DueDate, "2021-10-09T00:19:20.8445283+00:00")
	assert.Equal(t, activities[0].Completed, false)
}
