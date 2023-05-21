package base_repo

import (
	"AuthService/database"
	"AuthService/internal/exceptions"
	"AuthService/internal/repositories/base_repo/base_repo_utils"
	"AuthService/internal/utils"
	"context"
	"fmt"
	"log"
	"strings"
)

// placeholders generates a string of PostgreSQL placeholder syntax for SQL queries
func placeholders(n int) string {
	if n < 1 {
		return ""
	}

	// Start building the string with the first placeholder
	var buf strings.Builder
	buf.WriteString("$1")

	// Add the rest of the placeholders, each prefixed with a comma
	for i := 1; i < n; i++ {
		buf.WriteString(fmt.Sprintf(", $%d", i+1))
	}

	return buf.String()
}

// CreateOne inserts a new record into the specified table with the provided data
func CreateOne(tableName string, fields []string, values []interface{}) (int, error) {
	// Build the SQL query string
	sql := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s) RETURNING id;",
		tableName,
		strings.Join(fields, ", "),
		placeholders(len(values)),
	)

	// Execute the SQL query
	var id int
	err := database.Pool.QueryRow(context.Background(), sql, values...).Scan(&id)
	if err != nil {
		return 0, &exceptions.ErrDbConflict{
			Message: fmt.Sprintf("could not insert into %s table: %v", tableName, err),
		}
	}

	return id, nil
}

func GetMany(tableName string, limit *int, offset *int, orderBy *string, orderingDirection *string, filters *map[string]interface{}) ([]map[string]interface{}, error) {
	fmt.Printf("Called base_repo.GetMany filters=%v", filters)
	sql := fmt.Sprintf("SELECT * FROM %s", tableName)
	var args []interface{}

	filterStr, args, err := base_repo_utils.ParseSQLFilters(filters, &args)
	if err != nil {
		return nil, &exceptions.ErrInvalidEntity{
			Message: fmt.Sprintf("failed to validate filters: %v", err),
		}
	}
	if filterStr != "" {
		sql += " WHERE" + filterStr
	}

	if orderBy != nil {
		sql += fmt.Sprintf(" ORDER BY %s", *orderBy)
		if orderingDirection != nil {
			direction := strings.ToUpper(*orderingDirection)
			if direction != "ASC" && direction != "DESC" {
				return nil, &exceptions.ErrInvalidEntity{
					Message: fmt.Sprintf("Invalid order direction: %s", direction),
				}
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

	log.Printf("base_repo.GetMany: \n%s \nargs: %v", sql, args)
	rows, err := database.Pool.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, fmt.Errorf("could not get records from %s table: %v", tableName, err)
	}
	defer rows.Close()

	results, err := base_repo_utils.ParseSQLResults(&rows)
	if err != nil {
		errMsg := fmt.Sprintf("could not parse SQL results from %s table", tableName)
		return nil, utils.UpdateExceptionMsg(errMsg, err)
	}

	return *results, nil
}

// Retrieves 1 record from the table according filters, else returns an error.
func GetOne(tableName string, filters *map[string]interface{}) (map[string]interface{}, error) {
	fmt.Println("Called base_repo.GetOne")
	// retirving records using GetMany method
	records, err := GetMany(tableName, nil, nil, nil, nil, filters)
	log.Printf("base_repo.GetOne: Got %d records using GetMany", len(records))
	if err != nil {
		return nil, utils.UpdateExceptionMsg("could not perform GetMany method", err)
	}
	// we expect that we have only 1 record, so validate:
	if records == nil || len(records) == 0 {
		return nil, &exceptions.ErrNotFound{Message: "got no records according filters."}
	}
	if len(records) > 1 {
		return nil, &exceptions.ErrMultipleEntries{Message: "got multiple records according filters, but expected 1."}
	}
	return records[0], nil
}
