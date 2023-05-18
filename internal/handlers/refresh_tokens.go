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

func ValidateRefreshTokenCookie(r *http.Request) (*jwt.StandardClaims, error) {
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
	if claims.ExpiresAt < time.Now().Unix() {
		return nil, &exceptions.ErrUnauthorized{Message: "expired refresh token"}
	}
	return claims, nil
}

// RefreshTokens is a handler function for token refresh requests.
// @Summary Refresh JWT token
// @Description Use the refresh token to get a new access token and to set new refresh token in cookies.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {object} schemas.TokenResponse
// @Failure 401 {object} schemas.ErrorResponse "Error raturned when the provided auth data is invalid"
// @Failure 403 {object} schemas.ErrorResponse "Error raturned when auth data was not provided"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/refresh [post]
func RefreshTokens(w http.ResponseWriter, r *http.Request) {
	log.Println("Validating token")
	// Extract the refresh token from the request cookies
	claims, err := ValidateRefreshTokenCookie(r)
	if err != nil {
		HandleException(w, err)
		return
	}

	// Get the user associated with the refresh token
	userId, err := strconv.Atoi(claims.Id)
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

	// Set Refresh cookies
	cookies, err := utils.GenerateRefreshCookies(user)
	if err != nil {
		HandleException(w, err)
		return
	}
	http.SetCookie(w, &cookies)

	// Generate a new access token
	accessToken, err := utils.GenerateAccessToken(user)
	if err != nil {
		HandleException(w, err)
		return
	}
	// Return the new access token
	err = HandleJsonResponse(w, accessToken)
	if err != nil {
		HandleException(w, fmt.Errorf("Error while handling JSON response: %v", err))
	}
}
