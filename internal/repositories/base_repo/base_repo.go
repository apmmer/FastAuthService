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
