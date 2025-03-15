package check

import (
	"os"
	"path/filepath"
)

// Path - create path if not exists
func Path(path string) bool {

	_, err := os.Stat(path)

	if path != "" && err != nil {

		dir := filepath.Dir(path)

		err = os.MkdirAll(dir, os.ModePerm)
		IfError(err)

		_, err = os.Create(path)
		IfError(err)

		return false
	}

	return true
}

// Exists - check is file exists
func Exists(path string) bool {

	_, err := os.Stat(path)

	if path != "" && err != nil {

		return false
	}

	return true
}

// IsYaml - check if file got .yaml or .yml extension
func IsYaml(path string) bool {

	if Exists(path) {
		ext := filepath.Ext(path)
		if ext == ".yaml" || ext == ".yml" {
			return true
		}
	}

	return false
}

// IsEmpty - check if file is empty
func IsEmpty(path string) bool {

	if Exists(path) {
		stat, _ := os.Stat(path)
		size := stat.Size()
		if size > 0 {
			return false
		}
	}

	return true
}
