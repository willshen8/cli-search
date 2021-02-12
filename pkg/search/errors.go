package search

import "log"

// HandleError takes an error and log it to output gracefully
func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
