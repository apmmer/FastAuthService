package schemas

// CreateUserRequest godoc
// CreateUserRequest represents the request to create a new user.
// @Schema(name="CreateUserRequest")
// @Required screen_name email password
// @Description This request includes the necessary details to create a new user.
// @Example {"screen_name": "jonhDoe123", "email": "user@example.com", "password": "1234567", "company_id": 1, "rank": 1}
type CreateUserRequest struct {
	// The desired username for the new user. This field is required and must contain at least 4 characters.
	// @MinLength 4
	// @Example jonhDoe123
	ScreenName string `json:"screen_name,omitempty" validate:"required,min=4"`

	// The email address of the new user. This field is required and must be a valid email address.
	// @Format email
	// @Example john@example.com
	Email string `json:"email,omitempty" validate:"required,email"`

	// The password for the new user. This field is required and must contain at least 7 characters.
	// @MinLength 7
	// @Example 1234567
	Password string `json:"password,omitempty" validate:"required,min=7"`

	// The ID of the company to associate with the new user. This field is optional and must be greater than 0 if provided.
	// @Min 0
	// @Example 1
	CompanyId *int `json:"company_id" validate:"omitempty,gt=0"`

	// The rank of the new user. This field is optional and must be greater than 0 if provided.
	// @Min 0
	// @Example 1
	Rank *int `json:"rank" validate:"omitempty,gt=0"`
}

func (c CreateUserRequest) Validate() error {
	return Validate(c)
}
