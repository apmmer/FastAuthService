package base_repo

import (
	"AuthService/database"
	"context"
	"fmt"
	"strings"
)

// CreateOne inserts a new record into the specified table with the provided data
func CreateOne(tableName string, fields string, values string) (int, error) {
	sql := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s) RETURNING (id);",
		tableName,
		fields,
		values,
	)

	var id int
	err := database.Pool.QueryRow(context.Background(), sql).Scan(&id)
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
		return nil, fmt.Errorf("could not parse filters: %v", err)
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
		sql += fmt.Sprintf(" OFFSET %d", len(args))
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
