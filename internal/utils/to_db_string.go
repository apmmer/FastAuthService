package utils

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// MarshalToDBString converts an object into two strings that can be used in an SQL query.
// The first string contains the field names separated by commas.
// The second string contains the field values separated by commas.
func MarshalToDBString(data interface{}) (string, string, error) {
	v := reflect.ValueOf(data)
	t := v.Type()

	var fields []string
	var values []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// Если поле не экспортируемое, пропустить его
		if !value.CanInterface() {
			continue
		}

		// Get the field's value as an interface and type-assert.
		var val interface{}
		if value.Kind() == reflect.Ptr {
			if value.IsNil() {
				values = append(values, "NULL")
				fields = append(fields, field.Tag.Get("db"))
				continue
			} else {
				val = value.Elem().Interface()
			}
		} else {
			if value.IsZero() {
				continue
			}
			val = value.Interface()
		}

		fields = append(fields, field.Tag.Get("db"))

		switch v := val.(type) {
		case string:
			values = append(values, fmt.Sprintf("'%s'", v))
		case int, int32, int64, float32, float64, uint:
			values = append(values, fmt.Sprintf("%v", v))
		case bool:
			values = append(values, fmt.Sprintf("%t", v))
		case time.Time:
			values = append(values, fmt.Sprintf("'%s'", v.Format(time.RFC3339)))
		default:
			// Если тип не поддерживается, вернуть ошибку
			return "", "", fmt.Errorf("unsupported data type: %v", field.Type)
		}
	}

	return strings.Join(fields, ", "), strings.Join(values, ", "), nil
}
