package schemas

import "github.com/go-playground/validator/v10"

// CreateUserRequest godoc
// CreateUserRequest represents user data for creating a new user
// @Schema
// @ID CreateUserRequest
// @Required screen_name email password
// @Property screen_name string "ScreenName" minLength(4)
// @Property email string "Email" format(email)
// @Property password string "Password" minLength(7)
// @Property company_id int "CompanyId"
// @Property rank int "Rank"
type CreateUserRequest struct {
	ScreenName string `json:"screen_name,omitempty" validate:"required,min=4"`
	Email      string `json:"email,omitempty" validate:"required,email"`
	Password   string `json:"password,omitempty" validate:"required,min=7"`
	CompanyId  *int   `json:"company_id" validate:"omitempty,gt=0"`
	Rank       *int   `json:"rank" validate:"omitempty,gt=0"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (c CreateUserRequest) Validate() error {
	return validate.Struct(c)
}
