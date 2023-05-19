package handlers

import (
	"AuthService/internal/repositories/user_repo"
	"AuthService/internal/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// RefreshTokens is a handler function for token refresh requests.
// @Summary Refresh tokens
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
	headerAccessToken, err := utils.ExtractJWT(r)
	if err != nil {
		HandleException(w, err)
		return
	}
	refreshClaims, err := utils.ValidateRefreshTokenCookie(r, headerAccessToken)
	if err != nil {
		HandleException(w, err)
		return
	}
	ValidateSession or UpdateSession directly

	// Get the user associated with the refresh token
	userId, err := strconv.Atoi((*refreshClaims)["Id"].(string))
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

	// getting device info
	deviceInfo := utils.GetDeviceInfo(r)
	log.Printf("deviceInfo \n	IP: %s\n	UserAgent: %s", deviceInfo.IPAddress, deviceInfo.UserAgent)

	// Generate a new access token
	accessToken, err := utils.GenerateAccessToken(user, &deviceInfo)
	if err != nil {
		HandleException(w, err)
		return
	}

	// generate and set new Refresh cookies
	cookies, err := utils.GenerateRefreshCookies(user, accessToken.AccessToken, (*refreshClaims)["SessionToken"].(string))
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
