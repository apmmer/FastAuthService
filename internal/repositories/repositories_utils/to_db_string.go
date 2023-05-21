package repositories_utils

import (
	"reflect"
)

func GetFieldsAndValues(model interface{}, ignore_field string) ([]string, []interface{}) {
	val := reflect.ValueOf(model)

	// if the model is a pointer, then you need to dereference it to get the values
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
