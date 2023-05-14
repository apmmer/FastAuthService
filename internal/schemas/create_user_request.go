package schemas

// CreateUserRequest godoc
// CreateUserRequest represents user data for creating a new user
// @Schema
// @ID CreateUserRequest
// @Required username email password
// @Property username string "Username" minLength(1)
// @Property email string "Email" format(email)
// @Property password string "Password" minLength(1)
type CreateUserRequest struct {
	Username string `json:"username"`

	Email string `json:"email"`

	Password string `json:"password,omitempty"`
}
