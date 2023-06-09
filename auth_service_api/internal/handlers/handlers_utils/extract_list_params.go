package handlers_utils

import (
	"auth_service_api/internal/exceptions"
	"auth_service_api/internal/schemas"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// extract params from request and create validated schemas.ListParams object
func ExtractListParams(r *http.Request) (*schemas.ListParams, error) {
	// Parse params
	params, err := parseQueryListParams(r)
	if err != nil {
		return nil, err
	}
	// Validate parsed params
	err = (*params).Validate()
	if err != nil {
		return nil, exceptions.MakeInvalidEntityError(fmt.Sprintf("could not prepare validate params: %v", err))
	}

	return params, nil
}

// extract params from request and parse them to schemas.ListParams
func parseQueryListParams(r *http.Request) (*schemas.ListParams, error) {
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")
	sortStr := r.URL.Query().Get("sort")

	// Create ListParams
	var params schemas.ListParams
	// Get and set the limit if it exists
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			return nil, exceptions.MakeInvalidEntityError(fmt.Sprintf("could not validate limit: %v", err))
		}
		params.Limit = &limit
	}

	// Get and set the offset if it exists
	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			return nil, exceptions.MakeInvalidEntityError(fmt.Sprintf("could not validate offset: %v", err))
		}
		params.Offset = &offset
	}

	// Get and set the sort if it exists
	if sortStr != "" {
		// Use the ParseSorting function to extract the sorting field and direction
		fieldName, ordering, err := parseSorting(&sortStr)
		if err != nil {
			return nil, err
		}
		params.SortingField = fieldName
		params.SortingDirection = ordering
	}
	return &params, nil
}

// parseSorting parses specific syntax sort (*string) param
// returns:
//
//	fieldname (*string), ordering (*string), error
func parseSorting(sorting *string) (*string, *string, error) {
	if sorting == nil {
		return nil, nil, nil
	}

	parts := strings.SplitN(*sorting, "[", 2)
	if len(parts) != 2 {
		return nil, nil, nil
	}

	fieldName := parts[0]
	ordering := strings.TrimRight(parts[1], "]")
	errorMsg := ""
	if fieldName == "" || ordering == "" {
		errorMsg = "could not parse sorting: fieldname or ordering can not be empty."
	}
	if ordering != "ASC" && ordering != "DESC" {
		errorMsg = "could not parse sorting: ordering must be 'ASC' or 'DESC'."
	}
	if errorMsg != "" {
		return nil, nil, exceptions.MakeInvalidEntityError(errorMsg)
	}

	return &fieldName, &ordering, nil
}
