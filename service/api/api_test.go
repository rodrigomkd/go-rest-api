package api

import (
	"log"
	"testing"

	"github.com/rodrigomkd/go-rest-api/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockServiceApi struct {
	mock.Mock
}

func (m *mockServiceApi) GetActivities() []model.Activity {
	log.Println("Mocked GetActivities function")

	activities := []model.Activity{}
	act := model.Activity{}
	act.Id = 1
	act.Title = "Activity 1"
	act.DueDate = "2021-10-09T00:19:20.8445283+00:00"
	act.Completed = false
	activities = append(activities, act)
	return activities
}

func TestReadApi_Activities(t *testing.T) {
	mock := new(mockServiceApi)
	mock.On("GetActivities").Return()

	api := mock
	activities := api.GetActivities()

	log.Println("Activities: ", activities)
	assert.Equal(t, activities[0].Id, 1)
	assert.Equal(t, activities[0].Title, "Activity 1")
	assert.Equal(t, activities[0].DueDate, "2021-10-09T00:19:20.8445283+00:00")
	assert.Equal(t, activities[0].Completed, false)
}
