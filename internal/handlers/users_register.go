package handlers

import (
	"AuthService/internal/general_utils"
	"AuthService/internal/handlers/handlers_utils"
	"AuthService/internal/repositories/users_repo"
	"AuthService/internal/schemas"
	"encoding/json"
	"log"
	"net/http"
)

// RegisterUser godoc
// @Summary Register new user
// @Description Register a new user with email, screen_name, and password
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body schemas.CreateUserRequest true "Create user"
// @security ApiKeyAuth
// @Success 201 {object} models.User "User registered successfully"
// @Failure 400 {object} schemas.ErrorResponse "Bad request"
// @Failure 401 {object} schemas.ErrorResponse "Error returned when the provided auth data is invalid"
// @Failure 403 {object} schemas.ErrorResponse "Error returned when auth data was not provided"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/users [post]
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	log.Println("RegisterUser handler invoked")

	var userReq schemas.CreateUserRequest

	// Decoding the request body into the CreateUserRequest structure
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		general_utils.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("Decoded to CreateUserRequest")

	// Request ralidation
	err = userReq.Validate()
	if err != nil {
		general_utils.ErrorResponse(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// Saving the user in the database & prepare user response
	userRes, err := users_repo.CreateUser(userReq)
	if err != nil {
		log.Println("Faced an error in users_repo.CreateUser")
		general_utils.HandleExceptionResponse(w, err)
		return
	}
	// Setting the status 201
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	// Sending the data of the created user to the client in JSON format
	err = handlers_utils.HandleJsonResponse(w, userRes)
	if err != nil {
		log.Println("Error while handling JSON response:", err)
	}
}
