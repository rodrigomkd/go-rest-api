package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rodrigomkd/go-rest-api/model"

	"github.com/stretchr/testify/assert"
)

type mockItemService struct {
	activity    model.Activity
	getResponse []model.Activity
	getError    error
	workers     []model.Worker
}

func (m *mockItemService) GetItems() ([]model.Activity, error) {
	return m.getResponse, m.getError
}

func (m *mockItemService) GetItemsSync() ([]model.Activity, error) {
	return m.getResponse, m.getError
}

func (m *mockItemService) GetItem(taskID int) (model.Activity, error) {
	return m.activity, m.getError
}

func (m *mockItemService) GetItemsWorker(typ string, items int, itemsPerWork int) ([]model.Worker, error) {
	return m.workers, m.getError
}

func TestItemsController(t *testing.T) {
	// test table
	tests := []struct {
		name string
		// Expected Result
		StatusCode int
		// Expected Error
		Error error
		//mock usecase
		useCaseMock mockItemService
	}{
		{
			name:       "OK Status Code",
			StatusCode: http.StatusOK,
			Error:      nil,
			useCaseMock: mockItemService{
				getResponse: []model.Activity{
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
			},
		},
		{
			name:       "Error Status Code",
			StatusCode: http.StatusInternalServerError,
			Error:      errors.New("strconv.Atoi: parsing \"NO INTEGER\": invalid syntax"),
			useCaseMock: mockItemService{
				getError: errors.New("strconv.Atoi: parsing \"NO INTEGER\": invalid syntax"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			itemsController := New(&tt.useCaseMock)
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/items", nil)
			itemsController.GetItems(w, r)

			if tt.Error == nil {
				assert.Equal(t, tt.StatusCode, w.Result().StatusCode)
			} else {
				assert.Equal(t, tt.StatusCode, w.Result().StatusCode)
			}

			itemsController.GetItem(w, r)
			if tt.Error == nil {
				assert.Equal(t, tt.StatusCode, w.Result().StatusCode)
			} else {
				assert.Equal(t, tt.StatusCode, w.Result().StatusCode)
			}

			itemsController.GetItemsSync(w, r)
			if tt.Error == nil {
				assert.Equal(t, tt.StatusCode, w.Result().StatusCode)
			} else {
				assert.Equal(t, tt.StatusCode, w.Result().StatusCode)
			}

			itemsController.GetItemsWorkers(w, r)
			if tt.Error == nil {
				assert.Equal(t, tt.StatusCode, w.Result().StatusCode)
			} else {
				assert.Equal(t, tt.StatusCode, w.Result().StatusCode)
			}
		})

	}
}
