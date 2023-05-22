package handlers

import (
	"AuthService/configs"
	"AuthService/internal/exceptions"
	"AuthService/internal/general_utils"
	"AuthService/internal/handlers/handlers_utils"
	"AuthService/internal/repositories/sessions_repo"
	"AuthService/internal/repositories/user_repo"
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
		handlers_utils.HandleException(w, err)
		return
	}

	accessToken, cookies, err := generateUserSessionAndTokens(r, userId)
	if err != nil {
		handlers_utils.HandleException(w, err)
		return
	}
	// Set refresh token in cookies
	http.SetCookie(w, cookies)

	// Create response with token
	err = handlers_utils.HandleJsonResponse(w, accessToken)
	if err != nil {
		handlers_utils.HandleException(w, fmt.Errorf("Error while handling JSON response: %v", err))
	}
}

func processUser(r *http.Request) (int, error) {
	// Decode body
	var input schemas.LoginInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		return 0, &exceptions.ErrInputError{Message: fmt.Sprint("Invalid input")}
	}

	// Getting user from db
	filters := make(map[string]interface{})
	filters["email"] = input.Email
	filters["deleted_at"] = nil
	user, err := user_repo.GetUser(&filters)
	if err != nil {
		return 0, err
	}

	// Checking password
	isValid := general_utils.CheckHash(input.Password, user.Password)
	if !isValid {
		return 0, &exceptions.ErrUnauthorized{Message: fmt.Sprint("Invalid username or password")}
	}
	// return int UID
	return int((*user).ID), nil
}

func generateUserSessionAndTokens(r *http.Request, userId int) (*schemas.TokenResponse, *http.Cookie, error) {
	// getting device info
	deviceInfo := handlers_utils.GetDeviceInfo(r)
	log.Printf("deviceInfo \n	IP: %s\n	UserAgent: %s", deviceInfo.IPAddress, deviceInfo.UserAgent)

	// Генерируем токен доступа
	accessToken, err := handlers_utils.GenerateAccessToken(userId, &deviceInfo)
	if err != nil {
		return nil, nil, err
	}

	sessionToken, err := handlers_utils.GenerateSessionToken(&deviceInfo, configs.MainSettings.SessionSecret)
	// generate the expiration date for both session and refreshToken
	expiresAt := time.Now().Add(time.Minute * time.Duration(configs.MainSettings.RefreshTokenLifeMinutes))
	_, err = sessions_repo.CreateSession(uint(userId), sessionToken, &expiresAt)
	if err != nil {
		return nil, nil, err
	}
	// Set Refresh cookies
	cookies, err := handlers_utils.GenerateRefreshCookies(userId, accessToken.AccessToken, sessionToken, &expiresAt)
	if err != nil {
		return nil, nil, err
	}
	return accessToken, cookies, nil
}
