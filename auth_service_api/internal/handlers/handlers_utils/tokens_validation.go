package handlers_utils

import (
	"auth_service_api/configs"
	"auth_service_api/internal/exceptions"
	"auth_service_api/internal/general_utils"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// ValidateRefreshTokenCookie extracts refresh_token cookie and validates it's claims.
// Returns:
//
//	refresh token claims (*jwt.MapClaims), error
func ValidateRefreshTokenCookie(r *http.Request, headerAccessToken string) (*jwt.MapClaims, error) {
	c, err := r.Cookie("refresh_token")
	if err != nil {
		if err == http.ErrNoCookie {
			return nil, exceptions.MakeNoAuthDataError("refresh token is not provided in cookies.")
		}
		return nil, err
	}
	cookieRefreshToken := c.Value
	log.Printf("ValidateRefreshTokenCookie: got cookies value: %s.", cookieRefreshToken)

	// Parse the refresh token
	refreshClaims, err := ParseToken(cookieRefreshToken, configs.MainSettings.JwtRefreshSecret)
	if err != nil {
		err = general_utils.UpdateException("can not parse refresh token", err)
		return nil, exceptions.MakeUnauthorizedError(err.Error())
	}

	// Check if the refresh token has expired
	err = ValidateTokenExpiresAt(&refreshClaims)
	if err != nil {
		return nil, general_utils.UpdateException("Invalid refresh token", err)
	}
	// Compare access tokens
	cookieAccessToken := refreshClaims["AccessToken"].(string)
	if cookieAccessToken != headerAccessToken {
		log.Printf("ValidateRefreshTokenCookie: got different access tokens: \nFrom cookie: %s\nFrom header: %s.", cookieAccessToken, headerAccessToken)
		return nil, exceptions.MakeUnauthorizedError("a pair of access tokens are not suitable for each other.")
	}
	return &refreshClaims, nil
}

// Takes 'ExpiresAt' from provided claims and validates it.
// Note: Claims must contain 'ExpiresAt'.
func ValidateTokenExpiresAt(claims *jwt.MapClaims) error {
	if expiresAtFloat, ok := (*claims)["ExpiresAt"].(float64); ok {
		expiresAt := int64(expiresAtFloat)
		if expiresAt < time.Now().Unix() {
			return exceptions.MakeUnauthorizedError("token is expired.")
		}
	} else {
		// Log the unexpected type and value
		log.Printf("Unexpected type for ExpiresAt claim: %T. Value: %v", (*claims)["ExpiresAt"], (*claims)["ExpiresAt"])
		return exceptions.MakeUnauthorizedError("invalid ExpiresAt claim.")
	}
	return nil
}

func ValidateTokenDeviceInfo(r *http.Request, accessClaims *jwt.MapClaims) error {
	deviceInfo := GetDeviceInfo(r)
	if deviceInfo.IPAddress != (*accessClaims)["IPAddress"].(string) && deviceInfo.UserAgent != (*accessClaims)["UserAgent"].(string) {
		log.Printf("ValidateTokenDeviceInfo error: both IPAddress and UserAgent are different.")
		return exceptions.MakeUnauthorizedError("invalid access token")
	}
	return nil
}

// Extracts access token from header and validates it.
// As during this process claims must be extracted,
// returns claims for both tokens: access and refresh.
// Returns:
//
//	accessClaims (*jwt.MapClaims), refreshClaims (*jwt.MapClaims), error
func ValidateAccessTokenHeader(r *http.Request) (*jwt.MapClaims, *jwt.MapClaims, error) {
	headerAccessToken, err := ExtractJWTFromHeader(r)
	if err != nil {
		return nil, nil, err
	}
	refreshClaims, err := ValidateRefreshTokenCookie(r, headerAccessToken)
	if err != nil {
		return nil, nil, err
	}
	accessClaims, err := ParseToken(headerAccessToken, configs.MainSettings.JwtSecret)
	if err != nil {
		err = general_utils.UpdateException("can not parse access token", err)
		return nil, nil, exceptions.MakeUnauthorizedError(err.Error())
	}
	err = ValidateTokenExpiresAt(&accessClaims)
	if err != nil {
		return nil, nil, general_utils.UpdateException("invalid access token", err)
	}
	err = ValidateTokenDeviceInfo(r, &accessClaims)
	if err != nil {
		return nil, nil, err
	}
	return &accessClaims, refreshClaims, nil
}
