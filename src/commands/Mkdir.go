package commands

import (
	"os"
)

func Mkdir(directory, directoryName string) error {
	if err := ValidatePath(directoryName); err != nil {
		return err
	}
	path, err := getPath(directory)

	if err != nil {
		return err
	}

	err = os.Mkdir(path + "/" + directoryName, os.ModePerm)

	return err
}
