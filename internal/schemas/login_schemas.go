package schemas

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,min=7"`
}

type TokenResponse struct {
	AccessToken   string `json:"access_token"`
	AccessExpires int64  `json:"access_expires"`
}
