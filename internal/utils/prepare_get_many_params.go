package utils

import (
	"AuthService/internal/schemas"
	"fmt"
	"strconv"
)

func GetValidatedListRequestParams(limitStr string, offsetStr string, sortStr string) (*schemas.ListParams, error) {
	// Create ListParams
	var params schemas.ListParams
	// Get and set the limit if it exists
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			return nil, fmt.Errorf("could not prepare limit: %v", err)
		}
		params.Limit = &limit
	}

	// Get and set the offset if it exists
	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			return nil, fmt.Errorf("could not prepare offset: %v", err)
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
		return nil, fmt.Errorf("could not prepare validate params: %v", err)
	}

	return &params, nil
}
