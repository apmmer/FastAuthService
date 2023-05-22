package schemas

import (
	"AuthService/internal/exceptions"
	"fmt"
	"strconv"
)

// ListParams godoc
// ListParams represents the parameters for getting many objects.
// @Schema(name="ListParams")
// @Required
// @Description These parameters control the pagination and sorting of the results.
// @Example {"limit": 10, "offset": 0, "sorting": "email[asc]"}
type ListParams struct {
	Limit   *int    `json:"limit" validate:"omitempty,min=0"`
	Offset  *int    `json:"offset" validate:"omitempty,min=0"`
	Sorting *string `json:"sort"`
}

func (p ListParams) Validate() error {
	return Validate(p)
}

func GetValidatedListParams(limitStr string, offsetStr string, sortStr string) (*ListParams, error) {
	// Create ListParams
	var params ListParams
	// Get and set the limit if it exists
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			return nil, &exceptions.ErrInvalidEntity{
				Message: fmt.Sprintf("could not prepare limit: %v", err),
			}
		}
		params.Limit = &limit
	}

	// Get and set the offset if it exists
	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			return nil, &exceptions.ErrInvalidEntity{
				Message: fmt.Sprintf("could not prepare offset: %v", err),
			}
		}
		params.Offset = &offset
	}

	// Get and set the sort if it exists
	if sortStr != "" {
		params.Sorting = &sortStr
	}

	// Validate params
	err := params.Validate()
	if err != nil {
		return nil, &exceptions.ErrInvalidEntity{
			Message: fmt.Sprintf("could not prepare validate params: %v", err),
		}
	}

	return &params, nil
}
