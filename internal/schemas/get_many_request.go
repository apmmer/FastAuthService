package schemas

// GetManyRequestParams godoc
// GetManyRequestParams represents the parameters for getting many objects.
// @Schema(name="GetManyRequestParams")
// @Required
// @Description These parameters control the pagination and sorting of the results.
// @Example {"limit": 10, "offset": 0, "sorting": "email[asc]"}
type GetManyRequestParams struct {
	Limit   *int    `json:"limit" validate:"omitempty,min=0"`
	Offset  *int    `json:"offset" validate:"omitempty,min=0"`
	Sorting *string `json:"sort"`
}

func (p GetManyRequestParams) Validate() error {
	return Validate(p)
}
