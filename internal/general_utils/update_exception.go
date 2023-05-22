package general_utils

import (
	"AuthService/internal/exceptions"
	"fmt"
	"log"
)

// Updates exception message with provided prefix and returns the same type of exception.
func UpdateExceptionMsg(msg string, err error) error {
	switch err.(type) {
	case *exceptions.ErrNotFound:
		log.Println("Update ErrNotFound...")
		return &exceptions.ErrNotFound{Message: fmt.Sprintf("%s: %v", msg, err)}
	case *exceptions.ErrMultipleEntries:
		log.Println("Update ErrMultipleEntries...")
		return &exceptions.ErrMultipleEntries{Message: fmt.Sprintf("%s: %v", msg, err)}
	case *exceptions.ErrInvalidEntity:
		log.Println("Update ErrInvalidEntity...")
		return &exceptions.ErrInvalidEntity{Message: fmt.Sprintf("%s: %v", msg, err)}
	case *exceptions.ErrDbConflict:
		log.Println("Update ErrDbConflict...")
		return &exceptions.ErrDbConflict{Message: fmt.Sprintf("%s: %v", msg, err)}
	case *exceptions.ErrNoAuthData:
		log.Println("Update ErrNoAuthData...")
		return &exceptions.ErrNoAuthData{Message: fmt.Sprintf("%s: %v", msg, err)}
	case *exceptions.ErrUnauthorized:
		log.Println("Update ErrUnauthorized...")
		return &exceptions.ErrUnauthorized{Message: fmt.Sprintf("%s: %v", msg, err)}

	default:
		log.Println("Update default Exception...")
		return fmt.Errorf("%s: %v", msg, err)
	}
}
