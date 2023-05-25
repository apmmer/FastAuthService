package base_repo_utils

import (
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4"
)

func ParseSQLFilters(filters *map[string]interface{}, queryArgs *[]interface{}) (string, []interface{}) {
	filterStr := ""
	args := *queryArgs

	if filters != nil && len(*filters) > 0 {
		// Avoid SQL injection by using placeholders and passing values separately
		for field, value := range *filters {
			// If we filter by value=nil, it means we want to filter by field=NULL.
			if value == nil {
				filterStr += fmt.Sprintf(" %s IS NULL AND", field)
			} else {
				args = append(args, value)
				filterStr += fmt.Sprintf(" %s = $%d AND", field, len(args))
			}
		}
		filterStr = strings.TrimSuffix(filterStr, " AND") // Remove the trailing ' AND'
	}

	return filterStr, args
}

func ParseSQLResults(rowsPointer *pgx.Rows) (*[]map[string]interface{}, error) {
	rows := *rowsPointer
	var results []map[string]interface{}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, fmt.Errorf("could not get row values: %v", err)
		}

		item := make(map[string]interface{})
		for i, fd := range rows.FieldDescriptions() {
			item[string(fd.Name)] = values[i]
		}
		results = append(results, item)
	}
	return &results, nil
}
