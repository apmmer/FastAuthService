package exceptions

// When got no entries according filters
type ErrNotFound struct {
	Message string
}

func (e *ErrNotFound) Error() string {
	return e.Message
}

// When got too many entries and expected less
type ErrMultipleEntries struct {
	Message string
}

func (e *ErrMultipleEntries) Error() string {
	return e.Message
}

// input data validation error
type ErrInvalidEntity struct {
	Message string
}

func (e *ErrInvalidEntity) Error() string {
	return e.Message
}

// Db entries conflict error
type ErrDbConflict struct {
	Message string
}

func (e *ErrDbConflict) Error() string {
	return e.Message
}

// Auth data was not provided or can not be extracted
type ErrNoAuthData struct {
	Message string
}

func (e *ErrNoAuthData) Error() string {
	return e.Message
}

// Auth data is invalid
type ErrUnauthorized struct {
	Message string
}

func (e *ErrUnauthorized) Error() string {
	return e.Message
}
