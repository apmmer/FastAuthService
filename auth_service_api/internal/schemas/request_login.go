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
