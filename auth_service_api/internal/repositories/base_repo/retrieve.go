package base_repo

import (
	"auth_service_api/internal/exceptions"
	"auth_service_api/internal/general_utils"
	"auth_service_api/internal/repositories/base_repo/base_repo_utils"
	"fmt"
	"log"
	"strings"
)

// Performs SELECT operation to DB according provided params.
// Returns:
//
//	results ([]map[string]interface{}) - parsed list results.
func GetMany(tableName string, limit *int, offset *int, orderBy *string, orderingDirection *string, filters *map[string]interface{}) ([]map[string]interface{}, error) {
	log.Printf("Called base_repo.GetMany filters=%v", filters)
	sql := fmt.Sprintf("SELECT * FROM %s", tableName)
	var args []interface{}

	filterStr, args := base_repo_utils.ParseSQLFilters(filters, &args)
	if filterStr != "" {
		sql += " WHERE" + filterStr
	}

	sql, args, err := processListParams(sql, args, orderBy, orderingDirection, limit, offset)
	if err != nil {
		return nil, err
	}

	log.Printf("base_repo.GetMany: \n%s \nargs: %v", sql, args)
	results, err := ExecuteRowParseList(sql, args)
	if err != nil {
		return nil, err
	}
	return *results, nil
}

// updates SQL query with new params if they were provided (not nil)
// returns SQL query string and updated args
func processListParams(sql string, args []interface{}, orderBy *string, orderingDirection *string, limit *int, offset *int) (string, []interface{}, error) {
	if orderBy != nil {
		sql += fmt.Sprintf(" ORDER BY %s", *orderBy)
		if orderingDirection != nil {
			direction := strings.ToUpper(*orderingDirection)
			if direction != "ASC" && direction != "DESC" {
				return "", nil, exceptions.MakeInvalidEntityError(
					fmt.Sprintf("Invalid order direction: %s", direction),
				)
			}
			sql += fmt.Sprintf(" %s", direction)
		}
	}

	if limit != nil {
		args = append(args, *limit)
		sql += fmt.Sprintf(" LIMIT $%d", len(args))
	}

	if offset != nil {
		args = append(args, *offset)
		sql += fmt.Sprintf(" OFFSET $%d", len(args))
	}
	return sql, args, nil
}

// Retrieves 1 record from the table according filters, else returns an error.
func GetOne(tableName string, filters *map[string]interface{}) (map[string]interface{}, error) {
	log.Println("Called base_repo.GetOne")
	// retirving records using GetMany method
	records, err := GetMany(tableName, nil, nil, nil, nil, filters)
	log.Printf("base_repo.GetOne: Got %d records using GetMany", len(records))
	if err != nil {
		return nil, general_utils.UpdateException("could not perform GetMany method", err)
	}
	// we expect that we have only 1 record, so validate:
	if records == nil || len(records) == 0 {
		return nil, exceptions.MakeNotFoundError("got no records according filters.")
	}
	if len(records) > 1 {
		return nil, exceptions.MakeMultipleEntriesError("got multiple records according filters, but expected 1.")
	}
	return records[0], nil
}
