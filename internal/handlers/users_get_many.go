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
// @Failure 401 {object} schemas.ErrorResponse "Error returned when the provided auth data is invalid"
// @Failure 403 {object} schemas.ErrorResponse "Error returned when auth data was not provided"
// @Failure 422 {object} schemas.ErrorResponse "Unprocessable entity"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
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
		HandleException(w, err)
		return
	}

	log.Printf("Ready to call user_repo, params = %v.", *params)
	// Call GetManyUsers from the repo
	users, err := user_repo.GetManyUsers(*params)
	if err != nil {
		HandleException(w, err)
		return
	}
	log.Printf("Successfully got users = %v", users)

	// Setting the status 200
	w.WriteHeader(http.StatusOK)

	// Prepare response
	err = HandleJsonResponse(w, users)
	if err != nil {
		HandleException(w, fmt.Errorf("Error while handling JSON response: %v", err))
	}
}
