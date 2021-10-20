package api

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/rodrigomkd/go-rest-api/model"

	"github.com/stretchr/testify/assert"
)

type mockHttpClient struct {
	getResponse *http.Response
	getError    error
}

func (m *mockHttpClient) Get(url string) (resp *http.Response, err error) {
	log.Println("Mocked Get function")
	return m.getResponse, m.getError
}

func TestApiService(t *testing.T) {
	// test table
	tests := []struct {
		name string
		// Expected Result
		Result model.Activity
		// Expected Error
		Error error
		//mock usecase
		useCaseMock mockHttpClient
	}{
		{
			name: "Valid Response from API",
			Result: model.Activity{
				Id:        1,
				Title:     "test title",
				DueDate:   "2021-10-18T18:53:16.2413765+00:00",
				Completed: false,
			},
			Error: nil,
			useCaseMock: mockHttpClient{
				getResponse: &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewBufferString(string(`[{"id": 1,"title": "test title","dueDate": "2021-10-18T18:53:16.2413765+00:00","completed":false}]`))),
				},
			},
		},
		{
			name:   "Invalid Response from API",
			Result: model.Activity{},
			Error:  errors.New("Ocurred an error."),
			useCaseMock: mockHttpClient{
				getResponse: &http.Response{
					StatusCode: http.StatusInternalServerError,
				},
				getError: errors.New("Ocurred an error."),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := New("testing", &tt.useCaseMock)
			activities, err := api.GetActivities()

			log.Println("Activities: ", activities)
			log.Println("Error: ", err)
			if tt.Error == nil {
				assert.Equal(t, tt.Result, activities[0])
			} else {
				assert.Equal(t, tt.Error, err)
			}
		})

	}
}
