package service

import (
	"log"
	"testing"

	"github.com/rodrigomkd/go-rest-api/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockService struct {
	mock.Mock
}

type mockCsvService struct {
	mock.Mock
}

type mockApiService struct {
	mock.Mock
}

func (s mockService) GetItems() ([]model.Activity, error) {
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
	mock := new(mockService)
	mock.On("GetItems").Return()

	service := mock
	activities, err := service.GetItems()

	log.Println("Activities: ", activities)
	log.Println("err: ", err)
	assert.Equal(t, activities[0].Id, 1)
	assert.Equal(t, activities[0].Title, "Activity 1")
	assert.Equal(t, activities[0].DueDate, "2021-10-09T00:19:20.8445283+00:00")
	assert.Equal(t, activities[0].Completed, false)
}
