package utils

import (
	"AuthService/configs"
	"AuthService/internal/models"
	"AuthService/internal/schemas"
	"errors"
	"fmt"
	"net/http"
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

// generates exactly JWT token with provided data
func GenerateJWT(user *models.User, expiresAt int64, issuer string, secret string) (string, error) {
	// Create the claims for the JWT: these are the pieces of data we want to include in the token.
	// In this case, we are including the standard claims (expiry, issuer) plus the ID of the user.
	claims := &jwt.StandardClaims{
		ExpiresAt: expiresAt,
		Issuer:    issuer,
		Id:        strconv.Itoa(int(user.ID)),
	}

	// Create a new JWT and include the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with our secret
	tokenString, err := token.SignedString([]byte(secret))

	// If there was an error in the previous step, return the error
	if err != nil {
		return "", err
	}

	// Return the TokenResponse struct and no error
	return tokenString, nil
}

// GenerateAccessToken creates a new Access JWT for the provided user.
// The token includes the user's ID and an expiration timestamp, and it is signed with a secret key.
// The function returns a TokenResponse struct, which includes the JWT itself and its expiration timestamp.
func GenerateAccessToken(user *models.User) (*schemas.TokenResponse, error) {
	expiresAt := time.Now().Add(time.Minute * time.Duration(configs.MainSettings.TokenLifeMinutes)).Unix()
	tokenString, err := GenerateJWT(user, expiresAt, configs.MainSettings.ServiceName, configs.MainSettings.JwtSecret)
	if err != nil {
		return nil, err
	}
	// Create TokenResponse using expiresAt
	tokenRes := schemas.TokenResponse{
		AccessToken:   tokenString,
		AccessExpires: expiresAt,
	}

	// Return the TokenResponse struct and no error
	return &tokenRes, nil
}

// GenerateRefreshToken generates a new JWT refresh token for a given user
func GenerateRefreshCookies(user *models.User) (http.Cookie, error) {
	// Set expiration time for refresh token (longer than access token)
	expiresAt := time.Now().Add(time.Minute * time.Duration(configs.MainSettings.RefreshTokenLifeMinutes)).Unix()
	tokenString, err := GenerateJWT(user, expiresAt, configs.MainSettings.ServiceName, configs.MainSettings.JwtRefreshSecret)
	if err != nil {
		return http.Cookie{}, err
	}

	// Create and return the refresh token cookie
	cookie := http.Cookie{
		Name:     "refresh_token",
		Value:    tokenString,
		Expires:  time.Unix(expiresAt, 0),
		HttpOnly: true,
		Secure:   configs.MainSettings.SecureCookies,
		SameSite: http.SameSiteStrictMode,
	}
	return cookie, nil
}

// ParseToken parses a JWT token and returns the claims.
// The function expects a valid token and a secret key. If the token is not valid or does not contain the expected claims, the function will return an error.
func ParseToken(tokenString string, secret string) (*jwt.StandardClaims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the token method conforms to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	// If there was an error in parsing the token, return the error
	if err != nil {
		return nil, err
	}

	// If the token is valid, return the claims
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}
