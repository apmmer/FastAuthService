package general_utils

import "golang.org/x/crypto/bcrypt"

// HashPassword takes a raw password and returns a bcrypt hashed password
func HashPassword(password string) (string, error) {
	// Convert password string to byte slice so it can be used with bcrypt
	pass := []byte(password)

	// Hash password using bcrypt's DefaultCost
	hashedPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		// If there's an error, return it
		return "", err
	}

	// Return hashed password as a string
	return string(hashedPassword), nil
}
