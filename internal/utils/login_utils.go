package utils

import (
	"AuthService/configs"
	"AuthService/internal/models"
	"AuthService/internal/schemas"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

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

// CheckPasswordHash compares a provided password with a hashed password and returns a boolean result.
// It uses the bcrypt library's CompareHashAndPassword function to do this.
// The provided password is hashed and then compared to the stored hashed password.
// If the two hashed passwords match, the function returns true. Otherwise, it returns false.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateToken creates a new JWT for the provided user.
// The token includes the user's ID and an expiration timestamp, and it is signed with a secret key.
// The function returns a TokenResponse struct, which includes the JWT itself and its expiration timestamp.
func GenerateToken(user *models.User) (*schemas.TokenResponse, error) {

	// Calculate the expiration timestamp for the token: current time + the configured lifespan of the token
	expiresAt := time.Now().Add(time.Minute * time.Duration(configs.MainSettings.TokenLifeMinutes)).Unix()

	// Create the claims for the JWT: these are the pieces of data we want to include in the token.
	// In this case, we are including the standard claims (expiry, issuer) plus the ID of the user.
	claims := &jwt.StandardClaims{
		ExpiresAt: expiresAt,
		Issuer:    "AuthService",
		Id:        strconv.Itoa(int(user.ID)),
	}

	// Create a new JWT and include the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with our secret
	tokenString, err := token.SignedString([]byte(configs.MainSettings.JwtSecret))

	// If there was an error in the previous step, return the error
	if err != nil {
		return nil, err
	}

	// Create a new TokenResponse struct to hold the token and its expiration timestamp
	tokenRes := schemas.TokenResponse{
		AccessToken:   tokenString,
		AccessExpires: expiresAt,
	}

	// Return the TokenResponse struct and no error
	return &tokenRes, nil
}
