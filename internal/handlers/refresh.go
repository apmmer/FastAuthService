package handlers

import (
	"AuthService/configs"
	"AuthService/internal/repositories/user_repo"
	"AuthService/internal/utils"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// RefreshTokens is a handler function for token refresh requests.
// @Summary Refresh JWT token
// @Description Use the refresh token to get a new access token
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {object} schemas.TokenResponse
// @Failure 400 {object} string "bad request"
// @Failure 401 {object} string "unauthorized"
// @Failure 403 {object} string "Auth data was not provided"
// @Failure 500 {object} string "Internal server error"
// @Router /api/refresh [post]
func RefreshTokens(w http.ResponseWriter, r *http.Request) {
	// Extract the refresh token from the request cookies
	c, err := r.Cookie("refresh_token")
	if err != nil {
		if err == http.ErrNoCookie {
			ErrorResponse(w, "refresh token not provided", http.StatusForbidden)
			return
		}
		HandleException(w, err)
		return
	}
	refreshToken := c.Value

	// Parse the refresh token
	claims, err := utils.ParseToken(refreshToken, configs.MainSettings.JwtRefreshSecret)
	if err != nil {
		ErrorResponse(w, "invalid refresh token", http.StatusUnauthorized)
		return
	}

	// Check if the refresh token has expired
	if claims.ExpiresAt < time.Now().Unix() {
		ErrorResponse(w, "expired refresh token", http.StatusUnauthorized)
		return
	}

	// Get the user associated with the refresh token
	userId, err := strconv.Atoi(claims.Id)
	if err != nil {
		HandleException(w, err)
		return
	}
	filters := make(map[string]interface{})
	filters["id"] = userId
	user, err := user_repo.GetUser(&filters)
	if err != nil {
		HandleException(w, err)
		return
	}

	// Generate a new access token
	accessToken, err := utils.GenerateAccessToken(user)
	if err != nil {
		HandleException(w, err)
		return
	}

	// Set Refresh cookies
	cookies, err := utils.GenerateRefreshCookies(user)
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
