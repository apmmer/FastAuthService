package general_utils

import "golang.org/x/crypto/bcrypt"

// GetHash takes a raw string and returns a bcrypt hashed string
func GetHash(someString string) (string, error) {
	// Convert password string to byte slice so it can be used with bcrypt
	byteString := []byte(someString)

	// Hash password using bcrypt's DefaultCost
	hashedPassword, err := bcrypt.GenerateFromPassword(byteString, bcrypt.DefaultCost)
	if err != nil {
		// If there's an error, return it
		return "", err
	}

	// Return hashed password as a string
	return string(hashedPassword), nil
}

// CheckHash compares a provided string with a hashed string and returns a boolean result.
// It uses the bcrypt library's CompareHashAndPassword function to do this.
// The provided string is hashed and then compared to the stored hashed string.
// If the two hashed strings match, the function returns true. Otherwise false.
func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
