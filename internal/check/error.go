package check

import (
	"log"
)

// IfError prints error, if it is not nil
func IfError(err error) bool {
	if err == nil {
		return false
	}

	log.Println("ERROR:", err)
	return true
}
