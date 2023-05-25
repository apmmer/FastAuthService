package repositories_utils

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
	modelFields := GetModelFields(model)
	for key := range *filters {
		err := FieldInModelFields(key, modelFields)
		if err != nil {
			return err
		}
	}
	return nil
}

func FieldInModelFields(attribute string, modelFields []string) error {
	fieldFound := false
	for _, modelField := range modelFields {
		if attribute == modelField {
			fieldFound = true
			break
		}
	}
	if !fieldFound {
		return exceptions.MakeInvalidEntityError(
			fmt.Sprintf("attribute %s was not found in model.", attribute),
		)
	}
	return nil
}
