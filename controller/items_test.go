package controller

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockController struct {
	mock.Mock
}

func (m *mockController) GetItems(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock GetItems method")
}

func TestReadApi_Activities(t *testing.T) {
	c := new(mockController)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/items", strings.NewReader("{}"))

	c.GetItems(w, r)

	assert.Equal(t, 200, w.Code)
}
