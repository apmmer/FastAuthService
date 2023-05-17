package schemas

// ErrorResponse godoc
// A scheme for server errors responses
// @Schema
// @ID ErrorResponse
// @Property error integer "Error" example("Epic fail, because...")
type ErrorResponse struct {
	Error string `json:"error"`
}
