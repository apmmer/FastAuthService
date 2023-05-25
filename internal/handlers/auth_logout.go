package handlers

import (
	"AuthService/configs"
	"AuthService/internal/handlers/handlers_utils"
	"AuthService/internal/repositories/sessions_repo"
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
// @Failure 500 {object} schemas.ErrorResponse "Returns an error message if there is a server-side issue"
// @Router /api/logout [post]
func Logout(w http.ResponseWriter, r *http.Request) {
	// AccessToken validation
	log.Println("Logout: validating access token")
	userId, sessionToken, err := extractUidAndSessionToken(r)
	if err != nil {
		handlers_utils.HandleExceptionResponse(w, err)
		return
	}
	// here we will perform user session and cookies updation
	cookies, err := updateSessionAndCookies(sessionToken)
	if err != nil {
		handlers_utils.ErrorResponse(w, "Session is closed, expired or does not exist.", http.StatusUnauthorized)
		return
	}
	http.SetCookie(w, cookies)

	// prepare response
	responseMsg := fmt.Sprintf("Successfully logged out user with ID #%d", userId)
	err = handlers_utils.HandleJsonResponse(w, responseMsg)
	if err != nil {
		handlers_utils.HandleExceptionResponse(w, fmt.Errorf("Error while handling JSON response: %v", err))
	}
}

// updateSessionAndCookies invalidates a user session and resets the refresh token.
// It takes as input the session token string, and returns a pointer to an http.Cookie or an error.
func updateSessionAndCookies(sessionToken string) (*http.Cookie, error) {
	// Perform the session update operation. if error = session filters are invalid
	_, err := sessions_repo.UpdateSessions(
		&map[string]interface{}{
			"token":      sessionToken,
			"deleted_at": nil,
		},
		&map[string]interface{}{
			"deleted_at": time.Now(),
		},
	)
	if err != nil {
		return nil, err
	}

	// Create a new cookie
	cookies := http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		Path:     "/",
		HttpOnly: true,
		Secure:   configs.MainSettings.SecureCookies,
		SameSite: http.SameSiteStrictMode,
	}
	return &cookies, nil
}

// extractUidAndSessionToken extracts the user ID and session token from the request.
// It takes as input an http.Request, and returns the user ID, session token, or an error.
func extractUidAndSessionToken(r *http.Request) (int, string, error) {
	// Validate access token from request header
	accessClaims, refreshClaims, err := handlers_utils.ValidateAccessTokenHeader(r)
	if err != nil {
		return 0, "", err
	}
	// Convert user ID from string to integer
	userId, err := strconv.Atoi((*accessClaims)["Id"].(string))
	if err != nil {
		return 0, "", err
	}
	// Extract session token from refresh claims
	sessionToken := (*refreshClaims)["SessionToken"].(string)
	return userId, sessionToken, nil
}
