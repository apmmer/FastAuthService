package handlers

import (
	"AuthService/internal/repositories/user_repo"
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
// @Success 201 {object} models.User "User registered successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/users [post]
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	log.Println("RegisterUser handler invoked")

	var userReq schemas.CreateUserRequest

	// Decoding the request body into the CreateUserRequest structure
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("Decoded to CreateUserRequest")

	// Validation
	err = userReq.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// Saving the user in the database & prepare user response
	userRes, err := user_repo.CreateUser(userReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Setting the status 201
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	// Sending the data of the created user to the client in JSON format
	json.NewEncoder(w).Encode(userRes)
}
