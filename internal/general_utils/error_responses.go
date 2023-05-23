package general_utils

import (
	"AuthService/internal/exceptions"
	"encoding/json"
	"log"
	"net/http"
)

// Sets a body for error response with provided message and status code
func ErrorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	errMsg := map[string]string{"error": message}
	json.NewEncoder(w).Encode(errMsg)
}

// HandleException transforms exception to https response with valid status code.
func HandleException(w http.ResponseWriter, err error) {
	switch err.(type) {
	case *exceptions.ErrNotFound:
		log.Println("HandleException: ErrNotFound...")
		ErrorResponse(w, err.Error(), http.StatusNotFound)
	case *exceptions.ErrMultipleEntries:
		log.Println("HandleException: ErrMultipleEntries...")
		ErrorResponse(w, err.Error(), http.StatusNotAcceptable)
	case *exceptions.ErrInvalidEntity:
		log.Println("HandleException: ErrInvalidEntity...")
		ErrorResponse(w, err.Error(), http.StatusUnprocessableEntity)
	case *exceptions.ErrDbConflict:
		log.Println("HandleException: ErrDbConflict...")
		ErrorResponse(w, err.Error(), http.StatusConflict)
	case *exceptions.ErrNoAuthData:
		log.Println("HandleException: ErrNoAuthData...")
		ErrorResponse(w, err.Error(), http.StatusForbidden)
	case *exceptions.ErrUnauthorized:
		log.Println("HandleException: ErrUnauthorized...")
		ErrorResponse(w, err.Error(), http.StatusUnauthorized)
	case *exceptions.ErrInputError:
		log.Println("HandleException: ErrInputError...")
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
	default:
		log.Println("HandleException: default Exception...")
		ErrorResponse(w, "Server can not process your request", http.StatusInternalServerError)
	}
}

func HandleExceptionNew(w http.ResponseWriter, err error) {
	if customErr, ok := err.(*exceptions.DefaultError); ok {
		log.Println("HandleException: ErrNotFound...")
		ErrorResponse(w, customErr.Error(), customErr.StatusCode)
	} else {
		ErrorResponse(w, "Server can not process your request: unhandled error.", http.StatusInternalServerError)
	}
}
