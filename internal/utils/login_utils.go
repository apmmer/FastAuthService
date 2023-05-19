package utils

import (
	"AuthService/configs"
	"AuthService/internal/exceptions"
	"AuthService/internal/models"
	"AuthService/internal/schemas"
	"errors"
	"fmt"
	"log"
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
func GenerateJWT(data *map[string]interface{}, secret string) (string, error) {
	// Create the claims for the JWT: these are the pieces of data we want to include in the token.
	// In this case, we are including the standard claims (expiry, issuer) plus the ID of the user.
	claims := jwt.MapClaims(*data)

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
	claims := map[string]interface{}{
		"Id":        strconv.Itoa(int(user.ID)),
		"ExpiresAt": expiresAt,
		"Issuer":    configs.MainSettings.ServiceName,
	}
	tokenString, err := GenerateJWT(&claims, configs.MainSettings.JwtSecret)
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
func GenerateRefreshCookies(user *models.User, accessToken string) (http.Cookie, error) {
	// Set expiration time for refresh token (longer than access token)
	expiresAt := time.Now().Add(time.Minute * time.Duration(configs.MainSettings.RefreshTokenLifeMinutes)).Unix()
	// Note: claims include generated access token
	claims := map[string]interface{}{
		"Id":          strconv.Itoa(int(user.ID)),
		"ExpiresAt":   expiresAt,
		"Issuer":      configs.MainSettings.ServiceName,
		"AccessToken": accessToken,
	}
	tokenString, err := GenerateJWT(&claims, configs.MainSettings.JwtRefreshSecret)
	if err != nil {
		return http.Cookie{}, err
	}

	// Create and return the refresh token cookie
	log.Printf("GenerateRefreshCookies: generated cookie with \nAccess token: %s\nCookie value %s", accessToken, tokenString)
	cookie := http.Cookie{
		Name:     "refresh_token",
		Value:    tokenString,
		Expires:  time.Unix(expiresAt, 0),
		Path:     "/",
		HttpOnly: true,
		Secure:   configs.MainSettings.SecureCookies,
		SameSite: http.SameSiteStrictMode,
	}
	return cookie, nil
}

// ParseToken parses a JWT token and returns the claims.
// The function expects a valid token and a secret key.
// If the token is not valid or does not contain the expected claims,
// the function will return an error.
func ParseToken(tokenString string, secret string) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
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
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

func ExtractJWT(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", &exceptions.ErrNoAuthData{Message: "Missing Authorization header"}
	}
	return authHeader, nil
}

func ValidateRefreshTokenCookie(r *http.Request, headerAccessToken string) (*jwt.MapClaims, error) {
	c, err := r.Cookie("refresh_token")
	if err != nil {
		if err == http.ErrNoCookie {
			return nil, &exceptions.ErrNoAuthData{Message: "refresh token is not provided in cookies"}
		}
		return nil, err
	}
	cookieRefreshToken := c.Value
	log.Printf("ValidateRefreshTokenCookie: got cookies value: %s.", cookieRefreshToken)

	// Parse the refresh token
	refreshClaims, err := ParseToken(cookieRefreshToken, configs.MainSettings.JwtRefreshSecret)
	if err != nil {
		return nil, &exceptions.ErrUnauthorized{Message: "invalid refresh token"}
	}

	// Check if the refresh token has expired
	err = ValidateTokenExpiresAt(&refreshClaims)
	if err != nil {
		return nil, err
	}
	// Compare access tokens
	cookieAccessToken := refreshClaims["AccessToken"].(string)
	if cookieAccessToken != headerAccessToken {
		log.Printf("ValidateRefreshTokenCookie: got different access tokens: \nFrom cookie: %s\nFrom header: %s.", cookieAccessToken, headerAccessToken)
		return nil, &exceptions.ErrUnauthorized{Message: "a pair of access tokens are not suitable for each other"}
	}
	return &refreshClaims, nil
}

func ValidateTokenExpiresAt(claims *jwt.MapClaims) error {
	if expiresAtFloat, ok := (*claims)["ExpiresAt"].(float64); ok {
		expiresAt := int64(expiresAtFloat)
		if expiresAt < time.Now().Unix() {
			return &exceptions.ErrUnauthorized{Message: "expired refresh token"}
		}
	} else {
		// Log the unexpected type and value
		log.Printf("Unexpected type for ExpiresAt claim: %T. Value: %v", (*claims)["ExpiresAt"], (*claims)["ExpiresAt"])
		return &exceptions.ErrUnauthorized{Message: "invalid ExpiresAt claim in refresh token"}
	}
	return nil
}

func ValidateAccessToken(r *http.Request) (*jwt.MapClaims, error) {
	headerAccessToken, err := ExtractJWT(r)
	if err != nil {
		return nil, err
	}
	_, err = ValidateRefreshTokenCookie(r, headerAccessToken)
	if err != nil {
		return nil, err
	}
	accessClaims, err := ParseToken(headerAccessToken, configs.MainSettings.JwtSecret)
	if err != nil {
		return nil, &exceptions.ErrUnauthorized{Message: "invalid access token"}
	}
	err = ValidateTokenExpiresAt(&accessClaims)
	return &accessClaims, nil
}
