package commands

import (
	"fmt"
	"os"
)

func (command Command) Move() error{
	if err := command.validateBinaryParameters(); err != nil{
		return err
	}
	return move(command.Params[0], command.Params[1])
}

func move(source string, destination string) error {
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
