package handlers_utils

import (
	"AuthService/configs"
	"encoding/json"
	"net/http"
	"reflect"
)

// HandleJsonResponse converts object to a json server response
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
