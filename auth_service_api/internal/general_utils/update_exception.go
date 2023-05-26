package general_utils

import (
	"auth_service_api/internal/exceptions"
	"fmt"
)

// Updates exception message with provided prefix and returns the exceptions.DefaultError.
func UpdateException(msg string, err error) error {
	var newMsg string
	if msg != "" {
		newMsg = fmt.Sprintf("%s -> %v", msg, err)
	} else {
		newMsg = err.Error()
	}

	if customErr, ok := err.(*exceptions.DefaultError); ok {
		customErr.Message = newMsg
		return customErr
	} else {
		// unhandled exception, so set default status code 500
		return exceptions.MakeInternalError(newMsg)
	}
}
