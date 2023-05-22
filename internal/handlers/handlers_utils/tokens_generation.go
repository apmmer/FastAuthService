package handlers_utils

import (
	"AuthService/configs"
	"AuthService/internal/exceptions"
	"AuthService/internal/schemas"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

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
func GenerateAccessToken(userId int, deviceInfo *schemas.DeviceInfo) (*schemas.TokenResponse, error) {
	expiresAt := time.Now().Add(time.Minute * time.Duration(configs.MainSettings.TokenLifeMinutes)).Unix()
	claims := map[string]interface{}{
		"Id":        strconv.Itoa(userId),
		"ExpiresAt": expiresAt,
		"Issuer":    configs.MainSettings.ServiceName,
		"IPAddress": deviceInfo.IPAddress,
		"UserAgent": deviceInfo.UserAgent,
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
func GenerateRefreshCookies(userId int, accessToken string, sessionToken string, expiresAt *time.Time) (*http.Cookie, error) {
	// Note: claims include generated access token
	claims := map[string]interface{}{
		"Id":           strconv.Itoa(userId),
		"ExpiresAt":    expiresAt.Unix(),
		"Issuer":       configs.MainSettings.ServiceName,
		"AccessToken":  accessToken,
		"SessionToken": sessionToken,
	}
	tokenString, err := GenerateJWT(&claims, configs.MainSettings.JwtRefreshSecret)
	if err != nil {
		return nil, err
	}

	// Create and return the refresh token cookie
	log.Printf("GenerateRefreshCookies: generated cookie with \nAccess token: %s\nCookie value %s", accessToken, tokenString)
	cookie := http.Cookie{
		Name:     "refresh_token",
		Value:    tokenString,
		Expires:  time.Unix(expiresAt.Unix(), 0),
		Path:     "/",
		HttpOnly: true,
		Secure:   configs.MainSettings.SecureCookies,
		SameSite: http.SameSiteStrictMode,
	}
	return &cookie, nil
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

// Extracts JWT from request Authorization header
func ExtractJWTFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", &exceptions.ErrNoAuthData{Message: "Missing 'Authorization' header"}
	}
	return authHeader, nil
}
