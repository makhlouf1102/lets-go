package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	// unmarshal the response body
	var response map[string]interface{}
	json.NewDecoder(w.Body).Decode(&response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, response["message"], "pong")
}
