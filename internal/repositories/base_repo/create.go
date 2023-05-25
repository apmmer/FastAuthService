package base_repo

import (
	"AuthService/database"
	"AuthService/internal/exceptions"
	"context"
	"fmt"
	"strings"
)

// CreateOne inserts a new record into the specified table with the provided data
func CreateOne(tableName string, fields []string, values []interface{}) (int, error) {
	// Build the SQL query string
	sql := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s) RETURNING id;",
		tableName,
		strings.Join(fields, ", "),
		addPlaceholders(len(values)),
	)

	// Execute the SQL query
	var id int
	err := database.Pool.QueryRow(context.Background(), sql, values...).Scan(&id)
	if err != nil {
		return 0, exceptions.MakeDbConflictError(fmt.Sprintf("could not insert into %s table: %v", tableName, err))
	}

	return id, nil
}

// placeholders generates a string of PostgreSQL placeholder syntax for SQL queries
func addPlaceholders(n int) string {
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
