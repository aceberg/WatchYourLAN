package check

import (
	"fmt"
	"log/slog"
)

// IfError prints error, if it is not nil
func IfError(err error) bool {
	if err == nil {
		return false
	}

	slog.Error(fmt.Sprintf("%v", err))
	return true
}
