package commands

import (
	"fmt"
	"os"
)

func Move(source string, destination string) error {
	sourcePath, err := getPath(source)

	if err != nil {
		return err
	}

	destinationPath, err := getPath(destination)

	if err != nil {
		return err
	}

	if err := os.Rename(sourcePath, destinationPath); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
