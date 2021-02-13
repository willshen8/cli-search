package search

import (
	"errors"
	"log"
)

// HandleError takes an error and log it to output gracefully
func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var (
	ErrInvalidSearchField = errors.New("Invalid search field")
)
