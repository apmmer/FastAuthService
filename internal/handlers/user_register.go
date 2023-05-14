package handlers

import (
	"AuthService/internal/models"
	"AuthService/internal/repositories/user_repo"
	"AuthService/internal/schemas"
	"encoding/json"
	"log"
	"net/http"
)

// RegisterUser godoc
// @Summary Register new user
// @Description Register a new user with email, username, and password
// @Tags users
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

	// Validation of user data
	if userReq.Email == "" || userReq.Username == "" || userReq.Password == "" {
		http.Error(w, "Invalid user data", http.StatusBadRequest)
		return
	}

	user := &models.User{
		Username: userReq.Username,
		Email:    userReq.Email,
		Password: userReq.Password,
	}

	// Saving the user in the database
	newUser, err := user_repo.SaveUserToDB(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Prepare user response
	userRes := models.User{
		ID:        newUser.ID,
		Username:  newUser.Username,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
		DeletedAt: newUser.DeletedAt,
	}

	// Setting the status 201
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	// Sending the data of the created user to the client in JSON format
	json.NewEncoder(w).Encode(userRes)
}
