package handlers_test

import (
	"auth_service_api/internal/handlers"
	"auth_service_api/internal/schemas"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HealthCheck)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response schemas.HealthCheckResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Could not decode response: %v", err)
	}

	expected := schemas.HealthCheckResponse{Status: "OK"}
	if !reflect.DeepEqual(response, expected) {
		t.Errorf("handler returned unexpected body: got %+v want %+v",
			response, expected)
	}
}
