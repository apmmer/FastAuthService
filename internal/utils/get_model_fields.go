package utils

import (
	"AuthService/internal/exceptions"
	"fmt"
	"reflect"
)

// Returns a list of fields of provided model
func GetModelFields(model interface{}) []string {
	fields := []string{}
	reflectValue := reflect.ValueOf(model)
	reflectType := reflectValue.Type()
	for i := 0; i < reflectType.NumField(); i++ {
		field := reflectType.Field(i)
		dbTag := field.Tag.Get("db")
		if dbTag != "" {
			fields = append(fields, dbTag)
		}
	}
	return fields
}

// Validate if filters fieldnames are correct for provided model
func ValidateMapFields(filters *map[string]interface{}, model interface{}) error {
	model_fields := GetModelFields(model)
	for key := range *filters {
		field_found := false
		for _, fieldname := range model_fields {
			if key == fieldname {
				field_found = true
				break
			}
		}
		if !field_found {
			return &exceptions.ErrInvalidEntity{
				Message: fmt.Sprintf("field %s was not found in model", key),
			}
		}
	}
	return nil
}
