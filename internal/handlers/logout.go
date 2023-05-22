package handlers

import (
	"AuthService/configs"
	"AuthService/internal/repositories/sessions_repo"
	"AuthService/internal/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Logout is a handler function for user logout.
// @Summary Logs out a user
// @Description Restricts access to services for the active client until a new login occurs.
// @Tags Auth
// @Accept  json
// @Produce  json
// @security JWTAuth
// @security ApiKeyAuth
// @Success 200 {object} string "Successfully logged out user with ID #id"
// @Failure 401 {object} schemas.ErrorResponse "Error returned when the provided auth data is invalid"
// @Failure 403 {object} schemas.ErrorResponse "Error returned when auth data was not provided"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/logout [post]
func Logout(w http.ResponseWriter, r *http.Request) {
	// AccessToken validation
	log.Println("Logout: validating access token")
	accessClaims, refreshClaims, err := utils.ValidateAccessToken(r)
	if err != nil {
		HandleException(w, err)
		return
	}
	userId, err := strconv.Atoi((*accessClaims)["Id"].(string))
	log.Printf("Got userId = %d", userId)
	if err != nil {
		HandleException(w, err)
		return
	}
	refreshToken := (*refreshClaims)["SessionToken"].(string)
	// here we will perform user session update, but next time
	_, err = sessions_repo.UpdateSessions(
		&map[string]interface{}{
			"token":      refreshToken,
			"deleted_at": nil,
		},
		&map[string]interface{}{
			"deleted_at": time.Now(),
		},
	)
	if err != nil {
		ErrorResponse(w, "Session is closed, expired or not exists.", http.StatusUnauthorized)
		return
	}

	// here we delete cookies:
	cookies := http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		Path:     "/",
		HttpOnly: true,
		Secure:   configs.MainSettings.SecureCookies,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, &cookies)

	// prepare response
	responseMsg := fmt.Sprintf("Successfully logged out user with ID #%d", userId)
	err = HandleJsonResponse(w, responseMsg)
	if err != nil {
		HandleException(w, fmt.Errorf("Error while handling JSON response: %v", err))
	}
}
