package handlers

import (
	"AuthService/configs"
	"AuthService/internal/repositories/user_repo"
	"AuthService/internal/schemas"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// GetManyUsers godoc
// @Summary Get many users
// @Description get many users based on pagination and sorting parameters
// @Tags users
// @Accept json
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Param sort query string false "Sorting (format: field[direction])"
// @Success 200 {array} models.User
// @Router /api/users [get]
func GetManyUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Got request to fetch many users.")
	// Create GetManyRequestParams
	var params schemas.GetManyRequestParams

	// Get and set the limit if it exists
	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
			return
		}
		params.Limit = &limit
	}

	// Get and set the offset if it exists
	if offsetStr := r.URL.Query().Get("offset"); offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			http.Error(w, "Invalid offset parameter", http.StatusBadRequest)
			return
		}
		params.Offset = &offset
	}

	// Get and set the sort if it exists
	if sort := r.URL.Query().Get("sort"); sort != "" {
		params.Sorting = &sort
	}

	// Validate params
	err := params.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log_string := fmt.Sprintf("%d", params)
	log.Println("Ready to call repository, params = ." + log_string)
	// Call GetManyUsers from the repo
	users, err := user_repo.GetManyUsers(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Successfully got result from user_repo.GetManyUsers")

	// Setting the status 200
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// Prepare response
	if configs.MainSettings.Debug == "true" {
		prettyJSON, err := json.MarshalIndent(users, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(prettyJSON)
	} else {
		json.NewEncoder(w).Encode(users)
	}
}
