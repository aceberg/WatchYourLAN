package check

import (
	"os"
	"path/filepath"
)

// Path - create path if not exist
func Path(path string) {

	_, err := os.Stat(path)

	if path != "" && err != nil {

		dir := filepath.Dir(path)

		err = os.MkdirAll(dir, os.ModePerm)
		IfError(err)

		_, err = os.Create(path)
		IfError(err)
	}
}
