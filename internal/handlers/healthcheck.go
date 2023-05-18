package handlers

import (
	"AuthService/internal/schemas"
	"encoding/json" // go get encoding/json
	"fmt"
	"net/http"
)

// HealthCheck godoc
// @Summary Show server status
// @Description Get server status
// @Tags Health
// @Accept  json
// @Produce  json
// @Success 200 {object} schemas.HealthCheckResponse
// @Failure 401 {object} schemas.ErrorResponse "Wrong auth data"
// @Failure 403 {object} schemas.ErrorResponse "No auth data"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/healthcheck [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	// Setting the response header indicating the type of content
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("HealthCheck OK!")
	// Create an instance of the structure and set the status "OK"
	response := schemas.HealthCheckResponse{
		Status: "OK",
	}

	// Setting the response status and sending JSON with the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
