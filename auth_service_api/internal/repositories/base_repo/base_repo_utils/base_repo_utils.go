package base_repo_utils

import (
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4"
)

// ParseSQLFilters takes a pointer to a map[string]interface{} and transforms it into a SQL query string and a list of arguments.
// This function helps in creating SQL query string and its arguments for data filtering based on the given map object.
// SQL injection is prevented by using placeholders and passing values separately.
// If a map value is nil, it generates a SQL condition for the field to be NULL.
//
// Args:
//
// filters: Pointer to a map where each key-value pair represents a field and its value to be filtered.
// queryArgs: Pointer to an existing list of arguments, which may be used in the rest of the SQL query.
//
// Returns:
//
// filterStr: A string containing SQL conditions for each field-value pair in the filters map,
//
//	ready to be inserted in a SQL 'WHERE' clause. Each condition is separated by 'AND'.
//
// args:      A list of arguments corresponding to the values in the filters map. This list includes
//
//	the initial elements pointed to by queryArgs, followed by the values in the filters map.
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

// ParseSQLResults converts a pgx.Rows pointer into a slice of maps. Each map represents a row from the query result,
// with each key-value pair in the map representing a field and its value in the row.
//
// Args:
// rowsPointer: Pointer to pgx.Rows obtained from a SQL query.
//
// Returns:
// Pointer to a slice of maps, where each map represents a row from the SQL query result.
// Error, if any occurred during the operation. Otherwise, nil.
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
