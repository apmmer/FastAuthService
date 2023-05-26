package base_repo

import (
	"auth_service_api/database"
	"auth_service_api/internal/exceptions"
	"auth_service_api/internal/repositories/base_repo/base_repo_utils"
	"context"
	"fmt"
)

// Executes provided raw SQL query and parses sql result, expecting a list of 'rows'.
// Returns:
//
//	results (*[]map[string]interface{}) - already parsed sql pgx.rows to golang objects
func ExecuteRowParseList(sql string, args []interface{}) (*[]map[string]interface{}, error) {
	// execute SQL query (getting pgx.Rows)
	rows, err := database.Pool.Query(context.Background(), sql, args...)
	if err != nil {
		errMsg := fmt.Sprintf("could not get records from db: %v", err)
		return nil, exceptions.MakeInternalError(errMsg)
	}
	defer rows.Close()
	// parse results
	results, err := base_repo_utils.ParseSQLResults(&rows)
	if err != nil {
		errMsg := fmt.Sprintf("could not parse SQL results from db")
		return nil, exceptions.MakeInternalError(errMsg)
	}
	return results, nil
}
