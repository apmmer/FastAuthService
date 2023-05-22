package base_repo

import (
	"AuthService/database"
	"AuthService/internal/general_utils"
	"AuthService/internal/repositories/base_repo/base_repo_utils"
	"context"
	"fmt"
)

func ExecuteRowParseList(sql string, args []interface{}) (*[]map[string]interface{}, error) {
	// execute SQL query (getting pgx.Rows)
	rows, err := database.Pool.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, fmt.Errorf("could not get records from db: %v", err)
	}
	defer rows.Close()
	// parse results
	results, err := base_repo_utils.ParseSQLResults(&rows)
	if err != nil {
		errMsg := fmt.Sprintf("could not parse SQL results from db")
		return nil, general_utils.UpdateExceptionMsg(errMsg, err)
	}
	return results, nil
}
