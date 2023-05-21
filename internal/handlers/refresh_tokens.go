package handlers

import (
	"AuthService/configs"
	"AuthService/internal/handlers/handlers_utils"
	"AuthService/internal/repositories/sessions_repo"
	"AuthService/internal/repositories/user_repo"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
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
	headerAccessToken, err := handlers_utils.ExtractJWT(r)
	if err != nil {
		handlers_utils.HandleException(w, err)
		return
	}
	refreshClaims, err := handlers_utils.ValidateRefreshTokenCookie(r, headerAccessToken)
	if err != nil {
		handlers_utils.HandleException(w, err)
		return
	}

	// Get the user associated with the refresh token
	userId, err := strconv.Atoi((*refreshClaims)["Id"].(string))
	log.Printf("Got userId = %d", userId)
	if err != nil {
		handlers_utils.HandleException(w, err)
		return
	}

	// Retrieve user. Currently it is not necessary, but in future - yes.
	user, err := user_repo.GetActiveUserById(userId)
	if err != nil {
		handlers_utils.HandleException(w, err)
		return
	}
	log.Printf("Got user = %v", user)

	// Update session directly
	// Session expires in the same time as new refresh token
	expiresAt := time.Now().Add(time.Minute * time.Duration(configs.MainSettings.RefreshTokenLifeMinutes))
	sessions, err := sessions_repo.UpdateSessions(
		&map[string]interface{}{
			"token":      (*refreshClaims)["SessionToken"].(string),
			"deleted_at": nil,
			"user_id":    (*user).ID,
		},
		&map[string]interface{}{
			"expires_at": expiresAt,
		},
	)
	if err != nil {
		handlers_utils.ErrorResponse(w, "Session is closed, expired or not exists.", http.StatusUnauthorized)
		return
	}
	if len(*sessions) != 1 {
		handlers_utils.ErrorResponse(w, "Found unexpected user session, please log in again.", http.StatusInternalServerError)
	}

	// getting device info
	deviceInfo := handlers_utils.GetDeviceInfo(r)
	log.Printf("deviceInfo \n	IP: %s\n	UserAgent: %s", deviceInfo.IPAddress, deviceInfo.UserAgent)

	// Generate a new access token
	accessToken, err := handlers_utils.GenerateAccessToken(user, &deviceInfo)
	if err != nil {
		handlers_utils.HandleException(w, err)
		return
	}

	// generate and set new Refresh cookies using old session token
	cookies, err := handlers_utils.GenerateRefreshCookies(user, accessToken.AccessToken, (*refreshClaims)["SessionToken"].(string), &expiresAt)
	if err != nil {
		handlers_utils.HandleException(w, err)
		return
	}
	http.SetCookie(w, &cookies)

	// Return the new access token
	err = handlers_utils.HandleJsonResponse(w, accessToken)
	if err != nil {
		handlers_utils.HandleException(w, fmt.Errorf("Error while handling JSON response: %v", err))
	}
}
