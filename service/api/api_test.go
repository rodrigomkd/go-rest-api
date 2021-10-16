package api

import (
	"go-rest-api/model"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockHttpClient struct {
	mock.Mock
}

func (m *mockHttpClient) Get(url string) (resp *http.Response, err error) {
	log.Println("Mocked Get function")
	return httptest.NewRecorder().Result(), nil
}

func TestApiService(t *testing.T) {
	// test table
	tests := []struct {
		name string
		// Body mock the response body
		Body string
		// StatusCode mock the response statusCode
		StatusCode int

		// Expected Result
		Result *model.Activity
		// Expected Error
		Error error
	}{
		{
			name:       "valid response",
			Body:       `{"userId": 1,"id": 1,"title": "test title","body": "test body"}`,
			StatusCode: 200,
			Result: &model.Activity{
				Id:        1,
				Title:     "1",
				DueDate:   "test title",
				Completed: false,
			},
			Error: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := new(mockHttpClient)
			//mock.On("Get").Return("algo")

			//pass method mock
			api := New("testing", mock)
			activities := api.GetActivities()

			log.Println("Activities: ", activities)
			assert.Equal(t, len(activities), 1)
		})

	}
}

func TestReadApi_Activities(t *testing.T) {
	mock := new(mockHttpClient)
	//mock.On("Get").Return("algo")

	//pass method mock
	api := New("testing", mock)
	activities := api.GetActivities()

	log.Println("Activities: ", activities)
	assert.Equal(t, len(activities), 0)
	/*
		assert.Equal(t, activities[0].Id, 1)
		assert.Equal(t, activities[0].Title, "Activity 1")
		assert.Equal(t, activities[0].DueDate, "2021-10-09T00:19:20.8445283+00:00")
		assert.Equal(t, activities[0].Completed, false)
	*/
}
