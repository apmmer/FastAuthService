package handlers

import (
	"AuthService/internal/repositories/user_repo"
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
// @Failure 404 {object} schemas.ErrorResponse "User not found"
// @Failure 406 {object} schemas.ErrorResponse "Multiple users found"
// @Failure 422 {object} schemas.ErrorResponse "Unprocessable entity"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/users/{id} [get]
func GetUserById(w http.ResponseWriter, r *http.Request) {
	log.Println("Got request to fetch user by ID.")

	// declare filters with user_id
	filters := make(map[string]interface{})
	_, userID := path.Split(r.URL.Path)
	filters["id"] = userID

	user, err := user_repo.GetUser(&filters)
	if err != nil {
		HandleException(w, err)
		return
	}
	log.Println("Successfully got result from user_repo.GetUser")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	err = HandleJsonResponse(w, user)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error while handling JSON response: %v", err), http.StatusInternalServerError)
	}
}
