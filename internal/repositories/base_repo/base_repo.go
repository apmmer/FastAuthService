package base_repo

import (
	"AuthService/database"
	"AuthService/internal/exceptions"
	"context"
	"fmt"
	"log"
	"strings"
)

// CreateOne inserts a new record into the specified table with the provided data
func CreateOne(tableName string, fields string, values string) (int, error) {
	sql := fmt.Sprintf(
		"INSERT INTO %s ($1) VALUES ($2) RETURNING (id);",
		tableName,
	)

	var id int
	err := database.Pool.QueryRow(context.Background(), sql, fields, values).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("could not insert into %s table: %v", tableName, err)
	}

	return id, nil
}

func ParseSQLFilters(filters *map[string]interface{}) (string, []interface{}, error) {
	filterStr := ""
	var args []interface{}

	if filters != nil && len(*filters) > 0 {
		for field, value := range *filters {
			// Avoid SQL injection by using placeholders and passing values separately
			args = append(args, value)
			filterStr += fmt.Sprintf(" %s = $%d AND", field, len(args))
		}
		filterStr = strings.TrimSuffix(filterStr, " AND") // Remove the trailing ' AND'
	}

	return filterStr, args, nil
}

func GetMany(tableName string, limit *int, offset *int, orderBy *string, orderingDirection *string, filters *map[string]interface{}) ([]map[string]interface{}, error) {
	sql := fmt.Sprintf("SELECT * FROM %s", tableName)

	filterStr, args, err := ParseSQLFilters(filters)
	if err != nil {
		return nil, &exceptions.ErrInvalidEntity{
			Message: fmt.Sprintf("failed to validate filters: %v", err),
		}
	}

	if filterStr != "" {
		sql += " WHERE" + filterStr
	}

	if orderBy != nil {
		args = append(args, *orderBy)
		sql += fmt.Sprintf(" ORDER BY %s", *orderBy)

		if orderingDirection != nil {
			args = append(args, *orderingDirection)
			sql += fmt.Sprintf(" $%d", len(args))
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

	rows, err := database.Pool.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, fmt.Errorf("could not get records from %s table: %v", tableName, err)
	}
	defer rows.Close()

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

	return results, nil
}

// Retrieves 1 record from the table according filters, else returns an error.
func GetOne(tableName string, filters *map[string]interface{}) (map[string]interface{}, error) {
	// retirving records using GetMany method
	records, err := GetMany(tableName, nil, nil, nil, nil, filters)
	log.Println("Got records using GetMany")

	if err != nil {
		return nil, fmt.Errorf("could not perform GetMany method: %v", err)
	}
	// we expect that we have only 1 record, so validate:
	if records != nil && len(records) == 0 || len(records) == 0 {
		return nil, &exceptions.ErrNotFound{Message: "got no records according filters."}
	}
	if len(records) > 1 {
		return nil, &exceptions.ErrMultipleEntries{Message: "got multiple records according filters, but expected 1."}
	}
	return records[0], nil
}
