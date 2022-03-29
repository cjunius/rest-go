package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

var test_entity = map[string]interface{}{
	"name": "Test Entity",
}

func TestGetEntities_NotFound(t *testing.T) {

	request, _ := http.NewRequest("GET", "/persons", nil)
	response := httptest.NewRecorder()

	GetEntities(response, request)

	if response.Code != http.StatusNotFound {
		t.Fatalf("Non-expected status code %v:\n\tbody: %v", "404", response.Code)
	}

}

func TestGetEntity_NotFound(t *testing.T) {
	request, _ := http.NewRequest("GET", "/persons/1", nil)
	response := httptest.NewRecorder()

	vars := map[string]string{
		"id": "1",
	}

	request = mux.SetURLVars(request, vars)

	GetEntity(response, request)

	if response.Code != http.StatusNotFound {
		t.Fatalf("Non-expected status code %v:\n\tbody: %v", "404", response.Code)
	}

}

func TestReplaceEntity_NotFound(t *testing.T) {
	jsonBody, _ := json.Marshal(test_entity)
	request, _ := http.NewRequest("POST", "/persons/1", bytes.NewBuffer(jsonBody))
	response := httptest.NewRecorder()

	vars := map[string]string{
		"id": "1",
	}

	request = mux.SetURLVars(request, vars)

	ReplaceEntity(response, request)

	if response.Code != http.StatusNotFound {
		t.Fatalf("Non-expected status code %v:\n\tbody: %v", "404", response.Code)
	}

}

func TestUpdateEntity_NotFound(t *testing.T) {
	jsonBody, _ := json.Marshal(test_entity)
	request, _ := http.NewRequest("POST", "/persons/1", bytes.NewBuffer(jsonBody))
	response := httptest.NewRecorder()

	vars := map[string]string{
		"id": "1",
	}

	request = mux.SetURLVars(request, vars)

	UpdateEntity(response, request)

	if response.Code != http.StatusNotFound {
		t.Fatalf("Non-expected status code %v:\n\tbody: %v", "404", response.Code)
	}

}

func TestDeleteEntity_NotFound(t *testing.T) {
	request, _ := http.NewRequest("DELETE", "/persons/1", nil)
	response := httptest.NewRecorder()

	vars := map[string]string{
		"id": "1",
	}

	request = mux.SetURLVars(request, vars)

	DeleteEntity(response, request)

	if response.Code != http.StatusNotFound {
		t.Fatalf("Non-expected status code %v:\n\tbody: %v", "404", response.Code)
	}

}
