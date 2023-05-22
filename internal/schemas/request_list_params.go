package schemas

// ListParams godoc
// ListParams represents the parameters for getting many objects.
// @Schema(name="ListParams")
// @Required
// @Description These parameters control the pagination and sorting of the results.
// @Example {"limit": 10, "offset": 0, "sorting": "email[asc]"}
type ListParams struct {
	Limit            *int    `json:"limit" validate:"omitempty,min=0"`
	Offset           *int    `json:"offset" validate:"omitempty,min=0"`
	SortingField     *string `json:"sorting_field"`
	SortingDirection *string `json:"sorting_direction"`
}

func (p ListParams) Validate() error {
	return Validate(p)
}
