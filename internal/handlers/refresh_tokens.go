package handlers

import (
	"AuthService/configs"
	"AuthService/internal/exceptions"
	"AuthService/internal/repositories/user_repo"
	"AuthService/internal/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// RefreshTokens is a handler function for token refresh requests.
// @Summary Refresh JWT token
// @Description Use the refresh token to get a new access token and to set new refresh token in cookies.
// @Tags Auth
// @Accept  json
// @Produce  json
// @security JWTAuth
// @security ApiKeyAuth
// @Success 200 {object} schemas.TokenResponse
// @Failure 401 {object} schemas.ErrorResponse "Error returned when the provided auth data is invalid"
// @Failure 403 {object} schemas.ErrorResponse "Error returned when auth data was not provided"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/refresh [post]
func RefreshTokens(w http.ResponseWriter, r *http.Request) {
	log.Println("RefreshTokens: validating tokens")
	// Extract the refresh token from the request cookies
	claims, err := ValidateRefreshTokenCookie(r)
	if err != nil {
		HandleException(w, err)
		return
	}

	// Get the user associated with the refresh token
	userId, err := strconv.Atoi((*claims)["Id"].(string))
	log.Printf("Got userId = %d", userId)
	if err != nil {
		HandleException(w, err)
		return
	}
	// Retrieve user. Currently it is not necessary, but in future - yes.
	user, err := user_repo.GetActiveUserById(userId)
	if err != nil {
		HandleException(w, err)
		return
	}
	log.Printf("Got user = %v", user)

	// Generate a new access token
	accessToken, err := utils.GenerateAccessToken(user)
	if err != nil {
		HandleException(w, err)
		return
	}

	// generate and set new Refresh cookies
	cookies, err := utils.GenerateRefreshCookies(user, accessToken.AccessToken)
	if err != nil {
		HandleException(w, err)
		return
	}
	http.SetCookie(w, &cookies)

	// Return the new access token
	err = HandleJsonResponse(w, accessToken)
	if err != nil {
		HandleException(w, fmt.Errorf("Error while handling JSON response: %v", err))
	}
}

func ExtractJWT(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", &exceptions.ErrNoAuthData{Message: "Missing Authorization header"}
	}
	return authHeader, nil
}

func ValidateRefreshTokenCookie(r *http.Request) (*jwt.MapClaims, error) {
	c, err := r.Cookie("refresh_token")
	if err != nil {
		if err == http.ErrNoCookie {
			return nil, &exceptions.ErrNoAuthData{Message: "refresh token is not provided in cookies"}
		}
		return nil, err
	}
	refreshToken := c.Value

	// Parse the refresh token
	claims, err := utils.ParseToken(refreshToken, configs.MainSettings.JwtRefreshSecret)
	if err != nil {
		return nil, &exceptions.ErrUnauthorized{Message: "invalid refresh token"}
	}

	// Check if the refresh token has expired
	if expiresAtFloat, ok := claims["ExpiresAt"].(float64); ok {
		expiresAt := int64(expiresAtFloat)
		if expiresAt < time.Now().Unix() {
			return nil, &exceptions.ErrUnauthorized{Message: "expired refresh token"}
		}
	} else {
		// Log the unexpected type and value
		log.Printf("Unexpected type for ExpiresAt claim: %T. Value: %v", claims["ExpiresAt"], claims["ExpiresAt"])
		return nil, &exceptions.ErrUnauthorized{Message: "invalid ExpiresAt claim in refresh token"}
	}
	// compare access tokens
	cookieAccessToken := claims["AccessToken"].(string)
	headerAccessToken, err := ExtractJWT(r)
	if err != nil {
		return nil, err
	}
	if cookieAccessToken != headerAccessToken {
		return nil, &exceptions.ErrUnauthorized{Message: "a pair of access tokens are not suitable for each other"}
	}
	return &claims, nil
}
