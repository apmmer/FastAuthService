package handlers

import (
	"AuthService/internal/handlers/handlers_utils"
	"AuthService/internal/repositories/users_repo"
	"fmt"
	"log"
	"net/http"
	"path"
)

// @Summary Get user by ID
// @Description Retrieve a user from the database by ID
// @Tags Users
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 401 {object} schemas.ErrorResponse "Error returned when the provided auth data is invalid"
// @Failure 403 {object} schemas.ErrorResponse "Error returned when auth data was not provided"
// @Failure 404 {object} schemas.ErrorResponse "User not found"
// @Failure 406 {object} schemas.ErrorResponse "Multiple records found (internal error)"
// @Failure 422 {object} schemas.ErrorResponse "Unprocessable entity"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/users/{id} [get]
func GetUserById(w http.ResponseWriter, r *http.Request) {
	log.Println("Got request to fetch user by ID.")

	// declare filters with user_id
	filters := make(map[string]interface{})
	_, userID := path.Split(r.URL.Path)
	filters["id"] = userID

	user, err := users_repo.GetUser(&filters)
	if err != nil {
		handlers_utils.HandleException(w, err)
		return
	}
	log.Println("Successfully got result from users_repo.GetUser")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	err = handlers_utils.HandleJsonResponse(w, user)
	if err != nil {
		handlers_utils.ErrorResponse(w, fmt.Sprintf("Error while handling JSON response: %v", err), http.StatusInternalServerError)
	}
}
