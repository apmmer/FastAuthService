package schemas

// LoginInput godoc
// @Schema LoginInput
// @Required email password
// @Property1 email string email format
// @Property2 password string mininum 7 characters required
type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,min=7"`
}

// TokenResponse godoc
// @Schema TokenResponse
// @Property1 access_token string The generated JWT token
// @Property2 access_expires int64 The unix timestamp when the token will expire
type TokenResponse struct {
	AccessToken   string `json:"access_token"`
	AccessExpires int64  `json:"access_expires"`
}
