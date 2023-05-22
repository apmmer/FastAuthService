package handlers_utils

import (
	"AuthService/configs"
	"AuthService/internal/exceptions"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func ValidateRefreshTokenCookie(r *http.Request, headerAccessToken string) (*jwt.MapClaims, error) {
	c, err := r.Cookie("refresh_token")
	if err != nil {
		if err == http.ErrNoCookie {
			return nil, &exceptions.ErrNoAuthData{Message: "refresh token is not provided in cookies"}
		}
		return nil, err
	}
	cookieRefreshToken := c.Value
	log.Printf("ValidateRefreshTokenCookie: got cookies value: %s.", cookieRefreshToken)

	// Parse the refresh token
	refreshClaims, err := ParseToken(cookieRefreshToken, configs.MainSettings.JwtRefreshSecret)
	if err != nil {
		return nil, &exceptions.ErrUnauthorized{Message: "invalid refresh token"}
	}

	// Check if the refresh token has expired
	err = ValidateTokenExpiresAt(&refreshClaims)
	if err != nil {
		return nil, err
	}
	// Compare access tokens
	cookieAccessToken := refreshClaims["AccessToken"].(string)
	if cookieAccessToken != headerAccessToken {
		log.Printf("ValidateRefreshTokenCookie: got different access tokens: \nFrom cookie: %s\nFrom header: %s.", cookieAccessToken, headerAccessToken)
		return nil, &exceptions.ErrUnauthorized{Message: "a pair of access tokens are not suitable for each other"}
	}
	return &refreshClaims, nil
}

func ValidateTokenExpiresAt(claims *jwt.MapClaims) error {
	if expiresAtFloat, ok := (*claims)["ExpiresAt"].(float64); ok {
		expiresAt := int64(expiresAtFloat)
		if expiresAt < time.Now().Unix() {
			return &exceptions.ErrUnauthorized{Message: "expired refresh token"}
		}
	} else {
		// Log the unexpected type and value
		log.Printf("Unexpected type for ExpiresAt claim: %T. Value: %v", (*claims)["ExpiresAt"], (*claims)["ExpiresAt"])
		return &exceptions.ErrUnauthorized{Message: "invalid ExpiresAt claim in refresh token"}
	}
	return nil
}

func ValidateTokenDeviceInfo(r *http.Request, accessClaims *jwt.MapClaims) error {
	deviceInfo := GetDeviceInfo(r)
	if deviceInfo.IPAddress != (*accessClaims)["IPAddress"].(string) && deviceInfo.UserAgent != (*accessClaims)["UserAgent"].(string) {
		log.Printf("ValidateTokenDeviceInfo error: both IPAddress and UserAgent are different.")
		return &exceptions.ErrUnauthorized{Message: "invalid access token"}
	}
	return nil
}

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
		return nil, nil, &exceptions.ErrUnauthorized{Message: "invalid access token"}
	}
	err = ValidateTokenExpiresAt(&accessClaims)
	if err != nil {
		return nil, nil, err
	}
	err = ValidateTokenDeviceInfo(r, &accessClaims)
	if err != nil {
		return nil, nil, err
	}
	return &accessClaims, refreshClaims, nil
}
