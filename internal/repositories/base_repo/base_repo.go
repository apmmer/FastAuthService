package base_repo

import (
	"AuthService/database"
	"context"
	"fmt"
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

func GetMany(tableName string, limit *int, offset *int, orderBy *string, orderingDirection *string) ([]map[string]interface{}, error) {
	sql := fmt.Sprintf("SELECT * FROM %s", tableName)

	if orderBy != nil {
		sql += fmt.Sprintf(" ORDER BY %s", *orderBy)
		if orderingDirection != nil {
			sql += fmt.Sprintf(" %s", *orderingDirection)
		}
	}

	if limit != nil {
		sql += fmt.Sprintf(" LIMIT %d", *limit)
	}

	if offset != nil {
		sql += fmt.Sprintf(" OFFSET %d", *offset)
	}

	rows, err := database.Pool.Query(context.Background(), sql)
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
