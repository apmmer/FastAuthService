package handlers

import (
	"AuthService/internal/handlers/handlers_utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// ValidateAccess is a handler function to ensure that corrent client has access.
// @Summary Checks access for secure endpoints.
// @Description Checks if provided authorization data is valid.
// @Tags Auth
// @Accept  json
// @Produce  json
// @security JWTAuth
// @security ApiKeyAuth
// @Success 200 {object} string "Authorization data is valid for user with ID #id"
// @Failure 401 {object} schemas.ErrorResponse "Error returned when the provided auth data is invalid"
// @Failure 403 {object} schemas.ErrorResponse "Error returned when auth data was not provided"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/validate [post]
func ValidateAccess(w http.ResponseWriter, r *http.Request) {
	log.Println("ValidateAccess: validating access token")
	accessClaims, _, err := handlers_utils.ValidateAccessTokenHeader(r)
	if err != nil {
		handlers_utils.HandleException(w, err)
		return
	}
	userId, err := strconv.Atoi((*accessClaims)["Id"].(string))
	log.Printf("Got userId = %d", userId)
	if err != nil {
		handlers_utils.HandleException(w, err)
		return
	}
	responseMsg := fmt.Sprintf("Authorization data is valid for user with ID #%d", userId)
	err = handlers_utils.HandleJsonResponse(w, responseMsg)
	if err != nil {
		handlers_utils.HandleException(w, fmt.Errorf("Error while handling JSON response: %v", err))
	}
}
