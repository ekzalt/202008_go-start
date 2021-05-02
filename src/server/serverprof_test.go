package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	// "github.com/stretchr/testify/assert"
)

// TestHandleHome tests Home controller
func TestHandleHome(t *testing.T) {
	handler := http.HandlerFunc(HandleHome)

	t.Run("home", func(t *testing.T) {
		// response will be written here
		recorder := httptest.NewRecorder()
		// create test request
		request, _ := http.NewRequest("GET", "/", nil)
		// execute handler to get response
		handler.ServeHTTP(recorder, request)

		// assert response result here
		fmt.Println((recorder.Result().Status))
		// assert.Equal(t, "200 OK", recorder.Result().Status)
	})
}
