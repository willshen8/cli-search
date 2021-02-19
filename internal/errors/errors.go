package errors

import (
	"errors"
	"fmt"
	"log"
)

var logFatal = log.Fatal

// HandleError takes an error and log it to output gracefully
func HandleError(err error) {
	if err != nil {
		logFatal(err)
	}
}

type Error struct {
	Err     error
	Context string
}

// Error methods implements the error interface method by printing out the actual error
func (e *Error) Error() string {
	return fmt.Sprintf("%s: %v", e.Context, e.Err)
}

// NewError creates a new error
func NewError(err error, info string, args ...interface{}) *Error {
	return &Error{
		err,
		fmt.Sprintf(info, args...),
	}
}

var (
	ErrInvalidTable       = errors.New("You have specified an incorrect table: ")
	ErrInvalidSearchField = errors.New("Invalid search field: ")
)
