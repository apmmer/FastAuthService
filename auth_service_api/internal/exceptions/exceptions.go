// Exceptions package is based on patterns 'fabric' and 'strategy'
// Each exception here will contain a custom message (as well as default message)
// and associated status code for server response.
package exceptions

import "net/http"

// DefaultError is an error with default message and associated status code for server error response
// All errors constructors will use this class to return.
type DefaultError struct {
	Message    string
	StatusCode int
}

func (e *DefaultError) Error() string {
	return e.Message
}

func (e *DefaultError) GetStatusCode() int {
	return e.StatusCode
}

// Now declaring functions-fabrics (pattern)

// MakeNotFoundError is used to create an error, when no entries are found according to filters.
func MakeNotFoundError(message string) *DefaultError {
	if message == "" {
		message = "Object not found."
	}
	return &DefaultError{
		Message:    message,
		StatusCode: http.StatusNotFound,
	}
}

// MakeNotFoundError is used to create an error, when no entries are found according to filters.
func MakeMultipleEntriesError(message string) *DefaultError {
	if message == "" {
		message = "Found multiple records, which is unexpected."
	}
	return &DefaultError{
		Message:    message,
		StatusCode: http.StatusNotAcceptable,
	}
}

// MakeInvalidEntityError is used to create an error, when there is a problem with input data (in endpoint).
func MakeInvalidEntityError(message string) *DefaultError {
	if message == "" {
		message = "Input data is invalid."
	}
	return &DefaultError{
		Message:    message,
		StatusCode: http.StatusUnprocessableEntity,
	}
}

// MakeDbConflictError is used to create an error, when there is a conflict in db, like FKey exception etc.
func MakeDbConflictError(message string) *DefaultError {
	if message == "" {
		message = "Conflict error: can not create/modify db record."
	}
	return &DefaultError{
		Message:    message,
		StatusCode: http.StatusConflict,
	}
}

// MakeNoAuthDataError is used to create an error, when there is an auth required and no auth data provided.
func MakeNoAuthDataError(message string) *DefaultError {
	if message == "" {
		message = "Authentication data was not provided or can not be parsed."
	}
	return &DefaultError{
		Message:    message,
		StatusCode: http.StatusForbidden,
	}
}

// MakeUnauthorizedError is used to create an error, when there is an auth required,
// auth data was provided, but is invalid.
func MakeUnauthorizedError(message string) *DefaultError {
	if message == "" {
		message = "Authentication data is invalid."
	}
	return &DefaultError{
		Message:    message,
		StatusCode: http.StatusUnauthorized,
	}
}

// MakeValidationError is used to create an error, when data can not be validated somewhere in the internal code.
// for example, when the data successfully passed the handler, getting into the repository,
// where it turned out that it was impossible to create an instance of the model or something similar.
func MakeValidationError(message string) *DefaultError {
	if message == "" {
		message = "Validation for data was failed."
	}
	return &DefaultError{
		Message:    message,
		StatusCode: http.StatusUnprocessableEntity,
	}
}

// MakeInternalError is used to create an error, when reason can not be explained for client.
func MakeInternalError(message string) *DefaultError {
	if message == "" {
		message = "Server can not process your request."
	}
	return &DefaultError{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
	}
}
