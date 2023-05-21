package base_repo

import (
	"AuthService/database"
	"AuthService/internal/repositories/base_repo/base_repo_utils"
	"AuthService/internal/utils"
	"context"
	"fmt"
	"strings"
)

// Updates records in given table
func Update(tableName string, filters *map[string]interface{}, updateData *map[string]interface{}) (*[]map[string]interface{}, error) {
	sql := fmt.Sprintf("UPDATE %s", tableName)
	setFieldsStr, args, err := parseSQLUpdateData(updateData)
	if err != nil {
		return nil, utils.UpdateExceptionMsg("could not parse sql updateData", err)
	}
	filterStr, args, err := base_repo_utils.ParseSQLFilters(filters, &args)
	if err != nil {
		return nil, utils.UpdateExceptionMsg("could not parse sql filters", err)
	}
	// returnFieldsStr, args, err := base_repo_utils.ParseSQLReturningFields(returningFields)
	if err != nil {
		return nil, utils.UpdateExceptionMsg("could not parse sql returningFields", err)
	}
	if setFieldsStr != "" {
		sql += " SET" + setFieldsStr
	}
	if filterStr != "" {
		sql += " WHERE" + filterStr
	}

	sql += " RETURNING *"
	if err != nil {
		return nil, utils.UpdateExceptionMsg("could not parse sql returningFields", err)
	}
	// getting pgx.Rows
	rows, err := database.Pool.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, fmt.Errorf("could not get records from %s table: %v", tableName, err)
	}
	// parse results to []map[string]interface{}
	results, err := base_repo_utils.ParseSQLResults(&rows)
	if err != nil {
		return nil, fmt.Errorf("could not parse SQL results from %s table: %v", tableName, err)
	}
	return results, nil
}

// parses provided data to SQL Query string with placeholders and args (separately)
func parseSQLUpdateData(updateData *map[string]interface{}) (string, []interface{}, error) {
	if updateData == nil || len(*updateData) == 0 {
		return "", nil, nil
	}

	placeholders := make([]string, 0)
	args := make([]interface{}, 0)

	// Iterate over map to construct placeholders and args
	for field, value := range *updateData {
		if value == nil {
			placeholders = append(placeholders, fmt.Sprintf(" %s=NULL", field))
		} else {
			args = append(args, value) // the value is an argument only when it's not nil
			placeholders = append(placeholders, fmt.Sprintf(" %s=$%d", field, len(args)))
		}
	}

	// Join all placeholders with a comma
	updateStr := strings.Join(placeholders, ",")

	return updateStr, args, nil
}
