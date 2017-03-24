package commands

import (
	"os"
)

func Delete(file string) error {
	path, err := getPath(file)

	if err != nil {
		return err
	}

	if err := os.Remove(path); err != nil {
		return err
	}
	return nil
}
