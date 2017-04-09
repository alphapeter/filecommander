package commands

import "strings"

func ValidatePath(path string) error {
	if strings.Contains(path, "..") {
		return validationError{"using '..' in path to change directory is not allowed"}
	}
	if strings.Contains(path, "~") {
		return validationError{"using '~' in path to change directory is not allowed"}
	}
	return nil
}

type validationError struct {
	error string
}

func (e validationError) Error() string {
	return e.error
}
