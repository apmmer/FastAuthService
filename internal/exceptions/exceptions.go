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
type ErrUnprocessableEntity struct {
	Message string
}

func (e *ErrUnprocessableEntity) Error() string {
	return e.Message
}
