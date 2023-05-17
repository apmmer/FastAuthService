package handlers

import (
	"AuthService/internal/repositories/user_repo"
	"AuthService/internal/schemas"
	"fmt"
	"log"
	"net/http"
)

// GetManyUsers godoc
// @Summary Get many users
// @Description get many users based on pagination and sorting parameters
// @Tags Users
// @Accept json
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Param sort query string false "Sorting (format: field[direction])"
// @Success 200 {array} models.User
// @Router /api/users [get]
func GetManyUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Got request to fetch many users.")

	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")
	sortStr := r.URL.Query().Get("sort")

	params, err := schemas.GetValidatedListParams(
		limitStr, offsetStr, sortStr,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	log_string := fmt.Sprintf("%d", *params)
	log.Println("Ready to call repository, params = ." + log_string)
	// Call GetManyUsers from the repo
	users, err := user_repo.GetManyUsers(*params)
	if err != nil {
		HandleException(w, err)
	}

	log.Println("Successfully got result from user_repo.GetManyUsers")

	// Setting the status 200
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// Prepare response
	err = HandleJsonResponse(w, users)
	if err != nil {
		log.Println("Error while handling JSON response:", err)
	}
}
