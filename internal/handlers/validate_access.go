package handlers

import (
	"AuthService/internal/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// ValidateAccess is a handler function to ensure that corrent client has access.
// @Summary Validates access for current logged in client
// @Description Use the refresh token to get a new access token and to set new refresh token in cookies.
// @Tags Auth
// @Accept  json
// @Produce  json
// @security JWTAuth
// @security ApiKeyAuth
// @Success 200 {object} string
// @Failure 401 {object} schemas.ErrorResponse "Error returned when the provided auth data is invalid"
// @Failure 403 {object} schemas.ErrorResponse "Error returned when auth data was not provided"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/validate [post]
func ValidateAccess(w http.ResponseWriter, r *http.Request) {
	log.Println("ValidateAccess: validating access token")
	accessClaims, err := utils.ValidateAccessToken(r)
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
	responseMsg := fmt.Sprintf("Granted access for uid #%d", userId)
	err = HandleJsonResponse(w, responseMsg)
	if err != nil {
		HandleException(w, fmt.Errorf("Error while handling JSON response: %v", err))
	}
}
