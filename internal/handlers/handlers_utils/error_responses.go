package handlers_utils

import (
	"AuthService/internal/exceptions"
	"encoding/json"
	"net/http"
)

// Sets a body for error response with provided message and status code
func ErrorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	errMsg := map[string]string{"error": message}
	json.NewEncoder(w).Encode(errMsg)
}

// Converts custom error to http error response with correct status code.
func HandleExceptionResponse(w http.ResponseWriter, err error) {
	if customErr, ok := err.(*exceptions.DefaultError); ok {
		ErrorResponse(w, customErr.Error(), customErr.StatusCode)
	} else {
		ErrorResponse(w, "Server can not process your request: unhandled error.", http.StatusInternalServerError)
	}
}
