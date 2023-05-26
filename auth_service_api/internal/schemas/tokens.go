package schemas

// TokenResponse godoc
// @Schema TokenResponse
// @Property1 access_token string The generated JWT token
// @Property2 access_expires int64 The unix timestamp when the token will expire
type TokenResponse struct {
	AccessToken   string `json:"access_token"`
	AccessExpires int64  `json:"access_expires"`
}

// RefreshToken godoc
// @Schema RefreshToken
// @Property1 refresh_token string The generated JWT token
// @Property2 refresh_expires int64 The unix timestamp when the token will expire
type RefreshToken struct {
	RefreshToken   string `json:"refresh_token"`
	RefreshExpires int64  `json:"refresh_expires"`
}
