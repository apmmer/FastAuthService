// Exceptions package is based on patterns 'fabric' and 'strategy'
// Each exception here will contain a custom message (as well as default message)
// and associated status code for server response.
package exceptions

import "net/http"

// ErrNotFound is used when no entries are found according to filters.
type ErrNotFound struct {
	Message string
}

func (e *ErrNotFound) Error() string {
	return e.Message
}

// ErrMultipleEntries is used when too many entries are found and less were expected.
type ErrMultipleEntries struct {
	Message string
}

func (e *ErrMultipleEntries) Error() string {
	return e.Message
}

// ErrInvalidEntity is used when input data validation fails.
type ErrInvalidEntity struct {
	Message string
}

func (e *ErrInvalidEntity) Error() string {
	return e.Message
}

// ErrDbConflict is used when there is a conflict in database entries.
type ErrDbConflict struct {
	Message string
}

func (e *ErrDbConflict) Error() string {
	return e.Message
}

// ErrNoAuthData is used when authentication data was not provided.
type ErrNoAuthData struct {
	Message string
}

func (e *ErrNoAuthData) Error() string {
	return e.Message
}

// ErrUnauthorized is used when the authentication data is invalid.
type ErrUnauthorized struct {
	Message string
}

func (e *ErrUnauthorized) Error() string {
	return e.Message
}

// ErrInputError is used when there is a problem with input data (in endpoint).
type ErrInputError struct {
	Message string
}

func (e *ErrInputError) Error() string {
	return e.Message
}

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
func MakeNotFoundError(message string) error {
	if message == "" {
		message = "Object not found."
	}
	return &DefaultError{
		Message:    message,
		StatusCode: http.StatusNotFound,
	}
}

// MakeNotFoundError is used to create an error, when no entries are found according to filters.
func MakeMultipleEntriesError(message string) error {
	if message == "" {
		message = "Found multiple records, which is unexpected."
	}
	return &DefaultError{
		Message:    message,
		StatusCode: http.StatusNotAcceptable,
	}
}

// MakeInvalidEntityError is used to create an error, when there is a problem with input data (in endpoint).
func MakeInvalidEntityError(message string) error {
	if message == "" {
		message = "Input data is invalid."
	}
	return &DefaultError{
		Message:    message,
		StatusCode: http.StatusUnprocessableEntity,
	}
}

func MakeDbConflictError(message string) error {
	if message == "" {
		message = "Conflict error: can not create/modify db record."
	}
	return &DefaultError{
		Message:    message,
		StatusCode: http.StatusConflict,
	}
}

func MakeNoAuthDataError(message string) error {
	if message == "" {
		message = "Authentication data was not provided or can not be parsed."
	}
	return &DefaultError{
		Message:    message,
		StatusCode: http.StatusForbidden,
	}
}

func MakeUnauthorizedError(message string) error {
	if message == "" {
		message = "Authentication data is invalid."
	}
	return &DefaultError{
		Message:    message,
		StatusCode: http.StatusUnauthorized,
	}
}

func MakeValidationError(message string) error {
	if message == "" {
		message = "Validation for data was failed."
	}
	return &DefaultError{
		Message:    message,
		StatusCode: http.StatusUnprocessableEntity,
	}
}

// MakeInternalError is used to create an error, when reason can not be explained for client.
func MakeInternalError(message string) error {
	if message == "" {
		message = "Server can not process your request."
	}
	return &DefaultError{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
	}
}
