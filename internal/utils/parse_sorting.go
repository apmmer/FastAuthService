package utils

import (
	"strings"
)

func ParseSorting(sorting *string) (*string, *string) {
	if sorting == nil {
		return nil, nil
	}

	parts := strings.SplitN(*sorting, "[", 2)
	if len(parts) != 2 {
		return nil, nil
	}

	fieldName := parts[0]
	ordering := strings.TrimRight(parts[1], "]")

	return &fieldName, &ordering
}
