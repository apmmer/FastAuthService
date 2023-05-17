package utils

import "reflect"

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
