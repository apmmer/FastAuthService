package handlers_utils

import (
	"AuthService/configs"
	"AuthService/internal/exceptions"
	"encoding/json"
	"log"
	"net/http"
	"reflect"
)

func ErrorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	errMsg := map[string]string{"error": message}
	json.NewEncoder(w).Encode(errMsg)
}

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
	default:
		log.Println("HandleException: default Exception...")
		ErrorResponse(w, "Server can not process your request", http.StatusInternalServerError)
	}
}

func HandleJsonResponse(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")

	// Empty slice checking:
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Slice && v.IsNil() {
		data = reflect.MakeSlice(v.Type(), 0, 0).Interface()
	}

	if configs.MainSettings.Debug == "true" {
		prettyJSON, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
		w.Write(prettyJSON)
	} else {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
	}
	return nil
}
