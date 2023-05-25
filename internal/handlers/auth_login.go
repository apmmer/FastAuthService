package handlers

import (
	"AuthService/configs"
	"AuthService/internal/exceptions"
	"AuthService/internal/general_utils"
	"AuthService/internal/handlers/handlers_utils"
	"AuthService/internal/repositories/sessions_repo"
	"AuthService/internal/repositories/users_repo"
	"AuthService/internal/schemas"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Login godoc
// @Summary Logs in a user
// @Description Authenticates a user using email and password, and generates a new JWT. Also sets refresh token in cookies.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param InputBody body schemas.LoginInput true "The email and password of the user"
// @security ApiKeyAuth
// @Success 200 {object} schemas.TokenResponse "Returns a struct with the JWT and its expiration timestamp"
// @Failure 400 {object} schemas.ErrorResponse "Returns an error message if the request body cannot be parsed"
// @Failure 401 {object} schemas.ErrorResponse "Error returned when the provided auth data is invalid"
// @Failure 403 {object} schemas.ErrorResponse "Error returned when auth data was not provided"
// @Failure 500 {object} schemas.ErrorResponse "Returns an error message if there is a server-side issue"
// @Router /api/login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Got request to login.")

	userId, err := processUser(r)
	if err != nil {
		handlers_utils.HandleExceptionResponse(w, err)
		return
	}

	accessToken, cookies, err := generateUserSessionAndTokens(r, userId)
	if err != nil {
		handlers_utils.HandleExceptionResponse(w, err)
		return
	}
	// Set refresh token in cookies
	http.SetCookie(w, cookies)

	// Create response with token
	err = handlers_utils.HandleJsonResponse(w, accessToken)
	if err != nil {
		handlers_utils.HandleExceptionResponse(w, fmt.Errorf("Error while handling JSON response: %v", err))
	}
}

func processUser(r *http.Request) (int, error) {
	// Decode body
	var input schemas.LoginInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		return 0, exceptions.MakeInvalidEntityError("can not parse schemas.LoginInput")
	}

	// Getting user from db
	filters := make(map[string]interface{})
	filters["email"] = input.Email
	filters["deleted_at"] = nil
	user, err := users_repo.GetUser(&filters)
	if err != nil {
		return 0, err
	}

	// Checking password
	isValid := general_utils.CheckHash(input.Password, user.Password)
	if !isValid {
		return 0, exceptions.MakeUnauthorizedError("Invalid username or password")
	}
	// return int UID
	return int((*user).ID), nil
}

func generateUserSessionAndTokens(r *http.Request, userId int) (*schemas.TokenResponse, *http.Cookie, error) {
	// getting device info
	deviceInfo := handlers_utils.GetDeviceInfo(r)
	log.Printf("deviceInfo \n	IP: %s\n	UserAgent: %s", deviceInfo.IPAddress, deviceInfo.UserAgent)

	// Generate access token
	accessToken, err := handlers_utils.GenerateAccessToken(userId, &deviceInfo)
	if err != nil {
		return nil, nil, err
	}
	// Generate a session token for new session
	sessionToken, err := generateSessionToken(&deviceInfo, configs.MainSettings.SessionSecret)
	if err != nil {
		// hide origin message and log it.
		log.Printf("Can not create session token -> %v", err)
		return nil, nil, exceptions.MakeInternalError("Can not create session token.")
	}
	// generate the expiration date for both session and refreshToken
	expiresAt := time.Now().Add(time.Minute * time.Duration(configs.MainSettings.RefreshTokenLifeMinutes))
	// create session
	_, err = sessions_repo.CreateSession(uint(userId), sessionToken, &expiresAt)
	if err != nil {
		return nil, nil, err
	}
	// Get refresh token cookies
	cookies, err := handlers_utils.GenerateRefreshCookies(userId, accessToken.AccessToken, sessionToken, &expiresAt)
	if err != nil {
		return nil, nil, err
	}
	return accessToken, cookies, nil
}

// GenerateSessionToken creates an encrypted string that will be a token for the session
func generateSessionToken(deviceInfo *schemas.DeviceInfo, secretKey string) (string, error) {
	log.Println("Started session token encryption.")
	data := deviceInfo.IPAddress + "|" + deviceInfo.UserAgent
	cipherText, err := handlers_utils.CipherString(data, secretKey)
	if err != nil {
		log.Printf("Can not create session token -> %v", err)
		return "", exceptions.MakeInternalError("Can not create session token.")
	}
	return cipherText, nil
}
