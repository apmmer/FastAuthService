package handlers

import (
	"AuthService/internal/models"
	"AuthService/internal/repositories/user_repo"
	"encoding/json" // go get encoding/json
	"net/http"
)

func handleRegister(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Decoding the request body into the User structure
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validation of user data (additional checks can be added)
	if user.Email == "" || user.Username == "" || user.Password == "" {
		http.Error(w, "Invalid user data", http.StatusBadRequest)
		return
	}

	// Saving the user in the database
	newUser, err := user_repo.SaveUserToDB(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Setting the status 201
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	// Sending the data of the created user to the client in JSON format
	json.NewEncoder(w).Encode(newUser)
}
