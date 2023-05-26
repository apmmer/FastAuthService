package repositories_utils

import (
	"reflect"
)

// Extracts fields and values from model (or schema) instance.
// Returns:
//
//	fields ([]string) - a list of model's fields one by one
//	values ([]interface{}) - a list of model fields's values, one by one, any type.
func GetFieldsAndValues(model interface{}, ignore_field string) ([]string, []interface{}) {
	val := reflect.ValueOf(model)

	// if the model is a pointer, then we need to dereference it to get the values
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	num := val.NumField()
	var fields []string
	var values []interface{}

	for i := 0; i < num; i++ {
		// getting name of each field
		fieldName := val.Type().Field(i).Tag.Get("db")
		if fieldName == ignore_field {
			continue
		}
		// getting a value of each field
		fieldValue := val.Field(i).Interface()
		fields = append(fields, fieldName)
		values = append(values, fieldValue)
	}

	return fields, values
}
